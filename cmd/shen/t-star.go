package main

import . "github.com/pyrex41/shen-go/kl"

var TStarMain = MakeNative(func(__e *ControlFlow) {
tmp14520 := MakeNative(func(__e *ControlFlow) {
V4901 := __e.Get(1)
_ = V4901
V4902 := __e.Get(2)
_ = V4902
tmp14521 := Call(__e, PrimFunc(symshen_4extract_1vars), V4902)


let__13969 := tmp14521
_ = let__13969

tmp14522 := Call(__e, PrimFunc(symshen_4rectify_1type), V4902)


let__13970 := tmp14522
_ = let__13970

tmp14523 := Call(__e, PrimFunc(symshen_4curry), V4901)


let__13971 := tmp14523
_ = let__13971

tmp14524 := MakeNative(func(__e *ControlFlow) {
Z4906 := __e.Get(1)
_ = Z4906
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4907 := __e.Get(1)
_ = Z4907
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4908 := __e.Get(1)
_ = Z4908
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4909 := __e.Get(1)
_ = Z4909
tmp14525 := Call(__e, PrimFunc(symshen_4newpv), Z4906)


let__13972 := tmp14525
_ = let__13972

tmp14526 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14526

tmp14527 := Call(__e, PrimFunc(symshen_4deref), let__13969, Z4906)


tmp14528 := Call(__e, PrimFunc(symreceive), tmp14527)


tmp14529 := Call(__e, PrimFunc(symshen_4deref), let__13970, Z4906)


tmp14530 := Call(__e, PrimFunc(symreceive), tmp14529)


tmp14531 := MakeNative(func(__e *ControlFlow) {
tmp14532 := Call(__e, PrimFunc(symshen_4deref), let__13971, Z4906)


tmp14533 := Call(__e, PrimFunc(symreceive), tmp14532)


tmp14534 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symreturn), let__13972, Z4906, Z4907, Z4908, Z4909)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4toplevel_1forms), tmp14533, let__13972, Z4906, Z4907, Z4908, tmp14534)
return


}, 0)

tmp14535 := Call(__e, PrimFunc(symshen_4insert_1prolog_1variables), tmp14528, tmp14530, let__13972, Z4906, Z4907, Z4908, tmp14531)


__e.TailApply(PrimFunc(symshen_4gc), Z4906, tmp14535)
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
V4911 := __e.Get(1)
_ = V4911
V4912 := __e.Get(2)
_ = V4912
V4913 := __e.Get(3)
_ = V4913
V4914 := __e.Get(4)
_ = V4914
V4915 := __e.Get(5)
_ = V4915
V4916 := __e.Get(6)
_ = V4916
V4917 := __e.Get(7)
_ = V4917
tmp14552 := Call(__e, PrimFunc(symshen_4unlocked_2), V4915)


var ifres14546 Obj

if True == tmp14552 {
tmp14547 := Call(__e, PrimFunc(symshen_4lazyderef), V4911, V4914)


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

tmp14550 := Call(__e, PrimFunc(symis_b), V4912, V4913, V4914, V4915, V4916, V4917)


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
tmp14564 := Call(__e, PrimFunc(symshen_4unlocked_2), V4915)


if True == tmp14564 {
tmp14553 := Call(__e, PrimFunc(symshen_4lazyderef), V4911, V4914)


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

tmp14556 := Call(__e, PrimFunc(symshen_4newpv), V4914)


let__13978 := tmp14556
_ = let__13978

tmp14557 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14557

tmp14558 := Call(__e, PrimFunc(symshen_4deref), let__13978, V4914)


tmp14559 := Call(__e, PrimFunc(symsubst), tmp14558, let__13976, V4912)


tmp14560 := Call(__e, PrimFunc(symshen_4insert_1prolog_1variables), let__13977, tmp14559, V4913, V4914, V4915, V4916, V4917)


__e.TailApply(PrimFunc(symshen_4gc), V4914, tmp14560)
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
V4924 := __e.Get(1)
_ = V4924
V4925 := __e.Get(2)
_ = V4925
V4926 := __e.Get(3)
_ = V4926
V4927 := __e.Get(4)
_ = V4927
V4928 := __e.Get(5)
_ = V4928
V4929 := __e.Get(6)
_ = V4929
let__13979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4928)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4928
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__13979

tmp14593 := Call(__e, PrimFunc(symshen_4unlocked_2), V4927)


var ifres14570 Obj

if True == tmp14593 {
tmp14571 := Call(__e, PrimFunc(symshen_4lazyderef), V4924, V4926)


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

tmp14574 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14573, V4926)


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

tmp14577 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14576, V4926)


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

__e.TailApply(PrimFunc(symshen_4t_d), tmp14588, V4925, V4926, V4927, let__13979, V4929)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4signal_1def), tmp14585, let__13984, V4926, V4927, let__13979, tmp14586)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4926, V4927, let__13979, tmp14584)
return


}, 0)

tmp14589 := Call(__e, PrimFunc(symwhen), tmp14582, V4926, V4927, let__13979, tmp14583)


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
tmp14601 := Call(__e, PrimFunc(symshen_4unlocked_2), V4927)


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
return PrimCons(V4925, Nil)
}
__typedArg0 := V4925
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
return PrimCons(V4924, tmp14598)
}
__typedArg0 := V4924
__typedArg1 := tmp14598
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14600 := Call(__e, PrimFunc(symshen_4system_1S), tmp14599, Nil, V4926, V4927, let__13979, V4929)


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
__e.TailApply(PrimFunc(symshen_4unlock), V4927, let__13979)
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
tmp14614 := Call(__e, PrimFunc(symshen_4unlocked_2), V4941)


var ifres14608 Obj

if True == tmp14614 {
tmp14609 := Call(__e, PrimFunc(symshen_4lazyderef), V4938, V4940)


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

tmp14612 := Call(__e, PrimFunc(symthaw), V4943)


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
tmp14627 := Call(__e, PrimFunc(symshen_4unlocked_2), V4941)


if True == tmp14627 {
tmp14615 := Call(__e, PrimFunc(symshen_4lazyderef), V4938, V4940)


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
tmp14616 := Call(__e, PrimFunc(symshen_4newpv), V4940)


let__13990 := tmp14616
_ = let__13990

tmp14617 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14617

tmp14618 := Call(__e, PrimFunc(symshen_4deref), V4939, V4940)


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


tmp14623 := Call(__e, PrimFunc(symis), let__13990, tmp14622, V4940, V4941, V4942, V4943)


__e.TailApply(PrimFunc(symshen_4gc), V4940, tmp14623)
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
V4948 := __e.Get(1)
_ = V4948
tmp14632 := Call(__e, PrimFunc(symshen_4curry_1type), V4948)


__e.TailApply(PrimFunc(symshen_4demodulate), tmp14632)
return


}, 1)

tmp14633 := Call(__e, ns2_1set, symshen_4rectify_1type, tmp14631)


_ = tmp14633

tmp14634 := MakeNative(func(__e *ControlFlow) {
V4949 := __e.Get(1)
_ = V4949
tmp14635 := MakeNative(func(__e *ControlFlow) {
tmp14636 := MakeNative(func(__e *ControlFlow) {
Z4951 := __e.Get(1)
_ = Z4951
__e.TailApply(PrimFunc(symshen_4demod), Z4951)
return
}, 1)

tmp14637 := Call(__e, PrimFunc(symshen_4walk), tmp14636, V4949)


let__13991 := tmp14637
_ = let__13991

tmp14639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13991, V4949)
}
__typedArg0 := let__13991
__typedArg1 := V4949
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14639 {
__e.Return(V4949)
return
} else {
__e.TailApply(PrimFunc(symshen_4demodulate), let__13991)
return
}


}, 0)

tmp14640 := MakeNative(func(__e *ControlFlow) {
Z4952 := __e.Get(1)
_ = Z4952
__e.Return(V4949)
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
V4953 := __self1
_ = V4953
tmp14766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4953)
}
__typedArg0 := V4953
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14739 Obj

if True == tmp14766 {
tmp14764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4953)
}
__typedArg0 := V4953
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
return PrimTail(V4953)
}
__typedArg0 := V4953
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
return PrimTail(V4953)
}
__typedArg0 := V4953
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
return PrimTail(V4953)
}
__typedArg0 := V4953
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
return PrimTail(V4953)
}
__typedArg0 := V4953
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
return PrimHead(V4953)
}
__typedArg0 := V4953
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4953)
}
__typedArg0 := V4953
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
return PrimIsPair(V4953)
}
__typedArg0 := V4953
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14697 Obj

if True == tmp14737 {
tmp14735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4953)
}
__typedArg0 := V4953
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
return PrimHead(V4953)
}
__typedArg0 := V4953
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
return PrimHead(V4953)
}
__typedArg0 := V4953
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
return PrimHead(V4953)
}
__typedArg0 := V4953
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
return PrimTail(V4953)
}
__typedArg0 := V4953
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
return PrimTail(V4953)
}
__typedArg0 := V4953
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
return PrimTail(V4953)
}
__typedArg0 := V4953
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
return PrimTail(V4953)
}
__typedArg0 := V4953
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
return PrimHead(V4953)
}
__typedArg0 := V4953
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4953)
}
__typedArg0 := V4953
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4953)
}
__typedArg0 := V4953
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
return PrimIsPair(V4953)
}
__typedArg0 := V4953
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14668 Obj

if True == tmp14695 {
tmp14693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4953)
}
__typedArg0 := V4953
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
return PrimTail(V4953)
}
__typedArg0 := V4953
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
return PrimTail(V4953)
}
__typedArg0 := V4953
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
return PrimTail(V4953)
}
__typedArg0 := V4953
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
return PrimTail(V4953)
}
__typedArg0 := V4953
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
return PrimHead(V4953)
}
__typedArg0 := V4953
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4953)
}
__typedArg0 := V4953
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
return PrimIsPair(V4953)
}
__typedArg0 := V4953
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14666 {
tmp14664 := MakeNative(func(__e *ControlFlow) {
Z4954 := __e.Get(1)
_ = Z4954
__e.TailApply(PrimFunc(symshen_4curry_1type), Z4954)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp14664, V4953)
return


} else {
__e.Return(V4953)
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
V4955 := __self1
_ = V4955
tmp14857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4955)
}
__typedArg0 := V4955
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14849 Obj

if True == tmp14857 {
tmp14855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4955)
}
__typedArg0 := V4955
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
return PrimTail(V4955)
}
__typedArg0 := V4955
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
__e.Return(V4955)
return
} else {
tmp14847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4955)
}
__typedArg0 := V4955
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14828 Obj

if True == tmp14847 {
tmp14845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4955)
}
__typedArg0 := V4955
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
return PrimTail(V4955)
}
__typedArg0 := V4955
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
return PrimTail(V4955)
}
__typedArg0 := V4955
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
return PrimTail(V4955)
}
__typedArg0 := V4955
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
return PrimTail(V4955)
}
__typedArg0 := V4955
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
return PrimTail(V4955)
}
__typedArg0 := V4955
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
return PrimIsPair(V4955)
}
__typedArg0 := V4955
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14822 Obj

if True == tmp14826 {
tmp14824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4955)
}
__typedArg0 := V4955
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
return PrimHead(V4955)
}
__typedArg0 := V4955
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14776 := MakeNative(func(__e *ControlFlow) {
Z4956 := __e.Get(1)
_ = Z4956
__e.TailApply(PrimFunc(symshen_4curry), Z4956)
return
}, 1)

tmp14777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4955)
}
__typedArg0 := V4955
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
return PrimIsPair(V4955)
}
__typedArg0 := V4955
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14816 Obj

if True == tmp14820 {
tmp14818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4955)
}
__typedArg0 := V4955
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
__e.Return(V4955)
return
} else {
tmp14814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4955)
}
__typedArg0 := V4955
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14805 Obj

if True == tmp14814 {
tmp14812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4955)
}
__typedArg0 := V4955
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
return PrimTail(V4955)
}
__typedArg0 := V4955
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
return PrimHead(V4955)
}
__typedArg0 := V4955
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4955)
}
__typedArg0 := V4955
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
return PrimTail(V4955)
}
__typedArg0 := V4955
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
return PrimIsPair(V4955)
}
__typedArg0 := V4955
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14794 Obj

if True == tmp14803 {
tmp14801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4955)
}
__typedArg0 := V4955
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
return PrimTail(V4955)
}
__typedArg0 := V4955
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
return PrimHead(V4955)
}
__typedArg0 := V4955
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14788 := Call(__e, PrimFunc(symshen_4curry), tmp14787)


tmp14789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4955)
}
__typedArg0 := V4955
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
__e.Return(V4955)
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
V4957 := __e.Get(1)
_ = V4957
tmp14860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspecial_d)
}
__typedArg0 := symshen_4_dspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symelement_2), V4957, tmp14860)
return


}, 1)

tmp14861 := Call(__e, ns2_1set, symshen_4special_2, tmp14859)


_ = tmp14861

tmp14862 := MakeNative(func(__e *ControlFlow) {
V4958 := __e.Get(1)
_ = V4958
tmp14863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dextraspecial_d)
}
__typedArg0 := symshen_4_dextraspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symelement_2), V4958, tmp14863)
return


}, 1)

tmp14864 := Call(__e, ns2_1set, symshen_4extraspecial_2, tmp14862)


_ = tmp14864

tmp14865 := MakeNative(func(__e *ControlFlow) {
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
let__13992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4963)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4963
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__13992

tmp14871 := Call(__e, PrimFunc(symshen_4unlocked_2), V4962)


var ifres14867 Obj

if True == tmp14871 {
tmp14868 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14868

tmp14869 := Call(__e, PrimFunc(symshen_4maxinfexceeded_2))


tmp14870 := Call(__e, PrimFunc(symwhen), tmp14869, V4961, V4962, let__13992, V4964)


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
tmp14900 := Call(__e, PrimFunc(symshen_4unlocked_2), V4962)


var ifres14872 Obj

if True == tmp14900 {
tmp14873 := Call(__e, PrimFunc(symshen_4lazyderef), V4959, V4961)


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

tmp14877 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14876, V4961)


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

tmp14881 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14880, V4961)


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

tmp14885 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14884, V4961)


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

tmp14888 := Call(__e, PrimFunc(symshen_4deref), let__13998, V4961)


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
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__13996, let__14000, V4960, V4961, V4962, let__13992, V4964)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4961, V4962, let__13992, tmp14894)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14892, V4961, V4962, let__13992, tmp14893)
return


}, 0)

tmp14895 := Call(__e, PrimFunc(symwhen), tmp14890, V4961, V4962, let__13992, tmp14891)


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
tmp14906 := Call(__e, PrimFunc(symshen_4unlocked_2), V4962)


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
__e.TailApply(PrimFunc(symshen_4show), V4959, V4960, V4961, V4962, let__13992, V4964)
return
}, 0)

tmp14905 := Call(__e, PrimFunc(symwhen), tmp14903, V4961, V4962, let__13992, tmp14904)


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
tmp14911 := Call(__e, PrimFunc(symshen_4unlocked_2), V4962)


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

tmp14910 := Call(__e, PrimFunc(symshen_4search_1user_1datatypes), V4959, V4960, tmp14909, V4961, V4962, let__13992, V4964)


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
__e.TailApply(PrimFunc(symshen_4unlock), V4962, let__13992)
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
V4983 := __e.Get(1)
_ = V4983
V4984 := __e.Get(2)
_ = V4984
V4985 := __e.Get(3)
_ = V4985
V4986 := __e.Get(4)
_ = V4986
V4987 := __e.Get(5)
_ = V4987
V4988 := __e.Get(6)
_ = V4988
tmp14922 := Call(__e, PrimFunc(symshen_4line))


_ = tmp14922

tmp14923 := Call(__e, PrimFunc(symshen_4deref), V4983, V4985)


tmp14924 := Call(__e, PrimFunc(symshen_4show_1p), tmp14923)


_ = tmp14924

tmp14925 := Call(__e, PrimFunc(symnl), MakeInteger(2))


_ = tmp14925

tmp14926 := Call(__e, PrimFunc(symshen_4deref), V4984, V4985)


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
V4990 := __e.Get(1)
_ = V4990
tmp14972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4990)
}
__typedArg0 := V4990
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14951 Obj

if True == tmp14972 {
tmp14970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4990)
}
__typedArg0 := V4990
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
return PrimTail(V4990)
}
__typedArg0 := V4990
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
return PrimTail(V4990)
}
__typedArg0 := V4990
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
return PrimTail(V4990)
}
__typedArg0 := V4990
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
return PrimHead(V4990)
}
__typedArg0 := V4990
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14942 := Call(__e, PrimFunc(symshen_4prterm), tmp14941)


_ = tmp14942

tmp14943 := Call(__e, PrimFunc(symstoutput))


tmp14944 := Call(__e, PrimFunc(sympr), MakeString(" : "), tmp14943)


_ = tmp14944

tmp14945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4990)
}
__typedArg0 := V4990
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
__e.TailApply(PrimFunc(symshen_4prterm), V4990)
return
}


}, 1)

tmp14973 := Call(__e, ns2_1set, symshen_4show_1p, tmp14940)


_ = tmp14973

tmp14974 := MakeNative(func(__e *ControlFlow) {
V4991 := __e.Get(1)
_ = V4991
tmp15017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4991)
}
__typedArg0 := V4991
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14998 Obj

if True == tmp15017 {
tmp15015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4991)
}
__typedArg0 := V4991
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
return PrimTail(V4991)
}
__typedArg0 := V4991
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
return PrimTail(V4991)
}
__typedArg0 := V4991
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
return PrimTail(V4991)
}
__typedArg0 := V4991
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
return PrimTail(V4991)
}
__typedArg0 := V4991
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
return PrimTail(V4991)
}
__typedArg0 := V4991
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
return PrimIsPair(V4991)
}
__typedArg0 := V4991
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14996 {
tmp14985 := Call(__e, PrimFunc(symstoutput))


tmp14986 := Call(__e, PrimFunc(sympr), MakeString("("), tmp14985)


_ = tmp14986

tmp14987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4991)
}
__typedArg0 := V4991
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14988 := Call(__e, PrimFunc(symshen_4prterm), tmp14987)


_ = tmp14988

tmp14989 := MakeNative(func(__e *ControlFlow) {
Z4992 := __e.Get(1)
_ = Z4992
tmp14990 := Call(__e, PrimFunc(symstoutput))


tmp14991 := Call(__e, PrimFunc(sympr), MakeString(" "), tmp14990)


_ = tmp14991

__e.TailApply(PrimFunc(symshen_4prterm), Z4992)
return


}, 1)

tmp14992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4991)
}
__typedArg0 := V4991
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14993 := Call(__e, PrimFunc(symmap), tmp14989, tmp14992)


_ = tmp14993

tmp14994 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString(")"), tmp14994)
return


} else {
__e.TailApply(PrimFunc(symprint), V4991)
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
V4993 := __self1
_ = V4993
tmp15052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4993)
}
__typedArg0 := Nil
__typedArg1 := V4993
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15052 {
__e.Return(MakeString(""))
return
} else {
tmp15050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4993)
}
__typedArg0 := V4993
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15031 Obj

if True == tmp15050 {
tmp15048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4993)
}
__typedArg0 := V4993
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
return PrimTail(V4993)
}
__typedArg0 := V4993
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
return PrimTail(V4993)
}
__typedArg0 := V4993
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
return PrimTail(V4993)
}
__typedArg0 := V4993
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
return PrimTail(V4993)
}
__typedArg0 := V4993
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
return PrimTail(V4993)
}
__typedArg0 := V4993
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

__e.TailApply(PrimFunc(symshen_4prterm), V4993)
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
V5000 := __self1
_ = V5000
V5001 := __self2
_ = V5001
tmp15067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5000)
}
__typedArg0 := Nil
__typedArg1 := V5000
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15067 {
tmp15055 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString("\n> "), tmp15055)
return


} else {
tmp15065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5000)
}
__typedArg0 := V5000
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15065 {
tmp15056 := Call(__e, PrimFunc(symshen_4app), V5001, MakeString(". "), symshen_4a)


tmp15057 := Call(__e, PrimFunc(symstoutput))


tmp15058 := Call(__e, PrimFunc(sympr), tmp15056, tmp15057)


_ = tmp15058

tmp15059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5000)
}
__typedArg0 := V5000
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15060 := Call(__e, PrimFunc(symshen_4show_1p), tmp15059)


_ = tmp15060

tmp15061 := Call(__e, PrimFunc(symnl), MakeInteger(1))


_ = tmp15061

tmp15062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5000)
}
__typedArg0 := V5000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4show_1assumptions) == __self {
__self1, __self2 = tmp15062, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5001)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5001
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4show_1assumptions), tmp15062, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5001)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5001
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
V5003 := __e.Get(1)
_ = V5003
V5004 := __e.Get(2)
_ = V5004
V5005 := __e.Get(3)
_ = V5005
V5006 := __e.Get(4)
_ = V5006
V5007 := __e.Get(5)
_ = V5007
V5008 := __e.Get(6)
_ = V5008
V5009 := __e.Get(7)
_ = V5009
let__14006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5008)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5008
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__14006

tmp15094 := Call(__e, PrimFunc(symshen_4unlocked_2), V5007)


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
return PrimCons(V5004, Nil)
}
__typedArg0 := V5004
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
return PrimCons(V5003, tmp15091)
}
__typedArg0 := V5003
__typedArg1 := tmp15091
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4show), tmp15092, V5005, V5006, V5007, let__14006, V5009)
return


}, 0)

tmp15093 := Call(__e, PrimFunc(symwhen), tmp15087, V5006, V5007, let__14006, tmp15088)


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
tmp15102 := Call(__e, PrimFunc(symshen_4unlocked_2), V5007)


var ifres15095 Obj

if True == tmp15102 {
tmp15096 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15096

tmp15097 := Call(__e, PrimFunc(symshen_4lazyderef), V5003, V5006)


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
__e.TailApply(PrimFunc(symshen_4primitive), V5003, V5004, V5006, V5007, let__14006, V5009)
return
}, 0)

tmp15101 := Call(__e, PrimFunc(symwhen), tmp15099, V5006, V5007, let__14006, tmp15100)


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
tmp15106 := Call(__e, PrimFunc(symshen_4unlocked_2), V5007)


var ifres15103 Obj

if True == tmp15106 {
tmp15104 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15104

tmp15105 := Call(__e, PrimFunc(symshen_4by_1hypothesis), V5003, V5004, V5005, V5006, V5007, let__14006, V5009)


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
tmp15120 := Call(__e, PrimFunc(symshen_4unlocked_2), V5007)


var ifres15107 Obj

if True == tmp15120 {
tmp15108 := Call(__e, PrimFunc(symshen_4lazyderef), V5003, V5006)


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

tmp15112 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15111, V5006)


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
return PrimCons(V5004, Nil)
}
__typedArg0 := V5004
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

