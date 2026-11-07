package main

import . "github.com/pyrex41/shen-go/kl"

var TopLevelMain = MakeNative(func(__e *ControlFlow) {
tmp8207 := MakeNative(func(__e *ControlFlow) {
tmp8208 := Call(__e, PrimFunc(symshen_4credits))


_ = tmp8208

__e.TailApply(PrimFunc(symshen_4loop))
return


}, 0)

tmp8209 := Call(__e, ns2_1set, symshen_4repl, tmp8207)


_ = tmp8209

tmp8210 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__selftop:
tmp8211 := Call(__e, PrimFunc(symshen_4initialise__environment))


_ = tmp8211

tmp8212 := Call(__e, PrimFunc(symshen_4prompt))


_ = tmp8212

tmp8213 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4read_1evaluate_1print))
return
}, 0)

tmp8214 := MakeNative(func(__e *ControlFlow) {
Z5763 := __e.Get(1)
_ = Z5763
__e.TailApply(PrimFunc(symshen_4toplevel_1display_1exception), Z5763)
return
}, 1)

tmp8215 := Call(__e, try_1catch, tmp8213, tmp8214)


_ = tmp8215

if PrimFunc(symshen_4loop) == __self {
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4loop))
return


}, 0)

tmp8216 := Call(__e, ns2_1set, symshen_4loop, tmp8210)


_ = tmp8216

tmp8217 := MakeNative(func(__e *ControlFlow) {
V5764 := __e.Get(1)
_ = V5764
tmp8218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symerror_1to_1string) {
return PrimErrorToString(V5764)
}
__typedArg0 := V5764
return Call(__e, PrimFunc(symerror_1to_1string), __typedArg0)
})()

tmp8219 := Call(__e, PrimFunc(symstoutput))


tmp8220 := Call(__e, PrimFunc(sympr), tmp8218, tmp8219)


_ = tmp8220

__e.TailApply(PrimFunc(symnl), MakeInteger(0))
return


}, 1)

tmp8221 := Call(__e, ns2_1set, symshen_4toplevel_1display_1exception, tmp8217)


_ = tmp8221

tmp8222 := MakeNative(func(__e *ControlFlow) {
tmp8223 := Call(__e, PrimFunc(symstoutput))


tmp8224 := Call(__e, PrimFunc(sympr), MakeString("\nShen, www.shenlanguage.org, copyright (C) 2010-2026, Mark Tarver\n"), tmp8223)


_ = tmp8224

tmp8225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dversion_d)
}
__typedArg0 := sym_dversion_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dlanguage_d)
}
__typedArg0 := sym_dlanguage_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dimplementation_d)
}
__typedArg0 := sym_dimplementation_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_drelease_d)
}
__typedArg0 := sym_drelease_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8229 := Call(__e, PrimFunc(symshen_4app), tmp8228, MakeString("\n"), symshen_4a)


tmp8231 := Call(__e, PrimFunc(symshen_4app), tmp8227, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" "))
__typedS1, __typedOK1 := TypedString(tmp8229)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" ")
__typedArg1 := tmp8229
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp8233 := Call(__e, PrimFunc(symshen_4app), tmp8226, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(", platform: "))
__typedS1, __typedOK1 := TypedString(tmp8231)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(", platform: ")
__typedArg1 := tmp8231
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp8235 := Call(__e, PrimFunc(symshen_4app), tmp8225, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(", language: "))
__typedS1, __typedOK1 := TypedString(tmp8233)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(", language: ")
__typedArg1 := tmp8233
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp8236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("version: S"))
__typedS1, __typedOK1 := TypedString(tmp8235)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("version: S")
__typedArg1 := tmp8235
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp8237 := Call(__e, PrimFunc(symstoutput))


tmp8238 := Call(__e, PrimFunc(sympr), tmp8236, tmp8237)


_ = tmp8238

