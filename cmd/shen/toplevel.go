package main

import . "github.com/pyrex41/shen-go/kl"

var TopLevelMain = MakeNative(func(__e *ControlFlow) {
tmp8187 := MakeNative(func(__e *ControlFlow) {
tmp8188 := Call(__e, PrimFunc(symshen_4credits))


_ = tmp8188

__e.TailApply(PrimFunc(symshen_4loop))
return


}, 0)

tmp8189 := Call(__e, ns2_1set, symshen_4repl, tmp8187)


_ = tmp8189

tmp8190 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__selftop:
tmp8191 := Call(__e, PrimFunc(symshen_4initialise__environment))


_ = tmp8191

tmp8192 := Call(__e, PrimFunc(symshen_4prompt))


_ = tmp8192

tmp8193 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4read_1evaluate_1print))
return
}, 0)

tmp8194 := MakeNative(func(__e *ControlFlow) {
Z5741 := __e.Get(1)
_ = Z5741
__e.TailApply(PrimFunc(symshen_4toplevel_1display_1exception), Z5741)
return
}, 1)

tmp8195 := Call(__e, try_1catch, tmp8193, tmp8194)


_ = tmp8195

if PrimFunc(symshen_4loop) == __self {
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4loop))
return


}, 0)

tmp8196 := Call(__e, ns2_1set, symshen_4loop, tmp8190)


_ = tmp8196

tmp8197 := MakeNative(func(__e *ControlFlow) {
V5742 := __e.Get(1)
_ = V5742
tmp8198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symerror_1to_1string) {
return PrimErrorToString(V5742)
}
__typedArg0 := V5742
return Call(__e, PrimFunc(symerror_1to_1string), __typedArg0)
})()

tmp8199 := Call(__e, PrimFunc(symstoutput))


tmp8200 := Call(__e, PrimFunc(sympr), tmp8198, tmp8199)


_ = tmp8200

__e.TailApply(PrimFunc(symnl), MakeInteger(0))
return


}, 1)

tmp8201 := Call(__e, ns2_1set, symshen_4toplevel_1display_1exception, tmp8197)


_ = tmp8201

tmp8202 := MakeNative(func(__e *ControlFlow) {
tmp8203 := Call(__e, PrimFunc(symstoutput))


tmp8204 := Call(__e, PrimFunc(sympr), MakeString("\nShen, www.shenlanguage.org, copyright (C) 2010-2024, Mark Tarver\n"), tmp8203)


_ = tmp8204

tmp8205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dversion_d)
}
__typedArg0 := sym_dversion_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dlanguage_d)
}
__typedArg0 := sym_dlanguage_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dimplementation_d)
}
__typedArg0 := sym_dimplementation_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_drelease_d)
}
__typedArg0 := sym_drelease_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8209 := Call(__e, PrimFunc(symshen_4app), tmp8208, MakeString("\n"), symshen_4a)


tmp8211 := Call(__e, PrimFunc(symshen_4app), tmp8207, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" "))
__typedS1, __typedOK1 := TypedString(tmp8209)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" ")
__typedArg1 := tmp8209
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp8213 := Call(__e, PrimFunc(symshen_4app), tmp8206, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(", platform: "))
__typedS1, __typedOK1 := TypedString(tmp8211)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(", platform: ")
__typedArg1 := tmp8211
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp8215 := Call(__e, PrimFunc(symshen_4app), tmp8205, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(", language: "))
__typedS1, __typedOK1 := TypedString(tmp8213)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(", language: ")
__typedArg1 := tmp8213
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp8216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("version: S"))
__typedS1, __typedOK1 := TypedString(tmp8215)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("version: S")
__typedArg1 := tmp8215
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp8217 := Call(__e, PrimFunc(symstoutput))


tmp8218 := Call(__e, PrimFunc(sympr), tmp8216, tmp8217)


_ = tmp8218

tmp8219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dport_d)
}
__typedArg0 := sym_dport_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dporters_d)
}
__typedArg0 := sym_dporters_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8221 := Call(__e, PrimFunc(symshen_4app), tmp8220, MakeString("\n\n"), symshen_4a)


