package main

import . "github.com/pyrex41/shen-go/kl"

var TStarMain = MakeNative(func(__e *ControlFlow) {
tmp14520 := MakeNative(func(__e *ControlFlow) {
V4915 := __e.Get(1)
_ = V4915
V4916 := __e.Get(2)
_ = V4916
tmp14521 := Call(__e, PrimFunc(symshen_4extract_1vars), V4916)


let__13969 := tmp14521
_ = let__13969

tmp14522 := Call(__e, PrimFunc(symshen_4rectify_1type), V4916)


let__13970 := tmp14522
_ = let__13970

tmp14523 := Call(__e, PrimFunc(symshen_4curry), V4915)


let__13971 := tmp14523
_ = let__13971

tmp14524 := MakeNative(func(__e *ControlFlow) {
Z4920 := __e.Get(1)
_ = Z4920
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4921 := __e.Get(1)
_ = Z4921
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4922 := __e.Get(1)
_ = Z4922
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4923 := __e.Get(1)
_ = Z4923
tmp14525 := Call(__e, PrimFunc(symshen_4newpv), Z4920)


let__13972 := tmp14525
_ = let__13972

tmp14526 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14526

tmp14527 := Call(__e, PrimFunc(symshen_4deref), let__13969, Z4920)


tmp14528 := Call(__e, PrimFunc(symreceive), tmp14527)


tmp14529 := Call(__e, PrimFunc(symshen_4deref), let__13970, Z4920)


tmp14530 := Call(__e, PrimFunc(symreceive), tmp14529)


tmp14531 := MakeNative(func(__e *ControlFlow) {
tmp14532 := Call(__e, PrimFunc(symshen_4deref), let__13971, Z4920)


tmp14533 := Call(__e, PrimFunc(symreceive), tmp14532)


tmp14534 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symreturn), let__13972, Z4920, Z4921, Z4922, Z4923)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4toplevel_1forms), tmp14533, let__13972, Z4920, Z4921, Z4922, tmp14534)
return


}, 0)

tmp14535 := Call(__e, PrimFunc(symshen_4insert_1prolog_1variables), tmp14528, tmp14530, let__13972, Z4920, Z4921, Z4922, tmp14531)


__e.TailApply(PrimFunc(symshen_4gc), Z4920, tmp14535)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp14536 := Call(__e, PrimFunc(symshen_4prolog_1vector))


tmp14537 := Call(__e, tmp14524, tmp14536)


tmp14538 := Call(__e, PrimFunc(symvector), MakeInteger(0))


tmp14539 := Call(__e, PrimFunc(sym_8v), MakeInteger(0), tmp14538)


tmp14540 := Call(__e, PrimFunc(sym_8v), True, tmp14539)


tmp14541 := Call(__e, tmp14537, tmp14540)


tmp14542 := Call(__e, tmp14541, MakeInteger(0))


tmp14543 := MakeNative(func(__e *ControlFlow) {
__e.Return(True)
return
}, 0)

__e.TailApply(tmp14542, tmp14543)
return


}, 2)

tmp14544 := Call(__e, ns2_1set, symshen_4typecheck, tmp14520)


_ = tmp14544

tmp14545 := MakeNative(func(__e *ControlFlow) {
V4925 := __e.Get(1)
_ = V4925
V4926 := __e.Get(2)
_ = V4926
V4927 := __e.Get(3)
_ = V4927
V4928 := __e.Get(4)
_ = V4928
V4929 := __e.Get(5)
_ = V4929
V4930 := __e.Get(6)
_ = V4930
V4931 := __e.Get(7)
_ = V4931
tmp14552 := Call(__e, PrimFunc(symshen_4unlocked_2), V4929)


var ifres14546 Obj

if True == tmp14552 {
tmp14547 := Call(__e, PrimFunc(symshen_4lazyderef), V4925, V4928)


let__13974 := tmp14547
_ = let__13974

tmp14551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13974, Nil)
}
__typedArg0 := let__13974
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14548 Obj

if True == tmp14551 {
tmp14549 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14549

tmp14550 := Call(__e, PrimFunc(symis_b), V4926, V4927, V4928, V4929, V4930, V4931)


ifres14548 = tmp14550


} else {
ifres14548 = False


}

ifres14546 = ifres14548


} else {
ifres14546 = False


}

let__13973 := ifres14546
_ = let__13973

tmp14566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13973, False)
}
__typedArg0 := let__13973
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14566 {
tmp14564 := Call(__e, PrimFunc(symshen_4unlocked_2), V4929)


if True == tmp14564 {
tmp14553 := Call(__e, PrimFunc(symshen_4lazyderef), V4925, V4928)


let__13975 := tmp14553
_ = let__13975

tmp14562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__13975)
}
__typedArg0 := let__13975
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14562 {
tmp14554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__13975)
}
__typedArg0 := let__13975
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__13976 := tmp14554
_ = let__13976

tmp14555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__13975)
}
__typedArg0 := let__13975
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__13977 := tmp14555
_ = let__13977

tmp14556 := Call(__e, PrimFunc(symshen_4newpv), V4928)


let__13978 := tmp14556
_ = let__13978

tmp14557 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14557

tmp14558 := Call(__e, PrimFunc(symshen_4deref), let__13978, V4928)


tmp14559 := Call(__e, PrimFunc(symsubst), tmp14558, let__13976, V4926)


tmp14560 := Call(__e, PrimFunc(symshen_4insert_1prolog_1variables), let__13977, tmp14559, V4927, V4928, V4929, V4930, V4931)


__e.TailApply(PrimFunc(symshen_4gc), V4928, tmp14560)
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
__e.Return(let__13973)
return
}


}, 7)

tmp14567 := Call(__e, ns2_1set, symshen_4insert_1prolog_1variables, tmp14545)


_ = tmp14567

tmp14568 := MakeNative(func(__e *ControlFlow) {
V4938 := __e.Get(1)
_ = V4938
V4939 := __e.Get(2)
_ = V4939
V4940 := __e.Get(3)
_ = V4940
V4941 := __e.Get(4)
_ = V4941
V4942 := __e.Get(5)
_ = V4942
V4943 := __e.Get(6)
_ = V4943
let__13979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4942)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4942
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__13979

tmp14593 := Call(__e, PrimFunc(symshen_4unlocked_2), V4941)


var ifres14570 Obj

if True == tmp14593 {
tmp14571 := Call(__e, PrimFunc(symshen_4lazyderef), V4938, V4940)


let__13981 := tmp14571
_ = let__13981

tmp14592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__13981)
}
__typedArg0 := let__13981
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14572 Obj

if True == tmp14592 {
tmp14573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__13981)
}
__typedArg0 := let__13981
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14574 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14573, V4940)


let__13982 := tmp14574
_ = let__13982

tmp14591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13982, symdefine)
}
__typedArg0 := let__13982
__typedArg1 := symdefine
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14575 Obj

if True == tmp14591 {
tmp14576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__13981)
}
__typedArg0 := let__13981
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14577 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14576, V4940)


let__13983 := tmp14577
_ = let__13983

tmp14590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__13983)
}
__typedArg0 := let__13983
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14578 Obj

if True == tmp14590 {
tmp14579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__13983)
}
__typedArg0 := let__13983
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__13984 := tmp14579
_ = let__13984

tmp14580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__13983)
}
__typedArg0 := let__13983
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__13985 := tmp14580
_ = let__13985

tmp14581 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14581

tmp14582 := Call(__e, PrimFunc(symshen_4type_1theory_1enabled_2))


tmp14583 := MakeNative(func(__e *ControlFlow) {
tmp14584 := MakeNative(func(__e *ControlFlow) {
tmp14585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspy_d)
}
__typedArg0 := symshen_4_dspy_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp14586 := MakeNative(func(__e *ControlFlow) {
tmp14587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__13984, let__13985)
}
__typedArg0 := let__13984
__typedArg1 := let__13985
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp14587)
}
__typedArg0 := symdefine
__typedArg1 := tmp14587
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4t_d), tmp14588, V4939, V4940, V4941, let__13979, V4943)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4signal_1def), tmp14585, let__13984, V4940, V4941, let__13979, tmp14586)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4940, V4941, let__13979, tmp14584)
return


}, 0)

tmp14589 := Call(__e, PrimFunc(symwhen), tmp14582, V4940, V4941, let__13979, tmp14583)


ifres14578 = tmp14589


} else {
ifres14578 = False


}

ifres14575 = ifres14578


} else {
ifres14575 = False


}

ifres14572 = ifres14575


} else {
ifres14572 = False


}

ifres14570 = ifres14572


} else {
ifres14570 = False


}

let__13980 := ifres14570
_ = let__13980

tmp14605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13980, False)
}
__typedArg0 := let__13980
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14605 {
tmp14601 := Call(__e, PrimFunc(symshen_4unlocked_2), V4941)


var ifres14594 Obj

if True == tmp14601 {
tmp14595 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14595

tmp14596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp14597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4939, Nil)
}
__typedArg0 := V4939
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14596, tmp14597)
}
__typedArg0 := tmp14596
__typedArg1 := tmp14597
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4938, tmp14598)
}
__typedArg0 := V4938
__typedArg1 := tmp14598
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14600 := Call(__e, PrimFunc(symshen_4system_1S), tmp14599, Nil, V4940, V4941, let__13979, V4943)


ifres14594 = tmp14600


} else {
ifres14594 = False


}

let__13986 := ifres14594
_ = let__13986

tmp14603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13986, False)
}
__typedArg0 := let__13986
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14603 {
__e.TailApply(PrimFunc(symshen_4unlock), V4941, let__13979)
return
} else {
__e.Return(let__13986)
return
}


} else {
__e.Return(let__13980)
return
}


}, 6)

tmp14606 := Call(__e, ns2_1set, symshen_4toplevel_1forms, tmp14568)


_ = tmp14606

tmp14607 := MakeNative(func(__e *ControlFlow) {
V4952 := __e.Get(1)
_ = V4952
V4953 := __e.Get(2)
_ = V4953
V4954 := __e.Get(3)
_ = V4954
V4955 := __e.Get(4)
_ = V4955
V4956 := __e.Get(5)
_ = V4956
V4957 := __e.Get(6)
_ = V4957
tmp14614 := Call(__e, PrimFunc(symshen_4unlocked_2), V4955)


var ifres14608 Obj

if True == tmp14614 {
tmp14609 := Call(__e, PrimFunc(symshen_4lazyderef), V4952, V4954)


let__13988 := tmp14609
_ = let__13988

tmp14613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13988, False)
}
__typedArg0 := let__13988
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14610 Obj

if True == tmp14613 {
tmp14611 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14611

tmp14612 := Call(__e, PrimFunc(symthaw), V4957)


ifres14610 = tmp14612


} else {
ifres14610 = False


}

ifres14608 = ifres14610


} else {
ifres14608 = False


}

let__13987 := ifres14608
_ = let__13987

tmp14629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13987, False)
}
__typedArg0 := let__13987
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14629 {
tmp14627 := Call(__e, PrimFunc(symshen_4unlocked_2), V4955)


if True == tmp14627 {
tmp14615 := Call(__e, PrimFunc(symshen_4lazyderef), V4952, V4954)


let__13989 := tmp14615
_ = let__13989

tmp14625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13989, True)
}
__typedArg0 := let__13989
__typedArg1 := True
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14625 {
tmp14616 := Call(__e, PrimFunc(symshen_4newpv), V4954)


let__13990 := tmp14616
_ = let__13990

tmp14617 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14617

tmp14618 := Call(__e, PrimFunc(symshen_4deref), V4953, V4954)


tmp14619 := Call(__e, PrimFunc(symshen_4app), tmp14618, MakeString(")\n"), symshen_4a)


tmp14620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\ntypechecking (fn "))
__typedS1, __typedOK1 := TypedString(tmp14619)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\ntypechecking (fn ")
__typedArg1 := tmp14619
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp14621 := Call(__e, PrimFunc(symstoutput))


tmp14622 := Call(__e, PrimFunc(sympr), tmp14620, tmp14621)


tmp14623 := Call(__e, PrimFunc(symis), let__13990, tmp14622, V4954, V4955, V4956, V4957)


__e.TailApply(PrimFunc(symshen_4gc), V4954, tmp14623)
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
__e.Return(let__13987)
return
}


}, 6)

tmp14630 := Call(__e, ns2_1set, symshen_4signal_1def, tmp14607)


_ = tmp14630

tmp14631 := MakeNative(func(__e *ControlFlow) {
V4962 := __e.Get(1)
_ = V4962
tmp14632 := Call(__e, PrimFunc(symshen_4curry_1type), V4962)


__e.TailApply(PrimFunc(symshen_4demodulate), tmp14632)
return


}, 1)

tmp14633 := Call(__e, ns2_1set, symshen_4rectify_1type, tmp14631)


_ = tmp14633

tmp14634 := MakeNative(func(__e *ControlFlow) {
V4963 := __e.Get(1)
_ = V4963
tmp14635 := MakeNative(func(__e *ControlFlow) {
tmp14636 := MakeNative(func(__e *ControlFlow) {
Z4965 := __e.Get(1)
_ = Z4965
__e.TailApply(PrimFunc(symshen_4demod), Z4965)
return
}, 1)

tmp14637 := Call(__e, PrimFunc(symshen_4walk), tmp14636, V4963)


let__13991 := tmp14637
_ = let__13991

tmp14639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13991, V4963)
}
__typedArg0 := let__13991
__typedArg1 := V4963
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14639 {
__e.Return(V4963)
return
} else {
__e.TailApply(PrimFunc(symshen_4demodulate), let__13991)
return
}


}, 0)

tmp14640 := MakeNative(func(__e *ControlFlow) {
Z4966 := __e.Get(1)
_ = Z4966
__e.Return(V4963)
return
}, 1)

__e.TailApply(try_1catch, tmp14635, tmp14640)
return


}, 1)

tmp14641 := Call(__e, ns2_1set, symshen_4demodulate, tmp14634)


_ = tmp14641

tmp14642 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V4967 := __self1
_ = V4967
tmp14766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14739 Obj

if True == tmp14766 {
tmp14764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14764)
}
__typedArg0 := tmp14764
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14741 Obj

if True == tmp14765 {
tmp14761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14761)
}
__typedArg0 := tmp14761
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1_1_6, tmp14762)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp14762
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14743 Obj

if True == tmp14763 {
tmp14758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14758)
}
__typedArg0 := tmp14758
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14759)
}
__typedArg0 := tmp14759
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14745 Obj

if True == tmp14760 {
tmp14754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14754)
}
__typedArg0 := tmp14754
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14755)
}
__typedArg0 := tmp14755
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14756)
}
__typedArg0 := tmp14756
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14747 Obj

if True == tmp14757 {
tmp14749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14749)
}
__typedArg0 := tmp14749
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14750)
}
__typedArg0 := tmp14750
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14751)
}
__typedArg0 := tmp14751
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1_1_6, tmp14752)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp14752
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14748 Obj

if True == tmp14753 {
ifres14748 = True


} else {
ifres14748 = False


}

ifres14747 = ifres14748


} else {
ifres14747 = False


}

var ifres14746 Obj

if True == ifres14747 {
ifres14746 = True


} else {
ifres14746 = False


}

ifres14745 = ifres14746


} else {
ifres14745 = False


}

var ifres14744 Obj

if True == ifres14745 {
ifres14744 = True


} else {
ifres14744 = False


}

ifres14743 = ifres14744


} else {
ifres14743 = False


}

var ifres14742 Obj

if True == ifres14743 {
ifres14742 = True


} else {
ifres14742 = False


}

ifres14741 = ifres14742


} else {
ifres14741 = False


}

var ifres14740 Obj

if True == ifres14741 {
ifres14740 = True


} else {
ifres14740 = False


}

ifres14739 = ifres14740


} else {
ifres14739 = False


}

if True == ifres14739 {
tmp14643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14644)
}
__typedArg0 := tmp14644
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14645, Nil)
}
__typedArg0 := tmp14645
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp14646)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp14646
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14643, tmp14647)
}
__typedArg0 := tmp14643
__typedArg1 := tmp14647
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4curry_1type) == __self {
__self1 = tmp14648
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4curry_1type), tmp14648)
return


} else {
tmp14737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14697 Obj

if True == tmp14737 {
tmp14735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14735)
}
__typedArg0 := tmp14735
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14699 Obj

if True == tmp14736 {
tmp14732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14732)
}
__typedArg0 := tmp14732
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp14733)
}
__typedArg0 := symlist
__typedArg1 := tmp14733
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14701 Obj

if True == tmp14734 {
tmp14729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14729)
}
__typedArg0 := tmp14729
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14730)
}
__typedArg0 := tmp14730
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14703 Obj

if True == tmp14731 {
tmp14725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14725)
}
__typedArg0 := tmp14725
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14726)
}
__typedArg0 := tmp14726
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14727)
}
__typedArg0 := Nil
__typedArg1 := tmp14727
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14705 Obj

if True == tmp14728 {
tmp14723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14723)
}
__typedArg0 := tmp14723
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14707 Obj

if True == tmp14724 {
tmp14720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14720)
}
__typedArg0 := tmp14720
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_a_a_6, tmp14721)
}
__typedArg0 := sym_a_a_6
__typedArg1 := tmp14721
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14709 Obj

if True == tmp14722 {
tmp14717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14717)
}
__typedArg0 := tmp14717
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14718)
}
__typedArg0 := tmp14718
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14711 Obj

if True == tmp14719 {
tmp14713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14713)
}
__typedArg0 := tmp14713
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14714)
}
__typedArg0 := tmp14714
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14715)
}
__typedArg0 := Nil
__typedArg1 := tmp14715
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14712 Obj

if True == tmp14716 {
ifres14712 = True


} else {
ifres14712 = False


}

ifres14711 = ifres14712


} else {
ifres14711 = False


}

var ifres14710 Obj

if True == ifres14711 {
ifres14710 = True


} else {
ifres14710 = False


}

ifres14709 = ifres14710


} else {
ifres14709 = False


}

var ifres14708 Obj

if True == ifres14709 {
ifres14708 = True


} else {
ifres14708 = False


}

ifres14707 = ifres14708


} else {
ifres14707 = False


}

var ifres14706 Obj

if True == ifres14707 {
ifres14706 = True


} else {
ifres14706 = False


}

ifres14705 = ifres14706


} else {
ifres14705 = False


}

var ifres14704 Obj

if True == ifres14705 {
ifres14704 = True


} else {
ifres14704 = False


}

ifres14703 = ifres14704


} else {
ifres14703 = False


}

var ifres14702 Obj

if True == ifres14703 {
ifres14702 = True


} else {
ifres14702 = False


}

ifres14701 = ifres14702


} else {
ifres14701 = False


}

var ifres14700 Obj

if True == ifres14701 {
ifres14700 = True


} else {
ifres14700 = False


}

ifres14699 = ifres14700


} else {
ifres14699 = False


}

var ifres14698 Obj

if True == ifres14699 {
ifres14698 = True


} else {
ifres14698 = False


}

ifres14697 = ifres14698


} else {
ifres14697 = False


}

if True == ifres14697 {
tmp14649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14651)
}
__typedArg0 := tmp14651
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14650, tmp14652)
}
__typedArg0 := tmp14650
__typedArg1 := tmp14652
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp14653)
}
__typedArg0 := symstr
__typedArg1 := tmp14653
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14654, Nil)
}
__typedArg0 := tmp14654
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp14655)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp14655
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14649, tmp14656)
}
__typedArg0 := tmp14649
__typedArg1 := tmp14656
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4curry_1type) == __self {
__self1 = tmp14657
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4curry_1type), tmp14657)
return


} else {
tmp14695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14668 Obj

if True == tmp14695 {
tmp14693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14693)
}
__typedArg0 := tmp14693
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14670 Obj

if True == tmp14694 {
tmp14690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14690)
}
__typedArg0 := tmp14690
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_d, tmp14691)
}
__typedArg0 := sym_d
__typedArg1 := tmp14691
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14672 Obj

if True == tmp14692 {
tmp14687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14687)
}
__typedArg0 := tmp14687
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14688)
}
__typedArg0 := tmp14688
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14674 Obj

if True == tmp14689 {
tmp14683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14683)
}
__typedArg0 := tmp14683
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14684)
}
__typedArg0 := tmp14684
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14685)
}
__typedArg0 := tmp14685
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14676 Obj

if True == tmp14686 {
tmp14678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14678)
}
__typedArg0 := tmp14678
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14679)
}
__typedArg0 := tmp14679
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14680)
}
__typedArg0 := tmp14680
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_d, tmp14681)
}
__typedArg0 := sym_d
__typedArg1 := tmp14681
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14677 Obj

if True == tmp14682 {
ifres14677 = True


} else {
ifres14677 = False


}

ifres14676 = ifres14677


} else {
ifres14676 = False


}

var ifres14675 Obj

if True == ifres14676 {
ifres14675 = True


} else {
ifres14675 = False


}

ifres14674 = ifres14675


} else {
ifres14674 = False


}

var ifres14673 Obj

if True == ifres14674 {
ifres14673 = True


} else {
ifres14673 = False


}

ifres14672 = ifres14673


} else {
ifres14672 = False


}

var ifres14671 Obj

if True == ifres14672 {
ifres14671 = True


} else {
ifres14671 = False


}

ifres14670 = ifres14671


} else {
ifres14670 = False


}

var ifres14669 Obj

if True == ifres14670 {
ifres14669 = True


} else {
ifres14669 = False


}

ifres14668 = ifres14669


} else {
ifres14668 = False


}

if True == ifres14668 {
tmp14658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14659)
}
__typedArg0 := tmp14659
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14660, Nil)
}
__typedArg0 := tmp14660
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp14661)
}
__typedArg0 := sym_d
__typedArg1 := tmp14661
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14658, tmp14662)
}
__typedArg0 := tmp14658
__typedArg1 := tmp14662
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4curry_1type) == __self {
__self1 = tmp14663
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4curry_1type), tmp14663)
return


} else {
tmp14666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4967)
}
__typedArg0 := V4967
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14666 {
tmp14664 := MakeNative(func(__e *ControlFlow) {
Z4968 := __e.Get(1)
_ = Z4968
__e.TailApply(PrimFunc(symshen_4curry_1type), Z4968)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp14664, V4967)
return


} else {
__e.Return(V4967)
return
}


}


}


}


}, 1)

tmp14767 := Call(__e, ns2_1set, symshen_4curry_1type, tmp14642)


_ = tmp14767

tmp14768 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V4969 := __self1
_ = V4969
tmp14857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14849 Obj

if True == tmp14857 {
tmp14855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp14855)
}
__typedArg0 := symdefine
__typedArg1 := tmp14855
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14851 Obj

if True == tmp14856 {
tmp14853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14853)
}
__typedArg0 := tmp14853
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14852 Obj

if True == tmp14854 {
ifres14852 = True


} else {
ifres14852 = False


}

ifres14851 = ifres14852


} else {
ifres14851 = False


}

var ifres14850 Obj

if True == ifres14851 {
ifres14850 = True


} else {
ifres14850 = False


}

ifres14849 = ifres14850


} else {
ifres14849 = False


}

if True == ifres14849 {
__e.Return(V4969)
return
} else {
tmp14847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14828 Obj

if True == tmp14847 {
tmp14845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtype, tmp14845)
}
__typedArg0 := symtype
__typedArg1 := tmp14845
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14830 Obj

if True == tmp14846 {
tmp14843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14843)
}
__typedArg0 := tmp14843
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14832 Obj

if True == tmp14844 {
tmp14840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14840)
}
__typedArg0 := tmp14840
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14841)
}
__typedArg0 := tmp14841
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14834 Obj

if True == tmp14842 {
tmp14836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14836)
}
__typedArg0 := tmp14836
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14837)
}
__typedArg0 := tmp14837
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14838)
}
__typedArg0 := Nil
__typedArg1 := tmp14838
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14835 Obj

if True == tmp14839 {
ifres14835 = True


} else {
ifres14835 = False


}

ifres14834 = ifres14835


} else {
ifres14834 = False


}

var ifres14833 Obj

if True == ifres14834 {
ifres14833 = True


} else {
ifres14833 = False


}

ifres14832 = ifres14833


} else {
ifres14832 = False


}

var ifres14831 Obj

if True == ifres14832 {
ifres14831 = True


} else {
ifres14831 = False


}

ifres14830 = ifres14831


} else {
ifres14830 = False


}

var ifres14829 Obj

if True == ifres14830 {
ifres14829 = True


} else {
ifres14829 = False


}

ifres14828 = ifres14829


} else {
ifres14828 = False


}

if True == ifres14828 {
tmp14769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14769)
}
__typedArg0 := tmp14769
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14771 := Call(__e, PrimFunc(symshen_4curry), tmp14770)


tmp14772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14772)
}
__typedArg0 := tmp14772
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14771, tmp14773)
}
__typedArg0 := tmp14771
__typedArg1 := tmp14773
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp14774)
}
__typedArg0 := symtype
__typedArg1 := tmp14774
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp14826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14822 Obj

if True == tmp14826 {
tmp14824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14825 := Call(__e, PrimFunc(symshen_4special_2), tmp14824)


var ifres14823 Obj

if True == tmp14825 {
ifres14823 = True


} else {
ifres14823 = False


}

ifres14822 = ifres14823


} else {
ifres14822 = False


}

if True == ifres14822 {
tmp14775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14776 := MakeNative(func(__e *ControlFlow) {
Z4970 := __e.Get(1)
_ = Z4970
__e.TailApply(PrimFunc(symshen_4curry), Z4970)
return
}, 1)

tmp14777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14778 := Call(__e, PrimFunc(symmap), tmp14776, tmp14777)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14775, tmp14778)
}
__typedArg0 := tmp14775
__typedArg1 := tmp14778
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp14820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14816 Obj

if True == tmp14820 {
tmp14818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14819 := Call(__e, PrimFunc(symshen_4extraspecial_2), tmp14818)


var ifres14817 Obj

if True == tmp14819 {
ifres14817 = True


} else {
ifres14817 = False


}

ifres14816 = ifres14817


} else {
ifres14816 = False


}

if True == ifres14816 {
__e.Return(V4969)
return
} else {
tmp14814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14805 Obj

if True == tmp14814 {
tmp14812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14812)
}
__typedArg0 := tmp14812
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14807 Obj

if True == tmp14813 {
tmp14809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14809)
}
__typedArg0 := tmp14809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14810)
}
__typedArg0 := tmp14810
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14808 Obj

if True == tmp14811 {
ifres14808 = True


} else {
ifres14808 = False


}

ifres14807 = ifres14808


} else {
ifres14807 = False


}

var ifres14806 Obj

if True == ifres14807 {
ifres14806 = True


} else {
ifres14806 = False


}

ifres14805 = ifres14806


} else {
ifres14805 = False


}

if True == ifres14805 {
tmp14779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14780)
}
__typedArg0 := tmp14780
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14781, Nil)
}
__typedArg0 := tmp14781
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14779, tmp14782)
}
__typedArg0 := tmp14779
__typedArg1 := tmp14782
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14784)
}
__typedArg0 := tmp14784
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14783, tmp14785)
}
__typedArg0 := tmp14783
__typedArg1 := tmp14785
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4curry) == __self {
__self1 = tmp14786
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4curry), tmp14786)
return


} else {
tmp14803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14794 Obj

if True == tmp14803 {
tmp14801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14801)
}
__typedArg0 := tmp14801
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14796 Obj

if True == tmp14802 {
tmp14798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14798)
}
__typedArg0 := tmp14798
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14799)
}
__typedArg0 := Nil
__typedArg1 := tmp14799
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14797 Obj

if True == tmp14800 {
ifres14797 = True


} else {
ifres14797 = False


}

ifres14796 = ifres14797


} else {
ifres14796 = False


}

var ifres14795 Obj

if True == ifres14796 {
ifres14795 = True


} else {
ifres14795 = False


}

ifres14794 = ifres14795


} else {
ifres14794 = False


}

if True == ifres14794 {
tmp14787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14788 := Call(__e, PrimFunc(symshen_4curry), tmp14787)


tmp14789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4969)
}
__typedArg0 := V4969
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14789)
}
__typedArg0 := tmp14789
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14791 := Call(__e, PrimFunc(symshen_4curry), tmp14790)


tmp14792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14791, Nil)
}
__typedArg0 := tmp14791
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14788, tmp14792)
}
__typedArg0 := tmp14788
__typedArg1 := tmp14792
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V4969)
return
}


}


}


}


}


}


}, 1)

tmp14858 := Call(__e, ns2_1set, symshen_4curry, tmp14768)


_ = tmp14858

tmp14859 := MakeNative(func(__e *ControlFlow) {
V4971 := __e.Get(1)
_ = V4971
tmp14860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspecial_d)
}
__typedArg0 := symshen_4_dspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symelement_2), V4971, tmp14860)
return


}, 1)

tmp14861 := Call(__e, ns2_1set, symshen_4special_2, tmp14859)


_ = tmp14861

tmp14862 := MakeNative(func(__e *ControlFlow) {
V4972 := __e.Get(1)
_ = V4972
tmp14863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dextraspecial_d)
}
__typedArg0 := symshen_4_dextraspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symelement_2), V4972, tmp14863)
return


}, 1)

tmp14864 := Call(__e, ns2_1set, symshen_4extraspecial_2, tmp14862)


_ = tmp14864

tmp14865 := MakeNative(func(__e *ControlFlow) {
V4973 := __e.Get(1)
_ = V4973
V4974 := __e.Get(2)
_ = V4974
V4975 := __e.Get(3)
_ = V4975
V4976 := __e.Get(4)
_ = V4976
V4977 := __e.Get(5)
_ = V4977
V4978 := __e.Get(6)
_ = V4978
let__13992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4977)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4977
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__13992

tmp14871 := Call(__e, PrimFunc(symshen_4unlocked_2), V4976)


var ifres14867 Obj

if True == tmp14871 {
tmp14868 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14868

tmp14869 := Call(__e, PrimFunc(symshen_4maxinfexceeded_2))


tmp14870 := Call(__e, PrimFunc(symwhen), tmp14869, V4975, V4976, let__13992, V4978)


ifres14867 = tmp14870


} else {
ifres14867 = False


}

let__13993 := ifres14867
_ = let__13993

tmp14919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13993, False)
}
__typedArg0 := let__13993
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14919 {
tmp14900 := Call(__e, PrimFunc(symshen_4unlocked_2), V4976)


var ifres14872 Obj

if True == tmp14900 {
tmp14873 := Call(__e, PrimFunc(symshen_4lazyderef), V4973, V4975)


let__13995 := tmp14873
_ = let__13995

tmp14899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__13995)
}
__typedArg0 := let__13995
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14874 Obj

if True == tmp14899 {
tmp14875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__13995)
}
__typedArg0 := let__13995
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__13996 := tmp14875
_ = let__13996

tmp14876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__13995)
}
__typedArg0 := let__13995
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14877 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14876, V4975)


let__13997 := tmp14877
_ = let__13997

tmp14898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__13997)
}
__typedArg0 := let__13997
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14878 Obj

if True == tmp14898 {
tmp14879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__13997)
}
__typedArg0 := let__13997
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__13998 := tmp14879
_ = let__13998

tmp14880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__13997)
}
__typedArg0 := let__13997
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14881 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14880, V4975)


let__13999 := tmp14881
_ = let__13999

tmp14897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__13999)
}
__typedArg0 := let__13999
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14882 Obj

if True == tmp14897 {
tmp14883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__13999)
}
__typedArg0 := let__13999
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14000 := tmp14883
_ = let__14000

tmp14884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__13999)
}
__typedArg0 := let__13999
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14885 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14884, V4975)


let__14001 := tmp14885
_ = let__14001

tmp14896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14001, Nil)
}
__typedArg0 := let__14001
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14886 Obj

if True == tmp14896 {
tmp14887 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14887

tmp14888 := Call(__e, PrimFunc(symshen_4deref), let__13998, V4975)


tmp14889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp14890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp14888, tmp14889)
}
__typedArg0 := tmp14888
__typedArg1 := tmp14889
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp14891 := MakeNative(func(__e *ControlFlow) {
tmp14892 := Call(__e, PrimFunc(symshen_4type_1theory_1enabled_2))


tmp14893 := MakeNative(func(__e *ControlFlow) {
tmp14894 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__13996, let__14000, V4974, V4975, V4976, let__13992, V4978)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4975, V4976, let__13992, tmp14894)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14892, V4975, V4976, let__13992, tmp14893)
return


}, 0)

tmp14895 := Call(__e, PrimFunc(symwhen), tmp14890, V4975, V4976, let__13992, tmp14891)


ifres14886 = tmp14895


} else {
ifres14886 = False


}

ifres14882 = ifres14886


} else {
ifres14882 = False


}

ifres14878 = ifres14882


} else {
ifres14878 = False


}

ifres14874 = ifres14878


} else {
ifres14874 = False


}

ifres14872 = ifres14874


} else {
ifres14872 = False


}

let__13994 := ifres14872
_ = let__13994

