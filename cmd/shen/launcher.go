package main

import . "github.com/pyrex41/shen-go/kl"

var LauncherMain = MakeNative(func(__e *ControlFlow) {
_ = MakeString("Copyright (c) 2019 Bruno Deferrari.\nBSD 3-Clause License: http://opensource.org/licenses/BSD-3-Clause")

tmp22041 := MakeNative(func(__e *ControlFlow) {
V7337 := __e.Get(1)
_ = V7337
tmp22042 := Call(__e, PrimFunc(symread_1file), V7337)


let__22037 := tmp22042
_ = let__22037

tmp22043 := MakeNative(func(__e *ControlFlow) {
Z7339 := __e.Get(1)
_ = Z7339
__e.TailApply(PrimFunc(symeval), Z7339)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp22043, let__22037)
return


}, 1)

tmp22044 := Call(__e, ns2_1set, symshen_4x_4launcher_4quiet_1load, tmp22041)


_ = tmp22044

tmp22045 := MakeNative(func(__e *ControlFlow) {
tmp22046 := Call(__e, PrimFunc(symversion))


tmp22047 := Call(__e, PrimFunc(symlanguage))


tmp22048 := Call(__e, PrimFunc(symport))


tmp22049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp22048, Nil)
}
__typedArg0 := tmp22048
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp22047, tmp22049)
}
__typedArg0 := tmp22047
__typedArg1 := tmp22049
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22051 := Call(__e, PrimFunc(symimplementation))


tmp22052 := Call(__e, PrimFunc(symrelease))


tmp22053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp22052, Nil)
}
__typedArg0 := tmp22052
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp22051, tmp22053)
}
__typedArg0 := tmp22051
__typedArg1 := tmp22053
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp22054, Nil)
}
__typedArg0 := tmp22054
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symimplementation, tmp22055)
}
__typedArg0 := symimplementation
__typedArg1 := tmp22055
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp22050, tmp22056)
}
__typedArg0 := tmp22050
__typedArg1 := tmp22056
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symport, tmp22057)
}
__typedArg0 := symport
__typedArg1 := tmp22057
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22059 := Call(__e, PrimFunc(symshen_4app), tmp22058, MakeString("\n"), symshen_4r)


__e.TailApply(PrimFunc(symshen_4app), tmp22046, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" "))
__typedS1, __typedOK1 := TypedString(tmp22059)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" ")
__typedArg1 := tmp22059
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)
return


}, 0)

tmp22061 := Call(__e, ns2_1set, symshen_4x_4launcher_4version_1string, tmp22045)


_ = tmp22061

tmp22062 := MakeNative(func(__e *ControlFlow) {
V7340 := __e.Get(1)
_ = V7340
tmp22063 := Call(__e, PrimFunc(symshen_4app), V7340, MakeString(" [--version] [--help] <COMMAND> [<ARGS>]\n\ncommands:\n    repl\n        Launches the interactive REPL.\n        Default action if no command is supplied.\n\n    script <FILE> [<ARGS>]\n        Runs the script in FILE. *argv* is set to [FILE | ARGS].\n\n    eval <ARGS>\n        Evaluates expressions and files. ARGS are evaluated from\n        left to right and can be a combination of:\n            -e, --eval <EXPR>\n                Evaluates EXPR and prints result.\n            -l, --load <FILE>\n                Reads and evaluates FILE.\n            -q, --quiet\n                Silences interactive output.\n            -s, --set <KEY> <VALUE>\n                Evaluates KEY, VALUE and sets as global.\n            -r, --repl\n                Launches the interactive REPL after evaluating\n                all the previous expresions."), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Usage: "))
__typedS1, __typedOK1 := TypedString(tmp22063)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Usage: ")
__typedArg1 := tmp22063
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp22064 := Call(__e, ns2_1set, symshen_4x_4launcher_4help_1text, tmp22062)


_ = tmp22064

tmp22065 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V7341 := __self1
_ = V7341
tmp22072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V7341)
}
__typedArg0 := Nil
__typedArg1 := V7341
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp22072 {
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
tmp22070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7341)
}
__typedArg0 := V7341
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp22070 {
tmp22066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7341)
}
__typedArg0 := V7341
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22067 := Call(__e, PrimFunc(symthaw), tmp22066)


_ = tmp22067

tmp22068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7341)
}
__typedArg0 := V7341
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4x_4launcher_4execute_1all) == __self {
__self1 = tmp22068
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4x_4launcher_4execute_1all), tmp22068)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4x_4launcher_4execute_1all)
return
}


}


}, 1)

tmp22073 := Call(__e, ns2_1set, symshen_4x_4launcher_4execute_1all, tmp22065)


_ = tmp22073

tmp22074 := MakeNative(func(__e *ControlFlow) {
V7342 := __e.Get(1)
_ = V7342
tmp22075 := Call(__e, PrimFunc(symread_1from_1string), V7342)


tmp22076 := Call(__e, PrimFunc(symhead), tmp22075)


__e.TailApply(PrimFunc(symeval), tmp22076)
return


}, 1)

tmp22077 := Call(__e, ns2_1set, symshen_4x_4launcher_4eval_1string, tmp22074)


_ = tmp22077

tmp22078 := MakeNative(func(__e *ControlFlow) {
V7345 := __e.Get(1)
_ = V7345
tmp22088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("-e"), V7345)
}
__typedArg0 := MakeString("-e")
__typedArg1 := V7345
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp22088 {
__e.Return(MakeString("--eval"))
return
} else {
tmp22086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("-l"), V7345)
}
__typedArg0 := MakeString("-l")
__typedArg1 := V7345
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp22086 {
__e.Return(MakeString("--load"))
return
} else {
tmp22084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("-q"), V7345)
}
__typedArg0 := MakeString("-q")
__typedArg1 := V7345
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp22084 {
__e.Return(MakeString("--quiet"))
return
} else {
tmp22082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("-s"), V7345)
}
__typedArg0 := MakeString("-s")
__typedArg1 := V7345
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp22082 {
__e.Return(MakeString("--set"))
return
} else {
tmp22080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("-r"), V7345)
}
__typedArg0 := MakeString("-r")
__typedArg1 := V7345
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp22080 {
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

tmp22089 := Call(__e, ns2_1set, symshen_4x_4launcher_4eval_1flag_1map, tmp22078)


_ = tmp22089

tmp22090 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V7350 := __self1
_ = V7350
V7351 := __self2
_ = V7351
tmp22192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V7350)
}
__typedArg0 := Nil
__typedArg1 := V7350
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp22192 {
tmp22091 := Call(__e, PrimFunc(symreverse), V7351)


__e.TailApply(PrimFunc(symshen_4x_4launcher_4execute_1all), tmp22091)
return


} else {
tmp22190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22182 Obj

if True == tmp22190 {
tmp22188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--eval"), tmp22188)
}
__typedArg0 := MakeString("--eval")
__typedArg1 := tmp22188
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22184 Obj

if True == tmp22189 {
tmp22186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22186)
}
__typedArg0 := tmp22186
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22185 Obj

if True == tmp22187 {
ifres22185 = True


} else {
ifres22185 = False


}

ifres22184 = ifres22185


} else {
ifres22184 = False


}

var ifres22183 Obj

if True == ifres22184 {
ifres22183 = True


} else {
ifres22183 = False


}

ifres22182 = ifres22183


} else {
ifres22182 = False


}

if True == ifres22182 {
tmp22092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22092)
}
__typedArg0 := tmp22092
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22094 := MakeNative(func(__e *ControlFlow) {
tmp22095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22095)
}
__typedArg0 := tmp22095
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22097 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1string), tmp22096)


tmp22098 := Call(__e, PrimFunc(symshen_4app), tmp22097, MakeString("\n"), symshen_4a)


tmp22099 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp22098, tmp22099)
return


}, 0)

tmp22100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp22094, V7351)
}
__typedArg0 := tmp22094
__typedArg1 := V7351
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4x_4launcher_4eval_1command_1h) == __self {
__self1, __self2 = tmp22093, tmp22100
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp22093, tmp22100)
return


} else {
tmp22180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22172 Obj

if True == tmp22180 {
tmp22178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--load"), tmp22178)
}
__typedArg0 := MakeString("--load")
__typedArg1 := tmp22178
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22174 Obj

if True == tmp22179 {
tmp22176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22176)
}
__typedArg0 := tmp22176
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22175 Obj

if True == tmp22177 {
ifres22175 = True


} else {
ifres22175 = False


}

ifres22174 = ifres22175


} else {
ifres22174 = False


}

var ifres22173 Obj

if True == ifres22174 {
ifres22173 = True


} else {
ifres22173 = False


}

ifres22172 = ifres22173


} else {
ifres22172 = False


}

if True == ifres22172 {
tmp22101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22101)
}
__typedArg0 := tmp22101
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22103 := MakeNative(func(__e *ControlFlow) {
tmp22104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22104)
}
__typedArg0 := tmp22104
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symload), tmp22105)
return


}, 0)

tmp22106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp22103, V7351)
}
__typedArg0 := tmp22103
__typedArg1 := V7351
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4x_4launcher_4eval_1command_1h) == __self {
__self1, __self2 = tmp22102, tmp22106
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp22102, tmp22106)
return


} else {
tmp22170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22166 Obj

if True == tmp22170 {
tmp22168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--quiet"), tmp22168)
}
__typedArg0 := MakeString("--quiet")
__typedArg1 := tmp22168
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22167 Obj

if True == tmp22169 {
ifres22167 = True


} else {
ifres22167 = False


}

ifres22166 = ifres22167


} else {
ifres22166 = False


}

if True == ifres22166 {
tmp22107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22108 := MakeNative(func(__e *ControlFlow) {
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

tmp22109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp22108, V7351)
}
__typedArg0 := tmp22108
__typedArg1 := V7351
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4x_4launcher_4eval_1command_1h) == __self {
__self1, __self2 = tmp22107, tmp22109
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp22107, tmp22109)
return


} else {
tmp22164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22151 Obj

if True == tmp22164 {
tmp22162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--set"), tmp22162)
}
__typedArg0 := MakeString("--set")
__typedArg1 := tmp22162
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22153 Obj

if True == tmp22163 {
tmp22160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22160)
}
__typedArg0 := tmp22160
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22155 Obj

if True == tmp22161 {
tmp22157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22157)
}
__typedArg0 := tmp22157
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22158)
}
__typedArg0 := tmp22158
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22156 Obj

if True == tmp22159 {
ifres22156 = True


} else {
ifres22156 = False


}

ifres22155 = ifres22156


} else {
ifres22155 = False


}

var ifres22154 Obj

if True == ifres22155 {
ifres22154 = True


} else {
ifres22154 = False


}

ifres22153 = ifres22154


} else {
ifres22153 = False


}

var ifres22152 Obj

if True == ifres22153 {
ifres22152 = True


} else {
ifres22152 = False


}

ifres22151 = ifres22152


} else {
ifres22151 = False


}

if True == ifres22151 {
tmp22110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22110)
}
__typedArg0 := tmp22110
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22111)
}
__typedArg0 := tmp22111
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22113 := MakeNative(func(__e *ControlFlow) {
tmp22114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22114)
}
__typedArg0 := tmp22114
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22116 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1string), tmp22115)


tmp22117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22117)
}
__typedArg0 := tmp22117
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22118)
}
__typedArg0 := tmp22118
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22120 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1string), tmp22119)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(tmp22116, tmp22120)
}
__typedArg0 := tmp22116
__typedArg1 := tmp22120
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp22121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp22113, V7351)
}
__typedArg0 := tmp22113
__typedArg1 := V7351
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4x_4launcher_4eval_1command_1h) == __self {
__self1, __self2 = tmp22112, tmp22121
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp22112, tmp22121)
return


} else {
tmp22149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22145 Obj

if True == tmp22149 {
tmp22147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--repl"), tmp22147)
}
__typedArg0 := MakeString("--repl")
__typedArg1 := tmp22147
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22146 Obj

if True == tmp22148 {
ifres22146 = True


} else {
ifres22146 = False


}

ifres22145 = ifres22146


} else {
ifres22145 = False


}

