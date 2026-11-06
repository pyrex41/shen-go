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
tmp1280 := GVectorP(__e, V3846)

if True == tmp1280 {
tmp1272 := MakeNative0(func(__e *ControlFlow) Obj {
return GVecGet(__e, V3846, MakeInteger(0))
panic("unreachable")
})

tmp1273 := MakeNative1(func(__e *ControlFlow, Z3848 Obj) Obj {
return MakeInteger(-1)
panic("unreachable")
})

tmp1274 := Call2(__e, try_1catch, tmp1272, tmp1273)


let__1167 := tmp1274
_ = let__1167

tmp1278 := PrimGuard1(__e, symnumber_2, PrimIsNumber, let__1167)

var ifres1275 Obj

if True == tmp1278 {
tmp1277 := GGe(__e, let__1167, MakeInteger(0))

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
return True
} else {
return False
}


} else {
return False
}


panic("unreachable")
})

tmp1281 := Call2(__e, ns2_1set, symvector_2, tmp1270)


_ = tmp1281

tmp1282 := MakeNative3(func(__e *ControlFlow, V3849, V3850, V3851 Obj) Obj {
tmp1284 := GEq(__e, V3850, MakeInteger(0))

if True == tmp1284 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("cannot access 0th element of a vector\n"))
} else {
return GVecSet(__e, V3849, V3850, V3851)
}


panic("unreachable")
})

tmp1285 := Call2(__e, ns2_1set, symvector_1_6, tmp1282)


_ = tmp1285

tmp1286 := MakeNative2(func(__e *ControlFlow, V3852, V3853 Obj) Obj {
tmp1292 := GEq(__e, V3853, MakeInteger(0))

if True == tmp1292 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("cannot access 0th element of a vector\n"))
} else {
tmp1287 := GVecGet(__e, V3852, V3853)

let__1168 := tmp1287
_ = let__1168

tmp1289 := Call0(__e, SymFn(symfail))


tmp1290 := GEq(__e, let__1168, tmp1289)

if True == tmp1290 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("vector element not found\n"))
} else {
return let__1168
}


}


panic("unreachable")
})

tmp1293 := Call2(__e, ns2_1set, sym_5_1vector, tmp1286)


_ = tmp1293

tmp1294 := MakeNative1(func(__e *ControlFlow, V3855 Obj) Obj {
tmp1298 := PrimGuard1(__e, syminteger_2, PrimIsInteger, V3855)

if True == tmp1298 {
if True == GGe(__e, V3855, MakeInteger(0)) {
return True
} else {
return False
}


} else {
return False
}


panic("unreachable")
})

tmp1299 := Call2(__e, ns2_1set, symshen_4posint_2, tmp1294)


_ = tmp1299

tmp1300 := MakeNative1(func(__e *ControlFlow, V3856 Obj) Obj {
return GVecGet(__e, V3856, MakeInteger(0))
panic("unreachable")
})

tmp1301 := Call2(__e, ns2_1set, symlimit, tmp1300)


_ = tmp1301

tmp1302 := MakeNative1(func(__e *ControlFlow, V3857 Obj) Obj {
tmp1332 := Call1(__e, SymFn(symboolean_2), V3857)


var ifres1317 Obj

if True == tmp1332 {
ifres1317 = True


} else {
tmp1331 := PrimGuard1(__e, symnumber_2, PrimIsNumber, V3857)

var ifres1319 Obj

if True == tmp1331 {
ifres1319 = True


} else {
tmp1330 := PrimGuard1(__e, symstring_2, PrimIsString, V3857)

var ifres1321 Obj

if True == tmp1330 {
ifres1321 = True


} else {
tmp1329 := GConsP(__e, V3857)

var ifres1323 Obj

if True == tmp1329 {
ifres1323 = True


} else {
tmp1328 := Call1(__e, SymFn(symempty_2), V3857)


var ifres1325 Obj

if True == tmp1328 {
ifres1325 = True


} else {
tmp1327 := Call1(__e, SymFn(symvector_2), V3857)


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
return False
} else {
tmp1307 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp1308 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp1309 := PrimGuard1(__e, symintern, PrimIntern, MakeString(","))

tmp1310 := GCons(__e, tmp1309, Nil)

tmp1311 := GCons(__e, tmp1308, tmp1310)

tmp1312 := GCons(__e, tmp1307, tmp1311)

tmp1313 := GCons(__e, sym_j, tmp1312)

tmp1314 := GCons(__e, sym_i, tmp1313)

tmp1315 := Call2(__e, SymFn(symelement_2), V3857, tmp1314)


if True == tmp1315 {
return True
} else {
tmp1303 := MakeNative0(func(__e *ControlFlow) Obj {
tmp1304 := PrimGuard1(__e, symstr, PrimStr, V3857)

let__1169 := tmp1304
_ = let__1169

return TailCall1(__e, SymFn(symshen_4analyse_1symbol_2), let__1169)


panic("unreachable")
})

tmp1305 := MakeNative1(func(__e *ControlFlow, Z3859 Obj) Obj {
return False
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp1303, tmp1305)


}


}


panic("unreachable")
})

tmp1333 := Call2(__e, ns2_1set, symsymbol_2, tmp1302)


_ = tmp1333

tmp1334 := MakeNative1(func(__e *ControlFlow, V3862 Obj) Obj {
tmp1343 := Call1(__e, SymFn(symshen_4_7string_2), V3862)


if True == tmp1343 {
tmp1339 := Call1(__e, SymFn(symhdstr), V3862)


tmp1340 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp1339)

tmp1341 := Call1(__e, SymFn(symshen_4alpha_2), tmp1340)


if True == tmp1341 {
tmp1337 := Call1(__e, SymFn(symshen_4alphanums_2), PrimGuard1(__e, symtlstr, PrimTailString, V3862))


if True == tmp1337 {
return True
} else {
return False
}


} else {
return False
}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.analyse-symbol?"))
}


panic("unreachable")
})

tmp1344 := Call2(__e, ns2_1set, symshen_4analyse_1symbol_2, tmp1334)


_ = tmp1344

tmp1345 := MakeNative1(func(__e *ControlFlow, V3865 Obj) Obj {
tmp1359 := GEq(__e, MakeString(""), V3865)

if True == tmp1359 {
return True
} else {
tmp1357 := Call1(__e, SymFn(symshen_4_7string_2), V3865)


if True == tmp1357 {
tmp1346 := Call1(__e, SymFn(symhdstr), V3865)


tmp1347 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp1346)

let__1170 := tmp1347
_ = let__1170

tmp1355 := Call1(__e, SymFn(symshen_4alpha_2), let__1170)


var ifres1352 Obj

if True == tmp1355 {
ifres1352 = True


} else {
tmp1354 := Call1(__e, SymFn(symshen_4digit_2), let__1170)


var ifres1353 Obj

if True == tmp1354 {
ifres1353 = True


} else {
ifres1353 = False


}

ifres1352 = ifres1353


}

if True == ifres1352 {
tmp1350 := Call1(__e, SymFn(symshen_4alphanums_2), PrimGuard1(__e, symtlstr, PrimTailString, V3865))


if True == tmp1350 {
return True
} else {
return False
}


} else {
return False
}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.alphanums?"))
}


}


panic("unreachable")
})

tmp1360 := Call2(__e, ns2_1set, symshen_4alphanums_2, tmp1345)


_ = tmp1360

tmp1361 := MakeNative1(func(__e *ControlFlow, V3867 Obj) Obj {
tmp1372 := Call1(__e, SymFn(symboolean_2), V3867)


var ifres1366 Obj

if True == tmp1372 {
ifres1366 = True


} else {
tmp1371 := PrimGuard1(__e, symnumber_2, PrimIsNumber, V3867)

var ifres1368 Obj

if True == tmp1371 {
ifres1368 = True


} else {
tmp1370 := PrimGuard1(__e, symstring_2, PrimIsString, V3867)

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
return False
} else {
tmp1362 := MakeNative0(func(__e *ControlFlow) Obj {
tmp1363 := PrimGuard1(__e, symstr, PrimStr, V3867)

let__1171 := tmp1363
_ = let__1171

return TailCall1(__e, SymFn(symshen_4analyse_1variable_2), let__1171)


panic("unreachable")
})

tmp1364 := MakeNative1(func(__e *ControlFlow, Z3869 Obj) Obj {
return False
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp1362, tmp1364)


}


panic("unreachable")
})

tmp1373 := Call2(__e, ns2_1set, symvariable_2, tmp1361)


_ = tmp1373

tmp1374 := MakeNative1(func(__e *ControlFlow, V3872 Obj) Obj {
tmp1383 := Call1(__e, SymFn(symshen_4_7string_2), V3872)


if True == tmp1383 {
tmp1379 := Call1(__e, SymFn(symhdstr), V3872)


tmp1380 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp1379)

tmp1381 := Call1(__e, SymFn(symshen_4uppercase_2), tmp1380)


if True == tmp1381 {
tmp1377 := Call1(__e, SymFn(symshen_4alphanums_2), PrimGuard1(__e, symtlstr, PrimTailString, V3872))


if True == tmp1377 {
return True
} else {
return False
}


} else {
return False
}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.analyse-variable?"))
}


panic("unreachable")
})

tmp1384 := Call2(__e, ns2_1set, symshen_4analyse_1variable_2, tmp1374)


_ = tmp1384

tmp1385 := MakeNative1(func(__e *ControlFlow, V3873 Obj) Obj {
tmp1386 := GValue(__e, symshen_4_dgensym_d)

tmp1388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dgensym_d, GAdd(__e, MakeInteger(1), tmp1386))
}
__typedArg0 := symshen_4_dgensym_d
__typedArg1 := GAdd(__e, MakeInteger(1), tmp1386)
return Call2(__e, SymFn(symset), __typedArg0, __typedArg1)
})()

return TailCall2(__e, SymFn(symconcat), V3873, tmp1388)


panic("unreachable")
})

