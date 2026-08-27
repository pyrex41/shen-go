package main

import . "github.com/pyrex41/shen-go/kl"

var WriterMain = MakeNative(func(__e *ControlFlow) {
tmp2104 := MakeNative(func(__e *ControlFlow) {
V5602 := __e.Get(1)
_ = V5602
tmp2105 := MakeNative(func(__e *ControlFlow) {
W5603 := __e.Get(1)
_ = W5603
tmp2106 := MakeNative(func(__e *ControlFlow) {
W5604 := __e.Get(1)
_ = W5604
__e.Return(V5602)
return
}, 1)

tmp2107 := Call(__e, PrimFunc(symstoutput))


tmp2108 := Call(__e, PrimFunc(sympr), W5603, tmp2107)


__e.TailApply(tmp2106, tmp2108)
return


}, 1)

tmp2109 := Call(__e, PrimFunc(symshen_4insert), V5602, MakeString("~S"))


__e.TailApply(tmp2105, tmp2109)
return


}, 1)

tmp2110 := Call(__e, ns2_1set, symprint, tmp2104)


_ = tmp2110

tmp2111 := MakeNative(func(__e *ControlFlow) {
V5605 := __e.Get(1)
_ = V5605
V5606 := __e.Get(2)
_ = V5606
tmp2116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dhush_d)
}
__typedArg0 := sym_dhush_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

if True == tmp2116 {
__e.Return(V5605)
return
} else {
tmp2114 := Call(__e, PrimFunc(symshen_4char_1stoutput_2), V5606)


if True == tmp2114 {
__e.TailApply(PrimFunc(symshen_4write_1string), V5605, V5606)
return
} else {
tmp2112 := Call(__e, PrimFunc(symshen_4string_1_6byte), V5605, MakeNumber(0))


__e.TailApply(PrimFunc(symshen_4write_1chars), V5605, V5606, tmp2112, MakeNumber(1))
return


}


}


}, 2)

tmp2117 := Call(__e, ns2_1set, sympr, tmp2111)


_ = tmp2117

tmp2118 := MakeNative(func(__e *ControlFlow) {
V5607 := __e.Get(1)
_ = V5607
V5608 := __e.Get(2)
_ = V5608
tmp2119 := MakeNative(func(__e *ControlFlow) {
tmp2120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sympos) {
return PrimPos(V5607, V5608)
}
__typedArg0 := V5607
__typedArg1 := V5608
return Call(__e, PrimFunc(sympos), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp2120)
}
__typedArg0 := tmp2120
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})())
return


}, 0)

tmp2121 := MakeNative(func(__e *ControlFlow) {
Z5609 := __e.Get(1)
_ = Z5609
__e.Return(symshen_4eos)
return
}, 1)

__e.TailApply(try_1catch, tmp2119, tmp2121)
return


}, 2)

tmp2122 := Call(__e, ns2_1set, symshen_4string_1_6byte, tmp2118)


_ = tmp2122

tmp2123 := MakeNative(func(__e *ControlFlow) {
V5610 := __e.Get(1)
_ = V5610
V5611 := __e.Get(2)
_ = V5611
V5612 := __e.Get(3)
_ = V5612
V5613 := __e.Get(4)
_ = V5613
tmp2128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4eos, V5612)
}
__typedArg0 := symshen_4eos
__typedArg1 := V5612
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2128 {
__e.Return(V5610)
return
} else {
tmp2124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symwrite_1byte) {
return PrimWriteByte(V5612, V5611)
}
__typedArg0 := V5612
__typedArg1 := V5611
return Call(__e, PrimFunc(symwrite_1byte), __typedArg0, __typedArg1)
})()

_ = tmp2124

tmp2125 := Call(__e, PrimFunc(symshen_4string_1_6byte), V5610, V5613)


__e.TailApply(PrimFunc(symshen_4write_1chars), V5610, V5611, tmp2125, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5613)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5613
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


}


}, 4)

tmp2129 := Call(__e, ns2_1set, symshen_4write_1chars, tmp2123)


_ = tmp2129

tmp2130 := MakeNative(func(__e *ControlFlow) {
V5614 := __e.Get(1)
_ = V5614
V5615 := __e.Get(2)
_ = V5615
tmp2135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V5614)
}
__typedArg0 := V5614
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

if True == tmp2135 {
tmp2131 := Call(__e, PrimFunc(symshen_4proc_1nl), V5614)


__e.TailApply(PrimFunc(symshen_4mkstr_1l), tmp2131, V5615)
return


} else {
tmp2132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5614, Nil)
}
__typedArg0 := V5614
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4proc_1nl, tmp2132)
}
__typedArg0 := symshen_4proc_1nl
__typedArg1 := tmp2132
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4mkstr_1r), tmp2133, V5615)
return


}


}, 2)

tmp2136 := Call(__e, ns2_1set, symshen_4mkstr, tmp2130)


_ = tmp2136

tmp2137 := MakeNative(func(__e *ControlFlow) {
V5620 := __e.Get(1)
_ = V5620
V5621 := __e.Get(2)
_ = V5621
tmp2144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5621)
}
__typedArg0 := Nil
__typedArg1 := V5621
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2144 {
__e.Return(V5620)
return
} else {
tmp2142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5621)
}
__typedArg0 := V5621
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2142 {
tmp2138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5621)
}
__typedArg0 := V5621
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2139 := Call(__e, PrimFunc(symshen_4insert_1l), tmp2138, V5620)


tmp2140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5621)
}
__typedArg0 := V5621
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4mkstr_1l), tmp2139, tmp2140)
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

