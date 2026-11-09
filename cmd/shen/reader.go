package main

import . "github.com/pyrex41/shen-go/kl"

var ReaderMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp5169 := MakeNative1(func(__e *ControlFlow, V2556 Obj) Obj {
tmp5170 := PrimGuard1(__e, symread_1file_1as_1bytelist, PrimReadFileAsByteList, V2556)

let__4746 := tmp5170
_ = let__4746

tmp5171 := MakeNative0(func(__e *ControlFlow) Obj {
tmp5172 := MakeNative1(func(__e *ControlFlow, Z2559 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5s_1exprs_6), Z2559)
panic("unreachable")
})

return TailCall2(__e, SymFn(symcompile), tmp5172, let__4746)


panic("unreachable")
})

tmp5173 := MakeNative1(func(__e *ControlFlow, Z2560 Obj) Obj {
tmp5174 := GValue(__e, symshen_4_dresidue_d)

return TailCall1(__e, SymFn(symshen_4reader_1error), tmp5174)


panic("unreachable")
})

tmp5175 := Call2(__e, try_1catch, tmp5171, tmp5173)


let__4747 := tmp5175
_ = let__4747

tmp5176 := Call1(__e, SymFn(symshen_4process_1sexprs), let__4747)


let__4748 := tmp5176
_ = let__4748

return let__4748


panic("unreachable")
})

tmp5177 := Call2(__e, ns2_1set, symread_1file, tmp5169)


_ = tmp5177

tmp5178 := MakeNative1(func(__e *ControlFlow, V2562 Obj) Obj {
tmp5179 := GValue(__e, sym_dmaximum_1print_1sequence_1size_d)

tmp5180 := Call3(__e, SymFn(symshen_4reader_1error_1message), tmp5179, MakeInteger(0), V2562)


tmp5182 := Call1(__e, SymFn(symshen_4proc_1nl), PrimGuard2(__e, symcn, PrimStringConcat, MakeString("reader error near here: "), tmp5180))


return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp5182)


panic("unreachable")
})

tmp5183 := Call2(__e, ns2_1set, symshen_4reader_1error, tmp5178)


_ = tmp5183

tmp5184 := MakeNative3(func(__e *ControlFlow, V2570, V2571, V2572 Obj) Obj {
tmp5195 := GEq(__e, Nil, V2572)

if True == tmp5195 {
return MakeString("")
} else {
tmp5193 := GEq(__e, V2570, V2571)

if True == tmp5193 {
return MakeString("")
} else {
tmp5191 := GConsP(__e, V2572)

if True == tmp5191 {
tmp5185 := GHd(__e, V2572)

tmp5186 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, tmp5185)

tmp5187 := GAdd(__e, V2571, MakeInteger(1))

tmp5188 := GTl(__e, V2572)

tmp5189 := Call3(__e, SymFn(symshen_4reader_1error_1message), V2570, tmp5187, tmp5188)


return PrimGuard2(__e, symcn, PrimStringConcat, tmp5186, tmp5189)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4reader_1error_1message)
}


}


}


panic("unreachable")
})

tmp5196 := Call2(__e, ns2_1set, symshen_4reader_1error_1message, tmp5184)


_ = tmp5196

tmp5197 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dit_d)
panic("unreachable")
})

tmp5198 := Call2(__e, ns2_1set, symit, tmp5197)


_ = tmp5198

tmp5199 := MakeNative1(func(__e *ControlFlow, V2573 Obj) Obj {
tmp5200 := PrimGuard2(__e, symopen, PrimOpenStream, V2573, symin)

let__4749 := tmp5200
_ = let__4749

tmp5201 := PrimGuard1(__e, symread_1byte, PrimReadByte, let__4749)

let__4750 := tmp5201
_ = let__4750

tmp5202 := Call3(__e, SymFn(symshen_4read_1file_1as_1bytelist_1help), let__4749, let__4750, Nil)


let__4751 := tmp5202
_ = let__4751

tmp5203 := PrimGuard1(__e, symclose, PrimCloseStream, let__4749)

let__4752 := tmp5203
_ = let__4752

return TailCall1(__e, SymFn(symreverse), let__4751)


panic("unreachable")
})

tmp5204 := Call2(__e, ns2_1set, symread_1file_1as_1bytelist, tmp5199)


_ = tmp5204

tmp5205 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V2578 := __self1
_ = V2578
V2579 := __self2
_ = V2579
V2580 := __self3
_ = V2580
tmp5209 := GEq(__e, MakeInteger(-1), V2579)

if True == tmp5209 {
return V2580
} else {
tmp5206 := PrimGuard1(__e, symread_1byte, PrimReadByte, V2578)

tmp5207 := GCons(__e, V2579, V2580)

if SymFn(symshen_4read_1file_1as_1bytelist_1help) == __self {
__self1, __self2, __self3 = V2578, tmp5206, tmp5207
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4read_1file_1as_1bytelist_1help), V2578, tmp5206, tmp5207)


}


panic("unreachable")
})
return __self
})()

tmp5210 := Call2(__e, ns2_1set, symshen_4read_1file_1as_1bytelist_1help, tmp5205)


_ = tmp5210

tmp5211 := MakeNative1(func(__e *ControlFlow, V2581 Obj) Obj {
tmp5212 := PrimGuard2(__e, symopen, PrimOpenStream, V2581, symin)

let__4753 := tmp5212
_ = let__4753

tmp5213 := PrimGuard1(__e, symread_1byte, PrimReadByte, let__4753)

return TailCall3(__e, SymFn(symshen_4rfas_1h), let__4753, tmp5213, MakeString(""))


panic("unreachable")
})

tmp5214 := Call2(__e, ns2_1set, symread_1file_1as_1string, tmp5211)


_ = tmp5214

tmp5215 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V2583 := __self1
_ = V2583
V2584 := __self2
_ = V2584
V2585 := __self3
_ = V2585
tmp5221 := GEq(__e, MakeInteger(-1), V2584)

if True == tmp5221 {
tmp5216 := PrimGuard1(__e, symclose, PrimCloseStream, V2583)

_ = tmp5216

return V2585


} else {
tmp5217 := PrimGuard1(__e, symread_1byte, PrimReadByte, V2583)

tmp5218 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, V2584)

if SymFn(symshen_4rfas_1h) == __self {
__self1, __self2, __self3 = V2583, tmp5217, PrimGuard2(__e, symcn, PrimStringConcat, V2585, tmp5218)
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4rfas_1h), V2583, tmp5217, PrimGuard2(__e, symcn, PrimStringConcat, V2585, tmp5218))


}


panic("unreachable")
})
return __self
})()

tmp5222 := Call2(__e, ns2_1set, symshen_4rfas_1h, tmp5215)


_ = tmp5222

tmp5223 := MakeNative1(func(__e *ControlFlow, V2586 Obj) Obj {
tmp5224 := Call1(__e, SymFn(symread), V2586)


return TailCall1(__e, SymFn(symeval_1kl), tmp5224)


panic("unreachable")
})

tmp5225 := Call2(__e, ns2_1set, syminput, tmp5223)


_ = tmp5225

tmp5226 := MakeNative2(func(__e *ControlFlow, V2587, V2588 Obj) Obj {
tmp5227 := Call1(__e, SymFn(symshen_4monotype), V2587)


let__4754 := tmp5227
_ = let__4754

tmp5228 := Call1(__e, SymFn(symread), V2588)


let__4755 := tmp5228
_ = let__4755

tmp5234 := Call2(__e, SymFn(symshen_4typecheck), let__4755, V2587)


tmp5235 := GEq(__e, False, tmp5234)

if True == tmp5235 {
tmp5229 := Call3(__e, SymFn(symshen_4app), V2587, MakeString("\n"), symshen_4r)


tmp5231 := Call3(__e, SymFn(symshen_4app), let__4755, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" is not of type "), tmp5229), symshen_4r)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("type error: "), tmp5231))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("type error: "), tmp5231)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


} else {
return TailCall1(__e, SymFn(symeval_1kl), let__4755)
}


panic("unreachable")
})

tmp5236 := Call2(__e, ns2_1set, symshen_4input_1h_7, tmp5226)


_ = tmp5236

tmp5237 := MakeNative1(func(__e *ControlFlow, V2591 Obj) Obj {
tmp5244 := GConsP(__e, V2591)

if True == tmp5244 {
tmp5238 := MakeNative1(func(__e *ControlFlow, Z2592 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4monotype), Z2592)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp5238, V2591)


} else {
tmp5242 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V2591)

if True == tmp5242 {
tmp5239 := Call3(__e, SymFn(symshen_4app), V2591, MakeString("\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("input+ expects a monotype: not "), tmp5239))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("input+ expects a monotype: not "), tmp5239)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


} else {
return V2591
}


}


panic("unreachable")
})

tmp5245 := Call2(__e, ns2_1set, symshen_4monotype, tmp5237)


_ = tmp5245

tmp5246 := MakeNative1(func(__e *ControlFlow, V2593 Obj) Obj {
tmp5247 := Call1(__e, SymFn(symshen_4my_1read_1byte), V2593)


tmp5248 := MakeNative1(func(__e *ControlFlow, Z2594 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4return_2), Z2594)
panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4read_1loop), V2593, tmp5247, Nil, tmp5248)


panic("unreachable")
})

tmp5249 := Call2(__e, ns2_1set, symlineread, tmp5246)


_ = tmp5249

tmp5250 := MakeNative1(func(__e *ControlFlow, V2595 Obj) Obj {
tmp5251 := Call1(__e, SymFn(symshen_4str_1_6bytes), V2595)


let__4756 := tmp5251
_ = let__4756

tmp5252 := MakeNative1(func(__e *ControlFlow, Z2598 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5s_1exprs_6), Z2598)
panic("unreachable")
})

tmp5253 := Call2(__e, SymFn(symcompile), tmp5252, let__4756)


let__4757 := tmp5253
_ = let__4757

tmp5254 := Call1(__e, SymFn(symshen_4process_1sexprs), let__4757)


let__4758 := tmp5254
_ = let__4758

return let__4758


panic("unreachable")
})

tmp5255 := Call2(__e, ns2_1set, symread_1from_1string, tmp5250)


_ = tmp5255

tmp5256 := MakeNative1(func(__e *ControlFlow, V2600 Obj) Obj {
tmp5257 := Call1(__e, SymFn(symshen_4str_1_6bytes), V2600)


let__4759 := tmp5257
_ = let__4759

tmp5258 := MakeNative1(func(__e *ControlFlow, Z2603 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5s_1exprs_6), Z2603)
panic("unreachable")
})

tmp5259 := Call2(__e, SymFn(symcompile), tmp5258, let__4759)


let__4760 := tmp5259
_ = let__4760

return let__4760


panic("unreachable")
})

tmp5260 := Call2(__e, ns2_1set, symread_1from_1string_1unprocessed, tmp5256)


_ = tmp5260

tmp5261 := MakeNative1(func(__e *ControlFlow, V2604 Obj) Obj {
tmp5269 := GEq(__e, MakeString(""), V2604)

if True == tmp5269 {
return Nil
} else {
tmp5267 := Call1(__e, SymFn(symshen_4_7string_2), V2604)


if True == tmp5267 {
tmp5262 := Call1(__e, SymFn(symhdstr), V2604)


tmp5263 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp5262)

tmp5265 := Call1(__e, SymFn(symshen_4str_1_6bytes), PrimGuard1(__e, symtlstr, PrimTailString, V2604))


return GCons(__e, tmp5263, tmp5265)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4str_1_6bytes)
}


}


panic("unreachable")
})

tmp5270 := Call2(__e, ns2_1set, symshen_4str_1_6bytes, tmp5261)


_ = tmp5270

tmp5271 := MakeNative1(func(__e *ControlFlow, V2605 Obj) Obj {
tmp5272 := Call1(__e, SymFn(symshen_4my_1read_1byte), V2605)


tmp5273 := MakeNative1(func(__e *ControlFlow, Z2606 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4whitespace_2), Z2606)
panic("unreachable")
})

tmp5274 := Call4(__e, SymFn(symshen_4read_1loop), V2605, tmp5272, Nil, tmp5273)


return GHd(__e, tmp5274)


panic("unreachable")
})

tmp5275 := Call2(__e, ns2_1set, symread, tmp5271)


_ = tmp5275

tmp5276 := MakeNative1(func(__e *ControlFlow, V2607 Obj) Obj {
tmp5279 := Call1(__e, SymFn(symchar_1stream_2), V2607)


if True == tmp5279 {
tmp5277 := Call1(__e, SymFn(symread_1unit_1string), V2607)


return PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp5277)


} else {
return PrimGuard1(__e, symread_1byte, PrimReadByte, V2607)
}


panic("unreachable")
})

tmp5280 := Call2(__e, ns2_1set, symshen_4my_1read_1byte, tmp5276)


_ = tmp5280

tmp5281 := (func() Obj {
var __self Obj
__self = MakeNative4(func(__e *ControlFlow, __self1, __self2, __self3, __self4 Obj) Obj {
__selftop:
V2612 := __self1
_ = V2612
V2613 := __self2
_ = V2613
V2614 := __self3
_ = V2614
V2615 := __self4
_ = V2615
tmp5304 := GEq(__e, MakeInteger(94), V2613)

if True == tmp5304 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("read aborted"))
} else {
tmp5302 := GEq(__e, MakeInteger(-1), V2613)

if True == tmp5302 {
tmp5285 := Call1(__e, SymFn(symempty_2), V2614)


if True == tmp5285 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("error: empty stream"))
} else {
tmp5282 := MakeNative1(func(__e *ControlFlow, Z2616 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5s_1exprs_6), Z2616)
panic("unreachable")
})

tmp5283 := Call2(__e, SymFn(symcompile), tmp5282, V2614)


return TailCall1(__e, SymFn(symshen_4process_1sexprs), tmp5283)


}


} else {
tmp5300 := GEq(__e, MakeInteger(0), V2613)

if True == tmp5300 {
tmp5286 := Call1(__e, SymFn(symshen_4my_1read_1byte), V2612)


if SymFn(symshen_4read_1loop) == __self {
__self1, __self2, __self3, __self4 = V2612, tmp5286, V2614, V2615
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4read_1loop), V2612, tmp5286, V2614, V2615)


} else {
tmp5298 := Call1(__e, V2615, V2613)


if True == tmp5298 {
tmp5287 := Call1(__e, SymFn(symshen_4try_1parse), V2614)


let__4761 := tmp5287
_ = let__4761

tmp5293 := Call1(__e, SymFn(symshen_4nothing_1doing_2), let__4761)


if True == tmp5293 {
tmp5288 := Call1(__e, SymFn(symshen_4my_1read_1byte), V2612)


tmp5289 := GCons(__e, V2613, Nil)

tmp5290 := Call2(__e, SymFn(symappend), V2614, tmp5289)


if SymFn(symshen_4read_1loop) == __self {
__self1, __self2, __self3, __self4 = V2612, tmp5288, tmp5290, V2615
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4read_1loop), V2612, tmp5288, tmp5290, V2615)


} else {
tmp5291 := Call1(__e, SymFn(symshen_4record_1it), V2614)


_ = tmp5291

return TailCall1(__e, SymFn(symshen_4process_1sexprs), let__4761)


}


} else {
tmp5294 := Call1(__e, SymFn(symshen_4my_1read_1byte), V2612)


tmp5295 := GCons(__e, V2613, Nil)

tmp5296 := Call2(__e, SymFn(symappend), V2614, tmp5295)


if SymFn(symshen_4read_1loop) == __self {
__self1, __self2, __self3, __self4 = V2612, tmp5294, tmp5296, V2615
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4read_1loop), V2612, tmp5294, tmp5296, V2615)


}


}


}


}


panic("unreachable")
})
return __self
})()

tmp5305 := Call2(__e, ns2_1set, symshen_4read_1loop, tmp5281)


_ = tmp5305

tmp5306 := MakeNative1(func(__e *ControlFlow, V2618 Obj) Obj {
tmp5307 := MakeNative0(func(__e *ControlFlow) Obj {
tmp5308 := MakeNative1(func(__e *ControlFlow, Z2619 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5s_1exprs_6), Z2619)
panic("unreachable")
})

return TailCall2(__e, SymFn(symcompile), tmp5308, V2618)


panic("unreachable")
})

tmp5309 := MakeNative1(func(__e *ControlFlow, Z2620 Obj) Obj {
return symshen_4i_1failed_b
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp5307, tmp5309)


panic("unreachable")
})

tmp5310 := Call2(__e, ns2_1set, symshen_4try_1parse, tmp5306)


_ = tmp5310

tmp5311 := MakeNative1(func(__e *ControlFlow, V2623 Obj) Obj {
tmp5315 := GEq(__e, symshen_4i_1failed_b, V2623)

if True == tmp5315 {
return True
} else {
tmp5313 := GEq(__e, Nil, V2623)

if True == tmp5313 {
return True
} else {
return False
}


}


panic("unreachable")
})

tmp5316 := Call2(__e, ns2_1set, symshen_4nothing_1doing_2, tmp5311)


_ = tmp5316

tmp5317 := MakeNative1(func(__e *ControlFlow, V2624 Obj) Obj {
tmp5318 := Call1(__e, SymFn(symshen_4bytes_1_6string), V2624)


return PrimGuard2(__e, symset, PrimSet, symshen_4_dit_d, tmp5318)


panic("unreachable")
})

tmp5319 := Call2(__e, ns2_1set, symshen_4record_1it, tmp5317)


_ = tmp5319

tmp5320 := MakeNative1(func(__e *ControlFlow, V2625 Obj) Obj {
tmp5328 := GEq(__e, Nil, V2625)

if True == tmp5328 {
return MakeString("")
} else {
tmp5326 := GConsP(__e, V2625)

if True == tmp5326 {
tmp5321 := GHd(__e, V2625)

tmp5322 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, tmp5321)

tmp5323 := GTl(__e, V2625)

tmp5324 := Call1(__e, SymFn(symshen_4bytes_1_6string), tmp5323)


return PrimGuard2(__e, symcn, PrimStringConcat, tmp5322, tmp5324)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4bytes_1_6string)
}


}


panic("unreachable")
})

tmp5329 := Call2(__e, ns2_1set, symshen_4bytes_1_6string, tmp5320)


_ = tmp5329

tmp5330 := MakeNative1(func(__e *ControlFlow, V2626 Obj) Obj {
tmp5331 := Call1(__e, SymFn(symshen_4unpackage_emacroexpand), V2626)


let__4762 := tmp5331
_ = let__4762

tmp5332 := Call1(__e, SymFn(symshen_4find_1arities), let__4762)


let__4763 := tmp5332
_ = let__4763

tmp5333 := Call1(__e, SymFn(symshen_4find_1types), let__4762)


let__4764 := tmp5333
_ = let__4764

tmp5334 := MakeNative1(func(__e *ControlFlow, Z2630 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4process_1applications), Z2630, let__4764)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp5334, let__4762)


panic("unreachable")
})

tmp5335 := Call2(__e, ns2_1set, symshen_4process_1sexprs, tmp5330)


_ = tmp5335

tmp5336 := MakeNative1(func(__e *ControlFlow, V2631 Obj) Obj {
tmp5356 := GConsP(__e, V2631)

var ifres5349 Obj

if True == tmp5356 {
tmp5354 := GTl(__e, V2631)

tmp5355 := GConsP(__e, tmp5354)

var ifres5350 Obj

if True == tmp5355 {
tmp5351 := GHd(__e, V2631)

tmp5352 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp5353 := GEq(__e, tmp5351, tmp5352)

ifres5350 = tmp5353


} else {
ifres5350 = False


}

ifres5349 = ifres5350


} else {
ifres5349 = False


}

if True == ifres5349 {
tmp5337 := GTl(__e, V2631)

tmp5338 := GHd(__e, tmp5337)

tmp5339 := GTl(__e, V2631)

tmp5340 := GTl(__e, tmp5339)

tmp5341 := Call1(__e, SymFn(symshen_4find_1types), tmp5340)


return GCons(__e, tmp5338, tmp5341)


} else {
tmp5347 := GConsP(__e, V2631)

if True == tmp5347 {
tmp5342 := GHd(__e, V2631)

tmp5343 := Call1(__e, SymFn(symshen_4find_1types), tmp5342)


tmp5344 := GTl(__e, V2631)

tmp5345 := Call1(__e, SymFn(symshen_4find_1types), tmp5344)


return TailCall2(__e, SymFn(symappend), tmp5343, tmp5345)


} else {
return Nil
}


}


panic("unreachable")
})

tmp5357 := Call2(__e, ns2_1set, symshen_4find_1types, tmp5336)


_ = tmp5357

tmp5358 := MakeNative1(func(__e *ControlFlow, V2634 Obj) Obj {
tmp5401 := GConsP(__e, V2634)

var ifres5386 Obj

if True == tmp5401 {
tmp5399 := GHd(__e, V2634)

tmp5400 := GEq(__e, symdefine, tmp5399)

var ifres5387 Obj

if True == tmp5400 {
tmp5397 := GTl(__e, V2634)

tmp5398 := GConsP(__e, tmp5397)

var ifres5388 Obj

if True == tmp5398 {
tmp5394 := GTl(__e, V2634)

tmp5395 := GTl(__e, tmp5394)

tmp5396 := GConsP(__e, tmp5395)

var ifres5389 Obj

if True == tmp5396 {
tmp5390 := GTl(__e, V2634)

tmp5391 := GTl(__e, tmp5390)

tmp5392 := GHd(__e, tmp5391)

tmp5393 := GEq(__e, sym_i, tmp5392)

ifres5389 = tmp5393


} else {
ifres5389 = False


}

ifres5388 = ifres5389


} else {
ifres5388 = False


}

ifres5387 = ifres5388


} else {
ifres5387 = False


}

ifres5386 = ifres5387


} else {
ifres5386 = False


}

if True == ifres5386 {
tmp5359 := GTl(__e, V2634)

tmp5360 := GHd(__e, tmp5359)

tmp5361 := GTl(__e, V2634)

tmp5362 := GHd(__e, tmp5361)

tmp5363 := GTl(__e, V2634)

tmp5364 := GTl(__e, tmp5363)

tmp5365 := GTl(__e, tmp5364)

tmp5366 := Call3(__e, SymFn(symshen_4find_1arity), tmp5362, MakeInteger(1), tmp5365)


return TailCall2(__e, SymFn(symshen_4store_1arity), tmp5360, tmp5366)


} else {
tmp5384 := GConsP(__e, V2634)

var ifres5378 Obj

if True == tmp5384 {
tmp5382 := GHd(__e, V2634)

tmp5383 := GEq(__e, symdefine, tmp5382)

var ifres5379 Obj

if True == tmp5383 {
tmp5380 := GTl(__e, V2634)

tmp5381 := GConsP(__e, tmp5380)

ifres5379 = tmp5381


} else {
ifres5379 = False


}

ifres5378 = ifres5379


} else {
ifres5378 = False


}

if True == ifres5378 {
tmp5367 := GTl(__e, V2634)

tmp5368 := GHd(__e, tmp5367)

tmp5369 := GTl(__e, V2634)

tmp5370 := GHd(__e, tmp5369)

tmp5371 := GTl(__e, V2634)

tmp5372 := GTl(__e, tmp5371)

tmp5373 := Call3(__e, SymFn(symshen_4find_1arity), tmp5370, MakeInteger(0), tmp5372)


return TailCall2(__e, SymFn(symshen_4store_1arity), tmp5368, tmp5373)


} else {
tmp5376 := GConsP(__e, V2634)

if True == tmp5376 {
tmp5374 := MakeNative1(func(__e *ControlFlow, Z2635 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4find_1arities), Z2635)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp5374, V2634)


} else {
return symshen_4skip
}


}


}


panic("unreachable")
})

tmp5402 := Call2(__e, ns2_1set, symshen_4find_1arities, tmp5358)


_ = tmp5402

tmp5403 := MakeNative2(func(__e *ControlFlow, V2636, V2637 Obj) Obj {
tmp5404 := Call1(__e, SymFn(symarity), V2636)


let__4765 := tmp5404
_ = let__4765

tmp5415 := GEq(__e, let__4765, MakeInteger(-1))

if True == tmp5415 {
return TailCall2(__e, SymFn(symshen_4execute_1store_1arity), V2636, V2637)
} else {
tmp5413 := GEq(__e, let__4765, V2637)

if True == tmp5413 {
return symshen_4skip
} else {
tmp5411 := Call1(__e, SymFn(symshen_4sysfunc_2), V2636)


if True == tmp5411 {
tmp5405 := Call3(__e, SymFn(symshen_4app), V2636, MakeString(" is a system function\n"), symshen_4a)


return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp5405)


} else {
tmp5406 := Call3(__e, SymFn(symshen_4app), V2636, MakeString(" may cause errors\n"), symshen_4a)


tmp5407 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("changing the arity of "), tmp5406)

tmp5408 := Call0(__e, SymFn(symstoutput))


tmp5409 := Call2(__e, SymFn(sympr), tmp5407, tmp5408)


_ = tmp5409

return TailCall2(__e, SymFn(symshen_4execute_1store_1arity), V2636, V2637)


}


}


}


panic("unreachable")
})

tmp5416 := Call2(__e, ns2_1set, symshen_4store_1arity, tmp5403)


_ = tmp5416

tmp5417 := MakeNative2(func(__e *ControlFlow, V2639, V2640 Obj) Obj {
tmp5422 := GEq(__e, MakeInteger(0), V2640)

if True == tmp5422 {
tmp5418 := GValue(__e, sym_dproperty_1vector_d)

return TailCall4(__e, SymFn(symput), V2639, symarity, MakeInteger(0), tmp5418)


} else {
tmp5419 := GValue(__e, sym_dproperty_1vector_d)

tmp5420 := Call4(__e, SymFn(symput), V2639, symarity, V2640, tmp5419)


_ = tmp5420

return TailCall2(__e, SymFn(symshen_4update_1lambdatable), V2639, V2640)


}


panic("unreachable")
})

tmp5423 := Call2(__e, ns2_1set, symshen_4execute_1store_1arity, tmp5417)


_ = tmp5423

tmp5424 := MakeNative2(func(__e *ControlFlow, V2641, V2642 Obj) Obj {
tmp5425 := GCons(__e, V2641, Nil)

tmp5426 := Call2(__e, SymFn(symshen_4lambda_1function), tmp5425, V2642)


tmp5427 := Call1(__e, SymFn(symeval_1kl), tmp5426)


let__4766 := tmp5427
_ = let__4766

tmp5428 := GCons(__e, V2641, let__4766)

tmp5429 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp5428)


let__4767 := tmp5429
_ = let__4767

return V2642


panic("unreachable")
})

