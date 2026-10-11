package main

import . "github.com/pyrex41/shen-go/kl"

var SysMain = MakeNative(func(__e *ControlFlow) {
tmp1207 := MakeNative(func(__e *ControlFlow) {
V3830 := __e.Get(1)
_ = V3830
__e.TailApply(V3830)
return
}, 1)

tmp1208 := Call(__e, ns2_1set, symthaw, tmp1207)


_ = tmp1208

tmp1209 := MakeNative(func(__e *ControlFlow) {
V3831 := __e.Get(1)
_ = V3831
tmp1210 := Call(__e, PrimFunc(symmacroexpand), V3831)


tmp1211 := Call(__e, PrimFunc(symshen_4find_1types), V3831)


tmp1212 := Call(__e, PrimFunc(symshen_4process_1applications), tmp1210, tmp1211)


tmp1213 := Call(__e, PrimFunc(symshen_4shen_1_6kl), tmp1212)


__e.TailApply(PrimFunc(symeval_1kl), tmp1213)
return


}, 1)

tmp1214 := Call(__e, ns2_1set, symeval, tmp1209)


_ = tmp1214

tmp1215 := MakeNative(func(__e *ControlFlow) {
V3832 := __e.Get(1)
_ = V3832
tmp1222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnull, V3832)
}
__typedArg0 := symnull
__typedArg1 := V3832
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1222 {
__e.Return(Nil)
return
} else {
tmp1216 := MakeNative(func(__e *ControlFlow) {
tmp1217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V3832, symshen_4external_1symbols, tmp1217)
return


}, 0)

tmp1218 := MakeNative(func(__e *ControlFlow) {
Z3833 := __e.Get(1)
_ = Z3833
tmp1219 := Call(__e, PrimFunc(symshen_4app), V3832, MakeString(" does not exist.\n;"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("package "))
__typedS1, __typedOK1 := TypedString(tmp1219)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("package ")
__typedArg1 := tmp1219
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("package "))
__typedS1, __typedOK1 := TypedString(tmp1219)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("package ")
__typedArg1 := tmp1219
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

__e.TailApply(try_1catch, tmp1216, tmp1218)
return


}


}, 1)

tmp1223 := Call(__e, ns2_1set, symexternal, tmp1215)


_ = tmp1223

tmp1224 := MakeNative(func(__e *ControlFlow) {
V3834 := __e.Get(1)
_ = V3834
tmp1231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnull, V3834)
}
__typedArg0 := symnull
__typedArg1 := V3834
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1231 {
__e.Return(Nil)
return
} else {
tmp1225 := MakeNative(func(__e *ControlFlow) {
tmp1226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V3834, symshen_4internal_1symbols, tmp1226)
return


}, 0)

tmp1227 := MakeNative(func(__e *ControlFlow) {
Z3835 := __e.Get(1)
_ = Z3835
tmp1228 := Call(__e, PrimFunc(symshen_4app), V3834, MakeString(" does not exist.\n;"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("package "))
__typedS1, __typedOK1 := TypedString(tmp1228)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("package ")
__typedArg1 := tmp1228
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("package "))
__typedS1, __typedOK1 := TypedString(tmp1228)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("package ")
__typedArg1 := tmp1228
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

__e.TailApply(try_1catch, tmp1225, tmp1227)
return


}


}, 1)

tmp1232 := Call(__e, ns2_1set, syminternal, tmp1224)


_ = tmp1232

tmp1233 := MakeNative(func(__e *ControlFlow) {
V3836 := __e.Get(1)
_ = V3836
V3837 := __e.Get(2)
_ = V3837
tmp1235 := Call(__e, V3836, V3837)


if True == tmp1235 {
__e.TailApply(PrimFunc(symfail))
return
} else {
__e.Return(V3837)
return
}


}, 2)

tmp1236 := Call(__e, ns2_1set, symfail_1if, tmp1233)


_ = tmp1236

tmp1237 := MakeNative(func(__e *ControlFlow) {
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

tmp1238 := Call(__e, ns2_1set, sym_8s, tmp1237)


_ = tmp1238

tmp1239 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtc_d)
}
__typedArg0 := symshen_4_dtc_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1240 := Call(__e, ns2_1set, symtc_2, tmp1239)


_ = tmp1240

tmp1241 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_doccurs_d)
}
__typedArg0 := symshen_4_doccurs_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1242 := Call(__e, ns2_1set, symoccurs_2, tmp1241)


_ = tmp1242

tmp1243 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dfactorise_2_d)
}
__typedArg0 := symshen_4_dfactorise_2_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1244 := Call(__e, ns2_1set, symfactorise_2, tmp1243)


_ = tmp1244

tmp1245 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtracking_d)
}
__typedArg0 := symshen_4_dtracking_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1246 := Call(__e, ns2_1set, symtracked, tmp1245)


_ = tmp1246

tmp1247 := MakeNative(func(__e *ControlFlow) {
V3840 := __e.Get(1)
_ = V3840
tmp1248 := MakeNative(func(__e *ControlFlow) {
tmp1249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V3840, symshen_4source, tmp1249)
return


}, 0)

tmp1250 := MakeNative(func(__e *ControlFlow) {
Z3841 := __e.Get(1)
_ = Z3841
tmp1251 := Call(__e, PrimFunc(symshen_4app), V3840, MakeString(" not found.\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp1251)
}
__typedArg0 := tmp1251
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

__e.TailApply(try_1catch, tmp1248, tmp1250)
return


}, 1)

tmp1252 := Call(__e, ns2_1set, symps, tmp1247)


_ = tmp1252

tmp1253 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dstinput_d)
}
__typedArg0 := sym_dstinput_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1254 := Call(__e, ns2_1set, symstinput, tmp1253)


_ = tmp1254

tmp1255 := MakeNative(func(__e *ControlFlow) {
V3842 := __e.Get(1)
_ = V3842
tmp1257 := (func() Obj {
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

let__1164 := tmp1257
_ = let__1164

tmp1258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__1164, MakeInteger(0), V3842)
}
__typedArg0 := let__1164
__typedArg1 := MakeInteger(0)
__typedArg2 := V3842
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__1165 := tmp1258
_ = let__1165

tmp1262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3842, MakeInteger(0))
}
__typedArg0 := V3842
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1259 Obj

if True == tmp1262 {
ifres1259 = let__1165


} else {
tmp1260 := Call(__e, PrimFunc(symfail))


tmp1261 := Call(__e, PrimFunc(symshen_4fillvector), let__1165, MakeInteger(1), V3842, tmp1260)


ifres1259 = tmp1261


}

let__1166 := ifres1259
_ = let__1166

__e.Return(let__1166)
return


}, 1)

tmp1263 := Call(__e, ns2_1set, symvector, tmp1255)


_ = tmp1263

tmp1264 := MakeNative(func(__e *ControlFlow) {
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
tmp1268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3848, V3849)
}
__typedArg0 := V3848
__typedArg1 := V3849
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1268 {
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
tmp1265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V3847, V3848, V3850)
}
__typedArg0 := V3847
__typedArg1 := V3848
__typedArg2 := V3850
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

if PrimFunc(symshen_4fillvector) == __self {
__self1, __self2, __self3, __self4 = tmp1265, (func() Obj {
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
__e.TailApply(PrimFunc(symshen_4fillvector), tmp1265, (func() Obj {
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

tmp1269 := Call(__e, ns2_1set, symshen_4fillvector, tmp1264)


_ = tmp1269

tmp1270 := MakeNative(func(__e *ControlFlow) {
V3851 := __e.Get(1)
_ = V3851
tmp1280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V3851)
}
__typedArg0 := V3851
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp1280 {
tmp1272 := MakeNative(func(__e *ControlFlow) {
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

tmp1273 := MakeNative(func(__e *ControlFlow) {
Z3853 := __e.Get(1)
_ = Z3853
__e.Return(MakeInteger(-1))
return
}, 1)

tmp1274 := Call(__e, try_1catch, tmp1272, tmp1273)


let__1167 := tmp1274
_ = let__1167

tmp1278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(let__1167)
}
__typedArg0 := let__1167
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

var ifres1275 Obj

if True == tmp1278 {
tmp1277 := (func() Obj {
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

var ifres1276 Obj

if True == tmp1277 {
ifres1276 = True


} else {
ifres1276 = False


}

ifres1275 = ifres1276


} else {
ifres1275 = False


}

if True == ifres1275 {
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

tmp1281 := Call(__e, ns2_1set, symvector_2, tmp1270)


_ = tmp1281

tmp1282 := MakeNative(func(__e *ControlFlow) {
V3854 := __e.Get(1)
_ = V3854
V3855 := __e.Get(2)
_ = V3855
V3856 := __e.Get(3)
_ = V3856
tmp1284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3855, MakeInteger(0))
}
__typedArg0 := V3855
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1284 {
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

tmp1285 := Call(__e, ns2_1set, symvector_1_6, tmp1282)


_ = tmp1285

tmp1286 := MakeNative(func(__e *ControlFlow) {
V3857 := __e.Get(1)
_ = V3857
V3858 := __e.Get(2)
_ = V3858
tmp1292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3858, MakeInteger(0))
}
__typedArg0 := V3858
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1292 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("cannot access 0th element of a vector\n"))
}
__typedArg0 := MakeString("cannot access 0th element of a vector\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp1287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V3857, V3858)
}
__typedArg0 := V3857
__typedArg1 := V3858
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

let__1168 := tmp1287
_ = let__1168

tmp1289 := Call(__e, PrimFunc(symfail))


tmp1290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__1168, tmp1289)
}
__typedArg0 := let__1168
__typedArg1 := tmp1289
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1290 {
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

tmp1293 := Call(__e, ns2_1set, sym_5_1vector, tmp1286)


_ = tmp1293

tmp1294 := MakeNative(func(__e *ControlFlow) {
V3860 := __e.Get(1)
_ = V3860
tmp1298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V3860)
}
__typedArg0 := V3860
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})()

if True == tmp1298 {
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

tmp1299 := Call(__e, ns2_1set, symshen_4posint_2, tmp1294)


_ = tmp1299

tmp1300 := MakeNative(func(__e *ControlFlow) {
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

tmp1301 := Call(__e, ns2_1set, symlimit, tmp1300)


_ = tmp1301

tmp1302 := MakeNative(func(__e *ControlFlow) {
V3862 := __e.Get(1)
_ = V3862
tmp1332 := Call(__e, PrimFunc(symboolean_2), V3862)


var ifres1317 Obj

if True == tmp1332 {
ifres1317 = True


} else {
tmp1331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(V3862)
}
__typedArg0 := V3862
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

var ifres1319 Obj

if True == tmp1331 {
ifres1319 = True


} else {
tmp1330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V3862)
}
__typedArg0 := V3862
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres1321 Obj

if True == tmp1330 {
ifres1321 = True


} else {
tmp1329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3862)
}
__typedArg0 := V3862
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1323 Obj

if True == tmp1329 {
ifres1323 = True


} else {
tmp1328 := Call(__e, PrimFunc(symempty_2), V3862)


var ifres1325 Obj

if True == tmp1328 {
ifres1325 = True


} else {
tmp1327 := Call(__e, PrimFunc(symvector_2), V3862)


var ifres1326 Obj

if True == tmp1327 {
ifres1326 = True


} else {
ifres1326 = False


}

ifres1325 = ifres1326


}

var ifres1324 Obj

if True == ifres1325 {
ifres1324 = True


} else {
ifres1324 = False


}

ifres1323 = ifres1324


}

var ifres1322 Obj

if True == ifres1323 {
ifres1322 = True


} else {
ifres1322 = False


}

ifres1321 = ifres1322


}

var ifres1320 Obj

if True == ifres1321 {
ifres1320 = True


} else {
ifres1320 = False


}

ifres1319 = ifres1320


}

var ifres1318 Obj

if True == ifres1319 {
ifres1318 = True


} else {
ifres1318 = False


}

ifres1317 = ifres1318


}

if True == ifres1317 {
__e.Return(False)
return
} else {
tmp1307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp1308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp1309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(","))
}
__typedArg0 := MakeString(",")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp1310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1309, Nil)
}
__typedArg0 := tmp1309
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1308, tmp1310)
}
__typedArg0 := tmp1308
__typedArg1 := tmp1310
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1307, tmp1311)
}
__typedArg0 := tmp1307
__typedArg1 := tmp1311
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, tmp1312)
}
__typedArg0 := sym_j
__typedArg1 := tmp1312
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, tmp1313)
}
__typedArg0 := sym_i
__typedArg1 := tmp1313
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1315 := Call(__e, PrimFunc(symelement_2), V3862, tmp1314)


