package main

import . "github.com/pyrex41/shen-go/kl"

var SysMain = MakeNative(func(__e *ControlFlow) {
tmp1206 := MakeNative(func(__e *ControlFlow) {
V3444 := __e.Get(1)
_ = V3444
__e.TailApply(V3444)
return
}, 1)

tmp1207 := Call(__e, ns2_1set, symthaw, tmp1206)


_ = tmp1207

tmp1208 := MakeNative(func(__e *ControlFlow) {
V3445 := __e.Get(1)
_ = V3445
tmp1209 := Call(__e, PrimFunc(symmacroexpand), V3445)


tmp1210 := Call(__e, PrimFunc(symshen_4find_1types), V3445)


tmp1211 := Call(__e, PrimFunc(symshen_4process_1applications), tmp1209, tmp1210)


tmp1212 := Call(__e, PrimFunc(symshen_4shen_1_6kl), tmp1211)


__e.TailApply(PrimFunc(symeval_1kl), tmp1212)
return


}, 1)

tmp1213 := Call(__e, ns2_1set, symeval, tmp1208)


_ = tmp1213

tmp1214 := MakeNative(func(__e *ControlFlow) {
V3446 := __e.Get(1)
_ = V3446
tmp1221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnull, V3446)
}
__typedArg0 := symnull
__typedArg1 := V3446
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1221 {
__e.Return(Nil)
return
} else {
tmp1215 := MakeNative(func(__e *ControlFlow) {
tmp1216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V3446, symshen_4external_1symbols, tmp1216)
return


}, 0)

tmp1217 := MakeNative(func(__e *ControlFlow) {
Z3447 := __e.Get(1)
_ = Z3447
tmp1218 := Call(__e, PrimFunc(symshen_4app), V3446, MakeString(" does not exist.\n;"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("package "))
__typedS1, __typedOK1 := TypedString(tmp1218)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("package ")
__typedArg1 := tmp1218
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("package "))
__typedS1, __typedOK1 := TypedString(tmp1218)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("package ")
__typedArg1 := tmp1218
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

__e.TailApply(try_1catch, tmp1215, tmp1217)
return


}


}, 1)

tmp1222 := Call(__e, ns2_1set, symexternal, tmp1214)


_ = tmp1222

tmp1223 := MakeNative(func(__e *ControlFlow) {
V3448 := __e.Get(1)
_ = V3448
tmp1230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnull, V3448)
}
__typedArg0 := symnull
__typedArg1 := V3448
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1230 {
__e.Return(Nil)
return
} else {
tmp1224 := MakeNative(func(__e *ControlFlow) {
tmp1225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V3448, symshen_4internal_1symbols, tmp1225)
return


}, 0)

tmp1226 := MakeNative(func(__e *ControlFlow) {
Z3449 := __e.Get(1)
_ = Z3449
tmp1227 := Call(__e, PrimFunc(symshen_4app), V3448, MakeString(" does not exist.\n;"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("package "))
__typedS1, __typedOK1 := TypedString(tmp1227)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("package ")
__typedArg1 := tmp1227
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("package "))
__typedS1, __typedOK1 := TypedString(tmp1227)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("package ")
__typedArg1 := tmp1227
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

__e.TailApply(try_1catch, tmp1224, tmp1226)
return


}


}, 1)

tmp1231 := Call(__e, ns2_1set, syminternal, tmp1223)


_ = tmp1231

tmp1232 := MakeNative(func(__e *ControlFlow) {
V3450 := __e.Get(1)
_ = V3450
V3451 := __e.Get(2)
_ = V3451
tmp1234 := Call(__e, V3450, V3451)


if True == tmp1234 {
__e.TailApply(PrimFunc(symfail))
return
} else {
__e.Return(V3451)
return
}


}, 2)

tmp1235 := Call(__e, ns2_1set, symfail_1if, tmp1232)


_ = tmp1235

tmp1236 := MakeNative(func(__e *ControlFlow) {
V3452 := __e.Get(1)
_ = V3452
V3453 := __e.Get(2)
_ = V3453
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(V3452)
__typedS1, __typedOK1 := TypedString(V3453)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V3452
__typedArg1 := V3453
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return
}, 2)

tmp1237 := Call(__e, ns2_1set, sym_8s, tmp1236)


_ = tmp1237

tmp1238 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtc_d)
}
__typedArg0 := symshen_4_dtc_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1239 := Call(__e, ns2_1set, symtc_2, tmp1238)


_ = tmp1239

tmp1240 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_doccurs_d)
}
__typedArg0 := symshen_4_doccurs_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1241 := Call(__e, ns2_1set, symoccurs_2, tmp1240)


_ = tmp1241

tmp1242 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dfactorise_2_d)
}
__typedArg0 := symshen_4_dfactorise_2_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1243 := Call(__e, ns2_1set, symfactorise_2, tmp1242)


_ = tmp1243

tmp1244 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtracking_d)
}
__typedArg0 := symshen_4_dtracking_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1245 := Call(__e, ns2_1set, symtracked, tmp1244)


_ = tmp1245

tmp1246 := MakeNative(func(__e *ControlFlow) {
V3454 := __e.Get(1)
_ = V3454
tmp1247 := MakeNative(func(__e *ControlFlow) {
tmp1248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V3454, symshen_4source, tmp1248)
return


}, 0)

tmp1249 := MakeNative(func(__e *ControlFlow) {
Z3455 := __e.Get(1)
_ = Z3455
tmp1250 := Call(__e, PrimFunc(symshen_4app), V3454, MakeString(" not found.\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp1250)
}
__typedArg0 := tmp1250
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

__e.TailApply(try_1catch, tmp1247, tmp1249)
return


}, 1)

tmp1251 := Call(__e, ns2_1set, symps, tmp1246)


_ = tmp1251

tmp1252 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dstinput_d)
}
__typedArg0 := sym_dstinput_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1253 := Call(__e, ns2_1set, symstinput, tmp1252)


_ = tmp1253

tmp1254 := MakeNative(func(__e *ControlFlow) {
V3456 := __e.Get(1)
_ = V3456
tmp1256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V3456)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V3456
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V3456)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V3456
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

let__1164 := tmp1256
_ = let__1164

tmp1257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__1164, MakeNumber(0), V3456)
}
__typedArg0 := let__1164
__typedArg1 := MakeNumber(0)
__typedArg2 := V3456
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__1165 := tmp1257
_ = let__1165

tmp1261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3456, MakeNumber(0))
}
__typedArg0 := V3456
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1258 Obj

if True == tmp1261 {
ifres1258 = let__1165


} else {
tmp1259 := Call(__e, PrimFunc(symfail))


tmp1260 := Call(__e, PrimFunc(symshen_4fillvector), let__1165, MakeNumber(1), V3456, tmp1259)


ifres1258 = tmp1260


}

let__1166 := ifres1258
_ = let__1166

__e.Return(let__1166)
return


}, 1)

tmp1262 := Call(__e, ns2_1set, symvector, tmp1254)


_ = tmp1262

tmp1263 := MakeNative(func(__e *ControlFlow) {
V3461 := __e.Get(1)
_ = V3461
V3462 := __e.Get(2)
_ = V3462
V3463 := __e.Get(3)
_ = V3463
V3464 := __e.Get(4)
_ = V3464
tmp1267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3462, V3463)
}
__typedArg0 := V3462
__typedArg1 := V3463
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1267 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V3461, V3463, V3464)
}
__typedArg0 := V3461
__typedArg1 := V3463
__typedArg2 := V3464
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return
} else {
tmp1264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V3461, V3462, V3464)
}
__typedArg0 := V3461
__typedArg1 := V3462
__typedArg2 := V3464
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

__e.TailApply(PrimFunc(symshen_4fillvector), tmp1264, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(V3462)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeNumber(1)
__typedArg1 := V3462
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})(), V3463, V3464)
return


}


}, 4)

tmp1268 := Call(__e, ns2_1set, symshen_4fillvector, tmp1263)


_ = tmp1268

tmp1269 := MakeNative(func(__e *ControlFlow) {
V3465 := __e.Get(1)
_ = V3465
tmp1276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V3465)
}
__typedArg0 := V3465
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp1276 {
tmp1271 := MakeNative(func(__e *ControlFlow) {
tmp1272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V3465, MakeNumber(0))
}
__typedArg0 := V3465
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6_a) {
__typedN0, __typedOK0 := TypedFloat64(tmp1272)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(0))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6_a) {
return TypedMaterializeBoolean((__typedN0 >= __typedN1))
}}
__typedArg0 := tmp1272
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_6_a), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp1273 := MakeNative(func(__e *ControlFlow) {
Z3466 := __e.Get(1)
_ = Z3466
__e.Return(False)
return
}, 1)

tmp1274 := Call(__e, try_1catch, tmp1271, tmp1273)


if True == tmp1274 {
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

tmp1277 := Call(__e, ns2_1set, symvector_2, tmp1269)


_ = tmp1277

tmp1278 := MakeNative(func(__e *ControlFlow) {
V3467 := __e.Get(1)
_ = V3467
V3468 := __e.Get(2)
_ = V3468
V3469 := __e.Get(3)
_ = V3469
tmp1280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3468, MakeNumber(0))
}
__typedArg0 := V3468
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1280 {
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
return PrimVectorSet(V3467, V3468, V3469)
}
__typedArg0 := V3467
__typedArg1 := V3468
__typedArg2 := V3469
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return
}


}, 3)

tmp1281 := Call(__e, ns2_1set, symvector_1_6, tmp1278)


_ = tmp1281

tmp1282 := MakeNative(func(__e *ControlFlow) {
V3470 := __e.Get(1)
_ = V3470
V3471 := __e.Get(2)
_ = V3471
tmp1288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3471, MakeNumber(0))
}
__typedArg0 := V3471
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1288 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("cannot access 0th element of a vector\n"))
}
__typedArg0 := MakeString("cannot access 0th element of a vector\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp1283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V3470, V3471)
}
__typedArg0 := V3470
__typedArg1 := V3471
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

let__1167 := tmp1283
_ = let__1167

tmp1285 := Call(__e, PrimFunc(symfail))


tmp1286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__1167, tmp1285)
}
__typedArg0 := let__1167
__typedArg1 := tmp1285
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1286 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("vector element not found\n"))
}
__typedArg0 := MakeString("vector element not found\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
__e.Return(let__1167)
return
}


}


}, 2)

tmp1289 := Call(__e, ns2_1set, sym_5_1vector, tmp1282)


_ = tmp1289

tmp1290 := MakeNative(func(__e *ControlFlow) {
V3473 := __e.Get(1)
_ = V3473
tmp1294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V3473)
}
__typedArg0 := V3473
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})()

if True == tmp1294 {
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6_a) {
__typedN0, __typedOK0 := TypedFloat64(V3473)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(0))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6_a) {
return TypedMaterializeBoolean((__typedN0 >= __typedN1))
}}
__typedArg0 := V3473
__typedArg1 := MakeNumber(0)
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

tmp1295 := Call(__e, ns2_1set, symshen_4posint_2, tmp1290)


_ = tmp1295

tmp1296 := MakeNative(func(__e *ControlFlow) {
V3474 := __e.Get(1)
_ = V3474
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V3474, MakeNumber(0))
}
__typedArg0 := V3474
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp1297 := Call(__e, ns2_1set, symlimit, tmp1296)


_ = tmp1297