tmp1389 := Call2(__e, ns2_1set, symgensym, tmp1385)


_ = tmp1389

tmp1390 := MakeNative2(func(__e *ControlFlow, V3874, V3875 Obj) Obj {
tmp1391 := PrimGuard1(__e, symstr, PrimStr, V3874)

tmp1392 := PrimGuard1(__e, symstr, PrimStr, V3875)

return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(PrimGuard2(__e, symcn, PrimStringConcat, tmp1391, tmp1392))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, tmp1391, tmp1392)
return Call1(__e, SymFn(symintern), __typedArg0)
})()


panic("unreachable")
})

tmp1394 := Call2(__e, ns2_1set, symconcat, tmp1390)


_ = tmp1394

tmp1395 := MakeNative2(func(__e *ControlFlow, V3876, V3877 Obj) Obj {
tmp1396 := PrimGuard1(__e, symabsvector, PrimAbsvector, MakeInteger(3))

let__1172 := tmp1396
_ = let__1172

tmp1397 := GVecSet(__e, let__1172, MakeInteger(0), symshen_4tuple)

let__1173 := tmp1397
_ = let__1173

tmp1398 := GVecSet(__e, let__1172, MakeInteger(1), V3876)

let__1174 := tmp1398
_ = let__1174

tmp1399 := GVecSet(__e, let__1172, MakeInteger(2), V3877)

let__1175 := tmp1399
_ = let__1175

return let__1172


panic("unreachable")
})

tmp1400 := Call2(__e, ns2_1set, sym_8p, tmp1395)


_ = tmp1400

tmp1401 := MakeNative1(func(__e *ControlFlow, V3882 Obj) Obj {
return GVecGet(__e, V3882, MakeInteger(1))
panic("unreachable")
})

tmp1402 := Call2(__e, ns2_1set, symfst, tmp1401)


_ = tmp1402

tmp1403 := MakeNative1(func(__e *ControlFlow, V3883 Obj) Obj {
return GVecGet(__e, V3883, MakeInteger(2))
panic("unreachable")
})

tmp1404 := Call2(__e, ns2_1set, symsnd, tmp1403)


_ = tmp1404

tmp1405 := MakeNative1(func(__e *ControlFlow, V3884 Obj) Obj {
tmp1412 := GVectorP(__e, V3884)

if True == tmp1412 {
tmp1407 := MakeNative0(func(__e *ControlFlow) Obj {
return GVecGet(__e, V3884, MakeInteger(0))
panic("unreachable")
})

tmp1408 := MakeNative1(func(__e *ControlFlow, Z3885 Obj) Obj {
return symshen_4not_1tuple
panic("unreachable")
})

tmp1409 := Call2(__e, try_1catch, tmp1407, tmp1408)


tmp1410 := GEq(__e, symshen_4tuple, tmp1409)

if True == tmp1410 {
return True
} else {
return False
}


} else {
return False
}


panic("unreachable")
})

tmp1413 := Call2(__e, ns2_1set, symtuple_2, tmp1405)


_ = tmp1413

tmp1414 := MakeNative2(func(__e *ControlFlow, V3890, V3891 Obj) Obj {
tmp1421 := GEq(__e, Nil, V3890)

if True == tmp1421 {
return V3891
} else {
tmp1419 := GConsP(__e, V3890)

if True == tmp1419 {
tmp1415 := GHd(__e, V3890)

tmp1416 := GTl(__e, V3890)

tmp1417 := Call2(__e, SymFn(symappend), tmp1416, V3891)


return GCons(__e, tmp1415, tmp1417)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to append a non-list"))
}


}


panic("unreachable")
})

tmp1422 := Call2(__e, ns2_1set, symappend, tmp1414)


_ = tmp1422

tmp1423 := MakeNative2(func(__e *ControlFlow, V3892, V3893 Obj) Obj {
tmp1424 := Call1(__e, SymFn(symlimit), V3893)


let__1176 := tmp1424
_ = let__1176

tmp1426 := Call1(__e, SymFn(symvector), GAdd(__e, let__1176, MakeInteger(1)))


let__1177 := tmp1426
_ = let__1177

tmp1427 := Call3(__e, SymFn(symvector_1_6), let__1177, MakeInteger(1), V3892)


let__1178 := tmp1427
_ = let__1178

tmp1429 := GEq(__e, let__1176, MakeInteger(0))

if True == tmp1429 {
return let__1178
} else {
return TailCall4(__e, SymFn(symshen_4_8v_1help), V3893, MakeInteger(1), let__1176, let__1178)
}


panic("unreachable")
})

tmp1430 := Call2(__e, ns2_1set, sym_8v, tmp1423)


_ = tmp1430

tmp1431 := (func() Obj {
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
tmp1437 := GEq(__e, V3899, V3900)

if True == tmp1437 {
return TailCall4(__e, SymFn(symshen_4copyfromvector), V3898, V3901, V3900, GAdd(__e, V3900, MakeInteger(1)))


} else {
tmp1433 := GAdd(__e, V3899, MakeInteger(1))

tmp1435 := Call4(__e, SymFn(symshen_4copyfromvector), V3898, V3901, V3899, GAdd(__e, V3899, MakeInteger(1)))


if SymFn(symshen_4_8v_1help) == __self {
__self1, __self2, __self3, __self4 = V3898, tmp1433, V3900, tmp1435
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4_8v_1help), V3898, tmp1433, V3900, tmp1435)


}


panic("unreachable")
})
return __self
})()

tmp1438 := Call2(__e, ns2_1set, symshen_4_8v_1help, tmp1431)


_ = tmp1438

tmp1439 := MakeNative4(func(__e *ControlFlow, V3902, V3903, V3904, V3905 Obj) Obj {
tmp1440 := MakeNative0(func(__e *ControlFlow) Obj {
tmp1441 := Call2(__e, SymFn(sym_5_1vector), V3902, V3904)


return TailCall3(__e, SymFn(symvector_1_6), V3903, V3905, tmp1441)


panic("unreachable")
})

tmp1442 := MakeNative1(func(__e *ControlFlow, Z3906 Obj) Obj {
return V3903
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp1440, tmp1442)


panic("unreachable")
})

tmp1443 := Call2(__e, ns2_1set, symshen_4copyfromvector, tmp1439)


_ = tmp1443

tmp1444 := MakeNative1(func(__e *ControlFlow, V3907 Obj) Obj {
tmp1445 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall2(__e, SymFn(sym_5_1vector), V3907, MakeInteger(1))
panic("unreachable")
})

tmp1446 := MakeNative1(func(__e *ControlFlow, Z3908 Obj) Obj {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("hdv needs a non-empty vector as an argument\n"))
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp1445, tmp1446)


panic("unreachable")
})

tmp1447 := Call2(__e, ns2_1set, symhdv, tmp1444)


_ = tmp1447

tmp1448 := MakeNative1(func(__e *ControlFlow, V3909 Obj) Obj {
tmp1449 := Call1(__e, SymFn(symlimit), V3909)


let__1179 := tmp1449
_ = let__1179

tmp1457 := GEq(__e, let__1179, MakeInteger(0))

if True == tmp1457 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("cannot take the tail of the empty vector\n"))
} else {
tmp1455 := GEq(__e, let__1179, MakeInteger(1))

if True == tmp1455 {
return TailCall1(__e, SymFn(symvector), MakeInteger(0))
} else {
tmp1451 := Call1(__e, SymFn(symvector), GSub(__e, let__1179, MakeInteger(1)))


let__1180 := tmp1451
_ = let__1180

tmp1453 := Call1(__e, SymFn(symvector), GSub(__e, let__1179, MakeInteger(1)))


return TailCall4(__e, SymFn(symshen_4tlv_1help), V3909, MakeInteger(2), let__1179, tmp1453)


}


}


