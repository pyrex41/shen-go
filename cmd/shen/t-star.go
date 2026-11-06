package main

import . "github.com/pyrex41/shen-go/kl"

var TStarMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp14520 := MakeNative2(func(__e *ControlFlow, V4915, V4916 Obj) Obj {
tmp14521 := Call1(__e, SymFn(symshen_4extract_1vars), V4916)


let__13969 := tmp14521
_ = let__13969

tmp14522 := Call1(__e, SymFn(symshen_4rectify_1type), V4916)


let__13970 := tmp14522
_ = let__13970

tmp14523 := Call1(__e, SymFn(symshen_4curry), V4915)


let__13971 := tmp14523
_ = let__13971

tmp14524 := MakeNative1(func(__e *ControlFlow, Z4920 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z4921 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z4922 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z4923 Obj) Obj {
tmp14525 := Call1(__e, SymFn(symshen_4newpv), Z4920)


let__13972 := tmp14525
_ = let__13972

tmp14526 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14526

tmp14527 := Call2(__e, SymFn(symshen_4deref), let__13969, Z4920)


tmp14528 := Call1(__e, SymFn(symreceive), tmp14527)


tmp14529 := Call2(__e, SymFn(symshen_4deref), let__13970, Z4920)


tmp14530 := Call1(__e, SymFn(symreceive), tmp14529)


tmp14531 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14532 := Call2(__e, SymFn(symshen_4deref), let__13971, Z4920)


tmp14533 := Call1(__e, SymFn(symreceive), tmp14532)


tmp14534 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall5(__e, SymFn(symreturn), let__13972, Z4920, Z4921, Z4922, Z4923)
panic("unreachable")
})

return TailCall6(__e, SymFn(symshen_4toplevel_1forms), tmp14533, let__13972, Z4920, Z4921, Z4922, tmp14534)


panic("unreachable")
})

tmp14535 := Call(__e, SymFn(symshen_4insert_1prolog_1variables), tmp14528, tmp14530, let__13972, Z4920, Z4921, Z4922, tmp14531)


return TailCall2(__e, SymFn(symshen_4gc), Z4920, tmp14535)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp14536 := Call0(__e, SymFn(symshen_4prolog_1vector))


tmp14537 := Call1(__e, tmp14524, tmp14536)


tmp14538 := Call1(__e, SymFn(symvector), MakeInteger(0))


tmp14539 := Call2(__e, SymFn(sym_8v), MakeInteger(0), tmp14538)


tmp14540 := Call2(__e, SymFn(sym_8v), True, tmp14539)


tmp14541 := Call1(__e, tmp14537, tmp14540)


tmp14542 := Call1(__e, tmp14541, MakeInteger(0))


tmp14543 := MakeNative0(func(__e *ControlFlow) Obj {
return True
panic("unreachable")
})

return TailCall1(__e, tmp14542, tmp14543)


panic("unreachable")
})

tmp14544 := Call2(__e, ns2_1set, symshen_4typecheck, tmp14520)


_ = tmp14544

tmp14545 := MakeNative(func(__e *ControlFlow) {
V4925 := __e.Get(1)
V4926 := __e.Get(2)
V4927 := __e.Get(3)
V4928 := __e.Get(4)
V4929 := __e.Get(5)
V4930 := __e.Get(6)
V4931 := __e.Get(7)
__e.ReturnDirect(func() Obj {
tmp14552 := Call1(__e, SymFn(symshen_4unlocked_2), V4929)


var ifres14546 Obj

if True == tmp14552 {
tmp14547 := Call2(__e, SymFn(symshen_4lazyderef), V4925, V4928)


let__13974 := tmp14547
_ = let__13974

tmp14551 := GEq(__e, let__13974, Nil)

var ifres14548 Obj

if True == tmp14551 {
tmp14549 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14549

tmp14550 := Call6(__e, SymFn(symis_b), V4926, V4927, V4928, V4929, V4930, V4931)


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

tmp14566 := GEq(__e, let__13973, False)

if True == tmp14566 {
tmp14564 := Call1(__e, SymFn(symshen_4unlocked_2), V4929)


if True == tmp14564 {
tmp14553 := Call2(__e, SymFn(symshen_4lazyderef), V4925, V4928)


let__13975 := tmp14553
_ = let__13975

tmp14562 := GConsP(__e, let__13975)

if True == tmp14562 {
tmp14554 := GHd(__e, let__13975)

let__13976 := tmp14554
_ = let__13976

tmp14555 := GTl(__e, let__13975)

let__13977 := tmp14555
_ = let__13977

tmp14556 := Call1(__e, SymFn(symshen_4newpv), V4928)


let__13978 := tmp14556
_ = let__13978

tmp14557 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14557

tmp14558 := Call2(__e, SymFn(symshen_4deref), let__13978, V4928)


tmp14559 := Call3(__e, SymFn(symsubst), tmp14558, let__13976, V4926)


tmp14560 := Call(__e, SymFn(symshen_4insert_1prolog_1variables), let__13977, tmp14559, V4927, V4928, V4929, V4930, V4931)


return TailCall2(__e, SymFn(symshen_4gc), V4928, tmp14560)


} else {
return False
}


} else {
return False
}


} else {
return let__13973
}


panic("unreachable")
}())
}, 7)

tmp14567 := Call2(__e, ns2_1set, symshen_4insert_1prolog_1variables, tmp14545)


_ = tmp14567

tmp14568 := MakeNative6(func(__e *ControlFlow, V4938, V4939, V4940, V4941, V4942, V4943 Obj) Obj {
let__13979 := GAdd(__e, V4942, MakeInteger(1))
_ = let__13979

tmp14593 := Call1(__e, SymFn(symshen_4unlocked_2), V4941)


var ifres14570 Obj

if True == tmp14593 {
tmp14571 := Call2(__e, SymFn(symshen_4lazyderef), V4938, V4940)


let__13981 := tmp14571
_ = let__13981

tmp14592 := GConsP(__e, let__13981)

var ifres14572 Obj

if True == tmp14592 {
tmp14573 := GHd(__e, let__13981)

tmp14574 := Call2(__e, SymFn(symshen_4lazyderef), tmp14573, V4940)


let__13982 := tmp14574
_ = let__13982

tmp14591 := GEq(__e, let__13982, symdefine)

var ifres14575 Obj

if True == tmp14591 {
tmp14576 := GTl(__e, let__13981)

tmp14577 := Call2(__e, SymFn(symshen_4lazyderef), tmp14576, V4940)


let__13983 := tmp14577
_ = let__13983

tmp14590 := GConsP(__e, let__13983)

var ifres14578 Obj

if True == tmp14590 {
tmp14579 := GHd(__e, let__13983)

let__13984 := tmp14579
_ = let__13984

tmp14580 := GTl(__e, let__13983)

let__13985 := tmp14580
_ = let__13985

tmp14581 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14581

tmp14582 := Call0(__e, SymFn(symshen_4type_1theory_1enabled_2))


tmp14583 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14584 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14585 := GValue(__e, symshen_4_dspy_d)

tmp14586 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14587 := GCons(__e, let__13984, let__13985)

tmp14588 := GCons(__e, symdefine, tmp14587)

return TailCall6(__e, SymFn(symshen_4t_d), tmp14588, V4939, V4940, V4941, let__13979, V4943)


panic("unreachable")
})

return TailCall6(__e, SymFn(symshen_4signal_1def), tmp14585, let__13984, V4940, V4941, let__13979, tmp14586)


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V4940, V4941, let__13979, tmp14584)


panic("unreachable")
})

tmp14589 := Call5(__e, SymFn(symwhen), tmp14582, V4940, V4941, let__13979, tmp14583)


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

tmp14605 := GEq(__e, let__13980, False)

if True == tmp14605 {
tmp14601 := Call1(__e, SymFn(symshen_4unlocked_2), V4941)


var ifres14594 Obj

if True == tmp14601 {
tmp14595 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14595

tmp14596 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp14597 := GCons(__e, V4939, Nil)

tmp14598 := GCons(__e, tmp14596, tmp14597)

tmp14599 := GCons(__e, V4938, tmp14598)

tmp14600 := Call6(__e, SymFn(symshen_4system_1S), tmp14599, Nil, V4940, V4941, let__13979, V4943)


ifres14594 = tmp14600


} else {
ifres14594 = False


}

let__13986 := ifres14594
_ = let__13986

tmp14603 := GEq(__e, let__13986, False)

if True == tmp14603 {
return TailCall2(__e, SymFn(symshen_4unlock), V4941, let__13979)
} else {
return let__13986
}


} else {
return let__13980
}


panic("unreachable")
})

tmp14606 := Call2(__e, ns2_1set, symshen_4toplevel_1forms, tmp14568)


_ = tmp14606

tmp14607 := MakeNative6(func(__e *ControlFlow, V4952, V4953, V4954, V4955, V4956, V4957 Obj) Obj {
tmp14614 := Call1(__e, SymFn(symshen_4unlocked_2), V4955)


var ifres14608 Obj

if True == tmp14614 {
tmp14609 := Call2(__e, SymFn(symshen_4lazyderef), V4952, V4954)


let__13988 := tmp14609
_ = let__13988

tmp14613 := GEq(__e, let__13988, False)

var ifres14610 Obj

if True == tmp14613 {
tmp14611 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14611

tmp14612 := Call1(__e, SymFn(symthaw), V4957)


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

tmp14629 := GEq(__e, let__13987, False)

if True == tmp14629 {
tmp14627 := Call1(__e, SymFn(symshen_4unlocked_2), V4955)


if True == tmp14627 {
tmp14615 := Call2(__e, SymFn(symshen_4lazyderef), V4952, V4954)


let__13989 := tmp14615
_ = let__13989

tmp14625 := GEq(__e, let__13989, True)

if True == tmp14625 {
tmp14616 := Call1(__e, SymFn(symshen_4newpv), V4954)


let__13990 := tmp14616
_ = let__13990

tmp14617 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14617

tmp14618 := Call2(__e, SymFn(symshen_4deref), V4953, V4954)


tmp14619 := Call3(__e, SymFn(symshen_4app), tmp14618, MakeString(")\n"), symshen_4a)


tmp14620 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("\ntypechecking (fn "), tmp14619)

tmp14621 := Call0(__e, SymFn(symstoutput))


tmp14622 := Call2(__e, SymFn(sympr), tmp14620, tmp14621)


tmp14623 := Call6(__e, SymFn(symis), let__13990, tmp14622, V4954, V4955, V4956, V4957)


return TailCall2(__e, SymFn(symshen_4gc), V4954, tmp14623)


} else {
return False
}


} else {
return False
}


} else {
return let__13987
}


panic("unreachable")
})

tmp14630 := Call2(__e, ns2_1set, symshen_4signal_1def, tmp14607)


_ = tmp14630

tmp14631 := MakeNative1(func(__e *ControlFlow, V4962 Obj) Obj {
tmp14632 := Call1(__e, SymFn(symshen_4curry_1type), V4962)


return TailCall1(__e, SymFn(symshen_4demodulate), tmp14632)


panic("unreachable")
})

tmp14633 := Call2(__e, ns2_1set, symshen_4rectify_1type, tmp14631)


_ = tmp14633

tmp14634 := MakeNative1(func(__e *ControlFlow, V4963 Obj) Obj {
tmp14635 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14636 := MakeNative1(func(__e *ControlFlow, Z4965 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4demod), Z4965)
panic("unreachable")
})

tmp14637 := Call2(__e, SymFn(symshen_4walk), tmp14636, V4963)


let__13991 := tmp14637
_ = let__13991

tmp14639 := GEq(__e, let__13991, V4963)

if True == tmp14639 {
return V4963
} else {
return TailCall1(__e, SymFn(symshen_4demodulate), let__13991)
}


panic("unreachable")
})

tmp14640 := MakeNative1(func(__e *ControlFlow, Z4966 Obj) Obj {
return V4963
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp14635, tmp14640)


panic("unreachable")
})

tmp14641 := Call2(__e, ns2_1set, symshen_4demodulate, tmp14634)


_ = tmp14641

tmp14642 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V4967 := __self1
_ = V4967
tmp14766 := GConsP(__e, V4967)

var ifres14739 Obj

if True == tmp14766 {
tmp14764 := GTl(__e, V4967)

tmp14765 := GConsP(__e, tmp14764)

var ifres14741 Obj

if True == tmp14765 {
tmp14761 := GTl(__e, V4967)

tmp14762 := GHd(__e, tmp14761)

tmp14763 := GEq(__e, sym_1_1_6, tmp14762)

var ifres14743 Obj

if True == tmp14763 {
tmp14758 := GTl(__e, V4967)

tmp14759 := GTl(__e, tmp14758)

tmp14760 := GConsP(__e, tmp14759)

var ifres14745 Obj

if True == tmp14760 {
tmp14754 := GTl(__e, V4967)

tmp14755 := GTl(__e, tmp14754)

tmp14756 := GTl(__e, tmp14755)

tmp14757 := GConsP(__e, tmp14756)

var ifres14747 Obj

if True == tmp14757 {
tmp14749 := GTl(__e, V4967)

tmp14750 := GTl(__e, tmp14749)

tmp14751 := GTl(__e, tmp14750)

tmp14752 := GHd(__e, tmp14751)

tmp14753 := GEq(__e, sym_1_1_6, tmp14752)

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
tmp14643 := GHd(__e, V4967)

tmp14644 := GTl(__e, V4967)

tmp14645 := GTl(__e, tmp14644)

tmp14646 := GCons(__e, tmp14645, Nil)

tmp14647 := GCons(__e, sym_1_1_6, tmp14646)

tmp14648 := GCons(__e, tmp14643, tmp14647)

if SymFn(symshen_4curry_1type) == __self {
__self1 = tmp14648
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4curry_1type), tmp14648)


} else {
tmp14737 := GConsP(__e, V4967)

var ifres14697 Obj

if True == tmp14737 {
tmp14735 := GHd(__e, V4967)

tmp14736 := GConsP(__e, tmp14735)

var ifres14699 Obj

if True == tmp14736 {
tmp14732 := GHd(__e, V4967)

tmp14733 := GHd(__e, tmp14732)

tmp14734 := GEq(__e, symlist, tmp14733)

var ifres14701 Obj

if True == tmp14734 {
tmp14729 := GHd(__e, V4967)

tmp14730 := GTl(__e, tmp14729)

tmp14731 := GConsP(__e, tmp14730)

var ifres14703 Obj

if True == tmp14731 {
tmp14725 := GHd(__e, V4967)

tmp14726 := GTl(__e, tmp14725)

tmp14727 := GTl(__e, tmp14726)

tmp14728 := GEq(__e, Nil, tmp14727)

var ifres14705 Obj

if True == tmp14728 {
tmp14723 := GTl(__e, V4967)

tmp14724 := GConsP(__e, tmp14723)

var ifres14707 Obj

if True == tmp14724 {
tmp14720 := GTl(__e, V4967)

tmp14721 := GHd(__e, tmp14720)

tmp14722 := GEq(__e, sym_a_a_6, tmp14721)

var ifres14709 Obj

if True == tmp14722 {
tmp14717 := GTl(__e, V4967)

tmp14718 := GTl(__e, tmp14717)

tmp14719 := GConsP(__e, tmp14718)

var ifres14711 Obj

if True == tmp14719 {
tmp14713 := GTl(__e, V4967)

tmp14714 := GTl(__e, tmp14713)

tmp14715 := GTl(__e, tmp14714)

tmp14716 := GEq(__e, Nil, tmp14715)

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
tmp14649 := GHd(__e, V4967)

tmp14650 := GHd(__e, V4967)

tmp14651 := GTl(__e, V4967)

tmp14652 := GTl(__e, tmp14651)

tmp14653 := GCons(__e, tmp14650, tmp14652)

tmp14654 := GCons(__e, symstr, tmp14653)

tmp14655 := GCons(__e, tmp14654, Nil)

tmp14656 := GCons(__e, sym_1_1_6, tmp14655)

tmp14657 := GCons(__e, tmp14649, tmp14656)

if SymFn(symshen_4curry_1type) == __self {
__self1 = tmp14657
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4curry_1type), tmp14657)


} else {
tmp14695 := GConsP(__e, V4967)

var ifres14668 Obj

if True == tmp14695 {
tmp14693 := GTl(__e, V4967)

tmp14694 := GConsP(__e, tmp14693)

var ifres14670 Obj

if True == tmp14694 {
tmp14690 := GTl(__e, V4967)

tmp14691 := GHd(__e, tmp14690)

tmp14692 := GEq(__e, sym_d, tmp14691)

var ifres14672 Obj

if True == tmp14692 {
tmp14687 := GTl(__e, V4967)

tmp14688 := GTl(__e, tmp14687)

tmp14689 := GConsP(__e, tmp14688)

var ifres14674 Obj

if True == tmp14689 {
tmp14683 := GTl(__e, V4967)

tmp14684 := GTl(__e, tmp14683)

tmp14685 := GTl(__e, tmp14684)

tmp14686 := GConsP(__e, tmp14685)

var ifres14676 Obj

if True == tmp14686 {
tmp14678 := GTl(__e, V4967)

tmp14679 := GTl(__e, tmp14678)

tmp14680 := GTl(__e, tmp14679)

tmp14681 := GHd(__e, tmp14680)

tmp14682 := GEq(__e, sym_d, tmp14681)

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
tmp14658 := GHd(__e, V4967)

tmp14659 := GTl(__e, V4967)

tmp14660 := GTl(__e, tmp14659)

tmp14661 := GCons(__e, tmp14660, Nil)

tmp14662 := GCons(__e, sym_d, tmp14661)

tmp14663 := GCons(__e, tmp14658, tmp14662)

if SymFn(symshen_4curry_1type) == __self {
__self1 = tmp14663
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4curry_1type), tmp14663)


} else {
tmp14666 := GConsP(__e, V4967)

if True == tmp14666 {
tmp14664 := MakeNative1(func(__e *ControlFlow, Z4968 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4curry_1type), Z4968)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp14664, V4967)


} else {
return V4967
}


}


}


}


panic("unreachable")
})
return __self
})()

tmp14767 := Call2(__e, ns2_1set, symshen_4curry_1type, tmp14642)


_ = tmp14767

tmp14768 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V4969 := __self1
_ = V4969
tmp14857 := GConsP(__e, V4969)

var ifres14849 Obj

if True == tmp14857 {
tmp14855 := GHd(__e, V4969)

tmp14856 := GEq(__e, symdefine, tmp14855)

var ifres14851 Obj

if True == tmp14856 {
tmp14853 := GTl(__e, V4969)

tmp14854 := GConsP(__e, tmp14853)

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
return V4969
} else {
tmp14847 := GConsP(__e, V4969)

var ifres14828 Obj

if True == tmp14847 {
tmp14845 := GHd(__e, V4969)

tmp14846 := GEq(__e, symtype, tmp14845)

var ifres14830 Obj

if True == tmp14846 {
tmp14843 := GTl(__e, V4969)

tmp14844 := GConsP(__e, tmp14843)

var ifres14832 Obj

if True == tmp14844 {
tmp14840 := GTl(__e, V4969)

tmp14841 := GTl(__e, tmp14840)

tmp14842 := GConsP(__e, tmp14841)

var ifres14834 Obj

if True == tmp14842 {
tmp14836 := GTl(__e, V4969)

tmp14837 := GTl(__e, tmp14836)

tmp14838 := GTl(__e, tmp14837)

tmp14839 := GEq(__e, Nil, tmp14838)

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
tmp14769 := GTl(__e, V4969)

tmp14770 := GHd(__e, tmp14769)

tmp14771 := Call1(__e, SymFn(symshen_4curry), tmp14770)


tmp14772 := GTl(__e, V4969)

tmp14773 := GTl(__e, tmp14772)

tmp14774 := GCons(__e, tmp14771, tmp14773)

return GCons(__e, symtype, tmp14774)


} else {
tmp14826 := GConsP(__e, V4969)

var ifres14822 Obj

if True == tmp14826 {
tmp14824 := GHd(__e, V4969)

tmp14825 := Call1(__e, SymFn(symshen_4special_2), tmp14824)


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
tmp14775 := GHd(__e, V4969)

tmp14776 := MakeNative1(func(__e *ControlFlow, Z4970 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4curry), Z4970)
panic("unreachable")
})

tmp14777 := GTl(__e, V4969)

tmp14778 := Call2(__e, SymFn(symmap), tmp14776, tmp14777)


return GCons(__e, tmp14775, tmp14778)


} else {
tmp14820 := GConsP(__e, V4969)

var ifres14816 Obj

if True == tmp14820 {
tmp14818 := GHd(__e, V4969)

tmp14819 := Call1(__e, SymFn(symshen_4extraspecial_2), tmp14818)


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
return V4969
} else {
tmp14814 := GConsP(__e, V4969)

var ifres14805 Obj

if True == tmp14814 {
tmp14812 := GTl(__e, V4969)

tmp14813 := GConsP(__e, tmp14812)

var ifres14807 Obj

if True == tmp14813 {
tmp14809 := GTl(__e, V4969)

tmp14810 := GTl(__e, tmp14809)

tmp14811 := GConsP(__e, tmp14810)

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
tmp14779 := GHd(__e, V4969)

tmp14780 := GTl(__e, V4969)

tmp14781 := GHd(__e, tmp14780)

tmp14782 := GCons(__e, tmp14781, Nil)

tmp14783 := GCons(__e, tmp14779, tmp14782)

tmp14784 := GTl(__e, V4969)

tmp14785 := GTl(__e, tmp14784)

tmp14786 := GCons(__e, tmp14783, tmp14785)

if SymFn(symshen_4curry) == __self {
__self1 = tmp14786
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4curry), tmp14786)


} else {
tmp14803 := GConsP(__e, V4969)

var ifres14794 Obj

if True == tmp14803 {
tmp14801 := GTl(__e, V4969)

tmp14802 := GConsP(__e, tmp14801)

var ifres14796 Obj

if True == tmp14802 {
tmp14798 := GTl(__e, V4969)

tmp14799 := GTl(__e, tmp14798)

tmp14800 := GEq(__e, Nil, tmp14799)

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
tmp14787 := GHd(__e, V4969)

tmp14788 := Call1(__e, SymFn(symshen_4curry), tmp14787)


tmp14789 := GTl(__e, V4969)

tmp14790 := GHd(__e, tmp14789)

tmp14791 := Call1(__e, SymFn(symshen_4curry), tmp14790)


tmp14792 := GCons(__e, tmp14791, Nil)

return GCons(__e, tmp14788, tmp14792)


} else {
return V4969
}


}


}


}


}


}


panic("unreachable")
})
return __self
})()

tmp14858 := Call2(__e, ns2_1set, symshen_4curry, tmp14768)


_ = tmp14858

tmp14859 := MakeNative1(func(__e *ControlFlow, V4971 Obj) Obj {
tmp14860 := GValue(__e, symshen_4_dspecial_d)

return TailCall2(__e, SymFn(symelement_2), V4971, tmp14860)


panic("unreachable")
})

tmp14861 := Call2(__e, ns2_1set, symshen_4special_2, tmp14859)


_ = tmp14861

tmp14862 := MakeNative1(func(__e *ControlFlow, V4972 Obj) Obj {
tmp14863 := GValue(__e, symshen_4_dextraspecial_d)

return TailCall2(__e, SymFn(symelement_2), V4972, tmp14863)


panic("unreachable")
})

tmp14864 := Call2(__e, ns2_1set, symshen_4extraspecial_2, tmp14862)


_ = tmp14864

tmp14865 := MakeNative6(func(__e *ControlFlow, V4973, V4974, V4975, V4976, V4977, V4978 Obj) Obj {
let__13992 := GAdd(__e, V4977, MakeInteger(1))
_ = let__13992

tmp14871 := Call1(__e, SymFn(symshen_4unlocked_2), V4976)


var ifres14867 Obj

if True == tmp14871 {
tmp14868 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14868

tmp14869 := Call0(__e, SymFn(symshen_4maxinfexceeded_2))


tmp14870 := Call5(__e, SymFn(symwhen), tmp14869, V4975, V4976, let__13992, V4978)


ifres14867 = tmp14870


} else {
ifres14867 = False


}

let__13993 := ifres14867
_ = let__13993

tmp14919 := GEq(__e, let__13993, False)

if True == tmp14919 {
tmp14900 := Call1(__e, SymFn(symshen_4unlocked_2), V4976)


var ifres14872 Obj

if True == tmp14900 {
tmp14873 := Call2(__e, SymFn(symshen_4lazyderef), V4973, V4975)


let__13995 := tmp14873
_ = let__13995

tmp14899 := GConsP(__e, let__13995)

var ifres14874 Obj

if True == tmp14899 {
tmp14875 := GHd(__e, let__13995)

let__13996 := tmp14875
_ = let__13996

tmp14876 := GTl(__e, let__13995)

tmp14877 := Call2(__e, SymFn(symshen_4lazyderef), tmp14876, V4975)


let__13997 := tmp14877
_ = let__13997

tmp14898 := GConsP(__e, let__13997)

var ifres14878 Obj

if True == tmp14898 {
tmp14879 := GHd(__e, let__13997)

let__13998 := tmp14879
_ = let__13998

tmp14880 := GTl(__e, let__13997)

tmp14881 := Call2(__e, SymFn(symshen_4lazyderef), tmp14880, V4975)


let__13999 := tmp14881
_ = let__13999

tmp14897 := GConsP(__e, let__13999)

var ifres14882 Obj

if True == tmp14897 {
tmp14883 := GHd(__e, let__13999)

let__14000 := tmp14883
_ = let__14000

tmp14884 := GTl(__e, let__13999)

tmp14885 := Call2(__e, SymFn(symshen_4lazyderef), tmp14884, V4975)


let__14001 := tmp14885
_ = let__14001

tmp14896 := GEq(__e, let__14001, Nil)

var ifres14886 Obj

if True == tmp14896 {
tmp14887 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14887

tmp14888 := Call2(__e, SymFn(symshen_4deref), let__13998, V4975)


tmp14889 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp14890 := GEq(__e, tmp14888, tmp14889)

tmp14891 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14892 := Call0(__e, SymFn(symshen_4type_1theory_1enabled_2))


tmp14893 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14894 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), let__13996, let__14000, V4974, V4975, V4976, let__13992, V4978)
return nil
panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V4975, V4976, let__13992, tmp14894)


panic("unreachable")
})

return TailCall5(__e, SymFn(symwhen), tmp14892, V4975, V4976, let__13992, tmp14893)


panic("unreachable")
})

tmp14895 := Call5(__e, SymFn(symwhen), tmp14890, V4975, V4976, let__13992, tmp14891)


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

tmp14917 := GEq(__e, let__13994, False)

if True == tmp14917 {
tmp14906 := Call1(__e, SymFn(symshen_4unlocked_2), V4976)


var ifres14901 Obj

if True == tmp14906 {
tmp14902 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14902

tmp14903 := GValue(__e, symshen_4_dspy_d)

tmp14904 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall6(__e, SymFn(symshen_4show), V4973, V4974, V4975, V4976, let__13992, V4978)
panic("unreachable")
})

tmp14905 := Call5(__e, SymFn(symwhen), tmp14903, V4975, V4976, let__13992, tmp14904)


ifres14901 = tmp14905


} else {
ifres14901 = False


}

let__14002 := ifres14901
_ = let__14002

tmp14915 := GEq(__e, let__14002, False)

if True == tmp14915 {
tmp14911 := Call1(__e, SymFn(symshen_4unlocked_2), V4976)


var ifres14907 Obj

if True == tmp14911 {
tmp14908 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14908

tmp14909 := GValue(__e, symshen_4_ddatatypes_d)

tmp14910 := Call(__e, SymFn(symshen_4search_1user_1datatypes), V4973, V4974, tmp14909, V4975, V4976, let__13992, V4978)


ifres14907 = tmp14910


} else {
ifres14907 = False


}

let__14003 := ifres14907
_ = let__14003

tmp14913 := GEq(__e, let__14003, False)

if True == tmp14913 {
return TailCall2(__e, SymFn(symshen_4unlock), V4976, let__13992)
} else {
return let__14003
}


} else {
return let__14002
}


} else {
return let__13994
}


} else {
return let__13993
}


panic("unreachable")
})

tmp14920 := Call2(__e, ns2_1set, symshen_4system_1S, tmp14865)


_ = tmp14920

tmp14921 := MakeNative6(func(__e *ControlFlow, V4997, V4998, V4999, V5000, V5001, V5002 Obj) Obj {
tmp14922 := Call0(__e, SymFn(symshen_4line))


_ = tmp14922

tmp14923 := Call2(__e, SymFn(symshen_4deref), V4997, V4999)


tmp14924 := Call1(__e, SymFn(symshen_4show_1p), tmp14923)


_ = tmp14924

tmp14925 := Call1(__e, SymFn(symnl), MakeInteger(2))


_ = tmp14925

tmp14926 := Call2(__e, SymFn(symshen_4deref), V4998, V4999)


tmp14927 := Call2(__e, SymFn(symshen_4show_1assumptions), tmp14926, MakeInteger(1))


_ = tmp14927

tmp14928 := Call0(__e, SymFn(symshen_4pause_1for_1user))


_ = tmp14928

return False


panic("unreachable")
})

tmp14929 := Call2(__e, ns2_1set, symshen_4show, tmp14921)


_ = tmp14929

tmp14930 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14931 := Call0(__e, SymFn(syminferences))


let__14004 := tmp14931
_ = let__14004

tmp14933 := GEq(__e, MakeInteger(1), let__14004)

var ifres14932 Obj

if True == tmp14933 {
ifres14932 = MakeString("")


} else {
ifres14932 = MakeString("s")


}

tmp14934 := Call3(__e, SymFn(symshen_4app), ifres14932, MakeString(" \n?- "), symshen_4a)


tmp14936 := Call3(__e, SymFn(symshen_4app), let__14004, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" inference"), tmp14934), symshen_4a)


