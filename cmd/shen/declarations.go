package main

import . "github.com/pyrex41/shen-go/kl"

var DeclarationsMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp7622 := MakeNative0(func(__e *ControlFlow) Obj {
tmp7623 := MakeNative1(func(__e *ControlFlow, Z920 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4typename), Z920)
panic("unreachable")
})

tmp7624 := GValue(__e, symshen_4_dalldatatypes_d)

return TailCall2(__e, SymFn(symmap), tmp7623, tmp7624)


panic("unreachable")
})

tmp7625 := Call2(__e, ns2_1set, symdatatypes, tmp7622)


_ = tmp7625

tmp7626 := MakeNative0(func(__e *ControlFlow) Obj {
tmp7627 := MakeNative1(func(__e *ControlFlow, Z921 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4typename), Z921)
panic("unreachable")
})

tmp7628 := GValue(__e, symshen_4_ddatatypes_d)

return TailCall2(__e, SymFn(symmap), tmp7627, tmp7628)


panic("unreachable")
})

tmp7629 := Call2(__e, ns2_1set, symincluded, tmp7626)


_ = tmp7629

tmp7630 := MakeNative1(func(__e *ControlFlow, V924 Obj) Obj {
tmp7635 := GConsP(__e, V924)

if True == tmp7635 {
tmp7631 := GHd(__e, V924)

tmp7632 := PrimGuard1(__e, symstr, PrimStr, tmp7631)

tmp7633 := Call1(__e, SymFn(symshen_4typename_1h), tmp7632)


return PrimGuard1(__e, symintern, PrimIntern, tmp7633)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4typename)
}


panic("unreachable")
})

tmp7636 := Call2(__e, ns2_1set, symshen_4typename, tmp7630)


_ = tmp7636

tmp7637 := MakeNative1(func(__e *ControlFlow, V925 Obj) Obj {
tmp7644 := GEq(__e, MakeString("#type"), V925)

if True == tmp7644 {
return MakeString("")
} else {
tmp7642 := Call1(__e, SymFn(symshen_4_7string_2), V925)


if True == tmp7642 {
tmp7638 := Call1(__e, SymFn(symhdstr), V925)


tmp7640 := Call1(__e, SymFn(symshen_4typename_1h), PrimGuard1(__e, symtlstr, PrimTailString, V925))


return PrimGuard2(__e, symcn, PrimStringConcat, tmp7638, tmp7640)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4typename_1h)
}


}


panic("unreachable")
})

tmp7645 := Call2(__e, ns2_1set, symshen_4typename_1h, tmp7637)


_ = tmp7645

tmp7646 := MakeNative1(func(__e *ControlFlow, V926 Obj) Obj {
if True == GLt(__e, V926, MakeInteger(0)) {
return GValue(__e, symshen_4_dprolog_1memory_d)
} else {
tmp7648 := PrimGuard1(__e, syminteger_2, PrimIsInteger, V926)

if True == tmp7648 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dprolog_1memory_d, V926)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("prolog memory expects an integer value\n"))
}


}


panic("unreachable")
})

tmp7651 := Call2(__e, ns2_1set, symprolog_1memory, tmp7646)


_ = tmp7651

tmp7652 := MakeNative1(func(__e *ControlFlow, V927 Obj) Obj {
tmp7653 := MakeNative0(func(__e *ControlFlow) Obj {
tmp7654 := GValue(__e, sym_dproperty_1vector_d)

return TailCall3(__e, SymFn(symget), V927, symarity, tmp7654)


panic("unreachable")
})

tmp7655 := MakeNative1(func(__e *ControlFlow, Z928 Obj) Obj {
return MakeInteger(-1)
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp7653, tmp7655)


panic("unreachable")
})

tmp7656 := Call2(__e, ns2_1set, symarity, tmp7652)


_ = tmp7656

tmp7657 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V931 := __self1
_ = V931
tmp7671 := GEq(__e, Nil, V931)

if True == tmp7671 {
return Nil
} else {
tmp7669 := GConsP(__e, V931)

var ifres7666 Obj

if True == tmp7669 {
tmp7667 := GTl(__e, V931)

tmp7668 := GConsP(__e, tmp7667)

ifres7666 = tmp7668


} else {
ifres7666 = False


}

if True == ifres7666 {
tmp7658 := GHd(__e, V931)

tmp7659 := GTl(__e, V931)

tmp7660 := GHd(__e, tmp7659)

tmp7661 := GValue(__e, sym_dproperty_1vector_d)

tmp7662 := Call4(__e, SymFn(symput), tmp7658, symarity, tmp7660, tmp7661)


let__7617 := tmp7662
_ = let__7617

tmp7663 := GTl(__e, V931)

tmp7664 := GTl(__e, tmp7663)

if SymFn(symshen_4initialise_1arity_1table) == __self {
__self1 = tmp7664
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4initialise_1arity_1table), tmp7664)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.initialise-arity-table"))
}


}


panic("unreachable")
})
return __self
})()

tmp7672 := Call2(__e, ns2_1set, symshen_4initialise_1arity_1table, tmp7657)


_ = tmp7672

tmp7673 := MakeNative1(func(__e *ControlFlow, V933 Obj) Obj {
tmp7674 := GValue(__e, sym_dproperty_1vector_d)

tmp7675 := Call3(__e, SymFn(symget), symshen, symshen_4external_1symbols, tmp7674)


let__7618 := tmp7675
_ = let__7618

tmp7676 := Call2(__e, SymFn(symadjoin), V933, let__7618)


tmp7677 := GValue(__e, sym_dproperty_1vector_d)

tmp7678 := Call4(__e, SymFn(symput), symshen, symshen_4external_1symbols, tmp7676, tmp7677)


let__7619 := tmp7678
_ = let__7619

return V933


panic("unreachable")
})

