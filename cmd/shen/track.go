package main

import . "github.com/pyrex41/shen-go/kl"

var TrackMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp12797 := MakeNative1(func(__e *ControlFlow, V5886 Obj) Obj {
tmp12798 := Call3(__e, SymFn(symshen_4app), V5886, MakeString(";\n"), symshen_4a)


tmp12799 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("partial function "), tmp12798)

tmp12800 := Call0(__e, SymFn(symstoutput))


tmp12801 := Call2(__e, SymFn(sympr), tmp12799, tmp12800)


_ = tmp12801

tmp12809 := Call1(__e, SymFn(symshen_4tracked_2), V5886)


tmp12810 := GNot(__e, tmp12809)

var ifres12805 Obj

if True == tmp12810 {
tmp12806 := Call3(__e, SymFn(symshen_4app), V5886, MakeString("? "), symshen_4a)


tmp12808 := Call1(__e, SymFn(symy_1or_1n_2), PrimGuard2(__e, symcn, PrimStringConcat, MakeString("track "), tmp12806))


ifres12805 = tmp12808


} else {
ifres12805 = False


}

var ifres12802 Obj

if True == ifres12805 {
tmp12803 := Call1(__e, SymFn(symps), V5886)


tmp12804 := Call1(__e, SymFn(symshen_4track_1function), tmp12803)


ifres12802 = tmp12804


} else {
ifres12802 = symshen_4ok


}

_ = ifres12802

return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("aborted"))


panic("unreachable")
})

tmp12811 := Call2(__e, ns2_1set, symshen_4f_1error, tmp12797)


_ = tmp12811

tmp12812 := MakeNative1(func(__e *ControlFlow, V5887 Obj) Obj {
tmp12813 := GValue(__e, symshen_4_dtracking_d)

return TailCall2(__e, SymFn(symelement_2), V5887, tmp12813)


panic("unreachable")
})

tmp12814 := Call2(__e, ns2_1set, symshen_4tracked_2, tmp12812)


_ = tmp12814

tmp12815 := MakeNative1(func(__e *ControlFlow, V5888 Obj) Obj {
tmp12816 := Call1(__e, SymFn(symps), V5888)


let__12786 := tmp12816
_ = let__12786

return TailCall1(__e, SymFn(symshen_4track_1function), let__12786)


panic("unreachable")
})

tmp12817 := Call2(__e, ns2_1set, symtrack, tmp12815)


_ = tmp12817

tmp12818 := MakeNative1(func(__e *ControlFlow, V5892 Obj) Obj {
tmp12867 := GConsP(__e, V5892)

var ifres12846 Obj

if True == tmp12867 {
tmp12865 := GHd(__e, V5892)

tmp12866 := GEq(__e, symdefun, tmp12865)

var ifres12847 Obj

if True == tmp12866 {
tmp12863 := GTl(__e, V5892)

tmp12864 := GConsP(__e, tmp12863)

var ifres12848 Obj

if True == tmp12864 {
tmp12860 := GTl(__e, V5892)

tmp12861 := GTl(__e, tmp12860)

tmp12862 := GConsP(__e, tmp12861)

var ifres12849 Obj

if True == tmp12862 {
tmp12856 := GTl(__e, V5892)

tmp12857 := GTl(__e, tmp12856)

tmp12858 := GTl(__e, tmp12857)

tmp12859 := GConsP(__e, tmp12858)

var ifres12850 Obj

if True == tmp12859 {
tmp12851 := GTl(__e, V5892)

tmp12852 := GTl(__e, tmp12851)

tmp12853 := GTl(__e, tmp12852)

tmp12854 := GTl(__e, tmp12853)

tmp12855 := GEq(__e, Nil, tmp12854)

ifres12850 = tmp12855


} else {
ifres12850 = False


}

ifres12849 = ifres12850


} else {
ifres12849 = False


}

ifres12848 = ifres12849


} else {
ifres12848 = False


}

ifres12847 = ifres12848


} else {
ifres12847 = False


}

ifres12846 = ifres12847


} else {
ifres12846 = False


}

if True == ifres12846 {
tmp12819 := GTl(__e, V5892)

tmp12820 := GHd(__e, tmp12819)

tmp12821 := GTl(__e, V5892)

tmp12822 := GTl(__e, tmp12821)

tmp12823 := GHd(__e, tmp12822)

tmp12824 := GTl(__e, V5892)

tmp12825 := GHd(__e, tmp12824)

tmp12826 := GTl(__e, V5892)

tmp12827 := GTl(__e, tmp12826)

tmp12828 := GHd(__e, tmp12827)

tmp12829 := GTl(__e, V5892)

tmp12830 := GTl(__e, tmp12829)

tmp12831 := GTl(__e, tmp12830)

tmp12832 := GHd(__e, tmp12831)

tmp12833 := Call3(__e, SymFn(symshen_4insert_1tracking_1code), tmp12825, tmp12828, tmp12832)


tmp12834 := GCons(__e, tmp12833, Nil)

tmp12835 := GCons(__e, tmp12823, tmp12834)

tmp12836 := GCons(__e, tmp12820, tmp12835)

tmp12837 := GCons(__e, symdefun, tmp12836)

let__12787 := tmp12837
_ = let__12787

tmp12838 := Call1(__e, SymFn(symeval_1kl), let__12787)


let__12788 := tmp12838
_ = let__12788

tmp12839 := GTl(__e, V5892)

tmp12840 := GHd(__e, tmp12839)

tmp12841 := GValue(__e, symshen_4_dtracking_d)

tmp12842 := Call2(__e, SymFn(symadjoin), tmp12840, tmp12841)


tmp12843 := PrimGuard2(__e, symset, PrimSet, symshen_4_dtracking_d, tmp12842)

let__12789 := tmp12843
_ = let__12789

tmp12844 := GTl(__e, V5892)

return GHd(__e, tmp12844)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.track-function"))
}


panic("unreachable")
})

