package main

import . "github.com/pyrex41/shen-go/kl"

var WriterMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp2063 := MakeNative1(func(__e *ControlFlow, V6988 Obj) Obj {
tmp2064 := Call2(__e, SymFn(symshen_4insert), V6988, MakeString("~S"))


let__2058 := tmp2064
_ = let__2058

tmp2065 := Call0(__e, SymFn(symstoutput))


tmp2066 := Call2(__e, SymFn(sympr), let__2058, tmp2065)


let__2059 := tmp2066
_ = let__2059

return V6988


panic("unreachable")
})

tmp2067 := Call2(__e, ns2_1set, symprint, tmp2063)


_ = tmp2067

tmp2068 := MakeNative2(func(__e *ControlFlow, V6991, V6992 Obj) Obj {
tmp2073 := GValue(__e, sym_dhush_d)

if True == tmp2073 {
return V6991
} else {
tmp2071 := Call1(__e, SymFn(symchar_1stream_2), V6992)


if True == tmp2071 {
return TailCall2(__e, SymFn(symwrite_1string), V6991, V6992)
} else {
tmp2069 := Call2(__e, SymFn(symshen_4string_1_6byte), V6991, MakeInteger(0))


return TailCall4(__e, SymFn(symshen_4write_1chars), V6991, V6992, tmp2069, MakeInteger(1))


}


}


panic("unreachable")
})

tmp2074 := Call2(__e, ns2_1set, sympr, tmp2068)


_ = tmp2074

tmp2075 := MakeNative2(func(__e *ControlFlow, V6993, V6994 Obj) Obj {
tmp2076 := MakeNative0(func(__e *ControlFlow) Obj {
tmp2077 := PrimGuard2(__e, sympos, PrimPos, V6993, V6994)

return PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp2077)


panic("unreachable")
})

tmp2078 := MakeNative1(func(__e *ControlFlow, Z6995 Obj) Obj {
return symshen_4eos
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp2076, tmp2078)


panic("unreachable")
})

tmp2079 := Call2(__e, ns2_1set, symshen_4string_1_6byte, tmp2075)


_ = tmp2079

tmp2080 := (func() Obj {
var __self Obj
__self = MakeNative4(func(__e *ControlFlow, __self1, __self2, __self3, __self4 Obj) Obj {
__selftop:
V6996 := __self1
_ = V6996
V6997 := __self2
_ = V6997
V6998 := __self3
_ = V6998
V6999 := __self4
_ = V6999
tmp2085 := GEq(__e, symshen_4eos, V6998)

if True == tmp2085 {
return V6996
} else {
tmp2081 := PrimGuard2(__e, symwrite_1byte, PrimWriteByte, V6998, V6997)

_ = tmp2081

tmp2082 := Call2(__e, SymFn(symshen_4string_1_6byte), V6996, V6999)


if SymFn(symshen_4write_1chars) == __self {
__self1, __self2, __self3, __self4 = V6996, V6997, tmp2082, GAdd(__e, V6999, MakeInteger(1))
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4write_1chars), V6996, V6997, tmp2082, GAdd(__e, V6999, MakeInteger(1)))


}


panic("unreachable")
})
return __self
})()

tmp2086 := Call2(__e, ns2_1set, symshen_4write_1chars, tmp2080)


_ = tmp2086

tmp2087 := MakeNative2(func(__e *ControlFlow, V7000, V7001 Obj) Obj {
tmp2092 := PrimGuard1(__e, symstring_2, PrimIsString, V7000)

if True == tmp2092 {
tmp2088 := Call1(__e, SymFn(symshen_4proc_1nl), V7000)


return TailCall2(__e, SymFn(symshen_4mkstr_1l), tmp2088, V7001)


} else {
tmp2089 := GCons(__e, V7000, Nil)

tmp2090 := GCons(__e, symshen_4proc_1nl, tmp2089)

return TailCall2(__e, SymFn(symshen_4mkstr_1r), tmp2090, V7001)


}


panic("unreachable")
})

tmp2093 := Call2(__e, ns2_1set, symshen_4mkstr, tmp2087)


_ = tmp2093

tmp2094 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V7006 := __self1
_ = V7006
V7007 := __self2
_ = V7007
tmp2101 := GEq(__e, Nil, V7007)

if True == tmp2101 {
return V7006
} else {
tmp2099 := GConsP(__e, V7007)

if True == tmp2099 {
tmp2095 := GHd(__e, V7007)

tmp2096 := Call2(__e, SymFn(symshen_4insert_1l), tmp2095, V7006)


tmp2097 := GTl(__e, V7007)

if SymFn(symshen_4mkstr_1l) == __self {
__self1, __self2 = tmp2096, tmp2097
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4mkstr_1l), tmp2096, tmp2097)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.mkstr-l"))
}


}


