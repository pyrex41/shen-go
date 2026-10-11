package main

import . "github.com/pyrex41/shen-go/kl"

var LauncherMain = MakeNative(func(__e *ControlFlow) {
_ = MakeString("Copyright (c) 2019 Bruno Deferrari.\nBSD 3-Clause License: http://opensource.org/licenses/BSD-3-Clause")

tmp21968 := MakeNative(func(__e *ControlFlow) {
V7251 := __e.Get(1)
_ = V7251
tmp21969 := Call(__e, PrimFunc(symread_1file), V7251)


let__21964 := tmp21969
_ = let__21964

tmp21970 := MakeNative(func(__e *ControlFlow) {
Z7253 := __e.Get(1)
_ = Z7253
__e.TailApply(PrimFunc(symeval), Z7253)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp21970, let__21964)
return


}, 1)

tmp21971 := Call(__e, ns2_1set, symshen_4x_4launcher_4quiet_1load, tmp21968)


_ = tmp21971

tmp21972 := MakeNative(func(__e *ControlFlow) {
tmp21973 := Call(__e, PrimFunc(symversion))


tmp21974 := Call(__e, PrimFunc(symlanguage))


tmp21975 := Call(__e, PrimFunc(symport))


tmp21976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21975, Nil)
}
__typedArg0 := tmp21975
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21974, tmp21976)
}
__typedArg0 := tmp21974
__typedArg1 := tmp21976
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21978 := Call(__e, PrimFunc(symimplementation))


tmp21979 := Call(__e, PrimFunc(symrelease))


tmp21980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21979, Nil)
}
__typedArg0 := tmp21979
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21978, tmp21980)
}
__typedArg0 := tmp21978
__typedArg1 := tmp21980
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21981, Nil)
}
__typedArg0 := tmp21981
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symimplementation, tmp21982)
}
__typedArg0 := symimplementation
__typedArg1 := tmp21982
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp21977, tmp21983)
}
__typedArg0 := tmp21977
__typedArg1 := tmp21983
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symport, tmp21984)
}
__typedArg0 := symport
__typedArg1 := tmp21984
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp21986 := Call(__e, PrimFunc(symshen_4app), tmp21985, MakeString("\n"), symshen_4r)


__e.TailApply(PrimFunc(symshen_4app), tmp21973, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" "))
__typedS1, __typedOK1 := TypedString(tmp21986)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" ")
__typedArg1 := tmp21986
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)
return


}, 0)

tmp21988 := Call(__e, ns2_1set, symshen_4x_4launcher_4version_1string, tmp21972)


_ = tmp21988

tmp21989 := MakeNative(func(__e *ControlFlow) {
V7254 := __e.Get(1)
_ = V7254
tmp21990 := Call(__e, PrimFunc(symshen_4app), V7254, MakeString(" [--version] [--help] <COMMAND> [<ARGS>]\n\ncommands:\n    repl\n        Launches the interactive REPL.\n        Default action if no command is supplied.\n\n    script <FILE> [<ARGS>]\n        Runs the script in FILE. *argv* is set to [FILE | ARGS].\n\n    eval <ARGS>\n        Evaluates expressions and files. ARGS are evaluated from\n        left to right and can be a combination of:\n            -e, --eval <EXPR>\n                Evaluates EXPR and prints result.\n            -l, --load <FILE>\n                Reads and evaluates FILE.\n            -q, --quiet\n                Silences interactive output.\n            -s, --set <KEY> <VALUE>\n                Evaluates KEY, VALUE and sets as global.\n            -r, --repl\n                Launches the interactive REPL after evaluating\n                all the previous expresions."), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Usage: "))
__typedS1, __typedOK1 := TypedString(tmp21990)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Usage: ")
__typedArg1 := tmp21990
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp21991 := Call(__e, ns2_1set, symshen_4x_4launcher_4help_1text, tmp21989)


_ = tmp21991

tmp21992 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V7255 := __self1
_ = V7255
tmp21999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V7255)
}
__typedArg0 := Nil
__typedArg1 := V7255
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp21999 {
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
tmp21997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7255)
}
__typedArg0 := V7255
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp21997 {
tmp21993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7255)
}
__typedArg0 := V7255
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp21994 := Call(__e, PrimFunc(symthaw), tmp21993)