tmp5430 := Call2(__e, ns2_1set, symshen_4update_1lambdatable, tmp5424)


_ = tmp5430

tmp5431 := MakeNative2(func(__e *ControlFlow, V2647, V2648 Obj) Obj {
tmp5447 := GEq(__e, MakeInteger(0), V2648)

if True == tmp5447 {
return symshen_4skip
} else {
tmp5445 := GEq(__e, MakeInteger(1), V2648)

if True == tmp5445 {
tmp5432 := Call1(__e, SymFn(symgensym), symY)


let__4768 := tmp5432
_ = let__4768

tmp5433 := GCons(__e, let__4768, Nil)

tmp5434 := Call2(__e, SymFn(symappend), V2647, tmp5433)


tmp5435 := GCons(__e, tmp5434, Nil)

tmp5436 := GCons(__e, let__4768, tmp5435)

return GCons(__e, symlambda, tmp5436)


} else {
tmp5437 := Call1(__e, SymFn(symgensym), symY)


let__4769 := tmp5437
_ = let__4769

tmp5438 := GCons(__e, let__4769, Nil)

tmp5439 := Call2(__e, SymFn(symappend), V2647, tmp5438)


tmp5441 := Call2(__e, SymFn(symshen_4lambda_1function), tmp5439, GSub(__e, V2648, MakeInteger(1)))


tmp5442 := GCons(__e, tmp5441, Nil)

tmp5443 := GCons(__e, let__4769, tmp5442)

return GCons(__e, symlambda, tmp5443)


}


}


panic("unreachable")
})

tmp5448 := Call2(__e, ns2_1set, symshen_4lambda_1function, tmp5431)


_ = tmp5448

tmp5449 := MakeNative3(func(__e *ControlFlow, V2660, V2661, V2662 Obj) Obj {
tmp5470 := GEq(__e, Nil, V2662)

if True == tmp5470 {
tmp5450 := GCons(__e, V2660, V2661)

return GCons(__e, tmp5450, Nil)


} else {
tmp5468 := GConsP(__e, V2662)

var ifres5461 Obj

if True == tmp5468 {
tmp5466 := GHd(__e, V2662)

tmp5467 := GConsP(__e, tmp5466)

var ifres5462 Obj

if True == tmp5467 {
tmp5463 := GHd(__e, V2662)

tmp5464 := GHd(__e, tmp5463)

tmp5465 := GEq(__e, V2660, tmp5464)

ifres5462 = tmp5465


} else {
ifres5462 = False


}

ifres5461 = ifres5462


} else {
ifres5461 = False


}

if True == ifres5461 {
tmp5451 := GHd(__e, V2662)

tmp5452 := GHd(__e, tmp5451)

tmp5453 := GCons(__e, tmp5452, V2661)

tmp5454 := GTl(__e, V2662)

return GCons(__e, tmp5453, tmp5454)


} else {
tmp5459 := GConsP(__e, V2662)

if True == tmp5459 {
tmp5455 := GHd(__e, V2662)

tmp5456 := GTl(__e, V2662)

tmp5457 := Call3(__e, SymFn(symshen_4assoc_1_6), V2660, V2661, tmp5456)


return GCons(__e, tmp5455, tmp5457)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.assoc->"))
}


}


}


panic("unreachable")
})

tmp5471 := Call2(__e, ns2_1set, symshen_4assoc_1_6, tmp5449)


_ = tmp5471

tmp5472 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V2677 := __self1
_ = V2677
V2678 := __self2
_ = V2678
V2679 := __self3
_ = V2679
tmp5511 := GEq(__e, MakeInteger(0), V2678)

var ifres5506 Obj

if True == tmp5511 {
tmp5510 := GConsP(__e, V2679)

var ifres5507 Obj

if True == tmp5510 {
tmp5508 := GHd(__e, V2679)

tmp5509 := GEq(__e, tmp5508, sym_1_6)

ifres5507 = tmp5509


} else {
ifres5507 = False


}

ifres5506 = ifres5507


} else {
ifres5506 = False


}

if True == ifres5506 {
return MakeInteger(0)
} else {
tmp5504 := GEq(__e, MakeInteger(0), V2678)

var ifres5499 Obj

if True == tmp5504 {
tmp5503 := GConsP(__e, V2679)

var ifres5500 Obj

if True == tmp5503 {
tmp5501 := GHd(__e, V2679)

tmp5502 := GEq(__e, tmp5501, sym_5_1)

ifres5500 = tmp5502


} else {
ifres5500 = False


}

ifres5499 = ifres5500


} else {
ifres5499 = False


}

if True == ifres5499 {
return MakeInteger(0)
} else {
tmp5497 := GEq(__e, MakeInteger(0), V2678)

var ifres5495 Obj

if True == tmp5497 {
tmp5496 := GConsP(__e, V2679)

ifres5495 = tmp5496


} else {
ifres5495 = False


}

if True == ifres5495 {
tmp5473 := GTl(__e, V2679)

tmp5474 := Call3(__e, SymFn(symshen_4find_1arity), V2677, MakeInteger(0), tmp5473)


return GAdd(__e, MakeInteger(1), tmp5474)


} else {
tmp5493 := GEq(__e, MakeInteger(1), V2678)

var ifres5488 Obj

if True == tmp5493 {
tmp5492 := GConsP(__e, V2679)

var ifres5489 Obj

if True == tmp5492 {
tmp5490 := GHd(__e, V2679)

tmp5491 := GEq(__e, sym_j, tmp5490)

ifres5489 = tmp5491


} else {
ifres5489 = False


}

ifres5488 = ifres5489


} else {
ifres5488 = False


}

if True == ifres5488 {
tmp5475 := GTl(__e, V2679)

if SymFn(symshen_4find_1arity) == __self {
__self1, __self2, __self3 = V2677, MakeInteger(0), tmp5475
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4find_1arity), V2677, MakeInteger(0), tmp5475)


} else {
tmp5486 := GEq(__e, MakeInteger(1), V2678)

var ifres5484 Obj

if True == tmp5486 {
tmp5485 := GConsP(__e, V2679)

ifres5484 = tmp5485


} else {
ifres5484 = False


}

if True == ifres5484 {
tmp5476 := GTl(__e, V2679)

if SymFn(symshen_4find_1arity) == __self {
__self1, __self2, __self3 = V2677, MakeInteger(1), tmp5476
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4find_1arity), V2677, MakeInteger(1), tmp5476)


} else {
tmp5482 := GEq(__e, MakeInteger(1), V2678)

if True == tmp5482 {
tmp5477 := Call3(__e, SymFn(symshen_4app), V2677, MakeString(" definition: missing }\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("syntax error in "), tmp5477))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("syntax error in "), tmp5477)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


} else {
tmp5479 := Call3(__e, SymFn(symshen_4app), V2677, MakeString(" definition: missing -> or <-\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("syntax error in "), tmp5479))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("syntax error in "), tmp5479)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


}


}


}


}


}


panic("unreachable")
})
return __self
})()

tmp5512 := Call2(__e, ns2_1set, symshen_4find_1arity, tmp5472)


_ = tmp5512

tmp5513 := MakeNative1(func(__e *ControlFlow, V2680 Obj) Obj {
tmp5514 := Call1(__e, SymFn(symshen_4_5lsb_6), V2680)


let__4771 := tmp5514
_ = let__4771

tmp5538 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4771)


var ifres5515 Obj

if True == tmp5538 {
tmp5516 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5515 = tmp5516


} else {
tmp5517 := Call1(__e, SymFn(symshen_4in_1_6), let__4771)


let__4772 := tmp5517
_ = let__4772

tmp5518 := Call1(__e, SymFn(symshen_4_5s_1exprs1_6), let__4772)


let__4773 := tmp5518
_ = let__4773

tmp5537 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4773)


var ifres5519 Obj

if True == tmp5537 {
tmp5520 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5519 = tmp5520


} else {
tmp5521 := Call1(__e, SymFn(symshen_4_5_1out), let__4773)


let__4774 := tmp5521
_ = let__4774

tmp5522 := Call1(__e, SymFn(symshen_4in_1_6), let__4773)


let__4775 := tmp5522
_ = let__4775

tmp5523 := Call1(__e, SymFn(symshen_4_5rsb_6), let__4775)


let__4776 := tmp5523
_ = let__4776

tmp5536 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4776)


var ifres5524 Obj

if True == tmp5536 {
tmp5525 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5524 = tmp5525


} else {
tmp5526 := Call1(__e, SymFn(symshen_4in_1_6), let__4776)


let__4777 := tmp5526
_ = let__4777

tmp5527 := Call1(__e, SymFn(symshen_4_5s_1exprs2_6), let__4777)


let__4778 := tmp5527
_ = let__4778

tmp5535 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4778)


var ifres5528 Obj

if True == tmp5535 {
tmp5529 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5528 = tmp5529


} else {
tmp5530 := Call1(__e, SymFn(symshen_4_5_1out), let__4778)


let__4779 := tmp5530
_ = let__4779

tmp5531 := Call1(__e, SymFn(symshen_4in_1_6), let__4778)


let__4780 := tmp5531
_ = let__4780

tmp5532 := Call1(__e, SymFn(symshen_4cons_1form), let__4774)


tmp5533 := GCons(__e, tmp5532, let__4779)

tmp5534 := Call2(__e, SymFn(symshen_4comb), let__4780, tmp5533)


ifres5528 = tmp5534


}

ifres5524 = ifres5528


}

ifres5519 = ifres5524


}

ifres5515 = ifres5519


}

let__4770 := ifres5515
_ = let__4770

tmp5732 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4770)


if True == tmp5732 {
tmp5539 := Call1(__e, SymFn(symshen_4_5lrb_6), V2680)


let__4782 := tmp5539
_ = let__4782

tmp5562 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4782)


var ifres5540 Obj

if True == tmp5562 {
tmp5541 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5540 = tmp5541


} else {
tmp5542 := Call1(__e, SymFn(symshen_4in_1_6), let__4782)


let__4783 := tmp5542
_ = let__4783

tmp5543 := Call1(__e, SymFn(symshen_4_5s_1exprs1_6), let__4783)


let__4784 := tmp5543
_ = let__4784

tmp5561 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4784)


var ifres5544 Obj

if True == tmp5561 {
tmp5545 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5544 = tmp5545


} else {
tmp5546 := Call1(__e, SymFn(symshen_4_5_1out), let__4784)


let__4785 := tmp5546
_ = let__4785

tmp5547 := Call1(__e, SymFn(symshen_4in_1_6), let__4784)


let__4786 := tmp5547
_ = let__4786

tmp5548 := Call1(__e, SymFn(symshen_4_5rrb_6), let__4786)


let__4787 := tmp5548
_ = let__4787

tmp5560 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4787)


var ifres5549 Obj

if True == tmp5560 {
tmp5550 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5549 = tmp5550


} else {
tmp5551 := Call1(__e, SymFn(symshen_4in_1_6), let__4787)


let__4788 := tmp5551
_ = let__4788

tmp5552 := Call1(__e, SymFn(symshen_4_5s_1exprs2_6), let__4788)


let__4789 := tmp5552
_ = let__4789

tmp5559 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4789)


var ifres5553 Obj

if True == tmp5559 {
tmp5554 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5553 = tmp5554


} else {
tmp5555 := Call1(__e, SymFn(symshen_4_5_1out), let__4789)


let__4790 := tmp5555
_ = let__4790

tmp5556 := Call1(__e, SymFn(symshen_4in_1_6), let__4789)


let__4791 := tmp5556
_ = let__4791

tmp5557 := Call2(__e, SymFn(symshen_4add_1sexpr), let__4785, let__4790)


tmp5558 := Call2(__e, SymFn(symshen_4comb), let__4791, tmp5557)


ifres5553 = tmp5558


}

ifres5549 = ifres5553


}

ifres5544 = ifres5549


}

ifres5540 = ifres5544


}

let__4781 := ifres5540
_ = let__4781

tmp5730 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4781)


if True == tmp5730 {
tmp5563 := Call1(__e, SymFn(symshen_4_5lcurly_6), V2680)


let__4793 := tmp5563
_ = let__4793

tmp5575 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4793)


var ifres5564 Obj

if True == tmp5575 {
tmp5565 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5564 = tmp5565


} else {
tmp5566 := Call1(__e, SymFn(symshen_4in_1_6), let__4793)


let__4794 := tmp5566
_ = let__4794

tmp5567 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__4794)


let__4795 := tmp5567
_ = let__4795

tmp5574 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4795)


var ifres5568 Obj

if True == tmp5574 {
tmp5569 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5568 = tmp5569


} else {
tmp5570 := Call1(__e, SymFn(symshen_4_5_1out), let__4795)


let__4796 := tmp5570
_ = let__4796

tmp5571 := Call1(__e, SymFn(symshen_4in_1_6), let__4795)


let__4797 := tmp5571
_ = let__4797

tmp5572 := GCons(__e, sym_i, let__4796)

tmp5573 := Call2(__e, SymFn(symshen_4comb), let__4797, tmp5572)


ifres5568 = tmp5573


}

ifres5564 = ifres5568


}

let__4792 := ifres5564
_ = let__4792

tmp5728 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4792)


if True == tmp5728 {
tmp5576 := Call1(__e, SymFn(symshen_4_5rcurly_6), V2680)


let__4799 := tmp5576
_ = let__4799

tmp5588 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4799)


var ifres5577 Obj

if True == tmp5588 {
tmp5578 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5577 = tmp5578


} else {
tmp5579 := Call1(__e, SymFn(symshen_4in_1_6), let__4799)


let__4800 := tmp5579
_ = let__4800

tmp5580 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__4800)


let__4801 := tmp5580
_ = let__4801

tmp5587 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4801)


var ifres5581 Obj

if True == tmp5587 {
tmp5582 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5581 = tmp5582


} else {
tmp5583 := Call1(__e, SymFn(symshen_4_5_1out), let__4801)


let__4802 := tmp5583
_ = let__4802

tmp5584 := Call1(__e, SymFn(symshen_4in_1_6), let__4801)


let__4803 := tmp5584
_ = let__4803

tmp5585 := GCons(__e, sym_j, let__4802)

tmp5586 := Call2(__e, SymFn(symshen_4comb), let__4803, tmp5585)


ifres5581 = tmp5586


}

ifres5577 = ifres5581


}

let__4798 := ifres5577
_ = let__4798

tmp5726 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4798)


if True == tmp5726 {
tmp5589 := Call1(__e, SymFn(symshen_4_5bar_6), V2680)


let__4805 := tmp5589
_ = let__4805

tmp5601 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4805)


var ifres5590 Obj

if True == tmp5601 {
tmp5591 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5590 = tmp5591


} else {
tmp5592 := Call1(__e, SymFn(symshen_4in_1_6), let__4805)


let__4806 := tmp5592
_ = let__4806

tmp5593 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__4806)


let__4807 := tmp5593
_ = let__4807

tmp5600 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4807)


var ifres5594 Obj

if True == tmp5600 {
tmp5595 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5594 = tmp5595


} else {
tmp5596 := Call1(__e, SymFn(symshen_4_5_1out), let__4807)


let__4808 := tmp5596
_ = let__4808

tmp5597 := Call1(__e, SymFn(symshen_4in_1_6), let__4807)


let__4809 := tmp5597
_ = let__4809

tmp5598 := GCons(__e, symbar_b, let__4808)

tmp5599 := Call2(__e, SymFn(symshen_4comb), let__4809, tmp5598)


ifres5594 = tmp5599


}

ifres5590 = ifres5594


}

let__4804 := ifres5590
_ = let__4804

tmp5724 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4804)


if True == tmp5724 {
tmp5602 := Call1(__e, SymFn(symshen_4_5semicolon_6), V2680)


let__4811 := tmp5602
_ = let__4811

tmp5615 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4811)


var ifres5603 Obj

if True == tmp5615 {
tmp5604 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5603 = tmp5604


} else {
tmp5605 := Call1(__e, SymFn(symshen_4in_1_6), let__4811)


let__4812 := tmp5605
_ = let__4812

tmp5606 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__4812)


let__4813 := tmp5606
_ = let__4813

tmp5614 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4813)


var ifres5607 Obj

if True == tmp5614 {
tmp5608 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5607 = tmp5608


} else {
tmp5609 := Call1(__e, SymFn(symshen_4_5_1out), let__4813)


let__4814 := tmp5609
_ = let__4814

tmp5610 := Call1(__e, SymFn(symshen_4in_1_6), let__4813)


let__4815 := tmp5610
_ = let__4815

tmp5611 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp5612 := GCons(__e, tmp5611, let__4814)

tmp5613 := Call2(__e, SymFn(symshen_4comb), let__4815, tmp5612)


ifres5607 = tmp5613


}

ifres5603 = ifres5607


}

let__4810 := ifres5603
_ = let__4810

tmp5722 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4810)


if True == tmp5722 {
tmp5616 := Call1(__e, SymFn(symshen_4_5colon_6), V2680)


let__4817 := tmp5616
_ = let__4817

tmp5634 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4817)


var ifres5617 Obj

if True == tmp5634 {
tmp5618 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5617 = tmp5618


} else {
tmp5619 := Call1(__e, SymFn(symshen_4in_1_6), let__4817)


let__4818 := tmp5619
_ = let__4818

tmp5620 := Call1(__e, SymFn(symshen_4_5equal_6), let__4818)


let__4819 := tmp5620
_ = let__4819

tmp5633 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4819)


var ifres5621 Obj

if True == tmp5633 {
tmp5622 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5621 = tmp5622


} else {
tmp5623 := Call1(__e, SymFn(symshen_4in_1_6), let__4819)


let__4820 := tmp5623
_ = let__4820

tmp5624 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__4820)


let__4821 := tmp5624
_ = let__4821

tmp5632 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4821)


var ifres5625 Obj

if True == tmp5632 {
tmp5626 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5625 = tmp5626


} else {
tmp5627 := Call1(__e, SymFn(symshen_4_5_1out), let__4821)


let__4822 := tmp5627
_ = let__4822

tmp5628 := Call1(__e, SymFn(symshen_4in_1_6), let__4821)


let__4823 := tmp5628
_ = let__4823

tmp5629 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":="))

tmp5630 := GCons(__e, tmp5629, let__4822)

tmp5631 := Call2(__e, SymFn(symshen_4comb), let__4823, tmp5630)


ifres5625 = tmp5631


}

ifres5621 = ifres5625


}

ifres5617 = ifres5621


}

let__4816 := ifres5617
_ = let__4816

tmp5720 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4816)


if True == tmp5720 {
tmp5635 := Call1(__e, SymFn(symshen_4_5colon_6), V2680)


let__4825 := tmp5635
_ = let__4825

tmp5648 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4825)


var ifres5636 Obj

if True == tmp5648 {
tmp5637 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5636 = tmp5637


} else {
tmp5638 := Call1(__e, SymFn(symshen_4in_1_6), let__4825)


let__4826 := tmp5638
_ = let__4826

tmp5639 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__4826)


let__4827 := tmp5639
_ = let__4827

tmp5647 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4827)


var ifres5640 Obj

if True == tmp5647 {
tmp5641 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5640 = tmp5641


} else {
tmp5642 := Call1(__e, SymFn(symshen_4_5_1out), let__4827)


let__4828 := tmp5642
_ = let__4828

tmp5643 := Call1(__e, SymFn(symshen_4in_1_6), let__4827)


let__4829 := tmp5643
_ = let__4829

tmp5644 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp5645 := GCons(__e, tmp5644, let__4828)

tmp5646 := Call2(__e, SymFn(symshen_4comb), let__4829, tmp5645)


ifres5640 = tmp5646


}

ifres5636 = ifres5640


}

let__4824 := ifres5636
_ = let__4824

tmp5718 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4824)


if True == tmp5718 {
tmp5649 := Call1(__e, SymFn(symshen_4_5comma_6), V2680)


let__4831 := tmp5649
_ = let__4831

tmp5662 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4831)


var ifres5650 Obj

if True == tmp5662 {
tmp5651 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5650 = tmp5651


} else {
tmp5652 := Call1(__e, SymFn(symshen_4in_1_6), let__4831)


let__4832 := tmp5652
_ = let__4832

tmp5653 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__4832)


let__4833 := tmp5653
_ = let__4833

tmp5661 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4833)


var ifres5654 Obj

if True == tmp5661 {
tmp5655 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5654 = tmp5655


} else {
tmp5656 := Call1(__e, SymFn(symshen_4_5_1out), let__4833)


let__4834 := tmp5656
_ = let__4834

tmp5657 := Call1(__e, SymFn(symshen_4in_1_6), let__4833)


let__4835 := tmp5657
_ = let__4835

tmp5658 := PrimGuard1(__e, symintern, PrimIntern, MakeString(","))

tmp5659 := GCons(__e, tmp5658, let__4834)

tmp5660 := Call2(__e, SymFn(symshen_4comb), let__4835, tmp5659)


ifres5654 = tmp5660


}

ifres5650 = ifres5654


}

let__4830 := ifres5650
_ = let__4830

tmp5716 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4830)


if True == tmp5716 {
tmp5663 := Call1(__e, SymFn(symshen_4_5comment_6), V2680)


let__4837 := tmp5663
_ = let__4837

tmp5674 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4837)


var ifres5664 Obj

if True == tmp5674 {
tmp5665 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5664 = tmp5665


} else {
tmp5666 := Call1(__e, SymFn(symshen_4in_1_6), let__4837)


let__4838 := tmp5666
_ = let__4838

tmp5667 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__4838)


let__4839 := tmp5667
_ = let__4839

tmp5673 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4839)


var ifres5668 Obj

if True == tmp5673 {
tmp5669 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5668 = tmp5669


} else {
tmp5670 := Call1(__e, SymFn(symshen_4_5_1out), let__4839)


let__4840 := tmp5670
_ = let__4840

tmp5671 := Call1(__e, SymFn(symshen_4in_1_6), let__4839)


let__4841 := tmp5671
_ = let__4841

tmp5672 := Call2(__e, SymFn(symshen_4comb), let__4841, let__4840)


ifres5668 = tmp5672


}

ifres5664 = ifres5668


}

let__4836 := ifres5664
_ = let__4836

tmp5714 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4836)


if True == tmp5714 {
tmp5675 := Call1(__e, SymFn(symshen_4_5atom_6), V2680)


let__4843 := tmp5675
_ = let__4843

tmp5688 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4843)


var ifres5676 Obj

if True == tmp5688 {
tmp5677 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5676 = tmp5677


} else {
tmp5678 := Call1(__e, SymFn(symshen_4_5_1out), let__4843)


let__4844 := tmp5678
_ = let__4844

tmp5679 := Call1(__e, SymFn(symshen_4in_1_6), let__4843)


let__4845 := tmp5679
_ = let__4845

tmp5680 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__4845)


let__4846 := tmp5680
_ = let__4846

tmp5687 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4846)


var ifres5681 Obj

if True == tmp5687 {
tmp5682 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5681 = tmp5682


} else {
tmp5683 := Call1(__e, SymFn(symshen_4_5_1out), let__4846)


let__4847 := tmp5683
_ = let__4847

tmp5684 := Call1(__e, SymFn(symshen_4in_1_6), let__4846)


let__4848 := tmp5684
_ = let__4848

tmp5685 := GCons(__e, let__4844, let__4847)

tmp5686 := Call2(__e, SymFn(symshen_4comb), let__4848, tmp5685)


ifres5681 = tmp5686


}

ifres5676 = ifres5681


}

let__4842 := ifres5676
_ = let__4842

tmp5712 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4842)


if True == tmp5712 {
tmp5689 := Call1(__e, SymFn(symshen_4_5whitespaces_6), V2680)


let__4850 := tmp5689
_ = let__4850

tmp5700 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4850)


var ifres5690 Obj

if True == tmp5700 {
tmp5691 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5690 = tmp5691


} else {
tmp5692 := Call1(__e, SymFn(symshen_4in_1_6), let__4850)


let__4851 := tmp5692
_ = let__4851

tmp5693 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__4851)


let__4852 := tmp5693
_ = let__4852

tmp5699 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4852)


var ifres5694 Obj

if True == tmp5699 {
tmp5695 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5694 = tmp5695


} else {
tmp5696 := Call1(__e, SymFn(symshen_4_5_1out), let__4852)


let__4853 := tmp5696
_ = let__4853

tmp5697 := Call1(__e, SymFn(symshen_4in_1_6), let__4852)


let__4854 := tmp5697
_ = let__4854

tmp5698 := Call2(__e, SymFn(symshen_4comb), let__4854, let__4853)


ifres5694 = tmp5698


}

ifres5690 = ifres5694


}

let__4849 := ifres5690
_ = let__4849

tmp5710 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4849)


if True == tmp5710 {
tmp5701 := Call1(__e, SymFn(sym_5e_6), V2680)


let__4856 := tmp5701
_ = let__4856

tmp5706 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4856)


var ifres5702 Obj

if True == tmp5706 {
tmp5703 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5702 = tmp5703


} else {
tmp5704 := Call1(__e, SymFn(symshen_4in_1_6), let__4856)


let__4857 := tmp5704
_ = let__4857

tmp5705 := Call2(__e, SymFn(symshen_4comb), let__4857, Nil)


ifres5702 = tmp5705


}

let__4855 := ifres5702
_ = let__4855

tmp5708 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4855)


if True == tmp5708 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4855
}


} else {
return let__4849
}


} else {
return let__4842
}


} else {
return let__4836
}


} else {
return let__4830
}


} else {
return let__4824
}


} else {
return let__4816
}


} else {
return let__4810
}


} else {
return let__4804
}


} else {
return let__4798
}


} else {
return let__4792
}


} else {
return let__4781
}


} else {
return let__4770
}


panic("unreachable")
})

tmp5733 := Call2(__e, ns2_1set, symshen_4_5s_1exprs_6, tmp5513)


_ = tmp5733