panic("unreachable")
})
return __self
})()

tmp2102 := Call2(__e, ns2_1set, symshen_4mkstr_1l, tmp2094)


_ = tmp2102

tmp2103 := MakeNative2(func(__e *ControlFlow, V7014, V7015 Obj) Obj {
tmp2241 := GEq(__e, MakeString(""), V7015)

if True == tmp2241 {
return MakeString("")
} else {
tmp2239 := Call1(__e, SymFn(symshen_4_7string_2), V7015)


var ifres2226 Obj

if True == tmp2239 {
tmp2237 := Call1(__e, SymFn(symhdstr), V7015)


tmp2238 := GEq(__e, MakeString("~"), tmp2237)

var ifres2228 Obj

if True == tmp2238 {
tmp2236 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, V7015))


var ifres2230 Obj

if True == tmp2236 {
tmp2233 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, V7015))


tmp2234 := GEq(__e, MakeString("A"), tmp2233)

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
__typedS0, __typedOK0 := TypedString(V7015)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V7015)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})()

tmp2106 := GCons(__e, symshen_4a, Nil)

tmp2107 := GCons(__e, tmp2105, tmp2106)

tmp2108 := GCons(__e, V7014, tmp2107)

return GCons(__e, symshen_4app, tmp2108)


} else {
tmp2224 := Call1(__e, SymFn(symshen_4_7string_2), V7015)


var ifres2211 Obj

if True == tmp2224 {
tmp2222 := Call1(__e, SymFn(symhdstr), V7015)


tmp2223 := GEq(__e, MakeString("~"), tmp2222)

var ifres2213 Obj

if True == tmp2223 {
tmp2221 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, V7015))


var ifres2215 Obj

if True == tmp2221 {
tmp2218 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, V7015))


tmp2219 := GEq(__e, MakeString("R"), tmp2218)

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
__typedS0, __typedOK0 := TypedString(V7015)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V7015)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})()

tmp2111 := GCons(__e, symshen_4r, Nil)

tmp2112 := GCons(__e, tmp2110, tmp2111)

tmp2113 := GCons(__e, V7014, tmp2112)

return GCons(__e, symshen_4app, tmp2113)


} else {
tmp2209 := Call1(__e, SymFn(symshen_4_7string_2), V7015)


var ifres2196 Obj

if True == tmp2209 {
tmp2207 := Call1(__e, SymFn(symhdstr), V7015)


tmp2208 := GEq(__e, MakeString("~"), tmp2207)

var ifres2198 Obj

if True == tmp2208 {
tmp2206 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, V7015))


var ifres2200 Obj

if True == tmp2206 {
tmp2203 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, V7015))


tmp2204 := GEq(__e, MakeString("S"), tmp2203)

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
__typedS0, __typedOK0 := TypedString(V7015)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V7015)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})()

tmp2116 := GCons(__e, symshen_4s, Nil)

tmp2117 := GCons(__e, tmp2115, tmp2116)

tmp2118 := GCons(__e, V7014, tmp2117)

