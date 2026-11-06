package main

import . "github.com/pyrex41/shen-go/kl"

var TStarMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp13742 := MakeNative2(func(__e *ControlFlow, V4915, V4916 Obj) Obj {
tmp13743 := Call1(__e, SymFn(symshen_4extract_1vars), V4916)


let__13191 := tmp13743
_ = let__13191

tmp13744 := Call1(__e, SymFn(symshen_4rectify_1type), V4916)


let__13192 := tmp13744
_ = let__13192

tmp13745 := Call1(__e, SymFn(symshen_4curry), V4915)


let__13193 := tmp13745
_ = let__13193

tmp13746 := MakeNative1(func(__e *ControlFlow, Z4920 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z4921 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z4922 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z4923 Obj) Obj {
tmp13747 := Call1(__e, SymFn(symshen_4newpv), Z4920)


let__13194 := tmp13747
_ = let__13194

tmp13748 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp13748

tmp13749 := Call2(__e, SymFn(symshen_4deref), let__13191, Z4920)


tmp13750 := Call1(__e, SymFn(symreceive), tmp13749)


tmp13751 := Call2(__e, SymFn(symshen_4deref), let__13192, Z4920)


tmp13752 := Call1(__e, SymFn(symreceive), tmp13751)


tmp13753 := MakeNative0(func(__e *ControlFlow) Obj {
tmp13754 := Call2(__e, SymFn(symshen_4deref), let__13193, Z4920)


tmp13755 := Call1(__e, SymFn(symreceive), tmp13754)


tmp13756 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall5(__e, SymFn(symreturn), let__13194, Z4920, Z4921, Z4922, Z4923)
panic("unreachable")
})

return TailCall6(__e, SymFn(symshen_4toplevel_1forms), tmp13755, let__13194, Z4920, Z4921, Z4922, tmp13756)


panic("unreachable")
})

tmp13757 := Call(__e, SymFn(symshen_4insert_1prolog_1variables), tmp13750, tmp13752, let__13194, Z4920, Z4921, Z4922, tmp13753)


return TailCall2(__e, SymFn(symshen_4gc), Z4920, tmp13757)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp13758 := Call0(__e, SymFn(symshen_4prolog_1vector))


tmp13759 := Call1(__e, tmp13746, tmp13758)


tmp13760 := Call1(__e, SymFn(symvector), MakeInteger(0))


tmp13761 := Call2(__e, SymFn(sym_8v), MakeInteger(0), tmp13760)


tmp13762 := Call2(__e, SymFn(sym_8v), True, tmp13761)


tmp13763 := Call1(__e, tmp13759, tmp13762)


tmp13764 := Call1(__e, tmp13763, MakeInteger(0))


tmp13765 := MakeNative0(func(__e *ControlFlow) Obj {
return True
panic("unreachable")
})

return TailCall1(__e, tmp13764, tmp13765)


panic("unreachable")
})

tmp13766 := Call2(__e, ns2_1set, symshen_4typecheck, tmp13742)


_ = tmp13766

tmp13767 := MakeNative(func(__e *ControlFlow) {
V4925 := __e.Get(1)
V4926 := __e.Get(2)
V4927 := __e.Get(3)
V4928 := __e.Get(4)
V4929 := __e.Get(5)
V4930 := __e.Get(6)
V4931 := __e.Get(7)
__e.ReturnDirect(func() Obj {
tmp13774 := Call1(__e, SymFn(symshen_4unlocked_2), V4929)


var ifres13768 Obj

if True == tmp13774 {
tmp13769 := Call2(__e, SymFn(symshen_4lazyderef), V4925, V4928)


let__13196 := tmp13769
_ = let__13196

tmp13773 := GEq(__e, let__13196, Nil)

var ifres13770 Obj

if True == tmp13773 {
tmp13771 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp13771

tmp13772 := Call6(__e, SymFn(symis_b), V4926, V4927, V4928, V4929, V4930, V4931)


ifres13770 = tmp13772


} else {
ifres13770 = False


}

ifres13768 = ifres13770


} else {
ifres13768 = False


}

let__13195 := ifres13768
_ = let__13195

tmp13788 := GEq(__e, let__13195, False)

if True == tmp13788 {
tmp13786 := Call1(__e, SymFn(symshen_4unlocked_2), V4929)


if True == tmp13786 {
tmp13775 := Call2(__e, SymFn(symshen_4lazyderef), V4925, V4928)


let__13197 := tmp13775
_ = let__13197

tmp13784 := GConsP(__e, let__13197)

if True == tmp13784 {
tmp13776 := GHd(__e, let__13197)

let__13198 := tmp13776
_ = let__13198

tmp13777 := GTl(__e, let__13197)

let__13199 := tmp13777
_ = let__13199

tmp13778 := Call1(__e, SymFn(symshen_4newpv), V4928)


let__13200 := tmp13778
_ = let__13200

tmp13779 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp13779

tmp13780 := Call2(__e, SymFn(symshen_4deref), let__13200, V4928)


tmp13781 := Call3(__e, SymFn(symsubst), tmp13780, let__13198, V4926)


tmp13782 := Call(__e, SymFn(symshen_4insert_1prolog_1variables), let__13199, tmp13781, V4927, V4928, V4929, V4930, V4931)


return TailCall2(__e, SymFn(symshen_4gc), V4928, tmp13782)


} else {
return False
}


} else {
return False
}


} else {
return let__13195
}


panic("unreachable")
}())
}, 7)

tmp13789 := Call2(__e, ns2_1set, symshen_4insert_1prolog_1variables, tmp13767)


_ = tmp13789

tmp13790 := MakeNative6(func(__e *ControlFlow, V4938, V4939, V4940, V4941, V4942, V4943 Obj) Obj {
let__13201 := GAdd(__e, V4942, MakeInteger(1))
_ = let__13201

tmp13815 := Call1(__e, SymFn(symshen_4unlocked_2), V4941)


var ifres13792 Obj

if True == tmp13815 {
tmp13793 := Call2(__e, SymFn(symshen_4lazyderef), V4938, V4940)


let__13203 := tmp13793
_ = let__13203

tmp13814 := GConsP(__e, let__13203)

var ifres13794 Obj

if True == tmp13814 {
tmp13795 := GHd(__e, let__13203)

tmp13796 := Call2(__e, SymFn(symshen_4lazyderef), tmp13795, V4940)


let__13204 := tmp13796
_ = let__13204

tmp13813 := GEq(__e, let__13204, symdefine)

var ifres13797 Obj

if True == tmp13813 {
tmp13798 := GTl(__e, let__13203)

tmp13799 := Call2(__e, SymFn(symshen_4lazyderef), tmp13798, V4940)


let__13205 := tmp13799
_ = let__13205

tmp13812 := GConsP(__e, let__13205)

var ifres13800 Obj

if True == tmp13812 {
tmp13801 := GHd(__e, let__13205)

let__13206 := tmp13801
_ = let__13206

tmp13802 := GTl(__e, let__13205)

let__13207 := tmp13802
_ = let__13207

tmp13803 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp13803

tmp13804 := Call0(__e, SymFn(symshen_4type_1theory_1enabled_2))


tmp13805 := MakeNative0(func(__e *ControlFlow) Obj {
tmp13806 := MakeNative0(func(__e *ControlFlow) Obj {
tmp13807 := GValue(__e, symshen_4_dspy_d)

tmp13808 := MakeNative0(func(__e *ControlFlow) Obj {
tmp13809 := GCons(__e, let__13206, let__13207)

tmp13810 := GCons(__e, symdefine, tmp13809)

return TailCall6(__e, SymFn(symshen_4t_d), tmp13810, V4939, V4940, V4941, let__13201, V4943)


panic("unreachable")
})

return TailCall6(__e, SymFn(symshen_4signal_1def), tmp13807, let__13206, V4940, V4941, let__13201, tmp13808)


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V4940, V4941, let__13201, tmp13806)


panic("unreachable")
})

tmp13811 := Call5(__e, SymFn(symwhen), tmp13804, V4940, V4941, let__13201, tmp13805)


ifres13800 = tmp13811


} else {
ifres13800 = False


}

ifres13797 = ifres13800


} else {
ifres13797 = False


}

ifres13794 = ifres13797


} else {
ifres13794 = False


}

ifres13792 = ifres13794


} else {
ifres13792 = False


}

let__13202 := ifres13792
_ = let__13202

tmp13827 := GEq(__e, let__13202, False)

if True == tmp13827 {
tmp13823 := Call1(__e, SymFn(symshen_4unlocked_2), V4941)


var ifres13816 Obj

if True == tmp13823 {
tmp13817 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp13817

tmp13818 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp13819 := GCons(__e, V4939, Nil)

tmp13820 := GCons(__e, tmp13818, tmp13819)

tmp13821 := GCons(__e, V4938, tmp13820)

tmp13822 := Call6(__e, SymFn(symshen_4system_1S), tmp13821, Nil, V4940, V4941, let__13201, V4943)


ifres13816 = tmp13822


} else {
ifres13816 = False


}

let__13208 := ifres13816
_ = let__13208

tmp13825 := GEq(__e, let__13208, False)

if True == tmp13825 {
return TailCall2(__e, SymFn(symshen_4unlock), V4941, let__13201)
} else {
return let__13208
}


} else {
return let__13202
}


panic("unreachable")
})

tmp13828 := Call2(__e, ns2_1set, symshen_4toplevel_1forms, tmp13790)


_ = tmp13828

tmp13829 := MakeNative6(func(__e *ControlFlow, V4952, V4953, V4954, V4955, V4956, V4957 Obj) Obj {
tmp13836 := Call1(__e, SymFn(symshen_4unlocked_2), V4955)


var ifres13830 Obj

if True == tmp13836 {
tmp13831 := Call2(__e, SymFn(symshen_4lazyderef), V4952, V4954)


let__13210 := tmp13831
_ = let__13210

tmp13835 := GEq(__e, let__13210, False)

var ifres13832 Obj

if True == tmp13835 {
tmp13833 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp13833

tmp13834 := Call1(__e, SymFn(symthaw), V4957)


ifres13832 = tmp13834


} else {
ifres13832 = False


}

ifres13830 = ifres13832


} else {
ifres13830 = False


}

let__13209 := ifres13830
_ = let__13209

tmp13851 := GEq(__e, let__13209, False)

if True == tmp13851 {
tmp13849 := Call1(__e, SymFn(symshen_4unlocked_2), V4955)


if True == tmp13849 {
tmp13837 := Call2(__e, SymFn(symshen_4lazyderef), V4952, V4954)


let__13211 := tmp13837
_ = let__13211

tmp13847 := GEq(__e, let__13211, True)

if True == tmp13847 {
tmp13838 := Call1(__e, SymFn(symshen_4newpv), V4954)


let__13212 := tmp13838
_ = let__13212

tmp13839 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp13839

tmp13840 := Call2(__e, SymFn(symshen_4deref), V4953, V4954)


tmp13841 := Call3(__e, SymFn(symshen_4app), tmp13840, MakeString(")\n"), symshen_4a)


tmp13842 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("\ntypechecking (fn "), tmp13841)

tmp13843 := Call0(__e, SymFn(symstoutput))


tmp13844 := Call2(__e, SymFn(sympr), tmp13842, tmp13843)


tmp13845 := Call6(__e, SymFn(symis), let__13212, tmp13844, V4954, V4955, V4956, V4957)


return TailCall2(__e, SymFn(symshen_4gc), V4954, tmp13845)


} else {
return False
}


} else {
return False
}


} else {
return let__13209
}


panic("unreachable")
})

tmp13852 := Call2(__e, ns2_1set, symshen_4signal_1def, tmp13829)


_ = tmp13852

tmp13853 := MakeNative1(func(__e *ControlFlow, V4962 Obj) Obj {
tmp13854 := Call1(__e, SymFn(symshen_4curry_1type), V4962)


return TailCall1(__e, SymFn(symshen_4demodulate), tmp13854)


panic("unreachable")
})

tmp13855 := Call2(__e, ns2_1set, symshen_4rectify_1type, tmp13853)


_ = tmp13855

tmp13856 := MakeNative1(func(__e *ControlFlow, V4963 Obj) Obj {
tmp13857 := MakeNative0(func(__e *ControlFlow) Obj {
tmp13858 := MakeNative1(func(__e *ControlFlow, Z4965 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4demod), Z4965)
panic("unreachable")
})

tmp13859 := Call2(__e, SymFn(symshen_4walk), tmp13858, V4963)


let__13213 := tmp13859
_ = let__13213

tmp13861 := GEq(__e, let__13213, V4963)

if True == tmp13861 {
return V4963
} else {
return TailCall1(__e, SymFn(symshen_4demodulate), let__13213)
}


panic("unreachable")
})

tmp13862 := MakeNative1(func(__e *ControlFlow, Z4966 Obj) Obj {
return V4963
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp13857, tmp13862)


panic("unreachable")
})

tmp13863 := Call2(__e, ns2_1set, symshen_4demodulate, tmp13856)


_ = tmp13863

tmp13864 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V4967 := __self1
_ = V4967
tmp13970 := GConsP(__e, V4967)

var ifres13948 Obj

if True == tmp13970 {
tmp13968 := GTl(__e, V4967)

tmp13969 := GConsP(__e, tmp13968)

var ifres13949 Obj

if True == tmp13969 {
tmp13965 := GTl(__e, V4967)

tmp13966 := GHd(__e, tmp13965)

tmp13967 := GEq(__e, sym_1_1_6, tmp13966)

var ifres13950 Obj

if True == tmp13967 {
tmp13962 := GTl(__e, V4967)

tmp13963 := GTl(__e, tmp13962)

tmp13964 := GConsP(__e, tmp13963)

var ifres13951 Obj

if True == tmp13964 {
tmp13958 := GTl(__e, V4967)

tmp13959 := GTl(__e, tmp13958)

tmp13960 := GTl(__e, tmp13959)

tmp13961 := GConsP(__e, tmp13960)

var ifres13952 Obj

if True == tmp13961 {
tmp13953 := GTl(__e, V4967)

tmp13954 := GTl(__e, tmp13953)

tmp13955 := GTl(__e, tmp13954)

tmp13956 := GHd(__e, tmp13955)

tmp13957 := GEq(__e, sym_1_1_6, tmp13956)

ifres13952 = tmp13957


} else {
ifres13952 = False


}

ifres13951 = ifres13952


} else {
ifres13951 = False


}

ifres13950 = ifres13951


} else {
ifres13950 = False


}

ifres13949 = ifres13950


} else {
ifres13949 = False


}

ifres13948 = ifres13949


} else {
ifres13948 = False


}

if True == ifres13948 {
tmp13865 := GHd(__e, V4967)

tmp13866 := GTl(__e, V4967)

tmp13867 := GTl(__e, tmp13866)

tmp13868 := GCons(__e, tmp13867, Nil)

tmp13869 := GCons(__e, sym_1_1_6, tmp13868)

tmp13870 := GCons(__e, tmp13865, tmp13869)

if SymFn(symshen_4curry_1type) == __self {
__self1 = tmp13870
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4curry_1type), tmp13870)


} else {
tmp13946 := GConsP(__e, V4967)

var ifres13914 Obj

if True == tmp13946 {
tmp13944 := GHd(__e, V4967)

tmp13945 := GConsP(__e, tmp13944)

var ifres13915 Obj

if True == tmp13945 {
tmp13941 := GHd(__e, V4967)

tmp13942 := GHd(__e, tmp13941)

tmp13943 := GEq(__e, symlist, tmp13942)

var ifres13916 Obj

if True == tmp13943 {
tmp13938 := GHd(__e, V4967)

tmp13939 := GTl(__e, tmp13938)

tmp13940 := GConsP(__e, tmp13939)

var ifres13917 Obj

if True == tmp13940 {
tmp13934 := GHd(__e, V4967)

tmp13935 := GTl(__e, tmp13934)

tmp13936 := GTl(__e, tmp13935)

tmp13937 := GEq(__e, Nil, tmp13936)

var ifres13918 Obj

if True == tmp13937 {
tmp13932 := GTl(__e, V4967)

tmp13933 := GConsP(__e, tmp13932)

var ifres13919 Obj

if True == tmp13933 {
tmp13929 := GTl(__e, V4967)

tmp13930 := GHd(__e, tmp13929)

tmp13931 := GEq(__e, sym_a_a_6, tmp13930)

var ifres13920 Obj

if True == tmp13931 {
tmp13926 := GTl(__e, V4967)

tmp13927 := GTl(__e, tmp13926)

tmp13928 := GConsP(__e, tmp13927)

var ifres13921 Obj

if True == tmp13928 {
tmp13922 := GTl(__e, V4967)

tmp13923 := GTl(__e, tmp13922)

tmp13924 := GTl(__e, tmp13923)

tmp13925 := GEq(__e, Nil, tmp13924)

ifres13921 = tmp13925


} else {
ifres13921 = False


}

ifres13920 = ifres13921


} else {
ifres13920 = False


}

ifres13919 = ifres13920


} else {
ifres13919 = False


}

ifres13918 = ifres13919


} else {
ifres13918 = False


}

ifres13917 = ifres13918


} else {
ifres13917 = False


}

ifres13916 = ifres13917


} else {
ifres13916 = False


}

ifres13915 = ifres13916


} else {
ifres13915 = False


}

ifres13914 = ifres13915


} else {
ifres13914 = False


}

if True == ifres13914 {
tmp13871 := GHd(__e, V4967)

tmp13872 := GHd(__e, V4967)

tmp13873 := GTl(__e, V4967)

tmp13874 := GTl(__e, tmp13873)

tmp13875 := GCons(__e, tmp13872, tmp13874)

tmp13876 := GCons(__e, symstr, tmp13875)

tmp13877 := GCons(__e, tmp13876, Nil)

tmp13878 := GCons(__e, sym_1_1_6, tmp13877)

tmp13879 := GCons(__e, tmp13871, tmp13878)

if SymFn(symshen_4curry_1type) == __self {
__self1 = tmp13879
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4curry_1type), tmp13879)


} else {
tmp13912 := GConsP(__e, V4967)

var ifres13890 Obj

if True == tmp13912 {
tmp13910 := GTl(__e, V4967)

tmp13911 := GConsP(__e, tmp13910)

var ifres13891 Obj

if True == tmp13911 {
tmp13907 := GTl(__e, V4967)

tmp13908 := GHd(__e, tmp13907)

tmp13909 := GEq(__e, sym_d, tmp13908)

var ifres13892 Obj

if True == tmp13909 {
tmp13904 := GTl(__e, V4967)

tmp13905 := GTl(__e, tmp13904)

tmp13906 := GConsP(__e, tmp13905)

var ifres13893 Obj

if True == tmp13906 {
tmp13900 := GTl(__e, V4967)

tmp13901 := GTl(__e, tmp13900)

tmp13902 := GTl(__e, tmp13901)

tmp13903 := GConsP(__e, tmp13902)

var ifres13894 Obj

if True == tmp13903 {
tmp13895 := GTl(__e, V4967)

tmp13896 := GTl(__e, tmp13895)

tmp13897 := GTl(__e, tmp13896)

tmp13898 := GHd(__e, tmp13897)

tmp13899 := GEq(__e, sym_d, tmp13898)

ifres13894 = tmp13899


} else {
ifres13894 = False


}

ifres13893 = ifres13894


} else {
ifres13893 = False


}

ifres13892 = ifres13893


} else {
ifres13892 = False


}

ifres13891 = ifres13892


} else {
ifres13891 = False


}

ifres13890 = ifres13891


} else {
ifres13890 = False


}

if True == ifres13890 {
tmp13880 := GHd(__e, V4967)

tmp13881 := GTl(__e, V4967)

tmp13882 := GTl(__e, tmp13881)

tmp13883 := GCons(__e, tmp13882, Nil)

tmp13884 := GCons(__e, sym_d, tmp13883)

tmp13885 := GCons(__e, tmp13880, tmp13884)

if SymFn(symshen_4curry_1type) == __self {
__self1 = tmp13885
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4curry_1type), tmp13885)


} else {
tmp13888 := GConsP(__e, V4967)

if True == tmp13888 {
tmp13886 := MakeNative1(func(__e *ControlFlow, Z4968 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4curry_1type), Z4968)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp13886, V4967)


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

tmp13971 := Call2(__e, ns2_1set, symshen_4curry_1type, tmp13864)


_ = tmp13971

tmp13972 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V4969 := __self1
_ = V4969
tmp14049 := GConsP(__e, V4969)

var ifres14043 Obj

if True == tmp14049 {
tmp14047 := GHd(__e, V4969)

tmp14048 := GEq(__e, symdefine, tmp14047)

var ifres14044 Obj

if True == tmp14048 {
tmp14045 := GTl(__e, V4969)

tmp14046 := GConsP(__e, tmp14045)

ifres14044 = tmp14046


} else {
ifres14044 = False


}

ifres14043 = ifres14044


} else {
ifres14043 = False


}

if True == ifres14043 {
return V4969
} else {
tmp14041 := GConsP(__e, V4969)

var ifres14026 Obj

if True == tmp14041 {
tmp14039 := GHd(__e, V4969)

tmp14040 := GEq(__e, symtype, tmp14039)

var ifres14027 Obj

if True == tmp14040 {
tmp14037 := GTl(__e, V4969)

tmp14038 := GConsP(__e, tmp14037)

var ifres14028 Obj

if True == tmp14038 {
tmp14034 := GTl(__e, V4969)

tmp14035 := GTl(__e, tmp14034)

tmp14036 := GConsP(__e, tmp14035)

var ifres14029 Obj

if True == tmp14036 {
tmp14030 := GTl(__e, V4969)

tmp14031 := GTl(__e, tmp14030)

tmp14032 := GTl(__e, tmp14031)

tmp14033 := GEq(__e, Nil, tmp14032)

ifres14029 = tmp14033


} else {
ifres14029 = False


}

ifres14028 = ifres14029


} else {
ifres14028 = False


}

ifres14027 = ifres14028


} else {
ifres14027 = False


}

ifres14026 = ifres14027


} else {
ifres14026 = False


}

if True == ifres14026 {
tmp13973 := GTl(__e, V4969)

tmp13974 := GHd(__e, tmp13973)

tmp13975 := Call1(__e, SymFn(symshen_4curry), tmp13974)


tmp13976 := GTl(__e, V4969)

tmp13977 := GTl(__e, tmp13976)

tmp13978 := GCons(__e, tmp13975, tmp13977)

return GCons(__e, symtype, tmp13978)


} else {
tmp14024 := GConsP(__e, V4969)

var ifres14021 Obj

if True == tmp14024 {
tmp14022 := GHd(__e, V4969)

tmp14023 := Call1(__e, SymFn(symshen_4special_2), tmp14022)


ifres14021 = tmp14023


} else {
ifres14021 = False


}

if True == ifres14021 {
tmp13979 := GHd(__e, V4969)

tmp13980 := MakeNative1(func(__e *ControlFlow, Z4970 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4curry), Z4970)
panic("unreachable")
})

tmp13981 := GTl(__e, V4969)

tmp13982 := Call2(__e, SymFn(symmap), tmp13980, tmp13981)


return GCons(__e, tmp13979, tmp13982)


} else {
tmp14019 := GConsP(__e, V4969)

var ifres14016 Obj

if True == tmp14019 {
tmp14017 := GHd(__e, V4969)

tmp14018 := Call1(__e, SymFn(symshen_4extraspecial_2), tmp14017)


ifres14016 = tmp14018


} else {
ifres14016 = False


}

if True == ifres14016 {
return V4969
} else {
tmp14014 := GConsP(__e, V4969)

var ifres14007 Obj

if True == tmp14014 {
tmp14012 := GTl(__e, V4969)

tmp14013 := GConsP(__e, tmp14012)

var ifres14008 Obj

if True == tmp14013 {
tmp14009 := GTl(__e, V4969)

tmp14010 := GTl(__e, tmp14009)

tmp14011 := GConsP(__e, tmp14010)

ifres14008 = tmp14011


} else {
ifres14008 = False


}

ifres14007 = ifres14008


} else {
ifres14007 = False


}

if True == ifres14007 {
tmp13983 := GHd(__e, V4969)

tmp13984 := GTl(__e, V4969)

tmp13985 := GHd(__e, tmp13984)

tmp13986 := GCons(__e, tmp13985, Nil)

tmp13987 := GCons(__e, tmp13983, tmp13986)

tmp13988 := GTl(__e, V4969)

tmp13989 := GTl(__e, tmp13988)

tmp13990 := GCons(__e, tmp13987, tmp13989)

if SymFn(symshen_4curry) == __self {
__self1 = tmp13990
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4curry), tmp13990)


} else {
tmp14005 := GConsP(__e, V4969)

var ifres13998 Obj

if True == tmp14005 {
tmp14003 := GTl(__e, V4969)

tmp14004 := GConsP(__e, tmp14003)

var ifres13999 Obj

if True == tmp14004 {
tmp14000 := GTl(__e, V4969)

tmp14001 := GTl(__e, tmp14000)

tmp14002 := GEq(__e, Nil, tmp14001)

ifres13999 = tmp14002


} else {
ifres13999 = False


}

ifres13998 = ifres13999


} else {
ifres13998 = False


}

if True == ifres13998 {
tmp13991 := GHd(__e, V4969)

tmp13992 := Call1(__e, SymFn(symshen_4curry), tmp13991)


tmp13993 := GTl(__e, V4969)

tmp13994 := GHd(__e, tmp13993)

tmp13995 := Call1(__e, SymFn(symshen_4curry), tmp13994)


tmp13996 := GCons(__e, tmp13995, Nil)

return GCons(__e, tmp13992, tmp13996)


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

tmp14050 := Call2(__e, ns2_1set, symshen_4curry, tmp13972)


_ = tmp14050

tmp14051 := MakeNative1(func(__e *ControlFlow, V4971 Obj) Obj {
tmp14052 := GValue(__e, symshen_4_dspecial_d)

return TailCall2(__e, SymFn(symelement_2), V4971, tmp14052)


panic("unreachable")
})

tmp14053 := Call2(__e, ns2_1set, symshen_4special_2, tmp14051)


_ = tmp14053

tmp14054 := MakeNative1(func(__e *ControlFlow, V4972 Obj) Obj {
tmp14055 := GValue(__e, symshen_4_dextraspecial_d)

return TailCall2(__e, SymFn(symelement_2), V4972, tmp14055)


panic("unreachable")
})

tmp14056 := Call2(__e, ns2_1set, symshen_4extraspecial_2, tmp14054)


_ = tmp14056

tmp14057 := MakeNative6(func(__e *ControlFlow, V4973, V4974, V4975, V4976, V4977, V4978 Obj) Obj {
let__13214 := GAdd(__e, V4977, MakeInteger(1))
_ = let__13214

tmp14063 := Call1(__e, SymFn(symshen_4unlocked_2), V4976)


var ifres14059 Obj

if True == tmp14063 {
tmp14060 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14060

tmp14061 := Call0(__e, SymFn(symshen_4maxinfexceeded_2))


tmp14062 := Call5(__e, SymFn(symwhen), tmp14061, V4975, V4976, let__13214, V4978)


ifres14059 = tmp14062


} else {
ifres14059 = False


}

let__13215 := ifres14059
_ = let__13215

tmp14111 := GEq(__e, let__13215, False)

if True == tmp14111 {
tmp14092 := Call1(__e, SymFn(symshen_4unlocked_2), V4976)


var ifres14064 Obj

if True == tmp14092 {
tmp14065 := Call2(__e, SymFn(symshen_4lazyderef), V4973, V4975)


let__13217 := tmp14065
_ = let__13217

tmp14091 := GConsP(__e, let__13217)

var ifres14066 Obj

if True == tmp14091 {
tmp14067 := GHd(__e, let__13217)

let__13218 := tmp14067
_ = let__13218

tmp14068 := GTl(__e, let__13217)

tmp14069 := Call2(__e, SymFn(symshen_4lazyderef), tmp14068, V4975)


let__13219 := tmp14069
_ = let__13219

tmp14090 := GConsP(__e, let__13219)

var ifres14070 Obj

if True == tmp14090 {
tmp14071 := GHd(__e, let__13219)

let__13220 := tmp14071
_ = let__13220

tmp14072 := GTl(__e, let__13219)

tmp14073 := Call2(__e, SymFn(symshen_4lazyderef), tmp14072, V4975)


let__13221 := tmp14073
_ = let__13221

tmp14089 := GConsP(__e, let__13221)

var ifres14074 Obj

if True == tmp14089 {
tmp14075 := GHd(__e, let__13221)

let__13222 := tmp14075
_ = let__13222

tmp14076 := GTl(__e, let__13221)

tmp14077 := Call2(__e, SymFn(symshen_4lazyderef), tmp14076, V4975)


let__13223 := tmp14077
_ = let__13223

tmp14088 := GEq(__e, let__13223, Nil)

var ifres14078 Obj

if True == tmp14088 {
tmp14079 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14079

tmp14080 := Call2(__e, SymFn(symshen_4deref), let__13220, V4975)


tmp14081 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp14082 := GEq(__e, tmp14080, tmp14081)

tmp14083 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14084 := Call0(__e, SymFn(symshen_4type_1theory_1enabled_2))


tmp14085 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14086 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), let__13218, let__13222, V4974, V4975, V4976, let__13214, V4978)
return nil
panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V4975, V4976, let__13214, tmp14086)


panic("unreachable")
})

return TailCall5(__e, SymFn(symwhen), tmp14084, V4975, V4976, let__13214, tmp14085)


panic("unreachable")
})

tmp14087 := Call5(__e, SymFn(symwhen), tmp14082, V4975, V4976, let__13214, tmp14083)


ifres14078 = tmp14087


} else {
ifres14078 = False


}

ifres14074 = ifres14078


} else {
ifres14074 = False


}

ifres14070 = ifres14074


} else {
ifres14070 = False


}

ifres14066 = ifres14070


} else {
ifres14066 = False


}

ifres14064 = ifres14066


} else {
ifres14064 = False


}

let__13216 := ifres14064
_ = let__13216

tmp14109 := GEq(__e, let__13216, False)

if True == tmp14109 {
tmp14098 := Call1(__e, SymFn(symshen_4unlocked_2), V4976)


var ifres14093 Obj

if True == tmp14098 {
tmp14094 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14094

tmp14095 := GValue(__e, symshen_4_dspy_d)

tmp14096 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall6(__e, SymFn(symshen_4show), V4973, V4974, V4975, V4976, let__13214, V4978)
panic("unreachable")
})

tmp14097 := Call5(__e, SymFn(symwhen), tmp14095, V4975, V4976, let__13214, tmp14096)


ifres14093 = tmp14097


} else {
ifres14093 = False


}

let__13224 := ifres14093
_ = let__13224

tmp14107 := GEq(__e, let__13224, False)

if True == tmp14107 {
tmp14103 := Call1(__e, SymFn(symshen_4unlocked_2), V4976)


var ifres14099 Obj

if True == tmp14103 {
tmp14100 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14100

tmp14101 := GValue(__e, symshen_4_ddatatypes_d)

tmp14102 := Call(__e, SymFn(symshen_4search_1user_1datatypes), V4973, V4974, tmp14101, V4975, V4976, let__13214, V4978)


ifres14099 = tmp14102


} else {
ifres14099 = False


}

let__13225 := ifres14099
_ = let__13225

tmp14105 := GEq(__e, let__13225, False)

if True == tmp14105 {
return TailCall2(__e, SymFn(symshen_4unlock), V4976, let__13214)
} else {
return let__13225
}


} else {
return let__13224
}


} else {
return let__13216
}


} else {
return let__13215
}


