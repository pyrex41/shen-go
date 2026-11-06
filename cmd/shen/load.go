package main

import . "github.com/pyrex41/shen-go/kl"

var LoadMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp8931 := MakeNative1(func(__e *ControlFlow, V1262 Obj) Obj {
tmp8932 := GValue(__e, symshen_4_dtc_d)

let__8913 := tmp8932
_ = let__8913

tmp8933 := PrimGuard1(__e, symget_1time, PrimGetTime, symrun)

let__8915 := tmp8933
_ = let__8915

tmp8934 := Call1(__e, SymFn(symread_1file), V1262)


tmp8935 := Call2(__e, SymFn(symshen_4load_1help), let__8913, tmp8934)


let__8916 := tmp8935
_ = let__8916

tmp8936 := PrimGuard1(__e, symget_1time, PrimGetTime, symrun)

let__8917 := tmp8936
_ = let__8917

let__8918 := GSub(__e, let__8917, let__8915)
_ = let__8918

tmp8938 := PrimGuard1(__e, symstr, PrimStr, let__8918)

tmp8940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\nrun time: "))
__typedS1, __typedOK1 := TypedString(tmp8938)
__typedS2, __typedOK2 := TypedString(MakeString(" secs\n"))
if __typedOK0 && __typedOK1 && __typedOK2 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + (__typedS1 + __typedS2)))
}}
__typedArg0 := MakeString("\nrun time: ")
__typedArg1 := PrimGuard2(__e, symcn, PrimStringConcat, tmp8938, MakeString(" secs\n"))
return Call2(__e, SymFn(symcn), __typedArg0, __typedArg1)
})()

tmp8941 := Call0(__e, SymFn(symstoutput))


tmp8942 := Call2(__e, SymFn(sympr), tmp8940, tmp8941)


let__8919 := tmp8942
_ = let__8919

let__8914 := let__8916
_ = let__8914

var ifres8943 Obj

if True == let__8913 {
tmp8944 := Call0(__e, SymFn(syminferences))


tmp8945 := Call3(__e, SymFn(symshen_4app), tmp8944, MakeString(" inferences\n"), symshen_4a)


tmp8946 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("\ntypechecked in "), tmp8945)

tmp8947 := Call0(__e, SymFn(symstoutput))


tmp8948 := Call2(__e, SymFn(sympr), tmp8946, tmp8947)


ifres8943 = tmp8948


} else {
ifres8943 = symshen_4skip


}

let__8920 := ifres8943
_ = let__8920

return symloaded


panic("unreachable")
})

tmp8949 := Call2(__e, ns2_1set, symload, tmp8931)


_ = tmp8949

tmp8950 := MakeNative2(func(__e *ControlFlow, V1273, V1274 Obj) Obj {
tmp8952 := GEq(__e, False, V1273)

if True == tmp8952 {
return TailCall1(__e, SymFn(symshen_4eval_1and_1print), V1274)
} else {
return TailCall1(__e, SymFn(symshen_4check_1eval_1and_1print), V1274)
}


panic("unreachable")
})

tmp8953 := Call2(__e, ns2_1set, symshen_4load_1help, tmp8950)


_ = tmp8953

tmp8954 := MakeNative1(func(__e *ControlFlow, V1275 Obj) Obj {
tmp8955 := MakeNative1(func(__e *ControlFlow, Z1276 Obj) Obj {
tmp8956 := Call1(__e, SymFn(symshen_4shen_1_6kl), Z1276)


tmp8957 := Call1(__e, SymFn(symeval_1kl), tmp8956)


tmp8958 := Call3(__e, SymFn(symshen_4app), tmp8957, MakeString("\n"), symshen_4s)


tmp8959 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp8958, tmp8959)


panic("unreachable")
})

return TailCall2(__e, SymFn(symshen_4for_1each), tmp8955, V1275)


panic("unreachable")
})

tmp8960 := Call2(__e, ns2_1set, symshen_4eval_1and_1print, tmp8954)


