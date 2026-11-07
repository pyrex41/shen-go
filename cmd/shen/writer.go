package main

import . "github.com/pyrex41/shen-go/kl"

var WriterMain = MakeNative(func(__e *ControlFlow) {
tmp2080 := MakeNative(func(__e *ControlFlow) {
V7009 := __e.Get(1)
_ = V7009
tmp2081 := Call(__e, PrimFunc(symshen_4insert), V7009, MakeString("~S"))


let__2075 := tmp2081
_ = let__2075

tmp2082 := Call(__e, PrimFunc(symstoutput))


tmp2083 := Call(__e, PrimFunc(sympr), let__2075, tmp2082)


let__2076 := tmp2083
_ = let__2076

__e.Return(V7009)
return


}, 1)

tmp2084 := Call(__e, ns2_1set, symprint, tmp2080)


_ = tmp2084

tmp2085 := MakeNative(func(__e *ControlFlow) {
V7012 := __e.Get(1)
_ = V7012
V7013 := __e.Get(2)
_ = V7013
tmp2090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dhush_d)
}
__typedArg0 := sym_dhush_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

if True == tmp2090 {
__e.Return(V7012)
return
} else {
tmp2088 := Call(__e, PrimFunc(symchar_1stream_2), V7013)


if True == tmp2088 {
__e.TailApply(PrimFunc(symwrite_1string), V7012, V7013)
return
} else {
tmp2086 := Call(__e, PrimFunc(symshen_4string_1_6byte), V7012, MakeInteger(0))


__e.TailApply(PrimFunc(symshen_4write_1chars), V7012, V7013, tmp2086, MakeInteger(1))
return


}


}


}, 2)

tmp2091 := Call(__e, ns2_1set, sympr, tmp2085)


_ = tmp2091

tmp2092 := MakeNative(func(__e *ControlFlow) {
V7014 := __e.Get(1)
_ = V7014
V7015 := __e.Get(2)
_ = V7015
tmp2093 := MakeNative(func(__e *ControlFlow) {
tmp2094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sympos) {
return PrimPos(V7014, V7015)
}
__typedArg0 := V7014
__typedArg1 := V7015
return Call(__e, PrimFunc(sympos), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp2094)
}
__typedArg0 := tmp2094
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})())
return


}, 0)

tmp2095 := MakeNative(func(__e *ControlFlow) {
Z7016 := __e.Get(1)
_ = Z7016
__e.Return(symshen_4eos)
return
}, 1)

__e.TailApply(try_1catch, tmp2093, tmp2095)
return


}, 2)

tmp2096 := Call(__e, ns2_1set, symshen_4string_1_6byte, tmp2092)


_ = tmp2096

tmp2097 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__selftop:
V7017 := __self1
_ = V7017
V7018 := __self2
_ = V7018
V7019 := __self3
_ = V7019
V7020 := __self4
_ = V7020
tmp2102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4eos, V7019)
}
__typedArg0 := symshen_4eos
__typedArg1 := V7019
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2102 {
__e.Return(V7017)
return
} else {
tmp2098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symwrite_1byte) {
return PrimWriteByte(V7019, V7018)
}
__typedArg0 := V7019
__typedArg1 := V7018
return Call(__e, PrimFunc(symwrite_1byte), __typedArg0, __typedArg1)
})()

_ = tmp2098

tmp2099 := Call(__e, PrimFunc(symshen_4string_1_6byte), V7017, V7020)


if PrimFunc(symshen_4write_1chars) == __self {
__self1, __self2, __self3, __self4 = V7017, V7018, tmp2099, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V7020)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V7020
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4write_1chars), V7017, V7018, tmp2099, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V7020)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V7020
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


}


}, 4)

tmp2103 := Call(__e, ns2_1set, symshen_4write_1chars, tmp2097)


_ = tmp2103

tmp2104 := MakeNative(func(__e *ControlFlow) {
V7021 := __e.Get(1)
_ = V7021
V7022 := __e.Get(2)
_ = V7022
tmp2109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V7021)
}
__typedArg0 := V7021
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

if True == tmp2109 {
tmp2105 := Call(__e, PrimFunc(symshen_4proc_1nl), V7021)


__e.TailApply(PrimFunc(symshen_4mkstr_1l), tmp2105, V7022)
return


} else {
tmp2106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V7021, Nil)
}
__typedArg0 := V7021
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4proc_1nl, tmp2106)
}
__typedArg0 := symshen_4proc_1nl
__typedArg1 := tmp2106
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4mkstr_1r), tmp2107, V7022)
return


}


}, 2)

tmp2110 := Call(__e, ns2_1set, symshen_4mkstr, tmp2104)


_ = tmp2110

tmp2111 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V7027 := __self1
_ = V7027
V7028 := __self2
_ = V7028
tmp2118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V7028)
}
__typedArg0 := Nil
__typedArg1 := V7028
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2118 {
__e.Return(V7027)
return
} else {
tmp2116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7028)
}
__typedArg0 := V7028
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2116 {
tmp2112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7028)
}
__typedArg0 := V7028
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2113 := Call(__e, PrimFunc(symshen_4insert_1l), tmp2112, V7027)


tmp2114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7028)
}
__typedArg0 := V7028
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4mkstr_1l) == __self {
__self1, __self2 = tmp2113, tmp2114
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4mkstr_1l), tmp2113, tmp2114)
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

tmp2119 := Call(__e, ns2_1set, symshen_4mkstr_1l, tmp2111)


_ = tmp2119

