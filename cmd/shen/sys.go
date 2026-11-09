package main

import . "github.com/pyrex41/shen-go/kl"

var SysMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp1210 := MakeNative1(func(__e *ControlFlow, V3830 Obj) Obj {
return TailCall0(__e, V3830)
panic("unreachable")
})

tmp1211 := Call2(__e, ns2_1set, symthaw, tmp1210)


_ = tmp1211

tmp1212 := MakeNative1(func(__e *ControlFlow, V3831 Obj) Obj {
tmp1213 := Call1(__e, SymFn(symmacroexpand), V3831)


tmp1214 := Call1(__e, SymFn(symshen_4find_1types), V3831)


tmp1215 := Call2(__e, SymFn(symshen_4process_1applications), tmp1213, tmp1214)


tmp1216 := Call1(__e, SymFn(symshen_4shen_1_6kl), tmp1215)


return TailCall1(__e, SymFn(symeval_1kl), tmp1216)


panic("unreachable")
})

tmp1217 := Call2(__e, ns2_1set, symeval, tmp1212)


_ = tmp1217

tmp1218 := MakeNative1(func(__e *ControlFlow, V3832 Obj) Obj {
tmp1225 := GEq(__e, symnull, V3832)

if True == tmp1225 {
return Nil
} else {
tmp1219 := MakeNative0(func(__e *ControlFlow) Obj {
tmp1220 := GValue(__e, sym_dproperty_1vector_d)

return TailCall3(__e, SymFn(symget), V3832, symshen_4external_1symbols, tmp1220)


panic("unreachable")
})

tmp1221 := MakeNative1(func(__e *ControlFlow, Z3833 Obj) Obj {
tmp1222 := Call3(__e, SymFn(symshen_4app), V3832, MakeString(" does not exist.\n;"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("package "), tmp1222))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("package "), tmp1222)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp1219, tmp1221)


}


panic("unreachable")
})

tmp1226 := Call2(__e, ns2_1set, symexternal, tmp1218)


_ = tmp1226

tmp1227 := MakeNative1(func(__e *ControlFlow, V3834 Obj) Obj {
tmp1234 := GEq(__e, symnull, V3834)

if True == tmp1234 {
return Nil
} else {
tmp1228 := MakeNative0(func(__e *ControlFlow) Obj {
tmp1229 := GValue(__e, sym_dproperty_1vector_d)

return TailCall3(__e, SymFn(symget), V3834, symshen_4internal_1symbols, tmp1229)


panic("unreachable")
})

tmp1230 := MakeNative1(func(__e *ControlFlow, Z3835 Obj) Obj {
tmp1231 := Call3(__e, SymFn(symshen_4app), V3834, MakeString(" does not exist.\n;"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("package "), tmp1231))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("package "), tmp1231)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp1228, tmp1230)


}


panic("unreachable")
})

tmp1235 := Call2(__e, ns2_1set, syminternal, tmp1227)


_ = tmp1235

tmp1236 := MakeNative2(func(__e *ControlFlow, V3836, V3837 Obj) Obj {
tmp1238 := Call1(__e, V3836, V3837)


if True == tmp1238 {
return TailCall0(__e, SymFn(symfail))
} else {
return V3837
}


panic("unreachable")
})

tmp1239 := Call2(__e, ns2_1set, symfail_1if, tmp1236)


_ = tmp1239

tmp1240 := MakeNative2(func(__e *ControlFlow, V3838, V3839 Obj) Obj {
return PrimGuard2(__e, symcn, PrimStringConcat, V3838, V3839)
panic("unreachable")
})

tmp1241 := Call2(__e, ns2_1set, sym_8s, tmp1240)


_ = tmp1241

tmp1242 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dtc_d)
panic("unreachable")
})

tmp1243 := Call2(__e, ns2_1set, symtc_2, tmp1242)


_ = tmp1243

tmp1244 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_doccurs_d)
panic("unreachable")
})

tmp1245 := Call2(__e, ns2_1set, symoccurs_2, tmp1244)


_ = tmp1245

tmp1246 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dfactorise_2_d)
panic("unreachable")
})

tmp1247 := Call2(__e, ns2_1set, symfactorise_2, tmp1246)


_ = tmp1247

tmp1248 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dtracking_d)
panic("unreachable")
})

tmp1249 := Call2(__e, ns2_1set, symtracked, tmp1248)


_ = tmp1249

tmp1250 := MakeNative1(func(__e *ControlFlow, V3840 Obj) Obj {
tmp1251 := MakeNative0(func(__e *ControlFlow) Obj {
tmp1252 := GValue(__e, sym_dproperty_1vector_d)

return TailCall3(__e, SymFn(symget), V3840, symshen_4source, tmp1252)


panic("unreachable")
})

tmp1253 := MakeNative1(func(__e *ControlFlow, Z3841 Obj) Obj {
tmp1254 := Call3(__e, SymFn(symshen_4app), V3840, MakeString(" not found.\n"), symshen_4a)


return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp1254)


panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp1251, tmp1253)


panic("unreachable")
})

tmp1255 := Call2(__e, ns2_1set, symps, tmp1250)


_ = tmp1255

tmp1256 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dstinput_d)
panic("unreachable")
})

tmp1257 := Call2(__e, ns2_1set, symstinput, tmp1256)


_ = tmp1257

tmp1258 := MakeNative1(func(__e *ControlFlow, V3842 Obj) Obj {
tmp1260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(GAdd(__e, V3842, MakeInteger(1)))
}
__typedArg0 := GAdd(__e, V3842, MakeInteger(1))
return Call1(__e, SymFn(symabsvector), __typedArg0)
})()

let__1164 := tmp1260
_ = let__1164

tmp1261 := GVecSet(__e, let__1164, MakeInteger(0), V3842)

let__1165 := tmp1261
_ = let__1165

tmp1265 := GEq(__e, V3842, MakeInteger(0))

var ifres1262 Obj

if True == tmp1265 {
ifres1262 = let__1165


} else {
tmp1263 := Call0(__e, SymFn(symfail))


tmp1264 := Call4(__e, SymFn(symshen_4fillvector), let__1165, MakeInteger(1), V3842, tmp1263)


ifres1262 = tmp1264


}

let__1166 := ifres1262
_ = let__1166

return let__1166


panic("unreachable")
})

tmp1266 := Call2(__e, ns2_1set, symvector, tmp1258)


_ = tmp1266

tmp1267 := (func() Obj {
var __self Obj
__self = MakeNative4(func(__e *ControlFlow, __self1, __self2, __self3, __self4 Obj) Obj {
__selftop:
V3847 := __self1
_ = V3847
V3848 := __self2
_ = V3848
V3849 := __self3
_ = V3849
V3850 := __self4
_ = V3850
tmp1271 := GEq(__e, V3848, V3849)

if True == tmp1271 {
return GVecSet(__e, V3847, V3849, V3850)
} else {
tmp1268 := GVecSet(__e, V3847, V3848, V3850)

if SymFn(symshen_4fillvector) == __self {
__self1, __self2, __self3, __self4 = tmp1268, GAdd(__e, MakeInteger(1), V3848), V3849, V3850
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4fillvector), tmp1268, GAdd(__e, MakeInteger(1), V3848), V3849, V3850)


}


panic("unreachable")
})
return __self
})()

tmp1272 := Call2(__e, ns2_1set, symshen_4fillvector, tmp1267)


_ = tmp1272

tmp1273 := MakeNative1(func(__e *ControlFlow, V3851 Obj) Obj {
tmp1280 := GVectorP(__e, V3851)

if True == tmp1280 {
tmp1274 := MakeNative0(func(__e *ControlFlow) Obj {
return GVecGet(__e, V3851, MakeInteger(0))
panic("unreachable")
})

tmp1275 := MakeNative1(func(__e *ControlFlow, Z3853 Obj) Obj {
return MakeInteger(-1)
panic("unreachable")
})

tmp1276 := Call2(__e, try_1catch, tmp1274, tmp1275)


let__1167 := tmp1276
_ = let__1167

tmp1278 := PrimGuard1(__e, symnumber_2, PrimIsNumber, let__1167)

if True == tmp1278 {
return GGe(__e, let__1167, MakeInteger(0))
} else {
return False
}


} else {
return False
}


panic("unreachable")
})

tmp1281 := Call2(__e, ns2_1set, symvector_2, tmp1273)


_ = tmp1281

tmp1282 := MakeNative3(func(__e *ControlFlow, V3854, V3855, V3856 Obj) Obj {
tmp1284 := GEq(__e, V3855, MakeInteger(0))

if True == tmp1284 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("cannot access 0th element of a vector\n"))
} else {
return GVecSet(__e, V3854, V3855, V3856)
}


panic("unreachable")
})

tmp1285 := Call2(__e, ns2_1set, symvector_1_6, tmp1282)


_ = tmp1285

