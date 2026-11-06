package main

import . "github.com/pyrex41/shen-go/kl"

var SysMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp1207 := MakeNative1(func(__e *ControlFlow, V3825 Obj) Obj {
return TailCall0(__e, V3825)
panic("unreachable")
})

tmp1208 := Call2(__e, ns2_1set, symthaw, tmp1207)


_ = tmp1208

tmp1209 := MakeNative1(func(__e *ControlFlow, V3826 Obj) Obj {
tmp1210 := Call1(__e, SymFn(symmacroexpand), V3826)


tmp1211 := Call1(__e, SymFn(symshen_4find_1types), V3826)


tmp1212 := Call2(__e, SymFn(symshen_4process_1applications), tmp1210, tmp1211)


tmp1213 := Call1(__e, SymFn(symshen_4shen_1_6kl), tmp1212)


return TailCall1(__e, SymFn(symeval_1kl), tmp1213)


panic("unreachable")
})

tmp1214 := Call2(__e, ns2_1set, symeval, tmp1209)


_ = tmp1214

tmp1215 := MakeNative1(func(__e *ControlFlow, V3827 Obj) Obj {
tmp1222 := GEq(__e, symnull, V3827)

if True == tmp1222 {
return Nil
} else {
tmp1216 := MakeNative0(func(__e *ControlFlow) Obj {
tmp1217 := GValue(__e, sym_dproperty_1vector_d)

return TailCall3(__e, SymFn(symget), V3827, symshen_4external_1symbols, tmp1217)


panic("unreachable")
})

tmp1218 := MakeNative1(func(__e *ControlFlow, Z3828 Obj) Obj {
tmp1219 := Call3(__e, SymFn(symshen_4app), V3827, MakeString(" does not exist.\n;"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("package "), tmp1219))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("package "), tmp1219)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp1216, tmp1218)


}


panic("unreachable")
})

tmp1223 := Call2(__e, ns2_1set, symexternal, tmp1215)


_ = tmp1223

tmp1224 := MakeNative1(func(__e *ControlFlow, V3829 Obj) Obj {
tmp1231 := GEq(__e, symnull, V3829)

if True == tmp1231 {
return Nil
} else {
tmp1225 := MakeNative0(func(__e *ControlFlow) Obj {
tmp1226 := GValue(__e, sym_dproperty_1vector_d)

return TailCall3(__e, SymFn(symget), V3829, symshen_4internal_1symbols, tmp1226)


panic("unreachable")
})

tmp1227 := MakeNative1(func(__e *ControlFlow, Z3830 Obj) Obj {
tmp1228 := Call3(__e, SymFn(symshen_4app), V3829, MakeString(" does not exist.\n;"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("package "), tmp1228))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("package "), tmp1228)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp1225, tmp1227)


}


panic("unreachable")
})

tmp1232 := Call2(__e, ns2_1set, syminternal, tmp1224)


_ = tmp1232

tmp1233 := MakeNative2(func(__e *ControlFlow, V3831, V3832 Obj) Obj {
tmp1235 := Call1(__e, V3831, V3832)


if True == tmp1235 {
return TailCall0(__e, SymFn(symfail))
} else {
return V3832
}


panic("unreachable")
})

tmp1236 := Call2(__e, ns2_1set, symfail_1if, tmp1233)


_ = tmp1236

tmp1237 := MakeNative2(func(__e *ControlFlow, V3833, V3834 Obj) Obj {
return PrimGuard2(__e, symcn, PrimStringConcat, V3833, V3834)
panic("unreachable")
})

tmp1238 := Call2(__e, ns2_1set, sym_8s, tmp1237)


_ = tmp1238

tmp1239 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dtc_d)
panic("unreachable")
})

tmp1240 := Call2(__e, ns2_1set, symtc_2, tmp1239)


_ = tmp1240

tmp1241 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_doccurs_d)
panic("unreachable")
})

tmp1242 := Call2(__e, ns2_1set, symoccurs_2, tmp1241)


_ = tmp1242

tmp1243 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dfactorise_2_d)
panic("unreachable")
})

tmp1244 := Call2(__e, ns2_1set, symfactorise_2, tmp1243)


_ = tmp1244

tmp1245 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dtracking_d)
panic("unreachable")
})

tmp1246 := Call2(__e, ns2_1set, symtracked, tmp1245)


_ = tmp1246

tmp1247 := MakeNative1(func(__e *ControlFlow, V3835 Obj) Obj {
tmp1248 := MakeNative0(func(__e *ControlFlow) Obj {
tmp1249 := GValue(__e, sym_dproperty_1vector_d)

return TailCall3(__e, SymFn(symget), V3835, symshen_4source, tmp1249)


panic("unreachable")
})

tmp1250 := MakeNative1(func(__e *ControlFlow, Z3836 Obj) Obj {
tmp1251 := Call3(__e, SymFn(symshen_4app), V3835, MakeString(" not found.\n"), symshen_4a)


return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp1251)


panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp1248, tmp1250)


panic("unreachable")
})

tmp1252 := Call2(__e, ns2_1set, symps, tmp1247)


_ = tmp1252

tmp1253 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dstinput_d)
panic("unreachable")
})

tmp1254 := Call2(__e, ns2_1set, symstinput, tmp1253)


_ = tmp1254

tmp1255 := MakeNative1(func(__e *ControlFlow, V3837 Obj) Obj {
tmp1257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(GAdd(__e, V3837, MakeInteger(1)))
}
__typedArg0 := GAdd(__e, V3837, MakeInteger(1))
return Call1(__e, SymFn(symabsvector), __typedArg0)
})()

let__1164 := tmp1257
_ = let__1164

tmp1258 := GVecSet(__e, let__1164, MakeInteger(0), V3837)

let__1165 := tmp1258
_ = let__1165

tmp1262 := GEq(__e, V3837, MakeInteger(0))

var ifres1259 Obj

if True == tmp1262 {
ifres1259 = let__1165


} else {
tmp1260 := Call0(__e, SymFn(symfail))


tmp1261 := Call4(__e, SymFn(symshen_4fillvector), let__1165, MakeInteger(1), V3837, tmp1260)


ifres1259 = tmp1261


}

let__1166 := ifres1259
_ = let__1166

return let__1166


panic("unreachable")
})

tmp1263 := Call2(__e, ns2_1set, symvector, tmp1255)


_ = tmp1263

tmp1264 := (func() Obj {
var __self Obj
__self = MakeNative4(func(__e *ControlFlow, __self1, __self2, __self3, __self4 Obj) Obj {
__selftop:
V3842 := __self1
_ = V3842
V3843 := __self2
_ = V3843
V3844 := __self3
_ = V3844
V3845 := __self4
_ = V3845
tmp1268 := GEq(__e, V3843, V3844)

if True == tmp1268 {
return GVecSet(__e, V3842, V3844, V3845)
} else {
tmp1265 := GVecSet(__e, V3842, V3843, V3845)

if SymFn(symshen_4fillvector) == __self {
__self1, __self2, __self3, __self4 = tmp1265, GAdd(__e, MakeInteger(1), V3843), V3844, V3845
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4fillvector), tmp1265, GAdd(__e, MakeInteger(1), V3843), V3844, V3845)


}


panic("unreachable")
})
return __self
})()

tmp1269 := Call2(__e, ns2_1set, symshen_4fillvector, tmp1264)


_ = tmp1269

tmp1270 := MakeNative1(func(__e *ControlFlow, V3846 Obj) Obj {
tmp1277 := GVectorP(__e, V3846)

if True == tmp1277 {
tmp1271 := MakeNative0(func(__e *ControlFlow) Obj {
return GVecGet(__e, V3846, MakeInteger(0))
panic("unreachable")
})

tmp1272 := MakeNative1(func(__e *ControlFlow, Z3848 Obj) Obj {
return MakeInteger(-1)
panic("unreachable")
})

tmp1273 := Call2(__e, try_1catch, tmp1271, tmp1272)


let__1167 := tmp1273
_ = let__1167

tmp1275 := PrimGuard1(__e, symnumber_2, PrimIsNumber, let__1167)

if True == tmp1275 {
return GGe(__e, let__1167, MakeInteger(0))
} else {
return False
}


} else {
return False
}


panic("unreachable")
})

tmp1278 := Call2(__e, ns2_1set, symvector_2, tmp1270)


_ = tmp1278

tmp1279 := MakeNative3(func(__e *ControlFlow, V3849, V3850, V3851 Obj) Obj {
tmp1281 := GEq(__e, V3850, MakeInteger(0))

if True == tmp1281 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("cannot access 0th element of a vector\n"))
} else {
return GVecSet(__e, V3849, V3850, V3851)
}


panic("unreachable")
})

tmp1282 := Call2(__e, ns2_1set, symvector_1_6, tmp1279)


_ = tmp1282

tmp1283 := MakeNative2(func(__e *ControlFlow, V3852, V3853 Obj) Obj {
tmp1289 := GEq(__e, V3853, MakeInteger(0))

if True == tmp1289 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("cannot access 0th element of a vector\n"))
} else {
tmp1284 := GVecGet(__e, V3852, V3853)

let__1168 := tmp1284
_ = let__1168

tmp1286 := Call0(__e, SymFn(symfail))


tmp1287 := GEq(__e, let__1168, tmp1286)

if True == tmp1287 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("vector element not found\n"))
} else {
return let__1168
}


}


panic("unreachable")
})

tmp1290 := Call2(__e, ns2_1set, sym_5_1vector, tmp1283)


_ = tmp1290

tmp1291 := MakeNative1(func(__e *ControlFlow, V3855 Obj) Obj {
tmp1293 := PrimGuard1(__e, syminteger_2, PrimIsInteger, V3855)

if True == tmp1293 {
return GGe(__e, V3855, MakeInteger(0))
} else {
return False
}


panic("unreachable")
})

tmp1294 := Call2(__e, ns2_1set, symshen_4posint_2, tmp1291)


_ = tmp1294

tmp1295 := MakeNative1(func(__e *ControlFlow, V3856 Obj) Obj {
return GVecGet(__e, V3856, MakeInteger(0))
panic("unreachable")
})

