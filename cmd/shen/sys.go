package main

import . "github.com/pyrex41/shen-go/kl"

var SysMain = MakeNative(func(__e *ControlFlow) {
tmp1158 := MakeNative(func(__e *ControlFlow) {
V3430 := __e.Get(1)
_ = V3430
__e.TailApply(V3430)
return
}, 1)

tmp1159 := Call(__e, ns2_1set, symthaw, tmp1158)


_ = tmp1159

tmp1160 := MakeNative(func(__e *ControlFlow) {
V3431 := __e.Get(1)
_ = V3431
tmp1161 := Call(__e, PrimFunc(symmacroexpand), V3431)


tmp1162 := Call(__e, PrimFunc(symshen_4find_1types), V3431)


tmp1163 := Call(__e, PrimFunc(symshen_4process_1applications), tmp1161, tmp1162)


tmp1164 := Call(__e, PrimFunc(symshen_4shen_1_6kl), tmp1163)


__e.TailApply(PrimFunc(symeval_1kl), tmp1164)
return


}, 1)

tmp1165 := Call(__e, ns2_1set, symeval, tmp1160)


_ = tmp1165

tmp1166 := MakeNative(func(__e *ControlFlow) {
V3432 := __e.Get(1)
_ = V3432
tmp1173 := PrimEqual(symnull, V3432)

if True == tmp1173 {
__e.Return(Nil)
return
} else {
tmp1167 := MakeNative(func(__e *ControlFlow) {
tmp1168 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V3432, symshen_4external_1symbols, tmp1168)
return


}, 0)

tmp1169 := MakeNative(func(__e *ControlFlow) {
Z3433 := __e.Get(1)
_ = Z3433
tmp1170 := Call(__e, PrimFunc(symshen_4app), V3432, MakeString(" does not exist.\n;"), symshen_4a)


tmp1171 := PrimStringConcat(MakeString("package "), tmp1170)

__e.Return(PrimSimpleError(tmp1171))
return


}, 1)

__e.TailApply(try_1catch, tmp1167, tmp1169)
return


}


}, 1)

tmp1174 := Call(__e, ns2_1set, symexternal, tmp1166)


_ = tmp1174

tmp1175 := MakeNative(func(__e *ControlFlow) {
V3434 := __e.Get(1)
_ = V3434
tmp1182 := PrimEqual(symnull, V3434)

if True == tmp1182 {
__e.Return(Nil)
return
} else {
tmp1176 := MakeNative(func(__e *ControlFlow) {
tmp1177 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V3434, symshen_4internal_1symbols, tmp1177)
return


}, 0)

tmp1178 := MakeNative(func(__e *ControlFlow) {
Z3435 := __e.Get(1)
_ = Z3435
tmp1179 := Call(__e, PrimFunc(symshen_4app), V3434, MakeString(" does not exist.\n;"), symshen_4a)


tmp1180 := PrimStringConcat(MakeString("package "), tmp1179)

__e.Return(PrimSimpleError(tmp1180))
return


}, 1)

__e.TailApply(try_1catch, tmp1176, tmp1178)
return


}


}, 1)

tmp1183 := Call(__e, ns2_1set, syminternal, tmp1175)


_ = tmp1183

tmp1184 := MakeNative(func(__e *ControlFlow) {
V3436 := __e.Get(1)
_ = V3436
V3437 := __e.Get(2)
_ = V3437
tmp1186 := Call(__e, V3436, V3437)


if True == tmp1186 {
__e.TailApply(PrimFunc(symfail))
return
} else {
__e.Return(V3437)
return
}


}, 2)

tmp1187 := Call(__e, ns2_1set, symfail_1if, tmp1184)


_ = tmp1187

tmp1188 := MakeNative(func(__e *ControlFlow) {
V3438 := __e.Get(1)
_ = V3438
V3439 := __e.Get(2)
_ = V3439
__e.Return(PrimStringConcat(V3438, V3439))
return
}, 2)

tmp1189 := Call(__e, ns2_1set, sym_8s, tmp1188)


_ = tmp1189

tmp1190 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dtc_d))
return
}, 0)

tmp1191 := Call(__e, ns2_1set, symtc_2, tmp1190)


_ = tmp1191

tmp1192 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_doccurs_d))
return
}, 0)

tmp1193 := Call(__e, ns2_1set, symoccurs_2, tmp1192)


_ = tmp1193

tmp1194 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dfactorise_2_d))
return
}, 0)

tmp1195 := Call(__e, ns2_1set, symfactorise_2, tmp1194)


_ = tmp1195

tmp1196 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dtracking_d))
return
}, 0)

tmp1197 := Call(__e, ns2_1set, symtracked, tmp1196)


_ = tmp1197

tmp1198 := MakeNative(func(__e *ControlFlow) {
V3440 := __e.Get(1)
_ = V3440
tmp1199 := MakeNative(func(__e *ControlFlow) {
tmp1200 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V3440, symshen_4source, tmp1200)
return


}, 0)

tmp1201 := MakeNative(func(__e *ControlFlow) {
Z3441 := __e.Get(1)
_ = Z3441
tmp1202 := Call(__e, PrimFunc(symshen_4app), V3440, MakeString(" not found.\n"), symshen_4a)


__e.Return(PrimSimpleError(tmp1202))
return


}, 1)

__e.TailApply(try_1catch, tmp1199, tmp1201)
return


}, 1)

tmp1203 := Call(__e, ns2_1set, symps, tmp1198)


_ = tmp1203

tmp1204 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dstinput_d))
return
}, 0)

tmp1205 := Call(__e, ns2_1set, symstinput, tmp1204)


_ = tmp1205

tmp1206 := MakeNative(func(__e *ControlFlow) {
V3442 := __e.Get(1)
_ = V3442
tmp1207 := MakeNative(func(__e *ControlFlow) {
W3443 := __e.Get(1)
_ = W3443
tmp1208 := MakeNative(func(__e *ControlFlow) {
W3444 := __e.Get(1)
_ = W3444
tmp1209 := MakeNative(func(__e *ControlFlow) {
W3445 := __e.Get(1)
_ = W3445
__e.Return(W3445)
return
}, 1)

tmp1213 := PrimEqual(V3442, MakeNumber(0))

var ifres1210 Obj

if True == tmp1213 {
ifres1210 = W3444


} else {
tmp1211 := Call(__e, PrimFunc(symfail))


tmp1212 := Call(__e, PrimFunc(symshen_4fillvector), W3444, MakeNumber(1), V3442, tmp1211)


ifres1210 = tmp1212


}

__e.TailApply(tmp1209, ifres1210)
return


}, 1)

tmp1214 := PrimVectorSet(W3443, MakeNumber(0), V3442)

__e.TailApply(tmp1208, tmp1214)
return


}, 1)

tmp1215 := PrimNumberAdd(V3442, MakeNumber(1))

tmp1216 := PrimAbsvector(tmp1215)

__e.TailApply(tmp1207, tmp1216)
return


}, 1)

tmp1217 := Call(__e, ns2_1set, symvector, tmp1206)


_ = tmp1217

tmp1218 := MakeNative(func(__e *ControlFlow) {
V3447 := __e.Get(1)
_ = V3447
V3448 := __e.Get(2)
_ = V3448
V3449 := __e.Get(3)
_ = V3449
V3450 := __e.Get(4)
_ = V3450
tmp1222 := PrimEqual(V3448, V3449)

if True == tmp1222 {
__e.Return(PrimVectorSet(V3447, V3449, V3450))
return
} else {
tmp1219 := PrimVectorSet(V3447, V3448, V3450)

tmp1220 := PrimNumberAdd(MakeNumber(1), V3448)

__e.TailApply(PrimFunc(symshen_4fillvector), tmp1219, tmp1220, V3449, V3450)
return


}


}, 4)

tmp1223 := Call(__e, ns2_1set, symshen_4fillvector, tmp1218)


_ = tmp1223