return GCons(__e, symshen_4app, tmp2118)


} else {
tmp2194 := Call1(__e, SymFn(symshen_4_7string_2), V7015)


if True == tmp2194 {
tmp2119 := Call1(__e, SymFn(symhdstr), V7015)


tmp2121 := Call2(__e, SymFn(symshen_4insert_1l), V7014, PrimGuard1(__e, symtlstr, PrimTailString, V7015))


tmp2122 := GCons(__e, tmp2121, Nil)

tmp2123 := GCons(__e, tmp2119, tmp2122)

tmp2124 := GCons(__e, symcn, tmp2123)

return TailCall1(__e, SymFn(symshen_4factor_1cn), tmp2124)


} else {
tmp2192 := GConsP(__e, V7015)

var ifres2173 Obj

if True == tmp2192 {
tmp2190 := GHd(__e, V7015)

tmp2191 := GEq(__e, symcn, tmp2190)

var ifres2175 Obj

if True == tmp2191 {
tmp2188 := GTl(__e, V7015)

tmp2189 := GConsP(__e, tmp2188)

var ifres2177 Obj

if True == tmp2189 {
tmp2185 := GTl(__e, V7015)

tmp2186 := GTl(__e, tmp2185)

tmp2187 := GConsP(__e, tmp2186)

var ifres2179 Obj

if True == tmp2187 {
tmp2181 := GTl(__e, V7015)

tmp2182 := GTl(__e, tmp2181)

tmp2183 := GTl(__e, tmp2182)

tmp2184 := GEq(__e, Nil, tmp2183)

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
tmp2125 := GTl(__e, V7015)

tmp2126 := GHd(__e, tmp2125)

tmp2127 := GTl(__e, V7015)

tmp2128 := GTl(__e, tmp2127)

tmp2129 := GHd(__e, tmp2128)

tmp2130 := Call2(__e, SymFn(symshen_4insert_1l), V7014, tmp2129)


tmp2131 := GCons(__e, tmp2130, Nil)

tmp2132 := GCons(__e, tmp2126, tmp2131)

return GCons(__e, symcn, tmp2132)


} else {
tmp2171 := GConsP(__e, V7015)

var ifres2145 Obj

if True == tmp2171 {
tmp2169 := GHd(__e, V7015)

tmp2170 := GEq(__e, symshen_4app, tmp2169)

var ifres2147 Obj

if True == tmp2170 {
tmp2167 := GTl(__e, V7015)

tmp2168 := GConsP(__e, tmp2167)

var ifres2149 Obj

if True == tmp2168 {
tmp2164 := GTl(__e, V7015)

tmp2165 := GTl(__e, tmp2164)

tmp2166 := GConsP(__e, tmp2165)

var ifres2151 Obj

if True == tmp2166 {
tmp2160 := GTl(__e, V7015)

tmp2161 := GTl(__e, tmp2160)

tmp2162 := GTl(__e, tmp2161)

tmp2163 := GConsP(__e, tmp2162)

var ifres2153 Obj

if True == tmp2163 {
tmp2155 := GTl(__e, V7015)

tmp2156 := GTl(__e, tmp2155)

tmp2157 := GTl(__e, tmp2156)

tmp2158 := GTl(__e, tmp2157)

tmp2159 := GEq(__e, Nil, tmp2158)

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
tmp2133 := GTl(__e, V7015)

tmp2134 := GHd(__e, tmp2133)

tmp2135 := GTl(__e, V7015)

tmp2136 := GTl(__e, tmp2135)

tmp2137 := GHd(__e, tmp2136)

tmp2138 := Call2(__e, SymFn(symshen_4insert_1l), V7014, tmp2137)


tmp2139 := GTl(__e, V7015)

tmp2140 := GTl(__e, tmp2139)

tmp2141 := GTl(__e, tmp2140)

tmp2142 := GCons(__e, tmp2138, tmp2141)

tmp2143 := GCons(__e, tmp2134, tmp2142)

return GCons(__e, symshen_4app, tmp2143)


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

tmp2242 := Call2(__e, ns2_1set, symshen_4insert_1l, tmp2103)


_ = tmp2242

tmp2243 := MakeNative1(func(__e *ControlFlow, V7016 Obj) Obj {
tmp2328 := GConsP(__e, V7016)

var ifres2259 Obj

if True == tmp2328 {
tmp2326 := GHd(__e, V7016)

tmp2327 := GEq(__e, symcn, tmp2326)

var ifres2261 Obj

if True == tmp2327 {
tmp2324 := GTl(__e, V7016)

tmp2325 := GConsP(__e, tmp2324)

var ifres2263 Obj

if True == tmp2325 {
tmp2321 := GTl(__e, V7016)

tmp2322 := GTl(__e, tmp2321)

tmp2323 := GConsP(__e, tmp2322)

var ifres2265 Obj

if True == tmp2323 {
tmp2317 := GTl(__e, V7016)

tmp2318 := GTl(__e, tmp2317)

tmp2319 := GHd(__e, tmp2318)

tmp2320 := GConsP(__e, tmp2319)

var ifres2267 Obj

if True == tmp2320 {
tmp2312 := GTl(__e, V7016)

tmp2313 := GTl(__e, tmp2312)

tmp2314 := GHd(__e, tmp2313)

tmp2315 := GHd(__e, tmp2314)

tmp2316 := GEq(__e, symcn, tmp2315)

var ifres2269 Obj

if True == tmp2316 {
tmp2307 := GTl(__e, V7016)

tmp2308 := GTl(__e, tmp2307)

tmp2309 := GHd(__e, tmp2308)

tmp2310 := GTl(__e, tmp2309)

tmp2311 := GConsP(__e, tmp2310)

var ifres2271 Obj

if True == tmp2311 {
tmp2301 := GTl(__e, V7016)

tmp2302 := GTl(__e, tmp2301)

tmp2303 := GHd(__e, tmp2302)

tmp2304 := GTl(__e, tmp2303)

tmp2305 := GTl(__e, tmp2304)

tmp2306 := GConsP(__e, tmp2305)

var ifres2273 Obj

if True == tmp2306 {
tmp2294 := GTl(__e, V7016)

tmp2295 := GTl(__e, tmp2294)

tmp2296 := GHd(__e, tmp2295)

tmp2297 := GTl(__e, tmp2296)

tmp2298 := GTl(__e, tmp2297)

tmp2299 := GTl(__e, tmp2298)

tmp2300 := GEq(__e, Nil, tmp2299)

var ifres2275 Obj

if True == tmp2300 {
tmp2290 := GTl(__e, V7016)

tmp2291 := GTl(__e, tmp2290)

tmp2292 := GTl(__e, tmp2291)

tmp2293 := GEq(__e, Nil, tmp2292)

var ifres2277 Obj

if True == tmp2293 {
tmp2287 := GTl(__e, V7016)

tmp2288 := GHd(__e, tmp2287)

tmp2289 := PrimGuard1(__e, symstring_2, PrimIsString, tmp2288)

var ifres2279 Obj

if True == tmp2289 {
tmp2281 := GTl(__e, V7016)

tmp2282 := GTl(__e, tmp2281)

tmp2283 := GHd(__e, tmp2282)

tmp2284 := GTl(__e, tmp2283)

tmp2285 := GHd(__e, tmp2284)

tmp2286 := PrimGuard1(__e, symstring_2, PrimIsString, tmp2285)

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
tmp2244 := GTl(__e, V7016)

tmp2245 := GHd(__e, tmp2244)

tmp2246 := GTl(__e, V7016)

tmp2247 := GTl(__e, tmp2246)

tmp2248 := GHd(__e, tmp2247)

tmp2249 := GTl(__e, tmp2248)

tmp2250 := GHd(__e, tmp2249)

tmp2251 := PrimGuard2(__e, symcn, PrimStringConcat, tmp2245, tmp2250)

tmp2252 := GTl(__e, V7016)

tmp2253 := GTl(__e, tmp2252)

tmp2254 := GHd(__e, tmp2253)

tmp2255 := GTl(__e, tmp2254)

tmp2256 := GTl(__e, tmp2255)

tmp2257 := GCons(__e, tmp2251, tmp2256)

return GCons(__e, symcn, tmp2257)


} else {
return V7016
}


panic("unreachable")
})

tmp2329 := Call2(__e, ns2_1set, symshen_4factor_1cn, tmp2243)


_ = tmp2329

tmp2330 := MakeNative1(func(__e *ControlFlow, V7019 Obj) Obj {
tmp2356 := GEq(__e, MakeString(""), V7019)

if True == tmp2356 {
return MakeString("")
} else {
tmp2354 := Call1(__e, SymFn(symshen_4_7string_2), V7019)


var ifres2341 Obj

if True == tmp2354 {
tmp2352 := Call1(__e, SymFn(symhdstr), V7019)


tmp2353 := GEq(__e, MakeString("~"), tmp2352)

var ifres2343 Obj

if True == tmp2353 {
tmp2351 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, V7019))


var ifres2345 Obj

if True == tmp2351 {
tmp2348 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, V7019))


tmp2349 := GEq(__e, MakeString("%"), tmp2348)

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
tmp2331 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, MakeInteger(10))

tmp2334 := Call1(__e, SymFn(symshen_4proc_1nl), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7019)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V7019)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})())


return PrimGuard2(__e, symcn, PrimStringConcat, tmp2331, tmp2334)


} else {
tmp2339 := Call1(__e, SymFn(symshen_4_7string_2), V7019)


if True == tmp2339 {
tmp2335 := Call1(__e, SymFn(symhdstr), V7019)


tmp2337 := Call1(__e, SymFn(symshen_4proc_1nl), PrimGuard1(__e, symtlstr, PrimTailString, V7019))


return PrimGuard2(__e, symcn, PrimStringConcat, tmp2335, tmp2337)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.proc-nl"))
}


}


}


panic("unreachable")
})

tmp2357 := Call2(__e, ns2_1set, symshen_4proc_1nl, tmp2330)


_ = tmp2357

tmp2358 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V7024 := __self1
_ = V7024
V7025 := __self2
_ = V7025
tmp2367 := GEq(__e, Nil, V7025)

if True == tmp2367 {
return V7024
} else {
tmp2365 := GConsP(__e, V7025)

if True == tmp2365 {
tmp2359 := GHd(__e, V7025)

tmp2360 := GCons(__e, V7024, Nil)

tmp2361 := GCons(__e, tmp2359, tmp2360)

tmp2362 := GCons(__e, symshen_4insert, tmp2361)

tmp2363 := GTl(__e, V7025)

if SymFn(symshen_4mkstr_1r) == __self {
__self1, __self2 = tmp2362, tmp2363
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4mkstr_1r), tmp2362, tmp2363)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.mkstr-r"))
}


}


panic("unreachable")
})
return __self
})()

tmp2368 := Call2(__e, ns2_1set, symshen_4mkstr_1r, tmp2358)


_ = tmp2368

tmp2369 := MakeNative2(func(__e *ControlFlow, V7026, V7027 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4insert_1h), V7026, V7027, MakeString(""))
panic("unreachable")
})

tmp2370 := Call2(__e, ns2_1set, symshen_4insert, tmp2369)


_ = tmp2370

tmp2371 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V7036 := __self1
_ = V7036
V7037 := __self2
_ = V7037
V7038 := __self3
_ = V7038
tmp2432 := GEq(__e, MakeString(""), V7037)

if True == tmp2432 {
return V7038
} else {
tmp2430 := Call1(__e, SymFn(symshen_4_7string_2), V7037)


var ifres2417 Obj

if True == tmp2430 {
tmp2428 := Call1(__e, SymFn(symhdstr), V7037)


tmp2429 := GEq(__e, MakeString("~"), tmp2428)

var ifres2419 Obj

if True == tmp2429 {
tmp2427 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, V7037))


var ifres2421 Obj

if True == tmp2427 {
tmp2424 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, V7037))


tmp2425 := GEq(__e, MakeString("A"), tmp2424)

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
tmp2374 := Call3(__e, SymFn(symshen_4app), V7036, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7037)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V7037)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})(), symshen_4a)


return PrimGuard2(__e, symcn, PrimStringConcat, V7038, tmp2374)


} else {
tmp2415 := Call1(__e, SymFn(symshen_4_7string_2), V7037)


var ifres2402 Obj

if True == tmp2415 {
tmp2413 := Call1(__e, SymFn(symhdstr), V7037)


tmp2414 := GEq(__e, MakeString("~"), tmp2413)

var ifres2404 Obj

if True == tmp2414 {
tmp2412 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, V7037))


var ifres2406 Obj

if True == tmp2412 {
tmp2409 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, V7037))


tmp2410 := GEq(__e, MakeString("R"), tmp2409)

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
tmp2377 := Call3(__e, SymFn(symshen_4app), V7036, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7037)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V7037)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})(), symshen_4r)