tmp12868 := Call2(__e, ns2_1set, symshen_4track_1function, tmp12818)


_ = tmp12868

tmp12869 := MakeNative3(func(__e *ControlFlow, V5896, V5897, V5898 Obj) Obj {
tmp12870 := GCons(__e, symshen_4_dcall_d, Nil)

tmp12871 := GCons(__e, symvalue, tmp12870)

tmp12872 := GCons(__e, MakeInteger(1), Nil)

tmp12873 := GCons(__e, tmp12871, tmp12872)

tmp12874 := GCons(__e, sym_7, tmp12873)

tmp12875 := GCons(__e, tmp12874, Nil)

tmp12876 := GCons(__e, symshen_4_dcall_d, tmp12875)

tmp12877 := GCons(__e, symset, tmp12876)

tmp12878 := GCons(__e, symshen_4_dcall_d, Nil)

tmp12879 := GCons(__e, symvalue, tmp12878)

tmp12880 := Call2(__e, SymFn(symshen_4prolog_1track), V5898, V5897)


tmp12881 := Call1(__e, SymFn(symshen_4cons_1form), tmp12880)


tmp12882 := GCons(__e, tmp12881, Nil)

tmp12883 := GCons(__e, V5896, tmp12882)

tmp12884 := GCons(__e, tmp12879, tmp12883)

tmp12885 := GCons(__e, symshen_4input_1track, tmp12884)

tmp12886 := GCons(__e, symshen_4terpri_1or_1read_1char, Nil)

tmp12887 := GCons(__e, symshen_4_dcall_d, Nil)

tmp12888 := GCons(__e, symvalue, tmp12887)

tmp12889 := GCons(__e, symResult, Nil)

tmp12890 := GCons(__e, V5896, tmp12889)

tmp12891 := GCons(__e, tmp12888, tmp12890)

tmp12892 := GCons(__e, symshen_4output_1track, tmp12891)

tmp12893 := GCons(__e, symshen_4_dcall_d, Nil)

tmp12894 := GCons(__e, symvalue, tmp12893)

tmp12895 := GCons(__e, MakeInteger(1), Nil)

tmp12896 := GCons(__e, tmp12894, tmp12895)

tmp12897 := GCons(__e, sym_1, tmp12896)

tmp12898 := GCons(__e, tmp12897, Nil)

tmp12899 := GCons(__e, symshen_4_dcall_d, tmp12898)

tmp12900 := GCons(__e, symset, tmp12899)

tmp12901 := GCons(__e, symshen_4terpri_1or_1read_1char, Nil)

tmp12902 := GCons(__e, symResult, Nil)

tmp12903 := GCons(__e, tmp12901, tmp12902)

tmp12904 := GCons(__e, symdo, tmp12903)

tmp12905 := GCons(__e, tmp12904, Nil)

tmp12906 := GCons(__e, tmp12900, tmp12905)

tmp12907 := GCons(__e, symdo, tmp12906)

tmp12908 := GCons(__e, tmp12907, Nil)

tmp12909 := GCons(__e, tmp12892, tmp12908)

tmp12910 := GCons(__e, symdo, tmp12909)

tmp12911 := GCons(__e, tmp12910, Nil)

tmp12912 := GCons(__e, V5898, tmp12911)

tmp12913 := GCons(__e, symResult, tmp12912)

tmp12914 := GCons(__e, symlet, tmp12913)

tmp12915 := GCons(__e, tmp12914, Nil)

tmp12916 := GCons(__e, tmp12886, tmp12915)

tmp12917 := GCons(__e, symdo, tmp12916)

tmp12918 := GCons(__e, tmp12917, Nil)

tmp12919 := GCons(__e, tmp12885, tmp12918)

tmp12920 := GCons(__e, symdo, tmp12919)

tmp12921 := GCons(__e, tmp12920, Nil)

tmp12922 := GCons(__e, tmp12877, tmp12921)

return GCons(__e, symdo, tmp12922)


panic("unreachable")
})