panic("unreachable")
})

tmp1458 := Call2(__e, ns2_1set, symtlv, tmp1448)


_ = tmp1458

tmp1459 := (func() Obj {
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
tmp1465 := GEq(__e, V3914, V3915)

if True == tmp1465 {
return TailCall4(__e, SymFn(symshen_4copyfromvector), V3913, V3916, V3915, GSub(__e, V3915, MakeInteger(1)))


} else {
tmp1461 := GAdd(__e, V3914, MakeInteger(1))

tmp1463 := Call4(__e, SymFn(symshen_4copyfromvector), V3913, V3916, V3914, GSub(__e, V3914, MakeInteger(1)))


if SymFn(symshen_4tlv_1help) == __self {
__self1, __self2, __self3, __self4 = V3913, tmp1461, V3915, tmp1463
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4tlv_1help), V3913, tmp1461, V3915, tmp1463)


}


panic("unreachable")
})
return __self
})()

tmp1466 := Call2(__e, ns2_1set, symshen_4tlv_1help, tmp1459)


_ = tmp1466

tmp1467 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V3928 := __self1
_ = V3928
V3929 := __self2
_ = V3929
tmp1483 := GEq(__e, Nil, V3929)

if True == tmp1483 {
return Nil
} else {
tmp1481 := GConsP(__e, V3929)

var ifres1472 Obj

if True == tmp1481 {
tmp1479 := GHd(__e, V3929)

tmp1480 := GConsP(__e, tmp1479)

var ifres1474 Obj

if True == tmp1480 {
tmp1476 := GHd(__e, V3929)

tmp1477 := GHd(__e, tmp1476)

tmp1478 := GEq(__e, V3928, tmp1477)

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
return GHd(__e, V3929)
} else {
tmp1470 := GConsP(__e, V3929)

if True == tmp1470 {
tmp1468 := GTl(__e, V3929)

if SymFn(symassoc) == __self {
__self1, __self2 = V3928, tmp1468
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symassoc), V3928, tmp1468)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to search a non-list with assoc\n"))
}


}


}


panic("unreachable")
})
return __self
})()

tmp1484 := Call2(__e, ns2_1set, symassoc, tmp1467)


_ = tmp1484

tmp1485 := MakeNative3(func(__e *ControlFlow, V3933, V3934, V3935 Obj) Obj {
tmp1508 := GEq(__e, Nil, V3935)

if True == tmp1508 {
tmp1486 := GCons(__e, V3933, V3934)

return GCons(__e, tmp1486, Nil)


} else {
tmp1506 := GConsP(__e, V3935)

var ifres1497 Obj

if True == tmp1506 {
tmp1504 := GHd(__e, V3935)

tmp1505 := GConsP(__e, tmp1504)

var ifres1499 Obj

if True == tmp1505 {
tmp1501 := GHd(__e, V3935)

tmp1502 := GHd(__e, tmp1501)

tmp1503 := GEq(__e, V3933, tmp1502)

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
tmp1487 := GHd(__e, V3935)

tmp1488 := GHd(__e, tmp1487)

tmp1489 := GCons(__e, tmp1488, V3934)

tmp1490 := GTl(__e, V3935)

return GCons(__e, tmp1489, tmp1490)


} else {
tmp1495 := GConsP(__e, V3935)

if True == tmp1495 {
tmp1491 := GHd(__e, V3935)

tmp1492 := GTl(__e, V3935)

tmp1493 := Call3(__e, SymFn(symshen_4assoc_1set), V3933, V3934, tmp1492)


return GCons(__e, tmp1491, tmp1493)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4assoc_1set)
}


}


}


panic("unreachable")
})

tmp1509 := Call2(__e, ns2_1set, symshen_4assoc_1set, tmp1485)


_ = tmp1509

tmp1510 := MakeNative2(func(__e *ControlFlow, V3939, V3940 Obj) Obj {
tmp1528 := GEq(__e, Nil, V3940)

if True == tmp1528 {
return Nil
} else {
tmp1526 := GConsP(__e, V3940)

var ifres1517 Obj

if True == tmp1526 {
tmp1524 := GHd(__e, V3940)

tmp1525 := GConsP(__e, tmp1524)

var ifres1519 Obj

if True == tmp1525 {
tmp1521 := GHd(__e, V3940)

tmp1522 := GHd(__e, tmp1521)

tmp1523 := GEq(__e, V3939, tmp1522)

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
return GTl(__e, V3940)
} else {
tmp1515 := GConsP(__e, V3940)

if True == tmp1515 {
tmp1511 := GHd(__e, V3940)

tmp1512 := GTl(__e, V3940)

tmp1513 := Call2(__e, SymFn(symshen_4assoc_1rm), V3939, tmp1512)


return GCons(__e, tmp1511, tmp1513)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4assoc_1rm)
}


}


}


panic("unreachable")
})

tmp1529 := Call2(__e, ns2_1set, symshen_4assoc_1rm, tmp1510)


_ = tmp1529

tmp1530 := MakeNative1(func(__e *ControlFlow, V3943 Obj) Obj {
tmp1534 := GEq(__e, True, V3943)

if True == tmp1534 {
return True
} else {
tmp1532 := GEq(__e, False, V3943)

if True == tmp1532 {
return True
} else {
return False
}


}


panic("unreachable")
})

tmp1535 := Call2(__e, ns2_1set, symboolean_2, tmp1530)


_ = tmp1535

tmp1536 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V3944 := __self1
_ = V3944
tmp1541 := GEq(__e, MakeInteger(0), V3944)

if True == tmp1541 {
return MakeInteger(0)
} else {
tmp1537 := Call0(__e, SymFn(symstoutput))


tmp1538 := Call2(__e, SymFn(sympr), MakeString("\n"), tmp1537)


_ = tmp1538

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

tmp1542 := Call2(__e, ns2_1set, symnl, tmp1536)


_ = tmp1542

tmp1543 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V3951 := __self1
_ = V3951
V3952 := __self2
_ = V3952
tmp1554 := GEq(__e, Nil, V3951)

if True == tmp1554 {
return Nil
} else {
tmp1552 := GConsP(__e, V3951)

if True == tmp1552 {
tmp1549 := GHd(__e, V3951)

tmp1550 := Call2(__e, SymFn(symelement_2), tmp1549, V3952)


if True == tmp1550 {
tmp1544 := GTl(__e, V3951)

if SymFn(symdifference) == __self {
__self1, __self2 = tmp1544, V3952
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symdifference), tmp1544, V3952)


} else {
tmp1545 := GHd(__e, V3951)

tmp1546 := GTl(__e, V3951)

tmp1547 := Call2(__e, SymFn(symdifference), tmp1546, V3952)


return GCons(__e, tmp1545, tmp1547)


}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to find the difference with a non-list\n"))
}


}


panic("unreachable")
})
return __self
})()

tmp1555 := Call2(__e, ns2_1set, symdifference, tmp1543)


_ = tmp1555

tmp1556 := MakeNative2(func(__e *ControlFlow, V3953, V3954 Obj) Obj {
return V3954
panic("unreachable")
})

tmp1557 := Call2(__e, ns2_1set, symdo, tmp1556)


_ = tmp1557

tmp1558 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V3966 := __self1
_ = V3966
V3967 := __self2
_ = V3967
tmp1569 := GEq(__e, Nil, V3967)

if True == tmp1569 {
return False
} else {
tmp1567 := GConsP(__e, V3967)

var ifres1563 Obj

if True == tmp1567 {
tmp1565 := GHd(__e, V3967)

tmp1566 := GEq(__e, V3966, tmp1565)

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
return True
} else {
tmp1561 := GConsP(__e, V3967)

if True == tmp1561 {
tmp1559 := GTl(__e, V3967)

if SymFn(symelement_2) == __self {
__self1, __self2 = V3966, tmp1559
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symelement_2), V3966, tmp1559)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to find an element in a non-list\n"))
}


}


}


panic("unreachable")
})
return __self
})()

tmp1570 := Call2(__e, ns2_1set, symelement_2, tmp1558)


_ = tmp1570

tmp1571 := MakeNative1(func(__e *ControlFlow, V3970 Obj) Obj {
tmp1573 := GEq(__e, Nil, V3970)

if True == tmp1573 {
return True
} else {
return False
}


panic("unreachable")
})

tmp1574 := Call2(__e, ns2_1set, symempty_2, tmp1571)


_ = tmp1574

tmp1575 := MakeNative2(func(__e *ControlFlow, V3971, V3972 Obj) Obj {
tmp1576 := Call1(__e, V3971, V3972)


return TailCall3(__e, SymFn(symshen_4fix_1help), V3971, V3972, tmp1576)


panic("unreachable")
})