tmp5734 := MakeNative2(func(__e *ControlFlow, V2769, V2770 Obj) Obj {
tmp5749 := GConsP(__e, V2769)

var ifres5739 Obj

if True == tmp5749 {
tmp5747 := GHd(__e, V2769)

tmp5748 := GEq(__e, sym_3, tmp5747)

var ifres5740 Obj

if True == tmp5748 {
tmp5745 := GTl(__e, V2769)

tmp5746 := GConsP(__e, tmp5745)

var ifres5741 Obj

if True == tmp5746 {
tmp5742 := GTl(__e, V2769)

tmp5743 := GTl(__e, tmp5742)

tmp5744 := GEq(__e, Nil, tmp5743)

ifres5741 = tmp5744


} else {
ifres5741 = False


}

ifres5740 = ifres5741


} else {
ifres5740 = False


}

ifres5739 = ifres5740


} else {
ifres5739 = False


}

if True == ifres5739 {
tmp5735 := GTl(__e, V2769)

tmp5736 := GHd(__e, tmp5735)

tmp5737 := Call1(__e, SymFn(symexplode), tmp5736)


return TailCall2(__e, SymFn(symappend), tmp5737, V2770)


} else {
return GCons(__e, V2769, V2770)
}


panic("unreachable")
})

tmp5750 := Call2(__e, ns2_1set, symshen_4add_1sexpr, tmp5734)


_ = tmp5750

tmp5751 := MakeNative1(func(__e *ControlFlow, V2771 Obj) Obj {
tmp5756 := Call2(__e, SymFn(symshen_4hds_a_2), V2771, MakeInteger(91))


var ifres5752 Obj

if True == tmp5756 {
tmp5753 := Call1(__e, SymFn(symtail), V2771)


let__4859 := tmp5753
_ = let__4859

tmp5754 := Call2(__e, SymFn(symshen_4comb), let__4859, symshen_4skip)


ifres5752 = tmp5754


} else {
tmp5755 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5752 = tmp5755


}

let__4858 := ifres5752
_ = let__4858

tmp5758 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4858)


if True == tmp5758 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4858
}


panic("unreachable")
})

tmp5759 := Call2(__e, ns2_1set, symshen_4_5lsb_6, tmp5751)


_ = tmp5759

tmp5760 := MakeNative1(func(__e *ControlFlow, V2774 Obj) Obj {
tmp5765 := Call2(__e, SymFn(symshen_4hds_a_2), V2774, MakeInteger(93))


var ifres5761 Obj

if True == tmp5765 {
tmp5762 := Call1(__e, SymFn(symtail), V2774)


let__4861 := tmp5762
_ = let__4861

tmp5763 := Call2(__e, SymFn(symshen_4comb), let__4861, symshen_4skip)


ifres5761 = tmp5763


} else {
tmp5764 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5761 = tmp5764


}

let__4860 := ifres5761
_ = let__4860

tmp5767 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4860)


if True == tmp5767 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4860
}


panic("unreachable")
})

tmp5768 := Call2(__e, ns2_1set, symshen_4_5rsb_6, tmp5760)


_ = tmp5768

tmp5769 := MakeNative1(func(__e *ControlFlow, V2777 Obj) Obj {
tmp5770 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), V2777)


let__4863 := tmp5770
_ = let__4863

tmp5776 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4863)


var ifres5771 Obj

if True == tmp5776 {
tmp5772 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5771 = tmp5772


} else {
tmp5773 := Call1(__e, SymFn(symshen_4_5_1out), let__4863)


let__4864 := tmp5773
_ = let__4864

tmp5774 := Call1(__e, SymFn(symshen_4in_1_6), let__4863)


let__4865 := tmp5774
_ = let__4865

tmp5775 := Call2(__e, SymFn(symshen_4comb), let__4865, let__4864)


ifres5771 = tmp5775


}

let__4862 := ifres5771
_ = let__4862

tmp5778 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4862)


if True == tmp5778 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4862
}


panic("unreachable")
})

tmp5779 := Call2(__e, ns2_1set, symshen_4_5s_1exprs1_6, tmp5769)


_ = tmp5779

tmp5780 := MakeNative1(func(__e *ControlFlow, V2782 Obj) Obj {
tmp5781 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), V2782)


let__4867 := tmp5781
_ = let__4867

tmp5787 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4867)


var ifres5782 Obj

if True == tmp5787 {
tmp5783 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5782 = tmp5783


} else {
tmp5784 := Call1(__e, SymFn(symshen_4_5_1out), let__4867)


let__4868 := tmp5784
_ = let__4868

tmp5785 := Call1(__e, SymFn(symshen_4in_1_6), let__4867)


let__4869 := tmp5785
_ = let__4869

tmp5786 := Call2(__e, SymFn(symshen_4comb), let__4869, let__4868)


ifres5782 = tmp5786


}

let__4866 := ifres5782
_ = let__4866

tmp5789 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4866)


if True == tmp5789 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4866
}


panic("unreachable")
})

tmp5790 := Call2(__e, ns2_1set, symshen_4_5s_1exprs2_6, tmp5780)


_ = tmp5790

tmp5791 := MakeNative1(func(__e *ControlFlow, V2788 Obj) Obj {
tmp5840 := GEq(__e, Nil, V2788)

if True == tmp5840 {
return Nil
} else {
tmp5838 := GConsP(__e, V2788)

var ifres5822 Obj

if True == tmp5838 {
tmp5836 := GTl(__e, V2788)

tmp5837 := GConsP(__e, tmp5836)

var ifres5823 Obj

if True == tmp5837 {
tmp5833 := GTl(__e, V2788)

tmp5834 := GTl(__e, tmp5833)

tmp5835 := GConsP(__e, tmp5834)

var ifres5824 Obj

if True == tmp5835 {
tmp5829 := GTl(__e, V2788)

tmp5830 := GTl(__e, tmp5829)

tmp5831 := GTl(__e, tmp5830)

tmp5832 := GEq(__e, Nil, tmp5831)

var ifres5825 Obj

if True == tmp5832 {
tmp5826 := GTl(__e, V2788)

tmp5827 := GHd(__e, tmp5826)

tmp5828 := GEq(__e, tmp5827, symbar_b)

ifres5825 = tmp5828


} else {
ifres5825 = False


}

ifres5824 = ifres5825


} else {
ifres5824 = False


}

ifres5823 = ifres5824


} else {
ifres5823 = False


}

ifres5822 = ifres5823


} else {
ifres5822 = False


}

if True == ifres5822 {
tmp5792 := GHd(__e, V2788)

tmp5793 := GTl(__e, V2788)

tmp5794 := GTl(__e, tmp5793)

tmp5795 := GCons(__e, tmp5792, tmp5794)

return GCons(__e, symcons, tmp5795)


} else {
tmp5820 := GConsP(__e, V2788)

var ifres5804 Obj

if True == tmp5820 {
tmp5818 := GTl(__e, V2788)

tmp5819 := GConsP(__e, tmp5818)

var ifres5805 Obj

if True == tmp5819 {
tmp5815 := GTl(__e, V2788)

tmp5816 := GTl(__e, tmp5815)

tmp5817 := GConsP(__e, tmp5816)

var ifres5806 Obj

if True == tmp5817 {
tmp5811 := GTl(__e, V2788)

tmp5812 := GTl(__e, tmp5811)

tmp5813 := GTl(__e, tmp5812)

tmp5814 := GConsP(__e, tmp5813)

var ifres5807 Obj

if True == tmp5814 {
tmp5808 := GTl(__e, V2788)

tmp5809 := GHd(__e, tmp5808)

tmp5810 := GEq(__e, tmp5809, symbar_b)

ifres5807 = tmp5810


} else {
ifres5807 = False


}

ifres5806 = ifres5807


} else {
ifres5806 = False


}

ifres5805 = ifres5806


} else {
ifres5805 = False


}

ifres5804 = ifres5805


} else {
ifres5804 = False


}

if True == ifres5804 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("misapplication of |\n"))
} else {
tmp5802 := GConsP(__e, V2788)

if True == tmp5802 {
tmp5796 := GHd(__e, V2788)

tmp5797 := GTl(__e, V2788)

tmp5798 := Call1(__e, SymFn(symshen_4cons_1form), tmp5797)


tmp5799 := GCons(__e, tmp5798, Nil)

tmp5800 := GCons(__e, tmp5796, tmp5799)

return GCons(__e, symcons, tmp5800)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4cons_1form)
}


}


}


}


panic("unreachable")
})

tmp5841 := Call2(__e, ns2_1set, symshen_4cons_1form, tmp5791)


_ = tmp5841

tmp5842 := MakeNative1(func(__e *ControlFlow, V2789 Obj) Obj {
tmp5847 := Call2(__e, SymFn(symshen_4hds_a_2), V2789, MakeInteger(40))


var ifres5843 Obj

if True == tmp5847 {
tmp5844 := Call1(__e, SymFn(symtail), V2789)


let__4871 := tmp5844
_ = let__4871

tmp5845 := Call2(__e, SymFn(symshen_4comb), let__4871, symshen_4skip)


ifres5843 = tmp5845


} else {
tmp5846 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5843 = tmp5846


}

let__4870 := ifres5843
_ = let__4870

tmp5849 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4870)


if True == tmp5849 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4870
}


panic("unreachable")
})

tmp5850 := Call2(__e, ns2_1set, symshen_4_5lrb_6, tmp5842)


_ = tmp5850

tmp5851 := MakeNative1(func(__e *ControlFlow, V2792 Obj) Obj {
tmp5856 := Call2(__e, SymFn(symshen_4hds_a_2), V2792, MakeInteger(41))


var ifres5852 Obj

if True == tmp5856 {
tmp5853 := Call1(__e, SymFn(symtail), V2792)


let__4873 := tmp5853
_ = let__4873

tmp5854 := Call2(__e, SymFn(symshen_4comb), let__4873, symshen_4skip)


ifres5852 = tmp5854


} else {
tmp5855 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5852 = tmp5855


}

let__4872 := ifres5852
_ = let__4872

tmp5858 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4872)


if True == tmp5858 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4872
}


panic("unreachable")
})

tmp5859 := Call2(__e, ns2_1set, symshen_4_5rrb_6, tmp5851)


_ = tmp5859

tmp5860 := MakeNative1(func(__e *ControlFlow, V2795 Obj) Obj {
tmp5865 := Call2(__e, SymFn(symshen_4hds_a_2), V2795, MakeInteger(123))


var ifres5861 Obj

if True == tmp5865 {
tmp5862 := Call1(__e, SymFn(symtail), V2795)


let__4875 := tmp5862
_ = let__4875

tmp5863 := Call2(__e, SymFn(symshen_4comb), let__4875, symshen_4skip)


ifres5861 = tmp5863


} else {
tmp5864 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5861 = tmp5864


}

let__4874 := ifres5861
_ = let__4874

tmp5867 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4874)


if True == tmp5867 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4874
}


panic("unreachable")
})

tmp5868 := Call2(__e, ns2_1set, symshen_4_5lcurly_6, tmp5860)


_ = tmp5868

tmp5869 := MakeNative1(func(__e *ControlFlow, V2798 Obj) Obj {
tmp5874 := Call2(__e, SymFn(symshen_4hds_a_2), V2798, MakeInteger(125))


var ifres5870 Obj

if True == tmp5874 {
tmp5871 := Call1(__e, SymFn(symtail), V2798)


let__4877 := tmp5871
_ = let__4877

tmp5872 := Call2(__e, SymFn(symshen_4comb), let__4877, symshen_4skip)


ifres5870 = tmp5872


} else {
tmp5873 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5870 = tmp5873


}

let__4876 := ifres5870
_ = let__4876

tmp5876 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4876)


if True == tmp5876 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4876
}


panic("unreachable")
})

tmp5877 := Call2(__e, ns2_1set, symshen_4_5rcurly_6, tmp5869)


_ = tmp5877

tmp5878 := MakeNative1(func(__e *ControlFlow, V2801 Obj) Obj {
tmp5883 := Call2(__e, SymFn(symshen_4hds_a_2), V2801, MakeInteger(124))


var ifres5879 Obj

if True == tmp5883 {
tmp5880 := Call1(__e, SymFn(symtail), V2801)


let__4879 := tmp5880
_ = let__4879

tmp5881 := Call2(__e, SymFn(symshen_4comb), let__4879, symshen_4skip)


ifres5879 = tmp5881


} else {
tmp5882 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5879 = tmp5882


}

let__4878 := ifres5879
_ = let__4878

tmp5885 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4878)


if True == tmp5885 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4878
}


panic("unreachable")
})

tmp5886 := Call2(__e, ns2_1set, symshen_4_5bar_6, tmp5878)


_ = tmp5886

tmp5887 := MakeNative1(func(__e *ControlFlow, V2804 Obj) Obj {
tmp5892 := Call2(__e, SymFn(symshen_4hds_a_2), V2804, MakeInteger(59))


var ifres5888 Obj

if True == tmp5892 {
tmp5889 := Call1(__e, SymFn(symtail), V2804)


let__4881 := tmp5889
_ = let__4881

tmp5890 := Call2(__e, SymFn(symshen_4comb), let__4881, symshen_4skip)


ifres5888 = tmp5890


} else {
tmp5891 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5888 = tmp5891


}

let__4880 := ifres5888
_ = let__4880

tmp5894 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4880)


if True == tmp5894 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4880
}


panic("unreachable")
})

tmp5895 := Call2(__e, ns2_1set, symshen_4_5semicolon_6, tmp5887)


_ = tmp5895

tmp5896 := MakeNative1(func(__e *ControlFlow, V2807 Obj) Obj {
tmp5901 := Call2(__e, SymFn(symshen_4hds_a_2), V2807, MakeInteger(58))


var ifres5897 Obj

if True == tmp5901 {
tmp5898 := Call1(__e, SymFn(symtail), V2807)


let__4883 := tmp5898
_ = let__4883

tmp5899 := Call2(__e, SymFn(symshen_4comb), let__4883, symshen_4skip)


ifres5897 = tmp5899


} else {
tmp5900 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5897 = tmp5900


}

let__4882 := ifres5897
_ = let__4882

tmp5903 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4882)


if True == tmp5903 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4882
}


panic("unreachable")
})

tmp5904 := Call2(__e, ns2_1set, symshen_4_5colon_6, tmp5896)


_ = tmp5904

tmp5905 := MakeNative1(func(__e *ControlFlow, V2810 Obj) Obj {
tmp5910 := Call2(__e, SymFn(symshen_4hds_a_2), V2810, MakeInteger(44))


var ifres5906 Obj

if True == tmp5910 {
tmp5907 := Call1(__e, SymFn(symtail), V2810)


let__4885 := tmp5907
_ = let__4885

tmp5908 := Call2(__e, SymFn(symshen_4comb), let__4885, symshen_4skip)


ifres5906 = tmp5908


} else {
tmp5909 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5906 = tmp5909


}

let__4884 := ifres5906
_ = let__4884

tmp5912 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4884)


if True == tmp5912 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4884
}


panic("unreachable")
})

tmp5913 := Call2(__e, ns2_1set, symshen_4_5comma_6, tmp5905)


_ = tmp5913

tmp5914 := MakeNative1(func(__e *ControlFlow, V2813 Obj) Obj {
tmp5919 := Call2(__e, SymFn(symshen_4hds_a_2), V2813, MakeInteger(61))


var ifres5915 Obj

if True == tmp5919 {
tmp5916 := Call1(__e, SymFn(symtail), V2813)


let__4887 := tmp5916
_ = let__4887

tmp5917 := Call2(__e, SymFn(symshen_4comb), let__4887, symshen_4skip)


ifres5915 = tmp5917


} else {
tmp5918 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5915 = tmp5918


}

let__4886 := ifres5915
_ = let__4886

tmp5921 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4886)


if True == tmp5921 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4886
}


panic("unreachable")
})

tmp5922 := Call2(__e, ns2_1set, symshen_4_5equal_6, tmp5914)


_ = tmp5922

tmp5923 := MakeNative1(func(__e *ControlFlow, V2816 Obj) Obj {
tmp5924 := Call1(__e, SymFn(symshen_4_5singleline_6), V2816)


let__4889 := tmp5924
_ = let__4889

tmp5929 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4889)


var ifres5925 Obj

if True == tmp5929 {
tmp5926 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5925 = tmp5926


} else {
tmp5927 := Call1(__e, SymFn(symshen_4in_1_6), let__4889)


let__4890 := tmp5927
_ = let__4890

tmp5928 := Call2(__e, SymFn(symshen_4comb), let__4890, symshen_4skip)


ifres5925 = tmp5928


}

let__4888 := ifres5925
_ = let__4888

tmp5939 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4888)


if True == tmp5939 {
tmp5930 := Call1(__e, SymFn(symshen_4_5multiline_6), V2816)


let__4892 := tmp5930
_ = let__4892

tmp5935 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4892)


var ifres5931 Obj

if True == tmp5935 {
tmp5932 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5931 = tmp5932


} else {
tmp5933 := Call1(__e, SymFn(symshen_4in_1_6), let__4892)


let__4893 := tmp5933
_ = let__4893

tmp5934 := Call2(__e, SymFn(symshen_4comb), let__4893, symshen_4skip)


ifres5931 = tmp5934


}

let__4891 := ifres5931
_ = let__4891

tmp5937 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4891)


if True == tmp5937 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4891
}


} else {
return let__4888
}


panic("unreachable")
})

tmp5940 := Call2(__e, ns2_1set, symshen_4_5comment_6, tmp5923)


_ = tmp5940

tmp5941 := MakeNative1(func(__e *ControlFlow, V2823 Obj) Obj {
tmp5942 := Call1(__e, SymFn(symshen_4_5backslash_6), V2823)


let__4895 := tmp5942
_ = let__4895

tmp5962 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4895)


var ifres5943 Obj

if True == tmp5962 {
tmp5944 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5943 = tmp5944


} else {
tmp5945 := Call1(__e, SymFn(symshen_4in_1_6), let__4895)


let__4896 := tmp5945
_ = let__4896

tmp5946 := Call1(__e, SymFn(symshen_4_5backslash_6), let__4896)


let__4897 := tmp5946
_ = let__4897

tmp5961 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4897)


var ifres5947 Obj

if True == tmp5961 {
tmp5948 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5947 = tmp5948


} else {
tmp5949 := Call1(__e, SymFn(symshen_4in_1_6), let__4897)


let__4898 := tmp5949
_ = let__4898

tmp5950 := Call1(__e, SymFn(symshen_4_5shortnatters_6), let__4898)


let__4899 := tmp5950
_ = let__4899

tmp5960 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4899)


var ifres5951 Obj

if True == tmp5960 {
tmp5952 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5951 = tmp5952


} else {
tmp5953 := Call1(__e, SymFn(symshen_4in_1_6), let__4899)


let__4900 := tmp5953
_ = let__4900

tmp5954 := Call1(__e, SymFn(symshen_4_5returns_6), let__4900)


let__4901 := tmp5954
_ = let__4901

tmp5959 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4901)


var ifres5955 Obj

if True == tmp5959 {
tmp5956 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5955 = tmp5956


} else {
tmp5957 := Call1(__e, SymFn(symshen_4in_1_6), let__4901)


let__4902 := tmp5957
_ = let__4902

tmp5958 := Call2(__e, SymFn(symshen_4comb), let__4902, symshen_4skip)


ifres5955 = tmp5958


}

ifres5951 = ifres5955


}

ifres5947 = ifres5951


}

ifres5943 = ifres5947


}

let__4894 := ifres5943
_ = let__4894

tmp5964 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4894)


if True == tmp5964 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4894
}


panic("unreachable")
})

tmp5965 := Call2(__e, ns2_1set, symshen_4_5singleline_6, tmp5941)


_ = tmp5965

tmp5966 := MakeNative1(func(__e *ControlFlow, V2833 Obj) Obj {
tmp5971 := Call2(__e, SymFn(symshen_4hds_a_2), V2833, MakeInteger(92))


var ifres5967 Obj

if True == tmp5971 {
tmp5968 := Call1(__e, SymFn(symtail), V2833)


let__4904 := tmp5968
_ = let__4904

tmp5969 := Call2(__e, SymFn(symshen_4comb), let__4904, symshen_4skip)


ifres5967 = tmp5969


} else {
tmp5970 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5967 = tmp5970


}

let__4903 := ifres5967
_ = let__4903

tmp5973 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4903)


if True == tmp5973 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4903
}


panic("unreachable")
})

tmp5974 := Call2(__e, ns2_1set, symshen_4_5backslash_6, tmp5966)


_ = tmp5974

tmp5975 := MakeNative1(func(__e *ControlFlow, V2836 Obj) Obj {
tmp5976 := Call1(__e, SymFn(symshen_4_5shortnatter_6), V2836)


let__4906 := tmp5976
_ = let__4906

tmp5986 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4906)


var ifres5977 Obj

if True == tmp5986 {
tmp5978 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5977 = tmp5978


} else {
tmp5979 := Call1(__e, SymFn(symshen_4in_1_6), let__4906)


let__4907 := tmp5979
_ = let__4907

tmp5980 := Call1(__e, SymFn(symshen_4_5shortnatters_6), let__4907)


let__4908 := tmp5980
_ = let__4908

tmp5985 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4908)


var ifres5981 Obj

if True == tmp5985 {
tmp5982 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5981 = tmp5982


} else {
tmp5983 := Call1(__e, SymFn(symshen_4in_1_6), let__4908)


let__4909 := tmp5983
_ = let__4909

tmp5984 := Call2(__e, SymFn(symshen_4comb), let__4909, symshen_4skip)


ifres5981 = tmp5984


}

ifres5977 = ifres5981


}

let__4905 := ifres5977
_ = let__4905

tmp5996 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4905)


if True == tmp5996 {
tmp5987 := Call1(__e, SymFn(sym_5e_6), V2836)


let__4911 := tmp5987
_ = let__4911

tmp5992 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4911)


var ifres5988 Obj

if True == tmp5992 {
tmp5989 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5988 = tmp5989


} else {
tmp5990 := Call1(__e, SymFn(symshen_4in_1_6), let__4911)


let__4912 := tmp5990
_ = let__4912

tmp5991 := Call2(__e, SymFn(symshen_4comb), let__4912, symshen_4skip)


ifres5988 = tmp5991


}

let__4910 := ifres5988
_ = let__4910

tmp5994 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4910)


if True == tmp5994 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4910
}


} else {
return let__4905
}


panic("unreachable")
})

tmp5997 := Call2(__e, ns2_1set, symshen_4_5shortnatters_6, tmp5975)


_ = tmp5997

tmp5998 := MakeNative1(func(__e *ControlFlow, V2845 Obj) Obj {
tmp6008 := GConsP(__e, V2845)

var ifres5999 Obj

if True == tmp6008 {
tmp6000 := Call1(__e, SymFn(symhead), V2845)


let__4914 := tmp6000
_ = let__4914

tmp6001 := Call1(__e, SymFn(symtail), V2845)


let__4915 := tmp6001
_ = let__4915

tmp6005 := Call1(__e, SymFn(symshen_4return_2), let__4914)


tmp6006 := GNot(__e, tmp6005)

var ifres6002 Obj

if True == tmp6006 {
tmp6003 := Call2(__e, SymFn(symshen_4comb), let__4915, symshen_4skip)


ifres6002 = tmp6003


} else {
tmp6004 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6002 = tmp6004


}

ifres5999 = ifres6002


} else {
tmp6007 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5999 = tmp6007


}

let__4913 := ifres5999
_ = let__4913

tmp6010 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4913)


if True == tmp6010 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4913
}


panic("unreachable")
})

tmp6011 := Call2(__e, ns2_1set, symshen_4_5shortnatter_6, tmp5998)


_ = tmp6011

tmp6012 := MakeNative1(func(__e *ControlFlow, V2849 Obj) Obj {
tmp6013 := Call1(__e, SymFn(symshen_4_5return_6), V2849)


let__4917 := tmp6013
_ = let__4917

tmp6023 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4917)


var ifres6014 Obj

if True == tmp6023 {
tmp6015 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6014 = tmp6015


} else {
tmp6016 := Call1(__e, SymFn(symshen_4in_1_6), let__4917)


let__4918 := tmp6016
_ = let__4918

tmp6017 := Call1(__e, SymFn(symshen_4_5returns_6), let__4918)


let__4919 := tmp6017
_ = let__4919

tmp6022 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4919)


var ifres6018 Obj

if True == tmp6022 {
tmp6019 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6018 = tmp6019


} else {
tmp6020 := Call1(__e, SymFn(symshen_4in_1_6), let__4919)


let__4920 := tmp6020
_ = let__4920

tmp6021 := Call2(__e, SymFn(symshen_4comb), let__4920, symshen_4skip)


ifres6018 = tmp6021


}

ifres6014 = ifres6018


}

let__4916 := ifres6014
_ = let__4916

tmp6033 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4916)


if True == tmp6033 {
tmp6024 := Call1(__e, SymFn(symshen_4_5return_6), V2849)


let__4922 := tmp6024
_ = let__4922

tmp6029 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4922)


var ifres6025 Obj

if True == tmp6029 {
tmp6026 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6025 = tmp6026


} else {
tmp6027 := Call1(__e, SymFn(symshen_4in_1_6), let__4922)


let__4923 := tmp6027
_ = let__4923

tmp6028 := Call2(__e, SymFn(symshen_4comb), let__4923, symshen_4skip)


ifres6025 = tmp6028


}

let__4921 := ifres6025
_ = let__4921

tmp6031 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4921)


if True == tmp6031 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4921
}


} else {
return let__4916
}


panic("unreachable")
})

tmp6034 := Call2(__e, ns2_1set, symshen_4_5returns_6, tmp6012)


_ = tmp6034

tmp6035 := MakeNative1(func(__e *ControlFlow, V2858 Obj) Obj {
tmp6044 := GConsP(__e, V2858)

var ifres6036 Obj

if True == tmp6044 {
tmp6037 := Call1(__e, SymFn(symhead), V2858)


let__4925 := tmp6037
_ = let__4925

tmp6038 := Call1(__e, SymFn(symtail), V2858)


let__4926 := tmp6038
_ = let__4926

tmp6042 := Call1(__e, SymFn(symshen_4return_2), let__4925)


var ifres6039 Obj

if True == tmp6042 {
tmp6040 := Call2(__e, SymFn(symshen_4comb), let__4926, symshen_4skip)


ifres6039 = tmp6040


} else {
tmp6041 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6039 = tmp6041


}

ifres6036 = ifres6039


} else {
tmp6043 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6036 = tmp6043


}

let__4924 := ifres6036
_ = let__4924

tmp6046 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4924)


if True == tmp6046 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4924
}


panic("unreachable")
})

tmp6047 := Call2(__e, ns2_1set, symshen_4_5return_6, tmp6035)


_ = tmp6047

tmp6048 := MakeNative1(func(__e *ControlFlow, V2862 Obj) Obj {
tmp6049 := GCons(__e, MakeInteger(13), Nil)

tmp6050 := GCons(__e, MakeInteger(10), tmp6049)

tmp6051 := GCons(__e, MakeInteger(9), tmp6050)

return TailCall2(__e, SymFn(symelement_2), V2862, tmp6051)


panic("unreachable")
})