if True == tmp1315 {
__e.Return(True)
return
} else {
tmp1303 := MakeNative(func(__e *ControlFlow) {
tmp1304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3862)
}
__typedArg0 := V3862
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

let__1169 := tmp1304
_ = let__1169

__e.TailApply(PrimFunc(symshen_4analyse_1symbol_2), let__1169)
return


}, 0)

tmp1305 := MakeNative(func(__e *ControlFlow) {
Z3864 := __e.Get(1)
_ = Z3864
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1303, tmp1305)
return


}


}


}, 1)

tmp1333 := Call(__e, ns2_1set, symsymbol_2, tmp1302)


_ = tmp1333

tmp1334 := MakeNative(func(__e *ControlFlow) {
V3867 := __e.Get(1)
_ = V3867
tmp1343 := Call(__e, PrimFunc(symshen_4_7string_2), V3867)


if True == tmp1343 {
tmp1339 := Call(__e, PrimFunc(symhdstr), V3867)


tmp1340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp1339)
}
__typedArg0 := tmp1339
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp1341 := Call(__e, PrimFunc(symshen_4alpha_2), tmp1340)


if True == tmp1341 {
tmp1337 := Call(__e, PrimFunc(symshen_4alphanums_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3867)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3867
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


if True == tmp1337 {
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

tmp1344 := Call(__e, ns2_1set, symshen_4analyse_1symbol_2, tmp1334)


_ = tmp1344

tmp1345 := MakeNative(func(__e *ControlFlow) {
V3870 := __e.Get(1)
_ = V3870
tmp1359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V3870)
}
__typedArg0 := MakeString("")
__typedArg1 := V3870
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1359 {
__e.Return(True)
return
} else {
tmp1357 := Call(__e, PrimFunc(symshen_4_7string_2), V3870)


if True == tmp1357 {
tmp1346 := Call(__e, PrimFunc(symhdstr), V3870)


tmp1347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp1346)
}
__typedArg0 := tmp1346
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

let__1170 := tmp1347
_ = let__1170

tmp1355 := Call(__e, PrimFunc(symshen_4alpha_2), let__1170)


var ifres1352 Obj

if True == tmp1355 {
ifres1352 = True


} else {
tmp1354 := Call(__e, PrimFunc(symshen_4digit_2), let__1170)


var ifres1353 Obj

if True == tmp1354 {
ifres1353 = True


} else {
ifres1353 = False


}

ifres1352 = ifres1353


}

if True == ifres1352 {
tmp1350 := Call(__e, PrimFunc(symshen_4alphanums_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3870)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3870
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


if True == tmp1350 {
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

tmp1360 := Call(__e, ns2_1set, symshen_4alphanums_2, tmp1345)


_ = tmp1360

tmp1361 := MakeNative(func(__e *ControlFlow) {
V3872 := __e.Get(1)
_ = V3872
tmp1372 := Call(__e, PrimFunc(symboolean_2), V3872)


var ifres1366 Obj

if True == tmp1372 {
ifres1366 = True


} else {
tmp1371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(V3872)
}
__typedArg0 := V3872
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

var ifres1368 Obj

if True == tmp1371 {
ifres1368 = True


} else {
tmp1370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V3872)
}
__typedArg0 := V3872
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres1369 Obj

if True == tmp1370 {
ifres1369 = True


} else {
ifres1369 = False


}

ifres1368 = ifres1369


}

var ifres1367 Obj

if True == ifres1368 {
ifres1367 = True


} else {
ifres1367 = False


}

ifres1366 = ifres1367


}

if True == ifres1366 {
__e.Return(False)
return
} else {
tmp1362 := MakeNative(func(__e *ControlFlow) {
tmp1363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3872)
}
__typedArg0 := V3872
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

let__1171 := tmp1363
_ = let__1171

__e.TailApply(PrimFunc(symshen_4analyse_1variable_2), let__1171)
return


}, 0)

tmp1364 := MakeNative(func(__e *ControlFlow) {
Z3874 := __e.Get(1)
_ = Z3874
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1362, tmp1364)
return


}


}, 1)

tmp1373 := Call(__e, ns2_1set, symvariable_2, tmp1361)


_ = tmp1373

tmp1374 := MakeNative(func(__e *ControlFlow) {
V3877 := __e.Get(1)
_ = V3877
tmp1383 := Call(__e, PrimFunc(symshen_4_7string_2), V3877)


if True == tmp1383 {
tmp1379 := Call(__e, PrimFunc(symhdstr), V3877)


tmp1380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp1379)
}
__typedArg0 := tmp1379
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp1381 := Call(__e, PrimFunc(symshen_4uppercase_2), tmp1380)


if True == tmp1381 {
tmp1377 := Call(__e, PrimFunc(symshen_4alphanums_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3877)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3877
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


if True == tmp1377 {
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

tmp1384 := Call(__e, ns2_1set, symshen_4analyse_1variable_2, tmp1374)


_ = tmp1384

tmp1385 := MakeNative(func(__e *ControlFlow) {
V3878 := __e.Get(1)
_ = V3878
tmp1386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dgensym_d)
}
__typedArg0 := symshen_4_dgensym_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp1388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dgensym_d, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp1386)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp1386
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := symshen_4_dgensym_d
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp1386)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp1386
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symconcat), V3878, tmp1388)
return


}, 1)

tmp1389 := Call(__e, ns2_1set, symgensym, tmp1385)


_ = tmp1389

tmp1390 := MakeNative(func(__e *ControlFlow) {
V3879 := __e.Get(1)
_ = V3879
V3880 := __e.Get(2)
_ = V3880
tmp1391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3879)
}
__typedArg0 := V3879
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp1392 := (func() Obj {
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
__typedS0, __typedOK0 := TypedString(tmp1391)
__typedS1, __typedOK1 := TypedString(tmp1392)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp1391
__typedArg1 := tmp1392
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp1391)
__typedS1, __typedOK1 := TypedString(tmp1392)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp1391
__typedArg1 := tmp1392
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symintern), __typedArg0)
})())
return


}, 2)

tmp1394 := Call(__e, ns2_1set, symconcat, tmp1390)


_ = tmp1394

tmp1395 := MakeNative(func(__e *ControlFlow) {
V3881 := __e.Get(1)
_ = V3881
V3882 := __e.Get(2)
_ = V3882
tmp1396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(MakeInteger(3))
}
__typedArg0 := MakeInteger(3)
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

let__1172 := tmp1396
_ = let__1172

tmp1397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__1172, MakeInteger(0), symshen_4tuple)
}
__typedArg0 := let__1172
__typedArg1 := MakeInteger(0)
__typedArg2 := symshen_4tuple
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__1173 := tmp1397
_ = let__1173

tmp1398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__1172, MakeInteger(1), V3881)
}
__typedArg0 := let__1172
__typedArg1 := MakeInteger(1)
__typedArg2 := V3881
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__1174 := tmp1398
_ = let__1174

tmp1399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__1172, MakeInteger(2), V3882)
}
__typedArg0 := let__1172
__typedArg1 := MakeInteger(2)
__typedArg2 := V3882
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__1175 := tmp1399
_ = let__1175

__e.Return(let__1172)
return


}, 2)

tmp1400 := Call(__e, ns2_1set, sym_8p, tmp1395)


_ = tmp1400

tmp1401 := MakeNative(func(__e *ControlFlow) {
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

tmp1402 := Call(__e, ns2_1set, symfst, tmp1401)


_ = tmp1402

tmp1403 := MakeNative(func(__e *ControlFlow) {
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

tmp1404 := Call(__e, ns2_1set, symsnd, tmp1403)


_ = tmp1404

tmp1405 := MakeNative(func(__e *ControlFlow) {
V3889 := __e.Get(1)
_ = V3889
tmp1412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V3889)
}
__typedArg0 := V3889
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp1412 {
tmp1407 := MakeNative(func(__e *ControlFlow) {
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

tmp1408 := MakeNative(func(__e *ControlFlow) {
Z3890 := __e.Get(1)
_ = Z3890
__e.Return(symshen_4not_1tuple)
return
}, 1)

tmp1409 := Call(__e, try_1catch, tmp1407, tmp1408)


tmp1410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4tuple, tmp1409)
}
__typedArg0 := symshen_4tuple
__typedArg1 := tmp1409
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1410 {
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

tmp1413 := Call(__e, ns2_1set, symtuple_2, tmp1405)


_ = tmp1413

tmp1414 := MakeNative(func(__e *ControlFlow) {
V3895 := __e.Get(1)
_ = V3895
V3896 := __e.Get(2)
_ = V3896
tmp1421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3895)
}
__typedArg0 := Nil
__typedArg1 := V3895
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1421 {
__e.Return(V3896)
return
} else {
tmp1419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3895)
}
__typedArg0 := V3895
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1419 {
tmp1415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3895)
}
__typedArg0 := V3895
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3895)
}
__typedArg0 := V3895
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1417 := Call(__e, PrimFunc(symappend), tmp1416, V3896)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1415, tmp1417)
}
__typedArg0 := tmp1415
__typedArg1 := tmp1417
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

