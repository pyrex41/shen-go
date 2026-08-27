package main

import . "github.com/pyrex41/shen-go/kl"

var SysMain = MakeNative(func(__e *ControlFlow) {
tmp1164 := MakeNative(func(__e *ControlFlow) {
V3444 := __e.Get(1)
_ = V3444
__e.TailApply(V3444)
return
}, 1)

tmp1165 := Call(__e, ns2_1set, symthaw, tmp1164)


_ = tmp1165

tmp1166 := MakeNative(func(__e *ControlFlow) {
V3445 := __e.Get(1)
_ = V3445
tmp1167 := Call(__e, PrimFunc(symmacroexpand), V3445)


tmp1168 := Call(__e, PrimFunc(symshen_4find_1types), V3445)


tmp1169 := Call(__e, PrimFunc(symshen_4process_1applications), tmp1167, tmp1168)


tmp1170 := Call(__e, PrimFunc(symshen_4shen_1_6kl), tmp1169)


__e.TailApply(PrimFunc(symeval_1kl), tmp1170)
return


}, 1)

tmp1171 := Call(__e, ns2_1set, symeval, tmp1166)


_ = tmp1171

tmp1172 := MakeNative(func(__e *ControlFlow) {
V3446 := __e.Get(1)
_ = V3446
tmp1179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnull, V3446)
}
__typedArg0 := symnull
__typedArg1 := V3446
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1179 {
__e.Return(Nil)
return
} else {
tmp1173 := MakeNative(func(__e *ControlFlow) {
tmp1174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V3446, symshen_4external_1symbols, tmp1174)
return


}, 0)

tmp1175 := MakeNative(func(__e *ControlFlow) {
Z3447 := __e.Get(1)
_ = Z3447
tmp1176 := Call(__e, PrimFunc(symshen_4app), V3446, MakeString(" does not exist.\n;"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("package "))
__typedS1, __typedOK1 := TypedString(tmp1176)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("package ")
__typedArg1 := tmp1176
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("package "))
__typedS1, __typedOK1 := TypedString(tmp1176)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("package ")
__typedArg1 := tmp1176
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

__e.TailApply(try_1catch, tmp1173, tmp1175)
return


}


}, 1)

tmp1180 := Call(__e, ns2_1set, symexternal, tmp1172)


_ = tmp1180

tmp1181 := MakeNative(func(__e *ControlFlow) {
V3448 := __e.Get(1)
_ = V3448
tmp1188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnull, V3448)
}
__typedArg0 := symnull
__typedArg1 := V3448
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1188 {
__e.Return(Nil)
return
} else {
tmp1182 := MakeNative(func(__e *ControlFlow) {
tmp1183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V3448, symshen_4internal_1symbols, tmp1183)
return


}, 0)

tmp1184 := MakeNative(func(__e *ControlFlow) {
Z3449 := __e.Get(1)
_ = Z3449
tmp1185 := Call(__e, PrimFunc(symshen_4app), V3448, MakeString(" does not exist.\n;"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("package "))
__typedS1, __typedOK1 := TypedString(tmp1185)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("package ")
__typedArg1 := tmp1185
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("package "))
__typedS1, __typedOK1 := TypedString(tmp1185)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("package ")
__typedArg1 := tmp1185
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

__e.TailApply(try_1catch, tmp1182, tmp1184)
return


}


}, 1)

tmp1189 := Call(__e, ns2_1set, syminternal, tmp1181)


_ = tmp1189

tmp1190 := MakeNative(func(__e *ControlFlow) {
V3450 := __e.Get(1)
_ = V3450
V3451 := __e.Get(2)
_ = V3451
tmp1192 := Call(__e, V3450, V3451)


if True == tmp1192 {
__e.TailApply(PrimFunc(symfail))
return
} else {
__e.Return(V3451)
return
}


}, 2)

tmp1193 := Call(__e, ns2_1set, symfail_1if, tmp1190)


_ = tmp1193

tmp1194 := MakeNative(func(__e *ControlFlow) {
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

tmp1195 := Call(__e, ns2_1set, sym_8s, tmp1194)


_ = tmp1195

tmp1196 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtc_d)
}
__typedArg0 := symshen_4_dtc_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1197 := Call(__e, ns2_1set, symtc_2, tmp1196)


_ = tmp1197

tmp1198 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_doccurs_d)
}
__typedArg0 := symshen_4_doccurs_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1199 := Call(__e, ns2_1set, symoccurs_2, tmp1198)


_ = tmp1199

tmp1200 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dfactorise_2_d)
}
__typedArg0 := symshen_4_dfactorise_2_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1201 := Call(__e, ns2_1set, symfactorise_2, tmp1200)


_ = tmp1201

tmp1202 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtracking_d)
}
__typedArg0 := symshen_4_dtracking_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1203 := Call(__e, ns2_1set, symtracked, tmp1202)


_ = tmp1203

tmp1204 := MakeNative(func(__e *ControlFlow) {
V3454 := __e.Get(1)
_ = V3454
tmp1205 := MakeNative(func(__e *ControlFlow) {
tmp1206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V3454, symshen_4source, tmp1206)
return


}, 0)

tmp1207 := MakeNative(func(__e *ControlFlow) {
Z3455 := __e.Get(1)
_ = Z3455
tmp1208 := Call(__e, PrimFunc(symshen_4app), V3454, MakeString(" not found.\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp1208)
}
__typedArg0 := tmp1208
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

__e.TailApply(try_1catch, tmp1205, tmp1207)
return


}, 1)

tmp1209 := Call(__e, ns2_1set, symps, tmp1204)


_ = tmp1209

tmp1210 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dstinput_d)
}
__typedArg0 := sym_dstinput_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1211 := Call(__e, ns2_1set, symstinput, tmp1210)


_ = tmp1211

tmp1212 := MakeNative(func(__e *ControlFlow) {
V3456 := __e.Get(1)
_ = V3456
tmp1213 := MakeNative(func(__e *ControlFlow) {
W3457 := __e.Get(1)
_ = W3457
tmp1214 := MakeNative(func(__e *ControlFlow) {
W3458 := __e.Get(1)
_ = W3458
tmp1215 := MakeNative(func(__e *ControlFlow) {
W3459 := __e.Get(1)
_ = W3459
__e.Return(W3459)
return
}, 1)

tmp1219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3456, MakeNumber(0))
}
__typedArg0 := V3456
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1216 Obj

if True == tmp1219 {
ifres1216 = W3458


} else {
tmp1217 := Call(__e, PrimFunc(symfail))


tmp1218 := Call(__e, PrimFunc(symshen_4fillvector), W3458, MakeNumber(1), V3456, tmp1217)


ifres1216 = tmp1218


}

__e.TailApply(tmp1215, ifres1216)
return


}, 1)

tmp1220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(W3457, MakeNumber(0), V3456)
}
__typedArg0 := W3457
__typedArg1 := MakeNumber(0)
__typedArg2 := V3456
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

__e.TailApply(tmp1214, tmp1220)
return


}, 1)

tmp1222 := (func() Obj {
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

__e.TailApply(tmp1213, tmp1222)
return


}, 1)

tmp1223 := Call(__e, ns2_1set, symvector, tmp1212)


_ = tmp1223

tmp1224 := MakeNative(func(__e *ControlFlow) {
V3461 := __e.Get(1)
_ = V3461
V3462 := __e.Get(2)
_ = V3462
V3463 := __e.Get(3)
_ = V3463
V3464 := __e.Get(4)
_ = V3464
tmp1228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3462, V3463)
}
__typedArg0 := V3462
__typedArg1 := V3463
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1228 {
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
tmp1225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V3461, V3462, V3464)
}
__typedArg0 := V3461
__typedArg1 := V3462
__typedArg2 := V3464
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

__e.TailApply(PrimFunc(symshen_4fillvector), tmp1225, (func() Obj {
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

tmp1229 := Call(__e, ns2_1set, symshen_4fillvector, tmp1224)


_ = tmp1229

tmp1230 := MakeNative(func(__e *ControlFlow) {
V3465 := __e.Get(1)
_ = V3465
tmp1237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V3465)
}
__typedArg0 := V3465
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp1237 {
tmp1232 := MakeNative(func(__e *ControlFlow) {
tmp1233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V3465, MakeNumber(0))
}
__typedArg0 := V3465
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6_a) {
__typedN0, __typedOK0 := TypedFloat64(tmp1233)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(0))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6_a) {
return TypedMaterializeBoolean((__typedN0 >= __typedN1))
}}
__typedArg0 := tmp1233
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_6_a), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp1234 := MakeNative(func(__e *ControlFlow) {
Z3466 := __e.Get(1)
_ = Z3466
__e.Return(False)
return
}, 1)

tmp1235 := Call(__e, try_1catch, tmp1232, tmp1234)


if True == tmp1235 {
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

tmp1238 := Call(__e, ns2_1set, symvector_2, tmp1230)


_ = tmp1238

tmp1239 := MakeNative(func(__e *ControlFlow) {
V3467 := __e.Get(1)
_ = V3467
V3468 := __e.Get(2)
_ = V3468
V3469 := __e.Get(3)
_ = V3469
tmp1241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3468, MakeNumber(0))
}
__typedArg0 := V3468
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1241 {
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

tmp1242 := Call(__e, ns2_1set, symvector_1_6, tmp1239)


_ = tmp1242

tmp1243 := MakeNative(func(__e *ControlFlow) {
V3470 := __e.Get(1)
_ = V3470
V3471 := __e.Get(2)
_ = V3471
tmp1250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3471, MakeNumber(0))
}
__typedArg0 := V3471
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1250 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("cannot access 0th element of a vector\n"))
}
__typedArg0 := MakeString("cannot access 0th element of a vector\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp1244 := MakeNative(func(__e *ControlFlow) {
W3472 := __e.Get(1)
_ = W3472
tmp1246 := Call(__e, PrimFunc(symfail))


tmp1247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W3472, tmp1246)
}
__typedArg0 := W3472
__typedArg1 := tmp1246
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1247 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("vector element not found\n"))
}
__typedArg0 := MakeString("vector element not found\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
__e.Return(W3472)
return
}


}, 1)

tmp1248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V3470, V3471)
}
__typedArg0 := V3470
__typedArg1 := V3471
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp1244, tmp1248)
return


}


}, 2)

tmp1251 := Call(__e, ns2_1set, sym_5_1vector, tmp1243)


_ = tmp1251

tmp1252 := MakeNative(func(__e *ControlFlow) {
V3473 := __e.Get(1)
_ = V3473
tmp1256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V3473)
}
__typedArg0 := V3473
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})()

if True == tmp1256 {
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

tmp1257 := Call(__e, ns2_1set, symshen_4posint_2, tmp1252)


_ = tmp1257

tmp1258 := MakeNative(func(__e *ControlFlow) {
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

tmp1259 := Call(__e, ns2_1set, symlimit, tmp1258)


_ = tmp1259

tmp1260 := MakeNative(func(__e *ControlFlow) {
V3475 := __e.Get(1)
_ = V3475
tmp1291 := Call(__e, PrimFunc(symboolean_2), V3475)


var ifres1276 Obj

if True == tmp1291 {
ifres1276 = True


} else {
tmp1290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(V3475)
}
__typedArg0 := V3475
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

var ifres1278 Obj

if True == tmp1290 {
ifres1278 = True


} else {
tmp1289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V3475)
}
__typedArg0 := V3475
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres1280 Obj

if True == tmp1289 {
ifres1280 = True


} else {
tmp1288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3475)
}
__typedArg0 := V3475
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1282 Obj

if True == tmp1288 {
ifres1282 = True


} else {
tmp1287 := Call(__e, PrimFunc(symempty_2), V3475)


var ifres1284 Obj

if True == tmp1287 {
ifres1284 = True


} else {
tmp1286 := Call(__e, PrimFunc(symvector_2), V3475)


var ifres1285 Obj

if True == tmp1286 {
ifres1285 = True


} else {
ifres1285 = False


}

ifres1284 = ifres1285


}

var ifres1283 Obj

if True == ifres1284 {
ifres1283 = True


} else {
ifres1283 = False


}

ifres1282 = ifres1283


}

var ifres1281 Obj

if True == ifres1282 {
ifres1281 = True


} else {
ifres1281 = False


}

ifres1280 = ifres1281


}

var ifres1279 Obj

if True == ifres1280 {
ifres1279 = True


} else {
ifres1279 = False


}

ifres1278 = ifres1279


}

var ifres1277 Obj

if True == ifres1278 {
ifres1277 = True


} else {
ifres1277 = False


}

ifres1276 = ifres1277


}

