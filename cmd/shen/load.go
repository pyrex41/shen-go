package main

import . "github.com/pyrex41/shen-go/kl"

var LoadMain = MakeNative(func(__e *ControlFlow) {
tmp9936 := MakeNative(func(__e *ControlFlow) {
V907 := __e.Get(1)
_ = V907
tmp9937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtc_d)
}
__typedArg0 := symshen_4_dtc_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__9918 := tmp9937
_ = let__9918

tmp9938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symget_1time) {
return PrimGetTime(symrun)
}
__typedArg0 := symrun
return Call(__e, PrimFunc(symget_1time), __typedArg0)
})()

let__9920 := tmp9938
_ = let__9920

tmp9939 := Call(__e, PrimFunc(symread_1file), V907)


tmp9940 := Call(__e, PrimFunc(symshen_4load_1help), let__9918, tmp9939)


let__9921 := tmp9940
_ = let__9921

tmp9941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symget_1time) {
return PrimGetTime(symrun)
}
__typedArg0 := symrun
return Call(__e, PrimFunc(symget_1time), __typedArg0)
})()

let__9922 := tmp9941
_ = let__9922

let__9923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(let__9922)
__typedN1, __typedOK1 := TypedFloat64(let__9920)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := let__9922
__typedArg1 := let__9920
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})()
_ = let__9923

tmp9943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(let__9923)
}
__typedArg0 := let__9923
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp9945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\nrun time: "))
__typedS1, __typedOK1 := TypedString(tmp9943)
__typedS2, __typedOK2 := TypedString(MakeString(" secs\n"))
if __typedOK0 && __typedOK1 && __typedOK2 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + (__typedS1 + __typedS2)))
}}
__typedArg0 := MakeString("\nrun time: ")
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp9943)
__typedS1, __typedOK1 := TypedString(MakeString(" secs\n"))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp9943
__typedArg1 := MakeString(" secs\n")
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp9946 := Call(__e, PrimFunc(symstoutput))


tmp9947 := Call(__e, PrimFunc(sympr), tmp9945, tmp9946)


let__9924 := tmp9947
_ = let__9924

let__9919 := let__9921
_ = let__9919

var ifres9948 Obj

if True == let__9918 {
tmp9949 := Call(__e, PrimFunc(syminferences))


tmp9950 := Call(__e, PrimFunc(symshen_4app), tmp9949, MakeString(" inferences\n"), symshen_4a)


tmp9951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\ntypechecked in "))
__typedS1, __typedOK1 := TypedString(tmp9950)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\ntypechecked in ")
__typedArg1 := tmp9950
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp9952 := Call(__e, PrimFunc(symstoutput))


tmp9953 := Call(__e, PrimFunc(sympr), tmp9951, tmp9952)


ifres9948 = tmp9953


} else {
ifres9948 = symshen_4skip


}

let__9925 := ifres9948
_ = let__9925

__e.Return(symloaded)
return


}, 1)

tmp9954 := Call(__e, ns2_1set, symload, tmp9936)


_ = tmp9954

tmp9955 := MakeNative(func(__e *ControlFlow) {
V918 := __e.Get(1)
_ = V918
V919 := __e.Get(2)
_ = V919
tmp9957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, V918)
}
__typedArg0 := False
__typedArg1 := V918
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9957 {
__e.TailApply(PrimFunc(symshen_4eval_1and_1print), V919)
return
} else {
__e.TailApply(PrimFunc(symshen_4check_1eval_1and_1print), V919)
return
}


}, 2)

tmp9958 := Call(__e, ns2_1set, symshen_4load_1help, tmp9955)


_ = tmp9958

tmp9959 := MakeNative(func(__e *ControlFlow) {
V920 := __e.Get(1)
_ = V920
tmp9960 := MakeNative(func(__e *ControlFlow) {
Z921 := __e.Get(1)
_ = Z921
tmp9961 := Call(__e, PrimFunc(symshen_4shen_1_6kl), Z921)


tmp9962 := Call(__e, PrimFunc(symeval_1kl), tmp9961)


tmp9963 := Call(__e, PrimFunc(symshen_4app), tmp9962, MakeString("\n"), symshen_4s)


tmp9964 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp9963, tmp9964)
return


}, 1)

__e.TailApply(PrimFunc(symmap), tmp9960, V920)
return


}, 1)

tmp9965 := Call(__e, ns2_1set, symshen_4eval_1and_1print, tmp9959)


_ = tmp9965

tmp9966 := MakeNative(func(__e *ControlFlow) {
V922 := __e.Get(1)
_ = V922
tmp9967 := MakeNative(func(__e *ControlFlow) {
Z924 := __e.Get(1)
_ = Z924
__e.TailApply(PrimFunc(symshen_4typetable), Z924)
return
}, 1)

tmp9968 := Call(__e, PrimFunc(symmapcan), tmp9967, V922)


let__9926 := tmp9968
_ = let__9926

tmp9969 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4assumetypes), let__9926)
return
}, 0)