tmp1296 := Call2(__e, ns2_1set, symlimit, tmp1295)


_ = tmp1296

tmp1297 := MakeNative1(func(__e *ControlFlow, V3857 Obj) Obj {
tmp1322 := Call1(__e, SymFn(symboolean_2), V3857)


var ifres1312 Obj

if True == tmp1322 {
ifres1312 = True


} else {
tmp1321 := PrimGuard1(__e, symnumber_2, PrimIsNumber, V3857)

var ifres1313 Obj

if True == tmp1321 {
ifres1313 = True


} else {
tmp1320 := PrimGuard1(__e, symstring_2, PrimIsString, V3857)

var ifres1314 Obj

if True == tmp1320 {
ifres1314 = True


} else {
tmp1319 := GConsP(__e, V3857)

var ifres1315 Obj

if True == tmp1319 {
ifres1315 = True


} else {
tmp1318 := Call1(__e, SymFn(symempty_2), V3857)


var ifres1316 Obj

if True == tmp1318 {
ifres1316 = True


} else {
tmp1317 := Call1(__e, SymFn(symvector_2), V3857)


ifres1316 = tmp1317


}

ifres1315 = ifres1316


}

ifres1314 = ifres1315


}

ifres1313 = ifres1314


}

ifres1312 = ifres1313


}

if True == ifres1312 {
return False
} else {
tmp1302 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp1303 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp1304 := PrimGuard1(__e, symintern, PrimIntern, MakeString(","))

tmp1305 := GCons(__e, tmp1304, Nil)

tmp1306 := GCons(__e, tmp1303, tmp1305)

tmp1307 := GCons(__e, tmp1302, tmp1306)

tmp1308 := GCons(__e, sym_j, tmp1307)

tmp1309 := GCons(__e, sym_i, tmp1308)

tmp1310 := Call2(__e, SymFn(symelement_2), V3857, tmp1309)


if True == tmp1310 {
return True
} else {
tmp1298 := MakeNative0(func(__e *ControlFlow) Obj {
tmp1299 := PrimGuard1(__e, symstr, PrimStr, V3857)

let__1169 := tmp1299
_ = let__1169

return TailCall1(__e, SymFn(symshen_4analyse_1symbol_2), let__1169)


panic("unreachable")
})

tmp1300 := MakeNative1(func(__e *ControlFlow, Z3859 Obj) Obj {
return False
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp1298, tmp1300)


}


}


panic("unreachable")
})

tmp1323 := Call2(__e, ns2_1set, symsymbol_2, tmp1297)


_ = tmp1323

tmp1324 := MakeNative1(func(__e *ControlFlow, V3862 Obj) Obj {
tmp1331 := Call1(__e, SymFn(symshen_4_7string_2), V3862)


if True == tmp1331 {
tmp1327 := Call1(__e, SymFn(symhdstr), V3862)


tmp1328 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp1327)

tmp1329 := Call1(__e, SymFn(symshen_4alpha_2), tmp1328)


if True == tmp1329 {
return TailCall1(__e, SymFn(symshen_4alphanums_2), PrimGuard1(__e, symtlstr, PrimTailString, V3862))


} else {
return False
}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.analyse-symbol?"))
}


panic("unreachable")
})

tmp1332 := Call2(__e, ns2_1set, symshen_4analyse_1symbol_2, tmp1324)


_ = tmp1332

tmp1333 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V3865 := __self1
_ = V3865
tmp1344 := GEq(__e, MakeString(""), V3865)

if True == tmp1344 {
return True
} else {
tmp1342 := Call1(__e, SymFn(symshen_4_7string_2), V3865)


if True == tmp1342 {
tmp1334 := Call1(__e, SymFn(symhdstr), V3865)


tmp1335 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp1334)

let__1170 := tmp1335
_ = let__1170

tmp1340 := Call1(__e, SymFn(symshen_4alpha_2), let__1170)


var ifres1338 Obj

if True == tmp1340 {
ifres1338 = True


} else {
tmp1339 := Call1(__e, SymFn(symshen_4digit_2), let__1170)


ifres1338 = tmp1339


}

if True == ifres1338 {
if SymFn(symshen_4alphanums_2) == __self {
__self1 = PrimGuard1(__e, symtlstr, PrimTailString, V3865)
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4alphanums_2), PrimGuard1(__e, symtlstr, PrimTailString, V3865))


} else {
return False
}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.alphanums?"))
}


}


panic("unreachable")
})
return __self
})()

tmp1345 := Call2(__e, ns2_1set, symshen_4alphanums_2, tmp1333)


_ = tmp1345

tmp1346 := MakeNative1(func(__e *ControlFlow, V3867 Obj) Obj {
tmp1355 := Call1(__e, SymFn(symboolean_2), V3867)


var ifres1351 Obj

if True == tmp1355 {
ifres1351 = True


} else {
tmp1354 := PrimGuard1(__e, symnumber_2, PrimIsNumber, V3867)

var ifres1352 Obj

if True == tmp1354 {
ifres1352 = True


} else {
tmp1353 := PrimGuard1(__e, symstring_2, PrimIsString, V3867)

ifres1352 = tmp1353


}

ifres1351 = ifres1352


}

if True == ifres1351 {
return False
} else {
tmp1347 := MakeNative0(func(__e *ControlFlow) Obj {
tmp1348 := PrimGuard1(__e, symstr, PrimStr, V3867)

let__1171 := tmp1348
_ = let__1171

return TailCall1(__e, SymFn(symshen_4analyse_1variable_2), let__1171)


panic("unreachable")
})

tmp1349 := MakeNative1(func(__e *ControlFlow, Z3869 Obj) Obj {
return False
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp1347, tmp1349)


}


panic("unreachable")
})

tmp1356 := Call2(__e, ns2_1set, symvariable_2, tmp1346)


_ = tmp1356

tmp1357 := MakeNative1(func(__e *ControlFlow, V3872 Obj) Obj {
tmp1364 := Call1(__e, SymFn(symshen_4_7string_2), V3872)


if True == tmp1364 {
tmp1360 := Call1(__e, SymFn(symhdstr), V3872)


tmp1361 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp1360)

tmp1362 := Call1(__e, SymFn(symshen_4uppercase_2), tmp1361)


if True == tmp1362 {
return TailCall1(__e, SymFn(symshen_4alphanums_2), PrimGuard1(__e, symtlstr, PrimTailString, V3872))


} else {
return False
}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.analyse-variable?"))
}


panic("unreachable")
})

tmp1365 := Call2(__e, ns2_1set, symshen_4analyse_1variable_2, tmp1357)


_ = tmp1365

tmp1366 := MakeNative1(func(__e *ControlFlow, V3873 Obj) Obj {
tmp1367 := GValue(__e, symshen_4_dgensym_d)

tmp1369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dgensym_d, GAdd(__e, MakeInteger(1), tmp1367))
}
__typedArg0 := symshen_4_dgensym_d
__typedArg1 := GAdd(__e, MakeInteger(1), tmp1367)
return Call2(__e, SymFn(symset), __typedArg0, __typedArg1)
})()

return TailCall2(__e, SymFn(symconcat), V3873, tmp1369)


panic("unreachable")
})

tmp1370 := Call2(__e, ns2_1set, symgensym, tmp1366)


_ = tmp1370

tmp1371 := MakeNative2(func(__e *ControlFlow, V3874, V3875 Obj) Obj {
tmp1372 := PrimGuard1(__e, symstr, PrimStr, V3874)

tmp1373 := PrimGuard1(__e, symstr, PrimStr, V3875)

return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(PrimGuard2(__e, symcn, PrimStringConcat, tmp1372, tmp1373))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, tmp1372, tmp1373)
return Call1(__e, SymFn(symintern), __typedArg0)
})()


panic("unreachable")
})

tmp1375 := Call2(__e, ns2_1set, symconcat, tmp1371)


_ = tmp1375

tmp1376 := MakeNative2(func(__e *ControlFlow, V3876, V3877 Obj) Obj {
tmp1377 := PrimGuard1(__e, symabsvector, PrimAbsvector, MakeInteger(3))

let__1172 := tmp1377
_ = let__1172

tmp1378 := GVecSet(__e, let__1172, MakeInteger(0), symshen_4tuple)

let__1173 := tmp1378
_ = let__1173

tmp1379 := GVecSet(__e, let__1172, MakeInteger(1), V3876)

let__1174 := tmp1379
_ = let__1174

tmp1380 := GVecSet(__e, let__1172, MakeInteger(2), V3877)

let__1175 := tmp1380
_ = let__1175

return let__1172


panic("unreachable")
})

tmp1381 := Call2(__e, ns2_1set, sym_8p, tmp1376)


_ = tmp1381

tmp1382 := MakeNative1(func(__e *ControlFlow, V3882 Obj) Obj {
return GVecGet(__e, V3882, MakeInteger(1))
panic("unreachable")
})

tmp1383 := Call2(__e, ns2_1set, symfst, tmp1382)


_ = tmp1383

tmp1384 := MakeNative1(func(__e *ControlFlow, V3883 Obj) Obj {
return GVecGet(__e, V3883, MakeInteger(2))
panic("unreachable")
})

tmp1385 := Call2(__e, ns2_1set, symsnd, tmp1384)


_ = tmp1385

tmp1386 := MakeNative1(func(__e *ControlFlow, V3884 Obj) Obj {
tmp1391 := GVectorP(__e, V3884)

if True == tmp1391 {
tmp1387 := MakeNative0(func(__e *ControlFlow) Obj {
return GVecGet(__e, V3884, MakeInteger(0))
panic("unreachable")
})

tmp1388 := MakeNative1(func(__e *ControlFlow, Z3885 Obj) Obj {
return symshen_4not_1tuple
panic("unreachable")
})

tmp1389 := Call2(__e, try_1catch, tmp1387, tmp1388)


return GEq(__e, symshen_4tuple, tmp1389)


} else {
return False
}


panic("unreachable")
})