tmp1422 := Call(__e, ns2_1set, symappend, tmp1414)


_ = tmp1422

tmp1423 := MakeNative(func(__e *ControlFlow) {
V3897 := __e.Get(1)
_ = V3897
V3898 := __e.Get(2)
_ = V3898
tmp1424 := Call(__e, PrimFunc(symlimit), V3898)


let__1176 := tmp1424
_ = let__1176

tmp1426 := Call(__e, PrimFunc(symvector), (func() Obj {
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


let__1177 := tmp1426
_ = let__1177

tmp1427 := Call(__e, PrimFunc(symvector_1_6), let__1177, MakeInteger(1), V3897)


let__1178 := tmp1427
_ = let__1178

tmp1429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__1176, MakeInteger(0))
}
__typedArg0 := let__1176
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1429 {
__e.Return(let__1178)
return
} else {
__e.TailApply(PrimFunc(symshen_4_8v_1help), V3898, MakeInteger(1), let__1176, let__1178)
return
}


}, 2)

tmp1430 := Call(__e, ns2_1set, sym_8v, tmp1423)


_ = tmp1430

tmp1431 := MakeNative(func(__e *ControlFlow) {
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
tmp1437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3904, V3905)
}
__typedArg0 := V3904
__typedArg1 := V3905
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1437 {
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
tmp1433 := (func() Obj {
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

tmp1435 := Call(__e, PrimFunc(symshen_4copyfromvector), V3903, V3906, V3904, (func() Obj {
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
__self1, __self2, __self3, __self4 = V3903, tmp1433, V3905, tmp1435
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4_8v_1help), V3903, tmp1433, V3905, tmp1435)
return


}


}, 4)

tmp1438 := Call(__e, ns2_1set, symshen_4_8v_1help, tmp1431)


_ = tmp1438

tmp1439 := MakeNative(func(__e *ControlFlow) {
V3907 := __e.Get(1)
_ = V3907
V3908 := __e.Get(2)
_ = V3908
V3909 := __e.Get(3)
_ = V3909
V3910 := __e.Get(4)
_ = V3910
tmp1440 := MakeNative(func(__e *ControlFlow) {
tmp1441 := Call(__e, PrimFunc(sym_5_1vector), V3907, V3909)


__e.TailApply(PrimFunc(symvector_1_6), V3908, V3910, tmp1441)
return


}, 0)

tmp1442 := MakeNative(func(__e *ControlFlow) {
Z3911 := __e.Get(1)
_ = Z3911
__e.Return(V3908)
return
}, 1)

__e.TailApply(try_1catch, tmp1440, tmp1442)
return


}, 4)

tmp1443 := Call(__e, ns2_1set, symshen_4copyfromvector, tmp1439)


_ = tmp1443

tmp1444 := MakeNative(func(__e *ControlFlow) {
V3912 := __e.Get(1)
_ = V3912
tmp1445 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(sym_5_1vector), V3912, MakeInteger(1))
return
}, 0)

tmp1446 := MakeNative(func(__e *ControlFlow) {
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

__e.TailApply(try_1catch, tmp1445, tmp1446)
return


}, 1)

tmp1447 := Call(__e, ns2_1set, symhdv, tmp1444)


_ = tmp1447

tmp1448 := MakeNative(func(__e *ControlFlow) {
V3914 := __e.Get(1)
_ = V3914
tmp1449 := Call(__e, PrimFunc(symlimit), V3914)


let__1179 := tmp1449
_ = let__1179

tmp1457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__1179, MakeInteger(0))
}
__typedArg0 := let__1179
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1457 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("cannot take the tail of the empty vector\n"))
}
__typedArg0 := MakeString("cannot take the tail of the empty vector\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp1455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__1179, MakeInteger(1))
}
__typedArg0 := let__1179
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1455 {
__e.TailApply(PrimFunc(symvector), MakeInteger(0))
return
} else {
tmp1451 := Call(__e, PrimFunc(symvector), (func() Obj {
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


let__1180 := tmp1451
_ = let__1180

tmp1453 := Call(__e, PrimFunc(symvector), (func() Obj {
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


__e.TailApply(PrimFunc(symshen_4tlv_1help), V3914, MakeInteger(2), let__1179, tmp1453)
return


}


}


}, 1)

tmp1458 := Call(__e, ns2_1set, symtlv, tmp1448)


_ = tmp1458

tmp1459 := MakeNative(func(__e *ControlFlow) {
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
tmp1465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3919, V3920)
}
__typedArg0 := V3919
__typedArg1 := V3920
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1465 {
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
tmp1461 := (func() Obj {
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

tmp1463 := Call(__e, PrimFunc(symshen_4copyfromvector), V3918, V3921, V3919, (func() Obj {
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
__self1, __self2, __self3, __self4 = V3918, tmp1461, V3920, tmp1463
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4tlv_1help), V3918, tmp1461, V3920, tmp1463)
return


}


}, 4)

tmp1466 := Call(__e, ns2_1set, symshen_4tlv_1help, tmp1459)


_ = tmp1466

tmp1467 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V3933 := __self1
_ = V3933
V3934 := __self2
_ = V3934
tmp1483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3934)
}
__typedArg0 := Nil
__typedArg1 := V3934
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1483 {
__e.Return(Nil)
return
} else {
tmp1481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3934)
}
__typedArg0 := V3934
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1472 Obj

if True == tmp1481 {
tmp1479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3934)
}
__typedArg0 := V3934
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp1479)
}
__typedArg0 := tmp1479
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1474 Obj

if True == tmp1480 {
tmp1476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3934)
}
__typedArg0 := V3934
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1476)
}
__typedArg0 := tmp1476
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3933, tmp1477)
}
__typedArg0 := V3933
__typedArg1 := tmp1477
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1475 Obj

if True == tmp1478 {
ifres1475 = True


} else {
ifres1475 = False


}

ifres1474 = ifres1475


} else {
ifres1474 = False


}

var ifres1473 Obj

if True == ifres1474 {
ifres1473 = True


} else {
ifres1473 = False


}

ifres1472 = ifres1473


} else {
ifres1472 = False


}

if True == ifres1472 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3934)
}
__typedArg0 := V3934
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp1470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3934)
}
__typedArg0 := V3934
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1470 {
tmp1468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3934)
}
__typedArg0 := V3934
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symassoc) == __self {
__self1, __self2 = V3933, tmp1468
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symassoc), V3933, tmp1468)
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

tmp1484 := Call(__e, ns2_1set, symassoc, tmp1467)


_ = tmp1484

tmp1485 := MakeNative(func(__e *ControlFlow) {
V3938 := __e.Get(1)
_ = V3938
V3939 := __e.Get(2)
_ = V3939
V3940 := __e.Get(3)
_ = V3940
tmp1508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3940)
}
__typedArg0 := Nil
__typedArg1 := V3940
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1508 {
tmp1486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3938, V3939)
}
__typedArg0 := V3938
__typedArg1 := V3939
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1486, Nil)
}
__typedArg0 := tmp1486
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp1506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3940)
}
__typedArg0 := V3940
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1497 Obj

if True == tmp1506 {
tmp1504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3940)
}
__typedArg0 := V3940
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp1504)
}
__typedArg0 := tmp1504
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1499 Obj

if True == tmp1505 {
tmp1501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3940)
}
__typedArg0 := V3940
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1501)
}
__typedArg0 := tmp1501
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3938, tmp1502)
}
__typedArg0 := V3938
__typedArg1 := tmp1502
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1500 Obj

if True == tmp1503 {
ifres1500 = True


} else {
ifres1500 = False


}

ifres1499 = ifres1500


} else {
ifres1499 = False


}

var ifres1498 Obj

if True == ifres1499 {
ifres1498 = True


} else {
ifres1498 = False


}

ifres1497 = ifres1498


} else {
ifres1497 = False


}

if True == ifres1497 {
tmp1487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3940)
}
__typedArg0 := V3940
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1487)
}
__typedArg0 := tmp1487
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1488, V3939)
}
__typedArg0 := tmp1488
__typedArg1 := V3939
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3940)
}
__typedArg0 := V3940
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1489, tmp1490)
}
__typedArg0 := tmp1489
__typedArg1 := tmp1490
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp1495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3940)
}
__typedArg0 := V3940
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1495 {
tmp1491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3940)
}
__typedArg0 := V3940
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3940)
}
__typedArg0 := V3940
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1493 := Call(__e, PrimFunc(symshen_4assoc_1set), V3938, V3939, tmp1492)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1491, tmp1493)
}
__typedArg0 := tmp1491
__typedArg1 := tmp1493
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

tmp1509 := Call(__e, ns2_1set, symshen_4assoc_1set, tmp1485)


_ = tmp1509

tmp1510 := MakeNative(func(__e *ControlFlow) {
V3944 := __e.Get(1)
_ = V3944
V3945 := __e.Get(2)
_ = V3945
tmp1528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3945)
}
__typedArg0 := Nil
__typedArg1 := V3945
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1528 {
__e.Return(Nil)
return
} else {
tmp1526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3945)
}
__typedArg0 := V3945
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1517 Obj

if True == tmp1526 {
tmp1524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3945)
}
__typedArg0 := V3945
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp1524)
}
__typedArg0 := tmp1524
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1519 Obj

if True == tmp1525 {
tmp1521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3945)
}
__typedArg0 := V3945
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1521)
}
__typedArg0 := tmp1521
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3944, tmp1522)
}
__typedArg0 := V3944
__typedArg1 := tmp1522
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1520 Obj

if True == tmp1523 {
ifres1520 = True


} else {
ifres1520 = False


}