tmp9970 := MakeNative(func(__e *ControlFlow) {
Z926 := __e.Get(1)
_ = Z926
__e.TailApply(PrimFunc(symshen_4unwind_1types), Z926, let__9926)
return
}, 1)

tmp9971 := Call(__e, try_1catch, tmp9969, tmp9970)


let__9927 := tmp9971
_ = let__9927

tmp9972 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4work_1through), V922)
return
}, 0)

tmp9973 := MakeNative(func(__e *ControlFlow) {
Z927 := __e.Get(1)
_ = Z927
__e.TailApply(PrimFunc(symshen_4unwind_1types), Z927, let__9926)
return
}, 1)

__e.TailApply(try_1catch, tmp9972, tmp9973)
return


}, 1)

tmp9974 := Call(__e, ns2_1set, symshen_4check_1eval_1and_1print, tmp9966)


_ = tmp9974

tmp9975 := MakeNative(func(__e *ControlFlow) {
V932 := __e.Get(1)
_ = V932
tmp10020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V932)
}
__typedArg0 := V932
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10001 Obj

if True == tmp10020 {
tmp10018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V932)
}
__typedArg0 := V932
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp10018)
}
__typedArg0 := symdefine
__typedArg1 := tmp10018
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10003 Obj

if True == tmp10019 {
tmp10016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V932)
}
__typedArg0 := V932
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10016)
}
__typedArg0 := tmp10016
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10005 Obj

if True == tmp10017 {
tmp10013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V932)
}
__typedArg0 := V932
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10013)
}
__typedArg0 := tmp10013
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10014)
}
__typedArg0 := tmp10014
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10007 Obj

if True == tmp10015 {
tmp10009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V932)
}
__typedArg0 := V932
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10009)
}
__typedArg0 := tmp10009
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10010)
}
__typedArg0 := tmp10010
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, tmp10011)
}
__typedArg0 := sym_i
__typedArg1 := tmp10011
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10008 Obj

if True == tmp10012 {
ifres10008 = True


} else {
ifres10008 = False


}

ifres10007 = ifres10008


} else {
ifres10007 = False


}

var ifres10006 Obj

if True == ifres10007 {
ifres10006 = True


} else {
ifres10006 = False


}

ifres10005 = ifres10006


} else {
ifres10005 = False


}

var ifres10004 Obj

if True == ifres10005 {
ifres10004 = True


} else {
ifres10004 = False


}

ifres10003 = ifres10004


} else {
ifres10003 = False


}

var ifres10002 Obj

if True == ifres10003 {
ifres10002 = True


} else {
ifres10002 = False


}

ifres10001 = ifres10002


} else {
ifres10001 = False


}

if True == ifres10001 {
tmp9976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V932)
}
__typedArg0 := V932
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9976)
}
__typedArg0 := tmp9976
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V932)
}
__typedArg0 := V932
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9978)
}
__typedArg0 := tmp9978
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V932)
}
__typedArg0 := V932
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9980)
}
__typedArg0 := tmp9980
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9981)
}
__typedArg0 := tmp9981
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9983 := Call(__e, PrimFunc(symshen_4type_1F), tmp9979, tmp9982)


tmp9984 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp9983)


tmp9985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9984, Nil)
}
__typedArg0 := tmp9984
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9977, tmp9985)
}
__typedArg0 := tmp9977
__typedArg1 := tmp9985
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V932)
}
__typedArg0 := V932
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9991 Obj

if True == tmp9999 {
tmp9997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V932)
}
__typedArg0 := V932
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp9997)
}
__typedArg0 := symdefine
__typedArg1 := tmp9997
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9993 Obj

if True == tmp9998 {
tmp9995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V932)
}
__typedArg0 := V932
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9995)
}
__typedArg0 := tmp9995
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9994 Obj

if True == tmp9996 {
ifres9994 = True


} else {
ifres9994 = False


}

ifres9993 = ifres9994


} else {
ifres9993 = False


}

var ifres9992 Obj

if True == ifres9993 {
ifres9992 = True


} else {
ifres9992 = False


}

ifres9991 = ifres9992


} else {
ifres9991 = False


}