tmp2120 := MakeNative(func(__e *ControlFlow) {
V7035 := __e.Get(1)
_ = V7035
V7036 := __e.Get(2)
_ = V7036
tmp2258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V7036)
}
__typedArg0 := MakeString("")
__typedArg1 := V7036
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2258 {
__e.Return(MakeString(""))
return
} else {
tmp2256 := Call(__e, PrimFunc(symshen_4_7string_2), V7036)


var ifres2243 Obj

if True == tmp2256 {
tmp2254 := Call(__e, PrimFunc(symhdstr), V7036)


tmp2255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2254)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2254
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2245 Obj

if True == tmp2255 {
tmp2253 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7036)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7036
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2247 Obj

if True == tmp2253 {
tmp2250 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7036)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7036
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("A"), tmp2250)
}
__typedArg0 := MakeString("A")
__typedArg1 := tmp2250
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2248 Obj

if True == tmp2251 {
ifres2248 = True


} else {
ifres2248 = False


}

ifres2247 = ifres2248


} else {
ifres2247 = False


}

var ifres2246 Obj

if True == ifres2247 {
ifres2246 = True


} else {
ifres2246 = False


}

ifres2245 = ifres2246


} else {
ifres2245 = False


}

var ifres2244 Obj

if True == ifres2245 {
ifres2244 = True


} else {
ifres2244 = False


}

ifres2243 = ifres2244


} else {
ifres2243 = False


}

if True == ifres2243 {
tmp2122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7036)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7036)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7036
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

tmp2123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4a, Nil)
}
__typedArg0 := symshen_4a
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2122, tmp2123)
}
__typedArg0 := tmp2122
__typedArg1 := tmp2123
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V7035, tmp2124)
}
__typedArg0 := V7035
__typedArg1 := tmp2124
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4app, tmp2125)
}
__typedArg0 := symshen_4app
__typedArg1 := tmp2125
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp2241 := Call(__e, PrimFunc(symshen_4_7string_2), V7036)


var ifres2228 Obj

if True == tmp2241 {
tmp2239 := Call(__e, PrimFunc(symhdstr), V7036)


tmp2240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2239)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2239
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2230 Obj

if True == tmp2240 {
tmp2238 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7036)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7036
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2232 Obj

if True == tmp2238 {
tmp2235 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7036)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7036
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("R"), tmp2235)
}
__typedArg0 := MakeString("R")
__typedArg1 := tmp2235
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2233 Obj

if True == tmp2236 {
ifres2233 = True


} else {
ifres2233 = False


}

ifres2232 = ifres2233


} else {
ifres2232 = False


}

var ifres2231 Obj

if True == ifres2232 {
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

if True == ifres2228 {
tmp2127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7036)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7036)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7036
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

tmp2128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4r, Nil)
}
__typedArg0 := symshen_4r
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2127, tmp2128)
}
__typedArg0 := tmp2127
__typedArg1 := tmp2128
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V7035, tmp2129)
}
__typedArg0 := V7035
__typedArg1 := tmp2129
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4app, tmp2130)
}
__typedArg0 := symshen_4app
__typedArg1 := tmp2130
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp2226 := Call(__e, PrimFunc(symshen_4_7string_2), V7036)


var ifres2213 Obj

if True == tmp2226 {
tmp2224 := Call(__e, PrimFunc(symhdstr), V7036)


tmp2225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2224)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2224
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2215 Obj

if True == tmp2225 {
tmp2223 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7036)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7036
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2217 Obj

if True == tmp2223 {
tmp2220 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7036)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7036
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("S"), tmp2220)
}
__typedArg0 := MakeString("S")
__typedArg1 := tmp2220
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2218 Obj

if True == tmp2221 {
ifres2218 = True


} else {
ifres2218 = False


}

ifres2217 = ifres2218


} else {
ifres2217 = False


}

var ifres2216 Obj

if True == ifres2217 {
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

if True == ifres2213 {
tmp2132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7036)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7036)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7036
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

tmp2133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4s, Nil)
}
__typedArg0 := symshen_4s
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2132, tmp2133)
}
__typedArg0 := tmp2132
__typedArg1 := tmp2133
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V7035, tmp2134)
}
__typedArg0 := V7035
__typedArg1 := tmp2134
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4app, tmp2135)
}
__typedArg0 := symshen_4app
__typedArg1 := tmp2135
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp2211 := Call(__e, PrimFunc(symshen_4_7string_2), V7036)


if True == tmp2211 {
tmp2136 := Call(__e, PrimFunc(symhdstr), V7036)


tmp2138 := Call(__e, PrimFunc(symshen_4insert_1l), V7035, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7036)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7036
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2138, Nil)
}
__typedArg0 := tmp2138
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2136, tmp2139)
}
__typedArg0 := tmp2136
__typedArg1 := tmp2139
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp2140)
}
__typedArg0 := symcn
__typedArg1 := tmp2140
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4factor_1cn), tmp2141)
return


} else {
tmp2209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7036)
}
__typedArg0 := V7036
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2190 Obj

if True == tmp2209 {
tmp2207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7036)
}
__typedArg0 := V7036
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcn, tmp2207)
}
__typedArg0 := symcn
__typedArg1 := tmp2207
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2192 Obj

if True == tmp2208 {
tmp2205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7036)
}
__typedArg0 := V7036
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2205)
}
__typedArg0 := tmp2205
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2194 Obj

if True == tmp2206 {
tmp2202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7036)
}
__typedArg0 := V7036
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2202)
}
__typedArg0 := tmp2202
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2203)
}
__typedArg0 := tmp2203
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2196 Obj

if True == tmp2204 {
tmp2198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7036)
}
__typedArg0 := V7036
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2198)
}
__typedArg0 := tmp2198
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2199)
}
__typedArg0 := tmp2199
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2200)
}
__typedArg0 := Nil
__typedArg1 := tmp2200
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2197 Obj

if True == tmp2201 {
ifres2197 = True


} else {
ifres2197 = False


}

ifres2196 = ifres2197


} else {
ifres2196 = False


}

var ifres2195 Obj

if True == ifres2196 {
ifres2195 = True


} else {
ifres2195 = False


}

ifres2194 = ifres2195


} else {
ifres2194 = False


}

var ifres2193 Obj

if True == ifres2194 {
ifres2193 = True


} else {
ifres2193 = False


}