tmp1298 := MakeNative(func(__e *ControlFlow) {
V3475 := __e.Get(1)
_ = V3475
tmp1328 := Call(__e, PrimFunc(symboolean_2), V3475)


var ifres1313 Obj

if True == tmp1328 {
ifres1313 = True


} else {
tmp1327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(V3475)
}
__typedArg0 := V3475
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

var ifres1315 Obj

if True == tmp1327 {
ifres1315 = True


} else {
tmp1326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V3475)
}
__typedArg0 := V3475
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres1317 Obj

if True == tmp1326 {
ifres1317 = True


} else {
tmp1325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3475)
}
__typedArg0 := V3475
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1319 Obj

if True == tmp1325 {
ifres1319 = True


} else {
tmp1324 := Call(__e, PrimFunc(symempty_2), V3475)


var ifres1321 Obj

if True == tmp1324 {
ifres1321 = True


} else {
tmp1323 := Call(__e, PrimFunc(symvector_2), V3475)


var ifres1322 Obj

if True == tmp1323 {
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

var ifres1316 Obj

if True == ifres1317 {
ifres1316 = True


} else {
ifres1316 = False


}

ifres1315 = ifres1316


}

var ifres1314 Obj

if True == ifres1315 {
ifres1314 = True


} else {
ifres1314 = False


}

ifres1313 = ifres1314


}

if True == ifres1313 {
__e.Return(False)
return
} else {
tmp1303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp1304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp1305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(","))
}
__typedArg0 := MakeString(",")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp1306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1305, Nil)
}
__typedArg0 := tmp1305
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1304, tmp1306)
}
__typedArg0 := tmp1304
__typedArg1 := tmp1306
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1303, tmp1307)
}
__typedArg0 := tmp1303
__typedArg1 := tmp1307
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, tmp1308)
}
__typedArg0 := sym_j
__typedArg1 := tmp1308
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, tmp1309)
}
__typedArg0 := sym_i
__typedArg1 := tmp1309
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1311 := Call(__e, PrimFunc(symelement_2), V3475, tmp1310)


if True == tmp1311 {
__e.Return(True)
return
} else {
tmp1299 := MakeNative(func(__e *ControlFlow) {
tmp1300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3475)
}
__typedArg0 := V3475
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

let__1168 := tmp1300
_ = let__1168

__e.TailApply(PrimFunc(symshen_4analyse_1symbol_2), let__1168)
return


}, 0)

tmp1301 := MakeNative(func(__e *ControlFlow) {
Z3477 := __e.Get(1)
_ = Z3477
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1299, tmp1301)
return


}


}


}, 1)

tmp1329 := Call(__e, ns2_1set, symsymbol_2, tmp1298)


_ = tmp1329

tmp1330 := MakeNative(func(__e *ControlFlow) {
V3480 := __e.Get(1)
_ = V3480
tmp1339 := Call(__e, PrimFunc(symshen_4_7string_2), V3480)


if True == tmp1339 {
tmp1335 := Call(__e, PrimFunc(symhdstr), V3480)


tmp1336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp1335)
}
__typedArg0 := tmp1335
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp1337 := Call(__e, PrimFunc(symshen_4alpha_2), tmp1336)


if True == tmp1337 {
tmp1333 := Call(__e, PrimFunc(symshen_4alphanums_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3480)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3480
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


if True == tmp1333 {
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

tmp1340 := Call(__e, ns2_1set, symshen_4analyse_1symbol_2, tmp1330)


_ = tmp1340

tmp1341 := MakeNative(func(__e *ControlFlow) {
V3483 := __e.Get(1)
_ = V3483
tmp1355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V3483)
}
__typedArg0 := MakeString("")
__typedArg1 := V3483
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1355 {
__e.Return(True)
return
} else {
tmp1353 := Call(__e, PrimFunc(symshen_4_7string_2), V3483)


if True == tmp1353 {
tmp1342 := Call(__e, PrimFunc(symhdstr), V3483)


tmp1343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp1342)
}
__typedArg0 := tmp1342
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

let__1169 := tmp1343
_ = let__1169

tmp1351 := Call(__e, PrimFunc(symshen_4alpha_2), let__1169)


var ifres1348 Obj

if True == tmp1351 {
ifres1348 = True


} else {
tmp1350 := Call(__e, PrimFunc(symshen_4digit_2), let__1169)


var ifres1349 Obj

if True == tmp1350 {
ifres1349 = True


} else {
ifres1349 = False


}

ifres1348 = ifres1349


}

if True == ifres1348 {
tmp1346 := Call(__e, PrimFunc(symshen_4alphanums_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3483)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3483
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


if True == tmp1346 {
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

tmp1356 := Call(__e, ns2_1set, symshen_4alphanums_2, tmp1341)


_ = tmp1356

tmp1357 := MakeNative(func(__e *ControlFlow) {
V3485 := __e.Get(1)
_ = V3485
tmp1368 := Call(__e, PrimFunc(symboolean_2), V3485)


var ifres1362 Obj

if True == tmp1368 {
ifres1362 = True


} else {
tmp1367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(V3485)
}
__typedArg0 := V3485
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

var ifres1364 Obj

if True == tmp1367 {
ifres1364 = True


} else {
tmp1366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V3485)
}
__typedArg0 := V3485
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres1365 Obj

if True == tmp1366 {
ifres1365 = True


} else {
ifres1365 = False


}

ifres1364 = ifres1365


}

var ifres1363 Obj

if True == ifres1364 {
ifres1363 = True


} else {
ifres1363 = False


}

ifres1362 = ifres1363


}

if True == ifres1362 {
__e.Return(False)
return
} else {
tmp1358 := MakeNative(func(__e *ControlFlow) {
tmp1359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3485)
}
__typedArg0 := V3485
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

let__1170 := tmp1359
_ = let__1170

__e.TailApply(PrimFunc(symshen_4analyse_1variable_2), let__1170)
return


}, 0)

tmp1360 := MakeNative(func(__e *ControlFlow) {
Z3487 := __e.Get(1)
_ = Z3487
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1358, tmp1360)
return


}


}, 1)

tmp1369 := Call(__e, ns2_1set, symvariable_2, tmp1357)


_ = tmp1369

tmp1370 := MakeNative(func(__e *ControlFlow) {
V3490 := __e.Get(1)
_ = V3490
tmp1379 := Call(__e, PrimFunc(symshen_4_7string_2), V3490)


if True == tmp1379 {
tmp1375 := Call(__e, PrimFunc(symhdstr), V3490)


tmp1376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp1375)
}
__typedArg0 := tmp1375
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp1377 := Call(__e, PrimFunc(symshen_4uppercase_2), tmp1376)


if True == tmp1377 {
tmp1373 := Call(__e, PrimFunc(symshen_4alphanums_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3490)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3490
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


if True == tmp1373 {
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

tmp1380 := Call(__e, ns2_1set, symshen_4analyse_1variable_2, tmp1370)


_ = tmp1380

tmp1381 := MakeNative(func(__e *ControlFlow) {
V3491 := __e.Get(1)
_ = V3491
tmp1382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dgensym_d)
}
__typedArg0 := symshen_4_dgensym_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp1384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dgensym_d, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp1382)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp1382
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := symshen_4_dgensym_d
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp1382)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp1382
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symconcat), V3491, tmp1384)
return


}, 1)

tmp1385 := Call(__e, ns2_1set, symgensym, tmp1381)


_ = tmp1385

tmp1386 := MakeNative(func(__e *ControlFlow) {
V3492 := __e.Get(1)
_ = V3492
V3493 := __e.Get(2)
_ = V3493
tmp1387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3492)
}
__typedArg0 := V3492
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp1388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3493)
}
__typedArg0 := V3493
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp1387)
__typedS1, __typedOK1 := TypedString(tmp1388)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp1387
__typedArg1 := tmp1388
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp1387)
__typedS1, __typedOK1 := TypedString(tmp1388)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp1387
__typedArg1 := tmp1388
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symintern), __typedArg0)
})())
return


}, 2)

tmp1390 := Call(__e, ns2_1set, symconcat, tmp1386)


_ = tmp1390

tmp1391 := MakeNative(func(__e *ControlFlow) {
V3494 := __e.Get(1)
_ = V3494
V3495 := __e.Get(2)
_ = V3495
tmp1392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(MakeNumber(3))
}
__typedArg0 := MakeNumber(3)
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

let__1171 := tmp1392
_ = let__1171

tmp1393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__1171, MakeNumber(0), symshen_4tuple)
}
__typedArg0 := let__1171
__typedArg1 := MakeNumber(0)
__typedArg2 := symshen_4tuple
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__1172 := tmp1393
_ = let__1172

tmp1394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__1171, MakeNumber(1), V3494)
}
__typedArg0 := let__1171
__typedArg1 := MakeNumber(1)
__typedArg2 := V3494
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__1173 := tmp1394
_ = let__1173

tmp1395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__1171, MakeNumber(2), V3495)
}
__typedArg0 := let__1171
__typedArg1 := MakeNumber(2)
__typedArg2 := V3495
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__1174 := tmp1395
_ = let__1174

__e.Return(let__1171)
return


}, 2)

tmp1396 := Call(__e, ns2_1set, sym_8p, tmp1391)


_ = tmp1396

tmp1397 := MakeNative(func(__e *ControlFlow) {
V3500 := __e.Get(1)
_ = V3500
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V3500, MakeNumber(1))
}
__typedArg0 := V3500
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp1398 := Call(__e, ns2_1set, symfst, tmp1397)


_ = tmp1398

tmp1399 := MakeNative(func(__e *ControlFlow) {
V3501 := __e.Get(1)
_ = V3501
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V3501, MakeNumber(2))
}
__typedArg0 := V3501
__typedArg1 := MakeNumber(2)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp1400 := Call(__e, ns2_1set, symsnd, tmp1399)


_ = tmp1400

tmp1401 := MakeNative(func(__e *ControlFlow) {
V3502 := __e.Get(1)
_ = V3502
tmp1402 := MakeNative(func(__e *ControlFlow) {
tmp1407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V3502)
}
__typedArg0 := V3502
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp1407 {
tmp1404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V3502, MakeNumber(0))
}
__typedArg0 := V3502
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp1405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4tuple, tmp1404)
}
__typedArg0 := symshen_4tuple
__typedArg1 := tmp1404
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1405 {
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


}, 0)

tmp1408 := MakeNative(func(__e *ControlFlow) {
Z3503 := __e.Get(1)
_ = Z3503
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1402, tmp1408)
return


}, 1)

tmp1409 := Call(__e, ns2_1set, symtuple_2, tmp1401)


_ = tmp1409

tmp1410 := MakeNative(func(__e *ControlFlow) {
V3508 := __e.Get(1)
_ = V3508
V3509 := __e.Get(2)
_ = V3509
tmp1417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3508)
}
__typedArg0 := Nil
__typedArg1 := V3508
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1417 {
__e.Return(V3509)
return
} else {
tmp1415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3508)
}
__typedArg0 := V3508
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1415 {
tmp1411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3508)
}
__typedArg0 := V3508
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3508)
}
__typedArg0 := V3508
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1413 := Call(__e, PrimFunc(symappend), tmp1412, V3509)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1411, tmp1413)
}
__typedArg0 := tmp1411
__typedArg1 := tmp1413
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

tmp1418 := Call(__e, ns2_1set, symappend, tmp1410)


_ = tmp1418