ifres1519 = ifres1520


} else {
ifres1519 = False


}

var ifres1518 Obj

if True == ifres1519 {
ifres1518 = True


} else {
ifres1518 = False


}

ifres1517 = ifres1518


} else {
ifres1517 = False


}

if True == ifres1517 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3945)
}
__typedArg0 := V3945
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
} else {
tmp1515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3945)
}
__typedArg0 := V3945
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1515 {
tmp1511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3945)
}
__typedArg0 := V3945
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3945)
}
__typedArg0 := V3945
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1513 := Call(__e, PrimFunc(symshen_4assoc_1rm), V3944, tmp1512)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1511, tmp1513)
}
__typedArg0 := tmp1511
__typedArg1 := tmp1513
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

tmp1529 := Call(__e, ns2_1set, symshen_4assoc_1rm, tmp1510)


_ = tmp1529

tmp1530 := MakeNative(func(__e *ControlFlow) {
V3948 := __e.Get(1)
_ = V3948
tmp1534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, V3948)
}
__typedArg0 := True
__typedArg1 := V3948
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1534 {
__e.Return(True)
return
} else {
tmp1532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, V3948)
}
__typedArg0 := False
__typedArg1 := V3948
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1532 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp1535 := Call(__e, ns2_1set, symboolean_2, tmp1530)


_ = tmp1535

tmp1536 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V3949 := __self1
_ = V3949
tmp1541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V3949)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V3949
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1541 {
__e.Return(MakeInteger(0))
return
} else {
tmp1537 := Call(__e, PrimFunc(symstoutput))


tmp1538 := Call(__e, PrimFunc(sympr), MakeString("\n"), tmp1537)


_ = tmp1538

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

tmp1542 := Call(__e, ns2_1set, symnl, tmp1536)


_ = tmp1542

tmp1543 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V3956 := __self1
_ = V3956
V3957 := __self2
_ = V3957
tmp1554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3956)
}
__typedArg0 := Nil
__typedArg1 := V3956
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1554 {
__e.Return(Nil)
return
} else {
tmp1552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3956)
}
__typedArg0 := V3956
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1552 {
tmp1549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3956)
}
__typedArg0 := V3956
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1550 := Call(__e, PrimFunc(symelement_2), tmp1549, V3957)


if True == tmp1550 {
tmp1544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3956)
}
__typedArg0 := V3956
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symdifference) == __self {
__self1, __self2 = tmp1544, V3957
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symdifference), tmp1544, V3957)
return


} else {
tmp1545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3956)
}
__typedArg0 := V3956
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3956)
}
__typedArg0 := V3956
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1547 := Call(__e, PrimFunc(symdifference), tmp1546, V3957)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1545, tmp1547)
}
__typedArg0 := tmp1545
__typedArg1 := tmp1547
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

tmp1555 := Call(__e, ns2_1set, symdifference, tmp1543)


_ = tmp1555

tmp1556 := MakeNative(func(__e *ControlFlow) {
V3958 := __e.Get(1)
_ = V3958
V3959 := __e.Get(2)
_ = V3959
__e.Return(V3959)
return
}, 2)

tmp1557 := Call(__e, ns2_1set, symdo, tmp1556)


_ = tmp1557

tmp1558 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V3971 := __self1
_ = V3971
V3972 := __self2
_ = V3972
tmp1569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3972)
}
__typedArg0 := Nil
__typedArg1 := V3972
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1569 {
__e.Return(False)
return
} else {
tmp1567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3972)
}
__typedArg0 := V3972
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1563 Obj

if True == tmp1567 {
tmp1565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3972)
}
__typedArg0 := V3972
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3971, tmp1565)
}
__typedArg0 := V3971
__typedArg1 := tmp1565
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1564 Obj

if True == tmp1566 {
ifres1564 = True


} else {
ifres1564 = False


}

ifres1563 = ifres1564


} else {
ifres1563 = False


}

if True == ifres1563 {
__e.Return(True)
return
} else {
tmp1561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3972)
}
__typedArg0 := V3972
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1561 {
tmp1559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3972)
}
__typedArg0 := V3972
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symelement_2) == __self {
__self1, __self2 = V3971, tmp1559
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symelement_2), V3971, tmp1559)
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

tmp1570 := Call(__e, ns2_1set, symelement_2, tmp1558)


_ = tmp1570

tmp1571 := MakeNative(func(__e *ControlFlow) {
V3975 := __e.Get(1)
_ = V3975
tmp1573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3975)
}
__typedArg0 := Nil
__typedArg1 := V3975
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1573 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp1574 := Call(__e, ns2_1set, symempty_2, tmp1571)


_ = tmp1574

tmp1575 := MakeNative(func(__e *ControlFlow) {
V3976 := __e.Get(1)
_ = V3976
V3977 := __e.Get(2)
_ = V3977
tmp1576 := Call(__e, V3976, V3977)


__e.TailApply(PrimFunc(symshen_4fix_1help), V3976, V3977, tmp1576)
return


}, 2)

tmp1577 := Call(__e, ns2_1set, symfix, tmp1575)


_ = tmp1577

tmp1578 := MakeNative(func(__e *ControlFlow) {
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
tmp1581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3984, V3985)
}
__typedArg0 := V3984
__typedArg1 := V3985
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1581 {
__e.Return(V3985)
return
} else {
tmp1579 := Call(__e, V3983, V3985)


if PrimFunc(symshen_4fix_1help) == __self {
__self1, __self2, __self3 = V3983, V3985, tmp1579
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4fix_1help), V3983, V3985, tmp1579)
return


}


}, 3)

tmp1582 := Call(__e, ns2_1set, symshen_4fix_1help, tmp1578)


_ = tmp1582

tmp1583 := MakeNative(func(__e *ControlFlow) {
V3986 := __e.Get(1)
_ = V3986
V3987 := __e.Get(2)
_ = V3987
V3988 := __e.Get(3)
_ = V3988
V3989 := __e.Get(4)
_ = V3989
tmp1584 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4_5_1dict), V3989, V3986)
return
}, 0)

tmp1585 := MakeNative(func(__e *ControlFlow) {
Z3991 := __e.Get(1)
_ = Z3991
__e.Return(Nil)
return
}, 1)

tmp1586 := Call(__e, try_1catch, tmp1584, tmp1585)


let__1181 := tmp1586
_ = let__1181

tmp1587 := Call(__e, PrimFunc(symshen_4assoc_1set), V3987, V3988, let__1181)


let__1182 := tmp1587
_ = let__1182

tmp1588 := Call(__e, PrimFunc(symshen_4dict_1_6), V3989, V3986, let__1182)


let__1183 := tmp1588
_ = let__1183

__e.Return(V3988)
return


}, 4)

tmp1589 := Call(__e, ns2_1set, symput, tmp1583)


_ = tmp1589

tmp1590 := MakeNative(func(__e *ControlFlow) {
V3994 := __e.Get(1)
_ = V3994
V3995 := __e.Get(2)
_ = V3995
V3996 := __e.Get(3)
_ = V3996
tmp1591 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4_5_1dict), V3996, V3994)
return
}, 0)

tmp1592 := MakeNative(func(__e *ControlFlow) {
Z3998 := __e.Get(1)
_ = Z3998
__e.Return(Nil)
return
}, 1)

tmp1593 := Call(__e, try_1catch, tmp1591, tmp1592)


let__1184 := tmp1593
_ = let__1184

tmp1594 := Call(__e, PrimFunc(symshen_4assoc_1rm), V3995, let__1184)


let__1185 := tmp1594
_ = let__1185

tmp1595 := Call(__e, PrimFunc(symshen_4dict_1_6), V3996, V3994, let__1185)


let__1186 := tmp1595
_ = let__1186

__e.Return(V3994)
return


}, 3)

tmp1596 := Call(__e, ns2_1set, symunput, tmp1590)


_ = tmp1596

tmp1597 := MakeNative(func(__e *ControlFlow) {
V4001 := __e.Get(1)
_ = V4001
V4002 := __e.Get(2)
_ = V4002
V4003 := __e.Get(3)
_ = V4003
tmp1598 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4_5_1dict), V4003, V4001)
return
}, 0)

tmp1599 := MakeNative(func(__e *ControlFlow) {
Z4005 := __e.Get(1)
_ = Z4005
tmp1600 := Call(__e, PrimFunc(symshen_4app), V4002, MakeString("\n"), symshen_4s)


tmp1602 := Call(__e, PrimFunc(symshen_4app), V4001, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" has no attributes: "))
__typedS1, __typedOK1 := TypedString(tmp1600)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" has no attributes: ")
__typedArg1 := tmp1600
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp1602)
}
__typedArg0 := tmp1602
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

tmp1603 := Call(__e, try_1catch, tmp1598, tmp1599)


let__1187 := tmp1603
_ = let__1187

tmp1604 := Call(__e, PrimFunc(symassoc), V4002, let__1187)


let__1188 := tmp1604
_ = let__1188

tmp1610 := Call(__e, PrimFunc(symempty_2), let__1188)


if True == tmp1610 {
tmp1605 := Call(__e, PrimFunc(symshen_4app), V4001, MakeString("\n"), symshen_4s)


tmp1607 := Call(__e, PrimFunc(symshen_4app), V4002, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" not found for "))
__typedS1, __typedOK1 := TypedString(tmp1605)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" not found for ")
__typedArg1 := tmp1605
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attribute "))
__typedS1, __typedOK1 := TypedString(tmp1607)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attribute ")
__typedArg1 := tmp1607
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attribute "))
__typedS1, __typedOK1 := TypedString(tmp1607)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attribute ")
__typedArg1 := tmp1607
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

tmp1611 := Call(__e, ns2_1set, symget, tmp1597)


_ = tmp1611

tmp1612 := MakeNative(func(__e *ControlFlow) {
V4007 := __e.Get(1)
_ = V4007
V4008 := __e.Get(2)
_ = V4008
tmp1613 := Call(__e, PrimFunc(symshen_4hashkey), V4007)


tmp1614 := Call(__e, PrimFunc(symshen_4mod), tmp1613, V4008)


let__1189 := tmp1614
_ = let__1189

tmp1616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__1189, MakeInteger(0))
}
__typedArg0 := let__1189
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1616 {
__e.Return(MakeInteger(1))
return
} else {
__e.Return(let__1189)
return
}


}, 2)

tmp1617 := Call(__e, ns2_1set, symhash, tmp1612)


_ = tmp1617

