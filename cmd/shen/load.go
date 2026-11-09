package main

import . "github.com/pyrex41/shen-go/kl"

var LoadMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp8951 := MakeNative1(func(__e *ControlFlow, V1262 Obj) Obj {
tmp8952 := GValue(__e, symshen_4_dtc_d)

let__8933 := tmp8952
_ = let__8933

tmp8953 := PrimGuard1(__e, symget_1time, PrimGetTime, symrun)

let__8935 := tmp8953
_ = let__8935

tmp8954 := Call1(__e, SymFn(symread_1file), V1262)


tmp8955 := Call2(__e, SymFn(symshen_4load_1help), let__8933, tmp8954)


let__8936 := tmp8955
_ = let__8936

tmp8956 := PrimGuard1(__e, symget_1time, PrimGetTime, symrun)

let__8937 := tmp8956
_ = let__8937

let__8938 := GSub(__e, let__8937, let__8935)
_ = let__8938

tmp8958 := PrimGuard1(__e, symstr, PrimStr, let__8938)

tmp8960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\nrun time: "))
__typedS1, __typedOK1 := TypedString(tmp8958)
__typedS2, __typedOK2 := TypedString(MakeString(" secs\n"))
if __typedOK0 && __typedOK1 && __typedOK2 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + (__typedS1 + __typedS2)))
}}
__typedArg0 := MakeString("\nrun time: ")
__typedArg1 := PrimGuard2(__e, symcn, PrimStringConcat, tmp8958, MakeString(" secs\n"))
return Call2(__e, SymFn(symcn), __typedArg0, __typedArg1)
})()

tmp8961 := Call0(__e, SymFn(symstoutput))


tmp8962 := Call2(__e, SymFn(sympr), tmp8960, tmp8961)


let__8939 := tmp8962
_ = let__8939

let__8934 := let__8936
_ = let__8934

var ifres8963 Obj

if True == let__8933 {
tmp8964 := Call0(__e, SymFn(syminferences))


tmp8965 := Call3(__e, SymFn(symshen_4app), tmp8964, MakeString(" inferences\n"), symshen_4a)


tmp8966 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("\ntypechecked in "), tmp8965)

tmp8967 := Call0(__e, SymFn(symstoutput))


tmp8968 := Call2(__e, SymFn(sympr), tmp8966, tmp8967)


ifres8963 = tmp8968


} else {
ifres8963 = symshen_4skip


}

let__8940 := ifres8963
_ = let__8940

return symloaded


panic("unreachable")
})

tmp8969 := Call2(__e, ns2_1set, symload, tmp8951)


_ = tmp8969

tmp8970 := MakeNative2(func(__e *ControlFlow, V1273, V1274 Obj) Obj {
tmp8972 := GEq(__e, False, V1273)

if True == tmp8972 {
return TailCall1(__e, SymFn(symshen_4eval_1and_1print), V1274)
} else {
return TailCall1(__e, SymFn(symshen_4check_1eval_1and_1print), V1274)
}


panic("unreachable")
})

tmp8973 := Call2(__e, ns2_1set, symshen_4load_1help, tmp8970)


_ = tmp8973

tmp8974 := MakeNative1(func(__e *ControlFlow, V1275 Obj) Obj {
tmp8975 := MakeNative1(func(__e *ControlFlow, Z1276 Obj) Obj {
tmp8976 := Call1(__e, SymFn(symshen_4shen_1_6kl), Z1276)


tmp8977 := Call1(__e, SymFn(symeval_1kl), tmp8976)


tmp8978 := Call3(__e, SymFn(symshen_4app), tmp8977, MakeString("\n"), symshen_4s)


tmp8979 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp8978, tmp8979)


panic("unreachable")
})

return TailCall2(__e, SymFn(symshen_4for_1each), tmp8975, V1275)


panic("unreachable")
})

tmp8980 := Call2(__e, ns2_1set, symshen_4eval_1and_1print, tmp8974)


