package main

import . "github.com/tiancaiamao/shen-go/kl"

var PrologMain = MakeNative(func(__e *ControlFlow) {
tmp10088 := MakeNative(func(__e *ControlFlow) {
V1231 := __e.Get(1)
_ = V1231
__e.TailApply(PrimFunc(symshen_4assert_d), V1231, symshen_4top)
return
}, 1)

tmp10089 := Call(__e, ns2_1set, symasserta, tmp10088)


_ = tmp10089

tmp10090 := MakeNative(func(__e *ControlFlow) {
V1232 := __e.Get(1)
_ = V1232
__e.TailApply(PrimFunc(symshen_4assert_d), V1232, symshen_4bottom)
return
}, 1)

tmp10091 := Call(__e, ns2_1set, symassertz, tmp10090)


_ = tmp10091

tmp10092 := MakeNative(func(__e *ControlFlow) {
V1233 := __e.Get(1)
_ = V1233
V1234 := __e.Get(2)
_ = V1234
tmp10126 := PrimIsPair(V1233)

var ifres10117 Obj

if True == tmp10126 {
tmp10124 := PrimTail(V1233)

tmp10125 := PrimIsPair(tmp10124)

var ifres10119 Obj

if True == tmp10125 {
tmp10121 := PrimTail(V1233)

tmp10122 := PrimHead(tmp10121)

tmp10123 := PrimEqual(sym_5_1_1, tmp10122)

var ifres10120 Obj

if True == tmp10123 {
ifres10120 = True


} else {
ifres10120 = False


}

ifres10119 = ifres10120


} else {
ifres10119 = False


}

var ifres10118 Obj

if True == ifres10119 {
ifres10118 = True


} else {
ifres10118 = False


}

ifres10117 = ifres10118


} else {
ifres10117 = False


}

if True == ifres10117 {
tmp10093 := MakeNative(func(__e *ControlFlow) {
W1235 := __e.Get(1)
_ = W1235
tmp10094 := MakeNative(func(__e *ControlFlow) {
W1236 := __e.Get(1)
_ = W1236
tmp10095 := MakeNative(func(__e *ControlFlow) {
W1237 := __e.Get(1)
_ = W1237
tmp10096 := MakeNative(func(__e *ControlFlow) {
W1238 := __e.Get(1)
_ = W1238
tmp10097 := MakeNative(func(__e *ControlFlow) {
W1239 := __e.Get(1)
_ = W1239
tmp10098 := MakeNative(func(__e *ControlFlow) {
W1240 := __e.Get(1)
_ = W1240
tmp10099 := MakeNative(func(__e *ControlFlow) {
W1241 := __e.Get(1)
_ = W1241
__e.Return(W1235)
return
}, 1)

tmp10100 := PrimTail(V1233)

tmp10101 := PrimTail(tmp10100)

tmp10102 := Call(__e, PrimFunc(symshen_4insert_1info), W1235, W1236, tmp10101, V1233, V1234)


__e.TailApply(tmp10099, tmp10102)
return


}, 1)

tmp10108 := PrimEqual(W1239, MakeNumber(-1))

var ifres10103 Obj

if True == tmp10108 {
tmp10104 := Call(__e, PrimFunc(symshen_4create_1skeleton), W1235, W1238)


tmp10105 := Call(__e, PrimFunc(symeval), tmp10104)


_ = tmp10105

tmp10106 := PrimValue(sym_dproperty_1vector_d)

tmp10107 := Call(__e, PrimFunc(symput), W1235, symshen_4dynamic, Nil, tmp10106)


ifres10103 = tmp10107


} else {
ifres10103 = symshen_4skip


}

__e.TailApply(tmp10098, ifres10103)
return


}, 1)

tmp10109 := Call(__e, PrimFunc(symarity), W1235)


__e.TailApply(tmp10097, tmp10109)
return


}, 1)

tmp10110 := Call(__e, PrimFunc(symshen_4parameters), W1237)


__e.TailApply(tmp10096, tmp10110)
return


}, 1)

tmp10111 := Call(__e, PrimFunc(symlength), W1236)


__e.TailApply(tmp10095, tmp10111)
return


}, 1)

tmp10112 := PrimHead(V1233)

tmp10113 := Call(__e, PrimFunc(symshen_4terms), tmp10112)


__e.TailApply(tmp10094, tmp10113)
return


}, 1)

tmp10114 := PrimHead(V1233)

tmp10115 := Call(__e, PrimFunc(symshen_4predicate), tmp10114)


__e.TailApply(tmp10093, tmp10115)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.assert*")))
return
}


}, 2)

tmp10127 := Call(__e, ns2_1set, symshen_4assert_d, tmp10092)


_ = tmp10127

tmp10128 := MakeNative(func(__e *ControlFlow) {
V1244 := __e.Get(1)
_ = V1244
tmp10130 := PrimIsPair(V1244)

if True == tmp10130 {
__e.Return(PrimHead(V1244))
return
} else {
__e.Return(V1244)
return
}


}, 1)

tmp10131 := Call(__e, ns2_1set, symshen_4predicate, tmp10128)


_ = tmp10131

tmp10132 := MakeNative(func(__e *ControlFlow) {
V1249 := __e.Get(1)
_ = V1249
tmp10134 := PrimIsPair(V1249)

if True == tmp10134 {
__e.Return(PrimTail(V1249))
return
} else {
__e.Return(Nil)
return
}


}, 1)

tmp10135 := Call(__e, ns2_1set, symshen_4terms, tmp10132)


_ = tmp10135

tmp10136 := MakeNative(func(__e *ControlFlow) {
V1250 := __e.Get(1)
_ = V1250
V1251 := __e.Get(2)
_ = V1251
tmp10137 := Call(__e, PrimFunc(symshen_4dynamic_1default), V1250, V1251)


tmp10138 := PrimCons(V1250, tmp10137)

__e.Return(PrimCons(symdefprolog, tmp10138))
return


}, 2)

tmp10139 := Call(__e, ns2_1set, symshen_4create_1skeleton, tmp10136)


_ = tmp10139

tmp10140 := MakeNative(func(__e *ControlFlow) {
V1252 := __e.Get(1)
_ = V1252
V1253 := __e.Get(2)
_ = V1253
tmp10141 := Call(__e, PrimFunc(symshen_4cons_1form), V1253)


tmp10142 := PrimCons(symshen_4dynamic, Nil)

tmp10143 := PrimCons(V1252, tmp10142)

tmp10144 := PrimCons(symget, tmp10143)

tmp10145 := PrimCons(tmp10144, Nil)

tmp10146 := PrimCons(tmp10141, tmp10145)

tmp10147 := PrimCons(symshen_4call_1dynamic, tmp10146)

tmp10148 := PrimIntern(MakeString(";"))

tmp10149 := PrimCons(tmp10148, Nil)

tmp10150 := PrimCons(tmp10147, tmp10149)

tmp10151 := PrimCons(sym_5_1_1, tmp10150)

__e.TailApply(PrimFunc(symappend), V1253, tmp10151)
return


}, 2)

tmp10152 := Call(__e, ns2_1set, symshen_4dynamic_1default, tmp10140)


_ = tmp10152

tmp10153 := MakeNative(func(__e *ControlFlow) {
V1254 := __e.Get(1)
_ = V1254
V1255 := __e.Get(2)
_ = V1255
V1256 := __e.Get(3)
_ = V1256
V1257 := __e.Get(4)
_ = V1257
V1258 := __e.Get(5)
_ = V1258
tmp10154 := MakeNative(func(__e *ControlFlow) {
W1259 := __e.Get(1)
_ = W1259
tmp10155 := MakeNative(func(__e *ControlFlow) {
W1260 := __e.Get(1)
_ = W1260
tmp10156 := MakeNative(func(__e *ControlFlow) {
W1261 := __e.Get(1)
_ = W1261
tmp10157 := MakeNative(func(__e *ControlFlow) {
W1262 := __e.Get(1)
_ = W1262
tmp10158 := MakeNative(func(__e *ControlFlow) {
W1263 := __e.Get(1)
_ = W1263
tmp10159 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V1254, symshen_4dynamic, W1263, tmp10159)
return


}, 1)

tmp10164 := PrimEqual(V1258, symshen_4top)

var ifres10160 Obj

if True == tmp10164 {
tmp10161 := PrimCons(W1261, W1262)

ifres10160 = tmp10161


} else {
tmp10162 := PrimCons(W1261, Nil)

tmp10163 := Call(__e, PrimFunc(symappend), W1262, tmp10162)


ifres10160 = tmp10163


}

__e.TailApply(tmp10158, ifres10160)
return


}, 1)

tmp10165 := PrimValue(sym_dproperty_1vector_d)

tmp10166 := Call(__e, PrimFunc(symget), V1254, symshen_4dynamic, tmp10165)


__e.TailApply(tmp10157, tmp10166)
return


}, 1)

tmp10167 := Call(__e, PrimFunc(symfn), W1259)


tmp10168 := PrimCons(W1259, V1257)

tmp10169 := PrimCons(tmp10167, tmp10168)

__e.TailApply(tmp10156, tmp10169)
return


}, 1)

tmp10170 := PrimCons(W1259, Nil)

tmp10171 := PrimCons(symdefprolog, tmp10170)

tmp10172 := PrimCons(sym_5_1_1, V1256)

tmp10173 := Call(__e, PrimFunc(symappend), V1255, tmp10172)


tmp10174 := Call(__e, PrimFunc(symappend), tmp10171, tmp10173)


tmp10175 := Call(__e, PrimFunc(symeval), tmp10174)


__e.TailApply(tmp10155, tmp10175)
return


}, 1)

tmp10176 := Call(__e, PrimFunc(symgensym), symshen_4g)


__e.TailApply(tmp10154, tmp10176)
return


}, 5)

tmp10177 := Call(__e, ns2_1set, symshen_4insert_1info, tmp10153)


_ = tmp10177

tmp10178 := MakeNative(func(__e *ControlFlow) {
tmp10179 := MakeNative(func(__e *ControlFlow) {
W1264 := __e.Get(1)
_ = W1264
tmp10180 := MakeNative(func(__e *ControlFlow) {
W1265 := __e.Get(1)
_ = W1265
__e.Return(W1265)
return
}, 1)

tmp10186 := Call(__e, PrimFunc(symempty_2), W1264)


var ifres10181 Obj

if True == tmp10186 {
tmp10182 := Call(__e, PrimFunc(symgensym), symshen_4g)


ifres10181 = tmp10182


} else {
tmp10183 := PrimTail(W1264)

tmp10184 := PrimSet(symshen_4_dnames_d, tmp10183)

_ = tmp10184

tmp10185 := PrimHead(W1264)

ifres10181 = tmp10185


}

__e.TailApply(tmp10180, ifres10181)
return


}, 1)

tmp10187 := PrimValue(symshen_4_dnames_d)

__e.TailApply(tmp10179, tmp10187)
return


}, 0)

tmp10188 := Call(__e, ns2_1set, symshen_4newname, tmp10178)


_ = tmp10188

tmp10189 := MakeNative(func(__e *ControlFlow) {
V1266 := __e.Get(1)
_ = V1266
V1267 := __e.Get(2)
_ = V1267
V1268 := __e.Get(3)
_ = V1268
V1269 := __e.Get(4)
_ = V1269
V1270 := __e.Get(5)
_ = V1270
V1271 := __e.Get(6)
_ = V1271
tmp10190 := MakeNative(func(__e *ControlFlow) {
W1272 := __e.Get(1)
_ = W1272
tmp10201 := PrimEqual(W1272, False)

if True == tmp10201 {
tmp10199 := Call(__e, PrimFunc(symshen_4unlocked_2), V1269)


if True == tmp10199 {
tmp10191 := MakeNative(func(__e *ControlFlow) {
W1276 := __e.Get(1)
_ = W1276
tmp10196 := PrimIsPair(W1276)

if True == tmp10196 {
tmp10192 := MakeNative(func(__e *ControlFlow) {
W1277 := __e.Get(1)
_ = W1277
tmp10193 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp10193

__e.TailApply(PrimFunc(symshen_4call_1dynamic), V1266, W1277, V1268, V1269, V1270, V1271)
return


}, 1)

tmp10194 := PrimTail(W1276)

__e.TailApply(tmp10192, tmp10194)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp10197 := Call(__e, PrimFunc(symshen_4lazyderef), V1267, V1268)


__e.TailApply(tmp10191, tmp10197)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W1272)
return
}


}, 1)

tmp10216 := Call(__e, PrimFunc(symshen_4unlocked_2), V1269)


var ifres10202 Obj

if True == tmp10216 {
tmp10203 := MakeNative(func(__e *ControlFlow) {
W1273 := __e.Get(1)
_ = W1273
tmp10213 := PrimIsPair(W1273)

if True == tmp10213 {
tmp10204 := MakeNative(func(__e *ControlFlow) {
W1274 := __e.Get(1)
_ = W1274
tmp10209 := PrimIsPair(W1274)

if True == tmp10209 {
tmp10205 := MakeNative(func(__e *ControlFlow) {
W1275 := __e.Get(1)
_ = W1275
tmp10206 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp10206

__e.TailApply(PrimFunc(symshen_4callrec), W1275, V1266, V1268, V1269, V1270, V1271)
return


}, 1)

tmp10207 := PrimHead(W1274)

__e.TailApply(tmp10205, tmp10207)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp10210 := PrimHead(W1273)

tmp10211 := Call(__e, PrimFunc(symshen_4lazyderef), tmp10210, V1268)


__e.TailApply(tmp10204, tmp10211)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp10214 := Call(__e, PrimFunc(symshen_4lazyderef), V1267, V1268)


tmp10215 := Call(__e, tmp10203, tmp10214)


ifres10202 = tmp10215


} else {
ifres10202 = False


}

__e.TailApply(tmp10190, ifres10202)
return


}, 6)

tmp10217 := Call(__e, ns2_1set, symshen_4call_1dynamic, tmp10189)


_ = tmp10217

tmp10218 := MakeNative(func(__e *ControlFlow) {
V1278 := __e.Get(1)
_ = V1278
V1279 := __e.Get(2)
_ = V1279
V1280 := __e.Get(3)
_ = V1280
V1281 := __e.Get(4)
_ = V1281
V1282 := __e.Get(5)
_ = V1282
V1283 := __e.Get(6)
_ = V1283
tmp10228 := PrimEqual(Nil, V1279)

if True == tmp10228 {
tmp10219 := Call(__e, V1278, V1280)


tmp10220 := Call(__e, tmp10219, V1281)


tmp10221 := Call(__e, tmp10220, V1282)


__e.TailApply(tmp10221, V1283)
return


} else {
tmp10226 := PrimIsPair(V1279)

if True == tmp10226 {
tmp10222 := PrimHead(V1279)

tmp10223 := Call(__e, V1278, tmp10222)


tmp10224 := PrimTail(V1279)

__e.TailApply(PrimFunc(symshen_4callrec), tmp10223, tmp10224, V1280, V1281, V1282, V1283)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.callrec")))
return
}


}


}, 6)

tmp10229 := Call(__e, ns2_1set, symshen_4callrec, tmp10218)


_ = tmp10229

tmp10230 := MakeNative(func(__e *ControlFlow) {
V1284 := __e.Get(1)
_ = V1284
tmp10249 := PrimIsPair(V1284)

var ifres10240 Obj

if True == tmp10249 {
tmp10247 := PrimTail(V1284)

tmp10248 := PrimIsPair(tmp10247)

var ifres10242 Obj

if True == tmp10248 {
tmp10244 := PrimTail(V1284)

tmp10245 := PrimHead(tmp10244)

tmp10246 := PrimEqual(sym_5_1_1, tmp10245)

var ifres10243 Obj

if True == tmp10246 {
ifres10243 = True


} else {
ifres10243 = False


}

ifres10242 = ifres10243


} else {
ifres10242 = False


}

var ifres10241 Obj

if True == ifres10242 {
ifres10241 = True


} else {
ifres10241 = False


}

ifres10240 = ifres10241


} else {
ifres10240 = False


}

if True == ifres10240 {
tmp10231 := MakeNative(func(__e *ControlFlow) {
W1285 := __e.Get(1)
_ = W1285
tmp10232 := MakeNative(func(__e *ControlFlow) {
W1286 := __e.Get(1)
_ = W1286
tmp10233 := Call(__e, PrimFunc(symshen_4retract_1clause), V1284, W1286)


tmp10234 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), W1285, symshen_4dynamic, tmp10233, tmp10234)
return


}, 1)

tmp10235 := PrimValue(sym_dproperty_1vector_d)

tmp10236 := Call(__e, PrimFunc(symget), W1285, symshen_4dynamic, tmp10235)


__e.TailApply(tmp10232, tmp10236)
return


}, 1)

tmp10237 := PrimHead(V1284)

tmp10238 := Call(__e, PrimFunc(symshen_4predicate), tmp10237)


__e.TailApply(tmp10231, tmp10238)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function retract")))
return
}


}, 1)

tmp10250 := Call(__e, ns2_1set, symretract, tmp10230)


_ = tmp10250

tmp10251 := MakeNative(func(__e *ControlFlow) {
V1292 := __e.Get(1)
_ = V1292
V1293 := __e.Get(2)
_ = V1293
tmp10281 := PrimEqual(Nil, V1293)

if True == tmp10281 {
__e.Return(Nil)
return
} else {
tmp10279 := PrimIsPair(V1293)

var ifres10264 Obj

if True == tmp10279 {
tmp10277 := PrimHead(V1293)

tmp10278 := PrimIsPair(tmp10277)

var ifres10266 Obj

if True == tmp10278 {
tmp10274 := PrimHead(V1293)

tmp10275 := PrimTail(tmp10274)

tmp10276 := PrimIsPair(tmp10275)

var ifres10268 Obj

if True == tmp10276 {
tmp10270 := PrimHead(V1293)

tmp10271 := PrimTail(tmp10270)

tmp10272 := PrimTail(tmp10271)

tmp10273 := PrimEqual(V1292, tmp10272)

var ifres10269 Obj

if True == tmp10273 {
ifres10269 = True


} else {
ifres10269 = False


}

ifres10268 = ifres10269


} else {
ifres10268 = False


}

var ifres10267 Obj

if True == ifres10268 {
ifres10267 = True


} else {
ifres10267 = False


}

ifres10266 = ifres10267


} else {
ifres10266 = False


}

var ifres10265 Obj

if True == ifres10266 {
ifres10265 = True


} else {
ifres10265 = False


}

ifres10264 = ifres10265


} else {
ifres10264 = False


}

if True == ifres10264 {
tmp10252 := PrimHead(V1293)

tmp10253 := PrimTail(tmp10252)

tmp10254 := PrimHead(tmp10253)

tmp10255 := PrimValue(symshen_4_dnames_d)

tmp10256 := PrimCons(tmp10254, tmp10255)

tmp10257 := PrimSet(symshen_4_dnames_d, tmp10256)

_ = tmp10257

__e.Return(PrimTail(V1293))
return


} else {
tmp10262 := PrimIsPair(V1293)

if True == tmp10262 {
tmp10258 := PrimHead(V1293)

tmp10259 := PrimTail(V1293)

tmp10260 := Call(__e, PrimFunc(symshen_4retract_1clause), V1292, tmp10259)


__e.Return(PrimCons(tmp10258, tmp10260))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.retract-clause")))
return
}


}


}


}, 2)

tmp10282 := Call(__e, ns2_1set, symshen_4retract_1clause, tmp10251)


_ = tmp10282

tmp10283 := MakeNative(func(__e *ControlFlow) {
V1294 := __e.Get(1)
_ = V1294
V1295 := __e.Get(2)
_ = V1295
tmp10284 := MakeNative(func(__e *ControlFlow) {
Z1296 := __e.Get(1)
_ = Z1296
__e.TailApply(PrimFunc(symshen_4_5defprolog_6), Z1296)
return
}, 1)

tmp10285 := PrimCons(V1294, V1295)

__e.TailApply(PrimFunc(symcompile), tmp10284, tmp10285)
return


}, 2)

tmp10286 := Call(__e, ns2_1set, symshen_4compile_1prolog, tmp10283)


_ = tmp10286

tmp10287 := MakeNative(func(__e *ControlFlow) {
V1297 := __e.Get(1)
_ = V1297
tmp10288 := MakeNative(func(__e *ControlFlow) {
W1298 := __e.Get(1)
_ = W1298
tmp10290 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1298)


if True == tmp10290 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1298)
return
}


}, 1)

tmp10312 := PrimIsPair(V1297)

var ifres10291 Obj

if True == tmp10312 {
tmp10292 := MakeNative(func(__e *ControlFlow) {
W1299 := __e.Get(1)
_ = W1299
tmp10293 := MakeNative(func(__e *ControlFlow) {
W1300 := __e.Get(1)
_ = W1300
tmp10294 := MakeNative(func(__e *ControlFlow) {
W1301 := __e.Get(1)
_ = W1301
tmp10306 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1301)


if True == tmp10306 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10295 := MakeNative(func(__e *ControlFlow) {
W1302 := __e.Get(1)
_ = W1302
tmp10296 := MakeNative(func(__e *ControlFlow) {
W1303 := __e.Get(1)
_ = W1303
tmp10297 := MakeNative(func(__e *ControlFlow) {
W1304 := __e.Get(1)
_ = W1304
tmp10298 := MakeNative(func(__e *ControlFlow) {
W1305 := __e.Get(1)
_ = W1305
__e.TailApply(PrimFunc(symshen_4horn_1clause_1procedure), W1299, W1305)
return
}, 1)

tmp10299 := MakeNative(func(__e *ControlFlow) {
Z1306 := __e.Get(1)
_ = Z1306
__e.TailApply(PrimFunc(symshen_4linearise_1clause), Z1306)
return
}, 1)

tmp10300 := Call(__e, PrimFunc(symmap), tmp10299, W1302)


__e.TailApply(tmp10298, tmp10300)
return


}, 1)

tmp10301 := Call(__e, PrimFunc(symshen_4prolog_1arity_1check), W1299, W1302)


tmp10302 := Call(__e, tmp10297, tmp10301)


__e.TailApply(PrimFunc(symshen_4comb), W1303, tmp10302)
return


}, 1)

tmp10303 := Call(__e, PrimFunc(symshen_4in_1_6), W1301)


__e.TailApply(tmp10296, tmp10303)
return


}, 1)

tmp10304 := Call(__e, PrimFunc(symshen_4_5_1out), W1301)


__e.TailApply(tmp10295, tmp10304)
return


}


}, 1)

tmp10307 := Call(__e, PrimFunc(symshen_4_5clauses_6), W1300)


__e.TailApply(tmp10294, tmp10307)
return


}, 1)

tmp10308 := Call(__e, PrimFunc(symtail), V1297)


__e.TailApply(tmp10293, tmp10308)
return


}, 1)

tmp10309 := Call(__e, PrimFunc(symhead), V1297)


tmp10310 := Call(__e, tmp10292, tmp10309)


ifres10291 = tmp10310


} else {
tmp10311 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10291 = tmp10311


}

__e.TailApply(tmp10288, ifres10291)
return


}, 1)

tmp10313 := Call(__e, ns2_1set, symshen_4_5defprolog_6, tmp10287)


_ = tmp10313

tmp10314 := MakeNative(func(__e *ControlFlow) {
V1309 := __e.Get(1)
_ = V1309
V1310 := __e.Get(2)
_ = V1310
tmp10358 := PrimIsPair(V1310)

var ifres10339 Obj

if True == tmp10358 {
tmp10356 := PrimHead(V1310)

tmp10357 := PrimIsPair(tmp10356)

var ifres10341 Obj

if True == tmp10357 {
tmp10353 := PrimHead(V1310)

tmp10354 := PrimTail(tmp10353)

tmp10355 := PrimIsPair(tmp10354)

var ifres10343 Obj

if True == tmp10355 {
tmp10349 := PrimHead(V1310)

tmp10350 := PrimTail(tmp10349)

tmp10351 := PrimTail(tmp10350)

tmp10352 := PrimEqual(Nil, tmp10351)

var ifres10345 Obj

if True == tmp10352 {
tmp10347 := PrimTail(V1310)

tmp10348 := PrimEqual(Nil, tmp10347)

var ifres10346 Obj

if True == tmp10348 {
ifres10346 = True


} else {
ifres10346 = False


}

ifres10345 = ifres10346


} else {
ifres10345 = False


}

var ifres10344 Obj

if True == ifres10345 {
ifres10344 = True


} else {
ifres10344 = False


}

ifres10343 = ifres10344


} else {
ifres10343 = False


}

var ifres10342 Obj

if True == ifres10343 {
ifres10342 = True


} else {
ifres10342 = False


}

ifres10341 = ifres10342


} else {
ifres10341 = False


}

var ifres10340 Obj

if True == ifres10341 {
ifres10340 = True


} else {
ifres10340 = False


}

ifres10339 = ifres10340


} else {
ifres10339 = False


}

if True == ifres10339 {
tmp10315 := PrimHead(V1310)

tmp10316 := PrimHead(tmp10315)

__e.TailApply(PrimFunc(symlength), tmp10316)
return


} else {
tmp10337 := PrimIsPair(V1310)

var ifres10322 Obj

if True == tmp10337 {
tmp10335 := PrimHead(V1310)

tmp10336 := PrimIsPair(tmp10335)

var ifres10324 Obj

if True == tmp10336 {
tmp10332 := PrimHead(V1310)

tmp10333 := PrimTail(tmp10332)

tmp10334 := PrimIsPair(tmp10333)

var ifres10326 Obj

if True == tmp10334 {
tmp10328 := PrimHead(V1310)

tmp10329 := PrimTail(tmp10328)

tmp10330 := PrimTail(tmp10329)

tmp10331 := PrimEqual(Nil, tmp10330)

var ifres10327 Obj

if True == tmp10331 {
ifres10327 = True


} else {
ifres10327 = False


}

ifres10326 = ifres10327


} else {
ifres10326 = False


}

var ifres10325 Obj

if True == ifres10326 {
ifres10325 = True


} else {
ifres10325 = False


}

ifres10324 = ifres10325


} else {
ifres10324 = False


}

var ifres10323 Obj

if True == ifres10324 {
ifres10323 = True


} else {
ifres10323 = False


}

ifres10322 = ifres10323


} else {
ifres10322 = False


}

if True == ifres10322 {
tmp10317 := PrimHead(V1310)

tmp10318 := PrimHead(tmp10317)

tmp10319 := Call(__e, PrimFunc(symlength), tmp10318)


tmp10320 := PrimTail(V1310)

__e.TailApply(PrimFunc(symshen_4pac_1h), V1309, tmp10319, tmp10320)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.prolog-arity-check")))
return
}


}


}, 2)

tmp10359 := Call(__e, ns2_1set, symshen_4prolog_1arity_1check, tmp10314)


_ = tmp10359

tmp10360 := MakeNative(func(__e *ControlFlow) {
V1315 := __e.Get(1)
_ = V1315
V1316 := __e.Get(2)
_ = V1316
V1317 := __e.Get(3)
_ = V1317
tmp10376 := PrimEqual(Nil, V1317)

if True == tmp10376 {
__e.Return(V1316)
return
} else {
tmp10374 := PrimIsPair(V1317)

var ifres10370 Obj

if True == tmp10374 {
tmp10372 := PrimHead(V1317)

tmp10373 := PrimIsPair(tmp10372)

var ifres10371 Obj

if True == tmp10373 {
ifres10371 = True


} else {
ifres10371 = False


}

ifres10370 = ifres10371


} else {
ifres10370 = False


}

if True == ifres10370 {
tmp10365 := PrimHead(V1317)

tmp10366 := PrimHead(tmp10365)

tmp10367 := Call(__e, PrimFunc(symlength), tmp10366)


tmp10368 := PrimEqual(V1316, tmp10367)

if True == tmp10368 {
tmp10361 := PrimTail(V1317)

__e.TailApply(PrimFunc(symshen_4pac_1h), V1315, V1316, tmp10361)
return


} else {
tmp10362 := Call(__e, PrimFunc(symshen_4app), V1315, MakeString("\n"), symshen_4a)


tmp10363 := PrimStringConcat(MakeString("arity error in prolog procedure "), tmp10362)

__e.Return(PrimSimpleError(tmp10363))
return


}


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.pac-h")))
return
}


}


}, 3)

tmp10377 := Call(__e, ns2_1set, symshen_4pac_1h, tmp10360)


_ = tmp10377

tmp10378 := MakeNative(func(__e *ControlFlow) {
V1318 := __e.Get(1)
_ = V1318
tmp10379 := MakeNative(func(__e *ControlFlow) {
W1319 := __e.Get(1)
_ = W1319
tmp10398 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1319)


if True == tmp10398 {
tmp10380 := MakeNative(func(__e *ControlFlow) {
W1326 := __e.Get(1)
_ = W1326
tmp10382 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1326)


if True == tmp10382 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1326)
return
}


}, 1)

tmp10383 := MakeNative(func(__e *ControlFlow) {
W1327 := __e.Get(1)
_ = W1327
tmp10394 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1327)


if True == tmp10394 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10384 := MakeNative(func(__e *ControlFlow) {
W1328 := __e.Get(1)
_ = W1328
tmp10385 := MakeNative(func(__e *ControlFlow) {
W1329 := __e.Get(1)
_ = W1329
tmp10390 := Call(__e, PrimFunc(symempty_2), W1328)


var ifres10386 Obj

if True == tmp10390 {
ifres10386 = Nil


} else {
tmp10387 := Call(__e, PrimFunc(symshen_4app), W1328, MakeString("\n ..."), symshen_4r)


tmp10388 := PrimStringConcat(MakeString("Prolog syntax error here:\n "), tmp10387)

tmp10389 := PrimSimpleError(tmp10388)

ifres10386 = tmp10389


}

__e.TailApply(PrimFunc(symshen_4comb), W1329, ifres10386)
return


}, 1)

tmp10391 := Call(__e, PrimFunc(symshen_4in_1_6), W1327)


__e.TailApply(tmp10385, tmp10391)
return


}, 1)

tmp10392 := Call(__e, PrimFunc(symshen_4_5_1out), W1327)


__e.TailApply(tmp10384, tmp10392)
return


}


}, 1)

tmp10395 := Call(__e, PrimFunc(sym_5_b_6), V1318)


tmp10396 := Call(__e, tmp10383, tmp10395)


__e.TailApply(tmp10380, tmp10396)
return


} else {
__e.Return(W1319)
return
}


}, 1)

