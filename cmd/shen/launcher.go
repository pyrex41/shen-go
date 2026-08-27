package main

import . "github.com/pyrex41/shen-go/kl"

var LauncherMain = MakeNative(func(__e *ControlFlow) {
_ = MakeString("Copyright (c) 2019 Bruno Deferrari.\nBSD 3-Clause License: http://opensource.org/licenses/BSD-3-Clause")

tmp19823 := MakeNative(func(__e *ControlFlow) {
V7102 := __e.Get(1)
_ = V7102
tmp19824 := MakeNative(func(__e *ControlFlow) {
W7103 := __e.Get(1)
_ = W7103
tmp19825 := MakeNative(func(__e *ControlFlow) {
Z7104 := __e.Get(1)
_ = Z7104
__e.TailApply(PrimFunc(symeval), Z7104)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp19825, W7103)
return


}, 1)

tmp19826 := Call(__e, PrimFunc(symread_1file), V7102)


__e.TailApply(tmp19824, tmp19826)
return


}, 1)

tmp19827 := Call(__e, ns2_1set, symshen_4x_4launcher_4quiet_1load, tmp19823)


_ = tmp19827

tmp19828 := MakeNative(func(__e *ControlFlow) {
tmp19829 := Call(__e, PrimFunc(symversion))


tmp19830 := Call(__e, PrimFunc(symlanguage))


tmp19831 := Call(__e, PrimFunc(symport))


tmp19832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19831, Nil)
}
__typedArg0 := tmp19831
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19830, tmp19832)
}
__typedArg0 := tmp19830
__typedArg1 := tmp19832
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19834 := Call(__e, PrimFunc(symimplementation))


tmp19835 := Call(__e, PrimFunc(symrelease))


tmp19836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19835, Nil)
}
__typedArg0 := tmp19835
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19834, tmp19836)
}
__typedArg0 := tmp19834
__typedArg1 := tmp19836
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19837, Nil)
}
__typedArg0 := tmp19837
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symimplementation, tmp19838)
}
__typedArg0 := symimplementation
__typedArg1 := tmp19838
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19833, tmp19839)
}
__typedArg0 := tmp19833
__typedArg1 := tmp19839
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symport, tmp19840)
}
__typedArg0 := symport
__typedArg1 := tmp19840
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19842 := Call(__e, PrimFunc(symshen_4app), tmp19841, MakeString("\n"), symshen_4r)


__e.TailApply(PrimFunc(symshen_4app), tmp19829, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" "))
__typedS1, __typedOK1 := TypedString(tmp19842)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" ")
__typedArg1 := tmp19842
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)
return


}, 0)

tmp19844 := Call(__e, ns2_1set, symshen_4x_4launcher_4version_1string, tmp19828)


_ = tmp19844

tmp19845 := MakeNative(func(__e *ControlFlow) {
V7105 := __e.Get(1)
_ = V7105
tmp19846 := Call(__e, PrimFunc(symshen_4app), V7105, MakeString(" [--version] [--help] <COMMAND> [<ARGS>]\n\ncommands:\n    repl\n        Launches the interactive REPL.\n        Default action if no command is supplied.\n\n    script <FILE> [<ARGS>]\n        Runs the script in FILE. *argv* is set to [FILE | ARGS].\n\n    eval <ARGS>\n        Evaluates expressions and files. ARGS are evaluated from\n        left to right and can be a combination of:\n            -e, --eval <EXPR>\n                Evaluates EXPR and prints result.\n            -l, --load <FILE>\n                Reads and evaluates FILE.\n            -q, --quiet\n                Silences interactive output.\n            -s, --set <KEY> <VALUE>\n                Evaluates KEY, VALUE and sets as global.\n            -r, --repl\n                Launches the interactive REPL after evaluating\n                all the previous expresions."), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Usage: "))
__typedS1, __typedOK1 := TypedString(tmp19846)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Usage: ")
__typedArg1 := tmp19846
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp19847 := Call(__e, ns2_1set, symshen_4x_4launcher_4help_1text, tmp19845)


_ = tmp19847

tmp19848 := MakeNative(func(__e *ControlFlow) {
V7106 := __e.Get(1)
_ = V7106
tmp19855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V7106)
}
__typedArg0 := Nil
__typedArg1 := V7106
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp19855 {
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
tmp19853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7106)
}
__typedArg0 := V7106
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp19853 {
tmp19849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7106)
}
__typedArg0 := V7106
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp19850 := Call(__e, PrimFunc(symthaw), tmp19849)


_ = tmp19850

tmp19851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7106)
}
__typedArg0 := V7106
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4x_4launcher_4execute_1all), tmp19851)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4x_4launcher_4execute_1all)
return
}


}


}, 1)

