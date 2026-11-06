package main

import . "github.com/pyrex41/shen-go/kl"

var DeclarationsMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp8062 := MakeNative0(func(__e *ControlFlow) Obj {
tmp8063 := MakeNative1(func(__e *ControlFlow, Z920 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4typename), Z920)
panic("unreachable")
})

tmp8064 := GValue(__e, symshen_4_dalldatatypes_d)

return TailCall2(__e, SymFn(symmap), tmp8063, tmp8064)


panic("unreachable")
})

tmp8065 := Call2(__e, ns2_1set, symdatatypes, tmp8062)


_ = tmp8065

tmp8066 := MakeNative0(func(__e *ControlFlow) Obj {
tmp8067 := MakeNative1(func(__e *ControlFlow, Z921 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4typename), Z921)
panic("unreachable")
})

tmp8068 := GValue(__e, symshen_4_ddatatypes_d)

return TailCall2(__e, SymFn(symmap), tmp8067, tmp8068)


panic("unreachable")
})

tmp8069 := Call2(__e, ns2_1set, symincluded, tmp8066)


_ = tmp8069

tmp8070 := MakeNative1(func(__e *ControlFlow, V924 Obj) Obj {
tmp8075 := GConsP(__e, V924)

if True == tmp8075 {
tmp8071 := GHd(__e, V924)

tmp8072 := PrimGuard1(__e, symstr, PrimStr, tmp8071)

tmp8073 := Call1(__e, SymFn(symshen_4typename_1h), tmp8072)


return PrimGuard1(__e, symintern, PrimIntern, tmp8073)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4typename)
}


panic("unreachable")
})

tmp8076 := Call2(__e, ns2_1set, symshen_4typename, tmp8070)


_ = tmp8076

tmp8077 := MakeNative1(func(__e *ControlFlow, V925 Obj) Obj {
tmp8084 := GEq(__e, MakeString("#type"), V925)

if True == tmp8084 {
return MakeString("")
} else {
tmp8082 := Call1(__e, SymFn(symshen_4_7string_2), V925)


if True == tmp8082 {
tmp8078 := Call1(__e, SymFn(symhdstr), V925)


tmp8080 := Call1(__e, SymFn(symshen_4typename_1h), PrimGuard1(__e, symtlstr, PrimTailString, V925))


return PrimGuard2(__e, symcn, PrimStringConcat, tmp8078, tmp8080)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4typename_1h)
}


}


panic("unreachable")
})

tmp8085 := Call2(__e, ns2_1set, symshen_4typename_1h, tmp8077)


_ = tmp8085

tmp8086 := MakeNative1(func(__e *ControlFlow, V926 Obj) Obj {
if True == GLt(__e, V926, MakeInteger(0)) {
return GValue(__e, symshen_4_dprolog_1memory_d)
} else {
tmp8088 := PrimGuard1(__e, syminteger_2, PrimIsInteger, V926)

if True == tmp8088 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dprolog_1memory_d, V926)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("prolog memory expects an integer value\n"))
}


}


panic("unreachable")
})

tmp8091 := Call2(__e, ns2_1set, symprolog_1memory, tmp8086)


_ = tmp8091

tmp8092 := MakeNative1(func(__e *ControlFlow, V927 Obj) Obj {
tmp8093 := MakeNative0(func(__e *ControlFlow) Obj {
tmp8094 := GValue(__e, sym_dproperty_1vector_d)

return TailCall3(__e, SymFn(symget), V927, symarity, tmp8094)


panic("unreachable")
})

tmp8095 := MakeNative1(func(__e *ControlFlow, Z928 Obj) Obj {
return MakeInteger(-1)
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp8093, tmp8095)


panic("unreachable")
})

tmp8096 := Call2(__e, ns2_1set, symarity, tmp8092)


_ = tmp8096

tmp8097 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V931 := __self1
_ = V931
tmp8112 := GEq(__e, Nil, V931)

if True == tmp8112 {
return Nil
} else {
tmp8110 := GConsP(__e, V931)

var ifres8106 Obj

if True == tmp8110 {
tmp8108 := GTl(__e, V931)

tmp8109 := GConsP(__e, tmp8108)

var ifres8107 Obj

if True == tmp8109 {
ifres8107 = True


} else {
ifres8107 = False


}

ifres8106 = ifres8107


} else {
ifres8106 = False


}

if True == ifres8106 {
tmp8098 := GHd(__e, V931)

tmp8099 := GTl(__e, V931)

tmp8100 := GHd(__e, tmp8099)

tmp8101 := GValue(__e, sym_dproperty_1vector_d)

tmp8102 := Call4(__e, SymFn(symput), tmp8098, symarity, tmp8100, tmp8101)


let__8057 := tmp8102
_ = let__8057

tmp8103 := GTl(__e, V931)

tmp8104 := GTl(__e, tmp8103)

if SymFn(symshen_4initialise_1arity_1table) == __self {
__self1 = tmp8104
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4initialise_1arity_1table), tmp8104)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.initialise-arity-table"))
}


}


panic("unreachable")
})
return __self
})()

tmp8113 := Call2(__e, ns2_1set, symshen_4initialise_1arity_1table, tmp8097)


_ = tmp8113

tmp8114 := MakeNative1(func(__e *ControlFlow, V933 Obj) Obj {
tmp8115 := GValue(__e, sym_dproperty_1vector_d)

tmp8116 := Call3(__e, SymFn(symget), symshen, symshen_4external_1symbols, tmp8115)


let__8058 := tmp8116
_ = let__8058

tmp8117 := Call2(__e, SymFn(symadjoin), V933, let__8058)


tmp8118 := GValue(__e, sym_dproperty_1vector_d)

tmp8119 := Call4(__e, SymFn(symput), symshen, symshen_4external_1symbols, tmp8117, tmp8118)


let__8059 := tmp8119
_ = let__8059

return V933


panic("unreachable")
})