tmp14937 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("____________________________________________________________ "), tmp14936)

tmp14938 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp14937, tmp14938)


panic("unreachable")
})

tmp14939 := Call2(__e, ns2_1set, symshen_4line, tmp14930)


_ = tmp14939

tmp14940 := MakeNative1(func(__e *ControlFlow, V5004 Obj) Obj {
tmp14972 := GConsP(__e, V5004)

var ifres14951 Obj

if True == tmp14972 {
tmp14970 := GTl(__e, V5004)

tmp14971 := GConsP(__e, tmp14970)

var ifres14953 Obj

if True == tmp14971 {
tmp14967 := GTl(__e, V5004)

tmp14968 := GTl(__e, tmp14967)

tmp14969 := GConsP(__e, tmp14968)

var ifres14955 Obj

if True == tmp14969 {
tmp14963 := GTl(__e, V5004)

tmp14964 := GTl(__e, tmp14963)

tmp14965 := GTl(__e, tmp14964)

tmp14966 := GEq(__e, Nil, tmp14965)

var ifres14957 Obj

if True == tmp14966 {
tmp14959 := GTl(__e, V5004)

tmp14960 := GHd(__e, tmp14959)

tmp14961 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp14962 := GEq(__e, tmp14960, tmp14961)

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
tmp14941 := GHd(__e, V5004)

tmp14942 := Call1(__e, SymFn(symshen_4prterm), tmp14941)


_ = tmp14942

tmp14943 := Call0(__e, SymFn(symstoutput))


tmp14944 := Call2(__e, SymFn(sympr), MakeString(" : "), tmp14943)


_ = tmp14944

tmp14945 := GTl(__e, V5004)

tmp14946 := GTl(__e, tmp14945)

tmp14947 := GHd(__e, tmp14946)

tmp14948 := Call3(__e, SymFn(symshen_4app), tmp14947, MakeString(""), symshen_4r)


tmp14949 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp14948, tmp14949)


} else {
return TailCall1(__e, SymFn(symshen_4prterm), V5004)
}


panic("unreachable")
})

tmp14973 := Call2(__e, ns2_1set, symshen_4show_1p, tmp14940)


_ = tmp14973

tmp14974 := MakeNative1(func(__e *ControlFlow, V5005 Obj) Obj {
tmp15017 := GConsP(__e, V5005)

var ifres14998 Obj

if True == tmp15017 {
tmp15015 := GHd(__e, V5005)

tmp15016 := GEq(__e, symcons, tmp15015)

var ifres15000 Obj

if True == tmp15016 {
tmp15013 := GTl(__e, V5005)

tmp15014 := GConsP(__e, tmp15013)

var ifres15002 Obj

if True == tmp15014 {
tmp15010 := GTl(__e, V5005)

tmp15011 := GTl(__e, tmp15010)

tmp15012 := GConsP(__e, tmp15011)

var ifres15004 Obj

if True == tmp15012 {
tmp15006 := GTl(__e, V5005)

tmp15007 := GTl(__e, tmp15006)

tmp15008 := GTl(__e, tmp15007)

tmp15009 := GEq(__e, Nil, tmp15008)

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
tmp14975 := Call0(__e, SymFn(symstoutput))


tmp14976 := Call2(__e, SymFn(sympr), MakeString("["), tmp14975)


_ = tmp14976

tmp14977 := GTl(__e, V5005)

tmp14978 := GHd(__e, tmp14977)

tmp14979 := Call1(__e, SymFn(symshen_4prterm), tmp14978)


_ = tmp14979

tmp14980 := GTl(__e, V5005)

tmp14981 := GTl(__e, tmp14980)

tmp14982 := GHd(__e, tmp14981)

tmp14983 := Call1(__e, SymFn(symshen_4prtl), tmp14982)


_ = tmp14983

tmp14984 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), MakeString("]"), tmp14984)


} else {
tmp14996 := GConsP(__e, V5005)

if True == tmp14996 {
tmp14985 := Call0(__e, SymFn(symstoutput))


tmp14986 := Call2(__e, SymFn(sympr), MakeString("("), tmp14985)


_ = tmp14986

tmp14987 := GHd(__e, V5005)

tmp14988 := Call1(__e, SymFn(symshen_4prterm), tmp14987)


_ = tmp14988

tmp14989 := MakeNative1(func(__e *ControlFlow, Z5006 Obj) Obj {
tmp14990 := Call0(__e, SymFn(symstoutput))


tmp14991 := Call2(__e, SymFn(sympr), MakeString(" "), tmp14990)


_ = tmp14991

return TailCall1(__e, SymFn(symshen_4prterm), Z5006)


panic("unreachable")
})

tmp14992 := GTl(__e, V5005)

tmp14993 := Call2(__e, SymFn(symmap), tmp14989, tmp14992)


_ = tmp14993

tmp14994 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), MakeString(")"), tmp14994)


} else {
return TailCall1(__e, SymFn(symprint), V5005)
}


}


panic("unreachable")
})

tmp15018 := Call2(__e, ns2_1set, symshen_4prterm, tmp14974)


_ = tmp15018

tmp15019 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V5007 := __self1
_ = V5007
tmp15052 := GEq(__e, Nil, V5007)

if True == tmp15052 {
return MakeString("")
} else {
tmp15050 := GConsP(__e, V5007)

var ifres15031 Obj

if True == tmp15050 {
tmp15048 := GHd(__e, V5007)

tmp15049 := GEq(__e, symcons, tmp15048)

var ifres15033 Obj

if True == tmp15049 {
tmp15046 := GTl(__e, V5007)

tmp15047 := GConsP(__e, tmp15046)

var ifres15035 Obj

if True == tmp15047 {
tmp15043 := GTl(__e, V5007)

tmp15044 := GTl(__e, tmp15043)

tmp15045 := GConsP(__e, tmp15044)

var ifres15037 Obj

if True == tmp15045 {
tmp15039 := GTl(__e, V5007)

tmp15040 := GTl(__e, tmp15039)

tmp15041 := GTl(__e, tmp15040)

tmp15042 := GEq(__e, Nil, tmp15041)

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
tmp15020 := Call0(__e, SymFn(symstoutput))


tmp15021 := Call2(__e, SymFn(sympr), MakeString(" "), tmp15020)


_ = tmp15021

tmp15022 := GTl(__e, V5007)

tmp15023 := GHd(__e, tmp15022)

tmp15024 := Call1(__e, SymFn(symshen_4prterm), tmp15023)


_ = tmp15024

tmp15025 := GTl(__e, V5007)

tmp15026 := GTl(__e, tmp15025)

tmp15027 := GHd(__e, tmp15026)

if SymFn(symshen_4prtl) == __self {
__self1 = tmp15027
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4prtl), tmp15027)


} else {
tmp15028 := Call0(__e, SymFn(symstoutput))


tmp15029 := Call2(__e, SymFn(sympr), MakeString(" | "), tmp15028)


_ = tmp15029

return TailCall1(__e, SymFn(symshen_4prterm), V5007)


}


}


panic("unreachable")
})
return __self
})()

tmp15053 := Call2(__e, ns2_1set, symshen_4prtl, tmp15019)


_ = tmp15053

tmp15054 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V5014 := __self1
_ = V5014
V5015 := __self2
_ = V5015
tmp15067 := GEq(__e, Nil, V5014)

if True == tmp15067 {
tmp15055 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), MakeString("\n> "), tmp15055)


} else {
tmp15065 := GConsP(__e, V5014)

if True == tmp15065 {
tmp15056 := Call3(__e, SymFn(symshen_4app), V5015, MakeString(". "), symshen_4a)


tmp15057 := Call0(__e, SymFn(symstoutput))


tmp15058 := Call2(__e, SymFn(sympr), tmp15056, tmp15057)


_ = tmp15058

tmp15059 := GHd(__e, V5014)

tmp15060 := Call1(__e, SymFn(symshen_4show_1p), tmp15059)


_ = tmp15060

tmp15061 := Call1(__e, SymFn(symnl), MakeInteger(1))


_ = tmp15061

tmp15062 := GTl(__e, V5014)

if SymFn(symshen_4show_1assumptions) == __self {
__self1, __self2 = tmp15062, GAdd(__e, V5015, MakeInteger(1))
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4show_1assumptions), tmp15062, GAdd(__e, V5015, MakeInteger(1)))


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.show-assumptions"))
}


}


panic("unreachable")
})
return __self
})()

tmp15068 := Call2(__e, ns2_1set, symshen_4show_1assumptions, tmp15054)


_ = tmp15068

tmp15069 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15070 := Call0(__e, SymFn(symstinput))


tmp15071 := PrimGuard1(__e, symread_1byte, PrimReadByte, tmp15070)

let__14005 := tmp15071
_ = let__14005

tmp15073 := GEq(__e, let__14005, MakeInteger(94))

if True == tmp15073 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("input aborted\n"))
} else {
return TailCall1(__e, SymFn(symnl), MakeInteger(1))
}


panic("unreachable")
})

tmp15074 := Call2(__e, ns2_1set, symshen_4pause_1for_1user, tmp15069)


_ = tmp15074

tmp15075 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dshen_1type_1theory_1enabled_2_d)
panic("unreachable")
})

tmp15076 := Call2(__e, ns2_1set, symshen_4type_1theory_1enabled_2, tmp15075)


_ = tmp15076

tmp15077 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15079 := Call0(__e, SymFn(syminferences))


tmp15080 := GValue(__e, symshen_4_dmaxinferences_d)

if True == GGt(__e, tmp15079, tmp15080) {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("maximum inferences exceeded"))
} else {
return False
}


panic("unreachable")
})

tmp15082 := Call2(__e, ns2_1set, symshen_4maxinfexceeded_2, tmp15077)


_ = tmp15082

tmp15083 := MakeNative(func(__e *ControlFlow) {
V5017 := __e.Get(1)
V5018 := __e.Get(2)
V5019 := __e.Get(3)
V5020 := __e.Get(4)
V5021 := __e.Get(5)
V5022 := __e.Get(6)
V5023 := __e.Get(7)
__e.ReturnDirect(func() Obj {
let__14006 := GAdd(__e, V5022, MakeInteger(1))
_ = let__14006

tmp15094 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres15085 Obj

if True == tmp15094 {
tmp15086 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15086

tmp15087 := GValue(__e, symshen_4_dspy_d)

tmp15088 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15089 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp15090 := GCons(__e, V5018, Nil)

tmp15091 := GCons(__e, tmp15089, tmp15090)

tmp15092 := GCons(__e, V5017, tmp15091)

return TailCall6(__e, SymFn(symshen_4show), tmp15092, V5019, V5020, V5021, let__14006, V5023)


panic("unreachable")
})

tmp15093 := Call5(__e, SymFn(symwhen), tmp15087, V5020, V5021, let__14006, tmp15088)


ifres15085 = tmp15093


} else {
ifres15085 = False


}

let__14007 := ifres15085
_ = let__14007

tmp15833 := GEq(__e, let__14007, False)

if True == tmp15833 {
tmp15102 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres15095 Obj

if True == tmp15102 {
tmp15096 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15096

tmp15097 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


tmp15098 := GConsP(__e, tmp15097)

tmp15099 := GNot(__e, tmp15098)

tmp15100 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall6(__e, SymFn(symshen_4primitive), V5017, V5018, V5020, V5021, let__14006, V5023)
panic("unreachable")
})

tmp15101 := Call5(__e, SymFn(symwhen), tmp15099, V5020, V5021, let__14006, tmp15100)


ifres15095 = tmp15101


} else {
ifres15095 = False


}

let__14008 := ifres15095
_ = let__14008

tmp15831 := GEq(__e, let__14008, False)

if True == tmp15831 {
tmp15106 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres15103 Obj

if True == tmp15106 {
tmp15104 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15104

tmp15105 := Call(__e, SymFn(symshen_4by_1hypothesis), V5017, V5018, V5019, V5020, V5021, let__14006, V5023)


ifres15103 = tmp15105


} else {
ifres15103 = False


}

let__14009 := ifres15103
_ = let__14009

tmp15829 := GEq(__e, let__14009, False)

if True == tmp15829 {
tmp15120 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres15107 Obj

if True == tmp15120 {
tmp15108 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__14011 := tmp15108
_ = let__14011

tmp15119 := GConsP(__e, let__14011)

var ifres15109 Obj

if True == tmp15119 {
tmp15110 := GHd(__e, let__14011)

let__14012 := tmp15110
_ = let__14012

tmp15111 := GTl(__e, let__14011)

tmp15112 := Call2(__e, SymFn(symshen_4lazyderef), tmp15111, V5020)


let__14013 := tmp15112
_ = let__14013

tmp15118 := GEq(__e, let__14013, Nil)

var ifres15113 Obj

if True == tmp15118 {
tmp15114 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15114

tmp15115 := GCons(__e, V5018, Nil)

tmp15116 := GCons(__e, sym_1_1_6, tmp15115)

tmp15117 := Call6(__e, SymFn(symshen_4lookupsig), let__14012, tmp15116, V5020, V5021, let__14006, V5023)


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

tmp15827 := GEq(__e, let__14010, False)

if True == tmp15827 {
tmp15146 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres15121 Obj

if True == tmp15146 {
tmp15122 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__14015 := tmp15122
_ = let__14015

tmp15145 := GConsP(__e, let__14015)

var ifres15123 Obj

if True == tmp15145 {
tmp15124 := GHd(__e, let__14015)

tmp15125 := Call2(__e, SymFn(symshen_4lazyderef), tmp15124, V5020)


let__14016 := tmp15125
_ = let__14016

tmp15144 := GEq(__e, let__14016, symfn)

var ifres15126 Obj

if True == tmp15144 {
tmp15127 := GTl(__e, let__14015)

tmp15128 := Call2(__e, SymFn(symshen_4lazyderef), tmp15127, V5020)


let__14017 := tmp15128
_ = let__14017

tmp15143 := GConsP(__e, let__14017)

var ifres15129 Obj

if True == tmp15143 {
tmp15130 := GHd(__e, let__14017)

let__14018 := tmp15130
_ = let__14018

tmp15131 := GTl(__e, let__14017)

tmp15132 := Call2(__e, SymFn(symshen_4lazyderef), tmp15131, V5020)


let__14019 := tmp15132
_ = let__14019

tmp15142 := GEq(__e, let__14019, Nil)

var ifres15133 Obj

if True == tmp15142 {
tmp15134 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15134

tmp15135 := Call2(__e, SymFn(symshen_4deref), let__14018, V5020)


tmp15136 := Call1(__e, SymFn(symarity), tmp15135)


tmp15137 := GEq(__e, tmp15136, MakeInteger(0))

tmp15138 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15139 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15140 := GCons(__e, let__14018, Nil)

__e.TailApply(SymFn(symshen_4system_1S_1h), tmp15140, V5018, V5019, V5020, V5021, let__14006, V5023)
return nil


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5020, V5021, let__14006, tmp15139)


panic("unreachable")
})

tmp15141 := Call5(__e, SymFn(symwhen), tmp15137, V5020, V5021, let__14006, tmp15138)


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

tmp15825 := GEq(__e, let__14014, False)

if True == tmp15825 {
tmp15166 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres15147 Obj

if True == tmp15166 {
tmp15148 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__14021 := tmp15148
_ = let__14021

tmp15165 := GConsP(__e, let__14021)

var ifres15149 Obj

if True == tmp15165 {
tmp15150 := GHd(__e, let__14021)

tmp15151 := Call2(__e, SymFn(symshen_4lazyderef), tmp15150, V5020)


let__14022 := tmp15151
_ = let__14022

tmp15164 := GEq(__e, let__14022, symfn)

var ifres15152 Obj

if True == tmp15164 {
tmp15153 := GTl(__e, let__14021)

tmp15154 := Call2(__e, SymFn(symshen_4lazyderef), tmp15153, V5020)


let__14023 := tmp15154
_ = let__14023

tmp15163 := GConsP(__e, let__14023)

var ifres15155 Obj

if True == tmp15163 {
tmp15156 := GHd(__e, let__14023)

let__14024 := tmp15156
_ = let__14024

tmp15157 := GTl(__e, let__14023)

tmp15158 := Call2(__e, SymFn(symshen_4lazyderef), tmp15157, V5020)


let__14025 := tmp15158
_ = let__14025

tmp15162 := GEq(__e, let__14025, Nil)

var ifres15159 Obj

if True == tmp15162 {
tmp15160 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15160

tmp15161 := Call6(__e, SymFn(symshen_4lookupsig), let__14024, V5018, V5020, V5021, let__14006, V5023)


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

tmp15823 := GEq(__e, let__14020, False)

if True == tmp15823 {
tmp15193 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres15167 Obj

if True == tmp15193 {
tmp15168 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__14027 := tmp15168
_ = let__14027

tmp15192 := GConsP(__e, let__14027)

var ifres15169 Obj

if True == tmp15192 {
tmp15170 := GHd(__e, let__14027)

let__14028 := tmp15170
_ = let__14028

tmp15171 := GTl(__e, let__14027)

tmp15172 := Call2(__e, SymFn(symshen_4lazyderef), tmp15171, V5020)


let__14029 := tmp15172
_ = let__14029

tmp15191 := GConsP(__e, let__14029)

var ifres15173 Obj

if True == tmp15191 {
tmp15174 := GHd(__e, let__14029)

let__14030 := tmp15174
_ = let__14030

tmp15175 := GTl(__e, let__14029)

tmp15176 := Call2(__e, SymFn(symshen_4lazyderef), tmp15175, V5020)


let__14031 := tmp15176
_ = let__14031

tmp15190 := GEq(__e, let__14031, Nil)

var ifres15177 Obj

if True == tmp15190 {
tmp15178 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__14032 := tmp15178
_ = let__14032

tmp15179 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15179

tmp15180 := Call2(__e, SymFn(symshen_4lazyderef), let__14028, V5020)


tmp15181 := GConsP(__e, tmp15180)

tmp15182 := GNot(__e, tmp15181)

tmp15183 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15184 := GCons(__e, V5018, Nil)

tmp15185 := GCons(__e, sym_1_1_6, tmp15184)

tmp15186 := GCons(__e, let__14032, tmp15185)

tmp15187 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), let__14030, let__14032, V5019, V5020, V5021, let__14006, V5023)
return nil
panic("unreachable")
})

return TailCall6(__e, SymFn(symshen_4lookupsig), let__14028, tmp15186, V5020, V5021, let__14006, tmp15187)


panic("unreachable")
})

tmp15188 := Call5(__e, SymFn(symwhen), tmp15182, V5020, V5021, let__14006, tmp15183)


tmp15189 := Call2(__e, SymFn(symshen_4gc), V5020, tmp15188)


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

tmp15821 := GEq(__e, let__14026, False)

if True == tmp15821 {
tmp15216 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres15194 Obj

if True == tmp15216 {
tmp15195 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__14034 := tmp15195
_ = let__14034

tmp15215 := GConsP(__e, let__14034)

var ifres15196 Obj

if True == tmp15215 {
tmp15197 := GHd(__e, let__14034)

let__14035 := tmp15197
_ = let__14035

tmp15198 := GTl(__e, let__14034)

tmp15199 := Call2(__e, SymFn(symshen_4lazyderef), tmp15198, V5020)


let__14036 := tmp15199
_ = let__14036

tmp15214 := GConsP(__e, let__14036)

var ifres15200 Obj

if True == tmp15214 {
tmp15201 := GHd(__e, let__14036)

let__14037 := tmp15201
_ = let__14037

tmp15202 := GTl(__e, let__14036)

tmp15203 := Call2(__e, SymFn(symshen_4lazyderef), tmp15202, V5020)


let__14038 := tmp15203
_ = let__14038

tmp15213 := GEq(__e, let__14038, Nil)

var ifres15204 Obj

if True == tmp15213 {
tmp15205 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__14039 := tmp15205
_ = let__14039

tmp15206 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15206

tmp15207 := GCons(__e, V5018, Nil)

tmp15208 := GCons(__e, sym_1_1_6, tmp15207)

tmp15209 := GCons(__e, let__14039, tmp15208)

tmp15210 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), let__14037, let__14039, V5019, V5020, V5021, let__14006, V5023)
return nil
panic("unreachable")
})

tmp15211 := Call(__e, SymFn(symshen_4system_1S_1h), let__14035, tmp15209, V5019, V5020, V5021, let__14006, tmp15210)


tmp15212 := Call2(__e, SymFn(symshen_4gc), V5020, tmp15211)


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

tmp15819 := GEq(__e, let__14033, False)

if True == tmp15819 {
tmp15283 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres15217 Obj

if True == tmp15283 {
tmp15218 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__14041 := tmp15218
_ = let__14041

tmp15282 := GConsP(__e, let__14041)

var ifres15219 Obj

if True == tmp15282 {
tmp15220 := GHd(__e, let__14041)

tmp15221 := Call2(__e, SymFn(symshen_4lazyderef), tmp15220, V5020)


let__14042 := tmp15221
_ = let__14042

tmp15281 := GEq(__e, let__14042, symcons)

var ifres15222 Obj

if True == tmp15281 {
tmp15223 := GTl(__e, let__14041)

tmp15224 := Call2(__e, SymFn(symshen_4lazyderef), tmp15223, V5020)


let__14043 := tmp15224
_ = let__14043

tmp15280 := GConsP(__e, let__14043)

var ifres15225 Obj

if True == tmp15280 {
tmp15226 := GHd(__e, let__14043)

let__14044 := tmp15226
_ = let__14044

tmp15227 := GTl(__e, let__14043)

tmp15228 := Call2(__e, SymFn(symshen_4lazyderef), tmp15227, V5020)


let__14045 := tmp15228
_ = let__14045

tmp15279 := GConsP(__e, let__14045)

var ifres15229 Obj

if True == tmp15279 {
tmp15230 := GHd(__e, let__14045)

let__14046 := tmp15230
_ = let__14046

tmp15231 := GTl(__e, let__14045)

tmp15232 := Call2(__e, SymFn(symshen_4lazyderef), tmp15231, V5020)


let__14047 := tmp15232
_ = let__14047

tmp15278 := GEq(__e, let__14047, Nil)

var ifres15233 Obj

if True == tmp15278 {
tmp15234 := Call2(__e, SymFn(symshen_4lazyderef), V5018, V5020)


let__14048 := tmp15234
_ = let__14048

tmp15235 := MakeNative1(func(__e *ControlFlow, Z5068 Obj) Obj {
tmp15236 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15236

tmp15237 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15238 := GCons(__e, Z5068, Nil)

tmp15239 := GCons(__e, symlist, tmp15238)

__e.TailApply(SymFn(symshen_4system_1S_1h), let__14046, tmp15239, V5019, V5020, V5021, let__14006, V5023)
return nil


panic("unreachable")
})

__e.TailApply(SymFn(symshen_4system_1S_1h), let__14044, Z5068, V5019, V5020, V5021, let__14006, tmp15237)
return nil


panic("unreachable")
})

let__14049 := tmp15235
_ = let__14049

tmp15277 := GConsP(__e, let__14048)

var ifres15240 Obj

if True == tmp15277 {
tmp15241 := GHd(__e, let__14048)

tmp15242 := Call2(__e, SymFn(symshen_4lazyderef), tmp15241, V5020)


let__14050 := tmp15242
_ = let__14050

tmp15243 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15244 := GTl(__e, let__14048)

tmp15245 := Call2(__e, SymFn(symshen_4lazyderef), tmp15244, V5020)


let__14052 := tmp15245
_ = let__14052

tmp15246 := MakeNative1(func(__e *ControlFlow, Z5073 Obj) Obj {
return TailCall1(__e, let__14049, Z5073)
panic("unreachable")
})

let__14053 := tmp15246
_ = let__14053

tmp15262 := GConsP(__e, let__14052)

if True == tmp15262 {
tmp15247 := GHd(__e, let__14052)

let__14054 := tmp15247
_ = let__14054

tmp15248 := GTl(__e, let__14052)

tmp15249 := Call2(__e, SymFn(symshen_4lazyderef), tmp15248, V5020)


let__14055 := tmp15249
_ = let__14055

tmp15250 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__14053, let__14054)
panic("unreachable")
})

let__14056 := tmp15250
_ = let__14056

tmp15254 := GEq(__e, let__14055, Nil)

if True == tmp15254 {
return TailCall1(__e, SymFn(symthaw), let__14056)
} else {
tmp15252 := Call1(__e, SymFn(symshen_4pvar_2), let__14055)


if True == tmp15252 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__14055, Nil, V5020, let__14056)
} else {
return False
}


}


} else {
tmp15260 := Call1(__e, SymFn(symshen_4pvar_2), let__14052)


if True == tmp15260 {
tmp15255 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__14057 := tmp15255
_ = let__14057

tmp15256 := GCons(__e, let__14057, Nil)

tmp15257 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__14053, let__14057)
panic("unreachable")
})

tmp15258 := Call4(__e, SymFn(symshen_4bind_b), let__14052, tmp15256, V5020, tmp15257)


return TailCall2(__e, SymFn(symshen_4gc), V5020, tmp15258)


} else {
return False
}


}


panic("unreachable")
})

let__14051 := tmp15243
_ = let__14051

tmp15268 := GEq(__e, let__14050, symlist)

var ifres15263 Obj

if True == tmp15268 {
tmp15264 := Call1(__e, SymFn(symthaw), let__14051)


ifres15263 = tmp15264


} else {
tmp15267 := Call1(__e, SymFn(symshen_4pvar_2), let__14050)


var ifres15265 Obj

if True == tmp15267 {
tmp15266 := Call4(__e, SymFn(symshen_4bind_b), let__14050, symlist, V5020, let__14051)


ifres15265 = tmp15266


} else {
ifres15265 = False


}

ifres15263 = ifres15265


}

ifres15240 = ifres15263


} else {
tmp15276 := Call1(__e, SymFn(symshen_4pvar_2), let__14048)


var ifres15269 Obj

if True == tmp15276 {
tmp15270 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__14058 := tmp15270
_ = let__14058

tmp15271 := GCons(__e, let__14058, Nil)

tmp15272 := GCons(__e, symlist, tmp15271)

tmp15273 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__14049, let__14058)
panic("unreachable")
})

tmp15274 := Call4(__e, SymFn(symshen_4bind_b), let__14048, tmp15272, V5020, tmp15273)


tmp15275 := Call2(__e, SymFn(symshen_4gc), V5020, tmp15274)


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

tmp15817 := GEq(__e, let__14040, False)

if True == tmp15817 {
tmp15367 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres15284 Obj

if True == tmp15367 {
tmp15285 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__14060 := tmp15285
_ = let__14060

tmp15366 := GConsP(__e, let__14060)

var ifres15286 Obj

if True == tmp15366 {
tmp15287 := GHd(__e, let__14060)

tmp15288 := Call2(__e, SymFn(symshen_4lazyderef), tmp15287, V5020)


let__14061 := tmp15288
_ = let__14061

tmp15365 := GEq(__e, let__14061, sym_8p)

var ifres15289 Obj

if True == tmp15365 {
tmp15290 := GTl(__e, let__14060)

tmp15291 := Call2(__e, SymFn(symshen_4lazyderef), tmp15290, V5020)


let__14062 := tmp15291
_ = let__14062

tmp15364 := GConsP(__e, let__14062)

var ifres15292 Obj

if True == tmp15364 {
tmp15293 := GHd(__e, let__14062)

let__14063 := tmp15293
_ = let__14063

tmp15294 := GTl(__e, let__14062)

tmp15295 := Call2(__e, SymFn(symshen_4lazyderef), tmp15294, V5020)


let__14064 := tmp15295
_ = let__14064

tmp15363 := GConsP(__e, let__14064)

var ifres15296 Obj

if True == tmp15363 {
tmp15297 := GHd(__e, let__14064)

let__14065 := tmp15297
_ = let__14065

tmp15298 := GTl(__e, let__14064)

tmp15299 := Call2(__e, SymFn(symshen_4lazyderef), tmp15298, V5020)


let__14066 := tmp15299
_ = let__14066

tmp15362 := GEq(__e, let__14066, Nil)

var ifres15300 Obj

if True == tmp15362 {
tmp15301 := Call2(__e, SymFn(symshen_4lazyderef), V5018, V5020)


let__14067 := tmp15301
_ = let__14067

tmp15302 := MakeNative1(func(__e *ControlFlow, Z5089 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5090 Obj) Obj {
tmp15303 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15303

tmp15304 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), let__14065, Z5090, V5019, V5020, V5021, let__14006, V5023)
return nil
panic("unreachable")
})

__e.TailApply(SymFn(symshen_4system_1S_1h), let__14063, Z5089, V5019, V5020, V5021, let__14006, tmp15304)
return nil


panic("unreachable")
})
panic("unreachable")
})

let__14068 := tmp15302
_ = let__14068

tmp15361 := GConsP(__e, let__14067)

var ifres15305 Obj

if True == tmp15361 {
tmp15306 := GHd(__e, let__14067)

let__14069 := tmp15306
_ = let__14069

tmp15307 := GTl(__e, let__14067)

tmp15308 := Call2(__e, SymFn(symshen_4lazyderef), tmp15307, V5020)


let__14070 := tmp15308
_ = let__14070

tmp15309 := MakeNative1(func(__e *ControlFlow, Z5094 Obj) Obj {
tmp15310 := Call1(__e, let__14068, let__14069)


return TailCall1(__e, tmp15310, Z5094)


panic("unreachable")
})

let__14071 := tmp15309
_ = let__14071

tmp15348 := GConsP(__e, let__14070)

var ifres15311 Obj

if True == tmp15348 {
tmp15312 := GHd(__e, let__14070)

tmp15313 := Call2(__e, SymFn(symshen_4lazyderef), tmp15312, V5020)


let__14072 := tmp15313
_ = let__14072

tmp15314 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15315 := GTl(__e, let__14070)

tmp15316 := Call2(__e, SymFn(symshen_4lazyderef), tmp15315, V5020)


let__14074 := tmp15316
_ = let__14074

tmp15317 := MakeNative1(func(__e *ControlFlow, Z5099 Obj) Obj {
return TailCall1(__e, let__14071, Z5099)
panic("unreachable")
})

let__14075 := tmp15317
_ = let__14075

tmp15333 := GConsP(__e, let__14074)

if True == tmp15333 {
tmp15318 := GHd(__e, let__14074)

let__14076 := tmp15318
_ = let__14076

tmp15319 := GTl(__e, let__14074)

tmp15320 := Call2(__e, SymFn(symshen_4lazyderef), tmp15319, V5020)


let__14077 := tmp15320
_ = let__14077

tmp15321 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__14075, let__14076)
panic("unreachable")
})

let__14078 := tmp15321
_ = let__14078

tmp15325 := GEq(__e, let__14077, Nil)

if True == tmp15325 {
return TailCall1(__e, SymFn(symthaw), let__14078)
} else {
tmp15323 := Call1(__e, SymFn(symshen_4pvar_2), let__14077)


if True == tmp15323 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__14077, Nil, V5020, let__14078)
} else {
return False
}


}


} else {
tmp15331 := Call1(__e, SymFn(symshen_4pvar_2), let__14074)


if True == tmp15331 {
tmp15326 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__14079 := tmp15326
_ = let__14079

tmp15327 := GCons(__e, let__14079, Nil)

tmp15328 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__14075, let__14079)
panic("unreachable")
})

tmp15329 := Call4(__e, SymFn(symshen_4bind_b), let__14074, tmp15327, V5020, tmp15328)


return TailCall2(__e, SymFn(symshen_4gc), V5020, tmp15329)


} else {
return False
}


}


panic("unreachable")
})

let__14073 := tmp15314
_ = let__14073

tmp15339 := GEq(__e, let__14072, sym_d)

var ifres15334 Obj

if True == tmp15339 {
tmp15335 := Call1(__e, SymFn(symthaw), let__14073)


ifres15334 = tmp15335


} else {
tmp15338 := Call1(__e, SymFn(symshen_4pvar_2), let__14072)


var ifres15336 Obj

if True == tmp15338 {
tmp15337 := Call4(__e, SymFn(symshen_4bind_b), let__14072, sym_d, V5020, let__14073)


ifres15336 = tmp15337


} else {
ifres15336 = False


}

ifres15334 = ifres15336


}

ifres15311 = ifres15334


} else {
tmp15347 := Call1(__e, SymFn(symshen_4pvar_2), let__14070)


var ifres15340 Obj

if True == tmp15347 {
tmp15341 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__14080 := tmp15341
_ = let__14080

tmp15342 := GCons(__e, let__14080, Nil)

tmp15343 := GCons(__e, sym_d, tmp15342)

tmp15344 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__14071, let__14080)
panic("unreachable")
})