tmp8223 := Call(__e, PrimFunc(symshen_4app), tmp8219, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(", ported by "))
__typedS1, __typedOK1 := TypedString(tmp8221)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(", ported by ")
__typedArg1 := tmp8221
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp8224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("port "))
__typedS1, __typedOK1 := TypedString(tmp8223)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("port ")
__typedArg1 := tmp8223
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp8225 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8224, tmp8225)
return


}, 0)

tmp8226 := Call(__e, ns2_1set, symshen_4credits, tmp8202)


_ = tmp8226

tmp8227 := MakeNative(func(__e *ControlFlow) {
tmp8228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dcall_d, MakeInteger(0))
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp8228

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dinfs_d, MakeInteger(0))
}
__typedArg0 := symshen_4_dinfs_d
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp8229 := Call(__e, ns2_1set, symshen_4initialise__environment, tmp8227)


_ = tmp8229

tmp8230 := MakeNative(func(__e *ControlFlow) {
tmp8242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtc_d)
}
__typedArg0 := symshen_4_dtc_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

if True == tmp8242 {
tmp8231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dhistory_d)
}
__typedArg0 := symshen_4_dhistory_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8232 := Call(__e, PrimFunc(symlength), tmp8231)


tmp8233 := Call(__e, PrimFunc(symshen_4app), tmp8232, MakeString("+) "), symshen_4a)


tmp8234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\n("))
__typedS1, __typedOK1 := TypedString(tmp8233)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\n(")
__typedArg1 := tmp8233
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp8235 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8234, tmp8235)
return


} else {
tmp8236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dhistory_d)
}
__typedArg0 := symshen_4_dhistory_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8237 := Call(__e, PrimFunc(symlength), tmp8236)


tmp8238 := Call(__e, PrimFunc(symshen_4app), tmp8237, MakeString("-) "), symshen_4a)


tmp8239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\n("))
__typedS1, __typedOK1 := TypedString(tmp8238)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\n(")
__typedArg1 := tmp8238
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp8240 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8239, tmp8240)
return


}


}, 0)

tmp8243 := Call(__e, ns2_1set, symshen_4prompt, tmp8230)


_ = tmp8243

tmp8244 := MakeNative(func(__e *ControlFlow) {
tmp8245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dpackage_d)
}
__typedArg0 := symshen_4_dpackage_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__8171 := tmp8245
_ = let__8171

tmp8246 := Call(__e, PrimFunc(symstinput))


tmp8247 := Call(__e, PrimFunc(symlineread), tmp8246)


tmp8248 := Call(__e, PrimFunc(symshen_4package_1user_1input), let__8171, tmp8247)


let__8172 := tmp8248
_ = let__8172

tmp8249 := Call(__e, PrimFunc(symshen_4update_1history))


let__8173 := tmp8249
_ = let__8173

tmp8250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtc_d)
}
__typedArg0 := symshen_4_dtc_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), let__8172, let__8173, tmp8250)
return


}, 0)

tmp8251 := Call(__e, ns2_1set, symshen_4read_1evaluate_1print, tmp8244)


_ = tmp8251

tmp8252 := MakeNative(func(__e *ControlFlow) {
V5746 := __e.Get(1)
_ = V5746
V5747 := __e.Get(2)
_ = V5747
tmp8257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnull, V5746)
}
__typedArg0 := symnull
__typedArg1 := V5746
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8257 {
__e.Return(V5747)
return
} else {
tmp8253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V5746)
}
__typedArg0 := V5746
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

let__8174 := tmp8253
_ = let__8174

tmp8254 := Call(__e, PrimFunc(symexternal), V5746)


let__8175 := tmp8254
_ = let__8175

tmp8255 := MakeNative(func(__e *ControlFlow) {
Z5750 := __e.Get(1)
_ = Z5750
__e.TailApply(PrimFunc(symshen_4pui_1h), let__8174, let__8175, Z5750)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp8255, V5747)
return


}


}, 2)

tmp8258 := Call(__e, ns2_1set, symshen_4package_1user_1input, tmp8252)


_ = tmp8258