tmp8239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dport_d)
}
__typedArg0 := sym_dport_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dporters_d)
}
__typedArg0 := sym_dporters_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8241 := Call(__e, PrimFunc(symshen_4app), tmp8240, MakeString("\n\n"), symshen_4a)


tmp8243 := Call(__e, PrimFunc(symshen_4app), tmp8239, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(", ported by "))
__typedS1, __typedOK1 := TypedString(tmp8241)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(", ported by ")
__typedArg1 := tmp8241
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp8244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("port "))
__typedS1, __typedOK1 := TypedString(tmp8243)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("port ")
__typedArg1 := tmp8243
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp8245 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8244, tmp8245)
return


}, 0)

tmp8246 := Call(__e, ns2_1set, symshen_4credits, tmp8222)


_ = tmp8246

tmp8247 := MakeNative(func(__e *ControlFlow) {
tmp8248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dcall_d, MakeInteger(0))
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp8248

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

tmp8249 := Call(__e, ns2_1set, symshen_4initialise__environment, tmp8247)


_ = tmp8249

tmp8250 := MakeNative(func(__e *ControlFlow) {
tmp8262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtc_d)
}
__typedArg0 := symshen_4_dtc_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

if True == tmp8262 {
tmp8251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dhistory_d)
}
__typedArg0 := symshen_4_dhistory_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8252 := Call(__e, PrimFunc(symlength), tmp8251)


tmp8253 := Call(__e, PrimFunc(symshen_4app), tmp8252, MakeString("+) "), symshen_4a)


tmp8254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\n("))
__typedS1, __typedOK1 := TypedString(tmp8253)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\n(")
__typedArg1 := tmp8253
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp8255 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8254, tmp8255)
return


} else {
tmp8256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dhistory_d)
}
__typedArg0 := symshen_4_dhistory_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8257 := Call(__e, PrimFunc(symlength), tmp8256)


tmp8258 := Call(__e, PrimFunc(symshen_4app), tmp8257, MakeString("-) "), symshen_4a)


tmp8259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\n("))
__typedS1, __typedOK1 := TypedString(tmp8258)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\n(")
__typedArg1 := tmp8258
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp8260 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8259, tmp8260)
return


}


}, 0)

tmp8263 := Call(__e, ns2_1set, symshen_4prompt, tmp8250)


_ = tmp8263

tmp8264 := MakeNative(func(__e *ControlFlow) {
tmp8265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dpackage_d)
}
__typedArg0 := symshen_4_dpackage_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__8191 := tmp8265
_ = let__8191

tmp8266 := Call(__e, PrimFunc(symstinput))


tmp8267 := Call(__e, PrimFunc(symlineread), tmp8266)


tmp8268 := Call(__e, PrimFunc(symshen_4package_1user_1input), let__8191, tmp8267)


let__8192 := tmp8268
_ = let__8192

tmp8269 := Call(__e, PrimFunc(symshen_4update_1history))


let__8193 := tmp8269
_ = let__8193

tmp8270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtc_d)
}
__typedArg0 := symshen_4_dtc_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), let__8192, let__8193, tmp8270)
return


}, 0)

tmp8271 := Call(__e, ns2_1set, symshen_4read_1evaluate_1print, tmp8264)


_ = tmp8271

tmp8272 := MakeNative(func(__e *ControlFlow) {
V5768 := __e.Get(1)
_ = V5768
V5769 := __e.Get(2)
_ = V5769
tmp8277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnull, V5768)
}
__typedArg0 := symnull
__typedArg1 := V5768
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8277 {
__e.Return(V5769)
return
} else {
tmp8273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V5768)
}
__typedArg0 := V5768
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

let__8194 := tmp8273
_ = let__8194

tmp8274 := Call(__e, PrimFunc(symexternal), V5768)


let__8195 := tmp8274
_ = let__8195

