package main

import . "github.com/pyrex41/shen-go/kl"

var WriterMain = MakeNative(func(__e *ControlFlow) {
tmp2108 := MakeNative(func(__e *ControlFlow) {
V5602 := __e.Get(1)
_ = V5602
tmp2109 := Call(__e, PrimFunc(symshen_4insert), V5602, MakeString("~S"))


W56032103 := tmp2109
_ = W56032103

tmp2110 := Call(__e, PrimFunc(symstoutput))


tmp2111 := Call(__e, PrimFunc(sympr), W56032103, tmp2110)


W56042104 := tmp2111
_ = W56042104

__e.Return(V5602)
return


}, 1)

tmp2112 := Call(__e, ns2_1set, symprint, tmp2108)


_ = tmp2112

tmp2113 := MakeNative(func(__e *ControlFlow) {
V5605 := __e.Get(1)
_ = V5605
V5606 := __e.Get(2)
_ = V5606
tmp2118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dhush_d)
}
__typedArg0 := sym_dhush_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

if True == tmp2118 {
__e.Return(V5605)
return
} else {
tmp2116 := Call(__e, PrimFunc(symshen_4char_1stoutput_2), V5606)


if True == tmp2116 {
__e.TailApply(PrimFunc(symshen_4write_1string), V5605, V5606)
return
} else {
tmp2114 := Call(__e, PrimFunc(symshen_4string_1_6byte), V5605, MakeNumber(0))


__e.TailApply(PrimFunc(symshen_4write_1chars), V5605, V5606, tmp2114, MakeNumber(1))
return


}


}


}, 2)

tmp2119 := Call(__e, ns2_1set, sympr, tmp2113)


_ = tmp2119

tmp2120 := MakeNative(func(__e *ControlFlow) {
V5607 := __e.Get(1)
_ = V5607
V5608 := __e.Get(2)
_ = V5608
tmp2121 := MakeNative(func(__e *ControlFlow) {
tmp2122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sympos) {
return PrimPos(V5607, V5608)
}
__typedArg0 := V5607
__typedArg1 := V5608
return Call(__e, PrimFunc(sympos), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp2122)
}
__typedArg0 := tmp2122
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})())
return


}, 0)

tmp2123 := MakeNative(func(__e *ControlFlow) {
Z5609 := __e.Get(1)
_ = Z5609
__e.Return(symshen_4eos)
return
}, 1)

__e.TailApply(try_1catch, tmp2121, tmp2123)
return


}, 2)

tmp2124 := Call(__e, ns2_1set, symshen_4string_1_6byte, tmp2120)


_ = tmp2124

tmp2125 := MakeNative(func(__e *ControlFlow) {
V5610 := __e.Get(1)
_ = V5610
V5611 := __e.Get(2)
_ = V5611
V5612 := __e.Get(3)
_ = V5612
V5613 := __e.Get(4)
_ = V5613
tmp2130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4eos, V5612)
}
__typedArg0 := symshen_4eos
__typedArg1 := V5612
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2130 {
__e.Return(V5610)
return
} else {
tmp2126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symwrite_1byte) {
return PrimWriteByte(V5612, V5611)
}
__typedArg0 := V5612
__typedArg1 := V5611
return Call(__e, PrimFunc(symwrite_1byte), __typedArg0, __typedArg1)
})()

_ = tmp2126

tmp2127 := Call(__e, PrimFunc(symshen_4string_1_6byte), V5610, V5613)


__e.TailApply(PrimFunc(symshen_4write_1chars), V5610, V5611, tmp2127, (func() Obj {
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

tmp2131 := Call(__e, ns2_1set, symshen_4write_1chars, tmp2125)


_ = tmp2131

tmp2132 := MakeNative(func(__e *ControlFlow) {
V5614 := __e.Get(1)
_ = V5614
V5615 := __e.Get(2)
_ = V5615
tmp2137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V5614)
}
__typedArg0 := V5614
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

if True == tmp2137 {
tmp2133 := Call(__e, PrimFunc(symshen_4proc_1nl), V5614)


__e.TailApply(PrimFunc(symshen_4mkstr_1l), tmp2133, V5615)
return


} else {
tmp2134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5614, Nil)
}
__typedArg0 := V5614
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4proc_1nl, tmp2134)
}
__typedArg0 := symshen_4proc_1nl
__typedArg1 := tmp2134
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4mkstr_1r), tmp2135, V5615)
return


}


}, 2)

tmp2138 := Call(__e, ns2_1set, symshen_4mkstr, tmp2132)


_ = tmp2138

tmp2139 := MakeNative(func(__e *ControlFlow) {
V5620 := __e.Get(1)
_ = V5620
V5621 := __e.Get(2)
_ = V5621
tmp2146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5621)
}
__typedArg0 := Nil
__typedArg1 := V5621
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2146 {
__e.Return(V5620)
return
} else {
tmp2144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5621)
}
__typedArg0 := V5621
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2144 {
tmp2140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5621)
}
__typedArg0 := V5621
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2141 := Call(__e, PrimFunc(symshen_4insert_1l), tmp2140, V5620)