tmp8259 := MakeNative(func(__e *ControlFlow) {
V5755 := __e.Get(1)
_ = V5755
V5756 := __e.Get(2)
_ = V5756
V5757 := __e.Get(3)
_ = V5757
tmp8300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5757)
}
__typedArg0 := V5757
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8287 Obj

if True == tmp8300 {
tmp8298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5757)
}
__typedArg0 := V5757
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfn, tmp8298)
}
__typedArg0 := symfn
__typedArg1 := tmp8298
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8289 Obj

if True == tmp8299 {
tmp8296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5757)
}
__typedArg0 := V5757
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8296)
}
__typedArg0 := tmp8296
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8291 Obj

if True == tmp8297 {
tmp8293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5757)
}
__typedArg0 := V5757
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp8293)
}
__typedArg0 := tmp8293
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8294)
}
__typedArg0 := Nil
__typedArg1 := tmp8294
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8292 Obj

if True == tmp8295 {
ifres8292 = True


} else {
ifres8292 = False


}

ifres8291 = ifres8292


} else {
ifres8291 = False


}

var ifres8290 Obj

if True == ifres8291 {
ifres8290 = True


} else {
ifres8290 = False


}

ifres8289 = ifres8290


} else {
ifres8289 = False


}

var ifres8288 Obj

if True == ifres8289 {
ifres8288 = True


} else {
ifres8288 = False


}

ifres8287 = ifres8288


} else {
ifres8287 = False


}

if True == ifres8287 {
tmp8265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5757)
}
__typedArg0 := V5757
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8265)
}
__typedArg0 := tmp8265
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8267 := Call(__e, PrimFunc(symshen_4internal_2), tmp8266, V5755, V5756)


if True == tmp8267 {
tmp8260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5757)
}
__typedArg0 := V5757
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8260)
}
__typedArg0 := tmp8260
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8262 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V5755, tmp8261)


tmp8263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8262, Nil)
}
__typedArg0 := tmp8262
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp8263)
}
__typedArg0 := symfn
__typedArg1 := tmp8263
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V5757)
return
}


} else {
tmp8285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5757)
}
__typedArg0 := V5757
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8285 {
tmp8282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5757)
}
__typedArg0 := V5757
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8283 := Call(__e, PrimFunc(symshen_4internal_2), tmp8282, V5755, V5756)


if True == tmp8283 {
tmp8268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5757)
}
__typedArg0 := V5757
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8269 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V5755, tmp8268)


tmp8270 := MakeNative(func(__e *ControlFlow) {
Z5758 := __e.Get(1)
_ = Z5758
__e.TailApply(PrimFunc(symshen_4pui_1h), V5755, V5756, Z5758)
return
}, 1)

tmp8271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5757)
}
__typedArg0 := V5757
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8272 := Call(__e, PrimFunc(symmap), tmp8270, tmp8271)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8269, tmp8272)
}
__typedArg0 := tmp8269
__typedArg1 := tmp8272
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5757)
}
__typedArg0 := V5757
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8279)
}
__typedArg0 := tmp8279
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8280 {
tmp8273 := MakeNative(func(__e *ControlFlow) {
Z5759 := __e.Get(1)
_ = Z5759
__e.TailApply(PrimFunc(symshen_4pui_1h), V5755, V5756, Z5759)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp8273, V5757)
return


} else {
tmp8274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5757)
}
__typedArg0 := V5757
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8275 := MakeNative(func(__e *ControlFlow) {
Z5760 := __e.Get(1)
_ = Z5760
__e.TailApply(PrimFunc(symshen_4pui_1h), V5755, V5756, Z5760)
return
}, 1)

tmp8276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5757)
}
__typedArg0 := V5757
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8277 := Call(__e, PrimFunc(symmap), tmp8275, tmp8276)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8274, tmp8277)
}
__typedArg0 := tmp8274
__typedArg1 := tmp8277
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}


} else {
__e.Return(V5757)
return
}


}


}, 3)

tmp8301 := Call(__e, ns2_1set, symshen_4pui_1h, tmp8259)


_ = tmp8301

