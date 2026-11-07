package main

import . "github.com/pyrex41/shen-go/kl"

var TStarMain = MakeNative(func(__e *ControlFlow) {
tmp14540 := MakeNative(func(__e *ControlFlow) {
V4930 := __e.Get(1)
_ = V4930
V4931 := __e.Get(2)
_ = V4931
tmp14541 := Call(__e, PrimFunc(symshen_4extract_1vars), V4931)


let__13989 := tmp14541
_ = let__13989

tmp14542 := Call(__e, PrimFunc(symshen_4rectify_1type), V4931)


let__13990 := tmp14542
_ = let__13990

tmp14543 := Call(__e, PrimFunc(symshen_4curry), V4930)


let__13991 := tmp14543
_ = let__13991

tmp14544 := MakeNative(func(__e *ControlFlow) {
Z4935 := __e.Get(1)
_ = Z4935
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4936 := __e.Get(1)
_ = Z4936
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4937 := __e.Get(1)
_ = Z4937
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4938 := __e.Get(1)
_ = Z4938
tmp14545 := Call(__e, PrimFunc(symshen_4newpv), Z4935)


let__13992 := tmp14545
_ = let__13992

tmp14546 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14546

tmp14547 := Call(__e, PrimFunc(symshen_4deref), let__13989, Z4935)


tmp14548 := Call(__e, PrimFunc(symreceive), tmp14547)


tmp14549 := Call(__e, PrimFunc(symshen_4deref), let__13990, Z4935)


tmp14550 := Call(__e, PrimFunc(symreceive), tmp14549)


tmp14551 := MakeNative(func(__e *ControlFlow) {
tmp14552 := Call(__e, PrimFunc(symshen_4deref), let__13991, Z4935)


tmp14553 := Call(__e, PrimFunc(symreceive), tmp14552)


tmp14554 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symreturn), let__13992, Z4935, Z4936, Z4937, Z4938)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4toplevel_1forms), tmp14553, let__13992, Z4935, Z4936, Z4937, tmp14554)
return


}, 0)

tmp14555 := Call(__e, PrimFunc(symshen_4insert_1prolog_1variables), tmp14548, tmp14550, let__13992, Z4935, Z4936, Z4937, tmp14551)


__e.TailApply(PrimFunc(symshen_4gc), Z4935, tmp14555)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp14556 := Call(__e, PrimFunc(symshen_4prolog_1vector))


tmp14557 := Call(__e, tmp14544, tmp14556)


tmp14558 := Call(__e, PrimFunc(symvector), MakeInteger(0))


tmp14559 := Call(__e, PrimFunc(sym_8v), MakeInteger(0), tmp14558)


tmp14560 := Call(__e, PrimFunc(sym_8v), True, tmp14559)


tmp14561 := Call(__e, tmp14557, tmp14560)


tmp14562 := Call(__e, tmp14561, MakeInteger(0))


tmp14563 := MakeNative(func(__e *ControlFlow) {
__e.Return(True)
return
}, 0)

__e.TailApply(tmp14562, tmp14563)
return


}, 2)

tmp14564 := Call(__e, ns2_1set, symshen_4typecheck, tmp14540)


_ = tmp14564

tmp14565 := MakeNative(func(__e *ControlFlow) {
V4940 := __e.Get(1)
_ = V4940
V4941 := __e.Get(2)
_ = V4941
V4942 := __e.Get(3)
_ = V4942
V4943 := __e.Get(4)
_ = V4943
V4944 := __e.Get(5)
_ = V4944
V4945 := __e.Get(6)
_ = V4945
V4946 := __e.Get(7)
_ = V4946
tmp14572 := Call(__e, PrimFunc(symshen_4unlocked_2), V4944)


var ifres14566 Obj

if True == tmp14572 {
tmp14567 := Call(__e, PrimFunc(symshen_4lazyderef), V4940, V4943)


let__13994 := tmp14567
_ = let__13994

tmp14571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13994, Nil)
}
__typedArg0 := let__13994
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14568 Obj

if True == tmp14571 {
tmp14569 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14569

tmp14570 := Call(__e, PrimFunc(symis_b), V4941, V4942, V4943, V4944, V4945, V4946)


ifres14568 = tmp14570


} else {
ifres14568 = False


}

ifres14566 = ifres14568


} else {
ifres14566 = False


}

let__13993 := ifres14566
_ = let__13993

tmp14586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__13993, False)
}
__typedArg0 := let__13993
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14586 {
tmp14584 := Call(__e, PrimFunc(symshen_4unlocked_2), V4944)


if True == tmp14584 {
tmp14573 := Call(__e, PrimFunc(symshen_4lazyderef), V4940, V4943)


let__13995 := tmp14573
_ = let__13995

tmp14582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__13995)
}
__typedArg0 := let__13995
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14582 {
tmp14574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__13995)
}
__typedArg0 := let__13995
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__13996 := tmp14574
_ = let__13996

tmp14575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__13995)
}
__typedArg0 := let__13995
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__13997 := tmp14575
_ = let__13997

tmp14576 := Call(__e, PrimFunc(symshen_4newpv), V4943)


let__13998 := tmp14576
_ = let__13998

tmp14577 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14577

tmp14578 := Call(__e, PrimFunc(symshen_4deref), let__13998, V4943)


tmp14579 := Call(__e, PrimFunc(symsubst), tmp14578, let__13996, V4941)


tmp14580 := Call(__e, PrimFunc(symshen_4insert_1prolog_1variables), let__13997, tmp14579, V4942, V4943, V4944, V4945, V4946)


__e.TailApply(PrimFunc(symshen_4gc), V4943, tmp14580)
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
__e.Return(let__13993)
return
}


}, 7)

tmp14587 := Call(__e, ns2_1set, symshen_4insert_1prolog_1variables, tmp14565)


_ = tmp14587

tmp14588 := MakeNative(func(__e *ControlFlow) {
V4953 := __e.Get(1)
_ = V4953
V4954 := __e.Get(2)
_ = V4954
V4955 := __e.Get(3)
_ = V4955
V4956 := __e.Get(4)
_ = V4956
V4957 := __e.Get(5)
_ = V4957
V4958 := __e.Get(6)
_ = V4958
let__13999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4957)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4957
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__13999

tmp14613 := Call(__e, PrimFunc(symshen_4unlocked_2), V4956)


var ifres14590 Obj

if True == tmp14613 {
tmp14591 := Call(__e, PrimFunc(symshen_4lazyderef), V4953, V4955)


let__14001 := tmp14591
_ = let__14001

tmp14612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14001)
}
__typedArg0 := let__14001
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14592 Obj

if True == tmp14612 {
tmp14593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14001)
}
__typedArg0 := let__14001
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14594 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14593, V4955)


let__14002 := tmp14594
_ = let__14002

tmp14611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14002, symdefine)
}
__typedArg0 := let__14002
__typedArg1 := symdefine
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14595 Obj

if True == tmp14611 {
tmp14596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14001)
}
__typedArg0 := let__14001
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14597 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14596, V4955)


let__14003 := tmp14597
_ = let__14003

tmp14610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14003)
}
__typedArg0 := let__14003
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14598 Obj

if True == tmp14610 {
tmp14599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14003)
}
__typedArg0 := let__14003
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14004 := tmp14599
_ = let__14004

tmp14600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14003)
}
__typedArg0 := let__14003
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14005 := tmp14600
_ = let__14005

tmp14601 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14601

tmp14602 := Call(__e, PrimFunc(symshen_4type_1theory_1enabled_2))


tmp14603 := MakeNative(func(__e *ControlFlow) {
tmp14604 := MakeNative(func(__e *ControlFlow) {
tmp14605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspy_d)
}
__typedArg0 := symshen_4_dspy_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp14606 := MakeNative(func(__e *ControlFlow) {
tmp14607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14004, let__14005)
}
__typedArg0 := let__14004
__typedArg1 := let__14005
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp14607)
}
__typedArg0 := symdefine
__typedArg1 := tmp14607
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4t_d), tmp14608, V4954, V4955, V4956, let__13999, V4958)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4signal_1def), tmp14605, let__14004, V4955, V4956, let__13999, tmp14606)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4955, V4956, let__13999, tmp14604)
return


}, 0)

tmp14609 := Call(__e, PrimFunc(symwhen), tmp14602, V4955, V4956, let__13999, tmp14603)


ifres14598 = tmp14609


} else {
ifres14598 = False


}

ifres14595 = ifres14598


} else {
ifres14595 = False


}

ifres14592 = ifres14595


} else {
ifres14592 = False


}

ifres14590 = ifres14592


} else {
ifres14590 = False


}

let__14000 := ifres14590
_ = let__14000

tmp14625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14000, False)
}
__typedArg0 := let__14000
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14625 {
tmp14621 := Call(__e, PrimFunc(symshen_4unlocked_2), V4956)


var ifres14614 Obj

if True == tmp14621 {
tmp14615 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14615

tmp14616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp14617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4954, Nil)
}
__typedArg0 := V4954
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14616, tmp14617)
}
__typedArg0 := tmp14616
__typedArg1 := tmp14617
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4953, tmp14618)
}
__typedArg0 := V4953
__typedArg1 := tmp14618
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14620 := Call(__e, PrimFunc(symshen_4system_1S), tmp14619, Nil, V4955, V4956, let__13999, V4958)


ifres14614 = tmp14620


} else {
ifres14614 = False


}

let__14006 := ifres14614
_ = let__14006

tmp14623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14006, False)
}
__typedArg0 := let__14006
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14623 {
__e.TailApply(PrimFunc(symshen_4unlock), V4956, let__13999)
return
} else {
__e.Return(let__14006)
return
}


} else {
__e.Return(let__14000)
return
}


}, 6)

tmp14626 := Call(__e, ns2_1set, symshen_4toplevel_1forms, tmp14588)


_ = tmp14626

tmp14627 := MakeNative(func(__e *ControlFlow) {
V4967 := __e.Get(1)
_ = V4967
V4968 := __e.Get(2)
_ = V4968
V4969 := __e.Get(3)
_ = V4969
V4970 := __e.Get(4)
_ = V4970
V4971 := __e.Get(5)
_ = V4971
V4972 := __e.Get(6)
_ = V4972
tmp14634 := Call(__e, PrimFunc(symshen_4unlocked_2), V4970)


var ifres14628 Obj

if True == tmp14634 {
tmp14629 := Call(__e, PrimFunc(symshen_4lazyderef), V4967, V4969)


let__14008 := tmp14629
_ = let__14008

tmp14633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14008, False)
}
__typedArg0 := let__14008
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14630 Obj

if True == tmp14633 {
tmp14631 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14631

tmp14632 := Call(__e, PrimFunc(symthaw), V4972)


ifres14630 = tmp14632


} else {
ifres14630 = False


}

ifres14628 = ifres14630


} else {
ifres14628 = False


}

let__14007 := ifres14628
_ = let__14007

tmp14649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14007, False)
}
__typedArg0 := let__14007
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14649 {
tmp14647 := Call(__e, PrimFunc(symshen_4unlocked_2), V4970)


if True == tmp14647 {
tmp14635 := Call(__e, PrimFunc(symshen_4lazyderef), V4967, V4969)


let__14009 := tmp14635
_ = let__14009

tmp14645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14009, True)
}
__typedArg0 := let__14009
__typedArg1 := True
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14645 {
tmp14636 := Call(__e, PrimFunc(symshen_4newpv), V4969)


let__14010 := tmp14636
_ = let__14010

tmp14637 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14637

tmp14638 := Call(__e, PrimFunc(symshen_4deref), V4968, V4969)


tmp14639 := Call(__e, PrimFunc(symshen_4app), tmp14638, MakeString(")\n"), symshen_4a)


tmp14640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\ntypechecking (fn "))
__typedS1, __typedOK1 := TypedString(tmp14639)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\ntypechecking (fn ")
__typedArg1 := tmp14639
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp14641 := Call(__e, PrimFunc(symstoutput))


tmp14642 := Call(__e, PrimFunc(sympr), tmp14640, tmp14641)


tmp14643 := Call(__e, PrimFunc(symis), let__14010, tmp14642, V4969, V4970, V4971, V4972)


__e.TailApply(PrimFunc(symshen_4gc), V4969, tmp14643)
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
__e.Return(let__14007)
return
}


}, 6)

tmp14650 := Call(__e, ns2_1set, symshen_4signal_1def, tmp14627)


_ = tmp14650

tmp14651 := MakeNative(func(__e *ControlFlow) {
V4977 := __e.Get(1)
_ = V4977
tmp14652 := Call(__e, PrimFunc(symshen_4curry_1type), V4977)


__e.TailApply(PrimFunc(symshen_4demodulate), tmp14652)
return


}, 1)

tmp14653 := Call(__e, ns2_1set, symshen_4rectify_1type, tmp14651)


_ = tmp14653

tmp14654 := MakeNative(func(__e *ControlFlow) {
V4978 := __e.Get(1)
_ = V4978
tmp14655 := MakeNative(func(__e *ControlFlow) {
tmp14656 := MakeNative(func(__e *ControlFlow) {
Z4980 := __e.Get(1)
_ = Z4980
__e.TailApply(PrimFunc(symshen_4demod), Z4980)
return
}, 1)

tmp14657 := Call(__e, PrimFunc(symshen_4walk), tmp14656, V4978)


let__14011 := tmp14657
_ = let__14011

tmp14659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14011, V4978)
}
__typedArg0 := let__14011
__typedArg1 := V4978
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14659 {
__e.Return(V4978)
return
} else {
__e.TailApply(PrimFunc(symshen_4demodulate), let__14011)
return
}


}, 0)

tmp14660 := MakeNative(func(__e *ControlFlow) {
Z4981 := __e.Get(1)
_ = Z4981
__e.Return(V4978)
return
}, 1)

__e.TailApply(try_1catch, tmp14655, tmp14660)
return


}, 1)

tmp14661 := Call(__e, ns2_1set, symshen_4demodulate, tmp14654)


_ = tmp14661

tmp14662 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V4982 := __self1
_ = V4982
tmp14786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14759 Obj

if True == tmp14786 {
tmp14784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14784)
}
__typedArg0 := tmp14784
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14761 Obj

if True == tmp14785 {
tmp14781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14781)
}
__typedArg0 := tmp14781
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1_1_6, tmp14782)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp14782
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14763 Obj

if True == tmp14783 {
tmp14778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14778)
}
__typedArg0 := tmp14778
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14779)
}
__typedArg0 := tmp14779
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14765 Obj

if True == tmp14780 {
tmp14774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14774)
}
__typedArg0 := tmp14774
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14775)
}
__typedArg0 := tmp14775
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14776)
}
__typedArg0 := tmp14776
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14767 Obj

if True == tmp14777 {
tmp14769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14769)
}
__typedArg0 := tmp14769
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14770)
}
__typedArg0 := tmp14770
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14771)
}
__typedArg0 := tmp14771
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1_1_6, tmp14772)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp14772
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14768 Obj

if True == tmp14773 {
ifres14768 = True


} else {
ifres14768 = False


}

ifres14767 = ifres14768


} else {
ifres14767 = False


}

var ifres14766 Obj

if True == ifres14767 {
ifres14766 = True


} else {
ifres14766 = False


}

ifres14765 = ifres14766


} else {
ifres14765 = False


}

var ifres14764 Obj

if True == ifres14765 {
ifres14764 = True


} else {
ifres14764 = False


}

ifres14763 = ifres14764


} else {
ifres14763 = False


}

var ifres14762 Obj

if True == ifres14763 {
ifres14762 = True


} else {
ifres14762 = False


}

ifres14761 = ifres14762


} else {
ifres14761 = False


}

var ifres14760 Obj

if True == ifres14761 {
ifres14760 = True


} else {
ifres14760 = False


}

ifres14759 = ifres14760


} else {
ifres14759 = False


}

if True == ifres14759 {
tmp14663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14664)
}
__typedArg0 := tmp14664
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14665, Nil)
}
__typedArg0 := tmp14665
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp14666)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp14666
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14663, tmp14667)
}
__typedArg0 := tmp14663
__typedArg1 := tmp14667
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4curry_1type) == __self {
__self1 = tmp14668
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4curry_1type), tmp14668)
return


} else {
tmp14757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14717 Obj

if True == tmp14757 {
tmp14755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14755)
}
__typedArg0 := tmp14755
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14719 Obj

if True == tmp14756 {
tmp14752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14752)
}
__typedArg0 := tmp14752
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp14753)
}
__typedArg0 := symlist
__typedArg1 := tmp14753
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14721 Obj

if True == tmp14754 {
tmp14749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14749)
}
__typedArg0 := tmp14749
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14750)
}
__typedArg0 := tmp14750
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14723 Obj

if True == tmp14751 {
tmp14745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14745)
}
__typedArg0 := tmp14745
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14746)
}
__typedArg0 := tmp14746
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14747)
}
__typedArg0 := Nil
__typedArg1 := tmp14747
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14725 Obj

if True == tmp14748 {
tmp14743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14743)
}
__typedArg0 := tmp14743
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14727 Obj

if True == tmp14744 {
tmp14740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14740)
}
__typedArg0 := tmp14740
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_a_a_6, tmp14741)
}
__typedArg0 := sym_a_a_6
__typedArg1 := tmp14741
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14729 Obj

if True == tmp14742 {
tmp14737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14737)
}
__typedArg0 := tmp14737
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14738)
}
__typedArg0 := tmp14738
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14731 Obj

if True == tmp14739 {
tmp14733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14733)
}
__typedArg0 := tmp14733
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14734)
}
__typedArg0 := tmp14734
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14735)
}
__typedArg0 := Nil
__typedArg1 := tmp14735
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14732 Obj

if True == tmp14736 {
ifres14732 = True


} else {
ifres14732 = False


}

ifres14731 = ifres14732


} else {
ifres14731 = False


}

var ifres14730 Obj

if True == ifres14731 {
ifres14730 = True


} else {
ifres14730 = False


}

ifres14729 = ifres14730


} else {
ifres14729 = False


}

var ifres14728 Obj

if True == ifres14729 {
ifres14728 = True


} else {
ifres14728 = False


}

ifres14727 = ifres14728


} else {
ifres14727 = False


}

var ifres14726 Obj

if True == ifres14727 {
ifres14726 = True


} else {
ifres14726 = False


}

ifres14725 = ifres14726


} else {
ifres14725 = False


}

var ifres14724 Obj

if True == ifres14725 {
ifres14724 = True


} else {
ifres14724 = False


}

ifres14723 = ifres14724


} else {
ifres14723 = False


}

var ifres14722 Obj

if True == ifres14723 {
ifres14722 = True


} else {
ifres14722 = False


}

ifres14721 = ifres14722


} else {
ifres14721 = False


}

var ifres14720 Obj

if True == ifres14721 {
ifres14720 = True


} else {
ifres14720 = False


}

ifres14719 = ifres14720


} else {
ifres14719 = False


}

var ifres14718 Obj

if True == ifres14719 {
ifres14718 = True


} else {
ifres14718 = False


}

ifres14717 = ifres14718


} else {
ifres14717 = False


}

if True == ifres14717 {
tmp14669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4982)
}
__typedArg0 := V4982
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14670, tmp14672)
}
__typedArg0 := tmp14670
__typedArg1 := tmp14672
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstr, tmp14673)
}
__typedArg0 := symstr
__typedArg1 := tmp14673
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14674, Nil)
}
__typedArg0 := tmp14674
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp14675)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp14675
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14669, tmp14676)
}
__typedArg0 := tmp14669
__typedArg1 := tmp14676
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4curry_1type) == __self {
__self1 = tmp14677
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4curry_1type), tmp14677)
return


} else {
tmp14715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14688 Obj

if True == tmp14715 {
tmp14713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14713)
}
__typedArg0 := tmp14713
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14690 Obj

if True == tmp14714 {
tmp14710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14710)
}
__typedArg0 := tmp14710
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_d, tmp14711)
}
__typedArg0 := sym_d
__typedArg1 := tmp14711
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14692 Obj

if True == tmp14712 {
tmp14707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4982)
}
__typedArg0 := V4982
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14708)
}
__typedArg0 := tmp14708
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14694 Obj

if True == tmp14709 {
tmp14703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14703)
}
__typedArg0 := tmp14703
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14704)
}
__typedArg0 := tmp14704
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14705)
}
__typedArg0 := tmp14705
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14696 Obj

if True == tmp14706 {
tmp14698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14698)
}
__typedArg0 := tmp14698
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14699)
}
__typedArg0 := tmp14699
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14700)
}
__typedArg0 := tmp14700
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_d, tmp14701)
}
__typedArg0 := sym_d
__typedArg1 := tmp14701
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14697 Obj

if True == tmp14702 {
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

var ifres14689 Obj

if True == ifres14690 {
ifres14689 = True


} else {
ifres14689 = False


}

ifres14688 = ifres14689


} else {
ifres14688 = False


}

if True == ifres14688 {
tmp14678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4982)
}
__typedArg0 := V4982
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14680, Nil)
}
__typedArg0 := tmp14680
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp14681)
}
__typedArg0 := sym_d
__typedArg1 := tmp14681
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14678, tmp14682)
}
__typedArg0 := tmp14678
__typedArg1 := tmp14682
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4curry_1type) == __self {
__self1 = tmp14683
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4curry_1type), tmp14683)
return


} else {
tmp14686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4982)
}
__typedArg0 := V4982
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp14686 {
tmp14684 := MakeNative(func(__e *ControlFlow) {
Z4983 := __e.Get(1)
_ = Z4983
__e.TailApply(PrimFunc(symshen_4curry_1type), Z4983)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp14684, V4982)
return


} else {
__e.Return(V4982)
return
}


}


}


}


}, 1)

tmp14787 := Call(__e, ns2_1set, symshen_4curry_1type, tmp14662)


_ = tmp14787

tmp14788 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V4984 := __self1
_ = V4984
tmp14877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14869 Obj

if True == tmp14877 {
tmp14875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp14875)
}
__typedArg0 := symdefine
__typedArg1 := tmp14875
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14871 Obj

if True == tmp14876 {
tmp14873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14873)
}
__typedArg0 := tmp14873
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14872 Obj

if True == tmp14874 {
ifres14872 = True


} else {
ifres14872 = False


}

ifres14871 = ifres14872


} else {
ifres14871 = False


}

var ifres14870 Obj

if True == ifres14871 {
ifres14870 = True


} else {
ifres14870 = False


}

ifres14869 = ifres14870


} else {
ifres14869 = False


}

if True == ifres14869 {
__e.Return(V4984)
return
} else {
tmp14867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14848 Obj

if True == tmp14867 {
tmp14865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtype, tmp14865)
}
__typedArg0 := symtype
__typedArg1 := tmp14865
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14850 Obj

if True == tmp14866 {
tmp14863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14863)
}
__typedArg0 := tmp14863
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14852 Obj

if True == tmp14864 {
tmp14860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14860)
}
__typedArg0 := tmp14860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14861)
}
__typedArg0 := tmp14861
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14854 Obj

if True == tmp14862 {
tmp14856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14856)
}
__typedArg0 := tmp14856
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14857)
}
__typedArg0 := tmp14857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14858)
}
__typedArg0 := Nil
__typedArg1 := tmp14858
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14855 Obj

if True == tmp14859 {
ifres14855 = True


} else {
ifres14855 = False


}

ifres14854 = ifres14855


} else {
ifres14854 = False


}

var ifres14853 Obj

if True == ifres14854 {
ifres14853 = True


} else {
ifres14853 = False


}

ifres14852 = ifres14853


} else {
ifres14852 = False


}

var ifres14851 Obj

if True == ifres14852 {
ifres14851 = True


} else {
ifres14851 = False


}

ifres14850 = ifres14851


} else {
ifres14850 = False


}

var ifres14849 Obj

if True == ifres14850 {
ifres14849 = True


} else {
ifres14849 = False


}

ifres14848 = ifres14849


} else {
ifres14848 = False


}

if True == ifres14848 {
tmp14789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4984)
}
__typedArg0 := V4984
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14792)
}
__typedArg0 := tmp14792
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14791, tmp14793)
}
__typedArg0 := tmp14791
__typedArg1 := tmp14793
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp14794)
}
__typedArg0 := symtype
__typedArg1 := tmp14794
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp14846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14842 Obj

if True == tmp14846 {
tmp14844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14845 := Call(__e, PrimFunc(symshen_4special_2), tmp14844)


var ifres14843 Obj

if True == tmp14845 {
ifres14843 = True


} else {
ifres14843 = False


}

ifres14842 = ifres14843


} else {
ifres14842 = False


}

if True == ifres14842 {
tmp14795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14796 := MakeNative(func(__e *ControlFlow) {
Z4985 := __e.Get(1)
_ = Z4985
__e.TailApply(PrimFunc(symshen_4curry), Z4985)
return
}, 1)

tmp14797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14798 := Call(__e, PrimFunc(symmap), tmp14796, tmp14797)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14795, tmp14798)
}
__typedArg0 := tmp14795
__typedArg1 := tmp14798
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp14840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14836 Obj

if True == tmp14840 {
tmp14838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14839 := Call(__e, PrimFunc(symshen_4extraspecial_2), tmp14838)


var ifres14837 Obj

if True == tmp14839 {
ifres14837 = True


} else {
ifres14837 = False


}

ifres14836 = ifres14837


} else {
ifres14836 = False


}

if True == ifres14836 {
__e.Return(V4984)
return
} else {
tmp14834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14825 Obj

if True == tmp14834 {
tmp14832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14832)
}
__typedArg0 := tmp14832
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14827 Obj

if True == tmp14833 {
tmp14829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4984)
}
__typedArg0 := V4984
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14830)
}
__typedArg0 := tmp14830
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14828 Obj

if True == tmp14831 {
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

if True == ifres14825 {
tmp14799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14800)
}
__typedArg0 := tmp14800
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14801, Nil)
}
__typedArg0 := tmp14801
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14799, tmp14802)
}
__typedArg0 := tmp14799
__typedArg1 := tmp14802
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp14804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14804)
}
__typedArg0 := tmp14804
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14803, tmp14805)
}
__typedArg0 := tmp14803
__typedArg1 := tmp14805
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4curry) == __self {
__self1 = tmp14806
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4curry), tmp14806)
return


} else {
tmp14823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14814 Obj

if True == tmp14823 {
tmp14821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14821)
}
__typedArg0 := tmp14821
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14816 Obj

if True == tmp14822 {
tmp14818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14818)
}
__typedArg0 := tmp14818
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14819)
}
__typedArg0 := Nil
__typedArg1 := tmp14819
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14817 Obj

if True == tmp14820 {
ifres14817 = True


} else {
ifres14817 = False


}

ifres14816 = ifres14817


} else {
ifres14816 = False


}

var ifres14815 Obj

if True == ifres14816 {
ifres14815 = True


} else {
ifres14815 = False


}

ifres14814 = ifres14815


} else {
ifres14814 = False


}

if True == ifres14814 {
tmp14807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14808 := Call(__e, PrimFunc(symshen_4curry), tmp14807)


tmp14809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4984)
}
__typedArg0 := V4984
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14809)
}
__typedArg0 := tmp14809
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14811 := Call(__e, PrimFunc(symshen_4curry), tmp14810)


tmp14812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14811, Nil)
}
__typedArg0 := tmp14811
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp14808, tmp14812)
}
__typedArg0 := tmp14808
__typedArg1 := tmp14812
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V4984)
return
}


}


}


}


}


}


}, 1)

tmp14878 := Call(__e, ns2_1set, symshen_4curry, tmp14788)


_ = tmp14878

tmp14879 := MakeNative(func(__e *ControlFlow) {
V4986 := __e.Get(1)
_ = V4986
tmp14880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspecial_d)
}
__typedArg0 := symshen_4_dspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symelement_2), V4986, tmp14880)
return


}, 1)

tmp14881 := Call(__e, ns2_1set, symshen_4special_2, tmp14879)


_ = tmp14881

tmp14882 := MakeNative(func(__e *ControlFlow) {
V4987 := __e.Get(1)
_ = V4987
tmp14883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dextraspecial_d)
}
__typedArg0 := symshen_4_dextraspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symelement_2), V4987, tmp14883)
return


}, 1)

tmp14884 := Call(__e, ns2_1set, symshen_4extraspecial_2, tmp14882)


_ = tmp14884

tmp14885 := MakeNative(func(__e *ControlFlow) {
V4988 := __e.Get(1)
_ = V4988
V4989 := __e.Get(2)
_ = V4989
V4990 := __e.Get(3)
_ = V4990
V4991 := __e.Get(4)
_ = V4991
V4992 := __e.Get(5)
_ = V4992
V4993 := __e.Get(6)
_ = V4993
let__14012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4992)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4992
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__14012

tmp14891 := Call(__e, PrimFunc(symshen_4unlocked_2), V4991)


var ifres14887 Obj

if True == tmp14891 {
tmp14888 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14888

tmp14889 := Call(__e, PrimFunc(symshen_4maxinfexceeded_2))


tmp14890 := Call(__e, PrimFunc(symwhen), tmp14889, V4990, V4991, let__14012, V4993)


ifres14887 = tmp14890


} else {
ifres14887 = False


}

let__14013 := ifres14887
_ = let__14013

tmp14939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14013, False)
}
__typedArg0 := let__14013
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14939 {
tmp14920 := Call(__e, PrimFunc(symshen_4unlocked_2), V4991)


var ifres14892 Obj

if True == tmp14920 {
tmp14893 := Call(__e, PrimFunc(symshen_4lazyderef), V4988, V4990)


let__14015 := tmp14893
_ = let__14015

tmp14919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14015)
}
__typedArg0 := let__14015
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14894 Obj

if True == tmp14919 {
tmp14895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14015)
}
__typedArg0 := let__14015
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14016 := tmp14895
_ = let__14016

tmp14896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14015)
}
__typedArg0 := let__14015
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14897 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14896, V4990)


let__14017 := tmp14897
_ = let__14017

tmp14918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14017)
}
__typedArg0 := let__14017
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14898 Obj

if True == tmp14918 {
tmp14899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14017)
}
__typedArg0 := let__14017
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14018 := tmp14899
_ = let__14018

tmp14900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14017)
}
__typedArg0 := let__14017
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14901 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14900, V4990)


let__14019 := tmp14901
_ = let__14019

tmp14917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14019)
}
__typedArg0 := let__14019
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14902 Obj

if True == tmp14917 {
tmp14903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14019)
}
__typedArg0 := let__14019
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14020 := tmp14903
_ = let__14020

tmp14904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14019)
}
__typedArg0 := let__14019
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14905 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14904, V4990)


let__14021 := tmp14905
_ = let__14021

tmp14916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14021, Nil)
}
__typedArg0 := let__14021
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14906 Obj

if True == tmp14916 {
tmp14907 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14907

tmp14908 := Call(__e, PrimFunc(symshen_4deref), let__14018, V4990)


tmp14909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp14910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp14908, tmp14909)
}
__typedArg0 := tmp14908
__typedArg1 := tmp14909
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp14911 := MakeNative(func(__e *ControlFlow) {
tmp14912 := Call(__e, PrimFunc(symshen_4type_1theory_1enabled_2))


tmp14913 := MakeNative(func(__e *ControlFlow) {
tmp14914 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14016, let__14020, V4989, V4990, V4991, let__14012, V4993)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4990, V4991, let__14012, tmp14914)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14912, V4990, V4991, let__14012, tmp14913)
return


}, 0)

tmp14915 := Call(__e, PrimFunc(symwhen), tmp14910, V4990, V4991, let__14012, tmp14911)


ifres14906 = tmp14915


} else {
ifres14906 = False


}

ifres14902 = ifres14906


} else {
ifres14902 = False


}

ifres14898 = ifres14902


} else {
ifres14898 = False


}

ifres14894 = ifres14898


} else {
ifres14894 = False


}

ifres14892 = ifres14894


} else {
ifres14892 = False


}

let__14014 := ifres14892
_ = let__14014

tmp14937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14014, False)
}
__typedArg0 := let__14014
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14937 {
tmp14926 := Call(__e, PrimFunc(symshen_4unlocked_2), V4991)


var ifres14921 Obj

if True == tmp14926 {
tmp14922 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14922

tmp14923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspy_d)
}
__typedArg0 := symshen_4_dspy_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp14924 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4show), V4988, V4989, V4990, V4991, let__14012, V4993)
return
}, 0)

tmp14925 := Call(__e, PrimFunc(symwhen), tmp14923, V4990, V4991, let__14012, tmp14924)


ifres14921 = tmp14925


} else {
ifres14921 = False


}

let__14022 := ifres14921
_ = let__14022

