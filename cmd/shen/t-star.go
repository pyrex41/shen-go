package main

import . "github.com/pyrex41/shen-go/kl"

var TStarMain = MakeNative(func(__e *ControlFlow) {
tmp14513 := MakeNative(func(__e *ControlFlow) {
V4921 := __e.Get(1)
_ = V4921
V4922 := __e.Get(2)
_ = V4922
tmp14514 := Call(__e, PrimFunc(symshen_4extract_1vars), V4922)


let__13962 := tmp14514
_ = let__13962

tmp14515 := Call(__e, PrimFunc(symshen_4rectify_1type), V4922)


let__13963 := tmp14515
_ = let__13963

tmp14516 := Call(__e, PrimFunc(symshen_4curry), V4921)


let__13964 := tmp14516
_ = let__13964

tmp14517 := MakeNative(func(__e *ControlFlow) {
Z4926 := __e.Get(1)
_ = Z4926
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4927 := __e.Get(1)
_ = Z4927
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4928 := __e.Get(1)
_ = Z4928
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4929 := __e.Get(1)
_ = Z4929
tmp14518 := Call(__e, PrimFunc(symshen_4newpv), Z4926)


let__13965 := tmp14518
_ = let__13965

tmp14519 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14519

tmp14520 := Call(__e, PrimFunc(symshen_4deref), let__13962, Z4926)


tmp14521 := Call(__e, PrimFunc(symreceive), tmp14520)


tmp14522 := Call(__e, PrimFunc(symshen_4deref), let__13963, Z4926)


tmp14523 := Call(__e, PrimFunc(symreceive), tmp14522)


tmp14524 := MakeNative(func(__e *ControlFlow) {
tmp14525 := Call(__e, PrimFunc(symshen_4deref), let__13964, Z4926)


tmp14526 := Call(__e, PrimFunc(symreceive), tmp14525)


tmp14527 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symreturn), let__13965, Z4926, Z4927, Z4928, Z4929)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4toplevel_1forms), tmp14526, let__13965, Z4926, Z4927, Z4928, tmp14527)
return


}, 0)

tmp14528 := Call(__e, PrimFunc(symshen_4insert_1prolog_1variables), tmp14521, tmp14523, let__13965, Z4926, Z4927, Z4928, tmp14524)


__e.TailApply(PrimFunc(symshen_4gc), Z4926, tmp14528)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp14529 := Call(__e, PrimFunc(symshen_4prolog_1vector))


tmp14530 := Call(__e, tmp14517, tmp14529)


tmp14531 := Call(__e, PrimFunc(symvector), MakeInteger(0))


tmp14532 := Call(__e, PrimFunc(sym_8v), MakeInteger(0), tmp14531)


tmp14533 := Call(__e, PrimFunc(sym_8v), True, tmp14532)


tmp14534 := Call(__e, tmp14530, tmp14533)


tmp14535 := Call(__e, tmp14534, MakeInteger(0))


tmp14536 := MakeNative(func(__e *ControlFlow) {
__e.Return(True)
return
}, 0)

__e.TailApply(tmp14535, tmp14536)
return


}, 2)

tmp14537 := Call(__e, ns2_1set, symshen_4typecheck, tmp14513)


_ = tmp14537

tmp14538 := MakeNative(func(__e *ControlFlow) {
V4931 := __e.Get(1)
_ = V4931
V4932 := __e.Get(2)
_ = V4932
V4933 := __e.Get(3)
_ = V4933
V4934 := __e.Get(4)
_ = V4934
V4935 := __e.Get(5)
_ = V4935
V4936 := __e.Get(6)
_ = V4936
V4937 := __e.Get(7)
_ = V4937
tmp14545 := Call(__e, PrimFunc(symshen_4unlocked_2), V4935)


var ifres14539 Obj

if True == tmp14545 {
tmp14540 := Call(__e, PrimFunc(symshen_4lazyderef), V4931, V4934)


let__13967 := tmp14540
_ = let__13967

tmp14544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13967, Nil)
}
__typedArg0 := let__13967
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14541 Obj

if True == tmp14544 {
tmp14542 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14542

tmp14543 := Call(__e, PrimFunc(symis_b), V4932, V4933, V4934, V4935, V4936, V4937)


ifres14541 = tmp14543


} else {
ifres14541 = False


}

ifres14539 = ifres14541


} else {
ifres14539 = False


}

let__13966 := ifres14539
_ = let__13966

tmp14559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13966, False)
}
__typedArg0 := let__13966
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14559 {
tmp14557 := Call(__e, PrimFunc(symshen_4unlocked_2), V4935)


if True == tmp14557 {
tmp14546 := Call(__e, PrimFunc(symshen_4lazyderef), V4931, V4934)


let__13968 := tmp14546
_ = let__13968

tmp14555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__13968)
}
__typedArg0 := let__13968
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14555 {
tmp14547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__13968)
}
__typedArg0 := let__13968
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__13969 := tmp14547
_ = let__13969

tmp14548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__13968)
}
__typedArg0 := let__13968
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__13970 := tmp14548
_ = let__13970

tmp14549 := Call(__e, PrimFunc(symshen_4newpv), V4934)


let__13971 := tmp14549
_ = let__13971

tmp14550 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14550

tmp14551 := Call(__e, PrimFunc(symshen_4deref), let__13971, V4934)


tmp14552 := Call(__e, PrimFunc(symsubst), tmp14551, let__13969, V4932)


tmp14553 := Call(__e, PrimFunc(symshen_4insert_1prolog_1variables), let__13970, tmp14552, V4933, V4934, V4935, V4936, V4937)


__e.TailApply(PrimFunc(symshen_4gc), V4934, tmp14553)
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
__e.Return(let__13966)
return
}


}, 7)

tmp14560 := Call(__e, ns2_1set, symshen_4insert_1prolog_1variables, tmp14538)


_ = tmp14560

tmp14561 := MakeNative(func(__e *ControlFlow) {
V4944 := __e.Get(1)
_ = V4944
V4945 := __e.Get(2)
_ = V4945
V4946 := __e.Get(3)
_ = V4946
V4947 := __e.Get(4)
_ = V4947
V4948 := __e.Get(5)
_ = V4948
V4949 := __e.Get(6)
_ = V4949
let__13972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4948)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4948
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__13972

tmp14586 := Call(__e, PrimFunc(symshen_4unlocked_2), V4947)


var ifres14563 Obj

if True == tmp14586 {
tmp14564 := Call(__e, PrimFunc(symshen_4lazyderef), V4944, V4946)


let__13974 := tmp14564
_ = let__13974

tmp14585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__13974)
}
__typedArg0 := let__13974
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14565 Obj

if True == tmp14585 {
tmp14566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__13974)
}
__typedArg0 := let__13974
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14567 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14566, V4946)


let__13975 := tmp14567
_ = let__13975

tmp14584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13975, symdefine)
}
__typedArg0 := let__13975
__typedArg1 := symdefine
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14568 Obj

if True == tmp14584 {
tmp14569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__13974)
}
__typedArg0 := let__13974
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14570 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14569, V4946)


let__13976 := tmp14570
_ = let__13976

tmp14583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__13976)
}
__typedArg0 := let__13976
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14571 Obj

if True == tmp14583 {
tmp14572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__13976)
}
__typedArg0 := let__13976
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__13977 := tmp14572
_ = let__13977

tmp14573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__13976)
}
__typedArg0 := let__13976
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__13978 := tmp14573
_ = let__13978

tmp14574 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14574

tmp14575 := Call(__e, PrimFunc(symshen_4type_1theory_1enabled_2))


tmp14576 := MakeNative(func(__e *ControlFlow) {
tmp14577 := MakeNative(func(__e *ControlFlow) {
tmp14578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspy_d)
}
__typedArg0 := symshen_4_dspy_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp14579 := MakeNative(func(__e *ControlFlow) {
tmp14580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__13977, let__13978)
}
__typedArg0 := let__13977
__typedArg1 := let__13978
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp14580)
}
__typedArg0 := symdefine
__typedArg1 := tmp14580
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4t_d), tmp14581, V4945, V4946, V4947, let__13972, V4949)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4signal_1def), tmp14578, let__13977, V4946, V4947, let__13972, tmp14579)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4946, V4947, let__13972, tmp14577)
return


}, 0)

tmp14582 := Call(__e, PrimFunc(symwhen), tmp14575, V4946, V4947, let__13972, tmp14576)


ifres14571 = tmp14582


} else {
ifres14571 = False


}

ifres14568 = ifres14571


} else {
ifres14568 = False


}

ifres14565 = ifres14568


} else {
ifres14565 = False


}

ifres14563 = ifres14565


} else {
ifres14563 = False


}

let__13973 := ifres14563
_ = let__13973

tmp14598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13973, False)
}
__typedArg0 := let__13973
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14598 {
tmp14594 := Call(__e, PrimFunc(symshen_4unlocked_2), V4947)


var ifres14587 Obj

if True == tmp14594 {
tmp14588 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14588

tmp14589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp14590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4945, Nil)
}
__typedArg0 := V4945
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14589, tmp14590)
}
__typedArg0 := tmp14589
__typedArg1 := tmp14590
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4944, tmp14591)
}
__typedArg0 := V4944
__typedArg1 := tmp14591
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14593 := Call(__e, PrimFunc(symshen_4system_1S), tmp14592, Nil, V4946, V4947, let__13972, V4949)


ifres14587 = tmp14593


} else {
ifres14587 = False


}

let__13979 := ifres14587
_ = let__13979

tmp14596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13979, False)
}
__typedArg0 := let__13979
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14596 {
__e.TailApply(PrimFunc(symshen_4unlock), V4947, let__13972)
return
} else {
__e.Return(let__13979)
return
}


} else {
__e.Return(let__13973)
return
}


}, 6)

tmp14599 := Call(__e, ns2_1set, symshen_4toplevel_1forms, tmp14561)


_ = tmp14599

tmp14600 := MakeNative(func(__e *ControlFlow) {
V4958 := __e.Get(1)
_ = V4958
V4959 := __e.Get(2)
_ = V4959
V4960 := __e.Get(3)
_ = V4960
V4961 := __e.Get(4)
_ = V4961
V4962 := __e.Get(5)
_ = V4962
V4963 := __e.Get(6)
_ = V4963
tmp14607 := Call(__e, PrimFunc(symshen_4unlocked_2), V4961)


var ifres14601 Obj

if True == tmp14607 {
tmp14602 := Call(__e, PrimFunc(symshen_4lazyderef), V4958, V4960)


let__13981 := tmp14602
_ = let__13981

tmp14606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13981, False)
}
__typedArg0 := let__13981
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14603 Obj

if True == tmp14606 {
tmp14604 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14604

tmp14605 := Call(__e, PrimFunc(symthaw), V4963)


ifres14603 = tmp14605


} else {
ifres14603 = False


}

ifres14601 = ifres14603


} else {
ifres14601 = False


}

let__13980 := ifres14601
_ = let__13980

tmp14622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13980, False)
}
__typedArg0 := let__13980
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14622 {
tmp14620 := Call(__e, PrimFunc(symshen_4unlocked_2), V4961)


if True == tmp14620 {
tmp14608 := Call(__e, PrimFunc(symshen_4lazyderef), V4958, V4960)


let__13982 := tmp14608
_ = let__13982

tmp14618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13982, True)
}
__typedArg0 := let__13982
__typedArg1 := True
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14618 {
tmp14609 := Call(__e, PrimFunc(symshen_4newpv), V4960)


let__13983 := tmp14609
_ = let__13983

tmp14610 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14610

tmp14611 := Call(__e, PrimFunc(symshen_4deref), V4959, V4960)


tmp14612 := Call(__e, PrimFunc(symshen_4app), tmp14611, MakeString(")\n"), symshen_4a)


tmp14613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\ntypechecking (fn "))
__typedS1, __typedOK1 := TypedString(tmp14612)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\ntypechecking (fn ")
__typedArg1 := tmp14612
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp14614 := Call(__e, PrimFunc(symstoutput))


tmp14615 := Call(__e, PrimFunc(sympr), tmp14613, tmp14614)


tmp14616 := Call(__e, PrimFunc(symis), let__13983, tmp14615, V4960, V4961, V4962, V4963)


__e.TailApply(PrimFunc(symshen_4gc), V4960, tmp14616)
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
__e.Return(let__13980)
return
}


}, 6)

tmp14623 := Call(__e, ns2_1set, symshen_4signal_1def, tmp14600)


_ = tmp14623

tmp14624 := MakeNative(func(__e *ControlFlow) {
V4968 := __e.Get(1)
_ = V4968
tmp14625 := Call(__e, PrimFunc(symshen_4curry_1type), V4968)


__e.TailApply(PrimFunc(symshen_4demodulate), tmp14625)
return


}, 1)

tmp14626 := Call(__e, ns2_1set, symshen_4rectify_1type, tmp14624)


_ = tmp14626

tmp14627 := MakeNative(func(__e *ControlFlow) {
V4969 := __e.Get(1)
_ = V4969
tmp14628 := MakeNative(func(__e *ControlFlow) {
tmp14629 := MakeNative(func(__e *ControlFlow) {
Z4971 := __e.Get(1)
_ = Z4971
__e.TailApply(PrimFunc(symshen_4demod), Z4971)
return
}, 1)

tmp14630 := Call(__e, PrimFunc(symshen_4walk), tmp14629, V4969)


let__13984 := tmp14630
_ = let__13984

tmp14632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13984, V4969)
}
__typedArg0 := let__13984
__typedArg1 := V4969
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14632 {
__e.Return(V4969)
return
} else {
__e.TailApply(PrimFunc(symshen_4demodulate), let__13984)
return
}


}, 0)

tmp14633 := MakeNative(func(__e *ControlFlow) {
Z4972 := __e.Get(1)
_ = Z4972
__e.Return(V4969)
return
}, 1)

__e.TailApply(try_1catch, tmp14628, tmp14633)
return


}, 1)

tmp14634 := Call(__e, ns2_1set, symshen_4demodulate, tmp14627)


_ = tmp14634

tmp14635 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V4973 := __self1
_ = V4973
tmp14759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14732 Obj

if True == tmp14759 {
tmp14757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14757)
}
__typedArg0 := tmp14757
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14734 Obj

if True == tmp14758 {
tmp14754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14754)
}
__typedArg0 := tmp14754
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1_1_6, tmp14755)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp14755
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14736 Obj

if True == tmp14756 {
tmp14751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14751)
}
__typedArg0 := tmp14751
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14752)
}
__typedArg0 := tmp14752
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14738 Obj

if True == tmp14753 {
tmp14747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14747)
}
__typedArg0 := tmp14747
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14748)
}
__typedArg0 := tmp14748
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14749)
}
__typedArg0 := tmp14749
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14740 Obj

if True == tmp14750 {
tmp14742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14742)
}
__typedArg0 := tmp14742
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14743)
}
__typedArg0 := tmp14743
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14744)
}
__typedArg0 := tmp14744
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1_1_6, tmp14745)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp14745
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14741 Obj

if True == tmp14746 {
ifres14741 = True


} else {
ifres14741 = False


}

ifres14740 = ifres14741


} else {
ifres14740 = False


}

var ifres14739 Obj

if True == ifres14740 {
ifres14739 = True


} else {
ifres14739 = False


}

ifres14738 = ifres14739


} else {
ifres14738 = False


}

var ifres14737 Obj

if True == ifres14738 {
ifres14737 = True


} else {
ifres14737 = False


}

ifres14736 = ifres14737


} else {
ifres14736 = False


}

var ifres14735 Obj

if True == ifres14736 {
ifres14735 = True


} else {
ifres14735 = False


}

ifres14734 = ifres14735


} else {
ifres14734 = False


}

var ifres14733 Obj

if True == ifres14734 {
ifres14733 = True


} else {
ifres14733 = False


}

ifres14732 = ifres14733


} else {
ifres14732 = False


}

if True == ifres14732 {
tmp14636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14637)
}
__typedArg0 := tmp14637
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14638, Nil)
}
__typedArg0 := tmp14638
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp14639)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp14639
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14636, tmp14640)
}
__typedArg0 := tmp14636
__typedArg1 := tmp14640
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4curry_1type) == __self {
__self1 = tmp14641
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4curry_1type), tmp14641)
return


} else {
tmp14730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14690 Obj

if True == tmp14730 {
tmp14728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14728)
}
__typedArg0 := tmp14728
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14692 Obj

if True == tmp14729 {
tmp14725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14725)
}
__typedArg0 := tmp14725
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp14726)
}
__typedArg0 := symlist
__typedArg1 := tmp14726
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14694 Obj

if True == tmp14727 {
tmp14722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14722)
}
__typedArg0 := tmp14722
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14723)
}
__typedArg0 := tmp14723
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14696 Obj

if True == tmp14724 {
tmp14718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14718)
}
__typedArg0 := tmp14718
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14719)
}
__typedArg0 := tmp14719
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14720)
}
__typedArg0 := Nil
__typedArg1 := tmp14720
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14698 Obj

if True == tmp14721 {
tmp14716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14716)
}
__typedArg0 := tmp14716
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14700 Obj

if True == tmp14717 {
tmp14713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14713)
}
__typedArg0 := tmp14713
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_a_a_6, tmp14714)
}
__typedArg0 := sym_a_a_6
__typedArg1 := tmp14714
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14702 Obj

if True == tmp14715 {
tmp14710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14710)
}
__typedArg0 := tmp14710
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14711)
}
__typedArg0 := tmp14711
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14704 Obj

if True == tmp14712 {
tmp14706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14706)
}
__typedArg0 := tmp14706
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14707)
}
__typedArg0 := tmp14707
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14708)
}
__typedArg0 := Nil
__typedArg1 := tmp14708
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14705 Obj

if True == tmp14709 {
ifres14705 = True


} else {
ifres14705 = False


}

ifres14704 = ifres14705


} else {
ifres14704 = False


}

var ifres14703 Obj

if True == ifres14704 {
ifres14703 = True


} else {
ifres14703 = False


}

ifres14702 = ifres14703


} else {
ifres14702 = False


}

var ifres14701 Obj

if True == ifres14702 {
ifres14701 = True


} else {
ifres14701 = False


}

ifres14700 = ifres14701


} else {
ifres14700 = False


}

var ifres14699 Obj

if True == ifres14700 {
ifres14699 = True


} else {
ifres14699 = False


}

ifres14698 = ifres14699


} else {
ifres14698 = False


}

var ifres14697 Obj

if True == ifres14698 {
ifres14697 = True


} else {
ifres14697 = False


}

ifres14696 = ifres14697


} else {
ifres14696 = False


}

var ifres14695 Obj

if True == ifres14696 {
ifres14695 = True


} else {
ifres14695 = False


}

ifres14694 = ifres14695


} else {
ifres14694 = False


}

var ifres14693 Obj

if True == ifres14694 {
ifres14693 = True


} else {
ifres14693 = False


}

ifres14692 = ifres14693


} else {
ifres14692 = False


}

var ifres14691 Obj

if True == ifres14692 {
ifres14691 = True


} else {
ifres14691 = False


}

ifres14690 = ifres14691


} else {
ifres14690 = False


}

if True == ifres14690 {
tmp14642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4973)
}
__typedArg0 := V4973
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
return PrimCons(tmp14643, tmp14645)
}
__typedArg0 := tmp14643
__typedArg1 := tmp14645
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp14646)
}
__typedArg0 := symstr
__typedArg1 := tmp14646
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14647, Nil)
}
__typedArg0 := tmp14647
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp14648)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp14648
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14642, tmp14649)
}
__typedArg0 := tmp14642
__typedArg1 := tmp14649
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4curry_1type) == __self {
__self1 = tmp14650
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4curry_1type), tmp14650)
return


} else {
tmp14688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14661 Obj

if True == tmp14688 {
tmp14686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14686)
}
__typedArg0 := tmp14686
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14663 Obj

if True == tmp14687 {
tmp14683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14683)
}
__typedArg0 := tmp14683
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_d, tmp14684)
}
__typedArg0 := sym_d
__typedArg1 := tmp14684
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14665 Obj

if True == tmp14685 {
tmp14680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14680)
}
__typedArg0 := tmp14680
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14681)
}
__typedArg0 := tmp14681
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14667 Obj

if True == tmp14682 {
tmp14676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14676)
}
__typedArg0 := tmp14676
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14677)
}
__typedArg0 := tmp14677
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14678)
}
__typedArg0 := tmp14678
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14669 Obj

if True == tmp14679 {
tmp14671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14671)
}
__typedArg0 := tmp14671
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14672)
}
__typedArg0 := tmp14672
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14673)
}
__typedArg0 := tmp14673
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_d, tmp14674)
}
__typedArg0 := sym_d
__typedArg1 := tmp14674
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14670 Obj

if True == tmp14675 {
ifres14670 = True


} else {
ifres14670 = False


}

ifres14669 = ifres14670


} else {
ifres14669 = False


}

var ifres14668 Obj

if True == ifres14669 {
ifres14668 = True


} else {
ifres14668 = False


}

ifres14667 = ifres14668


} else {
ifres14667 = False


}

var ifres14666 Obj

if True == ifres14667 {
ifres14666 = True


} else {
ifres14666 = False


}

ifres14665 = ifres14666


} else {
ifres14665 = False


}

var ifres14664 Obj

if True == ifres14665 {
ifres14664 = True


} else {
ifres14664 = False


}

ifres14663 = ifres14664


} else {
ifres14663 = False


}

var ifres14662 Obj

if True == ifres14663 {
ifres14662 = True


} else {
ifres14662 = False


}

ifres14661 = ifres14662


} else {
ifres14661 = False


}

if True == ifres14661 {
tmp14651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14652)
}
__typedArg0 := tmp14652
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14653, Nil)
}
__typedArg0 := tmp14653
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp14654)
}
__typedArg0 := sym_d
__typedArg1 := tmp14654
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14651, tmp14655)
}
__typedArg0 := tmp14651
__typedArg1 := tmp14655
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4curry_1type) == __self {
__self1 = tmp14656
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4curry_1type), tmp14656)
return


} else {
tmp14659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4973)
}
__typedArg0 := V4973
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14659 {
tmp14657 := MakeNative(func(__e *ControlFlow) {
Z4974 := __e.Get(1)
_ = Z4974
__e.TailApply(PrimFunc(symshen_4curry_1type), Z4974)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp14657, V4973)
return


} else {
__e.Return(V4973)
return
}


}


}


}


}, 1)

tmp14760 := Call(__e, ns2_1set, symshen_4curry_1type, tmp14635)


_ = tmp14760

tmp14761 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V4975 := __self1
_ = V4975
tmp14850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14842 Obj

if True == tmp14850 {
tmp14848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp14848)
}
__typedArg0 := symdefine
__typedArg1 := tmp14848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14844 Obj

if True == tmp14849 {
tmp14846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14846)
}
__typedArg0 := tmp14846
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14845 Obj

if True == tmp14847 {
ifres14845 = True


} else {
ifres14845 = False


}

ifres14844 = ifres14845


} else {
ifres14844 = False


}

var ifres14843 Obj

if True == ifres14844 {
ifres14843 = True


} else {
ifres14843 = False


}

ifres14842 = ifres14843


} else {
ifres14842 = False


}

if True == ifres14842 {
__e.Return(V4975)
return
} else {
tmp14840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14821 Obj

if True == tmp14840 {
tmp14838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtype, tmp14838)
}
__typedArg0 := symtype
__typedArg1 := tmp14838
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14823 Obj

if True == tmp14839 {
tmp14836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14836)
}
__typedArg0 := tmp14836
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14825 Obj

if True == tmp14837 {
tmp14833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14833)
}
__typedArg0 := tmp14833
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14834)
}
__typedArg0 := tmp14834
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14827 Obj

if True == tmp14835 {
tmp14829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14829)
}
__typedArg0 := tmp14829
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14830)
}
__typedArg0 := tmp14830
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14831)
}
__typedArg0 := Nil
__typedArg1 := tmp14831
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14828 Obj

if True == tmp14832 {
ifres14828 = True


} else {
ifres14828 = False


}

ifres14827 = ifres14828


} else {
ifres14827 = False


}

var ifres14826 Obj

if True == ifres14827 {
ifres14826 = True


} else {
ifres14826 = False


}

ifres14825 = ifres14826


} else {
ifres14825 = False


}

var ifres14824 Obj

if True == ifres14825 {
ifres14824 = True


} else {
ifres14824 = False


}

ifres14823 = ifres14824


} else {
ifres14823 = False


}

var ifres14822 Obj

if True == ifres14823 {
ifres14822 = True


} else {
ifres14822 = False


}

ifres14821 = ifres14822


} else {
ifres14821 = False


}

if True == ifres14821 {
tmp14762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14762)
}
__typedArg0 := tmp14762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14764 := Call(__e, PrimFunc(symshen_4curry), tmp14763)


tmp14765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14765)
}
__typedArg0 := tmp14765
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14764, tmp14766)
}
__typedArg0 := tmp14764
__typedArg1 := tmp14766
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp14767)
}
__typedArg0 := symtype
__typedArg1 := tmp14767
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp14819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14815 Obj

if True == tmp14819 {
tmp14817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14818 := Call(__e, PrimFunc(symshen_4special_2), tmp14817)


var ifres14816 Obj

if True == tmp14818 {
ifres14816 = True


} else {
ifres14816 = False


}

ifres14815 = ifres14816


} else {
ifres14815 = False


}

if True == ifres14815 {
tmp14768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14769 := MakeNative(func(__e *ControlFlow) {
Z4976 := __e.Get(1)
_ = Z4976
__e.TailApply(PrimFunc(symshen_4curry), Z4976)
return
}, 1)

tmp14770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14771 := Call(__e, PrimFunc(symmap), tmp14769, tmp14770)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14768, tmp14771)
}
__typedArg0 := tmp14768
__typedArg1 := tmp14771
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp14813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14809 Obj

if True == tmp14813 {
tmp14811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14812 := Call(__e, PrimFunc(symshen_4extraspecial_2), tmp14811)


var ifres14810 Obj

if True == tmp14812 {
ifres14810 = True


} else {
ifres14810 = False


}

ifres14809 = ifres14810


} else {
ifres14809 = False


}

if True == ifres14809 {
__e.Return(V4975)
return
} else {
tmp14807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14798 Obj

if True == tmp14807 {
tmp14805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14805)
}
__typedArg0 := tmp14805
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14800 Obj

if True == tmp14806 {
tmp14802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14802)
}
__typedArg0 := tmp14802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14803)
}
__typedArg0 := tmp14803
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14801 Obj

if True == tmp14804 {
ifres14801 = True


} else {
ifres14801 = False


}

ifres14800 = ifres14801


} else {
ifres14800 = False


}

var ifres14799 Obj

if True == ifres14800 {
ifres14799 = True


} else {
ifres14799 = False


}

ifres14798 = ifres14799


} else {
ifres14798 = False


}

if True == ifres14798 {
tmp14772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14773)
}
__typedArg0 := tmp14773
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14774, Nil)
}
__typedArg0 := tmp14774
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14772, tmp14775)
}
__typedArg0 := tmp14772
__typedArg1 := tmp14775
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14777)
}
__typedArg0 := tmp14777
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14776, tmp14778)
}
__typedArg0 := tmp14776
__typedArg1 := tmp14778
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4curry) == __self {
__self1 = tmp14779
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4curry), tmp14779)
return


} else {
tmp14796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14787 Obj

if True == tmp14796 {
tmp14794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14794)
}
__typedArg0 := tmp14794
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14789 Obj

if True == tmp14795 {
tmp14791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14791)
}
__typedArg0 := tmp14791
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14792)
}
__typedArg0 := Nil
__typedArg1 := tmp14792
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14790 Obj

if True == tmp14793 {
ifres14790 = True


} else {
ifres14790 = False


}

ifres14789 = ifres14790


} else {
ifres14789 = False


}

var ifres14788 Obj

if True == ifres14789 {
ifres14788 = True


} else {
ifres14788 = False


}

ifres14787 = ifres14788


} else {
ifres14787 = False


}

if True == ifres14787 {
tmp14780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14781 := Call(__e, PrimFunc(symshen_4curry), tmp14780)


tmp14782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4975)
}
__typedArg0 := V4975
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14782)
}
__typedArg0 := tmp14782
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14784 := Call(__e, PrimFunc(symshen_4curry), tmp14783)


tmp14785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14784, Nil)
}
__typedArg0 := tmp14784
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14781, tmp14785)
}
__typedArg0 := tmp14781
__typedArg1 := tmp14785
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V4975)
return
}


}


}


}


}


}


}, 1)

tmp14851 := Call(__e, ns2_1set, symshen_4curry, tmp14761)


_ = tmp14851

tmp14852 := MakeNative(func(__e *ControlFlow) {
V4977 := __e.Get(1)
_ = V4977
tmp14853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspecial_d)
}
__typedArg0 := symshen_4_dspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symelement_2), V4977, tmp14853)
return


}, 1)

tmp14854 := Call(__e, ns2_1set, symshen_4special_2, tmp14852)


_ = tmp14854

tmp14855 := MakeNative(func(__e *ControlFlow) {
V4978 := __e.Get(1)
_ = V4978
tmp14856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dextraspecial_d)
}
__typedArg0 := symshen_4_dextraspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symelement_2), V4978, tmp14856)
return


}, 1)

tmp14857 := Call(__e, ns2_1set, symshen_4extraspecial_2, tmp14855)


_ = tmp14857

tmp14858 := MakeNative(func(__e *ControlFlow) {
V4979 := __e.Get(1)
_ = V4979
V4980 := __e.Get(2)
_ = V4980
V4981 := __e.Get(3)
_ = V4981
V4982 := __e.Get(4)
_ = V4982
V4983 := __e.Get(5)
_ = V4983
V4984 := __e.Get(6)
_ = V4984
let__13985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4983)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4983
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__13985

tmp14864 := Call(__e, PrimFunc(symshen_4unlocked_2), V4982)


var ifres14860 Obj

if True == tmp14864 {
tmp14861 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14861

tmp14862 := Call(__e, PrimFunc(symshen_4maxinfexceeded_2))


tmp14863 := Call(__e, PrimFunc(symwhen), tmp14862, V4981, V4982, let__13985, V4984)


ifres14860 = tmp14863


} else {
ifres14860 = False


}

let__13986 := ifres14860
_ = let__13986

tmp14912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13986, False)
}
__typedArg0 := let__13986
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14912 {
tmp14893 := Call(__e, PrimFunc(symshen_4unlocked_2), V4982)


var ifres14865 Obj

if True == tmp14893 {
tmp14866 := Call(__e, PrimFunc(symshen_4lazyderef), V4979, V4981)


let__13988 := tmp14866
_ = let__13988

tmp14892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__13988)
}
__typedArg0 := let__13988
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14867 Obj

if True == tmp14892 {
tmp14868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__13988)
}
__typedArg0 := let__13988
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__13989 := tmp14868
_ = let__13989

tmp14869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__13988)
}
__typedArg0 := let__13988
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14870 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14869, V4981)


let__13990 := tmp14870
_ = let__13990

tmp14891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__13990)
}
__typedArg0 := let__13990
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14871 Obj

if True == tmp14891 {
tmp14872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__13990)
}
__typedArg0 := let__13990
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__13991 := tmp14872
_ = let__13991

tmp14873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__13990)
}
__typedArg0 := let__13990
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14874 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14873, V4981)


let__13992 := tmp14874
_ = let__13992

tmp14890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__13992)
}
__typedArg0 := let__13992
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14875 Obj

if True == tmp14890 {
tmp14876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__13992)
}
__typedArg0 := let__13992
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__13993 := tmp14876
_ = let__13993

tmp14877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__13992)
}
__typedArg0 := let__13992
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14878 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14877, V4981)


let__13994 := tmp14878
_ = let__13994

tmp14889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13994, Nil)
}
__typedArg0 := let__13994
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14879 Obj

if True == tmp14889 {
tmp14880 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14880

tmp14881 := Call(__e, PrimFunc(symshen_4deref), let__13991, V4981)


tmp14882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp14883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp14881, tmp14882)
}
__typedArg0 := tmp14881
__typedArg1 := tmp14882
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp14884 := MakeNative(func(__e *ControlFlow) {
tmp14885 := Call(__e, PrimFunc(symshen_4type_1theory_1enabled_2))


tmp14886 := MakeNative(func(__e *ControlFlow) {
tmp14887 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__13989, let__13993, V4980, V4981, V4982, let__13985, V4984)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4981, V4982, let__13985, tmp14887)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14885, V4981, V4982, let__13985, tmp14886)
return


}, 0)

tmp14888 := Call(__e, PrimFunc(symwhen), tmp14883, V4981, V4982, let__13985, tmp14884)