_ = tmp21994

tmp21995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7255)
}
__typedArg0 := V7255
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4x_4launcher_4execute_1all) == __self {
__self1 = tmp21995
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4x_4launcher_4execute_1all), tmp21995)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4x_4launcher_4execute_1all)
return
}


}


}, 1)

tmp22000 := Call(__e, ns2_1set, symshen_4x_4launcher_4execute_1all, tmp21992)


_ = tmp22000

tmp22001 := MakeNative(func(__e *ControlFlow) {
V7256 := __e.Get(1)
_ = V7256
tmp22002 := Call(__e, PrimFunc(symread_1from_1string), V7256)


tmp22003 := Call(__e, PrimFunc(symhead), tmp22002)


__e.TailApply(PrimFunc(symeval), tmp22003)
return


}, 1)

tmp22004 := Call(__e, ns2_1set, symshen_4x_4launcher_4eval_1string, tmp22001)


_ = tmp22004

tmp22005 := MakeNative(func(__e *ControlFlow) {
V7259 := __e.Get(1)
_ = V7259
tmp22015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("-e"), V7259)
}
__typedArg0 := MakeString("-e")
__typedArg1 := V7259
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp22015 {
__e.Return(MakeString("--eval"))
return
} else {
tmp22013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("-l"), V7259)
}
__typedArg0 := MakeString("-l")
__typedArg1 := V7259
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp22013 {
__e.Return(MakeString("--load"))
return
} else {
tmp22011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("-q"), V7259)
}
__typedArg0 := MakeString("-q")
__typedArg1 := V7259
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp22011 {
__e.Return(MakeString("--quiet"))
return
} else {
tmp22009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("-s"), V7259)
}
__typedArg0 := MakeString("-s")
__typedArg1 := V7259
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp22009 {
__e.Return(MakeString("--set"))
return
} else {
tmp22007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("-r"), V7259)
}
__typedArg0 := MakeString("-r")
__typedArg1 := V7259
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp22007 {
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

tmp22016 := Call(__e, ns2_1set, symshen_4x_4launcher_4eval_1flag_1map, tmp22005)


_ = tmp22016

tmp22017 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V7264 := __self1
_ = V7264
V7265 := __self2
_ = V7265
tmp22119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V7264)
}
__typedArg0 := Nil
__typedArg1 := V7264
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp22119 {
tmp22018 := Call(__e, PrimFunc(symreverse), V7265)


__e.TailApply(PrimFunc(symshen_4x_4launcher_4execute_1all), tmp22018)
return


} else {
tmp22117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22109 Obj

if True == tmp22117 {
tmp22115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--eval"), tmp22115)
}
__typedArg0 := MakeString("--eval")
__typedArg1 := tmp22115
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22111 Obj

if True == tmp22116 {
tmp22113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22113)
}
__typedArg0 := tmp22113
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22112 Obj

if True == tmp22114 {
ifres22112 = True


} else {
ifres22112 = False


}

ifres22111 = ifres22112


} else {
ifres22111 = False


}

var ifres22110 Obj

if True == ifres22111 {
ifres22110 = True


} else {
ifres22110 = False


}

ifres22109 = ifres22110


} else {
ifres22109 = False


}

if True == ifres22109 {
tmp22019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22019)
}
__typedArg0 := tmp22019
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22021 := MakeNative(func(__e *ControlFlow) {
tmp22022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22022)
}
__typedArg0 := tmp22022
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22024 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1string), tmp22023)


tmp22025 := Call(__e, PrimFunc(symshen_4app), tmp22024, MakeString("\n"), symshen_4a)