panic("unreachable")
})

tmp14112 := Call2(__e, ns2_1set, symshen_4system_1S, tmp14057)


_ = tmp14112

tmp14113 := MakeNative6(func(__e *ControlFlow, V4997, V4998, V4999, V5000, V5001, V5002 Obj) Obj {
tmp14114 := Call0(__e, SymFn(symshen_4line))


_ = tmp14114

tmp14115 := Call2(__e, SymFn(symshen_4deref), V4997, V4999)


tmp14116 := Call1(__e, SymFn(symshen_4show_1p), tmp14115)


_ = tmp14116

tmp14117 := Call1(__e, SymFn(symnl), MakeInteger(2))


_ = tmp14117

tmp14118 := Call2(__e, SymFn(symshen_4deref), V4998, V4999)


tmp14119 := Call2(__e, SymFn(symshen_4show_1assumptions), tmp14118, MakeInteger(1))


_ = tmp14119

tmp14120 := Call0(__e, SymFn(symshen_4pause_1for_1user))


_ = tmp14120

return False


panic("unreachable")
})

tmp14121 := Call2(__e, ns2_1set, symshen_4show, tmp14113)


_ = tmp14121

tmp14122 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14123 := Call0(__e, SymFn(syminferences))


let__13226 := tmp14123
_ = let__13226

tmp14125 := GEq(__e, MakeInteger(1), let__13226)

var ifres14124 Obj

if True == tmp14125 {
ifres14124 = MakeString("")


} else {
ifres14124 = MakeString("s")


}

tmp14126 := Call3(__e, SymFn(symshen_4app), ifres14124, MakeString(" \n?- "), symshen_4a)


tmp14128 := Call3(__e, SymFn(symshen_4app), let__13226, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" inference"), tmp14126), symshen_4a)


tmp14129 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("____________________________________________________________ "), tmp14128)

tmp14130 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp14129, tmp14130)


panic("unreachable")
})

tmp14131 := Call2(__e, ns2_1set, symshen_4line, tmp14122)


_ = tmp14131

tmp14132 := MakeNative1(func(__e *ControlFlow, V5004 Obj) Obj {
tmp14160 := GConsP(__e, V5004)

var ifres14143 Obj

if True == tmp14160 {
tmp14158 := GTl(__e, V5004)

tmp14159 := GConsP(__e, tmp14158)

var ifres14144 Obj

if True == tmp14159 {
tmp14155 := GTl(__e, V5004)

tmp14156 := GTl(__e, tmp14155)

tmp14157 := GConsP(__e, tmp14156)

var ifres14145 Obj

if True == tmp14157 {
tmp14151 := GTl(__e, V5004)

tmp14152 := GTl(__e, tmp14151)

tmp14153 := GTl(__e, tmp14152)

tmp14154 := GEq(__e, Nil, tmp14153)

var ifres14146 Obj

if True == tmp14154 {
tmp14147 := GTl(__e, V5004)

tmp14148 := GHd(__e, tmp14147)

tmp14149 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp14150 := GEq(__e, tmp14148, tmp14149)

ifres14146 = tmp14150


} else {
ifres14146 = False


}

ifres14145 = ifres14146


} else {
ifres14145 = False


}

ifres14144 = ifres14145


} else {
ifres14144 = False


}

ifres14143 = ifres14144


} else {
ifres14143 = False


}

if True == ifres14143 {
tmp14133 := GHd(__e, V5004)

tmp14134 := Call1(__e, SymFn(symshen_4prterm), tmp14133)


_ = tmp14134

tmp14135 := Call0(__e, SymFn(symstoutput))


tmp14136 := Call2(__e, SymFn(sympr), MakeString(" : "), tmp14135)


_ = tmp14136

tmp14137 := GTl(__e, V5004)

tmp14138 := GTl(__e, tmp14137)

tmp14139 := GHd(__e, tmp14138)

tmp14140 := Call3(__e, SymFn(symshen_4app), tmp14139, MakeString(""), symshen_4r)


tmp14141 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp14140, tmp14141)


} else {
return TailCall1(__e, SymFn(symshen_4prterm), V5004)
}


panic("unreachable")
})

tmp14161 := Call2(__e, ns2_1set, symshen_4show_1p, tmp14132)


_ = tmp14161

tmp14162 := MakeNative1(func(__e *ControlFlow, V5005 Obj) Obj {
tmp14201 := GConsP(__e, V5005)

var ifres14186 Obj

if True == tmp14201 {
tmp14199 := GHd(__e, V5005)

tmp14200 := GEq(__e, symcons, tmp14199)

var ifres14187 Obj

if True == tmp14200 {
tmp14197 := GTl(__e, V5005)

tmp14198 := GConsP(__e, tmp14197)

var ifres14188 Obj

if True == tmp14198 {
tmp14194 := GTl(__e, V5005)

tmp14195 := GTl(__e, tmp14194)

tmp14196 := GConsP(__e, tmp14195)

var ifres14189 Obj

if True == tmp14196 {
tmp14190 := GTl(__e, V5005)

tmp14191 := GTl(__e, tmp14190)

tmp14192 := GTl(__e, tmp14191)

tmp14193 := GEq(__e, Nil, tmp14192)

ifres14189 = tmp14193


} else {
ifres14189 = False


}

ifres14188 = ifres14189


} else {
ifres14188 = False


}

ifres14187 = ifres14188


} else {
ifres14187 = False


}

ifres14186 = ifres14187


} else {
ifres14186 = False


}

if True == ifres14186 {
tmp14163 := Call0(__e, SymFn(symstoutput))


tmp14164 := Call2(__e, SymFn(sympr), MakeString("["), tmp14163)


_ = tmp14164

tmp14165 := GTl(__e, V5005)

tmp14166 := GHd(__e, tmp14165)

tmp14167 := Call1(__e, SymFn(symshen_4prterm), tmp14166)


_ = tmp14167

tmp14168 := GTl(__e, V5005)

tmp14169 := GTl(__e, tmp14168)

tmp14170 := GHd(__e, tmp14169)

tmp14171 := Call1(__e, SymFn(symshen_4prtl), tmp14170)


_ = tmp14171

tmp14172 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), MakeString("]"), tmp14172)


} else {
tmp14184 := GConsP(__e, V5005)

if True == tmp14184 {
tmp14173 := Call0(__e, SymFn(symstoutput))


tmp14174 := Call2(__e, SymFn(sympr), MakeString("("), tmp14173)


_ = tmp14174

tmp14175 := GHd(__e, V5005)

tmp14176 := Call1(__e, SymFn(symshen_4prterm), tmp14175)


_ = tmp14176

tmp14177 := MakeNative1(func(__e *ControlFlow, Z5006 Obj) Obj {
tmp14178 := Call0(__e, SymFn(symstoutput))


tmp14179 := Call2(__e, SymFn(sympr), MakeString(" "), tmp14178)


_ = tmp14179

return TailCall1(__e, SymFn(symshen_4prterm), Z5006)


panic("unreachable")
})

tmp14180 := GTl(__e, V5005)

tmp14181 := Call2(__e, SymFn(symmap), tmp14177, tmp14180)


_ = tmp14181

tmp14182 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), MakeString(")"), tmp14182)


} else {
return TailCall1(__e, SymFn(symprint), V5005)
}


}


panic("unreachable")
})

tmp14202 := Call2(__e, ns2_1set, symshen_4prterm, tmp14162)


_ = tmp14202

tmp14203 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V5007 := __self1
_ = V5007
tmp14232 := GEq(__e, Nil, V5007)

if True == tmp14232 {
return MakeString("")
} else {
tmp14230 := GConsP(__e, V5007)

var ifres14215 Obj

if True == tmp14230 {
tmp14228 := GHd(__e, V5007)

tmp14229 := GEq(__e, symcons, tmp14228)

var ifres14216 Obj

if True == tmp14229 {
tmp14226 := GTl(__e, V5007)

tmp14227 := GConsP(__e, tmp14226)

var ifres14217 Obj

if True == tmp14227 {
tmp14223 := GTl(__e, V5007)

tmp14224 := GTl(__e, tmp14223)

tmp14225 := GConsP(__e, tmp14224)

var ifres14218 Obj

if True == tmp14225 {
tmp14219 := GTl(__e, V5007)

tmp14220 := GTl(__e, tmp14219)

tmp14221 := GTl(__e, tmp14220)

tmp14222 := GEq(__e, Nil, tmp14221)

ifres14218 = tmp14222


} else {
ifres14218 = False


}

ifres14217 = ifres14218


} else {
ifres14217 = False


}

ifres14216 = ifres14217


} else {
ifres14216 = False


}

ifres14215 = ifres14216


} else {
ifres14215 = False


}

if True == ifres14215 {
tmp14204 := Call0(__e, SymFn(symstoutput))


tmp14205 := Call2(__e, SymFn(sympr), MakeString(" "), tmp14204)


_ = tmp14205

tmp14206 := GTl(__e, V5007)

tmp14207 := GHd(__e, tmp14206)

tmp14208 := Call1(__e, SymFn(symshen_4prterm), tmp14207)


_ = tmp14208

tmp14209 := GTl(__e, V5007)

tmp14210 := GTl(__e, tmp14209)

tmp14211 := GHd(__e, tmp14210)

if SymFn(symshen_4prtl) == __self {
__self1 = tmp14211
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4prtl), tmp14211)


} else {
tmp14212 := Call0(__e, SymFn(symstoutput))


tmp14213 := Call2(__e, SymFn(sympr), MakeString(" | "), tmp14212)


_ = tmp14213

return TailCall1(__e, SymFn(symshen_4prterm), V5007)


}


}


panic("unreachable")
})
return __self
})()

tmp14233 := Call2(__e, ns2_1set, symshen_4prtl, tmp14203)


_ = tmp14233

tmp14234 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V5014 := __self1
_ = V5014
V5015 := __self2
_ = V5015
tmp14247 := GEq(__e, Nil, V5014)

if True == tmp14247 {
tmp14235 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), MakeString("\n> "), tmp14235)


} else {
tmp14245 := GConsP(__e, V5014)

if True == tmp14245 {
tmp14236 := Call3(__e, SymFn(symshen_4app), V5015, MakeString(". "), symshen_4a)


tmp14237 := Call0(__e, SymFn(symstoutput))


tmp14238 := Call2(__e, SymFn(sympr), tmp14236, tmp14237)


_ = tmp14238

tmp14239 := GHd(__e, V5014)

tmp14240 := Call1(__e, SymFn(symshen_4show_1p), tmp14239)


_ = tmp14240

tmp14241 := Call1(__e, SymFn(symnl), MakeInteger(1))


_ = tmp14241

tmp14242 := GTl(__e, V5014)

if SymFn(symshen_4show_1assumptions) == __self {
__self1, __self2 = tmp14242, GAdd(__e, V5015, MakeInteger(1))
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4show_1assumptions), tmp14242, GAdd(__e, V5015, MakeInteger(1)))


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.show-assumptions"))
}


}


panic("unreachable")
})
return __self
})()

tmp14248 := Call2(__e, ns2_1set, symshen_4show_1assumptions, tmp14234)


_ = tmp14248

tmp14249 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14250 := Call0(__e, SymFn(symstinput))


tmp14251 := PrimGuard1(__e, symread_1byte, PrimReadByte, tmp14250)

let__13227 := tmp14251
_ = let__13227

tmp14253 := GEq(__e, let__13227, MakeInteger(94))

if True == tmp14253 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("input aborted\n"))
} else {
return TailCall1(__e, SymFn(symnl), MakeInteger(1))
}


panic("unreachable")
})

tmp14254 := Call2(__e, ns2_1set, symshen_4pause_1for_1user, tmp14249)


_ = tmp14254

tmp14255 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dshen_1type_1theory_1enabled_2_d)
panic("unreachable")
})

tmp14256 := Call2(__e, ns2_1set, symshen_4type_1theory_1enabled_2, tmp14255)


_ = tmp14256

tmp14257 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14259 := Call0(__e, SymFn(syminferences))


tmp14260 := GValue(__e, symshen_4_dmaxinferences_d)

if True == GGt(__e, tmp14259, tmp14260) {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("maximum inferences exceeded"))
} else {
return False
}


panic("unreachable")
})

tmp14262 := Call2(__e, ns2_1set, symshen_4maxinfexceeded_2, tmp14257)


_ = tmp14262

tmp14263 := MakeNative(func(__e *ControlFlow) {
V5017 := __e.Get(1)
V5018 := __e.Get(2)
V5019 := __e.Get(3)
V5020 := __e.Get(4)
V5021 := __e.Get(5)
V5022 := __e.Get(6)
V5023 := __e.Get(7)
__e.ReturnDirect(func() Obj {
let__13228 := GAdd(__e, V5022, MakeInteger(1))
_ = let__13228

tmp14274 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres14265 Obj

if True == tmp14274 {
tmp14266 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14266

tmp14267 := GValue(__e, symshen_4_dspy_d)

tmp14268 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14269 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp14270 := GCons(__e, V5018, Nil)

tmp14271 := GCons(__e, tmp14269, tmp14270)

tmp14272 := GCons(__e, V5017, tmp14271)

return TailCall6(__e, SymFn(symshen_4show), tmp14272, V5019, V5020, V5021, let__13228, V5023)


panic("unreachable")
})

tmp14273 := Call5(__e, SymFn(symwhen), tmp14267, V5020, V5021, let__13228, tmp14268)


ifres14265 = tmp14273


} else {
ifres14265 = False


}

let__13229 := ifres14265
_ = let__13229

tmp15013 := GEq(__e, let__13229, False)

if True == tmp15013 {
tmp14282 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres14275 Obj

if True == tmp14282 {
tmp14276 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14276

tmp14277 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


tmp14278 := GConsP(__e, tmp14277)

tmp14279 := GNot(__e, tmp14278)

tmp14280 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall6(__e, SymFn(symshen_4primitive), V5017, V5018, V5020, V5021, let__13228, V5023)
panic("unreachable")
})

tmp14281 := Call5(__e, SymFn(symwhen), tmp14279, V5020, V5021, let__13228, tmp14280)


ifres14275 = tmp14281


} else {
ifres14275 = False


}

let__13230 := ifres14275
_ = let__13230

tmp15011 := GEq(__e, let__13230, False)

if True == tmp15011 {
tmp14286 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres14283 Obj

if True == tmp14286 {
tmp14284 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14284

tmp14285 := Call(__e, SymFn(symshen_4by_1hypothesis), V5017, V5018, V5019, V5020, V5021, let__13228, V5023)


ifres14283 = tmp14285


} else {
ifres14283 = False


}

let__13231 := ifres14283
_ = let__13231

tmp15009 := GEq(__e, let__13231, False)

if True == tmp15009 {
tmp14300 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres14287 Obj

if True == tmp14300 {
tmp14288 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__13233 := tmp14288
_ = let__13233

tmp14299 := GConsP(__e, let__13233)

var ifres14289 Obj

if True == tmp14299 {
tmp14290 := GHd(__e, let__13233)

let__13234 := tmp14290
_ = let__13234

tmp14291 := GTl(__e, let__13233)

tmp14292 := Call2(__e, SymFn(symshen_4lazyderef), tmp14291, V5020)


let__13235 := tmp14292
_ = let__13235

tmp14298 := GEq(__e, let__13235, Nil)

var ifres14293 Obj

if True == tmp14298 {
tmp14294 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14294

tmp14295 := GCons(__e, V5018, Nil)

tmp14296 := GCons(__e, sym_1_1_6, tmp14295)

tmp14297 := Call6(__e, SymFn(symshen_4lookupsig), let__13234, tmp14296, V5020, V5021, let__13228, V5023)


ifres14293 = tmp14297


} else {
ifres14293 = False


}

ifres14289 = ifres14293


} else {
ifres14289 = False


}

ifres14287 = ifres14289


} else {
ifres14287 = False


}

let__13232 := ifres14287
_ = let__13232

tmp15007 := GEq(__e, let__13232, False)

if True == tmp15007 {
tmp14326 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres14301 Obj

if True == tmp14326 {
tmp14302 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__13237 := tmp14302
_ = let__13237

tmp14325 := GConsP(__e, let__13237)

var ifres14303 Obj

if True == tmp14325 {
tmp14304 := GHd(__e, let__13237)

tmp14305 := Call2(__e, SymFn(symshen_4lazyderef), tmp14304, V5020)


let__13238 := tmp14305
_ = let__13238

tmp14324 := GEq(__e, let__13238, symfn)

var ifres14306 Obj

if True == tmp14324 {
tmp14307 := GTl(__e, let__13237)

tmp14308 := Call2(__e, SymFn(symshen_4lazyderef), tmp14307, V5020)


let__13239 := tmp14308
_ = let__13239

tmp14323 := GConsP(__e, let__13239)

var ifres14309 Obj

if True == tmp14323 {
tmp14310 := GHd(__e, let__13239)

let__13240 := tmp14310
_ = let__13240

tmp14311 := GTl(__e, let__13239)

tmp14312 := Call2(__e, SymFn(symshen_4lazyderef), tmp14311, V5020)


let__13241 := tmp14312
_ = let__13241

tmp14322 := GEq(__e, let__13241, Nil)

var ifres14313 Obj

if True == tmp14322 {
tmp14314 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14314

tmp14315 := Call2(__e, SymFn(symshen_4deref), let__13240, V5020)


tmp14316 := Call1(__e, SymFn(symarity), tmp14315)


tmp14317 := GEq(__e, tmp14316, MakeInteger(0))

tmp14318 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14319 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14320 := GCons(__e, let__13240, Nil)

__e.TailApply(SymFn(symshen_4system_1S_1h), tmp14320, V5018, V5019, V5020, V5021, let__13228, V5023)
return nil


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5020, V5021, let__13228, tmp14319)


panic("unreachable")
})

tmp14321 := Call5(__e, SymFn(symwhen), tmp14317, V5020, V5021, let__13228, tmp14318)


ifres14313 = tmp14321


} else {
ifres14313 = False


}

ifres14309 = ifres14313


} else {
ifres14309 = False


}

ifres14306 = ifres14309


} else {
ifres14306 = False


}

ifres14303 = ifres14306


} else {
ifres14303 = False


}

ifres14301 = ifres14303


} else {
ifres14301 = False


}

let__13236 := ifres14301
_ = let__13236

tmp15005 := GEq(__e, let__13236, False)

if True == tmp15005 {
tmp14346 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres14327 Obj

if True == tmp14346 {
tmp14328 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__13243 := tmp14328
_ = let__13243

tmp14345 := GConsP(__e, let__13243)

var ifres14329 Obj

if True == tmp14345 {
tmp14330 := GHd(__e, let__13243)

tmp14331 := Call2(__e, SymFn(symshen_4lazyderef), tmp14330, V5020)


let__13244 := tmp14331
_ = let__13244

tmp14344 := GEq(__e, let__13244, symfn)

var ifres14332 Obj

if True == tmp14344 {
tmp14333 := GTl(__e, let__13243)

tmp14334 := Call2(__e, SymFn(symshen_4lazyderef), tmp14333, V5020)


let__13245 := tmp14334
_ = let__13245

tmp14343 := GConsP(__e, let__13245)

var ifres14335 Obj

if True == tmp14343 {
tmp14336 := GHd(__e, let__13245)

let__13246 := tmp14336
_ = let__13246

tmp14337 := GTl(__e, let__13245)

tmp14338 := Call2(__e, SymFn(symshen_4lazyderef), tmp14337, V5020)


let__13247 := tmp14338
_ = let__13247

tmp14342 := GEq(__e, let__13247, Nil)

var ifres14339 Obj

if True == tmp14342 {
tmp14340 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14340

tmp14341 := Call6(__e, SymFn(symshen_4lookupsig), let__13246, V5018, V5020, V5021, let__13228, V5023)


ifres14339 = tmp14341


} else {
ifres14339 = False


}

ifres14335 = ifres14339


} else {
ifres14335 = False


}

ifres14332 = ifres14335


} else {
ifres14332 = False


}

ifres14329 = ifres14332


} else {
ifres14329 = False


}

ifres14327 = ifres14329


} else {
ifres14327 = False


}

let__13242 := ifres14327
_ = let__13242

tmp15003 := GEq(__e, let__13242, False)

if True == tmp15003 {
tmp14373 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres14347 Obj

if True == tmp14373 {
tmp14348 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__13249 := tmp14348
_ = let__13249

tmp14372 := GConsP(__e, let__13249)

var ifres14349 Obj

if True == tmp14372 {
tmp14350 := GHd(__e, let__13249)

let__13250 := tmp14350
_ = let__13250

tmp14351 := GTl(__e, let__13249)

tmp14352 := Call2(__e, SymFn(symshen_4lazyderef), tmp14351, V5020)


let__13251 := tmp14352
_ = let__13251

tmp14371 := GConsP(__e, let__13251)

var ifres14353 Obj

if True == tmp14371 {
tmp14354 := GHd(__e, let__13251)

let__13252 := tmp14354
_ = let__13252

tmp14355 := GTl(__e, let__13251)

tmp14356 := Call2(__e, SymFn(symshen_4lazyderef), tmp14355, V5020)


let__13253 := tmp14356
_ = let__13253

tmp14370 := GEq(__e, let__13253, Nil)

var ifres14357 Obj

if True == tmp14370 {
tmp14358 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__13254 := tmp14358
_ = let__13254

tmp14359 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14359

tmp14360 := Call2(__e, SymFn(symshen_4lazyderef), let__13250, V5020)


tmp14361 := GConsP(__e, tmp14360)

tmp14362 := GNot(__e, tmp14361)

tmp14363 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14364 := GCons(__e, V5018, Nil)

tmp14365 := GCons(__e, sym_1_1_6, tmp14364)

tmp14366 := GCons(__e, let__13254, tmp14365)

tmp14367 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), let__13252, let__13254, V5019, V5020, V5021, let__13228, V5023)
return nil
panic("unreachable")
})

return TailCall6(__e, SymFn(symshen_4lookupsig), let__13250, tmp14366, V5020, V5021, let__13228, tmp14367)


panic("unreachable")
})

tmp14368 := Call5(__e, SymFn(symwhen), tmp14362, V5020, V5021, let__13228, tmp14363)


tmp14369 := Call2(__e, SymFn(symshen_4gc), V5020, tmp14368)


ifres14357 = tmp14369


} else {
ifres14357 = False


}

ifres14353 = ifres14357


} else {
ifres14353 = False


}

ifres14349 = ifres14353


} else {
ifres14349 = False


}

ifres14347 = ifres14349


} else {
ifres14347 = False


}

let__13248 := ifres14347
_ = let__13248

tmp15001 := GEq(__e, let__13248, False)

if True == tmp15001 {
tmp14396 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres14374 Obj

if True == tmp14396 {
tmp14375 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__13256 := tmp14375
_ = let__13256

tmp14395 := GConsP(__e, let__13256)

var ifres14376 Obj

if True == tmp14395 {
tmp14377 := GHd(__e, let__13256)

let__13257 := tmp14377
_ = let__13257

tmp14378 := GTl(__e, let__13256)

tmp14379 := Call2(__e, SymFn(symshen_4lazyderef), tmp14378, V5020)


let__13258 := tmp14379
_ = let__13258

tmp14394 := GConsP(__e, let__13258)

var ifres14380 Obj

if True == tmp14394 {
tmp14381 := GHd(__e, let__13258)

let__13259 := tmp14381
_ = let__13259

tmp14382 := GTl(__e, let__13258)

tmp14383 := Call2(__e, SymFn(symshen_4lazyderef), tmp14382, V5020)


let__13260 := tmp14383
_ = let__13260

tmp14393 := GEq(__e, let__13260, Nil)

var ifres14384 Obj

if True == tmp14393 {
tmp14385 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__13261 := tmp14385
_ = let__13261

tmp14386 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14386

tmp14387 := GCons(__e, V5018, Nil)

tmp14388 := GCons(__e, sym_1_1_6, tmp14387)

tmp14389 := GCons(__e, let__13261, tmp14388)

tmp14390 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), let__13259, let__13261, V5019, V5020, V5021, let__13228, V5023)
return nil
panic("unreachable")
})

tmp14391 := Call(__e, SymFn(symshen_4system_1S_1h), let__13257, tmp14389, V5019, V5020, V5021, let__13228, tmp14390)


tmp14392 := Call2(__e, SymFn(symshen_4gc), V5020, tmp14391)


ifres14384 = tmp14392


} else {
ifres14384 = False


}

ifres14380 = ifres14384


} else {
ifres14380 = False


}

ifres14376 = ifres14380


} else {
ifres14376 = False


}

ifres14374 = ifres14376


} else {
ifres14374 = False


}

let__13255 := ifres14374
_ = let__13255

tmp14999 := GEq(__e, let__13255, False)

if True == tmp14999 {
tmp14463 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres14397 Obj

if True == tmp14463 {
tmp14398 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__13263 := tmp14398
_ = let__13263

tmp14462 := GConsP(__e, let__13263)

var ifres14399 Obj

if True == tmp14462 {
tmp14400 := GHd(__e, let__13263)

tmp14401 := Call2(__e, SymFn(symshen_4lazyderef), tmp14400, V5020)


let__13264 := tmp14401
_ = let__13264

tmp14461 := GEq(__e, let__13264, symcons)

var ifres14402 Obj

if True == tmp14461 {
tmp14403 := GTl(__e, let__13263)

tmp14404 := Call2(__e, SymFn(symshen_4lazyderef), tmp14403, V5020)


let__13265 := tmp14404
_ = let__13265

tmp14460 := GConsP(__e, let__13265)

var ifres14405 Obj

if True == tmp14460 {
tmp14406 := GHd(__e, let__13265)

let__13266 := tmp14406
_ = let__13266

tmp14407 := GTl(__e, let__13265)

tmp14408 := Call2(__e, SymFn(symshen_4lazyderef), tmp14407, V5020)


let__13267 := tmp14408
_ = let__13267

tmp14459 := GConsP(__e, let__13267)

var ifres14409 Obj

if True == tmp14459 {
tmp14410 := GHd(__e, let__13267)

let__13268 := tmp14410
_ = let__13268

tmp14411 := GTl(__e, let__13267)

tmp14412 := Call2(__e, SymFn(symshen_4lazyderef), tmp14411, V5020)


let__13269 := tmp14412
_ = let__13269

tmp14458 := GEq(__e, let__13269, Nil)

var ifres14413 Obj

if True == tmp14458 {
tmp14414 := Call2(__e, SymFn(symshen_4lazyderef), V5018, V5020)


let__13270 := tmp14414
_ = let__13270

tmp14415 := MakeNative1(func(__e *ControlFlow, Z5068 Obj) Obj {
tmp14416 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14416

tmp14417 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14418 := GCons(__e, Z5068, Nil)

tmp14419 := GCons(__e, symlist, tmp14418)

__e.TailApply(SymFn(symshen_4system_1S_1h), let__13268, tmp14419, V5019, V5020, V5021, let__13228, V5023)
return nil


panic("unreachable")
})

__e.TailApply(SymFn(symshen_4system_1S_1h), let__13266, Z5068, V5019, V5020, V5021, let__13228, tmp14417)
return nil


panic("unreachable")
})

let__13271 := tmp14415
_ = let__13271

tmp14457 := GConsP(__e, let__13270)

var ifres14420 Obj

if True == tmp14457 {
tmp14421 := GHd(__e, let__13270)

tmp14422 := Call2(__e, SymFn(symshen_4lazyderef), tmp14421, V5020)


let__13272 := tmp14422
_ = let__13272

tmp14423 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14424 := GTl(__e, let__13270)

tmp14425 := Call2(__e, SymFn(symshen_4lazyderef), tmp14424, V5020)


let__13274 := tmp14425
_ = let__13274

tmp14426 := MakeNative1(func(__e *ControlFlow, Z5073 Obj) Obj {
return TailCall1(__e, let__13271, Z5073)
panic("unreachable")
})

let__13275 := tmp14426
_ = let__13275

tmp14442 := GConsP(__e, let__13274)

if True == tmp14442 {
tmp14427 := GHd(__e, let__13274)

let__13276 := tmp14427
_ = let__13276

tmp14428 := GTl(__e, let__13274)

tmp14429 := Call2(__e, SymFn(symshen_4lazyderef), tmp14428, V5020)


let__13277 := tmp14429
_ = let__13277

tmp14430 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13275, let__13276)
panic("unreachable")
})

let__13278 := tmp14430
_ = let__13278

tmp14434 := GEq(__e, let__13277, Nil)

if True == tmp14434 {
return TailCall1(__e, SymFn(symthaw), let__13278)
} else {
tmp14432 := Call1(__e, SymFn(symshen_4pvar_2), let__13277)


if True == tmp14432 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__13277, Nil, V5020, let__13278)
} else {
return False
}


}


} else {
tmp14440 := Call1(__e, SymFn(symshen_4pvar_2), let__13274)


if True == tmp14440 {
tmp14435 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__13279 := tmp14435
_ = let__13279

tmp14436 := GCons(__e, let__13279, Nil)

tmp14437 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13275, let__13279)
panic("unreachable")
})

tmp14438 := Call4(__e, SymFn(symshen_4bind_b), let__13274, tmp14436, V5020, tmp14437)


return TailCall2(__e, SymFn(symshen_4gc), V5020, tmp14438)


} else {
return False
}


}


panic("unreachable")
})

let__13273 := tmp14423
_ = let__13273

tmp14448 := GEq(__e, let__13272, symlist)

var ifres14443 Obj

if True == tmp14448 {
tmp14444 := Call1(__e, SymFn(symthaw), let__13273)


ifres14443 = tmp14444


} else {
tmp14447 := Call1(__e, SymFn(symshen_4pvar_2), let__13272)


var ifres14445 Obj

if True == tmp14447 {
tmp14446 := Call4(__e, SymFn(symshen_4bind_b), let__13272, symlist, V5020, let__13273)


ifres14445 = tmp14446


} else {
ifres14445 = False


}

ifres14443 = ifres14445


}

ifres14420 = ifres14443


} else {
tmp14456 := Call1(__e, SymFn(symshen_4pvar_2), let__13270)


var ifres14449 Obj

if True == tmp14456 {
tmp14450 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__13280 := tmp14450
_ = let__13280

tmp14451 := GCons(__e, let__13280, Nil)

tmp14452 := GCons(__e, symlist, tmp14451)

tmp14453 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13271, let__13280)
panic("unreachable")
})

tmp14454 := Call4(__e, SymFn(symshen_4bind_b), let__13270, tmp14452, V5020, tmp14453)


tmp14455 := Call2(__e, SymFn(symshen_4gc), V5020, tmp14454)


ifres14449 = tmp14455


} else {
ifres14449 = False


}

ifres14420 = ifres14449


}

ifres14413 = ifres14420


} else {
ifres14413 = False


}

ifres14409 = ifres14413


} else {
ifres14409 = False


}

ifres14405 = ifres14409


} else {
ifres14405 = False


}

ifres14402 = ifres14405


} else {
ifres14402 = False


}

ifres14399 = ifres14402


} else {
ifres14399 = False


}

ifres14397 = ifres14399


} else {
ifres14397 = False


}

let__13262 := ifres14397
_ = let__13262

tmp14997 := GEq(__e, let__13262, False)