tmp1577 := Call2(__e, ns2_1set, symfix, tmp1575)


_ = tmp1577

tmp1578 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V3978 := __self1
_ = V3978
V3979 := __self2
_ = V3979
V3980 := __self3
_ = V3980
tmp1581 := GEq(__e, V3979, V3980)

if True == tmp1581 {
return V3980
} else {
tmp1579 := Call1(__e, V3978, V3980)


if SymFn(symshen_4fix_1help) == __self {
__self1, __self2, __self3 = V3978, V3980, tmp1579
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4fix_1help), V3978, V3980, tmp1579)


}


panic("unreachable")
})
return __self
})()

tmp1582 := Call2(__e, ns2_1set, symshen_4fix_1help, tmp1578)


_ = tmp1582

tmp1583 := MakeNative4(func(__e *ControlFlow, V3981, V3982, V3983, V3984 Obj) Obj {
tmp1584 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall2(__e, SymFn(symshen_4_5_1dict), V3984, V3981)
panic("unreachable")
})

tmp1585 := MakeNative1(func(__e *ControlFlow, Z3986 Obj) Obj {
return Nil
panic("unreachable")
})

tmp1586 := Call2(__e, try_1catch, tmp1584, tmp1585)


let__1181 := tmp1586
_ = let__1181

tmp1587 := Call3(__e, SymFn(symshen_4assoc_1set), V3982, V3983, let__1181)


let__1182 := tmp1587
_ = let__1182

tmp1588 := Call3(__e, SymFn(symshen_4dict_1_6), V3984, V3981, let__1182)


let__1183 := tmp1588
_ = let__1183

return V3983


panic("unreachable")
})

tmp1589 := Call2(__e, ns2_1set, symput, tmp1583)


_ = tmp1589

tmp1590 := MakeNative3(func(__e *ControlFlow, V3989, V3990, V3991 Obj) Obj {
tmp1591 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall2(__e, SymFn(symshen_4_5_1dict), V3991, V3989)
panic("unreachable")
})

tmp1592 := MakeNative1(func(__e *ControlFlow, Z3993 Obj) Obj {
return Nil
panic("unreachable")
})

tmp1593 := Call2(__e, try_1catch, tmp1591, tmp1592)


let__1184 := tmp1593
_ = let__1184

tmp1594 := Call2(__e, SymFn(symshen_4assoc_1rm), V3990, let__1184)


let__1185 := tmp1594
_ = let__1185

tmp1595 := Call3(__e, SymFn(symshen_4dict_1_6), V3991, V3989, let__1185)


let__1186 := tmp1595
_ = let__1186

return V3989


panic("unreachable")
})

tmp1596 := Call2(__e, ns2_1set, symunput, tmp1590)


_ = tmp1596

tmp1597 := MakeNative3(func(__e *ControlFlow, V3996, V3997, V3998 Obj) Obj {
tmp1598 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall2(__e, SymFn(symshen_4_5_1dict), V3998, V3996)
panic("unreachable")
})

tmp1599 := MakeNative1(func(__e *ControlFlow, Z4000 Obj) Obj {
tmp1600 := Call3(__e, SymFn(symshen_4app), V3997, MakeString("\n"), symshen_4s)


tmp1602 := Call3(__e, SymFn(symshen_4app), V3996, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" has no attributes: "), tmp1600), symshen_4a)


return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp1602)


panic("unreachable")
})

tmp1603 := Call2(__e, try_1catch, tmp1598, tmp1599)


let__1187 := tmp1603
_ = let__1187

tmp1604 := Call2(__e, SymFn(symassoc), V3997, let__1187)


let__1188 := tmp1604
_ = let__1188

tmp1610 := Call1(__e, SymFn(symempty_2), let__1188)


if True == tmp1610 {
tmp1605 := Call3(__e, SymFn(symshen_4app), V3996, MakeString("\n"), symshen_4s)


tmp1607 := Call3(__e, SymFn(symshen_4app), V3997, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" not found for "), tmp1605), symshen_4s)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("attribute "), tmp1607))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("attribute "), tmp1607)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


} else {
return GTl(__e, let__1188)
}


panic("unreachable")
})

tmp1611 := Call2(__e, ns2_1set, symget, tmp1597)


_ = tmp1611

tmp1612 := MakeNative2(func(__e *ControlFlow, V4002, V4003 Obj) Obj {
tmp1613 := Call1(__e, SymFn(symshen_4hashkey), V4002)


tmp1614 := Call2(__e, SymFn(symshen_4mod), tmp1613, V4003)


let__1189 := tmp1614
_ = let__1189

tmp1616 := GEq(__e, let__1189, MakeInteger(0))

if True == tmp1616 {
return MakeInteger(1)
} else {
return let__1189
}


panic("unreachable")
})

tmp1617 := Call2(__e, ns2_1set, symhash, tmp1612)


_ = tmp1617

tmp1618 := MakeNative1(func(__e *ControlFlow, V4005 Obj) Obj {
tmp1619 := MakeNative1(func(__e *ControlFlow, Z4007 Obj) Obj {
return PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, Z4007)
panic("unreachable")
})

tmp1620 := Call1(__e, SymFn(symexplode), V4005)


tmp1621 := Call2(__e, SymFn(symmap), tmp1619, tmp1620)


let__1190 := tmp1621
_ = let__1190

return TailCall2(__e, SymFn(symshen_4prodbutzero), let__1190, MakeInteger(1))


panic("unreachable")
})

tmp1622 := Call2(__e, ns2_1set, symshen_4hashkey, tmp1618)


_ = tmp1622

tmp1623 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4008 := __self1
_ = V4008
V4009 := __self2
_ = V4009
tmp1642 := GEq(__e, Nil, V4008)

if True == tmp1642 {
return V4009
} else {
tmp1640 := GConsP(__e, V4008)

var ifres1636 Obj

if True == tmp1640 {
tmp1638 := GHd(__e, V4008)

tmp1639 := GEq(__e, MakeInteger(0), tmp1638)

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
tmp1624 := GTl(__e, V4008)

if SymFn(symshen_4prodbutzero) == __self {
__self1, __self2 = tmp1624, V4009
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4prodbutzero), tmp1624, V4009)


} else {
tmp1634 := GConsP(__e, V4008)

if True == tmp1634 {
if True == GGt(__e, V4009, MakeNumber(1e+10)) {
tmp1625 := GTl(__e, V4008)

tmp1626 := GHd(__e, V4008)

if SymFn(symshen_4prodbutzero) == __self {
__self1, __self2 = tmp1625, GAdd(__e, V4009, tmp1626)
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4prodbutzero), tmp1625, GAdd(__e, V4009, tmp1626))


} else {
tmp1628 := GTl(__e, V4008)

tmp1629 := GHd(__e, V4008)

if SymFn(symshen_4prodbutzero) == __self {
__self1, __self2 = tmp1628, GMul(__e, V4009, tmp1629)
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4prodbutzero), tmp1628, GMul(__e, V4009, tmp1629))


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

tmp1643 := Call2(__e, ns2_1set, symshen_4prodbutzero, tmp1623)


_ = tmp1643

tmp1644 := MakeNative2(func(__e *ControlFlow, V4010, V4011 Obj) Obj {
tmp1645 := GCons(__e, V4011, Nil)

tmp1646 := Call2(__e, SymFn(symshen_4multiples), V4010, tmp1645)


return TailCall2(__e, SymFn(symshen_4modh), V4010, tmp1646)


panic("unreachable")
})

tmp1647 := Call2(__e, ns2_1set, symshen_4mod, tmp1644)


_ = tmp1647

tmp1648 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4016 := __self1
_ = V4016
V4017 := __self2
_ = V4017
tmp1659 := GConsP(__e, V4017)

var ifres1655 Obj

if True == tmp1659 {
tmp1657 := GHd(__e, V4017)

tmp1658 := GGt(__e, tmp1657, V4016)

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
return GTl(__e, V4017)
} else {
tmp1653 := GConsP(__e, V4017)

if True == tmp1653 {
tmp1649 := GHd(__e, V4017)

tmp1651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(GMul(__e, MakeInteger(2), tmp1649), V4017)
}
__typedArg0 := GMul(__e, MakeInteger(2), tmp1649)
__typedArg1 := V4017
return Call2(__e, SymFn(symcons), __typedArg0, __typedArg1)
})()

if SymFn(symshen_4multiples) == __self {
__self1, __self2 = V4016, tmp1651
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4multiples), V4016, tmp1651)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.multiples"))
}


}


panic("unreachable")
})
return __self
})()

tmp1660 := Call2(__e, ns2_1set, symshen_4multiples, tmp1648)


_ = tmp1660

tmp1661 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4024 := __self1
_ = V4024
V4025 := __self2
_ = V4025
tmp1679 := GEq(__e, MakeInteger(0), V4024)