tmp22026 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp22025, tmp22026)
return


}, 0)

tmp22027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp22021, V7265)
}
__typedArg0 := tmp22021
__typedArg1 := V7265
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4x_4launcher_4eval_1command_1h) == __self {
__self1, __self2 = tmp22020, tmp22027
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp22020, tmp22027)
return


} else {
tmp22107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22099 Obj

if True == tmp22107 {
tmp22105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--load"), tmp22105)
}
__typedArg0 := MakeString("--load")
__typedArg1 := tmp22105
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22101 Obj

if True == tmp22106 {
tmp22103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22103)
}
__typedArg0 := tmp22103
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22102 Obj

if True == tmp22104 {
ifres22102 = True


} else {
ifres22102 = False


}

ifres22101 = ifres22102


} else {
ifres22101 = False


}

var ifres22100 Obj

if True == ifres22101 {
ifres22100 = True


} else {
ifres22100 = False


}

ifres22099 = ifres22100


} else {
ifres22099 = False


}

if True == ifres22099 {
tmp22028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22028)
}
__typedArg0 := tmp22028
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22030 := MakeNative(func(__e *ControlFlow) {
tmp22031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22031)
}
__typedArg0 := tmp22031
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symload), tmp22032)
return


}, 0)

tmp22033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp22030, V7265)
}
__typedArg0 := tmp22030
__typedArg1 := V7265
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4x_4launcher_4eval_1command_1h) == __self {
__self1, __self2 = tmp22029, tmp22033
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp22029, tmp22033)
return


} else {
tmp22097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22093 Obj

if True == tmp22097 {
tmp22095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--quiet"), tmp22095)
}
__typedArg0 := MakeString("--quiet")
__typedArg1 := tmp22095
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22094 Obj

if True == tmp22096 {
ifres22094 = True


} else {
ifres22094 = False


}

ifres22093 = ifres22094


} else {
ifres22093 = False


}

if True == ifres22093 {
tmp22034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22035 := MakeNative(func(__e *ControlFlow) {
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

tmp22036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp22035, V7265)
}
__typedArg0 := tmp22035
__typedArg1 := V7265
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4x_4launcher_4eval_1command_1h) == __self {
__self1, __self2 = tmp22034, tmp22036
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp22034, tmp22036)
return


} else {
tmp22091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22078 Obj

if True == tmp22091 {
tmp22089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--set"), tmp22089)
}
__typedArg0 := MakeString("--set")
__typedArg1 := tmp22089
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22080 Obj

if True == tmp22090 {
tmp22087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22087)
}
__typedArg0 := tmp22087
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22082 Obj

if True == tmp22088 {
tmp22084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22084)
}
__typedArg0 := tmp22084
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22085)
}
__typedArg0 := tmp22085
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22083 Obj

if True == tmp22086 {
ifres22083 = True


} else {
ifres22083 = False


}

ifres22082 = ifres22083


} else {
ifres22082 = False


}

var ifres22081 Obj

if True == ifres22082 {
ifres22081 = True


} else {
ifres22081 = False


}

ifres22080 = ifres22081


} else {
ifres22080 = False


}

var ifres22079 Obj

if True == ifres22080 {
ifres22079 = True


} else {
ifres22079 = False


}

ifres22078 = ifres22079


} else {
ifres22078 = False


}

if True == ifres22078 {
tmp22037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22037)
}
__typedArg0 := tmp22037
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22038)
}
__typedArg0 := tmp22038
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22040 := MakeNative(func(__e *ControlFlow) {
tmp22041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22041)
}
__typedArg0 := tmp22041
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22043 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1string), tmp22042)