tmp14917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13994, False)
}
__typedArg0 := let__13994
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14917 {
tmp14906 := Call(__e, PrimFunc(symshen_4unlocked_2), V4976)


var ifres14901 Obj

if True == tmp14906 {
tmp14902 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14902

tmp14903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspy_d)
}
__typedArg0 := symshen_4_dspy_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp14904 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4show), V4973, V4974, V4975, V4976, let__13992, V4978)
return
}, 0)

tmp14905 := Call(__e, PrimFunc(symwhen), tmp14903, V4975, V4976, let__13992, tmp14904)


ifres14901 = tmp14905


} else {
ifres14901 = False


}

let__14002 := ifres14901
_ = let__14002

tmp14915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14002, False)
}
__typedArg0 := let__14002
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14915 {
tmp14911 := Call(__e, PrimFunc(symshen_4unlocked_2), V4976)


var ifres14907 Obj

if True == tmp14911 {
tmp14908 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14908

tmp14909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp14910 := Call(__e, PrimFunc(symshen_4search_1user_1datatypes), V4973, V4974, tmp14909, V4975, V4976, let__13992, V4978)


ifres14907 = tmp14910


} else {
ifres14907 = False


}

let__14003 := ifres14907
_ = let__14003

tmp14913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14003, False)
}
__typedArg0 := let__14003
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14913 {
__e.TailApply(PrimFunc(symshen_4unlock), V4976, let__13992)
return
} else {
__e.Return(let__14003)
return
}


} else {
__e.Return(let__14002)
return
}


} else {
__e.Return(let__13994)
return
}


} else {
__e.Return(let__13993)
return
}


}, 6)

tmp14920 := Call(__e, ns2_1set, symshen_4system_1S, tmp14865)


_ = tmp14920

tmp14921 := MakeNative(func(__e *ControlFlow) {
V4997 := __e.Get(1)
_ = V4997
V4998 := __e.Get(2)
_ = V4998
V4999 := __e.Get(3)
_ = V4999
V5000 := __e.Get(4)
_ = V5000
V5001 := __e.Get(5)
_ = V5001
V5002 := __e.Get(6)
_ = V5002
tmp14922 := Call(__e, PrimFunc(symshen_4line))


_ = tmp14922

tmp14923 := Call(__e, PrimFunc(symshen_4deref), V4997, V4999)


tmp14924 := Call(__e, PrimFunc(symshen_4show_1p), tmp14923)


_ = tmp14924

tmp14925 := Call(__e, PrimFunc(symnl), MakeInteger(2))


_ = tmp14925

tmp14926 := Call(__e, PrimFunc(symshen_4deref), V4998, V4999)


tmp14927 := Call(__e, PrimFunc(symshen_4show_1assumptions), tmp14926, MakeInteger(1))


_ = tmp14927

tmp14928 := Call(__e, PrimFunc(symshen_4pause_1for_1user))


_ = tmp14928

__e.Return(False)
return


}, 6)

tmp14929 := Call(__e, ns2_1set, symshen_4show, tmp14921)


_ = tmp14929

tmp14930 := MakeNative(func(__e *ControlFlow) {
tmp14931 := Call(__e, PrimFunc(syminferences))


let__14004 := tmp14931
_ = let__14004

tmp14933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(1), let__14004)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := let__14004
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14932 Obj

if True == tmp14933 {
ifres14932 = MakeString("")


} else {
ifres14932 = MakeString("s")


}

tmp14934 := Call(__e, PrimFunc(symshen_4app), ifres14932, MakeString(" \n?- "), symshen_4a)


tmp14936 := Call(__e, PrimFunc(symshen_4app), let__14004, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" inference"))
__typedS1, __typedOK1 := TypedString(tmp14934)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" inference")
__typedArg1 := tmp14934
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp14937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("____________________________________________________________ "))
__typedS1, __typedOK1 := TypedString(tmp14936)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("____________________________________________________________ ")
__typedArg1 := tmp14936
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp14938 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp14937, tmp14938)
return


}, 0)

tmp14939 := Call(__e, ns2_1set, symshen_4line, tmp14930)


_ = tmp14939

tmp14940 := MakeNative(func(__e *ControlFlow) {
V5004 := __e.Get(1)
_ = V5004
tmp14972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5004)
}
__typedArg0 := V5004
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14951 Obj

if True == tmp14972 {
tmp14970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5004)
}
__typedArg0 := V5004
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14970)
}
__typedArg0 := tmp14970
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14953 Obj

if True == tmp14971 {
tmp14967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5004)
}
__typedArg0 := V5004
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14967)
}
__typedArg0 := tmp14967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14968)
}
__typedArg0 := tmp14968
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14955 Obj

if True == tmp14969 {
tmp14963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5004)
}
__typedArg0 := V5004
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14963)
}
__typedArg0 := tmp14963
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14964)
}
__typedArg0 := tmp14964
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14965)
}
__typedArg0 := Nil
__typedArg1 := tmp14965
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14957 Obj

if True == tmp14966 {
tmp14959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5004)
}
__typedArg0 := V5004
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14959)
}
__typedArg0 := tmp14959
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp14962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp14960, tmp14961)
}
__typedArg0 := tmp14960
__typedArg1 := tmp14961
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14958 Obj

if True == tmp14962 {
ifres14958 = True


} else {
ifres14958 = False


}

ifres14957 = ifres14958


} else {
ifres14957 = False


}

var ifres14956 Obj

if True == ifres14957 {
ifres14956 = True


} else {
ifres14956 = False


}

ifres14955 = ifres14956


} else {
ifres14955 = False


}

var ifres14954 Obj

if True == ifres14955 {
ifres14954 = True


} else {
ifres14954 = False


}

ifres14953 = ifres14954


} else {
ifres14953 = False


}

var ifres14952 Obj

if True == ifres14953 {
ifres14952 = True


} else {
ifres14952 = False


}

ifres14951 = ifres14952


} else {
ifres14951 = False


}

if True == ifres14951 {
tmp14941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5004)
}
__typedArg0 := V5004
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14942 := Call(__e, PrimFunc(symshen_4prterm), tmp14941)


_ = tmp14942

tmp14943 := Call(__e, PrimFunc(symstoutput))


tmp14944 := Call(__e, PrimFunc(sympr), MakeString(" : "), tmp14943)


_ = tmp14944

tmp14945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5004)
}
__typedArg0 := V5004
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14945)
}
__typedArg0 := tmp14945
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14946)
}
__typedArg0 := tmp14946
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14948 := Call(__e, PrimFunc(symshen_4app), tmp14947, MakeString(""), symshen_4r)


tmp14949 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp14948, tmp14949)
return


} else {
__e.TailApply(PrimFunc(symshen_4prterm), V5004)
return
}


}, 1)

tmp14973 := Call(__e, ns2_1set, symshen_4show_1p, tmp14940)


_ = tmp14973

tmp14974 := MakeNative(func(__e *ControlFlow) {
V5005 := __e.Get(1)
_ = V5005
tmp15017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5005)
}
__typedArg0 := V5005
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14998 Obj

if True == tmp15017 {
tmp15015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5005)
}
__typedArg0 := V5005
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp15015)
}
__typedArg0 := symcons
__typedArg1 := tmp15015
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15000 Obj

if True == tmp15016 {
tmp15013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5005)
}
__typedArg0 := V5005
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15013)
}
__typedArg0 := tmp15013
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15002 Obj

if True == tmp15014 {
tmp15010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5005)
}
__typedArg0 := V5005
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15010)
}
__typedArg0 := tmp15010
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15011)
}
__typedArg0 := tmp15011
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15004 Obj

if True == tmp15012 {
tmp15006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5005)
}
__typedArg0 := V5005
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15006)
}
__typedArg0 := tmp15006
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15007)
}
__typedArg0 := tmp15007
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp15008)
}
__typedArg0 := Nil
__typedArg1 := tmp15008
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15005 Obj

if True == tmp15009 {
ifres15005 = True


} else {
ifres15005 = False


}

ifres15004 = ifres15005


} else {
ifres15004 = False


}

var ifres15003 Obj

if True == ifres15004 {
ifres15003 = True


} else {
ifres15003 = False


}

ifres15002 = ifres15003


} else {
ifres15002 = False


}

var ifres15001 Obj

if True == ifres15002 {
ifres15001 = True


} else {
ifres15001 = False


}

ifres15000 = ifres15001


} else {
ifres15000 = False


}

var ifres14999 Obj

if True == ifres15000 {
ifres14999 = True


} else {
ifres14999 = False


}

ifres14998 = ifres14999


} else {
ifres14998 = False


}

if True == ifres14998 {
tmp14975 := Call(__e, PrimFunc(symstoutput))


tmp14976 := Call(__e, PrimFunc(sympr), MakeString("["), tmp14975)


_ = tmp14976

tmp14977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5005)
}
__typedArg0 := V5005
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14977)
}
__typedArg0 := tmp14977
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14979 := Call(__e, PrimFunc(symshen_4prterm), tmp14978)


_ = tmp14979

tmp14980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5005)
}
__typedArg0 := V5005
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14980)
}
__typedArg0 := tmp14980
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14981)
}
__typedArg0 := tmp14981
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14983 := Call(__e, PrimFunc(symshen_4prtl), tmp14982)


_ = tmp14983

tmp14984 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString("]"), tmp14984)
return


} else {
tmp14996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5005)
}
__typedArg0 := V5005
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14996 {
tmp14985 := Call(__e, PrimFunc(symstoutput))


tmp14986 := Call(__e, PrimFunc(sympr), MakeString("("), tmp14985)


_ = tmp14986

tmp14987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5005)
}
__typedArg0 := V5005
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14988 := Call(__e, PrimFunc(symshen_4prterm), tmp14987)


_ = tmp14988

tmp14989 := MakeNative(func(__e *ControlFlow) {
Z5006 := __e.Get(1)
_ = Z5006
tmp14990 := Call(__e, PrimFunc(symstoutput))


tmp14991 := Call(__e, PrimFunc(sympr), MakeString(" "), tmp14990)


_ = tmp14991

__e.TailApply(PrimFunc(symshen_4prterm), Z5006)
return


}, 1)

tmp14992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5005)
}
__typedArg0 := V5005
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14993 := Call(__e, PrimFunc(symmap), tmp14989, tmp14992)


_ = tmp14993

tmp14994 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString(")"), tmp14994)
return


} else {
__e.TailApply(PrimFunc(symprint), V5005)
return
}


}


}, 1)

tmp15018 := Call(__e, ns2_1set, symshen_4prterm, tmp14974)


_ = tmp15018

tmp15019 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V5007 := __self1
_ = V5007
tmp15052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5007)
}
__typedArg0 := Nil
__typedArg1 := V5007
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15052 {
__e.Return(MakeString(""))
return
} else {
tmp15050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5007)
}
__typedArg0 := V5007
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15031 Obj

if True == tmp15050 {
tmp15048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5007)
}
__typedArg0 := V5007
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp15048)
}
__typedArg0 := symcons
__typedArg1 := tmp15048
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15033 Obj

if True == tmp15049 {
tmp15046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5007)
}
__typedArg0 := V5007
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15046)
}
__typedArg0 := tmp15046
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15035 Obj

if True == tmp15047 {
tmp15043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5007)
}
__typedArg0 := V5007
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15043)
}
__typedArg0 := tmp15043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15044)
}
__typedArg0 := tmp15044
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15037 Obj

if True == tmp15045 {
tmp15039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5007)
}
__typedArg0 := V5007
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15039)
}
__typedArg0 := tmp15039
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15040)
}
__typedArg0 := tmp15040
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp15041)
}
__typedArg0 := Nil
__typedArg1 := tmp15041
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15038 Obj

if True == tmp15042 {
ifres15038 = True


} else {
ifres15038 = False


}

ifres15037 = ifres15038


} else {
ifres15037 = False


}

var ifres15036 Obj

if True == ifres15037 {
ifres15036 = True


} else {
ifres15036 = False


}

ifres15035 = ifres15036


} else {
ifres15035 = False


}

var ifres15034 Obj

if True == ifres15035 {
ifres15034 = True


} else {
ifres15034 = False


}

ifres15033 = ifres15034


} else {
ifres15033 = False


}

var ifres15032 Obj

if True == ifres15033 {
ifres15032 = True


} else {
ifres15032 = False


}

ifres15031 = ifres15032


} else {
ifres15031 = False


}

if True == ifres15031 {
tmp15020 := Call(__e, PrimFunc(symstoutput))


tmp15021 := Call(__e, PrimFunc(sympr), MakeString(" "), tmp15020)


_ = tmp15021

tmp15022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5007)
}
__typedArg0 := V5007
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15022)
}
__typedArg0 := tmp15022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15024 := Call(__e, PrimFunc(symshen_4prterm), tmp15023)


_ = tmp15024

tmp15025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5007)
}
__typedArg0 := V5007
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15025)
}
__typedArg0 := tmp15025
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15026)
}
__typedArg0 := tmp15026
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

if PrimFunc(symshen_4prtl) == __self {
__self1 = tmp15027
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4prtl), tmp15027)
return


} else {
tmp15028 := Call(__e, PrimFunc(symstoutput))


tmp15029 := Call(__e, PrimFunc(sympr), MakeString(" | "), tmp15028)


_ = tmp15029

__e.TailApply(PrimFunc(symshen_4prterm), V5007)
return


}


}


}, 1)

tmp15053 := Call(__e, ns2_1set, symshen_4prtl, tmp15019)


_ = tmp15053

tmp15054 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V5014 := __self1
_ = V5014
V5015 := __self2
_ = V5015
tmp15067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5014)
}
__typedArg0 := Nil
__typedArg1 := V5014
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15067 {
tmp15055 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString("\n> "), tmp15055)
return


} else {
tmp15065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5014)
}
__typedArg0 := V5014
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15065 {
tmp15056 := Call(__e, PrimFunc(symshen_4app), V5015, MakeString(". "), symshen_4a)


tmp15057 := Call(__e, PrimFunc(symstoutput))


tmp15058 := Call(__e, PrimFunc(sympr), tmp15056, tmp15057)


_ = tmp15058

tmp15059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5014)
}
__typedArg0 := V5014
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15060 := Call(__e, PrimFunc(symshen_4show_1p), tmp15059)


_ = tmp15060

tmp15061 := Call(__e, PrimFunc(symnl), MakeInteger(1))


_ = tmp15061

tmp15062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5014)
}
__typedArg0 := V5014
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4show_1assumptions) == __self {
__self1, __self2 = tmp15062, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5015)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5015
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4show_1assumptions), tmp15062, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5015)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5015
__typedArg1 := MakeInteger(1)
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

tmp15068 := Call(__e, ns2_1set, symshen_4show_1assumptions, tmp15054)


_ = tmp15068

tmp15069 := MakeNative(func(__e *ControlFlow) {
tmp15070 := Call(__e, PrimFunc(symstinput))


tmp15071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(tmp15070)
}
__typedArg0 := tmp15070
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

let__14005 := tmp15071
_ = let__14005

tmp15073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14005, MakeInteger(94))
}
__typedArg0 := let__14005
__typedArg1 := MakeInteger(94)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15073 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("input aborted\n"))
}
__typedArg0 := MakeString("input aborted\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
__e.TailApply(PrimFunc(symnl), MakeInteger(1))
return
}


}, 0)

tmp15074 := Call(__e, ns2_1set, symshen_4pause_1for_1user, tmp15069)


_ = tmp15074

tmp15075 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dshen_1type_1theory_1enabled_2_d)
}
__typedArg0 := symshen_4_dshen_1type_1theory_1enabled_2_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp15076 := Call(__e, ns2_1set, symshen_4type_1theory_1enabled_2, tmp15075)


_ = tmp15076

tmp15077 := MakeNative(func(__e *ControlFlow) {
tmp15079 := Call(__e, PrimFunc(syminferences))


tmp15080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dmaxinferences_d)
}
__typedArg0 := symshen_4_dmaxinferences_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(tmp15079)
__typedN1, __typedOK1 := TypedFloat64(tmp15080)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := tmp15079
__typedArg1 := tmp15080
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

tmp15082 := Call(__e, ns2_1set, symshen_4maxinfexceeded_2, tmp15077)


_ = tmp15082

tmp15083 := MakeNative(func(__e *ControlFlow) {
V5017 := __e.Get(1)
_ = V5017
V5018 := __e.Get(2)
_ = V5018
V5019 := __e.Get(3)
_ = V5019
V5020 := __e.Get(4)
_ = V5020
V5021 := __e.Get(5)
_ = V5021
V5022 := __e.Get(6)
_ = V5022
V5023 := __e.Get(7)
_ = V5023
let__14006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5022)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5022
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__14006

tmp15094 := Call(__e, PrimFunc(symshen_4unlocked_2), V5021)


var ifres15085 Obj

if True == tmp15094 {
tmp15086 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15086

tmp15087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspy_d)
}
__typedArg0 := symshen_4_dspy_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp15088 := MakeNative(func(__e *ControlFlow) {
tmp15089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp15090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5018, Nil)
}
__typedArg0 := V5018
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15089, tmp15090)
}
__typedArg0 := tmp15089
__typedArg1 := tmp15090
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5017, tmp15091)
}
__typedArg0 := V5017
__typedArg1 := tmp15091
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4show), tmp15092, V5019, V5020, V5021, let__14006, V5023)
return


}, 0)

tmp15093 := Call(__e, PrimFunc(symwhen), tmp15087, V5020, V5021, let__14006, tmp15088)


ifres15085 = tmp15093


} else {
ifres15085 = False


}

let__14007 := ifres15085
_ = let__14007

tmp15833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14007, False)
}
__typedArg0 := let__14007
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15833 {
tmp15102 := Call(__e, PrimFunc(symshen_4unlocked_2), V5021)


var ifres15095 Obj

if True == tmp15102 {
tmp15096 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15096

tmp15097 := Call(__e, PrimFunc(symshen_4lazyderef), V5017, V5020)


tmp15098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15097)
}
__typedArg0 := tmp15097
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

tmp15099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp15098)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp15098
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

tmp15100 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4primitive), V5017, V5018, V5020, V5021, let__14006, V5023)
return
}, 0)

tmp15101 := Call(__e, PrimFunc(symwhen), tmp15099, V5020, V5021, let__14006, tmp15100)


ifres15095 = tmp15101


} else {
ifres15095 = False


}

let__14008 := ifres15095
_ = let__14008

tmp15831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14008, False)
}
__typedArg0 := let__14008
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15831 {
tmp15106 := Call(__e, PrimFunc(symshen_4unlocked_2), V5021)


var ifres15103 Obj

if True == tmp15106 {
tmp15104 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15104

tmp15105 := Call(__e, PrimFunc(symshen_4by_1hypothesis), V5017, V5018, V5019, V5020, V5021, let__14006, V5023)


ifres15103 = tmp15105


} else {
ifres15103 = False


}

let__14009 := ifres15103
_ = let__14009

tmp15829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14009, False)
}
__typedArg0 := let__14009
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15829 {
tmp15120 := Call(__e, PrimFunc(symshen_4unlocked_2), V5021)


var ifres15107 Obj

if True == tmp15120 {
tmp15108 := Call(__e, PrimFunc(symshen_4lazyderef), V5017, V5020)


let__14011 := tmp15108
_ = let__14011

tmp15119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14011)
}
__typedArg0 := let__14011
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15109 Obj

if True == tmp15119 {
tmp15110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14011)
}
__typedArg0 := let__14011
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14012 := tmp15110
_ = let__14012

tmp15111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14011)
}
__typedArg0 := let__14011
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15112 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15111, V5020)


let__14013 := tmp15112
_ = let__14013

tmp15118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14013, Nil)
}
__typedArg0 := let__14013
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15113 Obj

if True == tmp15118 {
tmp15114 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15114

tmp15115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5018, Nil)
}
__typedArg0 := V5018
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15115)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15115
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15117 := Call(__e, PrimFunc(symshen_4lookupsig), let__14012, tmp15116, V5020, V5021, let__14006, V5023)


ifres15113 = tmp15117


} else {
ifres15113 = False


}

ifres15109 = ifres15113


} else {
ifres15109 = False


}

ifres15107 = ifres15109


} else {
ifres15107 = False


}

let__14010 := ifres15107
_ = let__14010

tmp15827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14010, False)
}
__typedArg0 := let__14010
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15827 {
tmp15146 := Call(__e, PrimFunc(symshen_4unlocked_2), V5021)


var ifres15121 Obj

if True == tmp15146 {
tmp15122 := Call(__e, PrimFunc(symshen_4lazyderef), V5017, V5020)


let__14015 := tmp15122
_ = let__14015

tmp15145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14015)
}
__typedArg0 := let__14015
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15123 Obj

if True == tmp15145 {
tmp15124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14015)
}
__typedArg0 := let__14015
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15125 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15124, V5020)


let__14016 := tmp15125
_ = let__14016

tmp15144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14016, symfn)
}
__typedArg0 := let__14016
__typedArg1 := symfn
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15126 Obj

if True == tmp15144 {
tmp15127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14015)
}
__typedArg0 := let__14015
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15128 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15127, V5020)


let__14017 := tmp15128
_ = let__14017

tmp15143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14017)
}
__typedArg0 := let__14017
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15129 Obj

if True == tmp15143 {
tmp15130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14017)
}
__typedArg0 := let__14017
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14018 := tmp15130
_ = let__14018

tmp15131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14017)
}
__typedArg0 := let__14017
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15132 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15131, V5020)


let__14019 := tmp15132
_ = let__14019

tmp15142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14019, Nil)
}
__typedArg0 := let__14019
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15133 Obj

if True == tmp15142 {
tmp15134 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15134

tmp15135 := Call(__e, PrimFunc(symshen_4deref), let__14018, V5020)


tmp15136 := Call(__e, PrimFunc(symarity), tmp15135)


tmp15137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp15136, MakeInteger(0))
}
__typedArg0 := tmp15136
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp15138 := MakeNative(func(__e *ControlFlow) {
tmp15139 := MakeNative(func(__e *ControlFlow) {
tmp15140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14018, Nil)
}
__typedArg0 := let__14018
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp15140, V5018, V5019, V5020, V5021, let__14006, V5023)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5020, V5021, let__14006, tmp15139)
return


}, 0)

tmp15141 := Call(__e, PrimFunc(symwhen), tmp15137, V5020, V5021, let__14006, tmp15138)


ifres15133 = tmp15141


} else {
ifres15133 = False


}

ifres15129 = ifres15133


} else {
ifres15129 = False


}

ifres15126 = ifres15129


} else {
ifres15126 = False


}

ifres15123 = ifres15126


} else {
ifres15123 = False


}

ifres15121 = ifres15123


} else {
ifres15121 = False


}

let__14014 := ifres15121
_ = let__14014

tmp15825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14014, False)
}
__typedArg0 := let__14014
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15825 {
tmp15166 := Call(__e, PrimFunc(symshen_4unlocked_2), V5021)


var ifres15147 Obj

if True == tmp15166 {
tmp15148 := Call(__e, PrimFunc(symshen_4lazyderef), V5017, V5020)


let__14021 := tmp15148
_ = let__14021

tmp15165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14021)
}
__typedArg0 := let__14021
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15149 Obj

if True == tmp15165 {
tmp15150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14021)
}
__typedArg0 := let__14021
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15151 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15150, V5020)


let__14022 := tmp15151
_ = let__14022

tmp15164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14022, symfn)
}
__typedArg0 := let__14022
__typedArg1 := symfn
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15152 Obj

if True == tmp15164 {
tmp15153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14021)
}
__typedArg0 := let__14021
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15154 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15153, V5020)


let__14023 := tmp15154
_ = let__14023

tmp15163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14023)
}
__typedArg0 := let__14023
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15155 Obj

if True == tmp15163 {
tmp15156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14023)
}
__typedArg0 := let__14023
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14024 := tmp15156
_ = let__14024

tmp15157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14023)
}
__typedArg0 := let__14023
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15158 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15157, V5020)


let__14025 := tmp15158
_ = let__14025

tmp15162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14025, Nil)
}
__typedArg0 := let__14025
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15159 Obj

if True == tmp15162 {
tmp15160 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15160

tmp15161 := Call(__e, PrimFunc(symshen_4lookupsig), let__14024, V5018, V5020, V5021, let__14006, V5023)


ifres15159 = tmp15161


} else {
ifres15159 = False


}

ifres15155 = ifres15159


} else {
ifres15155 = False


}

ifres15152 = ifres15155


} else {
ifres15152 = False


}

ifres15149 = ifres15152


} else {
ifres15149 = False


}

ifres15147 = ifres15149


} else {
ifres15147 = False


}

let__14020 := ifres15147
_ = let__14020

tmp15823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14020, False)
}
__typedArg0 := let__14020
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15823 {
tmp15193 := Call(__e, PrimFunc(symshen_4unlocked_2), V5021)


var ifres15167 Obj

if True == tmp15193 {
tmp15168 := Call(__e, PrimFunc(symshen_4lazyderef), V5017, V5020)


let__14027 := tmp15168
_ = let__14027

tmp15192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14027)
}
__typedArg0 := let__14027
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15169 Obj

if True == tmp15192 {
tmp15170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14027)
}
__typedArg0 := let__14027
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14028 := tmp15170
_ = let__14028

tmp15171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14027)
}
__typedArg0 := let__14027
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15172 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15171, V5020)


let__14029 := tmp15172
_ = let__14029

tmp15191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14029)
}
__typedArg0 := let__14029
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15173 Obj

if True == tmp15191 {
tmp15174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14029)
}
__typedArg0 := let__14029
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14030 := tmp15174
_ = let__14030

tmp15175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14029)
}
__typedArg0 := let__14029
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15176 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15175, V5020)


let__14031 := tmp15176
_ = let__14031

tmp15190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14031, Nil)
}
__typedArg0 := let__14031
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15177 Obj

if True == tmp15190 {
tmp15178 := Call(__e, PrimFunc(symshen_4newpv), V5020)


let__14032 := tmp15178
_ = let__14032

tmp15179 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15179

tmp15180 := Call(__e, PrimFunc(symshen_4lazyderef), let__14028, V5020)


tmp15181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15180)
}
__typedArg0 := tmp15180
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

tmp15182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp15181)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp15181
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

tmp15183 := MakeNative(func(__e *ControlFlow) {
tmp15184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5018, Nil)
}
__typedArg0 := V5018
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15184)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15184
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14032, tmp15185)
}
__typedArg0 := let__14032
__typedArg1 := tmp15185
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15187 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14030, let__14032, V5019, V5020, V5021, let__14006, V5023)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4lookupsig), let__14028, tmp15186, V5020, V5021, let__14006, tmp15187)
return


}, 0)

tmp15188 := Call(__e, PrimFunc(symwhen), tmp15182, V5020, V5021, let__14006, tmp15183)


tmp15189 := Call(__e, PrimFunc(symshen_4gc), V5020, tmp15188)


ifres15177 = tmp15189


} else {
ifres15177 = False


}

ifres15173 = ifres15177


} else {
ifres15173 = False


}

ifres15169 = ifres15173


} else {
ifres15169 = False


}

ifres15167 = ifres15169


} else {
ifres15167 = False


}

let__14026 := ifres15167
_ = let__14026

tmp15821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14026, False)
}
__typedArg0 := let__14026
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15821 {
tmp15216 := Call(__e, PrimFunc(symshen_4unlocked_2), V5021)


var ifres15194 Obj

if True == tmp15216 {
tmp15195 := Call(__e, PrimFunc(symshen_4lazyderef), V5017, V5020)


let__14034 := tmp15195
_ = let__14034

tmp15215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14034)
}
__typedArg0 := let__14034
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15196 Obj

if True == tmp15215 {
tmp15197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14034)
}
__typedArg0 := let__14034
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14035 := tmp15197
_ = let__14035

tmp15198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14034)
}
__typedArg0 := let__14034
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15199 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15198, V5020)


let__14036 := tmp15199
_ = let__14036

tmp15214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14036)
}
__typedArg0 := let__14036
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15200 Obj

if True == tmp15214 {
tmp15201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14036)
}
__typedArg0 := let__14036
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14037 := tmp15201
_ = let__14037

tmp15202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14036)
}
__typedArg0 := let__14036
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15203 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15202, V5020)


let__14038 := tmp15203
_ = let__14038

tmp15213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14038, Nil)
}
__typedArg0 := let__14038
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15204 Obj

if True == tmp15213 {
tmp15205 := Call(__e, PrimFunc(symshen_4newpv), V5020)


let__14039 := tmp15205
_ = let__14039

tmp15206 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15206

tmp15207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5018, Nil)
}
__typedArg0 := V5018
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15207)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15207
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14039, tmp15208)
}
__typedArg0 := let__14039
__typedArg1 := tmp15208
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15210 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14037, let__14039, V5019, V5020, V5021, let__14006, V5023)
return
}, 0)

tmp15211 := Call(__e, PrimFunc(symshen_4system_1S_1h), let__14035, tmp15209, V5019, V5020, V5021, let__14006, tmp15210)


tmp15212 := Call(__e, PrimFunc(symshen_4gc), V5020, tmp15211)


ifres15204 = tmp15212


} else {
ifres15204 = False


}

ifres15200 = ifres15204


} else {
ifres15200 = False


}

ifres15196 = ifres15200


} else {
ifres15196 = False


}

ifres15194 = ifres15196


} else {
ifres15194 = False


}

let__14033 := ifres15194
_ = let__14033

tmp15819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14033, False)
}
__typedArg0 := let__14033
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15819 {
tmp15283 := Call(__e, PrimFunc(symshen_4unlocked_2), V5021)


var ifres15217 Obj

if True == tmp15283 {
tmp15218 := Call(__e, PrimFunc(symshen_4lazyderef), V5017, V5020)


let__14041 := tmp15218
_ = let__14041

tmp15282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14041)
}
__typedArg0 := let__14041
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15219 Obj

if True == tmp15282 {
tmp15220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14041)
}
__typedArg0 := let__14041
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15221 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15220, V5020)


let__14042 := tmp15221
_ = let__14042

tmp15281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14042, symcons)
}
__typedArg0 := let__14042
__typedArg1 := symcons
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15222 Obj

if True == tmp15281 {
tmp15223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14041)
}
__typedArg0 := let__14041
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15224 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15223, V5020)


let__14043 := tmp15224
_ = let__14043

tmp15280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14043)
}
__typedArg0 := let__14043
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15225 Obj

if True == tmp15280 {
tmp15226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14043)
}
__typedArg0 := let__14043
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14044 := tmp15226
_ = let__14044

tmp15227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14043)
}
__typedArg0 := let__14043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15228 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15227, V5020)


let__14045 := tmp15228
_ = let__14045

tmp15279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14045)
}
__typedArg0 := let__14045
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15229 Obj

if True == tmp15279 {
tmp15230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14045)
}
__typedArg0 := let__14045
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14046 := tmp15230
_ = let__14046

tmp15231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14045)
}
__typedArg0 := let__14045
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15232 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15231, V5020)


let__14047 := tmp15232
_ = let__14047

tmp15278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14047, Nil)
}
__typedArg0 := let__14047
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15233 Obj

if True == tmp15278 {
tmp15234 := Call(__e, PrimFunc(symshen_4lazyderef), V5018, V5020)


let__14048 := tmp15234
_ = let__14048

tmp15235 := MakeNative(func(__e *ControlFlow) {
Z5068 := __e.Get(1)
_ = Z5068
tmp15236 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15236

tmp15237 := MakeNative(func(__e *ControlFlow) {
tmp15238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z5068, Nil)
}
__typedArg0 := Z5068
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp15238)
}
__typedArg0 := symlist
__typedArg1 := tmp15238
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14046, tmp15239, V5019, V5020, V5021, let__14006, V5023)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14044, Z5068, V5019, V5020, V5021, let__14006, tmp15237)
return


}, 1)

let__14049 := tmp15235
_ = let__14049

tmp15277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14048)
}
__typedArg0 := let__14048
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15240 Obj

if True == tmp15277 {
tmp15241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14048)
}
__typedArg0 := let__14048
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15242 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15241, V5020)


let__14050 := tmp15242
_ = let__14050

tmp15243 := MakeNative(func(__e *ControlFlow) {
tmp15244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14048)
}
__typedArg0 := let__14048
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15245 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15244, V5020)


let__14052 := tmp15245
_ = let__14052

tmp15246 := MakeNative(func(__e *ControlFlow) {
Z5073 := __e.Get(1)
_ = Z5073
__e.TailApply(let__14049, Z5073)
return
}, 1)

let__14053 := tmp15246
_ = let__14053

tmp15262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14052)
}
__typedArg0 := let__14052
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15262 {
tmp15247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14052)
}
__typedArg0 := let__14052
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14054 := tmp15247
_ = let__14054

tmp15248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14052)
}
__typedArg0 := let__14052
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15249 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15248, V5020)


let__14055 := tmp15249
_ = let__14055

tmp15250 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14053, let__14054)
return
}, 0)

let__14056 := tmp15250
_ = let__14056

tmp15254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14055, Nil)
}
__typedArg0 := let__14055
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15254 {
__e.TailApply(PrimFunc(symthaw), let__14056)
return
} else {
tmp15252 := Call(__e, PrimFunc(symshen_4pvar_2), let__14055)


if True == tmp15252 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14055, Nil, V5020, let__14056)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15260 := Call(__e, PrimFunc(symshen_4pvar_2), let__14052)


