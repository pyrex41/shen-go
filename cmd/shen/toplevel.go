package main

import . "github.com/pyrex41/shen-go/kl"

var TopLevelMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp7765 := MakeNative0(func(__e *ControlFlow) Obj {
tmp7766 := Call0(__e, SymFn(symshen_4credits))


_ = tmp7766

return TailCall0(__e, SymFn(symshen_4loop))


panic("unreachable")
})

tmp7767 := Call2(__e, ns2_1set, symshen_4repl, tmp7765)


_ = tmp7767

tmp7768 := (func() Obj {
var __self Obj
__self = MakeNative0(func(__e *ControlFlow) Obj {
__selftop:
tmp7769 := Call0(__e, SymFn(symshen_4initialise__environment))


_ = tmp7769

tmp7770 := Call0(__e, SymFn(symshen_4prompt))


_ = tmp7770

tmp7771 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall0(__e, SymFn(symshen_4read_1evaluate_1print))
panic("unreachable")
})

tmp7772 := MakeNative1(func(__e *ControlFlow, Z5763 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4toplevel_1display_1exception), Z5763)
panic("unreachable")
})

tmp7773 := Call2(__e, try_1catch, tmp7771, tmp7772)


_ = tmp7773

if SymFn(symshen_4loop) == __self {
__e.Tick()
goto __selftop
}
return TailCall0(__e, SymFn(symshen_4loop))


panic("unreachable")
})
return __self
})()

tmp7774 := Call2(__e, ns2_1set, symshen_4loop, tmp7768)


_ = tmp7774

tmp7775 := MakeNative1(func(__e *ControlFlow, V5764 Obj) Obj {
tmp7776 := PrimGuard1(__e, symerror_1to_1string, PrimErrorToString, V5764)

tmp7777 := Call0(__e, SymFn(symstoutput))


tmp7778 := Call2(__e, SymFn(sympr), tmp7776, tmp7777)


_ = tmp7778

return TailCall1(__e, SymFn(symnl), MakeInteger(0))


panic("unreachable")
})

tmp7779 := Call2(__e, ns2_1set, symshen_4toplevel_1display_1exception, tmp7775)


_ = tmp7779

tmp7780 := MakeNative0(func(__e *ControlFlow) Obj {
tmp7781 := Call0(__e, SymFn(symstoutput))


tmp7782 := Call2(__e, SymFn(sympr), MakeString("\nShen, www.shenlanguage.org, copyright (C) 2010-2026, Mark Tarver\n"), tmp7781)


_ = tmp7782

tmp7783 := GValue(__e, sym_dversion_d)

tmp7784 := GValue(__e, sym_dlanguage_d)

tmp7785 := GValue(__e, sym_dimplementation_d)

tmp7786 := GValue(__e, sym_drelease_d)

tmp7787 := Call3(__e, SymFn(symshen_4app), tmp7786, MakeString("\n"), symshen_4a)


tmp7789 := Call3(__e, SymFn(symshen_4app), tmp7785, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" "), tmp7787), symshen_4a)


tmp7791 := Call3(__e, SymFn(symshen_4app), tmp7784, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(", platform: "), tmp7789), symshen_4a)


tmp7793 := Call3(__e, SymFn(symshen_4app), tmp7783, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(", language: "), tmp7791), symshen_4a)


tmp7794 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("version: S"), tmp7793)

tmp7795 := Call0(__e, SymFn(symstoutput))


tmp7796 := Call2(__e, SymFn(sympr), tmp7794, tmp7795)


_ = tmp7796

tmp7797 := GValue(__e, sym_dport_d)

tmp7798 := GValue(__e, sym_dporters_d)

tmp7799 := Call3(__e, SymFn(symshen_4app), tmp7798, MakeString("\n\n"), symshen_4a)


tmp7801 := Call3(__e, SymFn(symshen_4app), tmp7797, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(", ported by "), tmp7799), symshen_4a)


tmp7802 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("port "), tmp7801)

tmp7803 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp7802, tmp7803)


panic("unreachable")
})

tmp7804 := Call2(__e, ns2_1set, symshen_4credits, tmp7780)


_ = tmp7804

