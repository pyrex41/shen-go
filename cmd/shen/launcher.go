package main

import . "github.com/pyrex41/shen-go/kl"

var LauncherMain = MakeNative(func(__e *ControlFlow) {
_ = MakeString("Copyright (c) 2019 Bruno Deferrari.\nBSD 3-Clause License: http://opensource.org/licenses/BSD-3-Clause")

tmp20178 := MakeNative(func(__e *ControlFlow) {
V7102 := __e.Get(1)
_ = V7102
tmp20179 := Call(__e, PrimFunc(symread_1file), V7102)


let__20174 := tmp20179
_ = let__20174

tmp20180 := MakeNative(func(__e *ControlFlow) {
Z7104 := __e.Get(1)
_ = Z7104
__e.TailApply(PrimFunc(symeval), Z7104)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp20180, let__20174)
return


}, 1)

tmp20181 := Call(__e, ns2_1set, symshen_4x_4launcher_4quiet_1load, tmp20178)


_ = tmp20181

tmp20182 := MakeNative(func(__e *ControlFlow) {
tmp20183 := Call(__e, PrimFunc(symversion))


tmp20184 := Call(__e, PrimFunc(symlanguage))


tmp20185 := Call(__e, PrimFunc(symport))


tmp20186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20185, Nil)
}
__typedArg0 := tmp20185
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20184, tmp20186)
}
__typedArg0 := tmp20184
__typedArg1 := tmp20186
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20188 := Call(__e, PrimFunc(symimplementation))


tmp20189 := Call(__e, PrimFunc(symrelease))


tmp20190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20189, Nil)
}
__typedArg0 := tmp20189
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20188, tmp20190)
}
__typedArg0 := tmp20188
__typedArg1 := tmp20190
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20191, Nil)
}
__typedArg0 := tmp20191
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symimplementation, tmp20192)
}
__typedArg0 := symimplementation
__typedArg1 := tmp20192
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20187, tmp20193)
}
__typedArg0 := tmp20187
__typedArg1 := tmp20193
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symport, tmp20194)
}
__typedArg0 := symport
__typedArg1 := tmp20194
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20196 := Call(__e, PrimFunc(symshen_4app), tmp20195, MakeString("\n"), symshen_4r)


__e.TailApply(PrimFunc(symshen_4app), tmp20183, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" "))
__typedS1, __typedOK1 := TypedString(tmp20196)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" ")
__typedArg1 := tmp20196
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)
return


}, 0)

tmp20198 := Call(__e, ns2_1set, symshen_4x_4launcher_4version_1string, tmp20182)


_ = tmp20198

tmp20199 := MakeNative(func(__e *ControlFlow) {
V7105 := __e.Get(1)
_ = V7105
tmp20200 := Call(__e, PrimFunc(symshen_4app), V7105, MakeString(" [--version] [--help] <COMMAND> [<ARGS>]\n\ncommands:\n    repl\n        Launches the interactive REPL.\n        Default action if no command is supplied.\n\n    script <FILE> [<ARGS>]\n        Runs the script in FILE. *argv* is set to [FILE | ARGS].\n\n    eval <ARGS>\n        Evaluates expressions and files. ARGS are evaluated from\n        left to right and can be a combination of:\n            -e, --eval <EXPR>\n                Evaluates EXPR and prints result.\n            -l, --load <FILE>\n                Reads and evaluates FILE.\n            -q, --quiet\n                Silences interactive output.\n            -s, --set <KEY> <VALUE>\n                Evaluates KEY, VALUE and sets as global.\n            -r, --repl\n                Launches the interactive REPL after evaluating\n                all the previous expresions."), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Usage: "))
__typedS1, __typedOK1 := TypedString(tmp20200)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Usage: ")
__typedArg1 := tmp20200
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp20201 := Call(__e, ns2_1set, symshen_4x_4launcher_4help_1text, tmp20199)


_ = tmp20201

tmp20202 := MakeNative(func(__e *ControlFlow) {
V7106 := __e.Get(1)
_ = V7106
tmp20209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V7106)
}
__typedArg0 := Nil
__typedArg1 := V7106
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp20209 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsuccess, Nil)
}
__typedArg0 := symsuccess
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
} else {
tmp20207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7106)
}
__typedArg0 := V7106
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp20207 {
tmp20203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7106)
}
__typedArg0 := V7106
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20204 := Call(__e, PrimFunc(symthaw), tmp20203)


_ = tmp20204

tmp20205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7106)
}
__typedArg0 := V7106
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4x_4launcher_4execute_1all), tmp20205)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4x_4launcher_4execute_1all)
return
}


}


}, 1)

tmp20210 := Call(__e, ns2_1set, symshen_4x_4launcher_4execute_1all, tmp20202)


_ = tmp20210

tmp20211 := MakeNative(func(__e *ControlFlow) {
V7107 := __e.Get(1)
_ = V7107
tmp20212 := Call(__e, PrimFunc(symread_1from_1string), V7107)


tmp20213 := Call(__e, PrimFunc(symhead), tmp20212)


__e.TailApply(PrimFunc(symeval), tmp20213)
return


}, 1)

tmp20214 := Call(__e, ns2_1set, symshen_4x_4launcher_4eval_1string, tmp20211)


_ = tmp20214

tmp20215 := MakeNative(func(__e *ControlFlow) {
V7110 := __e.Get(1)
_ = V7110
tmp20225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("-e"), V7110)
}
__typedArg0 := MakeString("-e")
__typedArg1 := V7110
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp20225 {
__e.Return(MakeString("--eval"))
return
} else {
tmp20223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("-l"), V7110)
}
__typedArg0 := MakeString("-l")
__typedArg1 := V7110
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp20223 {
__e.Return(MakeString("--load"))
return
} else {
tmp20221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("-q"), V7110)
}
__typedArg0 := MakeString("-q")
__typedArg1 := V7110
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp20221 {
__e.Return(MakeString("--quiet"))
return
} else {
tmp20219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("-s"), V7110)
}
__typedArg0 := MakeString("-s")
__typedArg1 := V7110
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp20219 {
__e.Return(MakeString("--set"))
return
} else {
tmp20217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("-r"), V7110)
}
__typedArg0 := MakeString("-r")
__typedArg1 := V7110
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp20217 {
__e.Return(MakeString("--repl"))
return
} else {
__e.Return(False)
return
}


}


}


}


}


}, 1)

tmp20226 := Call(__e, ns2_1set, symshen_4x_4launcher_4eval_1flag_1map, tmp20215)


_ = tmp20226

tmp20227 := MakeNative(func(__e *ControlFlow) {
V7115 := __e.Get(1)
_ = V7115
V7116 := __e.Get(2)
_ = V7116
tmp20329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V7115)
}
__typedArg0 := Nil
__typedArg1 := V7115
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp20329 {
tmp20228 := Call(__e, PrimFunc(symreverse), V7116)


__e.TailApply(PrimFunc(symshen_4x_4launcher_4execute_1all), tmp20228)
return


} else {
tmp20327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20319 Obj

if True == tmp20327 {
tmp20325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--eval"), tmp20325)
}
__typedArg0 := MakeString("--eval")
__typedArg1 := tmp20325
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20321 Obj

if True == tmp20326 {
tmp20323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20323)
}
__typedArg0 := tmp20323
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20322 Obj

if True == tmp20324 {
ifres20322 = True


} else {
ifres20322 = False


}

ifres20321 = ifres20322


} else {
ifres20321 = False


}

var ifres20320 Obj

if True == ifres20321 {
ifres20320 = True


} else {
ifres20320 = False


}

ifres20319 = ifres20320


} else {
ifres20319 = False


}

if True == ifres20319 {
tmp20229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp20229)
}
__typedArg0 := tmp20229
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20231 := MakeNative(func(__e *ControlFlow) {
tmp20232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20232)
}
__typedArg0 := tmp20232
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20234 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1string), tmp20233)


tmp20235 := Call(__e, PrimFunc(symshen_4app), tmp20234, MakeString("\n"), symshen_4a)


tmp20236 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp20235, tmp20236)
return


}, 0)

tmp20237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20231, V7116)
}
__typedArg0 := tmp20231
__typedArg1 := V7116
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp20230, tmp20237)
return


} else {
tmp20317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20309 Obj

if True == tmp20317 {
tmp20315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--load"), tmp20315)
}
__typedArg0 := MakeString("--load")
__typedArg1 := tmp20315
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20311 Obj

if True == tmp20316 {
tmp20313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20313)
}
__typedArg0 := tmp20313
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20312 Obj

if True == tmp20314 {
ifres20312 = True


} else {
ifres20312 = False


}

ifres20311 = ifres20312


} else {
ifres20311 = False


}

var ifres20310 Obj

if True == ifres20311 {
ifres20310 = True


} else {
ifres20310 = False


}

ifres20309 = ifres20310


} else {
ifres20309 = False


}

if True == ifres20309 {
tmp20238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp20238)
}
__typedArg0 := tmp20238
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20240 := MakeNative(func(__e *ControlFlow) {
tmp20241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20241)
}
__typedArg0 := tmp20241
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symload), tmp20242)
return


}, 0)

tmp20243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20240, V7116)
}
__typedArg0 := tmp20240
__typedArg1 := V7116
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp20239, tmp20243)
return


} else {
tmp20307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20303 Obj

if True == tmp20307 {
tmp20305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--quiet"), tmp20305)
}
__typedArg0 := MakeString("--quiet")
__typedArg1 := tmp20305
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20304 Obj

if True == tmp20306 {
ifres20304 = True


} else {
ifres20304 = False


}

ifres20303 = ifres20304


} else {
ifres20303 = False


}

if True == ifres20303 {
tmp20244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20245 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dhush_d, True)
}
__typedArg0 := sym_dhush_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
}, 0)

tmp20246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20245, V7116)
}
__typedArg0 := tmp20245
__typedArg1 := V7116
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp20244, tmp20246)
return


} else {
tmp20301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20288 Obj

if True == tmp20301 {
tmp20299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--set"), tmp20299)
}
__typedArg0 := MakeString("--set")
__typedArg1 := tmp20299
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20290 Obj

if True == tmp20300 {
tmp20297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20297)
}
__typedArg0 := tmp20297
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20292 Obj

if True == tmp20298 {
tmp20294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp20294)
}
__typedArg0 := tmp20294
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20295)
}
__typedArg0 := tmp20295
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20293 Obj

if True == tmp20296 {
ifres20293 = True


} else {
ifres20293 = False


}

ifres20292 = ifres20293


} else {
ifres20292 = False


}

var ifres20291 Obj

if True == ifres20292 {
ifres20291 = True


} else {
ifres20291 = False


}

ifres20290 = ifres20291


} else {
ifres20290 = False


}

var ifres20289 Obj

if True == ifres20290 {
ifres20289 = True


} else {
ifres20289 = False


}

ifres20288 = ifres20289


} else {
ifres20288 = False


}

if True == ifres20288 {
tmp20247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp20247)
}
__typedArg0 := tmp20247
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp20248)
}
__typedArg0 := tmp20248
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20250 := MakeNative(func(__e *ControlFlow) {
tmp20251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20251)
}
__typedArg0 := tmp20251
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20253 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1string), tmp20252)


