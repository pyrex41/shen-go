package main

import . "github.com/pyrex41/shen-go/kl"

var WriterMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp2026 := MakeNative1(func(__e *ControlFlow, V6988 Obj) Obj {
tmp2027 := Call2(__e, SymFn(symshen_4insert), V6988, MakeString("~S"))


let__2021 := tmp2027
_ = let__2021

tmp2028 := Call0(__e, SymFn(symstoutput))


tmp2029 := Call2(__e, SymFn(sympr), let__2021, tmp2028)


let__2022 := tmp2029
_ = let__2022

return V6988


panic("unreachable")
})

tmp2030 := Call2(__e, ns2_1set, symprint, tmp2026)


_ = tmp2030

tmp2031 := MakeNative2(func(__e *ControlFlow, V6991, V6992 Obj) Obj {
tmp2036 := GValue(__e, sym_dhush_d)

if True == tmp2036 {
return V6991
} else {
tmp2034 := Call1(__e, SymFn(symchar_1stream_2), V6992)


if True == tmp2034 {
return TailCall2(__e, SymFn(symwrite_1string), V6991, V6992)
} else {
tmp2032 := Call2(__e, SymFn(symshen_4string_1_6byte), V6991, MakeInteger(0))


return TailCall4(__e, SymFn(symshen_4write_1chars), V6991, V6992, tmp2032, MakeInteger(1))


}


}


panic("unreachable")
})

tmp2037 := Call2(__e, ns2_1set, sympr, tmp2031)


_ = tmp2037

tmp2038 := MakeNative2(func(__e *ControlFlow, V6993, V6994 Obj) Obj {
tmp2039 := MakeNative0(func(__e *ControlFlow) Obj {
tmp2040 := PrimGuard2(__e, sympos, PrimPos, V6993, V6994)

return PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp2040)


panic("unreachable")
})

tmp2041 := MakeNative1(func(__e *ControlFlow, Z6995 Obj) Obj {
return symshen_4eos
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp2039, tmp2041)


panic("unreachable")
})

tmp2042 := Call2(__e, ns2_1set, symshen_4string_1_6byte, tmp2038)


_ = tmp2042

tmp2043 := (func() Obj {
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
tmp2048 := GEq(__e, symshen_4eos, V6998)

if True == tmp2048 {
return V6996
} else {
tmp2044 := PrimGuard2(__e, symwrite_1byte, PrimWriteByte, V6998, V6997)

_ = tmp2044

tmp2045 := Call2(__e, SymFn(symshen_4string_1_6byte), V6996, V6999)


if SymFn(symshen_4write_1chars) == __self {
__self1, __self2, __self3, __self4 = V6996, V6997, tmp2045, GAdd(__e, V6999, MakeInteger(1))
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4write_1chars), V6996, V6997, tmp2045, GAdd(__e, V6999, MakeInteger(1)))


}


panic("unreachable")
})
return __self
})()

tmp2049 := Call2(__e, ns2_1set, symshen_4write_1chars, tmp2043)


_ = tmp2049

tmp2050 := MakeNative2(func(__e *ControlFlow, V7000, V7001 Obj) Obj {
tmp2055 := PrimGuard1(__e, symstring_2, PrimIsString, V7000)

if True == tmp2055 {
tmp2051 := Call1(__e, SymFn(symshen_4proc_1nl), V7000)


return TailCall2(__e, SymFn(symshen_4mkstr_1l), tmp2051, V7001)


} else {
tmp2052 := GCons(__e, V7000, Nil)

tmp2053 := GCons(__e, symshen_4proc_1nl, tmp2052)

return TailCall2(__e, SymFn(symshen_4mkstr_1r), tmp2053, V7001)


}


panic("unreachable")
})

tmp2056 := Call2(__e, ns2_1set, symshen_4mkstr, tmp2050)


_ = tmp2056

tmp2057 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V7006 := __self1
_ = V7006
V7007 := __self2
_ = V7007
tmp2064 := GEq(__e, Nil, V7007)

if True == tmp2064 {
return V7006
} else {
tmp2062 := GConsP(__e, V7007)

if True == tmp2062 {
tmp2058 := GHd(__e, V7007)

tmp2059 := Call2(__e, SymFn(symshen_4insert_1l), tmp2058, V7006)


tmp2060 := GTl(__e, V7007)

if SymFn(symshen_4mkstr_1l) == __self {
__self1, __self2 = tmp2059, tmp2060
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4mkstr_1l), tmp2059, tmp2060)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.mkstr-l"))
}


}


panic("unreachable")
})
return __self
})()

tmp2065 := Call2(__e, ns2_1set, symshen_4mkstr_1l, tmp2057)


_ = tmp2065

