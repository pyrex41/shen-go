package main

import . "github.com/pyrex41/shen-go/kl"

var TStarMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp13762 := MakeNative2(func(__e *ControlFlow, V4930, V4931 Obj) Obj {
tmp13763 := Call1(__e, SymFn(symshen_4extract_1vars), V4931)


let__13211 := tmp13763
_ = let__13211

tmp13764 := Call1(__e, SymFn(symshen_4rectify_1type), V4931)


let__13212 := tmp13764
_ = let__13212

tmp13765 := Call1(__e, SymFn(symshen_4curry), V4930)


let__13213 := tmp13765
_ = let__13213

tmp13766 := MakeNative1(func(__e *ControlFlow, Z4935 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z4936 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z4937 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z4938 Obj) Obj {
tmp13767 := Call1(__e, SymFn(symshen_4newpv), Z4935)


let__13214 := tmp13767
_ = let__13214

tmp13768 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp13768

tmp13769 := Call2(__e, SymFn(symshen_4deref), let__13211, Z4935)


tmp13770 := Call1(__e, SymFn(symreceive), tmp13769)


tmp13771 := Call2(__e, SymFn(symshen_4deref), let__13212, Z4935)


tmp13772 := Call1(__e, SymFn(symreceive), tmp13771)


tmp13773 := MakeNative0(func(__e *ControlFlow) Obj {
tmp13774 := Call2(__e, SymFn(symshen_4deref), let__13213, Z4935)


tmp13775 := Call1(__e, SymFn(symreceive), tmp13774)


tmp13776 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall5(__e, SymFn(symreturn), let__13214, Z4935, Z4936, Z4937, Z4938)
panic("unreachable")
})

return TailCall6(__e, SymFn(symshen_4toplevel_1forms), tmp13775, let__13214, Z4935, Z4936, Z4937, tmp13776)


panic("unreachable")
})

tmp13777 := Call(__e, SymFn(symshen_4insert_1prolog_1variables), tmp13770, tmp13772, let__13214, Z4935, Z4936, Z4937, tmp13773)


return TailCall2(__e, SymFn(symshen_4gc), Z4935, tmp13777)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp13778 := Call0(__e, SymFn(symshen_4prolog_1vector))


tmp13779 := Call1(__e, tmp13766, tmp13778)


tmp13780 := Call1(__e, SymFn(symvector), MakeInteger(0))


tmp13781 := Call2(__e, SymFn(sym_8v), MakeInteger(0), tmp13780)


tmp13782 := Call2(__e, SymFn(sym_8v), True, tmp13781)


tmp13783 := Call1(__e, tmp13779, tmp13782)


tmp13784 := Call1(__e, tmp13783, MakeInteger(0))


tmp13785 := MakeNative0(func(__e *ControlFlow) Obj {
return True
panic("unreachable")
})

return TailCall1(__e, tmp13784, tmp13785)


panic("unreachable")
})

tmp13786 := Call2(__e, ns2_1set, symshen_4typecheck, tmp13762)


_ = tmp13786

tmp13787 := MakeNative(func(__e *ControlFlow) {
V4940 := __e.Get(1)
V4941 := __e.Get(2)
V4942 := __e.Get(3)
V4943 := __e.Get(4)
V4944 := __e.Get(5)
V4945 := __e.Get(6)
V4946 := __e.Get(7)
__e.ReturnDirect(func() Obj {
tmp13794 := Call1(__e, SymFn(symshen_4unlocked_2), V4944)


var ifres13788 Obj

if True == tmp13794 {
tmp13789 := Call2(__e, SymFn(symshen_4lazyderef), V4940, V4943)


let__13216 := tmp13789
_ = let__13216

tmp13793 := GEq(__e, let__13216, Nil)

var ifres13790 Obj

if True == tmp13793 {
tmp13791 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp13791

tmp13792 := Call6(__e, SymFn(symis_b), V4941, V4942, V4943, V4944, V4945, V4946)


ifres13790 = tmp13792


} else {
ifres13790 = False


}

ifres13788 = ifres13790


} else {
ifres13788 = False


}

let__13215 := ifres13788
_ = let__13215

tmp13808 := GEq(__e, let__13215, False)

if True == tmp13808 {
tmp13806 := Call1(__e, SymFn(symshen_4unlocked_2), V4944)


if True == tmp13806 {
tmp13795 := Call2(__e, SymFn(symshen_4lazyderef), V4940, V4943)


let__13217 := tmp13795
_ = let__13217

tmp13804 := GConsP(__e, let__13217)

if True == tmp13804 {
tmp13796 := GHd(__e, let__13217)

let__13218 := tmp13796
_ = let__13218

tmp13797 := GTl(__e, let__13217)

let__13219 := tmp13797
_ = let__13219

tmp13798 := Call1(__e, SymFn(symshen_4newpv), V4943)


let__13220 := tmp13798
_ = let__13220

tmp13799 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp13799

tmp13800 := Call2(__e, SymFn(symshen_4deref), let__13220, V4943)


tmp13801 := Call3(__e, SymFn(symsubst), tmp13800, let__13218, V4941)


tmp13802 := Call(__e, SymFn(symshen_4insert_1prolog_1variables), let__13219, tmp13801, V4942, V4943, V4944, V4945, V4946)


return TailCall2(__e, SymFn(symshen_4gc), V4943, tmp13802)


} else {
return False
}


} else {
return False
}


} else {
return let__13215
}


panic("unreachable")
}())
}, 7)

tmp13809 := Call2(__e, ns2_1set, symshen_4insert_1prolog_1variables, tmp13787)


_ = tmp13809

tmp13810 := MakeNative6(func(__e *ControlFlow, V4953, V4954, V4955, V4956, V4957, V4958 Obj) Obj {
let__13221 := GAdd(__e, V4957, MakeInteger(1))
_ = let__13221

tmp13835 := Call1(__e, SymFn(symshen_4unlocked_2), V4956)


var ifres13812 Obj

if True == tmp13835 {
tmp13813 := Call2(__e, SymFn(symshen_4lazyderef), V4953, V4955)


let__13223 := tmp13813
_ = let__13223

tmp13834 := GConsP(__e, let__13223)

var ifres13814 Obj

if True == tmp13834 {
tmp13815 := GHd(__e, let__13223)

tmp13816 := Call2(__e, SymFn(symshen_4lazyderef), tmp13815, V4955)


let__13224 := tmp13816
_ = let__13224

tmp13833 := GEq(__e, let__13224, symdefine)

var ifres13817 Obj

if True == tmp13833 {
tmp13818 := GTl(__e, let__13223)

tmp13819 := Call2(__e, SymFn(symshen_4lazyderef), tmp13818, V4955)


let__13225 := tmp13819
_ = let__13225

tmp13832 := GConsP(__e, let__13225)

var ifres13820 Obj

if True == tmp13832 {
tmp13821 := GHd(__e, let__13225)

let__13226 := tmp13821
_ = let__13226

tmp13822 := GTl(__e, let__13225)

let__13227 := tmp13822
_ = let__13227

tmp13823 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp13823

tmp13824 := Call0(__e, SymFn(symshen_4type_1theory_1enabled_2))


tmp13825 := MakeNative0(func(__e *ControlFlow) Obj {
tmp13826 := MakeNative0(func(__e *ControlFlow) Obj {
tmp13827 := GValue(__e, symshen_4_dspy_d)

tmp13828 := MakeNative0(func(__e *ControlFlow) Obj {
tmp13829 := GCons(__e, let__13226, let__13227)

tmp13830 := GCons(__e, symdefine, tmp13829)

return TailCall6(__e, SymFn(symshen_4t_d), tmp13830, V4954, V4955, V4956, let__13221, V4958)


panic("unreachable")
})

return TailCall6(__e, SymFn(symshen_4signal_1def), tmp13827, let__13226, V4955, V4956, let__13221, tmp13828)


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V4955, V4956, let__13221, tmp13826)


panic("unreachable")
})

tmp13831 := Call5(__e, SymFn(symwhen), tmp13824, V4955, V4956, let__13221, tmp13825)


ifres13820 = tmp13831


} else {
ifres13820 = False


}

ifres13817 = ifres13820


} else {
ifres13817 = False


}

ifres13814 = ifres13817


} else {
ifres13814 = False


}

ifres13812 = ifres13814


} else {
ifres13812 = False


}

let__13222 := ifres13812
_ = let__13222

tmp13847 := GEq(__e, let__13222, False)

if True == tmp13847 {
tmp13843 := Call1(__e, SymFn(symshen_4unlocked_2), V4956)


var ifres13836 Obj

if True == tmp13843 {
tmp13837 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp13837

tmp13838 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp13839 := GCons(__e, V4954, Nil)

tmp13840 := GCons(__e, tmp13838, tmp13839)

tmp13841 := GCons(__e, V4953, tmp13840)

tmp13842 := Call6(__e, SymFn(symshen_4system_1S), tmp13841, Nil, V4955, V4956, let__13221, V4958)


ifres13836 = tmp13842


} else {
ifres13836 = False


}

let__13228 := ifres13836
_ = let__13228

tmp13845 := GEq(__e, let__13228, False)

if True == tmp13845 {
return TailCall2(__e, SymFn(symshen_4unlock), V4956, let__13221)
} else {
return let__13228
}


} else {
return let__13222
}


panic("unreachable")
})

tmp13848 := Call2(__e, ns2_1set, symshen_4toplevel_1forms, tmp13810)


_ = tmp13848

tmp13849 := MakeNative6(func(__e *ControlFlow, V4967, V4968, V4969, V4970, V4971, V4972 Obj) Obj {
tmp13856 := Call1(__e, SymFn(symshen_4unlocked_2), V4970)


var ifres13850 Obj

if True == tmp13856 {
tmp13851 := Call2(__e, SymFn(symshen_4lazyderef), V4967, V4969)


let__13230 := tmp13851
_ = let__13230

tmp13855 := GEq(__e, let__13230, False)

var ifres13852 Obj

if True == tmp13855 {
tmp13853 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp13853

tmp13854 := Call1(__e, SymFn(symthaw), V4972)


ifres13852 = tmp13854


} else {
ifres13852 = False


}

ifres13850 = ifres13852


} else {
ifres13850 = False


}

let__13229 := ifres13850
_ = let__13229

tmp13871 := GEq(__e, let__13229, False)

if True == tmp13871 {
tmp13869 := Call1(__e, SymFn(symshen_4unlocked_2), V4970)


if True == tmp13869 {
tmp13857 := Call2(__e, SymFn(symshen_4lazyderef), V4967, V4969)


let__13231 := tmp13857
_ = let__13231

tmp13867 := GEq(__e, let__13231, True)

if True == tmp13867 {
tmp13858 := Call1(__e, SymFn(symshen_4newpv), V4969)


let__13232 := tmp13858
_ = let__13232

tmp13859 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp13859

tmp13860 := Call2(__e, SymFn(symshen_4deref), V4968, V4969)


tmp13861 := Call3(__e, SymFn(symshen_4app), tmp13860, MakeString(")\n"), symshen_4a)


tmp13862 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("\ntypechecking (fn "), tmp13861)

tmp13863 := Call0(__e, SymFn(symstoutput))


tmp13864 := Call2(__e, SymFn(sympr), tmp13862, tmp13863)


tmp13865 := Call6(__e, SymFn(symis), let__13232, tmp13864, V4969, V4970, V4971, V4972)


return TailCall2(__e, SymFn(symshen_4gc), V4969, tmp13865)


} else {
return False
}


} else {
return False
}


} else {
return let__13229
}


panic("unreachable")
})

tmp13872 := Call2(__e, ns2_1set, symshen_4signal_1def, tmp13849)


_ = tmp13872

tmp13873 := MakeNative1(func(__e *ControlFlow, V4977 Obj) Obj {
tmp13874 := Call1(__e, SymFn(symshen_4curry_1type), V4977)


return TailCall1(__e, SymFn(symshen_4demodulate), tmp13874)


panic("unreachable")
})

tmp13875 := Call2(__e, ns2_1set, symshen_4rectify_1type, tmp13873)


_ = tmp13875

tmp13876 := MakeNative1(func(__e *ControlFlow, V4978 Obj) Obj {
tmp13877 := MakeNative0(func(__e *ControlFlow) Obj {
tmp13878 := MakeNative1(func(__e *ControlFlow, Z4980 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4demod), Z4980)
panic("unreachable")
})

tmp13879 := Call2(__e, SymFn(symshen_4walk), tmp13878, V4978)


let__13233 := tmp13879
_ = let__13233

tmp13881 := GEq(__e, let__13233, V4978)

if True == tmp13881 {
return V4978
} else {
return TailCall1(__e, SymFn(symshen_4demodulate), let__13233)
}


panic("unreachable")
})

tmp13882 := MakeNative1(func(__e *ControlFlow, Z4981 Obj) Obj {
return V4978
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp13877, tmp13882)


panic("unreachable")
})

tmp13883 := Call2(__e, ns2_1set, symshen_4demodulate, tmp13876)


_ = tmp13883

tmp13884 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V4982 := __self1
_ = V4982
tmp13990 := GConsP(__e, V4982)

var ifres13968 Obj

if True == tmp13990 {
tmp13988 := GTl(__e, V4982)

tmp13989 := GConsP(__e, tmp13988)

var ifres13969 Obj

if True == tmp13989 {
tmp13985 := GTl(__e, V4982)

tmp13986 := GHd(__e, tmp13985)

tmp13987 := GEq(__e, sym_1_1_6, tmp13986)

var ifres13970 Obj

if True == tmp13987 {
tmp13982 := GTl(__e, V4982)

tmp13983 := GTl(__e, tmp13982)

tmp13984 := GConsP(__e, tmp13983)

var ifres13971 Obj

if True == tmp13984 {
tmp13978 := GTl(__e, V4982)

tmp13979 := GTl(__e, tmp13978)

tmp13980 := GTl(__e, tmp13979)

tmp13981 := GConsP(__e, tmp13980)

var ifres13972 Obj

if True == tmp13981 {
tmp13973 := GTl(__e, V4982)

tmp13974 := GTl(__e, tmp13973)

tmp13975 := GTl(__e, tmp13974)

tmp13976 := GHd(__e, tmp13975)

tmp13977 := GEq(__e, sym_1_1_6, tmp13976)

ifres13972 = tmp13977


} else {
ifres13972 = False


}

ifres13971 = ifres13972


} else {
ifres13971 = False


}

ifres13970 = ifres13971


} else {
ifres13970 = False


}

ifres13969 = ifres13970


} else {
ifres13969 = False


}

ifres13968 = ifres13969


} else {
ifres13968 = False


}

if True == ifres13968 {
tmp13885 := GHd(__e, V4982)

tmp13886 := GTl(__e, V4982)

tmp13887 := GTl(__e, tmp13886)

tmp13888 := GCons(__e, tmp13887, Nil)

tmp13889 := GCons(__e, sym_1_1_6, tmp13888)

tmp13890 := GCons(__e, tmp13885, tmp13889)

if SymFn(symshen_4curry_1type) == __self {
__self1 = tmp13890
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4curry_1type), tmp13890)


} else {
tmp13966 := GConsP(__e, V4982)

var ifres13934 Obj

if True == tmp13966 {
tmp13964 := GHd(__e, V4982)

tmp13965 := GConsP(__e, tmp13964)

var ifres13935 Obj

if True == tmp13965 {
tmp13961 := GHd(__e, V4982)

tmp13962 := GHd(__e, tmp13961)

tmp13963 := GEq(__e, symlist, tmp13962)

var ifres13936 Obj

if True == tmp13963 {
tmp13958 := GHd(__e, V4982)

tmp13959 := GTl(__e, tmp13958)

tmp13960 := GConsP(__e, tmp13959)

var ifres13937 Obj

if True == tmp13960 {
tmp13954 := GHd(__e, V4982)

tmp13955 := GTl(__e, tmp13954)

tmp13956 := GTl(__e, tmp13955)

tmp13957 := GEq(__e, Nil, tmp13956)

var ifres13938 Obj

if True == tmp13957 {
tmp13952 := GTl(__e, V4982)

tmp13953 := GConsP(__e, tmp13952)

var ifres13939 Obj

if True == tmp13953 {
tmp13949 := GTl(__e, V4982)

tmp13950 := GHd(__e, tmp13949)

tmp13951 := GEq(__e, sym_a_a_6, tmp13950)

var ifres13940 Obj

if True == tmp13951 {
tmp13946 := GTl(__e, V4982)

tmp13947 := GTl(__e, tmp13946)

tmp13948 := GConsP(__e, tmp13947)

var ifres13941 Obj

if True == tmp13948 {
tmp13942 := GTl(__e, V4982)

tmp13943 := GTl(__e, tmp13942)

tmp13944 := GTl(__e, tmp13943)

tmp13945 := GEq(__e, Nil, tmp13944)

ifres13941 = tmp13945


} else {
ifres13941 = False


}

ifres13940 = ifres13941


} else {
ifres13940 = False


}

ifres13939 = ifres13940


} else {
ifres13939 = False


}

ifres13938 = ifres13939


} else {
ifres13938 = False


}

ifres13937 = ifres13938


} else {
ifres13937 = False


}

ifres13936 = ifres13937


} else {
ifres13936 = False


}

ifres13935 = ifres13936


} else {
ifres13935 = False


}

ifres13934 = ifres13935


} else {
ifres13934 = False


}

if True == ifres13934 {
tmp13891 := GHd(__e, V4982)

tmp13892 := GHd(__e, V4982)

tmp13893 := GTl(__e, V4982)

tmp13894 := GTl(__e, tmp13893)

tmp13895 := GCons(__e, tmp13892, tmp13894)

tmp13896 := GCons(__e, symstr, tmp13895)

tmp13897 := GCons(__e, tmp13896, Nil)

tmp13898 := GCons(__e, sym_1_1_6, tmp13897)

tmp13899 := GCons(__e, tmp13891, tmp13898)

if SymFn(symshen_4curry_1type) == __self {
__self1 = tmp13899
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4curry_1type), tmp13899)


} else {
tmp13932 := GConsP(__e, V4982)

var ifres13910 Obj

if True == tmp13932 {
tmp13930 := GTl(__e, V4982)

tmp13931 := GConsP(__e, tmp13930)

var ifres13911 Obj

if True == tmp13931 {
tmp13927 := GTl(__e, V4982)

tmp13928 := GHd(__e, tmp13927)

tmp13929 := GEq(__e, sym_d, tmp13928)

var ifres13912 Obj

if True == tmp13929 {
tmp13924 := GTl(__e, V4982)

tmp13925 := GTl(__e, tmp13924)

tmp13926 := GConsP(__e, tmp13925)

var ifres13913 Obj

if True == tmp13926 {
tmp13920 := GTl(__e, V4982)

tmp13921 := GTl(__e, tmp13920)

tmp13922 := GTl(__e, tmp13921)

tmp13923 := GConsP(__e, tmp13922)

var ifres13914 Obj

if True == tmp13923 {
tmp13915 := GTl(__e, V4982)

tmp13916 := GTl(__e, tmp13915)

tmp13917 := GTl(__e, tmp13916)

tmp13918 := GHd(__e, tmp13917)

tmp13919 := GEq(__e, sym_d, tmp13918)

ifres13914 = tmp13919


} else {
ifres13914 = False


}

ifres13913 = ifres13914


} else {
ifres13913 = False


}

ifres13912 = ifres13913


} else {
ifres13912 = False


}

ifres13911 = ifres13912


} else {
ifres13911 = False


}

ifres13910 = ifres13911


} else {
ifres13910 = False


}

if True == ifres13910 {
tmp13900 := GHd(__e, V4982)

tmp13901 := GTl(__e, V4982)

tmp13902 := GTl(__e, tmp13901)

tmp13903 := GCons(__e, tmp13902, Nil)

tmp13904 := GCons(__e, sym_d, tmp13903)

tmp13905 := GCons(__e, tmp13900, tmp13904)

if SymFn(symshen_4curry_1type) == __self {
__self1 = tmp13905
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4curry_1type), tmp13905)


} else {
tmp13908 := GConsP(__e, V4982)

if True == tmp13908 {
tmp13906 := MakeNative1(func(__e *ControlFlow, Z4983 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4curry_1type), Z4983)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp13906, V4982)


} else {
return V4982
}


}


}


}


panic("unreachable")
})
return __self
})()

tmp13991 := Call2(__e, ns2_1set, symshen_4curry_1type, tmp13884)


_ = tmp13991

tmp13992 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V4984 := __self1
_ = V4984
tmp14069 := GConsP(__e, V4984)

var ifres14063 Obj

if True == tmp14069 {
tmp14067 := GHd(__e, V4984)

tmp14068 := GEq(__e, symdefine, tmp14067)

var ifres14064 Obj

if True == tmp14068 {
tmp14065 := GTl(__e, V4984)

tmp14066 := GConsP(__e, tmp14065)

ifres14064 = tmp14066


} else {
ifres14064 = False


}

ifres14063 = ifres14064


} else {
ifres14063 = False


}

if True == ifres14063 {
return V4984
} else {
tmp14061 := GConsP(__e, V4984)

var ifres14046 Obj

if True == tmp14061 {
tmp14059 := GHd(__e, V4984)

tmp14060 := GEq(__e, symtype, tmp14059)

var ifres14047 Obj

if True == tmp14060 {
tmp14057 := GTl(__e, V4984)

tmp14058 := GConsP(__e, tmp14057)

var ifres14048 Obj

if True == tmp14058 {
tmp14054 := GTl(__e, V4984)

tmp14055 := GTl(__e, tmp14054)

tmp14056 := GConsP(__e, tmp14055)

var ifres14049 Obj

if True == tmp14056 {
tmp14050 := GTl(__e, V4984)

tmp14051 := GTl(__e, tmp14050)

tmp14052 := GTl(__e, tmp14051)

tmp14053 := GEq(__e, Nil, tmp14052)

ifres14049 = tmp14053


} else {
ifres14049 = False


}

ifres14048 = ifres14049


} else {
ifres14048 = False


}

ifres14047 = ifres14048


} else {
ifres14047 = False


}

ifres14046 = ifres14047


} else {
ifres14046 = False


}

if True == ifres14046 {
tmp13993 := GTl(__e, V4984)

tmp13994 := GHd(__e, tmp13993)

tmp13995 := Call1(__e, SymFn(symshen_4curry), tmp13994)


tmp13996 := GTl(__e, V4984)

tmp13997 := GTl(__e, tmp13996)

tmp13998 := GCons(__e, tmp13995, tmp13997)

return GCons(__e, symtype, tmp13998)


} else {
tmp14044 := GConsP(__e, V4984)

var ifres14041 Obj

if True == tmp14044 {
tmp14042 := GHd(__e, V4984)

tmp14043 := Call1(__e, SymFn(symshen_4special_2), tmp14042)


ifres14041 = tmp14043


} else {
ifres14041 = False


}

if True == ifres14041 {
tmp13999 := GHd(__e, V4984)

tmp14000 := MakeNative1(func(__e *ControlFlow, Z4985 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4curry), Z4985)
panic("unreachable")
})

tmp14001 := GTl(__e, V4984)

tmp14002 := Call2(__e, SymFn(symmap), tmp14000, tmp14001)


return GCons(__e, tmp13999, tmp14002)


} else {
tmp14039 := GConsP(__e, V4984)

var ifres14036 Obj

if True == tmp14039 {
tmp14037 := GHd(__e, V4984)

tmp14038 := Call1(__e, SymFn(symshen_4extraspecial_2), tmp14037)


ifres14036 = tmp14038


} else {
ifres14036 = False


}

if True == ifres14036 {
return V4984
} else {
tmp14034 := GConsP(__e, V4984)

var ifres14027 Obj

if True == tmp14034 {
tmp14032 := GTl(__e, V4984)

tmp14033 := GConsP(__e, tmp14032)

var ifres14028 Obj

if True == tmp14033 {
tmp14029 := GTl(__e, V4984)

tmp14030 := GTl(__e, tmp14029)

tmp14031 := GConsP(__e, tmp14030)

ifres14028 = tmp14031


} else {
ifres14028 = False


}

ifres14027 = ifres14028


} else {
ifres14027 = False


}

if True == ifres14027 {
tmp14003 := GHd(__e, V4984)

tmp14004 := GTl(__e, V4984)

tmp14005 := GHd(__e, tmp14004)

tmp14006 := GCons(__e, tmp14005, Nil)

tmp14007 := GCons(__e, tmp14003, tmp14006)

tmp14008 := GTl(__e, V4984)

tmp14009 := GTl(__e, tmp14008)

tmp14010 := GCons(__e, tmp14007, tmp14009)

if SymFn(symshen_4curry) == __self {
__self1 = tmp14010
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4curry), tmp14010)


} else {
tmp14025 := GConsP(__e, V4984)

var ifres14018 Obj

if True == tmp14025 {
tmp14023 := GTl(__e, V4984)

tmp14024 := GConsP(__e, tmp14023)

var ifres14019 Obj

if True == tmp14024 {
tmp14020 := GTl(__e, V4984)

tmp14021 := GTl(__e, tmp14020)

tmp14022 := GEq(__e, Nil, tmp14021)

ifres14019 = tmp14022


} else {
ifres14019 = False


}

ifres14018 = ifres14019


} else {
ifres14018 = False


}

if True == ifres14018 {
tmp14011 := GHd(__e, V4984)

tmp14012 := Call1(__e, SymFn(symshen_4curry), tmp14011)


tmp14013 := GTl(__e, V4984)

tmp14014 := GHd(__e, tmp14013)

tmp14015 := Call1(__e, SymFn(symshen_4curry), tmp14014)


tmp14016 := GCons(__e, tmp14015, Nil)

return GCons(__e, tmp14012, tmp14016)


} else {
return V4984
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

tmp14070 := Call2(__e, ns2_1set, symshen_4curry, tmp13992)


_ = tmp14070

tmp14071 := MakeNative1(func(__e *ControlFlow, V4986 Obj) Obj {
tmp14072 := GValue(__e, symshen_4_dspecial_d)

return TailCall2(__e, SymFn(symelement_2), V4986, tmp14072)


panic("unreachable")
})

tmp14073 := Call2(__e, ns2_1set, symshen_4special_2, tmp14071)


_ = tmp14073

tmp14074 := MakeNative1(func(__e *ControlFlow, V4987 Obj) Obj {
tmp14075 := GValue(__e, symshen_4_dextraspecial_d)

return TailCall2(__e, SymFn(symelement_2), V4987, tmp14075)


panic("unreachable")
})

tmp14076 := Call2(__e, ns2_1set, symshen_4extraspecial_2, tmp14074)


_ = tmp14076

tmp14077 := MakeNative6(func(__e *ControlFlow, V4988, V4989, V4990, V4991, V4992, V4993 Obj) Obj {
let__13234 := GAdd(__e, V4992, MakeInteger(1))
_ = let__13234

tmp14083 := Call1(__e, SymFn(symshen_4unlocked_2), V4991)


var ifres14079 Obj

if True == tmp14083 {
tmp14080 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14080

tmp14081 := Call0(__e, SymFn(symshen_4maxinfexceeded_2))


tmp14082 := Call5(__e, SymFn(symwhen), tmp14081, V4990, V4991, let__13234, V4993)


ifres14079 = tmp14082


} else {
ifres14079 = False


}

let__13235 := ifres14079
_ = let__13235

tmp14131 := GEq(__e, let__13235, False)

if True == tmp14131 {
tmp14112 := Call1(__e, SymFn(symshen_4unlocked_2), V4991)


var ifres14084 Obj

if True == tmp14112 {
tmp14085 := Call2(__e, SymFn(symshen_4lazyderef), V4988, V4990)


let__13237 := tmp14085
_ = let__13237

tmp14111 := GConsP(__e, let__13237)

var ifres14086 Obj

if True == tmp14111 {
tmp14087 := GHd(__e, let__13237)

let__13238 := tmp14087
_ = let__13238

tmp14088 := GTl(__e, let__13237)

tmp14089 := Call2(__e, SymFn(symshen_4lazyderef), tmp14088, V4990)


let__13239 := tmp14089
_ = let__13239

tmp14110 := GConsP(__e, let__13239)

var ifres14090 Obj

if True == tmp14110 {
tmp14091 := GHd(__e, let__13239)

let__13240 := tmp14091
_ = let__13240

tmp14092 := GTl(__e, let__13239)

tmp14093 := Call2(__e, SymFn(symshen_4lazyderef), tmp14092, V4990)


let__13241 := tmp14093
_ = let__13241

tmp14109 := GConsP(__e, let__13241)

var ifres14094 Obj

if True == tmp14109 {
tmp14095 := GHd(__e, let__13241)

let__13242 := tmp14095
_ = let__13242

tmp14096 := GTl(__e, let__13241)

tmp14097 := Call2(__e, SymFn(symshen_4lazyderef), tmp14096, V4990)


let__13243 := tmp14097
_ = let__13243

tmp14108 := GEq(__e, let__13243, Nil)

var ifres14098 Obj

if True == tmp14108 {
tmp14099 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14099

tmp14100 := Call2(__e, SymFn(symshen_4deref), let__13240, V4990)


tmp14101 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp14102 := GEq(__e, tmp14100, tmp14101)

tmp14103 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14104 := Call0(__e, SymFn(symshen_4type_1theory_1enabled_2))


tmp14105 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14106 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), let__13238, let__13242, V4989, V4990, V4991, let__13234, V4993)
return nil
panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V4990, V4991, let__13234, tmp14106)


panic("unreachable")
})

return TailCall5(__e, SymFn(symwhen), tmp14104, V4990, V4991, let__13234, tmp14105)


panic("unreachable")
})

tmp14107 := Call5(__e, SymFn(symwhen), tmp14102, V4990, V4991, let__13234, tmp14103)


ifres14098 = tmp14107


} else {
ifres14098 = False


}

ifres14094 = ifres14098


} else {
ifres14094 = False


}

ifres14090 = ifres14094


} else {
ifres14090 = False


}

ifres14086 = ifres14090


} else {
ifres14086 = False


}

ifres14084 = ifres14086


} else {
ifres14084 = False


}

let__13236 := ifres14084
_ = let__13236

tmp14129 := GEq(__e, let__13236, False)

if True == tmp14129 {
tmp14118 := Call1(__e, SymFn(symshen_4unlocked_2), V4991)


var ifres14113 Obj

if True == tmp14118 {
tmp14114 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14114

tmp14115 := GValue(__e, symshen_4_dspy_d)

tmp14116 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall6(__e, SymFn(symshen_4show), V4988, V4989, V4990, V4991, let__13234, V4993)
panic("unreachable")
})

tmp14117 := Call5(__e, SymFn(symwhen), tmp14115, V4990, V4991, let__13234, tmp14116)


ifres14113 = tmp14117


} else {
ifres14113 = False


}

let__13244 := ifres14113
_ = let__13244

tmp14127 := GEq(__e, let__13244, False)

if True == tmp14127 {
tmp14123 := Call1(__e, SymFn(symshen_4unlocked_2), V4991)


var ifres14119 Obj

if True == tmp14123 {
tmp14120 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14120

tmp14121 := GValue(__e, symshen_4_ddatatypes_d)

tmp14122 := Call(__e, SymFn(symshen_4search_1user_1datatypes), V4988, V4989, tmp14121, V4990, V4991, let__13234, V4993)


ifres14119 = tmp14122


} else {
ifres14119 = False


}

let__13245 := ifres14119
_ = let__13245

tmp14125 := GEq(__e, let__13245, False)

if True == tmp14125 {
return TailCall2(__e, SymFn(symshen_4unlock), V4991, let__13234)
} else {
return let__13245
}


} else {
return let__13244
}


} else {
return let__13236
}


} else {
return let__13235
}


panic("unreachable")
})

tmp14132 := Call2(__e, ns2_1set, symshen_4system_1S, tmp14077)


_ = tmp14132

