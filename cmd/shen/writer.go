package main

import . "github.com/pyrex41/shen-go/kl"

var WriterMain = MakeNative(func(__e *ControlFlow) {
tmp2063 := MakeNative(func(__e *ControlFlow) {
V6928 := __e.Get(1)
_ = V6928
tmp2064 := Call(__e, PrimFunc(symshen_4insert), V6928, MakeString("~S"))


let__2058 := tmp2064
_ = let__2058

tmp2065 := Call(__e, PrimFunc(symstoutput))


tmp2066 := Call(__e, PrimFunc(sympr), let__2058, tmp2065)


let__2059 := tmp2066
_ = let__2059

__e.Return(V6928)
return


}, 1)

tmp2067 := Call(__e, ns2_1set, symprint, tmp2063)


_ = tmp2067

tmp2068 := MakeNative(func(__e *ControlFlow) {
V6931 := __e.Get(1)
_ = V6931
V6932 := __e.Get(2)
_ = V6932
tmp2073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dhush_d)
}
__typedArg0 := sym_dhush_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

if True == tmp2073 {
__e.Return(V6931)
return
} else {
tmp2071 := Call(__e, PrimFunc(symshen_4char_1stoutput_2), V6932)


if True == tmp2071 {
__e.TailApply(PrimFunc(symshen_4write_1string), V6931, V6932)
return
} else {
tmp2069 := Call(__e, PrimFunc(symshen_4string_1_6byte), V6931, MakeInteger(0))


__e.TailApply(PrimFunc(symshen_4write_1chars), V6931, V6932, tmp2069, MakeInteger(1))
return


}


}


}, 2)

tmp2074 := Call(__e, ns2_1set, sympr, tmp2068)


_ = tmp2074

tmp2075 := MakeNative(func(__e *ControlFlow) {
V6933 := __e.Get(1)
_ = V6933
V6934 := __e.Get(2)
_ = V6934
tmp2076 := MakeNative(func(__e *ControlFlow) {
tmp2077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sympos) {
return PrimPos(V6933, V6934)
}
__typedArg0 := V6933
__typedArg1 := V6934
return Call(__e, PrimFunc(sympos), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp2077)
}
__typedArg0 := tmp2077
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})())
return


}, 0)

tmp2078 := MakeNative(func(__e *ControlFlow) {
Z6935 := __e.Get(1)
_ = Z6935
__e.Return(symshen_4eos)
return
}, 1)

__e.TailApply(try_1catch, tmp2076, tmp2078)
return


}, 2)

tmp2079 := Call(__e, ns2_1set, symshen_4string_1_6byte, tmp2075)


_ = tmp2079

tmp2080 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__selftop:
V6936 := __self1
_ = V6936
V6937 := __self2
_ = V6937
V6938 := __self3
_ = V6938
V6939 := __self4
_ = V6939
tmp2085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4eos, V6938)
}
__typedArg0 := symshen_4eos
__typedArg1 := V6938
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2085 {
__e.Return(V6936)
return
} else {
tmp2081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symwrite_1byte) {
return PrimWriteByte(V6938, V6937)
}
__typedArg0 := V6938
__typedArg1 := V6937
return Call(__e, PrimFunc(symwrite_1byte), __typedArg0, __typedArg1)
})()

_ = tmp2081

tmp2082 := Call(__e, PrimFunc(symshen_4string_1_6byte), V6936, V6939)


if PrimFunc(symshen_4write_1chars) == __self {
__self1, __self2, __self3, __self4 = V6936, V6937, tmp2082, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V6939)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V6939
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4write_1chars), V6936, V6937, tmp2082, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V6939)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V6939
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


}


}, 4)

tmp2086 := Call(__e, ns2_1set, symshen_4write_1chars, tmp2080)


_ = tmp2086

tmp2087 := MakeNative(func(__e *ControlFlow) {
V6940 := __e.Get(1)
_ = V6940
V6941 := __e.Get(2)
_ = V6941
tmp2092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V6940)
}
__typedArg0 := V6940
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

if True == tmp2092 {
tmp2088 := Call(__e, PrimFunc(symshen_4proc_1nl), V6940)


__e.TailApply(PrimFunc(symshen_4mkstr_1l), tmp2088, V6941)
return


} else {
tmp2089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V6940, Nil)
}
__typedArg0 := V6940
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4proc_1nl, tmp2089)
}
__typedArg0 := symshen_4proc_1nl
__typedArg1 := tmp2089
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4mkstr_1r), tmp2090, V6941)
return


}


}, 2)

tmp2093 := Call(__e, ns2_1set, symshen_4mkstr, tmp2087)


_ = tmp2093

tmp2094 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V6946 := __self1
_ = V6946
V6947 := __self2
_ = V6947
tmp2101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V6947)
}
__typedArg0 := Nil
__typedArg1 := V6947
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2101 {
__e.Return(V6946)
return
} else {
tmp2099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V6947)
}
__typedArg0 := V6947
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2099 {
tmp2095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V6947)
}
__typedArg0 := V6947
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2096 := Call(__e, PrimFunc(symshen_4insert_1l), tmp2095, V6946)


tmp2097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6947)
}
__typedArg0 := V6947
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4mkstr_1l) == __self {
__self1, __self2 = tmp2096, tmp2097
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4mkstr_1l), tmp2096, tmp2097)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.mkstr-l"))
}
__typedArg0 := MakeString("implementation error in shen.mkstr-l")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp2102 := Call(__e, ns2_1set, symshen_4mkstr_1l, tmp2094)


_ = tmp2102

tmp2103 := MakeNative(func(__e *ControlFlow) {
V6954 := __e.Get(1)
_ = V6954
V6955 := __e.Get(2)
_ = V6955
tmp2241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V6955)
}
__typedArg0 := MakeString("")
__typedArg1 := V6955
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2241 {
__e.Return(MakeString(""))
return
} else {
tmp2239 := Call(__e, PrimFunc(symshen_4_7string_2), V6955)


var ifres2226 Obj

if True == tmp2239 {
tmp2237 := Call(__e, PrimFunc(symhdstr), V6955)


tmp2238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2237)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2237
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2228 Obj

if True == tmp2238 {
tmp2236 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6955)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V6955
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2230 Obj

if True == tmp2236 {
tmp2233 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6955)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V6955
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("A"), tmp2233)
}
__typedArg0 := MakeString("A")
__typedArg1 := tmp2233
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2231 Obj