tmp2066 := MakeNative2(func(__e *ControlFlow, V7014, V7015 Obj) Obj {
tmp2186 := GEq(__e, MakeString(""), V7015)

if True == tmp2186 {
return MakeString("")
} else {
tmp2184 := Call1(__e, SymFn(symshen_4_7string_2), V7015)


var ifres2174 Obj

if True == tmp2184 {
tmp2182 := Call1(__e, SymFn(symhdstr), V7015)


tmp2183 := GEq(__e, MakeString("~"), tmp2182)

var ifres2175 Obj

if True == tmp2183 {
tmp2181 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, V7015))


var ifres2176 Obj

if True == tmp2181 {
tmp2178 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, V7015))


tmp2179 := GEq(__e, MakeString("A"), tmp2178)

ifres2176 = tmp2179


} else {
ifres2176 = False


}

ifres2175 = ifres2176


} else {
ifres2175 = False


}

ifres2174 = ifres2175


} else {
ifres2174 = False


}

if True == ifres2174 {
tmp2068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7015)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V7015)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})()

tmp2069 := GCons(__e, symshen_4a, Nil)

tmp2070 := GCons(__e, tmp2068, tmp2069)

tmp2071 := GCons(__e, V7014, tmp2070)

return GCons(__e, symshen_4app, tmp2071)


} else {
tmp2172 := Call1(__e, SymFn(symshen_4_7string_2), V7015)


var ifres2162 Obj

if True == tmp2172 {
tmp2170 := Call1(__e, SymFn(symhdstr), V7015)


tmp2171 := GEq(__e, MakeString("~"), tmp2170)

var ifres2163 Obj

if True == tmp2171 {
tmp2169 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, V7015))


var ifres2164 Obj

if True == tmp2169 {
tmp2166 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, V7015))


tmp2167 := GEq(__e, MakeString("R"), tmp2166)

ifres2164 = tmp2167


} else {
ifres2164 = False


}

ifres2163 = ifres2164


} else {
ifres2163 = False


}

ifres2162 = ifres2163


} else {
ifres2162 = False


}

if True == ifres2162 {
tmp2073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7015)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V7015)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})()

tmp2074 := GCons(__e, symshen_4r, Nil)

tmp2075 := GCons(__e, tmp2073, tmp2074)

tmp2076 := GCons(__e, V7014, tmp2075)

return GCons(__e, symshen_4app, tmp2076)


} else {
tmp2160 := Call1(__e, SymFn(symshen_4_7string_2), V7015)


var ifres2150 Obj

if True == tmp2160 {
tmp2158 := Call1(__e, SymFn(symhdstr), V7015)


tmp2159 := GEq(__e, MakeString("~"), tmp2158)

var ifres2151 Obj

if True == tmp2159 {
tmp2157 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, V7015))


var ifres2152 Obj

if True == tmp2157 {
tmp2154 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, V7015))


tmp2155 := GEq(__e, MakeString("S"), tmp2154)

ifres2152 = tmp2155


} else {
ifres2152 = False


}

ifres2151 = ifres2152


} else {
ifres2151 = False


}

ifres2150 = ifres2151


} else {
ifres2150 = False


}

if True == ifres2150 {
tmp2078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7015)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V7015)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})()

tmp2079 := GCons(__e, symshen_4s, Nil)

tmp2080 := GCons(__e, tmp2078, tmp2079)

tmp2081 := GCons(__e, V7014, tmp2080)

