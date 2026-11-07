package main

import . "github.com/pyrex41/shen-go/kl"

var DictMain = MakeNative(func(__e *ControlFlow) {
tmp18829 := MakeNative(func(__e *ControlFlow) {
V4220 := __e.Get(1)
_ = V4220
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5) {
__typedN0, __typedOK0 := TypedFloat64(V4220)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5) {
return TypedMaterializeBoolean((__typedN0 < __typedN1))
}}
__typedArg0 := V4220
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5), __typedArg0, __typedArg1)
})() {
tmp18830 := Call(__e, PrimFunc(symshen_4app), V4220, MakeString(""), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("invalid initial dict size: "))
__typedS1, __typedOK1 := TypedString(tmp18830)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("invalid initial dict size: ")
__typedArg1 := tmp18830
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("invalid initial dict size: "))
__typedS1, __typedOK1 := TypedString(tmp18830)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("invalid initial dict size: ")
__typedArg1 := tmp18830
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
tmp18833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(3))
__typedN1, __typedOK1 := TypedFloat64(V4220)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(3)
__typedArg1 := V4220
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(3))
__typedN1, __typedOK1 := TypedFloat64(V4220)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(3)
__typedArg1 := V4220
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

let__18807 := tmp18833
_ = let__18807

tmp18834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__18807, MakeInteger(0), symshen_4dictionary)
}
__typedArg0 := let__18807
__typedArg1 := MakeInteger(0)
__typedArg2 := symshen_4dictionary
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__18808 := tmp18834
_ = let__18808

tmp18835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__18807, MakeInteger(1), V4220)
}
__typedArg0 := let__18807
__typedArg1 := MakeInteger(1)
__typedArg2 := V4220
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__18809 := tmp18835
_ = let__18809

tmp18836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__18807, MakeInteger(2), MakeInteger(0))
}
__typedArg0 := let__18807
__typedArg1 := MakeInteger(2)
__typedArg2 := MakeInteger(0)
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__18810 := tmp18836
_ = let__18810

tmp18838 := Call(__e, PrimFunc(symshen_4fillvector), let__18807, MakeInteger(3), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(2))
__typedN1, __typedOK1 := TypedFloat64(V4220)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(2)
__typedArg1 := V4220
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})(), Nil)


let__18811 := tmp18838
_ = let__18811

__e.Return(let__18807)
return


}


}, 1)

tmp18841 := Call(__e, ns2_1set, symshen_4dict, tmp18829)


_ = tmp18841

tmp18842 := MakeNative(func(__e *ControlFlow) {
V4226 := __e.Get(1)
_ = V4226
tmp18849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V4226)
}
__typedArg0 := V4226
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp18849 {
tmp18844 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V4226, MakeInteger(0))
}
__typedArg0 := V4226
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 0)

tmp18845 := MakeNative(func(__e *ControlFlow) {
Z4227 := __e.Get(1)
_ = Z4227
__e.Return(symshen_4not_1dictionary)
return
}, 1)

tmp18846 := Call(__e, try_1catch, tmp18844, tmp18845)


tmp18847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp18846, symshen_4dictionary)
}
__typedArg0 := tmp18846
__typedArg1 := symshen_4dictionary
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18847 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


} else {
__e.Return(False)
return
}


}, 1)

tmp18850 := Call(__e, ns2_1set, symshen_4dict_2, tmp18842)


_ = tmp18850

tmp18851 := MakeNative(func(__e *ControlFlow) {
V4228 := __e.Get(1)
_ = V4228
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V4228, MakeInteger(1))
}
__typedArg0 := V4228
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp18852 := Call(__e, ns2_1set, symshen_4dict_1capacity, tmp18851)


_ = tmp18852

tmp18853 := MakeNative(func(__e *ControlFlow) {
V4229 := __e.Get(1)
_ = V4229
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V4229, MakeInteger(2))
}
__typedArg0 := V4229
__typedArg1 := MakeInteger(2)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp18854 := Call(__e, ns2_1set, symshen_4dict_1count, tmp18853)


_ = tmp18854

tmp18855 := MakeNative(func(__e *ControlFlow) {
V4230 := __e.Get(1)
_ = V4230
V4231 := __e.Get(2)
_ = V4231
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V4230, MakeInteger(2), V4231)
}
__typedArg0 := V4230
__typedArg1 := MakeInteger(2)
__typedArg2 := V4231
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return
}, 2)

tmp18856 := Call(__e, ns2_1set, symshen_4dict_1count_1_6, tmp18855)


_ = tmp18856

tmp18857 := MakeNative(func(__e *ControlFlow) {
V4232 := __e.Get(1)
_ = V4232
V4233 := __e.Get(2)
_ = V4233
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V4232, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(3))
__typedN1, __typedOK1 := TypedFloat64(V4233)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(3)
__typedArg1 := V4233
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := V4232
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(3))
__typedN1, __typedOK1 := TypedFloat64(V4233)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(3)
__typedArg1 := V4233
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp18859 := Call(__e, ns2_1set, symshen_4_5_1dict_1bucket, tmp18857)


_ = tmp18859