if True == tmp2234 {
ifres2231 = True


} else {
ifres2231 = False


}

ifres2230 = ifres2231


} else {
ifres2230 = False


}

var ifres2229 Obj

if True == ifres2230 {
ifres2229 = True


} else {
ifres2229 = False


}

ifres2228 = ifres2229


} else {
ifres2228 = False


}

var ifres2227 Obj

if True == ifres2228 {
ifres2227 = True


} else {
ifres2227 = False


}

ifres2226 = ifres2227


} else {
ifres2226 = False


}

if True == ifres2226 {
tmp2105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6955)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6955)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V6955
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

tmp2106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4a, Nil)
}
__typedArg0 := symshen_4a
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2105, tmp2106)
}
__typedArg0 := tmp2105
__typedArg1 := tmp2106
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V6954, tmp2107)
}
__typedArg0 := V6954
__typedArg1 := tmp2107
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4app, tmp2108)
}
__typedArg0 := symshen_4app
__typedArg1 := tmp2108
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp2224 := Call(__e, PrimFunc(symshen_4_7string_2), V6955)


var ifres2211 Obj

if True == tmp2224 {
tmp2222 := Call(__e, PrimFunc(symhdstr), V6955)


tmp2223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2222)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2222
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2213 Obj

if True == tmp2223 {
tmp2221 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6955)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V6955
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2215 Obj

if True == tmp2221 {
tmp2218 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6955)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V6955
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("R"), tmp2218)
}
__typedArg0 := MakeString("R")
__typedArg1 := tmp2218
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2216 Obj

if True == tmp2219 {
ifres2216 = True


} else {
ifres2216 = False


}

ifres2215 = ifres2216


} else {
ifres2215 = False


}

var ifres2214 Obj

if True == ifres2215 {
ifres2214 = True


} else {
ifres2214 = False


}

ifres2213 = ifres2214


} else {
ifres2213 = False


}

var ifres2212 Obj

if True == ifres2213 {
ifres2212 = True


} else {
ifres2212 = False


}

ifres2211 = ifres2212


} else {
ifres2211 = False


}

if True == ifres2211 {
tmp2110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6955)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6955)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V6955
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

tmp2111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4r, Nil)
}
__typedArg0 := symshen_4r
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2110, tmp2111)
}
__typedArg0 := tmp2110
__typedArg1 := tmp2111
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V6954, tmp2112)
}
__typedArg0 := V6954
__typedArg1 := tmp2112
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4app, tmp2113)
}
__typedArg0 := symshen_4app
__typedArg1 := tmp2113
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp2209 := Call(__e, PrimFunc(symshen_4_7string_2), V6955)


var ifres2196 Obj

if True == tmp2209 {
tmp2207 := Call(__e, PrimFunc(symhdstr), V6955)


tmp2208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2207)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2207
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2198 Obj

if True == tmp2208 {
tmp2206 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6955)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V6955
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2200 Obj

if True == tmp2206 {
tmp2203 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6955)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V6955
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("S"), tmp2203)
}
__typedArg0 := MakeString("S")
__typedArg1 := tmp2203
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2201 Obj

if True == tmp2204 {
ifres2201 = True


} else {
ifres2201 = False


}

ifres2200 = ifres2201


} else {
ifres2200 = False


}

var ifres2199 Obj

if True == ifres2200 {
ifres2199 = True


} else {
ifres2199 = False


}

ifres2198 = ifres2199


} else {
ifres2198 = False


}

var ifres2197 Obj

if True == ifres2198 {
ifres2197 = True


} else {
ifres2197 = False


}

ifres2196 = ifres2197


} else {
ifres2196 = False


}

if True == ifres2196 {
tmp2115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6955)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6955)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V6955
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

tmp2116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4s, Nil)
}
__typedArg0 := symshen_4s
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2115, tmp2116)
}
__typedArg0 := tmp2115
__typedArg1 := tmp2116
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V6954, tmp2117)
}
__typedArg0 := V6954
__typedArg1 := tmp2117
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4app, tmp2118)
}
__typedArg0 := symshen_4app
__typedArg1 := tmp2118
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp2194 := Call(__e, PrimFunc(symshen_4_7string_2), V6955)


if True == tmp2194 {
tmp2119 := Call(__e, PrimFunc(symhdstr), V6955)


tmp2121 := Call(__e, PrimFunc(symshen_4insert_1l), V6954, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6955)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V6955
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2121, Nil)
}
__typedArg0 := tmp2121
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2119, tmp2122)
}
__typedArg0 := tmp2119
__typedArg1 := tmp2122
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp2123)
}
__typedArg0 := symcn
__typedArg1 := tmp2123
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4factor_1cn), tmp2124)
return


} else {
tmp2192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V6955)
}
__typedArg0 := V6955
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2173 Obj

if True == tmp2192 {
tmp2190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V6955)
}
__typedArg0 := V6955
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcn, tmp2190)
}
__typedArg0 := symcn
__typedArg1 := tmp2190
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2175 Obj

if True == tmp2191 {
tmp2188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6955)
}
__typedArg0 := V6955
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2188)
}
__typedArg0 := tmp2188
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2177 Obj

if True == tmp2189 {
tmp2185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6955)
}
__typedArg0 := V6955
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2185)
}
__typedArg0 := tmp2185
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2186)
}
__typedArg0 := tmp2186
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2179 Obj

if True == tmp2187 {
tmp2181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6955)
}
__typedArg0 := V6955
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2181)
}
__typedArg0 := tmp2181
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2182)
}
__typedArg0 := tmp2182
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2183)
}
__typedArg0 := Nil
__typedArg1 := tmp2183
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2180 Obj

if True == tmp2184 {
ifres2180 = True


} else {
ifres2180 = False


}

ifres2179 = ifres2180


} else {
ifres2179 = False


}

var ifres2178 Obj

if True == ifres2179 {
ifres2178 = True


} else {
ifres2178 = False


}