tmp22044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22044)
}
__typedArg0 := tmp22044
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22045)
}
__typedArg0 := tmp22045
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22047 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1string), tmp22046)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(tmp22043, tmp22047)
}
__typedArg0 := tmp22043
__typedArg1 := tmp22047
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp22048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp22040, V7265)
}
__typedArg0 := tmp22040
__typedArg1 := V7265
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4x_4launcher_4eval_1command_1h) == __self {
__self1, __self2 = tmp22039, tmp22048
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp22039, tmp22048)
return


} else {
tmp22076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22072 Obj

if True == tmp22076 {
tmp22074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--repl"), tmp22074)
}
__typedArg0 := MakeString("--repl")
__typedArg1 := tmp22074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22073 Obj

if True == tmp22075 {
ifres22073 = True


} else {
ifres22073 = False


}

ifres22072 = ifres22073


} else {
ifres22072 = False


}

if True == ifres22072 {
tmp22049 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1command_1h), Nil, V7265)


_ = tmp22049

tmp22050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlaunch_1repl, tmp22050)
}
__typedArg0 := symlaunch_1repl
__typedArg1 := tmp22050
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp22051 := MakeNative(func(__e *ControlFlow) {
tmp22057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp22057 {
tmp22052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22053 := Call(__e, PrimFunc(symshen_4app), tmp22052, MakeString(""), symshen_4a)


tmp22055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Invalid eval argument: "))
__typedS1, __typedOK1 := TypedString(tmp22053)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Invalid eval argument: ")
__typedArg1 := tmp22053
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), Nil)
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Invalid eval argument: "))
__typedS1, __typedOK1 := TypedString(tmp22053)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Invalid eval argument: ")
__typedArg1 := tmp22053
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symerror, tmp22055)
}
__typedArg0 := symerror
__typedArg1 := tmp22055
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4x_4launcher_4eval_1command_1h)
return
}


}, 0)

let__21965 := tmp22051
_ = let__21965

tmp22070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp22070 {
tmp22058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22059 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1flag_1map), tmp22058)


let__21967 := tmp22059
_ = let__21967

tmp22065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, let__21967)
}
__typedArg0 := False
__typedArg1 := let__21967
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22060 Obj

if True == tmp22065 {
tmp22061 := Call(__e, PrimFunc(symfail))


ifres22060 = tmp22061


} else {
tmp22062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7264)
}
__typedArg0 := V7264
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__21967, tmp22062)
}
__typedArg0 := let__21967
__typedArg1 := tmp22062
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22064 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp22063, V7265)


ifres22060 = tmp22064


}

let__21966 := ifres22060
_ = let__21966

tmp22067 := Call(__e, PrimFunc(symfail))


tmp22068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__21966, tmp22067)
}
__typedArg0 := let__21966
__typedArg1 := tmp22067
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp22068 {
__e.TailApply(PrimFunc(symthaw), let__21965)
return
} else {
__e.Return(let__21966)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), let__21965)
return
}


}


}


}


}


}


}


}, 2)

tmp22120 := Call(__e, ns2_1set, symshen_4x_4launcher_4eval_1command_1h, tmp22017)


_ = tmp22120

tmp22121 := MakeNative(func(__e *ControlFlow) {
V7269 := __e.Get(1)
_ = V7269
__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), V7269, Nil)
return
}, 1)

tmp22122 := Call(__e, ns2_1set, symshen_4x_4launcher_4eval_1command, tmp22121)


_ = tmp22122

tmp22123 := MakeNative(func(__e *ControlFlow) {
V7270 := __e.Get(1)
_ = V7270
V7271 := __e.Get(2)
_ = V7271
tmp22124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V7270, V7271)
}
__typedArg0 := V7270
__typedArg1 := V7271
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp22125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dargv_d, tmp22124)
}
__typedArg0 := sym_dargv_d
__typedArg1 := tmp22124
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp22125

tmp22126 := Call(__e, PrimFunc(symshen_4x_4launcher_4quiet_1load), V7270)


_ = tmp22126

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

tmp22127 := Call(__e, ns2_1set, symshen_4x_4launcher_4script_1command, tmp22123)


_ = tmp22127

