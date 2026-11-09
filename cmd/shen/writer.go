package main

import . "github.com/pyrex41/shen-go/kl"

var WriterMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp2043 := MakeNative1(func(__e *ControlFlow, V7009 Obj) Obj {
tmp2044 := Call2(__e, SymFn(symshen_4insert), V7009, MakeString("~S"))


let__2038 := tmp2044
_ = let__2038

tmp2045 := Call0(__e, SymFn(symstoutput))


tmp2046 := Call2(__e, SymFn(sympr), let__2038, tmp2045)


let__2039 := tmp2046
_ = let__2039

return V7009


panic("unreachable")
})

tmp2047 := Call2(__e, ns2_1set, symprint, tmp2043)


_ = tmp2047

tmp2048 := MakeNative2(func(__e *ControlFlow, V7012, V7013 Obj) Obj {
tmp2053 := GValue(__e, sym_dhush_d)

if True == tmp2053 {
return V7012
} else {
tmp2051 := Call1(__e, SymFn(symchar_1stream_2), V7013)


if True == tmp2051 {
return TailCall2(__e, SymFn(symwrite_1string), V7012, V7013)
} else {
tmp2049 := Call2(__e, SymFn(symshen_4string_1_6byte), V7012, MakeInteger(0))


return TailCall4(__e, SymFn(symshen_4write_1chars), V7012, V7013, tmp2049, MakeInteger(1))


}


}


panic("unreachable")
})

tmp2054 := Call2(__e, ns2_1set, sympr, tmp2048)


_ = tmp2054

tmp2055 := MakeNative2(func(__e *ControlFlow, V7014, V7015 Obj) Obj {
tmp2056 := MakeNative0(func(__e *ControlFlow) Obj {
tmp2057 := PrimGuard2(__e, sympos, PrimPos, V7014, V7015)

return PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp2057)


panic("unreachable")
})

tmp2058 := MakeNative1(func(__e *ControlFlow, Z7016 Obj) Obj {
return symshen_4eos
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp2056, tmp2058)


panic("unreachable")
})

tmp2059 := Call2(__e, ns2_1set, symshen_4string_1_6byte, tmp2055)


_ = tmp2059

tmp2060 := (func() Obj {
var __self Obj
__self = MakeNative4(func(__e *ControlFlow, __self1, __self2, __self3, __self4 Obj) Obj {
__selftop:
V7017 := __self1
_ = V7017
V7018 := __self2
_ = V7018
V7019 := __self3
_ = V7019
V7020 := __self4
_ = V7020
tmp2065 := GEq(__e, symshen_4eos, V7019)

if True == tmp2065 {
return V7017
} else {
tmp2061 := PrimGuard2(__e, symwrite_1byte, PrimWriteByte, V7019, V7018)

_ = tmp2061

tmp2062 := Call2(__e, SymFn(symshen_4string_1_6byte), V7017, V7020)


if SymFn(symshen_4write_1chars) == __self {
__self1, __self2, __self3, __self4 = V7017, V7018, tmp2062, GAdd(__e, V7020, MakeInteger(1))
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4write_1chars), V7017, V7018, tmp2062, GAdd(__e, V7020, MakeInteger(1)))


}


panic("unreachable")
})
return __self
})()

tmp2066 := Call2(__e, ns2_1set, symshen_4write_1chars, tmp2060)


_ = tmp2066

tmp2067 := MakeNative2(func(__e *ControlFlow, V7021, V7022 Obj) Obj {
tmp2072 := PrimGuard1(__e, symstring_2, PrimIsString, V7021)

if True == tmp2072 {
tmp2068 := Call1(__e, SymFn(symshen_4proc_1nl), V7021)


return TailCall2(__e, SymFn(symshen_4mkstr_1l), tmp2068, V7022)


} else {
tmp2069 := GCons(__e, V7021, Nil)

tmp2070 := GCons(__e, symshen_4proc_1nl, tmp2069)

return TailCall2(__e, SymFn(symshen_4mkstr_1r), tmp2070, V7022)


}


panic("unreachable")
})

tmp2073 := Call2(__e, ns2_1set, symshen_4mkstr, tmp2067)


_ = tmp2073

tmp2074 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V7027 := __self1
_ = V7027
V7028 := __self2
_ = V7028
tmp2081 := GEq(__e, Nil, V7028)

if True == tmp2081 {
return V7027
} else {
tmp2079 := GConsP(__e, V7028)

if True == tmp2079 {
tmp2075 := GHd(__e, V7028)

tmp2076 := Call2(__e, SymFn(symshen_4insert_1l), tmp2075, V7027)


tmp2077 := GTl(__e, V7028)

if SymFn(symshen_4mkstr_1l) == __self {
__self1, __self2 = tmp2076, tmp2077
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4mkstr_1l), tmp2076, tmp2077)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.mkstr-l"))
}


}


panic("unreachable")
})
return __self
})()

tmp2082 := Call2(__e, ns2_1set, symshen_4mkstr_1l, tmp2074)


_ = tmp2082