if True == ifres9991 {
tmp9986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V932)
}
__typedArg0 := V932
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9986)
}
__typedArg0 := tmp9986
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9988 := Call(__e, PrimFunc(symshen_4app), tmp9987, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing { in "))
__typedS1, __typedOK1 := TypedString(tmp9988)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing { in ")
__typedArg1 := tmp9988
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing { in "))
__typedS1, __typedOK1 := TypedString(tmp9988)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing { in ")
__typedArg1 := tmp9988
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
__e.Return(Nil)
return
}


}


}, 1)

tmp10021 := Call(__e, ns2_1set, symshen_4typetable, tmp9975)


_ = tmp10021

tmp10022 := MakeNative(func(__e *ControlFlow) {
V939 := __e.Get(1)
_ = V939
V940 := __e.Get(2)
_ = V940
tmp10035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V940)
}
__typedArg0 := V940
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10031 Obj

if True == tmp10035 {
tmp10033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V940)
}
__typedArg0 := V940
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, tmp10033)
}
__typedArg0 := sym_j
__typedArg1 := tmp10033
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10032 Obj

if True == tmp10034 {
ifres10032 = True


} else {
ifres10032 = False


}

ifres10031 = ifres10032


} else {
ifres10031 = False


}

if True == ifres10031 {
__e.Return(Nil)
return
} else {
tmp10029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V940)
}
__typedArg0 := V940
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10029 {
tmp10023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V940)
}
__typedArg0 := V940
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V940)
}
__typedArg0 := V940
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10025 := Call(__e, PrimFunc(symshen_4type_1F), V939, tmp10024)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10023, tmp10025)
}
__typedArg0 := tmp10023
__typedArg1 := tmp10025
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp10026 := Call(__e, PrimFunc(symshen_4app), V939, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing } in "))
__typedS1, __typedOK1 := TypedString(tmp10026)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing } in ")
__typedArg1 := tmp10026
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing } in "))
__typedS1, __typedOK1 := TypedString(tmp10026)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing } in ")
__typedArg1 := tmp10026
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}


}, 2)

tmp10036 := Call(__e, ns2_1set, symshen_4type_1F, tmp10022)


_ = tmp10036

tmp10037 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V943 := __self1
_ = V943
tmp10051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V943)
}
__typedArg0 := Nil
__typedArg1 := V943
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10051 {
__e.Return(Nil)
return
} else {
tmp10049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V943)
}
__typedArg0 := V943
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10045 Obj

if True == tmp10049 {
tmp10047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V943)
}
__typedArg0 := V943
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10047)
}
__typedArg0 := tmp10047
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10046 Obj

if True == tmp10048 {
ifres10046 = True


} else {
ifres10046 = False


}

ifres10045 = ifres10046


} else {
ifres10045 = False


}

if True == ifres10045 {
tmp10038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V943)
}
__typedArg0 := V943
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V943)
}
__typedArg0 := V943
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10039)
}
__typedArg0 := tmp10039
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10041 := Call(__e, PrimFunc(symdeclare), tmp10038, tmp10040)


_ = tmp10041

tmp10042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V943)
}
__typedArg0 := V943
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10042)
}
__typedArg0 := tmp10042
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4assumetypes) == __self {
__self1 = tmp10043
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4assumetypes), tmp10043)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.assumetype"))
}
__typedArg0 := MakeString("implementation error in shen.assumetype")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp10052 := Call(__e, ns2_1set, symshen_4assumetypes, tmp10037)


_ = tmp10052

tmp10053 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V948 := __self1
_ = V948
V949 := __self2
_ = V949
tmp10064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V949)
}
__typedArg0 := V949
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10060 Obj

if True == tmp10064 {
tmp10062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V949)
}
__typedArg0 := V949
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10062)
}
__typedArg0 := tmp10062
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10061 Obj

if True == tmp10063 {
ifres10061 = True


} else {
ifres10061 = False


}

ifres10060 = ifres10061


} else {
ifres10060 = False


}

if True == ifres10060 {
tmp10054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V949)
}
__typedArg0 := V949
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10055 := Call(__e, PrimFunc(symdestroy), tmp10054)


_ = tmp10055

tmp10056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V949)
}
__typedArg0 := V949
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10056)
}
__typedArg0 := tmp10056
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4unwind_1types) == __self {
__self1, __self2 = V948, tmp10057
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4unwind_1types), V948, tmp10057)
return


} else {
tmp10058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symerror_1to_1string) {
return PrimErrorToString(V948)
}
__typedArg0 := V948
return Call(__e, PrimFunc(symerror_1to_1string), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp10058)
}
__typedArg0 := tmp10058
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 2)

tmp10065 := Call(__e, ns2_1set, symshen_4unwind_1types, tmp10053)


_ = tmp10065

tmp10066 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V952 := __self1
_ = V952
tmp10112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V952)
}
__typedArg0 := Nil
__typedArg1 := V952
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10112 {
__e.Return(Nil)
return
} else {
tmp10110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V952)
}
__typedArg0 := V952
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10095 Obj

if True == tmp10110 {
tmp10108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V952)
}
__typedArg0 := V952
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10108)
}
__typedArg0 := tmp10108
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10097 Obj