tmp22128 := MakeNative(func(__e *ControlFlow) {
V7272 := __e.Get(1)
_ = V7272
tmp22215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22211 Obj

if True == tmp22215 {
tmp22213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp22213)
}
__typedArg0 := Nil
__typedArg1 := tmp22213
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22212 Obj

if True == tmp22214 {
ifres22212 = True


} else {
ifres22212 = False


}

ifres22211 = ifres22212


} else {
ifres22211 = False


}

if True == ifres22211 {
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
tmp22209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22200 Obj

if True == tmp22209 {
tmp22207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22207)
}
__typedArg0 := tmp22207
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22202 Obj

if True == tmp22208 {
tmp22204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22204)
}
__typedArg0 := tmp22204
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--help"), tmp22205)
}
__typedArg0 := MakeString("--help")
__typedArg1 := tmp22205
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22203 Obj

if True == tmp22206 {
ifres22203 = True


} else {
ifres22203 = False


}

ifres22202 = ifres22203


} else {
ifres22202 = False


}

var ifres22201 Obj

if True == ifres22202 {
ifres22201 = True


} else {
ifres22201 = False


}

ifres22200 = ifres22201


} else {
ifres22200 = False


}

if True == ifres22200 {
tmp22129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22130 := Call(__e, PrimFunc(symshen_4x_4launcher_4help_1text), tmp22129)


tmp22131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp22130, Nil)
}
__typedArg0 := tmp22130
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshow_1help, tmp22131)
}
__typedArg0 := symshow_1help
__typedArg1 := tmp22131
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp22198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22189 Obj

if True == tmp22198 {
tmp22196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22196)
}
__typedArg0 := tmp22196
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22191 Obj

if True == tmp22197 {
tmp22193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22193)
}
__typedArg0 := tmp22193
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--version"), tmp22194)
}
__typedArg0 := MakeString("--version")
__typedArg1 := tmp22194
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22192 Obj

if True == tmp22195 {
ifres22192 = True


} else {
ifres22192 = False


}

ifres22191 = ifres22192


} else {
ifres22191 = False


}

var ifres22190 Obj

if True == ifres22191 {
ifres22190 = True


} else {
ifres22190 = False


}

ifres22189 = ifres22190


} else {
ifres22189 = False


}

if True == ifres22189 {
tmp22132 := Call(__e, PrimFunc(symshen_4x_4launcher_4version_1string))


tmp22133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp22132, Nil)
}
__typedArg0 := tmp22132
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsuccess, tmp22133)
}
__typedArg0 := symsuccess
__typedArg1 := tmp22133
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp22187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22178 Obj

if True == tmp22187 {
tmp22185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22185)
}
__typedArg0 := tmp22185
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22180 Obj

if True == tmp22186 {
tmp22182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22182)
}
__typedArg0 := tmp22182
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("repl"), tmp22183)
}
__typedArg0 := MakeString("repl")
__typedArg1 := tmp22183
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22181 Obj

if True == tmp22184 {
ifres22181 = True


} else {
ifres22181 = False


}

ifres22180 = ifres22181


} else {
ifres22180 = False


}

var ifres22179 Obj

if True == ifres22180 {
ifres22179 = True


} else {
ifres22179 = False


}

ifres22178 = ifres22179


} else {
ifres22178 = False


}

if True == ifres22178 {
tmp22134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22134)
}
__typedArg0 := tmp22134
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlaunch_1repl, tmp22135)
}
__typedArg0 := symlaunch_1repl
__typedArg1 := tmp22135
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp22176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22162 Obj

if True == tmp22176 {
tmp22174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22174)
}
__typedArg0 := tmp22174
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22164 Obj

if True == tmp22175 {
tmp22171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22171)
}
__typedArg0 := tmp22171
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("script"), tmp22172)
}
__typedArg0 := MakeString("script")
__typedArg1 := tmp22172
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22166 Obj

if True == tmp22173 {
tmp22168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22168)
}
__typedArg0 := tmp22168
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22169)
}
__typedArg0 := tmp22169
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22167 Obj