tmp14133 := MakeNative6(func(__e *ControlFlow, V5012, V5013, V5014, V5015, V5016, V5017 Obj) Obj {
tmp14134 := Call0(__e, SymFn(symshen_4line))


_ = tmp14134

tmp14135 := Call2(__e, SymFn(symshen_4deref), V5012, V5014)


tmp14136 := Call1(__e, SymFn(symshen_4show_1p), tmp14135)


_ = tmp14136

tmp14137 := Call1(__e, SymFn(symnl), MakeInteger(2))


_ = tmp14137

tmp14138 := Call2(__e, SymFn(symshen_4deref), V5013, V5014)


tmp14139 := Call2(__e, SymFn(symshen_4show_1assumptions), tmp14138, MakeInteger(1))


_ = tmp14139

tmp14140 := Call0(__e, SymFn(symshen_4pause_1for_1user))


_ = tmp14140

return False


panic("unreachable")
})

tmp14141 := Call2(__e, ns2_1set, symshen_4show, tmp14133)


_ = tmp14141

tmp14142 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14143 := Call0(__e, SymFn(syminferences))


let__13246 := tmp14143
_ = let__13246

tmp14145 := GEq(__e, MakeInteger(1), let__13246)

var ifres14144 Obj

if True == tmp14145 {
ifres14144 = MakeString("")


} else {
ifres14144 = MakeString("s")


}

tmp14146 := Call3(__e, SymFn(symshen_4app), ifres14144, MakeString(" \n?- "), symshen_4a)


tmp14148 := Call3(__e, SymFn(symshen_4app), let__13246, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" inference"), tmp14146), symshen_4a)


tmp14149 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("____________________________________________________________ "), tmp14148)

tmp14150 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp14149, tmp14150)


panic("unreachable")
})

tmp14151 := Call2(__e, ns2_1set, symshen_4line, tmp14142)


_ = tmp14151

tmp14152 := MakeNative1(func(__e *ControlFlow, V5019 Obj) Obj {
tmp14180 := GConsP(__e, V5019)

var ifres14163 Obj

if True == tmp14180 {
tmp14178 := GTl(__e, V5019)

tmp14179 := GConsP(__e, tmp14178)

var ifres14164 Obj

if True == tmp14179 {
tmp14175 := GTl(__e, V5019)

tmp14176 := GTl(__e, tmp14175)

tmp14177 := GConsP(__e, tmp14176)

var ifres14165 Obj

if True == tmp14177 {
tmp14171 := GTl(__e, V5019)

tmp14172 := GTl(__e, tmp14171)

tmp14173 := GTl(__e, tmp14172)

tmp14174 := GEq(__e, Nil, tmp14173)

var ifres14166 Obj

if True == tmp14174 {
tmp14167 := GTl(__e, V5019)

tmp14168 := GHd(__e, tmp14167)

tmp14169 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp14170 := GEq(__e, tmp14168, tmp14169)

ifres14166 = tmp14170


} else {
ifres14166 = False


}

ifres14165 = ifres14166


} else {
ifres14165 = False


}

ifres14164 = ifres14165


} else {
ifres14164 = False


}

ifres14163 = ifres14164


} else {
ifres14163 = False


}

if True == ifres14163 {
tmp14153 := GHd(__e, V5019)

tmp14154 := Call1(__e, SymFn(symshen_4prterm), tmp14153)


_ = tmp14154

tmp14155 := Call0(__e, SymFn(symstoutput))


tmp14156 := Call2(__e, SymFn(sympr), MakeString(" : "), tmp14155)


_ = tmp14156

tmp14157 := GTl(__e, V5019)

tmp14158 := GTl(__e, tmp14157)

tmp14159 := GHd(__e, tmp14158)

tmp14160 := Call3(__e, SymFn(symshen_4app), tmp14159, MakeString(""), symshen_4r)


tmp14161 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp14160, tmp14161)


} else {
return TailCall1(__e, SymFn(symshen_4prterm), V5019)
}


panic("unreachable")
})

tmp14181 := Call2(__e, ns2_1set, symshen_4show_1p, tmp14152)


_ = tmp14181

tmp14182 := MakeNative1(func(__e *ControlFlow, V5020 Obj) Obj {
tmp14221 := GConsP(__e, V5020)

var ifres14206 Obj

if True == tmp14221 {
tmp14219 := GHd(__e, V5020)

tmp14220 := GEq(__e, symcons, tmp14219)

var ifres14207 Obj

if True == tmp14220 {
tmp14217 := GTl(__e, V5020)

tmp14218 := GConsP(__e, tmp14217)

var ifres14208 Obj

if True == tmp14218 {
tmp14214 := GTl(__e, V5020)

tmp14215 := GTl(__e, tmp14214)

tmp14216 := GConsP(__e, tmp14215)

var ifres14209 Obj

if True == tmp14216 {
tmp14210 := GTl(__e, V5020)

tmp14211 := GTl(__e, tmp14210)

tmp14212 := GTl(__e, tmp14211)

tmp14213 := GEq(__e, Nil, tmp14212)

ifres14209 = tmp14213


} else {
ifres14209 = False


}

ifres14208 = ifres14209


} else {
ifres14208 = False


}

ifres14207 = ifres14208


} else {
ifres14207 = False


}

ifres14206 = ifres14207


} else {
ifres14206 = False


}

if True == ifres14206 {
tmp14183 := Call0(__e, SymFn(symstoutput))


tmp14184 := Call2(__e, SymFn(sympr), MakeString("["), tmp14183)


_ = tmp14184

tmp14185 := GTl(__e, V5020)

tmp14186 := GHd(__e, tmp14185)

tmp14187 := Call1(__e, SymFn(symshen_4prterm), tmp14186)


_ = tmp14187

tmp14188 := GTl(__e, V5020)

tmp14189 := GTl(__e, tmp14188)

tmp14190 := GHd(__e, tmp14189)

tmp14191 := Call1(__e, SymFn(symshen_4prtl), tmp14190)


_ = tmp14191

tmp14192 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), MakeString("]"), tmp14192)


} else {
tmp14204 := GConsP(__e, V5020)

if True == tmp14204 {
tmp14193 := Call0(__e, SymFn(symstoutput))


tmp14194 := Call2(__e, SymFn(sympr), MakeString("("), tmp14193)


_ = tmp14194

tmp14195 := GHd(__e, V5020)

tmp14196 := Call1(__e, SymFn(symshen_4prterm), tmp14195)


_ = tmp14196

tmp14197 := MakeNative1(func(__e *ControlFlow, Z5021 Obj) Obj {
tmp14198 := Call0(__e, SymFn(symstoutput))


tmp14199 := Call2(__e, SymFn(sympr), MakeString(" "), tmp14198)


_ = tmp14199

return TailCall1(__e, SymFn(symshen_4prterm), Z5021)


panic("unreachable")
})

tmp14200 := GTl(__e, V5020)

tmp14201 := Call2(__e, SymFn(symmap), tmp14197, tmp14200)


_ = tmp14201

tmp14202 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), MakeString(")"), tmp14202)


} else {
return TailCall1(__e, SymFn(symprint), V5020)
}


}


panic("unreachable")
})

tmp14222 := Call2(__e, ns2_1set, symshen_4prterm, tmp14182)


_ = tmp14222

tmp14223 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V5022 := __self1
_ = V5022
tmp14252 := GEq(__e, Nil, V5022)

if True == tmp14252 {
return MakeString("")
} else {
tmp14250 := GConsP(__e, V5022)

var ifres14235 Obj

if True == tmp14250 {
tmp14248 := GHd(__e, V5022)

tmp14249 := GEq(__e, symcons, tmp14248)

var ifres14236 Obj

if True == tmp14249 {
tmp14246 := GTl(__e, V5022)

tmp14247 := GConsP(__e, tmp14246)

var ifres14237 Obj

if True == tmp14247 {
tmp14243 := GTl(__e, V5022)

tmp14244 := GTl(__e, tmp14243)

tmp14245 := GConsP(__e, tmp14244)

var ifres14238 Obj

if True == tmp14245 {
tmp14239 := GTl(__e, V5022)

tmp14240 := GTl(__e, tmp14239)

tmp14241 := GTl(__e, tmp14240)

tmp14242 := GEq(__e, Nil, tmp14241)

ifres14238 = tmp14242


} else {
ifres14238 = False


}

ifres14237 = ifres14238


} else {
ifres14237 = False


}

ifres14236 = ifres14237


} else {
ifres14236 = False


}

ifres14235 = ifres14236


} else {
ifres14235 = False


}

if True == ifres14235 {
tmp14224 := Call0(__e, SymFn(symstoutput))


tmp14225 := Call2(__e, SymFn(sympr), MakeString(" "), tmp14224)


_ = tmp14225

tmp14226 := GTl(__e, V5022)

tmp14227 := GHd(__e, tmp14226)

tmp14228 := Call1(__e, SymFn(symshen_4prterm), tmp14227)


_ = tmp14228

tmp14229 := GTl(__e, V5022)

tmp14230 := GTl(__e, tmp14229)

tmp14231 := GHd(__e, tmp14230)

if SymFn(symshen_4prtl) == __self {
__self1 = tmp14231
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4prtl), tmp14231)


} else {
tmp14232 := Call0(__e, SymFn(symstoutput))


tmp14233 := Call2(__e, SymFn(sympr), MakeString(" | "), tmp14232)


_ = tmp14233

return TailCall1(__e, SymFn(symshen_4prterm), V5022)


}


}


panic("unreachable")
})
return __self
})()

tmp14253 := Call2(__e, ns2_1set, symshen_4prtl, tmp14223)


_ = tmp14253

tmp14254 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V5029 := __self1
_ = V5029
V5030 := __self2
_ = V5030
tmp14267 := GEq(__e, Nil, V5029)

if True == tmp14267 {
tmp14255 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), MakeString("\n> "), tmp14255)


} else {
tmp14265 := GConsP(__e, V5029)

if True == tmp14265 {
tmp14256 := Call3(__e, SymFn(symshen_4app), V5030, MakeString(". "), symshen_4a)


tmp14257 := Call0(__e, SymFn(symstoutput))


tmp14258 := Call2(__e, SymFn(sympr), tmp14256, tmp14257)


_ = tmp14258

tmp14259 := GHd(__e, V5029)

tmp14260 := Call1(__e, SymFn(symshen_4show_1p), tmp14259)


_ = tmp14260

tmp14261 := Call1(__e, SymFn(symnl), MakeInteger(1))


_ = tmp14261

tmp14262 := GTl(__e, V5029)

if SymFn(symshen_4show_1assumptions) == __self {
__self1, __self2 = tmp14262, GAdd(__e, V5030, MakeInteger(1))
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4show_1assumptions), tmp14262, GAdd(__e, V5030, MakeInteger(1)))


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.show-assumptions"))
}


}


panic("unreachable")
})
return __self
})()

tmp14268 := Call2(__e, ns2_1set, symshen_4show_1assumptions, tmp14254)


_ = tmp14268

tmp14269 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14270 := Call0(__e, SymFn(symstinput))


tmp14271 := PrimGuard1(__e, symread_1byte, PrimReadByte, tmp14270)

let__13247 := tmp14271
_ = let__13247

tmp14273 := GEq(__e, let__13247, MakeInteger(94))

if True == tmp14273 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("input aborted\n"))
} else {
return TailCall1(__e, SymFn(symnl), MakeInteger(1))
}


panic("unreachable")
})

tmp14274 := Call2(__e, ns2_1set, symshen_4pause_1for_1user, tmp14269)


_ = tmp14274

tmp14275 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dshen_1type_1theory_1enabled_2_d)
panic("unreachable")
})

tmp14276 := Call2(__e, ns2_1set, symshen_4type_1theory_1enabled_2, tmp14275)


_ = tmp14276

tmp14277 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14279 := Call0(__e, SymFn(syminferences))


tmp14280 := GValue(__e, symshen_4_dmaxinferences_d)

if True == GGt(__e, tmp14279, tmp14280) {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("maximum inferences exceeded"))
} else {
return False
}


panic("unreachable")
})

tmp14282 := Call2(__e, ns2_1set, symshen_4maxinfexceeded_2, tmp14277)


_ = tmp14282

tmp14283 := MakeNative(func(__e *ControlFlow) {
V5032 := __e.Get(1)
V5033 := __e.Get(2)
V5034 := __e.Get(3)
V5035 := __e.Get(4)
V5036 := __e.Get(5)
V5037 := __e.Get(6)
V5038 := __e.Get(7)
__e.ReturnDirect(func() Obj {
let__13248 := GAdd(__e, V5037, MakeInteger(1))
_ = let__13248

tmp14294 := Call1(__e, SymFn(symshen_4unlocked_2), V5036)


var ifres14285 Obj

if True == tmp14294 {
tmp14286 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14286

tmp14287 := GValue(__e, symshen_4_dspy_d)

tmp14288 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14289 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp14290 := GCons(__e, V5033, Nil)

tmp14291 := GCons(__e, tmp14289, tmp14290)

tmp14292 := GCons(__e, V5032, tmp14291)

return TailCall6(__e, SymFn(symshen_4show), tmp14292, V5034, V5035, V5036, let__13248, V5038)


panic("unreachable")
})

tmp14293 := Call5(__e, SymFn(symwhen), tmp14287, V5035, V5036, let__13248, tmp14288)


ifres14285 = tmp14293


} else {
ifres14285 = False


}

let__13249 := ifres14285
_ = let__13249

tmp15033 := GEq(__e, let__13249, False)

if True == tmp15033 {
tmp14302 := Call1(__e, SymFn(symshen_4unlocked_2), V5036)


var ifres14295 Obj

if True == tmp14302 {
tmp14296 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14296

tmp14297 := Call2(__e, SymFn(symshen_4lazyderef), V5032, V5035)


tmp14298 := GConsP(__e, tmp14297)

tmp14299 := GNot(__e, tmp14298)

tmp14300 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall6(__e, SymFn(symshen_4primitive), V5032, V5033, V5035, V5036, let__13248, V5038)
panic("unreachable")
})

tmp14301 := Call5(__e, SymFn(symwhen), tmp14299, V5035, V5036, let__13248, tmp14300)


ifres14295 = tmp14301


} else {
ifres14295 = False


}

let__13250 := ifres14295
_ = let__13250

tmp15031 := GEq(__e, let__13250, False)

if True == tmp15031 {
tmp14306 := Call1(__e, SymFn(symshen_4unlocked_2), V5036)


var ifres14303 Obj

if True == tmp14306 {
tmp14304 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14304

tmp14305 := Call(__e, SymFn(symshen_4by_1hypothesis), V5032, V5033, V5034, V5035, V5036, let__13248, V5038)


ifres14303 = tmp14305


} else {
ifres14303 = False


}

let__13251 := ifres14303
_ = let__13251

tmp15029 := GEq(__e, let__13251, False)

if True == tmp15029 {
tmp14320 := Call1(__e, SymFn(symshen_4unlocked_2), V5036)


var ifres14307 Obj

if True == tmp14320 {
tmp14308 := Call2(__e, SymFn(symshen_4lazyderef), V5032, V5035)


let__13253 := tmp14308
_ = let__13253

tmp14319 := GConsP(__e, let__13253)

var ifres14309 Obj

if True == tmp14319 {
tmp14310 := GHd(__e, let__13253)

let__13254 := tmp14310
_ = let__13254

tmp14311 := GTl(__e, let__13253)

tmp14312 := Call2(__e, SymFn(symshen_4lazyderef), tmp14311, V5035)


let__13255 := tmp14312
_ = let__13255

tmp14318 := GEq(__e, let__13255, Nil)

var ifres14313 Obj

if True == tmp14318 {
tmp14314 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14314

tmp14315 := GCons(__e, V5033, Nil)

tmp14316 := GCons(__e, sym_1_1_6, tmp14315)

tmp14317 := Call6(__e, SymFn(symshen_4lookupsig), let__13254, tmp14316, V5035, V5036, let__13248, V5038)


ifres14313 = tmp14317


} else {
ifres14313 = False


}

ifres14309 = ifres14313


} else {
ifres14309 = False


}

ifres14307 = ifres14309


} else {
ifres14307 = False


}

let__13252 := ifres14307
_ = let__13252

tmp15027 := GEq(__e, let__13252, False)

if True == tmp15027 {
tmp14346 := Call1(__e, SymFn(symshen_4unlocked_2), V5036)


var ifres14321 Obj

if True == tmp14346 {
tmp14322 := Call2(__e, SymFn(symshen_4lazyderef), V5032, V5035)


let__13257 := tmp14322
_ = let__13257

tmp14345 := GConsP(__e, let__13257)

var ifres14323 Obj

if True == tmp14345 {
tmp14324 := GHd(__e, let__13257)

tmp14325 := Call2(__e, SymFn(symshen_4lazyderef), tmp14324, V5035)


let__13258 := tmp14325
_ = let__13258

tmp14344 := GEq(__e, let__13258, symfn)

var ifres14326 Obj

if True == tmp14344 {
tmp14327 := GTl(__e, let__13257)

tmp14328 := Call2(__e, SymFn(symshen_4lazyderef), tmp14327, V5035)


let__13259 := tmp14328
_ = let__13259

tmp14343 := GConsP(__e, let__13259)

var ifres14329 Obj

if True == tmp14343 {
tmp14330 := GHd(__e, let__13259)

let__13260 := tmp14330
_ = let__13260

tmp14331 := GTl(__e, let__13259)

tmp14332 := Call2(__e, SymFn(symshen_4lazyderef), tmp14331, V5035)


let__13261 := tmp14332
_ = let__13261

tmp14342 := GEq(__e, let__13261, Nil)

var ifres14333 Obj

if True == tmp14342 {
tmp14334 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14334

tmp14335 := Call2(__e, SymFn(symshen_4deref), let__13260, V5035)


tmp14336 := Call1(__e, SymFn(symarity), tmp14335)


tmp14337 := GEq(__e, tmp14336, MakeInteger(0))

tmp14338 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14339 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14340 := GCons(__e, let__13260, Nil)

__e.TailApply(SymFn(symshen_4system_1S_1h), tmp14340, V5033, V5034, V5035, V5036, let__13248, V5038)
return nil


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5035, V5036, let__13248, tmp14339)


panic("unreachable")
})

tmp14341 := Call5(__e, SymFn(symwhen), tmp14337, V5035, V5036, let__13248, tmp14338)


ifres14333 = tmp14341


} else {
ifres14333 = False


}

ifres14329 = ifres14333


} else {
ifres14329 = False


}

ifres14326 = ifres14329


} else {
ifres14326 = False


}

ifres14323 = ifres14326


} else {
ifres14323 = False


}

ifres14321 = ifres14323


} else {
ifres14321 = False


}

let__13256 := ifres14321
_ = let__13256

tmp15025 := GEq(__e, let__13256, False)

if True == tmp15025 {
tmp14366 := Call1(__e, SymFn(symshen_4unlocked_2), V5036)


var ifres14347 Obj

if True == tmp14366 {
tmp14348 := Call2(__e, SymFn(symshen_4lazyderef), V5032, V5035)


let__13263 := tmp14348
_ = let__13263

tmp14365 := GConsP(__e, let__13263)

var ifres14349 Obj

if True == tmp14365 {
tmp14350 := GHd(__e, let__13263)

tmp14351 := Call2(__e, SymFn(symshen_4lazyderef), tmp14350, V5035)


let__13264 := tmp14351
_ = let__13264

tmp14364 := GEq(__e, let__13264, symfn)

var ifres14352 Obj

if True == tmp14364 {
tmp14353 := GTl(__e, let__13263)

tmp14354 := Call2(__e, SymFn(symshen_4lazyderef), tmp14353, V5035)


let__13265 := tmp14354
_ = let__13265

tmp14363 := GConsP(__e, let__13265)

var ifres14355 Obj

if True == tmp14363 {
tmp14356 := GHd(__e, let__13265)

let__13266 := tmp14356
_ = let__13266

tmp14357 := GTl(__e, let__13265)

tmp14358 := Call2(__e, SymFn(symshen_4lazyderef), tmp14357, V5035)


let__13267 := tmp14358
_ = let__13267

tmp14362 := GEq(__e, let__13267, Nil)

var ifres14359 Obj

if True == tmp14362 {
tmp14360 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14360

tmp14361 := Call6(__e, SymFn(symshen_4lookupsig), let__13266, V5033, V5035, V5036, let__13248, V5038)


ifres14359 = tmp14361


} else {
ifres14359 = False


}

ifres14355 = ifres14359


} else {
ifres14355 = False


}

ifres14352 = ifres14355


} else {
ifres14352 = False


}

ifres14349 = ifres14352


} else {
ifres14349 = False


}

ifres14347 = ifres14349


} else {
ifres14347 = False


}

let__13262 := ifres14347
_ = let__13262

tmp15023 := GEq(__e, let__13262, False)

if True == tmp15023 {
tmp14393 := Call1(__e, SymFn(symshen_4unlocked_2), V5036)


var ifres14367 Obj

if True == tmp14393 {
tmp14368 := Call2(__e, SymFn(symshen_4lazyderef), V5032, V5035)


let__13269 := tmp14368
_ = let__13269

tmp14392 := GConsP(__e, let__13269)

var ifres14369 Obj

if True == tmp14392 {
tmp14370 := GHd(__e, let__13269)

let__13270 := tmp14370
_ = let__13270

tmp14371 := GTl(__e, let__13269)

tmp14372 := Call2(__e, SymFn(symshen_4lazyderef), tmp14371, V5035)


let__13271 := tmp14372
_ = let__13271

tmp14391 := GConsP(__e, let__13271)

var ifres14373 Obj

if True == tmp14391 {
tmp14374 := GHd(__e, let__13271)

let__13272 := tmp14374
_ = let__13272

tmp14375 := GTl(__e, let__13271)

tmp14376 := Call2(__e, SymFn(symshen_4lazyderef), tmp14375, V5035)


let__13273 := tmp14376
_ = let__13273

tmp14390 := GEq(__e, let__13273, Nil)

var ifres14377 Obj

if True == tmp14390 {
tmp14378 := Call1(__e, SymFn(symshen_4newpv), V5035)


let__13274 := tmp14378
_ = let__13274

tmp14379 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14379

tmp14380 := Call2(__e, SymFn(symshen_4lazyderef), let__13270, V5035)


tmp14381 := GConsP(__e, tmp14380)

tmp14382 := GNot(__e, tmp14381)

tmp14383 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14384 := GCons(__e, V5033, Nil)

tmp14385 := GCons(__e, sym_1_1_6, tmp14384)

tmp14386 := GCons(__e, let__13274, tmp14385)

tmp14387 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), let__13272, let__13274, V5034, V5035, V5036, let__13248, V5038)
return nil
panic("unreachable")
})

return TailCall6(__e, SymFn(symshen_4lookupsig), let__13270, tmp14386, V5035, V5036, let__13248, tmp14387)


panic("unreachable")
})

tmp14388 := Call5(__e, SymFn(symwhen), tmp14382, V5035, V5036, let__13248, tmp14383)


tmp14389 := Call2(__e, SymFn(symshen_4gc), V5035, tmp14388)


ifres14377 = tmp14389


} else {
ifres14377 = False


}

ifres14373 = ifres14377


} else {
ifres14373 = False


}

ifres14369 = ifres14373


} else {
ifres14369 = False


}

ifres14367 = ifres14369


} else {
ifres14367 = False


}

let__13268 := ifres14367
_ = let__13268

tmp15021 := GEq(__e, let__13268, False)

if True == tmp15021 {
tmp14416 := Call1(__e, SymFn(symshen_4unlocked_2), V5036)


var ifres14394 Obj

if True == tmp14416 {
tmp14395 := Call2(__e, SymFn(symshen_4lazyderef), V5032, V5035)


let__13276 := tmp14395
_ = let__13276

tmp14415 := GConsP(__e, let__13276)

var ifres14396 Obj

if True == tmp14415 {
tmp14397 := GHd(__e, let__13276)

let__13277 := tmp14397
_ = let__13277

tmp14398 := GTl(__e, let__13276)

tmp14399 := Call2(__e, SymFn(symshen_4lazyderef), tmp14398, V5035)


let__13278 := tmp14399
_ = let__13278

tmp14414 := GConsP(__e, let__13278)

var ifres14400 Obj

if True == tmp14414 {
tmp14401 := GHd(__e, let__13278)

let__13279 := tmp14401
_ = let__13279

tmp14402 := GTl(__e, let__13278)

tmp14403 := Call2(__e, SymFn(symshen_4lazyderef), tmp14402, V5035)


let__13280 := tmp14403
_ = let__13280

tmp14413 := GEq(__e, let__13280, Nil)

var ifres14404 Obj

if True == tmp14413 {
tmp14405 := Call1(__e, SymFn(symshen_4newpv), V5035)


let__13281 := tmp14405
_ = let__13281

tmp14406 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14406

tmp14407 := GCons(__e, V5033, Nil)

tmp14408 := GCons(__e, sym_1_1_6, tmp14407)

tmp14409 := GCons(__e, let__13281, tmp14408)

tmp14410 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), let__13279, let__13281, V5034, V5035, V5036, let__13248, V5038)
return nil
panic("unreachable")
})

tmp14411 := Call(__e, SymFn(symshen_4system_1S_1h), let__13277, tmp14409, V5034, V5035, V5036, let__13248, tmp14410)


tmp14412 := Call2(__e, SymFn(symshen_4gc), V5035, tmp14411)


ifres14404 = tmp14412


} else {
ifres14404 = False


}

ifres14400 = ifres14404


} else {
ifres14400 = False


}

ifres14396 = ifres14400


} else {
ifres14396 = False


}

ifres14394 = ifres14396


} else {
ifres14394 = False


}

let__13275 := ifres14394
_ = let__13275

tmp15019 := GEq(__e, let__13275, False)

if True == tmp15019 {
tmp14483 := Call1(__e, SymFn(symshen_4unlocked_2), V5036)


var ifres14417 Obj

if True == tmp14483 {
tmp14418 := Call2(__e, SymFn(symshen_4lazyderef), V5032, V5035)


let__13283 := tmp14418
_ = let__13283

tmp14482 := GConsP(__e, let__13283)

var ifres14419 Obj

if True == tmp14482 {
tmp14420 := GHd(__e, let__13283)

tmp14421 := Call2(__e, SymFn(symshen_4lazyderef), tmp14420, V5035)


let__13284 := tmp14421
_ = let__13284

tmp14481 := GEq(__e, let__13284, symcons)

var ifres14422 Obj

if True == tmp14481 {
tmp14423 := GTl(__e, let__13283)

tmp14424 := Call2(__e, SymFn(symshen_4lazyderef), tmp14423, V5035)


let__13285 := tmp14424
_ = let__13285

tmp14480 := GConsP(__e, let__13285)

var ifres14425 Obj

if True == tmp14480 {
tmp14426 := GHd(__e, let__13285)

let__13286 := tmp14426
_ = let__13286

tmp14427 := GTl(__e, let__13285)

tmp14428 := Call2(__e, SymFn(symshen_4lazyderef), tmp14427, V5035)


let__13287 := tmp14428
_ = let__13287

tmp14479 := GConsP(__e, let__13287)

var ifres14429 Obj

if True == tmp14479 {
tmp14430 := GHd(__e, let__13287)

let__13288 := tmp14430
_ = let__13288

tmp14431 := GTl(__e, let__13287)

tmp14432 := Call2(__e, SymFn(symshen_4lazyderef), tmp14431, V5035)


let__13289 := tmp14432
_ = let__13289

tmp14478 := GEq(__e, let__13289, Nil)

var ifres14433 Obj

if True == tmp14478 {
tmp14434 := Call2(__e, SymFn(symshen_4lazyderef), V5033, V5035)


let__13290 := tmp14434
_ = let__13290

tmp14435 := MakeNative1(func(__e *ControlFlow, Z5083 Obj) Obj {
tmp14436 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14436

tmp14437 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14438 := GCons(__e, Z5083, Nil)

tmp14439 := GCons(__e, symlist, tmp14438)

__e.TailApply(SymFn(symshen_4system_1S_1h), let__13288, tmp14439, V5034, V5035, V5036, let__13248, V5038)
return nil


panic("unreachable")
})

__e.TailApply(SymFn(symshen_4system_1S_1h), let__13286, Z5083, V5034, V5035, V5036, let__13248, tmp14437)
return nil


panic("unreachable")
})

let__13291 := tmp14435
_ = let__13291

tmp14477 := GConsP(__e, let__13290)

var ifres14440 Obj

if True == tmp14477 {
tmp14441 := GHd(__e, let__13290)

tmp14442 := Call2(__e, SymFn(symshen_4lazyderef), tmp14441, V5035)


let__13292 := tmp14442
_ = let__13292

tmp14443 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14444 := GTl(__e, let__13290)

tmp14445 := Call2(__e, SymFn(symshen_4lazyderef), tmp14444, V5035)


let__13294 := tmp14445
_ = let__13294

tmp14446 := MakeNative1(func(__e *ControlFlow, Z5088 Obj) Obj {
return TailCall1(__e, let__13291, Z5088)
panic("unreachable")
})

let__13295 := tmp14446
_ = let__13295

tmp14462 := GConsP(__e, let__13294)

if True == tmp14462 {
tmp14447 := GHd(__e, let__13294)

let__13296 := tmp14447
_ = let__13296

tmp14448 := GTl(__e, let__13294)

tmp14449 := Call2(__e, SymFn(symshen_4lazyderef), tmp14448, V5035)


let__13297 := tmp14449
_ = let__13297

tmp14450 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13295, let__13296)
panic("unreachable")
})

let__13298 := tmp14450
_ = let__13298

tmp14454 := GEq(__e, let__13297, Nil)

if True == tmp14454 {
return TailCall1(__e, SymFn(symthaw), let__13298)
} else {
tmp14452 := Call1(__e, SymFn(symshen_4pvar_2), let__13297)


if True == tmp14452 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__13297, Nil, V5035, let__13298)
} else {
return False
}


}


} else {
tmp14460 := Call1(__e, SymFn(symshen_4pvar_2), let__13294)


if True == tmp14460 {
tmp14455 := Call1(__e, SymFn(symshen_4newpv), V5035)


let__13299 := tmp14455
_ = let__13299

tmp14456 := GCons(__e, let__13299, Nil)

tmp14457 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13295, let__13299)
panic("unreachable")
})

tmp14458 := Call4(__e, SymFn(symshen_4bind_b), let__13294, tmp14456, V5035, tmp14457)


return TailCall2(__e, SymFn(symshen_4gc), V5035, tmp14458)


} else {
return False
}


}


panic("unreachable")
})

let__13293 := tmp14443
_ = let__13293

tmp14468 := GEq(__e, let__13292, symlist)

var ifres14463 Obj

if True == tmp14468 {
tmp14464 := Call1(__e, SymFn(symthaw), let__13293)


ifres14463 = tmp14464


} else {
tmp14467 := Call1(__e, SymFn(symshen_4pvar_2), let__13292)


var ifres14465 Obj

if True == tmp14467 {
tmp14466 := Call4(__e, SymFn(symshen_4bind_b), let__13292, symlist, V5035, let__13293)


ifres14465 = tmp14466


} else {
ifres14465 = False


}

ifres14463 = ifres14465


}

ifres14440 = ifres14463


} else {
tmp14476 := Call1(__e, SymFn(symshen_4pvar_2), let__13290)


var ifres14469 Obj

if True == tmp14476 {
tmp14470 := Call1(__e, SymFn(symshen_4newpv), V5035)


let__13300 := tmp14470
_ = let__13300

tmp14471 := GCons(__e, let__13300, Nil)

tmp14472 := GCons(__e, symlist, tmp14471)

tmp14473 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13291, let__13300)
panic("unreachable")
})

tmp14474 := Call4(__e, SymFn(symshen_4bind_b), let__13290, tmp14472, V5035, tmp14473)


tmp14475 := Call2(__e, SymFn(symshen_4gc), V5035, tmp14474)


ifres14469 = tmp14475


} else {
ifres14469 = False


}

ifres14440 = ifres14469


}

