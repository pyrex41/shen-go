package main

import . "github.com/tiancaiamao/shen-go/kl"

var LoadMain = MakeNative(func(__e *ControlFlow) {
tmp7989 := MakeNative(func(__e *ControlFlow) {
V1239 := __e.Get(1)
_ = V1239
tmp7990 := MakeNative(func(__e *ControlFlow) {
W1240 := __e.Get(1)
_ = W1240
tmp7991 := MakeNative(func(__e *ControlFlow) {
W1241 := __e.Get(1)
_ = W1241
tmp7992 := MakeNative(func(__e *ControlFlow) {
W1247 := __e.Get(1)
_ = W1247
__e.Return(symloaded)
return
}, 1)

var ifres7993 Obj

if True == W1240 {
tmp7994 := Call(__e, PrimFunc(syminferences))


tmp7995 := Call(__e, PrimFunc(symshen_4app), tmp7994, MakeString(" inferences\n"), symshen_4a)


tmp7996 := PrimStringConcat(MakeString("\ntypechecked in "), tmp7995)

tmp7997 := Call(__e, PrimFunc(symstoutput))


tmp7998 := Call(__e, PrimFunc(sympr), tmp7996, tmp7997)


ifres7993 = tmp7998


} else {
ifres7993 = symshen_4skip


}

__e.TailApply(tmp7992, ifres7993)
return


}, 1)

tmp7999 := MakeNative(func(__e *ControlFlow) {
W1242 := __e.Get(1)
_ = W1242
tmp8000 := MakeNative(func(__e *ControlFlow) {
W1243 := __e.Get(1)
_ = W1243
tmp8001 := MakeNative(func(__e *ControlFlow) {
W1244 := __e.Get(1)
_ = W1244
tmp8002 := MakeNative(func(__e *ControlFlow) {
W1245 := __e.Get(1)
_ = W1245
tmp8003 := MakeNative(func(__e *ControlFlow) {
W1246 := __e.Get(1)
_ = W1246
__e.Return(W1243)
return
}, 1)

tmp8004 := PrimStr(W1245)

tmp8005 := PrimStringConcat(tmp8004, MakeString(" secs\n"))

tmp8006 := PrimStringConcat(MakeString("\nrun time: "), tmp8005)

tmp8007 := Call(__e, PrimFunc(symstoutput))


tmp8008 := Call(__e, PrimFunc(sympr), tmp8006, tmp8007)


__e.TailApply(tmp8003, tmp8008)
return


}, 1)

tmp8009 := PrimNumberSubtract(W1244, W1242)

__e.TailApply(tmp8002, tmp8009)
return


}, 1)

tmp8010 := PrimGetTime(symrun)

__e.TailApply(tmp8001, tmp8010)
return


}, 1)

tmp8011 := Call(__e, PrimFunc(symread_1file), V1239)


tmp8012 := Call(__e, PrimFunc(symshen_4load_1help), W1240, tmp8011)


__e.TailApply(tmp8000, tmp8012)
return


}, 1)

tmp8013 := PrimGetTime(symrun)

tmp8014 := Call(__e, tmp7999, tmp8013)


__e.TailApply(tmp7991, tmp8014)
return


}, 1)

tmp8015 := PrimValue(symshen_4_dtc_d)

__e.TailApply(tmp7990, tmp8015)
return


}, 1)

tmp8016 := Call(__e, ns2_1set, symload, tmp7989)


_ = tmp8016

tmp8017 := MakeNative(func(__e *ControlFlow) {
V1250 := __e.Get(1)
_ = V1250
V1251 := __e.Get(2)
_ = V1251
tmp8019 := PrimEqual(False, V1250)

if True == tmp8019 {
__e.TailApply(PrimFunc(symshen_4eval_1and_1print), V1251)
return
} else {
__e.TailApply(PrimFunc(symshen_4check_1eval_1and_1print), V1251)
return
}


}, 2)

tmp8020 := Call(__e, ns2_1set, symshen_4load_1help, tmp8017)


_ = tmp8020

tmp8021 := MakeNative(func(__e *ControlFlow) {
V1252 := __e.Get(1)
_ = V1252
tmp8022 := MakeNative(func(__e *ControlFlow) {
Z1253 := __e.Get(1)
_ = Z1253
tmp8023 := Call(__e, PrimFunc(symshen_4shen_1_6kl), Z1253)


tmp8024 := Call(__e, PrimFunc(symeval_1kl), tmp8023)


tmp8025 := Call(__e, PrimFunc(symshen_4app), tmp8024, MakeString("\n"), symshen_4s)


tmp8026 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8025, tmp8026)
return


}, 1)

__e.TailApply(PrimFunc(symshen_4for_1each), tmp8022, V1252)
return


}, 1)

tmp8027 := Call(__e, ns2_1set, symshen_4eval_1and_1print, tmp8021)


_ = tmp8027

tmp8028 := MakeNative(func(__e *ControlFlow) {
V1254 := __e.Get(1)
_ = V1254
tmp8029 := MakeNative(func(__e *ControlFlow) {
W1255 := __e.Get(1)
_ = W1255
tmp8030 := MakeNative(func(__e *ControlFlow) {
W1257 := __e.Get(1)
_ = W1257
tmp8031 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4work_1through), V1254)
return
}, 0)

tmp8032 := MakeNative(func(__e *ControlFlow) {
Z1259 := __e.Get(1)
_ = Z1259
__e.TailApply(PrimFunc(symshen_4unwind_1types), Z1259, W1255)
return
}, 1)

__e.TailApply(try_1catch, tmp8031, tmp8032)
return


}, 1)

tmp8033 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4assumetypes), W1255)
return
}, 0)