tmp12923 := Call2(__e, ns2_1set, symshen_4insert_1tracking_1code, tmp12869)


_ = tmp12923

tmp12924 := MakeNative2(func(__e *ControlFlow, V5899, V5900 Obj) Obj {
tmp12927 := Call2(__e, SymFn(symoccurrences), symshen_4incinfs, V5899)


tmp12928 := GEq(__e, tmp12927, MakeInteger(0))

if True == tmp12928 {
return V5900
} else {
tmp12925 := Call1(__e, SymFn(symshen_4vector_1parameter), V5900)


return TailCall2(__e, SymFn(symshen_4vector_1dereference), V5900, tmp12925)


}


panic("unreachable")
})

tmp12929 := Call2(__e, ns2_1set, symshen_4prolog_1track, tmp12924)


_ = tmp12929

tmp12930 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V5903 := __self1
_ = V5903
tmp12955 := GEq(__e, Nil, V5903)

if True == tmp12955 {
return Nil
} else {
tmp12953 := GConsP(__e, V5903)

var ifres12935 Obj

if True == tmp12953 {
tmp12951 := GTl(__e, V5903)

tmp12952 := GConsP(__e, tmp12951)

var ifres12936 Obj

if True == tmp12952 {
tmp12948 := GTl(__e, V5903)

tmp12949 := GTl(__e, tmp12948)

tmp12950 := GConsP(__e, tmp12949)

var ifres12937 Obj

if True == tmp12950 {
tmp12944 := GTl(__e, V5903)

tmp12945 := GTl(__e, tmp12944)

tmp12946 := GTl(__e, tmp12945)

tmp12947 := GConsP(__e, tmp12946)

var ifres12938 Obj

if True == tmp12947 {
tmp12939 := GTl(__e, V5903)

tmp12940 := GTl(__e, tmp12939)

tmp12941 := GTl(__e, tmp12940)

tmp12942 := GTl(__e, tmp12941)

tmp12943 := GEq(__e, Nil, tmp12942)

ifres12938 = tmp12943


} else {
ifres12938 = False


}

ifres12937 = ifres12938


} else {
ifres12937 = False


}

ifres12936 = ifres12937


} else {
ifres12936 = False


}

ifres12935 = ifres12936


} else {
ifres12935 = False


}

if True == ifres12935 {
return GHd(__e, V5903)
} else {
tmp12933 := GConsP(__e, V5903)

if True == tmp12933 {
tmp12931 := GTl(__e, V5903)

if SymFn(symshen_4vector_1parameter) == __self {
__self1 = tmp12931
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4vector_1parameter), tmp12931)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4vector_1parameter)
}


}


}


panic("unreachable")
})
return __self
})()

tmp12956 := Call2(__e, ns2_1set, symshen_4vector_1parameter, tmp12930)


_ = tmp12956