ifres14433 = ifres14440


} else {
ifres14433 = False


}

ifres14429 = ifres14433


} else {
ifres14429 = False


}

ifres14425 = ifres14429


} else {
ifres14425 = False


}

ifres14422 = ifres14425


} else {
ifres14422 = False


}

ifres14419 = ifres14422


} else {
ifres14419 = False


}

ifres14417 = ifres14419


} else {
ifres14417 = False


}

let__13282 := ifres14417
_ = let__13282

tmp15017 := GEq(__e, let__13282, False)

if True == tmp15017 {
tmp14567 := Call1(__e, SymFn(symshen_4unlocked_2), V5036)


var ifres14484 Obj

if True == tmp14567 {
tmp14485 := Call2(__e, SymFn(symshen_4lazyderef), V5032, V5035)


let__13302 := tmp14485
_ = let__13302

tmp14566 := GConsP(__e, let__13302)

var ifres14486 Obj

if True == tmp14566 {
tmp14487 := GHd(__e, let__13302)

tmp14488 := Call2(__e, SymFn(symshen_4lazyderef), tmp14487, V5035)


let__13303 := tmp14488
_ = let__13303

tmp14565 := GEq(__e, let__13303, sym_8p)

var ifres14489 Obj

if True == tmp14565 {
tmp14490 := GTl(__e, let__13302)

tmp14491 := Call2(__e, SymFn(symshen_4lazyderef), tmp14490, V5035)


let__13304 := tmp14491
_ = let__13304

tmp14564 := GConsP(__e, let__13304)

var ifres14492 Obj

if True == tmp14564 {
tmp14493 := GHd(__e, let__13304)

let__13305 := tmp14493
_ = let__13305

tmp14494 := GTl(__e, let__13304)

tmp14495 := Call2(__e, SymFn(symshen_4lazyderef), tmp14494, V5035)


let__13306 := tmp14495
_ = let__13306

tmp14563 := GConsP(__e, let__13306)

var ifres14496 Obj

if True == tmp14563 {
tmp14497 := GHd(__e, let__13306)

let__13307 := tmp14497
_ = let__13307

tmp14498 := GTl(__e, let__13306)

tmp14499 := Call2(__e, SymFn(symshen_4lazyderef), tmp14498, V5035)


let__13308 := tmp14499
_ = let__13308

tmp14562 := GEq(__e, let__13308, Nil)

var ifres14500 Obj

if True == tmp14562 {
tmp14501 := Call2(__e, SymFn(symshen_4lazyderef), V5033, V5035)


let__13309 := tmp14501
_ = let__13309

tmp14502 := MakeNative1(func(__e *ControlFlow, Z5104 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5105 Obj) Obj {
tmp14503 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14503

tmp14504 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), let__13307, Z5105, V5034, V5035, V5036, let__13248, V5038)
return nil
panic("unreachable")
})

__e.TailApply(SymFn(symshen_4system_1S_1h), let__13305, Z5104, V5034, V5035, V5036, let__13248, tmp14504)
return nil


panic("unreachable")
})
panic("unreachable")
})

let__13310 := tmp14502
_ = let__13310

tmp14561 := GConsP(__e, let__13309)

var ifres14505 Obj

if True == tmp14561 {
tmp14506 := GHd(__e, let__13309)

let__13311 := tmp14506
_ = let__13311

tmp14507 := GTl(__e, let__13309)

tmp14508 := Call2(__e, SymFn(symshen_4lazyderef), tmp14507, V5035)


let__13312 := tmp14508
_ = let__13312

tmp14509 := MakeNative1(func(__e *ControlFlow, Z5109 Obj) Obj {
tmp14510 := Call1(__e, let__13310, let__13311)


return TailCall1(__e, tmp14510, Z5109)


panic("unreachable")
})

let__13313 := tmp14509
_ = let__13313

tmp14548 := GConsP(__e, let__13312)

var ifres14511 Obj

if True == tmp14548 {
tmp14512 := GHd(__e, let__13312)

tmp14513 := Call2(__e, SymFn(symshen_4lazyderef), tmp14512, V5035)


let__13314 := tmp14513
_ = let__13314

tmp14514 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14515 := GTl(__e, let__13312)

tmp14516 := Call2(__e, SymFn(symshen_4lazyderef), tmp14515, V5035)


let__13316 := tmp14516
_ = let__13316

tmp14517 := MakeNative1(func(__e *ControlFlow, Z5114 Obj) Obj {
return TailCall1(__e, let__13313, Z5114)
panic("unreachable")
})

let__13317 := tmp14517
_ = let__13317

tmp14533 := GConsP(__e, let__13316)

if True == tmp14533 {
tmp14518 := GHd(__e, let__13316)

let__13318 := tmp14518
_ = let__13318

tmp14519 := GTl(__e, let__13316)

tmp14520 := Call2(__e, SymFn(symshen_4lazyderef), tmp14519, V5035)


let__13319 := tmp14520
_ = let__13319

tmp14521 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13317, let__13318)
panic("unreachable")
})

let__13320 := tmp14521
_ = let__13320

tmp14525 := GEq(__e, let__13319, Nil)

if True == tmp14525 {
return TailCall1(__e, SymFn(symthaw), let__13320)
} else {
tmp14523 := Call1(__e, SymFn(symshen_4pvar_2), let__13319)


if True == tmp14523 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__13319, Nil, V5035, let__13320)
} else {
return False
}


}


} else {
tmp14531 := Call1(__e, SymFn(symshen_4pvar_2), let__13316)


if True == tmp14531 {
tmp14526 := Call1(__e, SymFn(symshen_4newpv), V5035)


let__13321 := tmp14526
_ = let__13321

tmp14527 := GCons(__e, let__13321, Nil)

tmp14528 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13317, let__13321)
panic("unreachable")
})

tmp14529 := Call4(__e, SymFn(symshen_4bind_b), let__13316, tmp14527, V5035, tmp14528)


return TailCall2(__e, SymFn(symshen_4gc), V5035, tmp14529)


} else {
return False
}


}


panic("unreachable")
})

let__13315 := tmp14514
_ = let__13315

tmp14539 := GEq(__e, let__13314, sym_d)

var ifres14534 Obj

if True == tmp14539 {
tmp14535 := Call1(__e, SymFn(symthaw), let__13315)


ifres14534 = tmp14535


} else {
tmp14538 := Call1(__e, SymFn(symshen_4pvar_2), let__13314)


var ifres14536 Obj

if True == tmp14538 {
tmp14537 := Call4(__e, SymFn(symshen_4bind_b), let__13314, sym_d, V5035, let__13315)


ifres14536 = tmp14537


} else {
ifres14536 = False


}

ifres14534 = ifres14536


}

ifres14511 = ifres14534


} else {
tmp14547 := Call1(__e, SymFn(symshen_4pvar_2), let__13312)


var ifres14540 Obj

if True == tmp14547 {
tmp14541 := Call1(__e, SymFn(symshen_4newpv), V5035)


let__13322 := tmp14541
_ = let__13322

tmp14542 := GCons(__e, let__13322, Nil)

tmp14543 := GCons(__e, sym_d, tmp14542)

tmp14544 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13313, let__13322)
panic("unreachable")
})

tmp14545 := Call4(__e, SymFn(symshen_4bind_b), let__13312, tmp14543, V5035, tmp14544)


tmp14546 := Call2(__e, SymFn(symshen_4gc), V5035, tmp14545)


ifres14540 = tmp14546


} else {
ifres14540 = False


}

ifres14511 = ifres14540


}

ifres14505 = ifres14511


} else {
tmp14560 := Call1(__e, SymFn(symshen_4pvar_2), let__13309)


var ifres14549 Obj

if True == tmp14560 {
tmp14550 := Call1(__e, SymFn(symshen_4newpv), V5035)


let__13323 := tmp14550
_ = let__13323

tmp14551 := Call1(__e, SymFn(symshen_4newpv), V5035)


let__13324 := tmp14551
_ = let__13324

tmp14552 := GCons(__e, let__13324, Nil)

tmp14553 := GCons(__e, sym_d, tmp14552)

tmp14554 := GCons(__e, let__13323, tmp14553)

tmp14555 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14556 := Call1(__e, let__13310, let__13323)


return TailCall1(__e, tmp14556, let__13324)


panic("unreachable")
})

tmp14557 := Call4(__e, SymFn(symshen_4bind_b), let__13309, tmp14554, V5035, tmp14555)


tmp14558 := Call2(__e, SymFn(symshen_4gc), V5035, tmp14557)


tmp14559 := Call2(__e, SymFn(symshen_4gc), V5035, tmp14558)


ifres14549 = tmp14559


} else {
ifres14549 = False


}

ifres14505 = ifres14549


}

ifres14500 = ifres14505


} else {
ifres14500 = False


}

ifres14496 = ifres14500


} else {
ifres14496 = False


}

ifres14492 = ifres14496


} else {
ifres14492 = False


}

ifres14489 = ifres14492


} else {
ifres14489 = False


}

ifres14486 = ifres14489


} else {
ifres14486 = False


}

ifres14484 = ifres14486


} else {
ifres14484 = False


}

let__13301 := ifres14484
_ = let__13301

tmp15015 := GEq(__e, let__13301, False)

if True == tmp15015 {
tmp14634 := Call1(__e, SymFn(symshen_4unlocked_2), V5036)


var ifres14568 Obj

if True == tmp14634 {
tmp14569 := Call2(__e, SymFn(symshen_4lazyderef), V5032, V5035)


let__13326 := tmp14569
_ = let__13326

tmp14633 := GConsP(__e, let__13326)

var ifres14570 Obj

if True == tmp14633 {
tmp14571 := GHd(__e, let__13326)

tmp14572 := Call2(__e, SymFn(symshen_4lazyderef), tmp14571, V5035)


let__13327 := tmp14572
_ = let__13327

tmp14632 := GEq(__e, let__13327, sym_8v)

var ifres14573 Obj

if True == tmp14632 {
tmp14574 := GTl(__e, let__13326)

tmp14575 := Call2(__e, SymFn(symshen_4lazyderef), tmp14574, V5035)


let__13328 := tmp14575
_ = let__13328

tmp14631 := GConsP(__e, let__13328)

var ifres14576 Obj

if True == tmp14631 {
tmp14577 := GHd(__e, let__13328)

let__13329 := tmp14577
_ = let__13329

tmp14578 := GTl(__e, let__13328)

tmp14579 := Call2(__e, SymFn(symshen_4lazyderef), tmp14578, V5035)


let__13330 := tmp14579
_ = let__13330

tmp14630 := GConsP(__e, let__13330)

var ifres14580 Obj

if True == tmp14630 {
tmp14581 := GHd(__e, let__13330)

let__13331 := tmp14581
_ = let__13331

tmp14582 := GTl(__e, let__13330)

tmp14583 := Call2(__e, SymFn(symshen_4lazyderef), tmp14582, V5035)


let__13332 := tmp14583
_ = let__13332

tmp14629 := GEq(__e, let__13332, Nil)

var ifres14584 Obj

if True == tmp14629 {
tmp14585 := Call2(__e, SymFn(symshen_4lazyderef), V5033, V5035)


let__13333 := tmp14585
_ = let__13333

tmp14586 := MakeNative1(func(__e *ControlFlow, Z5132 Obj) Obj {
tmp14587 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14587

tmp14588 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14589 := GCons(__e, Z5132, Nil)

tmp14590 := GCons(__e, symvector, tmp14589)

__e.TailApply(SymFn(symshen_4system_1S_1h), let__13331, tmp14590, V5034, V5035, V5036, let__13248, V5038)
return nil


panic("unreachable")
})

__e.TailApply(SymFn(symshen_4system_1S_1h), let__13329, Z5132, V5034, V5035, V5036, let__13248, tmp14588)
return nil


panic("unreachable")
})

let__13334 := tmp14586
_ = let__13334

tmp14628 := GConsP(__e, let__13333)

var ifres14591 Obj

if True == tmp14628 {
tmp14592 := GHd(__e, let__13333)

tmp14593 := Call2(__e, SymFn(symshen_4lazyderef), tmp14592, V5035)


let__13335 := tmp14593
_ = let__13335

tmp14594 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14595 := GTl(__e, let__13333)

tmp14596 := Call2(__e, SymFn(symshen_4lazyderef), tmp14595, V5035)


let__13337 := tmp14596
_ = let__13337

tmp14597 := MakeNative1(func(__e *ControlFlow, Z5137 Obj) Obj {
return TailCall1(__e, let__13334, Z5137)
panic("unreachable")
})

let__13338 := tmp14597
_ = let__13338

tmp14613 := GConsP(__e, let__13337)

if True == tmp14613 {
tmp14598 := GHd(__e, let__13337)

let__13339 := tmp14598
_ = let__13339

tmp14599 := GTl(__e, let__13337)

tmp14600 := Call2(__e, SymFn(symshen_4lazyderef), tmp14599, V5035)


let__13340 := tmp14600
_ = let__13340

tmp14601 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13338, let__13339)
panic("unreachable")
})

let__13341 := tmp14601
_ = let__13341

tmp14605 := GEq(__e, let__13340, Nil)

if True == tmp14605 {
return TailCall1(__e, SymFn(symthaw), let__13341)
} else {
tmp14603 := Call1(__e, SymFn(symshen_4pvar_2), let__13340)


if True == tmp14603 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__13340, Nil, V5035, let__13341)
} else {
return False
}


}


} else {
tmp14611 := Call1(__e, SymFn(symshen_4pvar_2), let__13337)


if True == tmp14611 {
tmp14606 := Call1(__e, SymFn(symshen_4newpv), V5035)


let__13342 := tmp14606
_ = let__13342

tmp14607 := GCons(__e, let__13342, Nil)

tmp14608 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13338, let__13342)
panic("unreachable")
})

tmp14609 := Call4(__e, SymFn(symshen_4bind_b), let__13337, tmp14607, V5035, tmp14608)


return TailCall2(__e, SymFn(symshen_4gc), V5035, tmp14609)


} else {
return False
}


}


panic("unreachable")
})

let__13336 := tmp14594
_ = let__13336

tmp14619 := GEq(__e, let__13335, symvector)

var ifres14614 Obj

if True == tmp14619 {
tmp14615 := Call1(__e, SymFn(symthaw), let__13336)


ifres14614 = tmp14615


} else {
tmp14618 := Call1(__e, SymFn(symshen_4pvar_2), let__13335)


var ifres14616 Obj

if True == tmp14618 {
tmp14617 := Call4(__e, SymFn(symshen_4bind_b), let__13335, symvector, V5035, let__13336)


ifres14616 = tmp14617


} else {
ifres14616 = False


}

ifres14614 = ifres14616


}

ifres14591 = ifres14614


} else {
tmp14627 := Call1(__e, SymFn(symshen_4pvar_2), let__13333)


var ifres14620 Obj

if True == tmp14627 {
tmp14621 := Call1(__e, SymFn(symshen_4newpv), V5035)


let__13343 := tmp14621
_ = let__13343

tmp14622 := GCons(__e, let__13343, Nil)

tmp14623 := GCons(__e, symvector, tmp14622)

tmp14624 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13334, let__13343)
panic("unreachable")
})

tmp14625 := Call4(__e, SymFn(symshen_4bind_b), let__13333, tmp14623, V5035, tmp14624)


tmp14626 := Call2(__e, SymFn(symshen_4gc), V5035, tmp14625)


ifres14620 = tmp14626


} else {
ifres14620 = False


}

ifres14591 = ifres14620


}

ifres14584 = ifres14591


} else {
ifres14584 = False


}

ifres14580 = ifres14584


} else {
ifres14580 = False


}

ifres14576 = ifres14580


} else {
ifres14576 = False


}

ifres14573 = ifres14576


} else {
ifres14573 = False


}

ifres14570 = ifres14573


} else {
ifres14570 = False


}

ifres14568 = ifres14570


} else {
ifres14568 = False


}

let__13325 := ifres14568
_ = let__13325

tmp15013 := GEq(__e, let__13325, False)

if True == tmp15013 {
tmp14667 := Call1(__e, SymFn(symshen_4unlocked_2), V5036)


var ifres14635 Obj

if True == tmp14667 {
tmp14636 := Call2(__e, SymFn(symshen_4lazyderef), V5032, V5035)


let__13345 := tmp14636
_ = let__13345

tmp14666 := GConsP(__e, let__13345)

var ifres14637 Obj

if True == tmp14666 {
tmp14638 := GHd(__e, let__13345)

tmp14639 := Call2(__e, SymFn(symshen_4lazyderef), tmp14638, V5035)


let__13346 := tmp14639
_ = let__13346

tmp14665 := GEq(__e, let__13346, sym_8s)

var ifres14640 Obj

if True == tmp14665 {
tmp14641 := GTl(__e, let__13345)

tmp14642 := Call2(__e, SymFn(symshen_4lazyderef), tmp14641, V5035)


let__13347 := tmp14642
_ = let__13347

tmp14664 := GConsP(__e, let__13347)

var ifres14643 Obj

if True == tmp14664 {
tmp14644 := GHd(__e, let__13347)

let__13348 := tmp14644
_ = let__13348

tmp14645 := GTl(__e, let__13347)

tmp14646 := Call2(__e, SymFn(symshen_4lazyderef), tmp14645, V5035)


let__13349 := tmp14646
_ = let__13349

tmp14663 := GConsP(__e, let__13349)

var ifres14647 Obj

if True == tmp14663 {
tmp14648 := GHd(__e, let__13349)

let__13350 := tmp14648
_ = let__13350

tmp14649 := GTl(__e, let__13349)

tmp14650 := Call2(__e, SymFn(symshen_4lazyderef), tmp14649, V5035)


let__13351 := tmp14650
_ = let__13351

tmp14662 := GEq(__e, let__13351, Nil)

var ifres14651 Obj

if True == tmp14662 {
tmp14652 := Call2(__e, SymFn(symshen_4lazyderef), V5033, V5035)


let__13352 := tmp14652
_ = let__13352

tmp14653 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14654 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14654

tmp14655 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), let__13350, symstring, V5034, V5035, V5036, let__13248, V5038)
return nil
panic("unreachable")
})

__e.TailApply(SymFn(symshen_4system_1S_1h), let__13348, symstring, V5034, V5035, V5036, let__13248, tmp14655)
return nil


panic("unreachable")
})

let__13353 := tmp14653
_ = let__13353

tmp14661 := GEq(__e, let__13352, symstring)

var ifres14656 Obj

if True == tmp14661 {
tmp14657 := Call1(__e, SymFn(symthaw), let__13353)


ifres14656 = tmp14657


} else {
tmp14660 := Call1(__e, SymFn(symshen_4pvar_2), let__13352)


var ifres14658 Obj

if True == tmp14660 {
tmp14659 := Call4(__e, SymFn(symshen_4bind_b), let__13352, symstring, V5035, let__13353)


ifres14658 = tmp14659


} else {
ifres14658 = False


}

ifres14656 = ifres14658


}

ifres14651 = ifres14656


} else {
ifres14651 = False


}

ifres14647 = ifres14651


} else {
ifres14647 = False


}

ifres14643 = ifres14647


} else {
ifres14643 = False


}

ifres14640 = ifres14643


} else {
ifres14640 = False


}

ifres14637 = ifres14640


} else {
ifres14637 = False


}

ifres14635 = ifres14637


} else {
ifres14635 = False


}

let__13344 := ifres14635
_ = let__13344

tmp15011 := GEq(__e, let__13344, False)

if True == tmp15011 {
tmp14767 := Call1(__e, SymFn(symshen_4unlocked_2), V5036)


var ifres14668 Obj

if True == tmp14767 {
tmp14669 := Call2(__e, SymFn(symshen_4lazyderef), V5032, V5035)


let__13355 := tmp14669
_ = let__13355

tmp14766 := GConsP(__e, let__13355)

var ifres14670 Obj

if True == tmp14766 {
tmp14671 := GHd(__e, let__13355)

tmp14672 := Call2(__e, SymFn(symshen_4lazyderef), tmp14671, V5035)


let__13356 := tmp14672
_ = let__13356

tmp14765 := GEq(__e, let__13356, symlambda)

var ifres14673 Obj

if True == tmp14765 {
tmp14674 := GTl(__e, let__13355)

tmp14675 := Call2(__e, SymFn(symshen_4lazyderef), tmp14674, V5035)


let__13357 := tmp14675
_ = let__13357

tmp14764 := GConsP(__e, let__13357)

var ifres14676 Obj

if True == tmp14764 {
tmp14677 := GHd(__e, let__13357)

let__13358 := tmp14677
_ = let__13358

tmp14678 := GTl(__e, let__13357)

tmp14679 := Call2(__e, SymFn(symshen_4lazyderef), tmp14678, V5035)


let__13359 := tmp14679
_ = let__13359

tmp14763 := GConsP(__e, let__13359)

var ifres14680 Obj

if True == tmp14763 {
tmp14681 := GHd(__e, let__13359)

let__13360 := tmp14681
_ = let__13360

tmp14682 := GTl(__e, let__13359)

tmp14683 := Call2(__e, SymFn(symshen_4lazyderef), tmp14682, V5035)


let__13361 := tmp14683
_ = let__13361

tmp14762 := GEq(__e, let__13361, Nil)

var ifres14684 Obj

if True == tmp14762 {
tmp14685 := Call2(__e, SymFn(symshen_4lazyderef), V5033, V5035)


let__13362 := tmp14685
_ = let__13362

tmp14686 := MakeNative1(func(__e *ControlFlow, Z5163 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5164 Obj) Obj {
tmp14687 := Call1(__e, SymFn(symshen_4newpv), V5035)


let__13364 := tmp14687
_ = let__13364

tmp14688 := Call1(__e, SymFn(symshen_4newpv), V5035)


let__13365 := tmp14688
_ = let__13365

tmp14689 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14689

tmp14690 := Call2(__e, SymFn(symshen_4lazyderef), let__13358, V5035)


tmp14691 := Call1(__e, SymFn(symshen_4freshterm), tmp14690)


tmp14692 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14693 := Call2(__e, SymFn(symshen_4lazyderef), let__13358, V5035)


tmp14694 := Call2(__e, SymFn(symshen_4deref), let__13365, V5035)


tmp14695 := Call2(__e, SymFn(symshen_4deref), let__13360, V5035)


tmp14696 := Call3(__e, SymFn(symshen_4beta), tmp14693, tmp14694, tmp14695)


tmp14697 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14698 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp14699 := GCons(__e, Z5163, Nil)

tmp14700 := GCons(__e, tmp14698, tmp14699)

tmp14701 := GCons(__e, let__13365, tmp14700)

tmp14702 := GCons(__e, tmp14701, V5034)

__e.TailApply(SymFn(symshen_4system_1S_1h), let__13364, Z5164, tmp14702, V5035, V5036, let__13248, V5038)
return nil


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), let__13364, tmp14696, V5035, V5036, let__13248, tmp14697)


panic("unreachable")
})

tmp14703 := Call6(__e, SymFn(symbind), let__13365, tmp14691, V5035, V5036, let__13248, tmp14692)


tmp14704 := Call2(__e, SymFn(symshen_4gc), V5035, tmp14703)


return TailCall2(__e, SymFn(symshen_4gc), V5035, tmp14704)


panic("unreachable")
})
panic("unreachable")
})

let__13363 := tmp14686
_ = let__13363

tmp14761 := GConsP(__e, let__13362)

var ifres14705 Obj

if True == tmp14761 {
tmp14706 := GHd(__e, let__13362)

let__13366 := tmp14706
_ = let__13366

tmp14707 := GTl(__e, let__13362)

tmp14708 := Call2(__e, SymFn(symshen_4lazyderef), tmp14707, V5035)


let__13367 := tmp14708
_ = let__13367

tmp14709 := MakeNative1(func(__e *ControlFlow, Z5170 Obj) Obj {
tmp14710 := Call1(__e, let__13363, let__13366)


return TailCall1(__e, tmp14710, Z5170)


panic("unreachable")
})

let__13368 := tmp14709
_ = let__13368

tmp14748 := GConsP(__e, let__13367)

var ifres14711 Obj

if True == tmp14748 {
tmp14712 := GHd(__e, let__13367)

tmp14713 := Call2(__e, SymFn(symshen_4lazyderef), tmp14712, V5035)


let__13369 := tmp14713
_ = let__13369

tmp14714 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14715 := GTl(__e, let__13367)

tmp14716 := Call2(__e, SymFn(symshen_4lazyderef), tmp14715, V5035)


let__13371 := tmp14716
_ = let__13371

tmp14717 := MakeNative1(func(__e *ControlFlow, Z5175 Obj) Obj {
return TailCall1(__e, let__13368, Z5175)
panic("unreachable")
})

let__13372 := tmp14717
_ = let__13372

tmp14733 := GConsP(__e, let__13371)

if True == tmp14733 {
tmp14718 := GHd(__e, let__13371)

let__13373 := tmp14718
_ = let__13373

tmp14719 := GTl(__e, let__13371)

tmp14720 := Call2(__e, SymFn(symshen_4lazyderef), tmp14719, V5035)


let__13374 := tmp14720
_ = let__13374

tmp14721 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13372, let__13373)
panic("unreachable")
})

let__13375 := tmp14721
_ = let__13375

tmp14725 := GEq(__e, let__13374, Nil)

if True == tmp14725 {
return TailCall1(__e, SymFn(symthaw), let__13375)
} else {
tmp14723 := Call1(__e, SymFn(symshen_4pvar_2), let__13374)


if True == tmp14723 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__13374, Nil, V5035, let__13375)
} else {
return False
}


}


} else {
tmp14731 := Call1(__e, SymFn(symshen_4pvar_2), let__13371)


if True == tmp14731 {
tmp14726 := Call1(__e, SymFn(symshen_4newpv), V5035)


let__13376 := tmp14726
_ = let__13376

tmp14727 := GCons(__e, let__13376, Nil)

tmp14728 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13372, let__13376)
panic("unreachable")
})

tmp14729 := Call4(__e, SymFn(symshen_4bind_b), let__13371, tmp14727, V5035, tmp14728)


return TailCall2(__e, SymFn(symshen_4gc), V5035, tmp14729)


} else {
return False
}


}


panic("unreachable")
})

let__13370 := tmp14714
_ = let__13370

tmp14739 := GEq(__e, let__13369, sym_1_1_6)

var ifres14734 Obj

if True == tmp14739 {
tmp14735 := Call1(__e, SymFn(symthaw), let__13370)


ifres14734 = tmp14735


} else {
tmp14738 := Call1(__e, SymFn(symshen_4pvar_2), let__13369)


var ifres14736 Obj

if True == tmp14738 {
tmp14737 := Call4(__e, SymFn(symshen_4bind_b), let__13369, sym_1_1_6, V5035, let__13370)


ifres14736 = tmp14737


} else {
ifres14736 = False


}

ifres14734 = ifres14736


}

ifres14711 = ifres14734


} else {
tmp14747 := Call1(__e, SymFn(symshen_4pvar_2), let__13367)


var ifres14740 Obj

if True == tmp14747 {
tmp14741 := Call1(__e, SymFn(symshen_4newpv), V5035)


let__13377 := tmp14741
_ = let__13377

tmp14742 := GCons(__e, let__13377, Nil)

tmp14743 := GCons(__e, sym_1_1_6, tmp14742)

tmp14744 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13368, let__13377)
panic("unreachable")
})

tmp14745 := Call4(__e, SymFn(symshen_4bind_b), let__13367, tmp14743, V5035, tmp14744)


tmp14746 := Call2(__e, SymFn(symshen_4gc), V5035, tmp14745)


ifres14740 = tmp14746


} else {
ifres14740 = False


}

ifres14711 = ifres14740


}

ifres14705 = ifres14711


} else {
tmp14760 := Call1(__e, SymFn(symshen_4pvar_2), let__13362)


var ifres14749 Obj

if True == tmp14760 {
tmp14750 := Call1(__e, SymFn(symshen_4newpv), V5035)


let__13378 := tmp14750
_ = let__13378

tmp14751 := Call1(__e, SymFn(symshen_4newpv), V5035)


let__13379 := tmp14751
_ = let__13379

tmp14752 := GCons(__e, let__13379, Nil)

tmp14753 := GCons(__e, sym_1_1_6, tmp14752)

tmp14754 := GCons(__e, let__13378, tmp14753)

tmp14755 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14756 := Call1(__e, let__13363, let__13378)


return TailCall1(__e, tmp14756, let__13379)


panic("unreachable")
})

tmp14757 := Call4(__e, SymFn(symshen_4bind_b), let__13362, tmp14754, V5035, tmp14755)


tmp14758 := Call2(__e, SymFn(symshen_4gc), V5035, tmp14757)


tmp14759 := Call2(__e, SymFn(symshen_4gc), V5035, tmp14758)


ifres14749 = tmp14759


} else {
ifres14749 = False


}

ifres14705 = ifres14749


}

ifres14684 = ifres14705


} else {
ifres14684 = False


}

ifres14680 = ifres14684


} else {
ifres14680 = False


}

ifres14676 = ifres14680


} else {
ifres14676 = False


}

ifres14673 = ifres14676


} else {
ifres14673 = False


}

ifres14670 = ifres14673


} else {
ifres14670 = False


}

ifres14668 = ifres14670


} else {
ifres14668 = False


}

let__13354 := ifres14668
_ = let__13354

tmp15009 := GEq(__e, let__13354, False)

if True == tmp15009 {
tmp14817 := Call1(__e, SymFn(symshen_4unlocked_2), V5036)


var ifres14768 Obj

if True == tmp14817 {
tmp14769 := Call2(__e, SymFn(symshen_4lazyderef), V5032, V5035)


let__13381 := tmp14769
_ = let__13381

tmp14816 := GConsP(__e, let__13381)

var ifres14770 Obj

if True == tmp14816 {
tmp14771 := GHd(__e, let__13381)

tmp14772 := Call2(__e, SymFn(symshen_4lazyderef), tmp14771, V5035)


let__13382 := tmp14772
_ = let__13382

tmp14815 := GEq(__e, let__13382, symlet)

var ifres14773 Obj

if True == tmp14815 {
tmp14774 := GTl(__e, let__13381)

tmp14775 := Call2(__e, SymFn(symshen_4lazyderef), tmp14774, V5035)


let__13383 := tmp14775
_ = let__13383

tmp14814 := GConsP(__e, let__13383)

var ifres14776 Obj

if True == tmp14814 {
tmp14777 := GHd(__e, let__13383)

let__13384 := tmp14777
_ = let__13384

tmp14778 := GTl(__e, let__13383)

tmp14779 := Call2(__e, SymFn(symshen_4lazyderef), tmp14778, V5035)


let__13385 := tmp14779
_ = let__13385

tmp14813 := GConsP(__e, let__13385)

var ifres14780 Obj

if True == tmp14813 {
tmp14781 := GHd(__e, let__13385)

let__13386 := tmp14781
_ = let__13386

tmp14782 := GTl(__e, let__13385)

tmp14783 := Call2(__e, SymFn(symshen_4lazyderef), tmp14782, V5035)


let__13387 := tmp14783
_ = let__13387

tmp14812 := GConsP(__e, let__13387)

var ifres14784 Obj

if True == tmp14812 {
tmp14785 := GHd(__e, let__13387)

let__13388 := tmp14785
_ = let__13388

tmp14786 := GTl(__e, let__13387)

tmp14787 := Call2(__e, SymFn(symshen_4lazyderef), tmp14786, V5035)


let__13389 := tmp14787
_ = let__13389

tmp14811 := GEq(__e, let__13389, Nil)

var ifres14788 Obj

if True == tmp14811 {
tmp14789 := Call1(__e, SymFn(symshen_4newpv), V5035)


let__13390 := tmp14789
_ = let__13390

tmp14790 := Call1(__e, SymFn(symshen_4newpv), V5035)


let__13391 := tmp14790
_ = let__13391

tmp14791 := Call1(__e, SymFn(symshen_4newpv), V5035)


let__13392 := tmp14791
_ = let__13392

tmp14792 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14792

tmp14793 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14794 := Call2(__e, SymFn(symshen_4lazyderef), let__13384, V5035)


tmp14795 := Call1(__e, SymFn(symshen_4freshterm), tmp14794)


tmp14796 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14797 := Call2(__e, SymFn(symshen_4lazyderef), let__13384, V5035)


tmp14798 := Call2(__e, SymFn(symshen_4lazyderef), let__13391, V5035)


tmp14799 := Call2(__e, SymFn(symshen_4lazyderef), let__13388, V5035)


tmp14800 := Call3(__e, SymFn(symshen_4beta), tmp14797, tmp14798, tmp14799)


tmp14801 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14802 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp14803 := GCons(__e, let__13392, Nil)

tmp14804 := GCons(__e, tmp14802, tmp14803)

tmp14805 := GCons(__e, let__13391, tmp14804)

tmp14806 := GCons(__e, tmp14805, V5034)

__e.TailApply(SymFn(symshen_4system_1S_1h), let__13390, V5033, tmp14806, V5035, V5036, let__13248, V5038)
return nil


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), let__13390, tmp14800, V5035, V5036, let__13248, tmp14801)


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), let__13391, tmp14795, V5035, V5036, let__13248, tmp14796)


panic("unreachable")
})

tmp14807 := Call(__e, SymFn(symshen_4system_1S_1h), let__13386, let__13392, V5034, V5035, V5036, let__13248, tmp14793)


tmp14808 := Call2(__e, SymFn(symshen_4gc), V5035, tmp14807)


tmp14809 := Call2(__e, SymFn(symshen_4gc), V5035, tmp14808)


tmp14810 := Call2(__e, SymFn(symshen_4gc), V5035, tmp14809)


ifres14788 = tmp14810


} else {
ifres14788 = False


}

ifres14784 = ifres14788


} else {
ifres14784 = False


}