ifres2192 = ifres2193


} else {
ifres2192 = False


}

var ifres2191 Obj

if True == ifres2192 {
ifres2191 = True


} else {
ifres2191 = False


}

ifres2190 = ifres2191


} else {
ifres2190 = False


}

if True == ifres2190 {
tmp2142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7036)
}
__typedArg0 := V7036
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2142)
}
__typedArg0 := tmp2142
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7036)
}
__typedArg0 := V7036
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2144)
}
__typedArg0 := tmp2144
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2145)
}
__typedArg0 := tmp2145
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2147 := Call(__e, PrimFunc(symshen_4insert_1l), V7035, tmp2146)


tmp2148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2147, Nil)
}
__typedArg0 := tmp2147
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2143, tmp2148)
}
__typedArg0 := tmp2143
__typedArg1 := tmp2148
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp2149)
}
__typedArg0 := symcn
__typedArg1 := tmp2149
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp2188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7036)
}
__typedArg0 := V7036
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2162 Obj

if True == tmp2188 {
tmp2186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7036)
}
__typedArg0 := V7036
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4app, tmp2186)
}
__typedArg0 := symshen_4app
__typedArg1 := tmp2186
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2164 Obj

if True == tmp2187 {
tmp2184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7036)
}
__typedArg0 := V7036
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2184)
}
__typedArg0 := tmp2184
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2166 Obj

if True == tmp2185 {
tmp2181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7036)
}
__typedArg0 := V7036
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2182)
}
__typedArg0 := tmp2182
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2168 Obj

if True == tmp2183 {
tmp2177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7036)
}
__typedArg0 := V7036
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2177)
}
__typedArg0 := tmp2177
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2178)
}
__typedArg0 := tmp2178
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2179)
}
__typedArg0 := tmp2179
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2170 Obj

if True == tmp2180 {
tmp2172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7036)
}
__typedArg0 := V7036
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2172)
}
__typedArg0 := tmp2172
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2173)
}
__typedArg0 := tmp2173
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2174)
}
__typedArg0 := tmp2174
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2175)
}
__typedArg0 := Nil
__typedArg1 := tmp2175
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2171 Obj

if True == tmp2176 {
ifres2171 = True


} else {
ifres2171 = False


}

ifres2170 = ifres2171


} else {
ifres2170 = False


}

var ifres2169 Obj

if True == ifres2170 {
ifres2169 = True


} else {
ifres2169 = False


}

ifres2168 = ifres2169


} else {
ifres2168 = False


}

var ifres2167 Obj

if True == ifres2168 {
ifres2167 = True


} else {
ifres2167 = False


}

ifres2166 = ifres2167


} else {
ifres2166 = False


}

var ifres2165 Obj

if True == ifres2166 {
ifres2165 = True


} else {
ifres2165 = False


}

ifres2164 = ifres2165


} else {
ifres2164 = False


}

var ifres2163 Obj

if True == ifres2164 {
ifres2163 = True


} else {
ifres2163 = False


}

ifres2162 = ifres2163


} else {
ifres2162 = False


}

if True == ifres2162 {
tmp2150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7036)
}
__typedArg0 := V7036
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2150)
}
__typedArg0 := tmp2150
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7036)
}
__typedArg0 := V7036
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2152)
}
__typedArg0 := tmp2152
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2153)
}
__typedArg0 := tmp2153
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2155 := Call(__e, PrimFunc(symshen_4insert_1l), V7035, tmp2154)


tmp2156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7036)
}
__typedArg0 := V7036
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2155, tmp2158)
}
__typedArg0 := tmp2155
__typedArg1 := tmp2158
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2151, tmp2159)
}
__typedArg0 := tmp2151
__typedArg1 := tmp2159
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4app, tmp2160)
}
__typedArg0 := symshen_4app
__typedArg1 := tmp2160
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

tmp2259 := Call(__e, ns2_1set, symshen_4insert_1l, tmp2120)


_ = tmp2259

tmp2260 := MakeNative(func(__e *ControlFlow) {
V7037 := __e.Get(1)
_ = V7037
tmp2345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7037)
}
__typedArg0 := V7037
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2276 Obj

if True == tmp2345 {
tmp2343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7037)
}
__typedArg0 := V7037
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcn, tmp2343)
}
__typedArg0 := symcn
__typedArg1 := tmp2343
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2278 Obj

if True == tmp2344 {
tmp2341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7037)
}
__typedArg0 := V7037
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2341)
}
__typedArg0 := tmp2341
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2280 Obj

if True == tmp2342 {
tmp2338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7037)
}
__typedArg0 := V7037
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2338)
}
__typedArg0 := tmp2338
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2339)
}
__typedArg0 := tmp2339
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2282 Obj

if True == tmp2340 {
tmp2334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7037)
}
__typedArg0 := V7037
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2334)
}
__typedArg0 := tmp2334
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2335)
}
__typedArg0 := tmp2335
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2336)
}
__typedArg0 := tmp2336
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2284 Obj

if True == tmp2337 {
tmp2329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7037)
}
__typedArg0 := V7037
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2329)
}
__typedArg0 := tmp2329
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2330)
}
__typedArg0 := tmp2330
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2331)
}
__typedArg0 := tmp2331
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcn, tmp2332)
}
__typedArg0 := symcn
__typedArg1 := tmp2332
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2286 Obj

if True == tmp2333 {
tmp2324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7037)
}
__typedArg0 := V7037
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2324)
}
__typedArg0 := tmp2324
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2325)
}
__typedArg0 := tmp2325
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2326)
}
__typedArg0 := tmp2326
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2327)
}
__typedArg0 := tmp2327
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2288 Obj

if True == tmp2328 {
tmp2318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7037)
}
__typedArg0 := V7037
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2318)
}
__typedArg0 := tmp2318
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2319)
}
__typedArg0 := tmp2319
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2320)
}
__typedArg0 := tmp2320
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