if True == tmp15260 {
tmp15255 := Call(__e, PrimFunc(symshen_4newpv), V5020)


let__14057 := tmp15255
_ = let__14057

tmp15256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14057, Nil)
}
__typedArg0 := let__14057
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15257 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14053, let__14057)
return
}, 0)

tmp15258 := Call(__e, PrimFunc(symshen_4bind_b), let__14052, tmp15256, V5020, tmp15257)


__e.TailApply(PrimFunc(symshen_4gc), V5020, tmp15258)
return


} else {
__e.Return(False)
return
}


}


}, 0)

let__14051 := tmp15243
_ = let__14051

tmp15268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14050, symlist)
}
__typedArg0 := let__14050
__typedArg1 := symlist
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15263 Obj

if True == tmp15268 {
tmp15264 := Call(__e, PrimFunc(symthaw), let__14051)


ifres15263 = tmp15264


} else {
tmp15267 := Call(__e, PrimFunc(symshen_4pvar_2), let__14050)


var ifres15265 Obj

if True == tmp15267 {
tmp15266 := Call(__e, PrimFunc(symshen_4bind_b), let__14050, symlist, V5020, let__14051)


ifres15265 = tmp15266


} else {
ifres15265 = False


}

ifres15263 = ifres15265


}

ifres15240 = ifres15263


} else {
tmp15276 := Call(__e, PrimFunc(symshen_4pvar_2), let__14048)


var ifres15269 Obj

if True == tmp15276 {
tmp15270 := Call(__e, PrimFunc(symshen_4newpv), V5020)


let__14058 := tmp15270
_ = let__14058

tmp15271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14058, Nil)
}
__typedArg0 := let__14058
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp15271)
}
__typedArg0 := symlist
__typedArg1 := tmp15271
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15273 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14049, let__14058)
return
}, 0)

tmp15274 := Call(__e, PrimFunc(symshen_4bind_b), let__14048, tmp15272, V5020, tmp15273)


tmp15275 := Call(__e, PrimFunc(symshen_4gc), V5020, tmp15274)


ifres15269 = tmp15275


} else {
ifres15269 = False


}

ifres15240 = ifres15269


}

ifres15233 = ifres15240


} else {
ifres15233 = False


}

ifres15229 = ifres15233


} else {
ifres15229 = False


}

ifres15225 = ifres15229


} else {
ifres15225 = False


}

ifres15222 = ifres15225


} else {
ifres15222 = False


}

ifres15219 = ifres15222


} else {
ifres15219 = False


}

ifres15217 = ifres15219


} else {
ifres15217 = False


}

let__14040 := ifres15217
_ = let__14040

tmp15817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14040, False)
}
__typedArg0 := let__14040
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15817 {
tmp15367 := Call(__e, PrimFunc(symshen_4unlocked_2), V5021)


var ifres15284 Obj

if True == tmp15367 {
tmp15285 := Call(__e, PrimFunc(symshen_4lazyderef), V5017, V5020)


let__14060 := tmp15285
_ = let__14060

tmp15366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14060)
}
__typedArg0 := let__14060
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15286 Obj

if True == tmp15366 {
tmp15287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14060)
}
__typedArg0 := let__14060
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15288 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15287, V5020)


let__14061 := tmp15288
_ = let__14061

tmp15365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14061, sym_8p)
}
__typedArg0 := let__14061
__typedArg1 := sym_8p
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15289 Obj

if True == tmp15365 {
tmp15290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14060)
}
__typedArg0 := let__14060
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15291 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15290, V5020)


let__14062 := tmp15291
_ = let__14062

tmp15364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14062)
}
__typedArg0 := let__14062
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15292 Obj

if True == tmp15364 {
tmp15293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14062)
}
__typedArg0 := let__14062
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14063 := tmp15293
_ = let__14063

tmp15294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14062)
}
__typedArg0 := let__14062
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15295 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15294, V5020)


let__14064 := tmp15295
_ = let__14064

tmp15363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14064)
}
__typedArg0 := let__14064
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15296 Obj

if True == tmp15363 {
tmp15297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14064)
}
__typedArg0 := let__14064
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14065 := tmp15297
_ = let__14065

tmp15298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14064)
}
__typedArg0 := let__14064
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15299 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15298, V5020)


let__14066 := tmp15299
_ = let__14066

tmp15362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14066, Nil)
}
__typedArg0 := let__14066
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15300 Obj

if True == tmp15362 {
tmp15301 := Call(__e, PrimFunc(symshen_4lazyderef), V5018, V5020)


let__14067 := tmp15301
_ = let__14067

tmp15302 := MakeNative(func(__e *ControlFlow) {
Z5089 := __e.Get(1)
_ = Z5089
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5090 := __e.Get(1)
_ = Z5090
tmp15303 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15303

tmp15304 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14065, Z5090, V5019, V5020, V5021, let__14006, V5023)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14063, Z5089, V5019, V5020, V5021, let__14006, tmp15304)
return


}, 1))
return
}, 1)

let__14068 := tmp15302
_ = let__14068

tmp15361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14067)
}
__typedArg0 := let__14067
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15305 Obj

if True == tmp15361 {
tmp15306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14067)
}
__typedArg0 := let__14067
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14069 := tmp15306
_ = let__14069

tmp15307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14067)
}
__typedArg0 := let__14067
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15308 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15307, V5020)


let__14070 := tmp15308
_ = let__14070

tmp15309 := MakeNative(func(__e *ControlFlow) {
Z5094 := __e.Get(1)
_ = Z5094
tmp15310 := Call(__e, let__14068, let__14069)


__e.TailApply(tmp15310, Z5094)
return


}, 1)

let__14071 := tmp15309
_ = let__14071

tmp15348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14070)
}
__typedArg0 := let__14070
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15311 Obj

if True == tmp15348 {
tmp15312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14070)
}
__typedArg0 := let__14070
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15313 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15312, V5020)


let__14072 := tmp15313
_ = let__14072

tmp15314 := MakeNative(func(__e *ControlFlow) {
tmp15315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14070)
}
__typedArg0 := let__14070
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15316 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15315, V5020)


let__14074 := tmp15316
_ = let__14074

tmp15317 := MakeNative(func(__e *ControlFlow) {
Z5099 := __e.Get(1)
_ = Z5099
__e.TailApply(let__14071, Z5099)
return
}, 1)

let__14075 := tmp15317
_ = let__14075

tmp15333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14074)
}
__typedArg0 := let__14074
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15333 {
tmp15318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14074)
}
__typedArg0 := let__14074
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14076 := tmp15318
_ = let__14076

tmp15319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14074)
}
__typedArg0 := let__14074
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15320 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15319, V5020)


let__14077 := tmp15320
_ = let__14077

tmp15321 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14075, let__14076)
return
}, 0)

let__14078 := tmp15321
_ = let__14078

tmp15325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14077, Nil)
}
__typedArg0 := let__14077
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15325 {
__e.TailApply(PrimFunc(symthaw), let__14078)
return
} else {
tmp15323 := Call(__e, PrimFunc(symshen_4pvar_2), let__14077)


if True == tmp15323 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14077, Nil, V5020, let__14078)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15331 := Call(__e, PrimFunc(symshen_4pvar_2), let__14074)


if True == tmp15331 {
tmp15326 := Call(__e, PrimFunc(symshen_4newpv), V5020)


let__14079 := tmp15326
_ = let__14079

tmp15327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14079, Nil)
}
__typedArg0 := let__14079
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15328 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14075, let__14079)
return
}, 0)

tmp15329 := Call(__e, PrimFunc(symshen_4bind_b), let__14074, tmp15327, V5020, tmp15328)


__e.TailApply(PrimFunc(symshen_4gc), V5020, tmp15329)
return


} else {
__e.Return(False)
return
}


}


}, 0)

let__14073 := tmp15314
_ = let__14073

tmp15339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14072, sym_d)
}
__typedArg0 := let__14072
__typedArg1 := sym_d
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15334 Obj

if True == tmp15339 {
tmp15335 := Call(__e, PrimFunc(symthaw), let__14073)


ifres15334 = tmp15335


} else {
tmp15338 := Call(__e, PrimFunc(symshen_4pvar_2), let__14072)


var ifres15336 Obj

if True == tmp15338 {
tmp15337 := Call(__e, PrimFunc(symshen_4bind_b), let__14072, sym_d, V5020, let__14073)


ifres15336 = tmp15337


} else {
ifres15336 = False


}

ifres15334 = ifres15336


}

ifres15311 = ifres15334


} else {
tmp15347 := Call(__e, PrimFunc(symshen_4pvar_2), let__14070)


var ifres15340 Obj

if True == tmp15347 {
tmp15341 := Call(__e, PrimFunc(symshen_4newpv), V5020)


let__14080 := tmp15341
_ = let__14080

tmp15342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14080, Nil)
}
__typedArg0 := let__14080
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp15342)
}
__typedArg0 := sym_d
__typedArg1 := tmp15342
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15344 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14071, let__14080)
return
}, 0)

tmp15345 := Call(__e, PrimFunc(symshen_4bind_b), let__14070, tmp15343, V5020, tmp15344)


tmp15346 := Call(__e, PrimFunc(symshen_4gc), V5020, tmp15345)


ifres15340 = tmp15346


} else {
ifres15340 = False


}

ifres15311 = ifres15340


}

ifres15305 = ifres15311


} else {
tmp15360 := Call(__e, PrimFunc(symshen_4pvar_2), let__14067)


var ifres15349 Obj

if True == tmp15360 {
tmp15350 := Call(__e, PrimFunc(symshen_4newpv), V5020)


let__14081 := tmp15350
_ = let__14081

tmp15351 := Call(__e, PrimFunc(symshen_4newpv), V5020)


let__14082 := tmp15351
_ = let__14082

tmp15352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14082, Nil)
}
__typedArg0 := let__14082
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp15352)
}
__typedArg0 := sym_d
__typedArg1 := tmp15352
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14081, tmp15353)
}
__typedArg0 := let__14081
__typedArg1 := tmp15353
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15355 := MakeNative(func(__e *ControlFlow) {
tmp15356 := Call(__e, let__14068, let__14081)


__e.TailApply(tmp15356, let__14082)
return


}, 0)

tmp15357 := Call(__e, PrimFunc(symshen_4bind_b), let__14067, tmp15354, V5020, tmp15355)


tmp15358 := Call(__e, PrimFunc(symshen_4gc), V5020, tmp15357)


tmp15359 := Call(__e, PrimFunc(symshen_4gc), V5020, tmp15358)


ifres15349 = tmp15359


} else {
ifres15349 = False


}

ifres15305 = ifres15349


}

ifres15300 = ifres15305


} else {
ifres15300 = False


}

ifres15296 = ifres15300


} else {
ifres15296 = False


}

ifres15292 = ifres15296


} else {
ifres15292 = False


}

ifres15289 = ifres15292


} else {
ifres15289 = False


}

ifres15286 = ifres15289


} else {
ifres15286 = False


}

ifres15284 = ifres15286


} else {
ifres15284 = False


}

let__14059 := ifres15284
_ = let__14059

tmp15815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14059, False)
}
__typedArg0 := let__14059
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15815 {
tmp15434 := Call(__e, PrimFunc(symshen_4unlocked_2), V5021)


var ifres15368 Obj

if True == tmp15434 {
tmp15369 := Call(__e, PrimFunc(symshen_4lazyderef), V5017, V5020)


let__14084 := tmp15369
_ = let__14084

tmp15433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14084)
}
__typedArg0 := let__14084
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15370 Obj

if True == tmp15433 {
tmp15371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14084)
}
__typedArg0 := let__14084
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15372 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15371, V5020)


let__14085 := tmp15372
_ = let__14085

tmp15432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14085, sym_8v)
}
__typedArg0 := let__14085
__typedArg1 := sym_8v
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15373 Obj

if True == tmp15432 {
tmp15374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14084)
}
__typedArg0 := let__14084
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15375 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15374, V5020)


let__14086 := tmp15375
_ = let__14086

tmp15431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14086)
}
__typedArg0 := let__14086
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15376 Obj

if True == tmp15431 {
tmp15377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14086)
}
__typedArg0 := let__14086
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14087 := tmp15377
_ = let__14087

tmp15378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14086)
}
__typedArg0 := let__14086
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15379 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15378, V5020)


let__14088 := tmp15379
_ = let__14088

tmp15430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14088)
}
__typedArg0 := let__14088
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15380 Obj

if True == tmp15430 {
tmp15381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14088)
}
__typedArg0 := let__14088
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14089 := tmp15381
_ = let__14089

tmp15382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14088)
}
__typedArg0 := let__14088
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15383 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15382, V5020)


let__14090 := tmp15383
_ = let__14090

tmp15429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14090, Nil)
}
__typedArg0 := let__14090
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15384 Obj

if True == tmp15429 {
tmp15385 := Call(__e, PrimFunc(symshen_4lazyderef), V5018, V5020)


let__14091 := tmp15385
_ = let__14091

tmp15386 := MakeNative(func(__e *ControlFlow) {
Z5117 := __e.Get(1)
_ = Z5117
tmp15387 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15387

tmp15388 := MakeNative(func(__e *ControlFlow) {
tmp15389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z5117, Nil)
}
__typedArg0 := Z5117
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp15389)
}
__typedArg0 := symvector
__typedArg1 := tmp15389
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14089, tmp15390, V5019, V5020, V5021, let__14006, V5023)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14087, Z5117, V5019, V5020, V5021, let__14006, tmp15388)
return


}, 1)

let__14092 := tmp15386
_ = let__14092

tmp15428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14091)
}
__typedArg0 := let__14091
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15391 Obj

if True == tmp15428 {
tmp15392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14091)
}
__typedArg0 := let__14091
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15393 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15392, V5020)


let__14093 := tmp15393
_ = let__14093

tmp15394 := MakeNative(func(__e *ControlFlow) {
tmp15395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14091)
}
__typedArg0 := let__14091
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15396 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15395, V5020)


let__14095 := tmp15396
_ = let__14095

tmp15397 := MakeNative(func(__e *ControlFlow) {
Z5122 := __e.Get(1)
_ = Z5122
__e.TailApply(let__14092, Z5122)
return
}, 1)

let__14096 := tmp15397
_ = let__14096

tmp15413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14095)
}
__typedArg0 := let__14095
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15413 {
tmp15398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14095)
}
__typedArg0 := let__14095
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14097 := tmp15398
_ = let__14097

tmp15399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14095)
}
__typedArg0 := let__14095
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15400 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15399, V5020)


let__14098 := tmp15400
_ = let__14098

tmp15401 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14096, let__14097)
return
}, 0)

let__14099 := tmp15401
_ = let__14099

tmp15405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14098, Nil)
}
__typedArg0 := let__14098
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15405 {
__e.TailApply(PrimFunc(symthaw), let__14099)
return
} else {
tmp15403 := Call(__e, PrimFunc(symshen_4pvar_2), let__14098)


if True == tmp15403 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14098, Nil, V5020, let__14099)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15411 := Call(__e, PrimFunc(symshen_4pvar_2), let__14095)


if True == tmp15411 {
tmp15406 := Call(__e, PrimFunc(symshen_4newpv), V5020)


let__14100 := tmp15406
_ = let__14100

tmp15407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14100, Nil)
}
__typedArg0 := let__14100
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15408 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14096, let__14100)
return
}, 0)

tmp15409 := Call(__e, PrimFunc(symshen_4bind_b), let__14095, tmp15407, V5020, tmp15408)


__e.TailApply(PrimFunc(symshen_4gc), V5020, tmp15409)
return


} else {
__e.Return(False)
return
}


}


}, 0)

let__14094 := tmp15394
_ = let__14094

tmp15419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14093, symvector)
}
__typedArg0 := let__14093
__typedArg1 := symvector
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15414 Obj

if True == tmp15419 {
tmp15415 := Call(__e, PrimFunc(symthaw), let__14094)


ifres15414 = tmp15415


} else {
tmp15418 := Call(__e, PrimFunc(symshen_4pvar_2), let__14093)


var ifres15416 Obj

if True == tmp15418 {
tmp15417 := Call(__e, PrimFunc(symshen_4bind_b), let__14093, symvector, V5020, let__14094)


ifres15416 = tmp15417


} else {
ifres15416 = False


}

ifres15414 = ifres15416


}

ifres15391 = ifres15414


} else {
tmp15427 := Call(__e, PrimFunc(symshen_4pvar_2), let__14091)


var ifres15420 Obj

if True == tmp15427 {
tmp15421 := Call(__e, PrimFunc(symshen_4newpv), V5020)


let__14101 := tmp15421
_ = let__14101

tmp15422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14101, Nil)
}
__typedArg0 := let__14101
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp15422)
}
__typedArg0 := symvector
__typedArg1 := tmp15422
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15424 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14092, let__14101)
return
}, 0)

tmp15425 := Call(__e, PrimFunc(symshen_4bind_b), let__14091, tmp15423, V5020, tmp15424)


tmp15426 := Call(__e, PrimFunc(symshen_4gc), V5020, tmp15425)


ifres15420 = tmp15426


} else {
ifres15420 = False


}

ifres15391 = ifres15420


}

ifres15384 = ifres15391


} else {
ifres15384 = False


}

ifres15380 = ifres15384


} else {
ifres15380 = False


}

ifres15376 = ifres15380


} else {
ifres15376 = False


}

ifres15373 = ifres15376


} else {
ifres15373 = False


}

ifres15370 = ifres15373


} else {
ifres15370 = False


}

ifres15368 = ifres15370


} else {
ifres15368 = False


}

let__14083 := ifres15368
_ = let__14083

tmp15813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14083, False)
}
__typedArg0 := let__14083
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15813 {
tmp15467 := Call(__e, PrimFunc(symshen_4unlocked_2), V5021)


var ifres15435 Obj

if True == tmp15467 {
tmp15436 := Call(__e, PrimFunc(symshen_4lazyderef), V5017, V5020)


let__14103 := tmp15436
_ = let__14103

tmp15466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14103)
}
__typedArg0 := let__14103
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15437 Obj

if True == tmp15466 {
tmp15438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14103)
}
__typedArg0 := let__14103
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15439 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15438, V5020)


let__14104 := tmp15439
_ = let__14104

tmp15465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14104, sym_8s)
}
__typedArg0 := let__14104
__typedArg1 := sym_8s
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15440 Obj

if True == tmp15465 {
tmp15441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14103)
}
__typedArg0 := let__14103
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15442 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15441, V5020)


let__14105 := tmp15442
_ = let__14105

tmp15464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14105)
}
__typedArg0 := let__14105
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15443 Obj

if True == tmp15464 {
tmp15444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14105)
}
__typedArg0 := let__14105
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14106 := tmp15444
_ = let__14106

tmp15445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14105)
}
__typedArg0 := let__14105
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15446 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15445, V5020)


let__14107 := tmp15446
_ = let__14107

tmp15463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14107)
}
__typedArg0 := let__14107
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15447 Obj

if True == tmp15463 {
tmp15448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14107)
}
__typedArg0 := let__14107
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14108 := tmp15448
_ = let__14108

tmp15449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14107)
}
__typedArg0 := let__14107
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15450 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15449, V5020)


let__14109 := tmp15450
_ = let__14109

tmp15462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14109, Nil)
}
__typedArg0 := let__14109
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15451 Obj

if True == tmp15462 {
tmp15452 := Call(__e, PrimFunc(symshen_4lazyderef), V5018, V5020)


let__14110 := tmp15452
_ = let__14110

tmp15453 := MakeNative(func(__e *ControlFlow) {
tmp15454 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15454

tmp15455 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14108, symstring, V5019, V5020, V5021, let__14006, V5023)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14106, symstring, V5019, V5020, V5021, let__14006, tmp15455)
return


}, 0)

let__14111 := tmp15453
_ = let__14111

tmp15461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14110, symstring)
}
__typedArg0 := let__14110
__typedArg1 := symstring
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15456 Obj

if True == tmp15461 {
tmp15457 := Call(__e, PrimFunc(symthaw), let__14111)


ifres15456 = tmp15457


} else {
tmp15460 := Call(__e, PrimFunc(symshen_4pvar_2), let__14110)


var ifres15458 Obj

if True == tmp15460 {
tmp15459 := Call(__e, PrimFunc(symshen_4bind_b), let__14110, symstring, V5020, let__14111)


ifres15458 = tmp15459


} else {
ifres15458 = False


}

ifres15456 = ifres15458


}

ifres15451 = ifres15456


} else {
ifres15451 = False


}

ifres15447 = ifres15451


} else {
ifres15447 = False


}

ifres15443 = ifres15447


} else {
ifres15443 = False


}

ifres15440 = ifres15443


} else {
ifres15440 = False


}

ifres15437 = ifres15440


} else {
ifres15437 = False


}

ifres15435 = ifres15437


} else {
ifres15435 = False


}

let__14102 := ifres15435
_ = let__14102

tmp15811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14102, False)
}
__typedArg0 := let__14102
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15811 {
tmp15567 := Call(__e, PrimFunc(symshen_4unlocked_2), V5021)


var ifres15468 Obj

if True == tmp15567 {
tmp15469 := Call(__e, PrimFunc(symshen_4lazyderef), V5017, V5020)


let__14113 := tmp15469
_ = let__14113

tmp15566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14113)
}
__typedArg0 := let__14113
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15470 Obj

if True == tmp15566 {
tmp15471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14113)
}
__typedArg0 := let__14113
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15472 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15471, V5020)


let__14114 := tmp15472
_ = let__14114

tmp15565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14114, symlambda)
}
__typedArg0 := let__14114
__typedArg1 := symlambda
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15473 Obj

if True == tmp15565 {
tmp15474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14113)
}
__typedArg0 := let__14113
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15475 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15474, V5020)


let__14115 := tmp15475
_ = let__14115

tmp15564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14115)
}
__typedArg0 := let__14115
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15476 Obj

if True == tmp15564 {
tmp15477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14115)
}
__typedArg0 := let__14115
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14116 := tmp15477
_ = let__14116

tmp15478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14115)
}
__typedArg0 := let__14115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15479 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15478, V5020)


let__14117 := tmp15479
_ = let__14117

tmp15563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14117)
}
__typedArg0 := let__14117
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15480 Obj

if True == tmp15563 {
tmp15481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14117)
}
__typedArg0 := let__14117
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14118 := tmp15481
_ = let__14118

tmp15482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14117)
}
__typedArg0 := let__14117
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15483 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15482, V5020)


let__14119 := tmp15483
_ = let__14119

tmp15562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14119, Nil)
}
__typedArg0 := let__14119
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15484 Obj

if True == tmp15562 {
tmp15485 := Call(__e, PrimFunc(symshen_4lazyderef), V5018, V5020)


let__14120 := tmp15485
_ = let__14120

tmp15486 := MakeNative(func(__e *ControlFlow) {
Z5148 := __e.Get(1)
_ = Z5148
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5149 := __e.Get(1)
_ = Z5149
tmp15487 := Call(__e, PrimFunc(symshen_4newpv), V5020)


let__14122 := tmp15487
_ = let__14122

tmp15488 := Call(__e, PrimFunc(symshen_4newpv), V5020)


let__14123 := tmp15488
_ = let__14123

tmp15489 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15489

tmp15490 := Call(__e, PrimFunc(symshen_4lazyderef), let__14116, V5020)


tmp15491 := Call(__e, PrimFunc(symshen_4freshterm), tmp15490)


tmp15492 := MakeNative(func(__e *ControlFlow) {
tmp15493 := Call(__e, PrimFunc(symshen_4lazyderef), let__14116, V5020)


tmp15494 := Call(__e, PrimFunc(symshen_4deref), let__14123, V5020)


tmp15495 := Call(__e, PrimFunc(symshen_4deref), let__14118, V5020)


tmp15496 := Call(__e, PrimFunc(symshen_4beta), tmp15493, tmp15494, tmp15495)


tmp15497 := MakeNative(func(__e *ControlFlow) {
tmp15498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp15499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z5148, Nil)
}
__typedArg0 := Z5148
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15498, tmp15499)
}
__typedArg0 := tmp15498
__typedArg1 := tmp15499
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14123, tmp15500)
}
__typedArg0 := let__14123
__typedArg1 := tmp15500
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15501, V5019)
}
__typedArg0 := tmp15501
__typedArg1 := V5019
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14122, Z5149, tmp15502, V5020, V5021, let__14006, V5023)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14122, tmp15496, V5020, V5021, let__14006, tmp15497)
return


}, 0)

tmp15503 := Call(__e, PrimFunc(symbind), let__14123, tmp15491, V5020, V5021, let__14006, tmp15492)


tmp15504 := Call(__e, PrimFunc(symshen_4gc), V5020, tmp15503)


__e.TailApply(PrimFunc(symshen_4gc), V5020, tmp15504)
return


}, 1))
return
}, 1)

let__14121 := tmp15486
_ = let__14121

tmp15561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14120)
}
__typedArg0 := let__14120
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15505 Obj

if True == tmp15561 {
tmp15506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14120)
}
__typedArg0 := let__14120
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14124 := tmp15506
_ = let__14124

tmp15507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14120)
}
__typedArg0 := let__14120
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15508 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15507, V5020)


let__14125 := tmp15508
_ = let__14125

tmp15509 := MakeNative(func(__e *ControlFlow) {
Z5155 := __e.Get(1)
_ = Z5155
tmp15510 := Call(__e, let__14121, let__14124)


__e.TailApply(tmp15510, Z5155)
return


}, 1)

let__14126 := tmp15509
_ = let__14126

tmp15548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14125)
}
__typedArg0 := let__14125
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15511 Obj

if True == tmp15548 {
tmp15512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14125)
}
__typedArg0 := let__14125
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15513 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15512, V5020)


let__14127 := tmp15513
_ = let__14127

tmp15514 := MakeNative(func(__e *ControlFlow) {
tmp15515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14125)
}
__typedArg0 := let__14125
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15516 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15515, V5020)


let__14129 := tmp15516
_ = let__14129

tmp15517 := MakeNative(func(__e *ControlFlow) {
Z5160 := __e.Get(1)
_ = Z5160
__e.TailApply(let__14126, Z5160)
return
}, 1)

let__14130 := tmp15517
_ = let__14130

tmp15533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14129)
}
__typedArg0 := let__14129
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15533 {
tmp15518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14129)
}
__typedArg0 := let__14129
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14131 := tmp15518
_ = let__14131

tmp15519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14129)
}
__typedArg0 := let__14129
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15520 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15519, V5020)


let__14132 := tmp15520
_ = let__14132

tmp15521 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14130, let__14131)
return
}, 0)

let__14133 := tmp15521
_ = let__14133

tmp15525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14132, Nil)
}
__typedArg0 := let__14132
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15525 {
__e.TailApply(PrimFunc(symthaw), let__14133)
return
} else {
tmp15523 := Call(__e, PrimFunc(symshen_4pvar_2), let__14132)


if True == tmp15523 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14132, Nil, V5020, let__14133)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15531 := Call(__e, PrimFunc(symshen_4pvar_2), let__14129)


if True == tmp15531 {
tmp15526 := Call(__e, PrimFunc(symshen_4newpv), V5020)


let__14134 := tmp15526
_ = let__14134

tmp15527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14134, Nil)
}
__typedArg0 := let__14134
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15528 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14130, let__14134)
return
}, 0)

tmp15529 := Call(__e, PrimFunc(symshen_4bind_b), let__14129, tmp15527, V5020, tmp15528)


__e.TailApply(PrimFunc(symshen_4gc), V5020, tmp15529)
return


} else {
__e.Return(False)
return
}


}


}, 0)

let__14128 := tmp15514
_ = let__14128

tmp15539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14127, sym_1_1_6)
}
__typedArg0 := let__14127
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15534 Obj

if True == tmp15539 {
tmp15535 := Call(__e, PrimFunc(symthaw), let__14128)


ifres15534 = tmp15535


} else {
tmp15538 := Call(__e, PrimFunc(symshen_4pvar_2), let__14127)


var ifres15536 Obj

if True == tmp15538 {
tmp15537 := Call(__e, PrimFunc(symshen_4bind_b), let__14127, sym_1_1_6, V5020, let__14128)


ifres15536 = tmp15537


} else {
ifres15536 = False


}

ifres15534 = ifres15536


}

ifres15511 = ifres15534


} else {
tmp15547 := Call(__e, PrimFunc(symshen_4pvar_2), let__14125)


var ifres15540 Obj

if True == tmp15547 {
tmp15541 := Call(__e, PrimFunc(symshen_4newpv), V5020)


let__14135 := tmp15541
_ = let__14135

tmp15542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14135, Nil)
}
__typedArg0 := let__14135
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15542)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15542
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15544 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14126, let__14135)
return
}, 0)

tmp15545 := Call(__e, PrimFunc(symshen_4bind_b), let__14125, tmp15543, V5020, tmp15544)


tmp15546 := Call(__e, PrimFunc(symshen_4gc), V5020, tmp15545)


ifres15540 = tmp15546


} else {
ifres15540 = False


}

ifres15511 = ifres15540


}

ifres15505 = ifres15511


} else {
tmp15560 := Call(__e, PrimFunc(symshen_4pvar_2), let__14120)


var ifres15549 Obj

if True == tmp15560 {
tmp15550 := Call(__e, PrimFunc(symshen_4newpv), V5020)


let__14136 := tmp15550
_ = let__14136

tmp15551 := Call(__e, PrimFunc(symshen_4newpv), V5020)


let__14137 := tmp15551
_ = let__14137

tmp15552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14137, Nil)
}
__typedArg0 := let__14137
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15552)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15552
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14136, tmp15553)
}
__typedArg0 := let__14136
__typedArg1 := tmp15553
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15555 := MakeNative(func(__e *ControlFlow) {
tmp15556 := Call(__e, let__14121, let__14136)


__e.TailApply(tmp15556, let__14137)
return


}, 0)

tmp15557 := Call(__e, PrimFunc(symshen_4bind_b), let__14120, tmp15554, V5020, tmp15555)


tmp15558 := Call(__e, PrimFunc(symshen_4gc), V5020, tmp15557)


tmp15559 := Call(__e, PrimFunc(symshen_4gc), V5020, tmp15558)


ifres15549 = tmp15559


} else {
ifres15549 = False


}

ifres15505 = ifres15549


}

ifres15484 = ifres15505


} else {
ifres15484 = False


}

ifres15480 = ifres15484


} else {
ifres15480 = False


}

ifres15476 = ifres15480


} else {
ifres15476 = False


}

ifres15473 = ifres15476


} else {
ifres15473 = False


}

ifres15470 = ifres15473


} else {
ifres15470 = False


}

ifres15468 = ifres15470


} else {
ifres15468 = False


}

let__14112 := ifres15468
_ = let__14112

tmp15809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14112, False)
}
__typedArg0 := let__14112
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15809 {
tmp15617 := Call(__e, PrimFunc(symshen_4unlocked_2), V5021)


var ifres15568 Obj

if True == tmp15617 {
tmp15569 := Call(__e, PrimFunc(symshen_4lazyderef), V5017, V5020)


let__14139 := tmp15569
_ = let__14139

tmp15616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14139)
}
__typedArg0 := let__14139
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15570 Obj

if True == tmp15616 {
tmp15571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14139)
}
__typedArg0 := let__14139
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15572 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15571, V5020)


let__14140 := tmp15572
_ = let__14140

tmp15615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14140, symlet)
}
__typedArg0 := let__14140
__typedArg1 := symlet
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15573 Obj

if True == tmp15615 {
tmp15574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14139)
}
__typedArg0 := let__14139
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15575 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15574, V5020)


let__14141 := tmp15575
_ = let__14141

tmp15614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14141)
}
__typedArg0 := let__14141
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15576 Obj

if True == tmp15614 {
tmp15577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14141)
}
__typedArg0 := let__14141
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14142 := tmp15577
_ = let__14142

tmp15578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14141)
}
__typedArg0 := let__14141
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15579 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15578, V5020)


let__14143 := tmp15579
_ = let__14143

tmp15613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14143)
}
__typedArg0 := let__14143
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15580 Obj

if True == tmp15613 {
tmp15581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14143)
}
__typedArg0 := let__14143
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14144 := tmp15581
_ = let__14144

tmp15582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14143)
}
__typedArg0 := let__14143
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15583 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15582, V5020)


let__14145 := tmp15583
_ = let__14145

tmp15612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14145)
}
__typedArg0 := let__14145
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15584 Obj

if True == tmp15612 {
tmp15585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14145)
}
__typedArg0 := let__14145
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14146 := tmp15585
_ = let__14146

tmp15586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14145)
}
__typedArg0 := let__14145
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15587 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15586, V5020)


let__14147 := tmp15587
_ = let__14147

tmp15611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14147, Nil)
}
__typedArg0 := let__14147
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15588 Obj