ifres14780 = ifres14784


} else {
ifres14780 = False


}

ifres14776 = ifres14780


} else {
ifres14776 = False


}

ifres14773 = ifres14776


} else {
ifres14773 = False


}

ifres14770 = ifres14773


} else {
ifres14770 = False


}

ifres14768 = ifres14770


} else {
ifres14768 = False


}

let__13380 := ifres14768
_ = let__13380

tmp15007 := GEq(__e, let__13380, False)

if True == tmp15007 {
tmp14887 := Call1(__e, SymFn(symshen_4unlocked_2), V5036)


var ifres14818 Obj

if True == tmp14887 {
tmp14819 := Call2(__e, SymFn(symshen_4lazyderef), V5032, V5035)


let__13394 := tmp14819
_ = let__13394

tmp14886 := GConsP(__e, let__13394)

var ifres14820 Obj

if True == tmp14886 {
tmp14821 := GHd(__e, let__13394)

tmp14822 := Call2(__e, SymFn(symshen_4lazyderef), tmp14821, V5035)


let__13395 := tmp14822
_ = let__13395

tmp14885 := GEq(__e, let__13395, symopen)

var ifres14823 Obj

if True == tmp14885 {
tmp14824 := GTl(__e, let__13394)

tmp14825 := Call2(__e, SymFn(symshen_4lazyderef), tmp14824, V5035)


let__13396 := tmp14825
_ = let__13396

tmp14884 := GConsP(__e, let__13396)

var ifres14826 Obj

if True == tmp14884 {
tmp14827 := GHd(__e, let__13396)

let__13397 := tmp14827
_ = let__13397

tmp14828 := GTl(__e, let__13396)

tmp14829 := Call2(__e, SymFn(symshen_4lazyderef), tmp14828, V5035)


let__13398 := tmp14829
_ = let__13398

tmp14883 := GConsP(__e, let__13398)

var ifres14830 Obj

if True == tmp14883 {
tmp14831 := GHd(__e, let__13398)

let__13399 := tmp14831
_ = let__13399

tmp14832 := GTl(__e, let__13398)

tmp14833 := Call2(__e, SymFn(symshen_4lazyderef), tmp14832, V5035)


let__13400 := tmp14833
_ = let__13400

tmp14882 := GEq(__e, let__13400, Nil)

var ifres14834 Obj

if True == tmp14882 {
tmp14835 := Call2(__e, SymFn(symshen_4lazyderef), V5033, V5035)


let__13401 := tmp14835
_ = let__13401

tmp14836 := MakeNative1(func(__e *ControlFlow, Z5206 Obj) Obj {
tmp14837 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14837

tmp14838 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14839 := Call2(__e, SymFn(symshen_4lazyderef), Z5206, V5035)


tmp14840 := GCons(__e, symout, Nil)

tmp14841 := GCons(__e, symin, tmp14840)

tmp14842 := Call2(__e, SymFn(symelement_2), tmp14839, tmp14841)


tmp14843 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), let__13397, symstring, V5034, V5035, V5036, let__13248, V5038)
return nil
panic("unreachable")
})

return TailCall5(__e, SymFn(symwhen), tmp14842, V5035, V5036, let__13248, tmp14843)


panic("unreachable")
})

return TailCall6(__e, SymFn(symis_b), let__13399, Z5206, V5035, V5036, let__13248, tmp14838)


panic("unreachable")
})

let__13402 := tmp14836
_ = let__13402

tmp14881 := GConsP(__e, let__13401)

var ifres14844 Obj

if True == tmp14881 {
tmp14845 := GHd(__e, let__13401)

tmp14846 := Call2(__e, SymFn(symshen_4lazyderef), tmp14845, V5035)


let__13403 := tmp14846
_ = let__13403

tmp14847 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14848 := GTl(__e, let__13401)

tmp14849 := Call2(__e, SymFn(symshen_4lazyderef), tmp14848, V5035)


let__13405 := tmp14849
_ = let__13405

tmp14850 := MakeNative1(func(__e *ControlFlow, Z5211 Obj) Obj {
return TailCall1(__e, let__13402, Z5211)
panic("unreachable")
})

let__13406 := tmp14850
_ = let__13406

tmp14866 := GConsP(__e, let__13405)

if True == tmp14866 {
tmp14851 := GHd(__e, let__13405)

let__13407 := tmp14851
_ = let__13407

tmp14852 := GTl(__e, let__13405)

tmp14853 := Call2(__e, SymFn(symshen_4lazyderef), tmp14852, V5035)


let__13408 := tmp14853
_ = let__13408

tmp14854 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13406, let__13407)
panic("unreachable")
})

let__13409 := tmp14854
_ = let__13409

tmp14858 := GEq(__e, let__13408, Nil)

if True == tmp14858 {
return TailCall1(__e, SymFn(symthaw), let__13409)
} else {
tmp14856 := Call1(__e, SymFn(symshen_4pvar_2), let__13408)


if True == tmp14856 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__13408, Nil, V5035, let__13409)
} else {
return False
}


}


} else {
tmp14864 := Call1(__e, SymFn(symshen_4pvar_2), let__13405)


if True == tmp14864 {
tmp14859 := Call1(__e, SymFn(symshen_4newpv), V5035)


let__13410 := tmp14859
_ = let__13410

tmp14860 := GCons(__e, let__13410, Nil)

tmp14861 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13406, let__13410)
panic("unreachable")
})

tmp14862 := Call4(__e, SymFn(symshen_4bind_b), let__13405, tmp14860, V5035, tmp14861)


return TailCall2(__e, SymFn(symshen_4gc), V5035, tmp14862)


} else {
return False
}


}


panic("unreachable")
})

let__13404 := tmp14847
_ = let__13404

tmp14872 := GEq(__e, let__13403, symstream)

var ifres14867 Obj

if True == tmp14872 {
tmp14868 := Call1(__e, SymFn(symthaw), let__13404)


ifres14867 = tmp14868


} else {
tmp14871 := Call1(__e, SymFn(symshen_4pvar_2), let__13403)


var ifres14869 Obj

if True == tmp14871 {
tmp14870 := Call4(__e, SymFn(symshen_4bind_b), let__13403, symstream, V5035, let__13404)


ifres14869 = tmp14870


} else {
ifres14869 = False


}

ifres14867 = ifres14869


}

ifres14844 = ifres14867


} else {
tmp14880 := Call1(__e, SymFn(symshen_4pvar_2), let__13401)


var ifres14873 Obj

if True == tmp14880 {
tmp14874 := Call1(__e, SymFn(symshen_4newpv), V5035)


let__13411 := tmp14874
_ = let__13411

tmp14875 := GCons(__e, let__13411, Nil)

tmp14876 := GCons(__e, symstream, tmp14875)

tmp14877 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13402, let__13411)
panic("unreachable")
})

tmp14878 := Call4(__e, SymFn(symshen_4bind_b), let__13401, tmp14876, V5035, tmp14877)


tmp14879 := Call2(__e, SymFn(symshen_4gc), V5035, tmp14878)


ifres14873 = tmp14879


} else {
ifres14873 = False


}

ifres14844 = ifres14873


}

ifres14834 = ifres14844


} else {
ifres14834 = False


}

ifres14830 = ifres14834


} else {
ifres14830 = False


}

ifres14826 = ifres14830


} else {
ifres14826 = False


}

ifres14823 = ifres14826


} else {
ifres14823 = False


}

ifres14820 = ifres14823


} else {
ifres14820 = False


}

ifres14818 = ifres14820


} else {
ifres14818 = False


}

let__13393 := ifres14818
_ = let__13393

tmp15005 := GEq(__e, let__13393, False)

if True == tmp15005 {
tmp14916 := Call1(__e, SymFn(symshen_4unlocked_2), V5036)


var ifres14888 Obj

if True == tmp14916 {
tmp14889 := Call2(__e, SymFn(symshen_4lazyderef), V5032, V5035)


let__13413 := tmp14889
_ = let__13413

tmp14915 := GConsP(__e, let__13413)

var ifres14890 Obj

if True == tmp14915 {
tmp14891 := GHd(__e, let__13413)

tmp14892 := Call2(__e, SymFn(symshen_4lazyderef), tmp14891, V5035)


let__13414 := tmp14892
_ = let__13414

tmp14914 := GEq(__e, let__13414, symtype)

var ifres14893 Obj

if True == tmp14914 {
tmp14894 := GTl(__e, let__13413)

tmp14895 := Call2(__e, SymFn(symshen_4lazyderef), tmp14894, V5035)


let__13415 := tmp14895
_ = let__13415

tmp14913 := GConsP(__e, let__13415)

var ifres14896 Obj

if True == tmp14913 {
tmp14897 := GHd(__e, let__13415)

let__13416 := tmp14897
_ = let__13416

tmp14898 := GTl(__e, let__13415)

tmp14899 := Call2(__e, SymFn(symshen_4lazyderef), tmp14898, V5035)


let__13417 := tmp14899
_ = let__13417

tmp14912 := GConsP(__e, let__13417)

var ifres14900 Obj

if True == tmp14912 {
tmp14901 := GHd(__e, let__13417)

let__13418 := tmp14901
_ = let__13418

tmp14902 := GTl(__e, let__13417)

tmp14903 := Call2(__e, SymFn(symshen_4lazyderef), tmp14902, V5035)


let__13419 := tmp14903
_ = let__13419

tmp14911 := GEq(__e, let__13419, Nil)

var ifres14904 Obj

if True == tmp14911 {
tmp14905 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14905

tmp14906 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14907 := Call2(__e, SymFn(symshen_4deref), let__13418, V5035)


tmp14908 := Call1(__e, SymFn(symshen_4rectify_1type), tmp14907)


tmp14909 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), let__13416, V5033, V5034, V5035, V5036, let__13248, V5038)
return nil
panic("unreachable")
})

return TailCall6(__e, SymFn(symis_b), tmp14908, V5033, V5035, V5036, let__13248, tmp14909)


panic("unreachable")
})

tmp14910 := Call4(__e, SymFn(symshen_4cut), V5035, V5036, let__13248, tmp14906)


ifres14904 = tmp14910


} else {
ifres14904 = False


}

ifres14900 = ifres14904


} else {
ifres14900 = False


}

ifres14896 = ifres14900


} else {
ifres14896 = False


}

ifres14893 = ifres14896


} else {
ifres14893 = False


}

ifres14890 = ifres14893


} else {
ifres14890 = False


}

ifres14888 = ifres14890


} else {
ifres14888 = False


}

let__13412 := ifres14888
_ = let__13412

tmp15003 := GEq(__e, let__13412, False)

if True == tmp15003 {
tmp14947 := Call1(__e, SymFn(symshen_4unlocked_2), V5036)


var ifres14917 Obj

if True == tmp14947 {
tmp14918 := Call2(__e, SymFn(symshen_4lazyderef), V5032, V5035)


let__13421 := tmp14918
_ = let__13421

tmp14946 := GConsP(__e, let__13421)

var ifres14919 Obj

if True == tmp14946 {
tmp14920 := GHd(__e, let__13421)

tmp14921 := Call2(__e, SymFn(symshen_4lazyderef), tmp14920, V5035)


let__13422 := tmp14921
_ = let__13422

tmp14945 := GEq(__e, let__13422, symshen_4input_1h_7)

var ifres14922 Obj

if True == tmp14945 {
tmp14923 := GTl(__e, let__13421)

tmp14924 := Call2(__e, SymFn(symshen_4lazyderef), tmp14923, V5035)


let__13423 := tmp14924
_ = let__13423

tmp14944 := GConsP(__e, let__13423)

var ifres14925 Obj

if True == tmp14944 {
tmp14926 := GHd(__e, let__13423)

let__13424 := tmp14926
_ = let__13424

tmp14927 := GTl(__e, let__13423)

tmp14928 := Call2(__e, SymFn(symshen_4lazyderef), tmp14927, V5035)


let__13425 := tmp14928
_ = let__13425

tmp14943 := GConsP(__e, let__13425)

var ifres14929 Obj

if True == tmp14943 {
tmp14930 := GHd(__e, let__13425)

let__13426 := tmp14930
_ = let__13426

tmp14931 := GTl(__e, let__13425)

tmp14932 := Call2(__e, SymFn(symshen_4lazyderef), tmp14931, V5035)


let__13427 := tmp14932
_ = let__13427

tmp14942 := GEq(__e, let__13427, Nil)

var ifres14933 Obj

if True == tmp14942 {
tmp14934 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14934

tmp14935 := Call2(__e, SymFn(symshen_4deref), let__13424, V5035)


tmp14936 := Call1(__e, SymFn(symshen_4rdecons), tmp14935)


tmp14937 := Call1(__e, SymFn(symshen_4rectify_1type), tmp14936)


tmp14938 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14939 := GCons(__e, symin, Nil)

tmp14940 := GCons(__e, symstream, tmp14939)

__e.TailApply(SymFn(symshen_4system_1S_1h), let__13426, tmp14940, V5034, V5035, V5036, let__13248, V5038)
return nil


panic("unreachable")
})

tmp14941 := Call6(__e, SymFn(symis), V5033, tmp14937, V5035, V5036, let__13248, tmp14938)


ifres14933 = tmp14941


} else {
ifres14933 = False


}

ifres14929 = ifres14933


} else {
ifres14929 = False


}

ifres14925 = ifres14929


} else {
ifres14925 = False


}

ifres14922 = ifres14925


} else {
ifres14922 = False


}

ifres14919 = ifres14922


} else {
ifres14919 = False


}

ifres14917 = ifres14919


} else {
ifres14917 = False


}

let__13420 := ifres14917
_ = let__13420

tmp15001 := GEq(__e, let__13420, False)

if True == tmp15001 {
tmp14976 := Call1(__e, SymFn(symshen_4unlocked_2), V5036)


var ifres14948 Obj

if True == tmp14976 {
tmp14949 := Call2(__e, SymFn(symshen_4lazyderef), V5032, V5035)


let__13429 := tmp14949
_ = let__13429

tmp14975 := GConsP(__e, let__13429)

var ifres14950 Obj

if True == tmp14975 {
tmp14951 := GHd(__e, let__13429)

tmp14952 := Call2(__e, SymFn(symshen_4lazyderef), tmp14951, V5035)


let__13430 := tmp14952
_ = let__13430

tmp14974 := GEq(__e, let__13430, symset)

var ifres14953 Obj

if True == tmp14974 {
tmp14954 := GTl(__e, let__13429)

tmp14955 := Call2(__e, SymFn(symshen_4lazyderef), tmp14954, V5035)


let__13431 := tmp14955
_ = let__13431

tmp14973 := GConsP(__e, let__13431)

var ifres14956 Obj

if True == tmp14973 {
tmp14957 := GHd(__e, let__13431)

let__13432 := tmp14957
_ = let__13432

tmp14958 := GTl(__e, let__13431)

tmp14959 := Call2(__e, SymFn(symshen_4lazyderef), tmp14958, V5035)


let__13433 := tmp14959
_ = let__13433

tmp14972 := GConsP(__e, let__13433)

var ifres14960 Obj

if True == tmp14972 {
tmp14961 := GHd(__e, let__13433)

let__13434 := tmp14961
_ = let__13434

tmp14962 := GTl(__e, let__13433)

tmp14963 := Call2(__e, SymFn(symshen_4lazyderef), tmp14962, V5035)


let__13435 := tmp14963
_ = let__13435

tmp14971 := GEq(__e, let__13435, Nil)

var ifres14964 Obj

if True == tmp14971 {
tmp14965 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14965

tmp14966 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14967 := GCons(__e, let__13432, Nil)

tmp14968 := GCons(__e, symvalue, tmp14967)

tmp14969 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), let__13434, V5033, V5034, V5035, V5036, let__13248, V5038)
return nil
panic("unreachable")
})

__e.TailApply(SymFn(symshen_4system_1S_1h), tmp14968, V5033, V5034, V5035, V5036, let__13248, tmp14969)
return nil


panic("unreachable")
})

tmp14970 := Call(__e, SymFn(symshen_4system_1S_1h), let__13432, symsymbol, V5034, V5035, V5036, let__13248, tmp14966)


ifres14964 = tmp14970


} else {
ifres14964 = False


}

ifres14960 = ifres14964


} else {
ifres14960 = False


}

ifres14956 = ifres14960


} else {
ifres14956 = False


}

ifres14953 = ifres14956


} else {
ifres14953 = False


}

ifres14950 = ifres14953


} else {
ifres14950 = False


}

ifres14948 = ifres14950


} else {
ifres14948 = False


}

let__13428 := ifres14948
_ = let__13428

tmp14999 := GEq(__e, let__13428, False)

if True == tmp14999 {
tmp14984 := Call1(__e, SymFn(symshen_4unlocked_2), V5036)


var ifres14977 Obj

if True == tmp14984 {
tmp14978 := Call1(__e, SymFn(symshen_4newpv), V5035)


let__13437 := tmp14978
_ = let__13437

tmp14979 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14979

tmp14980 := MakeNative0(func(__e *ControlFlow) Obj {
tmp14981 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4system_1S_1h), V5032, V5033, let__13437, V5035, V5036, let__13248, V5038)
return nil
panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5035, V5036, let__13248, tmp14981)


panic("unreachable")
})

tmp14982 := Call(__e, SymFn(symshen_4l_1rules), V5034, let__13437, False, V5035, V5036, let__13248, tmp14980)


tmp14983 := Call2(__e, SymFn(symshen_4gc), V5035, tmp14982)


ifres14977 = tmp14983


} else {
ifres14977 = False


}

let__13436 := ifres14977
_ = let__13436

tmp14997 := GEq(__e, let__13436, False)

if True == tmp14997 {
tmp14993 := Call1(__e, SymFn(symshen_4unlocked_2), V5036)


var ifres14985 Obj

if True == tmp14993 {
tmp14986 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp14986

tmp14987 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp14988 := GCons(__e, V5033, Nil)

tmp14989 := GCons(__e, tmp14987, tmp14988)

tmp14990 := GCons(__e, V5032, tmp14989)

tmp14991 := GValue(__e, symshen_4_ddatatypes_d)

tmp14992 := Call(__e, SymFn(symshen_4search_1user_1datatypes), tmp14990, V5034, tmp14991, V5035, V5036, let__13248, V5038)


ifres14985 = tmp14992


} else {
ifres14985 = False


}

let__13438 := ifres14985
_ = let__13438

tmp14995 := GEq(__e, let__13438, False)

if True == tmp14995 {
return TailCall2(__e, SymFn(symshen_4unlock), V5036, let__13248)
} else {
return let__13438
}


} else {
return let__13436
}


} else {
return let__13428
}


} else {
return let__13420
}


} else {
return let__13412
}


} else {
return let__13393
}


} else {
return let__13380
}


} else {
return let__13354
}


} else {
return let__13344
}


} else {
return let__13325
}


} else {
return let__13301
}


} else {
return let__13282
}


} else {
return let__13275
}


} else {
return let__13268
}


} else {
return let__13262
}


} else {
return let__13256
}


} else {
return let__13252
}


} else {
return let__13251
}


} else {
return let__13250
}


} else {
return let__13249
}


panic("unreachable")
}())
}, 7)

tmp15034 := Call2(__e, ns2_1set, symshen_4system_1S_1h, tmp14283)


_ = tmp15034

tmp15035 := MakeNative1(func(__e *ControlFlow, V5244 Obj) Obj {
tmp15065 := GConsP(__e, V5244)

var ifres15050 Obj

if True == tmp15065 {
tmp15063 := GHd(__e, V5244)

tmp15064 := GEq(__e, symcons, tmp15063)

var ifres15051 Obj

if True == tmp15064 {
tmp15061 := GTl(__e, V5244)

tmp15062 := GConsP(__e, tmp15061)

var ifres15052 Obj

if True == tmp15062 {
tmp15058 := GTl(__e, V5244)

tmp15059 := GTl(__e, tmp15058)

tmp15060 := GConsP(__e, tmp15059)

var ifres15053 Obj

if True == tmp15060 {
tmp15054 := GTl(__e, V5244)

tmp15055 := GTl(__e, tmp15054)

tmp15056 := GTl(__e, tmp15055)

tmp15057 := GEq(__e, Nil, tmp15056)

ifres15053 = tmp15057


} else {
ifres15053 = False


}

ifres15052 = ifres15053


} else {
ifres15052 = False


}

ifres15051 = ifres15052


} else {
ifres15051 = False


}

ifres15050 = ifres15051


} else {
ifres15050 = False


}

if True == ifres15050 {
tmp15036 := GTl(__e, V5244)

tmp15037 := GHd(__e, tmp15036)

tmp15038 := Call1(__e, SymFn(symshen_4rdecons), tmp15037)


tmp15039 := GTl(__e, V5244)

tmp15040 := GTl(__e, tmp15039)

tmp15041 := GHd(__e, tmp15040)

tmp15042 := Call1(__e, SymFn(symshen_4rdecons), tmp15041)


return GCons(__e, tmp15038, tmp15042)


} else {
tmp15048 := GConsP(__e, V5244)

if True == tmp15048 {
tmp15043 := GHd(__e, V5244)

tmp15044 := Call1(__e, SymFn(symshen_4rdecons), tmp15043)


tmp15045 := GTl(__e, V5244)

tmp15046 := Call1(__e, SymFn(symshen_4rdecons), tmp15045)


return GCons(__e, tmp15044, tmp15046)


} else {
return V5244
}


}


panic("unreachable")
})

tmp15066 := Call2(__e, ns2_1set, symshen_4rdecons, tmp15035)


_ = tmp15066

tmp15067 := MakeNative6(func(__e *ControlFlow, V5245, V5246, V5247, V5248, V5249, V5250 Obj) Obj {
tmp15080 := Call1(__e, SymFn(symshen_4unlocked_2), V5248)


var ifres15068 Obj

if True == tmp15080 {
tmp15069 := Call2(__e, SymFn(symshen_4lazyderef), V5246, V5247)


let__13440 := tmp15069
_ = let__13440

tmp15070 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15071 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15071

tmp15072 := Call2(__e, SymFn(symshen_4lazyderef), V5245, V5247)


tmp15073 := PrimGuard1(__e, symnumber_2, PrimIsNumber, tmp15072)

return TailCall5(__e, SymFn(symwhen), tmp15073, V5247, V5248, V5249, V5250)


panic("unreachable")
})

let__13441 := tmp15070
_ = let__13441

tmp15079 := GEq(__e, let__13440, symnumber)

var ifres15074 Obj

if True == tmp15079 {
tmp15075 := Call1(__e, SymFn(symthaw), let__13441)


ifres15074 = tmp15075


} else {
tmp15078 := Call1(__e, SymFn(symshen_4pvar_2), let__13440)


var ifres15076 Obj

if True == tmp15078 {
tmp15077 := Call4(__e, SymFn(symshen_4bind_b), let__13440, symnumber, V5247, let__13441)


ifres15076 = tmp15077


} else {
ifres15076 = False


}

ifres15074 = ifres15076


}

ifres15068 = ifres15074


} else {
ifres15068 = False


}

let__13439 := ifres15068
_ = let__13439

tmp15170 := GEq(__e, let__13439, False)

if True == tmp15170 {
tmp15093 := Call1(__e, SymFn(symshen_4unlocked_2), V5248)


var ifres15081 Obj

if True == tmp15093 {
tmp15082 := Call2(__e, SymFn(symshen_4lazyderef), V5246, V5247)


let__13443 := tmp15082
_ = let__13443

tmp15083 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15084 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15084

tmp15085 := Call2(__e, SymFn(symshen_4lazyderef), V5245, V5247)


tmp15086 := Call1(__e, SymFn(symboolean_2), tmp15085)


return TailCall5(__e, SymFn(symwhen), tmp15086, V5247, V5248, V5249, V5250)


panic("unreachable")
})

let__13444 := tmp15083
_ = let__13444

tmp15092 := GEq(__e, let__13443, symboolean)

var ifres15087 Obj

if True == tmp15092 {
tmp15088 := Call1(__e, SymFn(symthaw), let__13444)


ifres15087 = tmp15088


} else {
tmp15091 := Call1(__e, SymFn(symshen_4pvar_2), let__13443)


var ifres15089 Obj

if True == tmp15091 {
tmp15090 := Call4(__e, SymFn(symshen_4bind_b), let__13443, symboolean, V5247, let__13444)


ifres15089 = tmp15090


} else {
ifres15089 = False


}

ifres15087 = ifres15089


}

ifres15081 = ifres15087


} else {
ifres15081 = False


}

let__13442 := ifres15081
_ = let__13442

tmp15168 := GEq(__e, let__13442, False)

if True == tmp15168 {
tmp15106 := Call1(__e, SymFn(symshen_4unlocked_2), V5248)


var ifres15094 Obj

if True == tmp15106 {
tmp15095 := Call2(__e, SymFn(symshen_4lazyderef), V5246, V5247)


let__13446 := tmp15095
_ = let__13446

tmp15096 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15097 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15097

tmp15098 := Call2(__e, SymFn(symshen_4lazyderef), V5245, V5247)


tmp15099 := PrimGuard1(__e, symstring_2, PrimIsString, tmp15098)

return TailCall5(__e, SymFn(symwhen), tmp15099, V5247, V5248, V5249, V5250)


panic("unreachable")
})

let__13447 := tmp15096
_ = let__13447

tmp15105 := GEq(__e, let__13446, symstring)

var ifres15100 Obj

if True == tmp15105 {
tmp15101 := Call1(__e, SymFn(symthaw), let__13447)


ifres15100 = tmp15101


} else {
tmp15104 := Call1(__e, SymFn(symshen_4pvar_2), let__13446)


var ifres15102 Obj

if True == tmp15104 {
tmp15103 := Call4(__e, SymFn(symshen_4bind_b), let__13446, symstring, V5247, let__13447)


ifres15102 = tmp15103


} else {
ifres15102 = False


}

ifres15100 = ifres15102


}

ifres15094 = ifres15100


} else {
ifres15094 = False


}

let__13445 := ifres15094
_ = let__13445

tmp15166 := GEq(__e, let__13445, False)

if True == tmp15166 {
tmp15119 := Call1(__e, SymFn(symshen_4unlocked_2), V5248)


var ifres15107 Obj

if True == tmp15119 {
tmp15108 := Call2(__e, SymFn(symshen_4lazyderef), V5246, V5247)


let__13449 := tmp15108
_ = let__13449

tmp15109 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15110 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15110

tmp15111 := Call2(__e, SymFn(symshen_4lazyderef), V5245, V5247)


tmp15112 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, tmp15111)

return TailCall5(__e, SymFn(symwhen), tmp15112, V5247, V5248, V5249, V5250)


panic("unreachable")
})

let__13450 := tmp15109
_ = let__13450

tmp15118 := GEq(__e, let__13449, symsymbol)

var ifres15113 Obj

if True == tmp15118 {
tmp15114 := Call1(__e, SymFn(symthaw), let__13450)


ifres15113 = tmp15114


} else {
tmp15117 := Call1(__e, SymFn(symshen_4pvar_2), let__13449)


var ifres15115 Obj

if True == tmp15117 {
tmp15116 := Call4(__e, SymFn(symshen_4bind_b), let__13449, symsymbol, V5247, let__13450)


ifres15115 = tmp15116


} else {
ifres15115 = False


}

ifres15113 = ifres15115


}

ifres15107 = ifres15113


} else {
ifres15107 = False


}

let__13448 := ifres15107
_ = let__13448

tmp15164 := GEq(__e, let__13448, False)

if True == tmp15164 {
tmp15162 := Call1(__e, SymFn(symshen_4unlocked_2), V5248)


if True == tmp15162 {
tmp15120 := Call2(__e, SymFn(symshen_4lazyderef), V5245, V5247)


let__13451 := tmp15120
_ = let__13451

tmp15160 := GEq(__e, let__13451, Nil)

if True == tmp15160 {
tmp15121 := Call2(__e, SymFn(symshen_4lazyderef), V5246, V5247)


let__13452 := tmp15121
_ = let__13452

tmp15122 := MakeNative1(func(__e *ControlFlow, Z5266 Obj) Obj {
tmp15123 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15123

return TailCall1(__e, SymFn(symthaw), V5250)


panic("unreachable")
})

let__13453 := tmp15122
_ = let__13453

tmp15158 := GConsP(__e, let__13452)

if True == tmp15158 {
tmp15124 := GHd(__e, let__13452)

tmp15125 := Call2(__e, SymFn(symshen_4lazyderef), tmp15124, V5247)


let__13454 := tmp15125
_ = let__13454

tmp15126 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15127 := GTl(__e, let__13452)

tmp15128 := Call2(__e, SymFn(symshen_4lazyderef), tmp15127, V5247)


let__13456 := tmp15128
_ = let__13456

tmp15129 := MakeNative1(func(__e *ControlFlow, Z5271 Obj) Obj {
return TailCall1(__e, let__13453, Z5271)
panic("unreachable")
})

let__13457 := tmp15129
_ = let__13457

tmp15145 := GConsP(__e, let__13456)

if True == tmp15145 {
tmp15130 := GHd(__e, let__13456)

let__13458 := tmp15130
_ = let__13458

tmp15131 := GTl(__e, let__13456)

tmp15132 := Call2(__e, SymFn(symshen_4lazyderef), tmp15131, V5247)


let__13459 := tmp15132
_ = let__13459

tmp15133 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13457, let__13458)
panic("unreachable")
})

let__13460 := tmp15133
_ = let__13460

tmp15137 := GEq(__e, let__13459, Nil)

if True == tmp15137 {
return TailCall1(__e, SymFn(symthaw), let__13460)
} else {
tmp15135 := Call1(__e, SymFn(symshen_4pvar_2), let__13459)


if True == tmp15135 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__13459, Nil, V5247, let__13460)
} else {
return False
}


}


} else {
tmp15143 := Call1(__e, SymFn(symshen_4pvar_2), let__13456)


if True == tmp15143 {
tmp15138 := Call1(__e, SymFn(symshen_4newpv), V5247)


let__13461 := tmp15138
_ = let__13461

tmp15139 := GCons(__e, let__13461, Nil)

tmp15140 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13457, let__13461)
panic("unreachable")
})

tmp15141 := Call4(__e, SymFn(symshen_4bind_b), let__13456, tmp15139, V5247, tmp15140)