var ifres2290 Obj

if True == tmp2323 {
tmp2311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7037)
}
__typedArg0 := V7037
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2311)
}
__typedArg0 := tmp2311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2312)
}
__typedArg0 := tmp2312
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2313)
}
__typedArg0 := tmp2313
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2314)
}
__typedArg0 := tmp2314
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2315)
}
__typedArg0 := tmp2315
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2316)
}
__typedArg0 := Nil
__typedArg1 := tmp2316
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2292 Obj

if True == tmp2317 {
tmp2307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7037)
}
__typedArg0 := V7037
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2308)
}
__typedArg0 := tmp2308
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2309)
}
__typedArg0 := Nil
__typedArg1 := tmp2309
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2294 Obj

if True == tmp2310 {
tmp2304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7037)
}
__typedArg0 := V7037
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2304)
}
__typedArg0 := tmp2304
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(tmp2305)
}
__typedArg0 := tmp2305
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres2296 Obj

if True == tmp2306 {
tmp2298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7037)
}
__typedArg0 := V7037
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2299)
}
__typedArg0 := tmp2299
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2300)
}
__typedArg0 := tmp2300
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2301)
}
__typedArg0 := tmp2301
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(tmp2302)
}
__typedArg0 := tmp2302
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres2297 Obj

if True == tmp2303 {
ifres2297 = True


} else {
ifres2297 = False


}

ifres2296 = ifres2297


} else {
ifres2296 = False


}

var ifres2295 Obj

if True == ifres2296 {
ifres2295 = True


} else {
ifres2295 = False


}

ifres2294 = ifres2295


} else {
ifres2294 = False


}

var ifres2293 Obj

if True == ifres2294 {
ifres2293 = True


} else {
ifres2293 = False


}

ifres2292 = ifres2293


} else {
ifres2292 = False


}

var ifres2291 Obj

if True == ifres2292 {
ifres2291 = True


} else {
ifres2291 = False


}

ifres2290 = ifres2291


} else {
ifres2290 = False


}

var ifres2289 Obj

if True == ifres2290 {
ifres2289 = True


} else {
ifres2289 = False


}

ifres2288 = ifres2289


} else {
ifres2288 = False


}

var ifres2287 Obj

if True == ifres2288 {
ifres2287 = True


} else {
ifres2287 = False


}

ifres2286 = ifres2287


} else {
ifres2286 = False


}

var ifres2285 Obj

if True == ifres2286 {
ifres2285 = True


} else {
ifres2285 = False


}

ifres2284 = ifres2285


} else {
ifres2284 = False


}

var ifres2283 Obj

if True == ifres2284 {
ifres2283 = True


} else {
ifres2283 = False


}

ifres2282 = ifres2283


} else {
ifres2282 = False


}

var ifres2281 Obj

if True == ifres2282 {
ifres2281 = True


} else {
ifres2281 = False


}

ifres2280 = ifres2281


} else {
ifres2280 = False


}

var ifres2279 Obj

if True == ifres2280 {
ifres2279 = True


} else {
ifres2279 = False


}

ifres2278 = ifres2279


} else {
ifres2278 = False


}

var ifres2277 Obj

if True == ifres2278 {
ifres2277 = True


} else {
ifres2277 = False


}

ifres2276 = ifres2277


} else {
ifres2276 = False


}

if True == ifres2276 {
tmp2261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7037)
}
__typedArg0 := V7037
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2261)
}
__typedArg0 := tmp2261
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7037)
}
__typedArg0 := V7037
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2263)
}
__typedArg0 := tmp2263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2264)
}
__typedArg0 := tmp2264
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2265)
}
__typedArg0 := tmp2265
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2266)
}
__typedArg0 := tmp2266
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp2262)
__typedS1, __typedOK1 := TypedString(tmp2267)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp2262
__typedArg1 := tmp2267
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp2269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7037)
}
__typedArg0 := V7037
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2269)
}
__typedArg0 := tmp2269
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2270)
}
__typedArg0 := tmp2270
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2271)
}
__typedArg0 := tmp2271
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2272)
}
__typedArg0 := tmp2272
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2268, tmp2273)
}
__typedArg0 := tmp2268
__typedArg1 := tmp2273
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp2274)
}
__typedArg0 := symcn
__typedArg1 := tmp2274
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V7037)
return
}


}, 1)

tmp2346 := Call(__e, ns2_1set, symshen_4factor_1cn, tmp2260)


_ = tmp2346

tmp2347 := MakeNative(func(__e *ControlFlow) {
V7040 := __e.Get(1)
_ = V7040
tmp2373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V7040)
}
__typedArg0 := MakeString("")
__typedArg1 := V7040
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2373 {
__e.Return(MakeString(""))
return
} else {
tmp2371 := Call(__e, PrimFunc(symshen_4_7string_2), V7040)


var ifres2358 Obj

if True == tmp2371 {
tmp2369 := Call(__e, PrimFunc(symhdstr), V7040)


tmp2370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2369)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2369
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2360 Obj

if True == tmp2370 {
tmp2368 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7040)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7040
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2362 Obj

if True == tmp2368 {
tmp2365 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7040)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7040
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("%"), tmp2365)
}
__typedArg0 := MakeString("%")
__typedArg1 := tmp2365
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2363 Obj

if True == tmp2366 {
ifres2363 = True


} else {
ifres2363 = False


}

ifres2362 = ifres2363


} else {
ifres2362 = False


}

var ifres2361 Obj

if True == ifres2362 {
ifres2361 = True


} else {
ifres2361 = False


}

ifres2360 = ifres2361


} else {
ifres2360 = False


}

var ifres2359 Obj

if True == ifres2360 {
ifres2359 = True


} else {
ifres2359 = False


}

ifres2358 = ifres2359


} else {
ifres2358 = False


}