tmp8302 := MakeNative(func(__e *ControlFlow) {
tmp8303 := Call(__e, PrimFunc(symit))


tmp8304 := Call(__e, PrimFunc(symshen_4trim_1it), tmp8303)


tmp8305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dhistory_d)
}
__typedArg0 := symshen_4_dhistory_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8304, tmp8305)
}
__typedArg0 := tmp8304
__typedArg1 := tmp8305
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dhistory_d, tmp8306)
}
__typedArg0 := symshen_4_dhistory_d
__typedArg1 := tmp8306
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp8307 := Call(__e, ns2_1set, symshen_4update_1history, tmp8302)


_ = tmp8307

tmp8308 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V5761 := __self1
_ = V5761
tmp8316 := Call(__e, PrimFunc(symshen_4_7string_2), V5761)


var ifres8311 Obj

if True == tmp8316 {
tmp8313 := Call(__e, PrimFunc(symhdstr), V5761)


tmp8314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp8313)
}
__typedArg0 := tmp8313
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp8315 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp8314)


var ifres8312 Obj

if True == tmp8315 {
ifres8312 = True


} else {
ifres8312 = False


}

ifres8311 = ifres8312


} else {
ifres8311 = False


}

if True == ifres8311 {
if PrimFunc(symshen_4trim_1it) == __self {
__self1 = (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5761)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5761
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4trim_1it), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5761)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5761
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())
return


} else {
__e.Return(V5761)
return
}


}, 1)

tmp8317 := Call(__e, ns2_1set, symshen_4trim_1it, tmp8308)


_ = tmp8317

tmp8318 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__selftop:
V5780 := __self1
_ = V5780
V5781 := __self2
_ = V5781
V5782 := __self3
_ = V5782
tmp8437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5780)
}
__typedArg0 := V5780
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8406 Obj

if True == tmp8437 {
tmp8435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5780)
}
__typedArg0 := V5780
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8435)
}
__typedArg0 := Nil
__typedArg1 := tmp8435
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8408 Obj

if True == tmp8436 {
tmp8434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8410 Obj

if True == tmp8434 {
tmp8432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8433 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8432)


var ifres8412 Obj

if True == tmp8433 {
tmp8429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8430 := Call(__e, PrimFunc(symhdstr), tmp8429)


tmp8431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("!"), tmp8430)
}
__typedArg0 := MakeString("!")
__typedArg1 := tmp8430
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8414 Obj

if True == tmp8431 {
tmp8426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8428 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8426)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8426
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres8416 Obj

if True == tmp8428 {
tmp8422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8424 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8422)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8422
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp8425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("!"), tmp8424)
}
__typedArg0 := MakeString("!")
__typedArg1 := tmp8424
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8418 Obj

if True == tmp8425 {
tmp8420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8420)
}
__typedArg0 := tmp8420
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8419 Obj

if True == tmp8421 {
ifres8419 = True


} else {
ifres8419 = False


}

ifres8418 = ifres8419


} else {
ifres8418 = False


}

var ifres8417 Obj

if True == ifres8418 {
ifres8417 = True


} else {
ifres8417 = False


}

ifres8416 = ifres8417


} else {
ifres8416 = False


}

var ifres8415 Obj

if True == ifres8416 {
ifres8415 = True


} else {
ifres8415 = False


}

ifres8414 = ifres8415


} else {
ifres8414 = False


}

var ifres8413 Obj

if True == ifres8414 {
ifres8413 = True


} else {
ifres8413 = False


}

ifres8412 = ifres8413


} else {
ifres8412 = False


}

var ifres8411 Obj

if True == ifres8412 {
ifres8411 = True


} else {
ifres8411 = False


}

ifres8410 = ifres8411


} else {
ifres8410 = False


}

var ifres8409 Obj

if True == ifres8410 {
ifres8409 = True


} else {
ifres8409 = False


}

ifres8408 = ifres8409


} else {
ifres8408 = False


}

var ifres8407 Obj

if True == ifres8408 {
ifres8407 = True


} else {
ifres8407 = False


}

ifres8406 = ifres8407


} else {
ifres8406 = False


}

if True == ifres8406 {
tmp8319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8319)
}
__typedArg0 := tmp8319
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8321 := Call(__e, PrimFunc(symread_1from_1string), tmp8320)


let__8176 := tmp8321
_ = let__8176