tmp1392 := Call2(__e, ns2_1set, symtuple_2, tmp1386)


_ = tmp1392

tmp1393 := MakeNative2(func(__e *ControlFlow, V3890, V3891 Obj) Obj {
tmp1400 := GEq(__e, Nil, V3890)

if True == tmp1400 {
return V3891
} else {
tmp1398 := GConsP(__e, V3890)

if True == tmp1398 {
tmp1394 := GHd(__e, V3890)

tmp1395 := GTl(__e, V3890)

tmp1396 := Call2(__e, SymFn(symappend), tmp1395, V3891)


return GCons(__e, tmp1394, tmp1396)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to append a non-list"))
}


}


panic("unreachable")
})

tmp1401 := Call2(__e, ns2_1set, symappend, tmp1393)


_ = tmp1401

tmp1402 := MakeNative2(func(__e *ControlFlow, V3892, V3893 Obj) Obj {
tmp1403 := Call1(__e, SymFn(symlimit), V3893)


let__1176 := tmp1403
_ = let__1176

tmp1405 := Call1(__e, SymFn(symvector), GAdd(__e, let__1176, MakeInteger(1)))


let__1177 := tmp1405
_ = let__1177

tmp1406 := Call3(__e, SymFn(symvector_1_6), let__1177, MakeInteger(1), V3892)


let__1178 := tmp1406
_ = let__1178

tmp1408 := GEq(__e, let__1176, MakeInteger(0))

if True == tmp1408 {
return let__1178
} else {
return TailCall4(__e, SymFn(symshen_4_8v_1help), V3893, MakeInteger(1), let__1176, let__1178)
}


panic("unreachable")
})

tmp1409 := Call2(__e, ns2_1set, sym_8v, tmp1402)


_ = tmp1409

tmp1410 := (func() Obj {
var __self Obj
__self = MakeNative4(func(__e *ControlFlow, __self1, __self2, __self3, __self4 Obj) Obj {
__selftop:
V3898 := __self1
_ = V3898
V3899 := __self2
_ = V3899
V3900 := __self3
_ = V3900
V3901 := __self4
_ = V3901
tmp1416 := GEq(__e, V3899, V3900)

if True == tmp1416 {
return TailCall4(__e, SymFn(symshen_4copyfromvector), V3898, V3901, V3900, GAdd(__e, V3900, MakeInteger(1)))


} else {
tmp1412 := GAdd(__e, V3899, MakeInteger(1))

tmp1414 := Call4(__e, SymFn(symshen_4copyfromvector), V3898, V3901, V3899, GAdd(__e, V3899, MakeInteger(1)))


if SymFn(symshen_4_8v_1help) == __self {
__self1, __self2, __self3, __self4 = V3898, tmp1412, V3900, tmp1414
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4_8v_1help), V3898, tmp1412, V3900, tmp1414)


}


panic("unreachable")
})
return __self
})()

tmp1417 := Call2(__e, ns2_1set, symshen_4_8v_1help, tmp1410)


_ = tmp1417

tmp1418 := MakeNative4(func(__e *ControlFlow, V3902, V3903, V3904, V3905 Obj) Obj {
tmp1419 := MakeNative0(func(__e *ControlFlow) Obj {
tmp1420 := Call2(__e, SymFn(sym_5_1vector), V3902, V3904)


return TailCall3(__e, SymFn(symvector_1_6), V3903, V3905, tmp1420)


panic("unreachable")
})

tmp1421 := MakeNative1(func(__e *ControlFlow, Z3906 Obj) Obj {
return V3903
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp1419, tmp1421)


panic("unreachable")
})

tmp1422 := Call2(__e, ns2_1set, symshen_4copyfromvector, tmp1418)


_ = tmp1422

tmp1423 := MakeNative1(func(__e *ControlFlow, V3907 Obj) Obj {
tmp1424 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall2(__e, SymFn(sym_5_1vector), V3907, MakeInteger(1))
panic("unreachable")
})

tmp1425 := MakeNative1(func(__e *ControlFlow, Z3908 Obj) Obj {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("hdv needs a non-empty vector as an argument\n"))
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp1424, tmp1425)


panic("unreachable")
})

tmp1426 := Call2(__e, ns2_1set, symhdv, tmp1423)


_ = tmp1426

tmp1427 := MakeNative1(func(__e *ControlFlow, V3909 Obj) Obj {
tmp1428 := Call1(__e, SymFn(symlimit), V3909)


let__1179 := tmp1428
_ = let__1179

tmp1436 := GEq(__e, let__1179, MakeInteger(0))

if True == tmp1436 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("cannot take the tail of the empty vector\n"))
} else {
tmp1434 := GEq(__e, let__1179, MakeInteger(1))

if True == tmp1434 {
return TailCall1(__e, SymFn(symvector), MakeInteger(0))
} else {
tmp1430 := Call1(__e, SymFn(symvector), GSub(__e, let__1179, MakeInteger(1)))


let__1180 := tmp1430
_ = let__1180

tmp1432 := Call1(__e, SymFn(symvector), GSub(__e, let__1179, MakeInteger(1)))


return TailCall4(__e, SymFn(symshen_4tlv_1help), V3909, MakeInteger(2), let__1179, tmp1432)


}


}


panic("unreachable")
})

tmp1437 := Call2(__e, ns2_1set, symtlv, tmp1427)


_ = tmp1437

tmp1438 := (func() Obj {
var __self Obj
__self = MakeNative4(func(__e *ControlFlow, __self1, __self2, __self3, __self4 Obj) Obj {
__selftop:
V3913 := __self1
_ = V3913
V3914 := __self2
_ = V3914
V3915 := __self3
_ = V3915
V3916 := __self4
_ = V3916
tmp1444 := GEq(__e, V3914, V3915)

if True == tmp1444 {
return TailCall4(__e, SymFn(symshen_4copyfromvector), V3913, V3916, V3915, GSub(__e, V3915, MakeInteger(1)))


} else {
tmp1440 := GAdd(__e, V3914, MakeInteger(1))

tmp1442 := Call4(__e, SymFn(symshen_4copyfromvector), V3913, V3916, V3914, GSub(__e, V3914, MakeInteger(1)))


if SymFn(symshen_4tlv_1help) == __self {
__self1, __self2, __self3, __self4 = V3913, tmp1440, V3915, tmp1442
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4tlv_1help), V3913, tmp1440, V3915, tmp1442)


}


panic("unreachable")
})
return __self
})()

tmp1445 := Call2(__e, ns2_1set, symshen_4tlv_1help, tmp1438)


_ = tmp1445

tmp1446 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V3928 := __self1
_ = V3928
V3929 := __self2
_ = V3929
tmp1460 := GEq(__e, Nil, V3929)

if True == tmp1460 {
return Nil
} else {
tmp1458 := GConsP(__e, V3929)

var ifres1451 Obj

if True == tmp1458 {
tmp1456 := GHd(__e, V3929)

tmp1457 := GConsP(__e, tmp1456)

var ifres1452 Obj

if True == tmp1457 {
tmp1453 := GHd(__e, V3929)

tmp1454 := GHd(__e, tmp1453)

tmp1455 := GEq(__e, V3928, tmp1454)

ifres1452 = tmp1455


} else {
ifres1452 = False


}

ifres1451 = ifres1452


} else {
ifres1451 = False


}

if True == ifres1451 {
return GHd(__e, V3929)
} else {
tmp1449 := GConsP(__e, V3929)

if True == tmp1449 {
tmp1447 := GTl(__e, V3929)

if SymFn(symassoc) == __self {
__self1, __self2 = V3928, tmp1447
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symassoc), V3928, tmp1447)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to search a non-list with assoc\n"))
}


}


}


panic("unreachable")
})
return __self
})()

tmp1461 := Call2(__e, ns2_1set, symassoc, tmp1446)


_ = tmp1461

tmp1462 := MakeNative3(func(__e *ControlFlow, V3933, V3934, V3935 Obj) Obj {
tmp1483 := GEq(__e, Nil, V3935)

if True == tmp1483 {
tmp1463 := GCons(__e, V3933, V3934)

return GCons(__e, tmp1463, Nil)


} else {
tmp1481 := GConsP(__e, V3935)

var ifres1474 Obj

if True == tmp1481 {
tmp1479 := GHd(__e, V3935)

tmp1480 := GConsP(__e, tmp1479)

var ifres1475 Obj

if True == tmp1480 {
tmp1476 := GHd(__e, V3935)

tmp1477 := GHd(__e, tmp1476)

tmp1478 := GEq(__e, V3933, tmp1477)

ifres1475 = tmp1478


} else {
ifres1475 = False


}

ifres1474 = ifres1475


} else {
ifres1474 = False


}

if True == ifres1474 {
tmp1464 := GHd(__e, V3935)

tmp1465 := GHd(__e, tmp1464)

tmp1466 := GCons(__e, tmp1465, V3934)

tmp1467 := GTl(__e, V3935)

return GCons(__e, tmp1466, tmp1467)


} else {
tmp1472 := GConsP(__e, V3935)

if True == tmp1472 {
tmp1468 := GHd(__e, V3935)

tmp1469 := GTl(__e, V3935)

tmp1470 := Call3(__e, SymFn(symshen_4assoc_1set), V3933, V3934, tmp1469)


return GCons(__e, tmp1468, tmp1470)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4assoc_1set)
}


}


}


panic("unreachable")
})

tmp1484 := Call2(__e, ns2_1set, symshen_4assoc_1set, tmp1462)


_ = tmp1484