_ = tmp8960

tmp8961 := MakeNative1(func(__e *ControlFlow, V1277 Obj) Obj {
tmp8962 := MakeNative1(func(__e *ControlFlow, Z1279 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4typetable), Z1279)
panic("unreachable")
})

tmp8963 := Call2(__e, SymFn(symmapcan), tmp8962, V1277)


let__8921 := tmp8963
_ = let__8921

tmp8964 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, SymFn(symshen_4assumetypes), let__8921)
panic("unreachable")
})

tmp8965 := MakeNative1(func(__e *ControlFlow, Z1281 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4unwind_1types), Z1281, let__8921)
panic("unreachable")
})

tmp8966 := Call2(__e, try_1catch, tmp8964, tmp8965)


let__8922 := tmp8966
_ = let__8922

tmp8967 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, SymFn(symshen_4work_1through), V1277)
panic("unreachable")
})

tmp8968 := MakeNative1(func(__e *ControlFlow, Z1282 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4unwind_1types), Z1282, let__8921)
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp8967, tmp8968)


panic("unreachable")
})

tmp8969 := Call2(__e, ns2_1set, symshen_4check_1eval_1and_1print, tmp8961)


_ = tmp8969

tmp8970 := MakeNative1(func(__e *ControlFlow, V1287 Obj) Obj {
tmp9009 := GConsP(__e, V1287)

var ifres8994 Obj

if True == tmp9009 {
tmp9007 := GHd(__e, V1287)

tmp9008 := GEq(__e, symdefine, tmp9007)

var ifres8995 Obj

if True == tmp9008 {
tmp9005 := GTl(__e, V1287)

tmp9006 := GConsP(__e, tmp9005)

var ifres8996 Obj

if True == tmp9006 {
tmp9002 := GTl(__e, V1287)

tmp9003 := GTl(__e, tmp9002)

tmp9004 := GConsP(__e, tmp9003)

var ifres8997 Obj

if True == tmp9004 {
tmp8998 := GTl(__e, V1287)

tmp8999 := GTl(__e, tmp8998)

tmp9000 := GHd(__e, tmp8999)

tmp9001 := GEq(__e, sym_i, tmp9000)

ifres8997 = tmp9001


} else {
ifres8997 = False


}

ifres8996 = ifres8997


} else {
ifres8996 = False


}

ifres8995 = ifres8996


} else {
ifres8995 = False


}

ifres8994 = ifres8995


} else {
ifres8994 = False


}

if True == ifres8994 {
tmp8971 := GTl(__e, V1287)

tmp8972 := GHd(__e, tmp8971)

tmp8973 := GTl(__e, V1287)

tmp8974 := GHd(__e, tmp8973)

tmp8975 := GTl(__e, V1287)

tmp8976 := GTl(__e, tmp8975)

tmp8977 := GTl(__e, tmp8976)

tmp8978 := Call2(__e, SymFn(symshen_4type_1F), tmp8974, tmp8977)


tmp8979 := Call1(__e, SymFn(symshen_4rectify_1type), tmp8978)


tmp8980 := GCons(__e, tmp8979, Nil)

return GCons(__e, tmp8972, tmp8980)


} else {
tmp8992 := GConsP(__e, V1287)

var ifres8986 Obj

if True == tmp8992 {
tmp8990 := GHd(__e, V1287)

tmp8991 := GEq(__e, symdefine, tmp8990)

var ifres8987 Obj

if True == tmp8991 {
tmp8988 := GTl(__e, V1287)

tmp8989 := GConsP(__e, tmp8988)

ifres8987 = tmp8989


} else {
ifres8987 = False


}

ifres8986 = ifres8987


} else {
ifres8986 = False


}

if True == ifres8986 {
tmp8981 := GTl(__e, V1287)

tmp8982 := GHd(__e, tmp8981)

tmp8983 := Call3(__e, SymFn(symshen_4app), tmp8982, MakeString("\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("missing { in "), tmp8983))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("missing { in "), tmp8983)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


} else {
return Nil
}


}


panic("unreachable")
})

tmp9010 := Call2(__e, ns2_1set, symshen_4typetable, tmp8970)


_ = tmp9010

tmp9011 := MakeNative2(func(__e *ControlFlow, V1294, V1295 Obj) Obj {
tmp9023 := GConsP(__e, V1295)

var ifres9020 Obj

if True == tmp9023 {
tmp9021 := GHd(__e, V1295)

tmp9022 := GEq(__e, sym_j, tmp9021)

ifres9020 = tmp9022


} else {
ifres9020 = False


}

if True == ifres9020 {
return Nil
} else {
tmp9018 := GConsP(__e, V1295)

if True == tmp9018 {
tmp9012 := GHd(__e, V1295)

tmp9013 := GTl(__e, V1295)

tmp9014 := Call2(__e, SymFn(symshen_4type_1F), V1294, tmp9013)


return GCons(__e, tmp9012, tmp9014)


} else {
tmp9015 := Call3(__e, SymFn(symshen_4app), V1294, MakeString("\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("missing } in "), tmp9015))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("missing } in "), tmp9015)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


}


panic("unreachable")
})

tmp9024 := Call2(__e, ns2_1set, symshen_4type_1F, tmp9011)


_ = tmp9024

tmp9025 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V1298 := __self1
_ = V1298
tmp9038 := GEq(__e, Nil, V1298)

if True == tmp9038 {
return Nil
} else {
tmp9036 := GConsP(__e, V1298)

var ifres9033 Obj

if True == tmp9036 {
tmp9034 := GTl(__e, V1298)

tmp9035 := GConsP(__e, tmp9034)

ifres9033 = tmp9035


} else {
ifres9033 = False


}

if True == ifres9033 {
tmp9026 := GHd(__e, V1298)

tmp9027 := GTl(__e, V1298)

tmp9028 := GHd(__e, tmp9027)

tmp9029 := Call2(__e, SymFn(symdeclare), tmp9026, tmp9028)


_ = tmp9029

tmp9030 := GTl(__e, V1298)

tmp9031 := GTl(__e, tmp9030)

if SymFn(symshen_4assumetypes) == __self {
__self1 = tmp9031
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4assumetypes), tmp9031)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.assumetype"))
}


}


panic("unreachable")
})
return __self
})()