return PrimGuard2(__e, symcn, PrimStringConcat, V7038, tmp2377)


} else {
tmp2400 := Call1(__e, SymFn(symshen_4_7string_2), V7037)


var ifres2387 Obj

if True == tmp2400 {
tmp2398 := Call1(__e, SymFn(symhdstr), V7037)


tmp2399 := GEq(__e, MakeString("~"), tmp2398)

var ifres2389 Obj

if True == tmp2399 {
tmp2397 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, V7037))


var ifres2391 Obj

if True == tmp2397 {
tmp2394 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, V7037))


tmp2395 := GEq(__e, MakeString("S"), tmp2394)

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
tmp2380 := Call3(__e, SymFn(symshen_4app), V7036, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7037)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V7037)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})(), symshen_4s)


return PrimGuard2(__e, symcn, PrimStringConcat, V7038, tmp2380)


} else {
tmp2385 := Call1(__e, SymFn(symshen_4_7string_2), V7037)


if True == tmp2385 {
tmp2381 := PrimGuard1(__e, symtlstr, PrimTailString, V7037)

tmp2382 := Call1(__e, SymFn(symhdstr), V7037)


if SymFn(symshen_4insert_1h) == __self {
__self1, __self2, __self3 = V7036, tmp2381, PrimGuard2(__e, symcn, PrimStringConcat, V7038, tmp2382)
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4insert_1h), V7036, tmp2381, PrimGuard2(__e, symcn, PrimStringConcat, V7038, tmp2382))


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

tmp2433 := Call2(__e, ns2_1set, symshen_4insert_1h, tmp2371)


_ = tmp2433

tmp2434 := MakeNative3(func(__e *ControlFlow, V7039, V7040, V7041 Obj) Obj {
tmp2435 := Call2(__e, SymFn(symshen_4arg_1_6str), V7039, V7041)


return PrimGuard2(__e, symcn, PrimStringConcat, tmp2435, V7040)


panic("unreachable")
})

tmp2436 := Call2(__e, ns2_1set, symshen_4app, tmp2434)


_ = tmp2436

tmp2437 := MakeNative2(func(__e *ControlFlow, V7045, V7046 Obj) Obj {
tmp2445 := Call0(__e, SymFn(symfail))


tmp2446 := GEq(__e, V7045, tmp2445)

if True == tmp2446 {
return MakeString("...")
} else {
tmp2443 := Call1(__e, SymFn(symshen_4list_2), V7045)


if True == tmp2443 {
return TailCall2(__e, SymFn(symshen_4list_1_6str), V7045, V7046)
} else {
tmp2441 := PrimGuard1(__e, symstring_2, PrimIsString, V7045)

if True == tmp2441 {
return TailCall2(__e, SymFn(symshen_4str_1_6str), V7045, V7046)
} else {
tmp2439 := GVectorP(__e, V7045)

if True == tmp2439 {
return TailCall2(__e, SymFn(symshen_4vector_1_6str), V7045, V7046)
} else {
return TailCall1(__e, SymFn(symshen_4atom_1_6str), V7045)
}


}


}


}


panic("unreachable")
})

tmp2447 := Call2(__e, ns2_1set, symshen_4arg_1_6str, tmp2437)


_ = tmp2447

tmp2448 := MakeNative2(func(__e *ControlFlow, V7047, V7048 Obj) Obj {
tmp2456 := GEq(__e, symshen_4r, V7048)

if True == tmp2456 {
tmp2449 := Call0(__e, SymFn(symshen_4maxseq))


tmp2450 := Call3(__e, SymFn(symshen_4iter_1list), V7047, symshen_4r, tmp2449)


tmp2451 := Call2(__e, SymFn(sym_8s), tmp2450, MakeString(")"))


return TailCall2(__e, SymFn(sym_8s), MakeString("("), tmp2451)


} else {
tmp2452 := Call0(__e, SymFn(symshen_4maxseq))


tmp2453 := Call3(__e, SymFn(symshen_4iter_1list), V7047, V7048, tmp2452)


tmp2454 := Call2(__e, SymFn(sym_8s), tmp2453, MakeString("]"))


return TailCall2(__e, SymFn(sym_8s), MakeString("["), tmp2454)


}


panic("unreachable")
})

tmp2457 := Call2(__e, ns2_1set, symshen_4list_1_6str, tmp2448)


_ = tmp2457

tmp2458 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dmaximum_1print_1sequence_1size_d)
panic("unreachable")
})