tmp8034 := MakeNative(func(__e *ControlFlow) {
Z1258 := __e.Get(1)
_ = Z1258
__e.TailApply(PrimFunc(symshen_4unwind_1types), Z1258, W1255)
return
}, 1)

tmp8035 := Call(__e, try_1catch, tmp8033, tmp8034)


__e.TailApply(tmp8030, tmp8035)
return


}, 1)

tmp8036 := MakeNative(func(__e *ControlFlow) {
Z1256 := __e.Get(1)
_ = Z1256
__e.TailApply(PrimFunc(symshen_4typetable), Z1256)
return
}, 1)

tmp8037 := Call(__e, PrimFunc(symmapcan), tmp8036, V1254)


__e.TailApply(tmp8029, tmp8037)
return


}, 1)

tmp8038 := Call(__e, ns2_1set, symshen_4check_1eval_1and_1print, tmp8028)


_ = tmp8038

tmp8039 := MakeNative(func(__e *ControlFlow) {
V1264 := __e.Get(1)
_ = V1264
tmp8084 := PrimIsPair(V1264)

var ifres8065 Obj

if True == tmp8084 {
tmp8082 := PrimHead(V1264)

tmp8083 := PrimEqual(symdefine, tmp8082)

var ifres8067 Obj

if True == tmp8083 {
tmp8080 := PrimTail(V1264)

tmp8081 := PrimIsPair(tmp8080)

var ifres8069 Obj

if True == tmp8081 {
tmp8077 := PrimTail(V1264)

tmp8078 := PrimTail(tmp8077)

tmp8079 := PrimIsPair(tmp8078)

var ifres8071 Obj

if True == tmp8079 {
tmp8073 := PrimTail(V1264)

tmp8074 := PrimTail(tmp8073)

tmp8075 := PrimHead(tmp8074)

tmp8076 := PrimEqual(sym_i, tmp8075)

var ifres8072 Obj

if True == tmp8076 {
ifres8072 = True


} else {
ifres8072 = False


}

ifres8071 = ifres8072


} else {
ifres8071 = False


}

var ifres8070 Obj

if True == ifres8071 {
ifres8070 = True


} else {
ifres8070 = False


}

ifres8069 = ifres8070


} else {
ifres8069 = False


}

var ifres8068 Obj

if True == ifres8069 {
ifres8068 = True


} else {
ifres8068 = False


}

ifres8067 = ifres8068


} else {
ifres8067 = False


}

var ifres8066 Obj

if True == ifres8067 {
ifres8066 = True


} else {
ifres8066 = False


}

ifres8065 = ifres8066


} else {
ifres8065 = False


}

if True == ifres8065 {
tmp8040 := PrimTail(V1264)

tmp8041 := PrimHead(tmp8040)

tmp8042 := PrimTail(V1264)

tmp8043 := PrimHead(tmp8042)

tmp8044 := PrimTail(V1264)

tmp8045 := PrimTail(tmp8044)

tmp8046 := PrimTail(tmp8045)

tmp8047 := Call(__e, PrimFunc(symshen_4type_1F), tmp8043, tmp8046)


tmp8048 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp8047)


tmp8049 := PrimCons(tmp8048, Nil)

__e.Return(PrimCons(tmp8041, tmp8049))
return


} else {
tmp8063 := PrimIsPair(V1264)

var ifres8055 Obj

if True == tmp8063 {
tmp8061 := PrimHead(V1264)

tmp8062 := PrimEqual(symdefine, tmp8061)

var ifres8057 Obj

if True == tmp8062 {
tmp8059 := PrimTail(V1264)

tmp8060 := PrimIsPair(tmp8059)

var ifres8058 Obj

if True == tmp8060 {
ifres8058 = True


} else {
ifres8058 = False


}

ifres8057 = ifres8058


} else {
ifres8057 = False


}

var ifres8056 Obj

if True == ifres8057 {
ifres8056 = True


} else {
ifres8056 = False


}

ifres8055 = ifres8056


} else {
ifres8055 = False


}

if True == ifres8055 {
tmp8050 := PrimTail(V1264)

tmp8051 := PrimHead(tmp8050)

tmp8052 := Call(__e, PrimFunc(symshen_4app), tmp8051, MakeString("\n"), symshen_4a)


tmp8053 := PrimStringConcat(MakeString("missing { in "), tmp8052)

__e.Return(PrimSimpleError(tmp8053))
return


} else {
__e.Return(Nil)
return
}


}


}, 1)