if True == ifres2358 {
tmp2348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(MakeInteger(10))
}
__typedArg0 := MakeInteger(10)
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp2351 := Call(__e, PrimFunc(symshen_4proc_1nl), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7040)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7040)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7040
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp2348)
__typedS1, __typedOK1 := TypedString(tmp2351)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp2348
__typedArg1 := tmp2351
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
tmp2356 := Call(__e, PrimFunc(symshen_4_7string_2), V7040)


if True == tmp2356 {
tmp2352 := Call(__e, PrimFunc(symhdstr), V7040)


tmp2354 := Call(__e, PrimFunc(symshen_4proc_1nl), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7040)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7040
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp2352)
__typedS1, __typedOK1 := TypedString(tmp2354)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp2352
__typedArg1 := tmp2354
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

tmp2374 := Call(__e, ns2_1set, symshen_4proc_1nl, tmp2347)


_ = tmp2374

tmp2375 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V7045 := __self1
_ = V7045
V7046 := __self2
_ = V7046
tmp2384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V7046)
}
__typedArg0 := Nil
__typedArg1 := V7046
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2384 {
__e.Return(V7045)
return
} else {
tmp2382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7046)
}
__typedArg0 := V7046
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2382 {
tmp2376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7046)
}
__typedArg0 := V7046
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V7045, Nil)
}
__typedArg0 := V7045
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2376, tmp2377)
}
__typedArg0 := tmp2376
__typedArg1 := tmp2377
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4insert, tmp2378)
}
__typedArg0 := symshen_4insert
__typedArg1 := tmp2378
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7046)
}
__typedArg0 := V7046
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4mkstr_1r) == __self {
__self1, __self2 = tmp2379, tmp2380
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4mkstr_1r), tmp2379, tmp2380)
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

tmp2385 := Call(__e, ns2_1set, symshen_4mkstr_1r, tmp2375)


_ = tmp2385

tmp2386 := MakeNative(func(__e *ControlFlow) {
V7047 := __e.Get(1)
_ = V7047
V7048 := __e.Get(2)
_ = V7048
__e.TailApply(PrimFunc(symshen_4insert_1h), V7047, V7048, MakeString(""))
return
}, 2)

tmp2387 := Call(__e, ns2_1set, symshen_4insert, tmp2386)


_ = tmp2387

tmp2388 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__selftop:
V7057 := __self1
_ = V7057
V7058 := __self2
_ = V7058
V7059 := __self3
_ = V7059
tmp2449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V7058)
}
__typedArg0 := MakeString("")
__typedArg1 := V7058
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2449 {
__e.Return(V7059)
return
} else {
tmp2447 := Call(__e, PrimFunc(symshen_4_7string_2), V7058)


var ifres2434 Obj

if True == tmp2447 {
tmp2445 := Call(__e, PrimFunc(symhdstr), V7058)


tmp2446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2445)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2445
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2436 Obj

if True == tmp2446 {
tmp2444 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7058)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7058
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2438 Obj

if True == tmp2444 {
tmp2441 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7058)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7058
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("A"), tmp2441)
}
__typedArg0 := MakeString("A")
__typedArg1 := tmp2441
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2439 Obj

if True == tmp2442 {
ifres2439 = True


} else {
ifres2439 = False


}

ifres2438 = ifres2439


} else {
ifres2438 = False


}

var ifres2437 Obj

if True == ifres2438 {
ifres2437 = True


} else {
ifres2437 = False


}

ifres2436 = ifres2437


} else {
ifres2436 = False


}

var ifres2435 Obj

if True == ifres2436 {
ifres2435 = True


} else {
ifres2435 = False


}

ifres2434 = ifres2435


} else {
ifres2434 = False


}

if True == ifres2434 {
tmp2391 := Call(__e, PrimFunc(symshen_4app), V7057, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7058)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7058)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7058
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})(), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(V7059)
__typedS1, __typedOK1 := TypedString(tmp2391)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V7059
__typedArg1 := tmp2391
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
tmp2432 := Call(__e, PrimFunc(symshen_4_7string_2), V7058)


var ifres2419 Obj

if True == tmp2432 {
tmp2430 := Call(__e, PrimFunc(symhdstr), V7058)


tmp2431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2430)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2430
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2421 Obj

if True == tmp2431 {
tmp2429 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7058)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7058
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2423 Obj

if True == tmp2429 {
tmp2426 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7058)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7058
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("R"), tmp2426)
}
__typedArg0 := MakeString("R")
__typedArg1 := tmp2426
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2424 Obj

if True == tmp2427 {
ifres2424 = True


} else {
ifres2424 = False


}

ifres2423 = ifres2424


} else {
ifres2423 = False


}

var ifres2422 Obj

if True == ifres2423 {
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

if True == ifres2419 {
tmp2394 := Call(__e, PrimFunc(symshen_4app), V7057, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7058)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7058)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7058
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})(), symshen_4r)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(V7059)
__typedS1, __typedOK1 := TypedString(tmp2394)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V7059
__typedArg1 := tmp2394
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
tmp2417 := Call(__e, PrimFunc(symshen_4_7string_2), V7058)


var ifres2404 Obj

if True == tmp2417 {
tmp2415 := Call(__e, PrimFunc(symhdstr), V7058)


tmp2416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2415)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2415
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2406 Obj

if True == tmp2416 {
tmp2414 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7058)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7058
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2408 Obj

if True == tmp2414 {
tmp2411 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7058)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7058
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("S"), tmp2411)
}
__typedArg0 := MakeString("S")
__typedArg1 := tmp2411
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2409 Obj

if True == tmp2412 {
ifres2409 = True


} else {
ifres2409 = False


}

ifres2408 = ifres2409


} else {
ifres2408 = False


}

var ifres2407 Obj

if True == ifres2408 {
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

if True == ifres2404 {
tmp2397 := Call(__e, PrimFunc(symshen_4app), V7057, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7058)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7058)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7058
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})(), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(V7059)
__typedS1, __typedOK1 := TypedString(tmp2397)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V7059
__typedArg1 := tmp2397
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
tmp2402 := Call(__e, PrimFunc(symshen_4_7string_2), V7058)