tmp6052 := Call2(__e, ns2_1set, symshen_4return_2, tmp6048)


_ = tmp6052

tmp6053 := MakeNative1(func(__e *ControlFlow, V2863 Obj) Obj {
tmp6054 := Call1(__e, SymFn(symshen_4_5backslash_6), V2863)


let__4928 := tmp6054
_ = let__4928

tmp6069 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4928)


var ifres6055 Obj

if True == tmp6069 {
tmp6056 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6055 = tmp6056


} else {
tmp6057 := Call1(__e, SymFn(symshen_4in_1_6), let__4928)


let__4929 := tmp6057
_ = let__4929

tmp6058 := Call1(__e, SymFn(symshen_4_5times_6), let__4929)


let__4930 := tmp6058
_ = let__4930

tmp6068 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4930)


var ifres6059 Obj

if True == tmp6068 {
tmp6060 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6059 = tmp6060


} else {
tmp6061 := Call1(__e, SymFn(symshen_4in_1_6), let__4930)


let__4931 := tmp6061
_ = let__4931

tmp6062 := Call1(__e, SymFn(symshen_4_5longnatter_6), let__4931)


let__4932 := tmp6062
_ = let__4932

tmp6067 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4932)


var ifres6063 Obj

if True == tmp6067 {
tmp6064 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6063 = tmp6064


} else {
tmp6065 := Call1(__e, SymFn(symshen_4in_1_6), let__4932)


let__4933 := tmp6065
_ = let__4933

tmp6066 := Call2(__e, SymFn(symshen_4comb), let__4933, symshen_4skip)


ifres6063 = tmp6066


}

ifres6059 = ifres6063


}

ifres6055 = ifres6059


}

let__4927 := ifres6055
_ = let__4927

tmp6071 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4927)


if True == tmp6071 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4927
}


panic("unreachable")
})

tmp6072 := Call2(__e, ns2_1set, symshen_4_5multiline_6, tmp6053)


_ = tmp6072

tmp6073 := MakeNative1(func(__e *ControlFlow, V2871 Obj) Obj {
tmp6078 := Call2(__e, SymFn(symshen_4hds_a_2), V2871, MakeInteger(42))


var ifres6074 Obj

if True == tmp6078 {
tmp6075 := Call1(__e, SymFn(symtail), V2871)


let__4935 := tmp6075
_ = let__4935

tmp6076 := Call2(__e, SymFn(symshen_4comb), let__4935, symshen_4skip)


ifres6074 = tmp6076


} else {
tmp6077 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6074 = tmp6077


}

let__4934 := ifres6074
_ = let__4934

tmp6080 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4934)


if True == tmp6080 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4934
}


panic("unreachable")
})

tmp6081 := Call2(__e, ns2_1set, symshen_4_5times_6, tmp6073)


_ = tmp6081

tmp6082 := MakeNative1(func(__e *ControlFlow, V2874 Obj) Obj {
tmp6083 := Call1(__e, SymFn(symshen_4_5comment_6), V2874)


let__4937 := tmp6083
_ = let__4937

tmp6093 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4937)


var ifres6084 Obj

if True == tmp6093 {
tmp6085 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6084 = tmp6085


} else {
tmp6086 := Call1(__e, SymFn(symshen_4in_1_6), let__4937)


let__4938 := tmp6086
_ = let__4938

tmp6087 := Call1(__e, SymFn(symshen_4_5longnatter_6), let__4938)


let__4939 := tmp6087
_ = let__4939

tmp6092 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4939)


var ifres6088 Obj

if True == tmp6092 {
tmp6089 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6088 = tmp6089


} else {
tmp6090 := Call1(__e, SymFn(symshen_4in_1_6), let__4939)


let__4940 := tmp6090
_ = let__4940

tmp6091 := Call2(__e, SymFn(symshen_4comb), let__4940, symshen_4skip)


ifres6088 = tmp6091


}

ifres6084 = ifres6088


}

let__4936 := ifres6084
_ = let__4936

tmp6120 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4936)


if True == tmp6120 {
tmp6094 := Call1(__e, SymFn(symshen_4_5times_6), V2874)


let__4942 := tmp6094
_ = let__4942

tmp6104 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4942)


var ifres6095 Obj

if True == tmp6104 {
tmp6096 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6095 = tmp6096


} else {
tmp6097 := Call1(__e, SymFn(symshen_4in_1_6), let__4942)


let__4943 := tmp6097
_ = let__4943

tmp6098 := Call1(__e, SymFn(symshen_4_5backslash_6), let__4943)


let__4944 := tmp6098
_ = let__4944

tmp6103 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4944)


var ifres6099 Obj

if True == tmp6103 {
tmp6100 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6099 = tmp6100


} else {
tmp6101 := Call1(__e, SymFn(symshen_4in_1_6), let__4944)


let__4945 := tmp6101
_ = let__4945

tmp6102 := Call2(__e, SymFn(symshen_4comb), let__4945, symshen_4skip)


ifres6099 = tmp6102


}

ifres6095 = ifres6099


}

let__4941 := ifres6095
_ = let__4941

tmp6118 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4941)


if True == tmp6118 {
tmp6114 := GConsP(__e, V2874)

var ifres6105 Obj

if True == tmp6114 {
tmp6106 := Call1(__e, SymFn(symtail), V2874)


let__4947 := tmp6106
_ = let__4947

tmp6107 := Call1(__e, SymFn(symshen_4_5longnatter_6), let__4947)


let__4948 := tmp6107
_ = let__4948

tmp6112 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4948)


var ifres6108 Obj

if True == tmp6112 {
tmp6109 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6108 = tmp6109


} else {
tmp6110 := Call1(__e, SymFn(symshen_4in_1_6), let__4948)


let__4949 := tmp6110
_ = let__4949

tmp6111 := Call2(__e, SymFn(symshen_4comb), let__4949, symshen_4skip)


ifres6108 = tmp6111


}

ifres6105 = ifres6108


} else {
tmp6113 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6105 = tmp6113


}

let__4946 := ifres6105
_ = let__4946

tmp6116 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4946)


if True == tmp6116 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4946
}


} else {
return let__4941
}


} else {
return let__4936
}


panic("unreachable")
})

tmp6121 := Call2(__e, ns2_1set, symshen_4_5longnatter_6, tmp6082)


_ = tmp6121

tmp6122 := MakeNative1(func(__e *ControlFlow, V2889 Obj) Obj {
tmp6123 := Call1(__e, SymFn(symshen_4_5str_6), V2889)


let__4951 := tmp6123
_ = let__4951

tmp6129 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4951)


var ifres6124 Obj

if True == tmp6129 {
tmp6125 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6124 = tmp6125


} else {
tmp6126 := Call1(__e, SymFn(symshen_4_5_1out), let__4951)


let__4952 := tmp6126
_ = let__4952

tmp6127 := Call1(__e, SymFn(symshen_4in_1_6), let__4951)


let__4953 := tmp6127
_ = let__4953

tmp6128 := Call2(__e, SymFn(symshen_4comb), let__4953, let__4952)


ifres6124 = tmp6128


}

let__4950 := ifres6124
_ = let__4950

tmp6154 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4950)


if True == tmp6154 {
tmp6130 := Call1(__e, SymFn(symshen_4_5number_6), V2889)


let__4955 := tmp6130
_ = let__4955

tmp6136 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4955)


var ifres6131 Obj

if True == tmp6136 {
tmp6132 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6131 = tmp6132


} else {
tmp6133 := Call1(__e, SymFn(symshen_4_5_1out), let__4955)


let__4956 := tmp6133
_ = let__4956

tmp6134 := Call1(__e, SymFn(symshen_4in_1_6), let__4955)


let__4957 := tmp6134
_ = let__4957

tmp6135 := Call2(__e, SymFn(symshen_4comb), let__4957, let__4956)


ifres6131 = tmp6135


}

let__4954 := ifres6131
_ = let__4954

tmp6152 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4954)


if True == tmp6152 {
tmp6137 := Call1(__e, SymFn(symshen_4_5sym_6), V2889)


let__4959 := tmp6137
_ = let__4959

tmp6148 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4959)


var ifres6138 Obj

if True == tmp6148 {
tmp6139 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6138 = tmp6139


} else {
tmp6140 := Call1(__e, SymFn(symshen_4_5_1out), let__4959)


let__4960 := tmp6140
_ = let__4960

tmp6141 := Call1(__e, SymFn(symshen_4in_1_6), let__4959)


let__4961 := tmp6141
_ = let__4961

tmp6146 := GEq(__e, let__4960, MakeString("<>"))

var ifres6142 Obj

if True == tmp6146 {
tmp6143 := GCons(__e, MakeInteger(0), Nil)

tmp6144 := GCons(__e, symvector, tmp6143)

ifres6142 = tmp6144


} else {
tmp6145 := PrimGuard1(__e, symintern, PrimIntern, let__4960)

ifres6142 = tmp6145


}

tmp6147 := Call2(__e, SymFn(symshen_4comb), let__4961, ifres6142)


ifres6138 = tmp6147


}

let__4958 := ifres6138
_ = let__4958

tmp6150 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4958)


if True == tmp6150 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4958
}


} else {
return let__4954
}


} else {
return let__4950
}


panic("unreachable")
})

tmp6155 := Call2(__e, ns2_1set, symshen_4_5atom_6, tmp6122)


_ = tmp6155

tmp6156 := MakeNative1(func(__e *ControlFlow, V2902 Obj) Obj {
tmp6157 := Call1(__e, SymFn(symshen_4_5alpha_6), V2902)


let__4963 := tmp6157
_ = let__4963

tmp6170 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4963)


var ifres6158 Obj

if True == tmp6170 {
tmp6159 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6158 = tmp6159


} else {
tmp6160 := Call1(__e, SymFn(symshen_4_5_1out), let__4963)


let__4964 := tmp6160
_ = let__4964

tmp6161 := Call1(__e, SymFn(symshen_4in_1_6), let__4963)


let__4965 := tmp6161
_ = let__4965

tmp6162 := Call1(__e, SymFn(symshen_4_5alphanums_6), let__4965)


let__4966 := tmp6162
_ = let__4966

tmp6169 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4966)


var ifres6163 Obj

if True == tmp6169 {
tmp6164 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6163 = tmp6164


} else {
tmp6165 := Call1(__e, SymFn(symshen_4_5_1out), let__4966)


let__4967 := tmp6165
_ = let__4967

tmp6166 := Call1(__e, SymFn(symshen_4in_1_6), let__4966)


let__4968 := tmp6166
_ = let__4968

tmp6168 := Call2(__e, SymFn(symshen_4comb), let__4968, PrimGuard2(__e, symcn, PrimStringConcat, let__4964, let__4967))


ifres6163 = tmp6168


}

ifres6158 = ifres6163


}

let__4962 := ifres6158
_ = let__4962

tmp6172 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4962)


if True == tmp6172 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4962
}


panic("unreachable")
})

tmp6173 := Call2(__e, ns2_1set, symshen_4_5sym_6, tmp6156)


_ = tmp6173

tmp6174 := MakeNative1(func(__e *ControlFlow, V2910 Obj) Obj {
tmp6184 := GConsP(__e, V2910)

var ifres6175 Obj

if True == tmp6184 {
tmp6176 := Call1(__e, SymFn(symhead), V2910)


let__4970 := tmp6176
_ = let__4970

tmp6177 := Call1(__e, SymFn(symtail), V2910)


let__4971 := tmp6177
_ = let__4971

tmp6182 := Call1(__e, SymFn(symshen_4alpha_2), let__4970)


var ifres6178 Obj

if True == tmp6182 {
tmp6179 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, let__4970)

tmp6180 := Call2(__e, SymFn(symshen_4comb), let__4971, tmp6179)


ifres6178 = tmp6180


} else {
tmp6181 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6178 = tmp6181


}

ifres6175 = ifres6178


} else {
tmp6183 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6175 = tmp6183


}

let__4969 := ifres6175
_ = let__4969

tmp6186 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4969)


if True == tmp6186 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4969
}


panic("unreachable")
})

tmp6187 := Call2(__e, ns2_1set, symshen_4_5alpha_6, tmp6174)


_ = tmp6187

tmp6188 := MakeNative1(func(__e *ControlFlow, V2914 Obj) Obj {
tmp6192 := Call1(__e, SymFn(symshen_4lowercase_2), V2914)


if True == tmp6192 {
return True
} else {
tmp6190 := Call1(__e, SymFn(symshen_4uppercase_2), V2914)


if True == tmp6190 {
return True
} else {
return TailCall1(__e, SymFn(symshen_4misc_2), V2914)
}


}


panic("unreachable")
})

tmp6193 := Call2(__e, ns2_1set, symshen_4alpha_2, tmp6188)


_ = tmp6193

tmp6194 := MakeNative1(func(__e *ControlFlow, V2915 Obj) Obj {
if True == GGe(__e, V2915, MakeInteger(97)) {
return GLe(__e, V2915, MakeInteger(122))
} else {
return False
}


panic("unreachable")
})

tmp6197 := Call2(__e, ns2_1set, symshen_4lowercase_2, tmp6194)


_ = tmp6197

tmp6198 := MakeNative1(func(__e *ControlFlow, V2916 Obj) Obj {
if True == GGe(__e, V2916, MakeInteger(65)) {
return GLe(__e, V2916, MakeInteger(90))
} else {
return False
}


panic("unreachable")
})

tmp6201 := Call2(__e, ns2_1set, symshen_4uppercase_2, tmp6198)


_ = tmp6201

tmp6202 := MakeNative1(func(__e *ControlFlow, V2917 Obj) Obj {
tmp6203 := GCons(__e, MakeInteger(96), Nil)

tmp6204 := GCons(__e, MakeInteger(35), tmp6203)

tmp6205 := GCons(__e, MakeInteger(39), tmp6204)

tmp6206 := GCons(__e, MakeInteger(37), tmp6205)

tmp6207 := GCons(__e, MakeInteger(38), tmp6206)

tmp6208 := GCons(__e, MakeInteger(60), tmp6207)

tmp6209 := GCons(__e, MakeInteger(62), tmp6208)

tmp6210 := GCons(__e, MakeInteger(46), tmp6209)

tmp6211 := GCons(__e, MakeInteger(126), tmp6210)

tmp6212 := GCons(__e, MakeInteger(64), tmp6211)

tmp6213 := GCons(__e, MakeInteger(33), tmp6212)

tmp6214 := GCons(__e, MakeInteger(36), tmp6213)

tmp6215 := GCons(__e, MakeInteger(63), tmp6214)

tmp6216 := GCons(__e, MakeInteger(95), tmp6215)

tmp6217 := GCons(__e, MakeInteger(43), tmp6216)

tmp6218 := GCons(__e, MakeInteger(47), tmp6217)

tmp6219 := GCons(__e, MakeInteger(42), tmp6218)

tmp6220 := GCons(__e, MakeInteger(45), tmp6219)

tmp6221 := GCons(__e, MakeInteger(61), tmp6220)

return TailCall2(__e, SymFn(symelement_2), V2917, tmp6221)


panic("unreachable")
})

tmp6222 := Call2(__e, ns2_1set, symshen_4misc_2, tmp6202)


_ = tmp6222

tmp6223 := MakeNative1(func(__e *ControlFlow, V2918 Obj) Obj {
tmp6224 := Call1(__e, SymFn(symshen_4_5alphanum_6), V2918)


let__4973 := tmp6224
_ = let__4973

tmp6237 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4973)


var ifres6225 Obj

if True == tmp6237 {
tmp6226 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6225 = tmp6226


} else {
tmp6227 := Call1(__e, SymFn(symshen_4_5_1out), let__4973)


let__4974 := tmp6227
_ = let__4974

tmp6228 := Call1(__e, SymFn(symshen_4in_1_6), let__4973)


let__4975 := tmp6228
_ = let__4975

tmp6229 := Call1(__e, SymFn(symshen_4_5alphanums_6), let__4975)


let__4976 := tmp6229
_ = let__4976

tmp6236 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4976)


var ifres6230 Obj

if True == tmp6236 {
tmp6231 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6230 = tmp6231


} else {
tmp6232 := Call1(__e, SymFn(symshen_4_5_1out), let__4976)


let__4977 := tmp6232
_ = let__4977

tmp6233 := Call1(__e, SymFn(symshen_4in_1_6), let__4976)


let__4978 := tmp6233
_ = let__4978

tmp6235 := Call2(__e, SymFn(symshen_4comb), let__4978, PrimGuard2(__e, symcn, PrimStringConcat, let__4974, let__4977))


ifres6230 = tmp6235


}

ifres6225 = ifres6230


}

let__4972 := ifres6225
_ = let__4972

tmp6247 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4972)


if True == tmp6247 {
tmp6238 := Call1(__e, SymFn(sym_5e_6), V2918)


let__4980 := tmp6238
_ = let__4980

tmp6243 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4980)


var ifres6239 Obj

if True == tmp6243 {
tmp6240 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6239 = tmp6240


} else {
tmp6241 := Call1(__e, SymFn(symshen_4in_1_6), let__4980)


let__4981 := tmp6241
_ = let__4981

tmp6242 := Call2(__e, SymFn(symshen_4comb), let__4981, MakeString(""))


ifres6239 = tmp6242


}

let__4979 := ifres6239
_ = let__4979

tmp6245 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4979)


if True == tmp6245 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4979
}


} else {
return let__4972
}


panic("unreachable")
})

tmp6248 := Call2(__e, ns2_1set, symshen_4_5alphanums_6, tmp6223)


_ = tmp6248

tmp6249 := MakeNative1(func(__e *ControlFlow, V2929 Obj) Obj {
tmp6250 := Call1(__e, SymFn(symshen_4_5alpha_6), V2929)


let__4983 := tmp6250
_ = let__4983

tmp6256 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4983)


var ifres6251 Obj

if True == tmp6256 {
tmp6252 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6251 = tmp6252


} else {
tmp6253 := Call1(__e, SymFn(symshen_4_5_1out), let__4983)


let__4984 := tmp6253
_ = let__4984

tmp6254 := Call1(__e, SymFn(symshen_4in_1_6), let__4983)


let__4985 := tmp6254
_ = let__4985

tmp6255 := Call2(__e, SymFn(symshen_4comb), let__4985, let__4984)


ifres6251 = tmp6255


}

let__4982 := ifres6251
_ = let__4982

tmp6267 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4982)


if True == tmp6267 {
tmp6257 := Call1(__e, SymFn(symshen_4_5numeral_6), V2929)


let__4987 := tmp6257
_ = let__4987

tmp6263 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4987)


var ifres6258 Obj

if True == tmp6263 {
tmp6259 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6258 = tmp6259


} else {
tmp6260 := Call1(__e, SymFn(symshen_4_5_1out), let__4987)


let__4988 := tmp6260
_ = let__4988

tmp6261 := Call1(__e, SymFn(symshen_4in_1_6), let__4987)


let__4989 := tmp6261
_ = let__4989

tmp6262 := Call2(__e, SymFn(symshen_4comb), let__4989, let__4988)


ifres6258 = tmp6262


}

let__4986 := ifres6258
_ = let__4986

tmp6265 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4986)


if True == tmp6265 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4986
}


} else {
return let__4982
}


panic("unreachable")
})

tmp6268 := Call2(__e, ns2_1set, symshen_4_5alphanum_6, tmp6249)


_ = tmp6268

tmp6269 := MakeNative1(func(__e *ControlFlow, V2938 Obj) Obj {
tmp6279 := GConsP(__e, V2938)

var ifres6270 Obj

if True == tmp6279 {
tmp6271 := Call1(__e, SymFn(symhead), V2938)


let__4991 := tmp6271
_ = let__4991

tmp6272 := Call1(__e, SymFn(symtail), V2938)


let__4992 := tmp6272
_ = let__4992

tmp6277 := Call1(__e, SymFn(symshen_4digit_2), let__4991)


var ifres6273 Obj

if True == tmp6277 {
tmp6274 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, let__4991)

tmp6275 := Call2(__e, SymFn(symshen_4comb), let__4992, tmp6274)


ifres6273 = tmp6275


} else {
tmp6276 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6273 = tmp6276


}

ifres6270 = ifres6273


} else {
tmp6278 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6270 = tmp6278


}

let__4990 := ifres6270
_ = let__4990

tmp6281 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4990)


if True == tmp6281 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4990
}


panic("unreachable")
})

tmp6282 := Call2(__e, ns2_1set, symshen_4_5numeral_6, tmp6269)


_ = tmp6282

tmp6283 := MakeNative1(func(__e *ControlFlow, V2942 Obj) Obj {
if True == GGe(__e, V2942, MakeInteger(48)) {
return GLe(__e, V2942, MakeInteger(57))
} else {
return False
}


panic("unreachable")
})

tmp6286 := Call2(__e, ns2_1set, symshen_4digit_2, tmp6283)


_ = tmp6286

tmp6287 := MakeNative1(func(__e *ControlFlow, V2943 Obj) Obj {
tmp6288 := Call1(__e, SymFn(symshen_4_5dbq_6), V2943)


let__4994 := tmp6288
_ = let__4994

tmp6304 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4994)


var ifres6289 Obj

if True == tmp6304 {
tmp6290 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6289 = tmp6290


} else {
tmp6291 := Call1(__e, SymFn(symshen_4in_1_6), let__4994)


let__4995 := tmp6291
_ = let__4995

tmp6292 := Call1(__e, SymFn(symshen_4_5strcontents_6), let__4995)


let__4996 := tmp6292
_ = let__4996

tmp6303 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4996)


var ifres6293 Obj

if True == tmp6303 {
tmp6294 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6293 = tmp6294


} else {
tmp6295 := Call1(__e, SymFn(symshen_4_5_1out), let__4996)


let__4997 := tmp6295
_ = let__4997

tmp6296 := Call1(__e, SymFn(symshen_4in_1_6), let__4996)


let__4998 := tmp6296
_ = let__4998

tmp6297 := Call1(__e, SymFn(symshen_4_5dbq_6), let__4998)


let__4999 := tmp6297
_ = let__4999

tmp6302 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4999)


var ifres6298 Obj

if True == tmp6302 {
tmp6299 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6298 = tmp6299


} else {
tmp6300 := Call1(__e, SymFn(symshen_4in_1_6), let__4999)


let__5000 := tmp6300
_ = let__5000

tmp6301 := Call2(__e, SymFn(symshen_4comb), let__5000, let__4997)


ifres6298 = tmp6301


}

ifres6293 = ifres6298


}

ifres6289 = ifres6293


}

let__4993 := ifres6289
_ = let__4993

tmp6306 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4993)


if True == tmp6306 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4993
}


panic("unreachable")
})

tmp6307 := Call2(__e, ns2_1set, symshen_4_5str_6, tmp6287)


_ = tmp6307

tmp6308 := MakeNative1(func(__e *ControlFlow, V2952 Obj) Obj {
tmp6313 := Call2(__e, SymFn(symshen_4hds_a_2), V2952, MakeInteger(34))


var ifres6309 Obj

if True == tmp6313 {
tmp6310 := Call1(__e, SymFn(symtail), V2952)


let__5002 := tmp6310
_ = let__5002

tmp6311 := Call2(__e, SymFn(symshen_4comb), let__5002, symshen_4skip)


ifres6309 = tmp6311


} else {
tmp6312 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6309 = tmp6312


}

let__5001 := ifres6309
_ = let__5001

tmp6315 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5001)


if True == tmp6315 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5001
}


panic("unreachable")
})

tmp6316 := Call2(__e, ns2_1set, symshen_4_5dbq_6, tmp6308)


_ = tmp6316

tmp6317 := MakeNative1(func(__e *ControlFlow, V2955 Obj) Obj {
tmp6318 := Call1(__e, SymFn(symshen_4_5strc_6), V2955)


let__5004 := tmp6318
_ = let__5004

tmp6331 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5004)


var ifres6319 Obj

if True == tmp6331 {
tmp6320 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6319 = tmp6320


} else {
tmp6321 := Call1(__e, SymFn(symshen_4_5_1out), let__5004)


let__5005 := tmp6321
_ = let__5005

tmp6322 := Call1(__e, SymFn(symshen_4in_1_6), let__5004)


let__5006 := tmp6322
_ = let__5006

tmp6323 := Call1(__e, SymFn(symshen_4_5strcontents_6), let__5006)


let__5007 := tmp6323
_ = let__5007

tmp6330 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5007)


var ifres6324 Obj

if True == tmp6330 {
tmp6325 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6324 = tmp6325


} else {
tmp6326 := Call1(__e, SymFn(symshen_4_5_1out), let__5007)


let__5008 := tmp6326
_ = let__5008

tmp6327 := Call1(__e, SymFn(symshen_4in_1_6), let__5007)


let__5009 := tmp6327
_ = let__5009

tmp6329 := Call2(__e, SymFn(symshen_4comb), let__5009, PrimGuard2(__e, symcn, PrimStringConcat, let__5005, let__5008))


ifres6324 = tmp6329


}

ifres6319 = ifres6324


}

let__5003 := ifres6319
_ = let__5003

tmp6341 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5003)


if True == tmp6341 {
tmp6332 := Call1(__e, SymFn(sym_5e_6), V2955)


let__5011 := tmp6332
_ = let__5011

tmp6337 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5011)


var ifres6333 Obj

if True == tmp6337 {
tmp6334 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6333 = tmp6334


} else {
tmp6335 := Call1(__e, SymFn(symshen_4in_1_6), let__5011)


let__5012 := tmp6335
_ = let__5012

tmp6336 := Call2(__e, SymFn(symshen_4comb), let__5012, MakeString(""))


ifres6333 = tmp6336


}

let__5010 := ifres6333
_ = let__5010

tmp6339 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5010)


if True == tmp6339 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5010
}


} else {
return let__5003
}


panic("unreachable")
})

tmp6342 := Call2(__e, ns2_1set, symshen_4_5strcontents_6, tmp6317)


_ = tmp6342

tmp6343 := MakeNative1(func(__e *ControlFlow, V2966 Obj) Obj {
tmp6344 := Call1(__e, SymFn(symshen_4_5control_6), V2966)


let__5014 := tmp6344
_ = let__5014

tmp6350 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5014)


var ifres6345 Obj

if True == tmp6350 {
tmp6346 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6345 = tmp6346


} else {
tmp6347 := Call1(__e, SymFn(symshen_4_5_1out), let__5014)


let__5015 := tmp6347
_ = let__5015

tmp6348 := Call1(__e, SymFn(symshen_4in_1_6), let__5014)


let__5016 := tmp6348
_ = let__5016

tmp6349 := Call2(__e, SymFn(symshen_4comb), let__5016, let__5015)


ifres6345 = tmp6349


}

let__5013 := ifres6345
_ = let__5013

tmp6361 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5013)