tmp7805 := MakeNative0(func(__e *ControlFlow) Obj {
tmp7806 := PrimGuard2(__e, symset, PrimSet, symshen_4_dcall_d, MakeInteger(0))

_ = tmp7806

return PrimGuard2(__e, symset, PrimSet, symshen_4_dinfs_d, MakeInteger(0))


panic("unreachable")
})

tmp7807 := Call2(__e, ns2_1set, symshen_4initialise__environment, tmp7805)


_ = tmp7807

tmp7808 := MakeNative0(func(__e *ControlFlow) Obj {
tmp7820 := GValue(__e, symshen_4_dtc_d)

if True == tmp7820 {
tmp7809 := GValue(__e, symshen_4_dhistory_d)

tmp7810 := Call1(__e, SymFn(symlength), tmp7809)


tmp7811 := Call3(__e, SymFn(symshen_4app), tmp7810, MakeString("+) "), symshen_4a)


tmp7812 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("\n("), tmp7811)

tmp7813 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp7812, tmp7813)


} else {
tmp7814 := GValue(__e, symshen_4_dhistory_d)

tmp7815 := Call1(__e, SymFn(symlength), tmp7814)


tmp7816 := Call3(__e, SymFn(symshen_4app), tmp7815, MakeString("-) "), symshen_4a)


tmp7817 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("\n("), tmp7816)

tmp7818 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp7817, tmp7818)


}


panic("unreachable")
})

tmp7821 := Call2(__e, ns2_1set, symshen_4prompt, tmp7808)


_ = tmp7821

tmp7822 := MakeNative0(func(__e *ControlFlow) Obj {
tmp7823 := GValue(__e, symshen_4_dpackage_d)

let__7749 := tmp7823
_ = let__7749

tmp7824 := Call0(__e, SymFn(symstinput))


tmp7825 := Call1(__e, SymFn(symlineread), tmp7824)


tmp7826 := Call2(__e, SymFn(symshen_4package_1user_1input), let__7749, tmp7825)


let__7750 := tmp7826
_ = let__7750

tmp7827 := Call0(__e, SymFn(symshen_4update_1history))


let__7751 := tmp7827
_ = let__7751

tmp7828 := GValue(__e, symshen_4_dtc_d)

return TailCall3(__e, SymFn(symshen_4evaluate_1lineread), let__7750, let__7751, tmp7828)


panic("unreachable")
})

tmp7829 := Call2(__e, ns2_1set, symshen_4read_1evaluate_1print, tmp7822)


_ = tmp7829

tmp7830 := MakeNative2(func(__e *ControlFlow, V5768, V5769 Obj) Obj {
tmp7835 := GEq(__e, symnull, V5768)

if True == tmp7835 {
return V5769
} else {
tmp7831 := PrimGuard1(__e, symstr, PrimStr, V5768)

let__7752 := tmp7831
_ = let__7752

tmp7832 := Call1(__e, SymFn(symexternal), V5768)


let__7753 := tmp7832
_ = let__7753

tmp7833 := MakeNative1(func(__e *ControlFlow, Z5772 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4pui_1h), let__7752, let__7753, Z5772)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp7833, V5769)


}


panic("unreachable")
})

tmp7836 := Call2(__e, ns2_1set, symshen_4package_1user_1input, tmp7830)


_ = tmp7836