tmp2145 := Call(__e, ns2_1set, symshen_4mkstr_1l, tmp2137)


_ = tmp2145

tmp2146 := MakeNative(func(__e *ControlFlow) {
V5628 := __e.Get(1)
_ = V5628
V5629 := __e.Get(2)
_ = V5629
tmp2284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V5629)
}
__typedArg0 := MakeString("")
__typedArg1 := V5629
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2284 {
__e.Return(MakeString(""))
return
} else {
tmp2282 := Call(__e, PrimFunc(symshen_4_7string_2), V5629)


var ifres2269 Obj

if True == tmp2282 {
tmp2280 := Call(__e, PrimFunc(symhdstr), V5629)


tmp2281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2280)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2280
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2271 Obj

if True == tmp2281 {
tmp2279 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5629)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2273 Obj

if True == tmp2279 {
tmp2276 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5629)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("A"), tmp2276)
}
__typedArg0 := MakeString("A")
__typedArg1 := tmp2276
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2274 Obj

if True == tmp2277 {
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

if True == ifres2269 {
tmp2148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5629)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5629)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

tmp2149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4a, Nil)
}
__typedArg0 := symshen_4a
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2148, tmp2149)
}
__typedArg0 := tmp2148
__typedArg1 := tmp2149
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5628, tmp2150)
}
__typedArg0 := V5628
__typedArg1 := tmp2150
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4app, tmp2151)
}
__typedArg0 := symshen_4app
__typedArg1 := tmp2151
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp2267 := Call(__e, PrimFunc(symshen_4_7string_2), V5629)


var ifres2254 Obj

if True == tmp2267 {
tmp2265 := Call(__e, PrimFunc(symhdstr), V5629)


tmp2266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2265)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2265
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2256 Obj

if True == tmp2266 {
tmp2264 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5629)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2258 Obj

if True == tmp2264 {
tmp2261 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5629)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("R"), tmp2261)
}
__typedArg0 := MakeString("R")
__typedArg1 := tmp2261
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2259 Obj

if True == tmp2262 {
ifres2259 = True


} else {
ifres2259 = False


}

ifres2258 = ifres2259


} else {
ifres2258 = False


}

var ifres2257 Obj

if True == ifres2258 {
ifres2257 = True


} else {
ifres2257 = False


}

ifres2256 = ifres2257


} else {
ifres2256 = False


}

var ifres2255 Obj

if True == ifres2256 {
ifres2255 = True


} else {
ifres2255 = False


}

ifres2254 = ifres2255


} else {
ifres2254 = False


}

if True == ifres2254 {
tmp2153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5629)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5629)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

tmp2154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4r, Nil)
}
__typedArg0 := symshen_4r
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2153, tmp2154)
}
__typedArg0 := tmp2153
__typedArg1 := tmp2154
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5628, tmp2155)
}
__typedArg0 := V5628
__typedArg1 := tmp2155
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4app, tmp2156)
}
__typedArg0 := symshen_4app
__typedArg1 := tmp2156
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp2252 := Call(__e, PrimFunc(symshen_4_7string_2), V5629)


var ifres2239 Obj

if True == tmp2252 {
tmp2250 := Call(__e, PrimFunc(symhdstr), V5629)


tmp2251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2250)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2250
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2241 Obj

if True == tmp2251 {
tmp2249 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5629)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2243 Obj

if True == tmp2249 {
tmp2246 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5629)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("S"), tmp2246)
}
__typedArg0 := MakeString("S")
__typedArg1 := tmp2246
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2244 Obj

if True == tmp2247 {
ifres2244 = True


} else {
ifres2244 = False


}

ifres2243 = ifres2244


} else {
ifres2243 = False


}

var ifres2242 Obj

if True == ifres2243 {
ifres2242 = True


} else {
ifres2242 = False


}

ifres2241 = ifres2242


} else {
ifres2241 = False


}

var ifres2240 Obj

if True == ifres2241 {
ifres2240 = True


} else {
ifres2240 = False


}

ifres2239 = ifres2240


} else {
ifres2239 = False


}

if True == ifres2239 {
tmp2158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5629)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5629)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

tmp2159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4s, Nil)
}
__typedArg0 := symshen_4s
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2158, tmp2159)
}
__typedArg0 := tmp2158
__typedArg1 := tmp2159
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5628, tmp2160)
}
__typedArg0 := V5628
__typedArg1 := tmp2160
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4app, tmp2161)
}
__typedArg0 := symshen_4app
__typedArg1 := tmp2161
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp2237 := Call(__e, PrimFunc(symshen_4_7string_2), V5629)


if True == tmp2237 {
tmp2162 := Call(__e, PrimFunc(symhdstr), V5629)


tmp2164 := Call(__e, PrimFunc(symshen_4insert_1l), V5628, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5629)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2164, Nil)
}
__typedArg0 := tmp2164
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2162, tmp2165)
}
__typedArg0 := tmp2162
__typedArg1 := tmp2165
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp2166)
}
__typedArg0 := symcn
__typedArg1 := tmp2166
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4factor_1cn), tmp2167)
return


} else {
tmp2235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2216 Obj

if True == tmp2235 {
tmp2233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcn, tmp2233)
}
__typedArg0 := symcn
__typedArg1 := tmp2233
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2218 Obj

if True == tmp2234 {
tmp2231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2231)
}
__typedArg0 := tmp2231
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2220 Obj

if True == tmp2232 {
tmp2228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2228)
}
__typedArg0 := tmp2228
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2229)
}
__typedArg0 := tmp2229
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2222 Obj