if True == ifres1276 {
__e.Return(False)
return
} else {
tmp1266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp1267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp1268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(","))
}
__typedArg0 := MakeString(",")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp1269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1268, Nil)
}
__typedArg0 := tmp1268
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1267, tmp1269)
}
__typedArg0 := tmp1267
__typedArg1 := tmp1269
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1266, tmp1270)
}
__typedArg0 := tmp1266
__typedArg1 := tmp1270
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, tmp1271)
}
__typedArg0 := sym_j
__typedArg1 := tmp1271
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, tmp1272)
}
__typedArg0 := sym_i
__typedArg1 := tmp1272
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1274 := Call(__e, PrimFunc(symelement_2), V3475, tmp1273)


if True == tmp1274 {
__e.Return(True)
return
} else {
tmp1261 := MakeNative(func(__e *ControlFlow) {
tmp1262 := MakeNative(func(__e *ControlFlow) {
W3476 := __e.Get(1)
_ = W3476
__e.TailApply(PrimFunc(symshen_4analyse_1symbol_2), W3476)
return
}, 1)

tmp1263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3475)
}
__typedArg0 := V3475
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

__e.TailApply(tmp1262, tmp1263)
return


}, 0)

tmp1264 := MakeNative(func(__e *ControlFlow) {
Z3477 := __e.Get(1)
_ = Z3477
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1261, tmp1264)
return


}


}


}, 1)

tmp1292 := Call(__e, ns2_1set, symsymbol_2, tmp1260)


_ = tmp1292

tmp1293 := MakeNative(func(__e *ControlFlow) {
V3480 := __e.Get(1)
_ = V3480
tmp1302 := Call(__e, PrimFunc(symshen_4_7string_2), V3480)


if True == tmp1302 {
tmp1298 := Call(__e, PrimFunc(symhdstr), V3480)


tmp1299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp1298)
}
__typedArg0 := tmp1298
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp1300 := Call(__e, PrimFunc(symshen_4alpha_2), tmp1299)


if True == tmp1300 {
tmp1296 := Call(__e, PrimFunc(symshen_4alphanums_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3480)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3480
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


if True == tmp1296 {
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

tmp1303 := Call(__e, ns2_1set, symshen_4analyse_1symbol_2, tmp1293)


_ = tmp1303

tmp1304 := MakeNative(func(__e *ControlFlow) {
V3483 := __e.Get(1)
_ = V3483
tmp1319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V3483)
}
__typedArg0 := MakeString("")
__typedArg1 := V3483
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1319 {
__e.Return(True)
return
} else {
tmp1317 := Call(__e, PrimFunc(symshen_4_7string_2), V3483)


if True == tmp1317 {
tmp1305 := MakeNative(func(__e *ControlFlow) {
W3484 := __e.Get(1)
_ = W3484
tmp1313 := Call(__e, PrimFunc(symshen_4alpha_2), W3484)


var ifres1310 Obj

if True == tmp1313 {
ifres1310 = True


} else {
tmp1312 := Call(__e, PrimFunc(symshen_4digit_2), W3484)


var ifres1311 Obj

if True == tmp1312 {
ifres1311 = True


} else {
ifres1311 = False


}

ifres1310 = ifres1311


}

if True == ifres1310 {
tmp1308 := Call(__e, PrimFunc(symshen_4alphanums_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3483)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3483
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


if True == tmp1308 {
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

tmp1314 := Call(__e, PrimFunc(symhdstr), V3483)


tmp1315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp1314)
}
__typedArg0 := tmp1314
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

__e.TailApply(tmp1305, tmp1315)
return


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

tmp1320 := Call(__e, ns2_1set, symshen_4alphanums_2, tmp1304)


_ = tmp1320

tmp1321 := MakeNative(func(__e *ControlFlow) {
V3485 := __e.Get(1)
_ = V3485
tmp1333 := Call(__e, PrimFunc(symboolean_2), V3485)


var ifres1327 Obj

if True == tmp1333 {
ifres1327 = True


} else {
tmp1332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(V3485)
}
__typedArg0 := V3485
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

var ifres1329 Obj

if True == tmp1332 {
ifres1329 = True


} else {
tmp1331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V3485)
}
__typedArg0 := V3485
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres1330 Obj

if True == tmp1331 {
ifres1330 = True


} else {
ifres1330 = False


}

ifres1329 = ifres1330


}

var ifres1328 Obj

if True == ifres1329 {
ifres1328 = True


} else {
ifres1328 = False


}

ifres1327 = ifres1328


}

if True == ifres1327 {
__e.Return(False)
return
} else {
tmp1322 := MakeNative(func(__e *ControlFlow) {
tmp1323 := MakeNative(func(__e *ControlFlow) {
W3486 := __e.Get(1)
_ = W3486
__e.TailApply(PrimFunc(symshen_4analyse_1variable_2), W3486)
return
}, 1)

tmp1324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3485)
}
__typedArg0 := V3485
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

__e.TailApply(tmp1323, tmp1324)
return


}, 0)

tmp1325 := MakeNative(func(__e *ControlFlow) {
Z3487 := __e.Get(1)
_ = Z3487
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1322, tmp1325)
return


}


}, 1)

tmp1334 := Call(__e, ns2_1set, symvariable_2, tmp1321)


_ = tmp1334

tmp1335 := MakeNative(func(__e *ControlFlow) {
V3490 := __e.Get(1)
_ = V3490
tmp1344 := Call(__e, PrimFunc(symshen_4_7string_2), V3490)


if True == tmp1344 {
tmp1340 := Call(__e, PrimFunc(symhdstr), V3490)


tmp1341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp1340)
}
__typedArg0 := tmp1340
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp1342 := Call(__e, PrimFunc(symshen_4uppercase_2), tmp1341)


if True == tmp1342 {
tmp1338 := Call(__e, PrimFunc(symshen_4alphanums_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3490)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3490
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


if True == tmp1338 {
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

tmp1345 := Call(__e, ns2_1set, symshen_4analyse_1variable_2, tmp1335)


_ = tmp1345

tmp1346 := MakeNative(func(__e *ControlFlow) {
V3491 := __e.Get(1)
_ = V3491
tmp1347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dgensym_d)
}
__typedArg0 := symshen_4_dgensym_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp1349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dgensym_d, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp1347)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp1347
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := symshen_4_dgensym_d
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp1347)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp1347
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symconcat), V3491, tmp1349)
return


}, 1)

tmp1350 := Call(__e, ns2_1set, symgensym, tmp1346)


_ = tmp1350

tmp1351 := MakeNative(func(__e *ControlFlow) {
V3492 := __e.Get(1)
_ = V3492
V3493 := __e.Get(2)
_ = V3493
tmp1352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3492)
}
__typedArg0 := V3492
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp1353 := (func() Obj {
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
__typedS0, __typedOK0 := TypedString(tmp1352)
__typedS1, __typedOK1 := TypedString(tmp1353)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp1352
__typedArg1 := tmp1353
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp1352)
__typedS1, __typedOK1 := TypedString(tmp1353)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp1352
__typedArg1 := tmp1353
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symintern), __typedArg0)
})())
return


}, 2)

tmp1355 := Call(__e, ns2_1set, symconcat, tmp1351)


_ = tmp1355

tmp1356 := MakeNative(func(__e *ControlFlow) {
V3494 := __e.Get(1)
_ = V3494
V3495 := __e.Get(2)
_ = V3495
tmp1357 := MakeNative(func(__e *ControlFlow) {
W3496 := __e.Get(1)
_ = W3496
tmp1358 := MakeNative(func(__e *ControlFlow) {
W3497 := __e.Get(1)
_ = W3497
tmp1359 := MakeNative(func(__e *ControlFlow) {
W3498 := __e.Get(1)
_ = W3498
tmp1360 := MakeNative(func(__e *ControlFlow) {
W3499 := __e.Get(1)
_ = W3499
__e.Return(W3496)
return
}, 1)

tmp1361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(W3496, MakeNumber(2), V3495)
}
__typedArg0 := W3496
__typedArg1 := MakeNumber(2)
__typedArg2 := V3495
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

__e.TailApply(tmp1360, tmp1361)
return


}, 1)

tmp1362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(W3496, MakeNumber(1), V3494)
}
__typedArg0 := W3496
__typedArg1 := MakeNumber(1)
__typedArg2 := V3494
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

__e.TailApply(tmp1359, tmp1362)
return


}, 1)

tmp1363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(W3496, MakeNumber(0), symshen_4tuple)
}
__typedArg0 := W3496
__typedArg1 := MakeNumber(0)
__typedArg2 := symshen_4tuple
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

__e.TailApply(tmp1358, tmp1363)
return


}, 1)

tmp1364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(MakeNumber(3))
}
__typedArg0 := MakeNumber(3)
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

__e.TailApply(tmp1357, tmp1364)
return


}, 2)

tmp1365 := Call(__e, ns2_1set, sym_8p, tmp1356)


_ = tmp1365

tmp1366 := MakeNative(func(__e *ControlFlow) {
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

tmp1367 := Call(__e, ns2_1set, symfst, tmp1366)


_ = tmp1367

tmp1368 := MakeNative(func(__e *ControlFlow) {
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

tmp1369 := Call(__e, ns2_1set, symsnd, tmp1368)


_ = tmp1369

tmp1370 := MakeNative(func(__e *ControlFlow) {
V3502 := __e.Get(1)
_ = V3502
tmp1371 := MakeNative(func(__e *ControlFlow) {
tmp1376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V3502)
}
__typedArg0 := V3502
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp1376 {
tmp1373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V3502, MakeNumber(0))
}
__typedArg0 := V3502
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp1374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4tuple, tmp1373)
}
__typedArg0 := symshen_4tuple
__typedArg1 := tmp1373
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1374 {
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

tmp1377 := MakeNative(func(__e *ControlFlow) {
Z3503 := __e.Get(1)
_ = Z3503
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1371, tmp1377)
return


}, 1)

tmp1378 := Call(__e, ns2_1set, symtuple_2, tmp1370)


_ = tmp1378

tmp1379 := MakeNative(func(__e *ControlFlow) {
V3508 := __e.Get(1)
_ = V3508
V3509 := __e.Get(2)
_ = V3509
tmp1386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3508)
}
__typedArg0 := Nil
__typedArg1 := V3508
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1386 {
__e.Return(V3509)
return
} else {
tmp1384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3508)
}
__typedArg0 := V3508
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1384 {
tmp1380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3508)
}
__typedArg0 := V3508
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3508)
}
__typedArg0 := V3508
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1382 := Call(__e, PrimFunc(symappend), tmp1381, V3509)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1380, tmp1382)
}
__typedArg0 := tmp1380
__typedArg1 := tmp1382
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

tmp1387 := Call(__e, ns2_1set, symappend, tmp1379)


_ = tmp1387

