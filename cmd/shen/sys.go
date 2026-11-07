package main

import . "github.com/pyrex41/shen-go/kl"

var SysMain = MakeNative(func(__e *ControlFlow) {
tmp1210 := MakeNative(func(__e *ControlFlow) {
V3830 := __e.Get(1)
_ = V3830
__e.TailApply(V3830)
return
}, 1)

tmp1211 := Call(__e, ns2_1set, symthaw, tmp1210)


_ = tmp1211

tmp1212 := MakeNative(func(__e *ControlFlow) {
V3831 := __e.Get(1)
_ = V3831
tmp1213 := Call(__e, PrimFunc(symmacroexpand), V3831)


tmp1214 := Call(__e, PrimFunc(symshen_4find_1types), V3831)


tmp1215 := Call(__e, PrimFunc(symshen_4process_1applications), tmp1213, tmp1214)


tmp1216 := Call(__e, PrimFunc(symshen_4shen_1_6kl), tmp1215)


__e.TailApply(PrimFunc(symeval_1kl), tmp1216)
return


}, 1)

tmp1217 := Call(__e, ns2_1set, symeval, tmp1212)


_ = tmp1217

tmp1218 := MakeNative(func(__e *ControlFlow) {
V3832 := __e.Get(1)
_ = V3832
tmp1225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnull, V3832)
}
__typedArg0 := symnull
__typedArg1 := V3832
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1225 {
__e.Return(Nil)
return
} else {
tmp1219 := MakeNative(func(__e *ControlFlow) {
tmp1220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V3832, symshen_4external_1symbols, tmp1220)
return


}, 0)

tmp1221 := MakeNative(func(__e *ControlFlow) {
Z3833 := __e.Get(1)
_ = Z3833
tmp1222 := Call(__e, PrimFunc(symshen_4app), V3832, MakeString(" does not exist.\n;"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("package "))
__typedS1, __typedOK1 := TypedString(tmp1222)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("package ")
__typedArg1 := tmp1222
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("package "))
__typedS1, __typedOK1 := TypedString(tmp1222)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("package ")
__typedArg1 := tmp1222
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

__e.TailApply(try_1catch, tmp1219, tmp1221)
return


}


}, 1)

tmp1226 := Call(__e, ns2_1set, symexternal, tmp1218)


_ = tmp1226

tmp1227 := MakeNative(func(__e *ControlFlow) {
V3834 := __e.Get(1)
_ = V3834
tmp1234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnull, V3834)
}
__typedArg0 := symnull
__typedArg1 := V3834
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1234 {
__e.Return(Nil)
return
} else {
tmp1228 := MakeNative(func(__e *ControlFlow) {
tmp1229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V3834, symshen_4internal_1symbols, tmp1229)
return


}, 0)

tmp1230 := MakeNative(func(__e *ControlFlow) {
Z3835 := __e.Get(1)
_ = Z3835
tmp1231 := Call(__e, PrimFunc(symshen_4app), V3834, MakeString(" does not exist.\n;"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("package "))
__typedS1, __typedOK1 := TypedString(tmp1231)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("package ")
__typedArg1 := tmp1231
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("package "))
__typedS1, __typedOK1 := TypedString(tmp1231)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("package ")
__typedArg1 := tmp1231
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

__e.TailApply(try_1catch, tmp1228, tmp1230)
return


}


}, 1)

tmp1235 := Call(__e, ns2_1set, syminternal, tmp1227)


_ = tmp1235

tmp1236 := MakeNative(func(__e *ControlFlow) {
V3836 := __e.Get(1)
_ = V3836
V3837 := __e.Get(2)
_ = V3837
tmp1238 := Call(__e, V3836, V3837)


if True == tmp1238 {
__e.TailApply(PrimFunc(symfail))
return
} else {
__e.Return(V3837)
return
}


}, 2)

tmp1239 := Call(__e, ns2_1set, symfail_1if, tmp1236)


_ = tmp1239

tmp1240 := MakeNative(func(__e *ControlFlow) {
V3838 := __e.Get(1)
_ = V3838
V3839 := __e.Get(2)
_ = V3839
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(V3838)
__typedS1, __typedOK1 := TypedString(V3839)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V3838
__typedArg1 := V3839
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return
}, 2)

tmp1241 := Call(__e, ns2_1set, sym_8s, tmp1240)


_ = tmp1241

tmp1242 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtc_d)
}
__typedArg0 := symshen_4_dtc_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1243 := Call(__e, ns2_1set, symtc_2, tmp1242)


_ = tmp1243

tmp1244 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_doccurs_d)
}
__typedArg0 := symshen_4_doccurs_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1245 := Call(__e, ns2_1set, symoccurs_2, tmp1244)


_ = tmp1245

tmp1246 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dfactorise_2_d)
}
__typedArg0 := symshen_4_dfactorise_2_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1247 := Call(__e, ns2_1set, symfactorise_2, tmp1246)


_ = tmp1247

tmp1248 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtracking_d)
}
__typedArg0 := symshen_4_dtracking_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1249 := Call(__e, ns2_1set, symtracked, tmp1248)


_ = tmp1249

tmp1250 := MakeNative(func(__e *ControlFlow) {
V3840 := __e.Get(1)
_ = V3840
tmp1251 := MakeNative(func(__e *ControlFlow) {
tmp1252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V3840, symshen_4source, tmp1252)
return


}, 0)

tmp1253 := MakeNative(func(__e *ControlFlow) {
Z3841 := __e.Get(1)
_ = Z3841
tmp1254 := Call(__e, PrimFunc(symshen_4app), V3840, MakeString(" not found.\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp1254)
}
__typedArg0 := tmp1254
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

__e.TailApply(try_1catch, tmp1251, tmp1253)
return


}, 1)

tmp1255 := Call(__e, ns2_1set, symps, tmp1250)


_ = tmp1255

tmp1256 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dstinput_d)
}
__typedArg0 := sym_dstinput_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1257 := Call(__e, ns2_1set, symstinput, tmp1256)


_ = tmp1257

tmp1258 := MakeNative(func(__e *ControlFlow) {
V3842 := __e.Get(1)
_ = V3842
tmp1260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V3842)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V3842
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V3842)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V3842
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

let__1164 := tmp1260
_ = let__1164

tmp1261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__1164, MakeInteger(0), V3842)
}
__typedArg0 := let__1164
__typedArg1 := MakeInteger(0)
__typedArg2 := V3842
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__1165 := tmp1261
_ = let__1165

tmp1265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3842, MakeInteger(0))
}
__typedArg0 := V3842
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1262 Obj

if True == tmp1265 {
ifres1262 = let__1165


} else {
tmp1263 := Call(__e, PrimFunc(symfail))


tmp1264 := Call(__e, PrimFunc(symshen_4fillvector), let__1165, MakeInteger(1), V3842, tmp1263)


ifres1262 = tmp1264


}

let__1166 := ifres1262
_ = let__1166

__e.Return(let__1166)
return


}, 1)

tmp1266 := Call(__e, ns2_1set, symvector, tmp1258)


_ = tmp1266

tmp1267 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__selftop:
V3847 := __self1
_ = V3847
V3848 := __self2
_ = V3848
V3849 := __self3
_ = V3849
V3850 := __self4
_ = V3850
tmp1271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3848, V3849)
}
__typedArg0 := V3848
__typedArg1 := V3849
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1271 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V3847, V3849, V3850)
}
__typedArg0 := V3847
__typedArg1 := V3849
__typedArg2 := V3850
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return
} else {
tmp1268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V3847, V3848, V3850)
}
__typedArg0 := V3847
__typedArg1 := V3848
__typedArg2 := V3850
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

if PrimFunc(symshen_4fillvector) == __self {
__self1, __self2, __self3, __self4 = tmp1268, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(V3848)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := V3848
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})(), V3849, V3850
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4fillvector), tmp1268, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(V3848)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := V3848
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})(), V3849, V3850)
return


}


}, 4)

tmp1272 := Call(__e, ns2_1set, symshen_4fillvector, tmp1267)


_ = tmp1272

tmp1273 := MakeNative(func(__e *ControlFlow) {
V3851 := __e.Get(1)
_ = V3851
tmp1283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V3851)
}
__typedArg0 := V3851
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp1283 {
tmp1275 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V3851, MakeInteger(0))
}
__typedArg0 := V3851
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 0)

tmp1276 := MakeNative(func(__e *ControlFlow) {
Z3853 := __e.Get(1)
_ = Z3853
__e.Return(MakeInteger(-1))
return
}, 1)

tmp1277 := Call(__e, try_1catch, tmp1275, tmp1276)


let__1167 := tmp1277
_ = let__1167

tmp1281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(let__1167)
}
__typedArg0 := let__1167
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

var ifres1278 Obj

if True == tmp1281 {
tmp1280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6_a) {
__typedN0, __typedOK0 := TypedFloat64(let__1167)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(0))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6_a) {
return TypedMaterializeBoolean((__typedN0 >= __typedN1))
}}
__typedArg0 := let__1167
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_6_a), __typedArg0, __typedArg1)
})()

var ifres1279 Obj

if True == tmp1280 {
ifres1279 = True


} else {
ifres1279 = False


}

ifres1278 = ifres1279


} else {
ifres1278 = False


}

if True == ifres1278 {
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

tmp1284 := Call(__e, ns2_1set, symvector_2, tmp1273)


_ = tmp1284

tmp1285 := MakeNative(func(__e *ControlFlow) {
V3854 := __e.Get(1)
_ = V3854
V3855 := __e.Get(2)
_ = V3855
V3856 := __e.Get(3)
_ = V3856
tmp1287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3855, MakeInteger(0))
}
__typedArg0 := V3855
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1287 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("cannot access 0th element of a vector\n"))
}
__typedArg0 := MakeString("cannot access 0th element of a vector\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V3854, V3855, V3856)
}
__typedArg0 := V3854
__typedArg1 := V3855
__typedArg2 := V3856
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return
}


}, 3)

tmp1288 := Call(__e, ns2_1set, symvector_1_6, tmp1285)


_ = tmp1288

tmp1289 := MakeNative(func(__e *ControlFlow) {
V3857 := __e.Get(1)
_ = V3857
V3858 := __e.Get(2)
_ = V3858
tmp1295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3858, MakeInteger(0))
}
__typedArg0 := V3858
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1295 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("cannot access 0th element of a vector\n"))
}
__typedArg0 := MakeString("cannot access 0th element of a vector\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp1290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V3857, V3858)
}
__typedArg0 := V3857
__typedArg1 := V3858
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

let__1168 := tmp1290
_ = let__1168

tmp1292 := Call(__e, PrimFunc(symfail))


tmp1293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__1168, tmp1292)
}
__typedArg0 := let__1168
__typedArg1 := tmp1292
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1293 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("vector element not found\n"))
}
__typedArg0 := MakeString("vector element not found\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
__e.Return(let__1168)
return
}


}


}, 2)

tmp1296 := Call(__e, ns2_1set, sym_5_1vector, tmp1289)


_ = tmp1296

tmp1297 := MakeNative(func(__e *ControlFlow) {
V3860 := __e.Get(1)
_ = V3860
tmp1301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V3860)
}
__typedArg0 := V3860
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})()