tmp19856 := Call(__e, ns2_1set, symshen_4x_4launcher_4execute_1all, tmp19848)


_ = tmp19856

tmp19857 := MakeNative(func(__e *ControlFlow) {
V7107 := __e.Get(1)
_ = V7107
tmp19858 := Call(__e, PrimFunc(symread_1from_1string), V7107)


tmp19859 := Call(__e, PrimFunc(symhead), tmp19858)


__e.TailApply(PrimFunc(symeval), tmp19859)
return


}, 1)

tmp19860 := Call(__e, ns2_1set, symshen_4x_4launcher_4eval_1string, tmp19857)


_ = tmp19860

tmp19861 := MakeNative(func(__e *ControlFlow) {
V7110 := __e.Get(1)
_ = V7110
tmp19871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("-e"), V7110)
}
__typedArg0 := MakeString("-e")
__typedArg1 := V7110
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp19871 {
__e.Return(MakeString("--eval"))
return
} else {
tmp19869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("-l"), V7110)
}
__typedArg0 := MakeString("-l")
__typedArg1 := V7110
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp19869 {
__e.Return(MakeString("--load"))
return
} else {
tmp19867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("-q"), V7110)
}
__typedArg0 := MakeString("-q")
__typedArg1 := V7110
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp19867 {
__e.Return(MakeString("--quiet"))
return
} else {
tmp19865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("-s"), V7110)
}
__typedArg0 := MakeString("-s")
__typedArg1 := V7110
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp19865 {
__e.Return(MakeString("--set"))
return
} else {
tmp19863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("-r"), V7110)
}
__typedArg0 := MakeString("-r")
__typedArg1 := V7110
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp19863 {
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

tmp19872 := Call(__e, ns2_1set, symshen_4x_4launcher_4eval_1flag_1map, tmp19861)


_ = tmp19872

tmp19873 := MakeNative(func(__e *ControlFlow) {
V7115 := __e.Get(1)
_ = V7115
V7116 := __e.Get(2)
_ = V7116
tmp19977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V7115)
}
__typedArg0 := Nil
__typedArg1 := V7115
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp19977 {
tmp19874 := Call(__e, PrimFunc(symreverse), V7116)


__e.TailApply(PrimFunc(symshen_4x_4launcher_4execute_1all), tmp19874)
return


} else {
tmp19975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres19967 Obj

if True == tmp19975 {
tmp19973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp19974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--eval"), tmp19973)
}
__typedArg0 := MakeString("--eval")
__typedArg1 := tmp19973
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres19969 Obj

if True == tmp19974 {
tmp19971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp19971)
}
__typedArg0 := tmp19971
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres19970 Obj

if True == tmp19972 {
ifres19970 = True


} else {
ifres19970 = False


}

ifres19969 = ifres19970


} else {
ifres19969 = False


}

var ifres19968 Obj

if True == ifres19969 {
ifres19968 = True


} else {
ifres19968 = False


}

ifres19967 = ifres19968


} else {
ifres19967 = False


}

if True == ifres19967 {
tmp19875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp19875)
}
__typedArg0 := tmp19875
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19877 := MakeNative(func(__e *ControlFlow) {
tmp19878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp19878)
}
__typedArg0 := tmp19878
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp19880 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1string), tmp19879)


tmp19881 := Call(__e, PrimFunc(symshen_4app), tmp19880, MakeString("\n"), symshen_4a)


tmp19882 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp19881, tmp19882)
return


}, 0)

tmp19883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19877, V7116)
}
__typedArg0 := tmp19877
__typedArg1 := V7116
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp19876, tmp19883)
return


} else {
tmp19965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres19957 Obj

if True == tmp19965 {
tmp19963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp19964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--load"), tmp19963)
}
__typedArg0 := MakeString("--load")
__typedArg1 := tmp19963
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres19959 Obj

if True == tmp19964 {
tmp19961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp19961)
}
__typedArg0 := tmp19961
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres19960 Obj