tmp1618 := MakeNative(func(__e *ControlFlow) {
V4010 := __e.Get(1)
_ = V4010
tmp1619 := MakeNative(func(__e *ControlFlow) {
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

tmp1620 := Call(__e, PrimFunc(symexplode), V4010)


tmp1621 := Call(__e, PrimFunc(symmap), tmp1619, tmp1620)


let__1190 := tmp1621
_ = let__1190

__e.TailApply(PrimFunc(symshen_4prodbutzero), let__1190, MakeInteger(1))
return


}, 1)

tmp1622 := Call(__e, ns2_1set, symshen_4hashkey, tmp1618)


_ = tmp1622

tmp1623 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V4013 := __self1
_ = V4013
V4014 := __self2
_ = V4014
tmp1642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4013)
}
__typedArg0 := Nil
__typedArg1 := V4013
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1642 {
__e.Return(V4014)
return
} else {
tmp1640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4013)
}
__typedArg0 := V4013
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1636 Obj

if True == tmp1640 {
tmp1638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4013)
}
__typedArg0 := V4013
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), tmp1638)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp1638
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1637 Obj

if True == tmp1639 {
ifres1637 = True


} else {
ifres1637 = False


}

ifres1636 = ifres1637


} else {
ifres1636 = False


}

if True == ifres1636 {
tmp1624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4013)
}
__typedArg0 := V4013
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4prodbutzero) == __self {
__self1, __self2 = tmp1624, V4014
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp1624, V4014)
return


} else {
tmp1634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4013)
}
__typedArg0 := V4013
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1634 {
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
tmp1625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4013)
}
__typedArg0 := V4013
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4013)
}
__typedArg0 := V4013
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

if PrimFunc(symshen_4prodbutzero) == __self {
__self1, __self2 = tmp1625, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4014)
__typedN1, __typedOK1 := TypedFloat64(tmp1626)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4014
__typedArg1 := tmp1626
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp1625, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V4014)
__typedN1, __typedOK1 := TypedFloat64(tmp1626)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V4014
__typedArg1 := tmp1626
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


} else {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(V4014)
__typedN1, __typedOK1 := TypedFloat64(tmp1629)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := V4014
__typedArg1 := tmp1629
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp1628, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(V4014)
__typedN1, __typedOK1 := TypedFloat64(tmp1629)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := V4014
__typedArg1 := tmp1629
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

tmp1643 := Call(__e, ns2_1set, symshen_4prodbutzero, tmp1623)


_ = tmp1643

tmp1644 := MakeNative(func(__e *ControlFlow) {
V4015 := __e.Get(1)
_ = V4015
V4016 := __e.Get(2)
_ = V4016
tmp1645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4016, Nil)
}
__typedArg0 := V4016
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1646 := Call(__e, PrimFunc(symshen_4multiples), V4015, tmp1645)


__e.TailApply(PrimFunc(symshen_4modh), V4015, tmp1646)
return


}, 2)

tmp1647 := Call(__e, ns2_1set, symshen_4mod, tmp1644)


_ = tmp1647

tmp1648 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V4021 := __self1
_ = V4021
V4022 := __self2
_ = V4022
tmp1659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4022)
}
__typedArg0 := V4022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1655 Obj

if True == tmp1659 {
tmp1657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4022)
}
__typedArg0 := V4022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(tmp1657)
__typedN1, __typedOK1 := TypedFloat64(V4021)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := tmp1657
__typedArg1 := V4021
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})()

var ifres1656 Obj

if True == tmp1658 {
ifres1656 = True


} else {
ifres1656 = False


}

ifres1655 = ifres1656


} else {
ifres1655 = False


}

if True == ifres1655 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4022)
}
__typedArg0 := V4022
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
} else {
tmp1653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4022)
}
__typedArg0 := V4022
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1653 {
tmp1649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4022)
}
__typedArg0 := V4022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(2))
__typedN1, __typedOK1 := TypedFloat64(tmp1649)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp1649
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})(), V4022)
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(2))
__typedN1, __typedOK1 := TypedFloat64(tmp1649)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := MakeInteger(2)
__typedArg1 := tmp1649
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})()
__typedArg1 := V4022
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4multiples) == __self {
__self1, __self2 = V4021, tmp1651
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4multiples), V4021, tmp1651)
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

tmp1660 := Call(__e, ns2_1set, symshen_4multiples, tmp1648)


_ = tmp1660

tmp1661 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V4029 := __self1
_ = V4029
V4030 := __self2
_ = V4030
tmp1679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V4029)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V4029
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1679 {
__e.Return(MakeInteger(0))
return
} else {
tmp1677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4030)
}
__typedArg0 := Nil
__typedArg1 := V4030
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1677 {
__e.Return(V4029)
return
} else {
tmp1675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4030)
}
__typedArg0 := V4030
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1671 Obj

if True == tmp1675 {
tmp1673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4030)
}
__typedArg0 := V4030
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(tmp1673)
__typedN1, __typedOK1 := TypedFloat64(V4029)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := tmp1673
__typedArg1 := V4029
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})()

var ifres1672 Obj

if True == tmp1674 {
ifres1672 = True


} else {
ifres1672 = False


}

ifres1671 = ifres1672


} else {
ifres1671 = False


}

if True == ifres1671 {
tmp1664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4030)
}
__typedArg0 := V4030
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1665 := Call(__e, PrimFunc(symempty_2), tmp1664)


if True == tmp1665 {
__e.Return(V4029)
return
} else {
tmp1662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4030)
}
__typedArg0 := V4030
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4modh) == __self {
__self1, __self2 = V4029, tmp1662
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4modh), V4029, tmp1662)
return


}


} else {
tmp1669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4030)
}
__typedArg0 := V4030
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1669 {
tmp1666 := (func() Obj {
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
__typedN1, __typedOK1 := TypedFloat64(tmp1666)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V4029
__typedArg1 := tmp1666
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})(), V4030
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4modh), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V4029)
__typedN1, __typedOK1 := TypedFloat64(tmp1666)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V4029
__typedArg1 := tmp1666
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

tmp1680 := Call(__e, ns2_1set, symshen_4modh, tmp1661)


_ = tmp1680

tmp1681 := MakeNative(func(__e *ControlFlow) {
V4033 := __e.Get(1)
_ = V4033
tmp1688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4033)
}
__typedArg0 := Nil
__typedArg1 := V4033
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1688 {
__e.Return(MakeInteger(0))
return
} else {
tmp1686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4033)
}
__typedArg0 := V4033
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1686 {
tmp1682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4033)
}
__typedArg0 := V4033
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4033)
}
__typedArg0 := V4033
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1684 := Call(__e, PrimFunc(symsum), tmp1683)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(tmp1682)
__typedN1, __typedOK1 := TypedFloat64(tmp1684)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := tmp1682
__typedArg1 := tmp1684
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

tmp1689 := Call(__e, ns2_1set, symsum, tmp1681)


_ = tmp1689

tmp1690 := MakeNative(func(__e *ControlFlow) {
V4038 := __e.Get(1)
_ = V4038
tmp1692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4038)
}
__typedArg0 := V4038
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1692 {
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

tmp1693 := Call(__e, ns2_1set, symhead, tmp1690)


_ = tmp1693

tmp1694 := MakeNative(func(__e *ControlFlow) {
V4043 := __e.Get(1)
_ = V4043
tmp1696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4043)
}
__typedArg0 := V4043
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1696 {
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

tmp1697 := Call(__e, ns2_1set, symtail, tmp1694)


_ = tmp1697

tmp1698 := MakeNative(func(__e *ControlFlow) {
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

tmp1699 := Call(__e, ns2_1set, symhdstr, tmp1698)


_ = tmp1699

tmp1700 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V4051 := __self1
_ = V4051
V4052 := __self2
_ = V4052
tmp1711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4051)
}
__typedArg0 := Nil
__typedArg1 := V4051
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1711 {
__e.Return(Nil)
return
} else {
tmp1709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4051)
}
__typedArg0 := V4051
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1709 {
tmp1706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4051)
}
__typedArg0 := V4051
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1707 := Call(__e, PrimFunc(symelement_2), tmp1706, V4052)


if True == tmp1707 {
tmp1701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4051)
}
__typedArg0 := V4051
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4051)
}
__typedArg0 := V4051
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1703 := Call(__e, PrimFunc(symintersection), tmp1702, V4052)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1701, tmp1703)
}
__typedArg0 := tmp1701
__typedArg1 := tmp1703
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp1704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4051)
}
__typedArg0 := V4051
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symintersection) == __self {
__self1, __self2 = tmp1704, V4052
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symintersection), tmp1704, V4052)
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

tmp1712 := Call(__e, ns2_1set, symintersection, tmp1700)


_ = tmp1712

tmp1713 := MakeNative(func(__e *ControlFlow) {
V4053 := __e.Get(1)
_ = V4053
__e.TailApply(PrimFunc(symshen_4reverse_1help), V4053, Nil)
return
}, 1)

tmp1714 := Call(__e, ns2_1set, symreverse, tmp1713)


_ = tmp1714

tmp1715 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V4058 := __self1
_ = V4058
V4059 := __self2
_ = V4059
tmp1722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4058)
}
__typedArg0 := Nil
__typedArg1 := V4058
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1722 {
__e.Return(V4059)
return
} else {
tmp1720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4058)
}
__typedArg0 := V4058
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1720 {
tmp1716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4058)
}
__typedArg0 := V4058
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4058)
}
__typedArg0 := V4058
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1717, V4059)
}
__typedArg0 := tmp1717
__typedArg1 := V4059
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4reverse_1help) == __self {
__self1, __self2 = tmp1716, tmp1718
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4reverse_1help), tmp1716, tmp1718)
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

tmp1723 := Call(__e, ns2_1set, symshen_4reverse_1help, tmp1715)


_ = tmp1723

tmp1724 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V4064 := __self1
_ = V4064
V4065 := __self2
_ = V4065
tmp1735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4064)
}
__typedArg0 := Nil
__typedArg1 := V4064
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1735 {
__e.Return(V4065)
return
} else {
tmp1733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4064)
}
__typedArg0 := V4064
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1733 {
tmp1730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4064)
}
__typedArg0 := V4064
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1731 := Call(__e, PrimFunc(symelement_2), tmp1730, V4065)


