package main

import . "github.com/pyrex41/shen-go/kl"

var PrologMain = MakeNative(func(__e *ControlFlow) {
tmp10206 := MakeNative(func(__e *ControlFlow) {
V1245 := __e.Get(1)
_ = V1245
__e.TailApply(PrimFunc(symshen_4assert_d), V1245, symshen_4top)
return
}, 1)

tmp10207 := Call(__e, ns2_1set, symasserta, tmp10206)


_ = tmp10207

tmp10208 := MakeNative(func(__e *ControlFlow) {
V1246 := __e.Get(1)
_ = V1246
__e.TailApply(PrimFunc(symshen_4assert_d), V1246, symshen_4bottom)
return
}, 1)

tmp10209 := Call(__e, ns2_1set, symassertz, tmp10208)


_ = tmp10209

tmp10210 := MakeNative(func(__e *ControlFlow) {
V1247 := __e.Get(1)
_ = V1247
V1248 := __e.Get(2)
_ = V1248
tmp10244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1247)
}
__typedArg0 := V1247
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10235 Obj

if True == tmp10244 {
tmp10242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1247)
}
__typedArg0 := V1247
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10242)
}
__typedArg0 := tmp10242
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10237 Obj

if True == tmp10243 {
tmp10239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1247)
}
__typedArg0 := V1247
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10239)
}
__typedArg0 := tmp10239
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_5_1_1, tmp10240)
}
__typedArg0 := sym_5_1_1
__typedArg1 := tmp10240
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10238 Obj

if True == tmp10241 {
ifres10238 = True


} else {
ifres10238 = False


}

ifres10237 = ifres10238


} else {
ifres10237 = False


}

var ifres10236 Obj

if True == ifres10237 {
ifres10236 = True


} else {
ifres10236 = False


}

ifres10235 = ifres10236


} else {
ifres10235 = False


}

if True == ifres10235 {
tmp10211 := MakeNative(func(__e *ControlFlow) {
W1249 := __e.Get(1)
_ = W1249
tmp10212 := MakeNative(func(__e *ControlFlow) {
W1250 := __e.Get(1)
_ = W1250
tmp10213 := MakeNative(func(__e *ControlFlow) {
W1251 := __e.Get(1)
_ = W1251
tmp10214 := MakeNative(func(__e *ControlFlow) {
W1252 := __e.Get(1)
_ = W1252
tmp10215 := MakeNative(func(__e *ControlFlow) {
W1253 := __e.Get(1)
_ = W1253
tmp10216 := MakeNative(func(__e *ControlFlow) {
W1254 := __e.Get(1)
_ = W1254
tmp10217 := MakeNative(func(__e *ControlFlow) {
W1255 := __e.Get(1)
_ = W1255
__e.Return(W1249)
return
}, 1)

tmp10218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1247)
}
__typedArg0 := V1247
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10218)
}
__typedArg0 := tmp10218
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10220 := Call(__e, PrimFunc(symshen_4insert_1info), W1249, W1250, tmp10219, V1247, V1248)


__e.TailApply(tmp10217, tmp10220)
return


}, 1)

tmp10226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W1253, MakeNumber(-1))
}
__typedArg0 := W1253
__typedArg1 := MakeNumber(-1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10221 Obj

if True == tmp10226 {
tmp10222 := Call(__e, PrimFunc(symshen_4create_1skeleton), W1249, W1252)


tmp10223 := Call(__e, PrimFunc(symeval), tmp10222)


_ = tmp10223

tmp10224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp10225 := Call(__e, PrimFunc(symput), W1249, symshen_4dynamic, Nil, tmp10224)


ifres10221 = tmp10225


} else {
ifres10221 = symshen_4skip


}

__e.TailApply(tmp10216, ifres10221)
return


}, 1)

tmp10227 := Call(__e, PrimFunc(symarity), W1249)


__e.TailApply(tmp10215, tmp10227)
return


}, 1)

tmp10228 := Call(__e, PrimFunc(symshen_4parameters), W1251)


__e.TailApply(tmp10214, tmp10228)
return


}, 1)

tmp10229 := Call(__e, PrimFunc(symlength), W1250)


__e.TailApply(tmp10213, tmp10229)
return


}, 1)

tmp10230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1247)
}
__typedArg0 := V1247
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10231 := Call(__e, PrimFunc(symshen_4terms), tmp10230)


__e.TailApply(tmp10212, tmp10231)
return


}, 1)

tmp10232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1247)
}
__typedArg0 := V1247
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10233 := Call(__e, PrimFunc(symshen_4predicate), tmp10232)


__e.TailApply(tmp10211, tmp10233)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.assert*"))
}
__typedArg0 := MakeString("partial function shen.assert*")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 2)

tmp10245 := Call(__e, ns2_1set, symshen_4assert_d, tmp10210)


_ = tmp10245

tmp10246 := MakeNative(func(__e *ControlFlow) {
V1258 := __e.Get(1)
_ = V1258
tmp10248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1258)
}
__typedArg0 := V1258
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10248 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1258)
}
__typedArg0 := V1258
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
__e.Return(V1258)
return
}


}, 1)

tmp10249 := Call(__e, ns2_1set, symshen_4predicate, tmp10246)


_ = tmp10249

tmp10250 := MakeNative(func(__e *ControlFlow) {
V1263 := __e.Get(1)
_ = V1263
tmp10252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1263)
}
__typedArg0 := V1263
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10252 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1263)
}
__typedArg0 := V1263
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
} else {
__e.Return(Nil)
return
}


}, 1)

tmp10253 := Call(__e, ns2_1set, symshen_4terms, tmp10250)


_ = tmp10253

tmp10254 := MakeNative(func(__e *ControlFlow) {
V1264 := __e.Get(1)
_ = V1264
V1265 := __e.Get(2)
_ = V1265
tmp10255 := Call(__e, PrimFunc(symshen_4dynamic_1default), V1264, V1265)


tmp10256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1264, tmp10255)
}
__typedArg0 := V1264
__typedArg1 := tmp10255
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp10256)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp10256
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp10257 := Call(__e, ns2_1set, symshen_4create_1skeleton, tmp10254)


_ = tmp10257

tmp10258 := MakeNative(func(__e *ControlFlow) {
V1266 := __e.Get(1)
_ = V1266
V1267 := __e.Get(2)
_ = V1267
tmp10259 := Call(__e, PrimFunc(symshen_4cons_1form), V1267)


tmp10260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4dynamic, Nil)
}
__typedArg0 := symshen_4dynamic
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1266, tmp10260)
}
__typedArg0 := V1266
__typedArg1 := tmp10260
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget, tmp10261)
}
__typedArg0 := symget
__typedArg1 := tmp10261
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10262, Nil)
}
__typedArg0 := tmp10262
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10259, tmp10263)
}
__typedArg0 := tmp10259
__typedArg1 := tmp10263
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4call_1dynamic, tmp10264)
}
__typedArg0 := symshen_4call_1dynamic
__typedArg1 := tmp10264
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp10267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10266, Nil)
}
__typedArg0 := tmp10266
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10265, tmp10267)
}
__typedArg0 := tmp10265
__typedArg1 := tmp10267
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, tmp10268)
}
__typedArg0 := sym_5_1_1
__typedArg1 := tmp10268
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symappend), V1267, tmp10269)
return


}, 2)

tmp10270 := Call(__e, ns2_1set, symshen_4dynamic_1default, tmp10258)


_ = tmp10270

tmp10271 := MakeNative(func(__e *ControlFlow) {
V1268 := __e.Get(1)
_ = V1268
V1269 := __e.Get(2)
_ = V1269
V1270 := __e.Get(3)
_ = V1270
V1271 := __e.Get(4)
_ = V1271
V1272 := __e.Get(5)
_ = V1272
tmp10272 := MakeNative(func(__e *ControlFlow) {
W1273 := __e.Get(1)
_ = W1273
tmp10273 := MakeNative(func(__e *ControlFlow) {
W1274 := __e.Get(1)
_ = W1274
tmp10274 := MakeNative(func(__e *ControlFlow) {
W1275 := __e.Get(1)
_ = W1275
tmp10275 := MakeNative(func(__e *ControlFlow) {
W1276 := __e.Get(1)
_ = W1276
tmp10276 := MakeNative(func(__e *ControlFlow) {
W1277 := __e.Get(1)
_ = W1277
tmp10277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V1268, symshen_4dynamic, W1277, tmp10277)
return


}, 1)

tmp10282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1272, symshen_4top)
}
__typedArg0 := V1272
__typedArg1 := symshen_4top
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10278 Obj

if True == tmp10282 {
tmp10279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1275, W1276)
}
__typedArg0 := W1275
__typedArg1 := W1276
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

ifres10278 = tmp10279


} else {
tmp10280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1275, Nil)
}
__typedArg0 := W1275
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10281 := Call(__e, PrimFunc(symappend), W1276, tmp10280)


ifres10278 = tmp10281


}

__e.TailApply(tmp10276, ifres10278)
return


}, 1)

tmp10283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp10284 := Call(__e, PrimFunc(symget), V1268, symshen_4dynamic, tmp10283)


__e.TailApply(tmp10275, tmp10284)
return


}, 1)

tmp10285 := Call(__e, PrimFunc(symfn), W1273)


tmp10286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1273, V1271)
}
__typedArg0 := W1273
__typedArg1 := V1271
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10285, tmp10286)
}
__typedArg0 := tmp10285
__typedArg1 := tmp10286
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp10274, tmp10287)
return


}, 1)

tmp10288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1273, Nil)
}
__typedArg0 := W1273
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp10288)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp10288
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, V1270)
}
__typedArg0 := sym_5_1_1
__typedArg1 := V1270
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10291 := Call(__e, PrimFunc(symappend), V1269, tmp10290)


tmp10292 := Call(__e, PrimFunc(symappend), tmp10289, tmp10291)


tmp10293 := Call(__e, PrimFunc(symeval), tmp10292)


__e.TailApply(tmp10273, tmp10293)
return


}, 1)

tmp10294 := Call(__e, PrimFunc(symgensym), symshen_4g)


__e.TailApply(tmp10272, tmp10294)
return


}, 5)

tmp10295 := Call(__e, ns2_1set, symshen_4insert_1info, tmp10271)


_ = tmp10295

tmp10296 := MakeNative(func(__e *ControlFlow) {
tmp10297 := MakeNative(func(__e *ControlFlow) {
W1278 := __e.Get(1)
_ = W1278
tmp10298 := MakeNative(func(__e *ControlFlow) {
W1279 := __e.Get(1)
_ = W1279
__e.Return(W1279)
return
}, 1)

tmp10304 := Call(__e, PrimFunc(symempty_2), W1278)


var ifres10299 Obj

if True == tmp10304 {
tmp10300 := Call(__e, PrimFunc(symgensym), symshen_4g)


ifres10299 = tmp10300


} else {
tmp10301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W1278)
}
__typedArg0 := W1278
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dnames_d, tmp10301)
}
__typedArg0 := symshen_4_dnames_d
__typedArg1 := tmp10301
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp10302

tmp10303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W1278)
}
__typedArg0 := W1278
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

ifres10299 = tmp10303


}

__e.TailApply(tmp10298, ifres10299)
return


}, 1)

tmp10305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dnames_d)
}
__typedArg0 := symshen_4_dnames_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(tmp10297, tmp10305)
return


}, 0)

tmp10306 := Call(__e, ns2_1set, symshen_4newname, tmp10296)


_ = tmp10306

tmp10307 := MakeNative(func(__e *ControlFlow) {
V1280 := __e.Get(1)
_ = V1280
V1281 := __e.Get(2)
_ = V1281
V1282 := __e.Get(3)
_ = V1282
V1283 := __e.Get(4)
_ = V1283
V1284 := __e.Get(5)
_ = V1284
V1285 := __e.Get(6)
_ = V1285
tmp10308 := MakeNative(func(__e *ControlFlow) {
W1286 := __e.Get(1)
_ = W1286
tmp10319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W1286, False)
}
__typedArg0 := W1286
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10319 {
tmp10317 := Call(__e, PrimFunc(symshen_4unlocked_2), V1283)


if True == tmp10317 {
tmp10309 := MakeNative(func(__e *ControlFlow) {
W1290 := __e.Get(1)
_ = W1290
tmp10314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W1290)
}
__typedArg0 := W1290
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10314 {
tmp10310 := MakeNative(func(__e *ControlFlow) {
W1291 := __e.Get(1)
_ = W1291
tmp10311 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp10311

__e.TailApply(PrimFunc(symshen_4call_1dynamic), V1280, W1291, V1282, V1283, V1284, V1285)
return


}, 1)

tmp10312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(W1290)
}
__typedArg0 := W1290
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(tmp10310, tmp10312)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp10315 := Call(__e, PrimFunc(symshen_4lazyderef), V1281, V1282)


__e.TailApply(tmp10309, tmp10315)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W1286)
return
}


}, 1)

tmp10334 := Call(__e, PrimFunc(symshen_4unlocked_2), V1283)


var ifres10320 Obj

if True == tmp10334 {
tmp10321 := MakeNative(func(__e *ControlFlow) {
W1287 := __e.Get(1)
_ = W1287
tmp10331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W1287)
}
__typedArg0 := W1287
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10331 {
tmp10322 := MakeNative(func(__e *ControlFlow) {
W1288 := __e.Get(1)
_ = W1288
tmp10327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W1288)
}
__typedArg0 := W1288
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10327 {
tmp10323 := MakeNative(func(__e *ControlFlow) {
W1289 := __e.Get(1)
_ = W1289
tmp10324 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp10324

__e.TailApply(PrimFunc(symshen_4callrec), W1289, V1280, V1282, V1283, V1284, V1285)
return


}, 1)

tmp10325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W1288)
}
__typedArg0 := W1288
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(tmp10323, tmp10325)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp10328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(W1287)
}
__typedArg0 := W1287
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10329 := Call(__e, PrimFunc(symshen_4lazyderef), tmp10328, V1282)


__e.TailApply(tmp10322, tmp10329)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp10332 := Call(__e, PrimFunc(symshen_4lazyderef), V1281, V1282)


tmp10333 := Call(__e, tmp10321, tmp10332)


ifres10320 = tmp10333


} else {
ifres10320 = False


}

__e.TailApply(tmp10308, ifres10320)
return


}, 6)

tmp10335 := Call(__e, ns2_1set, symshen_4call_1dynamic, tmp10307)


_ = tmp10335

tmp10336 := MakeNative(func(__e *ControlFlow) {
V1292 := __e.Get(1)
_ = V1292
V1293 := __e.Get(2)
_ = V1293
V1294 := __e.Get(3)
_ = V1294
V1295 := __e.Get(4)
_ = V1295
V1296 := __e.Get(5)
_ = V1296
V1297 := __e.Get(6)
_ = V1297
tmp10346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1293)
}
__typedArg0 := Nil
__typedArg1 := V1293
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10346 {
tmp10337 := Call(__e, V1292, V1294)


tmp10338 := Call(__e, tmp10337, V1295)


tmp10339 := Call(__e, tmp10338, V1296)


__e.TailApply(tmp10339, V1297)
return


} else {
tmp10344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1293)
}
__typedArg0 := V1293
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10344 {
tmp10340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1293)
}
__typedArg0 := V1293
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10341 := Call(__e, V1292, tmp10340)


tmp10342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1293)
}
__typedArg0 := V1293
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4callrec), tmp10341, tmp10342, V1294, V1295, V1296, V1297)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.callrec"))
}
__typedArg0 := MakeString("partial function shen.callrec")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 6)

tmp10347 := Call(__e, ns2_1set, symshen_4callrec, tmp10336)


_ = tmp10347

tmp10348 := MakeNative(func(__e *ControlFlow) {
V1298 := __e.Get(1)
_ = V1298
tmp10367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1298)
}
__typedArg0 := V1298
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10358 Obj

if True == tmp10367 {
tmp10365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1298)
}
__typedArg0 := V1298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10365)
}
__typedArg0 := tmp10365
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10360 Obj

if True == tmp10366 {
tmp10362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1298)
}
__typedArg0 := V1298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10362)
}
__typedArg0 := tmp10362
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_5_1_1, tmp10363)
}
__typedArg0 := sym_5_1_1
__typedArg1 := tmp10363
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10361 Obj

if True == tmp10364 {
ifres10361 = True


} else {
ifres10361 = False


}

ifres10360 = ifres10361


} else {
ifres10360 = False


}

var ifres10359 Obj

if True == ifres10360 {
ifres10359 = True


} else {
ifres10359 = False


}

ifres10358 = ifres10359


} else {
ifres10358 = False


}

if True == ifres10358 {
tmp10349 := MakeNative(func(__e *ControlFlow) {
W1299 := __e.Get(1)
_ = W1299
tmp10350 := MakeNative(func(__e *ControlFlow) {
W1300 := __e.Get(1)
_ = W1300
tmp10351 := Call(__e, PrimFunc(symshen_4retract_1clause), V1298, W1300)


tmp10352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), W1299, symshen_4dynamic, tmp10351, tmp10352)
return


}, 1)

tmp10353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp10354 := Call(__e, PrimFunc(symget), W1299, symshen_4dynamic, tmp10353)


__e.TailApply(tmp10350, tmp10354)
return


}, 1)

tmp10355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1298)
}
__typedArg0 := V1298
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10356 := Call(__e, PrimFunc(symshen_4predicate), tmp10355)


__e.TailApply(tmp10349, tmp10356)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function retract"))
}
__typedArg0 := MakeString("partial function retract")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp10368 := Call(__e, ns2_1set, symretract, tmp10348)


_ = tmp10368

tmp10369 := MakeNative(func(__e *ControlFlow) {
V1306 := __e.Get(1)
_ = V1306
V1307 := __e.Get(2)
_ = V1307
tmp10399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1307)
}
__typedArg0 := Nil
__typedArg1 := V1307
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10399 {
__e.Return(Nil)
return
} else {
tmp10397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10382 Obj

if True == tmp10397 {
tmp10395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10395)
}
__typedArg0 := tmp10395
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10384 Obj

if True == tmp10396 {
tmp10392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10392)
}
__typedArg0 := tmp10392
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10393)
}
__typedArg0 := tmp10393
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10386 Obj

if True == tmp10394 {
tmp10388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10388)
}
__typedArg0 := tmp10388
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10389)
}
__typedArg0 := tmp10389
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1306, tmp10390)
}
__typedArg0 := V1306
__typedArg1 := tmp10390
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10387 Obj

if True == tmp10391 {
ifres10387 = True


} else {
ifres10387 = False


}

ifres10386 = ifres10387


} else {
ifres10386 = False


}

var ifres10385 Obj

if True == ifres10386 {
ifres10385 = True


} else {
ifres10385 = False


}

ifres10384 = ifres10385


} else {
ifres10384 = False


}

var ifres10383 Obj

if True == ifres10384 {
ifres10383 = True


} else {
ifres10383 = False


}

ifres10382 = ifres10383


} else {
ifres10382 = False


}

if True == ifres10382 {
tmp10370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10370)
}
__typedArg0 := tmp10370
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10371)
}
__typedArg0 := tmp10371
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dnames_d)
}
__typedArg0 := symshen_4_dnames_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp10374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10372, tmp10373)
}
__typedArg0 := tmp10372
__typedArg1 := tmp10373
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dnames_d, tmp10374)
}
__typedArg0 := symshen_4_dnames_d
__typedArg1 := tmp10374
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp10375

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return


} else {
tmp10380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10380 {
tmp10376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10378 := Call(__e, PrimFunc(symshen_4retract_1clause), V1306, tmp10377)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10376, tmp10378)
}
__typedArg0 := tmp10376
__typedArg1 := tmp10378
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.retract-clause"))
}
__typedArg0 := MakeString("partial function shen.retract-clause")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 2)

tmp10400 := Call(__e, ns2_1set, symshen_4retract_1clause, tmp10369)


_ = tmp10400

tmp10401 := MakeNative(func(__e *ControlFlow) {
V1308 := __e.Get(1)
_ = V1308
V1309 := __e.Get(2)
_ = V1309
tmp10402 := MakeNative(func(__e *ControlFlow) {
Z1310 := __e.Get(1)
_ = Z1310
__e.TailApply(PrimFunc(symshen_4_5defprolog_6), Z1310)
return
}, 1)

tmp10403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1308, V1309)
}
__typedArg0 := V1308
__typedArg1 := V1309
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symcompile), tmp10402, tmp10403)
return


}, 2)

tmp10404 := Call(__e, ns2_1set, symshen_4compile_1prolog, tmp10401)


_ = tmp10404

tmp10405 := MakeNative(func(__e *ControlFlow) {
V1311 := __e.Get(1)
_ = V1311
tmp10406 := MakeNative(func(__e *ControlFlow) {
W1312 := __e.Get(1)
_ = W1312
tmp10408 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1312)


if True == tmp10408 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1312)
return
}


}, 1)

tmp10430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1311)
}
__typedArg0 := V1311
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10409 Obj

if True == tmp10430 {
tmp10410 := MakeNative(func(__e *ControlFlow) {
W1313 := __e.Get(1)
_ = W1313
tmp10411 := MakeNative(func(__e *ControlFlow) {
W1314 := __e.Get(1)
_ = W1314
tmp10412 := MakeNative(func(__e *ControlFlow) {
W1315 := __e.Get(1)
_ = W1315
tmp10424 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1315)


if True == tmp10424 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10413 := MakeNative(func(__e *ControlFlow) {
W1316 := __e.Get(1)
_ = W1316
tmp10414 := MakeNative(func(__e *ControlFlow) {
W1317 := __e.Get(1)
_ = W1317
tmp10415 := MakeNative(func(__e *ControlFlow) {
W1318 := __e.Get(1)
_ = W1318
tmp10416 := MakeNative(func(__e *ControlFlow) {
W1319 := __e.Get(1)
_ = W1319
__e.TailApply(PrimFunc(symshen_4horn_1clause_1procedure), W1313, W1319)
return
}, 1)

tmp10417 := MakeNative(func(__e *ControlFlow) {
Z1320 := __e.Get(1)
_ = Z1320
__e.TailApply(PrimFunc(symshen_4linearise_1clause), Z1320)
return
}, 1)

tmp10418 := Call(__e, PrimFunc(symmap), tmp10417, W1316)


__e.TailApply(tmp10416, tmp10418)
return


}, 1)

tmp10419 := Call(__e, PrimFunc(symshen_4prolog_1arity_1check), W1313, W1316)


tmp10420 := Call(__e, tmp10415, tmp10419)


__e.TailApply(PrimFunc(symshen_4comb), W1317, tmp10420)
return


}, 1)

tmp10421 := Call(__e, PrimFunc(symshen_4in_1_6), W1315)


__e.TailApply(tmp10414, tmp10421)
return


}, 1)

tmp10422 := Call(__e, PrimFunc(symshen_4_5_1out), W1315)


__e.TailApply(tmp10413, tmp10422)
return


}


}, 1)

tmp10425 := Call(__e, PrimFunc(symshen_4_5clauses_6), W1314)


__e.TailApply(tmp10412, tmp10425)
return


}, 1)

tmp10426 := Call(__e, PrimFunc(symtail), V1311)


__e.TailApply(tmp10411, tmp10426)
return


}, 1)

tmp10427 := Call(__e, PrimFunc(symhead), V1311)


tmp10428 := Call(__e, tmp10410, tmp10427)


ifres10409 = tmp10428


} else {
tmp10429 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10409 = tmp10429


}

__e.TailApply(tmp10406, ifres10409)
return


}, 1)

tmp10431 := Call(__e, ns2_1set, symshen_4_5defprolog_6, tmp10405)


_ = tmp10431

tmp10432 := MakeNative(func(__e *ControlFlow) {
V1323 := __e.Get(1)
_ = V1323
V1324 := __e.Get(2)
_ = V1324
tmp10476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10457 Obj

if True == tmp10476 {
tmp10474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10474)
}
__typedArg0 := tmp10474
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10459 Obj

if True == tmp10475 {
tmp10471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10471)
}
__typedArg0 := tmp10471
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10472)
}
__typedArg0 := tmp10472
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10461 Obj

if True == tmp10473 {
tmp10467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10467)
}
__typedArg0 := tmp10467
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10468)
}
__typedArg0 := tmp10468
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10469)
}
__typedArg0 := Nil
__typedArg1 := tmp10469
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10463 Obj

if True == tmp10470 {
tmp10465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10465)
}
__typedArg0 := Nil
__typedArg1 := tmp10465
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10464 Obj

if True == tmp10466 {
ifres10464 = True


} else {
ifres10464 = False


}

ifres10463 = ifres10464


} else {
ifres10463 = False


}

var ifres10462 Obj

if True == ifres10463 {
ifres10462 = True


} else {
ifres10462 = False


}

ifres10461 = ifres10462


} else {
ifres10461 = False


}

var ifres10460 Obj

if True == ifres10461 {
ifres10460 = True


} else {
ifres10460 = False


}

ifres10459 = ifres10460


} else {
ifres10459 = False


}

var ifres10458 Obj

if True == ifres10459 {
ifres10458 = True


} else {
ifres10458 = False


}

ifres10457 = ifres10458


} else {
ifres10457 = False


}

if True == ifres10457 {
tmp10433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10433)
}
__typedArg0 := tmp10433
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symlength), tmp10434)
return


} else {
tmp10455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10440 Obj

if True == tmp10455 {
tmp10453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10453)
}
__typedArg0 := tmp10453
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10442 Obj