tmp1224 := MakeNative(func(__e *ControlFlow) {
V3451 := __e.Get(1)
_ = V3451
tmp1231 := PrimIsVector(V3451)

if True == tmp1231 {
tmp1226 := MakeNative(func(__e *ControlFlow) {
tmp1227 := PrimVectorGet(V3451, MakeNumber(0))

__e.Return(PrimGreatEqual(tmp1227, MakeNumber(0)))
return


}, 0)

tmp1228 := MakeNative(func(__e *ControlFlow) {
Z3452 := __e.Get(1)
_ = Z3452
__e.Return(False)
return
}, 1)

tmp1229 := Call(__e, try_1catch, tmp1226, tmp1228)


if True == tmp1229 {
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

tmp1232 := Call(__e, ns2_1set, symvector_2, tmp1224)


_ = tmp1232

tmp1233 := MakeNative(func(__e *ControlFlow) {
V3453 := __e.Get(1)
_ = V3453
V3454 := __e.Get(2)
_ = V3454
V3455 := __e.Get(3)
_ = V3455
tmp1235 := PrimEqual(V3454, MakeNumber(0))

if True == tmp1235 {
__e.Return(PrimSimpleError(MakeString("cannot access 0th element of a vector\n")))
return
} else {
__e.Return(PrimVectorSet(V3453, V3454, V3455))
return
}


}, 3)

tmp1236 := Call(__e, ns2_1set, symvector_1_6, tmp1233)


_ = tmp1236

tmp1237 := MakeNative(func(__e *ControlFlow) {
V3456 := __e.Get(1)
_ = V3456
V3457 := __e.Get(2)
_ = V3457
tmp1244 := PrimEqual(V3457, MakeNumber(0))

if True == tmp1244 {
__e.Return(PrimSimpleError(MakeString("cannot access 0th element of a vector\n")))
return
} else {
tmp1238 := MakeNative(func(__e *ControlFlow) {
W3458 := __e.Get(1)
_ = W3458
tmp1240 := Call(__e, PrimFunc(symfail))


tmp1241 := PrimEqual(W3458, tmp1240)

if True == tmp1241 {
__e.Return(PrimSimpleError(MakeString("vector element not found\n")))
return
} else {
__e.Return(W3458)
return
}


}, 1)

tmp1242 := PrimVectorGet(V3456, V3457)

__e.TailApply(tmp1238, tmp1242)
return


}


}, 2)

tmp1245 := Call(__e, ns2_1set, sym_5_1vector, tmp1237)


_ = tmp1245

tmp1246 := MakeNative(func(__e *ControlFlow) {
V3459 := __e.Get(1)
_ = V3459
tmp1250 := PrimIsInteger(V3459)

if True == tmp1250 {
tmp1248 := PrimGreatEqual(V3459, MakeNumber(0))

if True == tmp1248 {
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

tmp1251 := Call(__e, ns2_1set, symshen_4posint_2, tmp1246)


_ = tmp1251

tmp1252 := MakeNative(func(__e *ControlFlow) {
V3460 := __e.Get(1)
_ = V3460
__e.Return(PrimVectorGet(V3460, MakeNumber(0)))
return
}, 1)

tmp1253 := Call(__e, ns2_1set, symlimit, tmp1252)


_ = tmp1253

tmp1254 := MakeNative(func(__e *ControlFlow) {
V3461 := __e.Get(1)
_ = V3461
tmp1285 := Call(__e, PrimFunc(symboolean_2), V3461)


var ifres1270 Obj

if True == tmp1285 {
ifres1270 = True


} else {
tmp1284 := PrimIsNumber(V3461)

var ifres1272 Obj

if True == tmp1284 {
ifres1272 = True


} else {
tmp1283 := PrimIsString(V3461)

var ifres1274 Obj

if True == tmp1283 {
ifres1274 = True


} else {
tmp1282 := PrimIsPair(V3461)

var ifres1276 Obj

if True == tmp1282 {
ifres1276 = True


} else {
tmp1281 := Call(__e, PrimFunc(symempty_2), V3461)


var ifres1278 Obj

if True == tmp1281 {
ifres1278 = True


} else {
tmp1280 := Call(__e, PrimFunc(symvector_2), V3461)


var ifres1279 Obj

if True == tmp1280 {
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

var ifres1275 Obj

if True == ifres1276 {
ifres1275 = True


} else {
ifres1275 = False


}

ifres1274 = ifres1275


}

var ifres1273 Obj

if True == ifres1274 {
ifres1273 = True


} else {
ifres1273 = False


}

ifres1272 = ifres1273


}

var ifres1271 Obj

if True == ifres1272 {
ifres1271 = True


} else {
ifres1271 = False


}

ifres1270 = ifres1271


}

if True == ifres1270 {
__e.Return(False)
return
} else {
tmp1260 := PrimIntern(MakeString(":"))

tmp1261 := PrimIntern(MakeString(";"))

tmp1262 := PrimIntern(MakeString(","))

tmp1263 := PrimCons(tmp1262, Nil)

tmp1264 := PrimCons(tmp1261, tmp1263)

tmp1265 := PrimCons(tmp1260, tmp1264)

tmp1266 := PrimCons(sym_j, tmp1265)

tmp1267 := PrimCons(sym_i, tmp1266)

tmp1268 := Call(__e, PrimFunc(symelement_2), V3461, tmp1267)


if True == tmp1268 {
__e.Return(True)
return
} else {
tmp1255 := MakeNative(func(__e *ControlFlow) {
tmp1256 := MakeNative(func(__e *ControlFlow) {
W3462 := __e.Get(1)
_ = W3462
__e.TailApply(PrimFunc(symshen_4analyse_1symbol_2), W3462)
return
}, 1)

tmp1257 := PrimStr(V3461)

__e.TailApply(tmp1256, tmp1257)
return


}, 0)

tmp1258 := MakeNative(func(__e *ControlFlow) {
Z3463 := __e.Get(1)
_ = Z3463
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1255, tmp1258)
return


}


}


}, 1)

tmp1286 := Call(__e, ns2_1set, symsymbol_2, tmp1254)


_ = tmp1286

tmp1287 := MakeNative(func(__e *ControlFlow) {
V3466 := __e.Get(1)
_ = V3466
tmp1296 := Call(__e, PrimFunc(symshen_4_7string_2), V3466)


if True == tmp1296 {
tmp1292 := Call(__e, PrimFunc(symhdstr), V3466)


tmp1293 := PrimStringToNumber(tmp1292)

tmp1294 := Call(__e, PrimFunc(symshen_4alpha_2), tmp1293)


if True == tmp1294 {
tmp1289 := PrimTailString(V3466)

tmp1290 := Call(__e, PrimFunc(symshen_4alphanums_2), tmp1289)


if True == tmp1290 {
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
__e.Return(PrimSimpleError(MakeString("implementation error in shen.analyse-symbol?")))
return
}


}, 1)

tmp1297 := Call(__e, ns2_1set, symshen_4analyse_1symbol_2, tmp1287)


_ = tmp1297

tmp1298 := MakeNative(func(__e *ControlFlow) {
V3469 := __e.Get(1)
_ = V3469
tmp1313 := PrimEqual(MakeString(""), V3469)

if True == tmp1313 {
__e.Return(True)
return
} else {
tmp1311 := Call(__e, PrimFunc(symshen_4_7string_2), V3469)


if True == tmp1311 {
tmp1299 := MakeNative(func(__e *ControlFlow) {
W3470 := __e.Get(1)
_ = W3470
tmp1307 := Call(__e, PrimFunc(symshen_4alpha_2), W3470)


var ifres1304 Obj

if True == tmp1307 {
ifres1304 = True


} else {
tmp1306 := Call(__e, PrimFunc(symshen_4digit_2), W3470)


var ifres1305 Obj

if True == tmp1306 {
ifres1305 = True


} else {
ifres1305 = False


}

ifres1304 = ifres1305


}

if True == ifres1304 {
tmp1301 := PrimTailString(V3469)

tmp1302 := Call(__e, PrimFunc(symshen_4alphanums_2), tmp1301)


if True == tmp1302 {
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

tmp1308 := Call(__e, PrimFunc(symhdstr), V3469)


tmp1309 := PrimStringToNumber(tmp1308)

__e.TailApply(tmp1299, tmp1309)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.alphanums?")))
return
}


}


}, 1)

tmp1314 := Call(__e, ns2_1set, symshen_4alphanums_2, tmp1298)


_ = tmp1314

tmp1315 := MakeNative(func(__e *ControlFlow) {
V3471 := __e.Get(1)
_ = V3471
tmp1327 := Call(__e, PrimFunc(symboolean_2), V3471)


var ifres1321 Obj

if True == tmp1327 {
ifres1321 = True


} else {
tmp1326 := PrimIsNumber(V3471)

var ifres1323 Obj

if True == tmp1326 {
ifres1323 = True


} else {
tmp1325 := PrimIsString(V3471)

var ifres1324 Obj

if True == tmp1325 {
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

if True == ifres1321 {
__e.Return(False)
return
} else {
tmp1316 := MakeNative(func(__e *ControlFlow) {
tmp1317 := MakeNative(func(__e *ControlFlow) {
W3472 := __e.Get(1)
_ = W3472
__e.TailApply(PrimFunc(symshen_4analyse_1variable_2), W3472)
return
}, 1)

tmp1318 := PrimStr(V3471)

__e.TailApply(tmp1317, tmp1318)
return


}, 0)

tmp1319 := MakeNative(func(__e *ControlFlow) {
Z3473 := __e.Get(1)
_ = Z3473
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1316, tmp1319)
return


}


}, 1)

tmp1328 := Call(__e, ns2_1set, symvariable_2, tmp1315)


_ = tmp1328

tmp1329 := MakeNative(func(__e *ControlFlow) {
V3476 := __e.Get(1)
_ = V3476
tmp1338 := Call(__e, PrimFunc(symshen_4_7string_2), V3476)


if True == tmp1338 {
tmp1334 := Call(__e, PrimFunc(symhdstr), V3476)


tmp1335 := PrimStringToNumber(tmp1334)

tmp1336 := Call(__e, PrimFunc(symshen_4uppercase_2), tmp1335)


if True == tmp1336 {
tmp1331 := PrimTailString(V3476)

tmp1332 := Call(__e, PrimFunc(symshen_4alphanums_2), tmp1331)


if True == tmp1332 {
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
__e.Return(PrimSimpleError(MakeString("implementation error in shen.analyse-variable?")))
return
}


}, 1)

tmp1339 := Call(__e, ns2_1set, symshen_4analyse_1variable_2, tmp1329)


_ = tmp1339

tmp1340 := MakeNative(func(__e *ControlFlow) {
V3477 := __e.Get(1)
_ = V3477
tmp1341 := PrimValue(symshen_4_dgensym_d)

tmp1342 := PrimNumberAdd(MakeNumber(1), tmp1341)

tmp1343 := PrimSet(symshen_4_dgensym_d, tmp1342)

__e.TailApply(PrimFunc(symconcat), V3477, tmp1343)
return


}, 1)

tmp1344 := Call(__e, ns2_1set, symgensym, tmp1340)


_ = tmp1344

tmp1345 := MakeNative(func(__e *ControlFlow) {
V3478 := __e.Get(1)
_ = V3478
V3479 := __e.Get(2)
_ = V3479
tmp1346 := PrimStr(V3478)

tmp1347 := PrimStr(V3479)

tmp1348 := PrimStringConcat(tmp1346, tmp1347)

__e.Return(PrimIntern(tmp1348))
return


}, 2)

tmp1349 := Call(__e, ns2_1set, symconcat, tmp1345)


_ = tmp1349

tmp1350 := MakeNative(func(__e *ControlFlow) {
V3480 := __e.Get(1)
_ = V3480
V3481 := __e.Get(2)
_ = V3481
tmp1351 := MakeNative(func(__e *ControlFlow) {
W3482 := __e.Get(1)
_ = W3482
tmp1352 := MakeNative(func(__e *ControlFlow) {
W3483 := __e.Get(1)
_ = W3483
tmp1353 := MakeNative(func(__e *ControlFlow) {
W3484 := __e.Get(1)
_ = W3484
tmp1354 := MakeNative(func(__e *ControlFlow) {
W3485 := __e.Get(1)
_ = W3485
__e.Return(W3482)
return
}, 1)

tmp1355 := PrimVectorSet(W3482, MakeNumber(2), V3481)

__e.TailApply(tmp1354, tmp1355)
return


}, 1)

tmp1356 := PrimVectorSet(W3482, MakeNumber(1), V3480)

__e.TailApply(tmp1353, tmp1356)
return


}, 1)

tmp1357 := PrimVectorSet(W3482, MakeNumber(0), symshen_4tuple)

__e.TailApply(tmp1352, tmp1357)
return


}, 1)

tmp1358 := PrimAbsvector(MakeNumber(3))

__e.TailApply(tmp1351, tmp1358)
return


}, 2)

tmp1359 := Call(__e, ns2_1set, sym_8p, tmp1350)


_ = tmp1359

tmp1360 := MakeNative(func(__e *ControlFlow) {
V3486 := __e.Get(1)
_ = V3486
__e.Return(PrimVectorGet(V3486, MakeNumber(1)))
return
}, 1)

tmp1361 := Call(__e, ns2_1set, symfst, tmp1360)


_ = tmp1361

tmp1362 := MakeNative(func(__e *ControlFlow) {
V3487 := __e.Get(1)
_ = V3487
__e.Return(PrimVectorGet(V3487, MakeNumber(2)))
return
}, 1)

tmp1363 := Call(__e, ns2_1set, symsnd, tmp1362)


_ = tmp1363

tmp1364 := MakeNative(func(__e *ControlFlow) {
V3488 := __e.Get(1)
_ = V3488
tmp1365 := MakeNative(func(__e *ControlFlow) {
tmp1370 := PrimIsVector(V3488)

if True == tmp1370 {
tmp1367 := PrimVectorGet(V3488, MakeNumber(0))

tmp1368 := PrimEqual(symshen_4tuple, tmp1367)

if True == tmp1368 {
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

tmp1371 := MakeNative(func(__e *ControlFlow) {
Z3489 := __e.Get(1)
_ = Z3489
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1365, tmp1371)
return


}, 1)

tmp1372 := Call(__e, ns2_1set, symtuple_2, tmp1364)


_ = tmp1372

tmp1373 := MakeNative(func(__e *ControlFlow) {
V3494 := __e.Get(1)
_ = V3494
V3495 := __e.Get(2)
_ = V3495
tmp1380 := PrimEqual(Nil, V3494)

if True == tmp1380 {
__e.Return(V3495)
return
} else {
tmp1378 := PrimIsPair(V3494)

if True == tmp1378 {
tmp1374 := PrimHead(V3494)

tmp1375 := PrimTail(V3494)

tmp1376 := Call(__e, PrimFunc(symappend), tmp1375, V3495)


__e.Return(PrimCons(tmp1374, tmp1376))
return


} else {
__e.Return(PrimSimpleError(MakeString("attempt to append a non-list")))
return
}


}


}, 2)

tmp1381 := Call(__e, ns2_1set, symappend, tmp1373)


_ = tmp1381

tmp1382 := MakeNative(func(__e *ControlFlow) {
V3496 := __e.Get(1)
_ = V3496
V3497 := __e.Get(2)
_ = V3497
tmp1383 := MakeNative(func(__e *ControlFlow) {
W3498 := __e.Get(1)
_ = W3498
tmp1384 := MakeNative(func(__e *ControlFlow) {
W3499 := __e.Get(1)
_ = W3499
tmp1385 := MakeNative(func(__e *ControlFlow) {
W3500 := __e.Get(1)
_ = W3500
tmp1387 := PrimEqual(W3498, MakeNumber(0))

if True == tmp1387 {
__e.Return(W3500)
return
} else {
__e.TailApply(PrimFunc(symshen_4_8v_1help), V3497, MakeNumber(1), W3498, W3500)
return
}


}, 1)

tmp1388 := Call(__e, PrimFunc(symvector_1_6), W3499, MakeNumber(1), V3496)


__e.TailApply(tmp1385, tmp1388)
return


}, 1)

tmp1389 := PrimNumberAdd(W3498, MakeNumber(1))

tmp1390 := Call(__e, PrimFunc(symvector), tmp1389)


__e.TailApply(tmp1384, tmp1390)
return


}, 1)

tmp1391 := Call(__e, PrimFunc(symlimit), V3497)


__e.TailApply(tmp1383, tmp1391)
return


}, 2)

tmp1392 := Call(__e, ns2_1set, sym_8v, tmp1382)


_ = tmp1392

tmp1393 := MakeNative(func(__e *ControlFlow) {
V3502 := __e.Get(1)
_ = V3502
V3503 := __e.Get(2)
_ = V3503
V3504 := __e.Get(3)
_ = V3504
V3505 := __e.Get(4)
_ = V3505
tmp1399 := PrimEqual(V3503, V3504)

if True == tmp1399 {
tmp1394 := PrimNumberAdd(V3504, MakeNumber(1))

__e.TailApply(PrimFunc(symshen_4copyfromvector), V3502, V3505, V3504, tmp1394)
return


} else {
tmp1395 := PrimNumberAdd(V3503, MakeNumber(1))

tmp1396 := PrimNumberAdd(V3503, MakeNumber(1))

tmp1397 := Call(__e, PrimFunc(symshen_4copyfromvector), V3502, V3505, V3503, tmp1396)


__e.TailApply(PrimFunc(symshen_4_8v_1help), V3502, tmp1395, V3504, tmp1397)
return


}


}, 4)

tmp1400 := Call(__e, ns2_1set, symshen_4_8v_1help, tmp1393)


_ = tmp1400

tmp1401 := MakeNative(func(__e *ControlFlow) {
V3506 := __e.Get(1)
_ = V3506
V3507 := __e.Get(2)
_ = V3507
V3508 := __e.Get(3)
_ = V3508
V3509 := __e.Get(4)
_ = V3509
tmp1402 := MakeNative(func(__e *ControlFlow) {
tmp1403 := Call(__e, PrimFunc(sym_5_1vector), V3506, V3508)


__e.TailApply(PrimFunc(symvector_1_6), V3507, V3509, tmp1403)
return


}, 0)

tmp1404 := MakeNative(func(__e *ControlFlow) {
Z3510 := __e.Get(1)
_ = Z3510
__e.Return(V3507)
return
}, 1)

__e.TailApply(try_1catch, tmp1402, tmp1404)
return


}, 4)

tmp1405 := Call(__e, ns2_1set, symshen_4copyfromvector, tmp1401)


_ = tmp1405

tmp1406 := MakeNative(func(__e *ControlFlow) {
V3511 := __e.Get(1)
_ = V3511
tmp1407 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(sym_5_1vector), V3511, MakeNumber(1))
return
}, 0)

tmp1408 := MakeNative(func(__e *ControlFlow) {
Z3512 := __e.Get(1)
_ = Z3512
__e.Return(PrimSimpleError(MakeString("hdv needs a non-empty vector as an argument\n")))
return
}, 1)

__e.TailApply(try_1catch, tmp1407, tmp1408)
return


}, 1)

tmp1409 := Call(__e, ns2_1set, symhdv, tmp1406)


_ = tmp1409

tmp1410 := MakeNative(func(__e *ControlFlow) {
V3513 := __e.Get(1)
_ = V3513
tmp1411 := MakeNative(func(__e *ControlFlow) {
W3514 := __e.Get(1)
_ = W3514
tmp1420 := PrimEqual(W3514, MakeNumber(0))

if True == tmp1420 {
__e.Return(PrimSimpleError(MakeString("cannot take the tail of the empty vector\n")))
return
} else {
tmp1418 := PrimEqual(W3514, MakeNumber(1))

if True == tmp1418 {
__e.TailApply(PrimFunc(symvector), MakeNumber(0))
return
} else {
tmp1412 := MakeNative(func(__e *ControlFlow) {
W3515 := __e.Get(1)
_ = W3515
tmp1413 := PrimNumberSubtract(W3514, MakeNumber(1))

tmp1414 := Call(__e, PrimFunc(symvector), tmp1413)


__e.TailApply(PrimFunc(symshen_4tlv_1help), V3513, MakeNumber(2), W3514, tmp1414)
return


}, 1)

tmp1415 := PrimNumberSubtract(W3514, MakeNumber(1))

tmp1416 := Call(__e, PrimFunc(symvector), tmp1415)


__e.TailApply(tmp1412, tmp1416)
return


}


}


}, 1)

tmp1421 := Call(__e, PrimFunc(symlimit), V3513)


__e.TailApply(tmp1411, tmp1421)
return


}, 1)

tmp1422 := Call(__e, ns2_1set, symtlv, tmp1410)


_ = tmp1422

tmp1423 := MakeNative(func(__e *ControlFlow) {
V3517 := __e.Get(1)
_ = V3517
V3518 := __e.Get(2)
_ = V3518
V3519 := __e.Get(3)
_ = V3519
V3520 := __e.Get(4)
_ = V3520
tmp1429 := PrimEqual(V3518, V3519)

if True == tmp1429 {
tmp1424 := PrimNumberSubtract(V3519, MakeNumber(1))

__e.TailApply(PrimFunc(symshen_4copyfromvector), V3517, V3520, V3519, tmp1424)
return


} else {
tmp1425 := PrimNumberAdd(V3518, MakeNumber(1))

tmp1426 := PrimNumberSubtract(V3518, MakeNumber(1))

tmp1427 := Call(__e, PrimFunc(symshen_4copyfromvector), V3517, V3520, V3518, tmp1426)


__e.TailApply(PrimFunc(symshen_4tlv_1help), V3517, tmp1425, V3519, tmp1427)
return


}


}, 4)

tmp1430 := Call(__e, ns2_1set, symshen_4tlv_1help, tmp1423)


_ = tmp1430

tmp1431 := MakeNative(func(__e *ControlFlow) {
V3532 := __e.Get(1)
_ = V3532
V3533 := __e.Get(2)
_ = V3533
tmp1447 := PrimEqual(Nil, V3533)

if True == tmp1447 {
__e.Return(Nil)
return
} else {
tmp1445 := PrimIsPair(V3533)

var ifres1436 Obj

if True == tmp1445 {
tmp1443 := PrimHead(V3533)

tmp1444 := PrimIsPair(tmp1443)

var ifres1438 Obj

if True == tmp1444 {
tmp1440 := PrimHead(V3533)

tmp1441 := PrimHead(tmp1440)

tmp1442 := PrimEqual(V3532, tmp1441)

var ifres1439 Obj

if True == tmp1442 {
ifres1439 = True


} else {
ifres1439 = False


}

ifres1438 = ifres1439


} else {
ifres1438 = False


}

var ifres1437 Obj

if True == ifres1438 {
ifres1437 = True


} else {
ifres1437 = False


}

ifres1436 = ifres1437


} else {
ifres1436 = False


}

if True == ifres1436 {
__e.Return(PrimHead(V3533))
return
} else {
tmp1434 := PrimIsPair(V3533)

if True == tmp1434 {
tmp1432 := PrimTail(V3533)

__e.TailApply(PrimFunc(symassoc), V3532, tmp1432)
return


} else {
__e.Return(PrimSimpleError(MakeString("attempt to search a non-list with assoc\n")))
return
}


}


}


}, 2)

tmp1448 := Call(__e, ns2_1set, symassoc, tmp1431)


_ = tmp1448

tmp1449 := MakeNative(func(__e *ControlFlow) {
V3536 := __e.Get(1)
_ = V3536
tmp1453 := PrimEqual(True, V3536)

if True == tmp1453 {
__e.Return(True)
return
} else {
tmp1451 := PrimEqual(False, V3536)

if True == tmp1451 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp1454 := Call(__e, ns2_1set, symboolean_2, tmp1449)


_ = tmp1454

tmp1455 := MakeNative(func(__e *ControlFlow) {
V3537 := __e.Get(1)
_ = V3537
tmp1460 := PrimEqual(MakeNumber(0), V3537)

if True == tmp1460 {
__e.Return(MakeNumber(0))
return
} else {
tmp1456 := Call(__e, PrimFunc(symstoutput))


tmp1457 := Call(__e, PrimFunc(sympr), MakeString("\n"), tmp1456)


_ = tmp1457

tmp1458 := PrimNumberSubtract(V3537, MakeNumber(1))

__e.TailApply(PrimFunc(symnl), tmp1458)
return


}


}, 1)

tmp1461 := Call(__e, ns2_1set, symnl, tmp1455)


_ = tmp1461

tmp1462 := MakeNative(func(__e *ControlFlow) {
V3544 := __e.Get(1)
_ = V3544
V3545 := __e.Get(2)
_ = V3545
tmp1473 := PrimEqual(Nil, V3544)

if True == tmp1473 {
__e.Return(Nil)
return
} else {
tmp1471 := PrimIsPair(V3544)

if True == tmp1471 {
tmp1468 := PrimHead(V3544)

tmp1469 := Call(__e, PrimFunc(symelement_2), tmp1468, V3545)


if True == tmp1469 {
tmp1463 := PrimTail(V3544)

__e.TailApply(PrimFunc(symdifference), tmp1463, V3545)
return


} else {
tmp1464 := PrimHead(V3544)

tmp1465 := PrimTail(V3544)

tmp1466 := Call(__e, PrimFunc(symdifference), tmp1465, V3545)


__e.Return(PrimCons(tmp1464, tmp1466))
return


}


} else {
__e.Return(PrimSimpleError(MakeString("attempt to find the difference with a non-list\n")))
return
}


}


}, 2)

tmp1474 := Call(__e, ns2_1set, symdifference, tmp1462)


_ = tmp1474

tmp1475 := MakeNative(func(__e *ControlFlow) {
V3546 := __e.Get(1)
_ = V3546
V3547 := __e.Get(2)
_ = V3547
__e.Return(V3547)
return
}, 2)

tmp1476 := Call(__e, ns2_1set, symdo, tmp1475)


_ = tmp1476

tmp1477 := MakeNative(func(__e *ControlFlow) {
V3559 := __e.Get(1)
_ = V3559
V3560 := __e.Get(2)
_ = V3560
tmp1488 := PrimEqual(Nil, V3560)

if True == tmp1488 {
__e.Return(False)
return
} else {
tmp1486 := PrimIsPair(V3560)

var ifres1482 Obj

if True == tmp1486 {
tmp1484 := PrimHead(V3560)

tmp1485 := PrimEqual(V3559, tmp1484)

var ifres1483 Obj

if True == tmp1485 {
ifres1483 = True


} else {
ifres1483 = False


}

ifres1482 = ifres1483


} else {
ifres1482 = False


}

if True == ifres1482 {
__e.Return(True)
return
} else {
tmp1480 := PrimIsPair(V3560)

if True == tmp1480 {
tmp1478 := PrimTail(V3560)

__e.TailApply(PrimFunc(symelement_2), V3559, tmp1478)
return


} else {
__e.Return(PrimSimpleError(MakeString("attempt to find an element in a non-list\n")))
return
}


}


}


}, 2)

tmp1489 := Call(__e, ns2_1set, symelement_2, tmp1477)


_ = tmp1489

tmp1490 := MakeNative(func(__e *ControlFlow) {
V3563 := __e.Get(1)
_ = V3563
tmp1492 := PrimEqual(Nil, V3563)

if True == tmp1492 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp1493 := Call(__e, ns2_1set, symempty_2, tmp1490)


_ = tmp1493

tmp1494 := MakeNative(func(__e *ControlFlow) {
V3564 := __e.Get(1)
_ = V3564
V3565 := __e.Get(2)
_ = V3565
tmp1495 := Call(__e, V3564, V3565)


__e.TailApply(PrimFunc(symshen_4fix_1help), V3564, V3565, tmp1495)
return


}, 2)

tmp1496 := Call(__e, ns2_1set, symfix, tmp1494)


_ = tmp1496

tmp1497 := MakeNative(func(__e *ControlFlow) {
V3571 := __e.Get(1)
_ = V3571
V3572 := __e.Get(2)
_ = V3572
V3573 := __e.Get(3)
_ = V3573
tmp1500 := PrimEqual(V3572, V3573)

if True == tmp1500 {
__e.Return(V3573)
return
} else {
tmp1498 := Call(__e, V3571, V3573)


__e.TailApply(PrimFunc(symshen_4fix_1help), V3571, V3573, tmp1498)
return


}


}, 3)

tmp1501 := Call(__e, ns2_1set, symshen_4fix_1help, tmp1497)


_ = tmp1501

tmp1502 := MakeNative(func(__e *ControlFlow) {
V3574 := __e.Get(1)
_ = V3574
V3575 := __e.Get(2)
_ = V3575
V3576 := __e.Get(3)
_ = V3576
V3577 := __e.Get(4)
_ = V3577
tmp1503 := MakeNative(func(__e *ControlFlow) {
W3578 := __e.Get(1)
_ = W3578
tmp1504 := MakeNative(func(__e *ControlFlow) {
W3579 := __e.Get(1)
_ = W3579
tmp1505 := MakeNative(func(__e *ControlFlow) {
W3581 := __e.Get(1)
_ = W3581
__e.Return(V3576)
return
}, 1)

tmp1506 := Call(__e, PrimFunc(symshen_4change_1pointer_1value), V3574, V3575, V3576, W3579)


tmp1507 := Call(__e, PrimFunc(symvector_1_6), V3577, W3578, tmp1506)


__e.TailApply(tmp1505, tmp1507)
return


}, 1)

tmp1508 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(sym_5_1vector), V3577, W3578)
return
}, 0)

tmp1509 := MakeNative(func(__e *ControlFlow) {
Z3580 := __e.Get(1)
_ = Z3580
__e.Return(Nil)
return
}, 1)

tmp1510 := Call(__e, try_1catch, tmp1508, tmp1509)


__e.TailApply(tmp1504, tmp1510)
return


}, 1)

tmp1511 := Call(__e, PrimFunc(symlimit), V3577)


tmp1512 := Call(__e, PrimFunc(symhash), V3574, tmp1511)


__e.TailApply(tmp1503, tmp1512)
return


}, 4)

tmp1513 := Call(__e, ns2_1set, symput, tmp1502)


_ = tmp1513

tmp1514 := MakeNative(func(__e *ControlFlow) {
V3582 := __e.Get(1)
_ = V3582
V3583 := __e.Get(2)
_ = V3583
V3584 := __e.Get(3)
_ = V3584
tmp1515 := MakeNative(func(__e *ControlFlow) {
W3585 := __e.Get(1)
_ = W3585
tmp1516 := MakeNative(func(__e *ControlFlow) {
W3586 := __e.Get(1)
_ = W3586
tmp1517 := MakeNative(func(__e *ControlFlow) {
W3588 := __e.Get(1)
_ = W3588
__e.Return(V3582)
return
}, 1)

tmp1518 := Call(__e, PrimFunc(symshen_4remove_1pointer), V3582, V3583, W3586)


tmp1519 := Call(__e, PrimFunc(symvector_1_6), V3584, W3585, tmp1518)


__e.TailApply(tmp1517, tmp1519)
return


}, 1)

tmp1520 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(sym_5_1vector), V3584, W3585)
return
}, 0)

tmp1521 := MakeNative(func(__e *ControlFlow) {
Z3587 := __e.Get(1)
_ = Z3587
__e.Return(Nil)
return
}, 1)

tmp1522 := Call(__e, try_1catch, tmp1520, tmp1521)


__e.TailApply(tmp1516, tmp1522)
return


}, 1)

tmp1523 := Call(__e, PrimFunc(symlimit), V3584)


tmp1524 := Call(__e, PrimFunc(symhash), V3582, tmp1523)


__e.TailApply(tmp1515, tmp1524)
return


}, 3)

tmp1525 := Call(__e, ns2_1set, symunput, tmp1514)


_ = tmp1525

tmp1526 := MakeNative(func(__e *ControlFlow) {
V3599 := __e.Get(1)
_ = V3599
V3600 := __e.Get(2)
_ = V3600
V3601 := __e.Get(3)
_ = V3601
tmp1570 := PrimEqual(Nil, V3601)

if True == tmp1570 {
__e.Return(Nil)
return
} else {
tmp1568 := PrimIsPair(V3601)

var ifres1533 Obj

if True == tmp1568 {
tmp1566 := PrimHead(V3601)

tmp1567 := PrimIsPair(tmp1566)

var ifres1535 Obj

if True == tmp1567 {
tmp1563 := PrimHead(V3601)

tmp1564 := PrimHead(tmp1563)

tmp1565 := PrimIsPair(tmp1564)

var ifres1537 Obj

if True == tmp1565 {
tmp1559 := PrimHead(V3601)

tmp1560 := PrimHead(tmp1559)

tmp1561 := PrimTail(tmp1560)

tmp1562 := PrimIsPair(tmp1561)

var ifres1539 Obj

if True == tmp1562 {
tmp1554 := PrimHead(V3601)

tmp1555 := PrimHead(tmp1554)

tmp1556 := PrimTail(tmp1555)

tmp1557 := PrimTail(tmp1556)

tmp1558 := PrimEqual(Nil, tmp1557)

var ifres1541 Obj

if True == tmp1558 {
tmp1549 := PrimHead(V3601)

tmp1550 := PrimHead(tmp1549)

tmp1551 := PrimTail(tmp1550)

tmp1552 := PrimHead(tmp1551)

tmp1553 := PrimEqual(V3600, tmp1552)

var ifres1543 Obj

if True == tmp1553 {
tmp1545 := PrimHead(V3601)

tmp1546 := PrimHead(tmp1545)

tmp1547 := PrimHead(tmp1546)

tmp1548 := PrimEqual(V3599, tmp1547)

var ifres1544 Obj

if True == tmp1548 {
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

var ifres1538 Obj

if True == ifres1539 {
ifres1538 = True


} else {
ifres1538 = False


}

ifres1537 = ifres1538


} else {
ifres1537 = False


}

var ifres1536 Obj

if True == ifres1537 {
ifres1536 = True


} else {
ifres1536 = False


}

ifres1535 = ifres1536


} else {
ifres1535 = False


}

var ifres1534 Obj

if True == ifres1535 {
ifres1534 = True


} else {
ifres1534 = False


}

ifres1533 = ifres1534


} else {
ifres1533 = False


}

if True == ifres1533 {
__e.Return(PrimTail(V3601))
return
} else {
tmp1531 := PrimIsPair(V3601)

if True == tmp1531 {
tmp1527 := PrimHead(V3601)

tmp1528 := PrimTail(V3601)

tmp1529 := Call(__e, PrimFunc(symshen_4remove_1pointer), V3599, V3600, tmp1528)


__e.Return(PrimCons(tmp1527, tmp1529))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.remove-pointer")))
return
}


}


}


}, 3)

tmp1571 := Call(__e, ns2_1set, symshen_4remove_1pointer, tmp1526)


_ = tmp1571

tmp1572 := MakeNative(func(__e *ControlFlow) {
V3614 := __e.Get(1)
_ = V3614
V3615 := __e.Get(2)
_ = V3615
V3616 := __e.Get(3)
_ = V3616
V3617 := __e.Get(4)
_ = V3617
tmp1623 := PrimEqual(Nil, V3617)

if True == tmp1623 {
tmp1573 := PrimCons(V3615, Nil)

tmp1574 := PrimCons(V3614, tmp1573)

tmp1575 := PrimCons(tmp1574, V3616)

__e.Return(PrimCons(tmp1575, Nil))
return


} else {
tmp1621 := PrimIsPair(V3617)

var ifres1586 Obj

if True == tmp1621 {
tmp1619 := PrimHead(V3617)

tmp1620 := PrimIsPair(tmp1619)

var ifres1588 Obj

if True == tmp1620 {
tmp1616 := PrimHead(V3617)

tmp1617 := PrimHead(tmp1616)

tmp1618 := PrimIsPair(tmp1617)

var ifres1590 Obj

if True == tmp1618 {
tmp1612 := PrimHead(V3617)

tmp1613 := PrimHead(tmp1612)

tmp1614 := PrimTail(tmp1613)

tmp1615 := PrimIsPair(tmp1614)

var ifres1592 Obj

if True == tmp1615 {
tmp1607 := PrimHead(V3617)

tmp1608 := PrimHead(tmp1607)

tmp1609 := PrimTail(tmp1608)

tmp1610 := PrimTail(tmp1609)

tmp1611 := PrimEqual(Nil, tmp1610)

var ifres1594 Obj

if True == tmp1611 {
tmp1602 := PrimHead(V3617)

tmp1603 := PrimHead(tmp1602)

tmp1604 := PrimTail(tmp1603)

tmp1605 := PrimHead(tmp1604)

tmp1606 := PrimEqual(V3615, tmp1605)

var ifres1596 Obj

if True == tmp1606 {
tmp1598 := PrimHead(V3617)

tmp1599 := PrimHead(tmp1598)

tmp1600 := PrimHead(tmp1599)

tmp1601 := PrimEqual(V3614, tmp1600)

var ifres1597 Obj

if True == tmp1601 {
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

var ifres1591 Obj

if True == ifres1592 {
ifres1591 = True


} else {
ifres1591 = False


}

ifres1590 = ifres1591


} else {
ifres1590 = False


}

var ifres1589 Obj

if True == ifres1590 {
ifres1589 = True


} else {
ifres1589 = False


}

ifres1588 = ifres1589


} else {
ifres1588 = False


}

var ifres1587 Obj

if True == ifres1588 {
ifres1587 = True


} else {
ifres1587 = False


}

ifres1586 = ifres1587


} else {
ifres1586 = False


}

if True == ifres1586 {
tmp1576 := PrimHead(V3617)

tmp1577 := PrimHead(tmp1576)

tmp1578 := PrimCons(tmp1577, V3616)

tmp1579 := PrimTail(V3617)

__e.Return(PrimCons(tmp1578, tmp1579))
return


} else {
tmp1584 := PrimIsPair(V3617)

if True == tmp1584 {
tmp1580 := PrimHead(V3617)

tmp1581 := PrimTail(V3617)

tmp1582 := Call(__e, PrimFunc(symshen_4change_1pointer_1value), V3614, V3615, V3616, tmp1581)


__e.Return(PrimCons(tmp1580, tmp1582))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.change-pointer-value")))
return
}


}


}


}, 4)

tmp1624 := Call(__e, ns2_1set, symshen_4change_1pointer_1value, tmp1572)


_ = tmp1624

tmp1625 := MakeNative(func(__e *ControlFlow) {
V3618 := __e.Get(1)
_ = V3618
V3619 := __e.Get(2)
_ = V3619
V3620 := __e.Get(3)
_ = V3620
tmp1626 := MakeNative(func(__e *ControlFlow) {
W3621 := __e.Get(1)
_ = W3621
tmp1627 := MakeNative(func(__e *ControlFlow) {
W3622 := __e.Get(1)
_ = W3622
tmp1628 := MakeNative(func(__e *ControlFlow) {
W3624 := __e.Get(1)
_ = W3624
tmp1634 := Call(__e, PrimFunc(symempty_2), W3624)


if True == tmp1634 {
tmp1629 := Call(__e, PrimFunc(symshen_4app), V3618, MakeString("\n"), symshen_4s)


tmp1630 := PrimStringConcat(MakeString(" not found for "), tmp1629)

tmp1631 := Call(__e, PrimFunc(symshen_4app), V3619, tmp1630, symshen_4s)


tmp1632 := PrimStringConcat(MakeString("attribute "), tmp1631)

__e.Return(PrimSimpleError(tmp1632))
return


} else {
__e.Return(PrimTail(W3624))
return
}


}, 1)

tmp1635 := PrimCons(V3619, Nil)

tmp1636 := PrimCons(V3618, tmp1635)

tmp1637 := Call(__e, PrimFunc(symassoc), tmp1636, W3622)


__e.TailApply(tmp1628, tmp1637)
return


}, 1)

tmp1638 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(sym_5_1vector), V3620, W3621)
return
}, 0)

tmp1639 := MakeNative(func(__e *ControlFlow) {
Z3623 := __e.Get(1)
_ = Z3623
tmp1640 := Call(__e, PrimFunc(symshen_4app), V3619, MakeString("\n"), symshen_4s)


tmp1641 := PrimStringConcat(MakeString(" has no attributes: "), tmp1640)

tmp1642 := Call(__e, PrimFunc(symshen_4app), V3618, tmp1641, symshen_4a)


__e.Return(PrimSimpleError(tmp1642))
return


}, 1)

tmp1643 := Call(__e, try_1catch, tmp1638, tmp1639)


__e.TailApply(tmp1627, tmp1643)
return


}, 1)

tmp1644 := Call(__e, PrimFunc(symlimit), V3620)


tmp1645 := Call(__e, PrimFunc(symhash), V3618, tmp1644)


__e.TailApply(tmp1626, tmp1645)
return


}, 3)

tmp1646 := Call(__e, ns2_1set, symget, tmp1625)


_ = tmp1646

tmp1647 := MakeNative(func(__e *ControlFlow) {
V3625 := __e.Get(1)
_ = V3625
V3626 := __e.Get(2)
_ = V3626
tmp1648 := MakeNative(func(__e *ControlFlow) {
W3627 := __e.Get(1)
_ = W3627
tmp1650 := PrimEqual(W3627, MakeNumber(0))

if True == tmp1650 {
__e.Return(MakeNumber(1))
return
} else {
__e.Return(W3627)
return
}


}, 1)

tmp1651 := Call(__e, PrimFunc(symshen_4hashkey), V3625)


tmp1652 := Call(__e, PrimFunc(symshen_4mod), tmp1651, V3626)


__e.TailApply(tmp1648, tmp1652)
return


}, 2)

tmp1653 := Call(__e, ns2_1set, symhash, tmp1647)


_ = tmp1653

tmp1654 := MakeNative(func(__e *ControlFlow) {
V3628 := __e.Get(1)
_ = V3628
tmp1655 := MakeNative(func(__e *ControlFlow) {
W3629 := __e.Get(1)
_ = W3629
__e.TailApply(PrimFunc(symshen_4prodbutzero), W3629, MakeNumber(1))
return
}, 1)

tmp1656 := MakeNative(func(__e *ControlFlow) {
Z3630 := __e.Get(1)
_ = Z3630
__e.Return(PrimStringToNumber(Z3630))
return
}, 1)

tmp1657 := Call(__e, PrimFunc(symexplode), V3628)


tmp1658 := Call(__e, PrimFunc(symmap), tmp1656, tmp1657)


__e.TailApply(tmp1655, tmp1658)
return


}, 1)

tmp1659 := Call(__e, ns2_1set, symshen_4hashkey, tmp1654)


_ = tmp1659

tmp1660 := MakeNative(func(__e *ControlFlow) {
V3631 := __e.Get(1)
_ = V3631
V3632 := __e.Get(2)
_ = V3632
tmp1679 := PrimEqual(Nil, V3631)

if True == tmp1679 {
__e.Return(V3632)
return
} else {
tmp1677 := PrimIsPair(V3631)

var ifres1673 Obj

if True == tmp1677 {
tmp1675 := PrimHead(V3631)

tmp1676 := PrimEqual(MakeNumber(0), tmp1675)

var ifres1674 Obj

if True == tmp1676 {
ifres1674 = True


} else {
ifres1674 = False


}

ifres1673 = ifres1674


} else {
ifres1673 = False


}

if True == ifres1673 {
tmp1661 := PrimTail(V3631)

__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp1661, V3632)
return


} else {
tmp1671 := PrimIsPair(V3631)

if True == tmp1671 {
tmp1669 := PrimGreatThan(V3632, MakeNumber(10000000000))

if True == tmp1669 {
tmp1662 := PrimTail(V3631)

tmp1663 := PrimHead(V3631)

tmp1664 := PrimNumberAdd(V3632, tmp1663)

__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp1662, tmp1664)
return


} else {
tmp1665 := PrimTail(V3631)

tmp1666 := PrimHead(V3631)

tmp1667 := PrimNumberMultiply(V3632, tmp1666)

__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp1665, tmp1667)
return


}


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.prodbutzero")))
return
}


}


}


}, 2)

tmp1680 := Call(__e, ns2_1set, symshen_4prodbutzero, tmp1660)


_ = tmp1680

tmp1681 := MakeNative(func(__e *ControlFlow) {
V3633 := __e.Get(1)
_ = V3633
V3634 := __e.Get(2)
_ = V3634
tmp1682 := PrimCons(V3634, Nil)

tmp1683 := Call(__e, PrimFunc(symshen_4multiples), V3633, tmp1682)


__e.TailApply(PrimFunc(symshen_4modh), V3633, tmp1683)
return


}, 2)

tmp1684 := Call(__e, ns2_1set, symshen_4mod, tmp1681)


_ = tmp1684

tmp1685 := MakeNative(func(__e *ControlFlow) {
V3639 := __e.Get(1)
_ = V3639
V3640 := __e.Get(2)
_ = V3640
tmp1696 := PrimIsPair(V3640)

var ifres1692 Obj

if True == tmp1696 {
tmp1694 := PrimHead(V3640)

tmp1695 := PrimGreatThan(tmp1694, V3639)

var ifres1693 Obj

if True == tmp1695 {
ifres1693 = True


} else {
ifres1693 = False


}

ifres1692 = ifres1693


} else {
ifres1692 = False


}

if True == ifres1692 {
__e.Return(PrimTail(V3640))
return
} else {
tmp1690 := PrimIsPair(V3640)

if True == tmp1690 {
tmp1686 := PrimHead(V3640)

tmp1687 := PrimNumberMultiply(MakeNumber(2), tmp1686)

tmp1688 := PrimCons(tmp1687, V3640)

__e.TailApply(PrimFunc(symshen_4multiples), V3639, tmp1688)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.multiples")))
return
}


}


}, 2)

tmp1697 := Call(__e, ns2_1set, symshen_4multiples, tmp1685)


_ = tmp1697

tmp1698 := MakeNative(func(__e *ControlFlow) {
V3647 := __e.Get(1)
_ = V3647
V3648 := __e.Get(2)
_ = V3648
tmp1716 := PrimEqual(MakeNumber(0), V3647)

if True == tmp1716 {
__e.Return(MakeNumber(0))
return
} else {
tmp1714 := PrimEqual(Nil, V3648)

if True == tmp1714 {
__e.Return(V3647)
return
} else {
tmp1712 := PrimIsPair(V3648)

var ifres1708 Obj

if True == tmp1712 {
tmp1710 := PrimHead(V3648)

tmp1711 := PrimGreatThan(tmp1710, V3647)

var ifres1709 Obj

if True == tmp1711 {
ifres1709 = True


} else {
ifres1709 = False


}

ifres1708 = ifres1709


} else {
ifres1708 = False


}

if True == ifres1708 {
tmp1701 := PrimTail(V3648)

tmp1702 := Call(__e, PrimFunc(symempty_2), tmp1701)


if True == tmp1702 {
__e.Return(V3647)
return
} else {
tmp1699 := PrimTail(V3648)

__e.TailApply(PrimFunc(symshen_4modh), V3647, tmp1699)
return


}


} else {
tmp1706 := PrimIsPair(V3648)

if True == tmp1706 {
tmp1703 := PrimHead(V3648)

tmp1704 := PrimNumberSubtract(V3647, tmp1703)

__e.TailApply(PrimFunc(symshen_4modh), tmp1704, V3648)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.modh")))
return
}


}


}


}


}, 2)

tmp1717 := Call(__e, ns2_1set, symshen_4modh, tmp1698)


_ = tmp1717

tmp1718 := MakeNative(func(__e *ControlFlow) {
V3651 := __e.Get(1)
_ = V3651
tmp1725 := PrimEqual(Nil, V3651)

if True == tmp1725 {
__e.Return(MakeNumber(0))
return
} else {
tmp1723 := PrimIsPair(V3651)

if True == tmp1723 {
tmp1719 := PrimHead(V3651)

tmp1720 := PrimTail(V3651)

tmp1721 := Call(__e, PrimFunc(symsum), tmp1720)


__e.Return(PrimNumberAdd(tmp1719, tmp1721))
return


} else {
__e.Return(PrimSimpleError(MakeString("attempt to sum a non-list\n")))
return
}


}


}, 1)

tmp1726 := Call(__e, ns2_1set, symsum, tmp1718)


_ = tmp1726

tmp1727 := MakeNative(func(__e *ControlFlow) {
V3656 := __e.Get(1)
_ = V3656
tmp1729 := PrimIsPair(V3656)

if True == tmp1729 {
__e.Return(PrimHead(V3656))
return
} else {
__e.Return(PrimSimpleError(MakeString("head expects a non-empty list\n")))
return
}


}, 1)

tmp1730 := Call(__e, ns2_1set, symhead, tmp1727)


_ = tmp1730

tmp1731 := MakeNative(func(__e *ControlFlow) {
V3661 := __e.Get(1)
_ = V3661
tmp1733 := PrimIsPair(V3661)

if True == tmp1733 {
__e.Return(PrimTail(V3661))
return
} else {
__e.Return(PrimSimpleError(MakeString("tail expects a non-empty list\n")))
return
}


}, 1)

tmp1734 := Call(__e, ns2_1set, symtail, tmp1731)


_ = tmp1734

tmp1735 := MakeNative(func(__e *ControlFlow) {
V3662 := __e.Get(1)
_ = V3662
__e.Return(PrimPos(V3662, MakeNumber(0)))
return
}, 1)

tmp1736 := Call(__e, ns2_1set, symhdstr, tmp1735)


_ = tmp1736

tmp1737 := MakeNative(func(__e *ControlFlow) {
V3669 := __e.Get(1)
_ = V3669
V3670 := __e.Get(2)
_ = V3670
tmp1748 := PrimEqual(Nil, V3669)

if True == tmp1748 {
__e.Return(Nil)
return
} else {
tmp1746 := PrimIsPair(V3669)

if True == tmp1746 {
tmp1743 := PrimHead(V3669)

tmp1744 := Call(__e, PrimFunc(symelement_2), tmp1743, V3670)


if True == tmp1744 {
tmp1738 := PrimHead(V3669)

tmp1739 := PrimTail(V3669)

tmp1740 := Call(__e, PrimFunc(symintersection), tmp1739, V3670)


__e.Return(PrimCons(tmp1738, tmp1740))
return


} else {
tmp1741 := PrimTail(V3669)

__e.TailApply(PrimFunc(symintersection), tmp1741, V3670)
return


}


} else {
__e.Return(PrimSimpleError(MakeString("attempt to find the intersection with a non-list\n")))
return
}


}


}, 2)

tmp1749 := Call(__e, ns2_1set, symintersection, tmp1737)


_ = tmp1749

tmp1750 := MakeNative(func(__e *ControlFlow) {
V3671 := __e.Get(1)
_ = V3671
__e.TailApply(PrimFunc(symshen_4reverse_1help), V3671, Nil)
return
}, 1)

tmp1751 := Call(__e, ns2_1set, symreverse, tmp1750)


_ = tmp1751

tmp1752 := MakeNative(func(__e *ControlFlow) {
V3676 := __e.Get(1)
_ = V3676
V3677 := __e.Get(2)
_ = V3677
tmp1759 := PrimEqual(Nil, V3676)

if True == tmp1759 {
__e.Return(V3677)
return
} else {
tmp1757 := PrimIsPair(V3676)

if True == tmp1757 {
tmp1753 := PrimTail(V3676)

tmp1754 := PrimHead(V3676)

tmp1755 := PrimCons(tmp1754, V3677)

__e.TailApply(PrimFunc(symshen_4reverse_1help), tmp1753, tmp1755)
return


} else {
__e.Return(PrimSimpleError(MakeString("attempt to reverse a non-list\n")))
return
}


}


}, 2)

tmp1760 := Call(__e, ns2_1set, symshen_4reverse_1help, tmp1752)


_ = tmp1760

tmp1761 := MakeNative(func(__e *ControlFlow) {
V3682 := __e.Get(1)
_ = V3682
V3683 := __e.Get(2)
_ = V3683
tmp1772 := PrimEqual(Nil, V3682)

if True == tmp1772 {
__e.Return(V3683)
return
} else {
tmp1770 := PrimIsPair(V3682)

if True == tmp1770 {
tmp1767 := PrimHead(V3682)

tmp1768 := Call(__e, PrimFunc(symelement_2), tmp1767, V3683)


if True == tmp1768 {
tmp1762 := PrimTail(V3682)

__e.TailApply(PrimFunc(symunion), tmp1762, V3683)
return


} else {
tmp1763 := PrimHead(V3682)

tmp1764 := PrimTail(V3682)

tmp1765 := Call(__e, PrimFunc(symunion), tmp1764, V3683)


__e.Return(PrimCons(tmp1763, tmp1765))
return


}


} else {
__e.Return(PrimSimpleError(MakeString("attempt to find the union with a non-list\n")))
return
}


}


}, 2)

tmp1773 := Call(__e, ns2_1set, symunion, tmp1761)


_ = tmp1773

tmp1774 := MakeNative(func(__e *ControlFlow) {
V3684 := __e.Get(1)
_ = V3684
tmp1775 := MakeNative(func(__e *ControlFlow) {
W3685 := __e.Get(1)
_ = W3685
tmp1776 := MakeNative(func(__e *ControlFlow) {
W3686 := __e.Get(1)
_ = W3686
tmp1777 := MakeNative(func(__e *ControlFlow) {
W3687 := __e.Get(1)
_ = W3687
tmp1783 := PrimEqual(MakeString("y"), W3687)

if True == tmp1783 {
__e.Return(True)
return
} else {
tmp1781 := PrimEqual(MakeString("n"), W3687)

if True == tmp1781 {
__e.Return(False)
return
} else {
tmp1778 := Call(__e, PrimFunc(symstoutput))


tmp1779 := Call(__e, PrimFunc(sympr), MakeString("please answer y or n\n"), tmp1778)


_ = tmp1779

__e.TailApply(PrimFunc(symy_1or_1n_2), V3684)
return


}


}


}, 1)

tmp1784 := Call(__e, PrimFunc(symstinput))


tmp1785 := Call(__e, PrimFunc(symread), tmp1784)


tmp1786 := Call(__e, PrimFunc(symshen_4app), tmp1785, MakeString(""), symshen_4s)


__e.TailApply(tmp1777, tmp1786)
return


}, 1)

tmp1787 := Call(__e, PrimFunc(symstoutput))


tmp1788 := Call(__e, PrimFunc(sympr), MakeString(" (y/n) "), tmp1787)


__e.TailApply(tmp1776, tmp1788)
return


}, 1)

tmp1789 := Call(__e, PrimFunc(symshen_4proc_1nl), V3684)


tmp1790 := Call(__e, PrimFunc(symstoutput))


tmp1791 := Call(__e, PrimFunc(sympr), tmp1789, tmp1790)


__e.TailApply(tmp1775, tmp1791)
return


}, 1)

tmp1792 := Call(__e, ns2_1set, symy_1or_1n_2, tmp1774)


_ = tmp1792

tmp1793 := MakeNative(func(__e *ControlFlow) {
V3688 := __e.Get(1)
_ = V3688
if True == V3688 {
__e.Return(False)
return
} else {
__e.Return(True)
return
}
}, 1)

tmp1795 := Call(__e, ns2_1set, symnot, tmp1793)


_ = tmp1795

tmp1796 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimSimpleError(MakeString("")))
return
}, 0)

tmp1797 := Call(__e, ns2_1set, symabort, tmp1796)


_ = tmp1797

tmp1798 := MakeNative(func(__e *ControlFlow) {
V3694 := __e.Get(1)
_ = V3694
V3695 := __e.Get(2)
_ = V3695
V3696 := __e.Get(3)
_ = V3696
tmp1806 := PrimEqual(V3695, V3696)

if True == tmp1806 {
__e.Return(V3694)
return
} else {
tmp1804 := PrimIsPair(V3696)

if True == tmp1804 {
tmp1799 := PrimHead(V3696)

tmp1800 := Call(__e, PrimFunc(symsubst), V3694, V3695, tmp1799)


tmp1801 := PrimTail(V3696)

tmp1802 := Call(__e, PrimFunc(symsubst), V3694, V3695, tmp1801)


__e.Return(PrimCons(tmp1800, tmp1802))
return


} else {
__e.Return(V3696)
return
}


}


}, 3)

tmp1807 := Call(__e, ns2_1set, symsubst, tmp1798)


_ = tmp1807

tmp1808 := MakeNative(func(__e *ControlFlow) {
V3697 := __e.Get(1)
_ = V3697
tmp1809 := Call(__e, PrimFunc(symshen_4app), V3697, MakeString(""), symshen_4a)


__e.TailApply(PrimFunc(symshen_4explode_1h), tmp1809)
return


}, 1)

tmp1810 := Call(__e, ns2_1set, symexplode, tmp1808)


_ = tmp1810

tmp1811 := MakeNative(func(__e *ControlFlow) {
V3700 := __e.Get(1)
_ = V3700
tmp1818 := PrimEqual(MakeString(""), V3700)

if True == tmp1818 {
__e.Return(Nil)
return
} else {
tmp1816 := Call(__e, PrimFunc(symshen_4_7string_2), V3700)


if True == tmp1816 {
tmp1812 := Call(__e, PrimFunc(symhdstr), V3700)


tmp1813 := PrimTailString(V3700)

tmp1814 := Call(__e, PrimFunc(symshen_4explode_1h), tmp1813)


__e.Return(PrimCons(tmp1812, tmp1814))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in explode-h")))
return
}


}


}, 1)

tmp1819 := Call(__e, ns2_1set, symshen_4explode_1h, tmp1811)


_ = tmp1819

tmp1820 := MakeNative(func(__e *ControlFlow) {
V3701 := __e.Get(1)
_ = V3701
tmp1823 := PrimEqual(V3701, MakeString(""))

var ifres1821 Obj

if True == tmp1823 {
ifres1821 = MakeString("")


} else {
tmp1822 := Call(__e, PrimFunc(symshen_4app), V3701, MakeString("/"), symshen_4a)


ifres1821 = tmp1822


}

__e.Return(PrimSet(sym_dhome_1directory_d, ifres1821))
return


}, 1)

tmp1824 := Call(__e, ns2_1set, symcd, tmp1820)


_ = tmp1824

tmp1825 := MakeNative(func(__e *ControlFlow) {
V3702 := __e.Get(1)
_ = V3702
V3703 := __e.Get(2)
_ = V3703
__e.TailApply(PrimFunc(symshen_4map_1h), V3702, V3703, Nil)
return
}, 2)

tmp1826 := Call(__e, ns2_1set, symmap, tmp1825)


_ = tmp1826

tmp1827 := MakeNative(func(__e *ControlFlow) {
V3704 := __e.Get(1)
_ = V3704
V3705 := __e.Get(2)
_ = V3705
V3706 := __e.Get(3)
_ = V3706
tmp1835 := PrimEqual(Nil, V3705)

if True == tmp1835 {
__e.TailApply(PrimFunc(symreverse), V3706)
return
} else {
tmp1833 := PrimIsPair(V3705)

if True == tmp1833 {
tmp1828 := PrimTail(V3705)

tmp1829 := PrimHead(V3705)

tmp1830 := Call(__e, V3704, tmp1829)


tmp1831 := PrimCons(tmp1830, V3706)

__e.TailApply(PrimFunc(symshen_4map_1h), V3704, tmp1828, tmp1831)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.map-h")))
return
}


}


}, 3)

tmp1836 := Call(__e, ns2_1set, symshen_4map_1h, tmp1827)


_ = tmp1836

tmp1837 := MakeNative(func(__e *ControlFlow) {
V3707 := __e.Get(1)
_ = V3707
__e.TailApply(PrimFunc(symshen_4length_1h), V3707, MakeNumber(0))
return
}, 1)

tmp1838 := Call(__e, ns2_1set, symlength, tmp1837)


_ = tmp1838

tmp1839 := MakeNative(func(__e *ControlFlow) {
V3712 := __e.Get(1)
_ = V3712
V3713 := __e.Get(2)
_ = V3713
tmp1843 := PrimEqual(Nil, V3712)

if True == tmp1843 {
__e.Return(V3713)
return
} else {
tmp1840 := PrimTail(V3712)

tmp1841 := PrimNumberAdd(V3713, MakeNumber(1))

__e.TailApply(PrimFunc(symshen_4length_1h), tmp1840, tmp1841)
return


}


}, 2)

tmp1844 := Call(__e, ns2_1set, symshen_4length_1h, tmp1839)


_ = tmp1844

tmp1845 := MakeNative(func(__e *ControlFlow) {
V3719 := __e.Get(1)
_ = V3719
V3720 := __e.Get(2)
_ = V3720
tmp1853 := PrimEqual(V3719, V3720)

if True == tmp1853 {
__e.Return(MakeNumber(1))
return
} else {
tmp1851 := PrimIsPair(V3720)

if True == tmp1851 {
tmp1846 := PrimHead(V3720)

tmp1847 := Call(__e, PrimFunc(symoccurrences), V3719, tmp1846)


tmp1848 := PrimTail(V3720)

tmp1849 := Call(__e, PrimFunc(symoccurrences), V3719, tmp1848)


__e.Return(PrimNumberAdd(tmp1847, tmp1849))
return


} else {
__e.Return(MakeNumber(0))
return
}


}


}, 2)

tmp1854 := Call(__e, ns2_1set, symoccurrences, tmp1845)


_ = tmp1854

tmp1855 := MakeNative(func(__e *ControlFlow) {
V3725 := __e.Get(1)
_ = V3725
V3726 := __e.Get(2)
_ = V3726
tmp1868 := PrimEqual(MakeNumber(1), V3725)

var ifres1865 Obj

if True == tmp1868 {
tmp1867 := PrimIsPair(V3726)

var ifres1866 Obj

if True == tmp1867 {
ifres1866 = True


} else {
ifres1866 = False


}

ifres1865 = ifres1866


} else {
ifres1865 = False


}

if True == ifres1865 {
__e.Return(PrimHead(V3726))
return
} else {
tmp1863 := PrimIsPair(V3726)

if True == tmp1863 {
tmp1856 := PrimNumberSubtract(V3725, MakeNumber(1))

tmp1857 := PrimTail(V3726)

__e.TailApply(PrimFunc(symnth), tmp1856, tmp1857)
return


} else {
tmp1858 := Call(__e, PrimFunc(symshen_4app), V3726, MakeString("\n"), symshen_4a)


tmp1859 := PrimStringConcat(MakeString(", "), tmp1858)

tmp1860 := Call(__e, PrimFunc(symshen_4app), V3725, tmp1859, symshen_4a)


tmp1861 := PrimStringConcat(MakeString("nth applied to "), tmp1860)

__e.Return(PrimSimpleError(tmp1861))
return


}


}


}, 2)

tmp1869 := Call(__e, ns2_1set, symnth, tmp1855)


_ = tmp1869

tmp1870 := MakeNative(func(__e *ControlFlow) {
V3727 := __e.Get(1)
_ = V3727
tmp1877 := PrimIsNumber(V3727)

if True == tmp1877 {
tmp1872 := MakeNative(func(__e *ControlFlow) {
W3728 := __e.Get(1)
_ = W3728
tmp1873 := Call(__e, PrimFunc(symshen_4magless), W3728, MakeNumber(1))


__e.TailApply(PrimFunc(symshen_4integer_1test_2), W3728, tmp1873)
return


}, 1)

tmp1874 := Call(__e, PrimFunc(symshen_4abs), V3727)


tmp1875 := Call(__e, tmp1872, tmp1874)


if True == tmp1875 {
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

tmp1878 := Call(__e, ns2_1set, syminteger_2, tmp1870)


_ = tmp1878

tmp1879 := MakeNative(func(__e *ControlFlow) {
V3729 := __e.Get(1)
_ = V3729
tmp1881 := PrimGreatThan(V3729, MakeNumber(0))

if True == tmp1881 {
__e.Return(V3729)
return
} else {
__e.Return(PrimNumberSubtract(MakeNumber(0), V3729))
return
}


}, 1)

tmp1882 := Call(__e, ns2_1set, symshen_4abs, tmp1879)


_ = tmp1882

tmp1883 := MakeNative(func(__e *ControlFlow) {
V3730 := __e.Get(1)
_ = V3730
V3731 := __e.Get(2)
_ = V3731
tmp1884 := MakeNative(func(__e *ControlFlow) {
W3732 := __e.Get(1)
_ = W3732
tmp1886 := PrimGreatThan(W3732, V3730)

if True == tmp1886 {
__e.Return(V3731)
return
} else {
__e.TailApply(PrimFunc(symshen_4magless), V3730, W3732)
return
}


}, 1)

tmp1887 := PrimNumberMultiply(V3731, MakeNumber(2))

__e.TailApply(tmp1884, tmp1887)
return


}, 2)

tmp1888 := Call(__e, ns2_1set, symshen_4magless, tmp1883)


_ = tmp1888

tmp1889 := MakeNative(func(__e *ControlFlow) {
V3736 := __e.Get(1)
_ = V3736
V3737 := __e.Get(2)
_ = V3737
tmp1897 := PrimEqual(MakeNumber(0), V3736)

if True == tmp1897 {
__e.Return(True)
return
} else {
tmp1895 := PrimGreatThan(MakeNumber(1), V3736)

if True == tmp1895 {
__e.Return(False)
return
} else {
tmp1890 := MakeNative(func(__e *ControlFlow) {
W3738 := __e.Get(1)
_ = W3738
tmp1892 := PrimGreatThan(MakeNumber(0), W3738)

if True == tmp1892 {
__e.Return(PrimIsInteger(V3736))
return
} else {
__e.TailApply(PrimFunc(symshen_4integer_1test_2), W3738, V3737)
return
}


}, 1)

tmp1893 := PrimNumberSubtract(V3736, V3737)

__e.TailApply(tmp1890, tmp1893)
return


}


}


}, 2)

tmp1898 := Call(__e, ns2_1set, symshen_4integer_1test_2, tmp1889)


_ = tmp1898

tmp1899 := MakeNative(func(__e *ControlFlow) {
V3745 := __e.Get(1)
_ = V3745
V3746 := __e.Get(2)
_ = V3746
tmp1907 := PrimEqual(Nil, V3746)

if True == tmp1907 {
__e.Return(Nil)
return
} else {
tmp1905 := PrimIsPair(V3746)

if True == tmp1905 {
tmp1900 := PrimHead(V3746)

tmp1901 := Call(__e, V3745, tmp1900)


tmp1902 := PrimTail(V3746)

tmp1903 := Call(__e, PrimFunc(symmapcan), V3745, tmp1902)


__e.TailApply(PrimFunc(symappend), tmp1901, tmp1903)
return


} else {
__e.Return(PrimSimpleError(MakeString("attempt to mapcan over a non-list\n")))
return
}


}


}, 2)

tmp1908 := Call(__e, ns2_1set, symmapcan, tmp1899)


_ = tmp1908

tmp1909 := MakeNative(func(__e *ControlFlow) {
V3752 := __e.Get(1)
_ = V3752
V3753 := __e.Get(2)
_ = V3753
tmp1911 := PrimEqual(V3752, V3753)

if True == tmp1911 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp1912 := Call(__e, ns2_1set, sym_a_a, tmp1909)


_ = tmp1912

tmp1913 := MakeNative(func(__e *ControlFlow) {
V3754 := __e.Get(1)
_ = V3754
tmp1923 := PrimIsSymbol(V3754)

if True == tmp1923 {
tmp1915 := MakeNative(func(__e *ControlFlow) {
W3755 := __e.Get(1)
_ = W3755
tmp1917 := PrimEqual(W3755, symshen_4this_1symbol_1is_1unbound)

if True == tmp1917 {
__e.Return(False)
return
} else {
__e.Return(True)
return
}


}, 1)

tmp1918 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(V3754))
return
}, 0)

tmp1919 := MakeNative(func(__e *ControlFlow) {
Z3756 := __e.Get(1)
_ = Z3756
__e.Return(symshen_4this_1symbol_1is_1unbound)
return
}, 1)

tmp1920 := Call(__e, try_1catch, tmp1918, tmp1919)


tmp1921 := Call(__e, tmp1915, tmp1920)


if True == tmp1921 {
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

tmp1924 := Call(__e, ns2_1set, symbound_2, tmp1913)


_ = tmp1924

tmp1925 := MakeNative(func(__e *ControlFlow) {
V3757 := __e.Get(1)
_ = V3757
tmp1931 := PrimEqual(MakeString(""), V3757)

if True == tmp1931 {
__e.Return(Nil)
return
} else {
tmp1926 := PrimPos(V3757, MakeNumber(0))

tmp1927 := PrimStringToNumber(tmp1926)

tmp1928 := PrimTailString(V3757)

tmp1929 := Call(__e, PrimFunc(symshen_4string_1_6bytes), tmp1928)


__e.Return(PrimCons(tmp1927, tmp1929))
return


}


}, 1)

tmp1932 := Call(__e, ns2_1set, symshen_4string_1_6bytes, tmp1925)


_ = tmp1932

tmp1933 := MakeNative(func(__e *ControlFlow) {
V3758 := __e.Get(1)
_ = V3758
tmp1937 := PrimLessThan(V3758, MakeNumber(0))

if True == tmp1937 {
__e.Return(PrimValue(symshen_4_dmaxinferences_d))
return
} else {
tmp1935 := PrimIsInteger(V3758)

if True == tmp1935 {
__e.Return(PrimSet(symshen_4_dmaxinferences_d, V3758))
return
} else {
__e.Return(PrimSimpleError(MakeString("maxinferences expects an integer value\n")))
return
}


}


}, 1)

tmp1938 := Call(__e, ns2_1set, symmaxinferences, tmp1933)


_ = tmp1938

tmp1939 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dinfs_d))
return
}, 0)

tmp1940 := Call(__e, ns2_1set, syminferences, tmp1939)


_ = tmp1940

tmp1941 := MakeNative(func(__e *ControlFlow) {
V3759 := __e.Get(1)
_ = V3759
__e.Return(V3759)
return
}, 1)

tmp1942 := Call(__e, ns2_1set, symprotect, tmp1941)


_ = tmp1942

tmp1943 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dstoutput_d))
return
}, 0)

tmp1944 := Call(__e, ns2_1set, symstoutput, tmp1943)


_ = tmp1944

tmp1945 := MakeNative(func(__e *ControlFlow) {
V3760 := __e.Get(1)
_ = V3760
tmp1946 := MakeNative(func(__e *ControlFlow) {
W3761 := __e.Get(1)
_ = W3761
tmp1950 := PrimIsSymbol(W3761)

if True == tmp1950 {
__e.Return(W3761)
return
} else {
tmp1947 := Call(__e, PrimFunc(symshen_4app), V3760, MakeString(" to a symbol"), symshen_4s)


tmp1948 := PrimStringConcat(MakeString("cannot intern "), tmp1947)

__e.Return(PrimSimpleError(tmp1948))
return


}


}, 1)

tmp1951 := PrimIntern(V3760)

__e.TailApply(tmp1946, tmp1951)
return


}, 1)

tmp1952 := Call(__e, ns2_1set, symstring_1_6symbol, tmp1945)


_ = tmp1952

tmp1953 := MakeNative(func(__e *ControlFlow) {
V3764 := __e.Get(1)
_ = V3764
tmp1957 := PrimEqual(sym_7, V3764)

if True == tmp1957 {
__e.Return(PrimSet(symshen_4_doptimise_d, True))
return
} else {
tmp1955 := PrimEqual(sym_1, V3764)

if True == tmp1955 {
__e.Return(PrimSet(symshen_4_doptimise_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("optimise expects a + or a -.\n")))
return
}


}


}, 1)

tmp1958 := Call(__e, ns2_1set, symoptimise, tmp1953)


_ = tmp1958

tmp1959 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dos_d))
return
}, 0)

tmp1960 := Call(__e, ns2_1set, symos, tmp1959)


_ = tmp1960

tmp1961 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dlanguage_d))
return
}, 0)

tmp1962 := Call(__e, ns2_1set, symlanguage, tmp1961)


_ = tmp1962

tmp1963 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dversion_d))
return
}, 0)

tmp1964 := Call(__e, ns2_1set, symversion, tmp1963)


_ = tmp1964

tmp1965 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dport_d))
return
}, 0)

tmp1966 := Call(__e, ns2_1set, symport, tmp1965)


_ = tmp1966

tmp1967 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dporters_d))
return
}, 0)

tmp1968 := Call(__e, ns2_1set, symporters, tmp1967)


_ = tmp1968

tmp1969 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dimplementation_d))
return
}, 0)