tmp7837 := MakeNative3(func(__e *ControlFlow, V5777, V5778, V5779 Obj) Obj {
tmp7875 := GConsP(__e, V5779)

var ifres7865 Obj

if True == tmp7875 {
tmp7873 := GHd(__e, V5779)

tmp7874 := GEq(__e, symfn, tmp7873)

var ifres7866 Obj

if True == tmp7874 {
tmp7871 := GTl(__e, V5779)

tmp7872 := GConsP(__e, tmp7871)

var ifres7867 Obj

if True == tmp7872 {
tmp7868 := GTl(__e, V5779)

tmp7869 := GTl(__e, tmp7868)

tmp7870 := GEq(__e, Nil, tmp7869)

ifres7867 = tmp7870


} else {
ifres7867 = False


}

ifres7866 = ifres7867


} else {
ifres7866 = False


}

ifres7865 = ifres7866


} else {
ifres7865 = False


}

if True == ifres7865 {
tmp7843 := GTl(__e, V5779)

tmp7844 := GHd(__e, tmp7843)

tmp7845 := Call3(__e, SymFn(symshen_4internal_2), tmp7844, V5777, V5778)


if True == tmp7845 {
tmp7838 := GTl(__e, V5779)

tmp7839 := GHd(__e, tmp7838)

tmp7840 := Call2(__e, SymFn(symshen_4intern_1in_1package), V5777, tmp7839)


tmp7841 := GCons(__e, tmp7840, Nil)

return GCons(__e, symfn, tmp7841)


} else {
return V5779
}


} else {
tmp7863 := GConsP(__e, V5779)

if True == tmp7863 {
tmp7860 := GHd(__e, V5779)

tmp7861 := Call3(__e, SymFn(symshen_4internal_2), tmp7860, V5777, V5778)


if True == tmp7861 {
tmp7846 := GHd(__e, V5779)

tmp7847 := Call2(__e, SymFn(symshen_4intern_1in_1package), V5777, tmp7846)


tmp7848 := MakeNative1(func(__e *ControlFlow, Z5780 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4pui_1h), V5777, V5778, Z5780)
panic("unreachable")
})

tmp7849 := GTl(__e, V5779)

tmp7850 := Call2(__e, SymFn(symmap), tmp7848, tmp7849)


return GCons(__e, tmp7847, tmp7850)


} else {
tmp7857 := GHd(__e, V5779)

tmp7858 := GConsP(__e, tmp7857)

if True == tmp7858 {
tmp7851 := MakeNative1(func(__e *ControlFlow, Z5781 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4pui_1h), V5777, V5778, Z5781)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp7851, V5779)


} else {
tmp7852 := GHd(__e, V5779)

tmp7853 := MakeNative1(func(__e *ControlFlow, Z5782 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4pui_1h), V5777, V5778, Z5782)
panic("unreachable")
})

tmp7854 := GTl(__e, V5779)

tmp7855 := Call2(__e, SymFn(symmap), tmp7853, tmp7854)


return GCons(__e, tmp7852, tmp7855)


}


}


} else {
return V5779
}


}


panic("unreachable")
})

tmp7876 := Call2(__e, ns2_1set, symshen_4pui_1h, tmp7837)


_ = tmp7876

tmp7877 := MakeNative0(func(__e *ControlFlow) Obj {
tmp7878 := Call0(__e, SymFn(symit))


tmp7879 := Call1(__e, SymFn(symshen_4trim_1it), tmp7878)


tmp7880 := GValue(__e, symshen_4_dhistory_d)

tmp7881 := GCons(__e, tmp7879, tmp7880)

return PrimGuard2(__e, symset, PrimSet, symshen_4_dhistory_d, tmp7881)


panic("unreachable")
})

tmp7882 := Call2(__e, ns2_1set, symshen_4update_1history, tmp7877)


_ = tmp7882

tmp7883 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V5783 := __self1
_ = V5783
tmp7890 := Call1(__e, SymFn(symshen_4_7string_2), V5783)


var ifres7886 Obj

if True == tmp7890 {
tmp7887 := Call1(__e, SymFn(symhdstr), V5783)


tmp7888 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp7887)

tmp7889 := Call1(__e, SymFn(symshen_4whitespace_2), tmp7888)


ifres7886 = tmp7889


} else {
ifres7886 = False


}

if True == ifres7886 {
if SymFn(symshen_4trim_1it) == __self {
__self1 = PrimGuard1(__e, symtlstr, PrimTailString, V5783)
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4trim_1it), PrimGuard1(__e, symtlstr, PrimTailString, V5783))


} else {
return V5783
}


panic("unreachable")
})
return __self
})()

tmp7891 := Call2(__e, ns2_1set, symshen_4trim_1it, tmp7883)


_ = tmp7891

tmp7892 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V5802 := __self1
_ = V5802
V5803 := __self2
_ = V5803
V5804 := __self3
_ = V5804
tmp7996 := GConsP(__e, V5802)

var ifres7972 Obj