if True == tmp10454 {
tmp10450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10450)
}
__typedArg0 := tmp10450
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10451)
}
__typedArg0 := tmp10451
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10444 Obj

if True == tmp10452 {
tmp10446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10446)
}
__typedArg0 := tmp10446
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10447)
}
__typedArg0 := tmp10447
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10448)
}
__typedArg0 := Nil
__typedArg1 := tmp10448
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10445 Obj

if True == tmp10449 {
ifres10445 = True


} else {
ifres10445 = False


}

ifres10444 = ifres10445


} else {
ifres10444 = False


}

var ifres10443 Obj

if True == ifres10444 {
ifres10443 = True


} else {
ifres10443 = False


}

ifres10442 = ifres10443


} else {
ifres10442 = False


}

var ifres10441 Obj

if True == ifres10442 {
ifres10441 = True


} else {
ifres10441 = False


}

ifres10440 = ifres10441


} else {
ifres10440 = False


}

if True == ifres10440 {
tmp10435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10435)
}
__typedArg0 := tmp10435
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10437 := Call(__e, PrimFunc(symlength), tmp10436)


tmp10438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4pac_1h), V1323, tmp10437, tmp10438)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.prolog-arity-check"))
}
__typedArg0 := MakeString("partial function shen.prolog-arity-check")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp10477 := Call(__e, ns2_1set, symshen_4prolog_1arity_1check, tmp10432)


_ = tmp10477

tmp10478 := MakeNative(func(__e *ControlFlow) {
V1329 := __e.Get(1)
_ = V1329
V1330 := __e.Get(2)
_ = V1330
V1331 := __e.Get(3)
_ = V1331
tmp10494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1331)
}
__typedArg0 := Nil
__typedArg1 := V1331
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10494 {
__e.Return(V1330)
return
} else {
tmp10492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1331)
}
__typedArg0 := V1331
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10488 Obj

if True == tmp10492 {
tmp10490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1331)
}
__typedArg0 := V1331
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10490)
}
__typedArg0 := tmp10490
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10489 Obj

if True == tmp10491 {
ifres10489 = True


} else {
ifres10489 = False


}

ifres10488 = ifres10489


} else {
ifres10488 = False


}

if True == ifres10488 {
tmp10483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1331)
}
__typedArg0 := V1331
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10483)
}
__typedArg0 := tmp10483
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10485 := Call(__e, PrimFunc(symlength), tmp10484)


tmp10486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1330, tmp10485)
}
__typedArg0 := V1330
__typedArg1 := tmp10485
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10486 {
tmp10479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1331)
}
__typedArg0 := V1331
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4pac_1h), V1329, V1330, tmp10479)
return


} else {
tmp10480 := Call(__e, PrimFunc(symshen_4app), V1329, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("arity error in prolog procedure "))
__typedS1, __typedOK1 := TypedString(tmp10480)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("arity error in prolog procedure ")
__typedArg1 := tmp10480
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("arity error in prolog procedure "))
__typedS1, __typedOK1 := TypedString(tmp10480)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("arity error in prolog procedure ")
__typedArg1 := tmp10480
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.pac-h"))
}
__typedArg0 := MakeString("partial function shen.pac-h")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 3)

tmp10495 := Call(__e, ns2_1set, symshen_4pac_1h, tmp10478)


_ = tmp10495

tmp10496 := MakeNative(func(__e *ControlFlow) {
V1332 := __e.Get(1)
_ = V1332
tmp10497 := MakeNative(func(__e *ControlFlow) {
W1333 := __e.Get(1)
_ = W1333
tmp10516 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1333)


if True == tmp10516 {
tmp10498 := MakeNative(func(__e *ControlFlow) {
W1340 := __e.Get(1)
_ = W1340
tmp10500 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1340)


if True == tmp10500 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1340)
return
}


}, 1)

tmp10501 := MakeNative(func(__e *ControlFlow) {
W1341 := __e.Get(1)
_ = W1341
tmp10512 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1341)


if True == tmp10512 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10502 := MakeNative(func(__e *ControlFlow) {
W1342 := __e.Get(1)
_ = W1342
tmp10503 := MakeNative(func(__e *ControlFlow) {
W1343 := __e.Get(1)
_ = W1343
tmp10508 := Call(__e, PrimFunc(symempty_2), W1342)


var ifres10504 Obj

if True == tmp10508 {
ifres10504 = Nil


} else {
tmp10505 := Call(__e, PrimFunc(symshen_4app), W1342, MakeString("\n ..."), symshen_4r)


tmp10507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Prolog syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp10505)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Prolog syntax error here:\n ")
__typedArg1 := tmp10505
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Prolog syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp10505)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Prolog syntax error here:\n ")
__typedArg1 := tmp10505
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres10504 = tmp10507


}

__e.TailApply(PrimFunc(symshen_4comb), W1343, ifres10504)
return


}, 1)

tmp10509 := Call(__e, PrimFunc(symshen_4in_1_6), W1341)


__e.TailApply(tmp10503, tmp10509)
return


}, 1)

tmp10510 := Call(__e, PrimFunc(symshen_4_5_1out), W1341)


__e.TailApply(tmp10502, tmp10510)
return


}


}, 1)

tmp10513 := Call(__e, PrimFunc(sym_5_b_6), V1332)


tmp10514 := Call(__e, tmp10501, tmp10513)


__e.TailApply(tmp10498, tmp10514)
return


} else {
__e.Return(W1333)
return
}


}, 1)

tmp10517 := MakeNative(func(__e *ControlFlow) {
W1334 := __e.Get(1)
_ = W1334
tmp10532 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1334)


if True == tmp10532 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10518 := MakeNative(func(__e *ControlFlow) {
W1335 := __e.Get(1)
_ = W1335
tmp10519 := MakeNative(func(__e *ControlFlow) {
W1336 := __e.Get(1)
_ = W1336
tmp10520 := MakeNative(func(__e *ControlFlow) {
W1337 := __e.Get(1)
_ = W1337
tmp10527 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1337)


if True == tmp10527 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10521 := MakeNative(func(__e *ControlFlow) {
W1338 := __e.Get(1)
_ = W1338
tmp10522 := MakeNative(func(__e *ControlFlow) {
W1339 := __e.Get(1)
_ = W1339
tmp10523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1335, W1338)
}
__typedArg0 := W1335
__typedArg1 := W1338
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W1339, tmp10523)
return


}, 1)

tmp10524 := Call(__e, PrimFunc(symshen_4in_1_6), W1337)


__e.TailApply(tmp10522, tmp10524)
return


}, 1)

tmp10525 := Call(__e, PrimFunc(symshen_4_5_1out), W1337)


__e.TailApply(tmp10521, tmp10525)
return


}


}, 1)

tmp10528 := Call(__e, PrimFunc(symshen_4_5clauses_6), W1336)


__e.TailApply(tmp10520, tmp10528)
return


}, 1)

tmp10529 := Call(__e, PrimFunc(symshen_4in_1_6), W1334)


__e.TailApply(tmp10519, tmp10529)
return


}, 1)

tmp10530 := Call(__e, PrimFunc(symshen_4_5_1out), W1334)


__e.TailApply(tmp10518, tmp10530)
return


}


}, 1)

tmp10533 := Call(__e, PrimFunc(symshen_4_5clause_6), V1332)


tmp10534 := Call(__e, tmp10517, tmp10533)


__e.TailApply(tmp10497, tmp10534)
return


}, 1)

tmp10535 := Call(__e, ns2_1set, symshen_4_5clauses_6, tmp10496)


_ = tmp10535

tmp10536 := MakeNative(func(__e *ControlFlow) {
V1344 := __e.Get(1)
_ = V1344
tmp10552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1344)
}
__typedArg0 := V1344
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10543 Obj

if True == tmp10552 {
tmp10550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1344)
}
__typedArg0 := V1344
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10550)
}
__typedArg0 := tmp10550
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10545 Obj

if True == tmp10551 {
tmp10547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1344)
}
__typedArg0 := V1344
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10547)
}
__typedArg0 := tmp10547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10548)
}
__typedArg0 := Nil
__typedArg1 := tmp10548
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10546 Obj

if True == tmp10549 {
ifres10546 = True


} else {
ifres10546 = False


}

ifres10545 = ifres10546


} else {
ifres10545 = False


}

var ifres10544 Obj

if True == ifres10545 {
ifres10544 = True


} else {
ifres10544 = False


}

ifres10543 = ifres10544


} else {
ifres10543 = False


}

if True == ifres10543 {
tmp10537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1344)
}
__typedArg0 := V1344
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1344)
}
__typedArg0 := V1344
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10538)
}
__typedArg0 := tmp10538
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10540 := Call(__e, PrimFunc(sym_8p), tmp10537, tmp10539)


tmp10541 := Call(__e, PrimFunc(symshen_4linearise), tmp10540)


__e.TailApply(PrimFunc(symshen_4lch), tmp10541)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.linearise-clause"))
}
__typedArg0 := MakeString("partial function shen.linearise-clause")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp10553 := Call(__e, ns2_1set, symshen_4linearise_1clause, tmp10536)


_ = tmp10553

tmp10554 := MakeNative(func(__e *ControlFlow) {
V1345 := __e.Get(1)
_ = V1345
tmp10560 := Call(__e, PrimFunc(symtuple_2), V1345)


if True == tmp10560 {
tmp10555 := Call(__e, PrimFunc(symfst), V1345)


tmp10556 := Call(__e, PrimFunc(symsnd), V1345)


tmp10557 := Call(__e, PrimFunc(symshen_4lchh), tmp10556)


tmp10558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10557, Nil)
}
__typedArg0 := tmp10557
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10555, tmp10558)
}
__typedArg0 := tmp10555
__typedArg1 := tmp10558
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.lch"))
}
__typedArg0 := MakeString("partial function shen.lch")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp10561 := Call(__e, ns2_1set, symshen_4lch, tmp10554)


_ = tmp10561

tmp10562 := MakeNative(func(__e *ControlFlow) {
V1346 := __e.Get(1)
_ = V1346
tmp10625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1346)
}
__typedArg0 := V1346
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10574 Obj

if True == tmp10625 {
tmp10623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1346)
}
__typedArg0 := V1346
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp10623)
}
__typedArg0 := symwhere
__typedArg1 := tmp10623
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10576 Obj

if True == tmp10624 {
tmp10621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1346)
}
__typedArg0 := V1346
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10621)
}
__typedArg0 := tmp10621
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10578 Obj

if True == tmp10622 {
tmp10618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1346)
}
__typedArg0 := V1346
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10618)
}
__typedArg0 := tmp10618
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10619)
}
__typedArg0 := tmp10619
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10580 Obj

if True == tmp10620 {
tmp10614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1346)
}
__typedArg0 := V1346
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10614)
}
__typedArg0 := tmp10614
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10615)
}
__typedArg0 := tmp10615
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_a, tmp10616)
}
__typedArg0 := sym_a
__typedArg1 := tmp10616
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10582 Obj

if True == tmp10617 {
tmp10610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1346)
}
__typedArg0 := V1346
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10610)
}
__typedArg0 := tmp10610
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10611)
}
__typedArg0 := tmp10611
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10612)
}
__typedArg0 := tmp10612
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10584 Obj

if True == tmp10613 {
tmp10605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1346)
}
__typedArg0 := V1346
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10605)
}
__typedArg0 := tmp10605
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10606)
}
__typedArg0 := tmp10606
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10607)
}
__typedArg0 := tmp10607
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10608)
}
__typedArg0 := tmp10608
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10586 Obj

if True == tmp10609 {
tmp10599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1346)
}
__typedArg0 := V1346
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10599)
}
__typedArg0 := tmp10599
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10600)
}
__typedArg0 := tmp10600
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10601)
}
__typedArg0 := tmp10601
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10602)
}
__typedArg0 := tmp10602
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10603)
}
__typedArg0 := Nil
__typedArg1 := tmp10603
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10588 Obj

if True == tmp10604 {
tmp10596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1346)
}
__typedArg0 := V1346
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10596)
}
__typedArg0 := tmp10596
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10597)
}
__typedArg0 := tmp10597
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10590 Obj

if True == tmp10598 {
tmp10592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1346)
}
__typedArg0 := V1346
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10592)
}
__typedArg0 := tmp10592
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10593)
}
__typedArg0 := tmp10593
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10594)
}
__typedArg0 := Nil
__typedArg1 := tmp10594
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10591 Obj

if True == tmp10595 {
ifres10591 = True


} else {
ifres10591 = False


}

ifres10590 = ifres10591


} else {
ifres10590 = False


}

var ifres10589 Obj

if True == ifres10590 {
ifres10589 = True


} else {
ifres10589 = False


}

ifres10588 = ifres10589


} else {
ifres10588 = False


}

var ifres10587 Obj

if True == ifres10588 {
ifres10587 = True


} else {
ifres10587 = False


}

ifres10586 = ifres10587


} else {
ifres10586 = False


}

var ifres10585 Obj

if True == ifres10586 {
ifres10585 = True


} else {
ifres10585 = False


}

ifres10584 = ifres10585


} else {
ifres10584 = False


}

var ifres10583 Obj

if True == ifres10584 {
ifres10583 = True


} else {
ifres10583 = False


}

ifres10582 = ifres10583


} else {
ifres10582 = False


}

var ifres10581 Obj

if True == ifres10582 {
ifres10581 = True


} else {
ifres10581 = False


}

ifres10580 = ifres10581


} else {
ifres10580 = False


}

var ifres10579 Obj

if True == ifres10580 {
ifres10579 = True


} else {
ifres10579 = False


}

ifres10578 = ifres10579


} else {
ifres10578 = False


}

var ifres10577 Obj

if True == ifres10578 {
ifres10577 = True


} else {
ifres10577 = False


}

ifres10576 = ifres10577


} else {
ifres10576 = False


}

var ifres10575 Obj

if True == ifres10576 {
ifres10575 = True


} else {
ifres10575 = False


}

ifres10574 = ifres10575


} else {
ifres10574 = False


}

if True == ifres10574 {
tmp10564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_doccurs_d)
}
__typedArg0 := symshen_4_doccurs_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

var ifres10563 Obj

if True == tmp10564 {
ifres10563 = symis_b


} else {
ifres10563 = symis


}

tmp10565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1346)
}
__typedArg0 := V1346
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10565)
}
__typedArg0 := tmp10565
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10566)
}
__typedArg0 := tmp10566
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(ifres10563, tmp10567)
}
__typedArg0 := ifres10563
__typedArg1 := tmp10567
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1346)
}
__typedArg0 := V1346
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10569)
}
__typedArg0 := tmp10569
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10570)
}
__typedArg0 := tmp10570
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10572 := Call(__e, PrimFunc(symshen_4lchh), tmp10571)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10568, tmp10572)
}
__typedArg0 := tmp10568
__typedArg1 := tmp10572
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V1346)
return
}


}, 1)

tmp10626 := Call(__e, ns2_1set, symshen_4lchh, tmp10562)


_ = tmp10626

tmp10627 := MakeNative(func(__e *ControlFlow) {
V1347 := __e.Get(1)
_ = V1347
tmp10628 := MakeNative(func(__e *ControlFlow) {
W1348 := __e.Get(1)
_ = W1348
tmp10630 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1348)


if True == tmp10630 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1348)
return
}


}, 1)

tmp10631 := MakeNative(func(__e *ControlFlow) {
W1349 := __e.Get(1)
_ = W1349
tmp10657 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1349)


if True == tmp10657 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10632 := MakeNative(func(__e *ControlFlow) {
W1350 := __e.Get(1)
_ = W1350
tmp10633 := MakeNative(func(__e *ControlFlow) {
W1351 := __e.Get(1)
_ = W1351
tmp10653 := Call(__e, PrimFunc(symshen_4hds_a_2), W1351, sym_5_1_1)


if True == tmp10653 {
tmp10634 := MakeNative(func(__e *ControlFlow) {
W1352 := __e.Get(1)
_ = W1352
tmp10635 := MakeNative(func(__e *ControlFlow) {
W1353 := __e.Get(1)
_ = W1353
tmp10649 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1353)


if True == tmp10649 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10636 := MakeNative(func(__e *ControlFlow) {
W1354 := __e.Get(1)
_ = W1354
tmp10637 := MakeNative(func(__e *ControlFlow) {
W1355 := __e.Get(1)
_ = W1355
tmp10638 := MakeNative(func(__e *ControlFlow) {
W1356 := __e.Get(1)
_ = W1356
tmp10644 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1356)


if True == tmp10644 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10639 := MakeNative(func(__e *ControlFlow) {
W1357 := __e.Get(1)
_ = W1357
tmp10640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1354, Nil)
}
__typedArg0 := W1354
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1350, tmp10640)
}
__typedArg0 := W1350
__typedArg1 := tmp10640
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W1357, tmp10641)
return


}, 1)

tmp10642 := Call(__e, PrimFunc(symshen_4in_1_6), W1356)


__e.TailApply(tmp10639, tmp10642)
return


}


}, 1)

tmp10645 := Call(__e, PrimFunc(symshen_4_5sc_6), W1355)


__e.TailApply(tmp10638, tmp10645)
return


}, 1)

tmp10646 := Call(__e, PrimFunc(symshen_4in_1_6), W1353)


__e.TailApply(tmp10637, tmp10646)
return


}, 1)

tmp10647 := Call(__e, PrimFunc(symshen_4_5_1out), W1353)


__e.TailApply(tmp10636, tmp10647)
return


}


}, 1)

tmp10650 := Call(__e, PrimFunc(symshen_4_5body_6), W1352)


__e.TailApply(tmp10635, tmp10650)
return


}, 1)

tmp10651 := Call(__e, PrimFunc(symtail), W1351)


__e.TailApply(tmp10634, tmp10651)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10654 := Call(__e, PrimFunc(symshen_4in_1_6), W1349)


__e.TailApply(tmp10633, tmp10654)
return


}, 1)

tmp10655 := Call(__e, PrimFunc(symshen_4_5_1out), W1349)


__e.TailApply(tmp10632, tmp10655)
return


}


}, 1)

tmp10658 := Call(__e, PrimFunc(symshen_4_5head_6), V1347)


tmp10659 := Call(__e, tmp10631, tmp10658)


__e.TailApply(tmp10628, tmp10659)
return


}, 1)

tmp10660 := Call(__e, ns2_1set, symshen_4_5clause_6, tmp10627)


_ = tmp10660

tmp10661 := MakeNative(func(__e *ControlFlow) {
V1358 := __e.Get(1)
_ = V1358
tmp10662 := MakeNative(func(__e *ControlFlow) {
W1359 := __e.Get(1)
_ = W1359
tmp10674 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1359)


if True == tmp10674 {
tmp10663 := MakeNative(func(__e *ControlFlow) {
W1366 := __e.Get(1)
_ = W1366
tmp10665 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1366)


if True == tmp10665 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1366)
return
}


}, 1)

tmp10666 := MakeNative(func(__e *ControlFlow) {
W1367 := __e.Get(1)
_ = W1367
tmp10670 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1367)


if True == tmp10670 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10667 := MakeNative(func(__e *ControlFlow) {
W1368 := __e.Get(1)
_ = W1368
__e.TailApply(PrimFunc(symshen_4comb), W1368, Nil)
return
}, 1)

tmp10668 := Call(__e, PrimFunc(symshen_4in_1_6), W1367)


__e.TailApply(tmp10667, tmp10668)
return


}


}, 1)

tmp10671 := Call(__e, PrimFunc(sym_5e_6), V1358)


tmp10672 := Call(__e, tmp10666, tmp10671)


__e.TailApply(tmp10663, tmp10672)
return


} else {
__e.Return(W1359)
return
}


}, 1)

tmp10675 := MakeNative(func(__e *ControlFlow) {
W1360 := __e.Get(1)
_ = W1360
tmp10690 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1360)


if True == tmp10690 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10676 := MakeNative(func(__e *ControlFlow) {
W1361 := __e.Get(1)
_ = W1361
tmp10677 := MakeNative(func(__e *ControlFlow) {
W1362 := __e.Get(1)
_ = W1362
tmp10678 := MakeNative(func(__e *ControlFlow) {
W1363 := __e.Get(1)
_ = W1363
tmp10685 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1363)


if True == tmp10685 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10679 := MakeNative(func(__e *ControlFlow) {
W1364 := __e.Get(1)
_ = W1364
tmp10680 := MakeNative(func(__e *ControlFlow) {
W1365 := __e.Get(1)
_ = W1365
tmp10681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1361, W1364)
}
__typedArg0 := W1361
__typedArg1 := W1364
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W1365, tmp10681)
return


}, 1)

tmp10682 := Call(__e, PrimFunc(symshen_4in_1_6), W1363)


__e.TailApply(tmp10680, tmp10682)
return


}, 1)

tmp10683 := Call(__e, PrimFunc(symshen_4_5_1out), W1363)


__e.TailApply(tmp10679, tmp10683)
return


}


}, 1)

tmp10686 := Call(__e, PrimFunc(symshen_4_5head_6), W1362)


__e.TailApply(tmp10678, tmp10686)
return


}, 1)

tmp10687 := Call(__e, PrimFunc(symshen_4in_1_6), W1360)


__e.TailApply(tmp10677, tmp10687)
return


}, 1)

tmp10688 := Call(__e, PrimFunc(symshen_4_5_1out), W1360)


__e.TailApply(tmp10676, tmp10688)
return


}


}, 1)

tmp10691 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1358)


tmp10692 := Call(__e, tmp10675, tmp10691)


__e.TailApply(tmp10662, tmp10692)
return


}, 1)

tmp10693 := Call(__e, ns2_1set, symshen_4_5head_6, tmp10661)


_ = tmp10693