tmp8322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8322)
}
__typedArg0 := tmp8322
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8323, tmp8324)
}
__typedArg0 := tmp8323
__typedArg1 := tmp8324
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dhistory_d, tmp8325)
}
__typedArg0 := symshen_4_dhistory_d
__typedArg1 := tmp8325
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__8177 := tmp8326
_ = let__8177

tmp8327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8327)
}
__typedArg0 := tmp8327
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8329 := Call(__e, PrimFunc(symshen_4app), tmp8328, MakeString("\n"), symshen_4a)


tmp8330 := Call(__e, PrimFunc(symstoutput))


tmp8331 := Call(__e, PrimFunc(sympr), tmp8329, tmp8330)


let__8178 := tmp8331
_ = let__8178

if PrimFunc(symshen_4evaluate_1lineread) == __self {
__self1, __self2, __self3 = let__8176, let__8177, V5782
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), let__8176, let__8177, V5782)
return


} else {
tmp8404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5780)
}
__typedArg0 := V5780
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8388 Obj

if True == tmp8404 {
tmp8402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5780)
}
__typedArg0 := V5780
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8402)
}
__typedArg0 := Nil
__typedArg1 := tmp8402
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8390 Obj

if True == tmp8403 {
tmp8401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8392 Obj

if True == tmp8401 {
tmp8399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8400 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8399)


var ifres8394 Obj

if True == tmp8400 {
tmp8396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8397 := Call(__e, PrimFunc(symhdstr), tmp8396)


tmp8398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("!"), tmp8397)
}
__typedArg0 := MakeString("!")
__typedArg1 := tmp8397
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8395 Obj

if True == tmp8398 {
ifres8395 = True


} else {
ifres8395 = False


}

ifres8394 = ifres8395


} else {
ifres8394 = False


}

var ifres8393 Obj

if True == ifres8394 {
ifres8393 = True


} else {
ifres8393 = False


}

ifres8392 = ifres8393


} else {
ifres8392 = False


}

var ifres8391 Obj

if True == ifres8392 {
ifres8391 = True


} else {
ifres8391 = False


}

ifres8390 = ifres8391


} else {
ifres8390 = False


}

var ifres8389 Obj

if True == ifres8390 {
ifres8389 = True


} else {
ifres8389 = False


}

ifres8388 = ifres8389


} else {
ifres8388 = False


}

if True == ifres8388 {
tmp8337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8337)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8337
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})(), MakeString(""))
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8337)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8337
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8332 Obj

if True == tmp8339 {
ifres8332 = Nil


} else {
tmp8333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8335 := Call(__e, PrimFunc(symread_1from_1string), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8333)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8333
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp8336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8335)
}
__typedArg0 := tmp8335
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

ifres8332 = tmp8336


}

let__8179 := ifres8332
_ = let__8179

tmp8340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8340)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8340
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

tmp8342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8343 := Call(__e, PrimFunc(symshen_4use_1history), let__8179, tmp8341, tmp8342)


let__8180 := tmp8343
_ = let__8180

tmp8344 := Call(__e, PrimFunc(symshen_4app), let__8180, MakeString("\n"), symshen_4a)


tmp8345 := Call(__e, PrimFunc(symstoutput))


tmp8346 := Call(__e, PrimFunc(sympr), tmp8344, tmp8345)


let__8181 := tmp8346
_ = let__8181

tmp8347 := Call(__e, PrimFunc(symread_1from_1string), let__8180)


let__8182 := tmp8347
_ = let__8182

tmp8348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8180, tmp8348)
}
__typedArg0 := let__8180
__typedArg1 := tmp8348
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dhistory_d, tmp8349)
}
__typedArg0 := symshen_4_dhistory_d
__typedArg1 := tmp8349
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__8183 := tmp8350
_ = let__8183

if PrimFunc(symshen_4evaluate_1lineread) == __self {
__self1, __self2, __self3 = let__8182, let__8183, V5782
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), let__8182, let__8183, V5782)
return


} else {
tmp8386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5780)
}
__typedArg0 := V5780
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8370 Obj

if True == tmp8386 {
tmp8384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5780)
}
__typedArg0 := V5780
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8384)
}
__typedArg0 := Nil
__typedArg1 := tmp8384
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8372 Obj