tmp1388 := MakeNative(func(__e *ControlFlow) {
V3510 := __e.Get(1)
_ = V3510
V3511 := __e.Get(2)
_ = V3511
tmp1389 := MakeNative(func(__e *ControlFlow) {
W3512 := __e.Get(1)
_ = W3512
tmp1390 := MakeNative(func(__e *ControlFlow) {
W3513 := __e.Get(1)
_ = W3513
tmp1391 := MakeNative(func(__e *ControlFlow) {
W3514 := __e.Get(1)
_ = W3514
tmp1393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W3512, MakeNumber(0))
}
__typedArg0 := W3512
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1393 {
__e.Return(W3514)
return
} else {
__e.TailApply(PrimFunc(symshen_4_8v_1help), V3511, MakeNumber(1), W3512, W3514)
return
}


}, 1)

tmp1394 := Call(__e, PrimFunc(symvector_1_6), W3513, MakeNumber(1), V3510)


__e.TailApply(tmp1391, tmp1394)
return


}, 1)

tmp1396 := Call(__e, PrimFunc(symvector), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(W3512)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := W3512
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())


__e.TailApply(tmp1390, tmp1396)
return


}, 1)

tmp1397 := Call(__e, PrimFunc(symlimit), V3511)


__e.TailApply(tmp1389, tmp1397)
return


}, 2)

tmp1398 := Call(__e, ns2_1set, sym_8v, tmp1388)


_ = tmp1398

tmp1399 := MakeNative(func(__e *ControlFlow) {
V3516 := __e.Get(1)
_ = V3516
V3517 := __e.Get(2)
_ = V3517
V3518 := __e.Get(3)
_ = V3518
V3519 := __e.Get(4)
_ = V3519
tmp1405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3517, V3518)
}
__typedArg0 := V3517
__typedArg1 := V3518
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1405 {
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
tmp1401 := (func() Obj {
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

tmp1403 := Call(__e, PrimFunc(symshen_4copyfromvector), V3516, V3519, V3517, (func() Obj {
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


__e.TailApply(PrimFunc(symshen_4_8v_1help), V3516, tmp1401, V3518, tmp1403)
return


}


}, 4)

tmp1406 := Call(__e, ns2_1set, symshen_4_8v_1help, tmp1399)


_ = tmp1406

tmp1407 := MakeNative(func(__e *ControlFlow) {
V3520 := __e.Get(1)
_ = V3520
V3521 := __e.Get(2)
_ = V3521
V3522 := __e.Get(3)
_ = V3522
V3523 := __e.Get(4)
_ = V3523
tmp1408 := MakeNative(func(__e *ControlFlow) {
tmp1409 := Call(__e, PrimFunc(sym_5_1vector), V3520, V3522)


__e.TailApply(PrimFunc(symvector_1_6), V3521, V3523, tmp1409)
return


}, 0)

tmp1410 := MakeNative(func(__e *ControlFlow) {
Z3524 := __e.Get(1)
_ = Z3524
__e.Return(V3521)
return
}, 1)

__e.TailApply(try_1catch, tmp1408, tmp1410)
return


}, 4)

tmp1411 := Call(__e, ns2_1set, symshen_4copyfromvector, tmp1407)


_ = tmp1411

tmp1412 := MakeNative(func(__e *ControlFlow) {
V3525 := __e.Get(1)
_ = V3525
tmp1413 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(sym_5_1vector), V3525, MakeNumber(1))
return
}, 0)

tmp1414 := MakeNative(func(__e *ControlFlow) {
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

__e.TailApply(try_1catch, tmp1413, tmp1414)
return


}, 1)

tmp1415 := Call(__e, ns2_1set, symhdv, tmp1412)


_ = tmp1415

tmp1416 := MakeNative(func(__e *ControlFlow) {
V3527 := __e.Get(1)
_ = V3527
tmp1417 := MakeNative(func(__e *ControlFlow) {
W3528 := __e.Get(1)
_ = W3528
tmp1426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W3528, MakeNumber(0))
}
__typedArg0 := W3528
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1426 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("cannot take the tail of the empty vector\n"))
}
__typedArg0 := MakeString("cannot take the tail of the empty vector\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp1424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W3528, MakeNumber(1))
}
__typedArg0 := W3528
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1424 {
__e.TailApply(PrimFunc(symvector), MakeNumber(0))
return
} else {
tmp1418 := MakeNative(func(__e *ControlFlow) {
W3529 := __e.Get(1)
_ = W3529
tmp1420 := Call(__e, PrimFunc(symvector), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(W3528)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := W3528
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


__e.TailApply(PrimFunc(symshen_4tlv_1help), V3527, MakeNumber(2), W3528, tmp1420)
return


}, 1)

tmp1422 := Call(__e, PrimFunc(symvector), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(W3528)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := W3528
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


__e.TailApply(tmp1418, tmp1422)
return


}


}


}, 1)

tmp1427 := Call(__e, PrimFunc(symlimit), V3527)


__e.TailApply(tmp1417, tmp1427)
return


}, 1)

tmp1428 := Call(__e, ns2_1set, symtlv, tmp1416)


_ = tmp1428

tmp1429 := MakeNative(func(__e *ControlFlow) {
V3531 := __e.Get(1)
_ = V3531
V3532 := __e.Get(2)
_ = V3532
V3533 := __e.Get(3)
_ = V3533
V3534 := __e.Get(4)
_ = V3534
tmp1435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3532, V3533)
}
__typedArg0 := V3532
__typedArg1 := V3533
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1435 {
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
tmp1431 := (func() Obj {
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

tmp1433 := Call(__e, PrimFunc(symshen_4copyfromvector), V3531, V3534, V3532, (func() Obj {
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


__e.TailApply(PrimFunc(symshen_4tlv_1help), V3531, tmp1431, V3533, tmp1433)
return


}


}, 4)

tmp1436 := Call(__e, ns2_1set, symshen_4tlv_1help, tmp1429)


_ = tmp1436

tmp1437 := MakeNative(func(__e *ControlFlow) {
V3546 := __e.Get(1)
_ = V3546
V3547 := __e.Get(2)
_ = V3547
tmp1453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3547)
}
__typedArg0 := Nil
__typedArg1 := V3547
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1453 {
__e.Return(Nil)
return
} else {
tmp1451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3547)
}
__typedArg0 := V3547
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1442 Obj

if True == tmp1451 {
tmp1449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3547)
}
__typedArg0 := V3547
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp1449)
}
__typedArg0 := tmp1449
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1444 Obj

if True == tmp1450 {
tmp1446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3547)
}
__typedArg0 := V3547
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1446)
}
__typedArg0 := tmp1446
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3546, tmp1447)
}
__typedArg0 := V3546
__typedArg1 := tmp1447
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1445 Obj

if True == tmp1448 {
ifres1445 = True


} else {
ifres1445 = False


}

ifres1444 = ifres1445


} else {
ifres1444 = False


}

var ifres1443 Obj

if True == ifres1444 {
ifres1443 = True


} else {
ifres1443 = False


}

ifres1442 = ifres1443


} else {
ifres1442 = False


}

if True == ifres1442 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3547)
}
__typedArg0 := V3547
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp1440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3547)
}
__typedArg0 := V3547
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1440 {
tmp1438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3547)
}
__typedArg0 := V3547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symassoc), V3546, tmp1438)
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

tmp1454 := Call(__e, ns2_1set, symassoc, tmp1437)


_ = tmp1454

tmp1455 := MakeNative(func(__e *ControlFlow) {
V3550 := __e.Get(1)
_ = V3550
tmp1459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, V3550)
}
__typedArg0 := True
__typedArg1 := V3550
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1459 {
__e.Return(True)
return
} else {
tmp1457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, V3550)
}
__typedArg0 := False
__typedArg1 := V3550
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1457 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp1460 := Call(__e, ns2_1set, symboolean_2, tmp1455)


_ = tmp1460

tmp1461 := MakeNative(func(__e *ControlFlow) {
V3551 := __e.Get(1)
_ = V3551
tmp1466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V3551)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V3551
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1466 {
__e.Return(MakeNumber(0))
return
} else {
tmp1462 := Call(__e, PrimFunc(symstoutput))


tmp1463 := Call(__e, PrimFunc(sympr), MakeString("\n"), tmp1462)


_ = tmp1463

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

tmp1467 := Call(__e, ns2_1set, symnl, tmp1461)


_ = tmp1467

tmp1468 := MakeNative(func(__e *ControlFlow) {
V3558 := __e.Get(1)
_ = V3558
V3559 := __e.Get(2)
_ = V3559
tmp1479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3558)
}
__typedArg0 := Nil
__typedArg1 := V3558
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1479 {
__e.Return(Nil)
return
} else {
tmp1477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3558)
}
__typedArg0 := V3558
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1477 {
tmp1474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3558)
}
__typedArg0 := V3558
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1475 := Call(__e, PrimFunc(symelement_2), tmp1474, V3559)


if True == tmp1475 {
tmp1469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3558)
}
__typedArg0 := V3558
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symdifference), tmp1469, V3559)
return


} else {
tmp1470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3558)
}
__typedArg0 := V3558
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3558)
}
__typedArg0 := V3558
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1472 := Call(__e, PrimFunc(symdifference), tmp1471, V3559)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1470, tmp1472)
}
__typedArg0 := tmp1470
__typedArg1 := tmp1472
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

tmp1480 := Call(__e, ns2_1set, symdifference, tmp1468)


_ = tmp1480

tmp1481 := MakeNative(func(__e *ControlFlow) {
V3560 := __e.Get(1)
_ = V3560
V3561 := __e.Get(2)
_ = V3561
__e.Return(V3561)
return
}, 2)

tmp1482 := Call(__e, ns2_1set, symdo, tmp1481)


_ = tmp1482

tmp1483 := MakeNative(func(__e *ControlFlow) {
V3573 := __e.Get(1)
_ = V3573
V3574 := __e.Get(2)
_ = V3574
tmp1494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3574)
}
__typedArg0 := Nil
__typedArg1 := V3574
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1494 {
__e.Return(False)
return
} else {
tmp1492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3574)
}
__typedArg0 := V3574
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1488 Obj

if True == tmp1492 {
tmp1490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3574)
}
__typedArg0 := V3574
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3573, tmp1490)
}
__typedArg0 := V3573
__typedArg1 := tmp1490
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1489 Obj

if True == tmp1491 {
ifres1489 = True


} else {
ifres1489 = False


}

ifres1488 = ifres1489


} else {
ifres1488 = False


}

if True == ifres1488 {
__e.Return(True)
return
} else {
tmp1486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3574)
}
__typedArg0 := V3574
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1486 {
tmp1484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3574)
}
__typedArg0 := V3574
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symelement_2), V3573, tmp1484)
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

tmp1495 := Call(__e, ns2_1set, symelement_2, tmp1483)


_ = tmp1495

tmp1496 := MakeNative(func(__e *ControlFlow) {
V3577 := __e.Get(1)
_ = V3577
tmp1498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3577)
}
__typedArg0 := Nil
__typedArg1 := V3577
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1498 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp1499 := Call(__e, ns2_1set, symempty_2, tmp1496)


_ = tmp1499

tmp1500 := MakeNative(func(__e *ControlFlow) {
V3578 := __e.Get(1)
_ = V3578
V3579 := __e.Get(2)
_ = V3579
tmp1501 := Call(__e, V3578, V3579)


__e.TailApply(PrimFunc(symshen_4fix_1help), V3578, V3579, tmp1501)
return


}, 2)

tmp1502 := Call(__e, ns2_1set, symfix, tmp1500)


_ = tmp1502

tmp1503 := MakeNative(func(__e *ControlFlow) {
V3585 := __e.Get(1)
_ = V3585
V3586 := __e.Get(2)
_ = V3586
V3587 := __e.Get(3)
_ = V3587
tmp1506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3586, V3587)
}
__typedArg0 := V3586
__typedArg1 := V3587
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1506 {
__e.Return(V3587)
return
} else {
tmp1504 := Call(__e, V3585, V3587)


__e.TailApply(PrimFunc(symshen_4fix_1help), V3585, V3587, tmp1504)
return


}


}, 3)