tmp14935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14022, False)
}
__typedArg0 := let__14022
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14935 {
tmp14931 := Call(__e, PrimFunc(symshen_4unlocked_2), V4991)


var ifres14927 Obj

if True == tmp14931 {
tmp14928 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14928

tmp14929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp14930 := Call(__e, PrimFunc(symshen_4search_1user_1datatypes), V4988, V4989, tmp14929, V4990, V4991, let__14012, V4993)


ifres14927 = tmp14930


} else {
ifres14927 = False


}

let__14023 := ifres14927
_ = let__14023

tmp14933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14023, False)
}
__typedArg0 := let__14023
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp14933 {
__e.TailApply(PrimFunc(symshen_4unlock), V4991, let__14012)
return
} else {
__e.Return(let__14023)
return
}


} else {
__e.Return(let__14022)
return
}


} else {
__e.Return(let__14014)
return
}


} else {
__e.Return(let__14013)
return
}


}, 6)

tmp14940 := Call(__e, ns2_1set, symshen_4system_1S, tmp14885)


_ = tmp14940

tmp14941 := MakeNative(func(__e *ControlFlow) {
V5012 := __e.Get(1)
_ = V5012
V5013 := __e.Get(2)
_ = V5013
V5014 := __e.Get(3)
_ = V5014
V5015 := __e.Get(4)
_ = V5015
V5016 := __e.Get(5)
_ = V5016
V5017 := __e.Get(6)
_ = V5017
tmp14942 := Call(__e, PrimFunc(symshen_4line))


_ = tmp14942

tmp14943 := Call(__e, PrimFunc(symshen_4deref), V5012, V5014)


tmp14944 := Call(__e, PrimFunc(symshen_4show_1p), tmp14943)


_ = tmp14944

tmp14945 := Call(__e, PrimFunc(symnl), MakeInteger(2))


_ = tmp14945

tmp14946 := Call(__e, PrimFunc(symshen_4deref), V5013, V5014)


tmp14947 := Call(__e, PrimFunc(symshen_4show_1assumptions), tmp14946, MakeInteger(1))


_ = tmp14947

tmp14948 := Call(__e, PrimFunc(symshen_4pause_1for_1user))


_ = tmp14948

__e.Return(False)
return


}, 6)

tmp14949 := Call(__e, ns2_1set, symshen_4show, tmp14941)


_ = tmp14949

tmp14950 := MakeNative(func(__e *ControlFlow) {
tmp14951 := Call(__e, PrimFunc(syminferences))


let__14024 := tmp14951
_ = let__14024

tmp14953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(1), let__14024)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := let__14024
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14952 Obj

if True == tmp14953 {
ifres14952 = MakeString("")


} else {
ifres14952 = MakeString("s")


}

tmp14954 := Call(__e, PrimFunc(symshen_4app), ifres14952, MakeString(" \n?- "), symshen_4a)


tmp14956 := Call(__e, PrimFunc(symshen_4app), let__14024, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" inference"))
__typedS1, __typedOK1 := TypedString(tmp14954)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" inference")
__typedArg1 := tmp14954
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp14957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("____________________________________________________________ "))
__typedS1, __typedOK1 := TypedString(tmp14956)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("____________________________________________________________ ")
__typedArg1 := tmp14956
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp14958 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp14957, tmp14958)
return


}, 0)

tmp14959 := Call(__e, ns2_1set, symshen_4line, tmp14950)


_ = tmp14959

tmp14960 := MakeNative(func(__e *ControlFlow) {
V5019 := __e.Get(1)
_ = V5019
tmp14992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5019)
}
__typedArg0 := V5019
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14971 Obj

if True == tmp14992 {
tmp14990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5019)
}
__typedArg0 := V5019
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14990)
}
__typedArg0 := tmp14990
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14973 Obj

if True == tmp14991 {
tmp14987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5019)
}
__typedArg0 := V5019
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14987)
}
__typedArg0 := tmp14987
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp14988)
}
__typedArg0 := tmp14988
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres14975 Obj

if True == tmp14989 {
tmp14983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5019)
}
__typedArg0 := V5019
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14983)
}
__typedArg0 := tmp14983
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14984)
}
__typedArg0 := tmp14984
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp14985)
}
__typedArg0 := Nil
__typedArg1 := tmp14985
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14977 Obj

if True == tmp14986 {
tmp14979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5019)
}
__typedArg0 := V5019
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14979)
}
__typedArg0 := tmp14979
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp14982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp14980, tmp14981)
}
__typedArg0 := tmp14980
__typedArg1 := tmp14981
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres14978 Obj

if True == tmp14982 {
ifres14978 = True


} else {
ifres14978 = False


}

ifres14977 = ifres14978


} else {
ifres14977 = False


}

var ifres14976 Obj

if True == ifres14977 {
ifres14976 = True


} else {
ifres14976 = False


}

ifres14975 = ifres14976


} else {
ifres14975 = False


}

var ifres14974 Obj

if True == ifres14975 {
ifres14974 = True


} else {
ifres14974 = False


}

ifres14973 = ifres14974


} else {
ifres14973 = False


}

var ifres14972 Obj

if True == ifres14973 {
ifres14972 = True


} else {
ifres14972 = False


}

ifres14971 = ifres14972


} else {
ifres14971 = False


}

if True == ifres14971 {
tmp14961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5019)
}
__typedArg0 := V5019
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14962 := Call(__e, PrimFunc(symshen_4prterm), tmp14961)


_ = tmp14962

tmp14963 := Call(__e, PrimFunc(symstoutput))


tmp14964 := Call(__e, PrimFunc(sympr), MakeString(" : "), tmp14963)


_ = tmp14964

tmp14965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5019)
}
__typedArg0 := V5019
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp14965)
}
__typedArg0 := tmp14965
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14966)
}
__typedArg0 := tmp14966
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14968 := Call(__e, PrimFunc(symshen_4app), tmp14967, MakeString(""), symshen_4r)


tmp14969 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp14968, tmp14969)
return


} else {
__e.TailApply(PrimFunc(symshen_4prterm), V5019)
return
}


}, 1)

tmp14993 := Call(__e, ns2_1set, symshen_4show_1p, tmp14960)


_ = tmp14993

tmp14994 := MakeNative(func(__e *ControlFlow) {
V5020 := __e.Get(1)
_ = V5020
tmp15037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5020)
}
__typedArg0 := V5020
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15018 Obj

if True == tmp15037 {
tmp15035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5020)
}
__typedArg0 := V5020
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp15035)
}
__typedArg0 := symcons
__typedArg1 := tmp15035
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15020 Obj

if True == tmp15036 {
tmp15033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5020)
}
__typedArg0 := V5020
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15033)
}
__typedArg0 := tmp15033
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15022 Obj

if True == tmp15034 {
tmp15030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5020)
}
__typedArg0 := V5020
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15030)
}
__typedArg0 := tmp15030
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15031)
}
__typedArg0 := tmp15031
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15024 Obj

if True == tmp15032 {
tmp15026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5020)
}
__typedArg0 := V5020
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15026)
}
__typedArg0 := tmp15026
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15027)
}
__typedArg0 := tmp15027
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp15028)
}
__typedArg0 := Nil
__typedArg1 := tmp15028
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15025 Obj

if True == tmp15029 {
ifres15025 = True


} else {
ifres15025 = False


}

ifres15024 = ifres15025


} else {
ifres15024 = False


}

var ifres15023 Obj

if True == ifres15024 {
ifres15023 = True


} else {
ifres15023 = False


}

ifres15022 = ifres15023


} else {
ifres15022 = False


}

var ifres15021 Obj

if True == ifres15022 {
ifres15021 = True


} else {
ifres15021 = False


}

ifres15020 = ifres15021


} else {
ifres15020 = False


}

var ifres15019 Obj

if True == ifres15020 {
ifres15019 = True


} else {
ifres15019 = False


}

ifres15018 = ifres15019


} else {
ifres15018 = False


}

if True == ifres15018 {
tmp14995 := Call(__e, PrimFunc(symstoutput))


tmp14996 := Call(__e, PrimFunc(sympr), MakeString("["), tmp14995)


_ = tmp14996

tmp14997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5020)
}
__typedArg0 := V5020
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp14998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp14997)
}
__typedArg0 := tmp14997
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp14999 := Call(__e, PrimFunc(symshen_4prterm), tmp14998)


_ = tmp14999

tmp15000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5020)
}
__typedArg0 := V5020
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15001)
}
__typedArg0 := tmp15001
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15003 := Call(__e, PrimFunc(symshen_4prtl), tmp15002)


_ = tmp15003

tmp15004 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString("]"), tmp15004)
return


} else {
tmp15016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5020)
}
__typedArg0 := V5020
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15016 {
tmp15005 := Call(__e, PrimFunc(symstoutput))


tmp15006 := Call(__e, PrimFunc(sympr), MakeString("("), tmp15005)


_ = tmp15006

tmp15007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5020)
}
__typedArg0 := V5020
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15008 := Call(__e, PrimFunc(symshen_4prterm), tmp15007)


_ = tmp15008

tmp15009 := MakeNative(func(__e *ControlFlow) {
Z5021 := __e.Get(1)
_ = Z5021
tmp15010 := Call(__e, PrimFunc(symstoutput))


tmp15011 := Call(__e, PrimFunc(sympr), MakeString(" "), tmp15010)


_ = tmp15011

__e.TailApply(PrimFunc(symshen_4prterm), Z5021)
return


}, 1)

tmp15012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5020)
}
__typedArg0 := V5020
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15013 := Call(__e, PrimFunc(symmap), tmp15009, tmp15012)


_ = tmp15013

tmp15014 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString(")"), tmp15014)
return


} else {
__e.TailApply(PrimFunc(symprint), V5020)
return
}


}


}, 1)

tmp15038 := Call(__e, ns2_1set, symshen_4prterm, tmp14994)


_ = tmp15038

tmp15039 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V5022 := __self1
_ = V5022
tmp15072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5022)
}
__typedArg0 := Nil
__typedArg1 := V5022
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15072 {
__e.Return(MakeString(""))
return
} else {
tmp15070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5022)
}
__typedArg0 := V5022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15051 Obj

if True == tmp15070 {
tmp15068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5022)
}
__typedArg0 := V5022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp15068)
}
__typedArg0 := symcons
__typedArg1 := tmp15068
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15053 Obj

if True == tmp15069 {
tmp15066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5022)
}
__typedArg0 := V5022
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15066)
}
__typedArg0 := tmp15066
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15055 Obj

if True == tmp15067 {
tmp15063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5022)
}
__typedArg0 := V5022
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15063)
}
__typedArg0 := tmp15063
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15064)
}
__typedArg0 := tmp15064
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15057 Obj

if True == tmp15065 {
tmp15059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5022)
}
__typedArg0 := V5022
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15059)
}
__typedArg0 := tmp15059
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15060)
}
__typedArg0 := tmp15060
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp15061)
}
__typedArg0 := Nil
__typedArg1 := tmp15061
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15058 Obj

if True == tmp15062 {
ifres15058 = True


} else {
ifres15058 = False


}

ifres15057 = ifres15058


} else {
ifres15057 = False


}

var ifres15056 Obj

if True == ifres15057 {
ifres15056 = True


} else {
ifres15056 = False


}

ifres15055 = ifres15056


} else {
ifres15055 = False


}

var ifres15054 Obj

if True == ifres15055 {
ifres15054 = True


} else {
ifres15054 = False


}

ifres15053 = ifres15054


} else {
ifres15053 = False


}

var ifres15052 Obj

if True == ifres15053 {
ifres15052 = True


} else {
ifres15052 = False


}

ifres15051 = ifres15052


} else {
ifres15051 = False


}

if True == ifres15051 {
tmp15040 := Call(__e, PrimFunc(symstoutput))


tmp15041 := Call(__e, PrimFunc(sympr), MakeString(" "), tmp15040)


_ = tmp15041

tmp15042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5022)
}
__typedArg0 := V5022
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15042)
}
__typedArg0 := tmp15042
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15044 := Call(__e, PrimFunc(symshen_4prterm), tmp15043)


_ = tmp15044

tmp15045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5022)
}
__typedArg0 := V5022
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15045)
}
__typedArg0 := tmp15045
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15046)
}
__typedArg0 := tmp15046
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

if PrimFunc(symshen_4prtl) == __self {
__self1 = tmp15047
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4prtl), tmp15047)
return


} else {
tmp15048 := Call(__e, PrimFunc(symstoutput))


tmp15049 := Call(__e, PrimFunc(sympr), MakeString(" | "), tmp15048)


_ = tmp15049

__e.TailApply(PrimFunc(symshen_4prterm), V5022)
return


}


}


}, 1)

tmp15073 := Call(__e, ns2_1set, symshen_4prtl, tmp15039)


_ = tmp15073

tmp15074 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V5029 := __self1
_ = V5029
V5030 := __self2
_ = V5030
tmp15087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5029)
}
__typedArg0 := Nil
__typedArg1 := V5029
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15087 {
tmp15075 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString("\n> "), tmp15075)
return


} else {
tmp15085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5029)
}
__typedArg0 := V5029
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15085 {
tmp15076 := Call(__e, PrimFunc(symshen_4app), V5030, MakeString(". "), symshen_4a)


tmp15077 := Call(__e, PrimFunc(symstoutput))


tmp15078 := Call(__e, PrimFunc(sympr), tmp15076, tmp15077)


_ = tmp15078

tmp15079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5029)
}
__typedArg0 := V5029
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15080 := Call(__e, PrimFunc(symshen_4show_1p), tmp15079)


_ = tmp15080

tmp15081 := Call(__e, PrimFunc(symnl), MakeInteger(1))


_ = tmp15081

tmp15082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5029)
}
__typedArg0 := V5029
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4show_1assumptions) == __self {
__self1, __self2 = tmp15082, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5030)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5030
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4show_1assumptions), tmp15082, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5030)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5030
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

tmp15088 := Call(__e, ns2_1set, symshen_4show_1assumptions, tmp15074)


_ = tmp15088

tmp15089 := MakeNative(func(__e *ControlFlow) {
tmp15090 := Call(__e, PrimFunc(symstinput))


tmp15091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(tmp15090)
}
__typedArg0 := tmp15090
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

let__14025 := tmp15091
_ = let__14025

tmp15093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14025, MakeInteger(94))
}
__typedArg0 := let__14025
__typedArg1 := MakeInteger(94)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15093 {
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

tmp15094 := Call(__e, ns2_1set, symshen_4pause_1for_1user, tmp15089)


_ = tmp15094

tmp15095 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dshen_1type_1theory_1enabled_2_d)
}
__typedArg0 := symshen_4_dshen_1type_1theory_1enabled_2_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp15096 := Call(__e, ns2_1set, symshen_4type_1theory_1enabled_2, tmp15095)


_ = tmp15096

tmp15097 := MakeNative(func(__e *ControlFlow) {
tmp15099 := Call(__e, PrimFunc(syminferences))


tmp15100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dmaxinferences_d)
}
__typedArg0 := symshen_4_dmaxinferences_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(tmp15099)
__typedN1, __typedOK1 := TypedFloat64(tmp15100)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := tmp15099
__typedArg1 := tmp15100
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

tmp15102 := Call(__e, ns2_1set, symshen_4maxinfexceeded_2, tmp15097)


_ = tmp15102

tmp15103 := MakeNative(func(__e *ControlFlow) {
V5032 := __e.Get(1)
_ = V5032
V5033 := __e.Get(2)
_ = V5033
V5034 := __e.Get(3)
_ = V5034
V5035 := __e.Get(4)
_ = V5035
V5036 := __e.Get(5)
_ = V5036
V5037 := __e.Get(6)
_ = V5037
V5038 := __e.Get(7)
_ = V5038
let__14026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5037)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5037
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__14026

tmp15114 := Call(__e, PrimFunc(symshen_4unlocked_2), V5036)


var ifres15105 Obj

if True == tmp15114 {
tmp15106 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15106

tmp15107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspy_d)
}
__typedArg0 := symshen_4_dspy_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp15108 := MakeNative(func(__e *ControlFlow) {
tmp15109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp15110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5033, Nil)
}
__typedArg0 := V5033
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15109, tmp15110)
}
__typedArg0 := tmp15109
__typedArg1 := tmp15110
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5032, tmp15111)
}
__typedArg0 := V5032
__typedArg1 := tmp15111
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4show), tmp15112, V5034, V5035, V5036, let__14026, V5038)
return


}, 0)

tmp15113 := Call(__e, PrimFunc(symwhen), tmp15107, V5035, V5036, let__14026, tmp15108)


ifres15105 = tmp15113


} else {
ifres15105 = False


}

let__14027 := ifres15105
_ = let__14027

tmp15853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14027, False)
}
__typedArg0 := let__14027
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15853 {
tmp15122 := Call(__e, PrimFunc(symshen_4unlocked_2), V5036)


var ifres15115 Obj

if True == tmp15122 {
tmp15116 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15116

tmp15117 := Call(__e, PrimFunc(symshen_4lazyderef), V5032, V5035)


tmp15118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15117)
}
__typedArg0 := tmp15117
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

tmp15119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp15118)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp15118
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

tmp15120 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4primitive), V5032, V5033, V5035, V5036, let__14026, V5038)
return
}, 0)

tmp15121 := Call(__e, PrimFunc(symwhen), tmp15119, V5035, V5036, let__14026, tmp15120)


ifres15115 = tmp15121


} else {
ifres15115 = False


}

let__14028 := ifres15115
_ = let__14028

tmp15851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14028, False)
}
__typedArg0 := let__14028
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15851 {
tmp15126 := Call(__e, PrimFunc(symshen_4unlocked_2), V5036)


var ifres15123 Obj

if True == tmp15126 {
tmp15124 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15124

tmp15125 := Call(__e, PrimFunc(symshen_4by_1hypothesis), V5032, V5033, V5034, V5035, V5036, let__14026, V5038)


ifres15123 = tmp15125


} else {
ifres15123 = False


}

let__14029 := ifres15123
_ = let__14029

tmp15849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14029, False)
}
__typedArg0 := let__14029
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15849 {
tmp15140 := Call(__e, PrimFunc(symshen_4unlocked_2), V5036)


var ifres15127 Obj

if True == tmp15140 {
tmp15128 := Call(__e, PrimFunc(symshen_4lazyderef), V5032, V5035)


let__14031 := tmp15128
_ = let__14031

tmp15139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14031)
}
__typedArg0 := let__14031
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15129 Obj

if True == tmp15139 {
tmp15130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14031)
}
__typedArg0 := let__14031
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14032 := tmp15130
_ = let__14032

tmp15131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14031)
}
__typedArg0 := let__14031
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15132 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15131, V5035)


let__14033 := tmp15132
_ = let__14033

tmp15138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14033, Nil)
}
__typedArg0 := let__14033
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15133 Obj

if True == tmp15138 {
tmp15134 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15134

tmp15135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5033, Nil)
}
__typedArg0 := V5033
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15135)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15135
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15137 := Call(__e, PrimFunc(symshen_4lookupsig), let__14032, tmp15136, V5035, V5036, let__14026, V5038)


ifres15133 = tmp15137


} else {
ifres15133 = False


}

ifres15129 = ifres15133


} else {
ifres15129 = False


}

ifres15127 = ifres15129


} else {
ifres15127 = False


}

let__14030 := ifres15127
_ = let__14030

tmp15847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14030, False)
}
__typedArg0 := let__14030
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15847 {
tmp15166 := Call(__e, PrimFunc(symshen_4unlocked_2), V5036)


var ifres15141 Obj

if True == tmp15166 {
tmp15142 := Call(__e, PrimFunc(symshen_4lazyderef), V5032, V5035)


let__14035 := tmp15142
_ = let__14035

tmp15165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14035)
}
__typedArg0 := let__14035
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15143 Obj

if True == tmp15165 {
tmp15144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14035)
}
__typedArg0 := let__14035
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15145 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15144, V5035)


let__14036 := tmp15145
_ = let__14036

tmp15164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14036, symfn)
}
__typedArg0 := let__14036
__typedArg1 := symfn
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15146 Obj

if True == tmp15164 {
tmp15147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14035)
}
__typedArg0 := let__14035
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15148 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15147, V5035)


let__14037 := tmp15148
_ = let__14037

tmp15163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14037)
}
__typedArg0 := let__14037
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15149 Obj

if True == tmp15163 {
tmp15150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14037)
}
__typedArg0 := let__14037
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14038 := tmp15150
_ = let__14038

tmp15151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14037)
}
__typedArg0 := let__14037
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15152 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15151, V5035)


let__14039 := tmp15152
_ = let__14039

tmp15162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14039, Nil)
}
__typedArg0 := let__14039
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15153 Obj

if True == tmp15162 {
tmp15154 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15154

tmp15155 := Call(__e, PrimFunc(symshen_4deref), let__14038, V5035)


tmp15156 := Call(__e, PrimFunc(symarity), tmp15155)


tmp15157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp15156, MakeInteger(0))
}
__typedArg0 := tmp15156
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp15158 := MakeNative(func(__e *ControlFlow) {
tmp15159 := MakeNative(func(__e *ControlFlow) {
tmp15160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14038, Nil)
}
__typedArg0 := let__14038
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp15160, V5033, V5034, V5035, V5036, let__14026, V5038)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5035, V5036, let__14026, tmp15159)
return


}, 0)

tmp15161 := Call(__e, PrimFunc(symwhen), tmp15157, V5035, V5036, let__14026, tmp15158)


ifres15153 = tmp15161


} else {
ifres15153 = False


}

ifres15149 = ifres15153


} else {
ifres15149 = False


}

ifres15146 = ifres15149


} else {
ifres15146 = False


}

ifres15143 = ifres15146


} else {
ifres15143 = False


}

ifres15141 = ifres15143


} else {
ifres15141 = False


}

let__14034 := ifres15141
_ = let__14034

tmp15845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14034, False)
}
__typedArg0 := let__14034
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15845 {
tmp15186 := Call(__e, PrimFunc(symshen_4unlocked_2), V5036)


var ifres15167 Obj

if True == tmp15186 {
tmp15168 := Call(__e, PrimFunc(symshen_4lazyderef), V5032, V5035)


let__14041 := tmp15168
_ = let__14041

tmp15185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14041)
}
__typedArg0 := let__14041
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15169 Obj

if True == tmp15185 {
tmp15170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14041)
}
__typedArg0 := let__14041
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15171 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15170, V5035)


let__14042 := tmp15171
_ = let__14042

tmp15184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14042, symfn)
}
__typedArg0 := let__14042
__typedArg1 := symfn
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15172 Obj

if True == tmp15184 {
tmp15173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14041)
}
__typedArg0 := let__14041
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15174 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15173, V5035)


let__14043 := tmp15174
_ = let__14043

tmp15183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14043)
}
__typedArg0 := let__14043
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15175 Obj

if True == tmp15183 {
tmp15176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14043)
}
__typedArg0 := let__14043
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14044 := tmp15176
_ = let__14044

tmp15177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14043)
}
__typedArg0 := let__14043
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15178 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15177, V5035)


let__14045 := tmp15178
_ = let__14045

tmp15182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14045, Nil)
}
__typedArg0 := let__14045
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15179 Obj

if True == tmp15182 {
tmp15180 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15180

tmp15181 := Call(__e, PrimFunc(symshen_4lookupsig), let__14044, V5033, V5035, V5036, let__14026, V5038)


ifres15179 = tmp15181


} else {
ifres15179 = False


}

ifres15175 = ifres15179


} else {
ifres15175 = False


}

ifres15172 = ifres15175


} else {
ifres15172 = False


}

ifres15169 = ifres15172


} else {
ifres15169 = False


}

ifres15167 = ifres15169


} else {
ifres15167 = False


}

let__14040 := ifres15167
_ = let__14040

tmp15843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14040, False)
}
__typedArg0 := let__14040
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15843 {
tmp15213 := Call(__e, PrimFunc(symshen_4unlocked_2), V5036)


var ifres15187 Obj

if True == tmp15213 {
tmp15188 := Call(__e, PrimFunc(symshen_4lazyderef), V5032, V5035)


let__14047 := tmp15188
_ = let__14047

tmp15212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14047)
}
__typedArg0 := let__14047
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15189 Obj

if True == tmp15212 {
tmp15190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14047)
}
__typedArg0 := let__14047
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14048 := tmp15190
_ = let__14048

tmp15191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14047)
}
__typedArg0 := let__14047
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15192 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15191, V5035)


let__14049 := tmp15192
_ = let__14049

tmp15211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14049)
}
__typedArg0 := let__14049
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15193 Obj

if True == tmp15211 {
tmp15194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14049)
}
__typedArg0 := let__14049
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14050 := tmp15194
_ = let__14050

tmp15195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14049)
}
__typedArg0 := let__14049
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15196 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15195, V5035)


let__14051 := tmp15196
_ = let__14051

tmp15210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14051, Nil)
}
__typedArg0 := let__14051
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15197 Obj

if True == tmp15210 {
tmp15198 := Call(__e, PrimFunc(symshen_4newpv), V5035)


let__14052 := tmp15198
_ = let__14052

tmp15199 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15199

tmp15200 := Call(__e, PrimFunc(symshen_4lazyderef), let__14048, V5035)


tmp15201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15200)
}
__typedArg0 := tmp15200
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

tmp15202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp15201)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp15201
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

tmp15203 := MakeNative(func(__e *ControlFlow) {
tmp15204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5033, Nil)
}
__typedArg0 := V5033
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15204)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15204
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14052, tmp15205)
}
__typedArg0 := let__14052
__typedArg1 := tmp15205
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15207 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14050, let__14052, V5034, V5035, V5036, let__14026, V5038)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4lookupsig), let__14048, tmp15206, V5035, V5036, let__14026, tmp15207)
return


}, 0)

tmp15208 := Call(__e, PrimFunc(symwhen), tmp15202, V5035, V5036, let__14026, tmp15203)


tmp15209 := Call(__e, PrimFunc(symshen_4gc), V5035, tmp15208)


ifres15197 = tmp15209


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

let__14046 := ifres15187
_ = let__14046

tmp15841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14046, False)
}
__typedArg0 := let__14046
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15841 {
tmp15236 := Call(__e, PrimFunc(symshen_4unlocked_2), V5036)


var ifres15214 Obj

if True == tmp15236 {
tmp15215 := Call(__e, PrimFunc(symshen_4lazyderef), V5032, V5035)


let__14054 := tmp15215
_ = let__14054

tmp15235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14054)
}
__typedArg0 := let__14054
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15216 Obj

if True == tmp15235 {
tmp15217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14054)
}
__typedArg0 := let__14054
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14055 := tmp15217
_ = let__14055

tmp15218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14054)
}
__typedArg0 := let__14054
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15219 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15218, V5035)


let__14056 := tmp15219
_ = let__14056

tmp15234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14056)
}
__typedArg0 := let__14056
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15220 Obj

if True == tmp15234 {
tmp15221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14056)
}
__typedArg0 := let__14056
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14057 := tmp15221
_ = let__14057

tmp15222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14056)
}
__typedArg0 := let__14056
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15223 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15222, V5035)


let__14058 := tmp15223
_ = let__14058

tmp15233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14058, Nil)
}
__typedArg0 := let__14058
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15224 Obj

if True == tmp15233 {
tmp15225 := Call(__e, PrimFunc(symshen_4newpv), V5035)


let__14059 := tmp15225
_ = let__14059

tmp15226 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15226

tmp15227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5033, Nil)
}
__typedArg0 := V5033
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15227)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15227
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14059, tmp15228)
}
__typedArg0 := let__14059
__typedArg1 := tmp15228
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15230 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14057, let__14059, V5034, V5035, V5036, let__14026, V5038)
return
}, 0)

tmp15231 := Call(__e, PrimFunc(symshen_4system_1S_1h), let__14055, tmp15229, V5034, V5035, V5036, let__14026, tmp15230)


tmp15232 := Call(__e, PrimFunc(symshen_4gc), V5035, tmp15231)


ifres15224 = tmp15232


} else {
ifres15224 = False


}

ifres15220 = ifres15224


} else {
ifres15220 = False


}

ifres15216 = ifres15220


} else {
ifres15216 = False


}

ifres15214 = ifres15216


} else {
ifres15214 = False


}

let__14053 := ifres15214
_ = let__14053

tmp15839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14053, False)
}
__typedArg0 := let__14053
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15839 {
tmp15303 := Call(__e, PrimFunc(symshen_4unlocked_2), V5036)


var ifres15237 Obj

if True == tmp15303 {
tmp15238 := Call(__e, PrimFunc(symshen_4lazyderef), V5032, V5035)


let__14061 := tmp15238
_ = let__14061

tmp15302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14061)
}
__typedArg0 := let__14061
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15239 Obj

if True == tmp15302 {
tmp15240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14061)
}
__typedArg0 := let__14061
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15241 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15240, V5035)


let__14062 := tmp15241
_ = let__14062

tmp15301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14062, symcons)
}
__typedArg0 := let__14062
__typedArg1 := symcons
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15242 Obj

if True == tmp15301 {
tmp15243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14061)
}
__typedArg0 := let__14061
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15244 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15243, V5035)


let__14063 := tmp15244
_ = let__14063

tmp15300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14063)
}
__typedArg0 := let__14063
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15245 Obj

if True == tmp15300 {
tmp15246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14063)
}
__typedArg0 := let__14063
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14064 := tmp15246
_ = let__14064

tmp15247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14063)
}
__typedArg0 := let__14063
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15248 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15247, V5035)


let__14065 := tmp15248
_ = let__14065

tmp15299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14065)
}
__typedArg0 := let__14065
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15249 Obj

if True == tmp15299 {
tmp15250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14065)
}
__typedArg0 := let__14065
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14066 := tmp15250
_ = let__14066

tmp15251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14065)
}
__typedArg0 := let__14065
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15252 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15251, V5035)


let__14067 := tmp15252
_ = let__14067

tmp15298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14067, Nil)
}
__typedArg0 := let__14067
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15253 Obj

if True == tmp15298 {
tmp15254 := Call(__e, PrimFunc(symshen_4lazyderef), V5033, V5035)


let__14068 := tmp15254
_ = let__14068

tmp15255 := MakeNative(func(__e *ControlFlow) {
Z5083 := __e.Get(1)
_ = Z5083
tmp15256 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15256

tmp15257 := MakeNative(func(__e *ControlFlow) {
tmp15258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z5083, Nil)
}
__typedArg0 := Z5083
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp15258)
}
__typedArg0 := symlist
__typedArg1 := tmp15258
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14066, tmp15259, V5034, V5035, V5036, let__14026, V5038)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14064, Z5083, V5034, V5035, V5036, let__14026, tmp15257)
return


}, 1)

let__14069 := tmp15255
_ = let__14069

tmp15297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14068)
}
__typedArg0 := let__14068
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15260 Obj

if True == tmp15297 {
tmp15261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14068)
}
__typedArg0 := let__14068
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15262 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15261, V5035)


let__14070 := tmp15262
_ = let__14070

tmp15263 := MakeNative(func(__e *ControlFlow) {
tmp15264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14068)
}
__typedArg0 := let__14068
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15265 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15264, V5035)


let__14072 := tmp15265
_ = let__14072

tmp15266 := MakeNative(func(__e *ControlFlow) {
Z5088 := __e.Get(1)
_ = Z5088
__e.TailApply(let__14069, Z5088)
return
}, 1)

let__14073 := tmp15266
_ = let__14073

tmp15282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14072)
}
__typedArg0 := let__14072
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15282 {
tmp15267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14072)
}
__typedArg0 := let__14072
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14074 := tmp15267
_ = let__14074

tmp15268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14072)
}
__typedArg0 := let__14072
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15269 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15268, V5035)


let__14075 := tmp15269
_ = let__14075

tmp15270 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14073, let__14074)
return
}, 0)

let__14076 := tmp15270
_ = let__14076

tmp15274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14075, Nil)
}
__typedArg0 := let__14075
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15274 {
__e.TailApply(PrimFunc(symthaw), let__14076)
return
} else {
tmp15272 := Call(__e, PrimFunc(symshen_4pvar_2), let__14075)


if True == tmp15272 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14075, Nil, V5035, let__14076)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15280 := Call(__e, PrimFunc(symshen_4pvar_2), let__14072)


if True == tmp15280 {
tmp15275 := Call(__e, PrimFunc(symshen_4newpv), V5035)


let__14077 := tmp15275
_ = let__14077

tmp15276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14077, Nil)
}
__typedArg0 := let__14077
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15277 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14073, let__14077)
return
}, 0)

tmp15278 := Call(__e, PrimFunc(symshen_4bind_b), let__14072, tmp15276, V5035, tmp15277)


__e.TailApply(PrimFunc(symshen_4gc), V5035, tmp15278)
return


} else {
__e.Return(False)
return
}


}


}, 0)

let__14071 := tmp15263
_ = let__14071

tmp15288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14070, symlist)
}
__typedArg0 := let__14070
__typedArg1 := symlist
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15283 Obj

if True == tmp15288 {
tmp15284 := Call(__e, PrimFunc(symthaw), let__14071)


ifres15283 = tmp15284


} else {
tmp15287 := Call(__e, PrimFunc(symshen_4pvar_2), let__14070)


var ifres15285 Obj

if True == tmp15287 {
tmp15286 := Call(__e, PrimFunc(symshen_4bind_b), let__14070, symlist, V5035, let__14071)


ifres15285 = tmp15286


} else {
ifres15285 = False


}

ifres15283 = ifres15285


}