if True == tmp2230 {
tmp2224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2224)
}
__typedArg0 := tmp2224
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2225)
}
__typedArg0 := tmp2225
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2226)
}
__typedArg0 := Nil
__typedArg1 := tmp2226
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2223 Obj

if True == tmp2227 {
ifres2223 = True


} else {
ifres2223 = False


}

ifres2222 = ifres2223


} else {
ifres2222 = False


}

var ifres2221 Obj

if True == ifres2222 {
ifres2221 = True


} else {
ifres2221 = False


}

ifres2220 = ifres2221


} else {
ifres2220 = False


}

var ifres2219 Obj

if True == ifres2220 {
ifres2219 = True


} else {
ifres2219 = False


}

ifres2218 = ifres2219


} else {
ifres2218 = False


}

var ifres2217 Obj

if True == ifres2218 {
ifres2217 = True


} else {
ifres2217 = False


}

ifres2216 = ifres2217


} else {
ifres2216 = False


}

if True == ifres2216 {
tmp2168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2168)
}
__typedArg0 := tmp2168
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2170)
}
__typedArg0 := tmp2170
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2171)
}
__typedArg0 := tmp2171
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2173 := Call(__e, PrimFunc(symshen_4insert_1l), V5628, tmp2172)


tmp2174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2173, Nil)
}
__typedArg0 := tmp2173
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2169, tmp2174)
}
__typedArg0 := tmp2169
__typedArg1 := tmp2174
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp2175)
}
__typedArg0 := symcn
__typedArg1 := tmp2175
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp2214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2188 Obj

if True == tmp2214 {
tmp2212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4app, tmp2212)
}
__typedArg0 := symshen_4app
__typedArg1 := tmp2212
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2190 Obj

if True == tmp2213 {
tmp2210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2210)
}
__typedArg0 := tmp2210
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2192 Obj

if True == tmp2211 {
tmp2207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2207)
}
__typedArg0 := tmp2207
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2208)
}
__typedArg0 := tmp2208
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2194 Obj

if True == tmp2209 {
tmp2203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2203)
}
__typedArg0 := tmp2203
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2204)
}
__typedArg0 := tmp2204
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2205)
}
__typedArg0 := tmp2205
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2196 Obj

if True == tmp2206 {
tmp2198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5629)
}
__typedArg0 := V5629
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2200)
}
__typedArg0 := tmp2200
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2201)
}
__typedArg0 := Nil
__typedArg1 := tmp2201
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2197 Obj

if True == tmp2202 {
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

var ifres2189 Obj

if True == ifres2190 {
ifres2189 = True


} else {
ifres2189 = False


}

ifres2188 = ifres2189


} else {
ifres2188 = False


}

if True == ifres2188 {
tmp2176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2176)
}
__typedArg0 := tmp2176
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5629)
}
__typedArg0 := V5629
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2179)
}
__typedArg0 := tmp2179
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2181 := Call(__e, PrimFunc(symshen_4insert_1l), V5628, tmp2180)


tmp2182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5629)
}
__typedArg0 := V5629
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2183)
}
__typedArg0 := tmp2183
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2181, tmp2184)
}
__typedArg0 := tmp2181
__typedArg1 := tmp2184
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2177, tmp2185)
}
__typedArg0 := tmp2177
__typedArg1 := tmp2185
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4app, tmp2186)
}
__typedArg0 := symshen_4app
__typedArg1 := tmp2186
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

tmp2285 := Call(__e, ns2_1set, symshen_4insert_1l, tmp2146)


_ = tmp2285

tmp2286 := MakeNative(func(__e *ControlFlow) {
V5630 := __e.Get(1)
_ = V5630
tmp2371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2302 Obj

if True == tmp2371 {
tmp2369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcn, tmp2369)
}
__typedArg0 := symcn
__typedArg1 := tmp2369
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2304 Obj

if True == tmp2370 {
tmp2367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2367)
}
__typedArg0 := tmp2367
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2306 Obj

if True == tmp2368 {
tmp2364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2364)
}
__typedArg0 := tmp2364
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2365)
}
__typedArg0 := tmp2365
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2308 Obj

if True == tmp2366 {
tmp2360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2360)
}
__typedArg0 := tmp2360
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2361)
}
__typedArg0 := tmp2361
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2362)
}
__typedArg0 := tmp2362
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2310 Obj

if True == tmp2363 {
tmp2355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2355)
}
__typedArg0 := tmp2355
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2356)
}
__typedArg0 := tmp2356
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2357)
}
__typedArg0 := tmp2357
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcn, tmp2358)
}
__typedArg0 := symcn
__typedArg1 := tmp2358
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2312 Obj

if True == tmp2359 {
tmp2350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2350)
}
__typedArg0 := tmp2350
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2351)
}
__typedArg0 := tmp2351
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2352)
}
__typedArg0 := tmp2352
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2353)
}
__typedArg0 := tmp2353
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2314 Obj

if True == tmp2354 {
tmp2344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2344)
}
__typedArg0 := tmp2344
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2345)
}
__typedArg0 := tmp2345
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2346)
}
__typedArg0 := tmp2346
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2347)
}
__typedArg0 := tmp2347
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2348)
}
__typedArg0 := tmp2348
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2316 Obj

if True == tmp2349 {
tmp2337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2337)
}
__typedArg0 := tmp2337
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2338)
}
__typedArg0 := tmp2338
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2339)
}
__typedArg0 := tmp2339
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2340)
}
__typedArg0 := tmp2340
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2341)
}
__typedArg0 := tmp2341
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2342)
}
__typedArg0 := Nil
__typedArg1 := tmp2342
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2318 Obj