tmp1507 := Call(__e, ns2_1set, symshen_4fix_1help, tmp1503)


_ = tmp1507

tmp1508 := MakeNative(func(__e *ControlFlow) {
V3588 := __e.Get(1)
_ = V3588
V3589 := __e.Get(2)
_ = V3589
V3590 := __e.Get(3)
_ = V3590
V3591 := __e.Get(4)
_ = V3591
tmp1509 := MakeNative(func(__e *ControlFlow) {
W3592 := __e.Get(1)
_ = W3592
tmp1510 := MakeNative(func(__e *ControlFlow) {
W3593 := __e.Get(1)
_ = W3593
tmp1511 := MakeNative(func(__e *ControlFlow) {
W3595 := __e.Get(1)
_ = W3595
__e.Return(V3590)
return
}, 1)

tmp1512 := Call(__e, PrimFunc(symshen_4change_1pointer_1value), V3588, V3589, V3590, W3593)


tmp1513 := Call(__e, PrimFunc(symvector_1_6), V3591, W3592, tmp1512)


__e.TailApply(tmp1511, tmp1513)
return


}, 1)

tmp1514 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(sym_5_1vector), V3591, W3592)
return
}, 0)

tmp1515 := MakeNative(func(__e *ControlFlow) {
Z3594 := __e.Get(1)
_ = Z3594
__e.Return(Nil)
return
}, 1)

tmp1516 := Call(__e, try_1catch, tmp1514, tmp1515)


__e.TailApply(tmp1510, tmp1516)
return


}, 1)

tmp1517 := Call(__e, PrimFunc(symlimit), V3591)


tmp1518 := Call(__e, PrimFunc(symhash), V3588, tmp1517)


__e.TailApply(tmp1509, tmp1518)
return


}, 4)

tmp1519 := Call(__e, ns2_1set, symput, tmp1508)


_ = tmp1519

tmp1520 := MakeNative(func(__e *ControlFlow) {
V3596 := __e.Get(1)
_ = V3596
V3597 := __e.Get(2)
_ = V3597
V3598 := __e.Get(3)
_ = V3598
tmp1521 := MakeNative(func(__e *ControlFlow) {
W3599 := __e.Get(1)
_ = W3599
tmp1522 := MakeNative(func(__e *ControlFlow) {
W3600 := __e.Get(1)
_ = W3600
tmp1523 := MakeNative(func(__e *ControlFlow) {
W3602 := __e.Get(1)
_ = W3602
__e.Return(V3596)
return
}, 1)

tmp1524 := Call(__e, PrimFunc(symshen_4remove_1pointer), V3596, V3597, W3600)


tmp1525 := Call(__e, PrimFunc(symvector_1_6), V3598, W3599, tmp1524)


__e.TailApply(tmp1523, tmp1525)
return


}, 1)

tmp1526 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(sym_5_1vector), V3598, W3599)
return
}, 0)

tmp1527 := MakeNative(func(__e *ControlFlow) {
Z3601 := __e.Get(1)
_ = Z3601
__e.Return(Nil)
return
}, 1)

tmp1528 := Call(__e, try_1catch, tmp1526, tmp1527)


__e.TailApply(tmp1522, tmp1528)
return


}, 1)

tmp1529 := Call(__e, PrimFunc(symlimit), V3598)


tmp1530 := Call(__e, PrimFunc(symhash), V3596, tmp1529)


__e.TailApply(tmp1521, tmp1530)
return


}, 3)

tmp1531 := Call(__e, ns2_1set, symunput, tmp1520)


_ = tmp1531

tmp1532 := MakeNative(func(__e *ControlFlow) {
V3613 := __e.Get(1)
_ = V3613
V3614 := __e.Get(2)
_ = V3614
V3615 := __e.Get(3)
_ = V3615
tmp1576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3615)
}
__typedArg0 := Nil
__typedArg1 := V3615
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1576 {
__e.Return(Nil)
return
} else {
tmp1574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3615)
}
__typedArg0 := V3615
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1539 Obj

if True == tmp1574 {
tmp1572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3615)
}
__typedArg0 := V3615
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp1572)
}
__typedArg0 := tmp1572
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1541 Obj

if True == tmp1573 {
tmp1569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3615)
}
__typedArg0 := V3615
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1569)
}
__typedArg0 := tmp1569
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp1570)
}
__typedArg0 := tmp1570
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1543 Obj

if True == tmp1571 {
tmp1565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3615)
}
__typedArg0 := V3615
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1565)
}
__typedArg0 := tmp1565
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp1566)
}
__typedArg0 := tmp1566
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp1567)
}
__typedArg0 := tmp1567
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1545 Obj

if True == tmp1568 {
tmp1560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3615)
}
__typedArg0 := V3615
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1560)
}
__typedArg0 := tmp1560
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp1561)
}
__typedArg0 := tmp1561
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp1562)
}
__typedArg0 := tmp1562
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp1563)
}
__typedArg0 := Nil
__typedArg1 := tmp1563
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1547 Obj

if True == tmp1564 {
tmp1555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3615)
}
__typedArg0 := V3615
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1555)
}
__typedArg0 := tmp1555
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp1556)
}
__typedArg0 := tmp1556
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1557)
}
__typedArg0 := tmp1557
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3614, tmp1558)
}
__typedArg0 := V3614
__typedArg1 := tmp1558
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1549 Obj

if True == tmp1559 {
tmp1551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3615)
}
__typedArg0 := V3615
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1551)
}
__typedArg0 := tmp1551
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1552)
}
__typedArg0 := tmp1552
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3613, tmp1553)
}
__typedArg0 := V3613
__typedArg1 := tmp1553
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1550 Obj

if True == tmp1554 {
ifres1550 = True


} else {
ifres1550 = False


}

ifres1549 = ifres1550


} else {
ifres1549 = False


}

var ifres1548 Obj

if True == ifres1549 {
ifres1548 = True


} else {
ifres1548 = False


}

ifres1547 = ifres1548


} else {
ifres1547 = False


}

var ifres1546 Obj

if True == ifres1547 {
ifres1546 = True


} else {
ifres1546 = False


}

ifres1545 = ifres1546


} else {
ifres1545 = False


}

var ifres1544 Obj

if True == ifres1545 {
ifres1544 = True


} else {
ifres1544 = False


}

ifres1543 = ifres1544


} else {
ifres1543 = False


}

var ifres1542 Obj

if True == ifres1543 {
ifres1542 = True


} else {
ifres1542 = False


}

ifres1541 = ifres1542


} else {
ifres1541 = False


}

var ifres1540 Obj

if True == ifres1541 {
ifres1540 = True


} else {
ifres1540 = False


}

ifres1539 = ifres1540


} else {
ifres1539 = False


}

if True == ifres1539 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3615)
}
__typedArg0 := V3615
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
} else {
tmp1537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3615)
}
__typedArg0 := V3615
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1537 {
tmp1533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3615)
}
__typedArg0 := V3615
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3615)
}
__typedArg0 := V3615
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1535 := Call(__e, PrimFunc(symshen_4remove_1pointer), V3613, V3614, tmp1534)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1533, tmp1535)
}
__typedArg0 := tmp1533
__typedArg1 := tmp1535
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

tmp1577 := Call(__e, ns2_1set, symshen_4remove_1pointer, tmp1532)


_ = tmp1577

tmp1578 := MakeNative(func(__e *ControlFlow) {
V3628 := __e.Get(1)
_ = V3628
V3629 := __e.Get(2)
_ = V3629
V3630 := __e.Get(3)
_ = V3630
V3631 := __e.Get(4)
_ = V3631
tmp1629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3631)
}
__typedArg0 := Nil
__typedArg1 := V3631
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1629 {
tmp1579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3629, Nil)
}
__typedArg0 := V3629
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3628, tmp1579)
}
__typedArg0 := V3628
__typedArg1 := tmp1579
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1580, V3630)
}
__typedArg0 := tmp1580
__typedArg1 := V3630
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1581, Nil)
}
__typedArg0 := tmp1581
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp1627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3631)
}
__typedArg0 := V3631
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1592 Obj

if True == tmp1627 {
tmp1625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3631)
}
__typedArg0 := V3631
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp1625)
}
__typedArg0 := tmp1625
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1594 Obj

if True == tmp1626 {
tmp1622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3631)
}
__typedArg0 := V3631
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1622)
}
__typedArg0 := tmp1622
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp1623)
}
__typedArg0 := tmp1623
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1596 Obj

if True == tmp1624 {
tmp1618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3631)
}
__typedArg0 := V3631
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1618)
}
__typedArg0 := tmp1618
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp1619)
}
__typedArg0 := tmp1619
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp1620)
}
__typedArg0 := tmp1620
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1598 Obj

if True == tmp1621 {
tmp1613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3631)
}
__typedArg0 := V3631
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1613)
}
__typedArg0 := tmp1613
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp1614)
}
__typedArg0 := tmp1614
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp1615)
}
__typedArg0 := tmp1615
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp1616)
}
__typedArg0 := Nil
__typedArg1 := tmp1616
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1600 Obj

if True == tmp1617 {
tmp1608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3631)
}
__typedArg0 := V3631
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1608)
}
__typedArg0 := tmp1608
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp1609)
}
__typedArg0 := tmp1609
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1610)
}
__typedArg0 := tmp1610
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3629, tmp1611)
}
__typedArg0 := V3629
__typedArg1 := tmp1611
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1602 Obj

if True == tmp1612 {
tmp1604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3631)
}
__typedArg0 := V3631
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1604)
}
__typedArg0 := tmp1604
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1605)
}
__typedArg0 := tmp1605
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3628, tmp1606)
}
__typedArg0 := V3628
__typedArg1 := tmp1606
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1603 Obj

if True == tmp1607 {
ifres1603 = True


} else {
ifres1603 = False


}

ifres1602 = ifres1603


} else {
ifres1602 = False


}

var ifres1601 Obj

if True == ifres1602 {
ifres1601 = True


} else {
ifres1601 = False


}

ifres1600 = ifres1601


} else {
ifres1600 = False


}

var ifres1599 Obj

if True == ifres1600 {
ifres1599 = True


} else {
ifres1599 = False


}

ifres1598 = ifres1599


} else {
ifres1598 = False


}

var ifres1597 Obj

if True == ifres1598 {
ifres1597 = True


} else {
ifres1597 = False


}

ifres1596 = ifres1597


} else {
ifres1596 = False


}

var ifres1595 Obj

if True == ifres1596 {
ifres1595 = True


} else {
ifres1595 = False


}

ifres1594 = ifres1595


} else {
ifres1594 = False


}

var ifres1593 Obj

if True == ifres1594 {
ifres1593 = True


} else {
ifres1593 = False


}

ifres1592 = ifres1593


} else {
ifres1592 = False


}

if True == ifres1592 {
tmp1582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3631)
}
__typedArg0 := V3631
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp1582)
}
__typedArg0 := tmp1582
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1583, V3630)
}
__typedArg0 := tmp1583
__typedArg1 := V3630
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3631)
}
__typedArg0 := V3631
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1584, tmp1585)
}
__typedArg0 := tmp1584
__typedArg1 := tmp1585
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp1590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3631)
}
__typedArg0 := V3631
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1590 {
tmp1586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3631)
}
__typedArg0 := V3631
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3631)
}
__typedArg0 := V3631
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1588 := Call(__e, PrimFunc(symshen_4change_1pointer_1value), V3628, V3629, V3630, tmp1587)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1586, tmp1588)
}
__typedArg0 := tmp1586
__typedArg1 := tmp1588
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