tmp8085 := Call(__e, ns2_1set, symshen_4typetable, tmp8039)


_ = tmp8085

tmp8086 := MakeNative(func(__e *ControlFlow) {
V1271 := __e.Get(1)
_ = V1271
V1272 := __e.Get(2)
_ = V1272
tmp8099 := PrimIsPair(V1272)

var ifres8095 Obj

if True == tmp8099 {
tmp8097 := PrimHead(V1272)

tmp8098 := PrimEqual(sym_j, tmp8097)

var ifres8096 Obj

if True == tmp8098 {
ifres8096 = True


} else {
ifres8096 = False


}

ifres8095 = ifres8096


} else {
ifres8095 = False


}

if True == ifres8095 {
__e.Return(Nil)
return
} else {
tmp8093 := PrimIsPair(V1272)

if True == tmp8093 {
tmp8087 := PrimHead(V1272)

tmp8088 := PrimTail(V1272)

tmp8089 := Call(__e, PrimFunc(symshen_4type_1F), V1271, tmp8088)


__e.Return(PrimCons(tmp8087, tmp8089))
return


} else {
tmp8090 := Call(__e, PrimFunc(symshen_4app), V1271, MakeString("\n"), symshen_4a)


tmp8091 := PrimStringConcat(MakeString("missing } in "), tmp8090)

__e.Return(PrimSimpleError(tmp8091))
return


}


}


}, 2)

tmp8100 := Call(__e, ns2_1set, symshen_4type_1F, tmp8086)


_ = tmp8100

tmp8101 := MakeNative(func(__e *ControlFlow) {
V1275 := __e.Get(1)
_ = V1275
tmp8115 := PrimEqual(Nil, V1275)

if True == tmp8115 {
__e.Return(Nil)
return
} else {
tmp8113 := PrimIsPair(V1275)

var ifres8109 Obj

if True == tmp8113 {
tmp8111 := PrimTail(V1275)

tmp8112 := PrimIsPair(tmp8111)

var ifres8110 Obj

if True == tmp8112 {
ifres8110 = True


} else {
ifres8110 = False


}

ifres8109 = ifres8110


} else {
ifres8109 = False


}

if True == ifres8109 {
tmp8102 := PrimHead(V1275)

tmp8103 := PrimTail(V1275)

tmp8104 := PrimHead(tmp8103)

tmp8105 := Call(__e, PrimFunc(symdeclare), tmp8102, tmp8104)


_ = tmp8105

tmp8106 := PrimTail(V1275)

tmp8107 := PrimTail(tmp8106)

__e.TailApply(PrimFunc(symshen_4assumetypes), tmp8107)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.assumetype")))
return
}


}


}, 1)

tmp8116 := Call(__e, ns2_1set, symshen_4assumetypes, tmp8101)


_ = tmp8116

tmp8117 := MakeNative(func(__e *ControlFlow) {
V1280 := __e.Get(1)
_ = V1280
V1281 := __e.Get(2)
_ = V1281
tmp8128 := PrimIsPair(V1281)

var ifres8124 Obj

if True == tmp8128 {
tmp8126 := PrimTail(V1281)

tmp8127 := PrimIsPair(tmp8126)

var ifres8125 Obj

if True == tmp8127 {
ifres8125 = True


} else {
ifres8125 = False


}

ifres8124 = ifres8125


} else {
ifres8124 = False


}

if True == ifres8124 {
tmp8118 := PrimHead(V1281)

tmp8119 := Call(__e, PrimFunc(symdestroy), tmp8118)


_ = tmp8119

tmp8120 := PrimTail(V1281)

tmp8121 := PrimTail(tmp8120)

__e.TailApply(PrimFunc(symshen_4unwind_1types), V1280, tmp8121)
return


} else {
tmp8122 := PrimErrorToString(V1280)

__e.Return(PrimSimpleError(tmp8122))
return


}


}, 2)

tmp8129 := Call(__e, ns2_1set, symshen_4unwind_1types, tmp8117)


_ = tmp8129