ifres2177 = ifres2178


} else {
ifres2177 = False


}

var ifres2176 Obj

if True == ifres2177 {
ifres2176 = True


} else {
ifres2176 = False


}

ifres2175 = ifres2176


} else {
ifres2175 = False


}

var ifres2174 Obj

if True == ifres2175 {
ifres2174 = True


} else {
ifres2174 = False


}

ifres2173 = ifres2174


} else {
ifres2173 = False


}

if True == ifres2173 {
tmp2125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6955)
}
__typedArg0 := V6955
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2125)
}
__typedArg0 := tmp2125
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6955)
}
__typedArg0 := V6955
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2127)
}
__typedArg0 := tmp2127
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2128)
}
__typedArg0 := tmp2128
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2130 := Call(__e, PrimFunc(symshen_4insert_1l), V6954, tmp2129)


tmp2131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2130, Nil)
}
__typedArg0 := tmp2130
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2126, tmp2131)
}
__typedArg0 := tmp2126
__typedArg1 := tmp2131
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp2132)
}
__typedArg0 := symcn
__typedArg1 := tmp2132
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp2171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V6955)
}
__typedArg0 := V6955
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2145 Obj

if True == tmp2171 {
tmp2169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V6955)
}
__typedArg0 := V6955
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4app, tmp2169)
}
__typedArg0 := symshen_4app
__typedArg1 := tmp2169
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2147 Obj

if True == tmp2170 {
tmp2167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6955)
}
__typedArg0 := V6955
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2167)
}
__typedArg0 := tmp2167
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2149 Obj

if True == tmp2168 {
tmp2164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6955)
}
__typedArg0 := V6955
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2164)
}
__typedArg0 := tmp2164
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2165)
}
__typedArg0 := tmp2165
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2151 Obj

if True == tmp2166 {
tmp2160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6955)
}
__typedArg0 := V6955
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2160)
}
__typedArg0 := tmp2160
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2161)
}
__typedArg0 := tmp2161
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2162)
}
__typedArg0 := tmp2162
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2153 Obj

if True == tmp2163 {
tmp2155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6955)
}
__typedArg0 := V6955
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2155)
}
__typedArg0 := tmp2155
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2156)
}
__typedArg0 := tmp2156
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2157)
}
__typedArg0 := tmp2157
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2158)
}
__typedArg0 := Nil
__typedArg1 := tmp2158
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2154 Obj

if True == tmp2159 {
ifres2154 = True


} else {
ifres2154 = False


}

ifres2153 = ifres2154


} else {
ifres2153 = False


}

var ifres2152 Obj

if True == ifres2153 {
ifres2152 = True


} else {
ifres2152 = False


}

ifres2151 = ifres2152


} else {
ifres2151 = False


}

var ifres2150 Obj

if True == ifres2151 {
ifres2150 = True


} else {
ifres2150 = False


}

ifres2149 = ifres2150


} else {
ifres2149 = False


}

var ifres2148 Obj

if True == ifres2149 {
ifres2148 = True


} else {
ifres2148 = False


}

ifres2147 = ifres2148


} else {
ifres2147 = False


}

var ifres2146 Obj

if True == ifres2147 {
ifres2146 = True


} else {
ifres2146 = False


}

ifres2145 = ifres2146


} else {
ifres2145 = False


}

if True == ifres2145 {
tmp2133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6955)
}
__typedArg0 := V6955
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2133)
}
__typedArg0 := tmp2133
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6955)
}
__typedArg0 := V6955
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2135)
}
__typedArg0 := tmp2135
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2136)
}
__typedArg0 := tmp2136
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2138 := Call(__e, PrimFunc(symshen_4insert_1l), V6954, tmp2137)


tmp2139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6955)
}
__typedArg0 := V6955
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2139)
}
__typedArg0 := tmp2139
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2140)
}
__typedArg0 := tmp2140
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2138, tmp2141)
}
__typedArg0 := tmp2138
__typedArg1 := tmp2141
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2134, tmp2142)
}
__typedArg0 := tmp2134
__typedArg1 := tmp2142
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4app, tmp2143)
}
__typedArg0 := symshen_4app
__typedArg1 := tmp2143
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.insert-l"))
}
__typedArg0 := MakeString("implementation error in shen.insert-l")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}


}


}


}


}, 2)

tmp2242 := Call(__e, ns2_1set, symshen_4insert_1l, tmp2103)


_ = tmp2242

tmp2243 := MakeNative(func(__e *ControlFlow) {
V6956 := __e.Get(1)
_ = V6956
tmp2328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V6956)
}
__typedArg0 := V6956
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2259 Obj

if True == tmp2328 {
tmp2326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V6956)
}
__typedArg0 := V6956
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcn, tmp2326)
}
__typedArg0 := symcn
__typedArg1 := tmp2326
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2261 Obj

if True == tmp2327 {
tmp2324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6956)
}
__typedArg0 := V6956
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2324)
}
__typedArg0 := tmp2324
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2263 Obj

if True == tmp2325 {
tmp2321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6956)
}
__typedArg0 := V6956
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2321)
}
__typedArg0 := tmp2321
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2322)
}
__typedArg0 := tmp2322
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2265 Obj

if True == tmp2323 {
tmp2317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6956)
}
__typedArg0 := V6956
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2317)
}
__typedArg0 := tmp2317
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2318)
}
__typedArg0 := tmp2318
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2319)
}
__typedArg0 := tmp2319
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2267 Obj

if True == tmp2320 {
tmp2312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6956)
}
__typedArg0 := V6956
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2312)
}
__typedArg0 := tmp2312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2313)
}
__typedArg0 := tmp2313
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2314)
}
__typedArg0 := tmp2314
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcn, tmp2315)
}
__typedArg0 := symcn
__typedArg1 := tmp2315
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2269 Obj

if True == tmp2316 {
tmp2307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6956)
}
__typedArg0 := V6956
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2307)
}
__typedArg0 := tmp2307
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2308)
}
__typedArg0 := tmp2308
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2309)
}
__typedArg0 := tmp2309
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2310)
}
__typedArg0 := tmp2310
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2271 Obj

