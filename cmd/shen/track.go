package main

import . "github.com/pyrex41/shen-go/kl"

var TrackMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp12777 := MakeNative1(func(__e *ControlFlow, V5871 Obj) Obj {
tmp12778 := Call3(__e, SymFn(symshen_4app), V5871, MakeString(";\n"), symshen_4a)


tmp12779 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("partial function "), tmp12778)

tmp12780 := Call0(__e, SymFn(symstoutput))


tmp12781 := Call2(__e, SymFn(sympr), tmp12779, tmp12780)


_ = tmp12781

tmp12789 := Call1(__e, SymFn(symshen_4tracked_2), V5871)


tmp12790 := GNot(__e, tmp12789)

var ifres12785 Obj

if True == tmp12790 {
tmp12786 := Call3(__e, SymFn(symshen_4app), V5871, MakeString("? "), symshen_4a)


tmp12788 := Call1(__e, SymFn(symy_1or_1n_2), PrimGuard2(__e, symcn, PrimStringConcat, MakeString("track "), tmp12786))


ifres12785 = tmp12788


} else {
ifres12785 = False


}

var ifres12782 Obj

if True == ifres12785 {
tmp12783 := Call1(__e, SymFn(symps), V5871)


tmp12784 := Call1(__e, SymFn(symshen_4track_1function), tmp12783)


ifres12782 = tmp12784


} else {
ifres12782 = symshen_4ok


}

_ = ifres12782

return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("aborted"))


panic("unreachable")
})

tmp12791 := Call2(__e, ns2_1set, symshen_4f_1error, tmp12777)


_ = tmp12791

tmp12792 := MakeNative1(func(__e *ControlFlow, V5872 Obj) Obj {
tmp12793 := GValue(__e, symshen_4_dtracking_d)

return TailCall2(__e, SymFn(symelement_2), V5872, tmp12793)


panic("unreachable")
})

tmp12794 := Call2(__e, ns2_1set, symshen_4tracked_2, tmp12792)


_ = tmp12794

tmp12795 := MakeNative1(func(__e *ControlFlow, V5873 Obj) Obj {
tmp12796 := Call1(__e, SymFn(symps), V5873)


let__12766 := tmp12796
_ = let__12766

return TailCall1(__e, SymFn(symshen_4track_1function), let__12766)


panic("unreachable")
})

tmp12797 := Call2(__e, ns2_1set, symtrack, tmp12795)


_ = tmp12797

tmp12798 := MakeNative1(func(__e *ControlFlow, V5877 Obj) Obj {
tmp12847 := GConsP(__e, V5877)

var ifres12826 Obj

if True == tmp12847 {
tmp12845 := GHd(__e, V5877)

tmp12846 := GEq(__e, symdefun, tmp12845)

var ifres12827 Obj

if True == tmp12846 {
tmp12843 := GTl(__e, V5877)

tmp12844 := GConsP(__e, tmp12843)

var ifres12828 Obj

if True == tmp12844 {
tmp12840 := GTl(__e, V5877)

tmp12841 := GTl(__e, tmp12840)

tmp12842 := GConsP(__e, tmp12841)

var ifres12829 Obj

if True == tmp12842 {
tmp12836 := GTl(__e, V5877)

tmp12837 := GTl(__e, tmp12836)

tmp12838 := GTl(__e, tmp12837)

tmp12839 := GConsP(__e, tmp12838)

var ifres12830 Obj

if True == tmp12839 {
tmp12831 := GTl(__e, V5877)

tmp12832 := GTl(__e, tmp12831)

tmp12833 := GTl(__e, tmp12832)

tmp12834 := GTl(__e, tmp12833)

tmp12835 := GEq(__e, Nil, tmp12834)

ifres12830 = tmp12835


} else {
ifres12830 = False


}

ifres12829 = ifres12830


} else {
ifres12829 = False


}

ifres12828 = ifres12829


} else {
ifres12828 = False


}

ifres12827 = ifres12828


} else {
ifres12827 = False


}

ifres12826 = ifres12827


} else {
ifres12826 = False


}

if True == ifres12826 {
tmp12799 := GTl(__e, V5877)

tmp12800 := GHd(__e, tmp12799)

tmp12801 := GTl(__e, V5877)

tmp12802 := GTl(__e, tmp12801)

tmp12803 := GHd(__e, tmp12802)

tmp12804 := GTl(__e, V5877)

tmp12805 := GHd(__e, tmp12804)

tmp12806 := GTl(__e, V5877)

tmp12807 := GTl(__e, tmp12806)

tmp12808 := GHd(__e, tmp12807)

tmp12809 := GTl(__e, V5877)

tmp12810 := GTl(__e, tmp12809)

tmp12811 := GTl(__e, tmp12810)

tmp12812 := GHd(__e, tmp12811)

tmp12813 := Call3(__e, SymFn(symshen_4insert_1tracking_1code), tmp12805, tmp12808, tmp12812)


tmp12814 := GCons(__e, tmp12813, Nil)

tmp12815 := GCons(__e, tmp12803, tmp12814)

tmp12816 := GCons(__e, tmp12800, tmp12815)

tmp12817 := GCons(__e, symdefun, tmp12816)

let__12767 := tmp12817
_ = let__12767

tmp12818 := Call1(__e, SymFn(symeval_1kl), let__12767)


let__12768 := tmp12818
_ = let__12768

tmp12819 := GTl(__e, V5877)

tmp12820 := GHd(__e, tmp12819)

tmp12821 := GValue(__e, symshen_4_dtracking_d)

tmp12822 := Call2(__e, SymFn(symadjoin), tmp12820, tmp12821)


tmp12823 := PrimGuard2(__e, symset, PrimSet, symshen_4_dtracking_d, tmp12822)

let__12769 := tmp12823
_ = let__12769

tmp12824 := GTl(__e, V5877)

return GHd(__e, tmp12824)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.track-function"))
}


panic("unreachable")
})