return GCons(__e, symshen_4app, tmp2081)


} else {
tmp2148 := Call1(__e, SymFn(symshen_4_7string_2), V7015)


if True == tmp2148 {
tmp2082 := Call1(__e, SymFn(symhdstr), V7015)


tmp2084 := Call2(__e, SymFn(symshen_4insert_1l), V7014, PrimGuard1(__e, symtlstr, PrimTailString, V7015))


tmp2085 := GCons(__e, tmp2084, Nil)

tmp2086 := GCons(__e, tmp2082, tmp2085)

tmp2087 := GCons(__e, symcn, tmp2086)

return TailCall1(__e, SymFn(symshen_4factor_1cn), tmp2087)


} else {
tmp2146 := GConsP(__e, V7015)

var ifres2131 Obj

if True == tmp2146 {
tmp2144 := GHd(__e, V7015)

tmp2145 := GEq(__e, symcn, tmp2144)

var ifres2132 Obj

if True == tmp2145 {
tmp2142 := GTl(__e, V7015)

tmp2143 := GConsP(__e, tmp2142)

var ifres2133 Obj

if True == tmp2143 {
tmp2139 := GTl(__e, V7015)

tmp2140 := GTl(__e, tmp2139)

tmp2141 := GConsP(__e, tmp2140)

var ifres2134 Obj

if True == tmp2141 {
tmp2135 := GTl(__e, V7015)

tmp2136 := GTl(__e, tmp2135)

tmp2137 := GTl(__e, tmp2136)

tmp2138 := GEq(__e, Nil, tmp2137)

ifres2134 = tmp2138


} else {
ifres2134 = False


}

ifres2133 = ifres2134


} else {
ifres2133 = False


}

ifres2132 = ifres2133


} else {
ifres2132 = False


}

ifres2131 = ifres2132


} else {
ifres2131 = False


}

if True == ifres2131 {
tmp2088 := GTl(__e, V7015)

tmp2089 := GHd(__e, tmp2088)

tmp2090 := GTl(__e, V7015)

tmp2091 := GTl(__e, tmp2090)

tmp2092 := GHd(__e, tmp2091)

tmp2093 := Call2(__e, SymFn(symshen_4insert_1l), V7014, tmp2092)


tmp2094 := GCons(__e, tmp2093, Nil)

tmp2095 := GCons(__e, tmp2089, tmp2094)

return GCons(__e, symcn, tmp2095)


} else {
tmp2129 := GConsP(__e, V7015)

var ifres2108 Obj

if True == tmp2129 {
tmp2127 := GHd(__e, V7015)

tmp2128 := GEq(__e, symshen_4app, tmp2127)

var ifres2109 Obj

if True == tmp2128 {
tmp2125 := GTl(__e, V7015)

tmp2126 := GConsP(__e, tmp2125)

var ifres2110 Obj

if True == tmp2126 {
tmp2122 := GTl(__e, V7015)

tmp2123 := GTl(__e, tmp2122)

tmp2124 := GConsP(__e, tmp2123)

var ifres2111 Obj

if True == tmp2124 {
tmp2118 := GTl(__e, V7015)

tmp2119 := GTl(__e, tmp2118)

tmp2120 := GTl(__e, tmp2119)

tmp2121 := GConsP(__e, tmp2120)

var ifres2112 Obj

if True == tmp2121 {
tmp2113 := GTl(__e, V7015)

tmp2114 := GTl(__e, tmp2113)

tmp2115 := GTl(__e, tmp2114)

tmp2116 := GTl(__e, tmp2115)

tmp2117 := GEq(__e, Nil, tmp2116)

ifres2112 = tmp2117


} else {
ifres2112 = False


}

ifres2111 = ifres2112


} else {
ifres2111 = False


}

ifres2110 = ifres2111


} else {
ifres2110 = False


}

ifres2109 = ifres2110


} else {
ifres2109 = False


}

ifres2108 = ifres2109


} else {
ifres2108 = False


}

if True == ifres2108 {
tmp2096 := GTl(__e, V7015)

tmp2097 := GHd(__e, tmp2096)

tmp2098 := GTl(__e, V7015)

tmp2099 := GTl(__e, tmp2098)

tmp2100 := GHd(__e, tmp2099)

tmp2101 := Call2(__e, SymFn(symshen_4insert_1l), V7014, tmp2100)


tmp2102 := GTl(__e, V7015)

tmp2103 := GTl(__e, tmp2102)

tmp2104 := GTl(__e, tmp2103)

tmp2105 := GCons(__e, tmp2101, tmp2104)

tmp2106 := GCons(__e, tmp2097, tmp2105)

return GCons(__e, symshen_4app, tmp2106)


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

tmp2187 := Call2(__e, ns2_1set, symshen_4insert_1l, tmp2066)


_ = tmp2187

tmp2188 := MakeNative1(func(__e *ControlFlow, V7016 Obj) Obj {
tmp2262 := GConsP(__e, V7016)

var ifres2204 Obj

if True == tmp2262 {
tmp2260 := GHd(__e, V7016)

tmp2261 := GEq(__e, symcn, tmp2260)

var ifres2205 Obj

if True == tmp2261 {
tmp2258 := GTl(__e, V7016)

tmp2259 := GConsP(__e, tmp2258)

var ifres2206 Obj

if True == tmp2259 {
tmp2255 := GTl(__e, V7016)

tmp2256 := GTl(__e, tmp2255)

tmp2257 := GConsP(__e, tmp2256)

var ifres2207 Obj

if True == tmp2257 {
tmp2251 := GTl(__e, V7016)

tmp2252 := GTl(__e, tmp2251)

tmp2253 := GHd(__e, tmp2252)

tmp2254 := GConsP(__e, tmp2253)

var ifres2208 Obj

if True == tmp2254 {
tmp2246 := GTl(__e, V7016)

tmp2247 := GTl(__e, tmp2246)

tmp2248 := GHd(__e, tmp2247)

tmp2249 := GHd(__e, tmp2248)

tmp2250 := GEq(__e, symcn, tmp2249)

var ifres2209 Obj

if True == tmp2250 {
tmp2241 := GTl(__e, V7016)

tmp2242 := GTl(__e, tmp2241)

tmp2243 := GHd(__e, tmp2242)

tmp2244 := GTl(__e, tmp2243)

tmp2245 := GConsP(__e, tmp2244)

var ifres2210 Obj

if True == tmp2245 {
tmp2235 := GTl(__e, V7016)

tmp2236 := GTl(__e, tmp2235)

tmp2237 := GHd(__e, tmp2236)

tmp2238 := GTl(__e, tmp2237)

tmp2239 := GTl(__e, tmp2238)

tmp2240 := GConsP(__e, tmp2239)

var ifres2211 Obj

if True == tmp2240 {
tmp2228 := GTl(__e, V7016)

tmp2229 := GTl(__e, tmp2228)

tmp2230 := GHd(__e, tmp2229)

tmp2231 := GTl(__e, tmp2230)

tmp2232 := GTl(__e, tmp2231)

tmp2233 := GTl(__e, tmp2232)

tmp2234 := GEq(__e, Nil, tmp2233)

var ifres2212 Obj

if True == tmp2234 {
tmp2224 := GTl(__e, V7016)

tmp2225 := GTl(__e, tmp2224)

tmp2226 := GTl(__e, tmp2225)

tmp2227 := GEq(__e, Nil, tmp2226)

var ifres2213 Obj

if True == tmp2227 {
tmp2221 := GTl(__e, V7016)

tmp2222 := GHd(__e, tmp2221)

tmp2223 := PrimGuard1(__e, symstring_2, PrimIsString, tmp2222)

var ifres2214 Obj

if True == tmp2223 {
tmp2215 := GTl(__e, V7016)

tmp2216 := GTl(__e, tmp2215)

tmp2217 := GHd(__e, tmp2216)

tmp2218 := GTl(__e, tmp2217)

tmp2219 := GHd(__e, tmp2218)

tmp2220 := PrimGuard1(__e, symstring_2, PrimIsString, tmp2219)

ifres2214 = tmp2220


} else {
ifres2214 = False


}

ifres2213 = ifres2214


} else {
ifres2213 = False


}

ifres2212 = ifres2213


} else {
ifres2212 = False


}

ifres2211 = ifres2212


} else {
ifres2211 = False


}

ifres2210 = ifres2211


} else {
ifres2210 = False


}

ifres2209 = ifres2210


} else {
ifres2209 = False


}

ifres2208 = ifres2209


} else {
ifres2208 = False


}

ifres2207 = ifres2208


} else {
ifres2207 = False


}

ifres2206 = ifres2207


} else {
ifres2206 = False


}

ifres2205 = ifres2206


} else {
ifres2205 = False


}

ifres2204 = ifres2205


} else {
ifres2204 = False


}

if True == ifres2204 {
tmp2189 := GTl(__e, V7016)

tmp2190 := GHd(__e, tmp2189)

tmp2191 := GTl(__e, V7016)

tmp2192 := GTl(__e, tmp2191)

tmp2193 := GHd(__e, tmp2192)

tmp2194 := GTl(__e, tmp2193)

tmp2195 := GHd(__e, tmp2194)

tmp2196 := PrimGuard2(__e, symcn, PrimStringConcat, tmp2190, tmp2195)

tmp2197 := GTl(__e, V7016)

tmp2198 := GTl(__e, tmp2197)

tmp2199 := GHd(__e, tmp2198)

tmp2200 := GTl(__e, tmp2199)

tmp2201 := GTl(__e, tmp2200)

tmp2202 := GCons(__e, tmp2196, tmp2201)

return GCons(__e, symcn, tmp2202)


} else {
return V7016
}


panic("unreachable")
})

tmp2263 := Call2(__e, ns2_1set, symshen_4factor_1cn, tmp2188)


_ = tmp2263

tmp2264 := MakeNative1(func(__e *ControlFlow, V7019 Obj) Obj {
tmp2287 := GEq(__e, MakeString(""), V7019)

if True == tmp2287 {
return MakeString("")
} else {
tmp2285 := Call1(__e, SymFn(symshen_4_7string_2), V7019)


var ifres2275 Obj

if True == tmp2285 {
tmp2283 := Call1(__e, SymFn(symhdstr), V7019)


tmp2284 := GEq(__e, MakeString("~"), tmp2283)

var ifres2276 Obj

if True == tmp2284 {
tmp2282 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, V7019))


var ifres2277 Obj

if True == tmp2282 {
tmp2279 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, V7019))


tmp2280 := GEq(__e, MakeString("%"), tmp2279)

ifres2277 = tmp2280


} else {
ifres2277 = False


}

ifres2276 = ifres2277


} else {
ifres2276 = False


}

ifres2275 = ifres2276


} else {
ifres2275 = False


}

if True == ifres2275 {
tmp2265 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, MakeInteger(10))

tmp2268 := Call1(__e, SymFn(symshen_4proc_1nl), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7019)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V7019)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})())