if True == tmp2311 {
tmp2301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6956)
}
__typedArg0 := V6956
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2301)
}
__typedArg0 := tmp2301
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2302)
}
__typedArg0 := tmp2302
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2303)
}
__typedArg0 := tmp2303
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2304)
}
__typedArg0 := tmp2304
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2305)
}
__typedArg0 := tmp2305
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2273 Obj

if True == tmp2306 {
tmp2294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6956)
}
__typedArg0 := V6956
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2294)
}
__typedArg0 := tmp2294
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2295)
}
__typedArg0 := tmp2295
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2296)
}
__typedArg0 := tmp2296
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2297)
}
__typedArg0 := tmp2297
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2298)
}
__typedArg0 := tmp2298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2299)
}
__typedArg0 := Nil
__typedArg1 := tmp2299
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2275 Obj

if True == tmp2300 {
tmp2290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6956)
}
__typedArg0 := V6956
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2290)
}
__typedArg0 := tmp2290
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2291)
}
__typedArg0 := tmp2291
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2292)
}
__typedArg0 := Nil
__typedArg1 := tmp2292
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2277 Obj

if True == tmp2293 {
tmp2287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6956)
}
__typedArg0 := V6956
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2287)
}
__typedArg0 := tmp2287
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(tmp2288)
}
__typedArg0 := tmp2288
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres2279 Obj

if True == tmp2289 {
tmp2281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6956)
}
__typedArg0 := V6956
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2281)
}
__typedArg0 := tmp2281
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2282)
}
__typedArg0 := tmp2282
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2283)
}
__typedArg0 := tmp2283
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2284)
}
__typedArg0 := tmp2284
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(tmp2285)
}
__typedArg0 := tmp2285
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres2280 Obj

if True == tmp2286 {
ifres2280 = True


} else {
ifres2280 = False


}

ifres2279 = ifres2280


} else {
ifres2279 = False


}

var ifres2278 Obj

if True == ifres2279 {
ifres2278 = True


} else {
ifres2278 = False


}

ifres2277 = ifres2278


} else {
ifres2277 = False


}

var ifres2276 Obj

if True == ifres2277 {
ifres2276 = True


} else {
ifres2276 = False


}

ifres2275 = ifres2276


} else {
ifres2275 = False


}

var ifres2274 Obj

if True == ifres2275 {
ifres2274 = True


} else {
ifres2274 = False


}

ifres2273 = ifres2274


} else {
ifres2273 = False


}

var ifres2272 Obj

if True == ifres2273 {
ifres2272 = True


} else {
ifres2272 = False


}

ifres2271 = ifres2272


} else {
ifres2271 = False


}

var ifres2270 Obj

if True == ifres2271 {
ifres2270 = True


} else {
ifres2270 = False


}

ifres2269 = ifres2270


} else {
ifres2269 = False


}

var ifres2268 Obj

if True == ifres2269 {
ifres2268 = True


} else {
ifres2268 = False


}

ifres2267 = ifres2268


} else {
ifres2267 = False


}

var ifres2266 Obj

if True == ifres2267 {
ifres2266 = True


} else {
ifres2266 = False


}

ifres2265 = ifres2266


} else {
ifres2265 = False


}

var ifres2264 Obj

if True == ifres2265 {
ifres2264 = True


} else {
ifres2264 = False


}

ifres2263 = ifres2264


} else {
ifres2263 = False


}

var ifres2262 Obj

if True == ifres2263 {
ifres2262 = True


} else {
ifres2262 = False


}

ifres2261 = ifres2262


} else {
ifres2261 = False


}

var ifres2260 Obj

if True == ifres2261 {
ifres2260 = True


} else {
ifres2260 = False


}

ifres2259 = ifres2260


} else {
ifres2259 = False


}

if True == ifres2259 {
tmp2244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6956)
}
__typedArg0 := V6956
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2244)
}
__typedArg0 := tmp2244
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6956)
}
__typedArg0 := V6956
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2246)
}
__typedArg0 := tmp2246
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2247)
}
__typedArg0 := tmp2247
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2248)
}
__typedArg0 := tmp2248
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2249)
}
__typedArg0 := tmp2249
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp2245)
__typedS1, __typedOK1 := TypedString(tmp2250)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp2245
__typedArg1 := tmp2250
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp2252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6956)
}
__typedArg0 := V6956
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2252)
}
__typedArg0 := tmp2252
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2253)
}
__typedArg0 := tmp2253
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2254)
}
__typedArg0 := tmp2254
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2255)
}
__typedArg0 := tmp2255
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2251, tmp2256)
}
__typedArg0 := tmp2251
__typedArg1 := tmp2256
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp2257)
}
__typedArg0 := symcn
__typedArg1 := tmp2257
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V6956)
return
}


}, 1)

tmp2329 := Call(__e, ns2_1set, symshen_4factor_1cn, tmp2243)


_ = tmp2329

tmp2330 := MakeNative(func(__e *ControlFlow) {
V6959 := __e.Get(1)
_ = V6959
tmp2356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V6959)
}
__typedArg0 := MakeString("")
__typedArg1 := V6959
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2356 {
__e.Return(MakeString(""))
return
} else {
tmp2354 := Call(__e, PrimFunc(symshen_4_7string_2), V6959)


var ifres2341 Obj

if True == tmp2354 {
tmp2352 := Call(__e, PrimFunc(symhdstr), V6959)


tmp2353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2352)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2352
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2343 Obj

if True == tmp2353 {
tmp2351 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6959)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V6959
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2345 Obj

if True == tmp2351 {
tmp2348 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6959)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V6959
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("%"), tmp2348)
}
__typedArg0 := MakeString("%")
__typedArg1 := tmp2348
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2346 Obj

if True == tmp2349 {
ifres2346 = True


} else {
ifres2346 = False


}

ifres2345 = ifres2346


} else {
ifres2345 = False


}

var ifres2344 Obj

if True == ifres2345 {
ifres2344 = True


} else {
ifres2344 = False


}

ifres2343 = ifres2344


} else {
ifres2343 = False


}

var ifres2342 Obj

if True == ifres2343 {
ifres2342 = True


} else {
ifres2342 = False


}

ifres2341 = ifres2342


} else {
ifres2341 = False


}