tmp12848 := Call2(__e, ns2_1set, symshen_4track_1function, tmp12798)


_ = tmp12848

tmp12849 := MakeNative3(func(__e *ControlFlow, V5881, V5882, V5883 Obj) Obj {
tmp12850 := GCons(__e, symshen_4_dcall_d, Nil)

tmp12851 := GCons(__e, symvalue, tmp12850)

tmp12852 := GCons(__e, MakeInteger(1), Nil)

tmp12853 := GCons(__e, tmp12851, tmp12852)

tmp12854 := GCons(__e, sym_7, tmp12853)

tmp12855 := GCons(__e, tmp12854, Nil)

tmp12856 := GCons(__e, symshen_4_dcall_d, tmp12855)

tmp12857 := GCons(__e, symset, tmp12856)

tmp12858 := GCons(__e, symshen_4_dcall_d, Nil)

tmp12859 := GCons(__e, symvalue, tmp12858)

tmp12860 := Call2(__e, SymFn(symshen_4prolog_1track), V5883, V5882)


tmp12861 := Call1(__e, SymFn(symshen_4cons_1form), tmp12860)


tmp12862 := GCons(__e, tmp12861, Nil)

tmp12863 := GCons(__e, V5881, tmp12862)

tmp12864 := GCons(__e, tmp12859, tmp12863)

tmp12865 := GCons(__e, symshen_4input_1track, tmp12864)

tmp12866 := GCons(__e, symshen_4terpri_1or_1read_1char, Nil)

tmp12867 := GCons(__e, symshen_4_dcall_d, Nil)

tmp12868 := GCons(__e, symvalue, tmp12867)

tmp12869 := GCons(__e, symResult, Nil)

tmp12870 := GCons(__e, V5881, tmp12869)

tmp12871 := GCons(__e, tmp12868, tmp12870)

tmp12872 := GCons(__e, symshen_4output_1track, tmp12871)

tmp12873 := GCons(__e, symshen_4_dcall_d, Nil)

tmp12874 := GCons(__e, symvalue, tmp12873)

tmp12875 := GCons(__e, MakeInteger(1), Nil)

tmp12876 := GCons(__e, tmp12874, tmp12875)

tmp12877 := GCons(__e, sym_1, tmp12876)

tmp12878 := GCons(__e, tmp12877, Nil)

tmp12879 := GCons(__e, symshen_4_dcall_d, tmp12878)

tmp12880 := GCons(__e, symset, tmp12879)

tmp12881 := GCons(__e, symshen_4terpri_1or_1read_1char, Nil)

tmp12882 := GCons(__e, symResult, Nil)

tmp12883 := GCons(__e, tmp12881, tmp12882)

tmp12884 := GCons(__e, symdo, tmp12883)

tmp12885 := GCons(__e, tmp12884, Nil)

tmp12886 := GCons(__e, tmp12880, tmp12885)

tmp12887 := GCons(__e, symdo, tmp12886)

tmp12888 := GCons(__e, tmp12887, Nil)

tmp12889 := GCons(__e, tmp12872, tmp12888)

tmp12890 := GCons(__e, symdo, tmp12889)

tmp12891 := GCons(__e, tmp12890, Nil)

tmp12892 := GCons(__e, V5883, tmp12891)

tmp12893 := GCons(__e, symResult, tmp12892)

tmp12894 := GCons(__e, symlet, tmp12893)

tmp12895 := GCons(__e, tmp12894, Nil)

tmp12896 := GCons(__e, tmp12866, tmp12895)

tmp12897 := GCons(__e, symdo, tmp12896)

tmp12898 := GCons(__e, tmp12897, Nil)

tmp12899 := GCons(__e, tmp12865, tmp12898)

tmp12900 := GCons(__e, symdo, tmp12899)

tmp12901 := GCons(__e, tmp12900, Nil)

tmp12902 := GCons(__e, tmp12857, tmp12901)

return GCons(__e, symdo, tmp12902)


panic("unreachable")
})