tmp15345 := Call4(__e, SymFn(symshen_4bind_b), let__14070, tmp15343, V5020, tmp15344)


tmp15346 := Call2(__e, SymFn(symshen_4gc), V5020, tmp15345)


ifres15340 = tmp15346


} else {
ifres15340 = False


}

ifres15311 = ifres15340


}

ifres15305 = ifres15311


} else {
tmp15360 := Call1(__e, SymFn(symshen_4pvar_2), let__14067)


var ifres15349 Obj

if True == tmp15360 {
tmp15350 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__14081 := tmp15350
_ = let__14081

tmp15351 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__14082 := tmp15351
_ = let__14082

tmp15352 := GCons(__e, let__14082, Nil)

tmp15353 := GCons(__e, sym_d, tmp15352)

tmp15354 := GCons(__e, let__14081, tmp15353)

tmp15355 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15356 := Call1(__e, let__14068, let__14081)


return TailCall1(__e, tmp15356, let__14082)


panic("unreachable")
})

tmp15357 := Call4(__e, SymFn(symshen_4bind_b), let__14067, tmp15354, V5020, tmp15355)


tmp15358 := Call2(__e, SymFn(symshen_4gc), V5020, tmp15357)


tmp15359 := Call2(__e, SymFn(symshen_4gc), V5020, tmp15358)


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

tmp15815 := GEq(__e, let__14059, False)

if True == tmp15815 {
tmp15434 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres15368 Obj

if True == tmp15434 {
tmp15369 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__14084 := tmp15369
_ = let__14084

tmp15433 := GConsP(__e, let__14084)

var ifres15370 Obj

if True == tmp15433 {
tmp15371 := GHd(__e, let__14084)

tmp15372 := Call2(__e, SymFn(symshen_4lazyderef), tmp15371, V5020)


let__14085 := tmp15372
_ = let__14085

tmp15432 := GEq(__e, let__14085, sym_8v)

var ifres15373 Obj

if True == tmp15432 {
tmp15374 := GTl(__e, let__14084)

tmp15375 := Call2(__e, SymFn(symshen_4lazyderef), tmp15374, V5020)


let__14086 := tmp15375
_ = let__14086

tmp15431 := GConsP(__e, let__14086)

var ifres15376 Obj

if True == tmp15431 {
tmp15377 := GHd(__e, let__14086)

let__14087 := tmp15377
_ = let__14087

tmp15378 := GTl(__e, let__14086)

tmp15379 := Call2(__e, SymFn(symshen_4lazyderef), tmp15378, V5020)


let__14088 := tmp15379
_ = let__14088

tmp15430 := GConsP(__e, let__14088)

var ifres15380 Obj

if True == tmp15430 {
tmp15381 := GHd(__e, let__14088)

let__14089 := tmp15381
_ = let__14089

tmp15382 := GTl(__e, let__14088)

tmp15383 := Call2(__e, SymFn(symshen_4lazyderef), tmp15382, V5020)


let__14090 := tmp15383
_ = let__14090

tmp15429 := GEq(__e, let__14090, Nil)

var ifres15384 Obj

if True == tmp15429 {
tmp15385 := Call2(__e, SymFn(symshen_4lazyderef), V5018, V5020)


let__14091 := tmp15385
_ = let__14091

tmp15386 := MakeNative1(func(__e *ControlFlow, Z5117 Obj) Obj {
tmp15387 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15387

tmp15388 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15389 := GCons(__e, Z5117, Nil)

tmp15390 := GCons(__e, symvector, tmp15389)

__e.TailApply(SymFn(symshen_4system_1S_1h), let__14089, tmp15390, V5019, V5020, V5021, let__14006, V5023)
return nil


panic("unreachable")
})

__e.TailApply(SymFn(symshen_4system_1S_1h), let__14087, Z5117, V5019, V5020, V5021, let__14006, tmp15388)
return nil


panic("unreachable")
})

let__14092 := tmp15386
_ = let__14092

tmp15428 := GConsP(__e, let__14091)

var ifres15391 Obj

if True == tmp15428 {
tmp15392 := GHd(__e, let__14091)

tmp15393 := Call2(__e, SymFn(symshen_4lazyderef), tmp15392, V5020)


let__14093 := tmp15393
_ = let__14093

tmp15394 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15395 := GTl(__e, let__14091)

tmp15396 := Call2(__e, SymFn(symshen_4lazyderef), tmp15395, V5020)


let__14095 := tmp15396
_ = let__14095

tmp15397 := MakeNative1(func(__e *ControlFlow, Z5122 Obj) Obj {
return TailCall1(__e, let__14092, Z5122)
panic("unreachable")
})

let__14096 := tmp15397
_ = let__14096

tmp15413 := GConsP(__e, let__14095)

if True == tmp15413 {
tmp15398 := GHd(__e, let__14095)

let__14097 := tmp15398
_ = let__14097

tmp15399 := GTl(__e, let__14095)

tmp15400 := Call2(__e, SymFn(symshen_4lazyderef), tmp15399, V5020)


let__14098 := tmp15400
_ = let__14098

tmp15401 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__14096, let__14097)
panic("unreachable")
})

let__14099 := tmp15401
_ = let__14099

tmp15405 := GEq(__e, let__14098, Nil)

if True == tmp15405 {
return TailCall1(__e, SymFn(symthaw), let__14099)
} else {
tmp15403 := Call1(__e, SymFn(symshen_4pvar_2), let__14098)


if True == tmp15403 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__14098, Nil, V5020, let__14099)
} else {
return False
}


}


} else {
tmp15411 := Call1(__e, SymFn(symshen_4pvar_2), let__14095)


if True == tmp15411 {
tmp15406 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__14100 := tmp15406
_ = let__14100

tmp15407 := GCons(__e, let__14100, Nil)

tmp15408 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__14096, let__14100)
panic("unreachable")
})

tmp15409 := Call4(__e, SymFn(symshen_4bind_b), let__14095, tmp15407, V5020, tmp15408)


return TailCall2(__e, SymFn(symshen_4gc), V5020, tmp15409)


} else {
return False
}


}


panic("unreachable")
})

let__14094 := tmp15394
_ = let__14094

tmp15419 := GEq(__e, let__14093, symvector)

var ifres15414 Obj

if True == tmp15419 {
tmp15415 := Call1(__e, SymFn(symthaw), let__14094)


ifres15414 = tmp15415


} else {
tmp15418 := Call1(__e, SymFn(symshen_4pvar_2), let__14093)


var ifres15416 Obj

if True == tmp15418 {
tmp15417 := Call4(__e, SymFn(symshen_4bind_b), let__14093, symvector, V5020, let__14094)


ifres15416 = tmp15417


} else {
ifres15416 = False


}

ifres15414 = ifres15416


}

ifres15391 = ifres15414


} else {
tmp15427 := Call1(__e, SymFn(symshen_4pvar_2), let__14091)


var ifres15420 Obj

if True == tmp15427 {
tmp15421 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__14101 := tmp15421
_ = let__14101

tmp15422 := GCons(__e, let__14101, Nil)

tmp15423 := GCons(__e, symvector, tmp15422)

tmp15424 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__14092, let__14101)
panic("unreachable")
})

tmp15425 := Call4(__e, SymFn(symshen_4bind_b), let__14091, tmp15423, V5020, tmp15424)


tmp15426 := Call2(__e, SymFn(symshen_4gc), V5020, tmp15425)


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

tmp15813 := GEq(__e, let__14083, False)

if True == tmp15813 {
tmp15467 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres15435 Obj

if True == tmp15467 {
tmp15436 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__14103 := tmp15436
_ = let__14103

tmp15466 := GConsP(__e, let__14103)

var ifres15437 Obj

if True == tmp15466 {
tmp15438 := GHd(__e, let__14103)

tmp15439 := Call2(__e, SymFn(symshen_4lazyderef), tmp15438, V5020)


let__14104 := tmp15439
_ = let__14104

tmp15465 := GEq(__e, let__14104, sym_8s)

var ifres15440 Obj

if True == tmp15465 {
tmp15441 := GTl(__e, let__14103)

tmp15442 := Call2(__e, SymFn(symshen_4lazyderef), tmp15441, V5020)


let__14105 := tmp15442
_ = let__14105

tmp15464 := GConsP(__e, let__14105)

var ifres15443 Obj

if True == tmp15464 {
tmp15444 := GHd(__e, let__14105)

let__14106 := tmp15444
_ = let__14106

tmp15445 := GTl(__e, let__14105)

tmp15446 := Call2(__e, SymFn(symshen_4lazyderef), tmp15445, V5020)


let__14107 := tmp15446
_ = let__14107

tmp15463 := GConsP(__e, let__14107)

var ifres15447 Obj

if True == tmp15463 {
tmp15448 := GHd(__e, let__14107)

let__14108 := tmp15448
_ = let__14108

tmp15449 := GTl(__e, let__14107)

tmp15450 := Call2(__e, SymFn(symshen_4lazyderef), tmp15449, V5020)


let__14109 := tmp15450
_ = let__14109

tmp15462 := GEq(__e, let__14109, Nil)

var ifres15451 Obj

if True == tmp15462 {
tmp15452 := Call2(__e, SymFn(symshen_4lazyderef), V5018, V5020)


let__14110 := tmp15452
_ = let__14110

tmp15453 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15454 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15454

tmp15455 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), let__14108, symstring, V5019, V5020, V5021, let__14006, V5023)
return nil
panic("unreachable")
})

__e.TailApply(SymFn(symshen_4system_1S_1h), let__14106, symstring, V5019, V5020, V5021, let__14006, tmp15455)
return nil


panic("unreachable")
})

let__14111 := tmp15453
_ = let__14111

tmp15461 := GEq(__e, let__14110, symstring)

var ifres15456 Obj

if True == tmp15461 {
tmp15457 := Call1(__e, SymFn(symthaw), let__14111)


ifres15456 = tmp15457


} else {
tmp15460 := Call1(__e, SymFn(symshen_4pvar_2), let__14110)


var ifres15458 Obj

if True == tmp15460 {
tmp15459 := Call4(__e, SymFn(symshen_4bind_b), let__14110, symstring, V5020, let__14111)


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

tmp15811 := GEq(__e, let__14102, False)

if True == tmp15811 {
tmp15567 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres15468 Obj

if True == tmp15567 {
tmp15469 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__14113 := tmp15469
_ = let__14113

tmp15566 := GConsP(__e, let__14113)

var ifres15470 Obj

if True == tmp15566 {
tmp15471 := GHd(__e, let__14113)

tmp15472 := Call2(__e, SymFn(symshen_4lazyderef), tmp15471, V5020)


let__14114 := tmp15472
_ = let__14114

tmp15565 := GEq(__e, let__14114, symlambda)

var ifres15473 Obj

if True == tmp15565 {
tmp15474 := GTl(__e, let__14113)

tmp15475 := Call2(__e, SymFn(symshen_4lazyderef), tmp15474, V5020)


let__14115 := tmp15475
_ = let__14115

tmp15564 := GConsP(__e, let__14115)

var ifres15476 Obj

if True == tmp15564 {
tmp15477 := GHd(__e, let__14115)

let__14116 := tmp15477
_ = let__14116

tmp15478 := GTl(__e, let__14115)

tmp15479 := Call2(__e, SymFn(symshen_4lazyderef), tmp15478, V5020)


let__14117 := tmp15479
_ = let__14117

tmp15563 := GConsP(__e, let__14117)

var ifres15480 Obj

if True == tmp15563 {
tmp15481 := GHd(__e, let__14117)

let__14118 := tmp15481
_ = let__14118

tmp15482 := GTl(__e, let__14117)

tmp15483 := Call2(__e, SymFn(symshen_4lazyderef), tmp15482, V5020)


let__14119 := tmp15483
_ = let__14119

tmp15562 := GEq(__e, let__14119, Nil)

var ifres15484 Obj

if True == tmp15562 {
tmp15485 := Call2(__e, SymFn(symshen_4lazyderef), V5018, V5020)


let__14120 := tmp15485
_ = let__14120

tmp15486 := MakeNative1(func(__e *ControlFlow, Z5148 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5149 Obj) Obj {
tmp15487 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__14122 := tmp15487
_ = let__14122

tmp15488 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__14123 := tmp15488
_ = let__14123

tmp15489 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15489

tmp15490 := Call2(__e, SymFn(symshen_4lazyderef), let__14116, V5020)


tmp15491 := Call1(__e, SymFn(symshen_4freshterm), tmp15490)


tmp15492 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15493 := Call2(__e, SymFn(symshen_4lazyderef), let__14116, V5020)


tmp15494 := Call2(__e, SymFn(symshen_4deref), let__14123, V5020)


tmp15495 := Call2(__e, SymFn(symshen_4deref), let__14118, V5020)


tmp15496 := Call3(__e, SymFn(symshen_4beta), tmp15493, tmp15494, tmp15495)


tmp15497 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15498 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp15499 := GCons(__e, Z5148, Nil)

tmp15500 := GCons(__e, tmp15498, tmp15499)

tmp15501 := GCons(__e, let__14123, tmp15500)

tmp15502 := GCons(__e, tmp15501, V5019)

__e.TailApply(SymFn(symshen_4system_1S_1h), let__14122, Z5149, tmp15502, V5020, V5021, let__14006, V5023)
return nil


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), let__14122, tmp15496, V5020, V5021, let__14006, tmp15497)


panic("unreachable")
})

tmp15503 := Call6(__e, SymFn(symbind), let__14123, tmp15491, V5020, V5021, let__14006, tmp15492)


tmp15504 := Call2(__e, SymFn(symshen_4gc), V5020, tmp15503)


return TailCall2(__e, SymFn(symshen_4gc), V5020, tmp15504)


panic("unreachable")
})
panic("unreachable")
})

let__14121 := tmp15486
_ = let__14121

tmp15561 := GConsP(__e, let__14120)

var ifres15505 Obj

if True == tmp15561 {
tmp15506 := GHd(__e, let__14120)

let__14124 := tmp15506
_ = let__14124

tmp15507 := GTl(__e, let__14120)

tmp15508 := Call2(__e, SymFn(symshen_4lazyderef), tmp15507, V5020)


let__14125 := tmp15508
_ = let__14125

tmp15509 := MakeNative1(func(__e *ControlFlow, Z5155 Obj) Obj {
tmp15510 := Call1(__e, let__14121, let__14124)


return TailCall1(__e, tmp15510, Z5155)


panic("unreachable")
})

let__14126 := tmp15509
_ = let__14126

tmp15548 := GConsP(__e, let__14125)

var ifres15511 Obj

if True == tmp15548 {
tmp15512 := GHd(__e, let__14125)

tmp15513 := Call2(__e, SymFn(symshen_4lazyderef), tmp15512, V5020)


let__14127 := tmp15513
_ = let__14127

tmp15514 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15515 := GTl(__e, let__14125)

tmp15516 := Call2(__e, SymFn(symshen_4lazyderef), tmp15515, V5020)


let__14129 := tmp15516
_ = let__14129

tmp15517 := MakeNative1(func(__e *ControlFlow, Z5160 Obj) Obj {
return TailCall1(__e, let__14126, Z5160)
panic("unreachable")
})

let__14130 := tmp15517
_ = let__14130

tmp15533 := GConsP(__e, let__14129)

if True == tmp15533 {
tmp15518 := GHd(__e, let__14129)

let__14131 := tmp15518
_ = let__14131

tmp15519 := GTl(__e, let__14129)

tmp15520 := Call2(__e, SymFn(symshen_4lazyderef), tmp15519, V5020)


let__14132 := tmp15520
_ = let__14132

tmp15521 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__14130, let__14131)
panic("unreachable")
})

let__14133 := tmp15521
_ = let__14133

tmp15525 := GEq(__e, let__14132, Nil)

if True == tmp15525 {
return TailCall1(__e, SymFn(symthaw), let__14133)
} else {
tmp15523 := Call1(__e, SymFn(symshen_4pvar_2), let__14132)


if True == tmp15523 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__14132, Nil, V5020, let__14133)
} else {
return False
}


}


} else {
tmp15531 := Call1(__e, SymFn(symshen_4pvar_2), let__14129)


if True == tmp15531 {
tmp15526 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__14134 := tmp15526
_ = let__14134

tmp15527 := GCons(__e, let__14134, Nil)

tmp15528 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__14130, let__14134)
panic("unreachable")
})

tmp15529 := Call4(__e, SymFn(symshen_4bind_b), let__14129, tmp15527, V5020, tmp15528)


return TailCall2(__e, SymFn(symshen_4gc), V5020, tmp15529)


} else {
return False
}


}


panic("unreachable")
})

let__14128 := tmp15514
_ = let__14128

tmp15539 := GEq(__e, let__14127, sym_1_1_6)

var ifres15534 Obj

if True == tmp15539 {
tmp15535 := Call1(__e, SymFn(symthaw), let__14128)


ifres15534 = tmp15535


} else {
tmp15538 := Call1(__e, SymFn(symshen_4pvar_2), let__14127)


var ifres15536 Obj

if True == tmp15538 {
tmp15537 := Call4(__e, SymFn(symshen_4bind_b), let__14127, sym_1_1_6, V5020, let__14128)


ifres15536 = tmp15537


} else {
ifres15536 = False


}

ifres15534 = ifres15536


}

ifres15511 = ifres15534


} else {
tmp15547 := Call1(__e, SymFn(symshen_4pvar_2), let__14125)


var ifres15540 Obj

if True == tmp15547 {
tmp15541 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__14135 := tmp15541
_ = let__14135

tmp15542 := GCons(__e, let__14135, Nil)

tmp15543 := GCons(__e, sym_1_1_6, tmp15542)

tmp15544 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__14126, let__14135)
panic("unreachable")
})

tmp15545 := Call4(__e, SymFn(symshen_4bind_b), let__14125, tmp15543, V5020, tmp15544)


tmp15546 := Call2(__e, SymFn(symshen_4gc), V5020, tmp15545)


ifres15540 = tmp15546


} else {
ifres15540 = False


}

ifres15511 = ifres15540


}

ifres15505 = ifres15511


} else {
tmp15560 := Call1(__e, SymFn(symshen_4pvar_2), let__14120)


var ifres15549 Obj

if True == tmp15560 {
tmp15550 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__14136 := tmp15550
_ = let__14136

tmp15551 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__14137 := tmp15551
_ = let__14137

tmp15552 := GCons(__e, let__14137, Nil)

tmp15553 := GCons(__e, sym_1_1_6, tmp15552)

tmp15554 := GCons(__e, let__14136, tmp15553)

tmp15555 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15556 := Call1(__e, let__14121, let__14136)


return TailCall1(__e, tmp15556, let__14137)


panic("unreachable")
})

tmp15557 := Call4(__e, SymFn(symshen_4bind_b), let__14120, tmp15554, V5020, tmp15555)


tmp15558 := Call2(__e, SymFn(symshen_4gc), V5020, tmp15557)


tmp15559 := Call2(__e, SymFn(symshen_4gc), V5020, tmp15558)


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

tmp15809 := GEq(__e, let__14112, False)

if True == tmp15809 {
tmp15617 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres15568 Obj

if True == tmp15617 {
tmp15569 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__14139 := tmp15569
_ = let__14139

tmp15616 := GConsP(__e, let__14139)

var ifres15570 Obj

if True == tmp15616 {
tmp15571 := GHd(__e, let__14139)

tmp15572 := Call2(__e, SymFn(symshen_4lazyderef), tmp15571, V5020)


let__14140 := tmp15572
_ = let__14140

tmp15615 := GEq(__e, let__14140, symlet)

var ifres15573 Obj

if True == tmp15615 {
tmp15574 := GTl(__e, let__14139)

tmp15575 := Call2(__e, SymFn(symshen_4lazyderef), tmp15574, V5020)


let__14141 := tmp15575
_ = let__14141

tmp15614 := GConsP(__e, let__14141)

var ifres15576 Obj

if True == tmp15614 {
tmp15577 := GHd(__e, let__14141)

let__14142 := tmp15577
_ = let__14142

tmp15578 := GTl(__e, let__14141)

tmp15579 := Call2(__e, SymFn(symshen_4lazyderef), tmp15578, V5020)


let__14143 := tmp15579
_ = let__14143

tmp15613 := GConsP(__e, let__14143)

var ifres15580 Obj

if True == tmp15613 {
tmp15581 := GHd(__e, let__14143)

let__14144 := tmp15581
_ = let__14144

tmp15582 := GTl(__e, let__14143)

tmp15583 := Call2(__e, SymFn(symshen_4lazyderef), tmp15582, V5020)


let__14145 := tmp15583
_ = let__14145

tmp15612 := GConsP(__e, let__14145)

var ifres15584 Obj

if True == tmp15612 {
tmp15585 := GHd(__e, let__14145)

let__14146 := tmp15585
_ = let__14146

tmp15586 := GTl(__e, let__14145)

tmp15587 := Call2(__e, SymFn(symshen_4lazyderef), tmp15586, V5020)


let__14147 := tmp15587
_ = let__14147

tmp15611 := GEq(__e, let__14147, Nil)

var ifres15588 Obj

if True == tmp15611 {
tmp15589 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__14148 := tmp15589
_ = let__14148

tmp15590 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__14149 := tmp15590
_ = let__14149

tmp15591 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__14150 := tmp15591
_ = let__14150

tmp15592 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15592

tmp15593 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15594 := Call2(__e, SymFn(symshen_4lazyderef), let__14142, V5020)


tmp15595 := Call1(__e, SymFn(symshen_4freshterm), tmp15594)


tmp15596 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15597 := Call2(__e, SymFn(symshen_4lazyderef), let__14142, V5020)


tmp15598 := Call2(__e, SymFn(symshen_4lazyderef), let__14149, V5020)


tmp15599 := Call2(__e, SymFn(symshen_4lazyderef), let__14146, V5020)


tmp15600 := Call3(__e, SymFn(symshen_4beta), tmp15597, tmp15598, tmp15599)


tmp15601 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15602 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp15603 := GCons(__e, let__14150, Nil)

tmp15604 := GCons(__e, tmp15602, tmp15603)

tmp15605 := GCons(__e, let__14149, tmp15604)

tmp15606 := GCons(__e, tmp15605, V5019)

__e.TailApply(SymFn(symshen_4system_1S_1h), let__14148, V5018, tmp15606, V5020, V5021, let__14006, V5023)
return nil


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), let__14148, tmp15600, V5020, V5021, let__14006, tmp15601)


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), let__14149, tmp15595, V5020, V5021, let__14006, tmp15596)


panic("unreachable")
})

tmp15607 := Call(__e, SymFn(symshen_4system_1S_1h), let__14144, let__14150, V5019, V5020, V5021, let__14006, tmp15593)


tmp15608 := Call2(__e, SymFn(symshen_4gc), V5020, tmp15607)


tmp15609 := Call2(__e, SymFn(symshen_4gc), V5020, tmp15608)


tmp15610 := Call2(__e, SymFn(symshen_4gc), V5020, tmp15609)


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

tmp15807 := GEq(__e, let__14138, False)

if True == tmp15807 {
tmp15687 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres15618 Obj

if True == tmp15687 {
tmp15619 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__14152 := tmp15619
_ = let__14152

tmp15686 := GConsP(__e, let__14152)

var ifres15620 Obj

if True == tmp15686 {
tmp15621 := GHd(__e, let__14152)

tmp15622 := Call2(__e, SymFn(symshen_4lazyderef), tmp15621, V5020)


let__14153 := tmp15622
_ = let__14153

tmp15685 := GEq(__e, let__14153, symopen)

var ifres15623 Obj

if True == tmp15685 {
tmp15624 := GTl(__e, let__14152)

tmp15625 := Call2(__e, SymFn(symshen_4lazyderef), tmp15624, V5020)


let__14154 := tmp15625
_ = let__14154

tmp15684 := GConsP(__e, let__14154)

var ifres15626 Obj

if True == tmp15684 {
tmp15627 := GHd(__e, let__14154)

let__14155 := tmp15627
_ = let__14155

tmp15628 := GTl(__e, let__14154)

tmp15629 := Call2(__e, SymFn(symshen_4lazyderef), tmp15628, V5020)


let__14156 := tmp15629
_ = let__14156

tmp15683 := GConsP(__e, let__14156)

var ifres15630 Obj

if True == tmp15683 {
tmp15631 := GHd(__e, let__14156)

let__14157 := tmp15631
_ = let__14157

tmp15632 := GTl(__e, let__14156)

tmp15633 := Call2(__e, SymFn(symshen_4lazyderef), tmp15632, V5020)


let__14158 := tmp15633
_ = let__14158

tmp15682 := GEq(__e, let__14158, Nil)

var ifres15634 Obj

if True == tmp15682 {
tmp15635 := Call2(__e, SymFn(symshen_4lazyderef), V5018, V5020)


let__14159 := tmp15635
_ = let__14159

tmp15636 := MakeNative1(func(__e *ControlFlow, Z5191 Obj) Obj {
tmp15637 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15637

tmp15638 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15639 := Call2(__e, SymFn(symshen_4lazyderef), Z5191, V5020)


tmp15640 := GCons(__e, symout, Nil)

tmp15641 := GCons(__e, symin, tmp15640)

tmp15642 := Call2(__e, SymFn(symelement_2), tmp15639, tmp15641)


tmp15643 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), let__14155, symstring, V5019, V5020, V5021, let__14006, V5023)
return nil
panic("unreachable")
})

return TailCall5(__e, SymFn(symwhen), tmp15642, V5020, V5021, let__14006, tmp15643)


panic("unreachable")
})

return TailCall6(__e, SymFn(symis_b), let__14157, Z5191, V5020, V5021, let__14006, tmp15638)


panic("unreachable")
})

let__14160 := tmp15636
_ = let__14160

tmp15681 := GConsP(__e, let__14159)

var ifres15644 Obj

if True == tmp15681 {
tmp15645 := GHd(__e, let__14159)

tmp15646 := Call2(__e, SymFn(symshen_4lazyderef), tmp15645, V5020)


let__14161 := tmp15646
_ = let__14161

tmp15647 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15648 := GTl(__e, let__14159)

tmp15649 := Call2(__e, SymFn(symshen_4lazyderef), tmp15648, V5020)


let__14163 := tmp15649
_ = let__14163

tmp15650 := MakeNative1(func(__e *ControlFlow, Z5196 Obj) Obj {
return TailCall1(__e, let__14160, Z5196)
panic("unreachable")
})

let__14164 := tmp15650
_ = let__14164

tmp15666 := GConsP(__e, let__14163)

if True == tmp15666 {
tmp15651 := GHd(__e, let__14163)

let__14165 := tmp15651
_ = let__14165

tmp15652 := GTl(__e, let__14163)

tmp15653 := Call2(__e, SymFn(symshen_4lazyderef), tmp15652, V5020)


let__14166 := tmp15653
_ = let__14166

tmp15654 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__14164, let__14165)
panic("unreachable")
})

let__14167 := tmp15654
_ = let__14167

tmp15658 := GEq(__e, let__14166, Nil)

if True == tmp15658 {
return TailCall1(__e, SymFn(symthaw), let__14167)
} else {
tmp15656 := Call1(__e, SymFn(symshen_4pvar_2), let__14166)


if True == tmp15656 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__14166, Nil, V5020, let__14167)
} else {
return False
}


}


} else {
tmp15664 := Call1(__e, SymFn(symshen_4pvar_2), let__14163)


if True == tmp15664 {
tmp15659 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__14168 := tmp15659
_ = let__14168

tmp15660 := GCons(__e, let__14168, Nil)

tmp15661 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__14164, let__14168)
panic("unreachable")
})