_ = tmp8980

tmp8981 := MakeNative1(func(__e *ControlFlow, V1277 Obj) Obj {
tmp8982 := MakeNative1(func(__e *ControlFlow, Z1279 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4typetable), Z1279)
panic("unreachable")
})

tmp8983 := Call2(__e, SymFn(symmapcan), tmp8982, V1277)


let__8941 := tmp8983
_ = let__8941

tmp8984 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, SymFn(symshen_4assumetypes), let__8941)
panic("unreachable")
})

tmp8985 := MakeNative1(func(__e *ControlFlow, Z1281 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4unwind_1types), Z1281, let__8941)
panic("unreachable")
})

tmp8986 := Call2(__e, try_1catch, tmp8984, tmp8985)


let__8942 := tmp8986
_ = let__8942

tmp8987 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, SymFn(symshen_4work_1through), V1277)
panic("unreachable")
})

tmp8988 := MakeNative1(func(__e *ControlFlow, Z1282 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4unwind_1types), Z1282, let__8941)
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp8987, tmp8988)


panic("unreachable")
})

tmp8989 := Call2(__e, ns2_1set, symshen_4check_1eval_1and_1print, tmp8981)


_ = tmp8989

tmp8990 := MakeNative1(func(__e *ControlFlow, V1287 Obj) Obj {
tmp9029 := GConsP(__e, V1287)

var ifres9014 Obj

if True == tmp9029 {
tmp9027 := GHd(__e, V1287)

tmp9028 := GEq(__e, symdefine, tmp9027)

var ifres9015 Obj

if True == tmp9028 {
tmp9025 := GTl(__e, V1287)

tmp9026 := GConsP(__e, tmp9025)

var ifres9016 Obj

if True == tmp9026 {
tmp9022 := GTl(__e, V1287)

tmp9023 := GTl(__e, tmp9022)

tmp9024 := GConsP(__e, tmp9023)

var ifres9017 Obj

if True == tmp9024 {
tmp9018 := GTl(__e, V1287)

tmp9019 := GTl(__e, tmp9018)

tmp9020 := GHd(__e, tmp9019)

tmp9021 := GEq(__e, sym_i, tmp9020)

ifres9017 = tmp9021


} else {
ifres9017 = False


}

ifres9016 = ifres9017


} else {
ifres9016 = False


}

ifres9015 = ifres9016


} else {
ifres9015 = False


}

ifres9014 = ifres9015


} else {
ifres9014 = False


}

if True == ifres9014 {
tmp8991 := GTl(__e, V1287)

tmp8992 := GHd(__e, tmp8991)

tmp8993 := GTl(__e, V1287)

tmp8994 := GHd(__e, tmp8993)

tmp8995 := GTl(__e, V1287)

tmp8996 := GTl(__e, tmp8995)

tmp8997 := GTl(__e, tmp8996)

tmp8998 := Call2(__e, SymFn(symshen_4type_1F), tmp8994, tmp8997)


tmp8999 := Call1(__e, SymFn(symshen_4rectify_1type), tmp8998)


tmp9000 := GCons(__e, tmp8999, Nil)

return GCons(__e, tmp8992, tmp9000)


} else {
tmp9012 := GConsP(__e, V1287)

var ifres9006 Obj

if True == tmp9012 {
tmp9010 := GHd(__e, V1287)

tmp9011 := GEq(__e, symdefine, tmp9010)

var ifres9007 Obj

if True == tmp9011 {
tmp9008 := GTl(__e, V1287)

tmp9009 := GConsP(__e, tmp9008)

ifres9007 = tmp9009


} else {
ifres9007 = False


}

ifres9006 = ifres9007


} else {
ifres9006 = False


}

if True == ifres9006 {
tmp9001 := GTl(__e, V1287)

tmp9002 := GHd(__e, tmp9001)

tmp9003 := Call3(__e, SymFn(symshen_4app), tmp9002, MakeString("\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("missing { in "), tmp9003))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("missing { in "), tmp9003)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


} else {
return Nil
}


}


panic("unreachable")
})

tmp9030 := Call2(__e, ns2_1set, symshen_4typetable, tmp8990)


_ = tmp9030

tmp9031 := MakeNative2(func(__e *ControlFlow, V1294, V1295 Obj) Obj {
tmp9043 := GConsP(__e, V1295)

var ifres9040 Obj

if True == tmp9043 {
tmp9041 := GHd(__e, V1295)

tmp9042 := GEq(__e, sym_j, tmp9041)

ifres9040 = tmp9042


} else {
ifres9040 = False


}

if True == ifres9040 {
return Nil
} else {
tmp9038 := GConsP(__e, V1295)

if True == tmp9038 {
tmp9032 := GHd(__e, V1295)

tmp9033 := GTl(__e, V1295)

tmp9034 := Call2(__e, SymFn(symshen_4type_1F), V1294, tmp9033)


return GCons(__e, tmp9032, tmp9034)


} else {
tmp9035 := Call3(__e, SymFn(symshen_4app), V1294, MakeString("\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("missing } in "), tmp9035))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("missing } in "), tmp9035)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


}


panic("unreachable")
})

tmp9044 := Call2(__e, ns2_1set, symshen_4type_1F, tmp9031)


_ = tmp9044

tmp9045 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V1298 := __self1
_ = V1298
tmp9058 := GEq(__e, Nil, V1298)

if True == tmp9058 {
return Nil
} else {
tmp9056 := GConsP(__e, V1298)

var ifres9053 Obj

if True == tmp9056 {
tmp9054 := GTl(__e, V1298)

tmp9055 := GConsP(__e, tmp9054)

ifres9053 = tmp9055


} else {
ifres9053 = False


}

if True == ifres9053 {
tmp9046 := GHd(__e, V1298)

tmp9047 := GTl(__e, V1298)

tmp9048 := GHd(__e, tmp9047)

tmp9049 := Call2(__e, SymFn(symdeclare), tmp9046, tmp9048)


_ = tmp9049

tmp9050 := GTl(__e, V1298)

tmp9051 := GTl(__e, tmp9050)

if SymFn(symshen_4assumetypes) == __self {
__self1 = tmp9051
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4assumetypes), tmp9051)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.assumetype"))
}


}


panic("unreachable")
})
return __self
})()