ifres15260 = ifres15283


} else {
tmp15296 := Call(__e, PrimFunc(symshen_4pvar_2), let__14068)


var ifres15289 Obj

if True == tmp15296 {
tmp15290 := Call(__e, PrimFunc(symshen_4newpv), V5035)


let__14078 := tmp15290
_ = let__14078

tmp15291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14078, Nil)
}
__typedArg0 := let__14078
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp15291)
}
__typedArg0 := symlist
__typedArg1 := tmp15291
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15293 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14069, let__14078)
return
}, 0)

tmp15294 := Call(__e, PrimFunc(symshen_4bind_b), let__14068, tmp15292, V5035, tmp15293)


tmp15295 := Call(__e, PrimFunc(symshen_4gc), V5035, tmp15294)


ifres15289 = tmp15295


} else {
ifres15289 = False


}

ifres15260 = ifres15289


}

ifres15253 = ifres15260


} else {
ifres15253 = False


}

ifres15249 = ifres15253


} else {
ifres15249 = False


}

ifres15245 = ifres15249


} else {
ifres15245 = False


}

ifres15242 = ifres15245


} else {
ifres15242 = False


}

ifres15239 = ifres15242


} else {
ifres15239 = False


}

ifres15237 = ifres15239


} else {
ifres15237 = False


}

let__14060 := ifres15237
_ = let__14060

tmp15837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14060, False)
}
__typedArg0 := let__14060
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15837 {
tmp15387 := Call(__e, PrimFunc(symshen_4unlocked_2), V5036)


var ifres15304 Obj

if True == tmp15387 {
tmp15305 := Call(__e, PrimFunc(symshen_4lazyderef), V5032, V5035)


let__14080 := tmp15305
_ = let__14080

tmp15386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14080)
}
__typedArg0 := let__14080
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15306 Obj

if True == tmp15386 {
tmp15307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14080)
}
__typedArg0 := let__14080
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15308 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15307, V5035)


let__14081 := tmp15308
_ = let__14081

tmp15385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14081, sym_8p)
}
__typedArg0 := let__14081
__typedArg1 := sym_8p
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15309 Obj

if True == tmp15385 {
tmp15310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14080)
}
__typedArg0 := let__14080
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15311 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15310, V5035)


let__14082 := tmp15311
_ = let__14082

tmp15384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14082)
}
__typedArg0 := let__14082
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15312 Obj

if True == tmp15384 {
tmp15313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14082)
}
__typedArg0 := let__14082
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14083 := tmp15313
_ = let__14083

tmp15314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14082)
}
__typedArg0 := let__14082
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15315 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15314, V5035)


let__14084 := tmp15315
_ = let__14084

tmp15383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14084)
}
__typedArg0 := let__14084
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15316 Obj

if True == tmp15383 {
tmp15317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14084)
}
__typedArg0 := let__14084
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14085 := tmp15317
_ = let__14085

tmp15318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14084)
}
__typedArg0 := let__14084
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15319 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15318, V5035)


let__14086 := tmp15319
_ = let__14086

tmp15382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14086, Nil)
}
__typedArg0 := let__14086
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15320 Obj

if True == tmp15382 {
tmp15321 := Call(__e, PrimFunc(symshen_4lazyderef), V5033, V5035)


let__14087 := tmp15321
_ = let__14087

tmp15322 := MakeNative(func(__e *ControlFlow) {
Z5104 := __e.Get(1)
_ = Z5104
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5105 := __e.Get(1)
_ = Z5105
tmp15323 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15323

tmp15324 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14085, Z5105, V5034, V5035, V5036, let__14026, V5038)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14083, Z5104, V5034, V5035, V5036, let__14026, tmp15324)
return


}, 1))
return
}, 1)

let__14088 := tmp15322
_ = let__14088

tmp15381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14087)
}
__typedArg0 := let__14087
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15325 Obj

if True == tmp15381 {
tmp15326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14087)
}
__typedArg0 := let__14087
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14089 := tmp15326
_ = let__14089

tmp15327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14087)
}
__typedArg0 := let__14087
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15328 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15327, V5035)


let__14090 := tmp15328
_ = let__14090

tmp15329 := MakeNative(func(__e *ControlFlow) {
Z5109 := __e.Get(1)
_ = Z5109
tmp15330 := Call(__e, let__14088, let__14089)


__e.TailApply(tmp15330, Z5109)
return


}, 1)

let__14091 := tmp15329
_ = let__14091

tmp15368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14090)
}
__typedArg0 := let__14090
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15331 Obj

if True == tmp15368 {
tmp15332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14090)
}
__typedArg0 := let__14090
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15333 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15332, V5035)


let__14092 := tmp15333
_ = let__14092

tmp15334 := MakeNative(func(__e *ControlFlow) {
tmp15335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14090)
}
__typedArg0 := let__14090
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15336 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15335, V5035)


let__14094 := tmp15336
_ = let__14094

tmp15337 := MakeNative(func(__e *ControlFlow) {
Z5114 := __e.Get(1)
_ = Z5114
__e.TailApply(let__14091, Z5114)
return
}, 1)

let__14095 := tmp15337
_ = let__14095

tmp15353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14094)
}
__typedArg0 := let__14094
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15353 {
tmp15338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14094)
}
__typedArg0 := let__14094
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14096 := tmp15338
_ = let__14096

tmp15339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14094)
}
__typedArg0 := let__14094
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15340 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15339, V5035)


let__14097 := tmp15340
_ = let__14097

tmp15341 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14095, let__14096)
return
}, 0)

let__14098 := tmp15341
_ = let__14098

tmp15345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14097, Nil)
}
__typedArg0 := let__14097
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15345 {
__e.TailApply(PrimFunc(symthaw), let__14098)
return
} else {
tmp15343 := Call(__e, PrimFunc(symshen_4pvar_2), let__14097)


if True == tmp15343 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14097, Nil, V5035, let__14098)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15351 := Call(__e, PrimFunc(symshen_4pvar_2), let__14094)


if True == tmp15351 {
tmp15346 := Call(__e, PrimFunc(symshen_4newpv), V5035)


let__14099 := tmp15346
_ = let__14099

tmp15347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14099, Nil)
}
__typedArg0 := let__14099
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15348 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14095, let__14099)
return
}, 0)

tmp15349 := Call(__e, PrimFunc(symshen_4bind_b), let__14094, tmp15347, V5035, tmp15348)


__e.TailApply(PrimFunc(symshen_4gc), V5035, tmp15349)
return


} else {
__e.Return(False)
return
}


}


}, 0)

let__14093 := tmp15334
_ = let__14093

tmp15359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14092, sym_d)
}
__typedArg0 := let__14092
__typedArg1 := sym_d
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15354 Obj

if True == tmp15359 {
tmp15355 := Call(__e, PrimFunc(symthaw), let__14093)


ifres15354 = tmp15355


} else {
tmp15358 := Call(__e, PrimFunc(symshen_4pvar_2), let__14092)


var ifres15356 Obj

if True == tmp15358 {
tmp15357 := Call(__e, PrimFunc(symshen_4bind_b), let__14092, sym_d, V5035, let__14093)


ifres15356 = tmp15357


} else {
ifres15356 = False


}

ifres15354 = ifres15356


}

ifres15331 = ifres15354


} else {
tmp15367 := Call(__e, PrimFunc(symshen_4pvar_2), let__14090)


var ifres15360 Obj

if True == tmp15367 {
tmp15361 := Call(__e, PrimFunc(symshen_4newpv), V5035)


let__14100 := tmp15361
_ = let__14100

tmp15362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14100, Nil)
}
__typedArg0 := let__14100
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp15362)
}
__typedArg0 := sym_d
__typedArg1 := tmp15362
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15364 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14091, let__14100)
return
}, 0)

tmp15365 := Call(__e, PrimFunc(symshen_4bind_b), let__14090, tmp15363, V5035, tmp15364)


tmp15366 := Call(__e, PrimFunc(symshen_4gc), V5035, tmp15365)


ifres15360 = tmp15366


} else {
ifres15360 = False


}

ifres15331 = ifres15360


}

ifres15325 = ifres15331


} else {
tmp15380 := Call(__e, PrimFunc(symshen_4pvar_2), let__14087)


var ifres15369 Obj

if True == tmp15380 {
tmp15370 := Call(__e, PrimFunc(symshen_4newpv), V5035)


let__14101 := tmp15370
_ = let__14101

tmp15371 := Call(__e, PrimFunc(symshen_4newpv), V5035)


let__14102 := tmp15371
_ = let__14102

tmp15372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14102, Nil)
}
__typedArg0 := let__14102
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_d, tmp15372)
}
__typedArg0 := sym_d
__typedArg1 := tmp15372
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14101, tmp15373)
}
__typedArg0 := let__14101
__typedArg1 := tmp15373
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15375 := MakeNative(func(__e *ControlFlow) {
tmp15376 := Call(__e, let__14088, let__14101)


__e.TailApply(tmp15376, let__14102)
return


}, 0)

tmp15377 := Call(__e, PrimFunc(symshen_4bind_b), let__14087, tmp15374, V5035, tmp15375)


tmp15378 := Call(__e, PrimFunc(symshen_4gc), V5035, tmp15377)


tmp15379 := Call(__e, PrimFunc(symshen_4gc), V5035, tmp15378)


ifres15369 = tmp15379


} else {
ifres15369 = False


}

ifres15325 = ifres15369


}

ifres15320 = ifres15325


} else {
ifres15320 = False


}

ifres15316 = ifres15320


} else {
ifres15316 = False


}

ifres15312 = ifres15316


} else {
ifres15312 = False


}

ifres15309 = ifres15312


} else {
ifres15309 = False


}

ifres15306 = ifres15309


} else {
ifres15306 = False


}

ifres15304 = ifres15306


} else {
ifres15304 = False


}

let__14079 := ifres15304
_ = let__14079

tmp15835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14079, False)
}
__typedArg0 := let__14079
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15835 {
tmp15454 := Call(__e, PrimFunc(symshen_4unlocked_2), V5036)


var ifres15388 Obj

if True == tmp15454 {
tmp15389 := Call(__e, PrimFunc(symshen_4lazyderef), V5032, V5035)


let__14104 := tmp15389
_ = let__14104

tmp15453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14104)
}
__typedArg0 := let__14104
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15390 Obj

if True == tmp15453 {
tmp15391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14104)
}
__typedArg0 := let__14104
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15392 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15391, V5035)


let__14105 := tmp15392
_ = let__14105

tmp15452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14105, sym_8v)
}
__typedArg0 := let__14105
__typedArg1 := sym_8v
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15393 Obj

if True == tmp15452 {
tmp15394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14104)
}
__typedArg0 := let__14104
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15395 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15394, V5035)


let__14106 := tmp15395
_ = let__14106

tmp15451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14106)
}
__typedArg0 := let__14106
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15396 Obj

if True == tmp15451 {
tmp15397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14106)
}
__typedArg0 := let__14106
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14107 := tmp15397
_ = let__14107

tmp15398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14106)
}
__typedArg0 := let__14106
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15399 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15398, V5035)


let__14108 := tmp15399
_ = let__14108

tmp15450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14108)
}
__typedArg0 := let__14108
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15400 Obj

if True == tmp15450 {
tmp15401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14108)
}
__typedArg0 := let__14108
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14109 := tmp15401
_ = let__14109

tmp15402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14108)
}
__typedArg0 := let__14108
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15403 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15402, V5035)


let__14110 := tmp15403
_ = let__14110

tmp15449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14110, Nil)
}
__typedArg0 := let__14110
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15404 Obj

if True == tmp15449 {
tmp15405 := Call(__e, PrimFunc(symshen_4lazyderef), V5033, V5035)


let__14111 := tmp15405
_ = let__14111

tmp15406 := MakeNative(func(__e *ControlFlow) {
Z5132 := __e.Get(1)
_ = Z5132
tmp15407 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15407

tmp15408 := MakeNative(func(__e *ControlFlow) {
tmp15409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z5132, Nil)
}
__typedArg0 := Z5132
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp15409)
}
__typedArg0 := symvector
__typedArg1 := tmp15409
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14109, tmp15410, V5034, V5035, V5036, let__14026, V5038)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14107, Z5132, V5034, V5035, V5036, let__14026, tmp15408)
return


}, 1)

let__14112 := tmp15406
_ = let__14112

tmp15448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14111)
}
__typedArg0 := let__14111
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15411 Obj

if True == tmp15448 {
tmp15412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14111)
}
__typedArg0 := let__14111
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15413 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15412, V5035)


let__14113 := tmp15413
_ = let__14113

tmp15414 := MakeNative(func(__e *ControlFlow) {
tmp15415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14111)
}
__typedArg0 := let__14111
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15416 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15415, V5035)


let__14115 := tmp15416
_ = let__14115

tmp15417 := MakeNative(func(__e *ControlFlow) {
Z5137 := __e.Get(1)
_ = Z5137
__e.TailApply(let__14112, Z5137)
return
}, 1)

let__14116 := tmp15417
_ = let__14116

tmp15433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14115)
}
__typedArg0 := let__14115
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15433 {
tmp15418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14115)
}
__typedArg0 := let__14115
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14117 := tmp15418
_ = let__14117

tmp15419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14115)
}
__typedArg0 := let__14115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15420 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15419, V5035)


let__14118 := tmp15420
_ = let__14118

tmp15421 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14116, let__14117)
return
}, 0)

let__14119 := tmp15421
_ = let__14119

tmp15425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14118, Nil)
}
__typedArg0 := let__14118
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15425 {
__e.TailApply(PrimFunc(symthaw), let__14119)
return
} else {
tmp15423 := Call(__e, PrimFunc(symshen_4pvar_2), let__14118)


if True == tmp15423 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14118, Nil, V5035, let__14119)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15431 := Call(__e, PrimFunc(symshen_4pvar_2), let__14115)


if True == tmp15431 {
tmp15426 := Call(__e, PrimFunc(symshen_4newpv), V5035)


let__14120 := tmp15426
_ = let__14120

tmp15427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14120, Nil)
}
__typedArg0 := let__14120
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15428 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14116, let__14120)
return
}, 0)

tmp15429 := Call(__e, PrimFunc(symshen_4bind_b), let__14115, tmp15427, V5035, tmp15428)


__e.TailApply(PrimFunc(symshen_4gc), V5035, tmp15429)
return


} else {
__e.Return(False)
return
}


}


}, 0)

let__14114 := tmp15414
_ = let__14114

tmp15439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14113, symvector)
}
__typedArg0 := let__14113
__typedArg1 := symvector
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15434 Obj

if True == tmp15439 {
tmp15435 := Call(__e, PrimFunc(symthaw), let__14114)


ifres15434 = tmp15435


} else {
tmp15438 := Call(__e, PrimFunc(symshen_4pvar_2), let__14113)


var ifres15436 Obj

if True == tmp15438 {
tmp15437 := Call(__e, PrimFunc(symshen_4bind_b), let__14113, symvector, V5035, let__14114)


ifres15436 = tmp15437


} else {
ifres15436 = False


}

ifres15434 = ifres15436


}

ifres15411 = ifres15434


} else {
tmp15447 := Call(__e, PrimFunc(symshen_4pvar_2), let__14111)


var ifres15440 Obj

if True == tmp15447 {
tmp15441 := Call(__e, PrimFunc(symshen_4newpv), V5035)


let__14121 := tmp15441
_ = let__14121

tmp15442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14121, Nil)
}
__typedArg0 := let__14121
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp15442)
}
__typedArg0 := symvector
__typedArg1 := tmp15442
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15444 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14112, let__14121)
return
}, 0)

tmp15445 := Call(__e, PrimFunc(symshen_4bind_b), let__14111, tmp15443, V5035, tmp15444)


tmp15446 := Call(__e, PrimFunc(symshen_4gc), V5035, tmp15445)


ifres15440 = tmp15446


} else {
ifres15440 = False


}

ifres15411 = ifres15440


}

ifres15404 = ifres15411


} else {
ifres15404 = False


}

ifres15400 = ifres15404


} else {
ifres15400 = False


}

ifres15396 = ifres15400


} else {
ifres15396 = False


}

ifres15393 = ifres15396


} else {
ifres15393 = False


}

ifres15390 = ifres15393


} else {
ifres15390 = False


}

ifres15388 = ifres15390


} else {
ifres15388 = False


}

let__14103 := ifres15388
_ = let__14103

tmp15833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14103, False)
}
__typedArg0 := let__14103
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15833 {
tmp15487 := Call(__e, PrimFunc(symshen_4unlocked_2), V5036)


var ifres15455 Obj

if True == tmp15487 {
tmp15456 := Call(__e, PrimFunc(symshen_4lazyderef), V5032, V5035)


let__14123 := tmp15456
_ = let__14123

tmp15486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14123)
}
__typedArg0 := let__14123
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15457 Obj

if True == tmp15486 {
tmp15458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14123)
}
__typedArg0 := let__14123
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15459 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15458, V5035)


let__14124 := tmp15459
_ = let__14124

tmp15485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14124, sym_8s)
}
__typedArg0 := let__14124
__typedArg1 := sym_8s
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15460 Obj

if True == tmp15485 {
tmp15461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14123)
}
__typedArg0 := let__14123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15462 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15461, V5035)


let__14125 := tmp15462
_ = let__14125

tmp15484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14125)
}
__typedArg0 := let__14125
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15463 Obj

if True == tmp15484 {
tmp15464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14125)
}
__typedArg0 := let__14125
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14126 := tmp15464
_ = let__14126

tmp15465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14125)
}
__typedArg0 := let__14125
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15466 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15465, V5035)


let__14127 := tmp15466
_ = let__14127

tmp15483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14127)
}
__typedArg0 := let__14127
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15467 Obj

if True == tmp15483 {
tmp15468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14127)
}
__typedArg0 := let__14127
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14128 := tmp15468
_ = let__14128

tmp15469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14127)
}
__typedArg0 := let__14127
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15470 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15469, V5035)


let__14129 := tmp15470
_ = let__14129

tmp15482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14129, Nil)
}
__typedArg0 := let__14129
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15471 Obj

if True == tmp15482 {
tmp15472 := Call(__e, PrimFunc(symshen_4lazyderef), V5033, V5035)


let__14130 := tmp15472
_ = let__14130

tmp15473 := MakeNative(func(__e *ControlFlow) {
tmp15474 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15474

tmp15475 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14128, symstring, V5034, V5035, V5036, let__14026, V5038)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14126, symstring, V5034, V5035, V5036, let__14026, tmp15475)
return


}, 0)

let__14131 := tmp15473
_ = let__14131

tmp15481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14130, symstring)
}
__typedArg0 := let__14130
__typedArg1 := symstring
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15476 Obj

if True == tmp15481 {
tmp15477 := Call(__e, PrimFunc(symthaw), let__14131)


ifres15476 = tmp15477


} else {
tmp15480 := Call(__e, PrimFunc(symshen_4pvar_2), let__14130)


var ifres15478 Obj

if True == tmp15480 {
tmp15479 := Call(__e, PrimFunc(symshen_4bind_b), let__14130, symstring, V5035, let__14131)


ifres15478 = tmp15479


} else {
ifres15478 = False


}

ifres15476 = ifres15478


}

ifres15471 = ifres15476


} else {
ifres15471 = False


}

ifres15467 = ifres15471


} else {
ifres15467 = False


}

ifres15463 = ifres15467


} else {
ifres15463 = False


}

ifres15460 = ifres15463


} else {
ifres15460 = False


}

ifres15457 = ifres15460


} else {
ifres15457 = False


}

ifres15455 = ifres15457


} else {
ifres15455 = False


}

let__14122 := ifres15455
_ = let__14122

tmp15831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14122, False)
}
__typedArg0 := let__14122
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15831 {
tmp15587 := Call(__e, PrimFunc(symshen_4unlocked_2), V5036)


var ifres15488 Obj

if True == tmp15587 {
tmp15489 := Call(__e, PrimFunc(symshen_4lazyderef), V5032, V5035)


let__14133 := tmp15489
_ = let__14133

tmp15586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14133)
}
__typedArg0 := let__14133
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15490 Obj

if True == tmp15586 {
tmp15491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14133)
}
__typedArg0 := let__14133
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15492 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15491, V5035)


let__14134 := tmp15492
_ = let__14134

tmp15585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14134, symlambda)
}
__typedArg0 := let__14134
__typedArg1 := symlambda
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15493 Obj

if True == tmp15585 {
tmp15494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14133)
}
__typedArg0 := let__14133
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15495 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15494, V5035)


let__14135 := tmp15495
_ = let__14135

tmp15584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14135)
}
__typedArg0 := let__14135
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15496 Obj

if True == tmp15584 {
tmp15497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14135)
}
__typedArg0 := let__14135
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14136 := tmp15497
_ = let__14136

tmp15498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14135)
}
__typedArg0 := let__14135
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15499 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15498, V5035)


let__14137 := tmp15499
_ = let__14137

tmp15583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14137)
}
__typedArg0 := let__14137
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15500 Obj

if True == tmp15583 {
tmp15501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14137)
}
__typedArg0 := let__14137
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14138 := tmp15501
_ = let__14138

tmp15502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14137)
}
__typedArg0 := let__14137
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15503 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15502, V5035)


let__14139 := tmp15503
_ = let__14139

tmp15582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14139, Nil)
}
__typedArg0 := let__14139
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15504 Obj

if True == tmp15582 {
tmp15505 := Call(__e, PrimFunc(symshen_4lazyderef), V5033, V5035)


let__14140 := tmp15505
_ = let__14140

tmp15506 := MakeNative(func(__e *ControlFlow) {
Z5163 := __e.Get(1)
_ = Z5163
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5164 := __e.Get(1)
_ = Z5164
tmp15507 := Call(__e, PrimFunc(symshen_4newpv), V5035)


let__14142 := tmp15507
_ = let__14142

tmp15508 := Call(__e, PrimFunc(symshen_4newpv), V5035)


let__14143 := tmp15508
_ = let__14143

tmp15509 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15509

tmp15510 := Call(__e, PrimFunc(symshen_4lazyderef), let__14136, V5035)


tmp15511 := Call(__e, PrimFunc(symshen_4freshterm), tmp15510)


tmp15512 := MakeNative(func(__e *ControlFlow) {
tmp15513 := Call(__e, PrimFunc(symshen_4lazyderef), let__14136, V5035)


tmp15514 := Call(__e, PrimFunc(symshen_4deref), let__14143, V5035)


tmp15515 := Call(__e, PrimFunc(symshen_4deref), let__14138, V5035)


tmp15516 := Call(__e, PrimFunc(symshen_4beta), tmp15513, tmp15514, tmp15515)


tmp15517 := MakeNative(func(__e *ControlFlow) {
tmp15518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp15519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z5163, Nil)
}
__typedArg0 := Z5163
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15518, tmp15519)
}
__typedArg0 := tmp15518
__typedArg1 := tmp15519
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14143, tmp15520)
}
__typedArg0 := let__14143
__typedArg1 := tmp15520
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15521, V5034)
}
__typedArg0 := tmp15521
__typedArg1 := V5034
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14142, Z5164, tmp15522, V5035, V5036, let__14026, V5038)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14142, tmp15516, V5035, V5036, let__14026, tmp15517)
return


}, 0)

tmp15523 := Call(__e, PrimFunc(symbind), let__14143, tmp15511, V5035, V5036, let__14026, tmp15512)


tmp15524 := Call(__e, PrimFunc(symshen_4gc), V5035, tmp15523)


__e.TailApply(PrimFunc(symshen_4gc), V5035, tmp15524)
return


}, 1))
return
}, 1)

let__14141 := tmp15506
_ = let__14141

tmp15581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14140)
}
__typedArg0 := let__14140
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15525 Obj

if True == tmp15581 {
tmp15526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14140)
}
__typedArg0 := let__14140
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14144 := tmp15526
_ = let__14144

tmp15527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14140)
}
__typedArg0 := let__14140
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15528 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15527, V5035)


let__14145 := tmp15528
_ = let__14145

tmp15529 := MakeNative(func(__e *ControlFlow) {
Z5170 := __e.Get(1)
_ = Z5170
tmp15530 := Call(__e, let__14141, let__14144)


__e.TailApply(tmp15530, Z5170)
return


}, 1)

let__14146 := tmp15529
_ = let__14146

tmp15568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14145)
}
__typedArg0 := let__14145
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15531 Obj

if True == tmp15568 {
tmp15532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14145)
}
__typedArg0 := let__14145
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15533 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15532, V5035)


let__14147 := tmp15533
_ = let__14147

tmp15534 := MakeNative(func(__e *ControlFlow) {
tmp15535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14145)
}
__typedArg0 := let__14145
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15536 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15535, V5035)


let__14149 := tmp15536
_ = let__14149

tmp15537 := MakeNative(func(__e *ControlFlow) {
Z5175 := __e.Get(1)
_ = Z5175
__e.TailApply(let__14146, Z5175)
return
}, 1)

let__14150 := tmp15537
_ = let__14150

tmp15553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14149)
}
__typedArg0 := let__14149
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15553 {
tmp15538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14149)
}
__typedArg0 := let__14149
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14151 := tmp15538
_ = let__14151

tmp15539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14149)
}
__typedArg0 := let__14149
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15540 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15539, V5035)


let__14152 := tmp15540
_ = let__14152

tmp15541 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14150, let__14151)
return
}, 0)

let__14153 := tmp15541
_ = let__14153

tmp15545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14152, Nil)
}
__typedArg0 := let__14152
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15545 {
__e.TailApply(PrimFunc(symthaw), let__14153)
return
} else {
tmp15543 := Call(__e, PrimFunc(symshen_4pvar_2), let__14152)


if True == tmp15543 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14152, Nil, V5035, let__14153)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15551 := Call(__e, PrimFunc(symshen_4pvar_2), let__14149)


if True == tmp15551 {
tmp15546 := Call(__e, PrimFunc(symshen_4newpv), V5035)


let__14154 := tmp15546
_ = let__14154

tmp15547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14154, Nil)
}
__typedArg0 := let__14154
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15548 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14150, let__14154)
return
}, 0)

tmp15549 := Call(__e, PrimFunc(symshen_4bind_b), let__14149, tmp15547, V5035, tmp15548)


__e.TailApply(PrimFunc(symshen_4gc), V5035, tmp15549)
return


} else {
__e.Return(False)
return
}


}


}, 0)

let__14148 := tmp15534
_ = let__14148

tmp15559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14147, sym_1_1_6)
}
__typedArg0 := let__14147
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15554 Obj

if True == tmp15559 {
tmp15555 := Call(__e, PrimFunc(symthaw), let__14148)


ifres15554 = tmp15555


} else {
tmp15558 := Call(__e, PrimFunc(symshen_4pvar_2), let__14147)


var ifres15556 Obj

if True == tmp15558 {
tmp15557 := Call(__e, PrimFunc(symshen_4bind_b), let__14147, sym_1_1_6, V5035, let__14148)


ifres15556 = tmp15557


} else {
ifres15556 = False


}

ifres15554 = ifres15556


}

ifres15531 = ifres15554


} else {
tmp15567 := Call(__e, PrimFunc(symshen_4pvar_2), let__14145)


var ifres15560 Obj

if True == tmp15567 {
tmp15561 := Call(__e, PrimFunc(symshen_4newpv), V5035)


let__14155 := tmp15561
_ = let__14155

tmp15562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14155, Nil)
}
__typedArg0 := let__14155
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15562)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15562
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15564 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14146, let__14155)
return
}, 0)

tmp15565 := Call(__e, PrimFunc(symshen_4bind_b), let__14145, tmp15563, V5035, tmp15564)


tmp15566 := Call(__e, PrimFunc(symshen_4gc), V5035, tmp15565)


ifres15560 = tmp15566


} else {
ifres15560 = False


}

ifres15531 = ifres15560


}

ifres15525 = ifres15531


} else {
tmp15580 := Call(__e, PrimFunc(symshen_4pvar_2), let__14140)


var ifres15569 Obj

if True == tmp15580 {
tmp15570 := Call(__e, PrimFunc(symshen_4newpv), V5035)


let__14156 := tmp15570
_ = let__14156

tmp15571 := Call(__e, PrimFunc(symshen_4newpv), V5035)


let__14157 := tmp15571
_ = let__14157

tmp15572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14157, Nil)
}
__typedArg0 := let__14157
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_1_6, tmp15572)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp15572
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14156, tmp15573)
}
__typedArg0 := let__14156
__typedArg1 := tmp15573
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15575 := MakeNative(func(__e *ControlFlow) {
tmp15576 := Call(__e, let__14141, let__14156)


__e.TailApply(tmp15576, let__14157)
return


}, 0)

tmp15577 := Call(__e, PrimFunc(symshen_4bind_b), let__14140, tmp15574, V5035, tmp15575)


tmp15578 := Call(__e, PrimFunc(symshen_4gc), V5035, tmp15577)


tmp15579 := Call(__e, PrimFunc(symshen_4gc), V5035, tmp15578)


ifres15569 = tmp15579


} else {
ifres15569 = False


}

ifres15525 = ifres15569


}

ifres15504 = ifres15525


} else {
ifres15504 = False


}

ifres15500 = ifres15504


} else {
ifres15500 = False


}

ifres15496 = ifres15500


} else {
ifres15496 = False


}

ifres15493 = ifres15496


} else {
ifres15493 = False


}

ifres15490 = ifres15493


} else {
ifres15490 = False


}

ifres15488 = ifres15490


} else {
ifres15488 = False


}

let__14132 := ifres15488
_ = let__14132

tmp15829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14132, False)
}
__typedArg0 := let__14132
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15829 {
tmp15637 := Call(__e, PrimFunc(symshen_4unlocked_2), V5036)


var ifres15588 Obj

if True == tmp15637 {
tmp15589 := Call(__e, PrimFunc(symshen_4lazyderef), V5032, V5035)


let__14159 := tmp15589
_ = let__14159

tmp15636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14159)
}
__typedArg0 := let__14159
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15590 Obj

if True == tmp15636 {
tmp15591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14159)
}
__typedArg0 := let__14159
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15592 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15591, V5035)


let__14160 := tmp15592
_ = let__14160

tmp15635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14160, symlet)
}
__typedArg0 := let__14160
__typedArg1 := symlet
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15593 Obj

if True == tmp15635 {
tmp15594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14159)
}
__typedArg0 := let__14159
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15595 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15594, V5035)


let__14161 := tmp15595
_ = let__14161

tmp15634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14161)
}
__typedArg0 := let__14161
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15596 Obj

if True == tmp15634 {
tmp15597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14161)
}
__typedArg0 := let__14161
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14162 := tmp15597
_ = let__14162

tmp15598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14161)
}
__typedArg0 := let__14161
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15599 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15598, V5035)


let__14163 := tmp15599
_ = let__14163

tmp15633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14163)
}
__typedArg0 := let__14163
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15600 Obj

if True == tmp15633 {
tmp15601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14163)
}
__typedArg0 := let__14163
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14164 := tmp15601
_ = let__14164

tmp15602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14163)
}
__typedArg0 := let__14163
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15603 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15602, V5035)


let__14165 := tmp15603
_ = let__14165

tmp15632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14165)
}
__typedArg0 := let__14165
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15604 Obj

if True == tmp15632 {
tmp15605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14165)
}
__typedArg0 := let__14165
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14166 := tmp15605
_ = let__14166

tmp15606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14165)
}
__typedArg0 := let__14165
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15607 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15606, V5035)


let__14167 := tmp15607
_ = let__14167

tmp15631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14167, Nil)
}
__typedArg0 := let__14167
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15608 Obj

if True == tmp15631 {
tmp15609 := Call(__e, PrimFunc(symshen_4newpv), V5035)


let__14168 := tmp15609
_ = let__14168

tmp15610 := Call(__e, PrimFunc(symshen_4newpv), V5035)


let__14169 := tmp15610
_ = let__14169

tmp15611 := Call(__e, PrimFunc(symshen_4newpv), V5035)


let__14170 := tmp15611
_ = let__14170

tmp15612 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15612

tmp15613 := MakeNative(func(__e *ControlFlow) {
tmp15614 := Call(__e, PrimFunc(symshen_4lazyderef), let__14162, V5035)


tmp15615 := Call(__e, PrimFunc(symshen_4freshterm), tmp15614)


tmp15616 := MakeNative(func(__e *ControlFlow) {
tmp15617 := Call(__e, PrimFunc(symshen_4lazyderef), let__14162, V5035)


tmp15618 := Call(__e, PrimFunc(symshen_4lazyderef), let__14169, V5035)


tmp15619 := Call(__e, PrimFunc(symshen_4lazyderef), let__14166, V5035)


tmp15620 := Call(__e, PrimFunc(symshen_4beta), tmp15617, tmp15618, tmp15619)


tmp15621 := MakeNative(func(__e *ControlFlow) {
tmp15622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp15623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14170, Nil)
}
__typedArg0 := let__14170
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15622, tmp15623)
}
__typedArg0 := tmp15622
__typedArg1 := tmp15623
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14169, tmp15624)
}
__typedArg0 := let__14169
__typedArg1 := tmp15624
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15625, V5034)
}
__typedArg0 := tmp15625
__typedArg1 := V5034
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14168, V5033, tmp15626, V5035, V5036, let__14026, V5038)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14168, tmp15620, V5035, V5036, let__14026, tmp15621)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14169, tmp15615, V5035, V5036, let__14026, tmp15616)
return


}, 0)