tmp20254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp20254)
}
__typedArg0 := tmp20254
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20255)
}
__typedArg0 := tmp20255
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20257 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1string), tmp20256)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(tmp20253, tmp20257)
}
__typedArg0 := tmp20253
__typedArg1 := tmp20257
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp20258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20250, V7116)
}
__typedArg0 := tmp20250
__typedArg1 := V7116
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp20249, tmp20258)
return


} else {
tmp20286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20282 Obj

if True == tmp20286 {
tmp20284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--repl"), tmp20284)
}
__typedArg0 := MakeString("--repl")
__typedArg1 := tmp20284
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20283 Obj

if True == tmp20285 {
ifres20283 = True


} else {
ifres20283 = False


}

ifres20282 = ifres20283


} else {
ifres20282 = False


}

if True == ifres20282 {
tmp20259 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1command_1h), Nil, V7116)


_ = tmp20259

tmp20260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlaunch_1repl, tmp20260)
}
__typedArg0 := symlaunch_1repl
__typedArg1 := tmp20260
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp20261 := MakeNative(func(__e *ControlFlow) {
tmp20267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp20267 {
tmp20262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20263 := Call(__e, PrimFunc(symshen_4app), tmp20262, MakeString(""), symshen_4a)


tmp20265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Invalid eval argument: "))
__typedS1, __typedOK1 := TypedString(tmp20263)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Invalid eval argument: ")
__typedArg1 := tmp20263
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), Nil)
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Invalid eval argument: "))
__typedS1, __typedOK1 := TypedString(tmp20263)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Invalid eval argument: ")
__typedArg1 := tmp20263
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symerror, tmp20265)
}
__typedArg0 := symerror
__typedArg1 := tmp20265
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4x_4launcher_4eval_1command_1h)
return
}


}, 0)

let__20175 := tmp20261
_ = let__20175

tmp20280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp20280 {
tmp20268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20269 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1flag_1map), tmp20268)


let__20177 := tmp20269
_ = let__20177

tmp20275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, let__20177)
}
__typedArg0 := False
__typedArg1 := let__20177
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20270 Obj

if True == tmp20275 {
tmp20271 := Call(__e, PrimFunc(symfail))


ifres20270 = tmp20271


} else {
tmp20272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__20177, tmp20272)
}
__typedArg0 := let__20177
__typedArg1 := tmp20272
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20274 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp20273, V7116)


ifres20270 = tmp20274


}

let__20176 := ifres20270
_ = let__20176

tmp20277 := Call(__e, PrimFunc(symfail))


tmp20278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__20176, tmp20277)
}
__typedArg0 := let__20176
__typedArg1 := tmp20277
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp20278 {
__e.TailApply(PrimFunc(symthaw), let__20175)
return
} else {
__e.Return(let__20176)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__20175)
return
}


}


}


}


}


}


}


}, 2)

tmp20330 := Call(__e, ns2_1set, symshen_4x_4launcher_4eval_1command_1h, tmp20227)


_ = tmp20330

tmp20331 := MakeNative(func(__e *ControlFlow) {
V7120 := __e.Get(1)
_ = V7120
__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), V7120, Nil)
return
}, 1)

tmp20332 := Call(__e, ns2_1set, symshen_4x_4launcher_4eval_1command, tmp20331)


_ = tmp20332

tmp20333 := MakeNative(func(__e *ControlFlow) {
V7121 := __e.Get(1)
_ = V7121
V7122 := __e.Get(2)
_ = V7122
tmp20334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V7121, V7122)
}
__typedArg0 := V7121
__typedArg1 := V7122
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp20335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dargv_d, tmp20334)
}
__typedArg0 := sym_dargv_d
__typedArg1 := tmp20334
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp20335

tmp20336 := Call(__e, PrimFunc(symshen_4x_4launcher_4quiet_1load), V7121)


_ = tmp20336

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsuccess, Nil)
}
__typedArg0 := symsuccess
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp20337 := Call(__e, ns2_1set, symshen_4x_4launcher_4script_1command, tmp20333)


_ = tmp20337

tmp20338 := MakeNative(func(__e *ControlFlow) {
V7123 := __e.Get(1)
_ = V7123
tmp20425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20421 Obj

if True == tmp20425 {
tmp20423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp20423)
}
__typedArg0 := Nil
__typedArg1 := tmp20423
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20422 Obj

if True == tmp20424 {
ifres20422 = True


} else {
ifres20422 = False


}

ifres20421 = ifres20422


} else {
ifres20421 = False


}

if True == ifres20421 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlaunch_1repl, Nil)
}
__typedArg0 := symlaunch_1repl
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
} else {
tmp20419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20410 Obj

if True == tmp20419 {
tmp20417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20417)
}
__typedArg0 := tmp20417
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20412 Obj

if True == tmp20418 {
tmp20414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20414)
}
__typedArg0 := tmp20414
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--help"), tmp20415)
}
__typedArg0 := MakeString("--help")
__typedArg1 := tmp20415
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20413 Obj

if True == tmp20416 {
ifres20413 = True


} else {
ifres20413 = False


}

ifres20412 = ifres20413


} else {
ifres20412 = False


}

var ifres20411 Obj

if True == ifres20412 {
ifres20411 = True


} else {
ifres20411 = False


}

ifres20410 = ifres20411


} else {
ifres20410 = False


}

if True == ifres20410 {
tmp20339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20340 := Call(__e, PrimFunc(symshen_4x_4launcher_4help_1text), tmp20339)


tmp20341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20340, Nil)
}
__typedArg0 := tmp20340
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshow_1help, tmp20341)
}
__typedArg0 := symshow_1help
__typedArg1 := tmp20341
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp20408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20399 Obj

if True == tmp20408 {
tmp20406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20406)
}
__typedArg0 := tmp20406
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20401 Obj

if True == tmp20407 {
tmp20403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20403)
}
__typedArg0 := tmp20403
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--version"), tmp20404)
}
__typedArg0 := MakeString("--version")
__typedArg1 := tmp20404
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20402 Obj

if True == tmp20405 {
ifres20402 = True


} else {
ifres20402 = False


}

ifres20401 = ifres20402


} else {
ifres20401 = False


}

var ifres20400 Obj

if True == ifres20401 {
ifres20400 = True


} else {
ifres20400 = False


}

ifres20399 = ifres20400


} else {
ifres20399 = False


}

if True == ifres20399 {
tmp20342 := Call(__e, PrimFunc(symshen_4x_4launcher_4version_1string))


tmp20343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp20342, Nil)
}
__typedArg0 := tmp20342
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsuccess, tmp20343)
}
__typedArg0 := symsuccess
__typedArg1 := tmp20343
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp20397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20388 Obj

if True == tmp20397 {
tmp20395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20395)
}
__typedArg0 := tmp20395
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20390 Obj

if True == tmp20396 {
tmp20392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20392)
}
__typedArg0 := tmp20392
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("repl"), tmp20393)
}
__typedArg0 := MakeString("repl")
__typedArg1 := tmp20393
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20391 Obj

if True == tmp20394 {
ifres20391 = True


} else {
ifres20391 = False


}

ifres20390 = ifres20391


} else {
ifres20390 = False


}

var ifres20389 Obj

if True == ifres20390 {
ifres20389 = True


} else {
ifres20389 = False


}

ifres20388 = ifres20389


} else {
ifres20388 = False


}

if True == ifres20388 {
tmp20344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp20344)
}
__typedArg0 := tmp20344
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlaunch_1repl, tmp20345)
}
__typedArg0 := symlaunch_1repl
__typedArg1 := tmp20345
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp20386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20372 Obj

if True == tmp20386 {
tmp20384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20384)
}
__typedArg0 := tmp20384
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20374 Obj

if True == tmp20385 {
tmp20381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20381)
}
__typedArg0 := tmp20381
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("script"), tmp20382)
}
__typedArg0 := MakeString("script")
__typedArg1 := tmp20382
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20376 Obj

if True == tmp20383 {
tmp20378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp20378)
}
__typedArg0 := tmp20378
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20379)
}
__typedArg0 := tmp20379
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20377 Obj

if True == tmp20380 {
ifres20377 = True


} else {
ifres20377 = False


}

ifres20376 = ifres20377


} else {
ifres20376 = False


}

var ifres20375 Obj

if True == ifres20376 {
ifres20375 = True


} else {
ifres20375 = False


}

ifres20374 = ifres20375


} else {
ifres20374 = False


}

var ifres20373 Obj

if True == ifres20374 {
ifres20373 = True


} else {
ifres20373 = False


}

ifres20372 = ifres20373


} else {
ifres20372 = False


}

if True == ifres20372 {
tmp20346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp20346)
}
__typedArg0 := tmp20346
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20347)
}
__typedArg0 := tmp20347
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp20349)
}
__typedArg0 := tmp20349
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp20350)
}
__typedArg0 := tmp20350
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4x_4launcher_4script_1command), tmp20348, tmp20351)
return


} else {
tmp20370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20361 Obj

if True == tmp20370 {
tmp20368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20368)
}
__typedArg0 := tmp20368
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20363 Obj

if True == tmp20369 {
tmp20365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20365)
}
__typedArg0 := tmp20365
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("eval"), tmp20366)
}
__typedArg0 := MakeString("eval")
__typedArg1 := tmp20366
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20364 Obj

if True == tmp20367 {
ifres20364 = True


} else {
ifres20364 = False


}

ifres20363 = ifres20364


} else {
ifres20363 = False


}

var ifres20362 Obj

if True == ifres20363 {
ifres20362 = True


} else {
ifres20362 = False


}

ifres20361 = ifres20362


} else {
ifres20361 = False


}

if True == ifres20361 {
tmp20352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp20352)
}
__typedArg0 := tmp20352
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command), tmp20353)
return


} else {
tmp20359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20355 Obj

if True == tmp20359 {
tmp20357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20357)
}
__typedArg0 := tmp20357
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20356 Obj

if True == tmp20358 {
ifres20356 = True


} else {
ifres20356 = False


}

ifres20355 = ifres20356


} else {
ifres20355 = False


}

if True == ifres20355 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunknown_1arguments, V7123)
}
__typedArg0 := symunknown_1arguments
__typedArg1 := V7123
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4x_4launcher_4launch_1shen)
return
}


}


}


}


}


}


}


}, 1)

tmp20426 := Call(__e, ns2_1set, symshen_4x_4launcher_4launch_1shen, tmp20338)


_ = tmp20426

tmp20427 := MakeNative(func(__e *ControlFlow) {
V7126 := __e.Get(1)
_ = V7126
tmp20526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20518 Obj

if True == tmp20526 {
tmp20524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsuccess, tmp20524)
}
__typedArg0 := symsuccess
__typedArg1 := tmp20524
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20520 Obj

if True == tmp20525 {
tmp20522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp20522)
}
__typedArg0 := Nil
__typedArg1 := tmp20522
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20521 Obj

if True == tmp20523 {
ifres20521 = True


} else {
ifres20521 = False


}

ifres20520 = ifres20521


} else {
ifres20520 = False


}

var ifres20519 Obj

if True == ifres20520 {
ifres20519 = True


} else {
ifres20519 = False


}

ifres20518 = ifres20519


} else {
ifres20518 = False


}

if True == ifres20518 {
__e.Return(symshen_4x_4launcher_4done)
return
} else {
tmp20516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20503 Obj

if True == tmp20516 {
tmp20514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsuccess, tmp20514)
}
__typedArg0 := symsuccess
__typedArg1 := tmp20514
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20505 Obj

if True == tmp20515 {
tmp20512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20512)
}
__typedArg0 := tmp20512
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20507 Obj

if True == tmp20513 {
tmp20509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp20509)
}
__typedArg0 := tmp20509
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp20510)
}
__typedArg0 := Nil
__typedArg1 := tmp20510
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20508 Obj