tmp9059 := Call2(__e, ns2_1set, symshen_4assumetypes, tmp9045)


_ = tmp9059

tmp9060 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V1303 := __self1
_ = V1303
V1304 := __self2
_ = V1304
tmp9070 := GConsP(__e, V1304)

var ifres9067 Obj

if True == tmp9070 {
tmp9068 := GTl(__e, V1304)

tmp9069 := GConsP(__e, tmp9068)

ifres9067 = tmp9069


} else {
ifres9067 = False


}

if True == ifres9067 {
tmp9061 := GHd(__e, V1304)

tmp9062 := Call1(__e, SymFn(symdestroy), tmp9061)


_ = tmp9062

tmp9063 := GTl(__e, V1304)

tmp9064 := GTl(__e, tmp9063)

if SymFn(symshen_4unwind_1types) == __self {
__self1, __self2 = V1303, tmp9064
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4unwind_1types), V1303, tmp9064)


} else {
tmp9065 := PrimGuard1(__e, symerror_1to_1string, PrimErrorToString, V1303)

return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp9065)


}


panic("unreachable")
})
return __self
})()

tmp9071 := Call2(__e, ns2_1set, symshen_4unwind_1types, tmp9060)


_ = tmp9071

tmp9072 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V1307 := __self1
_ = V1307
tmp9115 := GEq(__e, Nil, V1307)