tmp2083 := MakeNative2(func(__e *ControlFlow, V7035, V7036 Obj) Obj {
tmp2203 := GEq(__e, MakeString(""), V7036)

if True == tmp2203 {
return MakeString("")
} else {
tmp2201 := Call1(__e, SymFn(symshen_4_7string_2), V7036)


var ifres2191 Obj

if True == tmp2201 {
tmp2199 := Call1(__e, SymFn(symhdstr), V7036)


tmp2200 := GEq(__e, MakeString("~"), tmp2199)

var ifres2192 Obj

if True == tmp2200 {
tmp2198 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, V7036))


var ifres2193 Obj

if True == tmp2198 {
tmp2195 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, V7036))


tmp2196 := GEq(__e, MakeString("A"), tmp2195)

ifres2193 = tmp2196


} else {
ifres2193 = False


}

ifres2192 = ifres2193


} else {
ifres2192 = False


}

ifres2191 = ifres2192


} else {
ifres2191 = False


}

if True == ifres2191 {
tmp2085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7036)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V7036)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})()

tmp2086 := GCons(__e, symshen_4a, Nil)

tmp2087 := GCons(__e, tmp2085, tmp2086)

tmp2088 := GCons(__e, V7035, tmp2087)

return GCons(__e, symshen_4app, tmp2088)


} else {
tmp2189 := Call1(__e, SymFn(symshen_4_7string_2), V7036)


var ifres2179 Obj

if True == tmp2189 {
tmp2187 := Call1(__e, SymFn(symhdstr), V7036)


tmp2188 := GEq(__e, MakeString("~"), tmp2187)

var ifres2180 Obj

if True == tmp2188 {
tmp2186 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, V7036))


var ifres2181 Obj

if True == tmp2186 {
tmp2183 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, V7036))


tmp2184 := GEq(__e, MakeString("R"), tmp2183)

ifres2181 = tmp2184


} else {
ifres2181 = False


}

ifres2180 = ifres2181


} else {
ifres2180 = False


}

ifres2179 = ifres2180


} else {
ifres2179 = False


}

if True == ifres2179 {
tmp2090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7036)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V7036)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})()

tmp2091 := GCons(__e, symshen_4r, Nil)

tmp2092 := GCons(__e, tmp2090, tmp2091)

tmp2093 := GCons(__e, V7035, tmp2092)

return GCons(__e, symshen_4app, tmp2093)


} else {
tmp2177 := Call1(__e, SymFn(symshen_4_7string_2), V7036)


var ifres2167 Obj

if True == tmp2177 {
tmp2175 := Call1(__e, SymFn(symhdstr), V7036)


tmp2176 := GEq(__e, MakeString("~"), tmp2175)

var ifres2168 Obj

if True == tmp2176 {
tmp2174 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, V7036))


var ifres2169 Obj

if True == tmp2174 {
tmp2171 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, V7036))


tmp2172 := GEq(__e, MakeString("S"), tmp2171)

ifres2169 = tmp2172


} else {
ifres2169 = False


}

ifres2168 = ifres2169


} else {
ifres2168 = False


}

ifres2167 = ifres2168


} else {
ifres2167 = False


}

if True == ifres2167 {
tmp2095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7036)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V7036)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})()

tmp2096 := GCons(__e, symshen_4s, Nil)

tmp2097 := GCons(__e, tmp2095, tmp2096)

tmp2098 := GCons(__e, V7035, tmp2097)