tmp10694 := MakeNative(func(__e *ControlFlow) {
V1369 := __e.Get(1)
_ = V1369
tmp10695 := MakeNative(func(__e *ControlFlow) {
W1370 := __e.Get(1)
_ = W1370
tmp10883 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1370)


if True == tmp10883 {
tmp10696 := MakeNative(func(__e *ControlFlow) {
W1373 := __e.Get(1)
_ = W1373
tmp10870 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1373)


if True == tmp10870 {
tmp10697 := MakeNative(func(__e *ControlFlow) {
W1376 := __e.Get(1)
_ = W1376
tmp10831 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1376)


if True == tmp10831 {
tmp10698 := MakeNative(func(__e *ControlFlow) {
W1388 := __e.Get(1)
_ = W1388
tmp10801 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1388)


if True == tmp10801 {
tmp10699 := MakeNative(func(__e *ControlFlow) {
W1397 := __e.Get(1)
_ = W1397
tmp10771 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1397)


if True == tmp10771 {
tmp10700 := MakeNative(func(__e *ControlFlow) {
W1406 := __e.Get(1)
_ = W1406
tmp10737 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1406)


if True == tmp10737 {
tmp10701 := MakeNative(func(__e *ControlFlow) {
W1416 := __e.Get(1)
_ = W1416
tmp10703 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1416)


if True == tmp10703 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1416)
return
}


}, 1)

tmp10735 := Call(__e, PrimFunc(symshen_4ccons_2), V1369)


var ifres10704 Obj

if True == tmp10735 {
tmp10705 := MakeNative(func(__e *ControlFlow) {
W1417 := __e.Get(1)
_ = W1417
tmp10706 := MakeNative(func(__e *ControlFlow) {
W1418 := __e.Get(1)
_ = W1418
tmp10730 := Call(__e, PrimFunc(symshen_4hds_a_2), W1417, symmode)


if True == tmp10730 {
tmp10707 := MakeNative(func(__e *ControlFlow) {
W1419 := __e.Get(1)
_ = W1419
tmp10708 := MakeNative(func(__e *ControlFlow) {
W1420 := __e.Get(1)
_ = W1420
tmp10726 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1420)


if True == tmp10726 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10709 := MakeNative(func(__e *ControlFlow) {
W1421 := __e.Get(1)
_ = W1421
tmp10710 := MakeNative(func(__e *ControlFlow) {
W1422 := __e.Get(1)
_ = W1422
tmp10722 := Call(__e, PrimFunc(symshen_4hds_a_2), W1422, sym_1)


if True == tmp10722 {
tmp10711 := MakeNative(func(__e *ControlFlow) {
W1423 := __e.Get(1)
_ = W1423
tmp10712 := MakeNative(func(__e *ControlFlow) {
W1424 := __e.Get(1)
_ = W1424
tmp10718 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1424)


if True == tmp10718 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10713 := MakeNative(func(__e *ControlFlow) {
W1425 := __e.Get(1)
_ = W1425
tmp10714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1421, Nil)
}
__typedArg0 := W1421
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_1m, tmp10714)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp10714
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W1418, tmp10715)
return


}, 1)

tmp10716 := Call(__e, PrimFunc(symshen_4in_1_6), W1424)


__e.TailApply(tmp10713, tmp10716)
return


}


}, 1)

tmp10719 := Call(__e, PrimFunc(sym_5end_6), W1423)


__e.TailApply(tmp10712, tmp10719)
return


}, 1)

tmp10720 := Call(__e, PrimFunc(symtail), W1422)


__e.TailApply(tmp10711, tmp10720)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10723 := Call(__e, PrimFunc(symshen_4in_1_6), W1420)


__e.TailApply(tmp10710, tmp10723)
return


}, 1)

tmp10724 := Call(__e, PrimFunc(symshen_4_5_1out), W1420)


__e.TailApply(tmp10709, tmp10724)
return


}


}, 1)

tmp10727 := Call(__e, PrimFunc(symshen_4_5hterm_6), W1419)


__e.TailApply(tmp10708, tmp10727)
return


}, 1)

tmp10728 := Call(__e, PrimFunc(symtail), W1417)


__e.TailApply(tmp10707, tmp10728)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10731 := Call(__e, PrimFunc(symtail), V1369)


__e.TailApply(tmp10706, tmp10731)
return


}, 1)

tmp10732 := Call(__e, PrimFunc(symhead), V1369)


tmp10733 := Call(__e, tmp10705, tmp10732)


ifres10704 = tmp10733


} else {
tmp10734 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10704 = tmp10734


}

__e.TailApply(tmp10701, ifres10704)
return


} else {
__e.Return(W1406)
return
}


}, 1)

tmp10769 := Call(__e, PrimFunc(symshen_4ccons_2), V1369)


var ifres10738 Obj

if True == tmp10769 {
tmp10739 := MakeNative(func(__e *ControlFlow) {
W1407 := __e.Get(1)
_ = W1407
tmp10740 := MakeNative(func(__e *ControlFlow) {
W1408 := __e.Get(1)
_ = W1408
tmp10764 := Call(__e, PrimFunc(symshen_4hds_a_2), W1407, symmode)


if True == tmp10764 {
tmp10741 := MakeNative(func(__e *ControlFlow) {
W1409 := __e.Get(1)
_ = W1409
tmp10742 := MakeNative(func(__e *ControlFlow) {
W1410 := __e.Get(1)
_ = W1410
tmp10760 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1410)


if True == tmp10760 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10743 := MakeNative(func(__e *ControlFlow) {
W1411 := __e.Get(1)
_ = W1411
tmp10744 := MakeNative(func(__e *ControlFlow) {
W1412 := __e.Get(1)
_ = W1412
tmp10756 := Call(__e, PrimFunc(symshen_4hds_a_2), W1412, sym_7)


if True == tmp10756 {
tmp10745 := MakeNative(func(__e *ControlFlow) {
W1413 := __e.Get(1)
_ = W1413
tmp10746 := MakeNative(func(__e *ControlFlow) {
W1414 := __e.Get(1)
_ = W1414
tmp10752 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1414)


if True == tmp10752 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10747 := MakeNative(func(__e *ControlFlow) {
W1415 := __e.Get(1)
_ = W1415
tmp10748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1411, Nil)
}
__typedArg0 := W1411
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_7m, tmp10748)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp10748
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W1408, tmp10749)
return


}, 1)

tmp10750 := Call(__e, PrimFunc(symshen_4in_1_6), W1414)


__e.TailApply(tmp10747, tmp10750)
return


}


}, 1)

tmp10753 := Call(__e, PrimFunc(sym_5end_6), W1413)


__e.TailApply(tmp10746, tmp10753)
return


}, 1)

tmp10754 := Call(__e, PrimFunc(symtail), W1412)


__e.TailApply(tmp10745, tmp10754)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10757 := Call(__e, PrimFunc(symshen_4in_1_6), W1410)


__e.TailApply(tmp10744, tmp10757)
return


}, 1)

tmp10758 := Call(__e, PrimFunc(symshen_4_5_1out), W1410)


__e.TailApply(tmp10743, tmp10758)
return


}


}, 1)

tmp10761 := Call(__e, PrimFunc(symshen_4_5hterm_6), W1409)


__e.TailApply(tmp10742, tmp10761)
return


}, 1)

tmp10762 := Call(__e, PrimFunc(symtail), W1407)


__e.TailApply(tmp10741, tmp10762)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10765 := Call(__e, PrimFunc(symtail), V1369)


__e.TailApply(tmp10740, tmp10765)
return


}, 1)

tmp10766 := Call(__e, PrimFunc(symhead), V1369)


tmp10767 := Call(__e, tmp10739, tmp10766)


ifres10738 = tmp10767


} else {
tmp10768 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10738 = tmp10768


}

__e.TailApply(tmp10700, ifres10738)
return


} else {
__e.Return(W1397)
return
}


}, 1)

tmp10799 := Call(__e, PrimFunc(symshen_4ccons_2), V1369)


var ifres10772 Obj

if True == tmp10799 {
tmp10773 := MakeNative(func(__e *ControlFlow) {
W1398 := __e.Get(1)
_ = W1398
tmp10774 := MakeNative(func(__e *ControlFlow) {
W1399 := __e.Get(1)
_ = W1399
tmp10794 := Call(__e, PrimFunc(symshen_4hds_a_2), W1398, sym_1)


if True == tmp10794 {
tmp10775 := MakeNative(func(__e *ControlFlow) {
W1400 := __e.Get(1)
_ = W1400
tmp10776 := MakeNative(func(__e *ControlFlow) {
W1401 := __e.Get(1)
_ = W1401
tmp10790 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1401)


if True == tmp10790 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10777 := MakeNative(func(__e *ControlFlow) {
W1402 := __e.Get(1)
_ = W1402
tmp10778 := MakeNative(func(__e *ControlFlow) {
W1403 := __e.Get(1)
_ = W1403
tmp10779 := MakeNative(func(__e *ControlFlow) {
W1404 := __e.Get(1)
_ = W1404
tmp10785 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1404)


if True == tmp10785 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10780 := MakeNative(func(__e *ControlFlow) {
W1405 := __e.Get(1)
_ = W1405
tmp10781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1402, Nil)
}
__typedArg0 := W1402
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_1m, tmp10781)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp10781
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W1399, tmp10782)
return


}, 1)

tmp10783 := Call(__e, PrimFunc(symshen_4in_1_6), W1404)


__e.TailApply(tmp10780, tmp10783)
return


}


}, 1)

tmp10786 := Call(__e, PrimFunc(sym_5end_6), W1403)


__e.TailApply(tmp10779, tmp10786)
return


}, 1)

tmp10787 := Call(__e, PrimFunc(symshen_4in_1_6), W1401)


__e.TailApply(tmp10778, tmp10787)
return


}, 1)

tmp10788 := Call(__e, PrimFunc(symshen_4_5_1out), W1401)


__e.TailApply(tmp10777, tmp10788)
return


}


}, 1)

tmp10791 := Call(__e, PrimFunc(symshen_4_5hterm_6), W1400)


__e.TailApply(tmp10776, tmp10791)
return


}, 1)

tmp10792 := Call(__e, PrimFunc(symtail), W1398)


__e.TailApply(tmp10775, tmp10792)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10795 := Call(__e, PrimFunc(symtail), V1369)


__e.TailApply(tmp10774, tmp10795)
return


}, 1)

tmp10796 := Call(__e, PrimFunc(symhead), V1369)


tmp10797 := Call(__e, tmp10773, tmp10796)


ifres10772 = tmp10797


} else {
tmp10798 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10772 = tmp10798


}

__e.TailApply(tmp10699, ifres10772)
return


} else {
__e.Return(W1388)
return
}


}, 1)

tmp10829 := Call(__e, PrimFunc(symshen_4ccons_2), V1369)


var ifres10802 Obj

if True == tmp10829 {
tmp10803 := MakeNative(func(__e *ControlFlow) {
W1389 := __e.Get(1)
_ = W1389
tmp10804 := MakeNative(func(__e *ControlFlow) {
W1390 := __e.Get(1)
_ = W1390
tmp10824 := Call(__e, PrimFunc(symshen_4hds_a_2), W1389, sym_7)


if True == tmp10824 {
tmp10805 := MakeNative(func(__e *ControlFlow) {
W1391 := __e.Get(1)
_ = W1391
tmp10806 := MakeNative(func(__e *ControlFlow) {
W1392 := __e.Get(1)
_ = W1392
tmp10820 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1392)


if True == tmp10820 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10807 := MakeNative(func(__e *ControlFlow) {
W1393 := __e.Get(1)
_ = W1393
tmp10808 := MakeNative(func(__e *ControlFlow) {
W1394 := __e.Get(1)
_ = W1394
tmp10809 := MakeNative(func(__e *ControlFlow) {
W1395 := __e.Get(1)
_ = W1395
tmp10815 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1395)


if True == tmp10815 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10810 := MakeNative(func(__e *ControlFlow) {
W1396 := __e.Get(1)
_ = W1396
tmp10811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1393, Nil)
}
__typedArg0 := W1393
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_7m, tmp10811)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp10811
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W1390, tmp10812)
return


}, 1)

tmp10813 := Call(__e, PrimFunc(symshen_4in_1_6), W1395)


__e.TailApply(tmp10810, tmp10813)
return


}


}, 1)

tmp10816 := Call(__e, PrimFunc(sym_5end_6), W1394)


__e.TailApply(tmp10809, tmp10816)
return


}, 1)

tmp10817 := Call(__e, PrimFunc(symshen_4in_1_6), W1392)


__e.TailApply(tmp10808, tmp10817)
return


}, 1)

tmp10818 := Call(__e, PrimFunc(symshen_4_5_1out), W1392)


__e.TailApply(tmp10807, tmp10818)
return


}


}, 1)

tmp10821 := Call(__e, PrimFunc(symshen_4_5hterm_6), W1391)


__e.TailApply(tmp10806, tmp10821)
return


}, 1)

tmp10822 := Call(__e, PrimFunc(symtail), W1389)


__e.TailApply(tmp10805, tmp10822)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10825 := Call(__e, PrimFunc(symtail), V1369)


__e.TailApply(tmp10804, tmp10825)
return


}, 1)

tmp10826 := Call(__e, PrimFunc(symhead), V1369)


tmp10827 := Call(__e, tmp10803, tmp10826)


ifres10802 = tmp10827


} else {
tmp10828 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10802 = tmp10828


}

__e.TailApply(tmp10698, ifres10802)
return


} else {
__e.Return(W1376)
return
}


}, 1)

tmp10868 := Call(__e, PrimFunc(symshen_4ccons_2), V1369)


var ifres10832 Obj

if True == tmp10868 {
tmp10833 := MakeNative(func(__e *ControlFlow) {
W1377 := __e.Get(1)
_ = W1377
tmp10834 := MakeNative(func(__e *ControlFlow) {
W1378 := __e.Get(1)
_ = W1378
tmp10863 := Call(__e, PrimFunc(symshen_4hds_a_2), W1377, symcons)


if True == tmp10863 {
tmp10835 := MakeNative(func(__e *ControlFlow) {
W1379 := __e.Get(1)
_ = W1379
tmp10836 := MakeNative(func(__e *ControlFlow) {
W1380 := __e.Get(1)
_ = W1380
tmp10859 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1380)


if True == tmp10859 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10837 := MakeNative(func(__e *ControlFlow) {
W1381 := __e.Get(1)
_ = W1381
tmp10838 := MakeNative(func(__e *ControlFlow) {
W1382 := __e.Get(1)
_ = W1382
tmp10839 := MakeNative(func(__e *ControlFlow) {
W1383 := __e.Get(1)
_ = W1383
tmp10854 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1383)


if True == tmp10854 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10840 := MakeNative(func(__e *ControlFlow) {
W1384 := __e.Get(1)
_ = W1384
tmp10841 := MakeNative(func(__e *ControlFlow) {
W1385 := __e.Get(1)
_ = W1385
tmp10842 := MakeNative(func(__e *ControlFlow) {
W1386 := __e.Get(1)
_ = W1386
tmp10849 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1386)


if True == tmp10849 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10843 := MakeNative(func(__e *ControlFlow) {
W1387 := __e.Get(1)
_ = W1387
tmp10844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1384, Nil)
}
__typedArg0 := W1384
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1381, tmp10844)
}
__typedArg0 := W1381
__typedArg1 := tmp10844
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp10845)
}
__typedArg0 := symcons
__typedArg1 := tmp10845
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W1378, tmp10846)
return


}, 1)

tmp10847 := Call(__e, PrimFunc(symshen_4in_1_6), W1386)


__e.TailApply(tmp10843, tmp10847)
return


}


}, 1)

tmp10850 := Call(__e, PrimFunc(sym_5end_6), W1385)


__e.TailApply(tmp10842, tmp10850)
return


}, 1)

tmp10851 := Call(__e, PrimFunc(symshen_4in_1_6), W1383)


__e.TailApply(tmp10841, tmp10851)
return


}, 1)

tmp10852 := Call(__e, PrimFunc(symshen_4_5_1out), W1383)


__e.TailApply(tmp10840, tmp10852)
return


}


}, 1)

tmp10855 := Call(__e, PrimFunc(symshen_4_5hterm2_6), W1382)


__e.TailApply(tmp10839, tmp10855)
return


}, 1)

tmp10856 := Call(__e, PrimFunc(symshen_4in_1_6), W1380)


__e.TailApply(tmp10838, tmp10856)
return


}, 1)

tmp10857 := Call(__e, PrimFunc(symshen_4_5_1out), W1380)


__e.TailApply(tmp10837, tmp10857)
return


}


}, 1)

tmp10860 := Call(__e, PrimFunc(symshen_4_5hterm1_6), W1379)


__e.TailApply(tmp10836, tmp10860)
return


}, 1)

tmp10861 := Call(__e, PrimFunc(symtail), W1377)


__e.TailApply(tmp10835, tmp10861)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10864 := Call(__e, PrimFunc(symtail), V1369)


__e.TailApply(tmp10834, tmp10864)
return


}, 1)

tmp10865 := Call(__e, PrimFunc(symhead), V1369)


tmp10866 := Call(__e, tmp10833, tmp10865)


ifres10832 = tmp10866


} else {
tmp10867 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10832 = tmp10867


}

__e.TailApply(tmp10697, ifres10832)
return


} else {
__e.Return(W1373)
return
}


}, 1)

tmp10881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1369)
}
__typedArg0 := V1369
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10871 Obj

if True == tmp10881 {
tmp10872 := MakeNative(func(__e *ControlFlow) {
W1374 := __e.Get(1)
_ = W1374
tmp10873 := MakeNative(func(__e *ControlFlow) {
W1375 := __e.Get(1)
_ = W1375
tmp10875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp10876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W1374, tmp10875)
}
__typedArg0 := W1374
__typedArg1 := tmp10875
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10876 {
__e.TailApply(PrimFunc(symshen_4comb), W1375, W1374)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10877 := Call(__e, PrimFunc(symtail), V1369)


__e.TailApply(tmp10873, tmp10877)
return


}, 1)

tmp10878 := Call(__e, PrimFunc(symhead), V1369)


tmp10879 := Call(__e, tmp10872, tmp10878)


ifres10871 = tmp10879


} else {
tmp10880 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10871 = tmp10880


}

__e.TailApply(tmp10696, ifres10871)
return


} else {
__e.Return(W1370)
return
}


}, 1)

tmp10897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1369)
}
__typedArg0 := V1369
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10884 Obj

if True == tmp10897 {
tmp10885 := MakeNative(func(__e *ControlFlow) {
W1371 := __e.Get(1)
_ = W1371
tmp10886 := MakeNative(func(__e *ControlFlow) {
W1372 := __e.Get(1)
_ = W1372
tmp10892 := Call(__e, PrimFunc(symatom_2), W1371)


var ifres10888 Obj

if True == tmp10892 {
tmp10890 := Call(__e, PrimFunc(symshen_4prolog_1keyword_2), W1371)


tmp10891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp10890)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp10890
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres10889 Obj

if True == tmp10891 {
ifres10889 = True


} else {
ifres10889 = False


}

ifres10888 = ifres10889


} else {
ifres10888 = False


}

if True == ifres10888 {
__e.TailApply(PrimFunc(symshen_4comb), W1372, W1371)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10893 := Call(__e, PrimFunc(symtail), V1369)


__e.TailApply(tmp10886, tmp10893)
return


}, 1)

tmp10894 := Call(__e, PrimFunc(symhead), V1369)


tmp10895 := Call(__e, tmp10885, tmp10894)


ifres10884 = tmp10895


} else {
tmp10896 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10884 = tmp10896


}

__e.TailApply(tmp10695, ifres10884)
return


}, 1)

tmp10898 := Call(__e, ns2_1set, symshen_4_5hterm_6, tmp10694)


_ = tmp10898

tmp10899 := MakeNative(func(__e *ControlFlow) {
V1426 := __e.Get(1)
_ = V1426
tmp10900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp10901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10900, tmp10901)
}
__typedArg0 := tmp10900
__typedArg1 := tmp10901
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symelement_2), V1426, tmp10902)
return


}, 1)

tmp10903 := Call(__e, ns2_1set, symshen_4prolog_1keyword_2, tmp10899)


_ = tmp10903

tmp10904 := MakeNative(func(__e *ControlFlow) {
V1427 := __e.Get(1)
_ = V1427
tmp10917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V1427)
}
__typedArg0 := V1427
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp10917 {
__e.Return(True)
return
} else {
tmp10915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V1427)
}
__typedArg0 := V1427
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres10906 Obj

if True == tmp10915 {
ifres10906 = True


} else {
tmp10914 := Call(__e, PrimFunc(symboolean_2), V1427)


var ifres10908 Obj

if True == tmp10914 {
ifres10908 = True


} else {
tmp10913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(V1427)
}
__typedArg0 := V1427
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

var ifres10910 Obj

if True == tmp10913 {
ifres10910 = True


} else {
tmp10912 := Call(__e, PrimFunc(symempty_2), V1427)


var ifres10911 Obj

if True == tmp10912 {
ifres10911 = True


} else {
ifres10911 = False


}

ifres10910 = ifres10911


}

var ifres10909 Obj

if True == ifres10910 {
ifres10909 = True


} else {
ifres10909 = False


}

ifres10908 = ifres10909


}

var ifres10907 Obj

if True == ifres10908 {
ifres10907 = True


} else {
ifres10907 = False


}

ifres10906 = ifres10907


}

if True == ifres10906 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp10918 := Call(__e, ns2_1set, symatom_2, tmp10904)


_ = tmp10918

tmp10919 := MakeNative(func(__e *ControlFlow) {
V1428 := __e.Get(1)
_ = V1428
tmp10920 := MakeNative(func(__e *ControlFlow) {
W1429 := __e.Get(1)
_ = W1429
tmp10922 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1429)


if True == tmp10922 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1429)
return
}


}, 1)

tmp10923 := MakeNative(func(__e *ControlFlow) {
W1430 := __e.Get(1)
_ = W1430
tmp10929 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1430)


if True == tmp10929 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10924 := MakeNative(func(__e *ControlFlow) {
W1431 := __e.Get(1)
_ = W1431
tmp10925 := MakeNative(func(__e *ControlFlow) {
W1432 := __e.Get(1)
_ = W1432
__e.TailApply(PrimFunc(symshen_4comb), W1432, W1431)
return
}, 1)

tmp10926 := Call(__e, PrimFunc(symshen_4in_1_6), W1430)


__e.TailApply(tmp10925, tmp10926)
return


}, 1)

tmp10927 := Call(__e, PrimFunc(symshen_4_5_1out), W1430)


__e.TailApply(tmp10924, tmp10927)
return


}


}, 1)

tmp10930 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1428)


tmp10931 := Call(__e, tmp10923, tmp10930)


__e.TailApply(tmp10920, tmp10931)
return


}, 1)

tmp10932 := Call(__e, ns2_1set, symshen_4_5hterm1_6, tmp10919)


_ = tmp10932

tmp10933 := MakeNative(func(__e *ControlFlow) {
V1433 := __e.Get(1)
_ = V1433
tmp10934 := MakeNative(func(__e *ControlFlow) {
W1434 := __e.Get(1)
_ = W1434
tmp10936 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1434)


if True == tmp10936 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1434)
return
}


}, 1)

tmp10937 := MakeNative(func(__e *ControlFlow) {
W1435 := __e.Get(1)
_ = W1435
tmp10943 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1435)


if True == tmp10943 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10938 := MakeNative(func(__e *ControlFlow) {
W1436 := __e.Get(1)
_ = W1436
tmp10939 := MakeNative(func(__e *ControlFlow) {
W1437 := __e.Get(1)
_ = W1437
__e.TailApply(PrimFunc(symshen_4comb), W1437, W1436)
return
}, 1)

tmp10940 := Call(__e, PrimFunc(symshen_4in_1_6), W1435)


__e.TailApply(tmp10939, tmp10940)
return


}, 1)

tmp10941 := Call(__e, PrimFunc(symshen_4_5_1out), W1435)


__e.TailApply(tmp10938, tmp10941)
return


}


}, 1)

tmp10944 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1433)


tmp10945 := Call(__e, tmp10937, tmp10944)


__e.TailApply(tmp10934, tmp10945)
return


}, 1)

tmp10946 := Call(__e, ns2_1set, symshen_4_5hterm2_6, tmp10933)


_ = tmp10946

tmp10947 := MakeNative(func(__e *ControlFlow) {
V1438 := __e.Get(1)
_ = V1438
tmp10948 := MakeNative(func(__e *ControlFlow) {
W1439 := __e.Get(1)
_ = W1439
tmp10960 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1439)


if True == tmp10960 {
tmp10949 := MakeNative(func(__e *ControlFlow) {
W1446 := __e.Get(1)
_ = W1446
tmp10951 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1446)


if True == tmp10951 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1446)
return
}


}, 1)

tmp10952 := MakeNative(func(__e *ControlFlow) {
W1447 := __e.Get(1)
_ = W1447
tmp10956 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1447)


if True == tmp10956 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10953 := MakeNative(func(__e *ControlFlow) {
W1448 := __e.Get(1)
_ = W1448
__e.TailApply(PrimFunc(symshen_4comb), W1448, Nil)
return
}, 1)

tmp10954 := Call(__e, PrimFunc(symshen_4in_1_6), W1447)


__e.TailApply(tmp10953, tmp10954)
return


}


}, 1)

tmp10957 := Call(__e, PrimFunc(sym_5e_6), V1438)


tmp10958 := Call(__e, tmp10952, tmp10957)


__e.TailApply(tmp10949, tmp10958)
return


} else {
__e.Return(W1439)
return
}


}, 1)

tmp10961 := MakeNative(func(__e *ControlFlow) {
W1440 := __e.Get(1)
_ = W1440
tmp10976 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1440)


if True == tmp10976 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10962 := MakeNative(func(__e *ControlFlow) {
W1441 := __e.Get(1)
_ = W1441
tmp10963 := MakeNative(func(__e *ControlFlow) {
W1442 := __e.Get(1)
_ = W1442
tmp10964 := MakeNative(func(__e *ControlFlow) {
W1443 := __e.Get(1)
_ = W1443
tmp10971 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1443)


if True == tmp10971 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10965 := MakeNative(func(__e *ControlFlow) {
W1444 := __e.Get(1)
_ = W1444
tmp10966 := MakeNative(func(__e *ControlFlow) {
W1445 := __e.Get(1)
_ = W1445
tmp10967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1441, W1444)
}
__typedArg0 := W1441
__typedArg1 := W1444
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W1445, tmp10967)
return


}, 1)

tmp10968 := Call(__e, PrimFunc(symshen_4in_1_6), W1443)


__e.TailApply(tmp10966, tmp10968)
return


}, 1)

tmp10969 := Call(__e, PrimFunc(symshen_4_5_1out), W1443)


__e.TailApply(tmp10965, tmp10969)
return


}


}, 1)

tmp10972 := Call(__e, PrimFunc(symshen_4_5body_6), W1442)


__e.TailApply(tmp10964, tmp10972)
return


}, 1)

tmp10973 := Call(__e, PrimFunc(symshen_4in_1_6), W1440)


__e.TailApply(tmp10963, tmp10973)
return


}, 1)

tmp10974 := Call(__e, PrimFunc(symshen_4_5_1out), W1440)


__e.TailApply(tmp10962, tmp10974)
return


}


}, 1)

tmp10977 := Call(__e, PrimFunc(symshen_4_5literal_6), V1438)


tmp10978 := Call(__e, tmp10961, tmp10977)


__e.TailApply(tmp10948, tmp10978)
return


}, 1)

tmp10979 := Call(__e, ns2_1set, symshen_4_5body_6, tmp10947)


_ = tmp10979

tmp10980 := MakeNative(func(__e *ControlFlow) {
V1449 := __e.Get(1)
_ = V1449
tmp10981 := MakeNative(func(__e *ControlFlow) {
W1450 := __e.Get(1)
_ = W1450
tmp11008 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1450)


if True == tmp11008 {
tmp10982 := MakeNative(func(__e *ControlFlow) {
W1452 := __e.Get(1)
_ = W1452
tmp10984 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1452)


if True == tmp10984 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1452)
return
}


}, 1)

tmp11006 := Call(__e, PrimFunc(symshen_4ccons_2), V1449)


var ifres10985 Obj

if True == tmp11006 {
tmp10986 := MakeNative(func(__e *ControlFlow) {
W1453 := __e.Get(1)
_ = W1453
tmp10987 := MakeNative(func(__e *ControlFlow) {
W1454 := __e.Get(1)
_ = W1454
tmp10988 := MakeNative(func(__e *ControlFlow) {
W1455 := __e.Get(1)
_ = W1455
tmp11000 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1455)


if True == tmp11000 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10989 := MakeNative(func(__e *ControlFlow) {
W1456 := __e.Get(1)
_ = W1456
tmp10990 := MakeNative(func(__e *ControlFlow) {
W1457 := __e.Get(1)
_ = W1457
tmp10991 := MakeNative(func(__e *ControlFlow) {
W1458 := __e.Get(1)
_ = W1458
tmp10995 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1458)


if True == tmp10995 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10992 := MakeNative(func(__e *ControlFlow) {
W1459 := __e.Get(1)
_ = W1459
__e.TailApply(PrimFunc(symshen_4comb), W1454, W1456)
return
}, 1)

tmp10993 := Call(__e, PrimFunc(symshen_4in_1_6), W1458)


__e.TailApply(tmp10992, tmp10993)
return


}


}, 1)

tmp10996 := Call(__e, PrimFunc(sym_5end_6), W1457)


__e.TailApply(tmp10991, tmp10996)
return


}, 1)

tmp10997 := Call(__e, PrimFunc(symshen_4in_1_6), W1455)


__e.TailApply(tmp10990, tmp10997)
return


}, 1)

tmp10998 := Call(__e, PrimFunc(symshen_4_5_1out), W1455)


__e.TailApply(tmp10989, tmp10998)
return


}


}, 1)

tmp11001 := Call(__e, PrimFunc(symshen_4_5bterms_6), W1453)


__e.TailApply(tmp10988, tmp11001)
return


}, 1)