ifres14879 = tmp14888


} else {
ifres14879 = False


}

ifres14875 = ifres14879


} else {
ifres14875 = False


}

ifres14871 = ifres14875


} else {
ifres14871 = False


}

ifres14867 = ifres14871


} else {
ifres14867 = False


}

ifres14865 = ifres14867


} else {
ifres14865 = False


}

let__13987 := ifres14865
_ = let__13987

tmp14910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13987, False)
}
__typedArg0 := let__13987
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14910 {
tmp14899 := Call(__e, PrimFunc(symshen_4unlocked_2), V4982)


var ifres14894 Obj

if True == tmp14899 {
tmp14895 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14895

tmp14896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspy_d)
}
__typedArg0 := symshen_4_dspy_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp14897 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4show), V4979, V4980, V4981, V4982, let__13985, V4984)
return
}, 0)

tmp14898 := Call(__e, PrimFunc(symwhen), tmp14896, V4981, V4982, let__13985, tmp14897)


ifres14894 = tmp14898


} else {
ifres14894 = False


}

let__13995 := ifres14894
_ = let__13995

tmp14908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13995, False)
}
__typedArg0 := let__13995
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14908 {
tmp14904 := Call(__e, PrimFunc(symshen_4unlocked_2), V4982)


var ifres14900 Obj

if True == tmp14904 {
tmp14901 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14901

tmp14902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp14903 := Call(__e, PrimFunc(symshen_4search_1user_1datatypes), V4979, V4980, tmp14902, V4981, V4982, let__13985, V4984)


ifres14900 = tmp14903


} else {
ifres14900 = False


}

let__13996 := ifres14900
_ = let__13996

tmp14906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13996, False)
}
__typedArg0 := let__13996
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14906 {
__e.TailApply(PrimFunc(symshen_4unlock), V4982, let__13985)
return
} else {
__e.Return(let__13996)
return
}


} else {
__e.Return(let__13995)
return
}


} else {
__e.Return(let__13987)
return
}


} else {
__e.Return(let__13986)
return
}


}, 6)

tmp14913 := Call(__e, ns2_1set, symshen_4system_1S, tmp14858)


_ = tmp14913

tmp14914 := MakeNative(func(__e *ControlFlow) {
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
tmp14915 := Call(__e, PrimFunc(symshen_4line))


_ = tmp14915

tmp14916 := Call(__e, PrimFunc(symshen_4deref), V5003, V5005)


tmp14917 := Call(__e, PrimFunc(symshen_4show_1p), tmp14916)


_ = tmp14917

tmp14918 := Call(__e, PrimFunc(symnl), MakeInteger(2))


_ = tmp14918

tmp14919 := Call(__e, PrimFunc(symshen_4deref), V5004, V5005)


tmp14920 := Call(__e, PrimFunc(symshen_4show_1assumptions), tmp14919, MakeInteger(1))


_ = tmp14920

tmp14921 := Call(__e, PrimFunc(symshen_4pause_1for_1user))


_ = tmp14921

__e.Return(False)
return


}, 6)

tmp14922 := Call(__e, ns2_1set, symshen_4show, tmp14914)


_ = tmp14922

tmp14923 := MakeNative(func(__e *ControlFlow) {
tmp14924 := Call(__e, PrimFunc(syminferences))


let__13997 := tmp14924
_ = let__13997

tmp14926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(1), let__13997)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := let__13997
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14925 Obj

if True == tmp14926 {
ifres14925 = MakeString("")


} else {
ifres14925 = MakeString("s")


}

tmp14927 := Call(__e, PrimFunc(symshen_4app), ifres14925, MakeString(" \n?- "), symshen_4a)


tmp14929 := Call(__e, PrimFunc(symshen_4app), let__13997, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" inference"))
__typedS1, __typedOK1 := TypedString(tmp14927)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" inference")
__typedArg1 := tmp14927
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp14930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("____________________________________________________________ "))
__typedS1, __typedOK1 := TypedString(tmp14929)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("____________________________________________________________ ")
__typedArg1 := tmp14929
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp14931 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp14930, tmp14931)
return


}, 0)

tmp14932 := Call(__e, ns2_1set, symshen_4line, tmp14923)


_ = tmp14932

tmp14933 := MakeNative(func(__e *ControlFlow) {
V5010 := __e.Get(1)
_ = V5010
tmp14965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5010)
}
__typedArg0 := V5010
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14944 Obj

if True == tmp14965 {
tmp14963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5010)
}
__typedArg0 := V5010
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14963)
}
__typedArg0 := tmp14963
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14946 Obj

if True == tmp14964 {
tmp14960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5010)
}
__typedArg0 := V5010
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14960)
}
__typedArg0 := tmp14960
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14961)
}
__typedArg0 := tmp14961
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14948 Obj

if True == tmp14962 {
tmp14956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5010)
}
__typedArg0 := V5010
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14956)
}
__typedArg0 := tmp14956
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14957)
}
__typedArg0 := tmp14957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14958)
}
__typedArg0 := Nil
__typedArg1 := tmp14958
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14950 Obj

if True == tmp14959 {
tmp14952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5010)
}
__typedArg0 := V5010
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14952)
}
__typedArg0 := tmp14952
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp14955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp14953, tmp14954)
}
__typedArg0 := tmp14953
__typedArg1 := tmp14954
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14951 Obj

if True == tmp14955 {
ifres14951 = True


} else {
ifres14951 = False


}

ifres14950 = ifres14951


} else {
ifres14950 = False


}

var ifres14949 Obj

if True == ifres14950 {
ifres14949 = True


} else {
ifres14949 = False


}

ifres14948 = ifres14949


} else {
ifres14948 = False


}

var ifres14947 Obj

if True == ifres14948 {
ifres14947 = True


} else {
ifres14947 = False


}

ifres14946 = ifres14947


} else {
ifres14946 = False


}

var ifres14945 Obj

if True == ifres14946 {
ifres14945 = True


} else {
ifres14945 = False


}

ifres14944 = ifres14945


} else {
ifres14944 = False


}

if True == ifres14944 {
tmp14934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5010)
}
__typedArg0 := V5010
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14935 := Call(__e, PrimFunc(symshen_4prterm), tmp14934)


_ = tmp14935

tmp14936 := Call(__e, PrimFunc(symstoutput))


tmp14937 := Call(__e, PrimFunc(sympr), MakeString(" : "), tmp14936)


_ = tmp14937

tmp14938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5010)
}
__typedArg0 := V5010
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14938)
}
__typedArg0 := tmp14938
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14939)
}
__typedArg0 := tmp14939
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14941 := Call(__e, PrimFunc(symshen_4app), tmp14940, MakeString(""), symshen_4r)


tmp14942 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp14941, tmp14942)
return


} else {
__e.TailApply(PrimFunc(symshen_4prterm), V5010)
return
}


}, 1)

tmp14966 := Call(__e, ns2_1set, symshen_4show_1p, tmp14933)


_ = tmp14966

tmp14967 := MakeNative(func(__e *ControlFlow) {
V5011 := __e.Get(1)
_ = V5011
tmp15010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5011)
}
__typedArg0 := V5011
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14991 Obj

if True == tmp15010 {
tmp15008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5011)
}
__typedArg0 := V5011
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp15008)
}
__typedArg0 := symcons
__typedArg1 := tmp15008
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14993 Obj

if True == tmp15009 {
tmp15006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5011)
}
__typedArg0 := V5011
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15006)
}
__typedArg0 := tmp15006
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14995 Obj

if True == tmp15007 {
tmp15003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5011)
}
__typedArg0 := V5011
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15003)
}
__typedArg0 := tmp15003
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15004)
}
__typedArg0 := tmp15004
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14997 Obj

if True == tmp15005 {
tmp14999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5011)
}
__typedArg0 := V5011
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14999)
}
__typedArg0 := tmp14999
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15000)
}
__typedArg0 := tmp15000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp15001)
}
__typedArg0 := Nil
__typedArg1 := tmp15001
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14998 Obj

if True == tmp15002 {
ifres14998 = True


} else {
ifres14998 = False


}

ifres14997 = ifres14998


} else {
ifres14997 = False


}

var ifres14996 Obj

if True == ifres14997 {
ifres14996 = True


} else {
ifres14996 = False


}

ifres14995 = ifres14996


} else {
ifres14995 = False


}

var ifres14994 Obj

if True == ifres14995 {
ifres14994 = True


} else {
ifres14994 = False


}

ifres14993 = ifres14994


} else {
ifres14993 = False


}

var ifres14992 Obj

if True == ifres14993 {
ifres14992 = True


} else {
ifres14992 = False


}

ifres14991 = ifres14992


} else {
ifres14991 = False


}

if True == ifres14991 {
tmp14968 := Call(__e, PrimFunc(symstoutput))


tmp14969 := Call(__e, PrimFunc(sympr), MakeString("["), tmp14968)


_ = tmp14969

tmp14970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5011)
}
__typedArg0 := V5011
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14970)
}
__typedArg0 := tmp14970
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14972 := Call(__e, PrimFunc(symshen_4prterm), tmp14971)


_ = tmp14972

tmp14973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5011)
}
__typedArg0 := V5011
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14973)
}
__typedArg0 := tmp14973
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14974)
}
__typedArg0 := tmp14974
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14976 := Call(__e, PrimFunc(symshen_4prtl), tmp14975)


_ = tmp14976

tmp14977 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString("]"), tmp14977)
return


} else {
tmp14989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5011)
}
__typedArg0 := V5011
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14989 {
tmp14978 := Call(__e, PrimFunc(symstoutput))


tmp14979 := Call(__e, PrimFunc(sympr), MakeString("("), tmp14978)


_ = tmp14979

tmp14980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5011)
}
__typedArg0 := V5011
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14981 := Call(__e, PrimFunc(symshen_4prterm), tmp14980)


_ = tmp14981

tmp14982 := MakeNative(func(__e *ControlFlow) {
Z5012 := __e.Get(1)
_ = Z5012
tmp14983 := Call(__e, PrimFunc(symstoutput))


tmp14984 := Call(__e, PrimFunc(sympr), MakeString(" "), tmp14983)


_ = tmp14984

__e.TailApply(PrimFunc(symshen_4prterm), Z5012)
return


}, 1)

tmp14985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5011)
}
__typedArg0 := V5011
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14986 := Call(__e, PrimFunc(symmap), tmp14982, tmp14985)


_ = tmp14986

tmp14987 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString(")"), tmp14987)
return


} else {
__e.TailApply(PrimFunc(symprint), V5011)
return
}


}


}, 1)

tmp15011 := Call(__e, ns2_1set, symshen_4prterm, tmp14967)


_ = tmp15011

tmp15012 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V5013 := __self1
_ = V5013
tmp15045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5013)
}
__typedArg0 := Nil
__typedArg1 := V5013
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15045 {
__e.Return(MakeString(""))
return
} else {
tmp15043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5013)
}
__typedArg0 := V5013
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15024 Obj

if True == tmp15043 {
tmp15041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5013)
}
__typedArg0 := V5013
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp15041)
}
__typedArg0 := symcons
__typedArg1 := tmp15041
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15026 Obj

if True == tmp15042 {
tmp15039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5013)
}
__typedArg0 := V5013
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15039)
}
__typedArg0 := tmp15039
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15028 Obj

if True == tmp15040 {
tmp15036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5013)
}
__typedArg0 := V5013
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15036)
}
__typedArg0 := tmp15036
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15037)
}
__typedArg0 := tmp15037
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15030 Obj

if True == tmp15038 {
tmp15032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5013)
}
__typedArg0 := V5013
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15032)
}
__typedArg0 := tmp15032
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15033)
}
__typedArg0 := tmp15033
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp15034)
}
__typedArg0 := Nil
__typedArg1 := tmp15034
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15031 Obj

if True == tmp15035 {
ifres15031 = True


} else {
ifres15031 = False


}

ifres15030 = ifres15031


} else {
ifres15030 = False


}

var ifres15029 Obj

if True == ifres15030 {
ifres15029 = True


} else {
ifres15029 = False


}

ifres15028 = ifres15029


} else {
ifres15028 = False


}

var ifres15027 Obj

if True == ifres15028 {
ifres15027 = True


} else {
ifres15027 = False


}

ifres15026 = ifres15027


} else {
ifres15026 = False


}

var ifres15025 Obj

if True == ifres15026 {
ifres15025 = True


} else {
ifres15025 = False


}

ifres15024 = ifres15025


} else {
ifres15024 = False


}

if True == ifres15024 {
tmp15013 := Call(__e, PrimFunc(symstoutput))


tmp15014 := Call(__e, PrimFunc(sympr), MakeString(" "), tmp15013)


_ = tmp15014

tmp15015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5013)
}
__typedArg0 := V5013
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15015)
}
__typedArg0 := tmp15015
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15017 := Call(__e, PrimFunc(symshen_4prterm), tmp15016)


_ = tmp15017

tmp15018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5013)
}
__typedArg0 := V5013
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15018)
}
__typedArg0 := tmp15018
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15019)
}
__typedArg0 := tmp15019
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

if PrimFunc(symshen_4prtl) == __self {
__self1 = tmp15020
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4prtl), tmp15020)
return


} else {
tmp15021 := Call(__e, PrimFunc(symstoutput))


tmp15022 := Call(__e, PrimFunc(sympr), MakeString(" | "), tmp15021)


_ = tmp15022

__e.TailApply(PrimFunc(symshen_4prterm), V5013)
return


}


}


}, 1)

tmp15046 := Call(__e, ns2_1set, symshen_4prtl, tmp15012)


_ = tmp15046

tmp15047 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V5020 := __self1
_ = V5020
V5021 := __self2
_ = V5021
tmp15060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5020)
}
__typedArg0 := Nil
__typedArg1 := V5020
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15060 {
tmp15048 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString("\n> "), tmp15048)
return


} else {
tmp15058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5020)
}
__typedArg0 := V5020
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15058 {
tmp15049 := Call(__e, PrimFunc(symshen_4app), V5021, MakeString(". "), symshen_4a)


tmp15050 := Call(__e, PrimFunc(symstoutput))


tmp15051 := Call(__e, PrimFunc(sympr), tmp15049, tmp15050)


_ = tmp15051

tmp15052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5020)
}
__typedArg0 := V5020
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15053 := Call(__e, PrimFunc(symshen_4show_1p), tmp15052)


_ = tmp15053

tmp15054 := Call(__e, PrimFunc(symnl), MakeInteger(1))


_ = tmp15054

tmp15055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5020)
}
__typedArg0 := V5020
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4show_1assumptions) == __self {
__self1, __self2 = tmp15055, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5021)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5021
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4show_1assumptions), tmp15055, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5021)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5021
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

tmp15061 := Call(__e, ns2_1set, symshen_4show_1assumptions, tmp15047)


_ = tmp15061

tmp15062 := MakeNative(func(__e *ControlFlow) {
tmp15063 := Call(__e, PrimFunc(symstinput))


tmp15064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(tmp15063)
}
__typedArg0 := tmp15063
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

let__13998 := tmp15064
_ = let__13998

tmp15066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13998, MakeInteger(94))
}
__typedArg0 := let__13998
__typedArg1 := MakeInteger(94)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15066 {
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

tmp15067 := Call(__e, ns2_1set, symshen_4pause_1for_1user, tmp15062)


_ = tmp15067

tmp15068 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dshen_1type_1theory_1enabled_2_d)
}
__typedArg0 := symshen_4_dshen_1type_1theory_1enabled_2_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp15069 := Call(__e, ns2_1set, symshen_4type_1theory_1enabled_2, tmp15068)


_ = tmp15069

tmp15070 := MakeNative(func(__e *ControlFlow) {
tmp15072 := Call(__e, PrimFunc(syminferences))


tmp15073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dmaxinferences_d)
}
__typedArg0 := symshen_4_dmaxinferences_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(tmp15072)
__typedN1, __typedOK1 := TypedFloat64(tmp15073)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := tmp15072
__typedArg1 := tmp15073
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

tmp15075 := Call(__e, ns2_1set, symshen_4maxinfexceeded_2, tmp15070)


_ = tmp15075

tmp15076 := MakeNative(func(__e *ControlFlow) {
V5023 := __e.Get(1)
_ = V5023
V5024 := __e.Get(2)
_ = V5024
V5025 := __e.Get(3)
_ = V5025
V5026 := __e.Get(4)
_ = V5026
V5027 := __e.Get(5)
_ = V5027
V5028 := __e.Get(6)
_ = V5028
V5029 := __e.Get(7)
_ = V5029
let__13999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5028)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5028
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__13999

tmp15087 := Call(__e, PrimFunc(symshen_4unlocked_2), V5027)


var ifres15078 Obj

if True == tmp15087 {
tmp15079 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15079

tmp15080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspy_d)
}
__typedArg0 := symshen_4_dspy_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp15081 := MakeNative(func(__e *ControlFlow) {
tmp15082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp15083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5024, Nil)
}
__typedArg0 := V5024
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15082, tmp15083)
}
__typedArg0 := tmp15082
__typedArg1 := tmp15083
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5023, tmp15084)
}
__typedArg0 := V5023
__typedArg1 := tmp15084
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4show), tmp15085, V5025, V5026, V5027, let__13999, V5029)
return


}, 0)

tmp15086 := Call(__e, PrimFunc(symwhen), tmp15080, V5026, V5027, let__13999, tmp15081)


ifres15078 = tmp15086


} else {
ifres15078 = False


}

let__14000 := ifres15078
_ = let__14000

tmp15826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14000, False)
}
__typedArg0 := let__14000
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15826 {
tmp15095 := Call(__e, PrimFunc(symshen_4unlocked_2), V5027)


var ifres15088 Obj

if True == tmp15095 {
tmp15089 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15089

tmp15090 := Call(__e, PrimFunc(symshen_4lazyderef), V5023, V5026)


tmp15091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15090)
}
__typedArg0 := tmp15090
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

tmp15092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp15091)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp15091
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

tmp15093 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4primitive), V5023, V5024, V5026, V5027, let__13999, V5029)
return
}, 0)

tmp15094 := Call(__e, PrimFunc(symwhen), tmp15092, V5026, V5027, let__13999, tmp15093)


ifres15088 = tmp15094


} else {
ifres15088 = False


}

let__14001 := ifres15088
_ = let__14001

tmp15824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14001, False)
}
__typedArg0 := let__14001
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15824 {
tmp15099 := Call(__e, PrimFunc(symshen_4unlocked_2), V5027)


var ifres15096 Obj

if True == tmp15099 {
tmp15097 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15097

tmp15098 := Call(__e, PrimFunc(symshen_4by_1hypothesis), V5023, V5024, V5025, V5026, V5027, let__13999, V5029)


ifres15096 = tmp15098


} else {
ifres15096 = False


}

let__14002 := ifres15096
_ = let__14002

tmp15822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14002, False)
}
__typedArg0 := let__14002
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15822 {
tmp15113 := Call(__e, PrimFunc(symshen_4unlocked_2), V5027)


var ifres15100 Obj

if True == tmp15113 {
tmp15101 := Call(__e, PrimFunc(symshen_4lazyderef), V5023, V5026)


let__14004 := tmp15101
_ = let__14004

tmp15112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14004)
}
__typedArg0 := let__14004
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15102 Obj

if True == tmp15112 {
tmp15103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14004)
}
__typedArg0 := let__14004
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14005 := tmp15103
_ = let__14005

tmp15104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14004)
}
__typedArg0 := let__14004
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15105 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15104, V5026)


let__14006 := tmp15105
_ = let__14006

tmp15111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14006, Nil)
}
__typedArg0 := let__14006
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15106 Obj

if True == tmp15111 {
tmp15107 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15107

tmp15108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5024, Nil)
}
__typedArg0 := V5024
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15108)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15108
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15110 := Call(__e, PrimFunc(symshen_4lookupsig), let__14005, tmp15109, V5026, V5027, let__13999, V5029)


ifres15106 = tmp15110


} else {
ifres15106 = False


}

ifres15102 = ifres15106


} else {
ifres15102 = False


}

ifres15100 = ifres15102


} else {
ifres15100 = False


}

let__14003 := ifres15100
_ = let__14003

tmp15820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14003, False)
}
__typedArg0 := let__14003
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15820 {
tmp15139 := Call(__e, PrimFunc(symshen_4unlocked_2), V5027)


var ifres15114 Obj

if True == tmp15139 {
tmp15115 := Call(__e, PrimFunc(symshen_4lazyderef), V5023, V5026)


let__14008 := tmp15115
_ = let__14008

tmp15138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14008)
}
__typedArg0 := let__14008
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15116 Obj

if True == tmp15138 {
tmp15117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14008)
}
__typedArg0 := let__14008
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15118 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15117, V5026)


let__14009 := tmp15118
_ = let__14009

tmp15137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14009, symfn)
}
__typedArg0 := let__14009
__typedArg1 := symfn
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15119 Obj

if True == tmp15137 {
tmp15120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14008)
}
__typedArg0 := let__14008
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15121 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15120, V5026)


let__14010 := tmp15121
_ = let__14010

tmp15136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14010)
}
__typedArg0 := let__14010
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15122 Obj

if True == tmp15136 {
tmp15123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14010)
}
__typedArg0 := let__14010
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14011 := tmp15123
_ = let__14011

tmp15124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14010)
}
__typedArg0 := let__14010
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15125 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15124, V5026)


let__14012 := tmp15125
_ = let__14012

tmp15135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14012, Nil)
}
__typedArg0 := let__14012
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15126 Obj

if True == tmp15135 {
tmp15127 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15127

tmp15128 := Call(__e, PrimFunc(symshen_4deref), let__14011, V5026)


tmp15129 := Call(__e, PrimFunc(symarity), tmp15128)


tmp15130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp15129, MakeInteger(0))
}
__typedArg0 := tmp15129
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp15131 := MakeNative(func(__e *ControlFlow) {
tmp15132 := MakeNative(func(__e *ControlFlow) {
tmp15133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14011, Nil)
}
__typedArg0 := let__14011
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp15133, V5024, V5025, V5026, V5027, let__13999, V5029)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5026, V5027, let__13999, tmp15132)
return


}, 0)

tmp15134 := Call(__e, PrimFunc(symwhen), tmp15130, V5026, V5027, let__13999, tmp15131)


ifres15126 = tmp15134


} else {
ifres15126 = False


}

ifres15122 = ifres15126


} else {
ifres15122 = False


}

ifres15119 = ifres15122


} else {
ifres15119 = False


}

ifres15116 = ifres15119


} else {
ifres15116 = False


}

ifres15114 = ifres15116


} else {
ifres15114 = False


}

let__14007 := ifres15114
_ = let__14007

tmp15818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14007, False)
}
__typedArg0 := let__14007
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15818 {
tmp15159 := Call(__e, PrimFunc(symshen_4unlocked_2), V5027)


var ifres15140 Obj

if True == tmp15159 {
tmp15141 := Call(__e, PrimFunc(symshen_4lazyderef), V5023, V5026)


let__14014 := tmp15141
_ = let__14014

tmp15158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14014)
}
__typedArg0 := let__14014
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15142 Obj

if True == tmp15158 {
tmp15143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14014)
}
__typedArg0 := let__14014
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15144 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15143, V5026)


let__14015 := tmp15144
_ = let__14015

tmp15157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14015, symfn)
}
__typedArg0 := let__14015
__typedArg1 := symfn
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15145 Obj

if True == tmp15157 {
tmp15146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14014)
}
__typedArg0 := let__14014
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15147 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15146, V5026)


let__14016 := tmp15147
_ = let__14016

tmp15156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14016)
}
__typedArg0 := let__14016
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15148 Obj

if True == tmp15156 {
tmp15149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14016)
}
__typedArg0 := let__14016
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14017 := tmp15149
_ = let__14017

tmp15150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14016)
}
__typedArg0 := let__14016
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15151 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15150, V5026)


let__14018 := tmp15151
_ = let__14018

tmp15155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14018, Nil)
}
__typedArg0 := let__14018
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15152 Obj

if True == tmp15155 {
tmp15153 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15153

tmp15154 := Call(__e, PrimFunc(symshen_4lookupsig), let__14017, V5024, V5026, V5027, let__13999, V5029)


ifres15152 = tmp15154


} else {
ifres15152 = False


}

ifres15148 = ifres15152


} else {
ifres15148 = False


}

ifres15145 = ifres15148


} else {
ifres15145 = False


}

ifres15142 = ifres15145


} else {
ifres15142 = False


}

ifres15140 = ifres15142


} else {
ifres15140 = False


}

let__14013 := ifres15140
_ = let__14013

tmp15816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14013, False)
}
__typedArg0 := let__14013
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15816 {
tmp15186 := Call(__e, PrimFunc(symshen_4unlocked_2), V5027)


var ifres15160 Obj

if True == tmp15186 {
tmp15161 := Call(__e, PrimFunc(symshen_4lazyderef), V5023, V5026)


let__14020 := tmp15161
_ = let__14020

tmp15185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14020)
}
__typedArg0 := let__14020
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15162 Obj

if True == tmp15185 {
tmp15163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14020)
}
__typedArg0 := let__14020
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14021 := tmp15163
_ = let__14021

tmp15164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14020)
}
__typedArg0 := let__14020
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15165 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15164, V5026)


let__14022 := tmp15165
_ = let__14022

tmp15184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14022)
}
__typedArg0 := let__14022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15166 Obj

if True == tmp15184 {
tmp15167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14022)
}
__typedArg0 := let__14022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14023 := tmp15167
_ = let__14023

tmp15168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14022)
}
__typedArg0 := let__14022
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15169 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15168, V5026)


let__14024 := tmp15169
_ = let__14024

tmp15183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14024, Nil)
}
__typedArg0 := let__14024
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15170 Obj

if True == tmp15183 {
tmp15171 := Call(__e, PrimFunc(symshen_4newpv), V5026)


let__14025 := tmp15171
_ = let__14025

tmp15172 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15172

tmp15173 := Call(__e, PrimFunc(symshen_4lazyderef), let__14021, V5026)


tmp15174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15173)
}
__typedArg0 := tmp15173
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

tmp15175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp15174)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp15174
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

tmp15176 := MakeNative(func(__e *ControlFlow) {
tmp15177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5024, Nil)
}
__typedArg0 := V5024
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15177)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15177
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14025, tmp15178)
}
__typedArg0 := let__14025
__typedArg1 := tmp15178
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15180 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14023, let__14025, V5025, V5026, V5027, let__13999, V5029)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4lookupsig), let__14021, tmp15179, V5026, V5027, let__13999, tmp15180)
return


}, 0)

tmp15181 := Call(__e, PrimFunc(symwhen), tmp15175, V5026, V5027, let__13999, tmp15176)


tmp15182 := Call(__e, PrimFunc(symshen_4gc), V5026, tmp15181)


ifres15170 = tmp15182


} else {
ifres15170 = False


}

ifres15166 = ifres15170


} else {
ifres15166 = False


}

ifres15162 = ifres15166


} else {
ifres15162 = False


}

ifres15160 = ifres15162


} else {
ifres15160 = False


}

let__14019 := ifres15160
_ = let__14019

tmp15814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14019, False)
}
__typedArg0 := let__14019
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15814 {
tmp15209 := Call(__e, PrimFunc(symshen_4unlocked_2), V5027)


var ifres15187 Obj

if True == tmp15209 {
tmp15188 := Call(__e, PrimFunc(symshen_4lazyderef), V5023, V5026)


let__14027 := tmp15188
_ = let__14027

tmp15208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14027)
}
__typedArg0 := let__14027
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15189 Obj

if True == tmp15208 {
tmp15190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14027)
}
__typedArg0 := let__14027
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14028 := tmp15190
_ = let__14028

tmp15191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14027)
}
__typedArg0 := let__14027
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15192 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15191, V5026)


let__14029 := tmp15192
_ = let__14029

tmp15207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14029)
}
__typedArg0 := let__14029
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15193 Obj

if True == tmp15207 {
tmp15194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14029)
}
__typedArg0 := let__14029
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14030 := tmp15194
_ = let__14030

tmp15195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14029)
}
__typedArg0 := let__14029
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15196 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15195, V5026)


let__14031 := tmp15196
_ = let__14031

tmp15206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14031, Nil)
}
__typedArg0 := let__14031
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15197 Obj

if True == tmp15206 {
tmp15198 := Call(__e, PrimFunc(symshen_4newpv), V5026)


let__14032 := tmp15198
_ = let__14032

tmp15199 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15199

tmp15200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5024, Nil)
}
__typedArg0 := V5024
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15200)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15200
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14032, tmp15201)
}
__typedArg0 := let__14032
__typedArg1 := tmp15201
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15203 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14030, let__14032, V5025, V5026, V5027, let__13999, V5029)
return
}, 0)

tmp15204 := Call(__e, PrimFunc(symshen_4system_1S_1h), let__14028, tmp15202, V5025, V5026, V5027, let__13999, tmp15203)


tmp15205 := Call(__e, PrimFunc(symshen_4gc), V5026, tmp15204)


ifres15197 = tmp15205


} else {
ifres15197 = False


}

ifres15193 = ifres15197


} else {
ifres15193 = False


}

ifres15189 = ifres15193


} else {
ifres15189 = False


}

ifres15187 = ifres15189


} else {
ifres15187 = False


}

let__14026 := ifres15187
_ = let__14026

tmp15812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14026, False)
}
__typedArg0 := let__14026
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15812 {
tmp15276 := Call(__e, PrimFunc(symshen_4unlocked_2), V5027)


var ifres15210 Obj

if True == tmp15276 {
tmp15211 := Call(__e, PrimFunc(symshen_4lazyderef), V5023, V5026)


let__14034 := tmp15211
_ = let__14034

tmp15275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14034)
}
__typedArg0 := let__14034
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15212 Obj

if True == tmp15275 {
tmp15213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14034)
}
__typedArg0 := let__14034
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15214 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15213, V5026)


let__14035 := tmp15214
_ = let__14035

tmp15274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14035, symcons)
}
__typedArg0 := let__14035
__typedArg1 := symcons
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15215 Obj

if True == tmp15274 {
tmp15216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14034)
}
__typedArg0 := let__14034
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15217 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15216, V5026)


let__14036 := tmp15217
_ = let__14036

tmp15273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14036)
}
__typedArg0 := let__14036
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15218 Obj

if True == tmp15273 {
tmp15219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14036)
}
__typedArg0 := let__14036
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14037 := tmp15219
_ = let__14037

tmp15220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14036)
}
__typedArg0 := let__14036
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15221 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15220, V5026)


let__14038 := tmp15221
_ = let__14038

tmp15272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14038)
}
__typedArg0 := let__14038
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15222 Obj

if True == tmp15272 {
tmp15223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14038)
}
__typedArg0 := let__14038
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14039 := tmp15223
_ = let__14039

tmp15224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14038)
}
__typedArg0 := let__14038
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15225 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15224, V5026)


let__14040 := tmp15225
_ = let__14040

tmp15271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14040, Nil)
}
__typedArg0 := let__14040
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15226 Obj

if True == tmp15271 {
tmp15227 := Call(__e, PrimFunc(symshen_4lazyderef), V5024, V5026)


let__14041 := tmp15227
_ = let__14041

tmp15228 := MakeNative(func(__e *ControlFlow) {
Z5074 := __e.Get(1)
_ = Z5074
tmp15229 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15229

tmp15230 := MakeNative(func(__e *ControlFlow) {
tmp15231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z5074, Nil)
}
__typedArg0 := Z5074
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp15231)
}
__typedArg0 := symlist
__typedArg1 := tmp15231
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14039, tmp15232, V5025, V5026, V5027, let__13999, V5029)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14037, Z5074, V5025, V5026, V5027, let__13999, tmp15230)
return


}, 1)

let__14042 := tmp15228
_ = let__14042

tmp15270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14041)
}
__typedArg0 := let__14041
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15233 Obj

if True == tmp15270 {
tmp15234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14041)
}
__typedArg0 := let__14041
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15235 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15234, V5026)


let__14043 := tmp15235
_ = let__14043

tmp15236 := MakeNative(func(__e *ControlFlow) {
tmp15237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14041)
}
__typedArg0 := let__14041
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15238 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15237, V5026)


let__14045 := tmp15238
_ = let__14045

tmp15239 := MakeNative(func(__e *ControlFlow) {
Z5079 := __e.Get(1)
_ = Z5079
__e.TailApply(let__14042, Z5079)
return
}, 1)

let__14046 := tmp15239
_ = let__14046

tmp15255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14045)
}
__typedArg0 := let__14045
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15255 {
tmp15240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14045)
}
__typedArg0 := let__14045
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14047 := tmp15240
_ = let__14047

