package main

import . "github.com/pyrex41/shen-go/kl"

var DeclarationsMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp7642 := MakeNative0(func(__e *ControlFlow) Obj {
tmp7643 := MakeNative1(func(__e *ControlFlow, Z920 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4typename), Z920)
panic("unreachable")
})

tmp7644 := GValue(__e, symshen_4_dalldatatypes_d)

return TailCall2(__e, SymFn(symmap), tmp7643, tmp7644)


panic("unreachable")
})

tmp7645 := Call2(__e, ns2_1set, symdatatypes, tmp7642)


_ = tmp7645

tmp7646 := MakeNative0(func(__e *ControlFlow) Obj {
tmp7647 := MakeNative1(func(__e *ControlFlow, Z921 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4typename), Z921)
panic("unreachable")
})

tmp7648 := GValue(__e, symshen_4_ddatatypes_d)

return TailCall2(__e, SymFn(symmap), tmp7647, tmp7648)


panic("unreachable")
})

tmp7649 := Call2(__e, ns2_1set, symincluded, tmp7646)


_ = tmp7649

tmp7650 := MakeNative1(func(__e *ControlFlow, V924 Obj) Obj {
tmp7655 := GConsP(__e, V924)

if True == tmp7655 {
tmp7651 := GHd(__e, V924)

tmp7652 := PrimGuard1(__e, symstr, PrimStr, tmp7651)

tmp7653 := Call1(__e, SymFn(symshen_4typename_1h), tmp7652)


return PrimGuard1(__e, symintern, PrimIntern, tmp7653)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4typename)
}


panic("unreachable")
})

tmp7656 := Call2(__e, ns2_1set, symshen_4typename, tmp7650)


_ = tmp7656

tmp7657 := MakeNative1(func(__e *ControlFlow, V925 Obj) Obj {
tmp7664 := GEq(__e, MakeString("#type"), V925)

if True == tmp7664 {
return MakeString("")
} else {
tmp7662 := Call1(__e, SymFn(symshen_4_7string_2), V925)


if True == tmp7662 {
tmp7658 := Call1(__e, SymFn(symhdstr), V925)


tmp7660 := Call1(__e, SymFn(symshen_4typename_1h), PrimGuard1(__e, symtlstr, PrimTailString, V925))


return PrimGuard2(__e, symcn, PrimStringConcat, tmp7658, tmp7660)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4typename_1h)
}


}


panic("unreachable")
})

tmp7665 := Call2(__e, ns2_1set, symshen_4typename_1h, tmp7657)


_ = tmp7665

tmp7666 := MakeNative1(func(__e *ControlFlow, V926 Obj) Obj {
if True == GLt(__e, V926, MakeInteger(0)) {
return GValue(__e, symshen_4_dprolog_1memory_d)
} else {
tmp7668 := PrimGuard1(__e, syminteger_2, PrimIsInteger, V926)

if True == tmp7668 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dprolog_1memory_d, V926)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("prolog memory expects an integer value\n"))
}


}


panic("unreachable")
})

tmp7671 := Call2(__e, ns2_1set, symprolog_1memory, tmp7666)


_ = tmp7671

tmp7672 := MakeNative1(func(__e *ControlFlow, V927 Obj) Obj {
tmp7673 := MakeNative0(func(__e *ControlFlow) Obj {
tmp7674 := GValue(__e, sym_dproperty_1vector_d)

return TailCall3(__e, SymFn(symget), V927, symarity, tmp7674)


panic("unreachable")
})

tmp7675 := MakeNative1(func(__e *ControlFlow, Z928 Obj) Obj {
return MakeInteger(-1)
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp7673, tmp7675)


panic("unreachable")
})

tmp7676 := Call2(__e, ns2_1set, symarity, tmp7672)


_ = tmp7676

tmp7677 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V931 := __self1
_ = V931
tmp7691 := GEq(__e, Nil, V931)

if True == tmp7691 {
return Nil
} else {
tmp7689 := GConsP(__e, V931)

var ifres7686 Obj

if True == tmp7689 {
tmp7687 := GTl(__e, V931)

tmp7688 := GConsP(__e, tmp7687)

ifres7686 = tmp7688


} else {
ifres7686 = False


}

if True == ifres7686 {
tmp7678 := GHd(__e, V931)

tmp7679 := GTl(__e, V931)

tmp7680 := GHd(__e, tmp7679)

tmp7681 := GValue(__e, sym_dproperty_1vector_d)

tmp7682 := Call4(__e, SymFn(symput), tmp7678, symarity, tmp7680, tmp7681)


let__7637 := tmp7682
_ = let__7637

tmp7683 := GTl(__e, V931)

tmp7684 := GTl(__e, tmp7683)

if SymFn(symshen_4initialise_1arity_1table) == __self {
__self1 = tmp7684
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4initialise_1arity_1table), tmp7684)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.initialise-arity-table"))
}


}


panic("unreachable")
})
return __self
})()

tmp7692 := Call2(__e, ns2_1set, symshen_4initialise_1arity_1table, tmp7677)


_ = tmp7692

tmp7693 := MakeNative1(func(__e *ControlFlow, V933 Obj) Obj {
tmp7694 := GValue(__e, sym_dproperty_1vector_d)

tmp7695 := Call3(__e, SymFn(symget), symshen, symshen_4external_1symbols, tmp7694)


let__7638 := tmp7695
_ = let__7638

tmp7696 := Call2(__e, SymFn(symadjoin), V933, let__7638)


tmp7697 := GValue(__e, sym_dproperty_1vector_d)

tmp7698 := Call4(__e, SymFn(symput), symshen, symshen_4external_1symbols, tmp7696, tmp7697)


let__7639 := tmp7698
_ = let__7639

return V933


panic("unreachable")
})