return TailCall2(__e, SymFn(symshen_4gc), V5247, tmp15141)


} else {
return False
}


}


panic("unreachable")
})

let__13455 := tmp15126
_ = let__13455

tmp15149 := GEq(__e, let__13454, symlist)

if True == tmp15149 {
return TailCall1(__e, SymFn(symthaw), let__13455)
} else {
tmp15147 := Call1(__e, SymFn(symshen_4pvar_2), let__13454)


if True == tmp15147 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__13454, symlist, V5247, let__13455)
} else {
return False
}


}


} else {
tmp15156 := Call1(__e, SymFn(symshen_4pvar_2), let__13452)


if True == tmp15156 {
tmp15150 := Call1(__e, SymFn(symshen_4newpv), V5247)


let__13462 := tmp15150
_ = let__13462

tmp15151 := GCons(__e, let__13462, Nil)

tmp15152 := GCons(__e, symlist, tmp15151)

tmp15153 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13453, let__13462)
panic("unreachable")
})

tmp15154 := Call4(__e, SymFn(symshen_4bind_b), let__13452, tmp15152, V5247, tmp15153)


return TailCall2(__e, SymFn(symshen_4gc), V5247, tmp15154)


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
return let__13448
}


} else {
return let__13445
}


} else {
return let__13442
}


} else {
return let__13439
}


panic("unreachable")
})

tmp15171 := Call2(__e, ns2_1set, symshen_4primitive, tmp15067)


_ = tmp15171

tmp15172 := MakeNative(func(__e *ControlFlow) {
V5277 := __e.Get(1)
V5278 := __e.Get(2)
V5279 := __e.Get(3)
V5280 := __e.Get(4)
V5281 := __e.Get(5)
V5282 := __e.Get(6)
V5283 := __e.Get(7)
__e.ReturnDirect(func() Obj {
tmp15175 := Call2(__e, SymFn(symshen_4lazyderef), V5279, V5280)


tmp15176 := GEq(__e, tmp15175, Nil)

if True == tmp15176 {
return False
} else {
tmp15173 := Call2(__e, SymFn(symshen_4deref), V5277, V5280)


__e.TailApply(SymFn(symshen_4by_1hypothesis_1h), tmp15173, V5278, V5279, V5280, V5281, V5282, V5283)
return nil


}


panic("unreachable")
}())
}, 7)

tmp15177 := Call2(__e, ns2_1set, symshen_4by_1hypothesis, tmp15172)


_ = tmp15177

tmp15178 := MakeNative(func(__e *ControlFlow) {
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
tmp15211 := Call1(__e, SymFn(symshen_4unlocked_2), V5288)


var ifres15179 Obj

if True == tmp15211 {
tmp15180 := Call2(__e, SymFn(symshen_4lazyderef), V5286, V5287)


let__13464 := tmp15180
_ = let__13464

tmp15210 := GConsP(__e, let__13464)

var ifres15181 Obj

if True == tmp15210 {
tmp15182 := GHd(__e, let__13464)

tmp15183 := Call2(__e, SymFn(symshen_4lazyderef), tmp15182, V5287)


let__13465 := tmp15183
_ = let__13465

tmp15209 := GConsP(__e, let__13465)

var ifres15184 Obj

if True == tmp15209 {
tmp15185 := GHd(__e, let__13465)

let__13466 := tmp15185
_ = let__13466

tmp15186 := GTl(__e, let__13465)

tmp15187 := Call2(__e, SymFn(symshen_4lazyderef), tmp15186, V5287)


let__13467 := tmp15187
_ = let__13467

tmp15208 := GConsP(__e, let__13467)

var ifres15188 Obj

if True == tmp15208 {
tmp15189 := GHd(__e, let__13467)

let__13468 := tmp15189
_ = let__13468

tmp15190 := GTl(__e, let__13467)

tmp15191 := Call2(__e, SymFn(symshen_4lazyderef), tmp15190, V5287)


let__13469 := tmp15191
_ = let__13469

tmp15207 := GConsP(__e, let__13469)

var ifres15192 Obj

if True == tmp15207 {
tmp15193 := GHd(__e, let__13469)

let__13470 := tmp15193
_ = let__13470

tmp15194 := GTl(__e, let__13469)

tmp15195 := Call2(__e, SymFn(symshen_4lazyderef), tmp15194, V5287)


let__13471 := tmp15195
_ = let__13471

tmp15206 := GEq(__e, let__13471, Nil)

var ifres15196 Obj

if True == tmp15206 {
tmp15197 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15197

tmp15198 := Call2(__e, SymFn(symshen_4deref), let__13468, V5287)


tmp15199 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp15200 := GEq(__e, tmp15198, tmp15199)

tmp15201 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15202 := Call2(__e, SymFn(symshen_4deref), let__13466, V5287)


tmp15203 := GEq(__e, V5284, tmp15202)

tmp15204 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall6(__e, SymFn(symis_b), V5285, let__13470, V5287, V5288, V5289, V5290)
panic("unreachable")
})

return TailCall5(__e, SymFn(symwhen), tmp15203, V5287, V5288, V5289, tmp15204)


panic("unreachable")
})

tmp15205 := Call5(__e, SymFn(symwhen), tmp15200, V5287, V5288, V5289, tmp15201)


ifres15196 = tmp15205


} else {
ifres15196 = False


}

ifres15192 = ifres15196


} else {
ifres15192 = False


}

ifres15188 = ifres15192


} else {
ifres15188 = False


}

ifres15184 = ifres15188


} else {
ifres15184 = False


}

ifres15181 = ifres15184


} else {
ifres15181 = False


}

ifres15179 = ifres15181


} else {
ifres15179 = False


}

let__13463 := ifres15179
_ = let__13463

tmp15220 := GEq(__e, let__13463, False)

if True == tmp15220 {
tmp15218 := Call1(__e, SymFn(symshen_4unlocked_2), V5288)


if True == tmp15218 {
tmp15212 := Call2(__e, SymFn(symshen_4lazyderef), V5286, V5287)


let__13472 := tmp15212
_ = let__13472

tmp15216 := GConsP(__e, let__13472)

if True == tmp15216 {
tmp15213 := GTl(__e, let__13472)

let__13473 := tmp15213
_ = let__13473

tmp15214 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15214

if SymFn(symshen_4by_1hypothesis_1h) == __self {
__self1, __self2, __self3, __self4, __self5, __self6, __self7 = V5284, V5285, let__13473, V5287, V5288, V5289, V5290
__e.Tick()
goto __selftop
}
__e.TailApply(SymFn(symshen_4by_1hypothesis_1h), V5284, V5285, let__13473, V5287, V5288, V5289, V5290)
return nil


} else {
return False
}


} else {
return False
}


} else {
return let__13463
}


panic("unreachable")
}())
}, 7)

tmp15221 := Call2(__e, ns2_1set, symshen_4by_1hypothesis_1h, tmp15178)


_ = tmp15221

tmp15222 := MakeNative6(func(__e *ControlFlow, V5302, V5303, V5304, V5305, V5306, V5307 Obj) Obj {
tmp15227 := Call1(__e, SymFn(symshen_4unlocked_2), V5305)


if True == tmp15227 {
tmp15223 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15223

tmp15224 := GValue(__e, symshen_4_dsigf_d)

tmp15225 := Call2(__e, SymFn(symassoc), V5302, tmp15224)


return TailCall6(__e, SymFn(symshen_4sigf), tmp15225, V5303, V5304, V5305, V5306, V5307)


} else {
return False
}


panic("unreachable")
})

tmp15228 := Call2(__e, ns2_1set, symshen_4lookupsig, tmp15222)


_ = tmp15228

tmp15229 := MakeNative6(func(__e *ControlFlow, V5322, V5323, V5324, V5325, V5326, V5327 Obj) Obj {
tmp15236 := GConsP(__e, V5322)

if True == tmp15236 {
tmp15230 := GTl(__e, V5322)

tmp15231 := Call1(__e, tmp15230, V5323)


tmp15232 := Call1(__e, tmp15231, V5324)


tmp15233 := Call1(__e, tmp15232, V5325)


tmp15234 := Call1(__e, tmp15233, V5326)


return TailCall1(__e, tmp15234, V5327)


} else {
return False
}


panic("unreachable")
})

tmp15237 := Call2(__e, ns2_1set, symshen_4sigf, tmp15229)


_ = tmp15237

tmp15238 := MakeNative1(func(__e *ControlFlow, V5328 Obj) Obj {
tmp15239 := PrimGuard1(__e, symabsvector, PrimAbsvector, MakeInteger(3))

let__13474 := tmp15239
_ = let__13474

tmp15240 := GVecSet(__e, let__13474, MakeInteger(0), symshen_4print_1freshterm)

let__13475 := tmp15240
_ = let__13475

tmp15241 := GVecSet(__e, let__13475, MakeInteger(1), V5328)

let__13476 := tmp15241
_ = let__13476

tmp15242 := GValue(__e, symshen_4_dgensym_d)

tmp15244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dgensym_d, GAdd(__e, MakeInteger(1), tmp15242))
}
__typedArg0 := symshen_4_dgensym_d
__typedArg1 := GAdd(__e, MakeInteger(1), tmp15242)
return Call2(__e, SymFn(symset), __typedArg0, __typedArg1)
})()

tmp15245 := GVecSet(__e, let__13476, MakeInteger(2), tmp15244)

let__13477 := tmp15245
_ = let__13477

return let__13477


panic("unreachable")
})

tmp15246 := Call2(__e, ns2_1set, symshen_4freshterm, tmp15238)


_ = tmp15246

tmp15247 := MakeNative1(func(__e *ControlFlow, V5333 Obj) Obj {
tmp15248 := GVecGet(__e, V5333, MakeInteger(1))

tmp15249 := Call3(__e, SymFn(symshen_4app), tmp15248, MakeString(""), symshen_4a)


return PrimGuard2(__e, symcn, PrimStringConcat, MakeString("&&"), tmp15249)


panic("unreachable")
})

tmp15250 := Call2(__e, ns2_1set, symshen_4print_1freshterm, tmp15247)


_ = tmp15250

tmp15251 := MakeNative(func(__e *ControlFlow) {
V5334 := __e.Get(1)
V5335 := __e.Get(2)
V5336 := __e.Get(3)
V5337 := __e.Get(4)
V5338 := __e.Get(5)
V5339 := __e.Get(6)
V5340 := __e.Get(7)
__e.ReturnDirect(func() Obj {
tmp15255 := Call2(__e, SymFn(symshen_4lazyderef), V5336, V5337)


tmp15256 := GEq(__e, tmp15255, Nil)

if True == tmp15256 {
return False
} else {
tmp15252 := Call2(__e, SymFn(symshen_4deref), V5334, V5337)


tmp15253 := Call2(__e, SymFn(symshen_4deref), V5335, V5337)


__e.TailApply(SymFn(symshen_4search_1user_1datatypes_1h), tmp15252, tmp15253, V5336, V5337, V5338, V5339, V5340)
return nil


}


panic("unreachable")
}())
}, 7)

tmp15257 := Call2(__e, ns2_1set, symshen_4search_1user_1datatypes, tmp15251)


_ = tmp15257

tmp15258 := MakeNative(func(__e *ControlFlow) {
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
tmp15273 := Call1(__e, SymFn(symshen_4unlocked_2), V5345)


var ifres15259 Obj

if True == tmp15273 {
tmp15260 := Call2(__e, SymFn(symshen_4lazyderef), V5343, V5344)


let__13479 := tmp15260
_ = let__13479

tmp15272 := GConsP(__e, let__13479)

var ifres15261 Obj

if True == tmp15272 {
tmp15262 := GHd(__e, let__13479)

tmp15263 := Call2(__e, SymFn(symshen_4lazyderef), tmp15262, V5344)


let__13480 := tmp15263
_ = let__13480

tmp15271 := GConsP(__e, let__13480)

var ifres15264 Obj

if True == tmp15271 {
tmp15265 := GTl(__e, let__13480)

let__13481 := tmp15265
_ = let__13481

tmp15266 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15266

tmp15267 := Call2(__e, SymFn(symshen_4deref), let__13481, V5344)


tmp15268 := Call1(__e, tmp15267, V5341)


tmp15269 := Call1(__e, tmp15268, V5342)


tmp15270 := Call5(__e, SymFn(symcall), tmp15269, V5344, V5345, V5346, V5347)


ifres15264 = tmp15270


} else {
ifres15264 = False


}

ifres15261 = ifres15264


} else {
ifres15261 = False


}

ifres15259 = ifres15261


} else {
ifres15259 = False


}

let__13478 := ifres15259
_ = let__13478

tmp15282 := GEq(__e, let__13478, False)

if True == tmp15282 {
tmp15280 := Call1(__e, SymFn(symshen_4unlocked_2), V5345)


if True == tmp15280 {
tmp15274 := Call2(__e, SymFn(symshen_4lazyderef), V5343, V5344)


let__13482 := tmp15274
_ = let__13482

tmp15278 := GConsP(__e, let__13482)

if True == tmp15278 {
tmp15275 := GTl(__e, let__13482)

let__13483 := tmp15275
_ = let__13483

tmp15276 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15276

if SymFn(symshen_4search_1user_1datatypes_1h) == __self {
__self1, __self2, __self3, __self4, __self5, __self6, __self7 = V5341, V5342, let__13483, V5344, V5345, V5346, V5347
__e.Tick()
goto __selftop
}
__e.TailApply(SymFn(symshen_4search_1user_1datatypes_1h), V5341, V5342, let__13483, V5344, V5345, V5346, V5347)
return nil


} else {
return False
}


} else {
return False
}


} else {
return let__13478
}


panic("unreachable")
}())
}, 7)

tmp15283 := Call2(__e, ns2_1set, symshen_4search_1user_1datatypes_1h, tmp15258)


_ = tmp15283

tmp15284 := MakeNative(func(__e *ControlFlow) {
V5354 := __e.Get(1)
V5355 := __e.Get(2)
V5356 := __e.Get(3)
V5357 := __e.Get(4)
V5358 := __e.Get(5)
V5359 := __e.Get(6)
V5360 := __e.Get(7)
__e.ReturnDirect(func() Obj {
let__13484 := GAdd(__e, V5359, MakeInteger(1))
_ = let__13484

tmp15296 := Call1(__e, SymFn(symshen_4unlocked_2), V5358)


var ifres15286 Obj

if True == tmp15296 {
tmp15287 := Call2(__e, SymFn(symshen_4lazyderef), V5354, V5357)


let__13486 := tmp15287
_ = let__13486

tmp15295 := GEq(__e, let__13486, Nil)

var ifres15288 Obj

if True == tmp15295 {
tmp15289 := Call2(__e, SymFn(symshen_4lazyderef), V5356, V5357)


let__13487 := tmp15289
_ = let__13487

tmp15294 := GEq(__e, let__13487, True)

var ifres15290 Obj

if True == tmp15294 {
tmp15291 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15291

tmp15292 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall6(__e, SymFn(symbind), V5355, Nil, V5357, V5358, let__13484, V5360)
panic("unreachable")
})

tmp15293 := Call4(__e, SymFn(symshen_4cut), V5357, V5358, let__13484, tmp15292)


ifres15290 = tmp15293


} else {
ifres15290 = False


}

ifres15288 = ifres15290


} else {
ifres15288 = False


}

ifres15286 = ifres15288


} else {
ifres15286 = False


}

let__13485 := ifres15286
_ = let__13485

tmp15639 := GEq(__e, let__13485, False)

if True == tmp15639 {
tmp15375 := Call1(__e, SymFn(symshen_4unlocked_2), V5358)


var ifres15297 Obj

if True == tmp15375 {
tmp15298 := Call2(__e, SymFn(symshen_4lazyderef), V5354, V5357)


let__13489 := tmp15298
_ = let__13489

tmp15374 := GConsP(__e, let__13489)

var ifres15299 Obj

if True == tmp15374 {
tmp15300 := GHd(__e, let__13489)

tmp15301 := Call2(__e, SymFn(symshen_4lazyderef), tmp15300, V5357)


let__13490 := tmp15301
_ = let__13490

tmp15373 := GConsP(__e, let__13490)

var ifres15302 Obj

if True == tmp15373 {
tmp15303 := GHd(__e, let__13490)

tmp15304 := Call2(__e, SymFn(symshen_4lazyderef), tmp15303, V5357)


let__13491 := tmp15304
_ = let__13491

tmp15372 := GConsP(__e, let__13491)

var ifres15305 Obj

if True == tmp15372 {
tmp15306 := GHd(__e, let__13491)

tmp15307 := Call2(__e, SymFn(symshen_4lazyderef), tmp15306, V5357)


let__13492 := tmp15307
_ = let__13492

tmp15371 := GEq(__e, let__13492, symcons)

var ifres15308 Obj

if True == tmp15371 {
tmp15309 := GTl(__e, let__13491)

tmp15310 := Call2(__e, SymFn(symshen_4lazyderef), tmp15309, V5357)


let__13493 := tmp15310
_ = let__13493

tmp15370 := GConsP(__e, let__13493)

var ifres15311 Obj

if True == tmp15370 {
tmp15312 := GHd(__e, let__13493)

let__13494 := tmp15312
_ = let__13494

tmp15313 := GTl(__e, let__13493)

tmp15314 := Call2(__e, SymFn(symshen_4lazyderef), tmp15313, V5357)


let__13495 := tmp15314
_ = let__13495

tmp15369 := GConsP(__e, let__13495)

var ifres15315 Obj

if True == tmp15369 {
tmp15316 := GHd(__e, let__13495)

let__13496 := tmp15316
_ = let__13496

tmp15317 := GTl(__e, let__13495)

tmp15318 := Call2(__e, SymFn(symshen_4lazyderef), tmp15317, V5357)


let__13497 := tmp15318
_ = let__13497

tmp15368 := GEq(__e, let__13497, Nil)

var ifres15319 Obj

if True == tmp15368 {
tmp15320 := GTl(__e, let__13490)

tmp15321 := Call2(__e, SymFn(symshen_4lazyderef), tmp15320, V5357)


let__13498 := tmp15321
_ = let__13498

tmp15367 := GConsP(__e, let__13498)

var ifres15322 Obj

if True == tmp15367 {
tmp15323 := GHd(__e, let__13498)

let__13499 := tmp15323
_ = let__13499

tmp15324 := GTl(__e, let__13498)

tmp15325 := Call2(__e, SymFn(symshen_4lazyderef), tmp15324, V5357)


let__13500 := tmp15325
_ = let__13500

tmp15366 := GConsP(__e, let__13500)

var ifres15326 Obj

if True == tmp15366 {
tmp15327 := GHd(__e, let__13500)

tmp15328 := Call2(__e, SymFn(symshen_4lazyderef), tmp15327, V5357)


let__13501 := tmp15328
_ = let__13501

tmp15365 := GConsP(__e, let__13501)

var ifres15329 Obj

if True == tmp15365 {
tmp15330 := GHd(__e, let__13501)

tmp15331 := Call2(__e, SymFn(symshen_4lazyderef), tmp15330, V5357)


let__13502 := tmp15331
_ = let__13502

tmp15364 := GEq(__e, let__13502, symlist)

var ifres15332 Obj

if True == tmp15364 {
tmp15333 := GTl(__e, let__13501)

tmp15334 := Call2(__e, SymFn(symshen_4lazyderef), tmp15333, V5357)


let__13503 := tmp15334
_ = let__13503

tmp15363 := GConsP(__e, let__13503)

var ifres15335 Obj

if True == tmp15363 {
tmp15336 := GHd(__e, let__13503)

let__13504 := tmp15336
_ = let__13504

tmp15337 := GTl(__e, let__13503)

tmp15338 := Call2(__e, SymFn(symshen_4lazyderef), tmp15337, V5357)


let__13505 := tmp15338
_ = let__13505

tmp15362 := GEq(__e, let__13505, Nil)

var ifres15339 Obj

if True == tmp15362 {
tmp15340 := GTl(__e, let__13500)

tmp15341 := Call2(__e, SymFn(symshen_4lazyderef), tmp15340, V5357)


let__13506 := tmp15341
_ = let__13506

tmp15361 := GEq(__e, let__13506, Nil)

var ifres15342 Obj

if True == tmp15361 {
tmp15343 := GTl(__e, let__13489)

let__13507 := tmp15343
_ = let__13507

tmp15344 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15344

tmp15345 := Call2(__e, SymFn(symshen_4deref), let__13499, V5357)


tmp15346 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp15347 := GEq(__e, tmp15345, tmp15346)

tmp15348 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15349 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15350 := GCons(__e, let__13504, Nil)

tmp15351 := GCons(__e, let__13499, tmp15350)

tmp15352 := GCons(__e, let__13494, tmp15351)

tmp15353 := GCons(__e, let__13504, Nil)

tmp15354 := GCons(__e, symlist, tmp15353)

tmp15355 := GCons(__e, tmp15354, Nil)

tmp15356 := GCons(__e, let__13499, tmp15355)

tmp15357 := GCons(__e, let__13496, tmp15356)

tmp15358 := GCons(__e, tmp15357, let__13507)

tmp15359 := GCons(__e, tmp15352, tmp15358)

__e.TailApply(SymFn(symshen_4l_1rules), tmp15359, V5355, True, V5357, V5358, let__13484, V5360)
return nil


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5357, V5358, let__13484, tmp15349)


panic("unreachable")
})

tmp15360 := Call5(__e, SymFn(symwhen), tmp15347, V5357, V5358, let__13484, tmp15348)


ifres15342 = tmp15360


} else {
ifres15342 = False


}

ifres15339 = ifres15342


} else {
ifres15339 = False


}

ifres15335 = ifres15339


} else {
ifres15335 = False


}

ifres15332 = ifres15335


} else {
ifres15332 = False


}

ifres15329 = ifres15332


} else {
ifres15329 = False


}

ifres15326 = ifres15329


} else {
ifres15326 = False


}

ifres15322 = ifres15326


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

ifres15311 = ifres15315


} else {
ifres15311 = False


}

ifres15308 = ifres15311


} else {
ifres15308 = False


}

ifres15305 = ifres15308


} else {
ifres15305 = False


}

ifres15302 = ifres15305


} else {
ifres15302 = False


}

ifres15299 = ifres15302


} else {
ifres15299 = False


}

ifres15297 = ifres15299


} else {
ifres15297 = False


}

let__13488 := ifres15297
_ = let__13488

tmp15637 := GEq(__e, let__13488, False)

if True == tmp15637 {
tmp15457 := Call1(__e, SymFn(symshen_4unlocked_2), V5358)


var ifres15376 Obj

if True == tmp15457 {
tmp15377 := Call2(__e, SymFn(symshen_4lazyderef), V5354, V5357)


let__13509 := tmp15377
_ = let__13509

tmp15456 := GConsP(__e, let__13509)

var ifres15378 Obj

if True == tmp15456 {
tmp15379 := GHd(__e, let__13509)

tmp15380 := Call2(__e, SymFn(symshen_4lazyderef), tmp15379, V5357)


let__13510 := tmp15380
_ = let__13510

tmp15455 := GConsP(__e, let__13510)

var ifres15381 Obj

if True == tmp15455 {
tmp15382 := GHd(__e, let__13510)

tmp15383 := Call2(__e, SymFn(symshen_4lazyderef), tmp15382, V5357)


let__13511 := tmp15383
_ = let__13511

tmp15454 := GConsP(__e, let__13511)

var ifres15384 Obj

if True == tmp15454 {
tmp15385 := GHd(__e, let__13511)

tmp15386 := Call2(__e, SymFn(symshen_4lazyderef), tmp15385, V5357)


let__13512 := tmp15386
_ = let__13512

tmp15453 := GEq(__e, let__13512, sym_8p)

var ifres15387 Obj

if True == tmp15453 {
tmp15388 := GTl(__e, let__13511)

tmp15389 := Call2(__e, SymFn(symshen_4lazyderef), tmp15388, V5357)


let__13513 := tmp15389
_ = let__13513

tmp15452 := GConsP(__e, let__13513)

var ifres15390 Obj

if True == tmp15452 {
tmp15391 := GHd(__e, let__13513)

let__13514 := tmp15391
_ = let__13514

tmp15392 := GTl(__e, let__13513)

tmp15393 := Call2(__e, SymFn(symshen_4lazyderef), tmp15392, V5357)


let__13515 := tmp15393
_ = let__13515

tmp15451 := GConsP(__e, let__13515)

var ifres15394 Obj

if True == tmp15451 {
tmp15395 := GHd(__e, let__13515)

let__13516 := tmp15395
_ = let__13516

tmp15396 := GTl(__e, let__13515)

tmp15397 := Call2(__e, SymFn(symshen_4lazyderef), tmp15396, V5357)


let__13517 := tmp15397
_ = let__13517

tmp15450 := GEq(__e, let__13517, Nil)

var ifres15398 Obj

if True == tmp15450 {
tmp15399 := GTl(__e, let__13510)

tmp15400 := Call2(__e, SymFn(symshen_4lazyderef), tmp15399, V5357)


let__13518 := tmp15400
_ = let__13518

tmp15449 := GConsP(__e, let__13518)

var ifres15401 Obj

if True == tmp15449 {
tmp15402 := GHd(__e, let__13518)

let__13519 := tmp15402
_ = let__13519

tmp15403 := GTl(__e, let__13518)

tmp15404 := Call2(__e, SymFn(symshen_4lazyderef), tmp15403, V5357)


let__13520 := tmp15404
_ = let__13520

tmp15448 := GConsP(__e, let__13520)

var ifres15405 Obj

if True == tmp15448 {
tmp15406 := GHd(__e, let__13520)

tmp15407 := Call2(__e, SymFn(symshen_4lazyderef), tmp15406, V5357)


let__13521 := tmp15407
_ = let__13521

tmp15447 := GConsP(__e, let__13521)

var ifres15408 Obj

if True == tmp15447 {
tmp15409 := GHd(__e, let__13521)

let__13522 := tmp15409
_ = let__13522

tmp15410 := GTl(__e, let__13521)

tmp15411 := Call2(__e, SymFn(symshen_4lazyderef), tmp15410, V5357)


let__13523 := tmp15411
_ = let__13523

tmp15446 := GConsP(__e, let__13523)

var ifres15412 Obj

if True == tmp15446 {
tmp15413 := GHd(__e, let__13523)

tmp15414 := Call2(__e, SymFn(symshen_4lazyderef), tmp15413, V5357)


let__13524 := tmp15414
_ = let__13524

tmp15445 := GEq(__e, let__13524, sym_d)

var ifres15415 Obj

if True == tmp15445 {
tmp15416 := GTl(__e, let__13523)

tmp15417 := Call2(__e, SymFn(symshen_4lazyderef), tmp15416, V5357)


let__13525 := tmp15417
_ = let__13525

tmp15444 := GConsP(__e, let__13525)

var ifres15418 Obj

if True == tmp15444 {
tmp15419 := GHd(__e, let__13525)

let__13526 := tmp15419
_ = let__13526

tmp15420 := GTl(__e, let__13525)

tmp15421 := Call2(__e, SymFn(symshen_4lazyderef), tmp15420, V5357)


let__13527 := tmp15421
_ = let__13527

tmp15443 := GEq(__e, let__13527, Nil)

var ifres15422 Obj

if True == tmp15443 {
tmp15423 := GTl(__e, let__13520)

tmp15424 := Call2(__e, SymFn(symshen_4lazyderef), tmp15423, V5357)


let__13528 := tmp15424
_ = let__13528

tmp15442 := GEq(__e, let__13528, Nil)

var ifres15425 Obj

if True == tmp15442 {
tmp15426 := GTl(__e, let__13509)

let__13529 := tmp15426
_ = let__13529

tmp15427 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15427

tmp15428 := Call2(__e, SymFn(symshen_4deref), let__13519, V5357)


tmp15429 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp15430 := GEq(__e, tmp15428, tmp15429)

tmp15431 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15432 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15433 := GCons(__e, let__13522, Nil)

tmp15434 := GCons(__e, let__13519, tmp15433)

tmp15435 := GCons(__e, let__13514, tmp15434)

tmp15436 := GCons(__e, let__13526, Nil)

tmp15437 := GCons(__e, let__13519, tmp15436)

tmp15438 := GCons(__e, let__13516, tmp15437)

tmp15439 := GCons(__e, tmp15438, let__13529)

tmp15440 := GCons(__e, tmp15435, tmp15439)

__e.TailApply(SymFn(symshen_4l_1rules), tmp15440, V5355, True, V5357, V5358, let__13484, V5360)
return nil


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5357, V5358, let__13484, tmp15432)


panic("unreachable")
})

tmp15441 := Call5(__e, SymFn(symwhen), tmp15430, V5357, V5358, let__13484, tmp15431)


ifres15425 = tmp15441


} else {
ifres15425 = False


}

ifres15422 = ifres15425


} else {
ifres15422 = False


}

ifres15418 = ifres15422


} else {
ifres15418 = False


}

ifres15415 = ifres15418


} else {
ifres15415 = False


}

ifres15412 = ifres15415


} else {
ifres15412 = False


}

ifres15408 = ifres15412


} else {
ifres15408 = False


}

ifres15405 = ifres15408


} else {
ifres15405 = False


}

ifres15401 = ifres15405


} else {
ifres15401 = False


}

ifres15398 = ifres15401


} else {
ifres15398 = False


}

ifres15394 = ifres15398


} else {
ifres15394 = False


}

ifres15390 = ifres15394


} else {
ifres15390 = False


}

ifres15387 = ifres15390


} else {
ifres15387 = False


}

ifres15384 = ifres15387


} else {
ifres15384 = False


}

ifres15381 = ifres15384


} else {
ifres15381 = False


}

ifres15378 = ifres15381


} else {
ifres15378 = False


}

ifres15376 = ifres15378


} else {
ifres15376 = False


}

let__13508 := ifres15376
_ = let__13508

tmp15635 := GEq(__e, let__13508, False)

if True == tmp15635 {
tmp15521 := Call1(__e, SymFn(symshen_4unlocked_2), V5358)


var ifres15458 Obj

if True == tmp15521 {
tmp15459 := Call2(__e, SymFn(symshen_4lazyderef), V5354, V5357)


let__13531 := tmp15459
_ = let__13531

tmp15520 := GConsP(__e, let__13531)

var ifres15460 Obj

if True == tmp15520 {
tmp15461 := GHd(__e, let__13531)

tmp15462 := Call2(__e, SymFn(symshen_4lazyderef), tmp15461, V5357)


let__13532 := tmp15462
_ = let__13532

tmp15519 := GConsP(__e, let__13532)

var ifres15463 Obj

if True == tmp15519 {
tmp15464 := GHd(__e, let__13532)

tmp15465 := Call2(__e, SymFn(symshen_4lazyderef), tmp15464, V5357)


let__13533 := tmp15465
_ = let__13533

tmp15518 := GConsP(__e, let__13533)

var ifres15466 Obj

if True == tmp15518 {
tmp15467 := GHd(__e, let__13533)

tmp15468 := Call2(__e, SymFn(symshen_4lazyderef), tmp15467, V5357)


let__13534 := tmp15468
_ = let__13534

tmp15517 := GEq(__e, let__13534, sym_8s)

var ifres15469 Obj

if True == tmp15517 {
tmp15470 := GTl(__e, let__13533)

tmp15471 := Call2(__e, SymFn(symshen_4lazyderef), tmp15470, V5357)


let__13535 := tmp15471
_ = let__13535

tmp15516 := GConsP(__e, let__13535)

var ifres15472 Obj

if True == tmp15516 {
tmp15473 := GHd(__e, let__13535)

let__13536 := tmp15473
_ = let__13536

tmp15474 := GTl(__e, let__13535)

tmp15475 := Call2(__e, SymFn(symshen_4lazyderef), tmp15474, V5357)


let__13537 := tmp15475
_ = let__13537

tmp15515 := GConsP(__e, let__13537)

var ifres15476 Obj

if True == tmp15515 {
tmp15477 := GHd(__e, let__13537)

let__13538 := tmp15477
_ = let__13538

tmp15478 := GTl(__e, let__13537)

tmp15479 := Call2(__e, SymFn(symshen_4lazyderef), tmp15478, V5357)


let__13539 := tmp15479
_ = let__13539

tmp15514 := GEq(__e, let__13539, Nil)

var ifres15480 Obj

if True == tmp15514 {
tmp15481 := GTl(__e, let__13532)

tmp15482 := Call2(__e, SymFn(symshen_4lazyderef), tmp15481, V5357)


let__13540 := tmp15482
_ = let__13540

tmp15513 := GConsP(__e, let__13540)

var ifres15483 Obj

if True == tmp15513 {
tmp15484 := GHd(__e, let__13540)

let__13541 := tmp15484
_ = let__13541

tmp15485 := GTl(__e, let__13540)

tmp15486 := Call2(__e, SymFn(symshen_4lazyderef), tmp15485, V5357)


let__13542 := tmp15486
_ = let__13542

tmp15512 := GConsP(__e, let__13542)

var ifres15487 Obj

if True == tmp15512 {
tmp15488 := GHd(__e, let__13542)

tmp15489 := Call2(__e, SymFn(symshen_4lazyderef), tmp15488, V5357)


let__13543 := tmp15489
_ = let__13543

tmp15511 := GEq(__e, let__13543, symstring)

var ifres15490 Obj

if True == tmp15511 {
tmp15491 := GTl(__e, let__13542)

tmp15492 := Call2(__e, SymFn(symshen_4lazyderef), tmp15491, V5357)


let__13544 := tmp15492
_ = let__13544

tmp15510 := GEq(__e, let__13544, Nil)

var ifres15493 Obj

if True == tmp15510 {
tmp15494 := GTl(__e, let__13531)

let__13545 := tmp15494
_ = let__13545

tmp15495 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15495

tmp15496 := Call2(__e, SymFn(symshen_4deref), let__13541, V5357)


tmp15497 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp15498 := GEq(__e, tmp15496, tmp15497)

tmp15499 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15500 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15501 := GCons(__e, symstring, Nil)

tmp15502 := GCons(__e, let__13541, tmp15501)

tmp15503 := GCons(__e, let__13536, tmp15502)

tmp15504 := GCons(__e, symstring, Nil)

tmp15505 := GCons(__e, let__13541, tmp15504)

tmp15506 := GCons(__e, let__13538, tmp15505)

tmp15507 := GCons(__e, tmp15506, let__13545)

tmp15508 := GCons(__e, tmp15503, tmp15507)

__e.TailApply(SymFn(symshen_4l_1rules), tmp15508, V5355, True, V5357, V5358, let__13484, V5360)
return nil


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5357, V5358, let__13484, tmp15500)


panic("unreachable")
})

tmp15509 := Call5(__e, SymFn(symwhen), tmp15498, V5357, V5358, let__13484, tmp15499)


ifres15493 = tmp15509


} else {
ifres15493 = False


}

ifres15490 = ifres15493


} else {
ifres15490 = False


}

ifres15487 = ifres15490


} else {
ifres15487 = False


}

ifres15483 = ifres15487


} else {
ifres15483 = False


}

ifres15480 = ifres15483


} else {
ifres15480 = False


}

ifres15476 = ifres15480


} else {
ifres15476 = False


}

ifres15472 = ifres15476


} else {
ifres15472 = False


}

ifres15469 = ifres15472


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

ifres15460 = ifres15463


} else {
ifres15460 = False


}

ifres15458 = ifres15460


} else {
ifres15458 = False


}

let__13530 := ifres15458
_ = let__13530

tmp15633 := GEq(__e, let__13530, False)

if True == tmp15633 {
tmp15600 := Call1(__e, SymFn(symshen_4unlocked_2), V5358)


var ifres15522 Obj

if True == tmp15600 {
tmp15523 := Call2(__e, SymFn(symshen_4lazyderef), V5354, V5357)


let__13547 := tmp15523
_ = let__13547

tmp15599 := GConsP(__e, let__13547)

var ifres15524 Obj

if True == tmp15599 {
tmp15525 := GHd(__e, let__13547)

tmp15526 := Call2(__e, SymFn(symshen_4lazyderef), tmp15525, V5357)


let__13548 := tmp15526
_ = let__13548

tmp15598 := GConsP(__e, let__13548)

var ifres15527 Obj

if True == tmp15598 {
tmp15528 := GHd(__e, let__13548)

tmp15529 := Call2(__e, SymFn(symshen_4lazyderef), tmp15528, V5357)


let__13549 := tmp15529
_ = let__13549

tmp15597 := GConsP(__e, let__13549)

var ifres15530 Obj

if True == tmp15597 {
tmp15531 := GHd(__e, let__13549)

tmp15532 := Call2(__e, SymFn(symshen_4lazyderef), tmp15531, V5357)


let__13550 := tmp15532
_ = let__13550

tmp15596 := GEq(__e, let__13550, sym_8v)

var ifres15533 Obj

if True == tmp15596 {
tmp15534 := GTl(__e, let__13549)

tmp15535 := Call2(__e, SymFn(symshen_4lazyderef), tmp15534, V5357)


let__13551 := tmp15535
_ = let__13551

tmp15595 := GConsP(__e, let__13551)

var ifres15536 Obj

if True == tmp15595 {
tmp15537 := GHd(__e, let__13551)

let__13552 := tmp15537
_ = let__13552

tmp15538 := GTl(__e, let__13551)

tmp15539 := Call2(__e, SymFn(symshen_4lazyderef), tmp15538, V5357)


let__13553 := tmp15539
_ = let__13553

tmp15594 := GConsP(__e, let__13553)

var ifres15540 Obj

if True == tmp15594 {
tmp15541 := GHd(__e, let__13553)

let__13554 := tmp15541
_ = let__13554

tmp15542 := GTl(__e, let__13553)

tmp15543 := Call2(__e, SymFn(symshen_4lazyderef), tmp15542, V5357)


let__13555 := tmp15543
_ = let__13555

tmp15593 := GEq(__e, let__13555, Nil)

var ifres15544 Obj

if True == tmp15593 {
tmp15545 := GTl(__e, let__13548)

tmp15546 := Call2(__e, SymFn(symshen_4lazyderef), tmp15545, V5357)


let__13556 := tmp15546
_ = let__13556

tmp15592 := GConsP(__e, let__13556)

var ifres15547 Obj

if True == tmp15592 {
tmp15548 := GHd(__e, let__13556)

let__13557 := tmp15548
_ = let__13557

tmp15549 := GTl(__e, let__13556)

tmp15550 := Call2(__e, SymFn(symshen_4lazyderef), tmp15549, V5357)


let__13558 := tmp15550
_ = let__13558

tmp15591 := GConsP(__e, let__13558)

var ifres15551 Obj

if True == tmp15591 {
tmp15552 := GHd(__e, let__13558)

tmp15553 := Call2(__e, SymFn(symshen_4lazyderef), tmp15552, V5357)


let__13559 := tmp15553
_ = let__13559

tmp15590 := GConsP(__e, let__13559)

var ifres15554 Obj

if True == tmp15590 {
tmp15555 := GHd(__e, let__13559)

tmp15556 := Call2(__e, SymFn(symshen_4lazyderef), tmp15555, V5357)


let__13560 := tmp15556
_ = let__13560

tmp15589 := GEq(__e, let__13560, symvector)

var ifres15557 Obj

if True == tmp15589 {
tmp15558 := GTl(__e, let__13559)

tmp15559 := Call2(__e, SymFn(symshen_4lazyderef), tmp15558, V5357)


let__13561 := tmp15559
_ = let__13561

tmp15588 := GConsP(__e, let__13561)

var ifres15560 Obj

if True == tmp15588 {
tmp15561 := GHd(__e, let__13561)

let__13562 := tmp15561
_ = let__13562

tmp15562 := GTl(__e, let__13561)

tmp15563 := Call2(__e, SymFn(symshen_4lazyderef), tmp15562, V5357)


let__13563 := tmp15563
_ = let__13563

tmp15587 := GEq(__e, let__13563, Nil)

var ifres15564 Obj

if True == tmp15587 {
tmp15565 := GTl(__e, let__13558)

tmp15566 := Call2(__e, SymFn(symshen_4lazyderef), tmp15565, V5357)


let__13564 := tmp15566
_ = let__13564

tmp15586 := GEq(__e, let__13564, Nil)

var ifres15567 Obj

if True == tmp15586 {
tmp15568 := GTl(__e, let__13547)

let__13565 := tmp15568
_ = let__13565

tmp15569 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15569

tmp15570 := Call2(__e, SymFn(symshen_4deref), let__13557, V5357)


tmp15571 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp15572 := GEq(__e, tmp15570, tmp15571)

tmp15573 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15574 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15575 := GCons(__e, let__13562, Nil)

tmp15576 := GCons(__e, let__13557, tmp15575)

tmp15577 := GCons(__e, let__13552, tmp15576)

tmp15578 := GCons(__e, let__13562, Nil)

tmp15579 := GCons(__e, symvector, tmp15578)

tmp15580 := GCons(__e, tmp15579, Nil)

tmp15581 := GCons(__e, let__13557, tmp15580)

tmp15582 := GCons(__e, let__13554, tmp15581)

tmp15583 := GCons(__e, tmp15582, let__13565)

tmp15584 := GCons(__e, tmp15577, tmp15583)

__e.TailApply(SymFn(symshen_4l_1rules), tmp15584, V5355, True, V5357, V5358, let__13484, V5360)
return nil


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5357, V5358, let__13484, tmp15574)


panic("unreachable")
})

tmp15585 := Call5(__e, SymFn(symwhen), tmp15572, V5357, V5358, let__13484, tmp15573)


ifres15567 = tmp15585


} else {
ifres15567 = False


}

ifres15564 = ifres15567


} else {
ifres15564 = False


}

ifres15560 = ifres15564


} else {
ifres15560 = False


}

ifres15557 = ifres15560


} else {
ifres15557 = False


}

ifres15554 = ifres15557


} else {
ifres15554 = False


}

ifres15551 = ifres15554


} else {
ifres15551 = False


}

ifres15547 = ifres15551


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

ifres15536 = ifres15540


} else {
ifres15536 = False


}

ifres15533 = ifres15536


} else {
ifres15533 = False


}

ifres15530 = ifres15533


} else {
ifres15530 = False


}

ifres15527 = ifres15530


} else {
ifres15527 = False


}

ifres15524 = ifres15527


} else {
ifres15524 = False


}

ifres15522 = ifres15524


} else {
ifres15522 = False


}

let__13546 := ifres15522
_ = let__13546

tmp15631 := GEq(__e, let__13546, False)

if True == tmp15631 {
tmp15627 := Call1(__e, SymFn(symshen_4unlocked_2), V5358)


var ifres15601 Obj

if True == tmp15627 {
tmp15602 := Call2(__e, SymFn(symshen_4lazyderef), V5354, V5357)


let__13567 := tmp15602
_ = let__13567

tmp15626 := GConsP(__e, let__13567)

var ifres15603 Obj

if True == tmp15626 {
tmp15604 := GHd(__e, let__13567)

let__13568 := tmp15604
_ = let__13568

tmp15605 := GTl(__e, let__13567)

let__13569 := tmp15605
_ = let__13569

tmp15606 := Call2(__e, SymFn(symshen_4lazyderef), V5355, V5357)


let__13570 := tmp15606
_ = let__13570

tmp15607 := MakeNative1(func(__e *ControlFlow, Z5449 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5450 Obj) Obj {
tmp15608 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15608

tmp15609 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4l_1rules), let__13569, Z5450, V5356, V5357, V5358, let__13484, V5360)
return nil
panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), Z5449, let__13568, V5357, V5358, let__13484, tmp15609)


panic("unreachable")
})
panic("unreachable")
})