return PrimGuard2(__e, symcn, PrimStringConcat, tmp2265, tmp2268)


} else {
tmp2273 := Call1(__e, SymFn(symshen_4_7string_2), V7019)


if True == tmp2273 {
tmp2269 := Call1(__e, SymFn(symhdstr), V7019)


tmp2271 := Call1(__e, SymFn(symshen_4proc_1nl), PrimGuard1(__e, symtlstr, PrimTailString, V7019))


return PrimGuard2(__e, symcn, PrimStringConcat, tmp2269, tmp2271)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.proc-nl"))
}


}


}


panic("unreachable")
})

tmp2288 := Call2(__e, ns2_1set, symshen_4proc_1nl, tmp2264)


_ = tmp2288

tmp2289 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V7024 := __self1
_ = V7024
V7025 := __self2
_ = V7025
tmp2298 := GEq(__e, Nil, V7025)

if True == tmp2298 {
return V7024
} else {
tmp2296 := GConsP(__e, V7025)

if True == tmp2296 {
tmp2290 := GHd(__e, V7025)

tmp2291 := GCons(__e, V7024, Nil)

tmp2292 := GCons(__e, tmp2290, tmp2291)

tmp2293 := GCons(__e, symshen_4insert, tmp2292)

tmp2294 := GTl(__e, V7025)

if SymFn(symshen_4mkstr_1r) == __self {
__self1, __self2 = tmp2293, tmp2294
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4mkstr_1r), tmp2293, tmp2294)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.mkstr-r"))
}


}