if True == tmp1679 {
return MakeInteger(0)
} else {
tmp1677 := GEq(__e, Nil, V4025)

if True == tmp1677 {
return V4024
} else {
tmp1675 := GConsP(__e, V4025)

var ifres1671 Obj

if True == tmp1675 {
tmp1673 := GHd(__e, V4025)

tmp1674 := GGt(__e, tmp1673, V4024)

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
tmp1664 := GTl(__e, V4025)

tmp1665 := Call1(__e, SymFn(symempty_2), tmp1664)


if True == tmp1665 {
return V4024
} else {
tmp1662 := GTl(__e, V4025)

if SymFn(symshen_4modh) == __self {
__self1, __self2 = V4024, tmp1662
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4modh), V4024, tmp1662)


}


} else {
tmp1669 := GConsP(__e, V4025)

if True == tmp1669 {
tmp1666 := GHd(__e, V4025)

if SymFn(symshen_4modh) == __self {
__self1, __self2 = GSub(__e, V4024, tmp1666), V4025
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4modh), GSub(__e, V4024, tmp1666), V4025)


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

tmp1680 := Call2(__e, ns2_1set, symshen_4modh, tmp1661)


_ = tmp1680

tmp1681 := MakeNative1(func(__e *ControlFlow, V4028 Obj) Obj {
tmp1688 := GEq(__e, Nil, V4028)

if True == tmp1688 {
return MakeInteger(0)
} else {
tmp1686 := GConsP(__e, V4028)

if True == tmp1686 {
tmp1682 := GHd(__e, V4028)

tmp1683 := GTl(__e, V4028)

tmp1684 := Call1(__e, SymFn(symsum), tmp1683)


return GAdd(__e, tmp1682, tmp1684)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to sum a non-list\n"))
}


}


panic("unreachable")
})

tmp1689 := Call2(__e, ns2_1set, symsum, tmp1681)


_ = tmp1689

tmp1690 := MakeNative1(func(__e *ControlFlow, V4033 Obj) Obj {
tmp1692 := GConsP(__e, V4033)

if True == tmp1692 {
return GHd(__e, V4033)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("head expects a non-empty list\n"))
}


panic("unreachable")
})

tmp1693 := Call2(__e, ns2_1set, symhead, tmp1690)


_ = tmp1693

tmp1694 := MakeNative1(func(__e *ControlFlow, V4038 Obj) Obj {
tmp1696 := GConsP(__e, V4038)

if True == tmp1696 {
return GTl(__e, V4038)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("tail expects a non-empty list\n"))
}


panic("unreachable")
})

tmp1697 := Call2(__e, ns2_1set, symtail, tmp1694)


_ = tmp1697

tmp1698 := MakeNative1(func(__e *ControlFlow, V4039 Obj) Obj {
return PrimGuard2(__e, sympos, PrimPos, V4039, MakeInteger(0))
panic("unreachable")
})

tmp1699 := Call2(__e, ns2_1set, symhdstr, tmp1698)


_ = tmp1699

tmp1700 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4046 := __self1
_ = V4046
V4047 := __self2
_ = V4047
tmp1711 := GEq(__e, Nil, V4046)

if True == tmp1711 {
return Nil
} else {
tmp1709 := GConsP(__e, V4046)

if True == tmp1709 {
tmp1706 := GHd(__e, V4046)

tmp1707 := Call2(__e, SymFn(symelement_2), tmp1706, V4047)


if True == tmp1707 {
tmp1701 := GHd(__e, V4046)

tmp1702 := GTl(__e, V4046)

tmp1703 := Call2(__e, SymFn(symintersection), tmp1702, V4047)


return GCons(__e, tmp1701, tmp1703)


} else {
tmp1704 := GTl(__e, V4046)

if SymFn(symintersection) == __self {
__self1, __self2 = tmp1704, V4047
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symintersection), tmp1704, V4047)


}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to find the intersection with a non-list\n"))
}


}


panic("unreachable")
})
return __self
})()

tmp1712 := Call2(__e, ns2_1set, symintersection, tmp1700)


_ = tmp1712

tmp1713 := MakeNative1(func(__e *ControlFlow, V4048 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4reverse_1help), V4048, Nil)
panic("unreachable")
})

tmp1714 := Call2(__e, ns2_1set, symreverse, tmp1713)


_ = tmp1714

tmp1715 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4053 := __self1
_ = V4053
V4054 := __self2
_ = V4054
tmp1722 := GEq(__e, Nil, V4053)

if True == tmp1722 {
return V4054
} else {
tmp1720 := GConsP(__e, V4053)

if True == tmp1720 {
tmp1716 := GTl(__e, V4053)

tmp1717 := GHd(__e, V4053)

tmp1718 := GCons(__e, tmp1717, V4054)

if SymFn(symshen_4reverse_1help) == __self {
__self1, __self2 = tmp1716, tmp1718
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4reverse_1help), tmp1716, tmp1718)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to reverse a non-list\n"))
}


}


panic("unreachable")
})
return __self
})()

tmp1723 := Call2(__e, ns2_1set, symshen_4reverse_1help, tmp1715)


_ = tmp1723

tmp1724 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4059 := __self1
_ = V4059
V4060 := __self2
_ = V4060
tmp1735 := GEq(__e, Nil, V4059)

if True == tmp1735 {
return V4060
} else {
tmp1733 := GConsP(__e, V4059)

if True == tmp1733 {
tmp1730 := GHd(__e, V4059)

tmp1731 := Call2(__e, SymFn(symelement_2), tmp1730, V4060)


if True == tmp1731 {
tmp1725 := GTl(__e, V4059)

if SymFn(symunion) == __self {
__self1, __self2 = tmp1725, V4060
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symunion), tmp1725, V4060)


} else {
tmp1726 := GHd(__e, V4059)

tmp1727 := GTl(__e, V4059)

tmp1728 := Call2(__e, SymFn(symunion), tmp1727, V4060)


return GCons(__e, tmp1726, tmp1728)


}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to find the union with a non-list\n"))
}


}


panic("unreachable")
})
return __self
})()

tmp1736 := Call2(__e, ns2_1set, symunion, tmp1724)


_ = tmp1736

tmp1737 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V4061 := __self1
_ = V4061
tmp1738 := Call1(__e, SymFn(symshen_4proc_1nl), V4061)


tmp1739 := Call0(__e, SymFn(symstoutput))


tmp1740 := Call2(__e, SymFn(sympr), tmp1738, tmp1739)


let__1191 := tmp1740
_ = let__1191

tmp1741 := Call0(__e, SymFn(symstoutput))


tmp1742 := Call2(__e, SymFn(sympr), MakeString(" (y/n) "), tmp1741)


let__1192 := tmp1742
_ = let__1192

tmp1743 := Call0(__e, SymFn(symstinput))


tmp1744 := Call1(__e, SymFn(symread), tmp1743)


tmp1745 := Call3(__e, SymFn(symshen_4app), tmp1744, MakeString(""), symshen_4s)


let__1193 := tmp1745
_ = let__1193

tmp1751 := GEq(__e, MakeString("y"), let__1193)