tmp2142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5621)
}
__typedArg0 := V5621
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4mkstr_1l), tmp2141, tmp2142)
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

tmp2147 := Call(__e, ns2_1set, symshen_4mkstr_1l, tmp2139)


_ = tmp2147

tmp2148 := MakeNative(func(__e *ControlFlow) {
V5628 := __e.Get(1)
_ = V5628
V5629 := __e.Get(2)
_ = V5629
tmp2286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V5629)
}
__typedArg0 := MakeString("")
__typedArg1 := V5629
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2286 {
__e.Return(MakeString(""))
return
} else {
tmp2284 := Call(__e, PrimFunc(symshen_4_7string_2), V5629)


var ifres2271 Obj

if True == tmp2284 {
tmp2282 := Call(__e, PrimFunc(symhdstr), V5629)


tmp2283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2282)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2282
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2273 Obj

if True == tmp2283 {
tmp2281 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5629)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2275 Obj

if True == tmp2281 {
tmp2278 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5629)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("A"), tmp2278)
}
__typedArg0 := MakeString("A")
__typedArg1 := tmp2278
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2276 Obj

if True == tmp2279 {
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

if True == ifres2271 {
tmp2150 := (func() Obj {
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

tmp2151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4a, Nil)
}
__typedArg0 := symshen_4a
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2150, tmp2151)
}
__typedArg0 := tmp2150
__typedArg1 := tmp2151
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5628, tmp2152)
}
__typedArg0 := V5628
__typedArg1 := tmp2152
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4app, tmp2153)
}
__typedArg0 := symshen_4app
__typedArg1 := tmp2153
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp2269 := Call(__e, PrimFunc(symshen_4_7string_2), V5629)


var ifres2256 Obj

if True == tmp2269 {
tmp2267 := Call(__e, PrimFunc(symhdstr), V5629)


tmp2268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2267)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2267
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2258 Obj

if True == tmp2268 {
tmp2266 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5629)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2260 Obj

if True == tmp2266 {
tmp2263 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5629)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("R"), tmp2263)
}
__typedArg0 := MakeString("R")
__typedArg1 := tmp2263
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2261 Obj

if True == tmp2264 {
ifres2261 = True


} else {
ifres2261 = False


}

ifres2260 = ifres2261


} else {
ifres2260 = False


}

var ifres2259 Obj

if True == ifres2260 {
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

if True == ifres2256 {
tmp2155 := (func() Obj {
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

tmp2156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4r, Nil)
}
__typedArg0 := symshen_4r
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2155, tmp2156)
}
__typedArg0 := tmp2155
__typedArg1 := tmp2156
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5628, tmp2157)
}
__typedArg0 := V5628
__typedArg1 := tmp2157
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4app, tmp2158)
}
__typedArg0 := symshen_4app
__typedArg1 := tmp2158
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp2254 := Call(__e, PrimFunc(symshen_4_7string_2), V5629)


var ifres2241 Obj

if True == tmp2254 {
tmp2252 := Call(__e, PrimFunc(symhdstr), V5629)


tmp2253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2252)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2252
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2243 Obj

if True == tmp2253 {
tmp2251 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5629)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2245 Obj

if True == tmp2251 {
tmp2248 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5629)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("S"), tmp2248)
}
__typedArg0 := MakeString("S")
__typedArg1 := tmp2248
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2246 Obj

if True == tmp2249 {
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

if True == ifres2241 {
tmp2160 := (func() Obj {
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

tmp2161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4s, Nil)
}
__typedArg0 := symshen_4s
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2160, tmp2161)
}
__typedArg0 := tmp2160
__typedArg1 := tmp2161
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5628, tmp2162)
}
__typedArg0 := V5628
__typedArg1 := tmp2162
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4app, tmp2163)
}
__typedArg0 := symshen_4app
__typedArg1 := tmp2163
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp2239 := Call(__e, PrimFunc(symshen_4_7string_2), V5629)


if True == tmp2239 {
tmp2164 := Call(__e, PrimFunc(symhdstr), V5629)


tmp2166 := Call(__e, PrimFunc(symshen_4insert_1l), V5628, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5629)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2166, Nil)
}
__typedArg0 := tmp2166
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2164, tmp2167)
}
__typedArg0 := tmp2164
__typedArg1 := tmp2167
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp2168)
}
__typedArg0 := symcn
__typedArg1 := tmp2168
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4factor_1cn), tmp2169)
return


} else {
tmp2237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2218 Obj

if True == tmp2237 {
tmp2235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcn, tmp2235)
}
__typedArg0 := symcn
__typedArg1 := tmp2235
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2220 Obj