if True == tmp20511 {
ifres20508 = True


} else {
ifres20508 = False


}

ifres20507 = ifres20508


} else {
ifres20507 = False


}

var ifres20506 Obj

if True == ifres20507 {
ifres20506 = True


} else {
ifres20506 = False


}

ifres20505 = ifres20506


} else {
ifres20505 = False


}

var ifres20504 Obj

if True == ifres20505 {
ifres20504 = True


} else {
ifres20504 = False


}

ifres20503 = ifres20504


} else {
ifres20503 = False


}

if True == ifres20503 {
tmp20428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20428)
}
__typedArg0 := tmp20428
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20430 := Call(__e, PrimFunc(symshen_4app), tmp20429, MakeString("\n"), symshen_4a)


tmp20431 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp20430, tmp20431)
return


} else {
tmp20501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20488 Obj

if True == tmp20501 {
tmp20499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symerror, tmp20499)
}
__typedArg0 := symerror
__typedArg1 := tmp20499
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20490 Obj

if True == tmp20500 {
tmp20497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20497)
}
__typedArg0 := tmp20497
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20492 Obj

if True == tmp20498 {
tmp20494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp20494)
}
__typedArg0 := tmp20494
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp20495)
}
__typedArg0 := Nil
__typedArg1 := tmp20495
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20493 Obj

if True == tmp20496 {
ifres20493 = True


} else {
ifres20493 = False


}

ifres20492 = ifres20493


} else {
ifres20492 = False


}

var ifres20491 Obj

if True == ifres20492 {
ifres20491 = True


} else {
ifres20491 = False


}

ifres20490 = ifres20491


} else {
ifres20490 = False


}

var ifres20489 Obj

if True == ifres20490 {
ifres20489 = True


} else {
ifres20489 = False


}

ifres20488 = ifres20489


} else {
ifres20488 = False


}

if True == ifres20488 {
tmp20432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20432)
}
__typedArg0 := tmp20432
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20434 := Call(__e, PrimFunc(symshen_4app), tmp20433, MakeString("\n"), symshen_4a)


tmp20435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("ERROR: "))
__typedS1, __typedOK1 := TypedString(tmp20434)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("ERROR: ")
__typedArg1 := tmp20434
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp20436 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp20435, tmp20436)
return


} else {
tmp20486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20482 Obj

if True == tmp20486 {
tmp20484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlaunch_1repl, tmp20484)
}
__typedArg0 := symlaunch_1repl
__typedArg1 := tmp20484
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20483 Obj

if True == tmp20485 {
ifres20483 = True


} else {
ifres20483 = False


}

ifres20482 = ifres20483


} else {
ifres20482 = False


}

if True == ifres20482 {
__e.TailApply(PrimFunc(symshen_4repl))
return
} else {
tmp20480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20467 Obj

if True == tmp20480 {
tmp20478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshow_1help, tmp20478)
}
__typedArg0 := symshow_1help
__typedArg1 := tmp20478
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20469 Obj

if True == tmp20479 {
tmp20476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20476)
}
__typedArg0 := tmp20476
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20471 Obj

if True == tmp20477 {
tmp20473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp20473)
}
__typedArg0 := tmp20473
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp20474)
}
__typedArg0 := Nil
__typedArg1 := tmp20474
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20472 Obj

if True == tmp20475 {
ifres20472 = True


} else {
ifres20472 = False


}

ifres20471 = ifres20472


} else {
ifres20471 = False


}

var ifres20470 Obj

if True == ifres20471 {
ifres20470 = True


} else {
ifres20470 = False


}

ifres20469 = ifres20470


} else {
ifres20469 = False


}

var ifres20468 Obj

if True == ifres20469 {
ifres20468 = True


} else {
ifres20468 = False


}

ifres20467 = ifres20468


} else {
ifres20467 = False


}

if True == ifres20467 {
tmp20437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20437)
}
__typedArg0 := tmp20437
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20439 := Call(__e, PrimFunc(symshen_4app), tmp20438, MakeString("\n"), symshen_4a)


tmp20440 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp20439, tmp20440)
return


} else {
tmp20465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20452 Obj

if True == tmp20465 {
tmp20463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symunknown_1arguments, tmp20463)
}
__typedArg0 := symunknown_1arguments
__typedArg1 := tmp20463
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20454 Obj

if True == tmp20464 {
tmp20461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20461)
}
__typedArg0 := tmp20461
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20456 Obj

if True == tmp20462 {
tmp20458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp20458)
}
__typedArg0 := tmp20458
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20459)
}
__typedArg0 := tmp20459
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20457 Obj

if True == tmp20460 {
ifres20457 = True


} else {
ifres20457 = False


}

ifres20456 = ifres20457


} else {
ifres20456 = False


}

var ifres20455 Obj

if True == ifres20456 {
ifres20455 = True


} else {
ifres20455 = False


}

ifres20454 = ifres20455


} else {
ifres20454 = False


}

var ifres20453 Obj

if True == ifres20454 {
ifres20453 = True


} else {
ifres20453 = False


}

ifres20452 = ifres20453


} else {
ifres20452 = False


}

if True == ifres20452 {
tmp20441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp20441)
}
__typedArg0 := tmp20441
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20442)
}
__typedArg0 := tmp20442
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20444)
}
__typedArg0 := tmp20444
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20446 := Call(__e, PrimFunc(symshen_4app), tmp20445, MakeString(" --help' for more information.\n"), symshen_4a)


tmp20448 := Call(__e, PrimFunc(symshen_4app), tmp20443, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\nTry `"))
__typedS1, __typedOK1 := TypedString(tmp20446)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\nTry `")
__typedArg1 := tmp20446
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp20449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("ERROR: Invalid argument: "))
__typedS1, __typedOK1 := TypedString(tmp20448)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("ERROR: Invalid argument: ")
__typedArg1 := tmp20448
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp20450 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp20449, tmp20450)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4x_4launcher_4default_1handle_1result)
return
}


}


}


}


}


}


}, 1)

tmp20527 := Call(__e, ns2_1set, symshen_4x_4launcher_4default_1handle_1result, tmp20427)


_ = tmp20527

tmp20528 := MakeNative(func(__e *ControlFlow) {
V7127 := __e.Get(1)
_ = V7127
tmp20529 := Call(__e, PrimFunc(symshen_4x_4launcher_4launch_1shen), V7127)


__e.TailApply(PrimFunc(symshen_4x_4launcher_4default_1handle_1result), tmp20529)
return


}, 1)

__e.TailApply(ns2_1set, symshen_4x_4launcher_4main, tmp20528)
return




}, 0)