tmp9039 := Call2(__e, ns2_1set, symshen_4assumetypes, tmp9025)


_ = tmp9039

tmp9040 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V1303 := __self1
_ = V1303
V1304 := __self2
_ = V1304
tmp9050 := GConsP(__e, V1304)

var ifres9047 Obj

if True == tmp9050 {
tmp9048 := GTl(__e, V1304)

tmp9049 := GConsP(__e, tmp9048)

ifres9047 = tmp9049


} else {
ifres9047 = False


}

if True == ifres9047 {
tmp9041 := GHd(__e, V1304)

tmp9042 := Call1(__e, SymFn(symdestroy), tmp9041)


_ = tmp9042

tmp9043 := GTl(__e, V1304)

tmp9044 := GTl(__e, tmp9043)

if SymFn(symshen_4unwind_1types) == __self {
__self1, __self2 = V1303, tmp9044
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4unwind_1types), V1303, tmp9044)


} else {
tmp9045 := PrimGuard1(__e, symerror_1to_1string, PrimErrorToString, V1303)

return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp9045)


}


panic("unreachable")
})
return __self
})()

tmp9051 := Call2(__e, ns2_1set, symshen_4unwind_1types, tmp9040)


_ = tmp9051

tmp9052 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V1307 := __self1
_ = V1307
tmp9095 := GEq(__e, Nil, V1307)