tmp8275 := MakeNative(func(__e *ControlFlow) {
Z5772 := __e.Get(1)
_ = Z5772
__e.TailApply(PrimFunc(symshen_4pui_1h), let__8194, let__8195, Z5772)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp8275, V5769)
return


}


}, 2)

tmp8278 := Call(__e, ns2_1set, symshen_4package_1user_1input, tmp8272)


_ = tmp8278

tmp8279 := MakeNative(func(__e *ControlFlow) {
V5777 := __e.Get(1)
_ = V5777
V5778 := __e.Get(2)
_ = V5778
V5779 := __e.Get(3)
_ = V5779
tmp8320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5779)
}
__typedArg0 := V5779
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8307 Obj

if True == tmp8320 {
tmp8318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5779)
}
__typedArg0 := V5779
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfn, tmp8318)
}
__typedArg0 := symfn
__typedArg1 := tmp8318
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8309 Obj

if True == tmp8319 {
tmp8316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5779)
}
__typedArg0 := V5779
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8316)
}
__typedArg0 := tmp8316
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8311 Obj

if True == tmp8317 {
tmp8313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5779)
}
__typedArg0 := V5779
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp8313)
}
__typedArg0 := tmp8313
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8314)
}
__typedArg0 := Nil
__typedArg1 := tmp8314
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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

var ifres8310 Obj

if True == ifres8311 {
ifres8310 = True


} else {
ifres8310 = False


}

ifres8309 = ifres8310


} else {
ifres8309 = False


}

var ifres8308 Obj

if True == ifres8309 {
ifres8308 = True


} else {
ifres8308 = False


}

ifres8307 = ifres8308


} else {
ifres8307 = False


}

if True == ifres8307 {
tmp8285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5779)
}
__typedArg0 := V5779
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8285)
}
__typedArg0 := tmp8285
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8287 := Call(__e, PrimFunc(symshen_4internal_2), tmp8286, V5777, V5778)


if True == tmp8287 {
tmp8280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5779)
}
__typedArg0 := V5779
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8280)
}
__typedArg0 := tmp8280
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8282 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V5777, tmp8281)


tmp8283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8282, Nil)
}
__typedArg0 := tmp8282
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp8283)
}
__typedArg0 := symfn
__typedArg1 := tmp8283
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V5779)
return
}


} else {
tmp8305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5779)
}
__typedArg0 := V5779
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8305 {
tmp8302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5779)
}
__typedArg0 := V5779
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8303 := Call(__e, PrimFunc(symshen_4internal_2), tmp8302, V5777, V5778)


if True == tmp8303 {
tmp8288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5779)
}
__typedArg0 := V5779
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8289 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V5777, tmp8288)


tmp8290 := MakeNative(func(__e *ControlFlow) {
Z5780 := __e.Get(1)
_ = Z5780
__e.TailApply(PrimFunc(symshen_4pui_1h), V5777, V5778, Z5780)
return
}, 1)

tmp8291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5779)
}
__typedArg0 := V5779
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8292 := Call(__e, PrimFunc(symmap), tmp8290, tmp8291)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8289, tmp8292)
}
__typedArg0 := tmp8289
__typedArg1 := tmp8292
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp8299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5779)
}
__typedArg0 := V5779
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8299)
}
__typedArg0 := tmp8299
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8300 {
tmp8293 := MakeNative(func(__e *ControlFlow) {
Z5781 := __e.Get(1)
_ = Z5781
__e.TailApply(PrimFunc(symshen_4pui_1h), V5777, V5778, Z5781)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp8293, V5779)
return


} else {
tmp8294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5779)
}
__typedArg0 := V5779
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8295 := MakeNative(func(__e *ControlFlow) {
Z5782 := __e.Get(1)
_ = Z5782
__e.TailApply(PrimFunc(symshen_4pui_1h), V5777, V5778, Z5782)
return
}, 1)

