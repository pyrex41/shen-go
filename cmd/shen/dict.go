package main

import . "github.com/pyrex41/shen-go/kl"

var DictMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp18806 := MakeNative1(func(__e *ControlFlow, V4205 Obj) Obj {
if True == GLt(__e, V4205, MakeInteger(1)) {
tmp18807 := Call3(__e, SymFn(symshen_4app), V4205, MakeString(""), symshen_4s)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("invalid initial dict size: "), tmp18807))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("invalid initial dict size: "), tmp18807)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


} else {
tmp18810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(GAdd(__e, MakeInteger(3), V4205))
}
__typedArg0 := GAdd(__e, MakeInteger(3), V4205)
return Call1(__e, SymFn(symabsvector), __typedArg0)
})()

let__18784 := tmp18810
_ = let__18784

tmp18811 := GVecSet(__e, let__18784, MakeInteger(0), symshen_4dictionary)

let__18785 := tmp18811
_ = let__18785

tmp18812 := GVecSet(__e, let__18784, MakeInteger(1), V4205)

let__18786 := tmp18812
_ = let__18786

tmp18813 := GVecSet(__e, let__18784, MakeInteger(2), MakeInteger(0))

let__18787 := tmp18813
_ = let__18787

tmp18815 := Call4(__e, SymFn(symshen_4fillvector), let__18784, MakeInteger(3), GAdd(__e, MakeInteger(2), V4205), Nil)


let__18788 := tmp18815
_ = let__18788

return let__18784


}


panic("unreachable")
})

tmp18818 := Call2(__e, ns2_1set, symshen_4dict, tmp18806)


_ = tmp18818

tmp18819 := MakeNative1(func(__e *ControlFlow, V4211 Obj) Obj {
tmp18826 := GVectorP(__e, V4211)

if True == tmp18826 {
tmp18821 := MakeNative0(func(__e *ControlFlow) Obj {
return GVecGet(__e, V4211, MakeInteger(0))
panic("unreachable")
})

tmp18822 := MakeNative1(func(__e *ControlFlow, Z4212 Obj) Obj {
return symshen_4not_1dictionary
panic("unreachable")
})

tmp18823 := Call2(__e, try_1catch, tmp18821, tmp18822)


tmp18824 := GEq(__e, tmp18823, symshen_4dictionary)

if True == tmp18824 {
return True
} else {
return False
}


} else {
return False
}


panic("unreachable")
})

tmp18827 := Call2(__e, ns2_1set, symshen_4dict_2, tmp18819)


_ = tmp18827

tmp18828 := MakeNative1(func(__e *ControlFlow, V4213 Obj) Obj {
return GVecGet(__e, V4213, MakeInteger(1))
panic("unreachable")
})

tmp18829 := Call2(__e, ns2_1set, symshen_4dict_1capacity, tmp18828)


_ = tmp18829

tmp18830 := MakeNative1(func(__e *ControlFlow, V4214 Obj) Obj {
return GVecGet(__e, V4214, MakeInteger(2))
panic("unreachable")
})

tmp18831 := Call2(__e, ns2_1set, symshen_4dict_1count, tmp18830)


_ = tmp18831

tmp18832 := MakeNative2(func(__e *ControlFlow, V4215, V4216 Obj) Obj {
return GVecSet(__e, V4215, MakeInteger(2), V4216)
panic("unreachable")
})

tmp18833 := Call2(__e, ns2_1set, symshen_4dict_1count_1_6, tmp18832)


_ = tmp18833