if True == tmp10109 {
tmp10105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V952)
}
__typedArg0 := V952
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10105)
}
__typedArg0 := tmp10105
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10106)
}
__typedArg0 := tmp10106
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10099 Obj

if True == tmp10107 {
tmp10101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V952)
}
__typedArg0 := V952
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10101)
}
__typedArg0 := tmp10101
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp10104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp10102, tmp10103)
}
__typedArg0 := tmp10102
__typedArg1 := tmp10103
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10100 Obj

if True == tmp10104 {
ifres10100 = True


} else {
ifres10100 = False


}

ifres10099 = ifres10100


} else {
ifres10099 = False


}

var ifres10098 Obj

if True == ifres10099 {
ifres10098 = True


} else {
ifres10098 = False


}

ifres10097 = ifres10098


} else {
ifres10097 = False


}

var ifres10096 Obj

if True == ifres10097 {
ifres10096 = True


} else {
ifres10096 = False


}

ifres10095 = ifres10096


} else {
ifres10095 = False


}

if True == ifres10095 {
tmp10067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V952)
}
__typedArg0 := V952
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V952)
}
__typedArg0 := V952
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10068)
}
__typedArg0 := tmp10068
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10069)
}
__typedArg0 := tmp10069
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10071 := Call(__e, PrimFunc(symshen_4typecheck), tmp10067, tmp10070)


let__9928 := tmp10071
_ = let__9928

tmp10085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__9928, False)
}
__typedArg0 := let__9928
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10085 {
__e.TailApply(PrimFunc(symshen_4type_1error))
return
} else {
tmp10072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V952)
}
__typedArg0 := V952
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10073 := Call(__e, PrimFunc(symshen_4shen_1_6kl), tmp10072)


tmp10074 := Call(__e, PrimFunc(symeval_1kl), tmp10073)


let__9929 := tmp10074
_ = let__9929

tmp10075 := Call(__e, PrimFunc(symshen_4pretty_1type), let__9928)


tmp10076 := Call(__e, PrimFunc(symshen_4app), tmp10075, MakeString("\n"), symshen_4r)


tmp10078 := Call(__e, PrimFunc(symshen_4app), let__9929, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" : "))
__typedS1, __typedOK1 := TypedString(tmp10076)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" : ")
__typedArg1 := tmp10076
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4s)


tmp10079 := Call(__e, PrimFunc(symstoutput))


tmp10080 := Call(__e, PrimFunc(sympr), tmp10078, tmp10079)


let__9930 := tmp10080
_ = let__9930

tmp10081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V952)
}
__typedArg0 := V952
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10081)
}
__typedArg0 := tmp10081
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10082)
}
__typedArg0 := tmp10082
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4work_1through) == __self {
__self1 = tmp10083
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4work_1through), tmp10083)
return


}


} else {
tmp10093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V952)
}
__typedArg0 := V952
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10093 {
tmp10086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V952)
}
__typedArg0 := V952
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp10088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V952)
}
__typedArg0 := V952
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp10088)
}
__typedArg0 := symA
__typedArg1 := tmp10088
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10087, tmp10089)
}
__typedArg0 := tmp10087
__typedArg1 := tmp10089
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10086, tmp10090)
}
__typedArg0 := tmp10086
__typedArg1 := tmp10090
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4work_1through) == __self {
__self1 = tmp10091
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4work_1through), tmp10091)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.work-through"))
}
__typedArg0 := MakeString("implementation error in shen.work-through")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 1)

tmp10113 := Call(__e, ns2_1set, symshen_4work_1through, tmp10066)


_ = tmp10113

tmp10114 := MakeNative(func(__e *ControlFlow) {
V957 := __e.Get(1)
_ = V957
tmp10256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10130 Obj

if True == tmp10256 {
tmp10254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10254)
}
__typedArg0 := tmp10254
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10132 Obj

if True == tmp10255 {
tmp10251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10251)
}
__typedArg0 := tmp10251
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp10252)
}
__typedArg0 := symlist
__typedArg1 := tmp10252
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10134 Obj

if True == tmp10253 {
tmp10248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10248)
}
__typedArg0 := tmp10248
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10249)
}
__typedArg0 := tmp10249
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10136 Obj

if True == tmp10250 {
tmp10244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10244)
}
__typedArg0 := tmp10244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10245)
}
__typedArg0 := tmp10245
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10246)
}
__typedArg0 := Nil
__typedArg1 := tmp10246
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10138 Obj

if True == tmp10247 {
tmp10242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10242)
}
__typedArg0 := tmp10242
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10140 Obj