if True == tmp9095 {
return Nil
} else {
tmp9093 := GConsP(__e, V1307)

var ifres9081 Obj

if True == tmp9093 {
tmp9091 := GTl(__e, V1307)

tmp9092 := GConsP(__e, tmp9091)

var ifres9082 Obj

if True == tmp9092 {
tmp9088 := GTl(__e, V1307)

tmp9089 := GTl(__e, tmp9088)

tmp9090 := GConsP(__e, tmp9089)

var ifres9083 Obj

if True == tmp9090 {
tmp9084 := GTl(__e, V1307)

tmp9085 := GHd(__e, tmp9084)

tmp9086 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp9087 := GEq(__e, tmp9085, tmp9086)

ifres9083 = tmp9087


} else {
ifres9083 = False


}

ifres9082 = ifres9083


} else {
ifres9082 = False


}

ifres9081 = ifres9082


} else {
ifres9081 = False


}

if True == ifres9081 {
tmp9053 := GHd(__e, V1307)

tmp9054 := GTl(__e, V1307)

tmp9055 := GTl(__e, tmp9054)

tmp9056 := GHd(__e, tmp9055)

tmp9057 := Call2(__e, SymFn(symshen_4typecheck), tmp9053, tmp9056)


let__8923 := tmp9057
_ = let__8923

tmp9071 := GEq(__e, let__8923, False)

if True == tmp9071 {
return TailCall0(__e, SymFn(symshen_4type_1error))
} else {
tmp9058 := GHd(__e, V1307)

tmp9059 := Call1(__e, SymFn(symshen_4shen_1_6kl), tmp9058)


tmp9060 := Call1(__e, SymFn(symeval_1kl), tmp9059)


let__8924 := tmp9060
_ = let__8924

tmp9061 := Call1(__e, SymFn(symshen_4pretty_1type), let__8923)


tmp9062 := Call3(__e, SymFn(symshen_4app), tmp9061, MakeString("\n"), symshen_4r)


tmp9064 := Call3(__e, SymFn(symshen_4app), let__8924, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" : "), tmp9062), symshen_4s)


tmp9065 := Call0(__e, SymFn(symstoutput))


tmp9066 := Call2(__e, SymFn(sympr), tmp9064, tmp9065)


let__8925 := tmp9066
_ = let__8925

tmp9067 := GTl(__e, V1307)

tmp9068 := GTl(__e, tmp9067)

tmp9069 := GTl(__e, tmp9068)

if SymFn(symshen_4work_1through) == __self {
__self1 = tmp9069
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4work_1through), tmp9069)


}


} else {
tmp9079 := GConsP(__e, V1307)

if True == tmp9079 {
tmp9072 := GHd(__e, V1307)

tmp9073 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp9074 := GTl(__e, V1307)

tmp9075 := GCons(__e, symA, tmp9074)

tmp9076 := GCons(__e, tmp9073, tmp9075)

tmp9077 := GCons(__e, tmp9072, tmp9076)

if SymFn(symshen_4work_1through) == __self {
__self1 = tmp9077
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4work_1through), tmp9077)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.work-through"))
}


}


}


panic("unreachable")
})
return __self
})()

tmp9096 := Call2(__e, ns2_1set, symshen_4work_1through, tmp9052)


_ = tmp9096