tmp15241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14045)
}
__typedArg0 := let__14045
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15242 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15241, V5026)


let__14048 := tmp15242
_ = let__14048

tmp15243 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14046, let__14047)
return
}, 0)

let__14049 := tmp15243
_ = let__14049

tmp15247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14048, Nil)
}
__typedArg0 := let__14048
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15247 {
__e.TailApply(PrimFunc(symthaw), let__14049)
return
} else {
tmp15245 := Call(__e, PrimFunc(symshen_4pvar_2), let__14048)


if True == tmp15245 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14048, Nil, V5026, let__14049)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15253 := Call(__e, PrimFunc(symshen_4pvar_2), let__14045)


if True == tmp15253 {
tmp15248 := Call(__e, PrimFunc(symshen_4newpv), V5026)


let__14050 := tmp15248
_ = let__14050

tmp15249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14050, Nil)
}
__typedArg0 := let__14050
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15250 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14046, let__14050)
return
}, 0)

tmp15251 := Call(__e, PrimFunc(symshen_4bind_b), let__14045, tmp15249, V5026, tmp15250)


__e.TailApply(PrimFunc(symshen_4gc), V5026, tmp15251)
return


} else {
__e.Return(False)
return
}


}


}, 0)

let__14044 := tmp15236
_ = let__14044

tmp15261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14043, symlist)
}
__typedArg0 := let__14043
__typedArg1 := symlist
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15256 Obj

if True == tmp15261 {
tmp15257 := Call(__e, PrimFunc(symthaw), let__14044)


ifres15256 = tmp15257


} else {
tmp15260 := Call(__e, PrimFunc(symshen_4pvar_2), let__14043)


var ifres15258 Obj

if True == tmp15260 {
tmp15259 := Call(__e, PrimFunc(symshen_4bind_b), let__14043, symlist, V5026, let__14044)


ifres15258 = tmp15259


} else {
ifres15258 = False


}

ifres15256 = ifres15258


}

ifres15233 = ifres15256


} else {
tmp15269 := Call(__e, PrimFunc(symshen_4pvar_2), let__14041)


var ifres15262 Obj

if True == tmp15269 {
tmp15263 := Call(__e, PrimFunc(symshen_4newpv), V5026)


let__14051 := tmp15263
_ = let__14051

tmp15264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14051, Nil)
}
__typedArg0 := let__14051
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp15264)
}
__typedArg0 := symlist
__typedArg1 := tmp15264
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15266 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14042, let__14051)
return
}, 0)

tmp15267 := Call(__e, PrimFunc(symshen_4bind_b), let__14041, tmp15265, V5026, tmp15266)


tmp15268 := Call(__e, PrimFunc(symshen_4gc), V5026, tmp15267)


ifres15262 = tmp15268


} else {
ifres15262 = False


}

ifres15233 = ifres15262


}

ifres15226 = ifres15233


} else {
ifres15226 = False


}

ifres15222 = ifres15226


} else {
ifres15222 = False


}

ifres15218 = ifres15222


} else {
ifres15218 = False


}

ifres15215 = ifres15218


} else {
ifres15215 = False


}

ifres15212 = ifres15215


} else {
ifres15212 = False


}

ifres15210 = ifres15212


} else {
ifres15210 = False


}

let__14033 := ifres15210
_ = let__14033

tmp15810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14033, False)
}
__typedArg0 := let__14033
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15810 {
tmp15360 := Call(__e, PrimFunc(symshen_4unlocked_2), V5027)


var ifres15277 Obj

if True == tmp15360 {
tmp15278 := Call(__e, PrimFunc(symshen_4lazyderef), V5023, V5026)


let__14053 := tmp15278
_ = let__14053

tmp15359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14053)
}
__typedArg0 := let__14053
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15279 Obj

if True == tmp15359 {
tmp15280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14053)
}
__typedArg0 := let__14053
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15281 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15280, V5026)


let__14054 := tmp15281
_ = let__14054

tmp15358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14054, sym_8p)
}
__typedArg0 := let__14054
__typedArg1 := sym_8p
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15282 Obj

if True == tmp15358 {
tmp15283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14053)
}
__typedArg0 := let__14053
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15284 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15283, V5026)


let__14055 := tmp15284
_ = let__14055

tmp15357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14055)
}
__typedArg0 := let__14055
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15285 Obj

if True == tmp15357 {
tmp15286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14055)
}
__typedArg0 := let__14055
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14056 := tmp15286
_ = let__14056

tmp15287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14055)
}
__typedArg0 := let__14055
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15288 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15287, V5026)


let__14057 := tmp15288
_ = let__14057

tmp15356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14057)
}
__typedArg0 := let__14057
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15289 Obj

if True == tmp15356 {
tmp15290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14057)
}
__typedArg0 := let__14057
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14058 := tmp15290
_ = let__14058

tmp15291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14057)
}
__typedArg0 := let__14057
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15292 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15291, V5026)


let__14059 := tmp15292
_ = let__14059

tmp15355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14059, Nil)
}
__typedArg0 := let__14059
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15293 Obj

if True == tmp15355 {
tmp15294 := Call(__e, PrimFunc(symshen_4lazyderef), V5024, V5026)


let__14060 := tmp15294
_ = let__14060

tmp15295 := MakeNative(func(__e *ControlFlow) {
Z5095 := __e.Get(1)
_ = Z5095
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5096 := __e.Get(1)
_ = Z5096
tmp15296 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15296

tmp15297 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14058, Z5096, V5025, V5026, V5027, let__13999, V5029)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14056, Z5095, V5025, V5026, V5027, let__13999, tmp15297)
return


}, 1))
return
}, 1)

let__14061 := tmp15295
_ = let__14061

tmp15354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14060)
}
__typedArg0 := let__14060
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15298 Obj

if True == tmp15354 {
tmp15299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14060)
}
__typedArg0 := let__14060
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14062 := tmp15299
_ = let__14062

tmp15300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14060)
}
__typedArg0 := let__14060
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15301 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15300, V5026)


let__14063 := tmp15301
_ = let__14063

tmp15302 := MakeNative(func(__e *ControlFlow) {
Z5100 := __e.Get(1)
_ = Z5100
tmp15303 := Call(__e, let__14061, let__14062)


__e.TailApply(tmp15303, Z5100)
return


}, 1)

let__14064 := tmp15302
_ = let__14064

tmp15341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14063)
}
__typedArg0 := let__14063
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15304 Obj

if True == tmp15341 {
tmp15305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14063)
}
__typedArg0 := let__14063
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15306 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15305, V5026)


let__14065 := tmp15306
_ = let__14065

tmp15307 := MakeNative(func(__e *ControlFlow) {
tmp15308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14063)
}
__typedArg0 := let__14063
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15309 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15308, V5026)


let__14067 := tmp15309
_ = let__14067

tmp15310 := MakeNative(func(__e *ControlFlow) {
Z5105 := __e.Get(1)
_ = Z5105
__e.TailApply(let__14064, Z5105)
return
}, 1)

let__14068 := tmp15310
_ = let__14068

tmp15326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14067)
}
__typedArg0 := let__14067
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15326 {
tmp15311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14067)
}
__typedArg0 := let__14067
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14069 := tmp15311
_ = let__14069

tmp15312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14067)
}
__typedArg0 := let__14067
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15313 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15312, V5026)


let__14070 := tmp15313
_ = let__14070

tmp15314 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14068, let__14069)
return
}, 0)

let__14071 := tmp15314
_ = let__14071

tmp15318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14070, Nil)
}
__typedArg0 := let__14070
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15318 {
__e.TailApply(PrimFunc(symthaw), let__14071)
return
} else {
tmp15316 := Call(__e, PrimFunc(symshen_4pvar_2), let__14070)


if True == tmp15316 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14070, Nil, V5026, let__14071)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15324 := Call(__e, PrimFunc(symshen_4pvar_2), let__14067)


if True == tmp15324 {
tmp15319 := Call(__e, PrimFunc(symshen_4newpv), V5026)


let__14072 := tmp15319
_ = let__14072

tmp15320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14072, Nil)
}
__typedArg0 := let__14072
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15321 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14068, let__14072)
return
}, 0)

tmp15322 := Call(__e, PrimFunc(symshen_4bind_b), let__14067, tmp15320, V5026, tmp15321)


__e.TailApply(PrimFunc(symshen_4gc), V5026, tmp15322)
return


} else {
__e.Return(False)
return
}


}


}, 0)

let__14066 := tmp15307
_ = let__14066

tmp15332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14065, sym_d)
}
__typedArg0 := let__14065
__typedArg1 := sym_d
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15327 Obj

if True == tmp15332 {
tmp15328 := Call(__e, PrimFunc(symthaw), let__14066)


ifres15327 = tmp15328


} else {
tmp15331 := Call(__e, PrimFunc(symshen_4pvar_2), let__14065)


var ifres15329 Obj

if True == tmp15331 {
tmp15330 := Call(__e, PrimFunc(symshen_4bind_b), let__14065, sym_d, V5026, let__14066)


ifres15329 = tmp15330


} else {
ifres15329 = False


}

ifres15327 = ifres15329


}

ifres15304 = ifres15327


} else {
tmp15340 := Call(__e, PrimFunc(symshen_4pvar_2), let__14063)


var ifres15333 Obj

if True == tmp15340 {
tmp15334 := Call(__e, PrimFunc(symshen_4newpv), V5026)


let__14073 := tmp15334
_ = let__14073

tmp15335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14073, Nil)
}
__typedArg0 := let__14073
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp15335)
}
__typedArg0 := sym_d
__typedArg1 := tmp15335
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15337 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14064, let__14073)
return
}, 0)

tmp15338 := Call(__e, PrimFunc(symshen_4bind_b), let__14063, tmp15336, V5026, tmp15337)


tmp15339 := Call(__e, PrimFunc(symshen_4gc), V5026, tmp15338)


ifres15333 = tmp15339


} else {
ifres15333 = False


}

ifres15304 = ifres15333


}

ifres15298 = ifres15304


} else {
tmp15353 := Call(__e, PrimFunc(symshen_4pvar_2), let__14060)


var ifres15342 Obj

if True == tmp15353 {
tmp15343 := Call(__e, PrimFunc(symshen_4newpv), V5026)


let__14074 := tmp15343
_ = let__14074

tmp15344 := Call(__e, PrimFunc(symshen_4newpv), V5026)


let__14075 := tmp15344
_ = let__14075

tmp15345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14075, Nil)
}
__typedArg0 := let__14075
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp15345)
}
__typedArg0 := sym_d
__typedArg1 := tmp15345
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14074, tmp15346)
}
__typedArg0 := let__14074
__typedArg1 := tmp15346
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15348 := MakeNative(func(__e *ControlFlow) {
tmp15349 := Call(__e, let__14061, let__14074)


__e.TailApply(tmp15349, let__14075)
return


}, 0)

tmp15350 := Call(__e, PrimFunc(symshen_4bind_b), let__14060, tmp15347, V5026, tmp15348)


tmp15351 := Call(__e, PrimFunc(symshen_4gc), V5026, tmp15350)


tmp15352 := Call(__e, PrimFunc(symshen_4gc), V5026, tmp15351)


ifres15342 = tmp15352


} else {
ifres15342 = False


}

ifres15298 = ifres15342


}

ifres15293 = ifres15298


} else {
ifres15293 = False


}

ifres15289 = ifres15293


} else {
ifres15289 = False


}

ifres15285 = ifres15289


} else {
ifres15285 = False


}

ifres15282 = ifres15285


} else {
ifres15282 = False


}

ifres15279 = ifres15282


} else {
ifres15279 = False


}

ifres15277 = ifres15279


} else {
ifres15277 = False


}

let__14052 := ifres15277
_ = let__14052

tmp15808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14052, False)
}
__typedArg0 := let__14052
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15808 {
tmp15427 := Call(__e, PrimFunc(symshen_4unlocked_2), V5027)


var ifres15361 Obj

if True == tmp15427 {
tmp15362 := Call(__e, PrimFunc(symshen_4lazyderef), V5023, V5026)


let__14077 := tmp15362
_ = let__14077

tmp15426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14077)
}
__typedArg0 := let__14077
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15363 Obj

if True == tmp15426 {
tmp15364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14077)
}
__typedArg0 := let__14077
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15365 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15364, V5026)


let__14078 := tmp15365
_ = let__14078

tmp15425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14078, sym_8v)
}
__typedArg0 := let__14078
__typedArg1 := sym_8v
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15366 Obj

if True == tmp15425 {
tmp15367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14077)
}
__typedArg0 := let__14077
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15368 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15367, V5026)


let__14079 := tmp15368
_ = let__14079

tmp15424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14079)
}
__typedArg0 := let__14079
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15369 Obj

if True == tmp15424 {
tmp15370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14079)
}
__typedArg0 := let__14079
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14080 := tmp15370
_ = let__14080

tmp15371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14079)
}
__typedArg0 := let__14079
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15372 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15371, V5026)


let__14081 := tmp15372
_ = let__14081

tmp15423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14081)
}
__typedArg0 := let__14081
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15373 Obj

if True == tmp15423 {
tmp15374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14081)
}
__typedArg0 := let__14081
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14082 := tmp15374
_ = let__14082

tmp15375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14081)
}
__typedArg0 := let__14081
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15376 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15375, V5026)


let__14083 := tmp15376
_ = let__14083

tmp15422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14083, Nil)
}
__typedArg0 := let__14083
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15377 Obj

if True == tmp15422 {
tmp15378 := Call(__e, PrimFunc(symshen_4lazyderef), V5024, V5026)


let__14084 := tmp15378
_ = let__14084

tmp15379 := MakeNative(func(__e *ControlFlow) {
Z5123 := __e.Get(1)
_ = Z5123
tmp15380 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15380

tmp15381 := MakeNative(func(__e *ControlFlow) {
tmp15382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z5123, Nil)
}
__typedArg0 := Z5123
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp15382)
}
__typedArg0 := symvector
__typedArg1 := tmp15382
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14082, tmp15383, V5025, V5026, V5027, let__13999, V5029)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14080, Z5123, V5025, V5026, V5027, let__13999, tmp15381)
return


}, 1)

let__14085 := tmp15379
_ = let__14085

tmp15421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14084)
}
__typedArg0 := let__14084
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15384 Obj

if True == tmp15421 {
tmp15385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14084)
}
__typedArg0 := let__14084
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15386 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15385, V5026)


let__14086 := tmp15386
_ = let__14086

tmp15387 := MakeNative(func(__e *ControlFlow) {
tmp15388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14084)
}
__typedArg0 := let__14084
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15389 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15388, V5026)


let__14088 := tmp15389
_ = let__14088

tmp15390 := MakeNative(func(__e *ControlFlow) {
Z5128 := __e.Get(1)
_ = Z5128
__e.TailApply(let__14085, Z5128)
return
}, 1)

let__14089 := tmp15390
_ = let__14089

tmp15406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14088)
}
__typedArg0 := let__14088
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15406 {
tmp15391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14088)
}
__typedArg0 := let__14088
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14090 := tmp15391
_ = let__14090

tmp15392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14088)
}
__typedArg0 := let__14088
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15393 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15392, V5026)


let__14091 := tmp15393
_ = let__14091

tmp15394 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14089, let__14090)
return
}, 0)

let__14092 := tmp15394
_ = let__14092

tmp15398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14091, Nil)
}
__typedArg0 := let__14091
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15398 {
__e.TailApply(PrimFunc(symthaw), let__14092)
return
} else {
tmp15396 := Call(__e, PrimFunc(symshen_4pvar_2), let__14091)


if True == tmp15396 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14091, Nil, V5026, let__14092)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15404 := Call(__e, PrimFunc(symshen_4pvar_2), let__14088)


if True == tmp15404 {
tmp15399 := Call(__e, PrimFunc(symshen_4newpv), V5026)


let__14093 := tmp15399
_ = let__14093

tmp15400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14093, Nil)
}
__typedArg0 := let__14093
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15401 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14089, let__14093)
return
}, 0)

tmp15402 := Call(__e, PrimFunc(symshen_4bind_b), let__14088, tmp15400, V5026, tmp15401)


__e.TailApply(PrimFunc(symshen_4gc), V5026, tmp15402)
return


} else {
__e.Return(False)
return
}


}


}, 0)

let__14087 := tmp15387
_ = let__14087

tmp15412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14086, symvector)
}
__typedArg0 := let__14086
__typedArg1 := symvector
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15407 Obj

if True == tmp15412 {
tmp15408 := Call(__e, PrimFunc(symthaw), let__14087)


ifres15407 = tmp15408


} else {
tmp15411 := Call(__e, PrimFunc(symshen_4pvar_2), let__14086)


var ifres15409 Obj

if True == tmp15411 {
tmp15410 := Call(__e, PrimFunc(symshen_4bind_b), let__14086, symvector, V5026, let__14087)


ifres15409 = tmp15410


} else {
ifres15409 = False


}

ifres15407 = ifres15409


}

ifres15384 = ifres15407


} else {
tmp15420 := Call(__e, PrimFunc(symshen_4pvar_2), let__14084)


var ifres15413 Obj

if True == tmp15420 {
tmp15414 := Call(__e, PrimFunc(symshen_4newpv), V5026)


let__14094 := tmp15414
_ = let__14094

tmp15415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14094, Nil)
}
__typedArg0 := let__14094
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp15415)
}
__typedArg0 := symvector
__typedArg1 := tmp15415
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15417 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14085, let__14094)
return
}, 0)

tmp15418 := Call(__e, PrimFunc(symshen_4bind_b), let__14084, tmp15416, V5026, tmp15417)


tmp15419 := Call(__e, PrimFunc(symshen_4gc), V5026, tmp15418)


ifres15413 = tmp15419


} else {
ifres15413 = False


}

ifres15384 = ifres15413


}

ifres15377 = ifres15384


} else {
ifres15377 = False


}

ifres15373 = ifres15377


} else {
ifres15373 = False


}

ifres15369 = ifres15373


} else {
ifres15369 = False


}

ifres15366 = ifres15369


} else {
ifres15366 = False


}

ifres15363 = ifres15366


} else {
ifres15363 = False


}

ifres15361 = ifres15363


} else {
ifres15361 = False


}

let__14076 := ifres15361
_ = let__14076

tmp15806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14076, False)
}
__typedArg0 := let__14076
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15806 {
tmp15460 := Call(__e, PrimFunc(symshen_4unlocked_2), V5027)


var ifres15428 Obj

if True == tmp15460 {
tmp15429 := Call(__e, PrimFunc(symshen_4lazyderef), V5023, V5026)


let__14096 := tmp15429
_ = let__14096

tmp15459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14096)
}
__typedArg0 := let__14096
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15430 Obj

if True == tmp15459 {
tmp15431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14096)
}
__typedArg0 := let__14096
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15432 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15431, V5026)


let__14097 := tmp15432
_ = let__14097

tmp15458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14097, sym_8s)
}
__typedArg0 := let__14097
__typedArg1 := sym_8s
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15433 Obj

if True == tmp15458 {
tmp15434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14096)
}
__typedArg0 := let__14096
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15435 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15434, V5026)


let__14098 := tmp15435
_ = let__14098

tmp15457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14098)
}
__typedArg0 := let__14098
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15436 Obj

if True == tmp15457 {
tmp15437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14098)
}
__typedArg0 := let__14098
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14099 := tmp15437
_ = let__14099

tmp15438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14098)
}
__typedArg0 := let__14098
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15439 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15438, V5026)


let__14100 := tmp15439
_ = let__14100

tmp15456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14100)
}
__typedArg0 := let__14100
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15440 Obj

if True == tmp15456 {
tmp15441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14100)
}
__typedArg0 := let__14100
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14101 := tmp15441
_ = let__14101

tmp15442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14100)
}
__typedArg0 := let__14100
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15443 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15442, V5026)


let__14102 := tmp15443
_ = let__14102

tmp15455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14102, Nil)
}
__typedArg0 := let__14102
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15444 Obj

if True == tmp15455 {
tmp15445 := Call(__e, PrimFunc(symshen_4lazyderef), V5024, V5026)


let__14103 := tmp15445
_ = let__14103

tmp15446 := MakeNative(func(__e *ControlFlow) {
tmp15447 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15447

tmp15448 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14101, symstring, V5025, V5026, V5027, let__13999, V5029)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14099, symstring, V5025, V5026, V5027, let__13999, tmp15448)
return


}, 0)

let__14104 := tmp15446
_ = let__14104

tmp15454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14103, symstring)
}
__typedArg0 := let__14103
__typedArg1 := symstring
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15449 Obj

if True == tmp15454 {
tmp15450 := Call(__e, PrimFunc(symthaw), let__14104)


ifres15449 = tmp15450


} else {
tmp15453 := Call(__e, PrimFunc(symshen_4pvar_2), let__14103)


var ifres15451 Obj

if True == tmp15453 {
tmp15452 := Call(__e, PrimFunc(symshen_4bind_b), let__14103, symstring, V5026, let__14104)


ifres15451 = tmp15452


} else {
ifres15451 = False


}

ifres15449 = ifres15451


}

ifres15444 = ifres15449


} else {
ifres15444 = False


}

ifres15440 = ifres15444


} else {
ifres15440 = False


}

ifres15436 = ifres15440


} else {
ifres15436 = False


}

ifres15433 = ifres15436


} else {
ifres15433 = False


}

ifres15430 = ifres15433


} else {
ifres15430 = False


}

ifres15428 = ifres15430


} else {
ifres15428 = False


}

let__14095 := ifres15428
_ = let__14095

tmp15804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14095, False)
}
__typedArg0 := let__14095
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15804 {
tmp15560 := Call(__e, PrimFunc(symshen_4unlocked_2), V5027)


var ifres15461 Obj

if True == tmp15560 {
tmp15462 := Call(__e, PrimFunc(symshen_4lazyderef), V5023, V5026)


let__14106 := tmp15462
_ = let__14106

tmp15559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14106)
}
__typedArg0 := let__14106
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15463 Obj

if True == tmp15559 {
tmp15464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14106)
}
__typedArg0 := let__14106
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15465 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15464, V5026)


let__14107 := tmp15465
_ = let__14107

tmp15558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14107, symlambda)
}
__typedArg0 := let__14107
__typedArg1 := symlambda
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15466 Obj

if True == tmp15558 {
tmp15467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14106)
}
__typedArg0 := let__14106
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15468 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15467, V5026)


let__14108 := tmp15468
_ = let__14108

tmp15557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14108)
}
__typedArg0 := let__14108
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15469 Obj

if True == tmp15557 {
tmp15470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14108)
}
__typedArg0 := let__14108
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14109 := tmp15470
_ = let__14109

tmp15471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14108)
}
__typedArg0 := let__14108
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15472 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15471, V5026)


let__14110 := tmp15472
_ = let__14110

tmp15556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14110)
}
__typedArg0 := let__14110
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15473 Obj

if True == tmp15556 {
tmp15474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14110)
}
__typedArg0 := let__14110
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14111 := tmp15474
_ = let__14111

tmp15475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14110)
}
__typedArg0 := let__14110
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15476 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15475, V5026)


let__14112 := tmp15476
_ = let__14112

tmp15555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14112, Nil)
}
__typedArg0 := let__14112
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15477 Obj

if True == tmp15555 {
tmp15478 := Call(__e, PrimFunc(symshen_4lazyderef), V5024, V5026)


let__14113 := tmp15478
_ = let__14113

tmp15479 := MakeNative(func(__e *ControlFlow) {
Z5154 := __e.Get(1)
_ = Z5154
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5155 := __e.Get(1)
_ = Z5155
tmp15480 := Call(__e, PrimFunc(symshen_4newpv), V5026)


let__14115 := tmp15480
_ = let__14115

tmp15481 := Call(__e, PrimFunc(symshen_4newpv), V5026)


let__14116 := tmp15481
_ = let__14116

tmp15482 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15482

tmp15483 := Call(__e, PrimFunc(symshen_4lazyderef), let__14109, V5026)


tmp15484 := Call(__e, PrimFunc(symshen_4freshterm), tmp15483)


tmp15485 := MakeNative(func(__e *ControlFlow) {
tmp15486 := Call(__e, PrimFunc(symshen_4lazyderef), let__14109, V5026)


tmp15487 := Call(__e, PrimFunc(symshen_4deref), let__14116, V5026)


tmp15488 := Call(__e, PrimFunc(symshen_4deref), let__14111, V5026)


tmp15489 := Call(__e, PrimFunc(symshen_4beta), tmp15486, tmp15487, tmp15488)


tmp15490 := MakeNative(func(__e *ControlFlow) {
tmp15491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp15492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z5154, Nil)
}
__typedArg0 := Z5154
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15491, tmp15492)
}
__typedArg0 := tmp15491
__typedArg1 := tmp15492
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14116, tmp15493)
}
__typedArg0 := let__14116
__typedArg1 := tmp15493
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15494, V5025)
}
__typedArg0 := tmp15494
__typedArg1 := V5025
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14115, Z5155, tmp15495, V5026, V5027, let__13999, V5029)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14115, tmp15489, V5026, V5027, let__13999, tmp15490)
return


}, 0)

tmp15496 := Call(__e, PrimFunc(symbind), let__14116, tmp15484, V5026, V5027, let__13999, tmp15485)


tmp15497 := Call(__e, PrimFunc(symshen_4gc), V5026, tmp15496)


__e.TailApply(PrimFunc(symshen_4gc), V5026, tmp15497)
return


}, 1))
return
}, 1)

let__14114 := tmp15479
_ = let__14114

tmp15554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14113)
}
__typedArg0 := let__14113
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15498 Obj

if True == tmp15554 {
tmp15499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14113)
}
__typedArg0 := let__14113
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14117 := tmp15499
_ = let__14117

tmp15500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14113)
}
__typedArg0 := let__14113
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15501 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15500, V5026)


let__14118 := tmp15501
_ = let__14118

tmp15502 := MakeNative(func(__e *ControlFlow) {
Z5161 := __e.Get(1)
_ = Z5161
tmp15503 := Call(__e, let__14114, let__14117)


__e.TailApply(tmp15503, Z5161)
return


}, 1)

let__14119 := tmp15502
_ = let__14119

tmp15541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14118)
}
__typedArg0 := let__14118
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15504 Obj

if True == tmp15541 {
tmp15505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14118)
}
__typedArg0 := let__14118
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15506 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15505, V5026)


let__14120 := tmp15506
_ = let__14120

tmp15507 := MakeNative(func(__e *ControlFlow) {
tmp15508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14118)
}
__typedArg0 := let__14118
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15509 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15508, V5026)


let__14122 := tmp15509
_ = let__14122

tmp15510 := MakeNative(func(__e *ControlFlow) {
Z5166 := __e.Get(1)
_ = Z5166
__e.TailApply(let__14119, Z5166)
return
}, 1)

let__14123 := tmp15510
_ = let__14123

tmp15526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14122)
}
__typedArg0 := let__14122
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15526 {
tmp15511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14122)
}
__typedArg0 := let__14122
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14124 := tmp15511
_ = let__14124

tmp15512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14122)
}
__typedArg0 := let__14122
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15513 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15512, V5026)


let__14125 := tmp15513
_ = let__14125

tmp15514 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14123, let__14124)
return
}, 0)

let__14126 := tmp15514
_ = let__14126

tmp15518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14125, Nil)
}
__typedArg0 := let__14125
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15518 {
__e.TailApply(PrimFunc(symthaw), let__14126)
return
} else {
tmp15516 := Call(__e, PrimFunc(symshen_4pvar_2), let__14125)


if True == tmp15516 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14125, Nil, V5026, let__14126)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15524 := Call(__e, PrimFunc(symshen_4pvar_2), let__14122)


if True == tmp15524 {
tmp15519 := Call(__e, PrimFunc(symshen_4newpv), V5026)


let__14127 := tmp15519
_ = let__14127

tmp15520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14127, Nil)
}
__typedArg0 := let__14127
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15521 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14123, let__14127)
return
}, 0)

tmp15522 := Call(__e, PrimFunc(symshen_4bind_b), let__14122, tmp15520, V5026, tmp15521)


__e.TailApply(PrimFunc(symshen_4gc), V5026, tmp15522)
return


} else {
__e.Return(False)
return
}


}


}, 0)

let__14121 := tmp15507
_ = let__14121

tmp15532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14120, sym_1_1_6)
}
__typedArg0 := let__14120
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15527 Obj

if True == tmp15532 {
tmp15528 := Call(__e, PrimFunc(symthaw), let__14121)


ifres15527 = tmp15528


} else {
tmp15531 := Call(__e, PrimFunc(symshen_4pvar_2), let__14120)


var ifres15529 Obj

if True == tmp15531 {
tmp15530 := Call(__e, PrimFunc(symshen_4bind_b), let__14120, sym_1_1_6, V5026, let__14121)


ifres15529 = tmp15530


} else {
ifres15529 = False


}

ifres15527 = ifres15529


}

ifres15504 = ifres15527


} else {
tmp15540 := Call(__e, PrimFunc(symshen_4pvar_2), let__14118)


var ifres15533 Obj

if True == tmp15540 {
tmp15534 := Call(__e, PrimFunc(symshen_4newpv), V5026)


let__14128 := tmp15534
_ = let__14128

tmp15535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14128, Nil)
}
__typedArg0 := let__14128
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15535)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15535
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15537 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14119, let__14128)
return
}, 0)

tmp15538 := Call(__e, PrimFunc(symshen_4bind_b), let__14118, tmp15536, V5026, tmp15537)


tmp15539 := Call(__e, PrimFunc(symshen_4gc), V5026, tmp15538)


ifres15533 = tmp15539


} else {
ifres15533 = False


}

ifres15504 = ifres15533


}

ifres15498 = ifres15504


} else {
tmp15553 := Call(__e, PrimFunc(symshen_4pvar_2), let__14113)


var ifres15542 Obj

if True == tmp15553 {
tmp15543 := Call(__e, PrimFunc(symshen_4newpv), V5026)


let__14129 := tmp15543
_ = let__14129

tmp15544 := Call(__e, PrimFunc(symshen_4newpv), V5026)


let__14130 := tmp15544
_ = let__14130

tmp15545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14130, Nil)
}
__typedArg0 := let__14130
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15545)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15545
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14129, tmp15546)
}
__typedArg0 := let__14129
__typedArg1 := tmp15546
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15548 := MakeNative(func(__e *ControlFlow) {
tmp15549 := Call(__e, let__14114, let__14129)


__e.TailApply(tmp15549, let__14130)
return


}, 0)

tmp15550 := Call(__e, PrimFunc(symshen_4bind_b), let__14113, tmp15547, V5026, tmp15548)


tmp15551 := Call(__e, PrimFunc(symshen_4gc), V5026, tmp15550)


tmp15552 := Call(__e, PrimFunc(symshen_4gc), V5026, tmp15551)


ifres15542 = tmp15552


} else {
ifres15542 = False


}

ifres15498 = ifres15542


}

ifres15477 = ifres15498


} else {
ifres15477 = False


}

ifres15473 = ifres15477


} else {
ifres15473 = False


}

ifres15469 = ifres15473


} else {
ifres15469 = False


}

ifres15466 = ifres15469


} else {
ifres15466 = False


}

ifres15463 = ifres15466


} else {
ifres15463 = False


}

ifres15461 = ifres15463


} else {
ifres15461 = False


}

let__14105 := ifres15461
_ = let__14105

tmp15802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14105, False)
}
__typedArg0 := let__14105
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15802 {
tmp15610 := Call(__e, PrimFunc(symshen_4unlocked_2), V5027)


var ifres15561 Obj

if True == tmp15610 {
tmp15562 := Call(__e, PrimFunc(symshen_4lazyderef), V5023, V5026)


let__14132 := tmp15562
_ = let__14132

tmp15609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14132)
}
__typedArg0 := let__14132
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15563 Obj

if True == tmp15609 {
tmp15564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14132)
}
__typedArg0 := let__14132
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15565 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15564, V5026)


let__14133 := tmp15565
_ = let__14133

tmp15608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14133, symlet)
}
__typedArg0 := let__14133
__typedArg1 := symlet
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15566 Obj

if True == tmp15608 {
tmp15567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14132)
}
__typedArg0 := let__14132
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15568 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15567, V5026)


let__14134 := tmp15568
_ = let__14134

tmp15607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14134)
}
__typedArg0 := let__14134
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15569 Obj

if True == tmp15607 {
tmp15570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14134)
}
__typedArg0 := let__14134
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14135 := tmp15570
_ = let__14135

tmp15571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14134)
}
__typedArg0 := let__14134
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15572 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15571, V5026)


let__14136 := tmp15572
_ = let__14136