if True == tmp14997 {
tmp14547 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres14464 Obj

if True == tmp14547 {
tmp14465 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__13282 := tmp14465
_ = let__13282

tmp14546 := GConsP(__e, let__13282)

var ifres14466 Obj

if True == tmp14546 {
tmp14467 := GHd(__e, let__13282)

tmp14468 := Call2(__e, SymFn(symshen_4lazyderef), tmp14467, V5020)


let__13283 := tmp14468
_ = let__13283

tmp14545 := GEq(__e, let__13283, sym_8p)

var ifres14469 Obj

if True == tmp14545 {
tmp14470 := GTl(__e, let__13282)

tmp14471 := Call2(__e, SymFn(symshen_4lazyderef), tmp14470, V5020)


let__13284 := tmp14471
_ = let__13284

tmp14544 := GConsP(__e, let__13284)

var ifres14472 Obj

if True == tmp14544 {
tmp14473 := GHd(__e, let__13284)

let__13285 := tmp14473
_ = let__13285

tmp14474 := GTl(__e, let__13284)

tmp14475 := Call2(__e, SymFn(symshen_4lazyderef), tmp14474, V5020)


let__13286 := tmp14475
_ = let__13286

tmp14543 := GConsP(__e, let__13286)

var ifres14476 Obj

if True == tmp14543 {
tmp14477 := GHd(__e, let__13286)

let__13287 := tmp14477
_ = let__13287

tmp14478 := GTl(__e, let__13286)

tmp14479 := Call2(__e, SymFn(symshen_4lazyderef), tmp14478, V5020)


let__13288 := tmp14479
_ = let__13288

tmp14542 := GEq(__e, let__13288, Nil)

var ifres14480 Obj

if True == tmp14542 {
tmp14481 := Call2(__e, SymFn(symshen_4lazyderef), V5018, V5020)


let__13289 := tmp14481
_ = let__13289

tmp14482 := MakeNative1(func(__e *ControlFlow, Z5089 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5090 Obj) Obj {
tmp14483 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14483

tmp14484 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), let__13287, Z5090, V5019, V5020, V5021, let__13228, V5023)
return nil
panic("unreachable")
})

__e.TailApply(SymFn(symshen_4system_1S_1h), let__13285, Z5089, V5019, V5020, V5021, let__13228, tmp14484)
return nil


panic("unreachable")
})
panic("unreachable")
})

let__13290 := tmp14482
_ = let__13290

tmp14541 := GConsP(__e, let__13289)

var ifres14485 Obj

if True == tmp14541 {
tmp14486 := GHd(__e, let__13289)

let__13291 := tmp14486
_ = let__13291

tmp14487 := GTl(__e, let__13289)

tmp14488 := Call2(__e, SymFn(symshen_4lazyderef), tmp14487, V5020)


let__13292 := tmp14488
_ = let__13292

tmp14489 := MakeNative1(func(__e *ControlFlow, Z5094 Obj) Obj {
tmp14490 := Call1(__e, let__13290, let__13291)


return TailCall1(__e, tmp14490, Z5094)


panic("unreachable")
})

let__13293 := tmp14489
_ = let__13293

tmp14528 := GConsP(__e, let__13292)

var ifres14491 Obj

if True == tmp14528 {
tmp14492 := GHd(__e, let__13292)

tmp14493 := Call2(__e, SymFn(symshen_4lazyderef), tmp14492, V5020)


let__13294 := tmp14493
_ = let__13294

tmp14494 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14495 := GTl(__e, let__13292)

tmp14496 := Call2(__e, SymFn(symshen_4lazyderef), tmp14495, V5020)


let__13296 := tmp14496
_ = let__13296

tmp14497 := MakeNative1(func(__e *ControlFlow, Z5099 Obj) Obj {
return TailCall1(__e, let__13293, Z5099)
panic("unreachable")
})

let__13297 := tmp14497
_ = let__13297

tmp14513 := GConsP(__e, let__13296)

if True == tmp14513 {
tmp14498 := GHd(__e, let__13296)

let__13298 := tmp14498
_ = let__13298

tmp14499 := GTl(__e, let__13296)

tmp14500 := Call2(__e, SymFn(symshen_4lazyderef), tmp14499, V5020)


let__13299 := tmp14500
_ = let__13299

tmp14501 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13297, let__13298)
panic("unreachable")
})

let__13300 := tmp14501
_ = let__13300

tmp14505 := GEq(__e, let__13299, Nil)

if True == tmp14505 {
return TailCall1(__e, SymFn(symthaw), let__13300)
} else {
tmp14503 := Call1(__e, SymFn(symshen_4pvar_2), let__13299)


if True == tmp14503 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__13299, Nil, V5020, let__13300)
} else {
return False
}


}


} else {
tmp14511 := Call1(__e, SymFn(symshen_4pvar_2), let__13296)


if True == tmp14511 {
tmp14506 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__13301 := tmp14506
_ = let__13301

tmp14507 := GCons(__e, let__13301, Nil)

tmp14508 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13297, let__13301)
panic("unreachable")
})

tmp14509 := Call4(__e, SymFn(symshen_4bind_b), let__13296, tmp14507, V5020, tmp14508)


return TailCall2(__e, SymFn(symshen_4gc), V5020, tmp14509)


} else {
return False
}


}


panic("unreachable")
})

let__13295 := tmp14494
_ = let__13295

tmp14519 := GEq(__e, let__13294, sym_d)

var ifres14514 Obj

if True == tmp14519 {
tmp14515 := Call1(__e, SymFn(symthaw), let__13295)


ifres14514 = tmp14515


} else {
tmp14518 := Call1(__e, SymFn(symshen_4pvar_2), let__13294)


var ifres14516 Obj

if True == tmp14518 {
tmp14517 := Call4(__e, SymFn(symshen_4bind_b), let__13294, sym_d, V5020, let__13295)


ifres14516 = tmp14517


} else {
ifres14516 = False


}

ifres14514 = ifres14516


}

ifres14491 = ifres14514


} else {
tmp14527 := Call1(__e, SymFn(symshen_4pvar_2), let__13292)


var ifres14520 Obj

if True == tmp14527 {
tmp14521 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__13302 := tmp14521
_ = let__13302

tmp14522 := GCons(__e, let__13302, Nil)

tmp14523 := GCons(__e, sym_d, tmp14522)

tmp14524 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13293, let__13302)
panic("unreachable")
})

tmp14525 := Call4(__e, SymFn(symshen_4bind_b), let__13292, tmp14523, V5020, tmp14524)


tmp14526 := Call2(__e, SymFn(symshen_4gc), V5020, tmp14525)


ifres14520 = tmp14526


} else {
ifres14520 = False


}

ifres14491 = ifres14520


}

ifres14485 = ifres14491


} else {
tmp14540 := Call1(__e, SymFn(symshen_4pvar_2), let__13289)


var ifres14529 Obj

if True == tmp14540 {
tmp14530 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__13303 := tmp14530
_ = let__13303

tmp14531 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__13304 := tmp14531
_ = let__13304

tmp14532 := GCons(__e, let__13304, Nil)

tmp14533 := GCons(__e, sym_d, tmp14532)

tmp14534 := GCons(__e, let__13303, tmp14533)

tmp14535 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14536 := Call1(__e, let__13290, let__13303)


return TailCall1(__e, tmp14536, let__13304)


panic("unreachable")
})

tmp14537 := Call4(__e, SymFn(symshen_4bind_b), let__13289, tmp14534, V5020, tmp14535)


tmp14538 := Call2(__e, SymFn(symshen_4gc), V5020, tmp14537)


tmp14539 := Call2(__e, SymFn(symshen_4gc), V5020, tmp14538)


ifres14529 = tmp14539


} else {
ifres14529 = False


}

ifres14485 = ifres14529


}

ifres14480 = ifres14485


} else {
ifres14480 = False


}

ifres14476 = ifres14480


} else {
ifres14476 = False


}

ifres14472 = ifres14476


} else {
ifres14472 = False


}

ifres14469 = ifres14472


} else {
ifres14469 = False


}

ifres14466 = ifres14469


} else {
ifres14466 = False


}

ifres14464 = ifres14466


} else {
ifres14464 = False


}

let__13281 := ifres14464
_ = let__13281

tmp14995 := GEq(__e, let__13281, False)

if True == tmp14995 {
tmp14614 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres14548 Obj

if True == tmp14614 {
tmp14549 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__13306 := tmp14549
_ = let__13306

tmp14613 := GConsP(__e, let__13306)

var ifres14550 Obj

if True == tmp14613 {
tmp14551 := GHd(__e, let__13306)

tmp14552 := Call2(__e, SymFn(symshen_4lazyderef), tmp14551, V5020)


let__13307 := tmp14552
_ = let__13307

tmp14612 := GEq(__e, let__13307, sym_8v)

var ifres14553 Obj

if True == tmp14612 {
tmp14554 := GTl(__e, let__13306)

tmp14555 := Call2(__e, SymFn(symshen_4lazyderef), tmp14554, V5020)


let__13308 := tmp14555
_ = let__13308

tmp14611 := GConsP(__e, let__13308)

var ifres14556 Obj

if True == tmp14611 {
tmp14557 := GHd(__e, let__13308)

let__13309 := tmp14557
_ = let__13309

tmp14558 := GTl(__e, let__13308)

tmp14559 := Call2(__e, SymFn(symshen_4lazyderef), tmp14558, V5020)


let__13310 := tmp14559
_ = let__13310

tmp14610 := GConsP(__e, let__13310)

var ifres14560 Obj

if True == tmp14610 {
tmp14561 := GHd(__e, let__13310)

let__13311 := tmp14561
_ = let__13311

tmp14562 := GTl(__e, let__13310)

tmp14563 := Call2(__e, SymFn(symshen_4lazyderef), tmp14562, V5020)


let__13312 := tmp14563
_ = let__13312

tmp14609 := GEq(__e, let__13312, Nil)

var ifres14564 Obj

if True == tmp14609 {
tmp14565 := Call2(__e, SymFn(symshen_4lazyderef), V5018, V5020)


let__13313 := tmp14565
_ = let__13313

tmp14566 := MakeNative1(func(__e *ControlFlow, Z5117 Obj) Obj {
tmp14567 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14567

tmp14568 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14569 := GCons(__e, Z5117, Nil)

tmp14570 := GCons(__e, symvector, tmp14569)

__e.TailApply(SymFn(symshen_4system_1S_1h), let__13311, tmp14570, V5019, V5020, V5021, let__13228, V5023)
return nil


panic("unreachable")
})

__e.TailApply(SymFn(symshen_4system_1S_1h), let__13309, Z5117, V5019, V5020, V5021, let__13228, tmp14568)
return nil


panic("unreachable")
})

let__13314 := tmp14566
_ = let__13314

tmp14608 := GConsP(__e, let__13313)

var ifres14571 Obj

if True == tmp14608 {
tmp14572 := GHd(__e, let__13313)

tmp14573 := Call2(__e, SymFn(symshen_4lazyderef), tmp14572, V5020)


let__13315 := tmp14573
_ = let__13315

tmp14574 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14575 := GTl(__e, let__13313)

tmp14576 := Call2(__e, SymFn(symshen_4lazyderef), tmp14575, V5020)


let__13317 := tmp14576
_ = let__13317

tmp14577 := MakeNative1(func(__e *ControlFlow, Z5122 Obj) Obj {
return TailCall1(__e, let__13314, Z5122)
panic("unreachable")
})

let__13318 := tmp14577
_ = let__13318

tmp14593 := GConsP(__e, let__13317)

if True == tmp14593 {
tmp14578 := GHd(__e, let__13317)

let__13319 := tmp14578
_ = let__13319

tmp14579 := GTl(__e, let__13317)

tmp14580 := Call2(__e, SymFn(symshen_4lazyderef), tmp14579, V5020)


let__13320 := tmp14580
_ = let__13320

tmp14581 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13318, let__13319)
panic("unreachable")
})

let__13321 := tmp14581
_ = let__13321

tmp14585 := GEq(__e, let__13320, Nil)

if True == tmp14585 {
return TailCall1(__e, SymFn(symthaw), let__13321)
} else {
tmp14583 := Call1(__e, SymFn(symshen_4pvar_2), let__13320)


if True == tmp14583 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__13320, Nil, V5020, let__13321)
} else {
return False
}


}


} else {
tmp14591 := Call1(__e, SymFn(symshen_4pvar_2), let__13317)


if True == tmp14591 {
tmp14586 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__13322 := tmp14586
_ = let__13322

tmp14587 := GCons(__e, let__13322, Nil)

tmp14588 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13318, let__13322)
panic("unreachable")
})

tmp14589 := Call4(__e, SymFn(symshen_4bind_b), let__13317, tmp14587, V5020, tmp14588)


return TailCall2(__e, SymFn(symshen_4gc), V5020, tmp14589)


} else {
return False
}


}


panic("unreachable")
})

let__13316 := tmp14574
_ = let__13316

tmp14599 := GEq(__e, let__13315, symvector)

var ifres14594 Obj

if True == tmp14599 {
tmp14595 := Call1(__e, SymFn(symthaw), let__13316)


ifres14594 = tmp14595


} else {
tmp14598 := Call1(__e, SymFn(symshen_4pvar_2), let__13315)


var ifres14596 Obj

if True == tmp14598 {
tmp14597 := Call4(__e, SymFn(symshen_4bind_b), let__13315, symvector, V5020, let__13316)


ifres14596 = tmp14597


} else {
ifres14596 = False


}

ifres14594 = ifres14596


}

ifres14571 = ifres14594


} else {
tmp14607 := Call1(__e, SymFn(symshen_4pvar_2), let__13313)


var ifres14600 Obj

if True == tmp14607 {
tmp14601 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__13323 := tmp14601
_ = let__13323

tmp14602 := GCons(__e, let__13323, Nil)

tmp14603 := GCons(__e, symvector, tmp14602)

tmp14604 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13314, let__13323)
panic("unreachable")
})

tmp14605 := Call4(__e, SymFn(symshen_4bind_b), let__13313, tmp14603, V5020, tmp14604)


tmp14606 := Call2(__e, SymFn(symshen_4gc), V5020, tmp14605)


ifres14600 = tmp14606


} else {
ifres14600 = False


}

ifres14571 = ifres14600


}

ifres14564 = ifres14571


} else {
ifres14564 = False


}

ifres14560 = ifres14564


} else {
ifres14560 = False


}

ifres14556 = ifres14560


} else {
ifres14556 = False


}

ifres14553 = ifres14556


} else {
ifres14553 = False


}

ifres14550 = ifres14553


} else {
ifres14550 = False


}

ifres14548 = ifres14550


} else {
ifres14548 = False


}

let__13305 := ifres14548
_ = let__13305

tmp14993 := GEq(__e, let__13305, False)

if True == tmp14993 {
tmp14647 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres14615 Obj

if True == tmp14647 {
tmp14616 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__13325 := tmp14616
_ = let__13325

tmp14646 := GConsP(__e, let__13325)

var ifres14617 Obj

if True == tmp14646 {
tmp14618 := GHd(__e, let__13325)

tmp14619 := Call2(__e, SymFn(symshen_4lazyderef), tmp14618, V5020)


let__13326 := tmp14619
_ = let__13326

tmp14645 := GEq(__e, let__13326, sym_8s)

var ifres14620 Obj

if True == tmp14645 {
tmp14621 := GTl(__e, let__13325)

tmp14622 := Call2(__e, SymFn(symshen_4lazyderef), tmp14621, V5020)


let__13327 := tmp14622
_ = let__13327

tmp14644 := GConsP(__e, let__13327)

var ifres14623 Obj

if True == tmp14644 {
tmp14624 := GHd(__e, let__13327)

let__13328 := tmp14624
_ = let__13328

tmp14625 := GTl(__e, let__13327)

tmp14626 := Call2(__e, SymFn(symshen_4lazyderef), tmp14625, V5020)


let__13329 := tmp14626
_ = let__13329

tmp14643 := GConsP(__e, let__13329)

var ifres14627 Obj

if True == tmp14643 {
tmp14628 := GHd(__e, let__13329)

let__13330 := tmp14628
_ = let__13330

tmp14629 := GTl(__e, let__13329)

tmp14630 := Call2(__e, SymFn(symshen_4lazyderef), tmp14629, V5020)


let__13331 := tmp14630
_ = let__13331

tmp14642 := GEq(__e, let__13331, Nil)

var ifres14631 Obj

if True == tmp14642 {
tmp14632 := Call2(__e, SymFn(symshen_4lazyderef), V5018, V5020)


let__13332 := tmp14632
_ = let__13332

tmp14633 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14634 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14634

tmp14635 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), let__13330, symstring, V5019, V5020, V5021, let__13228, V5023)
return nil
panic("unreachable")
})

__e.TailApply(SymFn(symshen_4system_1S_1h), let__13328, symstring, V5019, V5020, V5021, let__13228, tmp14635)
return nil


panic("unreachable")
})

let__13333 := tmp14633
_ = let__13333

tmp14641 := GEq(__e, let__13332, symstring)

var ifres14636 Obj

if True == tmp14641 {
tmp14637 := Call1(__e, SymFn(symthaw), let__13333)


ifres14636 = tmp14637


} else {
tmp14640 := Call1(__e, SymFn(symshen_4pvar_2), let__13332)


var ifres14638 Obj

if True == tmp14640 {
tmp14639 := Call4(__e, SymFn(symshen_4bind_b), let__13332, symstring, V5020, let__13333)


ifres14638 = tmp14639


} else {
ifres14638 = False


}

ifres14636 = ifres14638


}

ifres14631 = ifres14636


} else {
ifres14631 = False


}

ifres14627 = ifres14631


} else {
ifres14627 = False


}

ifres14623 = ifres14627


} else {
ifres14623 = False


}

ifres14620 = ifres14623


} else {
ifres14620 = False


}

ifres14617 = ifres14620


} else {
ifres14617 = False


}

ifres14615 = ifres14617


} else {
ifres14615 = False


}

let__13324 := ifres14615
_ = let__13324

tmp14991 := GEq(__e, let__13324, False)

if True == tmp14991 {
tmp14747 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres14648 Obj

if True == tmp14747 {
tmp14649 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__13335 := tmp14649
_ = let__13335

tmp14746 := GConsP(__e, let__13335)

var ifres14650 Obj

if True == tmp14746 {
tmp14651 := GHd(__e, let__13335)

tmp14652 := Call2(__e, SymFn(symshen_4lazyderef), tmp14651, V5020)


let__13336 := tmp14652
_ = let__13336

tmp14745 := GEq(__e, let__13336, symlambda)

var ifres14653 Obj

if True == tmp14745 {
tmp14654 := GTl(__e, let__13335)

tmp14655 := Call2(__e, SymFn(symshen_4lazyderef), tmp14654, V5020)


let__13337 := tmp14655
_ = let__13337

tmp14744 := GConsP(__e, let__13337)

var ifres14656 Obj

if True == tmp14744 {
tmp14657 := GHd(__e, let__13337)

let__13338 := tmp14657
_ = let__13338

tmp14658 := GTl(__e, let__13337)

tmp14659 := Call2(__e, SymFn(symshen_4lazyderef), tmp14658, V5020)


let__13339 := tmp14659
_ = let__13339

tmp14743 := GConsP(__e, let__13339)

var ifres14660 Obj

if True == tmp14743 {
tmp14661 := GHd(__e, let__13339)

let__13340 := tmp14661
_ = let__13340

tmp14662 := GTl(__e, let__13339)

tmp14663 := Call2(__e, SymFn(symshen_4lazyderef), tmp14662, V5020)


let__13341 := tmp14663
_ = let__13341

tmp14742 := GEq(__e, let__13341, Nil)

var ifres14664 Obj

if True == tmp14742 {
tmp14665 := Call2(__e, SymFn(symshen_4lazyderef), V5018, V5020)


let__13342 := tmp14665
_ = let__13342

tmp14666 := MakeNative1(func(__e *ControlFlow, Z5148 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5149 Obj) Obj {
tmp14667 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__13344 := tmp14667
_ = let__13344

tmp14668 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__13345 := tmp14668
_ = let__13345

tmp14669 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14669

tmp14670 := Call2(__e, SymFn(symshen_4lazyderef), let__13338, V5020)


tmp14671 := Call1(__e, SymFn(symshen_4freshterm), tmp14670)


tmp14672 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14673 := Call2(__e, SymFn(symshen_4lazyderef), let__13338, V5020)


tmp14674 := Call2(__e, SymFn(symshen_4deref), let__13345, V5020)


tmp14675 := Call2(__e, SymFn(symshen_4deref), let__13340, V5020)


tmp14676 := Call3(__e, SymFn(symshen_4beta), tmp14673, tmp14674, tmp14675)


tmp14677 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14678 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp14679 := GCons(__e, Z5148, Nil)

tmp14680 := GCons(__e, tmp14678, tmp14679)

tmp14681 := GCons(__e, let__13345, tmp14680)

tmp14682 := GCons(__e, tmp14681, V5019)

__e.TailApply(SymFn(symshen_4system_1S_1h), let__13344, Z5149, tmp14682, V5020, V5021, let__13228, V5023)
return nil


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), let__13344, tmp14676, V5020, V5021, let__13228, tmp14677)


panic("unreachable")
})

tmp14683 := Call6(__e, SymFn(symbind), let__13345, tmp14671, V5020, V5021, let__13228, tmp14672)


tmp14684 := Call2(__e, SymFn(symshen_4gc), V5020, tmp14683)


return TailCall2(__e, SymFn(symshen_4gc), V5020, tmp14684)


panic("unreachable")
})
panic("unreachable")
})

let__13343 := tmp14666
_ = let__13343

tmp14741 := GConsP(__e, let__13342)

var ifres14685 Obj

if True == tmp14741 {
tmp14686 := GHd(__e, let__13342)

let__13346 := tmp14686
_ = let__13346

tmp14687 := GTl(__e, let__13342)

tmp14688 := Call2(__e, SymFn(symshen_4lazyderef), tmp14687, V5020)


let__13347 := tmp14688
_ = let__13347

tmp14689 := MakeNative1(func(__e *ControlFlow, Z5155 Obj) Obj {
tmp14690 := Call1(__e, let__13343, let__13346)


return TailCall1(__e, tmp14690, Z5155)


panic("unreachable")
})

let__13348 := tmp14689
_ = let__13348

tmp14728 := GConsP(__e, let__13347)

var ifres14691 Obj

if True == tmp14728 {
tmp14692 := GHd(__e, let__13347)

tmp14693 := Call2(__e, SymFn(symshen_4lazyderef), tmp14692, V5020)


let__13349 := tmp14693
_ = let__13349

tmp14694 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14695 := GTl(__e, let__13347)

tmp14696 := Call2(__e, SymFn(symshen_4lazyderef), tmp14695, V5020)


let__13351 := tmp14696
_ = let__13351

tmp14697 := MakeNative1(func(__e *ControlFlow, Z5160 Obj) Obj {
return TailCall1(__e, let__13348, Z5160)
panic("unreachable")
})

let__13352 := tmp14697
_ = let__13352

tmp14713 := GConsP(__e, let__13351)

if True == tmp14713 {
tmp14698 := GHd(__e, let__13351)

let__13353 := tmp14698
_ = let__13353

tmp14699 := GTl(__e, let__13351)

tmp14700 := Call2(__e, SymFn(symshen_4lazyderef), tmp14699, V5020)


let__13354 := tmp14700
_ = let__13354

tmp14701 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13352, let__13353)
panic("unreachable")
})

let__13355 := tmp14701
_ = let__13355

tmp14705 := GEq(__e, let__13354, Nil)

if True == tmp14705 {
return TailCall1(__e, SymFn(symthaw), let__13355)
} else {
tmp14703 := Call1(__e, SymFn(symshen_4pvar_2), let__13354)


if True == tmp14703 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__13354, Nil, V5020, let__13355)
} else {
return False
}


}


} else {
tmp14711 := Call1(__e, SymFn(symshen_4pvar_2), let__13351)


if True == tmp14711 {
tmp14706 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__13356 := tmp14706
_ = let__13356

tmp14707 := GCons(__e, let__13356, Nil)

tmp14708 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13352, let__13356)
panic("unreachable")
})

tmp14709 := Call4(__e, SymFn(symshen_4bind_b), let__13351, tmp14707, V5020, tmp14708)


return TailCall2(__e, SymFn(symshen_4gc), V5020, tmp14709)


} else {
return False
}


}


panic("unreachable")
})

let__13350 := tmp14694
_ = let__13350

tmp14719 := GEq(__e, let__13349, sym_1_1_6)

var ifres14714 Obj

if True == tmp14719 {
tmp14715 := Call1(__e, SymFn(symthaw), let__13350)


ifres14714 = tmp14715


} else {
tmp14718 := Call1(__e, SymFn(symshen_4pvar_2), let__13349)


var ifres14716 Obj

if True == tmp14718 {
tmp14717 := Call4(__e, SymFn(symshen_4bind_b), let__13349, sym_1_1_6, V5020, let__13350)


ifres14716 = tmp14717


} else {
ifres14716 = False


}

ifres14714 = ifres14716


}

ifres14691 = ifres14714


} else {
tmp14727 := Call1(__e, SymFn(symshen_4pvar_2), let__13347)


var ifres14720 Obj

if True == tmp14727 {
tmp14721 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__13357 := tmp14721
_ = let__13357

tmp14722 := GCons(__e, let__13357, Nil)

tmp14723 := GCons(__e, sym_1_1_6, tmp14722)

tmp14724 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13348, let__13357)
panic("unreachable")
})

tmp14725 := Call4(__e, SymFn(symshen_4bind_b), let__13347, tmp14723, V5020, tmp14724)


tmp14726 := Call2(__e, SymFn(symshen_4gc), V5020, tmp14725)


ifres14720 = tmp14726


} else {
ifres14720 = False


}

ifres14691 = ifres14720


}

ifres14685 = ifres14691


} else {
tmp14740 := Call1(__e, SymFn(symshen_4pvar_2), let__13342)


var ifres14729 Obj

if True == tmp14740 {
tmp14730 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__13358 := tmp14730
_ = let__13358

tmp14731 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__13359 := tmp14731
_ = let__13359

tmp14732 := GCons(__e, let__13359, Nil)

tmp14733 := GCons(__e, sym_1_1_6, tmp14732)

tmp14734 := GCons(__e, let__13358, tmp14733)

tmp14735 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14736 := Call1(__e, let__13343, let__13358)


return TailCall1(__e, tmp14736, let__13359)


panic("unreachable")
})

tmp14737 := Call4(__e, SymFn(symshen_4bind_b), let__13342, tmp14734, V5020, tmp14735)


tmp14738 := Call2(__e, SymFn(symshen_4gc), V5020, tmp14737)


tmp14739 := Call2(__e, SymFn(symshen_4gc), V5020, tmp14738)


ifres14729 = tmp14739


} else {
ifres14729 = False


}

ifres14685 = ifres14729


}

ifres14664 = ifres14685


} else {
ifres14664 = False


}

ifres14660 = ifres14664


} else {
ifres14660 = False


}

ifres14656 = ifres14660


} else {
ifres14656 = False


}

ifres14653 = ifres14656


} else {
ifres14653 = False


}

ifres14650 = ifres14653


} else {
ifres14650 = False


}

ifres14648 = ifres14650


} else {
ifres14648 = False


}

let__13334 := ifres14648
_ = let__13334

tmp14989 := GEq(__e, let__13334, False)

if True == tmp14989 {
tmp14797 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres14748 Obj

if True == tmp14797 {
tmp14749 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__13361 := tmp14749
_ = let__13361

tmp14796 := GConsP(__e, let__13361)

var ifres14750 Obj

if True == tmp14796 {
tmp14751 := GHd(__e, let__13361)

tmp14752 := Call2(__e, SymFn(symshen_4lazyderef), tmp14751, V5020)


let__13362 := tmp14752
_ = let__13362

tmp14795 := GEq(__e, let__13362, symlet)

var ifres14753 Obj

if True == tmp14795 {
tmp14754 := GTl(__e, let__13361)

tmp14755 := Call2(__e, SymFn(symshen_4lazyderef), tmp14754, V5020)


let__13363 := tmp14755
_ = let__13363

tmp14794 := GConsP(__e, let__13363)

var ifres14756 Obj

if True == tmp14794 {
tmp14757 := GHd(__e, let__13363)

let__13364 := tmp14757
_ = let__13364

tmp14758 := GTl(__e, let__13363)

tmp14759 := Call2(__e, SymFn(symshen_4lazyderef), tmp14758, V5020)


let__13365 := tmp14759
_ = let__13365

tmp14793 := GConsP(__e, let__13365)

var ifres14760 Obj

if True == tmp14793 {
tmp14761 := GHd(__e, let__13365)

let__13366 := tmp14761
_ = let__13366

tmp14762 := GTl(__e, let__13365)

tmp14763 := Call2(__e, SymFn(symshen_4lazyderef), tmp14762, V5020)


let__13367 := tmp14763
_ = let__13367

tmp14792 := GConsP(__e, let__13367)

var ifres14764 Obj

if True == tmp14792 {
tmp14765 := GHd(__e, let__13367)

let__13368 := tmp14765
_ = let__13368

tmp14766 := GTl(__e, let__13367)

tmp14767 := Call2(__e, SymFn(symshen_4lazyderef), tmp14766, V5020)


let__13369 := tmp14767
_ = let__13369

tmp14791 := GEq(__e, let__13369, Nil)

var ifres14768 Obj

if True == tmp14791 {
tmp14769 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__13370 := tmp14769
_ = let__13370

tmp14770 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__13371 := tmp14770
_ = let__13371

tmp14771 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__13372 := tmp14771
_ = let__13372

tmp14772 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14772

tmp14773 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14774 := Call2(__e, SymFn(symshen_4lazyderef), let__13364, V5020)


tmp14775 := Call1(__e, SymFn(symshen_4freshterm), tmp14774)


tmp14776 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14777 := Call2(__e, SymFn(symshen_4lazyderef), let__13364, V5020)


tmp14778 := Call2(__e, SymFn(symshen_4lazyderef), let__13371, V5020)


tmp14779 := Call2(__e, SymFn(symshen_4lazyderef), let__13368, V5020)


tmp14780 := Call3(__e, SymFn(symshen_4beta), tmp14777, tmp14778, tmp14779)


tmp14781 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14782 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp14783 := GCons(__e, let__13372, Nil)

tmp14784 := GCons(__e, tmp14782, tmp14783)

tmp14785 := GCons(__e, let__13371, tmp14784)

tmp14786 := GCons(__e, tmp14785, V5019)

__e.TailApply(SymFn(symshen_4system_1S_1h), let__13370, V5018, tmp14786, V5020, V5021, let__13228, V5023)
return nil


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), let__13370, tmp14780, V5020, V5021, let__13228, tmp14781)


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), let__13371, tmp14775, V5020, V5021, let__13228, tmp14776)


panic("unreachable")
})

tmp14787 := Call(__e, SymFn(symshen_4system_1S_1h), let__13366, let__13372, V5019, V5020, V5021, let__13228, tmp14773)


tmp14788 := Call2(__e, SymFn(symshen_4gc), V5020, tmp14787)


tmp14789 := Call2(__e, SymFn(symshen_4gc), V5020, tmp14788)


tmp14790 := Call2(__e, SymFn(symshen_4gc), V5020, tmp14789)


ifres14768 = tmp14790


} else {
ifres14768 = False


}

ifres14764 = ifres14768


} else {
ifres14764 = False


}