if True == tmp2343 {
tmp2333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2333)
}
__typedArg0 := tmp2333
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2335)
}
__typedArg0 := Nil
__typedArg1 := tmp2335
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2320 Obj

if True == tmp2336 {
tmp2330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(tmp2331)
}
__typedArg0 := tmp2331
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres2322 Obj

if True == tmp2332 {
tmp2324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2327)
}
__typedArg0 := tmp2327
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(tmp2328)
}
__typedArg0 := tmp2328
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres2323 Obj

if True == tmp2329 {
ifres2323 = True


} else {
ifres2323 = False


}

ifres2322 = ifres2323


} else {
ifres2322 = False


}

var ifres2321 Obj

if True == ifres2322 {
ifres2321 = True


} else {
ifres2321 = False


}

ifres2320 = ifres2321


} else {
ifres2320 = False


}

var ifres2319 Obj

if True == ifres2320 {
ifres2319 = True


} else {
ifres2319 = False


}

ifres2318 = ifres2319


} else {
ifres2318 = False


}

var ifres2317 Obj

if True == ifres2318 {
ifres2317 = True


} else {
ifres2317 = False


}

ifres2316 = ifres2317


} else {
ifres2316 = False


}

var ifres2315 Obj

if True == ifres2316 {
ifres2315 = True


} else {
ifres2315 = False


}

ifres2314 = ifres2315


} else {
ifres2314 = False


}

var ifres2313 Obj

if True == ifres2314 {
ifres2313 = True


} else {
ifres2313 = False


}

ifres2312 = ifres2313


} else {
ifres2312 = False


}

var ifres2311 Obj

if True == ifres2312 {
ifres2311 = True


} else {
ifres2311 = False


}

ifres2310 = ifres2311


} else {
ifres2310 = False


}

var ifres2309 Obj

if True == ifres2310 {
ifres2309 = True


} else {
ifres2309 = False


}

ifres2308 = ifres2309


} else {
ifres2308 = False


}

var ifres2307 Obj

if True == ifres2308 {
ifres2307 = True


} else {
ifres2307 = False


}

ifres2306 = ifres2307


} else {
ifres2306 = False


}

var ifres2305 Obj

if True == ifres2306 {
ifres2305 = True


} else {
ifres2305 = False


}

ifres2304 = ifres2305


} else {
ifres2304 = False


}

var ifres2303 Obj

if True == ifres2304 {
ifres2303 = True


} else {
ifres2303 = False


}

ifres2302 = ifres2303


} else {
ifres2302 = False


}

if True == ifres2302 {
tmp2287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2289)
}
__typedArg0 := tmp2289
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2290)
}
__typedArg0 := tmp2290
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2291)
}
__typedArg0 := tmp2291
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2292)
}
__typedArg0 := tmp2292
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp2288)
__typedS1, __typedOK1 := TypedString(tmp2293)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp2288
__typedArg1 := tmp2293
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp2295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2295)
}
__typedArg0 := tmp2295
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2296)
}
__typedArg0 := tmp2296
return Call(__e, PrimFunc(symhd), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2294, tmp2299)
}
__typedArg0 := tmp2294
__typedArg1 := tmp2299
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp2300)
}
__typedArg0 := symcn
__typedArg1 := tmp2300
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V5630)
return
}


}, 1)

tmp2372 := Call(__e, ns2_1set, symshen_4factor_1cn, tmp2286)


_ = tmp2372

tmp2373 := MakeNative(func(__e *ControlFlow) {
V5633 := __e.Get(1)
_ = V5633
tmp2399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V5633)
}
__typedArg0 := MakeString("")
__typedArg1 := V5633
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2399 {
__e.Return(MakeString(""))
return
} else {
tmp2397 := Call(__e, PrimFunc(symshen_4_7string_2), V5633)


var ifres2384 Obj

if True == tmp2397 {
tmp2395 := Call(__e, PrimFunc(symhdstr), V5633)


tmp2396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2395)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2395
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2386 Obj

if True == tmp2396 {
tmp2394 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5633)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5633
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2388 Obj

if True == tmp2394 {
tmp2391 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5633)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5633
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("%"), tmp2391)
}
__typedArg0 := MakeString("%")
__typedArg1 := tmp2391
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2389 Obj

if True == tmp2392 {
ifres2389 = True


} else {
ifres2389 = False


}

ifres2388 = ifres2389


} else {
ifres2388 = False


}

var ifres2387 Obj

if True == ifres2388 {
ifres2387 = True


} else {
ifres2387 = False


}

ifres2386 = ifres2387


} else {
ifres2386 = False


}

var ifres2385 Obj

if True == ifres2386 {
ifres2385 = True


} else {
ifres2385 = False


}

ifres2384 = ifres2385


} else {
ifres2384 = False


}

if True == ifres2384 {
tmp2374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(MakeNumber(10))
}
__typedArg0 := MakeNumber(10)
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp2377 := Call(__e, PrimFunc(symshen_4proc_1nl), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5633)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5633)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5633
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp2374)
__typedS1, __typedOK1 := TypedString(tmp2377)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp2374
__typedArg1 := tmp2377
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
tmp2382 := Call(__e, PrimFunc(symshen_4_7string_2), V5633)