if True == tmp19962 {
ifres19960 = True


} else {
ifres19960 = False


}

ifres19959 = ifres19960


} else {
ifres19959 = False


}

var ifres19958 Obj

if True == ifres19959 {
ifres19958 = True


} else {
ifres19958 = False


}

ifres19957 = ifres19958


} else {
ifres19957 = False


}

if True == ifres19957 {
tmp19884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp19884)
}
__typedArg0 := tmp19884
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19886 := MakeNative(func(__e *ControlFlow) {
tmp19887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp19887)
}
__typedArg0 := tmp19887
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symload), tmp19888)
return


}, 0)

tmp19889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19886, V7116)
}
__typedArg0 := tmp19886
__typedArg1 := V7116
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp19885, tmp19889)
return


} else {
tmp19955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres19951 Obj

if True == tmp19955 {
tmp19953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp19954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--quiet"), tmp19953)
}
__typedArg0 := MakeString("--quiet")
__typedArg1 := tmp19953
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres19952 Obj

if True == tmp19954 {
ifres19952 = True


} else {
ifres19952 = False


}

ifres19951 = ifres19952


} else {
ifres19951 = False


}

if True == ifres19951 {
tmp19890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19891 := MakeNative(func(__e *ControlFlow) {
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

tmp19892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19891, V7116)
}
__typedArg0 := tmp19891
__typedArg1 := V7116
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp19890, tmp19892)
return


} else {
tmp19949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres19936 Obj

if True == tmp19949 {
tmp19947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp19948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--set"), tmp19947)
}
__typedArg0 := MakeString("--set")
__typedArg1 := tmp19947
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres19938 Obj

if True == tmp19948 {
tmp19945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp19945)
}
__typedArg0 := tmp19945
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres19940 Obj

if True == tmp19946 {
tmp19942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp19942)
}
__typedArg0 := tmp19942
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp19943)
}
__typedArg0 := tmp19943
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres19941 Obj

if True == tmp19944 {
ifres19941 = True


} else {
ifres19941 = False


}

ifres19940 = ifres19941


} else {
ifres19940 = False


}

var ifres19939 Obj

if True == ifres19940 {
ifres19939 = True


} else {
ifres19939 = False


}

ifres19938 = ifres19939


} else {
ifres19938 = False


}

var ifres19937 Obj

if True == ifres19938 {
ifres19937 = True


} else {
ifres19937 = False


}

ifres19936 = ifres19937


} else {
ifres19936 = False


}

if True == ifres19936 {
tmp19893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp19893)
}
__typedArg0 := tmp19893
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp19894)
}
__typedArg0 := tmp19894
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19896 := MakeNative(func(__e *ControlFlow) {
tmp19897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp19897)
}
__typedArg0 := tmp19897
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp19899 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1string), tmp19898)


tmp19900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp19900)
}
__typedArg0 := tmp19900
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp19901)
}
__typedArg0 := tmp19901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp19903 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1string), tmp19902)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(tmp19899, tmp19903)
}
__typedArg0 := tmp19899
__typedArg1 := tmp19903
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp19904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19896, V7116)
}
__typedArg0 := tmp19896
__typedArg1 := V7116
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp19895, tmp19904)
return


} else {
tmp19934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres19930 Obj

if True == tmp19934 {
tmp19932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp19933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--repl"), tmp19932)
}
__typedArg0 := MakeString("--repl")
__typedArg1 := tmp19932
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres19931 Obj

if True == tmp19933 {
ifres19931 = True


} else {
ifres19931 = False


}

ifres19930 = ifres19931


} else {
ifres19930 = False


}

if True == ifres19930 {
tmp19905 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1command_1h), Nil, V7116)


_ = tmp19905

tmp19906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlaunch_1repl, tmp19906)
}
__typedArg0 := symlaunch_1repl
__typedArg1 := tmp19906
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp19907 := MakeNative(func(__e *ControlFlow) {
Freeze7119 := __e.Get(1)
_ = Freeze7119
tmp19921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp19921 {
tmp19908 := MakeNative(func(__e *ControlFlow) {
Result7118 := __e.Get(1)
_ = Result7118
tmp19910 := Call(__e, PrimFunc(symfail))


tmp19911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Result7118, tmp19910)
}
__typedArg0 := Result7118
__typedArg1 := tmp19910
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp19911 {
__e.TailApply(PrimFunc(symthaw), Freeze7119)
return
} else {
__e.Return(Result7118)
return
}


}, 1)

tmp19912 := MakeNative(func(__e *ControlFlow) {
W7117 := __e.Get(1)
_ = W7117
tmp19916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, W7117)
}
__typedArg0 := False
__typedArg1 := W7117
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp19916 {
__e.TailApply(PrimFunc(symfail))
return
} else {
tmp19913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W7117, tmp19913)
}
__typedArg0 := W7117
__typedArg1 := tmp19913
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), tmp19914, V7116)
return


}


}, 1)