tmp8130 := MakeNative(func(__e *ControlFlow) {
V1284 := __e.Get(1)
_ = V1284
tmp8179 := PrimEqual(Nil, V1284)

if True == tmp8179 {
__e.Return(Nil)
return
} else {
tmp8177 := PrimIsPair(V1284)

var ifres8162 Obj

if True == tmp8177 {
tmp8175 := PrimTail(V1284)

tmp8176 := PrimIsPair(tmp8175)

var ifres8164 Obj

if True == tmp8176 {
tmp8172 := PrimTail(V1284)

tmp8173 := PrimTail(tmp8172)

tmp8174 := PrimIsPair(tmp8173)

var ifres8166 Obj

if True == tmp8174 {
tmp8168 := PrimTail(V1284)

tmp8169 := PrimHead(tmp8168)

tmp8170 := PrimIntern(MakeString(":"))

tmp8171 := PrimEqual(tmp8169, tmp8170)

var ifres8167 Obj

if True == tmp8171 {
ifres8167 = True


} else {
ifres8167 = False


}

ifres8166 = ifres8167


} else {
ifres8166 = False


}

var ifres8165 Obj

if True == ifres8166 {
ifres8165 = True


} else {
ifres8165 = False


}

ifres8164 = ifres8165


} else {
ifres8164 = False


}

var ifres8163 Obj

if True == ifres8164 {
ifres8163 = True


} else {
ifres8163 = False


}

ifres8162 = ifres8163


} else {
ifres8162 = False


}

if True == ifres8162 {
tmp8131 := MakeNative(func(__e *ControlFlow) {
W1285 := __e.Get(1)
_ = W1285
tmp8147 := PrimEqual(W1285, False)

if True == tmp8147 {
__e.TailApply(PrimFunc(symshen_4type_1error))
return
} else {
tmp8132 := MakeNative(func(__e *ControlFlow) {
W1286 := __e.Get(1)
_ = W1286
tmp8133 := MakeNative(func(__e *ControlFlow) {
W1287 := __e.Get(1)
_ = W1287
tmp8134 := PrimTail(V1284)

tmp8135 := PrimTail(tmp8134)

tmp8136 := PrimTail(tmp8135)

__e.TailApply(PrimFunc(symshen_4work_1through), tmp8136)
return


}, 1)

tmp8137 := Call(__e, PrimFunc(symshen_4pretty_1type), W1285)


tmp8138 := Call(__e, PrimFunc(symshen_4app), tmp8137, MakeString("\n"), symshen_4r)


tmp8139 := PrimStringConcat(MakeString(" : "), tmp8138)

tmp8140 := Call(__e, PrimFunc(symshen_4app), W1286, tmp8139, symshen_4s)


tmp8141 := Call(__e, PrimFunc(symstoutput))


tmp8142 := Call(__e, PrimFunc(sympr), tmp8140, tmp8141)


__e.TailApply(tmp8133, tmp8142)
return


}, 1)

tmp8143 := PrimHead(V1284)

tmp8144 := Call(__e, PrimFunc(symshen_4shen_1_6kl), tmp8143)


tmp8145 := Call(__e, PrimFunc(symeval_1kl), tmp8144)


__e.TailApply(tmp8132, tmp8145)
return


}


}, 1)

tmp8148 := PrimHead(V1284)

tmp8149 := PrimTail(V1284)

tmp8150 := PrimTail(tmp8149)

tmp8151 := PrimHead(tmp8150)

tmp8152 := Call(__e, PrimFunc(symshen_4typecheck), tmp8148, tmp8151)


__e.TailApply(tmp8131, tmp8152)
return


} else {
tmp8160 := PrimIsPair(V1284)

if True == tmp8160 {
tmp8153 := PrimHead(V1284)

tmp8154 := PrimIntern(MakeString(":"))

tmp8155 := PrimTail(V1284)

tmp8156 := PrimCons(symA, tmp8155)

tmp8157 := PrimCons(tmp8154, tmp8156)

tmp8158 := PrimCons(tmp8153, tmp8157)

__e.TailApply(PrimFunc(symshen_4work_1through), tmp8158)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.work-through")))
return
}


}


}


}, 1)

tmp8180 := Call(__e, ns2_1set, symshen_4work_1through, tmp8130)


_ = tmp8180