tmp15627 := Call(__e, PrimFunc(symshen_4system_1S_1h), let__14164, let__14170, V5034, V5035, V5036, let__14026, tmp15613)


tmp15628 := Call(__e, PrimFunc(symshen_4gc), V5035, tmp15627)


tmp15629 := Call(__e, PrimFunc(symshen_4gc), V5035, tmp15628)


tmp15630 := Call(__e, PrimFunc(symshen_4gc), V5035, tmp15629)


ifres15608 = tmp15630


} else {
ifres15608 = False


}

ifres15604 = ifres15608


} else {
ifres15604 = False


}

ifres15600 = ifres15604


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

let__14158 := ifres15588
_ = let__14158

tmp15827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14158, False)
}
__typedArg0 := let__14158
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15827 {
tmp15707 := Call(__e, PrimFunc(symshen_4unlocked_2), V5036)


var ifres15638 Obj

if True == tmp15707 {
tmp15639 := Call(__e, PrimFunc(symshen_4lazyderef), V5032, V5035)


let__14172 := tmp15639
_ = let__14172

tmp15706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14172)
}
__typedArg0 := let__14172
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15640 Obj

if True == tmp15706 {
tmp15641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14172)
}
__typedArg0 := let__14172
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15642 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15641, V5035)


let__14173 := tmp15642
_ = let__14173

tmp15705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14173, symopen)
}
__typedArg0 := let__14173
__typedArg1 := symopen
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15643 Obj

if True == tmp15705 {
tmp15644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14172)
}
__typedArg0 := let__14172
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15645 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15644, V5035)


let__14174 := tmp15645
_ = let__14174

tmp15704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14174)
}
__typedArg0 := let__14174
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15646 Obj

if True == tmp15704 {
tmp15647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14174)
}
__typedArg0 := let__14174
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14175 := tmp15647
_ = let__14175

tmp15648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14174)
}
__typedArg0 := let__14174
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15649 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15648, V5035)


let__14176 := tmp15649
_ = let__14176

tmp15703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14176)
}
__typedArg0 := let__14176
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15650 Obj

if True == tmp15703 {
tmp15651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14176)
}
__typedArg0 := let__14176
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14177 := tmp15651
_ = let__14177

tmp15652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14176)
}
__typedArg0 := let__14176
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15653 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15652, V5035)


let__14178 := tmp15653
_ = let__14178

tmp15702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14178, Nil)
}
__typedArg0 := let__14178
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15654 Obj

if True == tmp15702 {
tmp15655 := Call(__e, PrimFunc(symshen_4lazyderef), V5033, V5035)


let__14179 := tmp15655
_ = let__14179

tmp15656 := MakeNative(func(__e *ControlFlow) {
Z5206 := __e.Get(1)
_ = Z5206
tmp15657 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15657

tmp15658 := MakeNative(func(__e *ControlFlow) {
tmp15659 := Call(__e, PrimFunc(symshen_4lazyderef), Z5206, V5035)


tmp15660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symout, Nil)
}
__typedArg0 := symout
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, tmp15660)
}
__typedArg0 := symin
__typedArg1 := tmp15660
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15662 := Call(__e, PrimFunc(symelement_2), tmp15659, tmp15661)


tmp15663 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14175, symstring, V5034, V5035, V5036, let__14026, V5038)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp15662, V5035, V5036, let__14026, tmp15663)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), let__14177, Z5206, V5035, V5036, let__14026, tmp15658)
return


}, 1)

let__14180 := tmp15656
_ = let__14180

tmp15701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14179)
}
__typedArg0 := let__14179
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15664 Obj

if True == tmp15701 {
tmp15665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14179)
}
__typedArg0 := let__14179
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15666 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15665, V5035)


let__14181 := tmp15666
_ = let__14181

tmp15667 := MakeNative(func(__e *ControlFlow) {
tmp15668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14179)
}
__typedArg0 := let__14179
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15669 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15668, V5035)


let__14183 := tmp15669
_ = let__14183

tmp15670 := MakeNative(func(__e *ControlFlow) {
Z5211 := __e.Get(1)
_ = Z5211
__e.TailApply(let__14180, Z5211)
return
}, 1)

let__14184 := tmp15670
_ = let__14184

tmp15686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14183)
}
__typedArg0 := let__14183
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15686 {
tmp15671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14183)
}
__typedArg0 := let__14183
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14185 := tmp15671
_ = let__14185

tmp15672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14183)
}
__typedArg0 := let__14183
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15673 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15672, V5035)


let__14186 := tmp15673
_ = let__14186

tmp15674 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14184, let__14185)
return
}, 0)

let__14187 := tmp15674
_ = let__14187

tmp15678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14186, Nil)
}
__typedArg0 := let__14186
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15678 {
__e.TailApply(PrimFunc(symthaw), let__14187)
return
} else {
tmp15676 := Call(__e, PrimFunc(symshen_4pvar_2), let__14186)


if True == tmp15676 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14186, Nil, V5035, let__14187)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15684 := Call(__e, PrimFunc(symshen_4pvar_2), let__14183)


if True == tmp15684 {
tmp15679 := Call(__e, PrimFunc(symshen_4newpv), V5035)


let__14188 := tmp15679
_ = let__14188

tmp15680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14188, Nil)
}
__typedArg0 := let__14188
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15681 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14184, let__14188)
return
}, 0)

tmp15682 := Call(__e, PrimFunc(symshen_4bind_b), let__14183, tmp15680, V5035, tmp15681)


__e.TailApply(PrimFunc(symshen_4gc), V5035, tmp15682)
return


} else {
__e.Return(False)
return
}


}


}, 0)

let__14182 := tmp15667
_ = let__14182

tmp15692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14181, symstream)
}
__typedArg0 := let__14181
__typedArg1 := symstream
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15687 Obj

if True == tmp15692 {
tmp15688 := Call(__e, PrimFunc(symthaw), let__14182)


ifres15687 = tmp15688


} else {
tmp15691 := Call(__e, PrimFunc(symshen_4pvar_2), let__14181)


var ifres15689 Obj

if True == tmp15691 {
tmp15690 := Call(__e, PrimFunc(symshen_4bind_b), let__14181, symstream, V5035, let__14182)


ifres15689 = tmp15690


} else {
ifres15689 = False


}

ifres15687 = ifres15689


}

ifres15664 = ifres15687


} else {
tmp15700 := Call(__e, PrimFunc(symshen_4pvar_2), let__14179)


var ifres15693 Obj

if True == tmp15700 {
tmp15694 := Call(__e, PrimFunc(symshen_4newpv), V5035)


let__14189 := tmp15694
_ = let__14189

tmp15695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14189, Nil)
}
__typedArg0 := let__14189
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp15695)
}
__typedArg0 := symstream
__typedArg1 := tmp15695
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15697 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14180, let__14189)
return
}, 0)

tmp15698 := Call(__e, PrimFunc(symshen_4bind_b), let__14179, tmp15696, V5035, tmp15697)


tmp15699 := Call(__e, PrimFunc(symshen_4gc), V5035, tmp15698)


ifres15693 = tmp15699


} else {
ifres15693 = False


}

ifres15664 = ifres15693


}

ifres15654 = ifres15664


} else {
ifres15654 = False


}

ifres15650 = ifres15654


} else {
ifres15650 = False


}

ifres15646 = ifres15650


} else {
ifres15646 = False


}

ifres15643 = ifres15646


} else {
ifres15643 = False


}

ifres15640 = ifres15643


} else {
ifres15640 = False


}

ifres15638 = ifres15640


} else {
ifres15638 = False


}

let__14171 := ifres15638
_ = let__14171

tmp15825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14171, False)
}
__typedArg0 := let__14171
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15825 {
tmp15736 := Call(__e, PrimFunc(symshen_4unlocked_2), V5036)


var ifres15708 Obj

if True == tmp15736 {
tmp15709 := Call(__e, PrimFunc(symshen_4lazyderef), V5032, V5035)


let__14191 := tmp15709
_ = let__14191

tmp15735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14191)
}
__typedArg0 := let__14191
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15710 Obj

if True == tmp15735 {
tmp15711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14191)
}
__typedArg0 := let__14191
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15712 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15711, V5035)


let__14192 := tmp15712
_ = let__14192

tmp15734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14192, symtype)
}
__typedArg0 := let__14192
__typedArg1 := symtype
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15713 Obj

if True == tmp15734 {
tmp15714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14191)
}
__typedArg0 := let__14191
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15715 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15714, V5035)


let__14193 := tmp15715
_ = let__14193

tmp15733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14193)
}
__typedArg0 := let__14193
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15716 Obj

if True == tmp15733 {
tmp15717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14193)
}
__typedArg0 := let__14193
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14194 := tmp15717
_ = let__14194

tmp15718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14193)
}
__typedArg0 := let__14193
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15719 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15718, V5035)


let__14195 := tmp15719
_ = let__14195

tmp15732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14195)
}
__typedArg0 := let__14195
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15720 Obj

if True == tmp15732 {
tmp15721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14195)
}
__typedArg0 := let__14195
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14196 := tmp15721
_ = let__14196

tmp15722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14195)
}
__typedArg0 := let__14195
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15723 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15722, V5035)


let__14197 := tmp15723
_ = let__14197

tmp15731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14197, Nil)
}
__typedArg0 := let__14197
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15724 Obj

if True == tmp15731 {
tmp15725 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15725

tmp15726 := MakeNative(func(__e *ControlFlow) {
tmp15727 := Call(__e, PrimFunc(symshen_4deref), let__14196, V5035)


tmp15728 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp15727)


tmp15729 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14194, V5033, V5034, V5035, V5036, let__14026, V5038)
return
}, 0)

__e.TailApply(PrimFunc(symis_b), tmp15728, V5033, V5035, V5036, let__14026, tmp15729)
return


}, 0)

tmp15730 := Call(__e, PrimFunc(symshen_4cut), V5035, V5036, let__14026, tmp15726)


ifres15724 = tmp15730


} else {
ifres15724 = False


}

ifres15720 = ifres15724


} else {
ifres15720 = False


}

ifres15716 = ifres15720


} else {
ifres15716 = False


}

ifres15713 = ifres15716


} else {
ifres15713 = False


}

ifres15710 = ifres15713


} else {
ifres15710 = False


}

ifres15708 = ifres15710


} else {
ifres15708 = False


}

let__14190 := ifres15708
_ = let__14190

tmp15823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14190, False)
}
__typedArg0 := let__14190
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15823 {
tmp15767 := Call(__e, PrimFunc(symshen_4unlocked_2), V5036)


var ifres15737 Obj

if True == tmp15767 {
tmp15738 := Call(__e, PrimFunc(symshen_4lazyderef), V5032, V5035)


let__14199 := tmp15738
_ = let__14199

tmp15766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14199)
}
__typedArg0 := let__14199
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15739 Obj

if True == tmp15766 {
tmp15740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14199)
}
__typedArg0 := let__14199
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15741 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15740, V5035)


let__14200 := tmp15741
_ = let__14200

tmp15765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14200, symshen_4input_1h_7)
}
__typedArg0 := let__14200
__typedArg1 := symshen_4input_1h_7
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15742 Obj

if True == tmp15765 {
tmp15743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14199)
}
__typedArg0 := let__14199
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15744 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15743, V5035)


let__14201 := tmp15744
_ = let__14201

tmp15764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14201)
}
__typedArg0 := let__14201
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15745 Obj

if True == tmp15764 {
tmp15746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14201)
}
__typedArg0 := let__14201
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14202 := tmp15746
_ = let__14202

tmp15747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14201)
}
__typedArg0 := let__14201
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15748 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15747, V5035)


let__14203 := tmp15748
_ = let__14203

tmp15763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14203)
}
__typedArg0 := let__14203
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15749 Obj

if True == tmp15763 {
tmp15750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14203)
}
__typedArg0 := let__14203
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14204 := tmp15750
_ = let__14204

tmp15751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14203)
}
__typedArg0 := let__14203
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15752 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15751, V5035)


let__14205 := tmp15752
_ = let__14205

tmp15762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14205, Nil)
}
__typedArg0 := let__14205
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15753 Obj

if True == tmp15762 {
tmp15754 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15754

tmp15755 := Call(__e, PrimFunc(symshen_4deref), let__14202, V5035)


tmp15756 := Call(__e, PrimFunc(symshen_4rdecons), tmp15755)


tmp15757 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp15756)


tmp15758 := MakeNative(func(__e *ControlFlow) {
tmp15759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symin, Nil)
}
__typedArg0 := symin
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstream, tmp15759)
}
__typedArg0 := symstream
__typedArg1 := tmp15759
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14204, tmp15760, V5034, V5035, V5036, let__14026, V5038)
return


}, 0)

tmp15761 := Call(__e, PrimFunc(symis), V5033, tmp15757, V5035, V5036, let__14026, tmp15758)


ifres15753 = tmp15761


} else {
ifres15753 = False


}

ifres15749 = ifres15753


} else {
ifres15749 = False


}

ifres15745 = ifres15749


} else {
ifres15745 = False


}

ifres15742 = ifres15745


} else {
ifres15742 = False


}

ifres15739 = ifres15742


} else {
ifres15739 = False


}

ifres15737 = ifres15739


} else {
ifres15737 = False


}

let__14198 := ifres15737
_ = let__14198

tmp15821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14198, False)
}
__typedArg0 := let__14198
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15821 {
tmp15796 := Call(__e, PrimFunc(symshen_4unlocked_2), V5036)


var ifres15768 Obj

if True == tmp15796 {
tmp15769 := Call(__e, PrimFunc(symshen_4lazyderef), V5032, V5035)


let__14207 := tmp15769
_ = let__14207

tmp15795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14207)
}
__typedArg0 := let__14207
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15770 Obj

if True == tmp15795 {
tmp15771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14207)
}
__typedArg0 := let__14207
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15772 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15771, V5035)


let__14208 := tmp15772
_ = let__14208

tmp15794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14208, symset)
}
__typedArg0 := let__14208
__typedArg1 := symset
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15773 Obj

if True == tmp15794 {
tmp15774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14207)
}
__typedArg0 := let__14207
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15775 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15774, V5035)


let__14209 := tmp15775
_ = let__14209

tmp15793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14209)
}
__typedArg0 := let__14209
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15776 Obj

if True == tmp15793 {
tmp15777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14209)
}
__typedArg0 := let__14209
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14210 := tmp15777
_ = let__14210

tmp15778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14209)
}
__typedArg0 := let__14209
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15779 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15778, V5035)


let__14211 := tmp15779
_ = let__14211

tmp15792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14211)
}
__typedArg0 := let__14211
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15780 Obj

if True == tmp15792 {
tmp15781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14211)
}
__typedArg0 := let__14211
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14212 := tmp15781
_ = let__14212

tmp15782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14211)
}
__typedArg0 := let__14211
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15783 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15782, V5035)


let__14213 := tmp15783
_ = let__14213

tmp15791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14213, Nil)
}
__typedArg0 := let__14213
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15784 Obj

if True == tmp15791 {
tmp15785 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15785

tmp15786 := MakeNative(func(__e *ControlFlow) {
tmp15787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14210, Nil)
}
__typedArg0 := let__14210
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp15787)
}
__typedArg0 := symvalue
__typedArg1 := tmp15787
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15789 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14212, V5033, V5034, V5035, V5036, let__14026, V5038)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp15788, V5033, V5034, V5035, V5036, let__14026, tmp15789)
return


}, 0)

tmp15790 := Call(__e, PrimFunc(symshen_4system_1S_1h), let__14210, symsymbol, V5034, V5035, V5036, let__14026, tmp15786)


ifres15784 = tmp15790


} else {
ifres15784 = False


}

ifres15780 = ifres15784


} else {
ifres15780 = False


}

ifres15776 = ifres15780


} else {
ifres15776 = False


}

ifres15773 = ifres15776


} else {
ifres15773 = False


}

ifres15770 = ifres15773


} else {
ifres15770 = False


}

ifres15768 = ifres15770


} else {
ifres15768 = False


}

let__14206 := ifres15768
_ = let__14206

tmp15819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14206, False)
}
__typedArg0 := let__14206
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15819 {
tmp15804 := Call(__e, PrimFunc(symshen_4unlocked_2), V5036)


var ifres15797 Obj

if True == tmp15804 {
tmp15798 := Call(__e, PrimFunc(symshen_4newpv), V5035)


let__14215 := tmp15798
_ = let__14215

tmp15799 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15799

tmp15800 := MakeNative(func(__e *ControlFlow) {
tmp15801 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), V5032, V5033, let__14215, V5035, V5036, let__14026, V5038)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5035, V5036, let__14026, tmp15801)
return


}, 0)

tmp15802 := Call(__e, PrimFunc(symshen_4l_1rules), V5034, let__14215, False, V5035, V5036, let__14026, tmp15800)


tmp15803 := Call(__e, PrimFunc(symshen_4gc), V5035, tmp15802)


ifres15797 = tmp15803


} else {
ifres15797 = False


}

let__14214 := ifres15797
_ = let__14214

tmp15817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14214, False)
}
__typedArg0 := let__14214
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15817 {
tmp15813 := Call(__e, PrimFunc(symshen_4unlocked_2), V5036)


var ifres15805 Obj

if True == tmp15813 {
tmp15806 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15806

tmp15807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp15808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5033, Nil)
}
__typedArg0 := V5033
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15807, tmp15808)
}
__typedArg0 := tmp15807
__typedArg1 := tmp15808
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5032, tmp15809)
}
__typedArg0 := V5032
__typedArg1 := tmp15809
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp15812 := Call(__e, PrimFunc(symshen_4search_1user_1datatypes), tmp15810, V5034, tmp15811, V5035, V5036, let__14026, V5038)


ifres15805 = tmp15812


} else {
ifres15805 = False


}

let__14216 := ifres15805
_ = let__14216

tmp15815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14216, False)
}
__typedArg0 := let__14216
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15815 {
__e.TailApply(PrimFunc(symshen_4unlock), V5036, let__14026)
return
} else {
__e.Return(let__14216)
return
}


} else {
__e.Return(let__14214)
return
}


} else {
__e.Return(let__14206)
return
}


} else {
__e.Return(let__14198)
return
}


} else {
__e.Return(let__14190)
return
}


} else {
__e.Return(let__14171)
return
}


} else {
__e.Return(let__14158)
return
}


} else {
__e.Return(let__14132)
return
}


} else {
__e.Return(let__14122)
return
}


} else {
__e.Return(let__14103)
return
}


} else {
__e.Return(let__14079)
return
}


} else {
__e.Return(let__14060)
return
}


} else {
__e.Return(let__14053)
return
}


} else {
__e.Return(let__14046)
return
}


} else {
__e.Return(let__14040)
return
}


} else {
__e.Return(let__14034)
return
}


} else {
__e.Return(let__14030)
return
}


} else {
__e.Return(let__14029)
return
}


} else {
__e.Return(let__14028)
return
}


} else {
__e.Return(let__14027)
return
}


}, 7)

tmp15854 := Call(__e, ns2_1set, symshen_4system_1S_1h, tmp15103)


_ = tmp15854

tmp15855 := MakeNative(func(__e *ControlFlow) {
V5244 := __e.Get(1)
_ = V5244
tmp15889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5244)
}
__typedArg0 := V5244
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15870 Obj

if True == tmp15889 {
tmp15887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5244)
}
__typedArg0 := V5244
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp15887)
}
__typedArg0 := symcons
__typedArg1 := tmp15887
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15872 Obj

if True == tmp15888 {
tmp15885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5244)
}
__typedArg0 := V5244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15885)
}
__typedArg0 := tmp15885
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15874 Obj

if True == tmp15886 {
tmp15882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5244)
}
__typedArg0 := V5244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15882)
}
__typedArg0 := tmp15882
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp15883)
}
__typedArg0 := tmp15883
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres15876 Obj

if True == tmp15884 {
tmp15878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5244)
}
__typedArg0 := V5244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15878)
}
__typedArg0 := tmp15878
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp15879)
}
__typedArg0 := tmp15879
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp15880)
}
__typedArg0 := Nil
__typedArg1 := tmp15880
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15877 Obj

if True == tmp15881 {
ifres15877 = True


} else {
ifres15877 = False


}

ifres15876 = ifres15877


} else {
ifres15876 = False


}

var ifres15875 Obj

if True == ifres15876 {
ifres15875 = True


} else {
ifres15875 = False


}

ifres15874 = ifres15875


} else {
ifres15874 = False


}

var ifres15873 Obj

if True == ifres15874 {
ifres15873 = True


} else {
ifres15873 = False


}

ifres15872 = ifres15873


} else {
ifres15872 = False


}

var ifres15871 Obj

if True == ifres15872 {
ifres15871 = True


} else {
ifres15871 = False


}

ifres15870 = ifres15871


} else {
ifres15870 = False


}

if True == ifres15870 {
tmp15856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5244)
}
__typedArg0 := V5244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15856)
}
__typedArg0 := tmp15856
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15858 := Call(__e, PrimFunc(symshen_4rdecons), tmp15857)


tmp15859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5244)
}
__typedArg0 := V5244
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp15860)
}
__typedArg0 := tmp15860
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15862 := Call(__e, PrimFunc(symshen_4rdecons), tmp15861)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15858, tmp15862)
}
__typedArg0 := tmp15858
__typedArg1 := tmp15862
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp15868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5244)
}
__typedArg0 := V5244
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15868 {
tmp15863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5244)
}
__typedArg0 := V5244
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15864 := Call(__e, PrimFunc(symshen_4rdecons), tmp15863)


tmp15865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5244)
}
__typedArg0 := V5244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15866 := Call(__e, PrimFunc(symshen_4rdecons), tmp15865)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp15864, tmp15866)
}
__typedArg0 := tmp15864
__typedArg1 := tmp15866
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V5244)
return
}


}


}, 1)

tmp15890 := Call(__e, ns2_1set, symshen_4rdecons, tmp15855)


_ = tmp15890

tmp15891 := MakeNative(func(__e *ControlFlow) {
V5245 := __e.Get(1)
_ = V5245
V5246 := __e.Get(2)
_ = V5246
V5247 := __e.Get(3)
_ = V5247
V5248 := __e.Get(4)
_ = V5248
V5249 := __e.Get(5)
_ = V5249
V5250 := __e.Get(6)
_ = V5250
tmp15904 := Call(__e, PrimFunc(symshen_4unlocked_2), V5248)


var ifres15892 Obj

if True == tmp15904 {
tmp15893 := Call(__e, PrimFunc(symshen_4lazyderef), V5246, V5247)


let__14218 := tmp15893
_ = let__14218

tmp15894 := MakeNative(func(__e *ControlFlow) {
tmp15895 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15895

tmp15896 := Call(__e, PrimFunc(symshen_4lazyderef), V5245, V5247)


tmp15897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(tmp15896)
}
__typedArg0 := tmp15896
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

__e.TailApply(PrimFunc(symwhen), tmp15897, V5247, V5248, V5249, V5250)
return


}, 0)

let__14219 := tmp15894
_ = let__14219

tmp15903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14218, symnumber)
}
__typedArg0 := let__14218
__typedArg1 := symnumber
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15898 Obj

if True == tmp15903 {
tmp15899 := Call(__e, PrimFunc(symthaw), let__14219)


ifres15898 = tmp15899


} else {
tmp15902 := Call(__e, PrimFunc(symshen_4pvar_2), let__14218)


var ifres15900 Obj

if True == tmp15902 {
tmp15901 := Call(__e, PrimFunc(symshen_4bind_b), let__14218, symnumber, V5247, let__14219)


ifres15900 = tmp15901


} else {
ifres15900 = False


}

ifres15898 = ifres15900


}

ifres15892 = ifres15898


} else {
ifres15892 = False


}

let__14217 := ifres15892
_ = let__14217

tmp15994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14217, False)
}
__typedArg0 := let__14217
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15994 {
tmp15917 := Call(__e, PrimFunc(symshen_4unlocked_2), V5248)


var ifres15905 Obj

if True == tmp15917 {
tmp15906 := Call(__e, PrimFunc(symshen_4lazyderef), V5246, V5247)


let__14221 := tmp15906
_ = let__14221

tmp15907 := MakeNative(func(__e *ControlFlow) {
tmp15908 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15908

tmp15909 := Call(__e, PrimFunc(symshen_4lazyderef), V5245, V5247)


tmp15910 := Call(__e, PrimFunc(symboolean_2), tmp15909)


__e.TailApply(PrimFunc(symwhen), tmp15910, V5247, V5248, V5249, V5250)
return


}, 0)

let__14222 := tmp15907
_ = let__14222

tmp15916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14221, symboolean)
}
__typedArg0 := let__14221
__typedArg1 := symboolean
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15911 Obj

if True == tmp15916 {
tmp15912 := Call(__e, PrimFunc(symthaw), let__14222)


ifres15911 = tmp15912


} else {
tmp15915 := Call(__e, PrimFunc(symshen_4pvar_2), let__14221)


var ifres15913 Obj

if True == tmp15915 {
tmp15914 := Call(__e, PrimFunc(symshen_4bind_b), let__14221, symboolean, V5247, let__14222)


ifres15913 = tmp15914


} else {
ifres15913 = False


}

ifres15911 = ifres15913


}

ifres15905 = ifres15911


} else {
ifres15905 = False


}

let__14220 := ifres15905
_ = let__14220

tmp15992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14220, False)
}
__typedArg0 := let__14220
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15992 {
tmp15930 := Call(__e, PrimFunc(symshen_4unlocked_2), V5248)


var ifres15918 Obj

if True == tmp15930 {
tmp15919 := Call(__e, PrimFunc(symshen_4lazyderef), V5246, V5247)


let__14224 := tmp15919
_ = let__14224

tmp15920 := MakeNative(func(__e *ControlFlow) {
tmp15921 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15921

tmp15922 := Call(__e, PrimFunc(symshen_4lazyderef), V5245, V5247)


tmp15923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(tmp15922)
}
__typedArg0 := tmp15922
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

__e.TailApply(PrimFunc(symwhen), tmp15923, V5247, V5248, V5249, V5250)
return


}, 0)

let__14225 := tmp15920
_ = let__14225

tmp15929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14224, symstring)
}
__typedArg0 := let__14224
__typedArg1 := symstring
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15924 Obj

if True == tmp15929 {
tmp15925 := Call(__e, PrimFunc(symthaw), let__14225)


ifres15924 = tmp15925


} else {
tmp15928 := Call(__e, PrimFunc(symshen_4pvar_2), let__14224)


var ifres15926 Obj

if True == tmp15928 {
tmp15927 := Call(__e, PrimFunc(symshen_4bind_b), let__14224, symstring, V5247, let__14225)


ifres15926 = tmp15927


} else {
ifres15926 = False


}

ifres15924 = ifres15926


}

ifres15918 = ifres15924


} else {
ifres15918 = False


}

let__14223 := ifres15918
_ = let__14223

tmp15990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14223, False)
}
__typedArg0 := let__14223
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15990 {
tmp15943 := Call(__e, PrimFunc(symshen_4unlocked_2), V5248)


var ifres15931 Obj

if True == tmp15943 {
tmp15932 := Call(__e, PrimFunc(symshen_4lazyderef), V5246, V5247)


let__14227 := tmp15932
_ = let__14227

tmp15933 := MakeNative(func(__e *ControlFlow) {
tmp15934 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15934

tmp15935 := Call(__e, PrimFunc(symshen_4lazyderef), V5245, V5247)


tmp15936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(tmp15935)
}
__typedArg0 := tmp15935
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

__e.TailApply(PrimFunc(symwhen), tmp15936, V5247, V5248, V5249, V5250)
return


}, 0)

let__14228 := tmp15933
_ = let__14228

tmp15942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14227, symsymbol)
}
__typedArg0 := let__14227
__typedArg1 := symsymbol
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres15937 Obj

if True == tmp15942 {
tmp15938 := Call(__e, PrimFunc(symthaw), let__14228)


ifres15937 = tmp15938


} else {
tmp15941 := Call(__e, PrimFunc(symshen_4pvar_2), let__14227)


var ifres15939 Obj

if True == tmp15941 {
tmp15940 := Call(__e, PrimFunc(symshen_4bind_b), let__14227, symsymbol, V5247, let__14228)


ifres15939 = tmp15940


} else {
ifres15939 = False


}

ifres15937 = ifres15939


}

ifres15931 = ifres15937


} else {
ifres15931 = False


}

let__14226 := ifres15931
_ = let__14226

tmp15988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14226, False)
}
__typedArg0 := let__14226
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15988 {
tmp15986 := Call(__e, PrimFunc(symshen_4unlocked_2), V5248)


if True == tmp15986 {
tmp15944 := Call(__e, PrimFunc(symshen_4lazyderef), V5245, V5247)


let__14229 := tmp15944
_ = let__14229

tmp15984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14229, Nil)
}
__typedArg0 := let__14229
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15984 {
tmp15945 := Call(__e, PrimFunc(symshen_4lazyderef), V5246, V5247)


let__14230 := tmp15945
_ = let__14230

tmp15946 := MakeNative(func(__e *ControlFlow) {
Z5266 := __e.Get(1)
_ = Z5266
tmp15947 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15947

__e.TailApply(PrimFunc(symthaw), V5250)
return


}, 1)

let__14231 := tmp15946
_ = let__14231

tmp15982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14230)
}
__typedArg0 := let__14230
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15982 {
tmp15948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14230)
}
__typedArg0 := let__14230
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp15949 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15948, V5247)


let__14232 := tmp15949
_ = let__14232

tmp15950 := MakeNative(func(__e *ControlFlow) {
tmp15951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14230)
}
__typedArg0 := let__14230
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15952 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15951, V5247)


let__14234 := tmp15952
_ = let__14234

tmp15953 := MakeNative(func(__e *ControlFlow) {
Z5271 := __e.Get(1)
_ = Z5271
__e.TailApply(let__14231, Z5271)
return
}, 1)

let__14235 := tmp15953
_ = let__14235

tmp15969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14234)
}
__typedArg0 := let__14234
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp15969 {
tmp15954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14234)
}
__typedArg0 := let__14234
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14236 := tmp15954
_ = let__14236

tmp15955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14234)
}
__typedArg0 := let__14234
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp15956 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15955, V5247)


let__14237 := tmp15956
_ = let__14237

tmp15957 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14235, let__14236)
return
}, 0)

let__14238 := tmp15957
_ = let__14238

tmp15961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14237, Nil)
}
__typedArg0 := let__14237
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15961 {
__e.TailApply(PrimFunc(symthaw), let__14238)
return
} else {
tmp15959 := Call(__e, PrimFunc(symshen_4pvar_2), let__14237)


if True == tmp15959 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14237, Nil, V5247, let__14238)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15967 := Call(__e, PrimFunc(symshen_4pvar_2), let__14234)


if True == tmp15967 {
tmp15962 := Call(__e, PrimFunc(symshen_4newpv), V5247)


let__14239 := tmp15962
_ = let__14239

tmp15963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14239, Nil)
}
__typedArg0 := let__14239
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15964 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14235, let__14239)
return
}, 0)

tmp15965 := Call(__e, PrimFunc(symshen_4bind_b), let__14234, tmp15963, V5247, tmp15964)


__e.TailApply(PrimFunc(symshen_4gc), V5247, tmp15965)
return


} else {
__e.Return(False)
return
}


}


}, 0)

let__14233 := tmp15950
_ = let__14233

tmp15973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14232, symlist)
}
__typedArg0 := let__14232
__typedArg1 := symlist
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp15973 {
__e.TailApply(PrimFunc(symthaw), let__14233)
return
} else {
tmp15971 := Call(__e, PrimFunc(symshen_4pvar_2), let__14232)


if True == tmp15971 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14232, symlist, V5247, let__14233)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp15980 := Call(__e, PrimFunc(symshen_4pvar_2), let__14230)


if True == tmp15980 {
tmp15974 := Call(__e, PrimFunc(symshen_4newpv), V5247)


let__14240 := tmp15974
_ = let__14240

tmp15975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14240, Nil)
}
__typedArg0 := let__14240
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp15975)
}
__typedArg0 := symlist
__typedArg1 := tmp15975
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp15977 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14231, let__14240)
return
}, 0)

tmp15978 := Call(__e, PrimFunc(symshen_4bind_b), let__14230, tmp15976, V5247, tmp15977)


__e.TailApply(PrimFunc(symshen_4gc), V5247, tmp15978)
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
__e.Return(let__14226)
return
}


} else {
__e.Return(let__14223)
return
}


} else {
__e.Return(let__14220)
return
}


} else {
__e.Return(let__14217)
return
}


}, 6)