tmp15117 := Call(__e, PrimFunc(symshen_4lookupsig), let__14012, tmp15116, V5006, V5007, let__14006, V5009)


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
tmp15146 := Call(__e, PrimFunc(symshen_4unlocked_2), V5007)


var ifres15121 Obj

if True == tmp15146 {
tmp15122 := Call(__e, PrimFunc(symshen_4lazyderef), V5003, V5006)


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

tmp15125 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15124, V5006)


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

tmp15128 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15127, V5006)


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

tmp15132 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15131, V5006)


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

tmp15135 := Call(__e, PrimFunc(symshen_4deref), let__14018, V5006)


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

__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp15140, V5004, V5005, V5006, V5007, let__14006, V5009)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5006, V5007, let__14006, tmp15139)
return


}, 0)

tmp15141 := Call(__e, PrimFunc(symwhen), tmp15137, V5006, V5007, let__14006, tmp15138)


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
tmp15166 := Call(__e, PrimFunc(symshen_4unlocked_2), V5007)


var ifres15147 Obj

if True == tmp15166 {
tmp15148 := Call(__e, PrimFunc(symshen_4lazyderef), V5003, V5006)


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

tmp15151 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15150, V5006)


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

tmp15154 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15153, V5006)


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

tmp15158 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15157, V5006)


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

tmp15161 := Call(__e, PrimFunc(symshen_4lookupsig), let__14024, V5004, V5006, V5007, let__14006, V5009)


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
tmp15193 := Call(__e, PrimFunc(symshen_4unlocked_2), V5007)


var ifres15167 Obj

if True == tmp15193 {
tmp15168 := Call(__e, PrimFunc(symshen_4lazyderef), V5003, V5006)


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

tmp15172 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15171, V5006)


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

tmp15176 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15175, V5006)


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
tmp15178 := Call(__e, PrimFunc(symshen_4newpv), V5006)


let__14032 := tmp15178
_ = let__14032

tmp15179 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15179

tmp15180 := Call(__e, PrimFunc(symshen_4lazyderef), let__14028, V5006)


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
return PrimCons(V5004, Nil)
}
__typedArg0 := V5004
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
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14030, let__14032, V5005, V5006, V5007, let__14006, V5009)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4lookupsig), let__14028, tmp15186, V5006, V5007, let__14006, tmp15187)
return


}, 0)

tmp15188 := Call(__e, PrimFunc(symwhen), tmp15182, V5006, V5007, let__14006, tmp15183)


tmp15189 := Call(__e, PrimFunc(symshen_4gc), V5006, tmp15188)


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
tmp15216 := Call(__e, PrimFunc(symshen_4unlocked_2), V5007)


var ifres15194 Obj

if True == tmp15216 {
tmp15195 := Call(__e, PrimFunc(symshen_4lazyderef), V5003, V5006)


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

tmp15199 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15198, V5006)


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

tmp15203 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15202, V5006)


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
tmp15205 := Call(__e, PrimFunc(symshen_4newpv), V5006)


let__14039 := tmp15205
_ = let__14039

tmp15206 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15206

tmp15207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5004, Nil)
}
__typedArg0 := V5004
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
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14037, let__14039, V5005, V5006, V5007, let__14006, V5009)
return
}, 0)

tmp15211 := Call(__e, PrimFunc(symshen_4system_1S_1h), let__14035, tmp15209, V5005, V5006, V5007, let__14006, tmp15210)


tmp15212 := Call(__e, PrimFunc(symshen_4gc), V5006, tmp15211)


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
tmp15283 := Call(__e, PrimFunc(symshen_4unlocked_2), V5007)


var ifres15217 Obj

if True == tmp15283 {
tmp15218 := Call(__e, PrimFunc(symshen_4lazyderef), V5003, V5006)


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

tmp15221 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15220, V5006)


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

tmp15224 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15223, V5006)


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

tmp15228 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15227, V5006)


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

tmp15232 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15231, V5006)


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
tmp15234 := Call(__e, PrimFunc(symshen_4lazyderef), V5004, V5006)


let__14048 := tmp15234
_ = let__14048

tmp15235 := MakeNative(func(__e *ControlFlow) {
Z5054 := __e.Get(1)
_ = Z5054
tmp15236 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15236

tmp15237 := MakeNative(func(__e *ControlFlow) {
tmp15238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z5054, Nil)
}
__typedArg0 := Z5054
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

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14046, tmp15239, V5005, V5006, V5007, let__14006, V5009)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14044, Z5054, V5005, V5006, V5007, let__14006, tmp15237)
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

tmp15242 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15241, V5006)


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

tmp15245 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15244, V5006)


let__14052 := tmp15245
_ = let__14052

tmp15246 := MakeNative(func(__e *ControlFlow) {
Z5059 := __e.Get(1)
_ = Z5059
__e.TailApply(let__14049, Z5059)
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

tmp15249 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15248, V5006)


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
__e.TailApply(PrimFunc(symshen_4bind_b), let__14055, Nil, V5006, let__14056)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15260 := Call(__e, PrimFunc(symshen_4pvar_2), let__14052)


if True == tmp15260 {
tmp15255 := Call(__e, PrimFunc(symshen_4newpv), V5006)


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

tmp15258 := Call(__e, PrimFunc(symshen_4bind_b), let__14052, tmp15256, V5006, tmp15257)


__e.TailApply(PrimFunc(symshen_4gc), V5006, tmp15258)
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
tmp15266 := Call(__e, PrimFunc(symshen_4bind_b), let__14050, symlist, V5006, let__14051)


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
tmp15270 := Call(__e, PrimFunc(symshen_4newpv), V5006)


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

tmp15274 := Call(__e, PrimFunc(symshen_4bind_b), let__14048, tmp15272, V5006, tmp15273)


tmp15275 := Call(__e, PrimFunc(symshen_4gc), V5006, tmp15274)


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
tmp15367 := Call(__e, PrimFunc(symshen_4unlocked_2), V5007)


var ifres15284 Obj

if True == tmp15367 {
tmp15285 := Call(__e, PrimFunc(symshen_4lazyderef), V5003, V5006)


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

tmp15288 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15287, V5006)


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

tmp15291 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15290, V5006)


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

tmp15295 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15294, V5006)


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

tmp15299 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15298, V5006)


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
tmp15301 := Call(__e, PrimFunc(symshen_4lazyderef), V5004, V5006)


let__14067 := tmp15301
_ = let__14067

tmp15302 := MakeNative(func(__e *ControlFlow) {
Z5075 := __e.Get(1)
_ = Z5075
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5076 := __e.Get(1)
_ = Z5076
tmp15303 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15303

tmp15304 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14065, Z5076, V5005, V5006, V5007, let__14006, V5009)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14063, Z5075, V5005, V5006, V5007, let__14006, tmp15304)
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

tmp15308 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15307, V5006)


let__14070 := tmp15308
_ = let__14070

tmp15309 := MakeNative(func(__e *ControlFlow) {
Z5080 := __e.Get(1)
_ = Z5080
tmp15310 := Call(__e, let__14068, let__14069)


__e.TailApply(tmp15310, Z5080)
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

tmp15313 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15312, V5006)


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

tmp15316 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15315, V5006)


let__14074 := tmp15316
_ = let__14074

tmp15317 := MakeNative(func(__e *ControlFlow) {
Z5085 := __e.Get(1)
_ = Z5085
__e.TailApply(let__14071, Z5085)
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

tmp15320 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15319, V5006)


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
__e.TailApply(PrimFunc(symshen_4bind_b), let__14077, Nil, V5006, let__14078)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15331 := Call(__e, PrimFunc(symshen_4pvar_2), let__14074)


if True == tmp15331 {
tmp15326 := Call(__e, PrimFunc(symshen_4newpv), V5006)


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

tmp15329 := Call(__e, PrimFunc(symshen_4bind_b), let__14074, tmp15327, V5006, tmp15328)


__e.TailApply(PrimFunc(symshen_4gc), V5006, tmp15329)
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
tmp15337 := Call(__e, PrimFunc(symshen_4bind_b), let__14072, sym_d, V5006, let__14073)


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
tmp15341 := Call(__e, PrimFunc(symshen_4newpv), V5006)


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

tmp15345 := Call(__e, PrimFunc(symshen_4bind_b), let__14070, tmp15343, V5006, tmp15344)


tmp15346 := Call(__e, PrimFunc(symshen_4gc), V5006, tmp15345)


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
tmp15350 := Call(__e, PrimFunc(symshen_4newpv), V5006)


let__14081 := tmp15350
_ = let__14081

tmp15351 := Call(__e, PrimFunc(symshen_4newpv), V5006)


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

tmp15357 := Call(__e, PrimFunc(symshen_4bind_b), let__14067, tmp15354, V5006, tmp15355)


tmp15358 := Call(__e, PrimFunc(symshen_4gc), V5006, tmp15357)


tmp15359 := Call(__e, PrimFunc(symshen_4gc), V5006, tmp15358)


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
tmp15434 := Call(__e, PrimFunc(symshen_4unlocked_2), V5007)


var ifres15368 Obj

if True == tmp15434 {
tmp15369 := Call(__e, PrimFunc(symshen_4lazyderef), V5003, V5006)


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

tmp15372 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15371, V5006)


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

tmp15375 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15374, V5006)


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

tmp15379 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15378, V5006)


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

tmp15383 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15382, V5006)


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
tmp15385 := Call(__e, PrimFunc(symshen_4lazyderef), V5004, V5006)


let__14091 := tmp15385
_ = let__14091

tmp15386 := MakeNative(func(__e *ControlFlow) {
Z5103 := __e.Get(1)
_ = Z5103
tmp15387 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15387

tmp15388 := MakeNative(func(__e *ControlFlow) {
tmp15389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z5103, Nil)
}
__typedArg0 := Z5103
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

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14089, tmp15390, V5005, V5006, V5007, let__14006, V5009)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14087, Z5103, V5005, V5006, V5007, let__14006, tmp15388)
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

tmp15393 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15392, V5006)


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

tmp15396 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15395, V5006)


let__14095 := tmp15396
_ = let__14095

tmp15397 := MakeNative(func(__e *ControlFlow) {
Z5108 := __e.Get(1)
_ = Z5108
__e.TailApply(let__14092, Z5108)
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

tmp15400 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15399, V5006)


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
__e.TailApply(PrimFunc(symshen_4bind_b), let__14098, Nil, V5006, let__14099)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15411 := Call(__e, PrimFunc(symshen_4pvar_2), let__14095)


if True == tmp15411 {
tmp15406 := Call(__e, PrimFunc(symshen_4newpv), V5006)


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

tmp15409 := Call(__e, PrimFunc(symshen_4bind_b), let__14095, tmp15407, V5006, tmp15408)


__e.TailApply(PrimFunc(symshen_4gc), V5006, tmp15409)
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
tmp15417 := Call(__e, PrimFunc(symshen_4bind_b), let__14093, symvector, V5006, let__14094)


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
tmp15421 := Call(__e, PrimFunc(symshen_4newpv), V5006)


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

tmp15425 := Call(__e, PrimFunc(symshen_4bind_b), let__14091, tmp15423, V5006, tmp15424)


tmp15426 := Call(__e, PrimFunc(symshen_4gc), V5006, tmp15425)


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
tmp15467 := Call(__e, PrimFunc(symshen_4unlocked_2), V5007)


var ifres15435 Obj

if True == tmp15467 {
tmp15436 := Call(__e, PrimFunc(symshen_4lazyderef), V5003, V5006)


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

tmp15439 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15438, V5006)


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

tmp15442 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15441, V5006)


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

tmp15446 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15445, V5006)


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

tmp15450 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15449, V5006)


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
tmp15452 := Call(__e, PrimFunc(symshen_4lazyderef), V5004, V5006)


let__14110 := tmp15452
_ = let__14110

tmp15453 := MakeNative(func(__e *ControlFlow) {
tmp15454 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15454

tmp15455 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14108, symstring, V5005, V5006, V5007, let__14006, V5009)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14106, symstring, V5005, V5006, V5007, let__14006, tmp15455)
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
tmp15459 := Call(__e, PrimFunc(symshen_4bind_b), let__14110, symstring, V5006, let__14111)


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
tmp15567 := Call(__e, PrimFunc(symshen_4unlocked_2), V5007)


var ifres15468 Obj

if True == tmp15567 {
tmp15469 := Call(__e, PrimFunc(symshen_4lazyderef), V5003, V5006)


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

tmp15472 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15471, V5006)


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

tmp15475 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15474, V5006)


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

tmp15479 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15478, V5006)


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

tmp15483 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15482, V5006)


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
tmp15485 := Call(__e, PrimFunc(symshen_4lazyderef), V5004, V5006)


let__14120 := tmp15485
_ = let__14120

tmp15486 := MakeNative(func(__e *ControlFlow) {
Z5134 := __e.Get(1)
_ = Z5134
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5135 := __e.Get(1)
_ = Z5135
tmp15487 := Call(__e, PrimFunc(symshen_4newpv), V5006)


let__14122 := tmp15487
_ = let__14122

tmp15488 := Call(__e, PrimFunc(symshen_4newpv), V5006)


let__14123 := tmp15488
_ = let__14123

tmp15489 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15489

tmp15490 := Call(__e, PrimFunc(symshen_4lazyderef), let__14116, V5006)


tmp15491 := Call(__e, PrimFunc(symshen_4freshterm), tmp15490)


tmp15492 := MakeNative(func(__e *ControlFlow) {
tmp15493 := Call(__e, PrimFunc(symshen_4lazyderef), let__14116, V5006)


tmp15494 := Call(__e, PrimFunc(symshen_4deref), let__14123, V5006)


tmp15495 := Call(__e, PrimFunc(symshen_4deref), let__14118, V5006)


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
return PrimCons(Z5134, Nil)
}
__typedArg0 := Z5134
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
return PrimCons(tmp15501, V5005)
}
__typedArg0 := tmp15501
__typedArg1 := V5005
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14122, Z5135, tmp15502, V5006, V5007, let__14006, V5009)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14122, tmp15496, V5006, V5007, let__14006, tmp15497)
return


}, 0)

tmp15503 := Call(__e, PrimFunc(symbind), let__14123, tmp15491, V5006, V5007, let__14006, tmp15492)


tmp15504 := Call(__e, PrimFunc(symshen_4gc), V5006, tmp15503)


__e.TailApply(PrimFunc(symshen_4gc), V5006, tmp15504)
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

tmp15508 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15507, V5006)


let__14125 := tmp15508
_ = let__14125

tmp15509 := MakeNative(func(__e *ControlFlow) {
Z5141 := __e.Get(1)
_ = Z5141
tmp15510 := Call(__e, let__14121, let__14124)


__e.TailApply(tmp15510, Z5141)
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

tmp15513 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15512, V5006)


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

tmp15516 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15515, V5006)


let__14129 := tmp15516
_ = let__14129

tmp15517 := MakeNative(func(__e *ControlFlow) {
Z5146 := __e.Get(1)
_ = Z5146
__e.TailApply(let__14126, Z5146)
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

tmp15520 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15519, V5006)


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
__e.TailApply(PrimFunc(symshen_4bind_b), let__14132, Nil, V5006, let__14133)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15531 := Call(__e, PrimFunc(symshen_4pvar_2), let__14129)


if True == tmp15531 {
tmp15526 := Call(__e, PrimFunc(symshen_4newpv), V5006)


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

tmp15529 := Call(__e, PrimFunc(symshen_4bind_b), let__14129, tmp15527, V5006, tmp15528)


__e.TailApply(PrimFunc(symshen_4gc), V5006, tmp15529)
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
tmp15537 := Call(__e, PrimFunc(symshen_4bind_b), let__14127, sym_1_1_6, V5006, let__14128)


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
tmp15541 := Call(__e, PrimFunc(symshen_4newpv), V5006)


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

tmp15545 := Call(__e, PrimFunc(symshen_4bind_b), let__14125, tmp15543, V5006, tmp15544)


tmp15546 := Call(__e, PrimFunc(symshen_4gc), V5006, tmp15545)


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
tmp15550 := Call(__e, PrimFunc(symshen_4newpv), V5006)


let__14136 := tmp15550
_ = let__14136

tmp15551 := Call(__e, PrimFunc(symshen_4newpv), V5006)


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

tmp15557 := Call(__e, PrimFunc(symshen_4bind_b), let__14120, tmp15554, V5006, tmp15555)


tmp15558 := Call(__e, PrimFunc(symshen_4gc), V5006, tmp15557)


tmp15559 := Call(__e, PrimFunc(symshen_4gc), V5006, tmp15558)


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
tmp15617 := Call(__e, PrimFunc(symshen_4unlocked_2), V5007)


var ifres15568 Obj

if True == tmp15617 {
tmp15569 := Call(__e, PrimFunc(symshen_4lazyderef), V5003, V5006)


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

tmp15572 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15571, V5006)


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

tmp15575 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15574, V5006)


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

tmp15579 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15578, V5006)


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

tmp15583 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15582, V5006)


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

tmp15587 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15586, V5006)


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
tmp15589 := Call(__e, PrimFunc(symshen_4newpv), V5006)


let__14148 := tmp15589
_ = let__14148

tmp15590 := Call(__e, PrimFunc(symshen_4newpv), V5006)


let__14149 := tmp15590
_ = let__14149

tmp15591 := Call(__e, PrimFunc(symshen_4newpv), V5006)


let__14150 := tmp15591
_ = let__14150

tmp15592 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15592

tmp15593 := MakeNative(func(__e *ControlFlow) {
tmp15594 := Call(__e, PrimFunc(symshen_4lazyderef), let__14142, V5006)


tmp15595 := Call(__e, PrimFunc(symshen_4freshterm), tmp15594)


tmp15596 := MakeNative(func(__e *ControlFlow) {
tmp15597 := Call(__e, PrimFunc(symshen_4lazyderef), let__14142, V5006)


tmp15598 := Call(__e, PrimFunc(symshen_4lazyderef), let__14149, V5006)


tmp15599 := Call(__e, PrimFunc(symshen_4lazyderef), let__14146, V5006)


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
return PrimCons(tmp15605, V5005)
}
__typedArg0 := tmp15605
__typedArg1 := V5005
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14148, V5004, tmp15606, V5006, V5007, let__14006, V5009)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14148, tmp15600, V5006, V5007, let__14006, tmp15601)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14149, tmp15595, V5006, V5007, let__14006, tmp15596)
return


}, 0)

tmp15607 := Call(__e, PrimFunc(symshen_4system_1S_1h), let__14144, let__14150, V5005, V5006, V5007, let__14006, tmp15593)


tmp15608 := Call(__e, PrimFunc(symshen_4gc), V5006, tmp15607)


tmp15609 := Call(__e, PrimFunc(symshen_4gc), V5006, tmp15608)


tmp15610 := Call(__e, PrimFunc(symshen_4gc), V5006, tmp15609)


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
tmp15687 := Call(__e, PrimFunc(symshen_4unlocked_2), V5007)


var ifres15618 Obj

if True == tmp15687 {
tmp15619 := Call(__e, PrimFunc(symshen_4lazyderef), V5003, V5006)


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

tmp15622 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15621, V5006)


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

tmp15625 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15624, V5006)


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

tmp15629 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15628, V5006)


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

tmp15633 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15632, V5006)


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
tmp15635 := Call(__e, PrimFunc(symshen_4lazyderef), V5004, V5006)


let__14159 := tmp15635
_ = let__14159

tmp15636 := MakeNative(func(__e *ControlFlow) {
Z5177 := __e.Get(1)
_ = Z5177
tmp15637 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15637

tmp15638 := MakeNative(func(__e *ControlFlow) {
tmp15639 := Call(__e, PrimFunc(symshen_4lazyderef), Z5177, V5006)


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
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14155, symstring, V5005, V5006, V5007, let__14006, V5009)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15642, V5006, V5007, let__14006, tmp15643)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), let__14157, Z5177, V5006, V5007, let__14006, tmp15638)
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

tmp15646 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15645, V5006)


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

tmp15649 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15648, V5006)


let__14163 := tmp15649
_ = let__14163

tmp15650 := MakeNative(func(__e *ControlFlow) {
Z5182 := __e.Get(1)
_ = Z5182
__e.TailApply(let__14160, Z5182)
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

tmp15653 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15652, V5006)


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
__e.TailApply(PrimFunc(symshen_4bind_b), let__14166, Nil, V5006, let__14167)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15664 := Call(__e, PrimFunc(symshen_4pvar_2), let__14163)


if True == tmp15664 {
tmp15659 := Call(__e, PrimFunc(symshen_4newpv), V5006)


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

tmp15662 := Call(__e, PrimFunc(symshen_4bind_b), let__14163, tmp15660, V5006, tmp15661)


__e.TailApply(PrimFunc(symshen_4gc), V5006, tmp15662)
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
tmp15670 := Call(__e, PrimFunc(symshen_4bind_b), let__14161, symstream, V5006, let__14162)


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
tmp15674 := Call(__e, PrimFunc(symshen_4newpv), V5006)


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

tmp15678 := Call(__e, PrimFunc(symshen_4bind_b), let__14159, tmp15676, V5006, tmp15677)


tmp15679 := Call(__e, PrimFunc(symshen_4gc), V5006, tmp15678)


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
tmp15716 := Call(__e, PrimFunc(symshen_4unlocked_2), V5007)


var ifres15688 Obj

if True == tmp15716 {
tmp15689 := Call(__e, PrimFunc(symshen_4lazyderef), V5003, V5006)


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

tmp15692 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15691, V5006)


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

tmp15695 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15694, V5006)


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

tmp15699 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15698, V5006)


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

tmp15703 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15702, V5006)


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
tmp15707 := Call(__e, PrimFunc(symshen_4deref), let__14176, V5006)


tmp15708 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp15707)


tmp15709 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14174, V5004, V5005, V5006, V5007, let__14006, V5009)
return
}, 0)

__e.TailApply(PrimFunc(symis_b), tmp15708, V5004, V5006, V5007, let__14006, tmp15709)
return


}, 0)

tmp15710 := Call(__e, PrimFunc(symshen_4cut), V5006, V5007, let__14006, tmp15706)


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
tmp15747 := Call(__e, PrimFunc(symshen_4unlocked_2), V5007)


var ifres15717 Obj

if True == tmp15747 {
tmp15718 := Call(__e, PrimFunc(symshen_4lazyderef), V5003, V5006)


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

tmp15721 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15720, V5006)


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

tmp15724 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15723, V5006)


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

tmp15728 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15727, V5006)


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

tmp15732 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15731, V5006)


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