panic("unreachable")
})
return __self
})()

tmp2299 := Call2(__e, ns2_1set, symshen_4mkstr_1r, tmp2289)


_ = tmp2299

tmp2300 := MakeNative2(func(__e *ControlFlow, V7026, V7027 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4insert_1h), V7026, V7027, MakeString(""))
panic("unreachable")
})

tmp2301 := Call2(__e, ns2_1set, symshen_4insert, tmp2300)


_ = tmp2301

tmp2302 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V7036 := __self1
_ = V7036
V7037 := __self2
_ = V7037
V7038 := __self3
_ = V7038
tmp2354 := GEq(__e, MakeString(""), V7037)

if True == tmp2354 {
return V7038
} else {
tmp2352 := Call1(__e, SymFn(symshen_4_7string_2), V7037)


var ifres2342 Obj

if True == tmp2352 {
tmp2350 := Call1(__e, SymFn(symhdstr), V7037)


tmp2351 := GEq(__e, MakeString("~"), tmp2350)

var ifres2343 Obj

if True == tmp2351 {
tmp2349 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, V7037))


var ifres2344 Obj

if True == tmp2349 {
tmp2346 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, V7037))


tmp2347 := GEq(__e, MakeString("A"), tmp2346)

ifres2344 = tmp2347


} else {
ifres2344 = False


}

ifres2343 = ifres2344


} else {
ifres2343 = False


}

ifres2342 = ifres2343


} else {
ifres2342 = False


}

if True == ifres2342 {
tmp2305 := Call3(__e, SymFn(symshen_4app), V7036, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7037)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V7037)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})(), symshen_4a)


return PrimGuard2(__e, symcn, PrimStringConcat, V7038, tmp2305)


} else {
tmp2340 := Call1(__e, SymFn(symshen_4_7string_2), V7037)


var ifres2330 Obj

if True == tmp2340 {
tmp2338 := Call1(__e, SymFn(symhdstr), V7037)


tmp2339 := GEq(__e, MakeString("~"), tmp2338)

var ifres2331 Obj

if True == tmp2339 {
tmp2337 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, V7037))


var ifres2332 Obj

if True == tmp2337 {
tmp2334 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, V7037))


tmp2335 := GEq(__e, MakeString("R"), tmp2334)

ifres2332 = tmp2335


} else {
ifres2332 = False


}

ifres2331 = ifres2332


} else {
ifres2331 = False


}

ifres2330 = ifres2331


} else {
ifres2330 = False


}

if True == ifres2330 {
tmp2308 := Call3(__e, SymFn(symshen_4app), V7036, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7037)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V7037)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})(), symshen_4r)


return PrimGuard2(__e, symcn, PrimStringConcat, V7038, tmp2308)


} else {
tmp2328 := Call1(__e, SymFn(symshen_4_7string_2), V7037)


var ifres2318 Obj

if True == tmp2328 {
tmp2326 := Call1(__e, SymFn(symhdstr), V7037)


tmp2327 := GEq(__e, MakeString("~"), tmp2326)

var ifres2319 Obj

if True == tmp2327 {
tmp2325 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, V7037))


var ifres2320 Obj

if True == tmp2325 {
tmp2322 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, V7037))


tmp2323 := GEq(__e, MakeString("S"), tmp2322)

ifres2320 = tmp2323


} else {
ifres2320 = False


}

ifres2319 = ifres2320


} else {
ifres2319 = False


}

ifres2318 = ifres2319


} else {
ifres2318 = False


}

if True == ifres2318 {
tmp2311 := Call3(__e, SymFn(symshen_4app), V7036, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V7037)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V7037)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})(), symshen_4s)