tmp10399 := MakeNative(func(__e *ControlFlow) {
W1320 := __e.Get(1)
_ = W1320
tmp10414 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1320)


if True == tmp10414 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10400 := MakeNative(func(__e *ControlFlow) {
W1321 := __e.Get(1)
_ = W1321
tmp10401 := MakeNative(func(__e *ControlFlow) {
W1322 := __e.Get(1)
_ = W1322
tmp10402 := MakeNative(func(__e *ControlFlow) {
W1323 := __e.Get(1)
_ = W1323
tmp10409 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1323)


if True == tmp10409 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10403 := MakeNative(func(__e *ControlFlow) {
W1324 := __e.Get(1)
_ = W1324
tmp10404 := MakeNative(func(__e *ControlFlow) {
W1325 := __e.Get(1)
_ = W1325
tmp10405 := PrimCons(W1321, W1324)

__e.TailApply(PrimFunc(symshen_4comb), W1325, tmp10405)
return


}, 1)

tmp10406 := Call(__e, PrimFunc(symshen_4in_1_6), W1323)


__e.TailApply(tmp10404, tmp10406)
return


}, 1)

tmp10407 := Call(__e, PrimFunc(symshen_4_5_1out), W1323)


__e.TailApply(tmp10403, tmp10407)
return


}


}, 1)

tmp10410 := Call(__e, PrimFunc(symshen_4_5clauses_6), W1322)


__e.TailApply(tmp10402, tmp10410)
return


}, 1)

tmp10411 := Call(__e, PrimFunc(symshen_4in_1_6), W1320)


__e.TailApply(tmp10401, tmp10411)
return


}, 1)

tmp10412 := Call(__e, PrimFunc(symshen_4_5_1out), W1320)


__e.TailApply(tmp10400, tmp10412)
return


}


}, 1)

tmp10415 := Call(__e, PrimFunc(symshen_4_5clause_6), V1318)


tmp10416 := Call(__e, tmp10399, tmp10415)


__e.TailApply(tmp10379, tmp10416)
return


}, 1)

tmp10417 := Call(__e, ns2_1set, symshen_4_5clauses_6, tmp10378)


_ = tmp10417

tmp10418 := MakeNative(func(__e *ControlFlow) {
V1330 := __e.Get(1)
_ = V1330
tmp10434 := PrimIsPair(V1330)

var ifres10425 Obj

if True == tmp10434 {
tmp10432 := PrimTail(V1330)

tmp10433 := PrimIsPair(tmp10432)

var ifres10427 Obj

if True == tmp10433 {
tmp10429 := PrimTail(V1330)

tmp10430 := PrimTail(tmp10429)

tmp10431 := PrimEqual(Nil, tmp10430)

var ifres10428 Obj

if True == tmp10431 {
ifres10428 = True


} else {
ifres10428 = False


}

ifres10427 = ifres10428


} else {
ifres10427 = False


}

var ifres10426 Obj

if True == ifres10427 {
ifres10426 = True


} else {
ifres10426 = False


}

ifres10425 = ifres10426


} else {
ifres10425 = False


}

if True == ifres10425 {
tmp10419 := PrimHead(V1330)

tmp10420 := PrimTail(V1330)

tmp10421 := PrimHead(tmp10420)

tmp10422 := Call(__e, PrimFunc(sym_8p), tmp10419, tmp10421)


tmp10423 := Call(__e, PrimFunc(symshen_4linearise), tmp10422)


__e.TailApply(PrimFunc(symshen_4lch), tmp10423)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.linearise-clause")))
return
}


}, 1)

tmp10435 := Call(__e, ns2_1set, symshen_4linearise_1clause, tmp10418)


_ = tmp10435

tmp10436 := MakeNative(func(__e *ControlFlow) {
V1331 := __e.Get(1)
_ = V1331
tmp10442 := Call(__e, PrimFunc(symtuple_2), V1331)


if True == tmp10442 {
tmp10437 := Call(__e, PrimFunc(symfst), V1331)


tmp10438 := Call(__e, PrimFunc(symsnd), V1331)


tmp10439 := Call(__e, PrimFunc(symshen_4lchh), tmp10438)


tmp10440 := PrimCons(tmp10439, Nil)

__e.Return(PrimCons(tmp10437, tmp10440))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.lch")))
return
}


}, 1)

tmp10443 := Call(__e, ns2_1set, symshen_4lch, tmp10436)


_ = tmp10443

tmp10444 := MakeNative(func(__e *ControlFlow) {
V1332 := __e.Get(1)
_ = V1332
tmp10507 := PrimIsPair(V1332)

var ifres10456 Obj

if True == tmp10507 {
tmp10505 := PrimHead(V1332)

tmp10506 := PrimEqual(symwhere, tmp10505)

var ifres10458 Obj

if True == tmp10506 {
tmp10503 := PrimTail(V1332)

tmp10504 := PrimIsPair(tmp10503)

var ifres10460 Obj

if True == tmp10504 {
tmp10500 := PrimTail(V1332)

tmp10501 := PrimHead(tmp10500)

tmp10502 := PrimIsPair(tmp10501)

var ifres10462 Obj

if True == tmp10502 {
tmp10496 := PrimTail(V1332)

tmp10497 := PrimHead(tmp10496)

tmp10498 := PrimHead(tmp10497)

tmp10499 := PrimEqual(sym_a, tmp10498)

var ifres10464 Obj

if True == tmp10499 {
tmp10492 := PrimTail(V1332)

tmp10493 := PrimHead(tmp10492)

tmp10494 := PrimTail(tmp10493)

tmp10495 := PrimIsPair(tmp10494)

var ifres10466 Obj

if True == tmp10495 {
tmp10487 := PrimTail(V1332)

tmp10488 := PrimHead(tmp10487)

tmp10489 := PrimTail(tmp10488)

tmp10490 := PrimTail(tmp10489)

tmp10491 := PrimIsPair(tmp10490)

var ifres10468 Obj

if True == tmp10491 {
tmp10481 := PrimTail(V1332)

tmp10482 := PrimHead(tmp10481)

tmp10483 := PrimTail(tmp10482)

tmp10484 := PrimTail(tmp10483)

tmp10485 := PrimTail(tmp10484)

tmp10486 := PrimEqual(Nil, tmp10485)

var ifres10470 Obj

if True == tmp10486 {
tmp10478 := PrimTail(V1332)

tmp10479 := PrimTail(tmp10478)

tmp10480 := PrimIsPair(tmp10479)

var ifres10472 Obj

if True == tmp10480 {
tmp10474 := PrimTail(V1332)

tmp10475 := PrimTail(tmp10474)

tmp10476 := PrimTail(tmp10475)

tmp10477 := PrimEqual(Nil, tmp10476)

var ifres10473 Obj

if True == tmp10477 {
ifres10473 = True


} else {
ifres10473 = False


}

ifres10472 = ifres10473


} else {
ifres10472 = False


}

var ifres10471 Obj

if True == ifres10472 {
ifres10471 = True


} else {
ifres10471 = False


}

ifres10470 = ifres10471


} else {
ifres10470 = False


}

var ifres10469 Obj

if True == ifres10470 {
ifres10469 = True


} else {
ifres10469 = False


}

ifres10468 = ifres10469


} else {
ifres10468 = False


}

var ifres10467 Obj

if True == ifres10468 {
ifres10467 = True


} else {
ifres10467 = False


}

ifres10466 = ifres10467


} else {
ifres10466 = False


}

var ifres10465 Obj

if True == ifres10466 {
ifres10465 = True


} else {
ifres10465 = False


}

ifres10464 = ifres10465


} else {
ifres10464 = False


}

var ifres10463 Obj

if True == ifres10464 {
ifres10463 = True


} else {
ifres10463 = False


}

ifres10462 = ifres10463


} else {
ifres10462 = False


}

var ifres10461 Obj

if True == ifres10462 {
ifres10461 = True


} else {
ifres10461 = False


}

ifres10460 = ifres10461


} else {
ifres10460 = False


}

var ifres10459 Obj

if True == ifres10460 {
ifres10459 = True


} else {
ifres10459 = False


}

ifres10458 = ifres10459


} else {
ifres10458 = False


}

var ifres10457 Obj

if True == ifres10458 {
ifres10457 = True


} else {
ifres10457 = False


}

ifres10456 = ifres10457


} else {
ifres10456 = False


}

if True == ifres10456 {
tmp10446 := PrimValue(symshen_4_doccurs_d)

var ifres10445 Obj

if True == tmp10446 {
ifres10445 = symis_b


} else {
ifres10445 = symis


}

tmp10447 := PrimTail(V1332)

tmp10448 := PrimHead(tmp10447)

tmp10449 := PrimTail(tmp10448)

tmp10450 := PrimCons(ifres10445, tmp10449)

tmp10451 := PrimTail(V1332)

tmp10452 := PrimTail(tmp10451)

tmp10453 := PrimHead(tmp10452)

tmp10454 := Call(__e, PrimFunc(symshen_4lchh), tmp10453)


__e.Return(PrimCons(tmp10450, tmp10454))
return


} else {
__e.Return(V1332)
return
}


}, 1)

tmp10508 := Call(__e, ns2_1set, symshen_4lchh, tmp10444)


_ = tmp10508

tmp10509 := MakeNative(func(__e *ControlFlow) {
V1333 := __e.Get(1)
_ = V1333
tmp10510 := MakeNative(func(__e *ControlFlow) {
W1334 := __e.Get(1)
_ = W1334
tmp10512 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1334)


if True == tmp10512 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1334)
return
}


}, 1)

tmp10513 := MakeNative(func(__e *ControlFlow) {
W1335 := __e.Get(1)
_ = W1335
tmp10539 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1335)


if True == tmp10539 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10514 := MakeNative(func(__e *ControlFlow) {
W1336 := __e.Get(1)
_ = W1336
tmp10515 := MakeNative(func(__e *ControlFlow) {
W1337 := __e.Get(1)
_ = W1337
tmp10535 := Call(__e, PrimFunc(symshen_4hds_a_2), W1337, sym_5_1_1)


if True == tmp10535 {
tmp10516 := MakeNative(func(__e *ControlFlow) {
W1338 := __e.Get(1)
_ = W1338
tmp10517 := MakeNative(func(__e *ControlFlow) {
W1339 := __e.Get(1)
_ = W1339
tmp10531 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1339)


if True == tmp10531 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10518 := MakeNative(func(__e *ControlFlow) {
W1340 := __e.Get(1)
_ = W1340
tmp10519 := MakeNative(func(__e *ControlFlow) {
W1341 := __e.Get(1)
_ = W1341
tmp10520 := MakeNative(func(__e *ControlFlow) {
W1342 := __e.Get(1)
_ = W1342
tmp10526 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1342)


if True == tmp10526 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10521 := MakeNative(func(__e *ControlFlow) {
W1343 := __e.Get(1)
_ = W1343
tmp10522 := PrimCons(W1340, Nil)

tmp10523 := PrimCons(W1336, tmp10522)

__e.TailApply(PrimFunc(symshen_4comb), W1343, tmp10523)
return


}, 1)

tmp10524 := Call(__e, PrimFunc(symshen_4in_1_6), W1342)


__e.TailApply(tmp10521, tmp10524)
return


}


}, 1)

tmp10527 := Call(__e, PrimFunc(symshen_4_5sc_6), W1341)


__e.TailApply(tmp10520, tmp10527)
return


}, 1)

tmp10528 := Call(__e, PrimFunc(symshen_4in_1_6), W1339)


__e.TailApply(tmp10519, tmp10528)
return


}, 1)

tmp10529 := Call(__e, PrimFunc(symshen_4_5_1out), W1339)


__e.TailApply(tmp10518, tmp10529)
return


}


}, 1)

tmp10532 := Call(__e, PrimFunc(symshen_4_5body_6), W1338)


__e.TailApply(tmp10517, tmp10532)
return


}, 1)

tmp10533 := Call(__e, PrimFunc(symtail), W1337)


__e.TailApply(tmp10516, tmp10533)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10536 := Call(__e, PrimFunc(symshen_4in_1_6), W1335)


__e.TailApply(tmp10515, tmp10536)
return


}, 1)

tmp10537 := Call(__e, PrimFunc(symshen_4_5_1out), W1335)


__e.TailApply(tmp10514, tmp10537)
return


}


}, 1)

tmp10540 := Call(__e, PrimFunc(symshen_4_5head_6), V1333)


tmp10541 := Call(__e, tmp10513, tmp10540)


__e.TailApply(tmp10510, tmp10541)
return


}, 1)

tmp10542 := Call(__e, ns2_1set, symshen_4_5clause_6, tmp10509)


_ = tmp10542

tmp10543 := MakeNative(func(__e *ControlFlow) {
V1344 := __e.Get(1)
_ = V1344
tmp10544 := MakeNative(func(__e *ControlFlow) {
W1345 := __e.Get(1)
_ = W1345
tmp10556 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1345)


if True == tmp10556 {
tmp10545 := MakeNative(func(__e *ControlFlow) {
W1352 := __e.Get(1)
_ = W1352
tmp10547 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1352)


if True == tmp10547 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1352)
return
}


}, 1)

tmp10548 := MakeNative(func(__e *ControlFlow) {
W1353 := __e.Get(1)
_ = W1353
tmp10552 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1353)


if True == tmp10552 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10549 := MakeNative(func(__e *ControlFlow) {
W1354 := __e.Get(1)
_ = W1354
__e.TailApply(PrimFunc(symshen_4comb), W1354, Nil)
return
}, 1)

tmp10550 := Call(__e, PrimFunc(symshen_4in_1_6), W1353)


__e.TailApply(tmp10549, tmp10550)
return


}


}, 1)

tmp10553 := Call(__e, PrimFunc(sym_5e_6), V1344)


tmp10554 := Call(__e, tmp10548, tmp10553)


__e.TailApply(tmp10545, tmp10554)
return


} else {
__e.Return(W1345)
return
}


}, 1)

tmp10557 := MakeNative(func(__e *ControlFlow) {
W1346 := __e.Get(1)
_ = W1346
tmp10572 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1346)


if True == tmp10572 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10558 := MakeNative(func(__e *ControlFlow) {
W1347 := __e.Get(1)
_ = W1347
tmp10559 := MakeNative(func(__e *ControlFlow) {
W1348 := __e.Get(1)
_ = W1348
tmp10560 := MakeNative(func(__e *ControlFlow) {
W1349 := __e.Get(1)
_ = W1349
tmp10567 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1349)


if True == tmp10567 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10561 := MakeNative(func(__e *ControlFlow) {
W1350 := __e.Get(1)
_ = W1350
tmp10562 := MakeNative(func(__e *ControlFlow) {
W1351 := __e.Get(1)
_ = W1351
tmp10563 := PrimCons(W1347, W1350)

__e.TailApply(PrimFunc(symshen_4comb), W1351, tmp10563)
return


}, 1)

tmp10564 := Call(__e, PrimFunc(symshen_4in_1_6), W1349)


__e.TailApply(tmp10562, tmp10564)
return


}, 1)

tmp10565 := Call(__e, PrimFunc(symshen_4_5_1out), W1349)


__e.TailApply(tmp10561, tmp10565)
return


}


}, 1)

tmp10568 := Call(__e, PrimFunc(symshen_4_5head_6), W1348)


__e.TailApply(tmp10560, tmp10568)
return


}, 1)

tmp10569 := Call(__e, PrimFunc(symshen_4in_1_6), W1346)


__e.TailApply(tmp10559, tmp10569)
return


}, 1)

tmp10570 := Call(__e, PrimFunc(symshen_4_5_1out), W1346)


__e.TailApply(tmp10558, tmp10570)
return


}


}, 1)

tmp10573 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1344)


tmp10574 := Call(__e, tmp10557, tmp10573)


__e.TailApply(tmp10544, tmp10574)
return


}, 1)

tmp10575 := Call(__e, ns2_1set, symshen_4_5head_6, tmp10543)


_ = tmp10575