let__13571 := tmp15607
_ = let__13571

tmp15625 := GConsP(__e, let__13570)

var ifres15610 Obj

if True == tmp15625 {
tmp15611 := GHd(__e, let__13570)

let__13572 := tmp15611
_ = let__13572

tmp15612 := GTl(__e, let__13570)

let__13573 := tmp15612
_ = let__13573

tmp15613 := Call1(__e, let__13571, let__13572)


tmp15614 := Call1(__e, tmp15613, let__13573)


ifres15610 = tmp15614


} else {
tmp15624 := Call1(__e, SymFn(symshen_4pvar_2), let__13570)


var ifres15615 Obj

if True == tmp15624 {
tmp15616 := Call1(__e, SymFn(symshen_4newpv), V5357)


let__13574 := tmp15616
_ = let__13574

tmp15617 := Call1(__e, SymFn(symshen_4newpv), V5357)


let__13575 := tmp15617
_ = let__13575

tmp15618 := GCons(__e, let__13574, let__13575)

tmp15619 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15620 := Call1(__e, let__13571, let__13574)


return TailCall1(__e, tmp15620, let__13575)


panic("unreachable")
})

tmp15621 := Call4(__e, SymFn(symshen_4bind_b), let__13570, tmp15618, V5357, tmp15619)


tmp15622 := Call2(__e, SymFn(symshen_4gc), V5357, tmp15621)


tmp15623 := Call2(__e, SymFn(symshen_4gc), V5357, tmp15622)


ifres15615 = tmp15623


} else {
ifres15615 = False


}

ifres15610 = ifres15615


}

ifres15603 = ifres15610


} else {
ifres15603 = False


}

ifres15601 = ifres15603


} else {
ifres15601 = False


}

let__13566 := ifres15601
_ = let__13566

tmp15629 := GEq(__e, let__13566, False)

if True == tmp15629 {
return TailCall2(__e, SymFn(symshen_4unlock), V5358, let__13484)
} else {
return let__13566
}


} else {
return let__13546
}


} else {
return let__13530
}


} else {
return let__13508
}


} else {
return let__13488
}


} else {
return let__13485
}


panic("unreachable")
}())
}, 7)

tmp15640 := Call2(__e, ns2_1set, symshen_4l_1rules, tmp15284)


_ = tmp15640

tmp15641 := MakeNative6(func(__e *ControlFlow, V5455, V5456, V5457, V5458, V5459, V5460 Obj) Obj {
let__13576 := GAdd(__e, V5459, MakeInteger(1))
_ = let__13576

tmp15681 := Call1(__e, SymFn(symshen_4unlocked_2), V5458)


var ifres15643 Obj

if True == tmp15681 {
tmp15644 := Call2(__e, SymFn(symshen_4lazyderef), V5455, V5457)


let__13578 := tmp15644
_ = let__13578

tmp15680 := GConsP(__e, let__13578)

var ifres15645 Obj

if True == tmp15680 {
tmp15646 := GHd(__e, let__13578)

tmp15647 := Call2(__e, SymFn(symshen_4lazyderef), tmp15646, V5457)


let__13579 := tmp15647
_ = let__13579

tmp15679 := GEq(__e, let__13579, symdefine)

var ifres15648 Obj

if True == tmp15679 {
tmp15649 := GTl(__e, let__13578)

tmp15650 := Call2(__e, SymFn(symshen_4lazyderef), tmp15649, V5457)


let__13580 := tmp15650
_ = let__13580

tmp15678 := GConsP(__e, let__13580)

var ifres15651 Obj

if True == tmp15678 {
tmp15652 := GHd(__e, let__13580)

let__13581 := tmp15652
_ = let__13581

tmp15653 := GTl(__e, let__13580)

let__13582 := tmp15653
_ = let__13582

tmp15654 := Call1(__e, SymFn(symshen_4newpv), V5457)


let__13583 := tmp15654
_ = let__13583

tmp15655 := Call1(__e, SymFn(symshen_4newpv), V5457)


let__13584 := tmp15655
_ = let__13584

tmp15656 := Call1(__e, SymFn(symshen_4newpv), V5457)


let__13585 := tmp15656
_ = let__13585

tmp15657 := Call1(__e, SymFn(symshen_4newpv), V5457)


let__13586 := tmp15657
_ = let__13586

tmp15658 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp15658

tmp15659 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15660 := GCons(__e, let__13581, let__13582)

tmp15661 := Call1(__e, SymFn(symshen_4sigxrules), tmp15660)


tmp15662 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15663 := Call2(__e, SymFn(symshen_4lazyderef), let__13583, V5457)


tmp15664 := Call1(__e, SymFn(symfst), tmp15663)


tmp15665 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15666 := Call2(__e, SymFn(symshen_4lazyderef), let__13583, V5457)


tmp15667 := Call1(__e, SymFn(symsnd), tmp15666)


tmp15668 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15669 := Call2(__e, SymFn(symshen_4deref), let__13586, V5457)


tmp15670 := Call1(__e, SymFn(symshen_4freshen_1sig), tmp15669)


tmp15671 := MakeNative0(func(__e *ControlFlow) Obj {
tmp15672 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall6(__e, SymFn(symis), let__13586, V5456, V5457, V5458, let__13576, V5460)
panic("unreachable")
})

__e.TailApply(SymFn(symshen_4t_d_1rules), let__13581, let__13584, let__13585, MakeInteger(1), V5457, V5458, let__13576, tmp15672)
return nil


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), let__13585, tmp15670, V5457, V5458, let__13576, tmp15671)


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), let__13584, tmp15667, V5457, V5458, let__13576, tmp15668)


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), let__13586, tmp15664, V5457, V5458, let__13576, tmp15665)


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), let__13583, tmp15661, V5457, V5458, let__13576, tmp15662)


panic("unreachable")
})

tmp15673 := Call4(__e, SymFn(symshen_4cut), V5457, V5458, let__13576, tmp15659)


tmp15674 := Call2(__e, SymFn(symshen_4gc), V5457, tmp15673)


tmp15675 := Call2(__e, SymFn(symshen_4gc), V5457, tmp15674)


tmp15676 := Call2(__e, SymFn(symshen_4gc), V5457, tmp15675)


tmp15677 := Call2(__e, SymFn(symshen_4gc), V5457, tmp15676)


ifres15651 = tmp15677


} else {
ifres15651 = False


}

ifres15648 = ifres15651


} else {
ifres15648 = False


}

ifres15645 = ifres15648


} else {
ifres15645 = False


}

ifres15643 = ifres15645


} else {
ifres15643 = False


}

let__13577 := ifres15643
_ = let__13577

tmp15683 := GEq(__e, let__13577, False)

if True == tmp15683 {
return TailCall2(__e, SymFn(symshen_4unlock), V5458, let__13576)
} else {
return let__13577
}


panic("unreachable")
})

tmp15684 := Call2(__e, ns2_1set, symshen_4t_d, tmp15641)


_ = tmp15684

tmp15685 := MakeNative1(func(__e *ControlFlow, V5472 Obj) Obj {
tmp15686 := MakeNative1(func(__e *ControlFlow, Z5474 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5sig_drules_6), Z5474)
panic("unreachable")
})

tmp15687 := Call2(__e, SymFn(symcompile), tmp15686, V5472)


let__13587 := tmp15687
_ = let__13587

tmp15688 := Call1(__e, SymFn(symfst), let__13587)


tmp15689 := MakeNative1(func(__e *ControlFlow, Z5475 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4unmark_1rule_1custom_1patterns), Z5475)
panic("unreachable")
})

tmp15690 := Call1(__e, SymFn(symsnd), let__13587)


tmp15691 := Call2(__e, SymFn(symmap), tmp15689, tmp15690)


return TailCall2(__e, SymFn(sym_8p), tmp15688, tmp15691)


panic("unreachable")
})

tmp15692 := Call2(__e, ns2_1set, symshen_4sigxrules, tmp15685)


_ = tmp15692

tmp15693 := MakeNative1(func(__e *ControlFlow, V5476 Obj) Obj {
tmp15699 := Call1(__e, SymFn(symtuple_2), V5476)


if True == tmp15699 {
tmp15694 := MakeNative1(func(__e *ControlFlow, Z5477 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4unmark_1custom_1pattern), Z5477)
panic("unreachable")
})

tmp15695 := Call1(__e, SymFn(symfst), V5476)


tmp15696 := Call2(__e, SymFn(symmap), tmp15694, tmp15695)


tmp15697 := Call1(__e, SymFn(symsnd), V5476)


return TailCall2(__e, SymFn(sym_8p), tmp15696, tmp15697)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4unmark_1rule_1custom_1patterns)
}


panic("unreachable")
})

tmp15700 := Call2(__e, ns2_1set, symshen_4unmark_1rule_1custom_1patterns, tmp15693)


_ = tmp15700

tmp15701 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V5478 := __self1
_ = V5478
tmp15732 := GConsP(__e, V5478)

var ifres15713 Obj

if True == tmp15732 {
tmp15730 := GHd(__e, V5478)

tmp15731 := GEq(__e, sym_8p, tmp15730)

var ifres15714 Obj

if True == tmp15731 {
tmp15728 := GTl(__e, V5478)

tmp15729 := GConsP(__e, tmp15728)

var ifres15715 Obj

if True == tmp15729 {
tmp15725 := GTl(__e, V5478)

tmp15726 := GHd(__e, tmp15725)

tmp15727 := GEq(__e, symshen_4custom_1pattern, tmp15726)

var ifres15716 Obj

if True == tmp15727 {
tmp15722 := GTl(__e, V5478)

tmp15723 := GTl(__e, tmp15722)

tmp15724 := GConsP(__e, tmp15723)

var ifres15717 Obj

if True == tmp15724 {
tmp15718 := GTl(__e, V5478)

tmp15719 := GTl(__e, tmp15718)

tmp15720 := GTl(__e, tmp15719)

tmp15721 := GEq(__e, Nil, tmp15720)

ifres15717 = tmp15721


} else {
ifres15717 = False


}

ifres15716 = ifres15717


} else {
ifres15716 = False


}

ifres15715 = ifres15716


} else {
ifres15715 = False


}

ifres15714 = ifres15715


} else {
ifres15714 = False


}

ifres15713 = ifres15714


} else {
ifres15713 = False


}

if True == ifres15713 {
tmp15702 := GTl(__e, V5478)

tmp15703 := GTl(__e, tmp15702)

tmp15704 := GHd(__e, tmp15703)

tmp15705 := Call1(__e, SymFn(symshen_4curry), tmp15704)


if SymFn(symshen_4unmark_1custom_1pattern) == __self {
__self1 = tmp15705
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4unmark_1custom_1pattern), tmp15705)


} else {
tmp15711 := GConsP(__e, V5478)

if True == tmp15711 {
tmp15706 := GHd(__e, V5478)

tmp15707 := Call1(__e, SymFn(symshen_4unmark_1custom_1pattern), tmp15706)


tmp15708 := GTl(__e, V5478)

tmp15709 := Call1(__e, SymFn(symshen_4unmark_1custom_1pattern), tmp15708)


return GCons(__e, tmp15707, tmp15709)


} else {
return V5478
}


}


panic("unreachable")
})
return __self
})()

tmp15733 := Call2(__e, ns2_1set, symshen_4unmark_1custom_1pattern, tmp15701)


_ = tmp15733

tmp15734 := MakeNative1(func(__e *ControlFlow, V5479 Obj) Obj {
tmp15761 := GConsP(__e, V5479)

var ifres15735 Obj

if True == tmp15761 {
tmp15736 := Call1(__e, SymFn(symtail), V5479)


let__13589 := tmp15736
_ = let__13589

tmp15759 := Call2(__e, SymFn(symshen_4hds_a_2), let__13589, sym_i)


var ifres15737 Obj

if True == tmp15759 {
tmp15738 := Call1(__e, SymFn(symtail), let__13589)


let__13590 := tmp15738
_ = let__13590

tmp15739 := Call1(__e, SymFn(symshen_4_5signature_6), let__13590)


let__13591 := tmp15739
_ = let__13591

tmp15757 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13591)


var ifres15740 Obj

if True == tmp15757 {
tmp15741 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15740 = tmp15741


} else {
tmp15742 := Call1(__e, SymFn(symshen_4_5_1out), let__13591)


let__13592 := tmp15742
_ = let__13592

tmp15743 := Call1(__e, SymFn(symshen_4in_1_6), let__13591)


let__13593 := tmp15743
_ = let__13593

tmp15756 := Call2(__e, SymFn(symshen_4hds_a_2), let__13593, sym_j)


var ifres15744 Obj

if True == tmp15756 {
tmp15745 := Call1(__e, SymFn(symtail), let__13593)


let__13594 := tmp15745
_ = let__13594

tmp15746 := Call1(__e, SymFn(symshen_4_5rules_d_6), let__13594)


let__13595 := tmp15746
_ = let__13595

tmp15754 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13595)


var ifres15747 Obj

if True == tmp15754 {
tmp15748 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15747 = tmp15748


} else {
tmp15749 := Call1(__e, SymFn(symshen_4_5_1out), let__13595)


let__13596 := tmp15749
_ = let__13596

tmp15750 := Call1(__e, SymFn(symshen_4in_1_6), let__13595)


let__13597 := tmp15750
_ = let__13597

tmp15751 := Call1(__e, SymFn(symshen_4rectify_1type), let__13592)


let__13598 := tmp15751
_ = let__13598

tmp15752 := Call2(__e, SymFn(sym_8p), let__13598, let__13596)


tmp15753 := Call2(__e, SymFn(symshen_4comb), let__13597, tmp15752)


ifres15747 = tmp15753


}

ifres15744 = ifres15747


} else {
tmp15755 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15744 = tmp15755


}

ifres15740 = ifres15744


}

ifres15737 = ifres15740


} else {
tmp15758 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15737 = tmp15758


}

ifres15735 = ifres15737


} else {
tmp15760 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15735 = tmp15760


}

let__13588 := ifres15735
_ = let__13588

tmp15763 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13588)


if True == tmp15763 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__13588
}


panic("unreachable")
})

tmp15764 := Call2(__e, ns2_1set, symshen_4_5sig_drules_6, tmp15734)


_ = tmp15764

tmp15765 := MakeNative1(func(__e *ControlFlow, V5491 Obj) Obj {
tmp15766 := Call1(__e, SymFn(symshen_4extract_1vars), V5491)


let__13599 := tmp15766
_ = let__13599

tmp15767 := MakeNative1(func(__e *ControlFlow, Z5494 Obj) Obj {
tmp15768 := Call2(__e, SymFn(symconcat), sym_e, Z5494)


tmp15769 := Call1(__e, SymFn(symshen_4freshterm), tmp15768)


return GCons(__e, Z5494, tmp15769)


panic("unreachable")
})

tmp15770 := Call2(__e, SymFn(symmap), tmp15767, let__13599)


let__13600 := tmp15770
_ = let__13600

return TailCall2(__e, SymFn(symshen_4freshen_1type), let__13600, V5491)


panic("unreachable")
})

tmp15771 := Call2(__e, ns2_1set, symshen_4freshen_1sig, tmp15765)


_ = tmp15771

tmp15772 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V5495 := __self1
_ = V5495
V5496 := __self2
_ = V5496
tmp15785 := GEq(__e, Nil, V5495)

if True == tmp15785 {
return V5496
} else {
tmp15783 := GConsP(__e, V5495)

var ifres15780 Obj

if True == tmp15783 {
tmp15781 := GHd(__e, V5495)

tmp15782 := GConsP(__e, tmp15781)

ifres15780 = tmp15782


} else {
ifres15780 = False


}

if True == ifres15780 {
tmp15773 := GTl(__e, V5495)

tmp15774 := GHd(__e, V5495)

tmp15775 := GTl(__e, tmp15774)

tmp15776 := GHd(__e, V5495)

tmp15777 := GHd(__e, tmp15776)

tmp15778 := Call3(__e, SymFn(symsubst), tmp15775, tmp15777, V5496)


if SymFn(symshen_4freshen_1type) == __self {
__self1, __self2 = tmp15773, tmp15778
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4freshen_1type), tmp15773, tmp15778)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4freshen_1type)
}


}


panic("unreachable")
})
return __self
})()

tmp15786 := Call2(__e, ns2_1set, symshen_4freshen_1type, tmp15772)


_ = tmp15786

tmp15787 := MakeNative1(func(__e *ControlFlow, V5497 Obj) Obj {
tmp15788 := Call1(__e, SymFn(symshen_4_5rule_d_6), V5497)


let__13602 := tmp15788
_ = let__13602

tmp15801 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13602)


var ifres15789 Obj

if True == tmp15801 {
tmp15790 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15789 = tmp15790


} else {
tmp15791 := Call1(__e, SymFn(symshen_4_5_1out), let__13602)


let__13603 := tmp15791
_ = let__13603

tmp15792 := Call1(__e, SymFn(symshen_4in_1_6), let__13602)


let__13604 := tmp15792
_ = let__13604

tmp15793 := Call1(__e, SymFn(symshen_4_5rules_d_6), let__13604)


let__13605 := tmp15793
_ = let__13605

tmp15800 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13605)


var ifres15794 Obj

if True == tmp15800 {
tmp15795 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15794 = tmp15795


} else {
tmp15796 := Call1(__e, SymFn(symshen_4_5_1out), let__13605)


let__13606 := tmp15796
_ = let__13606

tmp15797 := Call1(__e, SymFn(symshen_4in_1_6), let__13605)


let__13607 := tmp15797
_ = let__13607

tmp15798 := GCons(__e, let__13603, let__13606)

tmp15799 := Call2(__e, SymFn(symshen_4comb), let__13607, tmp15798)


ifres15794 = tmp15799


}

ifres15789 = ifres15794


}

let__13601 := ifres15789
_ = let__13601

tmp15813 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13601)


if True == tmp15813 {
tmp15802 := Call1(__e, SymFn(symshen_4_5rule_d_6), V5497)


let__13609 := tmp15802
_ = let__13609

tmp15809 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13609)


var ifres15803 Obj

if True == tmp15809 {
tmp15804 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15803 = tmp15804


} else {
tmp15805 := Call1(__e, SymFn(symshen_4_5_1out), let__13609)


let__13610 := tmp15805
_ = let__13610

tmp15806 := Call1(__e, SymFn(symshen_4in_1_6), let__13609)


let__13611 := tmp15806
_ = let__13611

tmp15807 := GCons(__e, let__13610, Nil)

tmp15808 := Call2(__e, SymFn(symshen_4comb), let__13611, tmp15807)


ifres15803 = tmp15808


}

let__13608 := ifres15803
_ = let__13608

tmp15811 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13608)


if True == tmp15811 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__13608
}


} else {
return let__13601
}


panic("unreachable")
})

tmp15814 := Call2(__e, ns2_1set, symshen_4_5rules_d_6, tmp15787)


_ = tmp15814

tmp15815 := MakeNative1(func(__e *ControlFlow, V5509 Obj) Obj {
tmp15816 := Call1(__e, SymFn(symshen_4_5patterns_6), V5509)


let__13613 := tmp15816
_ = let__13613

tmp15844 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13613)


var ifres15817 Obj

if True == tmp15844 {
tmp15818 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15817 = tmp15818


} else {
tmp15819 := Call1(__e, SymFn(symshen_4_5_1out), let__13613)


let__13614 := tmp15819
_ = let__13614

tmp15820 := Call1(__e, SymFn(symshen_4in_1_6), let__13613)


let__13615 := tmp15820
_ = let__13615

tmp15843 := Call2(__e, SymFn(symshen_4hds_a_2), let__13615, sym_1_6)


var ifres15821 Obj

if True == tmp15843 {
tmp15822 := Call1(__e, SymFn(symtail), let__13615)


let__13616 := tmp15822
_ = let__13616

tmp15841 := GConsP(__e, let__13616)

var ifres15823 Obj

if True == tmp15841 {
tmp15824 := Call1(__e, SymFn(symhead), let__13616)


let__13617 := tmp15824
_ = let__13617

tmp15825 := Call1(__e, SymFn(symtail), let__13616)


let__13618 := tmp15825
_ = let__13618

tmp15839 := Call2(__e, SymFn(symshen_4hds_a_2), let__13618, symwhere)


var ifres15826 Obj

if True == tmp15839 {
tmp15827 := Call1(__e, SymFn(symtail), let__13618)


let__13619 := tmp15827
_ = let__13619

tmp15837 := GConsP(__e, let__13619)

var ifres15828 Obj

if True == tmp15837 {
tmp15829 := Call1(__e, SymFn(symhead), let__13619)


let__13620 := tmp15829
_ = let__13620

tmp15830 := Call1(__e, SymFn(symtail), let__13619)


let__13621 := tmp15830
_ = let__13621

tmp15831 := GCons(__e, let__13617, Nil)

tmp15832 := GCons(__e, let__13620, tmp15831)

tmp15833 := GCons(__e, symwhere, tmp15832)

tmp15834 := Call2(__e, SymFn(sym_8p), let__13614, tmp15833)


tmp15835 := Call2(__e, SymFn(symshen_4comb), let__13621, tmp15834)


ifres15828 = tmp15835


} else {
tmp15836 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15828 = tmp15836


}

ifres15826 = ifres15828


} else {
tmp15838 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15826 = tmp15838


}

ifres15823 = ifres15826


} else {
tmp15840 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15823 = tmp15840


}

ifres15821 = ifres15823


} else {
tmp15842 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15821 = tmp15842


}

ifres15817 = ifres15821


}

let__13612 := ifres15817
_ = let__13612

tmp15917 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13612)