return GCons(__e, symshen_4app, tmp2098)


} else {
tmp2165 := Call1(__e, SymFn(symshen_4_7string_2), V7036)


if True == tmp2165 {
tmp2099 := Call1(__e, SymFn(symhdstr), V7036)


tmp2101 := Call2(__e, SymFn(symshen_4insert_1l), V7035, PrimGuard1(__e, symtlstr, PrimTailString, V7036))


tmp2102 := GCons(__e, tmp2101, Nil)

tmp2103 := GCons(__e, tmp2099, tmp2102)

tmp2104 := GCons(__e, symcn, tmp2103)

return TailCall1(__e, SymFn(symshen_4factor_1cn), tmp2104)


} else {
tmp2163 := GConsP(__e, V7036)

var ifres2148 Obj

if True == tmp2163 {
tmp2161 := GHd(__e, V7036)

tmp2162 := GEq(__e, symcn, tmp2161)

var ifres2149 Obj

if True == tmp2162 {
tmp2159 := GTl(__e, V7036)

tmp2160 := GConsP(__e, tmp2159)

var ifres2150 Obj

if True == tmp2160 {
tmp2156 := GTl(__e, V7036)

tmp2157 := GTl(__e, tmp2156)

tmp2158 := GConsP(__e, tmp2157)

var ifres2151 Obj

if True == tmp2158 {
tmp2152 := GTl(__e, V7036)

tmp2153 := GTl(__e, tmp2152)

tmp2154 := GTl(__e, tmp2153)

tmp2155 := GEq(__e, Nil, tmp2154)

ifres2151 = tmp2155


} else {
ifres2151 = False


}

ifres2150 = ifres2151


} else {
ifres2150 = False


}

ifres2149 = ifres2150


} else {
ifres2149 = False


}

ifres2148 = ifres2149


} else {
ifres2148 = False


}

if True == ifres2148 {
tmp2105 := GTl(__e, V7036)

tmp2106 := GHd(__e, tmp2105)

tmp2107 := GTl(__e, V7036)

tmp2108 := GTl(__e, tmp2107)

tmp2109 := GHd(__e, tmp2108)

tmp2110 := Call2(__e, SymFn(symshen_4insert_1l), V7035, tmp2109)


tmp2111 := GCons(__e, tmp2110, Nil)

tmp2112 := GCons(__e, tmp2106, tmp2111)

return GCons(__e, symcn, tmp2112)


} else {
tmp2146 := GConsP(__e, V7036)

var ifres2125 Obj

if True == tmp2146 {
tmp2144 := GHd(__e, V7036)

tmp2145 := GEq(__e, symshen_4app, tmp2144)

var ifres2126 Obj

if True == tmp2145 {
tmp2142 := GTl(__e, V7036)

tmp2143 := GConsP(__e, tmp2142)

var ifres2127 Obj

if True == tmp2143 {
tmp2139 := GTl(__e, V7036)

tmp2140 := GTl(__e, tmp2139)

tmp2141 := GConsP(__e, tmp2140)

var ifres2128 Obj

if True == tmp2141 {
tmp2135 := GTl(__e, V7036)

tmp2136 := GTl(__e, tmp2135)

tmp2137 := GTl(__e, tmp2136)

tmp2138 := GConsP(__e, tmp2137)

var ifres2129 Obj

if True == tmp2138 {
tmp2130 := GTl(__e, V7036)

tmp2131 := GTl(__e, tmp2130)

tmp2132 := GTl(__e, tmp2131)

tmp2133 := GTl(__e, tmp2132)

tmp2134 := GEq(__e, Nil, tmp2133)

ifres2129 = tmp2134


} else {
ifres2129 = False


}

ifres2128 = ifres2129


} else {
ifres2128 = False


}

ifres2127 = ifres2128


} else {
ifres2127 = False


}

ifres2126 = ifres2127


} else {
ifres2126 = False


}

ifres2125 = ifres2126


} else {
ifres2125 = False


}

if True == ifres2125 {
tmp2113 := GTl(__e, V7036)

tmp2114 := GHd(__e, tmp2113)

tmp2115 := GTl(__e, V7036)

tmp2116 := GTl(__e, tmp2115)

tmp2117 := GHd(__e, tmp2116)

tmp2118 := Call2(__e, SymFn(symshen_4insert_1l), V7035, tmp2117)


tmp2119 := GTl(__e, V7036)

tmp2120 := GTl(__e, tmp2119)

tmp2121 := GTl(__e, tmp2120)

tmp2122 := GCons(__e, tmp2118, tmp2121)

tmp2123 := GCons(__e, tmp2114, tmp2122)

return GCons(__e, symshen_4app, tmp2123)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.insert-l"))
}


}


}


}


}


}


}


panic("unreachable")
})

tmp2204 := Call2(__e, ns2_1set, symshen_4insert_1l, tmp2083)


_ = tmp2204