tmp15995 := Call(__e, ns2_1set, symshen_4primitive, tmp15891)


_ = tmp15995

tmp15996 := MakeNative(func(__e *ControlFlow) {
V5277 := __e.Get(1)
_ = V5277
V5278 := __e.Get(2)
_ = V5278
V5279 := __e.Get(3)
_ = V5279
V5280 := __e.Get(4)
_ = V5280
V5281 := __e.Get(5)
_ = V5281
V5282 := __e.Get(6)
_ = V5282
V5283 := __e.Get(7)
_ = V5283
tmp15999 := Call(__e, PrimFunc(symshen_4lazyderef), V5279, V5280)


tmp16000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp15999, Nil)
}
__typedArg0 := tmp15999
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16000 {
__e.Return(False)
return
} else {
tmp15997 := Call(__e, PrimFunc(symshen_4deref), V5277, V5280)


__e.TailApply(PrimFunc(symshen_4by_1hypothesis_1h), tmp15997, V5278, V5279, V5280, V5281, V5282, V5283)
return


}


}, 7)

tmp16001 := Call(__e, ns2_1set, symshen_4by_1hypothesis, tmp15996)


_ = tmp16001

tmp16002 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
__self6 := __e.Get(6)
__self7 := __e.Get(7)
__selftop:
V5284 := __self1
_ = V5284
V5285 := __self2
_ = V5285
V5286 := __self3
_ = V5286
V5287 := __self4
_ = V5287
V5288 := __self5
_ = V5288
V5289 := __self6
_ = V5289
V5290 := __self7
_ = V5290
tmp16035 := Call(__e, PrimFunc(symshen_4unlocked_2), V5288)


var ifres16003 Obj

if True == tmp16035 {
tmp16004 := Call(__e, PrimFunc(symshen_4lazyderef), V5286, V5287)


let__14242 := tmp16004
_ = let__14242

tmp16034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14242)
}
__typedArg0 := let__14242
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16005 Obj

if True == tmp16034 {
tmp16006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14242)
}
__typedArg0 := let__14242
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16007 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16006, V5287)


let__14243 := tmp16007
_ = let__14243

tmp16033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14243)
}
__typedArg0 := let__14243
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16008 Obj

if True == tmp16033 {
tmp16009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14243)
}
__typedArg0 := let__14243
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14244 := tmp16009
_ = let__14244

tmp16010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14243)
}
__typedArg0 := let__14243
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16011 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16010, V5287)


let__14245 := tmp16011
_ = let__14245

tmp16032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14245)
}
__typedArg0 := let__14245
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16012 Obj

if True == tmp16032 {
tmp16013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14245)
}
__typedArg0 := let__14245
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14246 := tmp16013
_ = let__14246

tmp16014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14245)
}
__typedArg0 := let__14245
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16015 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16014, V5287)


let__14247 := tmp16015
_ = let__14247

tmp16031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14247)
}
__typedArg0 := let__14247
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16016 Obj

if True == tmp16031 {
tmp16017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14247)
}
__typedArg0 := let__14247
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14248 := tmp16017
_ = let__14248

tmp16018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14247)
}
__typedArg0 := let__14247
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16019 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16018, V5287)


let__14249 := tmp16019
_ = let__14249

tmp16030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14249, Nil)
}
__typedArg0 := let__14249
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16020 Obj

if True == tmp16030 {
tmp16021 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16021

tmp16022 := Call(__e, PrimFunc(symshen_4deref), let__14246, V5287)


tmp16023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16022, tmp16023)
}
__typedArg0 := tmp16022
__typedArg1 := tmp16023
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16025 := MakeNative(func(__e *ControlFlow) {
tmp16026 := Call(__e, PrimFunc(symshen_4deref), let__14244, V5287)


tmp16027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V5284, tmp16026)
}
__typedArg0 := V5284
__typedArg1 := tmp16026
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16028 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symis_b), V5285, let__14248, V5287, V5288, V5289, V5290)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp16027, V5287, V5288, V5289, tmp16028)
return


}, 0)

tmp16029 := Call(__e, PrimFunc(symwhen), tmp16024, V5287, V5288, V5289, tmp16025)


ifres16020 = tmp16029


} else {
ifres16020 = False


}

ifres16016 = ifres16020


} else {
ifres16016 = False


}

ifres16012 = ifres16016


} else {
ifres16012 = False


}

ifres16008 = ifres16012


} else {
ifres16008 = False


}

ifres16005 = ifres16008


} else {
ifres16005 = False


}

ifres16003 = ifres16005


} else {
ifres16003 = False


}

let__14241 := ifres16003
_ = let__14241

tmp16044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14241, False)
}
__typedArg0 := let__14241
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16044 {
tmp16042 := Call(__e, PrimFunc(symshen_4unlocked_2), V5288)


if True == tmp16042 {
tmp16036 := Call(__e, PrimFunc(symshen_4lazyderef), V5286, V5287)


let__14250 := tmp16036
_ = let__14250

tmp16040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14250)
}
__typedArg0 := let__14250
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16040 {
tmp16037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14250)
}
__typedArg0 := let__14250
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14251 := tmp16037
_ = let__14251

tmp16038 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16038

if PrimFunc(symshen_4by_1hypothesis_1h) == __self {
__self1, __self2, __self3, __self4, __self5, __self6, __self7 = V5284, V5285, let__14251, V5287, V5288, V5289, V5290
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4by_1hypothesis_1h), V5284, V5285, let__14251, V5287, V5288, V5289, V5290)
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
__e.Return(let__14241)
return
}


}, 7)

tmp16045 := Call(__e, ns2_1set, symshen_4by_1hypothesis_1h, tmp16002)


_ = tmp16045

tmp16046 := MakeNative(func(__e *ControlFlow) {
V5302 := __e.Get(1)
_ = V5302
V5303 := __e.Get(2)
_ = V5303
V5304 := __e.Get(3)
_ = V5304
V5305 := __e.Get(4)
_ = V5305
V5306 := __e.Get(5)
_ = V5306
V5307 := __e.Get(6)
_ = V5307
tmp16051 := Call(__e, PrimFunc(symshen_4unlocked_2), V5305)


if True == tmp16051 {
tmp16047 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16047

tmp16048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp16049 := Call(__e, PrimFunc(symassoc), V5302, tmp16048)


__e.TailApply(PrimFunc(symshen_4sigf), tmp16049, V5303, V5304, V5305, V5306, V5307)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp16052 := Call(__e, ns2_1set, symshen_4lookupsig, tmp16046)


_ = tmp16052

tmp16053 := MakeNative(func(__e *ControlFlow) {
V5322 := __e.Get(1)
_ = V5322
V5323 := __e.Get(2)
_ = V5323
V5324 := __e.Get(3)
_ = V5324
V5325 := __e.Get(4)
_ = V5325
V5326 := __e.Get(5)
_ = V5326
V5327 := __e.Get(6)
_ = V5327
tmp16060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5322)
}
__typedArg0 := V5322
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16060 {
tmp16054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5322)
}
__typedArg0 := V5322
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16055 := Call(__e, tmp16054, V5323)


tmp16056 := Call(__e, tmp16055, V5324)


tmp16057 := Call(__e, tmp16056, V5325)


tmp16058 := Call(__e, tmp16057, V5326)


__e.TailApply(tmp16058, V5327)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp16061 := Call(__e, ns2_1set, symshen_4sigf, tmp16053)


_ = tmp16061

tmp16062 := MakeNative(func(__e *ControlFlow) {
V5328 := __e.Get(1)
_ = V5328
tmp16063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(MakeInteger(3))
}
__typedArg0 := MakeInteger(3)
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

let__14252 := tmp16063
_ = let__14252

tmp16064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__14252, MakeInteger(0), symshen_4print_1freshterm)
}
__typedArg0 := let__14252
__typedArg1 := MakeInteger(0)
__typedArg2 := symshen_4print_1freshterm
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__14253 := tmp16064
_ = let__14253

tmp16065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__14253, MakeInteger(1), V5328)
}
__typedArg0 := let__14253
__typedArg1 := MakeInteger(1)
__typedArg2 := V5328
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__14254 := tmp16065
_ = let__14254

tmp16066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dgensym_d)
}
__typedArg0 := symshen_4_dgensym_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp16068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dgensym_d, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp16066)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp16066
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := symshen_4_dgensym_d
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp16066)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp16066
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

tmp16069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__14254, MakeInteger(2), tmp16068)
}
__typedArg0 := let__14254
__typedArg1 := MakeInteger(2)
__typedArg2 := tmp16068
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__14255 := tmp16069
_ = let__14255

__e.Return(let__14255)
return


}, 1)

tmp16070 := Call(__e, ns2_1set, symshen_4freshterm, tmp16062)


_ = tmp16070

tmp16071 := MakeNative(func(__e *ControlFlow) {
V5333 := __e.Get(1)
_ = V5333
tmp16072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V5333, MakeInteger(1))
}
__typedArg0 := V5333
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp16073 := Call(__e, PrimFunc(symshen_4app), tmp16072, MakeString(""), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("&&"))
__typedS1, __typedOK1 := TypedString(tmp16073)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("&&")
__typedArg1 := tmp16073
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp16074 := Call(__e, ns2_1set, symshen_4print_1freshterm, tmp16071)


_ = tmp16074

tmp16075 := MakeNative(func(__e *ControlFlow) {
V5334 := __e.Get(1)
_ = V5334
V5335 := __e.Get(2)
_ = V5335
V5336 := __e.Get(3)
_ = V5336
V5337 := __e.Get(4)
_ = V5337
V5338 := __e.Get(5)
_ = V5338
V5339 := __e.Get(6)
_ = V5339
V5340 := __e.Get(7)
_ = V5340
tmp16079 := Call(__e, PrimFunc(symshen_4lazyderef), V5336, V5337)


tmp16080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16079, Nil)
}
__typedArg0 := tmp16079
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16080 {
__e.Return(False)
return
} else {
tmp16076 := Call(__e, PrimFunc(symshen_4deref), V5334, V5337)


tmp16077 := Call(__e, PrimFunc(symshen_4deref), V5335, V5337)


__e.TailApply(PrimFunc(symshen_4search_1user_1datatypes_1h), tmp16076, tmp16077, V5336, V5337, V5338, V5339, V5340)
return


}


}, 7)

tmp16081 := Call(__e, ns2_1set, symshen_4search_1user_1datatypes, tmp16075)


_ = tmp16081

tmp16082 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
__self6 := __e.Get(6)
__self7 := __e.Get(7)
__selftop:
V5341 := __self1
_ = V5341
V5342 := __self2
_ = V5342
V5343 := __self3
_ = V5343
V5344 := __self4
_ = V5344
V5345 := __self5
_ = V5345
V5346 := __self6
_ = V5346
V5347 := __self7
_ = V5347
tmp16097 := Call(__e, PrimFunc(symshen_4unlocked_2), V5345)


var ifres16083 Obj

if True == tmp16097 {
tmp16084 := Call(__e, PrimFunc(symshen_4lazyderef), V5343, V5344)


let__14257 := tmp16084
_ = let__14257

tmp16096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14257)
}
__typedArg0 := let__14257
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16085 Obj

if True == tmp16096 {
tmp16086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14257)
}
__typedArg0 := let__14257
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16087 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16086, V5344)


let__14258 := tmp16087
_ = let__14258

tmp16095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14258)
}
__typedArg0 := let__14258
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16088 Obj

if True == tmp16095 {
tmp16089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14258)
}
__typedArg0 := let__14258
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14259 := tmp16089
_ = let__14259

tmp16090 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16090

tmp16091 := Call(__e, PrimFunc(symshen_4deref), let__14259, V5344)


tmp16092 := Call(__e, tmp16091, V5341)


tmp16093 := Call(__e, tmp16092, V5342)


tmp16094 := Call(__e, PrimFunc(symcall), tmp16093, V5344, V5345, V5346, V5347)


ifres16088 = tmp16094


} else {
ifres16088 = False


}

ifres16085 = ifres16088


} else {
ifres16085 = False


}

ifres16083 = ifres16085


} else {
ifres16083 = False


}

let__14256 := ifres16083
_ = let__14256

tmp16106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14256, False)
}
__typedArg0 := let__14256
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16106 {
tmp16104 := Call(__e, PrimFunc(symshen_4unlocked_2), V5345)


if True == tmp16104 {
tmp16098 := Call(__e, PrimFunc(symshen_4lazyderef), V5343, V5344)


let__14260 := tmp16098
_ = let__14260

tmp16102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14260)
}
__typedArg0 := let__14260
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16102 {
tmp16099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14260)
}
__typedArg0 := let__14260
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14261 := tmp16099
_ = let__14261

tmp16100 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16100

if PrimFunc(symshen_4search_1user_1datatypes_1h) == __self {
__self1, __self2, __self3, __self4, __self5, __self6, __self7 = V5341, V5342, let__14261, V5344, V5345, V5346, V5347
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4search_1user_1datatypes_1h), V5341, V5342, let__14261, V5344, V5345, V5346, V5347)
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
__e.Return(let__14256)
return
}


}, 7)

tmp16107 := Call(__e, ns2_1set, symshen_4search_1user_1datatypes_1h, tmp16082)


_ = tmp16107

tmp16108 := MakeNative(func(__e *ControlFlow) {
V5354 := __e.Get(1)
_ = V5354
V5355 := __e.Get(2)
_ = V5355
V5356 := __e.Get(3)
_ = V5356
V5357 := __e.Get(4)
_ = V5357
V5358 := __e.Get(5)
_ = V5358
V5359 := __e.Get(6)
_ = V5359
V5360 := __e.Get(7)
_ = V5360
let__14262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5359)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5359
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__14262

tmp16120 := Call(__e, PrimFunc(symshen_4unlocked_2), V5358)


var ifres16110 Obj

if True == tmp16120 {
tmp16111 := Call(__e, PrimFunc(symshen_4lazyderef), V5354, V5357)


let__14264 := tmp16111
_ = let__14264

tmp16119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14264, Nil)
}
__typedArg0 := let__14264
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16112 Obj

if True == tmp16119 {
tmp16113 := Call(__e, PrimFunc(symshen_4lazyderef), V5356, V5357)


let__14265 := tmp16113
_ = let__14265

tmp16118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14265, True)
}
__typedArg0 := let__14265
__typedArg1 := True
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16114 Obj

if True == tmp16118 {
tmp16115 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16115

tmp16116 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symbind), V5355, Nil, V5357, V5358, let__14262, V5360)
return
}, 0)

tmp16117 := Call(__e, PrimFunc(symshen_4cut), V5357, V5358, let__14262, tmp16116)


ifres16114 = tmp16117


} else {
ifres16114 = False


}

ifres16112 = ifres16114


} else {
ifres16112 = False


}

ifres16110 = ifres16112


} else {
ifres16110 = False


}

let__14263 := ifres16110
_ = let__14263

tmp16463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14263, False)
}
__typedArg0 := let__14263
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16463 {
tmp16199 := Call(__e, PrimFunc(symshen_4unlocked_2), V5358)


var ifres16121 Obj

if True == tmp16199 {
tmp16122 := Call(__e, PrimFunc(symshen_4lazyderef), V5354, V5357)


let__14267 := tmp16122
_ = let__14267

tmp16198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14267)
}
__typedArg0 := let__14267
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16123 Obj

if True == tmp16198 {
tmp16124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14267)
}
__typedArg0 := let__14267
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16125 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16124, V5357)


let__14268 := tmp16125
_ = let__14268

tmp16197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14268)
}
__typedArg0 := let__14268
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16126 Obj

if True == tmp16197 {
tmp16127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14268)
}
__typedArg0 := let__14268
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16128 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16127, V5357)


let__14269 := tmp16128
_ = let__14269

tmp16196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14269)
}
__typedArg0 := let__14269
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16129 Obj

if True == tmp16196 {
tmp16130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14269)
}
__typedArg0 := let__14269
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16131 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16130, V5357)


let__14270 := tmp16131
_ = let__14270

tmp16195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14270, symcons)
}
__typedArg0 := let__14270
__typedArg1 := symcons
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16132 Obj

if True == tmp16195 {
tmp16133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14269)
}
__typedArg0 := let__14269
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16134 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16133, V5357)


let__14271 := tmp16134
_ = let__14271

tmp16194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14271)
}
__typedArg0 := let__14271
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16135 Obj

if True == tmp16194 {
tmp16136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14271)
}
__typedArg0 := let__14271
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14272 := tmp16136
_ = let__14272

tmp16137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14271)
}
__typedArg0 := let__14271
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16138 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16137, V5357)


let__14273 := tmp16138
_ = let__14273

tmp16193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14273)
}
__typedArg0 := let__14273
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16139 Obj

if True == tmp16193 {
tmp16140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14273)
}
__typedArg0 := let__14273
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14274 := tmp16140
_ = let__14274

tmp16141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14273)
}
__typedArg0 := let__14273
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16142 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16141, V5357)


let__14275 := tmp16142
_ = let__14275

tmp16192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14275, Nil)
}
__typedArg0 := let__14275
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16143 Obj

if True == tmp16192 {
tmp16144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14268)
}
__typedArg0 := let__14268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16145 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16144, V5357)


let__14276 := tmp16145
_ = let__14276

tmp16191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14276)
}
__typedArg0 := let__14276
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16146 Obj

if True == tmp16191 {
tmp16147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14276)
}
__typedArg0 := let__14276
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14277 := tmp16147
_ = let__14277

tmp16148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14276)
}
__typedArg0 := let__14276
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16149 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16148, V5357)


let__14278 := tmp16149
_ = let__14278

tmp16190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14278)
}
__typedArg0 := let__14278
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16150 Obj

if True == tmp16190 {
tmp16151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14278)
}
__typedArg0 := let__14278
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16152 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16151, V5357)


let__14279 := tmp16152
_ = let__14279

tmp16189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14279)
}
__typedArg0 := let__14279
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16153 Obj

if True == tmp16189 {
tmp16154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14279)
}
__typedArg0 := let__14279
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16155 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16154, V5357)


let__14280 := tmp16155
_ = let__14280

tmp16188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14280, symlist)
}
__typedArg0 := let__14280
__typedArg1 := symlist
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16156 Obj

if True == tmp16188 {
tmp16157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14279)
}
__typedArg0 := let__14279
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16158 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16157, V5357)


let__14281 := tmp16158
_ = let__14281

tmp16187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14281)
}
__typedArg0 := let__14281
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16159 Obj

if True == tmp16187 {
tmp16160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14281)
}
__typedArg0 := let__14281
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14282 := tmp16160
_ = let__14282

tmp16161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14281)
}
__typedArg0 := let__14281
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16162 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16161, V5357)


let__14283 := tmp16162
_ = let__14283

tmp16186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14283, Nil)
}
__typedArg0 := let__14283
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16163 Obj

if True == tmp16186 {
tmp16164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14278)
}
__typedArg0 := let__14278
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16165 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16164, V5357)


let__14284 := tmp16165
_ = let__14284

tmp16185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14284, Nil)
}
__typedArg0 := let__14284
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16166 Obj

if True == tmp16185 {
tmp16167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14267)
}
__typedArg0 := let__14267
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14285 := tmp16167
_ = let__14285

tmp16168 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16168

tmp16169 := Call(__e, PrimFunc(symshen_4deref), let__14277, V5357)


tmp16170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16169, tmp16170)
}
__typedArg0 := tmp16169
__typedArg1 := tmp16170
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16172 := MakeNative(func(__e *ControlFlow) {
tmp16173 := MakeNative(func(__e *ControlFlow) {
tmp16174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14282, Nil)
}
__typedArg0 := let__14282
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14277, tmp16174)
}
__typedArg0 := let__14277
__typedArg1 := tmp16174
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14272, tmp16175)
}
__typedArg0 := let__14272
__typedArg1 := tmp16175
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14282, Nil)
}
__typedArg0 := let__14282
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlist, tmp16177)
}
__typedArg0 := symlist
__typedArg1 := tmp16177
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16178, Nil)
}
__typedArg0 := tmp16178
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14277, tmp16179)
}
__typedArg0 := let__14277
__typedArg1 := tmp16179
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14274, tmp16180)
}
__typedArg0 := let__14274
__typedArg1 := tmp16180
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16181, let__14285)
}
__typedArg0 := tmp16181
__typedArg1 := let__14285
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16176, tmp16182)
}
__typedArg0 := tmp16176
__typedArg1 := tmp16182
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16183, V5355, True, V5357, V5358, let__14262, V5360)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5357, V5358, let__14262, tmp16173)
return


}, 0)

tmp16184 := Call(__e, PrimFunc(symwhen), tmp16171, V5357, V5358, let__14262, tmp16172)


ifres16166 = tmp16184


} else {
ifres16166 = False


}

ifres16163 = ifres16166


} else {
ifres16163 = False


}

ifres16159 = ifres16163


} else {
ifres16159 = False


}

ifres16156 = ifres16159


} else {
ifres16156 = False


}

ifres16153 = ifres16156


} else {
ifres16153 = False


}

ifres16150 = ifres16153


} else {
ifres16150 = False


}

ifres16146 = ifres16150


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

ifres16135 = ifres16139


} else {
ifres16135 = False


}

ifres16132 = ifres16135


} else {
ifres16132 = False


}

ifres16129 = ifres16132


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

ifres16121 = ifres16123


} else {
ifres16121 = False


}

let__14266 := ifres16121
_ = let__14266

tmp16461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14266, False)
}
__typedArg0 := let__14266
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16461 {
tmp16281 := Call(__e, PrimFunc(symshen_4unlocked_2), V5358)


var ifres16200 Obj

if True == tmp16281 {
tmp16201 := Call(__e, PrimFunc(symshen_4lazyderef), V5354, V5357)


let__14287 := tmp16201
_ = let__14287

tmp16280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14287)
}
__typedArg0 := let__14287
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16202 Obj

if True == tmp16280 {
tmp16203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14287)
}
__typedArg0 := let__14287
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16204 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16203, V5357)


let__14288 := tmp16204
_ = let__14288

tmp16279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14288)
}
__typedArg0 := let__14288
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16205 Obj

if True == tmp16279 {
tmp16206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14288)
}
__typedArg0 := let__14288
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16207 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16206, V5357)


let__14289 := tmp16207
_ = let__14289

tmp16278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14289)
}
__typedArg0 := let__14289
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16208 Obj

if True == tmp16278 {
tmp16209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14289)
}
__typedArg0 := let__14289
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16210 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16209, V5357)


let__14290 := tmp16210
_ = let__14290

tmp16277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14290, sym_8p)
}
__typedArg0 := let__14290
__typedArg1 := sym_8p
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16211 Obj

if True == tmp16277 {
tmp16212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14289)
}
__typedArg0 := let__14289
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16213 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16212, V5357)


let__14291 := tmp16213
_ = let__14291

tmp16276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14291)
}
__typedArg0 := let__14291
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16214 Obj

if True == tmp16276 {
tmp16215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14291)
}
__typedArg0 := let__14291
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14292 := tmp16215
_ = let__14292

tmp16216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14291)
}
__typedArg0 := let__14291
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16217 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16216, V5357)


let__14293 := tmp16217
_ = let__14293

tmp16275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14293)
}
__typedArg0 := let__14293
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16218 Obj

if True == tmp16275 {
tmp16219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14293)
}
__typedArg0 := let__14293
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14294 := tmp16219
_ = let__14294

tmp16220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14293)
}
__typedArg0 := let__14293
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16221 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16220, V5357)


let__14295 := tmp16221
_ = let__14295

tmp16274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14295, Nil)
}
__typedArg0 := let__14295
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16222 Obj

if True == tmp16274 {
tmp16223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14288)
}
__typedArg0 := let__14288
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16224 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16223, V5357)


let__14296 := tmp16224
_ = let__14296

tmp16273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14296)
}
__typedArg0 := let__14296
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16225 Obj

if True == tmp16273 {
tmp16226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14296)
}
__typedArg0 := let__14296
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14297 := tmp16226
_ = let__14297

tmp16227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14296)
}
__typedArg0 := let__14296
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16228 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16227, V5357)


let__14298 := tmp16228
_ = let__14298

tmp16272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14298)
}
__typedArg0 := let__14298
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16229 Obj

if True == tmp16272 {
tmp16230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14298)
}
__typedArg0 := let__14298
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16231 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16230, V5357)


let__14299 := tmp16231
_ = let__14299

tmp16271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14299)
}
__typedArg0 := let__14299
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16232 Obj

if True == tmp16271 {
tmp16233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14299)
}
__typedArg0 := let__14299
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14300 := tmp16233
_ = let__14300

tmp16234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14299)
}
__typedArg0 := let__14299
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16235 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16234, V5357)


let__14301 := tmp16235
_ = let__14301

tmp16270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14301)
}
__typedArg0 := let__14301
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16236 Obj

if True == tmp16270 {
tmp16237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14301)
}
__typedArg0 := let__14301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16238 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16237, V5357)


let__14302 := tmp16238
_ = let__14302

tmp16269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14302, sym_d)
}
__typedArg0 := let__14302
__typedArg1 := sym_d
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16239 Obj

if True == tmp16269 {
tmp16240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14301)
}
__typedArg0 := let__14301
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16241 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16240, V5357)


let__14303 := tmp16241
_ = let__14303

tmp16268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14303)
}
__typedArg0 := let__14303
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16242 Obj

if True == tmp16268 {
tmp16243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14303)
}
__typedArg0 := let__14303
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14304 := tmp16243
_ = let__14304

tmp16244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14303)
}
__typedArg0 := let__14303
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16245 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16244, V5357)


let__14305 := tmp16245
_ = let__14305

tmp16267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14305, Nil)
}
__typedArg0 := let__14305
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16246 Obj

if True == tmp16267 {
tmp16247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14298)
}
__typedArg0 := let__14298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16248 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16247, V5357)


let__14306 := tmp16248
_ = let__14306

tmp16266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14306, Nil)
}
__typedArg0 := let__14306
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16249 Obj

if True == tmp16266 {
tmp16250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14287)
}
__typedArg0 := let__14287
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14307 := tmp16250
_ = let__14307

tmp16251 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16251

tmp16252 := Call(__e, PrimFunc(symshen_4deref), let__14297, V5357)


tmp16253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16252, tmp16253)
}
__typedArg0 := tmp16252
__typedArg1 := tmp16253
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16255 := MakeNative(func(__e *ControlFlow) {
tmp16256 := MakeNative(func(__e *ControlFlow) {
tmp16257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14300, Nil)
}
__typedArg0 := let__14300
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14297, tmp16257)
}
__typedArg0 := let__14297
__typedArg1 := tmp16257
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14292, tmp16258)
}
__typedArg0 := let__14292
__typedArg1 := tmp16258
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14304, Nil)
}
__typedArg0 := let__14304
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14297, tmp16260)
}
__typedArg0 := let__14297
__typedArg1 := tmp16260
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14294, tmp16261)
}
__typedArg0 := let__14294
__typedArg1 := tmp16261
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16262, let__14307)
}
__typedArg0 := tmp16262
__typedArg1 := let__14307
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16259, tmp16263)
}
__typedArg0 := tmp16259
__typedArg1 := tmp16263
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16264, V5355, True, V5357, V5358, let__14262, V5360)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5357, V5358, let__14262, tmp16256)
return


}, 0)

tmp16265 := Call(__e, PrimFunc(symwhen), tmp16254, V5357, V5358, let__14262, tmp16255)


ifres16249 = tmp16265


} else {
ifres16249 = False


}

ifres16246 = ifres16249


} else {
ifres16246 = False


}

ifres16242 = ifres16246


} else {
ifres16242 = False


}

ifres16239 = ifres16242


} else {
ifres16239 = False


}

ifres16236 = ifres16239


} else {
ifres16236 = False


}

ifres16232 = ifres16236


} else {
ifres16232 = False


}

ifres16229 = ifres16232


} else {
ifres16229 = False


}

ifres16225 = ifres16229


} else {
ifres16225 = False


}

ifres16222 = ifres16225


} else {
ifres16222 = False


}

ifres16218 = ifres16222


} else {
ifres16218 = False


}

ifres16214 = ifres16218


} else {
ifres16214 = False


}

ifres16211 = ifres16214


} else {
ifres16211 = False


}

ifres16208 = ifres16211


} else {
ifres16208 = False


}

ifres16205 = ifres16208


} else {
ifres16205 = False


}

ifres16202 = ifres16205


} else {
ifres16202 = False


}

ifres16200 = ifres16202


} else {
ifres16200 = False


}

let__14286 := ifres16200
_ = let__14286

tmp16459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14286, False)
}
__typedArg0 := let__14286
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16459 {
tmp16345 := Call(__e, PrimFunc(symshen_4unlocked_2), V5358)


var ifres16282 Obj

if True == tmp16345 {
tmp16283 := Call(__e, PrimFunc(symshen_4lazyderef), V5354, V5357)


let__14309 := tmp16283
_ = let__14309

tmp16344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14309)
}
__typedArg0 := let__14309
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16284 Obj

if True == tmp16344 {
tmp16285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14309)
}
__typedArg0 := let__14309
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16286 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16285, V5357)


let__14310 := tmp16286
_ = let__14310

tmp16343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14310)
}
__typedArg0 := let__14310
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16287 Obj

if True == tmp16343 {
tmp16288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14310)
}
__typedArg0 := let__14310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16289 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16288, V5357)


let__14311 := tmp16289
_ = let__14311

tmp16342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14311)
}
__typedArg0 := let__14311
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16290 Obj

if True == tmp16342 {
tmp16291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14311)
}
__typedArg0 := let__14311
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16292 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16291, V5357)


let__14312 := tmp16292
_ = let__14312

tmp16341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14312, sym_8s)
}
__typedArg0 := let__14312
__typedArg1 := sym_8s
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16293 Obj

if True == tmp16341 {
tmp16294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14311)
}
__typedArg0 := let__14311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16295 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16294, V5357)


let__14313 := tmp16295
_ = let__14313

tmp16340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14313)
}
__typedArg0 := let__14313
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16296 Obj

if True == tmp16340 {
tmp16297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14313)
}
__typedArg0 := let__14313
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14314 := tmp16297
_ = let__14314

tmp16298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14313)
}
__typedArg0 := let__14313
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16299 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16298, V5357)


let__14315 := tmp16299
_ = let__14315

tmp16339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14315)
}
__typedArg0 := let__14315
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16300 Obj

if True == tmp16339 {
tmp16301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14315)
}
__typedArg0 := let__14315
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14316 := tmp16301
_ = let__14316

tmp16302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14315)
}
__typedArg0 := let__14315
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16303 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16302, V5357)


let__14317 := tmp16303
_ = let__14317

tmp16338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14317, Nil)
}
__typedArg0 := let__14317
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16304 Obj

if True == tmp16338 {
tmp16305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14310)
}
__typedArg0 := let__14310
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16306 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16305, V5357)


let__14318 := tmp16306
_ = let__14318

tmp16337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14318)
}
__typedArg0 := let__14318
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16307 Obj

if True == tmp16337 {
tmp16308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14318)
}
__typedArg0 := let__14318
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14319 := tmp16308
_ = let__14319

tmp16309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14318)
}
__typedArg0 := let__14318
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16310 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16309, V5357)


let__14320 := tmp16310
_ = let__14320

tmp16336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14320)
}
__typedArg0 := let__14320
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16311 Obj

if True == tmp16336 {
tmp16312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14320)
}
__typedArg0 := let__14320
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16313 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16312, V5357)


let__14321 := tmp16313
_ = let__14321

tmp16335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14321, symstring)
}
__typedArg0 := let__14321
__typedArg1 := symstring
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16314 Obj

if True == tmp16335 {
tmp16315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14320)
}
__typedArg0 := let__14320
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16316 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16315, V5357)


let__14322 := tmp16316
_ = let__14322

tmp16334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14322, Nil)
}
__typedArg0 := let__14322
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16317 Obj

if True == tmp16334 {
tmp16318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14309)
}
__typedArg0 := let__14309
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14323 := tmp16318
_ = let__14323

tmp16319 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16319

tmp16320 := Call(__e, PrimFunc(symshen_4deref), let__14319, V5357)


tmp16321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16320, tmp16321)
}
__typedArg0 := tmp16320
__typedArg1 := tmp16321
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16323 := MakeNative(func(__e *ControlFlow) {
tmp16324 := MakeNative(func(__e *ControlFlow) {
tmp16325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14319, tmp16325)
}
__typedArg0 := let__14319
__typedArg1 := tmp16325
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14314, tmp16326)
}
__typedArg0 := let__14314
__typedArg1 := tmp16326
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symstring, Nil)
}
__typedArg0 := symstring
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14319, tmp16328)
}
__typedArg0 := let__14319
__typedArg1 := tmp16328
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14316, tmp16329)
}
__typedArg0 := let__14316
__typedArg1 := tmp16329
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16330, let__14323)
}
__typedArg0 := tmp16330
__typedArg1 := let__14323
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16327, tmp16331)
}
__typedArg0 := tmp16327
__typedArg1 := tmp16331
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16332, V5355, True, V5357, V5358, let__14262, V5360)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5357, V5358, let__14262, tmp16324)
return


}, 0)