if True == tmp7996 {
tmp7994 := GTl(__e, V5802)

tmp7995 := GEq(__e, Nil, tmp7994)

var ifres7973 Obj

if True == tmp7995 {
tmp7993 := GConsP(__e, V5803)

var ifres7974 Obj

if True == tmp7993 {
tmp7991 := GHd(__e, V5803)

tmp7992 := Call1(__e, SymFn(symshen_4_7string_2), tmp7991)


var ifres7975 Obj

if True == tmp7992 {
tmp7988 := GHd(__e, V5803)

tmp7989 := Call1(__e, SymFn(symhdstr), tmp7988)


tmp7990 := GEq(__e, MakeString("!"), tmp7989)

var ifres7976 Obj

if True == tmp7990 {
tmp7985 := GHd(__e, V5803)

tmp7987 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, tmp7985))


var ifres7977 Obj

if True == tmp7987 {
tmp7981 := GHd(__e, V5803)

tmp7983 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, tmp7981))


tmp7984 := GEq(__e, MakeString("!"), tmp7983)

var ifres7978 Obj

if True == tmp7984 {
tmp7979 := GTl(__e, V5803)

tmp7980 := GConsP(__e, tmp7979)

ifres7978 = tmp7980


} else {
ifres7978 = False


}

ifres7977 = ifres7978


} else {
ifres7977 = False


}

ifres7976 = ifres7977


} else {
ifres7976 = False


}

ifres7975 = ifres7976


} else {
ifres7975 = False


}

ifres7974 = ifres7975


} else {
ifres7974 = False


}

ifres7973 = ifres7974


} else {
ifres7973 = False


}

ifres7972 = ifres7973


} else {
ifres7972 = False


}

if True == ifres7972 {
tmp7893 := GTl(__e, V5803)

tmp7894 := GHd(__e, tmp7893)

tmp7895 := Call1(__e, SymFn(symread_1from_1string), tmp7894)


let__7754 := tmp7895
_ = let__7754

tmp7896 := GTl(__e, V5803)

tmp7897 := GHd(__e, tmp7896)

tmp7898 := GTl(__e, V5803)

tmp7899 := GCons(__e, tmp7897, tmp7898)

tmp7900 := PrimGuard2(__e, symset, PrimSet, symshen_4_dhistory_d, tmp7899)

let__7755 := tmp7900
_ = let__7755

tmp7901 := GTl(__e, V5803)

tmp7902 := GHd(__e, tmp7901)

tmp7903 := Call3(__e, SymFn(symshen_4app), tmp7902, MakeString("\n"), symshen_4a)


tmp7904 := Call0(__e, SymFn(symstoutput))


tmp7905 := Call2(__e, SymFn(sympr), tmp7903, tmp7904)


let__7756 := tmp7905
_ = let__7756

if SymFn(symshen_4evaluate_1lineread) == __self {
__self1, __self2, __self3 = let__7754, let__7755, V5804
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4evaluate_1lineread), let__7754, let__7755, V5804)


} else {
tmp7970 := GConsP(__e, V5802)

var ifres7958 Obj

if True == tmp7970 {
tmp7968 := GTl(__e, V5802)

tmp7969 := GEq(__e, Nil, tmp7968)

var ifres7959 Obj

if True == tmp7969 {
tmp7967 := GConsP(__e, V5803)

var ifres7960 Obj

if True == tmp7967 {
tmp7965 := GHd(__e, V5803)

tmp7966 := Call1(__e, SymFn(symshen_4_7string_2), tmp7965)


var ifres7961 Obj

if True == tmp7966 {
tmp7962 := GHd(__e, V5803)

tmp7963 := Call1(__e, SymFn(symhdstr), tmp7962)


tmp7964 := GEq(__e, MakeString("!"), tmp7963)

ifres7961 = tmp7964


} else {
ifres7961 = False


}

ifres7960 = ifres7961


} else {
ifres7960 = False


}

ifres7959 = ifres7960


} else {
ifres7959 = False


}

ifres7958 = ifres7959


} else {
ifres7958 = False


}

if True == ifres7958 {
tmp7911 := GHd(__e, V5803)

tmp7913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(PrimGuard1(__e, symtlstr, PrimTailString, tmp7911), MakeString(""))
}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, tmp7911)
__typedArg1 := MakeString("")
return Call2(__e, SymFn(sym_a), __typedArg0, __typedArg1)
})()

var ifres7906 Obj

if True == tmp7913 {
ifres7906 = Nil


} else {
tmp7907 := GHd(__e, V5803)

tmp7909 := Call1(__e, SymFn(symread_1from_1string), PrimGuard1(__e, symtlstr, PrimTailString, tmp7907))


tmp7910 := GHd(__e, tmp7909)

ifres7906 = tmp7910


}