tmp11002 := Call(__e, PrimFunc(symtail), V1449)


__e.TailApply(tmp10987, tmp11002)
return


}, 1)

tmp11003 := Call(__e, PrimFunc(symhead), V1449)


tmp11004 := Call(__e, tmp10986, tmp11003)


ifres10985 = tmp11004


} else {
tmp11005 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10985 = tmp11005


}

__e.TailApply(tmp10982, ifres10985)
return


} else {
__e.Return(W1450)
return
}


}, 1)

tmp11014 := Call(__e, PrimFunc(symshen_4hds_a_2), V1449, sym_b)


var ifres11009 Obj

if True == tmp11014 {
tmp11010 := MakeNative(func(__e *ControlFlow) {
W1451 := __e.Get(1)
_ = W1451
__e.TailApply(PrimFunc(symshen_4comb), W1451, sym_b)
return
}, 1)

tmp11011 := Call(__e, PrimFunc(symtail), V1449)


tmp11012 := Call(__e, tmp11010, tmp11011)


ifres11009 = tmp11012


} else {
tmp11013 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11009 = tmp11013


}

__e.TailApply(tmp10981, ifres11009)
return


}, 1)

tmp11015 := Call(__e, ns2_1set, symshen_4_5literal_6, tmp10980)


_ = tmp11015

tmp11016 := MakeNative(func(__e *ControlFlow) {
V1460 := __e.Get(1)
_ = V1460
tmp11017 := MakeNative(func(__e *ControlFlow) {
W1461 := __e.Get(1)
_ = W1461
tmp11029 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1461)


if True == tmp11029 {
tmp11018 := MakeNative(func(__e *ControlFlow) {
W1468 := __e.Get(1)
_ = W1468
tmp11020 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1468)


if True == tmp11020 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1468)
return
}


}, 1)

tmp11021 := MakeNative(func(__e *ControlFlow) {
W1469 := __e.Get(1)
_ = W1469
tmp11025 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1469)


if True == tmp11025 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11022 := MakeNative(func(__e *ControlFlow) {
W1470 := __e.Get(1)
_ = W1470
__e.TailApply(PrimFunc(symshen_4comb), W1470, Nil)
return
}, 1)

tmp11023 := Call(__e, PrimFunc(symshen_4in_1_6), W1469)


__e.TailApply(tmp11022, tmp11023)
return


}


}, 1)

tmp11026 := Call(__e, PrimFunc(sym_5e_6), V1460)


tmp11027 := Call(__e, tmp11021, tmp11026)


__e.TailApply(tmp11018, tmp11027)
return


} else {
__e.Return(W1461)
return
}


}, 1)

tmp11030 := MakeNative(func(__e *ControlFlow) {
W1462 := __e.Get(1)
_ = W1462
tmp11045 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1462)


if True == tmp11045 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11031 := MakeNative(func(__e *ControlFlow) {
W1463 := __e.Get(1)
_ = W1463
tmp11032 := MakeNative(func(__e *ControlFlow) {
W1464 := __e.Get(1)
_ = W1464
tmp11033 := MakeNative(func(__e *ControlFlow) {
W1465 := __e.Get(1)
_ = W1465
tmp11040 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1465)


if True == tmp11040 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11034 := MakeNative(func(__e *ControlFlow) {
W1466 := __e.Get(1)
_ = W1466
tmp11035 := MakeNative(func(__e *ControlFlow) {
W1467 := __e.Get(1)
_ = W1467
tmp11036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1463, W1466)
}
__typedArg0 := W1463
__typedArg1 := W1466
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W1467, tmp11036)
return


}, 1)

tmp11037 := Call(__e, PrimFunc(symshen_4in_1_6), W1465)


__e.TailApply(tmp11035, tmp11037)
return


}, 1)

tmp11038 := Call(__e, PrimFunc(symshen_4_5_1out), W1465)


__e.TailApply(tmp11034, tmp11038)
return


}


}, 1)

tmp11041 := Call(__e, PrimFunc(symshen_4_5bterms_6), W1464)


__e.TailApply(tmp11033, tmp11041)
return


}, 1)

tmp11042 := Call(__e, PrimFunc(symshen_4in_1_6), W1462)


__e.TailApply(tmp11032, tmp11042)
return


}, 1)

tmp11043 := Call(__e, PrimFunc(symshen_4_5_1out), W1462)


__e.TailApply(tmp11031, tmp11043)
return


}


}, 1)

tmp11046 := Call(__e, PrimFunc(symshen_4_5bterm_6), V1460)


tmp11047 := Call(__e, tmp11030, tmp11046)


__e.TailApply(tmp11017, tmp11047)
return


}, 1)

tmp11048 := Call(__e, ns2_1set, symshen_4_5bterms_6, tmp11016)


_ = tmp11048

tmp11049 := MakeNative(func(__e *ControlFlow) {
V1471 := __e.Get(1)
_ = V1471
tmp11050 := MakeNative(func(__e *ControlFlow) {
W1472 := __e.Get(1)
_ = W1472
tmp11090 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1472)


if True == tmp11090 {
tmp11051 := MakeNative(func(__e *ControlFlow) {
W1476 := __e.Get(1)
_ = W1476
tmp11078 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1476)


if True == tmp11078 {
tmp11052 := MakeNative(func(__e *ControlFlow) {
W1479 := __e.Get(1)
_ = W1479
tmp11054 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1479)


if True == tmp11054 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1479)
return
}


}, 1)

tmp11076 := Call(__e, PrimFunc(symshen_4ccons_2), V1471)


var ifres11055 Obj

if True == tmp11076 {
tmp11056 := MakeNative(func(__e *ControlFlow) {
W1480 := __e.Get(1)
_ = W1480
tmp11057 := MakeNative(func(__e *ControlFlow) {
W1481 := __e.Get(1)
_ = W1481
tmp11058 := MakeNative(func(__e *ControlFlow) {
W1482 := __e.Get(1)
_ = W1482
tmp11070 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1482)


if True == tmp11070 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11059 := MakeNative(func(__e *ControlFlow) {
W1483 := __e.Get(1)
_ = W1483
tmp11060 := MakeNative(func(__e *ControlFlow) {
W1484 := __e.Get(1)
_ = W1484
tmp11061 := MakeNative(func(__e *ControlFlow) {
W1485 := __e.Get(1)
_ = W1485
tmp11065 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1485)


if True == tmp11065 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11062 := MakeNative(func(__e *ControlFlow) {
W1486 := __e.Get(1)
_ = W1486
__e.TailApply(PrimFunc(symshen_4comb), W1481, W1483)
return
}, 1)

tmp11063 := Call(__e, PrimFunc(symshen_4in_1_6), W1485)


__e.TailApply(tmp11062, tmp11063)
return


}


}, 1)

tmp11066 := Call(__e, PrimFunc(sym_5end_6), W1484)


__e.TailApply(tmp11061, tmp11066)
return


}, 1)

tmp11067 := Call(__e, PrimFunc(symshen_4in_1_6), W1482)


__e.TailApply(tmp11060, tmp11067)
return


}, 1)

tmp11068 := Call(__e, PrimFunc(symshen_4_5_1out), W1482)


__e.TailApply(tmp11059, tmp11068)
return


}


}, 1)

tmp11071 := Call(__e, PrimFunc(symshen_4_5bterms_6), W1480)


__e.TailApply(tmp11058, tmp11071)
return


}, 1)

tmp11072 := Call(__e, PrimFunc(symtail), V1471)


__e.TailApply(tmp11057, tmp11072)
return


}, 1)

tmp11073 := Call(__e, PrimFunc(symhead), V1471)


tmp11074 := Call(__e, tmp11056, tmp11073)


ifres11055 = tmp11074


} else {
tmp11075 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11055 = tmp11075


}

__e.TailApply(tmp11052, ifres11055)
return


} else {
__e.Return(W1476)
return
}


}, 1)

tmp11088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1471)
}
__typedArg0 := V1471
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11079 Obj

if True == tmp11088 {
tmp11080 := MakeNative(func(__e *ControlFlow) {
W1477 := __e.Get(1)
_ = W1477
tmp11081 := MakeNative(func(__e *ControlFlow) {
W1478 := __e.Get(1)
_ = W1478
tmp11083 := Call(__e, PrimFunc(symatom_2), W1477)


if True == tmp11083 {
__e.TailApply(PrimFunc(symshen_4comb), W1478, W1477)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp11084 := Call(__e, PrimFunc(symtail), V1471)


__e.TailApply(tmp11081, tmp11084)
return


}, 1)

tmp11085 := Call(__e, PrimFunc(symhead), V1471)


tmp11086 := Call(__e, tmp11080, tmp11085)


ifres11079 = tmp11086


} else {
tmp11087 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11079 = tmp11087


}

__e.TailApply(tmp11051, ifres11079)
return


} else {
__e.Return(W1472)
return
}


}, 1)

tmp11091 := MakeNative(func(__e *ControlFlow) {
W1473 := __e.Get(1)
_ = W1473
tmp11097 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1473)


if True == tmp11097 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11092 := MakeNative(func(__e *ControlFlow) {
W1474 := __e.Get(1)
_ = W1474
tmp11093 := MakeNative(func(__e *ControlFlow) {
W1475 := __e.Get(1)
_ = W1475
__e.TailApply(PrimFunc(symshen_4comb), W1475, W1474)
return
}, 1)

tmp11094 := Call(__e, PrimFunc(symshen_4in_1_6), W1473)


__e.TailApply(tmp11093, tmp11094)
return


}, 1)

tmp11095 := Call(__e, PrimFunc(symshen_4_5_1out), W1473)


__e.TailApply(tmp11092, tmp11095)
return


}


}, 1)

tmp11098 := Call(__e, PrimFunc(symshen_4_5wildcard_6), V1471)


tmp11099 := Call(__e, tmp11091, tmp11098)


__e.TailApply(tmp11050, tmp11099)
return


}, 1)

tmp11100 := Call(__e, ns2_1set, symshen_4_5bterm_6, tmp11049)


_ = tmp11100

tmp11101 := MakeNative(func(__e *ControlFlow) {
V1487 := __e.Get(1)
_ = V1487
tmp11102 := MakeNative(func(__e *ControlFlow) {
W1488 := __e.Get(1)
_ = W1488
tmp11104 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1488)


if True == tmp11104 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1488)
return
}


}, 1)

tmp11115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1487)
}
__typedArg0 := V1487
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11105 Obj

if True == tmp11115 {
tmp11106 := MakeNative(func(__e *ControlFlow) {
W1489 := __e.Get(1)
_ = W1489
tmp11107 := MakeNative(func(__e *ControlFlow) {
W1490 := __e.Get(1)
_ = W1490
tmp11110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W1489, sym__)
}
__typedArg0 := W1489
__typedArg1 := sym__
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11110 {
tmp11108 := Call(__e, PrimFunc(symgensym), symY)


__e.TailApply(PrimFunc(symshen_4comb), W1490, tmp11108)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp11111 := Call(__e, PrimFunc(symtail), V1487)


__e.TailApply(tmp11107, tmp11111)
return


}, 1)

tmp11112 := Call(__e, PrimFunc(symhead), V1487)


tmp11113 := Call(__e, tmp11106, tmp11112)


ifres11105 = tmp11113


} else {
tmp11114 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11105 = tmp11114


}

__e.TailApply(tmp11102, ifres11105)
return


}, 1)

tmp11116 := Call(__e, ns2_1set, symshen_4_5wildcard_6, tmp11101)


_ = tmp11116

tmp11117 := MakeNative(func(__e *ControlFlow) {
V1491 := __e.Get(1)
_ = V1491
tmp11118 := MakeNative(func(__e *ControlFlow) {
W1492 := __e.Get(1)
_ = W1492
tmp11120 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W1492)


if True == tmp11120 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W1492)
return
}


}, 1)

tmp11130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1491)
}
__typedArg0 := V1491
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11121 Obj

if True == tmp11130 {
tmp11122 := MakeNative(func(__e *ControlFlow) {
W1493 := __e.Get(1)
_ = W1493
tmp11123 := MakeNative(func(__e *ControlFlow) {
W1494 := __e.Get(1)
_ = W1494
tmp11125 := Call(__e, PrimFunc(symshen_4semicolon_2), W1493)


if True == tmp11125 {
__e.TailApply(PrimFunc(symshen_4comb), W1494, W1493)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp11126 := Call(__e, PrimFunc(symtail), V1491)


__e.TailApply(tmp11123, tmp11126)
return


}, 1)

tmp11127 := Call(__e, PrimFunc(symhead), V1491)


tmp11128 := Call(__e, tmp11122, tmp11127)


ifres11121 = tmp11128


} else {
tmp11129 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11121 = tmp11129


}

__e.TailApply(tmp11118, ifres11121)
return


}, 1)

tmp11131 := Call(__e, ns2_1set, symshen_4_5sc_6, tmp11117)


_ = tmp11131

tmp11132 := MakeNative(func(__e *ControlFlow) {
V1495 := __e.Get(1)
_ = V1495
V1496 := __e.Get(2)
_ = V1496
tmp11133 := MakeNative(func(__e *ControlFlow) {
W1497 := __e.Get(1)
_ = W1497
tmp11134 := MakeNative(func(__e *ControlFlow) {
W1498 := __e.Get(1)
_ = W1498
tmp11135 := MakeNative(func(__e *ControlFlow) {
W1499 := __e.Get(1)
_ = W1499
tmp11136 := MakeNative(func(__e *ControlFlow) {
W1500 := __e.Get(1)
_ = W1500
tmp11137 := MakeNative(func(__e *ControlFlow) {
W1501 := __e.Get(1)
_ = W1501
tmp11138 := MakeNative(func(__e *ControlFlow) {
W1502 := __e.Get(1)
_ = W1502
tmp11139 := MakeNative(func(__e *ControlFlow) {
W1503 := __e.Get(1)
_ = W1503
tmp11140 := MakeNative(func(__e *ControlFlow) {
W1504 := __e.Get(1)
_ = W1504
tmp11141 := MakeNative(func(__e *ControlFlow) {
W1505 := __e.Get(1)
_ = W1505
__e.Return(W1505)
return
}, 1)

tmp11142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, Nil)
}
__typedArg0 := sym_1_6
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1500, tmp11142)
}
__typedArg0 := W1500
__typedArg1 := tmp11142
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1499, tmp11143)
}
__typedArg0 := W1499
__typedArg1 := tmp11143
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1498, tmp11144)
}
__typedArg0 := W1498
__typedArg1 := tmp11144
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1497, tmp11145)
}
__typedArg0 := W1497
__typedArg1 := tmp11145
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1504, Nil)
}
__typedArg0 := W1504
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11148 := Call(__e, PrimFunc(symappend), tmp11146, tmp11147)


tmp11149 := Call(__e, PrimFunc(symappend), W1501, tmp11148)


tmp11150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1495, tmp11149)
}
__typedArg0 := V1495
__typedArg1 := tmp11149
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp11150)
}
__typedArg0 := symdefine
__typedArg1 := tmp11150
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp11141, tmp11151)
return


}, 1)

var ifres11152 Obj

if True == W1502 {
tmp11153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(1), Nil)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1499, tmp11153)
}
__typedArg0 := W1499
__typedArg1 := tmp11153
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp11154)
}
__typedArg0 := sym_7
__typedArg1 := tmp11154
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1503, Nil)
}
__typedArg0 := W1503
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11155, tmp11156)
}
__typedArg0 := tmp11155
__typedArg1 := tmp11156
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1499, tmp11157)
}
__typedArg0 := W1499
__typedArg1 := tmp11157
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11158)
}
__typedArg0 := symlet
__typedArg1 := tmp11158
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

ifres11152 = tmp11159


} else {
ifres11152 = W1503


}

__e.TailApply(tmp11140, ifres11152)
return


}, 1)

tmp11160 := Call(__e, PrimFunc(symshen_4prolog_1fbody), V1496, W1501, W1497, W1498, W1499, W1500, W1502)


__e.TailApply(tmp11139, tmp11160)
return


}, 1)

tmp11161 := Call(__e, PrimFunc(symshen_4hascut_2), V1496)


__e.TailApply(tmp11138, tmp11161)
return


}, 1)

tmp11162 := Call(__e, PrimFunc(symshen_4prolog_1parameters), V1496)


__e.TailApply(tmp11137, tmp11162)
return


}, 1)

tmp11163 := Call(__e, PrimFunc(symgensym), symC)


__e.TailApply(tmp11136, tmp11163)
return


}, 1)

tmp11164 := Call(__e, PrimFunc(symgensym), symK)


__e.TailApply(tmp11135, tmp11164)
return


}, 1)

tmp11165 := Call(__e, PrimFunc(symgensym), symL)


__e.TailApply(tmp11134, tmp11165)
return


}, 1)

tmp11166 := Call(__e, PrimFunc(symgensym), symB)


__e.TailApply(tmp11133, tmp11166)
return


}, 2)

tmp11167 := Call(__e, ns2_1set, symshen_4horn_1clause_1procedure, tmp11132)


_ = tmp11167

tmp11168 := MakeNative(func(__e *ControlFlow) {
V1508 := __e.Get(1)
_ = V1508
tmp11178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_b, V1508)
}
__typedArg0 := sym_b
__typedArg1 := V1508
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11178 {
__e.Return(True)
return
} else {
tmp11176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1508)
}
__typedArg0 := V1508
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11176 {
tmp11173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1508)
}
__typedArg0 := V1508
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11174 := Call(__e, PrimFunc(symshen_4hascut_2), tmp11173)


if True == tmp11174 {
__e.Return(True)
return
} else {
tmp11170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1508)
}
__typedArg0 := V1508
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11171 := Call(__e, PrimFunc(symshen_4hascut_2), tmp11170)


if True == tmp11171 {
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

tmp11179 := Call(__e, ns2_1set, symshen_4hascut_2, tmp11168)


_ = tmp11179

tmp11180 := MakeNative(func(__e *ControlFlow) {
V1513 := __e.Get(1)
_ = V1513
tmp11189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1513)
}
__typedArg0 := V1513
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11185 Obj

if True == tmp11189 {
tmp11187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1513)
}
__typedArg0 := V1513
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11187)
}
__typedArg0 := tmp11187
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11186 Obj

if True == tmp11188 {
ifres11186 = True


} else {
ifres11186 = False


}

ifres11185 = ifres11186


} else {
ifres11185 = False


}

if True == ifres11185 {
tmp11181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1513)
}
__typedArg0 := V1513
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11181)
}
__typedArg0 := tmp11181
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11183 := Call(__e, PrimFunc(symlength), tmp11182)


__e.TailApply(PrimFunc(symshen_4parameters), tmp11183)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.prolog-parameters"))
}
__typedArg0 := MakeString("partial function shen.prolog-parameters")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp11190 := Call(__e, ns2_1set, symshen_4prolog_1parameters, tmp11180)


_ = tmp11190

tmp11191 := MakeNative(func(__e *ControlFlow) {
V1534 := __e.Get(1)
_ = V1534
V1535 := __e.Get(2)
_ = V1535
V1536 := __e.Get(3)
_ = V1536
V1537 := __e.Get(4)
_ = V1537
V1538 := __e.Get(5)
_ = V1538
V1539 := __e.Get(6)
_ = V1539
V1540 := __e.Get(7)
_ = V1540
tmp11284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1534)
}
__typedArg0 := Nil
__typedArg1 := V1534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11281 Obj

if True == tmp11284 {
tmp11283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, V1540)
}
__typedArg0 := True
__typedArg1 := V1540
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11282 Obj

if True == tmp11283 {
ifres11282 = True


} else {
ifres11282 = False


}

ifres11281 = ifres11282


} else {
ifres11281 = False


}

if True == ifres11281 {
tmp11192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1538, Nil)
}
__typedArg0 := V1538
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1537, tmp11192)
}
__typedArg0 := V1537
__typedArg1 := tmp11192
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4unlock, tmp11193)
}
__typedArg0 := symshen_4unlock
__typedArg1 := tmp11193
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11257 Obj

if True == tmp11279 {
tmp11277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11277)
}
__typedArg0 := tmp11277
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11259 Obj

if True == tmp11278 {
tmp11274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11274)
}
__typedArg0 := tmp11274
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11275)
}
__typedArg0 := tmp11275
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11261 Obj

if True == tmp11276 {
tmp11270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11270)
}
__typedArg0 := tmp11270
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11271)
}
__typedArg0 := tmp11271
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11272)
}
__typedArg0 := Nil
__typedArg1 := tmp11272
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11263 Obj

if True == tmp11273 {
tmp11268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11268)
}
__typedArg0 := Nil
__typedArg1 := tmp11268
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11265 Obj

if True == tmp11269 {
tmp11267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, V1540)
}
__typedArg0 := False
__typedArg1 := V1540
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11266 Obj

if True == tmp11267 {
ifres11266 = True


} else {
ifres11266 = False


}

ifres11265 = ifres11266


} else {
ifres11265 = False


}

var ifres11264 Obj

if True == ifres11265 {
ifres11264 = True


} else {
ifres11264 = False


}

ifres11263 = ifres11264


} else {
ifres11263 = False


}

var ifres11262 Obj

if True == ifres11263 {
ifres11262 = True


} else {
ifres11262 = False


}

ifres11261 = ifres11262


} else {
ifres11261 = False


}

var ifres11260 Obj

if True == ifres11261 {
ifres11260 = True


} else {
ifres11260 = False


}

ifres11259 = ifres11260


} else {
ifres11259 = False


}

var ifres11258 Obj

if True == ifres11259 {
ifres11258 = True


} else {
ifres11258 = False


}

ifres11257 = ifres11258


} else {
ifres11257 = False


}

if True == ifres11257 {
tmp11194 := MakeNative(func(__e *ControlFlow) {
W1541 := __e.Get(1)
_ = W1541
tmp11195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1537, Nil)
}
__typedArg0 := V1537
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4unlocked_2, tmp11195)
}
__typedArg0 := symshen_4unlocked_2
__typedArg1 := tmp11195
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11197)
}
__typedArg0 := tmp11197
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11199 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11198, V1535, V1536, W1541)


tmp11200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11199, tmp11200)
}
__typedArg0 := tmp11199
__typedArg1 := tmp11200
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11196, tmp11201)
}
__typedArg0 := tmp11196
__typedArg1 := tmp11201
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11202)
}
__typedArg0 := symif
__typedArg1 := tmp11202
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp11203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11203)
}
__typedArg0 := tmp11203
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11205)
}
__typedArg0 := tmp11205
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11206)
}
__typedArg0 := tmp11206
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11208 := Call(__e, PrimFunc(symshen_4continue), tmp11204, tmp11207, V1536, V1537, V1538, V1539)


__e.TailApply(tmp11194, tmp11208)
return


} else {
tmp11255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11240 Obj

if True == tmp11255 {
tmp11253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11253)
}
__typedArg0 := tmp11253
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11242 Obj

if True == tmp11254 {
tmp11250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11250)
}
__typedArg0 := tmp11250
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11251)
}
__typedArg0 := tmp11251
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11244 Obj

if True == tmp11252 {
tmp11246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11246)
}
__typedArg0 := tmp11246
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11247)
}
__typedArg0 := tmp11247
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11248)
}
__typedArg0 := Nil
__typedArg1 := tmp11248
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11245 Obj

if True == tmp11249 {
ifres11245 = True


} else {
ifres11245 = False


}

ifres11244 = ifres11245


} else {
ifres11244 = False


}

var ifres11243 Obj

if True == ifres11244 {
ifres11243 = True


} else {
ifres11243 = False


}

ifres11242 = ifres11243


} else {
ifres11242 = False


}

var ifres11241 Obj

if True == ifres11242 {
ifres11241 = True


} else {
ifres11241 = False


}

ifres11240 = ifres11241


} else {
ifres11240 = False


}

if True == ifres11240 {
tmp11209 := MakeNative(func(__e *ControlFlow) {
W1542 := __e.Get(1)
_ = W1542
tmp11210 := MakeNative(func(__e *ControlFlow) {
W1543 := __e.Get(1)
_ = W1543
tmp11211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1537, Nil)
}
__typedArg0 := V1537
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4unlocked_2, tmp11211)
}
__typedArg0 := symshen_4unlocked_2
__typedArg1 := tmp11211
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11213)
}
__typedArg0 := tmp11213
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11215 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11214, V1535, V1536, W1543)


