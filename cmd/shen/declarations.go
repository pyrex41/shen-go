package main

import . "github.com/pyrex41/shen-go/kl"

var DeclarationsMain = MakeNative(func(__e *ControlFlow) {
tmp8062 := MakeNative(func(__e *ControlFlow) {
tmp8063 := MakeNative(func(__e *ControlFlow) {
Z920 := __e.Get(1)
_ = Z920
__e.TailApply(PrimFunc(symshen_4typename), Z920)
return
}, 1)

tmp8064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dalldatatypes_d)
}
__typedArg0 := symshen_4_dalldatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symmap), tmp8063, tmp8064)
return


}, 0)

tmp8065 := Call(__e, ns2_1set, symdatatypes, tmp8062)


_ = tmp8065

tmp8066 := MakeNative(func(__e *ControlFlow) {
tmp8067 := MakeNative(func(__e *ControlFlow) {
Z921 := __e.Get(1)
_ = Z921
__e.TailApply(PrimFunc(symshen_4typename), Z921)
return
}, 1)

tmp8068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symmap), tmp8067, tmp8068)
return


}, 0)

tmp8069 := Call(__e, ns2_1set, symshen_4included, tmp8066)


_ = tmp8069

tmp8070 := MakeNative(func(__e *ControlFlow) {
V924 := __e.Get(1)
_ = V924
tmp8075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V924)
}
__typedArg0 := V924
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8075 {
tmp8071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V924)
}
__typedArg0 := V924
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(tmp8071)
}
__typedArg0 := tmp8071
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp8073 := Call(__e, PrimFunc(symshen_4typename_1h), tmp8072)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(tmp8073)
}
__typedArg0 := tmp8073
return Call(__e, PrimFunc(symintern), __typedArg0)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4typename)
return
}


}, 1)

tmp8076 := Call(__e, ns2_1set, symshen_4typename, tmp8070)


_ = tmp8076

tmp8077 := MakeNative(func(__e *ControlFlow) {
V925 := __e.Get(1)
_ = V925
tmp8084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("#type"), V925)
}
__typedArg0 := MakeString("#type")
__typedArg1 := V925
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8084 {
__e.Return(MakeString(""))
return
} else {
tmp8082 := Call(__e, PrimFunc(symshen_4_7string_2), V925)


if True == tmp8082 {
tmp8078 := Call(__e, PrimFunc(symhdstr), V925)


tmp8080 := Call(__e, PrimFunc(symshen_4typename_1h), (func() Obj {
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
__typedS0, __typedOK0 := TypedString(tmp8078)
__typedS1, __typedOK1 := TypedString(tmp8080)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp8078
__typedArg1 := tmp8080
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4typename_1h)
return
}


}


}, 1)

tmp8085 := Call(__e, ns2_1set, symshen_4typename_1h, tmp8077)


_ = tmp8085

tmp8086 := MakeNative(func(__e *ControlFlow) {
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
tmp8088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V926)
}
__typedArg0 := V926
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})()

if True == tmp8088 {
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

tmp8091 := Call(__e, ns2_1set, symprolog_1memory, tmp8086)


_ = tmp8091

tmp8092 := MakeNative(func(__e *ControlFlow) {
V927 := __e.Get(1)
_ = V927
tmp8093 := MakeNative(func(__e *ControlFlow) {
tmp8094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V927, symarity, tmp8094)
return


}, 0)

tmp8095 := MakeNative(func(__e *ControlFlow) {
Z928 := __e.Get(1)
_ = Z928
__e.Return(MakeInteger(-1))
return
}, 1)

__e.TailApply(try_1catch, tmp8093, tmp8095)
return


}, 1)

tmp8096 := Call(__e, ns2_1set, symarity, tmp8092)


_ = tmp8096

tmp8097 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V931 := __self1
_ = V931
tmp8112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V931)
}
__typedArg0 := Nil
__typedArg1 := V931
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8112 {
__e.Return(Nil)
return
} else {
tmp8110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V931)
}
__typedArg0 := V931
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8106 Obj