tmp1630 := Call(__e, ns2_1set, symshen_4change_1pointer_1value, tmp1578)


_ = tmp1630

tmp1631 := MakeNative(func(__e *ControlFlow) {
V3632 := __e.Get(1)
_ = V3632
V3633 := __e.Get(2)
_ = V3633
V3634 := __e.Get(3)
_ = V3634
tmp1632 := MakeNative(func(__e *ControlFlow) {
W3635 := __e.Get(1)
_ = W3635
tmp1633 := MakeNative(func(__e *ControlFlow) {
W3636 := __e.Get(1)
_ = W3636
tmp1634 := MakeNative(func(__e *ControlFlow) {
W3638 := __e.Get(1)
_ = W3638
tmp1640 := Call(__e, PrimFunc(symempty_2), W3638)


if True == tmp1640 {
tmp1635 := Call(__e, PrimFunc(symshen_4app), V3632, MakeString("\n"), symshen_4s)


tmp1637 := Call(__e, PrimFunc(symshen_4app), V3633, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" not found for "))
__typedS1, __typedOK1 := TypedString(tmp1635)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" not found for ")
__typedArg1 := tmp1635
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attribute "))
__typedS1, __typedOK1 := TypedString(tmp1637)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attribute ")
__typedArg1 := tmp1637
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attribute "))
__typedS1, __typedOK1 := TypedString(tmp1637)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attribute ")
__typedArg1 := tmp1637
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W3638)
}
__typedArg0 := W3638
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
}


}, 1)

tmp1641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3633, Nil)
}
__typedArg0 := V3633
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3632, tmp1641)
}
__typedArg0 := V3632
__typedArg1 := tmp1641
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1643 := Call(__e, PrimFunc(symassoc), tmp1642, W3636)


__e.TailApply(tmp1634, tmp1643)
return


}, 1)

tmp1644 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(sym_5_1vector), V3634, W3635)
return
}, 0)

tmp1645 := MakeNative(func(__e *ControlFlow) {
Z3637 := __e.Get(1)
_ = Z3637
tmp1646 := Call(__e, PrimFunc(symshen_4app), V3633, MakeString("\n"), symshen_4s)


tmp1648 := Call(__e, PrimFunc(symshen_4app), V3632, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" has no attributes: "))
__typedS1, __typedOK1 := TypedString(tmp1646)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" has no attributes: ")
__typedArg1 := tmp1646
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp1648)
}
__typedArg0 := tmp1648
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

tmp1649 := Call(__e, try_1catch, tmp1644, tmp1645)


__e.TailApply(tmp1633, tmp1649)
return


}, 1)

tmp1650 := Call(__e, PrimFunc(symlimit), V3634)


tmp1651 := Call(__e, PrimFunc(symhash), V3632, tmp1650)


__e.TailApply(tmp1632, tmp1651)
return


}, 3)

tmp1652 := Call(__e, ns2_1set, symget, tmp1631)


_ = tmp1652

tmp1653 := MakeNative(func(__e *ControlFlow) {
V3639 := __e.Get(1)
_ = V3639
V3640 := __e.Get(2)
_ = V3640
tmp1654 := MakeNative(func(__e *ControlFlow) {
W3641 := __e.Get(1)
_ = W3641
tmp1656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W3641, MakeNumber(0))
}
__typedArg0 := W3641
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1656 {
__e.Return(MakeNumber(1))
return
} else {
__e.Return(W3641)
return
}


}, 1)

tmp1657 := Call(__e, PrimFunc(symshen_4hashkey), V3639)


tmp1658 := Call(__e, PrimFunc(symshen_4mod), tmp1657, V3640)


__e.TailApply(tmp1654, tmp1658)
return


}, 2)

tmp1659 := Call(__e, ns2_1set, symhash, tmp1653)


_ = tmp1659

tmp1660 := MakeNative(func(__e *ControlFlow) {
V3642 := __e.Get(1)
_ = V3642
tmp1661 := MakeNative(func(__e *ControlFlow) {
W3643 := __e.Get(1)
_ = W3643
__e.TailApply(PrimFunc(symshen_4prodbutzero), W3643, MakeNumber(1))
return
}, 1)

tmp1662 := MakeNative(func(__e *ControlFlow) {
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

tmp1663 := Call(__e, PrimFunc(symexplode), V3642)


tmp1664 := Call(__e, PrimFunc(symmap), tmp1662, tmp1663)


__e.TailApply(tmp1661, tmp1664)
return


}, 1)

tmp1665 := Call(__e, ns2_1set, symshen_4hashkey, tmp1660)


_ = tmp1665

tmp1666 := MakeNative(func(__e *ControlFlow) {
V3645 := __e.Get(1)
_ = V3645
V3646 := __e.Get(2)
_ = V3646
tmp1685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3645)
}
__typedArg0 := Nil
__typedArg1 := V3645
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1685 {
__e.Return(V3646)
return
} else {
tmp1683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3645)
}
__typedArg0 := V3645
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1679 Obj

if True == tmp1683 {
tmp1681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3645)
}
__typedArg0 := V3645
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), tmp1681)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp1681
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1680 Obj

if True == tmp1682 {
ifres1680 = True


} else {
ifres1680 = False


}

ifres1679 = ifres1680


} else {
ifres1679 = False


}

if True == ifres1679 {
tmp1667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3645)
}
__typedArg0 := V3645
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp1667, V3646)
return


} else {
tmp1677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3645)
}
__typedArg0 := V3645
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1677 {
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
tmp1668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3645)
}
__typedArg0 := V3645
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3645)
}
__typedArg0 := V3645
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp1668, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V3646)
__typedN1, __typedOK1 := TypedFloat64(tmp1669)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V3646
__typedArg1 := tmp1669
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


} else {
tmp1671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3645)
}
__typedArg0 := V3645
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3645)
}
__typedArg0 := V3645
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp1671, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(V3646)
__typedN1, __typedOK1 := TypedFloat64(tmp1672)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := V3646
__typedArg1 := tmp1672
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

tmp1686 := Call(__e, ns2_1set, symshen_4prodbutzero, tmp1666)


_ = tmp1686

tmp1687 := MakeNative(func(__e *ControlFlow) {
V3647 := __e.Get(1)
_ = V3647
V3648 := __e.Get(2)
_ = V3648
tmp1688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3648, Nil)
}
__typedArg0 := V3648
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp1689 := Call(__e, PrimFunc(symshen_4multiples), V3647, tmp1688)


__e.TailApply(PrimFunc(symshen_4modh), V3647, tmp1689)
return


}, 2)

tmp1690 := Call(__e, ns2_1set, symshen_4mod, tmp1687)


_ = tmp1690

tmp1691 := MakeNative(func(__e *ControlFlow) {
V3653 := __e.Get(1)
_ = V3653
V3654 := __e.Get(2)
_ = V3654
tmp1702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3654)
}
__typedArg0 := V3654
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1698 Obj

if True == tmp1702 {
tmp1700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3654)
}
__typedArg0 := V3654
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(tmp1700)
__typedN1, __typedOK1 := TypedFloat64(V3653)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := tmp1700
__typedArg1 := V3653
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})()

var ifres1699 Obj

if True == tmp1701 {
ifres1699 = True


} else {
ifres1699 = False


}

ifres1698 = ifres1699


} else {
ifres1698 = False


}

if True == ifres1698 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3654)
}
__typedArg0 := V3654
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
} else {
tmp1696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3654)
}
__typedArg0 := V3654
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1696 {
tmp1692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3654)
}
__typedArg0 := V3654
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(2))
__typedN1, __typedOK1 := TypedFloat64(tmp1692)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp1692
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})(), V3654)
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(2))
__typedN1, __typedOK1 := TypedFloat64(tmp1692)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := MakeNumber(2)
__typedArg1 := tmp1692
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})()
__typedArg1 := V3654
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4multiples), V3653, tmp1694)
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

tmp1703 := Call(__e, ns2_1set, symshen_4multiples, tmp1691)


_ = tmp1703

tmp1704 := MakeNative(func(__e *ControlFlow) {
V3661 := __e.Get(1)
_ = V3661
V3662 := __e.Get(2)
_ = V3662
tmp1722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V3661)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V3661
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1722 {
__e.Return(MakeNumber(0))
return
} else {
tmp1720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3662)
}
__typedArg0 := Nil
__typedArg1 := V3662
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1720 {
__e.Return(V3661)
return
} else {
tmp1718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1714 Obj

if True == tmp1718 {
tmp1716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(tmp1716)
__typedN1, __typedOK1 := TypedFloat64(V3661)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := tmp1716
__typedArg1 := V3661
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})()

var ifres1715 Obj

if True == tmp1717 {
ifres1715 = True


} else {
ifres1715 = False


}

ifres1714 = ifres1715


} else {
ifres1714 = False


}

if True == ifres1714 {
tmp1707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1708 := Call(__e, PrimFunc(symempty_2), tmp1707)


if True == tmp1708 {
__e.Return(V3661)
return
} else {
tmp1705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4modh), V3661, tmp1705)
return


}


} else {
tmp1712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1712 {
tmp1709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3662)
}
__typedArg0 := V3662
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4modh), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V3661)
__typedN1, __typedOK1 := TypedFloat64(tmp1709)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V3661
__typedArg1 := tmp1709
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

tmp1723 := Call(__e, ns2_1set, symshen_4modh, tmp1704)


_ = tmp1723

tmp1724 := MakeNative(func(__e *ControlFlow) {
V3665 := __e.Get(1)
_ = V3665
tmp1731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3665)
}
__typedArg0 := Nil
__typedArg1 := V3665
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1731 {
__e.Return(MakeNumber(0))
return
} else {
tmp1729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3665)
}
__typedArg0 := V3665
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1729 {
tmp1725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3665)
}
__typedArg0 := V3665
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3665)
}
__typedArg0 := V3665
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1727 := Call(__e, PrimFunc(symsum), tmp1726)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(tmp1725)
__typedN1, __typedOK1 := TypedFloat64(tmp1727)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := tmp1725
__typedArg1 := tmp1727
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

tmp1732 := Call(__e, ns2_1set, symsum, tmp1724)


_ = tmp1732

tmp1733 := MakeNative(func(__e *ControlFlow) {
V3670 := __e.Get(1)
_ = V3670
tmp1735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3670)
}
__typedArg0 := V3670
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1735 {
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

tmp1736 := Call(__e, ns2_1set, symhead, tmp1733)


_ = tmp1736

tmp1737 := MakeNative(func(__e *ControlFlow) {
V3675 := __e.Get(1)
_ = V3675
tmp1739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3675)
}
__typedArg0 := V3675
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1739 {
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

tmp1740 := Call(__e, ns2_1set, symtail, tmp1737)


_ = tmp1740

tmp1741 := MakeNative(func(__e *ControlFlow) {
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

tmp1742 := Call(__e, ns2_1set, symhdstr, tmp1741)


_ = tmp1742

tmp1743 := MakeNative(func(__e *ControlFlow) {
V3683 := __e.Get(1)
_ = V3683
V3684 := __e.Get(2)
_ = V3684
tmp1754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3683)
}
__typedArg0 := Nil
__typedArg1 := V3683
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1754 {
__e.Return(Nil)
return
} else {
tmp1752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3683)
}
__typedArg0 := V3683
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1752 {
tmp1749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3683)
}
__typedArg0 := V3683
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1750 := Call(__e, PrimFunc(symelement_2), tmp1749, V3684)


if True == tmp1750 {
tmp1744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3683)
}
__typedArg0 := V3683
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3683)
}
__typedArg0 := V3683
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1746 := Call(__e, PrimFunc(symintersection), tmp1745, V3684)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1744, tmp1746)
}
__typedArg0 := tmp1744
__typedArg1 := tmp1746
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp1747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3683)
}
__typedArg0 := V3683
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symintersection), tmp1747, V3684)
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