if True == tmp15611 {
tmp15589 := Call(__e, PrimFunc(symshen_4newpv), V5020)


let__14148 := tmp15589
_ = let__14148

tmp15590 := Call(__e, PrimFunc(symshen_4newpv), V5020)


let__14149 := tmp15590
_ = let__14149

tmp15591 := Call(__e, PrimFunc(symshen_4newpv), V5020)


let__14150 := tmp15591
_ = let__14150

tmp15592 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15592

tmp15593 := MakeNative(func(__e *ControlFlow) {
tmp15594 := Call(__e, PrimFunc(symshen_4lazyderef), let__14142, V5020)


tmp15595 := Call(__e, PrimFunc(symshen_4freshterm), tmp15594)


tmp15596 := MakeNative(func(__e *ControlFlow) {
tmp15597 := Call(__e, PrimFunc(symshen_4lazyderef), let__14142, V5020)


tmp15598 := Call(__e, PrimFunc(symshen_4lazyderef), let__14149, V5020)


tmp15599 := Call(__e, PrimFunc(symshen_4lazyderef), let__14146, V5020)


tmp15600 := Call(__e, PrimFunc(symshen_4beta), tmp15597, tmp15598, tmp15599)


tmp15601 := MakeNative(func(__e *ControlFlow) {
tmp15602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp15603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14150, Nil)
}
__typedArg0 := let__14150
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15602, tmp15603)
}
__typedArg0 := tmp15602
__typedArg1 := tmp15603
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14149, tmp15604)
}
__typedArg0 := let__14149
__typedArg1 := tmp15604
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15605, V5019)
}
__typedArg0 := tmp15605
__typedArg1 := V5019
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14148, V5018, tmp15606, V5020, V5021, let__14006, V5023)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14148, tmp15600, V5020, V5021, let__14006, tmp15601)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14149, tmp15595, V5020, V5021, let__14006, tmp15596)
return


}, 0)

tmp15607 := Call(__e, PrimFunc(symshen_4system_1S_1h), let__14144, let__14150, V5019, V5020, V5021, let__14006, tmp15593)


tmp15608 := Call(__e, PrimFunc(symshen_4gc), V5020, tmp15607)


tmp15609 := Call(__e, PrimFunc(symshen_4gc), V5020, tmp15608)


tmp15610 := Call(__e, PrimFunc(symshen_4gc), V5020, tmp15609)


ifres15588 = tmp15610


} else {
ifres15588 = False


}

ifres15584 = ifres15588


} else {
ifres15584 = False


}

ifres15580 = ifres15584


} else {
ifres15580 = False


}

ifres15576 = ifres15580


} else {
ifres15576 = False


}

ifres15573 = ifres15576


} else {
ifres15573 = False


}

ifres15570 = ifres15573


} else {
ifres15570 = False


}

ifres15568 = ifres15570


} else {
ifres15568 = False


}

let__14138 := ifres15568
_ = let__14138

tmp15807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14138, False)
}
__typedArg0 := let__14138
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15807 {
tmp15687 := Call(__e, PrimFunc(symshen_4unlocked_2), V5021)


var ifres15618 Obj

if True == tmp15687 {
tmp15619 := Call(__e, PrimFunc(symshen_4lazyderef), V5017, V5020)


let__14152 := tmp15619
_ = let__14152

tmp15686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14152)
}
__typedArg0 := let__14152
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15620 Obj

if True == tmp15686 {
tmp15621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14152)
}
__typedArg0 := let__14152
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15622 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15621, V5020)


let__14153 := tmp15622
_ = let__14153

tmp15685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14153, symopen)
}
__typedArg0 := let__14153
__typedArg1 := symopen
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15623 Obj

if True == tmp15685 {
tmp15624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14152)
}
__typedArg0 := let__14152
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15625 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15624, V5020)


let__14154 := tmp15625
_ = let__14154

tmp15684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14154)
}
__typedArg0 := let__14154
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15626 Obj

if True == tmp15684 {
tmp15627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14154)
}
__typedArg0 := let__14154
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14155 := tmp15627
_ = let__14155

tmp15628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14154)
}
__typedArg0 := let__14154
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15629 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15628, V5020)


let__14156 := tmp15629
_ = let__14156

tmp15683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14156)
}
__typedArg0 := let__14156
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15630 Obj

if True == tmp15683 {
tmp15631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14156)
}
__typedArg0 := let__14156
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14157 := tmp15631
_ = let__14157

tmp15632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14156)
}
__typedArg0 := let__14156
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15633 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15632, V5020)


let__14158 := tmp15633
_ = let__14158

tmp15682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14158, Nil)
}
__typedArg0 := let__14158
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15634 Obj

if True == tmp15682 {
tmp15635 := Call(__e, PrimFunc(symshen_4lazyderef), V5018, V5020)


let__14159 := tmp15635
_ = let__14159

tmp15636 := MakeNative(func(__e *ControlFlow) {
Z5191 := __e.Get(1)
_ = Z5191
tmp15637 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15637

tmp15638 := MakeNative(func(__e *ControlFlow) {
tmp15639 := Call(__e, PrimFunc(symshen_4lazyderef), Z5191, V5020)


tmp15640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, tmp15640)
}
__typedArg0 := symin
__typedArg1 := tmp15640
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15642 := Call(__e, PrimFunc(symelement_2), tmp15639, tmp15641)


tmp15643 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14155, symstring, V5019, V5020, V5021, let__14006, V5023)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15642, V5020, V5021, let__14006, tmp15643)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), let__14157, Z5191, V5020, V5021, let__14006, tmp15638)
return


}, 1)

let__14160 := tmp15636
_ = let__14160

tmp15681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14159)
}
__typedArg0 := let__14159
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15644 Obj

if True == tmp15681 {
tmp15645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14159)
}
__typedArg0 := let__14159
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15646 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15645, V5020)


let__14161 := tmp15646
_ = let__14161

tmp15647 := MakeNative(func(__e *ControlFlow) {
tmp15648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14159)
}
__typedArg0 := let__14159
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15649 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15648, V5020)


let__14163 := tmp15649
_ = let__14163

tmp15650 := MakeNative(func(__e *ControlFlow) {
Z5196 := __e.Get(1)
_ = Z5196
__e.TailApply(let__14160, Z5196)
return
}, 1)

let__14164 := tmp15650
_ = let__14164

tmp15666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14163)
}
__typedArg0 := let__14163
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15666 {
tmp15651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14163)
}
__typedArg0 := let__14163
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14165 := tmp15651
_ = let__14165

tmp15652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14163)
}
__typedArg0 := let__14163
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15653 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15652, V5020)


let__14166 := tmp15653
_ = let__14166

tmp15654 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14164, let__14165)
return
}, 0)

let__14167 := tmp15654
_ = let__14167

tmp15658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14166, Nil)
}
__typedArg0 := let__14166
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15658 {
__e.TailApply(PrimFunc(symthaw), let__14167)
return
} else {
tmp15656 := Call(__e, PrimFunc(symshen_4pvar_2), let__14166)


if True == tmp15656 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14166, Nil, V5020, let__14167)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15664 := Call(__e, PrimFunc(symshen_4pvar_2), let__14163)


if True == tmp15664 {
tmp15659 := Call(__e, PrimFunc(symshen_4newpv), V5020)


let__14168 := tmp15659
_ = let__14168

tmp15660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14168, Nil)
}
__typedArg0 := let__14168
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15661 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14164, let__14168)
return
}, 0)

tmp15662 := Call(__e, PrimFunc(symshen_4bind_b), let__14163, tmp15660, V5020, tmp15661)


__e.TailApply(PrimFunc(symshen_4gc), V5020, tmp15662)
return


} else {
__e.Return(False)
return
}


}


}, 0)

let__14162 := tmp15647
_ = let__14162

tmp15672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14161, symstream)
}
__typedArg0 := let__14161
__typedArg1 := symstream
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15667 Obj

if True == tmp15672 {
tmp15668 := Call(__e, PrimFunc(symthaw), let__14162)


ifres15667 = tmp15668


} else {
tmp15671 := Call(__e, PrimFunc(symshen_4pvar_2), let__14161)


var ifres15669 Obj

if True == tmp15671 {
tmp15670 := Call(__e, PrimFunc(symshen_4bind_b), let__14161, symstream, V5020, let__14162)


ifres15669 = tmp15670


} else {
ifres15669 = False


}

ifres15667 = ifres15669


}

ifres15644 = ifres15667


} else {
tmp15680 := Call(__e, PrimFunc(symshen_4pvar_2), let__14159)


var ifres15673 Obj

if True == tmp15680 {
tmp15674 := Call(__e, PrimFunc(symshen_4newpv), V5020)


let__14169 := tmp15674
_ = let__14169

tmp15675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14169, Nil)
}
__typedArg0 := let__14169
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp15675)
}
__typedArg0 := symstream
__typedArg1 := tmp15675
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15677 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14160, let__14169)
return
}, 0)

tmp15678 := Call(__e, PrimFunc(symshen_4bind_b), let__14159, tmp15676, V5020, tmp15677)


tmp15679 := Call(__e, PrimFunc(symshen_4gc), V5020, tmp15678)


ifres15673 = tmp15679


} else {
ifres15673 = False


}

ifres15644 = ifres15673


}

ifres15634 = ifres15644


} else {
ifres15634 = False


}

ifres15630 = ifres15634


} else {
ifres15630 = False


}

ifres15626 = ifres15630


} else {
ifres15626 = False


}

ifres15623 = ifres15626


} else {
ifres15623 = False


}

ifres15620 = ifres15623


} else {
ifres15620 = False


}

ifres15618 = ifres15620


} else {
ifres15618 = False


}

let__14151 := ifres15618
_ = let__14151

tmp15805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14151, False)
}
__typedArg0 := let__14151
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15805 {
tmp15716 := Call(__e, PrimFunc(symshen_4unlocked_2), V5021)


var ifres15688 Obj

if True == tmp15716 {
tmp15689 := Call(__e, PrimFunc(symshen_4lazyderef), V5017, V5020)


let__14171 := tmp15689
_ = let__14171

tmp15715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14171)
}
__typedArg0 := let__14171
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15690 Obj

if True == tmp15715 {
tmp15691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14171)
}
__typedArg0 := let__14171
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15692 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15691, V5020)


let__14172 := tmp15692
_ = let__14172

tmp15714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14172, symtype)
}
__typedArg0 := let__14172
__typedArg1 := symtype
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15693 Obj

if True == tmp15714 {
tmp15694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14171)
}
__typedArg0 := let__14171
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15695 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15694, V5020)


let__14173 := tmp15695
_ = let__14173

tmp15713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14173)
}
__typedArg0 := let__14173
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15696 Obj

if True == tmp15713 {
tmp15697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14173)
}
__typedArg0 := let__14173
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14174 := tmp15697
_ = let__14174

tmp15698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14173)
}
__typedArg0 := let__14173
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15699 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15698, V5020)


let__14175 := tmp15699
_ = let__14175

tmp15712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14175)
}
__typedArg0 := let__14175
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15700 Obj

if True == tmp15712 {
tmp15701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14175)
}
__typedArg0 := let__14175
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14176 := tmp15701
_ = let__14176

tmp15702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14175)
}
__typedArg0 := let__14175
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15703 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15702, V5020)


let__14177 := tmp15703
_ = let__14177

tmp15711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14177, Nil)
}
__typedArg0 := let__14177
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15704 Obj

if True == tmp15711 {
tmp15705 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15705

tmp15706 := MakeNative(func(__e *ControlFlow) {
tmp15707 := Call(__e, PrimFunc(symshen_4deref), let__14176, V5020)


tmp15708 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp15707)


tmp15709 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14174, V5018, V5019, V5020, V5021, let__14006, V5023)
return
}, 0)

__e.TailApply(PrimFunc(symis_b), tmp15708, V5018, V5020, V5021, let__14006, tmp15709)
return


}, 0)

tmp15710 := Call(__e, PrimFunc(symshen_4cut), V5020, V5021, let__14006, tmp15706)


ifres15704 = tmp15710


} else {
ifres15704 = False


}

ifres15700 = ifres15704


} else {
ifres15700 = False


}

ifres15696 = ifres15700


} else {
ifres15696 = False


}

ifres15693 = ifres15696


} else {
ifres15693 = False


}

ifres15690 = ifres15693


} else {
ifres15690 = False


}

ifres15688 = ifres15690


} else {
ifres15688 = False


}

let__14170 := ifres15688
_ = let__14170

tmp15803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14170, False)
}
__typedArg0 := let__14170
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15803 {
tmp15747 := Call(__e, PrimFunc(symshen_4unlocked_2), V5021)


var ifres15717 Obj

if True == tmp15747 {
tmp15718 := Call(__e, PrimFunc(symshen_4lazyderef), V5017, V5020)


let__14179 := tmp15718
_ = let__14179

tmp15746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14179)
}
__typedArg0 := let__14179
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15719 Obj

if True == tmp15746 {
tmp15720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14179)
}
__typedArg0 := let__14179
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15721 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15720, V5020)


let__14180 := tmp15721
_ = let__14180

tmp15745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14180, symshen_4input_1h_7)
}
__typedArg0 := let__14180
__typedArg1 := symshen_4input_1h_7
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15722 Obj

if True == tmp15745 {
tmp15723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14179)
}
__typedArg0 := let__14179
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15724 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15723, V5020)


let__14181 := tmp15724
_ = let__14181

tmp15744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14181)
}
__typedArg0 := let__14181
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15725 Obj

if True == tmp15744 {
tmp15726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14181)
}
__typedArg0 := let__14181
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14182 := tmp15726
_ = let__14182

tmp15727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14181)
}
__typedArg0 := let__14181
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15728 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15727, V5020)


let__14183 := tmp15728
_ = let__14183

tmp15743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14183)
}
__typedArg0 := let__14183
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15729 Obj

if True == tmp15743 {
tmp15730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14183)
}
__typedArg0 := let__14183
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14184 := tmp15730
_ = let__14184

tmp15731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14183)
}
__typedArg0 := let__14183
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15732 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15731, V5020)


let__14185 := tmp15732
_ = let__14185

tmp15742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14185, Nil)
}
__typedArg0 := let__14185
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15733 Obj

if True == tmp15742 {
tmp15734 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15734

tmp15735 := Call(__e, PrimFunc(symshen_4deref), let__14182, V5020)


tmp15736 := Call(__e, PrimFunc(symshen_4rdecons), tmp15735)


tmp15737 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp15736)


tmp15738 := MakeNative(func(__e *ControlFlow) {
tmp15739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp15739)
}
__typedArg0 := symstream
__typedArg1 := tmp15739
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14184, tmp15740, V5019, V5020, V5021, let__14006, V5023)
return


}, 0)

tmp15741 := Call(__e, PrimFunc(symis), V5018, tmp15737, V5020, V5021, let__14006, tmp15738)


ifres15733 = tmp15741


} else {
ifres15733 = False


}

ifres15729 = ifres15733


} else {
ifres15729 = False


}

ifres15725 = ifres15729


} else {
ifres15725 = False


}

ifres15722 = ifres15725


} else {
ifres15722 = False


}

ifres15719 = ifres15722


} else {
ifres15719 = False


}

ifres15717 = ifres15719


} else {
ifres15717 = False


}

let__14178 := ifres15717
_ = let__14178

tmp15801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14178, False)
}
__typedArg0 := let__14178
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15801 {
tmp15776 := Call(__e, PrimFunc(symshen_4unlocked_2), V5021)


var ifres15748 Obj

if True == tmp15776 {
tmp15749 := Call(__e, PrimFunc(symshen_4lazyderef), V5017, V5020)


let__14187 := tmp15749
_ = let__14187

tmp15775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14187)
}
__typedArg0 := let__14187
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15750 Obj

if True == tmp15775 {
tmp15751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14187)
}
__typedArg0 := let__14187
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15752 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15751, V5020)


let__14188 := tmp15752
_ = let__14188

tmp15774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14188, symset)
}
__typedArg0 := let__14188
__typedArg1 := symset
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15753 Obj

if True == tmp15774 {
tmp15754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14187)
}
__typedArg0 := let__14187
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15755 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15754, V5020)


let__14189 := tmp15755
_ = let__14189

tmp15773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14189)
}
__typedArg0 := let__14189
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15756 Obj

if True == tmp15773 {
tmp15757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14189)
}
__typedArg0 := let__14189
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14190 := tmp15757
_ = let__14190

tmp15758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14189)
}
__typedArg0 := let__14189
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15759 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15758, V5020)


let__14191 := tmp15759
_ = let__14191

tmp15772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14191)
}
__typedArg0 := let__14191
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15760 Obj

if True == tmp15772 {
tmp15761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14191)
}
__typedArg0 := let__14191
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14192 := tmp15761
_ = let__14192

tmp15762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14191)
}
__typedArg0 := let__14191
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15763 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15762, V5020)


let__14193 := tmp15763
_ = let__14193

tmp15771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14193, Nil)
}
__typedArg0 := let__14193
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15764 Obj

if True == tmp15771 {
tmp15765 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15765

tmp15766 := MakeNative(func(__e *ControlFlow) {
tmp15767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14190, Nil)
}
__typedArg0 := let__14190
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp15767)
}
__typedArg0 := symvalue
__typedArg1 := tmp15767
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15769 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14192, V5018, V5019, V5020, V5021, let__14006, V5023)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp15768, V5018, V5019, V5020, V5021, let__14006, tmp15769)
return


}, 0)

tmp15770 := Call(__e, PrimFunc(symshen_4system_1S_1h), let__14190, symsymbol, V5019, V5020, V5021, let__14006, tmp15766)


ifres15764 = tmp15770


} else {
ifres15764 = False


}

ifres15760 = ifres15764


} else {
ifres15760 = False


}

ifres15756 = ifres15760


} else {
ifres15756 = False


}

ifres15753 = ifres15756


} else {
ifres15753 = False


}

ifres15750 = ifres15753


} else {
ifres15750 = False


}

ifres15748 = ifres15750


} else {
ifres15748 = False


}

let__14186 := ifres15748
_ = let__14186

tmp15799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14186, False)
}
__typedArg0 := let__14186
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15799 {
tmp15784 := Call(__e, PrimFunc(symshen_4unlocked_2), V5021)


var ifres15777 Obj

if True == tmp15784 {
tmp15778 := Call(__e, PrimFunc(symshen_4newpv), V5020)


let__14195 := tmp15778
_ = let__14195

tmp15779 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15779

tmp15780 := MakeNative(func(__e *ControlFlow) {
tmp15781 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), V5017, V5018, let__14195, V5020, V5021, let__14006, V5023)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5020, V5021, let__14006, tmp15781)
return


}, 0)

tmp15782 := Call(__e, PrimFunc(symshen_4l_1rules), V5019, let__14195, False, V5020, V5021, let__14006, tmp15780)


tmp15783 := Call(__e, PrimFunc(symshen_4gc), V5020, tmp15782)


ifres15777 = tmp15783


} else {
ifres15777 = False


}

let__14194 := ifres15777
_ = let__14194

tmp15797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14194, False)
}
__typedArg0 := let__14194
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15797 {
tmp15793 := Call(__e, PrimFunc(symshen_4unlocked_2), V5021)


var ifres15785 Obj

if True == tmp15793 {
tmp15786 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15786

tmp15787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp15788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5018, Nil)
}
__typedArg0 := V5018
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15787, tmp15788)
}
__typedArg0 := tmp15787
__typedArg1 := tmp15788
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5017, tmp15789)
}
__typedArg0 := V5017
__typedArg1 := tmp15789
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp15792 := Call(__e, PrimFunc(symshen_4search_1user_1datatypes), tmp15790, V5019, tmp15791, V5020, V5021, let__14006, V5023)


ifres15785 = tmp15792


} else {
ifres15785 = False


}

let__14196 := ifres15785
_ = let__14196

tmp15795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14196, False)
}
__typedArg0 := let__14196
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15795 {
__e.TailApply(PrimFunc(symshen_4unlock), V5021, let__14006)
return
} else {
__e.Return(let__14196)
return
}


} else {
__e.Return(let__14194)
return
}


} else {
__e.Return(let__14186)
return
}


} else {
__e.Return(let__14178)
return
}


} else {
__e.Return(let__14170)
return
}


} else {
__e.Return(let__14151)
return
}


} else {
__e.Return(let__14138)
return
}


} else {
__e.Return(let__14112)
return
}


} else {
__e.Return(let__14102)
return
}


} else {
__e.Return(let__14083)
return
}


} else {
__e.Return(let__14059)
return
}


} else {
__e.Return(let__14040)
return
}


} else {
__e.Return(let__14033)
return
}


} else {
__e.Return(let__14026)
return
}


} else {
__e.Return(let__14020)
return
}


} else {
__e.Return(let__14014)
return
}


} else {
__e.Return(let__14010)
return
}


} else {
__e.Return(let__14009)
return
}


} else {
__e.Return(let__14008)
return
}


} else {
__e.Return(let__14007)
return
}


}, 7)

tmp15834 := Call(__e, ns2_1set, symshen_4system_1S_1h, tmp15083)


_ = tmp15834

tmp15835 := MakeNative(func(__e *ControlFlow) {
V5229 := __e.Get(1)
_ = V5229
tmp15869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5229)
}
__typedArg0 := V5229
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15850 Obj

if True == tmp15869 {
tmp15867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5229)
}
__typedArg0 := V5229
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp15867)
}
__typedArg0 := symcons
__typedArg1 := tmp15867
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15852 Obj

if True == tmp15868 {
tmp15865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5229)
}
__typedArg0 := V5229
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15865)
}
__typedArg0 := tmp15865
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15854 Obj

if True == tmp15866 {
tmp15862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5229)
}
__typedArg0 := V5229
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15862)
}
__typedArg0 := tmp15862
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15863)
}
__typedArg0 := tmp15863
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15856 Obj

if True == tmp15864 {
tmp15858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5229)
}
__typedArg0 := V5229
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15858)
}
__typedArg0 := tmp15858
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15859)
}
__typedArg0 := tmp15859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp15860)
}
__typedArg0 := Nil
__typedArg1 := tmp15860
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15857 Obj

if True == tmp15861 {
ifres15857 = True


} else {
ifres15857 = False


}

ifres15856 = ifres15857


} else {
ifres15856 = False


}

var ifres15855 Obj

if True == ifres15856 {
ifres15855 = True


} else {
ifres15855 = False


}

ifres15854 = ifres15855


} else {
ifres15854 = False


}

var ifres15853 Obj

if True == ifres15854 {
ifres15853 = True


} else {
ifres15853 = False


}

ifres15852 = ifres15853


} else {
ifres15852 = False


}

var ifres15851 Obj

if True == ifres15852 {
ifres15851 = True


} else {
ifres15851 = False


}

ifres15850 = ifres15851


} else {
ifres15850 = False


}

if True == ifres15850 {
tmp15836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5229)
}
__typedArg0 := V5229
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15836)
}
__typedArg0 := tmp15836
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15838 := Call(__e, PrimFunc(symshen_4rdecons), tmp15837)


tmp15839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5229)
}
__typedArg0 := V5229
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15839)
}
__typedArg0 := tmp15839
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15840)
}
__typedArg0 := tmp15840
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15842 := Call(__e, PrimFunc(symshen_4rdecons), tmp15841)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15838, tmp15842)
}
__typedArg0 := tmp15838
__typedArg1 := tmp15842
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp15848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5229)
}
__typedArg0 := V5229
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15848 {
tmp15843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5229)
}
__typedArg0 := V5229
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15844 := Call(__e, PrimFunc(symshen_4rdecons), tmp15843)


tmp15845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5229)
}
__typedArg0 := V5229
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15846 := Call(__e, PrimFunc(symshen_4rdecons), tmp15845)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15844, tmp15846)
}
__typedArg0 := tmp15844
__typedArg1 := tmp15846
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V5229)
return
}


}


}, 1)

tmp15870 := Call(__e, ns2_1set, symshen_4rdecons, tmp15835)


_ = tmp15870

tmp15871 := MakeNative(func(__e *ControlFlow) {
V5230 := __e.Get(1)
_ = V5230
V5231 := __e.Get(2)
_ = V5231
V5232 := __e.Get(3)
_ = V5232
V5233 := __e.Get(4)
_ = V5233
V5234 := __e.Get(5)
_ = V5234
V5235 := __e.Get(6)
_ = V5235
tmp15884 := Call(__e, PrimFunc(symshen_4unlocked_2), V5233)


var ifres15872 Obj

if True == tmp15884 {
tmp15873 := Call(__e, PrimFunc(symshen_4lazyderef), V5231, V5232)


let__14198 := tmp15873
_ = let__14198

tmp15874 := MakeNative(func(__e *ControlFlow) {
tmp15875 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15875

tmp15876 := Call(__e, PrimFunc(symshen_4lazyderef), V5230, V5232)


tmp15877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(tmp15876)
}
__typedArg0 := tmp15876
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

__e.TailApply(PrimFunc(symwhen), tmp15877, V5232, V5233, V5234, V5235)
return


}, 0)

let__14199 := tmp15874
_ = let__14199

tmp15883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14198, symnumber)
}
__typedArg0 := let__14198
__typedArg1 := symnumber
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15878 Obj

if True == tmp15883 {
tmp15879 := Call(__e, PrimFunc(symthaw), let__14199)


ifres15878 = tmp15879


} else {
tmp15882 := Call(__e, PrimFunc(symshen_4pvar_2), let__14198)


var ifres15880 Obj

if True == tmp15882 {
tmp15881 := Call(__e, PrimFunc(symshen_4bind_b), let__14198, symnumber, V5232, let__14199)


ifres15880 = tmp15881


} else {
ifres15880 = False


}

ifres15878 = ifres15880


}

ifres15872 = ifres15878


} else {
ifres15872 = False


}

let__14197 := ifres15872
_ = let__14197

tmp15974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14197, False)
}
__typedArg0 := let__14197
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15974 {
tmp15897 := Call(__e, PrimFunc(symshen_4unlocked_2), V5233)


var ifres15885 Obj

if True == tmp15897 {
tmp15886 := Call(__e, PrimFunc(symshen_4lazyderef), V5231, V5232)


let__14201 := tmp15886
_ = let__14201

tmp15887 := MakeNative(func(__e *ControlFlow) {
tmp15888 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15888

tmp15889 := Call(__e, PrimFunc(symshen_4lazyderef), V5230, V5232)


tmp15890 := Call(__e, PrimFunc(symboolean_2), tmp15889)


__e.TailApply(PrimFunc(symwhen), tmp15890, V5232, V5233, V5234, V5235)
return


}, 0)

let__14202 := tmp15887
_ = let__14202

tmp15896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14201, symboolean)
}
__typedArg0 := let__14201
__typedArg1 := symboolean
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15891 Obj

if True == tmp15896 {
tmp15892 := Call(__e, PrimFunc(symthaw), let__14202)


ifres15891 = tmp15892


} else {
tmp15895 := Call(__e, PrimFunc(symshen_4pvar_2), let__14201)


var ifres15893 Obj

if True == tmp15895 {
tmp15894 := Call(__e, PrimFunc(symshen_4bind_b), let__14201, symboolean, V5232, let__14202)


ifres15893 = tmp15894


} else {
ifres15893 = False


}

ifres15891 = ifres15893


}

ifres15885 = ifres15891


} else {
ifres15885 = False


}

let__14200 := ifres15885
_ = let__14200

tmp15972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14200, False)
}
__typedArg0 := let__14200
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15972 {
tmp15910 := Call(__e, PrimFunc(symshen_4unlocked_2), V5233)


var ifres15898 Obj

if True == tmp15910 {
tmp15899 := Call(__e, PrimFunc(symshen_4lazyderef), V5231, V5232)


let__14204 := tmp15899
_ = let__14204

tmp15900 := MakeNative(func(__e *ControlFlow) {
tmp15901 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15901

tmp15902 := Call(__e, PrimFunc(symshen_4lazyderef), V5230, V5232)


tmp15903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(tmp15902)
}
__typedArg0 := tmp15902
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

__e.TailApply(PrimFunc(symwhen), tmp15903, V5232, V5233, V5234, V5235)
return


}, 0)

let__14205 := tmp15900
_ = let__14205

tmp15909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14204, symstring)
}
__typedArg0 := let__14204
__typedArg1 := symstring
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15904 Obj

if True == tmp15909 {
tmp15905 := Call(__e, PrimFunc(symthaw), let__14205)


ifres15904 = tmp15905


} else {
tmp15908 := Call(__e, PrimFunc(symshen_4pvar_2), let__14204)


var ifres15906 Obj

if True == tmp15908 {
tmp15907 := Call(__e, PrimFunc(symshen_4bind_b), let__14204, symstring, V5232, let__14205)


ifres15906 = tmp15907


} else {
ifres15906 = False


}

ifres15904 = ifres15906


}

ifres15898 = ifres15904


} else {
ifres15898 = False


}

let__14203 := ifres15898
_ = let__14203

tmp15970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14203, False)
}
__typedArg0 := let__14203
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15970 {
tmp15923 := Call(__e, PrimFunc(symshen_4unlocked_2), V5233)


var ifres15911 Obj

if True == tmp15923 {
tmp15912 := Call(__e, PrimFunc(symshen_4lazyderef), V5231, V5232)


let__14207 := tmp15912
_ = let__14207

tmp15913 := MakeNative(func(__e *ControlFlow) {
tmp15914 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15914

tmp15915 := Call(__e, PrimFunc(symshen_4lazyderef), V5230, V5232)


tmp15916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(tmp15915)
}
__typedArg0 := tmp15915
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

__e.TailApply(PrimFunc(symwhen), tmp15916, V5232, V5233, V5234, V5235)
return


}, 0)

let__14208 := tmp15913
_ = let__14208

tmp15922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14207, symsymbol)
}
__typedArg0 := let__14207
__typedArg1 := symsymbol
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15917 Obj

if True == tmp15922 {
tmp15918 := Call(__e, PrimFunc(symthaw), let__14208)


ifres15917 = tmp15918


} else {
tmp15921 := Call(__e, PrimFunc(symshen_4pvar_2), let__14207)


var ifres15919 Obj

if True == tmp15921 {
tmp15920 := Call(__e, PrimFunc(symshen_4bind_b), let__14207, symsymbol, V5232, let__14208)


ifres15919 = tmp15920


} else {
ifres15919 = False


}

ifres15917 = ifres15919


}

ifres15911 = ifres15917


} else {
ifres15911 = False


}

let__14206 := ifres15911
_ = let__14206

tmp15968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14206, False)
}
__typedArg0 := let__14206
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15968 {
tmp15966 := Call(__e, PrimFunc(symshen_4unlocked_2), V5233)


if True == tmp15966 {
tmp15924 := Call(__e, PrimFunc(symshen_4lazyderef), V5230, V5232)


let__14209 := tmp15924
_ = let__14209

tmp15964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14209, Nil)
}
__typedArg0 := let__14209
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15964 {
tmp15925 := Call(__e, PrimFunc(symshen_4lazyderef), V5231, V5232)


let__14210 := tmp15925
_ = let__14210

tmp15926 := MakeNative(func(__e *ControlFlow) {
Z5251 := __e.Get(1)
_ = Z5251
tmp15927 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15927

__e.TailApply(PrimFunc(symthaw), V5235)
return


}, 1)

let__14211 := tmp15926
_ = let__14211

tmp15962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14210)
}
__typedArg0 := let__14210
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15962 {
tmp15928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14210)
}
__typedArg0 := let__14210
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15929 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15928, V5232)


let__14212 := tmp15929
_ = let__14212

tmp15930 := MakeNative(func(__e *ControlFlow) {
tmp15931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14210)
}
__typedArg0 := let__14210
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15932 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15931, V5232)


let__14214 := tmp15932
_ = let__14214

tmp15933 := MakeNative(func(__e *ControlFlow) {
Z5256 := __e.Get(1)
_ = Z5256
__e.TailApply(let__14211, Z5256)
return
}, 1)

let__14215 := tmp15933
_ = let__14215

tmp15949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14214)
}
__typedArg0 := let__14214
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15949 {
tmp15934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14214)
}
__typedArg0 := let__14214
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14216 := tmp15934
_ = let__14216

tmp15935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14214)
}
__typedArg0 := let__14214
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15936 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15935, V5232)


let__14217 := tmp15936
_ = let__14217

tmp15937 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14215, let__14216)
return
}, 0)

let__14218 := tmp15937
_ = let__14218

tmp15941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14217, Nil)
}
__typedArg0 := let__14217
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15941 {
__e.TailApply(PrimFunc(symthaw), let__14218)
return
} else {
tmp15939 := Call(__e, PrimFunc(symshen_4pvar_2), let__14217)


if True == tmp15939 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14217, Nil, V5232, let__14218)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15947 := Call(__e, PrimFunc(symshen_4pvar_2), let__14214)


if True == tmp15947 {
tmp15942 := Call(__e, PrimFunc(symshen_4newpv), V5232)


let__14219 := tmp15942
_ = let__14219

tmp15943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14219, Nil)
}
__typedArg0 := let__14219
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15944 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14215, let__14219)
return
}, 0)

tmp15945 := Call(__e, PrimFunc(symshen_4bind_b), let__14214, tmp15943, V5232, tmp15944)


__e.TailApply(PrimFunc(symshen_4gc), V5232, tmp15945)
return


} else {
__e.Return(False)
return
}


}


}, 0)

let__14213 := tmp15930
_ = let__14213