tmp2459 := Call2(__e, ns2_1set, symshen_4maxseq, tmp2458)


_ = tmp2459

tmp2460 := MakeNative3(func(__e *ControlFlow, V7059, V7060, V7061 Obj) Obj {
tmp2481 := GEq(__e, Nil, V7059)

if True == tmp2481 {
return MakeString("")
} else {
tmp2479 := GEq(__e, MakeInteger(0), V7061)

if True == tmp2479 {
return MakeString("... etc")
} else {
tmp2477 := GConsP(__e, V7059)

var ifres2473 Obj

if True == tmp2477 {
tmp2475 := GTl(__e, V7059)

tmp2476 := GEq(__e, Nil, tmp2475)

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
tmp2461 := GHd(__e, V7059)

return TailCall2(__e, SymFn(symshen_4arg_1_6str), tmp2461, V7060)


} else {
tmp2471 := GConsP(__e, V7059)

if True == tmp2471 {
tmp2462 := GHd(__e, V7059)

tmp2463 := Call2(__e, SymFn(symshen_4arg_1_6str), tmp2462, V7060)


tmp2464 := GTl(__e, V7059)

tmp2466 := Call3(__e, SymFn(symshen_4iter_1list), tmp2464, V7060, GSub(__e, V7061, MakeInteger(1)))


tmp2467 := Call2(__e, SymFn(sym_8s), MakeString(" "), tmp2466)


return TailCall2(__e, SymFn(sym_8s), tmp2463, tmp2467)


} else {
tmp2468 := Call2(__e, SymFn(symshen_4arg_1_6str), V7059, V7060)


tmp2469 := Call2(__e, SymFn(sym_8s), MakeString(" "), tmp2468)


return TailCall2(__e, SymFn(sym_8s), MakeString("|"), tmp2469)


}


}


}


}


panic("unreachable")
})