tmp19917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp19918 := Call(__e, PrimFunc(symshen_4x_4launcher_4eval_1flag_1map), tmp19917)


tmp19919 := Call(__e, tmp19912, tmp19918)


__e.TailApply(tmp19908, tmp19919)
return


} else {
__e.TailApply(PrimFunc(symthaw), Freeze7119)
return
}


}, 1)

tmp19922 := MakeNative(func(__e *ControlFlow) {
tmp19928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp19928 {
tmp19923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7115)
}
__typedArg0 := V7115
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp19924 := Call(__e, PrimFunc(symshen_4app), tmp19923, MakeString(""), symshen_4a)


tmp19926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Invalid eval argument: "))
__typedS1, __typedOK1 := TypedString(tmp19924)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Invalid eval argument: ")
__typedArg1 := tmp19924
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), Nil)
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Invalid eval argument: "))
__typedS1, __typedOK1 := TypedString(tmp19924)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Invalid eval argument: ")
__typedArg1 := tmp19924
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symerror, tmp19926)
}
__typedArg0 := symerror
__typedArg1 := tmp19926
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4x_4launcher_4eval_1command_1h)
return
}


}, 0)

__e.TailApply(tmp19907, tmp19922)
return


}


}


}


}


}


}


}, 2)

tmp19978 := Call(__e, ns2_1set, symshen_4x_4launcher_4eval_1command_1h, tmp19873)


_ = tmp19978

tmp19979 := MakeNative(func(__e *ControlFlow) {
V7120 := __e.Get(1)
_ = V7120
__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command_1h), V7120, Nil)
return
}, 1)

tmp19980 := Call(__e, ns2_1set, symshen_4x_4launcher_4eval_1command, tmp19979)


_ = tmp19980

tmp19981 := MakeNative(func(__e *ControlFlow) {
V7121 := __e.Get(1)
_ = V7121
V7122 := __e.Get(2)
_ = V7122
tmp19982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V7121, V7122)
}
__typedArg0 := V7121
__typedArg1 := V7122
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp19983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(sym_dargv_d, tmp19982)
}
__typedArg0 := sym_dargv_d
__typedArg1 := tmp19982
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp19983

tmp19984 := Call(__e, PrimFunc(symshen_4x_4launcher_4quiet_1load), V7121)


_ = tmp19984

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

tmp19985 := Call(__e, ns2_1set, symshen_4x_4launcher_4script_1command, tmp19981)


_ = tmp19985

tmp19986 := MakeNative(func(__e *ControlFlow) {
V7123 := __e.Get(1)
_ = V7123
tmp20073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20069 Obj

if True == tmp20073 {
tmp20071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp20071)
}
__typedArg0 := Nil
__typedArg1 := tmp20071
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20070 Obj

if True == tmp20072 {
ifres20070 = True


} else {
ifres20070 = False


}

ifres20069 = ifres20070


} else {
ifres20069 = False


}

if True == ifres20069 {
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
tmp20067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20058 Obj

if True == tmp20067 {
tmp20065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20065)
}
__typedArg0 := tmp20065
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20060 Obj

if True == tmp20066 {
tmp20062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20062)
}
__typedArg0 := tmp20062
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--help"), tmp20063)
}
__typedArg0 := MakeString("--help")
__typedArg1 := tmp20063
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20061 Obj

if True == tmp20064 {
ifres20061 = True


} else {
ifres20061 = False


}

ifres20060 = ifres20061


} else {
ifres20060 = False


}

var ifres20059 Obj

if True == ifres20060 {
ifres20059 = True


} else {
ifres20059 = False


}

ifres20058 = ifres20059


} else {
ifres20058 = False


}

