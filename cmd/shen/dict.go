package main

import . "github.com/pyrex41/shen-go/kl"

var DictMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp17889 := MakeNative1(func(__e *ControlFlow, V4205 Obj) Obj {
if True == GLt(__e, V4205, MakeInteger(1)) {
tmp17890 := Call3(__e, SymFn(symshen_4app), V4205, MakeString(""), symshen_4s)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("invalid initial dict size: "), tmp17890))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("invalid initial dict size: "), tmp17890)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


} else {
tmp17893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(GAdd(__e, MakeInteger(3), V4205))
}
__typedArg0 := GAdd(__e, MakeInteger(3), V4205)
return Call1(__e, SymFn(symabsvector), __typedArg0)
})()

let__17867 := tmp17893
_ = let__17867

tmp17894 := GVecSet(__e, let__17867, MakeInteger(0), symshen_4dictionary)

let__17868 := tmp17894
_ = let__17868

tmp17895 := GVecSet(__e, let__17867, MakeInteger(1), V4205)

let__17869 := tmp17895
_ = let__17869

tmp17896 := GVecSet(__e, let__17867, MakeInteger(2), MakeInteger(0))

let__17870 := tmp17896
_ = let__17870

tmp17898 := Call4(__e, SymFn(symshen_4fillvector), let__17867, MakeInteger(3), GAdd(__e, MakeInteger(2), V4205), Nil)


let__17871 := tmp17898
_ = let__17871

return let__17867


}


panic("unreachable")
})

tmp17901 := Call2(__e, ns2_1set, symshen_4dict, tmp17889)


_ = tmp17901

tmp17902 := MakeNative1(func(__e *ControlFlow, V4211 Obj) Obj {
tmp17907 := GVectorP(__e, V4211)

if True == tmp17907 {
tmp17903 := MakeNative0(func(__e *ControlFlow) Obj {
return GVecGet(__e, V4211, MakeInteger(0))
panic("unreachable")
})

tmp17904 := MakeNative1(func(__e *ControlFlow, Z4212 Obj) Obj {
return symshen_4not_1dictionary
panic("unreachable")
})

tmp17905 := Call2(__e, try_1catch, tmp17903, tmp17904)


return GEq(__e, tmp17905, symshen_4dictionary)


} else {
return False
}


panic("unreachable")
})

tmp17908 := Call2(__e, ns2_1set, symshen_4dict_2, tmp17902)


_ = tmp17908

tmp17909 := MakeNative1(func(__e *ControlFlow, V4213 Obj) Obj {
return GVecGet(__e, V4213, MakeInteger(1))
panic("unreachable")
})

tmp17910 := Call2(__e, ns2_1set, symshen_4dict_1capacity, tmp17909)


_ = tmp17910

tmp17911 := MakeNative1(func(__e *ControlFlow, V4214 Obj) Obj {
return GVecGet(__e, V4214, MakeInteger(2))
panic("unreachable")
})

tmp17912 := Call2(__e, ns2_1set, symshen_4dict_1count, tmp17911)


_ = tmp17912

tmp17913 := MakeNative2(func(__e *ControlFlow, V4215, V4216 Obj) Obj {
return GVecSet(__e, V4215, MakeInteger(2), V4216)
panic("unreachable")
})

tmp17914 := Call2(__e, ns2_1set, symshen_4dict_1count_1_6, tmp17913)


_ = tmp17914

tmp17915 := MakeNative2(func(__e *ControlFlow, V4217, V4218 Obj) Obj {
return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V4217, GAdd(__e, MakeInteger(3), V4218))
}
__typedArg0 := V4217
__typedArg1 := GAdd(__e, MakeInteger(3), V4218)
return Call2(__e, SymFn(sym_5_1address), __typedArg0, __typedArg1)
})()


panic("unreachable")
})

tmp17917 := Call2(__e, ns2_1set, symshen_4_5_1dict_1bucket, tmp17915)


_ = tmp17917

tmp17918 := MakeNative3(func(__e *ControlFlow, V4219, V4220, V4221 Obj) Obj {
return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V4219, GAdd(__e, MakeInteger(3), V4220), V4221)
}
__typedArg0 := V4219
__typedArg1 := GAdd(__e, MakeInteger(3), V4220)
__typedArg2 := V4221
return Call3(__e, SymFn(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()


panic("unreachable")
})

