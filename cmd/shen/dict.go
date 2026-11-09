package main

import . "github.com/pyrex41/shen-go/kl"

var DictMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp17912 := MakeNative1(func(__e *ControlFlow, V4220 Obj) Obj {
if True == GLt(__e, V4220, MakeInteger(1)) {
tmp17913 := Call3(__e, SymFn(symshen_4app), V4220, MakeString(""), symshen_4s)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("invalid initial dict size: "), tmp17913))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("invalid initial dict size: "), tmp17913)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


} else {
tmp17916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(GAdd(__e, MakeInteger(3), V4220))
}
__typedArg0 := GAdd(__e, MakeInteger(3), V4220)
return Call1(__e, SymFn(symabsvector), __typedArg0)
})()

let__17890 := tmp17916
_ = let__17890

tmp17917 := GVecSet(__e, let__17890, MakeInteger(0), symshen_4dictionary)

let__17891 := tmp17917
_ = let__17891

tmp17918 := GVecSet(__e, let__17890, MakeInteger(1), V4220)

let__17892 := tmp17918
_ = let__17892

tmp17919 := GVecSet(__e, let__17890, MakeInteger(2), MakeInteger(0))

let__17893 := tmp17919
_ = let__17893

tmp17921 := Call4(__e, SymFn(symshen_4fillvector), let__17890, MakeInteger(3), GAdd(__e, MakeInteger(2), V4220), Nil)


let__17894 := tmp17921
_ = let__17894

return let__17890


}


panic("unreachable")
})

tmp17924 := Call2(__e, ns2_1set, symshen_4dict, tmp17912)


_ = tmp17924

tmp17925 := MakeNative1(func(__e *ControlFlow, V4226 Obj) Obj {
tmp17930 := GVectorP(__e, V4226)

if True == tmp17930 {
tmp17926 := MakeNative0(func(__e *ControlFlow) Obj {
return GVecGet(__e, V4226, MakeInteger(0))
panic("unreachable")
})

tmp17927 := MakeNative1(func(__e *ControlFlow, Z4227 Obj) Obj {
return symshen_4not_1dictionary
panic("unreachable")
})

tmp17928 := Call2(__e, try_1catch, tmp17926, tmp17927)


return GEq(__e, tmp17928, symshen_4dictionary)


} else {
return False
}


panic("unreachable")
})

tmp17931 := Call2(__e, ns2_1set, symshen_4dict_2, tmp17925)


_ = tmp17931

tmp17932 := MakeNative1(func(__e *ControlFlow, V4228 Obj) Obj {
return GVecGet(__e, V4228, MakeInteger(1))
panic("unreachable")
})

tmp17933 := Call2(__e, ns2_1set, symshen_4dict_1capacity, tmp17932)


_ = tmp17933

tmp17934 := MakeNative1(func(__e *ControlFlow, V4229 Obj) Obj {
return GVecGet(__e, V4229, MakeInteger(2))
panic("unreachable")
})

tmp17935 := Call2(__e, ns2_1set, symshen_4dict_1count, tmp17934)


_ = tmp17935

tmp17936 := MakeNative2(func(__e *ControlFlow, V4230, V4231 Obj) Obj {
return GVecSet(__e, V4230, MakeInteger(2), V4231)
panic("unreachable")
})

tmp17937 := Call2(__e, ns2_1set, symshen_4dict_1count_1_6, tmp17936)


_ = tmp17937

tmp17938 := MakeNative2(func(__e *ControlFlow, V4232, V4233 Obj) Obj {
return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V4232, GAdd(__e, MakeInteger(3), V4233))
}
__typedArg0 := V4232
__typedArg1 := GAdd(__e, MakeInteger(3), V4233)
return Call2(__e, SymFn(sym_5_1address), __typedArg0, __typedArg1)
})()


panic("unreachable")
})

tmp17940 := Call2(__e, ns2_1set, symshen_4_5_1dict_1bucket, tmp17938)


_ = tmp17940

tmp17941 := MakeNative3(func(__e *ControlFlow, V4234, V4235, V4236 Obj) Obj {
return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V4234, GAdd(__e, MakeInteger(3), V4235), V4236)
}
__typedArg0 := V4234
__typedArg1 := GAdd(__e, MakeInteger(3), V4235)
__typedArg2 := V4236
return Call3(__e, SymFn(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()


panic("unreachable")
})