tmp8296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5779)
}
__typedArg0 := V5779
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8297 := Call(__e, PrimFunc(symmap), tmp8295, tmp8296)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8294, tmp8297)
}
__typedArg0 := tmp8294
__typedArg1 := tmp8297
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}


} else {
__e.Return(V5779)
return
}


}


}, 3)

tmp8321 := Call(__e, ns2_1set, symshen_4pui_1h, tmp8279)


_ = tmp8321

tmp8322 := MakeNative(func(__e *ControlFlow) {
tmp8323 := Call(__e, PrimFunc(symit))


tmp8324 := Call(__e, PrimFunc(symshen_4trim_1it), tmp8323)


tmp8325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dhistory_d)
}
__typedArg0 := symshen_4_dhistory_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8324, tmp8325)
}
__typedArg0 := tmp8324
__typedArg1 := tmp8325
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dhistory_d, tmp8326)
}
__typedArg0 := symshen_4_dhistory_d
__typedArg1 := tmp8326
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp8327 := Call(__e, ns2_1set, symshen_4update_1history, tmp8322)


_ = tmp8327

tmp8328 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V5783 := __self1
_ = V5783
tmp8336 := Call(__e, PrimFunc(symshen_4_7string_2), V5783)


var ifres8331 Obj

if True == tmp8336 {
tmp8333 := Call(__e, PrimFunc(symhdstr), V5783)


tmp8334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp8333)
}
__typedArg0 := tmp8333
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp8335 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp8334)


var ifres8332 Obj

if True == tmp8335 {
ifres8332 = True


} else {
ifres8332 = False


}

ifres8331 = ifres8332


} else {
ifres8331 = False


}

if True == ifres8331 {
if PrimFunc(symshen_4trim_1it) == __self {
__self1 = (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5783)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5783
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4trim_1it), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5783)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5783
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())
return


} else {
__e.Return(V5783)
return
}


}, 1)

tmp8337 := Call(__e, ns2_1set, symshen_4trim_1it, tmp8328)


_ = tmp8337

tmp8338 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__selftop:
V5802 := __self1
_ = V5802
V5803 := __self2
_ = V5803
V5804 := __self3
_ = V5804
tmp8457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5802)
}
__typedArg0 := V5802
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8426 Obj

if True == tmp8457 {
tmp8455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5802)
}
__typedArg0 := V5802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8455)
}
__typedArg0 := Nil
__typedArg1 := tmp8455
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8428 Obj

if True == tmp8456 {
tmp8454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8430 Obj

if True == tmp8454 {
tmp8452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8453 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8452)


var ifres8432 Obj

if True == tmp8453 {
tmp8449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8450 := Call(__e, PrimFunc(symhdstr), tmp8449)


tmp8451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("!"), tmp8450)
}
__typedArg0 := MakeString("!")
__typedArg1 := tmp8450
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8434 Obj

if True == tmp8451 {
tmp8446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8448 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8446)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8446
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres8436 Obj

if True == tmp8448 {
tmp8442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8444 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8442)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8442
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp8445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("!"), tmp8444)
}
__typedArg0 := MakeString("!")
__typedArg1 := tmp8444
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8438 Obj

if True == tmp8445 {
tmp8440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8440)
}
__typedArg0 := tmp8440
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8439 Obj

if True == tmp8441 {
ifres8439 = True


} else {
ifres8439 = False


}

ifres8438 = ifres8439


} else {
ifres8438 = False


}

var ifres8437 Obj

if True == ifres8438 {
ifres8437 = True


} else {
ifres8437 = False


}

ifres8436 = ifres8437


} else {
ifres8436 = False


}

var ifres8435 Obj

if True == ifres8436 {
ifres8435 = True


} else {
ifres8435 = False


}

ifres8434 = ifres8435


} else {
ifres8434 = False


}

var ifres8433 Obj

if True == ifres8434 {
ifres8433 = True


} else {
ifres8433 = False


}

ifres8432 = ifres8433


} else {
ifres8432 = False


}

var ifres8431 Obj