if True == tmp2382 {
tmp2378 := Call(__e, PrimFunc(symhdstr), V5633)


tmp2380 := Call(__e, PrimFunc(symshen_4proc_1nl), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5633)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5633
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp2378)
__typedS1, __typedOK1 := TypedString(tmp2380)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp2378
__typedArg1 := tmp2380
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

tmp2400 := Call(__e, ns2_1set, symshen_4proc_1nl, tmp2373)


_ = tmp2400

tmp2401 := MakeNative(func(__e *ControlFlow) {
V5638 := __e.Get(1)
_ = V5638
V5639 := __e.Get(2)
_ = V5639
tmp2410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5639)
}
__typedArg0 := Nil
__typedArg1 := V5639
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2410 {
__e.Return(V5638)
return
} else {
tmp2408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5639)
}
__typedArg0 := V5639
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2408 {
tmp2402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5639)
}
__typedArg0 := V5639
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5638, Nil)
}
__typedArg0 := V5638
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2402, tmp2403)
}
__typedArg0 := tmp2402
__typedArg1 := tmp2403
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4insert, tmp2404)
}
__typedArg0 := symshen_4insert
__typedArg1 := tmp2404
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5639)
}
__typedArg0 := V5639
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4mkstr_1r), tmp2405, tmp2406)
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

tmp2411 := Call(__e, ns2_1set, symshen_4mkstr_1r, tmp2401)


_ = tmp2411

tmp2412 := MakeNative(func(__e *ControlFlow) {
V5640 := __e.Get(1)
_ = V5640
V5641 := __e.Get(2)
_ = V5641
__e.TailApply(PrimFunc(symshen_4insert_1h), V5640, V5641, MakeString(""))
return
}, 2)

tmp2413 := Call(__e, ns2_1set, symshen_4insert, tmp2412)


_ = tmp2413

tmp2414 := MakeNative(func(__e *ControlFlow) {
V5650 := __e.Get(1)
_ = V5650
V5651 := __e.Get(2)
_ = V5651
V5652 := __e.Get(3)
_ = V5652
tmp2475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V5651)
}
__typedArg0 := MakeString("")
__typedArg1 := V5651
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2475 {
__e.Return(V5652)
return
} else {
tmp2473 := Call(__e, PrimFunc(symshen_4_7string_2), V5651)


var ifres2460 Obj

if True == tmp2473 {
tmp2471 := Call(__e, PrimFunc(symhdstr), V5651)


tmp2472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2471)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2471
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2462 Obj

if True == tmp2472 {
tmp2470 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5651)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5651
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2464 Obj

if True == tmp2470 {
tmp2467 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5651)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5651
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("A"), tmp2467)
}
__typedArg0 := MakeString("A")
__typedArg1 := tmp2467
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2465 Obj

if True == tmp2468 {
ifres2465 = True


} else {
ifres2465 = False


}

ifres2464 = ifres2465


} else {
ifres2464 = False


}

var ifres2463 Obj

if True == ifres2464 {
ifres2463 = True


} else {
ifres2463 = False


}

ifres2462 = ifres2463


} else {
ifres2462 = False


}

var ifres2461 Obj

if True == ifres2462 {
ifres2461 = True


} else {
ifres2461 = False


}

ifres2460 = ifres2461


} else {
ifres2460 = False


}

if True == ifres2460 {
tmp2417 := Call(__e, PrimFunc(symshen_4app), V5650, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5651)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5651)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5651
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})(), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(V5652)
__typedS1, __typedOK1 := TypedString(tmp2417)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V5652
__typedArg1 := tmp2417
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
tmp2458 := Call(__e, PrimFunc(symshen_4_7string_2), V5651)


var ifres2445 Obj

if True == tmp2458 {
tmp2456 := Call(__e, PrimFunc(symhdstr), V5651)


tmp2457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2456)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2456
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2447 Obj

if True == tmp2457 {
tmp2455 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5651)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5651
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2449 Obj

if True == tmp2455 {
tmp2452 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5651)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5651
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("R"), tmp2452)
}
__typedArg0 := MakeString("R")
__typedArg1 := tmp2452
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2450 Obj

if True == tmp2453 {
ifres2450 = True


} else {
ifres2450 = False


}

ifres2449 = ifres2450


} else {
ifres2449 = False


}

var ifres2448 Obj

if True == ifres2449 {
ifres2448 = True


} else {
ifres2448 = False


}

ifres2447 = ifres2448


} else {
ifres2447 = False


}

var ifres2446 Obj

if True == ifres2447 {
ifres2446 = True


} else {
ifres2446 = False


}

ifres2445 = ifres2446


} else {
ifres2445 = False


}

if True == ifres2445 {
tmp2420 := Call(__e, PrimFunc(symshen_4app), V5650, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5651)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5651)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5651
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})(), symshen_4r)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(V5652)
__typedS1, __typedOK1 := TypedString(tmp2420)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V5652
__typedArg1 := tmp2420
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
tmp2443 := Call(__e, PrimFunc(symshen_4_7string_2), V5651)


var ifres2430 Obj

if True == tmp2443 {
tmp2441 := Call(__e, PrimFunc(symhdstr), V5651)


tmp2442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2441)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2441
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2432 Obj

if True == tmp2442 {
tmp2440 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5651)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5651
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2434 Obj

if True == tmp2440 {
tmp2437 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5651)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5651
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("S"), tmp2437)
}
__typedArg0 := MakeString("S")
__typedArg1 := tmp2437
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2435 Obj

if True == tmp2438 {
ifres2435 = True


} else {
ifres2435 = False


}

ifres2434 = ifres2435


} else {
ifres2434 = False


}