tmp15735 := Call(__e, PrimFunc(symshen_4deref), let__14182, V5006)


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

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14184, tmp15740, V5005, V5006, V5007, let__14006, V5009)
return


}, 0)

tmp15741 := Call(__e, PrimFunc(symis), V5004, tmp15737, V5006, V5007, let__14006, tmp15738)


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
tmp15776 := Call(__e, PrimFunc(symshen_4unlocked_2), V5007)


var ifres15748 Obj

if True == tmp15776 {
tmp15749 := Call(__e, PrimFunc(symshen_4lazyderef), V5003, V5006)


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

tmp15752 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15751, V5006)


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

tmp15755 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15754, V5006)


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

tmp15759 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15758, V5006)


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

tmp15763 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15762, V5006)


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
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14192, V5004, V5005, V5006, V5007, let__14006, V5009)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp15768, V5004, V5005, V5006, V5007, let__14006, tmp15769)
return


}, 0)

tmp15770 := Call(__e, PrimFunc(symshen_4system_1S_1h), let__14190, symsymbol, V5005, V5006, V5007, let__14006, tmp15766)


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
tmp15784 := Call(__e, PrimFunc(symshen_4unlocked_2), V5007)


var ifres15777 Obj

if True == tmp15784 {
tmp15778 := Call(__e, PrimFunc(symshen_4newpv), V5006)


let__14195 := tmp15778
_ = let__14195

tmp15779 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15779

tmp15780 := MakeNative(func(__e *ControlFlow) {
tmp15781 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), V5003, V5004, let__14195, V5006, V5007, let__14006, V5009)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5006, V5007, let__14006, tmp15781)
return


}, 0)

tmp15782 := Call(__e, PrimFunc(symshen_4l_1rules), V5005, let__14195, False, V5006, V5007, let__14006, tmp15780)


tmp15783 := Call(__e, PrimFunc(symshen_4gc), V5006, tmp15782)


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
tmp15793 := Call(__e, PrimFunc(symshen_4unlocked_2), V5007)


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
return PrimCons(V5004, Nil)
}
__typedArg0 := V5004
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
return PrimCons(V5003, tmp15789)
}
__typedArg0 := V5003
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

tmp15792 := Call(__e, PrimFunc(symshen_4search_1user_1datatypes), tmp15790, V5005, tmp15791, V5006, V5007, let__14006, V5009)


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
__e.TailApply(PrimFunc(symshen_4unlock), V5007, let__14006)
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
V5215 := __e.Get(1)
_ = V5215
tmp15869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5215)
}
__typedArg0 := V5215
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15850 Obj

if True == tmp15869 {
tmp15867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5215)
}
__typedArg0 := V5215
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
return PrimTail(V5215)
}
__typedArg0 := V5215
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
return PrimTail(V5215)
}
__typedArg0 := V5215
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
return PrimTail(V5215)
}
__typedArg0 := V5215
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
return PrimTail(V5215)
}
__typedArg0 := V5215
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
return PrimTail(V5215)
}
__typedArg0 := V5215
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
return PrimIsPair(V5215)
}
__typedArg0 := V5215
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15848 {
tmp15843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5215)
}
__typedArg0 := V5215
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15844 := Call(__e, PrimFunc(symshen_4rdecons), tmp15843)


tmp15845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5215)
}
__typedArg0 := V5215
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
__e.Return(V5215)
return
}


}


}, 1)

tmp15870 := Call(__e, ns2_1set, symshen_4rdecons, tmp15835)


_ = tmp15870

tmp15871 := MakeNative(func(__e *ControlFlow) {
V5216 := __e.Get(1)
_ = V5216
V5217 := __e.Get(2)
_ = V5217
V5218 := __e.Get(3)
_ = V5218
V5219 := __e.Get(4)
_ = V5219
V5220 := __e.Get(5)
_ = V5220
V5221 := __e.Get(6)
_ = V5221
tmp15884 := Call(__e, PrimFunc(symshen_4unlocked_2), V5219)


var ifres15872 Obj

if True == tmp15884 {
tmp15873 := Call(__e, PrimFunc(symshen_4lazyderef), V5217, V5218)


let__14198 := tmp15873
_ = let__14198

tmp15874 := MakeNative(func(__e *ControlFlow) {
tmp15875 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15875

tmp15876 := Call(__e, PrimFunc(symshen_4lazyderef), V5216, V5218)


tmp15877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(tmp15876)
}
__typedArg0 := tmp15876
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

__e.TailApply(PrimFunc(symwhen), tmp15877, V5218, V5219, V5220, V5221)
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
tmp15881 := Call(__e, PrimFunc(symshen_4bind_b), let__14198, symnumber, V5218, let__14199)


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
tmp15897 := Call(__e, PrimFunc(symshen_4unlocked_2), V5219)


var ifres15885 Obj

if True == tmp15897 {
tmp15886 := Call(__e, PrimFunc(symshen_4lazyderef), V5217, V5218)


let__14201 := tmp15886
_ = let__14201

tmp15887 := MakeNative(func(__e *ControlFlow) {
tmp15888 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15888

tmp15889 := Call(__e, PrimFunc(symshen_4lazyderef), V5216, V5218)


tmp15890 := Call(__e, PrimFunc(symboolean_2), tmp15889)


__e.TailApply(PrimFunc(symwhen), tmp15890, V5218, V5219, V5220, V5221)
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
tmp15894 := Call(__e, PrimFunc(symshen_4bind_b), let__14201, symboolean, V5218, let__14202)


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
tmp15910 := Call(__e, PrimFunc(symshen_4unlocked_2), V5219)


var ifres15898 Obj

if True == tmp15910 {
tmp15899 := Call(__e, PrimFunc(symshen_4lazyderef), V5217, V5218)


let__14204 := tmp15899
_ = let__14204

tmp15900 := MakeNative(func(__e *ControlFlow) {
tmp15901 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15901

tmp15902 := Call(__e, PrimFunc(symshen_4lazyderef), V5216, V5218)


tmp15903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(tmp15902)
}
__typedArg0 := tmp15902
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

__e.TailApply(PrimFunc(symwhen), tmp15903, V5218, V5219, V5220, V5221)
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
tmp15907 := Call(__e, PrimFunc(symshen_4bind_b), let__14204, symstring, V5218, let__14205)


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
tmp15923 := Call(__e, PrimFunc(symshen_4unlocked_2), V5219)


var ifres15911 Obj

if True == tmp15923 {
tmp15912 := Call(__e, PrimFunc(symshen_4lazyderef), V5217, V5218)


let__14207 := tmp15912
_ = let__14207

tmp15913 := MakeNative(func(__e *ControlFlow) {
tmp15914 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15914

tmp15915 := Call(__e, PrimFunc(symshen_4lazyderef), V5216, V5218)


tmp15916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(tmp15915)
}
__typedArg0 := tmp15915
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

__e.TailApply(PrimFunc(symwhen), tmp15916, V5218, V5219, V5220, V5221)
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
tmp15920 := Call(__e, PrimFunc(symshen_4bind_b), let__14207, symsymbol, V5218, let__14208)


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
tmp15966 := Call(__e, PrimFunc(symshen_4unlocked_2), V5219)


if True == tmp15966 {
tmp15924 := Call(__e, PrimFunc(symshen_4lazyderef), V5216, V5218)


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
tmp15925 := Call(__e, PrimFunc(symshen_4lazyderef), V5217, V5218)


let__14210 := tmp15925
_ = let__14210

tmp15926 := MakeNative(func(__e *ControlFlow) {
Z5237 := __e.Get(1)
_ = Z5237
tmp15927 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15927

__e.TailApply(PrimFunc(symthaw), V5221)
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

tmp15929 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15928, V5218)


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

tmp15932 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15931, V5218)


let__14214 := tmp15932
_ = let__14214

tmp15933 := MakeNative(func(__e *ControlFlow) {
Z5242 := __e.Get(1)
_ = Z5242
__e.TailApply(let__14211, Z5242)
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

tmp15936 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15935, V5218)


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
__e.TailApply(PrimFunc(symshen_4bind_b), let__14217, Nil, V5218, let__14218)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15947 := Call(__e, PrimFunc(symshen_4pvar_2), let__14214)


if True == tmp15947 {
tmp15942 := Call(__e, PrimFunc(symshen_4newpv), V5218)


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

tmp15945 := Call(__e, PrimFunc(symshen_4bind_b), let__14214, tmp15943, V5218, tmp15944)


__e.TailApply(PrimFunc(symshen_4gc), V5218, tmp15945)
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
__e.TailApply(PrimFunc(symshen_4bind_b), let__14212, symlist, V5218, let__14213)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15960 := Call(__e, PrimFunc(symshen_4pvar_2), let__14210)


if True == tmp15960 {
tmp15954 := Call(__e, PrimFunc(symshen_4newpv), V5218)


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

tmp15958 := Call(__e, PrimFunc(symshen_4bind_b), let__14210, tmp15956, V5218, tmp15957)


__e.TailApply(PrimFunc(symshen_4gc), V5218, tmp15958)
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
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
__self6 := __e.Get(6)
__self7 := __e.Get(7)
__selftop:
V5248 := __self1
_ = V5248
V5249 := __self2
_ = V5249
V5250 := __self3
_ = V5250
V5251 := __self4
_ = V5251
V5252 := __self5
_ = V5252
V5253 := __self6
_ = V5253
V5254 := __self7
_ = V5254
tmp16010 := Call(__e, PrimFunc(symshen_4unlocked_2), V5252)


var ifres15977 Obj

if True == tmp16010 {
tmp15978 := Call(__e, PrimFunc(symshen_4lazyderef), V5250, V5251)


let__14222 := tmp15978
_ = let__14222

tmp16009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14222)
}
__typedArg0 := let__14222
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15979 Obj

if True == tmp16009 {
tmp15980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14222)
}
__typedArg0 := let__14222
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15981 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15980, V5251)


let__14223 := tmp15981
_ = let__14223

tmp16008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14223)
}
__typedArg0 := let__14223
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15982 Obj

if True == tmp16008 {
tmp15983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14223)
}
__typedArg0 := let__14223
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14224 := tmp15983
_ = let__14224

tmp15984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14223)
}
__typedArg0 := let__14223
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15985 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15984, V5251)


let__14225 := tmp15985
_ = let__14225

tmp16007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14225)
}
__typedArg0 := let__14225
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15986 Obj

if True == tmp16007 {
tmp15987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14225)
}
__typedArg0 := let__14225
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14226 := tmp15987
_ = let__14226

tmp15988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14225)
}
__typedArg0 := let__14225
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15989 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15988, V5251)


let__14227 := tmp15989
_ = let__14227

tmp16006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14227)
}
__typedArg0 := let__14227
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15990 Obj

if True == tmp16006 {
tmp15991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14227)
}
__typedArg0 := let__14227
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14228 := tmp15991
_ = let__14228

tmp15992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14227)
}
__typedArg0 := let__14227
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15993 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15992, V5251)


let__14229 := tmp15993
_ = let__14229

tmp16005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14229, Nil)
}
__typedArg0 := let__14229
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15994 Obj

if True == tmp16005 {
tmp15995 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15995

tmp15996 := Call(__e, PrimFunc(symshen_4deref), let__14226, V5251)


tmp15997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp15998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp15996, tmp15997)
}
__typedArg0 := tmp15996
__typedArg1 := tmp15997
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp15999 := MakeNative(func(__e *ControlFlow) {
tmp16000 := Call(__e, PrimFunc(symshen_4deref), V5248, V5251)


tmp16001 := Call(__e, PrimFunc(symshen_4deref), let__14224, V5251)


tmp16002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16000, tmp16001)
}
__typedArg0 := tmp16000
__typedArg1 := tmp16001
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16003 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symis_b), V5249, let__14228, V5251, V5252, V5253, V5254)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp16002, V5251, V5252, V5253, tmp16003)
return


}, 0)

tmp16004 := Call(__e, PrimFunc(symwhen), tmp15998, V5251, V5252, V5253, tmp15999)


ifres15994 = tmp16004


} else {
ifres15994 = False


}

ifres15990 = ifres15994


} else {
ifres15990 = False


}

ifres15986 = ifres15990


} else {
ifres15986 = False


}

ifres15982 = ifres15986


} else {
ifres15982 = False


}

ifres15979 = ifres15982


} else {
ifres15979 = False


}

ifres15977 = ifres15979


} else {
ifres15977 = False


}

let__14221 := ifres15977
_ = let__14221

tmp16019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14221, False)
}
__typedArg0 := let__14221
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16019 {
tmp16017 := Call(__e, PrimFunc(symshen_4unlocked_2), V5252)


if True == tmp16017 {
tmp16011 := Call(__e, PrimFunc(symshen_4lazyderef), V5250, V5251)


let__14230 := tmp16011
_ = let__14230

tmp16015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14230)
}
__typedArg0 := let__14230
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16015 {
tmp16012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14230)
}
__typedArg0 := let__14230
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14231 := tmp16012
_ = let__14231

tmp16013 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16013

if PrimFunc(symshen_4by_1hypothesis) == __self {
__self1, __self2, __self3, __self4, __self5, __self6, __self7 = V5248, V5249, let__14231, V5251, V5252, V5253, V5254
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4by_1hypothesis), V5248, V5249, let__14231, V5251, V5252, V5253, V5254)
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

tmp16020 := Call(__e, ns2_1set, symshen_4by_1hypothesis, tmp15976)


_ = tmp16020

tmp16021 := MakeNative(func(__e *ControlFlow) {
V5266 := __e.Get(1)
_ = V5266
V5267 := __e.Get(2)
_ = V5267
V5268 := __e.Get(3)
_ = V5268
V5269 := __e.Get(4)
_ = V5269
V5270 := __e.Get(5)
_ = V5270
V5271 := __e.Get(6)
_ = V5271
tmp16026 := Call(__e, PrimFunc(symshen_4unlocked_2), V5269)


if True == tmp16026 {
tmp16022 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16022

tmp16023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp16024 := Call(__e, PrimFunc(symassoc), V5266, tmp16023)


__e.TailApply(PrimFunc(symshen_4sigf), tmp16024, V5267, V5268, V5269, V5270, V5271)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp16027 := Call(__e, ns2_1set, symshen_4lookupsig, tmp16021)


_ = tmp16027

tmp16028 := MakeNative(func(__e *ControlFlow) {
V5286 := __e.Get(1)
_ = V5286
V5287 := __e.Get(2)
_ = V5287
V5288 := __e.Get(3)
_ = V5288
V5289 := __e.Get(4)
_ = V5289
V5290 := __e.Get(5)
_ = V5290
V5291 := __e.Get(6)
_ = V5291
tmp16035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5286)
}
__typedArg0 := V5286
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16035 {
tmp16029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5286)
}
__typedArg0 := V5286
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16030 := Call(__e, tmp16029, V5287)


tmp16031 := Call(__e, tmp16030, V5288)


tmp16032 := Call(__e, tmp16031, V5289)


tmp16033 := Call(__e, tmp16032, V5290)


__e.TailApply(tmp16033, V5291)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp16036 := Call(__e, ns2_1set, symshen_4sigf, tmp16028)


_ = tmp16036

tmp16037 := MakeNative(func(__e *ControlFlow) {
V5292 := __e.Get(1)
_ = V5292
tmp16038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(MakeInteger(3))
}
__typedArg0 := MakeInteger(3)
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

let__14232 := tmp16038
_ = let__14232

tmp16039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__14232, MakeInteger(0), symshen_4print_1freshterm)
}
__typedArg0 := let__14232
__typedArg1 := MakeInteger(0)
__typedArg2 := symshen_4print_1freshterm
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__14233 := tmp16039
_ = let__14233

tmp16040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__14233, MakeInteger(1), V5292)
}
__typedArg0 := let__14233
__typedArg1 := MakeInteger(1)
__typedArg2 := V5292
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__14234 := tmp16040
_ = let__14234

tmp16041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dgensym_d)
}
__typedArg0 := symshen_4_dgensym_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp16043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dgensym_d, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp16041)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp16041
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := symshen_4_dgensym_d
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp16041)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp16041
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

tmp16044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__14234, MakeInteger(2), tmp16043)
}
__typedArg0 := let__14234
__typedArg1 := MakeInteger(2)
__typedArg2 := tmp16043
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__14235 := tmp16044
_ = let__14235

__e.Return(let__14235)
return


}, 1)

tmp16045 := Call(__e, ns2_1set, symshen_4freshterm, tmp16037)


_ = tmp16045

tmp16046 := MakeNative(func(__e *ControlFlow) {
V5297 := __e.Get(1)
_ = V5297
tmp16047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V5297, MakeInteger(1))
}
__typedArg0 := V5297
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp16048 := Call(__e, PrimFunc(symshen_4app), tmp16047, MakeString(""), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("&&"))
__typedS1, __typedOK1 := TypedString(tmp16048)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("&&")
__typedArg1 := tmp16048
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp16049 := Call(__e, ns2_1set, symshen_4print_1freshterm, tmp16046)


_ = tmp16049

tmp16050 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
__self6 := __e.Get(6)
__self7 := __e.Get(7)
__selftop:
V5298 := __self1
_ = V5298
V5299 := __self2
_ = V5299
V5300 := __self3
_ = V5300
V5301 := __self4
_ = V5301
V5302 := __self5
_ = V5302
V5303 := __self6
_ = V5303
V5304 := __self7
_ = V5304
tmp16067 := Call(__e, PrimFunc(symshen_4unlocked_2), V5302)


var ifres16051 Obj

if True == tmp16067 {
tmp16052 := Call(__e, PrimFunc(symshen_4lazyderef), V5300, V5301)


let__14237 := tmp16052
_ = let__14237

tmp16066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14237)
}
__typedArg0 := let__14237
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16053 Obj

if True == tmp16066 {
tmp16054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14237)
}
__typedArg0 := let__14237
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16055 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16054, V5301)


let__14238 := tmp16055
_ = let__14238

tmp16065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14238)
}
__typedArg0 := let__14238
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16056 Obj

if True == tmp16065 {
tmp16057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14238)
}
__typedArg0 := let__14238
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14239 := tmp16057
_ = let__14239

tmp16058 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16058

tmp16059 := Call(__e, PrimFunc(symshen_4deref), let__14239, V5301)


tmp16060 := Call(__e, PrimFunc(symshen_4deref), V5298, V5301)


tmp16061 := Call(__e, tmp16059, tmp16060)


tmp16062 := Call(__e, PrimFunc(symshen_4deref), V5299, V5301)


tmp16063 := Call(__e, tmp16061, tmp16062)


tmp16064 := Call(__e, PrimFunc(symcall), tmp16063, V5301, V5302, V5303, V5304)


ifres16056 = tmp16064


} else {
ifres16056 = False


}

ifres16053 = ifres16056


} else {
ifres16053 = False


}

ifres16051 = ifres16053


} else {
ifres16051 = False


}

let__14236 := ifres16051
_ = let__14236

tmp16076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14236, False)
}
__typedArg0 := let__14236
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16076 {
tmp16074 := Call(__e, PrimFunc(symshen_4unlocked_2), V5302)


if True == tmp16074 {
tmp16068 := Call(__e, PrimFunc(symshen_4lazyderef), V5300, V5301)


let__14240 := tmp16068
_ = let__14240

tmp16072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14240)
}
__typedArg0 := let__14240
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16072 {
tmp16069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14240)
}
__typedArg0 := let__14240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14241 := tmp16069
_ = let__14241

tmp16070 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16070

if PrimFunc(symshen_4search_1user_1datatypes) == __self {
__self1, __self2, __self3, __self4, __self5, __self6, __self7 = V5298, V5299, let__14241, V5301, V5302, V5303, V5304
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4search_1user_1datatypes), V5298, V5299, let__14241, V5301, V5302, V5303, V5304)
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

tmp16077 := Call(__e, ns2_1set, symshen_4search_1user_1datatypes, tmp16050)


_ = tmp16077

tmp16078 := MakeNative(func(__e *ControlFlow) {
V5311 := __e.Get(1)
_ = V5311
V5312 := __e.Get(2)
_ = V5312
V5313 := __e.Get(3)
_ = V5313
V5314 := __e.Get(4)
_ = V5314
V5315 := __e.Get(5)
_ = V5315
V5316 := __e.Get(6)
_ = V5316
V5317 := __e.Get(7)
_ = V5317
let__14242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5316)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5316
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__14242

tmp16090 := Call(__e, PrimFunc(symshen_4unlocked_2), V5315)


var ifres16080 Obj

if True == tmp16090 {
tmp16081 := Call(__e, PrimFunc(symshen_4lazyderef), V5311, V5314)


let__14244 := tmp16081
_ = let__14244

tmp16089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14244, Nil)
}
__typedArg0 := let__14244
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16082 Obj

if True == tmp16089 {
tmp16083 := Call(__e, PrimFunc(symshen_4lazyderef), V5313, V5314)


let__14245 := tmp16083
_ = let__14245

tmp16088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14245, True)
}
__typedArg0 := let__14245
__typedArg1 := True
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16084 Obj

if True == tmp16088 {
tmp16085 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16085

tmp16086 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symbind), V5312, Nil, V5314, V5315, let__14242, V5317)
return
}, 0)

tmp16087 := Call(__e, PrimFunc(symshen_4cut), V5314, V5315, let__14242, tmp16086)


ifres16084 = tmp16087


} else {
ifres16084 = False


}

ifres16082 = ifres16084


} else {
ifres16082 = False


}

ifres16080 = ifres16082


} else {
ifres16080 = False


}

let__14243 := ifres16080
_ = let__14243

tmp16433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14243, False)
}
__typedArg0 := let__14243
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16433 {
tmp16169 := Call(__e, PrimFunc(symshen_4unlocked_2), V5315)


var ifres16091 Obj

if True == tmp16169 {
tmp16092 := Call(__e, PrimFunc(symshen_4lazyderef), V5311, V5314)


let__14247 := tmp16092
_ = let__14247

tmp16168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14247)
}
__typedArg0 := let__14247
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16093 Obj

if True == tmp16168 {
tmp16094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14247)
}
__typedArg0 := let__14247
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16095 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16094, V5314)


let__14248 := tmp16095
_ = let__14248

tmp16167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14248)
}
__typedArg0 := let__14248
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16096 Obj

if True == tmp16167 {
tmp16097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14248)
}
__typedArg0 := let__14248
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16098 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16097, V5314)


let__14249 := tmp16098
_ = let__14249

tmp16166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14249)
}
__typedArg0 := let__14249
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16099 Obj

if True == tmp16166 {
tmp16100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14249)
}
__typedArg0 := let__14249
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16101 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16100, V5314)


let__14250 := tmp16101
_ = let__14250

tmp16165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14250, symcons)
}
__typedArg0 := let__14250
__typedArg1 := symcons
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16102 Obj

