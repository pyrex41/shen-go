package main

import . "github.com/pyrex41/shen-go/kl"

var DictMain = MakeNative(func(__e *ControlFlow) {
tmp18796 := MakeNative(func(__e *ControlFlow) {
V4205 := __e.Get(1)
_ = V4205
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5) {
__typedN0, __typedOK0 := TypedFloat64(V4205)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5) {
return TypedMaterializeBoolean((__typedN0 < __typedN1))
}}
__typedArg0 := V4205
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5), __typedArg0, __typedArg1)
})() {
tmp18797 := Call(__e, PrimFunc(symshen_4app), V4205, MakeString(""), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("invalid initial dict size: "))
__typedS1, __typedOK1 := TypedString(tmp18797)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("invalid initial dict size: ")
__typedArg1 := tmp18797
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("invalid initial dict size: "))
__typedS1, __typedOK1 := TypedString(tmp18797)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("invalid initial dict size: ")
__typedArg1 := tmp18797
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
tmp18800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(3))
__typedN1, __typedOK1 := TypedFloat64(V4205)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(3)
__typedArg1 := V4205
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(3))
__typedN1, __typedOK1 := TypedFloat64(V4205)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(3)
__typedArg1 := V4205
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

let__18774 := tmp18800
_ = let__18774

tmp18801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__18774, MakeInteger(0), symshen_4dictionary)
}
__typedArg0 := let__18774
__typedArg1 := MakeInteger(0)
__typedArg2 := symshen_4dictionary
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__18775 := tmp18801
_ = let__18775

tmp18802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__18774, MakeInteger(1), V4205)
}
__typedArg0 := let__18774
__typedArg1 := MakeInteger(1)
__typedArg2 := V4205
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__18776 := tmp18802
_ = let__18776

tmp18803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__18774, MakeInteger(2), MakeInteger(0))
}
__typedArg0 := let__18774
__typedArg1 := MakeInteger(2)
__typedArg2 := MakeInteger(0)
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__18777 := tmp18803
_ = let__18777

tmp18805 := Call(__e, PrimFunc(symshen_4fillvector), let__18774, MakeInteger(3), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(2))
__typedN1, __typedOK1 := TypedFloat64(V4205)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(2)
__typedArg1 := V4205
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})(), Nil)


let__18778 := tmp18805
_ = let__18778

__e.Return(let__18774)
return


}


}, 1)

tmp18808 := Call(__e, ns2_1set, symshen_4dict, tmp18796)


_ = tmp18808

tmp18809 := MakeNative(func(__e *ControlFlow) {
V4211 := __e.Get(1)
_ = V4211
tmp18816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V4211)
}
__typedArg0 := V4211
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp18816 {
tmp18811 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V4211, MakeInteger(0))
}
__typedArg0 := V4211
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 0)

tmp18812 := MakeNative(func(__e *ControlFlow) {
Z4212 := __e.Get(1)
_ = Z4212
__e.Return(symshen_4not_1dictionary)
return
}, 1)

tmp18813 := Call(__e, try_1catch, tmp18811, tmp18812)


tmp18814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp18813, symshen_4dictionary)
}
__typedArg0 := tmp18813
__typedArg1 := symshen_4dictionary
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18814 {
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

tmp18817 := Call(__e, ns2_1set, symshen_4dict_2, tmp18809)


_ = tmp18817

tmp18818 := MakeNative(func(__e *ControlFlow) {
V4213 := __e.Get(1)
_ = V4213
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V4213, MakeInteger(1))
}
__typedArg0 := V4213
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp18819 := Call(__e, ns2_1set, symshen_4dict_1capacity, tmp18818)


_ = tmp18819

tmp18820 := MakeNative(func(__e *ControlFlow) {
V4214 := __e.Get(1)
_ = V4214
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V4214, MakeInteger(2))
}
__typedArg0 := V4214
__typedArg1 := MakeInteger(2)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp18821 := Call(__e, ns2_1set, symshen_4dict_1count, tmp18820)


_ = tmp18821

tmp18822 := MakeNative(func(__e *ControlFlow) {
V4215 := __e.Get(1)
_ = V4215
V4216 := __e.Get(2)
_ = V4216
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V4215, MakeInteger(2), V4216)
}
__typedArg0 := V4215
__typedArg1 := MakeInteger(2)
__typedArg2 := V4216
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return
}, 2)

tmp18823 := Call(__e, ns2_1set, symshen_4dict_1count_1_6, tmp18822)


_ = tmp18823

tmp18824 := MakeNative(func(__e *ControlFlow) {
V4217 := __e.Get(1)
_ = V4217
V4218 := __e.Get(2)
_ = V4218
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V4217, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(3))
__typedN1, __typedOK1 := TypedFloat64(V4218)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(3)
__typedArg1 := V4218
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := V4217
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(3))
__typedN1, __typedOK1 := TypedFloat64(V4218)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(3)
__typedArg1 := V4218
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp18826 := Call(__e, ns2_1set, symshen_4_5_1dict_1bucket, tmp18824)


_ = tmp18826