tmp1755 := Call(__e, ns2_1set, symintersection, tmp1743)


_ = tmp1755

tmp1756 := MakeNative(func(__e *ControlFlow) {
V3685 := __e.Get(1)
_ = V3685
__e.TailApply(PrimFunc(symshen_4reverse_1help), V3685, Nil)
return
}, 1)

tmp1757 := Call(__e, ns2_1set, symreverse, tmp1756)


_ = tmp1757

tmp1758 := MakeNative(func(__e *ControlFlow) {
V3690 := __e.Get(1)
_ = V3690
V3691 := __e.Get(2)
_ = V3691
tmp1765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3690)
}
__typedArg0 := Nil
__typedArg1 := V3690
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1765 {
__e.Return(V3691)
return
} else {
tmp1763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3690)
}
__typedArg0 := V3690
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1763 {
tmp1759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3690)
}
__typedArg0 := V3690
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3690)
}
__typedArg0 := V3690
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1760, V3691)
}
__typedArg0 := tmp1760
__typedArg1 := V3691
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4reverse_1help), tmp1759, tmp1761)
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

tmp1766 := Call(__e, ns2_1set, symshen_4reverse_1help, tmp1758)


_ = tmp1766

tmp1767 := MakeNative(func(__e *ControlFlow) {
V3696 := __e.Get(1)
_ = V3696
V3697 := __e.Get(2)
_ = V3697
tmp1778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3696)
}
__typedArg0 := Nil
__typedArg1 := V3696
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1778 {
__e.Return(V3697)
return
} else {
tmp1776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3696)
}
__typedArg0 := V3696
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1776 {
tmp1773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3696)
}
__typedArg0 := V3696
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1774 := Call(__e, PrimFunc(symelement_2), tmp1773, V3697)


if True == tmp1774 {
tmp1768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3696)
}
__typedArg0 := V3696
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symunion), tmp1768, V3697)
return


} else {
tmp1769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3696)
}
__typedArg0 := V3696
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3696)
}
__typedArg0 := V3696
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1771 := Call(__e, PrimFunc(symunion), tmp1770, V3697)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1769, tmp1771)
}
__typedArg0 := tmp1769
__typedArg1 := tmp1771
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

tmp1779 := Call(__e, ns2_1set, symunion, tmp1767)


_ = tmp1779

tmp1780 := MakeNative(func(__e *ControlFlow) {
V3698 := __e.Get(1)
_ = V3698
tmp1781 := MakeNative(func(__e *ControlFlow) {
W3699 := __e.Get(1)
_ = W3699
tmp1782 := MakeNative(func(__e *ControlFlow) {
W3700 := __e.Get(1)
_ = W3700
tmp1783 := MakeNative(func(__e *ControlFlow) {
W3701 := __e.Get(1)
_ = W3701
tmp1789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("y"), W3701)
}
__typedArg0 := MakeString("y")
__typedArg1 := W3701
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1789 {
__e.Return(True)
return
} else {
tmp1787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("n"), W3701)
}
__typedArg0 := MakeString("n")
__typedArg1 := W3701
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1787 {
__e.Return(False)
return
} else {
tmp1784 := Call(__e, PrimFunc(symstoutput))


tmp1785 := Call(__e, PrimFunc(sympr), MakeString("please answer y or n\n"), tmp1784)


_ = tmp1785

__e.TailApply(PrimFunc(symy_1or_1n_2), V3698)
return


}


}


}, 1)

tmp1790 := Call(__e, PrimFunc(symstinput))


tmp1791 := Call(__e, PrimFunc(symread), tmp1790)


tmp1792 := Call(__e, PrimFunc(symshen_4app), tmp1791, MakeString(""), symshen_4s)


__e.TailApply(tmp1783, tmp1792)
return


}, 1)

tmp1793 := Call(__e, PrimFunc(symstoutput))


tmp1794 := Call(__e, PrimFunc(sympr), MakeString(" (y/n) "), tmp1793)


__e.TailApply(tmp1782, tmp1794)
return


}, 1)

tmp1795 := Call(__e, PrimFunc(symshen_4proc_1nl), V3698)


tmp1796 := Call(__e, PrimFunc(symstoutput))


tmp1797 := Call(__e, PrimFunc(sympr), tmp1795, tmp1796)


__e.TailApply(tmp1781, tmp1797)
return


}, 1)

tmp1798 := Call(__e, ns2_1set, symy_1or_1n_2, tmp1780)


_ = tmp1798

tmp1799 := MakeNative(func(__e *ControlFlow) {
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

tmp1801 := Call(__e, ns2_1set, symnot, tmp1799)


_ = tmp1801

tmp1802 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString(""))
}
__typedArg0 := MakeString("")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}, 0)

tmp1803 := Call(__e, ns2_1set, symabort, tmp1802)


_ = tmp1803

tmp1804 := MakeNative(func(__e *ControlFlow) {
V3708 := __e.Get(1)
_ = V3708
V3709 := __e.Get(2)
_ = V3709
V3710 := __e.Get(3)
_ = V3710
tmp1812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3709, V3710)
}
__typedArg0 := V3709
__typedArg1 := V3710
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1812 {
__e.Return(V3708)
return
} else {
tmp1810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3710)
}
__typedArg0 := V3710
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1810 {
tmp1805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3710)
}
__typedArg0 := V3710
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1806 := Call(__e, PrimFunc(symsubst), V3708, V3709, tmp1805)


tmp1807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3710)
}
__typedArg0 := V3710
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1808 := Call(__e, PrimFunc(symsubst), V3708, V3709, tmp1807)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1806, tmp1808)
}
__typedArg0 := tmp1806
__typedArg1 := tmp1808
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V3710)
return
}


}


}, 3)

tmp1813 := Call(__e, ns2_1set, symsubst, tmp1804)


_ = tmp1813

tmp1814 := MakeNative(func(__e *ControlFlow) {
V3711 := __e.Get(1)
_ = V3711
tmp1815 := Call(__e, PrimFunc(symshen_4app), V3711, MakeString(""), symshen_4a)


__e.TailApply(PrimFunc(symshen_4explode_1h), tmp1815)
return


}, 1)

tmp1816 := Call(__e, ns2_1set, symexplode, tmp1814)


_ = tmp1816

tmp1817 := MakeNative(func(__e *ControlFlow) {
V3714 := __e.Get(1)
_ = V3714
tmp1824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V3714)
}
__typedArg0 := MakeString("")
__typedArg1 := V3714
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1824 {
__e.Return(Nil)
return
} else {
tmp1822 := Call(__e, PrimFunc(symshen_4_7string_2), V3714)


if True == tmp1822 {
tmp1818 := Call(__e, PrimFunc(symhdstr), V3714)


tmp1820 := Call(__e, PrimFunc(symshen_4explode_1h), (func() Obj {
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
return PrimCons(tmp1818, tmp1820)
}
__typedArg0 := tmp1818
__typedArg1 := tmp1820
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

tmp1825 := Call(__e, ns2_1set, symshen_4explode_1h, tmp1817)


_ = tmp1825

tmp1826 := MakeNative(func(__e *ControlFlow) {
V3715 := __e.Get(1)
_ = V3715
tmp1829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3715, MakeString(""))
}
__typedArg0 := V3715
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1827 Obj

if True == tmp1829 {
ifres1827 = MakeString("")


} else {
tmp1828 := Call(__e, PrimFunc(symshen_4app), V3715, MakeString("/"), symshen_4a)


ifres1827 = tmp1828


}

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dhome_1directory_d, ifres1827)
}
__typedArg0 := sym_dhome_1directory_d
__typedArg1 := ifres1827
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp1830 := Call(__e, ns2_1set, symcd, tmp1826)


_ = tmp1830

tmp1831 := MakeNative(func(__e *ControlFlow) {
V3716 := __e.Get(1)
_ = V3716
V3717 := __e.Get(2)
_ = V3717
__e.TailApply(PrimFunc(symshen_4map_1h), V3716, V3717, Nil)
return
}, 2)

tmp1832 := Call(__e, ns2_1set, symmap, tmp1831)


_ = tmp1832

tmp1833 := MakeNative(func(__e *ControlFlow) {
V3718 := __e.Get(1)
_ = V3718
V3719 := __e.Get(2)
_ = V3719
V3720 := __e.Get(3)
_ = V3720
tmp1841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3719)
}
__typedArg0 := Nil
__typedArg1 := V3719
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1841 {
__e.TailApply(PrimFunc(symreverse), V3720)
return
} else {
tmp1839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3719)
}
__typedArg0 := V3719
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1839 {
tmp1834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3719)
}
__typedArg0 := V3719
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3719)
}
__typedArg0 := V3719
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1836 := Call(__e, V3718, tmp1835)


tmp1837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp1836, V3720)
}
__typedArg0 := tmp1836
__typedArg1 := V3720
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4map_1h), V3718, tmp1834, tmp1837)
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

tmp1842 := Call(__e, ns2_1set, symshen_4map_1h, tmp1833)


_ = tmp1842

tmp1843 := MakeNative(func(__e *ControlFlow) {
V3721 := __e.Get(1)
_ = V3721
__e.TailApply(PrimFunc(symshen_4length_1h), V3721, MakeNumber(0))
return
}, 1)

tmp1844 := Call(__e, ns2_1set, symlength, tmp1843)


_ = tmp1844

tmp1845 := MakeNative(func(__e *ControlFlow) {
V3726 := __e.Get(1)
_ = V3726
V3727 := __e.Get(2)
_ = V3727
tmp1849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3726)
}
__typedArg0 := Nil
__typedArg1 := V3726
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1849 {
__e.Return(V3727)
return
} else {
tmp1846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3726)
}
__typedArg0 := V3726
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4length_1h), tmp1846, (func() Obj {
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

tmp1850 := Call(__e, ns2_1set, symshen_4length_1h, tmp1845)


_ = tmp1850

tmp1851 := MakeNative(func(__e *ControlFlow) {
V3733 := __e.Get(1)
_ = V3733
V3734 := __e.Get(2)
_ = V3734
tmp1859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3733, V3734)
}
__typedArg0 := V3733
__typedArg1 := V3734
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1859 {
__e.Return(MakeNumber(1))
return
} else {
tmp1857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3734)
}
__typedArg0 := V3734
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1857 {
tmp1852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3734)
}
__typedArg0 := V3734
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1853 := Call(__e, PrimFunc(symoccurrences), V3733, tmp1852)


tmp1854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3734)
}
__typedArg0 := V3734
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1855 := Call(__e, PrimFunc(symoccurrences), V3733, tmp1854)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(tmp1853)
__typedN1, __typedOK1 := TypedFloat64(tmp1855)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := tmp1853
__typedArg1 := tmp1855
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(MakeNumber(0))
return
}


}


}, 2)

tmp1860 := Call(__e, ns2_1set, symoccurrences, tmp1851)


_ = tmp1860

tmp1861 := MakeNative(func(__e *ControlFlow) {
V3739 := __e.Get(1)
_ = V3739
V3740 := __e.Get(2)
_ = V3740
tmp1874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(1), V3739)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := V3739
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres1871 Obj

if True == tmp1874 {
tmp1873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3740)
}
__typedArg0 := V3740
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres1872 Obj

if True == tmp1873 {
ifres1872 = True


} else {
ifres1872 = False


}

ifres1871 = ifres1872


} else {
ifres1871 = False


}