tmp16333 := Call(__e, PrimFunc(symwhen), tmp16322, V5357, V5358, let__14262, tmp16323)


ifres16317 = tmp16333


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

ifres16307 = ifres16311


} else {
ifres16307 = False


}

ifres16304 = ifres16307


} else {
ifres16304 = False


}

ifres16300 = ifres16304


} else {
ifres16300 = False


}

ifres16296 = ifres16300


} else {
ifres16296 = False


}

ifres16293 = ifres16296


} else {
ifres16293 = False


}

ifres16290 = ifres16293


} else {
ifres16290 = False


}

ifres16287 = ifres16290


} else {
ifres16287 = False


}

ifres16284 = ifres16287


} else {
ifres16284 = False


}

ifres16282 = ifres16284


} else {
ifres16282 = False


}

let__14308 := ifres16282
_ = let__14308

tmp16457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14308, False)
}
__typedArg0 := let__14308
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16457 {
tmp16424 := Call(__e, PrimFunc(symshen_4unlocked_2), V5358)


var ifres16346 Obj

if True == tmp16424 {
tmp16347 := Call(__e, PrimFunc(symshen_4lazyderef), V5354, V5357)


let__14325 := tmp16347
_ = let__14325

tmp16423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14325)
}
__typedArg0 := let__14325
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16348 Obj

if True == tmp16423 {
tmp16349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14325)
}
__typedArg0 := let__14325
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16350 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16349, V5357)


let__14326 := tmp16350
_ = let__14326

tmp16422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14326)
}
__typedArg0 := let__14326
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16351 Obj

if True == tmp16422 {
tmp16352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14326)
}
__typedArg0 := let__14326
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16353 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16352, V5357)


let__14327 := tmp16353
_ = let__14327

tmp16421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14327)
}
__typedArg0 := let__14327
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16354 Obj

if True == tmp16421 {
tmp16355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14327)
}
__typedArg0 := let__14327
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16356 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16355, V5357)


let__14328 := tmp16356
_ = let__14328

tmp16420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14328, sym_8v)
}
__typedArg0 := let__14328
__typedArg1 := sym_8v
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16357 Obj

if True == tmp16420 {
tmp16358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14327)
}
__typedArg0 := let__14327
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16359 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16358, V5357)


let__14329 := tmp16359
_ = let__14329

tmp16419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14329)
}
__typedArg0 := let__14329
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16360 Obj

if True == tmp16419 {
tmp16361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14329)
}
__typedArg0 := let__14329
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14330 := tmp16361
_ = let__14330

tmp16362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14329)
}
__typedArg0 := let__14329
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16363 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16362, V5357)


let__14331 := tmp16363
_ = let__14331

tmp16418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14331)
}
__typedArg0 := let__14331
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16364 Obj

if True == tmp16418 {
tmp16365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14331)
}
__typedArg0 := let__14331
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14332 := tmp16365
_ = let__14332

tmp16366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14331)
}
__typedArg0 := let__14331
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16367 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16366, V5357)


let__14333 := tmp16367
_ = let__14333

tmp16417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14333, Nil)
}
__typedArg0 := let__14333
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16368 Obj

if True == tmp16417 {
tmp16369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14326)
}
__typedArg0 := let__14326
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16370 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16369, V5357)


let__14334 := tmp16370
_ = let__14334

tmp16416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14334)
}
__typedArg0 := let__14334
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16371 Obj

if True == tmp16416 {
tmp16372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14334)
}
__typedArg0 := let__14334
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14335 := tmp16372
_ = let__14335

tmp16373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14334)
}
__typedArg0 := let__14334
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16374 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16373, V5357)


let__14336 := tmp16374
_ = let__14336

tmp16415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14336)
}
__typedArg0 := let__14336
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16375 Obj

if True == tmp16415 {
tmp16376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14336)
}
__typedArg0 := let__14336
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16377 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16376, V5357)


let__14337 := tmp16377
_ = let__14337

tmp16414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14337)
}
__typedArg0 := let__14337
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16378 Obj

if True == tmp16414 {
tmp16379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14337)
}
__typedArg0 := let__14337
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16380 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16379, V5357)


let__14338 := tmp16380
_ = let__14338

tmp16413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14338, symvector)
}
__typedArg0 := let__14338
__typedArg1 := symvector
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16381 Obj

if True == tmp16413 {
tmp16382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14337)
}
__typedArg0 := let__14337
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16383 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16382, V5357)


let__14339 := tmp16383
_ = let__14339

tmp16412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14339)
}
__typedArg0 := let__14339
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16384 Obj

if True == tmp16412 {
tmp16385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14339)
}
__typedArg0 := let__14339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14340 := tmp16385
_ = let__14340

tmp16386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14339)
}
__typedArg0 := let__14339
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16387 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16386, V5357)


let__14341 := tmp16387
_ = let__14341

tmp16411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14341, Nil)
}
__typedArg0 := let__14341
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16388 Obj

if True == tmp16411 {
tmp16389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14336)
}
__typedArg0 := let__14336
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16390 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16389, V5357)


let__14342 := tmp16390
_ = let__14342

tmp16410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14342, Nil)
}
__typedArg0 := let__14342
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16391 Obj

if True == tmp16410 {
tmp16392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14325)
}
__typedArg0 := let__14325
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14343 := tmp16392
_ = let__14343

tmp16393 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16393

tmp16394 := Call(__e, PrimFunc(symshen_4deref), let__14335, V5357)


tmp16395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp16396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp16394, tmp16395)
}
__typedArg0 := tmp16394
__typedArg1 := tmp16395
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp16397 := MakeNative(func(__e *ControlFlow) {
tmp16398 := MakeNative(func(__e *ControlFlow) {
tmp16399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14340, Nil)
}
__typedArg0 := let__14340
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14335, tmp16399)
}
__typedArg0 := let__14335
__typedArg1 := tmp16399
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14330, tmp16400)
}
__typedArg0 := let__14330
__typedArg1 := tmp16400
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14340, Nil)
}
__typedArg0 := let__14340
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp16402)
}
__typedArg0 := symvector
__typedArg1 := tmp16402
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16403, Nil)
}
__typedArg0 := tmp16403
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14335, tmp16404)
}
__typedArg0 := let__14335
__typedArg1 := tmp16404
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14332, tmp16405)
}
__typedArg0 := let__14332
__typedArg1 := tmp16405
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16406, let__14343)
}
__typedArg0 := tmp16406
__typedArg1 := let__14343
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16401, tmp16407)
}
__typedArg0 := tmp16401
__typedArg1 := tmp16407
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp16408, V5355, True, V5357, V5358, let__14262, V5360)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5357, V5358, let__14262, tmp16398)
return


}, 0)

tmp16409 := Call(__e, PrimFunc(symwhen), tmp16396, V5357, V5358, let__14262, tmp16397)


ifres16391 = tmp16409


} else {
ifres16391 = False


}

ifres16388 = ifres16391


} else {
ifres16388 = False


}

ifres16384 = ifres16388


} else {
ifres16384 = False


}

ifres16381 = ifres16384


} else {
ifres16381 = False


}

ifres16378 = ifres16381


} else {
ifres16378 = False


}

ifres16375 = ifres16378


} else {
ifres16375 = False


}

ifres16371 = ifres16375


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

ifres16360 = ifres16364


} else {
ifres16360 = False


}

ifres16357 = ifres16360


} else {
ifres16357 = False


}

ifres16354 = ifres16357


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

ifres16346 = ifres16348


} else {
ifres16346 = False


}

let__14324 := ifres16346
_ = let__14324

tmp16455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14324, False)
}
__typedArg0 := let__14324
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16455 {
tmp16451 := Call(__e, PrimFunc(symshen_4unlocked_2), V5358)


var ifres16425 Obj

if True == tmp16451 {
tmp16426 := Call(__e, PrimFunc(symshen_4lazyderef), V5354, V5357)


let__14345 := tmp16426
_ = let__14345

tmp16450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14345)
}
__typedArg0 := let__14345
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16427 Obj

if True == tmp16450 {
tmp16428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14345)
}
__typedArg0 := let__14345
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14346 := tmp16428
_ = let__14346

tmp16429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14345)
}
__typedArg0 := let__14345
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14347 := tmp16429
_ = let__14347

tmp16430 := Call(__e, PrimFunc(symshen_4lazyderef), V5355, V5357)


let__14348 := tmp16430
_ = let__14348

tmp16431 := MakeNative(func(__e *ControlFlow) {
Z5449 := __e.Get(1)
_ = Z5449
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5450 := __e.Get(1)
_ = Z5450
tmp16432 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16432

tmp16433 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4l_1rules), let__14347, Z5450, V5356, V5357, V5358, let__14262, V5360)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5449, let__14346, V5357, V5358, let__14262, tmp16433)
return


}, 1))
return
}, 1)

let__14349 := tmp16431
_ = let__14349

tmp16449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14348)
}
__typedArg0 := let__14348
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16434 Obj

if True == tmp16449 {
tmp16435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14348)
}
__typedArg0 := let__14348
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14350 := tmp16435
_ = let__14350

tmp16436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14348)
}
__typedArg0 := let__14348
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14351 := tmp16436
_ = let__14351

tmp16437 := Call(__e, let__14349, let__14350)


tmp16438 := Call(__e, tmp16437, let__14351)


ifres16434 = tmp16438


} else {
tmp16448 := Call(__e, PrimFunc(symshen_4pvar_2), let__14348)


var ifres16439 Obj

if True == tmp16448 {
tmp16440 := Call(__e, PrimFunc(symshen_4newpv), V5357)


let__14352 := tmp16440
_ = let__14352

tmp16441 := Call(__e, PrimFunc(symshen_4newpv), V5357)


let__14353 := tmp16441
_ = let__14353

tmp16442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14352, let__14353)
}
__typedArg0 := let__14352
__typedArg1 := let__14353
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16443 := MakeNative(func(__e *ControlFlow) {
tmp16444 := Call(__e, let__14349, let__14352)


__e.TailApply(tmp16444, let__14353)
return


}, 0)

tmp16445 := Call(__e, PrimFunc(symshen_4bind_b), let__14348, tmp16442, V5357, tmp16443)


tmp16446 := Call(__e, PrimFunc(symshen_4gc), V5357, tmp16445)


tmp16447 := Call(__e, PrimFunc(symshen_4gc), V5357, tmp16446)


ifres16439 = tmp16447


} else {
ifres16439 = False


}

ifres16434 = ifres16439


}

ifres16427 = ifres16434


} else {
ifres16427 = False


}

ifres16425 = ifres16427


} else {
ifres16425 = False


}

let__14344 := ifres16425
_ = let__14344

tmp16453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14344, False)
}
__typedArg0 := let__14344
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16453 {
__e.TailApply(PrimFunc(symshen_4unlock), V5358, let__14262)
return
} else {
__e.Return(let__14344)
return
}


} else {
__e.Return(let__14324)
return
}


} else {
__e.Return(let__14308)
return
}


} else {
__e.Return(let__14286)
return
}


} else {
__e.Return(let__14266)
return
}


} else {
__e.Return(let__14263)
return
}


}, 7)

tmp16464 := Call(__e, ns2_1set, symshen_4l_1rules, tmp16108)


_ = tmp16464

tmp16465 := MakeNative(func(__e *ControlFlow) {
V5455 := __e.Get(1)
_ = V5455
V5456 := __e.Get(2)
_ = V5456
V5457 := __e.Get(3)
_ = V5457
V5458 := __e.Get(4)
_ = V5458
V5459 := __e.Get(5)
_ = V5459
V5460 := __e.Get(6)
_ = V5460
let__14354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5459)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5459
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__14354

tmp16505 := Call(__e, PrimFunc(symshen_4unlocked_2), V5458)


var ifres16467 Obj

if True == tmp16505 {
tmp16468 := Call(__e, PrimFunc(symshen_4lazyderef), V5455, V5457)


let__14356 := tmp16468
_ = let__14356

tmp16504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14356)
}
__typedArg0 := let__14356
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16469 Obj

if True == tmp16504 {
tmp16470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14356)
}
__typedArg0 := let__14356
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16471 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16470, V5457)


let__14357 := tmp16471
_ = let__14357

tmp16503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14357, symdefine)
}
__typedArg0 := let__14357
__typedArg1 := symdefine
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16472 Obj

if True == tmp16503 {
tmp16473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14356)
}
__typedArg0 := let__14356
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16474 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16473, V5457)


let__14358 := tmp16474
_ = let__14358

tmp16502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14358)
}
__typedArg0 := let__14358
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16475 Obj

if True == tmp16502 {
tmp16476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14358)
}
__typedArg0 := let__14358
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14359 := tmp16476
_ = let__14359

tmp16477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14358)
}
__typedArg0 := let__14358
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14360 := tmp16477
_ = let__14360

tmp16478 := Call(__e, PrimFunc(symshen_4newpv), V5457)


let__14361 := tmp16478
_ = let__14361

tmp16479 := Call(__e, PrimFunc(symshen_4newpv), V5457)


let__14362 := tmp16479
_ = let__14362

tmp16480 := Call(__e, PrimFunc(symshen_4newpv), V5457)


let__14363 := tmp16480
_ = let__14363

tmp16481 := Call(__e, PrimFunc(symshen_4newpv), V5457)


let__14364 := tmp16481
_ = let__14364

tmp16482 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16482

tmp16483 := MakeNative(func(__e *ControlFlow) {
tmp16484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14359, let__14360)
}
__typedArg0 := let__14359
__typedArg1 := let__14360
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16485 := Call(__e, PrimFunc(symshen_4sigxrules), tmp16484)


tmp16486 := MakeNative(func(__e *ControlFlow) {
tmp16487 := Call(__e, PrimFunc(symshen_4lazyderef), let__14361, V5457)


tmp16488 := Call(__e, PrimFunc(symfst), tmp16487)


tmp16489 := MakeNative(func(__e *ControlFlow) {
tmp16490 := Call(__e, PrimFunc(symshen_4lazyderef), let__14361, V5457)


tmp16491 := Call(__e, PrimFunc(symsnd), tmp16490)


tmp16492 := MakeNative(func(__e *ControlFlow) {
tmp16493 := Call(__e, PrimFunc(symshen_4deref), let__14364, V5457)


tmp16494 := Call(__e, PrimFunc(symshen_4freshen_1sig), tmp16493)


tmp16495 := MakeNative(func(__e *ControlFlow) {
tmp16496 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symis), let__14364, V5456, V5457, V5458, let__14354, V5460)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1rules), let__14359, let__14362, let__14363, MakeInteger(1), V5457, V5458, let__14354, tmp16496)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14363, tmp16494, V5457, V5458, let__14354, tmp16495)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14362, tmp16491, V5457, V5458, let__14354, tmp16492)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14364, tmp16488, V5457, V5458, let__14354, tmp16489)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14361, tmp16485, V5457, V5458, let__14354, tmp16486)
return


}, 0)

tmp16497 := Call(__e, PrimFunc(symshen_4cut), V5457, V5458, let__14354, tmp16483)


tmp16498 := Call(__e, PrimFunc(symshen_4gc), V5457, tmp16497)


tmp16499 := Call(__e, PrimFunc(symshen_4gc), V5457, tmp16498)


tmp16500 := Call(__e, PrimFunc(symshen_4gc), V5457, tmp16499)


tmp16501 := Call(__e, PrimFunc(symshen_4gc), V5457, tmp16500)


ifres16475 = tmp16501


} else {
ifres16475 = False


}

ifres16472 = ifres16475


} else {
ifres16472 = False


}

ifres16469 = ifres16472


} else {
ifres16469 = False


}

ifres16467 = ifres16469


} else {
ifres16467 = False


}

let__14355 := ifres16467
_ = let__14355

tmp16507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14355, False)
}
__typedArg0 := let__14355
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16507 {
__e.TailApply(PrimFunc(symshen_4unlock), V5458, let__14354)
return
} else {
__e.Return(let__14355)
return
}


}, 6)

tmp16508 := Call(__e, ns2_1set, symshen_4t_d, tmp16465)


_ = tmp16508

tmp16509 := MakeNative(func(__e *ControlFlow) {
V5472 := __e.Get(1)
_ = V5472
tmp16510 := MakeNative(func(__e *ControlFlow) {
Z5474 := __e.Get(1)
_ = Z5474
__e.TailApply(PrimFunc(symshen_4_5sig_drules_6), Z5474)
return
}, 1)

tmp16511 := Call(__e, PrimFunc(symcompile), tmp16510, V5472)


let__14365 := tmp16511
_ = let__14365

tmp16512 := Call(__e, PrimFunc(symfst), let__14365)


tmp16513 := MakeNative(func(__e *ControlFlow) {
Z5475 := __e.Get(1)
_ = Z5475
__e.TailApply(PrimFunc(symshen_4unmark_1rule_1custom_1patterns), Z5475)
return
}, 1)

tmp16514 := Call(__e, PrimFunc(symsnd), let__14365)


tmp16515 := Call(__e, PrimFunc(symmap), tmp16513, tmp16514)


__e.TailApply(PrimFunc(sym_8p), tmp16512, tmp16515)
return


}, 1)

tmp16516 := Call(__e, ns2_1set, symshen_4sigxrules, tmp16509)


_ = tmp16516

tmp16517 := MakeNative(func(__e *ControlFlow) {
V5476 := __e.Get(1)
_ = V5476
tmp16523 := Call(__e, PrimFunc(symtuple_2), V5476)


if True == tmp16523 {
tmp16518 := MakeNative(func(__e *ControlFlow) {
Z5477 := __e.Get(1)
_ = Z5477
__e.TailApply(PrimFunc(symshen_4unmark_1custom_1pattern), Z5477)
return
}, 1)

tmp16519 := Call(__e, PrimFunc(symfst), V5476)


tmp16520 := Call(__e, PrimFunc(symmap), tmp16518, tmp16519)


tmp16521 := Call(__e, PrimFunc(symsnd), V5476)


__e.TailApply(PrimFunc(sym_8p), tmp16520, tmp16521)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4unmark_1rule_1custom_1patterns)
return
}


}, 1)

tmp16524 := Call(__e, ns2_1set, symshen_4unmark_1rule_1custom_1patterns, tmp16517)


_ = tmp16524

tmp16525 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V5478 := __self1
_ = V5478
tmp16561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5478)
}
__typedArg0 := V5478
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16537 Obj

if True == tmp16561 {
tmp16559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5478)
}
__typedArg0 := V5478
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8p, tmp16559)
}
__typedArg0 := sym_8p
__typedArg1 := tmp16559
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16539 Obj

if True == tmp16560 {
tmp16557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5478)
}
__typedArg0 := V5478
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16557)
}
__typedArg0 := tmp16557
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16541 Obj

if True == tmp16558 {
tmp16554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5478)
}
__typedArg0 := V5478
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16554)
}
__typedArg0 := tmp16554
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4custom_1pattern, tmp16555)
}
__typedArg0 := symshen_4custom_1pattern
__typedArg1 := tmp16555
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16543 Obj

if True == tmp16556 {
tmp16551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5478)
}
__typedArg0 := V5478
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16551)
}
__typedArg0 := tmp16551
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16552)
}
__typedArg0 := tmp16552
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16545 Obj

if True == tmp16553 {
tmp16547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5478)
}
__typedArg0 := V5478
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16547)
}
__typedArg0 := tmp16547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16548)
}
__typedArg0 := tmp16548
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp16549)
}
__typedArg0 := Nil
__typedArg1 := tmp16549
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16546 Obj

if True == tmp16550 {
ifres16546 = True


} else {
ifres16546 = False


}

ifres16545 = ifres16546


} else {
ifres16545 = False


}

var ifres16544 Obj

if True == ifres16545 {
ifres16544 = True


} else {
ifres16544 = False


}

ifres16543 = ifres16544


} else {
ifres16543 = False


}

var ifres16542 Obj

if True == ifres16543 {
ifres16542 = True


} else {
ifres16542 = False


}

ifres16541 = ifres16542


} else {
ifres16541 = False


}

var ifres16540 Obj

if True == ifres16541 {
ifres16540 = True


} else {
ifres16540 = False


}

ifres16539 = ifres16540


} else {
ifres16539 = False


}

var ifres16538 Obj

if True == ifres16539 {
ifres16538 = True


} else {
ifres16538 = False


}

ifres16537 = ifres16538


} else {
ifres16537 = False


}

if True == ifres16537 {
tmp16526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5478)
}
__typedArg0 := V5478
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16526)
}
__typedArg0 := tmp16526
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16527)
}
__typedArg0 := tmp16527
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16529 := Call(__e, PrimFunc(symshen_4curry), tmp16528)


if PrimFunc(symshen_4unmark_1custom_1pattern) == __self {
__self1 = tmp16529
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4unmark_1custom_1pattern), tmp16529)
return


} else {
tmp16535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5478)
}
__typedArg0 := V5478
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp16535 {
tmp16530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5478)
}
__typedArg0 := V5478
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16531 := Call(__e, PrimFunc(symshen_4unmark_1custom_1pattern), tmp16530)


tmp16532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5478)
}
__typedArg0 := V5478
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16533 := Call(__e, PrimFunc(symshen_4unmark_1custom_1pattern), tmp16532)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16531, tmp16533)
}
__typedArg0 := tmp16531
__typedArg1 := tmp16533
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V5478)
return
}


}


}, 1)

tmp16562 := Call(__e, ns2_1set, symshen_4unmark_1custom_1pattern, tmp16525)


_ = tmp16562

tmp16563 := MakeNative(func(__e *ControlFlow) {
V5479 := __e.Get(1)
_ = V5479
tmp16590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5479)
}
__typedArg0 := V5479
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16564 Obj

if True == tmp16590 {
tmp16565 := Call(__e, PrimFunc(symtail), V5479)


let__14367 := tmp16565
_ = let__14367

tmp16588 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14367, sym_i)


var ifres16566 Obj

if True == tmp16588 {
tmp16567 := Call(__e, PrimFunc(symtail), let__14367)


let__14368 := tmp16567
_ = let__14368

tmp16568 := Call(__e, PrimFunc(symshen_4_5signature_6), let__14368)


let__14369 := tmp16568
_ = let__14369

tmp16586 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14369)


var ifres16569 Obj

if True == tmp16586 {
tmp16570 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16569 = tmp16570


} else {
tmp16571 := Call(__e, PrimFunc(symshen_4_5_1out), let__14369)


let__14370 := tmp16571
_ = let__14370

tmp16572 := Call(__e, PrimFunc(symshen_4in_1_6), let__14369)


let__14371 := tmp16572
_ = let__14371

tmp16585 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14371, sym_j)


var ifres16573 Obj

if True == tmp16585 {
tmp16574 := Call(__e, PrimFunc(symtail), let__14371)


let__14372 := tmp16574
_ = let__14372

tmp16575 := Call(__e, PrimFunc(symshen_4_5rules_d_6), let__14372)


let__14373 := tmp16575
_ = let__14373

tmp16583 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14373)


var ifres16576 Obj

if True == tmp16583 {
tmp16577 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16576 = tmp16577


} else {
tmp16578 := Call(__e, PrimFunc(symshen_4_5_1out), let__14373)


let__14374 := tmp16578
_ = let__14374

tmp16579 := Call(__e, PrimFunc(symshen_4in_1_6), let__14373)


let__14375 := tmp16579
_ = let__14375

tmp16580 := Call(__e, PrimFunc(symshen_4rectify_1type), let__14370)


let__14376 := tmp16580
_ = let__14376

tmp16581 := Call(__e, PrimFunc(sym_8p), let__14376, let__14374)


tmp16582 := Call(__e, PrimFunc(symshen_4comb), let__14375, tmp16581)


ifres16576 = tmp16582


}

ifres16573 = ifres16576


} else {
tmp16584 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16573 = tmp16584


}

ifres16569 = ifres16573


}

ifres16566 = ifres16569


} else {
tmp16587 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16566 = tmp16587


}

ifres16564 = ifres16566


} else {
tmp16589 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16564 = tmp16589


}

let__14366 := ifres16564
_ = let__14366

tmp16592 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14366)


if True == tmp16592 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__14366)
return
}


}, 1)

tmp16593 := Call(__e, ns2_1set, symshen_4_5sig_drules_6, tmp16563)


_ = tmp16593

tmp16594 := MakeNative(func(__e *ControlFlow) {
V5491 := __e.Get(1)
_ = V5491
tmp16595 := Call(__e, PrimFunc(symshen_4extract_1vars), V5491)


let__14377 := tmp16595
_ = let__14377

tmp16596 := MakeNative(func(__e *ControlFlow) {
Z5494 := __e.Get(1)
_ = Z5494
tmp16597 := Call(__e, PrimFunc(symconcat), sym_e, Z5494)


tmp16598 := Call(__e, PrimFunc(symshen_4freshterm), tmp16597)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z5494, tmp16598)
}
__typedArg0 := Z5494
__typedArg1 := tmp16598
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp16599 := Call(__e, PrimFunc(symmap), tmp16596, let__14377)


let__14378 := tmp16599
_ = let__14378

__e.TailApply(PrimFunc(symshen_4freshen_1type), let__14378, V5491)
return


}, 1)

tmp16600 := Call(__e, ns2_1set, symshen_4freshen_1sig, tmp16594)


_ = tmp16600

tmp16601 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V5495 := __self1
_ = V5495
V5496 := __self2
_ = V5496
tmp16615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5495)
}
__typedArg0 := Nil
__typedArg1 := V5495
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16615 {
__e.Return(V5496)
return
} else {
tmp16613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5495)
}
__typedArg0 := V5495
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16609 Obj

if True == tmp16613 {
tmp16611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5495)
}
__typedArg0 := V5495
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16611)
}
__typedArg0 := tmp16611
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16610 Obj

if True == tmp16612 {
ifres16610 = True


} else {
ifres16610 = False


}

ifres16609 = ifres16610


} else {
ifres16609 = False


}

if True == ifres16609 {
tmp16602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5495)
}
__typedArg0 := V5495
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5495)
}
__typedArg0 := V5495
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16603)
}
__typedArg0 := tmp16603
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5495)
}
__typedArg0 := V5495
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16605)
}
__typedArg0 := tmp16605
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16607 := Call(__e, PrimFunc(symsubst), tmp16604, tmp16606, V5496)


if PrimFunc(symshen_4freshen_1type) == __self {
__self1, __self2 = tmp16602, tmp16607
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4freshen_1type), tmp16602, tmp16607)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshen_1type)
return
}


}


}, 2)

tmp16616 := Call(__e, ns2_1set, symshen_4freshen_1type, tmp16601)


_ = tmp16616

tmp16617 := MakeNative(func(__e *ControlFlow) {
V5497 := __e.Get(1)
_ = V5497
tmp16618 := Call(__e, PrimFunc(symshen_4_5rule_d_6), V5497)


let__14380 := tmp16618
_ = let__14380

tmp16631 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14380)


var ifres16619 Obj

if True == tmp16631 {
tmp16620 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16619 = tmp16620


} else {
tmp16621 := Call(__e, PrimFunc(symshen_4_5_1out), let__14380)


let__14381 := tmp16621
_ = let__14381

tmp16622 := Call(__e, PrimFunc(symshen_4in_1_6), let__14380)


let__14382 := tmp16622
_ = let__14382

tmp16623 := Call(__e, PrimFunc(symshen_4_5rules_d_6), let__14382)


let__14383 := tmp16623
_ = let__14383

tmp16630 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14383)


var ifres16624 Obj

if True == tmp16630 {
tmp16625 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16624 = tmp16625


} else {
tmp16626 := Call(__e, PrimFunc(symshen_4_5_1out), let__14383)


let__14384 := tmp16626
_ = let__14384

tmp16627 := Call(__e, PrimFunc(symshen_4in_1_6), let__14383)


let__14385 := tmp16627
_ = let__14385

tmp16628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14381, let__14384)
}
__typedArg0 := let__14381
__typedArg1 := let__14384
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16629 := Call(__e, PrimFunc(symshen_4comb), let__14385, tmp16628)


ifres16624 = tmp16629


}

ifres16619 = ifres16624


}

let__14379 := ifres16619
_ = let__14379

tmp16643 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14379)


if True == tmp16643 {
tmp16632 := Call(__e, PrimFunc(symshen_4_5rule_d_6), V5497)


let__14387 := tmp16632
_ = let__14387

tmp16639 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14387)


var ifres16633 Obj

if True == tmp16639 {
tmp16634 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16633 = tmp16634


} else {
tmp16635 := Call(__e, PrimFunc(symshen_4_5_1out), let__14387)


let__14388 := tmp16635
_ = let__14388

tmp16636 := Call(__e, PrimFunc(symshen_4in_1_6), let__14387)


let__14389 := tmp16636
_ = let__14389

tmp16637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14388, Nil)
}
__typedArg0 := let__14388
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16638 := Call(__e, PrimFunc(symshen_4comb), let__14389, tmp16637)


ifres16633 = tmp16638


}

let__14386 := ifres16633
_ = let__14386

tmp16641 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14386)


if True == tmp16641 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__14386)
return
}


} else {
__e.Return(let__14379)
return
}


}, 1)

tmp16644 := Call(__e, ns2_1set, symshen_4_5rules_d_6, tmp16617)


_ = tmp16644

tmp16645 := MakeNative(func(__e *ControlFlow) {
V5509 := __e.Get(1)
_ = V5509
tmp16646 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5509)


let__14391 := tmp16646
_ = let__14391

tmp16674 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14391)


var ifres16647 Obj

if True == tmp16674 {
tmp16648 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16647 = tmp16648


} else {
tmp16649 := Call(__e, PrimFunc(symshen_4_5_1out), let__14391)


let__14392 := tmp16649
_ = let__14392

tmp16650 := Call(__e, PrimFunc(symshen_4in_1_6), let__14391)


let__14393 := tmp16650
_ = let__14393

tmp16673 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14393, sym_1_6)


var ifres16651 Obj

if True == tmp16673 {
tmp16652 := Call(__e, PrimFunc(symtail), let__14393)


let__14394 := tmp16652
_ = let__14394

tmp16671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14394)
}
__typedArg0 := let__14394
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16653 Obj

if True == tmp16671 {
tmp16654 := Call(__e, PrimFunc(symhead), let__14394)


let__14395 := tmp16654
_ = let__14395

tmp16655 := Call(__e, PrimFunc(symtail), let__14394)


let__14396 := tmp16655
_ = let__14396

tmp16669 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14396, symwhere)


var ifres16656 Obj

if True == tmp16669 {
tmp16657 := Call(__e, PrimFunc(symtail), let__14396)


let__14397 := tmp16657
_ = let__14397

tmp16667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14397)
}
__typedArg0 := let__14397
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16658 Obj

if True == tmp16667 {
tmp16659 := Call(__e, PrimFunc(symhead), let__14397)


let__14398 := tmp16659
_ = let__14398

tmp16660 := Call(__e, PrimFunc(symtail), let__14397)


let__14399 := tmp16660
_ = let__14399

tmp16661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14395, Nil)
}
__typedArg0 := let__14395
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14398, tmp16661)
}
__typedArg0 := let__14398
__typedArg1 := tmp16661
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16662)
}
__typedArg0 := symwhere
__typedArg1 := tmp16662
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16664 := Call(__e, PrimFunc(sym_8p), let__14392, tmp16663)


tmp16665 := Call(__e, PrimFunc(symshen_4comb), let__14399, tmp16664)


ifres16658 = tmp16665


} else {
tmp16666 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16658 = tmp16666


}

ifres16656 = ifres16658


} else {
tmp16668 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16656 = tmp16668


}

ifres16653 = ifres16656


} else {
tmp16670 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16653 = tmp16670


}

ifres16651 = ifres16653


} else {
tmp16672 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16651 = tmp16672


}

ifres16647 = ifres16651


}

let__14390 := ifres16647
_ = let__14390

tmp16747 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14390)


if True == tmp16747 {
tmp16675 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5509)


let__14401 := tmp16675
_ = let__14401

tmp16704 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14401)


var ifres16676 Obj

if True == tmp16704 {
tmp16677 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16676 = tmp16677


} else {
tmp16678 := Call(__e, PrimFunc(symshen_4_5_1out), let__14401)


let__14402 := tmp16678
_ = let__14402

tmp16679 := Call(__e, PrimFunc(symshen_4in_1_6), let__14401)


let__14403 := tmp16679
_ = let__14403

tmp16703 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14403, sym_5_1)


var ifres16680 Obj

if True == tmp16703 {
tmp16681 := Call(__e, PrimFunc(symtail), let__14403)


let__14404 := tmp16681
_ = let__14404

tmp16701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14404)
}
__typedArg0 := let__14404
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16682 Obj

if True == tmp16701 {
tmp16683 := Call(__e, PrimFunc(symhead), let__14404)


let__14405 := tmp16683
_ = let__14405

tmp16684 := Call(__e, PrimFunc(symtail), let__14404)


let__14406 := tmp16684
_ = let__14406

tmp16699 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14406, symwhere)