tmp12903 := Call2(__e, ns2_1set, symshen_4insert_1tracking_1code, tmp12849)


_ = tmp12903

tmp12904 := MakeNative2(func(__e *ControlFlow, V5884, V5885 Obj) Obj {
tmp12907 := Call2(__e, SymFn(symoccurrences), symshen_4incinfs, V5884)


tmp12908 := GEq(__e, tmp12907, MakeInteger(0))

if True == tmp12908 {
return V5885
} else {
tmp12905 := Call1(__e, SymFn(symshen_4vector_1parameter), V5885)


return TailCall2(__e, SymFn(symshen_4vector_1dereference), V5885, tmp12905)


}


panic("unreachable")
})

tmp12909 := Call2(__e, ns2_1set, symshen_4prolog_1track, tmp12904)


_ = tmp12909

tmp12910 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V5888 := __self1
_ = V5888
tmp12935 := GEq(__e, Nil, V5888)

if True == tmp12935 {
return Nil
} else {
tmp12933 := GConsP(__e, V5888)

var ifres12915 Obj

if True == tmp12933 {
tmp12931 := GTl(__e, V5888)

tmp12932 := GConsP(__e, tmp12931)

var ifres12916 Obj

if True == tmp12932 {
tmp12928 := GTl(__e, V5888)

tmp12929 := GTl(__e, tmp12928)

tmp12930 := GConsP(__e, tmp12929)

var ifres12917 Obj

if True == tmp12930 {
tmp12924 := GTl(__e, V5888)

tmp12925 := GTl(__e, tmp12924)

tmp12926 := GTl(__e, tmp12925)

tmp12927 := GConsP(__e, tmp12926)

var ifres12918 Obj

if True == tmp12927 {
tmp12919 := GTl(__e, V5888)

tmp12920 := GTl(__e, tmp12919)

tmp12921 := GTl(__e, tmp12920)

tmp12922 := GTl(__e, tmp12921)

tmp12923 := GEq(__e, Nil, tmp12922)

ifres12918 = tmp12923


} else {
ifres12918 = False


}

ifres12917 = ifres12918


} else {
ifres12917 = False


}

ifres12916 = ifres12917


} else {
ifres12916 = False


}

ifres12915 = ifres12916


} else {
ifres12915 = False


}

if True == ifres12915 {
return GHd(__e, V5888)
} else {
tmp12913 := GConsP(__e, V5888)

if True == tmp12913 {
tmp12911 := GTl(__e, V5888)

if SymFn(symshen_4vector_1parameter) == __self {
__self1 = tmp12911
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4vector_1parameter), tmp12911)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4vector_1parameter)
}


}


}


panic("unreachable")
})
return __self
})()

tmp12936 := Call2(__e, ns2_1set, symshen_4vector_1parameter, tmp12910)


_ = tmp12936