if True == tmp2402 {
tmp2398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7058)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V7058
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

tmp2399 := Call(__e, PrimFunc(symhdstr), V7058)


if PrimFunc(symshen_4insert_1h) == __self {
__self1, __self2, __self3 = V7057, tmp2398, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(V7059)
__typedS1, __typedOK1 := TypedString(tmp2399)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V7059
__typedArg1 := tmp2399
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4insert_1h), V7057, tmp2398, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(V7059)
__typedS1, __typedOK1 := TypedString(tmp2399)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V7059
__typedArg1 := tmp2399
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

tmp2450 := Call(__e, ns2_1set, symshen_4insert_1h, tmp2388)


_ = tmp2450

tmp2451 := MakeNative(func(__e *ControlFlow) {
V7060 := __e.Get(1)
_ = V7060
V7061 := __e.Get(2)
_ = V7061
V7062 := __e.Get(3)
_ = V7062
tmp2452 := Call(__e, PrimFunc(symshen_4arg_1_6str), V7060, V7062)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp2452)
__typedS1, __typedOK1 := TypedString(V7061)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp2452
__typedArg1 := V7061
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 3)

tmp2453 := Call(__e, ns2_1set, symshen_4app, tmp2451)


_ = tmp2453

tmp2454 := MakeNative(func(__e *ControlFlow) {
V7066 := __e.Get(1)
_ = V7066
V7067 := __e.Get(2)
_ = V7067
tmp2462 := Call(__e, PrimFunc(symfail))


tmp2463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V7066, tmp2462)
}
__typedArg0 := V7066
__typedArg1 := tmp2462
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2463 {
__e.Return(MakeString("..."))
return
} else {
tmp2460 := Call(__e, PrimFunc(symshen_4list_2), V7066)


if True == tmp2460 {
__e.TailApply(PrimFunc(symshen_4list_1_6str), V7066, V7067)
return
} else {
tmp2458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V7066)
}
__typedArg0 := V7066
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

if True == tmp2458 {
__e.TailApply(PrimFunc(symshen_4str_1_6str), V7066, V7067)
return
} else {
tmp2456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V7066)
}
__typedArg0 := V7066
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp2456 {
__e.TailApply(PrimFunc(symshen_4vector_1_6str), V7066, V7067)
return
} else {
__e.TailApply(PrimFunc(symshen_4atom_1_6str), V7066)
return
}


}


}


}


}, 2)

tmp2464 := Call(__e, ns2_1set, symshen_4arg_1_6str, tmp2454)


_ = tmp2464

tmp2465 := MakeNative(func(__e *ControlFlow) {
V7068 := __e.Get(1)
_ = V7068
V7069 := __e.Get(2)
_ = V7069
tmp2473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4r, V7069)
}
__typedArg0 := symshen_4r
__typedArg1 := V7069
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2473 {
tmp2466 := Call(__e, PrimFunc(symshen_4maxseq))


tmp2467 := Call(__e, PrimFunc(symshen_4iter_1list), V7068, symshen_4r, tmp2466)


tmp2468 := Call(__e, PrimFunc(sym_8s), tmp2467, MakeString(")"))


__e.TailApply(PrimFunc(sym_8s), MakeString("("), tmp2468)
return


} else {
tmp2469 := Call(__e, PrimFunc(symshen_4maxseq))


tmp2470 := Call(__e, PrimFunc(symshen_4iter_1list), V7068, V7069, tmp2469)


tmp2471 := Call(__e, PrimFunc(sym_8s), tmp2470, MakeString("]"))


__e.TailApply(PrimFunc(sym_8s), MakeString("["), tmp2471)
return


}


}, 2)

tmp2474 := Call(__e, ns2_1set, symshen_4list_1_6str, tmp2465)


_ = tmp2474

tmp2475 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmaximum_1print_1sequence_1size_d)
}
__typedArg0 := sym_dmaximum_1print_1sequence_1size_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp2476 := Call(__e, ns2_1set, symshen_4maxseq, tmp2475)


_ = tmp2476

tmp2477 := MakeNative(func(__e *ControlFlow) {
V7080 := __e.Get(1)
_ = V7080
V7081 := __e.Get(2)
_ = V7081
V7082 := __e.Get(3)
_ = V7082
tmp2498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V7080)
}
__typedArg0 := Nil
__typedArg1 := V7080
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2498 {
__e.Return(MakeString(""))
return
} else {
tmp2496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V7082)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V7082
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2496 {
__e.Return(MakeString("... etc"))
return
} else {
tmp2494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7080)
}
__typedArg0 := V7080
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2490 Obj

if True == tmp2494 {
tmp2492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7080)
}
__typedArg0 := V7080
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2492)
}
__typedArg0 := Nil
__typedArg1 := tmp2492
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2491 Obj

if True == tmp2493 {
ifres2491 = True


} else {
ifres2491 = False


}

ifres2490 = ifres2491


} else {
ifres2490 = False


}

if True == ifres2490 {
tmp2478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7080)
}
__typedArg0 := V7080
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4arg_1_6str), tmp2478, V7081)
return


} else {
tmp2488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7080)
}
__typedArg0 := V7080
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2488 {
tmp2479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V7080)
}
__typedArg0 := V7080
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2480 := Call(__e, PrimFunc(symshen_4arg_1_6str), tmp2479, V7081)


tmp2481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V7080)
}
__typedArg0 := V7080
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2483 := Call(__e, PrimFunc(symshen_4iter_1list), tmp2481, V7081, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V7082)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V7082
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


tmp2484 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp2483)


__e.TailApply(PrimFunc(sym_8s), tmp2480, tmp2484)
return


} else {
tmp2485 := Call(__e, PrimFunc(symshen_4arg_1_6str), V7080, V7081)


tmp2486 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp2485)