tmp2205 := MakeNative1(func(__e *ControlFlow, V7037 Obj) Obj {
tmp2279 := GConsP(__e, V7037)

var ifres2221 Obj

if True == tmp2279 {
tmp2277 := GHd(__e, V7037)

tmp2278 := GEq(__e, symcn, tmp2277)

var ifres2222 Obj

if True == tmp2278 {
tmp2275 := GTl(__e, V7037)

tmp2276 := GConsP(__e, tmp2275)

var ifres2223 Obj

if True == tmp2276 {
tmp2272 := GTl(__e, V7037)

tmp2273 := GTl(__e, tmp2272)

tmp2274 := GConsP(__e, tmp2273)

var ifres2224 Obj

if True == tmp2274 {
tmp2268 := GTl(__e, V7037)

tmp2269 := GTl(__e, tmp2268)

tmp2270 := GHd(__e, tmp2269)

tmp2271 := GConsP(__e, tmp2270)

var ifres2225 Obj

if True == tmp2271 {
tmp2263 := GTl(__e, V7037)

tmp2264 := GTl(__e, tmp2263)

tmp2265 := GHd(__e, tmp2264)

tmp2266 := GHd(__e, tmp2265)

tmp2267 := GEq(__e, symcn, tmp2266)

var ifres2226 Obj

if True == tmp2267 {
tmp2258 := GTl(__e, V7037)

tmp2259 := GTl(__e, tmp2258)

tmp2260 := GHd(__e, tmp2259)

tmp2261 := GTl(__e, tmp2260)

tmp2262 := GConsP(__e, tmp2261)

var ifres2227 Obj

if True == tmp2262 {
tmp2252 := GTl(__e, V7037)

tmp2253 := GTl(__e, tmp2252)

tmp2254 := GHd(__e, tmp2253)

tmp2255 := GTl(__e, tmp2254)

tmp2256 := GTl(__e, tmp2255)

tmp2257 := GConsP(__e, tmp2256)

var ifres2228 Obj

if True == tmp2257 {
tmp2245 := GTl(__e, V7037)

tmp2246 := GTl(__e, tmp2245)

tmp2247 := GHd(__e, tmp2246)

tmp2248 := GTl(__e, tmp2247)

tmp2249 := GTl(__e, tmp2248)

tmp2250 := GTl(__e, tmp2249)

tmp2251 := GEq(__e, Nil, tmp2250)

var ifres2229 Obj

if True == tmp2251 {
tmp2241 := GTl(__e, V7037)

tmp2242 := GTl(__e, tmp2241)

tmp2243 := GTl(__e, tmp2242)

tmp2244 := GEq(__e, Nil, tmp2243)

var ifres2230 Obj

if True == tmp2244 {
tmp2238 := GTl(__e, V7037)

tmp2239 := GHd(__e, tmp2238)

tmp2240 := PrimGuard1(__e, symstring_2, PrimIsString, tmp2239)

var ifres2231 Obj

if True == tmp2240 {
tmp2232 := GTl(__e, V7037)

tmp2233 := GTl(__e, tmp2232)

tmp2234 := GHd(__e, tmp2233)

tmp2235 := GTl(__e, tmp2234)

tmp2236 := GHd(__e, tmp2235)

tmp2237 := PrimGuard1(__e, symstring_2, PrimIsString, tmp2236)

ifres2231 = tmp2237


} else {
ifres2231 = False


}

ifres2230 = ifres2231


} else {
ifres2230 = False


}

ifres2229 = ifres2230


} else {
ifres2229 = False


}

ifres2228 = ifres2229


} else {
ifres2228 = False


}

ifres2227 = ifres2228


} else {
ifres2227 = False


}

ifres2226 = ifres2227


} else {
ifres2226 = False


}

ifres2225 = ifres2226


} else {
ifres2225 = False


}

ifres2224 = ifres2225


} else {
ifres2224 = False


}

ifres2223 = ifres2224


} else {
ifres2223 = False


}

ifres2222 = ifres2223


} else {
ifres2222 = False


}

ifres2221 = ifres2222


} else {
ifres2221 = False


}

if True == ifres2221 {
tmp2206 := GTl(__e, V7037)

tmp2207 := GHd(__e, tmp2206)

tmp2208 := GTl(__e, V7037)

tmp2209 := GTl(__e, tmp2208)

tmp2210 := GHd(__e, tmp2209)

tmp2211 := GTl(__e, tmp2210)

tmp2212 := GHd(__e, tmp2211)

tmp2213 := PrimGuard2(__e, symcn, PrimStringConcat, tmp2207, tmp2212)

tmp2214 := GTl(__e, V7037)

tmp2215 := GTl(__e, tmp2214)

tmp2216 := GHd(__e, tmp2215)

tmp2217 := GTl(__e, tmp2216)

tmp2218 := GTl(__e, tmp2217)

tmp2219 := GCons(__e, tmp2213, tmp2218)

return GCons(__e, symcn, tmp2219)


} else {
return V7037
}


panic("unreachable")
})

tmp2280 := Call2(__e, ns2_1set, symshen_4factor_1cn, tmp2205)


_ = tmp2280

tmp2281 := MakeNative1(func(__e *ControlFlow, V7040 Obj) Obj {
tmp2304 := GEq(__e, MakeString(""), V7040)

if True == tmp2304 {
return MakeString("")
} else {
tmp2302 := Call1(__e, SymFn(symshen_4_7string_2), V7040)


var ifres2292 Obj

if True == tmp2302 {
tmp2300 := Call1(__e, SymFn(symhdstr), V7040)


tmp2301 := GEq(__e, MakeString("~"), tmp2300)

var ifres2293 Obj

if True == tmp2301 {
tmp2299 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, V7040))


var ifres2294 Obj

if True == tmp2299 {
tmp2296 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, V7040))


tmp2297 := GEq(__e, MakeString("%"), tmp2296)

ifres2294 = tmp2297


} else {
ifres2294 = False


}

ifres2293 = ifres2294


} else {
ifres2293 = False


}

ifres2292 = ifres2293


} else {
ifres2292 = False


}

if True == ifres2292 {
tmp2282 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, MakeInteger(10))

tmp2285 := Call1(__e, SymFn(symshen_4proc_1nl), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7040)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V7040)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})())


return PrimGuard2(__e, symcn, PrimStringConcat, tmp2282, tmp2285)


} else {
tmp2290 := Call1(__e, SymFn(symshen_4_7string_2), V7040)


if True == tmp2290 {
tmp2286 := Call1(__e, SymFn(symhdstr), V7040)


tmp2288 := Call1(__e, SymFn(symshen_4proc_1nl), PrimGuard1(__e, symtlstr, PrimTailString, V7040))


return PrimGuard2(__e, symcn, PrimStringConcat, tmp2286, tmp2288)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.proc-nl"))
}


}


}