if True == tmp1751 {
return True
} else {
tmp1749 := GEq(__e, MakeString("n"), let__1193)

if True == tmp1749 {
return False
} else {
tmp1746 := Call0(__e, SymFn(symstoutput))


tmp1747 := Call2(__e, SymFn(sympr), MakeString("please answer y or n\n"), tmp1746)


_ = tmp1747

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

tmp1752 := Call2(__e, ns2_1set, symy_1or_1n_2, tmp1737)


_ = tmp1752

tmp1753 := MakeNative1(func(__e *ControlFlow, V4065 Obj) Obj {
if True == V4065 {
return False
} else {
return True
}
panic("unreachable")
})

tmp1755 := Call2(__e, ns2_1set, symnot, tmp1753)


_ = tmp1755

tmp1756 := MakeNative0(func(__e *ControlFlow) Obj {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString(""))
panic("unreachable")
})

tmp1757 := Call2(__e, ns2_1set, symabort, tmp1756)


_ = tmp1757

tmp1758 := MakeNative3(func(__e *ControlFlow, V4071, V4072, V4073 Obj) Obj {
tmp1766 := GEq(__e, V4072, V4073)

if True == tmp1766 {
return V4071
} else {
tmp1764 := GConsP(__e, V4073)

if True == tmp1764 {
tmp1759 := GHd(__e, V4073)

tmp1760 := Call3(__e, SymFn(symsubst), V4071, V4072, tmp1759)


tmp1761 := GTl(__e, V4073)

tmp1762 := Call3(__e, SymFn(symsubst), V4071, V4072, tmp1761)


return GCons(__e, tmp1760, tmp1762)


} else {
return V4073
}


}


panic("unreachable")
})

tmp1767 := Call2(__e, ns2_1set, symsubst, tmp1758)


_ = tmp1767

tmp1768 := MakeNative1(func(__e *ControlFlow, V4074 Obj) Obj {
tmp1769 := Call3(__e, SymFn(symshen_4app), V4074, MakeString(""), symshen_4a)


return TailCall1(__e, SymFn(symshen_4explode_1h), tmp1769)


panic("unreachable")
})

tmp1770 := Call2(__e, ns2_1set, symexplode, tmp1768)


_ = tmp1770

tmp1771 := MakeNative1(func(__e *ControlFlow, V4077 Obj) Obj {
tmp1778 := GEq(__e, MakeString(""), V4077)

if True == tmp1778 {
return Nil
} else {
tmp1776 := Call1(__e, SymFn(symshen_4_7string_2), V4077)


if True == tmp1776 {
tmp1772 := Call1(__e, SymFn(symhdstr), V4077)


tmp1774 := Call1(__e, SymFn(symshen_4explode_1h), PrimGuard1(__e, symtlstr, PrimTailString, V4077))


return GCons(__e, tmp1772, tmp1774)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in explode-h"))
}


}


panic("unreachable")
})

tmp1779 := Call2(__e, ns2_1set, symshen_4explode_1h, tmp1771)


_ = tmp1779

tmp1780 := MakeNative1(func(__e *ControlFlow, V4078 Obj) Obj {
tmp1783 := GEq(__e, V4078, MakeString(""))

var ifres1781 Obj

if True == tmp1783 {
ifres1781 = MakeString("")


} else {
tmp1782 := Call3(__e, SymFn(symshen_4app), V4078, MakeString("/"), symshen_4a)


ifres1781 = tmp1782


}

return PrimGuard2(__e, symset, PrimSet, sym_dhome_1directory_d, ifres1781)


panic("unreachable")
})

tmp1784 := Call2(__e, ns2_1set, symcd, tmp1780)


_ = tmp1784

tmp1785 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4079 := __self1
_ = V4079
V4080 := __self2
_ = V4080
tmp1792 := GEq(__e, Nil, V4080)

if True == tmp1792 {
return True
} else {
tmp1790 := GConsP(__e, V4080)

if True == tmp1790 {
tmp1786 := GHd(__e, V4080)

tmp1787 := Call1(__e, V4079, tmp1786)


let__1194 := tmp1787
_ = let__1194

tmp1788 := GTl(__e, V4080)

if SymFn(symshen_4for_1each) == __self {
__self1, __self2 = V4079, tmp1788
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4for_1each), V4079, tmp1788)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4for_1each)
}


}


panic("unreachable")
})
return __self
})()

tmp1793 := Call2(__e, ns2_1set, symshen_4for_1each, tmp1785)


_ = tmp1793

tmp1794 := MakeNative2(func(__e *ControlFlow, V4082, V4083 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4map_1h), V4082, V4083, Nil)
panic("unreachable")
})

tmp1795 := Call2(__e, ns2_1set, symmap, tmp1794)


_ = tmp1795

tmp1796 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V4084 := __self1
_ = V4084
V4085 := __self2
_ = V4085
V4086 := __self3
_ = V4086
tmp1804 := GEq(__e, Nil, V4085)

if True == tmp1804 {
return TailCall1(__e, SymFn(symreverse), V4086)
} else {
tmp1802 := GConsP(__e, V4085)

if True == tmp1802 {
tmp1797 := GTl(__e, V4085)

tmp1798 := GHd(__e, V4085)

tmp1799 := Call1(__e, V4084, tmp1798)


tmp1800 := GCons(__e, tmp1799, V4086)

if SymFn(symshen_4map_1h) == __self {
__self1, __self2, __self3 = V4084, tmp1797, tmp1800
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4map_1h), V4084, tmp1797, tmp1800)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4map_1h)
}


}


panic("unreachable")
})
return __self
})()

tmp1805 := Call2(__e, ns2_1set, symshen_4map_1h, tmp1796)


_ = tmp1805

tmp1806 := MakeNative1(func(__e *ControlFlow, V4087 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4length_1h), V4087, MakeInteger(0))
panic("unreachable")
})

tmp1807 := Call2(__e, ns2_1set, symlength, tmp1806)


_ = tmp1807

tmp1808 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4092 := __self1
_ = V4092
V4093 := __self2
_ = V4093
tmp1812 := GEq(__e, Nil, V4092)

if True == tmp1812 {
return V4093
} else {
tmp1809 := GTl(__e, V4092)

if SymFn(symshen_4length_1h) == __self {
__self1, __self2 = tmp1809, GAdd(__e, V4093, MakeInteger(1))
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4length_1h), tmp1809, GAdd(__e, V4093, MakeInteger(1)))


}


panic("unreachable")
})
return __self
})()

tmp1813 := Call2(__e, ns2_1set, symshen_4length_1h, tmp1808)


_ = tmp1813

tmp1814 := MakeNative2(func(__e *ControlFlow, V4099, V4100 Obj) Obj {
tmp1822 := GEq(__e, V4099, V4100)

if True == tmp1822 {
return MakeInteger(1)
} else {
tmp1820 := GConsP(__e, V4100)

if True == tmp1820 {
tmp1815 := GHd(__e, V4100)

tmp1816 := Call2(__e, SymFn(symoccurrences), V4099, tmp1815)


tmp1817 := GTl(__e, V4100)

tmp1818 := Call2(__e, SymFn(symoccurrences), V4099, tmp1817)


return GAdd(__e, tmp1816, tmp1818)


} else {
return MakeInteger(0)
}


}


panic("unreachable")
})

tmp1823 := Call2(__e, ns2_1set, symoccurrences, tmp1814)


_ = tmp1823

tmp1824 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4105 := __self1
_ = V4105
V4106 := __self2
_ = V4106
tmp1837 := GEq(__e, MakeInteger(1), V4105)

var ifres1834 Obj

if True == tmp1837 {
tmp1836 := GConsP(__e, V4106)

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
return GHd(__e, V4106)
} else {
tmp1832 := GConsP(__e, V4106)

if True == tmp1832 {
tmp1825 := GSub(__e, V4105, MakeInteger(1))

tmp1826 := GTl(__e, V4106)

if SymFn(symnth) == __self {
__self1, __self2 = tmp1825, tmp1826
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symnth), tmp1825, tmp1826)


} else {
tmp1827 := Call3(__e, SymFn(symshen_4app), V4106, MakeString("\n"), symshen_4a)


tmp1829 := Call3(__e, SymFn(symshen_4app), V4105, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(", "), tmp1827), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("nth applied to "), tmp1829))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("nth applied to "), tmp1829)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


}


panic("unreachable")
})
return __self
})()

tmp1838 := Call2(__e, ns2_1set, symnth, tmp1824)


_ = tmp1838

tmp1839 := MakeNative1(func(__e *ControlFlow, V4107 Obj) Obj {
tmp1845 := PrimGuard1(__e, symnumber_2, PrimIsNumber, V4107)

if True == tmp1845 {
tmp1841 := Call1(__e, SymFn(symshen_4abs), V4107)


let__1195 := tmp1841
_ = let__1195

tmp1842 := Call2(__e, SymFn(symshen_4magless), let__1195, MakeInteger(1))


tmp1843 := Call2(__e, SymFn(symshen_4integer_1test_2), let__1195, tmp1842)


if True == tmp1843 {
return True
} else {
return False
}


} else {
return False
}


panic("unreachable")
})

tmp1846 := Call2(__e, ns2_1set, syminteger_2, tmp1839)


_ = tmp1846

tmp1847 := MakeNative1(func(__e *ControlFlow, V4109 Obj) Obj {
if True == GGt(__e, V4109, MakeInteger(0)) {
return V4109
} else {
return GSub(__e, MakeInteger(0), V4109)
}


panic("unreachable")
})

tmp1850 := Call2(__e, ns2_1set, symshen_4abs, tmp1847)


_ = tmp1850