if True == tmp1301 {
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6_a) {
__typedN0, __typedOK0 := TypedFloat64(V3860)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(0))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6_a) {
return TypedMaterializeBoolean((__typedN0 >= __typedN1))
}}
__typedArg0 := V3860
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_6_a), __typedArg0, __typedArg1)
})() {
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

tmp1302 := Call(__e, ns2_1set, symshen_4posint_2, tmp1297)


_ = tmp1302

tmp1303 := MakeNative(func(__e *ControlFlow) {
V3861 := __e.Get(1)
_ = V3861
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V3861, MakeInteger(0))
}
__typedArg0 := V3861
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp1304 := Call(__e, ns2_1set, symlimit, tmp1303)


_ = tmp1304

tmp1305 := MakeNative(func(__e *ControlFlow) {
V3862 := __e.Get(1)
_ = V3862
tmp1335 := Call(__e, PrimFunc(symboolean_2), V3862)


var ifres1320 Obj

if True == tmp1335 {
ifres1320 = True


} else {
tmp1334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(V3862)
}
__typedArg0 := V3862
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

var ifres1322 Obj

if True == tmp1334 {
ifres1322 = True


} else {
tmp1333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V3862)
}
__typedArg0 := V3862
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres1324 Obj

if True == tmp1333 {
ifres1324 = True


} else {
tmp1332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3862)
}
__typedArg0 := V3862
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1326 Obj

if True == tmp1332 {
ifres1326 = True


} else {
tmp1331 := Call(__e, PrimFunc(symempty_2), V3862)


var ifres1328 Obj

if True == tmp1331 {
ifres1328 = True


} else {
tmp1330 := Call(__e, PrimFunc(symvector_2), V3862)


var ifres1329 Obj

if True == tmp1330 {
ifres1329 = True


} else {
ifres1329 = False


}

ifres1328 = ifres1329


}

var ifres1327 Obj

if True == ifres1328 {
ifres1327 = True


} else {
ifres1327 = False


}

ifres1326 = ifres1327


}

var ifres1325 Obj

if True == ifres1326 {
ifres1325 = True


} else {
ifres1325 = False


}

ifres1324 = ifres1325


}

var ifres1323 Obj

if True == ifres1324 {
ifres1323 = True


} else {
ifres1323 = False


}

ifres1322 = ifres1323


}

var ifres1321 Obj

if True == ifres1322 {
ifres1321 = True


} else {
ifres1321 = False


}

ifres1320 = ifres1321


}

if True == ifres1320 {
__e.Return(False)
return
} else {
tmp1310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp1311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp1312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(","))
}
__typedArg0 := MakeString(",")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp1313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1312, Nil)
}
__typedArg0 := tmp1312
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1311, tmp1313)
}
__typedArg0 := tmp1311
__typedArg1 := tmp1313
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1310, tmp1314)
}
__typedArg0 := tmp1310
__typedArg1 := tmp1314
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, tmp1315)
}
__typedArg0 := sym_j
__typedArg1 := tmp1315
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, tmp1316)
}
__typedArg0 := sym_i
__typedArg1 := tmp1316
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1318 := Call(__e, PrimFunc(symelement_2), V3862, tmp1317)


if True == tmp1318 {
__e.Return(True)
return
} else {
tmp1306 := MakeNative(func(__e *ControlFlow) {
tmp1307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3862)
}
__typedArg0 := V3862
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

let__1169 := tmp1307
_ = let__1169

__e.TailApply(PrimFunc(symshen_4analyse_1symbol_2), let__1169)
return


}, 0)

tmp1308 := MakeNative(func(__e *ControlFlow) {
Z3864 := __e.Get(1)
_ = Z3864
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1306, tmp1308)
return


}


}


}, 1)

tmp1336 := Call(__e, ns2_1set, symsymbol_2, tmp1305)


_ = tmp1336

tmp1337 := MakeNative(func(__e *ControlFlow) {
V3867 := __e.Get(1)
_ = V3867
tmp1346 := Call(__e, PrimFunc(symshen_4_7string_2), V3867)


if True == tmp1346 {
tmp1342 := Call(__e, PrimFunc(symhdstr), V3867)


tmp1343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp1342)
}
__typedArg0 := tmp1342
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp1344 := Call(__e, PrimFunc(symshen_4alpha_2), tmp1343)


if True == tmp1344 {
tmp1340 := Call(__e, PrimFunc(symshen_4alphanums_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3867)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3867
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


if True == tmp1340 {
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


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.analyse-symbol?"))
}
__typedArg0 := MakeString("implementation error in shen.analyse-symbol?")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp1347 := Call(__e, ns2_1set, symshen_4analyse_1symbol_2, tmp1337)


_ = tmp1347

tmp1348 := MakeNative(func(__e *ControlFlow) {
V3870 := __e.Get(1)
_ = V3870
tmp1362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V3870)
}
__typedArg0 := MakeString("")
__typedArg1 := V3870
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1362 {
__e.Return(True)
return
} else {
tmp1360 := Call(__e, PrimFunc(symshen_4_7string_2), V3870)


if True == tmp1360 {
tmp1349 := Call(__e, PrimFunc(symhdstr), V3870)


tmp1350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp1349)
}
__typedArg0 := tmp1349
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

let__1170 := tmp1350
_ = let__1170

tmp1358 := Call(__e, PrimFunc(symshen_4alpha_2), let__1170)


var ifres1355 Obj

if True == tmp1358 {
ifres1355 = True


} else {
tmp1357 := Call(__e, PrimFunc(symshen_4digit_2), let__1170)


var ifres1356 Obj

if True == tmp1357 {
ifres1356 = True


} else {
ifres1356 = False


}

ifres1355 = ifres1356


}

if True == ifres1355 {
tmp1353 := Call(__e, PrimFunc(symshen_4alphanums_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3870)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3870
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


if True == tmp1353 {
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


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.alphanums?"))
}
__typedArg0 := MakeString("implementation error in shen.alphanums?")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp1363 := Call(__e, ns2_1set, symshen_4alphanums_2, tmp1348)


_ = tmp1363

tmp1364 := MakeNative(func(__e *ControlFlow) {
V3872 := __e.Get(1)
_ = V3872
tmp1375 := Call(__e, PrimFunc(symboolean_2), V3872)


var ifres1369 Obj

if True == tmp1375 {
ifres1369 = True


} else {
tmp1374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(V3872)
}
__typedArg0 := V3872
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

var ifres1371 Obj

if True == tmp1374 {
ifres1371 = True


} else {
tmp1373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V3872)
}
__typedArg0 := V3872
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres1372 Obj

if True == tmp1373 {
ifres1372 = True


} else {
ifres1372 = False


}

ifres1371 = ifres1372


}

var ifres1370 Obj

if True == ifres1371 {
ifres1370 = True


} else {
ifres1370 = False


}

ifres1369 = ifres1370


}

if True == ifres1369 {
__e.Return(False)
return
} else {
tmp1365 := MakeNative(func(__e *ControlFlow) {
tmp1366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3872)
}
__typedArg0 := V3872
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

let__1171 := tmp1366
_ = let__1171

__e.TailApply(PrimFunc(symshen_4analyse_1variable_2), let__1171)
return


}, 0)

tmp1367 := MakeNative(func(__e *ControlFlow) {
Z3874 := __e.Get(1)
_ = Z3874
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1365, tmp1367)
return


}


}, 1)

tmp1376 := Call(__e, ns2_1set, symvariable_2, tmp1364)


_ = tmp1376

tmp1377 := MakeNative(func(__e *ControlFlow) {
V3877 := __e.Get(1)
_ = V3877
tmp1386 := Call(__e, PrimFunc(symshen_4_7string_2), V3877)


if True == tmp1386 {
tmp1382 := Call(__e, PrimFunc(symhdstr), V3877)


tmp1383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp1382)
}
__typedArg0 := tmp1382
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp1384 := Call(__e, PrimFunc(symshen_4uppercase_2), tmp1383)


if True == tmp1384 {
tmp1380 := Call(__e, PrimFunc(symshen_4alphanums_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3877)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3877
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


if True == tmp1380 {
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


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.analyse-variable?"))
}
__typedArg0 := MakeString("implementation error in shen.analyse-variable?")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp1387 := Call(__e, ns2_1set, symshen_4analyse_1variable_2, tmp1377)


_ = tmp1387

tmp1388 := MakeNative(func(__e *ControlFlow) {
V3878 := __e.Get(1)
_ = V3878
tmp1389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dgensym_d)
}
__typedArg0 := symshen_4_dgensym_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp1391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dgensym_d, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp1389)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp1389
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := symshen_4_dgensym_d
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp1389)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp1389
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symconcat), V3878, tmp1391)
return


}, 1)

tmp1392 := Call(__e, ns2_1set, symgensym, tmp1388)


_ = tmp1392

tmp1393 := MakeNative(func(__e *ControlFlow) {
V3879 := __e.Get(1)
_ = V3879
V3880 := __e.Get(2)
_ = V3880
tmp1394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3879)
}
__typedArg0 := V3879
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp1395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3880)
}
__typedArg0 := V3880
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp1394)
__typedS1, __typedOK1 := TypedString(tmp1395)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp1394
__typedArg1 := tmp1395
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp1394)
__typedS1, __typedOK1 := TypedString(tmp1395)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp1394
__typedArg1 := tmp1395
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symintern), __typedArg0)
})())
return


}, 2)

tmp1397 := Call(__e, ns2_1set, symconcat, tmp1393)


_ = tmp1397

tmp1398 := MakeNative(func(__e *ControlFlow) {
V3881 := __e.Get(1)
_ = V3881
V3882 := __e.Get(2)
_ = V3882
tmp1399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(MakeInteger(3))
}
__typedArg0 := MakeInteger(3)
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

let__1172 := tmp1399
_ = let__1172

tmp1400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__1172, MakeInteger(0), symshen_4tuple)
}
__typedArg0 := let__1172
__typedArg1 := MakeInteger(0)
__typedArg2 := symshen_4tuple
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__1173 := tmp1400
_ = let__1173

tmp1401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__1172, MakeInteger(1), V3881)
}
__typedArg0 := let__1172
__typedArg1 := MakeInteger(1)
__typedArg2 := V3881
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__1174 := tmp1401
_ = let__1174

tmp1402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__1172, MakeInteger(2), V3882)
}
__typedArg0 := let__1172
__typedArg1 := MakeInteger(2)
__typedArg2 := V3882
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__1175 := tmp1402
_ = let__1175

__e.Return(let__1172)
return


}, 2)

tmp1403 := Call(__e, ns2_1set, sym_8p, tmp1398)


_ = tmp1403

tmp1404 := MakeNative(func(__e *ControlFlow) {
V3887 := __e.Get(1)
_ = V3887
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V3887, MakeInteger(1))
}
__typedArg0 := V3887
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp1405 := Call(__e, ns2_1set, symfst, tmp1404)


_ = tmp1405

tmp1406 := MakeNative(func(__e *ControlFlow) {
V3888 := __e.Get(1)
_ = V3888
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V3888, MakeInteger(2))
}
__typedArg0 := V3888
__typedArg1 := MakeInteger(2)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp1407 := Call(__e, ns2_1set, symsnd, tmp1406)


_ = tmp1407

tmp1408 := MakeNative(func(__e *ControlFlow) {
V3889 := __e.Get(1)
_ = V3889
tmp1415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V3889)
}
__typedArg0 := V3889
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp1415 {
tmp1410 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V3889, MakeInteger(0))
}
__typedArg0 := V3889
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 0)

tmp1411 := MakeNative(func(__e *ControlFlow) {
Z3890 := __e.Get(1)
_ = Z3890
__e.Return(symshen_4not_1tuple)
return
}, 1)

tmp1412 := Call(__e, try_1catch, tmp1410, tmp1411)


tmp1413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4tuple, tmp1412)
}
__typedArg0 := symshen_4tuple
__typedArg1 := tmp1412
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1413 {
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

tmp1416 := Call(__e, ns2_1set, symtuple_2, tmp1408)


_ = tmp1416

tmp1417 := MakeNative(func(__e *ControlFlow) {
V3895 := __e.Get(1)
_ = V3895
V3896 := __e.Get(2)
_ = V3896
tmp1424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3895)
}
__typedArg0 := Nil
__typedArg1 := V3895
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1424 {
__e.Return(V3896)
return
} else {
tmp1422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3895)
}
__typedArg0 := V3895
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1422 {
tmp1418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3895)
}
__typedArg0 := V3895
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3895)
}
__typedArg0 := V3895
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1420 := Call(__e, PrimFunc(symappend), tmp1419, V3896)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1418, tmp1420)
}
__typedArg0 := tmp1418
__typedArg1 := tmp1420
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("attempt to append a non-list"))
}
__typedArg0 := MakeString("attempt to append a non-list")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp1425 := Call(__e, ns2_1set, symappend, tmp1417)