tmp10576 := MakeNative(func(__e *ControlFlow) {
V1355 := __e.Get(1)
_ = V1355
tmp10577 := MakeNative(func(__e *ControlFlow) {
W1356 := __e.Get(1)
_ = W1356
tmp10765 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1356)


if True == tmp10765 {
tmp10578 := MakeNative(func(__e *ControlFlow) {
W1359 := __e.Get(1)
_ = W1359
tmp10752 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1359)


if True == tmp10752 {
tmp10579 := MakeNative(func(__e *ControlFlow) {
W1362 := __e.Get(1)
_ = W1362
tmp10713 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1362)


if True == tmp10713 {
tmp10580 := MakeNative(func(__e *ControlFlow) {
W1374 := __e.Get(1)
_ = W1374
tmp10683 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1374)


if True == tmp10683 {
tmp10581 := MakeNative(func(__e *ControlFlow) {
W1383 := __e.Get(1)
_ = W1383
tmp10653 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1383)


if True == tmp10653 {
tmp10582 := MakeNative(func(__e *ControlFlow) {
W1392 := __e.Get(1)
_ = W1392
tmp10619 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1392)


if True == tmp10619 {
tmp10583 := MakeNative(func(__e *ControlFlow) {
W1402 := __e.Get(1)
_ = W1402
tmp10585 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1402)


if True == tmp10585 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1402)
return
}


}, 1)

tmp10617 := Call(__e, PrimFunc(symshen_4ccons_2), V1355)


var ifres10586 Obj

if True == tmp10617 {
tmp10587 := MakeNative(func(__e *ControlFlow) {
W1403 := __e.Get(1)
_ = W1403
tmp10588 := MakeNative(func(__e *ControlFlow) {
W1404 := __e.Get(1)
_ = W1404
tmp10612 := Call(__e, PrimFunc(symshen_4hds_a_2), W1403, symmode)


if True == tmp10612 {
tmp10589 := MakeNative(func(__e *ControlFlow) {
W1405 := __e.Get(1)
_ = W1405
tmp10590 := MakeNative(func(__e *ControlFlow) {
W1406 := __e.Get(1)
_ = W1406
tmp10608 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1406)


if True == tmp10608 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10591 := MakeNative(func(__e *ControlFlow) {
W1407 := __e.Get(1)
_ = W1407
tmp10592 := MakeNative(func(__e *ControlFlow) {
W1408 := __e.Get(1)
_ = W1408
tmp10604 := Call(__e, PrimFunc(symshen_4hds_a_2), W1408, sym_1)


if True == tmp10604 {
tmp10593 := MakeNative(func(__e *ControlFlow) {
W1409 := __e.Get(1)
_ = W1409
tmp10594 := MakeNative(func(__e *ControlFlow) {
W1410 := __e.Get(1)
_ = W1410
tmp10600 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1410)


if True == tmp10600 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10595 := MakeNative(func(__e *ControlFlow) {
W1411 := __e.Get(1)
_ = W1411
tmp10596 := PrimCons(W1407, Nil)

tmp10597 := PrimCons(symshen_4_1m, tmp10596)

__e.TailApply(PrimFunc(symshen_4comb), W1404, tmp10597)
return


}, 1)

tmp10598 := Call(__e, PrimFunc(symshen_4in_1_6), W1410)


__e.TailApply(tmp10595, tmp10598)
return


}


}, 1)

tmp10601 := Call(__e, PrimFunc(sym_5end_6), W1409)


__e.TailApply(tmp10594, tmp10601)
return


}, 1)

tmp10602 := Call(__e, PrimFunc(symtail), W1408)


__e.TailApply(tmp10593, tmp10602)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10605 := Call(__e, PrimFunc(symshen_4in_1_6), W1406)


__e.TailApply(tmp10592, tmp10605)
return


}, 1)

tmp10606 := Call(__e, PrimFunc(symshen_4_5_1out), W1406)


__e.TailApply(tmp10591, tmp10606)
return


}


}, 1)

tmp10609 := Call(__e, PrimFunc(symshen_4_5hterm_6), W1405)


__e.TailApply(tmp10590, tmp10609)
return


}, 1)

tmp10610 := Call(__e, PrimFunc(symtail), W1403)


__e.TailApply(tmp10589, tmp10610)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10613 := Call(__e, PrimFunc(symtail), V1355)


__e.TailApply(tmp10588, tmp10613)
return


}, 1)

tmp10614 := Call(__e, PrimFunc(symhead), V1355)


tmp10615 := Call(__e, tmp10587, tmp10614)


ifres10586 = tmp10615


} else {
tmp10616 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10586 = tmp10616


}

__e.TailApply(tmp10583, ifres10586)
return


} else {
__e.Return(W1392)
return
}


}, 1)

tmp10651 := Call(__e, PrimFunc(symshen_4ccons_2), V1355)


var ifres10620 Obj

if True == tmp10651 {
tmp10621 := MakeNative(func(__e *ControlFlow) {
W1393 := __e.Get(1)
_ = W1393
tmp10622 := MakeNative(func(__e *ControlFlow) {
W1394 := __e.Get(1)
_ = W1394
tmp10646 := Call(__e, PrimFunc(symshen_4hds_a_2), W1393, symmode)


if True == tmp10646 {
tmp10623 := MakeNative(func(__e *ControlFlow) {
W1395 := __e.Get(1)
_ = W1395
tmp10624 := MakeNative(func(__e *ControlFlow) {
W1396 := __e.Get(1)
_ = W1396
tmp10642 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1396)


if True == tmp10642 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10625 := MakeNative(func(__e *ControlFlow) {
W1397 := __e.Get(1)
_ = W1397
tmp10626 := MakeNative(func(__e *ControlFlow) {
W1398 := __e.Get(1)
_ = W1398
tmp10638 := Call(__e, PrimFunc(symshen_4hds_a_2), W1398, sym_7)


if True == tmp10638 {
tmp10627 := MakeNative(func(__e *ControlFlow) {
W1399 := __e.Get(1)
_ = W1399
tmp10628 := MakeNative(func(__e *ControlFlow) {
W1400 := __e.Get(1)
_ = W1400
tmp10634 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1400)


if True == tmp10634 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10629 := MakeNative(func(__e *ControlFlow) {
W1401 := __e.Get(1)
_ = W1401
tmp10630 := PrimCons(W1397, Nil)

tmp10631 := PrimCons(symshen_4_7m, tmp10630)

__e.TailApply(PrimFunc(symshen_4comb), W1394, tmp10631)
return


}, 1)

tmp10632 := Call(__e, PrimFunc(symshen_4in_1_6), W1400)


__e.TailApply(tmp10629, tmp10632)
return


}


}, 1)

tmp10635 := Call(__e, PrimFunc(sym_5end_6), W1399)


__e.TailApply(tmp10628, tmp10635)
return


}, 1)

tmp10636 := Call(__e, PrimFunc(symtail), W1398)


__e.TailApply(tmp10627, tmp10636)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10639 := Call(__e, PrimFunc(symshen_4in_1_6), W1396)


__e.TailApply(tmp10626, tmp10639)
return


}, 1)

tmp10640 := Call(__e, PrimFunc(symshen_4_5_1out), W1396)


__e.TailApply(tmp10625, tmp10640)
return


}


}, 1)

tmp10643 := Call(__e, PrimFunc(symshen_4_5hterm_6), W1395)


__e.TailApply(tmp10624, tmp10643)
return


}, 1)

tmp10644 := Call(__e, PrimFunc(symtail), W1393)


__e.TailApply(tmp10623, tmp10644)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10647 := Call(__e, PrimFunc(symtail), V1355)


__e.TailApply(tmp10622, tmp10647)
return


}, 1)

tmp10648 := Call(__e, PrimFunc(symhead), V1355)


tmp10649 := Call(__e, tmp10621, tmp10648)


ifres10620 = tmp10649


} else {
tmp10650 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10620 = tmp10650


}

__e.TailApply(tmp10582, ifres10620)
return


} else {
__e.Return(W1383)
return
}


}, 1)

tmp10681 := Call(__e, PrimFunc(symshen_4ccons_2), V1355)


var ifres10654 Obj

if True == tmp10681 {
tmp10655 := MakeNative(func(__e *ControlFlow) {
W1384 := __e.Get(1)
_ = W1384
tmp10656 := MakeNative(func(__e *ControlFlow) {
W1385 := __e.Get(1)
_ = W1385
tmp10676 := Call(__e, PrimFunc(symshen_4hds_a_2), W1384, sym_1)


if True == tmp10676 {
tmp10657 := MakeNative(func(__e *ControlFlow) {
W1386 := __e.Get(1)
_ = W1386
tmp10658 := MakeNative(func(__e *ControlFlow) {
W1387 := __e.Get(1)
_ = W1387
tmp10672 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1387)


if True == tmp10672 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10659 := MakeNative(func(__e *ControlFlow) {
W1388 := __e.Get(1)
_ = W1388
tmp10660 := MakeNative(func(__e *ControlFlow) {
W1389 := __e.Get(1)
_ = W1389
tmp10661 := MakeNative(func(__e *ControlFlow) {
W1390 := __e.Get(1)
_ = W1390
tmp10667 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1390)


if True == tmp10667 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10662 := MakeNative(func(__e *ControlFlow) {
W1391 := __e.Get(1)
_ = W1391
tmp10663 := PrimCons(W1388, Nil)

tmp10664 := PrimCons(symshen_4_1m, tmp10663)

__e.TailApply(PrimFunc(symshen_4comb), W1385, tmp10664)
return


}, 1)

tmp10665 := Call(__e, PrimFunc(symshen_4in_1_6), W1390)


__e.TailApply(tmp10662, tmp10665)
return


}


}, 1)

tmp10668 := Call(__e, PrimFunc(sym_5end_6), W1389)


__e.TailApply(tmp10661, tmp10668)
return


}, 1)

tmp10669 := Call(__e, PrimFunc(symshen_4in_1_6), W1387)


__e.TailApply(tmp10660, tmp10669)
return


}, 1)

tmp10670 := Call(__e, PrimFunc(symshen_4_5_1out), W1387)


__e.TailApply(tmp10659, tmp10670)
return


}


}, 1)

tmp10673 := Call(__e, PrimFunc(symshen_4_5hterm_6), W1386)


__e.TailApply(tmp10658, tmp10673)
return


}, 1)

tmp10674 := Call(__e, PrimFunc(symtail), W1384)


__e.TailApply(tmp10657, tmp10674)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10677 := Call(__e, PrimFunc(symtail), V1355)


__e.TailApply(tmp10656, tmp10677)
return


}, 1)

tmp10678 := Call(__e, PrimFunc(symhead), V1355)


tmp10679 := Call(__e, tmp10655, tmp10678)


ifres10654 = tmp10679


} else {
tmp10680 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10654 = tmp10680


}

__e.TailApply(tmp10581, ifres10654)
return


} else {
__e.Return(W1374)
return
}


}, 1)

tmp10711 := Call(__e, PrimFunc(symshen_4ccons_2), V1355)


var ifres10684 Obj

if True == tmp10711 {
tmp10685 := MakeNative(func(__e *ControlFlow) {
W1375 := __e.Get(1)
_ = W1375
tmp10686 := MakeNative(func(__e *ControlFlow) {
W1376 := __e.Get(1)
_ = W1376
tmp10706 := Call(__e, PrimFunc(symshen_4hds_a_2), W1375, sym_7)


if True == tmp10706 {
tmp10687 := MakeNative(func(__e *ControlFlow) {
W1377 := __e.Get(1)
_ = W1377
tmp10688 := MakeNative(func(__e *ControlFlow) {
W1378 := __e.Get(1)
_ = W1378
tmp10702 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1378)


if True == tmp10702 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10689 := MakeNative(func(__e *ControlFlow) {
W1379 := __e.Get(1)
_ = W1379
tmp10690 := MakeNative(func(__e *ControlFlow) {
W1380 := __e.Get(1)
_ = W1380
tmp10691 := MakeNative(func(__e *ControlFlow) {
W1381 := __e.Get(1)
_ = W1381
tmp10697 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1381)


if True == tmp10697 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10692 := MakeNative(func(__e *ControlFlow) {
W1382 := __e.Get(1)
_ = W1382
tmp10693 := PrimCons(W1379, Nil)

tmp10694 := PrimCons(symshen_4_7m, tmp10693)

__e.TailApply(PrimFunc(symshen_4comb), W1376, tmp10694)
return


}, 1)

tmp10695 := Call(__e, PrimFunc(symshen_4in_1_6), W1381)


__e.TailApply(tmp10692, tmp10695)
return


}


}, 1)

tmp10698 := Call(__e, PrimFunc(sym_5end_6), W1380)


__e.TailApply(tmp10691, tmp10698)
return


}, 1)

tmp10699 := Call(__e, PrimFunc(symshen_4in_1_6), W1378)


__e.TailApply(tmp10690, tmp10699)
return


}, 1)

tmp10700 := Call(__e, PrimFunc(symshen_4_5_1out), W1378)


__e.TailApply(tmp10689, tmp10700)
return


}


}, 1)

tmp10703 := Call(__e, PrimFunc(symshen_4_5hterm_6), W1377)


__e.TailApply(tmp10688, tmp10703)
return


}, 1)

tmp10704 := Call(__e, PrimFunc(symtail), W1375)


__e.TailApply(tmp10687, tmp10704)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10707 := Call(__e, PrimFunc(symtail), V1355)


__e.TailApply(tmp10686, tmp10707)
return


}, 1)

tmp10708 := Call(__e, PrimFunc(symhead), V1355)


tmp10709 := Call(__e, tmp10685, tmp10708)


ifres10684 = tmp10709


} else {
tmp10710 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10684 = tmp10710


}

__e.TailApply(tmp10580, ifres10684)
return


} else {
__e.Return(W1362)
return
}


}, 1)

tmp10750 := Call(__e, PrimFunc(symshen_4ccons_2), V1355)


var ifres10714 Obj

if True == tmp10750 {
tmp10715 := MakeNative(func(__e *ControlFlow) {
W1363 := __e.Get(1)
_ = W1363
tmp10716 := MakeNative(func(__e *ControlFlow) {
W1364 := __e.Get(1)
_ = W1364
tmp10745 := Call(__e, PrimFunc(symshen_4hds_a_2), W1363, symcons)


if True == tmp10745 {
tmp10717 := MakeNative(func(__e *ControlFlow) {
W1365 := __e.Get(1)
_ = W1365
tmp10718 := MakeNative(func(__e *ControlFlow) {
W1366 := __e.Get(1)
_ = W1366
tmp10741 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1366)


if True == tmp10741 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10719 := MakeNative(func(__e *ControlFlow) {
W1367 := __e.Get(1)
_ = W1367
tmp10720 := MakeNative(func(__e *ControlFlow) {
W1368 := __e.Get(1)
_ = W1368
tmp10721 := MakeNative(func(__e *ControlFlow) {
W1369 := __e.Get(1)
_ = W1369
tmp10736 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1369)


if True == tmp10736 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10722 := MakeNative(func(__e *ControlFlow) {
W1370 := __e.Get(1)
_ = W1370
tmp10723 := MakeNative(func(__e *ControlFlow) {
W1371 := __e.Get(1)
_ = W1371
tmp10724 := MakeNative(func(__e *ControlFlow) {
W1372 := __e.Get(1)
_ = W1372
tmp10731 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1372)


if True == tmp10731 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10725 := MakeNative(func(__e *ControlFlow) {
W1373 := __e.Get(1)
_ = W1373
tmp10726 := PrimCons(W1370, Nil)

tmp10727 := PrimCons(W1367, tmp10726)

tmp10728 := PrimCons(symcons, tmp10727)

__e.TailApply(PrimFunc(symshen_4comb), W1364, tmp10728)
return


}, 1)

tmp10729 := Call(__e, PrimFunc(symshen_4in_1_6), W1372)


__e.TailApply(tmp10725, tmp10729)
return


}


}, 1)

tmp10732 := Call(__e, PrimFunc(sym_5end_6), W1371)


__e.TailApply(tmp10724, tmp10732)
return


}, 1)

tmp10733 := Call(__e, PrimFunc(symshen_4in_1_6), W1369)


__e.TailApply(tmp10723, tmp10733)
return


}, 1)

tmp10734 := Call(__e, PrimFunc(symshen_4_5_1out), W1369)


__e.TailApply(tmp10722, tmp10734)
return


}


}, 1)

tmp10737 := Call(__e, PrimFunc(symshen_4_5hterm2_6), W1368)


__e.TailApply(tmp10721, tmp10737)
return


}, 1)

tmp10738 := Call(__e, PrimFunc(symshen_4in_1_6), W1366)


__e.TailApply(tmp10720, tmp10738)
return


}, 1)

tmp10739 := Call(__e, PrimFunc(symshen_4_5_1out), W1366)


__e.TailApply(tmp10719, tmp10739)
return


}


}, 1)

tmp10742 := Call(__e, PrimFunc(symshen_4_5hterm1_6), W1365)


__e.TailApply(tmp10718, tmp10742)
return


}, 1)

tmp10743 := Call(__e, PrimFunc(symtail), W1363)


__e.TailApply(tmp10717, tmp10743)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10746 := Call(__e, PrimFunc(symtail), V1355)


__e.TailApply(tmp10716, tmp10746)
return


}, 1)

tmp10747 := Call(__e, PrimFunc(symhead), V1355)


tmp10748 := Call(__e, tmp10715, tmp10747)


ifres10714 = tmp10748


} else {
tmp10749 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10714 = tmp10749


}

__e.TailApply(tmp10579, ifres10714)
return


} else {
__e.Return(W1359)
return
}


}, 1)

tmp10763 := PrimIsPair(V1355)

var ifres10753 Obj

if True == tmp10763 {
tmp10754 := MakeNative(func(__e *ControlFlow) {
W1360 := __e.Get(1)
_ = W1360
tmp10755 := MakeNative(func(__e *ControlFlow) {
W1361 := __e.Get(1)
_ = W1361
tmp10757 := PrimIntern(MakeString(":"))

tmp10758 := PrimEqual(W1360, tmp10757)

if True == tmp10758 {
__e.TailApply(PrimFunc(symshen_4comb), W1361, W1360)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10759 := Call(__e, PrimFunc(symtail), V1355)


__e.TailApply(tmp10755, tmp10759)
return


}, 1)

tmp10760 := Call(__e, PrimFunc(symhead), V1355)


tmp10761 := Call(__e, tmp10754, tmp10760)


ifres10753 = tmp10761


} else {
tmp10762 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10753 = tmp10762


}

__e.TailApply(tmp10578, ifres10753)
return


} else {
__e.Return(W1356)
return
}


}, 1)

tmp10779 := PrimIsPair(V1355)

var ifres10766 Obj

if True == tmp10779 {
tmp10767 := MakeNative(func(__e *ControlFlow) {
W1357 := __e.Get(1)
_ = W1357
tmp10768 := MakeNative(func(__e *ControlFlow) {
W1358 := __e.Get(1)
_ = W1358
tmp10774 := Call(__e, PrimFunc(symatom_2), W1357)


var ifres10770 Obj

if True == tmp10774 {
tmp10772 := Call(__e, PrimFunc(symshen_4prolog_1keyword_2), W1357)


tmp10773 := PrimNot(tmp10772)

var ifres10771 Obj

if True == tmp10773 {
ifres10771 = True


} else {
ifres10771 = False


}

ifres10770 = ifres10771


} else {
ifres10770 = False


}

if True == ifres10770 {
__e.TailApply(PrimFunc(symshen_4comb), W1358, W1357)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10775 := Call(__e, PrimFunc(symtail), V1355)


__e.TailApply(tmp10768, tmp10775)
return


}, 1)

tmp10776 := Call(__e, PrimFunc(symhead), V1355)


tmp10777 := Call(__e, tmp10767, tmp10776)


ifres10766 = tmp10777


} else {
tmp10778 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10766 = tmp10778


}

__e.TailApply(tmp10577, ifres10766)
return


}, 1)

tmp10780 := Call(__e, ns2_1set, symshen_4_5hterm_6, tmp10576)


_ = tmp10780

tmp10781 := MakeNative(func(__e *ControlFlow) {
V1412 := __e.Get(1)
_ = V1412
tmp10782 := PrimIntern(MakeString(";"))

tmp10783 := PrimCons(sym_5_1_1, Nil)

tmp10784 := PrimCons(tmp10782, tmp10783)

__e.TailApply(PrimFunc(symelement_2), V1412, tmp10784)
return


}, 1)

tmp10785 := Call(__e, ns2_1set, symshen_4prolog_1keyword_2, tmp10781)


_ = tmp10785

tmp10786 := MakeNative(func(__e *ControlFlow) {
V1413 := __e.Get(1)
_ = V1413
tmp10799 := PrimIsSymbol(V1413)

if True == tmp10799 {
__e.Return(True)
return
} else {
tmp10797 := PrimIsString(V1413)

var ifres10788 Obj

if True == tmp10797 {
ifres10788 = True


} else {
tmp10796 := Call(__e, PrimFunc(symboolean_2), V1413)


var ifres10790 Obj

if True == tmp10796 {
ifres10790 = True


} else {
tmp10795 := PrimIsNumber(V1413)

var ifres10792 Obj

if True == tmp10795 {
ifres10792 = True


} else {
tmp10794 := Call(__e, PrimFunc(symempty_2), V1413)


var ifres10793 Obj

if True == tmp10794 {
ifres10793 = True


} else {
ifres10793 = False


}

ifres10792 = ifres10793


}

var ifres10791 Obj

if True == ifres10792 {
ifres10791 = True


} else {
ifres10791 = False


}

ifres10790 = ifres10791


}

var ifres10789 Obj

if True == ifres10790 {
ifres10789 = True


} else {
ifres10789 = False


}

ifres10788 = ifres10789


}

if True == ifres10788 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp10800 := Call(__e, ns2_1set, symatom_2, tmp10786)


_ = tmp10800

tmp10801 := MakeNative(func(__e *ControlFlow) {
V1414 := __e.Get(1)
_ = V1414
tmp10802 := MakeNative(func(__e *ControlFlow) {
W1415 := __e.Get(1)
_ = W1415
tmp10804 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1415)


if True == tmp10804 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1415)
return
}


}, 1)

tmp10805 := MakeNative(func(__e *ControlFlow) {
W1416 := __e.Get(1)
_ = W1416
tmp10811 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1416)


if True == tmp10811 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10806 := MakeNative(func(__e *ControlFlow) {
W1417 := __e.Get(1)
_ = W1417
tmp10807 := MakeNative(func(__e *ControlFlow) {
W1418 := __e.Get(1)
_ = W1418
__e.TailApply(PrimFunc(symshen_4comb), W1418, W1417)
return
}, 1)

tmp10808 := Call(__e, PrimFunc(symshen_4in_1_6), W1416)


__e.TailApply(tmp10807, tmp10808)
return


}, 1)

tmp10809 := Call(__e, PrimFunc(symshen_4_5_1out), W1416)


__e.TailApply(tmp10806, tmp10809)
return


}


}, 1)

tmp10812 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1414)


tmp10813 := Call(__e, tmp10805, tmp10812)


__e.TailApply(tmp10802, tmp10813)
return


}, 1)

tmp10814 := Call(__e, ns2_1set, symshen_4_5hterm1_6, tmp10801)


_ = tmp10814

tmp10815 := MakeNative(func(__e *ControlFlow) {
V1419 := __e.Get(1)
_ = V1419
tmp10816 := MakeNative(func(__e *ControlFlow) {
W1420 := __e.Get(1)
_ = W1420
tmp10818 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1420)


if True == tmp10818 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1420)
return
}


}, 1)

tmp10819 := MakeNative(func(__e *ControlFlow) {
W1421 := __e.Get(1)
_ = W1421
tmp10825 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1421)


if True == tmp10825 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10820 := MakeNative(func(__e *ControlFlow) {
W1422 := __e.Get(1)
_ = W1422
tmp10821 := MakeNative(func(__e *ControlFlow) {
W1423 := __e.Get(1)
_ = W1423
__e.TailApply(PrimFunc(symshen_4comb), W1423, W1422)
return
}, 1)

tmp10822 := Call(__e, PrimFunc(symshen_4in_1_6), W1421)


__e.TailApply(tmp10821, tmp10822)
return


}, 1)

tmp10823 := Call(__e, PrimFunc(symshen_4_5_1out), W1421)


__e.TailApply(tmp10820, tmp10823)
return


}


}, 1)

tmp10826 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1419)


tmp10827 := Call(__e, tmp10819, tmp10826)


__e.TailApply(tmp10816, tmp10827)
return


}, 1)

tmp10828 := Call(__e, ns2_1set, symshen_4_5hterm2_6, tmp10815)


_ = tmp10828

tmp10829 := MakeNative(func(__e *ControlFlow) {
V1424 := __e.Get(1)
_ = V1424
tmp10830 := MakeNative(func(__e *ControlFlow) {
W1425 := __e.Get(1)
_ = W1425
tmp10842 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1425)


if True == tmp10842 {
tmp10831 := MakeNative(func(__e *ControlFlow) {
W1432 := __e.Get(1)
_ = W1432
tmp10833 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1432)


if True == tmp10833 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1432)
return
}


}, 1)

tmp10834 := MakeNative(func(__e *ControlFlow) {
W1433 := __e.Get(1)
_ = W1433
tmp10838 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1433)


if True == tmp10838 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10835 := MakeNative(func(__e *ControlFlow) {
W1434 := __e.Get(1)
_ = W1434
__e.TailApply(PrimFunc(symshen_4comb), W1434, Nil)
return
}, 1)

tmp10836 := Call(__e, PrimFunc(symshen_4in_1_6), W1433)


__e.TailApply(tmp10835, tmp10836)
return


}


}, 1)

tmp10839 := Call(__e, PrimFunc(sym_5e_6), V1424)


tmp10840 := Call(__e, tmp10834, tmp10839)


__e.TailApply(tmp10831, tmp10840)
return


} else {
__e.Return(W1425)
return
}


}, 1)

tmp10843 := MakeNative(func(__e *ControlFlow) {
W1426 := __e.Get(1)
_ = W1426
tmp10858 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1426)


if True == tmp10858 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10844 := MakeNative(func(__e *ControlFlow) {
W1427 := __e.Get(1)
_ = W1427
tmp10845 := MakeNative(func(__e *ControlFlow) {
W1428 := __e.Get(1)
_ = W1428
tmp10846 := MakeNative(func(__e *ControlFlow) {
W1429 := __e.Get(1)
_ = W1429
tmp10853 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1429)


if True == tmp10853 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10847 := MakeNative(func(__e *ControlFlow) {
W1430 := __e.Get(1)
_ = W1430
tmp10848 := MakeNative(func(__e *ControlFlow) {
W1431 := __e.Get(1)
_ = W1431
tmp10849 := PrimCons(W1427, W1430)

__e.TailApply(PrimFunc(symshen_4comb), W1431, tmp10849)
return


}, 1)

tmp10850 := Call(__e, PrimFunc(symshen_4in_1_6), W1429)


__e.TailApply(tmp10848, tmp10850)
return


}, 1)

tmp10851 := Call(__e, PrimFunc(symshen_4_5_1out), W1429)


__e.TailApply(tmp10847, tmp10851)
return


}


}, 1)

tmp10854 := Call(__e, PrimFunc(symshen_4_5body_6), W1428)


__e.TailApply(tmp10846, tmp10854)
return


}, 1)

tmp10855 := Call(__e, PrimFunc(symshen_4in_1_6), W1426)


__e.TailApply(tmp10845, tmp10855)
return


}, 1)

tmp10856 := Call(__e, PrimFunc(symshen_4_5_1out), W1426)


__e.TailApply(tmp10844, tmp10856)
return


}


}, 1)

tmp10859 := Call(__e, PrimFunc(symshen_4_5literal_6), V1424)


tmp10860 := Call(__e, tmp10843, tmp10859)


__e.TailApply(tmp10830, tmp10860)
return


}, 1)

tmp10861 := Call(__e, ns2_1set, symshen_4_5body_6, tmp10829)


_ = tmp10861