tmp1286 := MakeNative2(func(__e *ControlFlow, V3857, V3858 Obj) Obj {
tmp1292 := GEq(__e, V3858, MakeInteger(0))

if True == tmp1292 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("cannot access 0th element of a vector\n"))
} else {
tmp1287 := GVecGet(__e, V3857, V3858)

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

tmp1294 := MakeNative1(func(__e *ControlFlow, V3860 Obj) Obj {
tmp1296 := PrimGuard1(__e, syminteger_2, PrimIsInteger, V3860)

if True == tmp1296 {
return GGe(__e, V3860, MakeInteger(0))
} else {
return False
}


panic("unreachable")
})

tmp1297 := Call2(__e, ns2_1set, symshen_4posint_2, tmp1294)


_ = tmp1297

tmp1298 := MakeNative1(func(__e *ControlFlow, V3861 Obj) Obj {
return GVecGet(__e, V3861, MakeInteger(0))
panic("unreachable")
})

tmp1299 := Call2(__e, ns2_1set, symlimit, tmp1298)


_ = tmp1299

tmp1300 := MakeNative1(func(__e *ControlFlow, V3862 Obj) Obj {
tmp1325 := Call1(__e, SymFn(symboolean_2), V3862)


var ifres1315 Obj

if True == tmp1325 {
ifres1315 = True


} else {
tmp1324 := PrimGuard1(__e, symnumber_2, PrimIsNumber, V3862)

var ifres1316 Obj

if True == tmp1324 {
ifres1316 = True


} else {
tmp1323 := PrimGuard1(__e, symstring_2, PrimIsString, V3862)

var ifres1317 Obj

if True == tmp1323 {
ifres1317 = True


} else {
tmp1322 := GConsP(__e, V3862)

var ifres1318 Obj

if True == tmp1322 {
ifres1318 = True


} else {
tmp1321 := Call1(__e, SymFn(symempty_2), V3862)


var ifres1319 Obj

if True == tmp1321 {
ifres1319 = True


} else {
tmp1320 := Call1(__e, SymFn(symvector_2), V3862)


ifres1319 = tmp1320


}

ifres1318 = ifres1319


}

ifres1317 = ifres1318


}

ifres1316 = ifres1317


}

ifres1315 = ifres1316


}

if True == ifres1315 {
return False
} else {
tmp1305 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp1306 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp1307 := PrimGuard1(__e, symintern, PrimIntern, MakeString(","))

tmp1308 := GCons(__e, tmp1307, Nil)

tmp1309 := GCons(__e, tmp1306, tmp1308)

tmp1310 := GCons(__e, tmp1305, tmp1309)

tmp1311 := GCons(__e, sym_j, tmp1310)

tmp1312 := GCons(__e, sym_i, tmp1311)

tmp1313 := Call2(__e, SymFn(symelement_2), V3862, tmp1312)


if True == tmp1313 {
return True
} else {
tmp1301 := MakeNative0(func(__e *ControlFlow) Obj {
tmp1302 := PrimGuard1(__e, symstr, PrimStr, V3862)

let__1169 := tmp1302
_ = let__1169

return TailCall1(__e, SymFn(symshen_4analyse_1symbol_2), let__1169)


panic("unreachable")
})

tmp1303 := MakeNative1(func(__e *ControlFlow, Z3864 Obj) Obj {
return False
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp1301, tmp1303)


}


}


panic("unreachable")
})

tmp1326 := Call2(__e, ns2_1set, symsymbol_2, tmp1300)


_ = tmp1326

tmp1327 := MakeNative1(func(__e *ControlFlow, V3867 Obj) Obj {
tmp1334 := Call1(__e, SymFn(symshen_4_7string_2), V3867)


if True == tmp1334 {
tmp1330 := Call1(__e, SymFn(symhdstr), V3867)


tmp1331 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp1330)

tmp1332 := Call1(__e, SymFn(symshen_4alpha_2), tmp1331)


if True == tmp1332 {
return TailCall1(__e, SymFn(symshen_4alphanums_2), PrimGuard1(__e, symtlstr, PrimTailString, V3867))


} else {
return False
}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.analyse-symbol?"))
}


panic("unreachable")
})

tmp1335 := Call2(__e, ns2_1set, symshen_4analyse_1symbol_2, tmp1327)


_ = tmp1335

tmp1336 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V3870 := __self1
_ = V3870
tmp1347 := GEq(__e, MakeString(""), V3870)

if True == tmp1347 {
return True
} else {
tmp1345 := Call1(__e, SymFn(symshen_4_7string_2), V3870)


if True == tmp1345 {
tmp1337 := Call1(__e, SymFn(symhdstr), V3870)


tmp1338 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp1337)

let__1170 := tmp1338
_ = let__1170

tmp1343 := Call1(__e, SymFn(symshen_4alpha_2), let__1170)


var ifres1341 Obj

if True == tmp1343 {
ifres1341 = True


} else {
tmp1342 := Call1(__e, SymFn(symshen_4digit_2), let__1170)


ifres1341 = tmp1342


}

if True == ifres1341 {
if SymFn(symshen_4alphanums_2) == __self {
__self1 = PrimGuard1(__e, symtlstr, PrimTailString, V3870)
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4alphanums_2), PrimGuard1(__e, symtlstr, PrimTailString, V3870))


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

tmp1348 := Call2(__e, ns2_1set, symshen_4alphanums_2, tmp1336)


_ = tmp1348

tmp1349 := MakeNative1(func(__e *ControlFlow, V3872 Obj) Obj {
tmp1358 := Call1(__e, SymFn(symboolean_2), V3872)


var ifres1354 Obj

if True == tmp1358 {
ifres1354 = True


} else {
tmp1357 := PrimGuard1(__e, symnumber_2, PrimIsNumber, V3872)

var ifres1355 Obj

if True == tmp1357 {
ifres1355 = True


} else {
tmp1356 := PrimGuard1(__e, symstring_2, PrimIsString, V3872)

ifres1355 = tmp1356


}

ifres1354 = ifres1355


}

if True == ifres1354 {
return False
} else {
tmp1350 := MakeNative0(func(__e *ControlFlow) Obj {
tmp1351 := PrimGuard1(__e, symstr, PrimStr, V3872)

let__1171 := tmp1351
_ = let__1171

return TailCall1(__e, SymFn(symshen_4analyse_1variable_2), let__1171)


panic("unreachable")
})

tmp1352 := MakeNative1(func(__e *ControlFlow, Z3874 Obj) Obj {
return False
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp1350, tmp1352)


}


panic("unreachable")
})

tmp1359 := Call2(__e, ns2_1set, symvariable_2, tmp1349)


_ = tmp1359

tmp1360 := MakeNative1(func(__e *ControlFlow, V3877 Obj) Obj {
tmp1367 := Call1(__e, SymFn(symshen_4_7string_2), V3877)


if True == tmp1367 {
tmp1363 := Call1(__e, SymFn(symhdstr), V3877)


tmp1364 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp1363)

tmp1365 := Call1(__e, SymFn(symshen_4uppercase_2), tmp1364)


if True == tmp1365 {
return TailCall1(__e, SymFn(symshen_4alphanums_2), PrimGuard1(__e, symtlstr, PrimTailString, V3877))


} else {
return False
}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.analyse-variable?"))
}


panic("unreachable")
})

tmp1368 := Call2(__e, ns2_1set, symshen_4analyse_1variable_2, tmp1360)


_ = tmp1368

tmp1369 := MakeNative1(func(__e *ControlFlow, V3878 Obj) Obj {
tmp1370 := GValue(__e, symshen_4_dgensym_d)

tmp1372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dgensym_d, GAdd(__e, MakeInteger(1), tmp1370))
}
__typedArg0 := symshen_4_dgensym_d
__typedArg1 := GAdd(__e, MakeInteger(1), tmp1370)
return Call2(__e, SymFn(symset), __typedArg0, __typedArg1)
})()

return TailCall2(__e, SymFn(symconcat), V3878, tmp1372)


panic("unreachable")
})

tmp1373 := Call2(__e, ns2_1set, symgensym, tmp1369)


_ = tmp1373

tmp1374 := MakeNative2(func(__e *ControlFlow, V3879, V3880 Obj) Obj {
tmp1375 := PrimGuard1(__e, symstr, PrimStr, V3879)

tmp1376 := PrimGuard1(__e, symstr, PrimStr, V3880)

return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(PrimGuard2(__e, symcn, PrimStringConcat, tmp1375, tmp1376))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, tmp1375, tmp1376)
return Call1(__e, SymFn(symintern), __typedArg0)
})()


panic("unreachable")
})

tmp1378 := Call2(__e, ns2_1set, symconcat, tmp1374)


_ = tmp1378

tmp1379 := MakeNative2(func(__e *ControlFlow, V3881, V3882 Obj) Obj {
tmp1380 := PrimGuard1(__e, symabsvector, PrimAbsvector, MakeInteger(3))

let__1172 := tmp1380
_ = let__1172

tmp1381 := GVecSet(__e, let__1172, MakeInteger(0), symshen_4tuple)

let__1173 := tmp1381
_ = let__1173

tmp1382 := GVecSet(__e, let__1172, MakeInteger(1), V3881)

let__1174 := tmp1382
_ = let__1174

tmp1383 := GVecSet(__e, let__1172, MakeInteger(2), V3882)

let__1175 := tmp1383
_ = let__1175

return let__1172


panic("unreachable")
})

tmp1384 := Call2(__e, ns2_1set, sym_8p, tmp1379)


_ = tmp1384

tmp1385 := MakeNative1(func(__e *ControlFlow, V3887 Obj) Obj {
return GVecGet(__e, V3887, MakeInteger(1))
panic("unreachable")
})

tmp1386 := Call2(__e, ns2_1set, symfst, tmp1385)


_ = tmp1386

tmp1387 := MakeNative1(func(__e *ControlFlow, V3888 Obj) Obj {
return GVecGet(__e, V3888, MakeInteger(2))
panic("unreachable")
})

tmp1388 := Call2(__e, ns2_1set, symsnd, tmp1387)


_ = tmp1388

tmp1389 := MakeNative1(func(__e *ControlFlow, V3889 Obj) Obj {
tmp1394 := GVectorP(__e, V3889)

if True == tmp1394 {
tmp1390 := MakeNative0(func(__e *ControlFlow) Obj {
return GVecGet(__e, V3889, MakeInteger(0))
panic("unreachable")
})

tmp1391 := MakeNative1(func(__e *ControlFlow, Z3890 Obj) Obj {
return symshen_4not_1tuple
panic("unreachable")
})

tmp1392 := Call2(__e, try_1catch, tmp1390, tmp1391)


return GEq(__e, symshen_4tuple, tmp1392)


} else {
return False
}


panic("unreachable")
})

tmp1395 := Call2(__e, ns2_1set, symtuple_2, tmp1389)


_ = tmp1395

tmp1396 := MakeNative2(func(__e *ControlFlow, V3895, V3896 Obj) Obj {
tmp1403 := GEq(__e, Nil, V3895)

if True == tmp1403 {
return V3896
} else {
tmp1401 := GConsP(__e, V3895)

if True == tmp1401 {
tmp1397 := GHd(__e, V3895)

tmp1398 := GTl(__e, V3895)

tmp1399 := Call2(__e, SymFn(symappend), tmp1398, V3896)


return GCons(__e, tmp1397, tmp1399)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to append a non-list"))
}


}


panic("unreachable")
})