if True == tmp15917 {
tmp15845 := Call1(__e, SymFn(symshen_4_5patterns_6), V5509)


let__13623 := tmp15845
_ = let__13623

tmp15874 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13623)


var ifres15846 Obj

if True == tmp15874 {
tmp15847 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15846 = tmp15847


} else {
tmp15848 := Call1(__e, SymFn(symshen_4_5_1out), let__13623)


let__13624 := tmp15848
_ = let__13624

tmp15849 := Call1(__e, SymFn(symshen_4in_1_6), let__13623)


let__13625 := tmp15849
_ = let__13625

tmp15873 := Call2(__e, SymFn(symshen_4hds_a_2), let__13625, sym_5_1)


var ifres15850 Obj

if True == tmp15873 {
tmp15851 := Call1(__e, SymFn(symtail), let__13625)


let__13626 := tmp15851
_ = let__13626

tmp15871 := GConsP(__e, let__13626)

var ifres15852 Obj

if True == tmp15871 {
tmp15853 := Call1(__e, SymFn(symhead), let__13626)


let__13627 := tmp15853
_ = let__13627

tmp15854 := Call1(__e, SymFn(symtail), let__13626)


let__13628 := tmp15854
_ = let__13628

tmp15869 := Call2(__e, SymFn(symshen_4hds_a_2), let__13628, symwhere)


var ifres15855 Obj

if True == tmp15869 {
tmp15856 := Call1(__e, SymFn(symtail), let__13628)


let__13629 := tmp15856
_ = let__13629

tmp15867 := GConsP(__e, let__13629)

var ifres15857 Obj

if True == tmp15867 {
tmp15858 := Call1(__e, SymFn(symhead), let__13629)


let__13630 := tmp15858
_ = let__13630

tmp15859 := Call1(__e, SymFn(symtail), let__13629)


let__13631 := tmp15859
_ = let__13631

tmp15860 := GCons(__e, let__13627, Nil)

tmp15861 := GCons(__e, let__13630, tmp15860)

tmp15862 := GCons(__e, symwhere, tmp15861)

tmp15863 := Call1(__e, SymFn(symshen_4correct), tmp15862)


tmp15864 := Call2(__e, SymFn(sym_8p), let__13624, tmp15863)


tmp15865 := Call2(__e, SymFn(symshen_4comb), let__13631, tmp15864)


ifres15857 = tmp15865


} else {
tmp15866 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15857 = tmp15866


}

ifres15855 = ifres15857


} else {
tmp15868 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15855 = tmp15868


}

ifres15852 = ifres15855


} else {
tmp15870 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15852 = tmp15870


}

ifres15850 = ifres15852


} else {
tmp15872 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15850 = tmp15872


}

ifres15846 = ifres15850


}

let__13622 := ifres15846
_ = let__13622

tmp15915 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13622)


if True == tmp15915 {
tmp15875 := Call1(__e, SymFn(symshen_4_5patterns_6), V5509)


let__13633 := tmp15875
_ = let__13633

tmp15892 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13633)


var ifres15876 Obj

if True == tmp15892 {
tmp15877 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15876 = tmp15877


} else {
tmp15878 := Call1(__e, SymFn(symshen_4_5_1out), let__13633)


let__13634 := tmp15878
_ = let__13634

tmp15879 := Call1(__e, SymFn(symshen_4in_1_6), let__13633)


let__13635 := tmp15879
_ = let__13635

tmp15891 := Call2(__e, SymFn(symshen_4hds_a_2), let__13635, sym_5_1)


var ifres15880 Obj

if True == tmp15891 {
tmp15881 := Call1(__e, SymFn(symtail), let__13635)


let__13636 := tmp15881
_ = let__13636

tmp15889 := GConsP(__e, let__13636)

var ifres15882 Obj

if True == tmp15889 {
tmp15883 := Call1(__e, SymFn(symhead), let__13636)


let__13637 := tmp15883
_ = let__13637

tmp15884 := Call1(__e, SymFn(symtail), let__13636)


let__13638 := tmp15884
_ = let__13638

tmp15885 := Call1(__e, SymFn(symshen_4correct), let__13637)


tmp15886 := Call2(__e, SymFn(sym_8p), let__13634, tmp15885)


tmp15887 := Call2(__e, SymFn(symshen_4comb), let__13638, tmp15886)


ifres15882 = tmp15887


} else {
tmp15888 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15882 = tmp15888


}

ifres15880 = ifres15882


} else {
tmp15890 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15880 = tmp15890


}

ifres15876 = ifres15880


}

let__13632 := ifres15876
_ = let__13632

tmp15913 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13632)


if True == tmp15913 {
tmp15893 := Call1(__e, SymFn(symshen_4_5patterns_6), V5509)


let__13640 := tmp15893
_ = let__13640

tmp15909 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13640)


var ifres15894 Obj

if True == tmp15909 {
tmp15895 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15894 = tmp15895


} else {
tmp15896 := Call1(__e, SymFn(symshen_4_5_1out), let__13640)


let__13641 := tmp15896
_ = let__13641

tmp15897 := Call1(__e, SymFn(symshen_4in_1_6), let__13640)


let__13642 := tmp15897
_ = let__13642

tmp15908 := Call2(__e, SymFn(symshen_4hds_a_2), let__13642, sym_1_6)


var ifres15898 Obj

if True == tmp15908 {
tmp15899 := Call1(__e, SymFn(symtail), let__13642)


let__13643 := tmp15899
_ = let__13643

tmp15906 := GConsP(__e, let__13643)

var ifres15900 Obj

if True == tmp15906 {
tmp15901 := Call1(__e, SymFn(symhead), let__13643)


let__13644 := tmp15901
_ = let__13644

tmp15902 := Call1(__e, SymFn(symtail), let__13643)


let__13645 := tmp15902
_ = let__13645

tmp15903 := Call2(__e, SymFn(sym_8p), let__13641, let__13644)


tmp15904 := Call2(__e, SymFn(symshen_4comb), let__13645, tmp15903)


ifres15900 = tmp15904


} else {
tmp15905 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15900 = tmp15905


}

ifres15898 = ifres15900


} else {
tmp15907 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres15898 = tmp15907


}

ifres15894 = ifres15898


}

let__13639 := ifres15894
_ = let__13639

tmp15911 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__13639)


if True == tmp15911 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__13639
}


} else {
return let__13632
}


} else {
return let__13622
}


} else {
return let__13612
}


panic("unreachable")
})

tmp15918 := Call2(__e, ns2_1set, symshen_4_5rule_d_6, tmp15815)


_ = tmp15918

tmp15919 := MakeNative1(func(__e *ControlFlow, V5544 Obj) Obj {
tmp16050 := GConsP(__e, V5544)

var ifres16003 Obj

if True == tmp16050 {
tmp16048 := GHd(__e, V5544)

tmp16049 := GEq(__e, symwhere, tmp16048)

var ifres16004 Obj

if True == tmp16049 {
tmp16046 := GTl(__e, V5544)

tmp16047 := GConsP(__e, tmp16046)

var ifres16005 Obj

if True == tmp16047 {
tmp16043 := GTl(__e, V5544)

tmp16044 := GTl(__e, tmp16043)

tmp16045 := GConsP(__e, tmp16044)

var ifres16006 Obj

if True == tmp16045 {
tmp16039 := GTl(__e, V5544)

tmp16040 := GTl(__e, tmp16039)

tmp16041 := GHd(__e, tmp16040)

tmp16042 := GConsP(__e, tmp16041)

var ifres16007 Obj

if True == tmp16042 {
tmp16034 := GTl(__e, V5544)

tmp16035 := GTl(__e, tmp16034)

tmp16036 := GHd(__e, tmp16035)

tmp16037 := GHd(__e, tmp16036)

tmp16038 := GEq(__e, symfail_1if, tmp16037)

var ifres16008 Obj

if True == tmp16038 {
tmp16029 := GTl(__e, V5544)

tmp16030 := GTl(__e, tmp16029)

tmp16031 := GHd(__e, tmp16030)

tmp16032 := GTl(__e, tmp16031)

tmp16033 := GConsP(__e, tmp16032)

var ifres16009 Obj

if True == tmp16033 {
tmp16023 := GTl(__e, V5544)

tmp16024 := GTl(__e, tmp16023)

tmp16025 := GHd(__e, tmp16024)

tmp16026 := GTl(__e, tmp16025)

tmp16027 := GTl(__e, tmp16026)

tmp16028 := GConsP(__e, tmp16027)

var ifres16010 Obj

if True == tmp16028 {
tmp16016 := GTl(__e, V5544)

tmp16017 := GTl(__e, tmp16016)

tmp16018 := GHd(__e, tmp16017)

tmp16019 := GTl(__e, tmp16018)

tmp16020 := GTl(__e, tmp16019)

tmp16021 := GTl(__e, tmp16020)

tmp16022 := GEq(__e, Nil, tmp16021)

var ifres16011 Obj

if True == tmp16022 {
tmp16012 := GTl(__e, V5544)

tmp16013 := GTl(__e, tmp16012)

tmp16014 := GTl(__e, tmp16013)

tmp16015 := GEq(__e, Nil, tmp16014)

ifres16011 = tmp16015


} else {
ifres16011 = False


}

ifres16010 = ifres16011


} else {
ifres16010 = False


}

ifres16009 = ifres16010


} else {
ifres16009 = False


}

ifres16008 = ifres16009


} else {
ifres16008 = False


}

ifres16007 = ifres16008


} else {
ifres16007 = False


}

ifres16006 = ifres16007


} else {
ifres16006 = False


}

ifres16005 = ifres16006


} else {
ifres16005 = False


}

ifres16004 = ifres16005


} else {
ifres16004 = False


}

ifres16003 = ifres16004


} else {
ifres16003 = False


}

if True == ifres16003 {
tmp15920 := GTl(__e, V5544)

tmp15921 := GHd(__e, tmp15920)

tmp15922 := GTl(__e, V5544)

tmp15923 := GTl(__e, tmp15922)

tmp15924 := GHd(__e, tmp15923)

tmp15925 := GTl(__e, tmp15924)

tmp15926 := GCons(__e, tmp15925, Nil)

tmp15927 := GCons(__e, symnot, tmp15926)

tmp15928 := GCons(__e, tmp15927, Nil)

tmp15929 := GCons(__e, tmp15921, tmp15928)

tmp15930 := GCons(__e, symand, tmp15929)

tmp15931 := GTl(__e, V5544)

tmp15932 := GTl(__e, tmp15931)

tmp15933 := GHd(__e, tmp15932)

tmp15934 := GTl(__e, tmp15933)

tmp15935 := GTl(__e, tmp15934)

tmp15936 := GCons(__e, tmp15930, tmp15935)

return GCons(__e, symwhere, tmp15936)


} else {
tmp16001 := GConsP(__e, V5544)

var ifres15986 Obj

if True == tmp16001 {
tmp15999 := GHd(__e, V5544)

tmp16000 := GEq(__e, symwhere, tmp15999)

var ifres15987 Obj

if True == tmp16000 {
tmp15997 := GTl(__e, V5544)

tmp15998 := GConsP(__e, tmp15997)

var ifres15988 Obj

if True == tmp15998 {
tmp15994 := GTl(__e, V5544)

tmp15995 := GTl(__e, tmp15994)

tmp15996 := GConsP(__e, tmp15995)

var ifres15989 Obj

if True == tmp15996 {
tmp15990 := GTl(__e, V5544)

tmp15991 := GTl(__e, tmp15990)

tmp15992 := GTl(__e, tmp15991)

tmp15993 := GEq(__e, Nil, tmp15992)

ifres15989 = tmp15993


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

if True == ifres15986 {
tmp15937 := GTl(__e, V5544)

tmp15938 := GHd(__e, tmp15937)

tmp15939 := GTl(__e, V5544)

tmp15940 := GTl(__e, tmp15939)

tmp15941 := GHd(__e, tmp15940)

tmp15942 := GCons(__e, symfail, Nil)

tmp15943 := GCons(__e, tmp15942, Nil)

tmp15944 := GCons(__e, tmp15941, tmp15943)

tmp15945 := GCons(__e, sym_a_a, tmp15944)

tmp15946 := GCons(__e, tmp15945, Nil)

tmp15947 := GCons(__e, symnot, tmp15946)

tmp15948 := GCons(__e, tmp15947, Nil)

tmp15949 := GCons(__e, tmp15938, tmp15948)

tmp15950 := GCons(__e, symand, tmp15949)

tmp15951 := GTl(__e, V5544)

tmp15952 := GTl(__e, tmp15951)

tmp15953 := GCons(__e, tmp15950, tmp15952)

return GCons(__e, symwhere, tmp15953)


} else {
tmp15984 := GConsP(__e, V5544)

var ifres15969 Obj

if True == tmp15984 {
tmp15982 := GHd(__e, V5544)

tmp15983 := GEq(__e, symfail_1if, tmp15982)

var ifres15970 Obj

if True == tmp15983 {
tmp15980 := GTl(__e, V5544)

tmp15981 := GConsP(__e, tmp15980)

var ifres15971 Obj

if True == tmp15981 {
tmp15977 := GTl(__e, V5544)

tmp15978 := GTl(__e, tmp15977)

tmp15979 := GConsP(__e, tmp15978)

var ifres15972 Obj

if True == tmp15979 {
tmp15973 := GTl(__e, V5544)

tmp15974 := GTl(__e, tmp15973)

tmp15975 := GTl(__e, tmp15974)

tmp15976 := GEq(__e, Nil, tmp15975)

ifres15972 = tmp15976


} else {
ifres15972 = False


}

ifres15971 = ifres15972


} else {
ifres15971 = False


}

ifres15970 = ifres15971


} else {
ifres15970 = False


}

ifres15969 = ifres15970


} else {
ifres15969 = False


}

if True == ifres15969 {
tmp15954 := GTl(__e, V5544)

tmp15955 := GCons(__e, tmp15954, Nil)

tmp15956 := GCons(__e, symnot, tmp15955)

tmp15957 := GTl(__e, V5544)

tmp15958 := GTl(__e, tmp15957)

tmp15959 := GCons(__e, tmp15956, tmp15958)

return GCons(__e, symwhere, tmp15959)


} else {
tmp15960 := GCons(__e, symfail, Nil)

tmp15961 := GCons(__e, tmp15960, Nil)

tmp15962 := GCons(__e, V5544, tmp15961)

tmp15963 := GCons(__e, sym_a_a, tmp15962)

tmp15964 := GCons(__e, tmp15963, Nil)

tmp15965 := GCons(__e, symnot, tmp15964)

tmp15966 := GCons(__e, V5544, Nil)

tmp15967 := GCons(__e, tmp15965, tmp15966)

return GCons(__e, symwhere, tmp15967)


}


}


}


panic("unreachable")
})

tmp16051 := Call2(__e, ns2_1set, symshen_4correct, tmp15919)


_ = tmp16051

tmp16052 := MakeNative(func(__e *ControlFlow) {
V5545 := __e.Get(1)
V5546 := __e.Get(2)
V5547 := __e.Get(3)
V5548 := __e.Get(4)
V5549 := __e.Get(5)
V5550 := __e.Get(6)
V5551 := __e.Get(7)
V5552 := __e.Get(8)
__e.ReturnDirect(func() Obj {
let__13646 := GAdd(__e, V5551, MakeInteger(1))
_ = let__13646

tmp16060 := Call1(__e, SymFn(symshen_4unlocked_2), V5550)


var ifres16054 Obj

if True == tmp16060 {
tmp16055 := Call2(__e, SymFn(symshen_4lazyderef), V5546, V5549)


let__13648 := tmp16055
_ = let__13648

tmp16059 := GEq(__e, let__13648, Nil)

var ifres16056 Obj

if True == tmp16059 {
tmp16057 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16057

tmp16058 := Call1(__e, SymFn(symthaw), V5552)


ifres16056 = tmp16058


} else {
ifres16056 = False


}

ifres16054 = ifres16056


} else {
ifres16054 = False


}

let__13647 := ifres16054
_ = let__13647

tmp16085 := GEq(__e, let__13647, False)

if True == tmp16085 {
tmp16081 := Call1(__e, SymFn(symshen_4unlocked_2), V5550)


var ifres16061 Obj

if True == tmp16081 {
tmp16062 := Call2(__e, SymFn(symshen_4lazyderef), V5546, V5549)


let__13650 := tmp16062
_ = let__13650

tmp16080 := GConsP(__e, let__13650)

var ifres16063 Obj

if True == tmp16080 {
tmp16064 := GHd(__e, let__13650)

let__13651 := tmp16064
_ = let__13651

tmp16065 := GTl(__e, let__13650)

let__13652 := tmp16065
_ = let__13652

tmp16066 := Call1(__e, SymFn(symshen_4newpv), V5549)


let__13653 := tmp16066
_ = let__13653

tmp16067 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16067

tmp16068 := Call2(__e, SymFn(symshen_4deref), let__13651, V5549)


tmp16069 := Call1(__e, SymFn(symshen_4freshen_1rule), tmp16068)


tmp16070 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16071 := Call2(__e, SymFn(symshen_4lazyderef), let__13653, V5549)


tmp16072 := Call1(__e, SymFn(symfst), tmp16071)


tmp16073 := Call2(__e, SymFn(symshen_4lazyderef), let__13653, V5549)


tmp16074 := Call1(__e, SymFn(symsnd), tmp16073)


tmp16075 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16076 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4t_d_1rules), V5545, let__13652, V5547, GAdd(__e, V5548, MakeInteger(1)), V5549, V5550, let__13646, V5552)
return nil


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5549, V5550, let__13646, tmp16076)


panic("unreachable")
})

__e.TailApply(SymFn(symshen_4t_d_1rule), V5545, V5548, tmp16072, tmp16074, V5547, V5549, V5550, let__13646, tmp16075)
return nil


panic("unreachable")
})

tmp16078 := Call6(__e, SymFn(symbind), let__13653, tmp16069, V5549, V5550, let__13646, tmp16070)


tmp16079 := Call2(__e, SymFn(symshen_4gc), V5549, tmp16078)


ifres16063 = tmp16079


} else {
ifres16063 = False


}

ifres16061 = ifres16063


} else {
ifres16061 = False


}

let__13649 := ifres16061
_ = let__13649

tmp16083 := GEq(__e, let__13649, False)

if True == tmp16083 {
return TailCall2(__e, SymFn(symshen_4unlock), V5550, let__13646)
} else {
return let__13649
}


} else {
return let__13647
}


panic("unreachable")
}())
}, 8)

tmp16086 := Call2(__e, ns2_1set, symshen_4t_d_1rules, tmp16052)


_ = tmp16086

tmp16087 := MakeNative1(func(__e *ControlFlow, V5561 Obj) Obj {
tmp16098 := Call1(__e, SymFn(symtuple_2), V5561)


if True == tmp16098 {
tmp16088 := Call1(__e, SymFn(symfst), V5561)


tmp16089 := Call1(__e, SymFn(symshen_4extract_1vars), tmp16088)


let__13654 := tmp16089
_ = let__13654

tmp16090 := MakeNative1(func(__e *ControlFlow, Z5564 Obj) Obj {
tmp16091 := Call1(__e, SymFn(symshen_4freshterm), Z5564)


return GCons(__e, Z5564, tmp16091)


panic("unreachable")
})

tmp16092 := Call2(__e, SymFn(symmap), tmp16090, let__13654)


let__13655 := tmp16092
_ = let__13655

tmp16093 := Call1(__e, SymFn(symfst), V5561)


tmp16094 := Call2(__e, SymFn(symshen_4freshen), let__13655, tmp16093)


tmp16095 := Call1(__e, SymFn(symsnd), V5561)


tmp16096 := Call2(__e, SymFn(symshen_4freshen), let__13655, tmp16095)


return TailCall2(__e, SymFn(sym_8p), tmp16094, tmp16096)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4freshen_1rule)
}


panic("unreachable")
})

tmp16099 := Call2(__e, ns2_1set, symshen_4freshen_1rule, tmp16087)


_ = tmp16099

tmp16100 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V5565 := __self1
_ = V5565
V5566 := __self2
_ = V5566
tmp16113 := GEq(__e, Nil, V5565)

if True == tmp16113 {
return V5566
} else {
tmp16111 := GConsP(__e, V5565)

var ifres16108 Obj

if True == tmp16111 {
tmp16109 := GHd(__e, V5565)

tmp16110 := GConsP(__e, tmp16109)

ifres16108 = tmp16110


} else {
ifres16108 = False


}

if True == ifres16108 {
tmp16101 := GTl(__e, V5565)

tmp16102 := GHd(__e, V5565)

tmp16103 := GHd(__e, tmp16102)

tmp16104 := GHd(__e, V5565)

tmp16105 := GTl(__e, tmp16104)

tmp16106 := Call3(__e, SymFn(symshen_4beta), tmp16103, tmp16105, V5566)


if SymFn(symshen_4freshen) == __self {
__self1, __self2 = tmp16101, tmp16106
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4freshen), tmp16101, tmp16106)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4freshen)
}


}


panic("unreachable")
})
return __self
})()

tmp16114 := Call2(__e, ns2_1set, symshen_4freshen, tmp16100)


_ = tmp16114

tmp16115 := MakeNative(func(__e *ControlFlow) {
V5567 := __e.Get(1)
V5568 := __e.Get(2)
V5569 := __e.Get(3)
V5570 := __e.Get(4)
V5571 := __e.Get(5)
V5572 := __e.Get(6)
V5573 := __e.Get(7)
V5574 := __e.Get(8)
V5575 := __e.Get(9)
__e.ReturnDirect(func() Obj {
tmp16119 := Call1(__e, SymFn(symshen_4unlocked_2), V5573)


var ifres16116 Obj

if True == tmp16119 {
tmp16117 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16117

tmp16118 := Call(__e, SymFn(symshen_4t_d_1rule_1h), V5569, V5570, V5571, V5572, V5573, V5574, V5575)


ifres16116 = tmp16118


} else {
ifres16116 = False


}

let__13656 := ifres16116
_ = let__13656

tmp16131 := GEq(__e, let__13656, False)

if True == tmp16131 {
tmp16129 := Call1(__e, SymFn(symshen_4unlocked_2), V5573)


if True == tmp16129 {
tmp16120 := Call1(__e, SymFn(symshen_4newpv), V5572)


let__13657 := tmp16120
_ = let__13657

tmp16121 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16121

tmp16122 := Call3(__e, SymFn(symshen_4app), V5567, MakeString("\n"), symshen_4a)


tmp16124 := Call3(__e, SymFn(symshen_4app), V5568, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" of "), tmp16122), symshen_4a)


tmp16126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("type error in rule "), tmp16124))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("type error in rule "), tmp16124)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()

tmp16127 := Call6(__e, SymFn(symbind), let__13657, tmp16126, V5572, V5573, V5574, V5575)


return TailCall2(__e, SymFn(symshen_4gc), V5572, tmp16127)


} else {
return False
}


} else {
return let__13656
}


panic("unreachable")
}())
}, 9)

tmp16132 := Call2(__e, ns2_1set, symshen_4t_d_1rule, tmp16115)


_ = tmp16132

tmp16133 := MakeNative(func(__e *ControlFlow) {
V5578 := __e.Get(1)
V5579 := __e.Get(2)
V5580 := __e.Get(3)
V5581 := __e.Get(4)
V5582 := __e.Get(5)
V5583 := __e.Get(6)
V5584 := __e.Get(7)
__e.ReturnDirect(func() Obj {
let__13658 := GAdd(__e, V5583, MakeInteger(1))
_ = let__13658

tmp16158 := Call1(__e, SymFn(symshen_4unlocked_2), V5582)


var ifres16135 Obj

if True == tmp16158 {
tmp16136 := Call2(__e, SymFn(symshen_4lazyderef), V5578, V5581)


let__13660 := tmp16136
_ = let__13660

tmp16157 := GEq(__e, let__13660, Nil)

var ifres16137 Obj

if True == tmp16157 {
tmp16138 := Call2(__e, SymFn(symshen_4lazyderef), V5580, V5581)


let__13661 := tmp16138
_ = let__13661

tmp16156 := GConsP(__e, let__13661)

var ifres16139 Obj

if True == tmp16156 {
tmp16140 := GHd(__e, let__13661)

tmp16141 := Call2(__e, SymFn(symshen_4lazyderef), tmp16140, V5581)


let__13662 := tmp16141
_ = let__13662

tmp16155 := GEq(__e, let__13662, sym_1_1_6)

var ifres16142 Obj

if True == tmp16155 {
tmp16143 := GTl(__e, let__13661)

tmp16144 := Call2(__e, SymFn(symshen_4lazyderef), tmp16143, V5581)


let__13663 := tmp16144
_ = let__13663

tmp16154 := GConsP(__e, let__13663)

var ifres16145 Obj

if True == tmp16154 {
tmp16146 := GHd(__e, let__13663)

let__13664 := tmp16146
_ = let__13664

tmp16147 := GTl(__e, let__13663)

tmp16148 := Call2(__e, SymFn(symshen_4lazyderef), tmp16147, V5581)


let__13665 := tmp16148
_ = let__13665

tmp16153 := GEq(__e, let__13665, Nil)

var ifres16149 Obj

if True == tmp16153 {
tmp16150 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16150

tmp16151 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4t_d_1correct), V5579, let__13664, Nil, V5581, V5582, let__13658, V5584)
return nil
panic("unreachable")
})

tmp16152 := Call4(__e, SymFn(symshen_4cut), V5581, V5582, let__13658, tmp16151)


ifres16149 = tmp16152


} else {
ifres16149 = False


}

ifres16145 = ifres16149


} else {
ifres16145 = False


}

ifres16142 = ifres16145


} else {
ifres16142 = False


}

ifres16139 = ifres16142


} else {
ifres16139 = False


}

ifres16137 = ifres16139


} else {
ifres16137 = False


}

ifres16135 = ifres16137


} else {
ifres16135 = False


}

let__13659 := ifres16135
_ = let__13659

tmp16177 := GEq(__e, let__13659, False)

if True == tmp16177 {
tmp16173 := Call1(__e, SymFn(symshen_4unlocked_2), V5582)


var ifres16159 Obj

if True == tmp16173 {
tmp16160 := Call1(__e, SymFn(symshen_4newpv), V5581)


let__13667 := tmp16160
_ = let__13667

tmp16161 := Call1(__e, SymFn(symshen_4newpv), V5581)


let__13668 := tmp16161
_ = let__13668

tmp16162 := Call1(__e, SymFn(symshen_4newpv), V5581)


let__13669 := tmp16162
_ = let__13669

tmp16163 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16163

tmp16164 := Call1(__e, SymFn(symshen_4freshterms), V5578)


tmp16165 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16166 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16167 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16168 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4t_d_1correct), V5579, let__13668, let__13669, V5581, V5582, let__13658, V5584)
return nil
panic("unreachable")
})

__e.TailApply(SymFn(symshen_4myassume), V5578, V5580, let__13669, V5581, V5582, let__13658, tmp16168)
return nil


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5581, V5582, let__13658, tmp16167)


panic("unreachable")
})

__e.TailApply(SymFn(symshen_4t_d_1integrity), V5578, V5580, let__13667, let__13668, V5581, V5582, let__13658, tmp16166)
return nil


panic("unreachable")
})

tmp16169 := Call6(__e, SymFn(symshen_4p_1hyps), tmp16164, let__13667, V5581, V5582, let__13658, tmp16165)


tmp16170 := Call2(__e, SymFn(symshen_4gc), V5581, tmp16169)


tmp16171 := Call2(__e, SymFn(symshen_4gc), V5581, tmp16170)


tmp16172 := Call2(__e, SymFn(symshen_4gc), V5581, tmp16171)


ifres16159 = tmp16172


} else {
ifres16159 = False


}

let__13666 := ifres16159
_ = let__13666

tmp16175 := GEq(__e, let__13666, False)

if True == tmp16175 {
return TailCall2(__e, SymFn(symshen_4unlock), V5582, let__13658)
} else {
return let__13666
}


} else {
return let__13659
}


panic("unreachable")
}())
}, 7)

tmp16178 := Call2(__e, ns2_1set, symshen_4t_d_1rule_1h, tmp16133)


_ = tmp16178