tmp7679 := Call2(__e, ns2_1set, symsystemf, tmp7673)


_ = tmp7679

tmp7680 := MakeNative2(func(__e *ControlFlow, V936, V937 Obj) Obj {
tmp7682 := Call2(__e, SymFn(symelement_2), V936, V937)


if True == tmp7682 {
return V937
} else {
return GCons(__e, V936, V937)
}


panic("unreachable")
})

tmp7683 := Call2(__e, ns2_1set, symadjoin, tmp7680)


_ = tmp7683

tmp7684 := MakeNative1(func(__e *ControlFlow, V938 Obj) Obj {
tmp7685 := Call1(__e, SymFn(symarity), V938)


let__7620 := tmp7685
_ = let__7620

tmp7692 := GEq(__e, let__7620, MakeInteger(-1))

var ifres7690 Obj

if True == tmp7692 {
ifres7690 = True


} else {
tmp7691 := GEq(__e, let__7620, MakeInteger(0))

ifres7690 = tmp7691


}

if True == ifres7690 {
return Nil
} else {
tmp7686 := GCons(__e, V938, Nil)

tmp7687 := Call2(__e, SymFn(symshen_4lambda_1function), tmp7686, let__7620)


tmp7688 := Call1(__e, SymFn(symeval_1kl), tmp7687)


return GCons(__e, V938, tmp7688)


}


panic("unreachable")
})

tmp7693 := Call2(__e, ns2_1set, symshen_4lambda_1entry, tmp7684)


_ = tmp7693

tmp7694 := MakeNative1(func(__e *ControlFlow, V940 Obj) Obj {
tmp7699 := GConsP(__e, V940)

if True == tmp7699 {
tmp7695 := GHd(__e, V940)

tmp7696 := GTl(__e, V940)

tmp7697 := GValue(__e, sym_dproperty_1vector_d)

return TailCall4(__e, SymFn(symput), tmp7695, symshen_4lambda_1form, tmp7696, tmp7697)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4set_1lambda_1form_1entry)
}


panic("unreachable")
})

tmp7700 := Call2(__e, ns2_1set, symshen_4set_1lambda_1form_1entry, tmp7694)


_ = tmp7700

tmp7701 := MakeNative1(func(__e *ControlFlow, V941 Obj) Obj {
tmp7702 := MakeNative1(func(__e *ControlFlow, Z943 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4lambda_1entry), Z943)
panic("unreachable")
})

tmp7703 := Call2(__e, SymFn(symmap), tmp7702, V941)


let__7621 := tmp7703
_ = let__7621

tmp7704 := MakeNative1(func(__e *ControlFlow, Z944 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4set_1lambda_1form_1entry), Z944)
panic("unreachable")
})

tmp7705 := MakeNative1(func(__e *ControlFlow, Z945 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z946 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4register_1source_1form), Z945, Z946)
panic("unreachable")
})
panic("unreachable")
})

tmp7706 := GCons(__e, symshen_4register_1source_1form, tmp7705)

tmp7707 := MakeNative1(func(__e *ControlFlow, Z947 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4unregister_1source_1form), Z947)
panic("unreachable")
})

tmp7708 := GCons(__e, symshen_4unregister_1source_1form, tmp7707)

tmp7709 := MakeNative1(func(__e *ControlFlow, Z948 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4tuple), Z948)
panic("unreachable")
})

tmp7710 := GCons(__e, symshen_4tuple, tmp7709)

tmp7711 := MakeNative1(func(__e *ControlFlow, Z949 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4pvar), Z949)
panic("unreachable")
})

tmp7712 := GCons(__e, symshen_4pvar, tmp7711)

tmp7713 := MakeNative1(func(__e *ControlFlow, Z950 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4dictionary), Z950)
panic("unreachable")
})

tmp7714 := GCons(__e, symshen_4dictionary, tmp7713)

tmp7715 := MakeNative1(func(__e *ControlFlow, Z951 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4print_1prolog_1vector), Z951)
panic("unreachable")
})

tmp7716 := GCons(__e, symshen_4print_1prolog_1vector, tmp7715)

tmp7717 := MakeNative1(func(__e *ControlFlow, Z952 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4print_1freshterm), Z952)
panic("unreachable")
})

tmp7718 := GCons(__e, symshen_4print_1freshterm, tmp7717)

tmp7719 := MakeNative1(func(__e *ControlFlow, Z953 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4printF), Z953)
panic("unreachable")
})

tmp7720 := GCons(__e, symshen_4printF, tmp7719)

tmp7721 := GCons(__e, tmp7720, let__7621)

tmp7722 := GCons(__e, tmp7718, tmp7721)

tmp7723 := GCons(__e, tmp7716, tmp7722)

tmp7724 := GCons(__e, tmp7714, tmp7723)

tmp7725 := GCons(__e, tmp7712, tmp7724)

tmp7726 := GCons(__e, tmp7710, tmp7725)

tmp7727 := GCons(__e, tmp7708, tmp7726)

tmp7728 := GCons(__e, tmp7706, tmp7727)

return TailCall2(__e, SymFn(symshen_4for_1each), tmp7704, tmp7728)


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4build_1lambda_1table, tmp7701)




panic("unreachable")
})