tmp1485 := MakeNative2(func(__e *ControlFlow, V3939, V3940 Obj) Obj {
tmp1501 := GEq(__e, Nil, V3940)

if True == tmp1501 {
return Nil
} else {
tmp1499 := GConsP(__e, V3940)

var ifres1492 Obj

if True == tmp1499 {
tmp1497 := GHd(__e, V3940)

tmp1498 := GConsP(__e, tmp1497)

var ifres1493 Obj

if True == tmp1498 {
tmp1494 := GHd(__e, V3940)

tmp1495 := GHd(__e, tmp1494)

tmp1496 := GEq(__e, V3939, tmp1495)

ifres1493 = tmp1496


} else {
ifres1493 = False


}

ifres1492 = ifres1493


} else {
ifres1492 = False


}

if True == ifres1492 {
return GTl(__e, V3940)
} else {
tmp1490 := GConsP(__e, V3940)

if True == tmp1490 {
tmp1486 := GHd(__e, V3940)

tmp1487 := GTl(__e, V3940)

tmp1488 := Call2(__e, SymFn(symshen_4assoc_1rm), V3939, tmp1487)


return GCons(__e, tmp1486, tmp1488)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4assoc_1rm)
}


}


}


panic("unreachable")
})

tmp1502 := Call2(__e, ns2_1set, symshen_4assoc_1rm, tmp1485)


_ = tmp1502

tmp1503 := MakeNative1(func(__e *ControlFlow, V3943 Obj) Obj {
tmp1507 := GEq(__e, True, V3943)

if True == tmp1507 {
return True
} else {
tmp1505 := GEq(__e, False, V3943)

if True == tmp1505 {
return True
} else {
return False
}


}


panic("unreachable")
})

tmp1508 := Call2(__e, ns2_1set, symboolean_2, tmp1503)


_ = tmp1508

tmp1509 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V3944 := __self1
_ = V3944
tmp1514 := GEq(__e, MakeInteger(0), V3944)

if True == tmp1514 {
return MakeInteger(0)
} else {
tmp1510 := Call0(__e, SymFn(symstoutput))


tmp1511 := Call2(__e, SymFn(sympr), MakeString("\n"), tmp1510)


_ = tmp1511

if SymFn(symnl) == __self {
__self1 = GSub(__e, V3944, MakeInteger(1))
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symnl), GSub(__e, V3944, MakeInteger(1)))


}


panic("unreachable")
})
return __self
})()

tmp1515 := Call2(__e, ns2_1set, symnl, tmp1509)


_ = tmp1515

tmp1516 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V3951 := __self1
_ = V3951
V3952 := __self2
_ = V3952
tmp1527 := GEq(__e, Nil, V3951)

if True == tmp1527 {
return Nil
} else {
tmp1525 := GConsP(__e, V3951)

if True == tmp1525 {
tmp1522 := GHd(__e, V3951)

tmp1523 := Call2(__e, SymFn(symelement_2), tmp1522, V3952)


if True == tmp1523 {
tmp1517 := GTl(__e, V3951)

if SymFn(symdifference) == __self {
__self1, __self2 = tmp1517, V3952
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symdifference), tmp1517, V3952)


} else {
tmp1518 := GHd(__e, V3951)

tmp1519 := GTl(__e, V3951)

tmp1520 := Call2(__e, SymFn(symdifference), tmp1519, V3952)


return GCons(__e, tmp1518, tmp1520)


}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to find the difference with a non-list\n"))
}


}


panic("unreachable")
})
return __self
})()

tmp1528 := Call2(__e, ns2_1set, symdifference, tmp1516)


_ = tmp1528

tmp1529 := MakeNative2(func(__e *ControlFlow, V3953, V3954 Obj) Obj {
return V3954
panic("unreachable")
})

tmp1530 := Call2(__e, ns2_1set, symdo, tmp1529)


_ = tmp1530

tmp1531 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V3966 := __self1
_ = V3966
V3967 := __self2
_ = V3967
tmp1541 := GEq(__e, Nil, V3967)

if True == tmp1541 {
return False
} else {
tmp1539 := GConsP(__e, V3967)

var ifres1536 Obj

if True == tmp1539 {
tmp1537 := GHd(__e, V3967)

tmp1538 := GEq(__e, V3966, tmp1537)

ifres1536 = tmp1538


} else {
ifres1536 = False


}

if True == ifres1536 {
return True
} else {
tmp1534 := GConsP(__e, V3967)

if True == tmp1534 {
tmp1532 := GTl(__e, V3967)

if SymFn(symelement_2) == __self {
__self1, __self2 = V3966, tmp1532
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symelement_2), V3966, tmp1532)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to find an element in a non-list\n"))
}


}


}


panic("unreachable")
})
return __self
})()

tmp1542 := Call2(__e, ns2_1set, symelement_2, tmp1531)


_ = tmp1542

tmp1543 := MakeNative1(func(__e *ControlFlow, V3970 Obj) Obj {
tmp1545 := GEq(__e, Nil, V3970)

if True == tmp1545 {
return True
} else {
return False
}


panic("unreachable")
})

tmp1546 := Call2(__e, ns2_1set, symempty_2, tmp1543)


_ = tmp1546

tmp1547 := MakeNative2(func(__e *ControlFlow, V3971, V3972 Obj) Obj {
tmp1548 := Call1(__e, V3971, V3972)


return TailCall3(__e, SymFn(symshen_4fix_1help), V3971, V3972, tmp1548)


panic("unreachable")
})

tmp1549 := Call2(__e, ns2_1set, symfix, tmp1547)


_ = tmp1549

tmp1550 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V3978 := __self1
_ = V3978
V3979 := __self2
_ = V3979
V3980 := __self3
_ = V3980
tmp1553 := GEq(__e, V3979, V3980)

if True == tmp1553 {
return V3980
} else {
tmp1551 := Call1(__e, V3978, V3980)


if SymFn(symshen_4fix_1help) == __self {
__self1, __self2, __self3 = V3978, V3980, tmp1551
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4fix_1help), V3978, V3980, tmp1551)


}


panic("unreachable")
})
return __self
})()

tmp1554 := Call2(__e, ns2_1set, symshen_4fix_1help, tmp1550)


_ = tmp1554

tmp1555 := MakeNative4(func(__e *ControlFlow, V3981, V3982, V3983, V3984 Obj) Obj {
tmp1556 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall2(__e, SymFn(symshen_4_5_1dict), V3984, V3981)
panic("unreachable")
})

tmp1557 := MakeNative1(func(__e *ControlFlow, Z3986 Obj) Obj {
return Nil
panic("unreachable")
})

tmp1558 := Call2(__e, try_1catch, tmp1556, tmp1557)


let__1181 := tmp1558
_ = let__1181

tmp1559 := Call3(__e, SymFn(symshen_4assoc_1set), V3982, V3983, let__1181)


let__1182 := tmp1559
_ = let__1182

tmp1560 := Call3(__e, SymFn(symshen_4dict_1_6), V3984, V3981, let__1182)


let__1183 := tmp1560
_ = let__1183

return V3983


panic("unreachable")
})

tmp1561 := Call2(__e, ns2_1set, symput, tmp1555)


_ = tmp1561

tmp1562 := MakeNative3(func(__e *ControlFlow, V3989, V3990, V3991 Obj) Obj {
tmp1563 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall2(__e, SymFn(symshen_4_5_1dict), V3991, V3989)
panic("unreachable")
})

tmp1564 := MakeNative1(func(__e *ControlFlow, Z3993 Obj) Obj {
return Nil
panic("unreachable")
})

tmp1565 := Call2(__e, try_1catch, tmp1563, tmp1564)


let__1184 := tmp1565
_ = let__1184

tmp1566 := Call2(__e, SymFn(symshen_4assoc_1rm), V3990, let__1184)


let__1185 := tmp1566
_ = let__1185

tmp1567 := Call3(__e, SymFn(symshen_4dict_1_6), V3991, V3989, let__1185)


let__1186 := tmp1567
_ = let__1186

return V3989


panic("unreachable")
})

tmp1568 := Call2(__e, ns2_1set, symunput, tmp1562)


_ = tmp1568

tmp1569 := MakeNative3(func(__e *ControlFlow, V3996, V3997, V3998 Obj) Obj {
tmp1570 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall2(__e, SymFn(symshen_4_5_1dict), V3998, V3996)
panic("unreachable")
})

tmp1571 := MakeNative1(func(__e *ControlFlow, Z4000 Obj) Obj {
tmp1572 := Call3(__e, SymFn(symshen_4app), V3997, MakeString("\n"), symshen_4s)


tmp1574 := Call3(__e, SymFn(symshen_4app), V3996, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" has no attributes: "), tmp1572), symshen_4a)


return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp1574)


panic("unreachable")
})

tmp1575 := Call2(__e, try_1catch, tmp1570, tmp1571)


let__1187 := tmp1575
_ = let__1187

tmp1576 := Call2(__e, SymFn(symassoc), V3997, let__1187)


let__1188 := tmp1576
_ = let__1188

tmp1582 := Call1(__e, SymFn(symempty_2), let__1188)


if True == tmp1582 {
tmp1577 := Call3(__e, SymFn(symshen_4app), V3996, MakeString("\n"), symshen_4s)


tmp1579 := Call3(__e, SymFn(symshen_4app), V3997, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" not found for "), tmp1577), symshen_4s)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("attribute "), tmp1579))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("attribute "), tmp1579)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


} else {
return GTl(__e, let__1188)
}


panic("unreachable")
})

tmp1583 := Call2(__e, ns2_1set, symget, tmp1569)


_ = tmp1583

tmp1584 := MakeNative2(func(__e *ControlFlow, V4002, V4003 Obj) Obj {
tmp1585 := Call1(__e, SymFn(symshen_4hashkey), V4002)


tmp1586 := Call2(__e, SymFn(symshen_4mod), tmp1585, V4003)


let__1189 := tmp1586
_ = let__1189

tmp1588 := GEq(__e, let__1189, MakeInteger(0))

if True == tmp1588 {
return MakeInteger(1)
} else {
return let__1189
}


panic("unreachable")
})

tmp1589 := Call2(__e, ns2_1set, symhash, tmp1584)


_ = tmp1589

tmp1590 := MakeNative1(func(__e *ControlFlow, V4005 Obj) Obj {
tmp1591 := MakeNative1(func(__e *ControlFlow, Z4007 Obj) Obj {
return PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, Z4007)
panic("unreachable")
})

tmp1592 := Call1(__e, SymFn(symexplode), V4005)


tmp1593 := Call2(__e, SymFn(symmap), tmp1591, tmp1592)


let__1190 := tmp1593
_ = let__1190

return TailCall2(__e, SymFn(symshen_4prodbutzero), let__1190, MakeInteger(1))


panic("unreachable")
})