ifres14760 = ifres14764


} else {
ifres14760 = False


}

ifres14756 = ifres14760


} else {
ifres14756 = False


}

ifres14753 = ifres14756


} else {
ifres14753 = False


}

ifres14750 = ifres14753


} else {
ifres14750 = False


}

ifres14748 = ifres14750


} else {
ifres14748 = False


}

let__13360 := ifres14748
_ = let__13360

tmp14987 := GEq(__e, let__13360, False)

if True == tmp14987 {
tmp14867 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres14798 Obj

if True == tmp14867 {
tmp14799 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__13374 := tmp14799
_ = let__13374

tmp14866 := GConsP(__e, let__13374)

var ifres14800 Obj

if True == tmp14866 {
tmp14801 := GHd(__e, let__13374)

tmp14802 := Call2(__e, SymFn(symshen_4lazyderef), tmp14801, V5020)


let__13375 := tmp14802
_ = let__13375

tmp14865 := GEq(__e, let__13375, symopen)

var ifres14803 Obj

if True == tmp14865 {
tmp14804 := GTl(__e, let__13374)

tmp14805 := Call2(__e, SymFn(symshen_4lazyderef), tmp14804, V5020)


let__13376 := tmp14805
_ = let__13376

tmp14864 := GConsP(__e, let__13376)

var ifres14806 Obj

if True == tmp14864 {
tmp14807 := GHd(__e, let__13376)

let__13377 := tmp14807
_ = let__13377

tmp14808 := GTl(__e, let__13376)

tmp14809 := Call2(__e, SymFn(symshen_4lazyderef), tmp14808, V5020)


let__13378 := tmp14809
_ = let__13378

tmp14863 := GConsP(__e, let__13378)

var ifres14810 Obj

if True == tmp14863 {
tmp14811 := GHd(__e, let__13378)

let__13379 := tmp14811
_ = let__13379

tmp14812 := GTl(__e, let__13378)

tmp14813 := Call2(__e, SymFn(symshen_4lazyderef), tmp14812, V5020)


let__13380 := tmp14813
_ = let__13380

tmp14862 := GEq(__e, let__13380, Nil)

var ifres14814 Obj

if True == tmp14862 {
tmp14815 := Call2(__e, SymFn(symshen_4lazyderef), V5018, V5020)


let__13381 := tmp14815
_ = let__13381

tmp14816 := MakeNative1(func(__e *ControlFlow, Z5191 Obj) Obj {
tmp14817 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14817

tmp14818 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14819 := Call2(__e, SymFn(symshen_4lazyderef), Z5191, V5020)


tmp14820 := GCons(__e, symout, Nil)

tmp14821 := GCons(__e, symin, tmp14820)

tmp14822 := Call2(__e, SymFn(symelement_2), tmp14819, tmp14821)


tmp14823 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), let__13377, symstring, V5019, V5020, V5021, let__13228, V5023)
return nil
panic("unreachable")
})

return TailCall5(__e, SymFn(symwhen), tmp14822, V5020, V5021, let__13228, tmp14823)


panic("unreachable")
})

return TailCall6(__e, SymFn(symis_b), let__13379, Z5191, V5020, V5021, let__13228, tmp14818)


panic("unreachable")
})

let__13382 := tmp14816
_ = let__13382

tmp14861 := GConsP(__e, let__13381)

var ifres14824 Obj

if True == tmp14861 {
tmp14825 := GHd(__e, let__13381)

tmp14826 := Call2(__e, SymFn(symshen_4lazyderef), tmp14825, V5020)


let__13383 := tmp14826
_ = let__13383

tmp14827 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14828 := GTl(__e, let__13381)

tmp14829 := Call2(__e, SymFn(symshen_4lazyderef), tmp14828, V5020)


let__13385 := tmp14829
_ = let__13385

tmp14830 := MakeNative1(func(__e *ControlFlow, Z5196 Obj) Obj {
return TailCall1(__e, let__13382, Z5196)
panic("unreachable")
})

let__13386 := tmp14830
_ = let__13386

tmp14846 := GConsP(__e, let__13385)

if True == tmp14846 {
tmp14831 := GHd(__e, let__13385)

let__13387 := tmp14831
_ = let__13387

tmp14832 := GTl(__e, let__13385)

tmp14833 := Call2(__e, SymFn(symshen_4lazyderef), tmp14832, V5020)


let__13388 := tmp14833
_ = let__13388

tmp14834 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13386, let__13387)
panic("unreachable")
})

let__13389 := tmp14834
_ = let__13389

tmp14838 := GEq(__e, let__13388, Nil)

if True == tmp14838 {
return TailCall1(__e, SymFn(symthaw), let__13389)
} else {
tmp14836 := Call1(__e, SymFn(symshen_4pvar_2), let__13388)


if True == tmp14836 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__13388, Nil, V5020, let__13389)
} else {
return False
}


}


} else {
tmp14844 := Call1(__e, SymFn(symshen_4pvar_2), let__13385)


if True == tmp14844 {
tmp14839 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__13390 := tmp14839
_ = let__13390

tmp14840 := GCons(__e, let__13390, Nil)

tmp14841 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13386, let__13390)
panic("unreachable")
})

tmp14842 := Call4(__e, SymFn(symshen_4bind_b), let__13385, tmp14840, V5020, tmp14841)


return TailCall2(__e, SymFn(symshen_4gc), V5020, tmp14842)


} else {
return False
}


}


panic("unreachable")
})

let__13384 := tmp14827
_ = let__13384

tmp14852 := GEq(__e, let__13383, symstream)

var ifres14847 Obj

if True == tmp14852 {
tmp14848 := Call1(__e, SymFn(symthaw), let__13384)


ifres14847 = tmp14848


} else {
tmp14851 := Call1(__e, SymFn(symshen_4pvar_2), let__13383)


var ifres14849 Obj

if True == tmp14851 {
tmp14850 := Call4(__e, SymFn(symshen_4bind_b), let__13383, symstream, V5020, let__13384)


ifres14849 = tmp14850


} else {
ifres14849 = False


}

ifres14847 = ifres14849


}

ifres14824 = ifres14847


} else {
tmp14860 := Call1(__e, SymFn(symshen_4pvar_2), let__13381)


var ifres14853 Obj

if True == tmp14860 {
tmp14854 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__13391 := tmp14854
_ = let__13391

tmp14855 := GCons(__e, let__13391, Nil)

tmp14856 := GCons(__e, symstream, tmp14855)

tmp14857 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13382, let__13391)
panic("unreachable")
})

tmp14858 := Call4(__e, SymFn(symshen_4bind_b), let__13381, tmp14856, V5020, tmp14857)


tmp14859 := Call2(__e, SymFn(symshen_4gc), V5020, tmp14858)


ifres14853 = tmp14859


} else {
ifres14853 = False


}

ifres14824 = ifres14853


}

ifres14814 = ifres14824


} else {
ifres14814 = False


}

ifres14810 = ifres14814


} else {
ifres14810 = False


}

ifres14806 = ifres14810


} else {
ifres14806 = False


}

ifres14803 = ifres14806


} else {
ifres14803 = False


}

ifres14800 = ifres14803


} else {
ifres14800 = False


}

ifres14798 = ifres14800


} else {
ifres14798 = False


}

let__13373 := ifres14798
_ = let__13373

tmp14985 := GEq(__e, let__13373, False)

if True == tmp14985 {
tmp14896 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres14868 Obj

if True == tmp14896 {
tmp14869 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__13393 := tmp14869
_ = let__13393

tmp14895 := GConsP(__e, let__13393)

var ifres14870 Obj

if True == tmp14895 {
tmp14871 := GHd(__e, let__13393)

tmp14872 := Call2(__e, SymFn(symshen_4lazyderef), tmp14871, V5020)


let__13394 := tmp14872
_ = let__13394

tmp14894 := GEq(__e, let__13394, symtype)

var ifres14873 Obj

if True == tmp14894 {
tmp14874 := GTl(__e, let__13393)

tmp14875 := Call2(__e, SymFn(symshen_4lazyderef), tmp14874, V5020)


let__13395 := tmp14875
_ = let__13395

tmp14893 := GConsP(__e, let__13395)

var ifres14876 Obj

if True == tmp14893 {
tmp14877 := GHd(__e, let__13395)

let__13396 := tmp14877
_ = let__13396

tmp14878 := GTl(__e, let__13395)

tmp14879 := Call2(__e, SymFn(symshen_4lazyderef), tmp14878, V5020)


let__13397 := tmp14879
_ = let__13397

tmp14892 := GConsP(__e, let__13397)

var ifres14880 Obj

if True == tmp14892 {
tmp14881 := GHd(__e, let__13397)

let__13398 := tmp14881
_ = let__13398

tmp14882 := GTl(__e, let__13397)

tmp14883 := Call2(__e, SymFn(symshen_4lazyderef), tmp14882, V5020)


let__13399 := tmp14883
_ = let__13399

tmp14891 := GEq(__e, let__13399, Nil)

var ifres14884 Obj

if True == tmp14891 {
tmp14885 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14885

tmp14886 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14887 := Call2(__e, SymFn(symshen_4deref), let__13398, V5020)


tmp14888 := Call1(__e, SymFn(symshen_4rectify_1type), tmp14887)


tmp14889 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), let__13396, V5018, V5019, V5020, V5021, let__13228, V5023)
return nil
panic("unreachable")
})

return TailCall6(__e, SymFn(symis_b), tmp14888, V5018, V5020, V5021, let__13228, tmp14889)


panic("unreachable")
})

tmp14890 := Call4(__e, SymFn(symshen_4cut), V5020, V5021, let__13228, tmp14886)


ifres14884 = tmp14890


} else {
ifres14884 = False


}

ifres14880 = ifres14884


} else {
ifres14880 = False


}

ifres14876 = ifres14880


} else {
ifres14876 = False


}

ifres14873 = ifres14876


} else {
ifres14873 = False


}

ifres14870 = ifres14873


} else {
ifres14870 = False


}

ifres14868 = ifres14870


} else {
ifres14868 = False


}

let__13392 := ifres14868
_ = let__13392

tmp14983 := GEq(__e, let__13392, False)

if True == tmp14983 {
tmp14927 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres14897 Obj

if True == tmp14927 {
tmp14898 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__13401 := tmp14898
_ = let__13401

tmp14926 := GConsP(__e, let__13401)

var ifres14899 Obj

if True == tmp14926 {
tmp14900 := GHd(__e, let__13401)

tmp14901 := Call2(__e, SymFn(symshen_4lazyderef), tmp14900, V5020)


let__13402 := tmp14901
_ = let__13402

tmp14925 := GEq(__e, let__13402, symshen_4input_1h_7)

var ifres14902 Obj

if True == tmp14925 {
tmp14903 := GTl(__e, let__13401)

tmp14904 := Call2(__e, SymFn(symshen_4lazyderef), tmp14903, V5020)


let__13403 := tmp14904
_ = let__13403

tmp14924 := GConsP(__e, let__13403)

var ifres14905 Obj

if True == tmp14924 {
tmp14906 := GHd(__e, let__13403)

let__13404 := tmp14906
_ = let__13404

tmp14907 := GTl(__e, let__13403)

tmp14908 := Call2(__e, SymFn(symshen_4lazyderef), tmp14907, V5020)


let__13405 := tmp14908
_ = let__13405

tmp14923 := GConsP(__e, let__13405)

var ifres14909 Obj

if True == tmp14923 {
tmp14910 := GHd(__e, let__13405)

let__13406 := tmp14910
_ = let__13406

tmp14911 := GTl(__e, let__13405)

tmp14912 := Call2(__e, SymFn(symshen_4lazyderef), tmp14911, V5020)


let__13407 := tmp14912
_ = let__13407

tmp14922 := GEq(__e, let__13407, Nil)

var ifres14913 Obj

if True == tmp14922 {
tmp14914 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14914

tmp14915 := Call2(__e, SymFn(symshen_4deref), let__13404, V5020)


tmp14916 := Call1(__e, SymFn(symshen_4rdecons), tmp14915)


tmp14917 := Call1(__e, SymFn(symshen_4rectify_1type), tmp14916)


tmp14918 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14919 := GCons(__e, symin, Nil)

tmp14920 := GCons(__e, symstream, tmp14919)

__e.TailApply(SymFn(symshen_4system_1S_1h), let__13406, tmp14920, V5019, V5020, V5021, let__13228, V5023)
return nil


panic("unreachable")
})

tmp14921 := Call6(__e, SymFn(symis), V5018, tmp14917, V5020, V5021, let__13228, tmp14918)


ifres14913 = tmp14921


} else {
ifres14913 = False


}

ifres14909 = ifres14913


} else {
ifres14909 = False


}

ifres14905 = ifres14909


} else {
ifres14905 = False


}

ifres14902 = ifres14905


} else {
ifres14902 = False


}

ifres14899 = ifres14902


} else {
ifres14899 = False


}

ifres14897 = ifres14899


} else {
ifres14897 = False


}

let__13400 := ifres14897
_ = let__13400

tmp14981 := GEq(__e, let__13400, False)

if True == tmp14981 {
tmp14956 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres14928 Obj

if True == tmp14956 {
tmp14929 := Call2(__e, SymFn(symshen_4lazyderef), V5017, V5020)


let__13409 := tmp14929
_ = let__13409

tmp14955 := GConsP(__e, let__13409)

var ifres14930 Obj

if True == tmp14955 {
tmp14931 := GHd(__e, let__13409)

tmp14932 := Call2(__e, SymFn(symshen_4lazyderef), tmp14931, V5020)


let__13410 := tmp14932
_ = let__13410

tmp14954 := GEq(__e, let__13410, symset)

var ifres14933 Obj

if True == tmp14954 {
tmp14934 := GTl(__e, let__13409)

tmp14935 := Call2(__e, SymFn(symshen_4lazyderef), tmp14934, V5020)


let__13411 := tmp14935
_ = let__13411

tmp14953 := GConsP(__e, let__13411)

var ifres14936 Obj

if True == tmp14953 {
tmp14937 := GHd(__e, let__13411)

let__13412 := tmp14937
_ = let__13412

tmp14938 := GTl(__e, let__13411)

tmp14939 := Call2(__e, SymFn(symshen_4lazyderef), tmp14938, V5020)


let__13413 := tmp14939
_ = let__13413

tmp14952 := GConsP(__e, let__13413)

var ifres14940 Obj

if True == tmp14952 {
tmp14941 := GHd(__e, let__13413)

let__13414 := tmp14941
_ = let__13414

tmp14942 := GTl(__e, let__13413)

tmp14943 := Call2(__e, SymFn(symshen_4lazyderef), tmp14942, V5020)


let__13415 := tmp14943
_ = let__13415

tmp14951 := GEq(__e, let__13415, Nil)

var ifres14944 Obj

if True == tmp14951 {
tmp14945 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14945

tmp14946 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14947 := GCons(__e, let__13412, Nil)

tmp14948 := GCons(__e, symvalue, tmp14947)

tmp14949 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), let__13414, V5018, V5019, V5020, V5021, let__13228, V5023)
return nil
panic("unreachable")
})

__e.TailApply(SymFn(symshen_4system_1S_1h), tmp14948, V5018, V5019, V5020, V5021, let__13228, tmp14949)
return nil


panic("unreachable")
})

tmp14950 := Call(__e, SymFn(symshen_4system_1S_1h), let__13412, symsymbol, V5019, V5020, V5021, let__13228, tmp14946)


ifres14944 = tmp14950


} else {
ifres14944 = False


}

ifres14940 = ifres14944


} else {
ifres14940 = False


}

ifres14936 = ifres14940


} else {
ifres14936 = False


}

ifres14933 = ifres14936


} else {
ifres14933 = False


}

ifres14930 = ifres14933


} else {
ifres14930 = False


}

ifres14928 = ifres14930


} else {
ifres14928 = False


}

let__13408 := ifres14928
_ = let__13408

tmp14979 := GEq(__e, let__13408, False)

if True == tmp14979 {
tmp14964 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres14957 Obj

if True == tmp14964 {
tmp14958 := Call1(__e, SymFn(symshen_4newpv), V5020)


let__13417 := tmp14958
_ = let__13417

tmp14959 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14959

tmp14960 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14961 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), V5017, V5018, let__13417, V5020, V5021, let__13228, V5023)
return nil
panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5020, V5021, let__13228, tmp14961)


panic("unreachable")
})

tmp14962 := Call(__e, SymFn(symshen_4l_1rules), V5019, let__13417, False, V5020, V5021, let__13228, tmp14960)


tmp14963 := Call2(__e, SymFn(symshen_4gc), V5020, tmp14962)


ifres14957 = tmp14963


} else {
ifres14957 = False


}

let__13416 := ifres14957
_ = let__13416

tmp14977 := GEq(__e, let__13416, False)

if True == tmp14977 {
tmp14973 := Call1(__e, SymFn(symshen_4unlocked_2), V5021)


var ifres14965 Obj

if True == tmp14973 {
tmp14966 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14966

tmp14967 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp14968 := GCons(__e, V5018, Nil)

tmp14969 := GCons(__e, tmp14967, tmp14968)

tmp14970 := GCons(__e, V5017, tmp14969)

tmp14971 := GValue(__e, symshen_4_ddatatypes_d)

tmp14972 := Call(__e, SymFn(symshen_4search_1user_1datatypes), tmp14970, V5019, tmp14971, V5020, V5021, let__13228, V5023)


ifres14965 = tmp14972


} else {
ifres14965 = False


}

let__13418 := ifres14965
_ = let__13418

tmp14975 := GEq(__e, let__13418, False)

if True == tmp14975 {
return TailCall2(__e, SymFn(symshen_4unlock), V5021, let__13228)
} else {
return let__13418
}


} else {
return let__13416
}


} else {
return let__13408
}


} else {
return let__13400
}


} else {
return let__13392
}


} else {
return let__13373
}


} else {
return let__13360
}


} else {
return let__13334
}


} else {
return let__13324
}


} else {
return let__13305
}


} else {
return let__13281
}


} else {
return let__13262
}


} else {
return let__13255
}


} else {
return let__13248
}


} else {
return let__13242
}


} else {
return let__13236
}


} else {
return let__13232
}


} else {
return let__13231
}


} else {
return let__13230
}


} else {
return let__13229
}


panic("unreachable")
}())
}, 7)

tmp15014 := Call2(__e, ns2_1set, symshen_4system_1S_1h, tmp14263)


_ = tmp15014

tmp15015 := MakeNative1(func(__e *ControlFlow, V5229 Obj) Obj {
tmp15045 := GConsP(__e, V5229)

var ifres15030 Obj

if True == tmp15045 {
tmp15043 := GHd(__e, V5229)

tmp15044 := GEq(__e, symcons, tmp15043)

var ifres15031 Obj

if True == tmp15044 {
tmp15041 := GTl(__e, V5229)

tmp15042 := GConsP(__e, tmp15041)

var ifres15032 Obj

if True == tmp15042 {
tmp15038 := GTl(__e, V5229)

tmp15039 := GTl(__e, tmp15038)

tmp15040 := GConsP(__e, tmp15039)

var ifres15033 Obj

if True == tmp15040 {
tmp15034 := GTl(__e, V5229)

tmp15035 := GTl(__e, tmp15034)

tmp15036 := GTl(__e, tmp15035)

tmp15037 := GEq(__e, Nil, tmp15036)

ifres15033 = tmp15037


} else {
ifres15033 = False


}

ifres15032 = ifres15033


} else {
ifres15032 = False


}

ifres15031 = ifres15032


} else {
ifres15031 = False


}

ifres15030 = ifres15031


} else {
ifres15030 = False


}

if True == ifres15030 {
tmp15016 := GTl(__e, V5229)

tmp15017 := GHd(__e, tmp15016)

tmp15018 := Call1(__e, SymFn(symshen_4rdecons), tmp15017)


tmp15019 := GTl(__e, V5229)

tmp15020 := GTl(__e, tmp15019)

tmp15021 := GHd(__e, tmp15020)

tmp15022 := Call1(__e, SymFn(symshen_4rdecons), tmp15021)


return GCons(__e, tmp15018, tmp15022)


} else {
tmp15028 := GConsP(__e, V5229)

if True == tmp15028 {
tmp15023 := GHd(__e, V5229)

tmp15024 := Call1(__e, SymFn(symshen_4rdecons), tmp15023)


tmp15025 := GTl(__e, V5229)

tmp15026 := Call1(__e, SymFn(symshen_4rdecons), tmp15025)


return GCons(__e, tmp15024, tmp15026)


} else {
return V5229
}


}


panic("unreachable")
})

tmp15046 := Call2(__e, ns2_1set, symshen_4rdecons, tmp15015)


_ = tmp15046

tmp15047 := MakeNative6(func(__e *ControlFlow, V5230, V5231, V5232, V5233, V5234, V5235 Obj) Obj {
tmp15060 := Call1(__e, SymFn(symshen_4unlocked_2), V5233)


var ifres15048 Obj

if True == tmp15060 {
tmp15049 := Call2(__e, SymFn(symshen_4lazyderef), V5231, V5232)


let__13420 := tmp15049
_ = let__13420

tmp15050 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15051 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15051

tmp15052 := Call2(__e, SymFn(symshen_4lazyderef), V5230, V5232)


tmp15053 := PrimGuard1(__e, symnumber_2, PrimIsNumber, tmp15052)

return TailCall5(__e, SymFn(symwhen), tmp15053, V5232, V5233, V5234, V5235)


panic("unreachable")
})

let__13421 := tmp15050
_ = let__13421

tmp15059 := GEq(__e, let__13420, symnumber)

var ifres15054 Obj

if True == tmp15059 {
tmp15055 := Call1(__e, SymFn(symthaw), let__13421)


ifres15054 = tmp15055


} else {
tmp15058 := Call1(__e, SymFn(symshen_4pvar_2), let__13420)


var ifres15056 Obj

if True == tmp15058 {
tmp15057 := Call4(__e, SymFn(symshen_4bind_b), let__13420, symnumber, V5232, let__13421)


ifres15056 = tmp15057


} else {
ifres15056 = False


}

ifres15054 = ifres15056


}

ifres15048 = ifres15054


} else {
ifres15048 = False


}

let__13419 := ifres15048
_ = let__13419

tmp15150 := GEq(__e, let__13419, False)

if True == tmp15150 {
tmp15073 := Call1(__e, SymFn(symshen_4unlocked_2), V5233)


var ifres15061 Obj

if True == tmp15073 {
tmp15062 := Call2(__e, SymFn(symshen_4lazyderef), V5231, V5232)


let__13423 := tmp15062
_ = let__13423

tmp15063 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15064 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15064

tmp15065 := Call2(__e, SymFn(symshen_4lazyderef), V5230, V5232)


tmp15066 := Call1(__e, SymFn(symboolean_2), tmp15065)


return TailCall5(__e, SymFn(symwhen), tmp15066, V5232, V5233, V5234, V5235)


panic("unreachable")
})

let__13424 := tmp15063
_ = let__13424

tmp15072 := GEq(__e, let__13423, symboolean)

var ifres15067 Obj

if True == tmp15072 {
tmp15068 := Call1(__e, SymFn(symthaw), let__13424)


ifres15067 = tmp15068


} else {
tmp15071 := Call1(__e, SymFn(symshen_4pvar_2), let__13423)


var ifres15069 Obj

if True == tmp15071 {
tmp15070 := Call4(__e, SymFn(symshen_4bind_b), let__13423, symboolean, V5232, let__13424)


ifres15069 = tmp15070


} else {
ifres15069 = False


}

ifres15067 = ifres15069


}

ifres15061 = ifres15067


} else {
ifres15061 = False


}

let__13422 := ifres15061
_ = let__13422

tmp15148 := GEq(__e, let__13422, False)

if True == tmp15148 {
tmp15086 := Call1(__e, SymFn(symshen_4unlocked_2), V5233)


var ifres15074 Obj

if True == tmp15086 {
tmp15075 := Call2(__e, SymFn(symshen_4lazyderef), V5231, V5232)


let__13426 := tmp15075
_ = let__13426

tmp15076 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15077 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15077

tmp15078 := Call2(__e, SymFn(symshen_4lazyderef), V5230, V5232)


tmp15079 := PrimGuard1(__e, symstring_2, PrimIsString, tmp15078)

return TailCall5(__e, SymFn(symwhen), tmp15079, V5232, V5233, V5234, V5235)


panic("unreachable")
})

let__13427 := tmp15076
_ = let__13427

tmp15085 := GEq(__e, let__13426, symstring)

var ifres15080 Obj

if True == tmp15085 {
tmp15081 := Call1(__e, SymFn(symthaw), let__13427)


ifres15080 = tmp15081


} else {
tmp15084 := Call1(__e, SymFn(symshen_4pvar_2), let__13426)


var ifres15082 Obj

if True == tmp15084 {
tmp15083 := Call4(__e, SymFn(symshen_4bind_b), let__13426, symstring, V5232, let__13427)


ifres15082 = tmp15083


} else {
ifres15082 = False


}

ifres15080 = ifres15082


}

ifres15074 = ifres15080


} else {
ifres15074 = False


}

let__13425 := ifres15074
_ = let__13425

tmp15146 := GEq(__e, let__13425, False)

if True == tmp15146 {
tmp15099 := Call1(__e, SymFn(symshen_4unlocked_2), V5233)


var ifres15087 Obj

if True == tmp15099 {
tmp15088 := Call2(__e, SymFn(symshen_4lazyderef), V5231, V5232)


let__13429 := tmp15088
_ = let__13429

tmp15089 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15090 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15090

tmp15091 := Call2(__e, SymFn(symshen_4lazyderef), V5230, V5232)


tmp15092 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, tmp15091)

return TailCall5(__e, SymFn(symwhen), tmp15092, V5232, V5233, V5234, V5235)


panic("unreachable")
})

let__13430 := tmp15089
_ = let__13430

tmp15098 := GEq(__e, let__13429, symsymbol)

var ifres15093 Obj

if True == tmp15098 {
tmp15094 := Call1(__e, SymFn(symthaw), let__13430)


ifres15093 = tmp15094


} else {
tmp15097 := Call1(__e, SymFn(symshen_4pvar_2), let__13429)


var ifres15095 Obj

if True == tmp15097 {
tmp15096 := Call4(__e, SymFn(symshen_4bind_b), let__13429, symsymbol, V5232, let__13430)


ifres15095 = tmp15096


} else {
ifres15095 = False


}

ifres15093 = ifres15095


}

ifres15087 = ifres15093


} else {
ifres15087 = False


}

let__13428 := ifres15087
_ = let__13428

tmp15144 := GEq(__e, let__13428, False)

if True == tmp15144 {
tmp15142 := Call1(__e, SymFn(symshen_4unlocked_2), V5233)


if True == tmp15142 {
tmp15100 := Call2(__e, SymFn(symshen_4lazyderef), V5230, V5232)


let__13431 := tmp15100
_ = let__13431

tmp15140 := GEq(__e, let__13431, Nil)

if True == tmp15140 {
tmp15101 := Call2(__e, SymFn(symshen_4lazyderef), V5231, V5232)


let__13432 := tmp15101
_ = let__13432

tmp15102 := MakeNative1(func(__e *ControlFlow, Z5251 Obj) Obj {
tmp15103 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15103

return TailCall1(__e, SymFn(symthaw), V5235)


panic("unreachable")
})

let__13433 := tmp15102
_ = let__13433

tmp15138 := GConsP(__e, let__13432)

if True == tmp15138 {
tmp15104 := GHd(__e, let__13432)

tmp15105 := Call2(__e, SymFn(symshen_4lazyderef), tmp15104, V5232)


let__13434 := tmp15105
_ = let__13434

tmp15106 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15107 := GTl(__e, let__13432)

tmp15108 := Call2(__e, SymFn(symshen_4lazyderef), tmp15107, V5232)


let__13436 := tmp15108
_ = let__13436

tmp15109 := MakeNative1(func(__e *ControlFlow, Z5256 Obj) Obj {
return TailCall1(__e, let__13433, Z5256)
panic("unreachable")
})

let__13437 := tmp15109
_ = let__13437

tmp15125 := GConsP(__e, let__13436)

if True == tmp15125 {
tmp15110 := GHd(__e, let__13436)

let__13438 := tmp15110
_ = let__13438

tmp15111 := GTl(__e, let__13436)

tmp15112 := Call2(__e, SymFn(symshen_4lazyderef), tmp15111, V5232)


let__13439 := tmp15112
_ = let__13439

tmp15113 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13437, let__13438)
panic("unreachable")
})

let__13440 := tmp15113
_ = let__13440

tmp15117 := GEq(__e, let__13439, Nil)

if True == tmp15117 {
return TailCall1(__e, SymFn(symthaw), let__13440)
} else {
tmp15115 := Call1(__e, SymFn(symshen_4pvar_2), let__13439)


if True == tmp15115 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__13439, Nil, V5232, let__13440)
} else {
return False
}


}


} else {
tmp15123 := Call1(__e, SymFn(symshen_4pvar_2), let__13436)


if True == tmp15123 {
tmp15118 := Call1(__e, SymFn(symshen_4newpv), V5232)


let__13441 := tmp15118
_ = let__13441

tmp15119 := GCons(__e, let__13441, Nil)

tmp15120 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13437, let__13441)
panic("unreachable")
})

tmp15121 := Call4(__e, SymFn(symshen_4bind_b), let__13436, tmp15119, V5232, tmp15120)


return TailCall2(__e, SymFn(symshen_4gc), V5232, tmp15121)


} else {
return False
}


}


panic("unreachable")
})

let__13435 := tmp15106
_ = let__13435

tmp15129 := GEq(__e, let__13434, symlist)

if True == tmp15129 {
return TailCall1(__e, SymFn(symthaw), let__13435)
} else {
tmp15127 := Call1(__e, SymFn(symshen_4pvar_2), let__13434)


if True == tmp15127 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__13434, symlist, V5232, let__13435)
} else {
return False
}


}


} else {
tmp15136 := Call1(__e, SymFn(symshen_4pvar_2), let__13432)


if True == tmp15136 {
tmp15130 := Call1(__e, SymFn(symshen_4newpv), V5232)


let__13442 := tmp15130
_ = let__13442

tmp15131 := GCons(__e, let__13442, Nil)

tmp15132 := GCons(__e, symlist, tmp15131)

tmp15133 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13433, let__13442)
panic("unreachable")
})

tmp15134 := Call4(__e, SymFn(symshen_4bind_b), let__13432, tmp15132, V5232, tmp15133)


return TailCall2(__e, SymFn(symshen_4gc), V5232, tmp15134)


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
return let__13428
}


} else {
return let__13425
}


} else {
return let__13422
}


} else {
return let__13419
}


panic("unreachable")
})

tmp15151 := Call2(__e, ns2_1set, symshen_4primitive, tmp15047)


_ = tmp15151

tmp15152 := MakeNative(func(__e *ControlFlow) {
V5262 := __e.Get(1)
V5263 := __e.Get(2)
V5264 := __e.Get(3)
V5265 := __e.Get(4)
V5266 := __e.Get(5)
V5267 := __e.Get(6)
V5268 := __e.Get(7)
__e.ReturnDirect(func() Obj {
tmp15155 := Call2(__e, SymFn(symshen_4lazyderef), V5264, V5265)


tmp15156 := GEq(__e, tmp15155, Nil)

if True == tmp15156 {
return False
} else {
tmp15153 := Call2(__e, SymFn(symshen_4deref), V5262, V5265)


__e.TailApply(SymFn(symshen_4by_1hypothesis_1h), tmp15153, V5263, V5264, V5265, V5266, V5267, V5268)
return nil


}


panic("unreachable")
}())
}, 7)