if True == tmp9115 {
return Nil
} else {
tmp9113 := GConsP(__e, V1307)

var ifres9101 Obj

if True == tmp9113 {
tmp9111 := GTl(__e, V1307)

tmp9112 := GConsP(__e, tmp9111)

var ifres9102 Obj

if True == tmp9112 {
tmp9108 := GTl(__e, V1307)

tmp9109 := GTl(__e, tmp9108)

tmp9110 := GConsP(__e, tmp9109)

var ifres9103 Obj

if True == tmp9110 {
tmp9104 := GTl(__e, V1307)

tmp9105 := GHd(__e, tmp9104)

tmp9106 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp9107 := GEq(__e, tmp9105, tmp9106)

ifres9103 = tmp9107


} else {
ifres9103 = False


}

ifres9102 = ifres9103


} else {
ifres9102 = False


}

ifres9101 = ifres9102


} else {
ifres9101 = False


}

if True == ifres9101 {
tmp9073 := GHd(__e, V1307)

tmp9074 := GTl(__e, V1307)

tmp9075 := GTl(__e, tmp9074)

tmp9076 := GHd(__e, tmp9075)

tmp9077 := Call2(__e, SymFn(symshen_4typecheck), tmp9073, tmp9076)


let__8943 := tmp9077
_ = let__8943

tmp9091 := GEq(__e, let__8943, False)

if True == tmp9091 {
return TailCall0(__e, SymFn(symshen_4type_1error))
} else {
tmp9078 := GHd(__e, V1307)

tmp9079 := Call1(__e, SymFn(symshen_4shen_1_6kl), tmp9078)


tmp9080 := Call1(__e, SymFn(symeval_1kl), tmp9079)


let__8944 := tmp9080
_ = let__8944

tmp9081 := Call1(__e, SymFn(symshen_4pretty_1type), let__8943)


tmp9082 := Call3(__e, SymFn(symshen_4app), tmp9081, MakeString("\n"), symshen_4r)


tmp9084 := Call3(__e, SymFn(symshen_4app), let__8944, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" : "), tmp9082), symshen_4s)


tmp9085 := Call0(__e, SymFn(symstoutput))


tmp9086 := Call2(__e, SymFn(sympr), tmp9084, tmp9085)


let__8945 := tmp9086
_ = let__8945

tmp9087 := GTl(__e, V1307)

tmp9088 := GTl(__e, tmp9087)

tmp9089 := GTl(__e, tmp9088)

if SymFn(symshen_4work_1through) == __self {
__self1 = tmp9089
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4work_1through), tmp9089)


}


} else {
tmp9099 := GConsP(__e, V1307)

if True == tmp9099 {
tmp9092 := GHd(__e, V1307)

tmp9093 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp9094 := GTl(__e, V1307)

tmp9095 := GCons(__e, symA, tmp9094)

tmp9096 := GCons(__e, tmp9093, tmp9095)

tmp9097 := GCons(__e, tmp9092, tmp9096)

if SymFn(symshen_4work_1through) == __self {
__self1 = tmp9097
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4work_1through), tmp9097)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.work-through"))
}


}


}


panic("unreachable")
})
return __self
})()

tmp9116 := Call2(__e, ns2_1set, symshen_4work_1through, tmp9072)


_ = tmp9116