var ifres16685 Obj

if True == tmp16699 {
tmp16686 := Call(__e, PrimFunc(symtail), let__14406)


let__14407 := tmp16686
_ = let__14407

tmp16697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14407)
}
__typedArg0 := let__14407
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16687 Obj

if True == tmp16697 {
tmp16688 := Call(__e, PrimFunc(symhead), let__14407)


let__14408 := tmp16688
_ = let__14408

tmp16689 := Call(__e, PrimFunc(symtail), let__14407)


let__14409 := tmp16689
_ = let__14409

tmp16690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14405, Nil)
}
__typedArg0 := let__14405
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14408, tmp16690)
}
__typedArg0 := let__14408
__typedArg1 := tmp16690
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16691)
}
__typedArg0 := symwhere
__typedArg1 := tmp16691
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16693 := Call(__e, PrimFunc(symshen_4correct), tmp16692)


tmp16694 := Call(__e, PrimFunc(sym_8p), let__14402, tmp16693)


tmp16695 := Call(__e, PrimFunc(symshen_4comb), let__14409, tmp16694)


ifres16687 = tmp16695


} else {
tmp16696 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16687 = tmp16696


}

ifres16685 = ifres16687


} else {
tmp16698 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16685 = tmp16698


}

ifres16682 = ifres16685


} else {
tmp16700 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16682 = tmp16700


}

ifres16680 = ifres16682


} else {
tmp16702 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16680 = tmp16702


}

ifres16676 = ifres16680


}

let__14400 := ifres16676
_ = let__14400

tmp16745 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14400)


if True == tmp16745 {
tmp16705 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5509)


let__14411 := tmp16705
_ = let__14411

tmp16722 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14411)


var ifres16706 Obj

if True == tmp16722 {
tmp16707 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16706 = tmp16707


} else {
tmp16708 := Call(__e, PrimFunc(symshen_4_5_1out), let__14411)


let__14412 := tmp16708
_ = let__14412

tmp16709 := Call(__e, PrimFunc(symshen_4in_1_6), let__14411)


let__14413 := tmp16709
_ = let__14413

tmp16721 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14413, sym_5_1)


var ifres16710 Obj

if True == tmp16721 {
tmp16711 := Call(__e, PrimFunc(symtail), let__14413)


let__14414 := tmp16711
_ = let__14414

tmp16719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14414)
}
__typedArg0 := let__14414
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16712 Obj

if True == tmp16719 {
tmp16713 := Call(__e, PrimFunc(symhead), let__14414)


let__14415 := tmp16713
_ = let__14415

tmp16714 := Call(__e, PrimFunc(symtail), let__14414)


let__14416 := tmp16714
_ = let__14416

tmp16715 := Call(__e, PrimFunc(symshen_4correct), let__14415)


tmp16716 := Call(__e, PrimFunc(sym_8p), let__14412, tmp16715)


tmp16717 := Call(__e, PrimFunc(symshen_4comb), let__14416, tmp16716)


ifres16712 = tmp16717


} else {
tmp16718 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16712 = tmp16718


}

ifres16710 = ifres16712


} else {
tmp16720 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16710 = tmp16720


}

ifres16706 = ifres16710


}

let__14410 := ifres16706
_ = let__14410

tmp16743 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14410)


if True == tmp16743 {
tmp16723 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5509)


let__14418 := tmp16723
_ = let__14418

tmp16739 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14418)


var ifres16724 Obj

if True == tmp16739 {
tmp16725 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16724 = tmp16725


} else {
tmp16726 := Call(__e, PrimFunc(symshen_4_5_1out), let__14418)


let__14419 := tmp16726
_ = let__14419

tmp16727 := Call(__e, PrimFunc(symshen_4in_1_6), let__14418)


let__14420 := tmp16727
_ = let__14420

tmp16738 := Call(__e, PrimFunc(symshen_4hds_a_2), let__14420, sym_1_6)


var ifres16728 Obj

if True == tmp16738 {
tmp16729 := Call(__e, PrimFunc(symtail), let__14420)


let__14421 := tmp16729
_ = let__14421

tmp16736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14421)
}
__typedArg0 := let__14421
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16730 Obj

if True == tmp16736 {
tmp16731 := Call(__e, PrimFunc(symhead), let__14421)


let__14422 := tmp16731
_ = let__14422

tmp16732 := Call(__e, PrimFunc(symtail), let__14421)


let__14423 := tmp16732
_ = let__14423

tmp16733 := Call(__e, PrimFunc(sym_8p), let__14419, let__14422)


tmp16734 := Call(__e, PrimFunc(symshen_4comb), let__14423, tmp16733)


ifres16730 = tmp16734


} else {
tmp16735 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16730 = tmp16735


}

ifres16728 = ifres16730


} else {
tmp16737 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16728 = tmp16737


}

ifres16724 = ifres16728


}

let__14417 := ifres16724
_ = let__14417

tmp16741 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__14417)


if True == tmp16741 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__14417)
return
}


} else {
__e.Return(let__14410)
return
}


} else {
__e.Return(let__14400)
return
}


} else {
__e.Return(let__14390)
return
}


}, 1)

tmp16748 := Call(__e, ns2_1set, symshen_4_5rule_d_6, tmp16645)


_ = tmp16748

tmp16749 := MakeNative(func(__e *ControlFlow) {
V5544 := __e.Get(1)
_ = V5544
tmp16897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5544)
}
__typedArg0 := V5544
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16841 Obj

if True == tmp16897 {
tmp16895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5544)
}
__typedArg0 := V5544
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp16895)
}
__typedArg0 := symwhere
__typedArg1 := tmp16895
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16843 Obj

if True == tmp16896 {
tmp16893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5544)
}
__typedArg0 := V5544
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16893)
}
__typedArg0 := tmp16893
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16845 Obj

if True == tmp16894 {
tmp16890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5544)
}
__typedArg0 := V5544
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16890)
}
__typedArg0 := tmp16890
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16891)
}
__typedArg0 := tmp16891
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16847 Obj

if True == tmp16892 {
tmp16886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5544)
}
__typedArg0 := V5544
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16886)
}
__typedArg0 := tmp16886
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16887)
}
__typedArg0 := tmp16887
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16888)
}
__typedArg0 := tmp16888
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16849 Obj

if True == tmp16889 {
tmp16881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5544)
}
__typedArg0 := V5544
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16881)
}
__typedArg0 := tmp16881
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16882)
}
__typedArg0 := tmp16882
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16883)
}
__typedArg0 := tmp16883
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail_1if, tmp16884)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp16884
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16851 Obj

if True == tmp16885 {
tmp16876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5544)
}
__typedArg0 := V5544
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16876)
}
__typedArg0 := tmp16876
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16877)
}
__typedArg0 := tmp16877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16878)
}
__typedArg0 := tmp16878
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16879)
}
__typedArg0 := tmp16879
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16853 Obj

if True == tmp16880 {
tmp16870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5544)
}
__typedArg0 := V5544
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16871)
}
__typedArg0 := tmp16871
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16872)
}
__typedArg0 := tmp16872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16873)
}
__typedArg0 := tmp16873
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16874)
}
__typedArg0 := tmp16874
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16855 Obj

if True == tmp16875 {
tmp16863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5544)
}
__typedArg0 := V5544
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16863)
}
__typedArg0 := tmp16863
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16864)
}
__typedArg0 := tmp16864
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16865)
}
__typedArg0 := tmp16865
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16867)
}
__typedArg0 := tmp16867
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp16868)
}
__typedArg0 := Nil
__typedArg1 := tmp16868
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16857 Obj

if True == tmp16869 {
tmp16859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5544)
}
__typedArg0 := V5544
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16859)
}
__typedArg0 := tmp16859
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp16861)
}
__typedArg0 := Nil
__typedArg1 := tmp16861
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16858 Obj

if True == tmp16862 {
ifres16858 = True


} else {
ifres16858 = False


}

ifres16857 = ifres16858


} else {
ifres16857 = False


}

var ifres16856 Obj

if True == ifres16857 {
ifres16856 = True


} else {
ifres16856 = False


}

ifres16855 = ifres16856


} else {
ifres16855 = False


}

var ifres16854 Obj

if True == ifres16855 {
ifres16854 = True


} else {
ifres16854 = False


}

ifres16853 = ifres16854


} else {
ifres16853 = False


}

var ifres16852 Obj

if True == ifres16853 {
ifres16852 = True


} else {
ifres16852 = False


}

ifres16851 = ifres16852


} else {
ifres16851 = False


}

var ifres16850 Obj

if True == ifres16851 {
ifres16850 = True


} else {
ifres16850 = False


}

ifres16849 = ifres16850


} else {
ifres16849 = False


}

var ifres16848 Obj

if True == ifres16849 {
ifres16848 = True


} else {
ifres16848 = False


}

ifres16847 = ifres16848


} else {
ifres16847 = False


}

var ifres16846 Obj

if True == ifres16847 {
ifres16846 = True


} else {
ifres16846 = False


}

ifres16845 = ifres16846


} else {
ifres16845 = False


}

var ifres16844 Obj

if True == ifres16845 {
ifres16844 = True


} else {
ifres16844 = False


}

ifres16843 = ifres16844


} else {
ifres16843 = False


}

var ifres16842 Obj

if True == ifres16843 {
ifres16842 = True


} else {
ifres16842 = False


}

ifres16841 = ifres16842


} else {
ifres16841 = False


}

if True == ifres16841 {
tmp16750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5544)
}
__typedArg0 := V5544
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5544)
}
__typedArg0 := V5544
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16752)
}
__typedArg0 := tmp16752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16753)
}
__typedArg0 := tmp16753
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16754)
}
__typedArg0 := tmp16754
return Call(__e, PrimFunc(symtl), __typedArg0)
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
return PrimCons(tmp16751, tmp16758)
}
__typedArg0 := tmp16751
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
return PrimTail(V5544)
}
__typedArg0 := V5544
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16762)
}
__typedArg0 := tmp16762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16763)
}
__typedArg0 := tmp16763
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16764)
}
__typedArg0 := tmp16764
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16760, tmp16765)
}
__typedArg0 := tmp16760
__typedArg1 := tmp16765
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16766)
}
__typedArg0 := symwhere
__typedArg1 := tmp16766
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp16839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5544)
}
__typedArg0 := V5544
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16820 Obj

if True == tmp16839 {
tmp16837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5544)
}
__typedArg0 := V5544
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp16837)
}
__typedArg0 := symwhere
__typedArg1 := tmp16837
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16822 Obj

if True == tmp16838 {
tmp16835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5544)
}
__typedArg0 := V5544
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16835)
}
__typedArg0 := tmp16835
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16824 Obj

if True == tmp16836 {
tmp16832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5544)
}
__typedArg0 := V5544
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16832)
}
__typedArg0 := tmp16832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16833)
}
__typedArg0 := tmp16833
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16826 Obj

if True == tmp16834 {
tmp16828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5544)
}
__typedArg0 := V5544
return Call(__e, PrimFunc(symtl), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp16830)
}
__typedArg0 := Nil
__typedArg1 := tmp16830
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16827 Obj

if True == tmp16831 {
ifres16827 = True


} else {
ifres16827 = False


}

ifres16826 = ifres16827


} else {
ifres16826 = False


}

var ifres16825 Obj

if True == ifres16826 {
ifres16825 = True


} else {
ifres16825 = False


}

ifres16824 = ifres16825


} else {
ifres16824 = False


}

var ifres16823 Obj

if True == ifres16824 {
ifres16823 = True


} else {
ifres16823 = False


}

ifres16822 = ifres16823


} else {
ifres16822 = False


}

var ifres16821 Obj

if True == ifres16822 {
ifres16821 = True


} else {
ifres16821 = False


}

ifres16820 = ifres16821


} else {
ifres16820 = False


}

if True == ifres16820 {
tmp16767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5544)
}
__typedArg0 := V5544
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16767)
}
__typedArg0 := tmp16767
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5544)
}
__typedArg0 := V5544
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16769)
}
__typedArg0 := tmp16769
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16770)
}
__typedArg0 := tmp16770
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, Nil)
}
__typedArg0 := symfail
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16772, Nil)
}
__typedArg0 := tmp16772
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16771, tmp16773)
}
__typedArg0 := tmp16771
__typedArg1 := tmp16773
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp16774)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp16774
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16775, Nil)
}
__typedArg0 := tmp16775
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp16776)
}
__typedArg0 := symnot
__typedArg1 := tmp16776
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16777, Nil)
}
__typedArg0 := tmp16777
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16768, tmp16778)
}
__typedArg0 := tmp16768
__typedArg1 := tmp16778
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp16779)
}
__typedArg0 := symand
__typedArg1 := tmp16779
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5544)
}
__typedArg0 := V5544
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16780, tmp16782)
}
__typedArg0 := tmp16780
__typedArg1 := tmp16782
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16783)
}
__typedArg0 := symwhere
__typedArg1 := tmp16783
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp16818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5544)
}
__typedArg0 := V5544
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16799 Obj

if True == tmp16818 {
tmp16816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5544)
}
__typedArg0 := V5544
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail_1if, tmp16816)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp16816
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16801 Obj

if True == tmp16817 {
tmp16814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5544)
}
__typedArg0 := V5544
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16814)
}
__typedArg0 := tmp16814
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16803 Obj

if True == tmp16815 {
tmp16811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5544)
}
__typedArg0 := V5544
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16811)
}
__typedArg0 := tmp16811
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16812)
}
__typedArg0 := tmp16812
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16805 Obj

if True == tmp16813 {
tmp16807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5544)
}
__typedArg0 := V5544
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16807)
}
__typedArg0 := tmp16807
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp16809)
}
__typedArg0 := Nil
__typedArg1 := tmp16809
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16806 Obj

if True == tmp16810 {
ifres16806 = True


} else {
ifres16806 = False


}

ifres16805 = ifres16806


} else {
ifres16805 = False


}

var ifres16804 Obj

if True == ifres16805 {
ifres16804 = True


} else {
ifres16804 = False


}

ifres16803 = ifres16804


} else {
ifres16803 = False


}

var ifres16802 Obj

if True == ifres16803 {
ifres16802 = True


} else {
ifres16802 = False


}

ifres16801 = ifres16802


} else {
ifres16801 = False


}

var ifres16800 Obj

if True == ifres16801 {
ifres16800 = True


} else {
ifres16800 = False


}

ifres16799 = ifres16800


} else {
ifres16799 = False


}

if True == ifres16799 {
tmp16784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5544)
}
__typedArg0 := V5544
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16784, Nil)
}
__typedArg0 := tmp16784
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp16785)
}
__typedArg0 := symnot
__typedArg1 := tmp16785
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5544)
}
__typedArg0 := V5544
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16786, tmp16788)
}
__typedArg0 := tmp16786
__typedArg1 := tmp16788
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16789)
}
__typedArg0 := symwhere
__typedArg1 := tmp16789
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp16790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, Nil)
}
__typedArg0 := symfail
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16790, Nil)
}
__typedArg0 := tmp16790
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5544, tmp16791)
}
__typedArg0 := V5544
__typedArg1 := tmp16791
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a, tmp16792)
}
__typedArg0 := sym_a_a
__typedArg1 := tmp16792
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16793, Nil)
}
__typedArg0 := tmp16793
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symnot, tmp16794)
}
__typedArg0 := symnot
__typedArg1 := tmp16794
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5544, Nil)
}
__typedArg0 := V5544
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp16797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp16795, tmp16796)
}
__typedArg0 := tmp16795
__typedArg1 := tmp16796
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp16797)
}
__typedArg0 := symwhere
__typedArg1 := tmp16797
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}


}


}, 1)

tmp16898 := Call(__e, ns2_1set, symshen_4correct, tmp16749)


_ = tmp16898

tmp16899 := MakeNative(func(__e *ControlFlow) {
V5545 := __e.Get(1)
_ = V5545
V5546 := __e.Get(2)
_ = V5546
V5547 := __e.Get(3)
_ = V5547
V5548 := __e.Get(4)
_ = V5548
V5549 := __e.Get(5)
_ = V5549
V5550 := __e.Get(6)
_ = V5550
V5551 := __e.Get(7)
_ = V5551
V5552 := __e.Get(8)
_ = V5552
let__14424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5551)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5551
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__14424

tmp16907 := Call(__e, PrimFunc(symshen_4unlocked_2), V5550)


var ifres16901 Obj

if True == tmp16907 {
tmp16902 := Call(__e, PrimFunc(symshen_4lazyderef), V5546, V5549)


let__14426 := tmp16902
_ = let__14426

tmp16906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14426, Nil)
}
__typedArg0 := let__14426
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16903 Obj

if True == tmp16906 {
tmp16904 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16904

tmp16905 := Call(__e, PrimFunc(symthaw), V5552)


ifres16903 = tmp16905


} else {
ifres16903 = False


}

ifres16901 = ifres16903


} else {
ifres16901 = False


}

let__14425 := ifres16901
_ = let__14425

tmp16932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14425, False)
}
__typedArg0 := let__14425
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16932 {
tmp16928 := Call(__e, PrimFunc(symshen_4unlocked_2), V5550)


var ifres16908 Obj

if True == tmp16928 {
tmp16909 := Call(__e, PrimFunc(symshen_4lazyderef), V5546, V5549)


let__14428 := tmp16909
_ = let__14428

tmp16927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14428)
}
__typedArg0 := let__14428
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16910 Obj

if True == tmp16927 {
tmp16911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14428)
}
__typedArg0 := let__14428
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14429 := tmp16911
_ = let__14429

tmp16912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14428)
}
__typedArg0 := let__14428
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14430 := tmp16912
_ = let__14430

tmp16913 := Call(__e, PrimFunc(symshen_4newpv), V5549)


let__14431 := tmp16913
_ = let__14431

tmp16914 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16914

tmp16915 := Call(__e, PrimFunc(symshen_4deref), let__14429, V5549)


tmp16916 := Call(__e, PrimFunc(symshen_4freshen_1rule), tmp16915)


tmp16917 := MakeNative(func(__e *ControlFlow) {
tmp16918 := Call(__e, PrimFunc(symshen_4lazyderef), let__14431, V5549)


tmp16919 := Call(__e, PrimFunc(symfst), tmp16918)


tmp16920 := Call(__e, PrimFunc(symshen_4lazyderef), let__14431, V5549)


tmp16921 := Call(__e, PrimFunc(symsnd), tmp16920)


tmp16922 := MakeNative(func(__e *ControlFlow) {
tmp16923 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1rules), V5545, let__14430, V5547, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5548)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5548
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})(), V5549, V5550, let__14424, V5552)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5549, V5550, let__14424, tmp16923)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1rule), V5545, V5548, tmp16919, tmp16921, V5547, V5549, V5550, let__14424, tmp16922)
return


}, 0)

tmp16925 := Call(__e, PrimFunc(symbind), let__14431, tmp16916, V5549, V5550, let__14424, tmp16917)


tmp16926 := Call(__e, PrimFunc(symshen_4gc), V5549, tmp16925)


ifres16910 = tmp16926


} else {
ifres16910 = False


}

ifres16908 = ifres16910


} else {
ifres16908 = False


}

let__14427 := ifres16908
_ = let__14427

tmp16930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14427, False)
}
__typedArg0 := let__14427
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16930 {
__e.TailApply(PrimFunc(symshen_4unlock), V5550, let__14424)
return
} else {
__e.Return(let__14427)
return
}


} else {
__e.Return(let__14425)
return
}


}, 8)

tmp16933 := Call(__e, ns2_1set, symshen_4t_d_1rules, tmp16899)


_ = tmp16933

tmp16934 := MakeNative(func(__e *ControlFlow) {
V5561 := __e.Get(1)
_ = V5561
tmp16945 := Call(__e, PrimFunc(symtuple_2), V5561)


if True == tmp16945 {
tmp16935 := Call(__e, PrimFunc(symfst), V5561)


tmp16936 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp16935)


let__14432 := tmp16936
_ = let__14432

tmp16937 := MakeNative(func(__e *ControlFlow) {
Z5564 := __e.Get(1)
_ = Z5564
tmp16938 := Call(__e, PrimFunc(symshen_4freshterm), Z5564)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z5564, tmp16938)
}
__typedArg0 := Z5564
__typedArg1 := tmp16938
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp16939 := Call(__e, PrimFunc(symmap), tmp16937, let__14432)


let__14433 := tmp16939
_ = let__14433

tmp16940 := Call(__e, PrimFunc(symfst), V5561)


tmp16941 := Call(__e, PrimFunc(symshen_4freshen), let__14433, tmp16940)


tmp16942 := Call(__e, PrimFunc(symsnd), V5561)


tmp16943 := Call(__e, PrimFunc(symshen_4freshen), let__14433, tmp16942)


__e.TailApply(PrimFunc(sym_8p), tmp16941, tmp16943)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshen_1rule)
return
}


}, 1)

tmp16946 := Call(__e, ns2_1set, symshen_4freshen_1rule, tmp16934)


_ = tmp16946

tmp16947 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V5565 := __self1
_ = V5565
V5566 := __self2
_ = V5566
tmp16961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5565)
}
__typedArg0 := Nil
__typedArg1 := V5565
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16961 {
__e.Return(V5566)
return
} else {
tmp16959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5565)
}
__typedArg0 := V5565
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16955 Obj

if True == tmp16959 {
tmp16957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5565)
}
__typedArg0 := V5565
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp16957)
}
__typedArg0 := tmp16957
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16956 Obj

if True == tmp16958 {
ifres16956 = True


} else {
ifres16956 = False


}

ifres16955 = ifres16956


} else {
ifres16955 = False


}

if True == ifres16955 {
tmp16948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5565)
}
__typedArg0 := V5565
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5565)
}
__typedArg0 := V5565
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp16949)
}
__typedArg0 := tmp16949
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5565)
}
__typedArg0 := V5565
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp16951)
}
__typedArg0 := tmp16951
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16953 := Call(__e, PrimFunc(symshen_4beta), tmp16950, tmp16952, V5566)


if PrimFunc(symshen_4freshen) == __self {
__self1, __self2 = tmp16948, tmp16953
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4freshen), tmp16948, tmp16953)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshen)
return
}


}


}, 2)

tmp16962 := Call(__e, ns2_1set, symshen_4freshen, tmp16947)


_ = tmp16962

tmp16963 := MakeNative(func(__e *ControlFlow) {
V5567 := __e.Get(1)
_ = V5567
V5568 := __e.Get(2)
_ = V5568
V5569 := __e.Get(3)
_ = V5569
V5570 := __e.Get(4)
_ = V5570
V5571 := __e.Get(5)
_ = V5571
V5572 := __e.Get(6)
_ = V5572
V5573 := __e.Get(7)
_ = V5573
V5574 := __e.Get(8)
_ = V5574
V5575 := __e.Get(9)
_ = V5575
tmp16967 := Call(__e, PrimFunc(symshen_4unlocked_2), V5573)


var ifres16964 Obj

if True == tmp16967 {
tmp16965 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16965

tmp16966 := Call(__e, PrimFunc(symshen_4t_d_1rule_1h), V5569, V5570, V5571, V5572, V5573, V5574, V5575)


ifres16964 = tmp16966


} else {
ifres16964 = False


}

let__14434 := ifres16964
_ = let__14434

tmp16979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14434, False)
}
__typedArg0 := let__14434
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp16979 {
tmp16977 := Call(__e, PrimFunc(symshen_4unlocked_2), V5573)


if True == tmp16977 {
tmp16968 := Call(__e, PrimFunc(symshen_4newpv), V5572)


let__14435 := tmp16968
_ = let__14435

tmp16969 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16969

tmp16970 := Call(__e, PrimFunc(symshen_4app), V5567, MakeString("\n"), symshen_4a)


tmp16972 := Call(__e, PrimFunc(symshen_4app), V5568, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" of "))
__typedS1, __typedOK1 := TypedString(tmp16970)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" of ")
__typedArg1 := tmp16970
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp16974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("type error in rule "))
__typedS1, __typedOK1 := TypedString(tmp16972)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("type error in rule ")
__typedArg1 := tmp16972
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("type error in rule "))
__typedS1, __typedOK1 := TypedString(tmp16972)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("type error in rule ")
__typedArg1 := tmp16972
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

tmp16975 := Call(__e, PrimFunc(symbind), let__14435, tmp16974, V5572, V5573, V5574, V5575)


__e.TailApply(PrimFunc(symshen_4gc), V5572, tmp16975)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(let__14434)
return
}


}, 9)

tmp16980 := Call(__e, ns2_1set, symshen_4t_d_1rule, tmp16963)


_ = tmp16980

tmp16981 := MakeNative(func(__e *ControlFlow) {
V5578 := __e.Get(1)
_ = V5578
V5579 := __e.Get(2)
_ = V5579
V5580 := __e.Get(3)
_ = V5580
V5581 := __e.Get(4)
_ = V5581
V5582 := __e.Get(5)
_ = V5582
V5583 := __e.Get(6)
_ = V5583
V5584 := __e.Get(7)
_ = V5584
let__14436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5583)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5583
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__14436

tmp17006 := Call(__e, PrimFunc(symshen_4unlocked_2), V5582)


var ifres16983 Obj

if True == tmp17006 {
tmp16984 := Call(__e, PrimFunc(symshen_4lazyderef), V5578, V5581)


let__14438 := tmp16984
_ = let__14438

tmp17005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14438, Nil)
}
__typedArg0 := let__14438
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16985 Obj

if True == tmp17005 {
tmp16986 := Call(__e, PrimFunc(symshen_4lazyderef), V5580, V5581)


let__14439 := tmp16986
_ = let__14439

tmp17004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14439)
}
__typedArg0 := let__14439
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16987 Obj

if True == tmp17004 {
tmp16988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14439)
}
__typedArg0 := let__14439
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp16989 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16988, V5581)


let__14440 := tmp16989
_ = let__14440

tmp17003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14440, sym_1_1_6)
}
__typedArg0 := let__14440
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16990 Obj

if True == tmp17003 {
tmp16991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14439)
}
__typedArg0 := let__14439
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16992 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16991, V5581)


let__14441 := tmp16992
_ = let__14441

tmp17002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14441)
}
__typedArg0 := let__14441
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres16993 Obj

if True == tmp17002 {
tmp16994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14441)
}
__typedArg0 := let__14441
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14442 := tmp16994
_ = let__14442

tmp16995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14441)
}
__typedArg0 := let__14441
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp16996 := Call(__e, PrimFunc(symshen_4lazyderef), tmp16995, V5581)


let__14443 := tmp16996
_ = let__14443

tmp17001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14443, Nil)
}
__typedArg0 := let__14443
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres16997 Obj

if True == tmp17001 {
tmp16998 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16998

tmp16999 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1correct), V5579, let__14442, Nil, V5581, V5582, let__14436, V5584)
return
}, 0)

tmp17000 := Call(__e, PrimFunc(symshen_4cut), V5581, V5582, let__14436, tmp16999)


ifres16997 = tmp17000


} else {
ifres16997 = False


}

ifres16993 = ifres16997


} else {
ifres16993 = False


}

ifres16990 = ifres16993


} else {
ifres16990 = False


}

ifres16987 = ifres16990


} else {
ifres16987 = False


}

ifres16985 = ifres16987


} else {
ifres16985 = False


}

ifres16983 = ifres16985


} else {
ifres16983 = False


}

let__14437 := ifres16983
_ = let__14437

tmp17025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14437, False)
}
__typedArg0 := let__14437
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17025 {
tmp17021 := Call(__e, PrimFunc(symshen_4unlocked_2), V5582)


var ifres17007 Obj

if True == tmp17021 {
tmp17008 := Call(__e, PrimFunc(symshen_4newpv), V5581)


let__14445 := tmp17008
_ = let__14445

tmp17009 := Call(__e, PrimFunc(symshen_4newpv), V5581)


let__14446 := tmp17009
_ = let__14446

tmp17010 := Call(__e, PrimFunc(symshen_4newpv), V5581)


let__14447 := tmp17010
_ = let__14447

tmp17011 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17011

tmp17012 := Call(__e, PrimFunc(symshen_4freshterms), V5578)


tmp17013 := MakeNative(func(__e *ControlFlow) {
tmp17014 := MakeNative(func(__e *ControlFlow) {
tmp17015 := MakeNative(func(__e *ControlFlow) {
tmp17016 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1correct), V5579, let__14446, let__14447, V5581, V5582, let__14436, V5584)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4myassume), V5578, V5580, let__14447, V5581, V5582, let__14436, tmp17016)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5581, V5582, let__14436, tmp17015)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1integrity), V5578, V5580, let__14445, let__14446, V5581, V5582, let__14436, tmp17014)
return


}, 0)

tmp17017 := Call(__e, PrimFunc(symshen_4p_1hyps), tmp17012, let__14445, V5581, V5582, let__14436, tmp17013)


tmp17018 := Call(__e, PrimFunc(symshen_4gc), V5581, tmp17017)


tmp17019 := Call(__e, PrimFunc(symshen_4gc), V5581, tmp17018)


tmp17020 := Call(__e, PrimFunc(symshen_4gc), V5581, tmp17019)


ifres17007 = tmp17020


} else {
ifres17007 = False


}

let__14444 := ifres17007
_ = let__14444

tmp17023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14444, False)
}
__typedArg0 := let__14444
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17023 {
__e.TailApply(PrimFunc(symshen_4unlock), V5582, let__14436)
return
} else {
__e.Return(let__14444)
return
}


} else {
__e.Return(let__14437)
return
}


}, 7)

tmp17026 := Call(__e, ns2_1set, symshen_4t_d_1rule_1h, tmp16981)


_ = tmp17026

tmp17027 := MakeNative(func(__e *ControlFlow) {
V5597 := __e.Get(1)
_ = V5597
V5598 := __e.Get(2)
_ = V5598
V5599 := __e.Get(3)
_ = V5599
V5600 := __e.Get(4)
_ = V5600
V5601 := __e.Get(5)
_ = V5601
V5602 := __e.Get(6)
_ = V5602
V5603 := __e.Get(7)
_ = V5603
tmp17041 := Call(__e, PrimFunc(symshen_4unlocked_2), V5601)


var ifres17028 Obj

if True == tmp17041 {
tmp17029 := Call(__e, PrimFunc(symshen_4lazyderef), V5597, V5600)


let__14449 := tmp17029
_ = let__14449

tmp17040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14449, Nil)
}
__typedArg0 := let__14449
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17030 Obj

if True == tmp17040 {
tmp17031 := Call(__e, PrimFunc(symshen_4lazyderef), V5599, V5600)


let__14450 := tmp17031
_ = let__14450

tmp17032 := MakeNative(func(__e *ControlFlow) {
tmp17033 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17033

__e.TailApply(PrimFunc(symthaw), V5603)
return


}, 0)

let__14451 := tmp17032
_ = let__14451

tmp17039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14450, Nil)
}
__typedArg0 := let__14450
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17034 Obj

if True == tmp17039 {
tmp17035 := Call(__e, PrimFunc(symthaw), let__14451)


ifres17034 = tmp17035


} else {
tmp17038 := Call(__e, PrimFunc(symshen_4pvar_2), let__14450)


var ifres17036 Obj

if True == tmp17038 {
tmp17037 := Call(__e, PrimFunc(symshen_4bind_b), let__14450, Nil, V5600, let__14451)


ifres17036 = tmp17037


} else {
ifres17036 = False


}

ifres17034 = ifres17036


}

ifres17030 = ifres17034


} else {
ifres17030 = False


}

ifres17028 = ifres17030


} else {
ifres17028 = False


}

let__14448 := ifres17028
_ = let__14448

tmp17160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14448, False)
}
__typedArg0 := let__14448
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17160 {
tmp17158 := Call(__e, PrimFunc(symshen_4unlocked_2), V5601)


if True == tmp17158 {
tmp17042 := Call(__e, PrimFunc(symshen_4lazyderef), V5597, V5600)


let__14452 := tmp17042
_ = let__14452

tmp17156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14452)
}
__typedArg0 := let__14452
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17156 {
tmp17043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14452)
}
__typedArg0 := let__14452
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14453 := tmp17043
_ = let__14453

tmp17044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14452)
}
__typedArg0 := let__14452
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14454 := tmp17044
_ = let__14454

tmp17045 := Call(__e, PrimFunc(symshen_4lazyderef), V5598, V5600)


let__14455 := tmp17045
_ = let__14455

tmp17154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14455)
}
__typedArg0 := let__14455
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17154 {
tmp17046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14455)
}
__typedArg0 := let__14455
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14456 := tmp17046
_ = let__14456

tmp17047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14455)
}
__typedArg0 := let__14455
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17048 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17047, V5600)


let__14457 := tmp17048
_ = let__14457

tmp17152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14457)
}
__typedArg0 := let__14457
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17152 {
tmp17049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14457)
}
__typedArg0 := let__14457
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17050 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17049, V5600)


let__14458 := tmp17050
_ = let__14458