if True == tmp16165 {
tmp16103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14249)
}
__typedArg0 := let__14249
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16104 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16103, V5314)


let__14251 := tmp16104
_ = let__14251

tmp16164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14251)
}
__typedArg0 := let__14251
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16105 Obj

if True == tmp16164 {
tmp16106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14251)
}
__typedArg0 := let__14251
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14252 := tmp16106
_ = let__14252

tmp16107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14251)
}
__typedArg0 := let__14251
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16108 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16107, V5314)


let__14253 := tmp16108
_ = let__14253

tmp16163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14253)
}
__typedArg0 := let__14253
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16109 Obj

if True == tmp16163 {
tmp16110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14253)
}
__typedArg0 := let__14253
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14254 := tmp16110
_ = let__14254

tmp16111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14253)
}
__typedArg0 := let__14253
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16112 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16111, V5314)


let__14255 := tmp16112
_ = let__14255

tmp16162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14255, Nil)
}
__typedArg0 := let__14255
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16113 Obj

if True == tmp16162 {
tmp16114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14248)
}
__typedArg0 := let__14248
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16115 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16114, V5314)


let__14256 := tmp16115
_ = let__14256

tmp16161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14256)
}
__typedArg0 := let__14256
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16116 Obj

if True == tmp16161 {
tmp16117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14256)
}
__typedArg0 := let__14256
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14257 := tmp16117
_ = let__14257

tmp16118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14256)
}
__typedArg0 := let__14256
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16119 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16118, V5314)


let__14258 := tmp16119
_ = let__14258

tmp16160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14258)
}
__typedArg0 := let__14258
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16120 Obj

if True == tmp16160 {
tmp16121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14258)
}
__typedArg0 := let__14258
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16122 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16121, V5314)


let__14259 := tmp16122
_ = let__14259

tmp16159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14259)
}
__typedArg0 := let__14259
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16123 Obj

if True == tmp16159 {
tmp16124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14259)
}
__typedArg0 := let__14259
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16125 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16124, V5314)


let__14260 := tmp16125
_ = let__14260

tmp16158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14260, symlist)
}
__typedArg0 := let__14260
__typedArg1 := symlist
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16126 Obj

if True == tmp16158 {
tmp16127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14259)
}
__typedArg0 := let__14259
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16128 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16127, V5314)


let__14261 := tmp16128
_ = let__14261

tmp16157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14261)
}
__typedArg0 := let__14261
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16129 Obj

if True == tmp16157 {
tmp16130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14261)
}
__typedArg0 := let__14261
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14262 := tmp16130
_ = let__14262

tmp16131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14261)
}
__typedArg0 := let__14261
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16132 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16131, V5314)


let__14263 := tmp16132
_ = let__14263

tmp16156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14263, Nil)
}
__typedArg0 := let__14263
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16133 Obj

if True == tmp16156 {
tmp16134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14258)
}
__typedArg0 := let__14258
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16135 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16134, V5314)


let__14264 := tmp16135
_ = let__14264

tmp16155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14264, Nil)
}
__typedArg0 := let__14264
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16136 Obj

if True == tmp16155 {
tmp16137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14247)
}
__typedArg0 := let__14247
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14265 := tmp16137
_ = let__14265

tmp16138 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16138

tmp16139 := Call(__e, PrimFunc(symshen_4deref), let__14257, V5314)


tmp16140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16139, tmp16140)
}
__typedArg0 := tmp16139
__typedArg1 := tmp16140
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16142 := MakeNative(func(__e *ControlFlow) {
tmp16143 := MakeNative(func(__e *ControlFlow) {
tmp16144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14262, Nil)
}
__typedArg0 := let__14262
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14257, tmp16144)
}
__typedArg0 := let__14257
__typedArg1 := tmp16144
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14252, tmp16145)
}
__typedArg0 := let__14252
__typedArg1 := tmp16145
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14262, Nil)
}
__typedArg0 := let__14262
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp16147)
}
__typedArg0 := symlist
__typedArg1 := tmp16147
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16148, Nil)
}
__typedArg0 := tmp16148
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14257, tmp16149)
}
__typedArg0 := let__14257
__typedArg1 := tmp16149
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14254, tmp16150)
}
__typedArg0 := let__14254
__typedArg1 := tmp16150
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16151, let__14265)
}
__typedArg0 := tmp16151
__typedArg1 := let__14265
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16146, tmp16152)
}
__typedArg0 := tmp16146
__typedArg1 := tmp16152
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16153, V5312, True, V5314, V5315, let__14242, V5317)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5314, V5315, let__14242, tmp16143)
return


}, 0)

tmp16154 := Call(__e, PrimFunc(symwhen), tmp16141, V5314, V5315, let__14242, tmp16142)


ifres16136 = tmp16154


} else {
ifres16136 = False


}

ifres16133 = ifres16136


} else {
ifres16133 = False


}

ifres16129 = ifres16133


} else {
ifres16129 = False


}

ifres16126 = ifres16129


} else {
ifres16126 = False


}

ifres16123 = ifres16126


} else {
ifres16123 = False


}

ifres16120 = ifres16123


} else {
ifres16120 = False


}

ifres16116 = ifres16120


} else {
ifres16116 = False


}

ifres16113 = ifres16116


} else {
ifres16113 = False


}

ifres16109 = ifres16113


} else {
ifres16109 = False


}

ifres16105 = ifres16109


} else {
ifres16105 = False


}

ifres16102 = ifres16105


} else {
ifres16102 = False


}

ifres16099 = ifres16102


} else {
ifres16099 = False


}

ifres16096 = ifres16099


} else {
ifres16096 = False


}

ifres16093 = ifres16096


} else {
ifres16093 = False


}

ifres16091 = ifres16093


} else {
ifres16091 = False


}

let__14246 := ifres16091
_ = let__14246

tmp16431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14246, False)
}
__typedArg0 := let__14246
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16431 {
tmp16251 := Call(__e, PrimFunc(symshen_4unlocked_2), V5315)


var ifres16170 Obj

if True == tmp16251 {
tmp16171 := Call(__e, PrimFunc(symshen_4lazyderef), V5311, V5314)


let__14267 := tmp16171
_ = let__14267

tmp16250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14267)
}
__typedArg0 := let__14267
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16172 Obj

if True == tmp16250 {
tmp16173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14267)
}
__typedArg0 := let__14267
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16174 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16173, V5314)


let__14268 := tmp16174
_ = let__14268

tmp16249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14268)
}
__typedArg0 := let__14268
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16175 Obj

if True == tmp16249 {
tmp16176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14268)
}
__typedArg0 := let__14268
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16177 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16176, V5314)


let__14269 := tmp16177
_ = let__14269

tmp16248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14269)
}
__typedArg0 := let__14269
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16178 Obj

if True == tmp16248 {
tmp16179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14269)
}
__typedArg0 := let__14269
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16180 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16179, V5314)


let__14270 := tmp16180
_ = let__14270

tmp16247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14270, sym_8p)
}
__typedArg0 := let__14270
__typedArg1 := sym_8p
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16181 Obj

if True == tmp16247 {
tmp16182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14269)
}
__typedArg0 := let__14269
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16183 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16182, V5314)


let__14271 := tmp16183
_ = let__14271

tmp16246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14271)
}
__typedArg0 := let__14271
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16184 Obj

if True == tmp16246 {
tmp16185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14271)
}
__typedArg0 := let__14271
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14272 := tmp16185
_ = let__14272

tmp16186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14271)
}
__typedArg0 := let__14271
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16187 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16186, V5314)


let__14273 := tmp16187
_ = let__14273

tmp16245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14273)
}
__typedArg0 := let__14273
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16188 Obj

if True == tmp16245 {
tmp16189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14273)
}
__typedArg0 := let__14273
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14274 := tmp16189
_ = let__14274

tmp16190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14273)
}
__typedArg0 := let__14273
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16191 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16190, V5314)


let__14275 := tmp16191
_ = let__14275

tmp16244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14275, Nil)
}
__typedArg0 := let__14275
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16192 Obj

if True == tmp16244 {
tmp16193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14268)
}
__typedArg0 := let__14268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16194 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16193, V5314)


let__14276 := tmp16194
_ = let__14276

tmp16243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14276)
}
__typedArg0 := let__14276
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16195 Obj

if True == tmp16243 {
tmp16196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14276)
}
__typedArg0 := let__14276
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14277 := tmp16196
_ = let__14277

tmp16197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14276)
}
__typedArg0 := let__14276
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16198 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16197, V5314)


let__14278 := tmp16198
_ = let__14278

tmp16242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14278)
}
__typedArg0 := let__14278
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16199 Obj

if True == tmp16242 {
tmp16200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14278)
}
__typedArg0 := let__14278
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16201 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16200, V5314)


let__14279 := tmp16201
_ = let__14279

tmp16241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14279)
}
__typedArg0 := let__14279
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16202 Obj

if True == tmp16241 {
tmp16203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14279)
}
__typedArg0 := let__14279
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14280 := tmp16203
_ = let__14280

tmp16204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14279)
}
__typedArg0 := let__14279
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16205 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16204, V5314)


let__14281 := tmp16205
_ = let__14281

tmp16240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14281)
}
__typedArg0 := let__14281
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16206 Obj

if True == tmp16240 {
tmp16207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14281)
}
__typedArg0 := let__14281
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16208 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16207, V5314)


let__14282 := tmp16208
_ = let__14282

tmp16239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14282, sym_d)
}
__typedArg0 := let__14282
__typedArg1 := sym_d
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16209 Obj

if True == tmp16239 {
tmp16210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14281)
}
__typedArg0 := let__14281
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16211 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16210, V5314)


let__14283 := tmp16211
_ = let__14283

tmp16238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14283)
}
__typedArg0 := let__14283
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16212 Obj

if True == tmp16238 {
tmp16213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14283)
}
__typedArg0 := let__14283
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14284 := tmp16213
_ = let__14284

tmp16214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14283)
}
__typedArg0 := let__14283
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16215 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16214, V5314)


let__14285 := tmp16215
_ = let__14285

tmp16237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14285, Nil)
}
__typedArg0 := let__14285
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16216 Obj

if True == tmp16237 {
tmp16217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14278)
}
__typedArg0 := let__14278
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16218 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16217, V5314)


let__14286 := tmp16218
_ = let__14286

tmp16236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14286, Nil)
}
__typedArg0 := let__14286
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16219 Obj

if True == tmp16236 {
tmp16220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14267)
}
__typedArg0 := let__14267
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14287 := tmp16220
_ = let__14287

tmp16221 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16221

tmp16222 := Call(__e, PrimFunc(symshen_4deref), let__14277, V5314)


tmp16223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16222, tmp16223)
}
__typedArg0 := tmp16222
__typedArg1 := tmp16223
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16225 := MakeNative(func(__e *ControlFlow) {
tmp16226 := MakeNative(func(__e *ControlFlow) {
tmp16227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14280, Nil)
}
__typedArg0 := let__14280
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14277, tmp16227)
}
__typedArg0 := let__14277
__typedArg1 := tmp16227
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14272, tmp16228)
}
__typedArg0 := let__14272
__typedArg1 := tmp16228
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14284, Nil)
}
__typedArg0 := let__14284
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14277, tmp16230)
}
__typedArg0 := let__14277
__typedArg1 := tmp16230
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14274, tmp16231)
}
__typedArg0 := let__14274
__typedArg1 := tmp16231
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16232, let__14287)
}
__typedArg0 := tmp16232
__typedArg1 := let__14287
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16229, tmp16233)
}
__typedArg0 := tmp16229
__typedArg1 := tmp16233
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16234, V5312, True, V5314, V5315, let__14242, V5317)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5314, V5315, let__14242, tmp16226)
return


}, 0)

tmp16235 := Call(__e, PrimFunc(symwhen), tmp16224, V5314, V5315, let__14242, tmp16225)


ifres16219 = tmp16235


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

ifres16206 = ifres16209


} else {
ifres16206 = False


}

ifres16202 = ifres16206


} else {
ifres16202 = False


}

ifres16199 = ifres16202


} else {
ifres16199 = False


}

ifres16195 = ifres16199


} else {
ifres16195 = False


}

ifres16192 = ifres16195


} else {
ifres16192 = False


}

ifres16188 = ifres16192


} else {
ifres16188 = False


}

ifres16184 = ifres16188


} else {
ifres16184 = False


}

ifres16181 = ifres16184


} else {
ifres16181 = False


}

ifres16178 = ifres16181


} else {
ifres16178 = False


}

ifres16175 = ifres16178


} else {
ifres16175 = False


}

ifres16172 = ifres16175


} else {
ifres16172 = False


}

ifres16170 = ifres16172


} else {
ifres16170 = False


}

let__14266 := ifres16170
_ = let__14266

tmp16429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14266, False)
}
__typedArg0 := let__14266
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16429 {
tmp16315 := Call(__e, PrimFunc(symshen_4unlocked_2), V5315)


var ifres16252 Obj

if True == tmp16315 {
tmp16253 := Call(__e, PrimFunc(symshen_4lazyderef), V5311, V5314)


let__14289 := tmp16253
_ = let__14289

tmp16314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14289)
}
__typedArg0 := let__14289
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16254 Obj

if True == tmp16314 {
tmp16255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14289)
}
__typedArg0 := let__14289
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16256 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16255, V5314)


let__14290 := tmp16256
_ = let__14290

tmp16313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14290)
}
__typedArg0 := let__14290
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16257 Obj

if True == tmp16313 {
tmp16258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14290)
}
__typedArg0 := let__14290
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16259 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16258, V5314)


let__14291 := tmp16259
_ = let__14291

tmp16312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14291)
}
__typedArg0 := let__14291
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16260 Obj

if True == tmp16312 {
tmp16261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14291)
}
__typedArg0 := let__14291
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16262 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16261, V5314)


let__14292 := tmp16262
_ = let__14292

tmp16311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14292, sym_8s)
}
__typedArg0 := let__14292
__typedArg1 := sym_8s
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16263 Obj

if True == tmp16311 {
tmp16264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14291)
}
__typedArg0 := let__14291
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16265 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16264, V5314)


let__14293 := tmp16265
_ = let__14293

tmp16310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14293)
}
__typedArg0 := let__14293
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16266 Obj

if True == tmp16310 {
tmp16267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14293)
}
__typedArg0 := let__14293
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14294 := tmp16267
_ = let__14294

tmp16268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14293)
}
__typedArg0 := let__14293
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16269 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16268, V5314)


let__14295 := tmp16269
_ = let__14295

tmp16309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14295)
}
__typedArg0 := let__14295
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16270 Obj

if True == tmp16309 {
tmp16271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14295)
}
__typedArg0 := let__14295
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14296 := tmp16271
_ = let__14296

tmp16272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14295)
}
__typedArg0 := let__14295
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16273 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16272, V5314)


let__14297 := tmp16273
_ = let__14297

tmp16308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14297, Nil)
}
__typedArg0 := let__14297
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16274 Obj

if True == tmp16308 {
tmp16275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14290)
}
__typedArg0 := let__14290
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16276 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16275, V5314)


let__14298 := tmp16276
_ = let__14298

tmp16307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14298)
}
__typedArg0 := let__14298
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16277 Obj

if True == tmp16307 {
tmp16278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14298)
}
__typedArg0 := let__14298
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14299 := tmp16278
_ = let__14299

tmp16279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14298)
}
__typedArg0 := let__14298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16280 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16279, V5314)


let__14300 := tmp16280
_ = let__14300

tmp16306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14300)
}
__typedArg0 := let__14300
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16281 Obj

if True == tmp16306 {
tmp16282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14300)
}
__typedArg0 := let__14300
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16283 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16282, V5314)


let__14301 := tmp16283
_ = let__14301

tmp16305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14301, symstring)
}
__typedArg0 := let__14301
__typedArg1 := symstring
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16284 Obj

if True == tmp16305 {
tmp16285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14300)
}
__typedArg0 := let__14300
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16286 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16285, V5314)


let__14302 := tmp16286
_ = let__14302

tmp16304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14302, Nil)
}
__typedArg0 := let__14302
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16287 Obj

if True == tmp16304 {
tmp16288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14289)
}
__typedArg0 := let__14289
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14303 := tmp16288
_ = let__14303

tmp16289 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16289

tmp16290 := Call(__e, PrimFunc(symshen_4deref), let__14299, V5314)


tmp16291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16290, tmp16291)
}
__typedArg0 := tmp16290
__typedArg1 := tmp16291
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16293 := MakeNative(func(__e *ControlFlow) {
tmp16294 := MakeNative(func(__e *ControlFlow) {
tmp16295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14299, tmp16295)
}
__typedArg0 := let__14299
__typedArg1 := tmp16295
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14294, tmp16296)
}
__typedArg0 := let__14294
__typedArg1 := tmp16296
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14299, tmp16298)
}
__typedArg0 := let__14299
__typedArg1 := tmp16298
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14296, tmp16299)
}
__typedArg0 := let__14296
__typedArg1 := tmp16299
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16300, let__14303)
}
__typedArg0 := tmp16300
__typedArg1 := let__14303
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16297, tmp16301)
}
__typedArg0 := tmp16297
__typedArg1 := tmp16301
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16302, V5312, True, V5314, V5315, let__14242, V5317)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5314, V5315, let__14242, tmp16294)
return


}, 0)

tmp16303 := Call(__e, PrimFunc(symwhen), tmp16292, V5314, V5315, let__14242, tmp16293)


ifres16287 = tmp16303


} else {
ifres16287 = False


}

ifres16284 = ifres16287


} else {
ifres16284 = False


}

ifres16281 = ifres16284


} else {
ifres16281 = False


}

ifres16277 = ifres16281


} else {
ifres16277 = False


}

ifres16274 = ifres16277


} else {
ifres16274 = False


}

ifres16270 = ifres16274


} else {
ifres16270 = False


}

ifres16266 = ifres16270


} else {
ifres16266 = False


}

ifres16263 = ifres16266


} else {
ifres16263 = False


}

ifres16260 = ifres16263


} else {
ifres16260 = False


}

ifres16257 = ifres16260


} else {
ifres16257 = False


}

ifres16254 = ifres16257


} else {
ifres16254 = False


}

ifres16252 = ifres16254


} else {
ifres16252 = False


}

let__14288 := ifres16252
_ = let__14288

tmp16427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14288, False)
}
__typedArg0 := let__14288
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16427 {
tmp16394 := Call(__e, PrimFunc(symshen_4unlocked_2), V5315)


var ifres16316 Obj

if True == tmp16394 {
tmp16317 := Call(__e, PrimFunc(symshen_4lazyderef), V5311, V5314)


let__14305 := tmp16317
_ = let__14305

tmp16393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14305)
}
__typedArg0 := let__14305
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16318 Obj

if True == tmp16393 {
tmp16319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14305)
}
__typedArg0 := let__14305
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16320 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16319, V5314)


let__14306 := tmp16320
_ = let__14306

tmp16392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14306)
}
__typedArg0 := let__14306
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16321 Obj

if True == tmp16392 {
tmp16322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14306)
}
__typedArg0 := let__14306
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16323 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16322, V5314)


let__14307 := tmp16323
_ = let__14307

tmp16391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14307)
}
__typedArg0 := let__14307
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16324 Obj

if True == tmp16391 {
tmp16325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14307)
}
__typedArg0 := let__14307
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16326 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16325, V5314)


let__14308 := tmp16326
_ = let__14308

tmp16390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14308, sym_8v)
}
__typedArg0 := let__14308
__typedArg1 := sym_8v
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16327 Obj

if True == tmp16390 {
tmp16328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14307)
}
__typedArg0 := let__14307
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16329 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16328, V5314)


let__14309 := tmp16329
_ = let__14309

tmp16389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14309)
}
__typedArg0 := let__14309
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16330 Obj

if True == tmp16389 {
tmp16331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14309)
}
__typedArg0 := let__14309
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14310 := tmp16331
_ = let__14310

tmp16332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14309)
}
__typedArg0 := let__14309
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16333 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16332, V5314)


let__14311 := tmp16333
_ = let__14311

tmp16388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14311)
}
__typedArg0 := let__14311
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16334 Obj

if True == tmp16388 {
tmp16335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14311)
}
__typedArg0 := let__14311
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14312 := tmp16335
_ = let__14312

tmp16336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14311)
}
__typedArg0 := let__14311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16337 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16336, V5314)


let__14313 := tmp16337
_ = let__14313

tmp16387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14313, Nil)
}
__typedArg0 := let__14313
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16338 Obj

if True == tmp16387 {
tmp16339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14306)
}
__typedArg0 := let__14306
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16340 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16339, V5314)


let__14314 := tmp16340
_ = let__14314

tmp16386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14314)
}
__typedArg0 := let__14314
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16341 Obj

if True == tmp16386 {
tmp16342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14314)
}
__typedArg0 := let__14314
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14315 := tmp16342
_ = let__14315

tmp16343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14314)
}
__typedArg0 := let__14314
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16344 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16343, V5314)


let__14316 := tmp16344
_ = let__14316

tmp16385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14316)
}
__typedArg0 := let__14316
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16345 Obj

if True == tmp16385 {
tmp16346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14316)
}
__typedArg0 := let__14316
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16347 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16346, V5314)


let__14317 := tmp16347
_ = let__14317

tmp16384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14317)
}
__typedArg0 := let__14317
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16348 Obj

if True == tmp16384 {
tmp16349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14317)
}
__typedArg0 := let__14317
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16350 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16349, V5314)


let__14318 := tmp16350
_ = let__14318

tmp16383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14318, symvector)
}
__typedArg0 := let__14318
__typedArg1 := symvector
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16351 Obj

if True == tmp16383 {
tmp16352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14317)
}
__typedArg0 := let__14317
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16353 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16352, V5314)


let__14319 := tmp16353
_ = let__14319

tmp16382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14319)
}
__typedArg0 := let__14319
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16354 Obj

if True == tmp16382 {
tmp16355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14319)
}
__typedArg0 := let__14319
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14320 := tmp16355
_ = let__14320

tmp16356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14319)
}
__typedArg0 := let__14319
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16357 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16356, V5314)


let__14321 := tmp16357
_ = let__14321

tmp16381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14321, Nil)
}
__typedArg0 := let__14321
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16358 Obj

if True == tmp16381 {
tmp16359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14316)
}
__typedArg0 := let__14316
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16360 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16359, V5314)


let__14322 := tmp16360
_ = let__14322

tmp16380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14322, Nil)
}
__typedArg0 := let__14322
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16361 Obj

if True == tmp16380 {
tmp16362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14305)
}
__typedArg0 := let__14305
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14323 := tmp16362
_ = let__14323

tmp16363 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16363

tmp16364 := Call(__e, PrimFunc(symshen_4deref), let__14315, V5314)