if True == tmp2236 {
tmp2233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2233)
}
__typedArg0 := tmp2233
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2222 Obj

if True == tmp2234 {
tmp2230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2230)
}
__typedArg0 := tmp2230
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2231)
}
__typedArg0 := tmp2231
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2224 Obj

if True == tmp2232 {
tmp2226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2226)
}
__typedArg0 := tmp2226
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2227)
}
__typedArg0 := tmp2227
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2228)
}
__typedArg0 := Nil
__typedArg1 := tmp2228
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2225 Obj

if True == tmp2229 {
ifres2225 = True


} else {
ifres2225 = False


}

ifres2224 = ifres2225


} else {
ifres2224 = False


}

var ifres2223 Obj

if True == ifres2224 {
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

if True == ifres2218 {
tmp2170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2170)
}
__typedArg0 := tmp2170
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5629)
}
__typedArg0 := V5629
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2173)
}
__typedArg0 := tmp2173
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2175 := Call(__e, PrimFunc(symshen_4insert_1l), V5628, tmp2174)


tmp2176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2175, Nil)
}
__typedArg0 := tmp2175
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2171, tmp2176)
}
__typedArg0 := tmp2171
__typedArg1 := tmp2176
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp2177)
}
__typedArg0 := symcn
__typedArg1 := tmp2177
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp2216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2190 Obj

if True == tmp2216 {
tmp2214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4app, tmp2214)
}
__typedArg0 := symshen_4app
__typedArg1 := tmp2214
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2192 Obj

if True == tmp2215 {
tmp2212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2212)
}
__typedArg0 := tmp2212
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2194 Obj

if True == tmp2213 {
tmp2209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2209)
}
__typedArg0 := tmp2209
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2210)
}
__typedArg0 := tmp2210
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2196 Obj

if True == tmp2211 {
tmp2205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2205)
}
__typedArg0 := tmp2205
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2206)
}
__typedArg0 := tmp2206
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2207)
}
__typedArg0 := tmp2207
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2198 Obj

if True == tmp2208 {
tmp2200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5629)
}
__typedArg0 := V5629
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2201)
}
__typedArg0 := tmp2201
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2203)
}
__typedArg0 := Nil
__typedArg1 := tmp2203
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2199 Obj

if True == tmp2204 {
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
tmp2178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2178)
}
__typedArg0 := tmp2178
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2180)
}
__typedArg0 := tmp2180
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2181)
}
__typedArg0 := tmp2181
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2183 := Call(__e, PrimFunc(symshen_4insert_1l), V5628, tmp2182)


tmp2184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5629)
}
__typedArg0 := V5629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2184)
}
__typedArg0 := tmp2184
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2183, tmp2186)
}
__typedArg0 := tmp2183
__typedArg1 := tmp2186
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2179, tmp2187)
}
__typedArg0 := tmp2179
__typedArg1 := tmp2187
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4app, tmp2188)
}
__typedArg0 := symshen_4app
__typedArg1 := tmp2188
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

tmp2287 := Call(__e, ns2_1set, symshen_4insert_1l, tmp2148)


_ = tmp2287

tmp2288 := MakeNative(func(__e *ControlFlow) {
V5630 := __e.Get(1)
_ = V5630
tmp2373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2304 Obj

if True == tmp2373 {
tmp2371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcn, tmp2371)
}
__typedArg0 := symcn
__typedArg1 := tmp2371
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2306 Obj

if True == tmp2372 {
tmp2369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2369)
}
__typedArg0 := tmp2369
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2308 Obj

if True == tmp2370 {
tmp2366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2366)
}
__typedArg0 := tmp2366
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2367)
}
__typedArg0 := tmp2367
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2310 Obj

if True == tmp2368 {
tmp2362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2362)
}
__typedArg0 := tmp2362
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2363)
}
__typedArg0 := tmp2363
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2364)
}
__typedArg0 := tmp2364
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2312 Obj

if True == tmp2365 {
tmp2357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2357)
}
__typedArg0 := tmp2357
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2358)
}
__typedArg0 := tmp2358
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2359)
}
__typedArg0 := tmp2359
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcn, tmp2360)
}
__typedArg0 := symcn
__typedArg1 := tmp2360
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2314 Obj

if True == tmp2361 {
tmp2352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2352)
}
__typedArg0 := tmp2352
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2353)
}
__typedArg0 := tmp2353
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2354)
}
__typedArg0 := tmp2354
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2355)
}
__typedArg0 := tmp2355
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2316 Obj

if True == tmp2356 {
tmp2346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2346)
}
__typedArg0 := tmp2346
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2347)
}
__typedArg0 := tmp2347
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2348)
}
__typedArg0 := tmp2348
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2349)
}
__typedArg0 := tmp2349
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2350)
}
__typedArg0 := tmp2350
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2318 Obj