tmp9097 := MakeNative1(func(__e *ControlFlow, V1312 Obj) Obj {
tmp9221 := GConsP(__e, V1312)

var ifres9113 Obj

if True == tmp9221 {
tmp9219 := GHd(__e, V1312)

tmp9220 := GConsP(__e, tmp9219)

var ifres9114 Obj

if True == tmp9220 {
tmp9216 := GHd(__e, V1312)

tmp9217 := GHd(__e, tmp9216)

tmp9218 := GEq(__e, symlist, tmp9217)

var ifres9115 Obj

if True == tmp9218 {
tmp9213 := GHd(__e, V1312)

tmp9214 := GTl(__e, tmp9213)

tmp9215 := GConsP(__e, tmp9214)

var ifres9116 Obj

if True == tmp9215 {
tmp9209 := GHd(__e, V1312)

tmp9210 := GTl(__e, tmp9209)

tmp9211 := GTl(__e, tmp9210)

tmp9212 := GEq(__e, Nil, tmp9211)

var ifres9117 Obj

if True == tmp9212 {
tmp9207 := GTl(__e, V1312)

tmp9208 := GConsP(__e, tmp9207)

var ifres9118 Obj

if True == tmp9208 {
tmp9204 := GTl(__e, V1312)

tmp9205 := GHd(__e, tmp9204)

tmp9206 := GEq(__e, sym_1_1_6, tmp9205)

var ifres9119 Obj

if True == tmp9206 {
tmp9201 := GTl(__e, V1312)

tmp9202 := GTl(__e, tmp9201)

tmp9203 := GConsP(__e, tmp9202)

var ifres9120 Obj

if True == tmp9203 {
tmp9197 := GTl(__e, V1312)

tmp9198 := GTl(__e, tmp9197)

tmp9199 := GHd(__e, tmp9198)

tmp9200 := GConsP(__e, tmp9199)

var ifres9121 Obj

if True == tmp9200 {
tmp9192 := GTl(__e, V1312)

tmp9193 := GTl(__e, tmp9192)

tmp9194 := GHd(__e, tmp9193)

tmp9195 := GHd(__e, tmp9194)

tmp9196 := GEq(__e, symstr, tmp9195)

var ifres9122 Obj

if True == tmp9196 {
tmp9187 := GTl(__e, V1312)

tmp9188 := GTl(__e, tmp9187)

tmp9189 := GHd(__e, tmp9188)

tmp9190 := GTl(__e, tmp9189)

tmp9191 := GConsP(__e, tmp9190)

var ifres9123 Obj

if True == tmp9191 {
tmp9181 := GTl(__e, V1312)

tmp9182 := GTl(__e, tmp9181)

tmp9183 := GHd(__e, tmp9182)

tmp9184 := GTl(__e, tmp9183)

tmp9185 := GHd(__e, tmp9184)

tmp9186 := GConsP(__e, tmp9185)

var ifres9124 Obj

if True == tmp9186 {
tmp9174 := GTl(__e, V1312)

tmp9175 := GTl(__e, tmp9174)

tmp9176 := GHd(__e, tmp9175)

tmp9177 := GTl(__e, tmp9176)

tmp9178 := GHd(__e, tmp9177)

tmp9179 := GHd(__e, tmp9178)

tmp9180 := GEq(__e, symlist, tmp9179)

var ifres9125 Obj

if True == tmp9180 {
tmp9167 := GTl(__e, V1312)

tmp9168 := GTl(__e, tmp9167)

tmp9169 := GHd(__e, tmp9168)

tmp9170 := GTl(__e, tmp9169)

tmp9171 := GHd(__e, tmp9170)

tmp9172 := GTl(__e, tmp9171)

tmp9173 := GConsP(__e, tmp9172)

var ifres9126 Obj

if True == tmp9173 {
tmp9159 := GTl(__e, V1312)

tmp9160 := GTl(__e, tmp9159)

tmp9161 := GHd(__e, tmp9160)

tmp9162 := GTl(__e, tmp9161)

tmp9163 := GHd(__e, tmp9162)

tmp9164 := GTl(__e, tmp9163)

tmp9165 := GTl(__e, tmp9164)

tmp9166 := GEq(__e, Nil, tmp9165)

var ifres9127 Obj

if True == tmp9166 {
tmp9153 := GTl(__e, V1312)

tmp9154 := GTl(__e, tmp9153)

tmp9155 := GHd(__e, tmp9154)

tmp9156 := GTl(__e, tmp9155)

tmp9157 := GTl(__e, tmp9156)

tmp9158 := GConsP(__e, tmp9157)

var ifres9128 Obj

if True == tmp9158 {
tmp9146 := GTl(__e, V1312)

tmp9147 := GTl(__e, tmp9146)

tmp9148 := GHd(__e, tmp9147)

tmp9149 := GTl(__e, tmp9148)

tmp9150 := GTl(__e, tmp9149)

tmp9151 := GTl(__e, tmp9150)

tmp9152 := GEq(__e, Nil, tmp9151)

var ifres9129 Obj

if True == tmp9152 {
tmp9142 := GTl(__e, V1312)

tmp9143 := GTl(__e, tmp9142)

tmp9144 := GTl(__e, tmp9143)

tmp9145 := GEq(__e, Nil, tmp9144)

var ifres9130 Obj

if True == tmp9145 {
tmp9131 := GHd(__e, V1312)

tmp9132 := GTl(__e, tmp9131)

tmp9133 := GHd(__e, tmp9132)

tmp9134 := GTl(__e, V1312)

tmp9135 := GTl(__e, tmp9134)

tmp9136 := GHd(__e, tmp9135)

tmp9137 := GTl(__e, tmp9136)

tmp9138 := GHd(__e, tmp9137)

tmp9139 := GTl(__e, tmp9138)

tmp9140 := GHd(__e, tmp9139)

tmp9141 := GEq(__e, tmp9133, tmp9140)

ifres9130 = tmp9141


} else {
ifres9130 = False


}

ifres9129 = ifres9130


} else {
ifres9129 = False


}

ifres9128 = ifres9129


} else {
ifres9128 = False


}

ifres9127 = ifres9128


} else {
ifres9127 = False


}

ifres9126 = ifres9127


} else {
ifres9126 = False


}

ifres9125 = ifres9126


} else {
ifres9125 = False


}

ifres9124 = ifres9125


} else {
ifres9124 = False


}

ifres9123 = ifres9124


} else {
ifres9123 = False


}

ifres9122 = ifres9123


} else {
ifres9122 = False


}

ifres9121 = ifres9122


} else {
ifres9121 = False


}

ifres9120 = ifres9121


} else {
ifres9120 = False


}

ifres9119 = ifres9120


} else {
ifres9119 = False


}

ifres9118 = ifres9119


} else {
ifres9118 = False


}

ifres9117 = ifres9118


} else {
ifres9117 = False


}

ifres9116 = ifres9117


} else {
ifres9116 = False


}

ifres9115 = ifres9116


} else {
ifres9115 = False


}

ifres9114 = ifres9115


} else {
ifres9114 = False


}

ifres9113 = ifres9114


} else {
ifres9113 = False


}

if True == ifres9113 {
tmp9098 := GTl(__e, V1312)

tmp9099 := GTl(__e, tmp9098)

tmp9100 := GHd(__e, tmp9099)

tmp9101 := GTl(__e, tmp9100)

tmp9102 := GHd(__e, tmp9101)

tmp9103 := GTl(__e, V1312)

tmp9104 := GTl(__e, tmp9103)

tmp9105 := GHd(__e, tmp9104)

tmp9106 := GTl(__e, tmp9105)

tmp9107 := GTl(__e, tmp9106)

tmp9108 := GCons(__e, sym_a_a_6, tmp9107)

return GCons(__e, tmp9102, tmp9108)


} else {
tmp9111 := GConsP(__e, V1312)

if True == tmp9111 {
tmp9109 := MakeNative1(func(__e *ControlFlow, Z1313 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4pretty_1type), Z1313)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp9109, V1312)


} else {
return V1312
}


}


panic("unreachable")
})