tmp16365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16364, tmp16365)
}
__typedArg0 := tmp16364
__typedArg1 := tmp16365
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16367 := MakeNative(func(__e *ControlFlow) {
tmp16368 := MakeNative(func(__e *ControlFlow) {
tmp16369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14320, Nil)
}
__typedArg0 := let__14320
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14315, tmp16369)
}
__typedArg0 := let__14315
__typedArg1 := tmp16369
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14310, tmp16370)
}
__typedArg0 := let__14310
__typedArg1 := tmp16370
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14320, Nil)
}
__typedArg0 := let__14320
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp16372)
}
__typedArg0 := symvector
__typedArg1 := tmp16372
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16373, Nil)
}
__typedArg0 := tmp16373
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14315, tmp16374)
}
__typedArg0 := let__14315
__typedArg1 := tmp16374
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14312, tmp16375)
}
__typedArg0 := let__14312
__typedArg1 := tmp16375
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16376, let__14323)
}
__typedArg0 := tmp16376
__typedArg1 := let__14323
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16371, tmp16377)
}
__typedArg0 := tmp16371
__typedArg1 := tmp16377
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16378, V5312, True, V5314, V5315, let__14242, V5317)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5314, V5315, let__14242, tmp16368)
return


}, 0)

tmp16379 := Call(__e, PrimFunc(symwhen), tmp16366, V5314, V5315, let__14242, tmp16367)


ifres16361 = tmp16379


} else {
ifres16361 = False


}

ifres16358 = ifres16361


} else {
ifres16358 = False


}

ifres16354 = ifres16358


} else {
ifres16354 = False


}

ifres16351 = ifres16354


} else {
ifres16351 = False


}

ifres16348 = ifres16351


} else {
ifres16348 = False


}

ifres16345 = ifres16348


} else {
ifres16345 = False


}

ifres16341 = ifres16345


} else {
ifres16341 = False


}

ifres16338 = ifres16341


} else {
ifres16338 = False


}

ifres16334 = ifres16338


} else {
ifres16334 = False


}

ifres16330 = ifres16334


} else {
ifres16330 = False


}

ifres16327 = ifres16330


} else {
ifres16327 = False


}

ifres16324 = ifres16327


} else {
ifres16324 = False


}

ifres16321 = ifres16324


} else {
ifres16321 = False


}

ifres16318 = ifres16321


} else {
ifres16318 = False


}

ifres16316 = ifres16318


} else {
ifres16316 = False


}

let__14304 := ifres16316
_ = let__14304

tmp16425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14304, False)
}
__typedArg0 := let__14304
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16425 {
tmp16421 := Call(__e, PrimFunc(symshen_4unlocked_2), V5315)


var ifres16395 Obj

if True == tmp16421 {
tmp16396 := Call(__e, PrimFunc(symshen_4lazyderef), V5311, V5314)


let__14325 := tmp16396
_ = let__14325

tmp16420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14325)
}
__typedArg0 := let__14325
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16397 Obj

if True == tmp16420 {
tmp16398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14325)
}
__typedArg0 := let__14325
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14326 := tmp16398
_ = let__14326

tmp16399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14325)
}
__typedArg0 := let__14325
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14327 := tmp16399
_ = let__14327

tmp16400 := Call(__e, PrimFunc(symshen_4lazyderef), V5312, V5314)


let__14328 := tmp16400
_ = let__14328

tmp16401 := MakeNative(func(__e *ControlFlow) {
Z5406 := __e.Get(1)
_ = Z5406
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5407 := __e.Get(1)
_ = Z5407
tmp16402 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16402

tmp16403 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4l_1rules), let__14327, Z5407, V5313, V5314, V5315, let__14242, V5317)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5406, let__14326, V5314, V5315, let__14242, tmp16403)
return


}, 1))
return
}, 1)

let__14329 := tmp16401
_ = let__14329

tmp16419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14328)
}
__typedArg0 := let__14328
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16404 Obj

if True == tmp16419 {
tmp16405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14328)
}
__typedArg0 := let__14328
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14330 := tmp16405
_ = let__14330

tmp16406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14328)
}
__typedArg0 := let__14328
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14331 := tmp16406
_ = let__14331

tmp16407 := Call(__e, let__14329, let__14330)


tmp16408 := Call(__e, tmp16407, let__14331)


ifres16404 = tmp16408


} else {
tmp16418 := Call(__e, PrimFunc(symshen_4pvar_2), let__14328)


var ifres16409 Obj

if True == tmp16418 {
tmp16410 := Call(__e, PrimFunc(symshen_4newpv), V5314)


let__14332 := tmp16410
_ = let__14332

tmp16411 := Call(__e, PrimFunc(symshen_4newpv), V5314)


let__14333 := tmp16411
_ = let__14333

tmp16412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14332, let__14333)
}
__typedArg0 := let__14332
__typedArg1 := let__14333
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16413 := MakeNative(func(__e *ControlFlow) {
tmp16414 := Call(__e, let__14329, let__14332)


__e.TailApply(tmp16414, let__14333)
return


}, 0)

tmp16415 := Call(__e, PrimFunc(symshen_4bind_b), let__14328, tmp16412, V5314, tmp16413)


tmp16416 := Call(__e, PrimFunc(symshen_4gc), V5314, tmp16415)


tmp16417 := Call(__e, PrimFunc(symshen_4gc), V5314, tmp16416)


ifres16409 = tmp16417


} else {
ifres16409 = False


}

ifres16404 = ifres16409


}

ifres16397 = ifres16404


} else {
ifres16397 = False


}

ifres16395 = ifres16397


} else {
ifres16395 = False


}

let__14324 := ifres16395
_ = let__14324

tmp16423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14324, False)
}
__typedArg0 := let__14324
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16423 {
__e.TailApply(PrimFunc(symshen_4unlock), V5315, let__14242)
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

tmp16434 := Call(__e, ns2_1set, symshen_4l_1rules, tmp16078)


_ = tmp16434

tmp16435 := MakeNative(func(__e *ControlFlow) {
V5412 := __e.Get(1)
_ = V5412
V5413 := __e.Get(2)
_ = V5413
V5414 := __e.Get(3)
_ = V5414
V5415 := __e.Get(4)
_ = V5415
V5416 := __e.Get(5)
_ = V5416
V5417 := __e.Get(6)
_ = V5417
let__14334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5416)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5416
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__14334

tmp16475 := Call(__e, PrimFunc(symshen_4unlocked_2), V5415)


var ifres16437 Obj

if True == tmp16475 {
tmp16438 := Call(__e, PrimFunc(symshen_4lazyderef), V5412, V5414)


let__14336 := tmp16438
_ = let__14336

tmp16474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14336)
}
__typedArg0 := let__14336
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16439 Obj

if True == tmp16474 {
tmp16440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14336)
}
__typedArg0 := let__14336
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16441 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16440, V5414)


let__14337 := tmp16441
_ = let__14337

tmp16473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14337, symdefine)
}
__typedArg0 := let__14337
__typedArg1 := symdefine
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16442 Obj

if True == tmp16473 {
tmp16443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14336)
}
__typedArg0 := let__14336
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16444 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16443, V5414)


let__14338 := tmp16444
_ = let__14338

tmp16472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14338)
}
__typedArg0 := let__14338
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16445 Obj

if True == tmp16472 {
tmp16446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14338)
}
__typedArg0 := let__14338
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14339 := tmp16446
_ = let__14339

tmp16447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14338)
}
__typedArg0 := let__14338
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14340 := tmp16447
_ = let__14340

tmp16448 := Call(__e, PrimFunc(symshen_4newpv), V5414)


let__14341 := tmp16448
_ = let__14341

tmp16449 := Call(__e, PrimFunc(symshen_4newpv), V5414)


let__14342 := tmp16449
_ = let__14342

tmp16450 := Call(__e, PrimFunc(symshen_4newpv), V5414)


let__14343 := tmp16450
_ = let__14343

tmp16451 := Call(__e, PrimFunc(symshen_4newpv), V5414)


let__14344 := tmp16451
_ = let__14344

tmp16452 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16452

tmp16453 := MakeNative(func(__e *ControlFlow) {
tmp16454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14339, let__14340)
}
__typedArg0 := let__14339
__typedArg1 := let__14340
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16455 := Call(__e, PrimFunc(symshen_4sigxrules), tmp16454)


tmp16456 := MakeNative(func(__e *ControlFlow) {
tmp16457 := Call(__e, PrimFunc(symshen_4lazyderef), let__14341, V5414)


tmp16458 := Call(__e, PrimFunc(symfst), tmp16457)


tmp16459 := MakeNative(func(__e *ControlFlow) {
tmp16460 := Call(__e, PrimFunc(symshen_4lazyderef), let__14341, V5414)


tmp16461 := Call(__e, PrimFunc(symsnd), tmp16460)


tmp16462 := MakeNative(func(__e *ControlFlow) {
tmp16463 := Call(__e, PrimFunc(symshen_4deref), let__14344, V5414)


tmp16464 := Call(__e, PrimFunc(symshen_4freshen_1sig), tmp16463)


tmp16465 := MakeNative(func(__e *ControlFlow) {
tmp16466 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symis), let__14344, V5413, V5414, V5415, let__14334, V5417)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1rules), let__14339, let__14342, let__14343, MakeInteger(1), V5414, V5415, let__14334, tmp16466)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14343, tmp16464, V5414, V5415, let__14334, tmp16465)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14342, tmp16461, V5414, V5415, let__14334, tmp16462)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14344, tmp16458, V5414, V5415, let__14334, tmp16459)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14341, tmp16455, V5414, V5415, let__14334, tmp16456)
return


}, 0)

tmp16467 := Call(__e, PrimFunc(symshen_4cut), V5414, V5415, let__14334, tmp16453)


tmp16468 := Call(__e, PrimFunc(symshen_4gc), V5414, tmp16467)


tmp16469 := Call(__e, PrimFunc(symshen_4gc), V5414, tmp16468)


tmp16470 := Call(__e, PrimFunc(symshen_4gc), V5414, tmp16469)


tmp16471 := Call(__e, PrimFunc(symshen_4gc), V5414, tmp16470)


ifres16445 = tmp16471


} else {
ifres16445 = False


}

ifres16442 = ifres16445


} else {
ifres16442 = False


}

ifres16439 = ifres16442


} else {
ifres16439 = False


}

ifres16437 = ifres16439


} else {
ifres16437 = False


}

let__14335 := ifres16437
_ = let__14335

tmp16477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14335, False)
}
__typedArg0 := let__14335
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16477 {
__e.TailApply(PrimFunc(symshen_4unlock), V5415, let__14334)
return
} else {
__e.Return(let__14335)
return
}


}, 6)

tmp16478 := Call(__e, ns2_1set, symshen_4t_d, tmp16435)


_ = tmp16478

tmp16479 := MakeNative(func(__e *ControlFlow) {
V5429 := __e.Get(1)
_ = V5429
tmp16480 := MakeNative(func(__e *ControlFlow) {
Z5431 := __e.Get(1)
_ = Z5431
__e.TailApply(PrimFunc(symshen_4_5sig_drules_6), Z5431)
return
}, 1)

tmp16481 := Call(__e, PrimFunc(symcompile), tmp16480, V5429)


let__14345 := tmp16481
_ = let__14345

tmp16482 := Call(__e, PrimFunc(symfst), let__14345)


tmp16483 := MakeNative(func(__e *ControlFlow) {
Z5432 := __e.Get(1)
_ = Z5432
__e.TailApply(PrimFunc(symshen_4unmark_1rule_1custom_1patterns), Z5432)
return
}, 1)

tmp16484 := Call(__e, PrimFunc(symsnd), let__14345)


tmp16485 := Call(__e, PrimFunc(symmap), tmp16483, tmp16484)


__e.TailApply(PrimFunc(sym_8p), tmp16482, tmp16485)
return


}, 1)

tmp16486 := Call(__e, ns2_1set, symshen_4sigxrules, tmp16479)


_ = tmp16486

tmp16487 := MakeNative(func(__e *ControlFlow) {
V5433 := __e.Get(1)
_ = V5433
tmp16493 := Call(__e, PrimFunc(symtuple_2), V5433)


if True == tmp16493 {
tmp16488 := MakeNative(func(__e *ControlFlow) {
Z5434 := __e.Get(1)
_ = Z5434
__e.TailApply(PrimFunc(symshen_4unmark_1custom_1pattern), Z5434)
return
}, 1)

tmp16489 := Call(__e, PrimFunc(symfst), V5433)


tmp16490 := Call(__e, PrimFunc(symmap), tmp16488, tmp16489)


tmp16491 := Call(__e, PrimFunc(symsnd), V5433)


__e.TailApply(PrimFunc(sym_8p), tmp16490, tmp16491)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4unmark_1rule_1custom_1patterns)
return
}


}, 1)

tmp16494 := Call(__e, ns2_1set, symshen_4unmark_1rule_1custom_1patterns, tmp16487)


_ = tmp16494

tmp16495 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V5435 := __self1
_ = V5435
tmp16531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5435)
}
__typedArg0 := V5435
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16507 Obj

if True == tmp16531 {
tmp16529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5435)
}
__typedArg0 := V5435
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8p, tmp16529)
}
__typedArg0 := sym_8p
__typedArg1 := tmp16529
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16509 Obj

if True == tmp16530 {
tmp16527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5435)
}
__typedArg0 := V5435
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16527)
}
__typedArg0 := tmp16527
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16511 Obj

if True == tmp16528 {
tmp16524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5435)
}
__typedArg0 := V5435
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16524)
}
__typedArg0 := tmp16524
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4custom_1pattern, tmp16525)
}
__typedArg0 := symshen_4custom_1pattern
__typedArg1 := tmp16525
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16513 Obj

if True == tmp16526 {
tmp16521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5435)
}
__typedArg0 := V5435
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16521)
}
__typedArg0 := tmp16521
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16522)
}
__typedArg0 := tmp16522
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16515 Obj

if True == tmp16523 {
tmp16517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5435)
}
__typedArg0 := V5435
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16517)
}
__typedArg0 := tmp16517
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16518)
}
__typedArg0 := tmp16518
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp16519)
}
__typedArg0 := Nil
__typedArg1 := tmp16519
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16516 Obj

if True == tmp16520 {
ifres16516 = True


} else {
ifres16516 = False


}

ifres16515 = ifres16516


} else {
ifres16515 = False


}

var ifres16514 Obj

if True == ifres16515 {
ifres16514 = True


} else {
ifres16514 = False


}

ifres16513 = ifres16514


} else {
ifres16513 = False


}

var ifres16512 Obj

if True == ifres16513 {
ifres16512 = True


} else {
ifres16512 = False


}

ifres16511 = ifres16512


} else {
ifres16511 = False


}

var ifres16510 Obj

if True == ifres16511 {
ifres16510 = True


} else {
ifres16510 = False


}

ifres16509 = ifres16510


} else {
ifres16509 = False


}

var ifres16508 Obj

if True == ifres16509 {
ifres16508 = True


} else {
ifres16508 = False


}

ifres16507 = ifres16508


} else {
ifres16507 = False


}

if True == ifres16507 {
tmp16496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5435)
}
__typedArg0 := V5435
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16496)
}
__typedArg0 := tmp16496
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16497)
}
__typedArg0 := tmp16497
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16499 := Call(__e, PrimFunc(symshen_4curry), tmp16498)


if PrimFunc(symshen_4unmark_1custom_1pattern) == __self {
__self1 = tmp16499
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4unmark_1custom_1pattern), tmp16499)
return


} else {
tmp16505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5435)
}
__typedArg0 := V5435
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16505 {
tmp16500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5435)
}
__typedArg0 := V5435
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16501 := Call(__e, PrimFunc(symshen_4unmark_1custom_1pattern), tmp16500)


tmp16502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5435)
}
__typedArg0 := V5435
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16503 := Call(__e, PrimFunc(symshen_4unmark_1custom_1pattern), tmp16502)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16501, tmp16503)
}
__typedArg0 := tmp16501
__typedArg1 := tmp16503
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V5435)
return
}


}


}, 1)

tmp16532 := Call(__e, ns2_1set, symshen_4unmark_1custom_1pattern, tmp16495)


_ = tmp16532

tmp16533 := MakeNative(func(__e *ControlFlow) {
V5436 := __e.Get(1)
_ = V5436
tmp16560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5436)
}
__typedArg0 := V5436
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16534 Obj

if True == tmp16560 {
tmp16535 := Call(__e, PrimFunc(symtail), V5436)


let__14347 := tmp16535
_ = let__14347

tmp16558 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14347, sym_i)


var ifres16536 Obj

if True == tmp16558 {
tmp16537 := Call(__e, PrimFunc(symtail), let__14347)


let__14348 := tmp16537
_ = let__14348

tmp16538 := Call(__e, PrimFunc(symshen_4_5signature_6), let__14348)


let__14349 := tmp16538
_ = let__14349

tmp16556 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14349)


var ifres16539 Obj

if True == tmp16556 {
tmp16540 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16539 = tmp16540


} else {
tmp16541 := Call(__e, PrimFunc(symshen_4_5_1out), let__14349)


let__14350 := tmp16541
_ = let__14350

tmp16542 := Call(__e, PrimFunc(symshen_4in_1_6), let__14349)


let__14351 := tmp16542
_ = let__14351

tmp16555 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14351, sym_j)


var ifres16543 Obj

if True == tmp16555 {
tmp16544 := Call(__e, PrimFunc(symtail), let__14351)


let__14352 := tmp16544
_ = let__14352

tmp16545 := Call(__e, PrimFunc(symshen_4_5rules_d_6), let__14352)


let__14353 := tmp16545
_ = let__14353

tmp16553 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14353)


var ifres16546 Obj

if True == tmp16553 {
tmp16547 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16546 = tmp16547


} else {
tmp16548 := Call(__e, PrimFunc(symshen_4_5_1out), let__14353)


let__14354 := tmp16548
_ = let__14354

tmp16549 := Call(__e, PrimFunc(symshen_4in_1_6), let__14353)


let__14355 := tmp16549
_ = let__14355

tmp16550 := Call(__e, PrimFunc(symshen_4rectify_1type), let__14350)


let__14356 := tmp16550
_ = let__14356

tmp16551 := Call(__e, PrimFunc(sym_8p), let__14356, let__14354)


tmp16552 := Call(__e, PrimFunc(symshen_4comb), let__14355, tmp16551)


ifres16546 = tmp16552


}

ifres16543 = ifres16546


} else {
tmp16554 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16543 = tmp16554


}

ifres16539 = ifres16543


}

ifres16536 = ifres16539


} else {
tmp16557 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16536 = tmp16557


}

ifres16534 = ifres16536


} else {
tmp16559 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16534 = tmp16559


}

let__14346 := ifres16534
_ = let__14346

tmp16562 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14346)


if True == tmp16562 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__14346)
return
}


}, 1)

tmp16563 := Call(__e, ns2_1set, symshen_4_5sig_drules_6, tmp16533)


_ = tmp16563

tmp16564 := MakeNative(func(__e *ControlFlow) {
V5448 := __e.Get(1)
_ = V5448
tmp16565 := Call(__e, PrimFunc(symshen_4extract_1vars), V5448)


let__14357 := tmp16565
_ = let__14357

tmp16566 := MakeNative(func(__e *ControlFlow) {
Z5451 := __e.Get(1)
_ = Z5451
tmp16567 := Call(__e, PrimFunc(symconcat), sym_e, Z5451)


tmp16568 := Call(__e, PrimFunc(symshen_4freshterm), tmp16567)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z5451, tmp16568)
}
__typedArg0 := Z5451
__typedArg1 := tmp16568
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp16569 := Call(__e, PrimFunc(symmap), tmp16566, let__14357)


let__14358 := tmp16569
_ = let__14358

__e.TailApply(PrimFunc(symshen_4freshen_1type), let__14358, V5448)
return


}, 1)

tmp16570 := Call(__e, ns2_1set, symshen_4freshen_1sig, tmp16564)


_ = tmp16570

tmp16571 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V5452 := __self1
_ = V5452
V5453 := __self2
_ = V5453
tmp16585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5452)
}
__typedArg0 := Nil
__typedArg1 := V5452
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16585 {
__e.Return(V5453)
return
} else {
tmp16583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5452)
}
__typedArg0 := V5452
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16579 Obj

if True == tmp16583 {
tmp16581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5452)
}
__typedArg0 := V5452
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16581)
}
__typedArg0 := tmp16581
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16580 Obj

if True == tmp16582 {
ifres16580 = True


} else {
ifres16580 = False


}

ifres16579 = ifres16580


} else {
ifres16579 = False


}

if True == ifres16579 {
tmp16572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5452)
}
__typedArg0 := V5452
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5452)
}
__typedArg0 := V5452
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16573)
}
__typedArg0 := tmp16573
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5452)
}
__typedArg0 := V5452
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16575)
}
__typedArg0 := tmp16575
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16577 := Call(__e, PrimFunc(symsubst), tmp16574, tmp16576, V5453)


if PrimFunc(symshen_4freshen_1type) == __self {
__self1, __self2 = tmp16572, tmp16577
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4freshen_1type), tmp16572, tmp16577)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshen_1type)
return
}


}


}, 2)

tmp16586 := Call(__e, ns2_1set, symshen_4freshen_1type, tmp16571)


_ = tmp16586

tmp16587 := MakeNative(func(__e *ControlFlow) {
V5454 := __e.Get(1)
_ = V5454
tmp16588 := Call(__e, PrimFunc(symshen_4_5rule_d_6), V5454)


let__14360 := tmp16588
_ = let__14360

tmp16601 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14360)


var ifres16589 Obj

if True == tmp16601 {
tmp16590 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16589 = tmp16590


} else {
tmp16591 := Call(__e, PrimFunc(symshen_4_5_1out), let__14360)


let__14361 := tmp16591
_ = let__14361

tmp16592 := Call(__e, PrimFunc(symshen_4in_1_6), let__14360)


let__14362 := tmp16592
_ = let__14362

tmp16593 := Call(__e, PrimFunc(symshen_4_5rules_d_6), let__14362)


let__14363 := tmp16593
_ = let__14363

tmp16600 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14363)


var ifres16594 Obj

if True == tmp16600 {
tmp16595 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16594 = tmp16595


} else {
tmp16596 := Call(__e, PrimFunc(symshen_4_5_1out), let__14363)


let__14364 := tmp16596
_ = let__14364

tmp16597 := Call(__e, PrimFunc(symshen_4in_1_6), let__14363)


let__14365 := tmp16597
_ = let__14365

tmp16598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14361, let__14364)
}
__typedArg0 := let__14361
__typedArg1 := let__14364
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16599 := Call(__e, PrimFunc(symshen_4comb), let__14365, tmp16598)


ifres16594 = tmp16599


}

ifres16589 = ifres16594


}