return PrimGuard2(__e, symcn, PrimStringConcat, V7038, tmp2311)


} else {
tmp2316 := Call1(__e, SymFn(symshen_4_7string_2), V7037)


if True == tmp2316 {
tmp2312 := PrimGuard1(__e, symtlstr, PrimTailString, V7037)

tmp2313 := Call1(__e, SymFn(symhdstr), V7037)


if SymFn(symshen_4insert_1h) == __self {
__self1, __self2, __self3 = V7036, tmp2312, PrimGuard2(__e, symcn, PrimStringConcat, V7038, tmp2313)
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4insert_1h), V7036, tmp2312, PrimGuard2(__e, symcn, PrimStringConcat, V7038, tmp2313))


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

tmp2355 := Call2(__e, ns2_1set, symshen_4insert_1h, tmp2302)


_ = tmp2355

tmp2356 := MakeNative3(func(__e *ControlFlow, V7039, V7040, V7041 Obj) Obj {
tmp2357 := Call2(__e, SymFn(symshen_4arg_1_6str), V7039, V7041)


return PrimGuard2(__e, symcn, PrimStringConcat, tmp2357, V7040)


panic("unreachable")
})

tmp2358 := Call2(__e, ns2_1set, symshen_4app, tmp2356)


_ = tmp2358

tmp2359 := MakeNative2(func(__e *ControlFlow, V7045, V7046 Obj) Obj {
tmp2367 := Call0(__e, SymFn(symfail))


tmp2368 := GEq(__e, V7045, tmp2367)

if True == tmp2368 {
return MakeString("...")
} else {
tmp2365 := Call1(__e, SymFn(symshen_4list_2), V7045)


if True == tmp2365 {
return TailCall2(__e, SymFn(symshen_4list_1_6str), V7045, V7046)
} else {
tmp2363 := PrimGuard1(__e, symstring_2, PrimIsString, V7045)

if True == tmp2363 {
return TailCall2(__e, SymFn(symshen_4str_1_6str), V7045, V7046)
} else {
tmp2361 := GVectorP(__e, V7045)

if True == tmp2361 {
return TailCall2(__e, SymFn(symshen_4vector_1_6str), V7045, V7046)
} else {
return TailCall1(__e, SymFn(symshen_4atom_1_6str), V7045)
}


}


}


}


panic("unreachable")
})

tmp2369 := Call2(__e, ns2_1set, symshen_4arg_1_6str, tmp2359)


_ = tmp2369

tmp2370 := MakeNative2(func(__e *ControlFlow, V7047, V7048 Obj) Obj {
tmp2378 := GEq(__e, symshen_4r, V7048)

if True == tmp2378 {
tmp2371 := Call0(__e, SymFn(symshen_4maxseq))


tmp2372 := Call3(__e, SymFn(symshen_4iter_1list), V7047, symshen_4r, tmp2371)


tmp2373 := Call2(__e, SymFn(sym_8s), tmp2372, MakeString(")"))


return TailCall2(__e, SymFn(sym_8s), MakeString("("), tmp2373)


} else {
tmp2374 := Call0(__e, SymFn(symshen_4maxseq))


tmp2375 := Call3(__e, SymFn(symshen_4iter_1list), V7047, V7048, tmp2374)


tmp2376 := Call2(__e, SymFn(sym_8s), tmp2375, MakeString("]"))


return TailCall2(__e, SymFn(sym_8s), MakeString("["), tmp2376)


}


panic("unreachable")
})

tmp2379 := Call2(__e, ns2_1set, symshen_4list_1_6str, tmp2370)


_ = tmp2379

tmp2380 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, sym_dmaximum_1print_1sequence_1size_d)
panic("unreachable")
})

tmp2381 := Call2(__e, ns2_1set, symshen_4maxseq, tmp2380)


_ = tmp2381

tmp2382 := MakeNative3(func(__e *ControlFlow, V7059, V7060, V7061 Obj) Obj {
tmp2402 := GEq(__e, Nil, V7059)

if True == tmp2402 {
return MakeString("")
} else {
tmp2400 := GEq(__e, MakeInteger(0), V7061)

if True == tmp2400 {
return MakeString("... etc")
} else {
tmp2398 := GConsP(__e, V7059)

var ifres2395 Obj

if True == tmp2398 {
tmp2396 := GTl(__e, V7059)

tmp2397 := GEq(__e, Nil, tmp2396)

ifres2395 = tmp2397


} else {
ifres2395 = False


}

if True == ifres2395 {
tmp2383 := GHd(__e, V7059)

return TailCall2(__e, SymFn(symshen_4arg_1_6str), tmp2383, V7060)


} else {
tmp2393 := GConsP(__e, V7059)

if True == tmp2393 {
tmp2384 := GHd(__e, V7059)

tmp2385 := Call2(__e, SymFn(symshen_4arg_1_6str), tmp2384, V7060)


tmp2386 := GTl(__e, V7059)

tmp2388 := Call3(__e, SymFn(symshen_4iter_1list), tmp2386, V7060, GSub(__e, V7061, MakeInteger(1)))


tmp2389 := Call2(__e, SymFn(sym_8s), MakeString(" "), tmp2388)


return TailCall2(__e, SymFn(sym_8s), tmp2385, tmp2389)


} else {
tmp2390 := Call2(__e, SymFn(symshen_4arg_1_6str), V7059, V7060)


tmp2391 := Call2(__e, SymFn(sym_8s), MakeString(" "), tmp2390)


return TailCall2(__e, SymFn(sym_8s), MakeString("|"), tmp2391)


}


}


}


}


panic("unreachable")
})