tmp18827 := MakeNative(func(__e *ControlFlow) {
V4219 := __e.Get(1)
_ = V4219
V4220 := __e.Get(2)
_ = V4220
V4221 := __e.Get(3)
_ = V4221
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V4219, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(3))
__typedN1, __typedOK1 := TypedFloat64(V4220)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(3)
__typedArg1 := V4220
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})(), V4221)
}
__typedArg0 := V4219
__typedArg1 := (func() Obj {
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
__typedArg2 := V4221
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return


}, 3)

tmp18829 := Call(__e, ns2_1set, symshen_4dict_1bucket_1_6, tmp18827)


_ = tmp18829

tmp18830 := MakeNative(func(__e *ControlFlow) {
V4222 := __e.Get(1)
_ = V4222
V4223 := __e.Get(2)
_ = V4223
V4224 := __e.Get(3)
_ = V4224
tmp18831 := Call(__e, PrimFunc(symlength), V4224)


tmp18832 := Call(__e, PrimFunc(symlength), V4223)


let__18779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(tmp18831)
__typedN1, __typedOK1 := TypedFloat64(tmp18832)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := tmp18831
__typedArg1 := tmp18832
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})()
_ = let__18779

tmp18834 := Call(__e, PrimFunc(symshen_4dict_1count), V4222)


__e.TailApply(PrimFunc(symshen_4dict_1count_1_6), V4222, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(let__18779)
__typedN1, __typedOK1 := TypedFloat64(tmp18834)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := let__18779
__typedArg1 := tmp18834
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


}, 3)

tmp18836 := Call(__e, ns2_1set, symshen_4dict_1update_1count, tmp18830)


_ = tmp18836

tmp18837 := MakeNative(func(__e *ControlFlow) {
V4226 := __e.Get(1)
_ = V4226
V4227 := __e.Get(2)
_ = V4227
V4228 := __e.Get(3)
_ = V4228
tmp18838 := Call(__e, PrimFunc(symshen_4dict_1capacity), V4226)


tmp18839 := Call(__e, PrimFunc(symhash), V4227, tmp18838)


let__18780 := tmp18839
_ = let__18780

tmp18840 := Call(__e, PrimFunc(symshen_4_5_1dict_1bucket), V4226, let__18780)


let__18781 := tmp18840
_ = let__18781

tmp18841 := Call(__e, PrimFunc(symshen_4assoc_1set), V4227, V4228, let__18781)


let__18782 := tmp18841
_ = let__18782

tmp18842 := Call(__e, PrimFunc(symshen_4dict_1bucket_1_6), V4226, let__18780, let__18782)


let__18783 := tmp18842
_ = let__18783

tmp18843 := Call(__e, PrimFunc(symshen_4dict_1update_1count), V4226, let__18781, let__18782)


let__18784 := tmp18843
_ = let__18784

__e.Return(V4228)
return


}, 3)

tmp18844 := Call(__e, ns2_1set, symshen_4dict_1_6, tmp18837)


_ = tmp18844

tmp18845 := MakeNative(func(__e *ControlFlow) {
V4234 := __e.Get(1)
_ = V4234
V4235 := __e.Get(2)
_ = V4235
tmp18846 := Call(__e, PrimFunc(symshen_4dict_1capacity), V4234)


tmp18847 := Call(__e, PrimFunc(symhash), V4235, tmp18846)


let__18785 := tmp18847
_ = let__18785

tmp18848 := Call(__e, PrimFunc(symshen_4_5_1dict_1bucket), V4234, let__18785)


let__18786 := tmp18848
_ = let__18786

tmp18849 := Call(__e, PrimFunc(symassoc), V4235, let__18786)


let__18787 := tmp18849
_ = let__18787

tmp18853 := Call(__e, PrimFunc(symempty_2), let__18787)


if True == tmp18853 {
tmp18850 := Call(__e, PrimFunc(symshen_4app), V4235, MakeString(" not found in dict\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("value "))
__typedS1, __typedOK1 := TypedString(tmp18850)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("value ")
__typedArg1 := tmp18850
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("value "))
__typedS1, __typedOK1 := TypedString(tmp18850)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("value ")
__typedArg1 := tmp18850
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__18787)
}
__typedArg0 := let__18787
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
}


}, 2)

tmp18854 := Call(__e, ns2_1set, symshen_4_5_1dict, tmp18845)


_ = tmp18854

tmp18855 := MakeNative(func(__e *ControlFlow) {
V4239 := __e.Get(1)
_ = V4239
V4240 := __e.Get(2)
_ = V4240
tmp18856 := Call(__e, PrimFunc(symshen_4dict_1capacity), V4239)


tmp18857 := Call(__e, PrimFunc(symhash), V4240, tmp18856)


let__18788 := tmp18857
_ = let__18788

tmp18858 := Call(__e, PrimFunc(symshen_4_5_1dict_1bucket), V4239, let__18788)


let__18789 := tmp18858
_ = let__18789

tmp18859 := Call(__e, PrimFunc(symshen_4assoc_1rm), V4240, let__18789)


let__18790 := tmp18859
_ = let__18790

tmp18860 := Call(__e, PrimFunc(symshen_4dict_1bucket_1_6), V4239, let__18788, let__18790)


let__18791 := tmp18860
_ = let__18791

tmp18861 := Call(__e, PrimFunc(symshen_4dict_1update_1count), V4239, let__18789, let__18790)


let__18792 := tmp18861
_ = let__18792

__e.Return(V4240)
return


}, 2)