if True == ifres22145 {
tmp22122 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1command_1h), Nil, V7351)


_ = tmp22122

tmp22123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlaunch_1repl, tmp22123)
}
__typedArg0 := symlaunch_1repl
__typedArg1 := tmp22123
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp22124 := MakeNative(func(__e *ControlFlow) {
tmp22130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp22130 {
tmp22125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22126 := Call(__e, PrimFunc(symshen_4app), tmp22125, MakeString(""), symshen_4a)


tmp22128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Invalid eval argument: "))
__typedS1, __typedOK1 := TypedString(tmp22126)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Invalid eval argument: ")
__typedArg1 := tmp22126
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), Nil)
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Invalid eval argument: "))
__typedS1, __typedOK1 := TypedString(tmp22126)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Invalid eval argument: ")
__typedArg1 := tmp22126
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symerror, tmp22128)
}
__typedArg0 := symerror
__typedArg1 := tmp22128
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4x_4launcher_4eval_1command_1h)
return
}


}, 0)

let__22038 := tmp22124
_ = let__22038

tmp22143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp22143 {
tmp22131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22132 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1flag_1map), tmp22131)


let__22040 := tmp22132
_ = let__22040

tmp22138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, let__22040)
}
__typedArg0 := False
__typedArg1 := let__22040
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22133 Obj

if True == tmp22138 {
tmp22134 := Call(__e, PrimFunc(symfail))


ifres22133 = tmp22134


} else {
tmp22135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7350)
}
__typedArg0 := V7350
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__22040, tmp22135)
}
__typedArg0 := let__22040
__typedArg1 := tmp22135
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22137 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp22136, V7351)


ifres22133 = tmp22137


}

let__22039 := ifres22133
_ = let__22039

tmp22140 := Call(__e, PrimFunc(symfail))


tmp22141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__22039, tmp22140)
}
__typedArg0 := let__22039
__typedArg1 := tmp22140
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp22141 {
__e.TailApply(PrimFunc(symthaw), let__22038)
return
} else {
__e.Return(let__22039)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__22038)
return
}


}


}


}


}


}


}


}, 2)

tmp22193 := Call(__e, ns2_1set, symshen_4x_4launcher_4eval_1command_1h, tmp22090)


_ = tmp22193

tmp22194 := MakeNative(func(__e *ControlFlow) {
V7355 := __e.Get(1)
_ = V7355
__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), V7355, Nil)
return
}, 1)

tmp22195 := Call(__e, ns2_1set, symshen_4x_4launcher_4eval_1command, tmp22194)


_ = tmp22195

tmp22196 := MakeNative(func(__e *ControlFlow) {
V7356 := __e.Get(1)
_ = V7356
V7357 := __e.Get(2)
_ = V7357
tmp22197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V7356, V7357)
}
__typedArg0 := V7356
__typedArg1 := V7357
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dargv_d, tmp22197)
}
__typedArg0 := sym_dargv_d
__typedArg1 := tmp22197
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp22198

tmp22199 := Call(__e, PrimFunc(symshen_4x_4launcher_4quiet_1load), V7356)


_ = tmp22199

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

tmp22200 := Call(__e, ns2_1set, symshen_4x_4launcher_4script_1command, tmp22196)


_ = tmp22200

tmp22201 := MakeNative(func(__e *ControlFlow) {
V7358 := __e.Get(1)
_ = V7358
tmp22288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22284 Obj

if True == tmp22288 {
tmp22286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp22286)
}
__typedArg0 := Nil
__typedArg1 := tmp22286
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22285 Obj

if True == tmp22287 {
ifres22285 = True


} else {
ifres22285 = False


}

ifres22284 = ifres22285


} else {
ifres22284 = False


}

if True == ifres22284 {
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
tmp22282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22273 Obj

if True == tmp22282 {
tmp22280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22280)
}
__typedArg0 := tmp22280
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22275 Obj

if True == tmp22281 {
tmp22277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22277)
}
__typedArg0 := tmp22277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--help"), tmp22278)
}
__typedArg0 := MakeString("--help")
__typedArg1 := tmp22278
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22276 Obj

if True == tmp22279 {
ifres22276 = True


} else {
ifres22276 = False


}

ifres22275 = ifres22276


} else {
ifres22275 = False


}

var ifres22274 Obj

if True == ifres22275 {
ifres22274 = True


} else {
ifres22274 = False


}

ifres22273 = ifres22274


} else {
ifres22273 = False


}

if True == ifres22273 {
tmp22202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22203 := Call(__e, PrimFunc(symshen_4x_4launcher_4help_1text), tmp22202)


tmp22204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp22203, Nil)
}
__typedArg0 := tmp22203
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshow_1help, tmp22204)
}
__typedArg0 := symshow_1help
__typedArg1 := tmp22204
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp22271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22262 Obj

if True == tmp22271 {
tmp22269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22269)
}
__typedArg0 := tmp22269
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22264 Obj

if True == tmp22270 {
tmp22266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22266)
}
__typedArg0 := tmp22266
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--version"), tmp22267)
}
__typedArg0 := MakeString("--version")
__typedArg1 := tmp22267
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22265 Obj

if True == tmp22268 {
ifres22265 = True


} else {
ifres22265 = False


}

ifres22264 = ifres22265


} else {
ifres22264 = False


}

var ifres22263 Obj

if True == ifres22264 {
ifres22263 = True


} else {
ifres22263 = False


}

ifres22262 = ifres22263


} else {
ifres22262 = False


}

if True == ifres22262 {
tmp22205 := Call(__e, PrimFunc(symshen_4x_4launcher_4version_1string))


tmp22206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp22205, Nil)
}
__typedArg0 := tmp22205
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsuccess, tmp22206)
}
__typedArg0 := symsuccess
__typedArg1 := tmp22206
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp22260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22251 Obj

if True == tmp22260 {
tmp22258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22258)
}
__typedArg0 := tmp22258
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22253 Obj

if True == tmp22259 {
tmp22255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22255)
}
__typedArg0 := tmp22255
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("repl"), tmp22256)
}
__typedArg0 := MakeString("repl")
__typedArg1 := tmp22256
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22254 Obj

if True == tmp22257 {
ifres22254 = True


} else {
ifres22254 = False


}

ifres22253 = ifres22254


} else {
ifres22253 = False


}

var ifres22252 Obj

if True == ifres22253 {
ifres22252 = True


} else {
ifres22252 = False


}

ifres22251 = ifres22252


} else {
ifres22251 = False


}

if True == ifres22251 {
tmp22207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22207)
}
__typedArg0 := tmp22207
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlaunch_1repl, tmp22208)
}
__typedArg0 := symlaunch_1repl
__typedArg1 := tmp22208
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp22249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22235 Obj

if True == tmp22249 {
tmp22247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22247)
}
__typedArg0 := tmp22247
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22237 Obj

if True == tmp22248 {
tmp22244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22244)
}
__typedArg0 := tmp22244
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("script"), tmp22245)
}
__typedArg0 := MakeString("script")
__typedArg1 := tmp22245
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22239 Obj

if True == tmp22246 {
tmp22241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22241)
}
__typedArg0 := tmp22241
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22242)
}
__typedArg0 := tmp22242
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22240 Obj

if True == tmp22243 {
ifres22240 = True


} else {
ifres22240 = False


}

ifres22239 = ifres22240


} else {
ifres22239 = False


}

var ifres22238 Obj

if True == ifres22239 {
ifres22238 = True


} else {
ifres22238 = False


}

ifres22237 = ifres22238


} else {
ifres22237 = False


}

var ifres22236 Obj

if True == ifres22237 {
ifres22236 = True


} else {
ifres22236 = False


}

ifres22235 = ifres22236


} else {
ifres22235 = False


}

if True == ifres22235 {
tmp22209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22209)
}
__typedArg0 := tmp22209
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22210)
}
__typedArg0 := tmp22210
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22212)
}
__typedArg0 := tmp22212
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22213)
}
__typedArg0 := tmp22213
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4x_4launcher_4script_1command), tmp22211, tmp22214)
return


} else {
tmp22233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22224 Obj

if True == tmp22233 {
tmp22231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22231)
}
__typedArg0 := tmp22231
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22226 Obj

if True == tmp22232 {
tmp22228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22228)
}
__typedArg0 := tmp22228
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("eval"), tmp22229)
}
__typedArg0 := MakeString("eval")
__typedArg1 := tmp22229
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22227 Obj

if True == tmp22230 {
ifres22227 = True


} else {
ifres22227 = False


}

ifres22226 = ifres22227


} else {
ifres22226 = False


}

var ifres22225 Obj

if True == ifres22226 {
ifres22225 = True


} else {
ifres22225 = False


}

ifres22224 = ifres22225


} else {
ifres22224 = False


}

if True == ifres22224 {
tmp22215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22215)
}
__typedArg0 := tmp22215
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command), tmp22216)
return


} else {
tmp22222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22218 Obj

if True == tmp22222 {
tmp22220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7358)
}
__typedArg0 := V7358
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22220)
}
__typedArg0 := tmp22220
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22219 Obj

if True == tmp22221 {
ifres22219 = True


} else {
ifres22219 = False


}

ifres22218 = ifres22219


} else {
ifres22218 = False


}

if True == ifres22218 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunknown_1arguments, V7358)
}
__typedArg0 := symunknown_1arguments
__typedArg1 := V7358
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

tmp22289 := Call(__e, ns2_1set, symshen_4x_4launcher_4launch_1shen, tmp22201)


_ = tmp22289

tmp22290 := MakeNative(func(__e *ControlFlow) {
V7361 := __e.Get(1)
_ = V7361
tmp22389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22381 Obj

if True == tmp22389 {
tmp22387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsuccess, tmp22387)
}
__typedArg0 := symsuccess
__typedArg1 := tmp22387
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22383 Obj

if True == tmp22388 {
tmp22385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp22385)
}
__typedArg0 := Nil
__typedArg1 := tmp22385
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22384 Obj

if True == tmp22386 {
ifres22384 = True


} else {
ifres22384 = False


}

ifres22383 = ifres22384


} else {
ifres22383 = False


}

var ifres22382 Obj

if True == ifres22383 {
ifres22382 = True


} else {
ifres22382 = False


}

ifres22381 = ifres22382


} else {
ifres22381 = False


}

if True == ifres22381 {
__e.Return(symshen_4x_4launcher_4done)
return
} else {
tmp22379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22366 Obj

if True == tmp22379 {
tmp22377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsuccess, tmp22377)
}
__typedArg0 := symsuccess
__typedArg1 := tmp22377
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22368 Obj

if True == tmp22378 {
tmp22375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22375)
}
__typedArg0 := tmp22375
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22370 Obj

if True == tmp22376 {
tmp22372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22372)
}
__typedArg0 := tmp22372
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp22373)
}
__typedArg0 := Nil
__typedArg1 := tmp22373
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22371 Obj

if True == tmp22374 {
ifres22371 = True


} else {
ifres22371 = False


}

ifres22370 = ifres22371


} else {
ifres22370 = False


}

var ifres22369 Obj

if True == ifres22370 {
ifres22369 = True


} else {
ifres22369 = False


}

ifres22368 = ifres22369


} else {
ifres22368 = False


}

var ifres22367 Obj

if True == ifres22368 {
ifres22367 = True


} else {
ifres22367 = False


}

ifres22366 = ifres22367


} else {
ifres22366 = False


}

if True == ifres22366 {
tmp22291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22291)
}
__typedArg0 := tmp22291
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22293 := Call(__e, PrimFunc(symshen_4app), tmp22292, MakeString("\n"), symshen_4a)


tmp22294 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp22293, tmp22294)
return


} else {
tmp22364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22351 Obj

if True == tmp22364 {
tmp22362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symerror, tmp22362)
}
__typedArg0 := symerror
__typedArg1 := tmp22362
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22353 Obj

if True == tmp22363 {
tmp22360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22360)
}
__typedArg0 := tmp22360
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22355 Obj