tmp9117 := MakeNative1(func(__e *ControlFlow, V1312 Obj) Obj {
tmp9241 := GConsP(__e, V1312)

var ifres9133 Obj

if True == tmp9241 {
tmp9239 := GHd(__e, V1312)

tmp9240 := GConsP(__e, tmp9239)

var ifres9134 Obj

if True == tmp9240 {
tmp9236 := GHd(__e, V1312)

tmp9237 := GHd(__e, tmp9236)

tmp9238 := GEq(__e, symlist, tmp9237)

var ifres9135 Obj

if True == tmp9238 {
tmp9233 := GHd(__e, V1312)

tmp9234 := GTl(__e, tmp9233)

tmp9235 := GConsP(__e, tmp9234)

var ifres9136 Obj

if True == tmp9235 {
tmp9229 := GHd(__e, V1312)

tmp9230 := GTl(__e, tmp9229)

tmp9231 := GTl(__e, tmp9230)

tmp9232 := GEq(__e, Nil, tmp9231)

var ifres9137 Obj

if True == tmp9232 {
tmp9227 := GTl(__e, V1312)

tmp9228 := GConsP(__e, tmp9227)

var ifres9138 Obj

if True == tmp9228 {
tmp9224 := GTl(__e, V1312)

tmp9225 := GHd(__e, tmp9224)

tmp9226 := GEq(__e, sym_1_1_6, tmp9225)

var ifres9139 Obj

if True == tmp9226 {
tmp9221 := GTl(__e, V1312)

tmp9222 := GTl(__e, tmp9221)

tmp9223 := GConsP(__e, tmp9222)

var ifres9140 Obj

if True == tmp9223 {
tmp9217 := GTl(__e, V1312)

tmp9218 := GTl(__e, tmp9217)

tmp9219 := GHd(__e, tmp9218)

tmp9220 := GConsP(__e, tmp9219)

var ifres9141 Obj

if True == tmp9220 {
tmp9212 := GTl(__e, V1312)

tmp9213 := GTl(__e, tmp9212)

tmp9214 := GHd(__e, tmp9213)

tmp9215 := GHd(__e, tmp9214)

tmp9216 := GEq(__e, symstr, tmp9215)

var ifres9142 Obj

if True == tmp9216 {
tmp9207 := GTl(__e, V1312)

tmp9208 := GTl(__e, tmp9207)

tmp9209 := GHd(__e, tmp9208)

tmp9210 := GTl(__e, tmp9209)

tmp9211 := GConsP(__e, tmp9210)

var ifres9143 Obj

if True == tmp9211 {
tmp9201 := GTl(__e, V1312)

tmp9202 := GTl(__e, tmp9201)

tmp9203 := GHd(__e, tmp9202)

tmp9204 := GTl(__e, tmp9203)

tmp9205 := GHd(__e, tmp9204)

tmp9206 := GConsP(__e, tmp9205)

var ifres9144 Obj

if True == tmp9206 {
tmp9194 := GTl(__e, V1312)

tmp9195 := GTl(__e, tmp9194)

tmp9196 := GHd(__e, tmp9195)

tmp9197 := GTl(__e, tmp9196)

tmp9198 := GHd(__e, tmp9197)

tmp9199 := GHd(__e, tmp9198)

tmp9200 := GEq(__e, symlist, tmp9199)

var ifres9145 Obj

if True == tmp9200 {
tmp9187 := GTl(__e, V1312)

tmp9188 := GTl(__e, tmp9187)

tmp9189 := GHd(__e, tmp9188)

tmp9190 := GTl(__e, tmp9189)

tmp9191 := GHd(__e, tmp9190)

tmp9192 := GTl(__e, tmp9191)

tmp9193 := GConsP(__e, tmp9192)

var ifres9146 Obj

if True == tmp9193 {
tmp9179 := GTl(__e, V1312)

tmp9180 := GTl(__e, tmp9179)

tmp9181 := GHd(__e, tmp9180)

tmp9182 := GTl(__e, tmp9181)

tmp9183 := GHd(__e, tmp9182)

tmp9184 := GTl(__e, tmp9183)

tmp9185 := GTl(__e, tmp9184)

tmp9186 := GEq(__e, Nil, tmp9185)

var ifres9147 Obj

if True == tmp9186 {
tmp9173 := GTl(__e, V1312)

tmp9174 := GTl(__e, tmp9173)

tmp9175 := GHd(__e, tmp9174)

tmp9176 := GTl(__e, tmp9175)

tmp9177 := GTl(__e, tmp9176)

tmp9178 := GConsP(__e, tmp9177)

var ifres9148 Obj

if True == tmp9178 {
tmp9166 := GTl(__e, V1312)

tmp9167 := GTl(__e, tmp9166)

tmp9168 := GHd(__e, tmp9167)

tmp9169 := GTl(__e, tmp9168)

tmp9170 := GTl(__e, tmp9169)

tmp9171 := GTl(__e, tmp9170)

tmp9172 := GEq(__e, Nil, tmp9171)

var ifres9149 Obj

if True == tmp9172 {
tmp9162 := GTl(__e, V1312)

tmp9163 := GTl(__e, tmp9162)

tmp9164 := GTl(__e, tmp9163)

tmp9165 := GEq(__e, Nil, tmp9164)

var ifres9150 Obj

if True == tmp9165 {
tmp9151 := GHd(__e, V1312)

tmp9152 := GTl(__e, tmp9151)

tmp9153 := GHd(__e, tmp9152)

tmp9154 := GTl(__e, V1312)

tmp9155 := GTl(__e, tmp9154)

tmp9156 := GHd(__e, tmp9155)

tmp9157 := GTl(__e, tmp9156)

tmp9158 := GHd(__e, tmp9157)

tmp9159 := GTl(__e, tmp9158)

tmp9160 := GHd(__e, tmp9159)

tmp9161 := GEq(__e, tmp9153, tmp9160)

ifres9150 = tmp9161


} else {
ifres9150 = False


}

ifres9149 = ifres9150


} else {
ifres9149 = False


}

ifres9148 = ifres9149


} else {
ifres9148 = False


}

ifres9147 = ifres9148


} else {
ifres9147 = False


}

ifres9146 = ifres9147


} else {
ifres9146 = False


}

ifres9145 = ifres9146


} else {
ifres9145 = False


}

ifres9144 = ifres9145


} else {
ifres9144 = False


}

ifres9143 = ifres9144


} else {
ifres9143 = False


}

ifres9142 = ifres9143


} else {
ifres9142 = False


}

ifres9141 = ifres9142


} else {
ifres9141 = False


}

ifres9140 = ifres9141


} else {
ifres9140 = False


}

ifres9139 = ifres9140


} else {
ifres9139 = False


}

ifres9138 = ifres9139


} else {
ifres9138 = False


}

ifres9137 = ifres9138


} else {
ifres9137 = False


}

ifres9136 = ifres9137


} else {
ifres9136 = False


}

ifres9135 = ifres9136


} else {
ifres9135 = False


}

ifres9134 = ifres9135


} else {
ifres9134 = False


}

ifres9133 = ifres9134


} else {
ifres9133 = False


}

if True == ifres9133 {
tmp9118 := GTl(__e, V1312)

tmp9119 := GTl(__e, tmp9118)

tmp9120 := GHd(__e, tmp9119)

tmp9121 := GTl(__e, tmp9120)

tmp9122 := GHd(__e, tmp9121)

tmp9123 := GTl(__e, V1312)

tmp9124 := GTl(__e, tmp9123)

tmp9125 := GHd(__e, tmp9124)

tmp9126 := GTl(__e, tmp9125)

tmp9127 := GTl(__e, tmp9126)

tmp9128 := GCons(__e, sym_a_a_6, tmp9127)

return GCons(__e, tmp9122, tmp9128)


} else {
tmp9131 := GConsP(__e, V1312)

if True == tmp9131 {
tmp9129 := MakeNative1(func(__e *ControlFlow, Z1313 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4pretty_1type), Z1313)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp9129, V1312)


} else {
return V1312
}


}


panic("unreachable")
})