panic("unreachable")
})

tmp2305 := Call2(__e, ns2_1set, symshen_4proc_1nl, tmp2281)


_ = tmp2305

tmp2306 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V7045 := __self1
_ = V7045
V7046 := __self2
_ = V7046
tmp2315 := GEq(__e, Nil, V7046)

if True == tmp2315 {
return V7045
} else {
tmp2313 := GConsP(__e, V7046)

if True == tmp2313 {
tmp2307 := GHd(__e, V7046)

tmp2308 := GCons(__e, V7045, Nil)

tmp2309 := GCons(__e, tmp2307, tmp2308)

tmp2310 := GCons(__e, symshen_4insert, tmp2309)

tmp2311 := GTl(__e, V7046)

if SymFn(symshen_4mkstr_1r) == __self {
__self1, __self2 = tmp2310, tmp2311
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4mkstr_1r), tmp2310, tmp2311)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.mkstr-r"))
}


}


panic("unreachable")
})
return __self
})()

tmp2316 := Call2(__e, ns2_1set, symshen_4mkstr_1r, tmp2306)


_ = tmp2316

tmp2317 := MakeNative2(func(__e *ControlFlow, V7047, V7048 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4insert_1h), V7047, V7048, MakeString(""))
panic("unreachable")
})

tmp2318 := Call2(__e, ns2_1set, symshen_4insert, tmp2317)


_ = tmp2318

tmp2319 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V7057 := __self1
_ = V7057
V7058 := __self2
_ = V7058
V7059 := __self3
_ = V7059
tmp2371 := GEq(__e, MakeString(""), V7058)

if True == tmp2371 {
return V7059
} else {
tmp2369 := Call1(__e, SymFn(symshen_4_7string_2), V7058)


var ifres2359 Obj

if True == tmp2369 {
tmp2367 := Call1(__e, SymFn(symhdstr), V7058)


tmp2368 := GEq(__e, MakeString("~"), tmp2367)

var ifres2360 Obj

if True == tmp2368 {
tmp2366 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, V7058))


var ifres2361 Obj

if True == tmp2366 {
tmp2363 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, V7058))


tmp2364 := GEq(__e, MakeString("A"), tmp2363)

ifres2361 = tmp2364


} else {
ifres2361 = False


}

ifres2360 = ifres2361


} else {
ifres2360 = False


}

ifres2359 = ifres2360


} else {
ifres2359 = False


}

if True == ifres2359 {
tmp2322 := Call3(__e, SymFn(symshen_4app), V7057, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7058)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V7058)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})(), symshen_4a)


return PrimGuard2(__e, symcn, PrimStringConcat, V7059, tmp2322)


} else {
tmp2357 := Call1(__e, SymFn(symshen_4_7string_2), V7058)


var ifres2347 Obj

if True == tmp2357 {
tmp2355 := Call1(__e, SymFn(symhdstr), V7058)


tmp2356 := GEq(__e, MakeString("~"), tmp2355)

var ifres2348 Obj

if True == tmp2356 {
tmp2354 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, V7058))


var ifres2349 Obj

if True == tmp2354 {
tmp2351 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, V7058))


tmp2352 := GEq(__e, MakeString("R"), tmp2351)

ifres2349 = tmp2352


} else {
ifres2349 = False


}

ifres2348 = ifres2349


} else {
ifres2348 = False


}

ifres2347 = ifres2348


} else {
ifres2347 = False


}

if True == ifres2347 {
tmp2325 := Call3(__e, SymFn(symshen_4app), V7057, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7058)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V7058)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})(), symshen_4r)


return PrimGuard2(__e, symcn, PrimStringConcat, V7059, tmp2325)


} else {
tmp2345 := Call1(__e, SymFn(symshen_4_7string_2), V7058)


var ifres2335 Obj

if True == tmp2345 {
tmp2343 := Call1(__e, SymFn(symhdstr), V7058)


tmp2344 := GEq(__e, MakeString("~"), tmp2343)

var ifres2336 Obj

if True == tmp2344 {
tmp2342 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, V7058))


var ifres2337 Obj

if True == tmp2342 {
tmp2339 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, V7058))


tmp2340 := GEq(__e, MakeString("S"), tmp2339)

ifres2337 = tmp2340


} else {
ifres2337 = False


}

ifres2336 = ifres2337


} else {
ifres2336 = False


}

ifres2335 = ifres2336


} else {
ifres2335 = False


}

if True == ifres2335 {
tmp2328 := Call3(__e, SymFn(symshen_4app), V7057, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7058)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V7058)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})(), symshen_4s)


return PrimGuard2(__e, symcn, PrimStringConcat, V7059, tmp2328)


} else {
tmp2333 := Call1(__e, SymFn(symshen_4_7string_2), V7058)


if True == tmp2333 {
tmp2329 := PrimGuard1(__e, symtlstr, PrimTailString, V7058)

tmp2330 := Call1(__e, SymFn(symhdstr), V7058)


if SymFn(symshen_4insert_1h) == __self {
__self1, __self2, __self3 = V7057, tmp2329, PrimGuard2(__e, symcn, PrimStringConcat, V7059, tmp2330)
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4insert_1h), V7057, tmp2329, PrimGuard2(__e, symcn, PrimStringConcat, V7059, tmp2330))


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.insert-h"))
}


}


}


}


}