tmp1419 := MakeNative(func(__e *ControlFlow) {
V3510 := __e.Get(1)
_ = V3510
V3511 := __e.Get(2)
_ = V3511
tmp1420 := Call(__e, PrimFunc(symlimit), V3511)


let__1175 := tmp1420
_ = let__1175

tmp1422 := Call(__e, PrimFunc(symvector), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(let__1175)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := let__1175
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())


let__1176 := tmp1422
_ = let__1176

tmp1423 := Call(__e, PrimFunc(symvector_1_6), let__1176, MakeNumber(1), V3510)


let__1177 := tmp1423
_ = let__1177

tmp1425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__1175, MakeNumber(0))
}
__typedArg0 := let__1175
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1425 {
__e.Return(let__1177)
return
} else {
__e.TailApply(PrimFunc(symshen_4_8v_1help), V3511, MakeNumber(1), let__1175, let__1177)
return
}


}, 2)

tmp1426 := Call(__e, ns2_1set, sym_8v, tmp1419)


_ = tmp1426

tmp1427 := MakeNative(func(__e *ControlFlow) {
V3516 := __e.Get(1)
_ = V3516
V3517 := __e.Get(2)
_ = V3517
V3518 := __e.Get(3)
_ = V3518
V3519 := __e.Get(4)
_ = V3519
tmp1433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3517, V3518)
}
__typedArg0 := V3517
__typedArg1 := V3518
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1433 {
__e.TailApply(PrimFunc(symshen_4copyfromvector), V3516, V3519, V3518, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V3518)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V3518
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


} else {
tmp1429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V3517)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V3517
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()

tmp1431 := Call(__e, PrimFunc(symshen_4copyfromvector), V3516, V3519, V3517, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V3517)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V3517
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())


__e.TailApply(PrimFunc(symshen_4_8v_1help), V3516, tmp1429, V3518, tmp1431)
return


}


}, 4)

tmp1434 := Call(__e, ns2_1set, symshen_4_8v_1help, tmp1427)


_ = tmp1434

tmp1435 := MakeNative(func(__e *ControlFlow) {
V3520 := __e.Get(1)
_ = V3520
V3521 := __e.Get(2)
_ = V3521
V3522 := __e.Get(3)
_ = V3522
V3523 := __e.Get(4)
_ = V3523
tmp1436 := MakeNative(func(__e *ControlFlow) {
tmp1437 := Call(__e, PrimFunc(sym_5_1vector), V3520, V3522)


__e.TailApply(PrimFunc(symvector_1_6), V3521, V3523, tmp1437)
return


}, 0)

tmp1438 := MakeNative(func(__e *ControlFlow) {
Z3524 := __e.Get(1)
_ = Z3524
__e.Return(V3521)
return
}, 1)

__e.TailApply(try_1catch, tmp1436, tmp1438)
return


}, 4)

tmp1439 := Call(__e, ns2_1set, symshen_4copyfromvector, tmp1435)


_ = tmp1439

tmp1440 := MakeNative(func(__e *ControlFlow) {
V3525 := __e.Get(1)
_ = V3525
tmp1441 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(sym_5_1vector), V3525, MakeNumber(1))
return
}, 0)

tmp1442 := MakeNative(func(__e *ControlFlow) {
Z3526 := __e.Get(1)
_ = Z3526
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("hdv needs a non-empty vector as an argument\n"))
}
__typedArg0 := MakeString("hdv needs a non-empty vector as an argument\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}, 1)

__e.TailApply(try_1catch, tmp1441, tmp1442)
return


}, 1)

tmp1443 := Call(__e, ns2_1set, symhdv, tmp1440)


_ = tmp1443

tmp1444 := MakeNative(func(__e *ControlFlow) {
V3527 := __e.Get(1)
_ = V3527
tmp1445 := Call(__e, PrimFunc(symlimit), V3527)


let__1178 := tmp1445
_ = let__1178

tmp1453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__1178, MakeNumber(0))
}
__typedArg0 := let__1178
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1453 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("cannot take the tail of the empty vector\n"))
}
__typedArg0 := MakeString("cannot take the tail of the empty vector\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp1451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__1178, MakeNumber(1))
}
__typedArg0 := let__1178
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1451 {
__e.TailApply(PrimFunc(symvector), MakeNumber(0))
return
} else {
tmp1447 := Call(__e, PrimFunc(symvector), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(let__1178)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := let__1178
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


let__1179 := tmp1447
_ = let__1179

tmp1449 := Call(__e, PrimFunc(symvector), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(let__1178)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := let__1178
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


__e.TailApply(PrimFunc(symshen_4tlv_1help), V3527, MakeNumber(2), let__1178, tmp1449)
return


}


}


}, 1)

tmp1454 := Call(__e, ns2_1set, symtlv, tmp1444)


_ = tmp1454

tmp1455 := MakeNative(func(__e *ControlFlow) {
V3531 := __e.Get(1)
_ = V3531
V3532 := __e.Get(2)
_ = V3532
V3533 := __e.Get(3)
_ = V3533
V3534 := __e.Get(4)
_ = V3534
tmp1461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3532, V3533)
}
__typedArg0 := V3532
__typedArg1 := V3533
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1461 {
__e.TailApply(PrimFunc(symshen_4copyfromvector), V3531, V3534, V3533, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V3533)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V3533
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())
return


} else {
tmp1457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V3532)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V3532
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()

tmp1459 := Call(__e, PrimFunc(symshen_4copyfromvector), V3531, V3534, V3532, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V3532)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V3532
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


__e.TailApply(PrimFunc(symshen_4tlv_1help), V3531, tmp1457, V3533, tmp1459)
return


}


}, 4)

tmp1462 := Call(__e, ns2_1set, symshen_4tlv_1help, tmp1455)


_ = tmp1462

tmp1463 := MakeNative(func(__e *ControlFlow) {
V3546 := __e.Get(1)
_ = V3546
V3547 := __e.Get(2)
_ = V3547
tmp1479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3547)
}
__typedArg0 := Nil
__typedArg1 := V3547
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1479 {
__e.Return(Nil)
return
} else {
tmp1477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3547)
}
__typedArg0 := V3547
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1468 Obj

if True == tmp1477 {
tmp1475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3547)
}
__typedArg0 := V3547
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp1475)
}
__typedArg0 := tmp1475
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1470 Obj

if True == tmp1476 {
tmp1472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3547)
}
__typedArg0 := V3547
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1472)
}
__typedArg0 := tmp1472
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3546, tmp1473)
}
__typedArg0 := V3546
__typedArg1 := tmp1473
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1471 Obj

if True == tmp1474 {
ifres1471 = True


} else {
ifres1471 = False


}

ifres1470 = ifres1471


} else {
ifres1470 = False


}

var ifres1469 Obj

if True == ifres1470 {
ifres1469 = True


} else {
ifres1469 = False


}

ifres1468 = ifres1469


} else {
ifres1468 = False


}

if True == ifres1468 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3547)
}
__typedArg0 := V3547
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp1466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3547)
}
__typedArg0 := V3547
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1466 {
tmp1464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3547)
}
__typedArg0 := V3547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symassoc), V3546, tmp1464)
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

tmp1480 := Call(__e, ns2_1set, symassoc, tmp1463)


_ = tmp1480

tmp1481 := MakeNative(func(__e *ControlFlow) {
V3550 := __e.Get(1)
_ = V3550
tmp1485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, V3550)
}
__typedArg0 := True
__typedArg1 := V3550
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1485 {
__e.Return(True)
return
} else {
tmp1483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, V3550)
}
__typedArg0 := False
__typedArg1 := V3550
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1483 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp1486 := Call(__e, ns2_1set, symboolean_2, tmp1481)


_ = tmp1486

tmp1487 := MakeNative(func(__e *ControlFlow) {
V3551 := __e.Get(1)
_ = V3551
tmp1492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V3551)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V3551
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1492 {
__e.Return(MakeNumber(0))
return
} else {
tmp1488 := Call(__e, PrimFunc(symstoutput))


tmp1489 := Call(__e, PrimFunc(sympr), MakeString("\n"), tmp1488)


_ = tmp1489

__e.TailApply(PrimFunc(symnl), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V3551)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V3551
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp1493 := Call(__e, ns2_1set, symnl, tmp1487)


_ = tmp1493

tmp1494 := MakeNative(func(__e *ControlFlow) {
V3558 := __e.Get(1)
_ = V3558
V3559 := __e.Get(2)
_ = V3559
tmp1505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3558)
}
__typedArg0 := Nil
__typedArg1 := V3558
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1505 {
__e.Return(Nil)
return
} else {
tmp1503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3558)
}
__typedArg0 := V3558
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1503 {
tmp1500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3558)
}
__typedArg0 := V3558
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1501 := Call(__e, PrimFunc(symelement_2), tmp1500, V3559)


if True == tmp1501 {
tmp1495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3558)
}
__typedArg0 := V3558
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symdifference), tmp1495, V3559)
return


} else {
tmp1496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3558)
}
__typedArg0 := V3558
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3558)
}
__typedArg0 := V3558
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1498 := Call(__e, PrimFunc(symdifference), tmp1497, V3559)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1496, tmp1498)
}
__typedArg0 := tmp1496
__typedArg1 := tmp1498
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

tmp1506 := Call(__e, ns2_1set, symdifference, tmp1494)


_ = tmp1506

tmp1507 := MakeNative(func(__e *ControlFlow) {
V3560 := __e.Get(1)
_ = V3560
V3561 := __e.Get(2)
_ = V3561
__e.Return(V3561)
return
}, 2)

tmp1508 := Call(__e, ns2_1set, symdo, tmp1507)


_ = tmp1508

tmp1509 := MakeNative(func(__e *ControlFlow) {
V3573 := __e.Get(1)
_ = V3573
V3574 := __e.Get(2)
_ = V3574
tmp1520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3574)
}
__typedArg0 := Nil
__typedArg1 := V3574
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1520 {
__e.Return(False)
return
} else {
tmp1518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3574)
}
__typedArg0 := V3574
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1514 Obj

if True == tmp1518 {
tmp1516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3574)
}
__typedArg0 := V3574
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3573, tmp1516)
}
__typedArg0 := V3573
__typedArg1 := tmp1516
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1515 Obj

if True == tmp1517 {
ifres1515 = True


} else {
ifres1515 = False


}

ifres1514 = ifres1515


} else {
ifres1514 = False


}

if True == ifres1514 {
__e.Return(True)
return
} else {
tmp1512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3574)
}
__typedArg0 := V3574
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1512 {
tmp1510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3574)
}
__typedArg0 := V3574
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symelement_2), V3573, tmp1510)
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

tmp1521 := Call(__e, ns2_1set, symelement_2, tmp1509)


_ = tmp1521

tmp1522 := MakeNative(func(__e *ControlFlow) {
V3577 := __e.Get(1)
_ = V3577
tmp1524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3577)
}
__typedArg0 := Nil
__typedArg1 := V3577
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1524 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp1525 := Call(__e, ns2_1set, symempty_2, tmp1522)


_ = tmp1525

tmp1526 := MakeNative(func(__e *ControlFlow) {
V3578 := __e.Get(1)
_ = V3578
V3579 := __e.Get(2)
_ = V3579
tmp1527 := Call(__e, V3578, V3579)


__e.TailApply(PrimFunc(symshen_4fix_1help), V3578, V3579, tmp1527)
return


}, 2)

tmp1528 := Call(__e, ns2_1set, symfix, tmp1526)


_ = tmp1528

