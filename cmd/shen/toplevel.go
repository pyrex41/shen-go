package main

import . "github.com/pyrex41/shen-go/kl"

var TopLevelMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp7745 := MakeNative0(func(__e *ControlFlow) Obj {
tmp7746 := Call0(__e, SymFn(symshen_4credits))


_ = tmp7746

return TailCall0(__e, SymFn(symshen_4loop))


panic("unreachable")
})

tmp7747 := Call2(__e, ns2_1set, symshen_4repl, tmp7745)


_ = tmp7747

tmp7748 := (func() Obj {
var __self Obj
__self = MakeNative0(func(__e *ControlFlow) Obj {
__selftop:
tmp7749 := Call0(__e, SymFn(symshen_4initialise__environment))


_ = tmp7749

tmp7750 := Call0(__e, SymFn(symshen_4prompt))


_ = tmp7750

tmp7751 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall0(__e, SymFn(symshen_4read_1evaluate_1print))
panic("unreachable")
})

tmp7752 := MakeNative1(func(__e *ControlFlow, Z5748 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4toplevel_1display_1exception), Z5748)
panic("unreachable")
})

tmp7753 := Call2(__e, try_1catch, tmp7751, tmp7752)


_ = tmp7753

if SymFn(symshen_4loop) == __self {
__e.Tick()
goto __selftop
}
return TailCall0(__e, SymFn(symshen_4loop))


panic("unreachable")
})
return __self
})()

tmp7754 := Call2(__e, ns2_1set, symshen_4loop, tmp7748)


_ = tmp7754

tmp7755 := MakeNative1(func(__e *ControlFlow, V5749 Obj) Obj {
tmp7756 := PrimGuard1(__e, symerror_1to_1string, PrimErrorToString, V5749)

tmp7757 := Call0(__e, SymFn(symstoutput))


tmp7758 := Call2(__e, SymFn(sympr), tmp7756, tmp7757)


_ = tmp7758

return TailCall1(__e, SymFn(symnl), MakeInteger(0))


panic("unreachable")
})

tmp7759 := Call2(__e, ns2_1set, symshen_4toplevel_1display_1exception, tmp7755)


_ = tmp7759

tmp7760 := MakeNative0(func(__e *ControlFlow) Obj {
tmp7761 := Call0(__e, SymFn(symstoutput))


tmp7762 := Call2(__e, SymFn(sympr), MakeString("\nShen, www.shenlanguage.org, copyright (C) 2010-2026, Mark Tarver\n"), tmp7761)


_ = tmp7762

tmp7763 := GValue(__e, sym_dversion_d)

tmp7764 := GValue(__e, sym_dlanguage_d)

tmp7765 := GValue(__e, sym_dimplementation_d)

tmp7766 := GValue(__e, sym_drelease_d)

tmp7767 := Call3(__e, SymFn(symshen_4app), tmp7766, MakeString("\n"), symshen_4a)


tmp7769 := Call3(__e, SymFn(symshen_4app), tmp7765, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" "), tmp7767), symshen_4a)


tmp7771 := Call3(__e, SymFn(symshen_4app), tmp7764, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(", platform: "), tmp7769), symshen_4a)


tmp7773 := Call3(__e, SymFn(symshen_4app), tmp7763, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(", language: "), tmp7771), symshen_4a)


tmp7774 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("version: S"), tmp7773)

tmp7775 := Call0(__e, SymFn(symstoutput))


tmp7776 := Call2(__e, SymFn(sympr), tmp7774, tmp7775)


_ = tmp7776

tmp7777 := GValue(__e, sym_dport_d)

tmp7778 := GValue(__e, sym_dporters_d)

tmp7779 := Call3(__e, SymFn(symshen_4app), tmp7778, MakeString("\n\n"), symshen_4a)


tmp7781 := Call3(__e, SymFn(symshen_4app), tmp7777, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(", ported by "), tmp7779), symshen_4a)


tmp7782 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("port "), tmp7781)

tmp7783 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp7782, tmp7783)


panic("unreachable")
})

tmp7784 := Call2(__e, ns2_1set, symshen_4credits, tmp7760)


_ = tmp7784