if True == tmp8110 {
tmp8108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V931)
}
__typedArg0 := V931
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8108)
}
__typedArg0 := tmp8108
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

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
tmp8098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V931)
}
__typedArg0 := V931
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V931)
}
__typedArg0 := V931
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8099)
}
__typedArg0 := tmp8099
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8102 := Call(__e, PrimFunc(symput), tmp8098, symarity, tmp8100, tmp8101)


let__8057 := tmp8102
_ = let__8057

tmp8103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V931)
}
__typedArg0 := V931
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp8103)
}
__typedArg0 := tmp8103
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4initialise_1arity_1table) == __self {
__self1 = tmp8104
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4initialise_1arity_1table), tmp8104)
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

tmp8113 := Call(__e, ns2_1set, symshen_4initialise_1arity_1table, tmp8097)


_ = tmp8113

tmp8114 := MakeNative(func(__e *ControlFlow) {
V933 := __e.Get(1)
_ = V933
tmp8115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8116 := Call(__e, PrimFunc(symget), symshen, symshen_4external_1symbols, tmp8115)


let__8058 := tmp8116
_ = let__8058

tmp8117 := Call(__e, PrimFunc(symadjoin), V933, let__8058)


tmp8118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp8119 := Call(__e, PrimFunc(symput), symshen, symshen_4external_1symbols, tmp8117, tmp8118)


let__8059 := tmp8119
_ = let__8059

__e.Return(V933)
return


}, 1)

tmp8120 := Call(__e, ns2_1set, symsystemf, tmp8114)


_ = tmp8120

tmp8121 := MakeNative(func(__e *ControlFlow) {
V936 := __e.Get(1)
_ = V936
V937 := __e.Get(2)
_ = V937
tmp8123 := Call(__e, PrimFunc(symelement_2), V936, V937)


if True == tmp8123 {
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

tmp8124 := Call(__e, ns2_1set, symadjoin, tmp8121)


_ = tmp8124

tmp8125 := MakeNative(func(__e *ControlFlow) {
V938 := __e.Get(1)
_ = V938
tmp8126 := Call(__e, PrimFunc(symarity), V938)


let__8060 := tmp8126
_ = let__8060

tmp8134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__8060, MakeInteger(-1))
}
__typedArg0 := let__8060
__typedArg1 := MakeInteger(-1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8131 Obj

if True == tmp8134 {
ifres8131 = True


} else {
tmp8133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__8060, MakeInteger(0))
}
__typedArg0 := let__8060
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8132 Obj

if True == tmp8133 {
ifres8132 = True


} else {
ifres8132 = False


}

ifres8131 = ifres8132


}

if True == ifres8131 {
__e.Return(Nil)
return
} else {
tmp8127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V938, Nil)
}
__typedArg0 := V938
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8128 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp8127, let__8060)


tmp8129 := Call(__e, PrimFunc(symeval_1kl), tmp8128)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V938, tmp8129)
}
__typedArg0 := V938
__typedArg1 := tmp8129
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp8135 := Call(__e, ns2_1set, symshen_4lambda_1entry, tmp8125)


_ = tmp8135

tmp8136 := MakeNative(func(__e *ControlFlow) {
V940 := __e.Get(1)
_ = V940
tmp8141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V940)
}
__typedArg0 := V940
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp8141 {
tmp8137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V940)
}
__typedArg0 := V940
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V940)
}
__typedArg0 := V940
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), tmp8137, symshen_4lambda_1form, tmp8138, tmp8139)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4set_1lambda_1form_1entry)
return
}


}, 1)

tmp8142 := Call(__e, ns2_1set, symshen_4set_1lambda_1form_1entry, tmp8136)


_ = tmp8142