if True == tmp2351 {
tmp2339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2339)
}
__typedArg0 := tmp2339
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2340)
}
__typedArg0 := tmp2340
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2341)
}
__typedArg0 := tmp2341
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2342)
}
__typedArg0 := tmp2342
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2343)
}
__typedArg0 := tmp2343
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2344)
}
__typedArg0 := Nil
__typedArg1 := tmp2344
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2320 Obj

if True == tmp2345 {
tmp2335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2335)
}
__typedArg0 := tmp2335
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2336)
}
__typedArg0 := tmp2336
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2337)
}
__typedArg0 := Nil
__typedArg1 := tmp2337
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2322 Obj

if True == tmp2338 {
tmp2332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2332)
}
__typedArg0 := tmp2332
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(tmp2333)
}
__typedArg0 := tmp2333
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres2324 Obj

if True == tmp2334 {
tmp2326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symtl), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2328)
}
__typedArg0 := tmp2328
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2329)
}
__typedArg0 := tmp2329
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(tmp2330)
}
__typedArg0 := tmp2330
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres2325 Obj

if True == tmp2331 {
ifres2325 = True


} else {
ifres2325 = False


}

ifres2324 = ifres2325


} else {
ifres2324 = False


}

var ifres2323 Obj

if True == ifres2324 {
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

if True == ifres2304 {
tmp2289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2289)
}
__typedArg0 := tmp2289
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2292)
}
__typedArg0 := tmp2292
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2293)
}
__typedArg0 := tmp2293
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2294)
}
__typedArg0 := tmp2294
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp2290)
__typedS1, __typedOK1 := TypedString(tmp2295)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp2290
__typedArg1 := tmp2295
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp2297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5630)
}
__typedArg0 := V5630
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2298)
}
__typedArg0 := tmp2298
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2299)
}
__typedArg0 := tmp2299
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2300)
}
__typedArg0 := tmp2300
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2296, tmp2301)
}
__typedArg0 := tmp2296
__typedArg1 := tmp2301
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcn, tmp2302)
}
__typedArg0 := symcn
__typedArg1 := tmp2302
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V5630)
return
}


}, 1)

tmp2374 := Call(__e, ns2_1set, symshen_4factor_1cn, tmp2288)


_ = tmp2374

tmp2375 := MakeNative(func(__e *ControlFlow) {
V5633 := __e.Get(1)
_ = V5633
tmp2401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V5633)
}
__typedArg0 := MakeString("")
__typedArg1 := V5633
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2401 {
__e.Return(MakeString(""))
return
} else {
tmp2399 := Call(__e, PrimFunc(symshen_4_7string_2), V5633)


var ifres2386 Obj

if True == tmp2399 {
tmp2397 := Call(__e, PrimFunc(symhdstr), V5633)


tmp2398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2397)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2397
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2388 Obj

if True == tmp2398 {
tmp2396 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5633)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5633
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2390 Obj

if True == tmp2396 {
tmp2393 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5633)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5633
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("%"), tmp2393)
}
__typedArg0 := MakeString("%")
__typedArg1 := tmp2393
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2391 Obj

if True == tmp2394 {
ifres2391 = True


} else {
ifres2391 = False


}

ifres2390 = ifres2391


} else {
ifres2390 = False


}

var ifres2389 Obj

if True == ifres2390 {
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

if True == ifres2386 {
tmp2376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(MakeNumber(10))
}
__typedArg0 := MakeNumber(10)
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp2379 := Call(__e, PrimFunc(symshen_4proc_1nl), (func() Obj {
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
__typedS0, __typedOK0 := TypedString(tmp2376)
__typedS1, __typedOK1 := TypedString(tmp2379)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp2376
__typedArg1 := tmp2379
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
tmp2384 := Call(__e, PrimFunc(symshen_4_7string_2), V5633)


if True == tmp2384 {
tmp2380 := Call(__e, PrimFunc(symhdstr), V5633)


tmp2382 := Call(__e, PrimFunc(symshen_4proc_1nl), (func() Obj {
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
__typedS0, __typedOK0 := TypedString(tmp2380)
__typedS1, __typedOK1 := TypedString(tmp2382)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp2380
__typedArg1 := tmp2382
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

tmp2402 := Call(__e, ns2_1set, symshen_4proc_1nl, tmp2375)


_ = tmp2402

tmp2403 := MakeNative(func(__e *ControlFlow) {
V5638 := __e.Get(1)
_ = V5638
V5639 := __e.Get(2)
_ = V5639
tmp2412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5639)
}
__typedArg0 := Nil
__typedArg1 := V5639
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2412 {
__e.Return(V5638)
return
} else {
tmp2410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5639)
}
__typedArg0 := V5639
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2410 {
tmp2404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5639)
}
__typedArg0 := V5639
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5638, Nil)
}
__typedArg0 := V5638
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2404, tmp2405)
}
__typedArg0 := tmp2404
__typedArg1 := tmp2405
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4insert, tmp2406)
}
__typedArg0 := symshen_4insert
__typedArg1 := tmp2406
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5639)
}
__typedArg0 := V5639
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4mkstr_1r), tmp2407, tmp2408)
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