var ifres2433 Obj

if True == ifres2434 {
ifres2433 = True


} else {
ifres2433 = False


}

ifres2432 = ifres2433


} else {
ifres2432 = False


}

var ifres2431 Obj

if True == ifres2432 {
ifres2431 = True


} else {
ifres2431 = False


}

ifres2430 = ifres2431


} else {
ifres2430 = False


}

if True == ifres2430 {
tmp2423 := Call(__e, PrimFunc(symshen_4app), V5650, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5651)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5651)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5651
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})(), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(V5652)
__typedS1, __typedOK1 := TypedString(tmp2423)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V5652
__typedArg1 := tmp2423
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
tmp2428 := Call(__e, PrimFunc(symshen_4_7string_2), V5651)


if True == tmp2428 {
tmp2424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5651)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5651
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

tmp2425 := Call(__e, PrimFunc(symhdstr), V5651)


__e.TailApply(PrimFunc(symshen_4insert_1h), V5650, tmp2424, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(V5652)
__typedS1, __typedOK1 := TypedString(tmp2425)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V5652
__typedArg1 := tmp2425
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

tmp2476 := Call(__e, ns2_1set, symshen_4insert_1h, tmp2414)


_ = tmp2476

tmp2477 := MakeNative(func(__e *ControlFlow) {
V5653 := __e.Get(1)
_ = V5653
V5654 := __e.Get(2)
_ = V5654
V5655 := __e.Get(3)
_ = V5655
tmp2478 := Call(__e, PrimFunc(symshen_4arg_1_6str), V5653, V5655)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp2478)
__typedS1, __typedOK1 := TypedString(V5654)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp2478
__typedArg1 := V5654
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 3)

tmp2479 := Call(__e, ns2_1set, symshen_4app, tmp2477)


_ = tmp2479

tmp2480 := MakeNative(func(__e *ControlFlow) {
V5659 := __e.Get(1)
_ = V5659
V5660 := __e.Get(2)
_ = V5660
tmp2488 := Call(__e, PrimFunc(symfail))


tmp2489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V5659, tmp2488)
}
__typedArg0 := V5659
__typedArg1 := tmp2488
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2489 {
__e.Return(MakeString("..."))
return
} else {
tmp2486 := Call(__e, PrimFunc(symshen_4list_2), V5659)


if True == tmp2486 {
__e.TailApply(PrimFunc(symshen_4list_1_6str), V5659, V5660)
return
} else {
tmp2484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V5659)
}
__typedArg0 := V5659
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

if True == tmp2484 {
__e.TailApply(PrimFunc(symshen_4str_1_6str), V5659, V5660)
return
} else {
tmp2482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V5659)
}
__typedArg0 := V5659
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp2482 {
__e.TailApply(PrimFunc(symshen_4vector_1_6str), V5659, V5660)
return
} else {
__e.TailApply(PrimFunc(symshen_4atom_1_6str), V5659)
return
}


}


}


}


}, 2)

tmp2490 := Call(__e, ns2_1set, symshen_4arg_1_6str, tmp2480)


_ = tmp2490

tmp2491 := MakeNative(func(__e *ControlFlow) {
V5661 := __e.Get(1)
_ = V5661
V5662 := __e.Get(2)
_ = V5662
tmp2499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4r, V5662)
}
__typedArg0 := symshen_4r
__typedArg1 := V5662
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2499 {
tmp2492 := Call(__e, PrimFunc(symshen_4maxseq))


tmp2493 := Call(__e, PrimFunc(symshen_4iter_1list), V5661, symshen_4r, tmp2492)


tmp2494 := Call(__e, PrimFunc(sym_8s), tmp2493, MakeString(")"))


__e.TailApply(PrimFunc(sym_8s), MakeString("("), tmp2494)
return


} else {
tmp2495 := Call(__e, PrimFunc(symshen_4maxseq))


tmp2496 := Call(__e, PrimFunc(symshen_4iter_1list), V5661, V5662, tmp2495)


tmp2497 := Call(__e, PrimFunc(sym_8s), tmp2496, MakeString("]"))


__e.TailApply(PrimFunc(sym_8s), MakeString("["), tmp2497)
return


}


}, 2)

tmp2500 := Call(__e, ns2_1set, symshen_4list_1_6str, tmp2491)


_ = tmp2500

tmp2501 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmaximum_1print_1sequence_1size_d)
}
__typedArg0 := sym_dmaximum_1print_1sequence_1size_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp2502 := Call(__e, ns2_1set, symshen_4maxseq, tmp2501)


_ = tmp2502

tmp2503 := MakeNative(func(__e *ControlFlow) {
V5673 := __e.Get(1)
_ = V5673
V5674 := __e.Get(2)
_ = V5674
V5675 := __e.Get(3)
_ = V5675
tmp2524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5673)
}
__typedArg0 := Nil
__typedArg1 := V5673
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2524 {
__e.Return(MakeString(""))
return
} else {
tmp2522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V5675)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V5675
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2522 {
__e.Return(MakeString("... etc"))
return
} else {
tmp2520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5673)
}
__typedArg0 := V5673
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2516 Obj

if True == tmp2520 {
tmp2518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5673)
}
__typedArg0 := V5673
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2518)
}
__typedArg0 := Nil
__typedArg1 := tmp2518
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2517 Obj

if True == tmp2519 {
ifres2517 = True


} else {
ifres2517 = False


}

ifres2516 = ifres2517


} else {
ifres2516 = False


}