tmp12937 := MakeNative2(func(__e *ControlFlow, V5891, V5892 Obj) Obj {
tmp12967 := GEq(__e, Nil, V5892)

if True == tmp12967 {
return V5891
} else {
tmp12965 := GConsP(__e, V5891)

var ifres12947 Obj

if True == tmp12965 {
tmp12963 := GTl(__e, V5891)

tmp12964 := GConsP(__e, tmp12963)

var ifres12948 Obj

if True == tmp12964 {
tmp12960 := GTl(__e, V5891)

tmp12961 := GTl(__e, tmp12960)

tmp12962 := GConsP(__e, tmp12961)

var ifres12949 Obj

if True == tmp12962 {
tmp12956 := GTl(__e, V5891)

tmp12957 := GTl(__e, tmp12956)

tmp12958 := GTl(__e, tmp12957)

tmp12959 := GConsP(__e, tmp12958)

var ifres12950 Obj

if True == tmp12959 {
tmp12951 := GTl(__e, V5891)

tmp12952 := GTl(__e, tmp12951)

tmp12953 := GTl(__e, tmp12952)

tmp12954 := GTl(__e, tmp12953)

tmp12955 := GEq(__e, Nil, tmp12954)

ifres12950 = tmp12955


} else {
ifres12950 = False


}

ifres12949 = ifres12950


} else {
ifres12949 = False


}

ifres12948 = ifres12949


} else {
ifres12948 = False


}

ifres12947 = ifres12948


} else {
ifres12947 = False


}

if True == ifres12947 {
return V5891
} else {
tmp12945 := GConsP(__e, V5891)

if True == tmp12945 {
tmp12938 := GHd(__e, V5891)

tmp12939 := GCons(__e, V5892, Nil)

tmp12940 := GCons(__e, tmp12938, tmp12939)

tmp12941 := GCons(__e, symshen_4deref, tmp12940)

tmp12942 := GTl(__e, V5891)

tmp12943 := Call2(__e, SymFn(symshen_4vector_1dereference), tmp12942, V5892)


return GCons(__e, tmp12941, tmp12943)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4vector_1dereference)
}


}


}


panic("unreachable")
})

tmp12968 := Call2(__e, ns2_1set, symshen_4vector_1dereference, tmp12937)


_ = tmp12968

tmp12969 := MakeNative1(func(__e *ControlFlow, V5895 Obj) Obj {
tmp12973 := GEq(__e, sym_7, V5895)

if True == tmp12973 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dstep_d, True)
} else {
tmp12971 := GEq(__e, sym_1, V5895)

if True == tmp12971 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dstep_d, False)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("step expects a + or a -.\n"))
}


}


panic("unreachable")
})

tmp12974 := Call2(__e, ns2_1set, symstep, tmp12969)


_ = tmp12974

tmp12975 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dstep_d)
panic("unreachable")
})

tmp12976 := Call2(__e, ns2_1set, symstep_2, tmp12975)


_ = tmp12976

tmp12977 := MakeNative1(func(__e *ControlFlow, V5898 Obj) Obj {
tmp12981 := GEq(__e, sym_7, V5898)

if True == tmp12981 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dspy_d, True)
} else {
tmp12979 := GEq(__e, sym_1, V5898)

if True == tmp12979 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dspy_d, False)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("spy expects a + or a -.\n"))
}


}


panic("unreachable")
})

tmp12982 := Call2(__e, ns2_1set, symspy, tmp12977)


_ = tmp12982

tmp12983 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dspy_d)
panic("unreachable")
})

tmp12984 := Call2(__e, ns2_1set, symspy_2, tmp12983)


_ = tmp12984

tmp12985 := MakeNative0(func(__e *ControlFlow) Obj {
tmp12989 := GValue(__e, symshen_4_dstep_d)

if True == tmp12989 {
tmp12986 := GValue(__e, sym_dstinput_d)

tmp12987 := PrimGuard1(__e, symread_1byte, PrimReadByte, tmp12986)

return TailCall1(__e, SymFn(symshen_4check_1byte), tmp12987)


} else {
return TailCall1(__e, SymFn(symnl), MakeInteger(1))
}


panic("unreachable")
})

tmp12990 := Call2(__e, ns2_1set, symshen_4terpri_1or_1read_1char, tmp12985)


_ = tmp12990

tmp12991 := MakeNative1(func(__e *ControlFlow, V5901 Obj) Obj {
tmp12993 := GEq(__e, MakeInteger(94), V5901)

if True == tmp12993 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("aborted"))
} else {
return True
}


panic("unreachable")
})

tmp12994 := Call2(__e, ns2_1set, symshen_4check_1byte, tmp12991)