if True == tmp8385 {
tmp8383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8374 Obj

if True == tmp8383 {
tmp8381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8382 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8381)


var ifres8376 Obj

if True == tmp8382 {
tmp8378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8379 := Call(__e, PrimFunc(symhdstr), tmp8378)


tmp8380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("%"), tmp8379)
}
__typedArg0 := MakeString("%")
__typedArg1 := tmp8379
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8377 Obj

if True == tmp8380 {
ifres8377 = True


} else {
ifres8377 = False


}

ifres8376 = ifres8377


} else {
ifres8376 = False


}

var ifres8375 Obj

if True == ifres8376 {
ifres8375 = True


} else {
ifres8375 = False


}

ifres8374 = ifres8375


} else {
ifres8374 = False


}

var ifres8373 Obj

if True == ifres8374 {
ifres8373 = True


} else {
ifres8373 = False


}

ifres8372 = ifres8373


} else {
ifres8372 = False


}

var ifres8371 Obj

if True == ifres8372 {
ifres8371 = True


} else {
ifres8371 = False


}

ifres8370 = ifres8371


} else {
ifres8370 = False


}

if True == ifres8370 {
tmp8356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8356)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8356
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})(), MakeString(""))
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8356)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8356
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8351 Obj

if True == tmp8358 {
ifres8351 = Nil


} else {
tmp8352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8354 := Call(__e, PrimFunc(symread_1from_1string), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8352)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8352
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp8355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8354)
}
__typedArg0 := tmp8354
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

ifres8351 = tmp8355


}

let__8184 := ifres8351
_ = let__8184

tmp8359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8359)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8359
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

tmp8361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8362 := Call(__e, PrimFunc(symshen_4peek_1history), let__8184, tmp8360, tmp8361)


let__8185 := tmp8362
_ = let__8185

tmp8363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5781)
}
__typedArg0 := V5781
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dhistory_d, tmp8363)
}
__typedArg0 := symshen_4_dhistory_d
__typedArg1 := tmp8363
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__8186 := tmp8364
_ = let__8186

__e.TailApply(PrimFunc(symabort))
return


} else {
tmp8368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, V5782)
}
__typedArg0 := True
__typedArg1 := V5782
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8368 {
__e.TailApply(PrimFunc(symshen_4check_1eval_1and_1print), V5780)
return
} else {
tmp8366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, V5782)
}
__typedArg0 := False
__typedArg1 := V5782
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8366 {
__e.TailApply(PrimFunc(symshen_4eval_1and_1print), V5780)
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.evaluate-lineread"))
}
__typedArg0 := MakeString("implementation error in shen.evaluate-lineread")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}


}


}, 3)

tmp8438 := Call(__e, ns2_1set, symshen_4evaluate_1lineread, tmp8318)


_ = tmp8438

tmp8439 := MakeNative(func(__e *ControlFlow) {
V5794 := __e.Get(1)
_ = V5794
V5795 := __e.Get(2)
_ = V5795
V5796 := __e.Get(3)
_ = V5796
tmp8445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V5794)
}
__typedArg0 := V5794
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})()

if True == tmp8445 {
tmp8440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(V5794)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := V5794
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()

tmp8441 := Call(__e, PrimFunc(symreverse), V5796)


__e.TailApply(PrimFunc(symnth), tmp8440, tmp8441)
return


} else {
tmp8443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V5794)
}
__typedArg0 := V5794
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp8443 {
__e.TailApply(PrimFunc(symshen_4string_1match), V5795, V5796)
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("! expects a number or a symbol\n"))
}
__typedArg0 := MakeString("! expects a number or a symbol\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 3)

tmp8446 := Call(__e, ns2_1set, symshen_4use_1history, tmp8439)


_ = tmp8446

tmp8447 := MakeNative(func(__e *ControlFlow) {
V5797 := __e.Get(1)
_ = V5797
V5798 := __e.Get(2)
_ = V5798
V5799 := __e.Get(3)
_ = V5799
tmp8461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V5797)
}
__typedArg0 := V5797
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})()