if True == ifres8432 {
ifres8431 = True


} else {
ifres8431 = False


}

ifres8430 = ifres8431


} else {
ifres8430 = False


}

var ifres8429 Obj

if True == ifres8430 {
ifres8429 = True


} else {
ifres8429 = False


}

ifres8428 = ifres8429


} else {
ifres8428 = False


}

var ifres8427 Obj

if True == ifres8428 {
ifres8427 = True


} else {
ifres8427 = False


}

ifres8426 = ifres8427


} else {
ifres8426 = False


}

if True == ifres8426 {
tmp8339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8339)
}
__typedArg0 := tmp8339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8341 := Call(__e, PrimFunc(symread_1from_1string), tmp8340)


let__8196 := tmp8341
_ = let__8196

tmp8342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8342)
}
__typedArg0 := tmp8342
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8343, tmp8344)
}
__typedArg0 := tmp8343
__typedArg1 := tmp8344
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dhistory_d, tmp8345)
}
__typedArg0 := symshen_4_dhistory_d
__typedArg1 := tmp8345
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__8197 := tmp8346
_ = let__8197

tmp8347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8347)
}
__typedArg0 := tmp8347
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8349 := Call(__e, PrimFunc(symshen_4app), tmp8348, MakeString("\n"), symshen_4a)


tmp8350 := Call(__e, PrimFunc(symstoutput))


tmp8351 := Call(__e, PrimFunc(sympr), tmp8349, tmp8350)


let__8198 := tmp8351
_ = let__8198

if PrimFunc(symshen_4evaluate_1lineread) == __self {
__self1, __self2, __self3 = let__8196, let__8197, V5804
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), let__8196, let__8197, V5804)
return


} else {
tmp8424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5802)
}
__typedArg0 := V5802
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8408 Obj

if True == tmp8424 {
tmp8422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5802)
}
__typedArg0 := V5802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8422)
}
__typedArg0 := Nil
__typedArg1 := tmp8422
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8410 Obj

if True == tmp8423 {
tmp8421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8412 Obj

if True == tmp8421 {
tmp8419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8420 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8419)


var ifres8414 Obj

if True == tmp8420 {
tmp8416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8417 := Call(__e, PrimFunc(symhdstr), tmp8416)


tmp8418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("!"), tmp8417)
}
__typedArg0 := MakeString("!")
__typedArg1 := tmp8417
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8415 Obj

if True == tmp8418 {
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

if True == ifres8408 {
tmp8357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8357)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8357
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})(), MakeString(""))
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8357)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8357
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8352 Obj

if True == tmp8359 {
ifres8352 = Nil


} else {
tmp8353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8355 := Call(__e, PrimFunc(symread_1from_1string), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8353)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8353
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp8356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8355)
}
__typedArg0 := tmp8355
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

ifres8352 = tmp8356


}

let__8199 := ifres8352
_ = let__8199

tmp8360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8360)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8360
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

tmp8362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8363 := Call(__e, PrimFunc(symshen_4use_1history), let__8199, tmp8361, tmp8362)


let__8200 := tmp8363
_ = let__8200

tmp8364 := Call(__e, PrimFunc(symshen_4app), let__8200, MakeString("\n"), symshen_4a)


tmp8365 := Call(__e, PrimFunc(symstoutput))


tmp8366 := Call(__e, PrimFunc(sympr), tmp8364, tmp8365)


let__8201 := tmp8366
_ = let__8201

tmp8367 := Call(__e, PrimFunc(symread_1from_1string), let__8200)


let__8202 := tmp8367
_ = let__8202

tmp8368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__8200, tmp8368)
}
__typedArg0 := let__8200
__typedArg1 := tmp8368
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dhistory_d, tmp8369)
}
__typedArg0 := symshen_4_dhistory_d
__typedArg1 := tmp8369
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__8203 := tmp8370
_ = let__8203