panic("unreachable")
})
return __self
})()

tmp2372 := Call2(__e, ns2_1set, symshen_4insert_1h, tmp2319)


_ = tmp2372

tmp2373 := MakeNative3(func(__e *ControlFlow, V7060, V7061, V7062 Obj) Obj {
tmp2374 := Call2(__e, SymFn(symshen_4arg_1_6str), V7060, V7062)


return PrimGuard2(__e, symcn, PrimStringConcat, tmp2374, V7061)


panic("unreachable")
})

tmp2375 := Call2(__e, ns2_1set, symshen_4app, tmp2373)


_ = tmp2375

tmp2376 := MakeNative2(func(__e *ControlFlow, V7066, V7067 Obj) Obj {
tmp2384 := Call0(__e, SymFn(symfail))


tmp2385 := GEq(__e, V7066, tmp2384)

if True == tmp2385 {
return MakeString("...")
} else {
tmp2382 := Call1(__e, SymFn(symshen_4list_2), V7066)


if True == tmp2382 {
return TailCall2(__e, SymFn(symshen_4list_1_6str), V7066, V7067)
} else {
tmp2380 := PrimGuard1(__e, symstring_2, PrimIsString, V7066)

if True == tmp2380 {
return TailCall2(__e, SymFn(symshen_4str_1_6str), V7066, V7067)
} else {
tmp2378 := GVectorP(__e, V7066)

if True == tmp2378 {
return TailCall2(__e, SymFn(symshen_4vector_1_6str), V7066, V7067)
} else {
return TailCall1(__e, SymFn(symshen_4atom_1_6str), V7066)
}


}


}


}


panic("unreachable")
})

tmp2386 := Call2(__e, ns2_1set, symshen_4arg_1_6str, tmp2376)


_ = tmp2386

tmp2387 := MakeNative2(func(__e *ControlFlow, V7068, V7069 Obj) Obj {
tmp2395 := GEq(__e, symshen_4r, V7069)

if True == tmp2395 {
tmp2388 := Call0(__e, SymFn(symshen_4maxseq))


tmp2389 := Call3(__e, SymFn(symshen_4iter_1list), V7068, symshen_4r, tmp2388)


tmp2390 := Call2(__e, SymFn(sym_8s), tmp2389, MakeString(")"))


return TailCall2(__e, SymFn(sym_8s), MakeString("("), tmp2390)


} else {
tmp2391 := Call0(__e, SymFn(symshen_4maxseq))


tmp2392 := Call3(__e, SymFn(symshen_4iter_1list), V7068, V7069, tmp2391)


tmp2393 := Call2(__e, SymFn(sym_8s), tmp2392, MakeString("]"))


return TailCall2(__e, SymFn(sym_8s), MakeString("["), tmp2393)


}


panic("unreachable")
})

tmp2396 := Call2(__e, ns2_1set, symshen_4list_1_6str, tmp2387)


_ = tmp2396

tmp2397 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dmaximum_1print_1sequence_1size_d)
panic("unreachable")
})

tmp2398 := Call2(__e, ns2_1set, symshen_4maxseq, tmp2397)


_ = tmp2398

tmp2399 := MakeNative3(func(__e *ControlFlow, V7080, V7081, V7082 Obj) Obj {
tmp2419 := GEq(__e, Nil, V7080)

if True == tmp2419 {
return MakeString("")
} else {
tmp2417 := GEq(__e, MakeInteger(0), V7082)

if True == tmp2417 {
return MakeString("... etc")
} else {
tmp2415 := GConsP(__e, V7080)

var ifres2412 Obj

if True == tmp2415 {
tmp2413 := GTl(__e, V7080)

tmp2414 := GEq(__e, Nil, tmp2413)

ifres2412 = tmp2414


} else {
ifres2412 = False


}

if True == ifres2412 {
tmp2400 := GHd(__e, V7080)

return TailCall2(__e, SymFn(symshen_4arg_1_6str), tmp2400, V7081)


} else {
tmp2410 := GConsP(__e, V7080)

if True == tmp2410 {
tmp2401 := GHd(__e, V7080)

tmp2402 := Call2(__e, SymFn(symshen_4arg_1_6str), tmp2401, V7081)


tmp2403 := GTl(__e, V7080)

tmp2405 := Call3(__e, SymFn(symshen_4iter_1list), tmp2403, V7081, GSub(__e, V7082, MakeInteger(1)))


tmp2406 := Call2(__e, SymFn(sym_8s), MakeString(" "), tmp2405)


return TailCall2(__e, SymFn(sym_8s), tmp2402, tmp2406)


} else {
tmp2407 := Call2(__e, SymFn(symshen_4arg_1_6str), V7080, V7081)


tmp2408 := Call2(__e, SymFn(sym_8s), MakeString(" "), tmp2407)


return TailCall2(__e, SymFn(sym_8s), MakeString("|"), tmp2408)


}


}


}


}


panic("unreachable")
})