tmp18862 := Call(__e, ns2_1set, symshen_4dict_1rm, tmp18855)


_ = tmp18862

tmp18863 := MakeNative(func(__e *ControlFlow) {
V4246 := __e.Get(1)
_ = V4246
V4247 := __e.Get(2)
_ = V4247
V4248 := __e.Get(3)
_ = V4248
tmp18864 := Call(__e, PrimFunc(symshen_4dict_1capacity), V4247)


let__18793 := tmp18864
_ = let__18793

__e.TailApply(PrimFunc(symshen_4dict_1fold_1h), V4246, V4247, V4248, MakeInteger(0), let__18793)
return


}, 3)

tmp18865 := Call(__e, ns2_1set, symshen_4dict_1fold, tmp18863)


_ = tmp18865

tmp18866 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
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
tmp18871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V4254, V4255)
}
__typedArg0 := V4254
__typedArg1 := V4255
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18871 {
__e.Return(V4253)
return
} else {
tmp18867 := Call(__e, PrimFunc(symshen_4_5_1dict_1bucket), V4252, V4254)


let__18794 := tmp18867
_ = let__18794

tmp18868 := Call(__e, PrimFunc(symshen_4bucket_1fold), V4251, let__18794, V4253)


let__18795 := tmp18868
_ = let__18795

if PrimFunc(symshen_4dict_1fold_1h) == __self {
__self1, __self2, __self3, __self4, __self5 = V4251, V4252, let__18795, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(V4254)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := V4254
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})(), V4255
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4dict_1fold_1h), V4251, V4252, let__18795, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(V4254)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := V4254
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})(), V4255)
return


}


}, 5)

tmp18872 := Call(__e, ns2_1set, symshen_4dict_1fold_1h, tmp18866)


_ = tmp18872

tmp18873 := MakeNative(func(__e *ControlFlow) {
V4258 := __e.Get(1)
_ = V4258
V4259 := __e.Get(2)
_ = V4259
V4260 := __e.Get(3)
_ = V4260
tmp18889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4259)
}
__typedArg0 := Nil
__typedArg1 := V4259
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp18889 {
__e.Return(V4260)
return
} else {
tmp18887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4259)
}
__typedArg0 := V4259
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18883 Obj

if True == tmp18887 {
tmp18885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4259)
}
__typedArg0 := V4259
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp18885)
}
__typedArg0 := tmp18885
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres18884 Obj

if True == tmp18886 {
ifres18884 = True


} else {
ifres18884 = False


}

ifres18883 = ifres18884


} else {
ifres18883 = False


}

if True == ifres18883 {
tmp18874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4259)
}
__typedArg0 := V4259
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp18874)
}
__typedArg0 := tmp18874
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18876 := Call(__e, V4258, tmp18875)


tmp18877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4259)
}
__typedArg0 := V4259
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp18878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp18877)
}
__typedArg0 := tmp18877
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18879 := Call(__e, tmp18876, tmp18878)


tmp18880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4259)
}
__typedArg0 := V4259
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp18881 := Call(__e, PrimFunc(symshen_4bucket_1fold), V4258, tmp18880, V4260)


__e.TailApply(tmp18879, tmp18881)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4bucket_1fold)
return
}


}


}, 3)

tmp18890 := Call(__e, ns2_1set, symshen_4bucket_1fold, tmp18873)


_ = tmp18890

tmp18891 := MakeNative(func(__e *ControlFlow) {
V4261 := __e.Get(1)
_ = V4261
tmp18892 := MakeNative(func(__e *ControlFlow) {
Z4262 := __e.Get(1)
_ = Z4262
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4263 := __e.Get(1)
_ = Z4263
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4264 := __e.Get(1)
_ = Z4264
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z4262, Z4264)
}
__typedArg0 := Z4262
__typedArg1 := Z4264
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(PrimFunc(symshen_4dict_1fold), tmp18892, V4261, Nil)
return


}, 1)

tmp18893 := Call(__e, ns2_1set, symshen_4dict_1keys, tmp18891)


_ = tmp18893

tmp18894 := MakeNative(func(__e *ControlFlow) {
V4265 := __e.Get(1)
_ = V4265
tmp18895 := MakeNative(func(__e *ControlFlow) {
Z4266 := __e.Get(1)
_ = Z4266
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4267 := __e.Get(1)
_ = Z4267
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4268 := __e.Get(1)
_ = Z4268
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Z4267, Z4268)
}
__typedArg0 := Z4267
__typedArg1 := Z4268
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(PrimFunc(symshen_4dict_1fold), tmp18895, V4265, Nil)
return


}, 1)

__e.TailApply(ns2_1set, symshen_4dict_1values, tmp18894)
return




}, 0)