if True == tmp22361 {
tmp22357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22357)
}
__typedArg0 := tmp22357
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp22358)
}
__typedArg0 := Nil
__typedArg1 := tmp22358
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22356 Obj

if True == tmp22359 {
ifres22356 = True


} else {
ifres22356 = False


}

ifres22355 = ifres22356


} else {
ifres22355 = False


}

var ifres22354 Obj

if True == ifres22355 {
ifres22354 = True


} else {
ifres22354 = False


}

ifres22353 = ifres22354


} else {
ifres22353 = False


}

var ifres22352 Obj

if True == ifres22353 {
ifres22352 = True


} else {
ifres22352 = False


}

ifres22351 = ifres22352


} else {
ifres22351 = False


}

if True == ifres22351 {
tmp22295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22295)
}
__typedArg0 := tmp22295
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22297 := Call(__e, PrimFunc(symshen_4app), tmp22296, MakeString("\n"), symshen_4a)


tmp22298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("ERROR: "))
__typedS1, __typedOK1 := TypedString(tmp22297)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("ERROR: ")
__typedArg1 := tmp22297
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp22299 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp22298, tmp22299)
return


} else {
tmp22349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22345 Obj

if True == tmp22349 {
tmp22347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlaunch_1repl, tmp22347)
}
__typedArg0 := symlaunch_1repl
__typedArg1 := tmp22347
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22346 Obj

if True == tmp22348 {
ifres22346 = True


} else {
ifres22346 = False


}

ifres22345 = ifres22346


} else {
ifres22345 = False


}

if True == ifres22345 {
__e.TailApply(PrimFunc(symshen_4repl))
return
} else {
tmp22343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22330 Obj

if True == tmp22343 {
tmp22341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshow_1help, tmp22341)
}
__typedArg0 := symshow_1help
__typedArg1 := tmp22341
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22332 Obj

if True == tmp22342 {
tmp22339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22339)
}
__typedArg0 := tmp22339
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22334 Obj

if True == tmp22340 {
tmp22336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22336)
}
__typedArg0 := tmp22336
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp22337)
}
__typedArg0 := Nil
__typedArg1 := tmp22337
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22335 Obj

if True == tmp22338 {
ifres22335 = True


} else {
ifres22335 = False


}

ifres22334 = ifres22335


} else {
ifres22334 = False


}

var ifres22333 Obj

if True == ifres22334 {
ifres22333 = True


} else {
ifres22333 = False


}

ifres22332 = ifres22333


} else {
ifres22332 = False


}

var ifres22331 Obj

if True == ifres22332 {
ifres22331 = True


} else {
ifres22331 = False


}

ifres22330 = ifres22331


} else {
ifres22330 = False


}

if True == ifres22330 {
tmp22300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22300)
}
__typedArg0 := tmp22300
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22302 := Call(__e, PrimFunc(symshen_4app), tmp22301, MakeString("\n"), symshen_4a)


tmp22303 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp22302, tmp22303)
return


} else {
tmp22328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22315 Obj

if True == tmp22328 {
tmp22326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symunknown_1arguments, tmp22326)
}
__typedArg0 := symunknown_1arguments
__typedArg1 := tmp22326
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22317 Obj

if True == tmp22327 {
tmp22324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22324)
}
__typedArg0 := tmp22324
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22319 Obj

if True == tmp22325 {
tmp22321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22321)
}
__typedArg0 := tmp22321
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22322)
}
__typedArg0 := tmp22322
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22320 Obj

if True == tmp22323 {
ifres22320 = True


} else {
ifres22320 = False


}

ifres22319 = ifres22320


} else {
ifres22319 = False


}

var ifres22318 Obj

if True == ifres22319 {
ifres22318 = True


} else {
ifres22318 = False


}

ifres22317 = ifres22318


} else {
ifres22317 = False


}

var ifres22316 Obj

if True == ifres22317 {
ifres22316 = True


} else {
ifres22316 = False


}

ifres22315 = ifres22316


} else {
ifres22315 = False


}

if True == ifres22315 {
tmp22304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22304)
}
__typedArg0 := tmp22304
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22305)
}
__typedArg0 := tmp22305
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7361)
}
__typedArg0 := V7361
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22307)
}
__typedArg0 := tmp22307
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22309 := Call(__e, PrimFunc(symshen_4app), tmp22308, MakeString(" --help' for more information.\n"), symshen_4a)


tmp22311 := Call(__e, PrimFunc(symshen_4app), tmp22306, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\nTry `"))
__typedS1, __typedOK1 := TypedString(tmp22309)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\nTry `")
__typedArg1 := tmp22309
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp22312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("ERROR: Invalid argument: "))
__typedS1, __typedOK1 := TypedString(tmp22311)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("ERROR: Invalid argument: ")
__typedArg1 := tmp22311
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp22313 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp22312, tmp22313)
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

tmp22390 := Call(__e, ns2_1set, symshen_4x_4launcher_4default_1handle_1result, tmp22290)


_ = tmp22390

tmp22391 := MakeNative(func(__e *ControlFlow) {
V7362 := __e.Get(1)
_ = V7362
tmp22392 := Call(__e, PrimFunc(symshen_4x_4launcher_4launch_1shen), V7362)


__e.TailApply(PrimFunc(symshen_4x_4launcher_4default_1handle_1result), tmp22392)
return


}, 1)

__e.TailApply(ns2_1set, symshen_4x_4launcher_4main, tmp22391)
return




}, 0)