tmp1404 := Call2(__e, ns2_1set, symappend, tmp1396)


_ = tmp1404

tmp1405 := MakeNative2(func(__e *ControlFlow, V3897, V3898 Obj) Obj {
tmp1406 := Call1(__e, SymFn(symlimit), V3898)


let__1176 := tmp1406
_ = let__1176

tmp1408 := Call1(__e, SymFn(symvector), GAdd(__e, let__1176, MakeInteger(1)))


let__1177 := tmp1408
_ = let__1177

tmp1409 := Call3(__e, SymFn(symvector_1_6), let__1177, MakeInteger(1), V3897)


let__1178 := tmp1409
_ = let__1178

tmp1411 := GEq(__e, let__1176, MakeInteger(0))

if True == tmp1411 {
return let__1178
} else {
return TailCall4(__e, SymFn(symshen_4_8v_1help), V3898, MakeInteger(1), let__1176, let__1178)
}


panic("unreachable")
})

tmp1412 := Call2(__e, ns2_1set, sym_8v, tmp1405)


_ = tmp1412

tmp1413 := (func() Obj {
var __self Obj
__self = MakeNative4(func(__e *ControlFlow, __self1, __self2, __self3, __self4 Obj) Obj {
__selftop:
V3903 := __self1
_ = V3903
V3904 := __self2
_ = V3904
V3905 := __self3
_ = V3905
V3906 := __self4
_ = V3906
tmp1419 := GEq(__e, V3904, V3905)

if True == tmp1419 {
return TailCall4(__e, SymFn(symshen_4copyfromvector), V3903, V3906, V3905, GAdd(__e, V3905, MakeInteger(1)))


} else {
tmp1415 := GAdd(__e, V3904, MakeInteger(1))

tmp1417 := Call4(__e, SymFn(symshen_4copyfromvector), V3903, V3906, V3904, GAdd(__e, V3904, MakeInteger(1)))


if SymFn(symshen_4_8v_1help) == __self {
__self1, __self2, __self3, __self4 = V3903, tmp1415, V3905, tmp1417
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4_8v_1help), V3903, tmp1415, V3905, tmp1417)


}


panic("unreachable")
})
return __self
})()

tmp1420 := Call2(__e, ns2_1set, symshen_4_8v_1help, tmp1413)


_ = tmp1420

tmp1421 := MakeNative4(func(__e *ControlFlow, V3907, V3908, V3909, V3910 Obj) Obj {
tmp1422 := MakeNative0(func(__e *ControlFlow) Obj {
tmp1423 := Call2(__e, SymFn(sym_5_1vector), V3907, V3909)


return TailCall3(__e, SymFn(symvector_1_6), V3908, V3910, tmp1423)


panic("unreachable")
})

tmp1424 := MakeNative1(func(__e *ControlFlow, Z3911 Obj) Obj {
return V3908
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp1422, tmp1424)


panic("unreachable")
})

tmp1425 := Call2(__e, ns2_1set, symshen_4copyfromvector, tmp1421)


_ = tmp1425

tmp1426 := MakeNative1(func(__e *ControlFlow, V3912 Obj) Obj {
tmp1427 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall2(__e, SymFn(sym_5_1vector), V3912, MakeInteger(1))
panic("unreachable")
})

tmp1428 := MakeNative1(func(__e *ControlFlow, Z3913 Obj) Obj {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("hdv needs a non-empty vector as an argument\n"))
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp1427, tmp1428)


panic("unreachable")
})

tmp1429 := Call2(__e, ns2_1set, symhdv, tmp1426)


_ = tmp1429

tmp1430 := MakeNative1(func(__e *ControlFlow, V3914 Obj) Obj {
tmp1431 := Call1(__e, SymFn(symlimit), V3914)


let__1179 := tmp1431
_ = let__1179

tmp1439 := GEq(__e, let__1179, MakeInteger(0))

if True == tmp1439 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("cannot take the tail of the empty vector\n"))
} else {
tmp1437 := GEq(__e, let__1179, MakeInteger(1))

if True == tmp1437 {
return TailCall1(__e, SymFn(symvector), MakeInteger(0))
} else {
tmp1433 := Call1(__e, SymFn(symvector), GSub(__e, let__1179, MakeInteger(1)))


let__1180 := tmp1433
_ = let__1180

tmp1435 := Call1(__e, SymFn(symvector), GSub(__e, let__1179, MakeInteger(1)))


return TailCall4(__e, SymFn(symshen_4tlv_1help), V3914, MakeInteger(2), let__1179, tmp1435)


}


}


panic("unreachable")
})

tmp1440 := Call2(__e, ns2_1set, symtlv, tmp1430)


_ = tmp1440

tmp1441 := (func() Obj {
var __self Obj
__self = MakeNative4(func(__e *ControlFlow, __self1, __self2, __self3, __self4 Obj) Obj {
__selftop:
V3918 := __self1
_ = V3918
V3919 := __self2
_ = V3919
V3920 := __self3
_ = V3920
V3921 := __self4
_ = V3921
tmp1447 := GEq(__e, V3919, V3920)

if True == tmp1447 {
return TailCall4(__e, SymFn(symshen_4copyfromvector), V3918, V3921, V3920, GSub(__e, V3920, MakeInteger(1)))


} else {
tmp1443 := GAdd(__e, V3919, MakeInteger(1))

tmp1445 := Call4(__e, SymFn(symshen_4copyfromvector), V3918, V3921, V3919, GSub(__e, V3919, MakeInteger(1)))


if SymFn(symshen_4tlv_1help) == __self {
__self1, __self2, __self3, __self4 = V3918, tmp1443, V3920, tmp1445
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4tlv_1help), V3918, tmp1443, V3920, tmp1445)


}


panic("unreachable")
})
return __self
})()

tmp1448 := Call2(__e, ns2_1set, symshen_4tlv_1help, tmp1441)


_ = tmp1448

tmp1449 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V3933 := __self1
_ = V3933
V3934 := __self2
_ = V3934
tmp1463 := GEq(__e, Nil, V3934)

if True == tmp1463 {
return Nil
} else {
tmp1461 := GConsP(__e, V3934)

var ifres1454 Obj

if True == tmp1461 {
tmp1459 := GHd(__e, V3934)

tmp1460 := GConsP(__e, tmp1459)

var ifres1455 Obj

if True == tmp1460 {
tmp1456 := GHd(__e, V3934)

tmp1457 := GHd(__e, tmp1456)

tmp1458 := GEq(__e, V3933, tmp1457)

ifres1455 = tmp1458


} else {
ifres1455 = False


}

ifres1454 = ifres1455


} else {
ifres1454 = False


}

if True == ifres1454 {
return GHd(__e, V3934)
} else {
tmp1452 := GConsP(__e, V3934)

if True == tmp1452 {
tmp1450 := GTl(__e, V3934)

if SymFn(symassoc) == __self {
__self1, __self2 = V3933, tmp1450
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symassoc), V3933, tmp1450)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to search a non-list with assoc\n"))
}


}


}


panic("unreachable")
})
return __self
})()

tmp1464 := Call2(__e, ns2_1set, symassoc, tmp1449)


_ = tmp1464

tmp1465 := MakeNative3(func(__e *ControlFlow, V3938, V3939, V3940 Obj) Obj {
tmp1486 := GEq(__e, Nil, V3940)

if True == tmp1486 {
tmp1466 := GCons(__e, V3938, V3939)

return GCons(__e, tmp1466, Nil)


} else {
tmp1484 := GConsP(__e, V3940)

var ifres1477 Obj

if True == tmp1484 {
tmp1482 := GHd(__e, V3940)

tmp1483 := GConsP(__e, tmp1482)

var ifres1478 Obj

if True == tmp1483 {
tmp1479 := GHd(__e, V3940)

tmp1480 := GHd(__e, tmp1479)

tmp1481 := GEq(__e, V3938, tmp1480)

ifres1478 = tmp1481


} else {
ifres1478 = False


}

ifres1477 = ifres1478


} else {
ifres1477 = False


}

if True == ifres1477 {
tmp1467 := GHd(__e, V3940)

tmp1468 := GHd(__e, tmp1467)

tmp1469 := GCons(__e, tmp1468, V3939)

tmp1470 := GTl(__e, V3940)

return GCons(__e, tmp1469, tmp1470)


} else {
tmp1475 := GConsP(__e, V3940)

if True == tmp1475 {
tmp1471 := GHd(__e, V3940)

tmp1472 := GTl(__e, V3940)

tmp1473 := Call3(__e, SymFn(symshen_4assoc_1set), V3938, V3939, tmp1472)


return GCons(__e, tmp1471, tmp1473)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4assoc_1set)
}


}


}


panic("unreachable")
})

tmp1487 := Call2(__e, ns2_1set, symshen_4assoc_1set, tmp1465)


_ = tmp1487

tmp1488 := MakeNative2(func(__e *ControlFlow, V3944, V3945 Obj) Obj {
tmp1504 := GEq(__e, Nil, V3945)

if True == tmp1504 {
return Nil
} else {
tmp1502 := GConsP(__e, V3945)

var ifres1495 Obj

if True == tmp1502 {
tmp1500 := GHd(__e, V3945)

tmp1501 := GConsP(__e, tmp1500)

var ifres1496 Obj

if True == tmp1501 {
tmp1497 := GHd(__e, V3945)

tmp1498 := GHd(__e, tmp1497)

tmp1499 := GEq(__e, V3944, tmp1498)

ifres1496 = tmp1499


} else {
ifres1496 = False


}

ifres1495 = ifres1496


} else {
ifres1495 = False


}

if True == ifres1495 {
return GTl(__e, V3945)
} else {
tmp1493 := GConsP(__e, V3945)

if True == tmp1493 {
tmp1489 := GHd(__e, V3945)

tmp1490 := GTl(__e, V3945)

tmp1491 := Call2(__e, SymFn(symshen_4assoc_1rm), V3944, tmp1490)


return GCons(__e, tmp1489, tmp1491)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4assoc_1rm)
}


}


}


panic("unreachable")
})

tmp1505 := Call2(__e, ns2_1set, symshen_4assoc_1rm, tmp1488)


_ = tmp1505