if True == ifres20058 {
tmp19987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp19988 := Call(__e, PrimFunc(symshen_4x_4launcher_4help_1text), tmp19987)


tmp19989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19988, Nil)
}
__typedArg0 := tmp19988
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshow_1help, tmp19989)
}
__typedArg0 := symshow_1help
__typedArg1 := tmp19989
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp20056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20047 Obj

if True == tmp20056 {
tmp20054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20054)
}
__typedArg0 := tmp20054
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20049 Obj

if True == tmp20055 {
tmp20051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20051)
}
__typedArg0 := tmp20051
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("--version"), tmp20052)
}
__typedArg0 := MakeString("--version")
__typedArg1 := tmp20052
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20050 Obj

if True == tmp20053 {
ifres20050 = True


} else {
ifres20050 = False


}

ifres20049 = ifres20050


} else {
ifres20049 = False


}

var ifres20048 Obj

if True == ifres20049 {
ifres20048 = True


} else {
ifres20048 = False


}

ifres20047 = ifres20048


} else {
ifres20047 = False


}

if True == ifres20047 {
tmp19990 := Call(__e, PrimFunc(symshen_4x_4launcher_4version_1string))


tmp19991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp19990, Nil)
}
__typedArg0 := tmp19990
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsuccess, tmp19991)
}
__typedArg0 := symsuccess
__typedArg1 := tmp19991
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp20045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20036 Obj

if True == tmp20045 {
tmp20043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20043)
}
__typedArg0 := tmp20043
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20038 Obj

if True == tmp20044 {
tmp20040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20040)
}
__typedArg0 := tmp20040
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("repl"), tmp20041)
}
__typedArg0 := MakeString("repl")
__typedArg1 := tmp20041
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20039 Obj

if True == tmp20042 {
ifres20039 = True


} else {
ifres20039 = False


}

ifres20038 = ifres20039


} else {
ifres20038 = False


}

var ifres20037 Obj

if True == ifres20038 {
ifres20037 = True


} else {
ifres20037 = False


}

ifres20036 = ifres20037


} else {
ifres20036 = False


}

if True == ifres20036 {
tmp19992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp19992)
}
__typedArg0 := tmp19992
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlaunch_1repl, tmp19993)
}
__typedArg0 := symlaunch_1repl
__typedArg1 := tmp19993
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp20034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20020 Obj

if True == tmp20034 {
tmp20032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20032)
}
__typedArg0 := tmp20032
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20022 Obj

if True == tmp20033 {
tmp20029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20029)
}
__typedArg0 := tmp20029
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("script"), tmp20030)
}
__typedArg0 := MakeString("script")
__typedArg1 := tmp20030
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20024 Obj

if True == tmp20031 {
tmp20026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp20026)
}
__typedArg0 := tmp20026
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20027)
}
__typedArg0 := tmp20027
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20025 Obj

if True == tmp20028 {
ifres20025 = True


} else {
ifres20025 = False


}

ifres20024 = ifres20025


} else {
ifres20024 = False


}

var ifres20023 Obj

if True == ifres20024 {
ifres20023 = True


} else {
ifres20023 = False


}

ifres20022 = ifres20023


} else {
ifres20022 = False


}

var ifres20021 Obj

if True == ifres20022 {
ifres20021 = True


} else {
ifres20021 = False


}

ifres20020 = ifres20021


} else {
ifres20020 = False


}

if True == ifres20020 {
tmp19994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp19994)
}
__typedArg0 := tmp19994
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp19995)
}
__typedArg0 := tmp19995
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp19997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp19997)
}
__typedArg0 := tmp19997
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp19999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp19998)
}
__typedArg0 := tmp19998
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4x_4launcher_4script_1command), tmp19996, tmp19999)
return


} else {
tmp20018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20009 Obj

if True == tmp20018 {
tmp20016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20016)
}
__typedArg0 := tmp20016
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20011 Obj

if True == tmp20017 {
tmp20013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20013)
}
__typedArg0 := tmp20013
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("eval"), tmp20014)
}
__typedArg0 := MakeString("eval")
__typedArg1 := tmp20014
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20012 Obj

if True == tmp20015 {
ifres20012 = True


} else {
ifres20012 = False


}

ifres20011 = ifres20012


} else {
ifres20011 = False


}

var ifres20010 Obj

if True == ifres20011 {
ifres20010 = True


} else {
ifres20010 = False


}

ifres20009 = ifres20010


} else {
ifres20009 = False


}