tmp10862 := MakeNative(func(__e *ControlFlow) {
V1435 := __e.Get(1)
_ = V1435
tmp10863 := MakeNative(func(__e *ControlFlow) {
W1436 := __e.Get(1)
_ = W1436
tmp10890 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1436)


if True == tmp10890 {
tmp10864 := MakeNative(func(__e *ControlFlow) {
W1438 := __e.Get(1)
_ = W1438
tmp10866 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1438)


if True == tmp10866 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1438)
return
}


}, 1)

tmp10888 := Call(__e, PrimFunc(symshen_4ccons_2), V1435)


var ifres10867 Obj

if True == tmp10888 {
tmp10868 := MakeNative(func(__e *ControlFlow) {
W1439 := __e.Get(1)
_ = W1439
tmp10869 := MakeNative(func(__e *ControlFlow) {
W1440 := __e.Get(1)
_ = W1440
tmp10870 := MakeNative(func(__e *ControlFlow) {
W1441 := __e.Get(1)
_ = W1441
tmp10882 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1441)


if True == tmp10882 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10871 := MakeNative(func(__e *ControlFlow) {
W1442 := __e.Get(1)
_ = W1442
tmp10872 := MakeNative(func(__e *ControlFlow) {
W1443 := __e.Get(1)
_ = W1443
tmp10873 := MakeNative(func(__e *ControlFlow) {
W1444 := __e.Get(1)
_ = W1444
tmp10877 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1444)


if True == tmp10877 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10874 := MakeNative(func(__e *ControlFlow) {
W1445 := __e.Get(1)
_ = W1445
__e.TailApply(PrimFunc(symshen_4comb), W1440, W1442)
return
}, 1)

tmp10875 := Call(__e, PrimFunc(symshen_4in_1_6), W1444)


__e.TailApply(tmp10874, tmp10875)
return


}


}, 1)

tmp10878 := Call(__e, PrimFunc(sym_5end_6), W1443)


__e.TailApply(tmp10873, tmp10878)
return


}, 1)

tmp10879 := Call(__e, PrimFunc(symshen_4in_1_6), W1441)


__e.TailApply(tmp10872, tmp10879)
return


}, 1)

tmp10880 := Call(__e, PrimFunc(symshen_4_5_1out), W1441)


__e.TailApply(tmp10871, tmp10880)
return


}


}, 1)

tmp10883 := Call(__e, PrimFunc(symshen_4_5bterms_6), W1439)


__e.TailApply(tmp10870, tmp10883)
return


}, 1)

tmp10884 := Call(__e, PrimFunc(symtail), V1435)


__e.TailApply(tmp10869, tmp10884)
return


}, 1)

tmp10885 := Call(__e, PrimFunc(symhead), V1435)


tmp10886 := Call(__e, tmp10868, tmp10885)


ifres10867 = tmp10886


} else {
tmp10887 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10867 = tmp10887


}

__e.TailApply(tmp10864, ifres10867)
return


} else {
__e.Return(W1436)
return
}


}, 1)

tmp10896 := Call(__e, PrimFunc(symshen_4hds_a_2), V1435, sym_b)


var ifres10891 Obj

if True == tmp10896 {
tmp10892 := MakeNative(func(__e *ControlFlow) {
W1437 := __e.Get(1)
_ = W1437
__e.TailApply(PrimFunc(symshen_4comb), W1437, sym_b)
return
}, 1)

tmp10893 := Call(__e, PrimFunc(symtail), V1435)


tmp10894 := Call(__e, tmp10892, tmp10893)


ifres10891 = tmp10894


} else {
tmp10895 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10891 = tmp10895


}

__e.TailApply(tmp10863, ifres10891)
return


}, 1)

tmp10897 := Call(__e, ns2_1set, symshen_4_5literal_6, tmp10862)


_ = tmp10897

tmp10898 := MakeNative(func(__e *ControlFlow) {
V1446 := __e.Get(1)
_ = V1446
tmp10899 := MakeNative(func(__e *ControlFlow) {
W1447 := __e.Get(1)
_ = W1447
tmp10911 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1447)


if True == tmp10911 {
tmp10900 := MakeNative(func(__e *ControlFlow) {
W1454 := __e.Get(1)
_ = W1454
tmp10902 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1454)


if True == tmp10902 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1454)
return
}


}, 1)

tmp10903 := MakeNative(func(__e *ControlFlow) {
W1455 := __e.Get(1)
_ = W1455
tmp10907 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1455)


if True == tmp10907 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10904 := MakeNative(func(__e *ControlFlow) {
W1456 := __e.Get(1)
_ = W1456
__e.TailApply(PrimFunc(symshen_4comb), W1456, Nil)
return
}, 1)

tmp10905 := Call(__e, PrimFunc(symshen_4in_1_6), W1455)


__e.TailApply(tmp10904, tmp10905)
return


}


}, 1)

tmp10908 := Call(__e, PrimFunc(sym_5e_6), V1446)


tmp10909 := Call(__e, tmp10903, tmp10908)


__e.TailApply(tmp10900, tmp10909)
return


} else {
__e.Return(W1447)
return
}


}, 1)

tmp10912 := MakeNative(func(__e *ControlFlow) {
W1448 := __e.Get(1)
_ = W1448
tmp10927 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1448)


if True == tmp10927 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10913 := MakeNative(func(__e *ControlFlow) {
W1449 := __e.Get(1)
_ = W1449
tmp10914 := MakeNative(func(__e *ControlFlow) {
W1450 := __e.Get(1)
_ = W1450
tmp10915 := MakeNative(func(__e *ControlFlow) {
W1451 := __e.Get(1)
_ = W1451
tmp10922 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1451)


if True == tmp10922 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10916 := MakeNative(func(__e *ControlFlow) {
W1452 := __e.Get(1)
_ = W1452
tmp10917 := MakeNative(func(__e *ControlFlow) {
W1453 := __e.Get(1)
_ = W1453
tmp10918 := PrimCons(W1449, W1452)

__e.TailApply(PrimFunc(symshen_4comb), W1453, tmp10918)
return


}, 1)

tmp10919 := Call(__e, PrimFunc(symshen_4in_1_6), W1451)


__e.TailApply(tmp10917, tmp10919)
return


}, 1)

tmp10920 := Call(__e, PrimFunc(symshen_4_5_1out), W1451)


__e.TailApply(tmp10916, tmp10920)
return


}


}, 1)

tmp10923 := Call(__e, PrimFunc(symshen_4_5bterms_6), W1450)


__e.TailApply(tmp10915, tmp10923)
return


}, 1)

tmp10924 := Call(__e, PrimFunc(symshen_4in_1_6), W1448)


__e.TailApply(tmp10914, tmp10924)
return


}, 1)

tmp10925 := Call(__e, PrimFunc(symshen_4_5_1out), W1448)


__e.TailApply(tmp10913, tmp10925)
return


}


}, 1)

tmp10928 := Call(__e, PrimFunc(symshen_4_5bterm_6), V1446)


tmp10929 := Call(__e, tmp10912, tmp10928)


__e.TailApply(tmp10899, tmp10929)
return


}, 1)

tmp10930 := Call(__e, ns2_1set, symshen_4_5bterms_6, tmp10898)


_ = tmp10930

tmp10931 := MakeNative(func(__e *ControlFlow) {
V1457 := __e.Get(1)
_ = V1457
tmp10932 := MakeNative(func(__e *ControlFlow) {
W1458 := __e.Get(1)
_ = W1458
tmp10972 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1458)


if True == tmp10972 {
tmp10933 := MakeNative(func(__e *ControlFlow) {
W1462 := __e.Get(1)
_ = W1462
tmp10960 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1462)


if True == tmp10960 {
tmp10934 := MakeNative(func(__e *ControlFlow) {
W1465 := __e.Get(1)
_ = W1465
tmp10936 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1465)


if True == tmp10936 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1465)
return
}


}, 1)

tmp10958 := Call(__e, PrimFunc(symshen_4ccons_2), V1457)


var ifres10937 Obj

if True == tmp10958 {
tmp10938 := MakeNative(func(__e *ControlFlow) {
W1466 := __e.Get(1)
_ = W1466
tmp10939 := MakeNative(func(__e *ControlFlow) {
W1467 := __e.Get(1)
_ = W1467
tmp10940 := MakeNative(func(__e *ControlFlow) {
W1468 := __e.Get(1)
_ = W1468
tmp10952 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1468)


if True == tmp10952 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10941 := MakeNative(func(__e *ControlFlow) {
W1469 := __e.Get(1)
_ = W1469
tmp10942 := MakeNative(func(__e *ControlFlow) {
W1470 := __e.Get(1)
_ = W1470
tmp10943 := MakeNative(func(__e *ControlFlow) {
W1471 := __e.Get(1)
_ = W1471
tmp10947 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1471)


if True == tmp10947 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10944 := MakeNative(func(__e *ControlFlow) {
W1472 := __e.Get(1)
_ = W1472
__e.TailApply(PrimFunc(symshen_4comb), W1467, W1469)
return
}, 1)

tmp10945 := Call(__e, PrimFunc(symshen_4in_1_6), W1471)


__e.TailApply(tmp10944, tmp10945)
return


}


}, 1)

tmp10948 := Call(__e, PrimFunc(sym_5end_6), W1470)


__e.TailApply(tmp10943, tmp10948)
return


}, 1)

tmp10949 := Call(__e, PrimFunc(symshen_4in_1_6), W1468)


__e.TailApply(tmp10942, tmp10949)
return


}, 1)

tmp10950 := Call(__e, PrimFunc(symshen_4_5_1out), W1468)


__e.TailApply(tmp10941, tmp10950)
return


}


}, 1)

tmp10953 := Call(__e, PrimFunc(symshen_4_5bterms_6), W1466)


__e.TailApply(tmp10940, tmp10953)
return


}, 1)

tmp10954 := Call(__e, PrimFunc(symtail), V1457)


__e.TailApply(tmp10939, tmp10954)
return


}, 1)

tmp10955 := Call(__e, PrimFunc(symhead), V1457)


tmp10956 := Call(__e, tmp10938, tmp10955)


ifres10937 = tmp10956


} else {
tmp10957 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10937 = tmp10957


}

__e.TailApply(tmp10934, ifres10937)
return


} else {
__e.Return(W1462)
return
}


}, 1)

tmp10970 := PrimIsPair(V1457)

var ifres10961 Obj

if True == tmp10970 {
tmp10962 := MakeNative(func(__e *ControlFlow) {
W1463 := __e.Get(1)
_ = W1463
tmp10963 := MakeNative(func(__e *ControlFlow) {
W1464 := __e.Get(1)
_ = W1464
tmp10965 := Call(__e, PrimFunc(symatom_2), W1463)


if True == tmp10965 {
__e.TailApply(PrimFunc(symshen_4comb), W1464, W1463)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10966 := Call(__e, PrimFunc(symtail), V1457)


__e.TailApply(tmp10963, tmp10966)
return


}, 1)

tmp10967 := Call(__e, PrimFunc(symhead), V1457)


tmp10968 := Call(__e, tmp10962, tmp10967)


ifres10961 = tmp10968


} else {
tmp10969 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10961 = tmp10969


}

__e.TailApply(tmp10933, ifres10961)
return


} else {
__e.Return(W1458)
return
}


}, 1)

tmp10973 := MakeNative(func(__e *ControlFlow) {
W1459 := __e.Get(1)
_ = W1459
tmp10979 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1459)


if True == tmp10979 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10974 := MakeNative(func(__e *ControlFlow) {
W1460 := __e.Get(1)
_ = W1460
tmp10975 := MakeNative(func(__e *ControlFlow) {
W1461 := __e.Get(1)
_ = W1461
__e.TailApply(PrimFunc(symshen_4comb), W1461, W1460)
return
}, 1)

tmp10976 := Call(__e, PrimFunc(symshen_4in_1_6), W1459)


__e.TailApply(tmp10975, tmp10976)
return


}, 1)

tmp10977 := Call(__e, PrimFunc(symshen_4_5_1out), W1459)


__e.TailApply(tmp10974, tmp10977)
return


}


}, 1)

tmp10980 := Call(__e, PrimFunc(symshen_4_5wildcard_6), V1457)


tmp10981 := Call(__e, tmp10973, tmp10980)


__e.TailApply(tmp10932, tmp10981)
return


}, 1)

tmp10982 := Call(__e, ns2_1set, symshen_4_5bterm_6, tmp10931)


_ = tmp10982

tmp10983 := MakeNative(func(__e *ControlFlow) {
V1473 := __e.Get(1)
_ = V1473
tmp10984 := MakeNative(func(__e *ControlFlow) {
W1474 := __e.Get(1)
_ = W1474
tmp10986 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1474)


if True == tmp10986 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1474)
return
}


}, 1)

tmp10997 := PrimIsPair(V1473)

var ifres10987 Obj

if True == tmp10997 {
tmp10988 := MakeNative(func(__e *ControlFlow) {
W1475 := __e.Get(1)
_ = W1475
tmp10989 := MakeNative(func(__e *ControlFlow) {
W1476 := __e.Get(1)
_ = W1476
tmp10992 := PrimEqual(W1475, sym__)

if True == tmp10992 {
tmp10990 := Call(__e, PrimFunc(symgensym), symY)


__e.TailApply(PrimFunc(symshen_4comb), W1476, tmp10990)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10993 := Call(__e, PrimFunc(symtail), V1473)


__e.TailApply(tmp10989, tmp10993)
return


}, 1)

tmp10994 := Call(__e, PrimFunc(symhead), V1473)


tmp10995 := Call(__e, tmp10988, tmp10994)


ifres10987 = tmp10995


} else {
tmp10996 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10987 = tmp10996


}

__e.TailApply(tmp10984, ifres10987)
return


}, 1)

tmp10998 := Call(__e, ns2_1set, symshen_4_5wildcard_6, tmp10983)


_ = tmp10998

tmp10999 := MakeNative(func(__e *ControlFlow) {
V1477 := __e.Get(1)
_ = V1477
tmp11000 := MakeNative(func(__e *ControlFlow) {
W1478 := __e.Get(1)
_ = W1478
tmp11002 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1478)


if True == tmp11002 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1478)
return
}


}, 1)

tmp11012 := PrimIsPair(V1477)

var ifres11003 Obj

if True == tmp11012 {
tmp11004 := MakeNative(func(__e *ControlFlow) {
W1479 := __e.Get(1)
_ = W1479
tmp11005 := MakeNative(func(__e *ControlFlow) {
W1480 := __e.Get(1)
_ = W1480
tmp11007 := Call(__e, PrimFunc(symshen_4semicolon_2), W1479)


if True == tmp11007 {
__e.TailApply(PrimFunc(symshen_4comb), W1480, W1479)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp11008 := Call(__e, PrimFunc(symtail), V1477)


__e.TailApply(tmp11005, tmp11008)
return


}, 1)

tmp11009 := Call(__e, PrimFunc(symhead), V1477)


tmp11010 := Call(__e, tmp11004, tmp11009)


ifres11003 = tmp11010


} else {
tmp11011 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11003 = tmp11011


}

__e.TailApply(tmp11000, ifres11003)
return


}, 1)

tmp11013 := Call(__e, ns2_1set, symshen_4_5sc_6, tmp10999)


_ = tmp11013

tmp11014 := MakeNative(func(__e *ControlFlow) {
V1481 := __e.Get(1)
_ = V1481
V1482 := __e.Get(2)
_ = V1482
tmp11015 := MakeNative(func(__e *ControlFlow) {
W1483 := __e.Get(1)
_ = W1483
tmp11016 := MakeNative(func(__e *ControlFlow) {
W1484 := __e.Get(1)
_ = W1484
tmp11017 := MakeNative(func(__e *ControlFlow) {
W1485 := __e.Get(1)
_ = W1485
tmp11018 := MakeNative(func(__e *ControlFlow) {
W1486 := __e.Get(1)
_ = W1486
tmp11019 := MakeNative(func(__e *ControlFlow) {
W1487 := __e.Get(1)
_ = W1487
tmp11020 := MakeNative(func(__e *ControlFlow) {
W1488 := __e.Get(1)
_ = W1488
tmp11021 := MakeNative(func(__e *ControlFlow) {
W1489 := __e.Get(1)
_ = W1489
tmp11022 := MakeNative(func(__e *ControlFlow) {
W1490 := __e.Get(1)
_ = W1490
tmp11023 := MakeNative(func(__e *ControlFlow) {
W1491 := __e.Get(1)
_ = W1491
__e.Return(W1491)
return
}, 1)

tmp11024 := PrimCons(sym_1_6, Nil)

tmp11025 := PrimCons(W1486, tmp11024)

tmp11026 := PrimCons(W1485, tmp11025)

tmp11027 := PrimCons(W1484, tmp11026)

tmp11028 := PrimCons(W1483, tmp11027)

tmp11029 := PrimCons(W1490, Nil)

tmp11030 := Call(__e, PrimFunc(symappend), tmp11028, tmp11029)


tmp11031 := Call(__e, PrimFunc(symappend), W1487, tmp11030)


tmp11032 := PrimCons(V1481, tmp11031)

tmp11033 := PrimCons(symdefine, tmp11032)

__e.TailApply(tmp11023, tmp11033)
return


}, 1)

var ifres11034 Obj

if True == W1488 {
tmp11035 := PrimCons(MakeNumber(1), Nil)

tmp11036 := PrimCons(W1485, tmp11035)

tmp11037 := PrimCons(sym_7, tmp11036)

tmp11038 := PrimCons(W1489, Nil)

tmp11039 := PrimCons(tmp11037, tmp11038)

tmp11040 := PrimCons(W1485, tmp11039)

tmp11041 := PrimCons(symlet, tmp11040)

ifres11034 = tmp11041


} else {
ifres11034 = W1489


}

__e.TailApply(tmp11022, ifres11034)
return


}, 1)

tmp11042 := Call(__e, PrimFunc(symshen_4prolog_1fbody), V1482, W1487, W1483, W1484, W1485, W1486, W1488)


__e.TailApply(tmp11021, tmp11042)
return


}, 1)

tmp11043 := Call(__e, PrimFunc(symshen_4hascut_2), V1482)


__e.TailApply(tmp11020, tmp11043)
return


}, 1)

tmp11044 := Call(__e, PrimFunc(symshen_4prolog_1parameters), V1482)


__e.TailApply(tmp11019, tmp11044)
return


}, 1)

tmp11045 := Call(__e, PrimFunc(symgensym), symC)


__e.TailApply(tmp11018, tmp11045)
return


}, 1)

tmp11046 := Call(__e, PrimFunc(symgensym), symK)


__e.TailApply(tmp11017, tmp11046)
return


}, 1)

tmp11047 := Call(__e, PrimFunc(symgensym), symL)


__e.TailApply(tmp11016, tmp11047)
return


}, 1)

tmp11048 := Call(__e, PrimFunc(symgensym), symB)


__e.TailApply(tmp11015, tmp11048)
return


}, 2)

tmp11049 := Call(__e, ns2_1set, symshen_4horn_1clause_1procedure, tmp11014)


_ = tmp11049

tmp11050 := MakeNative(func(__e *ControlFlow) {
V1494 := __e.Get(1)
_ = V1494
tmp11060 := PrimEqual(sym_b, V1494)

if True == tmp11060 {
__e.Return(True)
return
} else {
tmp11058 := PrimIsPair(V1494)

if True == tmp11058 {
tmp11055 := PrimHead(V1494)

tmp11056 := Call(__e, PrimFunc(symshen_4hascut_2), tmp11055)


if True == tmp11056 {
__e.Return(True)
return
} else {
tmp11052 := PrimTail(V1494)

tmp11053 := Call(__e, PrimFunc(symshen_4hascut_2), tmp11052)


if True == tmp11053 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


} else {
__e.Return(False)
return
}


}


}, 1)

tmp11061 := Call(__e, ns2_1set, symshen_4hascut_2, tmp11050)


_ = tmp11061

tmp11062 := MakeNative(func(__e *ControlFlow) {
V1499 := __e.Get(1)
_ = V1499
tmp11071 := PrimIsPair(V1499)

var ifres11067 Obj

if True == tmp11071 {
tmp11069 := PrimHead(V1499)

tmp11070 := PrimIsPair(tmp11069)

var ifres11068 Obj

if True == tmp11070 {
ifres11068 = True


} else {
ifres11068 = False


}

ifres11067 = ifres11068


} else {
ifres11067 = False


}

if True == ifres11067 {
tmp11063 := PrimHead(V1499)

tmp11064 := PrimHead(tmp11063)

tmp11065 := Call(__e, PrimFunc(symlength), tmp11064)


__e.TailApply(PrimFunc(symshen_4parameters), tmp11065)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.prolog-parameters")))
return
}


}, 1)

tmp11072 := Call(__e, ns2_1set, symshen_4prolog_1parameters, tmp11062)


_ = tmp11072

tmp11073 := MakeNative(func(__e *ControlFlow) {
V1520 := __e.Get(1)
_ = V1520
V1521 := __e.Get(2)
_ = V1521
V1522 := __e.Get(3)
_ = V1522
V1523 := __e.Get(4)
_ = V1523
V1524 := __e.Get(5)
_ = V1524
V1525 := __e.Get(6)
_ = V1525
V1526 := __e.Get(7)
_ = V1526
tmp11166 := PrimEqual(Nil, V1520)

var ifres11163 Obj

if True == tmp11166 {
tmp11165 := PrimEqual(True, V1526)

var ifres11164 Obj

if True == tmp11165 {
ifres11164 = True


} else {
ifres11164 = False


}

ifres11163 = ifres11164


} else {
ifres11163 = False


}

if True == ifres11163 {
tmp11074 := PrimCons(V1524, Nil)

tmp11075 := PrimCons(V1523, tmp11074)

__e.Return(PrimCons(symshen_4unlock, tmp11075))
return


} else {
tmp11161 := PrimIsPair(V1520)

var ifres11139 Obj

if True == tmp11161 {
tmp11159 := PrimHead(V1520)

tmp11160 := PrimIsPair(tmp11159)

var ifres11141 Obj

if True == tmp11160 {
tmp11156 := PrimHead(V1520)

tmp11157 := PrimTail(tmp11156)

tmp11158 := PrimIsPair(tmp11157)

var ifres11143 Obj

if True == tmp11158 {
tmp11152 := PrimHead(V1520)

tmp11153 := PrimTail(tmp11152)

tmp11154 := PrimTail(tmp11153)

tmp11155 := PrimEqual(Nil, tmp11154)

var ifres11145 Obj

if True == tmp11155 {
tmp11150 := PrimTail(V1520)

tmp11151 := PrimEqual(Nil, tmp11150)

var ifres11147 Obj

if True == tmp11151 {
tmp11149 := PrimEqual(False, V1526)

var ifres11148 Obj

if True == tmp11149 {
ifres11148 = True


} else {
ifres11148 = False


}

ifres11147 = ifres11148


} else {
ifres11147 = False


}

var ifres11146 Obj

if True == ifres11147 {
ifres11146 = True


} else {
ifres11146 = False


}

ifres11145 = ifres11146


} else {
ifres11145 = False


}

var ifres11144 Obj

if True == ifres11145 {
ifres11144 = True


} else {
ifres11144 = False


}

ifres11143 = ifres11144


} else {
ifres11143 = False


}

var ifres11142 Obj

if True == ifres11143 {
ifres11142 = True


} else {
ifres11142 = False


}

ifres11141 = ifres11142


} else {
ifres11141 = False


}

var ifres11140 Obj

if True == ifres11141 {
ifres11140 = True


} else {
ifres11140 = False


}

ifres11139 = ifres11140


} else {
ifres11139 = False


}

if True == ifres11139 {
tmp11076 := MakeNative(func(__e *ControlFlow) {
W1527 := __e.Get(1)
_ = W1527
tmp11077 := PrimCons(V1523, Nil)

tmp11078 := PrimCons(symshen_4unlocked_2, tmp11077)

tmp11079 := PrimHead(V1520)

tmp11080 := PrimHead(tmp11079)

tmp11081 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11080, V1521, V1522, W1527)


tmp11082 := PrimCons(False, Nil)

tmp11083 := PrimCons(tmp11081, tmp11082)

tmp11084 := PrimCons(tmp11078, tmp11083)

__e.Return(PrimCons(symif, tmp11084))
return


}, 1)

tmp11085 := PrimHead(V1520)

tmp11086 := PrimHead(tmp11085)

tmp11087 := PrimHead(V1520)

tmp11088 := PrimTail(tmp11087)

tmp11089 := PrimHead(tmp11088)

tmp11090 := Call(__e, PrimFunc(symshen_4continue), tmp11086, tmp11089, V1522, V1523, V1524, V1525)


__e.TailApply(tmp11076, tmp11090)
return


} else {
tmp11137 := PrimIsPair(V1520)

var ifres11122 Obj

if True == tmp11137 {
tmp11135 := PrimHead(V1520)

tmp11136 := PrimIsPair(tmp11135)

var ifres11124 Obj

if True == tmp11136 {
tmp11132 := PrimHead(V1520)

tmp11133 := PrimTail(tmp11132)

tmp11134 := PrimIsPair(tmp11133)

var ifres11126 Obj

if True == tmp11134 {
tmp11128 := PrimHead(V1520)

tmp11129 := PrimTail(tmp11128)

tmp11130 := PrimTail(tmp11129)

tmp11131 := PrimEqual(Nil, tmp11130)

var ifres11127 Obj

if True == tmp11131 {
ifres11127 = True


} else {
ifres11127 = False


}

ifres11126 = ifres11127


} else {
ifres11126 = False


}

var ifres11125 Obj

if True == ifres11126 {
ifres11125 = True


} else {
ifres11125 = False


}

ifres11124 = ifres11125


} else {
ifres11124 = False


}

var ifres11123 Obj

if True == ifres11124 {
ifres11123 = True


} else {
ifres11123 = False


}

ifres11122 = ifres11123


} else {
ifres11122 = False


}

if True == ifres11122 {
tmp11091 := MakeNative(func(__e *ControlFlow) {
W1528 := __e.Get(1)
_ = W1528
tmp11092 := MakeNative(func(__e *ControlFlow) {
W1529 := __e.Get(1)
_ = W1529
tmp11093 := PrimCons(V1523, Nil)

tmp11094 := PrimCons(symshen_4unlocked_2, tmp11093)

tmp11095 := PrimHead(V1520)

tmp11096 := PrimHead(tmp11095)

tmp11097 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11096, V1521, V1522, W1529)


tmp11098 := PrimCons(False, Nil)

tmp11099 := PrimCons(tmp11097, tmp11098)

tmp11100 := PrimCons(tmp11094, tmp11099)

tmp11101 := PrimCons(symif, tmp11100)

tmp11102 := PrimCons(False, Nil)

tmp11103 := PrimCons(W1528, tmp11102)

tmp11104 := PrimCons(sym_a, tmp11103)

tmp11105 := PrimTail(V1520)

tmp11106 := Call(__e, PrimFunc(symshen_4prolog_1fbody), tmp11105, V1521, V1522, V1523, V1524, V1525, V1526)


tmp11107 := PrimCons(W1528, Nil)

tmp11108 := PrimCons(tmp11106, tmp11107)

tmp11109 := PrimCons(tmp11104, tmp11108)

tmp11110 := PrimCons(symif, tmp11109)

tmp11111 := PrimCons(tmp11110, Nil)

tmp11112 := PrimCons(tmp11101, tmp11111)

tmp11113 := PrimCons(W1528, tmp11112)

__e.Return(PrimCons(symlet, tmp11113))
return


}, 1)

tmp11114 := PrimHead(V1520)

tmp11115 := PrimHead(tmp11114)

tmp11116 := PrimHead(V1520)

tmp11117 := PrimTail(tmp11116)

tmp11118 := PrimHead(tmp11117)

tmp11119 := Call(__e, PrimFunc(symshen_4continue), tmp11115, tmp11118, V1522, V1523, V1524, V1525)


__e.TailApply(tmp11092, tmp11119)
return


}, 1)

tmp11120 := Call(__e, PrimFunc(symgensym), symC)


__e.TailApply(tmp11091, tmp11120)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.prolog-fbody")))
return
}


}


}


}, 7)