tmp11216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11215, tmp11216)
}
__typedArg0 := tmp11215
__typedArg1 := tmp11216
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11212, tmp11217)
}
__typedArg0 := tmp11212
__typedArg1 := tmp11217
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11218)
}
__typedArg0 := symif
__typedArg1 := tmp11218
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1542, tmp11220)
}
__typedArg0 := W1542
__typedArg1 := tmp11220
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp11221)
}
__typedArg0 := sym_a
__typedArg1 := tmp11221
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11224 := Call(__e, PrimFunc(symshen_4prolog_1fbody), tmp11223, V1535, V1536, V1537, V1538, V1539, V1540)


tmp11225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1542, Nil)
}
__typedArg0 := W1542
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11224, tmp11225)
}
__typedArg0 := tmp11224
__typedArg1 := tmp11225
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11222, tmp11226)
}
__typedArg0 := tmp11222
__typedArg1 := tmp11226
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11227)
}
__typedArg0 := symif
__typedArg1 := tmp11227
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11228, Nil)
}
__typedArg0 := tmp11228
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11219, tmp11229)
}
__typedArg0 := tmp11219
__typedArg1 := tmp11229
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1542, tmp11230)
}
__typedArg0 := W1542
__typedArg1 := tmp11230
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11231)
}
__typedArg0 := symlet
__typedArg1 := tmp11231
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp11232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11232)
}
__typedArg0 := tmp11232
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11234)
}
__typedArg0 := tmp11234
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11235)
}
__typedArg0 := tmp11235
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11237 := Call(__e, PrimFunc(symshen_4continue), tmp11233, tmp11236, V1536, V1537, V1538, V1539)


__e.TailApply(tmp11210, tmp11237)
return


}, 1)

tmp11238 := Call(__e, PrimFunc(symgensym), symC)


__e.TailApply(tmp11209, tmp11238)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.prolog-fbody"))
}
__typedArg0 := MakeString("implementation error in shen.prolog-fbody")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 7)

tmp11285 := Call(__e, ns2_1set, symshen_4prolog_1fbody, tmp11191)


_ = tmp11285

tmp11286 := MakeNative(func(__e *ControlFlow) {
V1544 := __e.Get(1)
_ = V1544
V1545 := __e.Get(2)
_ = V1545
tmp11291 := Call(__e, PrimFunc(symshen_4locked_2), V1544)


var ifres11288 Obj

if True == tmp11291 {
tmp11290 := Call(__e, PrimFunc(symshen_4fits_2), V1545, V1544)


var ifres11289 Obj

if True == tmp11290 {
ifres11289 = True


} else {
ifres11289 = False


}

ifres11288 = ifres11289


} else {
ifres11288 = False


}

if True == ifres11288 {
__e.TailApply(PrimFunc(symshen_4openlock), V1544)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp11292 := Call(__e, ns2_1set, symshen_4unlock, tmp11286)


_ = tmp11292

tmp11293 := MakeNative(func(__e *ControlFlow) {
V1546 := __e.Get(1)
_ = V1546
tmp11294 := Call(__e, PrimFunc(symshen_4unlocked_2), V1546)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp11294)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp11294
return Call(__e, PrimFunc(symnot), __typedArg0)
})())
return


}, 1)

tmp11295 := Call(__e, ns2_1set, symshen_4locked_2, tmp11293)


_ = tmp11295

tmp11296 := MakeNative(func(__e *ControlFlow) {
V1547 := __e.Get(1)
_ = V1547
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1547, MakeNumber(1))
}
__typedArg0 := V1547
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp11297 := Call(__e, ns2_1set, symshen_4unlocked_2, tmp11296)


_ = tmp11297

tmp11298 := MakeNative(func(__e *ControlFlow) {
V1548 := __e.Get(1)
_ = V1548
tmp11299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V1548, MakeNumber(1), True)
}
__typedArg0 := V1548
__typedArg1 := MakeNumber(1)
__typedArg2 := True
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

_ = tmp11299

__e.Return(False)
return


}, 1)

tmp11300 := Call(__e, ns2_1set, symshen_4openlock, tmp11298)


_ = tmp11300

tmp11301 := MakeNative(func(__e *ControlFlow) {
V1549 := __e.Get(1)
_ = V1549
V1550 := __e.Get(2)
_ = V1550
tmp11302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1550, MakeNumber(2))
}
__typedArg0 := V1550
__typedArg1 := MakeNumber(2)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1549, tmp11302)
}
__typedArg0 := V1549
__typedArg1 := tmp11302
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp11303 := Call(__e, ns2_1set, symshen_4fits_2, tmp11301)


_ = tmp11303

tmp11304 := MakeNative(func(__e *ControlFlow) {
V1553 := __e.Get(1)
_ = V1553
V1554 := __e.Get(2)
_ = V1554
V1555 := __e.Get(3)
_ = V1555
V1556 := __e.Get(4)
_ = V1556
tmp11305 := MakeNative(func(__e *ControlFlow) {
W1557 := __e.Get(1)
_ = W1557
tmp11310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W1557, False)
}
__typedArg0 := W1557
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11307 Obj

if True == tmp11310 {
tmp11309 := Call(__e, PrimFunc(symshen_4unlocked_2), V1554)


var ifres11308 Obj

if True == tmp11309 {
ifres11308 = True


} else {
ifres11308 = False


}

ifres11307 = ifres11308


} else {
ifres11307 = False


}

if True == ifres11307 {
__e.TailApply(PrimFunc(symshen_4lock), V1555, V1554)
return
} else {
__e.Return(W1557)
return
}


}, 1)

tmp11311 := Call(__e, PrimFunc(symthaw), V1556)


__e.TailApply(tmp11305, tmp11311)
return


}, 4)

tmp11312 := Call(__e, ns2_1set, symshen_4cut, tmp11304)


_ = tmp11312

tmp11313 := MakeNative(func(__e *ControlFlow) {
V1558 := __e.Get(1)
_ = V1558
V1559 := __e.Get(2)
_ = V1559
tmp11314 := MakeNative(func(__e *ControlFlow) {
W1560 := __e.Get(1)
_ = W1560
tmp11315 := MakeNative(func(__e *ControlFlow) {
W1561 := __e.Get(1)
_ = W1561
__e.Return(False)
return
}, 1)

tmp11316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V1559, MakeNumber(2), V1558)
}
__typedArg0 := V1559
__typedArg1 := MakeNumber(2)
__typedArg2 := V1558
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

__e.TailApply(tmp11315, tmp11316)
return


}, 1)

tmp11317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V1559, MakeNumber(1), False)
}
__typedArg0 := V1559
__typedArg1 := MakeNumber(1)
__typedArg2 := False
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

__e.TailApply(tmp11314, tmp11317)
return


}, 2)

tmp11318 := Call(__e, ns2_1set, symshen_4lock, tmp11313)


_ = tmp11318

tmp11319 := MakeNative(func(__e *ControlFlow) {
V1562 := __e.Get(1)
_ = V1562
V1563 := __e.Get(2)
_ = V1563
V1564 := __e.Get(3)
_ = V1564
V1565 := __e.Get(4)
_ = V1565
V1566 := __e.Get(5)
_ = V1566
V1567 := __e.Get(6)
_ = V1567
tmp11320 := MakeNative(func(__e *ControlFlow) {
W1568 := __e.Get(1)
_ = W1568
tmp11321 := MakeNative(func(__e *ControlFlow) {
W1569 := __e.Get(1)
_ = W1569
tmp11322 := MakeNative(func(__e *ControlFlow) {
W1570 := __e.Get(1)
_ = W1570
tmp11323 := MakeNative(func(__e *ControlFlow) {
W1571 := __e.Get(1)
_ = W1571
__e.TailApply(PrimFunc(symshen_4stpart), W1570, W1571, V1564)
return
}, 1)

tmp11324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4incinfs, Nil)
}
__typedArg0 := symshen_4incinfs
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11325 := Call(__e, PrimFunc(symshen_4compile_1body), V1563, V1564, V1565, V1566, V1567)


tmp11326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11325, Nil)
}
__typedArg0 := tmp11325
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11324, tmp11326)
}
__typedArg0 := tmp11324
__typedArg1 := tmp11326
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp11327)
}
__typedArg0 := symdo
__typedArg1 := tmp11327
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp11323, tmp11328)
return


}, 1)

tmp11329 := Call(__e, PrimFunc(symdifference), W1569, W1568)


__e.TailApply(tmp11322, tmp11329)
return


}, 1)

tmp11330 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), V1563)


__e.TailApply(tmp11321, tmp11330)
return


}, 1)

tmp11331 := Call(__e, PrimFunc(symshen_4extract_1vars), V1562)


__e.TailApply(tmp11320, tmp11331)
return


}, 6)

tmp11332 := Call(__e, ns2_1set, symshen_4continue, tmp11319)


_ = tmp11332

tmp11333 := MakeNative(func(__e *ControlFlow) {
V1574 := __e.Get(1)
_ = V1574
tmp11368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1574)
}
__typedArg0 := V1574
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11349 Obj

if True == tmp11368 {
tmp11366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1574)
}
__typedArg0 := V1574
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp11366)
}
__typedArg0 := symlambda
__typedArg1 := tmp11366
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11351 Obj

if True == tmp11367 {
tmp11364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1574)
}
__typedArg0 := V1574
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11364)
}
__typedArg0 := tmp11364
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11353 Obj

if True == tmp11365 {
tmp11361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1574)
}
__typedArg0 := V1574
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11361)
}
__typedArg0 := tmp11361
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11362)
}
__typedArg0 := tmp11362
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11355 Obj

if True == tmp11363 {
tmp11357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1574)
}
__typedArg0 := V1574
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11357)
}
__typedArg0 := tmp11357
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11358)
}
__typedArg0 := tmp11358
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11359)
}
__typedArg0 := Nil
__typedArg1 := tmp11359
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11356 Obj

if True == tmp11360 {
ifres11356 = True


} else {
ifres11356 = False


}

ifres11355 = ifres11356


} else {
ifres11355 = False


}

var ifres11354 Obj

if True == ifres11355 {
ifres11354 = True


} else {
ifres11354 = False


}

ifres11353 = ifres11354


} else {
ifres11353 = False


}

var ifres11352 Obj

if True == ifres11353 {
ifres11352 = True


} else {
ifres11352 = False


}

ifres11351 = ifres11352


} else {
ifres11351 = False


}

var ifres11350 Obj

if True == ifres11351 {
ifres11350 = True


} else {
ifres11350 = False


}

ifres11349 = ifres11350


} else {
ifres11349 = False


}

if True == ifres11349 {
tmp11334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1574)
}
__typedArg0 := V1574
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11334)
}
__typedArg0 := tmp11334
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1574)
}
__typedArg0 := V1574
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11336)
}
__typedArg0 := tmp11336
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11337)
}
__typedArg0 := tmp11337
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11339 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp11338)


__e.TailApply(PrimFunc(symremove), tmp11335, tmp11339)
return


} else {
tmp11347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1574)
}
__typedArg0 := V1574
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11347 {
tmp11340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1574)
}
__typedArg0 := V1574
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11341 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp11340)


tmp11342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1574)
}
__typedArg0 := V1574
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11343 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp11342)


__e.TailApply(PrimFunc(symunion), tmp11341, tmp11343)
return


} else {
tmp11345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V1574)
}
__typedArg0 := V1574
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp11345 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1574, Nil)
}
__typedArg0 := V1574
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
} else {
__e.Return(Nil)
return
}


}


}


}, 1)

tmp11369 := Call(__e, ns2_1set, symshen_4extract_1free_1vars, tmp11333)


_ = tmp11369

tmp11370 := MakeNative(func(__e *ControlFlow) {
V1591 := __e.Get(1)
_ = V1591
V1592 := __e.Get(2)
_ = V1592
V1593 := __e.Get(3)
_ = V1593
V1594 := __e.Get(4)
_ = V1594
V1595 := __e.Get(5)
_ = V1595
tmp11405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1591)
}
__typedArg0 := Nil
__typedArg1 := V1591
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11405 {
tmp11371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1595, Nil)
}
__typedArg0 := V1595
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp11371)
}
__typedArg0 := symthaw
__typedArg1 := tmp11371
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1591)
}
__typedArg0 := V1591
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11399 Obj

if True == tmp11403 {
tmp11401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1591)
}
__typedArg0 := V1591
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_b, tmp11401)
}
__typedArg0 := sym_b
__typedArg1 := tmp11401
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11400 Obj

if True == tmp11402 {
ifres11400 = True


} else {
ifres11400 = False


}

ifres11399 = ifres11400


} else {
ifres11399 = False


}

if True == ifres11399 {
tmp11372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4cut, Nil)
}
__typedArg0 := symshen_4cut
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1591)
}
__typedArg0 := V1591
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11372, tmp11373)
}
__typedArg0 := tmp11372
__typedArg1 := tmp11373
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4compile_1body), tmp11374, V1592, V1593, V1594, V1595)
return


} else {
tmp11397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1591)
}
__typedArg0 := V1591
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11393 Obj

if True == tmp11397 {
tmp11395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1591)
}
__typedArg0 := V1591
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11395)
}
__typedArg0 := Nil
__typedArg1 := tmp11395
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11394 Obj

if True == tmp11396 {
ifres11394 = True


} else {
ifres11394 = False


}

ifres11393 = ifres11394


} else {
ifres11393 = False


}

if True == ifres11393 {
tmp11375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1591)
}
__typedArg0 := V1591
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11376 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11375, V1592)


tmp11377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1595, Nil)
}
__typedArg0 := V1595
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1594, tmp11377)
}
__typedArg0 := V1594
__typedArg1 := tmp11377
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1593, tmp11378)
}
__typedArg0 := V1593
__typedArg1 := tmp11378
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1592, tmp11379)
}
__typedArg0 := V1592
__typedArg1 := tmp11379
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symappend), tmp11376, tmp11380)
return


} else {
tmp11391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1591)
}
__typedArg0 := V1591
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11391 {
tmp11381 := MakeNative(func(__e *ControlFlow) {
W1596 := __e.Get(1)
_ = W1596
tmp11382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1591)
}
__typedArg0 := V1591
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11383 := Call(__e, PrimFunc(symshen_4freeze_1literals), tmp11382, V1592, V1593, V1594, V1595)


tmp11384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11383, Nil)
}
__typedArg0 := tmp11383
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1594, tmp11384)
}
__typedArg0 := V1594
__typedArg1 := tmp11384
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1593, tmp11385)
}
__typedArg0 := V1593
__typedArg1 := tmp11385
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1592, tmp11386)
}
__typedArg0 := V1592
__typedArg1 := tmp11386
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symappend), W1596, tmp11387)
return


}, 1)

tmp11388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1591)
}
__typedArg0 := V1591
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11389 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11388, V1592)


__e.TailApply(tmp11381, tmp11389)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.compile-fbody"))
}
__typedArg0 := MakeString("implementation error in shen.compile-fbody")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}


}, 5)

tmp11406 := Call(__e, ns2_1set, symshen_4compile_1body, tmp11370)


_ = tmp11406

tmp11407 := MakeNative(func(__e *ControlFlow) {
V1613 := __e.Get(1)
_ = V1613
V1614 := __e.Get(2)
_ = V1614
V1615 := __e.Get(3)
_ = V1615
V1616 := __e.Get(4)
_ = V1616
V1617 := __e.Get(5)
_ = V1617
tmp11431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1613)
}
__typedArg0 := Nil
__typedArg1 := V1613
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11431 {
__e.Return(V1617)
return
} else {
tmp11429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1613)
}
__typedArg0 := V1613
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11425 Obj

if True == tmp11429 {
tmp11427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1613)
}
__typedArg0 := V1613
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_b, tmp11427)
}
__typedArg0 := sym_b
__typedArg1 := tmp11427
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11426 Obj

if True == tmp11428 {
ifres11426 = True


} else {
ifres11426 = False


}

ifres11425 = ifres11426


} else {
ifres11425 = False


}

if True == ifres11425 {
tmp11408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4cut, Nil)
}
__typedArg0 := symshen_4cut
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1613)
}
__typedArg0 := V1613
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11408, tmp11409)
}
__typedArg0 := tmp11408
__typedArg1 := tmp11409
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4freeze_1literals), tmp11410, V1614, V1615, V1616, V1617)
return


} else {
tmp11423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1613)
}
__typedArg0 := V1613
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11423 {
tmp11411 := MakeNative(func(__e *ControlFlow) {
W1618 := __e.Get(1)
_ = W1618
tmp11412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1613)
}
__typedArg0 := V1613
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11413 := Call(__e, PrimFunc(symshen_4freeze_1literals), tmp11412, V1614, V1615, V1616, V1617)


tmp11414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11413, Nil)
}
__typedArg0 := tmp11413
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1616, tmp11414)
}
__typedArg0 := V1616
__typedArg1 := tmp11414
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1615, tmp11415)
}
__typedArg0 := V1615
__typedArg1 := tmp11415
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1614, tmp11416)
}
__typedArg0 := V1614
__typedArg1 := tmp11416
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11418 := Call(__e, PrimFunc(symappend), W1618, tmp11417)


tmp11419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11418, Nil)
}
__typedArg0 := tmp11418
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp11419)
}
__typedArg0 := symfreeze
__typedArg1 := tmp11419
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp11420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1613)
}
__typedArg0 := V1613
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11421 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11420, V1614)


__e.TailApply(tmp11411, tmp11421)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.freeze-literals"))
}
__typedArg0 := MakeString("implementation error in shen.freeze-literals")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 5)

tmp11432 := Call(__e, ns2_1set, symshen_4freeze_1literals, tmp11407)


_ = tmp11432

tmp11433 := MakeNative(func(__e *ControlFlow) {
V1623 := __e.Get(1)
_ = V1623
V1624 := __e.Get(2)
_ = V1624
tmp11448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1623)
}
__typedArg0 := V1623
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11444 Obj

if True == tmp11448 {
tmp11446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1623)
}
__typedArg0 := V1623
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfork, tmp11446)
}
__typedArg0 := symfork
__typedArg1 := tmp11446
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11445 Obj

if True == tmp11447 {
ifres11445 = True


} else {
ifres11445 = False


}

ifres11444 = ifres11445


} else {
ifres11444 = False


}

if True == ifres11444 {
tmp11434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1623)
}
__typedArg0 := V1623
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11435 := Call(__e, PrimFunc(symshen_4deref_1forked_1literals), tmp11434, V1624)


tmp11436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11435, Nil)
}
__typedArg0 := tmp11435
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfork, tmp11436)
}
__typedArg0 := symfork
__typedArg1 := tmp11436
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1623)
}
__typedArg0 := V1623
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11442 {
tmp11437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1623)
}
__typedArg0 := V1623
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11438 := MakeNative(func(__e *ControlFlow) {
Z1625 := __e.Get(1)
_ = Z1625
__e.TailApply(PrimFunc(symshen_4function_1calls), Z1625, V1624)
return
}, 1)

tmp11439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1623)
}
__typedArg0 := V1623
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11440 := Call(__e, PrimFunc(symmap), tmp11438, tmp11439)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11437, tmp11440)
}
__typedArg0 := tmp11437
__typedArg1 := tmp11440
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.deref-calls"))
}
__typedArg0 := MakeString("implementation error in shen.deref-calls")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp11449 := Call(__e, ns2_1set, symshen_4deref_1calls, tmp11433)


_ = tmp11449

tmp11450 := MakeNative(func(__e *ControlFlow) {
V1632 := __e.Get(1)
_ = V1632
V1633 := __e.Get(2)
_ = V1633
tmp11460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1632)
}
__typedArg0 := Nil
__typedArg1 := V1632
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11460 {
__e.Return(Nil)
return
} else {
tmp11458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1632)
}
__typedArg0 := V1632
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11458 {
tmp11451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1632)
}
__typedArg0 := V1632
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11452 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11451, V1633)


tmp11453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1632)
}
__typedArg0 := V1632
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11454 := Call(__e, PrimFunc(symshen_4deref_1forked_1literals), tmp11453, V1633)


tmp11455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11454, Nil)
}
__typedArg0 := tmp11454
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11452, tmp11455)
}
__typedArg0 := tmp11452
__typedArg1 := tmp11455
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp11456)
}
__typedArg0 := symcons
__typedArg1 := tmp11456
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("fork requires a list of literals\n"))
}
__typedArg0 := MakeString("fork requires a list of literals\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp11461 := Call(__e, ns2_1set, symshen_4deref_1forked_1literals, tmp11450)


_ = tmp11461

tmp11462 := MakeNative(func(__e *ControlFlow) {
V1636 := __e.Get(1)
_ = V1636
V1637 := __e.Get(2)
_ = V1637
tmp11494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1636)
}
__typedArg0 := V1636
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11475 Obj

if True == tmp11494 {
tmp11492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1636)
}
__typedArg0 := V1636
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp11492)
}
__typedArg0 := symcons
__typedArg1 := tmp11492
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11477 Obj

if True == tmp11493 {
tmp11490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1636)
}
__typedArg0 := V1636
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11490)
}
__typedArg0 := tmp11490
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11479 Obj

if True == tmp11491 {
tmp11487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1636)
}
__typedArg0 := V1636
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11487)
}
__typedArg0 := tmp11487
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11488)
}
__typedArg0 := tmp11488
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11481 Obj

if True == tmp11489 {
tmp11483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1636)
}
__typedArg0 := V1636
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11483)
}
__typedArg0 := tmp11483
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11484)
}
__typedArg0 := tmp11484
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11485)
}
__typedArg0 := Nil
__typedArg1 := tmp11485
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11482 Obj

if True == tmp11486 {
ifres11482 = True


} else {
ifres11482 = False


}

ifres11481 = ifres11482


} else {
ifres11481 = False


}

var ifres11480 Obj

if True == ifres11481 {
ifres11480 = True


} else {
ifres11480 = False


}

ifres11479 = ifres11480


} else {
ifres11479 = False


}

var ifres11478 Obj

if True == ifres11479 {
ifres11478 = True


} else {
ifres11478 = False


}

ifres11477 = ifres11478


} else {
ifres11477 = False


}

var ifres11476 Obj

if True == ifres11477 {
ifres11476 = True


} else {
ifres11476 = False


}

ifres11475 = ifres11476


} else {
ifres11475 = False


}

if True == ifres11475 {
tmp11463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1636)
}
__typedArg0 := V1636
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11463)
}
__typedArg0 := tmp11463
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11465 := Call(__e, PrimFunc(symshen_4function_1calls), tmp11464, V1637)


tmp11466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1636)
}
__typedArg0 := V1636
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11466)
}
__typedArg0 := tmp11466
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11467)
}
__typedArg0 := tmp11467
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11469 := Call(__e, PrimFunc(symshen_4function_1calls), tmp11468, V1637)


tmp11470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11469, Nil)
}
__typedArg0 := tmp11469
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11465, tmp11470)
}
__typedArg0 := tmp11465
__typedArg1 := tmp11470
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp11471)
}
__typedArg0 := symcons
__typedArg1 := tmp11471
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1636)
}
__typedArg0 := V1636
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11473 {
__e.TailApply(PrimFunc(symshen_4deref_1terms), V1636, V1637, Nil)
return
} else {
__e.Return(V1636)
return
}


}


}, 2)

tmp11495 := Call(__e, ns2_1set, symshen_4function_1calls, tmp11462)


_ = tmp11495

tmp11496 := MakeNative(func(__e *ControlFlow) {
V1646 := __e.Get(1)
_ = V1646
V1647 := __e.Get(2)
_ = V1647
V1648 := __e.Get(3)
_ = V1648
tmp11590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11577 Obj

if True == tmp11590 {
tmp11588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), tmp11588)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := tmp11588
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11579 Obj

if True == tmp11589 {
tmp11586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11586)
}
__typedArg0 := tmp11586
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11581 Obj

if True == tmp11587 {
tmp11583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11583)
}
__typedArg0 := tmp11583
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11584)
}
__typedArg0 := Nil
__typedArg1 := tmp11584
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11582 Obj

if True == tmp11585 {
ifres11582 = True


} else {
ifres11582 = False


}

ifres11581 = ifres11582


} else {
ifres11581 = False


}

var ifres11580 Obj

if True == ifres11581 {
ifres11580 = True


} else {
ifres11580 = False


}

ifres11579 = ifres11580


} else {
ifres11579 = False


}

var ifres11578 Obj

if True == ifres11579 {
ifres11578 = True


} else {
ifres11578 = False


}

ifres11577 = ifres11578


} else {
ifres11577 = False


}

if True == ifres11577 {
tmp11503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11503)
}
__typedArg0 := tmp11503
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp11504)
}
__typedArg0 := tmp11504
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp11505 {
tmp11497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11497)
}
__typedArg0 := tmp11497
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp11498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11498)
}
__typedArg0 := tmp11498
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11500 := Call(__e, PrimFunc(symshen_4app), tmp11499, MakeString("\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attempt to optimise a non-variable "))
__typedS1, __typedOK1 := TypedString(tmp11500)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attempt to optimise a non-variable ")
__typedArg1 := tmp11500
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attempt to optimise a non-variable "))
__typedS1, __typedOK1 := TypedString(tmp11500)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attempt to optimise a non-variable ")
__typedArg1 := tmp11500
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


} else {
tmp11575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11562 Obj

if True == tmp11575 {
tmp11573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(1), tmp11573)
}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp11573
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11564 Obj

if True == tmp11574 {
tmp11571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11571)
}
__typedArg0 := tmp11571
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11566 Obj

if True == tmp11572 {
tmp11568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11568)
}
__typedArg0 := tmp11568
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11569)
}
__typedArg0 := Nil
__typedArg1 := tmp11569
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11567 Obj