if True == ifres2516 {
tmp2504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5673)
}
__typedArg0 := V5673
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4arg_1_6str), tmp2504, V5674)
return


} else {
tmp2514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5673)
}
__typedArg0 := V5673
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2514 {
tmp2505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5673)
}
__typedArg0 := V5673
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2506 := Call(__e, PrimFunc(symshen_4arg_1_6str), tmp2505, V5674)


tmp2507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5673)
}
__typedArg0 := V5673
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2509 := Call(__e, PrimFunc(symshen_4iter_1list), tmp2507, V5674, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V5675)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V5675
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


tmp2510 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp2509)


__e.TailApply(PrimFunc(sym_8s), tmp2506, tmp2510)
return


} else {
tmp2511 := Call(__e, PrimFunc(symshen_4arg_1_6str), V5673, V5674)


tmp2512 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp2511)


__e.TailApply(PrimFunc(sym_8s), MakeString("|"), tmp2512)
return


}


}


}


}


}, 3)

tmp2525 := Call(__e, ns2_1set, symshen_4iter_1list, tmp2503)


_ = tmp2525

tmp2526 := MakeNative(func(__e *ControlFlow) {
V5678 := __e.Get(1)
_ = V5678
V5679 := __e.Get(2)
_ = V5679
tmp2531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4a, V5679)
}
__typedArg0 := symshen_4a
__typedArg1 := V5679
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2531 {
__e.Return(V5678)
return
} else {
tmp2527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(MakeNumber(34))
}
__typedArg0 := MakeNumber(34)
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp2528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(MakeNumber(34))
}
__typedArg0 := MakeNumber(34)
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp2529 := Call(__e, PrimFunc(sym_8s), V5678, tmp2528)


__e.TailApply(PrimFunc(sym_8s), tmp2527, tmp2529)
return


}


}, 2)

tmp2532 := Call(__e, ns2_1set, symshen_4str_1_6str, tmp2526)


_ = tmp2532

tmp2533 := MakeNative(func(__e *ControlFlow) {
V5680 := __e.Get(1)
_ = V5680
V5681 := __e.Get(2)
_ = V5681
tmp2546 := Call(__e, PrimFunc(symshen_4print_1vector_2), V5680)


if True == tmp2546 {
tmp2534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V5680, MakeNumber(0))
}
__typedArg0 := V5680
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp2535 := Call(__e, PrimFunc(symfn), tmp2534)


__e.TailApply(tmp2535, V5680)
return


} else {
tmp2544 := Call(__e, PrimFunc(symvector_2), V5680)


if True == tmp2544 {
tmp2536 := Call(__e, PrimFunc(symshen_4maxseq))


tmp2537 := Call(__e, PrimFunc(symshen_4iter_1vector), V5680, MakeNumber(1), V5681, tmp2536)


tmp2538 := Call(__e, PrimFunc(sym_8s), tmp2537, MakeString(">"))


__e.TailApply(PrimFunc(sym_8s), MakeString("<"), tmp2538)
return


} else {
tmp2539 := Call(__e, PrimFunc(symshen_4maxseq))


tmp2540 := Call(__e, PrimFunc(symshen_4iter_1vector), V5680, MakeNumber(0), V5681, tmp2539)


tmp2541 := Call(__e, PrimFunc(sym_8s), tmp2540, MakeString(">>"))


tmp2542 := Call(__e, PrimFunc(sym_8s), MakeString("<"), tmp2541)


__e.TailApply(PrimFunc(sym_8s), MakeString("<"), tmp2542)
return


}


}


}, 2)

tmp2547 := Call(__e, ns2_1set, symshen_4vector_1_6str, tmp2533)


_ = tmp2547

tmp2548 := MakeNative(func(__e *ControlFlow) {
V5682 := __e.Get(1)
_ = V5682
tmp2549 := MakeNative(func(__e *ControlFlow) {
W5683 := __e.Get(1)
_ = W5683
tmp2556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5683, symshen_4tuple)
}
__typedArg0 := W5683
__typedArg1 := symshen_4tuple
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2556 {
__e.Return(True)
return
} else {
tmp2554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5683, symshen_4pvar)
}
__typedArg0 := W5683
__typedArg1 := symshen_4pvar
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2554 {
__e.Return(True)
return
} else {
tmp2551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(W5683)
}
__typedArg0 := W5683
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp2551)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp2551
return Call(__e, PrimFunc(symnot), __typedArg0)
})() {
__e.TailApply(PrimFunc(symshen_4fbound_2), W5683)
return
} else {
__e.Return(False)
return
}


}


}


}, 1)

tmp2557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V5682, MakeNumber(0))
}
__typedArg0 := V5682
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp2549, tmp2557)
return


}, 1)

tmp2558 := Call(__e, ns2_1set, symshen_4print_1vector_2, tmp2548)


_ = tmp2558

tmp2559 := MakeNative(func(__e *ControlFlow) {
V5684 := __e.Get(1)
_ = V5684
tmp2560 := Call(__e, PrimFunc(symarity), V5684)


tmp2561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp2560, MakeNumber(-1))
}
__typedArg0 := tmp2560
__typedArg1 := MakeNumber(-1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp2561)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp2561
return Call(__e, PrimFunc(symnot), __typedArg0)
})())
return


}, 1)

tmp2562 := Call(__e, ns2_1set, symshen_4fbound_2, tmp2559)


_ = tmp2562