tmp1506 := MakeNative1(func(__e *ControlFlow, V3948 Obj) Obj {
tmp1510 := GEq(__e, True, V3948)

if True == tmp1510 {
return True
} else {
tmp1508 := GEq(__e, False, V3948)

if True == tmp1508 {
return True
} else {
return False
}


}


panic("unreachable")
})

tmp1511 := Call2(__e, ns2_1set, symboolean_2, tmp1506)


_ = tmp1511

tmp1512 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V3949 := __self1
_ = V3949
tmp1517 := GEq(__e, MakeInteger(0), V3949)

if True == tmp1517 {
return MakeInteger(0)
} else {
tmp1513 := Call0(__e, SymFn(symstoutput))


tmp1514 := Call2(__e, SymFn(sympr), MakeString("\n"), tmp1513)


_ = tmp1514

if SymFn(symnl) == __self {
__self1 = GSub(__e, V3949, MakeInteger(1))
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symnl), GSub(__e, V3949, MakeInteger(1)))


}


panic("unreachable")
})
return __self
})()

tmp1518 := Call2(__e, ns2_1set, symnl, tmp1512)


_ = tmp1518

tmp1519 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V3956 := __self1
_ = V3956
V3957 := __self2
_ = V3957
tmp1530 := GEq(__e, Nil, V3956)

if True == tmp1530 {
return Nil
} else {
tmp1528 := GConsP(__e, V3956)

if True == tmp1528 {
tmp1525 := GHd(__e, V3956)

tmp1526 := Call2(__e, SymFn(symelement_2), tmp1525, V3957)


if True == tmp1526 {
tmp1520 := GTl(__e, V3956)

if SymFn(symdifference) == __self {
__self1, __self2 = tmp1520, V3957
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symdifference), tmp1520, V3957)


} else {
tmp1521 := GHd(__e, V3956)

tmp1522 := GTl(__e, V3956)

tmp1523 := Call2(__e, SymFn(symdifference), tmp1522, V3957)


return GCons(__e, tmp1521, tmp1523)


}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to find the difference with a non-list\n"))
}


}


panic("unreachable")
})
return __self
})()

tmp1531 := Call2(__e, ns2_1set, symdifference, tmp1519)


_ = tmp1531

tmp1532 := MakeNative2(func(__e *ControlFlow, V3958, V3959 Obj) Obj {
return V3959
panic("unreachable")
})

tmp1533 := Call2(__e, ns2_1set, symdo, tmp1532)


_ = tmp1533

tmp1534 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V3971 := __self1
_ = V3971
V3972 := __self2
_ = V3972
tmp1544 := GEq(__e, Nil, V3972)

if True == tmp1544 {
return False
} else {
tmp1542 := GConsP(__e, V3972)

var ifres1539 Obj

if True == tmp1542 {
tmp1540 := GHd(__e, V3972)

tmp1541 := GEq(__e, V3971, tmp1540)

ifres1539 = tmp1541


} else {
ifres1539 = False


}

if True == ifres1539 {
return True
} else {
tmp1537 := GConsP(__e, V3972)

if True == tmp1537 {
tmp1535 := GTl(__e, V3972)

if SymFn(symelement_2) == __self {
__self1, __self2 = V3971, tmp1535
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symelement_2), V3971, tmp1535)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to find an element in a non-list\n"))
}


}


}


panic("unreachable")
})
return __self
})()

tmp1545 := Call2(__e, ns2_1set, symelement_2, tmp1534)


_ = tmp1545

tmp1546 := MakeNative1(func(__e *ControlFlow, V3975 Obj) Obj {
tmp1548 := GEq(__e, Nil, V3975)

if True == tmp1548 {
return True
} else {
return False
}


panic("unreachable")
})

tmp1549 := Call2(__e, ns2_1set, symempty_2, tmp1546)


_ = tmp1549

tmp1550 := MakeNative2(func(__e *ControlFlow, V3976, V3977 Obj) Obj {
tmp1551 := Call1(__e, V3976, V3977)


return TailCall3(__e, SymFn(symshen_4fix_1help), V3976, V3977, tmp1551)


panic("unreachable")
})

tmp1552 := Call2(__e, ns2_1set, symfix, tmp1550)


_ = tmp1552

tmp1553 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V3983 := __self1
_ = V3983
V3984 := __self2
_ = V3984
V3985 := __self3
_ = V3985
tmp1556 := GEq(__e, V3984, V3985)

if True == tmp1556 {
return V3985
} else {
tmp1554 := Call1(__e, V3983, V3985)


if SymFn(symshen_4fix_1help) == __self {
__self1, __self2, __self3 = V3983, V3985, tmp1554
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4fix_1help), V3983, V3985, tmp1554)


}


panic("unreachable")
})
return __self
})()

tmp1557 := Call2(__e, ns2_1set, symshen_4fix_1help, tmp1553)


_ = tmp1557

tmp1558 := MakeNative4(func(__e *ControlFlow, V3986, V3987, V3988, V3989 Obj) Obj {
tmp1559 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall2(__e, SymFn(symshen_4_5_1dict), V3989, V3986)
panic("unreachable")
})

tmp1560 := MakeNative1(func(__e *ControlFlow, Z3991 Obj) Obj {
return Nil
panic("unreachable")
})

tmp1561 := Call2(__e, try_1catch, tmp1559, tmp1560)


let__1181 := tmp1561
_ = let__1181

tmp1562 := Call3(__e, SymFn(symshen_4assoc_1set), V3987, V3988, let__1181)


let__1182 := tmp1562
_ = let__1182

tmp1563 := Call3(__e, SymFn(symshen_4dict_1_6), V3989, V3986, let__1182)


let__1183 := tmp1563
_ = let__1183

return V3988


panic("unreachable")
})

tmp1564 := Call2(__e, ns2_1set, symput, tmp1558)


_ = tmp1564

tmp1565 := MakeNative3(func(__e *ControlFlow, V3994, V3995, V3996 Obj) Obj {
tmp1566 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall2(__e, SymFn(symshen_4_5_1dict), V3996, V3994)
panic("unreachable")
})

tmp1567 := MakeNative1(func(__e *ControlFlow, Z3998 Obj) Obj {
return Nil
panic("unreachable")
})

tmp1568 := Call2(__e, try_1catch, tmp1566, tmp1567)


let__1184 := tmp1568
_ = let__1184

tmp1569 := Call2(__e, SymFn(symshen_4assoc_1rm), V3995, let__1184)


let__1185 := tmp1569
_ = let__1185

tmp1570 := Call3(__e, SymFn(symshen_4dict_1_6), V3996, V3994, let__1185)


let__1186 := tmp1570
_ = let__1186

return V3994


panic("unreachable")
})

tmp1571 := Call2(__e, ns2_1set, symunput, tmp1565)


_ = tmp1571

tmp1572 := MakeNative3(func(__e *ControlFlow, V4001, V4002, V4003 Obj) Obj {
tmp1573 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall2(__e, SymFn(symshen_4_5_1dict), V4003, V4001)
panic("unreachable")
})

tmp1574 := MakeNative1(func(__e *ControlFlow, Z4005 Obj) Obj {
tmp1575 := Call3(__e, SymFn(symshen_4app), V4002, MakeString("\n"), symshen_4s)


tmp1577 := Call3(__e, SymFn(symshen_4app), V4001, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" has no attributes: "), tmp1575), symshen_4a)


return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp1577)


panic("unreachable")
})

tmp1578 := Call2(__e, try_1catch, tmp1573, tmp1574)


let__1187 := tmp1578
_ = let__1187

tmp1579 := Call2(__e, SymFn(symassoc), V4002, let__1187)


let__1188 := tmp1579
_ = let__1188

tmp1585 := Call1(__e, SymFn(symempty_2), let__1188)


if True == tmp1585 {
tmp1580 := Call3(__e, SymFn(symshen_4app), V4001, MakeString("\n"), symshen_4s)


tmp1582 := Call3(__e, SymFn(symshen_4app), V4002, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" not found for "), tmp1580), symshen_4s)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("attribute "), tmp1582))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("attribute "), tmp1582)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


} else {
return GTl(__e, let__1188)
}


panic("unreachable")
})

tmp1586 := Call2(__e, ns2_1set, symget, tmp1572)


_ = tmp1586

tmp1587 := MakeNative2(func(__e *ControlFlow, V4007, V4008 Obj) Obj {
tmp1588 := Call1(__e, SymFn(symshen_4hashkey), V4007)


tmp1589 := Call2(__e, SymFn(symshen_4mod), tmp1588, V4008)


let__1189 := tmp1589
_ = let__1189

tmp1591 := GEq(__e, let__1189, MakeInteger(0))

if True == tmp1591 {
return MakeInteger(1)
} else {
return let__1189
}


panic("unreachable")
})

tmp1592 := Call2(__e, ns2_1set, symhash, tmp1587)


_ = tmp1592

tmp1593 := MakeNative1(func(__e *ControlFlow, V4010 Obj) Obj {
tmp1594 := MakeNative1(func(__e *ControlFlow, Z4012 Obj) Obj {
return PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, Z4012)
panic("unreachable")
})

tmp1595 := Call1(__e, SymFn(symexplode), V4010)


tmp1596 := Call2(__e, SymFn(symmap), tmp1594, tmp1595)


let__1190 := tmp1596
_ = let__1190

return TailCall2(__e, SymFn(symshen_4prodbutzero), let__1190, MakeInteger(1))


panic("unreachable")
})

tmp1597 := Call2(__e, ns2_1set, symshen_4hashkey, tmp1593)


_ = tmp1597

tmp1598 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4013 := __self1
_ = V4013
V4014 := __self2
_ = V4014
tmp1616 := GEq(__e, Nil, V4013)