tmp9222 := Call2(__e, ns2_1set, symshen_4pretty_1type, tmp9097)


_ = tmp9222

tmp9223 := MakeNative0(func(__e *ControlFlow) Obj {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("type error\n"))
panic("unreachable")
})

tmp9224 := Call2(__e, ns2_1set, symshen_4type_1error, tmp9223)


_ = tmp9224

tmp9225 := MakeNative1(func(__e *ControlFlow, V1314 Obj) Obj {
tmp9226 := Call1(__e, SymFn(symshen_4klfile), V1314)


let__8926 := tmp9226
_ = let__8926

tmp9227 := Call1(__e, SymFn(symread_1file), V1314)


let__8927 := tmp9227
_ = let__8927

tmp9228 := PrimGuard2(__e, symopen, PrimOpenStream, let__8926, symout)

let__8928 := tmp9228
_ = let__8928

tmp9229 := MakeNative1(func(__e *ControlFlow, Z1319 Obj) Obj {
tmp9230 := Call1(__e, SymFn(symshen_4shen_1_6kl_1h), Z1319)


return TailCall1(__e, SymFn(symshen_4partial), tmp9230)


panic("unreachable")
})

tmp9231 := Call2(__e, SymFn(symmap), tmp9229, let__8927)


let__8929 := tmp9231
_ = let__8929

tmp9232 := Call2(__e, SymFn(symshen_4write_1kl), let__8929, let__8928)


let__8930 := tmp9232
_ = let__8930

return let__8926


panic("unreachable")
})