if True == ifres2341 {
tmp2331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(MakeInteger(10))
}
__typedArg0 := MakeInteger(10)
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp2334 := Call(__e, PrimFunc(symshen_4proc_1nl), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6959)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6959)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V6959
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp2331)
__typedS1, __typedOK1 := TypedString(tmp2334)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp2331
__typedArg1 := tmp2334
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
tmp2339 := Call(__e, PrimFunc(symshen_4_7string_2), V6959)


if True == tmp2339 {
tmp2335 := Call(__e, PrimFunc(symhdstr), V6959)


tmp2337 := Call(__e, PrimFunc(symshen_4proc_1nl), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6959)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V6959
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp2335)
__typedS1, __typedOK1 := TypedString(tmp2337)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp2335
__typedArg1 := tmp2337
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.proc-nl"))
}
__typedArg0 := MakeString("implementation error in shen.proc-nl")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 1)

tmp2357 := Call(__e, ns2_1set, symshen_4proc_1nl, tmp2330)


_ = tmp2357

tmp2358 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V6964 := __self1
_ = V6964
V6965 := __self2
_ = V6965
tmp2367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V6965)
}
__typedArg0 := Nil
__typedArg1 := V6965
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2367 {
__e.Return(V6964)
return
} else {
tmp2365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V6965)
}
__typedArg0 := V6965
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2365 {
tmp2359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V6965)
}
__typedArg0 := V6965
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V6964, Nil)
}
__typedArg0 := V6964
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2359, tmp2360)
}
__typedArg0 := tmp2359
__typedArg1 := tmp2360
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4insert, tmp2361)
}
__typedArg0 := symshen_4insert
__typedArg1 := tmp2361
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6965)
}
__typedArg0 := V6965
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4mkstr_1r) == __self {
__self1, __self2 = tmp2362, tmp2363
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4mkstr_1r), tmp2362, tmp2363)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.mkstr-r"))
}
__typedArg0 := MakeString("implementation error in shen.mkstr-r")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp2368 := Call(__e, ns2_1set, symshen_4mkstr_1r, tmp2358)


_ = tmp2368

tmp2369 := MakeNative(func(__e *ControlFlow) {
V6966 := __e.Get(1)
_ = V6966
V6967 := __e.Get(2)
_ = V6967
__e.TailApply(PrimFunc(symshen_4insert_1h), V6966, V6967, MakeString(""))
return
}, 2)

tmp2370 := Call(__e, ns2_1set, symshen_4insert, tmp2369)


_ = tmp2370

tmp2371 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__selftop:
V6976 := __self1
_ = V6976
V6977 := __self2
_ = V6977
V6978 := __self3
_ = V6978
tmp2432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V6977)
}
__typedArg0 := MakeString("")
__typedArg1 := V6977
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2432 {
__e.Return(V6978)
return
} else {
tmp2430 := Call(__e, PrimFunc(symshen_4_7string_2), V6977)


var ifres2417 Obj

if True == tmp2430 {
tmp2428 := Call(__e, PrimFunc(symhdstr), V6977)


tmp2429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2428)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2428
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2419 Obj

if True == tmp2429 {
tmp2427 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6977)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V6977
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2421 Obj

if True == tmp2427 {
tmp2424 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6977)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V6977
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("A"), tmp2424)
}
__typedArg0 := MakeString("A")
__typedArg1 := tmp2424
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2422 Obj

if True == tmp2425 {
ifres2422 = True


} else {
ifres2422 = False


}

ifres2421 = ifres2422


} else {
ifres2421 = False


}

var ifres2420 Obj

if True == ifres2421 {
ifres2420 = True


} else {
ifres2420 = False


}

ifres2419 = ifres2420


} else {
ifres2419 = False


}

var ifres2418 Obj

if True == ifres2419 {
ifres2418 = True


} else {
ifres2418 = False


}

ifres2417 = ifres2418


} else {
ifres2417 = False


}

if True == ifres2417 {
tmp2374 := Call(__e, PrimFunc(symshen_4app), V6976, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6977)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6977)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V6977
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})(), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(V6978)
__typedS1, __typedOK1 := TypedString(tmp2374)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V6978
__typedArg1 := tmp2374
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
tmp2415 := Call(__e, PrimFunc(symshen_4_7string_2), V6977)


var ifres2402 Obj

if True == tmp2415 {
tmp2413 := Call(__e, PrimFunc(symhdstr), V6977)


tmp2414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2413)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2413
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2404 Obj

if True == tmp2414 {
tmp2412 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6977)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V6977
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2406 Obj

if True == tmp2412 {
tmp2409 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6977)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V6977
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("R"), tmp2409)
}
__typedArg0 := MakeString("R")
__typedArg1 := tmp2409
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2407 Obj

if True == tmp2410 {
ifres2407 = True


} else {
ifres2407 = False


}

ifres2406 = ifres2407


} else {
ifres2406 = False


}

var ifres2405 Obj

if True == ifres2406 {
ifres2405 = True


} else {
ifres2405 = False


}

ifres2404 = ifres2405


} else {
ifres2404 = False


}

var ifres2403 Obj

if True == ifres2404 {
ifres2403 = True


} else {
ifres2403 = False


}

ifres2402 = ifres2403


} else {
ifres2402 = False


}

if True == ifres2402 {
tmp2377 := Call(__e, PrimFunc(symshen_4app), V6976, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6977)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6977)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V6977
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})(), symshen_4r)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(V6978)
__typedS1, __typedOK1 := TypedString(tmp2377)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V6978
__typedArg1 := tmp2377
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
tmp2400 := Call(__e, PrimFunc(symshen_4_7string_2), V6977)


var ifres2387 Obj

if True == tmp2400 {
tmp2398 := Call(__e, PrimFunc(symhdstr), V6977)


tmp2399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2398)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2398
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2389 Obj

if True == tmp2399 {
tmp2397 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6977)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V6977
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2391 Obj

if True == tmp2397 {
tmp2394 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6977)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V6977
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("S"), tmp2394)
}
__typedArg0 := MakeString("S")
__typedArg1 := tmp2394
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2392 Obj

if True == tmp2395 {
ifres2392 = True


} else {
ifres2392 = False


}