tmp1970 := Call(__e, ns2_1set, symimplementation, tmp1969)


_ = tmp1970

tmp1971 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_drelease_d))
return
}, 0)

tmp1972 := Call(__e, ns2_1set, symrelease, tmp1971)


_ = tmp1972

tmp1973 := MakeNative(func(__e *ControlFlow) {
V3765 := __e.Get(1)
_ = V3765
tmp1978 := PrimEqual(symnull, V3765)

if True == tmp1978 {
__e.Return(True)
return
} else {
tmp1974 := MakeNative(func(__e *ControlFlow) {
tmp1975 := Call(__e, PrimFunc(symexternal), V3765)


_ = tmp1975

__e.Return(True)
return


}, 0)

tmp1976 := MakeNative(func(__e *ControlFlow) {
Z3766 := __e.Get(1)
_ = Z3766
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp1974, tmp1976)
return


}


}, 1)

tmp1979 := Call(__e, ns2_1set, sympackage_2, tmp1973)


_ = tmp1979

tmp1980 := MakeNative(func(__e *ControlFlow) {
__e.Return(sym_4_4_4)
return
}, 0)

tmp1981 := Call(__e, ns2_1set, symfail, tmp1980)


_ = tmp1981

tmp1982 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_duserdefs_d))
return
}, 0)