tmp7699 := Call2(__e, ns2_1set, symsystemf, tmp7693)


_ = tmp7699

tmp7700 := MakeNative2(func(__e *ControlFlow, V936, V937 Obj) Obj {
tmp7702 := Call2(__e, SymFn(symelement_2), V936, V937)


if True == tmp7702 {
return V937
} else {
return GCons(__e, V936, V937)
}


panic("unreachable")
})

tmp7703 := Call2(__e, ns2_1set, symadjoin, tmp7700)


_ = tmp7703

tmp7704 := MakeNative1(func(__e *ControlFlow, V938 Obj) Obj {
tmp7705 := Call1(__e, SymFn(symarity), V938)


let__7640 := tmp7705
_ = let__7640

tmp7712 := GEq(__e, let__7640, MakeInteger(-1))

var ifres7710 Obj

if True == tmp7712 {
ifres7710 = True


} else {
tmp7711 := GEq(__e, let__7640, MakeInteger(0))

ifres7710 = tmp7711


}

if True == ifres7710 {
return Nil
} else {
tmp7706 := GCons(__e, V938, Nil)

tmp7707 := Call2(__e, SymFn(symshen_4lambda_1function), tmp7706, let__7640)


tmp7708 := Call1(__e, SymFn(symeval_1kl), tmp7707)


return GCons(__e, V938, tmp7708)


}


panic("unreachable")
})

tmp7713 := Call2(__e, ns2_1set, symshen_4lambda_1entry, tmp7704)


_ = tmp7713

tmp7714 := MakeNative1(func(__e *ControlFlow, V940 Obj) Obj {
tmp7719 := GConsP(__e, V940)

if True == tmp7719 {
tmp7715 := GHd(__e, V940)

tmp7716 := GTl(__e, V940)

tmp7717 := GValue(__e, sym_dproperty_1vector_d)

return TailCall4(__e, SymFn(symput), tmp7715, symshen_4lambda_1form, tmp7716, tmp7717)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4set_1lambda_1form_1entry)
}


panic("unreachable")
})

tmp7720 := Call2(__e, ns2_1set, symshen_4set_1lambda_1form_1entry, tmp7714)


_ = tmp7720

tmp7721 := MakeNative1(func(__e *ControlFlow, V941 Obj) Obj {
tmp7722 := MakeNative1(func(__e *ControlFlow, Z943 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4lambda_1entry), Z943)
panic("unreachable")
})

tmp7723 := Call2(__e, SymFn(symmap), tmp7722, V941)


let__7641 := tmp7723
_ = let__7641

tmp7724 := MakeNative1(func(__e *ControlFlow, Z944 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4set_1lambda_1form_1entry), Z944)
panic("unreachable")
})

tmp7725 := MakeNative1(func(__e *ControlFlow, Z945 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z946 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4register_1source_1form), Z945, Z946)
panic("unreachable")
})
panic("unreachable")
})

tmp7726 := GCons(__e, symshen_4register_1source_1form, tmp7725)

tmp7727 := MakeNative1(func(__e *ControlFlow, Z947 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4unregister_1source_1form), Z947)
panic("unreachable")
})

tmp7728 := GCons(__e, symshen_4unregister_1source_1form, tmp7727)

tmp7729 := MakeNative1(func(__e *ControlFlow, Z948 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4tuple), Z948)
panic("unreachable")
})

tmp7730 := GCons(__e, symshen_4tuple, tmp7729)

tmp7731 := MakeNative1(func(__e *ControlFlow, Z949 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4pvar), Z949)
panic("unreachable")
})

tmp7732 := GCons(__e, symshen_4pvar, tmp7731)

tmp7733 := MakeNative1(func(__e *ControlFlow, Z950 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4dictionary), Z950)
panic("unreachable")
})

tmp7734 := GCons(__e, symshen_4dictionary, tmp7733)

tmp7735 := MakeNative1(func(__e *ControlFlow, Z951 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4print_1prolog_1vector), Z951)
panic("unreachable")
})

tmp7736 := GCons(__e, symshen_4print_1prolog_1vector, tmp7735)

tmp7737 := MakeNative1(func(__e *ControlFlow, Z952 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4print_1freshterm), Z952)
panic("unreachable")
})

tmp7738 := GCons(__e, symshen_4print_1freshterm, tmp7737)

tmp7739 := MakeNative1(func(__e *ControlFlow, Z953 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4printF), Z953)
panic("unreachable")
})

tmp7740 := GCons(__e, symshen_4printF, tmp7739)

tmp7741 := GCons(__e, tmp7740, let__7641)

tmp7742 := GCons(__e, tmp7738, tmp7741)

tmp7743 := GCons(__e, tmp7736, tmp7742)

tmp7744 := GCons(__e, tmp7734, tmp7743)

tmp7745 := GCons(__e, tmp7732, tmp7744)

tmp7746 := GCons(__e, tmp7730, tmp7745)

tmp7747 := GCons(__e, tmp7728, tmp7746)

tmp7748 := GCons(__e, tmp7726, tmp7747)

return TailCall2(__e, SymFn(symshen_4for_1each), tmp7724, tmp7748)


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4build_1lambda_1table, tmp7721)




panic("unreachable")
})