tmp1851 := (func() Obj {
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

tmp1855 := Call2(__e, ns2_1set, symshen_4magless, tmp1851)


_ = tmp1855

tmp1856 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4116 := __self1
_ = V4116
V4117 := __self2
_ = V4117
tmp1863 := GEq(__e, MakeInteger(0), V4116)

if True == tmp1863 {
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

tmp1864 := Call2(__e, ns2_1set, symshen_4integer_1test_2, tmp1856)


_ = tmp1864

tmp1865 := MakeNative2(func(__e *ControlFlow, V4125, V4126 Obj) Obj {
tmp1873 := GEq(__e, Nil, V4126)

if True == tmp1873 {
return Nil
} else {
tmp1871 := GConsP(__e, V4126)

if True == tmp1871 {
tmp1866 := GHd(__e, V4126)

tmp1867 := Call1(__e, V4125, tmp1866)


tmp1868 := GTl(__e, V4126)

tmp1869 := Call2(__e, SymFn(symmapcan), V4125, tmp1868)


return TailCall2(__e, SymFn(symappend), tmp1867, tmp1869)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to mapcan over a non-list\n"))
}


}


panic("unreachable")
})

tmp1874 := Call2(__e, ns2_1set, symmapcan, tmp1865)


_ = tmp1874

tmp1875 := MakeNative2(func(__e *ControlFlow, V4132, V4133 Obj) Obj {
tmp1877 := GEq(__e, V4132, V4133)

if True == tmp1877 {
return True
} else {
return False
}


panic("unreachable")
})

tmp1878 := Call2(__e, ns2_1set, sym_a_a, tmp1875)


_ = tmp1878

tmp1879 := MakeNative1(func(__e *ControlFlow, V4134 Obj) Obj {
tmp1887 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V4134)

if True == tmp1887 {
tmp1881 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, V4134)
panic("unreachable")
})

tmp1882 := MakeNative1(func(__e *ControlFlow, Z4136 Obj) Obj {
return symshen_4this_1symbol_1is_1unbound
panic("unreachable")
})

tmp1883 := Call2(__e, try_1catch, tmp1881, tmp1882)


let__1198 := tmp1883
_ = let__1198

tmp1885 := GEq(__e, let__1198, symshen_4this_1symbol_1is_1unbound)

var ifres1884 Obj

if True == tmp1885 {
ifres1884 = False


} else {
ifres1884 = True


}

if True == ifres1884 {
return True
} else {
return False
}


} else {
return False
}


panic("unreachable")
})

tmp1888 := Call2(__e, ns2_1set, symbound_2, tmp1879)


_ = tmp1888

tmp1889 := MakeNative1(func(__e *ControlFlow, V4137 Obj) Obj {
tmp1895 := GEq(__e, MakeString(""), V4137)

if True == tmp1895 {
return Nil
} else {
tmp1890 := PrimGuard2(__e, sympos, PrimPos, V4137, MakeInteger(0))

tmp1891 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp1890)

tmp1893 := Call1(__e, SymFn(symshen_4string_1_6bytes), PrimGuard1(__e, symtlstr, PrimTailString, V4137))


return GCons(__e, tmp1891, tmp1893)


}


panic("unreachable")
})

tmp1896 := Call2(__e, ns2_1set, symshen_4string_1_6bytes, tmp1889)


_ = tmp1896

tmp1897 := MakeNative1(func(__e *ControlFlow, V4138 Obj) Obj {
if True == GLt(__e, V4138, MakeInteger(0)) {
return GValue(__e, symshen_4_dmaxinferences_d)
} else {
tmp1899 := PrimGuard1(__e, syminteger_2, PrimIsInteger, V4138)

if True == tmp1899 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dmaxinferences_d, V4138)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("maxinferences expects an integer value\n"))
}


}


panic("unreachable")
})

tmp1902 := Call2(__e, ns2_1set, symmaxinferences, tmp1897)


_ = tmp1902

tmp1903 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dinfs_d)
panic("unreachable")
})

tmp1904 := Call2(__e, ns2_1set, syminferences, tmp1903)


_ = tmp1904

tmp1905 := MakeNative1(func(__e *ControlFlow, V4139 Obj) Obj {
return V4139
panic("unreachable")
})

tmp1906 := Call2(__e, ns2_1set, symprotect, tmp1905)


_ = tmp1906

tmp1907 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dsterror_d)
panic("unreachable")
})

tmp1908 := Call2(__e, ns2_1set, symsterror, tmp1907)


_ = tmp1908

tmp1909 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dstoutput_d)
panic("unreachable")
})

tmp1910 := Call2(__e, ns2_1set, symstoutput, tmp1909)


_ = tmp1910

tmp1911 := MakeNative1(func(__e *ControlFlow, V4140 Obj) Obj {
tmp1912 := PrimGuard1(__e, symintern, PrimIntern, V4140)

let__1199 := tmp1912
_ = let__1199

tmp1916 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, let__1199)

if True == tmp1916 {
return let__1199
} else {
tmp1913 := Call3(__e, SymFn(symshen_4app), V4140, MakeString(" to a symbol"), symshen_4s)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("cannot intern "), tmp1913))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("cannot intern "), tmp1913)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


panic("unreachable")
})

tmp1917 := Call2(__e, ns2_1set, symstring_1_6symbol, tmp1911)


_ = tmp1917

tmp1918 := MakeNative1(func(__e *ControlFlow, V4144 Obj) Obj {
tmp1922 := GEq(__e, sym_7, V4144)

if True == tmp1922 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_doptimise_d, True)
} else {
tmp1920 := GEq(__e, sym_1, V4144)

if True == tmp1920 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_doptimise_d, False)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("optimise expects a + or a -.\n"))
}


}


panic("unreachable")
})

tmp1923 := Call2(__e, ns2_1set, symoptimise, tmp1918)


_ = tmp1923

tmp1924 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dos_d)
panic("unreachable")
})

tmp1925 := Call2(__e, ns2_1set, symos, tmp1924)


_ = tmp1925

tmp1926 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dlanguage_d)
panic("unreachable")
})

tmp1927 := Call2(__e, ns2_1set, symlanguage, tmp1926)


_ = tmp1927

tmp1928 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dversion_d)
panic("unreachable")
})

tmp1929 := Call2(__e, ns2_1set, symversion, tmp1928)


_ = tmp1929

tmp1930 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dport_d)
panic("unreachable")
})

tmp1931 := Call2(__e, ns2_1set, symport, tmp1930)


_ = tmp1931

tmp1932 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dporters_d)
panic("unreachable")
})

tmp1933 := Call2(__e, ns2_1set, symporters, tmp1932)


_ = tmp1933

tmp1934 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dimplementation_d)
panic("unreachable")
})

tmp1935 := Call2(__e, ns2_1set, symimplementation, tmp1934)


_ = tmp1935

tmp1936 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_drelease_d)
panic("unreachable")
})

tmp1937 := Call2(__e, ns2_1set, symrelease, tmp1936)


_ = tmp1937

tmp1938 := MakeNative1(func(__e *ControlFlow, V4145 Obj) Obj {
tmp1943 := GEq(__e, symnull, V4145)

if True == tmp1943 {
return True
} else {
tmp1939 := MakeNative0(func(__e *ControlFlow) Obj {
tmp1940 := Call1(__e, SymFn(symexternal), V4145)


_ = tmp1940

return True


panic("unreachable")
})

tmp1941 := MakeNative1(func(__e *ControlFlow, Z4146 Obj) Obj {
return False
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp1939, tmp1941)


}


panic("unreachable")
})

tmp1944 := Call2(__e, ns2_1set, sympackage_2, tmp1938)


_ = tmp1944

tmp1945 := MakeNative0(func(__e *ControlFlow) Obj {
return symshen_4fail_b
panic("unreachable")
})

tmp1946 := Call2(__e, ns2_1set, symfail, tmp1945)


_ = tmp1946

tmp1947 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_duserdefs_d)
panic("unreachable")
})

tmp1948 := Call2(__e, ns2_1set, symuserdefs, tmp1947)


_ = tmp1948

tmp1949 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_doptimise_d)
panic("unreachable")
})

tmp1950 := Call2(__e, ns2_1set, symoptimise_2, tmp1949)


_ = tmp1950

tmp1951 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dhush_d)
panic("unreachable")
})

tmp1952 := Call2(__e, ns2_1set, symhush_2, tmp1951)


_ = tmp1952

tmp1953 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dshen_1type_1theory_1enabled_2_d)
panic("unreachable")
})

tmp1954 := Call2(__e, ns2_1set, symsystem_1S_2, tmp1953)


_ = tmp1954

tmp1955 := MakeNative1(func(__e *ControlFlow, V4149 Obj) Obj {
tmp1959 := GEq(__e, sym_7, V4149)

if True == tmp1959 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dshen_1type_1theory_1enabled_2_d, True)
} else {
tmp1957 := GEq(__e, sym_1, V4149)

if True == tmp1957 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dshen_1type_1theory_1enabled_2_d, False)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("enable-type-theory expects a + or a -\n"))
}


}


panic("unreachable")
})

tmp1960 := Call2(__e, ns2_1set, symenable_1type_1theory, tmp1955)


_ = tmp1960

tmp1961 := MakeNative1(func(__e *ControlFlow, V4152 Obj) Obj {
tmp1965 := GEq(__e, sym_7, V4152)

if True == tmp1965 {
return PrimGuard2(__e, symset, PrimSet, sym_dhush_d, True)
} else {
tmp1963 := GEq(__e, sym_1, V4152)

if True == tmp1963 {
return PrimGuard2(__e, symset, PrimSet, sym_dhush_d, False)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("hush expects a + or a -\n"))
}


}


panic("unreachable")
})