if True == tmp11570 {
ifres11567 = True


} else {
ifres11567 = False


}

ifres11566 = ifres11567


} else {
ifres11566 = False


}

var ifres11565 Obj

if True == ifres11566 {
ifres11565 = True


} else {
ifres11565 = False


}

ifres11564 = ifres11565


} else {
ifres11564 = False


}

var ifres11563 Obj

if True == ifres11564 {
ifres11563 = True


} else {
ifres11563 = False


}

ifres11562 = ifres11563


} else {
ifres11562 = False


}

if True == ifres11562 {
tmp11515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11515)
}
__typedArg0 := tmp11515
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp11516)
}
__typedArg0 := tmp11516
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp11517 {
tmp11506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11506)
}
__typedArg0 := tmp11506
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1647, Nil)
}
__typedArg0 := V1647
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11507, tmp11508)
}
__typedArg0 := tmp11507
__typedArg1 := tmp11508
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp11509)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp11509
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11510)
}
__typedArg0 := tmp11510
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11512 := Call(__e, PrimFunc(symshen_4app), tmp11511, MakeString("\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attempt to optimise a non-variable "))
__typedS1, __typedOK1 := TypedString(tmp11512)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attempt to optimise a non-variable ")
__typedArg1 := tmp11512
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attempt to optimise a non-variable "))
__typedS1, __typedOK1 := TypedString(tmp11512)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attempt to optimise a non-variable ")
__typedArg1 := tmp11512
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


} else {
tmp11559 := Call(__e, PrimFunc(symelement_2), V1646, V1648)


tmp11560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp11559)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp11559
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres11556 Obj

if True == tmp11560 {
tmp11558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres11557 Obj

if True == tmp11558 {
ifres11557 = True


} else {
ifres11557 = False


}

ifres11556 = ifres11557


} else {
ifres11556 = False


}

if True == ifres11556 {
tmp11518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1647, Nil)
}
__typedArg0 := V1647
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1646, tmp11518)
}
__typedArg0 := V1646
__typedArg1 := tmp11518
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4deref, tmp11519)
}
__typedArg0 := symshen_4deref
__typedArg1 := tmp11519
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11535 Obj

if True == tmp11554 {
tmp11552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp11552)
}
__typedArg0 := symlambda
__typedArg1 := tmp11552
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11537 Obj

if True == tmp11553 {
tmp11550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11550)
}
__typedArg0 := tmp11550
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11539 Obj

if True == tmp11551 {
tmp11547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11547)
}
__typedArg0 := tmp11547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11548)
}
__typedArg0 := tmp11548
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11541 Obj

if True == tmp11549 {
tmp11543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11543)
}
__typedArg0 := tmp11543
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11544)
}
__typedArg0 := tmp11544
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11545)
}
__typedArg0 := Nil
__typedArg1 := tmp11545
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11542 Obj

if True == tmp11546 {
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
tmp11520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11520)
}
__typedArg0 := tmp11520
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11522)
}
__typedArg0 := tmp11522
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11523)
}
__typedArg0 := tmp11523
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11525)
}
__typedArg0 := tmp11525
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11526, V1648)
}
__typedArg0 := tmp11526
__typedArg1 := V1648
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11528 := Call(__e, PrimFunc(symshen_4deref_1terms), tmp11524, V1647, tmp11527)


tmp11529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11528, Nil)
}
__typedArg0 := tmp11528
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11521, tmp11529)
}
__typedArg0 := tmp11521
__typedArg1 := tmp11529
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp11530)
}
__typedArg0 := symlambda
__typedArg1 := tmp11530
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11533 {
tmp11531 := MakeNative(func(__e *ControlFlow) {
Z1649 := __e.Get(1)
_ = Z1649
__e.TailApply(PrimFunc(symshen_4deref_1terms), Z1649, V1647, V1648)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp11531, V1646)
return


} else {
__e.Return(V1646)
return
}


}


}


}


}


}, 3)

tmp11591 := Call(__e, ns2_1set, symshen_4deref_1terms, tmp11496)


_ = tmp11591

tmp11592 := MakeNative(func(__e *ControlFlow) {
V1667 := __e.Get(1)
_ = V1667
V1668 := __e.Get(2)
_ = V1668
V1669 := __e.Get(3)
_ = V1669
V1670 := __e.Get(4)
_ = V1670
V1671 := __e.Get(5)
_ = V1671
tmp11768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1668)
}
__typedArg0 := Nil
__typedArg1 := V1668
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11765 Obj

if True == tmp11768 {
tmp11767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1669)
}
__typedArg0 := Nil
__typedArg1 := V1669
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11766 Obj

if True == tmp11767 {
ifres11766 = True


} else {
ifres11766 = False


}

ifres11765 = ifres11766


} else {
ifres11765 = False


}

if True == ifres11765 {
__e.Return(V1671)
return
} else {
tmp11763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11743 Obj

if True == tmp11763 {
tmp11761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11761)
}
__typedArg0 := tmp11761
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11745 Obj

if True == tmp11762 {
tmp11758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11758)
}
__typedArg0 := tmp11758
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, tmp11759)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp11759
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11747 Obj

if True == tmp11760 {
tmp11755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11755)
}
__typedArg0 := tmp11755
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11756)
}
__typedArg0 := tmp11756
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11749 Obj

if True == tmp11757 {
tmp11751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11751)
}
__typedArg0 := tmp11751
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11752)
}
__typedArg0 := tmp11752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11753)
}
__typedArg0 := Nil
__typedArg1 := tmp11753
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11750 Obj

if True == tmp11754 {
ifres11750 = True


} else {
ifres11750 = False


}

ifres11749 = ifres11750


} else {
ifres11749 = False


}

var ifres11748 Obj

if True == ifres11749 {
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

if True == ifres11743 {
tmp11593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11593)
}
__typedArg0 := tmp11593
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11594)
}
__typedArg0 := tmp11594
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1667, tmp11596)
}
__typedArg0 := V1667
__typedArg1 := tmp11596
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11595, tmp11597)
}
__typedArg0 := tmp11595
__typedArg1 := tmp11597
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_7m, tmp11598)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp11598
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4compile_1head), V1667, tmp11599, V1669, V1670, V1671)
return


} else {
tmp11741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11721 Obj

if True == tmp11741 {
tmp11739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11739)
}
__typedArg0 := tmp11739
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11723 Obj

if True == tmp11740 {
tmp11736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11736)
}
__typedArg0 := tmp11736
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_1m, tmp11737)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp11737
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11725 Obj

if True == tmp11738 {
tmp11733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11733)
}
__typedArg0 := tmp11733
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11734)
}
__typedArg0 := tmp11734
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11727 Obj

if True == tmp11735 {
tmp11729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11729)
}
__typedArg0 := tmp11729
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11730)
}
__typedArg0 := tmp11730
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11731)
}
__typedArg0 := Nil
__typedArg1 := tmp11731
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11728 Obj

if True == tmp11732 {
ifres11728 = True


} else {
ifres11728 = False


}

ifres11727 = ifres11728


} else {
ifres11727 = False


}

var ifres11726 Obj

if True == ifres11727 {
ifres11726 = True


} else {
ifres11726 = False


}

ifres11725 = ifres11726


} else {
ifres11725 = False


}

var ifres11724 Obj

if True == ifres11725 {
ifres11724 = True


} else {
ifres11724 = False


}

ifres11723 = ifres11724


} else {
ifres11723 = False


}

var ifres11722 Obj

if True == ifres11723 {
ifres11722 = True


} else {
ifres11722 = False


}

ifres11721 = ifres11722


} else {
ifres11721 = False


}

if True == ifres11721 {
tmp11600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11600)
}
__typedArg0 := tmp11600
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11601)
}
__typedArg0 := tmp11601
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1667, tmp11603)
}
__typedArg0 := V1667
__typedArg1 := tmp11603
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11602, tmp11604)
}
__typedArg0 := tmp11602
__typedArg1 := tmp11604
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_1m, tmp11605)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp11605
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4compile_1head), V1667, tmp11606, V1669, V1670, V1671)
return


} else {
tmp11719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11715 Obj

if True == tmp11719 {
tmp11717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_1m, tmp11717)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp11717
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11716 Obj

if True == tmp11718 {
ifres11716 = True


} else {
ifres11716 = False


}

ifres11715 = ifres11716


} else {
ifres11715 = False


}

if True == ifres11715 {
tmp11607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp11607, V1669, V1670, V1671)
return


} else {
tmp11713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11709 Obj

if True == tmp11713 {
tmp11711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, tmp11711)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp11711
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11710 Obj

if True == tmp11712 {
ifres11710 = True


} else {
ifres11710 = False


}

ifres11709 = ifres11710


} else {
ifres11709 = False


}

if True == ifres11709 {
tmp11608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11608, V1669, V1670, V1671)
return


} else {
tmp11707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11700 Obj

if True == tmp11707 {
tmp11706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1669)
}
__typedArg0 := V1669
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11702 Obj

if True == tmp11706 {
tmp11704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11705 := Call(__e, PrimFunc(symshen_4wildcard_2), tmp11704)


var ifres11703 Obj

if True == tmp11705 {
ifres11703 = True


} else {
ifres11703 = False


}

ifres11702 = ifres11703


} else {
ifres11702 = False


}

var ifres11701 Obj

if True == ifres11702 {
ifres11701 = True


} else {
ifres11701 = False


}

ifres11700 = ifres11701


} else {
ifres11700 = False


}

if True == ifres11700 {
tmp11609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1669)
}
__typedArg0 := V1669
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4compile_1head), V1667, tmp11609, tmp11610, V1670, V1671)
return


} else {
tmp11698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11694 Obj

if True == tmp11698 {
tmp11696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp11696)
}
__typedArg0 := tmp11696
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres11695 Obj

if True == tmp11697 {
ifres11695 = True


} else {
ifres11695 = False


}

ifres11694 = ifres11695


} else {
ifres11694 = False


}

if True == ifres11694 {
__e.TailApply(PrimFunc(symshen_4variable_1case), V1667, V1668, V1669, V1670, V1671)
return
} else {
tmp11692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_1m, V1667)
}
__typedArg0 := symshen_4_1m
__typedArg1 := V1667
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11685 Obj

if True == tmp11692 {
tmp11691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11687 Obj

if True == tmp11691 {
tmp11689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11690 := Call(__e, PrimFunc(symatom_2), tmp11689)


var ifres11688 Obj

if True == tmp11690 {
ifres11688 = True


} else {
ifres11688 = False


}

ifres11687 = ifres11688


} else {
ifres11687 = False


}

var ifres11686 Obj

if True == ifres11687 {
ifres11686 = True


} else {
ifres11686 = False


}

ifres11685 = ifres11686


} else {
ifres11685 = False


}

if True == ifres11685 {
__e.TailApply(PrimFunc(symshen_4atom_1case_1minus), V1668, V1669, V1670, V1671)
return
} else {
tmp11683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_1m, V1667)
}
__typedArg0 := symshen_4_1m
__typedArg1 := V1667
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11653 Obj

if True == tmp11683 {
tmp11682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11655 Obj

if True == tmp11682 {
tmp11680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11680)
}
__typedArg0 := tmp11680
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11657 Obj

if True == tmp11681 {
tmp11677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11677)
}
__typedArg0 := tmp11677
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp11678)
}
__typedArg0 := symcons
__typedArg1 := tmp11678
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11659 Obj

if True == tmp11679 {
tmp11674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11674)
}
__typedArg0 := tmp11674
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11675)
}
__typedArg0 := tmp11675
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11661 Obj

if True == tmp11676 {
tmp11670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11670)
}
__typedArg0 := tmp11670
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11671)
}
__typedArg0 := tmp11671
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11672)
}
__typedArg0 := tmp11672
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11663 Obj

if True == tmp11673 {
tmp11665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11665)
}
__typedArg0 := tmp11665
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11666)
}
__typedArg0 := tmp11666
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11667)
}
__typedArg0 := tmp11667
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11668)
}
__typedArg0 := Nil
__typedArg1 := tmp11668
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11664 Obj

if True == tmp11669 {
ifres11664 = True


} else {
ifres11664 = False


}

ifres11663 = ifres11664


} else {
ifres11663 = False


}

var ifres11662 Obj

if True == ifres11663 {
ifres11662 = True


} else {
ifres11662 = False


}

ifres11661 = ifres11662


} else {
ifres11661 = False


}

var ifres11660 Obj

if True == ifres11661 {
ifres11660 = True


} else {
ifres11660 = False


}

ifres11659 = ifres11660


} else {
ifres11659 = False


}

var ifres11658 Obj

if True == ifres11659 {
ifres11658 = True


} else {
ifres11658 = False


}

ifres11657 = ifres11658


} else {
ifres11657 = False


}

var ifres11656 Obj

if True == ifres11657 {
ifres11656 = True


} else {
ifres11656 = False


}

ifres11655 = ifres11656


} else {
ifres11655 = False


}

var ifres11654 Obj

if True == ifres11655 {
ifres11654 = True


} else {
ifres11654 = False


}

ifres11653 = ifres11654


} else {
ifres11653 = False


}

if True == ifres11653 {
__e.TailApply(PrimFunc(symshen_4cons_1case_1minus), V1668, V1669, V1670, V1671)
return
} else {
tmp11651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, V1667)
}
__typedArg0 := symshen_4_7m
__typedArg1 := V1667
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11644 Obj

if True == tmp11651 {
tmp11650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11646 Obj

if True == tmp11650 {
tmp11648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11649 := Call(__e, PrimFunc(symatom_2), tmp11648)


var ifres11647 Obj

if True == tmp11649 {
ifres11647 = True


} else {
ifres11647 = False


}

ifres11646 = ifres11647


} else {
ifres11646 = False


}

var ifres11645 Obj

if True == ifres11646 {
ifres11645 = True


} else {
ifres11645 = False


}

ifres11644 = ifres11645


} else {
ifres11644 = False


}

if True == ifres11644 {
__e.TailApply(PrimFunc(symshen_4atom_1case_1plus), V1668, V1669, V1670, V1671)
return
} else {
tmp11642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, V1667)
}
__typedArg0 := symshen_4_7m
__typedArg1 := V1667
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11612 Obj

if True == tmp11642 {
tmp11641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11614 Obj

if True == tmp11641 {
tmp11639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11639)
}
__typedArg0 := tmp11639
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11616 Obj

if True == tmp11640 {
tmp11636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11636)
}
__typedArg0 := tmp11636
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp11637)
}
__typedArg0 := symcons
__typedArg1 := tmp11637
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11618 Obj

if True == tmp11638 {
tmp11633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11633)
}
__typedArg0 := tmp11633
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11634)
}
__typedArg0 := tmp11634
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11620 Obj

if True == tmp11635 {
tmp11629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11629)
}
__typedArg0 := tmp11629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11630)
}
__typedArg0 := tmp11630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11631)
}
__typedArg0 := tmp11631
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11622 Obj

if True == tmp11632 {
tmp11624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11624)
}
__typedArg0 := tmp11624
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11625)
}
__typedArg0 := tmp11625
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11626)
}
__typedArg0 := tmp11626
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11627)
}
__typedArg0 := Nil
__typedArg1 := tmp11627
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11623 Obj

if True == tmp11628 {
ifres11623 = True


} else {
ifres11623 = False


}

ifres11622 = ifres11623


} else {
ifres11622 = False


}

var ifres11621 Obj

if True == ifres11622 {
ifres11621 = True


} else {
ifres11621 = False


}

ifres11620 = ifres11621


} else {
ifres11620 = False


}

var ifres11619 Obj

if True == ifres11620 {
ifres11619 = True


} else {
ifres11619 = False


}

ifres11618 = ifres11619


} else {
ifres11618 = False


}

var ifres11617 Obj

if True == ifres11618 {
ifres11617 = True


} else {
ifres11617 = False


}

ifres11616 = ifres11617


} else {
ifres11616 = False


}

var ifres11615 Obj

if True == ifres11616 {
ifres11615 = True


} else {
ifres11615 = False


}

ifres11614 = ifres11615


} else {
ifres11614 = False


}

var ifres11613 Obj

if True == ifres11614 {
ifres11613 = True


} else {
ifres11613 = False


}

ifres11612 = ifres11613


} else {
ifres11612 = False


}

if True == ifres11612 {
__e.TailApply(PrimFunc(symshen_4cons_1case_1plus), V1668, V1669, V1670, V1671)
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.compile-head"))
}
__typedArg0 := MakeString("implementation error in shen.compile-head")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
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

tmp11769 := Call(__e, ns2_1set, symshen_4compile_1head, tmp11592)


_ = tmp11769

tmp11770 := MakeNative(func(__e *ControlFlow) {
V1682 := __e.Get(1)
_ = V1682
V1683 := __e.Get(2)
_ = V1683
V1684 := __e.Get(3)
_ = V1684
V1685 := __e.Get(4)
_ = V1685
V1686 := __e.Get(5)
_ = V1686
tmp11791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1683)
}
__typedArg0 := V1683
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11788 Obj

if True == tmp11791 {
tmp11790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1684)
}
__typedArg0 := V1684
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11789 Obj

if True == tmp11790 {
ifres11789 = True


} else {
ifres11789 = False


}

ifres11788 = ifres11789


} else {
ifres11788 = False


}

if True == ifres11788 {
tmp11785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1684)
}
__typedArg0 := V1684
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp11785)
}
__typedArg0 := tmp11785
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp11786 {
tmp11771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1683)
}
__typedArg0 := V1683
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1684)
}
__typedArg0 := V1684
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1684)
}
__typedArg0 := V1684
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1683)
}
__typedArg0 := V1683
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11775 := Call(__e, PrimFunc(symsubst), tmp11773, tmp11774, V1686)


__e.TailApply(PrimFunc(symshen_4compile_1head), V1682, tmp11771, tmp11772, V1685, tmp11775)
return


} else {
tmp11776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1683)
}
__typedArg0 := V1683
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1684)
}
__typedArg0 := V1684
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1683)
}
__typedArg0 := V1683
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1684)
}
__typedArg0 := V1684
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11780 := Call(__e, PrimFunc(symshen_4compile_1head), V1682, tmp11778, tmp11779, V1685, V1686)


tmp11781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11780, Nil)
}
__typedArg0 := tmp11780
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11777, tmp11781)
}
__typedArg0 := tmp11777
__typedArg1 := tmp11781
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11776, tmp11782)
}
__typedArg0 := tmp11776
__typedArg1 := tmp11782
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11783)
}
__typedArg0 := symlet
__typedArg1 := tmp11783
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.variable-case"))
}
__typedArg0 := MakeString("implementation error in shen.variable-case")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 5)

tmp11792 := Call(__e, ns2_1set, symshen_4variable_1case, tmp11770)


_ = tmp11792

tmp11793 := MakeNative(func(__e *ControlFlow) {
V1695 := __e.Get(1)
_ = V1695
V1696 := __e.Get(2)
_ = V1696
V1697 := __e.Get(3)
_ = V1697
V1698 := __e.Get(4)
_ = V1698
tmp11818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1695)
}
__typedArg0 := V1695
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11815 Obj

if True == tmp11818 {
tmp11817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1696)
}
__typedArg0 := V1696
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11816 Obj

if True == tmp11817 {
ifres11816 = True


} else {
ifres11816 = False


}

ifres11815 = ifres11816


} else {
ifres11815 = False


}

if True == ifres11815 {
tmp11794 := MakeNative(func(__e *ControlFlow) {
W1699 := __e.Get(1)
_ = W1699
tmp11795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1696)
}
__typedArg0 := V1696
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1697, Nil)
}
__typedArg0 := V1697
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11795, tmp11796)
}
__typedArg0 := tmp11795
__typedArg1 := tmp11796
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp11797)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp11797
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1695)
}
__typedArg0 := V1695
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11799, Nil)
}
__typedArg0 := tmp11799
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1699, tmp11800)
}
__typedArg0 := W1699
__typedArg1 := tmp11800
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp11801)
}
__typedArg0 := sym_a
__typedArg1 := tmp11801
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1695)
}
__typedArg0 := V1695
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1696)
}
__typedArg0 := V1696
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11805 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp11803, tmp11804, V1697, V1698)


tmp11806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11805, tmp11806)
}
__typedArg0 := tmp11805
__typedArg1 := tmp11806
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11802, tmp11807)
}
__typedArg0 := tmp11802
__typedArg1 := tmp11807
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11808)
}
__typedArg0 := symif
__typedArg1 := tmp11808
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11809, Nil)
}
__typedArg0 := tmp11809
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11798, tmp11810)
}
__typedArg0 := tmp11798
__typedArg1 := tmp11810
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1699, tmp11811)
}
__typedArg0 := W1699
__typedArg1 := tmp11811
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11812)
}
__typedArg0 := symlet
__typedArg1 := tmp11812
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp11813 := Call(__e, PrimFunc(symgensym), symTm)


__e.TailApply(tmp11794, tmp11813)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.atom-case-minus"))
}
__typedArg0 := MakeString("implementation error in shen.atom-case-minus")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 4)

tmp11819 := Call(__e, ns2_1set, symshen_4atom_1case_1minus, tmp11793)


_ = tmp11819

tmp11820 := MakeNative(func(__e *ControlFlow) {
V1708 := __e.Get(1)
_ = V1708
V1709 := __e.Get(2)
_ = V1709
V1710 := __e.Get(3)
_ = V1710
V1711 := __e.Get(4)
_ = V1711
tmp11885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1708)
}
__typedArg0 := V1708
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11855 Obj

if True == tmp11885 {
tmp11883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1708)
}
__typedArg0 := V1708
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11883)
}
__typedArg0 := tmp11883
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11857 Obj

if True == tmp11884 {
tmp11880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1708)
}
__typedArg0 := V1708
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11880)
}
__typedArg0 := tmp11880
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp11881)
}
__typedArg0 := symcons
__typedArg1 := tmp11881
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11859 Obj

if True == tmp11882 {
tmp11877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1708)
}
__typedArg0 := V1708
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11877)
}
__typedArg0 := tmp11877
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11878)
}
__typedArg0 := tmp11878
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11861 Obj

if True == tmp11879 {
tmp11873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1708)
}
__typedArg0 := V1708
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11873)
}
__typedArg0 := tmp11873
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11874)
}
__typedArg0 := tmp11874
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11875)
}
__typedArg0 := tmp11875
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11863 Obj

if True == tmp11876 {
tmp11868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1708)
}
__typedArg0 := V1708
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11868)
}
__typedArg0 := tmp11868
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11869)
}
__typedArg0 := tmp11869
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11870)
}
__typedArg0 := tmp11870
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11871)
}
__typedArg0 := Nil
__typedArg1 := tmp11871
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11865 Obj

if True == tmp11872 {
tmp11867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1709)
}
__typedArg0 := V1709
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11866 Obj

if True == tmp11867 {
ifres11866 = True


} else {
ifres11866 = False


}

ifres11865 = ifres11866


} else {
ifres11865 = False


}

var ifres11864 Obj

if True == ifres11865 {
ifres11864 = True


} else {
ifres11864 = False


}

ifres11863 = ifres11864


} else {
ifres11863 = False


}

var ifres11862 Obj

if True == ifres11863 {
ifres11862 = True


} else {
ifres11862 = False


}

ifres11861 = ifres11862


} else {
ifres11861 = False


}

var ifres11860 Obj

if True == ifres11861 {
ifres11860 = True


} else {
ifres11860 = False


}

ifres11859 = ifres11860


} else {
ifres11859 = False


}

var ifres11858 Obj

if True == ifres11859 {
ifres11858 = True


} else {
ifres11858 = False


}

ifres11857 = ifres11858


} else {
ifres11857 = False


}

var ifres11856 Obj

if True == ifres11857 {
ifres11856 = True


} else {
ifres11856 = False


}

ifres11855 = ifres11856


} else {
ifres11855 = False


}

if True == ifres11855 {
tmp11821 := MakeNative(func(__e *ControlFlow) {
W1712 := __e.Get(1)
_ = W1712
tmp11822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1709)
}
__typedArg0 := V1709
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1710, Nil)
}
__typedArg0 := V1710
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11822, tmp11823)
}
__typedArg0 := tmp11822
__typedArg1 := tmp11823
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp11824)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp11824
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1712, Nil)
}
__typedArg0 := W1712
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp11826)
}
__typedArg0 := symcons_2
__typedArg1 := tmp11826
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1708)
}
__typedArg0 := V1708
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11828)
}
__typedArg0 := tmp11828
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11829)
}
__typedArg0 := tmp11829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1708)
}
__typedArg0 := V1708
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11831)
}
__typedArg0 := tmp11831
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11832)
}
__typedArg0 := tmp11832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11833)
}
__typedArg0 := tmp11833
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1708)
}
__typedArg0 := V1708
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11834, tmp11835)
}
__typedArg0 := tmp11834
__typedArg1 := tmp11835
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11830, tmp11836)
}
__typedArg0 := tmp11830
__typedArg1 := tmp11836
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1712, Nil)
}
__typedArg0 := W1712
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhd, tmp11838)
}
__typedArg0 := symhd
__typedArg1 := tmp11838
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1712, Nil)
}
__typedArg0 := W1712
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtl, tmp11840)
}
__typedArg0 := symtl
__typedArg1 := tmp11840
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1709)
}
__typedArg0 := V1709
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11841, tmp11842)
}
__typedArg0 := tmp11841
__typedArg1 := tmp11842
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11839, tmp11843)
}
__typedArg0 := tmp11839
__typedArg1 := tmp11843
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11845 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp11837, tmp11844, V1710, V1711)


