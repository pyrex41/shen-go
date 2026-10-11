package main

import . "github.com/pyrex41/shen-go/kl"

var DictMain = MakeNative(func(__e *ControlFlow) {
tmp18789 := MakeNative(func(__e *ControlFlow) {
V4225 := __e.Get(1)
_ = V4225
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5) {
__typedN0, __typedOK0 := TypedFloat64(V4225)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5) {
return TypedMaterializeBoolean((__typedN0 < __typedN1))
}}
__typedArg0 := V4225
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5), __typedArg0, __typedArg1)
})() {
tmp18790 := Call(__e, PrimFunc(symshen_4app), V4225, MakeString(""), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("invalid initial dict size: "))
__typedS1, __typedOK1 := TypedString(tmp18790)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("invalid initial dict size: ")
__typedArg1 := tmp18790
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("invalid initial dict size: "))
__typedS1, __typedOK1 := TypedString(tmp18790)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("invalid initial dict size: ")
__typedArg1 := tmp18790
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
tmp18793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(3))
__typedN1, __typedOK1 := TypedFloat64(V4225)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(3)
__typedArg1 := V4225
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(3))
__typedN1, __typedOK1 := TypedFloat64(V4225)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(3)
__typedArg1 := V4225
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

let__18767 := tmp18793
_ = let__18767

tmp18794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__18767, MakeInteger(0), symshen_4dictionary)
}
__typedArg0 := let__18767
__typedArg1 := MakeInteger(0)
__typedArg2 := symshen_4dictionary
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__18768 := tmp18794
_ = let__18768

tmp18795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__18767, MakeInteger(1), V4225)
}
__typedArg0 := let__18767
__typedArg1 := MakeInteger(1)
__typedArg2 := V4225
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__18769 := tmp18795
_ = let__18769

tmp18796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__18767, MakeInteger(2), MakeInteger(0))
}
__typedArg0 := let__18767
__typedArg1 := MakeInteger(2)
__typedArg2 := MakeInteger(0)
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__18770 := tmp18796
_ = let__18770

tmp18798 := Call(__e, PrimFunc(symshen_4fillvector), let__18767, MakeInteger(3), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(2))
__typedN1, __typedOK1 := TypedFloat64(V4225)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(2)
__typedArg1 := V4225
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})(), Nil)


let__18771 := tmp18798
_ = let__18771

__e.Return(let__18767)
return


}


}, 1)

tmp18801 := Call(__e, ns2_1set, symshen_4dict, tmp18789)


_ = tmp18801

tmp18802 := MakeNative(func(__e *ControlFlow) {
V4231 := __e.Get(1)
_ = V4231
tmp18809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V4231)
}
__typedArg0 := V4231
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp18809 {
tmp18804 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V4231, MakeInteger(0))
}
__typedArg0 := V4231
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 0)

tmp18805 := MakeNative(func(__e *ControlFlow) {
Z4232 := __e.Get(1)
_ = Z4232
__e.Return(symshen_4not_1dictionary)
return
}, 1)

tmp18806 := Call(__e, try_1catch, tmp18804, tmp18805)


tmp18807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp18806, symshen_4dictionary)
}
__typedArg0 := tmp18806
__typedArg1 := symshen_4dictionary
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18807 {
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

tmp18810 := Call(__e, ns2_1set, symshen_4dict_2, tmp18802)


_ = tmp18810

tmp18811 := MakeNative(func(__e *ControlFlow) {
V4233 := __e.Get(1)
_ = V4233
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V4233, MakeInteger(1))
}
__typedArg0 := V4233
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp18812 := Call(__e, ns2_1set, symshen_4dict_1capacity, tmp18811)


_ = tmp18812

tmp18813 := MakeNative(func(__e *ControlFlow) {
V4234 := __e.Get(1)
_ = V4234
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V4234, MakeInteger(2))
}
__typedArg0 := V4234
__typedArg1 := MakeInteger(2)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp18814 := Call(__e, ns2_1set, symshen_4dict_1count, tmp18813)


_ = tmp18814

tmp18815 := MakeNative(func(__e *ControlFlow) {
V4235 := __e.Get(1)
_ = V4235
V4236 := __e.Get(2)
_ = V4236
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V4235, MakeInteger(2), V4236)
}
__typedArg0 := V4235
__typedArg1 := MakeInteger(2)
__typedArg2 := V4236
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return
}, 2)