tmp1983 := Call(__e, ns2_1set, symuserdefs, tmp1982)


_ = tmp1983

tmp1984 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_doptimise_d))
return
}, 0)

tmp1985 := Call(__e, ns2_1set, symoptimise_2, tmp1984)


_ = tmp1985

tmp1986 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dhush_d))
return
}, 0)

tmp1987 := Call(__e, ns2_1set, symhush_2, tmp1986)


_ = tmp1987

tmp1988 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dshen_1type_1theory_1enabled_2_d))
return
}, 0)

tmp1989 := Call(__e, ns2_1set, symsystem_1S_2, tmp1988)


_ = tmp1989

tmp1990 := MakeNative(func(__e *ControlFlow) {
V3769 := __e.Get(1)
_ = V3769
tmp1994 := PrimEqual(sym_7, V3769)

if True == tmp1994 {
__e.Return(PrimSet(symshen_4_dshen_1type_1theory_1enabled_2_d, True))
return
} else {
tmp1992 := PrimEqual(sym_1, V3769)

if True == tmp1992 {
__e.Return(PrimSet(symshen_4_dshen_1type_1theory_1enabled_2_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("enable-type-theory expects a + or a -\n")))
return
}


}


}, 1)

tmp1995 := Call(__e, ns2_1set, symenable_1type_1theory, tmp1990)