let__14359 := ifres16589
_ = let__14359

tmp16613 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14359)


if True == tmp16613 {
tmp16602 := Call(__e, PrimFunc(symshen_4_5rule_d_6), V5454)


let__14367 := tmp16602
_ = let__14367

tmp16609 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14367)


var ifres16603 Obj

if True == tmp16609 {
tmp16604 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16603 = tmp16604


} else {
tmp16605 := Call(__e, PrimFunc(symshen_4_5_1out), let__14367)


let__14368 := tmp16605
_ = let__14368

tmp16606 := Call(__e, PrimFunc(symshen_4in_1_6), let__14367)


let__14369 := tmp16606
_ = let__14369

tmp16607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14368, Nil)
}
__typedArg0 := let__14368
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16608 := Call(__e, PrimFunc(symshen_4comb), let__14369, tmp16607)


ifres16603 = tmp16608


}

let__14366 := ifres16603
_ = let__14366

tmp16611 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14366)


if True == tmp16611 {
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

tmp16614 := Call(__e, ns2_1set, symshen_4_5rules_d_6, tmp16587)


_ = tmp16614

tmp16615 := MakeNative(func(__e *ControlFlow) {
V5466 := __e.Get(1)
_ = V5466
tmp16616 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5466)


let__14371 := tmp16616
_ = let__14371

tmp16644 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14371)


var ifres16617 Obj

if True == tmp16644 {
tmp16618 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16617 = tmp16618


} else {
tmp16619 := Call(__e, PrimFunc(symshen_4_5_1out), let__14371)


let__14372 := tmp16619
_ = let__14372

tmp16620 := Call(__e, PrimFunc(symshen_4in_1_6), let__14371)


let__14373 := tmp16620
_ = let__14373

tmp16643 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14373, sym_1_6)


var ifres16621 Obj

if True == tmp16643 {
tmp16622 := Call(__e, PrimFunc(symtail), let__14373)


let__14374 := tmp16622
_ = let__14374

tmp16641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14374)
}
__typedArg0 := let__14374
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16623 Obj

if True == tmp16641 {
tmp16624 := Call(__e, PrimFunc(symhead), let__14374)


let__14375 := tmp16624
_ = let__14375

tmp16625 := Call(__e, PrimFunc(symtail), let__14374)


let__14376 := tmp16625
_ = let__14376

tmp16639 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14376, symwhere)


var ifres16626 Obj

if True == tmp16639 {
tmp16627 := Call(__e, PrimFunc(symtail), let__14376)


let__14377 := tmp16627
_ = let__14377

tmp16637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14377)
}
__typedArg0 := let__14377
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16628 Obj

if True == tmp16637 {
tmp16629 := Call(__e, PrimFunc(symhead), let__14377)


let__14378 := tmp16629
_ = let__14378

tmp16630 := Call(__e, PrimFunc(symtail), let__14377)


let__14379 := tmp16630
_ = let__14379

tmp16631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14375, Nil)
}
__typedArg0 := let__14375
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14378, tmp16631)
}
__typedArg0 := let__14378
__typedArg1 := tmp16631
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16632)
}
__typedArg0 := symwhere
__typedArg1 := tmp16632
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16634 := Call(__e, PrimFunc(sym_8p), let__14372, tmp16633)


tmp16635 := Call(__e, PrimFunc(symshen_4comb), let__14379, tmp16634)


ifres16628 = tmp16635


} else {
tmp16636 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16628 = tmp16636


}

ifres16626 = ifres16628


} else {
tmp16638 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16626 = tmp16638


}

ifres16623 = ifres16626


} else {
tmp16640 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16623 = tmp16640


}

ifres16621 = ifres16623


} else {
tmp16642 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16621 = tmp16642


}

ifres16617 = ifres16621


}

let__14370 := ifres16617
_ = let__14370

tmp16717 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14370)


if True == tmp16717 {
tmp16645 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5466)


let__14381 := tmp16645
_ = let__14381

tmp16674 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14381)


var ifres16646 Obj

if True == tmp16674 {
tmp16647 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16646 = tmp16647


} else {
tmp16648 := Call(__e, PrimFunc(symshen_4_5_1out), let__14381)


let__14382 := tmp16648
_ = let__14382

tmp16649 := Call(__e, PrimFunc(symshen_4in_1_6), let__14381)


let__14383 := tmp16649
_ = let__14383

tmp16673 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14383, sym_5_1)


var ifres16650 Obj

if True == tmp16673 {
tmp16651 := Call(__e, PrimFunc(symtail), let__14383)


let__14384 := tmp16651
_ = let__14384

tmp16671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14384)
}
__typedArg0 := let__14384
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16652 Obj

if True == tmp16671 {
tmp16653 := Call(__e, PrimFunc(symhead), let__14384)


let__14385 := tmp16653
_ = let__14385

tmp16654 := Call(__e, PrimFunc(symtail), let__14384)


let__14386 := tmp16654
_ = let__14386

tmp16669 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14386, symwhere)


var ifres16655 Obj

if True == tmp16669 {
tmp16656 := Call(__e, PrimFunc(symtail), let__14386)


let__14387 := tmp16656
_ = let__14387

tmp16667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14387)
}
__typedArg0 := let__14387
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16657 Obj

if True == tmp16667 {
tmp16658 := Call(__e, PrimFunc(symhead), let__14387)


let__14388 := tmp16658
_ = let__14388

tmp16659 := Call(__e, PrimFunc(symtail), let__14387)


let__14389 := tmp16659
_ = let__14389

tmp16660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14385, Nil)
}
__typedArg0 := let__14385
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14388, tmp16660)
}
__typedArg0 := let__14388
__typedArg1 := tmp16660
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16661)
}
__typedArg0 := symwhere
__typedArg1 := tmp16661
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16663 := Call(__e, PrimFunc(symshen_4correct), tmp16662)


tmp16664 := Call(__e, PrimFunc(sym_8p), let__14382, tmp16663)


tmp16665 := Call(__e, PrimFunc(symshen_4comb), let__14389, tmp16664)


ifres16657 = tmp16665


} else {
tmp16666 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16657 = tmp16666


}

ifres16655 = ifres16657


} else {
tmp16668 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16655 = tmp16668


}

ifres16652 = ifres16655


} else {
tmp16670 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16652 = tmp16670


}

ifres16650 = ifres16652


} else {
tmp16672 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16650 = tmp16672


}

ifres16646 = ifres16650


}

let__14380 := ifres16646
_ = let__14380

tmp16715 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14380)


if True == tmp16715 {
tmp16675 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5466)


let__14391 := tmp16675
_ = let__14391

tmp16692 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14391)


var ifres16676 Obj

if True == tmp16692 {
tmp16677 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16676 = tmp16677


} else {
tmp16678 := Call(__e, PrimFunc(symshen_4_5_1out), let__14391)


let__14392 := tmp16678
_ = let__14392

tmp16679 := Call(__e, PrimFunc(symshen_4in_1_6), let__14391)


let__14393 := tmp16679
_ = let__14393

tmp16691 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14393, sym_5_1)


var ifres16680 Obj

if True == tmp16691 {
tmp16681 := Call(__e, PrimFunc(symtail), let__14393)


let__14394 := tmp16681
_ = let__14394

tmp16689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14394)
}
__typedArg0 := let__14394
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16682 Obj

if True == tmp16689 {
tmp16683 := Call(__e, PrimFunc(symhead), let__14394)


let__14395 := tmp16683
_ = let__14395

tmp16684 := Call(__e, PrimFunc(symtail), let__14394)


let__14396 := tmp16684
_ = let__14396

tmp16685 := Call(__e, PrimFunc(symshen_4correct), let__14395)


tmp16686 := Call(__e, PrimFunc(sym_8p), let__14392, tmp16685)


tmp16687 := Call(__e, PrimFunc(symshen_4comb), let__14396, tmp16686)


ifres16682 = tmp16687


} else {
tmp16688 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16682 = tmp16688


}

ifres16680 = ifres16682


} else {
tmp16690 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16680 = tmp16690


}

ifres16676 = ifres16680


}

let__14390 := ifres16676
_ = let__14390

tmp16713 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14390)


if True == tmp16713 {
tmp16693 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5466)


let__14398 := tmp16693
_ = let__14398

tmp16709 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14398)


var ifres16694 Obj

if True == tmp16709 {
tmp16695 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16694 = tmp16695


} else {
tmp16696 := Call(__e, PrimFunc(symshen_4_5_1out), let__14398)


let__14399 := tmp16696
_ = let__14399

tmp16697 := Call(__e, PrimFunc(symshen_4in_1_6), let__14398)


let__14400 := tmp16697
_ = let__14400

tmp16708 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14400, sym_1_6)


var ifres16698 Obj

if True == tmp16708 {
tmp16699 := Call(__e, PrimFunc(symtail), let__14400)


let__14401 := tmp16699
_ = let__14401

tmp16706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14401)
}
__typedArg0 := let__14401
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16700 Obj

if True == tmp16706 {
tmp16701 := Call(__e, PrimFunc(symhead), let__14401)


let__14402 := tmp16701
_ = let__14402

tmp16702 := Call(__e, PrimFunc(symtail), let__14401)


let__14403 := tmp16702
_ = let__14403

tmp16703 := Call(__e, PrimFunc(sym_8p), let__14399, let__14402)


tmp16704 := Call(__e, PrimFunc(symshen_4comb), let__14403, tmp16703)


ifres16700 = tmp16704


} else {
tmp16705 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16700 = tmp16705


}

ifres16698 = ifres16700


} else {
tmp16707 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16698 = tmp16707


}

ifres16694 = ifres16698


}

let__14397 := ifres16694
_ = let__14397

tmp16711 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14397)


if True == tmp16711 {
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

tmp16718 := Call(__e, ns2_1set, symshen_4_5rule_d_6, tmp16615)


_ = tmp16718

tmp16719 := MakeNative(func(__e *ControlFlow) {
V5501 := __e.Get(1)
_ = V5501
tmp16867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16811 Obj

if True == tmp16867 {
tmp16865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp16865)
}
__typedArg0 := symwhere
__typedArg1 := tmp16865
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16813 Obj

if True == tmp16866 {
tmp16863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16863)
}
__typedArg0 := tmp16863
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16815 Obj

if True == tmp16864 {
tmp16860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16860)
}
__typedArg0 := tmp16860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16861)
}
__typedArg0 := tmp16861
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16817 Obj

if True == tmp16862 {
tmp16856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5501)
}
__typedArg0 := V5501
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16858)
}
__typedArg0 := tmp16858
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16819 Obj

if True == tmp16859 {
tmp16851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16851)
}
__typedArg0 := tmp16851
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16852)
}
__typedArg0 := tmp16852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16853)
}
__typedArg0 := tmp16853
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail_1if, tmp16854)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp16854
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16821 Obj

if True == tmp16855 {
tmp16846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5501)
}
__typedArg0 := V5501
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16847)
}
__typedArg0 := tmp16847
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16848)
}
__typedArg0 := tmp16848
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16849)
}
__typedArg0 := tmp16849
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16823 Obj

if True == tmp16850 {
tmp16840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5501)
}
__typedArg0 := V5501
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16841)
}
__typedArg0 := tmp16841
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16842)
}
__typedArg0 := tmp16842
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16844)
}
__typedArg0 := tmp16844
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16825 Obj

if True == tmp16845 {
tmp16833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16833)
}
__typedArg0 := tmp16833
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16834)
}
__typedArg0 := tmp16834
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16835)
}
__typedArg0 := tmp16835
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16836)
}
__typedArg0 := tmp16836
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16837)
}
__typedArg0 := tmp16837
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp16838)
}
__typedArg0 := Nil
__typedArg1 := tmp16838
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16827 Obj

if True == tmp16839 {
tmp16829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16829)
}
__typedArg0 := tmp16829
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16830)
}
__typedArg0 := tmp16830
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp16831)
}
__typedArg0 := Nil
__typedArg1 := tmp16831
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16828 Obj

if True == tmp16832 {
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

var ifres16820 Obj

if True == ifres16821 {
ifres16820 = True


} else {
ifres16820 = False


}

ifres16819 = ifres16820


} else {
ifres16819 = False


}

var ifres16818 Obj

if True == ifres16819 {
ifres16818 = True


} else {
ifres16818 = False


}

ifres16817 = ifres16818


} else {
ifres16817 = False


}

var ifres16816 Obj

if True == ifres16817 {
ifres16816 = True


} else {
ifres16816 = False


}

ifres16815 = ifres16816


} else {
ifres16815 = False


}

var ifres16814 Obj

if True == ifres16815 {
ifres16814 = True


} else {
ifres16814 = False


}

ifres16813 = ifres16814


} else {
ifres16813 = False


}

var ifres16812 Obj

if True == ifres16813 {
ifres16812 = True


} else {
ifres16812 = False


}

ifres16811 = ifres16812


} else {
ifres16811 = False


}

if True == ifres16811 {
tmp16720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16720)
}
__typedArg0 := tmp16720
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16722)
}
__typedArg0 := tmp16722
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16723)
}
__typedArg0 := tmp16723
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16724)
}
__typedArg0 := tmp16724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16725, Nil)
}
__typedArg0 := tmp16725
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp16726)
}
__typedArg0 := symnot
__typedArg1 := tmp16726
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16727, Nil)
}
__typedArg0 := tmp16727
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16721, tmp16728)
}
__typedArg0 := tmp16721
__typedArg1 := tmp16728
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp16729)
}
__typedArg0 := symand
__typedArg1 := tmp16729
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16731)
}
__typedArg0 := tmp16731
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16732)
}
__typedArg0 := tmp16732
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16733)
}
__typedArg0 := tmp16733
return Call(__e, PrimFunc(symtl), __typedArg0)
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
return PrimCons(tmp16730, tmp16735)
}
__typedArg0 := tmp16730
__typedArg1 := tmp16735
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16736)
}
__typedArg0 := symwhere
__typedArg1 := tmp16736
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp16809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16790 Obj

if True == tmp16809 {
tmp16807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp16807)
}
__typedArg0 := symwhere
__typedArg1 := tmp16807
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16792 Obj

if True == tmp16808 {
tmp16805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16805)
}
__typedArg0 := tmp16805
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16794 Obj

if True == tmp16806 {
tmp16802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16802)
}
__typedArg0 := tmp16802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16803)
}
__typedArg0 := tmp16803
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16796 Obj

if True == tmp16804 {
tmp16798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16798)
}
__typedArg0 := tmp16798
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16799)
}
__typedArg0 := tmp16799
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp16800)
}
__typedArg0 := Nil
__typedArg1 := tmp16800
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16797 Obj

if True == tmp16801 {
ifres16797 = True


} else {
ifres16797 = False


}

ifres16796 = ifres16797


} else {
ifres16796 = False


}

var ifres16795 Obj

if True == ifres16796 {
ifres16795 = True


} else {
ifres16795 = False


}

ifres16794 = ifres16795


} else {
ifres16794 = False


}

var ifres16793 Obj

if True == ifres16794 {
ifres16793 = True


} else {
ifres16793 = False


}

ifres16792 = ifres16793


} else {
ifres16792 = False


}

var ifres16791 Obj

if True == ifres16792 {
ifres16791 = True


} else {
ifres16791 = False


}

ifres16790 = ifres16791


} else {
ifres16790 = False


}

if True == ifres16790 {
tmp16737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16737)
}
__typedArg0 := tmp16737
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16739)
}
__typedArg0 := tmp16739
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16740)
}
__typedArg0 := tmp16740
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, Nil)
}
__typedArg0 := symfail
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16742, Nil)
}
__typedArg0 := tmp16742
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16741, tmp16743)
}
__typedArg0 := tmp16741
__typedArg1 := tmp16743
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp16744)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp16744
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16745, Nil)
}
__typedArg0 := tmp16745
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp16746)
}
__typedArg0 := symnot
__typedArg1 := tmp16746
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16747, Nil)
}
__typedArg0 := tmp16747
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16738, tmp16748)
}
__typedArg0 := tmp16738
__typedArg1 := tmp16748
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp16749)
}
__typedArg0 := symand
__typedArg1 := tmp16749
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16751)
}
__typedArg0 := tmp16751
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16750, tmp16752)
}
__typedArg0 := tmp16750
__typedArg1 := tmp16752
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16753)
}
__typedArg0 := symwhere
__typedArg1 := tmp16753
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp16788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16769 Obj

if True == tmp16788 {
tmp16786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail_1if, tmp16786)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp16786
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16771 Obj

if True == tmp16787 {
tmp16784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16784)
}
__typedArg0 := tmp16784
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16773 Obj

if True == tmp16785 {
tmp16781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16781)
}
__typedArg0 := tmp16781
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16782)
}
__typedArg0 := tmp16782
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16775 Obj

if True == tmp16783 {
tmp16777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16777)
}
__typedArg0 := tmp16777
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16778)
}
__typedArg0 := tmp16778
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp16779)
}
__typedArg0 := Nil
__typedArg1 := tmp16779
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16776 Obj

if True == tmp16780 {
ifres16776 = True


} else {
ifres16776 = False


}

ifres16775 = ifres16776


} else {
ifres16775 = False


}

var ifres16774 Obj

if True == ifres16775 {
ifres16774 = True


} else {
ifres16774 = False


}

ifres16773 = ifres16774


} else {
ifres16773 = False


}

var ifres16772 Obj

if True == ifres16773 {
ifres16772 = True


} else {
ifres16772 = False


}

ifres16771 = ifres16772


} else {
ifres16771 = False


}

var ifres16770 Obj

if True == ifres16771 {
ifres16770 = True


} else {
ifres16770 = False


}

ifres16769 = ifres16770


} else {
ifres16769 = False


}

if True == ifres16769 {
tmp16754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16754, Nil)
}
__typedArg0 := tmp16754
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp16755)
}
__typedArg0 := symnot
__typedArg1 := tmp16755
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5501)
}
__typedArg0 := V5501
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16757)
}
__typedArg0 := tmp16757
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16756, tmp16758)
}
__typedArg0 := tmp16756
__typedArg1 := tmp16758
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16759)
}
__typedArg0 := symwhere
__typedArg1 := tmp16759
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp16760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, Nil)
}
__typedArg0 := symfail
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16760, Nil)
}
__typedArg0 := tmp16760
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5501, tmp16761)
}
__typedArg0 := V5501
__typedArg1 := tmp16761
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp16762)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp16762
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16763, Nil)
}
__typedArg0 := tmp16763
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp16764)
}
__typedArg0 := symnot
__typedArg1 := tmp16764
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5501, Nil)
}
__typedArg0 := V5501
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16765, tmp16766)
}
__typedArg0 := tmp16765
__typedArg1 := tmp16766
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16767)
}
__typedArg0 := symwhere
__typedArg1 := tmp16767
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}


}


}, 1)

tmp16868 := Call(__e, ns2_1set, symshen_4correct, tmp16719)


_ = tmp16868

tmp16869 := MakeNative(func(__e *ControlFlow) {
V5502 := __e.Get(1)
_ = V5502
V5503 := __e.Get(2)
_ = V5503
V5504 := __e.Get(3)
_ = V5504
V5505 := __e.Get(4)
_ = V5505
V5506 := __e.Get(5)
_ = V5506
V5507 := __e.Get(6)
_ = V5507
V5508 := __e.Get(7)
_ = V5508
V5509 := __e.Get(8)
_ = V5509
let__14404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5508)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5508
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__14404

tmp16877 := Call(__e, PrimFunc(symshen_4unlocked_2), V5507)


var ifres16871 Obj

if True == tmp16877 {
tmp16872 := Call(__e, PrimFunc(symshen_4lazyderef), V5503, V5506)


let__14406 := tmp16872
_ = let__14406

tmp16876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14406, Nil)
}
__typedArg0 := let__14406
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16873 Obj

if True == tmp16876 {
tmp16874 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16874

tmp16875 := Call(__e, PrimFunc(symthaw), V5509)


ifres16873 = tmp16875


} else {
ifres16873 = False


}

ifres16871 = ifres16873


} else {
ifres16871 = False


}

let__14405 := ifres16871
_ = let__14405

tmp16902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14405, False)
}
__typedArg0 := let__14405
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16902 {
tmp16898 := Call(__e, PrimFunc(symshen_4unlocked_2), V5507)


var ifres16878 Obj

if True == tmp16898 {
tmp16879 := Call(__e, PrimFunc(symshen_4lazyderef), V5503, V5506)


let__14408 := tmp16879
_ = let__14408

tmp16897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14408)
}
__typedArg0 := let__14408
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16880 Obj

if True == tmp16897 {
tmp16881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14408)
}
__typedArg0 := let__14408
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14409 := tmp16881
_ = let__14409

tmp16882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14408)
}
__typedArg0 := let__14408
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14410 := tmp16882
_ = let__14410

tmp16883 := Call(__e, PrimFunc(symshen_4newpv), V5506)


let__14411 := tmp16883
_ = let__14411

tmp16884 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16884

tmp16885 := Call(__e, PrimFunc(symshen_4deref), let__14409, V5506)


tmp16886 := Call(__e, PrimFunc(symshen_4freshen_1rule), tmp16885)


tmp16887 := MakeNative(func(__e *ControlFlow) {
tmp16888 := Call(__e, PrimFunc(symshen_4lazyderef), let__14411, V5506)


tmp16889 := Call(__e, PrimFunc(symfst), tmp16888)


tmp16890 := Call(__e, PrimFunc(symshen_4lazyderef), let__14411, V5506)


tmp16891 := Call(__e, PrimFunc(symsnd), tmp16890)


tmp16892 := MakeNative(func(__e *ControlFlow) {
tmp16893 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1rules), V5502, let__14410, V5504, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5505)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5505
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})(), V5506, V5507, let__14404, V5509)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5506, V5507, let__14404, tmp16893)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1rule), V5502, V5505, tmp16889, tmp16891, V5504, V5506, V5507, let__14404, tmp16892)
return


}, 0)

tmp16895 := Call(__e, PrimFunc(symbind), let__14411, tmp16886, V5506, V5507, let__14404, tmp16887)


tmp16896 := Call(__e, PrimFunc(symshen_4gc), V5506, tmp16895)


ifres16880 = tmp16896


} else {
ifres16880 = False


}

ifres16878 = ifres16880


} else {
ifres16878 = False


}

let__14407 := ifres16878
_ = let__14407