ifres2391 = ifres2392


} else {
ifres2391 = False


}

var ifres2390 Obj

if True == ifres2391 {
ifres2390 = True


} else {
ifres2390 = False


}

ifres2389 = ifres2390


} else {
ifres2389 = False


}

var ifres2388 Obj

if True == ifres2389 {
ifres2388 = True


} else {
ifres2388 = False


}

ifres2387 = ifres2388


} else {
ifres2387 = False


}

if True == ifres2387 {
tmp2380 := Call(__e, PrimFunc(symshen_4app), V6976, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6977)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6977)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V6977
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})(), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(V6978)
__typedS1, __typedOK1 := TypedString(tmp2380)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V6978
__typedArg1 := tmp2380
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
tmp2385 := Call(__e, PrimFunc(symshen_4_7string_2), V6977)


if True == tmp2385 {
tmp2381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V6977)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V6977
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

tmp2382 := Call(__e, PrimFunc(symhdstr), V6977)


if PrimFunc(symshen_4insert_1h) == __self {
__self1, __self2, __self3 = V6976, tmp2381, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(V6978)
__typedS1, __typedOK1 := TypedString(tmp2382)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V6978
__typedArg1 := tmp2382
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4insert_1h), V6976, tmp2381, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(V6978)
__typedS1, __typedOK1 := TypedString(tmp2382)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V6978
__typedArg1 := tmp2382
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.insert-h"))
}
__typedArg0 := MakeString("implementation error in shen.insert-h")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}


}


}, 3)

tmp2433 := Call(__e, ns2_1set, symshen_4insert_1h, tmp2371)


_ = tmp2433

tmp2434 := MakeNative(func(__e *ControlFlow) {
V6979 := __e.Get(1)
_ = V6979
V6980 := __e.Get(2)
_ = V6980
V6981 := __e.Get(3)
_ = V6981
tmp2435 := Call(__e, PrimFunc(symshen_4arg_1_6str), V6979, V6981)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp2435)
__typedS1, __typedOK1 := TypedString(V6980)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp2435
__typedArg1 := V6980
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 3)

tmp2436 := Call(__e, ns2_1set, symshen_4app, tmp2434)


_ = tmp2436

tmp2437 := MakeNative(func(__e *ControlFlow) {
V6985 := __e.Get(1)
_ = V6985
V6986 := __e.Get(2)
_ = V6986
tmp2445 := Call(__e, PrimFunc(symfail))


tmp2446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V6985, tmp2445)
}
__typedArg0 := V6985
__typedArg1 := tmp2445
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2446 {
__e.Return(MakeString("..."))
return
} else {
tmp2443 := Call(__e, PrimFunc(symshen_4list_2), V6985)


if True == tmp2443 {
__e.TailApply(PrimFunc(symshen_4list_1_6str), V6985, V6986)
return
} else {
tmp2441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V6985)
}
__typedArg0 := V6985
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

if True == tmp2441 {
__e.TailApply(PrimFunc(symshen_4str_1_6str), V6985, V6986)
return
} else {
tmp2439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V6985)
}
__typedArg0 := V6985
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp2439 {
__e.TailApply(PrimFunc(symshen_4vector_1_6str), V6985, V6986)
return
} else {
__e.TailApply(PrimFunc(symshen_4atom_1_6str), V6985)
return
}


}


}


}


}, 2)

tmp2447 := Call(__e, ns2_1set, symshen_4arg_1_6str, tmp2437)


_ = tmp2447

tmp2448 := MakeNative(func(__e *ControlFlow) {
V6987 := __e.Get(1)
_ = V6987
V6988 := __e.Get(2)
_ = V6988
tmp2456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4r, V6988)
}
__typedArg0 := symshen_4r
__typedArg1 := V6988
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2456 {
tmp2449 := Call(__e, PrimFunc(symshen_4maxseq))


tmp2450 := Call(__e, PrimFunc(symshen_4iter_1list), V6987, symshen_4r, tmp2449)


tmp2451 := Call(__e, PrimFunc(sym_8s), tmp2450, MakeString(")"))


__e.TailApply(PrimFunc(sym_8s), MakeString("("), tmp2451)
return


} else {
tmp2452 := Call(__e, PrimFunc(symshen_4maxseq))


tmp2453 := Call(__e, PrimFunc(symshen_4iter_1list), V6987, V6988, tmp2452)


tmp2454 := Call(__e, PrimFunc(sym_8s), tmp2453, MakeString("]"))


__e.TailApply(PrimFunc(sym_8s), MakeString("["), tmp2454)
return


}


}, 2)

tmp2457 := Call(__e, ns2_1set, symshen_4list_1_6str, tmp2448)


_ = tmp2457

tmp2458 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmaximum_1print_1sequence_1size_d)
}
__typedArg0 := sym_dmaximum_1print_1sequence_1size_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp2459 := Call(__e, ns2_1set, symshen_4maxseq, tmp2458)


_ = tmp2459

tmp2460 := MakeNative(func(__e *ControlFlow) {
V6999 := __e.Get(1)
_ = V6999
V7000 := __e.Get(2)
_ = V7000
V7001 := __e.Get(3)
_ = V7001
tmp2481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V6999)
}
__typedArg0 := Nil
__typedArg1 := V6999
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2481 {
__e.Return(MakeString(""))
return
} else {
tmp2479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V7001)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V7001
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2479 {
__e.Return(MakeString("... etc"))
return
} else {
tmp2477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V6999)
}
__typedArg0 := V6999
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2473 Obj

if True == tmp2477 {
tmp2475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6999)
}
__typedArg0 := V6999
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2475)
}
__typedArg0 := Nil
__typedArg1 := tmp2475
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2474 Obj

if True == tmp2476 {
ifres2474 = True


} else {
ifres2474 = False


}

ifres2473 = ifres2474


} else {
ifres2473 = False


}

if True == ifres2473 {
tmp2461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V6999)
}
__typedArg0 := V6999
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4arg_1_6str), tmp2461, V7000)
return


} else {
tmp2471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V6999)
}
__typedArg0 := V6999
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2471 {
tmp2462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V6999)
}
__typedArg0 := V6999
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2463 := Call(__e, PrimFunc(symshen_4arg_1_6str), tmp2462, V7000)