tmp15606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14136)
}
__typedArg0 := let__14136
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15573 Obj

if True == tmp15606 {
tmp15574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14136)
}
__typedArg0 := let__14136
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14137 := tmp15574
_ = let__14137

tmp15575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14136)
}
__typedArg0 := let__14136
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15576 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15575, V5026)


let__14138 := tmp15576
_ = let__14138

tmp15605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14138)
}
__typedArg0 := let__14138
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15577 Obj

if True == tmp15605 {
tmp15578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14138)
}
__typedArg0 := let__14138
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14139 := tmp15578
_ = let__14139

tmp15579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14138)
}
__typedArg0 := let__14138
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15580 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15579, V5026)


let__14140 := tmp15580
_ = let__14140

tmp15604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14140, Nil)
}
__typedArg0 := let__14140
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15581 Obj

if True == tmp15604 {
tmp15582 := Call(__e, PrimFunc(symshen_4newpv), V5026)


let__14141 := tmp15582
_ = let__14141

tmp15583 := Call(__e, PrimFunc(symshen_4newpv), V5026)


let__14142 := tmp15583
_ = let__14142

tmp15584 := Call(__e, PrimFunc(symshen_4newpv), V5026)


let__14143 := tmp15584
_ = let__14143

tmp15585 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15585

tmp15586 := MakeNative(func(__e *ControlFlow) {
tmp15587 := Call(__e, PrimFunc(symshen_4lazyderef), let__14135, V5026)


tmp15588 := Call(__e, PrimFunc(symshen_4freshterm), tmp15587)


tmp15589 := MakeNative(func(__e *ControlFlow) {
tmp15590 := Call(__e, PrimFunc(symshen_4lazyderef), let__14135, V5026)


tmp15591 := Call(__e, PrimFunc(symshen_4lazyderef), let__14142, V5026)


tmp15592 := Call(__e, PrimFunc(symshen_4lazyderef), let__14139, V5026)


tmp15593 := Call(__e, PrimFunc(symshen_4beta), tmp15590, tmp15591, tmp15592)


tmp15594 := MakeNative(func(__e *ControlFlow) {
tmp15595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp15596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14143, Nil)
}
__typedArg0 := let__14143
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15595, tmp15596)
}
__typedArg0 := tmp15595
__typedArg1 := tmp15596
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14142, tmp15597)
}
__typedArg0 := let__14142
__typedArg1 := tmp15597
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15598, V5025)
}
__typedArg0 := tmp15598
__typedArg1 := V5025
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14141, V5024, tmp15599, V5026, V5027, let__13999, V5029)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14141, tmp15593, V5026, V5027, let__13999, tmp15594)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14142, tmp15588, V5026, V5027, let__13999, tmp15589)
return


}, 0)

tmp15600 := Call(__e, PrimFunc(symshen_4system_1S_1h), let__14137, let__14143, V5025, V5026, V5027, let__13999, tmp15586)


tmp15601 := Call(__e, PrimFunc(symshen_4gc), V5026, tmp15600)


tmp15602 := Call(__e, PrimFunc(symshen_4gc), V5026, tmp15601)


tmp15603 := Call(__e, PrimFunc(symshen_4gc), V5026, tmp15602)


ifres15581 = tmp15603


} else {
ifres15581 = False


}

ifres15577 = ifres15581


} else {
ifres15577 = False


}

ifres15573 = ifres15577


} else {
ifres15573 = False


}

ifres15569 = ifres15573


} else {
ifres15569 = False


}

ifres15566 = ifres15569


} else {
ifres15566 = False


}

ifres15563 = ifres15566


} else {
ifres15563 = False


}

ifres15561 = ifres15563


} else {
ifres15561 = False


}

let__14131 := ifres15561
_ = let__14131

tmp15800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14131, False)
}
__typedArg0 := let__14131
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15800 {
tmp15680 := Call(__e, PrimFunc(symshen_4unlocked_2), V5027)


var ifres15611 Obj

if True == tmp15680 {
tmp15612 := Call(__e, PrimFunc(symshen_4lazyderef), V5023, V5026)


let__14145 := tmp15612
_ = let__14145

tmp15679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14145)
}
__typedArg0 := let__14145
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15613 Obj

if True == tmp15679 {
tmp15614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14145)
}
__typedArg0 := let__14145
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15615 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15614, V5026)


let__14146 := tmp15615
_ = let__14146

tmp15678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14146, symopen)
}
__typedArg0 := let__14146
__typedArg1 := symopen
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15616 Obj

if True == tmp15678 {
tmp15617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14145)
}
__typedArg0 := let__14145
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15618 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15617, V5026)


let__14147 := tmp15618
_ = let__14147

tmp15677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14147)
}
__typedArg0 := let__14147
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15619 Obj

if True == tmp15677 {
tmp15620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14147)
}
__typedArg0 := let__14147
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14148 := tmp15620
_ = let__14148

tmp15621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14147)
}
__typedArg0 := let__14147
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15622 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15621, V5026)


let__14149 := tmp15622
_ = let__14149

tmp15676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14149)
}
__typedArg0 := let__14149
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15623 Obj

if True == tmp15676 {
tmp15624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14149)
}
__typedArg0 := let__14149
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14150 := tmp15624
_ = let__14150

tmp15625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14149)
}
__typedArg0 := let__14149
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15626 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15625, V5026)


let__14151 := tmp15626
_ = let__14151

tmp15675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14151, Nil)
}
__typedArg0 := let__14151
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15627 Obj

if True == tmp15675 {
tmp15628 := Call(__e, PrimFunc(symshen_4lazyderef), V5024, V5026)


let__14152 := tmp15628
_ = let__14152

tmp15629 := MakeNative(func(__e *ControlFlow) {
Z5197 := __e.Get(1)
_ = Z5197
tmp15630 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15630

tmp15631 := MakeNative(func(__e *ControlFlow) {
tmp15632 := Call(__e, PrimFunc(symshen_4lazyderef), Z5197, V5026)


tmp15633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, tmp15633)
}
__typedArg0 := symin
__typedArg1 := tmp15633
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15635 := Call(__e, PrimFunc(symelement_2), tmp15632, tmp15634)


tmp15636 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14148, symstring, V5025, V5026, V5027, let__13999, V5029)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15635, V5026, V5027, let__13999, tmp15636)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), let__14150, Z5197, V5026, V5027, let__13999, tmp15631)
return


}, 1)

let__14153 := tmp15629
_ = let__14153

tmp15674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14152)
}
__typedArg0 := let__14152
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15637 Obj

if True == tmp15674 {
tmp15638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14152)
}
__typedArg0 := let__14152
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15639 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15638, V5026)


let__14154 := tmp15639
_ = let__14154

tmp15640 := MakeNative(func(__e *ControlFlow) {
tmp15641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14152)
}
__typedArg0 := let__14152
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15642 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15641, V5026)


let__14156 := tmp15642
_ = let__14156

tmp15643 := MakeNative(func(__e *ControlFlow) {
Z5202 := __e.Get(1)
_ = Z5202
__e.TailApply(let__14153, Z5202)
return
}, 1)

let__14157 := tmp15643
_ = let__14157

tmp15659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14156)
}
__typedArg0 := let__14156
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15659 {
tmp15644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14156)
}
__typedArg0 := let__14156
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14158 := tmp15644
_ = let__14158

tmp15645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14156)
}
__typedArg0 := let__14156
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15646 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15645, V5026)


let__14159 := tmp15646
_ = let__14159

tmp15647 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14157, let__14158)
return
}, 0)

let__14160 := tmp15647
_ = let__14160

tmp15651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14159, Nil)
}
__typedArg0 := let__14159
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15651 {
__e.TailApply(PrimFunc(symthaw), let__14160)
return
} else {
tmp15649 := Call(__e, PrimFunc(symshen_4pvar_2), let__14159)


if True == tmp15649 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14159, Nil, V5026, let__14160)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15657 := Call(__e, PrimFunc(symshen_4pvar_2), let__14156)


if True == tmp15657 {
tmp15652 := Call(__e, PrimFunc(symshen_4newpv), V5026)


let__14161 := tmp15652
_ = let__14161

tmp15653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14161, Nil)
}
__typedArg0 := let__14161
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15654 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14157, let__14161)
return
}, 0)

tmp15655 := Call(__e, PrimFunc(symshen_4bind_b), let__14156, tmp15653, V5026, tmp15654)


__e.TailApply(PrimFunc(symshen_4gc), V5026, tmp15655)
return


} else {
__e.Return(False)
return
}


}


}, 0)

let__14155 := tmp15640
_ = let__14155

tmp15665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14154, symstream)
}
__typedArg0 := let__14154
__typedArg1 := symstream
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15660 Obj

if True == tmp15665 {
tmp15661 := Call(__e, PrimFunc(symthaw), let__14155)


ifres15660 = tmp15661


} else {
tmp15664 := Call(__e, PrimFunc(symshen_4pvar_2), let__14154)


var ifres15662 Obj

if True == tmp15664 {
tmp15663 := Call(__e, PrimFunc(symshen_4bind_b), let__14154, symstream, V5026, let__14155)


ifres15662 = tmp15663


} else {
ifres15662 = False


}

ifres15660 = ifres15662


}

ifres15637 = ifres15660


} else {
tmp15673 := Call(__e, PrimFunc(symshen_4pvar_2), let__14152)


var ifres15666 Obj

if True == tmp15673 {
tmp15667 := Call(__e, PrimFunc(symshen_4newpv), V5026)


let__14162 := tmp15667
_ = let__14162

tmp15668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14162, Nil)
}
__typedArg0 := let__14162
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp15668)
}
__typedArg0 := symstream
__typedArg1 := tmp15668
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15670 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14153, let__14162)
return
}, 0)

tmp15671 := Call(__e, PrimFunc(symshen_4bind_b), let__14152, tmp15669, V5026, tmp15670)


tmp15672 := Call(__e, PrimFunc(symshen_4gc), V5026, tmp15671)


ifres15666 = tmp15672


} else {
ifres15666 = False


}

ifres15637 = ifres15666


}

ifres15627 = ifres15637


} else {
ifres15627 = False


}

ifres15623 = ifres15627


} else {
ifres15623 = False


}

ifres15619 = ifres15623


} else {
ifres15619 = False


}

ifres15616 = ifres15619


} else {
ifres15616 = False


}

ifres15613 = ifres15616


} else {
ifres15613 = False


}

ifres15611 = ifres15613


} else {
ifres15611 = False


}

let__14144 := ifres15611
_ = let__14144

tmp15798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14144, False)
}
__typedArg0 := let__14144
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15798 {
tmp15709 := Call(__e, PrimFunc(symshen_4unlocked_2), V5027)


var ifres15681 Obj

if True == tmp15709 {
tmp15682 := Call(__e, PrimFunc(symshen_4lazyderef), V5023, V5026)


let__14164 := tmp15682
_ = let__14164

tmp15708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14164)
}
__typedArg0 := let__14164
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15683 Obj

if True == tmp15708 {
tmp15684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14164)
}
__typedArg0 := let__14164
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15685 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15684, V5026)


let__14165 := tmp15685
_ = let__14165

tmp15707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14165, symtype)
}
__typedArg0 := let__14165
__typedArg1 := symtype
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15686 Obj

if True == tmp15707 {
tmp15687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14164)
}
__typedArg0 := let__14164
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15688 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15687, V5026)


let__14166 := tmp15688
_ = let__14166

tmp15706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14166)
}
__typedArg0 := let__14166
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15689 Obj

if True == tmp15706 {
tmp15690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14166)
}
__typedArg0 := let__14166
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14167 := tmp15690
_ = let__14167

tmp15691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14166)
}
__typedArg0 := let__14166
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15692 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15691, V5026)


let__14168 := tmp15692
_ = let__14168

tmp15705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14168)
}
__typedArg0 := let__14168
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15693 Obj

if True == tmp15705 {
tmp15694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14168)
}
__typedArg0 := let__14168
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14169 := tmp15694
_ = let__14169

tmp15695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14168)
}
__typedArg0 := let__14168
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15696 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15695, V5026)


let__14170 := tmp15696
_ = let__14170

tmp15704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14170, Nil)
}
__typedArg0 := let__14170
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15697 Obj

if True == tmp15704 {
tmp15698 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15698

tmp15699 := MakeNative(func(__e *ControlFlow) {
tmp15700 := Call(__e, PrimFunc(symshen_4deref), let__14169, V5026)


tmp15701 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp15700)


tmp15702 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14167, V5024, V5025, V5026, V5027, let__13999, V5029)
return
}, 0)

__e.TailApply(PrimFunc(symis_b), tmp15701, V5024, V5026, V5027, let__13999, tmp15702)
return


}, 0)

tmp15703 := Call(__e, PrimFunc(symshen_4cut), V5026, V5027, let__13999, tmp15699)


ifres15697 = tmp15703


} else {
ifres15697 = False


}

ifres15693 = ifres15697


} else {
ifres15693 = False


}

ifres15689 = ifres15693


} else {
ifres15689 = False


}

ifres15686 = ifres15689


} else {
ifres15686 = False


}

ifres15683 = ifres15686


} else {
ifres15683 = False


}

ifres15681 = ifres15683


} else {
ifres15681 = False


}

let__14163 := ifres15681
_ = let__14163

tmp15796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14163, False)
}
__typedArg0 := let__14163
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15796 {
tmp15740 := Call(__e, PrimFunc(symshen_4unlocked_2), V5027)


var ifres15710 Obj

if True == tmp15740 {
tmp15711 := Call(__e, PrimFunc(symshen_4lazyderef), V5023, V5026)


let__14172 := tmp15711
_ = let__14172

tmp15739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14172)
}
__typedArg0 := let__14172
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15712 Obj

if True == tmp15739 {
tmp15713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14172)
}
__typedArg0 := let__14172
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15714 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15713, V5026)


let__14173 := tmp15714
_ = let__14173

tmp15738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14173, symshen_4input_1h_7)
}
__typedArg0 := let__14173
__typedArg1 := symshen_4input_1h_7
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15715 Obj

if True == tmp15738 {
tmp15716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14172)
}
__typedArg0 := let__14172
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15717 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15716, V5026)


let__14174 := tmp15717
_ = let__14174

tmp15737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14174)
}
__typedArg0 := let__14174
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15718 Obj

if True == tmp15737 {
tmp15719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14174)
}
__typedArg0 := let__14174
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14175 := tmp15719
_ = let__14175

tmp15720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14174)
}
__typedArg0 := let__14174
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15721 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15720, V5026)


let__14176 := tmp15721
_ = let__14176

tmp15736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14176)
}
__typedArg0 := let__14176
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15722 Obj

if True == tmp15736 {
tmp15723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14176)
}
__typedArg0 := let__14176
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14177 := tmp15723
_ = let__14177

tmp15724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14176)
}
__typedArg0 := let__14176
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15725 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15724, V5026)


let__14178 := tmp15725
_ = let__14178

tmp15735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14178, Nil)
}
__typedArg0 := let__14178
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15726 Obj

if True == tmp15735 {
tmp15727 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15727

tmp15728 := Call(__e, PrimFunc(symshen_4deref), let__14175, V5026)


tmp15729 := Call(__e, PrimFunc(symshen_4rdecons), tmp15728)


tmp15730 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp15729)


tmp15731 := MakeNative(func(__e *ControlFlow) {
tmp15732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp15732)
}
__typedArg0 := symstream
__typedArg1 := tmp15732
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14177, tmp15733, V5025, V5026, V5027, let__13999, V5029)
return


}, 0)

tmp15734 := Call(__e, PrimFunc(symis), V5024, tmp15730, V5026, V5027, let__13999, tmp15731)


ifres15726 = tmp15734


} else {
ifres15726 = False


}

ifres15722 = ifres15726


} else {
ifres15722 = False


}

ifres15718 = ifres15722


} else {
ifres15718 = False


}

ifres15715 = ifres15718


} else {
ifres15715 = False


}

ifres15712 = ifres15715


} else {
ifres15712 = False


}

ifres15710 = ifres15712


} else {
ifres15710 = False


}

let__14171 := ifres15710
_ = let__14171

tmp15794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14171, False)
}
__typedArg0 := let__14171
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15794 {
tmp15769 := Call(__e, PrimFunc(symshen_4unlocked_2), V5027)


var ifres15741 Obj

if True == tmp15769 {
tmp15742 := Call(__e, PrimFunc(symshen_4lazyderef), V5023, V5026)


let__14180 := tmp15742
_ = let__14180

tmp15768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14180)
}
__typedArg0 := let__14180
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15743 Obj

if True == tmp15768 {
tmp15744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14180)
}
__typedArg0 := let__14180
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15745 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15744, V5026)


let__14181 := tmp15745
_ = let__14181

tmp15767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14181, symset)
}
__typedArg0 := let__14181
__typedArg1 := symset
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15746 Obj

if True == tmp15767 {
tmp15747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14180)
}
__typedArg0 := let__14180
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15748 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15747, V5026)


let__14182 := tmp15748
_ = let__14182

tmp15766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14182)
}
__typedArg0 := let__14182
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15749 Obj

if True == tmp15766 {
tmp15750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14182)
}
__typedArg0 := let__14182
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14183 := tmp15750
_ = let__14183

tmp15751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14182)
}
__typedArg0 := let__14182
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15752 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15751, V5026)


let__14184 := tmp15752
_ = let__14184

tmp15765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14184)
}
__typedArg0 := let__14184
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15753 Obj

if True == tmp15765 {
tmp15754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14184)
}
__typedArg0 := let__14184
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14185 := tmp15754
_ = let__14185

tmp15755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14184)
}
__typedArg0 := let__14184
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15756 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15755, V5026)


let__14186 := tmp15756
_ = let__14186

tmp15764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14186, Nil)
}
__typedArg0 := let__14186
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15757 Obj

if True == tmp15764 {
tmp15758 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15758

tmp15759 := MakeNative(func(__e *ControlFlow) {
tmp15760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14183, Nil)
}
__typedArg0 := let__14183
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp15760)
}
__typedArg0 := symvalue
__typedArg1 := tmp15760
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15762 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14185, V5024, V5025, V5026, V5027, let__13999, V5029)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp15761, V5024, V5025, V5026, V5027, let__13999, tmp15762)
return


}, 0)

tmp15763 := Call(__e, PrimFunc(symshen_4system_1S_1h), let__14183, symsymbol, V5025, V5026, V5027, let__13999, tmp15759)


ifres15757 = tmp15763


} else {
ifres15757 = False


}

ifres15753 = ifres15757


} else {
ifres15753 = False


}

ifres15749 = ifres15753


} else {
ifres15749 = False


}

ifres15746 = ifres15749


} else {
ifres15746 = False


}

ifres15743 = ifres15746


} else {
ifres15743 = False


}

ifres15741 = ifres15743


} else {
ifres15741 = False


}

let__14179 := ifres15741
_ = let__14179

tmp15792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14179, False)
}
__typedArg0 := let__14179
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15792 {
tmp15777 := Call(__e, PrimFunc(symshen_4unlocked_2), V5027)


var ifres15770 Obj

if True == tmp15777 {
tmp15771 := Call(__e, PrimFunc(symshen_4newpv), V5026)


let__14188 := tmp15771
_ = let__14188

tmp15772 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15772

tmp15773 := MakeNative(func(__e *ControlFlow) {
tmp15774 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), V5023, V5024, let__14188, V5026, V5027, let__13999, V5029)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5026, V5027, let__13999, tmp15774)
return


}, 0)

tmp15775 := Call(__e, PrimFunc(symshen_4l_1rules), V5025, let__14188, False, V5026, V5027, let__13999, tmp15773)


tmp15776 := Call(__e, PrimFunc(symshen_4gc), V5026, tmp15775)


ifres15770 = tmp15776


} else {
ifres15770 = False


}

let__14187 := ifres15770
_ = let__14187

tmp15790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14187, False)
}
__typedArg0 := let__14187
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15790 {
tmp15786 := Call(__e, PrimFunc(symshen_4unlocked_2), V5027)


var ifres15778 Obj

if True == tmp15786 {
tmp15779 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15779

tmp15780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp15781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5024, Nil)
}
__typedArg0 := V5024
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15780, tmp15781)
}
__typedArg0 := tmp15780
__typedArg1 := tmp15781
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5023, tmp15782)
}
__typedArg0 := V5023
__typedArg1 := tmp15782
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp15785 := Call(__e, PrimFunc(symshen_4search_1user_1datatypes), tmp15783, V5025, tmp15784, V5026, V5027, let__13999, V5029)


ifres15778 = tmp15785


} else {
ifres15778 = False


}

let__14189 := ifres15778
_ = let__14189

tmp15788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14189, False)
}
__typedArg0 := let__14189
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15788 {
__e.TailApply(PrimFunc(symshen_4unlock), V5027, let__13999)
return
} else {
__e.Return(let__14189)
return
}


} else {
__e.Return(let__14187)
return
}


} else {
__e.Return(let__14179)
return
}


} else {
__e.Return(let__14171)
return
}


} else {
__e.Return(let__14163)
return
}


} else {
__e.Return(let__14144)
return
}


} else {
__e.Return(let__14131)
return
}


} else {
__e.Return(let__14105)
return
}


} else {
__e.Return(let__14095)
return
}


} else {
__e.Return(let__14076)
return
}


} else {
__e.Return(let__14052)
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
__e.Return(let__14019)
return
}


} else {
__e.Return(let__14013)
return
}


} else {
__e.Return(let__14007)
return
}


} else {
__e.Return(let__14003)
return
}


} else {
__e.Return(let__14002)
return
}


} else {
__e.Return(let__14001)
return
}


} else {
__e.Return(let__14000)
return
}


}, 7)

tmp15827 := Call(__e, ns2_1set, symshen_4system_1S_1h, tmp15076)


_ = tmp15827

tmp15828 := MakeNative(func(__e *ControlFlow) {
V5235 := __e.Get(1)
_ = V5235
tmp15862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5235)
}
__typedArg0 := V5235
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15843 Obj

if True == tmp15862 {
tmp15860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5235)
}
__typedArg0 := V5235
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp15860)
}
__typedArg0 := symcons
__typedArg1 := tmp15860
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15845 Obj

if True == tmp15861 {
tmp15858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5235)
}
__typedArg0 := V5235
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15858)
}
__typedArg0 := tmp15858
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15847 Obj

if True == tmp15859 {
tmp15855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5235)
}
__typedArg0 := V5235
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15855)
}
__typedArg0 := tmp15855
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15856)
}
__typedArg0 := tmp15856
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15849 Obj

if True == tmp15857 {
tmp15851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5235)
}
__typedArg0 := V5235
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15851)
}
__typedArg0 := tmp15851
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15852)
}
__typedArg0 := tmp15852
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp15853)
}
__typedArg0 := Nil
__typedArg1 := tmp15853
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15850 Obj

if True == tmp15854 {
ifres15850 = True


} else {
ifres15850 = False


}

ifres15849 = ifres15850


} else {
ifres15849 = False


}

var ifres15848 Obj

if True == ifres15849 {
ifres15848 = True


} else {
ifres15848 = False


}

ifres15847 = ifres15848


} else {
ifres15847 = False


}

var ifres15846 Obj

if True == ifres15847 {
ifres15846 = True


} else {
ifres15846 = False


}

ifres15845 = ifres15846


} else {
ifres15845 = False


}

var ifres15844 Obj

if True == ifres15845 {
ifres15844 = True


} else {
ifres15844 = False


}

ifres15843 = ifres15844


} else {
ifres15843 = False


}

if True == ifres15843 {
tmp15829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5235)
}
__typedArg0 := V5235
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15829)
}
__typedArg0 := tmp15829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15831 := Call(__e, PrimFunc(symshen_4rdecons), tmp15830)


tmp15832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5235)
}
__typedArg0 := V5235
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15832)
}
__typedArg0 := tmp15832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15833)
}
__typedArg0 := tmp15833
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15835 := Call(__e, PrimFunc(symshen_4rdecons), tmp15834)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15831, tmp15835)
}
__typedArg0 := tmp15831
__typedArg1 := tmp15835
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp15841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5235)
}
__typedArg0 := V5235
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15841 {
tmp15836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5235)
}
__typedArg0 := V5235
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15837 := Call(__e, PrimFunc(symshen_4rdecons), tmp15836)


tmp15838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5235)
}
__typedArg0 := V5235
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15839 := Call(__e, PrimFunc(symshen_4rdecons), tmp15838)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15837, tmp15839)
}
__typedArg0 := tmp15837
__typedArg1 := tmp15839
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V5235)
return
}


}


}, 1)

tmp15863 := Call(__e, ns2_1set, symshen_4rdecons, tmp15828)


_ = tmp15863

tmp15864 := MakeNative(func(__e *ControlFlow) {
V5236 := __e.Get(1)
_ = V5236
V5237 := __e.Get(2)
_ = V5237
V5238 := __e.Get(3)
_ = V5238
V5239 := __e.Get(4)
_ = V5239
V5240 := __e.Get(5)
_ = V5240
V5241 := __e.Get(6)
_ = V5241
tmp15877 := Call(__e, PrimFunc(symshen_4unlocked_2), V5239)


var ifres15865 Obj

if True == tmp15877 {
tmp15866 := Call(__e, PrimFunc(symshen_4lazyderef), V5237, V5238)


let__14191 := tmp15866
_ = let__14191

tmp15867 := MakeNative(func(__e *ControlFlow) {
tmp15868 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15868

tmp15869 := Call(__e, PrimFunc(symshen_4lazyderef), V5236, V5238)


tmp15870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(tmp15869)
}
__typedArg0 := tmp15869
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

__e.TailApply(PrimFunc(symwhen), tmp15870, V5238, V5239, V5240, V5241)
return


}, 0)

let__14192 := tmp15867
_ = let__14192

tmp15876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14191, symnumber)
}
__typedArg0 := let__14191
__typedArg1 := symnumber
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15871 Obj

if True == tmp15876 {
tmp15872 := Call(__e, PrimFunc(symthaw), let__14192)


ifres15871 = tmp15872


} else {
tmp15875 := Call(__e, PrimFunc(symshen_4pvar_2), let__14191)


var ifres15873 Obj

if True == tmp15875 {
tmp15874 := Call(__e, PrimFunc(symshen_4bind_b), let__14191, symnumber, V5238, let__14192)


ifres15873 = tmp15874


} else {
ifres15873 = False


}

ifres15871 = ifres15873


}

ifres15865 = ifres15871


} else {
ifres15865 = False


}

let__14190 := ifres15865
_ = let__14190

tmp15967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14190, False)
}
__typedArg0 := let__14190
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15967 {
tmp15890 := Call(__e, PrimFunc(symshen_4unlocked_2), V5239)


var ifres15878 Obj

if True == tmp15890 {
tmp15879 := Call(__e, PrimFunc(symshen_4lazyderef), V5237, V5238)


let__14194 := tmp15879
_ = let__14194

tmp15880 := MakeNative(func(__e *ControlFlow) {
tmp15881 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15881

tmp15882 := Call(__e, PrimFunc(symshen_4lazyderef), V5236, V5238)


tmp15883 := Call(__e, PrimFunc(symboolean_2), tmp15882)


__e.TailApply(PrimFunc(symwhen), tmp15883, V5238, V5239, V5240, V5241)
return


}, 0)

let__14195 := tmp15880
_ = let__14195

tmp15889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14194, symboolean)
}
__typedArg0 := let__14194
__typedArg1 := symboolean
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15884 Obj

if True == tmp15889 {
tmp15885 := Call(__e, PrimFunc(symthaw), let__14195)


ifres15884 = tmp15885


} else {
tmp15888 := Call(__e, PrimFunc(symshen_4pvar_2), let__14194)


var ifres15886 Obj

if True == tmp15888 {
tmp15887 := Call(__e, PrimFunc(symshen_4bind_b), let__14194, symboolean, V5238, let__14195)


ifres15886 = tmp15887


} else {
ifres15886 = False


}

ifres15884 = ifres15886


}

ifres15878 = ifres15884


} else {
ifres15878 = False


}

let__14193 := ifres15878
_ = let__14193

tmp15965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14193, False)
}
__typedArg0 := let__14193
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15965 {
tmp15903 := Call(__e, PrimFunc(symshen_4unlocked_2), V5239)


var ifres15891 Obj

if True == tmp15903 {
tmp15892 := Call(__e, PrimFunc(symshen_4lazyderef), V5237, V5238)


let__14197 := tmp15892
_ = let__14197

tmp15893 := MakeNative(func(__e *ControlFlow) {
tmp15894 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15894

tmp15895 := Call(__e, PrimFunc(symshen_4lazyderef), V5236, V5238)


tmp15896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(tmp15895)
}
__typedArg0 := tmp15895
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

__e.TailApply(PrimFunc(symwhen), tmp15896, V5238, V5239, V5240, V5241)
return


}, 0)

let__14198 := tmp15893
_ = let__14198

tmp15902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14197, symstring)
}
__typedArg0 := let__14197
__typedArg1 := symstring
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15897 Obj

if True == tmp15902 {
tmp15898 := Call(__e, PrimFunc(symthaw), let__14198)


ifres15897 = tmp15898


} else {
tmp15901 := Call(__e, PrimFunc(symshen_4pvar_2), let__14197)


var ifres15899 Obj

if True == tmp15901 {
tmp15900 := Call(__e, PrimFunc(symshen_4bind_b), let__14197, symstring, V5238, let__14198)


ifres15899 = tmp15900


} else {
ifres15899 = False


}

ifres15897 = ifres15899


}

ifres15891 = ifres15897


} else {
ifres15891 = False


}

let__14196 := ifres15891
_ = let__14196

tmp15963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14196, False)
}
__typedArg0 := let__14196
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15963 {
tmp15916 := Call(__e, PrimFunc(symshen_4unlocked_2), V5239)


var ifres15904 Obj

if True == tmp15916 {
tmp15905 := Call(__e, PrimFunc(symshen_4lazyderef), V5237, V5238)


let__14200 := tmp15905
_ = let__14200

tmp15906 := MakeNative(func(__e *ControlFlow) {
tmp15907 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15907

tmp15908 := Call(__e, PrimFunc(symshen_4lazyderef), V5236, V5238)


tmp15909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(tmp15908)
}
__typedArg0 := tmp15908
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

__e.TailApply(PrimFunc(symwhen), tmp15909, V5238, V5239, V5240, V5241)
return


}, 0)

let__14201 := tmp15906
_ = let__14201

tmp15915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14200, symsymbol)
}
__typedArg0 := let__14200
__typedArg1 := symsymbol
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15910 Obj

if True == tmp15915 {
tmp15911 := Call(__e, PrimFunc(symthaw), let__14201)


ifres15910 = tmp15911


} else {
tmp15914 := Call(__e, PrimFunc(symshen_4pvar_2), let__14200)


var ifres15912 Obj

if True == tmp15914 {
tmp15913 := Call(__e, PrimFunc(symshen_4bind_b), let__14200, symsymbol, V5238, let__14201)


ifres15912 = tmp15913


} else {
ifres15912 = False


}

ifres15910 = ifres15912


}

ifres15904 = ifres15910


} else {
ifres15904 = False


}

let__14199 := ifres15904
_ = let__14199

tmp15961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14199, False)
}
__typedArg0 := let__14199
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15961 {
tmp15959 := Call(__e, PrimFunc(symshen_4unlocked_2), V5239)


if True == tmp15959 {
tmp15917 := Call(__e, PrimFunc(symshen_4lazyderef), V5236, V5238)


let__14202 := tmp15917
_ = let__14202

tmp15957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14202, Nil)
}
__typedArg0 := let__14202
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15957 {
tmp15918 := Call(__e, PrimFunc(symshen_4lazyderef), V5237, V5238)


let__14203 := tmp15918
_ = let__14203

tmp15919 := MakeNative(func(__e *ControlFlow) {
Z5257 := __e.Get(1)
_ = Z5257
tmp15920 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15920

__e.TailApply(PrimFunc(symthaw), V5241)
return


}, 1)

let__14204 := tmp15919
_ = let__14204

tmp15955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14203)
}
__typedArg0 := let__14203
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15955 {
tmp15921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14203)
}
__typedArg0 := let__14203
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15922 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15921, V5238)


let__14205 := tmp15922
_ = let__14205

tmp15923 := MakeNative(func(__e *ControlFlow) {
tmp15924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14203)
}
__typedArg0 := let__14203
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15925 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15924, V5238)


let__14207 := tmp15925
_ = let__14207

tmp15926 := MakeNative(func(__e *ControlFlow) {
Z5262 := __e.Get(1)
_ = Z5262
__e.TailApply(let__14204, Z5262)
return
}, 1)