tmp8181 := MakeNative(func(__e *ControlFlow) {
V1289 := __e.Get(1)
_ = V1289
tmp8323 := PrimIsPair(V1289)

var ifres8197 Obj

if True == tmp8323 {
tmp8321 := PrimHead(V1289)

tmp8322 := PrimIsPair(tmp8321)

var ifres8199 Obj

if True == tmp8322 {
tmp8318 := PrimHead(V1289)

tmp8319 := PrimHead(tmp8318)

tmp8320 := PrimEqual(symlist, tmp8319)

var ifres8201 Obj

if True == tmp8320 {
tmp8315 := PrimHead(V1289)

tmp8316 := PrimTail(tmp8315)

tmp8317 := PrimIsPair(tmp8316)

var ifres8203 Obj

if True == tmp8317 {
tmp8311 := PrimHead(V1289)

tmp8312 := PrimTail(tmp8311)

tmp8313 := PrimTail(tmp8312)

tmp8314 := PrimEqual(Nil, tmp8313)

var ifres8205 Obj

if True == tmp8314 {
tmp8309 := PrimTail(V1289)

tmp8310 := PrimIsPair(tmp8309)

var ifres8207 Obj

if True == tmp8310 {
tmp8306 := PrimTail(V1289)

tmp8307 := PrimHead(tmp8306)

tmp8308 := PrimEqual(sym_1_1_6, tmp8307)

var ifres8209 Obj

if True == tmp8308 {
tmp8303 := PrimTail(V1289)

tmp8304 := PrimTail(tmp8303)

tmp8305 := PrimIsPair(tmp8304)

var ifres8211 Obj

if True == tmp8305 {
tmp8299 := PrimTail(V1289)

tmp8300 := PrimTail(tmp8299)

tmp8301 := PrimHead(tmp8300)

tmp8302 := PrimIsPair(tmp8301)

var ifres8213 Obj

if True == tmp8302 {
tmp8294 := PrimTail(V1289)

tmp8295 := PrimTail(tmp8294)

tmp8296 := PrimHead(tmp8295)

tmp8297 := PrimHead(tmp8296)

tmp8298 := PrimEqual(symstr, tmp8297)

var ifres8215 Obj

if True == tmp8298 {
tmp8289 := PrimTail(V1289)

tmp8290 := PrimTail(tmp8289)

tmp8291 := PrimHead(tmp8290)

tmp8292 := PrimTail(tmp8291)

tmp8293 := PrimIsPair(tmp8292)

var ifres8217 Obj

if True == tmp8293 {
tmp8283 := PrimTail(V1289)

tmp8284 := PrimTail(tmp8283)

tmp8285 := PrimHead(tmp8284)

tmp8286 := PrimTail(tmp8285)

tmp8287 := PrimHead(tmp8286)

tmp8288 := PrimIsPair(tmp8287)

var ifres8219 Obj

if True == tmp8288 {
tmp8276 := PrimTail(V1289)

tmp8277 := PrimTail(tmp8276)

tmp8278 := PrimHead(tmp8277)

tmp8279 := PrimTail(tmp8278)

tmp8280 := PrimHead(tmp8279)

tmp8281 := PrimHead(tmp8280)

tmp8282 := PrimEqual(symlist, tmp8281)

var ifres8221 Obj

if True == tmp8282 {
tmp8269 := PrimTail(V1289)

tmp8270 := PrimTail(tmp8269)

tmp8271 := PrimHead(tmp8270)

tmp8272 := PrimTail(tmp8271)

tmp8273 := PrimHead(tmp8272)

tmp8274 := PrimTail(tmp8273)

tmp8275 := PrimIsPair(tmp8274)

var ifres8223 Obj

if True == tmp8275 {
tmp8261 := PrimTail(V1289)

tmp8262 := PrimTail(tmp8261)

tmp8263 := PrimHead(tmp8262)

tmp8264 := PrimTail(tmp8263)

tmp8265 := PrimHead(tmp8264)

tmp8266 := PrimTail(tmp8265)

tmp8267 := PrimTail(tmp8266)

tmp8268 := PrimEqual(Nil, tmp8267)

var ifres8225 Obj

if True == tmp8268 {
tmp8255 := PrimTail(V1289)

tmp8256 := PrimTail(tmp8255)

tmp8257 := PrimHead(tmp8256)

tmp8258 := PrimTail(tmp8257)

tmp8259 := PrimTail(tmp8258)

tmp8260 := PrimIsPair(tmp8259)

var ifres8227 Obj

if True == tmp8260 {
tmp8248 := PrimTail(V1289)

tmp8249 := PrimTail(tmp8248)

tmp8250 := PrimHead(tmp8249)

tmp8251 := PrimTail(tmp8250)

tmp8252 := PrimTail(tmp8251)

tmp8253 := PrimTail(tmp8252)

tmp8254 := PrimEqual(Nil, tmp8253)

var ifres8229 Obj

if True == tmp8254 {
tmp8244 := PrimTail(V1289)

tmp8245 := PrimTail(tmp8244)

tmp8246 := PrimTail(tmp8245)

tmp8247 := PrimEqual(Nil, tmp8246)

var ifres8231 Obj

if True == tmp8247 {
tmp8233 := PrimHead(V1289)

tmp8234 := PrimTail(tmp8233)

tmp8235 := PrimHead(tmp8234)

tmp8236 := PrimTail(V1289)

tmp8237 := PrimTail(tmp8236)

tmp8238 := PrimHead(tmp8237)

tmp8239 := PrimTail(tmp8238)

tmp8240 := PrimHead(tmp8239)

tmp8241 := PrimTail(tmp8240)

tmp8242 := PrimHead(tmp8241)

tmp8243 := PrimEqual(tmp8235, tmp8242)

var ifres8232 Obj

if True == tmp8243 {
ifres8232 = True


} else {
ifres8232 = False


}

ifres8231 = ifres8232


} else {
ifres8231 = False


}

var ifres8230 Obj

if True == ifres8231 {
ifres8230 = True


} else {
ifres8230 = False


}

ifres8229 = ifres8230


} else {
ifres8229 = False


}

var ifres8228 Obj

if True == ifres8229 {
ifres8228 = True


} else {
ifres8228 = False


}

ifres8227 = ifres8228


} else {
ifres8227 = False


}

var ifres8226 Obj

if True == ifres8227 {
ifres8226 = True


} else {
ifres8226 = False


}

ifres8225 = ifres8226


} else {
ifres8225 = False


}

var ifres8224 Obj

if True == ifres8225 {
ifres8224 = True


} else {
ifres8224 = False


}

ifres8223 = ifres8224


} else {
ifres8223 = False


}

var ifres8222 Obj

if True == ifres8223 {
ifres8222 = True


} else {
ifres8222 = False


}

ifres8221 = ifres8222


} else {
ifres8221 = False


}

var ifres8220 Obj

if True == ifres8221 {
ifres8220 = True


} else {
ifres8220 = False


}

ifres8219 = ifres8220


} else {
ifres8219 = False


}

var ifres8218 Obj

if True == ifres8219 {
ifres8218 = True


} else {
ifres8218 = False


}

ifres8217 = ifres8218


} else {
ifres8217 = False


}

var ifres8216 Obj

if True == ifres8217 {
ifres8216 = True


} else {
ifres8216 = False


}

ifres8215 = ifres8216


} else {
ifres8215 = False


}

var ifres8214 Obj

if True == ifres8215 {
ifres8214 = True


} else {
ifres8214 = False


}

ifres8213 = ifres8214


} else {
ifres8213 = False


}

var ifres8212 Obj

if True == ifres8213 {
ifres8212 = True


} else {
ifres8212 = False


}

ifres8211 = ifres8212


} else {
ifres8211 = False


}

var ifres8210 Obj

if True == ifres8211 {
ifres8210 = True


} else {
ifres8210 = False


}

ifres8209 = ifres8210


} else {
ifres8209 = False


}

var ifres8208 Obj

if True == ifres8209 {
ifres8208 = True


} else {
ifres8208 = False


}

ifres8207 = ifres8208


} else {
ifres8207 = False


}

var ifres8206 Obj

if True == ifres8207 {
ifres8206 = True


} else {
ifres8206 = False


}

ifres8205 = ifres8206


} else {
ifres8205 = False


}

var ifres8204 Obj

if True == ifres8205 {
ifres8204 = True


} else {
ifres8204 = False


}

ifres8203 = ifres8204


} else {
ifres8203 = False


}

var ifres8202 Obj

if True == ifres8203 {
ifres8202 = True


} else {
ifres8202 = False


}

ifres8201 = ifres8202


} else {
ifres8201 = False


}

var ifres8200 Obj

if True == ifres8201 {
ifres8200 = True


} else {
ifres8200 = False


}

ifres8199 = ifres8200


} else {
ifres8199 = False


}

var ifres8198 Obj

if True == ifres8199 {
ifres8198 = True


} else {
ifres8198 = False


}

ifres8197 = ifres8198


} else {
ifres8197 = False


}

if True == ifres8197 {
tmp8182 := PrimTail(V1289)

tmp8183 := PrimTail(tmp8182)

tmp8184 := PrimHead(tmp8183)

tmp8185 := PrimTail(tmp8184)

tmp8186 := PrimHead(tmp8185)

tmp8187 := PrimTail(V1289)

tmp8188 := PrimTail(tmp8187)

tmp8189 := PrimHead(tmp8188)

tmp8190 := PrimTail(tmp8189)

tmp8191 := PrimTail(tmp8190)

tmp8192 := PrimCons(sym_a_a_6, tmp8191)

__e.Return(PrimCons(tmp8186, tmp8192))
return


} else {
tmp8195 := PrimIsPair(V1289)

if True == tmp8195 {
tmp8193 := MakeNative(func(__e *ControlFlow) {
Z1290 := __e.Get(1)
_ = Z1290
__e.TailApply(PrimFunc(symshen_4pretty_1type), Z1290)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp8193, V1289)
return


} else {
__e.Return(V1289)
return
}


}


}, 1)

tmp8324 := Call(__e, ns2_1set, symshen_4pretty_1type, tmp8181)


_ = tmp8324

tmp8325 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimSimpleError(MakeString("type error\n")))
return
}, 0)