tmp9242 := Call2(__e, ns2_1set, symshen_4pretty_1type, tmp9117)


_ = tmp9242

tmp9243 := MakeNative0(func(__e *ControlFlow) Obj {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("type error\n"))
panic("unreachable")
})

tmp9244 := Call2(__e, ns2_1set, symshen_4type_1error, tmp9243)


_ = tmp9244

tmp9245 := MakeNative1(func(__e *ControlFlow, V1314 Obj) Obj {
tmp9246 := Call1(__e, SymFn(symshen_4klfile), V1314)


let__8946 := tmp9246
_ = let__8946

tmp9247 := Call1(__e, SymFn(symread_1file), V1314)


let__8947 := tmp9247
_ = let__8947

tmp9248 := PrimGuard2(__e, symopen, PrimOpenStream, let__8946, symout)

let__8948 := tmp9248
_ = let__8948

tmp9249 := MakeNative1(func(__e *ControlFlow, Z1319 Obj) Obj {
tmp9250 := Call1(__e, SymFn(symshen_4shen_1_6kl_1h), Z1319)


return TailCall1(__e, SymFn(symshen_4partial), tmp9250)


panic("unreachable")
})

tmp9251 := Call2(__e, SymFn(symmap), tmp9249, let__8947)


let__8949 := tmp9251
_ = let__8949

tmp9252 := Call2(__e, SymFn(symshen_4write_1kl), let__8949, let__8948)


let__8950 := tmp9252
_ = let__8950

return let__8946


panic("unreachable")
})