tmp15953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14212, symlist)
}
__typedArg0 := let__14212
__typedArg1 := symlist
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15953 {
__e.TailApply(PrimFunc(symthaw), let__14213)
return
} else {
tmp15951 := Call(__e, PrimFunc(symshen_4pvar_2), let__14212)


if True == tmp15951 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14212, symlist, V5232, let__14213)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15960 := Call(__e, PrimFunc(symshen_4pvar_2), let__14210)


if True == tmp15960 {
tmp15954 := Call(__e, PrimFunc(symshen_4newpv), V5232)


let__14220 := tmp15954
_ = let__14220

tmp15955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14220, Nil)
}
__typedArg0 := let__14220
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp15955)
}
__typedArg0 := symlist
__typedArg1 := tmp15955
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15957 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14211, let__14220)
return
}, 0)

tmp15958 := Call(__e, PrimFunc(symshen_4bind_b), let__14210, tmp15956, V5232, tmp15957)


__e.TailApply(PrimFunc(symshen_4gc), V5232, tmp15958)
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
__e.Return(let__14206)
return
}


} else {
__e.Return(let__14203)
return
}


} else {
__e.Return(let__14200)
return
}


} else {
__e.Return(let__14197)
return
}


}, 6)

tmp15975 := Call(__e, ns2_1set, symshen_4primitive, tmp15871)


_ = tmp15975

tmp15976 := MakeNative(func(__e *ControlFlow) {
V5262 := __e.Get(1)
_ = V5262
V5263 := __e.Get(2)
_ = V5263
V5264 := __e.Get(3)
_ = V5264
V5265 := __e.Get(4)
_ = V5265
V5266 := __e.Get(5)
_ = V5266
V5267 := __e.Get(6)
_ = V5267
V5268 := __e.Get(7)
_ = V5268
tmp15979 := Call(__e, PrimFunc(symshen_4lazyderef), V5264, V5265)


tmp15980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp15979, Nil)
}
__typedArg0 := tmp15979
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15980 {
__e.Return(False)
return
} else {
tmp15977 := Call(__e, PrimFunc(symshen_4deref), V5262, V5265)


__e.TailApply(PrimFunc(symshen_4by_1hypothesis_1h), tmp15977, V5263, V5264, V5265, V5266, V5267, V5268)
return


}


}, 7)

tmp15981 := Call(__e, ns2_1set, symshen_4by_1hypothesis, tmp15976)


_ = tmp15981

tmp15982 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
__self6 := __e.Get(6)
__self7 := __e.Get(7)
__selftop:
V5269 := __self1
_ = V5269
V5270 := __self2
_ = V5270
V5271 := __self3
_ = V5271
V5272 := __self4
_ = V5272
V5273 := __self5
_ = V5273
V5274 := __self6
_ = V5274
V5275 := __self7
_ = V5275
tmp16015 := Call(__e, PrimFunc(symshen_4unlocked_2), V5273)


var ifres15983 Obj

if True == tmp16015 {
tmp15984 := Call(__e, PrimFunc(symshen_4lazyderef), V5271, V5272)


let__14222 := tmp15984
_ = let__14222

tmp16014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14222)
}
__typedArg0 := let__14222
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15985 Obj

if True == tmp16014 {
tmp15986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14222)
}
__typedArg0 := let__14222
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15987 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15986, V5272)


let__14223 := tmp15987
_ = let__14223

tmp16013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14223)
}
__typedArg0 := let__14223
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15988 Obj

if True == tmp16013 {
tmp15989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14223)
}
__typedArg0 := let__14223
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14224 := tmp15989
_ = let__14224

tmp15990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14223)
}
__typedArg0 := let__14223
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15991 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15990, V5272)


let__14225 := tmp15991
_ = let__14225

tmp16012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14225)
}
__typedArg0 := let__14225
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15992 Obj

if True == tmp16012 {
tmp15993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14225)
}
__typedArg0 := let__14225
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14226 := tmp15993
_ = let__14226

tmp15994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14225)
}
__typedArg0 := let__14225
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15995 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15994, V5272)


let__14227 := tmp15995
_ = let__14227

tmp16011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14227)
}
__typedArg0 := let__14227
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15996 Obj

if True == tmp16011 {
tmp15997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14227)
}
__typedArg0 := let__14227
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14228 := tmp15997
_ = let__14228

tmp15998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14227)
}
__typedArg0 := let__14227
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15999 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15998, V5272)


let__14229 := tmp15999
_ = let__14229

tmp16010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14229, Nil)
}
__typedArg0 := let__14229
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16000 Obj

if True == tmp16010 {
tmp16001 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16001

tmp16002 := Call(__e, PrimFunc(symshen_4deref), let__14226, V5272)


tmp16003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16002, tmp16003)
}
__typedArg0 := tmp16002
__typedArg1 := tmp16003
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16005 := MakeNative(func(__e *ControlFlow) {
tmp16006 := Call(__e, PrimFunc(symshen_4deref), let__14224, V5272)


tmp16007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V5269, tmp16006)
}
__typedArg0 := V5269
__typedArg1 := tmp16006
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16008 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symis_b), V5270, let__14228, V5272, V5273, V5274, V5275)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp16007, V5272, V5273, V5274, tmp16008)
return


}, 0)

tmp16009 := Call(__e, PrimFunc(symwhen), tmp16004, V5272, V5273, V5274, tmp16005)


ifres16000 = tmp16009


} else {
ifres16000 = False


}

ifres15996 = ifres16000


} else {
ifres15996 = False


}

ifres15992 = ifres15996


} else {
ifres15992 = False


}

ifres15988 = ifres15992


} else {
ifres15988 = False


}

ifres15985 = ifres15988


} else {
ifres15985 = False


}

ifres15983 = ifres15985


} else {
ifres15983 = False


}

let__14221 := ifres15983
_ = let__14221

tmp16024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14221, False)
}
__typedArg0 := let__14221
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16024 {
tmp16022 := Call(__e, PrimFunc(symshen_4unlocked_2), V5273)


if True == tmp16022 {
tmp16016 := Call(__e, PrimFunc(symshen_4lazyderef), V5271, V5272)


let__14230 := tmp16016
_ = let__14230

tmp16020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14230)
}
__typedArg0 := let__14230
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16020 {
tmp16017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14230)
}
__typedArg0 := let__14230
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14231 := tmp16017
_ = let__14231

tmp16018 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16018

if PrimFunc(symshen_4by_1hypothesis_1h) == __self {
__self1, __self2, __self3, __self4, __self5, __self6, __self7 = V5269, V5270, let__14231, V5272, V5273, V5274, V5275
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4by_1hypothesis_1h), V5269, V5270, let__14231, V5272, V5273, V5274, V5275)
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
__e.Return(let__14221)
return
}


}, 7)

tmp16025 := Call(__e, ns2_1set, symshen_4by_1hypothesis_1h, tmp15982)


_ = tmp16025

tmp16026 := MakeNative(func(__e *ControlFlow) {
V5287 := __e.Get(1)
_ = V5287
V5288 := __e.Get(2)
_ = V5288
V5289 := __e.Get(3)
_ = V5289
V5290 := __e.Get(4)
_ = V5290
V5291 := __e.Get(5)
_ = V5291
V5292 := __e.Get(6)
_ = V5292
tmp16031 := Call(__e, PrimFunc(symshen_4unlocked_2), V5290)


if True == tmp16031 {
tmp16027 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16027

tmp16028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp16029 := Call(__e, PrimFunc(symassoc), V5287, tmp16028)


__e.TailApply(PrimFunc(symshen_4sigf), tmp16029, V5288, V5289, V5290, V5291, V5292)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp16032 := Call(__e, ns2_1set, symshen_4lookupsig, tmp16026)


_ = tmp16032

tmp16033 := MakeNative(func(__e *ControlFlow) {
V5307 := __e.Get(1)
_ = V5307
V5308 := __e.Get(2)
_ = V5308
V5309 := __e.Get(3)
_ = V5309
V5310 := __e.Get(4)
_ = V5310
V5311 := __e.Get(5)
_ = V5311
V5312 := __e.Get(6)
_ = V5312
tmp16040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5307)
}
__typedArg0 := V5307
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16040 {
tmp16034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5307)
}
__typedArg0 := V5307
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16035 := Call(__e, tmp16034, V5308)


tmp16036 := Call(__e, tmp16035, V5309)


tmp16037 := Call(__e, tmp16036, V5310)


tmp16038 := Call(__e, tmp16037, V5311)


__e.TailApply(tmp16038, V5312)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp16041 := Call(__e, ns2_1set, symshen_4sigf, tmp16033)


_ = tmp16041

tmp16042 := MakeNative(func(__e *ControlFlow) {
V5313 := __e.Get(1)
_ = V5313
tmp16043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(MakeInteger(3))
}
__typedArg0 := MakeInteger(3)
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

let__14232 := tmp16043
_ = let__14232

tmp16044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__14232, MakeInteger(0), symshen_4print_1freshterm)
}
__typedArg0 := let__14232
__typedArg1 := MakeInteger(0)
__typedArg2 := symshen_4print_1freshterm
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__14233 := tmp16044
_ = let__14233

tmp16045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__14233, MakeInteger(1), V5313)
}
__typedArg0 := let__14233
__typedArg1 := MakeInteger(1)
__typedArg2 := V5313
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__14234 := tmp16045
_ = let__14234

tmp16046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dgensym_d)
}
__typedArg0 := symshen_4_dgensym_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp16048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dgensym_d, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp16046)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp16046
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := symshen_4_dgensym_d
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp16046)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp16046
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

tmp16049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__14234, MakeInteger(2), tmp16048)
}
__typedArg0 := let__14234
__typedArg1 := MakeInteger(2)
__typedArg2 := tmp16048
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__14235 := tmp16049
_ = let__14235

__e.Return(let__14235)
return


}, 1)

tmp16050 := Call(__e, ns2_1set, symshen_4freshterm, tmp16042)


_ = tmp16050

tmp16051 := MakeNative(func(__e *ControlFlow) {
V5318 := __e.Get(1)
_ = V5318
tmp16052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V5318, MakeInteger(1))
}
__typedArg0 := V5318
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp16053 := Call(__e, PrimFunc(symshen_4app), tmp16052, MakeString(""), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("&&"))
__typedS1, __typedOK1 := TypedString(tmp16053)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("&&")
__typedArg1 := tmp16053
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp16054 := Call(__e, ns2_1set, symshen_4print_1freshterm, tmp16051)


_ = tmp16054

tmp16055 := MakeNative(func(__e *ControlFlow) {
V5319 := __e.Get(1)
_ = V5319
V5320 := __e.Get(2)
_ = V5320
V5321 := __e.Get(3)
_ = V5321
V5322 := __e.Get(4)
_ = V5322
V5323 := __e.Get(5)
_ = V5323
V5324 := __e.Get(6)
_ = V5324
V5325 := __e.Get(7)
_ = V5325
tmp16059 := Call(__e, PrimFunc(symshen_4lazyderef), V5321, V5322)


tmp16060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16059, Nil)
}
__typedArg0 := tmp16059
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16060 {
__e.Return(False)
return
} else {
tmp16056 := Call(__e, PrimFunc(symshen_4deref), V5319, V5322)


tmp16057 := Call(__e, PrimFunc(symshen_4deref), V5320, V5322)


__e.TailApply(PrimFunc(symshen_4search_1user_1datatypes_1h), tmp16056, tmp16057, V5321, V5322, V5323, V5324, V5325)
return


}


}, 7)

tmp16061 := Call(__e, ns2_1set, symshen_4search_1user_1datatypes, tmp16055)


_ = tmp16061

tmp16062 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
__self6 := __e.Get(6)
__self7 := __e.Get(7)
__selftop:
V5326 := __self1
_ = V5326
V5327 := __self2
_ = V5327
V5328 := __self3
_ = V5328
V5329 := __self4
_ = V5329
V5330 := __self5
_ = V5330
V5331 := __self6
_ = V5331
V5332 := __self7
_ = V5332
tmp16077 := Call(__e, PrimFunc(symshen_4unlocked_2), V5330)


var ifres16063 Obj

if True == tmp16077 {
tmp16064 := Call(__e, PrimFunc(symshen_4lazyderef), V5328, V5329)


let__14237 := tmp16064
_ = let__14237

tmp16076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14237)
}
__typedArg0 := let__14237
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16065 Obj

if True == tmp16076 {
tmp16066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14237)
}
__typedArg0 := let__14237
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16067 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16066, V5329)


let__14238 := tmp16067
_ = let__14238

tmp16075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14238)
}
__typedArg0 := let__14238
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16068 Obj

if True == tmp16075 {
tmp16069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14238)
}
__typedArg0 := let__14238
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14239 := tmp16069
_ = let__14239

tmp16070 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16070

tmp16071 := Call(__e, PrimFunc(symshen_4deref), let__14239, V5329)


tmp16072 := Call(__e, tmp16071, V5326)


tmp16073 := Call(__e, tmp16072, V5327)


tmp16074 := Call(__e, PrimFunc(symcall), tmp16073, V5329, V5330, V5331, V5332)


ifres16068 = tmp16074


} else {
ifres16068 = False


}

ifres16065 = ifres16068


} else {
ifres16065 = False


}

ifres16063 = ifres16065


} else {
ifres16063 = False


}

let__14236 := ifres16063
_ = let__14236

tmp16086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14236, False)
}
__typedArg0 := let__14236
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16086 {
tmp16084 := Call(__e, PrimFunc(symshen_4unlocked_2), V5330)


if True == tmp16084 {
tmp16078 := Call(__e, PrimFunc(symshen_4lazyderef), V5328, V5329)


let__14240 := tmp16078
_ = let__14240

tmp16082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14240)
}
__typedArg0 := let__14240
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16082 {
tmp16079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14240)
}
__typedArg0 := let__14240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14241 := tmp16079
_ = let__14241

tmp16080 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16080

if PrimFunc(symshen_4search_1user_1datatypes_1h) == __self {
__self1, __self2, __self3, __self4, __self5, __self6, __self7 = V5326, V5327, let__14241, V5329, V5330, V5331, V5332
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4search_1user_1datatypes_1h), V5326, V5327, let__14241, V5329, V5330, V5331, V5332)
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
__e.Return(let__14236)
return
}


}, 7)

tmp16087 := Call(__e, ns2_1set, symshen_4search_1user_1datatypes_1h, tmp16062)


_ = tmp16087

tmp16088 := MakeNative(func(__e *ControlFlow) {
V5339 := __e.Get(1)
_ = V5339
V5340 := __e.Get(2)
_ = V5340
V5341 := __e.Get(3)
_ = V5341
V5342 := __e.Get(4)
_ = V5342
V5343 := __e.Get(5)
_ = V5343
V5344 := __e.Get(6)
_ = V5344
V5345 := __e.Get(7)
_ = V5345
let__14242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5344)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5344
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__14242

tmp16100 := Call(__e, PrimFunc(symshen_4unlocked_2), V5343)


var ifres16090 Obj

if True == tmp16100 {
tmp16091 := Call(__e, PrimFunc(symshen_4lazyderef), V5339, V5342)


let__14244 := tmp16091
_ = let__14244

tmp16099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14244, Nil)
}
__typedArg0 := let__14244
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16092 Obj

if True == tmp16099 {
tmp16093 := Call(__e, PrimFunc(symshen_4lazyderef), V5341, V5342)


let__14245 := tmp16093
_ = let__14245

tmp16098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14245, True)
}
__typedArg0 := let__14245
__typedArg1 := True
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16094 Obj

if True == tmp16098 {
tmp16095 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16095

tmp16096 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symbind), V5340, Nil, V5342, V5343, let__14242, V5345)
return
}, 0)

tmp16097 := Call(__e, PrimFunc(symshen_4cut), V5342, V5343, let__14242, tmp16096)


ifres16094 = tmp16097


} else {
ifres16094 = False


}

ifres16092 = ifres16094


} else {
ifres16092 = False


}

ifres16090 = ifres16092


} else {
ifres16090 = False


}

let__14243 := ifres16090
_ = let__14243

tmp16443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14243, False)
}
__typedArg0 := let__14243
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16443 {
tmp16179 := Call(__e, PrimFunc(symshen_4unlocked_2), V5343)


var ifres16101 Obj

if True == tmp16179 {
tmp16102 := Call(__e, PrimFunc(symshen_4lazyderef), V5339, V5342)


let__14247 := tmp16102
_ = let__14247

tmp16178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14247)
}
__typedArg0 := let__14247
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16103 Obj

if True == tmp16178 {
tmp16104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14247)
}
__typedArg0 := let__14247
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16105 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16104, V5342)


let__14248 := tmp16105
_ = let__14248

tmp16177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14248)
}
__typedArg0 := let__14248
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16106 Obj

if True == tmp16177 {
tmp16107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14248)
}
__typedArg0 := let__14248
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16108 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16107, V5342)


let__14249 := tmp16108
_ = let__14249

tmp16176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14249)
}
__typedArg0 := let__14249
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16109 Obj

if True == tmp16176 {
tmp16110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14249)
}
__typedArg0 := let__14249
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16111 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16110, V5342)


let__14250 := tmp16111
_ = let__14250

tmp16175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14250, symcons)
}
__typedArg0 := let__14250
__typedArg1 := symcons
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16112 Obj

if True == tmp16175 {
tmp16113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14249)
}
__typedArg0 := let__14249
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16114 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16113, V5342)


let__14251 := tmp16114
_ = let__14251

tmp16174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14251)
}
__typedArg0 := let__14251
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16115 Obj

if True == tmp16174 {
tmp16116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14251)
}
__typedArg0 := let__14251
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14252 := tmp16116
_ = let__14252

tmp16117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14251)
}
__typedArg0 := let__14251
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16118 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16117, V5342)


let__14253 := tmp16118
_ = let__14253

tmp16173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14253)
}
__typedArg0 := let__14253
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16119 Obj

if True == tmp16173 {
tmp16120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14253)
}
__typedArg0 := let__14253
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14254 := tmp16120
_ = let__14254

tmp16121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14253)
}
__typedArg0 := let__14253
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16122 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16121, V5342)


let__14255 := tmp16122
_ = let__14255

tmp16172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14255, Nil)
}
__typedArg0 := let__14255
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16123 Obj

if True == tmp16172 {
tmp16124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14248)
}
__typedArg0 := let__14248
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16125 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16124, V5342)


let__14256 := tmp16125
_ = let__14256

tmp16171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14256)
}
__typedArg0 := let__14256
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16126 Obj

if True == tmp16171 {
tmp16127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14256)
}
__typedArg0 := let__14256
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14257 := tmp16127
_ = let__14257

tmp16128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14256)
}
__typedArg0 := let__14256
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16129 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16128, V5342)


let__14258 := tmp16129
_ = let__14258

tmp16170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14258)
}
__typedArg0 := let__14258
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16130 Obj

if True == tmp16170 {
tmp16131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14258)
}
__typedArg0 := let__14258
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16132 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16131, V5342)


let__14259 := tmp16132
_ = let__14259

tmp16169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14259)
}
__typedArg0 := let__14259
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16133 Obj

if True == tmp16169 {
tmp16134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14259)
}
__typedArg0 := let__14259
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16135 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16134, V5342)


let__14260 := tmp16135
_ = let__14260

tmp16168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14260, symlist)
}
__typedArg0 := let__14260
__typedArg1 := symlist
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16136 Obj

if True == tmp16168 {
tmp16137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14259)
}
__typedArg0 := let__14259
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16138 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16137, V5342)


let__14261 := tmp16138
_ = let__14261

tmp16167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14261)
}
__typedArg0 := let__14261
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16139 Obj

if True == tmp16167 {
tmp16140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14261)
}
__typedArg0 := let__14261
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14262 := tmp16140
_ = let__14262

tmp16141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14261)
}
__typedArg0 := let__14261
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16142 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16141, V5342)


let__14263 := tmp16142
_ = let__14263

tmp16166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14263, Nil)
}
__typedArg0 := let__14263
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16143 Obj

if True == tmp16166 {
tmp16144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14258)
}
__typedArg0 := let__14258
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16145 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16144, V5342)


let__14264 := tmp16145
_ = let__14264

tmp16165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14264, Nil)
}
__typedArg0 := let__14264
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16146 Obj

if True == tmp16165 {
tmp16147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14247)
}
__typedArg0 := let__14247
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14265 := tmp16147
_ = let__14265

tmp16148 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16148

tmp16149 := Call(__e, PrimFunc(symshen_4deref), let__14257, V5342)


tmp16150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16149, tmp16150)
}
__typedArg0 := tmp16149
__typedArg1 := tmp16150
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16152 := MakeNative(func(__e *ControlFlow) {
tmp16153 := MakeNative(func(__e *ControlFlow) {
tmp16154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14262, Nil)
}
__typedArg0 := let__14262
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14257, tmp16154)
}
__typedArg0 := let__14257
__typedArg1 := tmp16154
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14252, tmp16155)
}
__typedArg0 := let__14252
__typedArg1 := tmp16155
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14262, Nil)
}
__typedArg0 := let__14262
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp16157)
}
__typedArg0 := symlist
__typedArg1 := tmp16157
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16158, Nil)
}
__typedArg0 := tmp16158
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14257, tmp16159)
}
__typedArg0 := let__14257
__typedArg1 := tmp16159
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14254, tmp16160)
}
__typedArg0 := let__14254
__typedArg1 := tmp16160
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16161, let__14265)
}
__typedArg0 := tmp16161
__typedArg1 := let__14265
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16156, tmp16162)
}
__typedArg0 := tmp16156
__typedArg1 := tmp16162
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16163, V5340, True, V5342, V5343, let__14242, V5345)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5342, V5343, let__14242, tmp16153)
return


}, 0)

tmp16164 := Call(__e, PrimFunc(symwhen), tmp16151, V5342, V5343, let__14242, tmp16152)


ifres16146 = tmp16164


} else {
ifres16146 = False


}

ifres16143 = ifres16146


} else {
ifres16143 = False


}

ifres16139 = ifres16143


} else {
ifres16139 = False


}

ifres16136 = ifres16139


} else {
ifres16136 = False


}

ifres16133 = ifres16136


} else {
ifres16133 = False


}

ifres16130 = ifres16133


} else {
ifres16130 = False


}

ifres16126 = ifres16130


} else {
ifres16126 = False


}

ifres16123 = ifres16126


} else {
ifres16123 = False


}

ifres16119 = ifres16123


} else {
ifres16119 = False


}

ifres16115 = ifres16119


} else {
ifres16115 = False


}

ifres16112 = ifres16115


} else {
ifres16112 = False


}

ifres16109 = ifres16112


} else {
ifres16109 = False


}

ifres16106 = ifres16109


} else {
ifres16106 = False


}

ifres16103 = ifres16106


} else {
ifres16103 = False


}

ifres16101 = ifres16103


} else {
ifres16101 = False


}

let__14246 := ifres16101
_ = let__14246

tmp16441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14246, False)
}
__typedArg0 := let__14246
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16441 {
tmp16261 := Call(__e, PrimFunc(symshen_4unlocked_2), V5343)


var ifres16180 Obj

if True == tmp16261 {
tmp16181 := Call(__e, PrimFunc(symshen_4lazyderef), V5339, V5342)


let__14267 := tmp16181
_ = let__14267

tmp16260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14267)
}
__typedArg0 := let__14267
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16182 Obj

if True == tmp16260 {
tmp16183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14267)
}
__typedArg0 := let__14267
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16184 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16183, V5342)


let__14268 := tmp16184
_ = let__14268

tmp16259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14268)
}
__typedArg0 := let__14268
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16185 Obj

if True == tmp16259 {
tmp16186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14268)
}
__typedArg0 := let__14268
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16187 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16186, V5342)


let__14269 := tmp16187
_ = let__14269

tmp16258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14269)
}
__typedArg0 := let__14269
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16188 Obj

if True == tmp16258 {
tmp16189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14269)
}
__typedArg0 := let__14269
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16190 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16189, V5342)


let__14270 := tmp16190
_ = let__14270

tmp16257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14270, sym_8p)
}
__typedArg0 := let__14270
__typedArg1 := sym_8p
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16191 Obj

if True == tmp16257 {
tmp16192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14269)
}
__typedArg0 := let__14269
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16193 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16192, V5342)


let__14271 := tmp16193
_ = let__14271

tmp16256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14271)
}
__typedArg0 := let__14271
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16194 Obj

if True == tmp16256 {
tmp16195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14271)
}
__typedArg0 := let__14271
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14272 := tmp16195
_ = let__14272

tmp16196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14271)
}
__typedArg0 := let__14271
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16197 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16196, V5342)


let__14273 := tmp16197
_ = let__14273

tmp16255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14273)
}
__typedArg0 := let__14273
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16198 Obj

if True == tmp16255 {
tmp16199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14273)
}
__typedArg0 := let__14273
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14274 := tmp16199
_ = let__14274

tmp16200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14273)
}
__typedArg0 := let__14273
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16201 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16200, V5342)


let__14275 := tmp16201
_ = let__14275

tmp16254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14275, Nil)
}
__typedArg0 := let__14275
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16202 Obj

if True == tmp16254 {
tmp16203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14268)
}
__typedArg0 := let__14268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16204 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16203, V5342)


let__14276 := tmp16204
_ = let__14276

tmp16253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14276)
}
__typedArg0 := let__14276
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16205 Obj

if True == tmp16253 {
tmp16206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14276)
}
__typedArg0 := let__14276
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14277 := tmp16206
_ = let__14277

tmp16207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14276)
}
__typedArg0 := let__14276
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16208 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16207, V5342)


let__14278 := tmp16208
_ = let__14278

tmp16252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14278)
}
__typedArg0 := let__14278
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16209 Obj

if True == tmp16252 {
tmp16210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14278)
}
__typedArg0 := let__14278
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16211 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16210, V5342)


let__14279 := tmp16211
_ = let__14279

tmp16251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14279)
}
__typedArg0 := let__14279
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16212 Obj

if True == tmp16251 {
tmp16213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14279)
}
__typedArg0 := let__14279
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14280 := tmp16213
_ = let__14280

tmp16214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14279)
}
__typedArg0 := let__14279
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16215 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16214, V5342)


let__14281 := tmp16215
_ = let__14281

tmp16250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14281)
}
__typedArg0 := let__14281
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16216 Obj

if True == tmp16250 {
tmp16217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14281)
}
__typedArg0 := let__14281
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16218 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16217, V5342)


let__14282 := tmp16218
_ = let__14282

tmp16249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14282, sym_d)
}
__typedArg0 := let__14282
__typedArg1 := sym_d
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16219 Obj

if True == tmp16249 {
tmp16220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14281)
}
__typedArg0 := let__14281
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16221 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16220, V5342)


let__14283 := tmp16221
_ = let__14283

tmp16248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14283)
}
__typedArg0 := let__14283
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16222 Obj

if True == tmp16248 {
tmp16223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14283)
}
__typedArg0 := let__14283
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14284 := tmp16223
_ = let__14284

tmp16224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14283)
}
__typedArg0 := let__14283
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16225 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16224, V5342)


let__14285 := tmp16225
_ = let__14285

tmp16247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14285, Nil)
}
__typedArg0 := let__14285
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16226 Obj

if True == tmp16247 {
tmp16227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14278)
}
__typedArg0 := let__14278
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16228 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16227, V5342)


let__14286 := tmp16228
_ = let__14286

tmp16246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14286, Nil)
}
__typedArg0 := let__14286
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16229 Obj

if True == tmp16246 {
tmp16230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14267)
}
__typedArg0 := let__14267
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14287 := tmp16230
_ = let__14287

tmp16231 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16231

tmp16232 := Call(__e, PrimFunc(symshen_4deref), let__14277, V5342)


tmp16233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16232, tmp16233)
}
__typedArg0 := tmp16232
__typedArg1 := tmp16233
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16235 := MakeNative(func(__e *ControlFlow) {
tmp16236 := MakeNative(func(__e *ControlFlow) {
tmp16237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14280, Nil)
}
__typedArg0 := let__14280
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14277, tmp16237)
}
__typedArg0 := let__14277
__typedArg1 := tmp16237
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14272, tmp16238)
}
__typedArg0 := let__14272
__typedArg1 := tmp16238
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14284, Nil)
}
__typedArg0 := let__14284
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14277, tmp16240)
}
__typedArg0 := let__14277
__typedArg1 := tmp16240
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14274, tmp16241)
}
__typedArg0 := let__14274
__typedArg1 := tmp16241
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16242, let__14287)
}
__typedArg0 := tmp16242
__typedArg1 := let__14287
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16239, tmp16243)
}
__typedArg0 := tmp16239
__typedArg1 := tmp16243
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16244, V5340, True, V5342, V5343, let__14242, V5345)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5342, V5343, let__14242, tmp16236)
return


}, 0)

tmp16245 := Call(__e, PrimFunc(symwhen), tmp16234, V5342, V5343, let__14242, tmp16235)


ifres16229 = tmp16245


} else {
ifres16229 = False


}

ifres16226 = ifres16229


} else {
ifres16226 = False


}

ifres16222 = ifres16226


} else {
ifres16222 = False


}

ifres16219 = ifres16222


} else {
ifres16219 = False


}

ifres16216 = ifres16219


} else {
ifres16216 = False


}

ifres16212 = ifres16216


} else {
ifres16212 = False


}

ifres16209 = ifres16212


} else {
ifres16209 = False


}

ifres16205 = ifres16209


} else {
ifres16205 = False


}

ifres16202 = ifres16205


} else {
ifres16202 = False


}

ifres16198 = ifres16202


} else {
ifres16198 = False


}

ifres16194 = ifres16198


} else {
ifres16194 = False


}

ifres16191 = ifres16194


} else {
ifres16191 = False


}

ifres16188 = ifres16191


} else {
ifres16188 = False


}

ifres16185 = ifres16188


} else {
ifres16185 = False


}

ifres16182 = ifres16185


} else {
ifres16182 = False


}

ifres16180 = ifres16182


} else {
ifres16180 = False


}

let__14266 := ifres16180
_ = let__14266

tmp16439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14266, False)
}
__typedArg0 := let__14266
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16439 {
tmp16325 := Call(__e, PrimFunc(symshen_4unlocked_2), V5343)


var ifres16262 Obj

if True == tmp16325 {
tmp16263 := Call(__e, PrimFunc(symshen_4lazyderef), V5339, V5342)


let__14289 := tmp16263
_ = let__14289

tmp16324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14289)
}
__typedArg0 := let__14289
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16264 Obj

if True == tmp16324 {
tmp16265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14289)
}
__typedArg0 := let__14289
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16266 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16265, V5342)


let__14290 := tmp16266
_ = let__14290

tmp16323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14290)
}
__typedArg0 := let__14290
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16267 Obj

if True == tmp16323 {
tmp16268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14290)
}
__typedArg0 := let__14290
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16269 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16268, V5342)


let__14291 := tmp16269
_ = let__14291

tmp16322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14291)
}
__typedArg0 := let__14291
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16270 Obj

if True == tmp16322 {
tmp16271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14291)
}
__typedArg0 := let__14291
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16272 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16271, V5342)


let__14292 := tmp16272
_ = let__14292

tmp16321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14292, sym_8s)
}
__typedArg0 := let__14292
__typedArg1 := sym_8s
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16273 Obj

if True == tmp16321 {
tmp16274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14291)
}
__typedArg0 := let__14291
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16275 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16274, V5342)


let__14293 := tmp16275
_ = let__14293

tmp16320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14293)
}
__typedArg0 := let__14293
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16276 Obj

if True == tmp16320 {
tmp16277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14293)
}
__typedArg0 := let__14293
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14294 := tmp16277
_ = let__14294

tmp16278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14293)
}
__typedArg0 := let__14293
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16279 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16278, V5342)


let__14295 := tmp16279
_ = let__14295

tmp16319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14295)
}
__typedArg0 := let__14295
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16280 Obj

if True == tmp16319 {
tmp16281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14295)
}
__typedArg0 := let__14295
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14296 := tmp16281
_ = let__14296

tmp16282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14295)
}
__typedArg0 := let__14295
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16283 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16282, V5342)


let__14297 := tmp16283
_ = let__14297

tmp16318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14297, Nil)
}
__typedArg0 := let__14297
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16284 Obj

if True == tmp16318 {
tmp16285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14290)
}
__typedArg0 := let__14290
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16286 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16285, V5342)


let__14298 := tmp16286
_ = let__14298

tmp16317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14298)
}
__typedArg0 := let__14298
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16287 Obj

if True == tmp16317 {
tmp16288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14298)
}
__typedArg0 := let__14298
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14299 := tmp16288
_ = let__14299

tmp16289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14298)
}
__typedArg0 := let__14298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16290 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16289, V5342)


let__14300 := tmp16290
_ = let__14300

tmp16316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14300)
}
__typedArg0 := let__14300
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16291 Obj

if True == tmp16316 {
tmp16292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14300)
}
__typedArg0 := let__14300
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16293 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16292, V5342)


let__14301 := tmp16293
_ = let__14301

tmp16315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14301, symstring)
}
__typedArg0 := let__14301
__typedArg1 := symstring
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16294 Obj

if True == tmp16315 {
tmp16295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14300)
}
__typedArg0 := let__14300
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16296 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16295, V5342)