if True == tmp1731 {
tmp1725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4064)
}
__typedArg0 := V4064
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symunion) == __self {
__self1, __self2 = tmp1725, V4065
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symunion), tmp1725, V4065)
return


} else {
tmp1726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4064)
}
__typedArg0 := V4064
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4064)
}
__typedArg0 := V4064
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1728 := Call(__e, PrimFunc(symunion), tmp1727, V4065)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1726, tmp1728)
}
__typedArg0 := tmp1726
__typedArg1 := tmp1728
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

tmp1736 := Call(__e, ns2_1set, symunion, tmp1724)


_ = tmp1736

tmp1737 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V4066 := __self1
_ = V4066
tmp1738 := Call(__e, PrimFunc(symshen_4proc_1nl), V4066)


tmp1739 := Call(__e, PrimFunc(symstoutput))


tmp1740 := Call(__e, PrimFunc(sympr), tmp1738, tmp1739)


let__1191 := tmp1740
_ = let__1191

tmp1741 := Call(__e, PrimFunc(symstoutput))


tmp1742 := Call(__e, PrimFunc(sympr), MakeString(" (y/n) "), tmp1741)


let__1192 := tmp1742
_ = let__1192

tmp1743 := Call(__e, PrimFunc(symstinput))


tmp1744 := Call(__e, PrimFunc(symread), tmp1743)


tmp1745 := Call(__e, PrimFunc(symshen_4app), tmp1744, MakeString(""), symshen_4s)


let__1193 := tmp1745
_ = let__1193

tmp1751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("y"), let__1193)
}
__typedArg0 := MakeString("y")
__typedArg1 := let__1193
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1751 {
__e.Return(True)
return
} else {
tmp1749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("n"), let__1193)
}
__typedArg0 := MakeString("n")
__typedArg1 := let__1193
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1749 {
__e.Return(False)
return
} else {
tmp1746 := Call(__e, PrimFunc(symstoutput))


tmp1747 := Call(__e, PrimFunc(sympr), MakeString("please answer y or n\n"), tmp1746)


_ = tmp1747

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

tmp1752 := Call(__e, ns2_1set, symy_1or_1n_2, tmp1737)


_ = tmp1752

tmp1753 := MakeNative(func(__e *ControlFlow) {
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

tmp1755 := Call(__e, ns2_1set, symnot, tmp1753)


_ = tmp1755

tmp1756 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString(""))
}
__typedArg0 := MakeString("")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}, 0)

tmp1757 := Call(__e, ns2_1set, symabort, tmp1756)


_ = tmp1757

tmp1758 := MakeNative(func(__e *ControlFlow) {
V4076 := __e.Get(1)
_ = V4076
V4077 := __e.Get(2)
_ = V4077
V4078 := __e.Get(3)
_ = V4078
tmp1766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V4077, V4078)
}
__typedArg0 := V4077
__typedArg1 := V4078
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1766 {
__e.Return(V4076)
return
} else {
tmp1764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4078)
}
__typedArg0 := V4078
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1764 {
tmp1759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4078)
}
__typedArg0 := V4078
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1760 := Call(__e, PrimFunc(symsubst), V4076, V4077, tmp1759)


tmp1761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4078)
}
__typedArg0 := V4078
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1762 := Call(__e, PrimFunc(symsubst), V4076, V4077, tmp1761)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1760, tmp1762)
}
__typedArg0 := tmp1760
__typedArg1 := tmp1762
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V4078)
return
}


}


}, 3)

tmp1767 := Call(__e, ns2_1set, symsubst, tmp1758)


_ = tmp1767

tmp1768 := MakeNative(func(__e *ControlFlow) {
V4079 := __e.Get(1)
_ = V4079
tmp1769 := Call(__e, PrimFunc(symshen_4app), V4079, MakeString(""), symshen_4a)


__e.TailApply(PrimFunc(symshen_4explode_1h), tmp1769)
return


}, 1)

tmp1770 := Call(__e, ns2_1set, symexplode, tmp1768)


_ = tmp1770

tmp1771 := MakeNative(func(__e *ControlFlow) {
V4082 := __e.Get(1)
_ = V4082
tmp1778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V4082)
}
__typedArg0 := MakeString("")
__typedArg1 := V4082
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1778 {
__e.Return(Nil)
return
} else {
tmp1776 := Call(__e, PrimFunc(symshen_4_7string_2), V4082)


if True == tmp1776 {
tmp1772 := Call(__e, PrimFunc(symhdstr), V4082)


tmp1774 := Call(__e, PrimFunc(symshen_4explode_1h), (func() Obj {
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
return PrimCons(tmp1772, tmp1774)
}
__typedArg0 := tmp1772
__typedArg1 := tmp1774
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

tmp1779 := Call(__e, ns2_1set, symshen_4explode_1h, tmp1771)


_ = tmp1779

tmp1780 := MakeNative(func(__e *ControlFlow) {
V4083 := __e.Get(1)
_ = V4083
tmp1783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V4083, MakeString(""))
}
__typedArg0 := V4083
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1781 Obj

if True == tmp1783 {
ifres1781 = MakeString("")


} else {
tmp1782 := Call(__e, PrimFunc(symshen_4app), V4083, MakeString("/"), symshen_4a)


ifres1781 = tmp1782


}

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dhome_1directory_d, ifres1781)
}
__typedArg0 := sym_dhome_1directory_d
__typedArg1 := ifres1781
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp1784 := Call(__e, ns2_1set, symcd, tmp1780)


_ = tmp1784

tmp1785 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V4084 := __self1
_ = V4084
V4085 := __self2
_ = V4085
tmp1792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4085)
}
__typedArg0 := Nil
__typedArg1 := V4085
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1792 {
__e.Return(True)
return
} else {
tmp1790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4085)
}
__typedArg0 := V4085
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1790 {
tmp1786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4085)
}
__typedArg0 := V4085
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1787 := Call(__e, V4084, tmp1786)


let__1194 := tmp1787
_ = let__1194

tmp1788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4085)
}
__typedArg0 := V4085
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4for_1each) == __self {
__self1, __self2 = V4084, tmp1788
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4for_1each), V4084, tmp1788)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4for_1each)
return
}


}


}, 2)

tmp1793 := Call(__e, ns2_1set, symshen_4for_1each, tmp1785)


_ = tmp1793

tmp1794 := MakeNative(func(__e *ControlFlow) {
V4087 := __e.Get(1)
_ = V4087
V4088 := __e.Get(2)
_ = V4088
__e.TailApply(PrimFunc(symshen_4map_1h), V4087, V4088, Nil)
return
}, 2)

tmp1795 := Call(__e, ns2_1set, symmap, tmp1794)


_ = tmp1795

tmp1796 := MakeNative(func(__e *ControlFlow) {
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
tmp1804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4090)
}
__typedArg0 := Nil
__typedArg1 := V4090
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1804 {
__e.TailApply(PrimFunc(symreverse), V4091)
return
} else {
tmp1802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4090)
}
__typedArg0 := V4090
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1802 {
tmp1797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4090)
}
__typedArg0 := V4090
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4090)
}
__typedArg0 := V4090
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1799 := Call(__e, V4089, tmp1798)


tmp1800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1799, V4091)
}
__typedArg0 := tmp1799
__typedArg1 := V4091
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4map_1h) == __self {
__self1, __self2, __self3 = V4089, tmp1797, tmp1800
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4map_1h), V4089, tmp1797, tmp1800)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4map_1h)
return
}


}


}, 3)

tmp1805 := Call(__e, ns2_1set, symshen_4map_1h, tmp1796)


_ = tmp1805

tmp1806 := MakeNative(func(__e *ControlFlow) {
V4092 := __e.Get(1)
_ = V4092
__e.TailApply(PrimFunc(symshen_4length_1h), V4092, MakeInteger(0))
return
}, 1)

tmp1807 := Call(__e, ns2_1set, symlength, tmp1806)


_ = tmp1807

tmp1808 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V4097 := __self1
_ = V4097
V4098 := __self2
_ = V4098
tmp1812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4097)
}
__typedArg0 := Nil
__typedArg1 := V4097
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1812 {
__e.Return(V4098)
return
} else {
tmp1809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4097)
}
__typedArg0 := V4097
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4length_1h) == __self {
__self1, __self2 = tmp1809, (func() Obj {
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
__e.TailApply(PrimFunc(symshen_4length_1h), tmp1809, (func() Obj {
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

tmp1813 := Call(__e, ns2_1set, symshen_4length_1h, tmp1808)


_ = tmp1813

tmp1814 := MakeNative(func(__e *ControlFlow) {
V4104 := __e.Get(1)
_ = V4104
V4105 := __e.Get(2)
_ = V4105
tmp1822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V4104, V4105)
}
__typedArg0 := V4104
__typedArg1 := V4105
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1822 {
__e.Return(MakeInteger(1))
return
} else {
tmp1820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4105)
}
__typedArg0 := V4105
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1820 {
tmp1815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4105)
}
__typedArg0 := V4105
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1816 := Call(__e, PrimFunc(symoccurrences), V4104, tmp1815)


tmp1817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4105)
}
__typedArg0 := V4105
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1818 := Call(__e, PrimFunc(symoccurrences), V4104, tmp1817)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(tmp1816)
__typedN1, __typedOK1 := TypedFloat64(tmp1818)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := tmp1816
__typedArg1 := tmp1818
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(MakeInteger(0))
return
}


}


}, 2)

tmp1823 := Call(__e, ns2_1set, symoccurrences, tmp1814)


_ = tmp1823

tmp1824 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V4110 := __self1
_ = V4110
V4111 := __self2
_ = V4111
tmp1837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(1), V4110)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := V4110
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1834 Obj

if True == tmp1837 {
tmp1836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4111)
}
__typedArg0 := V4111
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1835 Obj

if True == tmp1836 {
ifres1835 = True


} else {
ifres1835 = False


}

ifres1834 = ifres1835


} else {
ifres1834 = False


}

if True == ifres1834 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4111)
}
__typedArg0 := V4111
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp1832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4111)
}
__typedArg0 := V4111
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1832 {
tmp1825 := (func() Obj {
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

tmp1826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4111)
}
__typedArg0 := V4111
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symnth) == __self {
__self1, __self2 = tmp1825, tmp1826
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symnth), tmp1825, tmp1826)
return


} else {
tmp1827 := Call(__e, PrimFunc(symshen_4app), V4111, MakeString("\n"), symshen_4a)


tmp1829 := Call(__e, PrimFunc(symshen_4app), V4110, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(", "))
__typedS1, __typedOK1 := TypedString(tmp1827)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(", ")
__typedArg1 := tmp1827
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("nth applied to "))
__typedS1, __typedOK1 := TypedString(tmp1829)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("nth applied to ")
__typedArg1 := tmp1829
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("nth applied to "))
__typedS1, __typedOK1 := TypedString(tmp1829)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("nth applied to ")
__typedArg1 := tmp1829
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}


}, 2)