tmp9253 := Call2(__e, ns2_1set, symbootstrap, tmp9245)


_ = tmp9253

tmp9254 := MakeNative1(func(__e *ControlFlow, V1321 Obj) Obj {
tmp9274 := GConsP(__e, V1321)

var ifres9264 Obj

if True == tmp9274 {
tmp9272 := GHd(__e, V1321)

tmp9273 := GEq(__e, symshen_4f_1error, tmp9272)

var ifres9265 Obj

if True == tmp9273 {
tmp9270 := GTl(__e, V1321)

tmp9271 := GConsP(__e, tmp9270)

var ifres9266 Obj

if True == tmp9271 {
tmp9267 := GTl(__e, V1321)

tmp9268 := GTl(__e, tmp9267)

tmp9269 := GEq(__e, Nil, tmp9268)

ifres9266 = tmp9269


} else {
ifres9266 = False


}

ifres9265 = ifres9266


} else {
ifres9265 = False


}

ifres9264 = ifres9265


} else {
ifres9264 = False


}

if True == ifres9264 {
tmp9255 := GTl(__e, V1321)

tmp9256 := GHd(__e, tmp9255)

tmp9257 := PrimGuard1(__e, symstr, PrimStr, tmp9256)

tmp9259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("partial function "), tmp9257), Nil)
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("partial function "), tmp9257)
__typedArg1 := Nil
return Call2(__e, SymFn(symcons), __typedArg0, __typedArg1)
})()

return GCons(__e, symsimple_1error, tmp9259)


} else {
tmp9262 := GConsP(__e, V1321)

if True == tmp9262 {
tmp9260 := MakeNative1(func(__e *ControlFlow, Z1322 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4partial), Z1322)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp9260, V1321)


} else {
return V1321
}


}


panic("unreachable")
})

tmp9275 := Call2(__e, ns2_1set, symshen_4partial, tmp9254)


_ = tmp9275

tmp9276 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V1325 := __self1
_ = V1325
V1326 := __self2
_ = V1326
tmp9289 := GEq(__e, Nil, V1325)

if True == tmp9289 {
return PrimGuard1(__e, symclose, PrimCloseStream, V1326)
} else {
tmp9287 := GConsP(__e, V1325)

var ifres9284 Obj

if True == tmp9287 {
tmp9285 := GHd(__e, V1325)

tmp9286 := GConsP(__e, tmp9285)

ifres9284 = tmp9286


} else {
ifres9284 = False


}

if True == ifres9284 {
tmp9277 := GTl(__e, V1325)

tmp9278 := GHd(__e, V1325)

tmp9279 := Call2(__e, SymFn(symshen_4write_1kl_1h), tmp9278, V1326)


_ = tmp9279

if SymFn(symshen_4write_1kl) == __self {
__self1, __self2 = tmp9277, V1326
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4write_1kl), tmp9277, V1326)


} else {
tmp9282 := GConsP(__e, V1325)

if True == tmp9282 {
tmp9280 := GTl(__e, V1325)

if SymFn(symshen_4write_1kl) == __self {
__self1, __self2 = tmp9280, V1326
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4write_1kl), tmp9280, V1326)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4write_1kl)
}


}


}