tmp2482 := Call2(__e, ns2_1set, symshen_4iter_1list, tmp2460)


_ = tmp2482

tmp2483 := MakeNative2(func(__e *ControlFlow, V7064, V7065 Obj) Obj {
tmp2488 := GEq(__e, symshen_4a, V7065)

if True == tmp2488 {
return V7064
} else {
tmp2484 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, MakeInteger(34))

tmp2485 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, MakeInteger(34))

tmp2486 := Call2(__e, SymFn(sym_8s), V7064, tmp2485)


return TailCall2(__e, SymFn(sym_8s), tmp2484, tmp2486)


}


panic("unreachable")
})

tmp2489 := Call2(__e, ns2_1set, symshen_4str_1_6str, tmp2483)


_ = tmp2489

tmp2490 := MakeNative2(func(__e *ControlFlow, V7066, V7067 Obj) Obj {
tmp2503 := Call1(__e, SymFn(symshen_4print_1vector_2), V7066)


if True == tmp2503 {
tmp2491 := GVecGet(__e, V7066, MakeInteger(0))

tmp2492 := Call1(__e, SymFn(symfn), tmp2491)


return TailCall1(__e, tmp2492, V7066)


} else {
tmp2501 := Call1(__e, SymFn(symvector_2), V7066)


if True == tmp2501 {
tmp2493 := Call0(__e, SymFn(symshen_4maxseq))


tmp2494 := Call4(__e, SymFn(symshen_4iter_1vector), V7066, MakeInteger(1), V7067, tmp2493)


tmp2495 := Call2(__e, SymFn(sym_8s), tmp2494, MakeString(">"))


return TailCall2(__e, SymFn(sym_8s), MakeString("<"), tmp2495)


} else {
tmp2496 := Call0(__e, SymFn(symshen_4maxseq))


tmp2497 := Call4(__e, SymFn(symshen_4iter_1vector), V7066, MakeInteger(0), V7067, tmp2496)


tmp2498 := Call2(__e, SymFn(sym_8s), tmp2497, MakeString(">>"))


tmp2499 := Call2(__e, SymFn(sym_8s), MakeString("<"), tmp2498)


return TailCall2(__e, SymFn(sym_8s), MakeString("<"), tmp2499)


}


}


panic("unreachable")
})