tmp17920 := Call2(__e, ns2_1set, symshen_4dict_1bucket_1_6, tmp17918)


_ = tmp17920

tmp17921 := MakeNative3(func(__e *ControlFlow, V4222, V4223, V4224 Obj) Obj {
tmp17922 := Call1(__e, SymFn(symlength), V4224)


tmp17923 := Call1(__e, SymFn(symlength), V4223)


let__17872 := GSub(__e, tmp17922, tmp17923)
_ = let__17872

tmp17925 := Call1(__e, SymFn(symshen_4dict_1count), V4222)


return TailCall2(__e, SymFn(symshen_4dict_1count_1_6), V4222, GAdd(__e, let__17872, tmp17925))


panic("unreachable")
})

tmp17927 := Call2(__e, ns2_1set, symshen_4dict_1update_1count, tmp17921)


_ = tmp17927

tmp17928 := MakeNative3(func(__e *ControlFlow, V4226, V4227, V4228 Obj) Obj {
tmp17929 := Call1(__e, SymFn(symshen_4dict_1capacity), V4226)


tmp17930 := Call2(__e, SymFn(symhash), V4227, tmp17929)


let__17873 := tmp17930
_ = let__17873

tmp17931 := Call2(__e, SymFn(symshen_4_5_1dict_1bucket), V4226, let__17873)


let__17874 := tmp17931
_ = let__17874

tmp17932 := Call3(__e, SymFn(symshen_4assoc_1set), V4227, V4228, let__17874)


let__17875 := tmp17932
_ = let__17875

tmp17933 := Call3(__e, SymFn(symshen_4dict_1bucket_1_6), V4226, let__17873, let__17875)


let__17876 := tmp17933
_ = let__17876

tmp17934 := Call3(__e, SymFn(symshen_4dict_1update_1count), V4226, let__17874, let__17875)


let__17877 := tmp17934
_ = let__17877

return V4228


panic("unreachable")
})

tmp17935 := Call2(__e, ns2_1set, symshen_4dict_1_6, tmp17928)


_ = tmp17935