tmp18816 := Call(__e, ns2_1set, symshen_4dict_1count_1_6, tmp18815)


_ = tmp18816

tmp18817 := MakeNative(func(__e *ControlFlow) {
V4237 := __e.Get(1)
_ = V4237
V4238 := __e.Get(2)
_ = V4238
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V4237, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(3))
__typedN1, __typedOK1 := TypedFloat64(V4238)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(3)
__typedArg1 := V4238
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := V4237
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(3))
__typedN1, __typedOK1 := TypedFloat64(V4238)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(3)
__typedArg1 := V4238
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp18819 := Call(__e, ns2_1set, symshen_4_5_1dict_1bucket, tmp18817)


_ = tmp18819

tmp18820 := MakeNative(func(__e *ControlFlow) {
V4239 := __e.Get(1)
_ = V4239
V4240 := __e.Get(2)
_ = V4240
V4241 := __e.Get(3)
_ = V4241
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V4239, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(3))
__typedN1, __typedOK1 := TypedFloat64(V4240)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(3)
__typedArg1 := V4240
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})(), V4241)
}
__typedArg0 := V4239
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(3))
__typedN1, __typedOK1 := TypedFloat64(V4240)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(3)
__typedArg1 := V4240
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
__typedArg2 := V4241
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return


}, 3)

tmp18822 := Call(__e, ns2_1set, symshen_4dict_1bucket_1_6, tmp18820)


_ = tmp18822

tmp18823 := MakeNative(func(__e *ControlFlow) {
V4242 := __e.Get(1)
_ = V4242
V4243 := __e.Get(2)
_ = V4243
V4244 := __e.Get(3)
_ = V4244
tmp18824 := Call(__e, PrimFunc(symlength), V4244)


tmp18825 := Call(__e, PrimFunc(symlength), V4243)


let__18772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(tmp18824)
__typedN1, __typedOK1 := TypedFloat64(tmp18825)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := tmp18824
__typedArg1 := tmp18825
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})()
_ = let__18772

tmp18827 := Call(__e, PrimFunc(symshen_4dict_1count), V4242)


__e.TailApply(PrimFunc(symshen_4dict_1count_1_6), V4242, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(let__18772)
__typedN1, __typedOK1 := TypedFloat64(tmp18827)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := let__18772
__typedArg1 := tmp18827
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


}, 3)

tmp18829 := Call(__e, ns2_1set, symshen_4dict_1update_1count, tmp18823)


_ = tmp18829

tmp18830 := MakeNative(func(__e *ControlFlow) {
V4246 := __e.Get(1)
_ = V4246
V4247 := __e.Get(2)
_ = V4247
V4248 := __e.Get(3)
_ = V4248
tmp18831 := Call(__e, PrimFunc(symshen_4dict_1capacity), V4246)


tmp18832 := Call(__e, PrimFunc(symhash), V4247, tmp18831)


let__18773 := tmp18832
_ = let__18773

tmp18833 := Call(__e, PrimFunc(symshen_4_5_1dict_1bucket), V4246, let__18773)


let__18774 := tmp18833
_ = let__18774

tmp18834 := Call(__e, PrimFunc(symshen_4assoc_1set), V4247, V4248, let__18774)


let__18775 := tmp18834
_ = let__18775

tmp18835 := Call(__e, PrimFunc(symshen_4dict_1bucket_1_6), V4246, let__18773, let__18775)


let__18776 := tmp18835
_ = let__18776

tmp18836 := Call(__e, PrimFunc(symshen_4dict_1update_1count), V4246, let__18774, let__18775)


let__18777 := tmp18836
_ = let__18777

__e.Return(V4248)
return


}, 3)