_ = tmp1425

tmp1426 := MakeNative(func(__e *ControlFlow) {
V3897 := __e.Get(1)
_ = V3897
V3898 := __e.Get(2)
_ = V3898
tmp1427 := Call(__e, PrimFunc(symlimit), V3898)


let__1176 := tmp1427
_ = let__1176

tmp1429 := Call(__e, PrimFunc(symvector), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(let__1176)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := let__1176
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())


let__1177 := tmp1429
_ = let__1177

tmp1430 := Call(__e, PrimFunc(symvector_1_6), let__1177, MakeInteger(1), V3897)


let__1178 := tmp1430
_ = let__1178

tmp1432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__1176, MakeInteger(0))
}
__typedArg0 := let__1176
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1432 {
__e.Return(let__1178)
return
} else {
__e.TailApply(PrimFunc(symshen_4_8v_1help), V3898, MakeInteger(1), let__1176, let__1178)
return
}


}, 2)

tmp1433 := Call(__e, ns2_1set, sym_8v, tmp1426)


_ = tmp1433

tmp1434 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__selftop:
V3903 := __self1
_ = V3903
V3904 := __self2
_ = V3904
V3905 := __self3
_ = V3905
V3906 := __self4
_ = V3906
tmp1440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3904, V3905)
}
__typedArg0 := V3904
__typedArg1 := V3905
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1440 {
__e.TailApply(PrimFunc(symshen_4copyfromvector), V3903, V3906, V3905, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V3905)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V3905
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


} else {
tmp1436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V3904)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V3904
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()

tmp1438 := Call(__e, PrimFunc(symshen_4copyfromvector), V3903, V3906, V3904, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V3904)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V3904
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())


if PrimFunc(symshen_4_8v_1help) == __self {
__self1, __self2, __self3, __self4 = V3903, tmp1436, V3905, tmp1438
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4_8v_1help), V3903, tmp1436, V3905, tmp1438)
return


}


}, 4)

tmp1441 := Call(__e, ns2_1set, symshen_4_8v_1help, tmp1434)


_ = tmp1441

tmp1442 := MakeNative(func(__e *ControlFlow) {
V3907 := __e.Get(1)
_ = V3907
V3908 := __e.Get(2)
_ = V3908
V3909 := __e.Get(3)
_ = V3909
V3910 := __e.Get(4)
_ = V3910
tmp1443 := MakeNative(func(__e *ControlFlow) {
tmp1444 := Call(__e, PrimFunc(sym_5_1vector), V3907, V3909)


__e.TailApply(PrimFunc(symvector_1_6), V3908, V3910, tmp1444)
return


}, 0)

tmp1445 := MakeNative(func(__e *ControlFlow) {
Z3911 := __e.Get(1)
_ = Z3911
__e.Return(V3908)
return
}, 1)

__e.TailApply(try_1catch, tmp1443, tmp1445)
return


}, 4)

tmp1446 := Call(__e, ns2_1set, symshen_4copyfromvector, tmp1442)


_ = tmp1446

tmp1447 := MakeNative(func(__e *ControlFlow) {
V3912 := __e.Get(1)
_ = V3912
tmp1448 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(sym_5_1vector), V3912, MakeInteger(1))
return
}, 0)

tmp1449 := MakeNative(func(__e *ControlFlow) {
Z3913 := __e.Get(1)
_ = Z3913
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("hdv needs a non-empty vector as an argument\n"))
}
__typedArg0 := MakeString("hdv needs a non-empty vector as an argument\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}, 1)

__e.TailApply(try_1catch, tmp1448, tmp1449)
return


}, 1)

tmp1450 := Call(__e, ns2_1set, symhdv, tmp1447)


_ = tmp1450

tmp1451 := MakeNative(func(__e *ControlFlow) {
V3914 := __e.Get(1)
_ = V3914
tmp1452 := Call(__e, PrimFunc(symlimit), V3914)


let__1179 := tmp1452
_ = let__1179

tmp1460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__1179, MakeInteger(0))
}
__typedArg0 := let__1179
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1460 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("cannot take the tail of the empty vector\n"))
}
__typedArg0 := MakeString("cannot take the tail of the empty vector\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp1458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__1179, MakeInteger(1))
}
__typedArg0 := let__1179
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1458 {
__e.TailApply(PrimFunc(symvector), MakeInteger(0))
return
} else {
tmp1454 := Call(__e, PrimFunc(symvector), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(let__1179)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := let__1179
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


let__1180 := tmp1454
_ = let__1180

tmp1456 := Call(__e, PrimFunc(symvector), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(let__1179)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := let__1179
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


__e.TailApply(PrimFunc(symshen_4tlv_1help), V3914, MakeInteger(2), let__1179, tmp1456)
return


}


}


}, 1)

tmp1461 := Call(__e, ns2_1set, symtlv, tmp1451)


_ = tmp1461

tmp1462 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__selftop:
V3918 := __self1
_ = V3918
V3919 := __self2
_ = V3919
V3920 := __self3
_ = V3920
V3921 := __self4
_ = V3921
tmp1468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3919, V3920)
}
__typedArg0 := V3919
__typedArg1 := V3920
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1468 {
__e.TailApply(PrimFunc(symshen_4copyfromvector), V3918, V3921, V3920, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V3920)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V3920
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())
return


} else {
tmp1464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V3919)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V3919
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()

tmp1466 := Call(__e, PrimFunc(symshen_4copyfromvector), V3918, V3921, V3919, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V3919)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V3919
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


if PrimFunc(symshen_4tlv_1help) == __self {
__self1, __self2, __self3, __self4 = V3918, tmp1464, V3920, tmp1466
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4tlv_1help), V3918, tmp1464, V3920, tmp1466)
return


}


}, 4)

tmp1469 := Call(__e, ns2_1set, symshen_4tlv_1help, tmp1462)


_ = tmp1469

tmp1470 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V3933 := __self1
_ = V3933
V3934 := __self2
_ = V3934
tmp1486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3934)
}
__typedArg0 := Nil
__typedArg1 := V3934
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1486 {
__e.Return(Nil)
return
} else {
tmp1484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3934)
}
__typedArg0 := V3934
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1475 Obj

if True == tmp1484 {
tmp1482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3934)
}
__typedArg0 := V3934
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp1482)
}
__typedArg0 := tmp1482
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1477 Obj

if True == tmp1483 {
tmp1479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3934)
}
__typedArg0 := V3934
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1479)
}
__typedArg0 := tmp1479
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3933, tmp1480)
}
__typedArg0 := V3933
__typedArg1 := tmp1480
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1478 Obj

if True == tmp1481 {
ifres1478 = True


} else {
ifres1478 = False


}

ifres1477 = ifres1478


} else {
ifres1477 = False


}

var ifres1476 Obj

if True == ifres1477 {
ifres1476 = True


} else {
ifres1476 = False


}

ifres1475 = ifres1476


} else {
ifres1475 = False


}

if True == ifres1475 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3934)
}
__typedArg0 := V3934
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp1473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3934)
}
__typedArg0 := V3934
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1473 {
tmp1471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3934)
}
__typedArg0 := V3934
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symassoc) == __self {
__self1, __self2 = V3933, tmp1471
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symassoc), V3933, tmp1471)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("attempt to search a non-list with assoc\n"))
}
__typedArg0 := MakeString("attempt to search a non-list with assoc\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 2)

tmp1487 := Call(__e, ns2_1set, symassoc, tmp1470)


_ = tmp1487

tmp1488 := MakeNative(func(__e *ControlFlow) {
V3938 := __e.Get(1)
_ = V3938
V3939 := __e.Get(2)
_ = V3939
V3940 := __e.Get(3)
_ = V3940
tmp1511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3940)
}
__typedArg0 := Nil
__typedArg1 := V3940
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1511 {
tmp1489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3938, V3939)
}
__typedArg0 := V3938
__typedArg1 := V3939
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1489, Nil)
}
__typedArg0 := tmp1489
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp1509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3940)
}
__typedArg0 := V3940
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1500 Obj

if True == tmp1509 {
tmp1507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3940)
}
__typedArg0 := V3940
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp1507)
}
__typedArg0 := tmp1507
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1502 Obj

if True == tmp1508 {
tmp1504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3940)
}
__typedArg0 := V3940
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1504)
}
__typedArg0 := tmp1504
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3938, tmp1505)
}
__typedArg0 := V3938
__typedArg1 := tmp1505
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1503 Obj

if True == tmp1506 {
ifres1503 = True


} else {
ifres1503 = False


}

ifres1502 = ifres1503


} else {
ifres1502 = False


}

var ifres1501 Obj

if True == ifres1502 {
ifres1501 = True


} else {
ifres1501 = False


}

ifres1500 = ifres1501


} else {
ifres1500 = False


}

if True == ifres1500 {
tmp1490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3940)
}
__typedArg0 := V3940
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1490)
}
__typedArg0 := tmp1490
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1491, V3939)
}
__typedArg0 := tmp1491
__typedArg1 := V3939
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3940)
}
__typedArg0 := V3940
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1492, tmp1493)
}
__typedArg0 := tmp1492
__typedArg1 := tmp1493
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp1498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3940)
}
__typedArg0 := V3940
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1498 {
tmp1494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3940)
}
__typedArg0 := V3940
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3940)
}
__typedArg0 := V3940
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1496 := Call(__e, PrimFunc(symshen_4assoc_1set), V3938, V3939, tmp1495)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1494, tmp1496)
}
__typedArg0 := tmp1494
__typedArg1 := tmp1496
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4assoc_1set)
return
}


}


}


}, 3)

tmp1512 := Call(__e, ns2_1set, symshen_4assoc_1set, tmp1488)


_ = tmp1512

tmp1513 := MakeNative(func(__e *ControlFlow) {
V3944 := __e.Get(1)
_ = V3944
V3945 := __e.Get(2)
_ = V3945
tmp1531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3945)
}
__typedArg0 := Nil
__typedArg1 := V3945
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1531 {
__e.Return(Nil)
return
} else {
tmp1529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3945)
}
__typedArg0 := V3945
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1520 Obj

if True == tmp1529 {
tmp1527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3945)
}
__typedArg0 := V3945
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp1527)
}
__typedArg0 := tmp1527
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1522 Obj

if True == tmp1528 {
tmp1524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3945)
}
__typedArg0 := V3945
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1524)
}
__typedArg0 := tmp1524
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3944, tmp1525)
}
__typedArg0 := V3944
__typedArg1 := tmp1525
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1523 Obj

if True == tmp1526 {
ifres1523 = True


} else {
ifres1523 = False


}

ifres1522 = ifres1523


} else {
ifres1522 = False


}

var ifres1521 Obj

if True == ifres1522 {
ifres1521 = True


} else {
ifres1521 = False


}

ifres1520 = ifres1521


} else {
ifres1520 = False


}

if True == ifres1520 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3945)
}
__typedArg0 := V3945
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
} else {
tmp1518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3945)
}
__typedArg0 := V3945
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1518 {
tmp1514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3945)
}
__typedArg0 := V3945
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3945)
}
__typedArg0 := V3945
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1516 := Call(__e, PrimFunc(symshen_4assoc_1rm), V3944, tmp1515)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1514, tmp1516)
}
__typedArg0 := tmp1514
__typedArg1 := tmp1516
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4assoc_1rm)
return
}


}


}


}, 2)