if True == ifres1871 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3740)
}
__typedArg0 := V3740
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp1869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3740)
}
__typedArg0 := V3740
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1869 {
tmp1862 := (func() Obj {
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

tmp1863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3740)
}
__typedArg0 := V3740
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symnth), tmp1862, tmp1863)
return


} else {
tmp1864 := Call(__e, PrimFunc(symshen_4app), V3740, MakeString("\n"), symshen_4a)


tmp1866 := Call(__e, PrimFunc(symshen_4app), V3739, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(", "))
__typedS1, __typedOK1 := TypedString(tmp1864)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(", ")
__typedArg1 := tmp1864
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("nth applied to "))
__typedS1, __typedOK1 := TypedString(tmp1866)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("nth applied to ")
__typedArg1 := tmp1866
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("nth applied to "))
__typedS1, __typedOK1 := TypedString(tmp1866)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("nth applied to ")
__typedArg1 := tmp1866
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}


}, 2)

tmp1875 := Call(__e, ns2_1set, symnth, tmp1861)


_ = tmp1875

tmp1876 := MakeNative(func(__e *ControlFlow) {
V3741 := __e.Get(1)
_ = V3741
tmp1883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(V3741)
}
__typedArg0 := V3741
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

if True == tmp1883 {
tmp1878 := MakeNative(func(__e *ControlFlow) {
W3742 := __e.Get(1)
_ = W3742
tmp1879 := Call(__e, PrimFunc(symshen_4magless), W3742, MakeNumber(1))


__e.TailApply(PrimFunc(symshen_4integer_1test_2), W3742, tmp1879)
return


}, 1)

tmp1880 := Call(__e, PrimFunc(symshen_4abs), V3741)


tmp1881 := Call(__e, tmp1878, tmp1880)


if True == tmp1881 {
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

tmp1884 := Call(__e, ns2_1set, syminteger_2, tmp1876)


_ = tmp1884

tmp1885 := MakeNative(func(__e *ControlFlow) {
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

tmp1888 := Call(__e, ns2_1set, symshen_4abs, tmp1885)


_ = tmp1888

tmp1889 := MakeNative(func(__e *ControlFlow) {
V3744 := __e.Get(1)
_ = V3744
V3745 := __e.Get(2)
_ = V3745
tmp1890 := MakeNative(func(__e *ControlFlow) {
W3746 := __e.Get(1)
_ = W3746
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(W3746)
__typedN1, __typedOK1 := TypedFloat64(V3744)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := W3746
__typedArg1 := V3744
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
__e.Return(V3745)
return
} else {
__e.TailApply(PrimFunc(symshen_4magless), V3744, W3746)
return
}


}, 1)

__e.TailApply(tmp1890, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(V3745)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(2))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := V3745
__typedArg1 := MakeNumber(2)
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp1894 := Call(__e, ns2_1set, symshen_4magless, tmp1889)


_ = tmp1894

tmp1895 := MakeNative(func(__e *ControlFlow) {
V3750 := __e.Get(1)
_ = V3750
V3751 := __e.Get(2)
_ = V3751
tmp1903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V3750)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V3750
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1903 {
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
tmp1896 := MakeNative(func(__e *ControlFlow) {
W3752 := __e.Get(1)
_ = W3752
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(0))
__typedN1, __typedOK1 := TypedFloat64(W3752)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := MakeNumber(0)
__typedArg1 := W3752
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
__e.TailApply(PrimFunc(symshen_4integer_1test_2), W3752, V3751)
return
}


}, 1)

__e.TailApply(tmp1896, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V3750)
__typedN1, __typedOK1 := TypedFloat64(V3751)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V3750
__typedArg1 := V3751
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())
return


}


}


}, 2)

tmp1904 := Call(__e, ns2_1set, symshen_4integer_1test_2, tmp1895)


_ = tmp1904

tmp1905 := MakeNative(func(__e *ControlFlow) {
V3759 := __e.Get(1)
_ = V3759
V3760 := __e.Get(2)
_ = V3760
tmp1913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3760)
}
__typedArg0 := Nil
__typedArg1 := V3760
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1913 {
__e.Return(Nil)
return
} else {
tmp1911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3760)
}
__typedArg0 := V3760
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp1911 {
tmp1906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3760)
}
__typedArg0 := V3760
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp1907 := Call(__e, V3759, tmp1906)


tmp1908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3760)
}
__typedArg0 := V3760
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp1909 := Call(__e, PrimFunc(symmapcan), V3759, tmp1908)


__e.TailApply(PrimFunc(symappend), tmp1907, tmp1909)
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

tmp1914 := Call(__e, ns2_1set, symmapcan, tmp1905)


_ = tmp1914

tmp1915 := MakeNative(func(__e *ControlFlow) {
V3766 := __e.Get(1)
_ = V3766
V3767 := __e.Get(2)
_ = V3767
tmp1917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3766, V3767)
}
__typedArg0 := V3766
__typedArg1 := V3767
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1917 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp1918 := Call(__e, ns2_1set, sym_a_a, tmp1915)


_ = tmp1918

tmp1919 := MakeNative(func(__e *ControlFlow) {
V3768 := __e.Get(1)
_ = V3768
tmp1929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V3768)
}
__typedArg0 := V3768
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp1929 {
tmp1921 := MakeNative(func(__e *ControlFlow) {
W3769 := __e.Get(1)
_ = W3769
tmp1923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W3769, symshen_4this_1symbol_1is_1unbound)
}
__typedArg0 := W3769
__typedArg1 := symshen_4this_1symbol_1is_1unbound
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1923 {
__e.Return(False)
return
} else {
__e.Return(True)
return
}


}, 1)

tmp1924 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(V3768)
}
__typedArg0 := V3768
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1925 := MakeNative(func(__e *ControlFlow) {
Z3770 := __e.Get(1)
_ = Z3770
__e.Return(symshen_4this_1symbol_1is_1unbound)
return
}, 1)

tmp1926 := Call(__e, try_1catch, tmp1924, tmp1925)


tmp1927 := Call(__e, tmp1921, tmp1926)


if True == tmp1927 {
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

tmp1930 := Call(__e, ns2_1set, symbound_2, tmp1919)


_ = tmp1930

tmp1931 := MakeNative(func(__e *ControlFlow) {
V3771 := __e.Get(1)
_ = V3771
tmp1937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V3771)
}
__typedArg0 := MakeString("")
__typedArg1 := V3771
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1937 {
__e.Return(Nil)
return
} else {
tmp1932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sympos) {
return PrimPos(V3771, MakeNumber(0))
}
__typedArg0 := V3771
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sympos), __typedArg0, __typedArg1)
})()

tmp1933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp1932)
}
__typedArg0 := tmp1932
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp1935 := Call(__e, PrimFunc(symshen_4string_1_6bytes), (func() Obj {
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
return PrimCons(tmp1933, tmp1935)
}
__typedArg0 := tmp1933
__typedArg1 := tmp1935
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp1938 := Call(__e, ns2_1set, symshen_4string_1_6bytes, tmp1931)


_ = tmp1938

tmp1939 := MakeNative(func(__e *ControlFlow) {
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
tmp1941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V3772)
}
__typedArg0 := V3772
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})()

if True == tmp1941 {
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

tmp1944 := Call(__e, ns2_1set, symmaxinferences, tmp1939)


_ = tmp1944

tmp1945 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dinfs_d)
}
__typedArg0 := symshen_4_dinfs_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1946 := Call(__e, ns2_1set, syminferences, tmp1945)


_ = tmp1946

tmp1947 := MakeNative(func(__e *ControlFlow) {
V3773 := __e.Get(1)
_ = V3773
__e.Return(V3773)
return
}, 1)

tmp1948 := Call(__e, ns2_1set, symprotect, tmp1947)


_ = tmp1948

tmp1949 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dstoutput_d)
}
__typedArg0 := sym_dstoutput_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1950 := Call(__e, ns2_1set, symstoutput, tmp1949)


_ = tmp1950

tmp1951 := MakeNative(func(__e *ControlFlow) {
V3774 := __e.Get(1)
_ = V3774
tmp1952 := MakeNative(func(__e *ControlFlow) {
W3775 := __e.Get(1)
_ = W3775
tmp1956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(W3775)
}
__typedArg0 := W3775
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp1956 {
__e.Return(W3775)
return
} else {
tmp1953 := Call(__e, PrimFunc(symshen_4app), V3774, MakeString(" to a symbol"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("cannot intern "))
__typedS1, __typedOK1 := TypedString(tmp1953)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("cannot intern ")
__typedArg1 := tmp1953
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("cannot intern "))
__typedS1, __typedOK1 := TypedString(tmp1953)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("cannot intern ")
__typedArg1 := tmp1953
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 1)

tmp1957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(V3774)
}
__typedArg0 := V3774
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

__e.TailApply(tmp1952, tmp1957)
return


}, 1)

tmp1958 := Call(__e, ns2_1set, symstring_1_6symbol, tmp1951)


_ = tmp1958

tmp1959 := MakeNative(func(__e *ControlFlow) {
V3778 := __e.Get(1)
_ = V3778
tmp1963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V3778)
}
__typedArg0 := sym_7
__typedArg1 := V3778
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1963 {
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
tmp1961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V3778)
}
__typedArg0 := sym_1
__typedArg1 := V3778
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1961 {
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

tmp1964 := Call(__e, ns2_1set, symoptimise, tmp1959)


_ = tmp1964

tmp1965 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dos_d)
}
__typedArg0 := sym_dos_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1966 := Call(__e, ns2_1set, symos, tmp1965)


_ = tmp1966

tmp1967 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dlanguage_d)
}
__typedArg0 := sym_dlanguage_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1968 := Call(__e, ns2_1set, symlanguage, tmp1967)


_ = tmp1968

tmp1969 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dversion_d)
}
__typedArg0 := sym_dversion_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1970 := Call(__e, ns2_1set, symversion, tmp1969)


_ = tmp1970

tmp1971 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dport_d)
}
__typedArg0 := sym_dport_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1972 := Call(__e, ns2_1set, symport, tmp1971)


_ = tmp1972

tmp1973 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dporters_d)
}
__typedArg0 := sym_dporters_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1974 := Call(__e, ns2_1set, symporters, tmp1973)


_ = tmp1974

tmp1975 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dimplementation_d)
}
__typedArg0 := sym_dimplementation_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1976 := Call(__e, ns2_1set, symimplementation, tmp1975)


_ = tmp1976

tmp1977 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_drelease_d)
}
__typedArg0 := sym_drelease_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1978 := Call(__e, ns2_1set, symrelease, tmp1977)


_ = tmp1978

tmp1979 := MakeNative(func(__e *ControlFlow) {
V3779 := __e.Get(1)
_ = V3779
tmp1984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnull, V3779)
}
__typedArg0 := symnull
__typedArg1 := V3779
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1984 {
__e.Return(True)
return
} else {
tmp1980 := MakeNative(func(__e *ControlFlow) {
tmp1981 := Call(__e, PrimFunc(symexternal), V3779)


_ = tmp1981

__e.Return(True)
return


}, 0)

tmp1982 := MakeNative(func(__e *ControlFlow) {
Z3780 := __e.Get(1)
_ = Z3780
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1980, tmp1982)
return


}


}, 1)

tmp1985 := Call(__e, ns2_1set, sympackage_2, tmp1979)


_ = tmp1985

tmp1986 := MakeNative(func(__e *ControlFlow) {
__e.Return(sym_4_4_4)
return
}, 0)

tmp1987 := Call(__e, ns2_1set, symfail, tmp1986)


_ = tmp1987

tmp1988 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_duserdefs_d)
}
__typedArg0 := symshen_4_duserdefs_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1989 := Call(__e, ns2_1set, symuserdefs, tmp1988)


_ = tmp1989

tmp1990 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_doptimise_d)
}
__typedArg0 := symshen_4_doptimise_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1991 := Call(__e, ns2_1set, symoptimise_2, tmp1990)


_ = tmp1991

tmp1992 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dhush_d)
}
__typedArg0 := sym_dhush_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1993 := Call(__e, ns2_1set, symhush_2, tmp1992)


_ = tmp1993