tmp8326 := Call(__e, ns2_1set, symshen_4type_1error, tmp8325)


_ = tmp8326

tmp8327 := MakeNative(func(__e *ControlFlow) {
V1291 := __e.Get(1)
_ = V1291
tmp8328 := MakeNative(func(__e *ControlFlow) {
W1292 := __e.Get(1)
_ = W1292
tmp8329 := MakeNative(func(__e *ControlFlow) {
W1293 := __e.Get(1)
_ = W1293
tmp8330 := MakeNative(func(__e *ControlFlow) {
W1294 := __e.Get(1)
_ = W1294
tmp8331 := MakeNative(func(__e *ControlFlow) {
W1295 := __e.Get(1)
_ = W1295
tmp8332 := MakeNative(func(__e *ControlFlow) {
W1297 := __e.Get(1)
_ = W1297
__e.Return(W1292)
return
}, 1)

tmp8333 := Call(__e, PrimFunc(symshen_4write_1kl), W1295, W1294)


__e.TailApply(tmp8332, tmp8333)
return


}, 1)

tmp8334 := MakeNative(func(__e *ControlFlow) {
Z1296 := __e.Get(1)
_ = Z1296
tmp8335 := Call(__e, PrimFunc(symshen_4shen_1_6kl_1h), Z1296)


__e.TailApply(PrimFunc(symshen_4partial), tmp8335)
return


}, 1)

tmp8336 := Call(__e, PrimFunc(symmap), tmp8334, W1293)


__e.TailApply(tmp8331, tmp8336)
return


}, 1)

tmp8337 := PrimOpenStream(W1292, symout)

__e.TailApply(tmp8330, tmp8337)
return


}, 1)

tmp8338 := Call(__e, PrimFunc(symread_1file), V1291)


__e.TailApply(tmp8329, tmp8338)
return


}, 1)

tmp8339 := Call(__e, PrimFunc(symshen_4klfile), V1291)


__e.TailApply(tmp8328, tmp8339)
return


}, 1)