let__7757 := ifres7906
_ = let__7757

tmp7914 := GHd(__e, V5803)

tmp7915 := PrimGuard1(__e, symtlstr, PrimTailString, tmp7914)

tmp7916 := GTl(__e, V5803)

tmp7917 := Call3(__e, SymFn(symshen_4use_1history), let__7757, tmp7915, tmp7916)


let__7758 := tmp7917
_ = let__7758

tmp7918 := Call3(__e, SymFn(symshen_4app), let__7758, MakeString("\n"), symshen_4a)


tmp7919 := Call0(__e, SymFn(symstoutput))


tmp7920 := Call2(__e, SymFn(sympr), tmp7918, tmp7919)


let__7759 := tmp7920
_ = let__7759

tmp7921 := Call1(__e, SymFn(symread_1from_1string), let__7758)


let__7760 := tmp7921
_ = let__7760

tmp7922 := GTl(__e, V5803)

tmp7923 := GCons(__e, let__7758, tmp7922)

tmp7924 := PrimGuard2(__e, symset, PrimSet, symshen_4_dhistory_d, tmp7923)

let__7761 := tmp7924
_ = let__7761

if SymFn(symshen_4evaluate_1lineread) == __self {
__self1, __self2, __self3 = let__7760, let__7761, V5804
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4evaluate_1lineread), let__7760, let__7761, V5804)


} else {
tmp7956 := GConsP(__e, V5802)

var ifres7944 Obj

if True == tmp7956 {
tmp7954 := GTl(__e, V5802)

tmp7955 := GEq(__e, Nil, tmp7954)

var ifres7945 Obj

if True == tmp7955 {
tmp7953 := GConsP(__e, V5803)

var ifres7946 Obj

if True == tmp7953 {
tmp7951 := GHd(__e, V5803)

tmp7952 := Call1(__e, SymFn(symshen_4_7string_2), tmp7951)


var ifres7947 Obj

if True == tmp7952 {
tmp7948 := GHd(__e, V5803)

tmp7949 := Call1(__e, SymFn(symhdstr), tmp7948)


tmp7950 := GEq(__e, MakeString("%"), tmp7949)

ifres7947 = tmp7950


} else {
ifres7947 = False


}

ifres7946 = ifres7947


} else {
ifres7946 = False


}

ifres7945 = ifres7946


} else {
ifres7945 = False


}

ifres7944 = ifres7945


} else {
ifres7944 = False


}

if True == ifres7944 {
tmp7930 := GHd(__e, V5803)

tmp7932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(PrimGuard1(__e, symtlstr, PrimTailString, tmp7930), MakeString(""))
}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, tmp7930)
__typedArg1 := MakeString("")
return Call2(__e, SymFn(sym_a), __typedArg0, __typedArg1)
})()

var ifres7925 Obj

if True == tmp7932 {
ifres7925 = Nil


} else {
tmp7926 := GHd(__e, V5803)

tmp7928 := Call1(__e, SymFn(symread_1from_1string), PrimGuard1(__e, symtlstr, PrimTailString, tmp7926))


tmp7929 := GHd(__e, tmp7928)

ifres7925 = tmp7929


}

let__7762 := ifres7925
_ = let__7762

tmp7933 := GHd(__e, V5803)

tmp7934 := PrimGuard1(__e, symtlstr, PrimTailString, tmp7933)

tmp7935 := GTl(__e, V5803)

tmp7936 := Call3(__e, SymFn(symshen_4peek_1history), let__7762, tmp7934, tmp7935)


let__7763 := tmp7936
_ = let__7763

tmp7937 := GTl(__e, V5803)

tmp7938 := PrimGuard2(__e, symset, PrimSet, symshen_4_dhistory_d, tmp7937)

let__7764 := tmp7938
_ = let__7764