tmp11846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11845, tmp11846)
}
__typedArg0 := tmp11845
__typedArg1 := tmp11846
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11827, tmp11847)
}
__typedArg0 := tmp11827
__typedArg1 := tmp11847
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11848)
}
__typedArg0 := symif
__typedArg1 := tmp11848
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11849, Nil)
}
__typedArg0 := tmp11849
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11825, tmp11850)
}
__typedArg0 := tmp11825
__typedArg1 := tmp11850
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1712, tmp11851)
}
__typedArg0 := W1712
__typedArg1 := tmp11851
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11852)
}
__typedArg0 := symlet
__typedArg1 := tmp11852
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp11853 := Call(__e, PrimFunc(symgensym), symTm)


__e.TailApply(tmp11821, tmp11853)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.cons-case-minus"))
}
__typedArg0 := MakeString("implementation error in shen.cons-case-minus")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 4)

tmp11886 := Call(__e, ns2_1set, symshen_4cons_1case_1minus, tmp11820)


_ = tmp11886

tmp11887 := MakeNative(func(__e *ControlFlow) {
V1721 := __e.Get(1)
_ = V1721
V1722 := __e.Get(2)
_ = V1722
V1723 := __e.Get(3)
_ = V1723
V1724 := __e.Get(4)
_ = V1724
tmp11933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1721)
}
__typedArg0 := V1721
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11930 Obj

if True == tmp11933 {
tmp11932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1722)
}
__typedArg0 := V1722
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11931 Obj

if True == tmp11932 {
ifres11931 = True


} else {
ifres11931 = False


}

ifres11930 = ifres11931


} else {
ifres11930 = False


}

if True == ifres11930 {
tmp11888 := MakeNative(func(__e *ControlFlow) {
W1725 := __e.Get(1)
_ = W1725
tmp11889 := MakeNative(func(__e *ControlFlow) {
W1726 := __e.Get(1)
_ = W1726
tmp11890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1722)
}
__typedArg0 := V1722
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1723, Nil)
}
__typedArg0 := V1723
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11890, tmp11891)
}
__typedArg0 := tmp11890
__typedArg1 := tmp11891
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp11892)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp11892
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1721)
}
__typedArg0 := V1721
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1722)
}
__typedArg0 := V1722
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11896 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11894, tmp11895, V1723, V1724)


tmp11897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11896, Nil)
}
__typedArg0 := tmp11896
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp11897)
}
__typedArg0 := symfreeze
__typedArg1 := tmp11897
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1721)
}
__typedArg0 := V1721
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11899, Nil)
}
__typedArg0 := tmp11899
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1725, tmp11900)
}
__typedArg0 := W1725
__typedArg1 := tmp11900
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp11901)
}
__typedArg0 := sym_a
__typedArg1 := tmp11901
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1726, Nil)
}
__typedArg0 := W1726
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp11903)
}
__typedArg0 := symthaw
__typedArg1 := tmp11903
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1725, Nil)
}
__typedArg0 := W1725
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4pvar_2, tmp11905)
}
__typedArg0 := symshen_4pvar_2
__typedArg1 := tmp11905
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1721)
}
__typedArg0 := V1721
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11908 := Call(__e, PrimFunc(symshen_4demode), tmp11907)


tmp11909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1726, Nil)
}
__typedArg0 := W1726
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1723, tmp11909)
}
__typedArg0 := V1723
__typedArg1 := tmp11909
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11908, tmp11910)
}
__typedArg0 := tmp11908
__typedArg1 := tmp11910
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1725, tmp11911)
}
__typedArg0 := W1725
__typedArg1 := tmp11911
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4bind_b, tmp11912)
}
__typedArg0 := symshen_4bind_b
__typedArg1 := tmp11912
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11913, tmp11914)
}
__typedArg0 := tmp11913
__typedArg1 := tmp11914
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11906, tmp11915)
}
__typedArg0 := tmp11906
__typedArg1 := tmp11915
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11916)
}
__typedArg0 := symif
__typedArg1 := tmp11916
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11917, Nil)
}
__typedArg0 := tmp11917
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11904, tmp11918)
}
__typedArg0 := tmp11904
__typedArg1 := tmp11918
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11902, tmp11919)
}
__typedArg0 := tmp11902
__typedArg1 := tmp11919
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11920)
}
__typedArg0 := symif
__typedArg1 := tmp11920
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11921, Nil)
}
__typedArg0 := tmp11921
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11898, tmp11922)
}
__typedArg0 := tmp11898
__typedArg1 := tmp11922
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1726, tmp11923)
}
__typedArg0 := W1726
__typedArg1 := tmp11923
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11893, tmp11924)
}
__typedArg0 := tmp11893
__typedArg1 := tmp11924
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1725, tmp11925)
}
__typedArg0 := W1725
__typedArg1 := tmp11925
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11926)
}
__typedArg0 := symlet
__typedArg1 := tmp11926
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp11927 := Call(__e, PrimFunc(symgensym), symGoTo)


__e.TailApply(tmp11889, tmp11927)
return


}, 1)

tmp11928 := Call(__e, PrimFunc(symgensym), symTm)


__e.TailApply(tmp11888, tmp11928)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.atom-case-plus"))
}
__typedArg0 := MakeString("implementation error in shen.atom-case-plus")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 4)

tmp11934 := Call(__e, ns2_1set, symshen_4atom_1case_1plus, tmp11887)


_ = tmp11934

tmp11935 := MakeNative(func(__e *ControlFlow) {
V1735 := __e.Get(1)
_ = V1735
V1736 := __e.Get(2)
_ = V1736
V1737 := __e.Get(3)
_ = V1737
V1738 := __e.Get(4)
_ = V1738
tmp12031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1735)
}
__typedArg0 := V1735
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12001 Obj

if True == tmp12031 {
tmp12029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1735)
}
__typedArg0 := V1735
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp12029)
}
__typedArg0 := tmp12029
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12003 Obj

if True == tmp12030 {
tmp12026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1735)
}
__typedArg0 := V1735
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp12026)
}
__typedArg0 := tmp12026
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp12027)
}
__typedArg0 := symcons
__typedArg1 := tmp12027
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12005 Obj

if True == tmp12028 {
tmp12023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1735)
}
__typedArg0 := V1735
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12023)
}
__typedArg0 := tmp12023
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp12024)
}
__typedArg0 := tmp12024
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12007 Obj

if True == tmp12025 {
tmp12019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1735)
}
__typedArg0 := V1735
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12019)
}
__typedArg0 := tmp12019
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12020)
}
__typedArg0 := tmp12020
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp12021)
}
__typedArg0 := tmp12021
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12009 Obj

if True == tmp12022 {
tmp12014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1735)
}
__typedArg0 := V1735
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12014)
}
__typedArg0 := tmp12014
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12015)
}
__typedArg0 := tmp12015
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12016)
}
__typedArg0 := tmp12016
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp12017)
}
__typedArg0 := Nil
__typedArg1 := tmp12017
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12011 Obj

if True == tmp12018 {
tmp12013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1736)
}
__typedArg0 := V1736
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12012 Obj

if True == tmp12013 {
ifres12012 = True


} else {
ifres12012 = False


}

ifres12011 = ifres12012


} else {
ifres12011 = False


}

var ifres12010 Obj

if True == ifres12011 {
ifres12010 = True


} else {
ifres12010 = False


}

ifres12009 = ifres12010


} else {
ifres12009 = False


}

var ifres12008 Obj

if True == ifres12009 {
ifres12008 = True


} else {
ifres12008 = False


}

ifres12007 = ifres12008


} else {
ifres12007 = False


}

var ifres12006 Obj

if True == ifres12007 {
ifres12006 = True


} else {
ifres12006 = False


}

ifres12005 = ifres12006


} else {
ifres12005 = False


}

var ifres12004 Obj

if True == ifres12005 {
ifres12004 = True


} else {
ifres12004 = False


}

ifres12003 = ifres12004


} else {
ifres12003 = False


}

var ifres12002 Obj

if True == ifres12003 {
ifres12002 = True


} else {
ifres12002 = False


}

ifres12001 = ifres12002


} else {
ifres12001 = False


}

if True == ifres12001 {
tmp11936 := MakeNative(func(__e *ControlFlow) {
W1739 := __e.Get(1)
_ = W1739
tmp11937 := MakeNative(func(__e *ControlFlow) {
W1740 := __e.Get(1)
_ = W1740
tmp11938 := MakeNative(func(__e *ControlFlow) {
W1741 := __e.Get(1)
_ = W1741
tmp11939 := MakeNative(func(__e *ControlFlow) {
W1742 := __e.Get(1)
_ = W1742
tmp11940 := MakeNative(func(__e *ControlFlow) {
W1743 := __e.Get(1)
_ = W1743
tmp11941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1736)
}
__typedArg0 := V1736
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1737, Nil)
}
__typedArg0 := V1737
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11941, tmp11942)
}
__typedArg0 := tmp11941
__typedArg1 := tmp11942
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp11943)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp11943
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1735)
}
__typedArg0 := V1735
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1736)
}
__typedArg0 := V1736
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11947 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11945, tmp11946, V1737, V1738)


tmp11948 := Call(__e, PrimFunc(symshen_4goto), W1741, tmp11947)


tmp11949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1739, Nil)
}
__typedArg0 := W1739
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp11949)
}
__typedArg0 := symcons_2
__typedArg1 := tmp11949
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1735)
}
__typedArg0 := V1735
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11951)
}
__typedArg0 := tmp11951
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1739, Nil)
}
__typedArg0 := W1739
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhd, tmp11953)
}
__typedArg0 := symhd
__typedArg1 := tmp11953
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1739, Nil)
}
__typedArg0 := W1739
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtl, tmp11955)
}
__typedArg0 := symtl
__typedArg1 := tmp11955
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11956, Nil)
}
__typedArg0 := tmp11956
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11954, tmp11957)
}
__typedArg0 := tmp11954
__typedArg1 := tmp11957
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11959 := Call(__e, PrimFunc(symshen_4invoke), W1740, W1741)


tmp11960 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11952, tmp11958, V1737, tmp11959)


tmp11961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1739, Nil)
}
__typedArg0 := W1739
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4pvar_2, tmp11961)
}
__typedArg0 := symshen_4pvar_2
__typedArg1 := tmp11961
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11963 := Call(__e, PrimFunc(symshen_4demode), W1742)


tmp11964 := Call(__e, PrimFunc(symshen_4invoke), W1740, W1741)


tmp11965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11964, Nil)
}
__typedArg0 := tmp11964
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp11965)
}
__typedArg0 := symfreeze
__typedArg1 := tmp11965
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11966, Nil)
}
__typedArg0 := tmp11966
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1737, tmp11967)
}
__typedArg0 := V1737
__typedArg1 := tmp11967
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11963, tmp11968)
}
__typedArg0 := tmp11963
__typedArg1 := tmp11968
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1739, tmp11969)
}
__typedArg0 := W1739
__typedArg1 := tmp11969
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4bind_b, tmp11970)
}
__typedArg0 := symshen_4bind_b
__typedArg1 := tmp11970
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11972 := Call(__e, PrimFunc(symshen_4stpart), W1743, tmp11971, V1737)


tmp11973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11972, tmp11973)
}
__typedArg0 := tmp11972
__typedArg1 := tmp11973
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11962, tmp11974)
}
__typedArg0 := tmp11962
__typedArg1 := tmp11974
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11975)
}
__typedArg0 := symif
__typedArg1 := tmp11975
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11976, Nil)
}
__typedArg0 := tmp11976
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11960, tmp11977)
}
__typedArg0 := tmp11960
__typedArg1 := tmp11977
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11950, tmp11978)
}
__typedArg0 := tmp11950
__typedArg1 := tmp11978
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11979)
}
__typedArg0 := symif
__typedArg1 := tmp11979
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11980, Nil)
}
__typedArg0 := tmp11980
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11948, tmp11981)
}
__typedArg0 := tmp11948
__typedArg1 := tmp11981
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1740, tmp11982)
}
__typedArg0 := W1740
__typedArg1 := tmp11982
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11944, tmp11983)
}
__typedArg0 := tmp11944
__typedArg1 := tmp11983
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W1739, tmp11984)
}
__typedArg0 := W1739
__typedArg1 := tmp11984
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11985)
}
__typedArg0 := symlet
__typedArg1 := tmp11985
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp11986 := Call(__e, PrimFunc(symshen_4extract_1vars), W1742)


__e.TailApply(tmp11940, tmp11986)
return


}, 1)

tmp11987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1735)
}
__typedArg0 := V1735
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11988 := Call(__e, PrimFunc(symshen_4tame), tmp11987)


__e.TailApply(tmp11939, tmp11988)
return


}, 1)

tmp11989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1735)
}
__typedArg0 := V1735
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11989)
}
__typedArg0 := tmp11989
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11990)
}
__typedArg0 := tmp11990
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1735)
}
__typedArg0 := V1735
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11992)
}
__typedArg0 := tmp11992
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11993)
}
__typedArg0 := tmp11993
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11994)
}
__typedArg0 := tmp11994
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11991, tmp11995)
}
__typedArg0 := tmp11991
__typedArg1 := tmp11995
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11997 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp11996)


__e.TailApply(tmp11938, tmp11997)
return


}, 1)

tmp11998 := Call(__e, PrimFunc(symgensym), symGoTo)


__e.TailApply(tmp11937, tmp11998)
return


}, 1)

tmp11999 := Call(__e, PrimFunc(symgensym), symTm)


__e.TailApply(tmp11936, tmp11999)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.cons-case-plus"))
}
__typedArg0 := MakeString("implementation error in shen.cons-case-plus")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 4)

tmp12032 := Call(__e, ns2_1set, symshen_4cons_1case_1plus, tmp11935)


_ = tmp12032

tmp12033 := MakeNative(func(__e *ControlFlow) {
V1744 := __e.Get(1)
_ = V1744
tmp12070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1744)
}
__typedArg0 := V1744
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12057 Obj

if True == tmp12070 {
tmp12068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1744)
}
__typedArg0 := V1744
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, tmp12068)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp12068
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12059 Obj

if True == tmp12069 {
tmp12066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1744)
}
__typedArg0 := V1744
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp12066)
}
__typedArg0 := tmp12066
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12061 Obj

if True == tmp12067 {
tmp12063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1744)
}
__typedArg0 := V1744
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12063)
}
__typedArg0 := tmp12063
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp12064)
}
__typedArg0 := Nil
__typedArg1 := tmp12064
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12062 Obj

if True == tmp12065 {
ifres12062 = True


} else {
ifres12062 = False


}

ifres12061 = ifres12062


} else {
ifres12061 = False


}

var ifres12060 Obj

if True == ifres12061 {
ifres12060 = True


} else {
ifres12060 = False


}

ifres12059 = ifres12060


} else {
ifres12059 = False


}

var ifres12058 Obj

if True == ifres12059 {
ifres12058 = True


} else {
ifres12058 = False


}

ifres12057 = ifres12058


} else {
ifres12057 = False


}

if True == ifres12057 {
tmp12034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1744)
}
__typedArg0 := V1744
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp12034)
}
__typedArg0 := tmp12034
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4demode), tmp12035)
return


} else {
tmp12055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1744)
}
__typedArg0 := V1744
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12042 Obj

if True == tmp12055 {
tmp12053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1744)
}
__typedArg0 := V1744
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_1m, tmp12053)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp12053
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12044 Obj

if True == tmp12054 {
tmp12051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1744)
}
__typedArg0 := V1744
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp12051)
}
__typedArg0 := tmp12051
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12046 Obj

if True == tmp12052 {
tmp12048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1744)
}
__typedArg0 := V1744
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12048)
}
__typedArg0 := tmp12048
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp12049)
}
__typedArg0 := Nil
__typedArg1 := tmp12049
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12047 Obj

if True == tmp12050 {
ifres12047 = True


} else {
ifres12047 = False


}

ifres12046 = ifres12047


} else {
ifres12046 = False


}

var ifres12045 Obj

if True == ifres12046 {
ifres12045 = True


} else {
ifres12045 = False


}

ifres12044 = ifres12045


} else {
ifres12044 = False


}

var ifres12043 Obj

if True == ifres12044 {
ifres12043 = True


} else {
ifres12043 = False


}

ifres12042 = ifres12043


} else {
ifres12042 = False


}

if True == ifres12042 {
tmp12036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1744)
}
__typedArg0 := V1744
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp12036)
}
__typedArg0 := tmp12036
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4demode), tmp12037)
return


} else {
tmp12040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1744)
}
__typedArg0 := V1744
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12040 {
tmp12038 := MakeNative(func(__e *ControlFlow) {
Z1745 := __e.Get(1)
_ = Z1745
__e.TailApply(PrimFunc(symshen_4demode), Z1745)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp12038, V1744)
return


} else {
__e.Return(V1744)
return
}


}


}


}, 1)

tmp12071 := Call(__e, ns2_1set, symshen_4demode, tmp12033)


_ = tmp12071

tmp12072 := MakeNative(func(__e *ControlFlow) {
V1746 := __e.Get(1)
_ = V1746
tmp12077 := Call(__e, PrimFunc(symshen_4wildcard_2), V1746)


if True == tmp12077 {
__e.TailApply(PrimFunc(symgensym), symY)
return
} else {
tmp12075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1746)
}
__typedArg0 := V1746
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12075 {
tmp12073 := MakeNative(func(__e *ControlFlow) {
Z1747 := __e.Get(1)
_ = Z1747
__e.TailApply(PrimFunc(symshen_4tame), Z1747)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp12073, V1746)
return


} else {
__e.Return(V1746)
return
}


}


}, 1)

tmp12078 := Call(__e, ns2_1set, symshen_4tame, tmp12072)


_ = tmp12078

tmp12079 := MakeNative(func(__e *ControlFlow) {
V1748 := __e.Get(1)
_ = V1748
V1749 := __e.Get(2)
_ = V1749
tmp12082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1748)
}
__typedArg0 := Nil
__typedArg1 := V1748
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12082 {
tmp12080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1749, Nil)
}
__typedArg0 := V1749
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp12080)
}
__typedArg0 := symfreeze
__typedArg1 := tmp12080
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4goto_1h), V1748, V1749)
return
}


}, 2)

tmp12083 := Call(__e, ns2_1set, symshen_4goto, tmp12079)


_ = tmp12083

tmp12084 := MakeNative(func(__e *ControlFlow) {
V1750 := __e.Get(1)
_ = V1750
V1751 := __e.Get(2)
_ = V1751
tmp12093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1750)
}
__typedArg0 := Nil
__typedArg1 := V1750
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12093 {
__e.Return(V1751)
return
} else {
tmp12091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1750)
}
__typedArg0 := V1750
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12091 {
tmp12085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1750)
}
__typedArg0 := V1750
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1750)
}
__typedArg0 := V1750
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12087 := Call(__e, PrimFunc(symshen_4goto_1h), tmp12086, V1751)


tmp12088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12087, Nil)
}
__typedArg0 := tmp12087
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12085, tmp12088)
}
__typedArg0 := tmp12085
__typedArg1 := tmp12088
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp12089)
}
__typedArg0 := symlambda
__typedArg1 := tmp12089
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.goto-h"))
}
__typedArg0 := MakeString("partial function shen.goto-h")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp12094 := Call(__e, ns2_1set, symshen_4goto_1h, tmp12084)


_ = tmp12094

tmp12095 := MakeNative(func(__e *ControlFlow) {
V1752 := __e.Get(1)
_ = V1752
V1753 := __e.Get(2)
_ = V1753
tmp12098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1753)
}
__typedArg0 := Nil
__typedArg1 := V1753
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12098 {
tmp12096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1752, Nil)
}
__typedArg0 := V1752
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp12096)
}
__typedArg0 := symthaw
__typedArg1 := tmp12096
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1752, V1753)
}
__typedArg0 := V1752
__typedArg1 := V1753
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
}


}, 2)

tmp12099 := Call(__e, ns2_1set, symshen_4invoke, tmp12095)


_ = tmp12099

tmp12100 := MakeNative(func(__e *ControlFlow) {
V1754 := __e.Get(1)
_ = V1754
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1754, sym__)
}
__typedArg0 := V1754
__typedArg1 := sym__
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp12101 := Call(__e, ns2_1set, symshen_4wildcard_2, tmp12100)


_ = tmp12101

tmp12102 := MakeNative(func(__e *ControlFlow) {
V1755 := __e.Get(1)
_ = V1755
tmp12103 := MakeNative(func(__e *ControlFlow) {
tmp12108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V1755)
}
__typedArg0 := V1755
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp12108 {
tmp12105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1755, MakeNumber(0))
}
__typedArg0 := V1755
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp12106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp12105, symshen_4pvar)
}
__typedArg0 := tmp12105
__typedArg1 := symshen_4pvar
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12106 {
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

tmp12109 := MakeNative(func(__e *ControlFlow) {
Z1756 := __e.Get(1)
_ = Z1756
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp12103, tmp12109)
return


}, 1)

tmp12110 := Call(__e, ns2_1set, symshen_4pvar_2, tmp12102)


_ = tmp12110

tmp12111 := MakeNative(func(__e *ControlFlow) {
V1757 := __e.Get(1)
_ = V1757
V1758 := __e.Get(2)
_ = V1758
tmp12118 := Call(__e, PrimFunc(symshen_4pvar_2), V1757)


if True == tmp12118 {
tmp12112 := MakeNative(func(__e *ControlFlow) {
W1759 := __e.Get(1)
_ = W1759
tmp12114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W1759, symshen_4_1null_1)
}
__typedArg0 := W1759
__typedArg1 := symshen_4_1null_1
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12114 {
__e.Return(V1757)
return
} else {
__e.TailApply(PrimFunc(symshen_4lazyderef), W1759, V1758)
return
}


}, 1)

tmp12115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1757, MakeNumber(1))
}
__typedArg0 := V1757
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp12116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1758, tmp12115)
}
__typedArg0 := V1758
__typedArg1 := tmp12115
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp12112, tmp12116)
return


} else {
__e.Return(V1757)
return
}


}, 2)

tmp12119 := Call(__e, ns2_1set, symshen_4lazyderef, tmp12111)


_ = tmp12119

tmp12120 := MakeNative(func(__e *ControlFlow) {
V1760 := __e.Get(1)
_ = V1760
V1761 := __e.Get(2)
_ = V1761
tmp12133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1760)
}
__typedArg0 := V1760
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12133 {
tmp12121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1760)
}
__typedArg0 := V1760
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12122 := Call(__e, PrimFunc(symshen_4deref), tmp12121, V1761)


tmp12123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1760)
}
__typedArg0 := V1760
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12124 := Call(__e, PrimFunc(symshen_4deref), tmp12123, V1761)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12122, tmp12124)
}
__typedArg0 := tmp12122
__typedArg1 := tmp12124
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp12131 := Call(__e, PrimFunc(symshen_4pvar_2), V1760)


if True == tmp12131 {
tmp12125 := MakeNative(func(__e *ControlFlow) {
W1762 := __e.Get(1)
_ = W1762
tmp12127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W1762, symshen_4_1null_1)
}
__typedArg0 := W1762
__typedArg1 := symshen_4_1null_1
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12127 {
__e.Return(V1760)
return
} else {
__e.TailApply(PrimFunc(symshen_4deref), W1762, V1761)
return
}


}, 1)

tmp12128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1760, MakeNumber(1))
}
__typedArg0 := V1760
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp12129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1761, tmp12128)
}
__typedArg0 := V1761
__typedArg1 := tmp12128
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp12125, tmp12129)
return


} else {
__e.Return(V1760)
return
}


}


}, 2)

tmp12134 := Call(__e, ns2_1set, symshen_4deref, tmp12120)


_ = tmp12134

tmp12135 := MakeNative(func(__e *ControlFlow) {
V1763 := __e.Get(1)
_ = V1763
V1764 := __e.Get(2)
_ = V1764
V1765 := __e.Get(3)
_ = V1765
V1766 := __e.Get(4)
_ = V1766
tmp12136 := MakeNative(func(__e *ControlFlow) {
W1767 := __e.Get(1)
_ = W1767
tmp12137 := MakeNative(func(__e *ControlFlow) {
W1768 := __e.Get(1)
_ = W1768
tmp12139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W1768, False)
}
__typedArg0 := W1768
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12139 {
__e.TailApply(PrimFunc(symshen_4unwind), V1763, V1765, W1768)
return
} else {
__e.Return(W1768)
return
}


}, 1)

tmp12140 := Call(__e, PrimFunc(symthaw), V1766)


__e.TailApply(tmp12137, tmp12140)
return


}, 1)

tmp12141 := Call(__e, PrimFunc(symshen_4bindv), V1763, V1764, V1765)


__e.TailApply(tmp12136, tmp12141)
return


}, 4)

tmp12142 := Call(__e, ns2_1set, symshen_4bind_b, tmp12135)


_ = tmp12142

tmp12143 := MakeNative(func(__e *ControlFlow) {
V1769 := __e.Get(1)
_ = V1769
V1770 := __e.Get(2)
_ = V1770
V1771 := __e.Get(3)
_ = V1771
tmp12144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1769, MakeNumber(1))
}
__typedArg0 := V1769
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V1771, tmp12144, V1770)
}
__typedArg0 := V1771
__typedArg1 := tmp12144
__typedArg2 := V1770
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return


}, 3)