tmp15662 := Call4(__e, SymFn(symshen_4bind_b), let__14163, tmp15660, V5020, tmp15661)


return TailCall2(__e, SymFn(symshen_4gc), V5020, tmp15662)


} else {
return False
}


}


panic("unreachable")
})

let__14162 := tmp15647
_ = let__14162

tmp15672 := GEq(__e, let__14161, symstream)

var ifres15667 Obj

if True == tmp15672 {
tmp15668 := Call1(__e, SymFn(symthaw), let__14162)


ifres15667 = tmp15668


} else {
tmp15671 := Call1(__e, SymFn(symshen_4pvar_2), let__14161)


var ifres15669 Obj

if True == tmp15671 {
tmp15670 := Call4(__e, SymFn(symshen_4bind_b), let__14161, symstream, V5020, let__14162)


ifres15669 = tmp15670


} else {
ifres15669 = False


}

ifres15667 = ifres15669


}

ifres15644 = ifres15667


} else {
tmp15680 := Call1(__e, SymFn(symshen_4pvar_2), let__14159)


var ifres15673 Obj

if True == tmp15680 {
tmp15674 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__14169 := tmp15674
_ = let__14169

tmp15675 := GCons(__e, let__14169, Nil)

tmp15676 := GCons(__e, symstream, tmp15675)

tmp15677 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__14160, let__14169)
panic("unreachable")
})

tmp15678 := Call4(__e, SymFn(symshen_4bind_b), let__14159, tmp15676, V5020, tmp15677)


tmp15679 := Call2(__e, SymFn(symshen_4gc), V5020, tmp15678)


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

tmp15805 := GEq(__e, let__14151, False)

if True == tmp15805 {
tmp15716 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres15688 Obj

if True == tmp15716 {
tmp15689 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__14171 := tmp15689
_ = let__14171

tmp15715 := GConsP(__e, let__14171)

var ifres15690 Obj

if True == tmp15715 {
tmp15691 := GHd(__e, let__14171)

tmp15692 := Call2(__e, SymFn(symshen_4lazyderef), tmp15691, V5020)


let__14172 := tmp15692
_ = let__14172

tmp15714 := GEq(__e, let__14172, symtype)

var ifres15693 Obj

if True == tmp15714 {
tmp15694 := GTl(__e, let__14171)

tmp15695 := Call2(__e, SymFn(symshen_4lazyderef), tmp15694, V5020)


let__14173 := tmp15695
_ = let__14173

tmp15713 := GConsP(__e, let__14173)

var ifres15696 Obj

if True == tmp15713 {
tmp15697 := GHd(__e, let__14173)

let__14174 := tmp15697
_ = let__14174

tmp15698 := GTl(__e, let__14173)

tmp15699 := Call2(__e, SymFn(symshen_4lazyderef), tmp15698, V5020)


let__14175 := tmp15699
_ = let__14175

tmp15712 := GConsP(__e, let__14175)

var ifres15700 Obj

if True == tmp15712 {
tmp15701 := GHd(__e, let__14175)

let__14176 := tmp15701
_ = let__14176

tmp15702 := GTl(__e, let__14175)

tmp15703 := Call2(__e, SymFn(symshen_4lazyderef), tmp15702, V5020)


let__14177 := tmp15703
_ = let__14177

tmp15711 := GEq(__e, let__14177, Nil)

var ifres15704 Obj

if True == tmp15711 {
tmp15705 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15705

tmp15706 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15707 := Call2(__e, SymFn(symshen_4deref), let__14176, V5020)


tmp15708 := Call1(__e, SymFn(symshen_4rectify_1type), tmp15707)


tmp15709 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), let__14174, V5018, V5019, V5020, V5021, let__14006, V5023)
return nil
panic("unreachable")
})

return TailCall6(__e, SymFn(symis_b), tmp15708, V5018, V5020, V5021, let__14006, tmp15709)


panic("unreachable")
})

tmp15710 := Call4(__e, SymFn(symshen_4cut), V5020, V5021, let__14006, tmp15706)


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

tmp15803 := GEq(__e, let__14170, False)

if True == tmp15803 {
tmp15747 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres15717 Obj

if True == tmp15747 {
tmp15718 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__14179 := tmp15718
_ = let__14179

tmp15746 := GConsP(__e, let__14179)

var ifres15719 Obj

if True == tmp15746 {
tmp15720 := GHd(__e, let__14179)

tmp15721 := Call2(__e, SymFn(symshen_4lazyderef), tmp15720, V5020)


let__14180 := tmp15721
_ = let__14180

tmp15745 := GEq(__e, let__14180, symshen_4input_1h_7)

var ifres15722 Obj

if True == tmp15745 {
tmp15723 := GTl(__e, let__14179)

tmp15724 := Call2(__e, SymFn(symshen_4lazyderef), tmp15723, V5020)


let__14181 := tmp15724
_ = let__14181

tmp15744 := GConsP(__e, let__14181)

var ifres15725 Obj

if True == tmp15744 {
tmp15726 := GHd(__e, let__14181)

let__14182 := tmp15726
_ = let__14182

tmp15727 := GTl(__e, let__14181)

tmp15728 := Call2(__e, SymFn(symshen_4lazyderef), tmp15727, V5020)


let__14183 := tmp15728
_ = let__14183

tmp15743 := GConsP(__e, let__14183)

var ifres15729 Obj

if True == tmp15743 {
tmp15730 := GHd(__e, let__14183)

let__14184 := tmp15730
_ = let__14184

tmp15731 := GTl(__e, let__14183)

tmp15732 := Call2(__e, SymFn(symshen_4lazyderef), tmp15731, V5020)


let__14185 := tmp15732
_ = let__14185

tmp15742 := GEq(__e, let__14185, Nil)

var ifres15733 Obj

if True == tmp15742 {
tmp15734 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15734

tmp15735 := Call2(__e, SymFn(symshen_4deref), let__14182, V5020)


tmp15736 := Call1(__e, SymFn(symshen_4rdecons), tmp15735)


tmp15737 := Call1(__e, SymFn(symshen_4rectify_1type), tmp15736)


tmp15738 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15739 := GCons(__e, symin, Nil)

tmp15740 := GCons(__e, symstream, tmp15739)

__e.TailApply(SymFn(symshen_4system_1S_1h), let__14184, tmp15740, V5019, V5020, V5021, let__14006, V5023)
return nil


panic("unreachable")
})

tmp15741 := Call6(__e, SymFn(symis), V5018, tmp15737, V5020, V5021, let__14006, tmp15738)


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

tmp15801 := GEq(__e, let__14178, False)

if True == tmp15801 {
tmp15776 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres15748 Obj

if True == tmp15776 {
tmp15749 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__14187 := tmp15749
_ = let__14187

tmp15775 := GConsP(__e, let__14187)

var ifres15750 Obj

if True == tmp15775 {
tmp15751 := GHd(__e, let__14187)

tmp15752 := Call2(__e, SymFn(symshen_4lazyderef), tmp15751, V5020)


let__14188 := tmp15752
_ = let__14188

tmp15774 := GEq(__e, let__14188, symset)

var ifres15753 Obj

if True == tmp15774 {
tmp15754 := GTl(__e, let__14187)

tmp15755 := Call2(__e, SymFn(symshen_4lazyderef), tmp15754, V5020)


let__14189 := tmp15755
_ = let__14189

tmp15773 := GConsP(__e, let__14189)

var ifres15756 Obj

if True == tmp15773 {
tmp15757 := GHd(__e, let__14189)

let__14190 := tmp15757
_ = let__14190

tmp15758 := GTl(__e, let__14189)

tmp15759 := Call2(__e, SymFn(symshen_4lazyderef), tmp15758, V5020)


let__14191 := tmp15759
_ = let__14191

tmp15772 := GConsP(__e, let__14191)

var ifres15760 Obj

if True == tmp15772 {
tmp15761 := GHd(__e, let__14191)

let__14192 := tmp15761
_ = let__14192

tmp15762 := GTl(__e, let__14191)

tmp15763 := Call2(__e, SymFn(symshen_4lazyderef), tmp15762, V5020)


let__14193 := tmp15763
_ = let__14193

tmp15771 := GEq(__e, let__14193, Nil)

var ifres15764 Obj

if True == tmp15771 {
tmp15765 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15765

tmp15766 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15767 := GCons(__e, let__14190, Nil)

tmp15768 := GCons(__e, symvalue, tmp15767)

tmp15769 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), let__14192, V5018, V5019, V5020, V5021, let__14006, V5023)
return nil
panic("unreachable")
})

__e.TailApply(SymFn(symshen_4system_1S_1h), tmp15768, V5018, V5019, V5020, V5021, let__14006, tmp15769)
return nil


panic("unreachable")
})

tmp15770 := Call(__e, SymFn(symshen_4system_1S_1h), let__14190, symsymbol, V5019, V5020, V5021, let__14006, tmp15766)


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

tmp15799 := GEq(__e, let__14186, False)

if True == tmp15799 {
tmp15784 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres15777 Obj

if True == tmp15784 {
tmp15778 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__14195 := tmp15778
_ = let__14195

tmp15779 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15779

tmp15780 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15781 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), V5017, V5018, let__14195, V5020, V5021, let__14006, V5023)
return nil
panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5020, V5021, let__14006, tmp15781)


panic("unreachable")
})

tmp15782 := Call(__e, SymFn(symshen_4l_1rules), V5019, let__14195, False, V5020, V5021, let__14006, tmp15780)


tmp15783 := Call2(__e, SymFn(symshen_4gc), V5020, tmp15782)


ifres15777 = tmp15783


} else {
ifres15777 = False


}

let__14194 := ifres15777
_ = let__14194

tmp15797 := GEq(__e, let__14194, False)

if True == tmp15797 {
tmp15793 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres15785 Obj

if True == tmp15793 {
tmp15786 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15786

tmp15787 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp15788 := GCons(__e, V5018, Nil)

tmp15789 := GCons(__e, tmp15787, tmp15788)

tmp15790 := GCons(__e, V5017, tmp15789)

tmp15791 := GValue(__e, symshen_4_ddatatypes_d)

tmp15792 := Call(__e, SymFn(symshen_4search_1user_1datatypes), tmp15790, V5019, tmp15791, V5020, V5021, let__14006, V5023)


ifres15785 = tmp15792


} else {
ifres15785 = False


}

let__14196 := ifres15785
_ = let__14196

tmp15795 := GEq(__e, let__14196, False)

if True == tmp15795 {
return TailCall2(__e, SymFn(symshen_4unlock), V5021, let__14006)
} else {
return let__14196
}


} else {
return let__14194
}


} else {
return let__14186
}


} else {
return let__14178
}


} else {
return let__14170
}


} else {
return let__14151
}


} else {
return let__14138
}


} else {
return let__14112
}


} else {
return let__14102
}


} else {
return let__14083
}


} else {
return let__14059
}


} else {
return let__14040
}


} else {
return let__14033
}


} else {
return let__14026
}


} else {
return let__14020
}


} else {
return let__14014
}


} else {
return let__14010
}


} else {
return let__14009
}


} else {
return let__14008
}


} else {
return let__14007
}


panic("unreachable")
}())
}, 7)

tmp15834 := Call2(__e, ns2_1set, symshen_4system_1S_1h, tmp15083)


_ = tmp15834

tmp15835 := MakeNative1(func(__e *ControlFlow, V5229 Obj) Obj {
tmp15869 := GConsP(__e, V5229)

var ifres15850 Obj

if True == tmp15869 {
tmp15867 := GHd(__e, V5229)

tmp15868 := GEq(__e, symcons, tmp15867)

var ifres15852 Obj

if True == tmp15868 {
tmp15865 := GTl(__e, V5229)

tmp15866 := GConsP(__e, tmp15865)

var ifres15854 Obj

if True == tmp15866 {
tmp15862 := GTl(__e, V5229)

tmp15863 := GTl(__e, tmp15862)

tmp15864 := GConsP(__e, tmp15863)

var ifres15856 Obj

if True == tmp15864 {
tmp15858 := GTl(__e, V5229)

tmp15859 := GTl(__e, tmp15858)

tmp15860 := GTl(__e, tmp15859)

tmp15861 := GEq(__e, Nil, tmp15860)

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
tmp15836 := GTl(__e, V5229)

tmp15837 := GHd(__e, tmp15836)

tmp15838 := Call1(__e, SymFn(symshen_4rdecons), tmp15837)


tmp15839 := GTl(__e, V5229)

tmp15840 := GTl(__e, tmp15839)

tmp15841 := GHd(__e, tmp15840)

tmp15842 := Call1(__e, SymFn(symshen_4rdecons), tmp15841)


return GCons(__e, tmp15838, tmp15842)


} else {
tmp15848 := GConsP(__e, V5229)

if True == tmp15848 {
tmp15843 := GHd(__e, V5229)

tmp15844 := Call1(__e, SymFn(symshen_4rdecons), tmp15843)


tmp15845 := GTl(__e, V5229)

tmp15846 := Call1(__e, SymFn(symshen_4rdecons), tmp15845)


return GCons(__e, tmp15844, tmp15846)


} else {
return V5229
}


}


panic("unreachable")
})

tmp15870 := Call2(__e, ns2_1set, symshen_4rdecons, tmp15835)


_ = tmp15870

tmp15871 := MakeNative6(func(__e *ControlFlow, V5230, V5231, V5232, V5233, V5234, V5235 Obj) Obj {
tmp15884 := Call1(__e, SymFn(symshen_4unlocked_2), V5233)


var ifres15872 Obj

if True == tmp15884 {
tmp15873 := Call2(__e, SymFn(symshen_4lazyderef), V5231, V5232)


let__14198 := tmp15873
_ = let__14198

tmp15874 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15875 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15875

tmp15876 := Call2(__e, SymFn(symshen_4lazyderef), V5230, V5232)


tmp15877 := PrimGuard1(__e, symnumber_2, PrimIsNumber, tmp15876)

return TailCall5(__e, SymFn(symwhen), tmp15877, V5232, V5233, V5234, V5235)


panic("unreachable")
})

let__14199 := tmp15874
_ = let__14199

tmp15883 := GEq(__e, let__14198, symnumber)

var ifres15878 Obj

if True == tmp15883 {
tmp15879 := Call1(__e, SymFn(symthaw), let__14199)


ifres15878 = tmp15879


} else {
tmp15882 := Call1(__e, SymFn(symshen_4pvar_2), let__14198)


var ifres15880 Obj

if True == tmp15882 {
tmp15881 := Call4(__e, SymFn(symshen_4bind_b), let__14198, symnumber, V5232, let__14199)


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

tmp15974 := GEq(__e, let__14197, False)

if True == tmp15974 {
tmp15897 := Call1(__e, SymFn(symshen_4unlocked_2), V5233)


var ifres15885 Obj

if True == tmp15897 {
tmp15886 := Call2(__e, SymFn(symshen_4lazyderef), V5231, V5232)


let__14201 := tmp15886
_ = let__14201

tmp15887 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15888 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15888

tmp15889 := Call2(__e, SymFn(symshen_4lazyderef), V5230, V5232)


tmp15890 := Call1(__e, SymFn(symboolean_2), tmp15889)


return TailCall5(__e, SymFn(symwhen), tmp15890, V5232, V5233, V5234, V5235)


panic("unreachable")
})

let__14202 := tmp15887
_ = let__14202

tmp15896 := GEq(__e, let__14201, symboolean)

var ifres15891 Obj

if True == tmp15896 {
tmp15892 := Call1(__e, SymFn(symthaw), let__14202)


ifres15891 = tmp15892


} else {
tmp15895 := Call1(__e, SymFn(symshen_4pvar_2), let__14201)


var ifres15893 Obj

if True == tmp15895 {
tmp15894 := Call4(__e, SymFn(symshen_4bind_b), let__14201, symboolean, V5232, let__14202)


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

tmp15972 := GEq(__e, let__14200, False)

if True == tmp15972 {
tmp15910 := Call1(__e, SymFn(symshen_4unlocked_2), V5233)


var ifres15898 Obj

if True == tmp15910 {
tmp15899 := Call2(__e, SymFn(symshen_4lazyderef), V5231, V5232)


let__14204 := tmp15899
_ = let__14204

tmp15900 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15901 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15901

tmp15902 := Call2(__e, SymFn(symshen_4lazyderef), V5230, V5232)


tmp15903 := PrimGuard1(__e, symstring_2, PrimIsString, tmp15902)

return TailCall5(__e, SymFn(symwhen), tmp15903, V5232, V5233, V5234, V5235)


panic("unreachable")
})

let__14205 := tmp15900
_ = let__14205

tmp15909 := GEq(__e, let__14204, symstring)

var ifres15904 Obj

if True == tmp15909 {
tmp15905 := Call1(__e, SymFn(symthaw), let__14205)


ifres15904 = tmp15905


} else {
tmp15908 := Call1(__e, SymFn(symshen_4pvar_2), let__14204)


var ifres15906 Obj

if True == tmp15908 {
tmp15907 := Call4(__e, SymFn(symshen_4bind_b), let__14204, symstring, V5232, let__14205)


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

tmp15970 := GEq(__e, let__14203, False)

if True == tmp15970 {
tmp15923 := Call1(__e, SymFn(symshen_4unlocked_2), V5233)


var ifres15911 Obj

if True == tmp15923 {
tmp15912 := Call2(__e, SymFn(symshen_4lazyderef), V5231, V5232)


let__14207 := tmp15912
_ = let__14207

tmp15913 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15914 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15914

tmp15915 := Call2(__e, SymFn(symshen_4lazyderef), V5230, V5232)


tmp15916 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, tmp15915)

return TailCall5(__e, SymFn(symwhen), tmp15916, V5232, V5233, V5234, V5235)


panic("unreachable")
})

let__14208 := tmp15913
_ = let__14208

tmp15922 := GEq(__e, let__14207, symsymbol)

var ifres15917 Obj

if True == tmp15922 {
tmp15918 := Call1(__e, SymFn(symthaw), let__14208)


ifres15917 = tmp15918


} else {
tmp15921 := Call1(__e, SymFn(symshen_4pvar_2), let__14207)


var ifres15919 Obj

if True == tmp15921 {
tmp15920 := Call4(__e, SymFn(symshen_4bind_b), let__14207, symsymbol, V5232, let__14208)


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

tmp15968 := GEq(__e, let__14206, False)

if True == tmp15968 {
tmp15966 := Call1(__e, SymFn(symshen_4unlocked_2), V5233)


if True == tmp15966 {
tmp15924 := Call2(__e, SymFn(symshen_4lazyderef), V5230, V5232)


let__14209 := tmp15924
_ = let__14209

tmp15964 := GEq(__e, let__14209, Nil)

if True == tmp15964 {
tmp15925 := Call2(__e, SymFn(symshen_4lazyderef), V5231, V5232)


let__14210 := tmp15925
_ = let__14210

tmp15926 := MakeNative1(func(__e *ControlFlow, Z5251 Obj) Obj {
tmp15927 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15927

return TailCall1(__e, SymFn(symthaw), V5235)


panic("unreachable")
})

let__14211 := tmp15926
_ = let__14211

tmp15962 := GConsP(__e, let__14210)

if True == tmp15962 {
tmp15928 := GHd(__e, let__14210)

tmp15929 := Call2(__e, SymFn(symshen_4lazyderef), tmp15928, V5232)


let__14212 := tmp15929
_ = let__14212

tmp15930 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15931 := GTl(__e, let__14210)

tmp15932 := Call2(__e, SymFn(symshen_4lazyderef), tmp15931, V5232)


let__14214 := tmp15932
_ = let__14214

tmp15933 := MakeNative1(func(__e *ControlFlow, Z5256 Obj) Obj {
return TailCall1(__e, let__14211, Z5256)
panic("unreachable")
})

let__14215 := tmp15933
_ = let__14215

tmp15949 := GConsP(__e, let__14214)

if True == tmp15949 {
tmp15934 := GHd(__e, let__14214)

let__14216 := tmp15934
_ = let__14216

tmp15935 := GTl(__e, let__14214)

tmp15936 := Call2(__e, SymFn(symshen_4lazyderef), tmp15935, V5232)


let__14217 := tmp15936
_ = let__14217

tmp15937 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__14215, let__14216)
panic("unreachable")
})

let__14218 := tmp15937
_ = let__14218

tmp15941 := GEq(__e, let__14217, Nil)

if True == tmp15941 {
return TailCall1(__e, SymFn(symthaw), let__14218)
} else {
tmp15939 := Call1(__e, SymFn(symshen_4pvar_2), let__14217)


if True == tmp15939 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__14217, Nil, V5232, let__14218)
} else {
return False
}


}


} else {
tmp15947 := Call1(__e, SymFn(symshen_4pvar_2), let__14214)


if True == tmp15947 {
tmp15942 := Call1(__e, SymFn(symshen_4newpv), V5232)


let__14219 := tmp15942
_ = let__14219

tmp15943 := GCons(__e, let__14219, Nil)

tmp15944 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__14215, let__14219)
panic("unreachable")
})

tmp15945 := Call4(__e, SymFn(symshen_4bind_b), let__14214, tmp15943, V5232, tmp15944)


return TailCall2(__e, SymFn(symshen_4gc), V5232, tmp15945)


} else {
return False
}


}


panic("unreachable")
})

let__14213 := tmp15930
_ = let__14213

tmp15953 := GEq(__e, let__14212, symlist)

if True == tmp15953 {
return TailCall1(__e, SymFn(symthaw), let__14213)
} else {
tmp15951 := Call1(__e, SymFn(symshen_4pvar_2), let__14212)


if True == tmp15951 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__14212, symlist, V5232, let__14213)
} else {
return False
}


}


} else {
tmp15960 := Call1(__e, SymFn(symshen_4pvar_2), let__14210)


if True == tmp15960 {
tmp15954 := Call1(__e, SymFn(symshen_4newpv), V5232)


let__14220 := tmp15954
_ = let__14220

tmp15955 := GCons(__e, let__14220, Nil)

tmp15956 := GCons(__e, symlist, tmp15955)

tmp15957 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__14211, let__14220)
panic("unreachable")
})

tmp15958 := Call4(__e, SymFn(symshen_4bind_b), let__14210, tmp15956, V5232, tmp15957)


return TailCall2(__e, SymFn(symshen_4gc), V5232, tmp15958)


} else {
return False
}


}


} else {
return False
}


} else {
return False
}


} else {
return let__14206
}


} else {
return let__14203
}


} else {
return let__14200
}


} else {
return let__14197
}


panic("unreachable")
})

tmp15975 := Call2(__e, ns2_1set, symshen_4primitive, tmp15871)


_ = tmp15975

tmp15976 := MakeNative(func(__e *ControlFlow) {
V5262 := __e.Get(1)
V5263 := __e.Get(2)
V5264 := __e.Get(3)
V5265 := __e.Get(4)
V5266 := __e.Get(5)
V5267 := __e.Get(6)
V5268 := __e.Get(7)
__e.ReturnDirect(func() Obj {
tmp15979 := Call2(__e, SymFn(symshen_4lazyderef), V5264, V5265)


tmp15980 := GEq(__e, tmp15979, Nil)

if True == tmp15980 {
return False
} else {
tmp15977 := Call2(__e, SymFn(symshen_4deref), V5262, V5265)


__e.TailApply(SymFn(symshen_4by_1hypothesis_1h), tmp15977, V5263, V5264, V5265, V5266, V5267, V5268)
return nil


}


panic("unreachable")
}())
}, 7)

tmp15981 := Call2(__e, ns2_1set, symshen_4by_1hypothesis, tmp15976)


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
__e.ReturnDirect(func() Obj {
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
tmp16015 := Call1(__e, SymFn(symshen_4unlocked_2), V5273)


var ifres15983 Obj

if True == tmp16015 {
tmp15984 := Call2(__e, SymFn(symshen_4lazyderef), V5271, V5272)


let__14222 := tmp15984
_ = let__14222

tmp16014 := GConsP(__e, let__14222)

var ifres15985 Obj

if True == tmp16014 {
tmp15986 := GHd(__e, let__14222)

tmp15987 := Call2(__e, SymFn(symshen_4lazyderef), tmp15986, V5272)


let__14223 := tmp15987
_ = let__14223

tmp16013 := GConsP(__e, let__14223)

var ifres15988 Obj

if True == tmp16013 {
tmp15989 := GHd(__e, let__14223)

let__14224 := tmp15989
_ = let__14224

tmp15990 := GTl(__e, let__14223)

tmp15991 := Call2(__e, SymFn(symshen_4lazyderef), tmp15990, V5272)


let__14225 := tmp15991
_ = let__14225

tmp16012 := GConsP(__e, let__14225)

var ifres15992 Obj

if True == tmp16012 {
tmp15993 := GHd(__e, let__14225)

let__14226 := tmp15993
_ = let__14226

tmp15994 := GTl(__e, let__14225)

tmp15995 := Call2(__e, SymFn(symshen_4lazyderef), tmp15994, V5272)


let__14227 := tmp15995
_ = let__14227

tmp16011 := GConsP(__e, let__14227)

var ifres15996 Obj

if True == tmp16011 {
tmp15997 := GHd(__e, let__14227)

let__14228 := tmp15997
_ = let__14228

tmp15998 := GTl(__e, let__14227)

tmp15999 := Call2(__e, SymFn(symshen_4lazyderef), tmp15998, V5272)


let__14229 := tmp15999
_ = let__14229

tmp16010 := GEq(__e, let__14229, Nil)

var ifres16000 Obj

if True == tmp16010 {
tmp16001 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16001

tmp16002 := Call2(__e, SymFn(symshen_4deref), let__14226, V5272)


tmp16003 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp16004 := GEq(__e, tmp16002, tmp16003)

tmp16005 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16006 := Call2(__e, SymFn(symshen_4deref), let__14224, V5272)


tmp16007 := GEq(__e, V5269, tmp16006)

tmp16008 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall6(__e, SymFn(symis_b), V5270, let__14228, V5272, V5273, V5274, V5275)
panic("unreachable")
})

return TailCall5(__e, SymFn(symwhen), tmp16007, V5272, V5273, V5274, tmp16008)


panic("unreachable")
})

tmp16009 := Call5(__e, SymFn(symwhen), tmp16004, V5272, V5273, V5274, tmp16005)


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

tmp16024 := GEq(__e, let__14221, False)

if True == tmp16024 {
tmp16022 := Call1(__e, SymFn(symshen_4unlocked_2), V5273)


if True == tmp16022 {
tmp16016 := Call2(__e, SymFn(symshen_4lazyderef), V5271, V5272)


let__14230 := tmp16016
_ = let__14230

tmp16020 := GConsP(__e, let__14230)

if True == tmp16020 {
tmp16017 := GTl(__e, let__14230)

let__14231 := tmp16017
_ = let__14231

tmp16018 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16018

if SymFn(symshen_4by_1hypothesis_1h) == __self {
__self1, __self2, __self3, __self4, __self5, __self6, __self7 = V5269, V5270, let__14231, V5272, V5273, V5274, V5275
__e.Tick()
goto __selftop
}
__e.TailApply(SymFn(symshen_4by_1hypothesis_1h), V5269, V5270, let__14231, V5272, V5273, V5274, V5275)
return nil


} else {
return False
}


} else {
return False
}


} else {
return let__14221
}


panic("unreachable")
}())
}, 7)