tmp17936 := MakeNative2(func(__e *ControlFlow, V4234, V4235 Obj) Obj {
tmp17937 := Call1(__e, SymFn(symshen_4dict_1capacity), V4234)


tmp17938 := Call2(__e, SymFn(symhash), V4235, tmp17937)


let__17878 := tmp17938
_ = let__17878

tmp17939 := Call2(__e, SymFn(symshen_4_5_1dict_1bucket), V4234, let__17878)


let__17879 := tmp17939
_ = let__17879

tmp17940 := Call2(__e, SymFn(symassoc), V4235, let__17879)


let__17880 := tmp17940
_ = let__17880

tmp17944 := Call1(__e, SymFn(symempty_2), let__17880)


if True == tmp17944 {
tmp17941 := Call3(__e, SymFn(symshen_4app), V4235, MakeString(" not found in dict\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("value "), tmp17941))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("value "), tmp17941)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


} else {
return GTl(__e, let__17880)
}


panic("unreachable")
})

tmp17945 := Call2(__e, ns2_1set, symshen_4_5_1dict, tmp17936)


_ = tmp17945

tmp17946 := MakeNative2(func(__e *ControlFlow, V4239, V4240 Obj) Obj {
tmp17947 := Call1(__e, SymFn(symshen_4dict_1capacity), V4239)


tmp17948 := Call2(__e, SymFn(symhash), V4240, tmp17947)


let__17881 := tmp17948
_ = let__17881

tmp17949 := Call2(__e, SymFn(symshen_4_5_1dict_1bucket), V4239, let__17881)


let__17882 := tmp17949
_ = let__17882

tmp17950 := Call2(__e, SymFn(symshen_4assoc_1rm), V4240, let__17882)


let__17883 := tmp17950
_ = let__17883

tmp17951 := Call3(__e, SymFn(symshen_4dict_1bucket_1_6), V4239, let__17881, let__17883)


let__17884 := tmp17951
_ = let__17884

tmp17952 := Call3(__e, SymFn(symshen_4dict_1update_1count), V4239, let__17882, let__17883)


let__17885 := tmp17952
_ = let__17885

return V4240


panic("unreachable")
})

tmp17953 := Call2(__e, ns2_1set, symshen_4dict_1rm, tmp17946)


_ = tmp17953

tmp17954 := MakeNative3(func(__e *ControlFlow, V4246, V4247, V4248 Obj) Obj {
tmp17955 := Call1(__e, SymFn(symshen_4dict_1capacity), V4247)


let__17886 := tmp17955
_ = let__17886

return TailCall5(__e, SymFn(symshen_4dict_1fold_1h), V4246, V4247, V4248, MakeInteger(0), let__17886)


panic("unreachable")
})

tmp17956 := Call2(__e, ns2_1set, symshen_4dict_1fold, tmp17954)


_ = tmp17956

tmp17957 := (func() Obj {
var __self Obj
__self = MakeNative5(func(__e *ControlFlow, __self1, __self2, __self3, __self4, __self5 Obj) Obj {
__selftop:
V4251 := __self1
_ = V4251
V4252 := __self2
_ = V4252
V4253 := __self3
_ = V4253
V4254 := __self4
_ = V4254
V4255 := __self5
_ = V4255
tmp17962 := GEq(__e, V4254, V4255)

if True == tmp17962 {
return V4253
} else {
tmp17958 := Call2(__e, SymFn(symshen_4_5_1dict_1bucket), V4252, V4254)


let__17887 := tmp17958
_ = let__17887

tmp17959 := Call3(__e, SymFn(symshen_4bucket_1fold), V4251, let__17887, V4253)


let__17888 := tmp17959
_ = let__17888

if SymFn(symshen_4dict_1fold_1h) == __self {
__self1, __self2, __self3, __self4, __self5 = V4251, V4252, let__17888, GAdd(__e, MakeInteger(1), V4254), V4255
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4dict_1fold_1h), V4251, V4252, let__17888, GAdd(__e, MakeInteger(1), V4254), V4255)


}


panic("unreachable")
})
return __self
})()

tmp17963 := Call2(__e, ns2_1set, symshen_4dict_1fold_1h, tmp17957)


_ = tmp17963

tmp17964 := MakeNative3(func(__e *ControlFlow, V4258, V4259, V4260 Obj) Obj {
tmp17979 := GEq(__e, Nil, V4259)

if True == tmp17979 {
return V4260
} else {
tmp17977 := GConsP(__e, V4259)

var ifres17974 Obj

if True == tmp17977 {
tmp17975 := GHd(__e, V4259)

tmp17976 := GConsP(__e, tmp17975)

ifres17974 = tmp17976


} else {
ifres17974 = False


}

if True == ifres17974 {
tmp17965 := GHd(__e, V4259)

tmp17966 := GHd(__e, tmp17965)

tmp17967 := Call1(__e, V4258, tmp17966)


tmp17968 := GHd(__e, V4259)

tmp17969 := GTl(__e, tmp17968)

tmp17970 := Call1(__e, tmp17967, tmp17969)


tmp17971 := GTl(__e, V4259)

tmp17972 := Call3(__e, SymFn(symshen_4bucket_1fold), V4258, tmp17971, V4260)


return TailCall1(__e, tmp17970, tmp17972)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4bucket_1fold)
}


}


panic("unreachable")
})

tmp17980 := Call2(__e, ns2_1set, symshen_4bucket_1fold, tmp17964)


_ = tmp17980

tmp17981 := MakeNative1(func(__e *ControlFlow, V4261 Obj) Obj {
tmp17982 := MakeNative1(func(__e *ControlFlow, Z4262 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z4263 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z4264 Obj) Obj {
return GCons(__e, Z4262, Z4264)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

return TailCall3(__e, SymFn(symshen_4dict_1fold), tmp17982, V4261, Nil)


panic("unreachable")
})

tmp17983 := Call2(__e, ns2_1set, symshen_4dict_1keys, tmp17981)


_ = tmp17983

tmp17984 := MakeNative1(func(__e *ControlFlow, V4265 Obj) Obj {
tmp17985 := MakeNative1(func(__e *ControlFlow, Z4266 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z4267 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z4268 Obj) Obj {
return GCons(__e, Z4267, Z4268)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

return TailCall3(__e, SymFn(symshen_4dict_1fold), tmp17985, V4265, Nil)


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4dict_1values, tmp17984)




panic("unreachable")
})