if True == ifres20009 {
tmp20000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp20000)
}
__typedArg0 := tmp20000
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4x_4launcher_4eval_1command), tmp20001)
return


} else {
tmp20007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20003 Obj

if True == tmp20007 {
tmp20005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7123)
}
__typedArg0 := V7123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20005)
}
__typedArg0 := tmp20005
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20004 Obj

if True == tmp20006 {
ifres20004 = True


} else {
ifres20004 = False


}

ifres20003 = ifres20004


} else {
ifres20003 = False


}

if True == ifres20003 {
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

tmp20074 := Call(__e, ns2_1set, symshen_4x_4launcher_4launch_1shen, tmp19986)


_ = tmp20074

tmp20075 := MakeNative(func(__e *ControlFlow) {
V7126 := __e.Get(1)
_ = V7126
tmp20174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20166 Obj

if True == tmp20174 {
tmp20172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsuccess, tmp20172)
}
__typedArg0 := symsuccess
__typedArg1 := tmp20172
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20168 Obj

if True == tmp20173 {
tmp20170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp20170)
}
__typedArg0 := Nil
__typedArg1 := tmp20170
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20169 Obj

if True == tmp20171 {
ifres20169 = True


} else {
ifres20169 = False


}

ifres20168 = ifres20169


} else {
ifres20168 = False


}

var ifres20167 Obj

if True == ifres20168 {
ifres20167 = True


} else {
ifres20167 = False


}

ifres20166 = ifres20167


} else {
ifres20166 = False


}

if True == ifres20166 {
__e.Return(symshen_4x_4launcher_4done)
return
} else {
tmp20164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20151 Obj

if True == tmp20164 {
tmp20162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsuccess, tmp20162)
}
__typedArg0 := symsuccess
__typedArg1 := tmp20162
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20153 Obj

if True == tmp20163 {
tmp20160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20160)
}
__typedArg0 := tmp20160
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20155 Obj

if True == tmp20161 {
tmp20157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp20157)
}
__typedArg0 := tmp20157
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp20158)
}
__typedArg0 := Nil
__typedArg1 := tmp20158
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20156 Obj

if True == tmp20159 {
ifres20156 = True


} else {
ifres20156 = False


}

ifres20155 = ifres20156


} else {
ifres20155 = False


}

var ifres20154 Obj

if True == ifres20155 {
ifres20154 = True


} else {
ifres20154 = False


}

ifres20153 = ifres20154


} else {
ifres20153 = False


}

var ifres20152 Obj

if True == ifres20153 {
ifres20152 = True


} else {
ifres20152 = False


}

ifres20151 = ifres20152


} else {
ifres20151 = False


}

if True == ifres20151 {
tmp20076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20076)
}
__typedArg0 := tmp20076
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20078 := Call(__e, PrimFunc(symshen_4app), tmp20077, MakeString("\n"), symshen_4a)


tmp20079 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp20078, tmp20079)
return


} else {
tmp20149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20136 Obj

if True == tmp20149 {
tmp20147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symerror, tmp20147)
}
__typedArg0 := symerror
__typedArg1 := tmp20147
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20138 Obj

if True == tmp20148 {
tmp20145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20145)
}
__typedArg0 := tmp20145
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20140 Obj

if True == tmp20146 {
tmp20142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp20142)
}
__typedArg0 := tmp20142
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp20143)
}
__typedArg0 := Nil
__typedArg1 := tmp20143
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20141 Obj

if True == tmp20144 {
ifres20141 = True


} else {
ifres20141 = False


}

ifres20140 = ifres20141


} else {
ifres20140 = False


}

var ifres20139 Obj

if True == ifres20140 {
ifres20139 = True


} else {
ifres20139 = False


}

ifres20138 = ifres20139


} else {
ifres20138 = False


}

var ifres20137 Obj

if True == ifres20138 {
ifres20137 = True


} else {
ifres20137 = False


}

ifres20136 = ifres20137


} else {
ifres20136 = False


}

if True == ifres20136 {
tmp20080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20080)
}
__typedArg0 := tmp20080
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20082 := Call(__e, PrimFunc(symshen_4app), tmp20081, MakeString("\n"), symshen_4a)


tmp20083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("ERROR: "))
__typedS1, __typedOK1 := TypedString(tmp20082)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("ERROR: ")
__typedArg1 := tmp20082
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp20084 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp20083, tmp20084)
return


} else {
tmp20134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20130 Obj

if True == tmp20134 {
tmp20132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlaunch_1repl, tmp20132)
}
__typedArg0 := symlaunch_1repl
__typedArg1 := tmp20132
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20131 Obj