tmp2420 := Call2(__e, ns2_1set, symshen_4iter_1list, tmp2399)


_ = tmp2420

tmp2421 := MakeNative2(func(__e *ControlFlow, V7085, V7086 Obj) Obj {
tmp2426 := GEq(__e, symshen_4a, V7086)

if True == tmp2426 {
return V7085
} else {
tmp2422 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, MakeInteger(34))

tmp2423 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, MakeInteger(34))

tmp2424 := Call2(__e, SymFn(sym_8s), V7085, tmp2423)


return TailCall2(__e, SymFn(sym_8s), tmp2422, tmp2424)


}


panic("unreachable")
})

tmp2427 := Call2(__e, ns2_1set, symshen_4str_1_6str, tmp2421)


_ = tmp2427

tmp2428 := MakeNative2(func(__e *ControlFlow, V7087, V7088 Obj) Obj {
tmp2441 := Call1(__e, SymFn(symshen_4print_1vector_2), V7087)


if True == tmp2441 {
tmp2429 := GVecGet(__e, V7087, MakeInteger(0))

tmp2430 := Call1(__e, SymFn(symfn), tmp2429)


return TailCall1(__e, tmp2430, V7087)


} else {
tmp2439 := Call1(__e, SymFn(symvector_2), V7087)


if True == tmp2439 {
tmp2431 := Call0(__e, SymFn(symshen_4maxseq))


tmp2432 := Call4(__e, SymFn(symshen_4iter_1vector), V7087, MakeInteger(1), V7088, tmp2431)


tmp2433 := Call2(__e, SymFn(sym_8s), tmp2432, MakeString(">"))


return TailCall2(__e, SymFn(sym_8s), MakeString("<"), tmp2433)


} else {
tmp2434 := Call0(__e, SymFn(symshen_4maxseq))


tmp2435 := Call4(__e, SymFn(symshen_4iter_1vector), V7087, MakeInteger(0), V7088, tmp2434)


tmp2436 := Call2(__e, SymFn(sym_8s), tmp2435, MakeString(">>"))


tmp2437 := Call2(__e, SymFn(sym_8s), MakeString("<"), tmp2436)


return TailCall2(__e, SymFn(sym_8s), MakeString("<"), tmp2437)


}


}


panic("unreachable")
})

tmp2442 := Call2(__e, ns2_1set, symshen_4vector_1_6str, tmp2428)


_ = tmp2442

tmp2443 := MakeNative1(func(__e *ControlFlow, V7089 Obj) Obj {
tmp2444 := GValue(__e, symshen_4_dempty_1absvector_d)

return GEq(__e, V7089, tmp2444)


panic("unreachable")
})

tmp2445 := Call2(__e, ns2_1set, symshen_4empty_1absvector_2, tmp2443)


_ = tmp2445

tmp2446 := MakeNative1(func(__e *ControlFlow, V7090 Obj) Obj {
tmp2458 := Call1(__e, SymFn(symshen_4empty_1absvector_2), V7090)


if True == GNot(__e, tmp2458) {
tmp2447 := GVecGet(__e, V7090, MakeInteger(0))

let__2040 := tmp2447
_ = let__2040

tmp2456 := GEq(__e, let__2040, symshen_4tuple)

if True == tmp2456 {
return True
} else {
tmp2454 := GEq(__e, let__2040, symshen_4pvar)

if True == tmp2454 {
return True
} else {
tmp2452 := GEq(__e, let__2040, symshen_4dictionary)

if True == tmp2452 {
return True
} else {
tmp2449 := PrimGuard1(__e, symnumber_2, PrimIsNumber, let__2040)

if True == GNot(__e, tmp2449) {
return TailCall1(__e, SymFn(symshen_4fbound_2), let__2040)
} else {
return False
}


}


}


}


} else {
return False
}


panic("unreachable")
})

tmp2460 := Call2(__e, ns2_1set, symshen_4print_1vector_2, tmp2446)


_ = tmp2460

tmp2461 := MakeNative1(func(__e *ControlFlow, V7092 Obj) Obj {
tmp2462 := Call1(__e, SymFn(symarity), V7092)


tmp2463 := GEq(__e, tmp2462, MakeInteger(-1))

return GNot(__e, tmp2463)


panic("unreachable")
})

tmp2464 := Call2(__e, ns2_1set, symshen_4fbound_2, tmp2461)


_ = tmp2464

tmp2465 := MakeNative1(func(__e *ControlFlow, V7093 Obj) Obj {
tmp2466 := GVecGet(__e, V7093, MakeInteger(1))

tmp2467 := GVecGet(__e, V7093, MakeInteger(2))

tmp2468 := Call3(__e, SymFn(symshen_4app), tmp2467, MakeString(")"), symshen_4s)


tmp2470 := Call3(__e, SymFn(symshen_4app), tmp2466, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" "), tmp2468), symshen_4s)