tmp1594 := Call2(__e, ns2_1set, symshen_4hashkey, tmp1590)


_ = tmp1594

tmp1595 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4008 := __self1
_ = V4008
V4009 := __self2
_ = V4009
tmp1613 := GEq(__e, Nil, V4008)

if True == tmp1613 {
return V4009
} else {
tmp1611 := GConsP(__e, V4008)

var ifres1608 Obj

if True == tmp1611 {
tmp1609 := GHd(__e, V4008)

tmp1610 := GEq(__e, MakeInteger(0), tmp1609)

ifres1608 = tmp1610


} else {
ifres1608 = False


}

if True == ifres1608 {
tmp1596 := GTl(__e, V4008)

if SymFn(symshen_4prodbutzero) == __self {
__self1, __self2 = tmp1596, V4009
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4prodbutzero), tmp1596, V4009)


} else {
tmp1606 := GConsP(__e, V4008)

if True == tmp1606 {
if True == GGt(__e, V4009, MakeNumber(1e+10)) {
tmp1597 := GTl(__e, V4008)

tmp1598 := GHd(__e, V4008)

if SymFn(symshen_4prodbutzero) == __self {
__self1, __self2 = tmp1597, GAdd(__e, V4009, tmp1598)
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4prodbutzero), tmp1597, GAdd(__e, V4009, tmp1598))


} else {
tmp1600 := GTl(__e, V4008)

tmp1601 := GHd(__e, V4008)

if SymFn(symshen_4prodbutzero) == __self {
__self1, __self2 = tmp1600, GMul(__e, V4009, tmp1601)
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4prodbutzero), tmp1600, GMul(__e, V4009, tmp1601))


}


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4prodbutzero)
}


}


}


panic("unreachable")
})
return __self
})()

tmp1614 := Call2(__e, ns2_1set, symshen_4prodbutzero, tmp1595)


_ = tmp1614

tmp1615 := MakeNative2(func(__e *ControlFlow, V4010, V4011 Obj) Obj {
tmp1616 := GCons(__e, V4011, Nil)

tmp1617 := Call2(__e, SymFn(symshen_4multiples), V4010, tmp1616)


return TailCall2(__e, SymFn(symshen_4modh), V4010, tmp1617)


panic("unreachable")
})

tmp1618 := Call2(__e, ns2_1set, symshen_4mod, tmp1615)


_ = tmp1618

tmp1619 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4016 := __self1
_ = V4016
V4017 := __self2
_ = V4017
tmp1629 := GConsP(__e, V4017)

var ifres1626 Obj

if True == tmp1629 {
tmp1627 := GHd(__e, V4017)

ifres1626 = GGt(__e, tmp1627, V4016)


} else {
ifres1626 = False


}

if True == ifres1626 {
return GTl(__e, V4017)
} else {
tmp1624 := GConsP(__e, V4017)

if True == tmp1624 {
tmp1620 := GHd(__e, V4017)

tmp1622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(GMul(__e, MakeInteger(2), tmp1620), V4017)
}
__typedArg0 := GMul(__e, MakeInteger(2), tmp1620)
__typedArg1 := V4017
return Call2(__e, SymFn(symcons), __typedArg0, __typedArg1)
})()

if SymFn(symshen_4multiples) == __self {
__self1, __self2 = V4016, tmp1622
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4multiples), V4016, tmp1622)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.multiples"))
}


}


panic("unreachable")
})
return __self
})()

tmp1630 := Call2(__e, ns2_1set, symshen_4multiples, tmp1619)


_ = tmp1630

tmp1631 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4024 := __self1
_ = V4024
V4025 := __self2
_ = V4025
tmp1648 := GEq(__e, MakeInteger(0), V4024)

if True == tmp1648 {
return MakeInteger(0)
} else {
tmp1646 := GEq(__e, Nil, V4025)

if True == tmp1646 {
return V4024
} else {
tmp1644 := GConsP(__e, V4025)

var ifres1641 Obj

if True == tmp1644 {
tmp1642 := GHd(__e, V4025)

ifres1641 = GGt(__e, tmp1642, V4024)


} else {
ifres1641 = False


}

if True == ifres1641 {
tmp1634 := GTl(__e, V4025)

tmp1635 := Call1(__e, SymFn(symempty_2), tmp1634)


if True == tmp1635 {
return V4024
} else {
tmp1632 := GTl(__e, V4025)

if SymFn(symshen_4modh) == __self {
__self1, __self2 = V4024, tmp1632
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4modh), V4024, tmp1632)


}


} else {
tmp1639 := GConsP(__e, V4025)

if True == tmp1639 {
tmp1636 := GHd(__e, V4025)

if SymFn(symshen_4modh) == __self {
__self1, __self2 = GSub(__e, V4024, tmp1636), V4025
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4modh), GSub(__e, V4024, tmp1636), V4025)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.modh"))
}


}


}


}


panic("unreachable")
})
return __self
})()

tmp1649 := Call2(__e, ns2_1set, symshen_4modh, tmp1631)


_ = tmp1649

tmp1650 := MakeNative1(func(__e *ControlFlow, V4028 Obj) Obj {
tmp1657 := GEq(__e, Nil, V4028)

if True == tmp1657 {
return MakeInteger(0)
} else {
tmp1655 := GConsP(__e, V4028)

if True == tmp1655 {
tmp1651 := GHd(__e, V4028)

tmp1652 := GTl(__e, V4028)

tmp1653 := Call1(__e, SymFn(symsum), tmp1652)


return GAdd(__e, tmp1651, tmp1653)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to sum a non-list\n"))
}


}


panic("unreachable")
})

tmp1658 := Call2(__e, ns2_1set, symsum, tmp1650)


_ = tmp1658

tmp1659 := MakeNative1(func(__e *ControlFlow, V4033 Obj) Obj {
tmp1661 := GConsP(__e, V4033)

if True == tmp1661 {
return GHd(__e, V4033)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("head expects a non-empty list\n"))
}


panic("unreachable")
})

tmp1662 := Call2(__e, ns2_1set, symhead, tmp1659)


_ = tmp1662

tmp1663 := MakeNative1(func(__e *ControlFlow, V4038 Obj) Obj {
tmp1665 := GConsP(__e, V4038)

if True == tmp1665 {
return GTl(__e, V4038)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("tail expects a non-empty list\n"))
}


panic("unreachable")
})

tmp1666 := Call2(__e, ns2_1set, symtail, tmp1663)


_ = tmp1666

tmp1667 := MakeNative1(func(__e *ControlFlow, V4039 Obj) Obj {
return PrimGuard2(__e, sympos, PrimPos, V4039, MakeInteger(0))
panic("unreachable")
})

tmp1668 := Call2(__e, ns2_1set, symhdstr, tmp1667)


_ = tmp1668

tmp1669 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4046 := __self1
_ = V4046
V4047 := __self2
_ = V4047
tmp1680 := GEq(__e, Nil, V4046)

if True == tmp1680 {
return Nil
} else {
tmp1678 := GConsP(__e, V4046)

if True == tmp1678 {
tmp1675 := GHd(__e, V4046)

tmp1676 := Call2(__e, SymFn(symelement_2), tmp1675, V4047)


if True == tmp1676 {
tmp1670 := GHd(__e, V4046)

tmp1671 := GTl(__e, V4046)

tmp1672 := Call2(__e, SymFn(symintersection), tmp1671, V4047)


return GCons(__e, tmp1670, tmp1672)


} else {
tmp1673 := GTl(__e, V4046)

if SymFn(symintersection) == __self {
__self1, __self2 = tmp1673, V4047
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symintersection), tmp1673, V4047)


}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to find the intersection with a non-list\n"))
}


}


panic("unreachable")
})
return __self
})()

tmp1681 := Call2(__e, ns2_1set, symintersection, tmp1669)


_ = tmp1681

tmp1682 := MakeNative1(func(__e *ControlFlow, V4048 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4reverse_1help), V4048, Nil)
panic("unreachable")
})

tmp1683 := Call2(__e, ns2_1set, symreverse, tmp1682)


_ = tmp1683

tmp1684 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4053 := __self1
_ = V4053
V4054 := __self2
_ = V4054
tmp1691 := GEq(__e, Nil, V4053)

if True == tmp1691 {
return V4054
} else {
tmp1689 := GConsP(__e, V4053)

if True == tmp1689 {
tmp1685 := GTl(__e, V4053)

tmp1686 := GHd(__e, V4053)

tmp1687 := GCons(__e, tmp1686, V4054)

if SymFn(symshen_4reverse_1help) == __self {
__self1, __self2 = tmp1685, tmp1687
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4reverse_1help), tmp1685, tmp1687)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to reverse a non-list\n"))
}


}


panic("unreachable")
})
return __self
})()

tmp1692 := Call2(__e, ns2_1set, symshen_4reverse_1help, tmp1684)


_ = tmp1692

tmp1693 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4059 := __self1
_ = V4059
V4060 := __self2
_ = V4060
tmp1704 := GEq(__e, Nil, V4059)

if True == tmp1704 {
return V4060
} else {
tmp1702 := GConsP(__e, V4059)

if True == tmp1702 {
tmp1699 := GHd(__e, V4059)

tmp1700 := Call2(__e, SymFn(symelement_2), tmp1699, V4060)


if True == tmp1700 {
tmp1694 := GTl(__e, V4059)

if SymFn(symunion) == __self {
__self1, __self2 = tmp1694, V4060
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symunion), tmp1694, V4060)


} else {
tmp1695 := GHd(__e, V4059)

tmp1696 := GTl(__e, V4059)

tmp1697 := Call2(__e, SymFn(symunion), tmp1696, V4060)


return GCons(__e, tmp1695, tmp1697)


}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to find the union with a non-list\n"))
}


}


panic("unreachable")
})
return __self
})()

tmp1705 := Call2(__e, ns2_1set, symunion, tmp1693)