tmp9233 := Call2(__e, ns2_1set, symbootstrap, tmp9225)


_ = tmp9233

tmp9234 := MakeNative1(func(__e *ControlFlow, V1321 Obj) Obj {
tmp9254 := GConsP(__e, V1321)

var ifres9244 Obj

if True == tmp9254 {
tmp9252 := GHd(__e, V1321)

tmp9253 := GEq(__e, symshen_4f_1error, tmp9252)

var ifres9245 Obj

if True == tmp9253 {
tmp9250 := GTl(__e, V1321)

tmp9251 := GConsP(__e, tmp9250)

var ifres9246 Obj

if True == tmp9251 {
tmp9247 := GTl(__e, V1321)

tmp9248 := GTl(__e, tmp9247)

tmp9249 := GEq(__e, Nil, tmp9248)

ifres9246 = tmp9249


} else {
ifres9246 = False


}

ifres9245 = ifres9246


} else {
ifres9245 = False


}

ifres9244 = ifres9245


} else {
ifres9244 = False


}

if True == ifres9244 {
tmp9235 := GTl(__e, V1321)

tmp9236 := GHd(__e, tmp9235)

tmp9237 := PrimGuard1(__e, symstr, PrimStr, tmp9236)

tmp9239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("partial function "), tmp9237), Nil)
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("partial function "), tmp9237)
__typedArg1 := Nil
return Call2(__e, SymFn(symcons), __typedArg0, __typedArg1)
})()

return GCons(__e, symsimple_1error, tmp9239)


} else {
tmp9242 := GConsP(__e, V1321)

if True == tmp9242 {
tmp9240 := MakeNative1(func(__e *ControlFlow, Z1322 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4partial), Z1322)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp9240, V1321)


} else {
return V1321
}


}


panic("unreachable")
})

tmp9255 := Call2(__e, ns2_1set, symshen_4partial, tmp9234)


_ = tmp9255

tmp9256 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V1325 := __self1
_ = V1325
V1326 := __self2
_ = V1326
tmp9269 := GEq(__e, Nil, V1325)

if True == tmp9269 {
return PrimGuard1(__e, symclose, PrimCloseStream, V1326)
} else {
tmp9267 := GConsP(__e, V1325)

var ifres9264 Obj

if True == tmp9267 {
tmp9265 := GHd(__e, V1325)

tmp9266 := GConsP(__e, tmp9265)

ifres9264 = tmp9266


} else {
ifres9264 = False


}

if True == ifres9264 {
tmp9257 := GTl(__e, V1325)

tmp9258 := GHd(__e, V1325)

tmp9259 := Call2(__e, SymFn(symshen_4write_1kl_1h), tmp9258, V1326)


_ = tmp9259

if SymFn(symshen_4write_1kl) == __self {
__self1, __self2 = tmp9257, V1326
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4write_1kl), tmp9257, V1326)


} else {
tmp9262 := GConsP(__e, V1325)

if True == tmp9262 {
tmp9260 := GTl(__e, V1325)

if SymFn(symshen_4write_1kl) == __self {
__self1, __self2 = tmp9260, V1326
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4write_1kl), tmp9260, V1326)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4write_1kl)
}


}


}