let__14208 := tmp15926
_ = let__14208

tmp15942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14207)
}
__typedArg0 := let__14207
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15942 {
tmp15927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14207)
}
__typedArg0 := let__14207
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14209 := tmp15927
_ = let__14209

tmp15928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14207)
}
__typedArg0 := let__14207
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15929 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15928, V5238)


let__14210 := tmp15929
_ = let__14210

tmp15930 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14208, let__14209)
return
}, 0)

let__14211 := tmp15930
_ = let__14211

tmp15934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14210, Nil)
}
__typedArg0 := let__14210
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15934 {
__e.TailApply(PrimFunc(symthaw), let__14211)
return
} else {
tmp15932 := Call(__e, PrimFunc(symshen_4pvar_2), let__14210)


if True == tmp15932 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14210, Nil, V5238, let__14211)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15940 := Call(__e, PrimFunc(symshen_4pvar_2), let__14207)


if True == tmp15940 {
tmp15935 := Call(__e, PrimFunc(symshen_4newpv), V5238)


let__14212 := tmp15935
_ = let__14212

tmp15936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14212, Nil)
}
__typedArg0 := let__14212
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15937 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14208, let__14212)
return
}, 0)

tmp15938 := Call(__e, PrimFunc(symshen_4bind_b), let__14207, tmp15936, V5238, tmp15937)


__e.TailApply(PrimFunc(symshen_4gc), V5238, tmp15938)
return


} else {
__e.Return(False)
return
}


}


}, 0)

let__14206 := tmp15923
_ = let__14206

tmp15946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14205, symlist)
}
__typedArg0 := let__14205
__typedArg1 := symlist
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15946 {
__e.TailApply(PrimFunc(symthaw), let__14206)
return
} else {
tmp15944 := Call(__e, PrimFunc(symshen_4pvar_2), let__14205)


if True == tmp15944 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14205, symlist, V5238, let__14206)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15953 := Call(__e, PrimFunc(symshen_4pvar_2), let__14203)


if True == tmp15953 {
tmp15947 := Call(__e, PrimFunc(symshen_4newpv), V5238)


let__14213 := tmp15947
_ = let__14213

tmp15948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14213, Nil)
}
__typedArg0 := let__14213
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp15948)
}
__typedArg0 := symlist
__typedArg1 := tmp15948
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15950 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14204, let__14213)
return
}, 0)

tmp15951 := Call(__e, PrimFunc(symshen_4bind_b), let__14203, tmp15949, V5238, tmp15950)


__e.TailApply(PrimFunc(symshen_4gc), V5238, tmp15951)
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
__e.Return(let__14199)
return
}


} else {
__e.Return(let__14196)
return
}


} else {
__e.Return(let__14193)
return
}


} else {
__e.Return(let__14190)
return
}


}, 6)

tmp15968 := Call(__e, ns2_1set, symshen_4primitive, tmp15864)


_ = tmp15968

tmp15969 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
__self6 := __e.Get(6)
__self7 := __e.Get(7)
__selftop:
V5268 := __self1
_ = V5268
V5269 := __self2
_ = V5269
V5270 := __self3
_ = V5270
V5271 := __self4
_ = V5271
V5272 := __self5
_ = V5272
V5273 := __self6
_ = V5273
V5274 := __self7
_ = V5274
tmp16003 := Call(__e, PrimFunc(symshen_4unlocked_2), V5272)


var ifres15970 Obj

if True == tmp16003 {
tmp15971 := Call(__e, PrimFunc(symshen_4lazyderef), V5270, V5271)


let__14215 := tmp15971
_ = let__14215

tmp16002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14215)
}
__typedArg0 := let__14215
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15972 Obj

if True == tmp16002 {
tmp15973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14215)
}
__typedArg0 := let__14215
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15974 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15973, V5271)


let__14216 := tmp15974
_ = let__14216

tmp16001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14216)
}
__typedArg0 := let__14216
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15975 Obj

if True == tmp16001 {
tmp15976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14216)
}
__typedArg0 := let__14216
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14217 := tmp15976
_ = let__14217

tmp15977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14216)
}
__typedArg0 := let__14216
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15978 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15977, V5271)


let__14218 := tmp15978
_ = let__14218

tmp16000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14218)
}
__typedArg0 := let__14218
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15979 Obj

if True == tmp16000 {
tmp15980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14218)
}
__typedArg0 := let__14218
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14219 := tmp15980
_ = let__14219

tmp15981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14218)
}
__typedArg0 := let__14218
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15982 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15981, V5271)


let__14220 := tmp15982
_ = let__14220

tmp15999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14220)
}
__typedArg0 := let__14220
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15983 Obj

if True == tmp15999 {
tmp15984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14220)
}
__typedArg0 := let__14220
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14221 := tmp15984
_ = let__14221

tmp15985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14220)
}
__typedArg0 := let__14220
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15986 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15985, V5271)


let__14222 := tmp15986
_ = let__14222

tmp15998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14222, Nil)
}
__typedArg0 := let__14222
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15987 Obj

if True == tmp15998 {
tmp15988 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15988

tmp15989 := Call(__e, PrimFunc(symshen_4deref), let__14219, V5271)


tmp15990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp15991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp15989, tmp15990)
}
__typedArg0 := tmp15989
__typedArg1 := tmp15990
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp15992 := MakeNative(func(__e *ControlFlow) {
tmp15993 := Call(__e, PrimFunc(symshen_4deref), V5268, V5271)


tmp15994 := Call(__e, PrimFunc(symshen_4deref), let__14217, V5271)


tmp15995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp15993, tmp15994)
}
__typedArg0 := tmp15993
__typedArg1 := tmp15994
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp15996 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symis_b), V5269, let__14221, V5271, V5272, V5273, V5274)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15995, V5271, V5272, V5273, tmp15996)
return


}, 0)

tmp15997 := Call(__e, PrimFunc(symwhen), tmp15991, V5271, V5272, V5273, tmp15992)


ifres15987 = tmp15997


} else {
ifres15987 = False


}

ifres15983 = ifres15987


} else {
ifres15983 = False


}

ifres15979 = ifres15983


} else {
ifres15979 = False


}

ifres15975 = ifres15979


} else {
ifres15975 = False


}

ifres15972 = ifres15975


} else {
ifres15972 = False


}

ifres15970 = ifres15972


} else {
ifres15970 = False


}

let__14214 := ifres15970
_ = let__14214

tmp16012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14214, False)
}
__typedArg0 := let__14214
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16012 {
tmp16010 := Call(__e, PrimFunc(symshen_4unlocked_2), V5272)


if True == tmp16010 {
tmp16004 := Call(__e, PrimFunc(symshen_4lazyderef), V5270, V5271)


let__14223 := tmp16004
_ = let__14223

tmp16008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14223)
}
__typedArg0 := let__14223
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16008 {
tmp16005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14223)
}
__typedArg0 := let__14223
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14224 := tmp16005
_ = let__14224

tmp16006 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16006

if PrimFunc(symshen_4by_1hypothesis) == __self {
__self1, __self2, __self3, __self4, __self5, __self6, __self7 = V5268, V5269, let__14224, V5271, V5272, V5273, V5274
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4by_1hypothesis), V5268, V5269, let__14224, V5271, V5272, V5273, V5274)
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
__e.Return(let__14214)
return
}


}, 7)

tmp16013 := Call(__e, ns2_1set, symshen_4by_1hypothesis, tmp15969)


_ = tmp16013

tmp16014 := MakeNative(func(__e *ControlFlow) {
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
tmp16019 := Call(__e, PrimFunc(symshen_4unlocked_2), V5289)


if True == tmp16019 {
tmp16015 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16015

tmp16016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp16017 := Call(__e, PrimFunc(symassoc), V5286, tmp16016)


__e.TailApply(PrimFunc(symshen_4sigf), tmp16017, V5287, V5288, V5289, V5290, V5291)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp16020 := Call(__e, ns2_1set, symshen_4lookupsig, tmp16014)


_ = tmp16020

tmp16021 := MakeNative(func(__e *ControlFlow) {
V5306 := __e.Get(1)
_ = V5306
V5307 := __e.Get(2)
_ = V5307
V5308 := __e.Get(3)
_ = V5308
V5309 := __e.Get(4)
_ = V5309
V5310 := __e.Get(5)
_ = V5310
V5311 := __e.Get(6)
_ = V5311
tmp16028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5306)
}
__typedArg0 := V5306
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16028 {
tmp16022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5306)
}
__typedArg0 := V5306
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16023 := Call(__e, tmp16022, V5307)


tmp16024 := Call(__e, tmp16023, V5308)


tmp16025 := Call(__e, tmp16024, V5309)


tmp16026 := Call(__e, tmp16025, V5310)


__e.TailApply(tmp16026, V5311)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp16029 := Call(__e, ns2_1set, symshen_4sigf, tmp16021)


_ = tmp16029

tmp16030 := MakeNative(func(__e *ControlFlow) {
V5312 := __e.Get(1)
_ = V5312
tmp16031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(MakeInteger(3))
}
__typedArg0 := MakeInteger(3)
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

let__14225 := tmp16031
_ = let__14225

tmp16032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__14225, MakeInteger(0), symshen_4print_1freshterm)
}
__typedArg0 := let__14225
__typedArg1 := MakeInteger(0)
__typedArg2 := symshen_4print_1freshterm
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__14226 := tmp16032
_ = let__14226

tmp16033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__14226, MakeInteger(1), V5312)
}
__typedArg0 := let__14226
__typedArg1 := MakeInteger(1)
__typedArg2 := V5312
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__14227 := tmp16033
_ = let__14227

tmp16034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dgensym_d)
}
__typedArg0 := symshen_4_dgensym_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp16036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dgensym_d, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp16034)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp16034
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := symshen_4_dgensym_d
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp16034)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp16034
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

tmp16037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__14227, MakeInteger(2), tmp16036)
}
__typedArg0 := let__14227
__typedArg1 := MakeInteger(2)
__typedArg2 := tmp16036
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__14228 := tmp16037
_ = let__14228

__e.Return(let__14228)
return


}, 1)

tmp16038 := Call(__e, ns2_1set, symshen_4freshterm, tmp16030)


_ = tmp16038

tmp16039 := MakeNative(func(__e *ControlFlow) {
V5317 := __e.Get(1)
_ = V5317
tmp16040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V5317, MakeInteger(1))
}
__typedArg0 := V5317
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp16041 := Call(__e, PrimFunc(symshen_4app), tmp16040, MakeString(""), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("&&"))
__typedS1, __typedOK1 := TypedString(tmp16041)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("&&")
__typedArg1 := tmp16041
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp16042 := Call(__e, ns2_1set, symshen_4print_1freshterm, tmp16039)


_ = tmp16042

tmp16043 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
__self6 := __e.Get(6)
__self7 := __e.Get(7)
__selftop:
V5318 := __self1
_ = V5318
V5319 := __self2
_ = V5319
V5320 := __self3
_ = V5320
V5321 := __self4
_ = V5321
V5322 := __self5
_ = V5322
V5323 := __self6
_ = V5323
V5324 := __self7
_ = V5324
tmp16060 := Call(__e, PrimFunc(symshen_4unlocked_2), V5322)


var ifres16044 Obj

if True == tmp16060 {
tmp16045 := Call(__e, PrimFunc(symshen_4lazyderef), V5320, V5321)


let__14230 := tmp16045
_ = let__14230

tmp16059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14230)
}
__typedArg0 := let__14230
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16046 Obj

if True == tmp16059 {
tmp16047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14230)
}
__typedArg0 := let__14230
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16048 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16047, V5321)


let__14231 := tmp16048
_ = let__14231

tmp16058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14231)
}
__typedArg0 := let__14231
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16049 Obj

if True == tmp16058 {
tmp16050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14231)
}
__typedArg0 := let__14231
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14232 := tmp16050
_ = let__14232

tmp16051 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16051

tmp16052 := Call(__e, PrimFunc(symshen_4deref), let__14232, V5321)


tmp16053 := Call(__e, PrimFunc(symshen_4deref), V5318, V5321)


tmp16054 := Call(__e, tmp16052, tmp16053)


tmp16055 := Call(__e, PrimFunc(symshen_4deref), V5319, V5321)


tmp16056 := Call(__e, tmp16054, tmp16055)


tmp16057 := Call(__e, PrimFunc(symcall), tmp16056, V5321, V5322, V5323, V5324)


ifres16049 = tmp16057


} else {
ifres16049 = False


}

ifres16046 = ifres16049


} else {
ifres16046 = False


}

ifres16044 = ifres16046


} else {
ifres16044 = False


}

let__14229 := ifres16044
_ = let__14229

tmp16069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14229, False)
}
__typedArg0 := let__14229
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16069 {
tmp16067 := Call(__e, PrimFunc(symshen_4unlocked_2), V5322)


if True == tmp16067 {
tmp16061 := Call(__e, PrimFunc(symshen_4lazyderef), V5320, V5321)


let__14233 := tmp16061
_ = let__14233

tmp16065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14233)
}
__typedArg0 := let__14233
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16065 {
tmp16062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14233)
}
__typedArg0 := let__14233
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14234 := tmp16062
_ = let__14234

tmp16063 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16063

if PrimFunc(symshen_4search_1user_1datatypes) == __self {
__self1, __self2, __self3, __self4, __self5, __self6, __self7 = V5318, V5319, let__14234, V5321, V5322, V5323, V5324
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4search_1user_1datatypes), V5318, V5319, let__14234, V5321, V5322, V5323, V5324)
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
__e.Return(let__14229)
return
}


}, 7)

tmp16070 := Call(__e, ns2_1set, symshen_4search_1user_1datatypes, tmp16043)


_ = tmp16070

tmp16071 := MakeNative(func(__e *ControlFlow) {
V5331 := __e.Get(1)
_ = V5331
V5332 := __e.Get(2)
_ = V5332
V5333 := __e.Get(3)
_ = V5333
V5334 := __e.Get(4)
_ = V5334
V5335 := __e.Get(5)
_ = V5335
V5336 := __e.Get(6)
_ = V5336
V5337 := __e.Get(7)
_ = V5337
let__14235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5336)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5336
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__14235

tmp16083 := Call(__e, PrimFunc(symshen_4unlocked_2), V5335)


var ifres16073 Obj

if True == tmp16083 {
tmp16074 := Call(__e, PrimFunc(symshen_4lazyderef), V5331, V5334)


let__14237 := tmp16074
_ = let__14237

tmp16082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14237, Nil)
}
__typedArg0 := let__14237
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16075 Obj

if True == tmp16082 {
tmp16076 := Call(__e, PrimFunc(symshen_4lazyderef), V5333, V5334)


let__14238 := tmp16076
_ = let__14238

tmp16081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14238, True)
}
__typedArg0 := let__14238
__typedArg1 := True
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16077 Obj

if True == tmp16081 {
tmp16078 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16078

tmp16079 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symbind), V5332, Nil, V5334, V5335, let__14235, V5337)
return
}, 0)

tmp16080 := Call(__e, PrimFunc(symshen_4cut), V5334, V5335, let__14235, tmp16079)


ifres16077 = tmp16080


} else {
ifres16077 = False


}

ifres16075 = ifres16077


} else {
ifres16075 = False


}

ifres16073 = ifres16075


} else {
ifres16073 = False


}

let__14236 := ifres16073
_ = let__14236

tmp16426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14236, False)
}
__typedArg0 := let__14236
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16426 {
tmp16162 := Call(__e, PrimFunc(symshen_4unlocked_2), V5335)


var ifres16084 Obj

if True == tmp16162 {
tmp16085 := Call(__e, PrimFunc(symshen_4lazyderef), V5331, V5334)


let__14240 := tmp16085
_ = let__14240

tmp16161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14240)
}
__typedArg0 := let__14240
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16086 Obj

if True == tmp16161 {
tmp16087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14240)
}
__typedArg0 := let__14240
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16088 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16087, V5334)


let__14241 := tmp16088
_ = let__14241

tmp16160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14241)
}
__typedArg0 := let__14241
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16089 Obj

if True == tmp16160 {
tmp16090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14241)
}
__typedArg0 := let__14241
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16091 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16090, V5334)


let__14242 := tmp16091
_ = let__14242

tmp16159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14242)
}
__typedArg0 := let__14242
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16092 Obj

if True == tmp16159 {
tmp16093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14242)
}
__typedArg0 := let__14242
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16094 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16093, V5334)


let__14243 := tmp16094
_ = let__14243

tmp16158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14243, symcons)
}
__typedArg0 := let__14243
__typedArg1 := symcons
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16095 Obj

if True == tmp16158 {
tmp16096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14242)
}
__typedArg0 := let__14242
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16097 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16096, V5334)


let__14244 := tmp16097
_ = let__14244

tmp16157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14244)
}
__typedArg0 := let__14244
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16098 Obj

if True == tmp16157 {
tmp16099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14244)
}
__typedArg0 := let__14244
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14245 := tmp16099
_ = let__14245

tmp16100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14244)
}
__typedArg0 := let__14244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16101 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16100, V5334)


let__14246 := tmp16101
_ = let__14246

tmp16156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14246)
}
__typedArg0 := let__14246
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16102 Obj

if True == tmp16156 {
tmp16103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14246)
}
__typedArg0 := let__14246
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14247 := tmp16103
_ = let__14247

tmp16104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14246)
}
__typedArg0 := let__14246
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16105 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16104, V5334)


let__14248 := tmp16105
_ = let__14248

tmp16155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14248, Nil)
}
__typedArg0 := let__14248
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16106 Obj

if True == tmp16155 {
tmp16107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14241)
}
__typedArg0 := let__14241
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16108 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16107, V5334)


let__14249 := tmp16108
_ = let__14249

tmp16154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14249)
}
__typedArg0 := let__14249
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16109 Obj

if True == tmp16154 {
tmp16110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14249)
}
__typedArg0 := let__14249
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14250 := tmp16110
_ = let__14250

tmp16111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14249)
}
__typedArg0 := let__14249
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16112 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16111, V5334)


let__14251 := tmp16112
_ = let__14251

tmp16153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14251)
}
__typedArg0 := let__14251
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16113 Obj

if True == tmp16153 {
tmp16114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14251)
}
__typedArg0 := let__14251
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16115 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16114, V5334)


let__14252 := tmp16115
_ = let__14252

tmp16152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14252)
}
__typedArg0 := let__14252
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16116 Obj

if True == tmp16152 {
tmp16117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14252)
}
__typedArg0 := let__14252
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16118 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16117, V5334)


let__14253 := tmp16118
_ = let__14253

tmp16151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14253, symlist)
}
__typedArg0 := let__14253
__typedArg1 := symlist
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16119 Obj

if True == tmp16151 {
tmp16120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14252)
}
__typedArg0 := let__14252
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16121 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16120, V5334)


let__14254 := tmp16121
_ = let__14254

tmp16150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14254)
}
__typedArg0 := let__14254
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16122 Obj

if True == tmp16150 {
tmp16123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14254)
}
__typedArg0 := let__14254
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14255 := tmp16123
_ = let__14255

tmp16124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14254)
}
__typedArg0 := let__14254
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16125 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16124, V5334)


let__14256 := tmp16125
_ = let__14256

tmp16149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14256, Nil)
}
__typedArg0 := let__14256
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16126 Obj

if True == tmp16149 {
tmp16127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14251)
}
__typedArg0 := let__14251
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16128 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16127, V5334)


let__14257 := tmp16128
_ = let__14257

tmp16148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14257, Nil)
}
__typedArg0 := let__14257
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16129 Obj

if True == tmp16148 {
tmp16130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14240)
}
__typedArg0 := let__14240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14258 := tmp16130
_ = let__14258

tmp16131 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16131

tmp16132 := Call(__e, PrimFunc(symshen_4deref), let__14250, V5334)


tmp16133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16132, tmp16133)
}
__typedArg0 := tmp16132
__typedArg1 := tmp16133
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16135 := MakeNative(func(__e *ControlFlow) {
tmp16136 := MakeNative(func(__e *ControlFlow) {
tmp16137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14255, Nil)
}
__typedArg0 := let__14255
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14250, tmp16137)
}
__typedArg0 := let__14250
__typedArg1 := tmp16137
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14245, tmp16138)
}
__typedArg0 := let__14245
__typedArg1 := tmp16138
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14255, Nil)
}
__typedArg0 := let__14255
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp16140)
}
__typedArg0 := symlist
__typedArg1 := tmp16140
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16141, Nil)
}
__typedArg0 := tmp16141
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14250, tmp16142)
}
__typedArg0 := let__14250
__typedArg1 := tmp16142
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14247, tmp16143)
}
__typedArg0 := let__14247
__typedArg1 := tmp16143
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16144, let__14258)
}
__typedArg0 := tmp16144
__typedArg1 := let__14258
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16139, tmp16145)
}
__typedArg0 := tmp16139
__typedArg1 := tmp16145
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16146, V5332, True, V5334, V5335, let__14235, V5337)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5334, V5335, let__14235, tmp16136)
return


}, 0)

tmp16147 := Call(__e, PrimFunc(symwhen), tmp16134, V5334, V5335, let__14235, tmp16135)


ifres16129 = tmp16147


} else {
ifres16129 = False


}

ifres16126 = ifres16129


} else {
ifres16126 = False


}

ifres16122 = ifres16126


} else {
ifres16122 = False


}

ifres16119 = ifres16122


} else {
ifres16119 = False


}

ifres16116 = ifres16119


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

ifres16106 = ifres16109


} else {
ifres16106 = False


}

ifres16102 = ifres16106


} else {
ifres16102 = False


}

ifres16098 = ifres16102


} else {
ifres16098 = False


}

ifres16095 = ifres16098


} else {
ifres16095 = False


}

ifres16092 = ifres16095


} else {
ifres16092 = False


}

ifres16089 = ifres16092


} else {
ifres16089 = False


}

ifres16086 = ifres16089


} else {
ifres16086 = False


}

ifres16084 = ifres16086


} else {
ifres16084 = False


}

let__14239 := ifres16084
_ = let__14239

tmp16424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14239, False)
}
__typedArg0 := let__14239
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16424 {
tmp16244 := Call(__e, PrimFunc(symshen_4unlocked_2), V5335)


var ifres16163 Obj

if True == tmp16244 {
tmp16164 := Call(__e, PrimFunc(symshen_4lazyderef), V5331, V5334)


let__14260 := tmp16164
_ = let__14260

tmp16243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14260)
}
__typedArg0 := let__14260
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16165 Obj

if True == tmp16243 {
tmp16166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14260)
}
__typedArg0 := let__14260
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16167 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16166, V5334)


let__14261 := tmp16167
_ = let__14261

tmp16242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14261)
}
__typedArg0 := let__14261
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16168 Obj

if True == tmp16242 {
tmp16169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14261)
}
__typedArg0 := let__14261
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16170 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16169, V5334)


let__14262 := tmp16170
_ = let__14262

tmp16241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14262)
}
__typedArg0 := let__14262
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16171 Obj

if True == tmp16241 {
tmp16172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14262)
}
__typedArg0 := let__14262
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16173 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16172, V5334)


let__14263 := tmp16173
_ = let__14263

tmp16240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14263, sym_8p)
}
__typedArg0 := let__14263
__typedArg1 := sym_8p
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16174 Obj

if True == tmp16240 {
tmp16175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14262)
}
__typedArg0 := let__14262
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16176 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16175, V5334)


let__14264 := tmp16176
_ = let__14264

tmp16239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14264)
}
__typedArg0 := let__14264
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16177 Obj

if True == tmp16239 {
tmp16178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14264)
}
__typedArg0 := let__14264
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14265 := tmp16178
_ = let__14265

tmp16179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14264)
}
__typedArg0 := let__14264
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16180 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16179, V5334)


let__14266 := tmp16180
_ = let__14266

tmp16238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14266)
}
__typedArg0 := let__14266
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16181 Obj

if True == tmp16238 {
tmp16182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14266)
}
__typedArg0 := let__14266
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14267 := tmp16182
_ = let__14267

tmp16183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14266)
}
__typedArg0 := let__14266
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16184 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16183, V5334)


let__14268 := tmp16184
_ = let__14268

tmp16237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14268, Nil)
}
__typedArg0 := let__14268
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16185 Obj

if True == tmp16237 {
tmp16186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14261)
}
__typedArg0 := let__14261
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16187 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16186, V5334)


let__14269 := tmp16187
_ = let__14269

tmp16236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14269)
}
__typedArg0 := let__14269
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16188 Obj

if True == tmp16236 {
tmp16189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14269)
}
__typedArg0 := let__14269
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14270 := tmp16189
_ = let__14270

tmp16190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14269)
}
__typedArg0 := let__14269
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16191 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16190, V5334)


let__14271 := tmp16191
_ = let__14271

tmp16235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14271)
}
__typedArg0 := let__14271
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16192 Obj

if True == tmp16235 {
tmp16193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14271)
}
__typedArg0 := let__14271
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16194 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16193, V5334)


let__14272 := tmp16194
_ = let__14272

tmp16234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14272)
}
__typedArg0 := let__14272
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16195 Obj

if True == tmp16234 {
tmp16196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14272)
}
__typedArg0 := let__14272
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14273 := tmp16196
_ = let__14273

tmp16197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14272)
}
__typedArg0 := let__14272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16198 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16197, V5334)


let__14274 := tmp16198
_ = let__14274

tmp16233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14274)
}
__typedArg0 := let__14274
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16199 Obj

if True == tmp16233 {
tmp16200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14274)
}
__typedArg0 := let__14274
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16201 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16200, V5334)


let__14275 := tmp16201
_ = let__14275

tmp16232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14275, sym_d)
}
__typedArg0 := let__14275
__typedArg1 := sym_d
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16202 Obj

if True == tmp16232 {
tmp16203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14274)
}
__typedArg0 := let__14274
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16204 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16203, V5334)


let__14276 := tmp16204
_ = let__14276

tmp16231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14276)
}
__typedArg0 := let__14276
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16205 Obj

if True == tmp16231 {
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

tmp16208 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16207, V5334)


let__14278 := tmp16208
_ = let__14278

tmp16230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14278, Nil)
}
__typedArg0 := let__14278
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16209 Obj

if True == tmp16230 {
tmp16210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14271)
}
__typedArg0 := let__14271
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16211 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16210, V5334)


let__14279 := tmp16211
_ = let__14279

tmp16229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14279, Nil)
}
__typedArg0 := let__14279
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16212 Obj

if True == tmp16229 {
tmp16213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14260)
}
__typedArg0 := let__14260
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14280 := tmp16213
_ = let__14280

tmp16214 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16214

tmp16215 := Call(__e, PrimFunc(symshen_4deref), let__14270, V5334)


tmp16216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16215, tmp16216)
}
__typedArg0 := tmp16215
__typedArg1 := tmp16216
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16218 := MakeNative(func(__e *ControlFlow) {
tmp16219 := MakeNative(func(__e *ControlFlow) {
tmp16220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14273, Nil)
}
__typedArg0 := let__14273
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14270, tmp16220)
}
__typedArg0 := let__14270
__typedArg1 := tmp16220
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14265, tmp16221)
}
__typedArg0 := let__14265
__typedArg1 := tmp16221
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14277, Nil)
}
__typedArg0 := let__14277
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14270, tmp16223)
}
__typedArg0 := let__14270
__typedArg1 := tmp16223
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14267, tmp16224)
}
__typedArg0 := let__14267
__typedArg1 := tmp16224
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16225, let__14280)
}
__typedArg0 := tmp16225
__typedArg1 := let__14280
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16222, tmp16226)
}
__typedArg0 := tmp16222
__typedArg1 := tmp16226
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16227, V5332, True, V5334, V5335, let__14235, V5337)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5334, V5335, let__14235, tmp16219)
return


}, 0)

tmp16228 := Call(__e, PrimFunc(symwhen), tmp16217, V5334, V5335, let__14235, tmp16218)


ifres16212 = tmp16228


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

ifres16185 = ifres16188


} else {
ifres16185 = False


}

ifres16181 = ifres16185


} else {
ifres16181 = False


}

ifres16177 = ifres16181


} else {
ifres16177 = False


}

ifres16174 = ifres16177


} else {
ifres16174 = False


}

ifres16171 = ifres16174


} else {
ifres16171 = False


}

ifres16168 = ifres16171


} else {
ifres16168 = False


}

ifres16165 = ifres16168


} else {
ifres16165 = False


}

ifres16163 = ifres16165


} else {
ifres16163 = False


}

let__14259 := ifres16163
_ = let__14259

tmp16422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14259, False)
}
__typedArg0 := let__14259
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16422 {
tmp16308 := Call(__e, PrimFunc(symshen_4unlocked_2), V5335)


var ifres16245 Obj

if True == tmp16308 {
tmp16246 := Call(__e, PrimFunc(symshen_4lazyderef), V5331, V5334)


let__14282 := tmp16246
_ = let__14282

tmp16307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14282)
}
__typedArg0 := let__14282
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16247 Obj

if True == tmp16307 {
tmp16248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14282)
}
__typedArg0 := let__14282
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16249 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16248, V5334)


let__14283 := tmp16249
_ = let__14283

tmp16306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14283)
}
__typedArg0 := let__14283
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16250 Obj

if True == tmp16306 {
tmp16251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14283)
}
__typedArg0 := let__14283
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16252 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16251, V5334)


let__14284 := tmp16252
_ = let__14284

tmp16305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14284)
}
__typedArg0 := let__14284
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16253 Obj

if True == tmp16305 {
tmp16254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14284)
}
__typedArg0 := let__14284
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16255 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16254, V5334)


let__14285 := tmp16255
_ = let__14285

tmp16304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14285, sym_8s)
}
__typedArg0 := let__14285
__typedArg1 := sym_8s
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16256 Obj

if True == tmp16304 {
tmp16257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14284)
}
__typedArg0 := let__14284
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16258 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16257, V5334)


let__14286 := tmp16258
_ = let__14286

tmp16303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14286)
}
__typedArg0 := let__14286
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16259 Obj

if True == tmp16303 {
tmp16260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14286)
}
__typedArg0 := let__14286
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14287 := tmp16260
_ = let__14287

tmp16261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14286)
}
__typedArg0 := let__14286
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16262 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16261, V5334)


let__14288 := tmp16262
_ = let__14288

tmp16302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14288)
}
__typedArg0 := let__14288
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16263 Obj

if True == tmp16302 {
tmp16264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14288)
}
__typedArg0 := let__14288
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14289 := tmp16264
_ = let__14289

tmp16265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14288)
}
__typedArg0 := let__14288
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16266 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16265, V5334)


let__14290 := tmp16266
_ = let__14290

tmp16301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14290, Nil)
}
__typedArg0 := let__14290
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16267 Obj

if True == tmp16301 {
tmp16268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14283)
}
__typedArg0 := let__14283
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16269 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16268, V5334)


let__14291 := tmp16269
_ = let__14291

tmp16300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14291)
}
__typedArg0 := let__14291
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16270 Obj

if True == tmp16300 {
tmp16271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14291)
}
__typedArg0 := let__14291
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14292 := tmp16271
_ = let__14292

tmp16272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14291)
}
__typedArg0 := let__14291
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16273 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16272, V5334)


let__14293 := tmp16273
_ = let__14293

tmp16299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14293)
}
__typedArg0 := let__14293
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16274 Obj

if True == tmp16299 {
tmp16275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14293)
}
__typedArg0 := let__14293
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16276 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16275, V5334)


let__14294 := tmp16276
_ = let__14294

tmp16298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14294, symstring)
}
__typedArg0 := let__14294
__typedArg1 := symstring
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16277 Obj

if True == tmp16298 {
tmp16278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14293)
}
__typedArg0 := let__14293
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16279 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16278, V5334)


let__14295 := tmp16279
_ = let__14295

tmp16297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14295, Nil)
}
__typedArg0 := let__14295
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16280 Obj

if True == tmp16297 {
tmp16281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14282)
}
__typedArg0 := let__14282
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14296 := tmp16281
_ = let__14296

tmp16282 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16282

tmp16283 := Call(__e, PrimFunc(symshen_4deref), let__14292, V5334)


tmp16284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16283, tmp16284)
}
__typedArg0 := tmp16283
__typedArg1 := tmp16284
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16286 := MakeNative(func(__e *ControlFlow) {
tmp16287 := MakeNative(func(__e *ControlFlow) {
tmp16288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14292, tmp16288)
}
__typedArg0 := let__14292
__typedArg1 := tmp16288
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14287, tmp16289)
}
__typedArg0 := let__14287
__typedArg1 := tmp16289
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14292, tmp16291)
}
__typedArg0 := let__14292
__typedArg1 := tmp16291
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14289, tmp16292)
}
__typedArg0 := let__14289
__typedArg1 := tmp16292
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16293, let__14296)
}
__typedArg0 := tmp16293
__typedArg1 := let__14296
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16290, tmp16294)
}
__typedArg0 := tmp16290
__typedArg1 := tmp16294
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16295, V5332, True, V5334, V5335, let__14235, V5337)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5334, V5335, let__14235, tmp16287)
return


}, 0)