if True == tmp22170 {
ifres22167 = True


} else {
ifres22167 = False


}

ifres22166 = ifres22167


} else {
ifres22166 = False


}

var ifres22165 Obj

if True == ifres22166 {
ifres22165 = True


} else {
ifres22165 = False


}

ifres22164 = ifres22165


} else {
ifres22164 = False


}

var ifres22163 Obj

if True == ifres22164 {
ifres22163 = True


} else {
ifres22163 = False


}

ifres22162 = ifres22163


} else {
ifres22162 = False


}

if True == ifres22162 {
tmp22136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22136)
}
__typedArg0 := tmp22136
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22137)
}
__typedArg0 := tmp22137
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22139)
}
__typedArg0 := tmp22139
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22140)
}
__typedArg0 := tmp22140
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4x_4launcher_4script_1command), tmp22138, tmp22141)
return


} else {
tmp22160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22151 Obj

if True == tmp22160 {
tmp22158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22158)
}
__typedArg0 := tmp22158
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22153 Obj

if True == tmp22159 {
tmp22155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22155)
}
__typedArg0 := tmp22155
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("eval"), tmp22156)
}
__typedArg0 := MakeString("eval")
__typedArg1 := tmp22156
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22154 Obj

if True == tmp22157 {
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
tmp22142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22142)
}
__typedArg0 := tmp22142
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command), tmp22143)
return


} else {
tmp22149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22145 Obj

if True == tmp22149 {
tmp22147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7272)
}
__typedArg0 := V7272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22147)
}
__typedArg0 := tmp22147
return Call(__e, PrimFunc(symcons_2), __typedArg0)
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
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symunknown_1arguments, V7272)
}
__typedArg0 := symunknown_1arguments
__typedArg1 := V7272
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

tmp22216 := Call(__e, ns2_1set, symshen_4x_4launcher_4launch_1shen, tmp22128)


_ = tmp22216

tmp22217 := MakeNative(func(__e *ControlFlow) {
V7275 := __e.Get(1)
_ = V7275
tmp22316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22308 Obj

if True == tmp22316 {
tmp22314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsuccess, tmp22314)
}
__typedArg0 := symsuccess
__typedArg1 := tmp22314
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22310 Obj

if True == tmp22315 {
tmp22312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp22312)
}
__typedArg0 := Nil
__typedArg1 := tmp22312
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22311 Obj

if True == tmp22313 {
ifres22311 = True


} else {
ifres22311 = False


}

ifres22310 = ifres22311


} else {
ifres22310 = False


}

var ifres22309 Obj

if True == ifres22310 {
ifres22309 = True


} else {
ifres22309 = False


}

ifres22308 = ifres22309


} else {
ifres22308 = False


}

if True == ifres22308 {
__e.Return(symshen_4x_4launcher_4done)
return
} else {
tmp22306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22293 Obj

if True == tmp22306 {
tmp22304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsuccess, tmp22304)
}
__typedArg0 := symsuccess
__typedArg1 := tmp22304
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22295 Obj

if True == tmp22305 {
tmp22302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22302)
}
__typedArg0 := tmp22302
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22297 Obj

if True == tmp22303 {
tmp22299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22299)
}
__typedArg0 := tmp22299
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp22300)
}
__typedArg0 := Nil
__typedArg1 := tmp22300
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22298 Obj

if True == tmp22301 {
ifres22298 = True


} else {
ifres22298 = False


}

ifres22297 = ifres22298


} else {
ifres22297 = False


}

var ifres22296 Obj

if True == ifres22297 {
ifres22296 = True


} else {
ifres22296 = False


}

ifres22295 = ifres22296


} else {
ifres22295 = False


}

var ifres22294 Obj

if True == ifres22295 {
ifres22294 = True


} else {
ifres22294 = False


}

ifres22293 = ifres22294


} else {
ifres22293 = False


}