tmp1529 := MakeNative(func(__e *ControlFlow) {
V3585 := __e.Get(1)
_ = V3585
V3586 := __e.Get(2)
_ = V3586
V3587 := __e.Get(3)
_ = V3587
tmp1532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3586, V3587)
}
__typedArg0 := V3586
__typedArg1 := V3587
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1532 {
__e.Return(V3587)
return
} else {
tmp1530 := Call(__e, V3585, V3587)


__e.TailApply(PrimFunc(symshen_4fix_1help), V3585, V3587, tmp1530)
return


}


}, 3)

tmp1533 := Call(__e, ns2_1set, symshen_4fix_1help, tmp1529)


_ = tmp1533

tmp1534 := MakeNative(func(__e *ControlFlow) {
V3588 := __e.Get(1)
_ = V3588
V3589 := __e.Get(2)
_ = V3589
V3590 := __e.Get(3)
_ = V3590
V3591 := __e.Get(4)
_ = V3591
tmp1535 := Call(__e, PrimFunc(symlimit), V3591)


tmp1536 := Call(__e, PrimFunc(symhash), V3588, tmp1535)


let__1180 := tmp1536
_ = let__1180

tmp1537 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(sym_5_1vector), V3591, let__1180)
return
}, 0)

tmp1538 := MakeNative(func(__e *ControlFlow) {
Z3594 := __e.Get(1)
_ = Z3594
__e.Return(Nil)
return
}, 1)

tmp1539 := Call(__e, try_1catch, tmp1537, tmp1538)


let__1181 := tmp1539
_ = let__1181

tmp1540 := Call(__e, PrimFunc(symshen_4change_1pointer_1value), V3588, V3589, V3590, let__1181)


tmp1541 := Call(__e, PrimFunc(symvector_1_6), V3591, let__1180, tmp1540)


let__1182 := tmp1541
_ = let__1182

__e.Return(V3590)
return


}, 4)

tmp1542 := Call(__e, ns2_1set, symput, tmp1534)


_ = tmp1542

tmp1543 := MakeNative(func(__e *ControlFlow) {
V3596 := __e.Get(1)
_ = V3596
V3597 := __e.Get(2)
_ = V3597
V3598 := __e.Get(3)
_ = V3598
tmp1544 := Call(__e, PrimFunc(symlimit), V3598)


tmp1545 := Call(__e, PrimFunc(symhash), V3596, tmp1544)


let__1183 := tmp1545
_ = let__1183

tmp1546 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(sym_5_1vector), V3598, let__1183)
return
}, 0)

tmp1547 := MakeNative(func(__e *ControlFlow) {
Z3601 := __e.Get(1)
_ = Z3601
__e.Return(Nil)
return
}, 1)

tmp1548 := Call(__e, try_1catch, tmp1546, tmp1547)


let__1184 := tmp1548
_ = let__1184

tmp1549 := Call(__e, PrimFunc(symshen_4remove_1pointer), V3596, V3597, let__1184)


tmp1550 := Call(__e, PrimFunc(symvector_1_6), V3598, let__1183, tmp1549)


let__1185 := tmp1550
_ = let__1185

__e.Return(V3596)
return


}, 3)

tmp1551 := Call(__e, ns2_1set, symunput, tmp1543)


_ = tmp1551

tmp1552 := MakeNative(func(__e *ControlFlow) {
V3613 := __e.Get(1)
_ = V3613
V3614 := __e.Get(2)
_ = V3614
V3615 := __e.Get(3)
_ = V3615
tmp1596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3615)
}
__typedArg0 := Nil
__typedArg1 := V3615
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1596 {
__e.Return(Nil)
return
} else {
tmp1594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3615)
}
__typedArg0 := V3615
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1559 Obj

if True == tmp1594 {
tmp1592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3615)
}
__typedArg0 := V3615
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp1592)
}
__typedArg0 := tmp1592
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1561 Obj

if True == tmp1593 {
tmp1589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3615)
}
__typedArg0 := V3615
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1589)
}
__typedArg0 := tmp1589
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp1590)
}
__typedArg0 := tmp1590
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1563 Obj

if True == tmp1591 {
tmp1585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3615)
}
__typedArg0 := V3615
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1585)
}
__typedArg0 := tmp1585
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp1586)
}
__typedArg0 := tmp1586
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp1587)
}
__typedArg0 := tmp1587
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1565 Obj

if True == tmp1588 {
tmp1580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3615)
}
__typedArg0 := V3615
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1580)
}
__typedArg0 := tmp1580
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp1581)
}
__typedArg0 := tmp1581
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp1582)
}
__typedArg0 := tmp1582
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp1583)
}
__typedArg0 := Nil
__typedArg1 := tmp1583
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1567 Obj

if True == tmp1584 {
tmp1575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3615)
}
__typedArg0 := V3615
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1575)
}
__typedArg0 := tmp1575
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp1576)
}
__typedArg0 := tmp1576
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1577)
}
__typedArg0 := tmp1577
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3614, tmp1578)
}
__typedArg0 := V3614
__typedArg1 := tmp1578
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1569 Obj

if True == tmp1579 {
tmp1571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3615)
}
__typedArg0 := V3615
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1571)
}
__typedArg0 := tmp1571
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1572)
}
__typedArg0 := tmp1572
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3613, tmp1573)
}
__typedArg0 := V3613
__typedArg1 := tmp1573
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1570 Obj

if True == tmp1574 {
ifres1570 = True


} else {
ifres1570 = False


}

ifres1569 = ifres1570


} else {
ifres1569 = False


}

var ifres1568 Obj

if True == ifres1569 {
ifres1568 = True


} else {
ifres1568 = False


}

ifres1567 = ifres1568


} else {
ifres1567 = False


}

var ifres1566 Obj

if True == ifres1567 {
ifres1566 = True


} else {
ifres1566 = False


}

ifres1565 = ifres1566


} else {
ifres1565 = False


}

var ifres1564 Obj

if True == ifres1565 {
ifres1564 = True


} else {
ifres1564 = False


}

ifres1563 = ifres1564


} else {
ifres1563 = False


}

var ifres1562 Obj

if True == ifres1563 {
ifres1562 = True


} else {
ifres1562 = False


}

ifres1561 = ifres1562


} else {
ifres1561 = False


}

var ifres1560 Obj

if True == ifres1561 {
ifres1560 = True


} else {
ifres1560 = False


}

ifres1559 = ifres1560


} else {
ifres1559 = False


}

if True == ifres1559 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3615)
}
__typedArg0 := V3615
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
} else {
tmp1557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3615)
}
__typedArg0 := V3615
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1557 {
tmp1553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3615)
}
__typedArg0 := V3615
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3615)
}
__typedArg0 := V3615
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1555 := Call(__e, PrimFunc(symshen_4remove_1pointer), V3613, V3614, tmp1554)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1553, tmp1555)
}
__typedArg0 := tmp1553
__typedArg1 := tmp1555
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.remove-pointer"))
}
__typedArg0 := MakeString("implementation error in shen.remove-pointer")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 3)

tmp1597 := Call(__e, ns2_1set, symshen_4remove_1pointer, tmp1552)


_ = tmp1597

tmp1598 := MakeNative(func(__e *ControlFlow) {
V3628 := __e.Get(1)
_ = V3628
V3629 := __e.Get(2)
_ = V3629
V3630 := __e.Get(3)
_ = V3630
V3631 := __e.Get(4)
_ = V3631
tmp1649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3631)
}
__typedArg0 := Nil
__typedArg1 := V3631
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1649 {
tmp1599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3629, Nil)
}
__typedArg0 := V3629
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3628, tmp1599)
}
__typedArg0 := V3628
__typedArg1 := tmp1599
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1600, V3630)
}
__typedArg0 := tmp1600
__typedArg1 := V3630
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1601, Nil)
}
__typedArg0 := tmp1601
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp1647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3631)
}
__typedArg0 := V3631
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1612 Obj

if True == tmp1647 {
tmp1645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3631)
}
__typedArg0 := V3631
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp1645)
}
__typedArg0 := tmp1645
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1614 Obj

if True == tmp1646 {
tmp1642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3631)
}
__typedArg0 := V3631
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1642)
}
__typedArg0 := tmp1642
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp1643)
}
__typedArg0 := tmp1643
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1616 Obj

if True == tmp1644 {
tmp1638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3631)
}
__typedArg0 := V3631
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1638)
}
__typedArg0 := tmp1638
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp1639)
}
__typedArg0 := tmp1639
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp1640)
}
__typedArg0 := tmp1640
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1618 Obj

if True == tmp1641 {
tmp1633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3631)
}
__typedArg0 := V3631
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1633)
}
__typedArg0 := tmp1633
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp1634)
}
__typedArg0 := tmp1634
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp1635)
}
__typedArg0 := tmp1635
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp1636)
}
__typedArg0 := Nil
__typedArg1 := tmp1636
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1620 Obj

if True == tmp1637 {
tmp1628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3631)
}
__typedArg0 := V3631
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1628)
}
__typedArg0 := tmp1628
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp1629)
}
__typedArg0 := tmp1629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1630)
}
__typedArg0 := tmp1630
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3629, tmp1631)
}
__typedArg0 := V3629
__typedArg1 := tmp1631
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1622 Obj

if True == tmp1632 {
tmp1624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3631)
}
__typedArg0 := V3631
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1624)
}
__typedArg0 := tmp1624
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1625)
}
__typedArg0 := tmp1625
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3628, tmp1626)
}
__typedArg0 := V3628
__typedArg1 := tmp1626
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1623 Obj

if True == tmp1627 {
ifres1623 = True


} else {
ifres1623 = False


}

ifres1622 = ifres1623


} else {
ifres1622 = False


}

var ifres1621 Obj

if True == ifres1622 {
ifres1621 = True


} else {
ifres1621 = False


}

ifres1620 = ifres1621


} else {
ifres1620 = False


}

var ifres1619 Obj

if True == ifres1620 {
ifres1619 = True


} else {
ifres1619 = False


}

ifres1618 = ifres1619


} else {
ifres1618 = False


}

var ifres1617 Obj

if True == ifres1618 {
ifres1617 = True


} else {
ifres1617 = False


}

ifres1616 = ifres1617


} else {
ifres1616 = False


}

var ifres1615 Obj

if True == ifres1616 {
ifres1615 = True


} else {
ifres1615 = False


}

ifres1614 = ifres1615


} else {
ifres1614 = False


}

var ifres1613 Obj

if True == ifres1614 {
ifres1613 = True


} else {
ifres1613 = False


}

ifres1612 = ifres1613


} else {
ifres1612 = False


}

if True == ifres1612 {
tmp1602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3631)
}
__typedArg0 := V3631
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1602)
}
__typedArg0 := tmp1602
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1603, V3630)
}
__typedArg0 := tmp1603
__typedArg1 := V3630
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3631)
}
__typedArg0 := V3631
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1604, tmp1605)
}
__typedArg0 := tmp1604
__typedArg1 := tmp1605
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp1610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3631)
}
__typedArg0 := V3631
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1610 {
tmp1606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3631)
}
__typedArg0 := V3631
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3631)
}
__typedArg0 := V3631
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1608 := Call(__e, PrimFunc(symshen_4change_1pointer_1value), V3628, V3629, V3630, tmp1607)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1606, tmp1608)
}
__typedArg0 := tmp1606
__typedArg1 := tmp1608
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.change-pointer-value"))
}
__typedArg0 := MakeString("implementation error in shen.change-pointer-value")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 4)

tmp1650 := Call(__e, ns2_1set, symshen_4change_1pointer_1value, tmp1598)


_ = tmp1650