tmp16296 := Call(__e, PrimFunc(symwhen), tmp16285, V5334, V5335, let__14235, tmp16286)


ifres16280 = tmp16296


} else {
ifres16280 = False


}

ifres16277 = ifres16280


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

ifres16267 = ifres16270


} else {
ifres16267 = False


}

ifres16263 = ifres16267


} else {
ifres16263 = False


}

ifres16259 = ifres16263


} else {
ifres16259 = False


}

ifres16256 = ifres16259


} else {
ifres16256 = False


}

ifres16253 = ifres16256


} else {
ifres16253 = False


}

ifres16250 = ifres16253


} else {
ifres16250 = False


}

ifres16247 = ifres16250


} else {
ifres16247 = False


}

ifres16245 = ifres16247


} else {
ifres16245 = False


}

let__14281 := ifres16245
_ = let__14281

tmp16420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14281, False)
}
__typedArg0 := let__14281
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16420 {
tmp16387 := Call(__e, PrimFunc(symshen_4unlocked_2), V5335)


var ifres16309 Obj

if True == tmp16387 {
tmp16310 := Call(__e, PrimFunc(symshen_4lazyderef), V5331, V5334)


let__14298 := tmp16310
_ = let__14298

tmp16386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14298)
}
__typedArg0 := let__14298
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16311 Obj

if True == tmp16386 {
tmp16312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14298)
}
__typedArg0 := let__14298
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16313 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16312, V5334)


let__14299 := tmp16313
_ = let__14299

tmp16385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14299)
}
__typedArg0 := let__14299
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16314 Obj

if True == tmp16385 {
tmp16315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14299)
}
__typedArg0 := let__14299
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16316 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16315, V5334)


let__14300 := tmp16316
_ = let__14300

tmp16384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14300)
}
__typedArg0 := let__14300
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16317 Obj

if True == tmp16384 {
tmp16318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14300)
}
__typedArg0 := let__14300
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16319 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16318, V5334)


let__14301 := tmp16319
_ = let__14301

tmp16383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14301, sym_8v)
}
__typedArg0 := let__14301
__typedArg1 := sym_8v
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16320 Obj

if True == tmp16383 {
tmp16321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14300)
}
__typedArg0 := let__14300
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16322 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16321, V5334)


let__14302 := tmp16322
_ = let__14302

tmp16382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14302)
}
__typedArg0 := let__14302
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16323 Obj

if True == tmp16382 {
tmp16324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14302)
}
__typedArg0 := let__14302
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14303 := tmp16324
_ = let__14303

tmp16325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14302)
}
__typedArg0 := let__14302
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16326 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16325, V5334)


let__14304 := tmp16326
_ = let__14304

tmp16381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14304)
}
__typedArg0 := let__14304
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16327 Obj

if True == tmp16381 {
tmp16328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14304)
}
__typedArg0 := let__14304
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14305 := tmp16328
_ = let__14305

tmp16329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14304)
}
__typedArg0 := let__14304
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16330 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16329, V5334)


let__14306 := tmp16330
_ = let__14306

tmp16380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14306, Nil)
}
__typedArg0 := let__14306
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16331 Obj

if True == tmp16380 {
tmp16332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14299)
}
__typedArg0 := let__14299
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16333 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16332, V5334)


let__14307 := tmp16333
_ = let__14307

tmp16379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14307)
}
__typedArg0 := let__14307
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16334 Obj

if True == tmp16379 {
tmp16335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14307)
}
__typedArg0 := let__14307
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14308 := tmp16335
_ = let__14308

tmp16336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14307)
}
__typedArg0 := let__14307
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16337 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16336, V5334)


let__14309 := tmp16337
_ = let__14309

tmp16378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14309)
}
__typedArg0 := let__14309
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16338 Obj

if True == tmp16378 {
tmp16339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14309)
}
__typedArg0 := let__14309
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16340 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16339, V5334)


let__14310 := tmp16340
_ = let__14310

tmp16377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14310)
}
__typedArg0 := let__14310
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16341 Obj

if True == tmp16377 {
tmp16342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14310)
}
__typedArg0 := let__14310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16343 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16342, V5334)


let__14311 := tmp16343
_ = let__14311

tmp16376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14311, symvector)
}
__typedArg0 := let__14311
__typedArg1 := symvector
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16344 Obj

if True == tmp16376 {
tmp16345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14310)
}
__typedArg0 := let__14310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16346 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16345, V5334)


let__14312 := tmp16346
_ = let__14312

tmp16375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14312)
}
__typedArg0 := let__14312
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16347 Obj

if True == tmp16375 {
tmp16348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14312)
}
__typedArg0 := let__14312
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14313 := tmp16348
_ = let__14313

tmp16349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14312)
}
__typedArg0 := let__14312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16350 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16349, V5334)


let__14314 := tmp16350
_ = let__14314

tmp16374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14314, Nil)
}
__typedArg0 := let__14314
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16351 Obj

if True == tmp16374 {
tmp16352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14309)
}
__typedArg0 := let__14309
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16353 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16352, V5334)


let__14315 := tmp16353
_ = let__14315

tmp16373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14315, Nil)
}
__typedArg0 := let__14315
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16354 Obj

if True == tmp16373 {
tmp16355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14298)
}
__typedArg0 := let__14298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14316 := tmp16355
_ = let__14316

tmp16356 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16356

tmp16357 := Call(__e, PrimFunc(symshen_4deref), let__14308, V5334)


tmp16358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16357, tmp16358)
}
__typedArg0 := tmp16357
__typedArg1 := tmp16358
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16360 := MakeNative(func(__e *ControlFlow) {
tmp16361 := MakeNative(func(__e *ControlFlow) {
tmp16362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14313, Nil)
}
__typedArg0 := let__14313
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14308, tmp16362)
}
__typedArg0 := let__14308
__typedArg1 := tmp16362
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14303, tmp16363)
}
__typedArg0 := let__14303
__typedArg1 := tmp16363
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14313, Nil)
}
__typedArg0 := let__14313
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp16365)
}
__typedArg0 := symvector
__typedArg1 := tmp16365
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16366, Nil)
}
__typedArg0 := tmp16366
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14308, tmp16367)
}
__typedArg0 := let__14308
__typedArg1 := tmp16367
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14305, tmp16368)
}
__typedArg0 := let__14305
__typedArg1 := tmp16368
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16369, let__14316)
}
__typedArg0 := tmp16369
__typedArg1 := let__14316
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16364, tmp16370)
}
__typedArg0 := tmp16364
__typedArg1 := tmp16370
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16371, V5332, True, V5334, V5335, let__14235, V5337)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5334, V5335, let__14235, tmp16361)
return


}, 0)

tmp16372 := Call(__e, PrimFunc(symwhen), tmp16359, V5334, V5335, let__14235, tmp16360)


ifres16354 = tmp16372


} else {
ifres16354 = False


}

ifres16351 = ifres16354


} else {
ifres16351 = False


}

ifres16347 = ifres16351


} else {
ifres16347 = False


}

ifres16344 = ifres16347


} else {
ifres16344 = False


}

ifres16341 = ifres16344


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

ifres16331 = ifres16334


} else {
ifres16331 = False


}

ifres16327 = ifres16331


} else {
ifres16327 = False


}

ifres16323 = ifres16327


} else {
ifres16323 = False


}

ifres16320 = ifres16323


} else {
ifres16320 = False


}

ifres16317 = ifres16320


} else {
ifres16317 = False


}

ifres16314 = ifres16317


} else {
ifres16314 = False


}

ifres16311 = ifres16314


} else {
ifres16311 = False


}

ifres16309 = ifres16311


} else {
ifres16309 = False


}

let__14297 := ifres16309
_ = let__14297

tmp16418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14297, False)
}
__typedArg0 := let__14297
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16418 {
tmp16414 := Call(__e, PrimFunc(symshen_4unlocked_2), V5335)


var ifres16388 Obj

if True == tmp16414 {
tmp16389 := Call(__e, PrimFunc(symshen_4lazyderef), V5331, V5334)


let__14318 := tmp16389
_ = let__14318

tmp16413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14318)
}
__typedArg0 := let__14318
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16390 Obj

if True == tmp16413 {
tmp16391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14318)
}
__typedArg0 := let__14318
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14319 := tmp16391
_ = let__14319

tmp16392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14318)
}
__typedArg0 := let__14318
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14320 := tmp16392
_ = let__14320

tmp16393 := Call(__e, PrimFunc(symshen_4lazyderef), V5332, V5334)


let__14321 := tmp16393
_ = let__14321

tmp16394 := MakeNative(func(__e *ControlFlow) {
Z5426 := __e.Get(1)
_ = Z5426
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5427 := __e.Get(1)
_ = Z5427
tmp16395 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16395

tmp16396 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4l_1rules), let__14320, Z5427, V5333, V5334, V5335, let__14235, V5337)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5426, let__14319, V5334, V5335, let__14235, tmp16396)
return


}, 1))
return
}, 1)

let__14322 := tmp16394
_ = let__14322

tmp16412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14321)
}
__typedArg0 := let__14321
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16397 Obj

if True == tmp16412 {
tmp16398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14321)
}
__typedArg0 := let__14321
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14323 := tmp16398
_ = let__14323

tmp16399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14321)
}
__typedArg0 := let__14321
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14324 := tmp16399
_ = let__14324

tmp16400 := Call(__e, let__14322, let__14323)


tmp16401 := Call(__e, tmp16400, let__14324)


ifres16397 = tmp16401


} else {
tmp16411 := Call(__e, PrimFunc(symshen_4pvar_2), let__14321)


var ifres16402 Obj

if True == tmp16411 {
tmp16403 := Call(__e, PrimFunc(symshen_4newpv), V5334)


let__14325 := tmp16403
_ = let__14325

tmp16404 := Call(__e, PrimFunc(symshen_4newpv), V5334)


let__14326 := tmp16404
_ = let__14326

tmp16405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14325, let__14326)
}
__typedArg0 := let__14325
__typedArg1 := let__14326
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16406 := MakeNative(func(__e *ControlFlow) {
tmp16407 := Call(__e, let__14322, let__14325)


__e.TailApply(tmp16407, let__14326)
return


}, 0)

tmp16408 := Call(__e, PrimFunc(symshen_4bind_b), let__14321, tmp16405, V5334, tmp16406)


tmp16409 := Call(__e, PrimFunc(symshen_4gc), V5334, tmp16408)


tmp16410 := Call(__e, PrimFunc(symshen_4gc), V5334, tmp16409)


ifres16402 = tmp16410


} else {
ifres16402 = False


}

ifres16397 = ifres16402


}

ifres16390 = ifres16397


} else {
ifres16390 = False


}

ifres16388 = ifres16390


} else {
ifres16388 = False


}

let__14317 := ifres16388
_ = let__14317

tmp16416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14317, False)
}
__typedArg0 := let__14317
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16416 {
__e.TailApply(PrimFunc(symshen_4unlock), V5335, let__14235)
return
} else {
__e.Return(let__14317)
return
}


} else {
__e.Return(let__14297)
return
}


} else {
__e.Return(let__14281)
return
}


} else {
__e.Return(let__14259)
return
}


} else {
__e.Return(let__14239)
return
}


} else {
__e.Return(let__14236)
return
}


}, 7)

tmp16427 := Call(__e, ns2_1set, symshen_4l_1rules, tmp16071)


_ = tmp16427

tmp16428 := MakeNative(func(__e *ControlFlow) {
V5432 := __e.Get(1)
_ = V5432
V5433 := __e.Get(2)
_ = V5433
V5434 := __e.Get(3)
_ = V5434
V5435 := __e.Get(4)
_ = V5435
V5436 := __e.Get(5)
_ = V5436
V5437 := __e.Get(6)
_ = V5437
let__14327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5436)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5436
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__14327

tmp16468 := Call(__e, PrimFunc(symshen_4unlocked_2), V5435)


var ifres16430 Obj

if True == tmp16468 {
tmp16431 := Call(__e, PrimFunc(symshen_4lazyderef), V5432, V5434)


let__14329 := tmp16431
_ = let__14329

tmp16467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14329)
}
__typedArg0 := let__14329
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16432 Obj

if True == tmp16467 {
tmp16433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14329)
}
__typedArg0 := let__14329
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16434 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16433, V5434)


let__14330 := tmp16434
_ = let__14330

tmp16466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14330, symdefine)
}
__typedArg0 := let__14330
__typedArg1 := symdefine
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16435 Obj

if True == tmp16466 {
tmp16436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14329)
}
__typedArg0 := let__14329
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16437 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16436, V5434)


let__14331 := tmp16437
_ = let__14331

tmp16465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14331)
}
__typedArg0 := let__14331
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16438 Obj

if True == tmp16465 {
tmp16439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14331)
}
__typedArg0 := let__14331
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14332 := tmp16439
_ = let__14332

tmp16440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14331)
}
__typedArg0 := let__14331
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14333 := tmp16440
_ = let__14333

tmp16441 := Call(__e, PrimFunc(symshen_4newpv), V5434)


let__14334 := tmp16441
_ = let__14334

tmp16442 := Call(__e, PrimFunc(symshen_4newpv), V5434)


let__14335 := tmp16442
_ = let__14335

tmp16443 := Call(__e, PrimFunc(symshen_4newpv), V5434)


let__14336 := tmp16443
_ = let__14336

tmp16444 := Call(__e, PrimFunc(symshen_4newpv), V5434)


let__14337 := tmp16444
_ = let__14337

tmp16445 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16445

tmp16446 := MakeNative(func(__e *ControlFlow) {
tmp16447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14332, let__14333)
}
__typedArg0 := let__14332
__typedArg1 := let__14333
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16448 := Call(__e, PrimFunc(symshen_4sigxrules), tmp16447)


tmp16449 := MakeNative(func(__e *ControlFlow) {
tmp16450 := Call(__e, PrimFunc(symshen_4lazyderef), let__14334, V5434)


tmp16451 := Call(__e, PrimFunc(symfst), tmp16450)


tmp16452 := MakeNative(func(__e *ControlFlow) {
tmp16453 := Call(__e, PrimFunc(symshen_4lazyderef), let__14334, V5434)


tmp16454 := Call(__e, PrimFunc(symsnd), tmp16453)


tmp16455 := MakeNative(func(__e *ControlFlow) {
tmp16456 := Call(__e, PrimFunc(symshen_4deref), let__14337, V5434)


tmp16457 := Call(__e, PrimFunc(symshen_4freshen_1sig), tmp16456)


tmp16458 := MakeNative(func(__e *ControlFlow) {
tmp16459 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symis), let__14337, V5433, V5434, V5435, let__14327, V5437)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1rules), let__14332, let__14335, let__14336, MakeInteger(1), V5434, V5435, let__14327, tmp16459)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14336, tmp16457, V5434, V5435, let__14327, tmp16458)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14335, tmp16454, V5434, V5435, let__14327, tmp16455)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14337, tmp16451, V5434, V5435, let__14327, tmp16452)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14334, tmp16448, V5434, V5435, let__14327, tmp16449)
return


}, 0)

tmp16460 := Call(__e, PrimFunc(symshen_4cut), V5434, V5435, let__14327, tmp16446)


tmp16461 := Call(__e, PrimFunc(symshen_4gc), V5434, tmp16460)


tmp16462 := Call(__e, PrimFunc(symshen_4gc), V5434, tmp16461)


tmp16463 := Call(__e, PrimFunc(symshen_4gc), V5434, tmp16462)


tmp16464 := Call(__e, PrimFunc(symshen_4gc), V5434, tmp16463)


ifres16438 = tmp16464


} else {
ifres16438 = False


}

ifres16435 = ifres16438


} else {
ifres16435 = False


}

ifres16432 = ifres16435


} else {
ifres16432 = False


}

ifres16430 = ifres16432


} else {
ifres16430 = False


}

let__14328 := ifres16430
_ = let__14328

tmp16470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14328, False)
}
__typedArg0 := let__14328
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16470 {
__e.TailApply(PrimFunc(symshen_4unlock), V5435, let__14327)
return
} else {
__e.Return(let__14328)
return
}


}, 6)

tmp16471 := Call(__e, ns2_1set, symshen_4t_d, tmp16428)


_ = tmp16471

tmp16472 := MakeNative(func(__e *ControlFlow) {
V5449 := __e.Get(1)
_ = V5449
tmp16473 := MakeNative(func(__e *ControlFlow) {
Z5451 := __e.Get(1)
_ = Z5451
__e.TailApply(PrimFunc(symshen_4_5sig_drules_6), Z5451)
return
}, 1)

tmp16474 := Call(__e, PrimFunc(symcompile), tmp16473, V5449)


let__14338 := tmp16474
_ = let__14338

tmp16475 := Call(__e, PrimFunc(symfst), let__14338)


tmp16476 := MakeNative(func(__e *ControlFlow) {
Z5452 := __e.Get(1)
_ = Z5452
__e.TailApply(PrimFunc(symshen_4unmark_1rule_1custom_1patterns), Z5452)
return
}, 1)

tmp16477 := Call(__e, PrimFunc(symsnd), let__14338)


tmp16478 := Call(__e, PrimFunc(symmap), tmp16476, tmp16477)


__e.TailApply(PrimFunc(sym_8p), tmp16475, tmp16478)
return


}, 1)

tmp16479 := Call(__e, ns2_1set, symshen_4sigxrules, tmp16472)


_ = tmp16479

tmp16480 := MakeNative(func(__e *ControlFlow) {
V5453 := __e.Get(1)
_ = V5453
tmp16486 := Call(__e, PrimFunc(symtuple_2), V5453)


if True == tmp16486 {
tmp16481 := MakeNative(func(__e *ControlFlow) {
Z5454 := __e.Get(1)
_ = Z5454
__e.TailApply(PrimFunc(symshen_4unmark_1custom_1pattern), Z5454)
return
}, 1)

tmp16482 := Call(__e, PrimFunc(symfst), V5453)


tmp16483 := Call(__e, PrimFunc(symmap), tmp16481, tmp16482)


tmp16484 := Call(__e, PrimFunc(symsnd), V5453)


__e.TailApply(PrimFunc(sym_8p), tmp16483, tmp16484)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4unmark_1rule_1custom_1patterns)
return
}


}, 1)

tmp16487 := Call(__e, ns2_1set, symshen_4unmark_1rule_1custom_1patterns, tmp16480)


_ = tmp16487

tmp16488 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V5455 := __self1
_ = V5455
tmp16524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5455)
}
__typedArg0 := V5455
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16500 Obj

if True == tmp16524 {
tmp16522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5455)
}
__typedArg0 := V5455
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8p, tmp16522)
}
__typedArg0 := sym_8p
__typedArg1 := tmp16522
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16502 Obj

if True == tmp16523 {
tmp16520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5455)
}
__typedArg0 := V5455
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16520)
}
__typedArg0 := tmp16520
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16504 Obj

if True == tmp16521 {
tmp16517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5455)
}
__typedArg0 := V5455
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16517)
}
__typedArg0 := tmp16517
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4custom_1pattern, tmp16518)
}
__typedArg0 := symshen_4custom_1pattern
__typedArg1 := tmp16518
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16506 Obj

if True == tmp16519 {
tmp16514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5455)
}
__typedArg0 := V5455
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16514)
}
__typedArg0 := tmp16514
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16515)
}
__typedArg0 := tmp16515
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16508 Obj

if True == tmp16516 {
tmp16510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5455)
}
__typedArg0 := V5455
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16510)
}
__typedArg0 := tmp16510
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16511)
}
__typedArg0 := tmp16511
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp16512)
}
__typedArg0 := Nil
__typedArg1 := tmp16512
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16509 Obj

if True == tmp16513 {
ifres16509 = True


} else {
ifres16509 = False


}

ifres16508 = ifres16509


} else {
ifres16508 = False


}

var ifres16507 Obj

if True == ifres16508 {
ifres16507 = True


} else {
ifres16507 = False


}

ifres16506 = ifres16507


} else {
ifres16506 = False


}

var ifres16505 Obj

if True == ifres16506 {
ifres16505 = True


} else {
ifres16505 = False


}

ifres16504 = ifres16505


} else {
ifres16504 = False


}

var ifres16503 Obj

if True == ifres16504 {
ifres16503 = True


} else {
ifres16503 = False


}

ifres16502 = ifres16503


} else {
ifres16502 = False


}

var ifres16501 Obj

if True == ifres16502 {
ifres16501 = True


} else {
ifres16501 = False


}

ifres16500 = ifres16501


} else {
ifres16500 = False


}

if True == ifres16500 {
tmp16489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5455)
}
__typedArg0 := V5455
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16489)
}
__typedArg0 := tmp16489
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16490)
}
__typedArg0 := tmp16490
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16492 := Call(__e, PrimFunc(symshen_4curry), tmp16491)


if PrimFunc(symshen_4unmark_1custom_1pattern) == __self {
__self1 = tmp16492
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4unmark_1custom_1pattern), tmp16492)
return


} else {
tmp16498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5455)
}
__typedArg0 := V5455
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16498 {
tmp16493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5455)
}
__typedArg0 := V5455
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16494 := Call(__e, PrimFunc(symshen_4unmark_1custom_1pattern), tmp16493)


tmp16495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5455)
}
__typedArg0 := V5455
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16496 := Call(__e, PrimFunc(symshen_4unmark_1custom_1pattern), tmp16495)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16494, tmp16496)
}
__typedArg0 := tmp16494
__typedArg1 := tmp16496
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V5455)
return
}


}


}, 1)

tmp16525 := Call(__e, ns2_1set, symshen_4unmark_1custom_1pattern, tmp16488)


_ = tmp16525

tmp16526 := MakeNative(func(__e *ControlFlow) {
V5456 := __e.Get(1)
_ = V5456
tmp16553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5456)
}
__typedArg0 := V5456
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16527 Obj

if True == tmp16553 {
tmp16528 := Call(__e, PrimFunc(symtail), V5456)


let__14340 := tmp16528
_ = let__14340

tmp16551 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14340, sym_i)


var ifres16529 Obj

if True == tmp16551 {
tmp16530 := Call(__e, PrimFunc(symtail), let__14340)


let__14341 := tmp16530
_ = let__14341

tmp16531 := Call(__e, PrimFunc(symshen_4_5signature_6), let__14341)


let__14342 := tmp16531
_ = let__14342

tmp16549 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14342)


var ifres16532 Obj

if True == tmp16549 {
tmp16533 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16532 = tmp16533


} else {
tmp16534 := Call(__e, PrimFunc(symshen_4_5_1out), let__14342)


let__14343 := tmp16534
_ = let__14343

tmp16535 := Call(__e, PrimFunc(symshen_4in_1_6), let__14342)


let__14344 := tmp16535
_ = let__14344

tmp16548 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14344, sym_j)


var ifres16536 Obj

if True == tmp16548 {
tmp16537 := Call(__e, PrimFunc(symtail), let__14344)


let__14345 := tmp16537
_ = let__14345

tmp16538 := Call(__e, PrimFunc(symshen_4_5rules_d_6), let__14345)


let__14346 := tmp16538
_ = let__14346

tmp16546 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14346)


var ifres16539 Obj

if True == tmp16546 {
tmp16540 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16539 = tmp16540


} else {
tmp16541 := Call(__e, PrimFunc(symshen_4_5_1out), let__14346)


let__14347 := tmp16541
_ = let__14347

tmp16542 := Call(__e, PrimFunc(symshen_4in_1_6), let__14346)


let__14348 := tmp16542
_ = let__14348

tmp16543 := Call(__e, PrimFunc(symshen_4rectify_1type), let__14343)


let__14349 := tmp16543
_ = let__14349

tmp16544 := Call(__e, PrimFunc(sym_8p), let__14349, let__14347)


tmp16545 := Call(__e, PrimFunc(symshen_4comb), let__14348, tmp16544)


ifres16539 = tmp16545


}

ifres16536 = ifres16539


} else {
tmp16547 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16536 = tmp16547


}

ifres16532 = ifres16536


}

ifres16529 = ifres16532


} else {
tmp16550 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16529 = tmp16550


}

ifres16527 = ifres16529


} else {
tmp16552 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16527 = tmp16552


}

let__14339 := ifres16527
_ = let__14339

tmp16555 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14339)


if True == tmp16555 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__14339)
return
}


}, 1)

tmp16556 := Call(__e, ns2_1set, symshen_4_5sig_drules_6, tmp16526)


_ = tmp16556

tmp16557 := MakeNative(func(__e *ControlFlow) {
V5468 := __e.Get(1)
_ = V5468
tmp16558 := Call(__e, PrimFunc(symshen_4extract_1vars), V5468)


let__14350 := tmp16558
_ = let__14350

tmp16559 := MakeNative(func(__e *ControlFlow) {
Z5471 := __e.Get(1)
_ = Z5471
tmp16560 := Call(__e, PrimFunc(symconcat), sym_e, Z5471)


tmp16561 := Call(__e, PrimFunc(symshen_4freshterm), tmp16560)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z5471, tmp16561)
}
__typedArg0 := Z5471
__typedArg1 := tmp16561
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp16562 := Call(__e, PrimFunc(symmap), tmp16559, let__14350)


let__14351 := tmp16562
_ = let__14351

__e.TailApply(PrimFunc(symshen_4freshen_1type), let__14351, V5468)
return


}, 1)

tmp16563 := Call(__e, ns2_1set, symshen_4freshen_1sig, tmp16557)


_ = tmp16563

tmp16564 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V5472 := __self1
_ = V5472
V5473 := __self2
_ = V5473
tmp16578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5472)
}
__typedArg0 := Nil
__typedArg1 := V5472
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16578 {
__e.Return(V5473)
return
} else {
tmp16576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5472)
}
__typedArg0 := V5472
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16572 Obj

if True == tmp16576 {
tmp16574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5472)
}
__typedArg0 := V5472
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16574)
}
__typedArg0 := tmp16574
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16573 Obj

if True == tmp16575 {
ifres16573 = True


} else {
ifres16573 = False


}

ifres16572 = ifres16573


} else {
ifres16572 = False


}

if True == ifres16572 {
tmp16565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5472)
}
__typedArg0 := V5472
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5472)
}
__typedArg0 := V5472
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16566)
}
__typedArg0 := tmp16566
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5472)
}
__typedArg0 := V5472
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16568)
}
__typedArg0 := tmp16568
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16570 := Call(__e, PrimFunc(symsubst), tmp16567, tmp16569, V5473)


if PrimFunc(symshen_4freshen_1type) == __self {
__self1, __self2 = tmp16565, tmp16570
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4freshen_1type), tmp16565, tmp16570)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshen_1type)
return
}


}


}, 2)

tmp16579 := Call(__e, ns2_1set, symshen_4freshen_1type, tmp16564)


_ = tmp16579

tmp16580 := MakeNative(func(__e *ControlFlow) {
V5474 := __e.Get(1)
_ = V5474
tmp16581 := Call(__e, PrimFunc(symshen_4_5rule_d_6), V5474)


let__14353 := tmp16581
_ = let__14353

tmp16594 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14353)


var ifres16582 Obj

if True == tmp16594 {
tmp16583 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16582 = tmp16583


} else {
tmp16584 := Call(__e, PrimFunc(symshen_4_5_1out), let__14353)


let__14354 := tmp16584
_ = let__14354

tmp16585 := Call(__e, PrimFunc(symshen_4in_1_6), let__14353)


let__14355 := tmp16585
_ = let__14355

tmp16586 := Call(__e, PrimFunc(symshen_4_5rules_d_6), let__14355)


let__14356 := tmp16586
_ = let__14356

tmp16593 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14356)


var ifres16587 Obj

if True == tmp16593 {
tmp16588 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16587 = tmp16588


} else {
tmp16589 := Call(__e, PrimFunc(symshen_4_5_1out), let__14356)


let__14357 := tmp16589
_ = let__14357

tmp16590 := Call(__e, PrimFunc(symshen_4in_1_6), let__14356)


let__14358 := tmp16590
_ = let__14358

tmp16591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14354, let__14357)
}
__typedArg0 := let__14354
__typedArg1 := let__14357
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16592 := Call(__e, PrimFunc(symshen_4comb), let__14358, tmp16591)


ifres16587 = tmp16592


}

ifres16582 = ifres16587


}

let__14352 := ifres16582
_ = let__14352

tmp16606 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14352)


if True == tmp16606 {
tmp16595 := Call(__e, PrimFunc(symshen_4_5rule_d_6), V5474)


let__14360 := tmp16595
_ = let__14360

tmp16602 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14360)


var ifres16596 Obj

if True == tmp16602 {
tmp16597 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16596 = tmp16597


} else {
tmp16598 := Call(__e, PrimFunc(symshen_4_5_1out), let__14360)


let__14361 := tmp16598
_ = let__14361

tmp16599 := Call(__e, PrimFunc(symshen_4in_1_6), let__14360)


let__14362 := tmp16599
_ = let__14362

tmp16600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14361, Nil)
}
__typedArg0 := let__14361
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16601 := Call(__e, PrimFunc(symshen_4comb), let__14362, tmp16600)


ifres16596 = tmp16601


}

let__14359 := ifres16596
_ = let__14359

tmp16604 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14359)


if True == tmp16604 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__14359)
return
}


} else {
__e.Return(let__14352)
return
}


}, 1)

tmp16607 := Call(__e, ns2_1set, symshen_4_5rules_d_6, tmp16580)


_ = tmp16607

tmp16608 := MakeNative(func(__e *ControlFlow) {
V5486 := __e.Get(1)
_ = V5486
tmp16609 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5486)


let__14364 := tmp16609
_ = let__14364

tmp16637 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14364)


var ifres16610 Obj

if True == tmp16637 {
tmp16611 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16610 = tmp16611


} else {
tmp16612 := Call(__e, PrimFunc(symshen_4_5_1out), let__14364)


let__14365 := tmp16612
_ = let__14365

tmp16613 := Call(__e, PrimFunc(symshen_4in_1_6), let__14364)


let__14366 := tmp16613
_ = let__14366

tmp16636 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14366, sym_1_6)


var ifres16614 Obj

if True == tmp16636 {
tmp16615 := Call(__e, PrimFunc(symtail), let__14366)


let__14367 := tmp16615
_ = let__14367

tmp16634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14367)
}
__typedArg0 := let__14367
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16616 Obj

if True == tmp16634 {
tmp16617 := Call(__e, PrimFunc(symhead), let__14367)


let__14368 := tmp16617
_ = let__14368

tmp16618 := Call(__e, PrimFunc(symtail), let__14367)


let__14369 := tmp16618
_ = let__14369

tmp16632 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14369, symwhere)


var ifres16619 Obj

if True == tmp16632 {
tmp16620 := Call(__e, PrimFunc(symtail), let__14369)


let__14370 := tmp16620
_ = let__14370

tmp16630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14370)
}
__typedArg0 := let__14370
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16621 Obj

if True == tmp16630 {
tmp16622 := Call(__e, PrimFunc(symhead), let__14370)


let__14371 := tmp16622
_ = let__14371

tmp16623 := Call(__e, PrimFunc(symtail), let__14370)


let__14372 := tmp16623
_ = let__14372

tmp16624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14368, Nil)
}
__typedArg0 := let__14368
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14371, tmp16624)
}
__typedArg0 := let__14371
__typedArg1 := tmp16624
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16625)
}
__typedArg0 := symwhere
__typedArg1 := tmp16625
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16627 := Call(__e, PrimFunc(sym_8p), let__14365, tmp16626)


tmp16628 := Call(__e, PrimFunc(symshen_4comb), let__14372, tmp16627)


ifres16621 = tmp16628


} else {
tmp16629 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16621 = tmp16629


}

ifres16619 = ifres16621


} else {
tmp16631 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16619 = tmp16631


}

ifres16616 = ifres16619


} else {
tmp16633 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16616 = tmp16633


}

ifres16614 = ifres16616


} else {
tmp16635 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16614 = tmp16635


}

ifres16610 = ifres16614


}

let__14363 := ifres16610
_ = let__14363

tmp16710 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14363)


if True == tmp16710 {
tmp16638 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5486)


let__14374 := tmp16638
_ = let__14374

tmp16667 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14374)


var ifres16639 Obj