tmp2464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V6999)
}
__typedArg0 := V6999
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2466 := Call(__e, PrimFunc(symshen_4iter_1list), tmp2464, V7000, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V7001)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V7001
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


tmp2467 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp2466)


__e.TailApply(PrimFunc(sym_8s), tmp2463, tmp2467)
return


} else {
tmp2468 := Call(__e, PrimFunc(symshen_4arg_1_6str), V6999, V7000)


tmp2469 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp2468)


__e.TailApply(PrimFunc(sym_8s), MakeString("|"), tmp2469)
return


}


}


}


}


}, 3)

tmp2482 := Call(__e, ns2_1set, symshen_4iter_1list, tmp2460)


_ = tmp2482

tmp2483 := MakeNative(func(__e *ControlFlow) {
V7004 := __e.Get(1)
_ = V7004
V7005 := __e.Get(2)
_ = V7005
tmp2488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4a, V7005)
}
__typedArg0 := symshen_4a
__typedArg1 := V7005
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2488 {
__e.Return(V7004)
return
} else {
tmp2484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(MakeInteger(34))
}
__typedArg0 := MakeInteger(34)
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp2485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(MakeInteger(34))
}
__typedArg0 := MakeInteger(34)
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp2486 := Call(__e, PrimFunc(sym_8s), V7004, tmp2485)


__e.TailApply(PrimFunc(sym_8s), tmp2484, tmp2486)
return


}


}, 2)

tmp2489 := Call(__e, ns2_1set, symshen_4str_1_6str, tmp2483)


_ = tmp2489

tmp2490 := MakeNative(func(__e *ControlFlow) {
V7006 := __e.Get(1)
_ = V7006
V7007 := __e.Get(2)
_ = V7007
tmp2503 := Call(__e, PrimFunc(symshen_4print_1vector_2), V7006)


if True == tmp2503 {
tmp2491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V7006, MakeInteger(0))
}
__typedArg0 := V7006
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp2492 := Call(__e, PrimFunc(symfn), tmp2491)


__e.TailApply(tmp2492, V7006)
return


} else {
tmp2501 := Call(__e, PrimFunc(symvector_2), V7006)


if True == tmp2501 {
tmp2493 := Call(__e, PrimFunc(symshen_4maxseq))


tmp2494 := Call(__e, PrimFunc(symshen_4iter_1vector), V7006, MakeInteger(1), V7007, tmp2493)


tmp2495 := Call(__e, PrimFunc(sym_8s), tmp2494, MakeString(">"))


__e.TailApply(PrimFunc(sym_8s), MakeString("<"), tmp2495)
return


} else {
tmp2496 := Call(__e, PrimFunc(symshen_4maxseq))


tmp2497 := Call(__e, PrimFunc(symshen_4iter_1vector), V7006, MakeInteger(0), V7007, tmp2496)


tmp2498 := Call(__e, PrimFunc(sym_8s), tmp2497, MakeString(">>"))


tmp2499 := Call(__e, PrimFunc(sym_8s), MakeString("<"), tmp2498)


__e.TailApply(PrimFunc(sym_8s), MakeString("<"), tmp2499)
return


}


}


}, 2)

tmp2504 := Call(__e, ns2_1set, symshen_4vector_1_6str, tmp2490)


_ = tmp2504

tmp2505 := MakeNative(func(__e *ControlFlow) {
V7008 := __e.Get(1)
_ = V7008
tmp2506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dempty_1absvector_d)
}
__typedArg0 := symshen_4_dempty_1absvector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V7008, tmp2506)
}
__typedArg0 := V7008
__typedArg1 := tmp2506
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp2507 := Call(__e, ns2_1set, symshen_4empty_1absvector_2, tmp2505)


_ = tmp2507

tmp2508 := MakeNative(func(__e *ControlFlow) {
V7009 := __e.Get(1)
_ = V7009
tmp2526 := Call(__e, PrimFunc(symshen_4empty_1absvector_2), V7009)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp2526)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp2526
return Call(__e, PrimFunc(symnot), __typedArg0)
})() {
tmp2510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V7009, MakeInteger(0))
}
__typedArg0 := V7009
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

let__2060 := tmp2510
_ = let__2060

tmp2524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__2060, symshen_4tuple)
}
__typedArg0 := let__2060
__typedArg1 := symshen_4tuple
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2511 Obj

if True == tmp2524 {
ifres2511 = True


} else {
tmp2523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__2060, symshen_4pvar)
}
__typedArg0 := let__2060
__typedArg1 := symshen_4pvar
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2513 Obj

if True == tmp2523 {
ifres2513 = True


} else {
tmp2522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__2060, symshen_4dictionary)
}
__typedArg0 := let__2060
__typedArg1 := symshen_4dictionary
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2515 Obj

if True == tmp2522 {
ifres2515 = True


} else {
tmp2520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(let__2060)
}
__typedArg0 := let__2060
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

tmp2521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp2520)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp2520
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres2517 Obj

if True == tmp2521 {
tmp2519 := Call(__e, PrimFunc(symshen_4fbound_2), let__2060)


var ifres2518 Obj

if True == tmp2519 {
ifres2518 = True


} else {
ifres2518 = False


}

ifres2517 = ifres2518


} else {
ifres2517 = False


}

var ifres2516 Obj

if True == ifres2517 {
ifres2516 = True


} else {
ifres2516 = False


}

ifres2515 = ifres2516


}

var ifres2514 Obj

if True == ifres2515 {
ifres2514 = True


} else {
ifres2514 = False


}

ifres2513 = ifres2514


}

var ifres2512 Obj

if True == ifres2513 {
ifres2512 = True


} else {
ifres2512 = False


}

ifres2511 = ifres2512


}

if True == ifres2511 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


} else {
__e.Return(False)
return
}


}, 1)

tmp2528 := Call(__e, ns2_1set, symshen_4print_1vector_2, tmp2508)


_ = tmp2528