let__14302 := tmp16296
_ = let__14302

tmp16314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14302, Nil)
}
__typedArg0 := let__14302
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16297 Obj

if True == tmp16314 {
tmp16298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14289)
}
__typedArg0 := let__14289
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14303 := tmp16298
_ = let__14303

tmp16299 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16299

tmp16300 := Call(__e, PrimFunc(symshen_4deref), let__14299, V5342)


tmp16301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16300, tmp16301)
}
__typedArg0 := tmp16300
__typedArg1 := tmp16301
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16303 := MakeNative(func(__e *ControlFlow) {
tmp16304 := MakeNative(func(__e *ControlFlow) {
tmp16305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14299, tmp16305)
}
__typedArg0 := let__14299
__typedArg1 := tmp16305
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14294, tmp16306)
}
__typedArg0 := let__14294
__typedArg1 := tmp16306
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14299, tmp16308)
}
__typedArg0 := let__14299
__typedArg1 := tmp16308
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14296, tmp16309)
}
__typedArg0 := let__14296
__typedArg1 := tmp16309
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16310, let__14303)
}
__typedArg0 := tmp16310
__typedArg1 := let__14303
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16307, tmp16311)
}
__typedArg0 := tmp16307
__typedArg1 := tmp16311
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16312, V5340, True, V5342, V5343, let__14242, V5345)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5342, V5343, let__14242, tmp16304)
return


}, 0)

tmp16313 := Call(__e, PrimFunc(symwhen), tmp16302, V5342, V5343, let__14242, tmp16303)


ifres16297 = tmp16313


} else {
ifres16297 = False


}

ifres16294 = ifres16297


} else {
ifres16294 = False


}

ifres16291 = ifres16294


} else {
ifres16291 = False


}

ifres16287 = ifres16291


} else {
ifres16287 = False


}

ifres16284 = ifres16287


} else {
ifres16284 = False


}

ifres16280 = ifres16284


} else {
ifres16280 = False


}

ifres16276 = ifres16280


} else {
ifres16276 = False


}

ifres16273 = ifres16276


} else {
ifres16273 = False


}

ifres16270 = ifres16273


} else {
ifres16270 = False


}

ifres16267 = ifres16270


} else {
ifres16267 = False


}

ifres16264 = ifres16267


} else {
ifres16264 = False


}

ifres16262 = ifres16264


} else {
ifres16262 = False


}

let__14288 := ifres16262
_ = let__14288

tmp16437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14288, False)
}
__typedArg0 := let__14288
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16437 {
tmp16404 := Call(__e, PrimFunc(symshen_4unlocked_2), V5343)


var ifres16326 Obj

if True == tmp16404 {
tmp16327 := Call(__e, PrimFunc(symshen_4lazyderef), V5339, V5342)


let__14305 := tmp16327
_ = let__14305

tmp16403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14305)
}
__typedArg0 := let__14305
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16328 Obj

if True == tmp16403 {
tmp16329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14305)
}
__typedArg0 := let__14305
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16330 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16329, V5342)


let__14306 := tmp16330
_ = let__14306

tmp16402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14306)
}
__typedArg0 := let__14306
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16331 Obj

if True == tmp16402 {
tmp16332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14306)
}
__typedArg0 := let__14306
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16333 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16332, V5342)


let__14307 := tmp16333
_ = let__14307

tmp16401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14307)
}
__typedArg0 := let__14307
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16334 Obj

if True == tmp16401 {
tmp16335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14307)
}
__typedArg0 := let__14307
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16336 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16335, V5342)


let__14308 := tmp16336
_ = let__14308

tmp16400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14308, sym_8v)
}
__typedArg0 := let__14308
__typedArg1 := sym_8v
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16337 Obj

if True == tmp16400 {
tmp16338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14307)
}
__typedArg0 := let__14307
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16339 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16338, V5342)


let__14309 := tmp16339
_ = let__14309

tmp16399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14309)
}
__typedArg0 := let__14309
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16340 Obj

if True == tmp16399 {
tmp16341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14309)
}
__typedArg0 := let__14309
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14310 := tmp16341
_ = let__14310

tmp16342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14309)
}
__typedArg0 := let__14309
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16343 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16342, V5342)


let__14311 := tmp16343
_ = let__14311

tmp16398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14311)
}
__typedArg0 := let__14311
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16344 Obj

if True == tmp16398 {
tmp16345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14311)
}
__typedArg0 := let__14311
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14312 := tmp16345
_ = let__14312

tmp16346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14311)
}
__typedArg0 := let__14311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16347 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16346, V5342)


let__14313 := tmp16347
_ = let__14313

tmp16397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14313, Nil)
}
__typedArg0 := let__14313
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16348 Obj

if True == tmp16397 {
tmp16349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14306)
}
__typedArg0 := let__14306
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16350 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16349, V5342)


let__14314 := tmp16350
_ = let__14314

tmp16396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14314)
}
__typedArg0 := let__14314
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16351 Obj

if True == tmp16396 {
tmp16352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14314)
}
__typedArg0 := let__14314
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14315 := tmp16352
_ = let__14315

tmp16353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14314)
}
__typedArg0 := let__14314
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16354 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16353, V5342)


let__14316 := tmp16354
_ = let__14316

tmp16395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14316)
}
__typedArg0 := let__14316
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16355 Obj

if True == tmp16395 {
tmp16356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14316)
}
__typedArg0 := let__14316
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16357 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16356, V5342)


let__14317 := tmp16357
_ = let__14317

tmp16394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14317)
}
__typedArg0 := let__14317
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16358 Obj

if True == tmp16394 {
tmp16359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14317)
}
__typedArg0 := let__14317
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16360 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16359, V5342)


let__14318 := tmp16360
_ = let__14318

tmp16393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14318, symvector)
}
__typedArg0 := let__14318
__typedArg1 := symvector
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16361 Obj

if True == tmp16393 {
tmp16362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14317)
}
__typedArg0 := let__14317
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16363 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16362, V5342)


let__14319 := tmp16363
_ = let__14319

tmp16392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14319)
}
__typedArg0 := let__14319
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16364 Obj

if True == tmp16392 {
tmp16365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14319)
}
__typedArg0 := let__14319
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14320 := tmp16365
_ = let__14320

tmp16366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14319)
}
__typedArg0 := let__14319
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16367 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16366, V5342)


let__14321 := tmp16367
_ = let__14321

tmp16391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14321, Nil)
}
__typedArg0 := let__14321
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16368 Obj

if True == tmp16391 {
tmp16369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14316)
}
__typedArg0 := let__14316
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16370 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16369, V5342)


let__14322 := tmp16370
_ = let__14322

tmp16390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14322, Nil)
}
__typedArg0 := let__14322
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16371 Obj

if True == tmp16390 {
tmp16372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14305)
}
__typedArg0 := let__14305
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14323 := tmp16372
_ = let__14323

tmp16373 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16373

tmp16374 := Call(__e, PrimFunc(symshen_4deref), let__14315, V5342)


tmp16375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16374, tmp16375)
}
__typedArg0 := tmp16374
__typedArg1 := tmp16375
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16377 := MakeNative(func(__e *ControlFlow) {
tmp16378 := MakeNative(func(__e *ControlFlow) {
tmp16379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14320, Nil)
}
__typedArg0 := let__14320
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14315, tmp16379)
}
__typedArg0 := let__14315
__typedArg1 := tmp16379
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14310, tmp16380)
}
__typedArg0 := let__14310
__typedArg1 := tmp16380
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14320, Nil)
}
__typedArg0 := let__14320
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp16382)
}
__typedArg0 := symvector
__typedArg1 := tmp16382
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16383, Nil)
}
__typedArg0 := tmp16383
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14315, tmp16384)
}
__typedArg0 := let__14315
__typedArg1 := tmp16384
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14312, tmp16385)
}
__typedArg0 := let__14312
__typedArg1 := tmp16385
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16386, let__14323)
}
__typedArg0 := tmp16386
__typedArg1 := let__14323
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16381, tmp16387)
}
__typedArg0 := tmp16381
__typedArg1 := tmp16387
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16388, V5340, True, V5342, V5343, let__14242, V5345)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5342, V5343, let__14242, tmp16378)
return


}, 0)

tmp16389 := Call(__e, PrimFunc(symwhen), tmp16376, V5342, V5343, let__14242, tmp16377)


ifres16371 = tmp16389


} else {
ifres16371 = False


}

ifres16368 = ifres16371


} else {
ifres16368 = False


}

ifres16364 = ifres16368


} else {
ifres16364 = False


}

ifres16361 = ifres16364


} else {
ifres16361 = False


}

ifres16358 = ifres16361


} else {
ifres16358 = False


}

ifres16355 = ifres16358


} else {
ifres16355 = False


}

ifres16351 = ifres16355


} else {
ifres16351 = False


}

ifres16348 = ifres16351


} else {
ifres16348 = False


}

ifres16344 = ifres16348


} else {
ifres16344 = False


}

ifres16340 = ifres16344


} else {
ifres16340 = False


}

ifres16337 = ifres16340


} else {
ifres16337 = False


}

ifres16334 = ifres16337


} else {
ifres16334 = False


}

ifres16331 = ifres16334


} else {
ifres16331 = False


}

ifres16328 = ifres16331


} else {
ifres16328 = False


}

ifres16326 = ifres16328


} else {
ifres16326 = False


}

let__14304 := ifres16326
_ = let__14304

tmp16435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14304, False)
}
__typedArg0 := let__14304
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16435 {
tmp16431 := Call(__e, PrimFunc(symshen_4unlocked_2), V5343)


var ifres16405 Obj

if True == tmp16431 {
tmp16406 := Call(__e, PrimFunc(symshen_4lazyderef), V5339, V5342)


let__14325 := tmp16406
_ = let__14325

tmp16430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14325)
}
__typedArg0 := let__14325
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16407 Obj

if True == tmp16430 {
tmp16408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14325)
}
__typedArg0 := let__14325
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14326 := tmp16408
_ = let__14326

tmp16409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14325)
}
__typedArg0 := let__14325
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14327 := tmp16409
_ = let__14327

tmp16410 := Call(__e, PrimFunc(symshen_4lazyderef), V5340, V5342)


let__14328 := tmp16410
_ = let__14328

tmp16411 := MakeNative(func(__e *ControlFlow) {
Z5434 := __e.Get(1)
_ = Z5434
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5435 := __e.Get(1)
_ = Z5435
tmp16412 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16412

tmp16413 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4l_1rules), let__14327, Z5435, V5341, V5342, V5343, let__14242, V5345)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5434, let__14326, V5342, V5343, let__14242, tmp16413)
return


}, 1))
return
}, 1)

let__14329 := tmp16411
_ = let__14329

tmp16429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14328)
}
__typedArg0 := let__14328
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16414 Obj

if True == tmp16429 {
tmp16415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14328)
}
__typedArg0 := let__14328
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14330 := tmp16415
_ = let__14330

tmp16416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14328)
}
__typedArg0 := let__14328
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14331 := tmp16416
_ = let__14331

tmp16417 := Call(__e, let__14329, let__14330)


tmp16418 := Call(__e, tmp16417, let__14331)


ifres16414 = tmp16418


} else {
tmp16428 := Call(__e, PrimFunc(symshen_4pvar_2), let__14328)


var ifres16419 Obj

if True == tmp16428 {
tmp16420 := Call(__e, PrimFunc(symshen_4newpv), V5342)


let__14332 := tmp16420
_ = let__14332

tmp16421 := Call(__e, PrimFunc(symshen_4newpv), V5342)


let__14333 := tmp16421
_ = let__14333

tmp16422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14332, let__14333)
}
__typedArg0 := let__14332
__typedArg1 := let__14333
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16423 := MakeNative(func(__e *ControlFlow) {
tmp16424 := Call(__e, let__14329, let__14332)


__e.TailApply(tmp16424, let__14333)
return


}, 0)

tmp16425 := Call(__e, PrimFunc(symshen_4bind_b), let__14328, tmp16422, V5342, tmp16423)


tmp16426 := Call(__e, PrimFunc(symshen_4gc), V5342, tmp16425)


tmp16427 := Call(__e, PrimFunc(symshen_4gc), V5342, tmp16426)


ifres16419 = tmp16427


} else {
ifres16419 = False


}

ifres16414 = ifres16419


}

ifres16407 = ifres16414


} else {
ifres16407 = False


}

ifres16405 = ifres16407


} else {
ifres16405 = False


}

let__14324 := ifres16405
_ = let__14324

tmp16433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14324, False)
}
__typedArg0 := let__14324
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16433 {
__e.TailApply(PrimFunc(symshen_4unlock), V5343, let__14242)
return
} else {
__e.Return(let__14324)
return
}


} else {
__e.Return(let__14304)
return
}


} else {
__e.Return(let__14288)
return
}


} else {
__e.Return(let__14266)
return
}


} else {
__e.Return(let__14246)
return
}


} else {
__e.Return(let__14243)
return
}


}, 7)

tmp16444 := Call(__e, ns2_1set, symshen_4l_1rules, tmp16088)


_ = tmp16444

tmp16445 := MakeNative(func(__e *ControlFlow) {
V5440 := __e.Get(1)
_ = V5440
V5441 := __e.Get(2)
_ = V5441
V5442 := __e.Get(3)
_ = V5442
V5443 := __e.Get(4)
_ = V5443
V5444 := __e.Get(5)
_ = V5444
V5445 := __e.Get(6)
_ = V5445
let__14334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5444)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5444
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__14334

tmp16485 := Call(__e, PrimFunc(symshen_4unlocked_2), V5443)


var ifres16447 Obj

if True == tmp16485 {
tmp16448 := Call(__e, PrimFunc(symshen_4lazyderef), V5440, V5442)


let__14336 := tmp16448
_ = let__14336

tmp16484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14336)
}
__typedArg0 := let__14336
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16449 Obj

if True == tmp16484 {
tmp16450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14336)
}
__typedArg0 := let__14336
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16451 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16450, V5442)


let__14337 := tmp16451
_ = let__14337

tmp16483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14337, symdefine)
}
__typedArg0 := let__14337
__typedArg1 := symdefine
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16452 Obj

if True == tmp16483 {
tmp16453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14336)
}
__typedArg0 := let__14336
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16454 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16453, V5442)


let__14338 := tmp16454
_ = let__14338

tmp16482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14338)
}
__typedArg0 := let__14338
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16455 Obj

if True == tmp16482 {
tmp16456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14338)
}
__typedArg0 := let__14338
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14339 := tmp16456
_ = let__14339

tmp16457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14338)
}
__typedArg0 := let__14338
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14340 := tmp16457
_ = let__14340

tmp16458 := Call(__e, PrimFunc(symshen_4newpv), V5442)


let__14341 := tmp16458
_ = let__14341

tmp16459 := Call(__e, PrimFunc(symshen_4newpv), V5442)


let__14342 := tmp16459
_ = let__14342

tmp16460 := Call(__e, PrimFunc(symshen_4newpv), V5442)


let__14343 := tmp16460
_ = let__14343

tmp16461 := Call(__e, PrimFunc(symshen_4newpv), V5442)


let__14344 := tmp16461
_ = let__14344

tmp16462 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16462

tmp16463 := MakeNative(func(__e *ControlFlow) {
tmp16464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14339, let__14340)
}
__typedArg0 := let__14339
__typedArg1 := let__14340
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16465 := Call(__e, PrimFunc(symshen_4sigxrules), tmp16464)


tmp16466 := MakeNative(func(__e *ControlFlow) {
tmp16467 := Call(__e, PrimFunc(symshen_4lazyderef), let__14341, V5442)


tmp16468 := Call(__e, PrimFunc(symfst), tmp16467)


tmp16469 := MakeNative(func(__e *ControlFlow) {
tmp16470 := Call(__e, PrimFunc(symshen_4lazyderef), let__14341, V5442)


tmp16471 := Call(__e, PrimFunc(symsnd), tmp16470)


tmp16472 := MakeNative(func(__e *ControlFlow) {
tmp16473 := Call(__e, PrimFunc(symshen_4deref), let__14344, V5442)


tmp16474 := Call(__e, PrimFunc(symshen_4freshen_1sig), tmp16473)


tmp16475 := MakeNative(func(__e *ControlFlow) {
tmp16476 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symis), let__14344, V5441, V5442, V5443, let__14334, V5445)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1rules), let__14339, let__14342, let__14343, MakeInteger(1), V5442, V5443, let__14334, tmp16476)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14343, tmp16474, V5442, V5443, let__14334, tmp16475)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14342, tmp16471, V5442, V5443, let__14334, tmp16472)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14344, tmp16468, V5442, V5443, let__14334, tmp16469)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14341, tmp16465, V5442, V5443, let__14334, tmp16466)
return


}, 0)

tmp16477 := Call(__e, PrimFunc(symshen_4cut), V5442, V5443, let__14334, tmp16463)


tmp16478 := Call(__e, PrimFunc(symshen_4gc), V5442, tmp16477)


tmp16479 := Call(__e, PrimFunc(symshen_4gc), V5442, tmp16478)


tmp16480 := Call(__e, PrimFunc(symshen_4gc), V5442, tmp16479)


tmp16481 := Call(__e, PrimFunc(symshen_4gc), V5442, tmp16480)


ifres16455 = tmp16481


} else {
ifres16455 = False


}

ifres16452 = ifres16455


} else {
ifres16452 = False


}

ifres16449 = ifres16452


} else {
ifres16449 = False


}

ifres16447 = ifres16449


} else {
ifres16447 = False


}

let__14335 := ifres16447
_ = let__14335

tmp16487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14335, False)
}
__typedArg0 := let__14335
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16487 {
__e.TailApply(PrimFunc(symshen_4unlock), V5443, let__14334)
return
} else {
__e.Return(let__14335)
return
}


}, 6)

tmp16488 := Call(__e, ns2_1set, symshen_4t_d, tmp16445)


_ = tmp16488

tmp16489 := MakeNative(func(__e *ControlFlow) {
V5457 := __e.Get(1)
_ = V5457
tmp16490 := MakeNative(func(__e *ControlFlow) {
Z5459 := __e.Get(1)
_ = Z5459
__e.TailApply(PrimFunc(symshen_4_5sig_drules_6), Z5459)
return
}, 1)

tmp16491 := Call(__e, PrimFunc(symcompile), tmp16490, V5457)


let__14345 := tmp16491
_ = let__14345

tmp16492 := Call(__e, PrimFunc(symfst), let__14345)


tmp16493 := MakeNative(func(__e *ControlFlow) {
Z5460 := __e.Get(1)
_ = Z5460
__e.TailApply(PrimFunc(symshen_4unmark_1rule_1custom_1patterns), Z5460)
return
}, 1)

tmp16494 := Call(__e, PrimFunc(symsnd), let__14345)


tmp16495 := Call(__e, PrimFunc(symmap), tmp16493, tmp16494)


__e.TailApply(PrimFunc(sym_8p), tmp16492, tmp16495)
return


}, 1)

tmp16496 := Call(__e, ns2_1set, symshen_4sigxrules, tmp16489)


_ = tmp16496

tmp16497 := MakeNative(func(__e *ControlFlow) {
V5461 := __e.Get(1)
_ = V5461
tmp16503 := Call(__e, PrimFunc(symtuple_2), V5461)


if True == tmp16503 {
tmp16498 := MakeNative(func(__e *ControlFlow) {
Z5462 := __e.Get(1)
_ = Z5462
__e.TailApply(PrimFunc(symshen_4unmark_1custom_1pattern), Z5462)
return
}, 1)

tmp16499 := Call(__e, PrimFunc(symfst), V5461)


tmp16500 := Call(__e, PrimFunc(symmap), tmp16498, tmp16499)


tmp16501 := Call(__e, PrimFunc(symsnd), V5461)


__e.TailApply(PrimFunc(sym_8p), tmp16500, tmp16501)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4unmark_1rule_1custom_1patterns)
return
}


}, 1)

tmp16504 := Call(__e, ns2_1set, symshen_4unmark_1rule_1custom_1patterns, tmp16497)


_ = tmp16504

tmp16505 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V5463 := __self1
_ = V5463
tmp16541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5463)
}
__typedArg0 := V5463
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16517 Obj

if True == tmp16541 {
tmp16539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5463)
}
__typedArg0 := V5463
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8p, tmp16539)
}
__typedArg0 := sym_8p
__typedArg1 := tmp16539
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16519 Obj

if True == tmp16540 {
tmp16537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5463)
}
__typedArg0 := V5463
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16537)
}
__typedArg0 := tmp16537
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16521 Obj

if True == tmp16538 {
tmp16534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5463)
}
__typedArg0 := V5463
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16534)
}
__typedArg0 := tmp16534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4custom_1pattern, tmp16535)
}
__typedArg0 := symshen_4custom_1pattern
__typedArg1 := tmp16535
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16523 Obj

if True == tmp16536 {
tmp16531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5463)
}
__typedArg0 := V5463
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16531)
}
__typedArg0 := tmp16531
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16532)
}
__typedArg0 := tmp16532
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16525 Obj

if True == tmp16533 {
tmp16527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5463)
}
__typedArg0 := V5463
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16527)
}
__typedArg0 := tmp16527
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16528)
}
__typedArg0 := tmp16528
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp16529)
}
__typedArg0 := Nil
__typedArg1 := tmp16529
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16526 Obj

if True == tmp16530 {
ifres16526 = True


} else {
ifres16526 = False


}

ifres16525 = ifres16526


} else {
ifres16525 = False


}

var ifres16524 Obj

if True == ifres16525 {
ifres16524 = True


} else {
ifres16524 = False


}

ifres16523 = ifres16524


} else {
ifres16523 = False


}

var ifres16522 Obj

if True == ifres16523 {
ifres16522 = True


} else {
ifres16522 = False


}

ifres16521 = ifres16522


} else {
ifres16521 = False


}

var ifres16520 Obj

if True == ifres16521 {
ifres16520 = True


} else {
ifres16520 = False


}

ifres16519 = ifres16520


} else {
ifres16519 = False


}

var ifres16518 Obj

if True == ifres16519 {
ifres16518 = True


} else {
ifres16518 = False


}

ifres16517 = ifres16518


} else {
ifres16517 = False


}

if True == ifres16517 {
tmp16506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5463)
}
__typedArg0 := V5463
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16506)
}
__typedArg0 := tmp16506
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16507)
}
__typedArg0 := tmp16507
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16509 := Call(__e, PrimFunc(symshen_4curry), tmp16508)


if PrimFunc(symshen_4unmark_1custom_1pattern) == __self {
__self1 = tmp16509
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4unmark_1custom_1pattern), tmp16509)
return


} else {
tmp16515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5463)
}
__typedArg0 := V5463
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16515 {
tmp16510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5463)
}
__typedArg0 := V5463
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16511 := Call(__e, PrimFunc(symshen_4unmark_1custom_1pattern), tmp16510)


tmp16512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5463)
}
__typedArg0 := V5463
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16513 := Call(__e, PrimFunc(symshen_4unmark_1custom_1pattern), tmp16512)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16511, tmp16513)
}
__typedArg0 := tmp16511
__typedArg1 := tmp16513
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V5463)
return
}


}


}, 1)

tmp16542 := Call(__e, ns2_1set, symshen_4unmark_1custom_1pattern, tmp16505)


_ = tmp16542

tmp16543 := MakeNative(func(__e *ControlFlow) {
V5464 := __e.Get(1)
_ = V5464
tmp16570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5464)
}
__typedArg0 := V5464
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16544 Obj

if True == tmp16570 {
tmp16545 := Call(__e, PrimFunc(symtail), V5464)


let__14347 := tmp16545
_ = let__14347

tmp16568 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14347, sym_i)


var ifres16546 Obj

if True == tmp16568 {
tmp16547 := Call(__e, PrimFunc(symtail), let__14347)


let__14348 := tmp16547
_ = let__14348

tmp16548 := Call(__e, PrimFunc(symshen_4_5signature_6), let__14348)


let__14349 := tmp16548
_ = let__14349

tmp16566 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14349)


var ifres16549 Obj

if True == tmp16566 {
tmp16550 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16549 = tmp16550


} else {
tmp16551 := Call(__e, PrimFunc(symshen_4_5_1out), let__14349)


let__14350 := tmp16551
_ = let__14350

tmp16552 := Call(__e, PrimFunc(symshen_4in_1_6), let__14349)


let__14351 := tmp16552
_ = let__14351

tmp16565 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14351, sym_j)


var ifres16553 Obj

if True == tmp16565 {
tmp16554 := Call(__e, PrimFunc(symtail), let__14351)


let__14352 := tmp16554
_ = let__14352

tmp16555 := Call(__e, PrimFunc(symshen_4_5rules_d_6), let__14352)


let__14353 := tmp16555
_ = let__14353

tmp16563 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14353)


var ifres16556 Obj

if True == tmp16563 {
tmp16557 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16556 = tmp16557


} else {
tmp16558 := Call(__e, PrimFunc(symshen_4_5_1out), let__14353)


let__14354 := tmp16558
_ = let__14354

tmp16559 := Call(__e, PrimFunc(symshen_4in_1_6), let__14353)


let__14355 := tmp16559
_ = let__14355

tmp16560 := Call(__e, PrimFunc(symshen_4rectify_1type), let__14350)


let__14356 := tmp16560
_ = let__14356

tmp16561 := Call(__e, PrimFunc(sym_8p), let__14356, let__14354)


tmp16562 := Call(__e, PrimFunc(symshen_4comb), let__14355, tmp16561)


ifres16556 = tmp16562


}

ifres16553 = ifres16556


} else {
tmp16564 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16553 = tmp16564


}

ifres16549 = ifres16553


}

ifres16546 = ifres16549


} else {
tmp16567 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16546 = tmp16567


}

ifres16544 = ifres16546


} else {
tmp16569 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16544 = tmp16569


}

let__14346 := ifres16544
_ = let__14346

tmp16572 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14346)


if True == tmp16572 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__14346)
return
}


}, 1)

tmp16573 := Call(__e, ns2_1set, symshen_4_5sig_drules_6, tmp16543)


_ = tmp16573

tmp16574 := MakeNative(func(__e *ControlFlow) {
V5476 := __e.Get(1)
_ = V5476
tmp16575 := Call(__e, PrimFunc(symshen_4extract_1vars), V5476)


let__14357 := tmp16575
_ = let__14357

tmp16576 := MakeNative(func(__e *ControlFlow) {
Z5479 := __e.Get(1)
_ = Z5479
tmp16577 := Call(__e, PrimFunc(symconcat), sym_e, Z5479)


tmp16578 := Call(__e, PrimFunc(symshen_4freshterm), tmp16577)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z5479, tmp16578)
}
__typedArg0 := Z5479
__typedArg1 := tmp16578
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp16579 := Call(__e, PrimFunc(symmap), tmp16576, let__14357)


let__14358 := tmp16579
_ = let__14358

__e.TailApply(PrimFunc(symshen_4freshen_1type), let__14358, V5476)
return


}, 1)

tmp16580 := Call(__e, ns2_1set, symshen_4freshen_1sig, tmp16574)


_ = tmp16580

tmp16581 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V5480 := __self1
_ = V5480
V5481 := __self2
_ = V5481
tmp16595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5480)
}
__typedArg0 := Nil
__typedArg1 := V5480
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16595 {
__e.Return(V5481)
return
} else {
tmp16593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5480)
}
__typedArg0 := V5480
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16589 Obj

if True == tmp16593 {
tmp16591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5480)
}
__typedArg0 := V5480
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16591)
}
__typedArg0 := tmp16591
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16590 Obj

if True == tmp16592 {
ifres16590 = True


} else {
ifres16590 = False


}

ifres16589 = ifres16590


} else {
ifres16589 = False


}

if True == ifres16589 {
tmp16582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5480)
}
__typedArg0 := V5480
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5480)
}
__typedArg0 := V5480
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16583)
}
__typedArg0 := tmp16583
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5480)
}
__typedArg0 := V5480
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16585)
}
__typedArg0 := tmp16585
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16587 := Call(__e, PrimFunc(symsubst), tmp16584, tmp16586, V5481)


if PrimFunc(symshen_4freshen_1type) == __self {
__self1, __self2 = tmp16582, tmp16587
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4freshen_1type), tmp16582, tmp16587)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshen_1type)
return
}


}


}, 2)

tmp16596 := Call(__e, ns2_1set, symshen_4freshen_1type, tmp16581)


_ = tmp16596

tmp16597 := MakeNative(func(__e *ControlFlow) {
V5482 := __e.Get(1)
_ = V5482
tmp16598 := Call(__e, PrimFunc(symshen_4_5rule_d_6), V5482)


let__14360 := tmp16598
_ = let__14360

tmp16611 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14360)


var ifres16599 Obj

if True == tmp16611 {
tmp16600 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16599 = tmp16600


} else {
tmp16601 := Call(__e, PrimFunc(symshen_4_5_1out), let__14360)


let__14361 := tmp16601
_ = let__14361

tmp16602 := Call(__e, PrimFunc(symshen_4in_1_6), let__14360)


let__14362 := tmp16602
_ = let__14362

tmp16603 := Call(__e, PrimFunc(symshen_4_5rules_d_6), let__14362)


let__14363 := tmp16603
_ = let__14363

tmp16610 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14363)


var ifres16604 Obj

if True == tmp16610 {
tmp16605 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16604 = tmp16605


} else {
tmp16606 := Call(__e, PrimFunc(symshen_4_5_1out), let__14363)


let__14364 := tmp16606
_ = let__14364

tmp16607 := Call(__e, PrimFunc(symshen_4in_1_6), let__14363)


let__14365 := tmp16607
_ = let__14365

tmp16608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14361, let__14364)
}
__typedArg0 := let__14361
__typedArg1 := let__14364
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16609 := Call(__e, PrimFunc(symshen_4comb), let__14365, tmp16608)


ifres16604 = tmp16609


}

ifres16599 = ifres16604


}

let__14359 := ifres16599
_ = let__14359

tmp16623 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14359)


if True == tmp16623 {
tmp16612 := Call(__e, PrimFunc(symshen_4_5rule_d_6), V5482)


let__14367 := tmp16612
_ = let__14367

tmp16619 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14367)


var ifres16613 Obj

if True == tmp16619 {
tmp16614 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16613 = tmp16614


} else {
tmp16615 := Call(__e, PrimFunc(symshen_4_5_1out), let__14367)


let__14368 := tmp16615
_ = let__14368

tmp16616 := Call(__e, PrimFunc(symshen_4in_1_6), let__14367)


let__14369 := tmp16616
_ = let__14369

tmp16617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14368, Nil)
}
__typedArg0 := let__14368
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16618 := Call(__e, PrimFunc(symshen_4comb), let__14369, tmp16617)


ifres16613 = tmp16618


}

let__14366 := ifres16613
_ = let__14366

tmp16621 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14366)


if True == tmp16621 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__14366)
return
}


} else {
__e.Return(let__14359)
return
}


}, 1)

tmp16624 := Call(__e, ns2_1set, symshen_4_5rules_d_6, tmp16597)


_ = tmp16624

tmp16625 := MakeNative(func(__e *ControlFlow) {
V5494 := __e.Get(1)
_ = V5494
tmp16626 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5494)


let__14371 := tmp16626
_ = let__14371

tmp16654 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14371)


var ifres16627 Obj

if True == tmp16654 {
tmp16628 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16627 = tmp16628


} else {
tmp16629 := Call(__e, PrimFunc(symshen_4_5_1out), let__14371)


let__14372 := tmp16629
_ = let__14372

tmp16630 := Call(__e, PrimFunc(symshen_4in_1_6), let__14371)


let__14373 := tmp16630
_ = let__14373

tmp16653 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14373, sym_1_6)


var ifres16631 Obj

if True == tmp16653 {
tmp16632 := Call(__e, PrimFunc(symtail), let__14373)


let__14374 := tmp16632
_ = let__14374

tmp16651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14374)
}
__typedArg0 := let__14374
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16633 Obj

if True == tmp16651 {
tmp16634 := Call(__e, PrimFunc(symhead), let__14374)


let__14375 := tmp16634
_ = let__14375

tmp16635 := Call(__e, PrimFunc(symtail), let__14374)


let__14376 := tmp16635
_ = let__14376

tmp16649 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14376, symwhere)


var ifres16636 Obj

if True == tmp16649 {
tmp16637 := Call(__e, PrimFunc(symtail), let__14376)


let__14377 := tmp16637
_ = let__14377

tmp16647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14377)
}
__typedArg0 := let__14377
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16638 Obj

if True == tmp16647 {
tmp16639 := Call(__e, PrimFunc(symhead), let__14377)


let__14378 := tmp16639
_ = let__14378

tmp16640 := Call(__e, PrimFunc(symtail), let__14377)


let__14379 := tmp16640
_ = let__14379

tmp16641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14375, Nil)
}
__typedArg0 := let__14375
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14378, tmp16641)
}
__typedArg0 := let__14378
__typedArg1 := tmp16641
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16642)
}
__typedArg0 := symwhere
__typedArg1 := tmp16642
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16644 := Call(__e, PrimFunc(sym_8p), let__14372, tmp16643)