return TailCall0(__e, SymFn(symabort))


} else {
tmp7942 := GEq(__e, True, V5804)

if True == tmp7942 {
return TailCall1(__e, SymFn(symshen_4check_1eval_1and_1print), V5802)
} else {
tmp7940 := GEq(__e, False, V5804)

if True == tmp7940 {
return TailCall1(__e, SymFn(symshen_4eval_1and_1print), V5802)
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

tmp7997 := Call2(__e, ns2_1set, symshen_4evaluate_1lineread, tmp7892)


_ = tmp7997

tmp7998 := MakeNative3(func(__e *ControlFlow, V5816, V5817, V5818 Obj) Obj {
tmp8004 := PrimGuard1(__e, syminteger_2, PrimIsInteger, V5816)

if True == tmp8004 {
tmp7999 := GAdd(__e, MakeInteger(1), V5816)

tmp8000 := Call1(__e, SymFn(symreverse), V5818)


return TailCall2(__e, SymFn(symnth), tmp7999, tmp8000)


} else {
tmp8002 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V5816)

if True == tmp8002 {
return TailCall2(__e, SymFn(symshen_4string_1match), V5817, V5818)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("! expects a number or a symbol\n"))
}


}


panic("unreachable")
})

tmp8005 := Call2(__e, ns2_1set, symshen_4use_1history, tmp7998)


_ = tmp8005

tmp8006 := MakeNative3(func(__e *ControlFlow, V5819, V5820, V5821 Obj) Obj {
tmp8019 := PrimGuard1(__e, syminteger_2, PrimIsInteger, V5819)

if True == tmp8019 {
tmp8007 := GAdd(__e, MakeInteger(1), V5819)

tmp8008 := Call1(__e, SymFn(symreverse), V5821)


tmp8009 := Call2(__e, SymFn(symnth), tmp8007, tmp8008)


tmp8010 := Call3(__e, SymFn(symshen_4app), tmp8009, MakeString(""), symshen_4a)


tmp8011 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("\n"), tmp8010)

tmp8012 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp8011, tmp8012)


} else {
tmp8017 := GEq(__e, V5820, MakeString(""))

var ifres8015 Obj

if True == tmp8017 {
ifres8015 = True


} else {
tmp8016 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V5819)

ifres8015 = tmp8016


}

if True == ifres8015 {
tmp8013 := Call1(__e, SymFn(symreverse), V5821)


return TailCall3(__e, SymFn(symshen_4recursive_1string_1match), MakeInteger(0), V5820, tmp8013)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("% expects a number or a symbol\n"))
}


}


panic("unreachable")
})

tmp8020 := Call2(__e, ns2_1set, symshen_4peek_1history, tmp8006)


_ = tmp8020

tmp8021 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V5831 := __self1
_ = V5831
V5832 := __self2
_ = V5832
tmp8031 := GEq(__e, Nil, V5832)

if True == tmp8031 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("\ninput not found"))
} else {
tmp8029 := GConsP(__e, V5832)

var ifres8026 Obj

if True == tmp8029 {
tmp8027 := GHd(__e, V5832)

tmp8028 := Call2(__e, SymFn(symshen_4string_1prefix_2), V5831, tmp8027)


ifres8026 = tmp8028


} else {
ifres8026 = False


}

if True == ifres8026 {
return GHd(__e, V5832)
} else {
tmp8024 := GConsP(__e, V5832)

if True == tmp8024 {
tmp8022 := GTl(__e, V5832)

if SymFn(symshen_4string_1match) == __self {
__self1, __self2 = V5831, tmp8022
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4string_1match), V5831, tmp8022)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.string-match"))
}


}


}


panic("unreachable")
})
return __self
})()

tmp8032 := Call2(__e, ns2_1set, symshen_4string_1match, tmp8021)


_ = tmp8032

tmp8033 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V5840 := __self1
_ = V5840
V5841 := __self2
_ = V5841
tmp8065 := GEq(__e, MakeString(""), V5840)

if True == tmp8065 {
return True
} else {
tmp8063 := Call1(__e, SymFn(symshen_4_7string_2), V5840)


var ifres8059 Obj

if True == tmp8063 {
tmp8060 := Call1(__e, SymFn(symhdstr), V5840)


tmp8061 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp8060)

tmp8062 := Call1(__e, SymFn(symshen_4whitespace_2), tmp8061)


ifres8059 = tmp8062


} else {
ifres8059 = False


}