tmp17943 := Call2(__e, ns2_1set, symshen_4dict_1bucket_1_6, tmp17941)


_ = tmp17943

tmp17944 := MakeNative3(func(__e *ControlFlow, V4237, V4238, V4239 Obj) Obj {
tmp17945 := Call1(__e, SymFn(symlength), V4239)


tmp17946 := Call1(__e, SymFn(symlength), V4238)


let__17895 := GSub(__e, tmp17945, tmp17946)
_ = let__17895

tmp17948 := Call1(__e, SymFn(symshen_4dict_1count), V4237)


return TailCall2(__e, SymFn(symshen_4dict_1count_1_6), V4237, GAdd(__e, let__17895, tmp17948))


panic("unreachable")
})

tmp17950 := Call2(__e, ns2_1set, symshen_4dict_1update_1count, tmp17944)


_ = tmp17950

tmp17951 := MakeNative3(func(__e *ControlFlow, V4241, V4242, V4243 Obj) Obj {
tmp17952 := Call1(__e, SymFn(symshen_4dict_1capacity), V4241)


tmp17953 := Call2(__e, SymFn(symhash), V4242, tmp17952)


let__17896 := tmp17953
_ = let__17896

tmp17954 := Call2(__e, SymFn(symshen_4_5_1dict_1bucket), V4241, let__17896)


let__17897 := tmp17954
_ = let__17897

tmp17955 := Call3(__e, SymFn(symshen_4assoc_1set), V4242, V4243, let__17897)


let__17898 := tmp17955
_ = let__17898

tmp17956 := Call3(__e, SymFn(symshen_4dict_1bucket_1_6), V4241, let__17896, let__17898)


let__17899 := tmp17956
_ = let__17899

tmp17957 := Call3(__e, SymFn(symshen_4dict_1update_1count), V4241, let__17897, let__17898)


let__17900 := tmp17957
_ = let__17900

return V4243


panic("unreachable")
})

tmp17958 := Call2(__e, ns2_1set, symshen_4dict_1_6, tmp17951)


_ = tmp17958