tmp2413 := Call(__e, ns2_1set, symshen_4mkstr_1r, tmp2403)


_ = tmp2413

tmp2414 := MakeNative(func(__e *ControlFlow) {
V5640 := __e.Get(1)
_ = V5640
V5641 := __e.Get(2)
_ = V5641
__e.TailApply(PrimFunc(symshen_4insert_1h), V5640, V5641, MakeString(""))
return
}, 2)

tmp2415 := Call(__e, ns2_1set, symshen_4insert, tmp2414)


_ = tmp2415

tmp2416 := MakeNative(func(__e *ControlFlow) {
V5650 := __e.Get(1)
_ = V5650
V5651 := __e.Get(2)
_ = V5651
V5652 := __e.Get(3)
_ = V5652
tmp2477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V5651)
}
__typedArg0 := MakeString("")
__typedArg1 := V5651
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2477 {
__e.Return(V5652)
return
} else {
tmp2475 := Call(__e, PrimFunc(symshen_4_7string_2), V5651)


var ifres2462 Obj

if True == tmp2475 {
tmp2473 := Call(__e, PrimFunc(symhdstr), V5651)


tmp2474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2473)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2473
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2464 Obj

if True == tmp2474 {
tmp2472 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5651)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5651
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2466 Obj

if True == tmp2472 {
tmp2469 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5651)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5651
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("A"), tmp2469)
}
__typedArg0 := MakeString("A")
__typedArg1 := tmp2469
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2467 Obj

if True == tmp2470 {
ifres2467 = True


} else {
ifres2467 = False


}

ifres2466 = ifres2467


} else {
ifres2466 = False


}

var ifres2465 Obj