return PrimGuard2(__e, symcn, PrimStringConcat, MakeString("(@p "), tmp2470)


panic("unreachable")
})

tmp2471 := Call2(__e, ns2_1set, symshen_4tuple, tmp2465)


_ = tmp2471

tmp2472 := MakeNative1(func(__e *ControlFlow, V7094 Obj) Obj {
return MakeString("(dict ...)")
panic("unreachable")
})

tmp2473 := Call2(__e, ns2_1set, symshen_4dictionary, tmp2472)


_ = tmp2473

tmp2474 := MakeNative4(func(__e *ControlFlow, V7101, V7102, V7103, V7104 Obj) Obj {
tmp2492 := GEq(__e, MakeInteger(0), V7104)

if True == tmp2492 {
return MakeString("... etc")
} else {
tmp2475 := MakeNative0(func(__e *ControlFlow) Obj {
return GVecGet(__e, V7101, V7102)
panic("unreachable")
})

tmp2476 := MakeNative1(func(__e *ControlFlow, Z7106 Obj) Obj {
return symshen_4out_1of_1bounds
panic("unreachable")
})

tmp2477 := Call2(__e, try_1catch, tmp2475, tmp2476)


let__2041 := tmp2477
_ = let__2041

tmp2478 := MakeNative0(func(__e *ControlFlow) Obj {
return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V7101, GAdd(__e, V7102, MakeInteger(1)))
}
__typedArg0 := V7101
__typedArg1 := GAdd(__e, V7102, MakeInteger(1))
return Call2(__e, SymFn(sym_5_1address), __typedArg0, __typedArg1)
})()


panic("unreachable")
})

tmp2480 := MakeNative1(func(__e *ControlFlow, Z7108 Obj) Obj {
return symshen_4out_1of_1bounds
panic("unreachable")
})

tmp2481 := Call2(__e, try_1catch, tmp2478, tmp2480)


let__2042 := tmp2481
_ = let__2042

tmp2490 := GEq(__e, let__2041, symshen_4out_1of_1bounds)

if True == tmp2490 {
return MakeString("")
} else {
tmp2488 := GEq(__e, let__2042, symshen_4out_1of_1bounds)

if True == tmp2488 {
return TailCall2(__e, SymFn(symshen_4arg_1_6str), let__2041, V7103)
} else {
tmp2482 := Call2(__e, SymFn(symshen_4arg_1_6str), let__2041, V7103)


tmp2483 := GAdd(__e, V7102, MakeInteger(1))

tmp2485 := Call4(__e, SymFn(symshen_4iter_1vector), V7101, tmp2483, V7103, GSub(__e, V7104, MakeInteger(1)))


tmp2486 := Call2(__e, SymFn(sym_8s), MakeString(" "), tmp2485)


return TailCall2(__e, SymFn(sym_8s), tmp2482, tmp2486)


}


}


}


panic("unreachable")
})

tmp2493 := Call2(__e, ns2_1set, symshen_4iter_1vector, tmp2474)


_ = tmp2493

tmp2494 := MakeNative1(func(__e *ControlFlow, V7109 Obj) Obj {
tmp2495 := MakeNative0(func(__e *ControlFlow) Obj {
return PrimGuard1(__e, symstr, PrimStr, V7109)
panic("unreachable")
})

tmp2496 := MakeNative1(func(__e *ControlFlow, Z7110 Obj) Obj {
return TailCall0(__e, SymFn(symshen_4funexstring))
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp2495, tmp2496)


panic("unreachable")
})

tmp2497 := Call2(__e, ns2_1set, symshen_4atom_1_6str, tmp2494)


_ = tmp2497

tmp2498 := MakeNative0(func(__e *ControlFlow) Obj {
tmp2499 := PrimGuard1(__e, symintern, PrimIntern, MakeString("x"))

tmp2500 := Call1(__e, SymFn(symgensym), tmp2499)


tmp2501 := Call2(__e, SymFn(symshen_4arg_1_6str), tmp2500, symshen_4a)


tmp2502 := Call2(__e, SymFn(sym_8s), tmp2501, MakeString("\x11"))


tmp2503 := Call2(__e, SymFn(sym_8s), MakeString("e"), tmp2502)


tmp2504 := Call2(__e, SymFn(sym_8s), MakeString("n"), tmp2503)


tmp2505 := Call2(__e, SymFn(sym_8s), MakeString("u"), tmp2504)


tmp2506 := Call2(__e, SymFn(sym_8s), MakeString("f"), tmp2505)


return TailCall2(__e, SymFn(sym_8s), MakeString("\x10"), tmp2506)


panic("unreachable")
})

tmp2507 := Call2(__e, ns2_1set, symshen_4funexstring, tmp2498)


_ = tmp2507

tmp2508 := MakeNative1(func(__e *ControlFlow, V7111 Obj) Obj {
tmp2510 := Call1(__e, SymFn(symempty_2), V7111)


if True == tmp2510 {
return True
} else {
return GConsP(__e, V7111)
}


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4list_2, tmp2508)




panic("unreachable")
})