tmp7785 := MakeNative0(func(__e *ControlFlow) Obj {
tmp7786 := PrimGuard2(__e, symset, PrimSet, symshen_4_dcall_d, MakeInteger(0))

_ = tmp7786

return PrimGuard2(__e, symset, PrimSet, symshen_4_dinfs_d, MakeInteger(0))


panic("unreachable")
})

tmp7787 := Call2(__e, ns2_1set, symshen_4initialise__environment, tmp7785)


_ = tmp7787

tmp7788 := MakeNative0(func(__e *ControlFlow) Obj {
tmp7800 := GValue(__e, symshen_4_dtc_d)

if True == tmp7800 {
tmp7789 := GValue(__e, symshen_4_dhistory_d)

tmp7790 := Call1(__e, SymFn(symlength), tmp7789)


tmp7791 := Call3(__e, SymFn(symshen_4app), tmp7790, MakeString("+) "), symshen_4a)


tmp7792 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("\n("), tmp7791)

tmp7793 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp7792, tmp7793)


} else {
tmp7794 := GValue(__e, symshen_4_dhistory_d)

tmp7795 := Call1(__e, SymFn(symlength), tmp7794)


tmp7796 := Call3(__e, SymFn(symshen_4app), tmp7795, MakeString("-) "), symshen_4a)


tmp7797 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("\n("), tmp7796)

tmp7798 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp7797, tmp7798)


}


panic("unreachable")
})

tmp7801 := Call2(__e, ns2_1set, symshen_4prompt, tmp7788)


_ = tmp7801

tmp7802 := MakeNative0(func(__e *ControlFlow) Obj {
tmp7803 := GValue(__e, symshen_4_dpackage_d)

let__7729 := tmp7803
_ = let__7729

tmp7804 := Call0(__e, SymFn(symstinput))


tmp7805 := Call1(__e, SymFn(symlineread), tmp7804)


tmp7806 := Call2(__e, SymFn(symshen_4package_1user_1input), let__7729, tmp7805)


let__7730 := tmp7806
_ = let__7730

tmp7807 := Call0(__e, SymFn(symshen_4update_1history))


let__7731 := tmp7807
_ = let__7731

tmp7808 := GValue(__e, symshen_4_dtc_d)

return TailCall3(__e, SymFn(symshen_4evaluate_1lineread), let__7730, let__7731, tmp7808)


panic("unreachable")
})

tmp7809 := Call2(__e, ns2_1set, symshen_4read_1evaluate_1print, tmp7802)


_ = tmp7809

tmp7810 := MakeNative2(func(__e *ControlFlow, V5753, V5754 Obj) Obj {
tmp7815 := GEq(__e, symnull, V5753)

if True == tmp7815 {
return V5754
} else {
tmp7811 := PrimGuard1(__e, symstr, PrimStr, V5753)

let__7732 := tmp7811
_ = let__7732

tmp7812 := Call1(__e, SymFn(symexternal), V5753)


let__7733 := tmp7812
_ = let__7733

tmp7813 := MakeNative1(func(__e *ControlFlow, Z5757 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4pui_1h), let__7732, let__7733, Z5757)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp7813, V5754)


}


panic("unreachable")
})

tmp7816 := Call2(__e, ns2_1set, symshen_4package_1user_1input, tmp7810)


_ = tmp7816