var sym_5_1vector = MakeSymbol("<-vector")
var symtuple_2 = MakeSymbol("tuple?")
var symstoutput = MakeSymbol("stoutput")
var symshen_4eos = MakeSymbol("shen.eos")
var symshen_4_5define_6 = MakeSymbol("shen.<define>")
var symshen_4input_1track = MakeSymbol("shen.input-track")
var symprint = MakeSymbol("print")
var symshen_4rfas_1h = MakeSymbol("shen.rfas-h")
var symshen_4_5single_6 = MakeSymbol("shen.<single>")
var symshen_4ccons_2 = MakeSymbol("shen.ccons?")
var symshen_4_dhistory_d = MakeSymbol("shen.*history*")
var symshen_4_dresidue_d = MakeSymbol("shen.*residue*")
var symshen_4unpackage_emacroexpand = MakeSymbol("shen.unpackage&macroexpand")
var symwhen = MakeSymbol("when")
var symshen_4toplevel_1forms = MakeSymbol("shen.toplevel-forms")
var symshen_4out_1of_1bounds = MakeSymbol("shen.out-of-bounds")
var symshen_4_5rules_6 = MakeSymbol("shen.<rules>")
var symshen_4factor = MakeSymbol("shen.factor")
var symshen_4_5integer_6 = MakeSymbol("shen.<integer>")
var symshen_4mu_1h = MakeSymbol("shen.mu-h")
var symshen_4construct_1context = MakeSymbol("shen.construct-context")
var symshen_4show = MakeSymbol("shen.show")
var symshen_4lookupsig = MakeSymbol("shen.lookupsig")
var symhd = MakeSymbol("hd")
var symshen_4_5alphanum_6 = MakeSymbol("shen.<alphanum>")
var symenable_1type_1theory = MakeSymbol("enable-type-theory")
var symshen_4prodbutzero = MakeSymbol("shen.prodbutzero")
var symread = MakeSymbol("read")
var symread_1file = MakeSymbol("read-file")
var symshen_4intern_1in_1package = MakeSymbol("shen.intern-in-package")
var symshen_4prterm = MakeSymbol("shen.prterm")
var sym_a_a = MakeSymbol("==")
var symcond = MakeSymbol("cond")
var symfreeze = MakeSymbol("freeze")
var symsimple_1error = MakeSymbol("simple-error")
var symshen_4_5clause_6 = MakeSymbol("shen.<clause>")
var symshen_4side_1conditions_1_6goals = MakeSymbol("shen.side-conditions->goals")
var syminternal = MakeSymbol("internal")
var symin_1package = MakeSymbol("in-package")
var symunspecialise = MakeSymbol("unspecialise")
var symsystem_1S_2 = MakeSymbol("system-S?")
var symshen_4pretty_1type = MakeSymbol("shen.pretty-type")
var symshen_4lzy_a = MakeSymbol("shen.lzy=")
var symshen_4terpri_1or_1read_1char = MakeSymbol("shen.terpri-or-read-char")
var symshen_4prtl = MakeSymbol("shen.prtl")
var symsymbol_2 = MakeSymbol("symbol?")
var symshen_4remove_1pointer = MakeSymbol("shen.remove-pointer")
var symshen_4read_1unit_1string = MakeSymbol("shen.read-unit-string")
var symshen_4pause_1for_1user = MakeSymbol("shen.pause-for-user")
var symshen_4check_1eval_1and_1print = MakeSymbol("shen.check-eval-and-print")
var symshen_4_5digit_6 = MakeSymbol("shen.<digit>")
var sym_6_6 = MakeSymbol(">>")
var sym_dargv_d = MakeSymbol("*argv*")
var symunabsolute = MakeSymbol("unabsolute")
var symshen_4parse_1failure_2 = MakeSymbol("shen.parse-failure?")
var symreceive = MakeSymbol("receive")
var symB = MakeSymbol("B")
var symshen_4bind_b = MakeSymbol("shen.bind!")
var symshen_4ticket_1number = MakeSymbol("shen.ticket-number")
var symshen_4_dstep_d = MakeSymbol("shen.*step*")
var symunix = MakeSymbol("unix")
var symshen_4trim_1it = MakeSymbol("shen.trim-it")
var symshen_4_dtracking_d = MakeSymbol("shen.*tracking*")
var symshen_4parameters = MakeSymbol("shen.parameters")
var symshen_4return_2 = MakeSymbol("shen.return?")
var symshen_4use_1history = MakeSymbol("shen.use-history")
var symshen_4bottom = MakeSymbol("shen.bottom")
var symshen_4rep_1X = MakeSymbol("shen.rep-X")
var symshen_4_5wildcard_6 = MakeSymbol("shen.<wildcard>")
var symshen_4type_1theory_1enabled_2 = MakeSymbol("shen.type-theory-enabled?")
var symabort = MakeSymbol("abort")
var symshen_4analyse_1symbol_2 = MakeSymbol("shen.analyse-symbol?")
var symport = MakeSymbol("port")
var symshen_4compile_1to_1kl = MakeSymbol("shen.compile-to-kl")
var symshen_4print_1freshterm = MakeSymbol("shen.print-freshterm")
var symshen_4kl_1body = MakeSymbol("shen.kl-body")
var symshen_4_5lowC_6 = MakeSymbol("shen.<lowC>")
var sym_dmacros_d = MakeSymbol("*macros*")
var symshen_4type_1error = MakeSymbol("shen.type-error")
var symshen_4_7m = MakeSymbol("shen.+m")
var symshen_4t_d_1correct = MakeSymbol("shen.t*-correct")
var symParse = MakeSymbol("Parse")
var symshen_4arity_1chk = MakeSymbol("shen.arity-chk")
var symshen_4_5pattern1_6 = MakeSymbol("shen.<pattern1>")
var symdo = MakeSymbol("do")
var symshen_4atom_1_6str = MakeSymbol("shen.atom->str")
var symshen_4constructor_2 = MakeSymbol("shen.constructor?")
var symshen_4_5rcurly_6 = MakeSymbol("shen.<rcurly>")
var symshen_4_dspy_d = MakeSymbol("shen.*spy*")
var symshen_4_5sng_6 = MakeSymbol("shen.<sng>")
var symshen_4_5packagechar_6 = MakeSymbol("shen.<packagechar>")
var symhead = MakeSymbol("head")
var symstep_2 = MakeSymbol("step?")
var symshen_4arg_1_6str = MakeSymbol("shen.arg->str")
var symand = MakeSymbol("and")
var symshen_4typename = MakeSymbol("shen.typename")
var symshen_4_5formulae_6 = MakeSymbol("shen.<formulae>")
var symshen_4monotype = MakeSymbol("shen.monotype")
var symshen_4packaged_2 = MakeSymbol("shen.packaged?")
var symshen_4pvar_2 = MakeSymbol("shen.pvar?")
var symshen_4p_1hyps = MakeSymbol("shen.p-hyps")
var symshen_4x_4launcher_4quiet_1load = MakeSymbol("shen.x.launcher.quiet-load")
var symshen_4op1 = MakeSymbol("shen.op1")
var symshen_4fn_1call_2 = MakeSymbol("shen.fn-call?")
var symerror_1to_1string = MakeSymbol("error-to-string")
var symrun = MakeSymbol("run")
var symshen_4x_4launcher_4help_1text = MakeSymbol("shen.x.launcher.help-text")
var symshen_4fbound_2 = MakeSymbol("shen.fbound?")
var symshen_4lowercase_1symbol_2 = MakeSymbol("shen.lowercase-symbol?")
var symshen_4_5yacc_6 = MakeSymbol("shen.<yacc>")
var symshen_4processed = MakeSymbol("shen.processed")
var symcd = MakeSymbol("cd")
var symshen_4_dinfs_d = MakeSymbol("shen.*infs*")
var symis = MakeSymbol("is")
var symout = MakeSymbol("out")
var symshen_4use_1type_1info = MakeSymbol("shen.use-type-info")
var sym_8p = MakeSymbol("@p")
var symit = MakeSymbol("it")
var symshen_4_5head_6 = MakeSymbol("shen.<head>")
var symshen_4_1m = MakeSymbol("shen.-m")
var symshen_4_5lrb_6 = MakeSymbol("shen.<lrb>")
var sympackage = MakeSymbol("package")
var symsave = MakeSymbol("save")
var symshen_4_5clauses_6 = MakeSymbol("shen.<clauses>")
var symshen_4show_1p = MakeSymbol("shen.show-p")
var symshen_4tuple_1up = MakeSymbol("shen.tuple-up")
var symshen_4_5type_6 = MakeSymbol("shen.<type>")
var symshen_4maxinfexceeded_2 = MakeSymbol("shen.maxinfexceeded?")
var symshen_4x_4launcher_4script_1command = MakeSymbol("shen.x.launcher.script-command")
var symfail_1if = MakeSymbol("fail-if")
var symin = MakeSymbol("in")
var symshen_4x_4launcher_4eval_1string = MakeSymbol("shen.x.launcher.eval-string")
var symbound_2 = MakeSymbol("bound?")
var symshen_4read_1loop = MakeSymbol("shen.read-loop")
var symdatatypes = MakeSymbol("datatypes")
var symshen_4hascut_2 = MakeSymbol("shen.hascut?")
var symhush_2 = MakeSymbol("hush?")
var symadjoin = MakeSymbol("adjoin")
var symshen_4constructor_1error = MakeSymbol("shen.constructor-error")
var symshen_4i_1failed_b = MakeSymbol("shen.i-failed!")
var symshen_4parse_1failure = MakeSymbol("shen.parse-failure")
var symshen_4_5float_6 = MakeSymbol("shen.<float>")
var symshen_4_dlambdatable_d = MakeSymbol("shen.*lambdatable*")
var symshen_4lambda_1function = MakeSymbol("shen.lambda-function")
var symshen_4_5bterms_6 = MakeSymbol("shen.<bterms>")
var symshen_4alphanums_2 = MakeSymbol("shen.alphanums?")
var symread_1file_1as_1bytelist = MakeSymbol("read-file-as-bytelist")
var symis_b = MakeSymbol("is!")
var symclose = MakeSymbol("close")
var symshen_4fits_2 = MakeSymbol("shen.fits?")
var symRemainder = MakeSymbol("Remainder")
var symshen_4beta = MakeSymbol("shen.beta")
var symlineread = MakeSymbol("lineread")
var symshen_4_5colon_6 = MakeSymbol("shen.<colon>")
var symshen_4loop = MakeSymbol("shen.loop")
var symshen_4package_1user_1input = MakeSymbol("shen.package-user-input")
var symshen_4passive_1variables = MakeSymbol("shen.passive-variables")
var symRecord = MakeSymbol("Record")
var symexternal = MakeSymbol("external")
var sym_dstinput_d = MakeSymbol("*stinput*")
var symshen_4pivot_1on = MakeSymbol("shen.pivot-on")
var symshen_4variable_1case = MakeSymbol("shen.variable-case")
var symshen_4consume = MakeSymbol("shen.consume")
var symunknown_1arguments = MakeSymbol("unknown-arguments")
var symshen_4_5pattern_6 = MakeSymbol("shen.<pattern>")
var symshen_4expt = MakeSymbol("shen.expt")
var sym_c = MakeSymbol("/")
var symboolean = MakeSymbol("boolean")
var symoutput = MakeSymbol("output")
var symshen_4_5hterm_6 = MakeSymbol("shen.<hterm>")
var symshen_4_5bterm_6 = MakeSymbol("shen.<bterm>")
var symshen_4lock = MakeSymbol("shen.lock")
var symshen_4reverse_1help = MakeSymbol("shen.reverse-help")
var symoptimise = MakeSymbol("optimise")
var symshen_4cond_1form = MakeSymbol("shen.cond-form")
var symshen_4_5comma_6 = MakeSymbol("shen.<comma>")
var symshen_4_5c_1rules_6 = MakeSymbol("shen.<c-rules>")
var symshen_4restore_1P = MakeSymbol("shen.restore-P")
var symshen_4_dsize_1prolog_1vector_d = MakeSymbol("shen.*size-prolog-vector*")
var sym_1 = MakeSymbol("-")
var symupdate_1lambda_1table = MakeSymbol("update-lambda-table")
var symsqts = MakeSymbol("sqts")
var symshen_4spaces = MakeSymbol("shen.spaces")
var symshen_4autocomplete = MakeSymbol("shen.autocomplete")
var symshen_4shen_1_6kl = MakeSymbol("shen.shen->kl")
var symtracked = MakeSymbol("tracked")
var symFinish = MakeSymbol("Finish")
var symshen_4incinfs = MakeSymbol("shen.incinfs")
var symshen_4invoke = MakeSymbol("shen.invoke")
var symshen_4unprotect = MakeSymbol("shen.unprotect")
var symshen_4reader_1error = MakeSymbol("shen.reader-error")
var symshen_4find_1arities = MakeSymbol("shen.find-arities")
var symwrite_1byte = MakeSymbol("write-byte")
var syminclude = MakeSymbol("include")
var symshen_4process_1cases = MakeSymbol("shen.process-cases")
var symTime = MakeSymbol("Time")
var symshen_4c_1rule_1_6shen = MakeSymbol("shen.c-rule->shen")
var symoccurs_2 = MakeSymbol("occurs?")
var symshen_4digit_2 = MakeSymbol("shen.digit?")
var symshen_4process_1sexprs = MakeSymbol("shen.process-sexprs")
var symshen_4_dmaxinferences_d = MakeSymbol("shen.*maxinferences*")
var symshen_4free_1variable_1error_1message = MakeSymbol("shen.free-variable-error-message")
var sym_5_1address = MakeSymbol("<-address")
var symdefcc = MakeSymbol("defcc")
var symshen_4_5double_6 = MakeSymbol("shen.<double>")
var symshen_4multiples = MakeSymbol("shen.multiples")
var symshen_4nothing_1doing_2 = MakeSymbol("shen.nothing-doing?")
var symshen_4findall_1h = MakeSymbol("shen.findall-h")
var symshen_4prolog_1abstraction = MakeSymbol("shen.prolog-abstraction")
var symsuccess = MakeSymbol("success")
var symshen_4sysfunc_2 = MakeSymbol("shen.sysfunc?")
var symshen_4skip = MakeSymbol("shen.skip")
var symshen_4_5e_1number_6 = MakeSymbol("shen.<e-number>")
var symshen_4application_2 = MakeSymbol("shen.application?")
var symshen_4predicate = MakeSymbol("shen.predicate")
var symshen_4curry = MakeSymbol("shen.curry")
var symshen_4change_1pointer_1value = MakeSymbol("shen.change-pointer-value")
var symaddress_1_6 = MakeSymbol("address->")
var symshen_4combine_1c_1code = MakeSymbol("shen.combine-c-code")
var symshen_4app = MakeSymbol("shen.app")
var symsubst = MakeSymbol("subst")
var symshen_4_dit_d = MakeSymbol("shen.*it*")
var symshen_4_8c = MakeSymbol("shen.@c")
var symshen_4_5times_6 = MakeSymbol("shen.<times>")
var symshen_4_dsystem_d = MakeSymbol("shen.*system*")
var symunprofile = MakeSymbol("unprofile")
var symshen_4decrement_1ticket = MakeSymbol("shen.decrement-ticket")
var symshen_4variancy = MakeSymbol("shen.variancy")
var sym_d = MakeSymbol("*")
var symshen_4read_1evaluate_1print = MakeSymbol("shen.read-evaluate-print")
var symshen_4variants_2 = MakeSymbol("shen.variants?")
var symshen_4_5stop_6 = MakeSymbol("shen.<stop>")
var symshen_4_5sig_drules_6 = MakeSymbol("shen.<sig*rules>")
var symshen_4_5simple_1pattern_6 = MakeSymbol("shen.<simple-pattern>")
var symshen_4record_1external = MakeSymbol("shen.record-external")
var symprolog_1memory = MakeSymbol("prolog-memory")
var symshen_4remove_1h = MakeSymbol("shen.remove-h")
var symshen_4fillvector = MakeSymbol("shen.fillvector")
var symhush = MakeSymbol("hush")
var symshen_4_5body_6 = MakeSymbol("shen.<body>")
var symshen_4_5dbl_6 = MakeSymbol("shen.<dbl>")
var symconcat = MakeSymbol("concat")
var symunion = MakeSymbol("union")
var symbind = MakeSymbol("bind")
var symshen_4passive_1bind = MakeSymbol("shen.passive-bind")
var symshen_4factor_1selectors = MakeSymbol("shen.factor-selectors")
var symthaw = MakeSymbol("thaw")
var symshen_4tuple = MakeSymbol("shen.tuple")
var symoccurrences = MakeSymbol("occurrences")
var symshen_4_dshen_1type_1theory_1enabled_2_d = MakeSymbol("shen.*shen-type-theory-enabled?*")
var symshen_4factorise_1code = MakeSymbol("shen.factorise-code")
var symshen_4deref_1terms = MakeSymbol("shen.deref-terms")
var symshen_4primitive = MakeSymbol("shen.primitive")
var symwrite_1to_1file = MakeSymbol("write-to-file")
var symshen_4insert_1l = MakeSymbol("shen.insert-l")
var symshen_4f_1error = MakeSymbol("shen.f-error")
var symW = MakeSymbol("W")
var symfunction = MakeSymbol("function")
var symsystemf = MakeSymbol("systemf")
var symshen_4call_1prolog = MakeSymbol("shen.call-prolog")
var symshen_4freshen_1type = MakeSymbol("shen.freshen-type")
var sym_5_1 = MakeSymbol("<-")
var sym_e = MakeSymbol("&")
var symshen_4_5non_1terminal_2_6 = MakeSymbol("shen.<non-terminal?>")
var symshen_4store_1arity = MakeSymbol("shen.store-arity")
var symshen_4terms = MakeSymbol("shen.terms")
var symshen_4unwind = MakeSymbol("shen.unwind")
var symshen_4continue = MakeSymbol("shen.continue")
var symA = MakeSymbol("A")
var symshen_4make_1prolog_1variable = MakeSymbol("shen.make-prolog-variable")
var symnth = MakeSymbol("nth")
var symshen_4this_1symbol_1is_1unbound = MakeSymbol("shen.this-symbol-is-unbound")
var symtype = MakeSymbol("type")
var symshen_4unpackage = MakeSymbol("shen.unpackage")
var symshen_4process_1input_7 = MakeSymbol("shen.process-input+")
var symshen_4deref = MakeSymbol("shen.deref")
var symTm = MakeSymbol("Tm")
var sym_i = MakeSymbol("{")
var symshen_4_5c_1rule_6 = MakeSymbol("shen.<c-rule>")
var symset = MakeSymbol("set")
var symshen_4_8ch = MakeSymbol("shen.@ch")
var symnl = MakeSymbol("nl")
var symshen_4choicepoint = MakeSymbol("shen.choicepoint")
var symshen_4remove_1indirection = MakeSymbol("shen.remove-indirection")
var symshen_4_5strcontents_6 = MakeSymbol("shen.<strcontents>")
var symshen_4freshterms = MakeSymbol("shen.freshterms")
var sym_8v = MakeSymbol("@v")
var symshen_4_5whitespace_6 = MakeSymbol("shen.<whitespace>")
var symasserta = MakeSymbol("asserta")
var symshen_4_7string_2 = MakeSymbol("shen.+string?")
var symfst = MakeSymbol("fst")
var symshen_4callrec = MakeSymbol("shen.callrec")
var symshen_4mod = MakeSymbol("shen.mod")
var sym__ = MakeSymbol("_")
var symFreeze = MakeSymbol("Freeze")
var symn_1_6string = MakeSymbol("n->string")
var symshen_4goto_1h = MakeSymbol("shen.goto-h")
var symtlv = MakeSymbol("tlv")
var symlength = MakeSymbol("length")
var symshen_4_5constructor_6 = MakeSymbol("shen.<constructor>")
var symread_1from_1string_1unprocessed = MakeSymbol("read-from-string-unprocessed")
var symor = MakeSymbol("or")
var symshen_4rule_1_6clause = MakeSymbol("shen.rule->clause")
var symshen_4abs = MakeSymbol("shen.abs")
var sym_dabsolute_d = MakeSymbol("*absolute*")
var symshen_4insert_1h = MakeSymbol("shen.insert-h")
var symshen_4try_1parse = MakeSymbol("shen.try-parse")
var symshen_4yacc_1_6shen = MakeSymbol("shen.yacc->shen")
var symTl = MakeSymbol("Tl")
var symshen_4record_1internal = MakeSymbol("shen.record-internal")
var symshen_4shen_1call_2 = MakeSymbol("shen.shen-call?")
var symshen_4_5defprolog_6 = MakeSymbol("shen.<defprolog>")
var symshen_4lzy_a_b = MakeSymbol("shen.lzy=!")
var sym_dversion_d = MakeSymbol("*version*")
var symlambda = MakeSymbol("lambda")
var symshen_4_5lowE_6 = MakeSymbol("shen.<lowE>")
var symshen_4non_1application_2 = MakeSymbol("shen.non-application?")
var symshen_4lchh = MakeSymbol("shen.lchh")
var symshen_4output_1track = MakeSymbol("shen.output-track")
var symarity = MakeSymbol("arity")
var symshen_4maxseq = MakeSymbol("shen.maxseq")
var symshen_4hds_a_2 = MakeSymbol("shen.hds=?")
var symshen_4find_1free_1vars = MakeSymbol("shen.find-free-vars")
var symshen_4unpack_1foreign = MakeSymbol("shen.unpack-foreign")
var symshen_4tame = MakeSymbol("shen.tame")
var symshen_4x_4launcher_4default_1handle_1result = MakeSymbol("shen.x.launcher.default-handle-result")
var symcn = MakeSymbol("cn")
var sym_j = MakeSymbol("}")
var symL = MakeSymbol("L")
var symshen_4consume_1clause = MakeSymbol("shen.consume-clause")
var symshen_4objectcode = MakeSymbol("shen.objectcode")
var symshen_4_dgensym_d = MakeSymbol("shen.*gensym*")
var symspecialise = MakeSymbol("specialise")
var symshen_4funexstring = MakeSymbol("shen.funexstring")
var symshen_4_5s_1exprs_6 = MakeSymbol("shen.<s-exprs>")
var symshen_4_5lsb_6 = MakeSymbol("shen.<lsb>")
var symu_b = MakeSymbol("u!")
var symshen_4make_1uppercase = MakeSymbol("shen.make-uppercase")
var symshen_4atom_1case_1plus = MakeSymbol("shen.atom-case-plus")
var symshen_4_5datatype_1rule_6 = MakeSymbol("shen.<datatype-rule>")
var symshen_4_5sides_6 = MakeSymbol("shen.<sides>")
var symHd = MakeSymbol("Hd")
var symappend = MakeSymbol("append")
var symmaxinferences = MakeSymbol("maxinferences")
var symshen_4_dsigf_d = MakeSymbol("shen.*sigf*")
var symshen_4rectify_1test = MakeSymbol("shen.rectify-test")
var symshen_4_5literal_6 = MakeSymbol("shen.<literal>")
var symshen_4free_1var_1chk = MakeSymbol("shen.free-var-chk")
var symshen_4print_1prolog_1vector = MakeSymbol("shen.print-prolog-vector")
var symshen_4prhush = MakeSymbol("shen.prhush")
var symshen_4shendef_1_6kldef_1h = MakeSymbol("shen.shendef->kldef-h")
var symshen_4_5s_1exprs2_6 = MakeSymbol("shen.<s-exprs2>")
var sym_3 = MakeSymbol("$")
var symshen_4unlock = MakeSymbol("shen.unlock")
var symshen_4iter_1vector = MakeSymbol("shen.iter-vector")
var symshen_4assert_d = MakeSymbol("shen.assert*")
var symvector = MakeSymbol("vector")
var symshen_4s = MakeSymbol("shen.s")
var symshen_4t = MakeSymbol("shen.t")
var symshen_4initialise_1lambda_1tables = MakeSymbol("shen.initialise-lambda-tables")
var symshen_4partial = MakeSymbol("shen.partial")
var symshen_4prolog_1track = MakeSymbol("shen.prolog-track")
var symshen_4x_4launcher_4eval_1flag_1map = MakeSymbol("shen.x.launcher.eval-flag-map")
var symfn = MakeSymbol("fn")
var symshen_4pvar = MakeSymbol("shen.pvar")
var sym_dimplementation_d = MakeSymbol("*implementation*")
var symshen_4dbl_2 = MakeSymbol("shen.dbl?")
var symshen_4_5colon_1equal_6 = MakeSymbol("shen.<colon-equal>")
var symshen_4uppercase_2 = MakeSymbol("shen.uppercase?")
var symZ = MakeSymbol("Z")
var symshen_4_5s_1exprs1_6 = MakeSymbol("shen.<s-exprs1>")
var symshen_4misc_2 = MakeSymbol("shen.misc?")
var symdeclare = MakeSymbol("declare")
var sym_dporters_d = MakeSymbol("*porters*")
var symshen_4_5return_6 = MakeSymbol("shen.<return>")
var symshen_4internal_1to_1shen_2 = MakeSymbol("shen.internal-to-shen?")
var symshen_4insert_1prolog_1variables = MakeSymbol("shen.insert-prolog-variables")
var symshen_4t_d_1rule = MakeSymbol("shen.t*-rule")
var symctxt = MakeSymbol("ctxt")
var symshen_4key_1in_1sequent_1calculus_2 = MakeSymbol("shen.key-in-sequent-calculus?")
var symshen_4hashkey = MakeSymbol("shen.hashkey")
var sympackage_2 = MakeSymbol("package?")
var symshen_4_5shortnatter_6 = MakeSymbol("shen.<shortnatter>")
var symshen_4_dalldatatypes_d = MakeSymbol("shen.*alldatatypes*")
var symshen_4compile_1synonyms = MakeSymbol("shen.compile-synonyms")
var symshen_4dynamic_1default = MakeSymbol("shen.dynamic-default")
var symshen_4line = MakeSymbol("shen.line")
var symshen_4cons_1case_1minus = MakeSymbol("shen.cons-case-minus")
var symshen_4bindv = MakeSymbol("shen.bindv")
var symshen_4undefined_1f_2 = MakeSymbol("shen.undefined-f?")
var symshen_4_dcall_d = MakeSymbol("shen.*call*")
var symshen_4_doptimise_d = MakeSymbol("shen.*optimise*")
var sym_5e_6 = MakeSymbol("<e>")
var symshen_4_5log10_6 = MakeSymbol("shen.<log10>")
var symspy_2 = MakeSymbol("spy?")
var symshen_4nvars = MakeSymbol("shen.nvars")
var symshen_4build_1lambda_1table = MakeSymbol("shen.build-lambda-table")
var symshen_4record_1kl = MakeSymbol("shen.record-kl")
var symshen_4process_1after_1type = MakeSymbol("shen.process-after-type")
var symshen_4load_1help = MakeSymbol("shen.load-help")
var symcons_2 = MakeSymbol("cons?")
var symshen_4_dprolog_1memory_d = MakeSymbol("shen.*prolog-memory*")
var sym_dproperty_1vector_d = MakeSymbol("*property-vector*")
var symshen_4_doccurs_d = MakeSymbol("shen.*occurs*")
var symshen_4_5iscolon_6 = MakeSymbol("shen.<iscolon>")
var symshen_4find_1types = MakeSymbol("shen.find-types")
var symshen_4explode_1h = MakeSymbol("shen.explode-h")
var symshen_4fn_1call = MakeSymbol("shen.fn-call")
var symintern = MakeSymbol("intern")
var symshen_4create_1skeleton = MakeSymbol("shen.create-skeleton")
var symshen_4mkstr_1r = MakeSymbol("shen.mkstr-r")
var symshen_4show_1datatypes = MakeSymbol("shen.show-datatypes")
var symgensym = MakeSymbol("gensym")
var symshen_4proc_1nl = MakeSymbol("shen.proc-nl")
var symshen_4integer_1test_2 = MakeSymbol("shen.integer-test?")
var symoptimise_2 = MakeSymbol("optimise?")
var symshen_4string_1_6byte = MakeSymbol("shen.string->byte")
var symlist = MakeSymbol("list")
var symshen_4dynamic = MakeSymbol("shen.dynamic")
var symshen_4compile_1body = MakeSymbol("shen.compile-body")
var symfork = MakeSymbol("fork")
var symshen_4remove_1bystanders = MakeSymbol("shen.remove-bystanders")
var symS = MakeSymbol("S")
var symshen_4choicepoint_b = MakeSymbol("shen.choicepoint!")
var symshen_4specialise_1member = MakeSymbol("shen.specialise-member")
var symshen_4terminalcode = MakeSymbol("shen.terminalcode")
var symsum = MakeSymbol("sum")
var sympr = MakeSymbol("pr")
var symshen_4function_1calls = MakeSymbol("shen.function-calls")
var symshen = MakeSymbol("shen")
var symshen_4peek_1history = MakeSymbol("shen.peek-history")
var symshen_4cons_1form_1respect_1modes = MakeSymbol("shen.cons-form-respect-modes")
var symshen_4demode = MakeSymbol("shen.demode")
var symshen_4system_1S_1h = MakeSymbol("shen.system-S-h")
var symshen_4alpha_1convert = MakeSymbol("shen.alpha-convert")
var symshen_4print_1vector_2 = MakeSymbol("shen.print-vector?")
var symshen_4special_2 = MakeSymbol("shen.special?")
var symopen = MakeSymbol("open")
var symshen_4bytes_1_6string = MakeSymbol("shen.bytes->string")
var symshen_4macros = MakeSymbol("shen.macros")
var symNewAssumptions = MakeSymbol("NewAssumptions")
var symshen_4write_1kl = MakeSymbol("shen.write-kl")
var symAssumptions = MakeSymbol("Assumptions")
var symshen_4string_1_6bytes = MakeSymbol("shen.string->bytes")
var symshen_4compute_1integer = MakeSymbol("shen.compute-integer")
var symshen_4rectify_1type = MakeSymbol("shen.rectify-type")
var symshen_4t_d_1integrity = MakeSymbol("shen.t*-integrity")
var symnot = MakeSymbol("not")
var symshen_4_duserdefs_d = MakeSymbol("shen.*userdefs*")
var syminput = MakeSymbol("input")
var symload = MakeSymbol("load")
var symshen_4deref_1forked_1literals = MakeSymbol("shen.deref-forked-literals")
var symshen_4lambda_1entry = MakeSymbol("shen.lambda-entry")
var symforeign = MakeSymbol("foreign")
var symshen_4_5prem_6 = MakeSymbol("shen.<prem>")
var symshen_4f = MakeSymbol("shen.f")
var symprofile = MakeSymbol("profile")
var symshen_4process_1time = MakeSymbol("shen.process-time")
var symshen_4colon_1equal_2 = MakeSymbol("shen.colon-equal?")
var symshen_4_5iscomma_6 = MakeSymbol("shen.<iscomma>")
var symshen_4my_1read_1byte = MakeSymbol("shen.my-read-byte")
var symshen_4recursive_1string_1match = MakeSymbol("shen.recursive-string-match")
var symshen_4pac_1h = MakeSymbol("shen.pac-h")
var symshen_4compile_1head = MakeSymbol("shen.compile-head")
var symy_1or_1n_2 = MakeSymbol("y-or-n?")
var symhdstr = MakeSymbol("hdstr")
var symshen_4freshterm = MakeSymbol("shen.freshterm")
var symshen_4mkstr_1l = MakeSymbol("shen.mkstr-l")
var symfix = MakeSymbol("fix")
var symabsolute = MakeSymbol("absolute")
var symshen_4rdecons = MakeSymbol("shen.rdecons")
var symlaunch_1repl = MakeSymbol("launch-repl")
var symshen_4_5plus_6 = MakeSymbol("shen.<plus>")
var symdefmacro = MakeSymbol("defmacro")
var symshen_4call_1dynamic = MakeSymbol("shen.call-dynamic")
var symshen_4system_1S = MakeSymbol("shen.system-S")
var sym_dmaximum_1print_1sequence_1size_d = MakeSymbol("*maximum-print-sequence-size*")
var sym_6_a = MakeSymbol(">=")
var sym_dport_d = MakeSymbol("*port*")
var symshen_4linearise_1h = MakeSymbol("shen.linearise-h")
var symshen_4_dnames_d = MakeSymbol("shen.*names*")
var symshen_4prolog_1parameters = MakeSymbol("shen.prolog-parameters")
var symwhere = MakeSymbol("where")
var symif = MakeSymbol("if")
var symshen_4typename_1h = MakeSymbol("shen.typename-h")
var symstr = MakeSymbol("str")
var symshen_4prolog_1arity_1check = MakeSymbol("shen.prolog-arity-check")
var symshen_4_5hterm1_6 = MakeSymbol("shen.<hterm1>")
var symshen_4_5notdbq_6 = MakeSymbol("shen.<notdbq>")
var symshen_4insert = MakeSymbol("shen.insert")
var symshen_4internal_1symbols = MakeSymbol("shen.internal-symbols")
var symshen_4analyse_1variable_2 = MakeSymbol("shen.analyse-variable?")
var symshen_4credits = MakeSymbol("shen.credits")
var symshen_4tracked_2 = MakeSymbol("shen.tracked?")
var symshen_4demodulate = MakeSymbol("shen.demodulate")
var symshen_4sigf = MakeSymbol("shen.sigf")
var symshen_4freshen = MakeSymbol("shen.freshen")
var symlanguage = MakeSymbol("language")
var symshen_4fail_b = MakeSymbol("shen.fail!")
var symshen_4factor_1selectors_1h = MakeSymbol("shen.factor-selectors-h")
var sym_dstoutput_d = MakeSymbol("*stoutput*")
var symuntrack = MakeSymbol("untrack")
var symshen_4lazyderef = MakeSymbol("shen.lazyderef")
var symshen_4sigxrules = MakeSymbol("shen.sigxrules")
var symshen_4insert_1tracking_1code = MakeSymbol("shen.insert-tracking-code")
var symtc = MakeSymbol("tc")
var symfactorise = MakeSymbol("factorise")
var symshen_4write_1kl_1h = MakeSymbol("shen.write-kl-h")
var symshen_4newpv = MakeSymbol("shen.newpv")
var symshen_4factor_1cn = MakeSymbol("shen.factor-cn")
var symshen_4_5fraction_6 = MakeSymbol("shen.<fraction>")
var symmapcan = MakeSymbol("mapcan")
var sym_e_e = MakeSymbol("&&")
var symboolean_2 = MakeSymbol("boolean?")
var symshen_4simple_1curry = MakeSymbol("shen.simple-curry")
var symshen_4assumetypes = MakeSymbol("shen.assumetypes")
var symshen_4_dtc_d = MakeSymbol("shen.*tc*")
var symshen_4record_1it = MakeSymbol("shen.record-it")
var symshen_4_5minus_6 = MakeSymbol("shen.<minus>")
var symstring_2 = MakeSymbol("string?")
var symshen_4prolog_1vector = MakeSymbol("shen.prolog-vector")
var symshen_4work_1through = MakeSymbol("shen.work-through")
var sym_drelease_d = MakeSymbol("*release*")
var symshen_4op_1test = MakeSymbol("shen.op-test")
var symshen_4input_1h_7 = MakeSymbol("shen.input-h+")
var symshen_4_dsynonyms_d = MakeSymbol("shen.*synonyms*")
var symcall = MakeSymbol("call")
var symmode = MakeSymbol("mode")
var symshen_4_5comment_6 = MakeSymbol("shen.<comment>")
var symsynonyms = MakeSymbol("synonyms")
var symfile = MakeSymbol("file")
var symshen_4recursively_1print = MakeSymbol("shen.recursively-print")
var symshow_1help = MakeSymbol("show-help")
var symtlstr = MakeSymbol("tlstr")
var symshen_4_5pattern2_6 = MakeSymbol("shen.<pattern2>")
var symlet = MakeSymbol("let")
var symshen_4recursively_1factor_1selectors = MakeSymbol("shen.recursively-factor-selectors")
var symshen_4_5equal_6 = MakeSymbol("shen.<equal>")
var symcases = MakeSymbol("cases")
var sym_8s = MakeSymbol("@s")
var symshen_4_dfactorise_2_d = MakeSymbol("shen.*factorise?*")
var symshen_4_dextraspecial_d = MakeSymbol("shen.*extraspecial*")
var symshen_4_5_1out = MakeSymbol("shen.<-out")
var symshen_4_5returns_6 = MakeSymbol("shen.<returns>")
var symshen_4check_1byte = MakeSymbol("shen.check-byte")
var symexplode = MakeSymbol("explode")
var symshen_4choicepoint_2 = MakeSymbol("shen.choicepoint?")
var symshen_4str_1_6bytes = MakeSymbol("shen.str->bytes")
var symshen_4add_1sexpr = MakeSymbol("shen.add-sexpr")
var symshen_4pui_1h = MakeSymbol("shen.pui-h")
var symshen_4atom_1case_1minus = MakeSymbol("shen.atom-case-minus")
var symP = MakeSymbol("P")
var symshen_4x_4launcher_4main = MakeSymbol("shen.x.launcher.main")
var syminput_7 = MakeSymbol("input+")
var symSelect = MakeSymbol("Select")
var symshen_4_5strc_6 = MakeSymbol("shen.<strc>")
var symshen_4sng_2 = MakeSymbol("shen.sng?")
var symshen_4t_d = MakeSymbol("shen.t*")
var symshen_4_5rule_d_6 = MakeSymbol("shen.<rule*>")
var symshen_4factor_1recognisors = MakeSymbol("shen.factor-recognisors")
var symshen_4reader_1error_1message = MakeSymbol("shen.reader-error-message")
var symshen_4_5rsb_6 = MakeSymbol("shen.<rsb>")
var symshen_4macroexpand_1h = MakeSymbol("shen.macroexpand-h")
var symshen_4process_1datatype = MakeSymbol("shen.process-datatype")
var symshen_4_5datatype_6 = MakeSymbol("shen.<datatype>")
var symshen_4unlocked_2 = MakeSymbol("shen.unlocked?")
var symshen_4variablecode = MakeSymbol("shen.variablecode")
var symshen_4source = MakeSymbol("shen.source")
var symread_1file_1as_1string = MakeSymbol("read-file-as-string")
var symshen_4_1null_1 = MakeSymbol("shen.-null-")
var symvector_1_6 = MakeSymbol("vector->")
var symshen_4unassoc = MakeSymbol("shen.unassoc")
var symshen_4g = MakeSymbol("shen.g")
var symget_1time = MakeSymbol("get-time")
var symexception = MakeSymbol("exception")
var symshen_4typetable = MakeSymbol("shen.typetable")
var symshen_4rule_1_6body = MakeSymbol("shen.rule->body")
var symshen_4_5atom_6 = MakeSymbol("shen.<atom>")
var symlazy = MakeSymbol("lazy")
var symshen_4evaluate_1lineread = MakeSymbol("shen.evaluate-lineread")
var symshen_4_5syntax_1item_6 = MakeSymbol("shen.<syntax-item>")
var symshen_4process_1yacc_1semantics = MakeSymbol("shen.process-yacc-semantics")
var symshen_4str_1_6str = MakeSymbol("shen.str->str")
var symshen_4bad_1pivot_2 = MakeSymbol("shen.bad-pivot?")
var symshen_4prolog_1keyword_2 = MakeSymbol("shen.prolog-keyword?")
var symshen_4extract_1free_1vars = MakeSymbol("shen.extract-free-vars")
var symshen_4nextticket = MakeSymbol("shen.nextticket")
var symshen_4raise_1syntax_1error = MakeSymbol("shen.raise-syntax-error")
var symshen_4linearise = MakeSymbol("shen.linearise")
var symshen_4lowercase_2 = MakeSymbol("shen.lowercase?")
var symshen_4initialise__environment = MakeSymbol("shen.initialise_environment")
var symshen_4intern_1type = MakeSymbol("shen.intern-type")
var symprotect = MakeSymbol("protect")
var symshen_4write_1string = MakeSymbol("shen.write-string")
var symshen_4profile_1help = MakeSymbol("shen.profile-help")
var symshen_4list_1_6str = MakeSymbol("shen.list->str")
var symshen_4iter_1list = MakeSymbol("shen.iter-list")
var symdefun = MakeSymbol("defun")
var sym_dlanguage_d = MakeSymbol("*language*")
var symshen_4scan_1body = MakeSymbol("shen.scan-body")
var symbar_b = MakeSymbol("bar!")
var symshen_4goto = MakeSymbol("shen.goto")
var symshen_4r = MakeSymbol("shen.r")
var symshen_4shen = MakeSymbol("shen.shen")
var symshen_4coll_1formulae = MakeSymbol("shen.coll-formulae")
var symshen_4get_1profile = MakeSymbol("shen.get-profile")
var symshen_4_5longnatter_6 = MakeSymbol("shen.<longnatter>")
var symshen_4compute_1integer_1h = MakeSymbol("shen.compute-integer-h")
var symatom_2 = MakeSymbol("atom?")
var symshen_4correct = MakeSymbol("shen.correct")
var symuserdefs = MakeSymbol("userdefs")
var symdestroy = MakeSymbol("destroy")
var symeval_1kl = MakeSymbol("eval-kl")
var symfail = MakeSymbol("fail")
var symshen_4internal_2 = MakeSymbol("shen.internal?")
var symstep = MakeSymbol("step")
var symshen_4process_1read_1byte = MakeSymbol("shen.process-read-byte")
var symshen_4_5rules_d_6 = MakeSymbol("shen.<rules*>")
var symshen_4shen_1_6kl_1h = MakeSymbol("shen.shen->kl-h")
var symshen_4walk = MakeSymbol("shen.walk")
var symshen_4macro_1_8ch = MakeSymbol("shen.macro-@ch")
var symshen_4ok = MakeSymbol("shen.ok")
var symoccurs_1check = MakeSymbol("occurs-check")
var symshen_4_5ass_6 = MakeSymbol("shen.<ass>")
var symassoc = MakeSymbol("assoc")
var symput = MakeSymbol("put")
var symshen_4record_1macro = MakeSymbol("shen.record-macro")
var symHypotheses = MakeSymbol("Hypotheses")
var sym_dhome_1directory_d = MakeSymbol("*home-directory*")
var symshen_4_5semicolon_6 = MakeSymbol("shen.<semicolon>")
var symshen_4_5str_6 = MakeSymbol("shen.<str>")
var symshen_4write_1chars = MakeSymbol("shen.write-chars")
var symGoTo = MakeSymbol("GoTo")
var symread_1from_1string = MakeSymbol("read-from-string")
var symshen_4conscode = MakeSymbol("shen.conscode")
var symshen_4tlv_1help = MakeSymbol("shen.tlv-help")
var symY = MakeSymbol("Y")
var symdatatype = MakeSymbol("datatype")
var symshen_4remember_1datatype = MakeSymbol("shen.remember-datatype")
var symstinput = MakeSymbol("stinput")
var symreverse = MakeSymbol("reverse")
var symshen_4macro_1_8c = MakeSymbol("shen.macro-@c")
var symshen_4update_1lambdatable = MakeSymbol("shen.update-lambdatable")
var sym_a_a_6 = MakeSymbol("==>")
var symshen_4_5expr_6 = MakeSymbol("shen.<expr>")
var symshen_4l_1rules = MakeSymbol("shen.l-rules")
var symunput = MakeSymbol("unput")
var symshen_4remove_1datatypes = MakeSymbol("shen.remove-datatypes")
var symshen_4update_1history = MakeSymbol("shen.update-history")
var symshen_4freshen_1rule = MakeSymbol("shen.freshen-rule")
var symAction = MakeSymbol("Action")
var symvector_2 = MakeSymbol("vector?")
var symnumber_2 = MakeSymbol("number?")
var symshen_4process_1_8s = MakeSymbol("shen.process-@s")
var symshen_4semicolon_2 = MakeSymbol("shen.semicolon?")
var symshen_4x_4launcher_4launch_1shen = MakeSymbol("shen.x.launcher.launch-shen")
var symshen_4lr_1rule = MakeSymbol("shen.lr-rule")
var symeval = MakeSymbol("eval")
var symreturn = MakeSymbol("return")
var symvalue = MakeSymbol("value")
var symK = MakeSymbol("K")
var symshen_4retract_1clause = MakeSymbol("shen.retract-clause")
var symshen_4wildcard_2 = MakeSymbol("shen.wildcard?")
var symKey = MakeSymbol("Key")
var symshen_4repl = MakeSymbol("shen.repl")
var symshen_4_5rule_6 = MakeSymbol("shen.<rule>")
var symshen_4horn_1clause_1procedure = MakeSymbol("shen.horn-clause-procedure")
var symshen_4linearise_1clause = MakeSymbol("shen.linearise-clause")
var symshen_4restore_1local = MakeSymbol("shen.restore-local")
var symerror = MakeSymbol("error")
var symstring = MakeSymbol("string")
var symshen_4locked_2 = MakeSymbol("shen.locked?")
var symremove = MakeSymbol("remove")
var symshen_4compute_1E = MakeSymbol("shen.compute-E")
var symshen_4x_4launcher_4eval_1command = MakeSymbol("shen.x.launcher.eval-command")
var symshen_4vector_1_6str = MakeSymbol("shen.vector->str")
var symV = MakeSymbol("V")
var symshen_4package_1symbols = MakeSymbol("shen.package-symbols")
var symshen_4process_1assoc = MakeSymbol("shen.process-assoc")
var symshen_4premises_1_6goals = MakeSymbol("shen.premises->goals")
var symcompile = MakeSymbol("compile")
var symshen_4extract_1vars = MakeSymbol("shen.extract-vars")
var symshen_4eval_1and_1print = MakeSymbol("shen.eval-and-print")
var symC = MakeSymbol("C")
var symshen_4put_1profile = MakeSymbol("shen.put-profile")
var symtc_2 = MakeSymbol("tc?")
var symshen_4printF = MakeSymbol("shen.printF")
var symshen_4_5digits_6 = MakeSymbol("shen.<digits>")
var symshen_4type_1F = MakeSymbol("shen.type-F")
var symshen_4partial_1parse_1failure_2 = MakeSymbol("shen.partial-parse-failure?")
var symshen_4non_1terminal_2 = MakeSymbol("shen.non-terminal?")
var symverified = MakeSymbol("verified")
var symshen_4compile_1prolog = MakeSymbol("shen.compile-prolog")
var symshen_4_dpackage_d = MakeSymbol("shen.*package*")
var symshen_4_5hash_6 = MakeSymbol("shen.<hash>")
var symshen_4initialise_1arity_1table = MakeSymbol("shen.initialise-arity-table")
var symshen_4process_1let = MakeSymbol("shen.process-let")
var symshen_4top = MakeSymbol("shen.top")
var symshen_4signal_1def = MakeSymbol("shen.signal-def")
var symshen_4syntax_1error_1message = MakeSymbol("shen.syntax-error-message")
var symcons = MakeSymbol("cons")
var symshen_4_5datatype_1rules_6 = MakeSymbol("shen.<datatype-rules>")
var symshen_4t_d_1rules = MakeSymbol("shen.t*-rules")
var symshen_4x_4launcher_4done = MakeSymbol("shen.x.launcher.done")
var sym_7 = MakeSymbol("+")
var symshen_4zero_1place_2 = MakeSymbol("shen.zero-place?")
var symempty_2 = MakeSymbol("empty?")
var symvariable_2 = MakeSymbol("variable?")
var symshen_4deref_1calls = MakeSymbol("shen.deref-calls")
var symshen_4x_4launcher_4version_1string = MakeSymbol("shen.x.launcher.version-string")
var symshen_4in_1_6 = MakeSymbol("shen.in->")
var symshen_4synonyms_1h = MakeSymbol("shen.synonyms-h")
var symshen_4fix_1help = MakeSymbol("shen.fix-help")
var symshen_4compute_1fraction = MakeSymbol("shen.compute-fraction")
var symshen_4copyfromvector = MakeSymbol("shen.copyfromvector")
var symshen_4search_1user_1datatypes = MakeSymbol("shen.search-user-datatypes")
var symincluded = MakeSymbol("included")
var symshen_4_5rrb_6 = MakeSymbol("shen.<rrb>")
var symshen_4vector_1dereference = MakeSymbol("shen.vector-dereference")
var symshen_4t_d_1rule_1h = MakeSymbol("shen.t*-rule-h")
var symshen_4_5yaccsig_6 = MakeSymbol("shen.<yaccsig>")
var symshen_4modh = MakeSymbol("shen.modh")
var symshen_4hush = MakeSymbol("shen.hush")
var symshen_4_5signature_6 = MakeSymbol("shen.<signature>")
var symshen_4received = MakeSymbol("shen.received")
var symshen_4specialise_1consume = MakeSymbol("shen.specialise-consume")
var symspy = MakeSymbol("spy")
var symnumber = MakeSymbol("number")
var sym_c_4 = MakeSymbol("/.")
var sym_5_1_1 = MakeSymbol("<--")
var symX = MakeSymbol("X")
var symshen_4cut = MakeSymbol("shen.cut")
var sym_5end_6 = MakeSymbol("<end>")
var symshen_4compute_1fraction_1h = MakeSymbol("shen.compute-fraction-h")
var symshen_4special_1case = MakeSymbol("shen.special-case")
var symAssumption = MakeSymbol("Assumption")
var symtail = MakeSymbol("tail")
var symporters = MakeSymbol("porters")
var symshen_4process_1def = MakeSymbol("shen.process-def")
var symshen_4external_1symbols = MakeSymbol("shen.external-symbols")
var symshen_4_dprofiled_d = MakeSymbol("shen.*profiled*")
var symshen_4by_1hypothesis = MakeSymbol("shen.by-hypothesis")
var symshen_4list_2 = MakeSymbol("shen.list?")
var symshen_4_5name_6 = MakeSymbol("shen.<name>")
var symshen_4curry_1type = MakeSymbol("shen.curry-type")
var symMessage = MakeSymbol("Message")
var symshen_4_8v_1help = MakeSymbol("shen.@v-help")
var symshen_4_5lcurly_6 = MakeSymbol("shen.<lcurly>")
var symundefmacro = MakeSymbol("undefmacro")
var symshen_4cons_1case_1plus = MakeSymbol("shen.cons-case-plus")
var symshen_4_7vector_2 = MakeSymbol("shen.+vector?")
var symshen_4internal_1to_1P_2 = MakeSymbol("shen.internal-to-P?")
var sym_1_1_6 = MakeSymbol("-->")
var symprolog_2 = MakeSymbol("prolog?")
var symshen_4free_1variable_2 = MakeSymbol("shen.free-variable?")
var symshen_4partial_1application_d_2 = MakeSymbol("shen.partial-application*?")
var symshen_4char_1stinput_2 = MakeSymbol("shen.char-stinput?")
var symshen_4unwind_1types = MakeSymbol("shen.unwind-types")
var symshen_4syntax_1item_2 = MakeSymbol("shen.syntax-item?")
var symshen_4monomorphic_2 = MakeSymbol("shen.monomorphic?")
var sym_5_a = MakeSymbol("<=")
var symread_1byte = MakeSymbol("read-byte")
var symbootstrap = MakeSymbol("bootstrap")
var symshen_4lch = MakeSymbol("shen.lch")
var symshen_4_5packagenames_6 = MakeSymbol("shen.<packagenames>")
var symps = MakeSymbol("ps")
var symshen_4_dspecial_d = MakeSymbol("shen.*special*")
var symshen_4_5patterns_6 = MakeSymbol("shen.<patterns>")
var symshen_4wildcardcode = MakeSymbol("shen.wildcardcode")
var symget = MakeSymbol("get")
var symshen_4_5shortnatters_6 = MakeSymbol("shen.<shortnatters>")
var symelement_2 = MakeSymbol("element?")
var symshen_4_5bar_6 = MakeSymbol("shen.<bar>")
var symshen_4_5alphanums_6 = MakeSymbol("shen.<alphanums>")
var symshen_4rules_1_6prolog = MakeSymbol("shen.rules->prolog")
var symshen_4profiled_2 = MakeSymbol("shen.profiled?")
var symfactorise_2 = MakeSymbol("factorise?")
var sym_6 = MakeSymbol(">")
var sym_a = MakeSymbol("=")
var symshen_4op = MakeSymbol("shen.op")
var symshen_4overapplication_2 = MakeSymbol("shen.overapplication?")
var symshen_4prolog_1vector_1size = MakeSymbol("shen.prolog-vector-size")
var symshen_4process_1applications = MakeSymbol("shen.process-applications")
var symtl = MakeSymbol("tl")
var symtrap_1error = MakeSymbol("trap-error")
var sympreclude = MakeSymbol("preclude")
var symshen_4member_1clause = MakeSymbol("shen.member-clause")
var symshen_4_5semantics_6 = MakeSymbol("shen.<semantics>")
var symshen_4alpha_2 = MakeSymbol("shen.alpha?")
var symversion = MakeSymbol("version")
var symshen_4read_1file_1as_1bytelist_1help = MakeSymbol("shen.read-file-as-bytelist-help")
var symabsvector_2 = MakeSymbol("absvector?")
var symshen_4x_4launcher_4execute_1all = MakeSymbol("shen.x.launcher.execute-all")
var symshen_4process_1application = MakeSymbol("shen.process-application")
var symshen_4vector_1parameter = MakeSymbol("shen.vector-parameter")
var symshen_4profile_1func = MakeSymbol("shen.profile-func")
var symmap = MakeSymbol("map")
var symfresh = MakeSymbol("fresh")
var symshen_4_5number_6 = MakeSymbol("shen.<number>")
var symshen_4prompt = MakeSymbol("shen.prompt")
var symshen_4update_1assoc = MakeSymbol("shen.update-assoc")
var symshen_4stpart = MakeSymbol("shen.stpart")
var symshen_4length_1h = MakeSymbol("shen.length-h")
var symshen_4comb = MakeSymbol("shen.comb")
var symshen_4execute_1store_1arity = MakeSymbol("shen.execute-store-arity")
var symshen_4process_1lambda = MakeSymbol("shen.process-lambda")
var symrelease = MakeSymbol("release")
var symshen_4_5sym_6 = MakeSymbol("shen.<sym>")
var symvar_2 = MakeSymbol("var?")
var symretract = MakeSymbol("retract")
var symshen_4non_1terminalcode = MakeSymbol("shen.non-terminalcode")
var symshen_4_5whitespaces_6 = MakeSymbol("shen.<whitespaces>")
var symshen_4demod = MakeSymbol("shen.demod")
var symshen_4freeze_1literals = MakeSymbol("shen.freeze-literals")
var symshen_4gc = MakeSymbol("shen.gc")
var symshen_4typecheck = MakeSymbol("shen.typecheck")
var symshen_4process_1synonyms = MakeSymbol("shen.process-synonyms")
var symshen_4sng_1h_2 = MakeSymbol("shen.sng-h?")
var symshen_4_5syntax_6 = MakeSymbol("shen.<syntax>")
var symshen_4_5non_1terminal_1name_6 = MakeSymbol("shen.<non-terminal-name>")
var symshen_4record_1and_1evaluate = MakeSymbol("shen.record-and-evaluate")
var sym_5 = MakeSymbol("<")
var symshen_4_5control_6 = MakeSymbol("shen.<control>")
var symshen_4insert_1info = MakeSymbol("shen.insert-info")
var symshen_4char_1stoutput_2 = MakeSymbol("shen.char-stoutput?")
var symdefprolog = MakeSymbol("defprolog")
var symshen_4a = MakeSymbol("shen.a")
var symshen_4_5numeral_6 = MakeSymbol("shen.<numeral>")
var symstream = MakeSymbol("stream")
var symshen_4_5formula_6 = MakeSymbol("shen.<formula>")
var symnull = MakeSymbol("null")
var symimplementation = MakeSymbol("implementation")
var symunit = MakeSymbol("unit")
var symshen_4prolog_1fbody = MakeSymbol("shen.prolog-fbody")
var symhdv = MakeSymbol("hdv")
var syminferences = MakeSymbol("inferences")
var sym_dhush_d = MakeSymbol("*hush*")
var symtrack = MakeSymbol("track")
var symprofile_1results = MakeSymbol("profile-results")
var symmake_1string = MakeSymbol("make-string")
var symhash = MakeSymbol("hash")
var symshen_4myassume = MakeSymbol("shen.myassume")
var symshen_4magless = MakeSymbol("shen.magless")
var symsnd = MakeSymbol("snd")
var symshen_4foreign_2 = MakeSymbol("shen.foreign?")
var symstring_1_6n = MakeSymbol("string->n")
var sympreclude_1all_1but = MakeSymbol("preclude-all-but")
var symfindall = MakeSymbol("findall")
var symshen_4overbind = MakeSymbol("shen.overbind")
var symshen_4_5conc_6 = MakeSymbol("shen.<conc>")
var symshen_4shendef_1_6kldef = MakeSymbol("shen.shendef->kldef")
var symdefine = MakeSymbol("define")
var symResult = MakeSymbol("Result")
var symshen_4_5multiline_6 = MakeSymbol("shen.<multiline>")
var symshen_4_dloading_2_d = MakeSymbol("shen.*loading?*")
var symshen_4yacc_1semantics = MakeSymbol("shen.yacc-semantics")
var syminteger_2 = MakeSymbol("integer?")
var symNewV = MakeSymbol("NewV")
var symshen_4_5backslash_6 = MakeSymbol("shen.<backslash>")
var sympos = MakeSymbol("pos")
var symshen_4freshterm_2 = MakeSymbol("shen.freshterm?")
var sym_dos_d = MakeSymbol("*os*")
var symabsvector = MakeSymbol("absvector")
var symStart = MakeSymbol("Start")
var symshen_4string_1prefix_2 = MakeSymbol("shen.string-prefix?")
var symshen_4newname = MakeSymbol("shen.newname")
var symshen_4openlock = MakeSymbol("shen.openlock")
var symshen_4posint_2 = MakeSymbol("shen.posint?")
var symshen_4find_1arity = MakeSymbol("shen.find-arity")
var symshen_4op2 = MakeSymbol("shen.op2")
var symshen_4_ddatatypes_d = MakeSymbol("shen.*datatypes*")
var symshen_4klfile = MakeSymbol("shen.klfile")
var symshen_4_5hterm2_6 = MakeSymbol("shen.<hterm2>")
var symshen_4_5side_6 = MakeSymbol("shen.<side>")
var symshen_4dbl_1h_2 = MakeSymbol("shen.dbl-h?")
var symshen_4rcons__form = MakeSymbol("shen.rcons_form")
var symshen_4extraspecial_2 = MakeSymbol("shen.extraspecial?")
var symtime = MakeSymbol("time")
var symshen_4_5prems_6 = MakeSymbol("shen.<prems>")
var symshen_4track_1function = MakeSymbol("shen.track-function")
var symstring_1_6symbol = MakeSymbol("string->symbol")
var symsymbol = MakeSymbol("symbol")
var symshen_4_5alpha_6 = MakeSymbol("shen.<alpha>")
var sym_1_6 = MakeSymbol("->")
var symshen_4triple_1stack = MakeSymbol("shen.triple-stack")
var symshen_4byte_1_6digit = MakeSymbol("shen.byte->digit")
var symintersection = MakeSymbol("intersection")
var symshen_4map_1h = MakeSymbol("shen.map-h")
var symshen_4yacc_1syntax = MakeSymbol("shen.yacc-syntax")
var symshen_4fn_1print = MakeSymbol("shen.fn-print")
var symshen_4_5singleline_6 = MakeSymbol("shen.<singleline>")
var symassertz = MakeSymbol("assertz")
var symshen_4occurs_1check_2 = MakeSymbol("shen.occurs-check?")
var symshen_4member = MakeSymbol("shen.member")
var symshen_4cons_1form = MakeSymbol("shen.cons-form")
var symmacroexpand = MakeSymbol("macroexpand")
var symdifference = MakeSymbol("difference")
var symos = MakeSymbol("os")
var symshen_4mkstr = MakeSymbol("shen.mkstr")
var sym_5_b_6 = MakeSymbol("<!>")
var symshen_4_5dbq_6 = MakeSymbol("shen.<dbq>")
var syminclude_1all_1but = MakeSymbol("include-all-but")
var symshen_4string_1match = MakeSymbol("shen.string-match")
var symshen_4show_1assumptions = MakeSymbol("shen.show-assumptions")
var symlimit = MakeSymbol("limit")
var symshen_4assoc_1_6 = MakeSymbol("shen.assoc->")
var sym_b = MakeSymbol("!")
var symshen_4rule_1_6head = MakeSymbol("shen.rule->head")
var symshen_4c_1rules_1_6shen = MakeSymbol("shen.c-rules->shen")
var symshen_4whitespace_2 = MakeSymbol("shen.whitespace?")
var symshen_4loading_2 = MakeSymbol("shen.loading?")
var symloaded = MakeSymbol("loaded")
var symshen_4freshen_1sig = MakeSymbol("shen.freshen-sig")
var symshen_4_5packagename_6 = MakeSymbol("shen.<packagename>")
var symshen_4decons = MakeSymbol("shen.decons")
var symshen_4_5sc_6 = MakeSymbol("shen.<sc>")
var symshen_4x_4launcher_4eval_1command_1h = MakeSymbol("shen.x.launcher.eval-command-h")