if True == tmp6361 {
tmp6351 := Call1(__e, SymFn(symshen_4_5notdbq_6), V2966)


let__5018 := tmp6351
_ = let__5018

tmp6357 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5018)


var ifres6352 Obj

if True == tmp6357 {
tmp6353 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6352 = tmp6353


} else {
tmp6354 := Call1(__e, SymFn(symshen_4_5_1out), let__5018)


let__5019 := tmp6354
_ = let__5019

tmp6355 := Call1(__e, SymFn(symshen_4in_1_6), let__5018)


let__5020 := tmp6355
_ = let__5020

tmp6356 := Call2(__e, SymFn(symshen_4comb), let__5020, let__5019)


ifres6352 = tmp6356


}

let__5017 := ifres6352
_ = let__5017

tmp6359 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5017)


if True == tmp6359 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5017
}


} else {
return let__5013
}


panic("unreachable")
})

tmp6362 := Call2(__e, ns2_1set, symshen_4_5strc_6, tmp6343)


_ = tmp6362

tmp6363 := MakeNative1(func(__e *ControlFlow, V2975 Obj) Obj {
tmp6364 := Call1(__e, SymFn(symshen_4_5lowC_6), V2975)


let__5022 := tmp6364
_ = let__5022

tmp6386 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5022)


var ifres6365 Obj

if True == tmp6386 {
tmp6366 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6365 = tmp6366


} else {
tmp6367 := Call1(__e, SymFn(symshen_4in_1_6), let__5022)


let__5023 := tmp6367
_ = let__5023

tmp6368 := Call1(__e, SymFn(symshen_4_5hash_6), let__5023)


let__5024 := tmp6368
_ = let__5024

tmp6385 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5024)


var ifres6369 Obj

if True == tmp6385 {
tmp6370 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6369 = tmp6370


} else {
tmp6371 := Call1(__e, SymFn(symshen_4in_1_6), let__5024)


let__5025 := tmp6371
_ = let__5025

tmp6372 := Call1(__e, SymFn(symshen_4_5integer_6), let__5025)


let__5026 := tmp6372
_ = let__5026

tmp6384 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5026)


var ifres6373 Obj

if True == tmp6384 {
tmp6374 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6373 = tmp6374


} else {
tmp6375 := Call1(__e, SymFn(symshen_4_5_1out), let__5026)


let__5027 := tmp6375
_ = let__5027

tmp6376 := Call1(__e, SymFn(symshen_4in_1_6), let__5026)


let__5028 := tmp6376
_ = let__5028

tmp6377 := Call1(__e, SymFn(symshen_4_5semicolon_6), let__5028)


let__5029 := tmp6377
_ = let__5029

tmp6383 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5029)


var ifres6378 Obj

if True == tmp6383 {
tmp6379 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6378 = tmp6379


} else {
tmp6380 := Call1(__e, SymFn(symshen_4in_1_6), let__5029)


let__5030 := tmp6380
_ = let__5030

tmp6381 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, let__5027)

tmp6382 := Call2(__e, SymFn(symshen_4comb), let__5030, tmp6381)


ifres6378 = tmp6382


}

ifres6373 = ifres6378


}

ifres6369 = ifres6373


}

ifres6365 = ifres6369


}

let__5021 := ifres6365
_ = let__5021

tmp6388 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5021)


if True == tmp6388 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5021
}


panic("unreachable")
})

tmp6389 := Call2(__e, ns2_1set, symshen_4_5control_6, tmp6363)


_ = tmp6389

tmp6390 := MakeNative1(func(__e *ControlFlow, V2986 Obj) Obj {
tmp6401 := GConsP(__e, V2986)

var ifres6391 Obj

if True == tmp6401 {
tmp6392 := Call1(__e, SymFn(symhead), V2986)


let__5032 := tmp6392
_ = let__5032

tmp6393 := Call1(__e, SymFn(symtail), V2986)


let__5033 := tmp6393
_ = let__5033

tmp6398 := GEq(__e, let__5032, MakeInteger(34))

tmp6399 := GNot(__e, tmp6398)

var ifres6394 Obj

if True == tmp6399 {
tmp6395 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, let__5032)

tmp6396 := Call2(__e, SymFn(symshen_4comb), let__5033, tmp6395)


ifres6394 = tmp6396


} else {
tmp6397 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6394 = tmp6397


}

ifres6391 = ifres6394


} else {
tmp6400 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6391 = tmp6400


}

let__5031 := ifres6391
_ = let__5031

tmp6403 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5031)


if True == tmp6403 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5031
}


panic("unreachable")
})

tmp6404 := Call2(__e, ns2_1set, symshen_4_5notdbq_6, tmp6390)


_ = tmp6404

tmp6405 := MakeNative1(func(__e *ControlFlow, V2990 Obj) Obj {
tmp6410 := Call2(__e, SymFn(symshen_4hds_a_2), V2990, MakeInteger(99))


var ifres6406 Obj

if True == tmp6410 {
tmp6407 := Call1(__e, SymFn(symtail), V2990)


let__5035 := tmp6407
_ = let__5035

tmp6408 := Call2(__e, SymFn(symshen_4comb), let__5035, symshen_4skip)


ifres6406 = tmp6408


} else {
tmp6409 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6406 = tmp6409


}

let__5034 := ifres6406
_ = let__5034

tmp6412 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5034)


if True == tmp6412 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5034
}


panic("unreachable")
})

tmp6413 := Call2(__e, ns2_1set, symshen_4_5lowC_6, tmp6405)


_ = tmp6413

tmp6414 := MakeNative1(func(__e *ControlFlow, V2993 Obj) Obj {
tmp6419 := Call2(__e, SymFn(symshen_4hds_a_2), V2993, MakeInteger(35))


var ifres6415 Obj

if True == tmp6419 {
tmp6416 := Call1(__e, SymFn(symtail), V2993)


let__5037 := tmp6416
_ = let__5037

tmp6417 := Call2(__e, SymFn(symshen_4comb), let__5037, symshen_4skip)


ifres6415 = tmp6417


} else {
tmp6418 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6415 = tmp6418


}

let__5036 := ifres6415
_ = let__5036

tmp6421 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5036)


if True == tmp6421 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5036
}


panic("unreachable")
})

tmp6422 := Call2(__e, ns2_1set, symshen_4_5hash_6, tmp6414)


_ = tmp6422

tmp6423 := MakeNative1(func(__e *ControlFlow, V2996 Obj) Obj {
tmp6424 := Call1(__e, SymFn(symshen_4_5minus_6), V2996)


let__5039 := tmp6424
_ = let__5039

tmp6436 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5039)


var ifres6425 Obj

if True == tmp6436 {
tmp6426 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6425 = tmp6426


} else {
tmp6427 := Call1(__e, SymFn(symshen_4in_1_6), let__5039)


let__5040 := tmp6427
_ = let__5040

tmp6428 := Call1(__e, SymFn(symshen_4_5number_6), let__5040)


let__5041 := tmp6428
_ = let__5041

tmp6435 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5041)


var ifres6429 Obj

if True == tmp6435 {
tmp6430 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6429 = tmp6430


} else {
tmp6431 := Call1(__e, SymFn(symshen_4_5_1out), let__5041)


let__5042 := tmp6431
_ = let__5042

tmp6432 := Call1(__e, SymFn(symshen_4in_1_6), let__5041)


let__5043 := tmp6432
_ = let__5043

tmp6434 := Call2(__e, SymFn(symshen_4comb), let__5043, GSub(__e, MakeInteger(0), let__5042))


ifres6429 = tmp6434


}

ifres6425 = ifres6429


}

let__5038 := ifres6425
_ = let__5038

tmp6479 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5038)


if True == tmp6479 {
tmp6437 := Call1(__e, SymFn(symshen_4_5plus_6), V2996)


let__5045 := tmp6437
_ = let__5045

tmp6448 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5045)


var ifres6438 Obj

if True == tmp6448 {
tmp6439 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6438 = tmp6439


} else {
tmp6440 := Call1(__e, SymFn(symshen_4in_1_6), let__5045)


let__5046 := tmp6440
_ = let__5046

tmp6441 := Call1(__e, SymFn(symshen_4_5number_6), let__5046)


let__5047 := tmp6441
_ = let__5047

tmp6447 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5047)


var ifres6442 Obj

if True == tmp6447 {
tmp6443 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6442 = tmp6443


} else {
tmp6444 := Call1(__e, SymFn(symshen_4_5_1out), let__5047)


let__5048 := tmp6444
_ = let__5048

tmp6445 := Call1(__e, SymFn(symshen_4in_1_6), let__5047)


let__5049 := tmp6445
_ = let__5049

tmp6446 := Call2(__e, SymFn(symshen_4comb), let__5049, let__5048)


ifres6442 = tmp6446


}

ifres6438 = ifres6442


}

let__5044 := ifres6438
_ = let__5044

tmp6477 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5044)


if True == tmp6477 {
tmp6449 := Call1(__e, SymFn(symshen_4_5e_1number_6), V2996)


let__5051 := tmp6449
_ = let__5051

tmp6455 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5051)


var ifres6450 Obj

if True == tmp6455 {
tmp6451 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6450 = tmp6451


} else {
tmp6452 := Call1(__e, SymFn(symshen_4_5_1out), let__5051)


let__5052 := tmp6452
_ = let__5052

tmp6453 := Call1(__e, SymFn(symshen_4in_1_6), let__5051)


let__5053 := tmp6453
_ = let__5053

tmp6454 := Call2(__e, SymFn(symshen_4comb), let__5053, let__5052)


ifres6450 = tmp6454


}

let__5050 := ifres6450
_ = let__5050

tmp6475 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5050)


if True == tmp6475 {
tmp6456 := Call1(__e, SymFn(symshen_4_5float_6), V2996)


let__5055 := tmp6456
_ = let__5055

tmp6462 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5055)


var ifres6457 Obj

if True == tmp6462 {
tmp6458 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6457 = tmp6458


} else {
tmp6459 := Call1(__e, SymFn(symshen_4_5_1out), let__5055)


let__5056 := tmp6459
_ = let__5056

tmp6460 := Call1(__e, SymFn(symshen_4in_1_6), let__5055)


let__5057 := tmp6460
_ = let__5057

tmp6461 := Call2(__e, SymFn(symshen_4comb), let__5057, let__5056)


ifres6457 = tmp6461


}

let__5054 := ifres6457
_ = let__5054

tmp6473 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5054)


if True == tmp6473 {
tmp6463 := Call1(__e, SymFn(symshen_4_5integer_6), V2996)


let__5059 := tmp6463
_ = let__5059

tmp6469 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5059)


var ifres6464 Obj

if True == tmp6469 {
tmp6465 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6464 = tmp6465


} else {
tmp6466 := Call1(__e, SymFn(symshen_4_5_1out), let__5059)


let__5060 := tmp6466
_ = let__5060

tmp6467 := Call1(__e, SymFn(symshen_4in_1_6), let__5059)


let__5061 := tmp6467
_ = let__5061

tmp6468 := Call2(__e, SymFn(symshen_4comb), let__5061, let__5060)


ifres6464 = tmp6468


}

let__5058 := ifres6464
_ = let__5058

tmp6471 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5058)


if True == tmp6471 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5058
}


} else {
return let__5054
}


} else {
return let__5050
}


} else {
return let__5044
}


} else {
return let__5038
}


panic("unreachable")
})

tmp6480 := Call2(__e, ns2_1set, symshen_4_5number_6, tmp6423)


_ = tmp6480

tmp6481 := MakeNative1(func(__e *ControlFlow, V3021 Obj) Obj {
tmp6486 := Call2(__e, SymFn(symshen_4hds_a_2), V3021, MakeInteger(45))


var ifres6482 Obj

if True == tmp6486 {
tmp6483 := Call1(__e, SymFn(symtail), V3021)


let__5063 := tmp6483
_ = let__5063

tmp6484 := Call2(__e, SymFn(symshen_4comb), let__5063, symshen_4skip)


ifres6482 = tmp6484


} else {
tmp6485 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6482 = tmp6485


}

let__5062 := ifres6482
_ = let__5062

tmp6488 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5062)


if True == tmp6488 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5062
}


panic("unreachable")
})

tmp6489 := Call2(__e, ns2_1set, symshen_4_5minus_6, tmp6481)


_ = tmp6489

tmp6490 := MakeNative1(func(__e *ControlFlow, V3024 Obj) Obj {
tmp6495 := Call2(__e, SymFn(symshen_4hds_a_2), V3024, MakeInteger(43))


var ifres6491 Obj

if True == tmp6495 {
tmp6492 := Call1(__e, SymFn(symtail), V3024)


let__5065 := tmp6492
_ = let__5065

tmp6493 := Call2(__e, SymFn(symshen_4comb), let__5065, symshen_4skip)


ifres6491 = tmp6493


} else {
tmp6494 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6491 = tmp6494


}

let__5064 := ifres6491
_ = let__5064

tmp6497 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5064)


if True == tmp6497 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5064
}


panic("unreachable")
})

tmp6498 := Call2(__e, ns2_1set, symshen_4_5plus_6, tmp6490)


_ = tmp6498

tmp6499 := MakeNative1(func(__e *ControlFlow, V3027 Obj) Obj {
tmp6500 := Call1(__e, SymFn(symshen_4_5digits_6), V3027)


let__5067 := tmp6500
_ = let__5067

tmp6507 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5067)


var ifres6501 Obj

if True == tmp6507 {
tmp6502 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6501 = tmp6502


} else {
tmp6503 := Call1(__e, SymFn(symshen_4_5_1out), let__5067)


let__5068 := tmp6503
_ = let__5068

tmp6504 := Call1(__e, SymFn(symshen_4in_1_6), let__5067)


let__5069 := tmp6504
_ = let__5069

tmp6505 := Call1(__e, SymFn(symshen_4compute_1integer), let__5068)


tmp6506 := Call2(__e, SymFn(symshen_4comb), let__5069, tmp6505)


ifres6501 = tmp6506


}

let__5066 := ifres6501
_ = let__5066

tmp6509 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5066)


if True == tmp6509 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5066
}


panic("unreachable")
})

tmp6510 := Call2(__e, ns2_1set, symshen_4_5integer_6, tmp6499)


_ = tmp6510

tmp6511 := MakeNative1(func(__e *ControlFlow, V3032 Obj) Obj {
tmp6512 := Call1(__e, SymFn(symshen_4_5digit_6), V3032)


let__5071 := tmp6512
_ = let__5071

tmp6525 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5071)


var ifres6513 Obj

if True == tmp6525 {
tmp6514 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6513 = tmp6514


} else {
tmp6515 := Call1(__e, SymFn(symshen_4_5_1out), let__5071)


let__5072 := tmp6515
_ = let__5072

tmp6516 := Call1(__e, SymFn(symshen_4in_1_6), let__5071)


let__5073 := tmp6516
_ = let__5073

tmp6517 := Call1(__e, SymFn(symshen_4_5digits_6), let__5073)


let__5074 := tmp6517
_ = let__5074

tmp6524 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5074)


var ifres6518 Obj

if True == tmp6524 {
tmp6519 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6518 = tmp6519


} else {
tmp6520 := Call1(__e, SymFn(symshen_4_5_1out), let__5074)


let__5075 := tmp6520
_ = let__5075

tmp6521 := Call1(__e, SymFn(symshen_4in_1_6), let__5074)


let__5076 := tmp6521
_ = let__5076

tmp6522 := GCons(__e, let__5072, let__5075)

tmp6523 := Call2(__e, SymFn(symshen_4comb), let__5076, tmp6522)


ifres6518 = tmp6523


}

ifres6513 = ifres6518


}

let__5070 := ifres6513
_ = let__5070

tmp6537 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5070)


if True == tmp6537 {
tmp6526 := Call1(__e, SymFn(symshen_4_5digit_6), V3032)


let__5078 := tmp6526
_ = let__5078

tmp6533 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5078)


var ifres6527 Obj

if True == tmp6533 {
tmp6528 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6527 = tmp6528


} else {
tmp6529 := Call1(__e, SymFn(symshen_4_5_1out), let__5078)


let__5079 := tmp6529
_ = let__5079

tmp6530 := Call1(__e, SymFn(symshen_4in_1_6), let__5078)


let__5080 := tmp6530
_ = let__5080

tmp6531 := GCons(__e, let__5079, Nil)

tmp6532 := Call2(__e, SymFn(symshen_4comb), let__5080, tmp6531)


ifres6527 = tmp6532


}

let__5077 := ifres6527
_ = let__5077

tmp6535 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5077)


if True == tmp6535 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5077
}


} else {
return let__5070
}


panic("unreachable")
})

tmp6538 := Call2(__e, ns2_1set, symshen_4_5digits_6, tmp6511)


_ = tmp6538

tmp6539 := MakeNative1(func(__e *ControlFlow, V3044 Obj) Obj {
tmp6549 := GConsP(__e, V3044)

var ifres6540 Obj

if True == tmp6549 {
tmp6541 := Call1(__e, SymFn(symhead), V3044)


let__5082 := tmp6541
_ = let__5082

tmp6542 := Call1(__e, SymFn(symtail), V3044)


let__5083 := tmp6542
_ = let__5083

tmp6547 := Call1(__e, SymFn(symshen_4digit_2), let__5082)


var ifres6543 Obj

if True == tmp6547 {
tmp6544 := Call1(__e, SymFn(symshen_4byte_1_6digit), let__5082)


tmp6545 := Call2(__e, SymFn(symshen_4comb), let__5083, tmp6544)


ifres6543 = tmp6545


} else {
tmp6546 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6543 = tmp6546


}

ifres6540 = ifres6543


} else {
tmp6548 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6540 = tmp6548


}

let__5081 := ifres6540
_ = let__5081

tmp6551 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5081)


if True == tmp6551 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5081
}


panic("unreachable")
})

tmp6552 := Call2(__e, ns2_1set, symshen_4_5digit_6, tmp6539)


_ = tmp6552

tmp6553 := MakeNative1(func(__e *ControlFlow, V3048 Obj) Obj {
return GSub(__e, V3048, MakeInteger(48))
panic("unreachable")
})

tmp6554 := Call2(__e, ns2_1set, symshen_4byte_1_6digit, tmp6553)


_ = tmp6554

tmp6555 := MakeNative1(func(__e *ControlFlow, V3049 Obj) Obj {
tmp6556 := Call1(__e, SymFn(symreverse), V3049)


return TailCall2(__e, SymFn(symshen_4compute_1integer_1h), tmp6556, MakeInteger(0))


panic("unreachable")
})

tmp6557 := Call2(__e, ns2_1set, symshen_4compute_1integer, tmp6555)


_ = tmp6557

tmp6558 := MakeNative2(func(__e *ControlFlow, V3052, V3053 Obj) Obj {
tmp6568 := GEq(__e, Nil, V3052)

if True == tmp6568 {
return MakeInteger(0)
} else {
tmp6566 := GConsP(__e, V3052)

if True == tmp6566 {
tmp6559 := Call2(__e, SymFn(symshen_4expt), MakeInteger(10), V3053)


tmp6560 := GHd(__e, V3052)

tmp6561 := GMul(__e, tmp6559, tmp6560)

tmp6562 := GTl(__e, V3052)

tmp6564 := Call2(__e, SymFn(symshen_4compute_1integer_1h), tmp6562, GAdd(__e, V3053, MakeInteger(1)))


return GAdd(__e, tmp6561, tmp6564)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4compute_1integer_1h)
}


}


panic("unreachable")
})

tmp6569 := Call2(__e, ns2_1set, symshen_4compute_1integer_1h, tmp6558)


_ = tmp6569

tmp6570 := MakeNative2(func(__e *ControlFlow, V3056, V3057 Obj) Obj {
tmp6578 := GEq(__e, MakeInteger(0), V3057)

if True == tmp6578 {
return MakeInteger(1)
} else {
if True == GGt(__e, V3057, MakeInteger(0)) {
tmp6572 := Call2(__e, SymFn(symshen_4expt), V3056, GSub(__e, V3057, MakeInteger(1)))


return GMul(__e, V3056, tmp6572)


} else {
tmp6574 := Call2(__e, SymFn(symshen_4expt), V3056, GAdd(__e, V3057, MakeInteger(1)))


return GDiv(__e, tmp6574, V3056)


}


}


panic("unreachable")
})

tmp6579 := Call2(__e, ns2_1set, symshen_4expt, tmp6570)


_ = tmp6579

tmp6580 := MakeNative1(func(__e *ControlFlow, V3058 Obj) Obj {
tmp6581 := Call1(__e, SymFn(symshen_4_5integer_6), V3058)


let__5085 := tmp6581
_ = let__5085

tmp6599 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5085)


var ifres6582 Obj

if True == tmp6599 {
tmp6583 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6582 = tmp6583


} else {
tmp6584 := Call1(__e, SymFn(symshen_4_5_1out), let__5085)


let__5086 := tmp6584
_ = let__5086

tmp6585 := Call1(__e, SymFn(symshen_4in_1_6), let__5085)


let__5087 := tmp6585
_ = let__5087

tmp6586 := Call1(__e, SymFn(symshen_4_5stop_6), let__5087)


let__5088 := tmp6586
_ = let__5088

tmp6598 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5088)


var ifres6587 Obj

if True == tmp6598 {
tmp6588 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6587 = tmp6588


} else {
tmp6589 := Call1(__e, SymFn(symshen_4in_1_6), let__5088)


let__5089 := tmp6589
_ = let__5089

tmp6590 := Call1(__e, SymFn(symshen_4_5fraction_6), let__5089)


let__5090 := tmp6590
_ = let__5090

tmp6597 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5090)


var ifres6591 Obj

if True == tmp6597 {
tmp6592 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6591 = tmp6592


} else {
tmp6593 := Call1(__e, SymFn(symshen_4_5_1out), let__5090)


let__5091 := tmp6593
_ = let__5091

tmp6594 := Call1(__e, SymFn(symshen_4in_1_6), let__5090)


let__5092 := tmp6594
_ = let__5092

tmp6596 := Call2(__e, SymFn(symshen_4comb), let__5092, GAdd(__e, let__5086, let__5091))


ifres6591 = tmp6596


}

ifres6587 = ifres6591


}

ifres6582 = ifres6587


}

let__5084 := ifres6582
_ = let__5084

tmp6615 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5084)


if True == tmp6615 {
tmp6600 := Call1(__e, SymFn(symshen_4_5stop_6), V3058)


let__5094 := tmp6600
_ = let__5094

tmp6611 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5094)


var ifres6601 Obj

if True == tmp6611 {
tmp6602 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6601 = tmp6602


} else {
tmp6603 := Call1(__e, SymFn(symshen_4in_1_6), let__5094)


let__5095 := tmp6603
_ = let__5095

tmp6604 := Call1(__e, SymFn(symshen_4_5fraction_6), let__5095)


let__5096 := tmp6604
_ = let__5096

tmp6610 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5096)


var ifres6605 Obj

if True == tmp6610 {
tmp6606 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6605 = tmp6606


} else {
tmp6607 := Call1(__e, SymFn(symshen_4_5_1out), let__5096)


let__5097 := tmp6607
_ = let__5097

tmp6608 := Call1(__e, SymFn(symshen_4in_1_6), let__5096)


let__5098 := tmp6608
_ = let__5098

tmp6609 := Call2(__e, SymFn(symshen_4comb), let__5098, let__5097)


ifres6605 = tmp6609


}

ifres6601 = ifres6605


}

let__5093 := ifres6601
_ = let__5093

tmp6613 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5093)


if True == tmp6613 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5093
}


} else {
return let__5084
}


panic("unreachable")
})

tmp6616 := Call2(__e, ns2_1set, symshen_4_5float_6, tmp6580)


_ = tmp6616

tmp6617 := MakeNative1(func(__e *ControlFlow, V3074 Obj) Obj {
tmp6622 := Call2(__e, SymFn(symshen_4hds_a_2), V3074, MakeInteger(46))


var ifres6618 Obj

if True == tmp6622 {
tmp6619 := Call1(__e, SymFn(symtail), V3074)


let__5100 := tmp6619
_ = let__5100

tmp6620 := Call2(__e, SymFn(symshen_4comb), let__5100, symshen_4skip)


ifres6618 = tmp6620


} else {
tmp6621 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6618 = tmp6621


}

let__5099 := ifres6618
_ = let__5099

tmp6624 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5099)


if True == tmp6624 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5099
}


panic("unreachable")
})

tmp6625 := Call2(__e, ns2_1set, symshen_4_5stop_6, tmp6617)


_ = tmp6625

tmp6626 := MakeNative1(func(__e *ControlFlow, V3077 Obj) Obj {
tmp6627 := Call1(__e, SymFn(symshen_4_5digits_6), V3077)


let__5102 := tmp6627
_ = let__5102

tmp6634 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5102)


var ifres6628 Obj

if True == tmp6634 {
tmp6629 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6628 = tmp6629


} else {
tmp6630 := Call1(__e, SymFn(symshen_4_5_1out), let__5102)


let__5103 := tmp6630
_ = let__5103

tmp6631 := Call1(__e, SymFn(symshen_4in_1_6), let__5102)


let__5104 := tmp6631
_ = let__5104

tmp6632 := Call1(__e, SymFn(symshen_4compute_1fraction), let__5103)


tmp6633 := Call2(__e, SymFn(symshen_4comb), let__5104, tmp6632)


ifres6628 = tmp6633


}

let__5101 := ifres6628
_ = let__5101

tmp6636 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5101)


if True == tmp6636 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5101
}


panic("unreachable")
})

tmp6637 := Call2(__e, ns2_1set, symshen_4_5fraction_6, tmp6626)


_ = tmp6637

tmp6638 := MakeNative1(func(__e *ControlFlow, V3082 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4compute_1fraction_1h), V3082, MakeInteger(-1))
panic("unreachable")
})

tmp6639 := Call2(__e, ns2_1set, symshen_4compute_1fraction, tmp6638)


_ = tmp6639