tmp16900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14407, False)
}
__typedArg0 := let__14407
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16900 {
__e.TailApply(PrimFunc(symshen_4unlock), V5507, let__14404)
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

tmp16903 := Call(__e, ns2_1set, symshen_4t_d_1rules, tmp16869)


_ = tmp16903

tmp16904 := MakeNative(func(__e *ControlFlow) {
V5518 := __e.Get(1)
_ = V5518
tmp16915 := Call(__e, PrimFunc(symtuple_2), V5518)


if True == tmp16915 {
tmp16905 := Call(__e, PrimFunc(symfst), V5518)


tmp16906 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp16905)


let__14412 := tmp16906
_ = let__14412

tmp16907 := MakeNative(func(__e *ControlFlow) {
Z5521 := __e.Get(1)
_ = Z5521
tmp16908 := Call(__e, PrimFunc(symshen_4freshterm), Z5521)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z5521, tmp16908)
}
__typedArg0 := Z5521
__typedArg1 := tmp16908
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp16909 := Call(__e, PrimFunc(symmap), tmp16907, let__14412)


let__14413 := tmp16909
_ = let__14413

tmp16910 := Call(__e, PrimFunc(symfst), V5518)


tmp16911 := Call(__e, PrimFunc(symshen_4freshen), let__14413, tmp16910)


tmp16912 := Call(__e, PrimFunc(symsnd), V5518)


tmp16913 := Call(__e, PrimFunc(symshen_4freshen), let__14413, tmp16912)


__e.TailApply(PrimFunc(sym_8p), tmp16911, tmp16913)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshen_1rule)
return
}


}, 1)

tmp16916 := Call(__e, ns2_1set, symshen_4freshen_1rule, tmp16904)


_ = tmp16916

tmp16917 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V5522 := __self1
_ = V5522
V5523 := __self2
_ = V5523
tmp16931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5522)
}
__typedArg0 := Nil
__typedArg1 := V5522
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16931 {
__e.Return(V5523)
return
} else {
tmp16929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5522)
}
__typedArg0 := V5522
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16925 Obj

if True == tmp16929 {
tmp16927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5522)
}
__typedArg0 := V5522
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16927)
}
__typedArg0 := tmp16927
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16926 Obj

if True == tmp16928 {
ifres16926 = True


} else {
ifres16926 = False


}

ifres16925 = ifres16926


} else {
ifres16925 = False


}

if True == ifres16925 {
tmp16918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5522)
}
__typedArg0 := V5522
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5522)
}
__typedArg0 := V5522
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16919)
}
__typedArg0 := tmp16919
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5522)
}
__typedArg0 := V5522
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16921)
}
__typedArg0 := tmp16921
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16923 := Call(__e, PrimFunc(symshen_4beta), tmp16920, tmp16922, V5523)


if PrimFunc(symshen_4freshen) == __self {
__self1, __self2 = tmp16918, tmp16923
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4freshen), tmp16918, tmp16923)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshen)
return
}


}


}, 2)

tmp16932 := Call(__e, ns2_1set, symshen_4freshen, tmp16917)


_ = tmp16932

tmp16933 := MakeNative(func(__e *ControlFlow) {
V5524 := __e.Get(1)
_ = V5524
V5525 := __e.Get(2)
_ = V5525
V5526 := __e.Get(3)
_ = V5526
V5527 := __e.Get(4)
_ = V5527
V5528 := __e.Get(5)
_ = V5528
V5529 := __e.Get(6)
_ = V5529
V5530 := __e.Get(7)
_ = V5530
V5531 := __e.Get(8)
_ = V5531
V5532 := __e.Get(9)
_ = V5532
tmp16937 := Call(__e, PrimFunc(symshen_4unlocked_2), V5530)


var ifres16934 Obj

if True == tmp16937 {
tmp16935 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16935

tmp16936 := Call(__e, PrimFunc(symshen_4t_d_1rule_1h), V5526, V5527, V5528, V5529, V5530, V5531, V5532)


ifres16934 = tmp16936


} else {
ifres16934 = False


}

let__14414 := ifres16934
_ = let__14414

tmp16949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14414, False)
}
__typedArg0 := let__14414
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16949 {
tmp16947 := Call(__e, PrimFunc(symshen_4unlocked_2), V5530)


if True == tmp16947 {
tmp16938 := Call(__e, PrimFunc(symshen_4newpv), V5529)


let__14415 := tmp16938
_ = let__14415

tmp16939 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16939

tmp16940 := Call(__e, PrimFunc(symshen_4app), V5524, MakeString("\n"), symshen_4a)


tmp16942 := Call(__e, PrimFunc(symshen_4app), V5525, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" of "))
__typedS1, __typedOK1 := TypedString(tmp16940)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" of ")
__typedArg1 := tmp16940
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp16944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("type error in rule "))
__typedS1, __typedOK1 := TypedString(tmp16942)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("type error in rule ")
__typedArg1 := tmp16942
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("type error in rule "))
__typedS1, __typedOK1 := TypedString(tmp16942)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("type error in rule ")
__typedArg1 := tmp16942
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

tmp16945 := Call(__e, PrimFunc(symbind), let__14415, tmp16944, V5529, V5530, V5531, V5532)


__e.TailApply(PrimFunc(symshen_4gc), V5529, tmp16945)
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

tmp16950 := Call(__e, ns2_1set, symshen_4t_d_1rule, tmp16933)


_ = tmp16950

tmp16951 := MakeNative(func(__e *ControlFlow) {
V5535 := __e.Get(1)
_ = V5535
V5536 := __e.Get(2)
_ = V5536
V5537 := __e.Get(3)
_ = V5537
V5538 := __e.Get(4)
_ = V5538
V5539 := __e.Get(5)
_ = V5539
V5540 := __e.Get(6)
_ = V5540
V5541 := __e.Get(7)
_ = V5541
let__14416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5540)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5540
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__14416

tmp16976 := Call(__e, PrimFunc(symshen_4unlocked_2), V5539)


var ifres16953 Obj

if True == tmp16976 {
tmp16954 := Call(__e, PrimFunc(symshen_4lazyderef), V5535, V5538)


let__14418 := tmp16954
_ = let__14418

tmp16975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14418, Nil)
}
__typedArg0 := let__14418
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16955 Obj

if True == tmp16975 {
tmp16956 := Call(__e, PrimFunc(symshen_4lazyderef), V5537, V5538)


let__14419 := tmp16956
_ = let__14419

tmp16974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14419)
}
__typedArg0 := let__14419
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16957 Obj

if True == tmp16974 {
tmp16958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14419)
}
__typedArg0 := let__14419
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16959 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16958, V5538)


let__14420 := tmp16959
_ = let__14420

tmp16973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14420, sym_1_1_6)
}
__typedArg0 := let__14420
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16960 Obj

if True == tmp16973 {
tmp16961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14419)
}
__typedArg0 := let__14419
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16962 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16961, V5538)


let__14421 := tmp16962
_ = let__14421

tmp16972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14421)
}
__typedArg0 := let__14421
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16963 Obj

if True == tmp16972 {
tmp16964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14421)
}
__typedArg0 := let__14421
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14422 := tmp16964
_ = let__14422

tmp16965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14421)
}
__typedArg0 := let__14421
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16966 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16965, V5538)


let__14423 := tmp16966
_ = let__14423

tmp16971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14423, Nil)
}
__typedArg0 := let__14423
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16967 Obj

if True == tmp16971 {
tmp16968 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16968

tmp16969 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1correct), V5536, let__14422, Nil, V5538, V5539, let__14416, V5541)
return
}, 0)

tmp16970 := Call(__e, PrimFunc(symshen_4cut), V5538, V5539, let__14416, tmp16969)


ifres16967 = tmp16970


} else {
ifres16967 = False


}

ifres16963 = ifres16967


} else {
ifres16963 = False


}

ifres16960 = ifres16963


} else {
ifres16960 = False


}

ifres16957 = ifres16960


} else {
ifres16957 = False


}

ifres16955 = ifres16957


} else {
ifres16955 = False


}

ifres16953 = ifres16955


} else {
ifres16953 = False


}

let__14417 := ifres16953
_ = let__14417

tmp16995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14417, False)
}
__typedArg0 := let__14417
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16995 {
tmp16991 := Call(__e, PrimFunc(symshen_4unlocked_2), V5539)


var ifres16977 Obj

if True == tmp16991 {
tmp16978 := Call(__e, PrimFunc(symshen_4newpv), V5538)


let__14425 := tmp16978
_ = let__14425

tmp16979 := Call(__e, PrimFunc(symshen_4newpv), V5538)


let__14426 := tmp16979
_ = let__14426

tmp16980 := Call(__e, PrimFunc(symshen_4newpv), V5538)


let__14427 := tmp16980
_ = let__14427

tmp16981 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16981

tmp16982 := Call(__e, PrimFunc(symshen_4freshterms), V5535)


tmp16983 := MakeNative(func(__e *ControlFlow) {
tmp16984 := MakeNative(func(__e *ControlFlow) {
tmp16985 := MakeNative(func(__e *ControlFlow) {
tmp16986 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1correct), V5536, let__14426, let__14427, V5538, V5539, let__14416, V5541)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4myassume), V5535, V5537, let__14427, V5538, V5539, let__14416, tmp16986)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5538, V5539, let__14416, tmp16985)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1integrity), V5535, V5537, let__14425, let__14426, V5538, V5539, let__14416, tmp16984)
return


}, 0)

tmp16987 := Call(__e, PrimFunc(symshen_4p_1hyps), tmp16982, let__14425, V5538, V5539, let__14416, tmp16983)


tmp16988 := Call(__e, PrimFunc(symshen_4gc), V5538, tmp16987)


tmp16989 := Call(__e, PrimFunc(symshen_4gc), V5538, tmp16988)


tmp16990 := Call(__e, PrimFunc(symshen_4gc), V5538, tmp16989)


ifres16977 = tmp16990


} else {
ifres16977 = False


}

let__14424 := ifres16977
_ = let__14424

tmp16993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14424, False)
}
__typedArg0 := let__14424
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16993 {
__e.TailApply(PrimFunc(symshen_4unlock), V5539, let__14416)
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

tmp16996 := Call(__e, ns2_1set, symshen_4t_d_1rule_1h, tmp16951)


_ = tmp16996

tmp16997 := MakeNative(func(__e *ControlFlow) {
V5554 := __e.Get(1)
_ = V5554
V5555 := __e.Get(2)
_ = V5555
V5556 := __e.Get(3)
_ = V5556
V5557 := __e.Get(4)
_ = V5557
V5558 := __e.Get(5)
_ = V5558
V5559 := __e.Get(6)
_ = V5559
V5560 := __e.Get(7)
_ = V5560
tmp17011 := Call(__e, PrimFunc(symshen_4unlocked_2), V5558)


var ifres16998 Obj

if True == tmp17011 {
tmp16999 := Call(__e, PrimFunc(symshen_4lazyderef), V5554, V5557)


let__14429 := tmp16999
_ = let__14429

tmp17010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14429, Nil)
}
__typedArg0 := let__14429
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17000 Obj

if True == tmp17010 {
tmp17001 := Call(__e, PrimFunc(symshen_4lazyderef), V5556, V5557)


let__14430 := tmp17001
_ = let__14430

tmp17002 := MakeNative(func(__e *ControlFlow) {
tmp17003 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17003

__e.TailApply(PrimFunc(symthaw), V5560)
return


}, 0)

let__14431 := tmp17002
_ = let__14431

tmp17009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14430, Nil)
}
__typedArg0 := let__14430
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17004 Obj

if True == tmp17009 {
tmp17005 := Call(__e, PrimFunc(symthaw), let__14431)


ifres17004 = tmp17005


} else {
tmp17008 := Call(__e, PrimFunc(symshen_4pvar_2), let__14430)


var ifres17006 Obj

if True == tmp17008 {
tmp17007 := Call(__e, PrimFunc(symshen_4bind_b), let__14430, Nil, V5557, let__14431)


ifres17006 = tmp17007


} else {
ifres17006 = False


}

ifres17004 = ifres17006


}

ifres17000 = ifres17004


} else {
ifres17000 = False


}

ifres16998 = ifres17000


} else {
ifres16998 = False


}

let__14428 := ifres16998
_ = let__14428

tmp17130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14428, False)
}
__typedArg0 := let__14428
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17130 {
tmp17128 := Call(__e, PrimFunc(symshen_4unlocked_2), V5558)


if True == tmp17128 {
tmp17012 := Call(__e, PrimFunc(symshen_4lazyderef), V5554, V5557)


let__14432 := tmp17012
_ = let__14432

tmp17126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14432)
}
__typedArg0 := let__14432
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17126 {
tmp17013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14432)
}
__typedArg0 := let__14432
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14433 := tmp17013
_ = let__14433

tmp17014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14432)
}
__typedArg0 := let__14432
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14434 := tmp17014
_ = let__14434

tmp17015 := Call(__e, PrimFunc(symshen_4lazyderef), V5555, V5557)


let__14435 := tmp17015
_ = let__14435

tmp17124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14435)
}
__typedArg0 := let__14435
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17124 {
tmp17016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14435)
}
__typedArg0 := let__14435
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14436 := tmp17016
_ = let__14436

tmp17017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14435)
}
__typedArg0 := let__14435
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17018 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17017, V5557)


let__14437 := tmp17018
_ = let__14437

tmp17122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14437)
}
__typedArg0 := let__14437
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17122 {
tmp17019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14437)
}
__typedArg0 := let__14437
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17020 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17019, V5557)


let__14438 := tmp17020
_ = let__14438

tmp17120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14438, sym_1_1_6)
}
__typedArg0 := let__14438
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17120 {
tmp17021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14437)
}
__typedArg0 := let__14437
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17022 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17021, V5557)


let__14439 := tmp17022
_ = let__14439

tmp17118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14439)
}
__typedArg0 := let__14439
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17118 {
tmp17023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14439)
}
__typedArg0 := let__14439
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14440 := tmp17023
_ = let__14440

tmp17024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14439)
}
__typedArg0 := let__14439
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17025 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17024, V5557)


let__14441 := tmp17025
_ = let__14441

tmp17116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14441, Nil)
}
__typedArg0 := let__14441
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17116 {
tmp17026 := Call(__e, PrimFunc(symshen_4lazyderef), V5556, V5557)


let__14442 := tmp17026
_ = let__14442

tmp17027 := MakeNative(func(__e *ControlFlow) {
Z5577 := __e.Get(1)
_ = Z5577
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5578 := __e.Get(1)
_ = Z5578
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5579 := __e.Get(1)
_ = Z5579
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5580 := __e.Get(1)
_ = Z5580
tmp17028 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17028

tmp17029 := MakeNative(func(__e *ControlFlow) {
tmp17030 := MakeNative(func(__e *ControlFlow) {
tmp17031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp17032 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4myassume), let__14434, let__14440, Z5580, V5557, V5558, V5559, V5560)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5578, tmp17031, V5557, V5558, V5559, tmp17032)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), let__14433, Z5577, V5557, V5558, V5559, tmp17030)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), let__14436, Z5579, V5557, V5558, V5559, tmp17029)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

let__14443 := tmp17027
_ = let__14443

tmp17114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14442)
}
__typedArg0 := let__14442
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17114 {
tmp17033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14442)
}
__typedArg0 := let__14442
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17034 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17033, V5557)


let__14444 := tmp17034
_ = let__14444

tmp17035 := MakeNative(func(__e *ControlFlow) {
Z5583 := __e.Get(1)
_ = Z5583
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5584 := __e.Get(1)
_ = Z5584
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5585 := __e.Get(1)
_ = Z5585
tmp17036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14442)
}
__typedArg0 := let__14442
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14446 := tmp17036
_ = let__14446

tmp17037 := Call(__e, let__14443, Z5583)


tmp17038 := Call(__e, tmp17037, Z5584)


tmp17039 := Call(__e, tmp17038, Z5585)


__e.TailApply(tmp17039, let__14446)
return


}, 1))
return
}, 1))
return
}, 1)

let__14445 := tmp17035
_ = let__14445

tmp17094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14444)
}
__typedArg0 := let__14444
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17094 {
tmp17040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14444)
}
__typedArg0 := let__14444
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14447 := tmp17040
_ = let__14447

tmp17041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14444)
}
__typedArg0 := let__14444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17042 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17041, V5557)


let__14448 := tmp17042
_ = let__14448

tmp17043 := MakeNative(func(__e *ControlFlow) {
Z5590 := __e.Get(1)
_ = Z5590
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5591 := __e.Get(1)
_ = Z5591
tmp17044 := Call(__e, let__14445, let__14447)


tmp17045 := Call(__e, tmp17044, Z5590)


__e.TailApply(tmp17045, Z5591)
return


}, 1))
return
}, 1)

let__14449 := tmp17043
_ = let__14449

tmp17078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14448)
}
__typedArg0 := let__14448
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17078 {
tmp17046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14448)
}
__typedArg0 := let__14448
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14450 := tmp17046
_ = let__14450

tmp17047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14448)
}
__typedArg0 := let__14448
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17048 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17047, V5557)


let__14451 := tmp17048
_ = let__14451

tmp17049 := MakeNative(func(__e *ControlFlow) {
Z5595 := __e.Get(1)
_ = Z5595
tmp17050 := Call(__e, let__14449, let__14450)


__e.TailApply(tmp17050, Z5595)
return


}, 1)

let__14452 := tmp17049
_ = let__14452

tmp17066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14451)
}
__typedArg0 := let__14451
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17066 {
tmp17051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14451)
}
__typedArg0 := let__14451
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14453 := tmp17051
_ = let__14453

tmp17052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14451)
}
__typedArg0 := let__14451
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17053 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17052, V5557)


let__14454 := tmp17053
_ = let__14454

tmp17054 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14452, let__14453)
return
}, 0)

let__14455 := tmp17054
_ = let__14455

tmp17058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14454, Nil)
}
__typedArg0 := let__14454
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17058 {
__e.TailApply(PrimFunc(symthaw), let__14455)
return
} else {
tmp17056 := Call(__e, PrimFunc(symshen_4pvar_2), let__14454)


if True == tmp17056 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14454, Nil, V5557, let__14455)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp17064 := Call(__e, PrimFunc(symshen_4pvar_2), let__14451)


if True == tmp17064 {
tmp17059 := Call(__e, PrimFunc(symshen_4newpv), V5557)


let__14456 := tmp17059
_ = let__14456

tmp17060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14456, Nil)
}
__typedArg0 := let__14456
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17061 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14452, let__14456)
return
}, 0)

tmp17062 := Call(__e, PrimFunc(symshen_4bind_b), let__14451, tmp17060, V5557, tmp17061)


__e.TailApply(PrimFunc(symshen_4gc), V5557, tmp17062)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17076 := Call(__e, PrimFunc(symshen_4pvar_2), let__14448)


if True == tmp17076 {
tmp17067 := Call(__e, PrimFunc(symshen_4newpv), V5557)


let__14457 := tmp17067
_ = let__14457

tmp17068 := Call(__e, PrimFunc(symshen_4newpv), V5557)


let__14458 := tmp17068
_ = let__14458

tmp17069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14458, Nil)
}
__typedArg0 := let__14458
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14457, tmp17069)
}
__typedArg0 := let__14457
__typedArg1 := tmp17069
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17071 := MakeNative(func(__e *ControlFlow) {
tmp17072 := Call(__e, let__14449, let__14457)


__e.TailApply(tmp17072, let__14458)
return


}, 0)

tmp17073 := Call(__e, PrimFunc(symshen_4bind_b), let__14448, tmp17070, V5557, tmp17071)


tmp17074 := Call(__e, PrimFunc(symshen_4gc), V5557, tmp17073)


__e.TailApply(PrimFunc(symshen_4gc), V5557, tmp17074)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17092 := Call(__e, PrimFunc(symshen_4pvar_2), let__14444)


if True == tmp17092 {
tmp17079 := Call(__e, PrimFunc(symshen_4newpv), V5557)


let__14459 := tmp17079
_ = let__14459

tmp17080 := Call(__e, PrimFunc(symshen_4newpv), V5557)


let__14460 := tmp17080
_ = let__14460

tmp17081 := Call(__e, PrimFunc(symshen_4newpv), V5557)


let__14461 := tmp17081
_ = let__14461

tmp17082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14461, Nil)
}
__typedArg0 := let__14461
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14460, tmp17082)
}
__typedArg0 := let__14460
__typedArg1 := tmp17082
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14459, tmp17083)
}
__typedArg0 := let__14459
__typedArg1 := tmp17083
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17085 := MakeNative(func(__e *ControlFlow) {
tmp17086 := Call(__e, let__14445, let__14459)


tmp17087 := Call(__e, tmp17086, let__14460)


__e.TailApply(tmp17087, let__14461)
return


}, 0)

tmp17088 := Call(__e, PrimFunc(symshen_4bind_b), let__14444, tmp17084, V5557, tmp17085)


tmp17089 := Call(__e, PrimFunc(symshen_4gc), V5557, tmp17088)


tmp17090 := Call(__e, PrimFunc(symshen_4gc), V5557, tmp17089)


__e.TailApply(PrimFunc(symshen_4gc), V5557, tmp17090)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17112 := Call(__e, PrimFunc(symshen_4pvar_2), let__14442)


if True == tmp17112 {
tmp17095 := Call(__e, PrimFunc(symshen_4newpv), V5557)


let__14462 := tmp17095
_ = let__14462

tmp17096 := Call(__e, PrimFunc(symshen_4newpv), V5557)


let__14463 := tmp17096
_ = let__14463

tmp17097 := Call(__e, PrimFunc(symshen_4newpv), V5557)


let__14464 := tmp17097
_ = let__14464

tmp17098 := Call(__e, PrimFunc(symshen_4newpv), V5557)


let__14465 := tmp17098
_ = let__14465

tmp17099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14464, Nil)
}
__typedArg0 := let__14464
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14463, tmp17099)
}
__typedArg0 := let__14463
__typedArg1 := tmp17099
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14462, tmp17100)
}
__typedArg0 := let__14462
__typedArg1 := tmp17100
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17101, let__14465)
}
__typedArg0 := tmp17101
__typedArg1 := let__14465
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17103 := MakeNative(func(__e *ControlFlow) {
tmp17104 := Call(__e, let__14443, let__14462)


tmp17105 := Call(__e, tmp17104, let__14463)


tmp17106 := Call(__e, tmp17105, let__14464)


__e.TailApply(tmp17106, let__14465)
return


}, 0)

tmp17107 := Call(__e, PrimFunc(symshen_4bind_b), let__14442, tmp17102, V5557, tmp17103)


tmp17108 := Call(__e, PrimFunc(symshen_4gc), V5557, tmp17107)


tmp17109 := Call(__e, PrimFunc(symshen_4gc), V5557, tmp17108)


tmp17110 := Call(__e, PrimFunc(symshen_4gc), V5557, tmp17109)


__e.TailApply(PrimFunc(symshen_4gc), V5557, tmp17110)
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

tmp17131 := Call(__e, ns2_1set, symshen_4myassume, tmp16997)


_ = tmp17131