tmp11167 := Call(__e, ns2_1set, symshen_4prolog_1fbody, tmp11073)


_ = tmp11167

tmp11168 := MakeNative(func(__e *ControlFlow) {
V1530 := __e.Get(1)
_ = V1530
V1531 := __e.Get(2)
_ = V1531
tmp11173 := Call(__e, PrimFunc(symshen_4locked_2), V1530)


var ifres11170 Obj

if True == tmp11173 {
tmp11172 := Call(__e, PrimFunc(symshen_4fits_2), V1531, V1530)


var ifres11171 Obj

if True == tmp11172 {
ifres11171 = True


} else {
ifres11171 = False


}

ifres11170 = ifres11171


} else {
ifres11170 = False


}

if True == ifres11170 {
__e.TailApply(PrimFunc(symshen_4openlock), V1530)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp11174 := Call(__e, ns2_1set, symshen_4unlock, tmp11168)


_ = tmp11174

tmp11175 := MakeNative(func(__e *ControlFlow) {
V1532 := __e.Get(1)
_ = V1532
tmp11176 := Call(__e, PrimFunc(symshen_4unlocked_2), V1532)


__e.Return(PrimNot(tmp11176))
return


}, 1)

tmp11177 := Call(__e, ns2_1set, symshen_4locked_2, tmp11175)


_ = tmp11177

tmp11178 := MakeNative(func(__e *ControlFlow) {
V1533 := __e.Get(1)
_ = V1533
__e.Return(PrimVectorGet(V1533, MakeNumber(1)))
return
}, 1)

tmp11179 := Call(__e, ns2_1set, symshen_4unlocked_2, tmp11178)


_ = tmp11179

tmp11180 := MakeNative(func(__e *ControlFlow) {
V1534 := __e.Get(1)
_ = V1534
tmp11181 := PrimVectorSet(V1534, MakeNumber(1), True)

_ = tmp11181

__e.Return(False)
return


}, 1)

tmp11182 := Call(__e, ns2_1set, symshen_4openlock, tmp11180)


_ = tmp11182

tmp11183 := MakeNative(func(__e *ControlFlow) {
V1535 := __e.Get(1)
_ = V1535
V1536 := __e.Get(2)
_ = V1536
tmp11184 := PrimVectorGet(V1536, MakeNumber(2))

__e.Return(PrimEqual(V1535, tmp11184))
return


}, 2)

tmp11185 := Call(__e, ns2_1set, symshen_4fits_2, tmp11183)


_ = tmp11185

tmp11186 := MakeNative(func(__e *ControlFlow) {
V1539 := __e.Get(1)
_ = V1539
V1540 := __e.Get(2)
_ = V1540
V1541 := __e.Get(3)
_ = V1541
V1542 := __e.Get(4)
_ = V1542
tmp11187 := MakeNative(func(__e *ControlFlow) {
W1543 := __e.Get(1)
_ = W1543
tmp11192 := PrimEqual(W1543, False)

var ifres11189 Obj

if True == tmp11192 {
tmp11191 := Call(__e, PrimFunc(symshen_4unlocked_2), V1540)


var ifres11190 Obj

if True == tmp11191 {
ifres11190 = True


} else {
ifres11190 = False


}

ifres11189 = ifres11190


} else {
ifres11189 = False


}

if True == ifres11189 {
__e.TailApply(PrimFunc(symshen_4lock), V1541, V1540)
return
} else {
__e.Return(W1543)
return
}


}, 1)

tmp11193 := Call(__e, PrimFunc(symthaw), V1542)


__e.TailApply(tmp11187, tmp11193)
return


}, 4)

tmp11194 := Call(__e, ns2_1set, symshen_4cut, tmp11186)


_ = tmp11194

tmp11195 := MakeNative(func(__e *ControlFlow) {
V1544 := __e.Get(1)
_ = V1544
V1545 := __e.Get(2)
_ = V1545
tmp11196 := MakeNative(func(__e *ControlFlow) {
W1546 := __e.Get(1)
_ = W1546
tmp11197 := MakeNative(func(__e *ControlFlow) {
W1547 := __e.Get(1)
_ = W1547
__e.Return(False)
return
}, 1)

tmp11198 := PrimVectorSet(V1545, MakeNumber(2), V1544)

__e.TailApply(tmp11197, tmp11198)
return


}, 1)

tmp11199 := PrimVectorSet(V1545, MakeNumber(1), False)

__e.TailApply(tmp11196, tmp11199)
return


}, 2)

tmp11200 := Call(__e, ns2_1set, symshen_4lock, tmp11195)


_ = tmp11200

tmp11201 := MakeNative(func(__e *ControlFlow) {
V1548 := __e.Get(1)
_ = V1548
V1549 := __e.Get(2)
_ = V1549
V1550 := __e.Get(3)
_ = V1550
V1551 := __e.Get(4)
_ = V1551
V1552 := __e.Get(5)
_ = V1552
V1553 := __e.Get(6)
_ = V1553
tmp11202 := MakeNative(func(__e *ControlFlow) {
W1554 := __e.Get(1)
_ = W1554
tmp11203 := MakeNative(func(__e *ControlFlow) {
W1555 := __e.Get(1)
_ = W1555
tmp11204 := MakeNative(func(__e *ControlFlow) {
W1556 := __e.Get(1)
_ = W1556
tmp11205 := MakeNative(func(__e *ControlFlow) {
W1557 := __e.Get(1)
_ = W1557
__e.TailApply(PrimFunc(symshen_4stpart), W1556, W1557, V1550)
return
}, 1)

tmp11206 := PrimCons(symshen_4incinfs, Nil)

tmp11207 := Call(__e, PrimFunc(symshen_4compile_1body), V1549, V1550, V1551, V1552, V1553)


tmp11208 := PrimCons(tmp11207, Nil)

tmp11209 := PrimCons(tmp11206, tmp11208)

tmp11210 := PrimCons(symdo, tmp11209)

__e.TailApply(tmp11205, tmp11210)
return


}, 1)

tmp11211 := Call(__e, PrimFunc(symdifference), W1555, W1554)


__e.TailApply(tmp11204, tmp11211)
return


}, 1)

tmp11212 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), V1549)


__e.TailApply(tmp11203, tmp11212)
return


}, 1)

tmp11213 := Call(__e, PrimFunc(symshen_4extract_1vars), V1548)


__e.TailApply(tmp11202, tmp11213)
return


}, 6)

tmp11214 := Call(__e, ns2_1set, symshen_4continue, tmp11201)


_ = tmp11214

tmp11215 := MakeNative(func(__e *ControlFlow) {
V1560 := __e.Get(1)
_ = V1560
tmp11250 := PrimIsPair(V1560)

var ifres11231 Obj

if True == tmp11250 {
tmp11248 := PrimHead(V1560)

tmp11249 := PrimEqual(symlambda, tmp11248)

var ifres11233 Obj

if True == tmp11249 {
tmp11246 := PrimTail(V1560)

tmp11247 := PrimIsPair(tmp11246)

var ifres11235 Obj

if True == tmp11247 {
tmp11243 := PrimTail(V1560)

tmp11244 := PrimTail(tmp11243)

tmp11245 := PrimIsPair(tmp11244)

var ifres11237 Obj

if True == tmp11245 {
tmp11239 := PrimTail(V1560)

tmp11240 := PrimTail(tmp11239)

tmp11241 := PrimTail(tmp11240)

tmp11242 := PrimEqual(Nil, tmp11241)

var ifres11238 Obj

if True == tmp11242 {
ifres11238 = True


} else {
ifres11238 = False


}

ifres11237 = ifres11238


} else {
ifres11237 = False


}

var ifres11236 Obj

if True == ifres11237 {
ifres11236 = True


} else {
ifres11236 = False


}

ifres11235 = ifres11236


} else {
ifres11235 = False


}

var ifres11234 Obj

if True == ifres11235 {
ifres11234 = True


} else {
ifres11234 = False


}

ifres11233 = ifres11234


} else {
ifres11233 = False


}

var ifres11232 Obj

if True == ifres11233 {
ifres11232 = True


} else {
ifres11232 = False


}

ifres11231 = ifres11232


} else {
ifres11231 = False


}

if True == ifres11231 {
tmp11216 := PrimTail(V1560)

tmp11217 := PrimHead(tmp11216)

tmp11218 := PrimTail(V1560)

tmp11219 := PrimTail(tmp11218)

tmp11220 := PrimHead(tmp11219)

tmp11221 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp11220)


__e.TailApply(PrimFunc(symremove), tmp11217, tmp11221)
return


} else {
tmp11229 := PrimIsPair(V1560)

if True == tmp11229 {
tmp11222 := PrimHead(V1560)

tmp11223 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp11222)


tmp11224 := PrimTail(V1560)

tmp11225 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp11224)


__e.TailApply(PrimFunc(symunion), tmp11223, tmp11225)
return


} else {
tmp11227 := PrimIsVariable(V1560)

if True == tmp11227 {
__e.Return(PrimCons(V1560, Nil))
return
} else {
__e.Return(Nil)
return
}


}


}


}, 1)

tmp11251 := Call(__e, ns2_1set, symshen_4extract_1free_1vars, tmp11215)


_ = tmp11251

tmp11252 := MakeNative(func(__e *ControlFlow) {
V1577 := __e.Get(1)
_ = V1577
V1578 := __e.Get(2)
_ = V1578
V1579 := __e.Get(3)
_ = V1579
V1580 := __e.Get(4)
_ = V1580
V1581 := __e.Get(5)
_ = V1581
tmp11287 := PrimEqual(Nil, V1577)

if True == tmp11287 {
tmp11253 := PrimCons(V1581, Nil)

__e.Return(PrimCons(symthaw, tmp11253))
return


} else {
tmp11285 := PrimIsPair(V1577)

var ifres11281 Obj

if True == tmp11285 {
tmp11283 := PrimHead(V1577)

tmp11284 := PrimEqual(sym_b, tmp11283)

var ifres11282 Obj

if True == tmp11284 {
ifres11282 = True


} else {
ifres11282 = False


}

ifres11281 = ifres11282


} else {
ifres11281 = False


}

if True == ifres11281 {
tmp11254 := PrimCons(symshen_4cut, Nil)

tmp11255 := PrimTail(V1577)

tmp11256 := PrimCons(tmp11254, tmp11255)

__e.TailApply(PrimFunc(symshen_4compile_1body), tmp11256, V1578, V1579, V1580, V1581)
return


} else {
tmp11279 := PrimIsPair(V1577)

var ifres11275 Obj

if True == tmp11279 {
tmp11277 := PrimTail(V1577)

tmp11278 := PrimEqual(Nil, tmp11277)

var ifres11276 Obj

if True == tmp11278 {
ifres11276 = True


} else {
ifres11276 = False


}

ifres11275 = ifres11276


} else {
ifres11275 = False


}

if True == ifres11275 {
tmp11257 := PrimHead(V1577)

tmp11258 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11257, V1578)


tmp11259 := PrimCons(V1581, Nil)

tmp11260 := PrimCons(V1580, tmp11259)

tmp11261 := PrimCons(V1579, tmp11260)

tmp11262 := PrimCons(V1578, tmp11261)

__e.TailApply(PrimFunc(symappend), tmp11258, tmp11262)
return


} else {
tmp11273 := PrimIsPair(V1577)

if True == tmp11273 {
tmp11263 := MakeNative(func(__e *ControlFlow) {
W1582 := __e.Get(1)
_ = W1582
tmp11264 := PrimTail(V1577)

tmp11265 := Call(__e, PrimFunc(symshen_4freeze_1literals), tmp11264, V1578, V1579, V1580, V1581)


tmp11266 := PrimCons(tmp11265, Nil)

tmp11267 := PrimCons(V1580, tmp11266)

tmp11268 := PrimCons(V1579, tmp11267)

tmp11269 := PrimCons(V1578, tmp11268)

__e.TailApply(PrimFunc(symappend), W1582, tmp11269)
return


}, 1)

tmp11270 := PrimHead(V1577)

tmp11271 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11270, V1578)


__e.TailApply(tmp11263, tmp11271)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-fbody")))
return
}


}


}


}


}, 5)

tmp11288 := Call(__e, ns2_1set, symshen_4compile_1body, tmp11252)


_ = tmp11288

tmp11289 := MakeNative(func(__e *ControlFlow) {
V1599 := __e.Get(1)
_ = V1599
V1600 := __e.Get(2)
_ = V1600
V1601 := __e.Get(3)
_ = V1601
V1602 := __e.Get(4)
_ = V1602
V1603 := __e.Get(5)
_ = V1603
tmp11313 := PrimEqual(Nil, V1599)

if True == tmp11313 {
__e.Return(V1603)
return
} else {
tmp11311 := PrimIsPair(V1599)

var ifres11307 Obj

if True == tmp11311 {
tmp11309 := PrimHead(V1599)

tmp11310 := PrimEqual(sym_b, tmp11309)

var ifres11308 Obj

if True == tmp11310 {
ifres11308 = True


} else {
ifres11308 = False


}

ifres11307 = ifres11308


} else {
ifres11307 = False


}

if True == ifres11307 {
tmp11290 := PrimCons(symshen_4cut, Nil)

tmp11291 := PrimTail(V1599)

tmp11292 := PrimCons(tmp11290, tmp11291)

__e.TailApply(PrimFunc(symshen_4freeze_1literals), tmp11292, V1600, V1601, V1602, V1603)
return


} else {
tmp11305 := PrimIsPair(V1599)

if True == tmp11305 {
tmp11293 := MakeNative(func(__e *ControlFlow) {
W1604 := __e.Get(1)
_ = W1604
tmp11294 := PrimTail(V1599)

tmp11295 := Call(__e, PrimFunc(symshen_4freeze_1literals), tmp11294, V1600, V1601, V1602, V1603)


tmp11296 := PrimCons(tmp11295, Nil)

tmp11297 := PrimCons(V1602, tmp11296)

tmp11298 := PrimCons(V1601, tmp11297)

tmp11299 := PrimCons(V1600, tmp11298)

tmp11300 := Call(__e, PrimFunc(symappend), W1604, tmp11299)


tmp11301 := PrimCons(tmp11300, Nil)

__e.Return(PrimCons(symfreeze, tmp11301))
return


}, 1)

tmp11302 := PrimHead(V1599)

tmp11303 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11302, V1600)


__e.TailApply(tmp11293, tmp11303)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.freeze-literals")))
return
}


}


}


}, 5)

tmp11314 := Call(__e, ns2_1set, symshen_4freeze_1literals, tmp11289)


_ = tmp11314

tmp11315 := MakeNative(func(__e *ControlFlow) {
V1609 := __e.Get(1)
_ = V1609
V1610 := __e.Get(2)
_ = V1610
tmp11330 := PrimIsPair(V1609)

var ifres11326 Obj

if True == tmp11330 {
tmp11328 := PrimHead(V1609)

tmp11329 := PrimEqual(symfork, tmp11328)

var ifres11327 Obj

if True == tmp11329 {
ifres11327 = True


} else {
ifres11327 = False


}

ifres11326 = ifres11327


} else {
ifres11326 = False


}

if True == ifres11326 {
tmp11316 := PrimTail(V1609)

tmp11317 := Call(__e, PrimFunc(symshen_4deref_1forked_1literals), tmp11316, V1610)


tmp11318 := PrimCons(tmp11317, Nil)

__e.Return(PrimCons(symfork, tmp11318))
return


} else {
tmp11324 := PrimIsPair(V1609)

if True == tmp11324 {
tmp11319 := PrimHead(V1609)

tmp11320 := MakeNative(func(__e *ControlFlow) {
Z1611 := __e.Get(1)
_ = Z1611
__e.TailApply(PrimFunc(symshen_4function_1calls), Z1611, V1610)
return
}, 1)

tmp11321 := PrimTail(V1609)

tmp11322 := Call(__e, PrimFunc(symmap), tmp11320, tmp11321)


__e.Return(PrimCons(tmp11319, tmp11322))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.deref-calls")))
return
}


}


}, 2)

tmp11331 := Call(__e, ns2_1set, symshen_4deref_1calls, tmp11315)


_ = tmp11331

tmp11332 := MakeNative(func(__e *ControlFlow) {
V1618 := __e.Get(1)
_ = V1618
V1619 := __e.Get(2)
_ = V1619
tmp11342 := PrimEqual(Nil, V1618)

if True == tmp11342 {
__e.Return(Nil)
return
} else {
tmp11340 := PrimIsPair(V1618)

if True == tmp11340 {
tmp11333 := PrimHead(V1618)

tmp11334 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11333, V1619)


tmp11335 := PrimTail(V1618)

tmp11336 := Call(__e, PrimFunc(symshen_4deref_1forked_1literals), tmp11335, V1619)


tmp11337 := PrimCons(tmp11336, Nil)

tmp11338 := PrimCons(tmp11334, tmp11337)

__e.Return(PrimCons(symcons, tmp11338))
return


} else {
__e.Return(PrimSimpleError(MakeString("fork requires a list of literals\n")))
return
}


}


}, 2)

tmp11343 := Call(__e, ns2_1set, symshen_4deref_1forked_1literals, tmp11332)


_ = tmp11343

tmp11344 := MakeNative(func(__e *ControlFlow) {
V1622 := __e.Get(1)
_ = V1622
V1623 := __e.Get(2)
_ = V1623
tmp11376 := PrimIsPair(V1622)

var ifres11357 Obj

if True == tmp11376 {
tmp11374 := PrimHead(V1622)

tmp11375 := PrimEqual(symcons, tmp11374)

var ifres11359 Obj

if True == tmp11375 {
tmp11372 := PrimTail(V1622)

tmp11373 := PrimIsPair(tmp11372)

var ifres11361 Obj

if True == tmp11373 {
tmp11369 := PrimTail(V1622)

tmp11370 := PrimTail(tmp11369)

tmp11371 := PrimIsPair(tmp11370)

var ifres11363 Obj

if True == tmp11371 {
tmp11365 := PrimTail(V1622)

tmp11366 := PrimTail(tmp11365)

tmp11367 := PrimTail(tmp11366)

tmp11368 := PrimEqual(Nil, tmp11367)

var ifres11364 Obj

if True == tmp11368 {
ifres11364 = True


} else {
ifres11364 = False


}

ifres11363 = ifres11364


} else {
ifres11363 = False


}

var ifres11362 Obj

if True == ifres11363 {
ifres11362 = True


} else {
ifres11362 = False


}

ifres11361 = ifres11362


} else {
ifres11361 = False


}

var ifres11360 Obj

if True == ifres11361 {
ifres11360 = True


} else {
ifres11360 = False


}

ifres11359 = ifres11360


} else {
ifres11359 = False


}

var ifres11358 Obj

if True == ifres11359 {
ifres11358 = True


} else {
ifres11358 = False


}

ifres11357 = ifres11358


} else {
ifres11357 = False


}

if True == ifres11357 {
tmp11345 := PrimTail(V1622)

tmp11346 := PrimHead(tmp11345)

tmp11347 := Call(__e, PrimFunc(symshen_4function_1calls), tmp11346, V1623)


tmp11348 := PrimTail(V1622)

tmp11349 := PrimTail(tmp11348)

tmp11350 := PrimHead(tmp11349)

tmp11351 := Call(__e, PrimFunc(symshen_4function_1calls), tmp11350, V1623)


tmp11352 := PrimCons(tmp11351, Nil)

tmp11353 := PrimCons(tmp11347, tmp11352)

__e.Return(PrimCons(symcons, tmp11353))
return


} else {
tmp11355 := PrimIsPair(V1622)

if True == tmp11355 {
__e.TailApply(PrimFunc(symshen_4deref_1terms), V1622, V1623, Nil)
return
} else {
__e.Return(V1622)
return
}


}


}, 2)

tmp11377 := Call(__e, ns2_1set, symshen_4function_1calls, tmp11344)


_ = tmp11377