tmp2403 := Call2(__e, ns2_1set, symshen_4iter_1list, tmp2382)


_ = tmp2403

tmp2404 := MakeNative2(func(__e *ControlFlow, V7064, V7065 Obj) Obj {
tmp2409 := GEq(__e, symshen_4a, V7065)

if True == tmp2409 {
return V7064
} else {
tmp2405 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, MakeInteger(34))

tmp2406 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, MakeInteger(34))

tmp2407 := Call2(__e, SymFn(sym_8s), V7064, tmp2406)


return TailCall2(__e, SymFn(sym_8s), tmp2405, tmp2407)


}


panic("unreachable")
})

tmp2410 := Call2(__e, ns2_1set, symshen_4str_1_6str, tmp2404)


_ = tmp2410

tmp2411 := MakeNative2(func(__e *ControlFlow, V7066, V7067 Obj) Obj {
tmp2424 := Call1(__e, SymFn(symshen_4print_1vector_2), V7066)


if True == tmp2424 {
tmp2412 := GVecGet(__e, V7066, MakeInteger(0))

tmp2413 := Call1(__e, SymFn(symfn), tmp2412)


return TailCall1(__e, tmp2413, V7066)


} else {
tmp2422 := Call1(__e, SymFn(symvector_2), V7066)


if True == tmp2422 {
tmp2414 := Call0(__e, SymFn(symshen_4maxseq))


tmp2415 := Call4(__e, SymFn(symshen_4iter_1vector), V7066, MakeInteger(1), V7067, tmp2414)


tmp2416 := Call2(__e, SymFn(sym_8s), tmp2415, MakeString(">"))


return TailCall2(__e, SymFn(sym_8s), MakeString("<"), tmp2416)


} else {
tmp2417 := Call0(__e, SymFn(symshen_4maxseq))


tmp2418 := Call4(__e, SymFn(symshen_4iter_1vector), V7066, MakeInteger(0), V7067, tmp2417)


tmp2419 := Call2(__e, SymFn(sym_8s), tmp2418, MakeString(">>"))


tmp2420 := Call2(__e, SymFn(sym_8s), MakeString("<"), tmp2419)


return TailCall2(__e, SymFn(sym_8s), MakeString("<"), tmp2420)


}


}


panic("unreachable")
})

tmp2425 := Call2(__e, ns2_1set, symshen_4vector_1_6str, tmp2411)


_ = tmp2425

tmp2426 := MakeNative1(func(__e *ControlFlow, V7068 Obj) Obj {
tmp2427 := GValue(__e, symshen_4_dempty_1absvector_d)

return GEq(__e, V7068, tmp2427)


panic("unreachable")
})

tmp2428 := Call2(__e, ns2_1set, symshen_4empty_1absvector_2, tmp2426)


_ = tmp2428