tmp1651 := MakeNative(func(__e *ControlFlow) {
V3632 := __e.Get(1)
_ = V3632
V3633 := __e.Get(2)
_ = V3633
V3634 := __e.Get(3)
_ = V3634
tmp1652 := Call(__e, PrimFunc(symlimit), V3634)


tmp1653 := Call(__e, PrimFunc(symhash), V3632, tmp1652)


let__1186 := tmp1653
_ = let__1186

tmp1654 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(sym_5_1vector), V3634, let__1186)
return
}, 0)

tmp1655 := MakeNative(func(__e *ControlFlow) {
Z3637 := __e.Get(1)
_ = Z3637
tmp1656 := Call(__e, PrimFunc(symshen_4app), V3633, MakeString("\n"), symshen_4s)


tmp1658 := Call(__e, PrimFunc(symshen_4app), V3632, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" has no attributes: "))
__typedS1, __typedOK1 := TypedString(tmp1656)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" has no attributes: ")
__typedArg1 := tmp1656
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp1658)
}
__typedArg0 := tmp1658
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

tmp1659 := Call(__e, try_1catch, tmp1654, tmp1655)


let__1187 := tmp1659
_ = let__1187

tmp1660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3633, Nil)
}
__typedArg0 := V3633
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3632, tmp1660)
}
__typedArg0 := V3632
__typedArg1 := tmp1660
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1662 := Call(__e, PrimFunc(symassoc), tmp1661, let__1187)


let__1188 := tmp1662
_ = let__1188

tmp1668 := Call(__e, PrimFunc(symempty_2), let__1188)


if True == tmp1668 {
tmp1663 := Call(__e, PrimFunc(symshen_4app), V3632, MakeString("\n"), symshen_4s)


tmp1665 := Call(__e, PrimFunc(symshen_4app), V3633, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" not found for "))
__typedS1, __typedOK1 := TypedString(tmp1663)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" not found for ")
__typedArg1 := tmp1663
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attribute "))
__typedS1, __typedOK1 := TypedString(tmp1665)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attribute ")
__typedArg1 := tmp1665
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attribute "))
__typedS1, __typedOK1 := TypedString(tmp1665)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attribute ")
__typedArg1 := tmp1665
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

tmp1669 := Call(__e, ns2_1set, symget, tmp1651)


_ = tmp1669

tmp1670 := MakeNative(func(__e *ControlFlow) {
V3639 := __e.Get(1)
_ = V3639
V3640 := __e.Get(2)
_ = V3640
tmp1671 := Call(__e, PrimFunc(symshen_4hashkey), V3639)


tmp1672 := Call(__e, PrimFunc(symshen_4mod), tmp1671, V3640)


let__1189 := tmp1672
_ = let__1189

tmp1674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__1189, MakeNumber(0))
}
__typedArg0 := let__1189
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1674 {
__e.Return(MakeNumber(1))
return
} else {
__e.Return(let__1189)
return
}


}, 2)

tmp1675 := Call(__e, ns2_1set, symhash, tmp1670)


_ = tmp1675

tmp1676 := MakeNative(func(__e *ControlFlow) {
V3642 := __e.Get(1)
_ = V3642
tmp1677 := MakeNative(func(__e *ControlFlow) {
Z3644 := __e.Get(1)
_ = Z3644
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(Z3644)
}
__typedArg0 := Z3644
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})())
return
}, 1)

tmp1678 := Call(__e, PrimFunc(symexplode), V3642)


tmp1679 := Call(__e, PrimFunc(symmap), tmp1677, tmp1678)


let__1190 := tmp1679
_ = let__1190

__e.TailApply(PrimFunc(symshen_4prodbutzero), let__1190, MakeNumber(1))
return


}, 1)

tmp1680 := Call(__e, ns2_1set, symshen_4hashkey, tmp1676)


_ = tmp1680

tmp1681 := MakeNative(func(__e *ControlFlow) {
V3645 := __e.Get(1)
_ = V3645
V3646 := __e.Get(2)
_ = V3646
tmp1700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3645)
}
__typedArg0 := Nil
__typedArg1 := V3645
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1700 {
__e.Return(V3646)
return
} else {
tmp1698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3645)
}
__typedArg0 := V3645
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1694 Obj

if True == tmp1698 {
tmp1696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3645)
}
__typedArg0 := V3645
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), tmp1696)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp1696
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1695 Obj

if True == tmp1697 {
ifres1695 = True


} else {
ifres1695 = False


}

ifres1694 = ifres1695


} else {
ifres1694 = False


}

if True == ifres1694 {
tmp1682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3645)
}
__typedArg0 := V3645
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp1682, V3646)
return


} else {
tmp1692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3645)
}
__typedArg0 := V3645
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1692 {
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(V3646)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1e+10))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := V3646
__typedArg1 := MakeNumber(1e+10)
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
tmp1683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3645)
}
__typedArg0 := V3645
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3645)
}
__typedArg0 := V3645
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp1683, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V3646)
__typedN1, __typedOK1 := TypedFloat64(tmp1684)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V3646
__typedArg1 := tmp1684
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


} else {
tmp1686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3645)
}
__typedArg0 := V3645
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3645)
}
__typedArg0 := V3645
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp1686, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(V3646)
__typedN1, __typedOK1 := TypedFloat64(tmp1687)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := V3646
__typedArg1 := tmp1687
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})())
return


}


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.prodbutzero"))
}
__typedArg0 := MakeString("partial function shen.prodbutzero")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 2)

tmp1701 := Call(__e, ns2_1set, symshen_4prodbutzero, tmp1681)


_ = tmp1701

tmp1702 := MakeNative(func(__e *ControlFlow) {
V3647 := __e.Get(1)
_ = V3647
V3648 := __e.Get(2)
_ = V3648
tmp1703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3648, Nil)
}
__typedArg0 := V3648
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1704 := Call(__e, PrimFunc(symshen_4multiples), V3647, tmp1703)


__e.TailApply(PrimFunc(symshen_4modh), V3647, tmp1704)
return


}, 2)

tmp1705 := Call(__e, ns2_1set, symshen_4mod, tmp1702)


_ = tmp1705

tmp1706 := MakeNative(func(__e *ControlFlow) {
V3653 := __e.Get(1)
_ = V3653
V3654 := __e.Get(2)
_ = V3654
tmp1717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3654)
}
__typedArg0 := V3654
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1713 Obj

if True == tmp1717 {
tmp1715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3654)
}
__typedArg0 := V3654
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(tmp1715)
__typedN1, __typedOK1 := TypedFloat64(V3653)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := tmp1715
__typedArg1 := V3653
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})()

var ifres1714 Obj

if True == tmp1716 {
ifres1714 = True


} else {
ifres1714 = False


}

ifres1713 = ifres1714


} else {
ifres1713 = False


}

if True == ifres1713 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3654)
}
__typedArg0 := V3654
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
} else {
tmp1711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3654)
}
__typedArg0 := V3654
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1711 {
tmp1707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3654)
}
__typedArg0 := V3654
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(2))
__typedN1, __typedOK1 := TypedFloat64(tmp1707)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp1707
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})(), V3654)
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(2))
__typedN1, __typedOK1 := TypedFloat64(tmp1707)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp1707
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})()
__typedArg1 := V3654
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4multiples), V3653, tmp1709)
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

tmp1718 := Call(__e, ns2_1set, symshen_4multiples, tmp1706)


_ = tmp1718

tmp1719 := MakeNative(func(__e *ControlFlow) {
V3661 := __e.Get(1)
_ = V3661
V3662 := __e.Get(2)
_ = V3662
tmp1737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V3661)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V3661
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1737 {
__e.Return(MakeNumber(0))
return
} else {
tmp1735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3662)
}
__typedArg0 := Nil
__typedArg1 := V3662
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1735 {
__e.Return(V3661)
return
} else {
tmp1733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1729 Obj

if True == tmp1733 {
tmp1731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(tmp1731)
__typedN1, __typedOK1 := TypedFloat64(V3661)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := tmp1731
__typedArg1 := V3661
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})()

var ifres1730 Obj

if True == tmp1732 {
ifres1730 = True


} else {
ifres1730 = False


}

ifres1729 = ifres1730


} else {
ifres1729 = False


}

if True == ifres1729 {
tmp1722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1723 := Call(__e, PrimFunc(symempty_2), tmp1722)


if True == tmp1723 {
__e.Return(V3661)
return
} else {
tmp1720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4modh), V3661, tmp1720)
return


}


} else {
tmp1727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1727 {
tmp1724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4modh), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V3661)
__typedN1, __typedOK1 := TypedFloat64(tmp1724)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V3661
__typedArg1 := tmp1724
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})(), V3662)
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

tmp1738 := Call(__e, ns2_1set, symshen_4modh, tmp1719)


_ = tmp1738

tmp1739 := MakeNative(func(__e *ControlFlow) {
V3665 := __e.Get(1)
_ = V3665
tmp1746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3665)
}
__typedArg0 := Nil
__typedArg1 := V3665
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1746 {
__e.Return(MakeNumber(0))
return
} else {
tmp1744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3665)
}
__typedArg0 := V3665
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1744 {
tmp1740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3665)
}
__typedArg0 := V3665
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3665)
}
__typedArg0 := V3665
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1742 := Call(__e, PrimFunc(symsum), tmp1741)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(tmp1740)
__typedN1, __typedOK1 := TypedFloat64(tmp1742)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := tmp1740
__typedArg1 := tmp1742
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

tmp1747 := Call(__e, ns2_1set, symsum, tmp1739)


_ = tmp1747

tmp1748 := MakeNative(func(__e *ControlFlow) {
V3670 := __e.Get(1)
_ = V3670
tmp1750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3670)
}
__typedArg0 := V3670
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1750 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3670)
}
__typedArg0 := V3670
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

tmp1751 := Call(__e, ns2_1set, symhead, tmp1748)


_ = tmp1751

tmp1752 := MakeNative(func(__e *ControlFlow) {
V3675 := __e.Get(1)
_ = V3675
tmp1754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3675)
}
__typedArg0 := V3675
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1754 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3675)
}
__typedArg0 := V3675
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

tmp1755 := Call(__e, ns2_1set, symtail, tmp1752)


_ = tmp1755

tmp1756 := MakeNative(func(__e *ControlFlow) {
V3676 := __e.Get(1)
_ = V3676
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sympos) {
return PrimPos(V3676, MakeNumber(0))
}
__typedArg0 := V3676
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sympos), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp1757 := Call(__e, ns2_1set, symhdstr, tmp1756)


_ = tmp1757

tmp1758 := MakeNative(func(__e *ControlFlow) {
V3683 := __e.Get(1)
_ = V3683
V3684 := __e.Get(2)
_ = V3684
tmp1769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3683)
}
__typedArg0 := Nil
__typedArg1 := V3683
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1769 {
__e.Return(Nil)
return
} else {
tmp1767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3683)
}
__typedArg0 := V3683
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1767 {
tmp1764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3683)
}
__typedArg0 := V3683
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1765 := Call(__e, PrimFunc(symelement_2), tmp1764, V3684)


if True == tmp1765 {
tmp1759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3683)
}
__typedArg0 := V3683
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3683)
}
__typedArg0 := V3683
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1761 := Call(__e, PrimFunc(symintersection), tmp1760, V3684)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1759, tmp1761)
}
__typedArg0 := tmp1759
__typedArg1 := tmp1761
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp1762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3683)
}
__typedArg0 := V3683
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symintersection), tmp1762, V3684)
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