if True == tmp1616 {
return V4014
} else {
tmp1614 := GConsP(__e, V4013)

var ifres1611 Obj

if True == tmp1614 {
tmp1612 := GHd(__e, V4013)

tmp1613 := GEq(__e, MakeInteger(0), tmp1612)

ifres1611 = tmp1613


} else {
ifres1611 = False


}

if True == ifres1611 {
tmp1599 := GTl(__e, V4013)

if SymFn(symshen_4prodbutzero) == __self {
__self1, __self2 = tmp1599, V4014
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4prodbutzero), tmp1599, V4014)


} else {
tmp1609 := GConsP(__e, V4013)

if True == tmp1609 {
if True == GGt(__e, V4014, MakeNumber(1e+10)) {
tmp1600 := GTl(__e, V4013)

tmp1601 := GHd(__e, V4013)

if SymFn(symshen_4prodbutzero) == __self {
__self1, __self2 = tmp1600, GAdd(__e, V4014, tmp1601)
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4prodbutzero), tmp1600, GAdd(__e, V4014, tmp1601))


} else {
tmp1603 := GTl(__e, V4013)

tmp1604 := GHd(__e, V4013)

if SymFn(symshen_4prodbutzero) == __self {
__self1, __self2 = tmp1603, GMul(__e, V4014, tmp1604)
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4prodbutzero), tmp1603, GMul(__e, V4014, tmp1604))


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

tmp1617 := Call2(__e, ns2_1set, symshen_4prodbutzero, tmp1598)


_ = tmp1617

tmp1618 := MakeNative2(func(__e *ControlFlow, V4015, V4016 Obj) Obj {
tmp1619 := GCons(__e, V4016, Nil)

tmp1620 := Call2(__e, SymFn(symshen_4multiples), V4015, tmp1619)


return TailCall2(__e, SymFn(symshen_4modh), V4015, tmp1620)


panic("unreachable")
})

tmp1621 := Call2(__e, ns2_1set, symshen_4mod, tmp1618)


_ = tmp1621

tmp1622 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4021 := __self1
_ = V4021
V4022 := __self2
_ = V4022
tmp1632 := GConsP(__e, V4022)

var ifres1629 Obj

if True == tmp1632 {
tmp1630 := GHd(__e, V4022)

ifres1629 = GGt(__e, tmp1630, V4021)


} else {
ifres1629 = False


}

if True == ifres1629 {
return GTl(__e, V4022)
} else {
tmp1627 := GConsP(__e, V4022)

if True == tmp1627 {
tmp1623 := GHd(__e, V4022)

tmp1625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(GMul(__e, MakeInteger(2), tmp1623), V4022)
}
__typedArg0 := GMul(__e, MakeInteger(2), tmp1623)
__typedArg1 := V4022
return Call2(__e, SymFn(symcons), __typedArg0, __typedArg1)
})()

if SymFn(symshen_4multiples) == __self {
__self1, __self2 = V4021, tmp1625
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4multiples), V4021, tmp1625)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.multiples"))
}


}


panic("unreachable")
})
return __self
})()

tmp1633 := Call2(__e, ns2_1set, symshen_4multiples, tmp1622)


_ = tmp1633

tmp1634 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4029 := __self1
_ = V4029
V4030 := __self2
_ = V4030
tmp1651 := GEq(__e, MakeInteger(0), V4029)

if True == tmp1651 {
return MakeInteger(0)
} else {
tmp1649 := GEq(__e, Nil, V4030)

if True == tmp1649 {
return V4029
} else {
tmp1647 := GConsP(__e, V4030)

var ifres1644 Obj

if True == tmp1647 {
tmp1645 := GHd(__e, V4030)

ifres1644 = GGt(__e, tmp1645, V4029)


} else {
ifres1644 = False


}

if True == ifres1644 {
tmp1637 := GTl(__e, V4030)

tmp1638 := Call1(__e, SymFn(symempty_2), tmp1637)


if True == tmp1638 {
return V4029
} else {
tmp1635 := GTl(__e, V4030)

if SymFn(symshen_4modh) == __self {
__self1, __self2 = V4029, tmp1635
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4modh), V4029, tmp1635)


}


} else {
tmp1642 := GConsP(__e, V4030)

if True == tmp1642 {
tmp1639 := GHd(__e, V4030)

if SymFn(symshen_4modh) == __self {
__self1, __self2 = GSub(__e, V4029, tmp1639), V4030
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4modh), GSub(__e, V4029, tmp1639), V4030)


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

tmp1652 := Call2(__e, ns2_1set, symshen_4modh, tmp1634)


_ = tmp1652

tmp1653 := MakeNative1(func(__e *ControlFlow, V4033 Obj) Obj {
tmp1660 := GEq(__e, Nil, V4033)

if True == tmp1660 {
return MakeInteger(0)
} else {
tmp1658 := GConsP(__e, V4033)

if True == tmp1658 {
tmp1654 := GHd(__e, V4033)

tmp1655 := GTl(__e, V4033)

tmp1656 := Call1(__e, SymFn(symsum), tmp1655)


return GAdd(__e, tmp1654, tmp1656)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to sum a non-list\n"))
}


}


panic("unreachable")
})

tmp1661 := Call2(__e, ns2_1set, symsum, tmp1653)


_ = tmp1661

tmp1662 := MakeNative1(func(__e *ControlFlow, V4038 Obj) Obj {
tmp1664 := GConsP(__e, V4038)

if True == tmp1664 {
return GHd(__e, V4038)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("head expects a non-empty list\n"))
}


panic("unreachable")
})

tmp1665 := Call2(__e, ns2_1set, symhead, tmp1662)


_ = tmp1665

tmp1666 := MakeNative1(func(__e *ControlFlow, V4043 Obj) Obj {
tmp1668 := GConsP(__e, V4043)

if True == tmp1668 {
return GTl(__e, V4043)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("tail expects a non-empty list\n"))
}


panic("unreachable")
})

tmp1669 := Call2(__e, ns2_1set, symtail, tmp1666)


_ = tmp1669

tmp1670 := MakeNative1(func(__e *ControlFlow, V4044 Obj) Obj {
return PrimGuard2(__e, sympos, PrimPos, V4044, MakeInteger(0))
panic("unreachable")
})

tmp1671 := Call2(__e, ns2_1set, symhdstr, tmp1670)


_ = tmp1671

tmp1672 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4051 := __self1
_ = V4051
V4052 := __self2
_ = V4052
tmp1683 := GEq(__e, Nil, V4051)

if True == tmp1683 {
return Nil
} else {
tmp1681 := GConsP(__e, V4051)

if True == tmp1681 {
tmp1678 := GHd(__e, V4051)

tmp1679 := Call2(__e, SymFn(symelement_2), tmp1678, V4052)


if True == tmp1679 {
tmp1673 := GHd(__e, V4051)

tmp1674 := GTl(__e, V4051)

tmp1675 := Call2(__e, SymFn(symintersection), tmp1674, V4052)


return GCons(__e, tmp1673, tmp1675)


} else {
tmp1676 := GTl(__e, V4051)

if SymFn(symintersection) == __self {
__self1, __self2 = tmp1676, V4052
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symintersection), tmp1676, V4052)


}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to find the intersection with a non-list\n"))
}


}


panic("unreachable")
})
return __self
})()

tmp1684 := Call2(__e, ns2_1set, symintersection, tmp1672)


_ = tmp1684

tmp1685 := MakeNative1(func(__e *ControlFlow, V4053 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4reverse_1help), V4053, Nil)
panic("unreachable")
})

tmp1686 := Call2(__e, ns2_1set, symreverse, tmp1685)


_ = tmp1686

tmp1687 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4058 := __self1
_ = V4058
V4059 := __self2
_ = V4059
tmp1694 := GEq(__e, Nil, V4058)

if True == tmp1694 {
return V4059
} else {
tmp1692 := GConsP(__e, V4058)

if True == tmp1692 {
tmp1688 := GTl(__e, V4058)

tmp1689 := GHd(__e, V4058)

tmp1690 := GCons(__e, tmp1689, V4059)

if SymFn(symshen_4reverse_1help) == __self {
__self1, __self2 = tmp1688, tmp1690
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4reverse_1help), tmp1688, tmp1690)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to reverse a non-list\n"))
}


}


panic("unreachable")
})
return __self
})()

tmp1695 := Call2(__e, ns2_1set, symshen_4reverse_1help, tmp1687)


_ = tmp1695

tmp1696 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4064 := __self1
_ = V4064
V4065 := __self2
_ = V4065
tmp1707 := GEq(__e, Nil, V4064)

if True == tmp1707 {
return V4065
} else {
tmp1705 := GConsP(__e, V4064)

if True == tmp1705 {
tmp1702 := GHd(__e, V4064)

tmp1703 := Call2(__e, SymFn(symelement_2), tmp1702, V4065)


if True == tmp1703 {
tmp1697 := GTl(__e, V4064)

if SymFn(symunion) == __self {
__self1, __self2 = tmp1697, V4065
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symunion), tmp1697, V4065)


} else {
tmp1698 := GHd(__e, V4064)

tmp1699 := GTl(__e, V4064)

tmp1700 := Call2(__e, SymFn(symunion), tmp1699, V4065)


return GCons(__e, tmp1698, tmp1700)


}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to find the union with a non-list\n"))
}


}


panic("unreachable")
})
return __self
})()

tmp1708 := Call2(__e, ns2_1set, symunion, tmp1696)


_ = tmp1708

tmp1709 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V4066 := __self1
_ = V4066
tmp1710 := Call1(__e, SymFn(symshen_4proc_1nl), V4066)


tmp1711 := Call0(__e, SymFn(symstoutput))


tmp1712 := Call2(__e, SymFn(sympr), tmp1710, tmp1711)


let__1191 := tmp1712
_ = let__1191

tmp1713 := Call0(__e, SymFn(symstoutput))


tmp1714 := Call2(__e, SymFn(sympr), MakeString(" (y/n) "), tmp1713)


let__1192 := tmp1714
_ = let__1192

tmp1715 := Call0(__e, SymFn(symstinput))


tmp1716 := Call1(__e, SymFn(symread), tmp1715)


tmp1717 := Call3(__e, SymFn(symshen_4app), tmp1716, MakeString(""), symshen_4s)


let__1193 := tmp1717
_ = let__1193

tmp1723 := GEq(__e, MakeString("y"), let__1193)