tmp16025 := Call2(__e, ns2_1set, symshen_4by_1hypothesis_1h, tmp15982)


_ = tmp16025

tmp16026 := MakeNative6(func(__e *ControlFlow, V5287, V5288, V5289, V5290, V5291, V5292 Obj) Obj {
tmp16031 := Call1(__e, SymFn(symshen_4unlocked_2), V5290)


if True == tmp16031 {
tmp16027 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16027

tmp16028 := GValue(__e, symshen_4_dsigf_d)

tmp16029 := Call2(__e, SymFn(symassoc), V5287, tmp16028)


return TailCall6(__e, SymFn(symshen_4sigf), tmp16029, V5288, V5289, V5290, V5291, V5292)


} else {
return False
}


panic("unreachable")
})

tmp16032 := Call2(__e, ns2_1set, symshen_4lookupsig, tmp16026)


_ = tmp16032

tmp16033 := MakeNative6(func(__e *ControlFlow, V5307, V5308, V5309, V5310, V5311, V5312 Obj) Obj {
tmp16040 := GConsP(__e, V5307)

if True == tmp16040 {
tmp16034 := GTl(__e, V5307)

tmp16035 := Call1(__e, tmp16034, V5308)


tmp16036 := Call1(__e, tmp16035, V5309)


tmp16037 := Call1(__e, tmp16036, V5310)


tmp16038 := Call1(__e, tmp16037, V5311)


return TailCall1(__e, tmp16038, V5312)


} else {
return False
}


panic("unreachable")
})

tmp16041 := Call2(__e, ns2_1set, symshen_4sigf, tmp16033)


_ = tmp16041

tmp16042 := MakeNative1(func(__e *ControlFlow, V5313 Obj) Obj {
tmp16043 := PrimGuard1(__e, symabsvector, PrimAbsvector, MakeInteger(3))

let__14232 := tmp16043
_ = let__14232

tmp16044 := GVecSet(__e, let__14232, MakeInteger(0), symshen_4print_1freshterm)

let__14233 := tmp16044
_ = let__14233

tmp16045 := GVecSet(__e, let__14233, MakeInteger(1), V5313)

let__14234 := tmp16045
_ = let__14234

tmp16046 := GValue(__e, symshen_4_dgensym_d)

tmp16048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dgensym_d, GAdd(__e, MakeInteger(1), tmp16046))
}
__typedArg0 := symshen_4_dgensym_d
__typedArg1 := GAdd(__e, MakeInteger(1), tmp16046)
return Call2(__e, SymFn(symset), __typedArg0, __typedArg1)
})()

tmp16049 := GVecSet(__e, let__14234, MakeInteger(2), tmp16048)

let__14235 := tmp16049
_ = let__14235

return let__14235


panic("unreachable")
})

tmp16050 := Call2(__e, ns2_1set, symshen_4freshterm, tmp16042)


_ = tmp16050

tmp16051 := MakeNative1(func(__e *ControlFlow, V5318 Obj) Obj {
tmp16052 := GVecGet(__e, V5318, MakeInteger(1))

tmp16053 := Call3(__e, SymFn(symshen_4app), tmp16052, MakeString(""), symshen_4a)


return PrimGuard2(__e, symcn, PrimStringConcat, MakeString("&&"), tmp16053)


panic("unreachable")
})

tmp16054 := Call2(__e, ns2_1set, symshen_4print_1freshterm, tmp16051)


_ = tmp16054

tmp16055 := MakeNative(func(__e *ControlFlow) {
V5319 := __e.Get(1)
V5320 := __e.Get(2)
V5321 := __e.Get(3)
V5322 := __e.Get(4)
V5323 := __e.Get(5)
V5324 := __e.Get(6)
V5325 := __e.Get(7)
__e.ReturnDirect(func() Obj {
tmp16059 := Call2(__e, SymFn(symshen_4lazyderef), V5321, V5322)


tmp16060 := GEq(__e, tmp16059, Nil)

if True == tmp16060 {
return False
} else {
tmp16056 := Call2(__e, SymFn(symshen_4deref), V5319, V5322)


tmp16057 := Call2(__e, SymFn(symshen_4deref), V5320, V5322)


__e.TailApply(SymFn(symshen_4search_1user_1datatypes_1h), tmp16056, tmp16057, V5321, V5322, V5323, V5324, V5325)
return nil


}


panic("unreachable")
}())
}, 7)