tmp15157 := Call2(__e, ns2_1set, symshen_4by_1hypothesis, tmp15152)


_ = tmp15157

tmp15158 := MakeNative(func(__e *ControlFlow) {
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
tmp15191 := Call1(__e, SymFn(symshen_4unlocked_2), V5273)


var ifres15159 Obj

if True == tmp15191 {
tmp15160 := Call2(__e, SymFn(symshen_4lazyderef), V5271, V5272)


let__13444 := tmp15160
_ = let__13444

tmp15190 := GConsP(__e, let__13444)

var ifres15161 Obj

if True == tmp15190 {
tmp15162 := GHd(__e, let__13444)

tmp15163 := Call2(__e, SymFn(symshen_4lazyderef), tmp15162, V5272)


let__13445 := tmp15163
_ = let__13445

tmp15189 := GConsP(__e, let__13445)

var ifres15164 Obj

if True == tmp15189 {
tmp15165 := GHd(__e, let__13445)

let__13446 := tmp15165
_ = let__13446

tmp15166 := GTl(__e, let__13445)

tmp15167 := Call2(__e, SymFn(symshen_4lazyderef), tmp15166, V5272)


let__13447 := tmp15167
_ = let__13447

tmp15188 := GConsP(__e, let__13447)

var ifres15168 Obj

if True == tmp15188 {
tmp15169 := GHd(__e, let__13447)

let__13448 := tmp15169
_ = let__13448

tmp15170 := GTl(__e, let__13447)

tmp15171 := Call2(__e, SymFn(symshen_4lazyderef), tmp15170, V5272)


let__13449 := tmp15171
_ = let__13449

tmp15187 := GConsP(__e, let__13449)

var ifres15172 Obj

if True == tmp15187 {
tmp15173 := GHd(__e, let__13449)

let__13450 := tmp15173
_ = let__13450

tmp15174 := GTl(__e, let__13449)

tmp15175 := Call2(__e, SymFn(symshen_4lazyderef), tmp15174, V5272)


let__13451 := tmp15175
_ = let__13451

tmp15186 := GEq(__e, let__13451, Nil)

var ifres15176 Obj

if True == tmp15186 {
tmp15177 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15177

tmp15178 := Call2(__e, SymFn(symshen_4deref), let__13448, V5272)


tmp15179 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp15180 := GEq(__e, tmp15178, tmp15179)

tmp15181 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15182 := Call2(__e, SymFn(symshen_4deref), let__13446, V5272)


tmp15183 := GEq(__e, V5269, tmp15182)

tmp15184 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall6(__e, SymFn(symis_b), V5270, let__13450, V5272, V5273, V5274, V5275)
panic("unreachable")
})

return TailCall5(__e, SymFn(symwhen), tmp15183, V5272, V5273, V5274, tmp15184)


panic("unreachable")
})

tmp15185 := Call5(__e, SymFn(symwhen), tmp15180, V5272, V5273, V5274, tmp15181)


ifres15176 = tmp15185


} else {
ifres15176 = False


}

ifres15172 = ifres15176


} else {
ifres15172 = False


}

ifres15168 = ifres15172


} else {
ifres15168 = False


}

ifres15164 = ifres15168


} else {
ifres15164 = False


}

ifres15161 = ifres15164


} else {
ifres15161 = False


}

ifres15159 = ifres15161


} else {
ifres15159 = False


}

let__13443 := ifres15159
_ = let__13443

tmp15200 := GEq(__e, let__13443, False)

if True == tmp15200 {
tmp15198 := Call1(__e, SymFn(symshen_4unlocked_2), V5273)


if True == tmp15198 {
tmp15192 := Call2(__e, SymFn(symshen_4lazyderef), V5271, V5272)


let__13452 := tmp15192
_ = let__13452

tmp15196 := GConsP(__e, let__13452)

if True == tmp15196 {
tmp15193 := GTl(__e, let__13452)

let__13453 := tmp15193
_ = let__13453

tmp15194 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15194

if SymFn(symshen_4by_1hypothesis_1h) == __self {
__self1, __self2, __self3, __self4, __self5, __self6, __self7 = V5269, V5270, let__13453, V5272, V5273, V5274, V5275
__e.Tick()
goto __selftop
}
__e.TailApply(SymFn(symshen_4by_1hypothesis_1h), V5269, V5270, let__13453, V5272, V5273, V5274, V5275)
return nil


} else {
return False
}


} else {
return False
}


} else {
return let__13443
}


panic("unreachable")
}())
}, 7)

tmp15201 := Call2(__e, ns2_1set, symshen_4by_1hypothesis_1h, tmp15158)


_ = tmp15201

tmp15202 := MakeNative6(func(__e *ControlFlow, V5287, V5288, V5289, V5290, V5291, V5292 Obj) Obj {
tmp15207 := Call1(__e, SymFn(symshen_4unlocked_2), V5290)


if True == tmp15207 {
tmp15203 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15203

tmp15204 := GValue(__e, symshen_4_dsigf_d)

tmp15205 := Call2(__e, SymFn(symassoc), V5287, tmp15204)


return TailCall6(__e, SymFn(symshen_4sigf), tmp15205, V5288, V5289, V5290, V5291, V5292)


} else {
return False
}


panic("unreachable")
})

tmp15208 := Call2(__e, ns2_1set, symshen_4lookupsig, tmp15202)


_ = tmp15208

tmp15209 := MakeNative6(func(__e *ControlFlow, V5307, V5308, V5309, V5310, V5311, V5312 Obj) Obj {
tmp15216 := GConsP(__e, V5307)

if True == tmp15216 {
tmp15210 := GTl(__e, V5307)

tmp15211 := Call1(__e, tmp15210, V5308)


tmp15212 := Call1(__e, tmp15211, V5309)


tmp15213 := Call1(__e, tmp15212, V5310)


tmp15214 := Call1(__e, tmp15213, V5311)


return TailCall1(__e, tmp15214, V5312)


} else {
return False
}


panic("unreachable")
})

tmp15217 := Call2(__e, ns2_1set, symshen_4sigf, tmp15209)


_ = tmp15217

tmp15218 := MakeNative1(func(__e *ControlFlow, V5313 Obj) Obj {
tmp15219 := PrimGuard1(__e, symabsvector, PrimAbsvector, MakeInteger(3))

let__13454 := tmp15219
_ = let__13454

tmp15220 := GVecSet(__e, let__13454, MakeInteger(0), symshen_4print_1freshterm)

let__13455 := tmp15220
_ = let__13455

tmp15221 := GVecSet(__e, let__13455, MakeInteger(1), V5313)

let__13456 := tmp15221
_ = let__13456

tmp15222 := GValue(__e, symshen_4_dgensym_d)

tmp15224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dgensym_d, GAdd(__e, MakeInteger(1), tmp15222))
}
__typedArg0 := symshen_4_dgensym_d
__typedArg1 := GAdd(__e, MakeInteger(1), tmp15222)
return Call2(__e, SymFn(symset), __typedArg0, __typedArg1)
})()

tmp15225 := GVecSet(__e, let__13456, MakeInteger(2), tmp15224)

let__13457 := tmp15225
_ = let__13457

return let__13457


panic("unreachable")
})

tmp15226 := Call2(__e, ns2_1set, symshen_4freshterm, tmp15218)


_ = tmp15226

tmp15227 := MakeNative1(func(__e *ControlFlow, V5318 Obj) Obj {
tmp15228 := GVecGet(__e, V5318, MakeInteger(1))

tmp15229 := Call3(__e, SymFn(symshen_4app), tmp15228, MakeString(""), symshen_4a)


return PrimGuard2(__e, symcn, PrimStringConcat, MakeString("&&"), tmp15229)


panic("unreachable")
})

tmp15230 := Call2(__e, ns2_1set, symshen_4print_1freshterm, tmp15227)


_ = tmp15230

tmp15231 := MakeNative(func(__e *ControlFlow) {
V5319 := __e.Get(1)
V5320 := __e.Get(2)
V5321 := __e.Get(3)
V5322 := __e.Get(4)
V5323 := __e.Get(5)
V5324 := __e.Get(6)
V5325 := __e.Get(7)
__e.ReturnDirect(func() Obj {
tmp15235 := Call2(__e, SymFn(symshen_4lazyderef), V5321, V5322)


tmp15236 := GEq(__e, tmp15235, Nil)

if True == tmp15236 {
return False
} else {
tmp15232 := Call2(__e, SymFn(symshen_4deref), V5319, V5322)


tmp15233 := Call2(__e, SymFn(symshen_4deref), V5320, V5322)


__e.TailApply(SymFn(symshen_4search_1user_1datatypes_1h), tmp15232, tmp15233, V5321, V5322, V5323, V5324, V5325)
return nil


}


panic("unreachable")
}())
}, 7)

tmp15237 := Call2(__e, ns2_1set, symshen_4search_1user_1datatypes, tmp15231)


_ = tmp15237

tmp15238 := MakeNative(func(__e *ControlFlow) {
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
tmp15253 := Call1(__e, SymFn(symshen_4unlocked_2), V5330)


var ifres15239 Obj

if True == tmp15253 {
tmp15240 := Call2(__e, SymFn(symshen_4lazyderef), V5328, V5329)


let__13459 := tmp15240
_ = let__13459

tmp15252 := GConsP(__e, let__13459)

var ifres15241 Obj

if True == tmp15252 {
tmp15242 := GHd(__e, let__13459)

tmp15243 := Call2(__e, SymFn(symshen_4lazyderef), tmp15242, V5329)


let__13460 := tmp15243
_ = let__13460

tmp15251 := GConsP(__e, let__13460)

var ifres15244 Obj

if True == tmp15251 {
tmp15245 := GTl(__e, let__13460)

let__13461 := tmp15245
_ = let__13461

tmp15246 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15246

tmp15247 := Call2(__e, SymFn(symshen_4deref), let__13461, V5329)


tmp15248 := Call1(__e, tmp15247, V5326)


tmp15249 := Call1(__e, tmp15248, V5327)


tmp15250 := Call5(__e, SymFn(symcall), tmp15249, V5329, V5330, V5331, V5332)


ifres15244 = tmp15250


} else {
ifres15244 = False


}

ifres15241 = ifres15244


} else {
ifres15241 = False


}

ifres15239 = ifres15241


} else {
ifres15239 = False


}

let__13458 := ifres15239
_ = let__13458

tmp15262 := GEq(__e, let__13458, False)

if True == tmp15262 {
tmp15260 := Call1(__e, SymFn(symshen_4unlocked_2), V5330)


if True == tmp15260 {
tmp15254 := Call2(__e, SymFn(symshen_4lazyderef), V5328, V5329)


let__13462 := tmp15254
_ = let__13462

tmp15258 := GConsP(__e, let__13462)

if True == tmp15258 {
tmp15255 := GTl(__e, let__13462)

let__13463 := tmp15255
_ = let__13463

tmp15256 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15256

if SymFn(symshen_4search_1user_1datatypes_1h) == __self {
__self1, __self2, __self3, __self4, __self5, __self6, __self7 = V5326, V5327, let__13463, V5329, V5330, V5331, V5332
__e.Tick()
goto __selftop
}
__e.TailApply(SymFn(symshen_4search_1user_1datatypes_1h), V5326, V5327, let__13463, V5329, V5330, V5331, V5332)
return nil


} else {
return False
}


} else {
return False
}


} else {
return let__13458
}


panic("unreachable")
}())
}, 7)

tmp15263 := Call2(__e, ns2_1set, symshen_4search_1user_1datatypes_1h, tmp15238)


_ = tmp15263

tmp15264 := MakeNative(func(__e *ControlFlow) {
V5339 := __e.Get(1)
V5340 := __e.Get(2)
V5341 := __e.Get(3)
V5342 := __e.Get(4)
V5343 := __e.Get(5)
V5344 := __e.Get(6)
V5345 := __e.Get(7)
__e.ReturnDirect(func() Obj {
let__13464 := GAdd(__e, V5344, MakeInteger(1))
_ = let__13464

tmp15276 := Call1(__e, SymFn(symshen_4unlocked_2), V5343)


var ifres15266 Obj

if True == tmp15276 {
tmp15267 := Call2(__e, SymFn(symshen_4lazyderef), V5339, V5342)


let__13466 := tmp15267
_ = let__13466

tmp15275 := GEq(__e, let__13466, Nil)

var ifres15268 Obj

if True == tmp15275 {
tmp15269 := Call2(__e, SymFn(symshen_4lazyderef), V5341, V5342)


let__13467 := tmp15269
_ = let__13467

tmp15274 := GEq(__e, let__13467, True)

var ifres15270 Obj

if True == tmp15274 {
tmp15271 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15271

tmp15272 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall6(__e, SymFn(symbind), V5340, Nil, V5342, V5343, let__13464, V5345)
panic("unreachable")
})

tmp15273 := Call4(__e, SymFn(symshen_4cut), V5342, V5343, let__13464, tmp15272)


ifres15270 = tmp15273


} else {
ifres15270 = False


}

ifres15268 = ifres15270


} else {
ifres15268 = False


}

ifres15266 = ifres15268


} else {
ifres15266 = False


}

let__13465 := ifres15266
_ = let__13465

tmp15619 := GEq(__e, let__13465, False)

if True == tmp15619 {
tmp15355 := Call1(__e, SymFn(symshen_4unlocked_2), V5343)


var ifres15277 Obj

if True == tmp15355 {
tmp15278 := Call2(__e, SymFn(symshen_4lazyderef), V5339, V5342)


let__13469 := tmp15278
_ = let__13469

tmp15354 := GConsP(__e, let__13469)

var ifres15279 Obj

if True == tmp15354 {
tmp15280 := GHd(__e, let__13469)

tmp15281 := Call2(__e, SymFn(symshen_4lazyderef), tmp15280, V5342)


let__13470 := tmp15281
_ = let__13470

tmp15353 := GConsP(__e, let__13470)

var ifres15282 Obj

if True == tmp15353 {
tmp15283 := GHd(__e, let__13470)

tmp15284 := Call2(__e, SymFn(symshen_4lazyderef), tmp15283, V5342)


let__13471 := tmp15284
_ = let__13471

tmp15352 := GConsP(__e, let__13471)

var ifres15285 Obj

if True == tmp15352 {
tmp15286 := GHd(__e, let__13471)

tmp15287 := Call2(__e, SymFn(symshen_4lazyderef), tmp15286, V5342)


let__13472 := tmp15287
_ = let__13472

tmp15351 := GEq(__e, let__13472, symcons)

var ifres15288 Obj

if True == tmp15351 {
tmp15289 := GTl(__e, let__13471)

tmp15290 := Call2(__e, SymFn(symshen_4lazyderef), tmp15289, V5342)


let__13473 := tmp15290
_ = let__13473

tmp15350 := GConsP(__e, let__13473)

var ifres15291 Obj

if True == tmp15350 {
tmp15292 := GHd(__e, let__13473)

let__13474 := tmp15292
_ = let__13474

tmp15293 := GTl(__e, let__13473)

tmp15294 := Call2(__e, SymFn(symshen_4lazyderef), tmp15293, V5342)


let__13475 := tmp15294
_ = let__13475

tmp15349 := GConsP(__e, let__13475)

var ifres15295 Obj

if True == tmp15349 {
tmp15296 := GHd(__e, let__13475)

let__13476 := tmp15296
_ = let__13476

tmp15297 := GTl(__e, let__13475)

tmp15298 := Call2(__e, SymFn(symshen_4lazyderef), tmp15297, V5342)


let__13477 := tmp15298
_ = let__13477

tmp15348 := GEq(__e, let__13477, Nil)

var ifres15299 Obj

if True == tmp15348 {
tmp15300 := GTl(__e, let__13470)

tmp15301 := Call2(__e, SymFn(symshen_4lazyderef), tmp15300, V5342)


let__13478 := tmp15301
_ = let__13478

tmp15347 := GConsP(__e, let__13478)

var ifres15302 Obj

if True == tmp15347 {
tmp15303 := GHd(__e, let__13478)

let__13479 := tmp15303
_ = let__13479

tmp15304 := GTl(__e, let__13478)

tmp15305 := Call2(__e, SymFn(symshen_4lazyderef), tmp15304, V5342)


let__13480 := tmp15305
_ = let__13480

tmp15346 := GConsP(__e, let__13480)

var ifres15306 Obj

if True == tmp15346 {
tmp15307 := GHd(__e, let__13480)

tmp15308 := Call2(__e, SymFn(symshen_4lazyderef), tmp15307, V5342)


let__13481 := tmp15308
_ = let__13481

tmp15345 := GConsP(__e, let__13481)

var ifres15309 Obj

if True == tmp15345 {
tmp15310 := GHd(__e, let__13481)

tmp15311 := Call2(__e, SymFn(symshen_4lazyderef), tmp15310, V5342)


let__13482 := tmp15311
_ = let__13482

tmp15344 := GEq(__e, let__13482, symlist)

var ifres15312 Obj

if True == tmp15344 {
tmp15313 := GTl(__e, let__13481)

tmp15314 := Call2(__e, SymFn(symshen_4lazyderef), tmp15313, V5342)


let__13483 := tmp15314
_ = let__13483

tmp15343 := GConsP(__e, let__13483)

var ifres15315 Obj

if True == tmp15343 {
tmp15316 := GHd(__e, let__13483)

let__13484 := tmp15316
_ = let__13484

tmp15317 := GTl(__e, let__13483)

tmp15318 := Call2(__e, SymFn(symshen_4lazyderef), tmp15317, V5342)


let__13485 := tmp15318
_ = let__13485

tmp15342 := GEq(__e, let__13485, Nil)

var ifres15319 Obj

if True == tmp15342 {
tmp15320 := GTl(__e, let__13480)

tmp15321 := Call2(__e, SymFn(symshen_4lazyderef), tmp15320, V5342)


let__13486 := tmp15321
_ = let__13486

tmp15341 := GEq(__e, let__13486, Nil)

var ifres15322 Obj

if True == tmp15341 {
tmp15323 := GTl(__e, let__13469)

let__13487 := tmp15323
_ = let__13487

tmp15324 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15324

tmp15325 := Call2(__e, SymFn(symshen_4deref), let__13479, V5342)


tmp15326 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp15327 := GEq(__e, tmp15325, tmp15326)

tmp15328 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15329 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15330 := GCons(__e, let__13484, Nil)

tmp15331 := GCons(__e, let__13479, tmp15330)

tmp15332 := GCons(__e, let__13474, tmp15331)

tmp15333 := GCons(__e, let__13484, Nil)

tmp15334 := GCons(__e, symlist, tmp15333)

tmp15335 := GCons(__e, tmp15334, Nil)

tmp15336 := GCons(__e, let__13479, tmp15335)

tmp15337 := GCons(__e, let__13476, tmp15336)

tmp15338 := GCons(__e, tmp15337, let__13487)

tmp15339 := GCons(__e, tmp15332, tmp15338)

__e.TailApply(SymFn(symshen_4l_1rules), tmp15339, V5340, True, V5342, V5343, let__13464, V5345)
return nil


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5342, V5343, let__13464, tmp15329)


panic("unreachable")
})

tmp15340 := Call5(__e, SymFn(symwhen), tmp15327, V5342, V5343, let__13464, tmp15328)


ifres15322 = tmp15340


} else {
ifres15322 = False


}

ifres15319 = ifres15322


} else {
ifres15319 = False


}

ifres15315 = ifres15319


} else {
ifres15315 = False


}

ifres15312 = ifres15315


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

ifres15302 = ifres15306


} else {
ifres15302 = False


}

ifres15299 = ifres15302


} else {
ifres15299 = False


}

ifres15295 = ifres15299


} else {
ifres15295 = False


}

ifres15291 = ifres15295


} else {
ifres15291 = False


}

ifres15288 = ifres15291


} else {
ifres15288 = False


}

ifres15285 = ifres15288


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

let__13468 := ifres15277
_ = let__13468

tmp15617 := GEq(__e, let__13468, False)

if True == tmp15617 {
tmp15437 := Call1(__e, SymFn(symshen_4unlocked_2), V5343)


var ifres15356 Obj

if True == tmp15437 {
tmp15357 := Call2(__e, SymFn(symshen_4lazyderef), V5339, V5342)


let__13489 := tmp15357
_ = let__13489

tmp15436 := GConsP(__e, let__13489)

var ifres15358 Obj

if True == tmp15436 {
tmp15359 := GHd(__e, let__13489)

tmp15360 := Call2(__e, SymFn(symshen_4lazyderef), tmp15359, V5342)


let__13490 := tmp15360
_ = let__13490

tmp15435 := GConsP(__e, let__13490)

var ifres15361 Obj

if True == tmp15435 {
tmp15362 := GHd(__e, let__13490)

tmp15363 := Call2(__e, SymFn(symshen_4lazyderef), tmp15362, V5342)


let__13491 := tmp15363
_ = let__13491

tmp15434 := GConsP(__e, let__13491)

var ifres15364 Obj

if True == tmp15434 {
tmp15365 := GHd(__e, let__13491)

tmp15366 := Call2(__e, SymFn(symshen_4lazyderef), tmp15365, V5342)


let__13492 := tmp15366
_ = let__13492

tmp15433 := GEq(__e, let__13492, sym_8p)

var ifres15367 Obj

if True == tmp15433 {
tmp15368 := GTl(__e, let__13491)

tmp15369 := Call2(__e, SymFn(symshen_4lazyderef), tmp15368, V5342)


let__13493 := tmp15369
_ = let__13493

tmp15432 := GConsP(__e, let__13493)

var ifres15370 Obj

if True == tmp15432 {
tmp15371 := GHd(__e, let__13493)

let__13494 := tmp15371
_ = let__13494

tmp15372 := GTl(__e, let__13493)

tmp15373 := Call2(__e, SymFn(symshen_4lazyderef), tmp15372, V5342)


let__13495 := tmp15373
_ = let__13495

tmp15431 := GConsP(__e, let__13495)

var ifres15374 Obj

if True == tmp15431 {
tmp15375 := GHd(__e, let__13495)

let__13496 := tmp15375
_ = let__13496

tmp15376 := GTl(__e, let__13495)

tmp15377 := Call2(__e, SymFn(symshen_4lazyderef), tmp15376, V5342)


let__13497 := tmp15377
_ = let__13497

tmp15430 := GEq(__e, let__13497, Nil)

var ifres15378 Obj

if True == tmp15430 {
tmp15379 := GTl(__e, let__13490)

tmp15380 := Call2(__e, SymFn(symshen_4lazyderef), tmp15379, V5342)


let__13498 := tmp15380
_ = let__13498

tmp15429 := GConsP(__e, let__13498)

var ifres15381 Obj

if True == tmp15429 {
tmp15382 := GHd(__e, let__13498)

let__13499 := tmp15382
_ = let__13499

tmp15383 := GTl(__e, let__13498)

tmp15384 := Call2(__e, SymFn(symshen_4lazyderef), tmp15383, V5342)


let__13500 := tmp15384
_ = let__13500

tmp15428 := GConsP(__e, let__13500)

var ifres15385 Obj

if True == tmp15428 {
tmp15386 := GHd(__e, let__13500)

tmp15387 := Call2(__e, SymFn(symshen_4lazyderef), tmp15386, V5342)


let__13501 := tmp15387
_ = let__13501

tmp15427 := GConsP(__e, let__13501)

var ifres15388 Obj

if True == tmp15427 {
tmp15389 := GHd(__e, let__13501)

let__13502 := tmp15389
_ = let__13502

tmp15390 := GTl(__e, let__13501)

tmp15391 := Call2(__e, SymFn(symshen_4lazyderef), tmp15390, V5342)


let__13503 := tmp15391
_ = let__13503

tmp15426 := GConsP(__e, let__13503)

var ifres15392 Obj

if True == tmp15426 {
tmp15393 := GHd(__e, let__13503)

tmp15394 := Call2(__e, SymFn(symshen_4lazyderef), tmp15393, V5342)


let__13504 := tmp15394
_ = let__13504

tmp15425 := GEq(__e, let__13504, sym_d)

var ifres15395 Obj

if True == tmp15425 {
tmp15396 := GTl(__e, let__13503)

tmp15397 := Call2(__e, SymFn(symshen_4lazyderef), tmp15396, V5342)


let__13505 := tmp15397
_ = let__13505

tmp15424 := GConsP(__e, let__13505)

var ifres15398 Obj

if True == tmp15424 {
tmp15399 := GHd(__e, let__13505)

let__13506 := tmp15399
_ = let__13506

tmp15400 := GTl(__e, let__13505)

tmp15401 := Call2(__e, SymFn(symshen_4lazyderef), tmp15400, V5342)


let__13507 := tmp15401
_ = let__13507

tmp15423 := GEq(__e, let__13507, Nil)

var ifres15402 Obj

if True == tmp15423 {
tmp15403 := GTl(__e, let__13500)

tmp15404 := Call2(__e, SymFn(symshen_4lazyderef), tmp15403, V5342)


let__13508 := tmp15404
_ = let__13508

tmp15422 := GEq(__e, let__13508, Nil)

var ifres15405 Obj

if True == tmp15422 {
tmp15406 := GTl(__e, let__13489)

let__13509 := tmp15406
_ = let__13509

tmp15407 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15407

tmp15408 := Call2(__e, SymFn(symshen_4deref), let__13499, V5342)


tmp15409 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp15410 := GEq(__e, tmp15408, tmp15409)

tmp15411 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15412 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15413 := GCons(__e, let__13502, Nil)

tmp15414 := GCons(__e, let__13499, tmp15413)

tmp15415 := GCons(__e, let__13494, tmp15414)

tmp15416 := GCons(__e, let__13506, Nil)

tmp15417 := GCons(__e, let__13499, tmp15416)

tmp15418 := GCons(__e, let__13496, tmp15417)

tmp15419 := GCons(__e, tmp15418, let__13509)

tmp15420 := GCons(__e, tmp15415, tmp15419)

__e.TailApply(SymFn(symshen_4l_1rules), tmp15420, V5340, True, V5342, V5343, let__13464, V5345)
return nil


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5342, V5343, let__13464, tmp15412)


panic("unreachable")
})

tmp15421 := Call5(__e, SymFn(symwhen), tmp15410, V5342, V5343, let__13464, tmp15411)


ifres15405 = tmp15421


} else {
ifres15405 = False


}

ifres15402 = ifres15405


} else {
ifres15402 = False


}

ifres15398 = ifres15402


} else {
ifres15398 = False


}

ifres15395 = ifres15398


} else {
ifres15395 = False


}

ifres15392 = ifres15395


} else {
ifres15392 = False


}

ifres15388 = ifres15392


} else {
ifres15388 = False


}

ifres15385 = ifres15388


} else {
ifres15385 = False


}

ifres15381 = ifres15385


} else {
ifres15381 = False


}

ifres15378 = ifres15381


} else {
ifres15378 = False


}

ifres15374 = ifres15378


} else {
ifres15374 = False


}

ifres15370 = ifres15374


} else {
ifres15370 = False


}

ifres15367 = ifres15370


} else {
ifres15367 = False


}

ifres15364 = ifres15367


} else {
ifres15364 = False


}

ifres15361 = ifres15364


} else {
ifres15361 = False


}

ifres15358 = ifres15361


} else {
ifres15358 = False


}

ifres15356 = ifres15358


} else {
ifres15356 = False


}

let__13488 := ifres15356
_ = let__13488

tmp15615 := GEq(__e, let__13488, False)

if True == tmp15615 {
tmp15501 := Call1(__e, SymFn(symshen_4unlocked_2), V5343)


var ifres15438 Obj

if True == tmp15501 {
tmp15439 := Call2(__e, SymFn(symshen_4lazyderef), V5339, V5342)


let__13511 := tmp15439
_ = let__13511

tmp15500 := GConsP(__e, let__13511)

var ifres15440 Obj

if True == tmp15500 {
tmp15441 := GHd(__e, let__13511)

tmp15442 := Call2(__e, SymFn(symshen_4lazyderef), tmp15441, V5342)


let__13512 := tmp15442
_ = let__13512

tmp15499 := GConsP(__e, let__13512)

var ifres15443 Obj

if True == tmp15499 {
tmp15444 := GHd(__e, let__13512)

tmp15445 := Call2(__e, SymFn(symshen_4lazyderef), tmp15444, V5342)


let__13513 := tmp15445
_ = let__13513

tmp15498 := GConsP(__e, let__13513)

var ifres15446 Obj

if True == tmp15498 {
tmp15447 := GHd(__e, let__13513)

tmp15448 := Call2(__e, SymFn(symshen_4lazyderef), tmp15447, V5342)


let__13514 := tmp15448
_ = let__13514

tmp15497 := GEq(__e, let__13514, sym_8s)

var ifres15449 Obj

if True == tmp15497 {
tmp15450 := GTl(__e, let__13513)

tmp15451 := Call2(__e, SymFn(symshen_4lazyderef), tmp15450, V5342)


let__13515 := tmp15451
_ = let__13515

tmp15496 := GConsP(__e, let__13515)

var ifres15452 Obj

if True == tmp15496 {
tmp15453 := GHd(__e, let__13515)

let__13516 := tmp15453
_ = let__13516

tmp15454 := GTl(__e, let__13515)

tmp15455 := Call2(__e, SymFn(symshen_4lazyderef), tmp15454, V5342)


let__13517 := tmp15455
_ = let__13517

tmp15495 := GConsP(__e, let__13517)

var ifres15456 Obj

if True == tmp15495 {
tmp15457 := GHd(__e, let__13517)

let__13518 := tmp15457
_ = let__13518

tmp15458 := GTl(__e, let__13517)

tmp15459 := Call2(__e, SymFn(symshen_4lazyderef), tmp15458, V5342)


let__13519 := tmp15459
_ = let__13519

tmp15494 := GEq(__e, let__13519, Nil)

var ifres15460 Obj

if True == tmp15494 {
tmp15461 := GTl(__e, let__13512)

tmp15462 := Call2(__e, SymFn(symshen_4lazyderef), tmp15461, V5342)


let__13520 := tmp15462
_ = let__13520

tmp15493 := GConsP(__e, let__13520)

var ifres15463 Obj

if True == tmp15493 {
tmp15464 := GHd(__e, let__13520)

let__13521 := tmp15464
_ = let__13521

tmp15465 := GTl(__e, let__13520)

tmp15466 := Call2(__e, SymFn(symshen_4lazyderef), tmp15465, V5342)


let__13522 := tmp15466
_ = let__13522

tmp15492 := GConsP(__e, let__13522)

var ifres15467 Obj

if True == tmp15492 {
tmp15468 := GHd(__e, let__13522)

tmp15469 := Call2(__e, SymFn(symshen_4lazyderef), tmp15468, V5342)


let__13523 := tmp15469
_ = let__13523

tmp15491 := GEq(__e, let__13523, symstring)

var ifres15470 Obj

if True == tmp15491 {
tmp15471 := GTl(__e, let__13522)

tmp15472 := Call2(__e, SymFn(symshen_4lazyderef), tmp15471, V5342)


let__13524 := tmp15472
_ = let__13524

tmp15490 := GEq(__e, let__13524, Nil)

var ifres15473 Obj

if True == tmp15490 {
tmp15474 := GTl(__e, let__13511)

let__13525 := tmp15474
_ = let__13525

tmp15475 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15475

tmp15476 := Call2(__e, SymFn(symshen_4deref), let__13521, V5342)


tmp15477 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp15478 := GEq(__e, tmp15476, tmp15477)

tmp15479 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15480 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15481 := GCons(__e, symstring, Nil)

tmp15482 := GCons(__e, let__13521, tmp15481)

tmp15483 := GCons(__e, let__13516, tmp15482)

tmp15484 := GCons(__e, symstring, Nil)

tmp15485 := GCons(__e, let__13521, tmp15484)

tmp15486 := GCons(__e, let__13518, tmp15485)

tmp15487 := GCons(__e, tmp15486, let__13525)

tmp15488 := GCons(__e, tmp15483, tmp15487)

__e.TailApply(SymFn(symshen_4l_1rules), tmp15488, V5340, True, V5342, V5343, let__13464, V5345)
return nil


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5342, V5343, let__13464, tmp15480)


panic("unreachable")
})

tmp15489 := Call5(__e, SymFn(symwhen), tmp15478, V5342, V5343, let__13464, tmp15479)


ifres15473 = tmp15489


} else {
ifres15473 = False


}

ifres15470 = ifres15473


} else {
ifres15470 = False


}

ifres15467 = ifres15470


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

ifres15456 = ifres15460


} else {
ifres15456 = False


}

ifres15452 = ifres15456


} else {
ifres15452 = False


}

ifres15449 = ifres15452


} else {
ifres15449 = False


}

ifres15446 = ifres15449


} else {
ifres15446 = False


}

ifres15443 = ifres15446


} else {
ifres15443 = False


}

ifres15440 = ifres15443


} else {
ifres15440 = False


}

ifres15438 = ifres15440


} else {
ifres15438 = False


}

let__13510 := ifres15438
_ = let__13510

tmp15613 := GEq(__e, let__13510, False)

if True == tmp15613 {
tmp15580 := Call1(__e, SymFn(symshen_4unlocked_2), V5343)


var ifres15502 Obj

if True == tmp15580 {
tmp15503 := Call2(__e, SymFn(symshen_4lazyderef), V5339, V5342)


let__13527 := tmp15503
_ = let__13527

tmp15579 := GConsP(__e, let__13527)

var ifres15504 Obj

if True == tmp15579 {
tmp15505 := GHd(__e, let__13527)

tmp15506 := Call2(__e, SymFn(symshen_4lazyderef), tmp15505, V5342)


let__13528 := tmp15506
_ = let__13528

tmp15578 := GConsP(__e, let__13528)

var ifres15507 Obj

if True == tmp15578 {
tmp15508 := GHd(__e, let__13528)

tmp15509 := Call2(__e, SymFn(symshen_4lazyderef), tmp15508, V5342)


let__13529 := tmp15509
_ = let__13529

tmp15577 := GConsP(__e, let__13529)

var ifres15510 Obj

if True == tmp15577 {
tmp15511 := GHd(__e, let__13529)

tmp15512 := Call2(__e, SymFn(symshen_4lazyderef), tmp15511, V5342)


let__13530 := tmp15512
_ = let__13530

tmp15576 := GEq(__e, let__13530, sym_8v)

var ifres15513 Obj

if True == tmp15576 {
tmp15514 := GTl(__e, let__13529)

tmp15515 := Call2(__e, SymFn(symshen_4lazyderef), tmp15514, V5342)


let__13531 := tmp15515
_ = let__13531

tmp15575 := GConsP(__e, let__13531)

var ifres15516 Obj

if True == tmp15575 {
tmp15517 := GHd(__e, let__13531)

let__13532 := tmp15517
_ = let__13532

tmp15518 := GTl(__e, let__13531)

tmp15519 := Call2(__e, SymFn(symshen_4lazyderef), tmp15518, V5342)


let__13533 := tmp15519
_ = let__13533

tmp15574 := GConsP(__e, let__13533)

var ifres15520 Obj

if True == tmp15574 {
tmp15521 := GHd(__e, let__13533)

let__13534 := tmp15521
_ = let__13534

tmp15522 := GTl(__e, let__13533)

tmp15523 := Call2(__e, SymFn(symshen_4lazyderef), tmp15522, V5342)


let__13535 := tmp15523
_ = let__13535

tmp15573 := GEq(__e, let__13535, Nil)

var ifres15524 Obj

if True == tmp15573 {
tmp15525 := GTl(__e, let__13528)

tmp15526 := Call2(__e, SymFn(symshen_4lazyderef), tmp15525, V5342)


let__13536 := tmp15526
_ = let__13536

tmp15572 := GConsP(__e, let__13536)

var ifres15527 Obj

if True == tmp15572 {
tmp15528 := GHd(__e, let__13536)

let__13537 := tmp15528
_ = let__13537

tmp15529 := GTl(__e, let__13536)

tmp15530 := Call2(__e, SymFn(symshen_4lazyderef), tmp15529, V5342)


let__13538 := tmp15530
_ = let__13538

tmp15571 := GConsP(__e, let__13538)

var ifres15531 Obj

if True == tmp15571 {
tmp15532 := GHd(__e, let__13538)

tmp15533 := Call2(__e, SymFn(symshen_4lazyderef), tmp15532, V5342)


let__13539 := tmp15533
_ = let__13539

tmp15570 := GConsP(__e, let__13539)

var ifres15534 Obj

if True == tmp15570 {
tmp15535 := GHd(__e, let__13539)

tmp15536 := Call2(__e, SymFn(symshen_4lazyderef), tmp15535, V5342)


let__13540 := tmp15536
_ = let__13540

tmp15569 := GEq(__e, let__13540, symvector)

var ifres15537 Obj

if True == tmp15569 {
tmp15538 := GTl(__e, let__13539)

tmp15539 := Call2(__e, SymFn(symshen_4lazyderef), tmp15538, V5342)


let__13541 := tmp15539
_ = let__13541

tmp15568 := GConsP(__e, let__13541)

var ifres15540 Obj

if True == tmp15568 {
tmp15541 := GHd(__e, let__13541)

let__13542 := tmp15541
_ = let__13542

tmp15542 := GTl(__e, let__13541)

tmp15543 := Call2(__e, SymFn(symshen_4lazyderef), tmp15542, V5342)


let__13543 := tmp15543
_ = let__13543

tmp15567 := GEq(__e, let__13543, Nil)

var ifres15544 Obj

if True == tmp15567 {
tmp15545 := GTl(__e, let__13538)

tmp15546 := Call2(__e, SymFn(symshen_4lazyderef), tmp15545, V5342)


let__13544 := tmp15546
_ = let__13544

tmp15566 := GEq(__e, let__13544, Nil)

var ifres15547 Obj

if True == tmp15566 {
tmp15548 := GTl(__e, let__13527)

let__13545 := tmp15548
_ = let__13545

tmp15549 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15549

tmp15550 := Call2(__e, SymFn(symshen_4deref), let__13537, V5342)


tmp15551 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp15552 := GEq(__e, tmp15550, tmp15551)

tmp15553 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15554 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15555 := GCons(__e, let__13542, Nil)

tmp15556 := GCons(__e, let__13537, tmp15555)

tmp15557 := GCons(__e, let__13532, tmp15556)

tmp15558 := GCons(__e, let__13542, Nil)

tmp15559 := GCons(__e, symvector, tmp15558)

tmp15560 := GCons(__e, tmp15559, Nil)

tmp15561 := GCons(__e, let__13537, tmp15560)

tmp15562 := GCons(__e, let__13534, tmp15561)

tmp15563 := GCons(__e, tmp15562, let__13545)

tmp15564 := GCons(__e, tmp15557, tmp15563)

__e.TailApply(SymFn(symshen_4l_1rules), tmp15564, V5340, True, V5342, V5343, let__13464, V5345)
return nil


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5342, V5343, let__13464, tmp15554)


panic("unreachable")
})

tmp15565 := Call5(__e, SymFn(symwhen), tmp15552, V5342, V5343, let__13464, tmp15553)


ifres15547 = tmp15565


} else {
ifres15547 = False


}

ifres15544 = ifres15547


} else {
ifres15544 = False


}

ifres15540 = ifres15544


} else {
ifres15540 = False


}

ifres15537 = ifres15540


} else {
ifres15537 = False


}

ifres15534 = ifres15537


} else {
ifres15534 = False


}

ifres15531 = ifres15534


} else {
ifres15531 = False


}

ifres15527 = ifres15531


} else {
ifres15527 = False


}

ifres15524 = ifres15527


} else {
ifres15524 = False


}

ifres15520 = ifres15524


} else {
ifres15520 = False


}

ifres15516 = ifres15520


} else {
ifres15516 = False


}

ifres15513 = ifres15516


} else {
ifres15513 = False


}

ifres15510 = ifres15513


} else {
ifres15510 = False


}

ifres15507 = ifres15510


} else {
ifres15507 = False


}

ifres15504 = ifres15507


} else {
ifres15504 = False


}

ifres15502 = ifres15504


} else {
ifres15502 = False


}

let__13526 := ifres15502
_ = let__13526

tmp15611 := GEq(__e, let__13526, False)

if True == tmp15611 {
tmp15607 := Call1(__e, SymFn(symshen_4unlocked_2), V5343)


var ifres15581 Obj

if True == tmp15607 {
tmp15582 := Call2(__e, SymFn(symshen_4lazyderef), V5339, V5342)


let__13547 := tmp15582
_ = let__13547

tmp15606 := GConsP(__e, let__13547)

var ifres15583 Obj

if True == tmp15606 {
tmp15584 := GHd(__e, let__13547)

let__13548 := tmp15584
_ = let__13548

tmp15585 := GTl(__e, let__13547)

let__13549 := tmp15585
_ = let__13549

tmp15586 := Call2(__e, SymFn(symshen_4lazyderef), V5340, V5342)


let__13550 := tmp15586
_ = let__13550

tmp15587 := MakeNative1(func(__e *ControlFlow, Z5434 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5435 Obj) Obj {
tmp15588 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15588

tmp15589 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4l_1rules), let__13549, Z5435, V5341, V5342, V5343, let__13464, V5345)
return nil
panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), Z5434, let__13548, V5342, V5343, let__13464, tmp15589)


panic("unreachable")
})
panic("unreachable")
})

let__13551 := tmp15587
_ = let__13551

tmp15605 := GConsP(__e, let__13550)

var ifres15590 Obj

if True == tmp15605 {
tmp15591 := GHd(__e, let__13550)

let__13552 := tmp15591
_ = let__13552

tmp15592 := GTl(__e, let__13550)

let__13553 := tmp15592
_ = let__13553

tmp15593 := Call1(__e, let__13551, let__13552)


tmp15594 := Call1(__e, tmp15593, let__13553)


ifres15590 = tmp15594


} else {
tmp15604 := Call1(__e, SymFn(symshen_4pvar_2), let__13550)


var ifres15595 Obj

if True == tmp15604 {
tmp15596 := Call1(__e, SymFn(symshen_4newpv), V5342)


let__13554 := tmp15596
_ = let__13554

tmp15597 := Call1(__e, SymFn(symshen_4newpv), V5342)


let__13555 := tmp15597
_ = let__13555

tmp15598 := GCons(__e, let__13554, let__13555)

tmp15599 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15600 := Call1(__e, let__13551, let__13554)


return TailCall1(__e, tmp15600, let__13555)


panic("unreachable")
})

tmp15601 := Call4(__e, SymFn(symshen_4bind_b), let__13550, tmp15598, V5342, tmp15599)


tmp15602 := Call2(__e, SymFn(symshen_4gc), V5342, tmp15601)


tmp15603 := Call2(__e, SymFn(symshen_4gc), V5342, tmp15602)


ifres15595 = tmp15603


} else {
ifres15595 = False


}