tmp2429 := MakeNative1(func(__e *ControlFlow, V7069 Obj) Obj {
tmp2441 := Call1(__e, SymFn(symshen_4empty_1absvector_2), V7069)


if True == GNot(__e, tmp2441) {
tmp2430 := GVecGet(__e, V7069, MakeInteger(0))

let__2023 := tmp2430
_ = let__2023

tmp2439 := GEq(__e, let__2023, symshen_4tuple)

if True == tmp2439 {
return True
} else {
tmp2437 := GEq(__e, let__2023, symshen_4pvar)

if True == tmp2437 {
return True
} else {
tmp2435 := GEq(__e, let__2023, symshen_4dictionary)

if True == tmp2435 {
return True
} else {
tmp2432 := PrimGuard1(__e, symnumber_2, PrimIsNumber, let__2023)

if True == GNot(__e, tmp2432) {
return TailCall1(__e, SymFn(symshen_4fbound_2), let__2023)
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

tmp2443 := Call2(__e, ns2_1set, symshen_4print_1vector_2, tmp2429)


_ = tmp2443

tmp2444 := MakeNative1(func(__e *ControlFlow, V7071 Obj) Obj {
tmp2445 := Call1(__e, SymFn(symarity), V7071)


tmp2446 := GEq(__e, tmp2445, MakeInteger(-1))

return GNot(__e, tmp2446)


panic("unreachable")
})

tmp2447 := Call2(__e, ns2_1set, symshen_4fbound_2, tmp2444)


_ = tmp2447

tmp2448 := MakeNative1(func(__e *ControlFlow, V7072 Obj) Obj {
tmp2449 := GVecGet(__e, V7072, MakeInteger(1))

tmp2450 := GVecGet(__e, V7072, MakeInteger(2))

tmp2451 := Call3(__e, SymFn(symshen_4app), tmp2450, MakeString(")"), symshen_4s)


tmp2453 := Call3(__e, SymFn(symshen_4app), tmp2449, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" "), tmp2451), symshen_4s)


return PrimGuard2(__e, symcn, PrimStringConcat, MakeString("(@p "), tmp2453)


panic("unreachable")
})

tmp2454 := Call2(__e, ns2_1set, symshen_4tuple, tmp2448)


_ = tmp2454

tmp2455 := MakeNative1(func(__e *ControlFlow, V7073 Obj) Obj {
return MakeString("(dict ...)")
panic("unreachable")
})

tmp2456 := Call2(__e, ns2_1set, symshen_4dictionary, tmp2455)


_ = tmp2456

tmp2457 := MakeNative4(func(__e *ControlFlow, V7080, V7081, V7082, V7083 Obj) Obj {
tmp2475 := GEq(__e, MakeInteger(0), V7083)

if True == tmp2475 {
return MakeString("... etc")
} else {
tmp2458 := MakeNative0(func(__e *ControlFlow) Obj {
return GVecGet(__e, V7080, V7081)
panic("unreachable")
})

tmp2459 := MakeNative1(func(__e *ControlFlow, Z7085 Obj) Obj {
return symshen_4out_1of_1bounds
panic("unreachable")
})

tmp2460 := Call2(__e, try_1catch, tmp2458, tmp2459)


let__2024 := tmp2460
_ = let__2024

tmp2461 := MakeNative0(func(__e *ControlFlow) Obj {
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

tmp2463 := MakeNative1(func(__e *ControlFlow, Z7087 Obj) Obj {
return symshen_4out_1of_1bounds
panic("unreachable")
})

tmp2464 := Call2(__e, try_1catch, tmp2461, tmp2463)


let__2025 := tmp2464
_ = let__2025

tmp2473 := GEq(__e, let__2024, symshen_4out_1of_1bounds)

if True == tmp2473 {
return MakeString("")
} else {
tmp2471 := GEq(__e, let__2025, symshen_4out_1of_1bounds)

if True == tmp2471 {
return TailCall2(__e, SymFn(symshen_4arg_1_6str), let__2024, V7082)
} else {
tmp2465 := Call2(__e, SymFn(symshen_4arg_1_6str), let__2024, V7082)


tmp2466 := GAdd(__e, V7081, MakeInteger(1))

tmp2468 := Call4(__e, SymFn(symshen_4iter_1vector), V7080, tmp2466, V7082, GSub(__e, V7083, MakeInteger(1)))


tmp2469 := Call2(__e, SymFn(sym_8s), MakeString(" "), tmp2468)


return TailCall2(__e, SymFn(sym_8s), tmp2465, tmp2469)


}


}


}


panic("unreachable")
})

tmp2476 := Call2(__e, ns2_1set, symshen_4iter_1vector, tmp2457)


_ = tmp2476

tmp2477 := MakeNative1(func(__e *ControlFlow, V7088 Obj) Obj {
tmp2478 := MakeNative0(func(__e *ControlFlow) Obj {
return PrimGuard1(__e, symstr, PrimStr, V7088)
panic("unreachable")
})

tmp2479 := MakeNative1(func(__e *ControlFlow, Z7089 Obj) Obj {
return TailCall0(__e, SymFn(symshen_4funexstring))
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp2478, tmp2479)


panic("unreachable")
})

tmp2480 := Call2(__e, ns2_1set, symshen_4atom_1_6str, tmp2477)


_ = tmp2480

tmp2481 := MakeNative0(func(__e *ControlFlow) Obj {
tmp2482 := PrimGuard1(__e, symintern, PrimIntern, MakeString("x"))

tmp2483 := Call1(__e, SymFn(symgensym), tmp2482)


tmp2484 := Call2(__e, SymFn(symshen_4arg_1_6str), tmp2483, symshen_4a)


tmp2485 := Call2(__e, SymFn(sym_8s), tmp2484, MakeString("\x11"))


tmp2486 := Call2(__e, SymFn(sym_8s), MakeString("e"), tmp2485)


tmp2487 := Call2(__e, SymFn(sym_8s), MakeString("n"), tmp2486)


tmp2488 := Call2(__e, SymFn(sym_8s), MakeString("u"), tmp2487)


tmp2489 := Call2(__e, SymFn(sym_8s), MakeString("f"), tmp2488)


return TailCall2(__e, SymFn(sym_8s), MakeString("\x10"), tmp2489)


panic("unreachable")
})

tmp2490 := Call2(__e, ns2_1set, symshen_4funexstring, tmp2481)


_ = tmp2490

tmp2491 := MakeNative1(func(__e *ControlFlow, V7090 Obj) Obj {
tmp2493 := Call1(__e, SymFn(symempty_2), V7090)


if True == tmp2493 {
return True
} else {
return GConsP(__e, V7090)
}


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4list_2, tmp2491)




panic("unreachable")
})