tmp2563 := MakeNative(func(__e *ControlFlow) {
V5685 := __e.Get(1)
_ = V5685
tmp2564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V5685, MakeNumber(1))
}
__typedArg0 := V5685
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp2565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V5685, MakeNumber(2))
}
__typedArg0 := V5685
__typedArg1 := MakeNumber(2)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp2566 := Call(__e, PrimFunc(symshen_4app), tmp2565, MakeString(")"), symshen_4s)


tmp2568 := Call(__e, PrimFunc(symshen_4app), tmp2564, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" "))
__typedS1, __typedOK1 := TypedString(tmp2566)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" ")
__typedArg1 := tmp2566
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("(@p "))
__typedS1, __typedOK1 := TypedString(tmp2568)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("(@p ")
__typedArg1 := tmp2568
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp2569 := Call(__e, ns2_1set, symshen_4tuple, tmp2563)


_ = tmp2569

tmp2570 := MakeNative(func(__e *ControlFlow) {
V5692 := __e.Get(1)
_ = V5692
V5693 := __e.Get(2)
_ = V5693
V5694 := __e.Get(3)
_ = V5694
V5695 := __e.Get(4)
_ = V5695
tmp2590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V5695)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V5695
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2590 {
__e.Return(MakeString("... etc"))
return
} else {
tmp2571 := MakeNative(func(__e *ControlFlow) {
W5696 := __e.Get(1)
_ = W5696
tmp2572 := MakeNative(func(__e *ControlFlow) {
W5698 := __e.Get(1)
_ = W5698
tmp2581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5696, symshen_4out_1of_1bounds)
}
__typedArg0 := W5696
__typedArg1 := symshen_4out_1of_1bounds
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2581 {
__e.Return(MakeString(""))
return
} else {
tmp2579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W5698, symshen_4out_1of_1bounds)
}
__typedArg0 := W5698
__typedArg1 := symshen_4out_1of_1bounds
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2579 {
__e.TailApply(PrimFunc(symshen_4arg_1_6str), W5696, V5694)
return
} else {
tmp2573 := Call(__e, PrimFunc(symshen_4arg_1_6str), W5696, V5694)


tmp2574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5693)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5693
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()

tmp2576 := Call(__e, PrimFunc(symshen_4iter_1vector), V5692, tmp2574, V5694, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V5695)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V5695
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


tmp2577 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp2576)


__e.TailApply(PrimFunc(sym_8s), tmp2573, tmp2577)
return


}


}


}, 1)

tmp2582 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V5692, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5693)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5693
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := V5692
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V5693)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V5693
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp2584 := MakeNative(func(__e *ControlFlow) {
Z5699 := __e.Get(1)
_ = Z5699
__e.Return(symshen_4out_1of_1bounds)
return
}, 1)

tmp2585 := Call(__e, try_1catch, tmp2582, tmp2584)


__e.TailApply(tmp2572, tmp2585)
return


}, 1)

tmp2586 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V5692, V5693)
}
__typedArg0 := V5692
__typedArg1 := V5693
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 0)

tmp2587 := MakeNative(func(__e *ControlFlow) {
Z5697 := __e.Get(1)
_ = Z5697
__e.Return(symshen_4out_1of_1bounds)
return
}, 1)

tmp2588 := Call(__e, try_1catch, tmp2586, tmp2587)


__e.TailApply(tmp2571, tmp2588)
return


}


}, 4)

tmp2591 := Call(__e, ns2_1set, symshen_4iter_1vector, tmp2570)


_ = tmp2591

tmp2592 := MakeNative(func(__e *ControlFlow) {
V5700 := __e.Get(1)
_ = V5700
tmp2593 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V5700)
}
__typedArg0 := V5700
return Call(__e, PrimFunc(symstr), __typedArg0)
})())
return
}, 0)

tmp2594 := MakeNative(func(__e *ControlFlow) {
Z5701 := __e.Get(1)
_ = Z5701
__e.TailApply(PrimFunc(symshen_4funexstring))
return
}, 1)

__e.TailApply(try_1catch, tmp2593, tmp2594)
return


}, 1)

tmp2595 := Call(__e, ns2_1set, symshen_4atom_1_6str, tmp2592)


_ = tmp2595

tmp2596 := MakeNative(func(__e *ControlFlow) {
tmp2597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString("x"))
}
__typedArg0 := MakeString("x")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp2598 := Call(__e, PrimFunc(symgensym), tmp2597)


tmp2599 := Call(__e, PrimFunc(symshen_4arg_1_6str), tmp2598, symshen_4a)


tmp2600 := Call(__e, PrimFunc(sym_8s), tmp2599, MakeString("\x11"))


tmp2601 := Call(__e, PrimFunc(sym_8s), MakeString("e"), tmp2600)


tmp2602 := Call(__e, PrimFunc(sym_8s), MakeString("n"), tmp2601)


tmp2603 := Call(__e, PrimFunc(sym_8s), MakeString("u"), tmp2602)


tmp2604 := Call(__e, PrimFunc(sym_8s), MakeString("f"), tmp2603)


__e.TailApply(PrimFunc(sym_8s), MakeString("\x10"), tmp2604)
return


}, 0)

tmp2605 := Call(__e, ns2_1set, symshen_4funexstring, tmp2596)


_ = tmp2605

tmp2606 := MakeNative(func(__e *ControlFlow) {
V5702 := __e.Get(1)
_ = V5702
tmp2610 := Call(__e, PrimFunc(symempty_2), V5702)


if True == tmp2610 {
__e.Return(True)
return
} else {
tmp2608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5702)
}
__typedArg0 := V5702
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2608 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

__e.TailApply(ns2_1set, symshen_4list_2, tmp2606)
return




}, 0)

