package main

import . "github.com/pyrex41/shen-go/kl"

var DeclarationsMain = MakeNative(func(__e *ControlFlow) {
tmp8082 := MakeNative(func(__e *ControlFlow) {
tmp8083 := MakeNative(func(__e *ControlFlow) {
Z920 := __e.Get(1)
_ = Z920
__e.TailApply(PrimFunc(symshen_4typename), Z920)
return
}, 1)

tmp8084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dalldatatypes_d)
}
__typedArg0 := symshen_4_dalldatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symmap), tmp8083, tmp8084)
return


}, 0)

tmp8085 := Call(__e, ns2_1set, symdatatypes, tmp8082)


_ = tmp8085

tmp8086 := MakeNative(func(__e *ControlFlow) {
tmp8087 := MakeNative(func(__e *ControlFlow) {
Z921 := __e.Get(1)
_ = Z921
__e.TailApply(PrimFunc(symshen_4typename), Z921)
return
}, 1)

tmp8088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symmap), tmp8087, tmp8088)
return


}, 0)

tmp8089 := Call(__e, ns2_1set, symincluded, tmp8086)


_ = tmp8089

tmp8090 := MakeNative(func(__e *ControlFlow) {
V924 := __e.Get(1)
_ = V924
tmp8095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V924)
}
__typedArg0 := V924
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8095 {
tmp8091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V924)
}
__typedArg0 := V924
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(tmp8091)
}
__typedArg0 := tmp8091
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp8093 := Call(__e, PrimFunc(symshen_4typename_1h), tmp8092)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(tmp8093)
}
__typedArg0 := tmp8093
return Call(__e, PrimFunc(symintern), __typedArg0)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4typename)
return
}


}, 1)

tmp8096 := Call(__e, ns2_1set, symshen_4typename, tmp8090)


_ = tmp8096

tmp8097 := MakeNative(func(__e *ControlFlow) {
V925 := __e.Get(1)
_ = V925
tmp8104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("#type"), V925)
}
__typedArg0 := MakeString("#type")
__typedArg1 := V925
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8104 {
__e.Return(MakeString(""))
return
} else {
tmp8102 := Call(__e, PrimFunc(symshen_4_7string_2), V925)


if True == tmp8102 {
tmp8098 := Call(__e, PrimFunc(symhdstr), V925)


tmp8100 := Call(__e, PrimFunc(symshen_4typename_1h), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V925)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V925
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp8098)
__typedS1, __typedOK1 := TypedString(tmp8100)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp8098
__typedArg1 := tmp8100
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4typename_1h)
return
}


}


}, 1)

tmp8105 := Call(__e, ns2_1set, symshen_4typename_1h, tmp8097)


_ = tmp8105

tmp8106 := MakeNative(func(__e *ControlFlow) {
V926 := __e.Get(1)
_ = V926
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5) {
__typedN0, __typedOK0 := TypedFloat64(V926)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(0))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5) {
return TypedMaterializeBoolean((__typedN0 < __typedN1))
}}
__typedArg0 := V926
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_5), __typedArg0, __typedArg1)
})() {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dprolog_1memory_d)
}
__typedArg0 := symshen_4_dprolog_1memory_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
} else {
tmp8108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V926)
}
__typedArg0 := V926
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})()

if True == tmp8108 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dprolog_1memory_d, V926)
}
__typedArg0 := symshen_4_dprolog_1memory_d
__typedArg1 := V926
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("prolog memory expects an integer value\n"))
}
__typedArg0 := MakeString("prolog memory expects an integer value\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp8111 := Call(__e, ns2_1set, symprolog_1memory, tmp8106)


_ = tmp8111

tmp8112 := MakeNative(func(__e *ControlFlow) {
V927 := __e.Get(1)
_ = V927
tmp8113 := MakeNative(func(__e *ControlFlow) {
tmp8114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V927, symarity, tmp8114)
return


}, 0)

tmp8115 := MakeNative(func(__e *ControlFlow) {
Z928 := __e.Get(1)
_ = Z928
__e.Return(MakeInteger(-1))
return
}, 1)

__e.TailApply(try_1catch, tmp8113, tmp8115)
return


}, 1)

tmp8116 := Call(__e, ns2_1set, symarity, tmp8112)


_ = tmp8116

tmp8117 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V931 := __self1
_ = V931
tmp8132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V931)
}
__typedArg0 := Nil
__typedArg1 := V931
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8132 {
__e.Return(Nil)
return
} else {
tmp8130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V931)
}
__typedArg0 := V931
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8126 Obj