tmp2529 := MakeNative(func(__e *ControlFlow) {
V7011 := __e.Get(1)
_ = V7011
tmp2530 := Call(__e, PrimFunc(symarity), V7011)


tmp2531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp2530, MakeInteger(-1))
}
__typedArg0 := tmp2530
__typedArg1 := MakeInteger(-1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp2531)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp2531
return Call(__e, PrimFunc(symnot), __typedArg0)
})())
return


}, 1)

tmp2532 := Call(__e, ns2_1set, symshen_4fbound_2, tmp2529)


_ = tmp2532

tmp2533 := MakeNative(func(__e *ControlFlow) {
V7012 := __e.Get(1)
_ = V7012
tmp2534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V7012, MakeInteger(1))
}
__typedArg0 := V7012
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp2535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V7012, MakeInteger(2))
}
__typedArg0 := V7012
__typedArg1 := MakeInteger(2)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp2536 := Call(__e, PrimFunc(symshen_4app), tmp2535, MakeString(")"), symshen_4s)


tmp2538 := Call(__e, PrimFunc(symshen_4app), tmp2534, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" "))
__typedS1, __typedOK1 := TypedString(tmp2536)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" ")
__typedArg1 := tmp2536
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("(@p "))
__typedS1, __typedOK1 := TypedString(tmp2538)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("(@p ")
__typedArg1 := tmp2538
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp2539 := Call(__e, ns2_1set, symshen_4tuple, tmp2533)


_ = tmp2539

tmp2540 := MakeNative(func(__e *ControlFlow) {
V7013 := __e.Get(1)
_ = V7013
__e.Return(MakeString("(dict ...)"))
return
}, 1)

tmp2541 := Call(__e, ns2_1set, symshen_4dictionary, tmp2540)


_ = tmp2541

tmp2542 := MakeNative(func(__e *ControlFlow) {
V7020 := __e.Get(1)
_ = V7020
V7021 := __e.Get(2)
_ = V7021
V7022 := __e.Get(3)
_ = V7022
V7023 := __e.Get(4)
_ = V7023
tmp2560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V7023)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V7023
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2560 {
__e.Return(MakeString("... etc"))
return
} else {
tmp2543 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V7020, V7021)
}
__typedArg0 := V7020
__typedArg1 := V7021
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 0)

tmp2544 := MakeNative(func(__e *ControlFlow) {
Z7025 := __e.Get(1)
_ = Z7025
__e.Return(symshen_4out_1of_1bounds)
return
}, 1)

tmp2545 := Call(__e, try_1catch, tmp2543, tmp2544)


let__2061 := tmp2545
_ = let__2061

tmp2546 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V7020, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V7021)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V7021
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := V7020
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V7021)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V7021
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp2548 := MakeNative(func(__e *ControlFlow) {
Z7027 := __e.Get(1)
_ = Z7027
__e.Return(symshen_4out_1of_1bounds)
return
}, 1)

tmp2549 := Call(__e, try_1catch, tmp2546, tmp2548)


let__2062 := tmp2549
_ = let__2062

tmp2558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__2061, symshen_4out_1of_1bounds)
}
__typedArg0 := let__2061
__typedArg1 := symshen_4out_1of_1bounds
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2558 {
__e.Return(MakeString(""))
return
} else {
tmp2556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__2062, symshen_4out_1of_1bounds)
}
__typedArg0 := let__2062
__typedArg1 := symshen_4out_1of_1bounds
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2556 {
__e.TailApply(PrimFunc(symshen_4arg_1_6str), let__2061, V7022)
return
} else {
tmp2550 := Call(__e, PrimFunc(symshen_4arg_1_6str), let__2061, V7022)


tmp2551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V7021)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V7021
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()

tmp2553 := Call(__e, PrimFunc(symshen_4iter_1vector), V7020, tmp2551, V7022, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V7023)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V7023
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


tmp2554 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp2553)


__e.TailApply(PrimFunc(sym_8s), tmp2550, tmp2554)
return


}


}


}


}, 4)

tmp2561 := Call(__e, ns2_1set, symshen_4iter_1vector, tmp2542)


_ = tmp2561

tmp2562 := MakeNative(func(__e *ControlFlow) {
V7028 := __e.Get(1)
_ = V7028
tmp2563 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V7028)
}
__typedArg0 := V7028
return Call(__e, PrimFunc(symstr), __typedArg0)
})())
return
}, 0)

tmp2564 := MakeNative(func(__e *ControlFlow) {
Z7029 := __e.Get(1)
_ = Z7029
__e.TailApply(PrimFunc(symshen_4funexstring))
return
}, 1)

__e.TailApply(try_1catch, tmp2563, tmp2564)
return


}, 1)

tmp2565 := Call(__e, ns2_1set, symshen_4atom_1_6str, tmp2562)


_ = tmp2565

tmp2566 := MakeNative(func(__e *ControlFlow) {
tmp2567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString("x"))
}
__typedArg0 := MakeString("x")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp2568 := Call(__e, PrimFunc(symgensym), tmp2567)


tmp2569 := Call(__e, PrimFunc(symshen_4arg_1_6str), tmp2568, symshen_4a)


tmp2570 := Call(__e, PrimFunc(sym_8s), tmp2569, MakeString("\x11"))


tmp2571 := Call(__e, PrimFunc(sym_8s), MakeString("e"), tmp2570)


tmp2572 := Call(__e, PrimFunc(sym_8s), MakeString("n"), tmp2571)


tmp2573 := Call(__e, PrimFunc(sym_8s), MakeString("u"), tmp2572)


tmp2574 := Call(__e, PrimFunc(sym_8s), MakeString("f"), tmp2573)


__e.TailApply(PrimFunc(sym_8s), MakeString("\x10"), tmp2574)
return


}, 0)

tmp2575 := Call(__e, ns2_1set, symshen_4funexstring, tmp2566)


_ = tmp2575

tmp2576 := MakeNative(func(__e *ControlFlow) {
V7030 := __e.Get(1)
_ = V7030
tmp2580 := Call(__e, PrimFunc(symempty_2), V7030)


if True == tmp2580 {
__e.Return(True)
return
} else {
tmp2578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7030)
}
__typedArg0 := V7030
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2578 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

__e.TailApply(ns2_1set, symshen_4list_2, tmp2576)
return




}, 0)