__e.TailApply(PrimFunc(sym_8s), MakeString("|"), tmp2486)
return


}


}


}


}


}, 3)

tmp2499 := Call(__e, ns2_1set, symshen_4iter_1list, tmp2477)


_ = tmp2499

tmp2500 := MakeNative(func(__e *ControlFlow) {
V7085 := __e.Get(1)
_ = V7085
V7086 := __e.Get(2)
_ = V7086
tmp2505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4a, V7086)
}
__typedArg0 := symshen_4a
__typedArg1 := V7086
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2505 {
__e.Return(V7085)
return
} else {
tmp2501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(MakeInteger(34))
}
__typedArg0 := MakeInteger(34)
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp2502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(MakeInteger(34))
}
__typedArg0 := MakeInteger(34)
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp2503 := Call(__e, PrimFunc(sym_8s), V7085, tmp2502)


__e.TailApply(PrimFunc(sym_8s), tmp2501, tmp2503)
return


}


}, 2)

tmp2506 := Call(__e, ns2_1set, symshen_4str_1_6str, tmp2500)


_ = tmp2506

tmp2507 := MakeNative(func(__e *ControlFlow) {
V7087 := __e.Get(1)
_ = V7087
V7088 := __e.Get(2)
_ = V7088
tmp2520 := Call(__e, PrimFunc(symshen_4print_1vector_2), V7087)


if True == tmp2520 {
tmp2508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V7087, MakeInteger(0))
}
__typedArg0 := V7087
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp2509 := Call(__e, PrimFunc(symfn), tmp2508)


__e.TailApply(tmp2509, V7087)
return


} else {
tmp2518 := Call(__e, PrimFunc(symvector_2), V7087)


if True == tmp2518 {
tmp2510 := Call(__e, PrimFunc(symshen_4maxseq))


tmp2511 := Call(__e, PrimFunc(symshen_4iter_1vector), V7087, MakeInteger(1), V7088, tmp2510)


tmp2512 := Call(__e, PrimFunc(sym_8s), tmp2511, MakeString(">"))


__e.TailApply(PrimFunc(sym_8s), MakeString("<"), tmp2512)
return


} else {
tmp2513 := Call(__e, PrimFunc(symshen_4maxseq))


tmp2514 := Call(__e, PrimFunc(symshen_4iter_1vector), V7087, MakeInteger(0), V7088, tmp2513)


tmp2515 := Call(__e, PrimFunc(sym_8s), tmp2514, MakeString(">>"))


tmp2516 := Call(__e, PrimFunc(sym_8s), MakeString("<"), tmp2515)


__e.TailApply(PrimFunc(sym_8s), MakeString("<"), tmp2516)
return


}


}


}, 2)

tmp2521 := Call(__e, ns2_1set, symshen_4vector_1_6str, tmp2507)


_ = tmp2521

tmp2522 := MakeNative(func(__e *ControlFlow) {
V7089 := __e.Get(1)
_ = V7089
tmp2523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dempty_1absvector_d)
}
__typedArg0 := symshen_4_dempty_1absvector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V7089, tmp2523)
}
__typedArg0 := V7089
__typedArg1 := tmp2523
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp2524 := Call(__e, ns2_1set, symshen_4empty_1absvector_2, tmp2522)


_ = tmp2524

tmp2525 := MakeNative(func(__e *ControlFlow) {
V7090 := __e.Get(1)
_ = V7090
tmp2543 := Call(__e, PrimFunc(symshen_4empty_1absvector_2), V7090)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp2543)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp2543
return Call(__e, PrimFunc(symnot), __typedArg0)
})() {
tmp2527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V7090, MakeInteger(0))
}
__typedArg0 := V7090
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

let__2077 := tmp2527
_ = let__2077

tmp2541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__2077, symshen_4tuple)
}
__typedArg0 := let__2077
__typedArg1 := symshen_4tuple
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2528 Obj

if True == tmp2541 {
ifres2528 = True


} else {
tmp2540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__2077, symshen_4pvar)
}
__typedArg0 := let__2077
__typedArg1 := symshen_4pvar
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2530 Obj

if True == tmp2540 {
ifres2530 = True


} else {
tmp2539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__2077, symshen_4dictionary)
}
__typedArg0 := let__2077
__typedArg1 := symshen_4dictionary
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2532 Obj

if True == tmp2539 {
ifres2532 = True


} else {
tmp2537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(let__2077)
}
__typedArg0 := let__2077
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

tmp2538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp2537)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp2537
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres2534 Obj

if True == tmp2538 {
tmp2536 := Call(__e, PrimFunc(symshen_4fbound_2), let__2077)


var ifres2535 Obj

if True == tmp2536 {
ifres2535 = True


} else {
ifres2535 = False


}

ifres2534 = ifres2535


} else {
ifres2534 = False


}

var ifres2533 Obj

if True == ifres2534 {
ifres2533 = True


} else {
ifres2533 = False


}

ifres2532 = ifres2533


}

var ifres2531 Obj

if True == ifres2532 {
ifres2531 = True


} else {
ifres2531 = False


}

ifres2530 = ifres2531


}

var ifres2529 Obj

if True == ifres2530 {
ifres2529 = True


} else {
ifres2529 = False


}

ifres2528 = ifres2529


}

if True == ifres2528 {
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

tmp2545 := Call(__e, ns2_1set, symshen_4print_1vector_2, tmp2525)


_ = tmp2545

tmp2546 := MakeNative(func(__e *ControlFlow) {
V7092 := __e.Get(1)
_ = V7092
tmp2547 := Call(__e, PrimFunc(symarity), V7092)


tmp2548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp2547, MakeInteger(-1))
}
__typedArg0 := tmp2547
__typedArg1 := MakeInteger(-1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp2548)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp2548
return Call(__e, PrimFunc(symnot), __typedArg0)
})())
return


}, 1)