if PrimFunc(symshen_4evaluate_1lineread) == __self {
__self1, __self2, __self3 = let__8202, let__8203, V5804
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), let__8202, let__8203, V5804)
return


} else {
tmp8406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5802)
}
__typedArg0 := V5802
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8390 Obj

if True == tmp8406 {
tmp8404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5802)
}
__typedArg0 := V5802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8404)
}
__typedArg0 := Nil
__typedArg1 := tmp8404
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8392 Obj

if True == tmp8405 {
tmp8403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8394 Obj

if True == tmp8403 {
tmp8401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8402 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8401)


var ifres8396 Obj

if True == tmp8402 {
tmp8398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8399 := Call(__e, PrimFunc(symhdstr), tmp8398)


tmp8400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("%"), tmp8399)
}
__typedArg0 := MakeString("%")
__typedArg1 := tmp8399
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8397 Obj

if True == tmp8400 {
ifres8397 = True


} else {
ifres8397 = False


}

ifres8396 = ifres8397


} else {
ifres8396 = False


}

var ifres8395 Obj

if True == ifres8396 {
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

if True == ifres8390 {
tmp8376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8376)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8376
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})(), MakeString(""))
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8376)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8376
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8371 Obj

if True == tmp8378 {
ifres8371 = Nil


} else {
tmp8372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8374 := Call(__e, PrimFunc(symread_1from_1string), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8372)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8372
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp8375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8374)
}
__typedArg0 := tmp8374
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

ifres8371 = tmp8375


}

let__8204 := ifres8371
_ = let__8204

tmp8379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(tmp8379)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := tmp8379
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

tmp8381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8382 := Call(__e, PrimFunc(symshen_4peek_1history), let__8204, tmp8380, tmp8381)


let__8205 := tmp8382
_ = let__8205

tmp8383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5803)
}
__typedArg0 := V5803
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dhistory_d, tmp8383)
}
__typedArg0 := symshen_4_dhistory_d
__typedArg1 := tmp8383
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__8206 := tmp8384
_ = let__8206

__e.TailApply(PrimFunc(symabort))
return


} else {
tmp8388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, V5804)
}
__typedArg0 := True
__typedArg1 := V5804
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8388 {
__e.TailApply(PrimFunc(symshen_4check_1eval_1and_1print), V5802)
return
} else {
tmp8386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, V5804)
}
__typedArg0 := False
__typedArg1 := V5804
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8386 {
__e.TailApply(PrimFunc(symshen_4eval_1and_1print), V5802)
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

tmp8458 := Call(__e, ns2_1set, symshen_4evaluate_1lineread, tmp8338)


_ = tmp8458

tmp8459 := MakeNative(func(__e *ControlFlow) {
V5816 := __e.Get(1)
_ = V5816
V5817 := __e.Get(2)
_ = V5817
V5818 := __e.Get(3)
_ = V5818
tmp8465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V5816)
}
__typedArg0 := V5816
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})()

if True == tmp8465 {
tmp8460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(V5816)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := V5816
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()

tmp8461 := Call(__e, PrimFunc(symreverse), V5818)


__e.TailApply(PrimFunc(symnth), tmp8460, tmp8461)
return


} else {
tmp8463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V5816)
}
__typedArg0 := V5816
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp8463 {
__e.TailApply(PrimFunc(symshen_4string_1match), V5817, V5818)
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

tmp8466 := Call(__e, ns2_1set, symshen_4use_1history, tmp8459)


_ = tmp8466

tmp8467 := MakeNative(func(__e *ControlFlow) {
V5819 := __e.Get(1)
_ = V5819
V5820 := __e.Get(2)
_ = V5820
V5821 := __e.Get(3)
_ = V5821
tmp8481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V5819)
}
__typedArg0 := V5819
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})()