if True == tmp8130 {
tmp8128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V931)
}
__typedArg0 := V931
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8128)
}
__typedArg0 := tmp8128
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8127 Obj

if True == tmp8129 {
ifres8127 = True


} else {
ifres8127 = False


}

ifres8126 = ifres8127


} else {
ifres8126 = False


}

if True == ifres8126 {
tmp8118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V931)
}
__typedArg0 := V931
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V931)
}
__typedArg0 := V931
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8119)
}
__typedArg0 := tmp8119
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8122 := Call(__e, PrimFunc(symput), tmp8118, symarity, tmp8120, tmp8121)


let__8077 := tmp8122
_ = let__8077

tmp8123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V931)
}
__typedArg0 := V931
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp8123)
}
__typedArg0 := tmp8123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4initialise_1arity_1table) == __self {
__self1 = tmp8124
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4initialise_1arity_1table), tmp8124)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.initialise-arity-table"))
}
__typedArg0 := MakeString("implementation error in shen.initialise-arity-table")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp8133 := Call(__e, ns2_1set, symshen_4initialise_1arity_1table, tmp8117)


_ = tmp8133

tmp8134 := MakeNative(func(__e *ControlFlow) {
V933 := __e.Get(1)
_ = V933
tmp8135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8136 := Call(__e, PrimFunc(symget), symshen, symshen_4external_1symbols, tmp8135)


let__8078 := tmp8136
_ = let__8078

tmp8137 := Call(__e, PrimFunc(symadjoin), V933, let__8078)


tmp8138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8139 := Call(__e, PrimFunc(symput), symshen, symshen_4external_1symbols, tmp8137, tmp8138)


let__8079 := tmp8139
_ = let__8079

__e.Return(V933)
return


}, 1)

tmp8140 := Call(__e, ns2_1set, symsystemf, tmp8134)


_ = tmp8140

tmp8141 := MakeNative(func(__e *ControlFlow) {
V936 := __e.Get(1)
_ = V936
V937 := __e.Get(2)
_ = V937
tmp8143 := Call(__e, PrimFunc(symelement_2), V936, V937)


if True == tmp8143 {
__e.Return(V937)
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V936, V937)
}
__typedArg0 := V936
__typedArg1 := V937
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
}


}, 2)

tmp8144 := Call(__e, ns2_1set, symadjoin, tmp8141)


_ = tmp8144

tmp8145 := MakeNative(func(__e *ControlFlow) {
V938 := __e.Get(1)
_ = V938
tmp8146 := Call(__e, PrimFunc(symarity), V938)


let__8080 := tmp8146
_ = let__8080

tmp8154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__8080, MakeInteger(-1))
}
__typedArg0 := let__8080
__typedArg1 := MakeInteger(-1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8151 Obj

if True == tmp8154 {
ifres8151 = True


} else {
tmp8153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__8080, MakeInteger(0))
}
__typedArg0 := let__8080
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8152 Obj

if True == tmp8153 {
ifres8152 = True


} else {
ifres8152 = False


}

ifres8151 = ifres8152


}

if True == ifres8151 {
__e.Return(Nil)
return
} else {
tmp8147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V938, Nil)
}
__typedArg0 := V938
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8148 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp8147, let__8080)


tmp8149 := Call(__e, PrimFunc(symeval_1kl), tmp8148)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V938, tmp8149)
}
__typedArg0 := V938
__typedArg1 := tmp8149
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp8155 := Call(__e, ns2_1set, symshen_4lambda_1entry, tmp8145)


_ = tmp8155

tmp8156 := MakeNative(func(__e *ControlFlow) {
V940 := __e.Get(1)
_ = V940
tmp8161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V940)
}
__typedArg0 := V940
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8161 {
tmp8157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V940)
}
__typedArg0 := V940
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V940)
}
__typedArg0 := V940
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), tmp8157, symshen_4lambda_1form, tmp8158, tmp8159)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4set_1lambda_1form_1entry)
return
}


}, 1)

tmp8162 := Call(__e, ns2_1set, symshen_4set_1lambda_1form_1entry, tmp8156)


_ = tmp8162