if True == ifres8059 {
if SymFn(symshen_4string_1prefix_2) == __self {
__self1, __self2 = PrimGuard1(__e, symtlstr, PrimTailString, V5840), V5841
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4string_1prefix_2), PrimGuard1(__e, symtlstr, PrimTailString, V5840), V5841)


} else {
tmp8057 := Call1(__e, SymFn(symshen_4_7string_2), V5841)


var ifres8053 Obj

if True == tmp8057 {
tmp8054 := Call1(__e, SymFn(symhdstr), V5841)


tmp8055 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp8054)

tmp8056 := Call1(__e, SymFn(symshen_4whitespace_2), tmp8055)


ifres8053 = tmp8056


} else {
ifres8053 = False


}

if True == ifres8053 {
if SymFn(symshen_4string_1prefix_2) == __self {
__self1, __self2 = V5840, PrimGuard1(__e, symtlstr, PrimTailString, V5841)
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4string_1prefix_2), V5840, PrimGuard1(__e, symtlstr, PrimTailString, V5841))


} else {
tmp8051 := Call1(__e, SymFn(symshen_4_7string_2), V5841)


var ifres8048 Obj

if True == tmp8051 {
tmp8049 := Call1(__e, SymFn(symhdstr), V5841)


tmp8050 := GEq(__e, MakeString("("), tmp8049)

ifres8048 = tmp8050


} else {
ifres8048 = False


}

if True == ifres8048 {
if SymFn(symshen_4string_1prefix_2) == __self {
__self1, __self2 = V5840, PrimGuard1(__e, symtlstr, PrimTailString, V5841)
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4string_1prefix_2), V5840, PrimGuard1(__e, symtlstr, PrimTailString, V5841))


} else {
tmp8046 := Call1(__e, SymFn(symshen_4_7string_2), V5840)


var ifres8040 Obj

if True == tmp8046 {
tmp8045 := Call1(__e, SymFn(symshen_4_7string_2), V5841)


var ifres8041 Obj

if True == tmp8045 {
tmp8042 := Call1(__e, SymFn(symhdstr), V5840)


tmp8043 := Call1(__e, SymFn(symhdstr), V5841)


tmp8044 := GEq(__e, tmp8042, tmp8043)

ifres8041 = tmp8044


} else {
ifres8041 = False


}

ifres8040 = ifres8041


} else {
ifres8040 = False


}

if True == ifres8040 {
tmp8037 := PrimGuard1(__e, symtlstr, PrimTailString, V5840)

if SymFn(symshen_4string_1prefix_2) == __self {
__self1, __self2 = tmp8037, PrimGuard1(__e, symtlstr, PrimTailString, V5841)
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4string_1prefix_2), tmp8037, PrimGuard1(__e, symtlstr, PrimTailString, V5841))


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

tmp8066 := Call2(__e, ns2_1set, symshen_4string_1prefix_2, tmp8033)


_ = tmp8066

tmp8067 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V5852 := __self1
_ = V5852
V5853 := __self2
_ = V5853
V5854 := __self3
_ = V5854
tmp8082 := GEq(__e, Nil, V5854)

if True == tmp8082 {
return symshen_4skip
} else {
tmp8080 := GConsP(__e, V5854)

if True == tmp8080 {
tmp8075 := GHd(__e, V5854)

tmp8076 := Call2(__e, SymFn(symshen_4string_1prefix_2), V5853, tmp8075)


var ifres8068 Obj

if True == tmp8076 {
tmp8069 := GHd(__e, V5854)

tmp8070 := Call3(__e, SymFn(symshen_4app), tmp8069, MakeString("\n"), symshen_4a)


tmp8072 := Call3(__e, SymFn(symshen_4app), V5852, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(". "), tmp8070), symshen_4a)


tmp8073 := Call0(__e, SymFn(symstoutput))


tmp8074 := Call2(__e, SymFn(sympr), tmp8072, tmp8073)


ifres8068 = tmp8074


} else {
ifres8068 = symshen_4skip


}

_ = ifres8068

tmp8077 := GAdd(__e, V5852, MakeInteger(1))

tmp8078 := GTl(__e, V5854)

if SymFn(symshen_4recursive_1string_1match) == __self {
__self1, __self2, __self3 = tmp8077, V5853, tmp8078
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4recursive_1string_1match), tmp8077, V5853, tmp8078)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.recursive-string-match"))
}


}


panic("unreachable")
})
return __self
})()

return TailCall2(__e, ns2_1set, symshen_4recursive_1string_1match, tmp8067)




panic("unreachable")
})