tmp7817 := MakeNative3(func(__e *ControlFlow, V5762, V5763, V5764 Obj) Obj {
tmp7855 := GConsP(__e, V5764)

var ifres7845 Obj

if True == tmp7855 {
tmp7853 := GHd(__e, V5764)

tmp7854 := GEq(__e, symfn, tmp7853)

var ifres7846 Obj

if True == tmp7854 {
tmp7851 := GTl(__e, V5764)

tmp7852 := GConsP(__e, tmp7851)

var ifres7847 Obj

if True == tmp7852 {
tmp7848 := GTl(__e, V5764)

tmp7849 := GTl(__e, tmp7848)

tmp7850 := GEq(__e, Nil, tmp7849)

ifres7847 = tmp7850


} else {
ifres7847 = False


}

ifres7846 = ifres7847


} else {
ifres7846 = False


}

ifres7845 = ifres7846


} else {
ifres7845 = False


}

if True == ifres7845 {
tmp7823 := GTl(__e, V5764)

tmp7824 := GHd(__e, tmp7823)

tmp7825 := Call3(__e, SymFn(symshen_4internal_2), tmp7824, V5762, V5763)


if True == tmp7825 {
tmp7818 := GTl(__e, V5764)

tmp7819 := GHd(__e, tmp7818)

tmp7820 := Call2(__e, SymFn(symshen_4intern_1in_1package), V5762, tmp7819)


tmp7821 := GCons(__e, tmp7820, Nil)

return GCons(__e, symfn, tmp7821)


} else {
return V5764
}


} else {
tmp7843 := GConsP(__e, V5764)

if True == tmp7843 {
tmp7840 := GHd(__e, V5764)

tmp7841 := Call3(__e, SymFn(symshen_4internal_2), tmp7840, V5762, V5763)


if True == tmp7841 {
tmp7826 := GHd(__e, V5764)

tmp7827 := Call2(__e, SymFn(symshen_4intern_1in_1package), V5762, tmp7826)


tmp7828 := MakeNative1(func(__e *ControlFlow, Z5765 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4pui_1h), V5762, V5763, Z5765)
panic("unreachable")
})

tmp7829 := GTl(__e, V5764)

tmp7830 := Call2(__e, SymFn(symmap), tmp7828, tmp7829)


return GCons(__e, tmp7827, tmp7830)


} else {
tmp7837 := GHd(__e, V5764)

tmp7838 := GConsP(__e, tmp7837)

if True == tmp7838 {
tmp7831 := MakeNative1(func(__e *ControlFlow, Z5766 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4pui_1h), V5762, V5763, Z5766)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp7831, V5764)


} else {
tmp7832 := GHd(__e, V5764)

tmp7833 := MakeNative1(func(__e *ControlFlow, Z5767 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4pui_1h), V5762, V5763, Z5767)
panic("unreachable")
})

tmp7834 := GTl(__e, V5764)

tmp7835 := Call2(__e, SymFn(symmap), tmp7833, tmp7834)


return GCons(__e, tmp7832, tmp7835)


}


}


} else {
return V5764
}


}


panic("unreachable")
})

tmp7856 := Call2(__e, ns2_1set, symshen_4pui_1h, tmp7817)


_ = tmp7856

tmp7857 := MakeNative0(func(__e *ControlFlow) Obj {
tmp7858 := Call0(__e, SymFn(symit))


tmp7859 := Call1(__e, SymFn(symshen_4trim_1it), tmp7858)


tmp7860 := GValue(__e, symshen_4_dhistory_d)

tmp7861 := GCons(__e, tmp7859, tmp7860)

return PrimGuard2(__e, symset, PrimSet, symshen_4_dhistory_d, tmp7861)


panic("unreachable")
})

tmp7862 := Call2(__e, ns2_1set, symshen_4update_1history, tmp7857)


_ = tmp7862

tmp7863 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V5768 := __self1
_ = V5768
tmp7870 := Call1(__e, SymFn(symshen_4_7string_2), V5768)


var ifres7866 Obj

if True == tmp7870 {
tmp7867 := Call1(__e, SymFn(symhdstr), V5768)


tmp7868 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp7867)

tmp7869 := Call1(__e, SymFn(symshen_4whitespace_2), tmp7868)


ifres7866 = tmp7869


} else {
ifres7866 = False


}

if True == ifres7866 {
if SymFn(symshen_4trim_1it) == __self {
__self1 = PrimGuard1(__e, symtlstr, PrimTailString, V5768)
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4trim_1it), PrimGuard1(__e, symtlstr, PrimTailString, V5768))


} else {
return V5768
}


panic("unreachable")
})
return __self
})()

tmp7871 := Call2(__e, ns2_1set, symshen_4trim_1it, tmp7863)


_ = tmp7871

tmp7872 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V5787 := __self1
_ = V5787
V5788 := __self2
_ = V5788
V5789 := __self3
_ = V5789
tmp7976 := GConsP(__e, V5787)

var ifres7952 Obj