tmp18837 := Call(__e, ns2_1set, symshen_4dict_1_6, tmp18830)


_ = tmp18837

tmp18838 := MakeNative(func(__e *ControlFlow) {
V4254 := __e.Get(1)
_ = V4254
V4255 := __e.Get(2)
_ = V4255
tmp18839 := Call(__e, PrimFunc(symshen_4dict_1capacity), V4254)


tmp18840 := Call(__e, PrimFunc(symhash), V4255, tmp18839)


let__18778 := tmp18840
_ = let__18778

tmp18841 := Call(__e, PrimFunc(symshen_4_5_1dict_1bucket), V4254, let__18778)


let__18779 := tmp18841
_ = let__18779

tmp18842 := Call(__e, PrimFunc(symassoc), V4255, let__18779)


let__18780 := tmp18842
_ = let__18780

tmp18846 := Call(__e, PrimFunc(symempty_2), let__18780)


if True == tmp18846 {
tmp18843 := Call(__e, PrimFunc(symshen_4app), V4255, MakeString(" not found in dict\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("value "))
__typedS1, __typedOK1 := TypedString(tmp18843)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("value ")
__typedArg1 := tmp18843
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("value "))
__typedS1, __typedOK1 := TypedString(tmp18843)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("value ")
__typedArg1 := tmp18843
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__18780)
}
__typedArg0 := let__18780
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
}


}, 2)

tmp18847 := Call(__e, ns2_1set, symshen_4_5_1dict, tmp18838)


_ = tmp18847

tmp18848 := MakeNative(func(__e *ControlFlow) {
V4259 := __e.Get(1)
_ = V4259
V4260 := __e.Get(2)
_ = V4260
tmp18849 := Call(__e, PrimFunc(symshen_4dict_1capacity), V4259)


tmp18850 := Call(__e, PrimFunc(symhash), V4260, tmp18849)


let__18781 := tmp18850
_ = let__18781

tmp18851 := Call(__e, PrimFunc(symshen_4_5_1dict_1bucket), V4259, let__18781)


let__18782 := tmp18851
_ = let__18782

tmp18852 := Call(__e, PrimFunc(symshen_4assoc_1rm), V4260, let__18782)


let__18783 := tmp18852
_ = let__18783

tmp18853 := Call(__e, PrimFunc(symshen_4dict_1bucket_1_6), V4259, let__18781, let__18783)


let__18784 := tmp18853
_ = let__18784

tmp18854 := Call(__e, PrimFunc(symshen_4dict_1update_1count), V4259, let__18782, let__18783)


let__18785 := tmp18854
_ = let__18785

__e.Return(V4260)
return


}, 2)

tmp18855 := Call(__e, ns2_1set, symshen_4dict_1rm, tmp18848)


_ = tmp18855

tmp18856 := MakeNative(func(__e *ControlFlow) {
V4266 := __e.Get(1)
_ = V4266
V4267 := __e.Get(2)
_ = V4267
V4268 := __e.Get(3)
_ = V4268
tmp18857 := Call(__e, PrimFunc(symshen_4dict_1capacity), V4267)


let__18786 := tmp18857
_ = let__18786

__e.TailApply(PrimFunc(symshen_4dict_1fold_1h), V4266, V4267, V4268, MakeInteger(0), let__18786)
return


}, 3)

tmp18858 := Call(__e, ns2_1set, symshen_4dict_1fold, tmp18856)


_ = tmp18858