panic("unreachable")
})
return __self
})()

tmp9290 := Call2(__e, ns2_1set, symshen_4write_1kl, tmp9276)


_ = tmp9290

tmp9291 := MakeNative2(func(__e *ControlFlow, V1329, V1330 Obj) Obj {
tmp9324 := GConsP(__e, V1329)

var ifres9294 Obj

if True == tmp9324 {
tmp9322 := GHd(__e, V1329)

tmp9323 := GEq(__e, symdefun, tmp9322)

var ifres9295 Obj

if True == tmp9323 {
tmp9320 := GTl(__e, V1329)

tmp9321 := GConsP(__e, tmp9320)

var ifres9296 Obj

if True == tmp9321 {
tmp9317 := GTl(__e, V1329)

tmp9318 := GHd(__e, tmp9317)

tmp9319 := GEq(__e, symfail, tmp9318)

var ifres9297 Obj

if True == tmp9319 {
tmp9314 := GTl(__e, V1329)

tmp9315 := GTl(__e, tmp9314)

tmp9316 := GConsP(__e, tmp9315)

var ifres9298 Obj

if True == tmp9316 {
tmp9310 := GTl(__e, V1329)

tmp9311 := GTl(__e, tmp9310)

tmp9312 := GHd(__e, tmp9311)

tmp9313 := GEq(__e, Nil, tmp9312)

var ifres9299 Obj

if True == tmp9313 {
tmp9306 := GTl(__e, V1329)

tmp9307 := GTl(__e, tmp9306)

tmp9308 := GTl(__e, tmp9307)

tmp9309 := GConsP(__e, tmp9308)

var ifres9300 Obj

if True == tmp9309 {
tmp9301 := GTl(__e, V1329)

tmp9302 := GTl(__e, tmp9301)

tmp9303 := GTl(__e, tmp9302)

tmp9304 := GTl(__e, tmp9303)

tmp9305 := GEq(__e, Nil, tmp9304)

ifres9300 = tmp9305


} else {
ifres9300 = False


}

ifres9299 = ifres9300


} else {
ifres9299 = False


}

ifres9298 = ifres9299


} else {
ifres9298 = False


}

ifres9297 = ifres9298


} else {
ifres9297 = False


}

ifres9296 = ifres9297


} else {
ifres9296 = False


}

ifres9295 = ifres9296


} else {
ifres9295 = False


}

ifres9294 = ifres9295


} else {
ifres9294 = False


}

if True == ifres9294 {
return TailCall2(__e, SymFn(sympr), MakeString("(defun fail () shen.fail!)"), V1330)
} else {
tmp9292 := Call3(__e, SymFn(symshen_4app), V1329, MakeString("\n\n"), symshen_4r)


return TailCall2(__e, SymFn(sympr), tmp9292, V1330)


}


panic("unreachable")
})

tmp9325 := Call2(__e, ns2_1set, symshen_4write_1kl_1h, tmp9291)


_ = tmp9325

tmp9326 := MakeNative1(func(__e *ControlFlow, V1331 Obj) Obj {
tmp9335 := GEq(__e, MakeString(""), V1331)

if True == tmp9335 {
return MakeString(".kl")
} else {
tmp9333 := GEq(__e, MakeString(".shen"), V1331)

if True == tmp9333 {
return MakeString(".kl")
} else {
tmp9331 := Call1(__e, SymFn(symshen_4_7string_2), V1331)


if True == tmp9331 {
tmp9327 := Call1(__e, SymFn(symhdstr), V1331)


tmp9329 := Call1(__e, SymFn(symshen_4klfile), PrimGuard1(__e, symtlstr, PrimTailString, V1331))


return TailCall2(__e, SymFn(sym_8s), tmp9327, tmp9329)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4klfile)
}


}


}


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4klfile, tmp9326)




panic("unreachable")
})