tmp12957 := MakeNative2(func(__e *ControlFlow, V5906, V5907 Obj) Obj {
tmp12987 := GEq(__e, Nil, V5907)

if True == tmp12987 {
return V5906
} else {
tmp12985 := GConsP(__e, V5906)

var ifres12967 Obj

if True == tmp12985 {
tmp12983 := GTl(__e, V5906)

tmp12984 := GConsP(__e, tmp12983)

var ifres12968 Obj

if True == tmp12984 {
tmp12980 := GTl(__e, V5906)

tmp12981 := GTl(__e, tmp12980)

tmp12982 := GConsP(__e, tmp12981)

var ifres12969 Obj

if True == tmp12982 {
tmp12976 := GTl(__e, V5906)

tmp12977 := GTl(__e, tmp12976)

tmp12978 := GTl(__e, tmp12977)

tmp12979 := GConsP(__e, tmp12978)

var ifres12970 Obj

if True == tmp12979 {
tmp12971 := GTl(__e, V5906)

tmp12972 := GTl(__e, tmp12971)

tmp12973 := GTl(__e, tmp12972)

tmp12974 := GTl(__e, tmp12973)

tmp12975 := GEq(__e, Nil, tmp12974)

ifres12970 = tmp12975


} else {
ifres12970 = False


}

ifres12969 = ifres12970


} else {
ifres12969 = False


}

ifres12968 = ifres12969


} else {
ifres12968 = False


}

ifres12967 = ifres12968


} else {
ifres12967 = False


}

if True == ifres12967 {
return V5906
} else {
tmp12965 := GConsP(__e, V5906)

if True == tmp12965 {
tmp12958 := GHd(__e, V5906)

tmp12959 := GCons(__e, V5907, Nil)

tmp12960 := GCons(__e, tmp12958, tmp12959)

tmp12961 := GCons(__e, symshen_4deref, tmp12960)

tmp12962 := GTl(__e, V5906)

tmp12963 := Call2(__e, SymFn(symshen_4vector_1dereference), tmp12962, V5907)


return GCons(__e, tmp12961, tmp12963)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4vector_1dereference)
}


}


}


panic("unreachable")
})

tmp12988 := Call2(__e, ns2_1set, symshen_4vector_1dereference, tmp12957)


_ = tmp12988

tmp12989 := MakeNative1(func(__e *ControlFlow, V5910 Obj) Obj {
tmp12993 := GEq(__e, sym_7, V5910)

if True == tmp12993 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dstep_d, True)
} else {
tmp12991 := GEq(__e, sym_1, V5910)

if True == tmp12991 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dstep_d, False)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("step expects a + or a -.\n"))
}


}


panic("unreachable")
})

tmp12994 := Call2(__e, ns2_1set, symstep, tmp12989)


_ = tmp12994

tmp12995 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dstep_d)
panic("unreachable")
})

tmp12996 := Call2(__e, ns2_1set, symstep_2, tmp12995)


_ = tmp12996

tmp12997 := MakeNative1(func(__e *ControlFlow, V5913 Obj) Obj {
tmp13001 := GEq(__e, sym_7, V5913)

if True == tmp13001 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dspy_d, True)
} else {
tmp12999 := GEq(__e, sym_1, V5913)

if True == tmp12999 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dspy_d, False)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("spy expects a + or a -.\n"))
}


}


panic("unreachable")
})

tmp13002 := Call2(__e, ns2_1set, symspy, tmp12997)


_ = tmp13002

tmp13003 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dspy_d)
panic("unreachable")
})

tmp13004 := Call2(__e, ns2_1set, symspy_2, tmp13003)


_ = tmp13004

tmp13005 := MakeNative0(func(__e *ControlFlow) Obj {
tmp13009 := GValue(__e, symshen_4_dstep_d)

if True == tmp13009 {
tmp13006 := GValue(__e, sym_dstinput_d)

tmp13007 := PrimGuard1(__e, symread_1byte, PrimReadByte, tmp13006)

return TailCall1(__e, SymFn(symshen_4check_1byte), tmp13007)


} else {
return TailCall1(__e, SymFn(symnl), MakeInteger(1))
}


panic("unreachable")
})

tmp13010 := Call2(__e, ns2_1set, symshen_4terpri_1or_1read_1char, tmp13005)


_ = tmp13010

tmp13011 := MakeNative1(func(__e *ControlFlow, V5916 Obj) Obj {
tmp13013 := GEq(__e, MakeInteger(94), V5916)

if True == tmp13013 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("aborted"))
} else {
return True
}


panic("unreachable")
})

tmp13014 := Call2(__e, ns2_1set, symshen_4check_1byte, tmp13011)