tmp12145 := Call(__e, ns2_1set, symshen_4bindv, tmp12143)


_ = tmp12145

tmp12146 := MakeNative(func(__e *ControlFlow) {
V1772 := __e.Get(1)
_ = V1772
V1773 := __e.Get(2)
_ = V1773
V1774 := __e.Get(3)
_ = V1774
tmp12147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1772, MakeNumber(1))
}
__typedArg0 := V1772
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp12148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V1773, tmp12147, symshen_4_1null_1)
}
__typedArg0 := V1773
__typedArg1 := tmp12147
__typedArg2 := symshen_4_1null_1
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

_ = tmp12148

__e.Return(V1774)
return


}, 3)

tmp12149 := Call(__e, ns2_1set, symshen_4unwind, tmp12146)


_ = tmp12149

tmp12150 := MakeNative(func(__e *ControlFlow) {
V1783 := __e.Get(1)
_ = V1783
V1784 := __e.Get(2)
_ = V1784
V1785 := __e.Get(3)
_ = V1785
tmp12165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1783)
}
__typedArg0 := Nil
__typedArg1 := V1783
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12165 {
__e.Return(V1784)
return
} else {
tmp12163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1783)
}
__typedArg0 := V1783
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12163 {
tmp12151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1783)
}
__typedArg0 := V1783
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1785, Nil)
}
__typedArg0 := V1785
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4newpv, tmp12152)
}
__typedArg0 := symshen_4newpv
__typedArg1 := tmp12152
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1783)
}
__typedArg0 := V1783
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12155 := Call(__e, PrimFunc(symshen_4stpart), tmp12154, V1784, V1785)


tmp12156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12155, Nil)
}
__typedArg0 := tmp12155
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1785, tmp12156)
}
__typedArg0 := V1785
__typedArg1 := tmp12156
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4gc, tmp12157)
}
__typedArg0 := symshen_4gc
__typedArg1 := tmp12157
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12158, Nil)
}
__typedArg0 := tmp12158
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12153, tmp12159)
}
__typedArg0 := tmp12153
__typedArg1 := tmp12159
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12151, tmp12160)
}
__typedArg0 := tmp12151
__typedArg1 := tmp12160
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp12161)
}
__typedArg0 := symlet
__typedArg1 := tmp12161
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.stpart"))
}
__typedArg0 := MakeString("implementation error in shen.stpart")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 3)

tmp12166 := Call(__e, ns2_1set, symshen_4stpart, tmp12150)


_ = tmp12166

tmp12167 := MakeNative(func(__e *ControlFlow) {
V1786 := __e.Get(1)
_ = V1786
V1787 := __e.Get(2)
_ = V1787
tmp12172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1787, False)
}
__typedArg0 := V1787
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12172 {
tmp12168 := MakeNative(func(__e *ControlFlow) {
W1788 := __e.Get(1)
_ = W1788
tmp12169 := Call(__e, PrimFunc(symshen_4decrement_1ticket), W1788, V1786)


_ = tmp12169

__e.Return(V1787)
return


}, 1)

tmp12170 := Call(__e, PrimFunc(symshen_4ticket_1number), V1786)


__e.TailApply(tmp12168, tmp12170)
return


} else {
__e.Return(V1787)
return
}


}, 2)

tmp12173 := Call(__e, ns2_1set, symshen_4gc, tmp12167)


_ = tmp12173

tmp12174 := MakeNative(func(__e *ControlFlow) {
V1789 := __e.Get(1)
_ = V1789
V1790 := __e.Get(2)
_ = V1790
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V1790, MakeNumber(1), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V1789)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V1789
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())
}
__typedArg0 := V1790
__typedArg1 := MakeNumber(1)
__typedArg2 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V1789)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V1789
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return


}, 2)

tmp12176 := Call(__e, ns2_1set, symshen_4decrement_1ticket, tmp12174)


_ = tmp12176

tmp12177 := MakeNative(func(__e *ControlFlow) {
V1791 := __e.Get(1)
_ = V1791
tmp12178 := MakeNative(func(__e *ControlFlow) {
W1792 := __e.Get(1)
_ = W1792
tmp12179 := MakeNative(func(__e *ControlFlow) {
W1793 := __e.Get(1)
_ = W1793
tmp12180 := MakeNative(func(__e *ControlFlow) {
W1794 := __e.Get(1)
_ = W1794
__e.Return(W1793)
return
}, 1)

tmp12181 := Call(__e, PrimFunc(symshen_4nextticket), V1791, W1792)


__e.TailApply(tmp12180, tmp12181)
return


}, 1)

tmp12182 := Call(__e, PrimFunc(symshen_4make_1prolog_1variable), W1792)


__e.TailApply(tmp12179, tmp12182)
return


}, 1)

tmp12183 := Call(__e, PrimFunc(symshen_4ticket_1number), V1791)


__e.TailApply(tmp12178, tmp12183)
return


}, 1)

tmp12184 := Call(__e, ns2_1set, symshen_4newpv, tmp12177)


_ = tmp12184

tmp12185 := MakeNative(func(__e *ControlFlow) {
V1795 := __e.Get(1)
_ = V1795
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1795, MakeNumber(1))
}
__typedArg0 := V1795
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp12186 := Call(__e, ns2_1set, symshen_4ticket_1number, tmp12185)


_ = tmp12186

tmp12187 := MakeNative(func(__e *ControlFlow) {
V1796 := __e.Get(1)
_ = V1796
V1797 := __e.Get(2)
_ = V1797
tmp12188 := MakeNative(func(__e *ControlFlow) {
W1798 := __e.Get(1)
_ = W1798
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(W1798, MakeNumber(1), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V1797)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V1797
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := W1798
__typedArg1 := MakeNumber(1)
__typedArg2 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V1797)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V1797
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return


}, 1)

tmp12190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V1796, V1797, symshen_4_1null_1)
}
__typedArg0 := V1796
__typedArg1 := V1797
__typedArg2 := symshen_4_1null_1
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

__e.TailApply(tmp12188, tmp12190)
return


}, 2)

tmp12191 := Call(__e, ns2_1set, symshen_4nextticket, tmp12187)


_ = tmp12191

tmp12192 := MakeNative(func(__e *ControlFlow) {
V1799 := __e.Get(1)
_ = V1799
tmp12193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(MakeNumber(2))
}
__typedArg0 := MakeNumber(2)
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

tmp12194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(tmp12193, MakeNumber(0), symshen_4pvar)
}
__typedArg0 := tmp12193
__typedArg1 := MakeNumber(0)
__typedArg2 := symshen_4pvar
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(tmp12194, MakeNumber(1), V1799)
}
__typedArg0 := tmp12194
__typedArg1 := MakeNumber(1)
__typedArg2 := V1799
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return


}, 1)

tmp12195 := Call(__e, ns2_1set, symshen_4make_1prolog_1variable, tmp12192)


_ = tmp12195

tmp12196 := MakeNative(func(__e *ControlFlow) {
V1800 := __e.Get(1)
_ = V1800
tmp12197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1800, MakeNumber(1))
}
__typedArg0 := V1800
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp12198 := Call(__e, PrimFunc(symshen_4app), tmp12197, MakeString(""), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Var"))
__typedS1, __typedOK1 := TypedString(tmp12198)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Var")
__typedArg1 := tmp12198
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp12199 := Call(__e, ns2_1set, symshen_4pvar, tmp12196)


_ = tmp12199

tmp12200 := MakeNative(func(__e *ControlFlow) {
tmp12201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dinfs_d)
}
__typedArg0 := symshen_4_dinfs_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dinfs_d, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp12201)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp12201
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := symshen_4_dinfs_d
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp12201)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeNumber(1)
__typedArg1 := tmp12201
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp12203 := Call(__e, ns2_1set, symshen_4incinfs, tmp12200)


_ = tmp12203

tmp12204 := MakeNative(func(__e *ControlFlow) {
V1801 := __e.Get(1)
_ = V1801
tmp12211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V1801)
}
__typedArg0 := V1801
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})()

var ifres12208 Obj

if True == tmp12211 {
tmp12210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(V1801)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(0))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := V1801
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})()

var ifres12209 Obj

if True == tmp12210 {
ifres12209 = True


} else {
ifres12209 = False


}

ifres12208 = ifres12209


} else {
ifres12208 = False


}

if True == ifres12208 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsize_1prolog_1vector_d, V1801)
}
__typedArg0 := symshen_4_dsize_1prolog_1vector_d
__typedArg1 := V1801
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
tmp12205 := Call(__e, PrimFunc(symshen_4app), V1801, MakeString(""), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("prolog vector size: size should be a positive integer; not "))
__typedS1, __typedOK1 := TypedString(tmp12205)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("prolog vector size: size should be a positive integer; not ")
__typedArg1 := tmp12205
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("prolog vector size: size should be a positive integer; not "))
__typedS1, __typedOK1 := TypedString(tmp12205)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("prolog vector size: size should be a positive integer; not ")
__typedArg1 := tmp12205
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 1)

tmp12212 := Call(__e, ns2_1set, symshen_4prolog_1vector_1size, tmp12204)


_ = tmp12212

tmp12213 := MakeNative(func(__e *ControlFlow) {
V1813 := __e.Get(1)
_ = V1813
V1814 := __e.Get(2)
_ = V1814
V1815 := __e.Get(3)
_ = V1815
V1816 := __e.Get(4)
_ = V1816
tmp12243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1813, V1814)
}
__typedArg0 := V1813
__typedArg1 := V1814
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12243 {
__e.TailApply(PrimFunc(symthaw), V1816)
return
} else {
tmp12241 := Call(__e, PrimFunc(symshen_4pvar_2), V1813)


var ifres12236 Obj

if True == tmp12241 {
tmp12238 := Call(__e, PrimFunc(symshen_4deref), V1814, V1815)


tmp12239 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V1813, tmp12238)


tmp12240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp12239)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp12239
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres12237 Obj

if True == tmp12240 {
ifres12237 = True


} else {
ifres12237 = False


}

ifres12236 = ifres12237


} else {
ifres12236 = False


}

if True == ifres12236 {
__e.TailApply(PrimFunc(symshen_4bind_b), V1813, V1814, V1815, V1816)
return
} else {
tmp12234 := Call(__e, PrimFunc(symshen_4pvar_2), V1814)


var ifres12229 Obj

if True == tmp12234 {
tmp12231 := Call(__e, PrimFunc(symshen_4deref), V1813, V1815)


tmp12232 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V1814, tmp12231)


tmp12233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp12232)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp12232
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres12230 Obj

if True == tmp12233 {
ifres12230 = True


} else {
ifres12230 = False


}

ifres12229 = ifres12230


} else {
ifres12229 = False


}

if True == ifres12229 {
__e.TailApply(PrimFunc(symshen_4bind_b), V1814, V1813, V1815, V1816)
return
} else {
tmp12227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1813)
}
__typedArg0 := V1813
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12224 Obj

if True == tmp12227 {
tmp12226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1814)
}
__typedArg0 := V1814
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12225 Obj

if True == tmp12226 {
ifres12225 = True


} else {
ifres12225 = False


}

ifres12224 = ifres12225


} else {
ifres12224 = False


}

if True == ifres12224 {
tmp12214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1813)
}
__typedArg0 := V1813
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12215 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12214, V1815)


tmp12216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1814)
}
__typedArg0 := V1814
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12217 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12216, V1815)


tmp12218 := MakeNative(func(__e *ControlFlow) {
tmp12219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1813)
}
__typedArg0 := V1813
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12220 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12219, V1815)


tmp12221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1814)
}
__typedArg0 := V1814
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12222 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12221, V1815)


__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp12220, tmp12222, V1815, V1816)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp12215, tmp12217, V1815, tmp12218)
return


} else {
__e.Return(False)
return
}


}


}


}


}, 4)

tmp12244 := Call(__e, ns2_1set, symshen_4lzy_a_b, tmp12213)


_ = tmp12244

tmp12245 := MakeNative(func(__e *ControlFlow) {
V1828 := __e.Get(1)
_ = V1828
V1829 := __e.Get(2)
_ = V1829
V1830 := __e.Get(3)
_ = V1830
V1831 := __e.Get(4)
_ = V1831
tmp12265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1828, V1829)
}
__typedArg0 := V1828
__typedArg1 := V1829
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12265 {
__e.TailApply(PrimFunc(symthaw), V1831)
return
} else {
tmp12263 := Call(__e, PrimFunc(symshen_4pvar_2), V1828)


if True == tmp12263 {
__e.TailApply(PrimFunc(symshen_4bind_b), V1828, V1829, V1830, V1831)
return
} else {
tmp12261 := Call(__e, PrimFunc(symshen_4pvar_2), V1829)


if True == tmp12261 {
__e.TailApply(PrimFunc(symshen_4bind_b), V1829, V1828, V1830, V1831)
return
} else {
tmp12259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1828)
}
__typedArg0 := V1828
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12256 Obj

if True == tmp12259 {
tmp12258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1829)
}
__typedArg0 := V1829
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12257 Obj

if True == tmp12258 {
ifres12257 = True


} else {
ifres12257 = False


}

ifres12256 = ifres12257


} else {
ifres12256 = False


}

if True == ifres12256 {
tmp12246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1828)
}
__typedArg0 := V1828
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12247 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12246, V1830)


tmp12248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1829)
}
__typedArg0 := V1829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12249 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12248, V1830)


tmp12250 := MakeNative(func(__e *ControlFlow) {
tmp12251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1828)
}
__typedArg0 := V1828
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12252 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12251, V1830)


tmp12253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1829)
}
__typedArg0 := V1829
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12254 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12253, V1830)


__e.TailApply(PrimFunc(symshen_4lzy_a), tmp12252, tmp12254, V1830, V1831)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4lzy_a), tmp12247, tmp12249, V1830, tmp12250)
return


} else {
__e.Return(False)
return
}


}


}


}


}, 4)

tmp12266 := Call(__e, ns2_1set, symshen_4lzy_a, tmp12245)


_ = tmp12266

tmp12267 := MakeNative(func(__e *ControlFlow) {
V1837 := __e.Get(1)
_ = V1837
V1838 := __e.Get(2)
_ = V1838
tmp12277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1837, V1838)
}
__typedArg0 := V1837
__typedArg1 := V1838
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12277 {
__e.Return(True)
return
} else {
tmp12275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1838)
}
__typedArg0 := V1838
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12275 {
tmp12272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1838)
}
__typedArg0 := V1838
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12273 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V1837, tmp12272)


if True == tmp12273 {
__e.Return(True)
return
} else {
tmp12269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1838)
}
__typedArg0 := V1838
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12270 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V1837, tmp12269)


if True == tmp12270 {
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

tmp12278 := Call(__e, ns2_1set, symshen_4occurs_1check_2, tmp12267)


_ = tmp12278

tmp12279 := MakeNative(func(__e *ControlFlow) {
V1839 := __e.Get(1)
_ = V1839
V1840 := __e.Get(2)
_ = V1840
V1841 := __e.Get(3)
_ = V1841
V1842 := __e.Get(4)
_ = V1842
V1843 := __e.Get(5)
_ = V1843
tmp12280 := Call(__e, V1839, V1840)


tmp12281 := Call(__e, tmp12280, V1841)


tmp12282 := Call(__e, tmp12281, V1842)


__e.TailApply(tmp12282, V1843)
return


}, 5)

tmp12283 := Call(__e, ns2_1set, symcall, tmp12279)


_ = tmp12283

tmp12284 := MakeNative(func(__e *ControlFlow) {
V1850 := __e.Get(1)
_ = V1850
V1851 := __e.Get(2)
_ = V1851
V1852 := __e.Get(3)
_ = V1852
V1853 := __e.Get(4)
_ = V1853
V1854 := __e.Get(5)
_ = V1854
__e.TailApply(PrimFunc(symshen_4deref), V1850, V1851)
return
}, 5)

tmp12285 := Call(__e, ns2_1set, symreturn, tmp12284)


_ = tmp12285

tmp12286 := MakeNative(func(__e *ControlFlow) {
V1861 := __e.Get(1)
_ = V1861
V1862 := __e.Get(2)
_ = V1862
V1863 := __e.Get(3)
_ = V1863
V1864 := __e.Get(4)
_ = V1864
V1865 := __e.Get(5)
_ = V1865
if True == V1861 {
__e.TailApply(PrimFunc(symthaw), V1865)
return
} else {
__e.Return(False)
return
}
}, 5)

tmp12288 := Call(__e, ns2_1set, symwhen, tmp12286)


_ = tmp12288

tmp12289 := MakeNative(func(__e *ControlFlow) {
V1866 := __e.Get(1)
_ = V1866
V1867 := __e.Get(2)
_ = V1867
V1868 := __e.Get(3)
_ = V1868
V1869 := __e.Get(4)
_ = V1869
V1870 := __e.Get(5)
_ = V1870
V1871 := __e.Get(6)
_ = V1871
tmp12290 := Call(__e, PrimFunc(symshen_4lazyderef), V1866, V1868)


tmp12291 := Call(__e, PrimFunc(symshen_4lazyderef), V1867, V1868)


__e.TailApply(PrimFunc(symshen_4lzy_a), tmp12290, tmp12291, V1868, V1871)
return


}, 6)

tmp12292 := Call(__e, ns2_1set, symis, tmp12289)


_ = tmp12292

tmp12293 := MakeNative(func(__e *ControlFlow) {
V1872 := __e.Get(1)
_ = V1872
V1873 := __e.Get(2)
_ = V1873
V1874 := __e.Get(3)
_ = V1874
V1875 := __e.Get(4)
_ = V1875
V1876 := __e.Get(5)
_ = V1876
V1877 := __e.Get(6)
_ = V1877
tmp12294 := Call(__e, PrimFunc(symshen_4lazyderef), V1872, V1874)


tmp12295 := Call(__e, PrimFunc(symshen_4lazyderef), V1873, V1874)


__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp12294, tmp12295, V1874, V1877)
return


}, 6)

tmp12296 := Call(__e, ns2_1set, symis_b, tmp12293)


_ = tmp12296

tmp12297 := MakeNative(func(__e *ControlFlow) {
V1882 := __e.Get(1)
_ = V1882
V1883 := __e.Get(2)
_ = V1883
V1884 := __e.Get(3)
_ = V1884
V1885 := __e.Get(4)
_ = V1885
V1886 := __e.Get(5)
_ = V1886
V1887 := __e.Get(6)
_ = V1887
__e.TailApply(PrimFunc(symshen_4bind_b), V1882, V1883, V1884, V1887)
return
}, 6)

tmp12298 := Call(__e, ns2_1set, symbind, tmp12297)


_ = tmp12298

tmp12299 := MakeNative(func(__e *ControlFlow) {
V1888 := __e.Get(1)
_ = V1888
V1889 := __e.Get(2)
_ = V1889
V1890 := __e.Get(3)
_ = V1890
V1891 := __e.Get(4)
_ = V1891
V1892 := __e.Get(5)
_ = V1892
tmp12301 := Call(__e, PrimFunc(symshen_4lazyderef), V1888, V1889)


tmp12302 := Call(__e, PrimFunc(symshen_4pvar_2), tmp12301)


if True == tmp12302 {
__e.TailApply(PrimFunc(symthaw), V1892)
return
} else {
__e.Return(False)
return
}


}, 5)

tmp12303 := Call(__e, ns2_1set, symvar_2, tmp12299)


_ = tmp12303

tmp12304 := MakeNative(func(__e *ControlFlow) {
V1895 := __e.Get(1)
_ = V1895
__e.Return(MakeString("|prolog vector|"))
return
}, 1)

tmp12305 := Call(__e, ns2_1set, symshen_4print_1prolog_1vector, tmp12304)


_ = tmp12305

tmp12306 := MakeNative(func(__e *ControlFlow) {
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
tmp12319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1914)
}
__typedArg0 := Nil
__typedArg1 := V1914
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12319 {
__e.Return(False)
return
} else {
tmp12317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1914)
}
__typedArg0 := V1914
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12317 {
tmp12307 := MakeNative(func(__e *ControlFlow) {
W1919 := __e.Get(1)
_ = W1919
tmp12310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W1919, False)
}
__typedArg0 := W1919
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12310 {
tmp12308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1914)
}
__typedArg0 := V1914
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symfork), tmp12308, V1915, V1916, V1917, V1918)
return


} else {
__e.Return(W1919)
return
}


}, 1)

tmp12311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1914)
}
__typedArg0 := V1914
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12312 := Call(__e, tmp12311, V1915)


tmp12313 := Call(__e, tmp12312, V1916)


tmp12314 := Call(__e, tmp12313, V1917)


tmp12315 := Call(__e, tmp12314, V1918)


__e.TailApply(tmp12307, tmp12315)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("fork expects a list of literals\n"))
}
__typedArg0 := MakeString("fork expects a list of literals\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 5)

tmp12320 := Call(__e, ns2_1set, symfork, tmp12306)


_ = tmp12320

tmp12321 := MakeNative(func(__e *ControlFlow) {
V1920 := __e.Get(1)
_ = V1920
V1921 := __e.Get(2)
_ = V1921
V1922 := __e.Get(3)
_ = V1922
V1923 := __e.Get(4)
_ = V1923
V1924 := __e.Get(5)
_ = V1924
V1925 := __e.Get(6)
_ = V1925
V1926 := __e.Get(7)
_ = V1926
tmp12328 := Call(__e, PrimFunc(symshen_4unlocked_2), V1924)


if True == tmp12328 {
tmp12322 := MakeNative(func(__e *ControlFlow) {
W1927 := __e.Get(1)
_ = W1927
tmp12323 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12323

tmp12324 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4findall_1h), V1920, V1921, V1922, W1927, V1923, V1924, V1925, V1926)
return
}, 0)

tmp12325 := Call(__e, PrimFunc(symis), W1927, Nil, V1923, V1924, V1925, tmp12324)


__e.TailApply(PrimFunc(symshen_4gc), V1923, tmp12325)
return


}, 1)

tmp12326 := Call(__e, PrimFunc(symshen_4newpv), V1923)


__e.TailApply(tmp12322, tmp12326)
return


} else {
__e.Return(False)
return
}


}, 7)

tmp12329 := Call(__e, ns2_1set, symfindall, tmp12321)


_ = tmp12329

tmp12330 := MakeNative(func(__e *ControlFlow) {
V1928 := __e.Get(1)
_ = V1928
V1929 := __e.Get(2)
_ = V1929
V1930 := __e.Get(3)
_ = V1930
V1931 := __e.Get(4)
_ = V1931
V1932 := __e.Get(5)
_ = V1932
V1933 := __e.Get(6)
_ = V1933
V1934 := __e.Get(7)
_ = V1934
V1935 := __e.Get(8)
_ = V1935
tmp12331 := MakeNative(func(__e *ControlFlow) {
W1936 := __e.Get(1)
_ = W1936
tmp12336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W1936, False)
}
__typedArg0 := W1936
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12336 {
tmp12334 := Call(__e, PrimFunc(symshen_4unlocked_2), V1933)


if True == tmp12334 {
tmp12332 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12332

__e.TailApply(PrimFunc(symis_b), V1930, V1931, V1932, V1933, V1934, V1935)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W1936)
return
}


}, 1)

tmp12341 := Call(__e, PrimFunc(symshen_4unlocked_2), V1933)


var ifres12337 Obj

if True == tmp12341 {
tmp12338 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12338

tmp12339 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4overbind), V1928, V1931, V1932, V1933, V1934, V1935)
return
}, 0)

tmp12340 := Call(__e, PrimFunc(symcall), V1929, V1932, V1933, V1934, tmp12339)


ifres12337 = tmp12340


} else {
ifres12337 = False


}

__e.TailApply(tmp12331, ifres12337)
return


}, 8)

tmp12342 := Call(__e, ns2_1set, symshen_4findall_1h, tmp12330)


_ = tmp12342

tmp12343 := MakeNative(func(__e *ControlFlow) {
V1943 := __e.Get(1)
_ = V1943
V1944 := __e.Get(2)
_ = V1944
V1945 := __e.Get(3)
_ = V1945
V1946 := __e.Get(4)
_ = V1946
V1947 := __e.Get(5)
_ = V1947
V1948 := __e.Get(6)
_ = V1948
tmp12344 := Call(__e, PrimFunc(symshen_4deref), V1943, V1945)


tmp12345 := Call(__e, PrimFunc(symshen_4lazyderef), V1944, V1945)


tmp12346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12344, tmp12345)
}
__typedArg0 := tmp12344
__typedArg1 := tmp12345
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12347 := Call(__e, PrimFunc(symshen_4bindv), V1944, tmp12346, V1945)


_ = tmp12347

__e.Return(False)
return


}, 6)

tmp12348 := Call(__e, ns2_1set, symshen_4overbind, tmp12343)


_ = tmp12348

tmp12349 := MakeNative(func(__e *ControlFlow) {
V1951 := __e.Get(1)
_ = V1951
tmp12353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V1951)
}
__typedArg0 := sym_7
__typedArg1 := V1951
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12353 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_doccurs_d, True)
}
__typedArg0 := symshen_4_doccurs_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
tmp12351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V1951)
}
__typedArg0 := sym_1
__typedArg1 := V1951
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12351 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_doccurs_d, False)
}
__typedArg0 := symshen_4_doccurs_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("occurs-check expects a + or a -.\n"))
}
__typedArg0 := MakeString("occurs-check expects a + or a -.\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

__e.TailApply(ns2_1set, symoccurs_1check, tmp12349)
return




}, 0)