tmp8120 := Call2(__e, ns2_1set, symsystemf, tmp8114)


_ = tmp8120

tmp8121 := MakeNative2(func(__e *ControlFlow, V936, V937 Obj) Obj {
tmp8123 := Call2(__e, SymFn(symelement_2), V936, V937)


if True == tmp8123 {
return V937
} else {
return GCons(__e, V936, V937)
}


panic("unreachable")
})

tmp8124 := Call2(__e, ns2_1set, symadjoin, tmp8121)


_ = tmp8124

tmp8125 := MakeNative1(func(__e *ControlFlow, V938 Obj) Obj {
tmp8126 := Call1(__e, SymFn(symarity), V938)


let__8060 := tmp8126
_ = let__8060

tmp8134 := GEq(__e, let__8060, MakeInteger(-1))

var ifres8131 Obj

if True == tmp8134 {
ifres8131 = True


} else {
tmp8133 := GEq(__e, let__8060, MakeInteger(0))

var ifres8132 Obj

if True == tmp8133 {
ifres8132 = True


} else {
ifres8132 = False


}

ifres8131 = ifres8132


}

if True == ifres8131 {
return Nil
} else {
tmp8127 := GCons(__e, V938, Nil)

tmp8128 := Call2(__e, SymFn(symshen_4lambda_1function), tmp8127, let__8060)


tmp8129 := Call1(__e, SymFn(symeval_1kl), tmp8128)


return GCons(__e, V938, tmp8129)


}


panic("unreachable")
})

tmp8135 := Call2(__e, ns2_1set, symshen_4lambda_1entry, tmp8125)


_ = tmp8135

tmp8136 := MakeNative1(func(__e *ControlFlow, V940 Obj) Obj {
tmp8141 := GConsP(__e, V940)

if True == tmp8141 {
tmp8137 := GHd(__e, V940)

tmp8138 := GTl(__e, V940)

tmp8139 := GValue(__e, sym_dproperty_1vector_d)

return TailCall4(__e, SymFn(symput), tmp8137, symshen_4lambda_1form, tmp8138, tmp8139)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4set_1lambda_1form_1entry)
}


panic("unreachable")
})

tmp8142 := Call2(__e, ns2_1set, symshen_4set_1lambda_1form_1entry, tmp8136)


_ = tmp8142

tmp8143 := MakeNative1(func(__e *ControlFlow, V941 Obj) Obj {
tmp8144 := MakeNative1(func(__e *ControlFlow, Z943 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4lambda_1entry), Z943)
panic("unreachable")
})

tmp8145 := Call2(__e, SymFn(symmap), tmp8144, V941)


let__8061 := tmp8145
_ = let__8061

tmp8146 := MakeNative1(func(__e *ControlFlow, Z944 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4set_1lambda_1form_1entry), Z944)
panic("unreachable")
})

tmp8147 := MakeNative1(func(__e *ControlFlow, Z945 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z946 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4register_1source_1form), Z945, Z946)
panic("unreachable")
})
panic("unreachable")
})

tmp8148 := GCons(__e, symshen_4register_1source_1form, tmp8147)

tmp8149 := MakeNative1(func(__e *ControlFlow, Z947 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4unregister_1source_1form), Z947)
panic("unreachable")
})

tmp8150 := GCons(__e, symshen_4unregister_1source_1form, tmp8149)

tmp8151 := MakeNative1(func(__e *ControlFlow, Z948 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4tuple), Z948)
panic("unreachable")
})

tmp8152 := GCons(__e, symshen_4tuple, tmp8151)

tmp8153 := MakeNative1(func(__e *ControlFlow, Z949 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4pvar), Z949)
panic("unreachable")
})

tmp8154 := GCons(__e, symshen_4pvar, tmp8153)

tmp8155 := MakeNative1(func(__e *ControlFlow, Z950 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4dictionary), Z950)
panic("unreachable")
})

tmp8156 := GCons(__e, symshen_4dictionary, tmp8155)

tmp8157 := MakeNative1(func(__e *ControlFlow, Z951 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4print_1prolog_1vector), Z951)
panic("unreachable")
})

tmp8158 := GCons(__e, symshen_4print_1prolog_1vector, tmp8157)

tmp8159 := MakeNative1(func(__e *ControlFlow, Z952 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4print_1freshterm), Z952)
panic("unreachable")
})

tmp8160 := GCons(__e, symshen_4print_1freshterm, tmp8159)

tmp8161 := MakeNative1(func(__e *ControlFlow, Z953 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4printF), Z953)
panic("unreachable")
})

tmp8162 := GCons(__e, symshen_4printF, tmp8161)

tmp8163 := GCons(__e, tmp8162, let__8061)

tmp8164 := GCons(__e, tmp8160, tmp8163)

tmp8165 := GCons(__e, tmp8158, tmp8164)

tmp8166 := GCons(__e, tmp8156, tmp8165)

tmp8167 := GCons(__e, tmp8154, tmp8166)

tmp8168 := GCons(__e, tmp8152, tmp8167)

tmp8169 := GCons(__e, tmp8150, tmp8168)

tmp8170 := GCons(__e, tmp8148, tmp8169)

return TailCall2(__e, SymFn(symshen_4for_1each), tmp8146, tmp8170)


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4build_1lambda_1table, tmp8143)




panic("unreachable")
})