tmp8163 := MakeNative(func(__e *ControlFlow) {
V941 := __e.Get(1)
_ = V941
tmp8164 := MakeNative(func(__e *ControlFlow) {
Z943 := __e.Get(1)
_ = Z943
__e.TailApply(PrimFunc(symshen_4lambda_1entry), Z943)
return
}, 1)

tmp8165 := Call(__e, PrimFunc(symmap), tmp8164, V941)


let__8081 := tmp8165
_ = let__8081

tmp8166 := MakeNative(func(__e *ControlFlow) {
Z944 := __e.Get(1)
_ = Z944
__e.TailApply(PrimFunc(symshen_4set_1lambda_1form_1entry), Z944)
return
}, 1)

tmp8167 := MakeNative(func(__e *ControlFlow) {
Z945 := __e.Get(1)
_ = Z945
__e.Return(MakeNative(func(__e *ControlFlow) {
Z946 := __e.Get(1)
_ = Z946
__e.TailApply(PrimFunc(symshen_4register_1source_1form), Z945, Z946)
return
}, 1))
return
}, 1)

tmp8168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4register_1source_1form, tmp8167)
}
__typedArg0 := symshen_4register_1source_1form
__typedArg1 := tmp8167
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8169 := MakeNative(func(__e *ControlFlow) {
Z947 := __e.Get(1)
_ = Z947
__e.TailApply(PrimFunc(symshen_4unregister_1source_1form), Z947)
return
}, 1)

tmp8170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4unregister_1source_1form, tmp8169)
}
__typedArg0 := symshen_4unregister_1source_1form
__typedArg1 := tmp8169
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8171 := MakeNative(func(__e *ControlFlow) {
Z948 := __e.Get(1)
_ = Z948
__e.TailApply(PrimFunc(symshen_4tuple), Z948)
return
}, 1)

tmp8172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4tuple, tmp8171)
}
__typedArg0 := symshen_4tuple
__typedArg1 := tmp8171
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8173 := MakeNative(func(__e *ControlFlow) {
Z949 := __e.Get(1)
_ = Z949
__e.TailApply(PrimFunc(symshen_4pvar), Z949)
return
}, 1)

tmp8174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4pvar, tmp8173)
}
__typedArg0 := symshen_4pvar
__typedArg1 := tmp8173
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8175 := MakeNative(func(__e *ControlFlow) {
Z950 := __e.Get(1)
_ = Z950
__e.TailApply(PrimFunc(symshen_4dictionary), Z950)
return
}, 1)

tmp8176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4dictionary, tmp8175)
}
__typedArg0 := symshen_4dictionary
__typedArg1 := tmp8175
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8177 := MakeNative(func(__e *ControlFlow) {
Z951 := __e.Get(1)
_ = Z951
__e.TailApply(PrimFunc(symshen_4print_1prolog_1vector), Z951)
return
}, 1)

tmp8178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1prolog_1vector, tmp8177)
}
__typedArg0 := symshen_4print_1prolog_1vector
__typedArg1 := tmp8177
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8179 := MakeNative(func(__e *ControlFlow) {
Z952 := __e.Get(1)
_ = Z952
__e.TailApply(PrimFunc(symshen_4print_1freshterm), Z952)
return
}, 1)

tmp8180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1freshterm, tmp8179)
}
__typedArg0 := symshen_4print_1freshterm
__typedArg1 := tmp8179
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8181 := MakeNative(func(__e *ControlFlow) {
Z953 := __e.Get(1)
_ = Z953
__e.TailApply(PrimFunc(symshen_4printF), Z953)
return
}, 1)

tmp8182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4printF, tmp8181)
}
__typedArg0 := symshen_4printF
__typedArg1 := tmp8181
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8182, let__8081)
}
__typedArg0 := tmp8182
__typedArg1 := let__8081
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8180, tmp8183)
}
__typedArg0 := tmp8180
__typedArg1 := tmp8183
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8178, tmp8184)
}
__typedArg0 := tmp8178
__typedArg1 := tmp8184
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8176, tmp8185)
}
__typedArg0 := tmp8176
__typedArg1 := tmp8185
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8174, tmp8186)
}
__typedArg0 := tmp8174
__typedArg1 := tmp8186
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8172, tmp8187)
}
__typedArg0 := tmp8172
__typedArg1 := tmp8187
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8170, tmp8188)
}
__typedArg0 := tmp8170
__typedArg1 := tmp8188
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8168, tmp8189)
}
__typedArg0 := tmp8168
__typedArg1 := tmp8189
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4for_1each), tmp8166, tmp8190)
return


}, 1)

__e.TailApply(ns2_1set, symshen_4build_1lambda_1table, tmp8163)
return




}, 0)