if True == tmp8481 {
tmp8468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(V5819)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := V5819
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()

tmp8469 := Call(__e, PrimFunc(symreverse), V5821)


tmp8470 := Call(__e, PrimFunc(symnth), tmp8468, tmp8469)


tmp8471 := Call(__e, PrimFunc(symshen_4app), tmp8470, MakeString(""), symshen_4a)


tmp8472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\n"))
__typedS1, __typedOK1 := TypedString(tmp8471)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\n")
__typedArg1 := tmp8471
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp8473 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8472, tmp8473)
return


} else {
tmp8479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V5820, MakeString(""))
}
__typedArg0 := V5820
__typedArg1 := MakeString("")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8476 Obj

if True == tmp8479 {
ifres8476 = True


} else {
tmp8478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V5819)
}
__typedArg0 := V5819
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

var ifres8477 Obj

if True == tmp8478 {
ifres8477 = True


} else {
ifres8477 = False


}

ifres8476 = ifres8477


}

if True == ifres8476 {
tmp8474 := Call(__e, PrimFunc(symreverse), V5821)


__e.TailApply(PrimFunc(symshen_4recursive_1string_1match), MakeInteger(0), V5820, tmp8474)
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

tmp8482 := Call(__e, ns2_1set, symshen_4peek_1history, tmp8467)


_ = tmp8482

tmp8483 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V5831 := __self1
_ = V5831
V5832 := __self2
_ = V5832
tmp8494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5832)
}
__typedArg0 := Nil
__typedArg1 := V5832
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8494 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("\ninput not found"))
}
__typedArg0 := MakeString("\ninput not found")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp8492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5832)
}
__typedArg0 := V5832
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8488 Obj

if True == tmp8492 {
tmp8490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5832)
}
__typedArg0 := V5832
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8491 := Call(__e, PrimFunc(symshen_4string_1prefix_2), V5831, tmp8490)


var ifres8489 Obj

if True == tmp8491 {
ifres8489 = True


} else {
ifres8489 = False


}

ifres8488 = ifres8489


} else {
ifres8488 = False


}

if True == ifres8488 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5832)
}
__typedArg0 := V5832
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp8486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5832)
}
__typedArg0 := V5832
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8486 {
tmp8484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5832)
}
__typedArg0 := V5832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4string_1match) == __self {
__self1, __self2 = V5831, tmp8484
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4string_1match), V5831, tmp8484)
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

tmp8495 := Call(__e, ns2_1set, symshen_4string_1match, tmp8483)


_ = tmp8495

tmp8496 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V5840 := __self1
_ = V5840
V5841 := __self2
_ = V5841
tmp8533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V5840)
}
__typedArg0 := MakeString("")
__typedArg1 := V5840
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8533 {
__e.Return(True)
return
} else {
tmp8531 := Call(__e, PrimFunc(symshen_4_7string_2), V5840)


var ifres8526 Obj

if True == tmp8531 {
tmp8528 := Call(__e, PrimFunc(symhdstr), V5840)


tmp8529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp8528)
}
__typedArg0 := tmp8528
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp8530 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp8529)


var ifres8527 Obj

if True == tmp8530 {
ifres8527 = True


} else {
ifres8527 = False


}

ifres8526 = ifres8527


} else {
ifres8526 = False


}

if True == ifres8526 {
if PrimFunc(symshen_4string_1prefix_2) == __self {
__self1, __self2 = (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5840)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5840
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})(), V5841
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4string_1prefix_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5840)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5840
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})(), V5841)
return


} else {
tmp8524 := Call(__e, PrimFunc(symshen_4_7string_2), V5841)


var ifres8519 Obj

if True == tmp8524 {
tmp8521 := Call(__e, PrimFunc(symhdstr), V5841)


tmp8522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp8521)
}
__typedArg0 := tmp8521
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp8523 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp8522)


var ifres8520 Obj

if True == tmp8523 {
ifres8520 = True


} else {
ifres8520 = False


}

ifres8519 = ifres8520


} else {
ifres8519 = False


}