tmp17959 := MakeNative2(func(__e *ControlFlow, V4249, V4250 Obj) Obj {
tmp17960 := Call1(__e, SymFn(symshen_4dict_1capacity), V4249)


tmp17961 := Call2(__e, SymFn(symhash), V4250, tmp17960)


let__17901 := tmp17961
_ = let__17901

tmp17962 := Call2(__e, SymFn(symshen_4_5_1dict_1bucket), V4249, let__17901)


let__17902 := tmp17962
_ = let__17902

tmp17963 := Call2(__e, SymFn(symassoc), V4250, let__17902)


let__17903 := tmp17963
_ = let__17903

tmp17967 := Call1(__e, SymFn(symempty_2), let__17903)


if True == tmp17967 {
tmp17964 := Call3(__e, SymFn(symshen_4app), V4250, MakeString(" not found in dict\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("value "), tmp17964))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("value "), tmp17964)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


} else {
return GTl(__e, let__17903)
}


panic("unreachable")
})

tmp17968 := Call2(__e, ns2_1set, symshen_4_5_1dict, tmp17959)


_ = tmp17968

tmp17969 := MakeNative2(func(__e *ControlFlow, V4254, V4255 Obj) Obj {
tmp17970 := Call1(__e, SymFn(symshen_4dict_1capacity), V4254)


tmp17971 := Call2(__e, SymFn(symhash), V4255, tmp17970)


let__17904 := tmp17971
_ = let__17904

tmp17972 := Call2(__e, SymFn(symshen_4_5_1dict_1bucket), V4254, let__17904)


let__17905 := tmp17972
_ = let__17905

tmp17973 := Call2(__e, SymFn(symshen_4assoc_1rm), V4255, let__17905)


let__17906 := tmp17973
_ = let__17906

tmp17974 := Call3(__e, SymFn(symshen_4dict_1bucket_1_6), V4254, let__17904, let__17906)


let__17907 := tmp17974
_ = let__17907

tmp17975 := Call3(__e, SymFn(symshen_4dict_1update_1count), V4254, let__17905, let__17906)


let__17908 := tmp17975
_ = let__17908

return V4255


panic("unreachable")
})

tmp17976 := Call2(__e, ns2_1set, symshen_4dict_1rm, tmp17969)


_ = tmp17976

tmp17977 := MakeNative3(func(__e *ControlFlow, V4261, V4262, V4263 Obj) Obj {
tmp17978 := Call1(__e, SymFn(symshen_4dict_1capacity), V4262)


let__17909 := tmp17978
_ = let__17909

return TailCall5(__e, SymFn(symshen_4dict_1fold_1h), V4261, V4262, V4263, MakeInteger(0), let__17909)


panic("unreachable")
})

tmp17979 := Call2(__e, ns2_1set, symshen_4dict_1fold, tmp17977)


_ = tmp17979

tmp17980 := (func() Obj {
var __self Obj
__self = MakeNative5(func(__e *ControlFlow, __self1, __self2, __self3, __self4, __self5 Obj) Obj {
__selftop:
V4266 := __self1
_ = V4266
V4267 := __self2
_ = V4267
V4268 := __self3
_ = V4268
V4269 := __self4
_ = V4269
V4270 := __self5
_ = V4270
tmp17985 := GEq(__e, V4269, V4270)

if True == tmp17985 {
return V4268
} else {
tmp17981 := Call2(__e, SymFn(symshen_4_5_1dict_1bucket), V4267, V4269)


let__17910 := tmp17981
_ = let__17910

tmp17982 := Call3(__e, SymFn(symshen_4bucket_1fold), V4266, let__17910, V4268)


let__17911 := tmp17982
_ = let__17911

if SymFn(symshen_4dict_1fold_1h) == __self {
__self1, __self2, __self3, __self4, __self5 = V4266, V4267, let__17911, GAdd(__e, MakeInteger(1), V4269), V4270
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4dict_1fold_1h), V4266, V4267, let__17911, GAdd(__e, MakeInteger(1), V4269), V4270)


}


panic("unreachable")
})
return __self
})()

tmp17986 := Call2(__e, ns2_1set, symshen_4dict_1fold_1h, tmp17980)


_ = tmp17986

tmp17987 := MakeNative3(func(__e *ControlFlow, V4273, V4274, V4275 Obj) Obj {
tmp18002 := GEq(__e, Nil, V4274)

if True == tmp18002 {
return V4275
} else {
tmp18000 := GConsP(__e, V4274)

var ifres17997 Obj

if True == tmp18000 {
tmp17998 := GHd(__e, V4274)

tmp17999 := GConsP(__e, tmp17998)

ifres17997 = tmp17999


} else {
ifres17997 = False


}

if True == ifres17997 {
tmp17988 := GHd(__e, V4274)

tmp17989 := GHd(__e, tmp17988)

tmp17990 := Call1(__e, V4273, tmp17989)


tmp17991 := GHd(__e, V4274)

tmp17992 := GTl(__e, tmp17991)

tmp17993 := Call1(__e, tmp17990, tmp17992)


tmp17994 := GTl(__e, V4274)

tmp17995 := Call3(__e, SymFn(symshen_4bucket_1fold), V4273, tmp17994, V4275)


return TailCall1(__e, tmp17993, tmp17995)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4bucket_1fold)
}


}


panic("unreachable")
})

tmp18003 := Call2(__e, ns2_1set, symshen_4bucket_1fold, tmp17987)


_ = tmp18003

tmp18004 := MakeNative1(func(__e *ControlFlow, V4276 Obj) Obj {
tmp18005 := MakeNative1(func(__e *ControlFlow, Z4277 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z4278 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z4279 Obj) Obj {
return GCons(__e, Z4277, Z4279)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

return TailCall3(__e, SymFn(symshen_4dict_1fold), tmp18005, V4276, Nil)


panic("unreachable")
})

tmp18006 := Call2(__e, ns2_1set, symshen_4dict_1keys, tmp18004)


_ = tmp18006

tmp18007 := MakeNative1(func(__e *ControlFlow, V4280 Obj) Obj {
tmp18008 := MakeNative1(func(__e *ControlFlow, Z4281 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z4282 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z4283 Obj) Obj {
return GCons(__e, Z4282, Z4283)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

return TailCall3(__e, SymFn(symshen_4dict_1fold), tmp18008, V4280, Nil)


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4dict_1values, tmp18007)




panic("unreachable")
})