tmp16645 := Call(__e, PrimFunc(symshen_4comb), let__14379, tmp16644)


ifres16638 = tmp16645


} else {
tmp16646 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16638 = tmp16646


}

ifres16636 = ifres16638


} else {
tmp16648 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16636 = tmp16648


}

ifres16633 = ifres16636


} else {
tmp16650 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16633 = tmp16650


}

ifres16631 = ifres16633


} else {
tmp16652 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16631 = tmp16652


}

ifres16627 = ifres16631


}

let__14370 := ifres16627
_ = let__14370

tmp16727 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14370)


if True == tmp16727 {
tmp16655 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5494)


let__14381 := tmp16655
_ = let__14381

tmp16684 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14381)


var ifres16656 Obj

if True == tmp16684 {
tmp16657 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16656 = tmp16657


} else {
tmp16658 := Call(__e, PrimFunc(symshen_4_5_1out), let__14381)


let__14382 := tmp16658
_ = let__14382

tmp16659 := Call(__e, PrimFunc(symshen_4in_1_6), let__14381)


let__14383 := tmp16659
_ = let__14383

tmp16683 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14383, sym_5_1)


var ifres16660 Obj

if True == tmp16683 {
tmp16661 := Call(__e, PrimFunc(symtail), let__14383)


let__14384 := tmp16661
_ = let__14384

tmp16681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14384)
}
__typedArg0 := let__14384
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16662 Obj

if True == tmp16681 {
tmp16663 := Call(__e, PrimFunc(symhead), let__14384)


let__14385 := tmp16663
_ = let__14385

tmp16664 := Call(__e, PrimFunc(symtail), let__14384)


let__14386 := tmp16664
_ = let__14386

tmp16679 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14386, symwhere)


var ifres16665 Obj

if True == tmp16679 {
tmp16666 := Call(__e, PrimFunc(symtail), let__14386)


let__14387 := tmp16666
_ = let__14387

tmp16677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14387)
}
__typedArg0 := let__14387
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16667 Obj

if True == tmp16677 {
tmp16668 := Call(__e, PrimFunc(symhead), let__14387)


let__14388 := tmp16668
_ = let__14388

tmp16669 := Call(__e, PrimFunc(symtail), let__14387)


let__14389 := tmp16669
_ = let__14389

tmp16670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14385, Nil)
}
__typedArg0 := let__14385
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14388, tmp16670)
}
__typedArg0 := let__14388
__typedArg1 := tmp16670
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16671)
}
__typedArg0 := symwhere
__typedArg1 := tmp16671
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16673 := Call(__e, PrimFunc(symshen_4correct), tmp16672)


tmp16674 := Call(__e, PrimFunc(sym_8p), let__14382, tmp16673)


tmp16675 := Call(__e, PrimFunc(symshen_4comb), let__14389, tmp16674)


ifres16667 = tmp16675


} else {
tmp16676 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16667 = tmp16676


}

ifres16665 = ifres16667


} else {
tmp16678 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16665 = tmp16678


}

ifres16662 = ifres16665


} else {
tmp16680 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16662 = tmp16680


}

ifres16660 = ifres16662


} else {
tmp16682 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16660 = tmp16682


}

ifres16656 = ifres16660


}

let__14380 := ifres16656
_ = let__14380

tmp16725 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14380)


if True == tmp16725 {
tmp16685 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5494)


let__14391 := tmp16685
_ = let__14391

tmp16702 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14391)


var ifres16686 Obj

if True == tmp16702 {
tmp16687 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16686 = tmp16687


} else {
tmp16688 := Call(__e, PrimFunc(symshen_4_5_1out), let__14391)


let__14392 := tmp16688
_ = let__14392

tmp16689 := Call(__e, PrimFunc(symshen_4in_1_6), let__14391)


let__14393 := tmp16689
_ = let__14393

tmp16701 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14393, sym_5_1)


var ifres16690 Obj

if True == tmp16701 {
tmp16691 := Call(__e, PrimFunc(symtail), let__14393)


let__14394 := tmp16691
_ = let__14394

tmp16699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14394)
}
__typedArg0 := let__14394
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16692 Obj

if True == tmp16699 {
tmp16693 := Call(__e, PrimFunc(symhead), let__14394)


let__14395 := tmp16693
_ = let__14395

tmp16694 := Call(__e, PrimFunc(symtail), let__14394)


let__14396 := tmp16694
_ = let__14396

tmp16695 := Call(__e, PrimFunc(symshen_4correct), let__14395)


tmp16696 := Call(__e, PrimFunc(sym_8p), let__14392, tmp16695)


tmp16697 := Call(__e, PrimFunc(symshen_4comb), let__14396, tmp16696)


ifres16692 = tmp16697


} else {
tmp16698 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16692 = tmp16698


}

ifres16690 = ifres16692


} else {
tmp16700 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16690 = tmp16700


}

ifres16686 = ifres16690


}

let__14390 := ifres16686
_ = let__14390

tmp16723 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14390)


if True == tmp16723 {
tmp16703 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5494)


let__14398 := tmp16703
_ = let__14398

tmp16719 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14398)


var ifres16704 Obj

if True == tmp16719 {
tmp16705 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16704 = tmp16705


} else {
tmp16706 := Call(__e, PrimFunc(symshen_4_5_1out), let__14398)


let__14399 := tmp16706
_ = let__14399

tmp16707 := Call(__e, PrimFunc(symshen_4in_1_6), let__14398)


let__14400 := tmp16707
_ = let__14400

tmp16718 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14400, sym_1_6)


var ifres16708 Obj

if True == tmp16718 {
tmp16709 := Call(__e, PrimFunc(symtail), let__14400)


let__14401 := tmp16709
_ = let__14401

tmp16716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14401)
}
__typedArg0 := let__14401
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16710 Obj

if True == tmp16716 {
tmp16711 := Call(__e, PrimFunc(symhead), let__14401)


let__14402 := tmp16711
_ = let__14402

tmp16712 := Call(__e, PrimFunc(symtail), let__14401)


let__14403 := tmp16712
_ = let__14403

tmp16713 := Call(__e, PrimFunc(sym_8p), let__14399, let__14402)


tmp16714 := Call(__e, PrimFunc(symshen_4comb), let__14403, tmp16713)


ifres16710 = tmp16714


} else {
tmp16715 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16710 = tmp16715


}

ifres16708 = ifres16710


} else {
tmp16717 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16708 = tmp16717


}

ifres16704 = ifres16708


}

let__14397 := ifres16704
_ = let__14397

tmp16721 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14397)


if True == tmp16721 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__14397)
return
}


} else {
__e.Return(let__14390)
return
}


} else {
__e.Return(let__14380)
return
}


} else {
__e.Return(let__14370)
return
}


}, 1)

tmp16728 := Call(__e, ns2_1set, symshen_4_5rule_d_6, tmp16625)


_ = tmp16728

tmp16729 := MakeNative(func(__e *ControlFlow) {
V5529 := __e.Get(1)
_ = V5529
tmp16877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16821 Obj

if True == tmp16877 {
tmp16875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp16875)
}
__typedArg0 := symwhere
__typedArg1 := tmp16875
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16823 Obj

if True == tmp16876 {
tmp16873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16873)
}
__typedArg0 := tmp16873
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16825 Obj

if True == tmp16874 {
tmp16870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16870)
}
__typedArg0 := tmp16870
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16871)
}
__typedArg0 := tmp16871
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16827 Obj

if True == tmp16872 {
tmp16866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16866)
}
__typedArg0 := tmp16866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16867)
}
__typedArg0 := tmp16867
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16868)
}
__typedArg0 := tmp16868
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16829 Obj

if True == tmp16869 {
tmp16861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16861)
}
__typedArg0 := tmp16861
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16862)
}
__typedArg0 := tmp16862
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16863)
}
__typedArg0 := tmp16863
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail_1if, tmp16864)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp16864
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16831 Obj

if True == tmp16865 {
tmp16856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16856)
}
__typedArg0 := tmp16856
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16857)
}
__typedArg0 := tmp16857
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16858)
}
__typedArg0 := tmp16858
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16859)
}
__typedArg0 := tmp16859
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16833 Obj

if True == tmp16860 {
tmp16850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16850)
}
__typedArg0 := tmp16850
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16851)
}
__typedArg0 := tmp16851
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16852)
}
__typedArg0 := tmp16852
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16853)
}
__typedArg0 := tmp16853
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16854)
}
__typedArg0 := tmp16854
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16835 Obj

if True == tmp16855 {
tmp16843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16843)
}
__typedArg0 := tmp16843
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16844)
}
__typedArg0 := tmp16844
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16845)
}
__typedArg0 := tmp16845
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16846)
}
__typedArg0 := tmp16846
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16847)
}
__typedArg0 := tmp16847
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp16848)
}
__typedArg0 := Nil
__typedArg1 := tmp16848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16837 Obj

if True == tmp16849 {
tmp16839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16839)
}
__typedArg0 := tmp16839
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16840)
}
__typedArg0 := tmp16840
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp16841)
}
__typedArg0 := Nil
__typedArg1 := tmp16841
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16838 Obj

if True == tmp16842 {
ifres16838 = True


} else {
ifres16838 = False


}

ifres16837 = ifres16838


} else {
ifres16837 = False


}

var ifres16836 Obj

if True == ifres16837 {
ifres16836 = True


} else {
ifres16836 = False


}

ifres16835 = ifres16836


} else {
ifres16835 = False


}

var ifres16834 Obj

if True == ifres16835 {
ifres16834 = True


} else {
ifres16834 = False


}

ifres16833 = ifres16834


} else {
ifres16833 = False


}

var ifres16832 Obj

if True == ifres16833 {
ifres16832 = True


} else {
ifres16832 = False


}

ifres16831 = ifres16832


} else {
ifres16831 = False


}

var ifres16830 Obj

if True == ifres16831 {
ifres16830 = True


} else {
ifres16830 = False


}

ifres16829 = ifres16830


} else {
ifres16829 = False


}

var ifres16828 Obj

if True == ifres16829 {
ifres16828 = True


} else {
ifres16828 = False


}

ifres16827 = ifres16828


} else {
ifres16827 = False


}

var ifres16826 Obj

if True == ifres16827 {
ifres16826 = True


} else {
ifres16826 = False


}

ifres16825 = ifres16826


} else {
ifres16825 = False


}

var ifres16824 Obj

if True == ifres16825 {
ifres16824 = True


} else {
ifres16824 = False


}

ifres16823 = ifres16824


} else {
ifres16823 = False


}

var ifres16822 Obj

if True == ifres16823 {
ifres16822 = True


} else {
ifres16822 = False


}

ifres16821 = ifres16822


} else {
ifres16821 = False


}

if True == ifres16821 {
tmp16730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16730)
}
__typedArg0 := tmp16730
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16732)
}
__typedArg0 := tmp16732
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16733)
}
__typedArg0 := tmp16733
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16734)
}
__typedArg0 := tmp16734
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16735, Nil)
}
__typedArg0 := tmp16735
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp16736)
}
__typedArg0 := symnot
__typedArg1 := tmp16736
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16737, Nil)
}
__typedArg0 := tmp16737
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16731, tmp16738)
}
__typedArg0 := tmp16731
__typedArg1 := tmp16738
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp16739)
}
__typedArg0 := symand
__typedArg1 := tmp16739
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16741)
}
__typedArg0 := tmp16741
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16742)
}
__typedArg0 := tmp16742
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16743)
}
__typedArg0 := tmp16743
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16744)
}
__typedArg0 := tmp16744
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16740, tmp16745)
}
__typedArg0 := tmp16740
__typedArg1 := tmp16745
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16746)
}
__typedArg0 := symwhere
__typedArg1 := tmp16746
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp16819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16800 Obj

if True == tmp16819 {
tmp16817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp16817)
}
__typedArg0 := symwhere
__typedArg1 := tmp16817
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16802 Obj

if True == tmp16818 {
tmp16815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16815)
}
__typedArg0 := tmp16815
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16804 Obj

if True == tmp16816 {
tmp16812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16812)
}
__typedArg0 := tmp16812
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16813)
}
__typedArg0 := tmp16813
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16806 Obj

if True == tmp16814 {
tmp16808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16808)
}
__typedArg0 := tmp16808
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16809)
}
__typedArg0 := tmp16809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp16810)
}
__typedArg0 := Nil
__typedArg1 := tmp16810
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16807 Obj

if True == tmp16811 {
ifres16807 = True


} else {
ifres16807 = False


}

ifres16806 = ifres16807


} else {
ifres16806 = False


}

var ifres16805 Obj

if True == ifres16806 {
ifres16805 = True


} else {
ifres16805 = False


}

ifres16804 = ifres16805


} else {
ifres16804 = False


}

var ifres16803 Obj

if True == ifres16804 {
ifres16803 = True


} else {
ifres16803 = False


}

ifres16802 = ifres16803


} else {
ifres16802 = False


}

var ifres16801 Obj

if True == ifres16802 {
ifres16801 = True


} else {
ifres16801 = False


}

ifres16800 = ifres16801


} else {
ifres16800 = False


}

if True == ifres16800 {
tmp16747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16747)
}
__typedArg0 := tmp16747
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16749)
}
__typedArg0 := tmp16749
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16750)
}
__typedArg0 := tmp16750
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, Nil)
}
__typedArg0 := symfail
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16752, Nil)
}
__typedArg0 := tmp16752
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16751, tmp16753)
}
__typedArg0 := tmp16751
__typedArg1 := tmp16753
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp16754)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp16754
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16755, Nil)
}
__typedArg0 := tmp16755
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp16756)
}
__typedArg0 := symnot
__typedArg1 := tmp16756
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16757, Nil)
}
__typedArg0 := tmp16757
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16748, tmp16758)
}
__typedArg0 := tmp16748
__typedArg1 := tmp16758
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp16759)
}
__typedArg0 := symand
__typedArg1 := tmp16759
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16761)
}
__typedArg0 := tmp16761
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16760, tmp16762)
}
__typedArg0 := tmp16760
__typedArg1 := tmp16762
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16763)
}
__typedArg0 := symwhere
__typedArg1 := tmp16763
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp16798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16779 Obj

if True == tmp16798 {
tmp16796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail_1if, tmp16796)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp16796
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16781 Obj

if True == tmp16797 {
tmp16794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16794)
}
__typedArg0 := tmp16794
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16783 Obj

if True == tmp16795 {
tmp16791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16791)
}
__typedArg0 := tmp16791
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16792)
}
__typedArg0 := tmp16792
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16785 Obj

if True == tmp16793 {
tmp16787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16787)
}
__typedArg0 := tmp16787
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16788)
}
__typedArg0 := tmp16788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp16789)
}
__typedArg0 := Nil
__typedArg1 := tmp16789
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16786 Obj

if True == tmp16790 {
ifres16786 = True


} else {
ifres16786 = False


}

ifres16785 = ifres16786


} else {
ifres16785 = False


}

var ifres16784 Obj

if True == ifres16785 {
ifres16784 = True


} else {
ifres16784 = False


}

ifres16783 = ifres16784


} else {
ifres16783 = False


}

var ifres16782 Obj

if True == ifres16783 {
ifres16782 = True


} else {
ifres16782 = False


}

ifres16781 = ifres16782


} else {
ifres16781 = False


}

var ifres16780 Obj

if True == ifres16781 {
ifres16780 = True


} else {
ifres16780 = False


}

ifres16779 = ifres16780


} else {
ifres16779 = False


}

if True == ifres16779 {
tmp16764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16764, Nil)
}
__typedArg0 := tmp16764
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp16765)
}
__typedArg0 := symnot
__typedArg1 := tmp16765
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5529)
}
__typedArg0 := V5529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16767)
}
__typedArg0 := tmp16767
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16766, tmp16768)
}
__typedArg0 := tmp16766
__typedArg1 := tmp16768
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16769)
}
__typedArg0 := symwhere
__typedArg1 := tmp16769
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp16770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, Nil)
}
__typedArg0 := symfail
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16770, Nil)
}
__typedArg0 := tmp16770
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5529, tmp16771)
}
__typedArg0 := V5529
__typedArg1 := tmp16771
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp16772)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp16772
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16773, Nil)
}
__typedArg0 := tmp16773
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp16774)
}
__typedArg0 := symnot
__typedArg1 := tmp16774
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5529, Nil)
}
__typedArg0 := V5529
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16775, tmp16776)
}
__typedArg0 := tmp16775
__typedArg1 := tmp16776
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16777)
}
__typedArg0 := symwhere
__typedArg1 := tmp16777
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}


}


}, 1)

tmp16878 := Call(__e, ns2_1set, symshen_4correct, tmp16729)


_ = tmp16878

tmp16879 := MakeNative(func(__e *ControlFlow) {
V5530 := __e.Get(1)
_ = V5530
V5531 := __e.Get(2)
_ = V5531
V5532 := __e.Get(3)
_ = V5532
V5533 := __e.Get(4)
_ = V5533
V5534 := __e.Get(5)
_ = V5534
V5535 := __e.Get(6)
_ = V5535
V5536 := __e.Get(7)
_ = V5536
V5537 := __e.Get(8)
_ = V5537
let__14404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5536)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5536
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__14404

tmp16887 := Call(__e, PrimFunc(symshen_4unlocked_2), V5535)


var ifres16881 Obj

if True == tmp16887 {
tmp16882 := Call(__e, PrimFunc(symshen_4lazyderef), V5531, V5534)


let__14406 := tmp16882
_ = let__14406

tmp16886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14406, Nil)
}
__typedArg0 := let__14406
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16883 Obj

if True == tmp16886 {
tmp16884 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16884

tmp16885 := Call(__e, PrimFunc(symthaw), V5537)


ifres16883 = tmp16885


} else {
ifres16883 = False


}

ifres16881 = ifres16883


} else {
ifres16881 = False


}

let__14405 := ifres16881
_ = let__14405

tmp16912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14405, False)
}
__typedArg0 := let__14405
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16912 {
tmp16908 := Call(__e, PrimFunc(symshen_4unlocked_2), V5535)


var ifres16888 Obj

if True == tmp16908 {
tmp16889 := Call(__e, PrimFunc(symshen_4lazyderef), V5531, V5534)


let__14408 := tmp16889
_ = let__14408

tmp16907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14408)
}
__typedArg0 := let__14408
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16890 Obj

if True == tmp16907 {
tmp16891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14408)
}
__typedArg0 := let__14408
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14409 := tmp16891
_ = let__14409

tmp16892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14408)
}
__typedArg0 := let__14408
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14410 := tmp16892
_ = let__14410

tmp16893 := Call(__e, PrimFunc(symshen_4newpv), V5534)


let__14411 := tmp16893
_ = let__14411

tmp16894 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16894

tmp16895 := Call(__e, PrimFunc(symshen_4deref), let__14409, V5534)


tmp16896 := Call(__e, PrimFunc(symshen_4freshen_1rule), tmp16895)


tmp16897 := MakeNative(func(__e *ControlFlow) {
tmp16898 := Call(__e, PrimFunc(symshen_4lazyderef), let__14411, V5534)


tmp16899 := Call(__e, PrimFunc(symfst), tmp16898)


tmp16900 := Call(__e, PrimFunc(symshen_4lazyderef), let__14411, V5534)


tmp16901 := Call(__e, PrimFunc(symsnd), tmp16900)


tmp16902 := MakeNative(func(__e *ControlFlow) {
tmp16903 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1rules), V5530, let__14410, V5532, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5533)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5533
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})(), V5534, V5535, let__14404, V5537)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5534, V5535, let__14404, tmp16903)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1rule), V5530, V5533, tmp16899, tmp16901, V5532, V5534, V5535, let__14404, tmp16902)
return


}, 0)

tmp16905 := Call(__e, PrimFunc(symbind), let__14411, tmp16896, V5534, V5535, let__14404, tmp16897)


tmp16906 := Call(__e, PrimFunc(symshen_4gc), V5534, tmp16905)


ifres16890 = tmp16906


} else {
ifres16890 = False


}

ifres16888 = ifres16890


} else {
ifres16888 = False


}

let__14407 := ifres16888
_ = let__14407

tmp16910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14407, False)
}
__typedArg0 := let__14407
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16910 {
__e.TailApply(PrimFunc(symshen_4unlock), V5535, let__14404)
return
} else {
__e.Return(let__14407)
return
}


} else {
__e.Return(let__14405)
return
}


}, 8)

tmp16913 := Call(__e, ns2_1set, symshen_4t_d_1rules, tmp16879)


_ = tmp16913

tmp16914 := MakeNative(func(__e *ControlFlow) {
V5546 := __e.Get(1)
_ = V5546
tmp16925 := Call(__e, PrimFunc(symtuple_2), V5546)


if True == tmp16925 {
tmp16915 := Call(__e, PrimFunc(symfst), V5546)


tmp16916 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp16915)


let__14412 := tmp16916
_ = let__14412

tmp16917 := MakeNative(func(__e *ControlFlow) {
Z5549 := __e.Get(1)
_ = Z5549
tmp16918 := Call(__e, PrimFunc(symshen_4freshterm), Z5549)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z5549, tmp16918)
}
__typedArg0 := Z5549
__typedArg1 := tmp16918
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp16919 := Call(__e, PrimFunc(symmap), tmp16917, let__14412)


let__14413 := tmp16919
_ = let__14413

tmp16920 := Call(__e, PrimFunc(symfst), V5546)


tmp16921 := Call(__e, PrimFunc(symshen_4freshen), let__14413, tmp16920)


tmp16922 := Call(__e, PrimFunc(symsnd), V5546)


tmp16923 := Call(__e, PrimFunc(symshen_4freshen), let__14413, tmp16922)


__e.TailApply(PrimFunc(sym_8p), tmp16921, tmp16923)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshen_1rule)
return
}


}, 1)

tmp16926 := Call(__e, ns2_1set, symshen_4freshen_1rule, tmp16914)


_ = tmp16926

tmp16927 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V5550 := __self1
_ = V5550
V5551 := __self2
_ = V5551
tmp16941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5550)
}
__typedArg0 := Nil
__typedArg1 := V5550
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16941 {
__e.Return(V5551)
return
} else {
tmp16939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5550)
}
__typedArg0 := V5550
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16935 Obj

if True == tmp16939 {
tmp16937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5550)
}
__typedArg0 := V5550
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16937)
}
__typedArg0 := tmp16937
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16936 Obj

if True == tmp16938 {
ifres16936 = True


} else {
ifres16936 = False


}

ifres16935 = ifres16936


} else {
ifres16935 = False


}

if True == ifres16935 {
tmp16928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5550)
}
__typedArg0 := V5550
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5550)
}
__typedArg0 := V5550
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16929)
}
__typedArg0 := tmp16929
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5550)
}
__typedArg0 := V5550
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16931)
}
__typedArg0 := tmp16931
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16933 := Call(__e, PrimFunc(symshen_4beta), tmp16930, tmp16932, V5551)


if PrimFunc(symshen_4freshen) == __self {
__self1, __self2 = tmp16928, tmp16933
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4freshen), tmp16928, tmp16933)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshen)
return
}


}


}, 2)

tmp16942 := Call(__e, ns2_1set, symshen_4freshen, tmp16927)


_ = tmp16942

tmp16943 := MakeNative(func(__e *ControlFlow) {
V5552 := __e.Get(1)
_ = V5552
V5553 := __e.Get(2)
_ = V5553
V5554 := __e.Get(3)
_ = V5554
V5555 := __e.Get(4)
_ = V5555
V5556 := __e.Get(5)
_ = V5556
V5557 := __e.Get(6)
_ = V5557
V5558 := __e.Get(7)
_ = V5558
V5559 := __e.Get(8)
_ = V5559
V5560 := __e.Get(9)
_ = V5560
tmp16947 := Call(__e, PrimFunc(symshen_4unlocked_2), V5558)


var ifres16944 Obj

if True == tmp16947 {
tmp16945 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16945

tmp16946 := Call(__e, PrimFunc(symshen_4t_d_1rule_1h), V5554, V5555, V5556, V5557, V5558, V5559, V5560)


ifres16944 = tmp16946


} else {
ifres16944 = False


}

let__14414 := ifres16944
_ = let__14414

tmp16959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14414, False)
}
__typedArg0 := let__14414
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16959 {
tmp16957 := Call(__e, PrimFunc(symshen_4unlocked_2), V5558)


if True == tmp16957 {
tmp16948 := Call(__e, PrimFunc(symshen_4newpv), V5557)


let__14415 := tmp16948
_ = let__14415

tmp16949 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16949

tmp16950 := Call(__e, PrimFunc(symshen_4app), V5552, MakeString("\n"), symshen_4a)


tmp16952 := Call(__e, PrimFunc(symshen_4app), V5553, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" of "))
__typedS1, __typedOK1 := TypedString(tmp16950)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" of ")
__typedArg1 := tmp16950
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp16954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("type error in rule "))
__typedS1, __typedOK1 := TypedString(tmp16952)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("type error in rule ")
__typedArg1 := tmp16952
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("type error in rule "))
__typedS1, __typedOK1 := TypedString(tmp16952)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("type error in rule ")
__typedArg1 := tmp16952
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

tmp16955 := Call(__e, PrimFunc(symbind), let__14415, tmp16954, V5557, V5558, V5559, V5560)


__e.TailApply(PrimFunc(symshen_4gc), V5557, tmp16955)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(let__14414)
return
}


}, 9)

tmp16960 := Call(__e, ns2_1set, symshen_4t_d_1rule, tmp16943)


_ = tmp16960

tmp16961 := MakeNative(func(__e *ControlFlow) {
V5563 := __e.Get(1)
_ = V5563
V5564 := __e.Get(2)
_ = V5564
V5565 := __e.Get(3)
_ = V5565
V5566 := __e.Get(4)
_ = V5566
V5567 := __e.Get(5)
_ = V5567
V5568 := __e.Get(6)
_ = V5568
V5569 := __e.Get(7)
_ = V5569
let__14416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5568)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5568
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__14416

tmp16986 := Call(__e, PrimFunc(symshen_4unlocked_2), V5567)


var ifres16963 Obj

if True == tmp16986 {
tmp16964 := Call(__e, PrimFunc(symshen_4lazyderef), V5563, V5566)


let__14418 := tmp16964
_ = let__14418

tmp16985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14418, Nil)
}
__typedArg0 := let__14418
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16965 Obj

if True == tmp16985 {
tmp16966 := Call(__e, PrimFunc(symshen_4lazyderef), V5565, V5566)


let__14419 := tmp16966
_ = let__14419

tmp16984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14419)
}
__typedArg0 := let__14419
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16967 Obj

if True == tmp16984 {
tmp16968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14419)
}
__typedArg0 := let__14419
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16969 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16968, V5566)


let__14420 := tmp16969
_ = let__14420

tmp16983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14420, sym_1_1_6)
}
__typedArg0 := let__14420
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16970 Obj

if True == tmp16983 {
tmp16971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14419)
}
__typedArg0 := let__14419
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16972 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16971, V5566)


let__14421 := tmp16972
_ = let__14421

tmp16982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14421)
}
__typedArg0 := let__14421
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16973 Obj

if True == tmp16982 {
tmp16974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14421)
}
__typedArg0 := let__14421
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14422 := tmp16974
_ = let__14422

tmp16975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14421)
}
__typedArg0 := let__14421
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16976 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16975, V5566)


let__14423 := tmp16976
_ = let__14423

tmp16981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14423, Nil)
}
__typedArg0 := let__14423
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16977 Obj

if True == tmp16981 {
tmp16978 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16978

tmp16979 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1correct), V5564, let__14422, Nil, V5566, V5567, let__14416, V5569)
return
}, 0)

tmp16980 := Call(__e, PrimFunc(symshen_4cut), V5566, V5567, let__14416, tmp16979)


ifres16977 = tmp16980


} else {
ifres16977 = False


}

ifres16973 = ifres16977


} else {
ifres16973 = False


}

ifres16970 = ifres16973


} else {
ifres16970 = False


}

ifres16967 = ifres16970


} else {
ifres16967 = False


}

ifres16965 = ifres16967


} else {
ifres16965 = False


}

ifres16963 = ifres16965


} else {
ifres16963 = False


}

let__14417 := ifres16963
_ = let__14417

tmp17005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14417, False)
}
__typedArg0 := let__14417
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17005 {
tmp17001 := Call(__e, PrimFunc(symshen_4unlocked_2), V5567)


var ifres16987 Obj

if True == tmp17001 {
tmp16988 := Call(__e, PrimFunc(symshen_4newpv), V5566)


let__14425 := tmp16988
_ = let__14425

tmp16989 := Call(__e, PrimFunc(symshen_4newpv), V5566)


let__14426 := tmp16989
_ = let__14426

tmp16990 := Call(__e, PrimFunc(symshen_4newpv), V5566)


let__14427 := tmp16990
_ = let__14427

tmp16991 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16991

tmp16992 := Call(__e, PrimFunc(symshen_4freshterms), V5563)


tmp16993 := MakeNative(func(__e *ControlFlow) {
tmp16994 := MakeNative(func(__e *ControlFlow) {
tmp16995 := MakeNative(func(__e *ControlFlow) {
tmp16996 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1correct), V5564, let__14426, let__14427, V5566, V5567, let__14416, V5569)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4myassume), V5563, V5565, let__14427, V5566, V5567, let__14416, tmp16996)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5566, V5567, let__14416, tmp16995)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1integrity), V5563, V5565, let__14425, let__14426, V5566, V5567, let__14416, tmp16994)
return


}, 0)

tmp16997 := Call(__e, PrimFunc(symshen_4p_1hyps), tmp16992, let__14425, V5566, V5567, let__14416, tmp16993)


tmp16998 := Call(__e, PrimFunc(symshen_4gc), V5566, tmp16997)


tmp16999 := Call(__e, PrimFunc(symshen_4gc), V5566, tmp16998)


tmp17000 := Call(__e, PrimFunc(symshen_4gc), V5566, tmp16999)


ifres16987 = tmp17000


} else {
ifres16987 = False


}

let__14424 := ifres16987
_ = let__14424

tmp17003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14424, False)
}
__typedArg0 := let__14424
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17003 {
__e.TailApply(PrimFunc(symshen_4unlock), V5567, let__14416)
return
} else {
__e.Return(let__14424)
return
}


} else {
__e.Return(let__14417)
return
}


}, 7)

tmp17006 := Call(__e, ns2_1set, symshen_4t_d_1rule_1h, tmp16961)


_ = tmp17006

tmp17007 := MakeNative(func(__e *ControlFlow) {
V5582 := __e.Get(1)
_ = V5582
V5583 := __e.Get(2)
_ = V5583
V5584 := __e.Get(3)
_ = V5584
V5585 := __e.Get(4)
_ = V5585
V5586 := __e.Get(5)
_ = V5586
V5587 := __e.Get(6)
_ = V5587
V5588 := __e.Get(7)
_ = V5588
tmp17021 := Call(__e, PrimFunc(symshen_4unlocked_2), V5586)


var ifres17008 Obj

if True == tmp17021 {
tmp17009 := Call(__e, PrimFunc(symshen_4lazyderef), V5582, V5585)


let__14429 := tmp17009
_ = let__14429

tmp17020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14429, Nil)
}
__typedArg0 := let__14429
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17010 Obj

if True == tmp17020 {
tmp17011 := Call(__e, PrimFunc(symshen_4lazyderef), V5584, V5585)


let__14430 := tmp17011
_ = let__14430

tmp17012 := MakeNative(func(__e *ControlFlow) {
tmp17013 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17013

__e.TailApply(PrimFunc(symthaw), V5588)
return


}, 0)

let__14431 := tmp17012
_ = let__14431

tmp17019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14430, Nil)
}
__typedArg0 := let__14430
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17014 Obj

if True == tmp17019 {
tmp17015 := Call(__e, PrimFunc(symthaw), let__14431)


ifres17014 = tmp17015


} else {
tmp17018 := Call(__e, PrimFunc(symshen_4pvar_2), let__14430)


var ifres17016 Obj

if True == tmp17018 {
tmp17017 := Call(__e, PrimFunc(symshen_4bind_b), let__14430, Nil, V5585, let__14431)


ifres17016 = tmp17017


} else {
ifres17016 = False


}

ifres17014 = ifres17016


}

ifres17010 = ifres17014


} else {
ifres17010 = False


}

ifres17008 = ifres17010


} else {
ifres17008 = False


}

let__14428 := ifres17008
_ = let__14428

tmp17140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14428, False)
}
__typedArg0 := let__14428
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17140 {
tmp17138 := Call(__e, PrimFunc(symshen_4unlocked_2), V5586)


if True == tmp17138 {
tmp17022 := Call(__e, PrimFunc(symshen_4lazyderef), V5582, V5585)


let__14432 := tmp17022
_ = let__14432

tmp17136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14432)
}
__typedArg0 := let__14432
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17136 {
tmp17023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14432)
}
__typedArg0 := let__14432
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14433 := tmp17023
_ = let__14433

tmp17024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14432)
}
__typedArg0 := let__14432
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14434 := tmp17024
_ = let__14434

tmp17025 := Call(__e, PrimFunc(symshen_4lazyderef), V5583, V5585)


let__14435 := tmp17025
_ = let__14435