ifres15590 = ifres15595


}

ifres15583 = ifres15590


} else {
ifres15583 = False


}

ifres15581 = ifres15583


} else {
ifres15581 = False


}

let__13546 := ifres15581
_ = let__13546

tmp15609 := GEq(__e, let__13546, False)

if True == tmp15609 {
return TailCall2(__e, SymFn(symshen_4unlock), V5343, let__13464)
} else {
return let__13546
}


} else {
return let__13526
}


} else {
return let__13510
}


} else {
return let__13488
}


} else {
return let__13468
}


} else {
return let__13465
}


panic("unreachable")
}())
}, 7)

tmp15620 := Call2(__e, ns2_1set, symshen_4l_1rules, tmp15264)


_ = tmp15620

tmp15621 := MakeNative6(func(__e *ControlFlow, V5440, V5441, V5442, V5443, V5444, V5445 Obj) Obj {
let__13556 := GAdd(__e, V5444, MakeInteger(1))
_ = let__13556

tmp15661 := Call1(__e, SymFn(symshen_4unlocked_2), V5443)


var ifres15623 Obj

if True == tmp15661 {
tmp15624 := Call2(__e, SymFn(symshen_4lazyderef), V5440, V5442)


let__13558 := tmp15624
_ = let__13558

tmp15660 := GConsP(__e, let__13558)

var ifres15625 Obj

if True == tmp15660 {
tmp15626 := GHd(__e, let__13558)

tmp15627 := Call2(__e, SymFn(symshen_4lazyderef), tmp15626, V5442)


let__13559 := tmp15627
_ = let__13559

tmp15659 := GEq(__e, let__13559, symdefine)

var ifres15628 Obj

if True == tmp15659 {
tmp15629 := GTl(__e, let__13558)

tmp15630 := Call2(__e, SymFn(symshen_4lazyderef), tmp15629, V5442)


let__13560 := tmp15630
_ = let__13560

tmp15658 := GConsP(__e, let__13560)

var ifres15631 Obj

if True == tmp15658 {
tmp15632 := GHd(__e, let__13560)

let__13561 := tmp15632
_ = let__13561

tmp15633 := GTl(__e, let__13560)

let__13562 := tmp15633
_ = let__13562

tmp15634 := Call1(__e, SymFn(symshen_4newpv), V5442)


let__13563 := tmp15634
_ = let__13563

tmp15635 := Call1(__e, SymFn(symshen_4newpv), V5442)


let__13564 := tmp15635
_ = let__13564

tmp15636 := Call1(__e, SymFn(symshen_4newpv), V5442)


let__13565 := tmp15636
_ = let__13565

tmp15637 := Call1(__e, SymFn(symshen_4newpv), V5442)


let__13566 := tmp15637
_ = let__13566

tmp15638 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15638

tmp15639 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15640 := GCons(__e, let__13561, let__13562)

tmp15641 := Call1(__e, SymFn(symshen_4sigxrules), tmp15640)


tmp15642 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15643 := Call2(__e, SymFn(symshen_4lazyderef), let__13563, V5442)


tmp15644 := Call1(__e, SymFn(symfst), tmp15643)


tmp15645 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15646 := Call2(__e, SymFn(symshen_4lazyderef), let__13563, V5442)


tmp15647 := Call1(__e, SymFn(symsnd), tmp15646)


tmp15648 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15649 := Call2(__e, SymFn(symshen_4deref), let__13566, V5442)


tmp15650 := Call1(__e, SymFn(symshen_4freshen_1sig), tmp15649)


tmp15651 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15652 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall6(__e, SymFn(symis), let__13566, V5441, V5442, V5443, let__13556, V5445)
panic("unreachable")
})

__e.TailApply(SymFn(symshen_4t_d_1rules), let__13561, let__13564, let__13565, MakeInteger(1), V5442, V5443, let__13556, tmp15652)
return nil


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), let__13565, tmp15650, V5442, V5443, let__13556, tmp15651)


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), let__13564, tmp15647, V5442, V5443, let__13556, tmp15648)


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), let__13566, tmp15644, V5442, V5443, let__13556, tmp15645)


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), let__13563, tmp15641, V5442, V5443, let__13556, tmp15642)


panic("unreachable")
})

tmp15653 := Call4(__e, SymFn(symshen_4cut), V5442, V5443, let__13556, tmp15639)


tmp15654 := Call2(__e, SymFn(symshen_4gc), V5442, tmp15653)


tmp15655 := Call2(__e, SymFn(symshen_4gc), V5442, tmp15654)


tmp15656 := Call2(__e, SymFn(symshen_4gc), V5442, tmp15655)


tmp15657 := Call2(__e, SymFn(symshen_4gc), V5442, tmp15656)


ifres15631 = tmp15657


} else {
ifres15631 = False


}

ifres15628 = ifres15631


} else {
ifres15628 = False


}

ifres15625 = ifres15628


} else {
ifres15625 = False


}

ifres15623 = ifres15625


} else {
ifres15623 = False


}

let__13557 := ifres15623
_ = let__13557

tmp15663 := GEq(__e, let__13557, False)

if True == tmp15663 {
return TailCall2(__e, SymFn(symshen_4unlock), V5443, let__13556)
} else {
return let__13557
}


panic("unreachable")
})

tmp15664 := Call2(__e, ns2_1set, symshen_4t_d, tmp15621)


_ = tmp15664

tmp15665 := MakeNative1(func(__e *ControlFlow, V5457 Obj) Obj {
tmp15666 := MakeNative1(func(__e *ControlFlow, Z5459 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5sig_drules_6), Z5459)
panic("unreachable")
})

tmp15667 := Call2(__e, SymFn(symcompile), tmp15666, V5457)


let__13567 := tmp15667
_ = let__13567

tmp15668 := Call1(__e, SymFn(symfst), let__13567)


tmp15669 := MakeNative1(func(__e *ControlFlow, Z5460 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4unmark_1rule_1custom_1patterns), Z5460)
panic("unreachable")
})

tmp15670 := Call1(__e, SymFn(symsnd), let__13567)


tmp15671 := Call2(__e, SymFn(symmap), tmp15669, tmp15670)


return TailCall2(__e, SymFn(sym_8p), tmp15668, tmp15671)


panic("unreachable")
})

tmp15672 := Call2(__e, ns2_1set, symshen_4sigxrules, tmp15665)


_ = tmp15672

tmp15673 := MakeNative1(func(__e *ControlFlow, V5461 Obj) Obj {
tmp15679 := Call1(__e, SymFn(symtuple_2), V5461)


if True == tmp15679 {
tmp15674 := MakeNative1(func(__e *ControlFlow, Z5462 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4unmark_1custom_1pattern), Z5462)
panic("unreachable")
})

tmp15675 := Call1(__e, SymFn(symfst), V5461)


tmp15676 := Call2(__e, SymFn(symmap), tmp15674, tmp15675)


tmp15677 := Call1(__e, SymFn(symsnd), V5461)


return TailCall2(__e, SymFn(sym_8p), tmp15676, tmp15677)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4unmark_1rule_1custom_1patterns)
}


panic("unreachable")
})

tmp15680 := Call2(__e, ns2_1set, symshen_4unmark_1rule_1custom_1patterns, tmp15673)


_ = tmp15680

tmp15681 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V5463 := __self1
_ = V5463
tmp15712 := GConsP(__e, V5463)

var ifres15693 Obj

if True == tmp15712 {
tmp15710 := GHd(__e, V5463)

tmp15711 := GEq(__e, sym_8p, tmp15710)

var ifres15694 Obj

if True == tmp15711 {
tmp15708 := GTl(__e, V5463)

tmp15709 := GConsP(__e, tmp15708)

var ifres15695 Obj

if True == tmp15709 {
tmp15705 := GTl(__e, V5463)

tmp15706 := GHd(__e, tmp15705)

tmp15707 := GEq(__e, symshen_4custom_1pattern, tmp15706)

var ifres15696 Obj

if True == tmp15707 {
tmp15702 := GTl(__e, V5463)

tmp15703 := GTl(__e, tmp15702)

tmp15704 := GConsP(__e, tmp15703)

var ifres15697 Obj

if True == tmp15704 {
tmp15698 := GTl(__e, V5463)

tmp15699 := GTl(__e, tmp15698)

tmp15700 := GTl(__e, tmp15699)

tmp15701 := GEq(__e, Nil, tmp15700)

ifres15697 = tmp15701


} else {
ifres15697 = False


}

ifres15696 = ifres15697


} else {
ifres15696 = False


}

ifres15695 = ifres15696


} else {
ifres15695 = False


}

ifres15694 = ifres15695


} else {
ifres15694 = False


}

ifres15693 = ifres15694


} else {
ifres15693 = False


}

if True == ifres15693 {
tmp15682 := GTl(__e, V5463)

tmp15683 := GTl(__e, tmp15682)

tmp15684 := GHd(__e, tmp15683)

tmp15685 := Call1(__e, SymFn(symshen_4curry), tmp15684)


if SymFn(symshen_4unmark_1custom_1pattern) == __self {
__self1 = tmp15685
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4unmark_1custom_1pattern), tmp15685)


} else {
tmp15691 := GConsP(__e, V5463)

if True == tmp15691 {
tmp15686 := GHd(__e, V5463)

tmp15687 := Call1(__e, SymFn(symshen_4unmark_1custom_1pattern), tmp15686)


tmp15688 := GTl(__e, V5463)

tmp15689 := Call1(__e, SymFn(symshen_4unmark_1custom_1pattern), tmp15688)


return GCons(__e, tmp15687, tmp15689)


} else {
return V5463
}


}


panic("unreachable")
})
return __self
})()

tmp15713 := Call2(__e, ns2_1set, symshen_4unmark_1custom_1pattern, tmp15681)


_ = tmp15713

tmp15714 := MakeNative1(func(__e *ControlFlow, V5464 Obj) Obj {
tmp15741 := GConsP(__e, V5464)

var ifres15715 Obj

if True == tmp15741 {
tmp15716 := Call1(__e, SymFn(symtail), V5464)


let__13569 := tmp15716
_ = let__13569

tmp15739 := Call2(__e, SymFn(symshen_4hds_a_2), let__13569, sym_i)


var ifres15717 Obj

if True == tmp15739 {
tmp15718 := Call1(__e, SymFn(symtail), let__13569)


let__13570 := tmp15718
_ = let__13570

tmp15719 := Call1(__e, SymFn(symshen_4_5signature_6), let__13570)


let__13571 := tmp15719
_ = let__13571

tmp15737 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13571)


var ifres15720 Obj

if True == tmp15737 {
tmp15721 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15720 = tmp15721


} else {
tmp15722 := Call1(__e, SymFn(symshen_4_5_1out), let__13571)


let__13572 := tmp15722
_ = let__13572

tmp15723 := Call1(__e, SymFn(symshen_4in_1_6), let__13571)


let__13573 := tmp15723
_ = let__13573

tmp15736 := Call2(__e, SymFn(symshen_4hds_a_2), let__13573, sym_j)


var ifres15724 Obj

if True == tmp15736 {
tmp15725 := Call1(__e, SymFn(symtail), let__13573)


let__13574 := tmp15725
_ = let__13574

tmp15726 := Call1(__e, SymFn(symshen_4_5rules_d_6), let__13574)


let__13575 := tmp15726
_ = let__13575

tmp15734 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13575)


var ifres15727 Obj

if True == tmp15734 {
tmp15728 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15727 = tmp15728


} else {
tmp15729 := Call1(__e, SymFn(symshen_4_5_1out), let__13575)


let__13576 := tmp15729
_ = let__13576

tmp15730 := Call1(__e, SymFn(symshen_4in_1_6), let__13575)


let__13577 := tmp15730
_ = let__13577

tmp15731 := Call1(__e, SymFn(symshen_4rectify_1type), let__13572)


let__13578 := tmp15731
_ = let__13578

tmp15732 := Call2(__e, SymFn(sym_8p), let__13578, let__13576)


tmp15733 := Call2(__e, SymFn(symshen_4comb), let__13577, tmp15732)


ifres15727 = tmp15733


}

ifres15724 = ifres15727


} else {
tmp15735 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15724 = tmp15735


}

ifres15720 = ifres15724


}

ifres15717 = ifres15720


} else {
tmp15738 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15717 = tmp15738


}

ifres15715 = ifres15717


} else {
tmp15740 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15715 = tmp15740


}

let__13568 := ifres15715
_ = let__13568

tmp15743 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13568)


if True == tmp15743 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__13568
}


panic("unreachable")
})

tmp15744 := Call2(__e, ns2_1set, symshen_4_5sig_drules_6, tmp15714)


_ = tmp15744

tmp15745 := MakeNative1(func(__e *ControlFlow, V5476 Obj) Obj {
tmp15746 := Call1(__e, SymFn(symshen_4extract_1vars), V5476)


let__13579 := tmp15746
_ = let__13579

tmp15747 := MakeNative1(func(__e *ControlFlow, Z5479 Obj) Obj {
tmp15748 := Call2(__e, SymFn(symconcat), sym_e, Z5479)


tmp15749 := Call1(__e, SymFn(symshen_4freshterm), tmp15748)


return GCons(__e, Z5479, tmp15749)


panic("unreachable")
})

tmp15750 := Call2(__e, SymFn(symmap), tmp15747, let__13579)


let__13580 := tmp15750
_ = let__13580

return TailCall2(__e, SymFn(symshen_4freshen_1type), let__13580, V5476)


panic("unreachable")
})

tmp15751 := Call2(__e, ns2_1set, symshen_4freshen_1sig, tmp15745)


_ = tmp15751

tmp15752 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V5480 := __self1
_ = V5480
V5481 := __self2
_ = V5481
tmp15765 := GEq(__e, Nil, V5480)

if True == tmp15765 {
return V5481
} else {
tmp15763 := GConsP(__e, V5480)

var ifres15760 Obj

if True == tmp15763 {
tmp15761 := GHd(__e, V5480)

tmp15762 := GConsP(__e, tmp15761)

ifres15760 = tmp15762


} else {
ifres15760 = False


}

if True == ifres15760 {
tmp15753 := GTl(__e, V5480)

tmp15754 := GHd(__e, V5480)

tmp15755 := GTl(__e, tmp15754)

tmp15756 := GHd(__e, V5480)

tmp15757 := GHd(__e, tmp15756)

tmp15758 := Call3(__e, SymFn(symsubst), tmp15755, tmp15757, V5481)


if SymFn(symshen_4freshen_1type) == __self {
__self1, __self2 = tmp15753, tmp15758
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4freshen_1type), tmp15753, tmp15758)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4freshen_1type)
}


}


panic("unreachable")
})
return __self
})()

tmp15766 := Call2(__e, ns2_1set, symshen_4freshen_1type, tmp15752)


_ = tmp15766

tmp15767 := MakeNative1(func(__e *ControlFlow, V5482 Obj) Obj {
tmp15768 := Call1(__e, SymFn(symshen_4_5rule_d_6), V5482)


let__13582 := tmp15768
_ = let__13582

tmp15781 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13582)


var ifres15769 Obj

if True == tmp15781 {
tmp15770 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15769 = tmp15770


} else {
tmp15771 := Call1(__e, SymFn(symshen_4_5_1out), let__13582)


let__13583 := tmp15771
_ = let__13583

tmp15772 := Call1(__e, SymFn(symshen_4in_1_6), let__13582)


let__13584 := tmp15772
_ = let__13584

tmp15773 := Call1(__e, SymFn(symshen_4_5rules_d_6), let__13584)


let__13585 := tmp15773
_ = let__13585

tmp15780 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13585)


var ifres15774 Obj

if True == tmp15780 {
tmp15775 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15774 = tmp15775


} else {
tmp15776 := Call1(__e, SymFn(symshen_4_5_1out), let__13585)


let__13586 := tmp15776
_ = let__13586

tmp15777 := Call1(__e, SymFn(symshen_4in_1_6), let__13585)


let__13587 := tmp15777
_ = let__13587

tmp15778 := GCons(__e, let__13583, let__13586)

tmp15779 := Call2(__e, SymFn(symshen_4comb), let__13587, tmp15778)


ifres15774 = tmp15779


}

ifres15769 = ifres15774


}

let__13581 := ifres15769
_ = let__13581

tmp15793 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13581)


if True == tmp15793 {
tmp15782 := Call1(__e, SymFn(symshen_4_5rule_d_6), V5482)


let__13589 := tmp15782
_ = let__13589

tmp15789 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13589)


var ifres15783 Obj

if True == tmp15789 {
tmp15784 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15783 = tmp15784


} else {
tmp15785 := Call1(__e, SymFn(symshen_4_5_1out), let__13589)


let__13590 := tmp15785
_ = let__13590

tmp15786 := Call1(__e, SymFn(symshen_4in_1_6), let__13589)


let__13591 := tmp15786
_ = let__13591

tmp15787 := GCons(__e, let__13590, Nil)

tmp15788 := Call2(__e, SymFn(symshen_4comb), let__13591, tmp15787)


ifres15783 = tmp15788


}

let__13588 := ifres15783
_ = let__13588

tmp15791 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13588)


if True == tmp15791 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__13588
}


} else {
return let__13581
}


panic("unreachable")
})

tmp15794 := Call2(__e, ns2_1set, symshen_4_5rules_d_6, tmp15767)


_ = tmp15794

tmp15795 := MakeNative1(func(__e *ControlFlow, V5494 Obj) Obj {
tmp15796 := Call1(__e, SymFn(symshen_4_5patterns_6), V5494)


let__13593 := tmp15796
_ = let__13593

tmp15824 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13593)


var ifres15797 Obj

if True == tmp15824 {
tmp15798 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15797 = tmp15798


} else {
tmp15799 := Call1(__e, SymFn(symshen_4_5_1out), let__13593)


let__13594 := tmp15799
_ = let__13594

tmp15800 := Call1(__e, SymFn(symshen_4in_1_6), let__13593)


let__13595 := tmp15800
_ = let__13595

tmp15823 := Call2(__e, SymFn(symshen_4hds_a_2), let__13595, sym_1_6)


var ifres15801 Obj

if True == tmp15823 {
tmp15802 := Call1(__e, SymFn(symtail), let__13595)


let__13596 := tmp15802
_ = let__13596

tmp15821 := GConsP(__e, let__13596)

var ifres15803 Obj

if True == tmp15821 {
tmp15804 := Call1(__e, SymFn(symhead), let__13596)


let__13597 := tmp15804
_ = let__13597

tmp15805 := Call1(__e, SymFn(symtail), let__13596)


let__13598 := tmp15805
_ = let__13598

tmp15819 := Call2(__e, SymFn(symshen_4hds_a_2), let__13598, symwhere)


var ifres15806 Obj

if True == tmp15819 {
tmp15807 := Call1(__e, SymFn(symtail), let__13598)


let__13599 := tmp15807
_ = let__13599

tmp15817 := GConsP(__e, let__13599)

var ifres15808 Obj

if True == tmp15817 {
tmp15809 := Call1(__e, SymFn(symhead), let__13599)


let__13600 := tmp15809
_ = let__13600

tmp15810 := Call1(__e, SymFn(symtail), let__13599)


let__13601 := tmp15810
_ = let__13601

tmp15811 := GCons(__e, let__13597, Nil)

tmp15812 := GCons(__e, let__13600, tmp15811)

tmp15813 := GCons(__e, symwhere, tmp15812)

tmp15814 := Call2(__e, SymFn(sym_8p), let__13594, tmp15813)


tmp15815 := Call2(__e, SymFn(symshen_4comb), let__13601, tmp15814)


ifres15808 = tmp15815


} else {
tmp15816 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15808 = tmp15816


}

ifres15806 = ifres15808


} else {
tmp15818 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15806 = tmp15818


}

ifres15803 = ifres15806


} else {
tmp15820 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15803 = tmp15820


}

ifres15801 = ifres15803


} else {
tmp15822 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15801 = tmp15822


}

ifres15797 = ifres15801


}

let__13592 := ifres15797
_ = let__13592

tmp15897 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13592)