tmp1966 := Call2(__e, ns2_1set, symhush, tmp1961)


_ = tmp1966

tmp1967 := MakeNative1(func(__e *ControlFlow, V4155 Obj) Obj {
tmp1971 := GEq(__e, sym_7, V4155)

if True == tmp1971 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dtc_d, True)
} else {
tmp1969 := GEq(__e, sym_1, V4155)

if True == tmp1969 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dtc_d, False)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("tc expects a + or -"))
}


}


panic("unreachable")
})

tmp1972 := Call2(__e, ns2_1set, symtc, tmp1967)


_ = tmp1972

tmp1973 := MakeNative1(func(__e *ControlFlow, V4156 Obj) Obj {
tmp1974 := GValue(__e, symshen_4_dsigf_d)

tmp1975 := Call2(__e, SymFn(symshen_4unassoc), V4156, tmp1974)


tmp1976 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp1975)

_ = tmp1976

return V4156


panic("unreachable")
})

tmp1977 := Call2(__e, ns2_1set, symdestroy, tmp1973)


_ = tmp1977

tmp1978 := MakeNative2(func(__e *ControlFlow, V4166, V4167 Obj) Obj {
tmp1996 := GEq(__e, Nil, V4167)

if True == tmp1996 {
return Nil
} else {
tmp1994 := GConsP(__e, V4167)

var ifres1985 Obj

if True == tmp1994 {
tmp1992 := GHd(__e, V4167)

tmp1993 := GConsP(__e, tmp1992)

var ifres1987 Obj

if True == tmp1993 {
tmp1989 := GHd(__e, V4167)

tmp1990 := GHd(__e, tmp1989)

tmp1991 := GEq(__e, V4166, tmp1990)

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
return GTl(__e, V4167)
} else {
tmp1983 := GConsP(__e, V4167)

if True == tmp1983 {
tmp1979 := GHd(__e, V4167)

tmp1980 := GTl(__e, V4167)

tmp1981 := Call2(__e, SymFn(symshen_4unassoc), V4166, tmp1980)


return GCons(__e, tmp1979, tmp1981)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.unassoc"))
}


}


}


panic("unreachable")
})

tmp1997 := Call2(__e, ns2_1set, symshen_4unassoc, tmp1978)


_ = tmp1997

tmp1998 := MakeNative1(func(__e *ControlFlow, V4168 Obj) Obj {
tmp2002 := Call1(__e, SymFn(sympackage_2), V4168)


if True == tmp2002 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dpackage_d, V4168)
} else {
tmp1999 := Call3(__e, SymFn(symshen_4app), V4168, MakeString(" does not exist\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("package "), tmp1999))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("package "), tmp1999)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


panic("unreachable")
})

tmp2003 := Call2(__e, ns2_1set, symin_1package, tmp1998)


_ = tmp2003

tmp2004 := MakeNative2(func(__e *ControlFlow, V4169, V4170 Obj) Obj {
tmp2005 := PrimGuard2(__e, symopen, PrimOpenStream, V4169, symout)

let__1200 := tmp2005
_ = let__1200

tmp2008 := PrimGuard1(__e, symstring_2, PrimIsString, V4170)

var ifres2006 Obj

if True == tmp2008 {
ifres2006 = V4170


} else {
tmp2007 := Call3(__e, SymFn(symshen_4app), V4170, MakeString(""), symshen_4s)


ifres2006 = tmp2007


}

let__1201 := ifres2006
_ = let__1201

tmp2009 := Call2(__e, SymFn(sympr), let__1201, let__1200)


let__1202 := tmp2009
_ = let__1202

tmp2010 := PrimGuard1(__e, symclose, PrimCloseStream, let__1200)

let__1203 := tmp2010
_ = let__1203

return V4170


panic("unreachable")
})

tmp2011 := Call2(__e, ns2_1set, symwrite_1to_1file, tmp2004)


_ = tmp2011

tmp2012 := MakeNative0(func(__e *ControlFlow) Obj {
tmp2013 := Call1(__e, SymFn(symgensym), symshen_4t)


return TailCall1(__e, SymFn(symshen_4freshterm), tmp2013)


panic("unreachable")
})

tmp2014 := Call2(__e, ns2_1set, symfresh, tmp2012)


_ = tmp2014

tmp2015 := MakeNative2(func(__e *ControlFlow, V4175, V4176 Obj) Obj {
tmp2016 := GValue(__e, sym_dproperty_1vector_d)

tmp2017 := Call4(__e, SymFn(symput), V4175, symarity, V4176, tmp2016)


let__1204 := tmp2017
_ = let__1204

tmp2018 := Call1(__e, SymFn(symshen_4lambda_1entry), V4175)


let__1205 := tmp2018
_ = let__1205

tmp2023 := Call1(__e, SymFn(symempty_2), let__1205)


var ifres2019 Obj

if True == tmp2023 {
tmp2020 := GValue(__e, sym_dproperty_1vector_d)

tmp2021 := Call3(__e, SymFn(symunput), V4175, symshen_4lambda_1form, tmp2020)


ifres2019 = tmp2021


} else {
tmp2022 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), let__1205)


ifres2019 = tmp2022


}

let__1206 := ifres2019
_ = let__1206

return V4175


panic("unreachable")
})

tmp2024 := Call2(__e, ns2_1set, symupdate_1lambda_1table, tmp2015)


_ = tmp2024

tmp2025 := MakeNative2(func(__e *ControlFlow, V4182, V4183 Obj) Obj {
tmp2049 := GEq(__e, MakeInteger(0), V4183)

if True == tmp2049 {
tmp2026 := GValue(__e, symshen_4_dspecial_d)

tmp2027 := Call2(__e, SymFn(symremove), V4182, tmp2026)


tmp2028 := PrimGuard2(__e, symset, PrimSet, symshen_4_dspecial_d, tmp2027)

_ = tmp2028

tmp2029 := GValue(__e, symshen_4_dextraspecial_d)

tmp2030 := Call2(__e, SymFn(symremove), V4182, tmp2029)


tmp2031 := PrimGuard2(__e, symset, PrimSet, symshen_4_dextraspecial_d, tmp2030)

_ = tmp2031

return V4182


} else {
tmp2047 := GEq(__e, MakeInteger(1), V4183)

if True == tmp2047 {
tmp2032 := GValue(__e, symshen_4_dspecial_d)

tmp2033 := Call2(__e, SymFn(symadjoin), V4182, tmp2032)


tmp2034 := PrimGuard2(__e, symset, PrimSet, symshen_4_dspecial_d, tmp2033)

_ = tmp2034

tmp2035 := GValue(__e, symshen_4_dextraspecial_d)

tmp2036 := Call2(__e, SymFn(symremove), V4182, tmp2035)


tmp2037 := PrimGuard2(__e, symset, PrimSet, symshen_4_dextraspecial_d, tmp2036)

_ = tmp2037

return V4182


} else {
tmp2045 := GEq(__e, MakeInteger(2), V4183)

if True == tmp2045 {
tmp2038 := GValue(__e, symshen_4_dspecial_d)

tmp2039 := Call2(__e, SymFn(symremove), V4182, tmp2038)


tmp2040 := PrimGuard2(__e, symset, PrimSet, symshen_4_dspecial_d, tmp2039)

_ = tmp2040

tmp2041 := GValue(__e, symshen_4_dextraspecial_d)

tmp2042 := Call2(__e, SymFn(symadjoin), V4182, tmp2041)


tmp2043 := PrimGuard2(__e, symset, PrimSet, symshen_4_dextraspecial_d, tmp2042)

_ = tmp2043

return V4182


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("specialise requires values of 0, 1 or 2\n"))
}


}


}


panic("unreachable")
})

tmp2050 := Call2(__e, ns2_1set, symspecialise, tmp2025)


_ = tmp2050

tmp2051 := MakeNative1(func(__e *ControlFlow, V4184 Obj) Obj {
tmp2052 := GValue(__e, sym_dabsolute_d)

tmp2053 := GCons(__e, V4184, tmp2052)

return PrimGuard2(__e, symset, PrimSet, sym_dabsolute_d, tmp2053)


panic("unreachable")
})

tmp2054 := Call2(__e, ns2_1set, symabsolute, tmp2051)


_ = tmp2054

tmp2055 := MakeNative1(func(__e *ControlFlow, V4185 Obj) Obj {
tmp2056 := GValue(__e, sym_dabsolute_d)

tmp2057 := Call2(__e, SymFn(symremove), V4185, tmp2056)


return PrimGuard2(__e, symset, PrimSet, sym_dabsolute_d, tmp2057)


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symunabsolute, tmp2055)




panic("unreachable")
})