if True == tmp10243 {
tmp10239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
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
return PrimEqual(sym_1_1_6, tmp10240)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp10240
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10142 Obj

if True == tmp10241 {
tmp10236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10236)
}
__typedArg0 := tmp10236
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10237)
}
__typedArg0 := tmp10237
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10144 Obj

if True == tmp10238 {
tmp10232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10232)
}
__typedArg0 := tmp10232
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10233)
}
__typedArg0 := tmp10233
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10234)
}
__typedArg0 := tmp10234
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10146 Obj

if True == tmp10235 {
tmp10227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10227)
}
__typedArg0 := tmp10227
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10228)
}
__typedArg0 := tmp10228
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10229)
}
__typedArg0 := tmp10229
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symstr, tmp10230)
}
__typedArg0 := symstr
__typedArg1 := tmp10230
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10148 Obj

if True == tmp10231 {
tmp10222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10222)
}
__typedArg0 := tmp10222
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10223)
}
__typedArg0 := tmp10223
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10224)
}
__typedArg0 := tmp10224
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10225)
}
__typedArg0 := tmp10225
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10150 Obj

if True == tmp10226 {
tmp10216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10216)
}
__typedArg0 := tmp10216
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10217)
}
__typedArg0 := tmp10217
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10218)
}
__typedArg0 := tmp10218
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10219)
}
__typedArg0 := tmp10219
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10220)
}
__typedArg0 := tmp10220
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10152 Obj

if True == tmp10221 {
tmp10209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10209)
}
__typedArg0 := tmp10209
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10210)
}
__typedArg0 := tmp10210
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10211)
}
__typedArg0 := tmp10211
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10212)
}
__typedArg0 := tmp10212
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10213)
}
__typedArg0 := tmp10213
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp10214)
}
__typedArg0 := symlist
__typedArg1 := tmp10214
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10154 Obj

if True == tmp10215 {
tmp10202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10202)
}
__typedArg0 := tmp10202
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10203)
}
__typedArg0 := tmp10203
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10204)
}
__typedArg0 := tmp10204
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10205)
}
__typedArg0 := tmp10205
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10206)
}
__typedArg0 := tmp10206
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10207)
}
__typedArg0 := tmp10207
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10156 Obj

if True == tmp10208 {
tmp10194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10194)
}
__typedArg0 := tmp10194
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10195)
}
__typedArg0 := tmp10195
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10196)
}
__typedArg0 := tmp10196
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10197)
}
__typedArg0 := tmp10197
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10198)
}
__typedArg0 := tmp10198
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10199)
}
__typedArg0 := tmp10199
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10200)
}
__typedArg0 := Nil
__typedArg1 := tmp10200
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10158 Obj

if True == tmp10201 {
tmp10188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10188)
}
__typedArg0 := tmp10188
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10189)
}
__typedArg0 := tmp10189
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10190)
}
__typedArg0 := tmp10190
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10191)
}
__typedArg0 := tmp10191
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10192)
}
__typedArg0 := tmp10192
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10160 Obj

if True == tmp10193 {
tmp10181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10181)
}
__typedArg0 := tmp10181
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10182)
}
__typedArg0 := tmp10182
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10183)
}
__typedArg0 := tmp10183
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10184)
}
__typedArg0 := tmp10184
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10185)
}
__typedArg0 := tmp10185
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10186)
}
__typedArg0 := Nil
__typedArg1 := tmp10186
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10162 Obj

if True == tmp10187 {
tmp10177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10177)
}
__typedArg0 := tmp10177
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10178)
}
__typedArg0 := tmp10178
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10179)
}
__typedArg0 := Nil
__typedArg1 := tmp10179
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10164 Obj

if True == tmp10180 {
tmp10166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10166)
}
__typedArg0 := tmp10166
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10167)
}
__typedArg0 := tmp10167
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10169)
}
__typedArg0 := tmp10169
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10170)
}
__typedArg0 := tmp10170
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10171)
}
__typedArg0 := tmp10171
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10172)
}
__typedArg0 := tmp10172
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10173)
}
__typedArg0 := tmp10173
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10174)
}
__typedArg0 := tmp10174
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp10168, tmp10175)
}
__typedArg0 := tmp10168
__typedArg1 := tmp10175
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10165 Obj

if True == tmp10176 {
ifres10165 = True


} else {
ifres10165 = False


}

ifres10164 = ifres10165


} else {
ifres10164 = False


}

var ifres10163 Obj

if True == ifres10164 {
ifres10163 = True


} else {
ifres10163 = False


}

ifres10162 = ifres10163


} else {
ifres10162 = False


}

var ifres10161 Obj

if True == ifres10162 {
ifres10161 = True


} else {
ifres10161 = False


}

ifres10160 = ifres10161


} else {
ifres10160 = False


}

var ifres10159 Obj

if True == ifres10160 {
ifres10159 = True


} else {
ifres10159 = False


}