if True == tmp15897 {
tmp15825 := Call1(__e, SymFn(symshen_4_5patterns_6), V5494)


let__13603 := tmp15825
_ = let__13603

tmp15854 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13603)


var ifres15826 Obj

if True == tmp15854 {
tmp15827 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15826 = tmp15827


} else {
tmp15828 := Call1(__e, SymFn(symshen_4_5_1out), let__13603)


let__13604 := tmp15828
_ = let__13604

tmp15829 := Call1(__e, SymFn(symshen_4in_1_6), let__13603)


let__13605 := tmp15829
_ = let__13605

tmp15853 := Call2(__e, SymFn(symshen_4hds_a_2), let__13605, sym_5_1)


var ifres15830 Obj

if True == tmp15853 {
tmp15831 := Call1(__e, SymFn(symtail), let__13605)


let__13606 := tmp15831
_ = let__13606

tmp15851 := GConsP(__e, let__13606)

var ifres15832 Obj

if True == tmp15851 {
tmp15833 := Call1(__e, SymFn(symhead), let__13606)


let__13607 := tmp15833
_ = let__13607

tmp15834 := Call1(__e, SymFn(symtail), let__13606)


let__13608 := tmp15834
_ = let__13608

tmp15849 := Call2(__e, SymFn(symshen_4hds_a_2), let__13608, symwhere)


var ifres15835 Obj

if True == tmp15849 {
tmp15836 := Call1(__e, SymFn(symtail), let__13608)


let__13609 := tmp15836
_ = let__13609

tmp15847 := GConsP(__e, let__13609)

var ifres15837 Obj

if True == tmp15847 {
tmp15838 := Call1(__e, SymFn(symhead), let__13609)


let__13610 := tmp15838
_ = let__13610

tmp15839 := Call1(__e, SymFn(symtail), let__13609)


let__13611 := tmp15839
_ = let__13611

tmp15840 := GCons(__e, let__13607, Nil)

tmp15841 := GCons(__e, let__13610, tmp15840)

tmp15842 := GCons(__e, symwhere, tmp15841)

tmp15843 := Call1(__e, SymFn(symshen_4correct), tmp15842)


tmp15844 := Call2(__e, SymFn(sym_8p), let__13604, tmp15843)


tmp15845 := Call2(__e, SymFn(symshen_4comb), let__13611, tmp15844)


ifres15837 = tmp15845


} else {
tmp15846 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15837 = tmp15846


}

ifres15835 = ifres15837


} else {
tmp15848 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15835 = tmp15848


}

ifres15832 = ifres15835


} else {
tmp15850 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15832 = tmp15850


}

ifres15830 = ifres15832


} else {
tmp15852 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15830 = tmp15852


}

ifres15826 = ifres15830


}

let__13602 := ifres15826
_ = let__13602

tmp15895 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13602)


if True == tmp15895 {
tmp15855 := Call1(__e, SymFn(symshen_4_5patterns_6), V5494)


let__13613 := tmp15855
_ = let__13613

tmp15872 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13613)


var ifres15856 Obj

if True == tmp15872 {
tmp15857 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15856 = tmp15857


} else {
tmp15858 := Call1(__e, SymFn(symshen_4_5_1out), let__13613)


let__13614 := tmp15858
_ = let__13614

tmp15859 := Call1(__e, SymFn(symshen_4in_1_6), let__13613)


let__13615 := tmp15859
_ = let__13615

tmp15871 := Call2(__e, SymFn(symshen_4hds_a_2), let__13615, sym_5_1)


var ifres15860 Obj

if True == tmp15871 {
tmp15861 := Call1(__e, SymFn(symtail), let__13615)


let__13616 := tmp15861
_ = let__13616

tmp15869 := GConsP(__e, let__13616)

var ifres15862 Obj

if True == tmp15869 {
tmp15863 := Call1(__e, SymFn(symhead), let__13616)


let__13617 := tmp15863
_ = let__13617

tmp15864 := Call1(__e, SymFn(symtail), let__13616)


let__13618 := tmp15864
_ = let__13618

tmp15865 := Call1(__e, SymFn(symshen_4correct), let__13617)


tmp15866 := Call2(__e, SymFn(sym_8p), let__13614, tmp15865)


tmp15867 := Call2(__e, SymFn(symshen_4comb), let__13618, tmp15866)


ifres15862 = tmp15867


} else {
tmp15868 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15862 = tmp15868


}

ifres15860 = ifres15862


} else {
tmp15870 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15860 = tmp15870


}

ifres15856 = ifres15860


}

let__13612 := ifres15856
_ = let__13612

tmp15893 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13612)


if True == tmp15893 {
tmp15873 := Call1(__e, SymFn(symshen_4_5patterns_6), V5494)


let__13620 := tmp15873
_ = let__13620

tmp15889 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13620)


var ifres15874 Obj

if True == tmp15889 {
tmp15875 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15874 = tmp15875


} else {
tmp15876 := Call1(__e, SymFn(symshen_4_5_1out), let__13620)


let__13621 := tmp15876
_ = let__13621

tmp15877 := Call1(__e, SymFn(symshen_4in_1_6), let__13620)


let__13622 := tmp15877
_ = let__13622

tmp15888 := Call2(__e, SymFn(symshen_4hds_a_2), let__13622, sym_1_6)


var ifres15878 Obj

if True == tmp15888 {
tmp15879 := Call1(__e, SymFn(symtail), let__13622)


let__13623 := tmp15879
_ = let__13623

tmp15886 := GConsP(__e, let__13623)

var ifres15880 Obj

if True == tmp15886 {
tmp15881 := Call1(__e, SymFn(symhead), let__13623)


let__13624 := tmp15881
_ = let__13624

tmp15882 := Call1(__e, SymFn(symtail), let__13623)


let__13625 := tmp15882
_ = let__13625

tmp15883 := Call2(__e, SymFn(sym_8p), let__13621, let__13624)


tmp15884 := Call2(__e, SymFn(symshen_4comb), let__13625, tmp15883)


ifres15880 = tmp15884


} else {
tmp15885 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15880 = tmp15885


}

ifres15878 = ifres15880


} else {
tmp15887 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15878 = tmp15887


}

ifres15874 = ifres15878


}

let__13619 := ifres15874
_ = let__13619

tmp15891 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13619)


if True == tmp15891 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__13619
}


} else {
return let__13612
}


} else {
return let__13602
}


} else {
return let__13592
}


panic("unreachable")
})

tmp15898 := Call2(__e, ns2_1set, symshen_4_5rule_d_6, tmp15795)


_ = tmp15898

tmp15899 := MakeNative1(func(__e *ControlFlow, V5529 Obj) Obj {
tmp16030 := GConsP(__e, V5529)

var ifres15983 Obj

if True == tmp16030 {
tmp16028 := GHd(__e, V5529)

tmp16029 := GEq(__e, symwhere, tmp16028)

var ifres15984 Obj

if True == tmp16029 {
tmp16026 := GTl(__e, V5529)

tmp16027 := GConsP(__e, tmp16026)

var ifres15985 Obj

if True == tmp16027 {
tmp16023 := GTl(__e, V5529)

tmp16024 := GTl(__e, tmp16023)

tmp16025 := GConsP(__e, tmp16024)

var ifres15986 Obj

if True == tmp16025 {
tmp16019 := GTl(__e, V5529)

tmp16020 := GTl(__e, tmp16019)

tmp16021 := GHd(__e, tmp16020)

tmp16022 := GConsP(__e, tmp16021)

var ifres15987 Obj

if True == tmp16022 {
tmp16014 := GTl(__e, V5529)

tmp16015 := GTl(__e, tmp16014)

tmp16016 := GHd(__e, tmp16015)

tmp16017 := GHd(__e, tmp16016)

tmp16018 := GEq(__e, symfail_1if, tmp16017)

var ifres15988 Obj

if True == tmp16018 {
tmp16009 := GTl(__e, V5529)

tmp16010 := GTl(__e, tmp16009)

tmp16011 := GHd(__e, tmp16010)

tmp16012 := GTl(__e, tmp16011)

tmp16013 := GConsP(__e, tmp16012)

var ifres15989 Obj

if True == tmp16013 {
tmp16003 := GTl(__e, V5529)

tmp16004 := GTl(__e, tmp16003)

tmp16005 := GHd(__e, tmp16004)

tmp16006 := GTl(__e, tmp16005)

tmp16007 := GTl(__e, tmp16006)

tmp16008 := GConsP(__e, tmp16007)

var ifres15990 Obj

if True == tmp16008 {
tmp15996 := GTl(__e, V5529)

tmp15997 := GTl(__e, tmp15996)

tmp15998 := GHd(__e, tmp15997)

tmp15999 := GTl(__e, tmp15998)

tmp16000 := GTl(__e, tmp15999)

tmp16001 := GTl(__e, tmp16000)

tmp16002 := GEq(__e, Nil, tmp16001)

var ifres15991 Obj

if True == tmp16002 {
tmp15992 := GTl(__e, V5529)

tmp15993 := GTl(__e, tmp15992)

tmp15994 := GTl(__e, tmp15993)

tmp15995 := GEq(__e, Nil, tmp15994)

ifres15991 = tmp15995


} else {
ifres15991 = False


}

ifres15990 = ifres15991


} else {
ifres15990 = False


}

ifres15989 = ifres15990


} else {
ifres15989 = False


}

ifres15988 = ifres15989


} else {
ifres15988 = False


}

ifres15987 = ifres15988


} else {
ifres15987 = False


}

ifres15986 = ifres15987


} else {
ifres15986 = False


}

ifres15985 = ifres15986


} else {
ifres15985 = False


}

ifres15984 = ifres15985


} else {
ifres15984 = False


}

ifres15983 = ifres15984


} else {
ifres15983 = False


}

if True == ifres15983 {
tmp15900 := GTl(__e, V5529)

tmp15901 := GHd(__e, tmp15900)

tmp15902 := GTl(__e, V5529)

tmp15903 := GTl(__e, tmp15902)

tmp15904 := GHd(__e, tmp15903)

tmp15905 := GTl(__e, tmp15904)

tmp15906 := GCons(__e, tmp15905, Nil)

tmp15907 := GCons(__e, symnot, tmp15906)

tmp15908 := GCons(__e, tmp15907, Nil)

tmp15909 := GCons(__e, tmp15901, tmp15908)

tmp15910 := GCons(__e, symand, tmp15909)

tmp15911 := GTl(__e, V5529)

tmp15912 := GTl(__e, tmp15911)

tmp15913 := GHd(__e, tmp15912)

tmp15914 := GTl(__e, tmp15913)

tmp15915 := GTl(__e, tmp15914)

tmp15916 := GCons(__e, tmp15910, tmp15915)

return GCons(__e, symwhere, tmp15916)


} else {
tmp15981 := GConsP(__e, V5529)

var ifres15966 Obj

if True == tmp15981 {
tmp15979 := GHd(__e, V5529)

tmp15980 := GEq(__e, symwhere, tmp15979)

var ifres15967 Obj

if True == tmp15980 {
tmp15977 := GTl(__e, V5529)

tmp15978 := GConsP(__e, tmp15977)

var ifres15968 Obj

if True == tmp15978 {
tmp15974 := GTl(__e, V5529)

tmp15975 := GTl(__e, tmp15974)

tmp15976 := GConsP(__e, tmp15975)

var ifres15969 Obj

if True == tmp15976 {
tmp15970 := GTl(__e, V5529)

tmp15971 := GTl(__e, tmp15970)

tmp15972 := GTl(__e, tmp15971)

tmp15973 := GEq(__e, Nil, tmp15972)

ifres15969 = tmp15973


} else {
ifres15969 = False


}

ifres15968 = ifres15969


} else {
ifres15968 = False


}

ifres15967 = ifres15968


} else {
ifres15967 = False


}

ifres15966 = ifres15967


} else {
ifres15966 = False


}

if True == ifres15966 {
tmp15917 := GTl(__e, V5529)

tmp15918 := GHd(__e, tmp15917)

tmp15919 := GTl(__e, V5529)

tmp15920 := GTl(__e, tmp15919)

tmp15921 := GHd(__e, tmp15920)

tmp15922 := GCons(__e, symfail, Nil)

tmp15923 := GCons(__e, tmp15922, Nil)

tmp15924 := GCons(__e, tmp15921, tmp15923)

tmp15925 := GCons(__e, sym_a_a, tmp15924)

tmp15926 := GCons(__e, tmp15925, Nil)

tmp15927 := GCons(__e, symnot, tmp15926)

tmp15928 := GCons(__e, tmp15927, Nil)

tmp15929 := GCons(__e, tmp15918, tmp15928)

tmp15930 := GCons(__e, symand, tmp15929)

tmp15931 := GTl(__e, V5529)

tmp15932 := GTl(__e, tmp15931)

tmp15933 := GCons(__e, tmp15930, tmp15932)

return GCons(__e, symwhere, tmp15933)


} else {
tmp15964 := GConsP(__e, V5529)

var ifres15949 Obj

if True == tmp15964 {
tmp15962 := GHd(__e, V5529)

tmp15963 := GEq(__e, symfail_1if, tmp15962)

var ifres15950 Obj

if True == tmp15963 {
tmp15960 := GTl(__e, V5529)

tmp15961 := GConsP(__e, tmp15960)

var ifres15951 Obj

if True == tmp15961 {
tmp15957 := GTl(__e, V5529)

tmp15958 := GTl(__e, tmp15957)

tmp15959 := GConsP(__e, tmp15958)

var ifres15952 Obj

if True == tmp15959 {
tmp15953 := GTl(__e, V5529)

tmp15954 := GTl(__e, tmp15953)

tmp15955 := GTl(__e, tmp15954)

tmp15956 := GEq(__e, Nil, tmp15955)

ifres15952 = tmp15956


} else {
ifres15952 = False


}

ifres15951 = ifres15952


} else {
ifres15951 = False


}

ifres15950 = ifres15951


} else {
ifres15950 = False


}

ifres15949 = ifres15950


} else {
ifres15949 = False


}

if True == ifres15949 {
tmp15934 := GTl(__e, V5529)

tmp15935 := GCons(__e, tmp15934, Nil)

tmp15936 := GCons(__e, symnot, tmp15935)

tmp15937 := GTl(__e, V5529)

tmp15938 := GTl(__e, tmp15937)

tmp15939 := GCons(__e, tmp15936, tmp15938)

return GCons(__e, symwhere, tmp15939)


} else {
tmp15940 := GCons(__e, symfail, Nil)

tmp15941 := GCons(__e, tmp15940, Nil)

tmp15942 := GCons(__e, V5529, tmp15941)

tmp15943 := GCons(__e, sym_a_a, tmp15942)

tmp15944 := GCons(__e, tmp15943, Nil)

tmp15945 := GCons(__e, symnot, tmp15944)

tmp15946 := GCons(__e, V5529, Nil)

tmp15947 := GCons(__e, tmp15945, tmp15946)

return GCons(__e, symwhere, tmp15947)


}


}


}


panic("unreachable")
})

tmp16031 := Call2(__e, ns2_1set, symshen_4correct, tmp15899)


_ = tmp16031

tmp16032 := MakeNative(func(__e *ControlFlow) {
V5530 := __e.Get(1)
V5531 := __e.Get(2)
V5532 := __e.Get(3)
V5533 := __e.Get(4)
V5534 := __e.Get(5)
V5535 := __e.Get(6)
V5536 := __e.Get(7)
V5537 := __e.Get(8)
__e.ReturnDirect(func() Obj {
let__13626 := GAdd(__e, V5536, MakeInteger(1))
_ = let__13626

tmp16040 := Call1(__e, SymFn(symshen_4unlocked_2), V5535)


var ifres16034 Obj

if True == tmp16040 {
tmp16035 := Call2(__e, SymFn(symshen_4lazyderef), V5531, V5534)


let__13628 := tmp16035
_ = let__13628

tmp16039 := GEq(__e, let__13628, Nil)

var ifres16036 Obj

if True == tmp16039 {
tmp16037 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16037

tmp16038 := Call1(__e, SymFn(symthaw), V5537)


ifres16036 = tmp16038


} else {
ifres16036 = False


}

ifres16034 = ifres16036


} else {
ifres16034 = False


}

let__13627 := ifres16034
_ = let__13627

tmp16065 := GEq(__e, let__13627, False)

if True == tmp16065 {
tmp16061 := Call1(__e, SymFn(symshen_4unlocked_2), V5535)


var ifres16041 Obj

if True == tmp16061 {
tmp16042 := Call2(__e, SymFn(symshen_4lazyderef), V5531, V5534)


let__13630 := tmp16042
_ = let__13630

tmp16060 := GConsP(__e, let__13630)

var ifres16043 Obj

if True == tmp16060 {
tmp16044 := GHd(__e, let__13630)

let__13631 := tmp16044
_ = let__13631

tmp16045 := GTl(__e, let__13630)

let__13632 := tmp16045
_ = let__13632

tmp16046 := Call1(__e, SymFn(symshen_4newpv), V5534)


let__13633 := tmp16046
_ = let__13633

tmp16047 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16047

tmp16048 := Call2(__e, SymFn(symshen_4deref), let__13631, V5534)


tmp16049 := Call1(__e, SymFn(symshen_4freshen_1rule), tmp16048)


tmp16050 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16051 := Call2(__e, SymFn(symshen_4lazyderef), let__13633, V5534)


tmp16052 := Call1(__e, SymFn(symfst), tmp16051)


tmp16053 := Call2(__e, SymFn(symshen_4lazyderef), let__13633, V5534)


tmp16054 := Call1(__e, SymFn(symsnd), tmp16053)


tmp16055 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16056 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4t_d_1rules), V5530, let__13632, V5532, GAdd(__e, V5533, MakeInteger(1)), V5534, V5535, let__13626, V5537)
return nil


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5534, V5535, let__13626, tmp16056)


panic("unreachable")
})

__e.TailApply(SymFn(symshen_4t_d_1rule), V5530, V5533, tmp16052, tmp16054, V5532, V5534, V5535, let__13626, tmp16055)
return nil


panic("unreachable")
})

tmp16058 := Call6(__e, SymFn(symbind), let__13633, tmp16049, V5534, V5535, let__13626, tmp16050)


tmp16059 := Call2(__e, SymFn(symshen_4gc), V5534, tmp16058)


ifres16043 = tmp16059


} else {
ifres16043 = False


}

ifres16041 = ifres16043


} else {
ifres16041 = False


}

let__13629 := ifres16041
_ = let__13629

tmp16063 := GEq(__e, let__13629, False)

if True == tmp16063 {
return TailCall2(__e, SymFn(symshen_4unlock), V5535, let__13626)
} else {
return let__13629
}


} else {
return let__13627
}


panic("unreachable")
}())
}, 8)

tmp16066 := Call2(__e, ns2_1set, symshen_4t_d_1rules, tmp16032)


_ = tmp16066

tmp16067 := MakeNative1(func(__e *ControlFlow, V5546 Obj) Obj {
tmp16078 := Call1(__e, SymFn(symtuple_2), V5546)


if True == tmp16078 {
tmp16068 := Call1(__e, SymFn(symfst), V5546)


tmp16069 := Call1(__e, SymFn(symshen_4extract_1vars), tmp16068)


let__13634 := tmp16069
_ = let__13634

tmp16070 := MakeNative1(func(__e *ControlFlow, Z5549 Obj) Obj {
tmp16071 := Call1(__e, SymFn(symshen_4freshterm), Z5549)


return GCons(__e, Z5549, tmp16071)


panic("unreachable")
})

tmp16072 := Call2(__e, SymFn(symmap), tmp16070, let__13634)


let__13635 := tmp16072
_ = let__13635

tmp16073 := Call1(__e, SymFn(symfst), V5546)


tmp16074 := Call2(__e, SymFn(symshen_4freshen), let__13635, tmp16073)


tmp16075 := Call1(__e, SymFn(symsnd), V5546)


tmp16076 := Call2(__e, SymFn(symshen_4freshen), let__13635, tmp16075)


return TailCall2(__e, SymFn(sym_8p), tmp16074, tmp16076)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4freshen_1rule)
}


panic("unreachable")
})

tmp16079 := Call2(__e, ns2_1set, symshen_4freshen_1rule, tmp16067)


_ = tmp16079

tmp16080 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V5550 := __self1
_ = V5550
V5551 := __self2
_ = V5551
tmp16093 := GEq(__e, Nil, V5550)

if True == tmp16093 {
return V5551
} else {
tmp16091 := GConsP(__e, V5550)

var ifres16088 Obj

if True == tmp16091 {
tmp16089 := GHd(__e, V5550)

tmp16090 := GConsP(__e, tmp16089)

ifres16088 = tmp16090


} else {
ifres16088 = False


}

if True == ifres16088 {
tmp16081 := GTl(__e, V5550)

tmp16082 := GHd(__e, V5550)

tmp16083 := GHd(__e, tmp16082)

tmp16084 := GHd(__e, V5550)

tmp16085 := GTl(__e, tmp16084)

tmp16086 := Call3(__e, SymFn(symshen_4beta), tmp16083, tmp16085, V5551)


if SymFn(symshen_4freshen) == __self {
__self1, __self2 = tmp16081, tmp16086
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4freshen), tmp16081, tmp16086)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4freshen)
}


}


panic("unreachable")
})
return __self
})()

tmp16094 := Call2(__e, ns2_1set, symshen_4freshen, tmp16080)


_ = tmp16094

tmp16095 := MakeNative(func(__e *ControlFlow) {
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
tmp16099 := Call1(__e, SymFn(symshen_4unlocked_2), V5558)


var ifres16096 Obj

if True == tmp16099 {
tmp16097 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16097

tmp16098 := Call(__e, SymFn(symshen_4t_d_1rule_1h), V5554, V5555, V5556, V5557, V5558, V5559, V5560)


ifres16096 = tmp16098


} else {
ifres16096 = False


}

let__13636 := ifres16096
_ = let__13636

tmp16111 := GEq(__e, let__13636, False)

if True == tmp16111 {
tmp16109 := Call1(__e, SymFn(symshen_4unlocked_2), V5558)


if True == tmp16109 {
tmp16100 := Call1(__e, SymFn(symshen_4newpv), V5557)


let__13637 := tmp16100
_ = let__13637

tmp16101 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16101

tmp16102 := Call3(__e, SymFn(symshen_4app), V5552, MakeString("\n"), symshen_4a)


tmp16104 := Call3(__e, SymFn(symshen_4app), V5553, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" of "), tmp16102), symshen_4a)


tmp16106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("type error in rule "), tmp16104))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("type error in rule "), tmp16104)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()

tmp16107 := Call6(__e, SymFn(symbind), let__13637, tmp16106, V5557, V5558, V5559, V5560)


return TailCall2(__e, SymFn(symshen_4gc), V5557, tmp16107)


} else {
return False
}


} else {
return let__13636
}


panic("unreachable")
}())
}, 9)

tmp16112 := Call2(__e, ns2_1set, symshen_4t_d_1rule, tmp16095)


_ = tmp16112

tmp16113 := MakeNative(func(__e *ControlFlow) {
V5563 := __e.Get(1)
V5564 := __e.Get(2)
V5565 := __e.Get(3)
V5566 := __e.Get(4)
V5567 := __e.Get(5)
V5568 := __e.Get(6)
V5569 := __e.Get(7)
__e.ReturnDirect(func() Obj {
let__13638 := GAdd(__e, V5568, MakeInteger(1))
_ = let__13638

tmp16138 := Call1(__e, SymFn(symshen_4unlocked_2), V5567)


var ifres16115 Obj

if True == tmp16138 {
tmp16116 := Call2(__e, SymFn(symshen_4lazyderef), V5563, V5566)


let__13640 := tmp16116
_ = let__13640

tmp16137 := GEq(__e, let__13640, Nil)

var ifres16117 Obj

if True == tmp16137 {
tmp16118 := Call2(__e, SymFn(symshen_4lazyderef), V5565, V5566)


let__13641 := tmp16118
_ = let__13641

tmp16136 := GConsP(__e, let__13641)

var ifres16119 Obj

if True == tmp16136 {
tmp16120 := GHd(__e, let__13641)

tmp16121 := Call2(__e, SymFn(symshen_4lazyderef), tmp16120, V5566)


let__13642 := tmp16121
_ = let__13642

tmp16135 := GEq(__e, let__13642, sym_1_1_6)

var ifres16122 Obj

if True == tmp16135 {
tmp16123 := GTl(__e, let__13641)

tmp16124 := Call2(__e, SymFn(symshen_4lazyderef), tmp16123, V5566)


let__13643 := tmp16124
_ = let__13643

tmp16134 := GConsP(__e, let__13643)

var ifres16125 Obj

if True == tmp16134 {
tmp16126 := GHd(__e, let__13643)

let__13644 := tmp16126
_ = let__13644

tmp16127 := GTl(__e, let__13643)

tmp16128 := Call2(__e, SymFn(symshen_4lazyderef), tmp16127, V5566)


let__13645 := tmp16128
_ = let__13645

tmp16133 := GEq(__e, let__13645, Nil)

var ifres16129 Obj

if True == tmp16133 {
tmp16130 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16130

tmp16131 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4t_d_1correct), V5564, let__13644, Nil, V5566, V5567, let__13638, V5569)
return nil
panic("unreachable")
})

tmp16132 := Call4(__e, SymFn(symshen_4cut), V5566, V5567, let__13638, tmp16131)


ifres16129 = tmp16132


} else {
ifres16129 = False


}

ifres16125 = ifres16129


} else {
ifres16125 = False


}

ifres16122 = ifres16125


} else {
ifres16122 = False


}

ifres16119 = ifres16122


} else {
ifres16119 = False


}

ifres16117 = ifres16119


} else {
ifres16117 = False


}

ifres16115 = ifres16117


} else {
ifres16115 = False


}

let__13639 := ifres16115
_ = let__13639

tmp16157 := GEq(__e, let__13639, False)

if True == tmp16157 {
tmp16153 := Call1(__e, SymFn(symshen_4unlocked_2), V5567)


var ifres16139 Obj

if True == tmp16153 {
tmp16140 := Call1(__e, SymFn(symshen_4newpv), V5566)


let__13647 := tmp16140
_ = let__13647

tmp16141 := Call1(__e, SymFn(symshen_4newpv), V5566)


let__13648 := tmp16141
_ = let__13648

tmp16142 := Call1(__e, SymFn(symshen_4newpv), V5566)


let__13649 := tmp16142
_ = let__13649

tmp16143 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16143

tmp16144 := Call1(__e, SymFn(symshen_4freshterms), V5563)


tmp16145 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16146 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16147 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16148 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4t_d_1correct), V5564, let__13648, let__13649, V5566, V5567, let__13638, V5569)
return nil
panic("unreachable")
})

__e.TailApply(SymFn(symshen_4myassume), V5563, V5565, let__13649, V5566, V5567, let__13638, tmp16148)
return nil


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5566, V5567, let__13638, tmp16147)


panic("unreachable")
})

__e.TailApply(SymFn(symshen_4t_d_1integrity), V5563, V5565, let__13647, let__13648, V5566, V5567, let__13638, tmp16146)
return nil


panic("unreachable")
})

tmp16149 := Call6(__e, SymFn(symshen_4p_1hyps), tmp16144, let__13647, V5566, V5567, let__13638, tmp16145)


tmp16150 := Call2(__e, SymFn(symshen_4gc), V5566, tmp16149)


tmp16151 := Call2(__e, SymFn(symshen_4gc), V5566, tmp16150)


tmp16152 := Call2(__e, SymFn(symshen_4gc), V5566, tmp16151)


ifres16139 = tmp16152


} else {
ifres16139 = False


}

let__13646 := ifres16139
_ = let__13646

tmp16155 := GEq(__e, let__13646, False)

if True == tmp16155 {
return TailCall2(__e, SymFn(symshen_4unlock), V5567, let__13638)
} else {
return let__13646
}


} else {
return let__13639
}


panic("unreachable")
}())
}, 7)

tmp16158 := Call2(__e, ns2_1set, symshen_4t_d_1rule_1h, tmp16113)


_ = tmp16158