tmp6640 := MakeNative2(func(__e *ControlFlow, V3085, V3086 Obj) Obj {
tmp6650 := GEq(__e, Nil, V3085)

if True == tmp6650 {
return MakeInteger(0)
} else {
tmp6648 := GConsP(__e, V3085)

if True == tmp6648 {
tmp6641 := Call2(__e, SymFn(symshen_4expt), MakeInteger(10), V3086)


tmp6642 := GHd(__e, V3085)

tmp6643 := GMul(__e, tmp6641, tmp6642)

tmp6644 := GTl(__e, V3085)

tmp6646 := Call2(__e, SymFn(symshen_4compute_1fraction_1h), tmp6644, GSub(__e, V3086, MakeInteger(1)))


return GAdd(__e, tmp6643, tmp6646)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4compute_1fraction_1h)
}


}


panic("unreachable")
})

tmp6651 := Call2(__e, ns2_1set, symshen_4compute_1fraction_1h, tmp6640)


_ = tmp6651

tmp6652 := MakeNative1(func(__e *ControlFlow, V3087 Obj) Obj {
tmp6653 := Call1(__e, SymFn(symshen_4_5float_6), V3087)


let__5106 := tmp6653
_ = let__5106

tmp6671 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5106)


var ifres6654 Obj

if True == tmp6671 {
tmp6655 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6654 = tmp6655


} else {
tmp6656 := Call1(__e, SymFn(symshen_4_5_1out), let__5106)


let__5107 := tmp6656
_ = let__5107

tmp6657 := Call1(__e, SymFn(symshen_4in_1_6), let__5106)


let__5108 := tmp6657
_ = let__5108

tmp6658 := Call1(__e, SymFn(symshen_4_5lowE_6), let__5108)


let__5109 := tmp6658
_ = let__5109

tmp6670 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5109)


var ifres6659 Obj

if True == tmp6670 {
tmp6660 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6659 = tmp6660


} else {
tmp6661 := Call1(__e, SymFn(symshen_4in_1_6), let__5109)


let__5110 := tmp6661
_ = let__5110

tmp6662 := Call1(__e, SymFn(symshen_4_5log10_6), let__5110)


let__5111 := tmp6662
_ = let__5111

tmp6669 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5111)


var ifres6663 Obj

if True == tmp6669 {
tmp6664 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6663 = tmp6664


} else {
tmp6665 := Call1(__e, SymFn(symshen_4_5_1out), let__5111)


let__5112 := tmp6665
_ = let__5112

tmp6666 := Call1(__e, SymFn(symshen_4in_1_6), let__5111)


let__5113 := tmp6666
_ = let__5113

tmp6667 := Call2(__e, SymFn(symshen_4compute_1E), let__5107, let__5112)


tmp6668 := Call2(__e, SymFn(symshen_4comb), let__5113, tmp6667)


ifres6663 = tmp6668


}

ifres6659 = ifres6663


}

ifres6654 = ifres6659


}

let__5105 := ifres6654
_ = let__5105

tmp6694 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5105)


if True == tmp6694 {
tmp6672 := Call1(__e, SymFn(symshen_4_5integer_6), V3087)


let__5115 := tmp6672
_ = let__5115

tmp6690 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5115)


var ifres6673 Obj

if True == tmp6690 {
tmp6674 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6673 = tmp6674


} else {
tmp6675 := Call1(__e, SymFn(symshen_4_5_1out), let__5115)


let__5116 := tmp6675
_ = let__5116

tmp6676 := Call1(__e, SymFn(symshen_4in_1_6), let__5115)


let__5117 := tmp6676
_ = let__5117

tmp6677 := Call1(__e, SymFn(symshen_4_5lowE_6), let__5117)


let__5118 := tmp6677
_ = let__5118

tmp6689 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5118)


var ifres6678 Obj

if True == tmp6689 {
tmp6679 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6678 = tmp6679


} else {
tmp6680 := Call1(__e, SymFn(symshen_4in_1_6), let__5118)


let__5119 := tmp6680
_ = let__5119

tmp6681 := Call1(__e, SymFn(symshen_4_5log10_6), let__5119)


let__5120 := tmp6681
_ = let__5120

tmp6688 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5120)


var ifres6682 Obj

if True == tmp6688 {
tmp6683 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6682 = tmp6683


} else {
tmp6684 := Call1(__e, SymFn(symshen_4_5_1out), let__5120)


let__5121 := tmp6684
_ = let__5121

tmp6685 := Call1(__e, SymFn(symshen_4in_1_6), let__5120)


let__5122 := tmp6685
_ = let__5122

tmp6686 := Call2(__e, SymFn(symshen_4compute_1E), let__5116, let__5121)


tmp6687 := Call2(__e, SymFn(symshen_4comb), let__5122, tmp6686)


ifres6682 = tmp6687


}

ifres6678 = ifres6682


}

ifres6673 = ifres6678


}

let__5114 := ifres6673
_ = let__5114

tmp6692 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5114)


if True == tmp6692 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5114
}


} else {
return let__5105
}


panic("unreachable")
})

tmp6695 := Call2(__e, ns2_1set, symshen_4_5e_1number_6, tmp6652)


_ = tmp6695

tmp6696 := MakeNative1(func(__e *ControlFlow, V3106 Obj) Obj {
tmp6697 := Call1(__e, SymFn(symshen_4_5plus_6), V3106)


let__5124 := tmp6697
_ = let__5124

tmp6708 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5124)


var ifres6698 Obj

if True == tmp6708 {
tmp6699 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6698 = tmp6699


} else {
tmp6700 := Call1(__e, SymFn(symshen_4in_1_6), let__5124)


let__5125 := tmp6700
_ = let__5125

tmp6701 := Call1(__e, SymFn(symshen_4_5log10_6), let__5125)


let__5126 := tmp6701
_ = let__5126

tmp6707 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5126)


var ifres6702 Obj

if True == tmp6707 {
tmp6703 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6702 = tmp6703


} else {
tmp6704 := Call1(__e, SymFn(symshen_4_5_1out), let__5126)


let__5127 := tmp6704
_ = let__5127

tmp6705 := Call1(__e, SymFn(symshen_4in_1_6), let__5126)


let__5128 := tmp6705
_ = let__5128

tmp6706 := Call2(__e, SymFn(symshen_4comb), let__5128, let__5127)


ifres6702 = tmp6706


}

ifres6698 = ifres6702


}

let__5123 := ifres6698
_ = let__5123

tmp6734 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5123)


if True == tmp6734 {
tmp6709 := Call1(__e, SymFn(symshen_4_5minus_6), V3106)


let__5130 := tmp6709
_ = let__5130

tmp6721 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5130)


var ifres6710 Obj

if True == tmp6721 {
tmp6711 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6710 = tmp6711


} else {
tmp6712 := Call1(__e, SymFn(symshen_4in_1_6), let__5130)


let__5131 := tmp6712
_ = let__5131

tmp6713 := Call1(__e, SymFn(symshen_4_5log10_6), let__5131)


let__5132 := tmp6713
_ = let__5132

tmp6720 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5132)


var ifres6714 Obj

if True == tmp6720 {
tmp6715 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6714 = tmp6715


} else {
tmp6716 := Call1(__e, SymFn(symshen_4_5_1out), let__5132)


let__5133 := tmp6716
_ = let__5133

tmp6717 := Call1(__e, SymFn(symshen_4in_1_6), let__5132)


let__5134 := tmp6717
_ = let__5134

tmp6719 := Call2(__e, SymFn(symshen_4comb), let__5134, GSub(__e, MakeInteger(0), let__5133))


ifres6714 = tmp6719


}

ifres6710 = ifres6714


}

let__5129 := ifres6710
_ = let__5129

tmp6732 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5129)


if True == tmp6732 {
tmp6722 := Call1(__e, SymFn(symshen_4_5integer_6), V3106)


let__5136 := tmp6722
_ = let__5136

tmp6728 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5136)


var ifres6723 Obj

if True == tmp6728 {
tmp6724 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6723 = tmp6724


} else {
tmp6725 := Call1(__e, SymFn(symshen_4_5_1out), let__5136)


let__5137 := tmp6725
_ = let__5137

tmp6726 := Call1(__e, SymFn(symshen_4in_1_6), let__5136)


let__5138 := tmp6726
_ = let__5138

tmp6727 := Call2(__e, SymFn(symshen_4comb), let__5138, let__5137)


ifres6723 = tmp6727


}

let__5135 := ifres6723
_ = let__5135

tmp6730 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5135)


if True == tmp6730 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5135
}


} else {
return let__5129
}


} else {
return let__5123
}


panic("unreachable")
})

tmp6735 := Call2(__e, ns2_1set, symshen_4_5log10_6, tmp6696)


_ = tmp6735

tmp6736 := MakeNative1(func(__e *ControlFlow, V3123 Obj) Obj {
tmp6741 := Call2(__e, SymFn(symshen_4hds_a_2), V3123, MakeInteger(101))


var ifres6737 Obj

if True == tmp6741 {
tmp6738 := Call1(__e, SymFn(symtail), V3123)


let__5140 := tmp6738
_ = let__5140

tmp6739 := Call2(__e, SymFn(symshen_4comb), let__5140, symshen_4skip)


ifres6737 = tmp6739


} else {
tmp6740 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6737 = tmp6740


}

let__5139 := ifres6737
_ = let__5139

tmp6743 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5139)


if True == tmp6743 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5139
}


panic("unreachable")
})

tmp6744 := Call2(__e, ns2_1set, symshen_4_5lowE_6, tmp6736)


_ = tmp6744

tmp6745 := MakeNative2(func(__e *ControlFlow, V3126, V3127 Obj) Obj {
tmp6746 := Call2(__e, SymFn(symshen_4expt), MakeInteger(10), V3127)


return GMul(__e, V3126, tmp6746)


panic("unreachable")
})

tmp6747 := Call2(__e, ns2_1set, symshen_4compute_1E, tmp6745)


_ = tmp6747

tmp6748 := MakeNative1(func(__e *ControlFlow, V3128 Obj) Obj {
tmp6749 := Call1(__e, SymFn(symshen_4_5whitespace_6), V3128)


let__5142 := tmp6749
_ = let__5142

tmp6759 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5142)


var ifres6750 Obj

if True == tmp6759 {
tmp6751 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6750 = tmp6751


} else {
tmp6752 := Call1(__e, SymFn(symshen_4in_1_6), let__5142)


let__5143 := tmp6752
_ = let__5143

tmp6753 := Call1(__e, SymFn(symshen_4_5whitespaces_6), let__5143)


let__5144 := tmp6753
_ = let__5144

tmp6758 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5144)


var ifres6754 Obj

if True == tmp6758 {
tmp6755 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6754 = tmp6755


} else {
tmp6756 := Call1(__e, SymFn(symshen_4in_1_6), let__5144)


let__5145 := tmp6756
_ = let__5145

tmp6757 := Call2(__e, SymFn(symshen_4comb), let__5145, symshen_4skip)


ifres6754 = tmp6757


}

ifres6750 = ifres6754


}

let__5141 := ifres6750
_ = let__5141

tmp6769 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5141)


if True == tmp6769 {
tmp6760 := Call1(__e, SymFn(symshen_4_5whitespace_6), V3128)


let__5147 := tmp6760
_ = let__5147

tmp6765 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5147)


var ifres6761 Obj

if True == tmp6765 {
tmp6762 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6761 = tmp6762


} else {
tmp6763 := Call1(__e, SymFn(symshen_4in_1_6), let__5147)


let__5148 := tmp6763
_ = let__5148

tmp6764 := Call2(__e, SymFn(symshen_4comb), let__5148, symshen_4skip)


ifres6761 = tmp6764


}

let__5146 := ifres6761
_ = let__5146

tmp6767 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5146)


if True == tmp6767 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5146
}


} else {
return let__5141
}


panic("unreachable")
})

tmp6770 := Call2(__e, ns2_1set, symshen_4_5whitespaces_6, tmp6748)


_ = tmp6770

tmp6771 := MakeNative1(func(__e *ControlFlow, V3137 Obj) Obj {
tmp6780 := GConsP(__e, V3137)

var ifres6772 Obj

if True == tmp6780 {
tmp6773 := Call1(__e, SymFn(symhead), V3137)


let__5150 := tmp6773
_ = let__5150

tmp6774 := Call1(__e, SymFn(symtail), V3137)


let__5151 := tmp6774
_ = let__5151

tmp6778 := Call1(__e, SymFn(symshen_4whitespace_2), let__5150)


var ifres6775 Obj

if True == tmp6778 {
tmp6776 := Call2(__e, SymFn(symshen_4comb), let__5151, symshen_4skip)


ifres6775 = tmp6776


} else {
tmp6777 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6775 = tmp6777


}

ifres6772 = ifres6775


} else {
tmp6779 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6772 = tmp6779


}

let__5149 := ifres6772
_ = let__5149

tmp6782 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5149)


if True == tmp6782 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5149
}


panic("unreachable")
})

tmp6783 := Call2(__e, ns2_1set, symshen_4_5whitespace_6, tmp6771)


_ = tmp6783

tmp6784 := MakeNative1(func(__e *ControlFlow, V3143 Obj) Obj {
tmp6792 := GEq(__e, MakeInteger(32), V3143)

if True == tmp6792 {
return True
} else {
tmp6790 := GEq(__e, MakeInteger(13), V3143)

if True == tmp6790 {
return True
} else {
tmp6788 := GEq(__e, MakeInteger(10), V3143)

if True == tmp6788 {
return True
} else {
tmp6786 := GEq(__e, MakeInteger(9), V3143)

if True == tmp6786 {
return True
} else {
return False
}


}


}


}


panic("unreachable")
})

tmp6793 := Call2(__e, ns2_1set, symshen_4whitespace_2, tmp6784)


_ = tmp6793

tmp6794 := MakeNative2(func(__e *ControlFlow, V3146, V3147 Obj) Obj {
tmp6807 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V3146)

var ifres6801 Obj

if True == tmp6807 {
tmp6805 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V3146)

tmp6806 := GNot(__e, tmp6805)

var ifres6802 Obj

if True == tmp6806 {
tmp6803 := GEq(__e, V3146, sympackage)

ifres6802 = GNot(__e, tmp6803)


} else {
ifres6802 = False


}

ifres6801 = ifres6802


} else {
ifres6801 = False


}

if True == ifres6801 {
tmp6795 := GValue(__e, symshen_4_dsource_1form_1handlers_d)

tmp6796 := Call3(__e, SymFn(symshen_4update_1assoc), V3146, V3147, tmp6795)


tmp6797 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsource_1form_1handlers_d, tmp6796)

_ = tmp6797

return V3146


} else {
tmp6798 := Call3(__e, SymFn(symshen_4app), V3146, MakeString("\n"), symshen_4s)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("invalid source-form head: "), tmp6798))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("invalid source-form head: "), tmp6798)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


panic("unreachable")
})

tmp6808 := Call2(__e, ns2_1set, symshen_4register_1source_1form, tmp6794)


_ = tmp6808

tmp6809 := MakeNative1(func(__e *ControlFlow, V3148 Obj) Obj {
tmp6810 := GValue(__e, symshen_4_dsource_1form_1handlers_d)

tmp6811 := Call2(__e, SymFn(symassoc), V3148, tmp6810)


tmp6812 := GValue(__e, symshen_4_dsource_1form_1handlers_d)

tmp6813 := Call2(__e, SymFn(symremove), tmp6811, tmp6812)


tmp6814 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsource_1form_1handlers_d, tmp6813)

_ = tmp6814

return V3148


panic("unreachable")
})

tmp6815 := Call2(__e, ns2_1set, symshen_4unregister_1source_1form, tmp6809)


_ = tmp6815

tmp6816 := MakeNative2(func(__e *ControlFlow, V3155, V3156 Obj) Obj {
tmp6819 := GConsP(__e, V3155)

if True == tmp6819 {
tmp6817 := GHd(__e, V3155)

return TailCall2(__e, SymFn(symassoc), tmp6817, V3156)


} else {
return Nil
}


panic("unreachable")
})

tmp6820 := Call2(__e, ns2_1set, symshen_4source_1form_1handler, tmp6816)


_ = tmp6820

tmp6821 := MakeNative2(func(__e *ControlFlow, V3159, V3160 Obj) Obj {
tmp6829 := GConsP(__e, V3160)

if True == tmp6829 {
tmp6822 := GTl(__e, V3160)

tmp6823 := Call1(__e, tmp6822, V3159)


let__5152 := tmp6823
_ = let__5152

tmp6827 := Call1(__e, SymFn(symshen_4source_1form_1list_2), let__5152)


if True == tmp6827 {
return let__5152
} else {
tmp6824 := Call3(__e, SymFn(symshen_4app), V3159, MakeString("\n"), symshen_4s)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("source-form handler must return a list of forms: "), tmp6824))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("source-form handler must return a list of forms: "), tmp6824)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4source_1form_1expansion)
}


panic("unreachable")
})

tmp6830 := Call2(__e, ns2_1set, symshen_4source_1form_1expansion, tmp6821)


_ = tmp6830

tmp6831 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V3166 := __self1
_ = V3166
tmp6836 := GEq(__e, Nil, V3166)

if True == tmp6836 {
return True
} else {
tmp6834 := GConsP(__e, V3166)

if True == tmp6834 {
tmp6832 := GTl(__e, V3166)

if SymFn(symshen_4source_1form_1list_2) == __self {
__self1 = tmp6832
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4source_1form_1list_2), tmp6832)


} else {
return False
}


}


panic("unreachable")
})
return __self
})()

tmp6837 := Call2(__e, ns2_1set, symshen_4source_1form_1list_2, tmp6831)


_ = tmp6837

tmp6838 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V3167 := __self1
_ = V3167
tmp6872 := GEq(__e, Nil, V3167)

if True == tmp6872 {
return Nil
} else {
tmp6870 := GConsP(__e, V3167)

var ifres6867 Obj

if True == tmp6870 {
tmp6868 := GHd(__e, V3167)

tmp6869 := Call1(__e, SymFn(symshen_4packaged_2), tmp6868)


ifres6867 = tmp6869


} else {
ifres6867 = False


}

if True == ifres6867 {
tmp6839 := GHd(__e, V3167)

tmp6840 := Call1(__e, SymFn(symshen_4unpackage), tmp6839)


tmp6841 := GTl(__e, V3167)

tmp6842 := Call2(__e, SymFn(symappend), tmp6840, tmp6841)


if SymFn(symshen_4unpackage_emacroexpand) == __self {
__self1 = tmp6842
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4unpackage_emacroexpand), tmp6842)


} else {
tmp6865 := GConsP(__e, V3167)

if True == tmp6865 {
tmp6843 := GHd(__e, V3167)

tmp6844 := GValue(__e, symshen_4_dsource_1form_1handlers_d)

tmp6845 := Call2(__e, SymFn(symshen_4source_1form_1handler), tmp6843, tmp6844)


let__5153 := tmp6845
_ = let__5153

tmp6863 := GConsP(__e, let__5153)

if True == tmp6863 {
tmp6846 := GHd(__e, V3167)

tmp6847 := Call2(__e, SymFn(symshen_4source_1form_1expansion), tmp6846, let__5153)


tmp6848 := GTl(__e, V3167)

tmp6849 := Call2(__e, SymFn(symappend), tmp6847, tmp6848)


if SymFn(symshen_4unpackage_emacroexpand) == __self {
__self1 = tmp6849
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4unpackage_emacroexpand), tmp6849)


} else {
tmp6850 := GHd(__e, V3167)

tmp6851 := Call1(__e, SymFn(symmacroexpand), tmp6850)


let__5154 := tmp6851
_ = let__5154

tmp6861 := Call1(__e, SymFn(symshen_4packaged_2), let__5154)


var ifres6857 Obj

if True == tmp6861 {
ifres6857 = True


} else {
tmp6858 := GValue(__e, symshen_4_dsource_1form_1handlers_d)

tmp6859 := Call2(__e, SymFn(symshen_4source_1form_1handler), let__5154, tmp6858)


tmp6860 := GConsP(__e, tmp6859)

ifres6857 = tmp6860


}

if True == ifres6857 {
tmp6852 := GTl(__e, V3167)

tmp6853 := GCons(__e, let__5154, tmp6852)

if SymFn(symshen_4unpackage_emacroexpand) == __self {
__self1 = tmp6853
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4unpackage_emacroexpand), tmp6853)


} else {
tmp6854 := GTl(__e, V3167)

tmp6855 := Call1(__e, SymFn(symshen_4unpackage_emacroexpand), tmp6854)


return GCons(__e, let__5154, tmp6855)


}


}


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4unpackage_emacroexpand)
}


}


}


panic("unreachable")
})
return __self
})()

tmp6873 := Call2(__e, ns2_1set, symshen_4unpackage_emacroexpand, tmp6838)


_ = tmp6873

tmp6874 := MakeNative1(func(__e *ControlFlow, V3172 Obj) Obj {
tmp6886 := GConsP(__e, V3172)

var ifres6876 Obj

if True == tmp6886 {
tmp6884 := GHd(__e, V3172)

tmp6885 := GEq(__e, sympackage, tmp6884)

var ifres6877 Obj

if True == tmp6885 {
tmp6882 := GTl(__e, V3172)

tmp6883 := GConsP(__e, tmp6882)

var ifres6878 Obj

if True == tmp6883 {
tmp6879 := GTl(__e, V3172)

tmp6880 := GTl(__e, tmp6879)

tmp6881 := GConsP(__e, tmp6880)

ifres6878 = tmp6881


} else {
ifres6878 = False


}

ifres6877 = ifres6878


} else {
ifres6877 = False


}

ifres6876 = ifres6877


} else {
ifres6876 = False


}

if True == ifres6876 {
return True
} else {
return False
}


panic("unreachable")
})

tmp6887 := Call2(__e, ns2_1set, symshen_4packaged_2, tmp6874)


_ = tmp6887

tmp6888 := MakeNative1(func(__e *ControlFlow, V3175 Obj) Obj {
tmp6938 := GConsP(__e, V3175)

var ifres6924 Obj

if True == tmp6938 {
tmp6936 := GHd(__e, V3175)

tmp6937 := GEq(__e, sympackage, tmp6936)

var ifres6925 Obj

if True == tmp6937 {
tmp6934 := GTl(__e, V3175)

tmp6935 := GConsP(__e, tmp6934)

var ifres6926 Obj

if True == tmp6935 {
tmp6931 := GTl(__e, V3175)

tmp6932 := GHd(__e, tmp6931)

tmp6933 := GEq(__e, symnull, tmp6932)

var ifres6927 Obj

if True == tmp6933 {
tmp6928 := GTl(__e, V3175)

tmp6929 := GTl(__e, tmp6928)

tmp6930 := GConsP(__e, tmp6929)

ifres6927 = tmp6930


} else {
ifres6927 = False


}

ifres6926 = ifres6927


} else {
ifres6926 = False


}

ifres6925 = ifres6926


} else {
ifres6925 = False


}

ifres6924 = ifres6925


} else {
ifres6924 = False


}

if True == ifres6924 {
tmp6889 := GTl(__e, V3175)

tmp6890 := GTl(__e, tmp6889)

return GTl(__e, tmp6890)


} else {
tmp6922 := GConsP(__e, V3175)

var ifres6912 Obj

if True == tmp6922 {
tmp6920 := GHd(__e, V3175)

tmp6921 := GEq(__e, sympackage, tmp6920)

var ifres6913 Obj

if True == tmp6921 {
tmp6918 := GTl(__e, V3175)

tmp6919 := GConsP(__e, tmp6918)

var ifres6914 Obj

if True == tmp6919 {
tmp6915 := GTl(__e, V3175)

tmp6916 := GTl(__e, tmp6915)

tmp6917 := GConsP(__e, tmp6916)

ifres6914 = tmp6917


} else {
ifres6914 = False


}

ifres6913 = ifres6914


} else {
ifres6913 = False


}

ifres6912 = ifres6913


} else {
ifres6912 = False


}

if True == ifres6912 {
tmp6891 := GTl(__e, V3175)

tmp6892 := GTl(__e, tmp6891)

tmp6893 := GHd(__e, tmp6892)

tmp6894 := Call1(__e, SymFn(symeval), tmp6893)


let__5155 := tmp6894
_ = let__5155

tmp6895 := GTl(__e, V3175)

tmp6896 := GHd(__e, tmp6895)

tmp6897 := PrimGuard1(__e, symstr, PrimStr, tmp6896)

tmp6898 := GTl(__e, V3175)

tmp6899 := GTl(__e, tmp6898)

tmp6900 := GTl(__e, tmp6899)

tmp6901 := Call3(__e, SymFn(symshen_4package_1symbols), tmp6897, let__5155, tmp6900)


let__5156 := tmp6901
_ = let__5156

tmp6902 := GTl(__e, V3175)

tmp6903 := GHd(__e, tmp6902)

tmp6904 := Call2(__e, SymFn(symshen_4record_1external), tmp6903, let__5155)


let__5157 := tmp6904
_ = let__5157

tmp6905 := GTl(__e, V3175)

tmp6906 := GHd(__e, tmp6905)

tmp6907 := GTl(__e, V3175)

tmp6908 := GTl(__e, tmp6907)

tmp6909 := GTl(__e, tmp6908)

tmp6910 := Call3(__e, SymFn(symshen_4record_1internal), tmp6906, let__5155, tmp6909)


let__5158 := tmp6910
_ = let__5158

return let__5156


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4unpackage)
}


}


panic("unreachable")
})

tmp6939 := Call2(__e, ns2_1set, symshen_4unpackage, tmp6888)


_ = tmp6939

tmp6940 := MakeNative3(func(__e *ControlFlow, V3180, V3181, V3182 Obj) Obj {
tmp6941 := MakeNative0(func(__e *ControlFlow) Obj {
tmp6942 := GValue(__e, sym_dproperty_1vector_d)

return TailCall3(__e, SymFn(symget), V3180, symshen_4internal_1symbols, tmp6942)


panic("unreachable")
})

tmp6943 := MakeNative1(func(__e *ControlFlow, Z3184 Obj) Obj {
return Nil
panic("unreachable")
})

tmp6944 := Call2(__e, try_1catch, tmp6941, tmp6943)


let__5159 := tmp6944
_ = let__5159

tmp6945 := PrimGuard1(__e, symstr, PrimStr, V3180)

tmp6946 := Call3(__e, SymFn(symshen_4internal_1symbols), tmp6945, V3181, V3182)


let__5160 := tmp6946
_ = let__5160

tmp6947 := Call2(__e, SymFn(symunion), let__5160, let__5159)


tmp6948 := GValue(__e, sym_dproperty_1vector_d)

return TailCall4(__e, SymFn(symput), V3180, symshen_4internal_1symbols, tmp6947, tmp6948)


panic("unreachable")
})

tmp6949 := Call2(__e, ns2_1set, symshen_4record_1internal, tmp6940)