if True == tmp8461 {
tmp8448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(V5797)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := V5797
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()

tmp8449 := Call(__e, PrimFunc(symreverse), V5799)


tmp8450 := Call(__e, PrimFunc(symnth), tmp8448, tmp8449)


tmp8451 := Call(__e, PrimFunc(symshen_4app), tmp8450, MakeString(""), symshen_4a)


tmp8452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\n"))
__typedS1, __typedOK1 := TypedString(tmp8451)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\n")
__typedArg1 := tmp8451
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp8453 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8452, tmp8453)
return


} else {
tmp8459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V5798, MakeString(""))
}
__typedArg0 := V5798
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8456 Obj

if True == tmp8459 {
ifres8456 = True


} else {
tmp8458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V5797)
}
__typedArg0 := V5797
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

var ifres8457 Obj

if True == tmp8458 {
ifres8457 = True


} else {
ifres8457 = False


}

ifres8456 = ifres8457


}

if True == ifres8456 {
tmp8454 := Call(__e, PrimFunc(symreverse), V5799)


__e.TailApply(PrimFunc(symshen_4recursive_1string_1match), MakeInteger(0), V5798, tmp8454)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("% expects a number or a symbol\n"))
}
__typedArg0 := MakeString("% expects a number or a symbol\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 3)

tmp8462 := Call(__e, ns2_1set, symshen_4peek_1history, tmp8447)


_ = tmp8462

tmp8463 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V5809 := __self1
_ = V5809
V5810 := __self2
_ = V5810
tmp8474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5810)
}
__typedArg0 := Nil
__typedArg1 := V5810
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8474 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("\ninput not found"))
}
__typedArg0 := MakeString("\ninput not found")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp8472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5810)
}
__typedArg0 := V5810
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8468 Obj

if True == tmp8472 {
tmp8470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5810)
}
__typedArg0 := V5810
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8471 := Call(__e, PrimFunc(symshen_4string_1prefix_2), V5809, tmp8470)


var ifres8469 Obj

if True == tmp8471 {
ifres8469 = True


} else {
ifres8469 = False


}

ifres8468 = ifres8469


} else {
ifres8468 = False


}

if True == ifres8468 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5810)
}
__typedArg0 := V5810
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp8466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5810)
}
__typedArg0 := V5810
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8466 {
tmp8464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5810)
}
__typedArg0 := V5810
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4string_1match) == __self {
__self1, __self2 = V5809, tmp8464
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4string_1match), V5809, tmp8464)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.string-match"))
}
__typedArg0 := MakeString("implementation error in shen.string-match")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 2)

tmp8475 := Call(__e, ns2_1set, symshen_4string_1match, tmp8463)


_ = tmp8475

tmp8476 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V5818 := __self1
_ = V5818
V5819 := __self2
_ = V5819
tmp8513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V5818)
}
__typedArg0 := MakeString("")
__typedArg1 := V5818
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8513 {
__e.Return(True)
return
} else {
tmp8511 := Call(__e, PrimFunc(symshen_4_7string_2), V5818)


var ifres8506 Obj

if True == tmp8511 {
tmp8508 := Call(__e, PrimFunc(symhdstr), V5818)


tmp8509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp8508)
}
__typedArg0 := tmp8508
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp8510 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp8509)


var ifres8507 Obj

if True == tmp8510 {
ifres8507 = True


} else {
ifres8507 = False


}

ifres8506 = ifres8507


} else {
ifres8506 = False


}

if True == ifres8506 {
if PrimFunc(symshen_4string_1prefix_2) == __self {
__self1, __self2 = (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5818)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5818
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})(), V5819
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4string_1prefix_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5818)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5818
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})(), V5819)
return


} else {
tmp8504 := Call(__e, PrimFunc(symshen_4_7string_2), V5819)


var ifres8499 Obj

if True == tmp8504 {
tmp8501 := Call(__e, PrimFunc(symhdstr), V5819)


tmp8502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp8501)
}
__typedArg0 := tmp8501
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp8503 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp8502)


var ifres8500 Obj

if True == tmp8503 {
ifres8500 = True


} else {
ifres8500 = False


}

ifres8499 = ifres8500


} else {
ifres8499 = False


}