if True == tmp7976 {
tmp7974 := GTl(__e, V5787)

tmp7975 := GEq(__e, Nil, tmp7974)

var ifres7953 Obj

if True == tmp7975 {
tmp7973 := GConsP(__e, V5788)

var ifres7954 Obj

if True == tmp7973 {
tmp7971 := GHd(__e, V5788)

tmp7972 := Call1(__e, SymFn(symshen_4_7string_2), tmp7971)


var ifres7955 Obj

if True == tmp7972 {
tmp7968 := GHd(__e, V5788)

tmp7969 := Call1(__e, SymFn(symhdstr), tmp7968)


tmp7970 := GEq(__e, MakeString("!"), tmp7969)

var ifres7956 Obj

if True == tmp7970 {
tmp7965 := GHd(__e, V5788)

tmp7967 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, tmp7965))


var ifres7957 Obj

if True == tmp7967 {
tmp7961 := GHd(__e, V5788)

tmp7963 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, tmp7961))


tmp7964 := GEq(__e, MakeString("!"), tmp7963)

var ifres7958 Obj

if True == tmp7964 {
tmp7959 := GTl(__e, V5788)

tmp7960 := GConsP(__e, tmp7959)

ifres7958 = tmp7960


} else {
ifres7958 = False


}

ifres7957 = ifres7958


} else {
ifres7957 = False


}

ifres7956 = ifres7957


} else {
ifres7956 = False


}

ifres7955 = ifres7956


} else {
ifres7955 = False


}

ifres7954 = ifres7955


} else {
ifres7954 = False


}

ifres7953 = ifres7954


} else {
ifres7953 = False


}

ifres7952 = ifres7953


} else {
ifres7952 = False


}

if True == ifres7952 {
tmp7873 := GTl(__e, V5788)

tmp7874 := GHd(__e, tmp7873)

tmp7875 := Call1(__e, SymFn(symread_1from_1string), tmp7874)


let__7734 := tmp7875
_ = let__7734

tmp7876 := GTl(__e, V5788)

tmp7877 := GHd(__e, tmp7876)

tmp7878 := GTl(__e, V5788)

tmp7879 := GCons(__e, tmp7877, tmp7878)

tmp7880 := PrimGuard2(__e, symset, PrimSet, symshen_4_dhistory_d, tmp7879)

let__7735 := tmp7880
_ = let__7735

tmp7881 := GTl(__e, V5788)

tmp7882 := GHd(__e, tmp7881)

tmp7883 := Call3(__e, SymFn(symshen_4app), tmp7882, MakeString("\n"), symshen_4a)


tmp7884 := Call0(__e, SymFn(symstoutput))


tmp7885 := Call2(__e, SymFn(sympr), tmp7883, tmp7884)


let__7736 := tmp7885
_ = let__7736

if SymFn(symshen_4evaluate_1lineread) == __self {
__self1, __self2, __self3 = let__7734, let__7735, V5789
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4evaluate_1lineread), let__7734, let__7735, V5789)


} else {
tmp7950 := GConsP(__e, V5787)

var ifres7938 Obj

if True == tmp7950 {
tmp7948 := GTl(__e, V5787)

tmp7949 := GEq(__e, Nil, tmp7948)

var ifres7939 Obj

if True == tmp7949 {
tmp7947 := GConsP(__e, V5788)

var ifres7940 Obj

if True == tmp7947 {
tmp7945 := GHd(__e, V5788)

tmp7946 := Call1(__e, SymFn(symshen_4_7string_2), tmp7945)


var ifres7941 Obj

if True == tmp7946 {
tmp7942 := GHd(__e, V5788)

tmp7943 := Call1(__e, SymFn(symhdstr), tmp7942)


tmp7944 := GEq(__e, MakeString("!"), tmp7943)

ifres7941 = tmp7944


} else {
ifres7941 = False


}

ifres7940 = ifres7941


} else {
ifres7940 = False


}

ifres7939 = ifres7940


} else {
ifres7939 = False


}

ifres7938 = ifres7939


} else {
ifres7938 = False


}

if True == ifres7938 {
tmp7891 := GHd(__e, V5788)

tmp7893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(PrimGuard1(__e, symtlstr, PrimTailString, tmp7891), MakeString(""))
}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, tmp7891)
__typedArg1 := MakeString("")
return Call2(__e, SymFn(sym_a), __typedArg0, __typedArg1)
})()

var ifres7886 Obj

if True == tmp7893 {
ifres7886 = Nil


} else {
tmp7887 := GHd(__e, V5788)

tmp7889 := Call1(__e, SymFn(symread_1from_1string), PrimGuard1(__e, symtlstr, PrimTailString, tmp7887))


tmp7890 := GHd(__e, tmp7889)

ifres7886 = tmp7890


}