_ = tmp12994

tmp12995 := MakeNative3(func(__e *ControlFlow, V5902, V5903, V5904 Obj) Obj {
tmp12996 := Call1(__e, SymFn(symshen_4spaces), V5902)


tmp12997 := Call1(__e, SymFn(symshen_4spaces), V5902)


tmp12998 := Call3(__e, SymFn(symshen_4app), tmp12997, MakeString(""), symshen_4a)


tmp13000 := Call3(__e, SymFn(symshen_4app), V5903, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" \n"), tmp12998), symshen_4a)


tmp13002 := Call3(__e, SymFn(symshen_4app), V5902, PrimGuard2(__e, symcn, PrimStringConcat, MakeString("> Inputs to "), tmp13000), symshen_4a)


tmp13004 := Call3(__e, SymFn(symshen_4app), tmp12996, PrimGuard2(__e, symcn, PrimStringConcat, MakeString("<"), tmp13002), symshen_4a)


tmp13005 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("\n"), tmp13004)

tmp13006 := Call0(__e, SymFn(symstoutput))


tmp13007 := Call2(__e, SymFn(sympr), tmp13005, tmp13006)


_ = tmp13007

return TailCall1(__e, SymFn(symshen_4recursively_1print), V5904)


panic("unreachable")
})

tmp13008 := Call2(__e, ns2_1set, symshen_4input_1track, tmp12995)


_ = tmp13008

tmp13009 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V5907 := __self1
_ = V5907
tmp13019 := GEq(__e, Nil, V5907)

if True == tmp13019 {
tmp13010 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), MakeString(" ==>"), tmp13010)


} else {
tmp13017 := GConsP(__e, V5907)

if True == tmp13017 {
tmp13011 := GHd(__e, V5907)

tmp13012 := Call1(__e, SymFn(symprint), tmp13011)


_ = tmp13012

tmp13013 := Call0(__e, SymFn(symstoutput))


tmp13014 := Call2(__e, SymFn(sympr), MakeString(", "), tmp13013)


_ = tmp13014

tmp13015 := GTl(__e, V5907)

if SymFn(symshen_4recursively_1print) == __self {
__self1 = tmp13015
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4recursively_1print), tmp13015)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.recursively-print"))
}


}


panic("unreachable")
})
return __self
})()

tmp13020 := Call2(__e, ns2_1set, symshen_4recursively_1print, tmp13009)


_ = tmp13020

tmp13021 := MakeNative1(func(__e *ControlFlow, V5908 Obj) Obj {
tmp13025 := GEq(__e, MakeInteger(0), V5908)

if True == tmp13025 {
return MakeString("")
} else {
tmp13023 := Call1(__e, SymFn(symshen_4spaces), GSub(__e, V5908, MakeInteger(1)))


return PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" "), tmp13023)


}


panic("unreachable")
})

tmp13026 := Call2(__e, ns2_1set, symshen_4spaces, tmp13021)


_ = tmp13026

tmp13027 := MakeNative3(func(__e *ControlFlow, V5909, V5910, V5911 Obj) Obj {
tmp13028 := Call1(__e, SymFn(symshen_4spaces), V5909)


tmp13029 := Call1(__e, SymFn(symshen_4spaces), V5909)


tmp13030 := Call3(__e, SymFn(symshen_4app), V5911, MakeString(""), symshen_4s)


tmp13032 := Call3(__e, SymFn(symshen_4app), tmp13029, PrimGuard2(__e, symcn, PrimStringConcat, MakeString("==> "), tmp13030), symshen_4a)


tmp13034 := Call3(__e, SymFn(symshen_4app), V5910, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" \n"), tmp13032), symshen_4a)


tmp13036 := Call3(__e, SymFn(symshen_4app), V5909, PrimGuard2(__e, symcn, PrimStringConcat, MakeString("> Output from "), tmp13034), symshen_4a)


tmp13038 := Call3(__e, SymFn(symshen_4app), tmp13028, PrimGuard2(__e, symcn, PrimStringConcat, MakeString("<"), tmp13036), symshen_4a)


tmp13039 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("\n"), tmp13038)

tmp13040 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp13039, tmp13040)


panic("unreachable")
})

tmp13041 := Call2(__e, ns2_1set, symshen_4output_1track, tmp13027)


_ = tmp13041