if True == tmp1723 {
return True
} else {
tmp1721 := GEq(__e, MakeString("n"), let__1193)

if True == tmp1721 {
return False
} else {
tmp1718 := Call0(__e, SymFn(symstoutput))


tmp1719 := Call2(__e, SymFn(sympr), MakeString("please answer y or n\n"), tmp1718)


_ = tmp1719

if SymFn(symy_1or_1n_2) == __self {
__self1 = V4066
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symy_1or_1n_2), V4066)


}


}


panic("unreachable")
})
return __self
})()

tmp1724 := Call2(__e, ns2_1set, symy_1or_1n_2, tmp1709)


_ = tmp1724

tmp1725 := MakeNative1(func(__e *ControlFlow, V4070 Obj) Obj {
if True == V4070 {
return False
} else {
return True
}
panic("unreachable")
})

tmp1727 := Call2(__e, ns2_1set, symnot, tmp1725)


_ = tmp1727

tmp1728 := MakeNative0(func(__e *ControlFlow) Obj {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString(""))
panic("unreachable")
})

tmp1729 := Call2(__e, ns2_1set, symabort, tmp1728)


_ = tmp1729

tmp1730 := MakeNative3(func(__e *ControlFlow, V4076, V4077, V4078 Obj) Obj {
tmp1738 := GEq(__e, V4077, V4078)

if True == tmp1738 {
return V4076
} else {
tmp1736 := GConsP(__e, V4078)

if True == tmp1736 {
tmp1731 := GHd(__e, V4078)

tmp1732 := Call3(__e, SymFn(symsubst), V4076, V4077, tmp1731)


tmp1733 := GTl(__e, V4078)

tmp1734 := Call3(__e, SymFn(symsubst), V4076, V4077, tmp1733)


return GCons(__e, tmp1732, tmp1734)


} else {
return V4078
}


}


panic("unreachable")
})

tmp1739 := Call2(__e, ns2_1set, symsubst, tmp1730)


_ = tmp1739

tmp1740 := MakeNative1(func(__e *ControlFlow, V4079 Obj) Obj {
tmp1741 := Call3(__e, SymFn(symshen_4app), V4079, MakeString(""), symshen_4a)


return TailCall1(__e, SymFn(symshen_4explode_1h), tmp1741)


panic("unreachable")
})

tmp1742 := Call2(__e, ns2_1set, symexplode, tmp1740)


_ = tmp1742

tmp1743 := MakeNative1(func(__e *ControlFlow, V4082 Obj) Obj {
tmp1750 := GEq(__e, MakeString(""), V4082)

if True == tmp1750 {
return Nil
} else {
tmp1748 := Call1(__e, SymFn(symshen_4_7string_2), V4082)


if True == tmp1748 {
tmp1744 := Call1(__e, SymFn(symhdstr), V4082)


tmp1746 := Call1(__e, SymFn(symshen_4explode_1h), PrimGuard1(__e, symtlstr, PrimTailString, V4082))


return GCons(__e, tmp1744, tmp1746)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in explode-h"))
}


}


panic("unreachable")
})

tmp1751 := Call2(__e, ns2_1set, symshen_4explode_1h, tmp1743)


_ = tmp1751

tmp1752 := MakeNative1(func(__e *ControlFlow, V4083 Obj) Obj {
tmp1755 := GEq(__e, V4083, MakeString(""))

var ifres1753 Obj

if True == tmp1755 {
ifres1753 = MakeString("")


} else {
tmp1754 := Call3(__e, SymFn(symshen_4app), V4083, MakeString("/"), symshen_4a)


ifres1753 = tmp1754


}

return PrimGuard2(__e, symset, PrimSet, sym_dhome_1directory_d, ifres1753)


panic("unreachable")
})

tmp1756 := Call2(__e, ns2_1set, symcd, tmp1752)


_ = tmp1756

tmp1757 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4084 := __self1
_ = V4084
V4085 := __self2
_ = V4085
tmp1764 := GEq(__e, Nil, V4085)

if True == tmp1764 {
return True
} else {
tmp1762 := GConsP(__e, V4085)

if True == tmp1762 {
tmp1758 := GHd(__e, V4085)

tmp1759 := Call1(__e, V4084, tmp1758)


let__1194 := tmp1759
_ = let__1194

tmp1760 := GTl(__e, V4085)

if SymFn(symshen_4for_1each) == __self {
__self1, __self2 = V4084, tmp1760
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4for_1each), V4084, tmp1760)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4for_1each)
}


}


panic("unreachable")
})
return __self
})()

tmp1765 := Call2(__e, ns2_1set, symshen_4for_1each, tmp1757)


_ = tmp1765

tmp1766 := MakeNative2(func(__e *ControlFlow, V4087, V4088 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4map_1h), V4087, V4088, Nil)
panic("unreachable")
})

tmp1767 := Call2(__e, ns2_1set, symmap, tmp1766)


_ = tmp1767

tmp1768 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V4089 := __self1
_ = V4089
V4090 := __self2
_ = V4090
V4091 := __self3
_ = V4091
tmp1776 := GEq(__e, Nil, V4090)

if True == tmp1776 {
return TailCall1(__e, SymFn(symreverse), V4091)
} else {
tmp1774 := GConsP(__e, V4090)

if True == tmp1774 {
tmp1769 := GTl(__e, V4090)

tmp1770 := GHd(__e, V4090)

tmp1771 := Call1(__e, V4089, tmp1770)


tmp1772 := GCons(__e, tmp1771, V4091)

if SymFn(symshen_4map_1h) == __self {
__self1, __self2, __self3 = V4089, tmp1769, tmp1772
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4map_1h), V4089, tmp1769, tmp1772)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4map_1h)
}


}


panic("unreachable")
})
return __self
})()

tmp1777 := Call2(__e, ns2_1set, symshen_4map_1h, tmp1768)


_ = tmp1777

tmp1778 := MakeNative1(func(__e *ControlFlow, V4092 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4length_1h), V4092, MakeInteger(0))
panic("unreachable")
})

tmp1779 := Call2(__e, ns2_1set, symlength, tmp1778)


_ = tmp1779

tmp1780 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4097 := __self1
_ = V4097
V4098 := __self2
_ = V4098
tmp1784 := GEq(__e, Nil, V4097)

if True == tmp1784 {
return V4098
} else {
tmp1781 := GTl(__e, V4097)

if SymFn(symshen_4length_1h) == __self {
__self1, __self2 = tmp1781, GAdd(__e, V4098, MakeInteger(1))
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4length_1h), tmp1781, GAdd(__e, V4098, MakeInteger(1)))


}


panic("unreachable")
})
return __self
})()

tmp1785 := Call2(__e, ns2_1set, symshen_4length_1h, tmp1780)


_ = tmp1785

tmp1786 := MakeNative2(func(__e *ControlFlow, V4104, V4105 Obj) Obj {
tmp1794 := GEq(__e, V4104, V4105)

if True == tmp1794 {
return MakeInteger(1)
} else {
tmp1792 := GConsP(__e, V4105)

if True == tmp1792 {
tmp1787 := GHd(__e, V4105)

tmp1788 := Call2(__e, SymFn(symoccurrences), V4104, tmp1787)


tmp1789 := GTl(__e, V4105)

tmp1790 := Call2(__e, SymFn(symoccurrences), V4104, tmp1789)


return GAdd(__e, tmp1788, tmp1790)


} else {
return MakeInteger(0)
}


}


panic("unreachable")
})

tmp1795 := Call2(__e, ns2_1set, symoccurrences, tmp1786)


_ = tmp1795

tmp1796 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4110 := __self1
_ = V4110
V4111 := __self2
_ = V4111
tmp1808 := GEq(__e, MakeInteger(1), V4110)

var ifres1806 Obj

if True == tmp1808 {
tmp1807 := GConsP(__e, V4111)

ifres1806 = tmp1807


} else {
ifres1806 = False


}

if True == ifres1806 {
return GHd(__e, V4111)
} else {
tmp1804 := GConsP(__e, V4111)

if True == tmp1804 {
tmp1797 := GSub(__e, V4110, MakeInteger(1))

tmp1798 := GTl(__e, V4111)

if SymFn(symnth) == __self {
__self1, __self2 = tmp1797, tmp1798
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symnth), tmp1797, tmp1798)


} else {
tmp1799 := Call3(__e, SymFn(symshen_4app), V4111, MakeString("\n"), symshen_4a)


tmp1801 := Call3(__e, SymFn(symshen_4app), V4110, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(", "), tmp1799), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("nth applied to "), tmp1801))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("nth applied to "), tmp1801)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


}


panic("unreachable")
})
return __self
})()

tmp1809 := Call2(__e, ns2_1set, symnth, tmp1796)


_ = tmp1809

tmp1810 := MakeNative1(func(__e *ControlFlow, V4112 Obj) Obj {
tmp1814 := PrimGuard1(__e, symnumber_2, PrimIsNumber, V4112)

if True == tmp1814 {
tmp1811 := Call1(__e, SymFn(symshen_4abs), V4112)


let__1195 := tmp1811
_ = let__1195

tmp1812 := Call2(__e, SymFn(symshen_4magless), let__1195, MakeInteger(1))


return TailCall2(__e, SymFn(symshen_4integer_1test_2), let__1195, tmp1812)


} else {
return False
}


panic("unreachable")
})

tmp1815 := Call2(__e, ns2_1set, syminteger_2, tmp1810)


_ = tmp1815

tmp1816 := MakeNative1(func(__e *ControlFlow, V4114 Obj) Obj {
if True == GGt(__e, V4114, MakeInteger(0)) {
return V4114
} else {
return GSub(__e, MakeInteger(0), V4114)
}


panic("unreachable")
})

tmp1819 := Call2(__e, ns2_1set, symshen_4abs, tmp1816)


_ = tmp1819

tmp1820 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4115 := __self1
_ = V4115
V4116 := __self2
_ = V4116
let__1196 := GMul(__e, V4116, MakeInteger(2))
_ = let__1196

if True == GGt(__e, let__1196, V4115) {
return V4116
} else {
if SymFn(symshen_4magless) == __self {
__self1, __self2 = V4115, let__1196
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4magless), V4115, let__1196)
}


panic("unreachable")
})
return __self
})()

tmp1824 := Call2(__e, ns2_1set, symshen_4magless, tmp1820)


_ = tmp1824