_ = tmp1995

tmp1996 := MakeNative(func(__e *ControlFlow) {
V3772 := __e.Get(1)
_ = V3772
tmp2000 := PrimEqual(sym_7, V3772)

if True == tmp2000 {
__e.Return(PrimSet(sym_dhush_d, True))
return
} else {
tmp1998 := PrimEqual(sym_1, V3772)

if True == tmp1998 {
__e.Return(PrimSet(sym_dhush_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("hush expects a + or a -\n")))
return
}


}


}, 1)

tmp2001 := Call(__e, ns2_1set, symshen_4hush, tmp1996)


_ = tmp2001

tmp2002 := MakeNative(func(__e *ControlFlow) {
V3775 := __e.Get(1)
_ = V3775
tmp2006 := PrimEqual(sym_7, V3775)

if True == tmp2006 {
__e.Return(PrimSet(symshen_4_dtc_d, True))
return
} else {
tmp2004 := PrimEqual(sym_1, V3775)

if True == tmp2004 {
__e.Return(PrimSet(symshen_4_dtc_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("tc expects a + or -")))
return
}


}


}, 1)

tmp2007 := Call(__e, ns2_1set, symtc, tmp2002)


_ = tmp2007

tmp2008 := MakeNative(func(__e *ControlFlow) {
V3776 := __e.Get(1)
_ = V3776
tmp2009 := PrimValue(symshen_4_dsigf_d)

tmp2010 := Call(__e, PrimFunc(symshen_4unassoc), V3776, tmp2009)


tmp2011 := PrimSet(symshen_4_dsigf_d, tmp2010)

_ = tmp2011

__e.Return(V3776)
return


}, 1)