tmp11378 := MakeNative(func(__e *ControlFlow) {
V1632 := __e.Get(1)
_ = V1632
V1633 := __e.Get(2)
_ = V1633
V1634 := __e.Get(3)
_ = V1634
tmp11472 := PrimIsPair(V1632)

var ifres11459 Obj

if True == tmp11472 {
tmp11470 := PrimHead(V1632)

tmp11471 := PrimEqual(MakeNumber(0), tmp11470)

var ifres11461 Obj

if True == tmp11471 {
tmp11468 := PrimTail(V1632)

tmp11469 := PrimIsPair(tmp11468)

var ifres11463 Obj

if True == tmp11469 {
tmp11465 := PrimTail(V1632)

tmp11466 := PrimTail(tmp11465)

tmp11467 := PrimEqual(Nil, tmp11466)

var ifres11464 Obj

if True == tmp11467 {
ifres11464 = True


} else {
ifres11464 = False


}

ifres11463 = ifres11464


} else {
ifres11463 = False


}

var ifres11462 Obj

if True == ifres11463 {
ifres11462 = True


} else {
ifres11462 = False


}

ifres11461 = ifres11462


} else {
ifres11461 = False


}

var ifres11460 Obj

if True == ifres11461 {
ifres11460 = True


} else {
ifres11460 = False


}

ifres11459 = ifres11460


} else {
ifres11459 = False


}

if True == ifres11459 {
tmp11385 := PrimTail(V1632)

tmp11386 := PrimHead(tmp11385)

tmp11387 := PrimIsVariable(tmp11386)

if True == tmp11387 {
tmp11379 := PrimTail(V1632)

__e.Return(PrimHead(tmp11379))
return


} else {
tmp11380 := PrimTail(V1632)

tmp11381 := PrimHead(tmp11380)

tmp11382 := Call(__e, PrimFunc(symshen_4app), tmp11381, MakeString("\n"), symshen_4s)


tmp11383 := PrimStringConcat(MakeString("attempt to optimise a non-variable "), tmp11382)

__e.Return(PrimSimpleError(tmp11383))
return


}


} else {
tmp11457 := PrimIsPair(V1632)

var ifres11444 Obj

if True == tmp11457 {
tmp11455 := PrimHead(V1632)

tmp11456 := PrimEqual(MakeNumber(1), tmp11455)

var ifres11446 Obj

if True == tmp11456 {
tmp11453 := PrimTail(V1632)

tmp11454 := PrimIsPair(tmp11453)

var ifres11448 Obj

if True == tmp11454 {
tmp11450 := PrimTail(V1632)

tmp11451 := PrimTail(tmp11450)

tmp11452 := PrimEqual(Nil, tmp11451)

var ifres11449 Obj

if True == tmp11452 {
ifres11449 = True


} else {
ifres11449 = False


}

ifres11448 = ifres11449


} else {
ifres11448 = False


}

var ifres11447 Obj

if True == ifres11448 {
ifres11447 = True


} else {
ifres11447 = False


}

ifres11446 = ifres11447


} else {
ifres11446 = False


}

var ifres11445 Obj

if True == ifres11446 {
ifres11445 = True


} else {
ifres11445 = False


}

ifres11444 = ifres11445


} else {
ifres11444 = False


}

if True == ifres11444 {
tmp11397 := PrimTail(V1632)

tmp11398 := PrimHead(tmp11397)

tmp11399 := PrimIsVariable(tmp11398)

if True == tmp11399 {
tmp11388 := PrimTail(V1632)

tmp11389 := PrimHead(tmp11388)

tmp11390 := PrimCons(V1633, Nil)

tmp11391 := PrimCons(tmp11389, tmp11390)

__e.Return(PrimCons(symshen_4lazyderef, tmp11391))
return


} else {
tmp11392 := PrimTail(V1632)

tmp11393 := PrimHead(tmp11392)

tmp11394 := Call(__e, PrimFunc(symshen_4app), tmp11393, MakeString("\n"), symshen_4s)


tmp11395 := PrimStringConcat(MakeString("attempt to optimise a non-variable "), tmp11394)

__e.Return(PrimSimpleError(tmp11395))
return


}


} else {
tmp11441 := Call(__e, PrimFunc(symelement_2), V1632, V1634)


tmp11442 := PrimNot(tmp11441)

var ifres11438 Obj

if True == tmp11442 {
tmp11440 := PrimIsVariable(V1632)

var ifres11439 Obj

if True == tmp11440 {
ifres11439 = True


} else {
ifres11439 = False


}

ifres11438 = ifres11439


} else {
ifres11438 = False


}

if True == ifres11438 {
tmp11400 := PrimCons(V1633, Nil)

tmp11401 := PrimCons(V1632, tmp11400)

__e.Return(PrimCons(symshen_4deref, tmp11401))
return


} else {
tmp11436 := PrimIsPair(V1632)

var ifres11417 Obj

if True == tmp11436 {
tmp11434 := PrimHead(V1632)

tmp11435 := PrimEqual(symlambda, tmp11434)

var ifres11419 Obj

if True == tmp11435 {
tmp11432 := PrimTail(V1632)

tmp11433 := PrimIsPair(tmp11432)

var ifres11421 Obj

if True == tmp11433 {
tmp11429 := PrimTail(V1632)

tmp11430 := PrimTail(tmp11429)

tmp11431 := PrimIsPair(tmp11430)

var ifres11423 Obj

if True == tmp11431 {
tmp11425 := PrimTail(V1632)

tmp11426 := PrimTail(tmp11425)

tmp11427 := PrimTail(tmp11426)

tmp11428 := PrimEqual(Nil, tmp11427)

var ifres11424 Obj

if True == tmp11428 {
ifres11424 = True


} else {
ifres11424 = False


}

ifres11423 = ifres11424


} else {
ifres11423 = False


}

var ifres11422 Obj

if True == ifres11423 {
ifres11422 = True


} else {
ifres11422 = False


}

ifres11421 = ifres11422


} else {
ifres11421 = False


}

var ifres11420 Obj

if True == ifres11421 {
ifres11420 = True


} else {
ifres11420 = False


}

ifres11419 = ifres11420


} else {
ifres11419 = False


}

var ifres11418 Obj

if True == ifres11419 {
ifres11418 = True


} else {
ifres11418 = False


}

ifres11417 = ifres11418


} else {
ifres11417 = False


}

if True == ifres11417 {
tmp11402 := PrimTail(V1632)

tmp11403 := PrimHead(tmp11402)

tmp11404 := PrimTail(V1632)

tmp11405 := PrimTail(tmp11404)

tmp11406 := PrimHead(tmp11405)

tmp11407 := PrimTail(V1632)

tmp11408 := PrimHead(tmp11407)

tmp11409 := PrimCons(tmp11408, V1634)

tmp11410 := Call(__e, PrimFunc(symshen_4deref_1terms), tmp11406, V1633, tmp11409)


tmp11411 := PrimCons(tmp11410, Nil)

tmp11412 := PrimCons(tmp11403, tmp11411)

__e.Return(PrimCons(symlambda, tmp11412))
return


} else {
tmp11415 := PrimIsPair(V1632)

if True == tmp11415 {
tmp11413 := MakeNative(func(__e *ControlFlow) {
Z1635 := __e.Get(1)
_ = Z1635
__e.TailApply(PrimFunc(symshen_4deref_1terms), Z1635, V1633, V1634)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp11413, V1632)
return


} else {
__e.Return(V1632)
return
}


}


}


}


}


}, 3)

tmp11473 := Call(__e, ns2_1set, symshen_4deref_1terms, tmp11378)


_ = tmp11473

tmp11474 := MakeNative(func(__e *ControlFlow) {
V1653 := __e.Get(1)
_ = V1653
V1654 := __e.Get(2)
_ = V1654
V1655 := __e.Get(3)
_ = V1655
V1656 := __e.Get(4)
_ = V1656
V1657 := __e.Get(5)
_ = V1657
tmp11650 := PrimEqual(Nil, V1654)

var ifres11647 Obj

if True == tmp11650 {
tmp11649 := PrimEqual(Nil, V1655)

var ifres11648 Obj

if True == tmp11649 {
ifres11648 = True


} else {
ifres11648 = False


}

ifres11647 = ifres11648


} else {
ifres11647 = False


}

if True == ifres11647 {
__e.Return(V1657)
return
} else {
tmp11645 := PrimIsPair(V1654)

var ifres11625 Obj

if True == tmp11645 {
tmp11643 := PrimHead(V1654)

tmp11644 := PrimIsPair(tmp11643)

var ifres11627 Obj

if True == tmp11644 {
tmp11640 := PrimHead(V1654)

tmp11641 := PrimHead(tmp11640)

tmp11642 := PrimEqual(symshen_4_7m, tmp11641)

var ifres11629 Obj

if True == tmp11642 {
tmp11637 := PrimHead(V1654)

tmp11638 := PrimTail(tmp11637)

tmp11639 := PrimIsPair(tmp11638)

var ifres11631 Obj

if True == tmp11639 {
tmp11633 := PrimHead(V1654)

tmp11634 := PrimTail(tmp11633)

tmp11635 := PrimTail(tmp11634)

tmp11636 := PrimEqual(Nil, tmp11635)

var ifres11632 Obj

if True == tmp11636 {
ifres11632 = True


} else {
ifres11632 = False


}

ifres11631 = ifres11632


} else {
ifres11631 = False


}

var ifres11630 Obj

if True == ifres11631 {
ifres11630 = True


} else {
ifres11630 = False


}

ifres11629 = ifres11630


} else {
ifres11629 = False


}

var ifres11628 Obj

if True == ifres11629 {
ifres11628 = True


} else {
ifres11628 = False


}

ifres11627 = ifres11628


} else {
ifres11627 = False


}

var ifres11626 Obj

if True == ifres11627 {
ifres11626 = True


} else {
ifres11626 = False


}

ifres11625 = ifres11626


} else {
ifres11625 = False


}

if True == ifres11625 {
tmp11475 := PrimHead(V1654)

tmp11476 := PrimTail(tmp11475)

tmp11477 := PrimHead(tmp11476)

tmp11478 := PrimTail(V1654)

tmp11479 := PrimCons(V1653, tmp11478)

tmp11480 := PrimCons(tmp11477, tmp11479)

tmp11481 := PrimCons(symshen_4_7m, tmp11480)

__e.TailApply(PrimFunc(symshen_4compile_1head), V1653, tmp11481, V1655, V1656, V1657)
return


} else {
tmp11623 := PrimIsPair(V1654)

var ifres11603 Obj

if True == tmp11623 {
tmp11621 := PrimHead(V1654)

tmp11622 := PrimIsPair(tmp11621)

var ifres11605 Obj

if True == tmp11622 {
tmp11618 := PrimHead(V1654)

tmp11619 := PrimHead(tmp11618)

tmp11620 := PrimEqual(symshen_4_1m, tmp11619)

var ifres11607 Obj

if True == tmp11620 {
tmp11615 := PrimHead(V1654)

tmp11616 := PrimTail(tmp11615)

tmp11617 := PrimIsPair(tmp11616)

var ifres11609 Obj

if True == tmp11617 {
tmp11611 := PrimHead(V1654)

tmp11612 := PrimTail(tmp11611)

tmp11613 := PrimTail(tmp11612)

tmp11614 := PrimEqual(Nil, tmp11613)

var ifres11610 Obj

if True == tmp11614 {
ifres11610 = True


} else {
ifres11610 = False


}

ifres11609 = ifres11610


} else {
ifres11609 = False


}

var ifres11608 Obj

if True == ifres11609 {
ifres11608 = True


} else {
ifres11608 = False


}

ifres11607 = ifres11608


} else {
ifres11607 = False


}

var ifres11606 Obj

if True == ifres11607 {
ifres11606 = True


} else {
ifres11606 = False


}

ifres11605 = ifres11606


} else {
ifres11605 = False


}

var ifres11604 Obj

if True == ifres11605 {
ifres11604 = True


} else {
ifres11604 = False


}

ifres11603 = ifres11604


} else {
ifres11603 = False


}

if True == ifres11603 {
tmp11482 := PrimHead(V1654)

tmp11483 := PrimTail(tmp11482)

tmp11484 := PrimHead(tmp11483)

tmp11485 := PrimTail(V1654)

tmp11486 := PrimCons(V1653, tmp11485)

tmp11487 := PrimCons(tmp11484, tmp11486)

tmp11488 := PrimCons(symshen_4_1m, tmp11487)

__e.TailApply(PrimFunc(symshen_4compile_1head), V1653, tmp11488, V1655, V1656, V1657)
return


} else {
tmp11601 := PrimIsPair(V1654)

var ifres11597 Obj

if True == tmp11601 {
tmp11599 := PrimHead(V1654)

tmp11600 := PrimEqual(symshen_4_1m, tmp11599)

var ifres11598 Obj

if True == tmp11600 {
ifres11598 = True


} else {
ifres11598 = False


}

ifres11597 = ifres11598


} else {
ifres11597 = False


}

if True == ifres11597 {
tmp11489 := PrimTail(V1654)

__e.TailApply(PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp11489, V1655, V1656, V1657)
return


} else {
tmp11595 := PrimIsPair(V1654)

var ifres11591 Obj

if True == tmp11595 {
tmp11593 := PrimHead(V1654)

tmp11594 := PrimEqual(symshen_4_7m, tmp11593)

var ifres11592 Obj

if True == tmp11594 {
ifres11592 = True


} else {
ifres11592 = False


}

ifres11591 = ifres11592


} else {
ifres11591 = False


}

if True == ifres11591 {
tmp11490 := PrimTail(V1654)

__e.TailApply(PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11490, V1655, V1656, V1657)
return


} else {
tmp11589 := PrimIsPair(V1654)

var ifres11582 Obj

if True == tmp11589 {
tmp11588 := PrimIsPair(V1655)

var ifres11584 Obj

if True == tmp11588 {
tmp11586 := PrimHead(V1654)

tmp11587 := Call(__e, PrimFunc(symshen_4wildcard_2), tmp11586)


var ifres11585 Obj

if True == tmp11587 {
ifres11585 = True


} else {
ifres11585 = False


}

ifres11584 = ifres11585


} else {
ifres11584 = False


}

var ifres11583 Obj

if True == ifres11584 {
ifres11583 = True


} else {
ifres11583 = False


}

ifres11582 = ifres11583


} else {
ifres11582 = False


}

if True == ifres11582 {
tmp11491 := PrimTail(V1654)

tmp11492 := PrimTail(V1655)

__e.TailApply(PrimFunc(symshen_4compile_1head), V1653, tmp11491, tmp11492, V1656, V1657)
return


} else {
tmp11580 := PrimIsPair(V1654)

var ifres11576 Obj

if True == tmp11580 {
tmp11578 := PrimHead(V1654)

tmp11579 := PrimIsVariable(tmp11578)

var ifres11577 Obj

if True == tmp11579 {
ifres11577 = True


} else {
ifres11577 = False


}

ifres11576 = ifres11577


} else {
ifres11576 = False


}

if True == ifres11576 {
__e.TailApply(PrimFunc(symshen_4variable_1case), V1653, V1654, V1655, V1656, V1657)
return
} else {
tmp11574 := PrimEqual(symshen_4_1m, V1653)

var ifres11567 Obj

if True == tmp11574 {
tmp11573 := PrimIsPair(V1654)

var ifres11569 Obj

if True == tmp11573 {
tmp11571 := PrimHead(V1654)

tmp11572 := Call(__e, PrimFunc(symatom_2), tmp11571)


var ifres11570 Obj

if True == tmp11572 {
ifres11570 = True


} else {
ifres11570 = False


}

ifres11569 = ifres11570


} else {
ifres11569 = False


}

var ifres11568 Obj

if True == ifres11569 {
ifres11568 = True


} else {
ifres11568 = False


}

ifres11567 = ifres11568


} else {
ifres11567 = False


}

if True == ifres11567 {
__e.TailApply(PrimFunc(symshen_4atom_1case_1minus), V1654, V1655, V1656, V1657)
return
} else {
tmp11565 := PrimEqual(symshen_4_1m, V1653)

var ifres11535 Obj

if True == tmp11565 {
tmp11564 := PrimIsPair(V1654)

var ifres11537 Obj

if True == tmp11564 {
tmp11562 := PrimHead(V1654)

tmp11563 := PrimIsPair(tmp11562)

var ifres11539 Obj

if True == tmp11563 {
tmp11559 := PrimHead(V1654)

tmp11560 := PrimHead(tmp11559)

tmp11561 := PrimEqual(symcons, tmp11560)

var ifres11541 Obj

if True == tmp11561 {
tmp11556 := PrimHead(V1654)

tmp11557 := PrimTail(tmp11556)

tmp11558 := PrimIsPair(tmp11557)

var ifres11543 Obj

if True == tmp11558 {
tmp11552 := PrimHead(V1654)

tmp11553 := PrimTail(tmp11552)

tmp11554 := PrimTail(tmp11553)

tmp11555 := PrimIsPair(tmp11554)

var ifres11545 Obj

if True == tmp11555 {
tmp11547 := PrimHead(V1654)

tmp11548 := PrimTail(tmp11547)

tmp11549 := PrimTail(tmp11548)

tmp11550 := PrimTail(tmp11549)

tmp11551 := PrimEqual(Nil, tmp11550)

var ifres11546 Obj

if True == tmp11551 {
ifres11546 = True


} else {
ifres11546 = False


}

ifres11545 = ifres11546


} else {
ifres11545 = False


}

var ifres11544 Obj

if True == ifres11545 {
ifres11544 = True


} else {
ifres11544 = False


}

ifres11543 = ifres11544


} else {
ifres11543 = False


}

var ifres11542 Obj

if True == ifres11543 {
ifres11542 = True


} else {
ifres11542 = False


}

ifres11541 = ifres11542


} else {
ifres11541 = False


}

var ifres11540 Obj

if True == ifres11541 {
ifres11540 = True


} else {
ifres11540 = False


}

ifres11539 = ifres11540


} else {
ifres11539 = False


}

var ifres11538 Obj

if True == ifres11539 {
ifres11538 = True


} else {
ifres11538 = False


}

ifres11537 = ifres11538


} else {
ifres11537 = False


}

var ifres11536 Obj

if True == ifres11537 {
ifres11536 = True


} else {
ifres11536 = False


}

ifres11535 = ifres11536


} else {
ifres11535 = False


}

if True == ifres11535 {
__e.TailApply(PrimFunc(symshen_4cons_1case_1minus), V1654, V1655, V1656, V1657)
return
} else {
tmp11533 := PrimEqual(symshen_4_7m, V1653)

var ifres11526 Obj

if True == tmp11533 {
tmp11532 := PrimIsPair(V1654)

var ifres11528 Obj

if True == tmp11532 {
tmp11530 := PrimHead(V1654)

tmp11531 := Call(__e, PrimFunc(symatom_2), tmp11530)


var ifres11529 Obj

if True == tmp11531 {
ifres11529 = True


} else {
ifres11529 = False


}

ifres11528 = ifres11529


} else {
ifres11528 = False


}

var ifres11527 Obj

if True == ifres11528 {
ifres11527 = True


} else {
ifres11527 = False


}

ifres11526 = ifres11527


} else {
ifres11526 = False


}

if True == ifres11526 {
__e.TailApply(PrimFunc(symshen_4atom_1case_1plus), V1654, V1655, V1656, V1657)
return
} else {
tmp11524 := PrimEqual(symshen_4_7m, V1653)

var ifres11494 Obj

if True == tmp11524 {
tmp11523 := PrimIsPair(V1654)

var ifres11496 Obj

if True == tmp11523 {
tmp11521 := PrimHead(V1654)

tmp11522 := PrimIsPair(tmp11521)

var ifres11498 Obj

if True == tmp11522 {
tmp11518 := PrimHead(V1654)

tmp11519 := PrimHead(tmp11518)

tmp11520 := PrimEqual(symcons, tmp11519)

var ifres11500 Obj

if True == tmp11520 {
tmp11515 := PrimHead(V1654)

tmp11516 := PrimTail(tmp11515)

tmp11517 := PrimIsPair(tmp11516)

var ifres11502 Obj

if True == tmp11517 {
tmp11511 := PrimHead(V1654)

tmp11512 := PrimTail(tmp11511)

tmp11513 := PrimTail(tmp11512)

tmp11514 := PrimIsPair(tmp11513)

var ifres11504 Obj

if True == tmp11514 {
tmp11506 := PrimHead(V1654)

tmp11507 := PrimTail(tmp11506)

tmp11508 := PrimTail(tmp11507)

tmp11509 := PrimTail(tmp11508)

tmp11510 := PrimEqual(Nil, tmp11509)

var ifres11505 Obj

if True == tmp11510 {
ifres11505 = True


} else {
ifres11505 = False


}

ifres11504 = ifres11505


} else {
ifres11504 = False


}

var ifres11503 Obj

if True == ifres11504 {
ifres11503 = True


} else {
ifres11503 = False


}

ifres11502 = ifres11503


} else {
ifres11502 = False


}

var ifres11501 Obj

if True == ifres11502 {
ifres11501 = True


} else {
ifres11501 = False


}

ifres11500 = ifres11501


} else {
ifres11500 = False


}

var ifres11499 Obj

if True == ifres11500 {
ifres11499 = True


} else {
ifres11499 = False


}

ifres11498 = ifres11499


} else {
ifres11498 = False


}

var ifres11497 Obj

if True == ifres11498 {
ifres11497 = True


} else {
ifres11497 = False


}

ifres11496 = ifres11497


} else {
ifres11496 = False


}

var ifres11495 Obj

if True == ifres11496 {
ifres11495 = True


} else {
ifres11495 = False


}

ifres11494 = ifres11495


} else {
ifres11494 = False


}

if True == ifres11494 {
__e.TailApply(PrimFunc(symshen_4cons_1case_1plus), V1654, V1655, V1656, V1657)
return
} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-head")))
return
}


}


}


}


}


}


}


}


}


}


}


}, 5)

tmp11651 := Call(__e, ns2_1set, symshen_4compile_1head, tmp11474)


_ = tmp11651

tmp11652 := MakeNative(func(__e *ControlFlow) {
V1668 := __e.Get(1)
_ = V1668
V1669 := __e.Get(2)
_ = V1669
V1670 := __e.Get(3)
_ = V1670
V1671 := __e.Get(4)
_ = V1671
V1672 := __e.Get(5)
_ = V1672
tmp11673 := PrimIsPair(V1669)

var ifres11670 Obj

if True == tmp11673 {
tmp11672 := PrimIsPair(V1670)

var ifres11671 Obj

if True == tmp11672 {
ifres11671 = True


} else {
ifres11671 = False


}

ifres11670 = ifres11671


} else {
ifres11670 = False


}

if True == ifres11670 {
tmp11667 := PrimHead(V1670)

tmp11668 := PrimIsVariable(tmp11667)

if True == tmp11668 {
tmp11653 := PrimTail(V1669)

tmp11654 := PrimTail(V1670)

tmp11655 := PrimHead(V1670)

tmp11656 := PrimHead(V1669)

tmp11657 := Call(__e, PrimFunc(symsubst), tmp11655, tmp11656, V1672)


__e.TailApply(PrimFunc(symshen_4compile_1head), V1668, tmp11653, tmp11654, V1671, tmp11657)
return


} else {
tmp11658 := PrimHead(V1669)

tmp11659 := PrimHead(V1670)

tmp11660 := PrimTail(V1669)

tmp11661 := PrimTail(V1670)

tmp11662 := Call(__e, PrimFunc(symshen_4compile_1head), V1668, tmp11660, tmp11661, V1671, V1672)


tmp11663 := PrimCons(tmp11662, Nil)

tmp11664 := PrimCons(tmp11659, tmp11663)

tmp11665 := PrimCons(tmp11658, tmp11664)

__e.Return(PrimCons(symlet, tmp11665))
return


}


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.variable-case")))
return
}


}, 5)

tmp11674 := Call(__e, ns2_1set, symshen_4variable_1case, tmp11652)


_ = tmp11674

tmp11675 := MakeNative(func(__e *ControlFlow) {
V1681 := __e.Get(1)
_ = V1681
V1682 := __e.Get(2)
_ = V1682
V1683 := __e.Get(3)
_ = V1683
V1684 := __e.Get(4)
_ = V1684
tmp11700 := PrimIsPair(V1681)

var ifres11697 Obj

if True == tmp11700 {
tmp11699 := PrimIsPair(V1682)

var ifres11698 Obj

if True == tmp11699 {
ifres11698 = True


} else {
ifres11698 = False


}

ifres11697 = ifres11698


} else {
ifres11697 = False


}

if True == ifres11697 {
tmp11676 := MakeNative(func(__e *ControlFlow) {
W1685 := __e.Get(1)
_ = W1685
tmp11677 := PrimHead(V1682)

tmp11678 := PrimCons(V1683, Nil)

tmp11679 := PrimCons(tmp11677, tmp11678)

tmp11680 := PrimCons(symshen_4lazyderef, tmp11679)

tmp11681 := PrimHead(V1681)

tmp11682 := PrimCons(tmp11681, Nil)

tmp11683 := PrimCons(W1685, tmp11682)

tmp11684 := PrimCons(sym_a, tmp11683)

tmp11685 := PrimTail(V1681)

tmp11686 := PrimTail(V1682)

tmp11687 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp11685, tmp11686, V1683, V1684)


tmp11688 := PrimCons(False, Nil)

tmp11689 := PrimCons(tmp11687, tmp11688)

tmp11690 := PrimCons(tmp11684, tmp11689)

tmp11691 := PrimCons(symif, tmp11690)

tmp11692 := PrimCons(tmp11691, Nil)

tmp11693 := PrimCons(tmp11680, tmp11692)

tmp11694 := PrimCons(W1685, tmp11693)

__e.Return(PrimCons(symlet, tmp11694))
return


}, 1)

tmp11695 := Call(__e, PrimFunc(symgensym), symTm)


__e.TailApply(tmp11676, tmp11695)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.atom-case-minus")))
return
}


}, 4)

tmp11701 := Call(__e, ns2_1set, symshen_4atom_1case_1minus, tmp11675)


_ = tmp11701