_ = tmp6949

tmp6950 := MakeNative3(func(__e *ControlFlow, V3192, V3193, V3194 Obj) Obj {
tmp6959 := GConsP(__e, V3194)

if True == tmp6959 {
tmp6951 := GHd(__e, V3194)

tmp6952 := Call3(__e, SymFn(symshen_4internal_1symbols), V3192, V3193, tmp6951)


tmp6953 := GTl(__e, V3194)

tmp6954 := Call3(__e, SymFn(symshen_4internal_1symbols), V3192, V3193, tmp6953)


return TailCall2(__e, SymFn(symunion), tmp6952, tmp6954)


} else {
tmp6957 := Call3(__e, SymFn(symshen_4internal_2), V3194, V3192, V3193)


if True == tmp6957 {
tmp6955 := Call2(__e, SymFn(symshen_4intern_1in_1package), V3192, V3194)


return GCons(__e, tmp6955, Nil)


} else {
return Nil
}


}


panic("unreachable")
})

tmp6960 := Call2(__e, ns2_1set, symshen_4internal_1symbols, tmp6950)


_ = tmp6960

tmp6961 := MakeNative2(func(__e *ControlFlow, V3195, V3196 Obj) Obj {
tmp6962 := MakeNative0(func(__e *ControlFlow) Obj {
tmp6963 := GValue(__e, sym_dproperty_1vector_d)

return TailCall3(__e, SymFn(symget), V3195, symshen_4external_1symbols, tmp6963)


panic("unreachable")
})

tmp6964 := MakeNative1(func(__e *ControlFlow, Z3198 Obj) Obj {
return Nil
panic("unreachable")
})

tmp6965 := Call2(__e, try_1catch, tmp6962, tmp6964)


let__5161 := tmp6965
_ = let__5161

tmp6966 := Call2(__e, SymFn(symunion), V3196, let__5161)


tmp6967 := GValue(__e, sym_dproperty_1vector_d)

return TailCall4(__e, SymFn(symput), V3195, symshen_4external_1symbols, tmp6966, tmp6967)


panic("unreachable")
})

tmp6968 := Call2(__e, ns2_1set, symshen_4record_1external, tmp6961)


_ = tmp6968

tmp6969 := MakeNative3(func(__e *ControlFlow, V3203, V3204, V3205 Obj) Obj {
tmp6974 := GConsP(__e, V3205)

if True == tmp6974 {
tmp6970 := MakeNative1(func(__e *ControlFlow, Z3206 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4package_1symbols), V3203, V3204, Z3206)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp6970, V3205)


} else {
tmp6972 := Call3(__e, SymFn(symshen_4internal_2), V3205, V3203, V3204)


if True == tmp6972 {
return TailCall2(__e, SymFn(symshen_4intern_1in_1package), V3203, V3205)
} else {
return V3205
}


}


panic("unreachable")
})

tmp6975 := Call2(__e, ns2_1set, symshen_4package_1symbols, tmp6969)


_ = tmp6975

tmp6976 := MakeNative2(func(__e *ControlFlow, V3207, V3208 Obj) Obj {
tmp6977 := PrimGuard1(__e, symstr, PrimStr, V3208)

tmp6978 := Call2(__e, SymFn(sym_8s), MakeString("."), tmp6977)


tmp6979 := Call2(__e, SymFn(sym_8s), V3207, tmp6978)


return PrimGuard1(__e, symintern, PrimIntern, tmp6979)


panic("unreachable")
})

tmp6980 := Call2(__e, ns2_1set, symshen_4intern_1in_1package, tmp6976)


_ = tmp6980

tmp6981 := MakeNative3(func(__e *ControlFlow, V3209, V3210, V3211 Obj) Obj {
tmp7003 := Call2(__e, SymFn(symelement_2), V3209, V3211)


if True == GNot(__e, tmp7003) {
tmp7000 := Call1(__e, SymFn(symshen_4sng_2), V3209)


if True == GNot(__e, tmp7000) {
tmp6997 := Call1(__e, SymFn(symshen_4dbl_2), V3209)


if True == GNot(__e, tmp6997) {
tmp6995 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V3209)

if True == tmp6995 {
tmp6992 := Call1(__e, SymFn(symshen_4sysfunc_2), V3209)


if True == GNot(__e, tmp6992) {
tmp6989 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V3209)

if True == GNot(__e, tmp6989) {
tmp6985 := PrimGuard1(__e, symstr, PrimStr, V3209)

tmp6986 := Call1(__e, SymFn(symshen_4internal_1to_1shen_2), tmp6985)


if True == GNot(__e, tmp6986) {
tmp6982 := PrimGuard1(__e, symstr, PrimStr, V3209)

tmp6983 := Call2(__e, SymFn(symshen_4internal_1to_1P_2), V3210, tmp6982)


return GNot(__e, tmp6983)


} else {
return False
}


} else {
return False
}


} else {
return False
}


} else {
return False
}


} else {
return False
}


} else {
return False
}


} else {
return False
}


panic("unreachable")
})

tmp7005 := Call2(__e, ns2_1set, symshen_4internal_2, tmp6981)


_ = tmp7005

tmp7006 := MakeNative1(func(__e *ControlFlow, V3216 Obj) Obj {
tmp7051 := Call1(__e, SymFn(symshen_4_7string_2), V3216)


var ifres7008 Obj

if True == tmp7051 {
tmp7049 := Call1(__e, SymFn(symhdstr), V3216)


tmp7050 := GEq(__e, MakeString("s"), tmp7049)

var ifres7009 Obj

if True == tmp7050 {
tmp7048 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, V3216))


var ifres7010 Obj

if True == tmp7048 {
tmp7045 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, V3216))


tmp7046 := GEq(__e, MakeString("h"), tmp7045)

var ifres7011 Obj

if True == tmp7046 {
tmp7043 := Call1(__e, SymFn(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V3216)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})())


var ifres7012 Obj

if True == tmp7043 {
tmp7039 := Call1(__e, SymFn(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V3216)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})())


tmp7040 := GEq(__e, MakeString("e"), tmp7039)

var ifres7013 Obj

if True == tmp7040 {
tmp7036 := Call1(__e, SymFn(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(TypedStringTailValue(__typedS0))))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V3216)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})()
return Call1(__e, SymFn(symtlstr), __typedArg0)
})())


var ifres7014 Obj

if True == tmp7036 {
tmp7031 := Call1(__e, SymFn(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(TypedStringTailValue(__typedS0))))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V3216)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})()
return Call1(__e, SymFn(symtlstr), __typedArg0)
})())


tmp7032 := GEq(__e, MakeString("n"), tmp7031)

var ifres7015 Obj

if True == tmp7032 {
tmp7027 := Call1(__e, SymFn(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(TypedStringTailValue(TypedStringTailValue(__typedS0)))))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(TypedStringTailValue(__typedS0))))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V3216)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})()
return Call1(__e, SymFn(symtlstr), __typedArg0)
})()
return Call1(__e, SymFn(symtlstr), __typedArg0)
})())


var ifres7016 Obj

if True == tmp7027 {
tmp7021 := Call1(__e, SymFn(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(TypedStringTailValue(TypedStringTailValue(__typedS0)))))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(TypedStringTailValue(__typedS0))))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V3216)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})()
return Call1(__e, SymFn(symtlstr), __typedArg0)
})()
return Call1(__e, SymFn(symtlstr), __typedArg0)
})())


tmp7022 := GEq(__e, MakeString("."), tmp7021)

ifres7016 = tmp7022


} else {
ifres7016 = False


}

ifres7015 = ifres7016


} else {
ifres7015 = False


}

ifres7014 = ifres7015


} else {
ifres7014 = False


}

ifres7013 = ifres7014


} else {
ifres7013 = False


}

ifres7012 = ifres7013


} else {
ifres7012 = False


}

ifres7011 = ifres7012


} else {
ifres7011 = False


}

ifres7010 = ifres7011


} else {
ifres7010 = False


}

ifres7009 = ifres7010


} else {
ifres7009 = False


}

ifres7008 = ifres7009


} else {
ifres7008 = False


}

if True == ifres7008 {
return True
} else {
return False
}


panic("unreachable")
})

tmp7052 := Call2(__e, ns2_1set, symshen_4internal_1to_1shen_2, tmp7006)


_ = tmp7052

tmp7053 := MakeNative1(func(__e *ControlFlow, V3217 Obj) Obj {
tmp7054 := GValue(__e, sym_dproperty_1vector_d)

tmp7055 := Call3(__e, SymFn(symget), symshen, symshen_4external_1symbols, tmp7054)


return TailCall2(__e, SymFn(symelement_2), V3217, tmp7055)


panic("unreachable")
})

tmp7056 := Call2(__e, ns2_1set, symshen_4sysfunc_2, tmp7053)


_ = tmp7056

tmp7057 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V3225 := __self1
_ = V3225
V3226 := __self2
_ = V3226
tmp7074 := GEq(__e, MakeString(""), V3225)

var ifres7069 Obj

if True == tmp7074 {
tmp7073 := Call1(__e, SymFn(symshen_4_7string_2), V3226)


var ifres7070 Obj

if True == tmp7073 {
tmp7071 := Call1(__e, SymFn(symhdstr), V3226)


tmp7072 := GEq(__e, MakeString("."), tmp7071)

ifres7070 = tmp7072


} else {
ifres7070 = False


}

ifres7069 = ifres7070


} else {
ifres7069 = False


}

if True == ifres7069 {
return True
} else {
tmp7067 := Call1(__e, SymFn(symshen_4_7string_2), V3225)


var ifres7061 Obj

if True == tmp7067 {
tmp7066 := Call1(__e, SymFn(symshen_4_7string_2), V3226)


var ifres7062 Obj

if True == tmp7066 {
tmp7063 := Call1(__e, SymFn(symhdstr), V3225)


tmp7064 := Call1(__e, SymFn(symhdstr), V3226)


tmp7065 := GEq(__e, tmp7063, tmp7064)

ifres7062 = tmp7065


} else {
ifres7062 = False


}

ifres7061 = ifres7062


} else {
ifres7061 = False


}

if True == ifres7061 {
tmp7058 := PrimGuard1(__e, symtlstr, PrimTailString, V3225)

if SymFn(symshen_4internal_1to_1P_2) == __self {
__self1, __self2 = tmp7058, PrimGuard1(__e, symtlstr, PrimTailString, V3226)
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4internal_1to_1P_2), tmp7058, PrimGuard1(__e, symtlstr, PrimTailString, V3226))


} else {
return False
}


}


panic("unreachable")
})
return __self
})()

tmp7075 := Call2(__e, ns2_1set, symshen_4internal_1to_1P_2, tmp7057)


_ = tmp7075

tmp7076 := MakeNative2(func(__e *ControlFlow, V3229, V3230 Obj) Obj {
tmp7095 := Call2(__e, SymFn(symelement_2), V3229, V3230)


if True == tmp7095 {
return V3229
} else {
tmp7093 := GConsP(__e, V3229)

var ifres7090 Obj

if True == tmp7093 {
tmp7091 := GHd(__e, V3229)

tmp7092 := GEq(__e, symcond, tmp7091)

ifres7090 = tmp7092


} else {
ifres7090 = False


}

if True == ifres7090 {
tmp7077 := GTl(__e, V3229)

tmp7078 := Call2(__e, SymFn(symshen_4process_1cond_1clauses), tmp7077, V3230)


return GCons(__e, symcond, tmp7078)


} else {
tmp7088 := GConsP(__e, V3229)

var ifres7085 Obj

if True == tmp7088 {
tmp7086 := GHd(__e, V3229)

tmp7087 := Call1(__e, SymFn(symshen_4non_1application_2), tmp7086)


ifres7085 = tmp7087


} else {
ifres7085 = False


}

if True == ifres7085 {
tmp7079 := GHd(__e, V3229)

return TailCall3(__e, SymFn(symshen_4special_1case), tmp7079, V3229, V3230)


} else {
tmp7083 := GConsP(__e, V3229)

if True == tmp7083 {
tmp7080 := MakeNative1(func(__e *ControlFlow, Z3231 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4process_1applications), Z3231, V3230)
panic("unreachable")
})

tmp7081 := Call2(__e, SymFn(symmap), tmp7080, V3229)


return TailCall2(__e, SymFn(symshen_4process_1application), tmp7081, V3230)


} else {
return V3229
}


}


}


}


panic("unreachable")
})

tmp7096 := Call2(__e, ns2_1set, symshen_4process_1applications, tmp7076)


_ = tmp7096

tmp7097 := MakeNative1(func(__e *ControlFlow, V3234 Obj) Obj {
tmp7107 := GEq(__e, symdefine, V3234)

if True == tmp7107 {
return True
} else {
tmp7105 := GEq(__e, symdefun, V3234)

if True == tmp7105 {
return True
} else {
tmp7103 := GEq(__e, symsynonyms, V3234)

if True == tmp7103 {
return True
} else {
tmp7101 := Call1(__e, SymFn(symshen_4special_2), V3234)


if True == tmp7101 {
return True
} else {
tmp7099 := Call1(__e, SymFn(symshen_4extraspecial_2), V3234)


if True == tmp7099 {
return True
} else {
return False
}


}


}


}


}


panic("unreachable")
})

tmp7108 := Call2(__e, ns2_1set, symshen_4non_1application_2, tmp7097)


_ = tmp7108

tmp7109 := MakeNative3(func(__e *ControlFlow, V3239, V3240, V3241 Obj) Obj {
tmp7314 := GEq(__e, symlambda, V3239)

var ifres7297 Obj

if True == tmp7314 {
tmp7313 := GConsP(__e, V3240)

var ifres7298 Obj

if True == tmp7313 {
tmp7311 := GHd(__e, V3240)

tmp7312 := GEq(__e, symlambda, tmp7311)

var ifres7299 Obj

if True == tmp7312 {
tmp7309 := GTl(__e, V3240)

tmp7310 := GConsP(__e, tmp7309)

var ifres7300 Obj

if True == tmp7310 {
tmp7306 := GTl(__e, V3240)

tmp7307 := GTl(__e, tmp7306)

tmp7308 := GConsP(__e, tmp7307)

var ifres7301 Obj

if True == tmp7308 {
tmp7302 := GTl(__e, V3240)

tmp7303 := GTl(__e, tmp7302)

tmp7304 := GTl(__e, tmp7303)

tmp7305 := GEq(__e, Nil, tmp7304)

ifres7301 = tmp7305


} else {
ifres7301 = False


}

ifres7300 = ifres7301


} else {
ifres7300 = False


}

ifres7299 = ifres7300


} else {
ifres7299 = False


}

ifres7298 = ifres7299


} else {
ifres7298 = False


}

ifres7297 = ifres7298


} else {
ifres7297 = False


}

if True == ifres7297 {
tmp7110 := GTl(__e, V3240)

tmp7111 := GHd(__e, tmp7110)

tmp7112 := GTl(__e, V3240)

tmp7113 := GTl(__e, tmp7112)

tmp7114 := GHd(__e, tmp7113)

tmp7115 := Call2(__e, SymFn(symshen_4process_1applications), tmp7114, V3241)


tmp7116 := GCons(__e, tmp7115, Nil)

tmp7117 := GCons(__e, tmp7111, tmp7116)

return GCons(__e, symlambda, tmp7117)


} else {
tmp7295 := GEq(__e, symlet, V3239)

var ifres7272 Obj

if True == tmp7295 {
tmp7294 := GConsP(__e, V3240)

var ifres7273 Obj

if True == tmp7294 {
tmp7292 := GHd(__e, V3240)

tmp7293 := GEq(__e, symlet, tmp7292)

var ifres7274 Obj

if True == tmp7293 {
tmp7290 := GTl(__e, V3240)

tmp7291 := GConsP(__e, tmp7290)

var ifres7275 Obj

if True == tmp7291 {
tmp7287 := GTl(__e, V3240)

tmp7288 := GTl(__e, tmp7287)

tmp7289 := GConsP(__e, tmp7288)

var ifres7276 Obj

if True == tmp7289 {
tmp7283 := GTl(__e, V3240)

tmp7284 := GTl(__e, tmp7283)

tmp7285 := GTl(__e, tmp7284)

tmp7286 := GConsP(__e, tmp7285)

var ifres7277 Obj

if True == tmp7286 {
tmp7278 := GTl(__e, V3240)

tmp7279 := GTl(__e, tmp7278)

tmp7280 := GTl(__e, tmp7279)

tmp7281 := GTl(__e, tmp7280)

tmp7282 := GEq(__e, Nil, tmp7281)

ifres7277 = tmp7282


} else {
ifres7277 = False


}

ifres7276 = ifres7277


} else {
ifres7276 = False


}

ifres7275 = ifres7276


} else {
ifres7275 = False


}

ifres7274 = ifres7275


} else {
ifres7274 = False


}

ifres7273 = ifres7274


} else {
ifres7273 = False


}

ifres7272 = ifres7273


} else {
ifres7272 = False


}

if True == ifres7272 {
tmp7118 := GTl(__e, V3240)

tmp7119 := GHd(__e, tmp7118)

tmp7120 := GTl(__e, V3240)

tmp7121 := GTl(__e, tmp7120)

tmp7122 := GHd(__e, tmp7121)

tmp7123 := Call2(__e, SymFn(symshen_4process_1applications), tmp7122, V3241)


tmp7124 := GTl(__e, V3240)

tmp7125 := GTl(__e, tmp7124)

tmp7126 := GTl(__e, tmp7125)

tmp7127 := GHd(__e, tmp7126)

tmp7128 := Call2(__e, SymFn(symshen_4process_1applications), tmp7127, V3241)


tmp7129 := GCons(__e, tmp7128, Nil)

tmp7130 := GCons(__e, tmp7123, tmp7129)

tmp7131 := GCons(__e, tmp7119, tmp7130)

return GCons(__e, symlet, tmp7131)


} else {
tmp7270 := GEq(__e, symdefun, V3239)

var ifres7247 Obj

if True == tmp7270 {
tmp7269 := GConsP(__e, V3240)

var ifres7248 Obj

if True == tmp7269 {
tmp7267 := GHd(__e, V3240)

tmp7268 := GEq(__e, symdefun, tmp7267)

var ifres7249 Obj

if True == tmp7268 {
tmp7265 := GTl(__e, V3240)

tmp7266 := GConsP(__e, tmp7265)

var ifres7250 Obj

if True == tmp7266 {
tmp7262 := GTl(__e, V3240)

tmp7263 := GTl(__e, tmp7262)

tmp7264 := GConsP(__e, tmp7263)

var ifres7251 Obj

if True == tmp7264 {
tmp7258 := GTl(__e, V3240)

tmp7259 := GTl(__e, tmp7258)

tmp7260 := GTl(__e, tmp7259)

tmp7261 := GConsP(__e, tmp7260)

var ifres7252 Obj

if True == tmp7261 {
tmp7253 := GTl(__e, V3240)

tmp7254 := GTl(__e, tmp7253)

tmp7255 := GTl(__e, tmp7254)

tmp7256 := GTl(__e, tmp7255)

tmp7257 := GEq(__e, Nil, tmp7256)

ifres7252 = tmp7257


} else {
ifres7252 = False


}

ifres7251 = ifres7252


} else {
ifres7251 = False


}

ifres7250 = ifres7251


} else {
ifres7250 = False


}

ifres7249 = ifres7250


} else {
ifres7249 = False


}

ifres7248 = ifres7249


} else {
ifres7248 = False


}

ifres7247 = ifres7248


} else {
ifres7247 = False


}

if True == ifres7247 {
return V3240
} else {
tmp7245 := GEq(__e, symdefine, V3239)

var ifres7228 Obj

if True == tmp7245 {
tmp7244 := GConsP(__e, V3240)

var ifres7229 Obj

if True == tmp7244 {
tmp7242 := GHd(__e, V3240)

tmp7243 := GEq(__e, symdefine, tmp7242)

var ifres7230 Obj

if True == tmp7243 {
tmp7240 := GTl(__e, V3240)

tmp7241 := GConsP(__e, tmp7240)

var ifres7231 Obj

if True == tmp7241 {
tmp7237 := GTl(__e, V3240)

tmp7238 := GTl(__e, tmp7237)

tmp7239 := GConsP(__e, tmp7238)

var ifres7232 Obj

if True == tmp7239 {
tmp7233 := GTl(__e, V3240)

tmp7234 := GTl(__e, tmp7233)

tmp7235 := GHd(__e, tmp7234)

tmp7236 := GEq(__e, sym_i, tmp7235)

ifres7232 = tmp7236


} else {
ifres7232 = False


}

ifres7231 = ifres7232


} else {
ifres7231 = False


}

ifres7230 = ifres7231


} else {
ifres7230 = False


}

ifres7229 = ifres7230


} else {
ifres7229 = False


}

ifres7228 = ifres7229


} else {
ifres7228 = False


}

if True == ifres7228 {
tmp7132 := GTl(__e, V3240)

tmp7133 := GHd(__e, tmp7132)

tmp7134 := GTl(__e, V3240)

tmp7135 := GHd(__e, tmp7134)

tmp7136 := GTl(__e, V3240)

tmp7137 := GTl(__e, tmp7136)

tmp7138 := GTl(__e, tmp7137)

tmp7139 := Call3(__e, SymFn(symshen_4process_1after_1type), tmp7135, tmp7138, V3241)


tmp7140 := GCons(__e, sym_i, tmp7139)

tmp7141 := GCons(__e, tmp7133, tmp7140)

return GCons(__e, symdefine, tmp7141)


} else {
tmp7226 := GEq(__e, symdefine, V3239)

var ifres7218 Obj

if True == tmp7226 {
tmp7225 := GConsP(__e, V3240)

var ifres7219 Obj

if True == tmp7225 {
tmp7223 := GHd(__e, V3240)

tmp7224 := GEq(__e, symdefine, tmp7223)

var ifres7220 Obj

if True == tmp7224 {
tmp7221 := GTl(__e, V3240)

tmp7222 := GConsP(__e, tmp7221)

ifres7220 = tmp7222


} else {
ifres7220 = False


}

ifres7219 = ifres7220


} else {
ifres7219 = False


}

ifres7218 = ifres7219


} else {
ifres7218 = False


}

if True == ifres7218 {
tmp7142 := GTl(__e, V3240)

tmp7143 := GHd(__e, tmp7142)

tmp7144 := MakeNative1(func(__e *ControlFlow, Z3242 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4process_1applications), Z3242, V3241)
panic("unreachable")
})

tmp7145 := GTl(__e, V3240)

tmp7146 := GTl(__e, tmp7145)

tmp7147 := Call2(__e, SymFn(symmap), tmp7144, tmp7146)


tmp7148 := GCons(__e, tmp7143, tmp7147)

return GCons(__e, symdefine, tmp7148)


} else {
tmp7216 := GEq(__e, symsynonyms, V3239)

if True == tmp7216 {
return GCons(__e, symsynonyms, V3240)
} else {
tmp7214 := GEq(__e, symtype, V3239)

var ifres7197 Obj

if True == tmp7214 {
tmp7213 := GConsP(__e, V3240)

var ifres7198 Obj

if True == tmp7213 {
tmp7211 := GHd(__e, V3240)

tmp7212 := GEq(__e, symtype, tmp7211)

var ifres7199 Obj

if True == tmp7212 {
tmp7209 := GTl(__e, V3240)

tmp7210 := GConsP(__e, tmp7209)

var ifres7200 Obj

if True == tmp7210 {
tmp7206 := GTl(__e, V3240)

tmp7207 := GTl(__e, tmp7206)

tmp7208 := GConsP(__e, tmp7207)

var ifres7201 Obj

if True == tmp7208 {
tmp7202 := GTl(__e, V3240)

tmp7203 := GTl(__e, tmp7202)

tmp7204 := GTl(__e, tmp7203)

tmp7205 := GEq(__e, Nil, tmp7204)

ifres7201 = tmp7205


} else {
ifres7201 = False


}

ifres7200 = ifres7201


} else {
ifres7200 = False


}

ifres7199 = ifres7200


} else {
ifres7199 = False


}

ifres7198 = ifres7199


} else {
ifres7198 = False


}

ifres7197 = ifres7198


} else {
ifres7197 = False


}

if True == ifres7197 {
tmp7149 := GTl(__e, V3240)

tmp7150 := GHd(__e, tmp7149)

tmp7151 := Call2(__e, SymFn(symshen_4process_1applications), tmp7150, V3241)


tmp7152 := GTl(__e, V3240)

tmp7153 := GTl(__e, tmp7152)

tmp7154 := GCons(__e, tmp7151, tmp7153)

return GCons(__e, symtype, tmp7154)


} else {
tmp7195 := GEq(__e, syminput_7, V3239)

var ifres7178 Obj

if True == tmp7195 {
tmp7194 := GConsP(__e, V3240)

var ifres7179 Obj

if True == tmp7194 {
tmp7192 := GHd(__e, V3240)

tmp7193 := GEq(__e, syminput_7, tmp7192)

var ifres7180 Obj

if True == tmp7193 {
tmp7190 := GTl(__e, V3240)

tmp7191 := GConsP(__e, tmp7190)

var ifres7181 Obj

if True == tmp7191 {
tmp7187 := GTl(__e, V3240)

tmp7188 := GTl(__e, tmp7187)

tmp7189 := GConsP(__e, tmp7188)

var ifres7182 Obj

if True == tmp7189 {
tmp7183 := GTl(__e, V3240)

tmp7184 := GTl(__e, tmp7183)

tmp7185 := GTl(__e, tmp7184)

tmp7186 := GEq(__e, Nil, tmp7185)

ifres7182 = tmp7186


} else {
ifres7182 = False


}

ifres7181 = ifres7182


} else {
ifres7181 = False


}

ifres7180 = ifres7181


} else {
ifres7180 = False


}

ifres7179 = ifres7180


} else {
ifres7179 = False


}

ifres7178 = ifres7179


} else {
ifres7178 = False


}

if True == ifres7178 {
tmp7155 := GTl(__e, V3240)

tmp7156 := GHd(__e, tmp7155)

tmp7157 := GTl(__e, V3240)

tmp7158 := GTl(__e, tmp7157)

tmp7159 := GHd(__e, tmp7158)

tmp7160 := Call2(__e, SymFn(symshen_4process_1applications), tmp7159, V3241)


tmp7161 := GCons(__e, tmp7160, Nil)

tmp7162 := GCons(__e, tmp7156, tmp7161)

return GCons(__e, syminput_7, tmp7162)


} else {
tmp7176 := GConsP(__e, V3240)

var ifres7173 Obj

if True == tmp7176 {
tmp7174 := GHd(__e, V3240)

tmp7175 := Call1(__e, SymFn(symshen_4special_2), tmp7174)


ifres7173 = tmp7175


} else {
ifres7173 = False


}

if True == ifres7173 {
tmp7163 := GHd(__e, V3240)

tmp7164 := MakeNative1(func(__e *ControlFlow, Z3243 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4process_1applications), Z3243, V3241)
panic("unreachable")
})

tmp7165 := GTl(__e, V3240)

tmp7166 := Call2(__e, SymFn(symmap), tmp7164, tmp7165)


return GCons(__e, tmp7163, tmp7166)


} else {
tmp7171 := GConsP(__e, V3240)

var ifres7168 Obj

if True == tmp7171 {
tmp7169 := GHd(__e, V3240)

tmp7170 := Call1(__e, SymFn(symshen_4extraspecial_2), tmp7169)


ifres7168 = tmp7170


} else {
ifres7168 = False


}

if True == ifres7168 {
return V3240
} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4special_1case)
}


}


}


}


}


}


}


}


}


}