tmp2504 := Call2(__e, ns2_1set, symshen_4vector_1_6str, tmp2490)


_ = tmp2504

tmp2505 := MakeNative1(func(__e *ControlFlow, V7068 Obj) Obj {
tmp2506 := GValue(__e, symshen_4_dempty_1absvector_d)

return GEq(__e, V7068, tmp2506)


panic("unreachable")
})

tmp2507 := Call2(__e, ns2_1set, symshen_4empty_1absvector_2, tmp2505)


_ = tmp2507

tmp2508 := MakeNative1(func(__e *ControlFlow, V7069 Obj) Obj {
tmp2526 := Call1(__e, SymFn(symshen_4empty_1absvector_2), V7069)


if True == GNot(__e, tmp2526) {
tmp2510 := GVecGet(__e, V7069, MakeInteger(0))

let__2060 := tmp2510
_ = let__2060

tmp2524 := GEq(__e, let__2060, symshen_4tuple)

var ifres2511 Obj

if True == tmp2524 {
ifres2511 = True


} else {
tmp2523 := GEq(__e, let__2060, symshen_4pvar)

var ifres2513 Obj

if True == tmp2523 {
ifres2513 = True


} else {
tmp2522 := GEq(__e, let__2060, symshen_4dictionary)

var ifres2515 Obj

if True == tmp2522 {
ifres2515 = True


} else {
tmp2520 := PrimGuard1(__e, symnumber_2, PrimIsNumber, let__2060)

tmp2521 := GNot(__e, tmp2520)

var ifres2517 Obj

if True == tmp2521 {
tmp2519 := Call1(__e, SymFn(symshen_4fbound_2), let__2060)


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
return True
} else {
return False
}


} else {
return False
}


panic("unreachable")
})

tmp2528 := Call2(__e, ns2_1set, symshen_4print_1vector_2, tmp2508)


_ = tmp2528

tmp2529 := MakeNative1(func(__e *ControlFlow, V7071 Obj) Obj {
tmp2530 := Call1(__e, SymFn(symarity), V7071)


tmp2531 := GEq(__e, tmp2530, MakeInteger(-1))

return GNot(__e, tmp2531)


panic("unreachable")
})

tmp2532 := Call2(__e, ns2_1set, symshen_4fbound_2, tmp2529)


_ = tmp2532

tmp2533 := MakeNative1(func(__e *ControlFlow, V7072 Obj) Obj {
tmp2534 := GVecGet(__e, V7072, MakeInteger(1))

tmp2535 := GVecGet(__e, V7072, MakeInteger(2))

tmp2536 := Call3(__e, SymFn(symshen_4app), tmp2535, MakeString(")"), symshen_4s)


tmp2538 := Call3(__e, SymFn(symshen_4app), tmp2534, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" "), tmp2536), symshen_4s)


return PrimGuard2(__e, symcn, PrimStringConcat, MakeString("(@p "), tmp2538)


panic("unreachable")
})

tmp2539 := Call2(__e, ns2_1set, symshen_4tuple, tmp2533)


_ = tmp2539

tmp2540 := MakeNative1(func(__e *ControlFlow, V7073 Obj) Obj {
return MakeString("(dict ...)")
panic("unreachable")
})

tmp2541 := Call2(__e, ns2_1set, symshen_4dictionary, tmp2540)


_ = tmp2541

tmp2542 := MakeNative4(func(__e *ControlFlow, V7080, V7081, V7082, V7083 Obj) Obj {
tmp2560 := GEq(__e, MakeInteger(0), V7083)

if True == tmp2560 {
return MakeString("... etc")
} else {
tmp2543 := MakeNative0(func(__e *ControlFlow) Obj {
return GVecGet(__e, V7080, V7081)
panic("unreachable")
})

tmp2544 := MakeNative1(func(__e *ControlFlow, Z7085 Obj) Obj {
return symshen_4out_1of_1bounds
panic("unreachable")
})

tmp2545 := Call2(__e, try_1catch, tmp2543, tmp2544)


let__2061 := tmp2545
_ = let__2061

tmp2546 := MakeNative0(func(__e *ControlFlow) Obj {
return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V7080, GAdd(__e, V7081, MakeInteger(1)))
}
__typedArg0 := V7080
__typedArg1 := GAdd(__e, V7081, MakeInteger(1))
return Call2(__e, SymFn(sym_5_1address), __typedArg0, __typedArg1)
})()


panic("unreachable")
})