var sym_b = MakeSymbol("!")
var sym_3 = MakeSymbol("$")
var sym_e = MakeSymbol("&")
var sym_e_e = MakeSymbol("&&")
var sym_d = MakeSymbol("*")
var sym_dabsolute_d = MakeSymbol("*absolute*")
var sym_dargv_d = MakeSymbol("*argv*")
var sym_dhome_1directory_d = MakeSymbol("*home-directory*")
var sym_dhush_d = MakeSymbol("*hush*")
var sym_dimplementation_d = MakeSymbol("*implementation*")
var sym_dlanguage_d = MakeSymbol("*language*")
var sym_dmacros_d = MakeSymbol("*macros*")
var sym_dmaximum_1print_1sequence_1size_d = MakeSymbol("*maximum-print-sequence-size*")
var sym_dos_d = MakeSymbol("*os*")
var sym_dport_d = MakeSymbol("*port*")
var sym_dporters_d = MakeSymbol("*porters*")
var sym_dproperty_1vector_d = MakeSymbol("*property-vector*")
var sym_drelease_d = MakeSymbol("*release*")
var sym_dsterror_d = MakeSymbol("*sterror*")
var sym_dstinput_d = MakeSymbol("*stinput*")
var sym_dstoutput_d = MakeSymbol("*stoutput*")
var sym_dversion_d = MakeSymbol("*version*")
var sym_7 = MakeSymbol("+")
var sym_1 = MakeSymbol("-")
var sym_1_1_6 = MakeSymbol("-->")
var sym_1_6 = MakeSymbol("->")
var sym_c = MakeSymbol("/")
var sym_c_4 = MakeSymbol("/.")
var sym_5 = MakeSymbol("<")
var sym_5_b_6 = MakeSymbol("<!>")
var sym_5_1 = MakeSymbol("<-")
var sym_5_1_1 = MakeSymbol("<--")
var sym_5_1address = MakeSymbol("<-address")
var sym_5_1vector = MakeSymbol("<-vector")
var sym_5_a = MakeSymbol("<=")
var sym_5e_6 = MakeSymbol("<e>")
var sym_5end_6 = MakeSymbol("<end>")
var sym_a = MakeSymbol("=")
var sym_a_a = MakeSymbol("==")
var sym_a_a_6 = MakeSymbol("==>")
var sym_6 = MakeSymbol(">")
var sym_6_a = MakeSymbol(">=")
var sym_6_6 = MakeSymbol(">>")
var sym_8p = MakeSymbol("@p")
var sym_8s = MakeSymbol("@s")
var sym_8v = MakeSymbol("@v")
var symA = MakeSymbol("A")
var symAction = MakeSymbol("Action")
var symAssumption = MakeSymbol("Assumption")
var symAssumptions = MakeSymbol("Assumptions")
var symB = MakeSymbol("B")
var symC = MakeSymbol("C")
var symFinish = MakeSymbol("Finish")
var symFreeze = MakeSymbol("Freeze")
var symGoTo = MakeSymbol("GoTo")
var symHd = MakeSymbol("Hd")
var symHypotheses = MakeSymbol("Hypotheses")
var symK = MakeSymbol("K")
var symKey = MakeSymbol("Key")
var symL = MakeSymbol("L")
var symMessage = MakeSymbol("Message")
var symNewAssumptions = MakeSymbol("NewAssumptions")
var symNewV = MakeSymbol("NewV")
var symP = MakeSymbol("P")
var symParse = MakeSymbol("Parse")
var symRecord = MakeSymbol("Record")
var symRemainder = MakeSymbol("Remainder")
var symResult = MakeSymbol("Result")
var symS = MakeSymbol("S")
var symSelect = MakeSymbol("Select")
var symStart = MakeSymbol("Start")
var symTime = MakeSymbol("Time")
var symTl = MakeSymbol("Tl")
var symTm = MakeSymbol("Tm")
var symV = MakeSymbol("V")
var symW = MakeSymbol("W")
var symX = MakeSymbol("X")
var symY = MakeSymbol("Y")
var symZ = MakeSymbol("Z")
var sym__ = MakeSymbol("_")
var symabort = MakeSymbol("abort")
var symabsolute = MakeSymbol("absolute")
var symabsvector = MakeSymbol("absvector")
var symabsvector_2 = MakeSymbol("absvector?")
var symaddress_1_6 = MakeSymbol("address->")
var symadjoin = MakeSymbol("adjoin")
var symand = MakeSymbol("and")
var symappend = MakeSymbol("append")
var symarity = MakeSymbol("arity")
var symasserta = MakeSymbol("asserta")
var symassertz = MakeSymbol("assertz")
var symassoc = MakeSymbol("assoc")
var symatom_2 = MakeSymbol("atom?")
var symbar_b = MakeSymbol("bar!")
var symbind = MakeSymbol("bind")
var symboolean = MakeSymbol("boolean")
var symboolean_2 = MakeSymbol("boolean?")
var symbootstrap = MakeSymbol("bootstrap")
var symbound_2 = MakeSymbol("bound?")
var symcall = MakeSymbol("call")
var symcases = MakeSymbol("cases")
var symcd = MakeSymbol("cd")
var symchar_1stream_2 = MakeSymbol("char-stream?")
var symclose = MakeSymbol("close")
var symcn = MakeSymbol("cn")
var symcompile = MakeSymbol("compile")
var symconcat = MakeSymbol("concat")
var symcond = MakeSymbol("cond")
var symcons = MakeSymbol("cons")
var symcons_2 = MakeSymbol("cons?")
var symctxt = MakeSymbol("ctxt")
var symdatatype = MakeSymbol("datatype")
var symdatatypes = MakeSymbol("datatypes")
var symdeclare = MakeSymbol("declare")
var symdefcc = MakeSymbol("defcc")
var symdefine = MakeSymbol("define")
var symdefmacro = MakeSymbol("defmacro")
var symdefprolog = MakeSymbol("defprolog")
var symdefun = MakeSymbol("defun")
var symdestroy = MakeSymbol("destroy")
var symdifference = MakeSymbol("difference")
var symdo = MakeSymbol("do")
var symelement_2 = MakeSymbol("element?")
var symempty_2 = MakeSymbol("empty?")
var symenable_1type_1theory = MakeSymbol("enable-type-theory")
var symerror = MakeSymbol("error")
var symerror_1to_1string = MakeSymbol("error-to-string")
var symeval = MakeSymbol("eval")
var symeval_1kl = MakeSymbol("eval-kl")
var symexception = MakeSymbol("exception")
var symexplode = MakeSymbol("explode")
var symexternal = MakeSymbol("external")
var symfactorise = MakeSymbol("factorise")
var symfactorise_2 = MakeSymbol("factorise?")
var symfail = MakeSymbol("fail")
var symfail_1if = MakeSymbol("fail-if")
var symfile = MakeSymbol("file")
var symfindall = MakeSymbol("findall")
var symfix = MakeSymbol("fix")
var symfn = MakeSymbol("fn")
var symforeign = MakeSymbol("foreign")
var symfork = MakeSymbol("fork")
var symfreeze = MakeSymbol("freeze")
var symfresh = MakeSymbol("fresh")
var symfst = MakeSymbol("fst")
var symfunction = MakeSymbol("function")
var symg = MakeSymbol("g")
var symgensym = MakeSymbol("gensym")
var symget = MakeSymbol("get")
var symget_1time = MakeSymbol("get-time")
var symhash = MakeSymbol("hash")
var symhd = MakeSymbol("hd")
var symhdstr = MakeSymbol("hdstr")
var symhdv = MakeSymbol("hdv")
var symhead = MakeSymbol("head")
var symhush = MakeSymbol("hush")
var symhush_2 = MakeSymbol("hush?")
var symif = MakeSymbol("if")
var symimplementation = MakeSymbol("implementation")
var symin = MakeSymbol("in")
var symin_1package = MakeSymbol("in-package")
var syminclude = MakeSymbol("include")
var syminclude_1all_1but = MakeSymbol("include-all-but")
var symincluded = MakeSymbol("included")
var syminferences = MakeSymbol("inferences")
var syminline = MakeSymbol("inline")
var syminput = MakeSymbol("input")
var syminput_7 = MakeSymbol("input+")
var syminteger_2 = MakeSymbol("integer?")
var symintern = MakeSymbol("intern")
var syminternal = MakeSymbol("internal")
var symintersection = MakeSymbol("intersection")
var symis = MakeSymbol("is")
var symis_b = MakeSymbol("is!")
var symit = MakeSymbol("it")
var symlambda = MakeSymbol("lambda")
var symlanguage = MakeSymbol("language")
var symlaunch_1repl = MakeSymbol("launch-repl")
var symlazy = MakeSymbol("lazy")
var symlength = MakeSymbol("length")
var symlet = MakeSymbol("let")
var symlimit = MakeSymbol("limit")
var symlineread = MakeSymbol("lineread")
var symlist = MakeSymbol("list")
var symload = MakeSymbol("load")
var symloaded = MakeSymbol("loaded")
var symmacroexpand = MakeSymbol("macroexpand")
var symmake_1string = MakeSymbol("make-string")
var symmap = MakeSymbol("map")
var symmapcan = MakeSymbol("mapcan")
var symmaxinferences = MakeSymbol("maxinferences")
var symmode = MakeSymbol("mode")
var symn_1_6string = MakeSymbol("n->string")
var symnl = MakeSymbol("nl")
var symnot = MakeSymbol("not")
var symnth = MakeSymbol("nth")
var symnull = MakeSymbol("null")
var symnumber = MakeSymbol("number")
var symnumber_2 = MakeSymbol("number?")
var symoccurrences = MakeSymbol("occurrences")
var symoccurs_1check = MakeSymbol("occurs-check")
var symoccurs_2 = MakeSymbol("occurs?")
var symopen = MakeSymbol("open")
var symoptimise = MakeSymbol("optimise")
var symoptimise_2 = MakeSymbol("optimise?")
var symor = MakeSymbol("or")
var symos = MakeSymbol("os")
var symout = MakeSymbol("out")
var symoutput = MakeSymbol("output")
var sympackage = MakeSymbol("package")
var sympackage_2 = MakeSymbol("package?")
var symport = MakeSymbol("port")
var symporters = MakeSymbol("porters")
var sympos = MakeSymbol("pos")
var sympr = MakeSymbol("pr")
var sympreclude = MakeSymbol("preclude")
var sympreclude_1all_1but = MakeSymbol("preclude-all-but")
var symprint = MakeSymbol("print")
var symprofile = MakeSymbol("profile")
var symprofile_1results = MakeSymbol("profile-results")
var symprolog_1memory = MakeSymbol("prolog-memory")
var symprolog_2 = MakeSymbol("prolog?")
var symprotect = MakeSymbol("protect")
var symps = MakeSymbol("ps")
var symput = MakeSymbol("put")
var symread = MakeSymbol("read")
var symread_1byte = MakeSymbol("read-byte")
var symread_1file = MakeSymbol("read-file")
var symread_1file_1as_1bytelist = MakeSymbol("read-file-as-bytelist")
var symread_1file_1as_1string = MakeSymbol("read-file-as-string")
var symread_1from_1string = MakeSymbol("read-from-string")
var symread_1from_1string_1unprocessed = MakeSymbol("read-from-string-unprocessed")
var symread_1unit_1string = MakeSymbol("read-unit-string")
var symreceive = MakeSymbol("receive")
var symrelease = MakeSymbol("release")
var symremove = MakeSymbol("remove")
var symretract = MakeSymbol("retract")
var symreturn = MakeSymbol("return")
var symreverse = MakeSymbol("reverse")
var symrun = MakeSymbol("run")
var symsave = MakeSymbol("save")
var symset = MakeSymbol("set")
var symshen = MakeSymbol("shen")
var symshen_4_dalldatatypes_d = MakeSymbol("shen.*alldatatypes*")
var symshen_4_dcall_d = MakeSymbol("shen.*call*")
var symshen_4_dcustom_1pattern_1compiler_d = MakeSymbol("shen.*custom-pattern-compiler*")
var symshen_4_dcustom_1pattern_1reducer_d = MakeSymbol("shen.*custom-pattern-reducer*")
var symshen_4_ddatatypes_d = MakeSymbol("shen.*datatypes*")
var symshen_4_ddemodulation_1function_d = MakeSymbol("shen.*demodulation-function*")
var symshen_4_dempty_1absvector_d = MakeSymbol("shen.*empty-absvector*")
var symshen_4_dextraspecial_d = MakeSymbol("shen.*extraspecial*")
var symshen_4_dfactorise_2_d = MakeSymbol("shen.*factorise?*")
var symshen_4_dgensym_d = MakeSymbol("shen.*gensym*")
var symshen_4_dhistory_d = MakeSymbol("shen.*history*")
var symshen_4_dinfs_d = MakeSymbol("shen.*infs*")
var symshen_4_dit_d = MakeSymbol("shen.*it*")
var symshen_4_dloading_2_d = MakeSymbol("shen.*loading?*")
var symshen_4_dmaxinferences_d = MakeSymbol("shen.*maxinferences*")
var symshen_4_dnames_d = MakeSymbol("shen.*names*")
var symshen_4_doccurs_d = MakeSymbol("shen.*occurs*")
var symshen_4_doptimise_d = MakeSymbol("shen.*optimise*")
var symshen_4_dpackage_d = MakeSymbol("shen.*package*")
var symshen_4_dprofiled_d = MakeSymbol("shen.*profiled*")
var symshen_4_dprolog_1memory_d = MakeSymbol("shen.*prolog-memory*")
var symshen_4_dresidue_d = MakeSymbol("shen.*residue*")
var symshen_4_dshen_1type_1theory_1enabled_2_d = MakeSymbol("shen.*shen-type-theory-enabled?*")
var symshen_4_dsigf_d = MakeSymbol("shen.*sigf*")
var symshen_4_dsize_1prolog_1vector_d = MakeSymbol("shen.*size-prolog-vector*")
var symshen_4_dsource_1form_1handlers_d = MakeSymbol("shen.*source-form-handlers*")
var symshen_4_dspecial_d = MakeSymbol("shen.*special*")
var symshen_4_dspy_d = MakeSymbol("shen.*spy*")
var symshen_4_dstep_d = MakeSymbol("shen.*step*")
var symshen_4_dsynonyms_d = MakeSymbol("shen.*synonyms*")
var symshen_4_dtc_d = MakeSymbol("shen.*tc*")
var symshen_4_dtracking_d = MakeSymbol("shen.*tracking*")
var symshen_4_duserdefs_d = MakeSymbol("shen.*userdefs*")
var symshen_4_7m = MakeSymbol("shen.+m")
var symshen_4_7string_2 = MakeSymbol("shen.+string?")
var symshen_4_7vector_2 = MakeSymbol("shen.+vector?")
var symshen_4_1m = MakeSymbol("shen.-m")
var symshen_4_1null_1 = MakeSymbol("shen.-null-")
var symshen_4_5_1dict = MakeSymbol("shen.<-dict")
var symshen_4_5_1dict_1bucket = MakeSymbol("shen.<-dict-bucket")
var symshen_4_5_1out = MakeSymbol("shen.<-out")
var symshen_4_5alpha_6 = MakeSymbol("shen.<alpha>")
var symshen_4_5alphanum_6 = MakeSymbol("shen.<alphanum>")
var symshen_4_5alphanums_6 = MakeSymbol("shen.<alphanums>")
var symshen_4_5ass_6 = MakeSymbol("shen.<ass>")
var symshen_4_5atom_6 = MakeSymbol("shen.<atom>")
var symshen_4_5backslash_6 = MakeSymbol("shen.<backslash>")
var symshen_4_5bar_6 = MakeSymbol("shen.<bar>")
var symshen_4_5body_6 = MakeSymbol("shen.<body>")
var symshen_4_5bterm_6 = MakeSymbol("shen.<bterm>")
var symshen_4_5bterms_6 = MakeSymbol("shen.<bterms>")
var symshen_4_5c_1rule_6 = MakeSymbol("shen.<c-rule>")
var symshen_4_5c_1rules_6 = MakeSymbol("shen.<c-rules>")
var symshen_4_5clause_6 = MakeSymbol("shen.<clause>")
var symshen_4_5clauses_6 = MakeSymbol("shen.<clauses>")
var symshen_4_5colon_1equal_6 = MakeSymbol("shen.<colon-equal>")
var symshen_4_5colon_6 = MakeSymbol("shen.<colon>")
var symshen_4_5comma_6 = MakeSymbol("shen.<comma>")
var symshen_4_5comment_6 = MakeSymbol("shen.<comment>")
var symshen_4_5conc_6 = MakeSymbol("shen.<conc>")
var symshen_4_5constructor_6 = MakeSymbol("shen.<constructor>")
var symshen_4_5control_6 = MakeSymbol("shen.<control>")
var symshen_4_5datatype_1rule_6 = MakeSymbol("shen.<datatype-rule>")
var symshen_4_5datatype_1rules_6 = MakeSymbol("shen.<datatype-rules>")
var symshen_4_5datatype_6 = MakeSymbol("shen.<datatype>")
var symshen_4_5dbl_6 = MakeSymbol("shen.<dbl>")
var symshen_4_5dbq_6 = MakeSymbol("shen.<dbq>")
var symshen_4_5define_6 = MakeSymbol("shen.<define>")
var symshen_4_5defprolog_6 = MakeSymbol("shen.<defprolog>")
var symshen_4_5digit_6 = MakeSymbol("shen.<digit>")
var symshen_4_5digits_6 = MakeSymbol("shen.<digits>")
var symshen_4_5double_6 = MakeSymbol("shen.<double>")
var symshen_4_5e_1number_6 = MakeSymbol("shen.<e-number>")
var symshen_4_5equal_6 = MakeSymbol("shen.<equal>")
var symshen_4_5expr_6 = MakeSymbol("shen.<expr>")
var symshen_4_5float_6 = MakeSymbol("shen.<float>")
var symshen_4_5formula_6 = MakeSymbol("shen.<formula>")
var symshen_4_5formulae_6 = MakeSymbol("shen.<formulae>")
var symshen_4_5fraction_6 = MakeSymbol("shen.<fraction>")
var symshen_4_5hash_6 = MakeSymbol("shen.<hash>")
var symshen_4_5head_6 = MakeSymbol("shen.<head>")
var symshen_4_5hterm1_6 = MakeSymbol("shen.<hterm1>")
var symshen_4_5hterm2_6 = MakeSymbol("shen.<hterm2>")
var symshen_4_5hterm_6 = MakeSymbol("shen.<hterm>")
var symshen_4_5integer_6 = MakeSymbol("shen.<integer>")
var symshen_4_5iscolon_6 = MakeSymbol("shen.<iscolon>")
var symshen_4_5iscomma_6 = MakeSymbol("shen.<iscomma>")
var symshen_4_5lcurly_6 = MakeSymbol("shen.<lcurly>")
var symshen_4_5literal_6 = MakeSymbol("shen.<literal>")
var symshen_4_5log10_6 = MakeSymbol("shen.<log10>")
var symshen_4_5longnatter_6 = MakeSymbol("shen.<longnatter>")
var symshen_4_5lowC_6 = MakeSymbol("shen.<lowC>")
var symshen_4_5lowE_6 = MakeSymbol("shen.<lowE>")
var symshen_4_5lrb_6 = MakeSymbol("shen.<lrb>")
var symshen_4_5lsb_6 = MakeSymbol("shen.<lsb>")
var symshen_4_5minus_6 = MakeSymbol("shen.<minus>")
var symshen_4_5multiline_6 = MakeSymbol("shen.<multiline>")
var symshen_4_5name_6 = MakeSymbol("shen.<name>")
var symshen_4_5non_1terminal_1name_6 = MakeSymbol("shen.<non-terminal-name>")
var symshen_4_5non_1terminal_2_6 = MakeSymbol("shen.<non-terminal?>")
var symshen_4_5notdbq_6 = MakeSymbol("shen.<notdbq>")
var symshen_4_5number_6 = MakeSymbol("shen.<number>")
var symshen_4_5numeral_6 = MakeSymbol("shen.<numeral>")
var symshen_4_5packagechar_6 = MakeSymbol("shen.<packagechar>")
var symshen_4_5packagename_6 = MakeSymbol("shen.<packagename>")
var symshen_4_5packagenames_6 = MakeSymbol("shen.<packagenames>")
var symshen_4_5pattern1_6 = MakeSymbol("shen.<pattern1>")
var symshen_4_5pattern2_6 = MakeSymbol("shen.<pattern2>")
var symshen_4_5pattern_6 = MakeSymbol("shen.<pattern>")
var symshen_4_5patterns_6 = MakeSymbol("shen.<patterns>")
var symshen_4_5plus_6 = MakeSymbol("shen.<plus>")
var symshen_4_5prem_6 = MakeSymbol("shen.<prem>")
var symshen_4_5prems_6 = MakeSymbol("shen.<prems>")
var symshen_4_5rcurly_6 = MakeSymbol("shen.<rcurly>")
var symshen_4_5return_6 = MakeSymbol("shen.<return>")
var symshen_4_5returns_6 = MakeSymbol("shen.<returns>")
var symshen_4_5rrb_6 = MakeSymbol("shen.<rrb>")
var symshen_4_5rsb_6 = MakeSymbol("shen.<rsb>")
var symshen_4_5rule_d_6 = MakeSymbol("shen.<rule*>")
var symshen_4_5rule_6 = MakeSymbol("shen.<rule>")
var symshen_4_5rules_d_6 = MakeSymbol("shen.<rules*>")
var symshen_4_5rules_6 = MakeSymbol("shen.<rules>")
var symshen_4_5s_1exprs1_6 = MakeSymbol("shen.<s-exprs1>")
var symshen_4_5s_1exprs2_6 = MakeSymbol("shen.<s-exprs2>")
var symshen_4_5s_1exprs_6 = MakeSymbol("shen.<s-exprs>")
var symshen_4_5sc_6 = MakeSymbol("shen.<sc>")
var symshen_4_5semantics_6 = MakeSymbol("shen.<semantics>")
var symshen_4_5semicolon_6 = MakeSymbol("shen.<semicolon>")
var symshen_4_5shortnatter_6 = MakeSymbol("shen.<shortnatter>")
var symshen_4_5shortnatters_6 = MakeSymbol("shen.<shortnatters>")
var symshen_4_5side_6 = MakeSymbol("shen.<side>")
var symshen_4_5sides_6 = MakeSymbol("shen.<sides>")
var symshen_4_5sig_drules_6 = MakeSymbol("shen.<sig*rules>")
var symshen_4_5signature_6 = MakeSymbol("shen.<signature>")
var symshen_4_5simple_1pattern_6 = MakeSymbol("shen.<simple-pattern>")
var symshen_4_5single_6 = MakeSymbol("shen.<single>")
var symshen_4_5singleline_6 = MakeSymbol("shen.<singleline>")
var symshen_4_5sng_6 = MakeSymbol("shen.<sng>")
var symshen_4_5stop_6 = MakeSymbol("shen.<stop>")
var symshen_4_5str_6 = MakeSymbol("shen.<str>")
var symshen_4_5strc_6 = MakeSymbol("shen.<strc>")
var symshen_4_5strcontents_6 = MakeSymbol("shen.<strcontents>")
var symshen_4_5sym_6 = MakeSymbol("shen.<sym>")
var symshen_4_5syntax_1item_6 = MakeSymbol("shen.<syntax-item>")
var symshen_4_5syntax_6 = MakeSymbol("shen.<syntax>")
var symshen_4_5times_6 = MakeSymbol("shen.<times>")
var symshen_4_5type_6 = MakeSymbol("shen.<type>")
var symshen_4_5whitespace_6 = MakeSymbol("shen.<whitespace>")
var symshen_4_5whitespaces_6 = MakeSymbol("shen.<whitespaces>")
var symshen_4_5wildcard_6 = MakeSymbol("shen.<wildcard>")
var symshen_4_5yacc_6 = MakeSymbol("shen.<yacc>")
var symshen_4_5yaccsig_6 = MakeSymbol("shen.<yaccsig>")
var symshen_4_8c = MakeSymbol("shen.@c")
var symshen_4_8ch = MakeSymbol("shen.@ch")
var symshen_4_8v_1help = MakeSymbol("shen.@v-help")
var symshen_4a = MakeSymbol("shen.a")
var symshen_4abs = MakeSymbol("shen.abs")
var symshen_4add_1sexpr = MakeSymbol("shen.add-sexpr")
var symshen_4alpha_1convert = MakeSymbol("shen.alpha-convert")
var symshen_4alpha_2 = MakeSymbol("shen.alpha?")
var symshen_4alphanums_2 = MakeSymbol("shen.alphanums?")
var symshen_4analyse_1symbol_2 = MakeSymbol("shen.analyse-symbol?")
var symshen_4analyse_1variable_2 = MakeSymbol("shen.analyse-variable?")
var symshen_4app = MakeSymbol("shen.app")
var symshen_4application_2 = MakeSymbol("shen.application?")
var symshen_4arg_1_6str = MakeSymbol("shen.arg->str")
var symshen_4arity_1chk = MakeSymbol("shen.arity-chk")
var symshen_4assert_d = MakeSymbol("shen.assert*")
var symshen_4assoc_1_6 = MakeSymbol("shen.assoc->")
var symshen_4assoc_1rm = MakeSymbol("shen.assoc-rm")
var symshen_4assoc_1set = MakeSymbol("shen.assoc-set")
var symshen_4assumetypes = MakeSymbol("shen.assumetypes")
var symshen_4atom_1_6str = MakeSymbol("shen.atom->str")
var symshen_4atom_1case_1minus = MakeSymbol("shen.atom-case-minus")
var symshen_4atom_1case_1plus = MakeSymbol("shen.atom-case-plus")
var symshen_4autocomplete = MakeSymbol("shen.autocomplete")
var symshen_4bad_1pivot_2 = MakeSymbol("shen.bad-pivot?")
var symshen_4beta = MakeSymbol("shen.beta")
var symshen_4bind_b = MakeSymbol("shen.bind!")
var symshen_4bindv = MakeSymbol("shen.bindv")
var symshen_4bottom = MakeSymbol("shen.bottom")
var symshen_4bucket_1fold = MakeSymbol("shen.bucket-fold")
var symshen_4build_1lambda_1table = MakeSymbol("shen.build-lambda-table")
var symshen_4by_1hypothesis = MakeSymbol("shen.by-hypothesis")
var symshen_4by_1hypothesis_1h = MakeSymbol("shen.by-hypothesis-h")
var symshen_4byte_1_6digit = MakeSymbol("shen.byte->digit")
var symshen_4bytes_1_6string = MakeSymbol("shen.bytes->string")
var symshen_4c_1rule_1_6shen = MakeSymbol("shen.c-rule->shen")
var symshen_4c_1rules_1_6shen = MakeSymbol("shen.c-rules->shen")
var symshen_4call_1dynamic = MakeSymbol("shen.call-dynamic")
var symshen_4call_1prolog = MakeSymbol("shen.call-prolog")
var symshen_4callrec = MakeSymbol("shen.callrec")
var symshen_4ccons_2 = MakeSymbol("shen.ccons?")
var symshen_4check_1byte = MakeSymbol("shen.check-byte")
var symshen_4check_1eval_1and_1print = MakeSymbol("shen.check-eval-and-print")
var symshen_4choicepoint = MakeSymbol("shen.choicepoint")
var symshen_4choicepoint_b = MakeSymbol("shen.choicepoint!")
var symshen_4choicepoint_2 = MakeSymbol("shen.choicepoint?")
var symshen_4coll_1formulae = MakeSymbol("shen.coll-formulae")
var symshen_4colon_1equal_2 = MakeSymbol("shen.colon-equal?")
var symshen_4comb = MakeSymbol("shen.comb")
var symshen_4combine_1c_1code = MakeSymbol("shen.combine-c-code")
var symshen_4compile_1body = MakeSymbol("shen.compile-body")
var symshen_4compile_1head = MakeSymbol("shen.compile-head")
var symshen_4compile_1pattern_1fragment = MakeSymbol("shen.compile-pattern-fragment")
var symshen_4compile_1prolog = MakeSymbol("shen.compile-prolog")
var symshen_4compile_1synonyms = MakeSymbol("shen.compile-synonyms")
var symshen_4compile_1to_1kl = MakeSymbol("shen.compile-to-kl")
var symshen_4compound_1pattern = MakeSymbol("shen.compound-pattern")
var symshen_4compound_1pattern_1h = MakeSymbol("shen.compound-pattern-h")
var symshen_4compute_1E = MakeSymbol("shen.compute-E")
var symshen_4compute_1fraction = MakeSymbol("shen.compute-fraction")
var symshen_4compute_1fraction_1h = MakeSymbol("shen.compute-fraction-h")
var symshen_4compute_1integer = MakeSymbol("shen.compute-integer")
var symshen_4compute_1integer_1h = MakeSymbol("shen.compute-integer-h")
var symshen_4cond_1form = MakeSymbol("shen.cond-form")
var symshen_4cons_1case_1minus = MakeSymbol("shen.cons-case-minus")
var symshen_4cons_1case_1plus = MakeSymbol("shen.cons-case-plus")
var symshen_4cons_1form = MakeSymbol("shen.cons-form")
var symshen_4cons_1form_1respect_1modes = MakeSymbol("shen.cons-form-respect-modes")
var symshen_4conscode = MakeSymbol("shen.conscode")
var symshen_4construct_1context = MakeSymbol("shen.construct-context")
var symshen_4constructor_1error = MakeSymbol("shen.constructor-error")
var symshen_4constructor_2 = MakeSymbol("shen.constructor?")
var symshen_4consume = MakeSymbol("shen.consume")
var symshen_4consume_1clause = MakeSymbol("shen.consume-clause")
var symshen_4continue = MakeSymbol("shen.continue")
var symshen_4copyfromvector = MakeSymbol("shen.copyfromvector")
var symshen_4correct = MakeSymbol("shen.correct")
var symshen_4create_1skeleton = MakeSymbol("shen.create-skeleton")
var symshen_4credits = MakeSymbol("shen.credits")
var symshen_4curry = MakeSymbol("shen.curry")
var symshen_4curry_1type = MakeSymbol("shen.curry-type")
var symshen_4custom_1pattern = MakeSymbol("shen.custom-pattern")
var symshen_4custom_1pattern_1body = MakeSymbol("shen.custom-pattern-body")
var symshen_4custom_1pattern_1compiler = MakeSymbol("shen.custom-pattern-compiler")
var symshen_4custom_1pattern_1reducer = MakeSymbol("shen.custom-pattern-reducer")
var symshen_4custom_1pattern_1triple_1stack = MakeSymbol("shen.custom-pattern-triple-stack")
var symshen_4custom_1pattern_2 = MakeSymbol("shen.custom-pattern?")
var symshen_4cut = MakeSymbol("shen.cut")
var symshen_4dbl_1h_2 = MakeSymbol("shen.dbl-h?")
var symshen_4dbl_2 = MakeSymbol("shen.dbl?")
var symshen_4decons = MakeSymbol("shen.decons")
var symshen_4decrement_1ticket = MakeSymbol("shen.decrement-ticket")
var symshen_4demod = MakeSymbol("shen.demod")
var symshen_4demode = MakeSymbol("shen.demode")
var symshen_4demodulate = MakeSymbol("shen.demodulate")
var symshen_4deref = MakeSymbol("shen.deref")
var symshen_4deref_1calls = MakeSymbol("shen.deref-calls")
var symshen_4deref_1forked_1literals = MakeSymbol("shen.deref-forked-literals")
var symshen_4deref_1terms = MakeSymbol("shen.deref-terms")
var symshen_4dict = MakeSymbol("shen.dict")
var symshen_4dict_1_6 = MakeSymbol("shen.dict->")
var symshen_4dict_1bucket_1_6 = MakeSymbol("shen.dict-bucket->")
var symshen_4dict_1capacity = MakeSymbol("shen.dict-capacity")
var symshen_4dict_1count = MakeSymbol("shen.dict-count")
var symshen_4dict_1count_1_6 = MakeSymbol("shen.dict-count->")
var symshen_4dict_1fold = MakeSymbol("shen.dict-fold")
var symshen_4dict_1fold_1h = MakeSymbol("shen.dict-fold-h")
var symshen_4dict_1keys = MakeSymbol("shen.dict-keys")
var symshen_4dict_1rm = MakeSymbol("shen.dict-rm")
var symshen_4dict_1update_1count = MakeSymbol("shen.dict-update-count")
var symshen_4dict_1values = MakeSymbol("shen.dict-values")
var symshen_4dict_2 = MakeSymbol("shen.dict?")
var symshen_4dictionary = MakeSymbol("shen.dictionary")
var symshen_4digit_2 = MakeSymbol("shen.digit?")
var symshen_4dynamic = MakeSymbol("shen.dynamic")
var symshen_4dynamic_1default = MakeSymbol("shen.dynamic-default")
var symshen_4empty_1absvector_2 = MakeSymbol("shen.empty-absvector?")
var symshen_4eos = MakeSymbol("shen.eos")
var symshen_4eval_1and_1print = MakeSymbol("shen.eval-and-print")
var symshen_4evaluate_1lineread = MakeSymbol("shen.evaluate-lineread")
var symshen_4execute_1store_1arity = MakeSymbol("shen.execute-store-arity")
var symshen_4explode_1h = MakeSymbol("shen.explode-h")
var symshen_4expt = MakeSymbol("shen.expt")
var symshen_4external_1symbols = MakeSymbol("shen.external-symbols")
var symshen_4extract_1free_1vars = MakeSymbol("shen.extract-free-vars")
var symshen_4extract_1vars = MakeSymbol("shen.extract-vars")
var symshen_4extraspecial_2 = MakeSymbol("shen.extraspecial?")
var symshen_4f = MakeSymbol("shen.f")
var symshen_4f_1error = MakeSymbol("shen.f-error")
var symshen_4factor = MakeSymbol("shen.factor")
var symshen_4factor_1cn = MakeSymbol("shen.factor-cn")
var symshen_4factor_1recognisors = MakeSymbol("shen.factor-recognisors")
var symshen_4factor_1selectors = MakeSymbol("shen.factor-selectors")
var symshen_4factor_1selectors_1h = MakeSymbol("shen.factor-selectors-h")
var symshen_4factorise_1code = MakeSymbol("shen.factorise-code")
var symshen_4fail_b = MakeSymbol("shen.fail!")
var symshen_4fbound_2 = MakeSymbol("shen.fbound?")
var symshen_4fillvector = MakeSymbol("shen.fillvector")
var symshen_4find_1arities = MakeSymbol("shen.find-arities")
var symshen_4find_1arity = MakeSymbol("shen.find-arity")
var symshen_4find_1free_1vars = MakeSymbol("shen.find-free-vars")
var symshen_4find_1types = MakeSymbol("shen.find-types")
var symshen_4findall_1h = MakeSymbol("shen.findall-h")
var symshen_4fits_2 = MakeSymbol("shen.fits?")
var symshen_4fix_1help = MakeSymbol("shen.fix-help")
var symshen_4fn_1call = MakeSymbol("shen.fn-call")
var symshen_4fn_1call_2 = MakeSymbol("shen.fn-call?")
var symshen_4fn_1print = MakeSymbol("shen.fn-print")
var symshen_4for_1each = MakeSymbol("shen.for-each")
var symshen_4foreign_2 = MakeSymbol("shen.foreign?")
var symshen_4free_1var_1chk = MakeSymbol("shen.free-var-chk")
var symshen_4free_1variable_1error_1message = MakeSymbol("shen.free-variable-error-message")
var symshen_4free_1variable_2 = MakeSymbol("shen.free-variable?")
var symshen_4freeze_1literals = MakeSymbol("shen.freeze-literals")
var symshen_4freshen = MakeSymbol("shen.freshen")
var symshen_4freshen_1rule = MakeSymbol("shen.freshen-rule")
var symshen_4freshen_1sig = MakeSymbol("shen.freshen-sig")
var symshen_4freshen_1type = MakeSymbol("shen.freshen-type")
var symshen_4freshterm = MakeSymbol("shen.freshterm")
var symshen_4freshterm_2 = MakeSymbol("shen.freshterm?")
var symshen_4freshterms = MakeSymbol("shen.freshterms")
var symshen_4function_1calls = MakeSymbol("shen.function-calls")
var symshen_4funexstring = MakeSymbol("shen.funexstring")
var symshen_4gc = MakeSymbol("shen.gc")
var symshen_4get_1profile = MakeSymbol("shen.get-profile")
var symshen_4goto = MakeSymbol("shen.goto")
var symshen_4goto_1h = MakeSymbol("shen.goto-h")
var symshen_4hascut_2 = MakeSymbol("shen.hascut?")
var symshen_4hashkey = MakeSymbol("shen.hashkey")
var symshen_4hds_a_2 = MakeSymbol("shen.hds=?")
var symshen_4horn_1clause_1procedure = MakeSymbol("shen.horn-clause-procedure")
var symshen_4i_1failed_b = MakeSymbol("shen.i-failed!")
var symshen_4in_1_6 = MakeSymbol("shen.in->")
var symshen_4incinfs = MakeSymbol("shen.incinfs")
var symshen_4initialise = MakeSymbol("shen.initialise")
var symshen_4initialise_1arity_1table = MakeSymbol("shen.initialise-arity-table")
var symshen_4initialise_1environment = MakeSymbol("shen.initialise-environment")
var symshen_4initialise_1lambda_1forms = MakeSymbol("shen.initialise-lambda-forms")
var symshen_4initialise_1signedfuncs = MakeSymbol("shen.initialise-signedfuncs")
var symshen_4initialise__environment = MakeSymbol("shen.initialise_environment")
var symshen_4input_1h_7 = MakeSymbol("shen.input-h+")
var symshen_4input_1track = MakeSymbol("shen.input-track")
var symshen_4insert = MakeSymbol("shen.insert")
var symshen_4insert_1h = MakeSymbol("shen.insert-h")
var symshen_4insert_1info = MakeSymbol("shen.insert-info")
var symshen_4insert_1l = MakeSymbol("shen.insert-l")
var symshen_4insert_1prolog_1variables = MakeSymbol("shen.insert-prolog-variables")
var symshen_4insert_1tracking_1code = MakeSymbol("shen.insert-tracking-code")
var symshen_4integer_1test_2 = MakeSymbol("shen.integer-test?")
var symshen_4intern_1in_1package = MakeSymbol("shen.intern-in-package")
var symshen_4intern_1type = MakeSymbol("shen.intern-type")
var symshen_4internal_1symbols = MakeSymbol("shen.internal-symbols")
var symshen_4internal_1to_1P_2 = MakeSymbol("shen.internal-to-P?")
var symshen_4internal_1to_1shen_2 = MakeSymbol("shen.internal-to-shen?")
var symshen_4internal_2 = MakeSymbol("shen.internal?")
var symshen_4invoke = MakeSymbol("shen.invoke")
var symshen_4iter_1list = MakeSymbol("shen.iter-list")
var symshen_4iter_1vector = MakeSymbol("shen.iter-vector")
var symshen_4key_1in_1sequent_1calculus_2 = MakeSymbol("shen.key-in-sequent-calculus?")
var symshen_4kl_1body = MakeSymbol("shen.kl-body")
var symshen_4klfile = MakeSymbol("shen.klfile")
var symshen_4l_1rules = MakeSymbol("shen.l-rules")
var symshen_4lambda_1entry = MakeSymbol("shen.lambda-entry")
var symshen_4lambda_1form = MakeSymbol("shen.lambda-form")
var symshen_4lambda_1function = MakeSymbol("shen.lambda-function")
var symshen_4lambda_1of_1defun = MakeSymbol("shen.lambda-of-defun")
var symshen_4lazyderef = MakeSymbol("shen.lazyderef")
var symshen_4lch = MakeSymbol("shen.lch")
var symshen_4lchh = MakeSymbol("shen.lchh")
var symshen_4length_1h = MakeSymbol("shen.length-h")
var symshen_4line = MakeSymbol("shen.line")
var symshen_4linearise = MakeSymbol("shen.linearise")
var symshen_4linearise_1clause = MakeSymbol("shen.linearise-clause")
var symshen_4linearise_1h = MakeSymbol("shen.linearise-h")
var symshen_4list_1_6str = MakeSymbol("shen.list->str")
var symshen_4list_2 = MakeSymbol("shen.list?")
var symshen_4load_1help = MakeSymbol("shen.load-help")
var symshen_4loading_2 = MakeSymbol("shen.loading?")
var symshen_4lock = MakeSymbol("shen.lock")
var symshen_4locked_2 = MakeSymbol("shen.locked?")
var symshen_4lookupsig = MakeSymbol("shen.lookupsig")
var symshen_4loop = MakeSymbol("shen.loop")
var symshen_4lowercase_1symbol_2 = MakeSymbol("shen.lowercase-symbol?")
var symshen_4lowercase_2 = MakeSymbol("shen.lowercase?")
var symshen_4lr_1rule = MakeSymbol("shen.lr-rule")
var symshen_4lzy_a = MakeSymbol("shen.lzy=")
var symshen_4lzy_a_b = MakeSymbol("shen.lzy=!")
var symshen_4macro_1_8c = MakeSymbol("shen.macro-@c")
var symshen_4macro_1_8ch = MakeSymbol("shen.macro-@ch")
var symshen_4macroexpand_1h = MakeSymbol("shen.macroexpand-h")
var symshen_4macros = MakeSymbol("shen.macros")
var symshen_4magless = MakeSymbol("shen.magless")
var symshen_4make_1prolog_1variable = MakeSymbol("shen.make-prolog-variable")
var symshen_4make_1uppercase = MakeSymbol("shen.make-uppercase")
var symshen_4map_1h = MakeSymbol("shen.map-h")
var symshen_4maxinfexceeded_2 = MakeSymbol("shen.maxinfexceeded?")
var symshen_4maxseq = MakeSymbol("shen.maxseq")
var symshen_4member = MakeSymbol("shen.member")
var symshen_4member_1clause = MakeSymbol("shen.member-clause")
var symshen_4misc_2 = MakeSymbol("shen.misc?")
var symshen_4mkstr = MakeSymbol("shen.mkstr")
var symshen_4mkstr_1l = MakeSymbol("shen.mkstr-l")
var symshen_4mkstr_1r = MakeSymbol("shen.mkstr-r")
var symshen_4mod = MakeSymbol("shen.mod")
var symshen_4modh = MakeSymbol("shen.modh")
var symshen_4monomorphic_2 = MakeSymbol("shen.monomorphic?")
var symshen_4monotype = MakeSymbol("shen.monotype")
var symshen_4mu_1h = MakeSymbol("shen.mu-h")
var symshen_4multiples = MakeSymbol("shen.multiples")
var symshen_4my_1read_1byte = MakeSymbol("shen.my-read-byte")
var symshen_4myassume = MakeSymbol("shen.myassume")
var symshen_4newname = MakeSymbol("shen.newname")
var symshen_4newpv = MakeSymbol("shen.newpv")
var symshen_4nextticket = MakeSymbol("shen.nextticket")
var symshen_4non_1application_2 = MakeSymbol("shen.non-application?")
var symshen_4non_1terminal_2 = MakeSymbol("shen.non-terminal?")
var symshen_4non_1terminalcode = MakeSymbol("shen.non-terminalcode")
var symshen_4not_1dictionary = MakeSymbol("shen.not-dictionary")
var symshen_4not_1pvar = MakeSymbol("shen.not-pvar")
var symshen_4not_1tuple = MakeSymbol("shen.not-tuple")
var symshen_4nothing_1doing_2 = MakeSymbol("shen.nothing-doing?")
var symshen_4nvars = MakeSymbol("shen.nvars")
var symshen_4objectcode = MakeSymbol("shen.objectcode")
var symshen_4occurs_1check_2 = MakeSymbol("shen.occurs-check?")
var symshen_4ok = MakeSymbol("shen.ok")
var symshen_4op = MakeSymbol("shen.op")
var symshen_4op_1test = MakeSymbol("shen.op-test")
var symshen_4op1 = MakeSymbol("shen.op1")
var symshen_4op2 = MakeSymbol("shen.op2")
var symshen_4openlock = MakeSymbol("shen.openlock")
var symshen_4out_1of_1bounds = MakeSymbol("shen.out-of-bounds")
var symshen_4output_1track = MakeSymbol("shen.output-track")
var symshen_4overapplication_2 = MakeSymbol("shen.overapplication?")
var symshen_4overbind = MakeSymbol("shen.overbind")
var symshen_4p_1hyps = MakeSymbol("shen.p-hyps")
var symshen_4pac_1h = MakeSymbol("shen.pac-h")
var symshen_4package_1symbols = MakeSymbol("shen.package-symbols")
var symshen_4package_1user_1input = MakeSymbol("shen.package-user-input")
var symshen_4packaged_2 = MakeSymbol("shen.packaged?")
var symshen_4parameters = MakeSymbol("shen.parameters")
var symshen_4parse_1failure = MakeSymbol("shen.parse-failure")
var symshen_4parse_1failure_2 = MakeSymbol("shen.parse-failure?")
var symshen_4partial = MakeSymbol("shen.partial")
var symshen_4partial_1application_d_2 = MakeSymbol("shen.partial-application*?")
var symshen_4partial_1parse_1failure_2 = MakeSymbol("shen.partial-parse-failure?")
var symshen_4passive_1bind = MakeSymbol("shen.passive-bind")
var symshen_4passive_1variables = MakeSymbol("shen.passive-variables")
var symshen_4pause_1for_1user = MakeSymbol("shen.pause-for-user")
var symshen_4peek_1history = MakeSymbol("shen.peek-history")
var symshen_4pivot_1on = MakeSymbol("shen.pivot-on")
var symshen_4posint_2 = MakeSymbol("shen.posint?")
var symshen_4predicate = MakeSymbol("shen.predicate")
var symshen_4premises_1_6goals = MakeSymbol("shen.premises->goals")
var symshen_4pretty_1type = MakeSymbol("shen.pretty-type")
var symshen_4prhush = MakeSymbol("shen.prhush")
var symshen_4primitive = MakeSymbol("shen.primitive")
var symshen_4print_1freshterm = MakeSymbol("shen.print-freshterm")
var symshen_4print_1prolog_1vector = MakeSymbol("shen.print-prolog-vector")
var symshen_4print_1vector_2 = MakeSymbol("shen.print-vector?")
var symshen_4printF = MakeSymbol("shen.printF")
var symshen_4proc_1nl = MakeSymbol("shen.proc-nl")
var symshen_4process_1_8s = MakeSymbol("shen.process-@s")
var symshen_4process_1after_1type = MakeSymbol("shen.process-after-type")
var symshen_4process_1application = MakeSymbol("shen.process-application")
var symshen_4process_1applications = MakeSymbol("shen.process-applications")
var symshen_4process_1assoc = MakeSymbol("shen.process-assoc")
var symshen_4process_1cases = MakeSymbol("shen.process-cases")
var symshen_4process_1cond_1clauses = MakeSymbol("shen.process-cond-clauses")
var symshen_4process_1datatype = MakeSymbol("shen.process-datatype")
var symshen_4process_1def = MakeSymbol("shen.process-def")
var symshen_4process_1input_7 = MakeSymbol("shen.process-input+")
var symshen_4process_1lambda = MakeSymbol("shen.process-lambda")
var symshen_4process_1let = MakeSymbol("shen.process-let")
var symshen_4process_1read_1byte = MakeSymbol("shen.process-read-byte")
var symshen_4process_1sexprs = MakeSymbol("shen.process-sexprs")
var symshen_4process_1synonyms = MakeSymbol("shen.process-synonyms")
var symshen_4process_1time = MakeSymbol("shen.process-time")
var symshen_4process_1yacc_1semantics = MakeSymbol("shen.process-yacc-semantics")
var symshen_4processed = MakeSymbol("shen.processed")
var symshen_4prodbutzero = MakeSymbol("shen.prodbutzero")
var symshen_4profile_1func = MakeSymbol("shen.profile-func")
var symshen_4profile_1help = MakeSymbol("shen.profile-help")
var symshen_4profiled_2 = MakeSymbol("shen.profiled?")
var symshen_4prolog_1abstraction = MakeSymbol("shen.prolog-abstraction")
var symshen_4prolog_1arity_1check = MakeSymbol("shen.prolog-arity-check")
var symshen_4prolog_1fbody = MakeSymbol("shen.prolog-fbody")
var symshen_4prolog_1keyword_2 = MakeSymbol("shen.prolog-keyword?")
var symshen_4prolog_1parameters = MakeSymbol("shen.prolog-parameters")
var symshen_4prolog_1track = MakeSymbol("shen.prolog-track")
var symshen_4prolog_1vector = MakeSymbol("shen.prolog-vector")
var symshen_4prolog_1vector_1size = MakeSymbol("shen.prolog-vector-size")
var symshen_4prompt = MakeSymbol("shen.prompt")
var symshen_4prterm = MakeSymbol("shen.prterm")
var symshen_4prtl = MakeSymbol("shen.prtl")
var symshen_4pui_1h = MakeSymbol("shen.pui-h")
var symshen_4put_1profile = MakeSymbol("shen.put-profile")
var symshen_4pvar = MakeSymbol("shen.pvar")
var symshen_4pvar_2 = MakeSymbol("shen.pvar?")
var symshen_4r = MakeSymbol("shen.r")
var symshen_4raise_1syntax_1error = MakeSymbol("shen.raise-syntax-error")
var symshen_4rcons__form = MakeSymbol("shen.rcons_form")
var symshen_4rdecons = MakeSymbol("shen.rdecons")
var symshen_4read_1evaluate_1print = MakeSymbol("shen.read-evaluate-print")
var symshen_4read_1file_1as_1bytelist_1help = MakeSymbol("shen.read-file-as-bytelist-help")
var symshen_4read_1loop = MakeSymbol("shen.read-loop")
var symshen_4reader_1error = MakeSymbol("shen.reader-error")
var symshen_4reader_1error_1message = MakeSymbol("shen.reader-error-message")
var symshen_4received = MakeSymbol("shen.received")
var symshen_4record_1and_1evaluate = MakeSymbol("shen.record-and-evaluate")
var symshen_4record_1external = MakeSymbol("shen.record-external")
var symshen_4record_1internal = MakeSymbol("shen.record-internal")
var symshen_4record_1it = MakeSymbol("shen.record-it")
var symshen_4record_1kl = MakeSymbol("shen.record-kl")
var symshen_4record_1macro = MakeSymbol("shen.record-macro")
var symshen_4rectify_1test = MakeSymbol("shen.rectify-test")
var symshen_4rectify_1type = MakeSymbol("shen.rectify-type")
var symshen_4recursive_1string_1match = MakeSymbol("shen.recursive-string-match")
var symshen_4recursively_1factor_1selectors = MakeSymbol("shen.recursively-factor-selectors")
var symshen_4recursively_1print = MakeSymbol("shen.recursively-print")
var symshen_4register_1source_1form = MakeSymbol("shen.register-source-form")
var symshen_4remember_1datatype = MakeSymbol("shen.remember-datatype")
var symshen_4remove_1bystanders = MakeSymbol("shen.remove-bystanders")
var symshen_4remove_1datatypes = MakeSymbol("shen.remove-datatypes")
var symshen_4remove_1h = MakeSymbol("shen.remove-h")
var symshen_4remove_1indirection = MakeSymbol("shen.remove-indirection")
var symshen_4rep_1X = MakeSymbol("shen.rep-X")
var symshen_4repl = MakeSymbol("shen.repl")
var symshen_4restore_1P = MakeSymbol("shen.restore-P")
var symshen_4restore_1local = MakeSymbol("shen.restore-local")
var symshen_4retract_1clause = MakeSymbol("shen.retract-clause")
var symshen_4return_2 = MakeSymbol("shen.return?")
var symshen_4reverse_1help = MakeSymbol("shen.reverse-help")
var symshen_4rfas_1h = MakeSymbol("shen.rfas-h")
var symshen_4rule_1_6body = MakeSymbol("shen.rule->body")
var symshen_4rule_1_6clause = MakeSymbol("shen.rule->clause")
var symshen_4rule_1_6head = MakeSymbol("shen.rule->head")
var symshen_4rules_1_6prolog = MakeSymbol("shen.rules->prolog")
var symshen_4s = MakeSymbol("shen.s")
var symshen_4scan_1body = MakeSymbol("shen.scan-body")
var symshen_4search_1user_1datatypes = MakeSymbol("shen.search-user-datatypes")
var symshen_4search_1user_1datatypes_1h = MakeSymbol("shen.search-user-datatypes-h")
var symshen_4semicolon_2 = MakeSymbol("shen.semicolon?")
var symshen_4set_1lambda_1form_1entry = MakeSymbol("shen.set-lambda-form-entry")
var symshen_4shen_1_6kl = MakeSymbol("shen.shen->kl")
var symshen_4shen_1_6kl_1h = MakeSymbol("shen.shen->kl-h")
var symshen_4shen_1call_2 = MakeSymbol("shen.shen-call?")
var symshen_4shendef_1_6kldef = MakeSymbol("shen.shendef->kldef")
var symshen_4shendef_1_6kldef_1h = MakeSymbol("shen.shendef->kldef-h")
var symshen_4show = MakeSymbol("shen.show")
var symshen_4show_1assumptions = MakeSymbol("shen.show-assumptions")
var symshen_4show_1datatypes = MakeSymbol("shen.show-datatypes")
var symshen_4show_1p = MakeSymbol("shen.show-p")
var symshen_4side_1conditions_1_6goals = MakeSymbol("shen.side-conditions->goals")
var symshen_4sigf = MakeSymbol("shen.sigf")
var symshen_4signal_1def = MakeSymbol("shen.signal-def")
var symshen_4sigxrules = MakeSymbol("shen.sigxrules")
var symshen_4simple_1curry = MakeSymbol("shen.simple-curry")
var symshen_4skip = MakeSymbol("shen.skip")
var symshen_4sng_1h_2 = MakeSymbol("shen.sng-h?")
var symshen_4sng_2 = MakeSymbol("shen.sng?")
var symshen_4source = MakeSymbol("shen.source")
var symshen_4source_1form_1expansion = MakeSymbol("shen.source-form-expansion")
var symshen_4source_1form_1handler = MakeSymbol("shen.source-form-handler")
var symshen_4source_1form_1list_2 = MakeSymbol("shen.source-form-list?")
var symshen_4source_1macro_1walk = MakeSymbol("shen.source-macro-walk")
var symshen_4spaces = MakeSymbol("shen.spaces")
var symshen_4special_1case = MakeSymbol("shen.special-case")
var symshen_4special_2 = MakeSymbol("shen.special?")
var symshen_4specialise_1consume = MakeSymbol("shen.specialise-consume")
var symshen_4specialise_1member = MakeSymbol("shen.specialise-member")
var symshen_4store_1arity = MakeSymbol("shen.store-arity")
var symshen_4stpart = MakeSymbol("shen.stpart")
var symshen_4str_1_6bytes = MakeSymbol("shen.str->bytes")
var symshen_4str_1_6str = MakeSymbol("shen.str->str")
var symshen_4string_1_6byte = MakeSymbol("shen.string->byte")
var symshen_4string_1_6bytes = MakeSymbol("shen.string->bytes")
var symshen_4string_1match = MakeSymbol("shen.string-match")
var symshen_4string_1prefix_2 = MakeSymbol("shen.string-prefix?")
var symshen_4synonyms_1h = MakeSymbol("shen.synonyms-h")
var symshen_4syntax_1error_1message = MakeSymbol("shen.syntax-error-message")
var symshen_4syntax_1item_2 = MakeSymbol("shen.syntax-item?")
var symshen_4sysfunc_2 = MakeSymbol("shen.sysfunc?")
var symshen_4system_1S = MakeSymbol("shen.system-S")
var symshen_4system_1S_1h = MakeSymbol("shen.system-S-h")
var symshen_4t = MakeSymbol("shen.t")
var symshen_4t_d = MakeSymbol("shen.t*")
var symshen_4t_d_1correct = MakeSymbol("shen.t*-correct")
var symshen_4t_d_1integrity = MakeSymbol("shen.t*-integrity")
var symshen_4t_d_1rule = MakeSymbol("shen.t*-rule")
var symshen_4t_d_1rule_1h = MakeSymbol("shen.t*-rule-h")
var symshen_4t_d_1rules = MakeSymbol("shen.t*-rules")
var symshen_4tame = MakeSymbol("shen.tame")
var symshen_4terminalcode = MakeSymbol("shen.terminalcode")
var symshen_4terms = MakeSymbol("shen.terms")
var symshen_4terpri_1or_1read_1char = MakeSymbol("shen.terpri-or-read-char")
var symshen_4this_1symbol_1is_1unbound = MakeSymbol("shen.this-symbol-is-unbound")
var symshen_4ticket_1number = MakeSymbol("shen.ticket-number")
var symshen_4tlv_1help = MakeSymbol("shen.tlv-help")
var symshen_4top = MakeSymbol("shen.top")
var symshen_4toplevel_1display_1exception = MakeSymbol("shen.toplevel-display-exception")
var symshen_4toplevel_1forms = MakeSymbol("shen.toplevel-forms")
var symshen_4track_1function = MakeSymbol("shen.track-function")
var symshen_4tracked_2 = MakeSymbol("shen.tracked?")
var symshen_4trim_1it = MakeSymbol("shen.trim-it")
var symshen_4triple_1stack = MakeSymbol("shen.triple-stack")
var symshen_4try_1parse = MakeSymbol("shen.try-parse")
var symshen_4tuple = MakeSymbol("shen.tuple")
var symshen_4tuple_1up = MakeSymbol("shen.tuple-up")
var symshen_4type_1F = MakeSymbol("shen.type-F")
var symshen_4type_1error = MakeSymbol("shen.type-error")
var symshen_4type_1theory_1enabled_2 = MakeSymbol("shen.type-theory-enabled?")
var symshen_4typecheck = MakeSymbol("shen.typecheck")
var symshen_4typename = MakeSymbol("shen.typename")
var symshen_4typename_1h = MakeSymbol("shen.typename-h")
var symshen_4typetable = MakeSymbol("shen.typetable")
var symshen_4unassoc = MakeSymbol("shen.unassoc")
var symshen_4undefined_1f_2 = MakeSymbol("shen.undefined-f?")
var symshen_4unlock = MakeSymbol("shen.unlock")
var symshen_4unlocked_2 = MakeSymbol("shen.unlocked?")
var symshen_4unmark_1custom_1pattern = MakeSymbol("shen.unmark-custom-pattern")
var symshen_4unmark_1rule_1custom_1patterns = MakeSymbol("shen.unmark-rule-custom-patterns")
var symshen_4unpack_1foreign = MakeSymbol("shen.unpack-foreign")
var symshen_4unpackage = MakeSymbol("shen.unpackage")
var symshen_4unpackage_emacroexpand = MakeSymbol("shen.unpackage&macroexpand")
var symshen_4unprotect = MakeSymbol("shen.unprotect")
var symshen_4unregister_1source_1form = MakeSymbol("shen.unregister-source-form")
var symshen_4unwind = MakeSymbol("shen.unwind")
var symshen_4unwind_1types = MakeSymbol("shen.unwind-types")
var symshen_4update_1assoc = MakeSymbol("shen.update-assoc")
var symshen_4update_1history = MakeSymbol("shen.update-history")
var symshen_4update_1lambdatable = MakeSymbol("shen.update-lambdatable")
var symshen_4uppercase_2 = MakeSymbol("shen.uppercase?")
var symshen_4use_1history = MakeSymbol("shen.use-history")
var symshen_4use_1type_1info = MakeSymbol("shen.use-type-info")
var symshen_4variable_1case = MakeSymbol("shen.variable-case")
var symshen_4variablecode = MakeSymbol("shen.variablecode")
var symshen_4variancy = MakeSymbol("shen.variancy")
var symshen_4variancy_1signature = MakeSymbol("shen.variancy-signature")
var symshen_4variants_2 = MakeSymbol("shen.variants?")
var symshen_4vector_1_6str = MakeSymbol("shen.vector->str")
var symshen_4vector_1dereference = MakeSymbol("shen.vector-dereference")
var symshen_4vector_1parameter = MakeSymbol("shen.vector-parameter")
var symshen_4walk = MakeSymbol("shen.walk")
var symshen_4whitespace_2 = MakeSymbol("shen.whitespace?")
var symshen_4wildcard_2 = MakeSymbol("shen.wildcard?")
var symshen_4wildcardcode = MakeSymbol("shen.wildcardcode")
var symshen_4work_1through = MakeSymbol("shen.work-through")
var symshen_4write_1chars = MakeSymbol("shen.write-chars")
var symshen_4write_1kl = MakeSymbol("shen.write-kl")
var symshen_4write_1kl_1h = MakeSymbol("shen.write-kl-h")
var symshen_4x_4launcher_4default_1handle_1result = MakeSymbol("shen.x.launcher.default-handle-result")
var symshen_4x_4launcher_4done = MakeSymbol("shen.x.launcher.done")
var symshen_4x_4launcher_4eval_1command = MakeSymbol("shen.x.launcher.eval-command")
var symshen_4x_4launcher_4eval_1command_1h = MakeSymbol("shen.x.launcher.eval-command-h")
var symshen_4x_4launcher_4eval_1flag_1map = MakeSymbol("shen.x.launcher.eval-flag-map")
var symshen_4x_4launcher_4eval_1string = MakeSymbol("shen.x.launcher.eval-string")
var symshen_4x_4launcher_4execute_1all = MakeSymbol("shen.x.launcher.execute-all")
var symshen_4x_4launcher_4help_1text = MakeSymbol("shen.x.launcher.help-text")
var symshen_4x_4launcher_4launch_1shen = MakeSymbol("shen.x.launcher.launch-shen")
var symshen_4x_4launcher_4main = MakeSymbol("shen.x.launcher.main")
var symshen_4x_4launcher_4quiet_1load = MakeSymbol("shen.x.launcher.quiet-load")
var symshen_4x_4launcher_4script_1command = MakeSymbol("shen.x.launcher.script-command")
var symshen_4x_4launcher_4version_1string = MakeSymbol("shen.x.launcher.version-string")
var symshen_4yacc_1_6shen = MakeSymbol("shen.yacc->shen")
var symshen_4yacc_1semantics = MakeSymbol("shen.yacc-semantics")
var symshen_4yacc_1syntax = MakeSymbol("shen.yacc-syntax")
var symshen_4zero_1place_2 = MakeSymbol("shen.zero-place?")
var symshow_1help = MakeSymbol("show-help")
var symsimple_1error = MakeSymbol("simple-error")
var symsnd = MakeSymbol("snd")
var symspecialise = MakeSymbol("specialise")
var symspy = MakeSymbol("spy")
var symspy_2 = MakeSymbol("spy?")
var symsqts = MakeSymbol("sqts")
var symstep = MakeSymbol("step")
var symstep_2 = MakeSymbol("step?")
var symsterror = MakeSymbol("sterror")
var symstinput = MakeSymbol("stinput")
var symstoutput = MakeSymbol("stoutput")
var symstr = MakeSymbol("str")
var symstream = MakeSymbol("stream")
var symstring = MakeSymbol("string")
var symstring_1_6n = MakeSymbol("string->n")
var symstring_1_6symbol = MakeSymbol("string->symbol")
var symstring_2 = MakeSymbol("string?")
var symsubst = MakeSymbol("subst")
var symsuccess = MakeSymbol("success")
var symsum = MakeSymbol("sum")
var symsymbol = MakeSymbol("symbol")
var symsymbol_2 = MakeSymbol("symbol?")
var symsynonyms = MakeSymbol("synonyms")
var symsystem_1S_2 = MakeSymbol("system-S?")
var symsystemf = MakeSymbol("systemf")
var symtail = MakeSymbol("tail")
var symtc = MakeSymbol("tc")
var symtc_2 = MakeSymbol("tc?")
var symthaw = MakeSymbol("thaw")
var symtime = MakeSymbol("time")
var symtl = MakeSymbol("tl")
var symtlstr = MakeSymbol("tlstr")
var symtlv = MakeSymbol("tlv")
var symtrack = MakeSymbol("track")
var symtracked = MakeSymbol("tracked")
var symtrap_1error = MakeSymbol("trap-error")
var symtuple_2 = MakeSymbol("tuple?")
var symtype = MakeSymbol("type")
var symu_b = MakeSymbol("u!")
var symunabsolute = MakeSymbol("unabsolute")
var symundefmacro = MakeSymbol("undefmacro")
var symunion = MakeSymbol("union")
var symunit = MakeSymbol("unit")
var symunix = MakeSymbol("unix")
var symunknown_1arguments = MakeSymbol("unknown-arguments")
var symunprofile = MakeSymbol("unprofile")
var symunput = MakeSymbol("unput")
var symunspecialise = MakeSymbol("unspecialise")
var symuntrack = MakeSymbol("untrack")
var symupdate_1lambda_1table = MakeSymbol("update-lambda-table")
var symuserdefs = MakeSymbol("userdefs")
var symvalue = MakeSymbol("value")
var symvar_2 = MakeSymbol("var?")
var symvariable_2 = MakeSymbol("variable?")
var symvector = MakeSymbol("vector")
var symvector_1_6 = MakeSymbol("vector->")
var symvector_2 = MakeSymbol("vector?")
var symverified = MakeSymbol("verified")
var symversion = MakeSymbol("version")
var symwarn = MakeSymbol("warn")
var symwhen = MakeSymbol("when")
var symwhere = MakeSymbol("where")
var symwrite_1byte = MakeSymbol("write-byte")
var symwrite_1string = MakeSymbol("write-string")
var symwrite_1to_1file = MakeSymbol("write-to-file")
var symy_1or_1n_2 = MakeSymbol("y-or-n?")
var sym_i = MakeSymbol("{")
var sym_j = MakeSymbol("}")