tmp8340 := Call(__e, ns2_1set, symbootstrap, tmp8327)


_ = tmp8340

tmp8341 := MakeNative(func(__e *ControlFlow) {
V1298 := __e.Get(1)
_ = V1298
tmp8364 := PrimIsPair(V1298)

var ifres8351 Obj

if True == tmp8364 {
tmp8362 := PrimHead(V1298)

tmp8363 := PrimEqual(symshen_4f_1error, tmp8362)

var ifres8353 Obj

if True == tmp8363 {
tmp8360 := PrimTail(V1298)

tmp8361 := PrimIsPair(tmp8360)

var ifres8355 Obj

if True == tmp8361 {
tmp8357 := PrimTail(V1298)

tmp8358 := PrimTail(tmp8357)

tmp8359 := PrimEqual(Nil, tmp8358)

var ifres8356 Obj

if True == tmp8359 {
ifres8356 = True


} else {
ifres8356 = False


}

ifres8355 = ifres8356


} else {
ifres8355 = False


}

var ifres8354 Obj

if True == ifres8355 {
ifres8354 = True


} else {
ifres8354 = False


}

ifres8353 = ifres8354


} else {
ifres8353 = False


}

var ifres8352 Obj

if True == ifres8353 {
ifres8352 = True


} else {
ifres8352 = False


}

ifres8351 = ifres8352


} else {
ifres8351 = False


}

if True == ifres8351 {
tmp8342 := PrimTail(V1298)

tmp8343 := PrimHead(tmp8342)

tmp8344 := PrimStr(tmp8343)

tmp8345 := PrimStringConcat(MakeString("partial function "), tmp8344)

tmp8346 := PrimCons(tmp8345, Nil)

__e.Return(PrimCons(symsimple_1error, tmp8346))
return


} else {
tmp8349 := PrimIsPair(V1298)

if True == tmp8349 {
tmp8347 := MakeNative(func(__e *ControlFlow) {
Z1299 := __e.Get(1)
_ = Z1299
__e.TailApply(PrimFunc(symshen_4partial), Z1299)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp8347, V1298)
return


} else {
__e.Return(V1298)
return
}


}


}, 1)

tmp8365 := Call(__e, ns2_1set, symshen_4partial, tmp8341)


_ = tmp8365

tmp8366 := MakeNative(func(__e *ControlFlow) {
V1302 := __e.Get(1)
_ = V1302
V1303 := __e.Get(2)
_ = V1303
tmp8380 := PrimEqual(Nil, V1302)

if True == tmp8380 {
__e.Return(PrimCloseStream(V1303))
return
} else {
tmp8378 := PrimIsPair(V1302)

var ifres8374 Obj

if True == tmp8378 {
tmp8376 := PrimHead(V1302)

tmp8377 := PrimIsPair(tmp8376)

var ifres8375 Obj

if True == tmp8377 {
ifres8375 = True


} else {
ifres8375 = False


}

ifres8374 = ifres8375


} else {
ifres8374 = False


}

if True == ifres8374 {
tmp8367 := PrimTail(V1302)

tmp8368 := PrimHead(V1302)

tmp8369 := Call(__e, PrimFunc(symshen_4write_1kl_1h), tmp8368, V1303)


_ = tmp8369

__e.TailApply(PrimFunc(symshen_4write_1kl), tmp8367, V1303)
return


} else {
tmp8372 := PrimIsPair(V1302)

if True == tmp8372 {
tmp8370 := PrimTail(V1302)

__e.TailApply(PrimFunc(symshen_4write_1kl), tmp8370, V1303)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4write_1kl)
return
}


}


}


}, 2)