tmp1532 := Call(__e, ns2_1set, symshen_4assoc_1rm, tmp1513)


_ = tmp1532

tmp1533 := MakeNative(func(__e *ControlFlow) {
V3948 := __e.Get(1)
_ = V3948
tmp1537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, V3948)
}
__typedArg0 := True
__typedArg1 := V3948
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1537 {
__e.Return(True)
return
} else {
tmp1535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, V3948)
}
__typedArg0 := False
__typedArg1 := V3948
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1535 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp1538 := Call(__e, ns2_1set, symboolean_2, tmp1533)


_ = tmp1538

tmp1539 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V3949 := __self1
_ = V3949
tmp1544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V3949)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V3949
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1544 {
__e.Return(MakeInteger(0))
return
} else {
tmp1540 := Call(__e, PrimFunc(symstoutput))


tmp1541 := Call(__e, PrimFunc(sympr), MakeString("\n"), tmp1540)


_ = tmp1541

if PrimFunc(symnl) == __self {
__self1 = (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V3949)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V3949
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symnl), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V3949)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V3949
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp1545 := Call(__e, ns2_1set, symnl, tmp1539)


_ = tmp1545

tmp1546 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V3956 := __self1
_ = V3956
V3957 := __self2
_ = V3957
tmp1557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3956)
}
__typedArg0 := Nil
__typedArg1 := V3956
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1557 {
__e.Return(Nil)
return
} else {
tmp1555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3956)
}
__typedArg0 := V3956
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1555 {
tmp1552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3956)
}
__typedArg0 := V3956
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1553 := Call(__e, PrimFunc(symelement_2), tmp1552, V3957)


if True == tmp1553 {
tmp1547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3956)
}
__typedArg0 := V3956
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symdifference) == __self {
__self1, __self2 = tmp1547, V3957
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symdifference), tmp1547, V3957)
return


} else {
tmp1548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3956)
}
__typedArg0 := V3956
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3956)
}
__typedArg0 := V3956
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1550 := Call(__e, PrimFunc(symdifference), tmp1549, V3957)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1548, tmp1550)
}
__typedArg0 := tmp1548
__typedArg1 := tmp1550
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("attempt to find the difference with a non-list\n"))
}
__typedArg0 := MakeString("attempt to find the difference with a non-list\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp1558 := Call(__e, ns2_1set, symdifference, tmp1546)


_ = tmp1558

tmp1559 := MakeNative(func(__e *ControlFlow) {
V3958 := __e.Get(1)
_ = V3958
V3959 := __e.Get(2)
_ = V3959
__e.Return(V3959)
return
}, 2)

tmp1560 := Call(__e, ns2_1set, symdo, tmp1559)


_ = tmp1560

tmp1561 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V3971 := __self1
_ = V3971
V3972 := __self2
_ = V3972
tmp1572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3972)
}
__typedArg0 := Nil
__typedArg1 := V3972
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1572 {
__e.Return(False)
return
} else {
tmp1570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3972)
}
__typedArg0 := V3972
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1566 Obj

if True == tmp1570 {
tmp1568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3972)
}
__typedArg0 := V3972
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3971, tmp1568)
}
__typedArg0 := V3971
__typedArg1 := tmp1568
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1567 Obj

if True == tmp1569 {
ifres1567 = True


} else {
ifres1567 = False


}

ifres1566 = ifres1567


} else {
ifres1566 = False


}

if True == ifres1566 {
__e.Return(True)
return
} else {
tmp1564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3972)
}
__typedArg0 := V3972
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1564 {
tmp1562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3972)
}
__typedArg0 := V3972
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symelement_2) == __self {
__self1, __self2 = V3971, tmp1562
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symelement_2), V3971, tmp1562)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("attempt to find an element in a non-list\n"))
}
__typedArg0 := MakeString("attempt to find an element in a non-list\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 2)

tmp1573 := Call(__e, ns2_1set, symelement_2, tmp1561)


_ = tmp1573

tmp1574 := MakeNative(func(__e *ControlFlow) {
V3975 := __e.Get(1)
_ = V3975
tmp1576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3975)
}
__typedArg0 := Nil
__typedArg1 := V3975
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1576 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp1577 := Call(__e, ns2_1set, symempty_2, tmp1574)


_ = tmp1577

tmp1578 := MakeNative(func(__e *ControlFlow) {
V3976 := __e.Get(1)
_ = V3976
V3977 := __e.Get(2)
_ = V3977
tmp1579 := Call(__e, V3976, V3977)


__e.TailApply(PrimFunc(symshen_4fix_1help), V3976, V3977, tmp1579)
return


}, 2)

tmp1580 := Call(__e, ns2_1set, symfix, tmp1578)


_ = tmp1580

tmp1581 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__selftop:
V3983 := __self1
_ = V3983
V3984 := __self2
_ = V3984
V3985 := __self3
_ = V3985
tmp1584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3984, V3985)
}
__typedArg0 := V3984
__typedArg1 := V3985
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1584 {
__e.Return(V3985)
return
} else {
tmp1582 := Call(__e, V3983, V3985)


if PrimFunc(symshen_4fix_1help) == __self {
__self1, __self2, __self3 = V3983, V3985, tmp1582
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4fix_1help), V3983, V3985, tmp1582)
return


}


}, 3)

tmp1585 := Call(__e, ns2_1set, symshen_4fix_1help, tmp1581)


_ = tmp1585

tmp1586 := MakeNative(func(__e *ControlFlow) {
V3986 := __e.Get(1)
_ = V3986
V3987 := __e.Get(2)
_ = V3987
V3988 := __e.Get(3)
_ = V3988
V3989 := __e.Get(4)
_ = V3989
tmp1587 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4_5_1dict), V3989, V3986)
return
}, 0)

tmp1588 := MakeNative(func(__e *ControlFlow) {
Z3991 := __e.Get(1)
_ = Z3991
__e.Return(Nil)
return
}, 1)

tmp1589 := Call(__e, try_1catch, tmp1587, tmp1588)


let__1181 := tmp1589
_ = let__1181

tmp1590 := Call(__e, PrimFunc(symshen_4assoc_1set), V3987, V3988, let__1181)


let__1182 := tmp1590
_ = let__1182

tmp1591 := Call(__e, PrimFunc(symshen_4dict_1_6), V3989, V3986, let__1182)


let__1183 := tmp1591
_ = let__1183

__e.Return(V3988)
return


}, 4)

tmp1592 := Call(__e, ns2_1set, symput, tmp1586)


_ = tmp1592

tmp1593 := MakeNative(func(__e *ControlFlow) {
V3994 := __e.Get(1)
_ = V3994
V3995 := __e.Get(2)
_ = V3995
V3996 := __e.Get(3)
_ = V3996
tmp1594 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4_5_1dict), V3996, V3994)
return
}, 0)

tmp1595 := MakeNative(func(__e *ControlFlow) {
Z3998 := __e.Get(1)
_ = Z3998
__e.Return(Nil)
return
}, 1)

tmp1596 := Call(__e, try_1catch, tmp1594, tmp1595)


let__1184 := tmp1596
_ = let__1184

tmp1597 := Call(__e, PrimFunc(symshen_4assoc_1rm), V3995, let__1184)


let__1185 := tmp1597
_ = let__1185

tmp1598 := Call(__e, PrimFunc(symshen_4dict_1_6), V3996, V3994, let__1185)


let__1186 := tmp1598
_ = let__1186

__e.Return(V3994)
return


}, 3)

tmp1599 := Call(__e, ns2_1set, symunput, tmp1593)


_ = tmp1599

tmp1600 := MakeNative(func(__e *ControlFlow) {
V4001 := __e.Get(1)
_ = V4001
V4002 := __e.Get(2)
_ = V4002
V4003 := __e.Get(3)
_ = V4003
tmp1601 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4_5_1dict), V4003, V4001)
return
}, 0)

tmp1602 := MakeNative(func(__e *ControlFlow) {
Z4005 := __e.Get(1)
_ = Z4005
tmp1603 := Call(__e, PrimFunc(symshen_4app), V4002, MakeString("\n"), symshen_4s)


tmp1605 := Call(__e, PrimFunc(symshen_4app), V4001, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" has no attributes: "))
__typedS1, __typedOK1 := TypedString(tmp1603)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" has no attributes: ")
__typedArg1 := tmp1603
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp1605)
}
__typedArg0 := tmp1605
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

tmp1606 := Call(__e, try_1catch, tmp1601, tmp1602)


let__1187 := tmp1606
_ = let__1187

tmp1607 := Call(__e, PrimFunc(symassoc), V4002, let__1187)


let__1188 := tmp1607
_ = let__1188

tmp1613 := Call(__e, PrimFunc(symempty_2), let__1188)


if True == tmp1613 {
tmp1608 := Call(__e, PrimFunc(symshen_4app), V4001, MakeString("\n"), symshen_4s)


tmp1610 := Call(__e, PrimFunc(symshen_4app), V4002, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" not found for "))
__typedS1, __typedOK1 := TypedString(tmp1608)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" not found for ")
__typedArg1 := tmp1608
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attribute "))
__typedS1, __typedOK1 := TypedString(tmp1610)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attribute ")
__typedArg1 := tmp1610
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attribute "))
__typedS1, __typedOK1 := TypedString(tmp1610)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attribute ")
__typedArg1 := tmp1610
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__1188)
}
__typedArg0 := let__1188
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
}


}, 3)

tmp1614 := Call(__e, ns2_1set, symget, tmp1600)


_ = tmp1614

tmp1615 := MakeNative(func(__e *ControlFlow) {
V4007 := __e.Get(1)
_ = V4007
V4008 := __e.Get(2)
_ = V4008
tmp1616 := Call(__e, PrimFunc(symshen_4hashkey), V4007)


tmp1617 := Call(__e, PrimFunc(symshen_4mod), tmp1616, V4008)


let__1189 := tmp1617
_ = let__1189

tmp1619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__1189, MakeInteger(0))
}
__typedArg0 := let__1189
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1619 {
__e.Return(MakeInteger(1))
return
} else {
__e.Return(let__1189)
return
}


}, 2)

tmp1620 := Call(__e, ns2_1set, symhash, tmp1615)


_ = tmp1620

tmp1621 := MakeNative(func(__e *ControlFlow) {
V4010 := __e.Get(1)
_ = V4010
tmp1622 := MakeNative(func(__e *ControlFlow) {
Z4012 := __e.Get(1)
_ = Z4012
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(Z4012)
}
__typedArg0 := Z4012
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})())
return
}, 1)

tmp1623 := Call(__e, PrimFunc(symexplode), V4010)


tmp1624 := Call(__e, PrimFunc(symmap), tmp1622, tmp1623)


let__1190 := tmp1624
_ = let__1190

__e.TailApply(PrimFunc(symshen_4prodbutzero), let__1190, MakeInteger(1))
return


}, 1)

tmp1625 := Call(__e, ns2_1set, symshen_4hashkey, tmp1621)


_ = tmp1625

tmp1626 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V4013 := __self1
_ = V4013
V4014 := __self2
_ = V4014
tmp1645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4013)
}
__typedArg0 := Nil
__typedArg1 := V4013
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1645 {
__e.Return(V4014)
return
} else {
tmp1643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4013)
}
__typedArg0 := V4013
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1639 Obj

if True == tmp1643 {
tmp1641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4013)
}
__typedArg0 := V4013
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), tmp1641)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp1641
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1640 Obj

if True == tmp1642 {
ifres1640 = True


} else {
ifres1640 = False


}

ifres1639 = ifres1640


} else {
ifres1639 = False


}

if True == ifres1639 {
tmp1627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4013)
}
__typedArg0 := V4013
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4prodbutzero) == __self {
__self1, __self2 = tmp1627, V4014
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp1627, V4014)
return


} else {
tmp1637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4013)
}
__typedArg0 := V4013
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1637 {
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(V4014)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1e+10))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := V4014
__typedArg1 := MakeNumber(1e+10)
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
tmp1628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4013)
}
__typedArg0 := V4013
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4013)
}
__typedArg0 := V4013
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