tmp8143 := MakeNative(func(__e *ControlFlow) {
V941 := __e.Get(1)
_ = V941
tmp8144 := MakeNative(func(__e *ControlFlow) {
Z943 := __e.Get(1)
_ = Z943
__e.TailApply(PrimFunc(symshen_4lambda_1entry), Z943)
return
}, 1)

tmp8145 := Call(__e, PrimFunc(symmap), tmp8144, V941)


let__8061 := tmp8145
_ = let__8061

tmp8146 := MakeNative(func(__e *ControlFlow) {
Z944 := __e.Get(1)
_ = Z944
__e.TailApply(PrimFunc(symshen_4set_1lambda_1form_1entry), Z944)
return
}, 1)

tmp8147 := MakeNative(func(__e *ControlFlow) {
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

tmp8148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4register_1source_1form, tmp8147)
}
__typedArg0 := symshen_4register_1source_1form
__typedArg1 := tmp8147
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8149 := MakeNative(func(__e *ControlFlow) {
Z947 := __e.Get(1)
_ = Z947
__e.TailApply(PrimFunc(symshen_4unregister_1source_1form), Z947)
return
}, 1)

tmp8150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4unregister_1source_1form, tmp8149)
}
__typedArg0 := symshen_4unregister_1source_1form
__typedArg1 := tmp8149
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8151 := MakeNative(func(__e *ControlFlow) {
Z948 := __e.Get(1)
_ = Z948
__e.TailApply(PrimFunc(symshen_4tuple), Z948)
return
}, 1)

tmp8152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4tuple, tmp8151)
}
__typedArg0 := symshen_4tuple
__typedArg1 := tmp8151
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8153 := MakeNative(func(__e *ControlFlow) {
Z949 := __e.Get(1)
_ = Z949
__e.TailApply(PrimFunc(symshen_4pvar), Z949)
return
}, 1)

tmp8154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4pvar, tmp8153)
}
__typedArg0 := symshen_4pvar
__typedArg1 := tmp8153
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8155 := MakeNative(func(__e *ControlFlow) {
Z950 := __e.Get(1)
_ = Z950
__e.TailApply(PrimFunc(symshen_4dictionary), Z950)
return
}, 1)

tmp8156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4dictionary, tmp8155)
}
__typedArg0 := symshen_4dictionary
__typedArg1 := tmp8155
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8157 := MakeNative(func(__e *ControlFlow) {
Z951 := __e.Get(1)
_ = Z951
__e.TailApply(PrimFunc(symshen_4print_1prolog_1vector), Z951)
return
}, 1)

tmp8158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1prolog_1vector, tmp8157)
}
__typedArg0 := symshen_4print_1prolog_1vector
__typedArg1 := tmp8157
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8159 := MakeNative(func(__e *ControlFlow) {
Z952 := __e.Get(1)
_ = Z952
__e.TailApply(PrimFunc(symshen_4print_1freshterm), Z952)
return
}, 1)

tmp8160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4print_1freshterm, tmp8159)
}
__typedArg0 := symshen_4print_1freshterm
__typedArg1 := tmp8159
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8161 := MakeNative(func(__e *ControlFlow) {
Z953 := __e.Get(1)
_ = Z953
__e.TailApply(PrimFunc(symshen_4printF), Z953)
return
}, 1)

tmp8162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4printF, tmp8161)
}
__typedArg0 := symshen_4printF
__typedArg1 := tmp8161
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8162, let__8061)
}
__typedArg0 := tmp8162
__typedArg1 := let__8061
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8160, tmp8163)
}
__typedArg0 := tmp8160
__typedArg1 := tmp8163
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8158, tmp8164)
}
__typedArg0 := tmp8158
__typedArg1 := tmp8164
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8156, tmp8165)
}
__typedArg0 := tmp8156
__typedArg1 := tmp8165
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8154, tmp8166)
}
__typedArg0 := tmp8154
__typedArg1 := tmp8166
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8152, tmp8167)
}
__typedArg0 := tmp8152
__typedArg1 := tmp8167
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8150, tmp8168)
}
__typedArg0 := tmp8150
__typedArg1 := tmp8168
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp8148, tmp8169)
}
__typedArg0 := tmp8148
__typedArg1 := tmp8169
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4for_1each), tmp8146, tmp8170)
return


}, 1)

__e.TailApply(ns2_1set, symshen_4build_1lambda_1table, tmp8143)
return




}, 0)