_ = tmp1705

tmp1706 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V4061 := __self1
_ = V4061
tmp1707 := Call1(__e, SymFn(symshen_4proc_1nl), V4061)


tmp1708 := Call0(__e, SymFn(symstoutput))


tmp1709 := Call2(__e, SymFn(sympr), tmp1707, tmp1708)


let__1191 := tmp1709
_ = let__1191

tmp1710 := Call0(__e, SymFn(symstoutput))


tmp1711 := Call2(__e, SymFn(sympr), MakeString(" (y/n) "), tmp1710)


let__1192 := tmp1711
_ = let__1192

tmp1712 := Call0(__e, SymFn(symstinput))


tmp1713 := Call1(__e, SymFn(symread), tmp1712)


tmp1714 := Call3(__e, SymFn(symshen_4app), tmp1713, MakeString(""), symshen_4s)


let__1193 := tmp1714
_ = let__1193

tmp1720 := GEq(__e, MakeString("y"), let__1193)

if True == tmp1720 {
return True
} else {
tmp1718 := GEq(__e, MakeString("n"), let__1193)

if True == tmp1718 {
return False
} else {
tmp1715 := Call0(__e, SymFn(symstoutput))


tmp1716 := Call2(__e, SymFn(sympr), MakeString("please answer y or n\n"), tmp1715)


_ = tmp1716

if SymFn(symy_1or_1n_2) == __self {
__self1 = V4061
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symy_1or_1n_2), V4061)


}


}


panic("unreachable")
})
return __self
})()

tmp1721 := Call2(__e, ns2_1set, symy_1or_1n_2, tmp1706)


_ = tmp1721

tmp1722 := MakeNative1(func(__e *ControlFlow, V4065 Obj) Obj {
if True == V4065 {
return False
} else {
return True
}
panic("unreachable")
})

tmp1724 := Call2(__e, ns2_1set, symnot, tmp1722)


_ = tmp1724

tmp1725 := MakeNative0(func(__e *ControlFlow) Obj {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString(""))
panic("unreachable")
})

tmp1726 := Call2(__e, ns2_1set, symabort, tmp1725)


_ = tmp1726

tmp1727 := MakeNative3(func(__e *ControlFlow, V4071, V4072, V4073 Obj) Obj {
tmp1735 := GEq(__e, V4072, V4073)

if True == tmp1735 {
return V4071
} else {
tmp1733 := GConsP(__e, V4073)

if True == tmp1733 {
tmp1728 := GHd(__e, V4073)

tmp1729 := Call3(__e, SymFn(symsubst), V4071, V4072, tmp1728)


tmp1730 := GTl(__e, V4073)

tmp1731 := Call3(__e, SymFn(symsubst), V4071, V4072, tmp1730)


return GCons(__e, tmp1729, tmp1731)


} else {
return V4073
}


}


panic("unreachable")
})

tmp1736 := Call2(__e, ns2_1set, symsubst, tmp1727)


_ = tmp1736

tmp1737 := MakeNative1(func(__e *ControlFlow, V4074 Obj) Obj {
tmp1738 := Call3(__e, SymFn(symshen_4app), V4074, MakeString(""), symshen_4a)


return TailCall1(__e, SymFn(symshen_4explode_1h), tmp1738)


panic("unreachable")
})

tmp1739 := Call2(__e, ns2_1set, symexplode, tmp1737)


_ = tmp1739

tmp1740 := MakeNative1(func(__e *ControlFlow, V4077 Obj) Obj {
tmp1747 := GEq(__e, MakeString(""), V4077)

if True == tmp1747 {
return Nil
} else {
tmp1745 := Call1(__e, SymFn(symshen_4_7string_2), V4077)


if True == tmp1745 {
tmp1741 := Call1(__e, SymFn(symhdstr), V4077)


tmp1743 := Call1(__e, SymFn(symshen_4explode_1h), PrimGuard1(__e, symtlstr, PrimTailString, V4077))


return GCons(__e, tmp1741, tmp1743)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in explode-h"))
}


}


panic("unreachable")
})

tmp1748 := Call2(__e, ns2_1set, symshen_4explode_1h, tmp1740)


_ = tmp1748

tmp1749 := MakeNative1(func(__e *ControlFlow, V4078 Obj) Obj {
tmp1752 := GEq(__e, V4078, MakeString(""))

var ifres1750 Obj

if True == tmp1752 {
ifres1750 = MakeString("")


} else {
tmp1751 := Call3(__e, SymFn(symshen_4app), V4078, MakeString("/"), symshen_4a)


ifres1750 = tmp1751


}

return PrimGuard2(__e, symset, PrimSet, sym_dhome_1directory_d, ifres1750)


panic("unreachable")
})

tmp1753 := Call2(__e, ns2_1set, symcd, tmp1749)


_ = tmp1753

tmp1754 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4079 := __self1
_ = V4079
V4080 := __self2
_ = V4080
tmp1761 := GEq(__e, Nil, V4080)

if True == tmp1761 {
return True
} else {
tmp1759 := GConsP(__e, V4080)

if True == tmp1759 {
tmp1755 := GHd(__e, V4080)

tmp1756 := Call1(__e, V4079, tmp1755)


let__1194 := tmp1756
_ = let__1194

tmp1757 := GTl(__e, V4080)

if SymFn(symshen_4for_1each) == __self {
__self1, __self2 = V4079, tmp1757
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4for_1each), V4079, tmp1757)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4for_1each)
}


}


panic("unreachable")
})
return __self
})()

tmp1762 := Call2(__e, ns2_1set, symshen_4for_1each, tmp1754)


_ = tmp1762

tmp1763 := MakeNative2(func(__e *ControlFlow, V4082, V4083 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4map_1h), V4082, V4083, Nil)
panic("unreachable")
})

tmp1764 := Call2(__e, ns2_1set, symmap, tmp1763)


_ = tmp1764

tmp1765 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V4084 := __self1
_ = V4084
V4085 := __self2
_ = V4085
V4086 := __self3
_ = V4086
tmp1773 := GEq(__e, Nil, V4085)

if True == tmp1773 {
return TailCall1(__e, SymFn(symreverse), V4086)
} else {
tmp1771 := GConsP(__e, V4085)

if True == tmp1771 {
tmp1766 := GTl(__e, V4085)

tmp1767 := GHd(__e, V4085)

tmp1768 := Call1(__e, V4084, tmp1767)


tmp1769 := GCons(__e, tmp1768, V4086)

if SymFn(symshen_4map_1h) == __self {
__self1, __self2, __self3 = V4084, tmp1766, tmp1769
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4map_1h), V4084, tmp1766, tmp1769)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4map_1h)
}


}


panic("unreachable")
})
return __self
})()

tmp1774 := Call2(__e, ns2_1set, symshen_4map_1h, tmp1765)


_ = tmp1774

tmp1775 := MakeNative1(func(__e *ControlFlow, V4087 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4length_1h), V4087, MakeInteger(0))
panic("unreachable")
})

tmp1776 := Call2(__e, ns2_1set, symlength, tmp1775)


_ = tmp1776

tmp1777 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4092 := __self1
_ = V4092
V4093 := __self2
_ = V4093
tmp1781 := GEq(__e, Nil, V4092)

if True == tmp1781 {
return V4093
} else {
tmp1778 := GTl(__e, V4092)

if SymFn(symshen_4length_1h) == __self {
__self1, __self2 = tmp1778, GAdd(__e, V4093, MakeInteger(1))
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4length_1h), tmp1778, GAdd(__e, V4093, MakeInteger(1)))


}


panic("unreachable")
})
return __self
})()

tmp1782 := Call2(__e, ns2_1set, symshen_4length_1h, tmp1777)


_ = tmp1782

tmp1783 := MakeNative2(func(__e *ControlFlow, V4099, V4100 Obj) Obj {
tmp1791 := GEq(__e, V4099, V4100)

if True == tmp1791 {
return MakeInteger(1)
} else {
tmp1789 := GConsP(__e, V4100)

if True == tmp1789 {
tmp1784 := GHd(__e, V4100)

tmp1785 := Call2(__e, SymFn(symoccurrences), V4099, tmp1784)


tmp1786 := GTl(__e, V4100)

tmp1787 := Call2(__e, SymFn(symoccurrences), V4099, tmp1786)


return GAdd(__e, tmp1785, tmp1787)


} else {
return MakeInteger(0)
}


}


panic("unreachable")
})

tmp1792 := Call2(__e, ns2_1set, symoccurrences, tmp1783)


_ = tmp1792

tmp1793 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4105 := __self1
_ = V4105
V4106 := __self2
_ = V4106
tmp1805 := GEq(__e, MakeInteger(1), V4105)

var ifres1803 Obj

if True == tmp1805 {
tmp1804 := GConsP(__e, V4106)

ifres1803 = tmp1804


} else {
ifres1803 = False


}

if True == ifres1803 {
return GHd(__e, V4106)
} else {
tmp1801 := GConsP(__e, V4106)

if True == tmp1801 {
tmp1794 := GSub(__e, V4105, MakeInteger(1))

tmp1795 := GTl(__e, V4106)

if SymFn(symnth) == __self {
__self1, __self2 = tmp1794, tmp1795
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symnth), tmp1794, tmp1795)


} else {
tmp1796 := Call3(__e, SymFn(symshen_4app), V4106, MakeString("\n"), symshen_4a)


tmp1798 := Call3(__e, SymFn(symshen_4app), V4105, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(", "), tmp1796), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("nth applied to "), tmp1798))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("nth applied to "), tmp1798)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


}


panic("unreachable")
})
return __self
})()

tmp1806 := Call2(__e, ns2_1set, symnth, tmp1793)


_ = tmp1806

tmp1807 := MakeNative1(func(__e *ControlFlow, V4107 Obj) Obj {
tmp1811 := PrimGuard1(__e, symnumber_2, PrimIsNumber, V4107)

if True == tmp1811 {
tmp1808 := Call1(__e, SymFn(symshen_4abs), V4107)


let__1195 := tmp1808
_ = let__1195

tmp1809 := Call2(__e, SymFn(symshen_4magless), let__1195, MakeInteger(1))


return TailCall2(__e, SymFn(symshen_4integer_1test_2), let__1195, tmp1809)


} else {
return False
}


panic("unreachable")
})