tmp16179 := MakeNative(func(__e *ControlFlow) {
V5597 := __e.Get(1)
V5598 := __e.Get(2)
V5599 := __e.Get(3)
V5600 := __e.Get(4)
V5601 := __e.Get(5)
V5602 := __e.Get(6)
V5603 := __e.Get(7)
__e.ReturnDirect(func() Obj {
tmp16193 := Call1(__e, SymFn(symshen_4unlocked_2), V5601)


var ifres16180 Obj

if True == tmp16193 {
tmp16181 := Call2(__e, SymFn(symshen_4lazyderef), V5597, V5600)


let__13671 := tmp16181
_ = let__13671

tmp16192 := GEq(__e, let__13671, Nil)

var ifres16182 Obj

if True == tmp16192 {
tmp16183 := Call2(__e, SymFn(symshen_4lazyderef), V5599, V5600)


let__13672 := tmp16183
_ = let__13672

tmp16184 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16185 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16185

return TailCall1(__e, SymFn(symthaw), V5603)


panic("unreachable")
})

let__13673 := tmp16184
_ = let__13673

tmp16191 := GEq(__e, let__13672, Nil)

var ifres16186 Obj

if True == tmp16191 {
tmp16187 := Call1(__e, SymFn(symthaw), let__13673)


ifres16186 = tmp16187


} else {
tmp16190 := Call1(__e, SymFn(symshen_4pvar_2), let__13672)


var ifres16188 Obj

if True == tmp16190 {
tmp16189 := Call4(__e, SymFn(symshen_4bind_b), let__13672, Nil, V5600, let__13673)


ifres16188 = tmp16189


} else {
ifres16188 = False


}

ifres16186 = ifres16188


}

ifres16182 = ifres16186


} else {
ifres16182 = False


}

ifres16180 = ifres16182


} else {
ifres16180 = False


}

let__13670 := ifres16180
_ = let__13670

tmp16312 := GEq(__e, let__13670, False)

if True == tmp16312 {
tmp16310 := Call1(__e, SymFn(symshen_4unlocked_2), V5601)


if True == tmp16310 {
tmp16194 := Call2(__e, SymFn(symshen_4lazyderef), V5597, V5600)


let__13674 := tmp16194
_ = let__13674

tmp16308 := GConsP(__e, let__13674)

if True == tmp16308 {
tmp16195 := GHd(__e, let__13674)

let__13675 := tmp16195
_ = let__13675

tmp16196 := GTl(__e, let__13674)

let__13676 := tmp16196
_ = let__13676

tmp16197 := Call2(__e, SymFn(symshen_4lazyderef), V5598, V5600)


let__13677 := tmp16197
_ = let__13677

tmp16306 := GConsP(__e, let__13677)

if True == tmp16306 {
tmp16198 := GHd(__e, let__13677)

let__13678 := tmp16198
_ = let__13678

tmp16199 := GTl(__e, let__13677)

tmp16200 := Call2(__e, SymFn(symshen_4lazyderef), tmp16199, V5600)


let__13679 := tmp16200
_ = let__13679

tmp16304 := GConsP(__e, let__13679)

if True == tmp16304 {
tmp16201 := GHd(__e, let__13679)

tmp16202 := Call2(__e, SymFn(symshen_4lazyderef), tmp16201, V5600)


let__13680 := tmp16202
_ = let__13680

tmp16302 := GEq(__e, let__13680, sym_1_1_6)

if True == tmp16302 {
tmp16203 := GTl(__e, let__13679)

tmp16204 := Call2(__e, SymFn(symshen_4lazyderef), tmp16203, V5600)


let__13681 := tmp16204
_ = let__13681

tmp16300 := GConsP(__e, let__13681)

if True == tmp16300 {
tmp16205 := GHd(__e, let__13681)

let__13682 := tmp16205
_ = let__13682

tmp16206 := GTl(__e, let__13681)

tmp16207 := Call2(__e, SymFn(symshen_4lazyderef), tmp16206, V5600)


let__13683 := tmp16207
_ = let__13683

tmp16298 := GEq(__e, let__13683, Nil)

if True == tmp16298 {
tmp16208 := Call2(__e, SymFn(symshen_4lazyderef), V5599, V5600)


let__13684 := tmp16208
_ = let__13684

tmp16209 := MakeNative1(func(__e *ControlFlow, Z5620 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5621 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5622 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5623 Obj) Obj {
tmp16210 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16210

tmp16211 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16212 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16213 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp16214 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4myassume), let__13676, let__13682, Z5623, V5600, V5601, V5602, V5603)
return nil
panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), Z5621, tmp16213, V5600, V5601, V5602, tmp16214)


panic("unreachable")
})

return TailCall6(__e, SymFn(symis_b), let__13675, Z5620, V5600, V5601, V5602, tmp16212)


panic("unreachable")
})

return TailCall6(__e, SymFn(symis_b), let__13678, Z5622, V5600, V5601, V5602, tmp16211)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

let__13685 := tmp16209
_ = let__13685

tmp16296 := GConsP(__e, let__13684)

if True == tmp16296 {
tmp16215 := GHd(__e, let__13684)

tmp16216 := Call2(__e, SymFn(symshen_4lazyderef), tmp16215, V5600)


let__13686 := tmp16216
_ = let__13686

tmp16217 := MakeNative1(func(__e *ControlFlow, Z5626 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5627 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5628 Obj) Obj {
tmp16218 := GTl(__e, let__13684)

let__13688 := tmp16218
_ = let__13688

tmp16219 := Call1(__e, let__13685, Z5626)


tmp16220 := Call1(__e, tmp16219, Z5627)


tmp16221 := Call1(__e, tmp16220, Z5628)


return TailCall1(__e, tmp16221, let__13688)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

let__13687 := tmp16217
_ = let__13687

tmp16276 := GConsP(__e, let__13686)

if True == tmp16276 {
tmp16222 := GHd(__e, let__13686)

let__13689 := tmp16222
_ = let__13689

tmp16223 := GTl(__e, let__13686)

tmp16224 := Call2(__e, SymFn(symshen_4lazyderef), tmp16223, V5600)


let__13690 := tmp16224
_ = let__13690

tmp16225 := MakeNative1(func(__e *ControlFlow, Z5633 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5634 Obj) Obj {
tmp16226 := Call1(__e, let__13687, let__13689)


tmp16227 := Call1(__e, tmp16226, Z5633)


return TailCall1(__e, tmp16227, Z5634)


panic("unreachable")
})
panic("unreachable")
})

let__13691 := tmp16225
_ = let__13691

tmp16260 := GConsP(__e, let__13690)

if True == tmp16260 {
tmp16228 := GHd(__e, let__13690)

let__13692 := tmp16228
_ = let__13692

tmp16229 := GTl(__e, let__13690)

tmp16230 := Call2(__e, SymFn(symshen_4lazyderef), tmp16229, V5600)


let__13693 := tmp16230
_ = let__13693

tmp16231 := MakeNative1(func(__e *ControlFlow, Z5638 Obj) Obj {
tmp16232 := Call1(__e, let__13691, let__13692)


return TailCall1(__e, tmp16232, Z5638)


panic("unreachable")
})

let__13694 := tmp16231
_ = let__13694

tmp16248 := GConsP(__e, let__13693)

if True == tmp16248 {
tmp16233 := GHd(__e, let__13693)

let__13695 := tmp16233
_ = let__13695

tmp16234 := GTl(__e, let__13693)

tmp16235 := Call2(__e, SymFn(symshen_4lazyderef), tmp16234, V5600)


let__13696 := tmp16235
_ = let__13696

tmp16236 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13694, let__13695)
panic("unreachable")
})

let__13697 := tmp16236
_ = let__13697

tmp16240 := GEq(__e, let__13696, Nil)

if True == tmp16240 {
return TailCall1(__e, SymFn(symthaw), let__13697)
} else {
tmp16238 := Call1(__e, SymFn(symshen_4pvar_2), let__13696)


if True == tmp16238 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__13696, Nil, V5600, let__13697)
} else {
return False
}


}


} else {
tmp16246 := Call1(__e, SymFn(symshen_4pvar_2), let__13693)


if True == tmp16246 {
tmp16241 := Call1(__e, SymFn(symshen_4newpv), V5600)


let__13698 := tmp16241
_ = let__13698

tmp16242 := GCons(__e, let__13698, Nil)

tmp16243 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13694, let__13698)
panic("unreachable")
})

tmp16244 := Call4(__e, SymFn(symshen_4bind_b), let__13693, tmp16242, V5600, tmp16243)


return TailCall2(__e, SymFn(symshen_4gc), V5600, tmp16244)


} else {
return False
}


}


} else {
tmp16258 := Call1(__e, SymFn(symshen_4pvar_2), let__13690)


if True == tmp16258 {
tmp16249 := Call1(__e, SymFn(symshen_4newpv), V5600)


let__13699 := tmp16249
_ = let__13699

tmp16250 := Call1(__e, SymFn(symshen_4newpv), V5600)


let__13700 := tmp16250
_ = let__13700

tmp16251 := GCons(__e, let__13700, Nil)

tmp16252 := GCons(__e, let__13699, tmp16251)

tmp16253 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16254 := Call1(__e, let__13691, let__13699)


return TailCall1(__e, tmp16254, let__13700)


panic("unreachable")
})

tmp16255 := Call4(__e, SymFn(symshen_4bind_b), let__13690, tmp16252, V5600, tmp16253)


tmp16256 := Call2(__e, SymFn(symshen_4gc), V5600, tmp16255)


return TailCall2(__e, SymFn(symshen_4gc), V5600, tmp16256)


} else {
return False
}


}


} else {
tmp16274 := Call1(__e, SymFn(symshen_4pvar_2), let__13686)


if True == tmp16274 {
tmp16261 := Call1(__e, SymFn(symshen_4newpv), V5600)


let__13701 := tmp16261
_ = let__13701

tmp16262 := Call1(__e, SymFn(symshen_4newpv), V5600)


let__13702 := tmp16262
_ = let__13702

tmp16263 := Call1(__e, SymFn(symshen_4newpv), V5600)


let__13703 := tmp16263
_ = let__13703

tmp16264 := GCons(__e, let__13703, Nil)

tmp16265 := GCons(__e, let__13702, tmp16264)

tmp16266 := GCons(__e, let__13701, tmp16265)

tmp16267 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16268 := Call1(__e, let__13687, let__13701)


tmp16269 := Call1(__e, tmp16268, let__13702)


return TailCall1(__e, tmp16269, let__13703)


panic("unreachable")
})

tmp16270 := Call4(__e, SymFn(symshen_4bind_b), let__13686, tmp16266, V5600, tmp16267)


tmp16271 := Call2(__e, SymFn(symshen_4gc), V5600, tmp16270)


tmp16272 := Call2(__e, SymFn(symshen_4gc), V5600, tmp16271)


return TailCall2(__e, SymFn(symshen_4gc), V5600, tmp16272)


} else {
return False
}


}


} else {
tmp16294 := Call1(__e, SymFn(symshen_4pvar_2), let__13684)


if True == tmp16294 {
tmp16277 := Call1(__e, SymFn(symshen_4newpv), V5600)


let__13704 := tmp16277
_ = let__13704

tmp16278 := Call1(__e, SymFn(symshen_4newpv), V5600)


let__13705 := tmp16278
_ = let__13705

tmp16279 := Call1(__e, SymFn(symshen_4newpv), V5600)


let__13706 := tmp16279
_ = let__13706

tmp16280 := Call1(__e, SymFn(symshen_4newpv), V5600)


let__13707 := tmp16280
_ = let__13707

tmp16281 := GCons(__e, let__13706, Nil)

tmp16282 := GCons(__e, let__13705, tmp16281)

tmp16283 := GCons(__e, let__13704, tmp16282)

tmp16284 := GCons(__e, tmp16283, let__13707)

tmp16285 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16286 := Call1(__e, let__13685, let__13704)


tmp16287 := Call1(__e, tmp16286, let__13705)


tmp16288 := Call1(__e, tmp16287, let__13706)


return TailCall1(__e, tmp16288, let__13707)


panic("unreachable")
})

tmp16289 := Call4(__e, SymFn(symshen_4bind_b), let__13684, tmp16284, V5600, tmp16285)


tmp16290 := Call2(__e, SymFn(symshen_4gc), V5600, tmp16289)


tmp16291 := Call2(__e, SymFn(symshen_4gc), V5600, tmp16290)


tmp16292 := Call2(__e, SymFn(symshen_4gc), V5600, tmp16291)


return TailCall2(__e, SymFn(symshen_4gc), V5600, tmp16292)


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
return let__13670
}


panic("unreachable")
}())
}, 7)

tmp16313 := Call2(__e, ns2_1set, symshen_4myassume, tmp16179)


_ = tmp16313

tmp16314 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V5654 := __self1
_ = V5654
tmp16335 := GEq(__e, Nil, V5654)

if True == tmp16335 {
return Nil
} else {
tmp16333 := GConsP(__e, V5654)

var ifres16330 Obj

if True == tmp16333 {
tmp16331 := GHd(__e, V5654)

tmp16332 := GConsP(__e, tmp16331)

ifres16330 = tmp16332


} else {
ifres16330 = False


}

if True == ifres16330 {
tmp16315 := GHd(__e, V5654)

tmp16316 := GTl(__e, V5654)

tmp16317 := Call2(__e, SymFn(symappend), tmp16315, tmp16316)


if SymFn(symshen_4freshterms) == __self {
__self1 = tmp16317
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4freshterms), tmp16317)


} else {
tmp16328 := GConsP(__e, V5654)

var ifres16325 Obj

if True == tmp16328 {
tmp16326 := GHd(__e, V5654)

tmp16327 := Call1(__e, SymFn(symshen_4freshterm_2), tmp16326)


ifres16325 = tmp16327


} else {
ifres16325 = False


}

if True == ifres16325 {
tmp16318 := GHd(__e, V5654)

tmp16319 := GTl(__e, V5654)

tmp16320 := Call1(__e, SymFn(symshen_4freshterms), tmp16319)


return TailCall2(__e, SymFn(symadjoin), tmp16318, tmp16320)


} else {
tmp16323 := GConsP(__e, V5654)

if True == tmp16323 {
tmp16321 := GTl(__e, V5654)

if SymFn(symshen_4freshterms) == __self {
__self1 = tmp16321
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4freshterms), tmp16321)


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

tmp16336 := Call2(__e, ns2_1set, symshen_4freshterms, tmp16314)


_ = tmp16336

tmp16337 := MakeNative6(func(__e *ControlFlow, V5655, V5656, V5657, V5658, V5659, V5660 Obj) Obj {
tmp16351 := Call1(__e, SymFn(symshen_4unlocked_2), V5658)


var ifres16338 Obj

if True == tmp16351 {
tmp16339 := Call2(__e, SymFn(symshen_4lazyderef), V5655, V5657)


let__13709 := tmp16339
_ = let__13709

tmp16350 := GEq(__e, let__13709, Nil)

var ifres16340 Obj

if True == tmp16350 {
tmp16341 := Call2(__e, SymFn(symshen_4lazyderef), V5656, V5657)


let__13710 := tmp16341
_ = let__13710

tmp16342 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16343 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16343

return TailCall1(__e, SymFn(symthaw), V5660)


panic("unreachable")
})

let__13711 := tmp16342
_ = let__13711

tmp16349 := GEq(__e, let__13710, Nil)

var ifres16344 Obj

if True == tmp16349 {
tmp16345 := Call1(__e, SymFn(symthaw), let__13711)


ifres16344 = tmp16345


} else {
tmp16348 := Call1(__e, SymFn(symshen_4pvar_2), let__13710)


var ifres16346 Obj

if True == tmp16348 {
tmp16347 := Call4(__e, SymFn(symshen_4bind_b), let__13710, Nil, V5657, let__13711)


ifres16346 = tmp16347


} else {
ifres16346 = False


}

ifres16344 = ifres16346


}

ifres16340 = ifres16344


} else {
ifres16340 = False


}

ifres16338 = ifres16340


} else {
ifres16338 = False


}

let__13708 := ifres16338
_ = let__13708

tmp16448 := GEq(__e, let__13708, False)

if True == tmp16448 {
tmp16446 := Call1(__e, SymFn(symshen_4unlocked_2), V5658)


if True == tmp16446 {
tmp16352 := Call2(__e, SymFn(symshen_4lazyderef), V5655, V5657)


let__13712 := tmp16352
_ = let__13712

tmp16444 := GConsP(__e, let__13712)

if True == tmp16444 {
tmp16353 := GHd(__e, let__13712)

let__13713 := tmp16353
_ = let__13713

tmp16354 := GTl(__e, let__13712)

let__13714 := tmp16354
_ = let__13714

tmp16355 := Call2(__e, SymFn(symshen_4lazyderef), V5656, V5657)


let__13715 := tmp16355
_ = let__13715

tmp16356 := MakeNative1(func(__e *ControlFlow, Z5670 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5671 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5672 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5673 Obj) Obj {
tmp16357 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16357

tmp16358 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16359 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp16360 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall6(__e, SymFn(symshen_4p_1hyps), let__13714, Z5673, V5657, V5658, V5659, V5660)
panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), Z5671, tmp16359, V5657, V5658, V5659, tmp16360)


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), Z5670, let__13713, V5657, V5658, V5659, tmp16358)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

let__13716 := tmp16356
_ = let__13716

tmp16442 := GConsP(__e, let__13715)

if True == tmp16442 {
tmp16361 := GHd(__e, let__13715)

tmp16362 := Call2(__e, SymFn(symshen_4lazyderef), tmp16361, V5657)


let__13717 := tmp16362
_ = let__13717

tmp16363 := MakeNative1(func(__e *ControlFlow, Z5676 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5677 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5678 Obj) Obj {
tmp16364 := GTl(__e, let__13715)

let__13719 := tmp16364
_ = let__13719

tmp16365 := Call1(__e, let__13716, Z5676)


tmp16366 := Call1(__e, tmp16365, Z5677)


tmp16367 := Call1(__e, tmp16366, Z5678)


return TailCall1(__e, tmp16367, let__13719)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

let__13718 := tmp16363
_ = let__13718

tmp16422 := GConsP(__e, let__13717)

if True == tmp16422 {
tmp16368 := GHd(__e, let__13717)

let__13720 := tmp16368
_ = let__13720

tmp16369 := GTl(__e, let__13717)

tmp16370 := Call2(__e, SymFn(symshen_4lazyderef), tmp16369, V5657)


let__13721 := tmp16370
_ = let__13721

tmp16371 := MakeNative1(func(__e *ControlFlow, Z5683 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z5684 Obj) Obj {
tmp16372 := Call1(__e, let__13718, let__13720)


tmp16373 := Call1(__e, tmp16372, Z5683)


return TailCall1(__e, tmp16373, Z5684)


panic("unreachable")
})
panic("unreachable")
})

let__13722 := tmp16371
_ = let__13722

tmp16406 := GConsP(__e, let__13721)

if True == tmp16406 {
tmp16374 := GHd(__e, let__13721)

let__13723 := tmp16374
_ = let__13723

tmp16375 := GTl(__e, let__13721)

tmp16376 := Call2(__e, SymFn(symshen_4lazyderef), tmp16375, V5657)


let__13724 := tmp16376
_ = let__13724

tmp16377 := MakeNative1(func(__e *ControlFlow, Z5688 Obj) Obj {
tmp16378 := Call1(__e, let__13722, let__13723)


return TailCall1(__e, tmp16378, Z5688)


panic("unreachable")
})

let__13725 := tmp16377
_ = let__13725

tmp16394 := GConsP(__e, let__13724)

if True == tmp16394 {
tmp16379 := GHd(__e, let__13724)

let__13726 := tmp16379
_ = let__13726

tmp16380 := GTl(__e, let__13724)

tmp16381 := Call2(__e, SymFn(symshen_4lazyderef), tmp16380, V5657)


let__13727 := tmp16381
_ = let__13727

tmp16382 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13725, let__13726)
panic("unreachable")
})

let__13728 := tmp16382
_ = let__13728

tmp16386 := GEq(__e, let__13727, Nil)

if True == tmp16386 {
return TailCall1(__e, SymFn(symthaw), let__13728)
} else {
tmp16384 := Call1(__e, SymFn(symshen_4pvar_2), let__13727)


if True == tmp16384 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__13727, Nil, V5657, let__13728)
} else {
return False
}


}


} else {
tmp16392 := Call1(__e, SymFn(symshen_4pvar_2), let__13724)


if True == tmp16392 {
tmp16387 := Call1(__e, SymFn(symshen_4newpv), V5657)


let__13729 := tmp16387
_ = let__13729

tmp16388 := GCons(__e, let__13729, Nil)

tmp16389 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__13725, let__13729)
panic("unreachable")
})

tmp16390 := Call4(__e, SymFn(symshen_4bind_b), let__13724, tmp16388, V5657, tmp16389)


return TailCall2(__e, SymFn(symshen_4gc), V5657, tmp16390)


} else {
return False
}


}


} else {
tmp16404 := Call1(__e, SymFn(symshen_4pvar_2), let__13721)


if True == tmp16404 {
tmp16395 := Call1(__e, SymFn(symshen_4newpv), V5657)


let__13730 := tmp16395
_ = let__13730

tmp16396 := Call1(__e, SymFn(symshen_4newpv), V5657)


let__13731 := tmp16396
_ = let__13731

tmp16397 := GCons(__e, let__13731, Nil)

tmp16398 := GCons(__e, let__13730, tmp16397)

tmp16399 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16400 := Call1(__e, let__13722, let__13730)


return TailCall1(__e, tmp16400, let__13731)


panic("unreachable")
})

tmp16401 := Call4(__e, SymFn(symshen_4bind_b), let__13721, tmp16398, V5657, tmp16399)


tmp16402 := Call2(__e, SymFn(symshen_4gc), V5657, tmp16401)


return TailCall2(__e, SymFn(symshen_4gc), V5657, tmp16402)


} else {
return False
}


}


} else {
tmp16420 := Call1(__e, SymFn(symshen_4pvar_2), let__13717)


if True == tmp16420 {
tmp16407 := Call1(__e, SymFn(symshen_4newpv), V5657)


let__13732 := tmp16407
_ = let__13732

tmp16408 := Call1(__e, SymFn(symshen_4newpv), V5657)


let__13733 := tmp16408
_ = let__13733

tmp16409 := Call1(__e, SymFn(symshen_4newpv), V5657)


let__13734 := tmp16409
_ = let__13734

tmp16410 := GCons(__e, let__13734, Nil)

tmp16411 := GCons(__e, let__13733, tmp16410)

tmp16412 := GCons(__e, let__13732, tmp16411)

tmp16413 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16414 := Call1(__e, let__13718, let__13732)


tmp16415 := Call1(__e, tmp16414, let__13733)


return TailCall1(__e, tmp16415, let__13734)


panic("unreachable")
})

tmp16416 := Call4(__e, SymFn(symshen_4bind_b), let__13717, tmp16412, V5657, tmp16413)


tmp16417 := Call2(__e, SymFn(symshen_4gc), V5657, tmp16416)


tmp16418 := Call2(__e, SymFn(symshen_4gc), V5657, tmp16417)


return TailCall2(__e, SymFn(symshen_4gc), V5657, tmp16418)


} else {
return False
}


}


} else {
tmp16440 := Call1(__e, SymFn(symshen_4pvar_2), let__13715)


if True == tmp16440 {
tmp16423 := Call1(__e, SymFn(symshen_4newpv), V5657)


let__13735 := tmp16423
_ = let__13735

tmp16424 := Call1(__e, SymFn(symshen_4newpv), V5657)


let__13736 := tmp16424
_ = let__13736

tmp16425 := Call1(__e, SymFn(symshen_4newpv), V5657)


let__13737 := tmp16425
_ = let__13737

tmp16426 := Call1(__e, SymFn(symshen_4newpv), V5657)


let__13738 := tmp16426
_ = let__13738

tmp16427 := GCons(__e, let__13737, Nil)

tmp16428 := GCons(__e, let__13736, tmp16427)

tmp16429 := GCons(__e, let__13735, tmp16428)

tmp16430 := GCons(__e, tmp16429, let__13738)

tmp16431 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16432 := Call1(__e, let__13716, let__13735)


tmp16433 := Call1(__e, tmp16432, let__13736)


tmp16434 := Call1(__e, tmp16433, let__13737)


return TailCall1(__e, tmp16434, let__13738)


panic("unreachable")
})

tmp16435 := Call4(__e, SymFn(symshen_4bind_b), let__13715, tmp16430, V5657, tmp16431)


tmp16436 := Call2(__e, SymFn(symshen_4gc), V5657, tmp16435)


tmp16437 := Call2(__e, SymFn(symshen_4gc), V5657, tmp16436)


tmp16438 := Call2(__e, SymFn(symshen_4gc), V5657, tmp16437)


return TailCall2(__e, SymFn(symshen_4gc), V5657, tmp16438)


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
return let__13708
}


panic("unreachable")
})

tmp16449 := Call2(__e, ns2_1set, symshen_4p_1hyps, tmp16337)


_ = tmp16449

tmp16450 := MakeNative(func(__e *ControlFlow) {
V5702 := __e.Get(1)
V5703 := __e.Get(2)
V5704 := __e.Get(3)
V5705 := __e.Get(4)
V5706 := __e.Get(5)
V5707 := __e.Get(6)
V5708 := __e.Get(7)
__e.ReturnDirect(func() Obj {
let__13739 := GAdd(__e, V5707, MakeInteger(1))
_ = let__13739

tmp16488 := Call1(__e, SymFn(symshen_4unlocked_2), V5706)


var ifres16452 Obj

if True == tmp16488 {
tmp16453 := Call2(__e, SymFn(symshen_4lazyderef), V5702, V5705)


let__13741 := tmp16453
_ = let__13741

tmp16487 := GConsP(__e, let__13741)

var ifres16454 Obj

if True == tmp16487 {
tmp16455 := GHd(__e, let__13741)

tmp16456 := Call2(__e, SymFn(symshen_4lazyderef), tmp16455, V5705)


let__13742 := tmp16456
_ = let__13742

tmp16486 := GEq(__e, let__13742, symwhere)

var ifres16457 Obj

if True == tmp16486 {
tmp16458 := GTl(__e, let__13741)

tmp16459 := Call2(__e, SymFn(symshen_4lazyderef), tmp16458, V5705)


let__13743 := tmp16459
_ = let__13743

tmp16485 := GConsP(__e, let__13743)

var ifres16460 Obj

if True == tmp16485 {
tmp16461 := GHd(__e, let__13743)

let__13744 := tmp16461
_ = let__13744

tmp16462 := GTl(__e, let__13743)

tmp16463 := Call2(__e, SymFn(symshen_4lazyderef), tmp16462, V5705)


let__13745 := tmp16463
_ = let__13745

tmp16484 := GConsP(__e, let__13745)

var ifres16464 Obj

if True == tmp16484 {
tmp16465 := GHd(__e, let__13745)

let__13746 := tmp16465
_ = let__13746

tmp16466 := GTl(__e, let__13745)

tmp16467 := Call2(__e, SymFn(symshen_4lazyderef), tmp16466, V5705)


let__13747 := tmp16467
_ = let__13747

tmp16483 := GEq(__e, let__13747, Nil)

var ifres16468 Obj

if True == tmp16483 {
tmp16469 := Call1(__e, SymFn(symshen_4newpv), V5705)


let__13748 := tmp16469
_ = let__13748

tmp16470 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16470

tmp16471 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16472 := Call1(__e, SymFn(symshen_4curry), let__13744)


tmp16473 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16474 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16475 := MakeNative0(func(__e *ControlFlow) Obj {
tmp16476 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp16477 := GCons(__e, symverified, Nil)

tmp16478 := GCons(__e, tmp16476, tmp16477)

tmp16479 := GCons(__e, let__13748, tmp16478)

tmp16480 := GCons(__e, tmp16479, V5704)

__e.TailApply(SymFn(symshen_4t_d_1correct), let__13746, V5703, tmp16480, V5705, V5706, let__13739, V5708)
return nil


panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4cut), V5705, V5706, let__13739, tmp16475)


panic("unreachable")
})

__e.TailApply(SymFn(symshen_4system_1S_1h), let__13748, symboolean, V5704, V5705, V5706, let__13739, tmp16474)
return nil


panic("unreachable")
})

return TailCall6(__e, SymFn(symbind), let__13748, tmp16472, V5705, V5706, let__13739, tmp16473)


panic("unreachable")
})

tmp16481 := Call4(__e, SymFn(symshen_4cut), V5705, V5706, let__13739, tmp16471)


tmp16482 := Call2(__e, SymFn(symshen_4gc), V5705, tmp16481)


ifres16468 = tmp16482


} else {
ifres16468 = False


}

ifres16464 = ifres16468


} else {
ifres16464 = False


}

ifres16460 = ifres16464


} else {
ifres16460 = False


}

ifres16457 = ifres16460


} else {
ifres16457 = False


}

ifres16454 = ifres16457


} else {
ifres16454 = False


}

ifres16452 = ifres16454


} else {
ifres16452 = False


}

let__13740 := ifres16452
_ = let__13740

tmp16497 := GEq(__e, let__13740, False)

if True == tmp16497 {
tmp16493 := Call1(__e, SymFn(symshen_4unlocked_2), V5706)


var ifres16489 Obj

if True == tmp16493 {
tmp16490 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16490

tmp16491 := Call1(__e, SymFn(symshen_4curry), V5702)


tmp16492 := Call(__e, SymFn(symshen_4system_1S_1h), tmp16491, V5703, V5704, V5705, V5706, let__13739, V5708)


ifres16489 = tmp16492


} else {
ifres16489 = False


}

let__13749 := ifres16489
_ = let__13749

tmp16495 := GEq(__e, let__13749, False)

if True == tmp16495 {
return TailCall2(__e, SymFn(symshen_4unlock), V5706, let__13739)
} else {
return let__13749
}


} else {
return let__13740
}


panic("unreachable")
}())
}, 7)

tmp16498 := Call2(__e, ns2_1set, symshen_4t_d_1correct, tmp16450)


_ = tmp16498

tmp16499 := MakeNative(func(__e *ControlFlow) {
V5720 := __e.Get(1)
V5721 := __e.Get(2)
V5722 := __e.Get(3)
V5723 := __e.Get(4)
V5724 := __e.Get(5)
V5725 := __e.Get(6)
V5726 := __e.Get(7)
V5727 := __e.Get(8)
__e.ReturnDirect(func() Obj {
tmp16506 := Call1(__e, SymFn(symshen_4unlocked_2), V5725)


var ifres16500 Obj

if True == tmp16506 {
tmp16501 := Call2(__e, SymFn(symshen_4lazyderef), V5720, V5724)


let__13751 := tmp16501
_ = let__13751

tmp16505 := GEq(__e, let__13751, Nil)

var ifres16502 Obj

if True == tmp16505 {
tmp16503 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16503

tmp16504 := Call6(__e, SymFn(symis_b), V5721, V5723, V5724, V5725, V5726, V5727)


ifres16502 = tmp16504


} else {
ifres16502 = False


}

ifres16500 = ifres16502


} else {
ifres16500 = False


}

let__13750 := ifres16500
_ = let__13750

tmp16538 := GEq(__e, let__13750, False)

if True == tmp16538 {
tmp16536 := Call1(__e, SymFn(symshen_4unlocked_2), V5725)


if True == tmp16536 {
tmp16507 := Call2(__e, SymFn(symshen_4lazyderef), V5720, V5724)


let__13752 := tmp16507
_ = let__13752

tmp16534 := GConsP(__e, let__13752)

if True == tmp16534 {
tmp16508 := GHd(__e, let__13752)

let__13753 := tmp16508
_ = let__13753

tmp16509 := GTl(__e, let__13752)

let__13754 := tmp16509
_ = let__13754

tmp16510 := Call2(__e, SymFn(symshen_4lazyderef), V5721, V5724)


let__13755 := tmp16510
_ = let__13755

tmp16532 := GConsP(__e, let__13755)

if True == tmp16532 {
tmp16511 := GHd(__e, let__13755)

let__13756 := tmp16511
_ = let__13756

tmp16512 := GTl(__e, let__13755)

tmp16513 := Call2(__e, SymFn(symshen_4lazyderef), tmp16512, V5724)


let__13757 := tmp16513
_ = let__13757

tmp16530 := GConsP(__e, let__13757)

if True == tmp16530 {
tmp16514 := GHd(__e, let__13757)

tmp16515 := Call2(__e, SymFn(symshen_4lazyderef), tmp16514, V5724)


let__13758 := tmp16515
_ = let__13758

tmp16528 := GEq(__e, let__13758, sym_1_1_6)

if True == tmp16528 {
tmp16516 := GTl(__e, let__13757)

tmp16517 := Call2(__e, SymFn(symshen_4lazyderef), tmp16516, V5724)


let__13759 := tmp16517
_ = let__13759

tmp16526 := GConsP(__e, let__13759)

if True == tmp16526 {
tmp16518 := GHd(__e, let__13759)

let__13760 := tmp16518
_ = let__13760

tmp16519 := GTl(__e, let__13759)

tmp16520 := Call2(__e, SymFn(symshen_4lazyderef), tmp16519, V5724)


let__13761 := tmp16520
_ = let__13761

tmp16524 := GEq(__e, let__13761, Nil)

if True == tmp16524 {
tmp16521 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp16521

tmp16522 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4t_d_1integrity), let__13754, let__13760, V5722, V5723, V5724, V5725, V5726, V5727)
return nil
panic("unreachable")
})

__e.TailApply(SymFn(symshen_4system_1S_1h), let__13753, let__13756, V5722, V5724, V5725, V5726, tmp16522)
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
return let__13750
}


panic("unreachable")
}())
}, 8)

tmp16539 := Call2(__e, ns2_1set, symshen_4t_d_1integrity, tmp16499)


_ = tmp16539

tmp16540 := MakeNative1(func(__e *ControlFlow, V5740 Obj) Obj {
tmp16546 := GVectorP(__e, V5740)

if True == tmp16546 {
tmp16543 := PrimGuard1(__e, symstring_2, PrimIsString, V5740)

if True == GNot(__e, tmp16543) {
tmp16541 := GVecGet(__e, V5740, MakeInteger(0))

return GEq(__e, tmp16541, symshen_4print_1freshterm)


} else {
return False
}


} else {
return False
}


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4freshterm_2, tmp16540)




panic("unreachable")
})