ifres10158 = ifres10159


} else {
ifres10158 = False


}

var ifres10157 Obj

if True == ifres10158 {
ifres10157 = True


} else {
ifres10157 = False


}

ifres10156 = ifres10157


} else {
ifres10156 = False


}

var ifres10155 Obj

if True == ifres10156 {
ifres10155 = True


} else {
ifres10155 = False


}

ifres10154 = ifres10155


} else {
ifres10154 = False


}

var ifres10153 Obj

if True == ifres10154 {
ifres10153 = True


} else {
ifres10153 = False


}

ifres10152 = ifres10153


} else {
ifres10152 = False


}

var ifres10151 Obj

if True == ifres10152 {
ifres10151 = True


} else {
ifres10151 = False


}

ifres10150 = ifres10151


} else {
ifres10150 = False


}

var ifres10149 Obj

if True == ifres10150 {
ifres10149 = True


} else {
ifres10149 = False


}

ifres10148 = ifres10149


} else {
ifres10148 = False


}

var ifres10147 Obj

if True == ifres10148 {
ifres10147 = True


} else {
ifres10147 = False


}

ifres10146 = ifres10147


} else {
ifres10146 = False


}

var ifres10145 Obj

if True == ifres10146 {
ifres10145 = True


} else {
ifres10145 = False


}

ifres10144 = ifres10145


} else {
ifres10144 = False


}

var ifres10143 Obj

if True == ifres10144 {
ifres10143 = True


} else {
ifres10143 = False


}

ifres10142 = ifres10143


} else {
ifres10142 = False


}

var ifres10141 Obj

if True == ifres10142 {
ifres10141 = True


} else {
ifres10141 = False


}

ifres10140 = ifres10141


} else {
ifres10140 = False


}

var ifres10139 Obj

if True == ifres10140 {
ifres10139 = True


} else {
ifres10139 = False


}

ifres10138 = ifres10139


} else {
ifres10138 = False


}

var ifres10137 Obj

if True == ifres10138 {
ifres10137 = True


} else {
ifres10137 = False


}

ifres10136 = ifres10137


} else {
ifres10136 = False


}

var ifres10135 Obj

if True == ifres10136 {
ifres10135 = True


} else {
ifres10135 = False


}

ifres10134 = ifres10135


} else {
ifres10134 = False


}

var ifres10133 Obj

if True == ifres10134 {
ifres10133 = True


} else {
ifres10133 = False


}

ifres10132 = ifres10133


} else {
ifres10132 = False


}

var ifres10131 Obj

if True == ifres10132 {
ifres10131 = True


} else {
ifres10131 = False


}

ifres10130 = ifres10131


} else {
ifres10130 = False


}

if True == ifres10130 {
tmp10115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10115)
}
__typedArg0 := tmp10115
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10116)
}
__typedArg0 := tmp10116
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10117)
}
__typedArg0 := tmp10117
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10118)
}
__typedArg0 := tmp10118
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10120)
}
__typedArg0 := tmp10120
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10121)
}
__typedArg0 := tmp10121
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10122)
}
__typedArg0 := tmp10122
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10123)
}
__typedArg0 := tmp10123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a_6, tmp10124)
}
__typedArg0 := sym_a_a_6
__typedArg1 := tmp10124
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10119, tmp10125)
}
__typedArg0 := tmp10119
__typedArg1 := tmp10125
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp10128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V957)
}
__typedArg0 := V957
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10128 {
tmp10126 := MakeNative(func(__e *ControlFlow) {
Z958 := __e.Get(1)
_ = Z958
__e.TailApply(PrimFunc(symshen_4pretty_1type), Z958)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp10126, V957)
return


} else {
__e.Return(V957)
return
}


}


}, 1)

tmp10257 := Call(__e, ns2_1set, symshen_4pretty_1type, tmp10114)


_ = tmp10257

tmp10258 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("type error\n"))
}
__typedArg0 := MakeString("type error\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}, 0)

tmp10259 := Call(__e, ns2_1set, symshen_4type_1error, tmp10258)


_ = tmp10259

tmp10260 := MakeNative(func(__e *ControlFlow) {
V959 := __e.Get(1)
_ = V959
tmp10261 := Call(__e, PrimFunc(symshen_4klfile), V959)


let__9931 := tmp10261
_ = let__9931

tmp10262 := Call(__e, PrimFunc(symread_1file), V959)


let__9932 := tmp10262
_ = let__9932

tmp10263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symopen) {
return PrimOpenStream(let__9931, symout)
}
__typedArg0 := let__9931
__typedArg1 := symout
return Call(__e, PrimFunc(symopen), __typedArg0, __typedArg1)
})()

let__9933 := tmp10263
_ = let__9933