tmp13042 := MakeNative1(func(__e *ControlFlow, V5912 Obj) Obj {
tmp13043 := GValue(__e, symshen_4_dtracking_d)

tmp13044 := Call2(__e, SymFn(symremove), V5912, tmp13043)


tmp13045 := PrimGuard2(__e, symset, PrimSet, symshen_4_dtracking_d, tmp13044)

_ = tmp13045

tmp13046 := MakeNative0(func(__e *ControlFlow) Obj {
tmp13047 := Call1(__e, SymFn(symps), V5912)


return TailCall1(__e, SymFn(symeval), tmp13047)


panic("unreachable")
})

tmp13048 := MakeNative1(func(__e *ControlFlow, Z5913 Obj) Obj {
return V5912
panic("unreachable")
})

tmp13049 := Call2(__e, try_1catch, tmp13046, tmp13048)


_ = tmp13049

return V5912


panic("unreachable")
})

tmp13050 := Call2(__e, ns2_1set, symuntrack, tmp13042)


_ = tmp13050

tmp13051 := MakeNative2(func(__e *ControlFlow, V5914, V5915 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4remove_1h), V5914, V5915, Nil)
panic("unreachable")
})

tmp13052 := Call2(__e, ns2_1set, symremove, tmp13051)


_ = tmp13052

tmp13053 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V5925 := __self1
_ = V5925
V5926 := __self2
_ = V5926
V5927 := __self3
_ = V5927
tmp13067 := GEq(__e, Nil, V5926)

if True == tmp13067 {
return TailCall1(__e, SymFn(symreverse), V5927)
} else {
tmp13065 := GConsP(__e, V5926)

var ifres13062 Obj

if True == tmp13065 {
tmp13063 := GHd(__e, V5926)

tmp13064 := GEq(__e, V5925, tmp13063)

ifres13062 = tmp13064


} else {
ifres13062 = False


}

if True == ifres13062 {
tmp13054 := GHd(__e, V5926)

tmp13055 := GTl(__e, V5926)

if SymFn(symshen_4remove_1h) == __self {
__self1, __self2, __self3 = tmp13054, tmp13055, V5927
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4remove_1h), tmp13054, tmp13055, V5927)


} else {
tmp13060 := GConsP(__e, V5926)

if True == tmp13060 {
tmp13056 := GTl(__e, V5926)

tmp13057 := GHd(__e, V5926)

tmp13058 := GCons(__e, tmp13057, V5927)

if SymFn(symshen_4remove_1h) == __self {
__self1, __self2, __self3 = V5925, tmp13056, tmp13058
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4remove_1h), V5925, tmp13056, tmp13058)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.remove-h"))
}


}


}


panic("unreachable")
})
return __self
})()

tmp13068 := Call2(__e, ns2_1set, symshen_4remove_1h, tmp13053)


_ = tmp13068

tmp13069 := MakeNative1(func(__e *ControlFlow, V5928 Obj) Obj {
tmp13070 := GValue(__e, symshen_4_dprofiled_d)

tmp13071 := GCons(__e, V5928, tmp13070)

tmp13072 := PrimGuard2(__e, symset, PrimSet, symshen_4_dprofiled_d, tmp13071)

_ = tmp13072

tmp13073 := Call1(__e, SymFn(symps), V5928)


return TailCall1(__e, SymFn(symshen_4profile_1help), tmp13073)


panic("unreachable")
})

tmp13074 := Call2(__e, ns2_1set, symprofile, tmp13069)


_ = tmp13074