if PrimFunc(symshen_4prodbutzero) == __self {
__self1, __self2 = tmp1628, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4014)
__typedN1, __typedOK1 := TypedFloat64(tmp1629)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4014
__typedArg1 := tmp1629
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp1628, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4014)
__typedN1, __typedOK1 := TypedFloat64(tmp1629)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4014
__typedArg1 := tmp1629
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


} else {
tmp1631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4013)
}
__typedArg0 := V4013
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4013)
}
__typedArg0 := V4013
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

if PrimFunc(symshen_4prodbutzero) == __self {
__self1, __self2 = tmp1631, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(V4014)
__typedN1, __typedOK1 := TypedFloat64(tmp1632)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := V4014
__typedArg1 := tmp1632
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp1631, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(V4014)
__typedN1, __typedOK1 := TypedFloat64(tmp1632)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := V4014
__typedArg1 := tmp1632
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})())
return


}


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4prodbutzero)
return
}


}


}


}, 2)

tmp1646 := Call(__e, ns2_1set, symshen_4prodbutzero, tmp1626)


_ = tmp1646

tmp1647 := MakeNative(func(__e *ControlFlow) {
V4015 := __e.Get(1)
_ = V4015
V4016 := __e.Get(2)
_ = V4016
tmp1648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4016, Nil)
}
__typedArg0 := V4016
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1649 := Call(__e, PrimFunc(symshen_4multiples), V4015, tmp1648)


__e.TailApply(PrimFunc(symshen_4modh), V4015, tmp1649)
return


}, 2)

tmp1650 := Call(__e, ns2_1set, symshen_4mod, tmp1647)


_ = tmp1650

tmp1651 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V4021 := __self1
_ = V4021
V4022 := __self2
_ = V4022
tmp1662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4022)
}
__typedArg0 := V4022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1658 Obj

if True == tmp1662 {
tmp1660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4022)
}
__typedArg0 := V4022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(tmp1660)
__typedN1, __typedOK1 := TypedFloat64(V4021)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := tmp1660
__typedArg1 := V4021
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})()

var ifres1659 Obj

if True == tmp1661 {
ifres1659 = True


} else {
ifres1659 = False


}

ifres1658 = ifres1659


} else {
ifres1658 = False


}

if True == ifres1658 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4022)
}
__typedArg0 := V4022
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
} else {
tmp1656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4022)
}
__typedArg0 := V4022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1656 {
tmp1652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4022)
}
__typedArg0 := V4022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(2))
__typedN1, __typedOK1 := TypedFloat64(tmp1652)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp1652
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})(), V4022)
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(2))
__typedN1, __typedOK1 := TypedFloat64(tmp1652)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp1652
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})()
__typedArg1 := V4022
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4multiples) == __self {
__self1, __self2 = V4021, tmp1654
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4multiples), V4021, tmp1654)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.multiples"))
}
__typedArg0 := MakeString("implementation error in shen.multiples")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp1663 := Call(__e, ns2_1set, symshen_4multiples, tmp1651)


_ = tmp1663

tmp1664 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V4029 := __self1
_ = V4029
V4030 := __self2
_ = V4030
tmp1682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V4029)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V4029
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1682 {
__e.Return(MakeInteger(0))
return
} else {
tmp1680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4030)
}
__typedArg0 := Nil
__typedArg1 := V4030
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1680 {
__e.Return(V4029)
return
} else {
tmp1678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4030)
}
__typedArg0 := V4030
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1674 Obj

if True == tmp1678 {
tmp1676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4030)
}
__typedArg0 := V4030
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(tmp1676)
__typedN1, __typedOK1 := TypedFloat64(V4029)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := tmp1676
__typedArg1 := V4029
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})()

var ifres1675 Obj

if True == tmp1677 {
ifres1675 = True


} else {
ifres1675 = False


}

ifres1674 = ifres1675


} else {
ifres1674 = False


}

if True == ifres1674 {
tmp1667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4030)
}
__typedArg0 := V4030
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1668 := Call(__e, PrimFunc(symempty_2), tmp1667)


if True == tmp1668 {
__e.Return(V4029)
return
} else {
tmp1665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4030)
}
__typedArg0 := V4030
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4modh) == __self {
__self1, __self2 = V4029, tmp1665
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4modh), V4029, tmp1665)
return


}


} else {
tmp1672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4030)
}
__typedArg0 := V4030
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1672 {
tmp1669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4030)
}
__typedArg0 := V4030
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

if PrimFunc(symshen_4modh) == __self {
__self1, __self2 = (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V4029)
__typedN1, __typedOK1 := TypedFloat64(tmp1669)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V4029
__typedArg1 := tmp1669
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})(), V4030
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4modh), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V4029)
__typedN1, __typedOK1 := TypedFloat64(tmp1669)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V4029
__typedArg1 := tmp1669
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})(), V4030)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.modh"))
}
__typedArg0 := MakeString("implementation error in shen.modh")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}


}, 2)

tmp1683 := Call(__e, ns2_1set, symshen_4modh, tmp1664)


_ = tmp1683

tmp1684 := MakeNative(func(__e *ControlFlow) {
V4033 := __e.Get(1)
_ = V4033
tmp1691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4033)
}
__typedArg0 := Nil
__typedArg1 := V4033
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1691 {
__e.Return(MakeInteger(0))
return
} else {
tmp1689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4033)
}
__typedArg0 := V4033
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1689 {
tmp1685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4033)
}
__typedArg0 := V4033
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4033)
}
__typedArg0 := V4033
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1687 := Call(__e, PrimFunc(symsum), tmp1686)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(tmp1685)
__typedN1, __typedOK1 := TypedFloat64(tmp1687)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := tmp1685
__typedArg1 := tmp1687
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("attempt to sum a non-list\n"))
}
__typedArg0 := MakeString("attempt to sum a non-list\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp1692 := Call(__e, ns2_1set, symsum, tmp1684)


_ = tmp1692

tmp1693 := MakeNative(func(__e *ControlFlow) {
V4038 := __e.Get(1)
_ = V4038
tmp1695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4038)
}
__typedArg0 := V4038
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1695 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4038)
}
__typedArg0 := V4038
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("head expects a non-empty list\n"))
}
__typedArg0 := MakeString("head expects a non-empty list\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp1696 := Call(__e, ns2_1set, symhead, tmp1693)


_ = tmp1696

tmp1697 := MakeNative(func(__e *ControlFlow) {
V4043 := __e.Get(1)
_ = V4043
tmp1699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4043)
}
__typedArg0 := V4043
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1699 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4043)
}
__typedArg0 := V4043
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("tail expects a non-empty list\n"))
}
__typedArg0 := MakeString("tail expects a non-empty list\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp1700 := Call(__e, ns2_1set, symtail, tmp1697)


_ = tmp1700

tmp1701 := MakeNative(func(__e *ControlFlow) {
V4044 := __e.Get(1)
_ = V4044
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sympos) {
return PrimPos(V4044, MakeInteger(0))
}
__typedArg0 := V4044
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sympos), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp1702 := Call(__e, ns2_1set, symhdstr, tmp1701)


_ = tmp1702

tmp1703 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V4051 := __self1
_ = V4051
V4052 := __self2
_ = V4052
tmp1714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4051)
}
__typedArg0 := Nil
__typedArg1 := V4051
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1714 {
__e.Return(Nil)
return
} else {
tmp1712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4051)
}
__typedArg0 := V4051
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1712 {
tmp1709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4051)
}
__typedArg0 := V4051
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1710 := Call(__e, PrimFunc(symelement_2), tmp1709, V4052)


if True == tmp1710 {
tmp1704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4051)
}
__typedArg0 := V4051
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4051)
}
__typedArg0 := V4051
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1706 := Call(__e, PrimFunc(symintersection), tmp1705, V4052)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1704, tmp1706)
}
__typedArg0 := tmp1704
__typedArg1 := tmp1706
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp1707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4051)
}
__typedArg0 := V4051
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symintersection) == __self {
__self1, __self2 = tmp1707, V4052
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symintersection), tmp1707, V4052)
return


}


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("attempt to find the intersection with a non-list\n"))
}
__typedArg0 := MakeString("attempt to find the intersection with a non-list\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp1715 := Call(__e, ns2_1set, symintersection, tmp1703)


_ = tmp1715

tmp1716 := MakeNative(func(__e *ControlFlow) {
V4053 := __e.Get(1)
_ = V4053
__e.TailApply(PrimFunc(symshen_4reverse_1help), V4053, Nil)
return
}, 1)

tmp1717 := Call(__e, ns2_1set, symreverse, tmp1716)


_ = tmp1717

tmp1718 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V4058 := __self1
_ = V4058
V4059 := __self2
_ = V4059
tmp1725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4058)
}
__typedArg0 := Nil
__typedArg1 := V4058
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1725 {
__e.Return(V4059)
return
} else {
tmp1723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4058)
}
__typedArg0 := V4058
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1723 {
tmp1719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4058)
}
__typedArg0 := V4058
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4058)
}
__typedArg0 := V4058
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1720, V4059)
}
__typedArg0 := tmp1720
__typedArg1 := V4059
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4reverse_1help) == __self {
__self1, __self2 = tmp1719, tmp1721
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4reverse_1help), tmp1719, tmp1721)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("attempt to reverse a non-list\n"))
}
__typedArg0 := MakeString("attempt to reverse a non-list\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp1726 := Call(__e, ns2_1set, symshen_4reverse_1help, tmp1718)


_ = tmp1726

tmp1727 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V4064 := __self1
_ = V4064
V4065 := __self2
_ = V4065
tmp1738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4064)
}
__typedArg0 := Nil
__typedArg1 := V4064
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1738 {
__e.Return(V4065)
return
} else {
tmp1736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4064)
}
__typedArg0 := V4064
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1736 {
tmp1733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4064)
}
__typedArg0 := V4064
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1734 := Call(__e, PrimFunc(symelement_2), tmp1733, V4065)


if True == tmp1734 {
tmp1728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4064)
}
__typedArg0 := V4064
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symunion) == __self {
__self1, __self2 = tmp1728, V4065
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symunion), tmp1728, V4065)
return


} else {
tmp1729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4064)
}
__typedArg0 := V4064
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4064)
}
__typedArg0 := V4064
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1731 := Call(__e, PrimFunc(symunion), tmp1730, V4065)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1729, tmp1731)
}
__typedArg0 := tmp1729
__typedArg1 := tmp1731
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("attempt to find the union with a non-list\n"))
}
__typedArg0 := MakeString("attempt to find the union with a non-list\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp1739 := Call(__e, ns2_1set, symunion, tmp1727)


_ = tmp1739

tmp1740 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V4066 := __self1
_ = V4066
tmp1741 := Call(__e, PrimFunc(symshen_4proc_1nl), V4066)


tmp1742 := Call(__e, PrimFunc(symstoutput))


tmp1743 := Call(__e, PrimFunc(sympr), tmp1741, tmp1742)


let__1191 := tmp1743
_ = let__1191

tmp1744 := Call(__e, PrimFunc(symstoutput))


tmp1745 := Call(__e, PrimFunc(sympr), MakeString(" (y/n) "), tmp1744)


let__1192 := tmp1745
_ = let__1192

tmp1746 := Call(__e, PrimFunc(symstinput))


tmp1747 := Call(__e, PrimFunc(symread), tmp1746)


tmp1748 := Call(__e, PrimFunc(symshen_4app), tmp1747, MakeString(""), symshen_4s)


let__1193 := tmp1748
_ = let__1193

tmp1754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("y"), let__1193)
}
__typedArg0 := MakeString("y")
__typedArg1 := let__1193
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1754 {
__e.Return(True)
return
} else {
tmp1752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("n"), let__1193)
}
__typedArg0 := MakeString("n")
__typedArg1 := let__1193
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1752 {
__e.Return(False)
return
} else {
tmp1749 := Call(__e, PrimFunc(symstoutput))


tmp1750 := Call(__e, PrimFunc(sympr), MakeString("please answer y or n\n"), tmp1749)


_ = tmp1750

if PrimFunc(symy_1or_1n_2) == __self {
__self1 = V4066
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symy_1or_1n_2), V4066)
return


}


}


}, 1)