if True == ifres8499 {
if PrimFunc(symshen_4string_1prefix_2) == __self {
__self1, __self2 = V5818, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5819)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5819
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4string_1prefix_2), V5818, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5819)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5819
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())
return


} else {
tmp8497 := Call(__e, PrimFunc(symshen_4_7string_2), V5819)


var ifres8493 Obj

if True == tmp8497 {
tmp8495 := Call(__e, PrimFunc(symhdstr), V5819)


tmp8496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("("), tmp8495)
}
__typedArg0 := MakeString("(")
__typedArg1 := tmp8495
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8494 Obj

if True == tmp8496 {
ifres8494 = True


} else {
ifres8494 = False


}

ifres8493 = ifres8494


} else {
ifres8493 = False


}

if True == ifres8493 {
if PrimFunc(symshen_4string_1prefix_2) == __self {
__self1, __self2 = V5818, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5819)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5819
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4string_1prefix_2), V5818, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5819)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5819
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())
return


} else {
tmp8491 := Call(__e, PrimFunc(symshen_4_7string_2), V5818)


var ifres8483 Obj

if True == tmp8491 {
tmp8490 := Call(__e, PrimFunc(symshen_4_7string_2), V5819)


var ifres8485 Obj

if True == tmp8490 {
tmp8487 := Call(__e, PrimFunc(symhdstr), V5818)


tmp8488 := Call(__e, PrimFunc(symhdstr), V5819)


tmp8489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp8487, tmp8488)
}
__typedArg0 := tmp8487
__typedArg1 := tmp8488
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8486 Obj

if True == tmp8489 {
ifres8486 = True


} else {
ifres8486 = False


}

ifres8485 = ifres8486


} else {
ifres8485 = False


}

var ifres8484 Obj

if True == ifres8485 {
ifres8484 = True


} else {
ifres8484 = False


}

ifres8483 = ifres8484


} else {
ifres8483 = False


}

if True == ifres8483 {
tmp8480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5818)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5818
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

if PrimFunc(symshen_4string_1prefix_2) == __self {
__self1, __self2 = tmp8480, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5819)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5819
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4string_1prefix_2), tmp8480, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5819)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5819
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())
return


} else {
__e.Return(False)
return
}


}


}


}


}


}, 2)

tmp8514 := Call(__e, ns2_1set, symshen_4string_1prefix_2, tmp8476)


_ = tmp8514

tmp8515 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__selftop:
V5830 := __self1
_ = V5830
V5831 := __self2
_ = V5831
V5832 := __self3
_ = V5832
tmp8530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5832)
}
__typedArg0 := Nil
__typedArg1 := V5832
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8530 {
__e.Return(symshen_4skip)
return
} else {
tmp8528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5832)
}
__typedArg0 := V5832
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8528 {
tmp8523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5832)
}
__typedArg0 := V5832
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8524 := Call(__e, PrimFunc(symshen_4string_1prefix_2), V5831, tmp8523)


var ifres8516 Obj

if True == tmp8524 {
tmp8517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5832)
}
__typedArg0 := V5832
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8518 := Call(__e, PrimFunc(symshen_4app), tmp8517, MakeString("\n"), symshen_4a)


tmp8520 := Call(__e, PrimFunc(symshen_4app), V5830, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(". "))
__typedS1, __typedOK1 := TypedString(tmp8518)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(". ")
__typedArg1 := tmp8518
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp8521 := Call(__e, PrimFunc(symstoutput))


tmp8522 := Call(__e, PrimFunc(sympr), tmp8520, tmp8521)


ifres8516 = tmp8522


} else {
ifres8516 = symshen_4skip


}

_ = ifres8516

tmp8525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5830)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5830
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()

tmp8526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5832)
}
__typedArg0 := V5832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4recursive_1string_1match) == __self {
__self1, __self2, __self3 = tmp8525, V5831, tmp8526
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4recursive_1string_1match), tmp8525, V5831, tmp8526)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.recursive-string-match"))
}
__typedArg0 := MakeString("implementation error in shen.recursive-string-match")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 3)

__e.TailApply(ns2_1set, symshen_4recursive_1string_1match, tmp8515)
return




}, 0)