tmp1838 := Call(__e, ns2_1set, symnth, tmp1824)


_ = tmp1838

tmp1839 := MakeNative(func(__e *ControlFlow) {
V4112 := __e.Get(1)
_ = V4112
tmp1845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(V4112)
}
__typedArg0 := V4112
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

if True == tmp1845 {
tmp1841 := Call(__e, PrimFunc(symshen_4abs), V4112)


let__1195 := tmp1841
_ = let__1195

tmp1842 := Call(__e, PrimFunc(symshen_4magless), let__1195, MakeInteger(1))


tmp1843 := Call(__e, PrimFunc(symshen_4integer_1test_2), let__1195, tmp1842)


if True == tmp1843 {
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

tmp1846 := Call(__e, ns2_1set, syminteger_2, tmp1839)


_ = tmp1846

tmp1847 := MakeNative(func(__e *ControlFlow) {
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

tmp1850 := Call(__e, ns2_1set, symshen_4abs, tmp1847)


_ = tmp1850

tmp1851 := MakeNative(func(__e *ControlFlow) {
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

tmp1855 := Call(__e, ns2_1set, symshen_4magless, tmp1851)


_ = tmp1855

tmp1856 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V4121 := __self1
_ = V4121
V4122 := __self2
_ = V4122
tmp1863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V4121)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V4121
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1863 {
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

tmp1864 := Call(__e, ns2_1set, symshen_4integer_1test_2, tmp1856)


_ = tmp1864

tmp1865 := MakeNative(func(__e *ControlFlow) {
V4130 := __e.Get(1)
_ = V4130
V4131 := __e.Get(2)
_ = V4131
tmp1873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4131)
}
__typedArg0 := Nil
__typedArg1 := V4131
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1873 {
__e.Return(Nil)
return
} else {
tmp1871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4131)
}
__typedArg0 := V4131
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1871 {
tmp1866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4131)
}
__typedArg0 := V4131
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1867 := Call(__e, V4130, tmp1866)


tmp1868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4131)
}
__typedArg0 := V4131
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1869 := Call(__e, PrimFunc(symmapcan), V4130, tmp1868)


__e.TailApply(PrimFunc(symappend), tmp1867, tmp1869)
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

tmp1874 := Call(__e, ns2_1set, symmapcan, tmp1865)


_ = tmp1874

tmp1875 := MakeNative(func(__e *ControlFlow) {
V4137 := __e.Get(1)
_ = V4137
V4138 := __e.Get(2)
_ = V4138
tmp1877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V4137, V4138)
}
__typedArg0 := V4137
__typedArg1 := V4138
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1877 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp1878 := Call(__e, ns2_1set, sym_a_a, tmp1875)


_ = tmp1878

tmp1879 := MakeNative(func(__e *ControlFlow) {
V4139 := __e.Get(1)
_ = V4139
tmp1887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V4139)
}
__typedArg0 := V4139
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp1887 {
tmp1881 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(V4139)
}
__typedArg0 := V4139
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1882 := MakeNative(func(__e *ControlFlow) {
Z4141 := __e.Get(1)
_ = Z4141
__e.Return(symshen_4this_1symbol_1is_1unbound)
return
}, 1)

tmp1883 := Call(__e, try_1catch, tmp1881, tmp1882)


let__1198 := tmp1883
_ = let__1198

tmp1885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__1198, symshen_4this_1symbol_1is_1unbound)
}
__typedArg0 := let__1198
__typedArg1 := symshen_4this_1symbol_1is_1unbound
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1884 Obj

if True == tmp1885 {
ifres1884 = False


} else {
ifres1884 = True


}

if True == ifres1884 {
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

tmp1888 := Call(__e, ns2_1set, symbound_2, tmp1879)


_ = tmp1888

tmp1889 := MakeNative(func(__e *ControlFlow) {
V4142 := __e.Get(1)
_ = V4142
tmp1895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V4142)
}
__typedArg0 := MakeString("")
__typedArg1 := V4142
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1895 {
__e.Return(Nil)
return
} else {
tmp1890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sympos) {
return PrimPos(V4142, MakeInteger(0))
}
__typedArg0 := V4142
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sympos), __typedArg0, __typedArg1)
})()

tmp1891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp1890)
}
__typedArg0 := tmp1890
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp1893 := Call(__e, PrimFunc(symshen_4string_1_6bytes), (func() Obj {
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
return PrimCons(tmp1891, tmp1893)
}
__typedArg0 := tmp1891
__typedArg1 := tmp1893
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp1896 := Call(__e, ns2_1set, symshen_4string_1_6bytes, tmp1889)


_ = tmp1896

tmp1897 := MakeNative(func(__e *ControlFlow) {
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
tmp1899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V4143)
}
__typedArg0 := V4143
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})()

if True == tmp1899 {
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

tmp1902 := Call(__e, ns2_1set, symmaxinferences, tmp1897)


_ = tmp1902

tmp1903 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dinfs_d)
}
__typedArg0 := symshen_4_dinfs_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1904 := Call(__e, ns2_1set, syminferences, tmp1903)


_ = tmp1904

tmp1905 := MakeNative(func(__e *ControlFlow) {
V4144 := __e.Get(1)
_ = V4144
__e.Return(V4144)
return
}, 1)

tmp1906 := Call(__e, ns2_1set, symprotect, tmp1905)


_ = tmp1906

tmp1907 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dsterror_d)
}
__typedArg0 := sym_dsterror_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1908 := Call(__e, ns2_1set, symsterror, tmp1907)


_ = tmp1908

tmp1909 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dstoutput_d)
}
__typedArg0 := sym_dstoutput_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1910 := Call(__e, ns2_1set, symstoutput, tmp1909)


_ = tmp1910

tmp1911 := MakeNative(func(__e *ControlFlow) {
V4145 := __e.Get(1)
_ = V4145
tmp1912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(V4145)
}
__typedArg0 := V4145
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

let__1199 := tmp1912
_ = let__1199

tmp1916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(let__1199)
}
__typedArg0 := let__1199
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp1916 {
__e.Return(let__1199)
return
} else {
tmp1913 := Call(__e, PrimFunc(symshen_4app), V4145, MakeString(" to a symbol"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("cannot intern "))
__typedS1, __typedOK1 := TypedString(tmp1913)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("cannot intern ")
__typedArg1 := tmp1913
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("cannot intern "))
__typedS1, __typedOK1 := TypedString(tmp1913)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("cannot intern ")
__typedArg1 := tmp1913
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 1)

tmp1917 := Call(__e, ns2_1set, symstring_1_6symbol, tmp1911)


_ = tmp1917

tmp1918 := MakeNative(func(__e *ControlFlow) {
V4149 := __e.Get(1)
_ = V4149
tmp1922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V4149)
}
__typedArg0 := sym_7
__typedArg1 := V4149
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1922 {
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
tmp1920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V4149)
}
__typedArg0 := sym_1
__typedArg1 := V4149
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1920 {
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

tmp1923 := Call(__e, ns2_1set, symoptimise, tmp1918)


_ = tmp1923

tmp1924 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dos_d)
}
__typedArg0 := sym_dos_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1925 := Call(__e, ns2_1set, symos, tmp1924)


_ = tmp1925

tmp1926 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dlanguage_d)
}
__typedArg0 := sym_dlanguage_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1927 := Call(__e, ns2_1set, symlanguage, tmp1926)


_ = tmp1927

tmp1928 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dversion_d)
}
__typedArg0 := sym_dversion_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1929 := Call(__e, ns2_1set, symversion, tmp1928)


_ = tmp1929

tmp1930 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dport_d)
}
__typedArg0 := sym_dport_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1931 := Call(__e, ns2_1set, symport, tmp1930)


_ = tmp1931

tmp1932 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dporters_d)
}
__typedArg0 := sym_dporters_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1933 := Call(__e, ns2_1set, symporters, tmp1932)


_ = tmp1933

tmp1934 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dimplementation_d)
}
__typedArg0 := sym_dimplementation_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1935 := Call(__e, ns2_1set, symimplementation, tmp1934)


_ = tmp1935

tmp1936 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_drelease_d)
}
__typedArg0 := sym_drelease_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1937 := Call(__e, ns2_1set, symrelease, tmp1936)


_ = tmp1937

tmp1938 := MakeNative(func(__e *ControlFlow) {
V4150 := __e.Get(1)
_ = V4150
tmp1943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnull, V4150)
}
__typedArg0 := symnull
__typedArg1 := V4150
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1943 {
__e.Return(True)
return
} else {
tmp1939 := MakeNative(func(__e *ControlFlow) {
tmp1940 := Call(__e, PrimFunc(symexternal), V4150)


_ = tmp1940

__e.Return(True)
return


}, 0)

tmp1941 := MakeNative(func(__e *ControlFlow) {
Z4151 := __e.Get(1)
_ = Z4151
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1939, tmp1941)
return


}


}, 1)

tmp1944 := Call(__e, ns2_1set, sympackage_2, tmp1938)


_ = tmp1944

tmp1945 := MakeNative(func(__e *ControlFlow) {
__e.Return(symshen_4fail_b)
return
}, 0)

tmp1946 := Call(__e, ns2_1set, symfail, tmp1945)


_ = tmp1946

tmp1947 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_duserdefs_d)
}
__typedArg0 := symshen_4_duserdefs_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1948 := Call(__e, ns2_1set, symuserdefs, tmp1947)


_ = tmp1948

tmp1949 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_doptimise_d)
}
__typedArg0 := symshen_4_doptimise_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1950 := Call(__e, ns2_1set, symoptimise_2, tmp1949)


_ = tmp1950

tmp1951 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dhush_d)
}
__typedArg0 := sym_dhush_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1952 := Call(__e, ns2_1set, symhush_2, tmp1951)


_ = tmp1952

tmp1953 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dshen_1type_1theory_1enabled_2_d)
}
__typedArg0 := symshen_4_dshen_1type_1theory_1enabled_2_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1954 := Call(__e, ns2_1set, symsystem_1S_2, tmp1953)


_ = tmp1954