if True == ifres22293 {
tmp22218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22218)
}
__typedArg0 := tmp22218
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22220 := Call(__e, PrimFunc(symshen_4app), tmp22219, MakeString("\n"), symshen_4a)


tmp22221 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp22220, tmp22221)
return


} else {
tmp22291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22278 Obj

if True == tmp22291 {
tmp22289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symerror, tmp22289)
}
__typedArg0 := symerror
__typedArg1 := tmp22289
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22280 Obj

if True == tmp22290 {
tmp22287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22287)
}
__typedArg0 := tmp22287
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22282 Obj

if True == tmp22288 {
tmp22284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22284)
}
__typedArg0 := tmp22284
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp22285)
}
__typedArg0 := Nil
__typedArg1 := tmp22285
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22283 Obj

if True == tmp22286 {
ifres22283 = True


} else {
ifres22283 = False


}

ifres22282 = ifres22283


} else {
ifres22282 = False


}

var ifres22281 Obj

if True == ifres22282 {
ifres22281 = True


} else {
ifres22281 = False


}

ifres22280 = ifres22281


} else {
ifres22280 = False


}

var ifres22279 Obj

if True == ifres22280 {
ifres22279 = True


} else {
ifres22279 = False


}

ifres22278 = ifres22279


} else {
ifres22278 = False


}

if True == ifres22278 {
tmp22222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22222)
}
__typedArg0 := tmp22222
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22224 := Call(__e, PrimFunc(symshen_4app), tmp22223, MakeString("\n"), symshen_4a)


tmp22225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("ERROR: "))
__typedS1, __typedOK1 := TypedString(tmp22224)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("ERROR: ")
__typedArg1 := tmp22224
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp22226 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp22225, tmp22226)
return


} else {
tmp22276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22272 Obj

if True == tmp22276 {
tmp22274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlaunch_1repl, tmp22274)
}
__typedArg0 := symlaunch_1repl
__typedArg1 := tmp22274
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22273 Obj

if True == tmp22275 {
ifres22273 = True


} else {
ifres22273 = False


}

ifres22272 = ifres22273


} else {
ifres22272 = False


}

if True == ifres22272 {
__e.TailApply(PrimFunc(symshen_4repl))
return
} else {
tmp22270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22257 Obj

if True == tmp22270 {
tmp22268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshow_1help, tmp22268)
}
__typedArg0 := symshow_1help
__typedArg1 := tmp22268
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22259 Obj

if True == tmp22269 {
tmp22266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22266)
}
__typedArg0 := tmp22266
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22261 Obj

if True == tmp22267 {
tmp22263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22263)
}
__typedArg0 := tmp22263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp22264)
}
__typedArg0 := Nil
__typedArg1 := tmp22264
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22262 Obj

if True == tmp22265 {
ifres22262 = True


} else {
ifres22262 = False


}

ifres22261 = ifres22262


} else {
ifres22261 = False


}

var ifres22260 Obj

if True == ifres22261 {
ifres22260 = True


} else {
ifres22260 = False


}

ifres22259 = ifres22260


} else {
ifres22259 = False


}

var ifres22258 Obj

if True == ifres22259 {
ifres22258 = True


} else {
ifres22258 = False


}

ifres22257 = ifres22258


} else {
ifres22257 = False


}

if True == ifres22257 {
tmp22227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22227)
}
__typedArg0 := tmp22227
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22229 := Call(__e, PrimFunc(symshen_4app), tmp22228, MakeString("\n"), symshen_4a)


tmp22230 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp22229, tmp22230)
return


} else {
tmp22255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22242 Obj

if True == tmp22255 {
tmp22253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symunknown_1arguments, tmp22253)
}
__typedArg0 := symunknown_1arguments
__typedArg1 := tmp22253
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres22244 Obj

if True == tmp22254 {
tmp22251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22251)
}
__typedArg0 := tmp22251
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22246 Obj