if True == ifres2466 {
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

if True == ifres2462 {
tmp2419 := Call(__e, PrimFunc(symshen_4app), V5650, (func() Obj {
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
__typedS1, __typedOK1 := TypedString(tmp2419)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V5652
__typedArg1 := tmp2419
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
tmp2460 := Call(__e, PrimFunc(symshen_4_7string_2), V5651)


var ifres2447 Obj

if True == tmp2460 {
tmp2458 := Call(__e, PrimFunc(symhdstr), V5651)


tmp2459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2458)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2458
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2449 Obj

if True == tmp2459 {
tmp2457 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5651)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5651
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2451 Obj

if True == tmp2457 {
tmp2454 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5651)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5651
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("R"), tmp2454)
}
__typedArg0 := MakeString("R")
__typedArg1 := tmp2454
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2452 Obj

if True == tmp2455 {
ifres2452 = True


} else {
ifres2452 = False


}

ifres2451 = ifres2452


} else {
ifres2451 = False


}

var ifres2450 Obj

if True == ifres2451 {
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

if True == ifres2447 {
tmp2422 := Call(__e, PrimFunc(symshen_4app), V5650, (func() Obj {
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
__typedS1, __typedOK1 := TypedString(tmp2422)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V5652
__typedArg1 := tmp2422
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
tmp2445 := Call(__e, PrimFunc(symshen_4_7string_2), V5651)


var ifres2432 Obj

if True == tmp2445 {
tmp2443 := Call(__e, PrimFunc(symhdstr), V5651)


tmp2444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("~"), tmp2443)
}
__typedArg0 := MakeString("~")
__typedArg1 := tmp2443
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2434 Obj

if True == tmp2444 {
tmp2442 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5651)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5651
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres2436 Obj

if True == tmp2442 {
tmp2439 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5651)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5651
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp2440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("S"), tmp2439)
}
__typedArg0 := MakeString("S")
__typedArg1 := tmp2439
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2437 Obj

if True == tmp2440 {
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

if True == ifres2432 {
tmp2425 := Call(__e, PrimFunc(symshen_4app), V5650, (func() Obj {
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
tmp2430 := Call(__e, PrimFunc(symshen_4_7string_2), V5651)


if True == tmp2430 {
tmp2426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V5651)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V5651
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

tmp2427 := Call(__e, PrimFunc(symhdstr), V5651)


__e.TailApply(PrimFunc(symshen_4insert_1h), V5650, tmp2426, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(V5652)
__typedS1, __typedOK1 := TypedString(tmp2427)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V5652
__typedArg1 := tmp2427
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

tmp2478 := Call(__e, ns2_1set, symshen_4insert_1h, tmp2416)


_ = tmp2478

tmp2479 := MakeNative(func(__e *ControlFlow) {
V5653 := __e.Get(1)
_ = V5653
V5654 := __e.Get(2)
_ = V5654
V5655 := __e.Get(3)
_ = V5655
tmp2480 := Call(__e, PrimFunc(symshen_4arg_1_6str), V5653, V5655)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp2480)
__typedS1, __typedOK1 := TypedString(V5654)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp2480
__typedArg1 := V5654
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 3)

tmp2481 := Call(__e, ns2_1set, symshen_4app, tmp2479)


_ = tmp2481

tmp2482 := MakeNative(func(__e *ControlFlow) {
V5659 := __e.Get(1)
_ = V5659
V5660 := __e.Get(2)
_ = V5660
tmp2490 := Call(__e, PrimFunc(symfail))


tmp2491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V5659, tmp2490)
}
__typedArg0 := V5659
__typedArg1 := tmp2490
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2491 {
__e.Return(MakeString("..."))
return
} else {
tmp2488 := Call(__e, PrimFunc(symshen_4list_2), V5659)


if True == tmp2488 {
__e.TailApply(PrimFunc(symshen_4list_1_6str), V5659, V5660)
return
} else {
tmp2486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V5659)
}
__typedArg0 := V5659
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

if True == tmp2486 {
__e.TailApply(PrimFunc(symshen_4str_1_6str), V5659, V5660)
return
} else {
tmp2484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V5659)
}
__typedArg0 := V5659
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp2484 {
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

tmp2492 := Call(__e, ns2_1set, symshen_4arg_1_6str, tmp2482)


_ = tmp2492

tmp2493 := MakeNative(func(__e *ControlFlow) {
V5661 := __e.Get(1)
_ = V5661
V5662 := __e.Get(2)
_ = V5662
tmp2501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4r, V5662)
}
__typedArg0 := symshen_4r
__typedArg1 := V5662
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2501 {
tmp2494 := Call(__e, PrimFunc(symshen_4maxseq))


tmp2495 := Call(__e, PrimFunc(symshen_4iter_1list), V5661, symshen_4r, tmp2494)


tmp2496 := Call(__e, PrimFunc(sym_8s), tmp2495, MakeString(")"))


__e.TailApply(PrimFunc(sym_8s), MakeString("("), tmp2496)
return


} else {
tmp2497 := Call(__e, PrimFunc(symshen_4maxseq))


tmp2498 := Call(__e, PrimFunc(symshen_4iter_1list), V5661, V5662, tmp2497)


tmp2499 := Call(__e, PrimFunc(sym_8s), tmp2498, MakeString("]"))


__e.TailApply(PrimFunc(sym_8s), MakeString("["), tmp2499)
return


}


}, 2)

tmp2502 := Call(__e, ns2_1set, symshen_4list_1_6str, tmp2493)


_ = tmp2502

tmp2503 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmaximum_1print_1sequence_1size_d)
}
__typedArg0 := sym_dmaximum_1print_1sequence_1size_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp2504 := Call(__e, ns2_1set, symshen_4maxseq, tmp2503)


_ = tmp2504

tmp2505 := MakeNative(func(__e *ControlFlow) {
V5673 := __e.Get(1)
_ = V5673
V5674 := __e.Get(2)
_ = V5674
V5675 := __e.Get(3)
_ = V5675
tmp2526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5673)
}
__typedArg0 := Nil
__typedArg1 := V5673
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2526 {
__e.Return(MakeString(""))
return
} else {
tmp2524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V5675)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V5675
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2524 {
__e.Return(MakeString("... etc"))
return
} else {
tmp2522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5673)
}
__typedArg0 := V5673
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2518 Obj

if True == tmp2522 {
tmp2520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5673)
}
__typedArg0 := V5673
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2520)
}
__typedArg0 := Nil
__typedArg1 := tmp2520
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2519 Obj

if True == tmp2521 {
ifres2519 = True


} else {
ifres2519 = False


}

ifres2518 = ifres2519


} else {
ifres2518 = False


}

if True == ifres2518 {
tmp2506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5673)
}
__typedArg0 := V5673
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4arg_1_6str), tmp2506, V5674)
return


} else {
tmp2516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5673)
}
__typedArg0 := V5673
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2516 {
tmp2507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5673)
}
__typedArg0 := V5673
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2508 := Call(__e, PrimFunc(symshen_4arg_1_6str), tmp2507, V5674)


tmp2509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5673)
}
__typedArg0 := V5673
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2511 := Call(__e, PrimFunc(symshen_4iter_1list), tmp2509, V5674, (func() Obj {
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


tmp2512 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp2511)


__e.TailApply(PrimFunc(sym_8s), tmp2508, tmp2512)
return


} else {
tmp2513 := Call(__e, PrimFunc(symshen_4arg_1_6str), V5673, V5674)


tmp2514 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp2513)


__e.TailApply(PrimFunc(sym_8s), MakeString("|"), tmp2514)
return


}


}


}


}


}, 3)

tmp2527 := Call(__e, ns2_1set, symshen_4iter_1list, tmp2505)