tmp1994 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dshen_1type_1theory_1enabled_2_d)
}
__typedArg0 := symshen_4_dshen_1type_1theory_1enabled_2_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp1995 := Call(__e, ns2_1set, symsystem_1S_2, tmp1994)


_ = tmp1995

tmp1996 := MakeNative(func(__e *ControlFlow) {
V3783 := __e.Get(1)
_ = V3783
tmp2000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V3783)
}
__typedArg0 := sym_7
__typedArg1 := V3783
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2000 {
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
tmp1998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V3783)
}
__typedArg0 := sym_1
__typedArg1 := V3783
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp1998 {
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

tmp2001 := Call(__e, ns2_1set, symenable_1type_1theory, tmp1996)


_ = tmp2001

tmp2002 := MakeNative(func(__e *ControlFlow) {
V3786 := __e.Get(1)
_ = V3786
tmp2006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V3786)
}
__typedArg0 := sym_7
__typedArg1 := V3786
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2006 {
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
tmp2004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V3786)
}
__typedArg0 := sym_1
__typedArg1 := V3786
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2004 {
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

tmp2007 := Call(__e, ns2_1set, symshen_4hush, tmp2002)


_ = tmp2007

tmp2008 := MakeNative(func(__e *ControlFlow) {
V3789 := __e.Get(1)
_ = V3789
tmp2012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V3789)
}
__typedArg0 := sym_7
__typedArg1 := V3789
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2012 {
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
tmp2010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V3789)
}
__typedArg0 := sym_1
__typedArg1 := V3789
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2010 {
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

tmp2013 := Call(__e, ns2_1set, symtc, tmp2008)


_ = tmp2013

tmp2014 := MakeNative(func(__e *ControlFlow) {
V3790 := __e.Get(1)
_ = V3790
tmp2015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsigf_d)
}
__typedArg0 := symshen_4_dsigf_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2016 := Call(__e, PrimFunc(symshen_4unassoc), V3790, tmp2015)


tmp2017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsigf_d, tmp2016)
}
__typedArg0 := symshen_4_dsigf_d
__typedArg1 := tmp2016
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2017

__e.Return(V3790)
return


}, 1)

tmp2018 := Call(__e, ns2_1set, symdestroy, tmp2014)


_ = tmp2018

tmp2019 := MakeNative(func(__e *ControlFlow) {
V3800 := __e.Get(1)
_ = V3800
V3801 := __e.Get(2)
_ = V3801
tmp2037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3801)
}
__typedArg0 := Nil
__typedArg1 := V3801
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2037 {
__e.Return(Nil)
return
} else {
tmp2035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3801)
}
__typedArg0 := V3801
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2026 Obj

if True == tmp2035 {
tmp2033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3801)
}
__typedArg0 := V3801
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2033)
}
__typedArg0 := tmp2033
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2028 Obj

if True == tmp2034 {
tmp2030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3801)
}
__typedArg0 := V3801
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2030)
}
__typedArg0 := tmp2030
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3800, tmp2031)
}
__typedArg0 := V3800
__typedArg1 := tmp2031
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2029 Obj

if True == tmp2032 {
ifres2029 = True


} else {
ifres2029 = False


}

ifres2028 = ifres2029


} else {
ifres2028 = False


}

var ifres2027 Obj

if True == ifres2028 {
ifres2027 = True


} else {
ifres2027 = False


}

ifres2026 = ifres2027


} else {
ifres2026 = False


}

if True == ifres2026 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3801)
}
__typedArg0 := V3801
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
} else {
tmp2024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3801)
}
__typedArg0 := V3801
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2024 {
tmp2020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3801)
}
__typedArg0 := V3801
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3801)
}
__typedArg0 := V3801
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2022 := Call(__e, PrimFunc(symshen_4unassoc), V3800, tmp2021)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2020, tmp2022)
}
__typedArg0 := tmp2020
__typedArg1 := tmp2022
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

tmp2038 := Call(__e, ns2_1set, symshen_4unassoc, tmp2019)


_ = tmp2038

tmp2039 := MakeNative(func(__e *ControlFlow) {
V3802 := __e.Get(1)
_ = V3802
tmp2043 := Call(__e, PrimFunc(sympackage_2), V3802)


if True == tmp2043 {
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
tmp2040 := Call(__e, PrimFunc(symshen_4app), V3802, MakeString(" does not exist\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("package "))
__typedS1, __typedOK1 := TypedString(tmp2040)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("package ")
__typedArg1 := tmp2040
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("package "))
__typedS1, __typedOK1 := TypedString(tmp2040)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("package ")
__typedArg1 := tmp2040
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 1)

tmp2044 := Call(__e, ns2_1set, symin_1package, tmp2039)


_ = tmp2044

tmp2045 := MakeNative(func(__e *ControlFlow) {
V3803 := __e.Get(1)
_ = V3803
V3804 := __e.Get(2)
_ = V3804
tmp2046 := MakeNative(func(__e *ControlFlow) {
W3805 := __e.Get(1)
_ = W3805
tmp2047 := MakeNative(func(__e *ControlFlow) {
W3806 := __e.Get(1)
_ = W3806
tmp2048 := MakeNative(func(__e *ControlFlow) {
W3807 := __e.Get(1)
_ = W3807
tmp2049 := MakeNative(func(__e *ControlFlow) {
W3808 := __e.Get(1)
_ = W3808
__e.Return(V3804)
return
}, 1)

tmp2050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symclose) {
return PrimCloseStream(W3805)
}
__typedArg0 := W3805
return Call(__e, PrimFunc(symclose), __typedArg0)
})()

__e.TailApply(tmp2049, tmp2050)
return


}, 1)

tmp2051 := Call(__e, PrimFunc(sympr), W3806, W3805)


__e.TailApply(tmp2048, tmp2051)
return


}, 1)

tmp2054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V3804)
}
__typedArg0 := V3804
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres2052 Obj

if True == tmp2054 {
ifres2052 = V3804


} else {
tmp2053 := Call(__e, PrimFunc(symshen_4app), V3804, MakeString(""), symshen_4s)


ifres2052 = tmp2053


}

__e.TailApply(tmp2047, ifres2052)
return


}, 1)

tmp2055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symopen) {
return PrimOpenStream(V3803, symout)
}
__typedArg0 := V3803
__typedArg1 := symout
return Call(__e, PrimFunc(symopen), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp2046, tmp2055)
return


}, 2)

tmp2056 := Call(__e, ns2_1set, symwrite_1to_1file, tmp2045)


_ = tmp2056

tmp2057 := MakeNative(func(__e *ControlFlow) {
tmp2058 := Call(__e, PrimFunc(symgensym), symshen_4t)


__e.TailApply(PrimFunc(symshen_4freshterm), tmp2058)
return


}, 0)

tmp2059 := Call(__e, ns2_1set, symfresh, tmp2057)


_ = tmp2059

tmp2060 := MakeNative(func(__e *ControlFlow) {
V3809 := __e.Get(1)
_ = V3809
V3810 := __e.Get(2)
_ = V3810
tmp2061 := MakeNative(func(__e *ControlFlow) {
W3811 := __e.Get(1)
_ = W3811
tmp2062 := MakeNative(func(__e *ControlFlow) {
W3812 := __e.Get(1)
_ = W3812
tmp2063 := MakeNative(func(__e *ControlFlow) {
W3813 := __e.Get(1)
_ = W3813
__e.Return(V3809)
return
}, 1)

tmp2064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dlambdatable_d)
}
__typedArg0 := symshen_4_dlambdatable_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W3812, tmp2064)
}
__typedArg0 := W3812
__typedArg1 := tmp2064
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dlambdatable_d, tmp2065)
}
__typedArg0 := symshen_4_dlambdatable_d
__typedArg1 := tmp2065
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp2063, tmp2066)
return


}, 1)

tmp2067 := Call(__e, PrimFunc(symshen_4lambda_1entry), V3809)


__e.TailApply(tmp2062, tmp2067)
return


}, 1)

tmp2068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2069 := Call(__e, PrimFunc(symput), V3809, symarity, V3810, tmp2068)


__e.TailApply(tmp2061, tmp2069)
return


}, 2)

tmp2070 := Call(__e, ns2_1set, symupdate_1lambda_1table, tmp2060)


_ = tmp2070

tmp2071 := MakeNative(func(__e *ControlFlow) {
V3816 := __e.Get(1)
_ = V3816
V3817 := __e.Get(2)
_ = V3817
tmp2095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V3817)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V3817
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2095 {
tmp2072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspecial_d)
}
__typedArg0 := symshen_4_dspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2073 := Call(__e, PrimFunc(symremove), V3816, tmp2072)


tmp2074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspecial_d, tmp2073)
}
__typedArg0 := symshen_4_dspecial_d
__typedArg1 := tmp2073
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2074

tmp2075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dextraspecial_d)
}
__typedArg0 := symshen_4_dextraspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2076 := Call(__e, PrimFunc(symremove), V3816, tmp2075)


tmp2077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dextraspecial_d, tmp2076)
}
__typedArg0 := symshen_4_dextraspecial_d
__typedArg1 := tmp2076
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2077

__e.Return(V3816)
return


} else {
tmp2093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(1), V3817)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := V3817
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2093 {
tmp2078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspecial_d)
}
__typedArg0 := symshen_4_dspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2079 := Call(__e, PrimFunc(symadjoin), V3816, tmp2078)


tmp2080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspecial_d, tmp2079)
}
__typedArg0 := symshen_4_dspecial_d
__typedArg1 := tmp2079
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2080

tmp2081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dextraspecial_d)
}
__typedArg0 := symshen_4_dextraspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2082 := Call(__e, PrimFunc(symremove), V3816, tmp2081)


tmp2083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dextraspecial_d, tmp2082)
}
__typedArg0 := symshen_4_dextraspecial_d
__typedArg1 := tmp2082
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2083

__e.Return(V3816)
return


} else {
tmp2091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(2), V3817)
}
__typedArg0 := MakeNumber(2)
__typedArg1 := V3817
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2091 {
tmp2084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspecial_d)
}
__typedArg0 := symshen_4_dspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2085 := Call(__e, PrimFunc(symremove), V3816, tmp2084)


tmp2086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspecial_d, tmp2085)
}
__typedArg0 := symshen_4_dspecial_d
__typedArg1 := tmp2085
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2086

tmp2087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dextraspecial_d)
}
__typedArg0 := symshen_4_dextraspecial_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2088 := Call(__e, PrimFunc(symadjoin), V3816, tmp2087)


tmp2089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dextraspecial_d, tmp2088)
}
__typedArg0 := symshen_4_dextraspecial_d
__typedArg1 := tmp2088
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp2089

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

tmp2096 := Call(__e, ns2_1set, symspecialise, tmp2071)


_ = tmp2096

tmp2097 := MakeNative(func(__e *ControlFlow) {
V3818 := __e.Get(1)
_ = V3818
tmp2098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dabsolute_d)
}
__typedArg0 := sym_dabsolute_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3818, tmp2098)
}
__typedArg0 := V3818
__typedArg1 := tmp2098
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dabsolute_d, tmp2099)
}
__typedArg0 := sym_dabsolute_d
__typedArg1 := tmp2099
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp2100 := Call(__e, ns2_1set, symabsolute, tmp2097)


_ = tmp2100

tmp2101 := MakeNative(func(__e *ControlFlow) {
V3819 := __e.Get(1)
_ = V3819
tmp2102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dabsolute_d)
}
__typedArg0 := sym_dabsolute_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp2103 := Call(__e, PrimFunc(symremove), V3819, tmp2102)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dabsolute_d, tmp2103)
}
__typedArg0 := sym_dabsolute_d
__typedArg1 := tmp2103
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 1)

__e.TailApply(ns2_1set, symunabsolute, tmp2101)
return




}, 0)