tmp2012 := Call(__e, ns2_1set, symdestroy, tmp2008)


_ = tmp2012

tmp2013 := MakeNative(func(__e *ControlFlow) {
V3786 := __e.Get(1)
_ = V3786
V3787 := __e.Get(2)
_ = V3787
tmp2031 := PrimEqual(Nil, V3787)

if True == tmp2031 {
__e.Return(Nil)
return
} else {
tmp2029 := PrimIsPair(V3787)

var ifres2020 Obj

if True == tmp2029 {
tmp2027 := PrimHead(V3787)

tmp2028 := PrimIsPair(tmp2027)

var ifres2022 Obj

if True == tmp2028 {
tmp2024 := PrimHead(V3787)

tmp2025 := PrimHead(tmp2024)

tmp2026 := PrimEqual(V3786, tmp2025)

var ifres2023 Obj

if True == tmp2026 {
ifres2023 = True


} else {
ifres2023 = False


}

ifres2022 = ifres2023


} else {
ifres2022 = False


}

var ifres2021 Obj

if True == ifres2022 {
ifres2021 = True


} else {
ifres2021 = False


}

ifres2020 = ifres2021


} else {
ifres2020 = False


}

if True == ifres2020 {
__e.Return(PrimTail(V3787))
return
} else {
tmp2018 := PrimIsPair(V3787)

if True == tmp2018 {
tmp2014 := PrimHead(V3787)

tmp2015 := PrimTail(V3787)

tmp2016 := Call(__e, PrimFunc(symshen_4unassoc), V3786, tmp2015)


__e.Return(PrimCons(tmp2014, tmp2016))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.unassoc")))
return
}


}


}


}, 2)