tmp16159 := MakeNative(func(__e *ControlFlow) {
V5582 := __e.Get(1)
V5583 := __e.Get(2)
V5584 := __e.Get(3)
V5585 := __e.Get(4)
V5586 := __e.Get(5)
V5587 := __e.Get(6)
V5588 := __e.Get(7)
__e.ReturnDirect(func() Obj {
tmp16173 := Call1(__e, SymFn(symshen_4unlocked_2), V5586)


var ifres16160 Obj

if True == tmp16173 {
tmp16161 := Call2(__e, SymFn(symshen_4lazyderef), V5582, V5585)


let__13651 := tmp16161
_ = let__13651

tmp16172 := GEq(__e, let__13651, Nil)

var ifres16162 Obj

if True == tmp16172 {
tmp16163 := Call2(__e, SymFn(symshen_4lazyderef), V5584, V5585)


let__13652 := tmp16163
_ = let__13652

tmp16164 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16165 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16165

return TailCall1(__e, SymFn(symthaw), V5588)


panic("unreachable")
})

let__13653 := tmp16164
_ = let__13653

tmp16171 := GEq(__e, let__13652, Nil)

var ifres16166 Obj

if True == tmp16171 {
tmp16167 := Call1(__e, SymFn(symthaw), let__13653)


ifres16166 = tmp16167


} else {
tmp16170 := Call1(__e, SymFn(symshen_4pvar_2), let__13652)


var ifres16168 Obj

if True == tmp16170 {
tmp16169 := Call4(__e, SymFn(symshen_4bind_b), let__13652, Nil, V5585, let__13653)


ifres16168 = tmp16169


} else {
ifres16168 = False


}

ifres16166 = ifres16168


}

ifres16162 = ifres16166


} else {
ifres16162 = False


}

ifres16160 = ifres16162


} else {
ifres16160 = False


}

let__13650 := ifres16160
_ = let__13650

tmp16292 := GEq(__e, let__13650, False)

if True == tmp16292 {
tmp16290 := Call1(__e, SymFn(symshen_4unlocked_2), V5586)


if True == tmp16290 {
tmp16174 := Call2(__e, SymFn(symshen_4lazyderef), V5582, V5585)


let__13654 := tmp16174
_ = let__13654

tmp16288 := GConsP(__e, let__13654)

if True == tmp16288 {
tmp16175 := GHd(__e, let__13654)

let__13655 := tmp16175
_ = let__13655

tmp16176 := GTl(__e, let__13654)

let__13656 := tmp16176
_ = let__13656

tmp16177 := Call2(__e, SymFn(symshen_4lazyderef), V5583, V5585)


let__13657 := tmp16177
_ = let__13657

tmp16286 := GConsP(__e, let__13657)

if True == tmp16286 {
tmp16178 := GHd(__e, let__13657)

let__13658 := tmp16178
_ = let__13658

tmp16179 := GTl(__e, let__13657)

tmp16180 := Call2(__e, SymFn(symshen_4lazyderef), tmp16179, V5585)


let__13659 := tmp16180
_ = let__13659

tmp16284 := GConsP(__e, let__13659)

if True == tmp16284 {
tmp16181 := GHd(__e, let__13659)

tmp16182 := Call2(__e, SymFn(symshen_4lazyderef), tmp16181, V5585)


let__13660 := tmp16182
_ = let__13660

tmp16282 := GEq(__e, let__13660, sym_1_1_6)

if True == tmp16282 {
tmp16183 := GTl(__e, let__13659)

tmp16184 := Call2(__e, SymFn(symshen_4lazyderef), tmp16183, V5585)


let__13661 := tmp16184
_ = let__13661

tmp16280 := GConsP(__e, let__13661)

if True == tmp16280 {
tmp16185 := GHd(__e, let__13661)

let__13662 := tmp16185
_ = let__13662

tmp16186 := GTl(__e, let__13661)

tmp16187 := Call2(__e, SymFn(symshen_4lazyderef), tmp16186, V5585)


let__13663 := tmp16187
_ = let__13663

tmp16278 := GEq(__e, let__13663, Nil)

if True == tmp16278 {
tmp16188 := Call2(__e, SymFn(symshen_4lazyderef), V5584, V5585)


let__13664 := tmp16188
_ = let__13664

tmp16189 := MakeNative1(func(__e *ControlFlow, Z5605 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5606 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5607 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5608 Obj) Obj {
tmp16190 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16190

tmp16191 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16192 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16193 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp16194 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4myassume), let__13656, let__13662, Z5608, V5585, V5586, V5587, V5588)
return nil
panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), Z5606, tmp16193, V5585, V5586, V5587, tmp16194)


panic("unreachable")
})

return TailCall6(__e, SymFn(symis_b), let__13655, Z5605, V5585, V5586, V5587, tmp16192)


panic("unreachable")
})

return TailCall6(__e, SymFn(symis_b), let__13658, Z5607, V5585, V5586, V5587, tmp16191)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

let__13665 := tmp16189
_ = let__13665

tmp16276 := GConsP(__e, let__13664)

if True == tmp16276 {
tmp16195 := GHd(__e, let__13664)

tmp16196 := Call2(__e, SymFn(symshen_4lazyderef), tmp16195, V5585)


let__13666 := tmp16196
_ = let__13666

tmp16197 := MakeNative1(func(__e *ControlFlow, Z5611 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5612 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5613 Obj) Obj {
tmp16198 := GTl(__e, let__13664)

let__13668 := tmp16198
_ = let__13668

tmp16199 := Call1(__e, let__13665, Z5611)


tmp16200 := Call1(__e, tmp16199, Z5612)


tmp16201 := Call1(__e, tmp16200, Z5613)


return TailCall1(__e, tmp16201, let__13668)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

let__13667 := tmp16197
_ = let__13667

tmp16256 := GConsP(__e, let__13666)

if True == tmp16256 {
tmp16202 := GHd(__e, let__13666)

let__13669 := tmp16202
_ = let__13669

tmp16203 := GTl(__e, let__13666)

tmp16204 := Call2(__e, SymFn(symshen_4lazyderef), tmp16203, V5585)


let__13670 := tmp16204
_ = let__13670

tmp16205 := MakeNative1(func(__e *ControlFlow, Z5618 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5619 Obj) Obj {
tmp16206 := Call1(__e, let__13667, let__13669)


tmp16207 := Call1(__e, tmp16206, Z5618)


return TailCall1(__e, tmp16207, Z5619)


panic("unreachable")
})
panic("unreachable")
})

let__13671 := tmp16205
_ = let__13671

tmp16240 := GConsP(__e, let__13670)

if True == tmp16240 {
tmp16208 := GHd(__e, let__13670)

let__13672 := tmp16208
_ = let__13672

tmp16209 := GTl(__e, let__13670)

tmp16210 := Call2(__e, SymFn(symshen_4lazyderef), tmp16209, V5585)


let__13673 := tmp16210
_ = let__13673

tmp16211 := MakeNative1(func(__e *ControlFlow, Z5623 Obj) Obj {
tmp16212 := Call1(__e, let__13671, let__13672)


return TailCall1(__e, tmp16212, Z5623)


panic("unreachable")
})

let__13674 := tmp16211
_ = let__13674

tmp16228 := GConsP(__e, let__13673)

if True == tmp16228 {
tmp16213 := GHd(__e, let__13673)

let__13675 := tmp16213
_ = let__13675

tmp16214 := GTl(__e, let__13673)

tmp16215 := Call2(__e, SymFn(symshen_4lazyderef), tmp16214, V5585)


let__13676 := tmp16215
_ = let__13676

tmp16216 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13674, let__13675)
panic("unreachable")
})

let__13677 := tmp16216
_ = let__13677

tmp16220 := GEq(__e, let__13676, Nil)

if True == tmp16220 {
return TailCall1(__e, SymFn(symthaw), let__13677)
} else {
tmp16218 := Call1(__e, SymFn(symshen_4pvar_2), let__13676)


if True == tmp16218 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__13676, Nil, V5585, let__13677)
} else {
return False
}


}


} else {
tmp16226 := Call1(__e, SymFn(symshen_4pvar_2), let__13673)


if True == tmp16226 {
tmp16221 := Call1(__e, SymFn(symshen_4newpv), V5585)


let__13678 := tmp16221
_ = let__13678

tmp16222 := GCons(__e, let__13678, Nil)

tmp16223 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13674, let__13678)
panic("unreachable")
})

tmp16224 := Call4(__e, SymFn(symshen_4bind_b), let__13673, tmp16222, V5585, tmp16223)


return TailCall2(__e, SymFn(symshen_4gc), V5585, tmp16224)


} else {
return False
}


}


} else {
tmp16238 := Call1(__e, SymFn(symshen_4pvar_2), let__13670)


if True == tmp16238 {
tmp16229 := Call1(__e, SymFn(symshen_4newpv), V5585)


let__13679 := tmp16229
_ = let__13679

tmp16230 := Call1(__e, SymFn(symshen_4newpv), V5585)


let__13680 := tmp16230
_ = let__13680

tmp16231 := GCons(__e, let__13680, Nil)

tmp16232 := GCons(__e, let__13679, tmp16231)

tmp16233 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16234 := Call1(__e, let__13671, let__13679)


return TailCall1(__e, tmp16234, let__13680)


panic("unreachable")
})

tmp16235 := Call4(__e, SymFn(symshen_4bind_b), let__13670, tmp16232, V5585, tmp16233)


tmp16236 := Call2(__e, SymFn(symshen_4gc), V5585, tmp16235)


return TailCall2(__e, SymFn(symshen_4gc), V5585, tmp16236)


} else {
return False
}


}


} else {
tmp16254 := Call1(__e, SymFn(symshen_4pvar_2), let__13666)


if True == tmp16254 {
tmp16241 := Call1(__e, SymFn(symshen_4newpv), V5585)


let__13681 := tmp16241
_ = let__13681

tmp16242 := Call1(__e, SymFn(symshen_4newpv), V5585)


let__13682 := tmp16242
_ = let__13682

tmp16243 := Call1(__e, SymFn(symshen_4newpv), V5585)


let__13683 := tmp16243
_ = let__13683

tmp16244 := GCons(__e, let__13683, Nil)

tmp16245 := GCons(__e, let__13682, tmp16244)

tmp16246 := GCons(__e, let__13681, tmp16245)

tmp16247 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16248 := Call1(__e, let__13667, let__13681)


tmp16249 := Call1(__e, tmp16248, let__13682)


return TailCall1(__e, tmp16249, let__13683)


panic("unreachable")
})

tmp16250 := Call4(__e, SymFn(symshen_4bind_b), let__13666, tmp16246, V5585, tmp16247)


tmp16251 := Call2(__e, SymFn(symshen_4gc), V5585, tmp16250)


tmp16252 := Call2(__e, SymFn(symshen_4gc), V5585, tmp16251)


return TailCall2(__e, SymFn(symshen_4gc), V5585, tmp16252)


} else {
return False
}


}


} else {
tmp16274 := Call1(__e, SymFn(symshen_4pvar_2), let__13664)


if True == tmp16274 {
tmp16257 := Call1(__e, SymFn(symshen_4newpv), V5585)


let__13684 := tmp16257
_ = let__13684

tmp16258 := Call1(__e, SymFn(symshen_4newpv), V5585)


let__13685 := tmp16258
_ = let__13685

tmp16259 := Call1(__e, SymFn(symshen_4newpv), V5585)


let__13686 := tmp16259
_ = let__13686

tmp16260 := Call1(__e, SymFn(symshen_4newpv), V5585)


let__13687 := tmp16260
_ = let__13687

tmp16261 := GCons(__e, let__13686, Nil)

tmp16262 := GCons(__e, let__13685, tmp16261)

tmp16263 := GCons(__e, let__13684, tmp16262)

tmp16264 := GCons(__e, tmp16263, let__13687)

tmp16265 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16266 := Call1(__e, let__13665, let__13684)


tmp16267 := Call1(__e, tmp16266, let__13685)


tmp16268 := Call1(__e, tmp16267, let__13686)


return TailCall1(__e, tmp16268, let__13687)


panic("unreachable")
})

tmp16269 := Call4(__e, SymFn(symshen_4bind_b), let__13664, tmp16264, V5585, tmp16265)


tmp16270 := Call2(__e, SymFn(symshen_4gc), V5585, tmp16269)


tmp16271 := Call2(__e, SymFn(symshen_4gc), V5585, tmp16270)


tmp16272 := Call2(__e, SymFn(symshen_4gc), V5585, tmp16271)


return TailCall2(__e, SymFn(symshen_4gc), V5585, tmp16272)


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
return let__13650
}


panic("unreachable")
}())
}, 7)

tmp16293 := Call2(__e, ns2_1set, symshen_4myassume, tmp16159)


_ = tmp16293

tmp16294 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V5639 := __self1
_ = V5639
tmp16315 := GEq(__e, Nil, V5639)

if True == tmp16315 {
return Nil
} else {
tmp16313 := GConsP(__e, V5639)

var ifres16310 Obj

if True == tmp16313 {
tmp16311 := GHd(__e, V5639)

tmp16312 := GConsP(__e, tmp16311)

ifres16310 = tmp16312


} else {
ifres16310 = False


}

if True == ifres16310 {
tmp16295 := GHd(__e, V5639)

tmp16296 := GTl(__e, V5639)

tmp16297 := Call2(__e, SymFn(symappend), tmp16295, tmp16296)


if SymFn(symshen_4freshterms) == __self {
__self1 = tmp16297
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4freshterms), tmp16297)


} else {
tmp16308 := GConsP(__e, V5639)

var ifres16305 Obj

if True == tmp16308 {
tmp16306 := GHd(__e, V5639)

tmp16307 := Call1(__e, SymFn(symshen_4freshterm_2), tmp16306)


ifres16305 = tmp16307


} else {
ifres16305 = False


}

if True == ifres16305 {
tmp16298 := GHd(__e, V5639)

tmp16299 := GTl(__e, V5639)

tmp16300 := Call1(__e, SymFn(symshen_4freshterms), tmp16299)


return TailCall2(__e, SymFn(symadjoin), tmp16298, tmp16300)


} else {
tmp16303 := GConsP(__e, V5639)

if True == tmp16303 {
tmp16301 := GTl(__e, V5639)

if SymFn(symshen_4freshterms) == __self {
__self1 = tmp16301
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4freshterms), tmp16301)


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

tmp16316 := Call2(__e, ns2_1set, symshen_4freshterms, tmp16294)


_ = tmp16316

tmp16317 := MakeNative6(func(__e *ControlFlow, V5640, V5641, V5642, V5643, V5644, V5645 Obj) Obj {
tmp16331 := Call1(__e, SymFn(symshen_4unlocked_2), V5643)


var ifres16318 Obj

if True == tmp16331 {
tmp16319 := Call2(__e, SymFn(symshen_4lazyderef), V5640, V5642)


let__13689 := tmp16319
_ = let__13689

tmp16330 := GEq(__e, let__13689, Nil)

var ifres16320 Obj

if True == tmp16330 {
tmp16321 := Call2(__e, SymFn(symshen_4lazyderef), V5641, V5642)


let__13690 := tmp16321
_ = let__13690

tmp16322 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16323 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16323

return TailCall1(__e, SymFn(symthaw), V5645)


panic("unreachable")
})

let__13691 := tmp16322
_ = let__13691

tmp16329 := GEq(__e, let__13690, Nil)

var ifres16324 Obj

if True == tmp16329 {
tmp16325 := Call1(__e, SymFn(symthaw), let__13691)


ifres16324 = tmp16325


} else {
tmp16328 := Call1(__e, SymFn(symshen_4pvar_2), let__13690)


var ifres16326 Obj

if True == tmp16328 {
tmp16327 := Call4(__e, SymFn(symshen_4bind_b), let__13690, Nil, V5642, let__13691)


ifres16326 = tmp16327


} else {
ifres16326 = False


}

ifres16324 = ifres16326


}

ifres16320 = ifres16324


} else {
ifres16320 = False


}

ifres16318 = ifres16320


} else {
ifres16318 = False


}

let__13688 := ifres16318
_ = let__13688

tmp16428 := GEq(__e, let__13688, False)

if True == tmp16428 {
tmp16426 := Call1(__e, SymFn(symshen_4unlocked_2), V5643)


if True == tmp16426 {
tmp16332 := Call2(__e, SymFn(symshen_4lazyderef), V5640, V5642)


let__13692 := tmp16332
_ = let__13692

tmp16424 := GConsP(__e, let__13692)

if True == tmp16424 {
tmp16333 := GHd(__e, let__13692)

let__13693 := tmp16333
_ = let__13693

tmp16334 := GTl(__e, let__13692)

let__13694 := tmp16334
_ = let__13694

tmp16335 := Call2(__e, SymFn(symshen_4lazyderef), V5641, V5642)


let__13695 := tmp16335
_ = let__13695

tmp16336 := MakeNative1(func(__e *ControlFlow, Z5655 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5656 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5657 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5658 Obj) Obj {
tmp16337 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16337

tmp16338 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16339 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp16340 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall6(__e, SymFn(symshen_4p_1hyps), let__13694, Z5658, V5642, V5643, V5644, V5645)
panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), Z5656, tmp16339, V5642, V5643, V5644, tmp16340)


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), Z5655, let__13693, V5642, V5643, V5644, tmp16338)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

let__13696 := tmp16336
_ = let__13696

tmp16422 := GConsP(__e, let__13695)

if True == tmp16422 {
tmp16341 := GHd(__e, let__13695)

tmp16342 := Call2(__e, SymFn(symshen_4lazyderef), tmp16341, V5642)


let__13697 := tmp16342
_ = let__13697

tmp16343 := MakeNative1(func(__e *ControlFlow, Z5661 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5662 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5663 Obj) Obj {
tmp16344 := GTl(__e, let__13695)

let__13699 := tmp16344
_ = let__13699

tmp16345 := Call1(__e, let__13696, Z5661)


tmp16346 := Call1(__e, tmp16345, Z5662)


tmp16347 := Call1(__e, tmp16346, Z5663)


return TailCall1(__e, tmp16347, let__13699)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

let__13698 := tmp16343
_ = let__13698

tmp16402 := GConsP(__e, let__13697)

if True == tmp16402 {
tmp16348 := GHd(__e, let__13697)

let__13700 := tmp16348
_ = let__13700

tmp16349 := GTl(__e, let__13697)

tmp16350 := Call2(__e, SymFn(symshen_4lazyderef), tmp16349, V5642)


let__13701 := tmp16350
_ = let__13701

tmp16351 := MakeNative1(func(__e *ControlFlow, Z5668 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5669 Obj) Obj {
tmp16352 := Call1(__e, let__13698, let__13700)


tmp16353 := Call1(__e, tmp16352, Z5668)


return TailCall1(__e, tmp16353, Z5669)


panic("unreachable")
})
panic("unreachable")
})

let__13702 := tmp16351
_ = let__13702

tmp16386 := GConsP(__e, let__13701)

if True == tmp16386 {
tmp16354 := GHd(__e, let__13701)

let__13703 := tmp16354
_ = let__13703

tmp16355 := GTl(__e, let__13701)

tmp16356 := Call2(__e, SymFn(symshen_4lazyderef), tmp16355, V5642)


let__13704 := tmp16356
_ = let__13704

tmp16357 := MakeNative1(func(__e *ControlFlow, Z5673 Obj) Obj {
tmp16358 := Call1(__e, let__13702, let__13703)


return TailCall1(__e, tmp16358, Z5673)


panic("unreachable")
})

let__13705 := tmp16357
_ = let__13705

tmp16374 := GConsP(__e, let__13704)

if True == tmp16374 {
tmp16359 := GHd(__e, let__13704)

let__13706 := tmp16359
_ = let__13706

tmp16360 := GTl(__e, let__13704)

tmp16361 := Call2(__e, SymFn(symshen_4lazyderef), tmp16360, V5642)


let__13707 := tmp16361
_ = let__13707

tmp16362 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13705, let__13706)
panic("unreachable")
})

let__13708 := tmp16362
_ = let__13708

tmp16366 := GEq(__e, let__13707, Nil)

if True == tmp16366 {
return TailCall1(__e, SymFn(symthaw), let__13708)
} else {
tmp16364 := Call1(__e, SymFn(symshen_4pvar_2), let__13707)


if True == tmp16364 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__13707, Nil, V5642, let__13708)
} else {
return False
}


}


} else {
tmp16372 := Call1(__e, SymFn(symshen_4pvar_2), let__13704)


if True == tmp16372 {
tmp16367 := Call1(__e, SymFn(symshen_4newpv), V5642)


let__13709 := tmp16367
_ = let__13709

tmp16368 := GCons(__e, let__13709, Nil)

tmp16369 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13705, let__13709)
panic("unreachable")
})

tmp16370 := Call4(__e, SymFn(symshen_4bind_b), let__13704, tmp16368, V5642, tmp16369)


return TailCall2(__e, SymFn(symshen_4gc), V5642, tmp16370)


} else {
return False
}


}


} else {
tmp16384 := Call1(__e, SymFn(symshen_4pvar_2), let__13701)


if True == tmp16384 {
tmp16375 := Call1(__e, SymFn(symshen_4newpv), V5642)


let__13710 := tmp16375
_ = let__13710

tmp16376 := Call1(__e, SymFn(symshen_4newpv), V5642)


let__13711 := tmp16376
_ = let__13711

tmp16377 := GCons(__e, let__13711, Nil)

tmp16378 := GCons(__e, let__13710, tmp16377)

tmp16379 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16380 := Call1(__e, let__13702, let__13710)


return TailCall1(__e, tmp16380, let__13711)


panic("unreachable")
})

tmp16381 := Call4(__e, SymFn(symshen_4bind_b), let__13701, tmp16378, V5642, tmp16379)


tmp16382 := Call2(__e, SymFn(symshen_4gc), V5642, tmp16381)


return TailCall2(__e, SymFn(symshen_4gc), V5642, tmp16382)


} else {
return False
}


}


} else {
tmp16400 := Call1(__e, SymFn(symshen_4pvar_2), let__13697)


if True == tmp16400 {
tmp16387 := Call1(__e, SymFn(symshen_4newpv), V5642)


let__13712 := tmp16387
_ = let__13712

tmp16388 := Call1(__e, SymFn(symshen_4newpv), V5642)


let__13713 := tmp16388
_ = let__13713

tmp16389 := Call1(__e, SymFn(symshen_4newpv), V5642)


let__13714 := tmp16389
_ = let__13714

tmp16390 := GCons(__e, let__13714, Nil)

tmp16391 := GCons(__e, let__13713, tmp16390)

tmp16392 := GCons(__e, let__13712, tmp16391)

tmp16393 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16394 := Call1(__e, let__13698, let__13712)


tmp16395 := Call1(__e, tmp16394, let__13713)


return TailCall1(__e, tmp16395, let__13714)


panic("unreachable")
})

tmp16396 := Call4(__e, SymFn(symshen_4bind_b), let__13697, tmp16392, V5642, tmp16393)


tmp16397 := Call2(__e, SymFn(symshen_4gc), V5642, tmp16396)


tmp16398 := Call2(__e, SymFn(symshen_4gc), V5642, tmp16397)


return TailCall2(__e, SymFn(symshen_4gc), V5642, tmp16398)


} else {
return False
}


}


} else {
tmp16420 := Call1(__e, SymFn(symshen_4pvar_2), let__13695)


if True == tmp16420 {
tmp16403 := Call1(__e, SymFn(symshen_4newpv), V5642)


let__13715 := tmp16403
_ = let__13715

tmp16404 := Call1(__e, SymFn(symshen_4newpv), V5642)


let__13716 := tmp16404
_ = let__13716

tmp16405 := Call1(__e, SymFn(symshen_4newpv), V5642)


let__13717 := tmp16405
_ = let__13717

tmp16406 := Call1(__e, SymFn(symshen_4newpv), V5642)


let__13718 := tmp16406
_ = let__13718

tmp16407 := GCons(__e, let__13717, Nil)

tmp16408 := GCons(__e, let__13716, tmp16407)

tmp16409 := GCons(__e, let__13715, tmp16408)

tmp16410 := GCons(__e, tmp16409, let__13718)

tmp16411 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16412 := Call1(__e, let__13696, let__13715)


tmp16413 := Call1(__e, tmp16412, let__13716)


tmp16414 := Call1(__e, tmp16413, let__13717)


return TailCall1(__e, tmp16414, let__13718)


panic("unreachable")
})

tmp16415 := Call4(__e, SymFn(symshen_4bind_b), let__13695, tmp16410, V5642, tmp16411)


tmp16416 := Call2(__e, SymFn(symshen_4gc), V5642, tmp16415)


tmp16417 := Call2(__e, SymFn(symshen_4gc), V5642, tmp16416)


tmp16418 := Call2(__e, SymFn(symshen_4gc), V5642, tmp16417)


return TailCall2(__e, SymFn(symshen_4gc), V5642, tmp16418)


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
return let__13688
}


panic("unreachable")
})

tmp16429 := Call2(__e, ns2_1set, symshen_4p_1hyps, tmp16317)


_ = tmp16429

tmp16430 := MakeNative(func(__e *ControlFlow) {
V5687 := __e.Get(1)
V5688 := __e.Get(2)
V5689 := __e.Get(3)
V5690 := __e.Get(4)
V5691 := __e.Get(5)
V5692 := __e.Get(6)
V5693 := __e.Get(7)
__e.ReturnDirect(func() Obj {
let__13719 := GAdd(__e, V5692, MakeInteger(1))
_ = let__13719

tmp16468 := Call1(__e, SymFn(symshen_4unlocked_2), V5691)


var ifres16432 Obj

if True == tmp16468 {
tmp16433 := Call2(__e, SymFn(symshen_4lazyderef), V5687, V5690)


let__13721 := tmp16433
_ = let__13721

tmp16467 := GConsP(__e, let__13721)

var ifres16434 Obj

if True == tmp16467 {
tmp16435 := GHd(__e, let__13721)

tmp16436 := Call2(__e, SymFn(symshen_4lazyderef), tmp16435, V5690)


let__13722 := tmp16436
_ = let__13722

tmp16466 := GEq(__e, let__13722, symwhere)

var ifres16437 Obj

if True == tmp16466 {
tmp16438 := GTl(__e, let__13721)

tmp16439 := Call2(__e, SymFn(symshen_4lazyderef), tmp16438, V5690)


let__13723 := tmp16439
_ = let__13723

tmp16465 := GConsP(__e, let__13723)

var ifres16440 Obj

if True == tmp16465 {
tmp16441 := GHd(__e, let__13723)

let__13724 := tmp16441
_ = let__13724

tmp16442 := GTl(__e, let__13723)

tmp16443 := Call2(__e, SymFn(symshen_4lazyderef), tmp16442, V5690)


let__13725 := tmp16443
_ = let__13725

tmp16464 := GConsP(__e, let__13725)

var ifres16444 Obj

if True == tmp16464 {
tmp16445 := GHd(__e, let__13725)

let__13726 := tmp16445
_ = let__13726

tmp16446 := GTl(__e, let__13725)

tmp16447 := Call2(__e, SymFn(symshen_4lazyderef), tmp16446, V5690)


let__13727 := tmp16447
_ = let__13727

tmp16463 := GEq(__e, let__13727, Nil)

var ifres16448 Obj

if True == tmp16463 {
tmp16449 := Call1(__e, SymFn(symshen_4newpv), V5690)


let__13728 := tmp16449
_ = let__13728

tmp16450 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16450

tmp16451 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16452 := Call1(__e, SymFn(symshen_4curry), let__13724)


tmp16453 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16454 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16455 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16456 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp16457 := GCons(__e, symverified, Nil)

tmp16458 := GCons(__e, tmp16456, tmp16457)

tmp16459 := GCons(__e, let__13728, tmp16458)

tmp16460 := GCons(__e, tmp16459, V5689)

__e.TailApply(SymFn(symshen_4t_d_1correct), let__13726, V5688, tmp16460, V5690, V5691, let__13719, V5693)
return nil


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5690, V5691, let__13719, tmp16455)


panic("unreachable")
})

__e.TailApply(SymFn(symshen_4system_1S_1h), let__13728, symboolean, V5689, V5690, V5691, let__13719, tmp16454)
return nil


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), let__13728, tmp16452, V5690, V5691, let__13719, tmp16453)


panic("unreachable")
})

tmp16461 := Call4(__e, SymFn(symshen_4cut), V5690, V5691, let__13719, tmp16451)


tmp16462 := Call2(__e, SymFn(symshen_4gc), V5690, tmp16461)


ifres16448 = tmp16462


} else {
ifres16448 = False


}

ifres16444 = ifres16448


} else {
ifres16444 = False


}

ifres16440 = ifres16444


} else {
ifres16440 = False


}

ifres16437 = ifres16440


} else {
ifres16437 = False


}

ifres16434 = ifres16437


} else {
ifres16434 = False


}

ifres16432 = ifres16434


} else {
ifres16432 = False


}

let__13720 := ifres16432
_ = let__13720

tmp16477 := GEq(__e, let__13720, False)

if True == tmp16477 {
tmp16473 := Call1(__e, SymFn(symshen_4unlocked_2), V5691)


var ifres16469 Obj

if True == tmp16473 {
tmp16470 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16470

tmp16471 := Call1(__e, SymFn(symshen_4curry), V5687)


tmp16472 := Call(__e, SymFn(symshen_4system_1S_1h), tmp16471, V5688, V5689, V5690, V5691, let__13719, V5693)


ifres16469 = tmp16472


} else {
ifres16469 = False


}

let__13729 := ifres16469
_ = let__13729

tmp16475 := GEq(__e, let__13729, False)

if True == tmp16475 {
return TailCall2(__e, SymFn(symshen_4unlock), V5691, let__13719)
} else {
return let__13729
}


} else {
return let__13720
}


panic("unreachable")
}())
}, 7)

tmp16478 := Call2(__e, ns2_1set, symshen_4t_d_1correct, tmp16430)


_ = tmp16478

tmp16479 := MakeNative(func(__e *ControlFlow) {
V5705 := __e.Get(1)
V5706 := __e.Get(2)
V5707 := __e.Get(3)
V5708 := __e.Get(4)
V5709 := __e.Get(5)
V5710 := __e.Get(6)
V5711 := __e.Get(7)
V5712 := __e.Get(8)
__e.ReturnDirect(func() Obj {
tmp16486 := Call1(__e, SymFn(symshen_4unlocked_2), V5710)


var ifres16480 Obj

if True == tmp16486 {
tmp16481 := Call2(__e, SymFn(symshen_4lazyderef), V5705, V5709)


let__13731 := tmp16481
_ = let__13731

tmp16485 := GEq(__e, let__13731, Nil)

var ifres16482 Obj

if True == tmp16485 {
tmp16483 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16483

tmp16484 := Call6(__e, SymFn(symis_b), V5706, V5708, V5709, V5710, V5711, V5712)


ifres16482 = tmp16484


} else {
ifres16482 = False


}

ifres16480 = ifres16482


} else {
ifres16480 = False


}

let__13730 := ifres16480
_ = let__13730

tmp16518 := GEq(__e, let__13730, False)

if True == tmp16518 {
tmp16516 := Call1(__e, SymFn(symshen_4unlocked_2), V5710)


if True == tmp16516 {
tmp16487 := Call2(__e, SymFn(symshen_4lazyderef), V5705, V5709)


let__13732 := tmp16487
_ = let__13732

tmp16514 := GConsP(__e, let__13732)

if True == tmp16514 {
tmp16488 := GHd(__e, let__13732)

let__13733 := tmp16488
_ = let__13733

tmp16489 := GTl(__e, let__13732)

let__13734 := tmp16489
_ = let__13734

tmp16490 := Call2(__e, SymFn(symshen_4lazyderef), V5706, V5709)


let__13735 := tmp16490
_ = let__13735

tmp16512 := GConsP(__e, let__13735)

if True == tmp16512 {
tmp16491 := GHd(__e, let__13735)

let__13736 := tmp16491
_ = let__13736

tmp16492 := GTl(__e, let__13735)

tmp16493 := Call2(__e, SymFn(symshen_4lazyderef), tmp16492, V5709)


let__13737 := tmp16493
_ = let__13737

tmp16510 := GConsP(__e, let__13737)

if True == tmp16510 {
tmp16494 := GHd(__e, let__13737)

tmp16495 := Call2(__e, SymFn(symshen_4lazyderef), tmp16494, V5709)


let__13738 := tmp16495
_ = let__13738

tmp16508 := GEq(__e, let__13738, sym_1_1_6)

if True == tmp16508 {
tmp16496 := GTl(__e, let__13737)

tmp16497 := Call2(__e, SymFn(symshen_4lazyderef), tmp16496, V5709)


let__13739 := tmp16497
_ = let__13739

tmp16506 := GConsP(__e, let__13739)

if True == tmp16506 {
tmp16498 := GHd(__e, let__13739)

let__13740 := tmp16498
_ = let__13740

tmp16499 := GTl(__e, let__13739)

tmp16500 := Call2(__e, SymFn(symshen_4lazyderef), tmp16499, V5709)


let__13741 := tmp16500
_ = let__13741

tmp16504 := GEq(__e, let__13741, Nil)

if True == tmp16504 {
tmp16501 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16501

tmp16502 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4t_d_1integrity), let__13734, let__13740, V5707, V5708, V5709, V5710, V5711, V5712)
return nil
panic("unreachable")
})

__e.TailApply(SymFn(symshen_4system_1S_1h), let__13733, let__13736, V5707, V5709, V5710, V5711, tmp16502)
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
return let__13730
}


panic("unreachable")
}())
}, 8)

tmp16519 := Call2(__e, ns2_1set, symshen_4t_d_1integrity, tmp16479)


_ = tmp16519

tmp16520 := MakeNative1(func(__e *ControlFlow, V5725 Obj) Obj {
tmp16526 := GVectorP(__e, V5725)

if True == tmp16526 {
tmp16523 := PrimGuard1(__e, symstring_2, PrimIsString, V5725)

if True == GNot(__e, tmp16523) {
tmp16521 := GVecGet(__e, V5725, MakeInteger(0))

return GEq(__e, tmp16521, symshen_4print_1freshterm)


} else {
return False
}


} else {
return False
}


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4freshterm_2, tmp16520)




panic("unreachable")
})