if True == tmp20133 {
ifres20131 = True


} else {
ifres20131 = False


}

ifres20130 = ifres20131


} else {
ifres20130 = False


}

if True == ifres20130 {
__e.TailApply(PrimFunc(symshen_4repl))
return
} else {
tmp20128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20115 Obj

if True == tmp20128 {
tmp20126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshow_1help, tmp20126)
}
__typedArg0 := symshow_1help
__typedArg1 := tmp20126
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20117 Obj

if True == tmp20127 {
tmp20124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20124)
}
__typedArg0 := tmp20124
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20119 Obj

if True == tmp20125 {
tmp20121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp20121)
}
__typedArg0 := tmp20121
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp20122)
}
__typedArg0 := Nil
__typedArg1 := tmp20122
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20120 Obj

if True == tmp20123 {
ifres20120 = True


} else {
ifres20120 = False


}

ifres20119 = ifres20120


} else {
ifres20119 = False


}

var ifres20118 Obj

if True == ifres20119 {
ifres20118 = True


} else {
ifres20118 = False


}

ifres20117 = ifres20118


} else {
ifres20117 = False


}

var ifres20116 Obj

if True == ifres20117 {
ifres20116 = True


} else {
ifres20116 = False


}

ifres20115 = ifres20116


} else {
ifres20115 = False


}

if True == ifres20115 {
tmp20085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20085)
}
__typedArg0 := tmp20085
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20087 := Call(__e, PrimFunc(symshen_4app), tmp20086, MakeString("\n"), symshen_4a)


tmp20088 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp20087, tmp20088)
return


} else {
tmp20113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20100 Obj

if True == tmp20113 {
tmp20111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symunknown_1arguments, tmp20111)
}
__typedArg0 := symunknown_1arguments
__typedArg1 := tmp20111
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres20102 Obj

if True == tmp20112 {
tmp20109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20109)
}
__typedArg0 := tmp20109
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20104 Obj

if True == tmp20110 {
tmp20106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp20106)
}
__typedArg0 := tmp20106
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp20107)
}
__typedArg0 := tmp20107
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres20105 Obj

if True == tmp20108 {
ifres20105 = True


} else {
ifres20105 = False


}

ifres20104 = ifres20105


} else {
ifres20104 = False


}

var ifres20103 Obj

if True == ifres20104 {
ifres20103 = True


} else {
ifres20103 = False


}

ifres20102 = ifres20103


} else {
ifres20102 = False


}

var ifres20101 Obj

if True == ifres20102 {
ifres20101 = True


} else {
ifres20101 = False


}

ifres20100 = ifres20101


} else {
ifres20100 = False


}

if True == ifres20100 {
tmp20089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp20089)
}
__typedArg0 := tmp20089
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20090)
}
__typedArg0 := tmp20090
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7126)
}
__typedArg0 := V7126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp20093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp20092)
}
__typedArg0 := tmp20092
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp20094 := Call(__e, PrimFunc(symshen_4app), tmp20093, MakeString(" --help' for more information.\n"), symshen_4a)


tmp20096 := Call(__e, PrimFunc(symshen_4app), tmp20091, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\nTry `"))
__typedS1, __typedOK1 := TypedString(tmp20094)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\nTry `")
__typedArg1 := tmp20094
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp20097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("ERROR: Invalid argument: "))
__typedS1, __typedOK1 := TypedString(tmp20096)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("ERROR: Invalid argument: ")
__typedArg1 := tmp20096
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp20098 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp20097, tmp20098)
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

tmp20175 := Call(__e, ns2_1set, symshen_4x_4launcher_4default_1handle_1result, tmp20075)


_ = tmp20175

tmp20176 := MakeNative(func(__e *ControlFlow) {
V7127 := __e.Get(1)
_ = V7127
tmp20177 := Call(__e, PrimFunc(symshen_4x_4launcher_4launch_1shen), V7127)


__e.TailApply(PrimFunc(symshen_4x_4launcher_4default_1handle_1result), tmp20177)
return


}, 1)

__e.TailApply(ns2_1set, symshen_4x_4launcher_4main, tmp20176)
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
var sym_4_4_4 = MakeSymbol("...")
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