tmp2549 := Call(__e, ns2_1set, symshen_4fbound_2, tmp2546)


_ = tmp2549

tmp2550 := MakeNative(func(__e *ControlFlow) {
V7093 := __e.Get(1)
_ = V7093
tmp2551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V7093, MakeInteger(1))
}
__typedArg0 := V7093
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp2552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V7093, MakeInteger(2))
}
__typedArg0 := V7093
__typedArg1 := MakeInteger(2)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp2553 := Call(__e, PrimFunc(symshen_4app), tmp2552, MakeString(")"), symshen_4s)


tmp2555 := Call(__e, PrimFunc(symshen_4app), tmp2551, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" "))
__typedS1, __typedOK1 := TypedString(tmp2553)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" ")
__typedArg1 := tmp2553
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("(@p "))
__typedS1, __typedOK1 := TypedString(tmp2555)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("(@p ")
__typedArg1 := tmp2555
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp2556 := Call(__e, ns2_1set, symshen_4tuple, tmp2550)


_ = tmp2556

tmp2557 := MakeNative(func(__e *ControlFlow) {
V7094 := __e.Get(1)
_ = V7094
__e.Return(MakeString("(dict ...)"))
return
}, 1)

tmp2558 := Call(__e, ns2_1set, symshen_4dictionary, tmp2557)


_ = tmp2558

tmp2559 := MakeNative(func(__e *ControlFlow) {
V7101 := __e.Get(1)
_ = V7101
V7102 := __e.Get(2)
_ = V7102
V7103 := __e.Get(3)
_ = V7103
V7104 := __e.Get(4)
_ = V7104
tmp2577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V7104)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V7104
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2577 {
__e.Return(MakeString("... etc"))
return
} else {
tmp2560 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V7101, V7102)
}
__typedArg0 := V7101
__typedArg1 := V7102
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 0)

tmp2561 := MakeNative(func(__e *ControlFlow) {
Z7106 := __e.Get(1)
_ = Z7106
__e.Return(symshen_4out_1of_1bounds)
return
}, 1)

tmp2562 := Call(__e, try_1catch, tmp2560, tmp2561)


let__2078 := tmp2562
_ = let__2078

tmp2563 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V7101, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V7102)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V7102
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := V7101
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V7102)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V7102
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp2565 := MakeNative(func(__e *ControlFlow) {
Z7108 := __e.Get(1)
_ = Z7108
__e.Return(symshen_4out_1of_1bounds)
return
}, 1)

tmp2566 := Call(__e, try_1catch, tmp2563, tmp2565)


let__2079 := tmp2566
_ = let__2079

tmp2575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__2078, symshen_4out_1of_1bounds)
}
__typedArg0 := let__2078
__typedArg1 := symshen_4out_1of_1bounds
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2575 {
__e.Return(MakeString(""))
return
} else {
tmp2573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__2079, symshen_4out_1of_1bounds)
}
__typedArg0 := let__2079
__typedArg1 := symshen_4out_1of_1bounds
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2573 {
__e.TailApply(PrimFunc(symshen_4arg_1_6str), let__2078, V7103)
return
} else {
tmp2567 := Call(__e, PrimFunc(symshen_4arg_1_6str), let__2078, V7103)


tmp2568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V7102)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V7102
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()

tmp2570 := Call(__e, PrimFunc(symshen_4iter_1vector), V7101, tmp2568, V7103, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V7104)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V7104
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


tmp2571 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp2570)


__e.TailApply(PrimFunc(sym_8s), tmp2567, tmp2571)
return


}


}


}


}, 4)

tmp2578 := Call(__e, ns2_1set, symshen_4iter_1vector, tmp2559)


_ = tmp2578

tmp2579 := MakeNative(func(__e *ControlFlow) {
V7109 := __e.Get(1)
_ = V7109
tmp2580 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V7109)
}
__typedArg0 := V7109
return Call(__e, PrimFunc(symstr), __typedArg0)
})())
return
}, 0)

tmp2581 := MakeNative(func(__e *ControlFlow) {
Z7110 := __e.Get(1)
_ = Z7110
__e.TailApply(PrimFunc(symshen_4funexstring))
return
}, 1)

__e.TailApply(try_1catch, tmp2580, tmp2581)
return


}, 1)

tmp2582 := Call(__e, ns2_1set, symshen_4atom_1_6str, tmp2579)


_ = tmp2582

tmp2583 := MakeNative(func(__e *ControlFlow) {
tmp2584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString("x"))
}
__typedArg0 := MakeString("x")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp2585 := Call(__e, PrimFunc(symgensym), tmp2584)


tmp2586 := Call(__e, PrimFunc(symshen_4arg_1_6str), tmp2585, symshen_4a)


tmp2587 := Call(__e, PrimFunc(sym_8s), tmp2586, MakeString("\x11"))


tmp2588 := Call(__e, PrimFunc(sym_8s), MakeString("e"), tmp2587)


tmp2589 := Call(__e, PrimFunc(sym_8s), MakeString("n"), tmp2588)


tmp2590 := Call(__e, PrimFunc(sym_8s), MakeString("u"), tmp2589)


tmp2591 := Call(__e, PrimFunc(sym_8s), MakeString("f"), tmp2590)


__e.TailApply(PrimFunc(sym_8s), MakeString("\x10"), tmp2591)
return


}, 0)

tmp2592 := Call(__e, ns2_1set, symshen_4funexstring, tmp2583)


_ = tmp2592

tmp2593 := MakeNative(func(__e *ControlFlow) {
V7111 := __e.Get(1)
_ = V7111
tmp2597 := Call(__e, PrimFunc(symempty_2), V7111)


if True == tmp2597 {
__e.Return(True)
return
} else {
tmp2595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V7111)
}
__typedArg0 := V7111
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2595 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

__e.TailApply(ns2_1set, symshen_4list_2, tmp2593)
return




}, 0)