tmp1755 := Call(__e, ns2_1set, symy_1or_1n_2, tmp1740)


_ = tmp1755

tmp1756 := MakeNative(func(__e *ControlFlow) {
V4070 := __e.Get(1)
_ = V4070
if True == V4070 {
__e.Return(False)
return
} else {
__e.Return(True)
return
}
}, 1)

tmp1758 := Call(__e, ns2_1set, symnot, tmp1756)


_ = tmp1758

tmp1759 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString(""))
}
__typedArg0 := MakeString("")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}, 0)

tmp1760 := Call(__e, ns2_1set, symabort, tmp1759)


_ = tmp1760

tmp1761 := MakeNative(func(__e *ControlFlow) {
V4076 := __e.Get(1)
_ = V4076
V4077 := __e.Get(2)
_ = V4077
V4078 := __e.Get(3)
_ = V4078
tmp1769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V4077, V4078)
}
__typedArg0 := V4077
__typedArg1 := V4078
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1769 {
__e.Return(V4076)
return
} else {
tmp1767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4078)
}
__typedArg0 := V4078
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1767 {
tmp1762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4078)
}
__typedArg0 := V4078
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1763 := Call(__e, PrimFunc(symsubst), V4076, V4077, tmp1762)


tmp1764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4078)
}
__typedArg0 := V4078
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1765 := Call(__e, PrimFunc(symsubst), V4076, V4077, tmp1764)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1763, tmp1765)
}
__typedArg0 := tmp1763
__typedArg1 := tmp1765
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V4078)
return
}


}


}, 3)

tmp1770 := Call(__e, ns2_1set, symsubst, tmp1761)


_ = tmp1770

tmp1771 := MakeNative(func(__e *ControlFlow) {
V4079 := __e.Get(1)
_ = V4079
tmp1772 := Call(__e, PrimFunc(symshen_4app), V4079, MakeString(""), symshen_4a)


__e.TailApply(PrimFunc(symshen_4explode_1h), tmp1772)
return


}, 1)

tmp1773 := Call(__e, ns2_1set, symexplode, tmp1771)


_ = tmp1773

tmp1774 := MakeNative(func(__e *ControlFlow) {
V4082 := __e.Get(1)
_ = V4082
tmp1781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V4082)
}
__typedArg0 := MakeString("")
__typedArg1 := V4082
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1781 {
__e.Return(Nil)
return
} else {
tmp1779 := Call(__e, PrimFunc(symshen_4_7string_2), V4082)


if True == tmp1779 {
tmp1775 := Call(__e, PrimFunc(symhdstr), V4082)


tmp1777 := Call(__e, PrimFunc(symshen_4explode_1h), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V4082)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V4082
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1775, tmp1777)
}
__typedArg0 := tmp1775
__typedArg1 := tmp1777
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in explode-h"))
}
__typedArg0 := MakeString("implementation error in explode-h")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp1782 := Call(__e, ns2_1set, symshen_4explode_1h, tmp1774)


_ = tmp1782

tmp1783 := MakeNative(func(__e *ControlFlow) {
V4083 := __e.Get(1)
_ = V4083
tmp1786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V4083, MakeString(""))
}
__typedArg0 := V4083
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1784 Obj

if True == tmp1786 {
ifres1784 = MakeString("")


} else {
tmp1785 := Call(__e, PrimFunc(symshen_4app), V4083, MakeString("/"), symshen_4a)


ifres1784 = tmp1785


}

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dhome_1directory_d, ifres1784)
}
__typedArg0 := sym_dhome_1directory_d
__typedArg1 := ifres1784
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp1787 := Call(__e, ns2_1set, symcd, tmp1783)


_ = tmp1787

tmp1788 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V4084 := __self1
_ = V4084
V4085 := __self2
_ = V4085
tmp1795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4085)
}
__typedArg0 := Nil
__typedArg1 := V4085
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1795 {
__e.Return(True)
return
} else {
tmp1793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4085)
}
__typedArg0 := V4085
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1793 {
tmp1789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4085)
}
__typedArg0 := V4085
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1790 := Call(__e, V4084, tmp1789)


let__1194 := tmp1790
_ = let__1194

tmp1791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4085)
}
__typedArg0 := V4085
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4for_1each) == __self {
__self1, __self2 = V4084, tmp1791
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4for_1each), V4084, tmp1791)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4for_1each)
return
}


}


}, 2)

tmp1796 := Call(__e, ns2_1set, symshen_4for_1each, tmp1788)


_ = tmp1796

tmp1797 := MakeNative(func(__e *ControlFlow) {
V4087 := __e.Get(1)
_ = V4087
V4088 := __e.Get(2)
_ = V4088
__e.TailApply(PrimFunc(symshen_4map_1h), V4087, V4088, Nil)
return
}, 2)

tmp1798 := Call(__e, ns2_1set, symmap, tmp1797)


_ = tmp1798

tmp1799 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__selftop:
V4089 := __self1
_ = V4089
V4090 := __self2
_ = V4090
V4091 := __self3
_ = V4091
tmp1807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4090)
}
__typedArg0 := Nil
__typedArg1 := V4090
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1807 {
__e.TailApply(PrimFunc(symreverse), V4091)
return
} else {
tmp1805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4090)
}
__typedArg0 := V4090
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1805 {
tmp1800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4090)
}
__typedArg0 := V4090
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4090)
}
__typedArg0 := V4090
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1802 := Call(__e, V4089, tmp1801)


tmp1803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1802, V4091)
}
__typedArg0 := tmp1802
__typedArg1 := V4091
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4map_1h) == __self {
__self1, __self2, __self3 = V4089, tmp1800, tmp1803
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4map_1h), V4089, tmp1800, tmp1803)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4map_1h)
return
}


}


}, 3)

tmp1808 := Call(__e, ns2_1set, symshen_4map_1h, tmp1799)


_ = tmp1808

tmp1809 := MakeNative(func(__e *ControlFlow) {
V4092 := __e.Get(1)
_ = V4092
__e.TailApply(PrimFunc(symshen_4length_1h), V4092, MakeInteger(0))
return
}, 1)

tmp1810 := Call(__e, ns2_1set, symlength, tmp1809)


_ = tmp1810

tmp1811 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V4097 := __self1
_ = V4097
V4098 := __self2
_ = V4098
tmp1815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4097)
}
__typedArg0 := Nil
__typedArg1 := V4097
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1815 {
__e.Return(V4098)
return
} else {
tmp1812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4097)
}
__typedArg0 := V4097
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4length_1h) == __self {
__self1, __self2 = tmp1812, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4098)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4098
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4length_1h), tmp1812, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4098)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4098
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


}


}, 2)

tmp1816 := Call(__e, ns2_1set, symshen_4length_1h, tmp1811)


_ = tmp1816

tmp1817 := MakeNative(func(__e *ControlFlow) {
V4104 := __e.Get(1)
_ = V4104
V4105 := __e.Get(2)
_ = V4105
tmp1825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V4104, V4105)
}
__typedArg0 := V4104
__typedArg1 := V4105
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1825 {
__e.Return(MakeInteger(1))
return
} else {
tmp1823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4105)
}
__typedArg0 := V4105
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1823 {
tmp1818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4105)
}
__typedArg0 := V4105
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1819 := Call(__e, PrimFunc(symoccurrences), V4104, tmp1818)


tmp1820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4105)
}
__typedArg0 := V4105
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1821 := Call(__e, PrimFunc(symoccurrences), V4104, tmp1820)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(tmp1819)
__typedN1, __typedOK1 := TypedFloat64(tmp1821)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := tmp1819
__typedArg1 := tmp1821
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(MakeInteger(0))
return
}


}


}, 2)

tmp1826 := Call(__e, ns2_1set, symoccurrences, tmp1817)


_ = tmp1826

tmp1827 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V4110 := __self1
_ = V4110
V4111 := __self2
_ = V4111
tmp1840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(1), V4110)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := V4110
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1837 Obj

if True == tmp1840 {
tmp1839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4111)
}
__typedArg0 := V4111
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1838 Obj

if True == tmp1839 {
ifres1838 = True


} else {
ifres1838 = False


}

ifres1837 = ifres1838


} else {
ifres1837 = False


}

if True == ifres1837 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4111)
}
__typedArg0 := V4111
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp1835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4111)
}
__typedArg0 := V4111
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1835 {
tmp1828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V4110)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V4110
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})()

tmp1829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4111)
}
__typedArg0 := V4111
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symnth) == __self {
__self1, __self2 = tmp1828, tmp1829
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symnth), tmp1828, tmp1829)
return


} else {
tmp1830 := Call(__e, PrimFunc(symshen_4app), V4111, MakeString("\n"), symshen_4a)


tmp1832 := Call(__e, PrimFunc(symshen_4app), V4110, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(", "))
__typedS1, __typedOK1 := TypedString(tmp1830)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(", ")
__typedArg1 := tmp1830
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("nth applied to "))
__typedS1, __typedOK1 := TypedString(tmp1832)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("nth applied to ")
__typedArg1 := tmp1832
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("nth applied to "))
__typedS1, __typedOK1 := TypedString(tmp1832)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("nth applied to ")
__typedArg1 := tmp1832
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}


}, 2)

tmp1841 := Call(__e, ns2_1set, symnth, tmp1827)


_ = tmp1841

tmp1842 := MakeNative(func(__e *ControlFlow) {
V4112 := __e.Get(1)
_ = V4112
tmp1848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(V4112)
}
__typedArg0 := V4112
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

if True == tmp1848 {
tmp1844 := Call(__e, PrimFunc(symshen_4abs), V4112)


let__1195 := tmp1844
_ = let__1195

tmp1845 := Call(__e, PrimFunc(symshen_4magless), let__1195, MakeInteger(1))


tmp1846 := Call(__e, PrimFunc(symshen_4integer_1test_2), let__1195, tmp1845)


if True == tmp1846 {
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

tmp1849 := Call(__e, ns2_1set, syminteger_2, tmp1842)


_ = tmp1849

tmp1850 := MakeNative(func(__e *ControlFlow) {
V4114 := __e.Get(1)
_ = V4114
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(V4114)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(0))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := V4114
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
__e.Return(V4114)
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(0))
__typedN1, __typedOK1 := TypedFloat64(V4114)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := MakeInteger(0)
__typedArg1 := V4114
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())
return
}


}, 1)

tmp1853 := Call(__e, ns2_1set, symshen_4abs, tmp1850)


_ = tmp1853

tmp1854 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V4115 := __self1
_ = V4115
V4116 := __self2
_ = V4116
let__1196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(V4116)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(2))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := V4116
__typedArg1 := MakeInteger(2)
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})()
_ = let__1196

if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(let__1196)
__typedN1, __typedOK1 := TypedFloat64(V4115)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := let__1196
__typedArg1 := V4115
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
__e.Return(V4116)
return
} else {
if PrimFunc(symshen_4magless) == __self {
__self1, __self2 = V4115, let__1196
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4magless), V4115, let__1196)
return
}


}, 2)

tmp1858 := Call(__e, ns2_1set, symshen_4magless, tmp1854)


_ = tmp1858