if True == tmp22252 {
tmp22248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22248)
}
__typedArg0 := tmp22248
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp22249)
}
__typedArg0 := tmp22249
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres22247 Obj

if True == tmp22250 {
ifres22247 = True


} else {
ifres22247 = False


}

ifres22246 = ifres22247


} else {
ifres22246 = False


}

var ifres22245 Obj

if True == ifres22246 {
ifres22245 = True


} else {
ifres22245 = False


}

ifres22244 = ifres22245


} else {
ifres22244 = False


}

var ifres22243 Obj

if True == ifres22244 {
ifres22243 = True


} else {
ifres22243 = False


}

ifres22242 = ifres22243


} else {
ifres22242 = False


}

if True == ifres22242 {
tmp22231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp22231)
}
__typedArg0 := tmp22231
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22232)
}
__typedArg0 := tmp22232
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7275)
}
__typedArg0 := V7275
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp22235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp22234)
}
__typedArg0 := tmp22234
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp22236 := Call(__e, PrimFunc(symshen_4app), tmp22235, MakeString(" --help' for more information.\n"), symshen_4a)


tmp22238 := Call(__e, PrimFunc(symshen_4app), tmp22233, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\nTry `"))
__typedS1, __typedOK1 := TypedString(tmp22236)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\nTry `")
__typedArg1 := tmp22236
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp22239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("ERROR: Invalid argument: "))
__typedS1, __typedOK1 := TypedString(tmp22238)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("ERROR: Invalid argument: ")
__typedArg1 := tmp22238
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp22240 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp22239, tmp22240)
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

tmp22317 := Call(__e, ns2_1set, symshen_4x_4launcher_4default_1handle_1result, tmp22217)


_ = tmp22317

tmp22318 := MakeNative(func(__e *ControlFlow) {
V7276 := __e.Get(1)
_ = V7276
tmp22319 := Call(__e, PrimFunc(symshen_4x_4launcher_4launch_1shen), V7276)


__e.TailApply(PrimFunc(symshen_4x_4launcher_4default_1handle_1result), tmp22319)
return


}, 1)

__e.TailApply(ns2_1set, symshen_4x_4launcher_4main, tmp22318)
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
var symshen_4_dsystem_d = MakeSymbol("shen.*system*")
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
var symshen_4byte_1_6digit = MakeSymbol("shen.byte->digit")
var symshen_4bytes_1_6string = MakeSymbol("shen.bytes->string")
var symshen_4c_1rule_1_6shen = MakeSymbol("shen.c-rule->shen")
var symshen_4c_1rules_1_6shen = MakeSymbol("shen.c-rules->shen")
var symshen_4call_1dynamic = MakeSymbol("shen.call-dynamic")
var symshen_4call_1prolog = MakeSymbol("shen.call-prolog")
var symshen_4callrec = MakeSymbol("shen.callrec")
var symshen_4ccons_2 = MakeSymbol("shen.ccons?")
var symshen_4char_1stinput_2 = MakeSymbol("shen.char-stinput?")
var symshen_4char_1stoutput_2 = MakeSymbol("shen.char-stoutput?")
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
var symshen_4ctxt = MakeSymbol("shen.ctxt")
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
var symshen_4g = MakeSymbol("shen.g")
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
var symshen_4included = MakeSymbol("shen.included")
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
var symshen_4read_1unit_1string = MakeSymbol("shen.read-unit-string")
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
var symshen_4spy_2 = MakeSymbol("shen.spy?")
var symshen_4step_2 = MakeSymbol("shen.step?")
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
var symshen_4unix = MakeSymbol("shen.unix")
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
var symshen_4write_1string = MakeSymbol("shen.write-string")
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
var symsqts = MakeSymbol("sqts")
var symstep = MakeSymbol("step")
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
var symwrite_1to_1file = MakeSymbol("write-to-file")
var symy_1or_1n_2 = MakeSymbol("y-or-n?")
var sym_i = MakeSymbol("{")
var sym_j = MakeSymbol("}")