tmp16061 := Call2(__e, ns2_1set, symshen_4search_1user_1datatypes, tmp16055)


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
__e.ReturnDirect(func() Obj {
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
tmp16077 := Call1(__e, SymFn(symshen_4unlocked_2), V5330)


var ifres16063 Obj

if True == tmp16077 {
tmp16064 := Call2(__e, SymFn(symshen_4lazyderef), V5328, V5329)


let__14237 := tmp16064
_ = let__14237

tmp16076 := GConsP(__e, let__14237)

var ifres16065 Obj

if True == tmp16076 {
tmp16066 := GHd(__e, let__14237)

tmp16067 := Call2(__e, SymFn(symshen_4lazyderef), tmp16066, V5329)


let__14238 := tmp16067
_ = let__14238

tmp16075 := GConsP(__e, let__14238)

var ifres16068 Obj

if True == tmp16075 {
tmp16069 := GTl(__e, let__14238)

let__14239 := tmp16069
_ = let__14239

tmp16070 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16070

tmp16071 := Call2(__e, SymFn(symshen_4deref), let__14239, V5329)


tmp16072 := Call1(__e, tmp16071, V5326)


tmp16073 := Call1(__e, tmp16072, V5327)


tmp16074 := Call5(__e, SymFn(symcall), tmp16073, V5329, V5330, V5331, V5332)


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

tmp16086 := GEq(__e, let__14236, False)

if True == tmp16086 {
tmp16084 := Call1(__e, SymFn(symshen_4unlocked_2), V5330)


if True == tmp16084 {
tmp16078 := Call2(__e, SymFn(symshen_4lazyderef), V5328, V5329)


let__14240 := tmp16078
_ = let__14240

tmp16082 := GConsP(__e, let__14240)

if True == tmp16082 {
tmp16079 := GTl(__e, let__14240)

let__14241 := tmp16079
_ = let__14241

tmp16080 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16080

if SymFn(symshen_4search_1user_1datatypes_1h) == __self {
__self1, __self2, __self3, __self4, __self5, __self6, __self7 = V5326, V5327, let__14241, V5329, V5330, V5331, V5332
__e.Tick()
goto __selftop
}
__e.TailApply(SymFn(symshen_4search_1user_1datatypes_1h), V5326, V5327, let__14241, V5329, V5330, V5331, V5332)
return nil


} else {
return False
}


} else {
return False
}


} else {
return let__14236
}


panic("unreachable")
}())
}, 7)

tmp16087 := Call2(__e, ns2_1set, symshen_4search_1user_1datatypes_1h, tmp16062)


_ = tmp16087

tmp16088 := MakeNative(func(__e *ControlFlow) {
V5339 := __e.Get(1)
V5340 := __e.Get(2)
V5341 := __e.Get(3)
V5342 := __e.Get(4)
V5343 := __e.Get(5)
V5344 := __e.Get(6)
V5345 := __e.Get(7)
__e.ReturnDirect(func() Obj {
let__14242 := GAdd(__e, V5344, MakeInteger(1))
_ = let__14242

tmp16100 := Call1(__e, SymFn(symshen_4unlocked_2), V5343)


var ifres16090 Obj

if True == tmp16100 {
tmp16091 := Call2(__e, SymFn(symshen_4lazyderef), V5339, V5342)


let__14244 := tmp16091
_ = let__14244

tmp16099 := GEq(__e, let__14244, Nil)

var ifres16092 Obj

if True == tmp16099 {
tmp16093 := Call2(__e, SymFn(symshen_4lazyderef), V5341, V5342)


let__14245 := tmp16093
_ = let__14245

tmp16098 := GEq(__e, let__14245, True)

var ifres16094 Obj

if True == tmp16098 {
tmp16095 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16095

tmp16096 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall6(__e, SymFn(symbind), V5340, Nil, V5342, V5343, let__14242, V5345)
panic("unreachable")
})

tmp16097 := Call4(__e, SymFn(symshen_4cut), V5342, V5343, let__14242, tmp16096)


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

tmp16443 := GEq(__e, let__14243, False)

if True == tmp16443 {
tmp16179 := Call1(__e, SymFn(symshen_4unlocked_2), V5343)


var ifres16101 Obj

if True == tmp16179 {
tmp16102 := Call2(__e, SymFn(symshen_4lazyderef), V5339, V5342)


let__14247 := tmp16102
_ = let__14247

tmp16178 := GConsP(__e, let__14247)

var ifres16103 Obj

if True == tmp16178 {
tmp16104 := GHd(__e, let__14247)

tmp16105 := Call2(__e, SymFn(symshen_4lazyderef), tmp16104, V5342)


let__14248 := tmp16105
_ = let__14248

tmp16177 := GConsP(__e, let__14248)

var ifres16106 Obj

if True == tmp16177 {
tmp16107 := GHd(__e, let__14248)

tmp16108 := Call2(__e, SymFn(symshen_4lazyderef), tmp16107, V5342)


let__14249 := tmp16108
_ = let__14249

tmp16176 := GConsP(__e, let__14249)

var ifres16109 Obj

if True == tmp16176 {
tmp16110 := GHd(__e, let__14249)

tmp16111 := Call2(__e, SymFn(symshen_4lazyderef), tmp16110, V5342)


let__14250 := tmp16111
_ = let__14250

tmp16175 := GEq(__e, let__14250, symcons)

var ifres16112 Obj

if True == tmp16175 {
tmp16113 := GTl(__e, let__14249)

tmp16114 := Call2(__e, SymFn(symshen_4lazyderef), tmp16113, V5342)


let__14251 := tmp16114
_ = let__14251

tmp16174 := GConsP(__e, let__14251)

var ifres16115 Obj

if True == tmp16174 {
tmp16116 := GHd(__e, let__14251)

let__14252 := tmp16116
_ = let__14252

tmp16117 := GTl(__e, let__14251)

tmp16118 := Call2(__e, SymFn(symshen_4lazyderef), tmp16117, V5342)


let__14253 := tmp16118
_ = let__14253

tmp16173 := GConsP(__e, let__14253)

var ifres16119 Obj

if True == tmp16173 {
tmp16120 := GHd(__e, let__14253)

let__14254 := tmp16120
_ = let__14254

tmp16121 := GTl(__e, let__14253)

tmp16122 := Call2(__e, SymFn(symshen_4lazyderef), tmp16121, V5342)


let__14255 := tmp16122
_ = let__14255

tmp16172 := GEq(__e, let__14255, Nil)

var ifres16123 Obj

if True == tmp16172 {
tmp16124 := GTl(__e, let__14248)

tmp16125 := Call2(__e, SymFn(symshen_4lazyderef), tmp16124, V5342)


let__14256 := tmp16125
_ = let__14256

tmp16171 := GConsP(__e, let__14256)

var ifres16126 Obj

if True == tmp16171 {
tmp16127 := GHd(__e, let__14256)

let__14257 := tmp16127
_ = let__14257

tmp16128 := GTl(__e, let__14256)

tmp16129 := Call2(__e, SymFn(symshen_4lazyderef), tmp16128, V5342)


let__14258 := tmp16129
_ = let__14258

tmp16170 := GConsP(__e, let__14258)

var ifres16130 Obj

if True == tmp16170 {
tmp16131 := GHd(__e, let__14258)

tmp16132 := Call2(__e, SymFn(symshen_4lazyderef), tmp16131, V5342)


let__14259 := tmp16132
_ = let__14259

tmp16169 := GConsP(__e, let__14259)

var ifres16133 Obj

if True == tmp16169 {
tmp16134 := GHd(__e, let__14259)

tmp16135 := Call2(__e, SymFn(symshen_4lazyderef), tmp16134, V5342)


let__14260 := tmp16135
_ = let__14260

tmp16168 := GEq(__e, let__14260, symlist)

var ifres16136 Obj

if True == tmp16168 {
tmp16137 := GTl(__e, let__14259)

tmp16138 := Call2(__e, SymFn(symshen_4lazyderef), tmp16137, V5342)


let__14261 := tmp16138
_ = let__14261

tmp16167 := GConsP(__e, let__14261)

var ifres16139 Obj

if True == tmp16167 {
tmp16140 := GHd(__e, let__14261)

let__14262 := tmp16140
_ = let__14262

tmp16141 := GTl(__e, let__14261)

tmp16142 := Call2(__e, SymFn(symshen_4lazyderef), tmp16141, V5342)


let__14263 := tmp16142
_ = let__14263

tmp16166 := GEq(__e, let__14263, Nil)

var ifres16143 Obj

if True == tmp16166 {
tmp16144 := GTl(__e, let__14258)

tmp16145 := Call2(__e, SymFn(symshen_4lazyderef), tmp16144, V5342)


let__14264 := tmp16145
_ = let__14264

tmp16165 := GEq(__e, let__14264, Nil)

var ifres16146 Obj

if True == tmp16165 {
tmp16147 := GTl(__e, let__14247)

let__14265 := tmp16147
_ = let__14265

tmp16148 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16148

tmp16149 := Call2(__e, SymFn(symshen_4deref), let__14257, V5342)


tmp16150 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp16151 := GEq(__e, tmp16149, tmp16150)

tmp16152 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16153 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16154 := GCons(__e, let__14262, Nil)

tmp16155 := GCons(__e, let__14257, tmp16154)

tmp16156 := GCons(__e, let__14252, tmp16155)

tmp16157 := GCons(__e, let__14262, Nil)

tmp16158 := GCons(__e, symlist, tmp16157)

tmp16159 := GCons(__e, tmp16158, Nil)

tmp16160 := GCons(__e, let__14257, tmp16159)

tmp16161 := GCons(__e, let__14254, tmp16160)

tmp16162 := GCons(__e, tmp16161, let__14265)

tmp16163 := GCons(__e, tmp16156, tmp16162)

__e.TailApply(SymFn(symshen_4l_1rules), tmp16163, V5340, True, V5342, V5343, let__14242, V5345)
return nil


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5342, V5343, let__14242, tmp16153)


panic("unreachable")
})

tmp16164 := Call5(__e, SymFn(symwhen), tmp16151, V5342, V5343, let__14242, tmp16152)


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

tmp16441 := GEq(__e, let__14246, False)

if True == tmp16441 {
tmp16261 := Call1(__e, SymFn(symshen_4unlocked_2), V5343)


var ifres16180 Obj

if True == tmp16261 {
tmp16181 := Call2(__e, SymFn(symshen_4lazyderef), V5339, V5342)


let__14267 := tmp16181
_ = let__14267

tmp16260 := GConsP(__e, let__14267)

var ifres16182 Obj

if True == tmp16260 {
tmp16183 := GHd(__e, let__14267)

tmp16184 := Call2(__e, SymFn(symshen_4lazyderef), tmp16183, V5342)


let__14268 := tmp16184
_ = let__14268

tmp16259 := GConsP(__e, let__14268)

var ifres16185 Obj

if True == tmp16259 {
tmp16186 := GHd(__e, let__14268)

tmp16187 := Call2(__e, SymFn(symshen_4lazyderef), tmp16186, V5342)


let__14269 := tmp16187
_ = let__14269

tmp16258 := GConsP(__e, let__14269)

var ifres16188 Obj

if True == tmp16258 {
tmp16189 := GHd(__e, let__14269)

tmp16190 := Call2(__e, SymFn(symshen_4lazyderef), tmp16189, V5342)


let__14270 := tmp16190
_ = let__14270

tmp16257 := GEq(__e, let__14270, sym_8p)

var ifres16191 Obj

if True == tmp16257 {
tmp16192 := GTl(__e, let__14269)

tmp16193 := Call2(__e, SymFn(symshen_4lazyderef), tmp16192, V5342)


let__14271 := tmp16193
_ = let__14271

tmp16256 := GConsP(__e, let__14271)

var ifres16194 Obj

if True == tmp16256 {
tmp16195 := GHd(__e, let__14271)

let__14272 := tmp16195
_ = let__14272

tmp16196 := GTl(__e, let__14271)

tmp16197 := Call2(__e, SymFn(symshen_4lazyderef), tmp16196, V5342)


let__14273 := tmp16197
_ = let__14273

tmp16255 := GConsP(__e, let__14273)

var ifres16198 Obj

if True == tmp16255 {
tmp16199 := GHd(__e, let__14273)

let__14274 := tmp16199
_ = let__14274

tmp16200 := GTl(__e, let__14273)

tmp16201 := Call2(__e, SymFn(symshen_4lazyderef), tmp16200, V5342)


let__14275 := tmp16201
_ = let__14275

tmp16254 := GEq(__e, let__14275, Nil)

var ifres16202 Obj

if True == tmp16254 {
tmp16203 := GTl(__e, let__14268)

tmp16204 := Call2(__e, SymFn(symshen_4lazyderef), tmp16203, V5342)


let__14276 := tmp16204
_ = let__14276

tmp16253 := GConsP(__e, let__14276)

var ifres16205 Obj

if True == tmp16253 {
tmp16206 := GHd(__e, let__14276)

let__14277 := tmp16206
_ = let__14277

tmp16207 := GTl(__e, let__14276)

tmp16208 := Call2(__e, SymFn(symshen_4lazyderef), tmp16207, V5342)


let__14278 := tmp16208
_ = let__14278

tmp16252 := GConsP(__e, let__14278)

var ifres16209 Obj

if True == tmp16252 {
tmp16210 := GHd(__e, let__14278)

tmp16211 := Call2(__e, SymFn(symshen_4lazyderef), tmp16210, V5342)


let__14279 := tmp16211
_ = let__14279

tmp16251 := GConsP(__e, let__14279)

var ifres16212 Obj

if True == tmp16251 {
tmp16213 := GHd(__e, let__14279)

let__14280 := tmp16213
_ = let__14280

tmp16214 := GTl(__e, let__14279)

tmp16215 := Call2(__e, SymFn(symshen_4lazyderef), tmp16214, V5342)


let__14281 := tmp16215
_ = let__14281

tmp16250 := GConsP(__e, let__14281)

var ifres16216 Obj

if True == tmp16250 {
tmp16217 := GHd(__e, let__14281)

tmp16218 := Call2(__e, SymFn(symshen_4lazyderef), tmp16217, V5342)


let__14282 := tmp16218
_ = let__14282

tmp16249 := GEq(__e, let__14282, sym_d)

var ifres16219 Obj

if True == tmp16249 {
tmp16220 := GTl(__e, let__14281)

tmp16221 := Call2(__e, SymFn(symshen_4lazyderef), tmp16220, V5342)


let__14283 := tmp16221
_ = let__14283

tmp16248 := GConsP(__e, let__14283)

var ifres16222 Obj

if True == tmp16248 {
tmp16223 := GHd(__e, let__14283)

let__14284 := tmp16223
_ = let__14284

tmp16224 := GTl(__e, let__14283)

tmp16225 := Call2(__e, SymFn(symshen_4lazyderef), tmp16224, V5342)


let__14285 := tmp16225
_ = let__14285

tmp16247 := GEq(__e, let__14285, Nil)

var ifres16226 Obj

if True == tmp16247 {
tmp16227 := GTl(__e, let__14278)

tmp16228 := Call2(__e, SymFn(symshen_4lazyderef), tmp16227, V5342)


let__14286 := tmp16228
_ = let__14286

tmp16246 := GEq(__e, let__14286, Nil)

var ifres16229 Obj

if True == tmp16246 {
tmp16230 := GTl(__e, let__14267)

let__14287 := tmp16230
_ = let__14287

tmp16231 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16231

tmp16232 := Call2(__e, SymFn(symshen_4deref), let__14277, V5342)


tmp16233 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp16234 := GEq(__e, tmp16232, tmp16233)

tmp16235 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16236 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16237 := GCons(__e, let__14280, Nil)

tmp16238 := GCons(__e, let__14277, tmp16237)

tmp16239 := GCons(__e, let__14272, tmp16238)

tmp16240 := GCons(__e, let__14284, Nil)

tmp16241 := GCons(__e, let__14277, tmp16240)

tmp16242 := GCons(__e, let__14274, tmp16241)

tmp16243 := GCons(__e, tmp16242, let__14287)

tmp16244 := GCons(__e, tmp16239, tmp16243)

__e.TailApply(SymFn(symshen_4l_1rules), tmp16244, V5340, True, V5342, V5343, let__14242, V5345)
return nil


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5342, V5343, let__14242, tmp16236)


panic("unreachable")
})

tmp16245 := Call5(__e, SymFn(symwhen), tmp16234, V5342, V5343, let__14242, tmp16235)


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

tmp16439 := GEq(__e, let__14266, False)

if True == tmp16439 {
tmp16325 := Call1(__e, SymFn(symshen_4unlocked_2), V5343)


var ifres16262 Obj

if True == tmp16325 {
tmp16263 := Call2(__e, SymFn(symshen_4lazyderef), V5339, V5342)


let__14289 := tmp16263
_ = let__14289

tmp16324 := GConsP(__e, let__14289)

var ifres16264 Obj

if True == tmp16324 {
tmp16265 := GHd(__e, let__14289)

tmp16266 := Call2(__e, SymFn(symshen_4lazyderef), tmp16265, V5342)


let__14290 := tmp16266
_ = let__14290

tmp16323 := GConsP(__e, let__14290)

var ifres16267 Obj

if True == tmp16323 {
tmp16268 := GHd(__e, let__14290)

tmp16269 := Call2(__e, SymFn(symshen_4lazyderef), tmp16268, V5342)


let__14291 := tmp16269
_ = let__14291

tmp16322 := GConsP(__e, let__14291)

var ifres16270 Obj

if True == tmp16322 {
tmp16271 := GHd(__e, let__14291)

tmp16272 := Call2(__e, SymFn(symshen_4lazyderef), tmp16271, V5342)


let__14292 := tmp16272
_ = let__14292

tmp16321 := GEq(__e, let__14292, sym_8s)

var ifres16273 Obj

if True == tmp16321 {
tmp16274 := GTl(__e, let__14291)

tmp16275 := Call2(__e, SymFn(symshen_4lazyderef), tmp16274, V5342)


let__14293 := tmp16275
_ = let__14293

tmp16320 := GConsP(__e, let__14293)

var ifres16276 Obj

if True == tmp16320 {
tmp16277 := GHd(__e, let__14293)

let__14294 := tmp16277
_ = let__14294

tmp16278 := GTl(__e, let__14293)

tmp16279 := Call2(__e, SymFn(symshen_4lazyderef), tmp16278, V5342)


let__14295 := tmp16279
_ = let__14295

tmp16319 := GConsP(__e, let__14295)

var ifres16280 Obj

if True == tmp16319 {
tmp16281 := GHd(__e, let__14295)

let__14296 := tmp16281
_ = let__14296

tmp16282 := GTl(__e, let__14295)

tmp16283 := Call2(__e, SymFn(symshen_4lazyderef), tmp16282, V5342)


let__14297 := tmp16283
_ = let__14297

tmp16318 := GEq(__e, let__14297, Nil)

var ifres16284 Obj

if True == tmp16318 {
tmp16285 := GTl(__e, let__14290)

tmp16286 := Call2(__e, SymFn(symshen_4lazyderef), tmp16285, V5342)


let__14298 := tmp16286
_ = let__14298

tmp16317 := GConsP(__e, let__14298)

var ifres16287 Obj

if True == tmp16317 {
tmp16288 := GHd(__e, let__14298)

let__14299 := tmp16288
_ = let__14299

tmp16289 := GTl(__e, let__14298)

tmp16290 := Call2(__e, SymFn(symshen_4lazyderef), tmp16289, V5342)


let__14300 := tmp16290
_ = let__14300

tmp16316 := GConsP(__e, let__14300)

var ifres16291 Obj

if True == tmp16316 {
tmp16292 := GHd(__e, let__14300)

tmp16293 := Call2(__e, SymFn(symshen_4lazyderef), tmp16292, V5342)


let__14301 := tmp16293
_ = let__14301

tmp16315 := GEq(__e, let__14301, symstring)

var ifres16294 Obj

if True == tmp16315 {
tmp16295 := GTl(__e, let__14300)

tmp16296 := Call2(__e, SymFn(symshen_4lazyderef), tmp16295, V5342)


let__14302 := tmp16296
_ = let__14302

tmp16314 := GEq(__e, let__14302, Nil)

var ifres16297 Obj

if True == tmp16314 {
tmp16298 := GTl(__e, let__14289)

let__14303 := tmp16298
_ = let__14303

tmp16299 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16299

tmp16300 := Call2(__e, SymFn(symshen_4deref), let__14299, V5342)


tmp16301 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp16302 := GEq(__e, tmp16300, tmp16301)

tmp16303 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16304 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16305 := GCons(__e, symstring, Nil)

tmp16306 := GCons(__e, let__14299, tmp16305)

tmp16307 := GCons(__e, let__14294, tmp16306)

tmp16308 := GCons(__e, symstring, Nil)

tmp16309 := GCons(__e, let__14299, tmp16308)

tmp16310 := GCons(__e, let__14296, tmp16309)

tmp16311 := GCons(__e, tmp16310, let__14303)

tmp16312 := GCons(__e, tmp16307, tmp16311)

__e.TailApply(SymFn(symshen_4l_1rules), tmp16312, V5340, True, V5342, V5343, let__14242, V5345)
return nil


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5342, V5343, let__14242, tmp16304)


panic("unreachable")
})

tmp16313 := Call5(__e, SymFn(symwhen), tmp16302, V5342, V5343, let__14242, tmp16303)


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

tmp16437 := GEq(__e, let__14288, False)

if True == tmp16437 {
tmp16404 := Call1(__e, SymFn(symshen_4unlocked_2), V5343)


var ifres16326 Obj

if True == tmp16404 {
tmp16327 := Call2(__e, SymFn(symshen_4lazyderef), V5339, V5342)


let__14305 := tmp16327
_ = let__14305

tmp16403 := GConsP(__e, let__14305)

var ifres16328 Obj

if True == tmp16403 {
tmp16329 := GHd(__e, let__14305)

tmp16330 := Call2(__e, SymFn(symshen_4lazyderef), tmp16329, V5342)


let__14306 := tmp16330
_ = let__14306

tmp16402 := GConsP(__e, let__14306)

var ifres16331 Obj

if True == tmp16402 {
tmp16332 := GHd(__e, let__14306)

tmp16333 := Call2(__e, SymFn(symshen_4lazyderef), tmp16332, V5342)


let__14307 := tmp16333
_ = let__14307

tmp16401 := GConsP(__e, let__14307)

var ifres16334 Obj

if True == tmp16401 {
tmp16335 := GHd(__e, let__14307)

tmp16336 := Call2(__e, SymFn(symshen_4lazyderef), tmp16335, V5342)


let__14308 := tmp16336
_ = let__14308

tmp16400 := GEq(__e, let__14308, sym_8v)

var ifres16337 Obj

if True == tmp16400 {
tmp16338 := GTl(__e, let__14307)

tmp16339 := Call2(__e, SymFn(symshen_4lazyderef), tmp16338, V5342)


let__14309 := tmp16339
_ = let__14309

tmp16399 := GConsP(__e, let__14309)

var ifres16340 Obj

if True == tmp16399 {
tmp16341 := GHd(__e, let__14309)

let__14310 := tmp16341
_ = let__14310

tmp16342 := GTl(__e, let__14309)

tmp16343 := Call2(__e, SymFn(symshen_4lazyderef), tmp16342, V5342)


let__14311 := tmp16343
_ = let__14311

tmp16398 := GConsP(__e, let__14311)

var ifres16344 Obj

if True == tmp16398 {
tmp16345 := GHd(__e, let__14311)

let__14312 := tmp16345
_ = let__14312

tmp16346 := GTl(__e, let__14311)

tmp16347 := Call2(__e, SymFn(symshen_4lazyderef), tmp16346, V5342)


let__14313 := tmp16347
_ = let__14313

tmp16397 := GEq(__e, let__14313, Nil)

var ifres16348 Obj

if True == tmp16397 {
tmp16349 := GTl(__e, let__14306)

tmp16350 := Call2(__e, SymFn(symshen_4lazyderef), tmp16349, V5342)


let__14314 := tmp16350
_ = let__14314

tmp16396 := GConsP(__e, let__14314)

var ifres16351 Obj

if True == tmp16396 {
tmp16352 := GHd(__e, let__14314)

let__14315 := tmp16352
_ = let__14315

tmp16353 := GTl(__e, let__14314)

tmp16354 := Call2(__e, SymFn(symshen_4lazyderef), tmp16353, V5342)


let__14316 := tmp16354
_ = let__14316

tmp16395 := GConsP(__e, let__14316)

var ifres16355 Obj

if True == tmp16395 {
tmp16356 := GHd(__e, let__14316)

tmp16357 := Call2(__e, SymFn(symshen_4lazyderef), tmp16356, V5342)


let__14317 := tmp16357
_ = let__14317

tmp16394 := GConsP(__e, let__14317)

var ifres16358 Obj

if True == tmp16394 {
tmp16359 := GHd(__e, let__14317)

tmp16360 := Call2(__e, SymFn(symshen_4lazyderef), tmp16359, V5342)


let__14318 := tmp16360
_ = let__14318

tmp16393 := GEq(__e, let__14318, symvector)

var ifres16361 Obj

if True == tmp16393 {
tmp16362 := GTl(__e, let__14317)

tmp16363 := Call2(__e, SymFn(symshen_4lazyderef), tmp16362, V5342)


let__14319 := tmp16363
_ = let__14319

tmp16392 := GConsP(__e, let__14319)

var ifres16364 Obj

if True == tmp16392 {
tmp16365 := GHd(__e, let__14319)

let__14320 := tmp16365
_ = let__14320

tmp16366 := GTl(__e, let__14319)

tmp16367 := Call2(__e, SymFn(symshen_4lazyderef), tmp16366, V5342)


let__14321 := tmp16367
_ = let__14321

tmp16391 := GEq(__e, let__14321, Nil)

var ifres16368 Obj

if True == tmp16391 {
tmp16369 := GTl(__e, let__14316)

tmp16370 := Call2(__e, SymFn(symshen_4lazyderef), tmp16369, V5342)


let__14322 := tmp16370
_ = let__14322

tmp16390 := GEq(__e, let__14322, Nil)

var ifres16371 Obj

if True == tmp16390 {
tmp16372 := GTl(__e, let__14305)

let__14323 := tmp16372
_ = let__14323

tmp16373 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16373

tmp16374 := Call2(__e, SymFn(symshen_4deref), let__14315, V5342)


tmp16375 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp16376 := GEq(__e, tmp16374, tmp16375)

tmp16377 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16378 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16379 := GCons(__e, let__14320, Nil)

tmp16380 := GCons(__e, let__14315, tmp16379)

tmp16381 := GCons(__e, let__14310, tmp16380)

tmp16382 := GCons(__e, let__14320, Nil)

tmp16383 := GCons(__e, symvector, tmp16382)

tmp16384 := GCons(__e, tmp16383, Nil)

tmp16385 := GCons(__e, let__14315, tmp16384)

tmp16386 := GCons(__e, let__14312, tmp16385)

tmp16387 := GCons(__e, tmp16386, let__14323)

tmp16388 := GCons(__e, tmp16381, tmp16387)

__e.TailApply(SymFn(symshen_4l_1rules), tmp16388, V5340, True, V5342, V5343, let__14242, V5345)
return nil


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5342, V5343, let__14242, tmp16378)


panic("unreachable")
})

tmp16389 := Call5(__e, SymFn(symwhen), tmp16376, V5342, V5343, let__14242, tmp16377)


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

tmp16435 := GEq(__e, let__14304, False)

if True == tmp16435 {
tmp16431 := Call1(__e, SymFn(symshen_4unlocked_2), V5343)


var ifres16405 Obj

if True == tmp16431 {
tmp16406 := Call2(__e, SymFn(symshen_4lazyderef), V5339, V5342)


let__14325 := tmp16406
_ = let__14325

tmp16430 := GConsP(__e, let__14325)

var ifres16407 Obj

if True == tmp16430 {
tmp16408 := GHd(__e, let__14325)

let__14326 := tmp16408
_ = let__14326

tmp16409 := GTl(__e, let__14325)

let__14327 := tmp16409
_ = let__14327

tmp16410 := Call2(__e, SymFn(symshen_4lazyderef), V5340, V5342)


let__14328 := tmp16410
_ = let__14328

tmp16411 := MakeNative1(func(__e *ControlFlow, Z5434 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5435 Obj) Obj {
tmp16412 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16412

tmp16413 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4l_1rules), let__14327, Z5435, V5341, V5342, V5343, let__14242, V5345)
return nil
panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), Z5434, let__14326, V5342, V5343, let__14242, tmp16413)


panic("unreachable")
})
panic("unreachable")
})

let__14329 := tmp16411
_ = let__14329

tmp16429 := GConsP(__e, let__14328)

var ifres16414 Obj

if True == tmp16429 {
tmp16415 := GHd(__e, let__14328)

let__14330 := tmp16415
_ = let__14330

tmp16416 := GTl(__e, let__14328)

let__14331 := tmp16416
_ = let__14331

tmp16417 := Call1(__e, let__14329, let__14330)


tmp16418 := Call1(__e, tmp16417, let__14331)


ifres16414 = tmp16418


} else {
tmp16428 := Call1(__e, SymFn(symshen_4pvar_2), let__14328)


var ifres16419 Obj

if True == tmp16428 {
tmp16420 := Call1(__e, SymFn(symshen_4newpv), V5342)


let__14332 := tmp16420
_ = let__14332

tmp16421 := Call1(__e, SymFn(symshen_4newpv), V5342)


let__14333 := tmp16421
_ = let__14333

tmp16422 := GCons(__e, let__14332, let__14333)

tmp16423 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16424 := Call1(__e, let__14329, let__14332)


return TailCall1(__e, tmp16424, let__14333)


panic("unreachable")
})

tmp16425 := Call4(__e, SymFn(symshen_4bind_b), let__14328, tmp16422, V5342, tmp16423)


tmp16426 := Call2(__e, SymFn(symshen_4gc), V5342, tmp16425)


tmp16427 := Call2(__e, SymFn(symshen_4gc), V5342, tmp16426)


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

tmp16433 := GEq(__e, let__14324, False)

if True == tmp16433 {
return TailCall2(__e, SymFn(symshen_4unlock), V5343, let__14242)
} else {
return let__14324
}


} else {
return let__14304
}


} else {
return let__14288
}


} else {
return let__14266
}


} else {
return let__14246
}


} else {
return let__14243
}


panic("unreachable")
}())
}, 7)

tmp16444 := Call2(__e, ns2_1set, symshen_4l_1rules, tmp16088)


_ = tmp16444

tmp16445 := MakeNative6(func(__e *ControlFlow, V5440, V5441, V5442, V5443, V5444, V5445 Obj) Obj {
let__14334 := GAdd(__e, V5444, MakeInteger(1))
_ = let__14334

tmp16485 := Call1(__e, SymFn(symshen_4unlocked_2), V5443)


var ifres16447 Obj

if True == tmp16485 {
tmp16448 := Call2(__e, SymFn(symshen_4lazyderef), V5440, V5442)


let__14336 := tmp16448
_ = let__14336

tmp16484 := GConsP(__e, let__14336)

var ifres16449 Obj

if True == tmp16484 {
tmp16450 := GHd(__e, let__14336)

tmp16451 := Call2(__e, SymFn(symshen_4lazyderef), tmp16450, V5442)


let__14337 := tmp16451
_ = let__14337

tmp16483 := GEq(__e, let__14337, symdefine)

var ifres16452 Obj

if True == tmp16483 {
tmp16453 := GTl(__e, let__14336)

tmp16454 := Call2(__e, SymFn(symshen_4lazyderef), tmp16453, V5442)


let__14338 := tmp16454
_ = let__14338

tmp16482 := GConsP(__e, let__14338)

var ifres16455 Obj

if True == tmp16482 {
tmp16456 := GHd(__e, let__14338)

let__14339 := tmp16456
_ = let__14339

tmp16457 := GTl(__e, let__14338)

let__14340 := tmp16457
_ = let__14340

tmp16458 := Call1(__e, SymFn(symshen_4newpv), V5442)


let__14341 := tmp16458
_ = let__14341

tmp16459 := Call1(__e, SymFn(symshen_4newpv), V5442)


let__14342 := tmp16459
_ = let__14342

tmp16460 := Call1(__e, SymFn(symshen_4newpv), V5442)


let__14343 := tmp16460
_ = let__14343

tmp16461 := Call1(__e, SymFn(symshen_4newpv), V5442)


let__14344 := tmp16461
_ = let__14344

tmp16462 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16462

tmp16463 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16464 := GCons(__e, let__14339, let__14340)

tmp16465 := Call1(__e, SymFn(symshen_4sigxrules), tmp16464)


tmp16466 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16467 := Call2(__e, SymFn(symshen_4lazyderef), let__14341, V5442)


tmp16468 := Call1(__e, SymFn(symfst), tmp16467)


tmp16469 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16470 := Call2(__e, SymFn(symshen_4lazyderef), let__14341, V5442)


tmp16471 := Call1(__e, SymFn(symsnd), tmp16470)


tmp16472 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16473 := Call2(__e, SymFn(symshen_4deref), let__14344, V5442)


tmp16474 := Call1(__e, SymFn(symshen_4freshen_1sig), tmp16473)


tmp16475 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16476 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall6(__e, SymFn(symis), let__14344, V5441, V5442, V5443, let__14334, V5445)
panic("unreachable")
})

__e.TailApply(SymFn(symshen_4t_d_1rules), let__14339, let__14342, let__14343, MakeInteger(1), V5442, V5443, let__14334, tmp16476)
return nil


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), let__14343, tmp16474, V5442, V5443, let__14334, tmp16475)


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), let__14342, tmp16471, V5442, V5443, let__14334, tmp16472)


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), let__14344, tmp16468, V5442, V5443, let__14334, tmp16469)


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), let__14341, tmp16465, V5442, V5443, let__14334, tmp16466)


panic("unreachable")
})

tmp16477 := Call4(__e, SymFn(symshen_4cut), V5442, V5443, let__14334, tmp16463)


tmp16478 := Call2(__e, SymFn(symshen_4gc), V5442, tmp16477)


tmp16479 := Call2(__e, SymFn(symshen_4gc), V5442, tmp16478)


tmp16480 := Call2(__e, SymFn(symshen_4gc), V5442, tmp16479)


tmp16481 := Call2(__e, SymFn(symshen_4gc), V5442, tmp16480)


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

tmp16487 := GEq(__e, let__14335, False)

if True == tmp16487 {
return TailCall2(__e, SymFn(symshen_4unlock), V5443, let__14334)
} else {
return let__14335
}


panic("unreachable")
})

tmp16488 := Call2(__e, ns2_1set, symshen_4t_d, tmp16445)


_ = tmp16488

tmp16489 := MakeNative1(func(__e *ControlFlow, V5457 Obj) Obj {
tmp16490 := MakeNative1(func(__e *ControlFlow, Z5459 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5sig_drules_6), Z5459)
panic("unreachable")
})

tmp16491 := Call2(__e, SymFn(symcompile), tmp16490, V5457)


let__14345 := tmp16491
_ = let__14345

tmp16492 := Call1(__e, SymFn(symfst), let__14345)


tmp16493 := MakeNative1(func(__e *ControlFlow, Z5460 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4unmark_1rule_1custom_1patterns), Z5460)
panic("unreachable")
})

tmp16494 := Call1(__e, SymFn(symsnd), let__14345)


tmp16495 := Call2(__e, SymFn(symmap), tmp16493, tmp16494)


return TailCall2(__e, SymFn(sym_8p), tmp16492, tmp16495)


panic("unreachable")
})

tmp16496 := Call2(__e, ns2_1set, symshen_4sigxrules, tmp16489)


_ = tmp16496

tmp16497 := MakeNative1(func(__e *ControlFlow, V5461 Obj) Obj {
tmp16503 := Call1(__e, SymFn(symtuple_2), V5461)


if True == tmp16503 {
tmp16498 := MakeNative1(func(__e *ControlFlow, Z5462 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4unmark_1custom_1pattern), Z5462)
panic("unreachable")
})

tmp16499 := Call1(__e, SymFn(symfst), V5461)


tmp16500 := Call2(__e, SymFn(symmap), tmp16498, tmp16499)


tmp16501 := Call1(__e, SymFn(symsnd), V5461)


return TailCall2(__e, SymFn(sym_8p), tmp16500, tmp16501)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4unmark_1rule_1custom_1patterns)
}


panic("unreachable")
})

tmp16504 := Call2(__e, ns2_1set, symshen_4unmark_1rule_1custom_1patterns, tmp16497)


_ = tmp16504

tmp16505 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V5463 := __self1
_ = V5463
tmp16541 := GConsP(__e, V5463)

var ifres16517 Obj

if True == tmp16541 {
tmp16539 := GHd(__e, V5463)

tmp16540 := GEq(__e, sym_8p, tmp16539)

var ifres16519 Obj

if True == tmp16540 {
tmp16537 := GTl(__e, V5463)

tmp16538 := GConsP(__e, tmp16537)

var ifres16521 Obj

if True == tmp16538 {
tmp16534 := GTl(__e, V5463)

tmp16535 := GHd(__e, tmp16534)

tmp16536 := GEq(__e, symshen_4custom_1pattern, tmp16535)

var ifres16523 Obj

if True == tmp16536 {
tmp16531 := GTl(__e, V5463)

tmp16532 := GTl(__e, tmp16531)

tmp16533 := GConsP(__e, tmp16532)

var ifres16525 Obj

if True == tmp16533 {
tmp16527 := GTl(__e, V5463)

tmp16528 := GTl(__e, tmp16527)

tmp16529 := GTl(__e, tmp16528)

tmp16530 := GEq(__e, Nil, tmp16529)

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
tmp16506 := GTl(__e, V5463)

tmp16507 := GTl(__e, tmp16506)

tmp16508 := GHd(__e, tmp16507)

tmp16509 := Call1(__e, SymFn(symshen_4curry), tmp16508)


if SymFn(symshen_4unmark_1custom_1pattern) == __self {
__self1 = tmp16509
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4unmark_1custom_1pattern), tmp16509)


} else {
tmp16515 := GConsP(__e, V5463)

if True == tmp16515 {
tmp16510 := GHd(__e, V5463)

tmp16511 := Call1(__e, SymFn(symshen_4unmark_1custom_1pattern), tmp16510)


tmp16512 := GTl(__e, V5463)

tmp16513 := Call1(__e, SymFn(symshen_4unmark_1custom_1pattern), tmp16512)


return GCons(__e, tmp16511, tmp16513)


} else {
return V5463
}


}


panic("unreachable")
})
return __self
})()

tmp16542 := Call2(__e, ns2_1set, symshen_4unmark_1custom_1pattern, tmp16505)


_ = tmp16542

tmp16543 := MakeNative1(func(__e *ControlFlow, V5464 Obj) Obj {
tmp16570 := GConsP(__e, V5464)

var ifres16544 Obj

if True == tmp16570 {
tmp16545 := Call1(__e, SymFn(symtail), V5464)


let__14347 := tmp16545
_ = let__14347

tmp16568 := Call2(__e, SymFn(symshen_4hds_a_2), let__14347, sym_i)


var ifres16546 Obj

if True == tmp16568 {
tmp16547 := Call1(__e, SymFn(symtail), let__14347)


let__14348 := tmp16547
_ = let__14348

tmp16548 := Call1(__e, SymFn(symshen_4_5signature_6), let__14348)


let__14349 := tmp16548
_ = let__14349

tmp16566 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__14349)


var ifres16549 Obj

if True == tmp16566 {
tmp16550 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16549 = tmp16550


} else {
tmp16551 := Call1(__e, SymFn(symshen_4_5_1out), let__14349)


let__14350 := tmp16551
_ = let__14350

tmp16552 := Call1(__e, SymFn(symshen_4in_1_6), let__14349)


let__14351 := tmp16552
_ = let__14351

tmp16565 := Call2(__e, SymFn(symshen_4hds_a_2), let__14351, sym_j)


var ifres16553 Obj

if True == tmp16565 {
tmp16554 := Call1(__e, SymFn(symtail), let__14351)


let__14352 := tmp16554
_ = let__14352

tmp16555 := Call1(__e, SymFn(symshen_4_5rules_d_6), let__14352)


let__14353 := tmp16555
_ = let__14353

tmp16563 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__14353)


var ifres16556 Obj

if True == tmp16563 {
tmp16557 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16556 = tmp16557


} else {
tmp16558 := Call1(__e, SymFn(symshen_4_5_1out), let__14353)


let__14354 := tmp16558
_ = let__14354

tmp16559 := Call1(__e, SymFn(symshen_4in_1_6), let__14353)


let__14355 := tmp16559
_ = let__14355

tmp16560 := Call1(__e, SymFn(symshen_4rectify_1type), let__14350)


let__14356 := tmp16560
_ = let__14356

tmp16561 := Call2(__e, SymFn(sym_8p), let__14356, let__14354)


tmp16562 := Call2(__e, SymFn(symshen_4comb), let__14355, tmp16561)


ifres16556 = tmp16562


}

ifres16553 = ifres16556


} else {
tmp16564 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16553 = tmp16564


}

ifres16549 = ifres16553


}

ifres16546 = ifres16549


} else {
tmp16567 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16546 = tmp16567


}

ifres16544 = ifres16546


} else {
tmp16569 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16544 = tmp16569


}

let__14346 := ifres16544
_ = let__14346

tmp16572 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__14346)


if True == tmp16572 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__14346
}


panic("unreachable")
})

tmp16573 := Call2(__e, ns2_1set, symshen_4_5sig_drules_6, tmp16543)


_ = tmp16573

tmp16574 := MakeNative1(func(__e *ControlFlow, V5476 Obj) Obj {
tmp16575 := Call1(__e, SymFn(symshen_4extract_1vars), V5476)


let__14357 := tmp16575
_ = let__14357

tmp16576 := MakeNative1(func(__e *ControlFlow, Z5479 Obj) Obj {
tmp16577 := Call2(__e, SymFn(symconcat), sym_e, Z5479)


tmp16578 := Call1(__e, SymFn(symshen_4freshterm), tmp16577)


return GCons(__e, Z5479, tmp16578)


panic("unreachable")
})

tmp16579 := Call2(__e, SymFn(symmap), tmp16576, let__14357)


let__14358 := tmp16579
_ = let__14358

return TailCall2(__e, SymFn(symshen_4freshen_1type), let__14358, V5476)


panic("unreachable")
})

tmp16580 := Call2(__e, ns2_1set, symshen_4freshen_1sig, tmp16574)


_ = tmp16580

tmp16581 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V5480 := __self1
_ = V5480
V5481 := __self2
_ = V5481
tmp16595 := GEq(__e, Nil, V5480)

if True == tmp16595 {
return V5481
} else {
tmp16593 := GConsP(__e, V5480)

var ifres16589 Obj

if True == tmp16593 {
tmp16591 := GHd(__e, V5480)

tmp16592 := GConsP(__e, tmp16591)

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
tmp16582 := GTl(__e, V5480)

tmp16583 := GHd(__e, V5480)

tmp16584 := GTl(__e, tmp16583)

tmp16585 := GHd(__e, V5480)

tmp16586 := GHd(__e, tmp16585)

tmp16587 := Call3(__e, SymFn(symsubst), tmp16584, tmp16586, V5481)


if SymFn(symshen_4freshen_1type) == __self {
__self1, __self2 = tmp16582, tmp16587
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4freshen_1type), tmp16582, tmp16587)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4freshen_1type)
}


}


panic("unreachable")
})
return __self
})()

tmp16596 := Call2(__e, ns2_1set, symshen_4freshen_1type, tmp16581)


_ = tmp16596

tmp16597 := MakeNative1(func(__e *ControlFlow, V5482 Obj) Obj {
tmp16598 := Call1(__e, SymFn(symshen_4_5rule_d_6), V5482)


let__14360 := tmp16598
_ = let__14360

tmp16611 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__14360)


var ifres16599 Obj

if True == tmp16611 {
tmp16600 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16599 = tmp16600


} else {
tmp16601 := Call1(__e, SymFn(symshen_4_5_1out), let__14360)


let__14361 := tmp16601
_ = let__14361

tmp16602 := Call1(__e, SymFn(symshen_4in_1_6), let__14360)


let__14362 := tmp16602
_ = let__14362

tmp16603 := Call1(__e, SymFn(symshen_4_5rules_d_6), let__14362)


let__14363 := tmp16603
_ = let__14363

tmp16610 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__14363)


var ifres16604 Obj

if True == tmp16610 {
tmp16605 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16604 = tmp16605


} else {
tmp16606 := Call1(__e, SymFn(symshen_4_5_1out), let__14363)


let__14364 := tmp16606
_ = let__14364

tmp16607 := Call1(__e, SymFn(symshen_4in_1_6), let__14363)


let__14365 := tmp16607
_ = let__14365

tmp16608 := GCons(__e, let__14361, let__14364)

tmp16609 := Call2(__e, SymFn(symshen_4comb), let__14365, tmp16608)


ifres16604 = tmp16609


}

ifres16599 = ifres16604


}

let__14359 := ifres16599
_ = let__14359

tmp16623 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__14359)


if True == tmp16623 {
tmp16612 := Call1(__e, SymFn(symshen_4_5rule_d_6), V5482)


let__14367 := tmp16612
_ = let__14367

tmp16619 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__14367)


var ifres16613 Obj

if True == tmp16619 {
tmp16614 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16613 = tmp16614


} else {
tmp16615 := Call1(__e, SymFn(symshen_4_5_1out), let__14367)


let__14368 := tmp16615
_ = let__14368

tmp16616 := Call1(__e, SymFn(symshen_4in_1_6), let__14367)


let__14369 := tmp16616
_ = let__14369

tmp16617 := GCons(__e, let__14368, Nil)

tmp16618 := Call2(__e, SymFn(symshen_4comb), let__14369, tmp16617)


ifres16613 = tmp16618


}

let__14366 := ifres16613
_ = let__14366

tmp16621 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__14366)


if True == tmp16621 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__14366
}


} else {
return let__14359
}


panic("unreachable")
})

tmp16624 := Call2(__e, ns2_1set, symshen_4_5rules_d_6, tmp16597)


_ = tmp16624

tmp16625 := MakeNative1(func(__e *ControlFlow, V5494 Obj) Obj {
tmp16626 := Call1(__e, SymFn(symshen_4_5patterns_6), V5494)


let__14371 := tmp16626
_ = let__14371

tmp16654 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__14371)


var ifres16627 Obj

if True == tmp16654 {
tmp16628 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16627 = tmp16628


} else {
tmp16629 := Call1(__e, SymFn(symshen_4_5_1out), let__14371)


let__14372 := tmp16629
_ = let__14372

tmp16630 := Call1(__e, SymFn(symshen_4in_1_6), let__14371)


let__14373 := tmp16630
_ = let__14373

tmp16653 := Call2(__e, SymFn(symshen_4hds_a_2), let__14373, sym_1_6)


var ifres16631 Obj

if True == tmp16653 {
tmp16632 := Call1(__e, SymFn(symtail), let__14373)


let__14374 := tmp16632
_ = let__14374

tmp16651 := GConsP(__e, let__14374)

var ifres16633 Obj

if True == tmp16651 {
tmp16634 := Call1(__e, SymFn(symhead), let__14374)


let__14375 := tmp16634
_ = let__14375

tmp16635 := Call1(__e, SymFn(symtail), let__14374)


let__14376 := tmp16635
_ = let__14376

tmp16649 := Call2(__e, SymFn(symshen_4hds_a_2), let__14376, symwhere)


var ifres16636 Obj

if True == tmp16649 {
tmp16637 := Call1(__e, SymFn(symtail), let__14376)


let__14377 := tmp16637
_ = let__14377

tmp16647 := GConsP(__e, let__14377)

var ifres16638 Obj

if True == tmp16647 {
tmp16639 := Call1(__e, SymFn(symhead), let__14377)


let__14378 := tmp16639
_ = let__14378

tmp16640 := Call1(__e, SymFn(symtail), let__14377)


let__14379 := tmp16640
_ = let__14379

tmp16641 := GCons(__e, let__14375, Nil)

tmp16642 := GCons(__e, let__14378, tmp16641)

tmp16643 := GCons(__e, symwhere, tmp16642)

tmp16644 := Call2(__e, SymFn(sym_8p), let__14372, tmp16643)


tmp16645 := Call2(__e, SymFn(symshen_4comb), let__14379, tmp16644)


ifres16638 = tmp16645


} else {
tmp16646 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16638 = tmp16646


}

ifres16636 = ifres16638


} else {
tmp16648 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16636 = tmp16648


}

ifres16633 = ifres16636


} else {
tmp16650 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16633 = tmp16650


}

ifres16631 = ifres16633


} else {
tmp16652 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16631 = tmp16652


}

ifres16627 = ifres16631


}

let__14370 := ifres16627
_ = let__14370

tmp16727 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__14370)


if True == tmp16727 {
tmp16655 := Call1(__e, SymFn(symshen_4_5patterns_6), V5494)


let__14381 := tmp16655
_ = let__14381

tmp16684 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__14381)


var ifres16656 Obj

if True == tmp16684 {
tmp16657 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16656 = tmp16657


} else {
tmp16658 := Call1(__e, SymFn(symshen_4_5_1out), let__14381)


let__14382 := tmp16658
_ = let__14382

tmp16659 := Call1(__e, SymFn(symshen_4in_1_6), let__14381)


let__14383 := tmp16659
_ = let__14383

tmp16683 := Call2(__e, SymFn(symshen_4hds_a_2), let__14383, sym_5_1)


var ifres16660 Obj

if True == tmp16683 {
tmp16661 := Call1(__e, SymFn(symtail), let__14383)


let__14384 := tmp16661
_ = let__14384

tmp16681 := GConsP(__e, let__14384)

var ifres16662 Obj

if True == tmp16681 {
tmp16663 := Call1(__e, SymFn(symhead), let__14384)


let__14385 := tmp16663
_ = let__14385

tmp16664 := Call1(__e, SymFn(symtail), let__14384)


let__14386 := tmp16664
_ = let__14386

tmp16679 := Call2(__e, SymFn(symshen_4hds_a_2), let__14386, symwhere)


var ifres16665 Obj

if True == tmp16679 {
tmp16666 := Call1(__e, SymFn(symtail), let__14386)


let__14387 := tmp16666
_ = let__14387

tmp16677 := GConsP(__e, let__14387)

var ifres16667 Obj

if True == tmp16677 {
tmp16668 := Call1(__e, SymFn(symhead), let__14387)


let__14388 := tmp16668
_ = let__14388

tmp16669 := Call1(__e, SymFn(symtail), let__14387)


let__14389 := tmp16669
_ = let__14389

tmp16670 := GCons(__e, let__14385, Nil)

tmp16671 := GCons(__e, let__14388, tmp16670)

tmp16672 := GCons(__e, symwhere, tmp16671)

tmp16673 := Call1(__e, SymFn(symshen_4correct), tmp16672)


tmp16674 := Call2(__e, SymFn(sym_8p), let__14382, tmp16673)


tmp16675 := Call2(__e, SymFn(symshen_4comb), let__14389, tmp16674)


ifres16667 = tmp16675


} else {
tmp16676 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16667 = tmp16676


}

ifres16665 = ifres16667


} else {
tmp16678 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16665 = tmp16678


}

ifres16662 = ifres16665


} else {
tmp16680 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16662 = tmp16680


}

ifres16660 = ifres16662


} else {
tmp16682 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16660 = tmp16682


}

ifres16656 = ifres16660


}

let__14380 := ifres16656
_ = let__14380

tmp16725 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__14380)


if True == tmp16725 {
tmp16685 := Call1(__e, SymFn(symshen_4_5patterns_6), V5494)


let__14391 := tmp16685
_ = let__14391

tmp16702 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__14391)


var ifres16686 Obj

if True == tmp16702 {
tmp16687 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16686 = tmp16687


} else {
tmp16688 := Call1(__e, SymFn(symshen_4_5_1out), let__14391)


let__14392 := tmp16688
_ = let__14392

tmp16689 := Call1(__e, SymFn(symshen_4in_1_6), let__14391)


let__14393 := tmp16689
_ = let__14393

tmp16701 := Call2(__e, SymFn(symshen_4hds_a_2), let__14393, sym_5_1)


var ifres16690 Obj

if True == tmp16701 {
tmp16691 := Call1(__e, SymFn(symtail), let__14393)


let__14394 := tmp16691
_ = let__14394

tmp16699 := GConsP(__e, let__14394)

var ifres16692 Obj

if True == tmp16699 {
tmp16693 := Call1(__e, SymFn(symhead), let__14394)


let__14395 := tmp16693
_ = let__14395

tmp16694 := Call1(__e, SymFn(symtail), let__14394)


let__14396 := tmp16694
_ = let__14396

tmp16695 := Call1(__e, SymFn(symshen_4correct), let__14395)


tmp16696 := Call2(__e, SymFn(sym_8p), let__14392, tmp16695)


tmp16697 := Call2(__e, SymFn(symshen_4comb), let__14396, tmp16696)


ifres16692 = tmp16697


} else {
tmp16698 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16692 = tmp16698


}

ifres16690 = ifres16692


} else {
tmp16700 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16690 = tmp16700


}

ifres16686 = ifres16690


}

let__14390 := ifres16686
_ = let__14390

tmp16723 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__14390)


if True == tmp16723 {
tmp16703 := Call1(__e, SymFn(symshen_4_5patterns_6), V5494)


let__14398 := tmp16703
_ = let__14398

tmp16719 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__14398)


var ifres16704 Obj

if True == tmp16719 {
tmp16705 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16704 = tmp16705


} else {
tmp16706 := Call1(__e, SymFn(symshen_4_5_1out), let__14398)


let__14399 := tmp16706
_ = let__14399

tmp16707 := Call1(__e, SymFn(symshen_4in_1_6), let__14398)


let__14400 := tmp16707
_ = let__14400

tmp16718 := Call2(__e, SymFn(symshen_4hds_a_2), let__14400, sym_1_6)


var ifres16708 Obj

if True == tmp16718 {
tmp16709 := Call1(__e, SymFn(symtail), let__14400)


let__14401 := tmp16709
_ = let__14401

tmp16716 := GConsP(__e, let__14401)

var ifres16710 Obj

if True == tmp16716 {
tmp16711 := Call1(__e, SymFn(symhead), let__14401)


let__14402 := tmp16711
_ = let__14402

tmp16712 := Call1(__e, SymFn(symtail), let__14401)


let__14403 := tmp16712
_ = let__14403

tmp16713 := Call2(__e, SymFn(sym_8p), let__14399, let__14402)


tmp16714 := Call2(__e, SymFn(symshen_4comb), let__14403, tmp16713)


ifres16710 = tmp16714


} else {
tmp16715 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16710 = tmp16715


}

ifres16708 = ifres16710


} else {
tmp16717 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16708 = tmp16717


}

ifres16704 = ifres16708


}

let__14397 := ifres16704
_ = let__14397

tmp16721 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__14397)


if True == tmp16721 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__14397
}


} else {
return let__14390
}


} else {
return let__14380
}


} else {
return let__14370
}


panic("unreachable")
})

tmp16728 := Call2(__e, ns2_1set, symshen_4_5rule_d_6, tmp16625)


_ = tmp16728

tmp16729 := MakeNative1(func(__e *ControlFlow, V5529 Obj) Obj {
tmp16877 := GConsP(__e, V5529)

var ifres16821 Obj

if True == tmp16877 {
tmp16875 := GHd(__e, V5529)

tmp16876 := GEq(__e, symwhere, tmp16875)

var ifres16823 Obj

if True == tmp16876 {
tmp16873 := GTl(__e, V5529)

tmp16874 := GConsP(__e, tmp16873)

var ifres16825 Obj

if True == tmp16874 {
tmp16870 := GTl(__e, V5529)

tmp16871 := GTl(__e, tmp16870)

tmp16872 := GConsP(__e, tmp16871)

var ifres16827 Obj

if True == tmp16872 {
tmp16866 := GTl(__e, V5529)

tmp16867 := GTl(__e, tmp16866)

tmp16868 := GHd(__e, tmp16867)

tmp16869 := GConsP(__e, tmp16868)

var ifres16829 Obj

if True == tmp16869 {
tmp16861 := GTl(__e, V5529)

tmp16862 := GTl(__e, tmp16861)

tmp16863 := GHd(__e, tmp16862)

tmp16864 := GHd(__e, tmp16863)

tmp16865 := GEq(__e, symfail_1if, tmp16864)

var ifres16831 Obj

if True == tmp16865 {
tmp16856 := GTl(__e, V5529)

tmp16857 := GTl(__e, tmp16856)

tmp16858 := GHd(__e, tmp16857)

tmp16859 := GTl(__e, tmp16858)

tmp16860 := GConsP(__e, tmp16859)

var ifres16833 Obj

if True == tmp16860 {
tmp16850 := GTl(__e, V5529)

tmp16851 := GTl(__e, tmp16850)

tmp16852 := GHd(__e, tmp16851)

tmp16853 := GTl(__e, tmp16852)

tmp16854 := GTl(__e, tmp16853)

tmp16855 := GConsP(__e, tmp16854)

var ifres16835 Obj

if True == tmp16855 {
tmp16843 := GTl(__e, V5529)

tmp16844 := GTl(__e, tmp16843)

tmp16845 := GHd(__e, tmp16844)

tmp16846 := GTl(__e, tmp16845)

tmp16847 := GTl(__e, tmp16846)

tmp16848 := GTl(__e, tmp16847)

tmp16849 := GEq(__e, Nil, tmp16848)

var ifres16837 Obj

if True == tmp16849 {
tmp16839 := GTl(__e, V5529)

tmp16840 := GTl(__e, tmp16839)

tmp16841 := GTl(__e, tmp16840)

tmp16842 := GEq(__e, Nil, tmp16841)

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
tmp16730 := GTl(__e, V5529)

tmp16731 := GHd(__e, tmp16730)

tmp16732 := GTl(__e, V5529)

tmp16733 := GTl(__e, tmp16732)

tmp16734 := GHd(__e, tmp16733)

tmp16735 := GTl(__e, tmp16734)

tmp16736 := GCons(__e, tmp16735, Nil)

tmp16737 := GCons(__e, symnot, tmp16736)

tmp16738 := GCons(__e, tmp16737, Nil)

tmp16739 := GCons(__e, tmp16731, tmp16738)

tmp16740 := GCons(__e, symand, tmp16739)

tmp16741 := GTl(__e, V5529)

tmp16742 := GTl(__e, tmp16741)

tmp16743 := GHd(__e, tmp16742)

tmp16744 := GTl(__e, tmp16743)

tmp16745 := GTl(__e, tmp16744)

tmp16746 := GCons(__e, tmp16740, tmp16745)

return GCons(__e, symwhere, tmp16746)


} else {
tmp16819 := GConsP(__e, V5529)

var ifres16800 Obj

if True == tmp16819 {
tmp16817 := GHd(__e, V5529)

tmp16818 := GEq(__e, symwhere, tmp16817)

var ifres16802 Obj

if True == tmp16818 {
tmp16815 := GTl(__e, V5529)

tmp16816 := GConsP(__e, tmp16815)

var ifres16804 Obj

if True == tmp16816 {
tmp16812 := GTl(__e, V5529)

tmp16813 := GTl(__e, tmp16812)

tmp16814 := GConsP(__e, tmp16813)

var ifres16806 Obj

if True == tmp16814 {
tmp16808 := GTl(__e, V5529)

tmp16809 := GTl(__e, tmp16808)

tmp16810 := GTl(__e, tmp16809)

tmp16811 := GEq(__e, Nil, tmp16810)

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
tmp16747 := GTl(__e, V5529)

tmp16748 := GHd(__e, tmp16747)

tmp16749 := GTl(__e, V5529)

tmp16750 := GTl(__e, tmp16749)

tmp16751 := GHd(__e, tmp16750)

tmp16752 := GCons(__e, symfail, Nil)

tmp16753 := GCons(__e, tmp16752, Nil)

tmp16754 := GCons(__e, tmp16751, tmp16753)

tmp16755 := GCons(__e, sym_a_a, tmp16754)

tmp16756 := GCons(__e, tmp16755, Nil)

tmp16757 := GCons(__e, symnot, tmp16756)

tmp16758 := GCons(__e, tmp16757, Nil)

tmp16759 := GCons(__e, tmp16748, tmp16758)

tmp16760 := GCons(__e, symand, tmp16759)

tmp16761 := GTl(__e, V5529)

tmp16762 := GTl(__e, tmp16761)

tmp16763 := GCons(__e, tmp16760, tmp16762)

return GCons(__e, symwhere, tmp16763)


} else {
tmp16798 := GConsP(__e, V5529)

var ifres16779 Obj

if True == tmp16798 {
tmp16796 := GHd(__e, V5529)

tmp16797 := GEq(__e, symfail_1if, tmp16796)

var ifres16781 Obj

if True == tmp16797 {
tmp16794 := GTl(__e, V5529)

tmp16795 := GConsP(__e, tmp16794)

var ifres16783 Obj

if True == tmp16795 {
tmp16791 := GTl(__e, V5529)

tmp16792 := GTl(__e, tmp16791)

tmp16793 := GConsP(__e, tmp16792)

var ifres16785 Obj

if True == tmp16793 {
tmp16787 := GTl(__e, V5529)

tmp16788 := GTl(__e, tmp16787)

tmp16789 := GTl(__e, tmp16788)

tmp16790 := GEq(__e, Nil, tmp16789)

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
tmp16764 := GTl(__e, V5529)

tmp16765 := GCons(__e, tmp16764, Nil)

tmp16766 := GCons(__e, symnot, tmp16765)

tmp16767 := GTl(__e, V5529)

tmp16768 := GTl(__e, tmp16767)

tmp16769 := GCons(__e, tmp16766, tmp16768)

return GCons(__e, symwhere, tmp16769)


} else {
tmp16770 := GCons(__e, symfail, Nil)

tmp16771 := GCons(__e, tmp16770, Nil)

tmp16772 := GCons(__e, V5529, tmp16771)

tmp16773 := GCons(__e, sym_a_a, tmp16772)

tmp16774 := GCons(__e, tmp16773, Nil)

tmp16775 := GCons(__e, symnot, tmp16774)

tmp16776 := GCons(__e, V5529, Nil)

tmp16777 := GCons(__e, tmp16775, tmp16776)

return GCons(__e, symwhere, tmp16777)


}


}


}


panic("unreachable")
})

tmp16878 := Call2(__e, ns2_1set, symshen_4correct, tmp16729)


_ = tmp16878

tmp16879 := MakeNative(func(__e *ControlFlow) {
V5530 := __e.Get(1)
V5531 := __e.Get(2)
V5532 := __e.Get(3)
V5533 := __e.Get(4)
V5534 := __e.Get(5)
V5535 := __e.Get(6)
V5536 := __e.Get(7)
V5537 := __e.Get(8)
__e.ReturnDirect(func() Obj {
let__14404 := GAdd(__e, V5536, MakeInteger(1))
_ = let__14404

tmp16887 := Call1(__e, SymFn(symshen_4unlocked_2), V5535)


var ifres16881 Obj

if True == tmp16887 {
tmp16882 := Call2(__e, SymFn(symshen_4lazyderef), V5531, V5534)


let__14406 := tmp16882
_ = let__14406

tmp16886 := GEq(__e, let__14406, Nil)

var ifres16883 Obj

if True == tmp16886 {
tmp16884 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16884

tmp16885 := Call1(__e, SymFn(symthaw), V5537)


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

tmp16912 := GEq(__e, let__14405, False)

if True == tmp16912 {
tmp16908 := Call1(__e, SymFn(symshen_4unlocked_2), V5535)


var ifres16888 Obj

if True == tmp16908 {
tmp16889 := Call2(__e, SymFn(symshen_4lazyderef), V5531, V5534)


let__14408 := tmp16889
_ = let__14408

tmp16907 := GConsP(__e, let__14408)

var ifres16890 Obj

if True == tmp16907 {
tmp16891 := GHd(__e, let__14408)

let__14409 := tmp16891
_ = let__14409

tmp16892 := GTl(__e, let__14408)

let__14410 := tmp16892
_ = let__14410

tmp16893 := Call1(__e, SymFn(symshen_4newpv), V5534)


let__14411 := tmp16893
_ = let__14411

tmp16894 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16894

tmp16895 := Call2(__e, SymFn(symshen_4deref), let__14409, V5534)


tmp16896 := Call1(__e, SymFn(symshen_4freshen_1rule), tmp16895)


tmp16897 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16898 := Call2(__e, SymFn(symshen_4lazyderef), let__14411, V5534)


tmp16899 := Call1(__e, SymFn(symfst), tmp16898)


tmp16900 := Call2(__e, SymFn(symshen_4lazyderef), let__14411, V5534)


tmp16901 := Call1(__e, SymFn(symsnd), tmp16900)


tmp16902 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16903 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4t_d_1rules), V5530, let__14410, V5532, GAdd(__e, V5533, MakeInteger(1)), V5534, V5535, let__14404, V5537)
return nil


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5534, V5535, let__14404, tmp16903)


panic("unreachable")
})

__e.TailApply(SymFn(symshen_4t_d_1rule), V5530, V5533, tmp16899, tmp16901, V5532, V5534, V5535, let__14404, tmp16902)
return nil


panic("unreachable")
})

tmp16905 := Call6(__e, SymFn(symbind), let__14411, tmp16896, V5534, V5535, let__14404, tmp16897)


tmp16906 := Call2(__e, SymFn(symshen_4gc), V5534, tmp16905)


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

tmp16910 := GEq(__e, let__14407, False)

if True == tmp16910 {
return TailCall2(__e, SymFn(symshen_4unlock), V5535, let__14404)
} else {
return let__14407
}


} else {
return let__14405
}


panic("unreachable")
}())
}, 8)

tmp16913 := Call2(__e, ns2_1set, symshen_4t_d_1rules, tmp16879)


_ = tmp16913

tmp16914 := MakeNative1(func(__e *ControlFlow, V5546 Obj) Obj {
tmp16925 := Call1(__e, SymFn(symtuple_2), V5546)


if True == tmp16925 {
tmp16915 := Call1(__e, SymFn(symfst), V5546)


tmp16916 := Call1(__e, SymFn(symshen_4extract_1vars), tmp16915)


let__14412 := tmp16916
_ = let__14412

tmp16917 := MakeNative1(func(__e *ControlFlow, Z5549 Obj) Obj {
tmp16918 := Call1(__e, SymFn(symshen_4freshterm), Z5549)


return GCons(__e, Z5549, tmp16918)


panic("unreachable")
})

tmp16919 := Call2(__e, SymFn(symmap), tmp16917, let__14412)


let__14413 := tmp16919
_ = let__14413

tmp16920 := Call1(__e, SymFn(symfst), V5546)


tmp16921 := Call2(__e, SymFn(symshen_4freshen), let__14413, tmp16920)


tmp16922 := Call1(__e, SymFn(symsnd), V5546)


tmp16923 := Call2(__e, SymFn(symshen_4freshen), let__14413, tmp16922)


return TailCall2(__e, SymFn(sym_8p), tmp16921, tmp16923)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4freshen_1rule)
}


panic("unreachable")
})

tmp16926 := Call2(__e, ns2_1set, symshen_4freshen_1rule, tmp16914)


_ = tmp16926

tmp16927 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V5550 := __self1
_ = V5550
V5551 := __self2
_ = V5551
tmp16941 := GEq(__e, Nil, V5550)

if True == tmp16941 {
return V5551
} else {
tmp16939 := GConsP(__e, V5550)

var ifres16935 Obj

if True == tmp16939 {
tmp16937 := GHd(__e, V5550)

tmp16938 := GConsP(__e, tmp16937)

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
tmp16928 := GTl(__e, V5550)

tmp16929 := GHd(__e, V5550)

tmp16930 := GHd(__e, tmp16929)

tmp16931 := GHd(__e, V5550)

tmp16932 := GTl(__e, tmp16931)

tmp16933 := Call3(__e, SymFn(symshen_4beta), tmp16930, tmp16932, V5551)


if SymFn(symshen_4freshen) == __self {
__self1, __self2 = tmp16928, tmp16933
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4freshen), tmp16928, tmp16933)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4freshen)
}


}


panic("unreachable")
})
return __self
})()

tmp16942 := Call2(__e, ns2_1set, symshen_4freshen, tmp16927)


_ = tmp16942

tmp16943 := MakeNative(func(__e *ControlFlow) {
V5552 := __e.Get(1)
V5553 := __e.Get(2)
V5554 := __e.Get(3)
V5555 := __e.Get(4)
V5556 := __e.Get(5)
V5557 := __e.Get(6)
V5558 := __e.Get(7)
V5559 := __e.Get(8)
V5560 := __e.Get(9)
__e.ReturnDirect(func() Obj {
tmp16947 := Call1(__e, SymFn(symshen_4unlocked_2), V5558)


var ifres16944 Obj

if True == tmp16947 {
tmp16945 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16945

tmp16946 := Call(__e, SymFn(symshen_4t_d_1rule_1h), V5554, V5555, V5556, V5557, V5558, V5559, V5560)


ifres16944 = tmp16946


} else {
ifres16944 = False


}

let__14414 := ifres16944
_ = let__14414

tmp16959 := GEq(__e, let__14414, False)

if True == tmp16959 {
tmp16957 := Call1(__e, SymFn(symshen_4unlocked_2), V5558)


if True == tmp16957 {
tmp16948 := Call1(__e, SymFn(symshen_4newpv), V5557)


let__14415 := tmp16948
_ = let__14415

tmp16949 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16949

tmp16950 := Call3(__e, SymFn(symshen_4app), V5552, MakeString("\n"), symshen_4a)


tmp16952 := Call3(__e, SymFn(symshen_4app), V5553, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" of "), tmp16950), symshen_4a)


tmp16954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("type error in rule "), tmp16952))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("type error in rule "), tmp16952)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()

tmp16955 := Call6(__e, SymFn(symbind), let__14415, tmp16954, V5557, V5558, V5559, V5560)


return TailCall2(__e, SymFn(symshen_4gc), V5557, tmp16955)


} else {
return False
}


} else {
return let__14414
}


panic("unreachable")
}())
}, 9)

tmp16960 := Call2(__e, ns2_1set, symshen_4t_d_1rule, tmp16943)


_ = tmp16960

tmp16961 := MakeNative(func(__e *ControlFlow) {
V5563 := __e.Get(1)
V5564 := __e.Get(2)
V5565 := __e.Get(3)
V5566 := __e.Get(4)
V5567 := __e.Get(5)
V5568 := __e.Get(6)
V5569 := __e.Get(7)
__e.ReturnDirect(func() Obj {
let__14416 := GAdd(__e, V5568, MakeInteger(1))
_ = let__14416

tmp16986 := Call1(__e, SymFn(symshen_4unlocked_2), V5567)


var ifres16963 Obj

if True == tmp16986 {
tmp16964 := Call2(__e, SymFn(symshen_4lazyderef), V5563, V5566)


let__14418 := tmp16964
_ = let__14418

tmp16985 := GEq(__e, let__14418, Nil)

var ifres16965 Obj

if True == tmp16985 {
tmp16966 := Call2(__e, SymFn(symshen_4lazyderef), V5565, V5566)


let__14419 := tmp16966
_ = let__14419

tmp16984 := GConsP(__e, let__14419)

var ifres16967 Obj

if True == tmp16984 {
tmp16968 := GHd(__e, let__14419)

tmp16969 := Call2(__e, SymFn(symshen_4lazyderef), tmp16968, V5566)


let__14420 := tmp16969
_ = let__14420

tmp16983 := GEq(__e, let__14420, sym_1_1_6)

var ifres16970 Obj

if True == tmp16983 {
tmp16971 := GTl(__e, let__14419)

tmp16972 := Call2(__e, SymFn(symshen_4lazyderef), tmp16971, V5566)


let__14421 := tmp16972
_ = let__14421

tmp16982 := GConsP(__e, let__14421)

var ifres16973 Obj

if True == tmp16982 {
tmp16974 := GHd(__e, let__14421)

let__14422 := tmp16974
_ = let__14422

tmp16975 := GTl(__e, let__14421)

tmp16976 := Call2(__e, SymFn(symshen_4lazyderef), tmp16975, V5566)


let__14423 := tmp16976
_ = let__14423

tmp16981 := GEq(__e, let__14423, Nil)

var ifres16977 Obj

if True == tmp16981 {
tmp16978 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16978

tmp16979 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4t_d_1correct), V5564, let__14422, Nil, V5566, V5567, let__14416, V5569)
return nil
panic("unreachable")
})

tmp16980 := Call4(__e, SymFn(symshen_4cut), V5566, V5567, let__14416, tmp16979)


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

tmp17005 := GEq(__e, let__14417, False)

if True == tmp17005 {
tmp17001 := Call1(__e, SymFn(symshen_4unlocked_2), V5567)


var ifres16987 Obj

if True == tmp17001 {
tmp16988 := Call1(__e, SymFn(symshen_4newpv), V5566)


let__14425 := tmp16988
_ = let__14425

tmp16989 := Call1(__e, SymFn(symshen_4newpv), V5566)


let__14426 := tmp16989
_ = let__14426

tmp16990 := Call1(__e, SymFn(symshen_4newpv), V5566)


let__14427 := tmp16990
_ = let__14427

tmp16991 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16991

tmp16992 := Call1(__e, SymFn(symshen_4freshterms), V5563)


tmp16993 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16994 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16995 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16996 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4t_d_1correct), V5564, let__14426, let__14427, V5566, V5567, let__14416, V5569)
return nil
panic("unreachable")
})

__e.TailApply(SymFn(symshen_4myassume), V5563, V5565, let__14427, V5566, V5567, let__14416, tmp16996)
return nil


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5566, V5567, let__14416, tmp16995)


panic("unreachable")
})

__e.TailApply(SymFn(symshen_4t_d_1integrity), V5563, V5565, let__14425, let__14426, V5566, V5567, let__14416, tmp16994)
return nil


panic("unreachable")
})

tmp16997 := Call6(__e, SymFn(symshen_4p_1hyps), tmp16992, let__14425, V5566, V5567, let__14416, tmp16993)


tmp16998 := Call2(__e, SymFn(symshen_4gc), V5566, tmp16997)


tmp16999 := Call2(__e, SymFn(symshen_4gc), V5566, tmp16998)


tmp17000 := Call2(__e, SymFn(symshen_4gc), V5566, tmp16999)


ifres16987 = tmp17000


} else {
ifres16987 = False


}