tmp18859 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
__selftop:
V4271 := __self1
_ = V4271
V4272 := __self2
_ = V4272
V4273 := __self3
_ = V4273
V4274 := __self4
_ = V4274
V4275 := __self5
_ = V4275
tmp18864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V4274, V4275)
}
__typedArg0 := V4274
__typedArg1 := V4275
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18864 {
__e.Return(V4273)
return
} else {
tmp18860 := Call(__e, PrimFunc(symshen_4_5_1dict_1bucket), V4272, V4274)


let__18787 := tmp18860
_ = let__18787

tmp18861 := Call(__e, PrimFunc(symshen_4bucket_1fold), V4271, let__18787, V4273)


let__18788 := tmp18861
_ = let__18788

if PrimFunc(symshen_4dict_1fold_1h) == __self {
__self1, __self2, __self3, __self4, __self5 = V4271, V4272, let__18788, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(V4274)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := V4274
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})(), V4275
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4dict_1fold_1h), V4271, V4272, let__18788, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(V4274)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := V4274
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})(), V4275)
return


}


}, 5)

tmp18865 := Call(__e, ns2_1set, symshen_4dict_1fold_1h, tmp18859)


_ = tmp18865

tmp18866 := MakeNative(func(__e *ControlFlow) {
V4278 := __e.Get(1)
_ = V4278
V4279 := __e.Get(2)
_ = V4279
V4280 := __e.Get(3)
_ = V4280
tmp18882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4279)
}
__typedArg0 := Nil
__typedArg1 := V4279
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18882 {
__e.Return(V4280)
return
} else {
tmp18880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4279)
}
__typedArg0 := V4279
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18876 Obj

if True == tmp18880 {
tmp18878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4279)
}
__typedArg0 := V4279
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18878)
}
__typedArg0 := tmp18878
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18877 Obj

if True == tmp18879 {
ifres18877 = True


} else {
ifres18877 = False


}

ifres18876 = ifres18877


} else {
ifres18876 = False


}

if True == ifres18876 {
tmp18867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4279)
}
__typedArg0 := V4279
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18867)
}
__typedArg0 := tmp18867
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18869 := Call(__e, V4278, tmp18868)


tmp18870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4279)
}
__typedArg0 := V4279
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18870)
}
__typedArg0 := tmp18870
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18872 := Call(__e, tmp18869, tmp18871)


tmp18873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4279)
}
__typedArg0 := V4279
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18874 := Call(__e, PrimFunc(symshen_4bucket_1fold), V4278, tmp18873, V4280)


__e.TailApply(tmp18872, tmp18874)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4bucket_1fold)
return
}


}


}, 3)

tmp18883 := Call(__e, ns2_1set, symshen_4bucket_1fold, tmp18866)


_ = tmp18883

tmp18884 := MakeNative(func(__e *ControlFlow) {
V4281 := __e.Get(1)
_ = V4281
tmp18885 := MakeNative(func(__e *ControlFlow) {
Z4282 := __e.Get(1)
_ = Z4282
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4283 := __e.Get(1)
_ = Z4283
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4284 := __e.Get(1)
_ = Z4284
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z4282, Z4284)
}
__typedArg0 := Z4282
__typedArg1 := Z4284
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(PrimFunc(symshen_4dict_1fold), tmp18885, V4281, Nil)
return


}, 1)

tmp18886 := Call(__e, ns2_1set, symshen_4dict_1keys, tmp18884)


_ = tmp18886

tmp18887 := MakeNative(func(__e *ControlFlow) {
V4285 := __e.Get(1)
_ = V4285
tmp18888 := MakeNative(func(__e *ControlFlow) {
Z4286 := __e.Get(1)
_ = Z4286
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4287 := __e.Get(1)
_ = Z4287
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4288 := __e.Get(1)
_ = Z4288
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z4287, Z4288)
}
__typedArg0 := Z4287
__typedArg1 := Z4288
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(PrimFunc(symshen_4dict_1fold), tmp18888, V4285, Nil)
return


}, 1)

__e.TailApply(ns2_1set, symshen_4dict_1values, tmp18887)
return




}, 0)