tmp17132 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V5611 := __self1
_ = V5611
tmp17155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5611)
}
__typedArg0 := Nil
__typedArg1 := V5611
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17155 {
__e.Return(Nil)
return
} else {
tmp17153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5611)
}
__typedArg0 := V5611
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17149 Obj

if True == tmp17153 {
tmp17151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5611)
}
__typedArg0 := V5611
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17151)
}
__typedArg0 := tmp17151
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17150 Obj

if True == tmp17152 {
ifres17150 = True


} else {
ifres17150 = False


}

ifres17149 = ifres17150


} else {
ifres17149 = False


}

if True == ifres17149 {
tmp17133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5611)
}
__typedArg0 := V5611
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5611)
}
__typedArg0 := V5611
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17135 := Call(__e, PrimFunc(symappend), tmp17133, tmp17134)


if PrimFunc(symshen_4freshterms) == __self {
__self1 = tmp17135
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4freshterms), tmp17135)
return


} else {
tmp17147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5611)
}
__typedArg0 := V5611
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17143 Obj

if True == tmp17147 {
tmp17145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5611)
}
__typedArg0 := V5611
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17146 := Call(__e, PrimFunc(symshen_4freshterm_2), tmp17145)


var ifres17144 Obj

if True == tmp17146 {
ifres17144 = True


} else {
ifres17144 = False


}

ifres17143 = ifres17144


} else {
ifres17143 = False


}

if True == ifres17143 {
tmp17136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5611)
}
__typedArg0 := V5611
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5611)
}
__typedArg0 := V5611
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17138 := Call(__e, PrimFunc(symshen_4freshterms), tmp17137)


__e.TailApply(PrimFunc(symadjoin), tmp17136, tmp17138)
return


} else {
tmp17141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5611)
}
__typedArg0 := V5611
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17141 {
tmp17139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5611)
}
__typedArg0 := V5611
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4freshterms) == __self {
__self1 = tmp17139
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4freshterms), tmp17139)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshterms)
return
}


}


}


}


}, 1)

tmp17156 := Call(__e, ns2_1set, symshen_4freshterms, tmp17132)


_ = tmp17156

tmp17157 := MakeNative(func(__e *ControlFlow) {
V5612 := __e.Get(1)
_ = V5612
V5613 := __e.Get(2)
_ = V5613
V5614 := __e.Get(3)
_ = V5614
V5615 := __e.Get(4)
_ = V5615
V5616 := __e.Get(5)
_ = V5616
V5617 := __e.Get(6)
_ = V5617
tmp17171 := Call(__e, PrimFunc(symshen_4unlocked_2), V5615)


var ifres17158 Obj

if True == tmp17171 {
tmp17159 := Call(__e, PrimFunc(symshen_4lazyderef), V5612, V5614)


let__14467 := tmp17159
_ = let__14467

tmp17170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14467, Nil)
}
__typedArg0 := let__14467
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17160 Obj

if True == tmp17170 {
tmp17161 := Call(__e, PrimFunc(symshen_4lazyderef), V5613, V5614)


let__14468 := tmp17161
_ = let__14468

tmp17162 := MakeNative(func(__e *ControlFlow) {
tmp17163 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17163

__e.TailApply(PrimFunc(symthaw), V5617)
return


}, 0)

let__14469 := tmp17162
_ = let__14469

tmp17169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14468, Nil)
}
__typedArg0 := let__14468
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17164 Obj

if True == tmp17169 {
tmp17165 := Call(__e, PrimFunc(symthaw), let__14469)


ifres17164 = tmp17165


} else {
tmp17168 := Call(__e, PrimFunc(symshen_4pvar_2), let__14468)


var ifres17166 Obj

if True == tmp17168 {
tmp17167 := Call(__e, PrimFunc(symshen_4bind_b), let__14468, Nil, V5614, let__14469)


ifres17166 = tmp17167


} else {
ifres17166 = False


}

ifres17164 = ifres17166


}

ifres17160 = ifres17164


} else {
ifres17160 = False


}

ifres17158 = ifres17160


} else {
ifres17158 = False


}

let__14466 := ifres17158
_ = let__14466

tmp17268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14466, False)
}
__typedArg0 := let__14466
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17268 {
tmp17266 := Call(__e, PrimFunc(symshen_4unlocked_2), V5615)


if True == tmp17266 {
tmp17172 := Call(__e, PrimFunc(symshen_4lazyderef), V5612, V5614)


let__14470 := tmp17172
_ = let__14470

tmp17264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14470)
}
__typedArg0 := let__14470
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17264 {
tmp17173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14470)
}
__typedArg0 := let__14470
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14471 := tmp17173
_ = let__14471

tmp17174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14470)
}
__typedArg0 := let__14470
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14472 := tmp17174
_ = let__14472

tmp17175 := Call(__e, PrimFunc(symshen_4lazyderef), V5613, V5614)


let__14473 := tmp17175
_ = let__14473

tmp17176 := MakeNative(func(__e *ControlFlow) {
Z5627 := __e.Get(1)
_ = Z5627
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5628 := __e.Get(1)
_ = Z5628
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5629 := __e.Get(1)
_ = Z5629
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5630 := __e.Get(1)
_ = Z5630
tmp17177 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17177

tmp17178 := MakeNative(func(__e *ControlFlow) {
tmp17179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp17180 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4p_1hyps), let__14472, Z5630, V5614, V5615, V5616, V5617)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5628, tmp17179, V5614, V5615, V5616, tmp17180)
return


}, 0)

__e.TailApply(PrimFunc(symbind), Z5627, let__14471, V5614, V5615, V5616, tmp17178)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

let__14474 := tmp17176
_ = let__14474

tmp17262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14473)
}
__typedArg0 := let__14473
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17262 {
tmp17181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14473)
}
__typedArg0 := let__14473
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17182 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17181, V5614)


let__14475 := tmp17182
_ = let__14475

tmp17183 := MakeNative(func(__e *ControlFlow) {
Z5633 := __e.Get(1)
_ = Z5633
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5634 := __e.Get(1)
_ = Z5634
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5635 := __e.Get(1)
_ = Z5635
tmp17184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14473)
}
__typedArg0 := let__14473
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14477 := tmp17184
_ = let__14477

tmp17185 := Call(__e, let__14474, Z5633)


tmp17186 := Call(__e, tmp17185, Z5634)


tmp17187 := Call(__e, tmp17186, Z5635)


__e.TailApply(tmp17187, let__14477)
return


}, 1))
return
}, 1))
return
}, 1)

let__14476 := tmp17183
_ = let__14476

tmp17242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14475)
}
__typedArg0 := let__14475
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17242 {
tmp17188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14475)
}
__typedArg0 := let__14475
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14478 := tmp17188
_ = let__14478

tmp17189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14475)
}
__typedArg0 := let__14475
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17190 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17189, V5614)


let__14479 := tmp17190
_ = let__14479

tmp17191 := MakeNative(func(__e *ControlFlow) {
Z5640 := __e.Get(1)
_ = Z5640
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5641 := __e.Get(1)
_ = Z5641
tmp17192 := Call(__e, let__14476, let__14478)


tmp17193 := Call(__e, tmp17192, Z5640)


__e.TailApply(tmp17193, Z5641)
return


}, 1))
return
}, 1)

let__14480 := tmp17191
_ = let__14480

tmp17226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14479)
}
__typedArg0 := let__14479
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17226 {
tmp17194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14479)
}
__typedArg0 := let__14479
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14481 := tmp17194
_ = let__14481

tmp17195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14479)
}
__typedArg0 := let__14479
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17196 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17195, V5614)


let__14482 := tmp17196
_ = let__14482

tmp17197 := MakeNative(func(__e *ControlFlow) {
Z5645 := __e.Get(1)
_ = Z5645
tmp17198 := Call(__e, let__14480, let__14481)


__e.TailApply(tmp17198, Z5645)
return


}, 1)

let__14483 := tmp17197
_ = let__14483

tmp17214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14482)
}
__typedArg0 := let__14482
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17214 {
tmp17199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14482)
}
__typedArg0 := let__14482
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14484 := tmp17199
_ = let__14484

tmp17200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14482)
}
__typedArg0 := let__14482
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17201 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17200, V5614)


let__14485 := tmp17201
_ = let__14485

tmp17202 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14483, let__14484)
return
}, 0)

let__14486 := tmp17202
_ = let__14486

tmp17206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14485, Nil)
}
__typedArg0 := let__14485
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17206 {
__e.TailApply(PrimFunc(symthaw), let__14486)
return
} else {
tmp17204 := Call(__e, PrimFunc(symshen_4pvar_2), let__14485)


if True == tmp17204 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14485, Nil, V5614, let__14486)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp17212 := Call(__e, PrimFunc(symshen_4pvar_2), let__14482)


if True == tmp17212 {
tmp17207 := Call(__e, PrimFunc(symshen_4newpv), V5614)


let__14487 := tmp17207
_ = let__14487

tmp17208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14487, Nil)
}
__typedArg0 := let__14487
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17209 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14483, let__14487)
return
}, 0)

tmp17210 := Call(__e, PrimFunc(symshen_4bind_b), let__14482, tmp17208, V5614, tmp17209)


__e.TailApply(PrimFunc(symshen_4gc), V5614, tmp17210)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17224 := Call(__e, PrimFunc(symshen_4pvar_2), let__14479)


if True == tmp17224 {
tmp17215 := Call(__e, PrimFunc(symshen_4newpv), V5614)


let__14488 := tmp17215
_ = let__14488

tmp17216 := Call(__e, PrimFunc(symshen_4newpv), V5614)


let__14489 := tmp17216
_ = let__14489

tmp17217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14489, Nil)
}
__typedArg0 := let__14489
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14488, tmp17217)
}
__typedArg0 := let__14488
__typedArg1 := tmp17217
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17219 := MakeNative(func(__e *ControlFlow) {
tmp17220 := Call(__e, let__14480, let__14488)


__e.TailApply(tmp17220, let__14489)
return


}, 0)

tmp17221 := Call(__e, PrimFunc(symshen_4bind_b), let__14479, tmp17218, V5614, tmp17219)


tmp17222 := Call(__e, PrimFunc(symshen_4gc), V5614, tmp17221)


__e.TailApply(PrimFunc(symshen_4gc), V5614, tmp17222)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17240 := Call(__e, PrimFunc(symshen_4pvar_2), let__14475)


if True == tmp17240 {
tmp17227 := Call(__e, PrimFunc(symshen_4newpv), V5614)


let__14490 := tmp17227
_ = let__14490

tmp17228 := Call(__e, PrimFunc(symshen_4newpv), V5614)


let__14491 := tmp17228
_ = let__14491

tmp17229 := Call(__e, PrimFunc(symshen_4newpv), V5614)


let__14492 := tmp17229
_ = let__14492

tmp17230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14492, Nil)
}
__typedArg0 := let__14492
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14491, tmp17230)
}
__typedArg0 := let__14491
__typedArg1 := tmp17230
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14490, tmp17231)
}
__typedArg0 := let__14490
__typedArg1 := tmp17231
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17233 := MakeNative(func(__e *ControlFlow) {
tmp17234 := Call(__e, let__14476, let__14490)


tmp17235 := Call(__e, tmp17234, let__14491)


__e.TailApply(tmp17235, let__14492)
return


}, 0)

tmp17236 := Call(__e, PrimFunc(symshen_4bind_b), let__14475, tmp17232, V5614, tmp17233)


tmp17237 := Call(__e, PrimFunc(symshen_4gc), V5614, tmp17236)


tmp17238 := Call(__e, PrimFunc(symshen_4gc), V5614, tmp17237)


__e.TailApply(PrimFunc(symshen_4gc), V5614, tmp17238)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17260 := Call(__e, PrimFunc(symshen_4pvar_2), let__14473)


if True == tmp17260 {
tmp17243 := Call(__e, PrimFunc(symshen_4newpv), V5614)


let__14493 := tmp17243
_ = let__14493

tmp17244 := Call(__e, PrimFunc(symshen_4newpv), V5614)


let__14494 := tmp17244
_ = let__14494

tmp17245 := Call(__e, PrimFunc(symshen_4newpv), V5614)


let__14495 := tmp17245
_ = let__14495

tmp17246 := Call(__e, PrimFunc(symshen_4newpv), V5614)


let__14496 := tmp17246
_ = let__14496

tmp17247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14495, Nil)
}
__typedArg0 := let__14495
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14494, tmp17247)
}
__typedArg0 := let__14494
__typedArg1 := tmp17247
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14493, tmp17248)
}
__typedArg0 := let__14493
__typedArg1 := tmp17248
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17249, let__14496)
}
__typedArg0 := tmp17249
__typedArg1 := let__14496
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17251 := MakeNative(func(__e *ControlFlow) {
tmp17252 := Call(__e, let__14474, let__14493)


tmp17253 := Call(__e, tmp17252, let__14494)


tmp17254 := Call(__e, tmp17253, let__14495)


__e.TailApply(tmp17254, let__14496)
return


}, 0)

tmp17255 := Call(__e, PrimFunc(symshen_4bind_b), let__14473, tmp17250, V5614, tmp17251)


tmp17256 := Call(__e, PrimFunc(symshen_4gc), V5614, tmp17255)


tmp17257 := Call(__e, PrimFunc(symshen_4gc), V5614, tmp17256)


tmp17258 := Call(__e, PrimFunc(symshen_4gc), V5614, tmp17257)


__e.TailApply(PrimFunc(symshen_4gc), V5614, tmp17258)
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

tmp17269 := Call(__e, ns2_1set, symshen_4p_1hyps, tmp17157)


_ = tmp17269

tmp17270 := MakeNative(func(__e *ControlFlow) {
V5659 := __e.Get(1)
_ = V5659
V5660 := __e.Get(2)
_ = V5660
V5661 := __e.Get(3)
_ = V5661
V5662 := __e.Get(4)
_ = V5662
V5663 := __e.Get(5)
_ = V5663
V5664 := __e.Get(6)
_ = V5664
V5665 := __e.Get(7)
_ = V5665
let__14497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5664)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5664
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__14497

tmp17308 := Call(__e, PrimFunc(symshen_4unlocked_2), V5663)


var ifres17272 Obj

if True == tmp17308 {
tmp17273 := Call(__e, PrimFunc(symshen_4lazyderef), V5659, V5662)


let__14499 := tmp17273
_ = let__14499

tmp17307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14499)
}
__typedArg0 := let__14499
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17274 Obj

if True == tmp17307 {
tmp17275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14499)
}
__typedArg0 := let__14499
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17276 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17275, V5662)


let__14500 := tmp17276
_ = let__14500

tmp17306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14500, symwhere)
}
__typedArg0 := let__14500
__typedArg1 := symwhere
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17277 Obj

if True == tmp17306 {
tmp17278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14499)
}
__typedArg0 := let__14499
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17279 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17278, V5662)


let__14501 := tmp17279
_ = let__14501

tmp17305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14501)
}
__typedArg0 := let__14501
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17280 Obj

if True == tmp17305 {
tmp17281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14501)
}
__typedArg0 := let__14501
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14502 := tmp17281
_ = let__14502

tmp17282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14501)
}
__typedArg0 := let__14501
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17283 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17282, V5662)


let__14503 := tmp17283
_ = let__14503

tmp17304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14503)
}
__typedArg0 := let__14503
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17284 Obj

if True == tmp17304 {
tmp17285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14503)
}
__typedArg0 := let__14503
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14504 := tmp17285
_ = let__14504

tmp17286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14503)
}
__typedArg0 := let__14503
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17287 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17286, V5662)


let__14505 := tmp17287
_ = let__14505

tmp17303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14505, Nil)
}
__typedArg0 := let__14505
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17288 Obj

if True == tmp17303 {
tmp17289 := Call(__e, PrimFunc(symshen_4newpv), V5662)


let__14506 := tmp17289
_ = let__14506

tmp17290 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17290

tmp17291 := MakeNative(func(__e *ControlFlow) {
tmp17292 := Call(__e, PrimFunc(symshen_4curry), let__14502)


tmp17293 := MakeNative(func(__e *ControlFlow) {
tmp17294 := MakeNative(func(__e *ControlFlow) {
tmp17295 := MakeNative(func(__e *ControlFlow) {
tmp17296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp17297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symverified, Nil)
}
__typedArg0 := symverified
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17296, tmp17297)
}
__typedArg0 := tmp17296
__typedArg1 := tmp17297
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14506, tmp17298)
}
__typedArg0 := let__14506
__typedArg1 := tmp17298
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17299, V5661)
}
__typedArg0 := tmp17299
__typedArg1 := V5661
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4t_d_1correct), let__14504, V5660, tmp17300, V5662, V5663, let__14497, V5665)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5662, V5663, let__14497, tmp17295)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14506, symboolean, V5661, V5662, V5663, let__14497, tmp17294)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14506, tmp17292, V5662, V5663, let__14497, tmp17293)
return


}, 0)

tmp17301 := Call(__e, PrimFunc(symshen_4cut), V5662, V5663, let__14497, tmp17291)


tmp17302 := Call(__e, PrimFunc(symshen_4gc), V5662, tmp17301)


ifres17288 = tmp17302


} else {
ifres17288 = False


}

ifres17284 = ifres17288


} else {
ifres17284 = False


}

ifres17280 = ifres17284


} else {
ifres17280 = False


}

ifres17277 = ifres17280


} else {
ifres17277 = False


}

ifres17274 = ifres17277


} else {
ifres17274 = False


}

ifres17272 = ifres17274


} else {
ifres17272 = False


}

let__14498 := ifres17272
_ = let__14498

tmp17317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14498, False)
}
__typedArg0 := let__14498
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17317 {
tmp17313 := Call(__e, PrimFunc(symshen_4unlocked_2), V5663)


var ifres17309 Obj

if True == tmp17313 {
tmp17310 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17310

tmp17311 := Call(__e, PrimFunc(symshen_4curry), V5659)


tmp17312 := Call(__e, PrimFunc(symshen_4system_1S_1h), tmp17311, V5660, V5661, V5662, V5663, let__14497, V5665)


ifres17309 = tmp17312


} else {
ifres17309 = False


}

let__14507 := ifres17309
_ = let__14507

tmp17315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14507, False)
}
__typedArg0 := let__14507
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17315 {
__e.TailApply(PrimFunc(symshen_4unlock), V5663, let__14497)
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

tmp17318 := Call(__e, ns2_1set, symshen_4t_d_1correct, tmp17270)


_ = tmp17318

tmp17319 := MakeNative(func(__e *ControlFlow) {
V5677 := __e.Get(1)
_ = V5677
V5678 := __e.Get(2)
_ = V5678
V5679 := __e.Get(3)
_ = V5679
V5680 := __e.Get(4)
_ = V5680
V5681 := __e.Get(5)
_ = V5681
V5682 := __e.Get(6)
_ = V5682
V5683 := __e.Get(7)
_ = V5683
V5684 := __e.Get(8)
_ = V5684
tmp17326 := Call(__e, PrimFunc(symshen_4unlocked_2), V5682)


var ifres17320 Obj

if True == tmp17326 {
tmp17321 := Call(__e, PrimFunc(symshen_4lazyderef), V5677, V5681)


let__14509 := tmp17321
_ = let__14509

tmp17325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14509, Nil)
}
__typedArg0 := let__14509
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17322 Obj

if True == tmp17325 {
tmp17323 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17323

tmp17324 := Call(__e, PrimFunc(symis_b), V5678, V5680, V5681, V5682, V5683, V5684)


ifres17322 = tmp17324


} else {
ifres17322 = False


}

ifres17320 = ifres17322


} else {
ifres17320 = False


}

let__14508 := ifres17320
_ = let__14508

tmp17358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14508, False)
}
__typedArg0 := let__14508
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17358 {
tmp17356 := Call(__e, PrimFunc(symshen_4unlocked_2), V5682)


if True == tmp17356 {
tmp17327 := Call(__e, PrimFunc(symshen_4lazyderef), V5677, V5681)


let__14510 := tmp17327
_ = let__14510

tmp17354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14510)
}
__typedArg0 := let__14510
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17354 {
tmp17328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14510)
}
__typedArg0 := let__14510
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14511 := tmp17328
_ = let__14511

tmp17329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14510)
}
__typedArg0 := let__14510
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14512 := tmp17329
_ = let__14512

tmp17330 := Call(__e, PrimFunc(symshen_4lazyderef), V5678, V5681)


let__14513 := tmp17330
_ = let__14513

tmp17352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14513)
}
__typedArg0 := let__14513
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17352 {
tmp17331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14513)
}
__typedArg0 := let__14513
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14514 := tmp17331
_ = let__14514

tmp17332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14513)
}
__typedArg0 := let__14513
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17333 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17332, V5681)


let__14515 := tmp17333
_ = let__14515

tmp17350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14515)
}
__typedArg0 := let__14515
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17350 {
tmp17334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14515)
}
__typedArg0 := let__14515
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17335 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17334, V5681)


let__14516 := tmp17335
_ = let__14516

tmp17348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14516, sym_1_1_6)
}
__typedArg0 := let__14516
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17348 {
tmp17336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14515)
}
__typedArg0 := let__14515
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17337 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17336, V5681)


let__14517 := tmp17337
_ = let__14517

tmp17346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14517)
}
__typedArg0 := let__14517
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17346 {
tmp17338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14517)
}
__typedArg0 := let__14517
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14518 := tmp17338
_ = let__14518

tmp17339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14517)
}
__typedArg0 := let__14517
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17340 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17339, V5681)


let__14519 := tmp17340
_ = let__14519

tmp17344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14519, Nil)
}
__typedArg0 := let__14519
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17344 {
tmp17341 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17341

tmp17342 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1integrity), let__14512, let__14518, V5679, V5680, V5681, V5682, V5683, V5684)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14511, let__14514, V5679, V5681, V5682, V5683, tmp17342)
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

tmp17359 := Call(__e, ns2_1set, symshen_4t_d_1integrity, tmp17319)


_ = tmp17359

tmp17360 := MakeNative(func(__e *ControlFlow) {
V5697 := __e.Get(1)
_ = V5697
tmp17369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V5697)
}
__typedArg0 := V5697
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp17369 {
tmp17366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V5697)
}
__typedArg0 := V5697
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

tmp17367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp17366)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp17366
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres17362 Obj

if True == tmp17367 {
tmp17364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V5697, MakeInteger(0))
}
__typedArg0 := V5697
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp17365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp17364, symshen_4print_1freshterm)
}
__typedArg0 := tmp17364
__typedArg1 := symshen_4print_1freshterm
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17363 Obj

if True == tmp17365 {
ifres17363 = True


} else {
ifres17363 = False


}

ifres17362 = ifres17363


} else {
ifres17362 = False


}

if True == ifres17362 {
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

__e.TailApply(ns2_1set, symshen_4freshterm_2, tmp17360)
return




}, 0)