if True == ifres8519 {
if PrimFunc(symshen_4string_1prefix_2) == __self {
__self1, __self2 = V5840, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5841)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5841
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4string_1prefix_2), V5840, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5841)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5841
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())
return


} else {
tmp8517 := Call(__e, PrimFunc(symshen_4_7string_2), V5841)


var ifres8513 Obj

if True == tmp8517 {
tmp8515 := Call(__e, PrimFunc(symhdstr), V5841)


tmp8516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("("), tmp8515)
}
__typedArg0 := MakeString("(")
__typedArg1 := tmp8515
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8514 Obj

if True == tmp8516 {
ifres8514 = True


} else {
ifres8514 = False


}

ifres8513 = ifres8514


} else {
ifres8513 = False


}

if True == ifres8513 {
if PrimFunc(symshen_4string_1prefix_2) == __self {
__self1, __self2 = V5840, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5841)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5841
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4string_1prefix_2), V5840, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5841)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5841
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())
return


} else {
tmp8511 := Call(__e, PrimFunc(symshen_4_7string_2), V5840)


var ifres8503 Obj

if True == tmp8511 {
tmp8510 := Call(__e, PrimFunc(symshen_4_7string_2), V5841)


var ifres8505 Obj

if True == tmp8510 {
tmp8507 := Call(__e, PrimFunc(symhdstr), V5840)


tmp8508 := Call(__e, PrimFunc(symhdstr), V5841)


tmp8509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp8507, tmp8508)
}
__typedArg0 := tmp8507
__typedArg1 := tmp8508
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8506 Obj

if True == tmp8509 {
ifres8506 = True


} else {
ifres8506 = False


}

ifres8505 = ifres8506


} else {
ifres8505 = False


}

var ifres8504 Obj

if True == ifres8505 {
ifres8504 = True


} else {
ifres8504 = False


}

ifres8503 = ifres8504


} else {
ifres8503 = False


}

if True == ifres8503 {
tmp8500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5840)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5840
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

if PrimFunc(symshen_4string_1prefix_2) == __self {
__self1, __self2 = tmp8500, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5841)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5841
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4string_1prefix_2), tmp8500, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5841)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5841
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

tmp8534 := Call(__e, ns2_1set, symshen_4string_1prefix_2, tmp8496)


_ = tmp8534

tmp8535 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__selftop:
V5852 := __self1
_ = V5852
V5853 := __self2
_ = V5853
V5854 := __self3
_ = V5854
tmp8550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5854)
}
__typedArg0 := Nil
__typedArg1 := V5854
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8550 {
__e.Return(symshen_4skip)
return
} else {
tmp8548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5854)
}
__typedArg0 := V5854
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8548 {
tmp8543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5854)
}
__typedArg0 := V5854
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8544 := Call(__e, PrimFunc(symshen_4string_1prefix_2), V5853, tmp8543)


var ifres8536 Obj

if True == tmp8544 {
tmp8537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5854)
}
__typedArg0 := V5854
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8538 := Call(__e, PrimFunc(symshen_4app), tmp8537, MakeString("\n"), symshen_4a)


tmp8540 := Call(__e, PrimFunc(symshen_4app), V5852, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(". "))
__typedS1, __typedOK1 := TypedString(tmp8538)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(". ")
__typedArg1 := tmp8538
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp8541 := Call(__e, PrimFunc(symstoutput))


tmp8542 := Call(__e, PrimFunc(sympr), tmp8540, tmp8541)


ifres8536 = tmp8542


} else {
ifres8536 = symshen_4skip


}

_ = ifres8536

tmp8545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5852)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5852
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()

tmp8546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5854)
}
__typedArg0 := V5854
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4recursive_1string_1match) == __self {
__self1, __self2, __self3 = tmp8545, V5853, tmp8546
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4recursive_1string_1match), tmp8545, V5853, tmp8546)
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

__e.TailApply(ns2_1set, symshen_4recursive_1string_1match, tmp8535)
return




}, 0)