let__14424 := ifres16987
_ = let__14424

tmp17003 := GEq(__e, let__14424, False)

if True == tmp17003 {
return TailCall2(__e, SymFn(symshen_4unlock), V5567, let__14416)
} else {
return let__14424
}


} else {
return let__14417
}


panic("unreachable")
}())
}, 7)

tmp17006 := Call2(__e, ns2_1set, symshen_4t_d_1rule_1h, tmp16961)


_ = tmp17006

tmp17007 := MakeNative(func(__e *ControlFlow) {
V5582 := __e.Get(1)
V5583 := __e.Get(2)
V5584 := __e.Get(3)
V5585 := __e.Get(4)
V5586 := __e.Get(5)
V5587 := __e.Get(6)
V5588 := __e.Get(7)
__e.ReturnDirect(func() Obj {
tmp17021 := Call1(__e, SymFn(symshen_4unlocked_2), V5586)


var ifres17008 Obj

if True == tmp17021 {
tmp17009 := Call2(__e, SymFn(symshen_4lazyderef), V5582, V5585)


let__14429 := tmp17009
_ = let__14429

tmp17020 := GEq(__e, let__14429, Nil)

var ifres17010 Obj

if True == tmp17020 {
tmp17011 := Call2(__e, SymFn(symshen_4lazyderef), V5584, V5585)


let__14430 := tmp17011
_ = let__14430

tmp17012 := MakeNative0(func(__e *ControlFlow) Obj {
tmp17013 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp17013

return TailCall1(__e, SymFn(symthaw), V5588)


panic("unreachable")
})

let__14431 := tmp17012
_ = let__14431

tmp17019 := GEq(__e, let__14430, Nil)

var ifres17014 Obj

if True == tmp17019 {
tmp17015 := Call1(__e, SymFn(symthaw), let__14431)


ifres17014 = tmp17015


} else {
tmp17018 := Call1(__e, SymFn(symshen_4pvar_2), let__14430)


var ifres17016 Obj

if True == tmp17018 {
tmp17017 := Call4(__e, SymFn(symshen_4bind_b), let__14430, Nil, V5585, let__14431)


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

tmp17140 := GEq(__e, let__14428, False)

if True == tmp17140 {
tmp17138 := Call1(__e, SymFn(symshen_4unlocked_2), V5586)


if True == tmp17138 {
tmp17022 := Call2(__e, SymFn(symshen_4lazyderef), V5582, V5585)


let__14432 := tmp17022
_ = let__14432

tmp17136 := GConsP(__e, let__14432)

if True == tmp17136 {
tmp17023 := GHd(__e, let__14432)

let__14433 := tmp17023
_ = let__14433

tmp17024 := GTl(__e, let__14432)

let__14434 := tmp17024
_ = let__14434

tmp17025 := Call2(__e, SymFn(symshen_4lazyderef), V5583, V5585)


let__14435 := tmp17025
_ = let__14435

tmp17134 := GConsP(__e, let__14435)

if True == tmp17134 {
tmp17026 := GHd(__e, let__14435)

let__14436 := tmp17026
_ = let__14436

tmp17027 := GTl(__e, let__14435)

tmp17028 := Call2(__e, SymFn(symshen_4lazyderef), tmp17027, V5585)


let__14437 := tmp17028
_ = let__14437

tmp17132 := GConsP(__e, let__14437)

if True == tmp17132 {
tmp17029 := GHd(__e, let__14437)

tmp17030 := Call2(__e, SymFn(symshen_4lazyderef), tmp17029, V5585)


let__14438 := tmp17030
_ = let__14438

tmp17130 := GEq(__e, let__14438, sym_1_1_6)

if True == tmp17130 {
tmp17031 := GTl(__e, let__14437)

tmp17032 := Call2(__e, SymFn(symshen_4lazyderef), tmp17031, V5585)


let__14439 := tmp17032
_ = let__14439

tmp17128 := GConsP(__e, let__14439)

if True == tmp17128 {
tmp17033 := GHd(__e, let__14439)

let__14440 := tmp17033
_ = let__14440

tmp17034 := GTl(__e, let__14439)

tmp17035 := Call2(__e, SymFn(symshen_4lazyderef), tmp17034, V5585)


let__14441 := tmp17035
_ = let__14441

tmp17126 := GEq(__e, let__14441, Nil)

if True == tmp17126 {
tmp17036 := Call2(__e, SymFn(symshen_4lazyderef), V5584, V5585)


let__14442 := tmp17036
_ = let__14442

tmp17037 := MakeNative1(func(__e *ControlFlow, Z5605 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5606 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5607 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5608 Obj) Obj {
tmp17038 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp17038

tmp17039 := MakeNative0(func(__e *ControlFlow) Obj {
tmp17040 := MakeNative0(func(__e *ControlFlow) Obj {
tmp17041 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp17042 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4myassume), let__14434, let__14440, Z5608, V5585, V5586, V5587, V5588)
return nil
panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), Z5606, tmp17041, V5585, V5586, V5587, tmp17042)


panic("unreachable")
})

return TailCall6(__e, SymFn(symis_b), let__14433, Z5605, V5585, V5586, V5587, tmp17040)


panic("unreachable")
})

return TailCall6(__e, SymFn(symis_b), let__14436, Z5607, V5585, V5586, V5587, tmp17039)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

let__14443 := tmp17037
_ = let__14443

tmp17124 := GConsP(__e, let__14442)

if True == tmp17124 {
tmp17043 := GHd(__e, let__14442)

tmp17044 := Call2(__e, SymFn(symshen_4lazyderef), tmp17043, V5585)


let__14444 := tmp17044
_ = let__14444

tmp17045 := MakeNative1(func(__e *ControlFlow, Z5611 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5612 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5613 Obj) Obj {
tmp17046 := GTl(__e, let__14442)

let__14446 := tmp17046
_ = let__14446

tmp17047 := Call1(__e, let__14443, Z5611)


tmp17048 := Call1(__e, tmp17047, Z5612)


tmp17049 := Call1(__e, tmp17048, Z5613)


return TailCall1(__e, tmp17049, let__14446)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

let__14445 := tmp17045
_ = let__14445

tmp17104 := GConsP(__e, let__14444)

if True == tmp17104 {
tmp17050 := GHd(__e, let__14444)

let__14447 := tmp17050
_ = let__14447

tmp17051 := GTl(__e, let__14444)

tmp17052 := Call2(__e, SymFn(symshen_4lazyderef), tmp17051, V5585)


let__14448 := tmp17052
_ = let__14448

tmp17053 := MakeNative1(func(__e *ControlFlow, Z5618 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5619 Obj) Obj {
tmp17054 := Call1(__e, let__14445, let__14447)


tmp17055 := Call1(__e, tmp17054, Z5618)


return TailCall1(__e, tmp17055, Z5619)


panic("unreachable")
})
panic("unreachable")
})

let__14449 := tmp17053
_ = let__14449

tmp17088 := GConsP(__e, let__14448)

if True == tmp17088 {
tmp17056 := GHd(__e, let__14448)

let__14450 := tmp17056
_ = let__14450

tmp17057 := GTl(__e, let__14448)

tmp17058 := Call2(__e, SymFn(symshen_4lazyderef), tmp17057, V5585)


let__14451 := tmp17058
_ = let__14451

tmp17059 := MakeNative1(func(__e *ControlFlow, Z5623 Obj) Obj {
tmp17060 := Call1(__e, let__14449, let__14450)


return TailCall1(__e, tmp17060, Z5623)


panic("unreachable")
})

let__14452 := tmp17059
_ = let__14452

tmp17076 := GConsP(__e, let__14451)

if True == tmp17076 {
tmp17061 := GHd(__e, let__14451)

let__14453 := tmp17061
_ = let__14453

tmp17062 := GTl(__e, let__14451)

tmp17063 := Call2(__e, SymFn(symshen_4lazyderef), tmp17062, V5585)


let__14454 := tmp17063
_ = let__14454

tmp17064 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__14452, let__14453)
panic("unreachable")
})

let__14455 := tmp17064
_ = let__14455

tmp17068 := GEq(__e, let__14454, Nil)

if True == tmp17068 {
return TailCall1(__e, SymFn(symthaw), let__14455)
} else {
tmp17066 := Call1(__e, SymFn(symshen_4pvar_2), let__14454)


if True == tmp17066 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__14454, Nil, V5585, let__14455)
} else {
return False
}


}


} else {
tmp17074 := Call1(__e, SymFn(symshen_4pvar_2), let__14451)


if True == tmp17074 {
tmp17069 := Call1(__e, SymFn(symshen_4newpv), V5585)


let__14456 := tmp17069
_ = let__14456

tmp17070 := GCons(__e, let__14456, Nil)

tmp17071 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__14452, let__14456)
panic("unreachable")
})

tmp17072 := Call4(__e, SymFn(symshen_4bind_b), let__14451, tmp17070, V5585, tmp17071)


return TailCall2(__e, SymFn(symshen_4gc), V5585, tmp17072)


} else {
return False
}


}


} else {
tmp17086 := Call1(__e, SymFn(symshen_4pvar_2), let__14448)


if True == tmp17086 {
tmp17077 := Call1(__e, SymFn(symshen_4newpv), V5585)


let__14457 := tmp17077
_ = let__14457

tmp17078 := Call1(__e, SymFn(symshen_4newpv), V5585)


let__14458 := tmp17078
_ = let__14458

tmp17079 := GCons(__e, let__14458, Nil)

tmp17080 := GCons(__e, let__14457, tmp17079)

tmp17081 := MakeNative0(func(__e *ControlFlow) Obj {
tmp17082 := Call1(__e, let__14449, let__14457)


return TailCall1(__e, tmp17082, let__14458)


panic("unreachable")
})

tmp17083 := Call4(__e, SymFn(symshen_4bind_b), let__14448, tmp17080, V5585, tmp17081)


tmp17084 := Call2(__e, SymFn(symshen_4gc), V5585, tmp17083)


return TailCall2(__e, SymFn(symshen_4gc), V5585, tmp17084)


} else {
return False
}


}


} else {
tmp17102 := Call1(__e, SymFn(symshen_4pvar_2), let__14444)


if True == tmp17102 {
tmp17089 := Call1(__e, SymFn(symshen_4newpv), V5585)


let__14459 := tmp17089
_ = let__14459

tmp17090 := Call1(__e, SymFn(symshen_4newpv), V5585)


let__14460 := tmp17090
_ = let__14460

tmp17091 := Call1(__e, SymFn(symshen_4newpv), V5585)


let__14461 := tmp17091
_ = let__14461

tmp17092 := GCons(__e, let__14461, Nil)

tmp17093 := GCons(__e, let__14460, tmp17092)

tmp17094 := GCons(__e, let__14459, tmp17093)

tmp17095 := MakeNative0(func(__e *ControlFlow) Obj {
tmp17096 := Call1(__e, let__14445, let__14459)


tmp17097 := Call1(__e, tmp17096, let__14460)


return TailCall1(__e, tmp17097, let__14461)


panic("unreachable")
})

tmp17098 := Call4(__e, SymFn(symshen_4bind_b), let__14444, tmp17094, V5585, tmp17095)


tmp17099 := Call2(__e, SymFn(symshen_4gc), V5585, tmp17098)


tmp17100 := Call2(__e, SymFn(symshen_4gc), V5585, tmp17099)


return TailCall2(__e, SymFn(symshen_4gc), V5585, tmp17100)


} else {
return False
}


}


} else {
tmp17122 := Call1(__e, SymFn(symshen_4pvar_2), let__14442)


if True == tmp17122 {
tmp17105 := Call1(__e, SymFn(symshen_4newpv), V5585)


let__14462 := tmp17105
_ = let__14462

tmp17106 := Call1(__e, SymFn(symshen_4newpv), V5585)


let__14463 := tmp17106
_ = let__14463

tmp17107 := Call1(__e, SymFn(symshen_4newpv), V5585)


let__14464 := tmp17107
_ = let__14464

tmp17108 := Call1(__e, SymFn(symshen_4newpv), V5585)


let__14465 := tmp17108
_ = let__14465

tmp17109 := GCons(__e, let__14464, Nil)

tmp17110 := GCons(__e, let__14463, tmp17109)

tmp17111 := GCons(__e, let__14462, tmp17110)

tmp17112 := GCons(__e, tmp17111, let__14465)

tmp17113 := MakeNative0(func(__e *ControlFlow) Obj {
tmp17114 := Call1(__e, let__14443, let__14462)


tmp17115 := Call1(__e, tmp17114, let__14463)


tmp17116 := Call1(__e, tmp17115, let__14464)


return TailCall1(__e, tmp17116, let__14465)


panic("unreachable")
})

tmp17117 := Call4(__e, SymFn(symshen_4bind_b), let__14442, tmp17112, V5585, tmp17113)


tmp17118 := Call2(__e, SymFn(symshen_4gc), V5585, tmp17117)


tmp17119 := Call2(__e, SymFn(symshen_4gc), V5585, tmp17118)


tmp17120 := Call2(__e, SymFn(symshen_4gc), V5585, tmp17119)


return TailCall2(__e, SymFn(symshen_4gc), V5585, tmp17120)


} else {
return False
}


}


} else {
return False
}


} else {
return False
}


} else {
return False
}


} else {
return False
}


} else {
return False
}


} else {
return False
}


} else {
return False
}


} else {
return let__14428
}