tmp2548 := MakeNative1(func(__e *ControlFlow, Z7087 Obj) Obj {
return symshen_4out_1of_1bounds
panic("unreachable")
})

tmp2549 := Call2(__e, try_1catch, tmp2546, tmp2548)


let__2062 := tmp2549
_ = let__2062

tmp2558 := GEq(__e, let__2061, symshen_4out_1of_1bounds)

if True == tmp2558 {
return MakeString("")
} else {
tmp2556 := GEq(__e, let__2062, symshen_4out_1of_1bounds)

if True == tmp2556 {
return TailCall2(__e, SymFn(symshen_4arg_1_6str), let__2061, V7082)
} else {
tmp2550 := Call2(__e, SymFn(symshen_4arg_1_6str), let__2061, V7082)


tmp2551 := GAdd(__e, V7081, MakeInteger(1))

tmp2553 := Call4(__e, SymFn(symshen_4iter_1vector), V7080, tmp2551, V7082, GSub(__e, V7083, MakeInteger(1)))


tmp2554 := Call2(__e, SymFn(sym_8s), MakeString(" "), tmp2553)


return TailCall2(__e, SymFn(sym_8s), tmp2550, tmp2554)


}


}


}


panic("unreachable")
})

tmp2561 := Call2(__e, ns2_1set, symshen_4iter_1vector, tmp2542)


_ = tmp2561

tmp2562 := MakeNative1(func(__e *ControlFlow, V7088 Obj) Obj {
tmp2563 := MakeNative0(func(__e *ControlFlow) Obj {
return PrimGuard1(__e, symstr, PrimStr, V7088)
panic("unreachable")
})

tmp2564 := MakeNative1(func(__e *ControlFlow, Z7089 Obj) Obj {
return TailCall0(__e, SymFn(symshen_4funexstring))
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp2563, tmp2564)


panic("unreachable")
})

tmp2565 := Call2(__e, ns2_1set, symshen_4atom_1_6str, tmp2562)


_ = tmp2565

tmp2566 := MakeNative0(func(__e *ControlFlow) Obj {
tmp2567 := PrimGuard1(__e, symintern, PrimIntern, MakeString("x"))

tmp2568 := Call1(__e, SymFn(symgensym), tmp2567)


tmp2569 := Call2(__e, SymFn(symshen_4arg_1_6str), tmp2568, symshen_4a)


tmp2570 := Call2(__e, SymFn(sym_8s), tmp2569, MakeString("\x11"))


tmp2571 := Call2(__e, SymFn(sym_8s), MakeString("e"), tmp2570)


tmp2572 := Call2(__e, SymFn(sym_8s), MakeString("n"), tmp2571)


tmp2573 := Call2(__e, SymFn(sym_8s), MakeString("u"), tmp2572)


tmp2574 := Call2(__e, SymFn(sym_8s), MakeString("f"), tmp2573)


return TailCall2(__e, SymFn(sym_8s), MakeString("\x10"), tmp2574)


panic("unreachable")
})

tmp2575 := Call2(__e, ns2_1set, symshen_4funexstring, tmp2566)


_ = tmp2575

tmp2576 := MakeNative1(func(__e *ControlFlow, V7090 Obj) Obj {
tmp2580 := Call1(__e, SymFn(symempty_2), V7090)


if True == tmp2580 {
return True
} else {
tmp2578 := GConsP(__e, V7090)

if True == tmp2578 {
return True
} else {
return False
}


}


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4list_2, tmp2576)




panic("unreachable")
})