tmp8381 := Call(__e, ns2_1set, symshen_4write_1kl, tmp8366)


_ = tmp8381

tmp8382 := MakeNative(func(__e *ControlFlow) {
V1306 := __e.Get(1)
_ = V1306
V1307 := __e.Get(2)
_ = V1307
tmp8422 := PrimIsPair(V1306)

var ifres8385 Obj

if True == tmp8422 {
tmp8420 := PrimHead(V1306)

tmp8421 := PrimEqual(symdefun, tmp8420)

var ifres8387 Obj

if True == tmp8421 {
tmp8418 := PrimTail(V1306)

tmp8419 := PrimIsPair(tmp8418)

var ifres8389 Obj

if True == tmp8419 {
tmp8415 := PrimTail(V1306)

tmp8416 := PrimHead(tmp8415)

tmp8417 := PrimEqual(symfail, tmp8416)

var ifres8391 Obj

if True == tmp8417 {
tmp8412 := PrimTail(V1306)

tmp8413 := PrimTail(tmp8412)

tmp8414 := PrimIsPair(tmp8413)

var ifres8393 Obj

if True == tmp8414 {
tmp8408 := PrimTail(V1306)

tmp8409 := PrimTail(tmp8408)

tmp8410 := PrimHead(tmp8409)

tmp8411 := PrimEqual(Nil, tmp8410)

var ifres8395 Obj

if True == tmp8411 {
tmp8404 := PrimTail(V1306)

tmp8405 := PrimTail(tmp8404)

tmp8406 := PrimTail(tmp8405)

tmp8407 := PrimIsPair(tmp8406)

var ifres8397 Obj

if True == tmp8407 {
tmp8399 := PrimTail(V1306)

tmp8400 := PrimTail(tmp8399)

tmp8401 := PrimTail(tmp8400)

tmp8402 := PrimTail(tmp8401)

tmp8403 := PrimEqual(Nil, tmp8402)

var ifres8398 Obj

if True == tmp8403 {
ifres8398 = True


} else {
ifres8398 = False


}

ifres8397 = ifres8398


} else {
ifres8397 = False


}

var ifres8396 Obj

if True == ifres8397 {
ifres8396 = True


} else {
ifres8396 = False


}

ifres8395 = ifres8396


} else {
ifres8395 = False


}

var ifres8394 Obj

if True == ifres8395 {
ifres8394 = True


} else {
ifres8394 = False


}

ifres8393 = ifres8394


} else {
ifres8393 = False


}

var ifres8392 Obj

if True == ifres8393 {
ifres8392 = True


} else {
ifres8392 = False


}

ifres8391 = ifres8392


} else {
ifres8391 = False


}

var ifres8390 Obj

if True == ifres8391 {
ifres8390 = True


} else {
ifres8390 = False


}

ifres8389 = ifres8390


} else {
ifres8389 = False


}

var ifres8388 Obj

if True == ifres8389 {
ifres8388 = True


} else {
ifres8388 = False


}

ifres8387 = ifres8388


} else {
ifres8387 = False


}

var ifres8386 Obj

if True == ifres8387 {
ifres8386 = True


} else {
ifres8386 = False


}

ifres8385 = ifres8386


} else {
ifres8385 = False


}

if True == ifres8385 {
__e.TailApply(PrimFunc(sympr), MakeString("(defun fail () shen.fail!)"), V1307)
return
} else {
tmp8383 := Call(__e, PrimFunc(symshen_4app), V1306, MakeString("\n\n"), symshen_4r)


__e.TailApply(PrimFunc(sympr), tmp8383, V1307)
return


}


}, 2)

tmp8423 := Call(__e, ns2_1set, symshen_4write_1kl_1h, tmp8382)


_ = tmp8423

tmp8424 := MakeNative(func(__e *ControlFlow) {
V1308 := __e.Get(1)
_ = V1308
tmp8433 := PrimEqual(MakeString(""), V1308)

if True == tmp8433 {
__e.Return(MakeString(".kl"))
return
} else {
tmp8431 := PrimEqual(MakeString(".shen"), V1308)

if True == tmp8431 {
__e.Return(MakeString(".kl"))
return
} else {
tmp8429 := Call(__e, PrimFunc(symshen_4_7string_2), V1308)


if True == tmp8429 {
tmp8425 := Call(__e, PrimFunc(symhdstr), V1308)


tmp8426 := PrimTailString(V1308)

tmp8427 := Call(__e, PrimFunc(symshen_4klfile), tmp8426)


__e.TailApply(PrimFunc(sym_8s), tmp8425, tmp8427)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4klfile)
return
}


}


}


}, 1)

__e.TailApply(ns2_1set, symshen_4klfile, tmp8424)
return




}, 0)