tmp17150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14458, sym_1_1_6)
}
__typedArg0 := let__14458
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17150 {
tmp17051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14457)
}
__typedArg0 := let__14457
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17052 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17051, V5600)


let__14459 := tmp17052
_ = let__14459

tmp17148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14459)
}
__typedArg0 := let__14459
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17148 {
tmp17053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14459)
}
__typedArg0 := let__14459
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14460 := tmp17053
_ = let__14460

tmp17054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14459)
}
__typedArg0 := let__14459
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17055 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17054, V5600)


let__14461 := tmp17055
_ = let__14461

tmp17146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14461, Nil)
}
__typedArg0 := let__14461
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17146 {
tmp17056 := Call(__e, PrimFunc(symshen_4lazyderef), V5599, V5600)


let__14462 := tmp17056
_ = let__14462

tmp17057 := MakeNative(func(__e *ControlFlow) {
Z5620 := __e.Get(1)
_ = Z5620
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5621 := __e.Get(1)
_ = Z5621
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5622 := __e.Get(1)
_ = Z5622
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5623 := __e.Get(1)
_ = Z5623
tmp17058 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17058

tmp17059 := MakeNative(func(__e *ControlFlow) {
tmp17060 := MakeNative(func(__e *ControlFlow) {
tmp17061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp17062 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4myassume), let__14454, let__14460, Z5623, V5600, V5601, V5602, V5603)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5621, tmp17061, V5600, V5601, V5602, tmp17062)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), let__14453, Z5620, V5600, V5601, V5602, tmp17060)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), let__14456, Z5622, V5600, V5601, V5602, tmp17059)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

let__14463 := tmp17057
_ = let__14463

tmp17144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14462)
}
__typedArg0 := let__14462
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17144 {
tmp17063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14462)
}
__typedArg0 := let__14462
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17064 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17063, V5600)


let__14464 := tmp17064
_ = let__14464

tmp17065 := MakeNative(func(__e *ControlFlow) {
Z5626 := __e.Get(1)
_ = Z5626
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5627 := __e.Get(1)
_ = Z5627
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5628 := __e.Get(1)
_ = Z5628
tmp17066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14462)
}
__typedArg0 := let__14462
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14466 := tmp17066
_ = let__14466

tmp17067 := Call(__e, let__14463, Z5626)


tmp17068 := Call(__e, tmp17067, Z5627)


tmp17069 := Call(__e, tmp17068, Z5628)


__e.TailApply(tmp17069, let__14466)
return


}, 1))
return
}, 1))
return
}, 1)

let__14465 := tmp17065
_ = let__14465

tmp17124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14464)
}
__typedArg0 := let__14464
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17124 {
tmp17070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14464)
}
__typedArg0 := let__14464
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14467 := tmp17070
_ = let__14467

tmp17071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14464)
}
__typedArg0 := let__14464
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17072 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17071, V5600)


let__14468 := tmp17072
_ = let__14468

tmp17073 := MakeNative(func(__e *ControlFlow) {
Z5633 := __e.Get(1)
_ = Z5633
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5634 := __e.Get(1)
_ = Z5634
tmp17074 := Call(__e, let__14465, let__14467)


tmp17075 := Call(__e, tmp17074, Z5633)


__e.TailApply(tmp17075, Z5634)
return


}, 1))
return
}, 1)

let__14469 := tmp17073
_ = let__14469

tmp17108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14468)
}
__typedArg0 := let__14468
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17108 {
tmp17076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14468)
}
__typedArg0 := let__14468
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14470 := tmp17076
_ = let__14470

tmp17077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14468)
}
__typedArg0 := let__14468
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17078 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17077, V5600)


let__14471 := tmp17078
_ = let__14471

tmp17079 := MakeNative(func(__e *ControlFlow) {
Z5638 := __e.Get(1)
_ = Z5638
tmp17080 := Call(__e, let__14469, let__14470)


__e.TailApply(tmp17080, Z5638)
return


}, 1)

let__14472 := tmp17079
_ = let__14472

tmp17096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14471)
}
__typedArg0 := let__14471
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17096 {
tmp17081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14471)
}
__typedArg0 := let__14471
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14473 := tmp17081
_ = let__14473

tmp17082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14471)
}
__typedArg0 := let__14471
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17083 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17082, V5600)


let__14474 := tmp17083
_ = let__14474

tmp17084 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14472, let__14473)
return
}, 0)

let__14475 := tmp17084
_ = let__14475

tmp17088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14474, Nil)
}
__typedArg0 := let__14474
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17088 {
__e.TailApply(PrimFunc(symthaw), let__14475)
return
} else {
tmp17086 := Call(__e, PrimFunc(symshen_4pvar_2), let__14474)


if True == tmp17086 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14474, Nil, V5600, let__14475)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp17094 := Call(__e, PrimFunc(symshen_4pvar_2), let__14471)


if True == tmp17094 {
tmp17089 := Call(__e, PrimFunc(symshen_4newpv), V5600)


let__14476 := tmp17089
_ = let__14476

tmp17090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14476, Nil)
}
__typedArg0 := let__14476
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17091 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14472, let__14476)
return
}, 0)

tmp17092 := Call(__e, PrimFunc(symshen_4bind_b), let__14471, tmp17090, V5600, tmp17091)


__e.TailApply(PrimFunc(symshen_4gc), V5600, tmp17092)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17106 := Call(__e, PrimFunc(symshen_4pvar_2), let__14468)


if True == tmp17106 {
tmp17097 := Call(__e, PrimFunc(symshen_4newpv), V5600)


let__14477 := tmp17097
_ = let__14477

tmp17098 := Call(__e, PrimFunc(symshen_4newpv), V5600)


let__14478 := tmp17098
_ = let__14478

tmp17099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14478, Nil)
}
__typedArg0 := let__14478
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14477, tmp17099)
}
__typedArg0 := let__14477
__typedArg1 := tmp17099
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17101 := MakeNative(func(__e *ControlFlow) {
tmp17102 := Call(__e, let__14469, let__14477)


__e.TailApply(tmp17102, let__14478)
return


}, 0)

tmp17103 := Call(__e, PrimFunc(symshen_4bind_b), let__14468, tmp17100, V5600, tmp17101)


tmp17104 := Call(__e, PrimFunc(symshen_4gc), V5600, tmp17103)


__e.TailApply(PrimFunc(symshen_4gc), V5600, tmp17104)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17122 := Call(__e, PrimFunc(symshen_4pvar_2), let__14464)


if True == tmp17122 {
tmp17109 := Call(__e, PrimFunc(symshen_4newpv), V5600)


let__14479 := tmp17109
_ = let__14479

tmp17110 := Call(__e, PrimFunc(symshen_4newpv), V5600)


let__14480 := tmp17110
_ = let__14480

tmp17111 := Call(__e, PrimFunc(symshen_4newpv), V5600)


let__14481 := tmp17111
_ = let__14481

tmp17112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14481, Nil)
}
__typedArg0 := let__14481
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14480, tmp17112)
}
__typedArg0 := let__14480
__typedArg1 := tmp17112
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14479, tmp17113)
}
__typedArg0 := let__14479
__typedArg1 := tmp17113
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17115 := MakeNative(func(__e *ControlFlow) {
tmp17116 := Call(__e, let__14465, let__14479)


tmp17117 := Call(__e, tmp17116, let__14480)


__e.TailApply(tmp17117, let__14481)
return


}, 0)

tmp17118 := Call(__e, PrimFunc(symshen_4bind_b), let__14464, tmp17114, V5600, tmp17115)


tmp17119 := Call(__e, PrimFunc(symshen_4gc), V5600, tmp17118)


tmp17120 := Call(__e, PrimFunc(symshen_4gc), V5600, tmp17119)


__e.TailApply(PrimFunc(symshen_4gc), V5600, tmp17120)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17142 := Call(__e, PrimFunc(symshen_4pvar_2), let__14462)


if True == tmp17142 {
tmp17125 := Call(__e, PrimFunc(symshen_4newpv), V5600)


let__14482 := tmp17125
_ = let__14482

tmp17126 := Call(__e, PrimFunc(symshen_4newpv), V5600)


let__14483 := tmp17126
_ = let__14483

tmp17127 := Call(__e, PrimFunc(symshen_4newpv), V5600)


let__14484 := tmp17127
_ = let__14484

tmp17128 := Call(__e, PrimFunc(symshen_4newpv), V5600)


let__14485 := tmp17128
_ = let__14485

tmp17129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14484, Nil)
}
__typedArg0 := let__14484
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14483, tmp17129)
}
__typedArg0 := let__14483
__typedArg1 := tmp17129
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14482, tmp17130)
}
__typedArg0 := let__14482
__typedArg1 := tmp17130
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17131, let__14485)
}
__typedArg0 := tmp17131
__typedArg1 := let__14485
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17133 := MakeNative(func(__e *ControlFlow) {
tmp17134 := Call(__e, let__14463, let__14482)


tmp17135 := Call(__e, tmp17134, let__14483)


tmp17136 := Call(__e, tmp17135, let__14484)


__e.TailApply(tmp17136, let__14485)
return


}, 0)

tmp17137 := Call(__e, PrimFunc(symshen_4bind_b), let__14462, tmp17132, V5600, tmp17133)


tmp17138 := Call(__e, PrimFunc(symshen_4gc), V5600, tmp17137)


tmp17139 := Call(__e, PrimFunc(symshen_4gc), V5600, tmp17138)


tmp17140 := Call(__e, PrimFunc(symshen_4gc), V5600, tmp17139)


__e.TailApply(PrimFunc(symshen_4gc), V5600, tmp17140)
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
__e.Return(let__14448)
return
}


}, 7)

tmp17161 := Call(__e, ns2_1set, symshen_4myassume, tmp17027)


_ = tmp17161

tmp17162 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V5654 := __self1
_ = V5654
tmp17185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5654)
}
__typedArg0 := Nil
__typedArg1 := V5654
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17185 {
__e.Return(Nil)
return
} else {
tmp17183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5654)
}
__typedArg0 := V5654
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17179 Obj

if True == tmp17183 {
tmp17181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5654)
}
__typedArg0 := V5654
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp17181)
}
__typedArg0 := tmp17181
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17180 Obj

if True == tmp17182 {
ifres17180 = True


} else {
ifres17180 = False


}

ifres17179 = ifres17180


} else {
ifres17179 = False


}

if True == ifres17179 {
tmp17163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5654)
}
__typedArg0 := V5654
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5654)
}
__typedArg0 := V5654
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17165 := Call(__e, PrimFunc(symappend), tmp17163, tmp17164)


if PrimFunc(symshen_4freshterms) == __self {
__self1 = tmp17165
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4freshterms), tmp17165)
return


} else {
tmp17177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5654)
}
__typedArg0 := V5654
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17173 Obj

if True == tmp17177 {
tmp17175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5654)
}
__typedArg0 := V5654
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17176 := Call(__e, PrimFunc(symshen_4freshterm_2), tmp17175)


var ifres17174 Obj

if True == tmp17176 {
ifres17174 = True


} else {
ifres17174 = False


}

ifres17173 = ifres17174


} else {
ifres17173 = False


}

if True == ifres17173 {
tmp17166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5654)
}
__typedArg0 := V5654
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5654)
}
__typedArg0 := V5654
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17168 := Call(__e, PrimFunc(symshen_4freshterms), tmp17167)


__e.TailApply(PrimFunc(symadjoin), tmp17166, tmp17168)
return


} else {
tmp17171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5654)
}
__typedArg0 := V5654
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17171 {
tmp17169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5654)
}
__typedArg0 := V5654
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4freshterms) == __self {
__self1 = tmp17169
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4freshterms), tmp17169)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshterms)
return
}


}


}


}


}, 1)

tmp17186 := Call(__e, ns2_1set, symshen_4freshterms, tmp17162)


_ = tmp17186

tmp17187 := MakeNative(func(__e *ControlFlow) {
V5655 := __e.Get(1)
_ = V5655
V5656 := __e.Get(2)
_ = V5656
V5657 := __e.Get(3)
_ = V5657
V5658 := __e.Get(4)
_ = V5658
V5659 := __e.Get(5)
_ = V5659
V5660 := __e.Get(6)
_ = V5660
tmp17201 := Call(__e, PrimFunc(symshen_4unlocked_2), V5658)


var ifres17188 Obj

if True == tmp17201 {
tmp17189 := Call(__e, PrimFunc(symshen_4lazyderef), V5655, V5657)


let__14487 := tmp17189
_ = let__14487

tmp17200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14487, Nil)
}
__typedArg0 := let__14487
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17190 Obj

if True == tmp17200 {
tmp17191 := Call(__e, PrimFunc(symshen_4lazyderef), V5656, V5657)


let__14488 := tmp17191
_ = let__14488

tmp17192 := MakeNative(func(__e *ControlFlow) {
tmp17193 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17193

__e.TailApply(PrimFunc(symthaw), V5660)
return


}, 0)

let__14489 := tmp17192
_ = let__14489

tmp17199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14488, Nil)
}
__typedArg0 := let__14488
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17194 Obj

if True == tmp17199 {
tmp17195 := Call(__e, PrimFunc(symthaw), let__14489)


ifres17194 = tmp17195


} else {
tmp17198 := Call(__e, PrimFunc(symshen_4pvar_2), let__14488)


var ifres17196 Obj

if True == tmp17198 {
tmp17197 := Call(__e, PrimFunc(symshen_4bind_b), let__14488, Nil, V5657, let__14489)


ifres17196 = tmp17197


} else {
ifres17196 = False


}

ifres17194 = ifres17196


}

ifres17190 = ifres17194


} else {
ifres17190 = False


}

ifres17188 = ifres17190


} else {
ifres17188 = False


}

let__14486 := ifres17188
_ = let__14486

tmp17298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14486, False)
}
__typedArg0 := let__14486
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17298 {
tmp17296 := Call(__e, PrimFunc(symshen_4unlocked_2), V5658)


if True == tmp17296 {
tmp17202 := Call(__e, PrimFunc(symshen_4lazyderef), V5655, V5657)


let__14490 := tmp17202
_ = let__14490

tmp17294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14490)
}
__typedArg0 := let__14490
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17294 {
tmp17203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14490)
}
__typedArg0 := let__14490
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14491 := tmp17203
_ = let__14491

tmp17204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14490)
}
__typedArg0 := let__14490
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14492 := tmp17204
_ = let__14492

tmp17205 := Call(__e, PrimFunc(symshen_4lazyderef), V5656, V5657)


let__14493 := tmp17205
_ = let__14493

tmp17206 := MakeNative(func(__e *ControlFlow) {
Z5670 := __e.Get(1)
_ = Z5670
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5671 := __e.Get(1)
_ = Z5671
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5672 := __e.Get(1)
_ = Z5672
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5673 := __e.Get(1)
_ = Z5673
tmp17207 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17207

tmp17208 := MakeNative(func(__e *ControlFlow) {
tmp17209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp17210 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4p_1hyps), let__14492, Z5673, V5657, V5658, V5659, V5660)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5671, tmp17209, V5657, V5658, V5659, tmp17210)
return


}, 0)

__e.TailApply(PrimFunc(symbind), Z5670, let__14491, V5657, V5658, V5659, tmp17208)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

let__14494 := tmp17206
_ = let__14494

tmp17292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14493)
}
__typedArg0 := let__14493
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17292 {
tmp17211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14493)
}
__typedArg0 := let__14493
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17212 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17211, V5657)


let__14495 := tmp17212
_ = let__14495

tmp17213 := MakeNative(func(__e *ControlFlow) {
Z5676 := __e.Get(1)
_ = Z5676
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5677 := __e.Get(1)
_ = Z5677
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5678 := __e.Get(1)
_ = Z5678
tmp17214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14493)
}
__typedArg0 := let__14493
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14497 := tmp17214
_ = let__14497

tmp17215 := Call(__e, let__14494, Z5676)


tmp17216 := Call(__e, tmp17215, Z5677)


tmp17217 := Call(__e, tmp17216, Z5678)


__e.TailApply(tmp17217, let__14497)
return


}, 1))
return
}, 1))
return
}, 1)

let__14496 := tmp17213
_ = let__14496

tmp17272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14495)
}
__typedArg0 := let__14495
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17272 {
tmp17218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14495)
}
__typedArg0 := let__14495
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14498 := tmp17218
_ = let__14498

tmp17219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14495)
}
__typedArg0 := let__14495
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17220 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17219, V5657)


let__14499 := tmp17220
_ = let__14499

tmp17221 := MakeNative(func(__e *ControlFlow) {
Z5683 := __e.Get(1)
_ = Z5683
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5684 := __e.Get(1)
_ = Z5684
tmp17222 := Call(__e, let__14496, let__14498)


tmp17223 := Call(__e, tmp17222, Z5683)


__e.TailApply(tmp17223, Z5684)
return


}, 1))
return
}, 1)

let__14500 := tmp17221
_ = let__14500

tmp17256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14499)
}
__typedArg0 := let__14499
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17256 {
tmp17224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14499)
}
__typedArg0 := let__14499
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14501 := tmp17224
_ = let__14501

tmp17225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14499)
}
__typedArg0 := let__14499
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17226 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17225, V5657)


let__14502 := tmp17226
_ = let__14502

tmp17227 := MakeNative(func(__e *ControlFlow) {
Z5688 := __e.Get(1)
_ = Z5688
tmp17228 := Call(__e, let__14500, let__14501)


__e.TailApply(tmp17228, Z5688)
return


}, 1)

let__14503 := tmp17227
_ = let__14503

tmp17244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14502)
}
__typedArg0 := let__14502
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17244 {
tmp17229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14502)
}
__typedArg0 := let__14502
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14504 := tmp17229
_ = let__14504

tmp17230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14502)
}
__typedArg0 := let__14502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17231 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17230, V5657)


let__14505 := tmp17231
_ = let__14505

tmp17232 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14503, let__14504)
return
}, 0)

let__14506 := tmp17232
_ = let__14506

tmp17236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14505, Nil)
}
__typedArg0 := let__14505
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17236 {
__e.TailApply(PrimFunc(symthaw), let__14506)
return
} else {
tmp17234 := Call(__e, PrimFunc(symshen_4pvar_2), let__14505)


if True == tmp17234 {
__e.TailApply(PrimFunc(symshen_4bind_b), let__14505, Nil, V5657, let__14506)
return
} else {
__e.Return(False)
return
}


}


} else {
tmp17242 := Call(__e, PrimFunc(symshen_4pvar_2), let__14502)


if True == tmp17242 {
tmp17237 := Call(__e, PrimFunc(symshen_4newpv), V5657)


let__14507 := tmp17237
_ = let__14507

tmp17238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14507, Nil)
}
__typedArg0 := let__14507
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17239 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(let__14503, let__14507)
return
}, 0)

tmp17240 := Call(__e, PrimFunc(symshen_4bind_b), let__14502, tmp17238, V5657, tmp17239)


__e.TailApply(PrimFunc(symshen_4gc), V5657, tmp17240)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17254 := Call(__e, PrimFunc(symshen_4pvar_2), let__14499)


if True == tmp17254 {
tmp17245 := Call(__e, PrimFunc(symshen_4newpv), V5657)


let__14508 := tmp17245
_ = let__14508

tmp17246 := Call(__e, PrimFunc(symshen_4newpv), V5657)


let__14509 := tmp17246
_ = let__14509

tmp17247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14509, Nil)
}
__typedArg0 := let__14509
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14508, tmp17247)
}
__typedArg0 := let__14508
__typedArg1 := tmp17247
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17249 := MakeNative(func(__e *ControlFlow) {
tmp17250 := Call(__e, let__14500, let__14508)


__e.TailApply(tmp17250, let__14509)
return


}, 0)

tmp17251 := Call(__e, PrimFunc(symshen_4bind_b), let__14499, tmp17248, V5657, tmp17249)


tmp17252 := Call(__e, PrimFunc(symshen_4gc), V5657, tmp17251)


__e.TailApply(PrimFunc(symshen_4gc), V5657, tmp17252)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17270 := Call(__e, PrimFunc(symshen_4pvar_2), let__14495)


if True == tmp17270 {
tmp17257 := Call(__e, PrimFunc(symshen_4newpv), V5657)


let__14510 := tmp17257
_ = let__14510

tmp17258 := Call(__e, PrimFunc(symshen_4newpv), V5657)


let__14511 := tmp17258
_ = let__14511

tmp17259 := Call(__e, PrimFunc(symshen_4newpv), V5657)


let__14512 := tmp17259
_ = let__14512

tmp17260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14512, Nil)
}
__typedArg0 := let__14512
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14511, tmp17260)
}
__typedArg0 := let__14511
__typedArg1 := tmp17260
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14510, tmp17261)
}
__typedArg0 := let__14510
__typedArg1 := tmp17261
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17263 := MakeNative(func(__e *ControlFlow) {
tmp17264 := Call(__e, let__14496, let__14510)


tmp17265 := Call(__e, tmp17264, let__14511)


__e.TailApply(tmp17265, let__14512)
return


}, 0)

tmp17266 := Call(__e, PrimFunc(symshen_4bind_b), let__14495, tmp17262, V5657, tmp17263)


tmp17267 := Call(__e, PrimFunc(symshen_4gc), V5657, tmp17266)


tmp17268 := Call(__e, PrimFunc(symshen_4gc), V5657, tmp17267)


__e.TailApply(PrimFunc(symshen_4gc), V5657, tmp17268)
return


} else {
__e.Return(False)
return
}


}


} else {
tmp17290 := Call(__e, PrimFunc(symshen_4pvar_2), let__14493)


if True == tmp17290 {
tmp17273 := Call(__e, PrimFunc(symshen_4newpv), V5657)


let__14513 := tmp17273
_ = let__14513

tmp17274 := Call(__e, PrimFunc(symshen_4newpv), V5657)


let__14514 := tmp17274
_ = let__14514

tmp17275 := Call(__e, PrimFunc(symshen_4newpv), V5657)


let__14515 := tmp17275
_ = let__14515

tmp17276 := Call(__e, PrimFunc(symshen_4newpv), V5657)


let__14516 := tmp17276
_ = let__14516

tmp17277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14515, Nil)
}
__typedArg0 := let__14515
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14514, tmp17277)
}
__typedArg0 := let__14514
__typedArg1 := tmp17277
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14513, tmp17278)
}
__typedArg0 := let__14513
__typedArg1 := tmp17278
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17279, let__14516)
}
__typedArg0 := tmp17279
__typedArg1 := let__14516
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17281 := MakeNative(func(__e *ControlFlow) {
tmp17282 := Call(__e, let__14494, let__14513)


tmp17283 := Call(__e, tmp17282, let__14514)


tmp17284 := Call(__e, tmp17283, let__14515)


__e.TailApply(tmp17284, let__14516)
return


}, 0)

tmp17285 := Call(__e, PrimFunc(symshen_4bind_b), let__14493, tmp17280, V5657, tmp17281)


tmp17286 := Call(__e, PrimFunc(symshen_4gc), V5657, tmp17285)


tmp17287 := Call(__e, PrimFunc(symshen_4gc), V5657, tmp17286)


tmp17288 := Call(__e, PrimFunc(symshen_4gc), V5657, tmp17287)


__e.TailApply(PrimFunc(symshen_4gc), V5657, tmp17288)
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
__e.Return(let__14486)
return
}


}, 6)

tmp17299 := Call(__e, ns2_1set, symshen_4p_1hyps, tmp17187)


_ = tmp17299

tmp17300 := MakeNative(func(__e *ControlFlow) {
V5702 := __e.Get(1)
_ = V5702
V5703 := __e.Get(2)
_ = V5703
V5704 := __e.Get(3)
_ = V5704
V5705 := __e.Get(4)
_ = V5705
V5706 := __e.Get(5)
_ = V5706
V5707 := __e.Get(6)
_ = V5707
V5708 := __e.Get(7)
_ = V5708
let__14517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5707)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5707
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
_ = let__14517

tmp17338 := Call(__e, PrimFunc(symshen_4unlocked_2), V5706)


var ifres17302 Obj

if True == tmp17338 {
tmp17303 := Call(__e, PrimFunc(symshen_4lazyderef), V5702, V5705)


let__14519 := tmp17303
_ = let__14519

tmp17337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14519)
}
__typedArg0 := let__14519
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17304 Obj

if True == tmp17337 {
tmp17305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14519)
}
__typedArg0 := let__14519
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17306 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17305, V5705)


let__14520 := tmp17306
_ = let__14520

tmp17336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14520, symwhere)
}
__typedArg0 := let__14520
__typedArg1 := symwhere
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17307 Obj

if True == tmp17336 {
tmp17308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14519)
}
__typedArg0 := let__14519
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17309 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17308, V5705)


let__14521 := tmp17309
_ = let__14521

tmp17335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14521)
}
__typedArg0 := let__14521
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17310 Obj

if True == tmp17335 {
tmp17311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14521)
}
__typedArg0 := let__14521
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14522 := tmp17311
_ = let__14522

tmp17312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14521)
}
__typedArg0 := let__14521
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17313 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17312, V5705)


let__14523 := tmp17313
_ = let__14523

tmp17334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14523)
}
__typedArg0 := let__14523
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres17314 Obj

if True == tmp17334 {
tmp17315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14523)
}
__typedArg0 := let__14523
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14524 := tmp17315
_ = let__14524

tmp17316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14523)
}
__typedArg0 := let__14523
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17317 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17316, V5705)


let__14525 := tmp17317
_ = let__14525

tmp17333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14525, Nil)
}
__typedArg0 := let__14525
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17318 Obj

if True == tmp17333 {
tmp17319 := Call(__e, PrimFunc(symshen_4newpv), V5705)


let__14526 := tmp17319
_ = let__14526

tmp17320 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17320

tmp17321 := MakeNative(func(__e *ControlFlow) {
tmp17322 := Call(__e, PrimFunc(symshen_4curry), let__14522)


tmp17323 := MakeNative(func(__e *ControlFlow) {
tmp17324 := MakeNative(func(__e *ControlFlow) {
tmp17325 := MakeNative(func(__e *ControlFlow) {
tmp17326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp17327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symverified, Nil)
}
__typedArg0 := symverified
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17326, tmp17327)
}
__typedArg0 := tmp17326
__typedArg1 := tmp17327
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__14526, tmp17328)
}
__typedArg0 := let__14526
__typedArg1 := tmp17328
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp17330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp17329, V5704)
}
__typedArg0 := tmp17329
__typedArg1 := V5704
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4t_d_1correct), let__14524, V5703, tmp17330, V5705, V5706, let__14517, V5708)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5705, V5706, let__14517, tmp17325)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14526, symboolean, V5704, V5705, V5706, let__14517, tmp17324)
return


}, 0)

__e.TailApply(PrimFunc(symbind), let__14526, tmp17322, V5705, V5706, let__14517, tmp17323)
return


}, 0)

tmp17331 := Call(__e, PrimFunc(symshen_4cut), V5705, V5706, let__14517, tmp17321)


tmp17332 := Call(__e, PrimFunc(symshen_4gc), V5705, tmp17331)


ifres17318 = tmp17332


} else {
ifres17318 = False


}

ifres17314 = ifres17318


} else {
ifres17314 = False


}

ifres17310 = ifres17314


} else {
ifres17310 = False


}

ifres17307 = ifres17310


} else {
ifres17307 = False


}

ifres17304 = ifres17307


} else {
ifres17304 = False


}

ifres17302 = ifres17304


} else {
ifres17302 = False


}

let__14518 := ifres17302
_ = let__14518

tmp17347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14518, False)
}
__typedArg0 := let__14518
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17347 {
tmp17343 := Call(__e, PrimFunc(symshen_4unlocked_2), V5706)


var ifres17339 Obj

if True == tmp17343 {
tmp17340 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17340

tmp17341 := Call(__e, PrimFunc(symshen_4curry), V5702)


tmp17342 := Call(__e, PrimFunc(symshen_4system_1S_1h), tmp17341, V5703, V5704, V5705, V5706, let__14517, V5708)


ifres17339 = tmp17342


} else {
ifres17339 = False


}

let__14527 := ifres17339
_ = let__14527

tmp17345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14527, False)
}
__typedArg0 := let__14527
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17345 {
__e.TailApply(PrimFunc(symshen_4unlock), V5706, let__14517)
return
} else {
__e.Return(let__14527)
return
}


} else {
__e.Return(let__14518)
return
}


}, 7)

tmp17348 := Call(__e, ns2_1set, symshen_4t_d_1correct, tmp17300)


_ = tmp17348

tmp17349 := MakeNative(func(__e *ControlFlow) {
V5720 := __e.Get(1)
_ = V5720
V5721 := __e.Get(2)
_ = V5721
V5722 := __e.Get(3)
_ = V5722
V5723 := __e.Get(4)
_ = V5723
V5724 := __e.Get(5)
_ = V5724
V5725 := __e.Get(6)
_ = V5725
V5726 := __e.Get(7)
_ = V5726
V5727 := __e.Get(8)
_ = V5727
tmp17356 := Call(__e, PrimFunc(symshen_4unlocked_2), V5725)


var ifres17350 Obj

if True == tmp17356 {
tmp17351 := Call(__e, PrimFunc(symshen_4lazyderef), V5720, V5724)


let__14529 := tmp17351
_ = let__14529

tmp17355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14529, Nil)
}
__typedArg0 := let__14529
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17352 Obj

if True == tmp17355 {
tmp17353 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17353

tmp17354 := Call(__e, PrimFunc(symis_b), V5721, V5723, V5724, V5725, V5726, V5727)


ifres17352 = tmp17354


} else {
ifres17352 = False


}

ifres17350 = ifres17352


} else {
ifres17350 = False


}

let__14528 := ifres17350
_ = let__14528

tmp17388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14528, False)
}
__typedArg0 := let__14528
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17388 {
tmp17386 := Call(__e, PrimFunc(symshen_4unlocked_2), V5725)


if True == tmp17386 {
tmp17357 := Call(__e, PrimFunc(symshen_4lazyderef), V5720, V5724)


let__14530 := tmp17357
_ = let__14530

tmp17384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14530)
}
__typedArg0 := let__14530
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17384 {
tmp17358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14530)
}
__typedArg0 := let__14530
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14531 := tmp17358
_ = let__14531

tmp17359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14530)
}
__typedArg0 := let__14530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__14532 := tmp17359
_ = let__14532

tmp17360 := Call(__e, PrimFunc(symshen_4lazyderef), V5721, V5724)


let__14533 := tmp17360
_ = let__14533

tmp17382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14533)
}
__typedArg0 := let__14533
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17382 {
tmp17361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14533)
}
__typedArg0 := let__14533
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14534 := tmp17361
_ = let__14534

tmp17362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14533)
}
__typedArg0 := let__14533
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17363 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17362, V5724)


let__14535 := tmp17363
_ = let__14535

tmp17380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14535)
}
__typedArg0 := let__14535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17380 {
tmp17364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14535)
}
__typedArg0 := let__14535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp17365 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17364, V5724)


let__14536 := tmp17365
_ = let__14536

tmp17378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14536, sym_1_1_6)
}
__typedArg0 := let__14536
__typedArg1 := sym_1_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17378 {
tmp17366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14535)
}
__typedArg0 := let__14535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17367 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17366, V5724)


let__14537 := tmp17367
_ = let__14537

tmp17376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__14537)
}
__typedArg0 := let__14537
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp17376 {
tmp17368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__14537)
}
__typedArg0 := let__14537
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__14538 := tmp17368
_ = let__14538

tmp17369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__14537)
}
__typedArg0 := let__14537
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp17370 := Call(__e, PrimFunc(symshen_4lazyderef), tmp17369, V5724)


let__14539 := tmp17370
_ = let__14539

tmp17374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__14539, Nil)
}
__typedArg0 := let__14539
__typedArg1 := Nil
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp17374 {
tmp17371 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17371

tmp17372 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1integrity), let__14532, let__14538, V5722, V5723, V5724, V5725, V5726, V5727)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), let__14531, let__14534, V5722, V5724, V5725, V5726, tmp17372)
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
__e.Return(let__14528)
return
}


}, 8)

tmp17389 := Call(__e, ns2_1set, symshen_4t_d_1integrity, tmp17349)


_ = tmp17389

tmp17390 := MakeNative(func(__e *ControlFlow) {
V5740 := __e.Get(1)
_ = V5740
tmp17399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V5740)
}
__typedArg0 := V5740
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp17399 {
tmp17396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V5740)
}
__typedArg0 := V5740
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

tmp17397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp17396)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp17396
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres17392 Obj

if True == tmp17397 {
tmp17394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V5740, MakeInteger(0))
}
__typedArg0 := V5740
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp17395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp17394, symshen_4print_1freshterm)
}
__typedArg0 := tmp17394
__typedArg1 := symshen_4print_1freshterm
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres17393 Obj

if True == tmp17395 {
ifres17393 = True


} else {
ifres17393 = False


}

ifres17392 = ifres17393


} else {
ifres17392 = False


}

if True == ifres17392 {
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

__e.TailApply(ns2_1set, symshen_4freshterm_2, tmp17390)
return




}, 0)