tmp17134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14435)
}
__typedArg0 := let__14435
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17134 {
tmp17026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14435)
}
__typedArg0 := let__14435
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14436 := tmp17026
_ = let__14436

tmp17027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14435)
}
__typedArg0 := let__14435
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17028 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17027, V5585)


let__14437 := tmp17028
_ = let__14437

tmp17132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14437)
}
__typedArg0 := let__14437
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17132 {
tmp17029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14437)
}
__typedArg0 := let__14437
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17030 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17029, V5585)


let__14438 := tmp17030
_ = let__14438

tmp17130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14438, sym_1_1_6)
}
__typedArg0 := let__14438
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17130 {
tmp17031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14437)
}
__typedArg0 := let__14437
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17032 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17031, V5585)


let__14439 := tmp17032
_ = let__14439

tmp17128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14439)
}
__typedArg0 := let__14439
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17128 {
tmp17033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14439)
}
__typedArg0 := let__14439
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14440 := tmp17033
_ = let__14440

tmp17034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14439)
}
__typedArg0 := let__14439
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17035 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17034, V5585)


let__14441 := tmp17035
_ = let__14441

tmp17126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14441, Nil)
}
__typedArg0 := let__14441
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17126 {
tmp17036 := Call(__e, PrimFunc(symshen_4lazyderef), V5584, V5585)


let__14442 := tmp17036
_ = let__14442

tmp17037 := MakeNative(func(__e *ControlFlow) {
Z5605 := __e.Get(1)
_ = Z5605
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5606 := __e.Get(1)
_ = Z5606
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5607 := __e.Get(1)
_ = Z5607
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5608 := __e.Get(1)
_ = Z5608
tmp17038 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17038

tmp17039 := MakeNative(func(__e *ControlFlow) {
tmp17040 := MakeNative(func(__e *ControlFlow) {
tmp17041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp17042 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4myassume), let__14434, let__14440, Z5608, V5585, V5586, V5587, V5588)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5606, tmp17041, V5585, V5586, V5587, tmp17042)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), let__14433, Z5605, V5585, V5586, V5587, tmp17040)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), let__14436, Z5607, V5585, V5586, V5587, tmp17039)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

let__14443 := tmp17037
_ = let__14443

tmp17124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14442)
}
__typedArg0 := let__14442
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17124 {
tmp17043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14442)
}
__typedArg0 := let__14442
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17044 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17043, V5585)


let__14444 := tmp17044
_ = let__14444

tmp17045 := MakeNative(func(__e *ControlFlow) {
Z5611 := __e.Get(1)
_ = Z5611
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5612 := __e.Get(1)
_ = Z5612
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5613 := __e.Get(1)
_ = Z5613
tmp17046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14442)
}
__typedArg0 := let__14442
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14446 := tmp17046
_ = let__14446

tmp17047 := Call(__e, let__14443, Z5611)


tmp17048 := Call(__e, tmp17047, Z5612)


tmp17049 := Call(__e, tmp17048, Z5613)


__e.TailApply(tmp17049, let__14446)
return


}, 1))
return
}, 1))
return
}, 1)

let__14445 := tmp17045
_ = let__14445

tmp17104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14444)
}
__typedArg0 := let__14444
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17104 {
tmp17050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14444)
}
__typedArg0 := let__14444
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14447 := tmp17050
_ = let__14447

tmp17051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14444)
}
__typedArg0 := let__14444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17052 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17051, V5585)


let__14448 := tmp17052
_ = let__14448

tmp17053 := MakeNative(func(__e *ControlFlow) {
Z5618 := __e.Get(1)
_ = Z5618
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5619 := __e.Get(1)
_ = Z5619
tmp17054 := Call(__e, let__14445, let__14447)


tmp17055 := Call(__e, tmp17054, Z5618)


__e.TailApply(tmp17055, Z5619)
return


}, 1))
return
}, 1)

let__14449 := tmp17053
_ = let__14449

tmp17088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14448)
}
__typedArg0 := let__14448
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17088 {
tmp17056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14448)
}
__typedArg0 := let__14448
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14450 := tmp17056
_ = let__14450

tmp17057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14448)
}
__typedArg0 := let__14448
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17058 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17057, V5585)


let__14451 := tmp17058
_ = let__14451

tmp17059 := MakeNative(func(__e *ControlFlow) {
Z5623 := __e.Get(1)
_ = Z5623
tmp17060 := Call(__e, let__14449, let__14450)


__e.TailApply(tmp17060, Z5623)
return


}, 1)

let__14452 := tmp17059
_ = let__14452

tmp17076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14451)
}
__typedArg0 := let__14451
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17076 {
tmp17061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14451)
}
__typedArg0 := let__14451
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14453 := tmp17061
_ = let__14453

tmp17062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14451)
}
__typedArg0 := let__14451
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17063 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17062, V5585)


let__14454 := tmp17063
_ = let__14454

tmp17064 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14452, let__14453)
return
}, 0)

let__14455 := tmp17064
_ = let__14455

tmp17068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14454, Nil)
}
__typedArg0 := let__14454
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17068 {
__e.TailApply(PrimFunc(symthaw), let__14455)
return
} else {
tmp17066 := Call(__e, PrimFunc(symshen_4pvar_2), let__14454)


if True == tmp17066 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14454, Nil, V5585, let__14455)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp17074 := Call(__e, PrimFunc(symshen_4pvar_2), let__14451)


if True == tmp17074 {
tmp17069 := Call(__e, PrimFunc(symshen_4newpv), V5585)


let__14456 := tmp17069
_ = let__14456

tmp17070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14456, Nil)
}
__typedArg0 := let__14456
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17071 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14452, let__14456)
return
}, 0)

tmp17072 := Call(__e, PrimFunc(symshen_4bind_b), let__14451, tmp17070, V5585, tmp17071)


__e.TailApply(PrimFunc(symshen_4gc), V5585, tmp17072)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17086 := Call(__e, PrimFunc(symshen_4pvar_2), let__14448)


if True == tmp17086 {
tmp17077 := Call(__e, PrimFunc(symshen_4newpv), V5585)


let__14457 := tmp17077
_ = let__14457

tmp17078 := Call(__e, PrimFunc(symshen_4newpv), V5585)


let__14458 := tmp17078
_ = let__14458

tmp17079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14458, Nil)
}
__typedArg0 := let__14458
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14457, tmp17079)
}
__typedArg0 := let__14457
__typedArg1 := tmp17079
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17081 := MakeNative(func(__e *ControlFlow) {
tmp17082 := Call(__e, let__14449, let__14457)


__e.TailApply(tmp17082, let__14458)
return


}, 0)

tmp17083 := Call(__e, PrimFunc(symshen_4bind_b), let__14448, tmp17080, V5585, tmp17081)


tmp17084 := Call(__e, PrimFunc(symshen_4gc), V5585, tmp17083)


__e.TailApply(PrimFunc(symshen_4gc), V5585, tmp17084)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17102 := Call(__e, PrimFunc(symshen_4pvar_2), let__14444)


if True == tmp17102 {
tmp17089 := Call(__e, PrimFunc(symshen_4newpv), V5585)


let__14459 := tmp17089
_ = let__14459

tmp17090 := Call(__e, PrimFunc(symshen_4newpv), V5585)


let__14460 := tmp17090
_ = let__14460

tmp17091 := Call(__e, PrimFunc(symshen_4newpv), V5585)


let__14461 := tmp17091
_ = let__14461

tmp17092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14461, Nil)
}
__typedArg0 := let__14461
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14460, tmp17092)
}
__typedArg0 := let__14460
__typedArg1 := tmp17092
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14459, tmp17093)
}
__typedArg0 := let__14459
__typedArg1 := tmp17093
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17095 := MakeNative(func(__e *ControlFlow) {
tmp17096 := Call(__e, let__14445, let__14459)


tmp17097 := Call(__e, tmp17096, let__14460)


__e.TailApply(tmp17097, let__14461)
return


}, 0)

tmp17098 := Call(__e, PrimFunc(symshen_4bind_b), let__14444, tmp17094, V5585, tmp17095)


tmp17099 := Call(__e, PrimFunc(symshen_4gc), V5585, tmp17098)


tmp17100 := Call(__e, PrimFunc(symshen_4gc), V5585, tmp17099)


__e.TailApply(PrimFunc(symshen_4gc), V5585, tmp17100)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17122 := Call(__e, PrimFunc(symshen_4pvar_2), let__14442)


if True == tmp17122 {
tmp17105 := Call(__e, PrimFunc(symshen_4newpv), V5585)


let__14462 := tmp17105
_ = let__14462

tmp17106 := Call(__e, PrimFunc(symshen_4newpv), V5585)


let__14463 := tmp17106
_ = let__14463

tmp17107 := Call(__e, PrimFunc(symshen_4newpv), V5585)


let__14464 := tmp17107
_ = let__14464

tmp17108 := Call(__e, PrimFunc(symshen_4newpv), V5585)


let__14465 := tmp17108
_ = let__14465

tmp17109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14464, Nil)
}
__typedArg0 := let__14464
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14463, tmp17109)
}
__typedArg0 := let__14463
__typedArg1 := tmp17109
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14462, tmp17110)
}
__typedArg0 := let__14462
__typedArg1 := tmp17110
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17111, let__14465)
}
__typedArg0 := tmp17111
__typedArg1 := let__14465
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17113 := MakeNative(func(__e *ControlFlow) {
tmp17114 := Call(__e, let__14443, let__14462)


tmp17115 := Call(__e, tmp17114, let__14463)


tmp17116 := Call(__e, tmp17115, let__14464)


__e.TailApply(tmp17116, let__14465)
return


}, 0)

tmp17117 := Call(__e, PrimFunc(symshen_4bind_b), let__14442, tmp17112, V5585, tmp17113)


tmp17118 := Call(__e, PrimFunc(symshen_4gc), V5585, tmp17117)


tmp17119 := Call(__e, PrimFunc(symshen_4gc), V5585, tmp17118)


tmp17120 := Call(__e, PrimFunc(symshen_4gc), V5585, tmp17119)


__e.TailApply(PrimFunc(symshen_4gc), V5585, tmp17120)
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
__e.Return(let__14428)
return
}


}, 7)

tmp17141 := Call(__e, ns2_1set, symshen_4myassume, tmp17007)


_ = tmp17141

tmp17142 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V5639 := __self1
_ = V5639
tmp17165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5639)
}
__typedArg0 := Nil
__typedArg1 := V5639
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17165 {
__e.Return(Nil)
return
} else {
tmp17163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5639)
}
__typedArg0 := V5639
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17159 Obj

if True == tmp17163 {
tmp17161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5639)
}
__typedArg0 := V5639
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17161)
}
__typedArg0 := tmp17161
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17160 Obj

if True == tmp17162 {
ifres17160 = True


} else {
ifres17160 = False


}

ifres17159 = ifres17160


} else {
ifres17159 = False


}

if True == ifres17159 {
tmp17143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5639)
}
__typedArg0 := V5639
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5639)
}
__typedArg0 := V5639
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17145 := Call(__e, PrimFunc(symappend), tmp17143, tmp17144)


if PrimFunc(symshen_4freshterms) == __self {
__self1 = tmp17145
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4freshterms), tmp17145)
return


} else {
tmp17157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5639)
}
__typedArg0 := V5639
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17153 Obj

if True == tmp17157 {
tmp17155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5639)
}
__typedArg0 := V5639
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17156 := Call(__e, PrimFunc(symshen_4freshterm_2), tmp17155)


var ifres17154 Obj

if True == tmp17156 {
ifres17154 = True


} else {
ifres17154 = False


}

ifres17153 = ifres17154


} else {
ifres17153 = False


}

if True == ifres17153 {
tmp17146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5639)
}
__typedArg0 := V5639
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5639)
}
__typedArg0 := V5639
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17148 := Call(__e, PrimFunc(symshen_4freshterms), tmp17147)


__e.TailApply(PrimFunc(symadjoin), tmp17146, tmp17148)
return


} else {
tmp17151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5639)
}
__typedArg0 := V5639
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17151 {
tmp17149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5639)
}
__typedArg0 := V5639
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4freshterms) == __self {
__self1 = tmp17149
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4freshterms), tmp17149)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshterms)
return
}


}


}


}


}, 1)

tmp17166 := Call(__e, ns2_1set, symshen_4freshterms, tmp17142)


_ = tmp17166

tmp17167 := MakeNative(func(__e *ControlFlow) {
V5640 := __e.Get(1)
_ = V5640
V5641 := __e.Get(2)
_ = V5641
V5642 := __e.Get(3)
_ = V5642
V5643 := __e.Get(4)
_ = V5643
V5644 := __e.Get(5)
_ = V5644
V5645 := __e.Get(6)
_ = V5645
tmp17181 := Call(__e, PrimFunc(symshen_4unlocked_2), V5643)


var ifres17168 Obj

if True == tmp17181 {
tmp17169 := Call(__e, PrimFunc(symshen_4lazyderef), V5640, V5642)


let__14467 := tmp17169
_ = let__14467

tmp17180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14467, Nil)
}
__typedArg0 := let__14467
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17170 Obj

if True == tmp17180 {
tmp17171 := Call(__e, PrimFunc(symshen_4lazyderef), V5641, V5642)


let__14468 := tmp17171
_ = let__14468

tmp17172 := MakeNative(func(__e *ControlFlow) {
tmp17173 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17173

__e.TailApply(PrimFunc(symthaw), V5645)
return


}, 0)

let__14469 := tmp17172
_ = let__14469

tmp17179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14468, Nil)
}
__typedArg0 := let__14468
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17174 Obj

if True == tmp17179 {
tmp17175 := Call(__e, PrimFunc(symthaw), let__14469)


ifres17174 = tmp17175


} else {
tmp17178 := Call(__e, PrimFunc(symshen_4pvar_2), let__14468)


var ifres17176 Obj

if True == tmp17178 {
tmp17177 := Call(__e, PrimFunc(symshen_4bind_b), let__14468, Nil, V5642, let__14469)


ifres17176 = tmp17177


} else {
ifres17176 = False


}

ifres17174 = ifres17176


}

ifres17170 = ifres17174


} else {
ifres17170 = False


}

ifres17168 = ifres17170


} else {
ifres17168 = False


}

let__14466 := ifres17168
_ = let__14466

tmp17278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14466, False)
}
__typedArg0 := let__14466
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17278 {
tmp17276 := Call(__e, PrimFunc(symshen_4unlocked_2), V5643)


if True == tmp17276 {
tmp17182 := Call(__e, PrimFunc(symshen_4lazyderef), V5640, V5642)


let__14470 := tmp17182
_ = let__14470

tmp17274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14470)
}
__typedArg0 := let__14470
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17274 {
tmp17183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14470)
}
__typedArg0 := let__14470
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14471 := tmp17183
_ = let__14471

tmp17184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14470)
}
__typedArg0 := let__14470
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14472 := tmp17184
_ = let__14472

tmp17185 := Call(__e, PrimFunc(symshen_4lazyderef), V5641, V5642)


let__14473 := tmp17185
_ = let__14473

tmp17186 := MakeNative(func(__e *ControlFlow) {
Z5655 := __e.Get(1)
_ = Z5655
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5656 := __e.Get(1)
_ = Z5656
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5657 := __e.Get(1)
_ = Z5657
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5658 := __e.Get(1)
_ = Z5658
tmp17187 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17187

tmp17188 := MakeNative(func(__e *ControlFlow) {
tmp17189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp17190 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4p_1hyps), let__14472, Z5658, V5642, V5643, V5644, V5645)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5656, tmp17189, V5642, V5643, V5644, tmp17190)
return


}, 0)

__e.TailApply(PrimFunc(symbind), Z5655, let__14471, V5642, V5643, V5644, tmp17188)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

let__14474 := tmp17186
_ = let__14474

tmp17272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14473)
}
__typedArg0 := let__14473
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17272 {
tmp17191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14473)
}
__typedArg0 := let__14473
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17192 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17191, V5642)


let__14475 := tmp17192
_ = let__14475

tmp17193 := MakeNative(func(__e *ControlFlow) {
Z5661 := __e.Get(1)
_ = Z5661
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5662 := __e.Get(1)
_ = Z5662
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5663 := __e.Get(1)
_ = Z5663
tmp17194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14473)
}
__typedArg0 := let__14473
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14477 := tmp17194
_ = let__14477

tmp17195 := Call(__e, let__14474, Z5661)


tmp17196 := Call(__e, tmp17195, Z5662)


tmp17197 := Call(__e, tmp17196, Z5663)


__e.TailApply(tmp17197, let__14477)
return


}, 1))
return
}, 1))
return
}, 1)

let__14476 := tmp17193
_ = let__14476

tmp17252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14475)
}
__typedArg0 := let__14475
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17252 {
tmp17198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14475)
}
__typedArg0 := let__14475
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14478 := tmp17198
_ = let__14478

tmp17199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14475)
}
__typedArg0 := let__14475
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17200 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17199, V5642)


let__14479 := tmp17200
_ = let__14479

tmp17201 := MakeNative(func(__e *ControlFlow) {
Z5668 := __e.Get(1)
_ = Z5668
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5669 := __e.Get(1)
_ = Z5669
tmp17202 := Call(__e, let__14476, let__14478)


tmp17203 := Call(__e, tmp17202, Z5668)


__e.TailApply(tmp17203, Z5669)
return


}, 1))
return
}, 1)

let__14480 := tmp17201
_ = let__14480

tmp17236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14479)
}
__typedArg0 := let__14479
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17236 {
tmp17204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14479)
}
__typedArg0 := let__14479
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14481 := tmp17204
_ = let__14481

tmp17205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14479)
}
__typedArg0 := let__14479
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17206 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17205, V5642)


let__14482 := tmp17206
_ = let__14482

tmp17207 := MakeNative(func(__e *ControlFlow) {
Z5673 := __e.Get(1)
_ = Z5673
tmp17208 := Call(__e, let__14480, let__14481)


__e.TailApply(tmp17208, Z5673)
return


}, 1)

let__14483 := tmp17207
_ = let__14483

tmp17224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14482)
}
__typedArg0 := let__14482
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17224 {
tmp17209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14482)
}
__typedArg0 := let__14482
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14484 := tmp17209
_ = let__14484

tmp17210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14482)
}
__typedArg0 := let__14482
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17211 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17210, V5642)


let__14485 := tmp17211
_ = let__14485

tmp17212 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14483, let__14484)
return
}, 0)

let__14486 := tmp17212
_ = let__14486

tmp17216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14485, Nil)
}
__typedArg0 := let__14485
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17216 {
__e.TailApply(PrimFunc(symthaw), let__14486)
return
} else {
tmp17214 := Call(__e, PrimFunc(symshen_4pvar_2), let__14485)


if True == tmp17214 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14485, Nil, V5642, let__14486)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp17222 := Call(__e, PrimFunc(symshen_4pvar_2), let__14482)


if True == tmp17222 {
tmp17217 := Call(__e, PrimFunc(symshen_4newpv), V5642)


let__14487 := tmp17217
_ = let__14487

tmp17218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14487, Nil)
}
__typedArg0 := let__14487
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17219 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14483, let__14487)
return
}, 0)

tmp17220 := Call(__e, PrimFunc(symshen_4bind_b), let__14482, tmp17218, V5642, tmp17219)


__e.TailApply(PrimFunc(symshen_4gc), V5642, tmp17220)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17234 := Call(__e, PrimFunc(symshen_4pvar_2), let__14479)


if True == tmp17234 {
tmp17225 := Call(__e, PrimFunc(symshen_4newpv), V5642)


let__14488 := tmp17225
_ = let__14488

tmp17226 := Call(__e, PrimFunc(symshen_4newpv), V5642)


let__14489 := tmp17226
_ = let__14489

tmp17227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14489, Nil)
}
__typedArg0 := let__14489
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14488, tmp17227)
}
__typedArg0 := let__14488
__typedArg1 := tmp17227
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17229 := MakeNative(func(__e *ControlFlow) {
tmp17230 := Call(__e, let__14480, let__14488)


__e.TailApply(tmp17230, let__14489)
return


}, 0)

tmp17231 := Call(__e, PrimFunc(symshen_4bind_b), let__14479, tmp17228, V5642, tmp17229)


tmp17232 := Call(__e, PrimFunc(symshen_4gc), V5642, tmp17231)


__e.TailApply(PrimFunc(symshen_4gc), V5642, tmp17232)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17250 := Call(__e, PrimFunc(symshen_4pvar_2), let__14475)


if True == tmp17250 {
tmp17237 := Call(__e, PrimFunc(symshen_4newpv), V5642)


let__14490 := tmp17237
_ = let__14490

tmp17238 := Call(__e, PrimFunc(symshen_4newpv), V5642)


let__14491 := tmp17238
_ = let__14491

tmp17239 := Call(__e, PrimFunc(symshen_4newpv), V5642)


let__14492 := tmp17239
_ = let__14492

tmp17240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14492, Nil)
}
__typedArg0 := let__14492
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14491, tmp17240)
}
__typedArg0 := let__14491
__typedArg1 := tmp17240
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14490, tmp17241)
}
__typedArg0 := let__14490
__typedArg1 := tmp17241
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17243 := MakeNative(func(__e *ControlFlow) {
tmp17244 := Call(__e, let__14476, let__14490)


tmp17245 := Call(__e, tmp17244, let__14491)


__e.TailApply(tmp17245, let__14492)
return


}, 0)

tmp17246 := Call(__e, PrimFunc(symshen_4bind_b), let__14475, tmp17242, V5642, tmp17243)


tmp17247 := Call(__e, PrimFunc(symshen_4gc), V5642, tmp17246)


tmp17248 := Call(__e, PrimFunc(symshen_4gc), V5642, tmp17247)


__e.TailApply(PrimFunc(symshen_4gc), V5642, tmp17248)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17270 := Call(__e, PrimFunc(symshen_4pvar_2), let__14473)


if True == tmp17270 {
tmp17253 := Call(__e, PrimFunc(symshen_4newpv), V5642)


let__14493 := tmp17253
_ = let__14493

tmp17254 := Call(__e, PrimFunc(symshen_4newpv), V5642)


let__14494 := tmp17254
_ = let__14494

tmp17255 := Call(__e, PrimFunc(symshen_4newpv), V5642)


let__14495 := tmp17255
_ = let__14495

tmp17256 := Call(__e, PrimFunc(symshen_4newpv), V5642)


let__14496 := tmp17256
_ = let__14496

tmp17257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14495, Nil)
}
__typedArg0 := let__14495
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14494, tmp17257)
}
__typedArg0 := let__14494
__typedArg1 := tmp17257
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14493, tmp17258)
}
__typedArg0 := let__14493
__typedArg1 := tmp17258
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17259, let__14496)
}
__typedArg0 := tmp17259
__typedArg1 := let__14496
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17261 := MakeNative(func(__e *ControlFlow) {
tmp17262 := Call(__e, let__14474, let__14493)


tmp17263 := Call(__e, tmp17262, let__14494)


tmp17264 := Call(__e, tmp17263, let__14495)


__e.TailApply(tmp17264, let__14496)
return


}, 0)

tmp17265 := Call(__e, PrimFunc(symshen_4bind_b), let__14473, tmp17260, V5642, tmp17261)


tmp17266 := Call(__e, PrimFunc(symshen_4gc), V5642, tmp17265)


tmp17267 := Call(__e, PrimFunc(symshen_4gc), V5642, tmp17266)


tmp17268 := Call(__e, PrimFunc(symshen_4gc), V5642, tmp17267)


__e.TailApply(PrimFunc(symshen_4gc), V5642, tmp17268)
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
__e.Return(let__14466)
return
}


}, 6)

tmp17279 := Call(__e, ns2_1set, symshen_4p_1hyps, tmp17167)


_ = tmp17279

tmp17280 := MakeNative(func(__e *ControlFlow) {
V5687 := __e.Get(1)
_ = V5687
V5688 := __e.Get(2)
_ = V5688
V5689 := __e.Get(3)
_ = V5689
V5690 := __e.Get(4)
_ = V5690
V5691 := __e.Get(5)
_ = V5691
V5692 := __e.Get(6)
_ = V5692
V5693 := __e.Get(7)
_ = V5693
let__14497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5692)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5692
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__14497

tmp17318 := Call(__e, PrimFunc(symshen_4unlocked_2), V5691)


var ifres17282 Obj

if True == tmp17318 {
tmp17283 := Call(__e, PrimFunc(symshen_4lazyderef), V5687, V5690)


let__14499 := tmp17283
_ = let__14499

tmp17317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14499)
}
__typedArg0 := let__14499
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17284 Obj

if True == tmp17317 {
tmp17285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14499)
}
__typedArg0 := let__14499
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17286 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17285, V5690)


let__14500 := tmp17286
_ = let__14500

tmp17316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14500, symwhere)
}
__typedArg0 := let__14500
__typedArg1 := symwhere
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17287 Obj

if True == tmp17316 {
tmp17288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14499)
}
__typedArg0 := let__14499
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17289 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17288, V5690)


let__14501 := tmp17289
_ = let__14501

tmp17315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14501)
}
__typedArg0 := let__14501
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17290 Obj

if True == tmp17315 {
tmp17291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14501)
}
__typedArg0 := let__14501
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14502 := tmp17291
_ = let__14502

tmp17292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14501)
}
__typedArg0 := let__14501
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17293 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17292, V5690)


let__14503 := tmp17293
_ = let__14503

tmp17314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14503)
}
__typedArg0 := let__14503
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17294 Obj

if True == tmp17314 {
tmp17295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14503)
}
__typedArg0 := let__14503
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14504 := tmp17295
_ = let__14504

tmp17296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14503)
}
__typedArg0 := let__14503
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17297 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17296, V5690)


let__14505 := tmp17297
_ = let__14505

tmp17313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14505, Nil)
}
__typedArg0 := let__14505
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17298 Obj

if True == tmp17313 {
tmp17299 := Call(__e, PrimFunc(symshen_4newpv), V5690)


let__14506 := tmp17299
_ = let__14506

tmp17300 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17300

tmp17301 := MakeNative(func(__e *ControlFlow) {
tmp17302 := Call(__e, PrimFunc(symshen_4curry), let__14502)


tmp17303 := MakeNative(func(__e *ControlFlow) {
tmp17304 := MakeNative(func(__e *ControlFlow) {
tmp17305 := MakeNative(func(__e *ControlFlow) {
tmp17306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp17307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symverified, Nil)
}
__typedArg0 := symverified
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17306, tmp17307)
}
__typedArg0 := tmp17306
__typedArg1 := tmp17307
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14506, tmp17308)
}
__typedArg0 := let__14506
__typedArg1 := tmp17308
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17309, V5689)
}
__typedArg0 := tmp17309
__typedArg1 := V5689
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4t_d_1correct), let__14504, V5688, tmp17310, V5690, V5691, let__14497, V5693)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5690, V5691, let__14497, tmp17305)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14506, symboolean, V5689, V5690, V5691, let__14497, tmp17304)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14506, tmp17302, V5690, V5691, let__14497, tmp17303)
return


}, 0)

tmp17311 := Call(__e, PrimFunc(symshen_4cut), V5690, V5691, let__14497, tmp17301)


tmp17312 := Call(__e, PrimFunc(symshen_4gc), V5690, tmp17311)


ifres17298 = tmp17312


} else {
ifres17298 = False


}

ifres17294 = ifres17298


} else {
ifres17294 = False


}

ifres17290 = ifres17294


} else {
ifres17290 = False


}

ifres17287 = ifres17290


} else {
ifres17287 = False


}

ifres17284 = ifres17287


} else {
ifres17284 = False


}

ifres17282 = ifres17284


} else {
ifres17282 = False


}

let__14498 := ifres17282
_ = let__14498

tmp17327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14498, False)
}
__typedArg0 := let__14498
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17327 {
tmp17323 := Call(__e, PrimFunc(symshen_4unlocked_2), V5691)


var ifres17319 Obj

if True == tmp17323 {
tmp17320 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17320

tmp17321 := Call(__e, PrimFunc(symshen_4curry), V5687)


tmp17322 := Call(__e, PrimFunc(symshen_4system_1S_1h), tmp17321, V5688, V5689, V5690, V5691, let__14497, V5693)


ifres17319 = tmp17322


} else {
ifres17319 = False


}

let__14507 := ifres17319
_ = let__14507

tmp17325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14507, False)
}
__typedArg0 := let__14507
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17325 {
__e.TailApply(PrimFunc(symshen_4unlock), V5691, let__14497)
return
} else {
__e.Return(let__14507)
return
}


} else {
__e.Return(let__14498)
return
}


}, 7)

tmp17328 := Call(__e, ns2_1set, symshen_4t_d_1correct, tmp17280)


_ = tmp17328

tmp17329 := MakeNative(func(__e *ControlFlow) {
V5705 := __e.Get(1)
_ = V5705
V5706 := __e.Get(2)
_ = V5706
V5707 := __e.Get(3)
_ = V5707
V5708 := __e.Get(4)
_ = V5708
V5709 := __e.Get(5)
_ = V5709
V5710 := __e.Get(6)
_ = V5710
V5711 := __e.Get(7)
_ = V5711
V5712 := __e.Get(8)
_ = V5712
tmp17336 := Call(__e, PrimFunc(symshen_4unlocked_2), V5710)


var ifres17330 Obj

if True == tmp17336 {
tmp17331 := Call(__e, PrimFunc(symshen_4lazyderef), V5705, V5709)


let__14509 := tmp17331
_ = let__14509

tmp17335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14509, Nil)
}
__typedArg0 := let__14509
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17332 Obj

if True == tmp17335 {
tmp17333 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17333

tmp17334 := Call(__e, PrimFunc(symis_b), V5706, V5708, V5709, V5710, V5711, V5712)


ifres17332 = tmp17334


} else {
ifres17332 = False


}

ifres17330 = ifres17332


} else {
ifres17330 = False


}

let__14508 := ifres17330
_ = let__14508

tmp17368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14508, False)
}
__typedArg0 := let__14508
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17368 {
tmp17366 := Call(__e, PrimFunc(symshen_4unlocked_2), V5710)


if True == tmp17366 {
tmp17337 := Call(__e, PrimFunc(symshen_4lazyderef), V5705, V5709)


let__14510 := tmp17337
_ = let__14510

tmp17364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14510)
}
__typedArg0 := let__14510
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17364 {
tmp17338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14510)
}
__typedArg0 := let__14510
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14511 := tmp17338
_ = let__14511

tmp17339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14510)
}
__typedArg0 := let__14510
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14512 := tmp17339
_ = let__14512

tmp17340 := Call(__e, PrimFunc(symshen_4lazyderef), V5706, V5709)


let__14513 := tmp17340
_ = let__14513

tmp17362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14513)
}
__typedArg0 := let__14513
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17362 {
tmp17341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14513)
}
__typedArg0 := let__14513
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14514 := tmp17341
_ = let__14514

tmp17342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14513)
}
__typedArg0 := let__14513
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17343 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17342, V5709)


let__14515 := tmp17343
_ = let__14515

tmp17360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14515)
}
__typedArg0 := let__14515
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17360 {
tmp17344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14515)
}
__typedArg0 := let__14515
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17345 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17344, V5709)


let__14516 := tmp17345
_ = let__14516

tmp17358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14516, sym_1_1_6)
}
__typedArg0 := let__14516
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17358 {
tmp17346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14515)
}
__typedArg0 := let__14515
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17347 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17346, V5709)


let__14517 := tmp17347
_ = let__14517

tmp17356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14517)
}
__typedArg0 := let__14517
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17356 {
tmp17348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14517)
}
__typedArg0 := let__14517
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14518 := tmp17348
_ = let__14518

tmp17349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14517)
}
__typedArg0 := let__14517
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17350 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17349, V5709)


let__14519 := tmp17350
_ = let__14519

tmp17354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14519, Nil)
}
__typedArg0 := let__14519
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17354 {
tmp17351 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17351

tmp17352 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1integrity), let__14512, let__14518, V5707, V5708, V5709, V5710, V5711, V5712)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14511, let__14514, V5707, V5709, V5710, V5711, tmp17352)
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
__e.Return(let__14508)
return
}


}, 8)

tmp17369 := Call(__e, ns2_1set, symshen_4t_d_1integrity, tmp17329)


_ = tmp17369

tmp17370 := MakeNative(func(__e *ControlFlow) {
V5725 := __e.Get(1)
_ = V5725
tmp17379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V5725)
}
__typedArg0 := V5725
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp17379 {
tmp17376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V5725)
}
__typedArg0 := V5725
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

tmp17377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp17376)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp17376
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres17372 Obj

if True == tmp17377 {
tmp17374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V5725, MakeInteger(0))
}
__typedArg0 := V5725
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp17375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp17374, symshen_4print_1freshterm)
}
__typedArg0 := tmp17374
__typedArg1 := symshen_4print_1freshterm
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17373 Obj

if True == tmp17375 {
ifres17373 = True


} else {
ifres17373 = False


}

ifres17372 = ifres17373


} else {
ifres17372 = False


}

if True == ifres17372 {
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

__e.TailApply(ns2_1set, symshen_4freshterm_2, tmp17370)
return




}, 0)