let__7737 := ifres7886
_ = let__7737

tmp7894 := GHd(__e, V5788)

tmp7895 := PrimGuard1(__e, symtlstr, PrimTailString, tmp7894)

tmp7896 := GTl(__e, V5788)

tmp7897 := Call3(__e, SymFn(symshen_4use_1history), let__7737, tmp7895, tmp7896)


let__7738 := tmp7897
_ = let__7738

tmp7898 := Call3(__e, SymFn(symshen_4app), let__7738, MakeString("\n"), symshen_4a)


tmp7899 := Call0(__e, SymFn(symstoutput))


tmp7900 := Call2(__e, SymFn(sympr), tmp7898, tmp7899)


let__7739 := tmp7900
_ = let__7739

tmp7901 := Call1(__e, SymFn(symread_1from_1string), let__7738)


let__7740 := tmp7901
_ = let__7740

tmp7902 := GTl(__e, V5788)

tmp7903 := GCons(__e, let__7738, tmp7902)

tmp7904 := PrimGuard2(__e, symset, PrimSet, symshen_4_dhistory_d, tmp7903)

let__7741 := tmp7904
_ = let__7741

if SymFn(symshen_4evaluate_1lineread) == __self {
__self1, __self2, __self3 = let__7740, let__7741, V5789
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4evaluate_1lineread), let__7740, let__7741, V5789)


} else {
tmp7936 := GConsP(__e, V5787)

var ifres7924 Obj

if True == tmp7936 {
tmp7934 := GTl(__e, V5787)

tmp7935 := GEq(__e, Nil, tmp7934)

var ifres7925 Obj

if True == tmp7935 {
tmp7933 := GConsP(__e, V5788)

var ifres7926 Obj

if True == tmp7933 {
tmp7931 := GHd(__e, V5788)

tmp7932 := Call1(__e, SymFn(symshen_4_7string_2), tmp7931)


var ifres7927 Obj

if True == tmp7932 {
tmp7928 := GHd(__e, V5788)

tmp7929 := Call1(__e, SymFn(symhdstr), tmp7928)


tmp7930 := GEq(__e, MakeString("%"), tmp7929)

ifres7927 = tmp7930


} else {
ifres7927 = False


}

ifres7926 = ifres7927


} else {
ifres7926 = False


}

ifres7925 = ifres7926


} else {
ifres7925 = False


}

ifres7924 = ifres7925


} else {
ifres7924 = False


}

if True == ifres7924 {
tmp7910 := GHd(__e, V5788)

tmp7912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(PrimGuard1(__e, symtlstr, PrimTailString, tmp7910), MakeString(""))
}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, tmp7910)
__typedArg1 := MakeString("")
return Call2(__e, SymFn(sym_a), __typedArg0, __typedArg1)
})()

var ifres7905 Obj

if True == tmp7912 {
ifres7905 = Nil


} else {
tmp7906 := GHd(__e, V5788)

tmp7908 := Call1(__e, SymFn(symread_1from_1string), PrimGuard1(__e, symtlstr, PrimTailString, tmp7906))


tmp7909 := GHd(__e, tmp7908)

ifres7905 = tmp7909


}

let__7742 := ifres7905
_ = let__7742

tmp7913 := GHd(__e, V5788)

tmp7914 := PrimGuard1(__e, symtlstr, PrimTailString, tmp7913)

tmp7915 := GTl(__e, V5788)

tmp7916 := Call3(__e, SymFn(symshen_4peek_1history), let__7742, tmp7914, tmp7915)


let__7743 := tmp7916
_ = let__7743

tmp7917 := GTl(__e, V5788)

tmp7918 := PrimGuard2(__e, symset, PrimSet, symshen_4_dhistory_d, tmp7917)

let__7744 := tmp7918
_ = let__7744

return TailCall0(__e, SymFn(symabort))


} else {
tmp7922 := GEq(__e, True, V5789)

if True == tmp7922 {
return TailCall1(__e, SymFn(symshen_4check_1eval_1and_1print), V5787)
} else {
tmp7920 := GEq(__e, False, V5789)

if True == tmp7920 {
return TailCall1(__e, SymFn(symshen_4eval_1and_1print), V5787)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.evaluate-lineread"))
}


}


}


}


}


panic("unreachable")
})
return __self
})()