if True == tmp16667 {
tmp16640 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16639 = tmp16640


} else {
tmp16641 := Call(__e, PrimFunc(symshen_4_5_1out), let__14374)


let__14375 := tmp16641
_ = let__14375

tmp16642 := Call(__e, PrimFunc(symshen_4in_1_6), let__14374)


let__14376 := tmp16642
_ = let__14376

tmp16666 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14376, sym_5_1)


var ifres16643 Obj

if True == tmp16666 {
tmp16644 := Call(__e, PrimFunc(symtail), let__14376)


let__14377 := tmp16644
_ = let__14377

tmp16664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14377)
}
__typedArg0 := let__14377
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16645 Obj

if True == tmp16664 {
tmp16646 := Call(__e, PrimFunc(symhead), let__14377)


let__14378 := tmp16646
_ = let__14378

tmp16647 := Call(__e, PrimFunc(symtail), let__14377)


let__14379 := tmp16647
_ = let__14379

tmp16662 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14379, symwhere)


var ifres16648 Obj

if True == tmp16662 {
tmp16649 := Call(__e, PrimFunc(symtail), let__14379)


let__14380 := tmp16649
_ = let__14380

tmp16660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14380)
}
__typedArg0 := let__14380
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16650 Obj

if True == tmp16660 {
tmp16651 := Call(__e, PrimFunc(symhead), let__14380)


let__14381 := tmp16651
_ = let__14381

tmp16652 := Call(__e, PrimFunc(symtail), let__14380)


let__14382 := tmp16652
_ = let__14382

tmp16653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14378, Nil)
}
__typedArg0 := let__14378
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14381, tmp16653)
}
__typedArg0 := let__14381
__typedArg1 := tmp16653
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16654)
}
__typedArg0 := symwhere
__typedArg1 := tmp16654
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16656 := Call(__e, PrimFunc(symshen_4correct), tmp16655)


tmp16657 := Call(__e, PrimFunc(sym_8p), let__14375, tmp16656)


tmp16658 := Call(__e, PrimFunc(symshen_4comb), let__14382, tmp16657)


ifres16650 = tmp16658


} else {
tmp16659 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16650 = tmp16659


}

ifres16648 = ifres16650


} else {
tmp16661 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16648 = tmp16661


}

ifres16645 = ifres16648


} else {
tmp16663 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16645 = tmp16663


}

ifres16643 = ifres16645


} else {
tmp16665 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16643 = tmp16665


}

ifres16639 = ifres16643


}

let__14373 := ifres16639
_ = let__14373

tmp16708 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14373)


if True == tmp16708 {
tmp16668 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5486)


let__14384 := tmp16668
_ = let__14384

tmp16685 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14384)


var ifres16669 Obj

if True == tmp16685 {
tmp16670 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16669 = tmp16670


} else {
tmp16671 := Call(__e, PrimFunc(symshen_4_5_1out), let__14384)


let__14385 := tmp16671
_ = let__14385

tmp16672 := Call(__e, PrimFunc(symshen_4in_1_6), let__14384)


let__14386 := tmp16672
_ = let__14386

tmp16684 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14386, sym_5_1)


var ifres16673 Obj

if True == tmp16684 {
tmp16674 := Call(__e, PrimFunc(symtail), let__14386)


let__14387 := tmp16674
_ = let__14387

tmp16682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14387)
}
__typedArg0 := let__14387
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16675 Obj

if True == tmp16682 {
tmp16676 := Call(__e, PrimFunc(symhead), let__14387)


let__14388 := tmp16676
_ = let__14388

tmp16677 := Call(__e, PrimFunc(symtail), let__14387)


let__14389 := tmp16677
_ = let__14389

tmp16678 := Call(__e, PrimFunc(symshen_4correct), let__14388)


tmp16679 := Call(__e, PrimFunc(sym_8p), let__14385, tmp16678)


tmp16680 := Call(__e, PrimFunc(symshen_4comb), let__14389, tmp16679)


ifres16675 = tmp16680


} else {
tmp16681 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16675 = tmp16681


}

ifres16673 = ifres16675


} else {
tmp16683 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16673 = tmp16683


}

ifres16669 = ifres16673


}

let__14383 := ifres16669
_ = let__14383

tmp16706 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14383)


if True == tmp16706 {
tmp16686 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5486)


let__14391 := tmp16686
_ = let__14391

tmp16702 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14391)


var ifres16687 Obj

if True == tmp16702 {
tmp16688 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16687 = tmp16688


} else {
tmp16689 := Call(__e, PrimFunc(symshen_4_5_1out), let__14391)


let__14392 := tmp16689
_ = let__14392

tmp16690 := Call(__e, PrimFunc(symshen_4in_1_6), let__14391)


let__14393 := tmp16690
_ = let__14393

tmp16701 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14393, sym_1_6)


var ifres16691 Obj

if True == tmp16701 {
tmp16692 := Call(__e, PrimFunc(symtail), let__14393)


let__14394 := tmp16692
_ = let__14394

tmp16699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14394)
}
__typedArg0 := let__14394
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16693 Obj

if True == tmp16699 {
tmp16694 := Call(__e, PrimFunc(symhead), let__14394)


let__14395 := tmp16694
_ = let__14395

tmp16695 := Call(__e, PrimFunc(symtail), let__14394)


let__14396 := tmp16695
_ = let__14396

tmp16696 := Call(__e, PrimFunc(sym_8p), let__14392, let__14395)


tmp16697 := Call(__e, PrimFunc(symshen_4comb), let__14396, tmp16696)


ifres16693 = tmp16697


} else {
tmp16698 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16693 = tmp16698


}

ifres16691 = ifres16693


} else {
tmp16700 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16691 = tmp16700


}

ifres16687 = ifres16691


}

let__14390 := ifres16687
_ = let__14390

tmp16704 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14390)


if True == tmp16704 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__14390)
return
}


} else {
__e.Return(let__14383)
return
}


} else {
__e.Return(let__14373)
return
}


} else {
__e.Return(let__14363)
return
}


}, 1)

tmp16711 := Call(__e, ns2_1set, symshen_4_5rule_d_6, tmp16608)


_ = tmp16711

tmp16712 := MakeNative(func(__e *ControlFlow) {
V5521 := __e.Get(1)
_ = V5521
tmp16860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5521)
}
__typedArg0 := V5521
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16804 Obj

if True == tmp16860 {
tmp16858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5521)
}
__typedArg0 := V5521
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp16858)
}
__typedArg0 := symwhere
__typedArg1 := tmp16858
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16806 Obj

if True == tmp16859 {
tmp16856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5521)
}
__typedArg0 := V5521
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16856)
}
__typedArg0 := tmp16856
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16808 Obj

if True == tmp16857 {
tmp16853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5521)
}
__typedArg0 := V5521
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

var ifres16810 Obj

if True == tmp16855 {
tmp16849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5521)
}
__typedArg0 := V5521
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16849)
}
__typedArg0 := tmp16849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16850)
}
__typedArg0 := tmp16850
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16851)
}
__typedArg0 := tmp16851
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16812 Obj

if True == tmp16852 {
tmp16844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5521)
}
__typedArg0 := V5521
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16844)
}
__typedArg0 := tmp16844
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16845)
}
__typedArg0 := tmp16845
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16846)
}
__typedArg0 := tmp16846
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail_1if, tmp16847)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp16847
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16814 Obj

if True == tmp16848 {
tmp16839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5521)
}
__typedArg0 := V5521
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16840)
}
__typedArg0 := tmp16840
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16841)
}
__typedArg0 := tmp16841
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16842)
}
__typedArg0 := tmp16842
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16816 Obj

if True == tmp16843 {
tmp16833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5521)
}
__typedArg0 := V5521
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16837)
}
__typedArg0 := tmp16837
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16818 Obj

if True == tmp16838 {
tmp16826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5521)
}
__typedArg0 := V5521
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16826)
}
__typedArg0 := tmp16826
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16827)
}
__typedArg0 := tmp16827
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16828)
}
__typedArg0 := tmp16828
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

var ifres16820 Obj

if True == tmp16832 {
tmp16822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5521)
}
__typedArg0 := V5521
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16822)
}
__typedArg0 := tmp16822
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16823)
}
__typedArg0 := tmp16823
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp16824)
}
__typedArg0 := Nil
__typedArg1 := tmp16824
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16821 Obj

if True == tmp16825 {
ifres16821 = True


} else {
ifres16821 = False


}

ifres16820 = ifres16821


} else {
ifres16820 = False


}

var ifres16819 Obj

if True == ifres16820 {
ifres16819 = True


} else {
ifres16819 = False


}

ifres16818 = ifres16819


} else {
ifres16818 = False


}

var ifres16817 Obj

if True == ifres16818 {
ifres16817 = True


} else {
ifres16817 = False


}

ifres16816 = ifres16817


} else {
ifres16816 = False


}

var ifres16815 Obj

if True == ifres16816 {
ifres16815 = True


} else {
ifres16815 = False


}

ifres16814 = ifres16815


} else {
ifres16814 = False


}

var ifres16813 Obj

if True == ifres16814 {
ifres16813 = True


} else {
ifres16813 = False


}

ifres16812 = ifres16813


} else {
ifres16812 = False


}

var ifres16811 Obj

if True == ifres16812 {
ifres16811 = True


} else {
ifres16811 = False


}

ifres16810 = ifres16811


} else {
ifres16810 = False


}

var ifres16809 Obj

if True == ifres16810 {
ifres16809 = True


} else {
ifres16809 = False


}

ifres16808 = ifres16809


} else {
ifres16808 = False


}

var ifres16807 Obj

if True == ifres16808 {
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

if True == ifres16804 {
tmp16713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5521)
}
__typedArg0 := V5521
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16713)
}
__typedArg0 := tmp16713
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5521)
}
__typedArg0 := V5521
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16715)
}
__typedArg0 := tmp16715
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16716)
}
__typedArg0 := tmp16716
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16717)
}
__typedArg0 := tmp16717
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16718, Nil)
}
__typedArg0 := tmp16718
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp16719)
}
__typedArg0 := symnot
__typedArg1 := tmp16719
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16720, Nil)
}
__typedArg0 := tmp16720
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16714, tmp16721)
}
__typedArg0 := tmp16714
__typedArg1 := tmp16721
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp16722)
}
__typedArg0 := symand
__typedArg1 := tmp16722
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5521)
}
__typedArg0 := V5521
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16724)
}
__typedArg0 := tmp16724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16725)
}
__typedArg0 := tmp16725
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16726)
}
__typedArg0 := tmp16726
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16727)
}
__typedArg0 := tmp16727
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16723, tmp16728)
}
__typedArg0 := tmp16723
__typedArg1 := tmp16728
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16729)
}
__typedArg0 := symwhere
__typedArg1 := tmp16729
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp16802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5521)
}
__typedArg0 := V5521
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16783 Obj

if True == tmp16802 {
tmp16800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5521)
}
__typedArg0 := V5521
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp16800)
}
__typedArg0 := symwhere
__typedArg1 := tmp16800
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16785 Obj

if True == tmp16801 {
tmp16798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5521)
}
__typedArg0 := V5521
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16798)
}
__typedArg0 := tmp16798
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16787 Obj

if True == tmp16799 {
tmp16795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5521)
}
__typedArg0 := V5521
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16795)
}
__typedArg0 := tmp16795
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16796)
}
__typedArg0 := tmp16796
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16789 Obj

if True == tmp16797 {
tmp16791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5521)
}
__typedArg0 := V5521
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16792)
}
__typedArg0 := tmp16792
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp16793)
}
__typedArg0 := Nil
__typedArg1 := tmp16793
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16790 Obj

if True == tmp16794 {
ifres16790 = True


} else {
ifres16790 = False


}

ifres16789 = ifres16790


} else {
ifres16789 = False


}

var ifres16788 Obj

if True == ifres16789 {
ifres16788 = True


} else {
ifres16788 = False


}

ifres16787 = ifres16788


} else {
ifres16787 = False


}

var ifres16786 Obj

if True == ifres16787 {
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

if True == ifres16783 {
tmp16730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5521)
}
__typedArg0 := V5521
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
return PrimTail(V5521)
}
__typedArg0 := V5521
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, Nil)
}
__typedArg0 := symfail
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
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
return PrimCons(tmp16734, tmp16736)
}
__typedArg0 := tmp16734
__typedArg1 := tmp16736
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp16737)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp16737
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16738, Nil)
}
__typedArg0 := tmp16738
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp16739)
}
__typedArg0 := symnot
__typedArg1 := tmp16739
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16740, Nil)
}
__typedArg0 := tmp16740
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16731, tmp16741)
}
__typedArg0 := tmp16731
__typedArg1 := tmp16741
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp16742)
}
__typedArg0 := symand
__typedArg1 := tmp16742
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5521)
}
__typedArg0 := V5521
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
return PrimCons(tmp16743, tmp16745)
}
__typedArg0 := tmp16743
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
tmp16781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5521)
}
__typedArg0 := V5521
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16762 Obj

if True == tmp16781 {
tmp16779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5521)
}
__typedArg0 := V5521
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail_1if, tmp16779)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp16779
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16764 Obj

if True == tmp16780 {
tmp16777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5521)
}
__typedArg0 := V5521
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16777)
}
__typedArg0 := tmp16777
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16766 Obj

if True == tmp16778 {
tmp16774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5521)
}
__typedArg0 := V5521
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16774)
}
__typedArg0 := tmp16774
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16775)
}
__typedArg0 := tmp16775
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16768 Obj

if True == tmp16776 {
tmp16770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5521)
}
__typedArg0 := V5521
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16770)
}
__typedArg0 := tmp16770
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16771)
}
__typedArg0 := tmp16771
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp16772)
}
__typedArg0 := Nil
__typedArg1 := tmp16772
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16769 Obj

if True == tmp16773 {
ifres16769 = True


} else {
ifres16769 = False


}

ifres16768 = ifres16769


} else {
ifres16768 = False


}

var ifres16767 Obj

if True == ifres16768 {
ifres16767 = True


} else {
ifres16767 = False


}

ifres16766 = ifres16767


} else {
ifres16766 = False


}

var ifres16765 Obj

if True == ifres16766 {
ifres16765 = True


} else {
ifres16765 = False


}

ifres16764 = ifres16765


} else {
ifres16764 = False


}

var ifres16763 Obj

if True == ifres16764 {
ifres16763 = True


} else {
ifres16763 = False


}

ifres16762 = ifres16763


} else {
ifres16762 = False


}

if True == ifres16762 {
tmp16747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5521)
}
__typedArg0 := V5521
return Call(__e, PrimFunc(symtl), __typedArg0)
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
return PrimCons(symnot, tmp16748)
}
__typedArg0 := symnot
__typedArg1 := tmp16748
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5521)
}
__typedArg0 := V5521
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16750)
}
__typedArg0 := tmp16750
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16749, tmp16751)
}
__typedArg0 := tmp16749
__typedArg1 := tmp16751
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16752)
}
__typedArg0 := symwhere
__typedArg1 := tmp16752
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp16753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, Nil)
}
__typedArg0 := symfail
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16753, Nil)
}
__typedArg0 := tmp16753
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5521, tmp16754)
}
__typedArg0 := V5521
__typedArg1 := tmp16754
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp16755)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp16755
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16756, Nil)
}
__typedArg0 := tmp16756
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp16757)
}
__typedArg0 := symnot
__typedArg1 := tmp16757
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5521, Nil)
}
__typedArg0 := V5521
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16758, tmp16759)
}
__typedArg0 := tmp16758
__typedArg1 := tmp16759
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16760)
}
__typedArg0 := symwhere
__typedArg1 := tmp16760
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}


}


}, 1)

tmp16861 := Call(__e, ns2_1set, symshen_4correct, tmp16712)


_ = tmp16861

tmp16862 := MakeNative(func(__e *ControlFlow) {
V5522 := __e.Get(1)
_ = V5522
V5523 := __e.Get(2)
_ = V5523
V5524 := __e.Get(3)
_ = V5524
V5525 := __e.Get(4)
_ = V5525
V5526 := __e.Get(5)
_ = V5526
V5527 := __e.Get(6)
_ = V5527
V5528 := __e.Get(7)
_ = V5528
V5529 := __e.Get(8)
_ = V5529
let__14397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5528)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5528
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__14397

tmp16870 := Call(__e, PrimFunc(symshen_4unlocked_2), V5527)


var ifres16864 Obj

if True == tmp16870 {
tmp16865 := Call(__e, PrimFunc(symshen_4lazyderef), V5523, V5526)


let__14399 := tmp16865
_ = let__14399

tmp16869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14399, Nil)
}
__typedArg0 := let__14399
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16866 Obj

if True == tmp16869 {
tmp16867 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16867

tmp16868 := Call(__e, PrimFunc(symthaw), V5529)


ifres16866 = tmp16868


} else {
ifres16866 = False


}

ifres16864 = ifres16866


} else {
ifres16864 = False


}

let__14398 := ifres16864
_ = let__14398

tmp16895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14398, False)
}
__typedArg0 := let__14398
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16895 {
tmp16891 := Call(__e, PrimFunc(symshen_4unlocked_2), V5527)


var ifres16871 Obj

if True == tmp16891 {
tmp16872 := Call(__e, PrimFunc(symshen_4lazyderef), V5523, V5526)


let__14401 := tmp16872
_ = let__14401

tmp16890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14401)
}
__typedArg0 := let__14401
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16873 Obj

if True == tmp16890 {
tmp16874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14401)
}
__typedArg0 := let__14401
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14402 := tmp16874
_ = let__14402

tmp16875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14401)
}
__typedArg0 := let__14401
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14403 := tmp16875
_ = let__14403

tmp16876 := Call(__e, PrimFunc(symshen_4newpv), V5526)


let__14404 := tmp16876
_ = let__14404

tmp16877 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16877

tmp16878 := Call(__e, PrimFunc(symshen_4deref), let__14402, V5526)


tmp16879 := Call(__e, PrimFunc(symshen_4freshen_1rule), tmp16878)


tmp16880 := MakeNative(func(__e *ControlFlow) {
tmp16881 := Call(__e, PrimFunc(symshen_4lazyderef), let__14404, V5526)


tmp16882 := Call(__e, PrimFunc(symfst), tmp16881)


tmp16883 := Call(__e, PrimFunc(symshen_4lazyderef), let__14404, V5526)


tmp16884 := Call(__e, PrimFunc(symsnd), tmp16883)


tmp16885 := MakeNative(func(__e *ControlFlow) {
tmp16886 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1rules), V5522, let__14403, V5524, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5525)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5525
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})(), V5526, V5527, let__14397, V5529)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5526, V5527, let__14397, tmp16886)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1rule), V5522, V5525, tmp16882, tmp16884, V5524, V5526, V5527, let__14397, tmp16885)
return


}, 0)

tmp16888 := Call(__e, PrimFunc(symbind), let__14404, tmp16879, V5526, V5527, let__14397, tmp16880)


tmp16889 := Call(__e, PrimFunc(symshen_4gc), V5526, tmp16888)


ifres16873 = tmp16889


} else {
ifres16873 = False


}

ifres16871 = ifres16873


} else {
ifres16871 = False


}

let__14400 := ifres16871
_ = let__14400

tmp16893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14400, False)
}
__typedArg0 := let__14400
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16893 {
__e.TailApply(PrimFunc(symshen_4unlock), V5527, let__14397)
return
} else {
__e.Return(let__14400)
return
}


} else {
__e.Return(let__14398)
return
}


}, 8)

tmp16896 := Call(__e, ns2_1set, symshen_4t_d_1rules, tmp16862)


_ = tmp16896

tmp16897 := MakeNative(func(__e *ControlFlow) {
V5538 := __e.Get(1)
_ = V5538
tmp16908 := Call(__e, PrimFunc(symtuple_2), V5538)


if True == tmp16908 {
tmp16898 := Call(__e, PrimFunc(symfst), V5538)


tmp16899 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp16898)


let__14405 := tmp16899
_ = let__14405

tmp16900 := MakeNative(func(__e *ControlFlow) {
Z5541 := __e.Get(1)
_ = Z5541
tmp16901 := Call(__e, PrimFunc(symshen_4freshterm), Z5541)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z5541, tmp16901)
}
__typedArg0 := Z5541
__typedArg1 := tmp16901
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp16902 := Call(__e, PrimFunc(symmap), tmp16900, let__14405)


let__14406 := tmp16902
_ = let__14406

tmp16903 := Call(__e, PrimFunc(symfst), V5538)


tmp16904 := Call(__e, PrimFunc(symshen_4freshen), let__14406, tmp16903)


tmp16905 := Call(__e, PrimFunc(symsnd), V5538)


tmp16906 := Call(__e, PrimFunc(symshen_4freshen), let__14406, tmp16905)


__e.TailApply(PrimFunc(sym_8p), tmp16904, tmp16906)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshen_1rule)
return
}


}, 1)

tmp16909 := Call(__e, ns2_1set, symshen_4freshen_1rule, tmp16897)


_ = tmp16909

tmp16910 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V5542 := __self1
_ = V5542
V5543 := __self2
_ = V5543
tmp16924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5542)
}
__typedArg0 := Nil
__typedArg1 := V5542
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16924 {
__e.Return(V5543)
return
} else {
tmp16922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5542)
}
__typedArg0 := V5542
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16918 Obj

if True == tmp16922 {
tmp16920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5542)
}
__typedArg0 := V5542
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16920)
}
__typedArg0 := tmp16920
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16919 Obj

if True == tmp16921 {
ifres16919 = True


} else {
ifres16919 = False


}

ifres16918 = ifres16919


} else {
ifres16918 = False


}

if True == ifres16918 {
tmp16911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5542)
}
__typedArg0 := V5542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5542)
}
__typedArg0 := V5542
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16912)
}
__typedArg0 := tmp16912
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5542)
}
__typedArg0 := V5542
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16914)
}
__typedArg0 := tmp16914
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16916 := Call(__e, PrimFunc(symshen_4beta), tmp16913, tmp16915, V5543)


if PrimFunc(symshen_4freshen) == __self {
__self1, __self2 = tmp16911, tmp16916
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4freshen), tmp16911, tmp16916)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshen)
return
}


}


}, 2)

tmp16925 := Call(__e, ns2_1set, symshen_4freshen, tmp16910)


_ = tmp16925

tmp16926 := MakeNative(func(__e *ControlFlow) {
V5544 := __e.Get(1)
_ = V5544
V5545 := __e.Get(2)
_ = V5545
V5546 := __e.Get(3)
_ = V5546
V5547 := __e.Get(4)
_ = V5547
V5548 := __e.Get(5)
_ = V5548
V5549 := __e.Get(6)
_ = V5549
V5550 := __e.Get(7)
_ = V5550
V5551 := __e.Get(8)
_ = V5551
V5552 := __e.Get(9)
_ = V5552
tmp16930 := Call(__e, PrimFunc(symshen_4unlocked_2), V5550)


var ifres16927 Obj

if True == tmp16930 {
tmp16928 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16928

tmp16929 := Call(__e, PrimFunc(symshen_4t_d_1rule_1h), V5546, V5547, V5548, V5549, V5550, V5551, V5552)


ifres16927 = tmp16929


} else {
ifres16927 = False


}

let__14407 := ifres16927
_ = let__14407

tmp16942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14407, False)
}
__typedArg0 := let__14407
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16942 {
tmp16940 := Call(__e, PrimFunc(symshen_4unlocked_2), V5550)


if True == tmp16940 {
tmp16931 := Call(__e, PrimFunc(symshen_4newpv), V5549)


let__14408 := tmp16931
_ = let__14408

tmp16932 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16932

tmp16933 := Call(__e, PrimFunc(symshen_4app), V5544, MakeString("\n"), symshen_4a)


tmp16935 := Call(__e, PrimFunc(symshen_4app), V5545, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" of "))
__typedS1, __typedOK1 := TypedString(tmp16933)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" of ")
__typedArg1 := tmp16933
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp16937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("type error in rule "))
__typedS1, __typedOK1 := TypedString(tmp16935)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("type error in rule ")
__typedArg1 := tmp16935
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("type error in rule "))
__typedS1, __typedOK1 := TypedString(tmp16935)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("type error in rule ")
__typedArg1 := tmp16935
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

tmp16938 := Call(__e, PrimFunc(symbind), let__14408, tmp16937, V5549, V5550, V5551, V5552)


__e.TailApply(PrimFunc(symshen_4gc), V5549, tmp16938)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(let__14407)
return
}


}, 9)

tmp16943 := Call(__e, ns2_1set, symshen_4t_d_1rule, tmp16926)


_ = tmp16943

tmp16944 := MakeNative(func(__e *ControlFlow) {
V5555 := __e.Get(1)
_ = V5555
V5556 := __e.Get(2)
_ = V5556
V5557 := __e.Get(3)
_ = V5557
V5558 := __e.Get(4)
_ = V5558
V5559 := __e.Get(5)
_ = V5559
V5560 := __e.Get(6)
_ = V5560
V5561 := __e.Get(7)
_ = V5561
let__14409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5560)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5560
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__14409

tmp16969 := Call(__e, PrimFunc(symshen_4unlocked_2), V5559)


var ifres16946 Obj

if True == tmp16969 {
tmp16947 := Call(__e, PrimFunc(symshen_4lazyderef), V5555, V5558)


let__14411 := tmp16947
_ = let__14411

tmp16968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14411, Nil)
}
__typedArg0 := let__14411
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16948 Obj

if True == tmp16968 {
tmp16949 := Call(__e, PrimFunc(symshen_4lazyderef), V5557, V5558)


let__14412 := tmp16949
_ = let__14412

tmp16967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14412)
}
__typedArg0 := let__14412
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16950 Obj

if True == tmp16967 {
tmp16951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14412)
}
__typedArg0 := let__14412
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16952 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16951, V5558)


let__14413 := tmp16952
_ = let__14413

tmp16966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14413, sym_1_1_6)
}
__typedArg0 := let__14413
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16953 Obj

if True == tmp16966 {
tmp16954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14412)
}
__typedArg0 := let__14412
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16955 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16954, V5558)


let__14414 := tmp16955
_ = let__14414

tmp16965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14414)
}
__typedArg0 := let__14414
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16956 Obj

if True == tmp16965 {
tmp16957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14414)
}
__typedArg0 := let__14414
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14415 := tmp16957
_ = let__14415

tmp16958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14414)
}
__typedArg0 := let__14414
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16959 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16958, V5558)


let__14416 := tmp16959
_ = let__14416

tmp16964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14416, Nil)
}
__typedArg0 := let__14416
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16960 Obj

if True == tmp16964 {
tmp16961 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16961

tmp16962 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1correct), V5556, let__14415, Nil, V5558, V5559, let__14409, V5561)
return
}, 0)

tmp16963 := Call(__e, PrimFunc(symshen_4cut), V5558, V5559, let__14409, tmp16962)


ifres16960 = tmp16963


} else {
ifres16960 = False


}

ifres16956 = ifres16960


} else {
ifres16956 = False


}

ifres16953 = ifres16956


} else {
ifres16953 = False


}

ifres16950 = ifres16953


} else {
ifres16950 = False


}

ifres16948 = ifres16950


} else {
ifres16948 = False


}

ifres16946 = ifres16948


} else {
ifres16946 = False


}

let__14410 := ifres16946
_ = let__14410

tmp16988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14410, False)
}
__typedArg0 := let__14410
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16988 {
tmp16984 := Call(__e, PrimFunc(symshen_4unlocked_2), V5559)


var ifres16970 Obj

if True == tmp16984 {
tmp16971 := Call(__e, PrimFunc(symshen_4newpv), V5558)


let__14418 := tmp16971
_ = let__14418

tmp16972 := Call(__e, PrimFunc(symshen_4newpv), V5558)


let__14419 := tmp16972
_ = let__14419

tmp16973 := Call(__e, PrimFunc(symshen_4newpv), V5558)


let__14420 := tmp16973
_ = let__14420

tmp16974 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16974

tmp16975 := Call(__e, PrimFunc(symshen_4freshterms), V5555)


tmp16976 := MakeNative(func(__e *ControlFlow) {
tmp16977 := MakeNative(func(__e *ControlFlow) {
tmp16978 := MakeNative(func(__e *ControlFlow) {
tmp16979 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1correct), V5556, let__14419, let__14420, V5558, V5559, let__14409, V5561)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4myassume), V5555, V5557, let__14420, V5558, V5559, let__14409, tmp16979)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5558, V5559, let__14409, tmp16978)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1integrity), V5555, V5557, let__14418, let__14419, V5558, V5559, let__14409, tmp16977)
return


}, 0)

tmp16980 := Call(__e, PrimFunc(symshen_4p_1hyps), tmp16975, let__14418, V5558, V5559, let__14409, tmp16976)


tmp16981 := Call(__e, PrimFunc(symshen_4gc), V5558, tmp16980)


tmp16982 := Call(__e, PrimFunc(symshen_4gc), V5558, tmp16981)


tmp16983 := Call(__e, PrimFunc(symshen_4gc), V5558, tmp16982)


ifres16970 = tmp16983


} else {
ifres16970 = False


}

let__14417 := ifres16970
_ = let__14417

tmp16986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14417, False)
}
__typedArg0 := let__14417
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16986 {
__e.TailApply(PrimFunc(symshen_4unlock), V5559, let__14409)
return
} else {
__e.Return(let__14417)
return
}


} else {
__e.Return(let__14410)
return
}


}, 7)

tmp16989 := Call(__e, ns2_1set, symshen_4t_d_1rule_1h, tmp16944)


_ = tmp16989

tmp16990 := MakeNative(func(__e *ControlFlow) {
V5574 := __e.Get(1)
_ = V5574
V5575 := __e.Get(2)
_ = V5575
V5576 := __e.Get(3)
_ = V5576
V5577 := __e.Get(4)
_ = V5577
V5578 := __e.Get(5)
_ = V5578
V5579 := __e.Get(6)
_ = V5579
V5580 := __e.Get(7)
_ = V5580
tmp17004 := Call(__e, PrimFunc(symshen_4unlocked_2), V5578)


var ifres16991 Obj

if True == tmp17004 {
tmp16992 := Call(__e, PrimFunc(symshen_4lazyderef), V5574, V5577)


let__14422 := tmp16992
_ = let__14422

tmp17003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14422, Nil)
}
__typedArg0 := let__14422
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16993 Obj

if True == tmp17003 {
tmp16994 := Call(__e, PrimFunc(symshen_4lazyderef), V5576, V5577)


let__14423 := tmp16994
_ = let__14423

tmp16995 := MakeNative(func(__e *ControlFlow) {
tmp16996 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16996

__e.TailApply(PrimFunc(symthaw), V5580)
return


}, 0)

let__14424 := tmp16995
_ = let__14424

tmp17002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14423, Nil)
}
__typedArg0 := let__14423
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16997 Obj

if True == tmp17002 {
tmp16998 := Call(__e, PrimFunc(symthaw), let__14424)


ifres16997 = tmp16998


} else {
tmp17001 := Call(__e, PrimFunc(symshen_4pvar_2), let__14423)


var ifres16999 Obj

if True == tmp17001 {
tmp17000 := Call(__e, PrimFunc(symshen_4bind_b), let__14423, Nil, V5577, let__14424)


ifres16999 = tmp17000


} else {
ifres16999 = False


}

ifres16997 = ifres16999


}

ifres16993 = ifres16997


} else {
ifres16993 = False


}

ifres16991 = ifres16993


} else {
ifres16991 = False


}

let__14421 := ifres16991
_ = let__14421

tmp17123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14421, False)
}
__typedArg0 := let__14421
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17123 {
tmp17121 := Call(__e, PrimFunc(symshen_4unlocked_2), V5578)


if True == tmp17121 {
tmp17005 := Call(__e, PrimFunc(symshen_4lazyderef), V5574, V5577)


let__14425 := tmp17005
_ = let__14425

tmp17119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14425)
}
__typedArg0 := let__14425
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17119 {
tmp17006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14425)
}
__typedArg0 := let__14425
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14426 := tmp17006
_ = let__14426

tmp17007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14425)
}
__typedArg0 := let__14425
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14427 := tmp17007
_ = let__14427

tmp17008 := Call(__e, PrimFunc(symshen_4lazyderef), V5575, V5577)


let__14428 := tmp17008
_ = let__14428

tmp17117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14428)
}
__typedArg0 := let__14428
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17117 {
tmp17009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14428)
}
__typedArg0 := let__14428
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14429 := tmp17009
_ = let__14429

tmp17010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14428)
}
__typedArg0 := let__14428
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17011 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17010, V5577)


let__14430 := tmp17011
_ = let__14430

tmp17115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14430)
}
__typedArg0 := let__14430
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17115 {
tmp17012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14430)
}
__typedArg0 := let__14430
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17013 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17012, V5577)