_ = tmp13014

tmp13015 := MakeNative3(func(__e *ControlFlow, V5917, V5918, V5919 Obj) Obj {
tmp13016 := Call1(__e, SymFn(symshen_4spaces), V5917)


tmp13017 := Call1(__e, SymFn(symshen_4spaces), V5917)


tmp13018 := Call3(__e, SymFn(symshen_4app), tmp13017, MakeString(""), symshen_4a)


tmp13020 := Call3(__e, SymFn(symshen_4app), V5918, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" \n"), tmp13018), symshen_4a)


tmp13022 := Call3(__e, SymFn(symshen_4app), V5917, PrimGuard2(__e, symcn, PrimStringConcat, MakeString("> Inputs to "), tmp13020), symshen_4a)


tmp13024 := Call3(__e, SymFn(symshen_4app), tmp13016, PrimGuard2(__e, symcn, PrimStringConcat, MakeString("<"), tmp13022), symshen_4a)


tmp13025 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("\n"), tmp13024)

tmp13026 := Call0(__e, SymFn(symstoutput))


tmp13027 := Call2(__e, SymFn(sympr), tmp13025, tmp13026)


_ = tmp13027

return TailCall1(__e, SymFn(symshen_4recursively_1print), V5919)


panic("unreachable")
})

tmp13028 := Call2(__e, ns2_1set, symshen_4input_1track, tmp13015)


_ = tmp13028

tmp13029 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V5922 := __self1
_ = V5922
tmp13039 := GEq(__e, Nil, V5922)

if True == tmp13039 {
tmp13030 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), MakeString(" ==>"), tmp13030)


} else {
tmp13037 := GConsP(__e, V5922)

if True == tmp13037 {
tmp13031 := GHd(__e, V5922)

tmp13032 := Call1(__e, SymFn(symprint), tmp13031)


_ = tmp13032

tmp13033 := Call0(__e, SymFn(symstoutput))


tmp13034 := Call2(__e, SymFn(sympr), MakeString(", "), tmp13033)


_ = tmp13034

tmp13035 := GTl(__e, V5922)

if SymFn(symshen_4recursively_1print) == __self {
__self1 = tmp13035
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4recursively_1print), tmp13035)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.recursively-print"))
}


}


panic("unreachable")
})
return __self
})()

tmp13040 := Call2(__e, ns2_1set, symshen_4recursively_1print, tmp13029)


_ = tmp13040

tmp13041 := MakeNative1(func(__e *ControlFlow, V5923 Obj) Obj {
tmp13045 := GEq(__e, MakeInteger(0), V5923)

if True == tmp13045 {
return MakeString("")
} else {
tmp13043 := Call1(__e, SymFn(symshen_4spaces), GSub(__e, V5923, MakeInteger(1)))


return PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" "), tmp13043)


}


panic("unreachable")
})

tmp13046 := Call2(__e, ns2_1set, symshen_4spaces, tmp13041)


_ = tmp13046

tmp13047 := MakeNative3(func(__e *ControlFlow, V5924, V5925, V5926 Obj) Obj {
tmp13048 := Call1(__e, SymFn(symshen_4spaces), V5924)


tmp13049 := Call1(__e, SymFn(symshen_4spaces), V5924)


tmp13050 := Call3(__e, SymFn(symshen_4app), V5926, MakeString(""), symshen_4s)


tmp13052 := Call3(__e, SymFn(symshen_4app), tmp13049, PrimGuard2(__e, symcn, PrimStringConcat, MakeString("==> "), tmp13050), symshen_4a)


tmp13054 := Call3(__e, SymFn(symshen_4app), V5925, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" \n"), tmp13052), symshen_4a)


tmp13056 := Call3(__e, SymFn(symshen_4app), V5924, PrimGuard2(__e, symcn, PrimStringConcat, MakeString("> Output from "), tmp13054), symshen_4a)


tmp13058 := Call3(__e, SymFn(symshen_4app), tmp13048, PrimGuard2(__e, symcn, PrimStringConcat, MakeString("<"), tmp13056), symshen_4a)


tmp13059 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("\n"), tmp13058)

tmp13060 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp13059, tmp13060)


panic("unreachable")
})

tmp13061 := Call2(__e, ns2_1set, symshen_4output_1track, tmp13047)


_ = tmp13061