tmp2032 := Call(__e, ns2_1set, symshen_4unassoc, tmp2013)


_ = tmp2032

tmp2033 := MakeNative(func(__e *ControlFlow) {
V3788 := __e.Get(1)
_ = V3788
tmp2037 := Call(__e, PrimFunc(sympackage_2), V3788)


if True == tmp2037 {
__e.Return(PrimSet(symshen_4_dpackage_d, V3788))
return
} else {
tmp2034 := Call(__e, PrimFunc(symshen_4app), V3788, MakeString(" does not exist\n"), symshen_4a)


tmp2035 := PrimStringConcat(MakeString("package "), tmp2034)

__e.Return(PrimSimpleError(tmp2035))
return


}


}, 1)

tmp2038 := Call(__e, ns2_1set, symin_1package, tmp2033)


_ = tmp2038

tmp2039 := MakeNative(func(__e *ControlFlow) {
V3789 := __e.Get(1)
_ = V3789
V3790 := __e.Get(2)
_ = V3790
tmp2040 := MakeNative(func(__e *ControlFlow) {
W3791 := __e.Get(1)
_ = W3791
tmp2041 := MakeNative(func(__e *ControlFlow) {
W3792 := __e.Get(1)
_ = W3792
tmp2042 := MakeNative(func(__e *ControlFlow) {
W3793 := __e.Get(1)
_ = W3793
tmp2043 := MakeNative(func(__e *ControlFlow) {
W3794 := __e.Get(1)
_ = W3794
__e.Return(V3790)
return
}, 1)

tmp2044 := PrimCloseStream(W3791)

__e.TailApply(tmp2043, tmp2044)
return


}, 1)

tmp2045 := Call(__e, PrimFunc(sympr), W3792, W3791)


__e.TailApply(tmp2042, tmp2045)
return


}, 1)

tmp2048 := PrimIsString(V3790)

var ifres2046 Obj

if True == tmp2048 {
ifres2046 = V3790


} else {
tmp2047 := Call(__e, PrimFunc(symshen_4app), V3790, MakeString(""), symshen_4s)


ifres2046 = tmp2047


}

__e.TailApply(tmp2041, ifres2046)
return


}, 1)

tmp2049 := PrimOpenStream(V3789, symout)

__e.TailApply(tmp2040, tmp2049)
return


}, 2)