tmp13075 := MakeNative1(func(__e *ControlFlow, V5931 Obj) Obj {
tmp13135 := GConsP(__e, V5931)

var ifres13114 Obj

if True == tmp13135 {
tmp13133 := GHd(__e, V5931)

tmp13134 := GEq(__e, symdefun, tmp13133)

var ifres13115 Obj

if True == tmp13134 {
tmp13131 := GTl(__e, V5931)

tmp13132 := GConsP(__e, tmp13131)

var ifres13116 Obj

if True == tmp13132 {
tmp13128 := GTl(__e, V5931)

tmp13129 := GTl(__e, tmp13128)

tmp13130 := GConsP(__e, tmp13129)

var ifres13117 Obj

if True == tmp13130 {
tmp13124 := GTl(__e, V5931)

tmp13125 := GTl(__e, tmp13124)

tmp13126 := GTl(__e, tmp13125)

tmp13127 := GConsP(__e, tmp13126)

var ifres13118 Obj

if True == tmp13127 {
tmp13119 := GTl(__e, V5931)

tmp13120 := GTl(__e, tmp13119)

tmp13121 := GTl(__e, tmp13120)

tmp13122 := GTl(__e, tmp13121)

tmp13123 := GEq(__e, Nil, tmp13122)

ifres13118 = tmp13123


} else {
ifres13118 = False


}

ifres13117 = ifres13118


} else {
ifres13117 = False


}

ifres13116 = ifres13117


} else {
ifres13116 = False


}

ifres13115 = ifres13116


} else {
ifres13115 = False


}

ifres13114 = ifres13115


} else {
ifres13114 = False


}

if True == ifres13114 {
tmp13076 := Call1(__e, SymFn(symgensym), symshen_4f)


let__12770 := tmp13076
_ = let__12770

tmp13077 := GTl(__e, V5931)

tmp13078 := GHd(__e, tmp13077)

tmp13079 := GTl(__e, V5931)

tmp13080 := GTl(__e, tmp13079)

tmp13081 := GHd(__e, tmp13080)

tmp13082 := GTl(__e, V5931)

tmp13083 := GHd(__e, tmp13082)

tmp13084 := GTl(__e, V5931)

tmp13085 := GTl(__e, tmp13084)

tmp13086 := GHd(__e, tmp13085)

tmp13087 := GTl(__e, V5931)

tmp13088 := GTl(__e, tmp13087)

tmp13089 := GHd(__e, tmp13088)

tmp13090 := GCons(__e, let__12770, tmp13089)

tmp13091 := Call3(__e, SymFn(symshen_4profile_1func), tmp13083, tmp13086, tmp13090)


tmp13092 := GCons(__e, tmp13091, Nil)

tmp13093 := GCons(__e, tmp13081, tmp13092)

tmp13094 := GCons(__e, tmp13078, tmp13093)

tmp13095 := GCons(__e, symdefun, tmp13094)

let__12771 := tmp13095
_ = let__12771

tmp13096 := GTl(__e, V5931)

tmp13097 := GTl(__e, tmp13096)

tmp13098 := GHd(__e, tmp13097)

tmp13099 := GTl(__e, V5931)

tmp13100 := GHd(__e, tmp13099)

tmp13101 := GTl(__e, V5931)

tmp13102 := GTl(__e, tmp13101)

tmp13103 := GTl(__e, tmp13102)

tmp13104 := GHd(__e, tmp13103)

tmp13105 := Call3(__e, SymFn(symsubst), let__12770, tmp13100, tmp13104)


tmp13106 := GCons(__e, tmp13105, Nil)

tmp13107 := GCons(__e, tmp13098, tmp13106)

tmp13108 := GCons(__e, let__12770, tmp13107)

tmp13109 := GCons(__e, symdefun, tmp13108)

let__12772 := tmp13109
_ = let__12772

tmp13110 := Call1(__e, SymFn(symeval_1kl), let__12771)


let__12773 := tmp13110
_ = let__12773

tmp13111 := Call1(__e, SymFn(symeval_1kl), let__12772)


let__12774 := tmp13111
_ = let__12774

tmp13112 := GTl(__e, V5931)

return GHd(__e, tmp13112)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("Cannot profile.\n"))
}


panic("unreachable")
})

tmp13136 := Call2(__e, ns2_1set, symshen_4profile_1help, tmp13075)


_ = tmp13136