tmp10264 := MakeNative(func(__e *ControlFlow) {
Z964 := __e.Get(1)
_ = Z964
tmp10265 := Call(__e, PrimFunc(symshen_4shen_1_6kl_1h), Z964)


__e.TailApply(PrimFunc(symshen_4partial), tmp10265)
return


}, 1)

tmp10266 := Call(__e, PrimFunc(symmap), tmp10264, let__9932)


let__9934 := tmp10266
_ = let__9934

tmp10267 := Call(__e, PrimFunc(symshen_4write_1kl), let__9934, let__9933)


let__9935 := tmp10267
_ = let__9935

__e.Return(let__9931)
return


}, 1)

tmp10268 := Call(__e, ns2_1set, symbootstrap, tmp10260)


_ = tmp10268

tmp10269 := MakeNative(func(__e *ControlFlow) {
V966 := __e.Get(1)
_ = V966
tmp10292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V966)
}
__typedArg0 := V966
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10279 Obj

if True == tmp10292 {
tmp10290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V966)
}
__typedArg0 := V966
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4f_1error, tmp10290)
}
__typedArg0 := symshen_4f_1error
__typedArg1 := tmp10290
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10281 Obj

if True == tmp10291 {
tmp10288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V966)
}
__typedArg0 := V966
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10288)
}
__typedArg0 := tmp10288
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10283 Obj

if True == tmp10289 {
tmp10285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V966)
}
__typedArg0 := V966
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10285)
}
__typedArg0 := tmp10285
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10286)
}
__typedArg0 := Nil
__typedArg1 := tmp10286
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10284 Obj

if True == tmp10287 {
ifres10284 = True


} else {
ifres10284 = False


}

ifres10283 = ifres10284


} else {
ifres10283 = False


}

var ifres10282 Obj

if True == ifres10283 {
ifres10282 = True


} else {
ifres10282 = False


}

ifres10281 = ifres10282


} else {
ifres10281 = False


}

var ifres10280 Obj

if True == ifres10281 {
ifres10280 = True


} else {
ifres10280 = False


}

ifres10279 = ifres10280


} else {
ifres10279 = False


}

if True == ifres10279 {
tmp10270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V966)
}
__typedArg0 := V966
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10270)
}
__typedArg0 := tmp10270
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(tmp10271)
}
__typedArg0 := tmp10271
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp10274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("partial function "))
__typedS1, __typedOK1 := TypedString(tmp10272)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("partial function ")
__typedArg1 := tmp10272
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), Nil)
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("partial function "))
__typedS1, __typedOK1 := TypedString(tmp10272)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("partial function ")
__typedArg1 := tmp10272
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp10274)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp10274
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp10277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V966)
}
__typedArg0 := V966
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10277 {
tmp10275 := MakeNative(func(__e *ControlFlow) {
Z967 := __e.Get(1)
_ = Z967
__e.TailApply(PrimFunc(symshen_4partial), Z967)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp10275, V966)
return


} else {
__e.Return(V966)
return
}


}


}, 1)

tmp10293 := Call(__e, ns2_1set, symshen_4partial, tmp10269)


_ = tmp10293

tmp10294 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V970 := __self1
_ = V970
V971 := __self2
_ = V971
tmp10308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V970)
}
__typedArg0 := Nil
__typedArg1 := V970
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10308 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symclose) {
return PrimCloseStream(V971)
}
__typedArg0 := V971
return Call(__e, PrimFunc(symclose), __typedArg0)
})())
return
} else {
tmp10306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V970)
}
__typedArg0 := V970
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10302 Obj

if True == tmp10306 {
tmp10304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V970)
}
__typedArg0 := V970
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10304)
}
__typedArg0 := tmp10304
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10303 Obj

if True == tmp10305 {
ifres10303 = True


} else {
ifres10303 = False


}

ifres10302 = ifres10303


} else {
ifres10302 = False


}

if True == ifres10302 {
tmp10295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V970)
}
__typedArg0 := V970
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V970)
}
__typedArg0 := V970
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10297 := Call(__e, PrimFunc(symshen_4write_1kl_1h), tmp10296, V971)


_ = tmp10297

if PrimFunc(symshen_4write_1kl) == __self {
__self1, __self2 = tmp10295, V971
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4write_1kl), tmp10295, V971)
return


} else {
tmp10300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V970)
}
__typedArg0 := V970
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10300 {
tmp10298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V970)
}
__typedArg0 := V970
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4write_1kl) == __self {
__self1, __self2 = tmp10298, V971
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4write_1kl), tmp10298, V971)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.write-kl"))
}
__typedArg0 := MakeString("partial function shen.write-kl")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 2)

tmp10309 := Call(__e, ns2_1set, symshen_4write_1kl, tmp10294)


_ = tmp10309