panic("unreachable")
}())
}, 7)

tmp17141 := Call2(__e, ns2_1set, symshen_4myassume, tmp17007)


_ = tmp17141

tmp17142 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V5639 := __self1
_ = V5639
tmp17165 := GEq(__e, Nil, V5639)

if True == tmp17165 {
return Nil
} else {
tmp17163 := GConsP(__e, V5639)

var ifres17159 Obj

if True == tmp17163 {
tmp17161 := GHd(__e, V5639)

tmp17162 := GConsP(__e, tmp17161)

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
tmp17143 := GHd(__e, V5639)

tmp17144 := GTl(__e, V5639)

tmp17145 := Call2(__e, SymFn(symappend), tmp17143, tmp17144)


if SymFn(symshen_4freshterms) == __self {
__self1 = tmp17145
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4freshterms), tmp17145)


} else {
tmp17157 := GConsP(__e, V5639)

var ifres17153 Obj

if True == tmp17157 {
tmp17155 := GHd(__e, V5639)

tmp17156 := Call1(__e, SymFn(symshen_4freshterm_2), tmp17155)


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
tmp17146 := GHd(__e, V5639)

tmp17147 := GTl(__e, V5639)

tmp17148 := Call1(__e, SymFn(symshen_4freshterms), tmp17147)


return TailCall2(__e, SymFn(symadjoin), tmp17146, tmp17148)


} else {
tmp17151 := GConsP(__e, V5639)

if True == tmp17151 {
tmp17149 := GTl(__e, V5639)

if SymFn(symshen_4freshterms) == __self {
__self1 = tmp17149
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4freshterms), tmp17149)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4freshterms)
}


}


}


}


panic("unreachable")
})
return __self
})()

tmp17166 := Call2(__e, ns2_1set, symshen_4freshterms, tmp17142)


_ = tmp17166

tmp17167 := MakeNative6(func(__e *ControlFlow, V5640, V5641, V5642, V5643, V5644, V5645 Obj) Obj {
tmp17181 := Call1(__e, SymFn(symshen_4unlocked_2), V5643)


var ifres17168 Obj

if True == tmp17181 {
tmp17169 := Call2(__e, SymFn(symshen_4lazyderef), V5640, V5642)


let__14467 := tmp17169
_ = let__14467

tmp17180 := GEq(__e, let__14467, Nil)

var ifres17170 Obj

if True == tmp17180 {
tmp17171 := Call2(__e, SymFn(symshen_4lazyderef), V5641, V5642)


let__14468 := tmp17171
_ = let__14468

tmp17172 := MakeNative0(func(__e *ControlFlow) Obj {
tmp17173 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp17173

return TailCall1(__e, SymFn(symthaw), V5645)


panic("unreachable")
})

let__14469 := tmp17172
_ = let__14469

tmp17179 := GEq(__e, let__14468, Nil)

var ifres17174 Obj

if True == tmp17179 {
tmp17175 := Call1(__e, SymFn(symthaw), let__14469)


ifres17174 = tmp17175


} else {
tmp17178 := Call1(__e, SymFn(symshen_4pvar_2), let__14468)


var ifres17176 Obj

if True == tmp17178 {
tmp17177 := Call4(__e, SymFn(symshen_4bind_b), let__14468, Nil, V5642, let__14469)


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

tmp17278 := GEq(__e, let__14466, False)

if True == tmp17278 {
tmp17276 := Call1(__e, SymFn(symshen_4unlocked_2), V5643)


if True == tmp17276 {
tmp17182 := Call2(__e, SymFn(symshen_4lazyderef), V5640, V5642)


let__14470 := tmp17182
_ = let__14470

tmp17274 := GConsP(__e, let__14470)

if True == tmp17274 {
tmp17183 := GHd(__e, let__14470)

let__14471 := tmp17183
_ = let__14471

tmp17184 := GTl(__e, let__14470)

let__14472 := tmp17184
_ = let__14472

tmp17185 := Call2(__e, SymFn(symshen_4lazyderef), V5641, V5642)


let__14473 := tmp17185
_ = let__14473

tmp17186 := MakeNative1(func(__e *ControlFlow, Z5655 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5656 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5657 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5658 Obj) Obj {
tmp17187 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp17187

tmp17188 := MakeNative0(func(__e *ControlFlow) Obj {
tmp17189 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp17190 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall6(__e, SymFn(symshen_4p_1hyps), let__14472, Z5658, V5642, V5643, V5644, V5645)
panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), Z5656, tmp17189, V5642, V5643, V5644, tmp17190)


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), Z5655, let__14471, V5642, V5643, V5644, tmp17188)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

let__14474 := tmp17186
_ = let__14474

tmp17272 := GConsP(__e, let__14473)

if True == tmp17272 {
tmp17191 := GHd(__e, let__14473)

tmp17192 := Call2(__e, SymFn(symshen_4lazyderef), tmp17191, V5642)


let__14475 := tmp17192
_ = let__14475

tmp17193 := MakeNative1(func(__e *ControlFlow, Z5661 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5662 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5663 Obj) Obj {
tmp17194 := GTl(__e, let__14473)

let__14477 := tmp17194
_ = let__14477

tmp17195 := Call1(__e, let__14474, Z5661)


tmp17196 := Call1(__e, tmp17195, Z5662)


tmp17197 := Call1(__e, tmp17196, Z5663)


return TailCall1(__e, tmp17197, let__14477)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

let__14476 := tmp17193
_ = let__14476

tmp17252 := GConsP(__e, let__14475)

if True == tmp17252 {
tmp17198 := GHd(__e, let__14475)

let__14478 := tmp17198
_ = let__14478

tmp17199 := GTl(__e, let__14475)

tmp17200 := Call2(__e, SymFn(symshen_4lazyderef), tmp17199, V5642)


let__14479 := tmp17200
_ = let__14479

tmp17201 := MakeNative1(func(__e *ControlFlow, Z5668 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5669 Obj) Obj {
tmp17202 := Call1(__e, let__14476, let__14478)


tmp17203 := Call1(__e, tmp17202, Z5668)


return TailCall1(__e, tmp17203, Z5669)


panic("unreachable")
})
panic("unreachable")
})

let__14480 := tmp17201
_ = let__14480

tmp17236 := GConsP(__e, let__14479)

if True == tmp17236 {
tmp17204 := GHd(__e, let__14479)

let__14481 := tmp17204
_ = let__14481

tmp17205 := GTl(__e, let__14479)

tmp17206 := Call2(__e, SymFn(symshen_4lazyderef), tmp17205, V5642)


let__14482 := tmp17206
_ = let__14482

tmp17207 := MakeNative1(func(__e *ControlFlow, Z5673 Obj) Obj {
tmp17208 := Call1(__e, let__14480, let__14481)


return TailCall1(__e, tmp17208, Z5673)


panic("unreachable")
})

let__14483 := tmp17207
_ = let__14483

tmp17224 := GConsP(__e, let__14482)

if True == tmp17224 {
tmp17209 := GHd(__e, let__14482)

let__14484 := tmp17209
_ = let__14484

tmp17210 := GTl(__e, let__14482)

tmp17211 := Call2(__e, SymFn(symshen_4lazyderef), tmp17210, V5642)


let__14485 := tmp17211
_ = let__14485

tmp17212 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__14483, let__14484)
panic("unreachable")
})

let__14486 := tmp17212
_ = let__14486

tmp17216 := GEq(__e, let__14485, Nil)

if True == tmp17216 {
return TailCall1(__e, SymFn(symthaw), let__14486)
} else {
tmp17214 := Call1(__e, SymFn(symshen_4pvar_2), let__14485)


if True == tmp17214 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__14485, Nil, V5642, let__14486)
} else {
return False
}


}


} else {
tmp17222 := Call1(__e, SymFn(symshen_4pvar_2), let__14482)


if True == tmp17222 {
tmp17217 := Call1(__e, SymFn(symshen_4newpv), V5642)


let__14487 := tmp17217
_ = let__14487

tmp17218 := GCons(__e, let__14487, Nil)

tmp17219 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__14483, let__14487)
panic("unreachable")
})

tmp17220 := Call4(__e, SymFn(symshen_4bind_b), let__14482, tmp17218, V5642, tmp17219)


return TailCall2(__e, SymFn(symshen_4gc), V5642, tmp17220)


} else {
return False
}


}


} else {
tmp17234 := Call1(__e, SymFn(symshen_4pvar_2), let__14479)


if True == tmp17234 {
tmp17225 := Call1(__e, SymFn(symshen_4newpv), V5642)


let__14488 := tmp17225
_ = let__14488

tmp17226 := Call1(__e, SymFn(symshen_4newpv), V5642)


let__14489 := tmp17226
_ = let__14489

tmp17227 := GCons(__e, let__14489, Nil)

tmp17228 := GCons(__e, let__14488, tmp17227)

tmp17229 := MakeNative0(func(__e *ControlFlow) Obj {
tmp17230 := Call1(__e, let__14480, let__14488)


return TailCall1(__e, tmp17230, let__14489)


panic("unreachable")
})

tmp17231 := Call4(__e, SymFn(symshen_4bind_b), let__14479, tmp17228, V5642, tmp17229)


tmp17232 := Call2(__e, SymFn(symshen_4gc), V5642, tmp17231)


return TailCall2(__e, SymFn(symshen_4gc), V5642, tmp17232)


} else {
return False
}


}


} else {
tmp17250 := Call1(__e, SymFn(symshen_4pvar_2), let__14475)


if True == tmp17250 {
tmp17237 := Call1(__e, SymFn(symshen_4newpv), V5642)


let__14490 := tmp17237
_ = let__14490

tmp17238 := Call1(__e, SymFn(symshen_4newpv), V5642)


let__14491 := tmp17238
_ = let__14491

tmp17239 := Call1(__e, SymFn(symshen_4newpv), V5642)


let__14492 := tmp17239
_ = let__14492

tmp17240 := GCons(__e, let__14492, Nil)

tmp17241 := GCons(__e, let__14491, tmp17240)

tmp17242 := GCons(__e, let__14490, tmp17241)

tmp17243 := MakeNative0(func(__e *ControlFlow) Obj {
tmp17244 := Call1(__e, let__14476, let__14490)


tmp17245 := Call1(__e, tmp17244, let__14491)


return TailCall1(__e, tmp17245, let__14492)


panic("unreachable")
})

tmp17246 := Call4(__e, SymFn(symshen_4bind_b), let__14475, tmp17242, V5642, tmp17243)


tmp17247 := Call2(__e, SymFn(symshen_4gc), V5642, tmp17246)


tmp17248 := Call2(__e, SymFn(symshen_4gc), V5642, tmp17247)


return TailCall2(__e, SymFn(symshen_4gc), V5642, tmp17248)


} else {
return False
}


}


} else {
tmp17270 := Call1(__e, SymFn(symshen_4pvar_2), let__14473)


if True == tmp17270 {
tmp17253 := Call1(__e, SymFn(symshen_4newpv), V5642)


let__14493 := tmp17253
_ = let__14493

tmp17254 := Call1(__e, SymFn(symshen_4newpv), V5642)


let__14494 := tmp17254
_ = let__14494

tmp17255 := Call1(__e, SymFn(symshen_4newpv), V5642)


let__14495 := tmp17255
_ = let__14495

tmp17256 := Call1(__e, SymFn(symshen_4newpv), V5642)


let__14496 := tmp17256
_ = let__14496

tmp17257 := GCons(__e, let__14495, Nil)

tmp17258 := GCons(__e, let__14494, tmp17257)

tmp17259 := GCons(__e, let__14493, tmp17258)

tmp17260 := GCons(__e, tmp17259, let__14496)

tmp17261 := MakeNative0(func(__e *ControlFlow) Obj {
tmp17262 := Call1(__e, let__14474, let__14493)


tmp17263 := Call1(__e, tmp17262, let__14494)


tmp17264 := Call1(__e, tmp17263, let__14495)


return TailCall1(__e, tmp17264, let__14496)


panic("unreachable")
})

tmp17265 := Call4(__e, SymFn(symshen_4bind_b), let__14473, tmp17260, V5642, tmp17261)


tmp17266 := Call2(__e, SymFn(symshen_4gc), V5642, tmp17265)


tmp17267 := Call2(__e, SymFn(symshen_4gc), V5642, tmp17266)


tmp17268 := Call2(__e, SymFn(symshen_4gc), V5642, tmp17267)


return TailCall2(__e, SymFn(symshen_4gc), V5642, tmp17268)


} else {
return False
}


}


} else {
return False
}


} else {
return False
}


} else {
return let__14466
}


panic("unreachable")
})

tmp17279 := Call2(__e, ns2_1set, symshen_4p_1hyps, tmp17167)


_ = tmp17279

tmp17280 := MakeNative(func(__e *ControlFlow) {
V5687 := __e.Get(1)
V5688 := __e.Get(2)
V5689 := __e.Get(3)
V5690 := __e.Get(4)
V5691 := __e.Get(5)
V5692 := __e.Get(6)
V5693 := __e.Get(7)
__e.ReturnDirect(func() Obj {
let__14497 := GAdd(__e, V5692, MakeInteger(1))
_ = let__14497

tmp17318 := Call1(__e, SymFn(symshen_4unlocked_2), V5691)


var ifres17282 Obj

if True == tmp17318 {
tmp17283 := Call2(__e, SymFn(symshen_4lazyderef), V5687, V5690)


let__14499 := tmp17283
_ = let__14499

tmp17317 := GConsP(__e, let__14499)

var ifres17284 Obj

if True == tmp17317 {
tmp17285 := GHd(__e, let__14499)

tmp17286 := Call2(__e, SymFn(symshen_4lazyderef), tmp17285, V5690)


let__14500 := tmp17286
_ = let__14500

tmp17316 := GEq(__e, let__14500, symwhere)

var ifres17287 Obj

if True == tmp17316 {
tmp17288 := GTl(__e, let__14499)

tmp17289 := Call2(__e, SymFn(symshen_4lazyderef), tmp17288, V5690)


let__14501 := tmp17289
_ = let__14501

tmp17315 := GConsP(__e, let__14501)

var ifres17290 Obj

if True == tmp17315 {
tmp17291 := GHd(__e, let__14501)

let__14502 := tmp17291
_ = let__14502

tmp17292 := GTl(__e, let__14501)

tmp17293 := Call2(__e, SymFn(symshen_4lazyderef), tmp17292, V5690)


let__14503 := tmp17293
_ = let__14503

tmp17314 := GConsP(__e, let__14503)

var ifres17294 Obj

if True == tmp17314 {
tmp17295 := GHd(__e, let__14503)

let__14504 := tmp17295
_ = let__14504

tmp17296 := GTl(__e, let__14503)

tmp17297 := Call2(__e, SymFn(symshen_4lazyderef), tmp17296, V5690)


let__14505 := tmp17297
_ = let__14505

tmp17313 := GEq(__e, let__14505, Nil)

var ifres17298 Obj

if True == tmp17313 {
tmp17299 := Call1(__e, SymFn(symshen_4newpv), V5690)


let__14506 := tmp17299
_ = let__14506

tmp17300 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp17300

tmp17301 := MakeNative0(func(__e *ControlFlow) Obj {
tmp17302 := Call1(__e, SymFn(symshen_4curry), let__14502)


tmp17303 := MakeNative0(func(__e *ControlFlow) Obj {
tmp17304 := MakeNative0(func(__e *ControlFlow) Obj {
tmp17305 := MakeNative0(func(__e *ControlFlow) Obj {
tmp17306 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp17307 := GCons(__e, symverified, Nil)

tmp17308 := GCons(__e, tmp17306, tmp17307)

tmp17309 := GCons(__e, let__14506, tmp17308)

tmp17310 := GCons(__e, tmp17309, V5689)

__e.TailApply(SymFn(symshen_4t_d_1correct), let__14504, V5688, tmp17310, V5690, V5691, let__14497, V5693)
return nil


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5690, V5691, let__14497, tmp17305)


panic("unreachable")
})

__e.TailApply(SymFn(symshen_4system_1S_1h), let__14506, symboolean, V5689, V5690, V5691, let__14497, tmp17304)
return nil


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), let__14506, tmp17302, V5690, V5691, let__14497, tmp17303)


panic("unreachable")
})

tmp17311 := Call4(__e, SymFn(symshen_4cut), V5690, V5691, let__14497, tmp17301)


tmp17312 := Call2(__e, SymFn(symshen_4gc), V5690, tmp17311)


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

tmp17327 := GEq(__e, let__14498, False)

if True == tmp17327 {
tmp17323 := Call1(__e, SymFn(symshen_4unlocked_2), V5691)


var ifres17319 Obj

if True == tmp17323 {
tmp17320 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp17320

tmp17321 := Call1(__e, SymFn(symshen_4curry), V5687)


tmp17322 := Call(__e, SymFn(symshen_4system_1S_1h), tmp17321, V5688, V5689, V5690, V5691, let__14497, V5693)


ifres17319 = tmp17322


} else {
ifres17319 = False


}

let__14507 := ifres17319
_ = let__14507

tmp17325 := GEq(__e, let__14507, False)

if True == tmp17325 {
return TailCall2(__e, SymFn(symshen_4unlock), V5691, let__14497)
} else {
return let__14507
}


} else {
return let__14498
}


panic("unreachable")
}())
}, 7)

tmp17328 := Call2(__e, ns2_1set, symshen_4t_d_1correct, tmp17280)


_ = tmp17328

tmp17329 := MakeNative(func(__e *ControlFlow) {
V5705 := __e.Get(1)
V5706 := __e.Get(2)
V5707 := __e.Get(3)
V5708 := __e.Get(4)
V5709 := __e.Get(5)
V5710 := __e.Get(6)
V5711 := __e.Get(7)
V5712 := __e.Get(8)
__e.ReturnDirect(func() Obj {
tmp17336 := Call1(__e, SymFn(symshen_4unlocked_2), V5710)


var ifres17330 Obj

if True == tmp17336 {
tmp17331 := Call2(__e, SymFn(symshen_4lazyderef), V5705, V5709)


let__14509 := tmp17331
_ = let__14509

tmp17335 := GEq(__e, let__14509, Nil)

var ifres17332 Obj

if True == tmp17335 {
tmp17333 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp17333

tmp17334 := Call6(__e, SymFn(symis_b), V5706, V5708, V5709, V5710, V5711, V5712)


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

tmp17368 := GEq(__e, let__14508, False)

if True == tmp17368 {
tmp17366 := Call1(__e, SymFn(symshen_4unlocked_2), V5710)


if True == tmp17366 {
tmp17337 := Call2(__e, SymFn(symshen_4lazyderef), V5705, V5709)


let__14510 := tmp17337
_ = let__14510

tmp17364 := GConsP(__e, let__14510)

if True == tmp17364 {
tmp17338 := GHd(__e, let__14510)

let__14511 := tmp17338
_ = let__14511

tmp17339 := GTl(__e, let__14510)

let__14512 := tmp17339
_ = let__14512

tmp17340 := Call2(__e, SymFn(symshen_4lazyderef), V5706, V5709)


let__14513 := tmp17340
_ = let__14513

tmp17362 := GConsP(__e, let__14513)

if True == tmp17362 {
tmp17341 := GHd(__e, let__14513)

let__14514 := tmp17341
_ = let__14514

tmp17342 := GTl(__e, let__14513)

tmp17343 := Call2(__e, SymFn(symshen_4lazyderef), tmp17342, V5709)


let__14515 := tmp17343
_ = let__14515

tmp17360 := GConsP(__e, let__14515)

if True == tmp17360 {
tmp17344 := GHd(__e, let__14515)

tmp17345 := Call2(__e, SymFn(symshen_4lazyderef), tmp17344, V5709)


let__14516 := tmp17345
_ = let__14516

tmp17358 := GEq(__e, let__14516, sym_1_1_6)

if True == tmp17358 {
tmp17346 := GTl(__e, let__14515)

tmp17347 := Call2(__e, SymFn(symshen_4lazyderef), tmp17346, V5709)


let__14517 := tmp17347
_ = let__14517

tmp17356 := GConsP(__e, let__14517)

if True == tmp17356 {
tmp17348 := GHd(__e, let__14517)

let__14518 := tmp17348
_ = let__14518

tmp17349 := GTl(__e, let__14517)

tmp17350 := Call2(__e, SymFn(symshen_4lazyderef), tmp17349, V5709)


let__14519 := tmp17350
_ = let__14519

tmp17354 := GEq(__e, let__14519, Nil)

if True == tmp17354 {
tmp17351 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp17351

tmp17352 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4t_d_1integrity), let__14512, let__14518, V5707, V5708, V5709, V5710, V5711, V5712)
return nil
panic("unreachable")
})

__e.TailApply(SymFn(symshen_4system_1S_1h), let__14511, let__14514, V5707, V5709, V5710, V5711, tmp17352)
return nil


} else {
return False
}


} else {
return False
}


} else {
return False
}


} else {
return False
}


} else {
return False
}


} else {
return False
}


} else {
return False
}


} else {
return let__14508
}


panic("unreachable")
}())
}, 8)

tmp17369 := Call2(__e, ns2_1set, symshen_4t_d_1integrity, tmp17329)


_ = tmp17369

tmp17370 := MakeNative1(func(__e *ControlFlow, V5725 Obj) Obj {
tmp17379 := GVectorP(__e, V5725)

if True == tmp17379 {
tmp17376 := PrimGuard1(__e, symstring_2, PrimIsString, V5725)

tmp17377 := GNot(__e, tmp17376)

var ifres17372 Obj

if True == tmp17377 {
tmp17374 := GVecGet(__e, V5725, MakeInteger(0))

tmp17375 := GEq(__e, tmp17374, symshen_4print_1freshterm)

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
return True
} else {
return False
}


} else {
return False
}


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4freshterm_2, tmp17370)




panic("unreachable")
})