panic("unreachable")
})

tmp7315 := Call2(__e, ns2_1set, symshen_4special_1case, tmp7109)


_ = tmp7315

tmp7316 := MakeNative2(func(__e *ControlFlow, V3246, V3247 Obj) Obj {
tmp7343 := GEq(__e, Nil, V3246)

if True == tmp7343 {
return Nil
} else {
tmp7341 := GConsP(__e, V3246)

var ifres7329 Obj

if True == tmp7341 {
tmp7339 := GHd(__e, V3246)

tmp7340 := GConsP(__e, tmp7339)

var ifres7330 Obj

if True == tmp7340 {
tmp7336 := GHd(__e, V3246)

tmp7337 := GTl(__e, tmp7336)

tmp7338 := GConsP(__e, tmp7337)

var ifres7331 Obj

if True == tmp7338 {
tmp7332 := GHd(__e, V3246)

tmp7333 := GTl(__e, tmp7332)

tmp7334 := GTl(__e, tmp7333)

tmp7335 := GEq(__e, Nil, tmp7334)

ifres7331 = tmp7335


} else {
ifres7331 = False


}

ifres7330 = ifres7331


} else {
ifres7330 = False


}

ifres7329 = ifres7330


} else {
ifres7329 = False


}

if True == ifres7329 {
tmp7317 := GHd(__e, V3246)

tmp7318 := GHd(__e, tmp7317)

tmp7319 := Call2(__e, SymFn(symshen_4process_1applications), tmp7318, V3247)


tmp7320 := GHd(__e, V3246)

tmp7321 := GTl(__e, tmp7320)

tmp7322 := GHd(__e, tmp7321)

tmp7323 := Call2(__e, SymFn(symshen_4process_1applications), tmp7322, V3247)


tmp7324 := GCons(__e, tmp7323, Nil)

tmp7325 := GCons(__e, tmp7319, tmp7324)

tmp7326 := GTl(__e, V3246)

tmp7327 := Call2(__e, SymFn(symshen_4process_1cond_1clauses), tmp7326, V3247)


return GCons(__e, tmp7325, tmp7327)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4process_1cond_1clauses)
}


}


panic("unreachable")
})

tmp7344 := Call2(__e, ns2_1set, symshen_4process_1cond_1clauses, tmp7316)


_ = tmp7344

tmp7345 := MakeNative3(func(__e *ControlFlow, V3250, V3251, V3252 Obj) Obj {
tmp7360 := GConsP(__e, V3251)

var ifres7357 Obj

if True == tmp7360 {
tmp7358 := GHd(__e, V3251)

tmp7359 := GEq(__e, sym_j, tmp7358)

ifres7357 = tmp7359


} else {
ifres7357 = False


}

if True == ifres7357 {
tmp7346 := MakeNative1(func(__e *ControlFlow, Z3253 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4process_1applications), Z3253, V3252)
panic("unreachable")
})

tmp7347 := GTl(__e, V3251)

tmp7348 := Call2(__e, SymFn(symmap), tmp7346, tmp7347)


return GCons(__e, sym_j, tmp7348)


} else {
tmp7355 := GConsP(__e, V3251)

if True == tmp7355 {
tmp7349 := GHd(__e, V3251)

tmp7350 := GTl(__e, V3251)

tmp7351 := Call3(__e, SymFn(symshen_4process_1after_1type), V3250, tmp7350, V3252)


return GCons(__e, tmp7349, tmp7351)


} else {
tmp7352 := Call3(__e, SymFn(symshen_4app), V3250, MakeString("\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("missing } in "), tmp7352))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("missing } in "), tmp7352)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


}


panic("unreachable")
})

tmp7361 := Call2(__e, ns2_1set, symshen_4process_1after_1type, tmp7345)


_ = tmp7361

tmp7362 := MakeNative2(func(__e *ControlFlow, V3254, V3255 Obj) Obj {
tmp7405 := GConsP(__e, V3254)

if True == tmp7405 {
tmp7363 := GHd(__e, V3254)

tmp7364 := Call1(__e, SymFn(symarity), tmp7363)


let__5162 := tmp7364
_ = let__5162

tmp7365 := GTl(__e, V3254)

tmp7366 := Call1(__e, SymFn(symlength), tmp7365)


let__5163 := tmp7366
_ = let__5163

tmp7403 := Call2(__e, SymFn(symelement_2), V3254, V3255)


if True == tmp7403 {
return V3254
} else {
tmp7400 := GHd(__e, V3254)

tmp7401 := Call1(__e, SymFn(symshen_4shen_1call_2), tmp7400)


if True == tmp7401 {
return V3254
} else {
tmp7398 := Call1(__e, SymFn(symshen_4foreign_2), V3254)


if True == tmp7398 {
return TailCall1(__e, SymFn(symshen_4unpack_1foreign), V3254)
} else {
tmp7396 := Call1(__e, SymFn(symshen_4fn_1call_2), V3254)


if True == tmp7396 {
return TailCall1(__e, SymFn(symshen_4fn_1call), V3254)
} else {
tmp7394 := Call1(__e, SymFn(symshen_4zero_1place_2), V3254)


if True == tmp7394 {
return V3254
} else {
tmp7391 := GHd(__e, V3254)

tmp7392 := Call2(__e, SymFn(symshen_4undefined_1f_2), tmp7391, let__5162)


if True == tmp7392 {
tmp7367 := GHd(__e, V3254)

tmp7368 := GCons(__e, tmp7367, Nil)

tmp7369 := GCons(__e, symfn, tmp7368)

tmp7370 := GTl(__e, V3254)

tmp7371 := GCons(__e, tmp7369, tmp7370)

return TailCall1(__e, SymFn(symshen_4simple_1curry), tmp7371)


} else {
tmp7388 := GHd(__e, V3254)

tmp7389 := PrimGuard1(__e, symvariable_2, PrimIsVariable, tmp7388)

if True == tmp7389 {
return TailCall1(__e, SymFn(symshen_4simple_1curry), V3254)
} else {
tmp7385 := GHd(__e, V3254)

tmp7386 := Call1(__e, SymFn(symshen_4application_2), tmp7385)


if True == tmp7386 {
return TailCall1(__e, SymFn(symshen_4simple_1curry), V3254)
} else {
tmp7382 := GHd(__e, V3254)

tmp7383 := Call3(__e, SymFn(symshen_4partial_1application_d_2), tmp7382, let__5162, let__5163)


if True == tmp7383 {
return TailCall2(__e, SymFn(symshen_4lambda_1function), V3254, GSub(__e, let__5162, let__5163))


} else {
tmp7379 := GHd(__e, V3254)

tmp7380 := Call3(__e, SymFn(symshen_4overapplication_2), tmp7379, let__5162, let__5163)


if True == tmp7380 {
tmp7373 := GHd(__e, V3254)

tmp7374 := GCons(__e, tmp7373, Nil)

tmp7375 := GCons(__e, symfn, tmp7374)

tmp7376 := GTl(__e, V3254)

tmp7377 := GCons(__e, tmp7375, tmp7376)

return TailCall1(__e, SymFn(symshen_4simple_1curry), tmp7377)


} else {
return V3254
}


}


}


}


}


}


}


}


}


}


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4process_1application)
}


panic("unreachable")
})

tmp7406 := Call2(__e, ns2_1set, symshen_4process_1application, tmp7362)


_ = tmp7406

tmp7407 := MakeNative1(func(__e *ControlFlow, V3258 Obj) Obj {
tmp7429 := GConsP(__e, V3258)

var ifres7413 Obj

if True == tmp7429 {
tmp7427 := GHd(__e, V3258)

tmp7428 := GConsP(__e, tmp7427)

var ifres7414 Obj

if True == tmp7428 {
tmp7424 := GHd(__e, V3258)

tmp7425 := GHd(__e, tmp7424)

tmp7426 := GEq(__e, symforeign, tmp7425)

var ifres7415 Obj

if True == tmp7426 {
tmp7421 := GHd(__e, V3258)

tmp7422 := GTl(__e, tmp7421)

tmp7423 := GConsP(__e, tmp7422)

var ifres7416 Obj

if True == tmp7423 {
tmp7417 := GHd(__e, V3258)

tmp7418 := GTl(__e, tmp7417)

tmp7419 := GTl(__e, tmp7418)

tmp7420 := GEq(__e, Nil, tmp7419)

ifres7416 = tmp7420


} else {
ifres7416 = False


}

ifres7415 = ifres7416


} else {
ifres7415 = False


}

ifres7414 = ifres7415


} else {
ifres7414 = False


}

ifres7413 = ifres7414


} else {
ifres7413 = False


}

if True == ifres7413 {
tmp7408 := GHd(__e, V3258)

tmp7409 := GTl(__e, tmp7408)

tmp7410 := GHd(__e, tmp7409)

tmp7411 := GTl(__e, V3258)

return GCons(__e, tmp7410, tmp7411)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4unpack_1foreign)
}


panic("unreachable")
})

tmp7430 := Call2(__e, ns2_1set, symshen_4unpack_1foreign, tmp7407)


_ = tmp7430

tmp7431 := MakeNative1(func(__e *ControlFlow, V3261 Obj) Obj {
tmp7449 := GConsP(__e, V3261)

var ifres7433 Obj

if True == tmp7449 {
tmp7447 := GHd(__e, V3261)

tmp7448 := GConsP(__e, tmp7447)

var ifres7434 Obj

if True == tmp7448 {
tmp7444 := GHd(__e, V3261)

tmp7445 := GHd(__e, tmp7444)

tmp7446 := GEq(__e, symforeign, tmp7445)

var ifres7435 Obj

if True == tmp7446 {
tmp7441 := GHd(__e, V3261)

tmp7442 := GTl(__e, tmp7441)

tmp7443 := GConsP(__e, tmp7442)

var ifres7436 Obj

if True == tmp7443 {
tmp7437 := GHd(__e, V3261)

tmp7438 := GTl(__e, tmp7437)

tmp7439 := GTl(__e, tmp7438)

tmp7440 := GEq(__e, Nil, tmp7439)

ifres7436 = tmp7440


} else {
ifres7436 = False


}

ifres7435 = ifres7436


} else {
ifres7435 = False


}

ifres7434 = ifres7435


} else {
ifres7434 = False


}

ifres7433 = ifres7434


} else {
ifres7433 = False


}

if True == ifres7433 {
return True
} else {
return False
}


panic("unreachable")
})

tmp7450 := Call2(__e, ns2_1set, symshen_4foreign_2, tmp7431)


_ = tmp7450

tmp7451 := MakeNative1(func(__e *ControlFlow, V3264 Obj) Obj {
tmp7456 := GConsP(__e, V3264)

var ifres7453 Obj

if True == tmp7456 {
tmp7454 := GTl(__e, V3264)

tmp7455 := GEq(__e, Nil, tmp7454)

ifres7453 = tmp7455


} else {
ifres7453 = False


}

if True == ifres7453 {
return True
} else {
return False
}


panic("unreachable")
})

tmp7457 := Call2(__e, ns2_1set, symshen_4zero_1place_2, tmp7451)


_ = tmp7457

tmp7458 := MakeNative1(func(__e *ControlFlow, V3265 Obj) Obj {
tmp7461 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V3265)

if True == tmp7461 {
tmp7459 := PrimGuard1(__e, symstr, PrimStr, V3265)

return TailCall1(__e, SymFn(symshen_4internal_1to_1shen_2), tmp7459)


} else {
return False
}


panic("unreachable")
})

tmp7462 := Call2(__e, ns2_1set, symshen_4shen_1call_2, tmp7458)


_ = tmp7462

tmp7463 := MakeNative1(func(__e *ControlFlow, V3270 Obj) Obj {
tmp7487 := GConsP(__e, V3270)

var ifres7477 Obj

if True == tmp7487 {
tmp7485 := GHd(__e, V3270)

tmp7486 := GEq(__e, symprotect, tmp7485)

var ifres7478 Obj

if True == tmp7486 {
tmp7483 := GTl(__e, V3270)

tmp7484 := GConsP(__e, tmp7483)

var ifres7479 Obj

if True == tmp7484 {
tmp7480 := GTl(__e, V3270)

tmp7481 := GTl(__e, tmp7480)

tmp7482 := GEq(__e, Nil, tmp7481)

ifres7479 = tmp7482


} else {
ifres7479 = False


}

ifres7478 = ifres7479


} else {
ifres7478 = False


}

ifres7477 = ifres7478


} else {
ifres7477 = False


}

if True == ifres7477 {
return False
} else {
tmp7475 := GConsP(__e, V3270)

var ifres7465 Obj

if True == tmp7475 {
tmp7473 := GHd(__e, V3270)

tmp7474 := GEq(__e, symforeign, tmp7473)

var ifres7466 Obj

if True == tmp7474 {
tmp7471 := GTl(__e, V3270)

tmp7472 := GConsP(__e, tmp7471)

var ifres7467 Obj

if True == tmp7472 {
tmp7468 := GTl(__e, V3270)

tmp7469 := GTl(__e, tmp7468)

tmp7470 := GEq(__e, Nil, tmp7469)

ifres7467 = tmp7470


} else {
ifres7467 = False


}

ifres7466 = ifres7467


} else {
ifres7466 = False


}

ifres7465 = ifres7466


} else {
ifres7465 = False


}

if True == ifres7465 {
return False
} else {
return GConsP(__e, V3270)
}


}


panic("unreachable")
})

tmp7488 := Call2(__e, ns2_1set, symshen_4application_2, tmp7463)


_ = tmp7488

tmp7489 := MakeNative2(func(__e *ControlFlow, V3275, V3276 Obj) Obj {
tmp7495 := GEq(__e, MakeInteger(-1), V3276)

if True == tmp7495 {
tmp7493 := Call1(__e, SymFn(symshen_4lowercase_1symbol_2), V3275)


if True == tmp7493 {
tmp7490 := Call1(__e, SymFn(symexternal), symshen)


tmp7491 := Call2(__e, SymFn(symelement_2), V3275, tmp7490)


return GNot(__e, tmp7491)


} else {
return False
}


} else {
return False
}


panic("unreachable")
})

tmp7496 := Call2(__e, ns2_1set, symshen_4undefined_1f_2, tmp7489)


_ = tmp7496

tmp7497 := MakeNative1(func(__e *ControlFlow, V3277 Obj) Obj {
tmp7500 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V3277)

if True == tmp7500 {
tmp7498 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V3277)

return GNot(__e, tmp7498)


} else {
return False
}


panic("unreachable")
})

tmp7501 := Call2(__e, ns2_1set, symshen_4lowercase_1symbol_2, tmp7497)


_ = tmp7501

tmp7502 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V3278 := __self1
_ = V3278
tmp7528 := GConsP(__e, V3278)

var ifres7521 Obj

if True == tmp7528 {
tmp7526 := GTl(__e, V3278)

tmp7527 := GConsP(__e, tmp7526)

var ifres7522 Obj

if True == tmp7527 {
tmp7523 := GTl(__e, V3278)

tmp7524 := GTl(__e, tmp7523)

tmp7525 := GEq(__e, Nil, tmp7524)

ifres7522 = tmp7525


} else {
ifres7522 = False


}

ifres7521 = ifres7522


} else {
ifres7521 = False


}

if True == ifres7521 {
return V3278
} else {
tmp7519 := GConsP(__e, V3278)

var ifres7512 Obj

if True == tmp7519 {
tmp7517 := GTl(__e, V3278)

tmp7518 := GConsP(__e, tmp7517)

var ifres7513 Obj

if True == tmp7518 {
tmp7514 := GTl(__e, V3278)

tmp7515 := GTl(__e, tmp7514)

tmp7516 := GConsP(__e, tmp7515)

ifres7513 = tmp7516


} else {
ifres7513 = False


}

ifres7512 = ifres7513


} else {
ifres7512 = False


}

if True == ifres7512 {
tmp7503 := GHd(__e, V3278)

tmp7504 := GTl(__e, V3278)

tmp7505 := GHd(__e, tmp7504)

tmp7506 := GCons(__e, tmp7505, Nil)

tmp7507 := GCons(__e, tmp7503, tmp7506)

tmp7508 := GTl(__e, V3278)

tmp7509 := GTl(__e, tmp7508)

tmp7510 := GCons(__e, tmp7507, tmp7509)

if SymFn(symshen_4simple_1curry) == __self {
__self1 = tmp7510
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4simple_1curry), tmp7510)


} else {
return V3278
}


}


panic("unreachable")
})
return __self
})()

tmp7529 := Call2(__e, ns2_1set, symshen_4simple_1curry, tmp7502)


_ = tmp7529

tmp7530 := MakeNative1(func(__e *ControlFlow, V3279 Obj) Obj {
return TailCall1(__e, SymFn(symfn), V3279)
panic("unreachable")
})

tmp7531 := Call2(__e, ns2_1set, symfunction, tmp7530)


_ = tmp7531

tmp7532 := MakeNative1(func(__e *ControlFlow, V3280 Obj) Obj {
tmp7539 := Call1(__e, SymFn(symarity), V3280)


tmp7540 := GEq(__e, tmp7539, MakeInteger(0))

if True == tmp7540 {
return TailCall0(__e, V3280)
} else {
tmp7533 := MakeNative0(func(__e *ControlFlow) Obj {
tmp7534 := GValue(__e, sym_dproperty_1vector_d)

return TailCall3(__e, SymFn(symget), V3280, symshen_4lambda_1form, tmp7534)


panic("unreachable")
})

tmp7535 := MakeNative1(func(__e *ControlFlow, Z3281 Obj) Obj {
tmp7536 := Call3(__e, SymFn(symshen_4app), V3280, MakeString(" is undefined\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("fn: "), tmp7536))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("fn: "), tmp7536)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp7533, tmp7535)


}


panic("unreachable")
})

tmp7541 := Call2(__e, ns2_1set, symfn, tmp7532)


_ = tmp7541

tmp7542 := MakeNative1(func(__e *ControlFlow, V3284 Obj) Obj {
tmp7566 := GConsP(__e, V3284)

var ifres7556 Obj

if True == tmp7566 {
tmp7564 := GHd(__e, V3284)

tmp7565 := GEq(__e, symfn, tmp7564)

var ifres7557 Obj

if True == tmp7565 {
tmp7562 := GTl(__e, V3284)

tmp7563 := GConsP(__e, tmp7562)

var ifres7558 Obj

if True == tmp7563 {
tmp7559 := GTl(__e, V3284)

tmp7560 := GTl(__e, tmp7559)

tmp7561 := GEq(__e, Nil, tmp7560)

ifres7558 = tmp7561


} else {
ifres7558 = False


}

ifres7557 = ifres7558


} else {
ifres7557 = False


}

ifres7556 = ifres7557


} else {
ifres7556 = False


}

if True == ifres7556 {
return True
} else {
tmp7554 := GConsP(__e, V3284)

var ifres7544 Obj

if True == tmp7554 {
tmp7552 := GHd(__e, V3284)

tmp7553 := GEq(__e, symfunction, tmp7552)

var ifres7545 Obj

if True == tmp7553 {
tmp7550 := GTl(__e, V3284)

tmp7551 := GConsP(__e, tmp7550)

var ifres7546 Obj

if True == tmp7551 {
tmp7547 := GTl(__e, V3284)

tmp7548 := GTl(__e, tmp7547)

tmp7549 := GEq(__e, Nil, tmp7548)

ifres7546 = tmp7549


} else {
ifres7546 = False


}

ifres7545 = ifres7546


} else {
ifres7545 = False


}

ifres7544 = ifres7545


} else {
ifres7544 = False


}

if True == ifres7544 {
return True
} else {
return False
}


}


panic("unreachable")
})

tmp7567 := Call2(__e, ns2_1set, symshen_4fn_1call_2, tmp7542)


_ = tmp7567

tmp7568 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V3285 := __self1
_ = V3285
tmp7602 := GConsP(__e, V3285)

var ifres7592 Obj

if True == tmp7602 {
tmp7600 := GHd(__e, V3285)

tmp7601 := GEq(__e, symfunction, tmp7600)

var ifres7593 Obj

if True == tmp7601 {
tmp7598 := GTl(__e, V3285)

tmp7599 := GConsP(__e, tmp7598)

var ifres7594 Obj

if True == tmp7599 {
tmp7595 := GTl(__e, V3285)

tmp7596 := GTl(__e, tmp7595)

tmp7597 := GEq(__e, Nil, tmp7596)

ifres7594 = tmp7597


} else {
ifres7594 = False


}

ifres7593 = ifres7594


} else {
ifres7593 = False


}

ifres7592 = ifres7593


} else {
ifres7592 = False


}

if True == ifres7592 {
tmp7569 := GTl(__e, V3285)

tmp7570 := GCons(__e, symfn, tmp7569)

if SymFn(symshen_4fn_1call) == __self {
__self1 = tmp7570
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4fn_1call), tmp7570)


} else {
tmp7590 := GConsP(__e, V3285)

var ifres7580 Obj

if True == tmp7590 {
tmp7588 := GHd(__e, V3285)

tmp7589 := GEq(__e, symfn, tmp7588)

var ifres7581 Obj

if True == tmp7589 {
tmp7586 := GTl(__e, V3285)

tmp7587 := GConsP(__e, tmp7586)

var ifres7582 Obj

if True == tmp7587 {
tmp7583 := GTl(__e, V3285)

tmp7584 := GTl(__e, tmp7583)

tmp7585 := GEq(__e, Nil, tmp7584)

ifres7582 = tmp7585


} else {
ifres7582 = False


}

ifres7581 = ifres7582


} else {
ifres7581 = False


}

ifres7580 = ifres7581


} else {
ifres7580 = False


}

if True == ifres7580 {
tmp7571 := GTl(__e, V3285)

tmp7572 := GHd(__e, tmp7571)

tmp7573 := Call1(__e, SymFn(symarity), tmp7572)


let__5164 := tmp7573
_ = let__5164

tmp7578 := GEq(__e, let__5164, MakeInteger(-1))

if True == tmp7578 {
return V3285
} else {
tmp7576 := GEq(__e, let__5164, MakeInteger(0))

if True == tmp7576 {
return GTl(__e, V3285)
} else {
tmp7574 := GTl(__e, V3285)

return TailCall2(__e, SymFn(symshen_4lambda_1function), tmp7574, let__5164)


}


}


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4fn_1call)
}


}


panic("unreachable")
})
return __self
})()

tmp7603 := Call2(__e, ns2_1set, symshen_4fn_1call, tmp7568)


_ = tmp7603

tmp7604 := MakeNative3(func(__e *ControlFlow, V3287, V3288, V3289 Obj) Obj {
let__5165 := GGt(__e, V3288, V3289)
_ = let__5165

var ifres7611 Obj

if True == let__5165 {
tmp7617 := Call0(__e, SymFn(symshen_4loading_2))


var ifres7612 Obj

if True == tmp7617 {
tmp7613 := GCons(__e, sym_1, Nil)

tmp7614 := GCons(__e, sym_7, tmp7613)

tmp7615 := Call2(__e, SymFn(symelement_2), V3287, tmp7614)


ifres7612 = GNot(__e, tmp7615)


} else {
ifres7612 = False


}

ifres7611 = ifres7612


} else {
ifres7611 = False


}

var ifres7606 Obj

if True == ifres7611 {
tmp7607 := Call3(__e, SymFn(symshen_4app), V3287, MakeString("\n"), symshen_4a)


tmp7608 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("partial application of "), tmp7607)

tmp7609 := Call0(__e, SymFn(symstoutput))


tmp7610 := Call2(__e, SymFn(sympr), tmp7608, tmp7609)


ifres7606 = tmp7610


} else {
ifres7606 = symshen_4skip


}

let__5166 := ifres7606
_ = let__5166

return let__5165


panic("unreachable")
})

tmp7618 := Call2(__e, ns2_1set, symshen_4partial_1application_d_2, tmp7604)


_ = tmp7618

tmp7619 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dloading_2_d)
panic("unreachable")
})

tmp7620 := Call2(__e, ns2_1set, symshen_4loading_2, tmp7619)


_ = tmp7620

tmp7621 := MakeNative3(func(__e *ControlFlow, V3296, V3297, V3298 Obj) Obj {
tmp7636 := GEq(__e, MakeInteger(-1), V3297)

if True == tmp7636 {
return False
} else {
let__5167 := GLt(__e, V3297, V3298)
_ = let__5167

var ifres7633 Obj

if True == let__5167 {
tmp7634 := Call0(__e, SymFn(symshen_4loading_2))


ifres7633 = tmp7634


} else {
ifres7633 = False


}

var ifres7623 Obj

if True == ifres7633 {
tmp7625 := GEq(__e, V3298, MakeInteger(1))

var ifres7624 Obj

if True == tmp7625 {
ifres7624 = MakeString("")


} else {
ifres7624 = MakeString("s")


}

tmp7626 := Call3(__e, SymFn(symshen_4app), ifres7624, MakeString("\n"), symshen_4a)


tmp7628 := Call3(__e, SymFn(symshen_4app), V3298, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" argument"), tmp7626), symshen_4a)


tmp7630 := Call3(__e, SymFn(symshen_4app), V3296, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" might not like "), tmp7628), symshen_4a)


tmp7631 := Call0(__e, SymFn(symstoutput))


tmp7632 := Call2(__e, SymFn(sympr), tmp7630, tmp7631)


ifres7623 = tmp7632


} else {
ifres7623 = symshen_4skip


}

let__5168 := ifres7623
_ = let__5168

return let__5167


}


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4overapplication_2, tmp7621)




panic("unreachable")
})