_ = tmp2527

tmp2528 := MakeNative(func(__e *ControlFlow) {
V5678 := __e.Get(1)
_ = V5678
V5679 := __e.Get(2)
_ = V5679
tmp2533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4a, V5679)
}
__typedArg0 := symshen_4a
__typedArg1 := V5679
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2533 {
__e.Return(V5678)
return
} else {
tmp2529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(MakeNumber(34))
}
__typedArg0 := MakeNumber(34)
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp2530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(MakeNumber(34))
}
__typedArg0 := MakeNumber(34)
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp2531 := Call(__e, PrimFunc(sym_8s), V5678, tmp2530)


__e.TailApply(PrimFunc(sym_8s), tmp2529, tmp2531)
return


}


}, 2)

tmp2534 := Call(__e, ns2_1set, symshen_4str_1_6str, tmp2528)


_ = tmp2534

tmp2535 := MakeNative(func(__e *ControlFlow) {
V5680 := __e.Get(1)
_ = V5680
V5681 := __e.Get(2)
_ = V5681
tmp2548 := Call(__e, PrimFunc(symshen_4print_1vector_2), V5680)


if True == tmp2548 {
tmp2536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V5680, MakeNumber(0))
}
__typedArg0 := V5680
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp2537 := Call(__e, PrimFunc(symfn), tmp2536)


__e.TailApply(tmp2537, V5680)
return


} else {
tmp2546 := Call(__e, PrimFunc(symvector_2), V5680)


if True == tmp2546 {
tmp2538 := Call(__e, PrimFunc(symshen_4maxseq))


tmp2539 := Call(__e, PrimFunc(symshen_4iter_1vector), V5680, MakeNumber(1), V5681, tmp2538)


tmp2540 := Call(__e, PrimFunc(sym_8s), tmp2539, MakeString(">"))


__e.TailApply(PrimFunc(sym_8s), MakeString("<"), tmp2540)
return


} else {
tmp2541 := Call(__e, PrimFunc(symshen_4maxseq))


tmp2542 := Call(__e, PrimFunc(symshen_4iter_1vector), V5680, MakeNumber(0), V5681, tmp2541)


tmp2543 := Call(__e, PrimFunc(sym_8s), tmp2542, MakeString(">>"))


tmp2544 := Call(__e, PrimFunc(sym_8s), MakeString("<"), tmp2543)


__e.TailApply(PrimFunc(sym_8s), MakeString("<"), tmp2544)
return


}


}


}, 2)

tmp2549 := Call(__e, ns2_1set, symshen_4vector_1_6str, tmp2535)


_ = tmp2549

tmp2550 := MakeNative(func(__e *ControlFlow) {
V5682 := __e.Get(1)
_ = V5682
tmp2551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V5682, MakeNumber(0))
}
__typedArg0 := V5682
__typedArg1 := MakeNumber(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

W56832105 := tmp2551
_ = W56832105

tmp2558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W56832105, symshen_4tuple)
}
__typedArg0 := W56832105
__typedArg1 := symshen_4tuple
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2558 {
__e.Return(True)
return
} else {
tmp2556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W56832105, symshen_4pvar)
}
__typedArg0 := W56832105
__typedArg1 := symshen_4pvar
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2556 {
__e.Return(True)
return
} else {
tmp2553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(W56832105)
}
__typedArg0 := W56832105
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp2553)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp2553
return Call(__e, PrimFunc(symnot), __typedArg0)
})() {
__e.TailApply(PrimFunc(symshen_4fbound_2), W56832105)
return
} else {
__e.Return(False)
return
}


}


}


}, 1)

tmp2559 := Call(__e, ns2_1set, symshen_4print_1vector_2, tmp2550)


_ = tmp2559

tmp2560 := MakeNative(func(__e *ControlFlow) {
V5684 := __e.Get(1)
_ = V5684
tmp2561 := Call(__e, PrimFunc(symarity), V5684)


tmp2562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp2561, MakeNumber(-1))
}
__typedArg0 := tmp2561
__typedArg1 := MakeNumber(-1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp2562)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp2562
return Call(__e, PrimFunc(symnot), __typedArg0)
})())
return


}, 1)

tmp2563 := Call(__e, ns2_1set, symshen_4fbound_2, tmp2560)


_ = tmp2563

tmp2564 := MakeNative(func(__e *ControlFlow) {
V5685 := __e.Get(1)
_ = V5685
tmp2565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V5685, MakeNumber(1))
}
__typedArg0 := V5685
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp2566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V5685, MakeNumber(2))
}
__typedArg0 := V5685
__typedArg1 := MakeNumber(2)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp2567 := Call(__e, PrimFunc(symshen_4app), tmp2566, MakeString(")"), symshen_4s)