tmp18834 := MakeNative2(func(__e *ControlFlow, V4217, V4218 Obj) Obj {
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

tmp18836 := Call2(__e, ns2_1set, symshen_4_5_1dict_1bucket, tmp18834)


_ = tmp18836

tmp18837 := MakeNative3(func(__e *ControlFlow, V4219, V4220, V4221 Obj) Obj {
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

tmp18839 := Call2(__e, ns2_1set, symshen_4dict_1bucket_1_6, tmp18837)


_ = tmp18839

tmp18840 := MakeNative3(func(__e *ControlFlow, V4222, V4223, V4224 Obj) Obj {
tmp18841 := Call1(__e, SymFn(symlength), V4224)


tmp18842 := Call1(__e, SymFn(symlength), V4223)


let__18789 := GSub(__e, tmp18841, tmp18842)
_ = let__18789

tmp18844 := Call1(__e, SymFn(symshen_4dict_1count), V4222)


return TailCall2(__e, SymFn(symshen_4dict_1count_1_6), V4222, GAdd(__e, let__18789, tmp18844))


panic("unreachable")
})

tmp18846 := Call2(__e, ns2_1set, symshen_4dict_1update_1count, tmp18840)


_ = tmp18846

tmp18847 := MakeNative3(func(__e *ControlFlow, V4226, V4227, V4228 Obj) Obj {
tmp18848 := Call1(__e, SymFn(symshen_4dict_1capacity), V4226)


tmp18849 := Call2(__e, SymFn(symhash), V4227, tmp18848)


let__18790 := tmp18849
_ = let__18790

tmp18850 := Call2(__e, SymFn(symshen_4_5_1dict_1bucket), V4226, let__18790)


let__18791 := tmp18850
_ = let__18791

tmp18851 := Call3(__e, SymFn(symshen_4assoc_1set), V4227, V4228, let__18791)


let__18792 := tmp18851
_ = let__18792

tmp18852 := Call3(__e, SymFn(symshen_4dict_1bucket_1_6), V4226, let__18790, let__18792)


let__18793 := tmp18852
_ = let__18793

tmp18853 := Call3(__e, SymFn(symshen_4dict_1update_1count), V4226, let__18791, let__18792)


let__18794 := tmp18853
_ = let__18794

return V4228


panic("unreachable")
})

tmp18854 := Call2(__e, ns2_1set, symshen_4dict_1_6, tmp18847)


_ = tmp18854

tmp18855 := MakeNative2(func(__e *ControlFlow, V4234, V4235 Obj) Obj {
tmp18856 := Call1(__e, SymFn(symshen_4dict_1capacity), V4234)


tmp18857 := Call2(__e, SymFn(symhash), V4235, tmp18856)


let__18795 := tmp18857
_ = let__18795

tmp18858 := Call2(__e, SymFn(symshen_4_5_1dict_1bucket), V4234, let__18795)


let__18796 := tmp18858
_ = let__18796

tmp18859 := Call2(__e, SymFn(symassoc), V4235, let__18796)


let__18797 := tmp18859
_ = let__18797

tmp18863 := Call1(__e, SymFn(symempty_2), let__18797)


if True == tmp18863 {
tmp18860 := Call3(__e, SymFn(symshen_4app), V4235, MakeString(" not found in dict\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("value "), tmp18860))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("value "), tmp18860)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


} else {
return GTl(__e, let__18797)
}


panic("unreachable")
})

tmp18864 := Call2(__e, ns2_1set, symshen_4_5_1dict, tmp18855)


_ = tmp18864

tmp18865 := MakeNative2(func(__e *ControlFlow, V4239, V4240 Obj) Obj {
tmp18866 := Call1(__e, SymFn(symshen_4dict_1capacity), V4239)


tmp18867 := Call2(__e, SymFn(symhash), V4240, tmp18866)


let__18798 := tmp18867
_ = let__18798

tmp18868 := Call2(__e, SymFn(symshen_4_5_1dict_1bucket), V4239, let__18798)


let__18799 := tmp18868
_ = let__18799

tmp18869 := Call2(__e, SymFn(symshen_4assoc_1rm), V4240, let__18799)


let__18800 := tmp18869
_ = let__18800

tmp18870 := Call3(__e, SymFn(symshen_4dict_1bucket_1_6), V4239, let__18798, let__18800)


let__18801 := tmp18870
_ = let__18801

tmp18871 := Call3(__e, SymFn(symshen_4dict_1update_1count), V4239, let__18799, let__18800)


let__18802 := tmp18871
_ = let__18802

return V4240


panic("unreachable")
})

tmp18872 := Call2(__e, ns2_1set, symshen_4dict_1rm, tmp18865)


_ = tmp18872

tmp18873 := MakeNative3(func(__e *ControlFlow, V4246, V4247, V4248 Obj) Obj {
tmp18874 := Call1(__e, SymFn(symshen_4dict_1capacity), V4247)


let__18803 := tmp18874
_ = let__18803

return TailCall5(__e, SymFn(symshen_4dict_1fold_1h), V4246, V4247, V4248, MakeInteger(0), let__18803)


panic("unreachable")
})

tmp18875 := Call2(__e, ns2_1set, symshen_4dict_1fold, tmp18873)


_ = tmp18875

tmp18876 := (func() Obj {
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
tmp18881 := GEq(__e, V4254, V4255)

if True == tmp18881 {
return V4253
} else {
tmp18877 := Call2(__e, SymFn(symshen_4_5_1dict_1bucket), V4252, V4254)


let__18804 := tmp18877
_ = let__18804

tmp18878 := Call3(__e, SymFn(symshen_4bucket_1fold), V4251, let__18804, V4253)


let__18805 := tmp18878
_ = let__18805

if SymFn(symshen_4dict_1fold_1h) == __self {
__self1, __self2, __self3, __self4, __self5 = V4251, V4252, let__18805, GAdd(__e, MakeInteger(1), V4254), V4255
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4dict_1fold_1h), V4251, V4252, let__18805, GAdd(__e, MakeInteger(1), V4254), V4255)


}


panic("unreachable")
})
return __self
})()

tmp18882 := Call2(__e, ns2_1set, symshen_4dict_1fold_1h, tmp18876)


_ = tmp18882

tmp18883 := MakeNative3(func(__e *ControlFlow, V4258, V4259, V4260 Obj) Obj {
tmp18899 := GEq(__e, Nil, V4259)

if True == tmp18899 {
return V4260
} else {
tmp18897 := GConsP(__e, V4259)

var ifres18893 Obj

if True == tmp18897 {
tmp18895 := GHd(__e, V4259)

tmp18896 := GConsP(__e, tmp18895)

var ifres18894 Obj

if True == tmp18896 {
ifres18894 = True


} else {
ifres18894 = False


}

ifres18893 = ifres18894


} else {
ifres18893 = False


}

if True == ifres18893 {
tmp18884 := GHd(__e, V4259)

tmp18885 := GHd(__e, tmp18884)

tmp18886 := Call1(__e, V4258, tmp18885)


tmp18887 := GHd(__e, V4259)

tmp18888 := GTl(__e, tmp18887)

tmp18889 := Call1(__e, tmp18886, tmp18888)


tmp18890 := GTl(__e, V4259)

tmp18891 := Call3(__e, SymFn(symshen_4bucket_1fold), V4258, tmp18890, V4260)


return TailCall1(__e, tmp18889, tmp18891)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4bucket_1fold)
}


}


panic("unreachable")
})

tmp18900 := Call2(__e, ns2_1set, symshen_4bucket_1fold, tmp18883)


_ = tmp18900

tmp18901 := MakeNative1(func(__e *ControlFlow, V4261 Obj) Obj {
tmp18902 := MakeNative1(func(__e *ControlFlow, Z4262 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z4263 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z4264 Obj) Obj {
return GCons(__e, Z4262, Z4264)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

return TailCall3(__e, SymFn(symshen_4dict_1fold), tmp18902, V4261, Nil)


panic("unreachable")
})

tmp18903 := Call2(__e, ns2_1set, symshen_4dict_1keys, tmp18901)


_ = tmp18903

tmp18904 := MakeNative1(func(__e *ControlFlow, V4265 Obj) Obj {
tmp18905 := MakeNative1(func(__e *ControlFlow, Z4266 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z4267 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z4268 Obj) Obj {
return GCons(__e, Z4267, Z4268)
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

return TailCall3(__e, SymFn(symshen_4dict_1fold), tmp18905, V4265, Nil)


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4dict_1values, tmp18904)




panic("unreachable")
})