tmp2050 := Call(__e, ns2_1set, symwrite_1to_1file, tmp2039)


_ = tmp2050

tmp2051 := MakeNative(func(__e *ControlFlow) {
tmp2052 := Call(__e, PrimFunc(symgensym), symshen_4t)


__e.TailApply(PrimFunc(symshen_4freshterm), tmp2052)
return


}, 0)

tmp2053 := Call(__e, ns2_1set, symfresh, tmp2051)


_ = tmp2053

tmp2054 := MakeNative(func(__e *ControlFlow) {
V3795 := __e.Get(1)
_ = V3795
V3796 := __e.Get(2)
_ = V3796
tmp2055 := MakeNative(func(__e *ControlFlow) {
W3797 := __e.Get(1)
_ = W3797
tmp2056 := MakeNative(func(__e *ControlFlow) {
W3798 := __e.Get(1)
_ = W3798
tmp2057 := MakeNative(func(__e *ControlFlow) {
W3799 := __e.Get(1)
_ = W3799
__e.Return(V3795)
return
}, 1)

tmp2058 := PrimValue(symshen_4_dlambdatable_d)

tmp2059 := PrimCons(W3798, tmp2058)

tmp2060 := PrimSet(symshen_4_dlambdatable_d, tmp2059)

__e.TailApply(tmp2057, tmp2060)
return


}, 1)

tmp2061 := Call(__e, PrimFunc(symshen_4lambda_1entry), V3795)


__e.TailApply(tmp2056, tmp2061)
return


}, 1)

tmp2062 := PrimValue(sym_dproperty_1vector_d)

tmp2063 := Call(__e, PrimFunc(symput), V3795, symarity, V3796, tmp2062)


__e.TailApply(tmp2055, tmp2063)
return


}, 2)

tmp2064 := Call(__e, ns2_1set, symupdate_1lambda_1table, tmp2054)


_ = tmp2064

tmp2065 := MakeNative(func(__e *ControlFlow) {
V3802 := __e.Get(1)
_ = V3802
V3803 := __e.Get(2)
_ = V3803
tmp2089 := PrimEqual(MakeNumber(0), V3803)

if True == tmp2089 {
tmp2066 := PrimValue(symshen_4_dspecial_d)

tmp2067 := Call(__e, PrimFunc(symremove), V3802, tmp2066)


tmp2068 := PrimSet(symshen_4_dspecial_d, tmp2067)

_ = tmp2068

tmp2069 := PrimValue(symshen_4_dextraspecial_d)

tmp2070 := Call(__e, PrimFunc(symremove), V3802, tmp2069)


tmp2071 := PrimSet(symshen_4_dextraspecial_d, tmp2070)

_ = tmp2071

__e.Return(V3802)
return


} else {
tmp2087 := PrimEqual(MakeNumber(1), V3803)

if True == tmp2087 {
tmp2072 := PrimValue(symshen_4_dspecial_d)

tmp2073 := Call(__e, PrimFunc(symadjoin), V3802, tmp2072)


tmp2074 := PrimSet(symshen_4_dspecial_d, tmp2073)

_ = tmp2074

tmp2075 := PrimValue(symshen_4_dextraspecial_d)

tmp2076 := Call(__e, PrimFunc(symremove), V3802, tmp2075)


tmp2077 := PrimSet(symshen_4_dextraspecial_d, tmp2076)

_ = tmp2077

__e.Return(V3802)
return


} else {
tmp2085 := PrimEqual(MakeNumber(2), V3803)

if True == tmp2085 {
tmp2078 := PrimValue(symshen_4_dspecial_d)

tmp2079 := Call(__e, PrimFunc(symremove), V3802, tmp2078)


tmp2080 := PrimSet(symshen_4_dspecial_d, tmp2079)

_ = tmp2080

tmp2081 := PrimValue(symshen_4_dextraspecial_d)

tmp2082 := Call(__e, PrimFunc(symadjoin), V3802, tmp2081)


tmp2083 := PrimSet(symshen_4_dextraspecial_d, tmp2082)

_ = tmp2083

__e.Return(V3802)
return


} else {
__e.Return(PrimSimpleError(MakeString("specialise requires values of 0, 1 or 2\n")))
return
}


}


}


}, 2)

tmp2090 := Call(__e, ns2_1set, symspecialise, tmp2065)


_ = tmp2090

tmp2091 := MakeNative(func(__e *ControlFlow) {
V3804 := __e.Get(1)
_ = V3804
tmp2092 := PrimValue(sym_dabsolute_d)

tmp2093 := PrimCons(V3804, tmp2092)

__e.Return(PrimSet(sym_dabsolute_d, tmp2093))
return


}, 1)

tmp2094 := Call(__e, ns2_1set, symabsolute, tmp2091)


_ = tmp2094

tmp2095 := MakeNative(func(__e *ControlFlow) {
V3805 := __e.Get(1)
_ = V3805
tmp2096 := PrimValue(sym_dabsolute_d)

tmp2097 := Call(__e, PrimFunc(symremove), V3805, tmp2096)


__e.Return(PrimSet(sym_dabsolute_d, tmp2097))
return


}, 1)

__e.TailApply(ns2_1set, symunabsolute, tmp2095)
return




}, 0)