tmp7977 := Call2(__e, ns2_1set, symshen_4evaluate_1lineread, tmp7872)


_ = tmp7977

tmp7978 := MakeNative3(func(__e *ControlFlow, V5801, V5802, V5803 Obj) Obj {
tmp7984 := PrimGuard1(__e, syminteger_2, PrimIsInteger, V5801)

if True == tmp7984 {
tmp7979 := GAdd(__e, MakeInteger(1), V5801)

tmp7980 := Call1(__e, SymFn(symreverse), V5803)


return TailCall2(__e, SymFn(symnth), tmp7979, tmp7980)


} else {
tmp7982 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V5801)

if True == tmp7982 {
return TailCall2(__e, SymFn(symshen_4string_1match), V5802, V5803)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("! expects a number or a symbol\n"))
}


}


panic("unreachable")
})

tmp7985 := Call2(__e, ns2_1set, symshen_4use_1history, tmp7978)


_ = tmp7985

tmp7986 := MakeNative3(func(__e *ControlFlow, V5804, V5805, V5806 Obj) Obj {
tmp7999 := PrimGuard1(__e, syminteger_2, PrimIsInteger, V5804)

if True == tmp7999 {
tmp7987 := GAdd(__e, MakeInteger(1), V5804)

tmp7988 := Call1(__e, SymFn(symreverse), V5806)


tmp7989 := Call2(__e, SymFn(symnth), tmp7987, tmp7988)


tmp7990 := Call3(__e, SymFn(symshen_4app), tmp7989, MakeString(""), symshen_4a)


tmp7991 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("\n"), tmp7990)

tmp7992 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp7991, tmp7992)


} else {
tmp7997 := GEq(__e, V5805, MakeString(""))

var ifres7995 Obj

if True == tmp7997 {
ifres7995 = True


} else {
tmp7996 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V5804)

ifres7995 = tmp7996


}

if True == ifres7995 {
tmp7993 := Call1(__e, SymFn(symreverse), V5806)


return TailCall3(__e, SymFn(symshen_4recursive_1string_1match), MakeInteger(0), V5805, tmp7993)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("% expects a number or a symbol\n"))
}


}


panic("unreachable")
})

tmp8000 := Call2(__e, ns2_1set, symshen_4peek_1history, tmp7986)


_ = tmp8000

tmp8001 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V5816 := __self1
_ = V5816
V5817 := __self2
_ = V5817
tmp8011 := GEq(__e, Nil, V5817)

if True == tmp8011 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("\ninput not found"))
} else {
tmp8009 := GConsP(__e, V5817)

var ifres8006 Obj

if True == tmp8009 {
tmp8007 := GHd(__e, V5817)

tmp8008 := Call2(__e, SymFn(symshen_4string_1prefix_2), V5816, tmp8007)


ifres8006 = tmp8008


} else {
ifres8006 = False


}

if True == ifres8006 {
return GHd(__e, V5817)
} else {
tmp8004 := GConsP(__e, V5817)

if True == tmp8004 {
tmp8002 := GTl(__e, V5817)

if SymFn(symshen_4string_1match) == __self {
__self1, __self2 = V5816, tmp8002
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4string_1match), V5816, tmp8002)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.string-match"))
}


}


}


panic("unreachable")
})
return __self
})()

tmp8012 := Call2(__e, ns2_1set, symshen_4string_1match, tmp8001)


_ = tmp8012

tmp8013 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V5825 := __self1
_ = V5825
V5826 := __self2
_ = V5826
tmp8045 := GEq(__e, MakeString(""), V5825)

if True == tmp8045 {
return True
} else {
tmp8043 := Call1(__e, SymFn(symshen_4_7string_2), V5825)


var ifres8039 Obj

if True == tmp8043 {
tmp8040 := Call1(__e, SymFn(symhdstr), V5825)


tmp8041 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp8040)

tmp8042 := Call1(__e, SymFn(symshen_4whitespace_2), tmp8041)


ifres8039 = tmp8042


} else {
ifres8039 = False


}