tmp1770 := Call(__e, ns2_1set, symintersection, tmp1758)


_ = tmp1770

tmp1771 := MakeNative(func(__e *ControlFlow) {
V3685 := __e.Get(1)
_ = V3685
__e.TailApply(PrimFunc(symshen_4reverse_1help), V3685, Nil)
return
}, 1)

tmp1772 := Call(__e, ns2_1set, symreverse, tmp1771)


_ = tmp1772

tmp1773 := MakeNative(func(__e *ControlFlow) {
V3690 := __e.Get(1)
_ = V3690
V3691 := __e.Get(2)
_ = V3691
tmp1780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3690)
}
__typedArg0 := Nil
__typedArg1 := V3690
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1780 {
__e.Return(V3691)
return
} else {
tmp1778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3690)
}
__typedArg0 := V3690
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1778 {
tmp1774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3690)
}
__typedArg0 := V3690
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3690)
}
__typedArg0 := V3690
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1775, V3691)
}
__typedArg0 := tmp1775
__typedArg1 := V3691
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4reverse_1help), tmp1774, tmp1776)
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

tmp1781 := Call(__e, ns2_1set, symshen_4reverse_1help, tmp1773)


_ = tmp1781

tmp1782 := MakeNative(func(__e *ControlFlow) {
V3696 := __e.Get(1)
_ = V3696
V3697 := __e.Get(2)
_ = V3697
tmp1793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3696)
}
__typedArg0 := Nil
__typedArg1 := V3696
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1793 {
__e.Return(V3697)
return
} else {
tmp1791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3696)
}
__typedArg0 := V3696
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1791 {
tmp1788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3696)
}
__typedArg0 := V3696
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1789 := Call(__e, PrimFunc(symelement_2), tmp1788, V3697)


if True == tmp1789 {
tmp1783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3696)
}
__typedArg0 := V3696
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symunion), tmp1783, V3697)
return


} else {
tmp1784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3696)
}
__typedArg0 := V3696
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3696)
}
__typedArg0 := V3696
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1786 := Call(__e, PrimFunc(symunion), tmp1785, V3697)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1784, tmp1786)
}
__typedArg0 := tmp1784
__typedArg1 := tmp1786
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

tmp1794 := Call(__e, ns2_1set, symunion, tmp1782)


_ = tmp1794

tmp1795 := MakeNative(func(__e *ControlFlow) {
V3698 := __e.Get(1)
_ = V3698
tmp1796 := Call(__e, PrimFunc(symshen_4proc_1nl), V3698)


tmp1797 := Call(__e, PrimFunc(symstoutput))


tmp1798 := Call(__e, PrimFunc(sympr), tmp1796, tmp1797)


let__1191 := tmp1798
_ = let__1191

tmp1799 := Call(__e, PrimFunc(symstoutput))


tmp1800 := Call(__e, PrimFunc(sympr), MakeString(" (y/n) "), tmp1799)


let__1192 := tmp1800
_ = let__1192

tmp1801 := Call(__e, PrimFunc(symstinput))


tmp1802 := Call(__e, PrimFunc(symread), tmp1801)


tmp1803 := Call(__e, PrimFunc(symshen_4app), tmp1802, MakeString(""), symshen_4s)


let__1193 := tmp1803
_ = let__1193

tmp1809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("y"), let__1193)
}
__typedArg0 := MakeString("y")
__typedArg1 := let__1193
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1809 {
__e.Return(True)
return
} else {
tmp1807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("n"), let__1193)
}
__typedArg0 := MakeString("n")
__typedArg1 := let__1193
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1807 {
__e.Return(False)
return
} else {
tmp1804 := Call(__e, PrimFunc(symstoutput))


tmp1805 := Call(__e, PrimFunc(sympr), MakeString("please answer y or n\n"), tmp1804)


_ = tmp1805

__e.TailApply(PrimFunc(symy_1or_1n_2), V3698)
return


}


}


}, 1)

tmp1810 := Call(__e, ns2_1set, symy_1or_1n_2, tmp1795)


_ = tmp1810

tmp1811 := MakeNative(func(__e *ControlFlow) {
V3702 := __e.Get(1)
_ = V3702
if True == V3702 {
__e.Return(False)
return
} else {
__e.Return(True)
return
}
}, 1)

tmp1813 := Call(__e, ns2_1set, symnot, tmp1811)


_ = tmp1813

tmp1814 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString(""))
}
__typedArg0 := MakeString("")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}, 0)

tmp1815 := Call(__e, ns2_1set, symabort, tmp1814)


_ = tmp1815

tmp1816 := MakeNative(func(__e *ControlFlow) {
V3708 := __e.Get(1)
_ = V3708
V3709 := __e.Get(2)
_ = V3709
V3710 := __e.Get(3)
_ = V3710
tmp1824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3709, V3710)
}
__typedArg0 := V3709
__typedArg1 := V3710
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1824 {
__e.Return(V3708)
return
} else {
tmp1822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3710)
}
__typedArg0 := V3710
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1822 {
tmp1817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3710)
}
__typedArg0 := V3710
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1818 := Call(__e, PrimFunc(symsubst), V3708, V3709, tmp1817)


tmp1819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3710)
}
__typedArg0 := V3710
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1820 := Call(__e, PrimFunc(symsubst), V3708, V3709, tmp1819)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1818, tmp1820)
}
__typedArg0 := tmp1818
__typedArg1 := tmp1820
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V3710)
return
}


}


}, 3)

tmp1825 := Call(__e, ns2_1set, symsubst, tmp1816)


_ = tmp1825

tmp1826 := MakeNative(func(__e *ControlFlow) {
V3711 := __e.Get(1)
_ = V3711
tmp1827 := Call(__e, PrimFunc(symshen_4app), V3711, MakeString(""), symshen_4a)


__e.TailApply(PrimFunc(symshen_4explode_1h), tmp1827)
return


}, 1)

tmp1828 := Call(__e, ns2_1set, symexplode, tmp1826)


_ = tmp1828

tmp1829 := MakeNative(func(__e *ControlFlow) {
V3714 := __e.Get(1)
_ = V3714
tmp1836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V3714)
}
__typedArg0 := MakeString("")
__typedArg1 := V3714
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1836 {
__e.Return(Nil)
return
} else {
tmp1834 := Call(__e, PrimFunc(symshen_4_7string_2), V3714)


if True == tmp1834 {
tmp1830 := Call(__e, PrimFunc(symhdstr), V3714)


tmp1832 := Call(__e, PrimFunc(symshen_4explode_1h), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3714)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3714
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1830, tmp1832)
}
__typedArg0 := tmp1830
__typedArg1 := tmp1832
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

tmp1837 := Call(__e, ns2_1set, symshen_4explode_1h, tmp1829)


_ = tmp1837

tmp1838 := MakeNative(func(__e *ControlFlow) {
V3715 := __e.Get(1)
_ = V3715
tmp1841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3715, MakeString(""))
}
__typedArg0 := V3715
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1839 Obj

if True == tmp1841 {
ifres1839 = MakeString("")


} else {
tmp1840 := Call(__e, PrimFunc(symshen_4app), V3715, MakeString("/"), symshen_4a)


ifres1839 = tmp1840


}

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dhome_1directory_d, ifres1839)
}
__typedArg0 := sym_dhome_1directory_d
__typedArg1 := ifres1839
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp1842 := Call(__e, ns2_1set, symcd, tmp1838)


_ = tmp1842

tmp1843 := MakeNative(func(__e *ControlFlow) {
V3716 := __e.Get(1)
_ = V3716
V3717 := __e.Get(2)
_ = V3717
__e.TailApply(PrimFunc(symshen_4map_1h), V3716, V3717, Nil)
return
}, 2)

tmp1844 := Call(__e, ns2_1set, symmap, tmp1843)


_ = tmp1844

tmp1845 := MakeNative(func(__e *ControlFlow) {
V3718 := __e.Get(1)
_ = V3718
V3719 := __e.Get(2)
_ = V3719
V3720 := __e.Get(3)
_ = V3720
tmp1853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3719)
}
__typedArg0 := Nil
__typedArg1 := V3719
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1853 {
__e.TailApply(PrimFunc(symreverse), V3720)
return
} else {
tmp1851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3719)
}
__typedArg0 := V3719
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1851 {
tmp1846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3719)
}
__typedArg0 := V3719
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3719)
}
__typedArg0 := V3719
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1848 := Call(__e, V3718, tmp1847)


tmp1849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1848, V3720)
}
__typedArg0 := tmp1848
__typedArg1 := V3720
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4map_1h), V3718, tmp1846, tmp1849)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.map-h"))
}
__typedArg0 := MakeString("partial function shen.map-h")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 3)

tmp1854 := Call(__e, ns2_1set, symshen_4map_1h, tmp1845)


_ = tmp1854

tmp1855 := MakeNative(func(__e *ControlFlow) {
V3721 := __e.Get(1)
_ = V3721
__e.TailApply(PrimFunc(symshen_4length_1h), V3721, MakeNumber(0))
return
}, 1)

tmp1856 := Call(__e, ns2_1set, symlength, tmp1855)


_ = tmp1856

tmp1857 := MakeNative(func(__e *ControlFlow) {
V3726 := __e.Get(1)
_ = V3726
V3727 := __e.Get(2)
_ = V3727
tmp1861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3726)
}
__typedArg0 := Nil
__typedArg1 := V3726
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1861 {
__e.Return(V3727)
return
} else {
tmp1858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3726)
}
__typedArg0 := V3726
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4length_1h), tmp1858, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V3727)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V3727
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


}


}, 2)

tmp1862 := Call(__e, ns2_1set, symshen_4length_1h, tmp1857)


_ = tmp1862

tmp1863 := MakeNative(func(__e *ControlFlow) {
V3733 := __e.Get(1)
_ = V3733
V3734 := __e.Get(2)
_ = V3734
tmp1871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3733, V3734)
}
__typedArg0 := V3733
__typedArg1 := V3734
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1871 {
__e.Return(MakeNumber(1))
return
} else {
tmp1869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3734)
}
__typedArg0 := V3734
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1869 {
tmp1864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3734)
}
__typedArg0 := V3734
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1865 := Call(__e, PrimFunc(symoccurrences), V3733, tmp1864)


tmp1866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3734)
}
__typedArg0 := V3734
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1867 := Call(__e, PrimFunc(symoccurrences), V3733, tmp1866)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(tmp1865)
__typedN1, __typedOK1 := TypedFloat64(tmp1867)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := tmp1865
__typedArg1 := tmp1867
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(MakeNumber(0))
return
}


}


}, 2)

tmp1872 := Call(__e, ns2_1set, symoccurrences, tmp1863)


_ = tmp1872

tmp1873 := MakeNative(func(__e *ControlFlow) {
V3739 := __e.Get(1)
_ = V3739
V3740 := __e.Get(2)
_ = V3740
tmp1886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(1), V3739)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := V3739
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1883 Obj

if True == tmp1886 {
tmp1885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3740)
}
__typedArg0 := V3740
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1884 Obj

if True == tmp1885 {
ifres1884 = True


} else {
ifres1884 = False


}

ifres1883 = ifres1884


} else {
ifres1883 = False


}

if True == ifres1883 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3740)
}
__typedArg0 := V3740
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp1881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3740)
}
__typedArg0 := V3740
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1881 {
tmp1874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V3739)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V3739
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})()