tmp1812 := Call2(__e, ns2_1set, syminteger_2, tmp1807)


_ = tmp1812

tmp1813 := MakeNative1(func(__e *ControlFlow, V4109 Obj) Obj {
if True == GGt(__e, V4109, MakeInteger(0)) {
return V4109
} else {
return GSub(__e, MakeInteger(0), V4109)
}


panic("unreachable")
})

tmp1816 := Call2(__e, ns2_1set, symshen_4abs, tmp1813)


_ = tmp1816

tmp1817 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4110 := __self1
_ = V4110
V4111 := __self2
_ = V4111
let__1196 := GMul(__e, V4111, MakeInteger(2))
_ = let__1196

if True == GGt(__e, let__1196, V4110) {
return V4111
} else {
if SymFn(symshen_4magless) == __self {
__self1, __self2 = V4110, let__1196
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4magless), V4110, let__1196)
}


panic("unreachable")
})
return __self
})()

tmp1821 := Call2(__e, ns2_1set, symshen_4magless, tmp1817)


_ = tmp1821

tmp1822 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4116 := __self1
_ = V4116
V4117 := __self2
_ = V4117
tmp1829 := GEq(__e, MakeInteger(0), V4116)

if True == tmp1829 {
return True
} else {
if True == GGt(__e, MakeInteger(1), V4116) {
return False
} else {
let__1197 := GSub(__e, V4116, V4117)
_ = let__1197

if True == GGt(__e, MakeInteger(0), let__1197) {
return PrimGuard1(__e, syminteger_2, PrimIsInteger, V4116)
} else {
if SymFn(symshen_4integer_1test_2) == __self {
__self1, __self2 = let__1197, V4117
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4integer_1test_2), let__1197, V4117)
}


}


}


panic("unreachable")
})
return __self
})()

tmp1830 := Call2(__e, ns2_1set, symshen_4integer_1test_2, tmp1822)


_ = tmp1830

tmp1831 := MakeNative2(func(__e *ControlFlow, V4125, V4126 Obj) Obj {
tmp1839 := GEq(__e, Nil, V4126)

if True == tmp1839 {
return Nil
} else {
tmp1837 := GConsP(__e, V4126)

if True == tmp1837 {
tmp1832 := GHd(__e, V4126)

tmp1833 := Call1(__e, V4125, tmp1832)


tmp1834 := GTl(__e, V4126)

tmp1835 := Call2(__e, SymFn(symmapcan), V4125, tmp1834)


return TailCall2(__e, SymFn(symappend), tmp1833, tmp1835)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to mapcan over a non-list\n"))
}


}


panic("unreachable")
})

tmp1840 := Call2(__e, ns2_1set, symmapcan, tmp1831)


_ = tmp1840

tmp1841 := MakeNative2(func(__e *ControlFlow, V4132, V4133 Obj) Obj {
tmp1843 := GEq(__e, V4132, V4133)

if True == tmp1843 {
return True
} else {
return False
}


panic("unreachable")
})

tmp1844 := Call2(__e, ns2_1set, sym_a_a, tmp1841)


_ = tmp1844

tmp1845 := MakeNative1(func(__e *ControlFlow, V4134 Obj) Obj {
tmp1852 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V4134)

if True == tmp1852 {
tmp1846 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, V4134)
panic("unreachable")
})

tmp1847 := MakeNative1(func(__e *ControlFlow, Z4136 Obj) Obj {
return symshen_4this_1symbol_1is_1unbound
panic("unreachable")
})

tmp1848 := Call2(__e, try_1catch, tmp1846, tmp1847)


let__1198 := tmp1848
_ = let__1198

tmp1850 := GEq(__e, let__1198, symshen_4this_1symbol_1is_1unbound)

if True == tmp1850 {
return False
} else {
return True
}


} else {
return False
}


panic("unreachable")
})

tmp1853 := Call2(__e, ns2_1set, symbound_2, tmp1845)


_ = tmp1853

tmp1854 := MakeNative1(func(__e *ControlFlow, V4137 Obj) Obj {
tmp1860 := GEq(__e, MakeString(""), V4137)

if True == tmp1860 {
return Nil
} else {
tmp1855 := PrimGuard2(__e, sympos, PrimPos, V4137, MakeInteger(0))

tmp1856 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp1855)

tmp1858 := Call1(__e, SymFn(symshen_4string_1_6bytes), PrimGuard1(__e, symtlstr, PrimTailString, V4137))


return GCons(__e, tmp1856, tmp1858)


}


panic("unreachable")
})

tmp1861 := Call2(__e, ns2_1set, symshen_4string_1_6bytes, tmp1854)


_ = tmp1861

tmp1862 := MakeNative1(func(__e *ControlFlow, V4138 Obj) Obj {
if True == GLt(__e, V4138, MakeInteger(0)) {
return GValue(__e, symshen_4_dmaxinferences_d)
} else {
tmp1864 := PrimGuard1(__e, syminteger_2, PrimIsInteger, V4138)

if True == tmp1864 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dmaxinferences_d, V4138)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("maxinferences expects an integer value\n"))
}


}


panic("unreachable")
})

tmp1867 := Call2(__e, ns2_1set, symmaxinferences, tmp1862)


_ = tmp1867

tmp1868 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dinfs_d)
panic("unreachable")
})

tmp1869 := Call2(__e, ns2_1set, syminferences, tmp1868)


_ = tmp1869

tmp1870 := MakeNative1(func(__e *ControlFlow, V4139 Obj) Obj {
return V4139
panic("unreachable")
})

tmp1871 := Call2(__e, ns2_1set, symprotect, tmp1870)


_ = tmp1871

tmp1872 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dsterror_d)
panic("unreachable")
})

tmp1873 := Call2(__e, ns2_1set, symsterror, tmp1872)


_ = tmp1873

tmp1874 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dstoutput_d)
panic("unreachable")
})

tmp1875 := Call2(__e, ns2_1set, symstoutput, tmp1874)


_ = tmp1875

tmp1876 := MakeNative1(func(__e *ControlFlow, V4140 Obj) Obj {
tmp1877 := PrimGuard1(__e, symintern, PrimIntern, V4140)

let__1199 := tmp1877
_ = let__1199

tmp1881 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, let__1199)

if True == tmp1881 {
return let__1199
} else {
tmp1878 := Call3(__e, SymFn(symshen_4app), V4140, MakeString(" to a symbol"), symshen_4s)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("cannot intern "), tmp1878))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("cannot intern "), tmp1878)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


panic("unreachable")
})

tmp1882 := Call2(__e, ns2_1set, symstring_1_6symbol, tmp1876)


_ = tmp1882

tmp1883 := MakeNative1(func(__e *ControlFlow, V4144 Obj) Obj {
tmp1887 := GEq(__e, sym_7, V4144)

if True == tmp1887 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_doptimise_d, True)
} else {
tmp1885 := GEq(__e, sym_1, V4144)

if True == tmp1885 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_doptimise_d, False)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("optimise expects a + or a -.\n"))
}


}


panic("unreachable")
})

tmp1888 := Call2(__e, ns2_1set, symoptimise, tmp1883)


_ = tmp1888

tmp1889 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dos_d)
panic("unreachable")
})

tmp1890 := Call2(__e, ns2_1set, symos, tmp1889)


_ = tmp1890

tmp1891 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dlanguage_d)
panic("unreachable")
})

tmp1892 := Call2(__e, ns2_1set, symlanguage, tmp1891)


_ = tmp1892

tmp1893 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dversion_d)
panic("unreachable")
})

tmp1894 := Call2(__e, ns2_1set, symversion, tmp1893)


_ = tmp1894

tmp1895 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dport_d)
panic("unreachable")
})

tmp1896 := Call2(__e, ns2_1set, symport, tmp1895)


_ = tmp1896

tmp1897 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dporters_d)
panic("unreachable")
})

tmp1898 := Call2(__e, ns2_1set, symporters, tmp1897)


_ = tmp1898

tmp1899 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dimplementation_d)
panic("unreachable")
})

tmp1900 := Call2(__e, ns2_1set, symimplementation, tmp1899)


_ = tmp1900

tmp1901 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_drelease_d)
panic("unreachable")
})

tmp1902 := Call2(__e, ns2_1set, symrelease, tmp1901)


_ = tmp1902

tmp1903 := MakeNative1(func(__e *ControlFlow, V4145 Obj) Obj {
tmp1908 := GEq(__e, symnull, V4145)

if True == tmp1908 {
return True
} else {
tmp1904 := MakeNative0(func(__e *ControlFlow) Obj {
tmp1905 := Call1(__e, SymFn(symexternal), V4145)


_ = tmp1905

return True


panic("unreachable")
})

tmp1906 := MakeNative1(func(__e *ControlFlow, Z4146 Obj) Obj {
return False
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp1904, tmp1906)


}


panic("unreachable")
})

tmp1909 := Call2(__e, ns2_1set, sympackage_2, tmp1903)


_ = tmp1909

tmp1910 := MakeNative0(func(__e *ControlFlow) Obj {
return symshen_4fail_b
panic("unreachable")
})

tmp1911 := Call2(__e, ns2_1set, symfail, tmp1910)


_ = tmp1911

tmp1912 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_duserdefs_d)
panic("unreachable")
})

tmp1913 := Call2(__e, ns2_1set, symuserdefs, tmp1912)


_ = tmp1913

tmp1914 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_doptimise_d)
panic("unreachable")
})

tmp1915 := Call2(__e, ns2_1set, symoptimise_2, tmp1914)


_ = tmp1915

tmp1916 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dhush_d)
panic("unreachable")
})

tmp1917 := Call2(__e, ns2_1set, symhush_2, tmp1916)


_ = tmp1917

tmp1918 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dshen_1type_1theory_1enabled_2_d)
panic("unreachable")
})