if True == ifres8039 {
if SymFn(symshen_4string_1prefix_2) == __self {
__self1, __self2 = PrimGuard1(__e, symtlstr, PrimTailString, V5825), V5826
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4string_1prefix_2), PrimGuard1(__e, symtlstr, PrimTailString, V5825), V5826)


} else {
tmp8037 := Call1(__e, SymFn(symshen_4_7string_2), V5826)


var ifres8033 Obj

if True == tmp8037 {
tmp8034 := Call1(__e, SymFn(symhdstr), V5826)


tmp8035 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp8034)

tmp8036 := Call1(__e, SymFn(symshen_4whitespace_2), tmp8035)


ifres8033 = tmp8036


} else {
ifres8033 = False


}

if True == ifres8033 {
if SymFn(symshen_4string_1prefix_2) == __self {
__self1, __self2 = V5825, PrimGuard1(__e, symtlstr, PrimTailString, V5826)
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4string_1prefix_2), V5825, PrimGuard1(__e, symtlstr, PrimTailString, V5826))


} else {
tmp8031 := Call1(__e, SymFn(symshen_4_7string_2), V5826)


var ifres8028 Obj

if True == tmp8031 {
tmp8029 := Call1(__e, SymFn(symhdstr), V5826)


tmp8030 := GEq(__e, MakeString("("), tmp8029)

ifres8028 = tmp8030


} else {
ifres8028 = False


}

if True == ifres8028 {
if SymFn(symshen_4string_1prefix_2) == __self {
__self1, __self2 = V5825, PrimGuard1(__e, symtlstr, PrimTailString, V5826)
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4string_1prefix_2), V5825, PrimGuard1(__e, symtlstr, PrimTailString, V5826))


} else {
tmp8026 := Call1(__e, SymFn(symshen_4_7string_2), V5825)


var ifres8020 Obj

if True == tmp8026 {
tmp8025 := Call1(__e, SymFn(symshen_4_7string_2), V5826)


var ifres8021 Obj

if True == tmp8025 {
tmp8022 := Call1(__e, SymFn(symhdstr), V5825)


tmp8023 := Call1(__e, SymFn(symhdstr), V5826)


tmp8024 := GEq(__e, tmp8022, tmp8023)

ifres8021 = tmp8024


} else {
ifres8021 = False


}

ifres8020 = ifres8021


} else {
ifres8020 = False


}

if True == ifres8020 {
tmp8017 := PrimGuard1(__e, symtlstr, PrimTailString, V5825)

if SymFn(symshen_4string_1prefix_2) == __self {
__self1, __self2 = tmp8017, PrimGuard1(__e, symtlstr, PrimTailString, V5826)
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4string_1prefix_2), tmp8017, PrimGuard1(__e, symtlstr, PrimTailString, V5826))


} else {
return False
}


}


}


}


}


panic("unreachable")
})
return __self
})()

tmp8046 := Call2(__e, ns2_1set, symshen_4string_1prefix_2, tmp8013)


_ = tmp8046

tmp8047 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V5837 := __self1
_ = V5837
V5838 := __self2
_ = V5838
V5839 := __self3
_ = V5839
tmp8062 := GEq(__e, Nil, V5839)

if True == tmp8062 {
return symshen_4skip
} else {
tmp8060 := GConsP(__e, V5839)

if True == tmp8060 {
tmp8055 := GHd(__e, V5839)

tmp8056 := Call2(__e, SymFn(symshen_4string_1prefix_2), V5838, tmp8055)


var ifres8048 Obj

if True == tmp8056 {
tmp8049 := GHd(__e, V5839)

tmp8050 := Call3(__e, SymFn(symshen_4app), tmp8049, MakeString("\n"), symshen_4a)


tmp8052 := Call3(__e, SymFn(symshen_4app), V5837, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(". "), tmp8050), symshen_4a)


tmp8053 := Call0(__e, SymFn(symstoutput))


tmp8054 := Call2(__e, SymFn(sympr), tmp8052, tmp8053)


ifres8048 = tmp8054


} else {
ifres8048 = symshen_4skip


}

_ = ifres8048

tmp8057 := GAdd(__e, V5837, MakeInteger(1))

tmp8058 := GTl(__e, V5839)

if SymFn(symshen_4recursive_1string_1match) == __self {
__self1, __self2, __self3 = tmp8057, V5838, tmp8058
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4recursive_1string_1match), tmp8057, V5838, tmp8058)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.recursive-string-match"))
}


}


panic("unreachable")
})
return __self
})()

return TailCall2(__e, ns2_1set, symshen_4recursive_1string_1match, tmp8047)




panic("unreachable")
})