tmp1825 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V4121 := __self1
_ = V4121
V4122 := __self2
_ = V4122
tmp1832 := GEq(__e, MakeInteger(0), V4121)

if True == tmp1832 {
return True
} else {
if True == GGt(__e, MakeInteger(1), V4121) {
return False
} else {
let__1197 := GSub(__e, V4121, V4122)
_ = let__1197

if True == GGt(__e, MakeInteger(0), let__1197) {
return PrimGuard1(__e, syminteger_2, PrimIsInteger, V4121)
} else {
if SymFn(symshen_4integer_1test_2) == __self {
__self1, __self2 = let__1197, V4122
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4integer_1test_2), let__1197, V4122)
}


}


}


panic("unreachable")
})
return __self
})()

tmp1833 := Call2(__e, ns2_1set, symshen_4integer_1test_2, tmp1825)


_ = tmp1833

tmp1834 := MakeNative2(func(__e *ControlFlow, V4130, V4131 Obj) Obj {
tmp1842 := GEq(__e, Nil, V4131)

if True == tmp1842 {
return Nil
} else {
tmp1840 := GConsP(__e, V4131)

if True == tmp1840 {
tmp1835 := GHd(__e, V4131)

tmp1836 := Call1(__e, V4130, tmp1835)


tmp1837 := GTl(__e, V4131)

tmp1838 := Call2(__e, SymFn(symmapcan), V4130, tmp1837)


return TailCall2(__e, SymFn(symappend), tmp1836, tmp1838)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("attempt to mapcan over a non-list\n"))
}


}


panic("unreachable")
})

tmp1843 := Call2(__e, ns2_1set, symmapcan, tmp1834)


_ = tmp1843

tmp1844 := MakeNative2(func(__e *ControlFlow, V4137, V4138 Obj) Obj {
tmp1846 := GEq(__e, V4137, V4138)

if True == tmp1846 {
return True
} else {
return False
}


panic("unreachable")
})

tmp1847 := Call2(__e, ns2_1set, sym_a_a, tmp1844)


_ = tmp1847

tmp1848 := MakeNative1(func(__e *ControlFlow, V4139 Obj) Obj {
tmp1855 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V4139)

if True == tmp1855 {
tmp1849 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, V4139)
panic("unreachable")
})

tmp1850 := MakeNative1(func(__e *ControlFlow, Z4141 Obj) Obj {
return symshen_4this_1symbol_1is_1unbound
panic("unreachable")
})

tmp1851 := Call2(__e, try_1catch, tmp1849, tmp1850)


let__1198 := tmp1851
_ = let__1198

tmp1853 := GEq(__e, let__1198, symshen_4this_1symbol_1is_1unbound)

if True == tmp1853 {
return False
} else {
return True
}


} else {
return False
}


panic("unreachable")
})

tmp1856 := Call2(__e, ns2_1set, symbound_2, tmp1848)


_ = tmp1856

tmp1857 := MakeNative1(func(__e *ControlFlow, V4142 Obj) Obj {
tmp1863 := GEq(__e, MakeString(""), V4142)

if True == tmp1863 {
return Nil
} else {
tmp1858 := PrimGuard2(__e, sympos, PrimPos, V4142, MakeInteger(0))

tmp1859 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp1858)

tmp1861 := Call1(__e, SymFn(symshen_4string_1_6bytes), PrimGuard1(__e, symtlstr, PrimTailString, V4142))


return GCons(__e, tmp1859, tmp1861)


}


panic("unreachable")
})

tmp1864 := Call2(__e, ns2_1set, symshen_4string_1_6bytes, tmp1857)


_ = tmp1864

tmp1865 := MakeNative1(func(__e *ControlFlow, V4143 Obj) Obj {
if True == GLt(__e, V4143, MakeInteger(0)) {
return GValue(__e, symshen_4_dmaxinferences_d)
} else {
tmp1867 := PrimGuard1(__e, syminteger_2, PrimIsInteger, V4143)

if True == tmp1867 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dmaxinferences_d, V4143)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("maxinferences expects an integer value\n"))
}


}


panic("unreachable")
})

tmp1870 := Call2(__e, ns2_1set, symmaxinferences, tmp1865)


_ = tmp1870

tmp1871 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dinfs_d)
panic("unreachable")
})

tmp1872 := Call2(__e, ns2_1set, syminferences, tmp1871)


_ = tmp1872

tmp1873 := MakeNative1(func(__e *ControlFlow, V4144 Obj) Obj {
return V4144
panic("unreachable")
})

tmp1874 := Call2(__e, ns2_1set, symprotect, tmp1873)


_ = tmp1874

tmp1875 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dsterror_d)
panic("unreachable")
})

tmp1876 := Call2(__e, ns2_1set, symsterror, tmp1875)


_ = tmp1876

tmp1877 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dstoutput_d)
panic("unreachable")
})

tmp1878 := Call2(__e, ns2_1set, symstoutput, tmp1877)


_ = tmp1878

tmp1879 := MakeNative1(func(__e *ControlFlow, V4145 Obj) Obj {
tmp1880 := PrimGuard1(__e, symintern, PrimIntern, V4145)

let__1199 := tmp1880
_ = let__1199

tmp1884 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, let__1199)

if True == tmp1884 {
return let__1199
} else {
tmp1881 := Call3(__e, SymFn(symshen_4app), V4145, MakeString(" to a symbol"), symshen_4s)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("cannot intern "), tmp1881))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("cannot intern "), tmp1881)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


panic("unreachable")
})

tmp1885 := Call2(__e, ns2_1set, symstring_1_6symbol, tmp1879)


_ = tmp1885

tmp1886 := MakeNative1(func(__e *ControlFlow, V4149 Obj) Obj {
tmp1890 := GEq(__e, sym_7, V4149)

if True == tmp1890 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_doptimise_d, True)
} else {
tmp1888 := GEq(__e, sym_1, V4149)

if True == tmp1888 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_doptimise_d, False)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("optimise expects a + or a -.\n"))
}


}


panic("unreachable")
})

tmp1891 := Call2(__e, ns2_1set, symoptimise, tmp1886)


_ = tmp1891

tmp1892 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dos_d)
panic("unreachable")
})

tmp1893 := Call2(__e, ns2_1set, symos, tmp1892)


_ = tmp1893

tmp1894 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dlanguage_d)
panic("unreachable")
})

tmp1895 := Call2(__e, ns2_1set, symlanguage, tmp1894)


_ = tmp1895

tmp1896 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dversion_d)
panic("unreachable")
})

tmp1897 := Call2(__e, ns2_1set, symversion, tmp1896)


_ = tmp1897

tmp1898 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dport_d)
panic("unreachable")
})

tmp1899 := Call2(__e, ns2_1set, symport, tmp1898)


_ = tmp1899

tmp1900 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dporters_d)
panic("unreachable")
})

tmp1901 := Call2(__e, ns2_1set, symporters, tmp1900)


_ = tmp1901

tmp1902 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dimplementation_d)
panic("unreachable")
})

tmp1903 := Call2(__e, ns2_1set, symimplementation, tmp1902)


_ = tmp1903

tmp1904 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_drelease_d)
panic("unreachable")
})

tmp1905 := Call2(__e, ns2_1set, symrelease, tmp1904)


_ = tmp1905

tmp1906 := MakeNative1(func(__e *ControlFlow, V4150 Obj) Obj {
tmp1911 := GEq(__e, symnull, V4150)

if True == tmp1911 {
return True
} else {
tmp1907 := MakeNative0(func(__e *ControlFlow) Obj {
tmp1908 := Call1(__e, SymFn(symexternal), V4150)


_ = tmp1908

return True


panic("unreachable")
})

tmp1909 := MakeNative1(func(__e *ControlFlow, Z4151 Obj) Obj {
return False
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp1907, tmp1909)


}


panic("unreachable")
})

tmp1912 := Call2(__e, ns2_1set, sympackage_2, tmp1906)


_ = tmp1912

tmp1913 := MakeNative0(func(__e *ControlFlow) Obj {
return symshen_4fail_b
panic("unreachable")
})

tmp1914 := Call2(__e, ns2_1set, symfail, tmp1913)


_ = tmp1914

tmp1915 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_duserdefs_d)
panic("unreachable")
})

tmp1916 := Call2(__e, ns2_1set, symuserdefs, tmp1915)


_ = tmp1916

tmp1917 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_doptimise_d)
panic("unreachable")
})

tmp1918 := Call2(__e, ns2_1set, symoptimise_2, tmp1917)


_ = tmp1918

tmp1919 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dhush_d)
panic("unreachable")
})

tmp1920 := Call2(__e, ns2_1set, symhush_2, tmp1919)


_ = tmp1920

tmp1921 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dshen_1type_1theory_1enabled_2_d)
panic("unreachable")
})

tmp1922 := Call2(__e, ns2_1set, symsystem_1S_2, tmp1921)


_ = tmp1922

tmp1923 := MakeNative1(func(__e *ControlFlow, V4154 Obj) Obj {
tmp1927 := GEq(__e, sym_7, V4154)

if True == tmp1927 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dshen_1type_1theory_1enabled_2_d, True)
} else {
tmp1925 := GEq(__e, sym_1, V4154)

if True == tmp1925 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dshen_1type_1theory_1enabled_2_d, False)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("enable-type-theory expects a + or a -\n"))
}


}


panic("unreachable")
})

tmp1928 := Call2(__e, ns2_1set, symenable_1type_1theory, tmp1923)


_ = tmp1928

tmp1929 := MakeNative1(func(__e *ControlFlow, V4157 Obj) Obj {
tmp1933 := GEq(__e, sym_7, V4157)

if True == tmp1933 {
return PrimGuard2(__e, symset, PrimSet, sym_dhush_d, True)
} else {
tmp1931 := GEq(__e, sym_1, V4157)

if True == tmp1931 {
return PrimGuard2(__e, symset, PrimSet, sym_dhush_d, False)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("hush expects a + or a -\n"))
}


}


panic("unreachable")
})

tmp1934 := Call2(__e, ns2_1set, symhush, tmp1929)


_ = tmp1934