panic("unreachable")
})
return __self
})()

tmp9270 := Call2(__e, ns2_1set, symshen_4write_1kl, tmp9256)


_ = tmp9270

tmp9271 := MakeNative2(func(__e *ControlFlow, V1329, V1330 Obj) Obj {
tmp9304 := GConsP(__e, V1329)

var ifres9274 Obj

if True == tmp9304 {
tmp9302 := GHd(__e, V1329)

tmp9303 := GEq(__e, symdefun, tmp9302)

var ifres9275 Obj

if True == tmp9303 {
tmp9300 := GTl(__e, V1329)

tmp9301 := GConsP(__e, tmp9300)

var ifres9276 Obj

if True == tmp9301 {
tmp9297 := GTl(__e, V1329)

tmp9298 := GHd(__e, tmp9297)

tmp9299 := GEq(__e, symfail, tmp9298)

var ifres9277 Obj

if True == tmp9299 {
tmp9294 := GTl(__e, V1329)

tmp9295 := GTl(__e, tmp9294)

tmp9296 := GConsP(__e, tmp9295)

var ifres9278 Obj

if True == tmp9296 {
tmp9290 := GTl(__e, V1329)

tmp9291 := GTl(__e, tmp9290)

tmp9292 := GHd(__e, tmp9291)

tmp9293 := GEq(__e, Nil, tmp9292)

var ifres9279 Obj

if True == tmp9293 {
tmp9286 := GTl(__e, V1329)

tmp9287 := GTl(__e, tmp9286)

tmp9288 := GTl(__e, tmp9287)

tmp9289 := GConsP(__e, tmp9288)

var ifres9280 Obj

if True == tmp9289 {
tmp9281 := GTl(__e, V1329)

tmp9282 := GTl(__e, tmp9281)

tmp9283 := GTl(__e, tmp9282)

tmp9284 := GTl(__e, tmp9283)

tmp9285 := GEq(__e, Nil, tmp9284)

ifres9280 = tmp9285


} else {
ifres9280 = False


}

ifres9279 = ifres9280


} else {
ifres9279 = False


}

ifres9278 = ifres9279


} else {
ifres9278 = False


}

ifres9277 = ifres9278


} else {
ifres9277 = False


}

ifres9276 = ifres9277


} else {
ifres9276 = False


}

ifres9275 = ifres9276


} else {
ifres9275 = False


}

ifres9274 = ifres9275


} else {
ifres9274 = False


}

if True == ifres9274 {
return TailCall2(__e, SymFn(sympr), MakeString("(defun fail () shen.fail!)"), V1330)
} else {
tmp9272 := Call3(__e, SymFn(symshen_4app), V1329, MakeString("\n\n"), symshen_4r)


return TailCall2(__e, SymFn(sympr), tmp9272, V1330)


}


panic("unreachable")
})

tmp9305 := Call2(__e, ns2_1set, symshen_4write_1kl_1h, tmp9271)


_ = tmp9305

tmp9306 := MakeNative1(func(__e *ControlFlow, V1331 Obj) Obj {
tmp9315 := GEq(__e, MakeString(""), V1331)

if True == tmp9315 {
return MakeString(".kl")
} else {
tmp9313 := GEq(__e, MakeString(".shen"), V1331)

if True == tmp9313 {
return MakeString(".kl")
} else {
tmp9311 := Call1(__e, SymFn(symshen_4_7string_2), V1331)


if True == tmp9311 {
tmp9307 := Call1(__e, SymFn(symhdstr), V1331)


tmp9309 := Call1(__e, SymFn(symshen_4klfile), PrimGuard1(__e, symtlstr, PrimTailString, V1331))


return TailCall2(__e, SymFn(sym_8s), tmp9307, tmp9309)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4klfile)
}


}


}


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4klfile, tmp9306)




panic("unreachable")
})