tmp11702 := MakeNative(func(__e *ControlFlow) {
V1694 := __e.Get(1)
_ = V1694
V1695 := __e.Get(2)
_ = V1695
V1696 := __e.Get(3)
_ = V1696
V1697 := __e.Get(4)
_ = V1697
tmp11767 := PrimIsPair(V1694)

var ifres11737 Obj

if True == tmp11767 {
tmp11765 := PrimHead(V1694)

tmp11766 := PrimIsPair(tmp11765)

var ifres11739 Obj

if True == tmp11766 {
tmp11762 := PrimHead(V1694)

tmp11763 := PrimHead(tmp11762)

tmp11764 := PrimEqual(symcons, tmp11763)

var ifres11741 Obj

if True == tmp11764 {
tmp11759 := PrimHead(V1694)

tmp11760 := PrimTail(tmp11759)

tmp11761 := PrimIsPair(tmp11760)

var ifres11743 Obj

if True == tmp11761 {
tmp11755 := PrimHead(V1694)

tmp11756 := PrimTail(tmp11755)

tmp11757 := PrimTail(tmp11756)

tmp11758 := PrimIsPair(tmp11757)

var ifres11745 Obj

if True == tmp11758 {
tmp11750 := PrimHead(V1694)

tmp11751 := PrimTail(tmp11750)

tmp11752 := PrimTail(tmp11751)

tmp11753 := PrimTail(tmp11752)

tmp11754 := PrimEqual(Nil, tmp11753)

var ifres11747 Obj

if True == tmp11754 {
tmp11749 := PrimIsPair(V1695)

var ifres11748 Obj

if True == tmp11749 {
ifres11748 = True


} else {
ifres11748 = False


}

ifres11747 = ifres11748


} else {
ifres11747 = False


}

var ifres11746 Obj

if True == ifres11747 {
ifres11746 = True


} else {
ifres11746 = False


}

ifres11745 = ifres11746


} else {
ifres11745 = False


}

var ifres11744 Obj

if True == ifres11745 {
ifres11744 = True


} else {
ifres11744 = False


}

ifres11743 = ifres11744


} else {
ifres11743 = False


}

var ifres11742 Obj

if True == ifres11743 {
ifres11742 = True


} else {
ifres11742 = False


}

ifres11741 = ifres11742


} else {
ifres11741 = False


}

var ifres11740 Obj

if True == ifres11741 {
ifres11740 = True


} else {
ifres11740 = False


}

ifres11739 = ifres11740


} else {
ifres11739 = False


}

var ifres11738 Obj

if True == ifres11739 {
ifres11738 = True


} else {
ifres11738 = False


}

ifres11737 = ifres11738


} else {
ifres11737 = False


}

if True == ifres11737 {
tmp11703 := MakeNative(func(__e *ControlFlow) {
W1698 := __e.Get(1)
_ = W1698
tmp11704 := PrimHead(V1695)

tmp11705 := PrimCons(V1696, Nil)

tmp11706 := PrimCons(tmp11704, tmp11705)

tmp11707 := PrimCons(symshen_4lazyderef, tmp11706)

tmp11708 := PrimCons(W1698, Nil)

tmp11709 := PrimCons(symcons_2, tmp11708)

tmp11710 := PrimHead(V1694)

tmp11711 := PrimTail(tmp11710)

tmp11712 := PrimHead(tmp11711)

tmp11713 := PrimHead(V1694)

tmp11714 := PrimTail(tmp11713)

tmp11715 := PrimTail(tmp11714)

tmp11716 := PrimHead(tmp11715)

tmp11717 := PrimTail(V1694)

tmp11718 := PrimCons(tmp11716, tmp11717)

tmp11719 := PrimCons(tmp11712, tmp11718)

tmp11720 := PrimCons(W1698, Nil)

tmp11721 := PrimCons(symhd, tmp11720)

tmp11722 := PrimCons(W1698, Nil)

tmp11723 := PrimCons(symtl, tmp11722)

tmp11724 := PrimTail(V1695)

tmp11725 := PrimCons(tmp11723, tmp11724)

tmp11726 := PrimCons(tmp11721, tmp11725)

tmp11727 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp11719, tmp11726, V1696, V1697)


tmp11728 := PrimCons(False, Nil)

tmp11729 := PrimCons(tmp11727, tmp11728)

tmp11730 := PrimCons(tmp11709, tmp11729)

tmp11731 := PrimCons(symif, tmp11730)

tmp11732 := PrimCons(tmp11731, Nil)

tmp11733 := PrimCons(tmp11707, tmp11732)

tmp11734 := PrimCons(W1698, tmp11733)

__e.Return(PrimCons(symlet, tmp11734))
return


}, 1)

tmp11735 := Call(__e, PrimFunc(symgensym), symTm)


__e.TailApply(tmp11703, tmp11735)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.cons-case-minus")))
return
}


}, 4)

tmp11768 := Call(__e, ns2_1set, symshen_4cons_1case_1minus, tmp11702)


_ = tmp11768

tmp11769 := MakeNative(func(__e *ControlFlow) {
V1707 := __e.Get(1)
_ = V1707
V1708 := __e.Get(2)
_ = V1708
V1709 := __e.Get(3)
_ = V1709
V1710 := __e.Get(4)
_ = V1710
tmp11815 := PrimIsPair(V1707)

var ifres11812 Obj

if True == tmp11815 {
tmp11814 := PrimIsPair(V1708)

var ifres11813 Obj

if True == tmp11814 {
ifres11813 = True


} else {
ifres11813 = False


}

ifres11812 = ifres11813


} else {
ifres11812 = False


}

if True == ifres11812 {
tmp11770 := MakeNative(func(__e *ControlFlow) {
W1711 := __e.Get(1)
_ = W1711
tmp11771 := MakeNative(func(__e *ControlFlow) {
W1712 := __e.Get(1)
_ = W1712
tmp11772 := PrimHead(V1708)

tmp11773 := PrimCons(V1709, Nil)

tmp11774 := PrimCons(tmp11772, tmp11773)

tmp11775 := PrimCons(symshen_4lazyderef, tmp11774)

tmp11776 := PrimTail(V1707)

tmp11777 := PrimTail(V1708)

tmp11778 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11776, tmp11777, V1709, V1710)


tmp11779 := PrimCons(tmp11778, Nil)

tmp11780 := PrimCons(symfreeze, tmp11779)

tmp11781 := PrimHead(V1707)

tmp11782 := PrimCons(tmp11781, Nil)

tmp11783 := PrimCons(W1711, tmp11782)

tmp11784 := PrimCons(sym_a, tmp11783)

tmp11785 := PrimCons(W1712, Nil)

tmp11786 := PrimCons(symthaw, tmp11785)

tmp11787 := PrimCons(W1711, Nil)

tmp11788 := PrimCons(symshen_4pvar_2, tmp11787)

tmp11789 := PrimHead(V1707)

tmp11790 := Call(__e, PrimFunc(symshen_4demode), tmp11789)


tmp11791 := PrimCons(W1712, Nil)

tmp11792 := PrimCons(V1709, tmp11791)

tmp11793 := PrimCons(tmp11790, tmp11792)

tmp11794 := PrimCons(W1711, tmp11793)

tmp11795 := PrimCons(symshen_4bind_b, tmp11794)

tmp11796 := PrimCons(False, Nil)

tmp11797 := PrimCons(tmp11795, tmp11796)

tmp11798 := PrimCons(tmp11788, tmp11797)

tmp11799 := PrimCons(symif, tmp11798)

tmp11800 := PrimCons(tmp11799, Nil)

tmp11801 := PrimCons(tmp11786, tmp11800)

tmp11802 := PrimCons(tmp11784, tmp11801)

tmp11803 := PrimCons(symif, tmp11802)

tmp11804 := PrimCons(tmp11803, Nil)

tmp11805 := PrimCons(tmp11780, tmp11804)

tmp11806 := PrimCons(W1712, tmp11805)

tmp11807 := PrimCons(tmp11775, tmp11806)

tmp11808 := PrimCons(W1711, tmp11807)

__e.Return(PrimCons(symlet, tmp11808))
return


}, 1)

tmp11809 := Call(__e, PrimFunc(symgensym), symGoTo)


__e.TailApply(tmp11771, tmp11809)
return


}, 1)

tmp11810 := Call(__e, PrimFunc(symgensym), symTm)


__e.TailApply(tmp11770, tmp11810)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.atom-case-plus")))
return
}


}, 4)

tmp11816 := Call(__e, ns2_1set, symshen_4atom_1case_1plus, tmp11769)


_ = tmp11816

tmp11817 := MakeNative(func(__e *ControlFlow) {
V1721 := __e.Get(1)
_ = V1721
V1722 := __e.Get(2)
_ = V1722
V1723 := __e.Get(3)
_ = V1723
V1724 := __e.Get(4)
_ = V1724
tmp11913 := PrimIsPair(V1721)

var ifres11883 Obj

if True == tmp11913 {
tmp11911 := PrimHead(V1721)

tmp11912 := PrimIsPair(tmp11911)

var ifres11885 Obj

if True == tmp11912 {
tmp11908 := PrimHead(V1721)

tmp11909 := PrimHead(tmp11908)

tmp11910 := PrimEqual(symcons, tmp11909)

var ifres11887 Obj

if True == tmp11910 {
tmp11905 := PrimHead(V1721)

tmp11906 := PrimTail(tmp11905)

tmp11907 := PrimIsPair(tmp11906)

var ifres11889 Obj

if True == tmp11907 {
tmp11901 := PrimHead(V1721)

tmp11902 := PrimTail(tmp11901)

tmp11903 := PrimTail(tmp11902)

tmp11904 := PrimIsPair(tmp11903)

var ifres11891 Obj

if True == tmp11904 {
tmp11896 := PrimHead(V1721)

tmp11897 := PrimTail(tmp11896)

tmp11898 := PrimTail(tmp11897)

tmp11899 := PrimTail(tmp11898)

tmp11900 := PrimEqual(Nil, tmp11899)

var ifres11893 Obj

if True == tmp11900 {
tmp11895 := PrimIsPair(V1722)

var ifres11894 Obj

if True == tmp11895 {
ifres11894 = True


} else {
ifres11894 = False


}

ifres11893 = ifres11894


} else {
ifres11893 = False


}

var ifres11892 Obj

if True == ifres11893 {
ifres11892 = True


} else {
ifres11892 = False


}

ifres11891 = ifres11892


} else {
ifres11891 = False


}

var ifres11890 Obj

if True == ifres11891 {
ifres11890 = True


} else {
ifres11890 = False


}

ifres11889 = ifres11890


} else {
ifres11889 = False


}

var ifres11888 Obj

if True == ifres11889 {
ifres11888 = True


} else {
ifres11888 = False


}

ifres11887 = ifres11888


} else {
ifres11887 = False


}

var ifres11886 Obj

if True == ifres11887 {
ifres11886 = True


} else {
ifres11886 = False


}

ifres11885 = ifres11886


} else {
ifres11885 = False


}

var ifres11884 Obj

if True == ifres11885 {
ifres11884 = True


} else {
ifres11884 = False


}

ifres11883 = ifres11884


} else {
ifres11883 = False


}

if True == ifres11883 {
tmp11818 := MakeNative(func(__e *ControlFlow) {
W1725 := __e.Get(1)
_ = W1725
tmp11819 := MakeNative(func(__e *ControlFlow) {
W1726 := __e.Get(1)
_ = W1726
tmp11820 := MakeNative(func(__e *ControlFlow) {
W1727 := __e.Get(1)
_ = W1727
tmp11821 := MakeNative(func(__e *ControlFlow) {
W1728 := __e.Get(1)
_ = W1728
tmp11822 := MakeNative(func(__e *ControlFlow) {
W1729 := __e.Get(1)
_ = W1729
tmp11823 := PrimHead(V1722)

tmp11824 := PrimCons(V1723, Nil)

tmp11825 := PrimCons(tmp11823, tmp11824)

tmp11826 := PrimCons(symshen_4lazyderef, tmp11825)

tmp11827 := PrimTail(V1721)

tmp11828 := PrimTail(V1722)

tmp11829 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11827, tmp11828, V1723, V1724)


tmp11830 := Call(__e, PrimFunc(symshen_4goto), W1727, tmp11829)


tmp11831 := PrimCons(W1725, Nil)

tmp11832 := PrimCons(symcons_2, tmp11831)

tmp11833 := PrimHead(V1721)

tmp11834 := PrimTail(tmp11833)

tmp11835 := PrimCons(W1725, Nil)

tmp11836 := PrimCons(symhd, tmp11835)

tmp11837 := PrimCons(W1725, Nil)

tmp11838 := PrimCons(symtl, tmp11837)

tmp11839 := PrimCons(tmp11838, Nil)

tmp11840 := PrimCons(tmp11836, tmp11839)

tmp11841 := Call(__e, PrimFunc(symshen_4invoke), W1726, W1727)


tmp11842 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11834, tmp11840, V1723, tmp11841)


tmp11843 := PrimCons(W1725, Nil)

tmp11844 := PrimCons(symshen_4pvar_2, tmp11843)

tmp11845 := Call(__e, PrimFunc(symshen_4demode), W1728)


tmp11846 := Call(__e, PrimFunc(symshen_4invoke), W1726, W1727)


tmp11847 := PrimCons(tmp11846, Nil)

tmp11848 := PrimCons(symfreeze, tmp11847)

tmp11849 := PrimCons(tmp11848, Nil)

tmp11850 := PrimCons(V1723, tmp11849)

tmp11851 := PrimCons(tmp11845, tmp11850)

tmp11852 := PrimCons(W1725, tmp11851)

tmp11853 := PrimCons(symshen_4bind_b, tmp11852)

tmp11854 := Call(__e, PrimFunc(symshen_4stpart), W1729, tmp11853, V1723)


tmp11855 := PrimCons(False, Nil)

tmp11856 := PrimCons(tmp11854, tmp11855)

tmp11857 := PrimCons(tmp11844, tmp11856)

tmp11858 := PrimCons(symif, tmp11857)

tmp11859 := PrimCons(tmp11858, Nil)

tmp11860 := PrimCons(tmp11842, tmp11859)

tmp11861 := PrimCons(tmp11832, tmp11860)

tmp11862 := PrimCons(symif, tmp11861)

tmp11863 := PrimCons(tmp11862, Nil)

tmp11864 := PrimCons(tmp11830, tmp11863)

tmp11865 := PrimCons(W1726, tmp11864)

tmp11866 := PrimCons(tmp11826, tmp11865)

tmp11867 := PrimCons(W1725, tmp11866)

__e.Return(PrimCons(symlet, tmp11867))
return


}, 1)

tmp11868 := Call(__e, PrimFunc(symshen_4extract_1vars), W1728)


__e.TailApply(tmp11822, tmp11868)
return


}, 1)

tmp11869 := PrimHead(V1721)

tmp11870 := Call(__e, PrimFunc(symshen_4tame), tmp11869)


__e.TailApply(tmp11821, tmp11870)
return


}, 1)

tmp11871 := PrimHead(V1721)

tmp11872 := PrimTail(tmp11871)

tmp11873 := PrimHead(tmp11872)

tmp11874 := PrimHead(V1721)

tmp11875 := PrimTail(tmp11874)

tmp11876 := PrimTail(tmp11875)

tmp11877 := PrimHead(tmp11876)

tmp11878 := PrimCons(tmp11873, tmp11877)

tmp11879 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp11878)


__e.TailApply(tmp11820, tmp11879)
return


}, 1)

tmp11880 := Call(__e, PrimFunc(symgensym), symGoTo)


__e.TailApply(tmp11819, tmp11880)
return


}, 1)

tmp11881 := Call(__e, PrimFunc(symgensym), symTm)


__e.TailApply(tmp11818, tmp11881)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.cons-case-plus")))
return
}


}, 4)

tmp11914 := Call(__e, ns2_1set, symshen_4cons_1case_1plus, tmp11817)


_ = tmp11914