tmp13062 := MakeNative1(func(__e *ControlFlow, V5927 Obj) Obj {
tmp13063 := GValue(__e, symshen_4_dtracking_d)

tmp13064 := Call2(__e, SymFn(symremove), V5927, tmp13063)


tmp13065 := PrimGuard2(__e, symset, PrimSet, symshen_4_dtracking_d, tmp13064)

_ = tmp13065

tmp13066 := MakeNative0(func(__e *ControlFlow) Obj {
tmp13067 := Call1(__e, SymFn(symps), V5927)


return TailCall1(__e, SymFn(symeval), tmp13067)


panic("unreachable")
})

tmp13068 := MakeNative1(func(__e *ControlFlow, Z5928 Obj) Obj {
return V5927
panic("unreachable")
})

tmp13069 := Call2(__e, try_1catch, tmp13066, tmp13068)


_ = tmp13069

return V5927


panic("unreachable")
})

tmp13070 := Call2(__e, ns2_1set, symuntrack, tmp13062)


_ = tmp13070

tmp13071 := MakeNative2(func(__e *ControlFlow, V5929, V5930 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4remove_1h), V5929, V5930, Nil)
panic("unreachable")
})

tmp13072 := Call2(__e, ns2_1set, symremove, tmp13071)


_ = tmp13072

tmp13073 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V5940 := __self1
_ = V5940
V5941 := __self2
_ = V5941
V5942 := __self3
_ = V5942
tmp13087 := GEq(__e, Nil, V5941)

if True == tmp13087 {
return TailCall1(__e, SymFn(symreverse), V5942)
} else {
tmp13085 := GConsP(__e, V5941)

var ifres13082 Obj

if True == tmp13085 {
tmp13083 := GHd(__e, V5941)

tmp13084 := GEq(__e, V5940, tmp13083)

ifres13082 = tmp13084


} else {
ifres13082 = False


}

if True == ifres13082 {
tmp13074 := GHd(__e, V5941)

tmp13075 := GTl(__e, V5941)

if SymFn(symshen_4remove_1h) == __self {
__self1, __self2, __self3 = tmp13074, tmp13075, V5942
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4remove_1h), tmp13074, tmp13075, V5942)


} else {
tmp13080 := GConsP(__e, V5941)

if True == tmp13080 {
tmp13076 := GTl(__e, V5941)

tmp13077 := GHd(__e, V5941)

tmp13078 := GCons(__e, tmp13077, V5942)

if SymFn(symshen_4remove_1h) == __self {
__self1, __self2, __self3 = V5940, tmp13076, tmp13078
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4remove_1h), V5940, tmp13076, tmp13078)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.remove-h"))
}


}


}


panic("unreachable")
})
return __self
})()

tmp13088 := Call2(__e, ns2_1set, symshen_4remove_1h, tmp13073)


_ = tmp13088

tmp13089 := MakeNative1(func(__e *ControlFlow, V5943 Obj) Obj {
tmp13090 := GValue(__e, symshen_4_dprofiled_d)

tmp13091 := GCons(__e, V5943, tmp13090)

tmp13092 := PrimGuard2(__e, symset, PrimSet, symshen_4_dprofiled_d, tmp13091)

_ = tmp13092

tmp13093 := Call1(__e, SymFn(symps), V5943)


return TailCall1(__e, SymFn(symshen_4profile_1help), tmp13093)


panic("unreachable")
})

tmp13094 := Call2(__e, ns2_1set, symprofile, tmp13089)


_ = tmp13094