tmp1935 := MakeNative1(func(__e *ControlFlow, V4160 Obj) Obj {
tmp1939 := GEq(__e, sym_7, V4160)

if True == tmp1939 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dtc_d, True)
} else {
tmp1937 := GEq(__e, sym_1, V4160)

if True == tmp1937 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dtc_d, False)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("tc expects a + or -"))
}


}


panic("unreachable")
})

tmp1940 := Call2(__e, ns2_1set, symtc, tmp1935)


_ = tmp1940

tmp1941 := MakeNative1(func(__e *ControlFlow, V4161 Obj) Obj {
tmp1942 := GValue(__e, symshen_4_dsigf_d)

tmp1943 := Call2(__e, SymFn(symshen_4unassoc), V4161, tmp1942)


tmp1944 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp1943)

_ = tmp1944

return V4161


panic("unreachable")
})

tmp1945 := Call2(__e, ns2_1set, symdestroy, tmp1941)


_ = tmp1945

tmp1946 := MakeNative2(func(__e *ControlFlow, V4171, V4172 Obj) Obj {
tmp1962 := GEq(__e, Nil, V4172)

if True == tmp1962 {
return Nil
} else {
tmp1960 := GConsP(__e, V4172)

var ifres1953 Obj

if True == tmp1960 {
tmp1958 := GHd(__e, V4172)

tmp1959 := GConsP(__e, tmp1958)

var ifres1954 Obj

if True == tmp1959 {
tmp1955 := GHd(__e, V4172)

tmp1956 := GHd(__e, tmp1955)

tmp1957 := GEq(__e, V4171, tmp1956)

ifres1954 = tmp1957


} else {
ifres1954 = False


}

ifres1953 = ifres1954


} else {
ifres1953 = False


}

if True == ifres1953 {
return GTl(__e, V4172)
} else {
tmp1951 := GConsP(__e, V4172)

if True == tmp1951 {
tmp1947 := GHd(__e, V4172)

tmp1948 := GTl(__e, V4172)

tmp1949 := Call2(__e, SymFn(symshen_4unassoc), V4171, tmp1948)


return GCons(__e, tmp1947, tmp1949)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.unassoc"))
}


}


}


panic("unreachable")
})

tmp1963 := Call2(__e, ns2_1set, symshen_4unassoc, tmp1946)


_ = tmp1963

tmp1964 := MakeNative1(func(__e *ControlFlow, V4173 Obj) Obj {
tmp1968 := Call1(__e, SymFn(sympackage_2), V4173)


if True == tmp1968 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dpackage_d, V4173)
} else {
tmp1965 := Call3(__e, SymFn(symshen_4app), V4173, MakeString(" does not exist\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("package "), tmp1965))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("package "), tmp1965)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


panic("unreachable")
})

tmp1969 := Call2(__e, ns2_1set, symin_1package, tmp1964)


_ = tmp1969

tmp1970 := MakeNative2(func(__e *ControlFlow, V4174, V4175 Obj) Obj {
tmp1971 := PrimGuard2(__e, symopen, PrimOpenStream, V4174, symout)

let__1200 := tmp1971
_ = let__1200

tmp1974 := PrimGuard1(__e, symstring_2, PrimIsString, V4175)

var ifres1972 Obj

if True == tmp1974 {
ifres1972 = V4175


} else {
tmp1973 := Call3(__e, SymFn(symshen_4app), V4175, MakeString(""), symshen_4s)


ifres1972 = tmp1973


}

let__1201 := ifres1972
_ = let__1201

tmp1975 := Call2(__e, SymFn(sympr), let__1201, let__1200)


let__1202 := tmp1975
_ = let__1202

tmp1976 := PrimGuard1(__e, symclose, PrimCloseStream, let__1200)

let__1203 := tmp1976
_ = let__1203

return V4175


panic("unreachable")
})

tmp1977 := Call2(__e, ns2_1set, symwrite_1to_1file, tmp1970)


_ = tmp1977

tmp1978 := MakeNative0(func(__e *ControlFlow) Obj {
tmp1979 := Call1(__e, SymFn(symgensym), symshen_4t)


return TailCall1(__e, SymFn(symshen_4freshterm), tmp1979)


panic("unreachable")
})

tmp1980 := Call2(__e, ns2_1set, symfresh, tmp1978)


_ = tmp1980

tmp1981 := MakeNative2(func(__e *ControlFlow, V4180, V4181 Obj) Obj {
tmp1982 := GValue(__e, sym_dproperty_1vector_d)

tmp1983 := Call4(__e, SymFn(symput), V4180, symarity, V4181, tmp1982)


let__1204 := tmp1983
_ = let__1204

tmp1984 := Call1(__e, SymFn(symshen_4lambda_1entry), V4180)


let__1205 := tmp1984
_ = let__1205

tmp1989 := Call1(__e, SymFn(symempty_2), let__1205)


var ifres1985 Obj

if True == tmp1989 {
tmp1986 := GValue(__e, sym_dproperty_1vector_d)

tmp1987 := Call3(__e, SymFn(symunput), V4180, symshen_4lambda_1form, tmp1986)


ifres1985 = tmp1987


} else {
tmp1988 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), let__1205)


ifres1985 = tmp1988


}

let__1206 := ifres1985
_ = let__1206

return V4180


panic("unreachable")
})

tmp1990 := Call2(__e, ns2_1set, symupdate_1lambda_1table, tmp1981)


_ = tmp1990

tmp1991 := MakeNative1(func(__e *ControlFlow, V4187 Obj) Obj {
return False
panic("unreachable")
})

tmp1992 := Call2(__e, ns2_1set, symshen_4precompiled_1function_2, tmp1991)


_ = tmp1992

tmp1993 := MakeNative4(func(__e *ControlFlow, V4188, V4189, V4190, V4191 Obj) Obj {
if True == V4191 {
tmp1994 := GValue(__e, sym_dproperty_1vector_d)

tmp1995 := Call4(__e, SymFn(symput), V4188, symarity, V4189, tmp1994)


let__1207 := tmp1995
_ = let__1207

tmp1996 := Call1(__e, SymFn(symthaw), V4190)


let__1208 := tmp1996
_ = let__1208

tmp2002 := GEq(__e, V4189, MakeInteger(0))

var ifres1997 Obj

if True == tmp2002 {
tmp1998 := GValue(__e, sym_dproperty_1vector_d)

tmp1999 := Call3(__e, SymFn(symunput), V4188, symshen_4lambda_1form, tmp1998)


ifres1997 = tmp1999


} else {
tmp2000 := GCons(__e, V4188, let__1208)

tmp2001 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp2000)


ifres1997 = tmp2001


}

let__1209 := ifres1997
_ = let__1209

return V4188


} else {
return TailCall2(__e, SymFn(symupdate_1lambda_1table), V4188, V4189)
}
panic("unreachable")
})

tmp2004 := Call2(__e, ns2_1set, symshen_4update_1lambda_1table_cprecompiled, tmp1993)


_ = tmp2004

tmp2005 := MakeNative2(func(__e *ControlFlow, V4197, V4198 Obj) Obj {
tmp2029 := GEq(__e, MakeInteger(0), V4198)

if True == tmp2029 {
tmp2006 := GValue(__e, symshen_4_dspecial_d)

tmp2007 := Call2(__e, SymFn(symremove), V4197, tmp2006)


tmp2008 := PrimGuard2(__e, symset, PrimSet, symshen_4_dspecial_d, tmp2007)

_ = tmp2008

tmp2009 := GValue(__e, symshen_4_dextraspecial_d)

tmp2010 := Call2(__e, SymFn(symremove), V4197, tmp2009)


tmp2011 := PrimGuard2(__e, symset, PrimSet, symshen_4_dextraspecial_d, tmp2010)

_ = tmp2011

return V4197


} else {
tmp2027 := GEq(__e, MakeInteger(1), V4198)

if True == tmp2027 {
tmp2012 := GValue(__e, symshen_4_dspecial_d)

tmp2013 := Call2(__e, SymFn(symadjoin), V4197, tmp2012)


tmp2014 := PrimGuard2(__e, symset, PrimSet, symshen_4_dspecial_d, tmp2013)

_ = tmp2014

tmp2015 := GValue(__e, symshen_4_dextraspecial_d)

tmp2016 := Call2(__e, SymFn(symremove), V4197, tmp2015)


tmp2017 := PrimGuard2(__e, symset, PrimSet, symshen_4_dextraspecial_d, tmp2016)

_ = tmp2017

return V4197


} else {
tmp2025 := GEq(__e, MakeInteger(2), V4198)

if True == tmp2025 {
tmp2018 := GValue(__e, symshen_4_dspecial_d)

tmp2019 := Call2(__e, SymFn(symremove), V4197, tmp2018)


tmp2020 := PrimGuard2(__e, symset, PrimSet, symshen_4_dspecial_d, tmp2019)

_ = tmp2020

tmp2021 := GValue(__e, symshen_4_dextraspecial_d)

tmp2022 := Call2(__e, SymFn(symadjoin), V4197, tmp2021)


tmp2023 := PrimGuard2(__e, symset, PrimSet, symshen_4_dextraspecial_d, tmp2022)

_ = tmp2023

return V4197


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("specialise requires values of 0, 1 or 2\n"))
}


}


}


panic("unreachable")
})

tmp2030 := Call2(__e, ns2_1set, symspecialise, tmp2005)


_ = tmp2030

tmp2031 := MakeNative1(func(__e *ControlFlow, V4199 Obj) Obj {
tmp2032 := GValue(__e, sym_dabsolute_d)

tmp2033 := GCons(__e, V4199, tmp2032)

return PrimGuard2(__e, symset, PrimSet, sym_dabsolute_d, tmp2033)


panic("unreachable")
})

tmp2034 := Call2(__e, ns2_1set, symabsolute, tmp2031)


_ = tmp2034

tmp2035 := MakeNative1(func(__e *ControlFlow, V4200 Obj) Obj {
tmp2036 := GValue(__e, sym_dabsolute_d)

tmp2037 := Call2(__e, SymFn(symremove), V4200, tmp2036)


return PrimGuard2(__e, symset, PrimSet, sym_dabsolute_d, tmp2037)


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symunabsolute, tmp2035)




panic("unreachable")
})