tmp1859 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V4121 := __self1
_ = V4121
V4122 := __self2
_ = V4122
tmp1866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V4121)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V4121
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1866 {
__e.Return(True)
return
} else {
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(V4121)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := V4121
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
__e.Return(False)
return
} else {
let__1197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V4121)
__typedN1, __typedOK1 := TypedFloat64(V4122)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V4121
__typedArg1 := V4122
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})()
_ = let__1197

if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(0))
__typedN1, __typedOK1 := TypedFloat64(let__1197)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := MakeInteger(0)
__typedArg1 := let__1197
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V4121)
}
__typedArg0 := V4121
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})())
return
} else {
if PrimFunc(symshen_4integer_1test_2) == __self {
__self1, __self2 = let__1197, V4122
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4integer_1test_2), let__1197, V4122)
return
}


}


}


}, 2)

tmp1867 := Call(__e, ns2_1set, symshen_4integer_1test_2, tmp1859)


_ = tmp1867

tmp1868 := MakeNative(func(__e *ControlFlow) {
V4130 := __e.Get(1)
_ = V4130
V4131 := __e.Get(2)
_ = V4131
tmp1876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4131)
}
__typedArg0 := Nil
__typedArg1 := V4131
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1876 {
__e.Return(Nil)
return
} else {
tmp1874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4131)
}
__typedArg0 := V4131
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1874 {
tmp1869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4131)
}
__typedArg0 := V4131
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1870 := Call(__e, V4130, tmp1869)


tmp1871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4131)
}
__typedArg0 := V4131
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1872 := Call(__e, PrimFunc(symmapcan), V4130, tmp1871)


__e.TailApply(PrimFunc(symappend), tmp1870, tmp1872)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("attempt to mapcan over a non-list\n"))
}
__typedArg0 := MakeString("attempt to mapcan over a non-list\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp1877 := Call(__e, ns2_1set, symmapcan, tmp1868)


_ = tmp1877

tmp1878 := MakeNative(func(__e *ControlFlow) {
V4137 := __e.Get(1)
_ = V4137
V4138 := __e.Get(2)
_ = V4138
tmp1880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V4137, V4138)
}
__typedArg0 := V4137
__typedArg1 := V4138
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1880 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp1881 := Call(__e, ns2_1set, sym_a_a, tmp1878)


_ = tmp1881

tmp1882 := MakeNative(func(__e *ControlFlow) {
V4139 := __e.Get(1)
_ = V4139
tmp1890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V4139)
}
__typedArg0 := V4139
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp1890 {
tmp1884 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(V4139)
}
__typedArg0 := V4139
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1885 := MakeNative(func(__e *ControlFlow) {
Z4141 := __e.Get(1)
_ = Z4141
__e.Return(symshen_4this_1symbol_1is_1unbound)
return
}, 1)

tmp1886 := Call(__e, try_1catch, tmp1884, tmp1885)


let__1198 := tmp1886
_ = let__1198

tmp1888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__1198, symshen_4this_1symbol_1is_1unbound)
}
__typedArg0 := let__1198
__typedArg1 := symshen_4this_1symbol_1is_1unbound
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1887 Obj

if True == tmp1888 {
ifres1887 = False


} else {
ifres1887 = True


}

if True == ifres1887 {
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

tmp1891 := Call(__e, ns2_1set, symbound_2, tmp1882)


_ = tmp1891

tmp1892 := MakeNative(func(__e *ControlFlow) {
V4142 := __e.Get(1)
_ = V4142
tmp1898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V4142)
}
__typedArg0 := MakeString("")
__typedArg1 := V4142
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1898 {
__e.Return(Nil)
return
} else {
tmp1893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sympos) {
return PrimPos(V4142, MakeInteger(0))
}
__typedArg0 := V4142
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sympos), __typedArg0, __typedArg1)
})()

tmp1894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp1893)
}
__typedArg0 := tmp1893
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp1896 := Call(__e, PrimFunc(symshen_4string_1_6bytes), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V4142)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V4142
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1894, tmp1896)
}
__typedArg0 := tmp1894
__typedArg1 := tmp1896
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp1899 := Call(__e, ns2_1set, symshen_4string_1_6bytes, tmp1892)


_ = tmp1899

tmp1900 := MakeNative(func(__e *ControlFlow) {
V4143 := __e.Get(1)
_ = V4143
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5) {
__typedN0, __typedOK0 := TypedFloat64(V4143)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(0))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5) {
return TypedMaterializeBoolean((__typedN0 < __typedN1))
}}
__typedArg0 := V4143
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_5), __typedArg0, __typedArg1)
})() {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dmaxinferences_d)
}
__typedArg0 := symshen_4_dmaxinferences_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
} else {
tmp1902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V4143)
}
__typedArg0 := V4143
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})()

if True == tmp1902 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dmaxinferences_d, V4143)
}
__typedArg0 := symshen_4_dmaxinferences_d
__typedArg1 := V4143
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("maxinferences expects an integer value\n"))
}
__typedArg0 := MakeString("maxinferences expects an integer value\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp1905 := Call(__e, ns2_1set, symmaxinferences, tmp1900)


_ = tmp1905

tmp1906 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dinfs_d)
}
__typedArg0 := symshen_4_dinfs_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1907 := Call(__e, ns2_1set, syminferences, tmp1906)


_ = tmp1907

tmp1908 := MakeNative(func(__e *ControlFlow) {
V4144 := __e.Get(1)
_ = V4144
__e.Return(V4144)
return
}, 1)

tmp1909 := Call(__e, ns2_1set, symprotect, tmp1908)


_ = tmp1909

tmp1910 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dsterror_d)
}
__typedArg0 := sym_dsterror_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1911 := Call(__e, ns2_1set, symsterror, tmp1910)


_ = tmp1911

tmp1912 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dstoutput_d)
}
__typedArg0 := sym_dstoutput_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1913 := Call(__e, ns2_1set, symstoutput, tmp1912)


_ = tmp1913

tmp1914 := MakeNative(func(__e *ControlFlow) {
V4145 := __e.Get(1)
_ = V4145
tmp1915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(V4145)
}
__typedArg0 := V4145
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

let__1199 := tmp1915
_ = let__1199

tmp1919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(let__1199)
}
__typedArg0 := let__1199
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp1919 {
__e.Return(let__1199)
return
} else {
tmp1916 := Call(__e, PrimFunc(symshen_4app), V4145, MakeString(" to a symbol"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("cannot intern "))
__typedS1, __typedOK1 := TypedString(tmp1916)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("cannot intern ")
__typedArg1 := tmp1916
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("cannot intern "))
__typedS1, __typedOK1 := TypedString(tmp1916)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("cannot intern ")
__typedArg1 := tmp1916
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 1)

tmp1920 := Call(__e, ns2_1set, symstring_1_6symbol, tmp1914)


_ = tmp1920

tmp1921 := MakeNative(func(__e *ControlFlow) {
V4149 := __e.Get(1)
_ = V4149
tmp1925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V4149)
}
__typedArg0 := sym_7
__typedArg1 := V4149
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1925 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_doptimise_d, True)
}
__typedArg0 := symshen_4_doptimise_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
tmp1923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V4149)
}
__typedArg0 := sym_1
__typedArg1 := V4149
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1923 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_doptimise_d, False)
}
__typedArg0 := symshen_4_doptimise_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("optimise expects a + or a -.\n"))
}
__typedArg0 := MakeString("optimise expects a + or a -.\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp1926 := Call(__e, ns2_1set, symoptimise, tmp1921)


_ = tmp1926

tmp1927 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dos_d)
}
__typedArg0 := sym_dos_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1928 := Call(__e, ns2_1set, symos, tmp1927)


_ = tmp1928

tmp1929 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dlanguage_d)
}
__typedArg0 := sym_dlanguage_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1930 := Call(__e, ns2_1set, symlanguage, tmp1929)


_ = tmp1930

tmp1931 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dversion_d)
}
__typedArg0 := sym_dversion_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1932 := Call(__e, ns2_1set, symversion, tmp1931)


_ = tmp1932

tmp1933 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dport_d)
}
__typedArg0 := sym_dport_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1934 := Call(__e, ns2_1set, symport, tmp1933)


_ = tmp1934

tmp1935 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dporters_d)
}
__typedArg0 := sym_dporters_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1936 := Call(__e, ns2_1set, symporters, tmp1935)


_ = tmp1936

tmp1937 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dimplementation_d)
}
__typedArg0 := sym_dimplementation_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1938 := Call(__e, ns2_1set, symimplementation, tmp1937)


_ = tmp1938

tmp1939 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_drelease_d)
}
__typedArg0 := sym_drelease_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1940 := Call(__e, ns2_1set, symrelease, tmp1939)


_ = tmp1940

tmp1941 := MakeNative(func(__e *ControlFlow) {
V4150 := __e.Get(1)
_ = V4150
tmp1946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnull, V4150)
}
__typedArg0 := symnull
__typedArg1 := V4150
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1946 {
__e.Return(True)
return
} else {
tmp1942 := MakeNative(func(__e *ControlFlow) {
tmp1943 := Call(__e, PrimFunc(symexternal), V4150)


_ = tmp1943

__e.Return(True)
return


}, 0)

tmp1944 := MakeNative(func(__e *ControlFlow) {
Z4151 := __e.Get(1)
_ = Z4151
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1942, tmp1944)
return


}


}, 1)

tmp1947 := Call(__e, ns2_1set, sympackage_2, tmp1941)


_ = tmp1947

tmp1948 := MakeNative(func(__e *ControlFlow) {
__e.Return(symshen_4fail_b)
return
}, 0)

tmp1949 := Call(__e, ns2_1set, symfail, tmp1948)


_ = tmp1949

tmp1950 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_duserdefs_d)
}
__typedArg0 := symshen_4_duserdefs_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1951 := Call(__e, ns2_1set, symuserdefs, tmp1950)


_ = tmp1951

tmp1952 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_doptimise_d)
}
__typedArg0 := symshen_4_doptimise_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1953 := Call(__e, ns2_1set, symoptimise_2, tmp1952)


_ = tmp1953

tmp1954 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dhush_d)
}
__typedArg0 := sym_dhush_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1955 := Call(__e, ns2_1set, symhush_2, tmp1954)


_ = tmp1955

tmp1956 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dshen_1type_1theory_1enabled_2_d)
}
__typedArg0 := symshen_4_dshen_1type_1theory_1enabled_2_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1957 := Call(__e, ns2_1set, symsystem_1S_2, tmp1956)


_ = tmp1957

tmp1958 := MakeNative(func(__e *ControlFlow) {
V4154 := __e.Get(1)
_ = V4154
tmp1962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V4154)
}
__typedArg0 := sym_7
__typedArg1 := V4154
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1962 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dshen_1type_1theory_1enabled_2_d, True)
}
__typedArg0 := symshen_4_dshen_1type_1theory_1enabled_2_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
tmp1960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V4154)
}
__typedArg0 := sym_1
__typedArg1 := V4154
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1960 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dshen_1type_1theory_1enabled_2_d, False)
}
__typedArg0 := symshen_4_dshen_1type_1theory_1enabled_2_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("enable-type-theory expects a + or a -\n"))
}
__typedArg0 := MakeString("enable-type-theory expects a + or a -\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp1963 := Call(__e, ns2_1set, symenable_1type_1theory, tmp1958)


_ = tmp1963

tmp1964 := MakeNative(func(__e *ControlFlow) {
V4157 := __e.Get(1)
_ = V4157
tmp1968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V4157)
}
__typedArg0 := sym_7
__typedArg1 := V4157
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1968 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dhush_d, True)
}
__typedArg0 := sym_dhush_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
tmp1966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V4157)
}
__typedArg0 := sym_1
__typedArg1 := V4157
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1966 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dhush_d, False)
}
__typedArg0 := sym_dhush_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("hush expects a + or a -\n"))
}
__typedArg0 := MakeString("hush expects a + or a -\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp1969 := Call(__e, ns2_1set, symhush, tmp1964)


_ = tmp1969

tmp1970 := MakeNative(func(__e *ControlFlow) {
V4160 := __e.Get(1)
_ = V4160
tmp1974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V4160)
}
__typedArg0 := sym_7
__typedArg1 := V4160
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1974 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dtc_d, True)
}
__typedArg0 := symshen_4_dtc_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
tmp1972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V4160)
}
__typedArg0 := sym_1
__typedArg1 := V4160
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1972 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dtc_d, False)
}
__typedArg0 := symshen_4_dtc_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("tc expects a + or -"))
}
__typedArg0 := MakeString("tc expects a + or -")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp1975 := Call(__e, ns2_1set, symtc, tmp1970)