tmp13095 := MakeNative1(func(__e *ControlFlow, V5946 Obj) Obj {
tmp13155 := GConsP(__e, V5946)

var ifres13134 Obj

if True == tmp13155 {
tmp13153 := GHd(__e, V5946)

tmp13154 := GEq(__e, symdefun, tmp13153)

var ifres13135 Obj

if True == tmp13154 {
tmp13151 := GTl(__e, V5946)

tmp13152 := GConsP(__e, tmp13151)

var ifres13136 Obj

if True == tmp13152 {
tmp13148 := GTl(__e, V5946)

tmp13149 := GTl(__e, tmp13148)

tmp13150 := GConsP(__e, tmp13149)

var ifres13137 Obj

if True == tmp13150 {
tmp13144 := GTl(__e, V5946)

tmp13145 := GTl(__e, tmp13144)

tmp13146 := GTl(__e, tmp13145)

tmp13147 := GConsP(__e, tmp13146)

var ifres13138 Obj

if True == tmp13147 {
tmp13139 := GTl(__e, V5946)

tmp13140 := GTl(__e, tmp13139)

tmp13141 := GTl(__e, tmp13140)

tmp13142 := GTl(__e, tmp13141)

tmp13143 := GEq(__e, Nil, tmp13142)

ifres13138 = tmp13143


} else {
ifres13138 = False


}

ifres13137 = ifres13138


} else {
ifres13137 = False


}

ifres13136 = ifres13137


} else {
ifres13136 = False


}

ifres13135 = ifres13136


} else {
ifres13135 = False


}

ifres13134 = ifres13135


} else {
ifres13134 = False


}

if True == ifres13134 {
tmp13096 := Call1(__e, SymFn(symgensym), symshen_4f)


let__12790 := tmp13096
_ = let__12790

tmp13097 := GTl(__e, V5946)

tmp13098 := GHd(__e, tmp13097)

tmp13099 := GTl(__e, V5946)

tmp13100 := GTl(__e, tmp13099)

tmp13101 := GHd(__e, tmp13100)

tmp13102 := GTl(__e, V5946)

tmp13103 := GHd(__e, tmp13102)

tmp13104 := GTl(__e, V5946)

tmp13105 := GTl(__e, tmp13104)

tmp13106 := GHd(__e, tmp13105)

tmp13107 := GTl(__e, V5946)

tmp13108 := GTl(__e, tmp13107)

tmp13109 := GHd(__e, tmp13108)

tmp13110 := GCons(__e, let__12790, tmp13109)

tmp13111 := Call3(__e, SymFn(symshen_4profile_1func), tmp13103, tmp13106, tmp13110)


tmp13112 := GCons(__e, tmp13111, Nil)

tmp13113 := GCons(__e, tmp13101, tmp13112)

tmp13114 := GCons(__e, tmp13098, tmp13113)

tmp13115 := GCons(__e, symdefun, tmp13114)

let__12791 := tmp13115
_ = let__12791

tmp13116 := GTl(__e, V5946)

tmp13117 := GTl(__e, tmp13116)

tmp13118 := GHd(__e, tmp13117)

tmp13119 := GTl(__e, V5946)

tmp13120 := GHd(__e, tmp13119)

tmp13121 := GTl(__e, V5946)

tmp13122 := GTl(__e, tmp13121)

tmp13123 := GTl(__e, tmp13122)

tmp13124 := GHd(__e, tmp13123)

tmp13125 := Call3(__e, SymFn(symsubst), let__12790, tmp13120, tmp13124)


tmp13126 := GCons(__e, tmp13125, Nil)

tmp13127 := GCons(__e, tmp13118, tmp13126)

tmp13128 := GCons(__e, let__12790, tmp13127)

tmp13129 := GCons(__e, symdefun, tmp13128)

let__12792 := tmp13129
_ = let__12792

tmp13130 := Call1(__e, SymFn(symeval_1kl), let__12791)


let__12793 := tmp13130
_ = let__12793

tmp13131 := Call1(__e, SymFn(symeval_1kl), let__12792)


let__12794 := tmp13131
_ = let__12794

tmp13132 := GTl(__e, V5946)

return GHd(__e, tmp13132)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("Cannot profile.\n"))
}


panic("unreachable")
})

tmp13156 := Call2(__e, ns2_1set, symshen_4profile_1help, tmp13095)


_ = tmp13156