tmp1955 := MakeNative(func(__e *ControlFlow) {
V4154 := __e.Get(1)
_ = V4154
tmp1959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V4154)
}
__typedArg0 := sym_7
__typedArg1 := V4154
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1959 {
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
tmp1957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V4154)
}
__typedArg0 := sym_1
__typedArg1 := V4154
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1957 {
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

tmp1960 := Call(__e, ns2_1set, symenable_1type_1theory, tmp1955)


_ = tmp1960

tmp1961 := MakeNative(func(__e *ControlFlow) {
V4157 := __e.Get(1)
_ = V4157
tmp1965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V4157)
}
__typedArg0 := sym_7
__typedArg1 := V4157
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1965 {
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
tmp1963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V4157)
}
__typedArg0 := sym_1
__typedArg1 := V4157
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1963 {
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

tmp1966 := Call(__e, ns2_1set, symhush, tmp1961)


_ = tmp1966

tmp1967 := MakeNative(func(__e *ControlFlow) {
V4160 := __e.Get(1)
_ = V4160
tmp1971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V4160)
}
__typedArg0 := sym_7
__typedArg1 := V4160
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1971 {
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
tmp1969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V4160)
}
__typedArg0 := sym_1
__typedArg1 := V4160
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1969 {
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

tmp1972 := Call(__e, ns2_1set, symtc, tmp1967)


_ = tmp1972

tmp1973 := MakeNative(func(__e *ControlFlow) {
V4161 := __e.Get(1)
_ = V4161
tmp1974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp1975 := Call(__e, PrimFunc(symshen_4unassoc), V4161, tmp1974)


tmp1976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp1975)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp1975
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp1976

__e.Return(V4161)
return


}, 1)

tmp1977 := Call(__e, ns2_1set, symdestroy, tmp1973)


_ = tmp1977

tmp1978 := MakeNative(func(__e *ControlFlow) {
V4171 := __e.Get(1)
_ = V4171
V4172 := __e.Get(2)
_ = V4172
tmp1996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V4172)
}
__typedArg0 := Nil
__typedArg1 := V4172
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1996 {
__e.Return(Nil)
return
} else {
tmp1994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4172)
}
__typedArg0 := V4172
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1985 Obj

if True == tmp1994 {
tmp1992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4172)
}
__typedArg0 := V4172
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp1992)
}
__typedArg0 := tmp1992
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1987 Obj

if True == tmp1993 {
tmp1989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4172)
}
__typedArg0 := V4172
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1989)
}
__typedArg0 := tmp1989
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V4171, tmp1990)
}
__typedArg0 := V4171
__typedArg1 := tmp1990
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1988 Obj

if True == tmp1991 {
ifres1988 = True


} else {
ifres1988 = False


}

ifres1987 = ifres1988


} else {
ifres1987 = False


}

var ifres1986 Obj

if True == ifres1987 {
ifres1986 = True


} else {
ifres1986 = False


}

ifres1985 = ifres1986


} else {
ifres1985 = False


}

if True == ifres1985 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4172)
}
__typedArg0 := V4172
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
} else {
tmp1983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V4172)
}
__typedArg0 := V4172
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1983 {
tmp1979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V4172)
}
__typedArg0 := V4172
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V4172)
}
__typedArg0 := V4172
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1981 := Call(__e, PrimFunc(symshen_4unassoc), V4171, tmp1980)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1979, tmp1981)
}
__typedArg0 := tmp1979
__typedArg1 := tmp1981
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

tmp1997 := Call(__e, ns2_1set, symshen_4unassoc, tmp1978)


_ = tmp1997

tmp1998 := MakeNative(func(__e *ControlFlow) {
V4173 := __e.Get(1)
_ = V4173
tmp2002 := Call(__e, PrimFunc(sympackage_2), V4173)


if True == tmp2002 {
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
tmp1999 := Call(__e, PrimFunc(symshen_4app), V4173, MakeString(" does not exist\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("package "))
__typedS1, __typedOK1 := TypedString(tmp1999)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("package ")
__typedArg1 := tmp1999
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("package "))
__typedS1, __typedOK1 := TypedString(tmp1999)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("package ")
__typedArg1 := tmp1999
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 1)

tmp2003 := Call(__e, ns2_1set, symin_1package, tmp1998)


_ = tmp2003

tmp2004 := MakeNative(func(__e *ControlFlow) {
V4174 := __e.Get(1)
_ = V4174
V4175 := __e.Get(2)
_ = V4175
tmp2005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symopen) {
return PrimOpenStream(V4174, symout)
}
__typedArg0 := V4174
__typedArg1 := symout
return Call(__e, PrimFunc(symopen), __typedArg0, __typedArg1)
})()

let__1200 := tmp2005
_ = let__1200

tmp2008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V4175)
}
__typedArg0 := V4175
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres2006 Obj

if True == tmp2008 {
ifres2006 = V4175


} else {
tmp2007 := Call(__e, PrimFunc(symshen_4app), V4175, MakeString(""), symshen_4s)


ifres2006 = tmp2007


}

let__1201 := ifres2006
_ = let__1201

tmp2009 := Call(__e, PrimFunc(sympr), let__1201, let__1200)


let__1202 := tmp2009
_ = let__1202

tmp2010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symclose) {
return PrimCloseStream(let__1200)
}
__typedArg0 := let__1200
return Call(__e, PrimFunc(symclose), __typedArg0)
})()

let__1203 := tmp2010
_ = let__1203

__e.Return(V4175)
return


}, 2)

tmp2011 := Call(__e, ns2_1set, symwrite_1to_1file, tmp2004)


_ = tmp2011

tmp2012 := MakeNative(func(__e *ControlFlow) {
tmp2013 := Call(__e, PrimFunc(symgensym), symshen_4t)


__e.TailApply(PrimFunc(symshen_4freshterm), tmp2013)
return


}, 0)

tmp2014 := Call(__e, ns2_1set, symfresh, tmp2012)


_ = tmp2014

tmp2015 := MakeNative(func(__e *ControlFlow) {
V4180 := __e.Get(1)
_ = V4180
V4181 := __e.Get(2)
_ = V4181
tmp2016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2017 := Call(__e, PrimFunc(symput), V4180, symarity, V4181, tmp2016)


let__1204 := tmp2017
_ = let__1204

tmp2018 := Call(__e, PrimFunc(symshen_4lambda_1entry), V4180)


let__1205 := tmp2018
_ = let__1205

tmp2023 := Call(__e, PrimFunc(symempty_2), let__1205)


var ifres2019 Obj

if True == tmp2023 {
tmp2020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2021 := Call(__e, PrimFunc(symunput), V4180, symshen_4lambda_1form, tmp2020)


ifres2019 = tmp2021


} else {
tmp2022 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), let__1205)


ifres2019 = tmp2022


}

let__1206 := ifres2019
_ = let__1206

__e.Return(V4180)
return


}, 2)

tmp2024 := Call(__e, ns2_1set, symupdate_1lambda_1table, tmp2015)


_ = tmp2024

tmp2025 := MakeNative(func(__e *ControlFlow) {
V4187 := __e.Get(1)
_ = V4187
V4188 := __e.Get(2)
_ = V4188
tmp2049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V4188)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V4188
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2049 {
tmp2026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspecial_d)
}
__typedArg0 := symshen_4_dspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2027 := Call(__e, PrimFunc(symremove), V4187, tmp2026)


tmp2028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspecial_d, tmp2027)
}
__typedArg0 := symshen_4_dspecial_d
__typedArg1 := tmp2027
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2028

tmp2029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dextraspecial_d)
}
__typedArg0 := symshen_4_dextraspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2030 := Call(__e, PrimFunc(symremove), V4187, tmp2029)


tmp2031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dextraspecial_d, tmp2030)
}
__typedArg0 := symshen_4_dextraspecial_d
__typedArg1 := tmp2030
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2031

__e.Return(V4187)
return


} else {
tmp2047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(1), V4188)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := V4188
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2047 {
tmp2032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspecial_d)
}
__typedArg0 := symshen_4_dspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2033 := Call(__e, PrimFunc(symadjoin), V4187, tmp2032)


tmp2034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspecial_d, tmp2033)
}
__typedArg0 := symshen_4_dspecial_d
__typedArg1 := tmp2033
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2034

tmp2035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dextraspecial_d)
}
__typedArg0 := symshen_4_dextraspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2036 := Call(__e, PrimFunc(symremove), V4187, tmp2035)


tmp2037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dextraspecial_d, tmp2036)
}
__typedArg0 := symshen_4_dextraspecial_d
__typedArg1 := tmp2036
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2037

__e.Return(V4187)
return


} else {
tmp2045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(2), V4188)
}
__typedArg0 := MakeInteger(2)
__typedArg1 := V4188
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2045 {
tmp2038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspecial_d)
}
__typedArg0 := symshen_4_dspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2039 := Call(__e, PrimFunc(symremove), V4187, tmp2038)


tmp2040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspecial_d, tmp2039)
}
__typedArg0 := symshen_4_dspecial_d
__typedArg1 := tmp2039
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2040

tmp2041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dextraspecial_d)
}
__typedArg0 := symshen_4_dextraspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2042 := Call(__e, PrimFunc(symadjoin), V4187, tmp2041)


tmp2043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dextraspecial_d, tmp2042)
}
__typedArg0 := symshen_4_dextraspecial_d
__typedArg1 := tmp2042
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2043

__e.Return(V4187)
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

tmp2050 := Call(__e, ns2_1set, symspecialise, tmp2025)


_ = tmp2050

tmp2051 := MakeNative(func(__e *ControlFlow) {
V4189 := __e.Get(1)
_ = V4189
tmp2052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dabsolute_d)
}
__typedArg0 := sym_dabsolute_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V4189, tmp2052)
}
__typedArg0 := V4189
__typedArg1 := tmp2052
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dabsolute_d, tmp2053)
}
__typedArg0 := sym_dabsolute_d
__typedArg1 := tmp2053
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp2054 := Call(__e, ns2_1set, symabsolute, tmp2051)


_ = tmp2054

tmp2055 := MakeNative(func(__e *ControlFlow) {
V4190 := __e.Get(1)
_ = V4190
tmp2056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dabsolute_d)
}
__typedArg0 := sym_dabsolute_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2057 := Call(__e, PrimFunc(symremove), V4190, tmp2056)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dabsolute_d, tmp2057)
}
__typedArg0 := sym_dabsolute_d
__typedArg1 := tmp2057
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 1)

__e.TailApply(ns2_1set, symunabsolute, tmp2055)
return




}, 0)