tmp1919 := Call2(__e, ns2_1set, symsystem_1S_2, tmp1918)


_ = tmp1919

tmp1920 := MakeNative1(func(__e *ControlFlow, V4149 Obj) Obj {
tmp1924 := GEq(__e, sym_7, V4149)

if True == tmp1924 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dshen_1type_1theory_1enabled_2_d, True)
} else {
tmp1922 := GEq(__e, sym_1, V4149)

if True == tmp1922 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dshen_1type_1theory_1enabled_2_d, False)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("enable-type-theory expects a + or a -\n"))
}


}


panic("unreachable")
})

tmp1925 := Call2(__e, ns2_1set, symenable_1type_1theory, tmp1920)


_ = tmp1925

tmp1926 := MakeNative1(func(__e *ControlFlow, V4152 Obj) Obj {
tmp1930 := GEq(__e, sym_7, V4152)

if True == tmp1930 {
return PrimGuard2(__e, symset, PrimSet, sym_dhush_d, True)
} else {
tmp1928 := GEq(__e, sym_1, V4152)

if True == tmp1928 {
return PrimGuard2(__e, symset, PrimSet, sym_dhush_d, False)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("hush expects a + or a -\n"))
}


}


panic("unreachable")
})

tmp1931 := Call2(__e, ns2_1set, symhush, tmp1926)


_ = tmp1931

tmp1932 := MakeNative1(func(__e *ControlFlow, V4155 Obj) Obj {
tmp1936 := GEq(__e, sym_7, V4155)

if True == tmp1936 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dtc_d, True)
} else {
tmp1934 := GEq(__e, sym_1, V4155)

if True == tmp1934 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dtc_d, False)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("tc expects a + or -"))
}


}


panic("unreachable")
})

tmp1937 := Call2(__e, ns2_1set, symtc, tmp1932)


_ = tmp1937

tmp1938 := MakeNative1(func(__e *ControlFlow, V4156 Obj) Obj {
tmp1939 := GValue(__e, symshen_4_dsigf_d)

tmp1940 := Call2(__e, SymFn(symshen_4unassoc), V4156, tmp1939)


tmp1941 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp1940)

_ = tmp1941

return V4156


panic("unreachable")
})

tmp1942 := Call2(__e, ns2_1set, symdestroy, tmp1938)


_ = tmp1942

tmp1943 := MakeNative2(func(__e *ControlFlow, V4166, V4167 Obj) Obj {
tmp1959 := GEq(__e, Nil, V4167)

if True == tmp1959 {
return Nil
} else {
tmp1957 := GConsP(__e, V4167)

var ifres1950 Obj

if True == tmp1957 {
tmp1955 := GHd(__e, V4167)

tmp1956 := GConsP(__e, tmp1955)

var ifres1951 Obj

if True == tmp1956 {
tmp1952 := GHd(__e, V4167)

tmp1953 := GHd(__e, tmp1952)

tmp1954 := GEq(__e, V4166, tmp1953)

ifres1951 = tmp1954


} else {
ifres1951 = False


}

ifres1950 = ifres1951


} else {
ifres1950 = False


}

if True == ifres1950 {
return GTl(__e, V4167)
} else {
tmp1948 := GConsP(__e, V4167)

if True == tmp1948 {
tmp1944 := GHd(__e, V4167)

tmp1945 := GTl(__e, V4167)

tmp1946 := Call2(__e, SymFn(symshen_4unassoc), V4166, tmp1945)


return GCons(__e, tmp1944, tmp1946)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.unassoc"))
}


}


}


panic("unreachable")
})

tmp1960 := Call2(__e, ns2_1set, symshen_4unassoc, tmp1943)


_ = tmp1960

tmp1961 := MakeNative1(func(__e *ControlFlow, V4168 Obj) Obj {
tmp1965 := Call1(__e, SymFn(sympackage_2), V4168)


if True == tmp1965 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dpackage_d, V4168)
} else {
tmp1962 := Call3(__e, SymFn(symshen_4app), V4168, MakeString(" does not exist\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("package "), tmp1962))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("package "), tmp1962)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


panic("unreachable")
})

tmp1966 := Call2(__e, ns2_1set, symin_1package, tmp1961)


_ = tmp1966

tmp1967 := MakeNative2(func(__e *ControlFlow, V4169, V4170 Obj) Obj {
tmp1968 := PrimGuard2(__e, symopen, PrimOpenStream, V4169, symout)

let__1200 := tmp1968
_ = let__1200

tmp1971 := PrimGuard1(__e, symstring_2, PrimIsString, V4170)

var ifres1969 Obj

if True == tmp1971 {
ifres1969 = V4170


} else {
tmp1970 := Call3(__e, SymFn(symshen_4app), V4170, MakeString(""), symshen_4s)


ifres1969 = tmp1970


}

let__1201 := ifres1969
_ = let__1201

tmp1972 := Call2(__e, SymFn(sympr), let__1201, let__1200)


let__1202 := tmp1972
_ = let__1202

tmp1973 := PrimGuard1(__e, symclose, PrimCloseStream, let__1200)

let__1203 := tmp1973
_ = let__1203

return V4170


panic("unreachable")
})

tmp1974 := Call2(__e, ns2_1set, symwrite_1to_1file, tmp1967)


_ = tmp1974

tmp1975 := MakeNative0(func(__e *ControlFlow) Obj {
tmp1976 := Call1(__e, SymFn(symgensym), symshen_4t)


return TailCall1(__e, SymFn(symshen_4freshterm), tmp1976)


panic("unreachable")
})

tmp1977 := Call2(__e, ns2_1set, symfresh, tmp1975)


_ = tmp1977

tmp1978 := MakeNative2(func(__e *ControlFlow, V4175, V4176 Obj) Obj {
tmp1979 := GValue(__e, sym_dproperty_1vector_d)

tmp1980 := Call4(__e, SymFn(symput), V4175, symarity, V4176, tmp1979)


let__1204 := tmp1980
_ = let__1204

tmp1981 := Call1(__e, SymFn(symshen_4lambda_1entry), V4175)


let__1205 := tmp1981
_ = let__1205

tmp1986 := Call1(__e, SymFn(symempty_2), let__1205)


var ifres1982 Obj

if True == tmp1986 {
tmp1983 := GValue(__e, sym_dproperty_1vector_d)

tmp1984 := Call3(__e, SymFn(symunput), V4175, symshen_4lambda_1form, tmp1983)


ifres1982 = tmp1984


} else {
tmp1985 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), let__1205)


ifres1982 = tmp1985


}

let__1206 := ifres1982
_ = let__1206

return V4175


panic("unreachable")
})

tmp1987 := Call2(__e, ns2_1set, symupdate_1lambda_1table, tmp1978)


_ = tmp1987

tmp1988 := MakeNative2(func(__e *ControlFlow, V4182, V4183 Obj) Obj {
tmp2012 := GEq(__e, MakeInteger(0), V4183)

if True == tmp2012 {
tmp1989 := GValue(__e, symshen_4_dspecial_d)

tmp1990 := Call2(__e, SymFn(symremove), V4182, tmp1989)


tmp1991 := PrimGuard2(__e, symset, PrimSet, symshen_4_dspecial_d, tmp1990)

_ = tmp1991

tmp1992 := GValue(__e, symshen_4_dextraspecial_d)

tmp1993 := Call2(__e, SymFn(symremove), V4182, tmp1992)


tmp1994 := PrimGuard2(__e, symset, PrimSet, symshen_4_dextraspecial_d, tmp1993)

_ = tmp1994

return V4182


} else {
tmp2010 := GEq(__e, MakeInteger(1), V4183)

if True == tmp2010 {
tmp1995 := GValue(__e, symshen_4_dspecial_d)

tmp1996 := Call2(__e, SymFn(symadjoin), V4182, tmp1995)


tmp1997 := PrimGuard2(__e, symset, PrimSet, symshen_4_dspecial_d, tmp1996)

_ = tmp1997

tmp1998 := GValue(__e, symshen_4_dextraspecial_d)

tmp1999 := Call2(__e, SymFn(symremove), V4182, tmp1998)


tmp2000 := PrimGuard2(__e, symset, PrimSet, symshen_4_dextraspecial_d, tmp1999)

_ = tmp2000

return V4182


} else {
tmp2008 := GEq(__e, MakeInteger(2), V4183)

if True == tmp2008 {
tmp2001 := GValue(__e, symshen_4_dspecial_d)

tmp2002 := Call2(__e, SymFn(symremove), V4182, tmp2001)


tmp2003 := PrimGuard2(__e, symset, PrimSet, symshen_4_dspecial_d, tmp2002)

_ = tmp2003

tmp2004 := GValue(__e, symshen_4_dextraspecial_d)

tmp2005 := Call2(__e, SymFn(symadjoin), V4182, tmp2004)


tmp2006 := PrimGuard2(__e, symset, PrimSet, symshen_4_dextraspecial_d, tmp2005)

_ = tmp2006

return V4182


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("specialise requires values of 0, 1 or 2\n"))
}


}


}


panic("unreachable")
})

tmp2013 := Call2(__e, ns2_1set, symspecialise, tmp1988)


_ = tmp2013

tmp2014 := MakeNative1(func(__e *ControlFlow, V4184 Obj) Obj {
tmp2015 := GValue(__e, sym_dabsolute_d)

tmp2016 := GCons(__e, V4184, tmp2015)

return PrimGuard2(__e, symset, PrimSet, sym_dabsolute_d, tmp2016)


panic("unreachable")
})

tmp2017 := Call2(__e, ns2_1set, symabsolute, tmp2014)


_ = tmp2017

tmp2018 := MakeNative1(func(__e *ControlFlow, V4185 Obj) Obj {
tmp2019 := GValue(__e, sym_dabsolute_d)

tmp2020 := Call2(__e, SymFn(symremove), V4185, tmp2019)


return PrimGuard2(__e, symset, PrimSet, sym_dabsolute_d, tmp2020)


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symunabsolute, tmp2018)




panic("unreachable")
})