tmp1875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3740)
}
__typedArg0 := V3740
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symnth), tmp1874, tmp1875)
return


} else {
tmp1876 := Call(__e, PrimFunc(symshen_4app), V3740, MakeString("\n"), symshen_4a)


tmp1878 := Call(__e, PrimFunc(symshen_4app), V3739, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(", "))
__typedS1, __typedOK1 := TypedString(tmp1876)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(", ")
__typedArg1 := tmp1876
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("nth applied to "))
__typedS1, __typedOK1 := TypedString(tmp1878)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("nth applied to ")
__typedArg1 := tmp1878
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("nth applied to "))
__typedS1, __typedOK1 := TypedString(tmp1878)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("nth applied to ")
__typedArg1 := tmp1878
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}


}, 2)

tmp1887 := Call(__e, ns2_1set, symnth, tmp1873)


_ = tmp1887

tmp1888 := MakeNative(func(__e *ControlFlow) {
V3741 := __e.Get(1)
_ = V3741
tmp1894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(V3741)
}
__typedArg0 := V3741
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

if True == tmp1894 {
tmp1890 := Call(__e, PrimFunc(symshen_4abs), V3741)


let__1194 := tmp1890
_ = let__1194

tmp1891 := Call(__e, PrimFunc(symshen_4magless), let__1194, MakeNumber(1))


tmp1892 := Call(__e, PrimFunc(symshen_4integer_1test_2), let__1194, tmp1891)


if True == tmp1892 {
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

tmp1895 := Call(__e, ns2_1set, syminteger_2, tmp1888)


_ = tmp1895

tmp1896 := MakeNative(func(__e *ControlFlow) {
V3743 := __e.Get(1)
_ = V3743
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(V3743)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(0))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := V3743
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
__e.Return(V3743)
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(0))
__typedN1, __typedOK1 := TypedFloat64(V3743)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := MakeNumber(0)
__typedArg1 := V3743
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())
return
}


}, 1)

tmp1899 := Call(__e, ns2_1set, symshen_4abs, tmp1896)


_ = tmp1899

tmp1900 := MakeNative(func(__e *ControlFlow) {
V3744 := __e.Get(1)
_ = V3744
V3745 := __e.Get(2)
_ = V3745
let__1195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(V3745)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(2))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := V3745
__typedArg1 := MakeNumber(2)
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})()
_ = let__1195

if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(let__1195)
__typedN1, __typedOK1 := TypedFloat64(V3744)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := let__1195
__typedArg1 := V3744
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
__e.Return(V3745)
return
} else {
__e.TailApply(PrimFunc(symshen_4magless), V3744, let__1195)
return
}


}, 2)

tmp1904 := Call(__e, ns2_1set, symshen_4magless, tmp1900)


_ = tmp1904

tmp1905 := MakeNative(func(__e *ControlFlow) {
V3750 := __e.Get(1)
_ = V3750
V3751 := __e.Get(2)
_ = V3751
tmp1912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V3750)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V3750
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1912 {
__e.Return(True)
return
} else {
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(V3750)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := MakeNumber(1)
__typedArg1 := V3750
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
__e.Return(False)
return
} else {
let__1196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V3750)
__typedN1, __typedOK1 := TypedFloat64(V3751)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V3750
__typedArg1 := V3751
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})()
_ = let__1196

if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(0))
__typedN1, __typedOK1 := TypedFloat64(let__1196)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := MakeNumber(0)
__typedArg1 := let__1196
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V3750)
}
__typedArg0 := V3750
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})())
return
} else {
__e.TailApply(PrimFunc(symshen_4integer_1test_2), let__1196, V3751)
return
}


}


}


}, 2)

tmp1913 := Call(__e, ns2_1set, symshen_4integer_1test_2, tmp1905)


_ = tmp1913

tmp1914 := MakeNative(func(__e *ControlFlow) {
V3759 := __e.Get(1)
_ = V3759
V3760 := __e.Get(2)
_ = V3760
tmp1922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3760)
}
__typedArg0 := Nil
__typedArg1 := V3760
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1922 {
__e.Return(Nil)
return
} else {
tmp1920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3760)
}
__typedArg0 := V3760
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1920 {
tmp1915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3760)
}
__typedArg0 := V3760
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1916 := Call(__e, V3759, tmp1915)


tmp1917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3760)
}
__typedArg0 := V3760
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1918 := Call(__e, PrimFunc(symmapcan), V3759, tmp1917)


__e.TailApply(PrimFunc(symappend), tmp1916, tmp1918)
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

tmp1923 := Call(__e, ns2_1set, symmapcan, tmp1914)


_ = tmp1923

tmp1924 := MakeNative(func(__e *ControlFlow) {
V3766 := __e.Get(1)
_ = V3766
V3767 := __e.Get(2)
_ = V3767
tmp1926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3766, V3767)
}
__typedArg0 := V3766
__typedArg1 := V3767
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1926 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp1927 := Call(__e, ns2_1set, sym_a_a, tmp1924)


_ = tmp1927

tmp1928 := MakeNative(func(__e *ControlFlow) {
V3768 := __e.Get(1)
_ = V3768
tmp1936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V3768)
}
__typedArg0 := V3768
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp1936 {
tmp1930 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(V3768)
}
__typedArg0 := V3768
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1931 := MakeNative(func(__e *ControlFlow) {
Z3770 := __e.Get(1)
_ = Z3770
__e.Return(symshen_4this_1symbol_1is_1unbound)
return
}, 1)

tmp1932 := Call(__e, try_1catch, tmp1930, tmp1931)


let__1197 := tmp1932
_ = let__1197

tmp1934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__1197, symshen_4this_1symbol_1is_1unbound)
}
__typedArg0 := let__1197
__typedArg1 := symshen_4this_1symbol_1is_1unbound
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1933 Obj

if True == tmp1934 {
ifres1933 = False


} else {
ifres1933 = True


}

if True == ifres1933 {
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

tmp1937 := Call(__e, ns2_1set, symbound_2, tmp1928)


_ = tmp1937

tmp1938 := MakeNative(func(__e *ControlFlow) {
V3771 := __e.Get(1)
_ = V3771
tmp1944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V3771)
}
__typedArg0 := MakeString("")
__typedArg1 := V3771
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1944 {
__e.Return(Nil)
return
} else {
tmp1939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sympos) {
return PrimPos(V3771, MakeNumber(0))
}
__typedArg0 := V3771
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sympos), __typedArg0, __typedArg1)
})()

tmp1940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp1939)
}
__typedArg0 := tmp1939
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp1942 := Call(__e, PrimFunc(symshen_4string_1_6bytes), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3771)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3771
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1940, tmp1942)
}
__typedArg0 := tmp1940
__typedArg1 := tmp1942
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp1945 := Call(__e, ns2_1set, symshen_4string_1_6bytes, tmp1938)


_ = tmp1945

tmp1946 := MakeNative(func(__e *ControlFlow) {
V3772 := __e.Get(1)
_ = V3772
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5) {
__typedN0, __typedOK0 := TypedFloat64(V3772)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(0))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5) {
return TypedMaterializeBoolean((__typedN0 < __typedN1))
}}
__typedArg0 := V3772
__typedArg1 := MakeNumber(0)
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
tmp1948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V3772)
}
__typedArg0 := V3772
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})()

if True == tmp1948 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dmaxinferences_d, V3772)
}
__typedArg0 := symshen_4_dmaxinferences_d
__typedArg1 := V3772
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

tmp1951 := Call(__e, ns2_1set, symmaxinferences, tmp1946)


_ = tmp1951

tmp1952 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dinfs_d)
}
__typedArg0 := symshen_4_dinfs_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1953 := Call(__e, ns2_1set, syminferences, tmp1952)


_ = tmp1953

tmp1954 := MakeNative(func(__e *ControlFlow) {
V3773 := __e.Get(1)
_ = V3773
__e.Return(V3773)
return
}, 1)

tmp1955 := Call(__e, ns2_1set, symprotect, tmp1954)


_ = tmp1955

tmp1956 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dstoutput_d)
}
__typedArg0 := sym_dstoutput_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1957 := Call(__e, ns2_1set, symstoutput, tmp1956)


_ = tmp1957

tmp1958 := MakeNative(func(__e *ControlFlow) {
V3774 := __e.Get(1)
_ = V3774
tmp1959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(V3774)
}
__typedArg0 := V3774
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

let__1198 := tmp1959
_ = let__1198

tmp1963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(let__1198)
}
__typedArg0 := let__1198
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp1963 {
__e.Return(let__1198)
return
} else {
tmp1960 := Call(__e, PrimFunc(symshen_4app), V3774, MakeString(" to a symbol"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("cannot intern "))
__typedS1, __typedOK1 := TypedString(tmp1960)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("cannot intern ")
__typedArg1 := tmp1960
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("cannot intern "))
__typedS1, __typedOK1 := TypedString(tmp1960)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("cannot intern ")
__typedArg1 := tmp1960
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 1)

tmp1964 := Call(__e, ns2_1set, symstring_1_6symbol, tmp1958)


_ = tmp1964

tmp1965 := MakeNative(func(__e *ControlFlow) {
V3778 := __e.Get(1)
_ = V3778
tmp1969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V3778)
}
__typedArg0 := sym_7
__typedArg1 := V3778
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1969 {
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
tmp1967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V3778)
}
__typedArg0 := sym_1
__typedArg1 := V3778
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1967 {
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

tmp1970 := Call(__e, ns2_1set, symoptimise, tmp1965)


_ = tmp1970

tmp1971 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dos_d)
}
__typedArg0 := sym_dos_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1972 := Call(__e, ns2_1set, symos, tmp1971)


_ = tmp1972

tmp1973 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dlanguage_d)
}
__typedArg0 := sym_dlanguage_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1974 := Call(__e, ns2_1set, symlanguage, tmp1973)


_ = tmp1974

tmp1975 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dversion_d)
}
__typedArg0 := sym_dversion_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1976 := Call(__e, ns2_1set, symversion, tmp1975)


_ = tmp1976

tmp1977 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dport_d)
}
__typedArg0 := sym_dport_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1978 := Call(__e, ns2_1set, symport, tmp1977)


_ = tmp1978

tmp1979 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dporters_d)
}
__typedArg0 := sym_dporters_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1980 := Call(__e, ns2_1set, symporters, tmp1979)


_ = tmp1980

tmp1981 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dimplementation_d)
}
__typedArg0 := sym_dimplementation_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1982 := Call(__e, ns2_1set, symimplementation, tmp1981)


_ = tmp1982

tmp1983 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_drelease_d)
}
__typedArg0 := sym_drelease_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1984 := Call(__e, ns2_1set, symrelease, tmp1983)


_ = tmp1984

tmp1985 := MakeNative(func(__e *ControlFlow) {
V3779 := __e.Get(1)
_ = V3779
tmp1990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnull, V3779)
}
__typedArg0 := symnull
__typedArg1 := V3779
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1990 {
__e.Return(True)
return
} else {
tmp1986 := MakeNative(func(__e *ControlFlow) {
tmp1987 := Call(__e, PrimFunc(symexternal), V3779)


_ = tmp1987

__e.Return(True)
return


}, 0)

tmp1988 := MakeNative(func(__e *ControlFlow) {
Z3780 := __e.Get(1)
_ = Z3780
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1986, tmp1988)
return


}


}, 1)

tmp1991 := Call(__e, ns2_1set, sympackage_2, tmp1985)


_ = tmp1991

tmp1992 := MakeNative(func(__e *ControlFlow) {
__e.Return(symshen_4fail_b)
return
}, 0)