tmp2569 := Call(__e, PrimFunc(symshen_4app), tmp2565, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" "))
__typedS1, __typedOK1 := TypedString(tmp2567)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" ")
__typedArg1 := tmp2567
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("(@p "))
__typedS1, __typedOK1 := TypedString(tmp2569)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("(@p ")
__typedArg1 := tmp2569
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp2570 := Call(__e, ns2_1set, symshen_4tuple, tmp2564)


_ = tmp2570

tmp2571 := MakeNative(func(__e *ControlFlow) {
V5692 := __e.Get(1)
_ = V5692
V5693 := __e.Get(2)
_ = V5693
V5694 := __e.Get(3)
_ = V5694
V5695 := __e.Get(4)
_ = V5695
tmp2589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V5695)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V5695
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2589 {
__e.Return(MakeString("... etc"))
return
} else {
tmp2572 := MakeNative(func(__e *ControlFlow) {
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

tmp2573 := MakeNative(func(__e *ControlFlow) {
Z5697 := __e.Get(1)
_ = Z5697
__e.Return(symshen_4out_1of_1bounds)
return
}, 1)

tmp2574 := Call(__e, try_1catch, tmp2572, tmp2573)


W56962106 := tmp2574
_ = W56962106

tmp2575 := MakeNative(func(__e *ControlFlow) {
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

tmp2577 := MakeNative(func(__e *ControlFlow) {
Z5699 := __e.Get(1)
_ = Z5699
__e.Return(symshen_4out_1of_1bounds)
return
}, 1)

tmp2578 := Call(__e, try_1catch, tmp2575, tmp2577)


W56982107 := tmp2578
_ = W56982107

tmp2587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W56962106, symshen_4out_1of_1bounds)
}
__typedArg0 := W56962106
__typedArg1 := symshen_4out_1of_1bounds
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2587 {
__e.Return(MakeString(""))
return
} else {
tmp2585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W56982107, symshen_4out_1of_1bounds)
}
__typedArg0 := W56982107
__typedArg1 := symshen_4out_1of_1bounds
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp2585 {
__e.TailApply(PrimFunc(symshen_4arg_1_6str), W56962106, V5694)
return
} else {
tmp2579 := Call(__e, PrimFunc(symshen_4arg_1_6str), W56962106, V5694)


tmp2580 := (func() Obj {
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

tmp2582 := Call(__e, PrimFunc(symshen_4iter_1vector), V5692, tmp2580, V5694, (func() Obj {
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


tmp2583 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp2582)


__e.TailApply(PrimFunc(sym_8s), tmp2579, tmp2583)
return


}


}


}


}, 4)

tmp2590 := Call(__e, ns2_1set, symshen_4iter_1vector, tmp2571)


_ = tmp2590

tmp2591 := MakeNative(func(__e *ControlFlow) {
V5700 := __e.Get(1)
_ = V5700
tmp2592 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V5700)
}
__typedArg0 := V5700
return Call(__e, PrimFunc(symstr), __typedArg0)
})())
return
}, 0)

tmp2593 := MakeNative(func(__e *ControlFlow) {
Z5701 := __e.Get(1)
_ = Z5701
__e.TailApply(PrimFunc(symshen_4funexstring))
return
}, 1)

__e.TailApply(try_1catch, tmp2592, tmp2593)
return


}, 1)

tmp2594 := Call(__e, ns2_1set, symshen_4atom_1_6str, tmp2591)


_ = tmp2594

tmp2595 := MakeNative(func(__e *ControlFlow) {
tmp2596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString("x"))
}
__typedArg0 := MakeString("x")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp2597 := Call(__e, PrimFunc(symgensym), tmp2596)


tmp2598 := Call(__e, PrimFunc(symshen_4arg_1_6str), tmp2597, symshen_4a)


tmp2599 := Call(__e, PrimFunc(sym_8s), tmp2598, MakeString("\x11"))


tmp2600 := Call(__e, PrimFunc(sym_8s), MakeString("e"), tmp2599)


tmp2601 := Call(__e, PrimFunc(sym_8s), MakeString("n"), tmp2600)


tmp2602 := Call(__e, PrimFunc(sym_8s), MakeString("u"), tmp2601)


tmp2603 := Call(__e, PrimFunc(sym_8s), MakeString("f"), tmp2602)


__e.TailApply(PrimFunc(sym_8s), MakeString("\x10"), tmp2603)
return


}, 0)

tmp2604 := Call(__e, ns2_1set, symshen_4funexstring, tmp2595)


_ = tmp2604

tmp2605 := MakeNative(func(__e *ControlFlow) {
V5702 := __e.Get(1)
_ = V5702
tmp2609 := Call(__e, PrimFunc(symempty_2), V5702)


if True == tmp2609 {
__e.Return(True)
return
} else {
tmp2607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5702)
}
__typedArg0 := V5702
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2607 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

__e.TailApply(ns2_1set, symshen_4list_2, tmp2605)
return




}, 0)