tmp18860 := MakeNative(func(__e *ControlFlow) {
V4234 := __e.Get(1)
_ = V4234
V4235 := __e.Get(2)
_ = V4235
V4236 := __e.Get(3)
_ = V4236
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V4234, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(3))
__typedN1, __typedOK1 := TypedFloat64(V4235)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(3)
__typedArg1 := V4235
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})(), V4236)
}
__typedArg0 := V4234
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(3))
__typedN1, __typedOK1 := TypedFloat64(V4235)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(3)
__typedArg1 := V4235
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
__typedArg2 := V4236
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return


}, 3)

tmp18862 := Call(__e, ns2_1set, symshen_4dict_1bucket_1_6, tmp18860)


_ = tmp18862

tmp18863 := MakeNative(func(__e *ControlFlow) {
V4237 := __e.Get(1)
_ = V4237
V4238 := __e.Get(2)
_ = V4238
V4239 := __e.Get(3)
_ = V4239
tmp18864 := Call(__e, PrimFunc(symlength), V4239)


tmp18865 := Call(__e, PrimFunc(symlength), V4238)


let__18812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(tmp18864)
__typedN1, __typedOK1 := TypedFloat64(tmp18865)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := tmp18864
__typedArg1 := tmp18865
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})()
_ = let__18812

tmp18867 := Call(__e, PrimFunc(symshen_4dict_1count), V4237)


__e.TailApply(PrimFunc(symshen_4dict_1count_1_6), V4237, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(let__18812)
__typedN1, __typedOK1 := TypedFloat64(tmp18867)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := let__18812
__typedArg1 := tmp18867
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


}, 3)

tmp18869 := Call(__e, ns2_1set, symshen_4dict_1update_1count, tmp18863)


_ = tmp18869

tmp18870 := MakeNative(func(__e *ControlFlow) {
V4241 := __e.Get(1)
_ = V4241
V4242 := __e.Get(2)
_ = V4242
V4243 := __e.Get(3)
_ = V4243
tmp18871 := Call(__e, PrimFunc(symshen_4dict_1capacity), V4241)


tmp18872 := Call(__e, PrimFunc(symhash), V4242, tmp18871)


let__18813 := tmp18872
_ = let__18813

tmp18873 := Call(__e, PrimFunc(symshen_4_5_1dict_1bucket), V4241, let__18813)


let__18814 := tmp18873
_ = let__18814

tmp18874 := Call(__e, PrimFunc(symshen_4assoc_1set), V4242, V4243, let__18814)


let__18815 := tmp18874
_ = let__18815

tmp18875 := Call(__e, PrimFunc(symshen_4dict_1bucket_1_6), V4241, let__18813, let__18815)


let__18816 := tmp18875
_ = let__18816

tmp18876 := Call(__e, PrimFunc(symshen_4dict_1update_1count), V4241, let__18814, let__18815)


let__18817 := tmp18876
_ = let__18817

__e.Return(V4243)
return


}, 3)

tmp18877 := Call(__e, ns2_1set, symshen_4dict_1_6, tmp18870)


_ = tmp18877

tmp18878 := MakeNative(func(__e *ControlFlow) {
V4249 := __e.Get(1)
_ = V4249
V4250 := __e.Get(2)
_ = V4250
tmp18879 := Call(__e, PrimFunc(symshen_4dict_1capacity), V4249)


tmp18880 := Call(__e, PrimFunc(symhash), V4250, tmp18879)


let__18818 := tmp18880
_ = let__18818

tmp18881 := Call(__e, PrimFunc(symshen_4_5_1dict_1bucket), V4249, let__18818)


let__18819 := tmp18881
_ = let__18819

tmp18882 := Call(__e, PrimFunc(symassoc), V4250, let__18819)


let__18820 := tmp18882
_ = let__18820

tmp18886 := Call(__e, PrimFunc(symempty_2), let__18820)


if True == tmp18886 {
tmp18883 := Call(__e, PrimFunc(symshen_4app), V4250, MakeString(" not found in dict\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("value "))
__typedS1, __typedOK1 := TypedString(tmp18883)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("value ")
__typedArg1 := tmp18883
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("value "))
__typedS1, __typedOK1 := TypedString(tmp18883)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("value ")
__typedArg1 := tmp18883
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__18820)
}
__typedArg0 := let__18820
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
}


}, 2)

tmp18887 := Call(__e, ns2_1set, symshen_4_5_1dict, tmp18878)


_ = tmp18887

tmp18888 := MakeNative(func(__e *ControlFlow) {
V4254 := __e.Get(1)
_ = V4254
V4255 := __e.Get(2)
_ = V4255
tmp18889 := Call(__e, PrimFunc(symshen_4dict_1capacity), V4254)


tmp18890 := Call(__e, PrimFunc(symhash), V4255, tmp18889)


let__18821 := tmp18890
_ = let__18821

tmp18891 := Call(__e, PrimFunc(symshen_4_5_1dict_1bucket), V4254, let__18821)


let__18822 := tmp18891
_ = let__18822

tmp18892 := Call(__e, PrimFunc(symshen_4assoc_1rm), V4255, let__18822)


let__18823 := tmp18892
_ = let__18823

tmp18893 := Call(__e, PrimFunc(symshen_4dict_1bucket_1_6), V4254, let__18821, let__18823)


let__18824 := tmp18893
_ = let__18824

tmp18894 := Call(__e, PrimFunc(symshen_4dict_1update_1count), V4254, let__18822, let__18823)


let__18825 := tmp18894
_ = let__18825

__e.Return(V4255)
return


}, 2)