tmp13137 := MakeNative1(func(__e *ControlFlow, V5937 Obj) Obj {
tmp13138 := GValue(__e, symshen_4_dprofiled_d)

tmp13139 := Call2(__e, SymFn(symremove), V5937, tmp13138)


tmp13140 := PrimGuard2(__e, symset, PrimSet, symshen_4_dprofiled_d, tmp13139)

_ = tmp13140

tmp13141 := MakeNative0(func(__e *ControlFlow) Obj {
tmp13142 := Call1(__e, SymFn(symps), V5937)


return TailCall1(__e, SymFn(symeval), tmp13142)


panic("unreachable")
})

tmp13143 := MakeNative1(func(__e *ControlFlow, Z5938 Obj) Obj {
return V5937
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp13141, tmp13143)


panic("unreachable")
})

tmp13144 := Call2(__e, ns2_1set, symunprofile, tmp13137)


_ = tmp13144

tmp13145 := MakeNative1(func(__e *ControlFlow, V5939 Obj) Obj {
tmp13146 := GValue(__e, symshen_4_dprofiled_d)

return TailCall2(__e, SymFn(symelement_2), V5939, tmp13146)


panic("unreachable")
})

tmp13147 := Call2(__e, ns2_1set, symshen_4profiled_2, tmp13145)


_ = tmp13147

tmp13148 := MakeNative3(func(__e *ControlFlow, V5940, V5941, V5942 Obj) Obj {
tmp13149 := GCons(__e, symrun, Nil)

tmp13150 := GCons(__e, symget_1time, tmp13149)

tmp13151 := GCons(__e, symrun, Nil)

tmp13152 := GCons(__e, symget_1time, tmp13151)

tmp13153 := GCons(__e, symStart, Nil)

tmp13154 := GCons(__e, tmp13152, tmp13153)

tmp13155 := GCons(__e, sym_1, tmp13154)

tmp13156 := GCons(__e, V5940, Nil)

tmp13157 := GCons(__e, symshen_4get_1profile, tmp13156)

tmp13158 := GCons(__e, symFinish, Nil)

tmp13159 := GCons(__e, tmp13157, tmp13158)

tmp13160 := GCons(__e, sym_7, tmp13159)

tmp13161 := GCons(__e, tmp13160, Nil)

tmp13162 := GCons(__e, V5940, tmp13161)

tmp13163 := GCons(__e, symshen_4put_1profile, tmp13162)

tmp13164 := GCons(__e, symResult, Nil)

tmp13165 := GCons(__e, tmp13163, tmp13164)

tmp13166 := GCons(__e, symRecord, tmp13165)

tmp13167 := GCons(__e, symlet, tmp13166)

tmp13168 := GCons(__e, tmp13167, Nil)

tmp13169 := GCons(__e, tmp13155, tmp13168)

tmp13170 := GCons(__e, symFinish, tmp13169)

tmp13171 := GCons(__e, symlet, tmp13170)

tmp13172 := GCons(__e, tmp13171, Nil)

tmp13173 := GCons(__e, V5942, tmp13172)

tmp13174 := GCons(__e, symResult, tmp13173)

tmp13175 := GCons(__e, symlet, tmp13174)

tmp13176 := GCons(__e, tmp13175, Nil)

tmp13177 := GCons(__e, tmp13150, tmp13176)

tmp13178 := GCons(__e, symStart, tmp13177)

return GCons(__e, symlet, tmp13178)


panic("unreachable")
})

tmp13179 := Call2(__e, ns2_1set, symshen_4profile_1func, tmp13148)


_ = tmp13179

tmp13180 := MakeNative1(func(__e *ControlFlow, V5943 Obj) Obj {
tmp13181 := Call1(__e, SymFn(symshen_4get_1profile), V5943)


let__12775 := tmp13181
_ = let__12775

tmp13182 := Call2(__e, SymFn(symshen_4put_1profile), V5943, MakeInteger(0))


let__12776 := tmp13182
_ = let__12776

return TailCall2(__e, SymFn(sym_8p), V5943, let__12775)


panic("unreachable")
})

tmp13183 := Call2(__e, ns2_1set, symprofile_1results, tmp13180)


_ = tmp13183

tmp13184 := MakeNative1(func(__e *ControlFlow, V5946 Obj) Obj {
tmp13185 := MakeNative0(func(__e *ControlFlow) Obj {
tmp13186 := GValue(__e, sym_dproperty_1vector_d)

return TailCall3(__e, SymFn(symget), V5946, symprofile, tmp13186)


panic("unreachable")
})

tmp13187 := MakeNative1(func(__e *ControlFlow, Z5947 Obj) Obj {
return MakeInteger(0)
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp13185, tmp13187)


panic("unreachable")
})

tmp13188 := Call2(__e, ns2_1set, symshen_4get_1profile, tmp13184)


_ = tmp13188

tmp13189 := MakeNative2(func(__e *ControlFlow, V5948, V5949 Obj) Obj {
tmp13190 := GValue(__e, sym_dproperty_1vector_d)

return TailCall4(__e, SymFn(symput), V5948, symprofile, V5949, tmp13190)


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4put_1profile, tmp13189)




panic("unreachable")
})