_ = tmp1975

tmp1976 := MakeNative(func(__e *ControlFlow) {
V4161 := __e.Get(1)
_ = V4161
tmp1977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp1978 := Call(__e, PrimFunc(symshen_4unassoc), V4161, tmp1977)


tmp1979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp1978)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp1978
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp1979

__e.Return(V4161)
return


}, 1)

tmp1980 := Call(__e, ns2_1set, symdestroy, tmp1976)


_ = tmp1980

tmp1981 := MakeNative(func(__e *ControlFlow) {
V4171 := __e.Get(1)
_ = V4171
V4172 := __e.Get(2)
_ = V4172
tmp1999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4172)
}
__typedArg0 := Nil
__typedArg1 := V4172
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1999 {
__e.Return(Nil)
return
} else {
tmp1997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4172)
}
__typedArg0 := V4172
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1988 Obj

if True == tmp1997 {
tmp1995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4172)
}
__typedArg0 := V4172
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp1995)
}
__typedArg0 := tmp1995
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1990 Obj

if True == tmp1996 {
tmp1992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4172)
}
__typedArg0 := V4172
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1992)
}
__typedArg0 := tmp1992
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V4171, tmp1993)
}
__typedArg0 := V4171
__typedArg1 := tmp1993
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1991 Obj

if True == tmp1994 {
ifres1991 = True


} else {
ifres1991 = False


}

ifres1990 = ifres1991


} else {
ifres1990 = False


}

var ifres1989 Obj

if True == ifres1990 {
ifres1989 = True


} else {
ifres1989 = False


}

ifres1988 = ifres1989


} else {
ifres1988 = False


}

if True == ifres1988 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4172)
}
__typedArg0 := V4172
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
} else {
tmp1986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4172)
}
__typedArg0 := V4172
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1986 {
tmp1982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4172)
}
__typedArg0 := V4172
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4172)
}
__typedArg0 := V4172
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1984 := Call(__e, PrimFunc(symshen_4unassoc), V4171, tmp1983)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1982, tmp1984)
}
__typedArg0 := tmp1982
__typedArg1 := tmp1984
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.unassoc"))
}
__typedArg0 := MakeString("implementation error in shen.unassoc")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 2)

tmp2000 := Call(__e, ns2_1set, symshen_4unassoc, tmp1981)


_ = tmp2000

tmp2001 := MakeNative(func(__e *ControlFlow) {
V4173 := __e.Get(1)
_ = V4173
tmp2005 := Call(__e, PrimFunc(sympackage_2), V4173)


if True == tmp2005 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dpackage_d, V4173)
}
__typedArg0 := symshen_4_dpackage_d
__typedArg1 := V4173
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
tmp2002 := Call(__e, PrimFunc(symshen_4app), V4173, MakeString(" does not exist\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("package "))
__typedS1, __typedOK1 := TypedString(tmp2002)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("package ")
__typedArg1 := tmp2002
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("package "))
__typedS1, __typedOK1 := TypedString(tmp2002)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("package ")
__typedArg1 := tmp2002
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 1)

tmp2006 := Call(__e, ns2_1set, symin_1package, tmp2001)


_ = tmp2006

tmp2007 := MakeNative(func(__e *ControlFlow) {
V4174 := __e.Get(1)
_ = V4174
V4175 := __e.Get(2)
_ = V4175
tmp2008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symopen) {
return PrimOpenStream(V4174, symout)
}
__typedArg0 := V4174
__typedArg1 := symout
return Call(__e, PrimFunc(symopen), __typedArg0, __typedArg1)
})()

let__1200 := tmp2008
_ = let__1200

tmp2011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V4175)
}
__typedArg0 := V4175
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres2009 Obj

if True == tmp2011 {
ifres2009 = V4175


} else {
tmp2010 := Call(__e, PrimFunc(symshen_4app), V4175, MakeString(""), symshen_4s)


ifres2009 = tmp2010


}

let__1201 := ifres2009
_ = let__1201

tmp2012 := Call(__e, PrimFunc(sympr), let__1201, let__1200)


let__1202 := tmp2012
_ = let__1202

tmp2013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symclose) {
return PrimCloseStream(let__1200)
}
__typedArg0 := let__1200
return Call(__e, PrimFunc(symclose), __typedArg0)
})()

let__1203 := tmp2013
_ = let__1203

__e.Return(V4175)
return


}, 2)

tmp2014 := Call(__e, ns2_1set, symwrite_1to_1file, tmp2007)


_ = tmp2014

tmp2015 := MakeNative(func(__e *ControlFlow) {
tmp2016 := Call(__e, PrimFunc(symgensym), symshen_4t)


__e.TailApply(PrimFunc(symshen_4freshterm), tmp2016)
return


}, 0)

tmp2017 := Call(__e, ns2_1set, symfresh, tmp2015)


_ = tmp2017

tmp2018 := MakeNative(func(__e *ControlFlow) {
V4180 := __e.Get(1)
_ = V4180
V4181 := __e.Get(2)
_ = V4181
tmp2019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2020 := Call(__e, PrimFunc(symput), V4180, symarity, V4181, tmp2019)


let__1204 := tmp2020
_ = let__1204

tmp2021 := Call(__e, PrimFunc(symshen_4lambda_1entry), V4180)


let__1205 := tmp2021
_ = let__1205

tmp2026 := Call(__e, PrimFunc(symempty_2), let__1205)


var ifres2022 Obj

if True == tmp2026 {
tmp2023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2024 := Call(__e, PrimFunc(symunput), V4180, symshen_4lambda_1form, tmp2023)


ifres2022 = tmp2024


} else {
tmp2025 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), let__1205)


ifres2022 = tmp2025


}

let__1206 := ifres2022
_ = let__1206

__e.Return(V4180)
return


}, 2)

tmp2027 := Call(__e, ns2_1set, symupdate_1lambda_1table, tmp2018)


_ = tmp2027

tmp2028 := MakeNative(func(__e *ControlFlow) {
V4187 := __e.Get(1)
_ = V4187
__e.Return(False)
return
}, 1)

tmp2029 := Call(__e, ns2_1set, symshen_4precompiled_1function_2, tmp2028)


_ = tmp2029

tmp2030 := MakeNative(func(__e *ControlFlow) {
V4188 := __e.Get(1)
_ = V4188
V4189 := __e.Get(2)
_ = V4189
V4190 := __e.Get(3)
_ = V4190
V4191 := __e.Get(4)
_ = V4191
if True == V4191 {
tmp2031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2032 := Call(__e, PrimFunc(symput), V4188, symarity, V4189, tmp2031)


let__1207 := tmp2032
_ = let__1207

tmp2033 := Call(__e, PrimFunc(symthaw), V4190)


let__1208 := tmp2033
_ = let__1208

tmp2039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V4189, MakeInteger(0))
}
__typedArg0 := V4189
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2034 Obj

if True == tmp2039 {
tmp2035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2036 := Call(__e, PrimFunc(symunput), V4188, symshen_4lambda_1form, tmp2035)


ifres2034 = tmp2036


} else {
tmp2037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4188, let__1208)
}
__typedArg0 := V4188
__typedArg1 := let__1208
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2038 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2037)


ifres2034 = tmp2038


}

let__1209 := ifres2034
_ = let__1209

__e.Return(V4188)
return


} else {
__e.TailApply(PrimFunc(symupdate_1lambda_1table), V4188, V4189)
return
}
}, 4)

tmp2041 := Call(__e, ns2_1set, symshen_4update_1lambda_1table_cprecompiled, tmp2030)


_ = tmp2041

tmp2042 := MakeNative(func(__e *ControlFlow) {
V4197 := __e.Get(1)
_ = V4197
V4198 := __e.Get(2)
_ = V4198
tmp2066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V4198)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V4198
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2066 {
tmp2043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspecial_d)
}
__typedArg0 := symshen_4_dspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2044 := Call(__e, PrimFunc(symremove), V4197, tmp2043)


tmp2045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspecial_d, tmp2044)
}
__typedArg0 := symshen_4_dspecial_d
__typedArg1 := tmp2044
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2045

tmp2046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dextraspecial_d)
}
__typedArg0 := symshen_4_dextraspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2047 := Call(__e, PrimFunc(symremove), V4197, tmp2046)


tmp2048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dextraspecial_d, tmp2047)
}
__typedArg0 := symshen_4_dextraspecial_d
__typedArg1 := tmp2047
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2048

__e.Return(V4197)
return


} else {
tmp2064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(1), V4198)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := V4198
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2064 {
tmp2049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspecial_d)
}
__typedArg0 := symshen_4_dspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2050 := Call(__e, PrimFunc(symadjoin), V4197, tmp2049)


tmp2051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspecial_d, tmp2050)
}
__typedArg0 := symshen_4_dspecial_d
__typedArg1 := tmp2050
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2051

tmp2052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dextraspecial_d)
}
__typedArg0 := symshen_4_dextraspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2053 := Call(__e, PrimFunc(symremove), V4197, tmp2052)


tmp2054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dextraspecial_d, tmp2053)
}
__typedArg0 := symshen_4_dextraspecial_d
__typedArg1 := tmp2053
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2054

__e.Return(V4197)
return


} else {
tmp2062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(2), V4198)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := V4198
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2062 {
tmp2055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspecial_d)
}
__typedArg0 := symshen_4_dspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2056 := Call(__e, PrimFunc(symremove), V4197, tmp2055)


tmp2057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspecial_d, tmp2056)
}
__typedArg0 := symshen_4_dspecial_d
__typedArg1 := tmp2056
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2057

tmp2058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dextraspecial_d)
}
__typedArg0 := symshen_4_dextraspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2059 := Call(__e, PrimFunc(symadjoin), V4197, tmp2058)


tmp2060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dextraspecial_d, tmp2059)
}
__typedArg0 := symshen_4_dextraspecial_d
__typedArg1 := tmp2059
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2060

__e.Return(V4197)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("specialise requires values of 0, 1 or 2\n"))
}
__typedArg0 := MakeString("specialise requires values of 0, 1 or 2\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 2)

tmp2067 := Call(__e, ns2_1set, symspecialise, tmp2042)


_ = tmp2067

tmp2068 := MakeNative(func(__e *ControlFlow) {
V4199 := __e.Get(1)
_ = V4199
tmp2069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dabsolute_d)
}
__typedArg0 := sym_dabsolute_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4199, tmp2069)
}
__typedArg0 := V4199
__typedArg1 := tmp2069
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dabsolute_d, tmp2070)
}
__typedArg0 := sym_dabsolute_d
__typedArg1 := tmp2070
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp2071 := Call(__e, ns2_1set, symabsolute, tmp2068)


_ = tmp2071

tmp2072 := MakeNative(func(__e *ControlFlow) {
V4200 := __e.Get(1)
_ = V4200
tmp2073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dabsolute_d)
}
__typedArg0 := sym_dabsolute_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2074 := Call(__e, PrimFunc(symremove), V4200, tmp2073)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dabsolute_d, tmp2074)
}
__typedArg0 := sym_dabsolute_d
__typedArg1 := tmp2074
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 1)

__e.TailApply(ns2_1set, symunabsolute, tmp2072)
return




}, 0)