tmp11915 := MakeNative(func(__e *ControlFlow) {
V1730 := __e.Get(1)
_ = V1730
tmp11952 := PrimIsPair(V1730)

var ifres11939 Obj

if True == tmp11952 {
tmp11950 := PrimHead(V1730)

tmp11951 := PrimEqual(symshen_4_7m, tmp11950)

var ifres11941 Obj

if True == tmp11951 {
tmp11948 := PrimTail(V1730)

tmp11949 := PrimIsPair(tmp11948)

var ifres11943 Obj

if True == tmp11949 {
tmp11945 := PrimTail(V1730)

tmp11946 := PrimTail(tmp11945)

tmp11947 := PrimEqual(Nil, tmp11946)

var ifres11944 Obj

if True == tmp11947 {
ifres11944 = True


} else {
ifres11944 = False


}

ifres11943 = ifres11944


} else {
ifres11943 = False


}

var ifres11942 Obj

if True == ifres11943 {
ifres11942 = True


} else {
ifres11942 = False


}

ifres11941 = ifres11942


} else {
ifres11941 = False


}

var ifres11940 Obj

if True == ifres11941 {
ifres11940 = True


} else {
ifres11940 = False


}

ifres11939 = ifres11940


} else {
ifres11939 = False


}

if True == ifres11939 {
tmp11916 := PrimTail(V1730)

tmp11917 := PrimHead(tmp11916)

__e.TailApply(PrimFunc(symshen_4demode), tmp11917)
return


} else {
tmp11937 := PrimIsPair(V1730)

var ifres11924 Obj

if True == tmp11937 {
tmp11935 := PrimHead(V1730)

tmp11936 := PrimEqual(symshen_4_1m, tmp11935)

var ifres11926 Obj

if True == tmp11936 {
tmp11933 := PrimTail(V1730)

tmp11934 := PrimIsPair(tmp11933)

var ifres11928 Obj

if True == tmp11934 {
tmp11930 := PrimTail(V1730)

tmp11931 := PrimTail(tmp11930)

tmp11932 := PrimEqual(Nil, tmp11931)

var ifres11929 Obj

if True == tmp11932 {
ifres11929 = True


} else {
ifres11929 = False


}

ifres11928 = ifres11929


} else {
ifres11928 = False


}

var ifres11927 Obj

if True == ifres11928 {
ifres11927 = True


} else {
ifres11927 = False


}

ifres11926 = ifres11927


} else {
ifres11926 = False


}

var ifres11925 Obj

if True == ifres11926 {
ifres11925 = True


} else {
ifres11925 = False


}

ifres11924 = ifres11925


} else {
ifres11924 = False


}

if True == ifres11924 {
tmp11918 := PrimTail(V1730)

tmp11919 := PrimHead(tmp11918)

__e.TailApply(PrimFunc(symshen_4demode), tmp11919)
return


} else {
tmp11922 := PrimIsPair(V1730)

if True == tmp11922 {
tmp11920 := MakeNative(func(__e *ControlFlow) {
Z1731 := __e.Get(1)
_ = Z1731
__e.TailApply(PrimFunc(symshen_4demode), Z1731)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp11920, V1730)
return


} else {
__e.Return(V1730)
return
}


}


}


}, 1)

tmp11953 := Call(__e, ns2_1set, symshen_4demode, tmp11915)


_ = tmp11953

tmp11954 := MakeNative(func(__e *ControlFlow) {
V1732 := __e.Get(1)
_ = V1732
tmp11959 := Call(__e, PrimFunc(symshen_4wildcard_2), V1732)


if True == tmp11959 {
__e.TailApply(PrimFunc(symgensym), symY)
return
} else {
tmp11957 := PrimIsPair(V1732)

if True == tmp11957 {
tmp11955 := MakeNative(func(__e *ControlFlow) {
Z1733 := __e.Get(1)
_ = Z1733
__e.TailApply(PrimFunc(symshen_4tame), Z1733)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp11955, V1732)
return


} else {
__e.Return(V1732)
return
}


}


}, 1)

tmp11960 := Call(__e, ns2_1set, symshen_4tame, tmp11954)


_ = tmp11960

tmp11961 := MakeNative(func(__e *ControlFlow) {
V1734 := __e.Get(1)
_ = V1734
V1735 := __e.Get(2)
_ = V1735
tmp11964 := PrimEqual(Nil, V1734)

if True == tmp11964 {
tmp11962 := PrimCons(V1735, Nil)

__e.Return(PrimCons(symfreeze, tmp11962))
return


} else {
__e.TailApply(PrimFunc(symshen_4goto_1h), V1734, V1735)
return
}


}, 2)

tmp11965 := Call(__e, ns2_1set, symshen_4goto, tmp11961)


_ = tmp11965

tmp11966 := MakeNative(func(__e *ControlFlow) {
V1736 := __e.Get(1)
_ = V1736
V1737 := __e.Get(2)
_ = V1737
tmp11975 := PrimEqual(Nil, V1736)

if True == tmp11975 {
__e.Return(V1737)
return
} else {
tmp11973 := PrimIsPair(V1736)

if True == tmp11973 {
tmp11967 := PrimHead(V1736)

tmp11968 := PrimTail(V1736)

tmp11969 := Call(__e, PrimFunc(symshen_4goto_1h), tmp11968, V1737)


tmp11970 := PrimCons(tmp11969, Nil)

tmp11971 := PrimCons(tmp11967, tmp11970)

__e.Return(PrimCons(symlambda, tmp11971))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.goto-h")))
return
}


}


}, 2)

tmp11976 := Call(__e, ns2_1set, symshen_4goto_1h, tmp11966)


_ = tmp11976

tmp11977 := MakeNative(func(__e *ControlFlow) {
V1738 := __e.Get(1)
_ = V1738
V1739 := __e.Get(2)
_ = V1739
tmp11980 := PrimEqual(Nil, V1739)

if True == tmp11980 {
tmp11978 := PrimCons(V1738, Nil)

__e.Return(PrimCons(symthaw, tmp11978))
return


} else {
__e.Return(PrimCons(V1738, V1739))
return
}


}, 2)

tmp11981 := Call(__e, ns2_1set, symshen_4invoke, tmp11977)


_ = tmp11981

tmp11982 := MakeNative(func(__e *ControlFlow) {
V1740 := __e.Get(1)
_ = V1740
__e.Return(PrimEqual(V1740, sym__))
return
}, 1)

tmp11983 := Call(__e, ns2_1set, symshen_4wildcard_2, tmp11982)


_ = tmp11983

tmp11984 := MakeNative(func(__e *ControlFlow) {
V1741 := __e.Get(1)
_ = V1741
tmp11985 := MakeNative(func(__e *ControlFlow) {
tmp11990 := PrimIsVector(V1741)

if True == tmp11990 {
tmp11987 := PrimVectorGet(V1741, MakeNumber(0))

tmp11988 := PrimEqual(tmp11987, symshen_4pvar)

if True == tmp11988 {
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

tmp11991 := MakeNative(func(__e *ControlFlow) {
Z1742 := __e.Get(1)
_ = Z1742
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp11985, tmp11991)
return


}, 1)

tmp11992 := Call(__e, ns2_1set, symshen_4pvar_2, tmp11984)


_ = tmp11992

tmp11993 := MakeNative(func(__e *ControlFlow) {
V1743 := __e.Get(1)
_ = V1743
V1744 := __e.Get(2)
_ = V1744
tmp12000 := Call(__e, PrimFunc(symshen_4pvar_2), V1743)


if True == tmp12000 {
tmp11994 := MakeNative(func(__e *ControlFlow) {
W1745 := __e.Get(1)
_ = W1745
tmp11996 := PrimEqual(W1745, symshen_4_1null_1)

if True == tmp11996 {
__e.Return(V1743)
return
} else {
__e.TailApply(PrimFunc(symshen_4lazyderef), W1745, V1744)
return
}


}, 1)

tmp11997 := PrimVectorGet(V1743, MakeNumber(1))

tmp11998 := PrimVectorGet(V1744, tmp11997)

__e.TailApply(tmp11994, tmp11998)
return


} else {
__e.Return(V1743)
return
}


}, 2)

tmp12001 := Call(__e, ns2_1set, symshen_4lazyderef, tmp11993)


_ = tmp12001

tmp12002 := MakeNative(func(__e *ControlFlow) {
V1746 := __e.Get(1)
_ = V1746
V1747 := __e.Get(2)
_ = V1747
tmp12015 := PrimIsPair(V1746)

if True == tmp12015 {
tmp12003 := PrimHead(V1746)

tmp12004 := Call(__e, PrimFunc(symshen_4deref), tmp12003, V1747)


tmp12005 := PrimTail(V1746)

tmp12006 := Call(__e, PrimFunc(symshen_4deref), tmp12005, V1747)


__e.Return(PrimCons(tmp12004, tmp12006))
return


} else {
tmp12013 := Call(__e, PrimFunc(symshen_4pvar_2), V1746)


if True == tmp12013 {
tmp12007 := MakeNative(func(__e *ControlFlow) {
W1748 := __e.Get(1)
_ = W1748
tmp12009 := PrimEqual(W1748, symshen_4_1null_1)

if True == tmp12009 {
__e.Return(V1746)
return
} else {
__e.TailApply(PrimFunc(symshen_4deref), W1748, V1747)
return
}


}, 1)

tmp12010 := PrimVectorGet(V1746, MakeNumber(1))

tmp12011 := PrimVectorGet(V1747, tmp12010)

__e.TailApply(tmp12007, tmp12011)
return


} else {
__e.Return(V1746)
return
}


}


}, 2)

tmp12016 := Call(__e, ns2_1set, symshen_4deref, tmp12002)


_ = tmp12016

tmp12017 := MakeNative(func(__e *ControlFlow) {
V1749 := __e.Get(1)
_ = V1749
V1750 := __e.Get(2)
_ = V1750
V1751 := __e.Get(3)
_ = V1751
V1752 := __e.Get(4)
_ = V1752
tmp12018 := MakeNative(func(__e *ControlFlow) {
W1753 := __e.Get(1)
_ = W1753
tmp12019 := MakeNative(func(__e *ControlFlow) {
W1754 := __e.Get(1)
_ = W1754
tmp12021 := PrimEqual(W1754, False)

if True == tmp12021 {
__e.TailApply(PrimFunc(symshen_4unwind), V1749, V1751, W1754)
return
} else {
__e.Return(W1754)
return
}


}, 1)

tmp12022 := Call(__e, PrimFunc(symthaw), V1752)


__e.TailApply(tmp12019, tmp12022)
return


}, 1)

tmp12023 := Call(__e, PrimFunc(symshen_4bindv), V1749, V1750, V1751)


__e.TailApply(tmp12018, tmp12023)
return


}, 4)

tmp12024 := Call(__e, ns2_1set, symshen_4bind_b, tmp12017)


_ = tmp12024

tmp12025 := MakeNative(func(__e *ControlFlow) {
V1755 := __e.Get(1)
_ = V1755
V1756 := __e.Get(2)
_ = V1756
V1757 := __e.Get(3)
_ = V1757
tmp12026 := PrimVectorGet(V1755, MakeNumber(1))

__e.Return(PrimVectorSet(V1757, tmp12026, V1756))
return


}, 3)

tmp12027 := Call(__e, ns2_1set, symshen_4bindv, tmp12025)


_ = tmp12027

tmp12028 := MakeNative(func(__e *ControlFlow) {
V1758 := __e.Get(1)
_ = V1758
V1759 := __e.Get(2)
_ = V1759
V1760 := __e.Get(3)
_ = V1760
tmp12029 := PrimVectorGet(V1758, MakeNumber(1))

tmp12030 := PrimVectorSet(V1759, tmp12029, symshen_4_1null_1)

_ = tmp12030

__e.Return(V1760)
return


}, 3)

tmp12031 := Call(__e, ns2_1set, symshen_4unwind, tmp12028)


_ = tmp12031

tmp12032 := MakeNative(func(__e *ControlFlow) {
V1769 := __e.Get(1)
_ = V1769
V1770 := __e.Get(2)
_ = V1770
V1771 := __e.Get(3)
_ = V1771
tmp12047 := PrimEqual(Nil, V1769)

if True == tmp12047 {
__e.Return(V1770)
return
} else {
tmp12045 := PrimIsPair(V1769)

if True == tmp12045 {
tmp12033 := PrimHead(V1769)

tmp12034 := PrimCons(V1771, Nil)

tmp12035 := PrimCons(symshen_4newpv, tmp12034)

tmp12036 := PrimTail(V1769)

tmp12037 := Call(__e, PrimFunc(symshen_4stpart), tmp12036, V1770, V1771)


tmp12038 := PrimCons(tmp12037, Nil)

tmp12039 := PrimCons(V1771, tmp12038)

tmp12040 := PrimCons(symshen_4gc, tmp12039)

tmp12041 := PrimCons(tmp12040, Nil)

tmp12042 := PrimCons(tmp12035, tmp12041)

tmp12043 := PrimCons(tmp12033, tmp12042)

__e.Return(PrimCons(symlet, tmp12043))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.stpart")))
return
}


}


}, 3)

tmp12048 := Call(__e, ns2_1set, symshen_4stpart, tmp12032)


_ = tmp12048

tmp12049 := MakeNative(func(__e *ControlFlow) {
V1772 := __e.Get(1)
_ = V1772
V1773 := __e.Get(2)
_ = V1773
tmp12054 := PrimEqual(V1773, False)

if True == tmp12054 {
tmp12050 := MakeNative(func(__e *ControlFlow) {
W1774 := __e.Get(1)
_ = W1774
tmp12051 := Call(__e, PrimFunc(symshen_4decrement_1ticket), W1774, V1772)


_ = tmp12051

__e.Return(V1773)
return


}, 1)

tmp12052 := Call(__e, PrimFunc(symshen_4ticket_1number), V1772)


__e.TailApply(tmp12050, tmp12052)
return


} else {
__e.Return(V1773)
return
}


}, 2)

tmp12055 := Call(__e, ns2_1set, symshen_4gc, tmp12049)


_ = tmp12055

tmp12056 := MakeNative(func(__e *ControlFlow) {
V1775 := __e.Get(1)
_ = V1775
V1776 := __e.Get(2)
_ = V1776
tmp12057 := PrimNumberSubtract(V1775, MakeNumber(1))

__e.Return(PrimVectorSet(V1776, MakeNumber(1), tmp12057))
return


}, 2)

tmp12058 := Call(__e, ns2_1set, symshen_4decrement_1ticket, tmp12056)


_ = tmp12058

tmp12059 := MakeNative(func(__e *ControlFlow) {
V1777 := __e.Get(1)
_ = V1777
tmp12060 := MakeNative(func(__e *ControlFlow) {
W1778 := __e.Get(1)
_ = W1778
tmp12061 := MakeNative(func(__e *ControlFlow) {
W1779 := __e.Get(1)
_ = W1779
tmp12062 := MakeNative(func(__e *ControlFlow) {
W1780 := __e.Get(1)
_ = W1780
__e.Return(W1779)
return
}, 1)

tmp12063 := Call(__e, PrimFunc(symshen_4nextticket), V1777, W1778)


__e.TailApply(tmp12062, tmp12063)
return


}, 1)

tmp12064 := Call(__e, PrimFunc(symshen_4make_1prolog_1variable), W1778)


__e.TailApply(tmp12061, tmp12064)
return


}, 1)

tmp12065 := Call(__e, PrimFunc(symshen_4ticket_1number), V1777)


__e.TailApply(tmp12060, tmp12065)
return


}, 1)

tmp12066 := Call(__e, ns2_1set, symshen_4newpv, tmp12059)


_ = tmp12066

tmp12067 := MakeNative(func(__e *ControlFlow) {
V1781 := __e.Get(1)
_ = V1781
__e.Return(PrimVectorGet(V1781, MakeNumber(1)))
return
}, 1)

tmp12068 := Call(__e, ns2_1set, symshen_4ticket_1number, tmp12067)


_ = tmp12068

tmp12069 := MakeNative(func(__e *ControlFlow) {
V1782 := __e.Get(1)
_ = V1782
V1783 := __e.Get(2)
_ = V1783
tmp12070 := MakeNative(func(__e *ControlFlow) {
W1784 := __e.Get(1)
_ = W1784
tmp12071 := PrimNumberAdd(V1783, MakeNumber(1))

__e.Return(PrimVectorSet(W1784, MakeNumber(1), tmp12071))
return


}, 1)

tmp12072 := PrimVectorSet(V1782, V1783, symshen_4_1null_1)

__e.TailApply(tmp12070, tmp12072)
return


}, 2)

tmp12073 := Call(__e, ns2_1set, symshen_4nextticket, tmp12069)


_ = tmp12073

tmp12074 := MakeNative(func(__e *ControlFlow) {
V1785 := __e.Get(1)
_ = V1785
tmp12075 := PrimAbsvector(MakeNumber(2))

tmp12076 := PrimVectorSet(tmp12075, MakeNumber(0), symshen_4pvar)

__e.Return(PrimVectorSet(tmp12076, MakeNumber(1), V1785))
return


}, 1)

tmp12077 := Call(__e, ns2_1set, symshen_4make_1prolog_1variable, tmp12074)


_ = tmp12077

tmp12078 := MakeNative(func(__e *ControlFlow) {
V1786 := __e.Get(1)
_ = V1786
tmp12079 := PrimVectorGet(V1786, MakeNumber(1))

tmp12080 := Call(__e, PrimFunc(symshen_4app), tmp12079, MakeString(""), symshen_4a)


__e.Return(PrimStringConcat(MakeString("Var"), tmp12080))
return


}, 1)

tmp12081 := Call(__e, ns2_1set, symshen_4pvar, tmp12078)


_ = tmp12081

tmp12082 := MakeNative(func(__e *ControlFlow) {
tmp12083 := PrimValue(symshen_4_dinfs_d)

tmp12084 := PrimNumberAdd(MakeNumber(1), tmp12083)

__e.Return(PrimSet(symshen_4_dinfs_d, tmp12084))
return


}, 0)

tmp12085 := Call(__e, ns2_1set, symshen_4incinfs, tmp12082)


_ = tmp12085

tmp12086 := MakeNative(func(__e *ControlFlow) {
V1787 := __e.Get(1)
_ = V1787
tmp12093 := PrimIsInteger(V1787)

var ifres12090 Obj

if True == tmp12093 {
tmp12092 := PrimGreatThan(V1787, MakeNumber(0))

var ifres12091 Obj

if True == tmp12092 {
ifres12091 = True


} else {
ifres12091 = False


}

ifres12090 = ifres12091


} else {
ifres12090 = False


}

if True == ifres12090 {
__e.Return(PrimSet(symshen_4_dsize_1prolog_1vector_d, V1787))
return
} else {
tmp12087 := Call(__e, PrimFunc(symshen_4app), V1787, MakeString(""), symshen_4a)


tmp12088 := PrimStringConcat(MakeString("prolog vector size: size should be a positive integer; not "), tmp12087)

__e.Return(PrimSimpleError(tmp12088))
return


}


}, 1)

tmp12094 := Call(__e, ns2_1set, symshen_4prolog_1vector_1size, tmp12086)


_ = tmp12094

tmp12095 := MakeNative(func(__e *ControlFlow) {
V1799 := __e.Get(1)
_ = V1799
V1800 := __e.Get(2)
_ = V1800
V1801 := __e.Get(3)
_ = V1801
V1802 := __e.Get(4)
_ = V1802
tmp12125 := PrimEqual(V1799, V1800)

if True == tmp12125 {
__e.TailApply(PrimFunc(symthaw), V1802)
return
} else {
tmp12123 := Call(__e, PrimFunc(symshen_4pvar_2), V1799)


var ifres12118 Obj

if True == tmp12123 {
tmp12120 := Call(__e, PrimFunc(symshen_4deref), V1800, V1801)


tmp12121 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V1799, tmp12120)


tmp12122 := PrimNot(tmp12121)

var ifres12119 Obj

if True == tmp12122 {
ifres12119 = True


} else {
ifres12119 = False


}

ifres12118 = ifres12119


} else {
ifres12118 = False


}

if True == ifres12118 {
__e.TailApply(PrimFunc(symshen_4bind_b), V1799, V1800, V1801, V1802)
return
} else {
tmp12116 := Call(__e, PrimFunc(symshen_4pvar_2), V1800)


var ifres12111 Obj

if True == tmp12116 {
tmp12113 := Call(__e, PrimFunc(symshen_4deref), V1799, V1801)


tmp12114 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V1800, tmp12113)


tmp12115 := PrimNot(tmp12114)

var ifres12112 Obj

if True == tmp12115 {
ifres12112 = True


} else {
ifres12112 = False


}

ifres12111 = ifres12112


} else {
ifres12111 = False


}

if True == ifres12111 {
__e.TailApply(PrimFunc(symshen_4bind_b), V1800, V1799, V1801, V1802)
return
} else {
tmp12109 := PrimIsPair(V1799)

var ifres12106 Obj

if True == tmp12109 {
tmp12108 := PrimIsPair(V1800)

var ifres12107 Obj

if True == tmp12108 {
ifres12107 = True


} else {
ifres12107 = False


}

ifres12106 = ifres12107


} else {
ifres12106 = False


}

if True == ifres12106 {
tmp12096 := PrimHead(V1799)

tmp12097 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12096, V1801)


tmp12098 := PrimHead(V1800)

tmp12099 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12098, V1801)


tmp12100 := MakeNative(func(__e *ControlFlow) {
tmp12101 := PrimTail(V1799)

tmp12102 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12101, V1801)


tmp12103 := PrimTail(V1800)

tmp12104 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12103, V1801)


__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp12102, tmp12104, V1801, V1802)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp12097, tmp12099, V1801, tmp12100)
return


} else {
__e.Return(False)
return
}


}


}


}


}, 4)

tmp12126 := Call(__e, ns2_1set, symshen_4lzy_a_b, tmp12095)


_ = tmp12126

tmp12127 := MakeNative(func(__e *ControlFlow) {
V1814 := __e.Get(1)
_ = V1814
V1815 := __e.Get(2)
_ = V1815
V1816 := __e.Get(3)
_ = V1816
V1817 := __e.Get(4)
_ = V1817
tmp12147 := PrimEqual(V1814, V1815)

if True == tmp12147 {
__e.TailApply(PrimFunc(symthaw), V1817)
return
} else {
tmp12145 := Call(__e, PrimFunc(symshen_4pvar_2), V1814)


if True == tmp12145 {
__e.TailApply(PrimFunc(symshen_4bind_b), V1814, V1815, V1816, V1817)
return
} else {
tmp12143 := Call(__e, PrimFunc(symshen_4pvar_2), V1815)


if True == tmp12143 {
__e.TailApply(PrimFunc(symshen_4bind_b), V1815, V1814, V1816, V1817)
return
} else {
tmp12141 := PrimIsPair(V1814)

var ifres12138 Obj

if True == tmp12141 {
tmp12140 := PrimIsPair(V1815)

var ifres12139 Obj

if True == tmp12140 {
ifres12139 = True


} else {
ifres12139 = False


}

ifres12138 = ifres12139


} else {
ifres12138 = False


}

if True == ifres12138 {
tmp12128 := PrimHead(V1814)

tmp12129 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12128, V1816)


tmp12130 := PrimHead(V1815)

tmp12131 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12130, V1816)


tmp12132 := MakeNative(func(__e *ControlFlow) {
tmp12133 := PrimTail(V1814)

tmp12134 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12133, V1816)


tmp12135 := PrimTail(V1815)

tmp12136 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12135, V1816)


__e.TailApply(PrimFunc(symshen_4lzy_a), tmp12134, tmp12136, V1816, V1817)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4lzy_a), tmp12129, tmp12131, V1816, tmp12132)
return


} else {
__e.Return(False)
return
}


}


}


}


}, 4)

tmp12148 := Call(__e, ns2_1set, symshen_4lzy_a, tmp12127)


_ = tmp12148

tmp12149 := MakeNative(func(__e *ControlFlow) {
V1823 := __e.Get(1)
_ = V1823
V1824 := __e.Get(2)
_ = V1824
tmp12159 := PrimEqual(V1823, V1824)

if True == tmp12159 {
__e.Return(True)
return
} else {
tmp12157 := PrimIsPair(V1824)

if True == tmp12157 {
tmp12154 := PrimHead(V1824)

tmp12155 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V1823, tmp12154)


if True == tmp12155 {
__e.Return(True)
return
} else {
tmp12151 := PrimTail(V1824)

tmp12152 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V1823, tmp12151)


if True == tmp12152 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


} else {
__e.Return(False)
return
}


}


}, 2)

tmp12160 := Call(__e, ns2_1set, symshen_4occurs_1check_2, tmp12149)


_ = tmp12160

tmp12161 := MakeNative(func(__e *ControlFlow) {
V1825 := __e.Get(1)
_ = V1825
V1826 := __e.Get(2)
_ = V1826
V1827 := __e.Get(3)
_ = V1827
V1828 := __e.Get(4)
_ = V1828
V1829 := __e.Get(5)
_ = V1829
tmp12162 := Call(__e, V1825, V1826)


tmp12163 := Call(__e, tmp12162, V1827)


tmp12164 := Call(__e, tmp12163, V1828)


__e.TailApply(tmp12164, V1829)
return


}, 5)

tmp12165 := Call(__e, ns2_1set, symcall, tmp12161)


_ = tmp12165

tmp12166 := MakeNative(func(__e *ControlFlow) {
V1836 := __e.Get(1)
_ = V1836
V1837 := __e.Get(2)
_ = V1837
V1838 := __e.Get(3)
_ = V1838
V1839 := __e.Get(4)
_ = V1839
V1840 := __e.Get(5)
_ = V1840
__e.TailApply(PrimFunc(symshen_4deref), V1836, V1837)
return
}, 5)

tmp12167 := Call(__e, ns2_1set, symreturn, tmp12166)


_ = tmp12167

tmp12168 := MakeNative(func(__e *ControlFlow) {
V1847 := __e.Get(1)
_ = V1847
V1848 := __e.Get(2)
_ = V1848
V1849 := __e.Get(3)
_ = V1849
V1850 := __e.Get(4)
_ = V1850
V1851 := __e.Get(5)
_ = V1851
if True == V1847 {
__e.TailApply(PrimFunc(symthaw), V1851)
return
} else {
__e.Return(False)
return
}
}, 5)

tmp12170 := Call(__e, ns2_1set, symwhen, tmp12168)


_ = tmp12170

tmp12171 := MakeNative(func(__e *ControlFlow) {
V1852 := __e.Get(1)
_ = V1852
V1853 := __e.Get(2)
_ = V1853
V1854 := __e.Get(3)
_ = V1854
V1855 := __e.Get(4)
_ = V1855
V1856 := __e.Get(5)
_ = V1856
V1857 := __e.Get(6)
_ = V1857
tmp12172 := Call(__e, PrimFunc(symshen_4lazyderef), V1852, V1854)


tmp12173 := Call(__e, PrimFunc(symshen_4lazyderef), V1853, V1854)


__e.TailApply(PrimFunc(symshen_4lzy_a), tmp12172, tmp12173, V1854, V1857)
return


}, 6)

tmp12174 := Call(__e, ns2_1set, symis, tmp12171)


_ = tmp12174

tmp12175 := MakeNative(func(__e *ControlFlow) {
V1858 := __e.Get(1)
_ = V1858
V1859 := __e.Get(2)
_ = V1859
V1860 := __e.Get(3)
_ = V1860
V1861 := __e.Get(4)
_ = V1861
V1862 := __e.Get(5)
_ = V1862
V1863 := __e.Get(6)
_ = V1863
tmp12176 := Call(__e, PrimFunc(symshen_4lazyderef), V1858, V1860)


tmp12177 := Call(__e, PrimFunc(symshen_4lazyderef), V1859, V1860)


__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp12176, tmp12177, V1860, V1863)
return


}, 6)

tmp12178 := Call(__e, ns2_1set, symis_b, tmp12175)


_ = tmp12178

tmp12179 := MakeNative(func(__e *ControlFlow) {
V1868 := __e.Get(1)
_ = V1868
V1869 := __e.Get(2)
_ = V1869
V1870 := __e.Get(3)
_ = V1870
V1871 := __e.Get(4)
_ = V1871
V1872 := __e.Get(5)
_ = V1872
V1873 := __e.Get(6)
_ = V1873
__e.TailApply(PrimFunc(symshen_4bind_b), V1868, V1869, V1870, V1873)
return
}, 6)

tmp12180 := Call(__e, ns2_1set, symbind, tmp12179)


_ = tmp12180

tmp12181 := MakeNative(func(__e *ControlFlow) {
V1874 := __e.Get(1)
_ = V1874
V1875 := __e.Get(2)
_ = V1875
V1876 := __e.Get(3)
_ = V1876
V1877 := __e.Get(4)
_ = V1877
V1878 := __e.Get(5)
_ = V1878
tmp12183 := Call(__e, PrimFunc(symshen_4lazyderef), V1874, V1875)


tmp12184 := Call(__e, PrimFunc(symshen_4pvar_2), tmp12183)


if True == tmp12184 {
__e.TailApply(PrimFunc(symthaw), V1878)
return
} else {
__e.Return(False)
return
}


}, 5)

tmp12185 := Call(__e, ns2_1set, symvar_2, tmp12181)


_ = tmp12185

tmp12186 := MakeNative(func(__e *ControlFlow) {
V1881 := __e.Get(1)
_ = V1881
__e.Return(MakeString("|prolog vector|"))
return
}, 1)

tmp12187 := Call(__e, ns2_1set, symshen_4print_1prolog_1vector, tmp12186)


_ = tmp12187

tmp12188 := MakeNative(func(__e *ControlFlow) {
V1900 := __e.Get(1)
_ = V1900
V1901 := __e.Get(2)
_ = V1901
V1902 := __e.Get(3)
_ = V1902
V1903 := __e.Get(4)
_ = V1903
V1904 := __e.Get(5)
_ = V1904
tmp12201 := PrimEqual(Nil, V1900)

if True == tmp12201 {
__e.Return(False)
return
} else {
tmp12199 := PrimIsPair(V1900)

if True == tmp12199 {
tmp12189 := MakeNative(func(__e *ControlFlow) {
W1905 := __e.Get(1)
_ = W1905
tmp12192 := PrimEqual(W1905, False)

if True == tmp12192 {
tmp12190 := PrimTail(V1900)

__e.TailApply(PrimFunc(symfork), tmp12190, V1901, V1902, V1903, V1904)
return


} else {
__e.Return(W1905)
return
}


}, 1)

tmp12193 := PrimHead(V1900)

tmp12194 := Call(__e, tmp12193, V1901)


tmp12195 := Call(__e, tmp12194, V1902)


tmp12196 := Call(__e, tmp12195, V1903)


tmp12197 := Call(__e, tmp12196, V1904)


__e.TailApply(tmp12189, tmp12197)
return


} else {
__e.Return(PrimSimpleError(MakeString("fork expects a list of literals\n")))
return
}


}


}, 5)

tmp12202 := Call(__e, ns2_1set, symfork, tmp12188)


_ = tmp12202

tmp12203 := MakeNative(func(__e *ControlFlow) {
V1906 := __e.Get(1)
_ = V1906
V1907 := __e.Get(2)
_ = V1907
V1908 := __e.Get(3)
_ = V1908
V1909 := __e.Get(4)
_ = V1909
V1910 := __e.Get(5)
_ = V1910
V1911 := __e.Get(6)
_ = V1911
V1912 := __e.Get(7)
_ = V1912
tmp12210 := Call(__e, PrimFunc(symshen_4unlocked_2), V1910)


if True == tmp12210 {
tmp12204 := MakeNative(func(__e *ControlFlow) {
W1913 := __e.Get(1)
_ = W1913
tmp12205 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12205

tmp12206 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4findall_1h), V1906, V1907, V1908, W1913, V1909, V1910, V1911, V1912)
return
}, 0)

tmp12207 := Call(__e, PrimFunc(symis), W1913, Nil, V1909, V1910, V1911, tmp12206)


__e.TailApply(PrimFunc(symshen_4gc), V1909, tmp12207)
return


}, 1)

tmp12208 := Call(__e, PrimFunc(symshen_4newpv), V1909)


__e.TailApply(tmp12204, tmp12208)
return


} else {
__e.Return(False)
return
}


}, 7)

tmp12211 := Call(__e, ns2_1set, symfindall, tmp12203)


_ = tmp12211

tmp12212 := MakeNative(func(__e *ControlFlow) {
V1914 := __e.Get(1)
_ = V1914
V1915 := __e.Get(2)
_ = V1915
V1916 := __e.Get(3)
_ = V1916
V1917 := __e.Get(4)
_ = V1917
V1918 := __e.Get(5)
_ = V1918
V1919 := __e.Get(6)
_ = V1919
V1920 := __e.Get(7)
_ = V1920
V1921 := __e.Get(8)
_ = V1921
tmp12213 := MakeNative(func(__e *ControlFlow) {
W1922 := __e.Get(1)
_ = W1922
tmp12218 := PrimEqual(W1922, False)

if True == tmp12218 {
tmp12216 := Call(__e, PrimFunc(symshen_4unlocked_2), V1919)


if True == tmp12216 {
tmp12214 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12214

__e.TailApply(PrimFunc(symis_b), V1916, V1917, V1918, V1919, V1920, V1921)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W1922)
return
}


}, 1)

tmp12223 := Call(__e, PrimFunc(symshen_4unlocked_2), V1919)


var ifres12219 Obj

if True == tmp12223 {
tmp12220 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12220

tmp12221 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4overbind), V1914, V1917, V1918, V1919, V1920, V1921)
return
}, 0)

tmp12222 := Call(__e, PrimFunc(symcall), V1915, V1918, V1919, V1920, tmp12221)


ifres12219 = tmp12222


} else {
ifres12219 = False


}

__e.TailApply(tmp12213, ifres12219)
return


}, 8)

tmp12224 := Call(__e, ns2_1set, symshen_4findall_1h, tmp12212)


_ = tmp12224

tmp12225 := MakeNative(func(__e *ControlFlow) {
V1929 := __e.Get(1)
_ = V1929
V1930 := __e.Get(2)
_ = V1930
V1931 := __e.Get(3)
_ = V1931
V1932 := __e.Get(4)
_ = V1932
V1933 := __e.Get(5)
_ = V1933
V1934 := __e.Get(6)
_ = V1934
tmp12226 := Call(__e, PrimFunc(symshen_4deref), V1929, V1931)


tmp12227 := Call(__e, PrimFunc(symshen_4lazyderef), V1930, V1931)


tmp12228 := PrimCons(tmp12226, tmp12227)

tmp12229 := Call(__e, PrimFunc(symshen_4bindv), V1930, tmp12228, V1931)


_ = tmp12229

__e.Return(False)
return


}, 6)

tmp12230 := Call(__e, ns2_1set, symshen_4overbind, tmp12225)


_ = tmp12230

tmp12231 := MakeNative(func(__e *ControlFlow) {
V1937 := __e.Get(1)
_ = V1937
tmp12235 := PrimEqual(sym_7, V1937)

if True == tmp12235 {
__e.Return(PrimSet(symshen_4_doccurs_d, True))
return
} else {
tmp12233 := PrimEqual(sym_1, V1937)

if True == tmp12233 {
__e.Return(PrimSet(symshen_4_doccurs_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("occurs-check expects a + or a -.\n")))
return
}


}


}, 1)

__e.TailApply(ns2_1set, symoccurs_1check, tmp12231)
return




}, 0)