tmp18895 := Call(__e, ns2_1set, symshen_4dict_1rm, tmp18888)


_ = tmp18895

tmp18896 := MakeNative(func(__e *ControlFlow) {
V4261 := __e.Get(1)
_ = V4261
V4262 := __e.Get(2)
_ = V4262
V4263 := __e.Get(3)
_ = V4263
tmp18897 := Call(__e, PrimFunc(symshen_4dict_1capacity), V4262)


let__18826 := tmp18897
_ = let__18826

__e.TailApply(PrimFunc(symshen_4dict_1fold_1h), V4261, V4262, V4263, MakeInteger(0), let__18826)
return


}, 3)

tmp18898 := Call(__e, ns2_1set, symshen_4dict_1fold, tmp18896)


_ = tmp18898

tmp18899 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
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
tmp18904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V4269, V4270)
}
__typedArg0 := V4269
__typedArg1 := V4270
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18904 {
__e.Return(V4268)
return
} else {
tmp18900 := Call(__e, PrimFunc(symshen_4_5_1dict_1bucket), V4267, V4269)


let__18827 := tmp18900
_ = let__18827

tmp18901 := Call(__e, PrimFunc(symshen_4bucket_1fold), V4266, let__18827, V4268)


let__18828 := tmp18901
_ = let__18828

if PrimFunc(symshen_4dict_1fold_1h) == __self {
__self1, __self2, __self3, __self4, __self5 = V4266, V4267, let__18828, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(V4269)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := V4269
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})(), V4270
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4dict_1fold_1h), V4266, V4267, let__18828, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(V4269)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := V4269
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})(), V4270)
return


}


}, 5)

tmp18905 := Call(__e, ns2_1set, symshen_4dict_1fold_1h, tmp18899)


_ = tmp18905

tmp18906 := MakeNative(func(__e *ControlFlow) {
V4273 := __e.Get(1)
_ = V4273
V4274 := __e.Get(2)
_ = V4274
V4275 := __e.Get(3)
_ = V4275
tmp18922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4274)
}
__typedArg0 := Nil
__typedArg1 := V4274
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18922 {
__e.Return(V4275)
return
} else {
tmp18920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4274)
}
__typedArg0 := V4274
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18916 Obj

if True == tmp18920 {
tmp18918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4274)
}
__typedArg0 := V4274
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18918)
}
__typedArg0 := tmp18918
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18917 Obj

if True == tmp18919 {
ifres18917 = True


} else {
ifres18917 = False


}

ifres18916 = ifres18917


} else {
ifres18916 = False


}

if True == ifres18916 {
tmp18907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4274)
}
__typedArg0 := V4274
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18907)
}
__typedArg0 := tmp18907
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18909 := Call(__e, V4273, tmp18908)


tmp18910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4274)
}
__typedArg0 := V4274
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18910)
}
__typedArg0 := tmp18910
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18912 := Call(__e, tmp18909, tmp18911)


tmp18913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4274)
}
__typedArg0 := V4274
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18914 := Call(__e, PrimFunc(symshen_4bucket_1fold), V4273, tmp18913, V4275)


__e.TailApply(tmp18912, tmp18914)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4bucket_1fold)
return
}


}


}, 3)

tmp18923 := Call(__e, ns2_1set, symshen_4bucket_1fold, tmp18906)


_ = tmp18923

tmp18924 := MakeNative(func(__e *ControlFlow) {
V4276 := __e.Get(1)
_ = V4276
tmp18925 := MakeNative(func(__e *ControlFlow) {
Z4277 := __e.Get(1)
_ = Z4277
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4278 := __e.Get(1)
_ = Z4278
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4279 := __e.Get(1)
_ = Z4279
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z4277, Z4279)
}
__typedArg0 := Z4277
__typedArg1 := Z4279
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(PrimFunc(symshen_4dict_1fold), tmp18925, V4276, Nil)
return


}, 1)

tmp18926 := Call(__e, ns2_1set, symshen_4dict_1keys, tmp18924)


_ = tmp18926

tmp18927 := MakeNative(func(__e *ControlFlow) {
V4280 := __e.Get(1)
_ = V4280
tmp18928 := MakeNative(func(__e *ControlFlow) {
Z4281 := __e.Get(1)
_ = Z4281
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4282 := __e.Get(1)
_ = Z4282
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4283 := __e.Get(1)
_ = Z4283
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z4282, Z4283)
}
__typedArg0 := Z4282
__typedArg1 := Z4283
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(PrimFunc(symshen_4dict_1fold), tmp18928, V4280, Nil)
return


}, 1)

__e.TailApply(ns2_1set, symshen_4dict_1values, tmp18927)
return




}, 0)