tmp13157 := MakeNative1(func(__e *ControlFlow, V5952 Obj) Obj {
tmp13158 := GValue(__e, symshen_4_dprofiled_d)

tmp13159 := Call2(__e, SymFn(symremove), V5952, tmp13158)


tmp13160 := PrimGuard2(__e, symset, PrimSet, symshen_4_dprofiled_d, tmp13159)

_ = tmp13160

tmp13161 := MakeNative0(func(__e *ControlFlow) Obj {
tmp13162 := Call1(__e, SymFn(symps), V5952)


return TailCall1(__e, SymFn(symeval), tmp13162)


panic("unreachable")
})

tmp13163 := MakeNative1(func(__e *ControlFlow, Z5953 Obj) Obj {
return V5952
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp13161, tmp13163)


panic("unreachable")
})

tmp13164 := Call2(__e, ns2_1set, symunprofile, tmp13157)


_ = tmp13164

tmp13165 := MakeNative1(func(__e *ControlFlow, V5954 Obj) Obj {
tmp13166 := GValue(__e, symshen_4_dprofiled_d)

return TailCall2(__e, SymFn(symelement_2), V5954, tmp13166)


panic("unreachable")
})

tmp13167 := Call2(__e, ns2_1set, symshen_4profiled_2, tmp13165)


_ = tmp13167

tmp13168 := MakeNative3(func(__e *ControlFlow, V5955, V5956, V5957 Obj) Obj {
tmp13169 := GCons(__e, symrun, Nil)

tmp13170 := GCons(__e, symget_1time, tmp13169)

tmp13171 := GCons(__e, symrun, Nil)

tmp13172 := GCons(__e, symget_1time, tmp13171)

tmp13173 := GCons(__e, symStart, Nil)

tmp13174 := GCons(__e, tmp13172, tmp13173)

tmp13175 := GCons(__e, sym_1, tmp13174)

tmp13176 := GCons(__e, V5955, Nil)

tmp13177 := GCons(__e, symshen_4get_1profile, tmp13176)

tmp13178 := GCons(__e, symFinish, Nil)

tmp13179 := GCons(__e, tmp13177, tmp13178)

tmp13180 := GCons(__e, sym_7, tmp13179)

tmp13181 := GCons(__e, tmp13180, Nil)

tmp13182 := GCons(__e, V5955, tmp13181)

tmp13183 := GCons(__e, symshen_4put_1profile, tmp13182)

tmp13184 := GCons(__e, symResult, Nil)

tmp13185 := GCons(__e, tmp13183, tmp13184)

tmp13186 := GCons(__e, symRecord, tmp13185)

tmp13187 := GCons(__e, symlet, tmp13186)

tmp13188 := GCons(__e, tmp13187, Nil)

tmp13189 := GCons(__e, tmp13175, tmp13188)

tmp13190 := GCons(__e, symFinish, tmp13189)

tmp13191 := GCons(__e, symlet, tmp13190)

tmp13192 := GCons(__e, tmp13191, Nil)

tmp13193 := GCons(__e, V5957, tmp13192)

tmp13194 := GCons(__e, symResult, tmp13193)

tmp13195 := GCons(__e, symlet, tmp13194)

tmp13196 := GCons(__e, tmp13195, Nil)

tmp13197 := GCons(__e, tmp13170, tmp13196)

tmp13198 := GCons(__e, symStart, tmp13197)

return GCons(__e, symlet, tmp13198)


panic("unreachable")
})

tmp13199 := Call2(__e, ns2_1set, symshen_4profile_1func, tmp13168)


_ = tmp13199

tmp13200 := MakeNative1(func(__e *ControlFlow, V5958 Obj) Obj {
tmp13201 := Call1(__e, SymFn(symshen_4get_1profile), V5958)


let__12795 := tmp13201
_ = let__12795

tmp13202 := Call2(__e, SymFn(symshen_4put_1profile), V5958, MakeInteger(0))


let__12796 := tmp13202
_ = let__12796

return TailCall2(__e, SymFn(sym_8p), V5958, let__12795)


panic("unreachable")
})

tmp13203 := Call2(__e, ns2_1set, symprofile_1results, tmp13200)


_ = tmp13203

tmp13204 := MakeNative1(func(__e *ControlFlow, V5961 Obj) Obj {
tmp13205 := MakeNative0(func(__e *ControlFlow) Obj {
tmp13206 := GValue(__e, sym_dproperty_1vector_d)

return TailCall3(__e, SymFn(symget), V5961, symprofile, tmp13206)


panic("unreachable")
})

tmp13207 := MakeNative1(func(__e *ControlFlow, Z5962 Obj) Obj {
return MakeInteger(0)
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp13205, tmp13207)


panic("unreachable")
})

tmp13208 := Call2(__e, ns2_1set, symshen_4get_1profile, tmp13204)


_ = tmp13208

tmp13209 := MakeNative2(func(__e *ControlFlow, V5963, V5964 Obj) Obj {
tmp13210 := GValue(__e, sym_dproperty_1vector_d)

return TailCall4(__e, SymFn(symput), V5963, symprofile, V5964, tmp13210)


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4put_1profile, tmp13209)




panic("unreachable")
})