tmp1993 := Call(__e, ns2_1set, symfail, tmp1992)


_ = tmp1993

tmp1994 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_duserdefs_d)
}
__typedArg0 := symshen_4_duserdefs_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1995 := Call(__e, ns2_1set, symuserdefs, tmp1994)


_ = tmp1995

tmp1996 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_doptimise_d)
}
__typedArg0 := symshen_4_doptimise_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1997 := Call(__e, ns2_1set, symoptimise_2, tmp1996)


_ = tmp1997

tmp1998 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dhush_d)
}
__typedArg0 := sym_dhush_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1999 := Call(__e, ns2_1set, symhush_2, tmp1998)


_ = tmp1999

tmp2000 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dshen_1type_1theory_1enabled_2_d)
}
__typedArg0 := symshen_4_dshen_1type_1theory_1enabled_2_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp2001 := Call(__e, ns2_1set, symsystem_1S_2, tmp2000)


_ = tmp2001

tmp2002 := MakeNative(func(__e *ControlFlow) {
V3783 := __e.Get(1)
_ = V3783
tmp2006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V3783)
}
__typedArg0 := sym_7
__typedArg1 := V3783
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2006 {
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
tmp2004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V3783)
}
__typedArg0 := sym_1
__typedArg1 := V3783
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2004 {
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

tmp2007 := Call(__e, ns2_1set, symenable_1type_1theory, tmp2002)


_ = tmp2007

tmp2008 := MakeNative(func(__e *ControlFlow) {
V3786 := __e.Get(1)
_ = V3786
tmp2012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V3786)
}
__typedArg0 := sym_7
__typedArg1 := V3786
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2012 {
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
tmp2010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V3786)
}
__typedArg0 := sym_1
__typedArg1 := V3786
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2010 {
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

tmp2013 := Call(__e, ns2_1set, symshen_4hush, tmp2008)


_ = tmp2013

tmp2014 := MakeNative(func(__e *ControlFlow) {
V3789 := __e.Get(1)
_ = V3789
tmp2018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V3789)
}
__typedArg0 := sym_7
__typedArg1 := V3789
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2018 {
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
tmp2016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V3789)
}
__typedArg0 := sym_1
__typedArg1 := V3789
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2016 {
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

tmp2019 := Call(__e, ns2_1set, symtc, tmp2014)


_ = tmp2019

tmp2020 := MakeNative(func(__e *ControlFlow) {
V3790 := __e.Get(1)
_ = V3790
tmp2021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2022 := Call(__e, PrimFunc(symshen_4unassoc), V3790, tmp2021)


tmp2023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp2022)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp2022
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2023

__e.Return(V3790)
return


}, 1)

tmp2024 := Call(__e, ns2_1set, symdestroy, tmp2020)


_ = tmp2024

tmp2025 := MakeNative(func(__e *ControlFlow) {
V3800 := __e.Get(1)
_ = V3800
V3801 := __e.Get(2)
_ = V3801
tmp2043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3801)
}
__typedArg0 := Nil
__typedArg1 := V3801
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2043 {
__e.Return(Nil)
return
} else {
tmp2041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3801)
}
__typedArg0 := V3801
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2032 Obj

if True == tmp2041 {
tmp2039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3801)
}
__typedArg0 := V3801
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2039)
}
__typedArg0 := tmp2039
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2034 Obj

if True == tmp2040 {
tmp2036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3801)
}
__typedArg0 := V3801
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2036)
}
__typedArg0 := tmp2036
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3800, tmp2037)
}
__typedArg0 := V3800
__typedArg1 := tmp2037
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2035 Obj

if True == tmp2038 {
ifres2035 = True


} else {
ifres2035 = False


}

ifres2034 = ifres2035


} else {
ifres2034 = False


}

var ifres2033 Obj

if True == ifres2034 {
ifres2033 = True


} else {
ifres2033 = False


}

ifres2032 = ifres2033


} else {
ifres2032 = False


}

if True == ifres2032 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3801)
}
__typedArg0 := V3801
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
} else {
tmp2030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3801)
}
__typedArg0 := V3801
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2030 {
tmp2026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3801)
}
__typedArg0 := V3801
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3801)
}
__typedArg0 := V3801
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2028 := Call(__e, PrimFunc(symshen_4unassoc), V3800, tmp2027)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2026, tmp2028)
}
__typedArg0 := tmp2026
__typedArg1 := tmp2028
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

tmp2044 := Call(__e, ns2_1set, symshen_4unassoc, tmp2025)


_ = tmp2044

tmp2045 := MakeNative(func(__e *ControlFlow) {
V3802 := __e.Get(1)
_ = V3802
tmp2049 := Call(__e, PrimFunc(sympackage_2), V3802)


if True == tmp2049 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dpackage_d, V3802)
}
__typedArg0 := symshen_4_dpackage_d
__typedArg1 := V3802
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
tmp2046 := Call(__e, PrimFunc(symshen_4app), V3802, MakeString(" does not exist\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("package "))
__typedS1, __typedOK1 := TypedString(tmp2046)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("package ")
__typedArg1 := tmp2046
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("package "))
__typedS1, __typedOK1 := TypedString(tmp2046)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("package ")
__typedArg1 := tmp2046
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 1)

tmp2050 := Call(__e, ns2_1set, symin_1package, tmp2045)


_ = tmp2050

tmp2051 := MakeNative(func(__e *ControlFlow) {
V3803 := __e.Get(1)
_ = V3803
V3804 := __e.Get(2)
_ = V3804
tmp2052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symopen) {
return PrimOpenStream(V3803, symout)
}
__typedArg0 := V3803
__typedArg1 := symout
return Call(__e, PrimFunc(symopen), __typedArg0, __typedArg1)
})()

let__1199 := tmp2052
_ = let__1199

tmp2055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V3804)
}
__typedArg0 := V3804
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres2053 Obj

if True == tmp2055 {
ifres2053 = V3804


} else {
tmp2054 := Call(__e, PrimFunc(symshen_4app), V3804, MakeString(""), symshen_4s)


ifres2053 = tmp2054


}

let__1200 := ifres2053
_ = let__1200

tmp2056 := Call(__e, PrimFunc(sympr), let__1200, let__1199)


let__1201 := tmp2056
_ = let__1201

tmp2057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symclose) {
return PrimCloseStream(let__1199)
}
__typedArg0 := let__1199
return Call(__e, PrimFunc(symclose), __typedArg0)
})()

let__1202 := tmp2057
_ = let__1202

__e.Return(V3804)
return


}, 2)

tmp2058 := Call(__e, ns2_1set, symwrite_1to_1file, tmp2051)


_ = tmp2058

tmp2059 := MakeNative(func(__e *ControlFlow) {
tmp2060 := Call(__e, PrimFunc(symgensym), symshen_4t)


__e.TailApply(PrimFunc(symshen_4freshterm), tmp2060)
return


}, 0)

tmp2061 := Call(__e, ns2_1set, symfresh, tmp2059)


_ = tmp2061

tmp2062 := MakeNative(func(__e *ControlFlow) {
V3809 := __e.Get(1)
_ = V3809
V3810 := __e.Get(2)
_ = V3810
tmp2063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2064 := Call(__e, PrimFunc(symput), V3809, symarity, V3810, tmp2063)


let__1203 := tmp2064
_ = let__1203

tmp2065 := Call(__e, PrimFunc(symshen_4lambda_1entry), V3809)


let__1204 := tmp2065
_ = let__1204

tmp2066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dlambdatable_d)
}
__typedArg0 := symshen_4_dlambdatable_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__1204, tmp2066)
}
__typedArg0 := let__1204
__typedArg1 := tmp2066
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dlambdatable_d, tmp2067)
}
__typedArg0 := symshen_4_dlambdatable_d
__typedArg1 := tmp2067
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__1205 := tmp2068
_ = let__1205

__e.Return(V3809)
return


}, 2)

tmp2069 := Call(__e, ns2_1set, symupdate_1lambda_1table, tmp2062)


_ = tmp2069

tmp2070 := MakeNative(func(__e *ControlFlow) {
V3816 := __e.Get(1)
_ = V3816
V3817 := __e.Get(2)
_ = V3817
tmp2094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V3817)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V3817
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2094 {
tmp2071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspecial_d)
}
__typedArg0 := symshen_4_dspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2072 := Call(__e, PrimFunc(symremove), V3816, tmp2071)


tmp2073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspecial_d, tmp2072)
}
__typedArg0 := symshen_4_dspecial_d
__typedArg1 := tmp2072
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2073

tmp2074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dextraspecial_d)
}
__typedArg0 := symshen_4_dextraspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2075 := Call(__e, PrimFunc(symremove), V3816, tmp2074)


tmp2076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dextraspecial_d, tmp2075)
}
__typedArg0 := symshen_4_dextraspecial_d
__typedArg1 := tmp2075
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2076

__e.Return(V3816)
return


} else {
tmp2092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(1), V3817)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := V3817
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2092 {
tmp2077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspecial_d)
}
__typedArg0 := symshen_4_dspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2078 := Call(__e, PrimFunc(symadjoin), V3816, tmp2077)


tmp2079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspecial_d, tmp2078)
}
__typedArg0 := symshen_4_dspecial_d
__typedArg1 := tmp2078
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2079

tmp2080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dextraspecial_d)
}
__typedArg0 := symshen_4_dextraspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2081 := Call(__e, PrimFunc(symremove), V3816, tmp2080)


tmp2082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dextraspecial_d, tmp2081)
}
__typedArg0 := symshen_4_dextraspecial_d
__typedArg1 := tmp2081
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2082

__e.Return(V3816)
return


} else {
tmp2090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(2), V3817)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := V3817
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2090 {
tmp2083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspecial_d)
}
__typedArg0 := symshen_4_dspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2084 := Call(__e, PrimFunc(symremove), V3816, tmp2083)


tmp2085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspecial_d, tmp2084)
}
__typedArg0 := symshen_4_dspecial_d
__typedArg1 := tmp2084
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2085

tmp2086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dextraspecial_d)
}
__typedArg0 := symshen_4_dextraspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2087 := Call(__e, PrimFunc(symadjoin), V3816, tmp2086)


tmp2088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dextraspecial_d, tmp2087)
}
__typedArg0 := symshen_4_dextraspecial_d
__typedArg1 := tmp2087
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2088

__e.Return(V3816)
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

tmp2095 := Call(__e, ns2_1set, symspecialise, tmp2070)


_ = tmp2095

tmp2096 := MakeNative(func(__e *ControlFlow) {
V3818 := __e.Get(1)
_ = V3818
tmp2097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dabsolute_d)
}
__typedArg0 := sym_dabsolute_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3818, tmp2097)
}
__typedArg0 := V3818
__typedArg1 := tmp2097
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dabsolute_d, tmp2098)
}
__typedArg0 := sym_dabsolute_d
__typedArg1 := tmp2098
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp2099 := Call(__e, ns2_1set, symabsolute, tmp2096)


_ = tmp2099

tmp2100 := MakeNative(func(__e *ControlFlow) {
V3819 := __e.Get(1)
_ = V3819
tmp2101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dabsolute_d)
}
__typedArg0 := sym_dabsolute_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2102 := Call(__e, PrimFunc(symremove), V3819, tmp2101)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dabsolute_d, tmp2102)
}
__typedArg0 := sym_dabsolute_d
__typedArg1 := tmp2102
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 1)

__e.TailApply(ns2_1set, symunabsolute, tmp2100)
return




}, 0)