tmp10310 := MakeNative(func(__e *ControlFlow) {
V974 := __e.Get(1)
_ = V974
V975 := __e.Get(2)
_ = V975
tmp10350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V974)
}
__typedArg0 := V974
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10313 Obj

if True == tmp10350 {
tmp10348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V974)
}
__typedArg0 := V974
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp10348)
}
__typedArg0 := symdefun
__typedArg1 := tmp10348
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10315 Obj

if True == tmp10349 {
tmp10346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V974)
}
__typedArg0 := V974
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10346)
}
__typedArg0 := tmp10346
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10317 Obj

if True == tmp10347 {
tmp10343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V974)
}
__typedArg0 := V974
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10343)
}
__typedArg0 := tmp10343
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail, tmp10344)
}
__typedArg0 := symfail
__typedArg1 := tmp10344
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10319 Obj

if True == tmp10345 {
tmp10340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V974)
}
__typedArg0 := V974
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10340)
}
__typedArg0 := tmp10340
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10341)
}
__typedArg0 := tmp10341
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10321 Obj

if True == tmp10342 {
tmp10336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V974)
}
__typedArg0 := V974
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10336)
}
__typedArg0 := tmp10336
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10337)
}
__typedArg0 := tmp10337
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10338)
}
__typedArg0 := Nil
__typedArg1 := tmp10338
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10323 Obj

if True == tmp10339 {
tmp10332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V974)
}
__typedArg0 := V974
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10332)
}
__typedArg0 := tmp10332
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10333)
}
__typedArg0 := tmp10333
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10334)
}
__typedArg0 := tmp10334
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10325 Obj

if True == tmp10335 {
tmp10327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V974)
}
__typedArg0 := V974
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10327)
}
__typedArg0 := tmp10327
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10328)
}
__typedArg0 := tmp10328
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10329)
}
__typedArg0 := tmp10329
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10330)
}
__typedArg0 := Nil
__typedArg1 := tmp10330
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10326 Obj

if True == tmp10331 {
ifres10326 = True


} else {
ifres10326 = False


}

ifres10325 = ifres10326


} else {
ifres10325 = False


}

var ifres10324 Obj

if True == ifres10325 {
ifres10324 = True


} else {
ifres10324 = False


}

ifres10323 = ifres10324


} else {
ifres10323 = False


}

var ifres10322 Obj

if True == ifres10323 {
ifres10322 = True


} else {
ifres10322 = False


}

ifres10321 = ifres10322


} else {
ifres10321 = False


}

var ifres10320 Obj

if True == ifres10321 {
ifres10320 = True


} else {
ifres10320 = False


}

ifres10319 = ifres10320


} else {
ifres10319 = False


}

var ifres10318 Obj

if True == ifres10319 {
ifres10318 = True


} else {
ifres10318 = False


}

ifres10317 = ifres10318


} else {
ifres10317 = False


}

var ifres10316 Obj

if True == ifres10317 {
ifres10316 = True


} else {
ifres10316 = False


}

ifres10315 = ifres10316


} else {
ifres10315 = False


}

var ifres10314 Obj

if True == ifres10315 {
ifres10314 = True


} else {
ifres10314 = False


}

ifres10313 = ifres10314


} else {
ifres10313 = False


}

if True == ifres10313 {
__e.TailApply(PrimFunc(sympr), MakeString("(defun fail () shen.fail!)"), V975)
return
} else {
tmp10311 := Call(__e, PrimFunc(symshen_4app), V974, MakeString("\n\n"), symshen_4r)


__e.TailApply(PrimFunc(sympr), tmp10311, V975)
return


}


}, 2)

tmp10351 := Call(__e, ns2_1set, symshen_4write_1kl_1h, tmp10310)


_ = tmp10351

tmp10352 := MakeNative(func(__e *ControlFlow) {
V976 := __e.Get(1)
_ = V976
tmp10361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V976)
}
__typedArg0 := MakeString("")
__typedArg1 := V976
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10361 {
__e.Return(MakeString(".kl"))
return
} else {
tmp10359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(".shen"), V976)
}
__typedArg0 := MakeString(".shen")
__typedArg1 := V976
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10359 {
__e.Return(MakeString(".kl"))
return
} else {
tmp10357 := Call(__e, PrimFunc(symshen_4_7string_2), V976)


if True == tmp10357 {
tmp10353 := Call(__e, PrimFunc(symhdstr), V976)


tmp10355 := Call(__e, PrimFunc(symshen_4klfile), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V976)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V976
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.TailApply(PrimFunc(sym_8s), tmp10353, tmp10355)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.klfile"))
}
__typedArg0 := MakeString("partial function shen.klfile")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 1)

__e.TailApply(ns2_1set, symshen_4klfile, tmp10352)
return




}, 0)

