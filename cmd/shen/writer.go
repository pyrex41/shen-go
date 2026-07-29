package main

import . "github.com/pyrex41/shen-go/kl"

var WriterMain = MakeNative(func(__e *ControlFlow) {
tmp2098 := MakeNative(func(__e *ControlFlow) {
V5561 := __e.Get(1)
_ = V5561
tmp2099 := MakeNative(func(__e *ControlFlow) {
W5562 := __e.Get(1)
_ = W5562
tmp2100 := MakeNative(func(__e *ControlFlow) {
W5563 := __e.Get(1)
_ = W5563
__e.Return(V5561)
return
}, 1)

tmp2101 := Call(__e, PrimFunc(symstoutput))


tmp2102 := Call(__e, PrimFunc(sympr), W5562, tmp2101)


__e.TailApply(tmp2100, tmp2102)
return


}, 1)

tmp2103 := Call(__e, PrimFunc(symshen_4insert), V5561, MakeString("~S"))


__e.TailApply(tmp2099, tmp2103)
return


}, 1)

tmp2104 := Call(__e, ns2_1set, symprint, tmp2098)


_ = tmp2104

tmp2105 := MakeNative(func(__e *ControlFlow) {
V5564 := __e.Get(1)
_ = V5564
V5565 := __e.Get(2)
_ = V5565
tmp2110 := PrimValue(sym_dhush_d)

if True == tmp2110 {
__e.Return(V5564)
return
} else {
tmp2108 := Call(__e, PrimFunc(symshen_4char_1stoutput_2), V5565)


if True == tmp2108 {
__e.TailApply(PrimFunc(symshen_4write_1string), V5564, V5565)
return
} else {
tmp2106 := Call(__e, PrimFunc(symshen_4string_1_6byte), V5564, MakeNumber(0))


__e.TailApply(PrimFunc(symshen_4write_1chars), V5564, V5565, tmp2106, MakeNumber(1))
return


}


}


}, 2)

tmp2111 := Call(__e, ns2_1set, sympr, tmp2105)


_ = tmp2111

tmp2112 := MakeNative(func(__e *ControlFlow) {
V5566 := __e.Get(1)
_ = V5566
V5567 := __e.Get(2)
_ = V5567
tmp2113 := MakeNative(func(__e *ControlFlow) {
tmp2114 := PrimPos(V5566, V5567)

__e.Return(PrimStringToNumber(tmp2114))
return


}, 0)

tmp2115 := MakeNative(func(__e *ControlFlow) {
Z5568 := __e.Get(1)
_ = Z5568
__e.Return(symshen_4eos)
return
}, 1)

__e.TailApply(try_1catch, tmp2113, tmp2115)
return


}, 2)

tmp2116 := Call(__e, ns2_1set, symshen_4string_1_6byte, tmp2112)


_ = tmp2116

tmp2117 := MakeNative(func(__e *ControlFlow) {
V5569 := __e.Get(1)
_ = V5569
V5570 := __e.Get(2)
_ = V5570
V5571 := __e.Get(3)
_ = V5571
V5572 := __e.Get(4)
_ = V5572
tmp2122 := PrimEqual(symshen_4eos, V5571)

if True == tmp2122 {
__e.Return(V5569)
return
} else {
tmp2118 := PrimWriteByte(V5571, V5570)

_ = tmp2118

tmp2119 := Call(__e, PrimFunc(symshen_4string_1_6byte), V5569, V5572)


tmp2120 := PrimNumberAdd(V5572, MakeNumber(1))

__e.TailApply(PrimFunc(symshen_4write_1chars), V5569, V5570, tmp2119, tmp2120)
return


}


}, 4)

tmp2123 := Call(__e, ns2_1set, symshen_4write_1chars, tmp2117)


_ = tmp2123

tmp2124 := MakeNative(func(__e *ControlFlow) {
V5573 := __e.Get(1)
_ = V5573
V5574 := __e.Get(2)
_ = V5574
tmp2129 := PrimIsString(V5573)

if True == tmp2129 {
tmp2125 := Call(__e, PrimFunc(symshen_4proc_1nl), V5573)


__e.TailApply(PrimFunc(symshen_4mkstr_1l), tmp2125, V5574)
return


} else {
tmp2126 := PrimCons(V5573, Nil)

tmp2127 := PrimCons(symshen_4proc_1nl, tmp2126)

__e.TailApply(PrimFunc(symshen_4mkstr_1r), tmp2127, V5574)
return


}


}, 2)

tmp2130 := Call(__e, ns2_1set, symshen_4mkstr, tmp2124)


_ = tmp2130

tmp2131 := MakeNative(func(__e *ControlFlow) {
V5579 := __e.Get(1)
_ = V5579
V5580 := __e.Get(2)
_ = V5580
tmp2138 := PrimEqual(Nil, V5580)

if True == tmp2138 {
__e.Return(V5579)
return
} else {
tmp2136 := PrimIsPair(V5580)

if True == tmp2136 {
tmp2132 := PrimHead(V5580)

tmp2133 := Call(__e, PrimFunc(symshen_4insert_1l), tmp2132, V5579)


tmp2134 := PrimTail(V5580)

__e.TailApply(PrimFunc(symshen_4mkstr_1l), tmp2133, tmp2134)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.mkstr-l")))
return
}


}


}, 2)

tmp2139 := Call(__e, ns2_1set, symshen_4mkstr_1l, tmp2131)


_ = tmp2139

tmp2140 := MakeNative(func(__e *ControlFlow) {
V5587 := __e.Get(1)
_ = V5587
V5588 := __e.Get(2)
_ = V5588
tmp2278 := PrimEqual(MakeString(""), V5588)

if True == tmp2278 {
__e.Return(MakeString(""))
return
} else {
tmp2276 := Call(__e, PrimFunc(symshen_4_7string_2), V5588)


var ifres2263 Obj

if True == tmp2276 {
tmp2274 := Call(__e, PrimFunc(symhdstr), V5588)


tmp2275 := PrimEqual(MakeString("~"), tmp2274)

var ifres2265 Obj

if True == tmp2275 {
tmp2272 := PrimTailString(V5588)

tmp2273 := Call(__e, PrimFunc(symshen_4_7string_2), tmp2272)


var ifres2267 Obj

if True == tmp2273 {
tmp2269 := PrimTailString(V5588)

tmp2270 := Call(__e, PrimFunc(symhdstr), tmp2269)


tmp2271 := PrimEqual(MakeString("A"), tmp2270)

var ifres2268 Obj

if True == tmp2271 {
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

if True == ifres2263 {
tmp2141 := PrimTailString(V5588)

tmp2142 := PrimTailString(tmp2141)

tmp2143 := PrimCons(symshen_4a, Nil)

tmp2144 := PrimCons(tmp2142, tmp2143)

tmp2145 := PrimCons(V5587, tmp2144)

__e.Return(PrimCons(symshen_4app, tmp2145))
return


} else {
tmp2261 := Call(__e, PrimFunc(symshen_4_7string_2), V5588)


var ifres2248 Obj

if True == tmp2261 {
tmp2259 := Call(__e, PrimFunc(symhdstr), V5588)


tmp2260 := PrimEqual(MakeString("~"), tmp2259)

var ifres2250 Obj

if True == tmp2260 {
tmp2257 := PrimTailString(V5588)

tmp2258 := Call(__e, PrimFunc(symshen_4_7string_2), tmp2257)


var ifres2252 Obj

if True == tmp2258 {
tmp2254 := PrimTailString(V5588)

tmp2255 := Call(__e, PrimFunc(symhdstr), tmp2254)


tmp2256 := PrimEqual(MakeString("R"), tmp2255)

var ifres2253 Obj

if True == tmp2256 {
ifres2253 = True


} else {
ifres2253 = False


}

ifres2252 = ifres2253


} else {
ifres2252 = False


}

var ifres2251 Obj

if True == ifres2252 {
ifres2251 = True


} else {
ifres2251 = False


}

ifres2250 = ifres2251


} else {
ifres2250 = False


}

var ifres2249 Obj

if True == ifres2250 {
ifres2249 = True


} else {
ifres2249 = False


}

ifres2248 = ifres2249


} else {
ifres2248 = False


}

if True == ifres2248 {
tmp2146 := PrimTailString(V5588)

tmp2147 := PrimTailString(tmp2146)

tmp2148 := PrimCons(symshen_4r, Nil)

tmp2149 := PrimCons(tmp2147, tmp2148)

tmp2150 := PrimCons(V5587, tmp2149)

__e.Return(PrimCons(symshen_4app, tmp2150))
return


} else {
tmp2246 := Call(__e, PrimFunc(symshen_4_7string_2), V5588)


var ifres2233 Obj

if True == tmp2246 {
tmp2244 := Call(__e, PrimFunc(symhdstr), V5588)


tmp2245 := PrimEqual(MakeString("~"), tmp2244)

var ifres2235 Obj

if True == tmp2245 {
tmp2242 := PrimTailString(V5588)

tmp2243 := Call(__e, PrimFunc(symshen_4_7string_2), tmp2242)


var ifres2237 Obj

if True == tmp2243 {
tmp2239 := PrimTailString(V5588)

tmp2240 := Call(__e, PrimFunc(symhdstr), tmp2239)


tmp2241 := PrimEqual(MakeString("S"), tmp2240)

var ifres2238 Obj

if True == tmp2241 {
ifres2238 = True


} else {
ifres2238 = False


}

ifres2237 = ifres2238


} else {
ifres2237 = False


}

var ifres2236 Obj

if True == ifres2237 {
ifres2236 = True


} else {
ifres2236 = False


}

ifres2235 = ifres2236


} else {
ifres2235 = False


}

var ifres2234 Obj

if True == ifres2235 {
ifres2234 = True


} else {
ifres2234 = False


}

ifres2233 = ifres2234


} else {
ifres2233 = False


}

if True == ifres2233 {
tmp2151 := PrimTailString(V5588)

tmp2152 := PrimTailString(tmp2151)

tmp2153 := PrimCons(symshen_4s, Nil)

tmp2154 := PrimCons(tmp2152, tmp2153)

tmp2155 := PrimCons(V5587, tmp2154)

__e.Return(PrimCons(symshen_4app, tmp2155))
return


} else {
tmp2231 := Call(__e, PrimFunc(symshen_4_7string_2), V5588)


if True == tmp2231 {
tmp2156 := Call(__e, PrimFunc(symhdstr), V5588)


tmp2157 := PrimTailString(V5588)

tmp2158 := Call(__e, PrimFunc(symshen_4insert_1l), V5587, tmp2157)


tmp2159 := PrimCons(tmp2158, Nil)

tmp2160 := PrimCons(tmp2156, tmp2159)

tmp2161 := PrimCons(symcn, tmp2160)

__e.TailApply(PrimFunc(symshen_4factor_1cn), tmp2161)
return


} else {
tmp2229 := PrimIsPair(V5588)

var ifres2210 Obj

if True == tmp2229 {
tmp2227 := PrimHead(V5588)

tmp2228 := PrimEqual(symcn, tmp2227)

var ifres2212 Obj

if True == tmp2228 {
tmp2225 := PrimTail(V5588)

tmp2226 := PrimIsPair(tmp2225)

var ifres2214 Obj

if True == tmp2226 {
tmp2222 := PrimTail(V5588)

tmp2223 := PrimTail(tmp2222)

tmp2224 := PrimIsPair(tmp2223)

var ifres2216 Obj

if True == tmp2224 {
tmp2218 := PrimTail(V5588)

tmp2219 := PrimTail(tmp2218)

tmp2220 := PrimTail(tmp2219)

tmp2221 := PrimEqual(Nil, tmp2220)

var ifres2217 Obj

if True == tmp2221 {
ifres2217 = True


} else {
ifres2217 = False


}

ifres2216 = ifres2217


} else {
ifres2216 = False


}

var ifres2215 Obj

if True == ifres2216 {
ifres2215 = True


} else {
ifres2215 = False


}

ifres2214 = ifres2215


} else {
ifres2214 = False


}

var ifres2213 Obj

if True == ifres2214 {
ifres2213 = True


} else {
ifres2213 = False


}

ifres2212 = ifres2213


} else {
ifres2212 = False


}

var ifres2211 Obj

if True == ifres2212 {
ifres2211 = True


} else {
ifres2211 = False


}

ifres2210 = ifres2211


} else {
ifres2210 = False


}

if True == ifres2210 {
tmp2162 := PrimTail(V5588)

tmp2163 := PrimHead(tmp2162)

tmp2164 := PrimTail(V5588)

tmp2165 := PrimTail(tmp2164)

tmp2166 := PrimHead(tmp2165)

tmp2167 := Call(__e, PrimFunc(symshen_4insert_1l), V5587, tmp2166)


tmp2168 := PrimCons(tmp2167, Nil)

tmp2169 := PrimCons(tmp2163, tmp2168)

__e.Return(PrimCons(symcn, tmp2169))
return


} else {
tmp2208 := PrimIsPair(V5588)

var ifres2182 Obj

if True == tmp2208 {
tmp2206 := PrimHead(V5588)

tmp2207 := PrimEqual(symshen_4app, tmp2206)

var ifres2184 Obj

if True == tmp2207 {
tmp2204 := PrimTail(V5588)

tmp2205 := PrimIsPair(tmp2204)

var ifres2186 Obj

if True == tmp2205 {
tmp2201 := PrimTail(V5588)

tmp2202 := PrimTail(tmp2201)

tmp2203 := PrimIsPair(tmp2202)

var ifres2188 Obj

if True == tmp2203 {
tmp2197 := PrimTail(V5588)

tmp2198 := PrimTail(tmp2197)

tmp2199 := PrimTail(tmp2198)

tmp2200 := PrimIsPair(tmp2199)

var ifres2190 Obj

if True == tmp2200 {
tmp2192 := PrimTail(V5588)

tmp2193 := PrimTail(tmp2192)

tmp2194 := PrimTail(tmp2193)

tmp2195 := PrimTail(tmp2194)

tmp2196 := PrimEqual(Nil, tmp2195)

var ifres2191 Obj

if True == tmp2196 {
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

var ifres2187 Obj

if True == ifres2188 {
ifres2187 = True


} else {
ifres2187 = False


}

ifres2186 = ifres2187


} else {
ifres2186 = False


}

var ifres2185 Obj

if True == ifres2186 {
ifres2185 = True


} else {
ifres2185 = False


}

ifres2184 = ifres2185


} else {
ifres2184 = False


}

var ifres2183 Obj

if True == ifres2184 {
ifres2183 = True


} else {
ifres2183 = False


}

ifres2182 = ifres2183


} else {
ifres2182 = False


}

if True == ifres2182 {
tmp2170 := PrimTail(V5588)

tmp2171 := PrimHead(tmp2170)

tmp2172 := PrimTail(V5588)

tmp2173 := PrimTail(tmp2172)

tmp2174 := PrimHead(tmp2173)

tmp2175 := Call(__e, PrimFunc(symshen_4insert_1l), V5587, tmp2174)


tmp2176 := PrimTail(V5588)

tmp2177 := PrimTail(tmp2176)

tmp2178 := PrimTail(tmp2177)

tmp2179 := PrimCons(tmp2175, tmp2178)

tmp2180 := PrimCons(tmp2171, tmp2179)

__e.Return(PrimCons(symshen_4app, tmp2180))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.insert-l")))
return
}


}


}


}


}


}


}


}, 2)

tmp2279 := Call(__e, ns2_1set, symshen_4insert_1l, tmp2140)


_ = tmp2279

tmp2280 := MakeNative(func(__e *ControlFlow) {
V5589 := __e.Get(1)
_ = V5589
tmp2365 := PrimIsPair(V5589)

var ifres2296 Obj

if True == tmp2365 {
tmp2363 := PrimHead(V5589)

tmp2364 := PrimEqual(symcn, tmp2363)

var ifres2298 Obj

if True == tmp2364 {
tmp2361 := PrimTail(V5589)

tmp2362 := PrimIsPair(tmp2361)

var ifres2300 Obj

if True == tmp2362 {
tmp2358 := PrimTail(V5589)

tmp2359 := PrimTail(tmp2358)

tmp2360 := PrimIsPair(tmp2359)

var ifres2302 Obj

if True == tmp2360 {
tmp2354 := PrimTail(V5589)

tmp2355 := PrimTail(tmp2354)

tmp2356 := PrimHead(tmp2355)

tmp2357 := PrimIsPair(tmp2356)

var ifres2304 Obj

if True == tmp2357 {
tmp2349 := PrimTail(V5589)

tmp2350 := PrimTail(tmp2349)

tmp2351 := PrimHead(tmp2350)

tmp2352 := PrimHead(tmp2351)

tmp2353 := PrimEqual(symcn, tmp2352)

var ifres2306 Obj

if True == tmp2353 {
tmp2344 := PrimTail(V5589)

tmp2345 := PrimTail(tmp2344)

tmp2346 := PrimHead(tmp2345)

tmp2347 := PrimTail(tmp2346)

tmp2348 := PrimIsPair(tmp2347)

var ifres2308 Obj

if True == tmp2348 {
tmp2338 := PrimTail(V5589)

tmp2339 := PrimTail(tmp2338)

tmp2340 := PrimHead(tmp2339)

tmp2341 := PrimTail(tmp2340)

tmp2342 := PrimTail(tmp2341)

tmp2343 := PrimIsPair(tmp2342)

var ifres2310 Obj

if True == tmp2343 {
tmp2331 := PrimTail(V5589)

tmp2332 := PrimTail(tmp2331)

tmp2333 := PrimHead(tmp2332)

tmp2334 := PrimTail(tmp2333)

tmp2335 := PrimTail(tmp2334)

tmp2336 := PrimTail(tmp2335)

tmp2337 := PrimEqual(Nil, tmp2336)

var ifres2312 Obj

if True == tmp2337 {
tmp2327 := PrimTail(V5589)

tmp2328 := PrimTail(tmp2327)

tmp2329 := PrimTail(tmp2328)

tmp2330 := PrimEqual(Nil, tmp2329)

var ifres2314 Obj

if True == tmp2330 {
tmp2324 := PrimTail(V5589)

tmp2325 := PrimHead(tmp2324)

tmp2326 := PrimIsString(tmp2325)

var ifres2316 Obj

if True == tmp2326 {
tmp2318 := PrimTail(V5589)

tmp2319 := PrimTail(tmp2318)

tmp2320 := PrimHead(tmp2319)

tmp2321 := PrimTail(tmp2320)

tmp2322 := PrimHead(tmp2321)

tmp2323 := PrimIsString(tmp2322)

var ifres2317 Obj

if True == tmp2323 {
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

var ifres2301 Obj

if True == ifres2302 {
ifres2301 = True


} else {
ifres2301 = False


}

ifres2300 = ifres2301


} else {
ifres2300 = False


}

var ifres2299 Obj

if True == ifres2300 {
ifres2299 = True


} else {
ifres2299 = False


}

ifres2298 = ifres2299


} else {
ifres2298 = False


}

var ifres2297 Obj

if True == ifres2298 {
ifres2297 = True


} else {
ifres2297 = False


}

ifres2296 = ifres2297


} else {
ifres2296 = False


}

if True == ifres2296 {
tmp2281 := PrimTail(V5589)

tmp2282 := PrimHead(tmp2281)

tmp2283 := PrimTail(V5589)

tmp2284 := PrimTail(tmp2283)

tmp2285 := PrimHead(tmp2284)

tmp2286 := PrimTail(tmp2285)

tmp2287 := PrimHead(tmp2286)

tmp2288 := PrimStringConcat(tmp2282, tmp2287)

tmp2289 := PrimTail(V5589)

tmp2290 := PrimTail(tmp2289)

tmp2291 := PrimHead(tmp2290)

tmp2292 := PrimTail(tmp2291)

tmp2293 := PrimTail(tmp2292)

tmp2294 := PrimCons(tmp2288, tmp2293)

__e.Return(PrimCons(symcn, tmp2294))
return


} else {
__e.Return(V5589)
return
}


}, 1)

tmp2366 := Call(__e, ns2_1set, symshen_4factor_1cn, tmp2280)


_ = tmp2366

tmp2367 := MakeNative(func(__e *ControlFlow) {
V5592 := __e.Get(1)
_ = V5592
tmp2393 := PrimEqual(MakeString(""), V5592)

if True == tmp2393 {
__e.Return(MakeString(""))
return
} else {
tmp2391 := Call(__e, PrimFunc(symshen_4_7string_2), V5592)


var ifres2378 Obj

if True == tmp2391 {
tmp2389 := Call(__e, PrimFunc(symhdstr), V5592)


tmp2390 := PrimEqual(MakeString("~"), tmp2389)

var ifres2380 Obj

if True == tmp2390 {
tmp2387 := PrimTailString(V5592)

tmp2388 := Call(__e, PrimFunc(symshen_4_7string_2), tmp2387)


var ifres2382 Obj

if True == tmp2388 {
tmp2384 := PrimTailString(V5592)

tmp2385 := Call(__e, PrimFunc(symhdstr), tmp2384)


tmp2386 := PrimEqual(MakeString("%"), tmp2385)

var ifres2383 Obj

if True == tmp2386 {
ifres2383 = True


} else {
ifres2383 = False


}

ifres2382 = ifres2383


} else {
ifres2382 = False


}

var ifres2381 Obj

if True == ifres2382 {
ifres2381 = True


} else {
ifres2381 = False


}

ifres2380 = ifres2381


} else {
ifres2380 = False


}

var ifres2379 Obj

if True == ifres2380 {
ifres2379 = True


} else {
ifres2379 = False


}

ifres2378 = ifres2379


} else {
ifres2378 = False


}

if True == ifres2378 {
tmp2368 := PrimNumberToString(MakeNumber(10))

tmp2369 := PrimTailString(V5592)

tmp2370 := PrimTailString(tmp2369)

tmp2371 := Call(__e, PrimFunc(symshen_4proc_1nl), tmp2370)


__e.Return(PrimStringConcat(tmp2368, tmp2371))
return


} else {
tmp2376 := Call(__e, PrimFunc(symshen_4_7string_2), V5592)


if True == tmp2376 {
tmp2372 := Call(__e, PrimFunc(symhdstr), V5592)


tmp2373 := PrimTailString(V5592)

tmp2374 := Call(__e, PrimFunc(symshen_4proc_1nl), tmp2373)


__e.Return(PrimStringConcat(tmp2372, tmp2374))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.proc-nl")))
return
}


}


}


}, 1)

tmp2394 := Call(__e, ns2_1set, symshen_4proc_1nl, tmp2367)


_ = tmp2394

tmp2395 := MakeNative(func(__e *ControlFlow) {
V5597 := __e.Get(1)
_ = V5597
V5598 := __e.Get(2)
_ = V5598
tmp2404 := PrimEqual(Nil, V5598)

if True == tmp2404 {
__e.Return(V5597)
return
} else {
tmp2402 := PrimIsPair(V5598)

if True == tmp2402 {
tmp2396 := PrimHead(V5598)

tmp2397 := PrimCons(V5597, Nil)

tmp2398 := PrimCons(tmp2396, tmp2397)

tmp2399 := PrimCons(symshen_4insert, tmp2398)

tmp2400 := PrimTail(V5598)

__e.TailApply(PrimFunc(symshen_4mkstr_1r), tmp2399, tmp2400)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.mkstr-r")))
return
}


}


}, 2)

tmp2405 := Call(__e, ns2_1set, symshen_4mkstr_1r, tmp2395)


_ = tmp2405

tmp2406 := MakeNative(func(__e *ControlFlow) {
V5599 := __e.Get(1)
_ = V5599
V5600 := __e.Get(2)
_ = V5600
__e.TailApply(PrimFunc(symshen_4insert_1h), V5599, V5600, MakeString(""))
return
}, 2)

tmp2407 := Call(__e, ns2_1set, symshen_4insert, tmp2406)


_ = tmp2407

tmp2408 := MakeNative(func(__e *ControlFlow) {
V5609 := __e.Get(1)
_ = V5609
V5610 := __e.Get(2)
_ = V5610
V5611 := __e.Get(3)
_ = V5611
tmp2469 := PrimEqual(MakeString(""), V5610)

if True == tmp2469 {
__e.Return(V5611)
return
} else {
tmp2467 := Call(__e, PrimFunc(symshen_4_7string_2), V5610)


var ifres2454 Obj

if True == tmp2467 {
tmp2465 := Call(__e, PrimFunc(symhdstr), V5610)


tmp2466 := PrimEqual(MakeString("~"), tmp2465)

var ifres2456 Obj

if True == tmp2466 {
tmp2463 := PrimTailString(V5610)

tmp2464 := Call(__e, PrimFunc(symshen_4_7string_2), tmp2463)


var ifres2458 Obj

if True == tmp2464 {
tmp2460 := PrimTailString(V5610)

tmp2461 := Call(__e, PrimFunc(symhdstr), tmp2460)


tmp2462 := PrimEqual(MakeString("A"), tmp2461)

var ifres2459 Obj

if True == tmp2462 {
ifres2459 = True


} else {
ifres2459 = False


}

ifres2458 = ifres2459


} else {
ifres2458 = False


}

var ifres2457 Obj

if True == ifres2458 {
ifres2457 = True


} else {
ifres2457 = False


}

ifres2456 = ifres2457


} else {
ifres2456 = False


}

var ifres2455 Obj

if True == ifres2456 {
ifres2455 = True


} else {
ifres2455 = False


}

ifres2454 = ifres2455


} else {
ifres2454 = False


}

if True == ifres2454 {
tmp2409 := PrimTailString(V5610)

tmp2410 := PrimTailString(tmp2409)

tmp2411 := Call(__e, PrimFunc(symshen_4app), V5609, tmp2410, symshen_4a)


__e.Return(PrimStringConcat(V5611, tmp2411))
return


} else {
tmp2452 := Call(__e, PrimFunc(symshen_4_7string_2), V5610)


var ifres2439 Obj

if True == tmp2452 {
tmp2450 := Call(__e, PrimFunc(symhdstr), V5610)


tmp2451 := PrimEqual(MakeString("~"), tmp2450)

var ifres2441 Obj

if True == tmp2451 {
tmp2448 := PrimTailString(V5610)

tmp2449 := Call(__e, PrimFunc(symshen_4_7string_2), tmp2448)


var ifres2443 Obj

if True == tmp2449 {
tmp2445 := PrimTailString(V5610)

tmp2446 := Call(__e, PrimFunc(symhdstr), tmp2445)


tmp2447 := PrimEqual(MakeString("R"), tmp2446)

var ifres2444 Obj

if True == tmp2447 {
ifres2444 = True


} else {
ifres2444 = False


}

ifres2443 = ifres2444


} else {
ifres2443 = False


}

var ifres2442 Obj

if True == ifres2443 {
ifres2442 = True


} else {
ifres2442 = False


}

ifres2441 = ifres2442


} else {
ifres2441 = False


}

var ifres2440 Obj

if True == ifres2441 {
ifres2440 = True


} else {
ifres2440 = False


}

ifres2439 = ifres2440


} else {
ifres2439 = False


}

if True == ifres2439 {
tmp2412 := PrimTailString(V5610)

tmp2413 := PrimTailString(tmp2412)

tmp2414 := Call(__e, PrimFunc(symshen_4app), V5609, tmp2413, symshen_4r)


__e.Return(PrimStringConcat(V5611, tmp2414))
return


} else {
tmp2437 := Call(__e, PrimFunc(symshen_4_7string_2), V5610)


var ifres2424 Obj

if True == tmp2437 {
tmp2435 := Call(__e, PrimFunc(symhdstr), V5610)


tmp2436 := PrimEqual(MakeString("~"), tmp2435)

var ifres2426 Obj

if True == tmp2436 {
tmp2433 := PrimTailString(V5610)

tmp2434 := Call(__e, PrimFunc(symshen_4_7string_2), tmp2433)


var ifres2428 Obj

if True == tmp2434 {
tmp2430 := PrimTailString(V5610)

tmp2431 := Call(__e, PrimFunc(symhdstr), tmp2430)


tmp2432 := PrimEqual(MakeString("S"), tmp2431)

var ifres2429 Obj

if True == tmp2432 {
ifres2429 = True


} else {
ifres2429 = False


}

ifres2428 = ifres2429


} else {
ifres2428 = False


}

var ifres2427 Obj

if True == ifres2428 {
ifres2427 = True


} else {
ifres2427 = False


}

ifres2426 = ifres2427


} else {
ifres2426 = False


}

var ifres2425 Obj

if True == ifres2426 {
ifres2425 = True


} else {
ifres2425 = False


}

ifres2424 = ifres2425


} else {
ifres2424 = False


}

if True == ifres2424 {
tmp2415 := PrimTailString(V5610)

tmp2416 := PrimTailString(tmp2415)

tmp2417 := Call(__e, PrimFunc(symshen_4app), V5609, tmp2416, symshen_4s)


__e.Return(PrimStringConcat(V5611, tmp2417))
return


} else {
tmp2422 := Call(__e, PrimFunc(symshen_4_7string_2), V5610)


if True == tmp2422 {
tmp2418 := PrimTailString(V5610)

tmp2419 := Call(__e, PrimFunc(symhdstr), V5610)


tmp2420 := PrimStringConcat(V5611, tmp2419)

__e.TailApply(PrimFunc(symshen_4insert_1h), V5609, tmp2418, tmp2420)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.insert-h")))
return
}


}


}


}


}


}, 3)

tmp2470 := Call(__e, ns2_1set, symshen_4insert_1h, tmp2408)


_ = tmp2470

tmp2471 := MakeNative(func(__e *ControlFlow) {
V5612 := __e.Get(1)
_ = V5612
V5613 := __e.Get(2)
_ = V5613
V5614 := __e.Get(3)
_ = V5614
tmp2472 := Call(__e, PrimFunc(symshen_4arg_1_6str), V5612, V5614)


__e.Return(PrimStringConcat(tmp2472, V5613))
return


}, 3)

tmp2473 := Call(__e, ns2_1set, symshen_4app, tmp2471)


_ = tmp2473

tmp2474 := MakeNative(func(__e *ControlFlow) {
V5618 := __e.Get(1)
_ = V5618
V5619 := __e.Get(2)
_ = V5619
tmp2482 := Call(__e, PrimFunc(symfail))


tmp2483 := PrimEqual(V5618, tmp2482)

if True == tmp2483 {
__e.Return(MakeString("..."))
return
} else {
tmp2480 := Call(__e, PrimFunc(symshen_4list_2), V5618)


if True == tmp2480 {
__e.TailApply(PrimFunc(symshen_4list_1_6str), V5618, V5619)
return
} else {
tmp2478 := PrimIsString(V5618)

if True == tmp2478 {
__e.TailApply(PrimFunc(symshen_4str_1_6str), V5618, V5619)
return
} else {
tmp2476 := PrimIsVector(V5618)

if True == tmp2476 {
__e.TailApply(PrimFunc(symshen_4vector_1_6str), V5618, V5619)
return
} else {
__e.TailApply(PrimFunc(symshen_4atom_1_6str), V5618)
return
}


}


}


}


}, 2)

tmp2484 := Call(__e, ns2_1set, symshen_4arg_1_6str, tmp2474)


_ = tmp2484

tmp2485 := MakeNative(func(__e *ControlFlow) {
V5620 := __e.Get(1)
_ = V5620
V5621 := __e.Get(2)
_ = V5621
tmp2493 := PrimEqual(symshen_4r, V5621)

if True == tmp2493 {
tmp2486 := Call(__e, PrimFunc(symshen_4maxseq))


tmp2487 := Call(__e, PrimFunc(symshen_4iter_1list), V5620, symshen_4r, tmp2486)


tmp2488 := Call(__e, PrimFunc(sym_8s), tmp2487, MakeString(")"))


__e.TailApply(PrimFunc(sym_8s), MakeString("("), tmp2488)
return


} else {
tmp2489 := Call(__e, PrimFunc(symshen_4maxseq))


tmp2490 := Call(__e, PrimFunc(symshen_4iter_1list), V5620, V5621, tmp2489)


tmp2491 := Call(__e, PrimFunc(sym_8s), tmp2490, MakeString("]"))


__e.TailApply(PrimFunc(sym_8s), MakeString("["), tmp2491)
return


}


}, 2)

tmp2494 := Call(__e, ns2_1set, symshen_4list_1_6str, tmp2485)


_ = tmp2494

tmp2495 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dmaximum_1print_1sequence_1size_d))
return
}, 0)

tmp2496 := Call(__e, ns2_1set, symshen_4maxseq, tmp2495)


_ = tmp2496

tmp2497 := MakeNative(func(__e *ControlFlow) {
V5632 := __e.Get(1)
_ = V5632
V5633 := __e.Get(2)
_ = V5633
V5634 := __e.Get(3)
_ = V5634
tmp2518 := PrimEqual(Nil, V5632)

if True == tmp2518 {
__e.Return(MakeString(""))
return
} else {
tmp2516 := PrimEqual(MakeNumber(0), V5634)

if True == tmp2516 {
__e.Return(MakeString("... etc"))
return
} else {
tmp2514 := PrimIsPair(V5632)

var ifres2510 Obj

if True == tmp2514 {
tmp2512 := PrimTail(V5632)

tmp2513 := PrimEqual(Nil, tmp2512)

var ifres2511 Obj

if True == tmp2513 {
ifres2511 = True


} else {
ifres2511 = False


}

ifres2510 = ifres2511


} else {
ifres2510 = False


}

if True == ifres2510 {
tmp2498 := PrimHead(V5632)

__e.TailApply(PrimFunc(symshen_4arg_1_6str), tmp2498, V5633)
return


} else {
tmp2508 := PrimIsPair(V5632)

if True == tmp2508 {
tmp2499 := PrimHead(V5632)

tmp2500 := Call(__e, PrimFunc(symshen_4arg_1_6str), tmp2499, V5633)


tmp2501 := PrimTail(V5632)

tmp2502 := PrimNumberSubtract(V5634, MakeNumber(1))

tmp2503 := Call(__e, PrimFunc(symshen_4iter_1list), tmp2501, V5633, tmp2502)


tmp2504 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp2503)


__e.TailApply(PrimFunc(sym_8s), tmp2500, tmp2504)
return


} else {
tmp2505 := Call(__e, PrimFunc(symshen_4arg_1_6str), V5632, V5633)


tmp2506 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp2505)


__e.TailApply(PrimFunc(sym_8s), MakeString("|"), tmp2506)
return


}


}


}


}


}, 3)

tmp2519 := Call(__e, ns2_1set, symshen_4iter_1list, tmp2497)


_ = tmp2519

tmp2520 := MakeNative(func(__e *ControlFlow) {
V5637 := __e.Get(1)
_ = V5637
V5638 := __e.Get(2)
_ = V5638
tmp2525 := PrimEqual(symshen_4a, V5638)

if True == tmp2525 {
__e.Return(V5637)
return
} else {
tmp2521 := PrimNumberToString(MakeNumber(34))

tmp2522 := PrimNumberToString(MakeNumber(34))

tmp2523 := Call(__e, PrimFunc(sym_8s), V5637, tmp2522)


__e.TailApply(PrimFunc(sym_8s), tmp2521, tmp2523)
return


}


}, 2)

tmp2526 := Call(__e, ns2_1set, symshen_4str_1_6str, tmp2520)


_ = tmp2526

tmp2527 := MakeNative(func(__e *ControlFlow) {
V5639 := __e.Get(1)
_ = V5639
V5640 := __e.Get(2)
_ = V5640
tmp2540 := Call(__e, PrimFunc(symshen_4print_1vector_2), V5639)


if True == tmp2540 {
tmp2528 := PrimVectorGet(V5639, MakeNumber(0))

tmp2529 := Call(__e, PrimFunc(symfn), tmp2528)


__e.TailApply(tmp2529, V5639)
return


} else {
tmp2538 := Call(__e, PrimFunc(symvector_2), V5639)


if True == tmp2538 {
tmp2530 := Call(__e, PrimFunc(symshen_4maxseq))


tmp2531 := Call(__e, PrimFunc(symshen_4iter_1vector), V5639, MakeNumber(1), V5640, tmp2530)


tmp2532 := Call(__e, PrimFunc(sym_8s), tmp2531, MakeString(">"))


__e.TailApply(PrimFunc(sym_8s), MakeString("<"), tmp2532)
return


} else {
tmp2533 := Call(__e, PrimFunc(symshen_4maxseq))


tmp2534 := Call(__e, PrimFunc(symshen_4iter_1vector), V5639, MakeNumber(0), V5640, tmp2533)


tmp2535 := Call(__e, PrimFunc(sym_8s), tmp2534, MakeString(">>"))


tmp2536 := Call(__e, PrimFunc(sym_8s), MakeString("<"), tmp2535)


__e.TailApply(PrimFunc(sym_8s), MakeString("<"), tmp2536)
return


}


}


}, 2)

tmp2541 := Call(__e, ns2_1set, symshen_4vector_1_6str, tmp2527)


_ = tmp2541

tmp2542 := MakeNative(func(__e *ControlFlow) {
V5641 := __e.Get(1)
_ = V5641
tmp2543 := MakeNative(func(__e *ControlFlow) {
W5642 := __e.Get(1)
_ = W5642
tmp2550 := PrimEqual(W5642, symshen_4tuple)

if True == tmp2550 {
__e.Return(True)
return
} else {
tmp2548 := PrimEqual(W5642, symshen_4pvar)

if True == tmp2548 {
__e.Return(True)
return
} else {
tmp2545 := PrimIsNumber(W5642)

tmp2546 := PrimNot(tmp2545)

if True == tmp2546 {
__e.TailApply(PrimFunc(symshen_4fbound_2), W5642)
return
} else {
__e.Return(False)
return
}


}


}


}, 1)

tmp2551 := PrimVectorGet(V5641, MakeNumber(0))

__e.TailApply(tmp2543, tmp2551)
return


}, 1)

tmp2552 := Call(__e, ns2_1set, symshen_4print_1vector_2, tmp2542)


_ = tmp2552

tmp2553 := MakeNative(func(__e *ControlFlow) {
V5643 := __e.Get(1)
_ = V5643
tmp2554 := Call(__e, PrimFunc(symarity), V5643)


tmp2555 := PrimEqual(tmp2554, MakeNumber(-1))

__e.Return(PrimNot(tmp2555))
return


}, 1)

tmp2556 := Call(__e, ns2_1set, symshen_4fbound_2, tmp2553)


_ = tmp2556

tmp2557 := MakeNative(func(__e *ControlFlow) {
V5644 := __e.Get(1)
_ = V5644
tmp2558 := PrimVectorGet(V5644, MakeNumber(1))

tmp2559 := PrimVectorGet(V5644, MakeNumber(2))

tmp2560 := Call(__e, PrimFunc(symshen_4app), tmp2559, MakeString(")"), symshen_4s)


tmp2561 := PrimStringConcat(MakeString(" "), tmp2560)

tmp2562 := Call(__e, PrimFunc(symshen_4app), tmp2558, tmp2561, symshen_4s)


__e.Return(PrimStringConcat(MakeString("(@p "), tmp2562))
return


}, 1)

tmp2563 := Call(__e, ns2_1set, symshen_4tuple, tmp2557)


_ = tmp2563

tmp2564 := MakeNative(func(__e *ControlFlow) {
V5651 := __e.Get(1)
_ = V5651
V5652 := __e.Get(2)
_ = V5652
V5653 := __e.Get(3)
_ = V5653
V5654 := __e.Get(4)
_ = V5654
tmp2584 := PrimEqual(MakeNumber(0), V5654)

if True == tmp2584 {
__e.Return(MakeString("... etc"))
return
} else {
tmp2565 := MakeNative(func(__e *ControlFlow) {
W5655 := __e.Get(1)
_ = W5655
tmp2566 := MakeNative(func(__e *ControlFlow) {
W5657 := __e.Get(1)
_ = W5657
tmp2575 := PrimEqual(W5655, symshen_4out_1of_1bounds)

if True == tmp2575 {
__e.Return(MakeString(""))
return
} else {
tmp2573 := PrimEqual(W5657, symshen_4out_1of_1bounds)

if True == tmp2573 {
__e.TailApply(PrimFunc(symshen_4arg_1_6str), W5655, V5653)
return
} else {
tmp2567 := Call(__e, PrimFunc(symshen_4arg_1_6str), W5655, V5653)


tmp2568 := PrimNumberAdd(V5652, MakeNumber(1))

tmp2569 := PrimNumberSubtract(V5654, MakeNumber(1))

tmp2570 := Call(__e, PrimFunc(symshen_4iter_1vector), V5651, tmp2568, V5653, tmp2569)


tmp2571 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp2570)


__e.TailApply(PrimFunc(sym_8s), tmp2567, tmp2571)
return


}


}


}, 1)

tmp2576 := MakeNative(func(__e *ControlFlow) {
tmp2577 := PrimNumberAdd(V5652, MakeNumber(1))

__e.Return(PrimVectorGet(V5651, tmp2577))
return


}, 0)

tmp2578 := MakeNative(func(__e *ControlFlow) {
Z5658 := __e.Get(1)
_ = Z5658
__e.Return(symshen_4out_1of_1bounds)
return
}, 1)

tmp2579 := Call(__e, try_1catch, tmp2576, tmp2578)


__e.TailApply(tmp2566, tmp2579)
return


}, 1)

tmp2580 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimVectorGet(V5651, V5652))
return
}, 0)

tmp2581 := MakeNative(func(__e *ControlFlow) {
Z5656 := __e.Get(1)
_ = Z5656
__e.Return(symshen_4out_1of_1bounds)
return
}, 1)

tmp2582 := Call(__e, try_1catch, tmp2580, tmp2581)


__e.TailApply(tmp2565, tmp2582)
return


}


}, 4)

tmp2585 := Call(__e, ns2_1set, symshen_4iter_1vector, tmp2564)


_ = tmp2585

tmp2586 := MakeNative(func(__e *ControlFlow) {
V5659 := __e.Get(1)
_ = V5659
tmp2587 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimStr(V5659))
return
}, 0)

tmp2588 := MakeNative(func(__e *ControlFlow) {
Z5660 := __e.Get(1)
_ = Z5660
__e.TailApply(PrimFunc(symshen_4funexstring))
return
}, 1)

__e.TailApply(try_1catch, tmp2587, tmp2588)
return


}, 1)

tmp2589 := Call(__e, ns2_1set, symshen_4atom_1_6str, tmp2586)


_ = tmp2589

tmp2590 := MakeNative(func(__e *ControlFlow) {
tmp2591 := PrimIntern(MakeString("x"))

tmp2592 := Call(__e, PrimFunc(symgensym), tmp2591)


tmp2593 := Call(__e, PrimFunc(symshen_4arg_1_6str), tmp2592, symshen_4a)


tmp2594 := Call(__e, PrimFunc(sym_8s), tmp2593, MakeString("\x11"))


tmp2595 := Call(__e, PrimFunc(sym_8s), MakeString("e"), tmp2594)


tmp2596 := Call(__e, PrimFunc(sym_8s), MakeString("n"), tmp2595)


tmp2597 := Call(__e, PrimFunc(sym_8s), MakeString("u"), tmp2596)


tmp2598 := Call(__e, PrimFunc(sym_8s), MakeString("f"), tmp2597)


__e.TailApply(PrimFunc(sym_8s), MakeString("\x10"), tmp2598)
return


}, 0)

tmp2599 := Call(__e, ns2_1set, symshen_4funexstring, tmp2590)


_ = tmp2599

tmp2600 := MakeNative(func(__e *ControlFlow) {
V5661 := __e.Get(1)
_ = V5661
tmp2604 := Call(__e, PrimFunc(symempty_2), V5661)


if True == tmp2604 {
__e.Return(True)
return
} else {
tmp2602 := PrimIsPair(V5661)

if True == tmp2602 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

__e.TailApply(ns2_1set, symshen_4list_2, tmp2600)
return




}, 0)