let__14431 := tmp17013
_ = let__14431

tmp17113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14431, sym_1_1_6)
}
__typedArg0 := let__14431
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17113 {
tmp17014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14430)
}
__typedArg0 := let__14430
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17015 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17014, V5577)


let__14432 := tmp17015
_ = let__14432

tmp17111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14432)
}
__typedArg0 := let__14432
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17111 {
tmp17016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14432)
}
__typedArg0 := let__14432
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14433 := tmp17016
_ = let__14433

tmp17017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14432)
}
__typedArg0 := let__14432
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17018 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17017, V5577)


let__14434 := tmp17018
_ = let__14434

tmp17109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14434, Nil)
}
__typedArg0 := let__14434
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17109 {
tmp17019 := Call(__e, PrimFunc(symshen_4lazyderef), V5576, V5577)


let__14435 := tmp17019
_ = let__14435

tmp17020 := MakeNative(func(__e *ControlFlow) {
Z5597 := __e.Get(1)
_ = Z5597
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5598 := __e.Get(1)
_ = Z5598
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5599 := __e.Get(1)
_ = Z5599
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5600 := __e.Get(1)
_ = Z5600
tmp17021 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17021

tmp17022 := MakeNative(func(__e *ControlFlow) {
tmp17023 := MakeNative(func(__e *ControlFlow) {
tmp17024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp17025 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4myassume), let__14427, let__14433, Z5600, V5577, V5578, V5579, V5580)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5598, tmp17024, V5577, V5578, V5579, tmp17025)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), let__14426, Z5597, V5577, V5578, V5579, tmp17023)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), let__14429, Z5599, V5577, V5578, V5579, tmp17022)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

let__14436 := tmp17020
_ = let__14436

tmp17107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14435)
}
__typedArg0 := let__14435
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17107 {
tmp17026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14435)
}
__typedArg0 := let__14435
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17027 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17026, V5577)


let__14437 := tmp17027
_ = let__14437

tmp17028 := MakeNative(func(__e *ControlFlow) {
Z5603 := __e.Get(1)
_ = Z5603
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5604 := __e.Get(1)
_ = Z5604
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5605 := __e.Get(1)
_ = Z5605
tmp17029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14435)
}
__typedArg0 := let__14435
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14439 := tmp17029
_ = let__14439

tmp17030 := Call(__e, let__14436, Z5603)


tmp17031 := Call(__e, tmp17030, Z5604)


tmp17032 := Call(__e, tmp17031, Z5605)


__e.TailApply(tmp17032, let__14439)
return


}, 1))
return
}, 1))
return
}, 1)

let__14438 := tmp17028
_ = let__14438

tmp17087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14437)
}
__typedArg0 := let__14437
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17087 {
tmp17033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14437)
}
__typedArg0 := let__14437
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14440 := tmp17033
_ = let__14440

tmp17034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14437)
}
__typedArg0 := let__14437
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17035 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17034, V5577)


let__14441 := tmp17035
_ = let__14441

tmp17036 := MakeNative(func(__e *ControlFlow) {
Z5610 := __e.Get(1)
_ = Z5610
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5611 := __e.Get(1)
_ = Z5611
tmp17037 := Call(__e, let__14438, let__14440)


tmp17038 := Call(__e, tmp17037, Z5610)


__e.TailApply(tmp17038, Z5611)
return


}, 1))
return
}, 1)

let__14442 := tmp17036
_ = let__14442

tmp17071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14441)
}
__typedArg0 := let__14441
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17071 {
tmp17039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14441)
}
__typedArg0 := let__14441
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14443 := tmp17039
_ = let__14443

tmp17040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14441)
}
__typedArg0 := let__14441
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17041 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17040, V5577)


let__14444 := tmp17041
_ = let__14444

tmp17042 := MakeNative(func(__e *ControlFlow) {
Z5615 := __e.Get(1)
_ = Z5615
tmp17043 := Call(__e, let__14442, let__14443)


__e.TailApply(tmp17043, Z5615)
return


}, 1)

let__14445 := tmp17042
_ = let__14445

tmp17059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14444)
}
__typedArg0 := let__14444
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17059 {
tmp17044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14444)
}
__typedArg0 := let__14444
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14446 := tmp17044
_ = let__14446

tmp17045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14444)
}
__typedArg0 := let__14444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17046 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17045, V5577)


let__14447 := tmp17046
_ = let__14447

tmp17047 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14445, let__14446)
return
}, 0)

let__14448 := tmp17047
_ = let__14448

tmp17051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14447, Nil)
}
__typedArg0 := let__14447
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17051 {
__e.TailApply(PrimFunc(symthaw), let__14448)
return
} else {
tmp17049 := Call(__e, PrimFunc(symshen_4pvar_2), let__14447)


if True == tmp17049 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14447, Nil, V5577, let__14448)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp17057 := Call(__e, PrimFunc(symshen_4pvar_2), let__14444)


if True == tmp17057 {
tmp17052 := Call(__e, PrimFunc(symshen_4newpv), V5577)


let__14449 := tmp17052
_ = let__14449

tmp17053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14449, Nil)
}
__typedArg0 := let__14449
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17054 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14445, let__14449)
return
}, 0)

tmp17055 := Call(__e, PrimFunc(symshen_4bind_b), let__14444, tmp17053, V5577, tmp17054)


__e.TailApply(PrimFunc(symshen_4gc), V5577, tmp17055)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17069 := Call(__e, PrimFunc(symshen_4pvar_2), let__14441)


if True == tmp17069 {
tmp17060 := Call(__e, PrimFunc(symshen_4newpv), V5577)


let__14450 := tmp17060
_ = let__14450

tmp17061 := Call(__e, PrimFunc(symshen_4newpv), V5577)


let__14451 := tmp17061
_ = let__14451

tmp17062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14451, Nil)
}
__typedArg0 := let__14451
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14450, tmp17062)
}
__typedArg0 := let__14450
__typedArg1 := tmp17062
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17064 := MakeNative(func(__e *ControlFlow) {
tmp17065 := Call(__e, let__14442, let__14450)


__e.TailApply(tmp17065, let__14451)
return


}, 0)

tmp17066 := Call(__e, PrimFunc(symshen_4bind_b), let__14441, tmp17063, V5577, tmp17064)


tmp17067 := Call(__e, PrimFunc(symshen_4gc), V5577, tmp17066)


__e.TailApply(PrimFunc(symshen_4gc), V5577, tmp17067)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17085 := Call(__e, PrimFunc(symshen_4pvar_2), let__14437)


if True == tmp17085 {
tmp17072 := Call(__e, PrimFunc(symshen_4newpv), V5577)


let__14452 := tmp17072
_ = let__14452

tmp17073 := Call(__e, PrimFunc(symshen_4newpv), V5577)


let__14453 := tmp17073
_ = let__14453

tmp17074 := Call(__e, PrimFunc(symshen_4newpv), V5577)


let__14454 := tmp17074
_ = let__14454

tmp17075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14454, Nil)
}
__typedArg0 := let__14454
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14453, tmp17075)
}
__typedArg0 := let__14453
__typedArg1 := tmp17075
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14452, tmp17076)
}
__typedArg0 := let__14452
__typedArg1 := tmp17076
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17078 := MakeNative(func(__e *ControlFlow) {
tmp17079 := Call(__e, let__14438, let__14452)


tmp17080 := Call(__e, tmp17079, let__14453)


__e.TailApply(tmp17080, let__14454)
return


}, 0)

tmp17081 := Call(__e, PrimFunc(symshen_4bind_b), let__14437, tmp17077, V5577, tmp17078)


tmp17082 := Call(__e, PrimFunc(symshen_4gc), V5577, tmp17081)


tmp17083 := Call(__e, PrimFunc(symshen_4gc), V5577, tmp17082)


__e.TailApply(PrimFunc(symshen_4gc), V5577, tmp17083)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17105 := Call(__e, PrimFunc(symshen_4pvar_2), let__14435)


if True == tmp17105 {
tmp17088 := Call(__e, PrimFunc(symshen_4newpv), V5577)


let__14455 := tmp17088
_ = let__14455

tmp17089 := Call(__e, PrimFunc(symshen_4newpv), V5577)


let__14456 := tmp17089
_ = let__14456

tmp17090 := Call(__e, PrimFunc(symshen_4newpv), V5577)


let__14457 := tmp17090
_ = let__14457

tmp17091 := Call(__e, PrimFunc(symshen_4newpv), V5577)


let__14458 := tmp17091
_ = let__14458

tmp17092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14457, Nil)
}
__typedArg0 := let__14457
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14456, tmp17092)
}
__typedArg0 := let__14456
__typedArg1 := tmp17092
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14455, tmp17093)
}
__typedArg0 := let__14455
__typedArg1 := tmp17093
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17094, let__14458)
}
__typedArg0 := tmp17094
__typedArg1 := let__14458
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17096 := MakeNative(func(__e *ControlFlow) {
tmp17097 := Call(__e, let__14436, let__14455)


tmp17098 := Call(__e, tmp17097, let__14456)


tmp17099 := Call(__e, tmp17098, let__14457)


__e.TailApply(tmp17099, let__14458)
return


}, 0)

tmp17100 := Call(__e, PrimFunc(symshen_4bind_b), let__14435, tmp17095, V5577, tmp17096)


tmp17101 := Call(__e, PrimFunc(symshen_4gc), V5577, tmp17100)


tmp17102 := Call(__e, PrimFunc(symshen_4gc), V5577, tmp17101)


tmp17103 := Call(__e, PrimFunc(symshen_4gc), V5577, tmp17102)


__e.TailApply(PrimFunc(symshen_4gc), V5577, tmp17103)
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
__e.Return(let__14421)
return
}


}, 7)

tmp17124 := Call(__e, ns2_1set, symshen_4myassume, tmp16990)


_ = tmp17124

tmp17125 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V5631 := __self1
_ = V5631
tmp17148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5631)
}
__typedArg0 := Nil
__typedArg1 := V5631
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17148 {
__e.Return(Nil)
return
} else {
tmp17146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5631)
}
__typedArg0 := V5631
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17142 Obj

if True == tmp17146 {
tmp17144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5631)
}
__typedArg0 := V5631
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17144)
}
__typedArg0 := tmp17144
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17143 Obj

if True == tmp17145 {
ifres17143 = True


} else {
ifres17143 = False


}

ifres17142 = ifres17143


} else {
ifres17142 = False


}

if True == ifres17142 {
tmp17126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5631)
}
__typedArg0 := V5631
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5631)
}
__typedArg0 := V5631
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17128 := Call(__e, PrimFunc(symappend), tmp17126, tmp17127)


if PrimFunc(symshen_4freshterms) == __self {
__self1 = tmp17128
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4freshterms), tmp17128)
return


} else {
tmp17140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5631)
}
__typedArg0 := V5631
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17136 Obj

if True == tmp17140 {
tmp17138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5631)
}
__typedArg0 := V5631
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17139 := Call(__e, PrimFunc(symshen_4freshterm_2), tmp17138)


var ifres17137 Obj

if True == tmp17139 {
ifres17137 = True


} else {
ifres17137 = False


}

ifres17136 = ifres17137


} else {
ifres17136 = False


}

if True == ifres17136 {
tmp17129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5631)
}
__typedArg0 := V5631
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5631)
}
__typedArg0 := V5631
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17131 := Call(__e, PrimFunc(symshen_4freshterms), tmp17130)


__e.TailApply(PrimFunc(symadjoin), tmp17129, tmp17131)
return


} else {
tmp17134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5631)
}
__typedArg0 := V5631
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17134 {
tmp17132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5631)
}
__typedArg0 := V5631
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4freshterms) == __self {
__self1 = tmp17132
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4freshterms), tmp17132)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshterms)
return
}


}


}


}


}, 1)

tmp17149 := Call(__e, ns2_1set, symshen_4freshterms, tmp17125)


_ = tmp17149

tmp17150 := MakeNative(func(__e *ControlFlow) {
V5632 := __e.Get(1)
_ = V5632
V5633 := __e.Get(2)
_ = V5633
V5634 := __e.Get(3)
_ = V5634
V5635 := __e.Get(4)
_ = V5635
V5636 := __e.Get(5)
_ = V5636
V5637 := __e.Get(6)
_ = V5637
tmp17164 := Call(__e, PrimFunc(symshen_4unlocked_2), V5635)


var ifres17151 Obj

if True == tmp17164 {
tmp17152 := Call(__e, PrimFunc(symshen_4lazyderef), V5632, V5634)


let__14460 := tmp17152
_ = let__14460

tmp17163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14460, Nil)
}
__typedArg0 := let__14460
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17153 Obj

if True == tmp17163 {
tmp17154 := Call(__e, PrimFunc(symshen_4lazyderef), V5633, V5634)


let__14461 := tmp17154
_ = let__14461

tmp17155 := MakeNative(func(__e *ControlFlow) {
tmp17156 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17156

__e.TailApply(PrimFunc(symthaw), V5637)
return


}, 0)

let__14462 := tmp17155
_ = let__14462

tmp17162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14461, Nil)
}
__typedArg0 := let__14461
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17157 Obj

if True == tmp17162 {
tmp17158 := Call(__e, PrimFunc(symthaw), let__14462)


ifres17157 = tmp17158


} else {
tmp17161 := Call(__e, PrimFunc(symshen_4pvar_2), let__14461)


var ifres17159 Obj

if True == tmp17161 {
tmp17160 := Call(__e, PrimFunc(symshen_4bind_b), let__14461, Nil, V5634, let__14462)


ifres17159 = tmp17160


} else {
ifres17159 = False


}

ifres17157 = ifres17159


}

ifres17153 = ifres17157


} else {
ifres17153 = False


}

ifres17151 = ifres17153


} else {
ifres17151 = False


}

let__14459 := ifres17151
_ = let__14459

tmp17261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14459, False)
}
__typedArg0 := let__14459
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17261 {
tmp17259 := Call(__e, PrimFunc(symshen_4unlocked_2), V5635)


if True == tmp17259 {
tmp17165 := Call(__e, PrimFunc(symshen_4lazyderef), V5632, V5634)


let__14463 := tmp17165
_ = let__14463

tmp17257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14463)
}
__typedArg0 := let__14463
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17257 {
tmp17166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14463)
}
__typedArg0 := let__14463
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14464 := tmp17166
_ = let__14464

tmp17167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14463)
}
__typedArg0 := let__14463
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14465 := tmp17167
_ = let__14465

tmp17168 := Call(__e, PrimFunc(symshen_4lazyderef), V5633, V5634)


let__14466 := tmp17168
_ = let__14466

tmp17169 := MakeNative(func(__e *ControlFlow) {
Z5647 := __e.Get(1)
_ = Z5647
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5648 := __e.Get(1)
_ = Z5648
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5649 := __e.Get(1)
_ = Z5649
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5650 := __e.Get(1)
_ = Z5650
tmp17170 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17170

tmp17171 := MakeNative(func(__e *ControlFlow) {
tmp17172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp17173 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4p_1hyps), let__14465, Z5650, V5634, V5635, V5636, V5637)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5648, tmp17172, V5634, V5635, V5636, tmp17173)
return


}, 0)

__e.TailApply(PrimFunc(symbind), Z5647, let__14464, V5634, V5635, V5636, tmp17171)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

let__14467 := tmp17169
_ = let__14467

tmp17255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14466)
}
__typedArg0 := let__14466
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17255 {
tmp17174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14466)
}
__typedArg0 := let__14466
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17175 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17174, V5634)


let__14468 := tmp17175
_ = let__14468

tmp17176 := MakeNative(func(__e *ControlFlow) {
Z5653 := __e.Get(1)
_ = Z5653
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5654 := __e.Get(1)
_ = Z5654
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5655 := __e.Get(1)
_ = Z5655
tmp17177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14466)
}
__typedArg0 := let__14466
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14470 := tmp17177
_ = let__14470

tmp17178 := Call(__e, let__14467, Z5653)


tmp17179 := Call(__e, tmp17178, Z5654)


tmp17180 := Call(__e, tmp17179, Z5655)


__e.TailApply(tmp17180, let__14470)
return


}, 1))
return
}, 1))
return
}, 1)

let__14469 := tmp17176
_ = let__14469

tmp17235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14468)
}
__typedArg0 := let__14468
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17235 {
tmp17181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14468)
}
__typedArg0 := let__14468
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14471 := tmp17181
_ = let__14471

tmp17182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14468)
}
__typedArg0 := let__14468
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17183 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17182, V5634)


let__14472 := tmp17183
_ = let__14472

tmp17184 := MakeNative(func(__e *ControlFlow) {
Z5660 := __e.Get(1)
_ = Z5660
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5661 := __e.Get(1)
_ = Z5661
tmp17185 := Call(__e, let__14469, let__14471)


tmp17186 := Call(__e, tmp17185, Z5660)


__e.TailApply(tmp17186, Z5661)
return


}, 1))
return
}, 1)

let__14473 := tmp17184
_ = let__14473

tmp17219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14472)
}
__typedArg0 := let__14472
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17219 {
tmp17187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14472)
}
__typedArg0 := let__14472
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14474 := tmp17187
_ = let__14474

tmp17188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14472)
}
__typedArg0 := let__14472
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17189 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17188, V5634)


let__14475 := tmp17189
_ = let__14475

tmp17190 := MakeNative(func(__e *ControlFlow) {
Z5665 := __e.Get(1)
_ = Z5665
tmp17191 := Call(__e, let__14473, let__14474)


__e.TailApply(tmp17191, Z5665)
return


}, 1)

let__14476 := tmp17190
_ = let__14476

tmp17207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14475)
}
__typedArg0 := let__14475
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17207 {
tmp17192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14475)
}
__typedArg0 := let__14475
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14477 := tmp17192
_ = let__14477

tmp17193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14475)
}
__typedArg0 := let__14475
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17194 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17193, V5634)


let__14478 := tmp17194
_ = let__14478

tmp17195 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14476, let__14477)
return
}, 0)

let__14479 := tmp17195
_ = let__14479

tmp17199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14478, Nil)
}
__typedArg0 := let__14478
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17199 {
__e.TailApply(PrimFunc(symthaw), let__14479)
return
} else {
tmp17197 := Call(__e, PrimFunc(symshen_4pvar_2), let__14478)


if True == tmp17197 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14478, Nil, V5634, let__14479)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp17205 := Call(__e, PrimFunc(symshen_4pvar_2), let__14475)


if True == tmp17205 {
tmp17200 := Call(__e, PrimFunc(symshen_4newpv), V5634)


let__14480 := tmp17200
_ = let__14480

tmp17201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14480, Nil)
}
__typedArg0 := let__14480
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17202 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14476, let__14480)
return
}, 0)

tmp17203 := Call(__e, PrimFunc(symshen_4bind_b), let__14475, tmp17201, V5634, tmp17202)


__e.TailApply(PrimFunc(symshen_4gc), V5634, tmp17203)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17217 := Call(__e, PrimFunc(symshen_4pvar_2), let__14472)


if True == tmp17217 {
tmp17208 := Call(__e, PrimFunc(symshen_4newpv), V5634)


let__14481 := tmp17208
_ = let__14481

tmp17209 := Call(__e, PrimFunc(symshen_4newpv), V5634)


let__14482 := tmp17209
_ = let__14482

tmp17210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14482, Nil)
}
__typedArg0 := let__14482
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14481, tmp17210)
}
__typedArg0 := let__14481
__typedArg1 := tmp17210
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17212 := MakeNative(func(__e *ControlFlow) {
tmp17213 := Call(__e, let__14473, let__14481)


__e.TailApply(tmp17213, let__14482)
return


}, 0)

tmp17214 := Call(__e, PrimFunc(symshen_4bind_b), let__14472, tmp17211, V5634, tmp17212)


tmp17215 := Call(__e, PrimFunc(symshen_4gc), V5634, tmp17214)


__e.TailApply(PrimFunc(symshen_4gc), V5634, tmp17215)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17233 := Call(__e, PrimFunc(symshen_4pvar_2), let__14468)


if True == tmp17233 {
tmp17220 := Call(__e, PrimFunc(symshen_4newpv), V5634)


let__14483 := tmp17220
_ = let__14483

tmp17221 := Call(__e, PrimFunc(symshen_4newpv), V5634)


let__14484 := tmp17221
_ = let__14484

tmp17222 := Call(__e, PrimFunc(symshen_4newpv), V5634)


let__14485 := tmp17222
_ = let__14485

tmp17223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14485, Nil)
}
__typedArg0 := let__14485
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14484, tmp17223)
}
__typedArg0 := let__14484
__typedArg1 := tmp17223
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14483, tmp17224)
}
__typedArg0 := let__14483
__typedArg1 := tmp17224
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17226 := MakeNative(func(__e *ControlFlow) {
tmp17227 := Call(__e, let__14469, let__14483)


tmp17228 := Call(__e, tmp17227, let__14484)


__e.TailApply(tmp17228, let__14485)
return


}, 0)

tmp17229 := Call(__e, PrimFunc(symshen_4bind_b), let__14468, tmp17225, V5634, tmp17226)


tmp17230 := Call(__e, PrimFunc(symshen_4gc), V5634, tmp17229)


tmp17231 := Call(__e, PrimFunc(symshen_4gc), V5634, tmp17230)


__e.TailApply(PrimFunc(symshen_4gc), V5634, tmp17231)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17253 := Call(__e, PrimFunc(symshen_4pvar_2), let__14466)


if True == tmp17253 {
tmp17236 := Call(__e, PrimFunc(symshen_4newpv), V5634)


let__14486 := tmp17236
_ = let__14486

tmp17237 := Call(__e, PrimFunc(symshen_4newpv), V5634)


let__14487 := tmp17237
_ = let__14487

tmp17238 := Call(__e, PrimFunc(symshen_4newpv), V5634)


let__14488 := tmp17238
_ = let__14488

tmp17239 := Call(__e, PrimFunc(symshen_4newpv), V5634)


let__14489 := tmp17239
_ = let__14489

tmp17240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14488, Nil)
}
__typedArg0 := let__14488
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14487, tmp17240)
}
__typedArg0 := let__14487
__typedArg1 := tmp17240
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14486, tmp17241)
}
__typedArg0 := let__14486
__typedArg1 := tmp17241
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17242, let__14489)
}
__typedArg0 := tmp17242
__typedArg1 := let__14489
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17244 := MakeNative(func(__e *ControlFlow) {
tmp17245 := Call(__e, let__14467, let__14486)


tmp17246 := Call(__e, tmp17245, let__14487)


tmp17247 := Call(__e, tmp17246, let__14488)


__e.TailApply(tmp17247, let__14489)
return


}, 0)

tmp17248 := Call(__e, PrimFunc(symshen_4bind_b), let__14466, tmp17243, V5634, tmp17244)


tmp17249 := Call(__e, PrimFunc(symshen_4gc), V5634, tmp17248)


tmp17250 := Call(__e, PrimFunc(symshen_4gc), V5634, tmp17249)


tmp17251 := Call(__e, PrimFunc(symshen_4gc), V5634, tmp17250)


__e.TailApply(PrimFunc(symshen_4gc), V5634, tmp17251)
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
__e.Return(let__14459)
return
}


}, 6)

tmp17262 := Call(__e, ns2_1set, symshen_4p_1hyps, tmp17150)


_ = tmp17262

tmp17263 := MakeNative(func(__e *ControlFlow) {
V5679 := __e.Get(1)
_ = V5679
V5680 := __e.Get(2)
_ = V5680
V5681 := __e.Get(3)
_ = V5681
V5682 := __e.Get(4)
_ = V5682
V5683 := __e.Get(5)
_ = V5683
V5684 := __e.Get(6)
_ = V5684
V5685 := __e.Get(7)
_ = V5685
let__14490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5684)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5684
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__14490

tmp17301 := Call(__e, PrimFunc(symshen_4unlocked_2), V5683)


var ifres17265 Obj

if True == tmp17301 {
tmp17266 := Call(__e, PrimFunc(symshen_4lazyderef), V5679, V5682)


let__14492 := tmp17266
_ = let__14492

tmp17300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14492)
}
__typedArg0 := let__14492
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17267 Obj

if True == tmp17300 {
tmp17268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14492)
}
__typedArg0 := let__14492
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17269 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17268, V5682)


let__14493 := tmp17269
_ = let__14493

tmp17299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14493, symwhere)
}
__typedArg0 := let__14493
__typedArg1 := symwhere
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17270 Obj

if True == tmp17299 {
tmp17271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14492)
}
__typedArg0 := let__14492
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17272 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17271, V5682)


let__14494 := tmp17272
_ = let__14494

tmp17298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14494)
}
__typedArg0 := let__14494
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17273 Obj

if True == tmp17298 {
tmp17274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14494)
}
__typedArg0 := let__14494
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14495 := tmp17274
_ = let__14495

tmp17275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14494)
}
__typedArg0 := let__14494
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17276 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17275, V5682)


let__14496 := tmp17276
_ = let__14496

tmp17297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14496)
}
__typedArg0 := let__14496
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17277 Obj

if True == tmp17297 {
tmp17278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14496)
}
__typedArg0 := let__14496
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14497 := tmp17278
_ = let__14497

tmp17279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14496)
}
__typedArg0 := let__14496
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17280 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17279, V5682)


let__14498 := tmp17280
_ = let__14498

tmp17296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14498, Nil)
}
__typedArg0 := let__14498
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17281 Obj

if True == tmp17296 {
tmp17282 := Call(__e, PrimFunc(symshen_4newpv), V5682)


let__14499 := tmp17282
_ = let__14499

tmp17283 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17283

tmp17284 := MakeNative(func(__e *ControlFlow) {
tmp17285 := Call(__e, PrimFunc(symshen_4curry), let__14495)


tmp17286 := MakeNative(func(__e *ControlFlow) {
tmp17287 := MakeNative(func(__e *ControlFlow) {
tmp17288 := MakeNative(func(__e *ControlFlow) {
tmp17289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp17290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symverified, Nil)
}
__typedArg0 := symverified
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17289, tmp17290)
}
__typedArg0 := tmp17289
__typedArg1 := tmp17290
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14499, tmp17291)
}
__typedArg0 := let__14499
__typedArg1 := tmp17291
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17292, V5681)
}
__typedArg0 := tmp17292
__typedArg1 := V5681
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4t_d_1correct), let__14497, V5680, tmp17293, V5682, V5683, let__14490, V5685)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5682, V5683, let__14490, tmp17288)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14499, symboolean, V5681, V5682, V5683, let__14490, tmp17287)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14499, tmp17285, V5682, V5683, let__14490, tmp17286)
return


}, 0)

tmp17294 := Call(__e, PrimFunc(symshen_4cut), V5682, V5683, let__14490, tmp17284)


tmp17295 := Call(__e, PrimFunc(symshen_4gc), V5682, tmp17294)


ifres17281 = tmp17295


} else {
ifres17281 = False


}

ifres17277 = ifres17281


} else {
ifres17277 = False


}

ifres17273 = ifres17277


} else {
ifres17273 = False


}

ifres17270 = ifres17273


} else {
ifres17270 = False


}

ifres17267 = ifres17270


} else {
ifres17267 = False


}

ifres17265 = ifres17267


} else {
ifres17265 = False


}

let__14491 := ifres17265
_ = let__14491

tmp17310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14491, False)
}
__typedArg0 := let__14491
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17310 {
tmp17306 := Call(__e, PrimFunc(symshen_4unlocked_2), V5683)


var ifres17302 Obj

if True == tmp17306 {
tmp17303 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17303

tmp17304 := Call(__e, PrimFunc(symshen_4curry), V5679)


tmp17305 := Call(__e, PrimFunc(symshen_4system_1S_1h), tmp17304, V5680, V5681, V5682, V5683, let__14490, V5685)


ifres17302 = tmp17305


} else {
ifres17302 = False


}

let__14500 := ifres17302
_ = let__14500

tmp17308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14500, False)
}
__typedArg0 := let__14500
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17308 {
__e.TailApply(PrimFunc(symshen_4unlock), V5683, let__14490)
return
} else {
__e.Return(let__14500)
return
}


} else {
__e.Return(let__14491)
return
}


}, 7)

tmp17311 := Call(__e, ns2_1set, symshen_4t_d_1correct, tmp17263)


_ = tmp17311

tmp17312 := MakeNative(func(__e *ControlFlow) {
V5697 := __e.Get(1)
_ = V5697
V5698 := __e.Get(2)
_ = V5698
V5699 := __e.Get(3)
_ = V5699
V5700 := __e.Get(4)
_ = V5700
V5701 := __e.Get(5)
_ = V5701
V5702 := __e.Get(6)
_ = V5702
V5703 := __e.Get(7)
_ = V5703
V5704 := __e.Get(8)
_ = V5704
tmp17319 := Call(__e, PrimFunc(symshen_4unlocked_2), V5702)


var ifres17313 Obj

if True == tmp17319 {
tmp17314 := Call(__e, PrimFunc(symshen_4lazyderef), V5697, V5701)


let__14502 := tmp17314
_ = let__14502

tmp17318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14502, Nil)
}
__typedArg0 := let__14502
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17315 Obj

if True == tmp17318 {
tmp17316 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17316

tmp17317 := Call(__e, PrimFunc(symis_b), V5698, V5700, V5701, V5702, V5703, V5704)


ifres17315 = tmp17317


} else {
ifres17315 = False


}

ifres17313 = ifres17315


} else {
ifres17313 = False


}

let__14501 := ifres17313
_ = let__14501

tmp17351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14501, False)
}
__typedArg0 := let__14501
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17351 {
tmp17349 := Call(__e, PrimFunc(symshen_4unlocked_2), V5702)


if True == tmp17349 {
tmp17320 := Call(__e, PrimFunc(symshen_4lazyderef), V5697, V5701)


let__14503 := tmp17320
_ = let__14503

tmp17347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14503)
}
__typedArg0 := let__14503
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17347 {
tmp17321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14503)
}
__typedArg0 := let__14503
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14504 := tmp17321
_ = let__14504

tmp17322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14503)
}
__typedArg0 := let__14503
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14505 := tmp17322
_ = let__14505

tmp17323 := Call(__e, PrimFunc(symshen_4lazyderef), V5698, V5701)


let__14506 := tmp17323
_ = let__14506

tmp17345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14506)
}
__typedArg0 := let__14506
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17345 {
tmp17324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14506)
}
__typedArg0 := let__14506
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14507 := tmp17324
_ = let__14507

tmp17325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14506)
}
__typedArg0 := let__14506
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17326 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17325, V5701)


let__14508 := tmp17326
_ = let__14508

tmp17343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14508)
}
__typedArg0 := let__14508
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17343 {
tmp17327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14508)
}
__typedArg0 := let__14508
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17328 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17327, V5701)


let__14509 := tmp17328
_ = let__14509

tmp17341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14509, sym_1_1_6)
}
__typedArg0 := let__14509
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17341 {
tmp17329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14508)
}
__typedArg0 := let__14508
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17330 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17329, V5701)


let__14510 := tmp17330
_ = let__14510

tmp17339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14510)
}
__typedArg0 := let__14510
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17339 {
tmp17331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14510)
}
__typedArg0 := let__14510
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14511 := tmp17331
_ = let__14511

tmp17332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14510)
}
__typedArg0 := let__14510
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17333 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17332, V5701)


let__14512 := tmp17333
_ = let__14512

tmp17337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14512, Nil)
}
__typedArg0 := let__14512
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17337 {
tmp17334 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17334

tmp17335 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1integrity), let__14505, let__14511, V5699, V5700, V5701, V5702, V5703, V5704)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14504, let__14507, V5699, V5701, V5702, V5703, tmp17335)
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
__e.Return(let__14501)
return
}


}, 8)

tmp17352 := Call(__e, ns2_1set, symshen_4t_d_1integrity, tmp17312)


_ = tmp17352

tmp17353 := MakeNative(func(__e *ControlFlow) {
V5717 := __e.Get(1)
_ = V5717
tmp17362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V5717)
}
__typedArg0 := V5717
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp17362 {
tmp17359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V5717)
}
__typedArg0 := V5717
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

tmp17360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp17359)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp17359
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres17355 Obj

if True == tmp17360 {
tmp17357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V5717, MakeInteger(0))
}
__typedArg0 := V5717
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp17358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp17357, symshen_4print_1freshterm)
}
__typedArg0 := tmp17357
__typedArg1 := symshen_4print_1freshterm
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17356 Obj

if True == tmp17358 {
ifres17356 = True


} else {
ifres17356 = False


}

ifres17355 = ifres17356


} else {
ifres17355 = False


}

if True == ifres17355 {
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

__e.TailApply(ns2_1set, symshen_4freshterm_2, tmp17353)
return




}, 0)

