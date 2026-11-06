package main

import . "github.com/pyrex41/shen-go/kl"

var ReaderMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp5149 := MakeNative1(func(__e *ControlFlow, V2556 Obj) Obj {
tmp5150 := PrimGuard1(__e, symread_1file_1as_1bytelist, PrimReadFileAsByteList, V2556)

let__4726 := tmp5150
_ = let__4726

tmp5151 := MakeNative0(func(__e *ControlFlow) Obj {
tmp5152 := MakeNative1(func(__e *ControlFlow, Z2559 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5s_1exprs_6), Z2559)
panic("unreachable")
})

return TailCall2(__e, SymFn(symcompile), tmp5152, let__4726)


panic("unreachable")
})

tmp5153 := MakeNative1(func(__e *ControlFlow, Z2560 Obj) Obj {
tmp5154 := GValue(__e, symshen_4_dresidue_d)

return TailCall1(__e, SymFn(symshen_4reader_1error), tmp5154)


panic("unreachable")
})

tmp5155 := Call2(__e, try_1catch, tmp5151, tmp5153)


let__4727 := tmp5155
_ = let__4727

tmp5156 := Call1(__e, SymFn(symshen_4process_1sexprs), let__4727)


let__4728 := tmp5156
_ = let__4728

return let__4728


panic("unreachable")
})

tmp5157 := Call2(__e, ns2_1set, symread_1file, tmp5149)


_ = tmp5157

tmp5158 := MakeNative1(func(__e *ControlFlow, V2562 Obj) Obj {
tmp5159 := GValue(__e, sym_dmaximum_1print_1sequence_1size_d)

tmp5160 := Call3(__e, SymFn(symshen_4reader_1error_1message), tmp5159, MakeInteger(0), V2562)


tmp5162 := Call1(__e, SymFn(symshen_4proc_1nl), PrimGuard2(__e, symcn, PrimStringConcat, MakeString("reader error near here: "), tmp5160))


return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp5162)


panic("unreachable")
})

tmp5163 := Call2(__e, ns2_1set, symshen_4reader_1error, tmp5158)


_ = tmp5163

tmp5164 := MakeNative3(func(__e *ControlFlow, V2570, V2571, V2572 Obj) Obj {
tmp5175 := GEq(__e, Nil, V2572)

if True == tmp5175 {
return MakeString("")
} else {
tmp5173 := GEq(__e, V2570, V2571)

if True == tmp5173 {
return MakeString("")
} else {
tmp5171 := GConsP(__e, V2572)

if True == tmp5171 {
tmp5165 := GHd(__e, V2572)

tmp5166 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, tmp5165)

tmp5167 := GAdd(__e, V2571, MakeInteger(1))

tmp5168 := GTl(__e, V2572)

tmp5169 := Call3(__e, SymFn(symshen_4reader_1error_1message), V2570, tmp5167, tmp5168)


return PrimGuard2(__e, symcn, PrimStringConcat, tmp5166, tmp5169)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4reader_1error_1message)
}


}


}


panic("unreachable")
})

tmp5176 := Call2(__e, ns2_1set, symshen_4reader_1error_1message, tmp5164)


_ = tmp5176

tmp5177 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dit_d)
panic("unreachable")
})

tmp5178 := Call2(__e, ns2_1set, symit, tmp5177)


_ = tmp5178

tmp5179 := MakeNative1(func(__e *ControlFlow, V2573 Obj) Obj {
tmp5180 := PrimGuard2(__e, symopen, PrimOpenStream, V2573, symin)

let__4729 := tmp5180
_ = let__4729

tmp5181 := PrimGuard1(__e, symread_1byte, PrimReadByte, let__4729)

let__4730 := tmp5181
_ = let__4730

tmp5182 := Call3(__e, SymFn(symshen_4read_1file_1as_1bytelist_1help), let__4729, let__4730, Nil)


let__4731 := tmp5182
_ = let__4731

tmp5183 := PrimGuard1(__e, symclose, PrimCloseStream, let__4729)

let__4732 := tmp5183
_ = let__4732

return TailCall1(__e, SymFn(symreverse), let__4731)


panic("unreachable")
})

tmp5184 := Call2(__e, ns2_1set, symread_1file_1as_1bytelist, tmp5179)


_ = tmp5184

tmp5185 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V2578 := __self1
_ = V2578
V2579 := __self2
_ = V2579
V2580 := __self3
_ = V2580
tmp5189 := GEq(__e, MakeInteger(-1), V2579)

if True == tmp5189 {
return V2580
} else {
tmp5186 := PrimGuard1(__e, symread_1byte, PrimReadByte, V2578)

tmp5187 := GCons(__e, V2579, V2580)

if SymFn(symshen_4read_1file_1as_1bytelist_1help) == __self {
__self1, __self2, __self3 = V2578, tmp5186, tmp5187
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4read_1file_1as_1bytelist_1help), V2578, tmp5186, tmp5187)


}


panic("unreachable")
})
return __self
})()

tmp5190 := Call2(__e, ns2_1set, symshen_4read_1file_1as_1bytelist_1help, tmp5185)


_ = tmp5190

tmp5191 := MakeNative1(func(__e *ControlFlow, V2581 Obj) Obj {
tmp5192 := PrimGuard2(__e, symopen, PrimOpenStream, V2581, symin)

let__4733 := tmp5192
_ = let__4733

tmp5193 := PrimGuard1(__e, symread_1byte, PrimReadByte, let__4733)

return TailCall3(__e, SymFn(symshen_4rfas_1h), let__4733, tmp5193, MakeString(""))


panic("unreachable")
})

tmp5194 := Call2(__e, ns2_1set, symread_1file_1as_1string, tmp5191)


_ = tmp5194

tmp5195 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V2583 := __self1
_ = V2583
V2584 := __self2
_ = V2584
V2585 := __self3
_ = V2585
tmp5201 := GEq(__e, MakeInteger(-1), V2584)

if True == tmp5201 {
tmp5196 := PrimGuard1(__e, symclose, PrimCloseStream, V2583)

_ = tmp5196

return V2585


} else {
tmp5197 := PrimGuard1(__e, symread_1byte, PrimReadByte, V2583)

tmp5198 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, V2584)

if SymFn(symshen_4rfas_1h) == __self {
__self1, __self2, __self3 = V2583, tmp5197, PrimGuard2(__e, symcn, PrimStringConcat, V2585, tmp5198)
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4rfas_1h), V2583, tmp5197, PrimGuard2(__e, symcn, PrimStringConcat, V2585, tmp5198))


}


panic("unreachable")
})
return __self
})()

tmp5202 := Call2(__e, ns2_1set, symshen_4rfas_1h, tmp5195)


_ = tmp5202

tmp5203 := MakeNative1(func(__e *ControlFlow, V2586 Obj) Obj {
tmp5204 := Call1(__e, SymFn(symread), V2586)


return TailCall1(__e, SymFn(symeval_1kl), tmp5204)


panic("unreachable")
})

tmp5205 := Call2(__e, ns2_1set, syminput, tmp5203)


_ = tmp5205

tmp5206 := MakeNative2(func(__e *ControlFlow, V2587, V2588 Obj) Obj {
tmp5207 := Call1(__e, SymFn(symshen_4monotype), V2587)


let__4734 := tmp5207
_ = let__4734

tmp5208 := Call1(__e, SymFn(symread), V2588)


let__4735 := tmp5208
_ = let__4735

tmp5214 := Call2(__e, SymFn(symshen_4typecheck), let__4735, V2587)


tmp5215 := GEq(__e, False, tmp5214)

if True == tmp5215 {
tmp5209 := Call3(__e, SymFn(symshen_4app), V2587, MakeString("\n"), symshen_4r)


tmp5211 := Call3(__e, SymFn(symshen_4app), let__4735, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" is not of type "), tmp5209), symshen_4r)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("type error: "), tmp5211))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("type error: "), tmp5211)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


} else {
return TailCall1(__e, SymFn(symeval_1kl), let__4735)
}


panic("unreachable")
})

tmp5216 := Call2(__e, ns2_1set, symshen_4input_1h_7, tmp5206)


_ = tmp5216

tmp5217 := MakeNative1(func(__e *ControlFlow, V2591 Obj) Obj {
tmp5224 := GConsP(__e, V2591)

if True == tmp5224 {
tmp5218 := MakeNative1(func(__e *ControlFlow, Z2592 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4monotype), Z2592)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp5218, V2591)


} else {
tmp5222 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V2591)

if True == tmp5222 {
tmp5219 := Call3(__e, SymFn(symshen_4app), V2591, MakeString("\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("input+ expects a monotype: not "), tmp5219))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("input+ expects a monotype: not "), tmp5219)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


} else {
return V2591
}


}


panic("unreachable")
})

tmp5225 := Call2(__e, ns2_1set, symshen_4monotype, tmp5217)


_ = tmp5225

tmp5226 := MakeNative1(func(__e *ControlFlow, V2593 Obj) Obj {
tmp5227 := Call1(__e, SymFn(symshen_4my_1read_1byte), V2593)


tmp5228 := MakeNative1(func(__e *ControlFlow, Z2594 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4return_2), Z2594)
panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4read_1loop), V2593, tmp5227, Nil, tmp5228)


panic("unreachable")
})

tmp5229 := Call2(__e, ns2_1set, symlineread, tmp5226)


_ = tmp5229

tmp5230 := MakeNative1(func(__e *ControlFlow, V2595 Obj) Obj {
tmp5231 := Call1(__e, SymFn(symshen_4str_1_6bytes), V2595)


let__4736 := tmp5231
_ = let__4736

tmp5232 := MakeNative1(func(__e *ControlFlow, Z2598 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5s_1exprs_6), Z2598)
panic("unreachable")
})

tmp5233 := Call2(__e, SymFn(symcompile), tmp5232, let__4736)


let__4737 := tmp5233
_ = let__4737

tmp5234 := Call1(__e, SymFn(symshen_4process_1sexprs), let__4737)


let__4738 := tmp5234
_ = let__4738

return let__4738


panic("unreachable")
})

tmp5235 := Call2(__e, ns2_1set, symread_1from_1string, tmp5230)


_ = tmp5235

tmp5236 := MakeNative1(func(__e *ControlFlow, V2600 Obj) Obj {
tmp5237 := Call1(__e, SymFn(symshen_4str_1_6bytes), V2600)


let__4739 := tmp5237
_ = let__4739

tmp5238 := MakeNative1(func(__e *ControlFlow, Z2603 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5s_1exprs_6), Z2603)
panic("unreachable")
})

tmp5239 := Call2(__e, SymFn(symcompile), tmp5238, let__4739)


let__4740 := tmp5239
_ = let__4740

return let__4740


panic("unreachable")
})

tmp5240 := Call2(__e, ns2_1set, symread_1from_1string_1unprocessed, tmp5236)


_ = tmp5240

tmp5241 := MakeNative1(func(__e *ControlFlow, V2604 Obj) Obj {
tmp5249 := GEq(__e, MakeString(""), V2604)

if True == tmp5249 {
return Nil
} else {
tmp5247 := Call1(__e, SymFn(symshen_4_7string_2), V2604)


if True == tmp5247 {
tmp5242 := Call1(__e, SymFn(symhdstr), V2604)


tmp5243 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp5242)

tmp5245 := Call1(__e, SymFn(symshen_4str_1_6bytes), PrimGuard1(__e, symtlstr, PrimTailString, V2604))


return GCons(__e, tmp5243, tmp5245)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4str_1_6bytes)
}


}


panic("unreachable")
})

tmp5250 := Call2(__e, ns2_1set, symshen_4str_1_6bytes, tmp5241)


_ = tmp5250

tmp5251 := MakeNative1(func(__e *ControlFlow, V2605 Obj) Obj {
tmp5252 := Call1(__e, SymFn(symshen_4my_1read_1byte), V2605)


tmp5253 := MakeNative1(func(__e *ControlFlow, Z2606 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4whitespace_2), Z2606)
panic("unreachable")
})

tmp5254 := Call4(__e, SymFn(symshen_4read_1loop), V2605, tmp5252, Nil, tmp5253)


return GHd(__e, tmp5254)


panic("unreachable")
})

tmp5255 := Call2(__e, ns2_1set, symread, tmp5251)


_ = tmp5255

tmp5256 := MakeNative1(func(__e *ControlFlow, V2607 Obj) Obj {
tmp5259 := Call1(__e, SymFn(symchar_1stream_2), V2607)


if True == tmp5259 {
tmp5257 := Call1(__e, SymFn(symread_1unit_1string), V2607)


return PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp5257)


} else {
return PrimGuard1(__e, symread_1byte, PrimReadByte, V2607)
}


panic("unreachable")
})

tmp5260 := Call2(__e, ns2_1set, symshen_4my_1read_1byte, tmp5256)


_ = tmp5260

tmp5261 := (func() Obj {
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
tmp5284 := GEq(__e, MakeInteger(94), V2613)

if True == tmp5284 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("read aborted"))
} else {
tmp5282 := GEq(__e, MakeInteger(-1), V2613)

if True == tmp5282 {
tmp5265 := Call1(__e, SymFn(symempty_2), V2614)


if True == tmp5265 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("error: empty stream"))
} else {
tmp5262 := MakeNative1(func(__e *ControlFlow, Z2616 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5s_1exprs_6), Z2616)
panic("unreachable")
})

tmp5263 := Call2(__e, SymFn(symcompile), tmp5262, V2614)


return TailCall1(__e, SymFn(symshen_4process_1sexprs), tmp5263)


}


} else {
tmp5280 := GEq(__e, MakeInteger(0), V2613)

if True == tmp5280 {
tmp5266 := Call1(__e, SymFn(symshen_4my_1read_1byte), V2612)


if SymFn(symshen_4read_1loop) == __self {
__self1, __self2, __self3, __self4 = V2612, tmp5266, V2614, V2615
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4read_1loop), V2612, tmp5266, V2614, V2615)


} else {
tmp5278 := Call1(__e, V2615, V2613)


if True == tmp5278 {
tmp5267 := Call1(__e, SymFn(symshen_4try_1parse), V2614)


let__4741 := tmp5267
_ = let__4741

tmp5273 := Call1(__e, SymFn(symshen_4nothing_1doing_2), let__4741)


if True == tmp5273 {
tmp5268 := Call1(__e, SymFn(symshen_4my_1read_1byte), V2612)


tmp5269 := GCons(__e, V2613, Nil)

tmp5270 := Call2(__e, SymFn(symappend), V2614, tmp5269)


if SymFn(symshen_4read_1loop) == __self {
__self1, __self2, __self3, __self4 = V2612, tmp5268, tmp5270, V2615
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4read_1loop), V2612, tmp5268, tmp5270, V2615)


} else {
tmp5271 := Call1(__e, SymFn(symshen_4record_1it), V2614)


_ = tmp5271

return TailCall1(__e, SymFn(symshen_4process_1sexprs), let__4741)


}


} else {
tmp5274 := Call1(__e, SymFn(symshen_4my_1read_1byte), V2612)


tmp5275 := GCons(__e, V2613, Nil)

tmp5276 := Call2(__e, SymFn(symappend), V2614, tmp5275)


if SymFn(symshen_4read_1loop) == __self {
__self1, __self2, __self3, __self4 = V2612, tmp5274, tmp5276, V2615
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4read_1loop), V2612, tmp5274, tmp5276, V2615)


}


}


}


}


panic("unreachable")
})
return __self
})()

tmp5285 := Call2(__e, ns2_1set, symshen_4read_1loop, tmp5261)


_ = tmp5285

tmp5286 := MakeNative1(func(__e *ControlFlow, V2618 Obj) Obj {
tmp5287 := MakeNative0(func(__e *ControlFlow) Obj {
tmp5288 := MakeNative1(func(__e *ControlFlow, Z2619 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5s_1exprs_6), Z2619)
panic("unreachable")
})

return TailCall2(__e, SymFn(symcompile), tmp5288, V2618)


panic("unreachable")
})

tmp5289 := MakeNative1(func(__e *ControlFlow, Z2620 Obj) Obj {
return symshen_4i_1failed_b
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp5287, tmp5289)


panic("unreachable")
})

tmp5290 := Call2(__e, ns2_1set, symshen_4try_1parse, tmp5286)


_ = tmp5290

tmp5291 := MakeNative1(func(__e *ControlFlow, V2623 Obj) Obj {
tmp5295 := GEq(__e, symshen_4i_1failed_b, V2623)

if True == tmp5295 {
return True
} else {
tmp5293 := GEq(__e, Nil, V2623)

if True == tmp5293 {
return True
} else {
return False
}


}


panic("unreachable")
})

tmp5296 := Call2(__e, ns2_1set, symshen_4nothing_1doing_2, tmp5291)


_ = tmp5296

tmp5297 := MakeNative1(func(__e *ControlFlow, V2624 Obj) Obj {
tmp5298 := Call1(__e, SymFn(symshen_4bytes_1_6string), V2624)


return PrimGuard2(__e, symset, PrimSet, symshen_4_dit_d, tmp5298)


panic("unreachable")
})

tmp5299 := Call2(__e, ns2_1set, symshen_4record_1it, tmp5297)


_ = tmp5299

tmp5300 := MakeNative1(func(__e *ControlFlow, V2625 Obj) Obj {
tmp5308 := GEq(__e, Nil, V2625)

if True == tmp5308 {
return MakeString("")
} else {
tmp5306 := GConsP(__e, V2625)

if True == tmp5306 {
tmp5301 := GHd(__e, V2625)

tmp5302 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, tmp5301)

tmp5303 := GTl(__e, V2625)

tmp5304 := Call1(__e, SymFn(symshen_4bytes_1_6string), tmp5303)


return PrimGuard2(__e, symcn, PrimStringConcat, tmp5302, tmp5304)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4bytes_1_6string)
}


}


panic("unreachable")
})

tmp5309 := Call2(__e, ns2_1set, symshen_4bytes_1_6string, tmp5300)


_ = tmp5309

tmp5310 := MakeNative1(func(__e *ControlFlow, V2626 Obj) Obj {
tmp5311 := Call1(__e, SymFn(symshen_4unpackage_emacroexpand), V2626)


let__4742 := tmp5311
_ = let__4742

tmp5312 := Call1(__e, SymFn(symshen_4find_1arities), let__4742)


let__4743 := tmp5312
_ = let__4743

tmp5313 := Call1(__e, SymFn(symshen_4find_1types), let__4742)


let__4744 := tmp5313
_ = let__4744

tmp5314 := MakeNative1(func(__e *ControlFlow, Z2630 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4process_1applications), Z2630, let__4744)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp5314, let__4742)


panic("unreachable")
})

tmp5315 := Call2(__e, ns2_1set, symshen_4process_1sexprs, tmp5310)


_ = tmp5315

tmp5316 := MakeNative1(func(__e *ControlFlow, V2631 Obj) Obj {
tmp5336 := GConsP(__e, V2631)

var ifres5329 Obj

if True == tmp5336 {
tmp5334 := GTl(__e, V2631)

tmp5335 := GConsP(__e, tmp5334)

var ifres5330 Obj

if True == tmp5335 {
tmp5331 := GHd(__e, V2631)

tmp5332 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp5333 := GEq(__e, tmp5331, tmp5332)

ifres5330 = tmp5333


} else {
ifres5330 = False


}

ifres5329 = ifres5330


} else {
ifres5329 = False


}

if True == ifres5329 {
tmp5317 := GTl(__e, V2631)

tmp5318 := GHd(__e, tmp5317)

tmp5319 := GTl(__e, V2631)

tmp5320 := GTl(__e, tmp5319)

tmp5321 := Call1(__e, SymFn(symshen_4find_1types), tmp5320)


return GCons(__e, tmp5318, tmp5321)


} else {
tmp5327 := GConsP(__e, V2631)

if True == tmp5327 {
tmp5322 := GHd(__e, V2631)

tmp5323 := Call1(__e, SymFn(symshen_4find_1types), tmp5322)


tmp5324 := GTl(__e, V2631)

tmp5325 := Call1(__e, SymFn(symshen_4find_1types), tmp5324)


return TailCall2(__e, SymFn(symappend), tmp5323, tmp5325)


} else {
return Nil
}


}


panic("unreachable")
})

tmp5337 := Call2(__e, ns2_1set, symshen_4find_1types, tmp5316)


_ = tmp5337

tmp5338 := MakeNative1(func(__e *ControlFlow, V2634 Obj) Obj {
tmp5381 := GConsP(__e, V2634)

var ifres5366 Obj

if True == tmp5381 {
tmp5379 := GHd(__e, V2634)

tmp5380 := GEq(__e, symdefine, tmp5379)

var ifres5367 Obj

if True == tmp5380 {
tmp5377 := GTl(__e, V2634)

tmp5378 := GConsP(__e, tmp5377)

var ifres5368 Obj

if True == tmp5378 {
tmp5374 := GTl(__e, V2634)

tmp5375 := GTl(__e, tmp5374)

tmp5376 := GConsP(__e, tmp5375)

var ifres5369 Obj

if True == tmp5376 {
tmp5370 := GTl(__e, V2634)

tmp5371 := GTl(__e, tmp5370)

tmp5372 := GHd(__e, tmp5371)

tmp5373 := GEq(__e, sym_i, tmp5372)

ifres5369 = tmp5373


} else {
ifres5369 = False


}

ifres5368 = ifres5369


} else {
ifres5368 = False


}

ifres5367 = ifres5368


} else {
ifres5367 = False


}

ifres5366 = ifres5367


} else {
ifres5366 = False


}

if True == ifres5366 {
tmp5339 := GTl(__e, V2634)

tmp5340 := GHd(__e, tmp5339)

tmp5341 := GTl(__e, V2634)

tmp5342 := GHd(__e, tmp5341)

tmp5343 := GTl(__e, V2634)

tmp5344 := GTl(__e, tmp5343)

tmp5345 := GTl(__e, tmp5344)

tmp5346 := Call3(__e, SymFn(symshen_4find_1arity), tmp5342, MakeInteger(1), tmp5345)


return TailCall2(__e, SymFn(symshen_4store_1arity), tmp5340, tmp5346)


} else {
tmp5364 := GConsP(__e, V2634)

var ifres5358 Obj

if True == tmp5364 {
tmp5362 := GHd(__e, V2634)

tmp5363 := GEq(__e, symdefine, tmp5362)

var ifres5359 Obj

if True == tmp5363 {
tmp5360 := GTl(__e, V2634)

tmp5361 := GConsP(__e, tmp5360)

ifres5359 = tmp5361


} else {
ifres5359 = False


}

ifres5358 = ifres5359


} else {
ifres5358 = False


}

if True == ifres5358 {
tmp5347 := GTl(__e, V2634)

tmp5348 := GHd(__e, tmp5347)

tmp5349 := GTl(__e, V2634)

tmp5350 := GHd(__e, tmp5349)

tmp5351 := GTl(__e, V2634)

tmp5352 := GTl(__e, tmp5351)

tmp5353 := Call3(__e, SymFn(symshen_4find_1arity), tmp5350, MakeInteger(0), tmp5352)


return TailCall2(__e, SymFn(symshen_4store_1arity), tmp5348, tmp5353)


} else {
tmp5356 := GConsP(__e, V2634)

if True == tmp5356 {
tmp5354 := MakeNative1(func(__e *ControlFlow, Z2635 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4find_1arities), Z2635)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp5354, V2634)


} else {
return symshen_4skip
}


}


}


panic("unreachable")
})

tmp5382 := Call2(__e, ns2_1set, symshen_4find_1arities, tmp5338)


_ = tmp5382

tmp5383 := MakeNative2(func(__e *ControlFlow, V2636, V2637 Obj) Obj {
tmp5384 := Call1(__e, SymFn(symarity), V2636)


let__4745 := tmp5384
_ = let__4745

tmp5395 := GEq(__e, let__4745, MakeInteger(-1))

if True == tmp5395 {
return TailCall2(__e, SymFn(symshen_4execute_1store_1arity), V2636, V2637)
} else {
tmp5393 := GEq(__e, let__4745, V2637)

if True == tmp5393 {
return symshen_4skip
} else {
tmp5391 := Call1(__e, SymFn(symshen_4sysfunc_2), V2636)


if True == tmp5391 {
tmp5385 := Call3(__e, SymFn(symshen_4app), V2636, MakeString(" is a system function\n"), symshen_4a)


return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp5385)


} else {
tmp5386 := Call3(__e, SymFn(symshen_4app), V2636, MakeString(" may cause errors\n"), symshen_4a)


tmp5387 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("changing the arity of "), tmp5386)

tmp5388 := Call0(__e, SymFn(symstoutput))


tmp5389 := Call2(__e, SymFn(sympr), tmp5387, tmp5388)


_ = tmp5389

return TailCall2(__e, SymFn(symshen_4execute_1store_1arity), V2636, V2637)


}


}


}


panic("unreachable")
})

tmp5396 := Call2(__e, ns2_1set, symshen_4store_1arity, tmp5383)


_ = tmp5396

tmp5397 := MakeNative2(func(__e *ControlFlow, V2639, V2640 Obj) Obj {
tmp5402 := GEq(__e, MakeInteger(0), V2640)

if True == tmp5402 {
tmp5398 := GValue(__e, sym_dproperty_1vector_d)

return TailCall4(__e, SymFn(symput), V2639, symarity, MakeInteger(0), tmp5398)


} else {
tmp5399 := GValue(__e, sym_dproperty_1vector_d)

tmp5400 := Call4(__e, SymFn(symput), V2639, symarity, V2640, tmp5399)


_ = tmp5400

return TailCall2(__e, SymFn(symshen_4update_1lambdatable), V2639, V2640)


}


panic("unreachable")
})

tmp5403 := Call2(__e, ns2_1set, symshen_4execute_1store_1arity, tmp5397)


_ = tmp5403

tmp5404 := MakeNative2(func(__e *ControlFlow, V2641, V2642 Obj) Obj {
tmp5405 := GCons(__e, V2641, Nil)

tmp5406 := Call2(__e, SymFn(symshen_4lambda_1function), tmp5405, V2642)


tmp5407 := Call1(__e, SymFn(symeval_1kl), tmp5406)


let__4746 := tmp5407
_ = let__4746

tmp5408 := GCons(__e, V2641, let__4746)

tmp5409 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp5408)


let__4747 := tmp5409
_ = let__4747

return V2642


panic("unreachable")
})

tmp5410 := Call2(__e, ns2_1set, symshen_4update_1lambdatable, tmp5404)


_ = tmp5410

tmp5411 := MakeNative2(func(__e *ControlFlow, V2647, V2648 Obj) Obj {
tmp5427 := GEq(__e, MakeInteger(0), V2648)

if True == tmp5427 {
return symshen_4skip
} else {
tmp5425 := GEq(__e, MakeInteger(1), V2648)

if True == tmp5425 {
tmp5412 := Call1(__e, SymFn(symgensym), symY)


let__4748 := tmp5412
_ = let__4748

tmp5413 := GCons(__e, let__4748, Nil)

tmp5414 := Call2(__e, SymFn(symappend), V2647, tmp5413)


tmp5415 := GCons(__e, tmp5414, Nil)

tmp5416 := GCons(__e, let__4748, tmp5415)

return GCons(__e, symlambda, tmp5416)


} else {
tmp5417 := Call1(__e, SymFn(symgensym), symY)


let__4749 := tmp5417
_ = let__4749

tmp5418 := GCons(__e, let__4749, Nil)

tmp5419 := Call2(__e, SymFn(symappend), V2647, tmp5418)


tmp5421 := Call2(__e, SymFn(symshen_4lambda_1function), tmp5419, GSub(__e, V2648, MakeInteger(1)))


tmp5422 := GCons(__e, tmp5421, Nil)

tmp5423 := GCons(__e, let__4749, tmp5422)

return GCons(__e, symlambda, tmp5423)


}


}


panic("unreachable")
})

tmp5428 := Call2(__e, ns2_1set, symshen_4lambda_1function, tmp5411)


_ = tmp5428

tmp5429 := MakeNative3(func(__e *ControlFlow, V2660, V2661, V2662 Obj) Obj {
tmp5450 := GEq(__e, Nil, V2662)

if True == tmp5450 {
tmp5430 := GCons(__e, V2660, V2661)

return GCons(__e, tmp5430, Nil)


} else {
tmp5448 := GConsP(__e, V2662)

var ifres5441 Obj

if True == tmp5448 {
tmp5446 := GHd(__e, V2662)

tmp5447 := GConsP(__e, tmp5446)

var ifres5442 Obj

if True == tmp5447 {
tmp5443 := GHd(__e, V2662)

tmp5444 := GHd(__e, tmp5443)

tmp5445 := GEq(__e, V2660, tmp5444)

ifres5442 = tmp5445


} else {
ifres5442 = False


}

ifres5441 = ifres5442


} else {
ifres5441 = False


}

if True == ifres5441 {
tmp5431 := GHd(__e, V2662)

tmp5432 := GHd(__e, tmp5431)

tmp5433 := GCons(__e, tmp5432, V2661)

tmp5434 := GTl(__e, V2662)

return GCons(__e, tmp5433, tmp5434)


} else {
tmp5439 := GConsP(__e, V2662)

if True == tmp5439 {
tmp5435 := GHd(__e, V2662)

tmp5436 := GTl(__e, V2662)

tmp5437 := Call3(__e, SymFn(symshen_4assoc_1_6), V2660, V2661, tmp5436)


return GCons(__e, tmp5435, tmp5437)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.assoc->"))
}


}


}


panic("unreachable")
})

tmp5451 := Call2(__e, ns2_1set, symshen_4assoc_1_6, tmp5429)


_ = tmp5451

tmp5452 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V2677 := __self1
_ = V2677
V2678 := __self2
_ = V2678
V2679 := __self3
_ = V2679
tmp5491 := GEq(__e, MakeInteger(0), V2678)

var ifres5486 Obj

if True == tmp5491 {
tmp5490 := GConsP(__e, V2679)

var ifres5487 Obj

if True == tmp5490 {
tmp5488 := GHd(__e, V2679)

tmp5489 := GEq(__e, tmp5488, sym_1_6)

ifres5487 = tmp5489


} else {
ifres5487 = False


}

ifres5486 = ifres5487


} else {
ifres5486 = False


}

if True == ifres5486 {
return MakeInteger(0)
} else {
tmp5484 := GEq(__e, MakeInteger(0), V2678)

var ifres5479 Obj

if True == tmp5484 {
tmp5483 := GConsP(__e, V2679)

var ifres5480 Obj

if True == tmp5483 {
tmp5481 := GHd(__e, V2679)

tmp5482 := GEq(__e, tmp5481, sym_5_1)

ifres5480 = tmp5482


} else {
ifres5480 = False


}

ifres5479 = ifres5480


} else {
ifres5479 = False


}

if True == ifres5479 {
return MakeInteger(0)
} else {
tmp5477 := GEq(__e, MakeInteger(0), V2678)

var ifres5475 Obj

if True == tmp5477 {
tmp5476 := GConsP(__e, V2679)

ifres5475 = tmp5476


} else {
ifres5475 = False


}

if True == ifres5475 {
tmp5453 := GTl(__e, V2679)

tmp5454 := Call3(__e, SymFn(symshen_4find_1arity), V2677, MakeInteger(0), tmp5453)


return GAdd(__e, MakeInteger(1), tmp5454)


} else {
tmp5473 := GEq(__e, MakeInteger(1), V2678)

var ifres5468 Obj

if True == tmp5473 {
tmp5472 := GConsP(__e, V2679)

var ifres5469 Obj

if True == tmp5472 {
tmp5470 := GHd(__e, V2679)

tmp5471 := GEq(__e, sym_j, tmp5470)

ifres5469 = tmp5471


} else {
ifres5469 = False


}

ifres5468 = ifres5469


} else {
ifres5468 = False


}

if True == ifres5468 {
tmp5455 := GTl(__e, V2679)

if SymFn(symshen_4find_1arity) == __self {
__self1, __self2, __self3 = V2677, MakeInteger(0), tmp5455
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4find_1arity), V2677, MakeInteger(0), tmp5455)


} else {
tmp5466 := GEq(__e, MakeInteger(1), V2678)

var ifres5464 Obj

if True == tmp5466 {
tmp5465 := GConsP(__e, V2679)

ifres5464 = tmp5465


} else {
ifres5464 = False


}

if True == ifres5464 {
tmp5456 := GTl(__e, V2679)

if SymFn(symshen_4find_1arity) == __self {
__self1, __self2, __self3 = V2677, MakeInteger(1), tmp5456
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4find_1arity), V2677, MakeInteger(1), tmp5456)


} else {
tmp5462 := GEq(__e, MakeInteger(1), V2678)

if True == tmp5462 {
tmp5457 := Call3(__e, SymFn(symshen_4app), V2677, MakeString(" definition: missing }\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("syntax error in "), tmp5457))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("syntax error in "), tmp5457)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


} else {
tmp5459 := Call3(__e, SymFn(symshen_4app), V2677, MakeString(" definition: missing -> or <-\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("syntax error in "), tmp5459))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("syntax error in "), tmp5459)
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

tmp5492 := Call2(__e, ns2_1set, symshen_4find_1arity, tmp5452)


_ = tmp5492

tmp5493 := MakeNative1(func(__e *ControlFlow, V2680 Obj) Obj {
tmp5494 := Call1(__e, SymFn(symshen_4_5lsb_6), V2680)


let__4751 := tmp5494
_ = let__4751

tmp5518 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4751)


var ifres5495 Obj

if True == tmp5518 {
tmp5496 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5495 = tmp5496


} else {
tmp5497 := Call1(__e, SymFn(symshen_4in_1_6), let__4751)


let__4752 := tmp5497
_ = let__4752

tmp5498 := Call1(__e, SymFn(symshen_4_5s_1exprs1_6), let__4752)


let__4753 := tmp5498
_ = let__4753

tmp5517 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4753)


var ifres5499 Obj

if True == tmp5517 {
tmp5500 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5499 = tmp5500


} else {
tmp5501 := Call1(__e, SymFn(symshen_4_5_1out), let__4753)


let__4754 := tmp5501
_ = let__4754

tmp5502 := Call1(__e, SymFn(symshen_4in_1_6), let__4753)


let__4755 := tmp5502
_ = let__4755

tmp5503 := Call1(__e, SymFn(symshen_4_5rsb_6), let__4755)


let__4756 := tmp5503
_ = let__4756

tmp5516 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4756)


var ifres5504 Obj

if True == tmp5516 {
tmp5505 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5504 = tmp5505


} else {
tmp5506 := Call1(__e, SymFn(symshen_4in_1_6), let__4756)


let__4757 := tmp5506
_ = let__4757

tmp5507 := Call1(__e, SymFn(symshen_4_5s_1exprs2_6), let__4757)


let__4758 := tmp5507
_ = let__4758

tmp5515 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4758)


var ifres5508 Obj

if True == tmp5515 {
tmp5509 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5508 = tmp5509


} else {
tmp5510 := Call1(__e, SymFn(symshen_4_5_1out), let__4758)


let__4759 := tmp5510
_ = let__4759

tmp5511 := Call1(__e, SymFn(symshen_4in_1_6), let__4758)


let__4760 := tmp5511
_ = let__4760

tmp5512 := Call1(__e, SymFn(symshen_4cons_1form), let__4754)


tmp5513 := GCons(__e, tmp5512, let__4759)

tmp5514 := Call2(__e, SymFn(symshen_4comb), let__4760, tmp5513)


ifres5508 = tmp5514


}

ifres5504 = ifres5508


}

ifres5499 = ifres5504


}

ifres5495 = ifres5499


}

let__4750 := ifres5495
_ = let__4750

tmp5712 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4750)


if True == tmp5712 {
tmp5519 := Call1(__e, SymFn(symshen_4_5lrb_6), V2680)


let__4762 := tmp5519
_ = let__4762

tmp5542 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4762)


var ifres5520 Obj

if True == tmp5542 {
tmp5521 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5520 = tmp5521


} else {
tmp5522 := Call1(__e, SymFn(symshen_4in_1_6), let__4762)


let__4763 := tmp5522
_ = let__4763

tmp5523 := Call1(__e, SymFn(symshen_4_5s_1exprs1_6), let__4763)


let__4764 := tmp5523
_ = let__4764

tmp5541 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4764)


var ifres5524 Obj

if True == tmp5541 {
tmp5525 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5524 = tmp5525


} else {
tmp5526 := Call1(__e, SymFn(symshen_4_5_1out), let__4764)


let__4765 := tmp5526
_ = let__4765

tmp5527 := Call1(__e, SymFn(symshen_4in_1_6), let__4764)


let__4766 := tmp5527
_ = let__4766

tmp5528 := Call1(__e, SymFn(symshen_4_5rrb_6), let__4766)


let__4767 := tmp5528
_ = let__4767

tmp5540 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4767)


var ifres5529 Obj

if True == tmp5540 {
tmp5530 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5529 = tmp5530


} else {
tmp5531 := Call1(__e, SymFn(symshen_4in_1_6), let__4767)


let__4768 := tmp5531
_ = let__4768

tmp5532 := Call1(__e, SymFn(symshen_4_5s_1exprs2_6), let__4768)


let__4769 := tmp5532
_ = let__4769

tmp5539 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4769)


var ifres5533 Obj

if True == tmp5539 {
tmp5534 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5533 = tmp5534


} else {
tmp5535 := Call1(__e, SymFn(symshen_4_5_1out), let__4769)


let__4770 := tmp5535
_ = let__4770

tmp5536 := Call1(__e, SymFn(symshen_4in_1_6), let__4769)


let__4771 := tmp5536
_ = let__4771

tmp5537 := Call2(__e, SymFn(symshen_4add_1sexpr), let__4765, let__4770)


tmp5538 := Call2(__e, SymFn(symshen_4comb), let__4771, tmp5537)


ifres5533 = tmp5538


}

ifres5529 = ifres5533


}

ifres5524 = ifres5529


}

ifres5520 = ifres5524


}

let__4761 := ifres5520
_ = let__4761

tmp5710 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4761)


if True == tmp5710 {
tmp5543 := Call1(__e, SymFn(symshen_4_5lcurly_6), V2680)


let__4773 := tmp5543
_ = let__4773

tmp5555 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4773)


var ifres5544 Obj

if True == tmp5555 {
tmp5545 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5544 = tmp5545


} else {
tmp5546 := Call1(__e, SymFn(symshen_4in_1_6), let__4773)


let__4774 := tmp5546
_ = let__4774

tmp5547 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__4774)


let__4775 := tmp5547
_ = let__4775

tmp5554 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4775)


var ifres5548 Obj

if True == tmp5554 {
tmp5549 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5548 = tmp5549


} else {
tmp5550 := Call1(__e, SymFn(symshen_4_5_1out), let__4775)


let__4776 := tmp5550
_ = let__4776

tmp5551 := Call1(__e, SymFn(symshen_4in_1_6), let__4775)


let__4777 := tmp5551
_ = let__4777

tmp5552 := GCons(__e, sym_i, let__4776)

tmp5553 := Call2(__e, SymFn(symshen_4comb), let__4777, tmp5552)


ifres5548 = tmp5553


}

ifres5544 = ifres5548


}

let__4772 := ifres5544
_ = let__4772

tmp5708 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4772)


if True == tmp5708 {
tmp5556 := Call1(__e, SymFn(symshen_4_5rcurly_6), V2680)


let__4779 := tmp5556
_ = let__4779

tmp5568 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4779)


var ifres5557 Obj

if True == tmp5568 {
tmp5558 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5557 = tmp5558


} else {
tmp5559 := Call1(__e, SymFn(symshen_4in_1_6), let__4779)


let__4780 := tmp5559
_ = let__4780

tmp5560 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__4780)


let__4781 := tmp5560
_ = let__4781

tmp5567 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4781)


var ifres5561 Obj

if True == tmp5567 {
tmp5562 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5561 = tmp5562


} else {
tmp5563 := Call1(__e, SymFn(symshen_4_5_1out), let__4781)


let__4782 := tmp5563
_ = let__4782

tmp5564 := Call1(__e, SymFn(symshen_4in_1_6), let__4781)


let__4783 := tmp5564
_ = let__4783

tmp5565 := GCons(__e, sym_j, let__4782)

tmp5566 := Call2(__e, SymFn(symshen_4comb), let__4783, tmp5565)


ifres5561 = tmp5566


}

ifres5557 = ifres5561


}

let__4778 := ifres5557
_ = let__4778

tmp5706 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4778)


if True == tmp5706 {
tmp5569 := Call1(__e, SymFn(symshen_4_5bar_6), V2680)


let__4785 := tmp5569
_ = let__4785

tmp5581 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4785)


var ifres5570 Obj

if True == tmp5581 {
tmp5571 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5570 = tmp5571


} else {
tmp5572 := Call1(__e, SymFn(symshen_4in_1_6), let__4785)


let__4786 := tmp5572
_ = let__4786

tmp5573 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__4786)


let__4787 := tmp5573
_ = let__4787

tmp5580 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4787)


var ifres5574 Obj

if True == tmp5580 {
tmp5575 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5574 = tmp5575


} else {
tmp5576 := Call1(__e, SymFn(symshen_4_5_1out), let__4787)


let__4788 := tmp5576
_ = let__4788

tmp5577 := Call1(__e, SymFn(symshen_4in_1_6), let__4787)


let__4789 := tmp5577
_ = let__4789

tmp5578 := GCons(__e, symbar_b, let__4788)

tmp5579 := Call2(__e, SymFn(symshen_4comb), let__4789, tmp5578)


ifres5574 = tmp5579


}

ifres5570 = ifres5574


}

let__4784 := ifres5570
_ = let__4784

tmp5704 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4784)


if True == tmp5704 {
tmp5582 := Call1(__e, SymFn(symshen_4_5semicolon_6), V2680)


let__4791 := tmp5582
_ = let__4791

tmp5595 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4791)


var ifres5583 Obj

if True == tmp5595 {
tmp5584 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5583 = tmp5584


} else {
tmp5585 := Call1(__e, SymFn(symshen_4in_1_6), let__4791)


let__4792 := tmp5585
_ = let__4792

tmp5586 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__4792)


let__4793 := tmp5586
_ = let__4793

tmp5594 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4793)


var ifres5587 Obj

if True == tmp5594 {
tmp5588 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5587 = tmp5588


} else {
tmp5589 := Call1(__e, SymFn(symshen_4_5_1out), let__4793)


let__4794 := tmp5589
_ = let__4794

tmp5590 := Call1(__e, SymFn(symshen_4in_1_6), let__4793)


let__4795 := tmp5590
_ = let__4795

tmp5591 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp5592 := GCons(__e, tmp5591, let__4794)

tmp5593 := Call2(__e, SymFn(symshen_4comb), let__4795, tmp5592)


ifres5587 = tmp5593


}

ifres5583 = ifres5587


}

let__4790 := ifres5583
_ = let__4790

tmp5702 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4790)


if True == tmp5702 {
tmp5596 := Call1(__e, SymFn(symshen_4_5colon_6), V2680)


let__4797 := tmp5596
_ = let__4797

tmp5614 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4797)


var ifres5597 Obj

if True == tmp5614 {
tmp5598 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5597 = tmp5598


} else {
tmp5599 := Call1(__e, SymFn(symshen_4in_1_6), let__4797)


let__4798 := tmp5599
_ = let__4798

tmp5600 := Call1(__e, SymFn(symshen_4_5equal_6), let__4798)


let__4799 := tmp5600
_ = let__4799

tmp5613 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4799)


var ifres5601 Obj

if True == tmp5613 {
tmp5602 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5601 = tmp5602


} else {
tmp5603 := Call1(__e, SymFn(symshen_4in_1_6), let__4799)


let__4800 := tmp5603
_ = let__4800

tmp5604 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__4800)


let__4801 := tmp5604
_ = let__4801

tmp5612 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4801)


var ifres5605 Obj

if True == tmp5612 {
tmp5606 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5605 = tmp5606


} else {
tmp5607 := Call1(__e, SymFn(symshen_4_5_1out), let__4801)


let__4802 := tmp5607
_ = let__4802

tmp5608 := Call1(__e, SymFn(symshen_4in_1_6), let__4801)


let__4803 := tmp5608
_ = let__4803

tmp5609 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":="))

tmp5610 := GCons(__e, tmp5609, let__4802)

tmp5611 := Call2(__e, SymFn(symshen_4comb), let__4803, tmp5610)


ifres5605 = tmp5611


}

ifres5601 = ifres5605


}

ifres5597 = ifres5601


}

let__4796 := ifres5597
_ = let__4796

tmp5700 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4796)


if True == tmp5700 {
tmp5615 := Call1(__e, SymFn(symshen_4_5colon_6), V2680)


let__4805 := tmp5615
_ = let__4805

tmp5628 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4805)


var ifres5616 Obj

if True == tmp5628 {
tmp5617 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5616 = tmp5617


} else {
tmp5618 := Call1(__e, SymFn(symshen_4in_1_6), let__4805)


let__4806 := tmp5618
_ = let__4806

tmp5619 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__4806)


let__4807 := tmp5619
_ = let__4807

tmp5627 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4807)


var ifres5620 Obj

if True == tmp5627 {
tmp5621 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5620 = tmp5621


} else {
tmp5622 := Call1(__e, SymFn(symshen_4_5_1out), let__4807)


let__4808 := tmp5622
_ = let__4808

tmp5623 := Call1(__e, SymFn(symshen_4in_1_6), let__4807)


let__4809 := tmp5623
_ = let__4809

tmp5624 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp5625 := GCons(__e, tmp5624, let__4808)

tmp5626 := Call2(__e, SymFn(symshen_4comb), let__4809, tmp5625)


ifres5620 = tmp5626


}

ifres5616 = ifres5620


}

let__4804 := ifres5616
_ = let__4804

tmp5698 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4804)


if True == tmp5698 {
tmp5629 := Call1(__e, SymFn(symshen_4_5comma_6), V2680)


let__4811 := tmp5629
_ = let__4811

tmp5642 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4811)


var ifres5630 Obj

if True == tmp5642 {
tmp5631 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5630 = tmp5631


} else {
tmp5632 := Call1(__e, SymFn(symshen_4in_1_6), let__4811)


let__4812 := tmp5632
_ = let__4812

tmp5633 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__4812)


let__4813 := tmp5633
_ = let__4813

tmp5641 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4813)


var ifres5634 Obj

if True == tmp5641 {
tmp5635 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5634 = tmp5635


} else {
tmp5636 := Call1(__e, SymFn(symshen_4_5_1out), let__4813)


let__4814 := tmp5636
_ = let__4814

tmp5637 := Call1(__e, SymFn(symshen_4in_1_6), let__4813)


let__4815 := tmp5637
_ = let__4815

tmp5638 := PrimGuard1(__e, symintern, PrimIntern, MakeString(","))

tmp5639 := GCons(__e, tmp5638, let__4814)

tmp5640 := Call2(__e, SymFn(symshen_4comb), let__4815, tmp5639)


ifres5634 = tmp5640


}

ifres5630 = ifres5634


}

let__4810 := ifres5630
_ = let__4810

tmp5696 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4810)


if True == tmp5696 {
tmp5643 := Call1(__e, SymFn(symshen_4_5comment_6), V2680)


let__4817 := tmp5643
_ = let__4817

tmp5654 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4817)


var ifres5644 Obj

if True == tmp5654 {
tmp5645 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5644 = tmp5645


} else {
tmp5646 := Call1(__e, SymFn(symshen_4in_1_6), let__4817)


let__4818 := tmp5646
_ = let__4818

tmp5647 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__4818)


let__4819 := tmp5647
_ = let__4819

tmp5653 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4819)


var ifres5648 Obj

if True == tmp5653 {
tmp5649 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5648 = tmp5649


} else {
tmp5650 := Call1(__e, SymFn(symshen_4_5_1out), let__4819)


let__4820 := tmp5650
_ = let__4820

tmp5651 := Call1(__e, SymFn(symshen_4in_1_6), let__4819)


let__4821 := tmp5651
_ = let__4821

tmp5652 := Call2(__e, SymFn(symshen_4comb), let__4821, let__4820)


ifres5648 = tmp5652


}

ifres5644 = ifres5648


}

let__4816 := ifres5644
_ = let__4816

tmp5694 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4816)


if True == tmp5694 {
tmp5655 := Call1(__e, SymFn(symshen_4_5atom_6), V2680)


let__4823 := tmp5655
_ = let__4823

tmp5668 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4823)


var ifres5656 Obj

if True == tmp5668 {
tmp5657 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5656 = tmp5657


} else {
tmp5658 := Call1(__e, SymFn(symshen_4_5_1out), let__4823)


let__4824 := tmp5658
_ = let__4824

tmp5659 := Call1(__e, SymFn(symshen_4in_1_6), let__4823)


let__4825 := tmp5659
_ = let__4825

tmp5660 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__4825)


let__4826 := tmp5660
_ = let__4826

tmp5667 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4826)


var ifres5661 Obj

if True == tmp5667 {
tmp5662 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5661 = tmp5662


} else {
tmp5663 := Call1(__e, SymFn(symshen_4_5_1out), let__4826)


let__4827 := tmp5663
_ = let__4827

tmp5664 := Call1(__e, SymFn(symshen_4in_1_6), let__4826)


let__4828 := tmp5664
_ = let__4828

tmp5665 := GCons(__e, let__4824, let__4827)

tmp5666 := Call2(__e, SymFn(symshen_4comb), let__4828, tmp5665)


ifres5661 = tmp5666


}

ifres5656 = ifres5661


}

let__4822 := ifres5656
_ = let__4822

tmp5692 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4822)


if True == tmp5692 {
tmp5669 := Call1(__e, SymFn(symshen_4_5whitespaces_6), V2680)


let__4830 := tmp5669
_ = let__4830

tmp5680 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4830)


var ifres5670 Obj

if True == tmp5680 {
tmp5671 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5670 = tmp5671


} else {
tmp5672 := Call1(__e, SymFn(symshen_4in_1_6), let__4830)


let__4831 := tmp5672
_ = let__4831

tmp5673 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__4831)


let__4832 := tmp5673
_ = let__4832

tmp5679 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4832)


var ifres5674 Obj

if True == tmp5679 {
tmp5675 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5674 = tmp5675


} else {
tmp5676 := Call1(__e, SymFn(symshen_4_5_1out), let__4832)


let__4833 := tmp5676
_ = let__4833

tmp5677 := Call1(__e, SymFn(symshen_4in_1_6), let__4832)


let__4834 := tmp5677
_ = let__4834

tmp5678 := Call2(__e, SymFn(symshen_4comb), let__4834, let__4833)


ifres5674 = tmp5678


}

ifres5670 = ifres5674


}

let__4829 := ifres5670
_ = let__4829

tmp5690 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4829)


if True == tmp5690 {
tmp5681 := Call1(__e, SymFn(sym_5e_6), V2680)


let__4836 := tmp5681
_ = let__4836

tmp5686 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4836)


var ifres5682 Obj

if True == tmp5686 {
tmp5683 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5682 = tmp5683


} else {
tmp5684 := Call1(__e, SymFn(symshen_4in_1_6), let__4836)


let__4837 := tmp5684
_ = let__4837

tmp5685 := Call2(__e, SymFn(symshen_4comb), let__4837, Nil)


ifres5682 = tmp5685


}

let__4835 := ifres5682
_ = let__4835

tmp5688 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4835)


if True == tmp5688 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4835
}


} else {
return let__4829
}


} else {
return let__4822
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
return let__4796
}


} else {
return let__4790
}


} else {
return let__4784
}


} else {
return let__4778
}


} else {
return let__4772
}


} else {
return let__4761
}


} else {
return let__4750
}


panic("unreachable")
})

tmp5713 := Call2(__e, ns2_1set, symshen_4_5s_1exprs_6, tmp5493)


_ = tmp5713

tmp5714 := MakeNative2(func(__e *ControlFlow, V2769, V2770 Obj) Obj {
tmp5729 := GConsP(__e, V2769)

var ifres5719 Obj

if True == tmp5729 {
tmp5727 := GHd(__e, V2769)

tmp5728 := GEq(__e, sym_3, tmp5727)

var ifres5720 Obj

if True == tmp5728 {
tmp5725 := GTl(__e, V2769)

tmp5726 := GConsP(__e, tmp5725)

var ifres5721 Obj

if True == tmp5726 {
tmp5722 := GTl(__e, V2769)

tmp5723 := GTl(__e, tmp5722)

tmp5724 := GEq(__e, Nil, tmp5723)

ifres5721 = tmp5724


} else {
ifres5721 = False


}

ifres5720 = ifres5721


} else {
ifres5720 = False


}

ifres5719 = ifres5720


} else {
ifres5719 = False


}

if True == ifres5719 {
tmp5715 := GTl(__e, V2769)

tmp5716 := GHd(__e, tmp5715)

tmp5717 := Call1(__e, SymFn(symexplode), tmp5716)


return TailCall2(__e, SymFn(symappend), tmp5717, V2770)


} else {
return GCons(__e, V2769, V2770)
}


panic("unreachable")
})

tmp5730 := Call2(__e, ns2_1set, symshen_4add_1sexpr, tmp5714)


_ = tmp5730

tmp5731 := MakeNative1(func(__e *ControlFlow, V2771 Obj) Obj {
tmp5736 := Call2(__e, SymFn(symshen_4hds_a_2), V2771, MakeInteger(91))


var ifres5732 Obj

if True == tmp5736 {
tmp5733 := Call1(__e, SymFn(symtail), V2771)


let__4839 := tmp5733
_ = let__4839

tmp5734 := Call2(__e, SymFn(symshen_4comb), let__4839, symshen_4skip)


ifres5732 = tmp5734


} else {
tmp5735 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5732 = tmp5735


}

let__4838 := ifres5732
_ = let__4838

tmp5738 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4838)


if True == tmp5738 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4838
}


panic("unreachable")
})

tmp5739 := Call2(__e, ns2_1set, symshen_4_5lsb_6, tmp5731)


_ = tmp5739

tmp5740 := MakeNative1(func(__e *ControlFlow, V2774 Obj) Obj {
tmp5745 := Call2(__e, SymFn(symshen_4hds_a_2), V2774, MakeInteger(93))


var ifres5741 Obj

if True == tmp5745 {
tmp5742 := Call1(__e, SymFn(symtail), V2774)


let__4841 := tmp5742
_ = let__4841

tmp5743 := Call2(__e, SymFn(symshen_4comb), let__4841, symshen_4skip)


ifres5741 = tmp5743


} else {
tmp5744 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5741 = tmp5744


}

let__4840 := ifres5741
_ = let__4840

tmp5747 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4840)


if True == tmp5747 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4840
}


panic("unreachable")
})

tmp5748 := Call2(__e, ns2_1set, symshen_4_5rsb_6, tmp5740)


_ = tmp5748

tmp5749 := MakeNative1(func(__e *ControlFlow, V2777 Obj) Obj {
tmp5750 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), V2777)


let__4843 := tmp5750
_ = let__4843

tmp5756 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4843)


var ifres5751 Obj

if True == tmp5756 {
tmp5752 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5751 = tmp5752


} else {
tmp5753 := Call1(__e, SymFn(symshen_4_5_1out), let__4843)


let__4844 := tmp5753
_ = let__4844

tmp5754 := Call1(__e, SymFn(symshen_4in_1_6), let__4843)


let__4845 := tmp5754
_ = let__4845

tmp5755 := Call2(__e, SymFn(symshen_4comb), let__4845, let__4844)


ifres5751 = tmp5755


}

let__4842 := ifres5751
_ = let__4842

tmp5758 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4842)


if True == tmp5758 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4842
}


panic("unreachable")
})

tmp5759 := Call2(__e, ns2_1set, symshen_4_5s_1exprs1_6, tmp5749)


_ = tmp5759

tmp5760 := MakeNative1(func(__e *ControlFlow, V2782 Obj) Obj {
tmp5761 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), V2782)


let__4847 := tmp5761
_ = let__4847

tmp5767 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4847)


var ifres5762 Obj

if True == tmp5767 {
tmp5763 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5762 = tmp5763


} else {
tmp5764 := Call1(__e, SymFn(symshen_4_5_1out), let__4847)


let__4848 := tmp5764
_ = let__4848

tmp5765 := Call1(__e, SymFn(symshen_4in_1_6), let__4847)


let__4849 := tmp5765
_ = let__4849

tmp5766 := Call2(__e, SymFn(symshen_4comb), let__4849, let__4848)


ifres5762 = tmp5766


}

let__4846 := ifres5762
_ = let__4846

tmp5769 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4846)


if True == tmp5769 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4846
}


panic("unreachable")
})

tmp5770 := Call2(__e, ns2_1set, symshen_4_5s_1exprs2_6, tmp5760)


_ = tmp5770

tmp5771 := MakeNative1(func(__e *ControlFlow, V2788 Obj) Obj {
tmp5820 := GEq(__e, Nil, V2788)

if True == tmp5820 {
return Nil
} else {
tmp5818 := GConsP(__e, V2788)

var ifres5802 Obj

if True == tmp5818 {
tmp5816 := GTl(__e, V2788)

tmp5817 := GConsP(__e, tmp5816)

var ifres5803 Obj

if True == tmp5817 {
tmp5813 := GTl(__e, V2788)

tmp5814 := GTl(__e, tmp5813)

tmp5815 := GConsP(__e, tmp5814)

var ifres5804 Obj

if True == tmp5815 {
tmp5809 := GTl(__e, V2788)

tmp5810 := GTl(__e, tmp5809)

tmp5811 := GTl(__e, tmp5810)

tmp5812 := GEq(__e, Nil, tmp5811)

var ifres5805 Obj

if True == tmp5812 {
tmp5806 := GTl(__e, V2788)

tmp5807 := GHd(__e, tmp5806)

tmp5808 := GEq(__e, tmp5807, symbar_b)

ifres5805 = tmp5808


} else {
ifres5805 = False


}

ifres5804 = ifres5805


} else {
ifres5804 = False


}

ifres5803 = ifres5804


} else {
ifres5803 = False


}

ifres5802 = ifres5803


} else {
ifres5802 = False


}

if True == ifres5802 {
tmp5772 := GHd(__e, V2788)

tmp5773 := GTl(__e, V2788)

tmp5774 := GTl(__e, tmp5773)

tmp5775 := GCons(__e, tmp5772, tmp5774)

return GCons(__e, symcons, tmp5775)


} else {
tmp5800 := GConsP(__e, V2788)

var ifres5784 Obj

if True == tmp5800 {
tmp5798 := GTl(__e, V2788)

tmp5799 := GConsP(__e, tmp5798)

var ifres5785 Obj

if True == tmp5799 {
tmp5795 := GTl(__e, V2788)

tmp5796 := GTl(__e, tmp5795)

tmp5797 := GConsP(__e, tmp5796)

var ifres5786 Obj

if True == tmp5797 {
tmp5791 := GTl(__e, V2788)

tmp5792 := GTl(__e, tmp5791)

tmp5793 := GTl(__e, tmp5792)

tmp5794 := GConsP(__e, tmp5793)

var ifres5787 Obj

if True == tmp5794 {
tmp5788 := GTl(__e, V2788)

tmp5789 := GHd(__e, tmp5788)

tmp5790 := GEq(__e, tmp5789, symbar_b)

ifres5787 = tmp5790


} else {
ifres5787 = False


}

ifres5786 = ifres5787


} else {
ifres5786 = False


}

ifres5785 = ifres5786


} else {
ifres5785 = False


}

ifres5784 = ifres5785


} else {
ifres5784 = False


}

if True == ifres5784 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("misapplication of |\n"))
} else {
tmp5782 := GConsP(__e, V2788)

if True == tmp5782 {
tmp5776 := GHd(__e, V2788)

tmp5777 := GTl(__e, V2788)

tmp5778 := Call1(__e, SymFn(symshen_4cons_1form), tmp5777)


tmp5779 := GCons(__e, tmp5778, Nil)

tmp5780 := GCons(__e, tmp5776, tmp5779)

return GCons(__e, symcons, tmp5780)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4cons_1form)
}


}


}


}


panic("unreachable")
})

tmp5821 := Call2(__e, ns2_1set, symshen_4cons_1form, tmp5771)


_ = tmp5821

tmp5822 := MakeNative1(func(__e *ControlFlow, V2789 Obj) Obj {
tmp5827 := Call2(__e, SymFn(symshen_4hds_a_2), V2789, MakeInteger(40))


var ifres5823 Obj

if True == tmp5827 {
tmp5824 := Call1(__e, SymFn(symtail), V2789)


let__4851 := tmp5824
_ = let__4851

tmp5825 := Call2(__e, SymFn(symshen_4comb), let__4851, symshen_4skip)


ifres5823 = tmp5825


} else {
tmp5826 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5823 = tmp5826


}

let__4850 := ifres5823
_ = let__4850

tmp5829 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4850)


if True == tmp5829 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4850
}


panic("unreachable")
})

tmp5830 := Call2(__e, ns2_1set, symshen_4_5lrb_6, tmp5822)


_ = tmp5830

tmp5831 := MakeNative1(func(__e *ControlFlow, V2792 Obj) Obj {
tmp5836 := Call2(__e, SymFn(symshen_4hds_a_2), V2792, MakeInteger(41))


var ifres5832 Obj

if True == tmp5836 {
tmp5833 := Call1(__e, SymFn(symtail), V2792)


let__4853 := tmp5833
_ = let__4853

tmp5834 := Call2(__e, SymFn(symshen_4comb), let__4853, symshen_4skip)


ifres5832 = tmp5834


} else {
tmp5835 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5832 = tmp5835


}

let__4852 := ifres5832
_ = let__4852

tmp5838 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4852)


if True == tmp5838 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4852
}


panic("unreachable")
})

tmp5839 := Call2(__e, ns2_1set, symshen_4_5rrb_6, tmp5831)


_ = tmp5839

tmp5840 := MakeNative1(func(__e *ControlFlow, V2795 Obj) Obj {
tmp5845 := Call2(__e, SymFn(symshen_4hds_a_2), V2795, MakeInteger(123))


var ifres5841 Obj

if True == tmp5845 {
tmp5842 := Call1(__e, SymFn(symtail), V2795)


let__4855 := tmp5842
_ = let__4855

tmp5843 := Call2(__e, SymFn(symshen_4comb), let__4855, symshen_4skip)


ifres5841 = tmp5843


} else {
tmp5844 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5841 = tmp5844


}

let__4854 := ifres5841
_ = let__4854

tmp5847 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4854)


if True == tmp5847 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4854
}


panic("unreachable")
})

tmp5848 := Call2(__e, ns2_1set, symshen_4_5lcurly_6, tmp5840)


_ = tmp5848

tmp5849 := MakeNative1(func(__e *ControlFlow, V2798 Obj) Obj {
tmp5854 := Call2(__e, SymFn(symshen_4hds_a_2), V2798, MakeInteger(125))


var ifres5850 Obj

if True == tmp5854 {
tmp5851 := Call1(__e, SymFn(symtail), V2798)


let__4857 := tmp5851
_ = let__4857

tmp5852 := Call2(__e, SymFn(symshen_4comb), let__4857, symshen_4skip)


ifres5850 = tmp5852


} else {
tmp5853 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5850 = tmp5853


}

let__4856 := ifres5850
_ = let__4856

tmp5856 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4856)


if True == tmp5856 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4856
}


panic("unreachable")
})

tmp5857 := Call2(__e, ns2_1set, symshen_4_5rcurly_6, tmp5849)


_ = tmp5857

tmp5858 := MakeNative1(func(__e *ControlFlow, V2801 Obj) Obj {
tmp5863 := Call2(__e, SymFn(symshen_4hds_a_2), V2801, MakeInteger(124))


var ifres5859 Obj

if True == tmp5863 {
tmp5860 := Call1(__e, SymFn(symtail), V2801)


let__4859 := tmp5860
_ = let__4859

tmp5861 := Call2(__e, SymFn(symshen_4comb), let__4859, symshen_4skip)


ifres5859 = tmp5861


} else {
tmp5862 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5859 = tmp5862


}

let__4858 := ifres5859
_ = let__4858

tmp5865 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4858)


if True == tmp5865 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4858
}


panic("unreachable")
})

tmp5866 := Call2(__e, ns2_1set, symshen_4_5bar_6, tmp5858)


_ = tmp5866

tmp5867 := MakeNative1(func(__e *ControlFlow, V2804 Obj) Obj {
tmp5872 := Call2(__e, SymFn(symshen_4hds_a_2), V2804, MakeInteger(59))


var ifres5868 Obj

if True == tmp5872 {
tmp5869 := Call1(__e, SymFn(symtail), V2804)


let__4861 := tmp5869
_ = let__4861

tmp5870 := Call2(__e, SymFn(symshen_4comb), let__4861, symshen_4skip)


ifres5868 = tmp5870


} else {
tmp5871 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5868 = tmp5871


}

let__4860 := ifres5868
_ = let__4860

tmp5874 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4860)


if True == tmp5874 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4860
}


panic("unreachable")
})

tmp5875 := Call2(__e, ns2_1set, symshen_4_5semicolon_6, tmp5867)


_ = tmp5875

tmp5876 := MakeNative1(func(__e *ControlFlow, V2807 Obj) Obj {
tmp5881 := Call2(__e, SymFn(symshen_4hds_a_2), V2807, MakeInteger(58))


var ifres5877 Obj

if True == tmp5881 {
tmp5878 := Call1(__e, SymFn(symtail), V2807)


let__4863 := tmp5878
_ = let__4863

tmp5879 := Call2(__e, SymFn(symshen_4comb), let__4863, symshen_4skip)


ifres5877 = tmp5879


} else {
tmp5880 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5877 = tmp5880


}

let__4862 := ifres5877
_ = let__4862

tmp5883 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4862)


if True == tmp5883 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4862
}


panic("unreachable")
})

tmp5884 := Call2(__e, ns2_1set, symshen_4_5colon_6, tmp5876)


_ = tmp5884

tmp5885 := MakeNative1(func(__e *ControlFlow, V2810 Obj) Obj {
tmp5890 := Call2(__e, SymFn(symshen_4hds_a_2), V2810, MakeInteger(44))


var ifres5886 Obj

if True == tmp5890 {
tmp5887 := Call1(__e, SymFn(symtail), V2810)


let__4865 := tmp5887
_ = let__4865

tmp5888 := Call2(__e, SymFn(symshen_4comb), let__4865, symshen_4skip)


ifres5886 = tmp5888


} else {
tmp5889 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5886 = tmp5889


}

let__4864 := ifres5886
_ = let__4864

tmp5892 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4864)


if True == tmp5892 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4864
}


panic("unreachable")
})

tmp5893 := Call2(__e, ns2_1set, symshen_4_5comma_6, tmp5885)


_ = tmp5893

tmp5894 := MakeNative1(func(__e *ControlFlow, V2813 Obj) Obj {
tmp5899 := Call2(__e, SymFn(symshen_4hds_a_2), V2813, MakeInteger(61))


var ifres5895 Obj

if True == tmp5899 {
tmp5896 := Call1(__e, SymFn(symtail), V2813)


let__4867 := tmp5896
_ = let__4867

tmp5897 := Call2(__e, SymFn(symshen_4comb), let__4867, symshen_4skip)


ifres5895 = tmp5897


} else {
tmp5898 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5895 = tmp5898


}

let__4866 := ifres5895
_ = let__4866

tmp5901 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4866)


if True == tmp5901 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4866
}


panic("unreachable")
})

tmp5902 := Call2(__e, ns2_1set, symshen_4_5equal_6, tmp5894)


_ = tmp5902

tmp5903 := MakeNative1(func(__e *ControlFlow, V2816 Obj) Obj {
tmp5904 := Call1(__e, SymFn(symshen_4_5singleline_6), V2816)


let__4869 := tmp5904
_ = let__4869

tmp5909 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4869)


var ifres5905 Obj

if True == tmp5909 {
tmp5906 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5905 = tmp5906


} else {
tmp5907 := Call1(__e, SymFn(symshen_4in_1_6), let__4869)


let__4870 := tmp5907
_ = let__4870

tmp5908 := Call2(__e, SymFn(symshen_4comb), let__4870, symshen_4skip)


ifres5905 = tmp5908


}

let__4868 := ifres5905
_ = let__4868

tmp5919 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4868)


if True == tmp5919 {
tmp5910 := Call1(__e, SymFn(symshen_4_5multiline_6), V2816)


let__4872 := tmp5910
_ = let__4872

tmp5915 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4872)


var ifres5911 Obj

if True == tmp5915 {
tmp5912 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5911 = tmp5912


} else {
tmp5913 := Call1(__e, SymFn(symshen_4in_1_6), let__4872)


let__4873 := tmp5913
_ = let__4873

tmp5914 := Call2(__e, SymFn(symshen_4comb), let__4873, symshen_4skip)


ifres5911 = tmp5914


}

let__4871 := ifres5911
_ = let__4871

tmp5917 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4871)


if True == tmp5917 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4871
}


} else {
return let__4868
}


panic("unreachable")
})

tmp5920 := Call2(__e, ns2_1set, symshen_4_5comment_6, tmp5903)


_ = tmp5920

tmp5921 := MakeNative1(func(__e *ControlFlow, V2823 Obj) Obj {
tmp5922 := Call1(__e, SymFn(symshen_4_5backslash_6), V2823)


let__4875 := tmp5922
_ = let__4875

tmp5942 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4875)


var ifres5923 Obj

if True == tmp5942 {
tmp5924 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5923 = tmp5924


} else {
tmp5925 := Call1(__e, SymFn(symshen_4in_1_6), let__4875)


let__4876 := tmp5925
_ = let__4876

tmp5926 := Call1(__e, SymFn(symshen_4_5backslash_6), let__4876)


let__4877 := tmp5926
_ = let__4877

tmp5941 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4877)


var ifres5927 Obj

if True == tmp5941 {
tmp5928 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5927 = tmp5928


} else {
tmp5929 := Call1(__e, SymFn(symshen_4in_1_6), let__4877)


let__4878 := tmp5929
_ = let__4878

tmp5930 := Call1(__e, SymFn(symshen_4_5shortnatters_6), let__4878)


let__4879 := tmp5930
_ = let__4879

tmp5940 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4879)


var ifres5931 Obj

if True == tmp5940 {
tmp5932 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5931 = tmp5932


} else {
tmp5933 := Call1(__e, SymFn(symshen_4in_1_6), let__4879)


let__4880 := tmp5933
_ = let__4880

tmp5934 := Call1(__e, SymFn(symshen_4_5returns_6), let__4880)


let__4881 := tmp5934
_ = let__4881

tmp5939 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4881)


var ifres5935 Obj

if True == tmp5939 {
tmp5936 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5935 = tmp5936


} else {
tmp5937 := Call1(__e, SymFn(symshen_4in_1_6), let__4881)


let__4882 := tmp5937
_ = let__4882

tmp5938 := Call2(__e, SymFn(symshen_4comb), let__4882, symshen_4skip)


ifres5935 = tmp5938


}

ifres5931 = ifres5935


}

ifres5927 = ifres5931


}

ifres5923 = ifres5927


}

let__4874 := ifres5923
_ = let__4874

tmp5944 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4874)


if True == tmp5944 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4874
}


panic("unreachable")
})

tmp5945 := Call2(__e, ns2_1set, symshen_4_5singleline_6, tmp5921)


_ = tmp5945

tmp5946 := MakeNative1(func(__e *ControlFlow, V2833 Obj) Obj {
tmp5951 := Call2(__e, SymFn(symshen_4hds_a_2), V2833, MakeInteger(92))


var ifres5947 Obj

if True == tmp5951 {
tmp5948 := Call1(__e, SymFn(symtail), V2833)


let__4884 := tmp5948
_ = let__4884

tmp5949 := Call2(__e, SymFn(symshen_4comb), let__4884, symshen_4skip)


ifres5947 = tmp5949


} else {
tmp5950 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5947 = tmp5950


}

let__4883 := ifres5947
_ = let__4883

tmp5953 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4883)


if True == tmp5953 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4883
}


panic("unreachable")
})

tmp5954 := Call2(__e, ns2_1set, symshen_4_5backslash_6, tmp5946)


_ = tmp5954

tmp5955 := MakeNative1(func(__e *ControlFlow, V2836 Obj) Obj {
tmp5956 := Call1(__e, SymFn(symshen_4_5shortnatter_6), V2836)


let__4886 := tmp5956
_ = let__4886

tmp5966 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4886)


var ifres5957 Obj

if True == tmp5966 {
tmp5958 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5957 = tmp5958


} else {
tmp5959 := Call1(__e, SymFn(symshen_4in_1_6), let__4886)


let__4887 := tmp5959
_ = let__4887

tmp5960 := Call1(__e, SymFn(symshen_4_5shortnatters_6), let__4887)


let__4888 := tmp5960
_ = let__4888

tmp5965 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4888)


var ifres5961 Obj

if True == tmp5965 {
tmp5962 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5961 = tmp5962


} else {
tmp5963 := Call1(__e, SymFn(symshen_4in_1_6), let__4888)


let__4889 := tmp5963
_ = let__4889

tmp5964 := Call2(__e, SymFn(symshen_4comb), let__4889, symshen_4skip)


ifres5961 = tmp5964


}

ifres5957 = ifres5961


}

let__4885 := ifres5957
_ = let__4885

tmp5976 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4885)


if True == tmp5976 {
tmp5967 := Call1(__e, SymFn(sym_5e_6), V2836)


let__4891 := tmp5967
_ = let__4891

tmp5972 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4891)


var ifres5968 Obj

if True == tmp5972 {
tmp5969 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5968 = tmp5969


} else {
tmp5970 := Call1(__e, SymFn(symshen_4in_1_6), let__4891)


let__4892 := tmp5970
_ = let__4892

tmp5971 := Call2(__e, SymFn(symshen_4comb), let__4892, symshen_4skip)


ifres5968 = tmp5971


}

let__4890 := ifres5968
_ = let__4890

tmp5974 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4890)


if True == tmp5974 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4890
}


} else {
return let__4885
}


panic("unreachable")
})

tmp5977 := Call2(__e, ns2_1set, symshen_4_5shortnatters_6, tmp5955)


_ = tmp5977

tmp5978 := MakeNative1(func(__e *ControlFlow, V2845 Obj) Obj {
tmp5988 := GConsP(__e, V2845)

var ifres5979 Obj

if True == tmp5988 {
tmp5980 := Call1(__e, SymFn(symhead), V2845)


let__4894 := tmp5980
_ = let__4894

tmp5981 := Call1(__e, SymFn(symtail), V2845)


let__4895 := tmp5981
_ = let__4895

tmp5985 := Call1(__e, SymFn(symshen_4return_2), let__4894)


tmp5986 := GNot(__e, tmp5985)

var ifres5982 Obj

if True == tmp5986 {
tmp5983 := Call2(__e, SymFn(symshen_4comb), let__4895, symshen_4skip)


ifres5982 = tmp5983


} else {
tmp5984 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5982 = tmp5984


}

ifres5979 = ifres5982


} else {
tmp5987 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5979 = tmp5987


}

let__4893 := ifres5979
_ = let__4893

tmp5990 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4893)


if True == tmp5990 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4893
}


panic("unreachable")
})

tmp5991 := Call2(__e, ns2_1set, symshen_4_5shortnatter_6, tmp5978)


_ = tmp5991

tmp5992 := MakeNative1(func(__e *ControlFlow, V2849 Obj) Obj {
tmp5993 := Call1(__e, SymFn(symshen_4_5return_6), V2849)


let__4897 := tmp5993
_ = let__4897

tmp6003 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4897)


var ifres5994 Obj

if True == tmp6003 {
tmp5995 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5994 = tmp5995


} else {
tmp5996 := Call1(__e, SymFn(symshen_4in_1_6), let__4897)


let__4898 := tmp5996
_ = let__4898

tmp5997 := Call1(__e, SymFn(symshen_4_5returns_6), let__4898)


let__4899 := tmp5997
_ = let__4899

tmp6002 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4899)


var ifres5998 Obj

if True == tmp6002 {
tmp5999 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5998 = tmp5999


} else {
tmp6000 := Call1(__e, SymFn(symshen_4in_1_6), let__4899)


let__4900 := tmp6000
_ = let__4900

tmp6001 := Call2(__e, SymFn(symshen_4comb), let__4900, symshen_4skip)


ifres5998 = tmp6001


}

ifres5994 = ifres5998


}

let__4896 := ifres5994
_ = let__4896

tmp6013 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4896)


if True == tmp6013 {
tmp6004 := Call1(__e, SymFn(symshen_4_5return_6), V2849)


let__4902 := tmp6004
_ = let__4902

tmp6009 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4902)


var ifres6005 Obj

if True == tmp6009 {
tmp6006 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6005 = tmp6006


} else {
tmp6007 := Call1(__e, SymFn(symshen_4in_1_6), let__4902)


let__4903 := tmp6007
_ = let__4903

tmp6008 := Call2(__e, SymFn(symshen_4comb), let__4903, symshen_4skip)


ifres6005 = tmp6008


}

let__4901 := ifres6005
_ = let__4901

tmp6011 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4901)


if True == tmp6011 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4901
}


} else {
return let__4896
}


panic("unreachable")
})

tmp6014 := Call2(__e, ns2_1set, symshen_4_5returns_6, tmp5992)


_ = tmp6014

tmp6015 := MakeNative1(func(__e *ControlFlow, V2858 Obj) Obj {
tmp6024 := GConsP(__e, V2858)

var ifres6016 Obj

if True == tmp6024 {
tmp6017 := Call1(__e, SymFn(symhead), V2858)


let__4905 := tmp6017
_ = let__4905

tmp6018 := Call1(__e, SymFn(symtail), V2858)


let__4906 := tmp6018
_ = let__4906

tmp6022 := Call1(__e, SymFn(symshen_4return_2), let__4905)


var ifres6019 Obj

if True == tmp6022 {
tmp6020 := Call2(__e, SymFn(symshen_4comb), let__4906, symshen_4skip)


ifres6019 = tmp6020


} else {
tmp6021 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6019 = tmp6021


}

ifres6016 = ifres6019


} else {
tmp6023 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6016 = tmp6023


}

let__4904 := ifres6016
_ = let__4904

tmp6026 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4904)


if True == tmp6026 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4904
}


panic("unreachable")
})

tmp6027 := Call2(__e, ns2_1set, symshen_4_5return_6, tmp6015)


_ = tmp6027

tmp6028 := MakeNative1(func(__e *ControlFlow, V2862 Obj) Obj {
tmp6029 := GCons(__e, MakeInteger(13), Nil)

tmp6030 := GCons(__e, MakeInteger(10), tmp6029)

tmp6031 := GCons(__e, MakeInteger(9), tmp6030)

return TailCall2(__e, SymFn(symelement_2), V2862, tmp6031)


panic("unreachable")
})

tmp6032 := Call2(__e, ns2_1set, symshen_4return_2, tmp6028)


_ = tmp6032

tmp6033 := MakeNative1(func(__e *ControlFlow, V2863 Obj) Obj {
tmp6034 := Call1(__e, SymFn(symshen_4_5backslash_6), V2863)


let__4908 := tmp6034
_ = let__4908

tmp6049 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4908)


var ifres6035 Obj

if True == tmp6049 {
tmp6036 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6035 = tmp6036


} else {
tmp6037 := Call1(__e, SymFn(symshen_4in_1_6), let__4908)


let__4909 := tmp6037
_ = let__4909

tmp6038 := Call1(__e, SymFn(symshen_4_5times_6), let__4909)


let__4910 := tmp6038
_ = let__4910

tmp6048 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4910)


var ifres6039 Obj

if True == tmp6048 {
tmp6040 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6039 = tmp6040


} else {
tmp6041 := Call1(__e, SymFn(symshen_4in_1_6), let__4910)


let__4911 := tmp6041
_ = let__4911

tmp6042 := Call1(__e, SymFn(symshen_4_5longnatter_6), let__4911)


let__4912 := tmp6042
_ = let__4912

tmp6047 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4912)


var ifres6043 Obj

if True == tmp6047 {
tmp6044 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6043 = tmp6044


} else {
tmp6045 := Call1(__e, SymFn(symshen_4in_1_6), let__4912)


let__4913 := tmp6045
_ = let__4913

tmp6046 := Call2(__e, SymFn(symshen_4comb), let__4913, symshen_4skip)


ifres6043 = tmp6046


}

ifres6039 = ifres6043


}

ifres6035 = ifres6039


}

let__4907 := ifres6035
_ = let__4907

tmp6051 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4907)


if True == tmp6051 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4907
}


panic("unreachable")
})

tmp6052 := Call2(__e, ns2_1set, symshen_4_5multiline_6, tmp6033)


_ = tmp6052

tmp6053 := MakeNative1(func(__e *ControlFlow, V2871 Obj) Obj {
tmp6058 := Call2(__e, SymFn(symshen_4hds_a_2), V2871, MakeInteger(42))


var ifres6054 Obj

if True == tmp6058 {
tmp6055 := Call1(__e, SymFn(symtail), V2871)


let__4915 := tmp6055
_ = let__4915

tmp6056 := Call2(__e, SymFn(symshen_4comb), let__4915, symshen_4skip)


ifres6054 = tmp6056


} else {
tmp6057 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6054 = tmp6057


}

let__4914 := ifres6054
_ = let__4914

tmp6060 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4914)


if True == tmp6060 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4914
}


panic("unreachable")
})

tmp6061 := Call2(__e, ns2_1set, symshen_4_5times_6, tmp6053)


_ = tmp6061

tmp6062 := MakeNative1(func(__e *ControlFlow, V2874 Obj) Obj {
tmp6063 := Call1(__e, SymFn(symshen_4_5comment_6), V2874)


let__4917 := tmp6063
_ = let__4917

tmp6073 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4917)


var ifres6064 Obj

if True == tmp6073 {
tmp6065 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6064 = tmp6065


} else {
tmp6066 := Call1(__e, SymFn(symshen_4in_1_6), let__4917)


let__4918 := tmp6066
_ = let__4918

tmp6067 := Call1(__e, SymFn(symshen_4_5longnatter_6), let__4918)


let__4919 := tmp6067
_ = let__4919

tmp6072 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4919)


var ifres6068 Obj

if True == tmp6072 {
tmp6069 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6068 = tmp6069


} else {
tmp6070 := Call1(__e, SymFn(symshen_4in_1_6), let__4919)


let__4920 := tmp6070
_ = let__4920

tmp6071 := Call2(__e, SymFn(symshen_4comb), let__4920, symshen_4skip)


ifres6068 = tmp6071


}

ifres6064 = ifres6068


}

let__4916 := ifres6064
_ = let__4916

tmp6100 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4916)


if True == tmp6100 {
tmp6074 := Call1(__e, SymFn(symshen_4_5times_6), V2874)


let__4922 := tmp6074
_ = let__4922

tmp6084 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4922)


var ifres6075 Obj

if True == tmp6084 {
tmp6076 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6075 = tmp6076


} else {
tmp6077 := Call1(__e, SymFn(symshen_4in_1_6), let__4922)


let__4923 := tmp6077
_ = let__4923

tmp6078 := Call1(__e, SymFn(symshen_4_5backslash_6), let__4923)


let__4924 := tmp6078
_ = let__4924

tmp6083 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4924)


var ifres6079 Obj

if True == tmp6083 {
tmp6080 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6079 = tmp6080


} else {
tmp6081 := Call1(__e, SymFn(symshen_4in_1_6), let__4924)


let__4925 := tmp6081
_ = let__4925

tmp6082 := Call2(__e, SymFn(symshen_4comb), let__4925, symshen_4skip)


ifres6079 = tmp6082


}

ifres6075 = ifres6079


}

let__4921 := ifres6075
_ = let__4921

tmp6098 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4921)


if True == tmp6098 {
tmp6094 := GConsP(__e, V2874)

var ifres6085 Obj

if True == tmp6094 {
tmp6086 := Call1(__e, SymFn(symtail), V2874)


let__4927 := tmp6086
_ = let__4927

tmp6087 := Call1(__e, SymFn(symshen_4_5longnatter_6), let__4927)


let__4928 := tmp6087
_ = let__4928

tmp6092 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4928)


var ifres6088 Obj

if True == tmp6092 {
tmp6089 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6088 = tmp6089


} else {
tmp6090 := Call1(__e, SymFn(symshen_4in_1_6), let__4928)


let__4929 := tmp6090
_ = let__4929

tmp6091 := Call2(__e, SymFn(symshen_4comb), let__4929, symshen_4skip)


ifres6088 = tmp6091


}

ifres6085 = ifres6088


} else {
tmp6093 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6085 = tmp6093


}

let__4926 := ifres6085
_ = let__4926

tmp6096 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4926)


if True == tmp6096 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4926
}


} else {
return let__4921
}


} else {
return let__4916
}


panic("unreachable")
})

tmp6101 := Call2(__e, ns2_1set, symshen_4_5longnatter_6, tmp6062)


_ = tmp6101

tmp6102 := MakeNative1(func(__e *ControlFlow, V2889 Obj) Obj {
tmp6103 := Call1(__e, SymFn(symshen_4_5str_6), V2889)


let__4931 := tmp6103
_ = let__4931

tmp6109 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4931)


var ifres6104 Obj

if True == tmp6109 {
tmp6105 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6104 = tmp6105


} else {
tmp6106 := Call1(__e, SymFn(symshen_4_5_1out), let__4931)


let__4932 := tmp6106
_ = let__4932

tmp6107 := Call1(__e, SymFn(symshen_4in_1_6), let__4931)


let__4933 := tmp6107
_ = let__4933

tmp6108 := Call2(__e, SymFn(symshen_4comb), let__4933, let__4932)


ifres6104 = tmp6108


}

let__4930 := ifres6104
_ = let__4930

tmp6134 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4930)


if True == tmp6134 {
tmp6110 := Call1(__e, SymFn(symshen_4_5number_6), V2889)


let__4935 := tmp6110
_ = let__4935

tmp6116 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4935)


var ifres6111 Obj

if True == tmp6116 {
tmp6112 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6111 = tmp6112


} else {
tmp6113 := Call1(__e, SymFn(symshen_4_5_1out), let__4935)


let__4936 := tmp6113
_ = let__4936

tmp6114 := Call1(__e, SymFn(symshen_4in_1_6), let__4935)


let__4937 := tmp6114
_ = let__4937

tmp6115 := Call2(__e, SymFn(symshen_4comb), let__4937, let__4936)


ifres6111 = tmp6115


}

let__4934 := ifres6111
_ = let__4934

tmp6132 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4934)


if True == tmp6132 {
tmp6117 := Call1(__e, SymFn(symshen_4_5sym_6), V2889)


let__4939 := tmp6117
_ = let__4939

tmp6128 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4939)


var ifres6118 Obj

if True == tmp6128 {
tmp6119 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6118 = tmp6119


} else {
tmp6120 := Call1(__e, SymFn(symshen_4_5_1out), let__4939)


let__4940 := tmp6120
_ = let__4940

tmp6121 := Call1(__e, SymFn(symshen_4in_1_6), let__4939)


let__4941 := tmp6121
_ = let__4941

tmp6126 := GEq(__e, let__4940, MakeString("<>"))

var ifres6122 Obj

if True == tmp6126 {
tmp6123 := GCons(__e, MakeInteger(0), Nil)

tmp6124 := GCons(__e, symvector, tmp6123)

ifres6122 = tmp6124


} else {
tmp6125 := PrimGuard1(__e, symintern, PrimIntern, let__4940)

ifres6122 = tmp6125


}

tmp6127 := Call2(__e, SymFn(symshen_4comb), let__4941, ifres6122)


ifres6118 = tmp6127


}

let__4938 := ifres6118
_ = let__4938

tmp6130 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4938)


if True == tmp6130 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4938
}


} else {
return let__4934
}


} else {
return let__4930
}


panic("unreachable")
})

tmp6135 := Call2(__e, ns2_1set, symshen_4_5atom_6, tmp6102)


_ = tmp6135

tmp6136 := MakeNative1(func(__e *ControlFlow, V2902 Obj) Obj {
tmp6137 := Call1(__e, SymFn(symshen_4_5alpha_6), V2902)


let__4943 := tmp6137
_ = let__4943

tmp6150 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4943)


var ifres6138 Obj

if True == tmp6150 {
tmp6139 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6138 = tmp6139


} else {
tmp6140 := Call1(__e, SymFn(symshen_4_5_1out), let__4943)


let__4944 := tmp6140
_ = let__4944

tmp6141 := Call1(__e, SymFn(symshen_4in_1_6), let__4943)


let__4945 := tmp6141
_ = let__4945

tmp6142 := Call1(__e, SymFn(symshen_4_5alphanums_6), let__4945)


let__4946 := tmp6142
_ = let__4946

tmp6149 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4946)


var ifres6143 Obj

if True == tmp6149 {
tmp6144 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6143 = tmp6144


} else {
tmp6145 := Call1(__e, SymFn(symshen_4_5_1out), let__4946)


let__4947 := tmp6145
_ = let__4947

tmp6146 := Call1(__e, SymFn(symshen_4in_1_6), let__4946)


let__4948 := tmp6146
_ = let__4948

tmp6148 := Call2(__e, SymFn(symshen_4comb), let__4948, PrimGuard2(__e, symcn, PrimStringConcat, let__4944, let__4947))


ifres6143 = tmp6148


}

ifres6138 = ifres6143


}

let__4942 := ifres6138
_ = let__4942

tmp6152 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4942)


if True == tmp6152 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4942
}


panic("unreachable")
})

tmp6153 := Call2(__e, ns2_1set, symshen_4_5sym_6, tmp6136)


_ = tmp6153

tmp6154 := MakeNative1(func(__e *ControlFlow, V2910 Obj) Obj {
tmp6164 := GConsP(__e, V2910)

var ifres6155 Obj

if True == tmp6164 {
tmp6156 := Call1(__e, SymFn(symhead), V2910)


let__4950 := tmp6156
_ = let__4950

tmp6157 := Call1(__e, SymFn(symtail), V2910)


let__4951 := tmp6157
_ = let__4951

tmp6162 := Call1(__e, SymFn(symshen_4alpha_2), let__4950)


var ifres6158 Obj

if True == tmp6162 {
tmp6159 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, let__4950)

tmp6160 := Call2(__e, SymFn(symshen_4comb), let__4951, tmp6159)


ifres6158 = tmp6160


} else {
tmp6161 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6158 = tmp6161


}

ifres6155 = ifres6158


} else {
tmp6163 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6155 = tmp6163


}

let__4949 := ifres6155
_ = let__4949

tmp6166 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4949)


if True == tmp6166 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4949
}


panic("unreachable")
})

tmp6167 := Call2(__e, ns2_1set, symshen_4_5alpha_6, tmp6154)


_ = tmp6167

tmp6168 := MakeNative1(func(__e *ControlFlow, V2914 Obj) Obj {
tmp6172 := Call1(__e, SymFn(symshen_4lowercase_2), V2914)


if True == tmp6172 {
return True
} else {
tmp6170 := Call1(__e, SymFn(symshen_4uppercase_2), V2914)


if True == tmp6170 {
return True
} else {
return TailCall1(__e, SymFn(symshen_4misc_2), V2914)
}


}


panic("unreachable")
})

tmp6173 := Call2(__e, ns2_1set, symshen_4alpha_2, tmp6168)


_ = tmp6173

tmp6174 := MakeNative1(func(__e *ControlFlow, V2915 Obj) Obj {
if True == GGe(__e, V2915, MakeInteger(97)) {
return GLe(__e, V2915, MakeInteger(122))
} else {
return False
}


panic("unreachable")
})

tmp6177 := Call2(__e, ns2_1set, symshen_4lowercase_2, tmp6174)


_ = tmp6177

tmp6178 := MakeNative1(func(__e *ControlFlow, V2916 Obj) Obj {
if True == GGe(__e, V2916, MakeInteger(65)) {
return GLe(__e, V2916, MakeInteger(90))
} else {
return False
}


panic("unreachable")
})

tmp6181 := Call2(__e, ns2_1set, symshen_4uppercase_2, tmp6178)


_ = tmp6181

tmp6182 := MakeNative1(func(__e *ControlFlow, V2917 Obj) Obj {
tmp6183 := GCons(__e, MakeInteger(96), Nil)

tmp6184 := GCons(__e, MakeInteger(35), tmp6183)

tmp6185 := GCons(__e, MakeInteger(39), tmp6184)

tmp6186 := GCons(__e, MakeInteger(37), tmp6185)

tmp6187 := GCons(__e, MakeInteger(38), tmp6186)

tmp6188 := GCons(__e, MakeInteger(60), tmp6187)

tmp6189 := GCons(__e, MakeInteger(62), tmp6188)

tmp6190 := GCons(__e, MakeInteger(46), tmp6189)

tmp6191 := GCons(__e, MakeInteger(126), tmp6190)

tmp6192 := GCons(__e, MakeInteger(64), tmp6191)

tmp6193 := GCons(__e, MakeInteger(33), tmp6192)

tmp6194 := GCons(__e, MakeInteger(36), tmp6193)

tmp6195 := GCons(__e, MakeInteger(63), tmp6194)

tmp6196 := GCons(__e, MakeInteger(95), tmp6195)

tmp6197 := GCons(__e, MakeInteger(43), tmp6196)

tmp6198 := GCons(__e, MakeInteger(47), tmp6197)

tmp6199 := GCons(__e, MakeInteger(42), tmp6198)

tmp6200 := GCons(__e, MakeInteger(45), tmp6199)

tmp6201 := GCons(__e, MakeInteger(61), tmp6200)

return TailCall2(__e, SymFn(symelement_2), V2917, tmp6201)


panic("unreachable")
})

tmp6202 := Call2(__e, ns2_1set, symshen_4misc_2, tmp6182)


_ = tmp6202

tmp6203 := MakeNative1(func(__e *ControlFlow, V2918 Obj) Obj {
tmp6204 := Call1(__e, SymFn(symshen_4_5alphanum_6), V2918)


let__4953 := tmp6204
_ = let__4953

tmp6217 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4953)


var ifres6205 Obj

if True == tmp6217 {
tmp6206 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6205 = tmp6206


} else {
tmp6207 := Call1(__e, SymFn(symshen_4_5_1out), let__4953)


let__4954 := tmp6207
_ = let__4954

tmp6208 := Call1(__e, SymFn(symshen_4in_1_6), let__4953)


let__4955 := tmp6208
_ = let__4955

tmp6209 := Call1(__e, SymFn(symshen_4_5alphanums_6), let__4955)


let__4956 := tmp6209
_ = let__4956

tmp6216 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4956)


var ifres6210 Obj

if True == tmp6216 {
tmp6211 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6210 = tmp6211


} else {
tmp6212 := Call1(__e, SymFn(symshen_4_5_1out), let__4956)


let__4957 := tmp6212
_ = let__4957

tmp6213 := Call1(__e, SymFn(symshen_4in_1_6), let__4956)


let__4958 := tmp6213
_ = let__4958

tmp6215 := Call2(__e, SymFn(symshen_4comb), let__4958, PrimGuard2(__e, symcn, PrimStringConcat, let__4954, let__4957))


ifres6210 = tmp6215


}

ifres6205 = ifres6210


}

let__4952 := ifres6205
_ = let__4952

tmp6227 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4952)


if True == tmp6227 {
tmp6218 := Call1(__e, SymFn(sym_5e_6), V2918)


let__4960 := tmp6218
_ = let__4960

tmp6223 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4960)


var ifres6219 Obj

if True == tmp6223 {
tmp6220 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6219 = tmp6220


} else {
tmp6221 := Call1(__e, SymFn(symshen_4in_1_6), let__4960)


let__4961 := tmp6221
_ = let__4961

tmp6222 := Call2(__e, SymFn(symshen_4comb), let__4961, MakeString(""))


ifres6219 = tmp6222


}

let__4959 := ifres6219
_ = let__4959

tmp6225 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4959)


if True == tmp6225 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4959
}


} else {
return let__4952
}


panic("unreachable")
})

tmp6228 := Call2(__e, ns2_1set, symshen_4_5alphanums_6, tmp6203)


_ = tmp6228

tmp6229 := MakeNative1(func(__e *ControlFlow, V2929 Obj) Obj {
tmp6230 := Call1(__e, SymFn(symshen_4_5alpha_6), V2929)


let__4963 := tmp6230
_ = let__4963

tmp6236 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4963)


var ifres6231 Obj

if True == tmp6236 {
tmp6232 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6231 = tmp6232


} else {
tmp6233 := Call1(__e, SymFn(symshen_4_5_1out), let__4963)


let__4964 := tmp6233
_ = let__4964

tmp6234 := Call1(__e, SymFn(symshen_4in_1_6), let__4963)


let__4965 := tmp6234
_ = let__4965

tmp6235 := Call2(__e, SymFn(symshen_4comb), let__4965, let__4964)


ifres6231 = tmp6235


}

let__4962 := ifres6231
_ = let__4962

tmp6247 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4962)


if True == tmp6247 {
tmp6237 := Call1(__e, SymFn(symshen_4_5numeral_6), V2929)


let__4967 := tmp6237
_ = let__4967

tmp6243 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4967)


var ifres6238 Obj

if True == tmp6243 {
tmp6239 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6238 = tmp6239


} else {
tmp6240 := Call1(__e, SymFn(symshen_4_5_1out), let__4967)


let__4968 := tmp6240
_ = let__4968

tmp6241 := Call1(__e, SymFn(symshen_4in_1_6), let__4967)


let__4969 := tmp6241
_ = let__4969

tmp6242 := Call2(__e, SymFn(symshen_4comb), let__4969, let__4968)


ifres6238 = tmp6242


}

let__4966 := ifres6238
_ = let__4966

tmp6245 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4966)


if True == tmp6245 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4966
}


} else {
return let__4962
}


panic("unreachable")
})

tmp6248 := Call2(__e, ns2_1set, symshen_4_5alphanum_6, tmp6229)


_ = tmp6248

tmp6249 := MakeNative1(func(__e *ControlFlow, V2938 Obj) Obj {
tmp6259 := GConsP(__e, V2938)

var ifres6250 Obj

if True == tmp6259 {
tmp6251 := Call1(__e, SymFn(symhead), V2938)


let__4971 := tmp6251
_ = let__4971

tmp6252 := Call1(__e, SymFn(symtail), V2938)


let__4972 := tmp6252
_ = let__4972

tmp6257 := Call1(__e, SymFn(symshen_4digit_2), let__4971)


var ifres6253 Obj

if True == tmp6257 {
tmp6254 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, let__4971)

tmp6255 := Call2(__e, SymFn(symshen_4comb), let__4972, tmp6254)


ifres6253 = tmp6255


} else {
tmp6256 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6253 = tmp6256


}

ifres6250 = ifres6253


} else {
tmp6258 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6250 = tmp6258


}

let__4970 := ifres6250
_ = let__4970

tmp6261 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4970)


if True == tmp6261 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4970
}


panic("unreachable")
})

tmp6262 := Call2(__e, ns2_1set, symshen_4_5numeral_6, tmp6249)


_ = tmp6262

tmp6263 := MakeNative1(func(__e *ControlFlow, V2942 Obj) Obj {
if True == GGe(__e, V2942, MakeInteger(48)) {
return GLe(__e, V2942, MakeInteger(57))
} else {
return False
}


panic("unreachable")
})

tmp6266 := Call2(__e, ns2_1set, symshen_4digit_2, tmp6263)


_ = tmp6266

tmp6267 := MakeNative1(func(__e *ControlFlow, V2943 Obj) Obj {
tmp6268 := Call1(__e, SymFn(symshen_4_5dbq_6), V2943)


let__4974 := tmp6268
_ = let__4974

tmp6284 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4974)


var ifres6269 Obj

if True == tmp6284 {
tmp6270 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6269 = tmp6270


} else {
tmp6271 := Call1(__e, SymFn(symshen_4in_1_6), let__4974)


let__4975 := tmp6271
_ = let__4975

tmp6272 := Call1(__e, SymFn(symshen_4_5strcontents_6), let__4975)


let__4976 := tmp6272
_ = let__4976

tmp6283 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4976)


var ifres6273 Obj

if True == tmp6283 {
tmp6274 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6273 = tmp6274


} else {
tmp6275 := Call1(__e, SymFn(symshen_4_5_1out), let__4976)


let__4977 := tmp6275
_ = let__4977

tmp6276 := Call1(__e, SymFn(symshen_4in_1_6), let__4976)


let__4978 := tmp6276
_ = let__4978

tmp6277 := Call1(__e, SymFn(symshen_4_5dbq_6), let__4978)


let__4979 := tmp6277
_ = let__4979

tmp6282 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4979)


var ifres6278 Obj

if True == tmp6282 {
tmp6279 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6278 = tmp6279


} else {
tmp6280 := Call1(__e, SymFn(symshen_4in_1_6), let__4979)


let__4980 := tmp6280
_ = let__4980

tmp6281 := Call2(__e, SymFn(symshen_4comb), let__4980, let__4977)


ifres6278 = tmp6281


}

ifres6273 = ifres6278


}

ifres6269 = ifres6273


}

let__4973 := ifres6269
_ = let__4973

tmp6286 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4973)


if True == tmp6286 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4973
}


panic("unreachable")
})

tmp6287 := Call2(__e, ns2_1set, symshen_4_5str_6, tmp6267)


_ = tmp6287

tmp6288 := MakeNative1(func(__e *ControlFlow, V2952 Obj) Obj {
tmp6293 := Call2(__e, SymFn(symshen_4hds_a_2), V2952, MakeInteger(34))


var ifres6289 Obj

if True == tmp6293 {
tmp6290 := Call1(__e, SymFn(symtail), V2952)


let__4982 := tmp6290
_ = let__4982

tmp6291 := Call2(__e, SymFn(symshen_4comb), let__4982, symshen_4skip)


ifres6289 = tmp6291


} else {
tmp6292 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6289 = tmp6292


}

let__4981 := ifres6289
_ = let__4981

tmp6295 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4981)


if True == tmp6295 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4981
}


panic("unreachable")
})

tmp6296 := Call2(__e, ns2_1set, symshen_4_5dbq_6, tmp6288)


_ = tmp6296

tmp6297 := MakeNative1(func(__e *ControlFlow, V2955 Obj) Obj {
tmp6298 := Call1(__e, SymFn(symshen_4_5strc_6), V2955)


let__4984 := tmp6298
_ = let__4984

tmp6311 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4984)


var ifres6299 Obj

if True == tmp6311 {
tmp6300 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6299 = tmp6300


} else {
tmp6301 := Call1(__e, SymFn(symshen_4_5_1out), let__4984)


let__4985 := tmp6301
_ = let__4985

tmp6302 := Call1(__e, SymFn(symshen_4in_1_6), let__4984)


let__4986 := tmp6302
_ = let__4986

tmp6303 := Call1(__e, SymFn(symshen_4_5strcontents_6), let__4986)


let__4987 := tmp6303
_ = let__4987

tmp6310 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4987)


var ifres6304 Obj

if True == tmp6310 {
tmp6305 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6304 = tmp6305


} else {
tmp6306 := Call1(__e, SymFn(symshen_4_5_1out), let__4987)


let__4988 := tmp6306
_ = let__4988

tmp6307 := Call1(__e, SymFn(symshen_4in_1_6), let__4987)


let__4989 := tmp6307
_ = let__4989

tmp6309 := Call2(__e, SymFn(symshen_4comb), let__4989, PrimGuard2(__e, symcn, PrimStringConcat, let__4985, let__4988))


ifres6304 = tmp6309


}

ifres6299 = ifres6304


}

let__4983 := ifres6299
_ = let__4983

tmp6321 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4983)


if True == tmp6321 {
tmp6312 := Call1(__e, SymFn(sym_5e_6), V2955)


let__4991 := tmp6312
_ = let__4991

tmp6317 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4991)


var ifres6313 Obj

if True == tmp6317 {
tmp6314 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6313 = tmp6314


} else {
tmp6315 := Call1(__e, SymFn(symshen_4in_1_6), let__4991)


let__4992 := tmp6315
_ = let__4992

tmp6316 := Call2(__e, SymFn(symshen_4comb), let__4992, MakeString(""))


ifres6313 = tmp6316


}

let__4990 := ifres6313
_ = let__4990

tmp6319 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4990)


if True == tmp6319 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4990
}


} else {
return let__4983
}


panic("unreachable")
})

tmp6322 := Call2(__e, ns2_1set, symshen_4_5strcontents_6, tmp6297)


_ = tmp6322

tmp6323 := MakeNative1(func(__e *ControlFlow, V2966 Obj) Obj {
tmp6324 := Call1(__e, SymFn(symshen_4_5control_6), V2966)


let__4994 := tmp6324
_ = let__4994

tmp6330 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4994)


var ifres6325 Obj

if True == tmp6330 {
tmp6326 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6325 = tmp6326


} else {
tmp6327 := Call1(__e, SymFn(symshen_4_5_1out), let__4994)


let__4995 := tmp6327
_ = let__4995

tmp6328 := Call1(__e, SymFn(symshen_4in_1_6), let__4994)


let__4996 := tmp6328
_ = let__4996

tmp6329 := Call2(__e, SymFn(symshen_4comb), let__4996, let__4995)


ifres6325 = tmp6329


}

let__4993 := ifres6325
_ = let__4993

tmp6341 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4993)


if True == tmp6341 {
tmp6331 := Call1(__e, SymFn(symshen_4_5notdbq_6), V2966)


let__4998 := tmp6331
_ = let__4998

tmp6337 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4998)


var ifres6332 Obj

if True == tmp6337 {
tmp6333 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6332 = tmp6333


} else {
tmp6334 := Call1(__e, SymFn(symshen_4_5_1out), let__4998)


let__4999 := tmp6334
_ = let__4999

tmp6335 := Call1(__e, SymFn(symshen_4in_1_6), let__4998)


let__5000 := tmp6335
_ = let__5000

tmp6336 := Call2(__e, SymFn(symshen_4comb), let__5000, let__4999)


ifres6332 = tmp6336


}

let__4997 := ifres6332
_ = let__4997

tmp6339 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__4997)


if True == tmp6339 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__4997
}


} else {
return let__4993
}


panic("unreachable")
})

tmp6342 := Call2(__e, ns2_1set, symshen_4_5strc_6, tmp6323)


_ = tmp6342

tmp6343 := MakeNative1(func(__e *ControlFlow, V2975 Obj) Obj {
tmp6344 := Call1(__e, SymFn(symshen_4_5lowC_6), V2975)


let__5002 := tmp6344
_ = let__5002

tmp6366 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5002)


var ifres6345 Obj

if True == tmp6366 {
tmp6346 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6345 = tmp6346


} else {
tmp6347 := Call1(__e, SymFn(symshen_4in_1_6), let__5002)


let__5003 := tmp6347
_ = let__5003

tmp6348 := Call1(__e, SymFn(symshen_4_5hash_6), let__5003)


let__5004 := tmp6348
_ = let__5004

tmp6365 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5004)


var ifres6349 Obj

if True == tmp6365 {
tmp6350 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6349 = tmp6350


} else {
tmp6351 := Call1(__e, SymFn(symshen_4in_1_6), let__5004)


let__5005 := tmp6351
_ = let__5005

tmp6352 := Call1(__e, SymFn(symshen_4_5integer_6), let__5005)


let__5006 := tmp6352
_ = let__5006

tmp6364 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5006)


var ifres6353 Obj

if True == tmp6364 {
tmp6354 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6353 = tmp6354


} else {
tmp6355 := Call1(__e, SymFn(symshen_4_5_1out), let__5006)


let__5007 := tmp6355
_ = let__5007

tmp6356 := Call1(__e, SymFn(symshen_4in_1_6), let__5006)


let__5008 := tmp6356
_ = let__5008

tmp6357 := Call1(__e, SymFn(symshen_4_5semicolon_6), let__5008)


let__5009 := tmp6357
_ = let__5009

tmp6363 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5009)


var ifres6358 Obj

if True == tmp6363 {
tmp6359 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6358 = tmp6359


} else {
tmp6360 := Call1(__e, SymFn(symshen_4in_1_6), let__5009)


let__5010 := tmp6360
_ = let__5010

tmp6361 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, let__5007)

tmp6362 := Call2(__e, SymFn(symshen_4comb), let__5010, tmp6361)


ifres6358 = tmp6362


}

ifres6353 = ifres6358


}

ifres6349 = ifres6353


}

ifres6345 = ifres6349


}

let__5001 := ifres6345
_ = let__5001

tmp6368 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5001)


if True == tmp6368 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5001
}


panic("unreachable")
})

tmp6369 := Call2(__e, ns2_1set, symshen_4_5control_6, tmp6343)


_ = tmp6369

tmp6370 := MakeNative1(func(__e *ControlFlow, V2986 Obj) Obj {
tmp6381 := GConsP(__e, V2986)

var ifres6371 Obj

if True == tmp6381 {
tmp6372 := Call1(__e, SymFn(symhead), V2986)


let__5012 := tmp6372
_ = let__5012

tmp6373 := Call1(__e, SymFn(symtail), V2986)


let__5013 := tmp6373
_ = let__5013

tmp6378 := GEq(__e, let__5012, MakeInteger(34))

tmp6379 := GNot(__e, tmp6378)

var ifres6374 Obj

if True == tmp6379 {
tmp6375 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, let__5012)

tmp6376 := Call2(__e, SymFn(symshen_4comb), let__5013, tmp6375)


ifres6374 = tmp6376


} else {
tmp6377 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6374 = tmp6377


}

ifres6371 = ifres6374


} else {
tmp6380 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6371 = tmp6380


}

let__5011 := ifres6371
_ = let__5011

tmp6383 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5011)


if True == tmp6383 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5011
}


panic("unreachable")
})

tmp6384 := Call2(__e, ns2_1set, symshen_4_5notdbq_6, tmp6370)


_ = tmp6384

tmp6385 := MakeNative1(func(__e *ControlFlow, V2990 Obj) Obj {
tmp6390 := Call2(__e, SymFn(symshen_4hds_a_2), V2990, MakeInteger(99))


var ifres6386 Obj

if True == tmp6390 {
tmp6387 := Call1(__e, SymFn(symtail), V2990)


let__5015 := tmp6387
_ = let__5015

tmp6388 := Call2(__e, SymFn(symshen_4comb), let__5015, symshen_4skip)


ifres6386 = tmp6388


} else {
tmp6389 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6386 = tmp6389


}

let__5014 := ifres6386
_ = let__5014

tmp6392 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5014)


if True == tmp6392 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5014
}


panic("unreachable")
})

tmp6393 := Call2(__e, ns2_1set, symshen_4_5lowC_6, tmp6385)


_ = tmp6393

tmp6394 := MakeNative1(func(__e *ControlFlow, V2993 Obj) Obj {
tmp6399 := Call2(__e, SymFn(symshen_4hds_a_2), V2993, MakeInteger(35))


var ifres6395 Obj

if True == tmp6399 {
tmp6396 := Call1(__e, SymFn(symtail), V2993)


let__5017 := tmp6396
_ = let__5017

tmp6397 := Call2(__e, SymFn(symshen_4comb), let__5017, symshen_4skip)


ifres6395 = tmp6397


} else {
tmp6398 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6395 = tmp6398


}

let__5016 := ifres6395
_ = let__5016

tmp6401 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5016)


if True == tmp6401 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5016
}


panic("unreachable")
})

tmp6402 := Call2(__e, ns2_1set, symshen_4_5hash_6, tmp6394)


_ = tmp6402

tmp6403 := MakeNative1(func(__e *ControlFlow, V2996 Obj) Obj {
tmp6404 := Call1(__e, SymFn(symshen_4_5minus_6), V2996)


let__5019 := tmp6404
_ = let__5019

tmp6416 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5019)


var ifres6405 Obj

if True == tmp6416 {
tmp6406 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6405 = tmp6406


} else {
tmp6407 := Call1(__e, SymFn(symshen_4in_1_6), let__5019)


let__5020 := tmp6407
_ = let__5020

tmp6408 := Call1(__e, SymFn(symshen_4_5number_6), let__5020)


let__5021 := tmp6408
_ = let__5021

tmp6415 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5021)


var ifres6409 Obj

if True == tmp6415 {
tmp6410 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6409 = tmp6410


} else {
tmp6411 := Call1(__e, SymFn(symshen_4_5_1out), let__5021)


let__5022 := tmp6411
_ = let__5022

tmp6412 := Call1(__e, SymFn(symshen_4in_1_6), let__5021)


let__5023 := tmp6412
_ = let__5023

tmp6414 := Call2(__e, SymFn(symshen_4comb), let__5023, GSub(__e, MakeInteger(0), let__5022))


ifres6409 = tmp6414


}

ifres6405 = ifres6409


}

let__5018 := ifres6405
_ = let__5018

tmp6459 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5018)


if True == tmp6459 {
tmp6417 := Call1(__e, SymFn(symshen_4_5plus_6), V2996)


let__5025 := tmp6417
_ = let__5025

tmp6428 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5025)


var ifres6418 Obj

if True == tmp6428 {
tmp6419 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6418 = tmp6419


} else {
tmp6420 := Call1(__e, SymFn(symshen_4in_1_6), let__5025)


let__5026 := tmp6420
_ = let__5026

tmp6421 := Call1(__e, SymFn(symshen_4_5number_6), let__5026)


let__5027 := tmp6421
_ = let__5027

tmp6427 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5027)


var ifres6422 Obj

if True == tmp6427 {
tmp6423 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6422 = tmp6423


} else {
tmp6424 := Call1(__e, SymFn(symshen_4_5_1out), let__5027)


let__5028 := tmp6424
_ = let__5028

tmp6425 := Call1(__e, SymFn(symshen_4in_1_6), let__5027)


let__5029 := tmp6425
_ = let__5029

tmp6426 := Call2(__e, SymFn(symshen_4comb), let__5029, let__5028)


ifres6422 = tmp6426


}

ifres6418 = ifres6422


}

let__5024 := ifres6418
_ = let__5024

tmp6457 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5024)


if True == tmp6457 {
tmp6429 := Call1(__e, SymFn(symshen_4_5e_1number_6), V2996)


let__5031 := tmp6429
_ = let__5031

tmp6435 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5031)


var ifres6430 Obj

if True == tmp6435 {
tmp6431 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6430 = tmp6431


} else {
tmp6432 := Call1(__e, SymFn(symshen_4_5_1out), let__5031)


let__5032 := tmp6432
_ = let__5032

tmp6433 := Call1(__e, SymFn(symshen_4in_1_6), let__5031)


let__5033 := tmp6433
_ = let__5033

tmp6434 := Call2(__e, SymFn(symshen_4comb), let__5033, let__5032)


ifres6430 = tmp6434


}

let__5030 := ifres6430
_ = let__5030

tmp6455 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5030)


if True == tmp6455 {
tmp6436 := Call1(__e, SymFn(symshen_4_5float_6), V2996)


let__5035 := tmp6436
_ = let__5035

tmp6442 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5035)


var ifres6437 Obj

if True == tmp6442 {
tmp6438 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6437 = tmp6438


} else {
tmp6439 := Call1(__e, SymFn(symshen_4_5_1out), let__5035)


let__5036 := tmp6439
_ = let__5036

tmp6440 := Call1(__e, SymFn(symshen_4in_1_6), let__5035)


let__5037 := tmp6440
_ = let__5037

tmp6441 := Call2(__e, SymFn(symshen_4comb), let__5037, let__5036)


ifres6437 = tmp6441


}

let__5034 := ifres6437
_ = let__5034

tmp6453 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5034)


if True == tmp6453 {
tmp6443 := Call1(__e, SymFn(symshen_4_5integer_6), V2996)


let__5039 := tmp6443
_ = let__5039

tmp6449 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5039)


var ifres6444 Obj

if True == tmp6449 {
tmp6445 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6444 = tmp6445


} else {
tmp6446 := Call1(__e, SymFn(symshen_4_5_1out), let__5039)


let__5040 := tmp6446
_ = let__5040

tmp6447 := Call1(__e, SymFn(symshen_4in_1_6), let__5039)


let__5041 := tmp6447
_ = let__5041

tmp6448 := Call2(__e, SymFn(symshen_4comb), let__5041, let__5040)


ifres6444 = tmp6448


}

let__5038 := ifres6444
_ = let__5038

tmp6451 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5038)


if True == tmp6451 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5038
}


} else {
return let__5034
}


} else {
return let__5030
}


} else {
return let__5024
}


} else {
return let__5018
}


panic("unreachable")
})

tmp6460 := Call2(__e, ns2_1set, symshen_4_5number_6, tmp6403)


_ = tmp6460

tmp6461 := MakeNative1(func(__e *ControlFlow, V3021 Obj) Obj {
tmp6466 := Call2(__e, SymFn(symshen_4hds_a_2), V3021, MakeInteger(45))


var ifres6462 Obj

if True == tmp6466 {
tmp6463 := Call1(__e, SymFn(symtail), V3021)


let__5043 := tmp6463
_ = let__5043

tmp6464 := Call2(__e, SymFn(symshen_4comb), let__5043, symshen_4skip)


ifres6462 = tmp6464


} else {
tmp6465 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6462 = tmp6465


}

let__5042 := ifres6462
_ = let__5042

tmp6468 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5042)


if True == tmp6468 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5042
}


panic("unreachable")
})

tmp6469 := Call2(__e, ns2_1set, symshen_4_5minus_6, tmp6461)


_ = tmp6469

tmp6470 := MakeNative1(func(__e *ControlFlow, V3024 Obj) Obj {
tmp6475 := Call2(__e, SymFn(symshen_4hds_a_2), V3024, MakeInteger(43))


var ifres6471 Obj

if True == tmp6475 {
tmp6472 := Call1(__e, SymFn(symtail), V3024)


let__5045 := tmp6472
_ = let__5045

tmp6473 := Call2(__e, SymFn(symshen_4comb), let__5045, symshen_4skip)


ifres6471 = tmp6473


} else {
tmp6474 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6471 = tmp6474


}

let__5044 := ifres6471
_ = let__5044

tmp6477 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5044)


if True == tmp6477 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5044
}


panic("unreachable")
})

tmp6478 := Call2(__e, ns2_1set, symshen_4_5plus_6, tmp6470)


_ = tmp6478

tmp6479 := MakeNative1(func(__e *ControlFlow, V3027 Obj) Obj {
tmp6480 := Call1(__e, SymFn(symshen_4_5digits_6), V3027)


let__5047 := tmp6480
_ = let__5047

tmp6487 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5047)


var ifres6481 Obj

if True == tmp6487 {
tmp6482 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6481 = tmp6482


} else {
tmp6483 := Call1(__e, SymFn(symshen_4_5_1out), let__5047)


let__5048 := tmp6483
_ = let__5048

tmp6484 := Call1(__e, SymFn(symshen_4in_1_6), let__5047)


let__5049 := tmp6484
_ = let__5049

tmp6485 := Call1(__e, SymFn(symshen_4compute_1integer), let__5048)


tmp6486 := Call2(__e, SymFn(symshen_4comb), let__5049, tmp6485)


ifres6481 = tmp6486


}

let__5046 := ifres6481
_ = let__5046

tmp6489 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5046)


if True == tmp6489 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5046
}


panic("unreachable")
})

tmp6490 := Call2(__e, ns2_1set, symshen_4_5integer_6, tmp6479)


_ = tmp6490

tmp6491 := MakeNative1(func(__e *ControlFlow, V3032 Obj) Obj {
tmp6492 := Call1(__e, SymFn(symshen_4_5digit_6), V3032)


let__5051 := tmp6492
_ = let__5051

tmp6505 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5051)


var ifres6493 Obj

if True == tmp6505 {
tmp6494 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6493 = tmp6494


} else {
tmp6495 := Call1(__e, SymFn(symshen_4_5_1out), let__5051)


let__5052 := tmp6495
_ = let__5052

tmp6496 := Call1(__e, SymFn(symshen_4in_1_6), let__5051)


let__5053 := tmp6496
_ = let__5053

tmp6497 := Call1(__e, SymFn(symshen_4_5digits_6), let__5053)


let__5054 := tmp6497
_ = let__5054

tmp6504 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5054)


var ifres6498 Obj

if True == tmp6504 {
tmp6499 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6498 = tmp6499


} else {
tmp6500 := Call1(__e, SymFn(symshen_4_5_1out), let__5054)


let__5055 := tmp6500
_ = let__5055

tmp6501 := Call1(__e, SymFn(symshen_4in_1_6), let__5054)


let__5056 := tmp6501
_ = let__5056

tmp6502 := GCons(__e, let__5052, let__5055)

tmp6503 := Call2(__e, SymFn(symshen_4comb), let__5056, tmp6502)


ifres6498 = tmp6503


}

ifres6493 = ifres6498


}

let__5050 := ifres6493
_ = let__5050

tmp6517 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5050)


if True == tmp6517 {
tmp6506 := Call1(__e, SymFn(symshen_4_5digit_6), V3032)


let__5058 := tmp6506
_ = let__5058

tmp6513 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5058)


var ifres6507 Obj

if True == tmp6513 {
tmp6508 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6507 = tmp6508


} else {
tmp6509 := Call1(__e, SymFn(symshen_4_5_1out), let__5058)


let__5059 := tmp6509
_ = let__5059

tmp6510 := Call1(__e, SymFn(symshen_4in_1_6), let__5058)


let__5060 := tmp6510
_ = let__5060

tmp6511 := GCons(__e, let__5059, Nil)

tmp6512 := Call2(__e, SymFn(symshen_4comb), let__5060, tmp6511)


ifres6507 = tmp6512


}

let__5057 := ifres6507
_ = let__5057

tmp6515 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5057)


if True == tmp6515 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5057
}


} else {
return let__5050
}


panic("unreachable")
})

tmp6518 := Call2(__e, ns2_1set, symshen_4_5digits_6, tmp6491)


_ = tmp6518

tmp6519 := MakeNative1(func(__e *ControlFlow, V3044 Obj) Obj {
tmp6529 := GConsP(__e, V3044)

var ifres6520 Obj

if True == tmp6529 {
tmp6521 := Call1(__e, SymFn(symhead), V3044)


let__5062 := tmp6521
_ = let__5062

tmp6522 := Call1(__e, SymFn(symtail), V3044)


let__5063 := tmp6522
_ = let__5063

tmp6527 := Call1(__e, SymFn(symshen_4digit_2), let__5062)


var ifres6523 Obj

if True == tmp6527 {
tmp6524 := Call1(__e, SymFn(symshen_4byte_1_6digit), let__5062)


tmp6525 := Call2(__e, SymFn(symshen_4comb), let__5063, tmp6524)


ifres6523 = tmp6525


} else {
tmp6526 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6523 = tmp6526


}

ifres6520 = ifres6523


} else {
tmp6528 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6520 = tmp6528


}

let__5061 := ifres6520
_ = let__5061

tmp6531 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5061)


if True == tmp6531 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5061
}


panic("unreachable")
})

tmp6532 := Call2(__e, ns2_1set, symshen_4_5digit_6, tmp6519)


_ = tmp6532

tmp6533 := MakeNative1(func(__e *ControlFlow, V3048 Obj) Obj {
return GSub(__e, V3048, MakeInteger(48))
panic("unreachable")
})

tmp6534 := Call2(__e, ns2_1set, symshen_4byte_1_6digit, tmp6533)


_ = tmp6534

tmp6535 := MakeNative1(func(__e *ControlFlow, V3049 Obj) Obj {
tmp6536 := Call1(__e, SymFn(symreverse), V3049)


return TailCall2(__e, SymFn(symshen_4compute_1integer_1h), tmp6536, MakeInteger(0))


panic("unreachable")
})

tmp6537 := Call2(__e, ns2_1set, symshen_4compute_1integer, tmp6535)


_ = tmp6537

tmp6538 := MakeNative2(func(__e *ControlFlow, V3052, V3053 Obj) Obj {
tmp6548 := GEq(__e, Nil, V3052)

if True == tmp6548 {
return MakeInteger(0)
} else {
tmp6546 := GConsP(__e, V3052)

if True == tmp6546 {
tmp6539 := Call2(__e, SymFn(symshen_4expt), MakeInteger(10), V3053)


tmp6540 := GHd(__e, V3052)

tmp6541 := GMul(__e, tmp6539, tmp6540)

tmp6542 := GTl(__e, V3052)

tmp6544 := Call2(__e, SymFn(symshen_4compute_1integer_1h), tmp6542, GAdd(__e, V3053, MakeInteger(1)))


return GAdd(__e, tmp6541, tmp6544)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4compute_1integer_1h)
}


}


panic("unreachable")
})

tmp6549 := Call2(__e, ns2_1set, symshen_4compute_1integer_1h, tmp6538)


_ = tmp6549

tmp6550 := MakeNative2(func(__e *ControlFlow, V3056, V3057 Obj) Obj {
tmp6558 := GEq(__e, MakeInteger(0), V3057)

if True == tmp6558 {
return MakeInteger(1)
} else {
if True == GGt(__e, V3057, MakeInteger(0)) {
tmp6552 := Call2(__e, SymFn(symshen_4expt), V3056, GSub(__e, V3057, MakeInteger(1)))


return GMul(__e, V3056, tmp6552)


} else {
tmp6554 := Call2(__e, SymFn(symshen_4expt), V3056, GAdd(__e, V3057, MakeInteger(1)))


return GDiv(__e, tmp6554, V3056)


}


}


panic("unreachable")
})

tmp6559 := Call2(__e, ns2_1set, symshen_4expt, tmp6550)


_ = tmp6559

tmp6560 := MakeNative1(func(__e *ControlFlow, V3058 Obj) Obj {
tmp6561 := Call1(__e, SymFn(symshen_4_5integer_6), V3058)


let__5065 := tmp6561
_ = let__5065

tmp6579 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5065)


var ifres6562 Obj

if True == tmp6579 {
tmp6563 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6562 = tmp6563


} else {
tmp6564 := Call1(__e, SymFn(symshen_4_5_1out), let__5065)


let__5066 := tmp6564
_ = let__5066

tmp6565 := Call1(__e, SymFn(symshen_4in_1_6), let__5065)


let__5067 := tmp6565
_ = let__5067

tmp6566 := Call1(__e, SymFn(symshen_4_5stop_6), let__5067)


let__5068 := tmp6566
_ = let__5068

tmp6578 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5068)


var ifres6567 Obj

if True == tmp6578 {
tmp6568 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6567 = tmp6568


} else {
tmp6569 := Call1(__e, SymFn(symshen_4in_1_6), let__5068)


let__5069 := tmp6569
_ = let__5069

tmp6570 := Call1(__e, SymFn(symshen_4_5fraction_6), let__5069)


let__5070 := tmp6570
_ = let__5070

tmp6577 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5070)


var ifres6571 Obj

if True == tmp6577 {
tmp6572 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6571 = tmp6572


} else {
tmp6573 := Call1(__e, SymFn(symshen_4_5_1out), let__5070)


let__5071 := tmp6573
_ = let__5071

tmp6574 := Call1(__e, SymFn(symshen_4in_1_6), let__5070)


let__5072 := tmp6574
_ = let__5072

tmp6576 := Call2(__e, SymFn(symshen_4comb), let__5072, GAdd(__e, let__5066, let__5071))


ifres6571 = tmp6576


}

ifres6567 = ifres6571


}

ifres6562 = ifres6567


}

let__5064 := ifres6562
_ = let__5064

tmp6595 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5064)


if True == tmp6595 {
tmp6580 := Call1(__e, SymFn(symshen_4_5stop_6), V3058)


let__5074 := tmp6580
_ = let__5074

tmp6591 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5074)


var ifres6581 Obj

if True == tmp6591 {
tmp6582 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6581 = tmp6582


} else {
tmp6583 := Call1(__e, SymFn(symshen_4in_1_6), let__5074)


let__5075 := tmp6583
_ = let__5075

tmp6584 := Call1(__e, SymFn(symshen_4_5fraction_6), let__5075)


let__5076 := tmp6584
_ = let__5076

tmp6590 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5076)


var ifres6585 Obj

if True == tmp6590 {
tmp6586 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6585 = tmp6586


} else {
tmp6587 := Call1(__e, SymFn(symshen_4_5_1out), let__5076)


let__5077 := tmp6587
_ = let__5077

tmp6588 := Call1(__e, SymFn(symshen_4in_1_6), let__5076)


let__5078 := tmp6588
_ = let__5078

tmp6589 := Call2(__e, SymFn(symshen_4comb), let__5078, let__5077)


ifres6585 = tmp6589


}

ifres6581 = ifres6585


}

let__5073 := ifres6581
_ = let__5073

tmp6593 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5073)


if True == tmp6593 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5073
}


} else {
return let__5064
}


panic("unreachable")
})

tmp6596 := Call2(__e, ns2_1set, symshen_4_5float_6, tmp6560)


_ = tmp6596

tmp6597 := MakeNative1(func(__e *ControlFlow, V3074 Obj) Obj {
tmp6602 := Call2(__e, SymFn(symshen_4hds_a_2), V3074, MakeInteger(46))


var ifres6598 Obj

if True == tmp6602 {
tmp6599 := Call1(__e, SymFn(symtail), V3074)


let__5080 := tmp6599
_ = let__5080

tmp6600 := Call2(__e, SymFn(symshen_4comb), let__5080, symshen_4skip)


ifres6598 = tmp6600


} else {
tmp6601 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6598 = tmp6601


}

let__5079 := ifres6598
_ = let__5079

tmp6604 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5079)


if True == tmp6604 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5079
}


panic("unreachable")
})

tmp6605 := Call2(__e, ns2_1set, symshen_4_5stop_6, tmp6597)


_ = tmp6605

tmp6606 := MakeNative1(func(__e *ControlFlow, V3077 Obj) Obj {
tmp6607 := Call1(__e, SymFn(symshen_4_5digits_6), V3077)


let__5082 := tmp6607
_ = let__5082

tmp6614 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5082)


var ifres6608 Obj

if True == tmp6614 {
tmp6609 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6608 = tmp6609


} else {
tmp6610 := Call1(__e, SymFn(symshen_4_5_1out), let__5082)


let__5083 := tmp6610
_ = let__5083

tmp6611 := Call1(__e, SymFn(symshen_4in_1_6), let__5082)


let__5084 := tmp6611
_ = let__5084

tmp6612 := Call1(__e, SymFn(symshen_4compute_1fraction), let__5083)


tmp6613 := Call2(__e, SymFn(symshen_4comb), let__5084, tmp6612)


ifres6608 = tmp6613


}

let__5081 := ifres6608
_ = let__5081

tmp6616 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5081)


if True == tmp6616 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5081
}


panic("unreachable")
})

tmp6617 := Call2(__e, ns2_1set, symshen_4_5fraction_6, tmp6606)


_ = tmp6617

tmp6618 := MakeNative1(func(__e *ControlFlow, V3082 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4compute_1fraction_1h), V3082, MakeInteger(-1))
panic("unreachable")
})

tmp6619 := Call2(__e, ns2_1set, symshen_4compute_1fraction, tmp6618)


_ = tmp6619

tmp6620 := MakeNative2(func(__e *ControlFlow, V3085, V3086 Obj) Obj {
tmp6630 := GEq(__e, Nil, V3085)

if True == tmp6630 {
return MakeInteger(0)
} else {
tmp6628 := GConsP(__e, V3085)

if True == tmp6628 {
tmp6621 := Call2(__e, SymFn(symshen_4expt), MakeInteger(10), V3086)


tmp6622 := GHd(__e, V3085)

tmp6623 := GMul(__e, tmp6621, tmp6622)

tmp6624 := GTl(__e, V3085)

tmp6626 := Call2(__e, SymFn(symshen_4compute_1fraction_1h), tmp6624, GSub(__e, V3086, MakeInteger(1)))


return GAdd(__e, tmp6623, tmp6626)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4compute_1fraction_1h)
}


}


panic("unreachable")
})

tmp6631 := Call2(__e, ns2_1set, symshen_4compute_1fraction_1h, tmp6620)


_ = tmp6631

tmp6632 := MakeNative1(func(__e *ControlFlow, V3087 Obj) Obj {
tmp6633 := Call1(__e, SymFn(symshen_4_5float_6), V3087)


let__5086 := tmp6633
_ = let__5086

tmp6651 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5086)


var ifres6634 Obj

if True == tmp6651 {
tmp6635 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6634 = tmp6635


} else {
tmp6636 := Call1(__e, SymFn(symshen_4_5_1out), let__5086)


let__5087 := tmp6636
_ = let__5087

tmp6637 := Call1(__e, SymFn(symshen_4in_1_6), let__5086)


let__5088 := tmp6637
_ = let__5088

tmp6638 := Call1(__e, SymFn(symshen_4_5lowE_6), let__5088)


let__5089 := tmp6638
_ = let__5089

tmp6650 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5089)


var ifres6639 Obj

if True == tmp6650 {
tmp6640 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6639 = tmp6640


} else {
tmp6641 := Call1(__e, SymFn(symshen_4in_1_6), let__5089)


let__5090 := tmp6641
_ = let__5090

tmp6642 := Call1(__e, SymFn(symshen_4_5log10_6), let__5090)


let__5091 := tmp6642
_ = let__5091

tmp6649 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5091)


var ifres6643 Obj

if True == tmp6649 {
tmp6644 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6643 = tmp6644


} else {
tmp6645 := Call1(__e, SymFn(symshen_4_5_1out), let__5091)


let__5092 := tmp6645
_ = let__5092

tmp6646 := Call1(__e, SymFn(symshen_4in_1_6), let__5091)


let__5093 := tmp6646
_ = let__5093

tmp6647 := Call2(__e, SymFn(symshen_4compute_1E), let__5087, let__5092)


tmp6648 := Call2(__e, SymFn(symshen_4comb), let__5093, tmp6647)


ifres6643 = tmp6648


}

ifres6639 = ifres6643


}

ifres6634 = ifres6639


}

let__5085 := ifres6634
_ = let__5085

tmp6674 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5085)


if True == tmp6674 {
tmp6652 := Call1(__e, SymFn(symshen_4_5integer_6), V3087)


let__5095 := tmp6652
_ = let__5095

tmp6670 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5095)


var ifres6653 Obj

if True == tmp6670 {
tmp6654 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6653 = tmp6654


} else {
tmp6655 := Call1(__e, SymFn(symshen_4_5_1out), let__5095)


let__5096 := tmp6655
_ = let__5096

tmp6656 := Call1(__e, SymFn(symshen_4in_1_6), let__5095)


let__5097 := tmp6656
_ = let__5097

tmp6657 := Call1(__e, SymFn(symshen_4_5lowE_6), let__5097)


let__5098 := tmp6657
_ = let__5098

tmp6669 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5098)


var ifres6658 Obj

if True == tmp6669 {
tmp6659 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6658 = tmp6659


} else {
tmp6660 := Call1(__e, SymFn(symshen_4in_1_6), let__5098)


let__5099 := tmp6660
_ = let__5099

tmp6661 := Call1(__e, SymFn(symshen_4_5log10_6), let__5099)


let__5100 := tmp6661
_ = let__5100

tmp6668 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5100)


var ifres6662 Obj

if True == tmp6668 {
tmp6663 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6662 = tmp6663


} else {
tmp6664 := Call1(__e, SymFn(symshen_4_5_1out), let__5100)


let__5101 := tmp6664
_ = let__5101

tmp6665 := Call1(__e, SymFn(symshen_4in_1_6), let__5100)


let__5102 := tmp6665
_ = let__5102

tmp6666 := Call2(__e, SymFn(symshen_4compute_1E), let__5096, let__5101)


tmp6667 := Call2(__e, SymFn(symshen_4comb), let__5102, tmp6666)


ifres6662 = tmp6667


}

ifres6658 = ifres6662


}

ifres6653 = ifres6658


}

let__5094 := ifres6653
_ = let__5094

tmp6672 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5094)


if True == tmp6672 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5094
}


} else {
return let__5085
}


panic("unreachable")
})

tmp6675 := Call2(__e, ns2_1set, symshen_4_5e_1number_6, tmp6632)


_ = tmp6675

tmp6676 := MakeNative1(func(__e *ControlFlow, V3106 Obj) Obj {
tmp6677 := Call1(__e, SymFn(symshen_4_5plus_6), V3106)


let__5104 := tmp6677
_ = let__5104

tmp6688 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5104)


var ifres6678 Obj

if True == tmp6688 {
tmp6679 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6678 = tmp6679


} else {
tmp6680 := Call1(__e, SymFn(symshen_4in_1_6), let__5104)


let__5105 := tmp6680
_ = let__5105

tmp6681 := Call1(__e, SymFn(symshen_4_5log10_6), let__5105)


let__5106 := tmp6681
_ = let__5106

tmp6687 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5106)


var ifres6682 Obj

if True == tmp6687 {
tmp6683 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6682 = tmp6683


} else {
tmp6684 := Call1(__e, SymFn(symshen_4_5_1out), let__5106)


let__5107 := tmp6684
_ = let__5107

tmp6685 := Call1(__e, SymFn(symshen_4in_1_6), let__5106)


let__5108 := tmp6685
_ = let__5108

tmp6686 := Call2(__e, SymFn(symshen_4comb), let__5108, let__5107)


ifres6682 = tmp6686


}

ifres6678 = ifres6682


}

let__5103 := ifres6678
_ = let__5103

tmp6714 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5103)


if True == tmp6714 {
tmp6689 := Call1(__e, SymFn(symshen_4_5minus_6), V3106)


let__5110 := tmp6689
_ = let__5110

tmp6701 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5110)


var ifres6690 Obj

if True == tmp6701 {
tmp6691 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6690 = tmp6691


} else {
tmp6692 := Call1(__e, SymFn(symshen_4in_1_6), let__5110)


let__5111 := tmp6692
_ = let__5111

tmp6693 := Call1(__e, SymFn(symshen_4_5log10_6), let__5111)


let__5112 := tmp6693
_ = let__5112

tmp6700 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5112)


var ifres6694 Obj

if True == tmp6700 {
tmp6695 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6694 = tmp6695


} else {
tmp6696 := Call1(__e, SymFn(symshen_4_5_1out), let__5112)


let__5113 := tmp6696
_ = let__5113

tmp6697 := Call1(__e, SymFn(symshen_4in_1_6), let__5112)


let__5114 := tmp6697
_ = let__5114

tmp6699 := Call2(__e, SymFn(symshen_4comb), let__5114, GSub(__e, MakeInteger(0), let__5113))


ifres6694 = tmp6699


}

ifres6690 = ifres6694


}

let__5109 := ifres6690
_ = let__5109

tmp6712 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5109)


if True == tmp6712 {
tmp6702 := Call1(__e, SymFn(symshen_4_5integer_6), V3106)


let__5116 := tmp6702
_ = let__5116

tmp6708 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5116)


var ifres6703 Obj

if True == tmp6708 {
tmp6704 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6703 = tmp6704


} else {
tmp6705 := Call1(__e, SymFn(symshen_4_5_1out), let__5116)


let__5117 := tmp6705
_ = let__5117

tmp6706 := Call1(__e, SymFn(symshen_4in_1_6), let__5116)


let__5118 := tmp6706
_ = let__5118

tmp6707 := Call2(__e, SymFn(symshen_4comb), let__5118, let__5117)


ifres6703 = tmp6707


}

let__5115 := ifres6703
_ = let__5115

tmp6710 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5115)


if True == tmp6710 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5115
}


} else {
return let__5109
}


} else {
return let__5103
}


panic("unreachable")
})

tmp6715 := Call2(__e, ns2_1set, symshen_4_5log10_6, tmp6676)


_ = tmp6715

tmp6716 := MakeNative1(func(__e *ControlFlow, V3123 Obj) Obj {
tmp6721 := Call2(__e, SymFn(symshen_4hds_a_2), V3123, MakeInteger(101))


var ifres6717 Obj

if True == tmp6721 {
tmp6718 := Call1(__e, SymFn(symtail), V3123)


let__5120 := tmp6718
_ = let__5120

tmp6719 := Call2(__e, SymFn(symshen_4comb), let__5120, symshen_4skip)


ifres6717 = tmp6719


} else {
tmp6720 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6717 = tmp6720


}

let__5119 := ifres6717
_ = let__5119

tmp6723 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5119)


if True == tmp6723 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5119
}


panic("unreachable")
})

tmp6724 := Call2(__e, ns2_1set, symshen_4_5lowE_6, tmp6716)


_ = tmp6724

tmp6725 := MakeNative2(func(__e *ControlFlow, V3126, V3127 Obj) Obj {
tmp6726 := Call2(__e, SymFn(symshen_4expt), MakeInteger(10), V3127)


return GMul(__e, V3126, tmp6726)


panic("unreachable")
})

tmp6727 := Call2(__e, ns2_1set, symshen_4compute_1E, tmp6725)


_ = tmp6727

tmp6728 := MakeNative1(func(__e *ControlFlow, V3128 Obj) Obj {
tmp6729 := Call1(__e, SymFn(symshen_4_5whitespace_6), V3128)


let__5122 := tmp6729
_ = let__5122

tmp6739 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5122)


var ifres6730 Obj

if True == tmp6739 {
tmp6731 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6730 = tmp6731


} else {
tmp6732 := Call1(__e, SymFn(symshen_4in_1_6), let__5122)


let__5123 := tmp6732
_ = let__5123

tmp6733 := Call1(__e, SymFn(symshen_4_5whitespaces_6), let__5123)


let__5124 := tmp6733
_ = let__5124

tmp6738 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5124)


var ifres6734 Obj

if True == tmp6738 {
tmp6735 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6734 = tmp6735


} else {
tmp6736 := Call1(__e, SymFn(symshen_4in_1_6), let__5124)


let__5125 := tmp6736
_ = let__5125

tmp6737 := Call2(__e, SymFn(symshen_4comb), let__5125, symshen_4skip)


ifres6734 = tmp6737


}

ifres6730 = ifres6734


}

let__5121 := ifres6730
_ = let__5121

tmp6749 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5121)


if True == tmp6749 {
tmp6740 := Call1(__e, SymFn(symshen_4_5whitespace_6), V3128)


let__5127 := tmp6740
_ = let__5127

tmp6745 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5127)


var ifres6741 Obj

if True == tmp6745 {
tmp6742 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6741 = tmp6742


} else {
tmp6743 := Call1(__e, SymFn(symshen_4in_1_6), let__5127)


let__5128 := tmp6743
_ = let__5128

tmp6744 := Call2(__e, SymFn(symshen_4comb), let__5128, symshen_4skip)


ifres6741 = tmp6744


}

let__5126 := ifres6741
_ = let__5126

tmp6747 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5126)


if True == tmp6747 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5126
}


} else {
return let__5121
}


panic("unreachable")
})

tmp6750 := Call2(__e, ns2_1set, symshen_4_5whitespaces_6, tmp6728)


_ = tmp6750

tmp6751 := MakeNative1(func(__e *ControlFlow, V3137 Obj) Obj {
tmp6760 := GConsP(__e, V3137)

var ifres6752 Obj

if True == tmp6760 {
tmp6753 := Call1(__e, SymFn(symhead), V3137)


let__5130 := tmp6753
_ = let__5130

tmp6754 := Call1(__e, SymFn(symtail), V3137)


let__5131 := tmp6754
_ = let__5131

tmp6758 := Call1(__e, SymFn(symshen_4whitespace_2), let__5130)


var ifres6755 Obj

if True == tmp6758 {
tmp6756 := Call2(__e, SymFn(symshen_4comb), let__5131, symshen_4skip)


ifres6755 = tmp6756


} else {
tmp6757 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6755 = tmp6757


}

ifres6752 = ifres6755


} else {
tmp6759 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6752 = tmp6759


}

let__5129 := ifres6752
_ = let__5129

tmp6762 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5129)


if True == tmp6762 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5129
}


panic("unreachable")
})

tmp6763 := Call2(__e, ns2_1set, symshen_4_5whitespace_6, tmp6751)


_ = tmp6763

tmp6764 := MakeNative1(func(__e *ControlFlow, V3143 Obj) Obj {
tmp6772 := GEq(__e, MakeInteger(32), V3143)

if True == tmp6772 {
return True
} else {
tmp6770 := GEq(__e, MakeInteger(13), V3143)

if True == tmp6770 {
return True
} else {
tmp6768 := GEq(__e, MakeInteger(10), V3143)

if True == tmp6768 {
return True
} else {
tmp6766 := GEq(__e, MakeInteger(9), V3143)

if True == tmp6766 {
return True
} else {
return False
}


}


}


}


panic("unreachable")
})

tmp6773 := Call2(__e, ns2_1set, symshen_4whitespace_2, tmp6764)


_ = tmp6773

tmp6774 := MakeNative2(func(__e *ControlFlow, V3146, V3147 Obj) Obj {
tmp6787 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V3146)

var ifres6781 Obj

if True == tmp6787 {
tmp6785 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V3146)

tmp6786 := GNot(__e, tmp6785)

var ifres6782 Obj

if True == tmp6786 {
tmp6783 := GEq(__e, V3146, sympackage)

ifres6782 = GNot(__e, tmp6783)


} else {
ifres6782 = False


}

ifres6781 = ifres6782


} else {
ifres6781 = False


}

if True == ifres6781 {
tmp6775 := GValue(__e, symshen_4_dsource_1form_1handlers_d)

tmp6776 := Call3(__e, SymFn(symshen_4update_1assoc), V3146, V3147, tmp6775)


tmp6777 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsource_1form_1handlers_d, tmp6776)

_ = tmp6777

return V3146


} else {
tmp6778 := Call3(__e, SymFn(symshen_4app), V3146, MakeString("\n"), symshen_4s)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("invalid source-form head: "), tmp6778))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("invalid source-form head: "), tmp6778)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


panic("unreachable")
})

tmp6788 := Call2(__e, ns2_1set, symshen_4register_1source_1form, tmp6774)


_ = tmp6788

tmp6789 := MakeNative1(func(__e *ControlFlow, V3148 Obj) Obj {
tmp6790 := GValue(__e, symshen_4_dsource_1form_1handlers_d)

tmp6791 := Call2(__e, SymFn(symassoc), V3148, tmp6790)


tmp6792 := GValue(__e, symshen_4_dsource_1form_1handlers_d)

tmp6793 := Call2(__e, SymFn(symremove), tmp6791, tmp6792)


tmp6794 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsource_1form_1handlers_d, tmp6793)

_ = tmp6794

return V3148


panic("unreachable")
})

tmp6795 := Call2(__e, ns2_1set, symshen_4unregister_1source_1form, tmp6789)


_ = tmp6795

tmp6796 := MakeNative2(func(__e *ControlFlow, V3155, V3156 Obj) Obj {
tmp6799 := GConsP(__e, V3155)

if True == tmp6799 {
tmp6797 := GHd(__e, V3155)

return TailCall2(__e, SymFn(symassoc), tmp6797, V3156)


} else {
return Nil
}


panic("unreachable")
})

tmp6800 := Call2(__e, ns2_1set, symshen_4source_1form_1handler, tmp6796)


_ = tmp6800

tmp6801 := MakeNative2(func(__e *ControlFlow, V3159, V3160 Obj) Obj {
tmp6809 := GConsP(__e, V3160)

if True == tmp6809 {
tmp6802 := GTl(__e, V3160)

tmp6803 := Call1(__e, tmp6802, V3159)


let__5132 := tmp6803
_ = let__5132

tmp6807 := Call1(__e, SymFn(symshen_4source_1form_1list_2), let__5132)


if True == tmp6807 {
return let__5132
} else {
tmp6804 := Call3(__e, SymFn(symshen_4app), V3159, MakeString("\n"), symshen_4s)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("source-form handler must return a list of forms: "), tmp6804))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("source-form handler must return a list of forms: "), tmp6804)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4source_1form_1expansion)
}


panic("unreachable")
})

tmp6810 := Call2(__e, ns2_1set, symshen_4source_1form_1expansion, tmp6801)


_ = tmp6810

tmp6811 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V3166 := __self1
_ = V3166
tmp6816 := GEq(__e, Nil, V3166)

if True == tmp6816 {
return True
} else {
tmp6814 := GConsP(__e, V3166)

if True == tmp6814 {
tmp6812 := GTl(__e, V3166)

if SymFn(symshen_4source_1form_1list_2) == __self {
__self1 = tmp6812
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4source_1form_1list_2), tmp6812)


} else {
return False
}


}


panic("unreachable")
})
return __self
})()

tmp6817 := Call2(__e, ns2_1set, symshen_4source_1form_1list_2, tmp6811)


_ = tmp6817

tmp6818 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V3167 := __self1
_ = V3167
tmp6852 := GEq(__e, Nil, V3167)

if True == tmp6852 {
return Nil
} else {
tmp6850 := GConsP(__e, V3167)

var ifres6847 Obj

if True == tmp6850 {
tmp6848 := GHd(__e, V3167)

tmp6849 := Call1(__e, SymFn(symshen_4packaged_2), tmp6848)


ifres6847 = tmp6849


} else {
ifres6847 = False


}

if True == ifres6847 {
tmp6819 := GHd(__e, V3167)

tmp6820 := Call1(__e, SymFn(symshen_4unpackage), tmp6819)


tmp6821 := GTl(__e, V3167)

tmp6822 := Call2(__e, SymFn(symappend), tmp6820, tmp6821)


if SymFn(symshen_4unpackage_emacroexpand) == __self {
__self1 = tmp6822
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4unpackage_emacroexpand), tmp6822)


} else {
tmp6845 := GConsP(__e, V3167)

if True == tmp6845 {
tmp6823 := GHd(__e, V3167)

tmp6824 := GValue(__e, symshen_4_dsource_1form_1handlers_d)

tmp6825 := Call2(__e, SymFn(symshen_4source_1form_1handler), tmp6823, tmp6824)


let__5133 := tmp6825
_ = let__5133

tmp6843 := GConsP(__e, let__5133)

if True == tmp6843 {
tmp6826 := GHd(__e, V3167)

tmp6827 := Call2(__e, SymFn(symshen_4source_1form_1expansion), tmp6826, let__5133)


tmp6828 := GTl(__e, V3167)

tmp6829 := Call2(__e, SymFn(symappend), tmp6827, tmp6828)


if SymFn(symshen_4unpackage_emacroexpand) == __self {
__self1 = tmp6829
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4unpackage_emacroexpand), tmp6829)


} else {
tmp6830 := GHd(__e, V3167)

tmp6831 := Call1(__e, SymFn(symmacroexpand), tmp6830)


let__5134 := tmp6831
_ = let__5134

tmp6841 := Call1(__e, SymFn(symshen_4packaged_2), let__5134)


var ifres6837 Obj

if True == tmp6841 {
ifres6837 = True


} else {
tmp6838 := GValue(__e, symshen_4_dsource_1form_1handlers_d)

tmp6839 := Call2(__e, SymFn(symshen_4source_1form_1handler), let__5134, tmp6838)


tmp6840 := GConsP(__e, tmp6839)

ifres6837 = tmp6840


}

if True == ifres6837 {
tmp6832 := GTl(__e, V3167)

tmp6833 := GCons(__e, let__5134, tmp6832)

if SymFn(symshen_4unpackage_emacroexpand) == __self {
__self1 = tmp6833
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4unpackage_emacroexpand), tmp6833)


} else {
tmp6834 := GTl(__e, V3167)

tmp6835 := Call1(__e, SymFn(symshen_4unpackage_emacroexpand), tmp6834)


return GCons(__e, let__5134, tmp6835)


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

tmp6853 := Call2(__e, ns2_1set, symshen_4unpackage_emacroexpand, tmp6818)


_ = tmp6853

tmp6854 := MakeNative1(func(__e *ControlFlow, V3172 Obj) Obj {
tmp6866 := GConsP(__e, V3172)

var ifres6856 Obj

if True == tmp6866 {
tmp6864 := GHd(__e, V3172)

tmp6865 := GEq(__e, sympackage, tmp6864)

var ifres6857 Obj

if True == tmp6865 {
tmp6862 := GTl(__e, V3172)

tmp6863 := GConsP(__e, tmp6862)

var ifres6858 Obj

if True == tmp6863 {
tmp6859 := GTl(__e, V3172)

tmp6860 := GTl(__e, tmp6859)

tmp6861 := GConsP(__e, tmp6860)

ifres6858 = tmp6861


} else {
ifres6858 = False


}

ifres6857 = ifres6858


} else {
ifres6857 = False


}

ifres6856 = ifres6857


} else {
ifres6856 = False


}

if True == ifres6856 {
return True
} else {
return False
}


panic("unreachable")
})

tmp6867 := Call2(__e, ns2_1set, symshen_4packaged_2, tmp6854)


_ = tmp6867

tmp6868 := MakeNative1(func(__e *ControlFlow, V3175 Obj) Obj {
tmp6918 := GConsP(__e, V3175)

var ifres6904 Obj

if True == tmp6918 {
tmp6916 := GHd(__e, V3175)

tmp6917 := GEq(__e, sympackage, tmp6916)

var ifres6905 Obj

if True == tmp6917 {
tmp6914 := GTl(__e, V3175)

tmp6915 := GConsP(__e, tmp6914)

var ifres6906 Obj

if True == tmp6915 {
tmp6911 := GTl(__e, V3175)

tmp6912 := GHd(__e, tmp6911)

tmp6913 := GEq(__e, symnull, tmp6912)

var ifres6907 Obj

if True == tmp6913 {
tmp6908 := GTl(__e, V3175)

tmp6909 := GTl(__e, tmp6908)

tmp6910 := GConsP(__e, tmp6909)

ifres6907 = tmp6910


} else {
ifres6907 = False


}

ifres6906 = ifres6907


} else {
ifres6906 = False


}

ifres6905 = ifres6906


} else {
ifres6905 = False


}

ifres6904 = ifres6905


} else {
ifres6904 = False


}

if True == ifres6904 {
tmp6869 := GTl(__e, V3175)

tmp6870 := GTl(__e, tmp6869)

return GTl(__e, tmp6870)


} else {
tmp6902 := GConsP(__e, V3175)

var ifres6892 Obj

if True == tmp6902 {
tmp6900 := GHd(__e, V3175)

tmp6901 := GEq(__e, sympackage, tmp6900)

var ifres6893 Obj

if True == tmp6901 {
tmp6898 := GTl(__e, V3175)

tmp6899 := GConsP(__e, tmp6898)

var ifres6894 Obj

if True == tmp6899 {
tmp6895 := GTl(__e, V3175)

tmp6896 := GTl(__e, tmp6895)

tmp6897 := GConsP(__e, tmp6896)

ifres6894 = tmp6897


} else {
ifres6894 = False


}

ifres6893 = ifres6894


} else {
ifres6893 = False


}

ifres6892 = ifres6893


} else {
ifres6892 = False


}

if True == ifres6892 {
tmp6871 := GTl(__e, V3175)

tmp6872 := GTl(__e, tmp6871)

tmp6873 := GHd(__e, tmp6872)

tmp6874 := Call1(__e, SymFn(symeval), tmp6873)


let__5135 := tmp6874
_ = let__5135

tmp6875 := GTl(__e, V3175)

tmp6876 := GHd(__e, tmp6875)

tmp6877 := PrimGuard1(__e, symstr, PrimStr, tmp6876)

tmp6878 := GTl(__e, V3175)

tmp6879 := GTl(__e, tmp6878)

tmp6880 := GTl(__e, tmp6879)

tmp6881 := Call3(__e, SymFn(symshen_4package_1symbols), tmp6877, let__5135, tmp6880)


let__5136 := tmp6881
_ = let__5136

tmp6882 := GTl(__e, V3175)

tmp6883 := GHd(__e, tmp6882)

tmp6884 := Call2(__e, SymFn(symshen_4record_1external), tmp6883, let__5135)


let__5137 := tmp6884
_ = let__5137

tmp6885 := GTl(__e, V3175)

tmp6886 := GHd(__e, tmp6885)

tmp6887 := GTl(__e, V3175)

tmp6888 := GTl(__e, tmp6887)

tmp6889 := GTl(__e, tmp6888)

tmp6890 := Call3(__e, SymFn(symshen_4record_1internal), tmp6886, let__5135, tmp6889)


let__5138 := tmp6890
_ = let__5138

return let__5136


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4unpackage)
}


}


panic("unreachable")
})

tmp6919 := Call2(__e, ns2_1set, symshen_4unpackage, tmp6868)


_ = tmp6919

tmp6920 := MakeNative3(func(__e *ControlFlow, V3180, V3181, V3182 Obj) Obj {
tmp6921 := MakeNative0(func(__e *ControlFlow) Obj {
tmp6922 := GValue(__e, sym_dproperty_1vector_d)

return TailCall3(__e, SymFn(symget), V3180, symshen_4internal_1symbols, tmp6922)


panic("unreachable")
})

tmp6923 := MakeNative1(func(__e *ControlFlow, Z3184 Obj) Obj {
return Nil
panic("unreachable")
})

tmp6924 := Call2(__e, try_1catch, tmp6921, tmp6923)


let__5139 := tmp6924
_ = let__5139

tmp6925 := PrimGuard1(__e, symstr, PrimStr, V3180)

tmp6926 := Call3(__e, SymFn(symshen_4internal_1symbols), tmp6925, V3181, V3182)


let__5140 := tmp6926
_ = let__5140

tmp6927 := Call2(__e, SymFn(symunion), let__5140, let__5139)


tmp6928 := GValue(__e, sym_dproperty_1vector_d)

return TailCall4(__e, SymFn(symput), V3180, symshen_4internal_1symbols, tmp6927, tmp6928)


panic("unreachable")
})

tmp6929 := Call2(__e, ns2_1set, symshen_4record_1internal, tmp6920)


_ = tmp6929

tmp6930 := MakeNative3(func(__e *ControlFlow, V3192, V3193, V3194 Obj) Obj {
tmp6939 := GConsP(__e, V3194)

if True == tmp6939 {
tmp6931 := GHd(__e, V3194)

tmp6932 := Call3(__e, SymFn(symshen_4internal_1symbols), V3192, V3193, tmp6931)


tmp6933 := GTl(__e, V3194)

tmp6934 := Call3(__e, SymFn(symshen_4internal_1symbols), V3192, V3193, tmp6933)


return TailCall2(__e, SymFn(symunion), tmp6932, tmp6934)


} else {
tmp6937 := Call3(__e, SymFn(symshen_4internal_2), V3194, V3192, V3193)


if True == tmp6937 {
tmp6935 := Call2(__e, SymFn(symshen_4intern_1in_1package), V3192, V3194)


return GCons(__e, tmp6935, Nil)


} else {
return Nil
}


}


panic("unreachable")
})

tmp6940 := Call2(__e, ns2_1set, symshen_4internal_1symbols, tmp6930)


_ = tmp6940

tmp6941 := MakeNative2(func(__e *ControlFlow, V3195, V3196 Obj) Obj {
tmp6942 := MakeNative0(func(__e *ControlFlow) Obj {
tmp6943 := GValue(__e, sym_dproperty_1vector_d)

return TailCall3(__e, SymFn(symget), V3195, symshen_4external_1symbols, tmp6943)


panic("unreachable")
})

tmp6944 := MakeNative1(func(__e *ControlFlow, Z3198 Obj) Obj {
return Nil
panic("unreachable")
})

tmp6945 := Call2(__e, try_1catch, tmp6942, tmp6944)


let__5141 := tmp6945
_ = let__5141

tmp6946 := Call2(__e, SymFn(symunion), V3196, let__5141)


tmp6947 := GValue(__e, sym_dproperty_1vector_d)

return TailCall4(__e, SymFn(symput), V3195, symshen_4external_1symbols, tmp6946, tmp6947)


panic("unreachable")
})

tmp6948 := Call2(__e, ns2_1set, symshen_4record_1external, tmp6941)


_ = tmp6948

tmp6949 := MakeNative3(func(__e *ControlFlow, V3203, V3204, V3205 Obj) Obj {
tmp6954 := GConsP(__e, V3205)

if True == tmp6954 {
tmp6950 := MakeNative1(func(__e *ControlFlow, Z3206 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4package_1symbols), V3203, V3204, Z3206)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp6950, V3205)


} else {
tmp6952 := Call3(__e, SymFn(symshen_4internal_2), V3205, V3203, V3204)


if True == tmp6952 {
return TailCall2(__e, SymFn(symshen_4intern_1in_1package), V3203, V3205)
} else {
return V3205
}


}


panic("unreachable")
})

tmp6955 := Call2(__e, ns2_1set, symshen_4package_1symbols, tmp6949)


_ = tmp6955

tmp6956 := MakeNative2(func(__e *ControlFlow, V3207, V3208 Obj) Obj {
tmp6957 := PrimGuard1(__e, symstr, PrimStr, V3208)

tmp6958 := Call2(__e, SymFn(sym_8s), MakeString("."), tmp6957)


tmp6959 := Call2(__e, SymFn(sym_8s), V3207, tmp6958)


return PrimGuard1(__e, symintern, PrimIntern, tmp6959)


panic("unreachable")
})

tmp6960 := Call2(__e, ns2_1set, symshen_4intern_1in_1package, tmp6956)


_ = tmp6960

tmp6961 := MakeNative3(func(__e *ControlFlow, V3209, V3210, V3211 Obj) Obj {
tmp6983 := Call2(__e, SymFn(symelement_2), V3209, V3211)


if True == GNot(__e, tmp6983) {
tmp6980 := Call1(__e, SymFn(symshen_4sng_2), V3209)


if True == GNot(__e, tmp6980) {
tmp6977 := Call1(__e, SymFn(symshen_4dbl_2), V3209)


if True == GNot(__e, tmp6977) {
tmp6975 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V3209)

if True == tmp6975 {
tmp6972 := Call1(__e, SymFn(symshen_4sysfunc_2), V3209)


if True == GNot(__e, tmp6972) {
tmp6969 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V3209)

if True == GNot(__e, tmp6969) {
tmp6965 := PrimGuard1(__e, symstr, PrimStr, V3209)

tmp6966 := Call1(__e, SymFn(symshen_4internal_1to_1shen_2), tmp6965)


if True == GNot(__e, tmp6966) {
tmp6962 := PrimGuard1(__e, symstr, PrimStr, V3209)

tmp6963 := Call2(__e, SymFn(symshen_4internal_1to_1P_2), V3210, tmp6962)


return GNot(__e, tmp6963)


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

tmp6985 := Call2(__e, ns2_1set, symshen_4internal_2, tmp6961)


_ = tmp6985

tmp6986 := MakeNative1(func(__e *ControlFlow, V3216 Obj) Obj {
tmp7031 := Call1(__e, SymFn(symshen_4_7string_2), V3216)


var ifres6988 Obj

if True == tmp7031 {
tmp7029 := Call1(__e, SymFn(symhdstr), V3216)


tmp7030 := GEq(__e, MakeString("s"), tmp7029)

var ifres6989 Obj

if True == tmp7030 {
tmp7028 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, V3216))


var ifres6990 Obj

if True == tmp7028 {
tmp7025 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, V3216))


tmp7026 := GEq(__e, MakeString("h"), tmp7025)

var ifres6991 Obj

if True == tmp7026 {
tmp7023 := Call1(__e, SymFn(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V3216)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})())


var ifres6992 Obj

if True == tmp7023 {
tmp7019 := Call1(__e, SymFn(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V3216)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})())


tmp7020 := GEq(__e, MakeString("e"), tmp7019)

var ifres6993 Obj

if True == tmp7020 {
tmp7016 := Call1(__e, SymFn(symshen_4_7string_2), (func() Obj {
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


var ifres6994 Obj

if True == tmp7016 {
tmp7011 := Call1(__e, SymFn(symhdstr), (func() Obj {
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


tmp7012 := GEq(__e, MakeString("n"), tmp7011)

var ifres6995 Obj

if True == tmp7012 {
tmp7007 := Call1(__e, SymFn(symshen_4_7string_2), (func() Obj {
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


var ifres6996 Obj

if True == tmp7007 {
tmp7001 := Call1(__e, SymFn(symhdstr), (func() Obj {
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


tmp7002 := GEq(__e, MakeString("."), tmp7001)

ifres6996 = tmp7002


} else {
ifres6996 = False


}

ifres6995 = ifres6996


} else {
ifres6995 = False


}

ifres6994 = ifres6995


} else {
ifres6994 = False


}

ifres6993 = ifres6994


} else {
ifres6993 = False


}

ifres6992 = ifres6993


} else {
ifres6992 = False


}

ifres6991 = ifres6992


} else {
ifres6991 = False


}

ifres6990 = ifres6991


} else {
ifres6990 = False


}

ifres6989 = ifres6990


} else {
ifres6989 = False


}

ifres6988 = ifres6989


} else {
ifres6988 = False


}

if True == ifres6988 {
return True
} else {
return False
}


panic("unreachable")
})

tmp7032 := Call2(__e, ns2_1set, symshen_4internal_1to_1shen_2, tmp6986)


_ = tmp7032

tmp7033 := MakeNative1(func(__e *ControlFlow, V3217 Obj) Obj {
tmp7034 := GValue(__e, sym_dproperty_1vector_d)

tmp7035 := Call3(__e, SymFn(symget), symshen, symshen_4external_1symbols, tmp7034)


return TailCall2(__e, SymFn(symelement_2), V3217, tmp7035)


panic("unreachable")
})

tmp7036 := Call2(__e, ns2_1set, symshen_4sysfunc_2, tmp7033)


_ = tmp7036

tmp7037 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V3225 := __self1
_ = V3225
V3226 := __self2
_ = V3226
tmp7054 := GEq(__e, MakeString(""), V3225)

var ifres7049 Obj

if True == tmp7054 {
tmp7053 := Call1(__e, SymFn(symshen_4_7string_2), V3226)


var ifres7050 Obj

if True == tmp7053 {
tmp7051 := Call1(__e, SymFn(symhdstr), V3226)


tmp7052 := GEq(__e, MakeString("."), tmp7051)

ifres7050 = tmp7052


} else {
ifres7050 = False


}

ifres7049 = ifres7050


} else {
ifres7049 = False


}

if True == ifres7049 {
return True
} else {
tmp7047 := Call1(__e, SymFn(symshen_4_7string_2), V3225)


var ifres7041 Obj

if True == tmp7047 {
tmp7046 := Call1(__e, SymFn(symshen_4_7string_2), V3226)


var ifres7042 Obj

if True == tmp7046 {
tmp7043 := Call1(__e, SymFn(symhdstr), V3225)


tmp7044 := Call1(__e, SymFn(symhdstr), V3226)


tmp7045 := GEq(__e, tmp7043, tmp7044)

ifres7042 = tmp7045


} else {
ifres7042 = False


}

ifres7041 = ifres7042


} else {
ifres7041 = False


}

if True == ifres7041 {
tmp7038 := PrimGuard1(__e, symtlstr, PrimTailString, V3225)

if SymFn(symshen_4internal_1to_1P_2) == __self {
__self1, __self2 = tmp7038, PrimGuard1(__e, symtlstr, PrimTailString, V3226)
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4internal_1to_1P_2), tmp7038, PrimGuard1(__e, symtlstr, PrimTailString, V3226))


} else {
return False
}


}


panic("unreachable")
})
return __self
})()

tmp7055 := Call2(__e, ns2_1set, symshen_4internal_1to_1P_2, tmp7037)


_ = tmp7055

tmp7056 := MakeNative2(func(__e *ControlFlow, V3229, V3230 Obj) Obj {
tmp7075 := Call2(__e, SymFn(symelement_2), V3229, V3230)


if True == tmp7075 {
return V3229
} else {
tmp7073 := GConsP(__e, V3229)

var ifres7070 Obj

if True == tmp7073 {
tmp7071 := GHd(__e, V3229)

tmp7072 := GEq(__e, symcond, tmp7071)

ifres7070 = tmp7072


} else {
ifres7070 = False


}

if True == ifres7070 {
tmp7057 := GTl(__e, V3229)

tmp7058 := Call2(__e, SymFn(symshen_4process_1cond_1clauses), tmp7057, V3230)


return GCons(__e, symcond, tmp7058)


} else {
tmp7068 := GConsP(__e, V3229)

var ifres7065 Obj

if True == tmp7068 {
tmp7066 := GHd(__e, V3229)

tmp7067 := Call1(__e, SymFn(symshen_4non_1application_2), tmp7066)


ifres7065 = tmp7067


} else {
ifres7065 = False


}

if True == ifres7065 {
tmp7059 := GHd(__e, V3229)

return TailCall3(__e, SymFn(symshen_4special_1case), tmp7059, V3229, V3230)


} else {
tmp7063 := GConsP(__e, V3229)

if True == tmp7063 {
tmp7060 := MakeNative1(func(__e *ControlFlow, Z3231 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4process_1applications), Z3231, V3230)
panic("unreachable")
})

tmp7061 := Call2(__e, SymFn(symmap), tmp7060, V3229)


return TailCall2(__e, SymFn(symshen_4process_1application), tmp7061, V3230)


} else {
return V3229
}


}


}


}


panic("unreachable")
})

tmp7076 := Call2(__e, ns2_1set, symshen_4process_1applications, tmp7056)


_ = tmp7076

tmp7077 := MakeNative1(func(__e *ControlFlow, V3234 Obj) Obj {
tmp7087 := GEq(__e, symdefine, V3234)

if True == tmp7087 {
return True
} else {
tmp7085 := GEq(__e, symdefun, V3234)

if True == tmp7085 {
return True
} else {
tmp7083 := GEq(__e, symsynonyms, V3234)

if True == tmp7083 {
return True
} else {
tmp7081 := Call1(__e, SymFn(symshen_4special_2), V3234)


if True == tmp7081 {
return True
} else {
tmp7079 := Call1(__e, SymFn(symshen_4extraspecial_2), V3234)


if True == tmp7079 {
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

tmp7088 := Call2(__e, ns2_1set, symshen_4non_1application_2, tmp7077)


_ = tmp7088

tmp7089 := MakeNative3(func(__e *ControlFlow, V3239, V3240, V3241 Obj) Obj {
tmp7294 := GEq(__e, symlambda, V3239)

var ifres7277 Obj

if True == tmp7294 {
tmp7293 := GConsP(__e, V3240)

var ifres7278 Obj

if True == tmp7293 {
tmp7291 := GHd(__e, V3240)

tmp7292 := GEq(__e, symlambda, tmp7291)

var ifres7279 Obj

if True == tmp7292 {
tmp7289 := GTl(__e, V3240)

tmp7290 := GConsP(__e, tmp7289)

var ifres7280 Obj

if True == tmp7290 {
tmp7286 := GTl(__e, V3240)

tmp7287 := GTl(__e, tmp7286)

tmp7288 := GConsP(__e, tmp7287)

var ifres7281 Obj

if True == tmp7288 {
tmp7282 := GTl(__e, V3240)

tmp7283 := GTl(__e, tmp7282)

tmp7284 := GTl(__e, tmp7283)

tmp7285 := GEq(__e, Nil, tmp7284)

ifres7281 = tmp7285


} else {
ifres7281 = False


}

ifres7280 = ifres7281


} else {
ifres7280 = False


}

ifres7279 = ifres7280


} else {
ifres7279 = False


}

ifres7278 = ifres7279


} else {
ifres7278 = False


}

ifres7277 = ifres7278


} else {
ifres7277 = False


}

if True == ifres7277 {
tmp7090 := GTl(__e, V3240)

tmp7091 := GHd(__e, tmp7090)

tmp7092 := GTl(__e, V3240)

tmp7093 := GTl(__e, tmp7092)

tmp7094 := GHd(__e, tmp7093)

tmp7095 := Call2(__e, SymFn(symshen_4process_1applications), tmp7094, V3241)


tmp7096 := GCons(__e, tmp7095, Nil)

tmp7097 := GCons(__e, tmp7091, tmp7096)

return GCons(__e, symlambda, tmp7097)


} else {
tmp7275 := GEq(__e, symlet, V3239)

var ifres7252 Obj

if True == tmp7275 {
tmp7274 := GConsP(__e, V3240)

var ifres7253 Obj

if True == tmp7274 {
tmp7272 := GHd(__e, V3240)

tmp7273 := GEq(__e, symlet, tmp7272)

var ifres7254 Obj

if True == tmp7273 {
tmp7270 := GTl(__e, V3240)

tmp7271 := GConsP(__e, tmp7270)

var ifres7255 Obj

if True == tmp7271 {
tmp7267 := GTl(__e, V3240)

tmp7268 := GTl(__e, tmp7267)

tmp7269 := GConsP(__e, tmp7268)

var ifres7256 Obj

if True == tmp7269 {
tmp7263 := GTl(__e, V3240)

tmp7264 := GTl(__e, tmp7263)

tmp7265 := GTl(__e, tmp7264)

tmp7266 := GConsP(__e, tmp7265)

var ifres7257 Obj

if True == tmp7266 {
tmp7258 := GTl(__e, V3240)

tmp7259 := GTl(__e, tmp7258)

tmp7260 := GTl(__e, tmp7259)

tmp7261 := GTl(__e, tmp7260)

tmp7262 := GEq(__e, Nil, tmp7261)

ifres7257 = tmp7262


} else {
ifres7257 = False


}

ifres7256 = ifres7257


} else {
ifres7256 = False


}

ifres7255 = ifres7256


} else {
ifres7255 = False


}

ifres7254 = ifres7255


} else {
ifres7254 = False


}

ifres7253 = ifres7254


} else {
ifres7253 = False


}

ifres7252 = ifres7253


} else {
ifres7252 = False


}

if True == ifres7252 {
tmp7098 := GTl(__e, V3240)

tmp7099 := GHd(__e, tmp7098)

tmp7100 := GTl(__e, V3240)

tmp7101 := GTl(__e, tmp7100)

tmp7102 := GHd(__e, tmp7101)

tmp7103 := Call2(__e, SymFn(symshen_4process_1applications), tmp7102, V3241)


tmp7104 := GTl(__e, V3240)

tmp7105 := GTl(__e, tmp7104)

tmp7106 := GTl(__e, tmp7105)

tmp7107 := GHd(__e, tmp7106)

tmp7108 := Call2(__e, SymFn(symshen_4process_1applications), tmp7107, V3241)


tmp7109 := GCons(__e, tmp7108, Nil)

tmp7110 := GCons(__e, tmp7103, tmp7109)

tmp7111 := GCons(__e, tmp7099, tmp7110)

return GCons(__e, symlet, tmp7111)


} else {
tmp7250 := GEq(__e, symdefun, V3239)

var ifres7227 Obj

if True == tmp7250 {
tmp7249 := GConsP(__e, V3240)

var ifres7228 Obj

if True == tmp7249 {
tmp7247 := GHd(__e, V3240)

tmp7248 := GEq(__e, symdefun, tmp7247)

var ifres7229 Obj

if True == tmp7248 {
tmp7245 := GTl(__e, V3240)

tmp7246 := GConsP(__e, tmp7245)

var ifres7230 Obj

if True == tmp7246 {
tmp7242 := GTl(__e, V3240)

tmp7243 := GTl(__e, tmp7242)

tmp7244 := GConsP(__e, tmp7243)

var ifres7231 Obj

if True == tmp7244 {
tmp7238 := GTl(__e, V3240)

tmp7239 := GTl(__e, tmp7238)

tmp7240 := GTl(__e, tmp7239)

tmp7241 := GConsP(__e, tmp7240)

var ifres7232 Obj

if True == tmp7241 {
tmp7233 := GTl(__e, V3240)

tmp7234 := GTl(__e, tmp7233)

tmp7235 := GTl(__e, tmp7234)

tmp7236 := GTl(__e, tmp7235)

tmp7237 := GEq(__e, Nil, tmp7236)

ifres7232 = tmp7237


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

ifres7227 = ifres7228


} else {
ifres7227 = False


}

if True == ifres7227 {
return V3240
} else {
tmp7225 := GEq(__e, symdefine, V3239)

var ifres7208 Obj

if True == tmp7225 {
tmp7224 := GConsP(__e, V3240)

var ifres7209 Obj

if True == tmp7224 {
tmp7222 := GHd(__e, V3240)

tmp7223 := GEq(__e, symdefine, tmp7222)

var ifres7210 Obj

if True == tmp7223 {
tmp7220 := GTl(__e, V3240)

tmp7221 := GConsP(__e, tmp7220)

var ifres7211 Obj

if True == tmp7221 {
tmp7217 := GTl(__e, V3240)

tmp7218 := GTl(__e, tmp7217)

tmp7219 := GConsP(__e, tmp7218)

var ifres7212 Obj

if True == tmp7219 {
tmp7213 := GTl(__e, V3240)

tmp7214 := GTl(__e, tmp7213)

tmp7215 := GHd(__e, tmp7214)

tmp7216 := GEq(__e, sym_i, tmp7215)

ifres7212 = tmp7216


} else {
ifres7212 = False


}

ifres7211 = ifres7212


} else {
ifres7211 = False


}

ifres7210 = ifres7211


} else {
ifres7210 = False


}

ifres7209 = ifres7210


} else {
ifres7209 = False


}

ifres7208 = ifres7209


} else {
ifres7208 = False


}

if True == ifres7208 {
tmp7112 := GTl(__e, V3240)

tmp7113 := GHd(__e, tmp7112)

tmp7114 := GTl(__e, V3240)

tmp7115 := GHd(__e, tmp7114)

tmp7116 := GTl(__e, V3240)

tmp7117 := GTl(__e, tmp7116)

tmp7118 := GTl(__e, tmp7117)

tmp7119 := Call3(__e, SymFn(symshen_4process_1after_1type), tmp7115, tmp7118, V3241)


tmp7120 := GCons(__e, sym_i, tmp7119)

tmp7121 := GCons(__e, tmp7113, tmp7120)

return GCons(__e, symdefine, tmp7121)


} else {
tmp7206 := GEq(__e, symdefine, V3239)

var ifres7198 Obj

if True == tmp7206 {
tmp7205 := GConsP(__e, V3240)

var ifres7199 Obj

if True == tmp7205 {
tmp7203 := GHd(__e, V3240)

tmp7204 := GEq(__e, symdefine, tmp7203)

var ifres7200 Obj

if True == tmp7204 {
tmp7201 := GTl(__e, V3240)

tmp7202 := GConsP(__e, tmp7201)

ifres7200 = tmp7202


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

if True == ifres7198 {
tmp7122 := GTl(__e, V3240)

tmp7123 := GHd(__e, tmp7122)

tmp7124 := MakeNative1(func(__e *ControlFlow, Z3242 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4process_1applications), Z3242, V3241)
panic("unreachable")
})

tmp7125 := GTl(__e, V3240)

tmp7126 := GTl(__e, tmp7125)

tmp7127 := Call2(__e, SymFn(symmap), tmp7124, tmp7126)


tmp7128 := GCons(__e, tmp7123, tmp7127)

return GCons(__e, symdefine, tmp7128)


} else {
tmp7196 := GEq(__e, symsynonyms, V3239)

if True == tmp7196 {
return GCons(__e, symsynonyms, V3240)
} else {
tmp7194 := GEq(__e, symtype, V3239)

var ifres7177 Obj

if True == tmp7194 {
tmp7193 := GConsP(__e, V3240)

var ifres7178 Obj

if True == tmp7193 {
tmp7191 := GHd(__e, V3240)

tmp7192 := GEq(__e, symtype, tmp7191)

var ifres7179 Obj

if True == tmp7192 {
tmp7189 := GTl(__e, V3240)

tmp7190 := GConsP(__e, tmp7189)

var ifres7180 Obj

if True == tmp7190 {
tmp7186 := GTl(__e, V3240)

tmp7187 := GTl(__e, tmp7186)

tmp7188 := GConsP(__e, tmp7187)

var ifres7181 Obj

if True == tmp7188 {
tmp7182 := GTl(__e, V3240)

tmp7183 := GTl(__e, tmp7182)

tmp7184 := GTl(__e, tmp7183)

tmp7185 := GEq(__e, Nil, tmp7184)

ifres7181 = tmp7185


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

ifres7177 = ifres7178


} else {
ifres7177 = False


}

if True == ifres7177 {
tmp7129 := GTl(__e, V3240)

tmp7130 := GHd(__e, tmp7129)

tmp7131 := Call2(__e, SymFn(symshen_4process_1applications), tmp7130, V3241)


tmp7132 := GTl(__e, V3240)

tmp7133 := GTl(__e, tmp7132)

tmp7134 := GCons(__e, tmp7131, tmp7133)

return GCons(__e, symtype, tmp7134)


} else {
tmp7175 := GEq(__e, syminput_7, V3239)

var ifres7158 Obj

if True == tmp7175 {
tmp7174 := GConsP(__e, V3240)

var ifres7159 Obj

if True == tmp7174 {
tmp7172 := GHd(__e, V3240)

tmp7173 := GEq(__e, syminput_7, tmp7172)

var ifres7160 Obj

if True == tmp7173 {
tmp7170 := GTl(__e, V3240)

tmp7171 := GConsP(__e, tmp7170)

var ifres7161 Obj

if True == tmp7171 {
tmp7167 := GTl(__e, V3240)

tmp7168 := GTl(__e, tmp7167)

tmp7169 := GConsP(__e, tmp7168)

var ifres7162 Obj

if True == tmp7169 {
tmp7163 := GTl(__e, V3240)

tmp7164 := GTl(__e, tmp7163)

tmp7165 := GTl(__e, tmp7164)

tmp7166 := GEq(__e, Nil, tmp7165)

ifres7162 = tmp7166


} else {
ifres7162 = False


}

ifres7161 = ifres7162


} else {
ifres7161 = False


}

ifres7160 = ifres7161


} else {
ifres7160 = False


}

ifres7159 = ifres7160


} else {
ifres7159 = False


}

ifres7158 = ifres7159


} else {
ifres7158 = False


}

if True == ifres7158 {
tmp7135 := GTl(__e, V3240)

tmp7136 := GHd(__e, tmp7135)

tmp7137 := GTl(__e, V3240)

tmp7138 := GTl(__e, tmp7137)

tmp7139 := GHd(__e, tmp7138)

tmp7140 := Call2(__e, SymFn(symshen_4process_1applications), tmp7139, V3241)


tmp7141 := GCons(__e, tmp7140, Nil)

tmp7142 := GCons(__e, tmp7136, tmp7141)

return GCons(__e, syminput_7, tmp7142)


} else {
tmp7156 := GConsP(__e, V3240)

var ifres7153 Obj

if True == tmp7156 {
tmp7154 := GHd(__e, V3240)

tmp7155 := Call1(__e, SymFn(symshen_4special_2), tmp7154)


ifres7153 = tmp7155


} else {
ifres7153 = False


}

if True == ifres7153 {
tmp7143 := GHd(__e, V3240)

tmp7144 := MakeNative1(func(__e *ControlFlow, Z3243 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4process_1applications), Z3243, V3241)
panic("unreachable")
})

tmp7145 := GTl(__e, V3240)

tmp7146 := Call2(__e, SymFn(symmap), tmp7144, tmp7145)


return GCons(__e, tmp7143, tmp7146)


} else {
tmp7151 := GConsP(__e, V3240)

var ifres7148 Obj

if True == tmp7151 {
tmp7149 := GHd(__e, V3240)

tmp7150 := Call1(__e, SymFn(symshen_4extraspecial_2), tmp7149)


ifres7148 = tmp7150


} else {
ifres7148 = False


}

if True == ifres7148 {
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

tmp7295 := Call2(__e, ns2_1set, symshen_4special_1case, tmp7089)


_ = tmp7295

tmp7296 := MakeNative2(func(__e *ControlFlow, V3246, V3247 Obj) Obj {
tmp7323 := GEq(__e, Nil, V3246)

if True == tmp7323 {
return Nil
} else {
tmp7321 := GConsP(__e, V3246)

var ifres7309 Obj

if True == tmp7321 {
tmp7319 := GHd(__e, V3246)

tmp7320 := GConsP(__e, tmp7319)

var ifres7310 Obj

if True == tmp7320 {
tmp7316 := GHd(__e, V3246)

tmp7317 := GTl(__e, tmp7316)

tmp7318 := GConsP(__e, tmp7317)

var ifres7311 Obj

if True == tmp7318 {
tmp7312 := GHd(__e, V3246)

tmp7313 := GTl(__e, tmp7312)

tmp7314 := GTl(__e, tmp7313)

tmp7315 := GEq(__e, Nil, tmp7314)

ifres7311 = tmp7315


} else {
ifres7311 = False


}

ifres7310 = ifres7311


} else {
ifres7310 = False


}

ifres7309 = ifres7310


} else {
ifres7309 = False


}

if True == ifres7309 {
tmp7297 := GHd(__e, V3246)

tmp7298 := GHd(__e, tmp7297)

tmp7299 := Call2(__e, SymFn(symshen_4process_1applications), tmp7298, V3247)


tmp7300 := GHd(__e, V3246)

tmp7301 := GTl(__e, tmp7300)

tmp7302 := GHd(__e, tmp7301)

tmp7303 := Call2(__e, SymFn(symshen_4process_1applications), tmp7302, V3247)


tmp7304 := GCons(__e, tmp7303, Nil)

tmp7305 := GCons(__e, tmp7299, tmp7304)

tmp7306 := GTl(__e, V3246)

tmp7307 := Call2(__e, SymFn(symshen_4process_1cond_1clauses), tmp7306, V3247)


return GCons(__e, tmp7305, tmp7307)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4process_1cond_1clauses)
}


}


panic("unreachable")
})

tmp7324 := Call2(__e, ns2_1set, symshen_4process_1cond_1clauses, tmp7296)


_ = tmp7324

tmp7325 := MakeNative3(func(__e *ControlFlow, V3250, V3251, V3252 Obj) Obj {
tmp7340 := GConsP(__e, V3251)

var ifres7337 Obj

if True == tmp7340 {
tmp7338 := GHd(__e, V3251)

tmp7339 := GEq(__e, sym_j, tmp7338)

ifres7337 = tmp7339


} else {
ifres7337 = False


}

if True == ifres7337 {
tmp7326 := MakeNative1(func(__e *ControlFlow, Z3253 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4process_1applications), Z3253, V3252)
panic("unreachable")
})

tmp7327 := GTl(__e, V3251)

tmp7328 := Call2(__e, SymFn(symmap), tmp7326, tmp7327)


return GCons(__e, sym_j, tmp7328)


} else {
tmp7335 := GConsP(__e, V3251)

if True == tmp7335 {
tmp7329 := GHd(__e, V3251)

tmp7330 := GTl(__e, V3251)

tmp7331 := Call3(__e, SymFn(symshen_4process_1after_1type), V3250, tmp7330, V3252)


return GCons(__e, tmp7329, tmp7331)


} else {
tmp7332 := Call3(__e, SymFn(symshen_4app), V3250, MakeString("\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("missing } in "), tmp7332))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("missing } in "), tmp7332)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


}


panic("unreachable")
})

tmp7341 := Call2(__e, ns2_1set, symshen_4process_1after_1type, tmp7325)


_ = tmp7341

tmp7342 := MakeNative2(func(__e *ControlFlow, V3254, V3255 Obj) Obj {
tmp7385 := GConsP(__e, V3254)

if True == tmp7385 {
tmp7343 := GHd(__e, V3254)

tmp7344 := Call1(__e, SymFn(symarity), tmp7343)


let__5142 := tmp7344
_ = let__5142

tmp7345 := GTl(__e, V3254)

tmp7346 := Call1(__e, SymFn(symlength), tmp7345)


let__5143 := tmp7346
_ = let__5143

tmp7383 := Call2(__e, SymFn(symelement_2), V3254, V3255)


if True == tmp7383 {
return V3254
} else {
tmp7380 := GHd(__e, V3254)

tmp7381 := Call1(__e, SymFn(symshen_4shen_1call_2), tmp7380)


if True == tmp7381 {
return V3254
} else {
tmp7378 := Call1(__e, SymFn(symshen_4foreign_2), V3254)


if True == tmp7378 {
return TailCall1(__e, SymFn(symshen_4unpack_1foreign), V3254)
} else {
tmp7376 := Call1(__e, SymFn(symshen_4fn_1call_2), V3254)


if True == tmp7376 {
return TailCall1(__e, SymFn(symshen_4fn_1call), V3254)
} else {
tmp7374 := Call1(__e, SymFn(symshen_4zero_1place_2), V3254)


if True == tmp7374 {
return V3254
} else {
tmp7371 := GHd(__e, V3254)

tmp7372 := Call2(__e, SymFn(symshen_4undefined_1f_2), tmp7371, let__5142)


if True == tmp7372 {
tmp7347 := GHd(__e, V3254)

tmp7348 := GCons(__e, tmp7347, Nil)

tmp7349 := GCons(__e, symfn, tmp7348)

tmp7350 := GTl(__e, V3254)

tmp7351 := GCons(__e, tmp7349, tmp7350)

return TailCall1(__e, SymFn(symshen_4simple_1curry), tmp7351)


} else {
tmp7368 := GHd(__e, V3254)

tmp7369 := PrimGuard1(__e, symvariable_2, PrimIsVariable, tmp7368)

if True == tmp7369 {
return TailCall1(__e, SymFn(symshen_4simple_1curry), V3254)
} else {
tmp7365 := GHd(__e, V3254)

tmp7366 := Call1(__e, SymFn(symshen_4application_2), tmp7365)


if True == tmp7366 {
return TailCall1(__e, SymFn(symshen_4simple_1curry), V3254)
} else {
tmp7362 := GHd(__e, V3254)

tmp7363 := Call3(__e, SymFn(symshen_4partial_1application_d_2), tmp7362, let__5142, let__5143)


if True == tmp7363 {
return TailCall2(__e, SymFn(symshen_4lambda_1function), V3254, GSub(__e, let__5142, let__5143))


} else {
tmp7359 := GHd(__e, V3254)

tmp7360 := Call3(__e, SymFn(symshen_4overapplication_2), tmp7359, let__5142, let__5143)


if True == tmp7360 {
tmp7353 := GHd(__e, V3254)

tmp7354 := GCons(__e, tmp7353, Nil)

tmp7355 := GCons(__e, symfn, tmp7354)

tmp7356 := GTl(__e, V3254)

tmp7357 := GCons(__e, tmp7355, tmp7356)

return TailCall1(__e, SymFn(symshen_4simple_1curry), tmp7357)


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

tmp7386 := Call2(__e, ns2_1set, symshen_4process_1application, tmp7342)


_ = tmp7386

tmp7387 := MakeNative1(func(__e *ControlFlow, V3258 Obj) Obj {
tmp7409 := GConsP(__e, V3258)

var ifres7393 Obj

if True == tmp7409 {
tmp7407 := GHd(__e, V3258)

tmp7408 := GConsP(__e, tmp7407)

var ifres7394 Obj

if True == tmp7408 {
tmp7404 := GHd(__e, V3258)

tmp7405 := GHd(__e, tmp7404)

tmp7406 := GEq(__e, symforeign, tmp7405)

var ifres7395 Obj

if True == tmp7406 {
tmp7401 := GHd(__e, V3258)

tmp7402 := GTl(__e, tmp7401)

tmp7403 := GConsP(__e, tmp7402)

var ifres7396 Obj

if True == tmp7403 {
tmp7397 := GHd(__e, V3258)

tmp7398 := GTl(__e, tmp7397)

tmp7399 := GTl(__e, tmp7398)

tmp7400 := GEq(__e, Nil, tmp7399)

ifres7396 = tmp7400


} else {
ifres7396 = False


}

ifres7395 = ifres7396


} else {
ifres7395 = False


}

ifres7394 = ifres7395


} else {
ifres7394 = False


}

ifres7393 = ifres7394


} else {
ifres7393 = False


}

if True == ifres7393 {
tmp7388 := GHd(__e, V3258)

tmp7389 := GTl(__e, tmp7388)

tmp7390 := GHd(__e, tmp7389)

tmp7391 := GTl(__e, V3258)

return GCons(__e, tmp7390, tmp7391)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4unpack_1foreign)
}


panic("unreachable")
})

tmp7410 := Call2(__e, ns2_1set, symshen_4unpack_1foreign, tmp7387)


_ = tmp7410

tmp7411 := MakeNative1(func(__e *ControlFlow, V3261 Obj) Obj {
tmp7429 := GConsP(__e, V3261)

var ifres7413 Obj

if True == tmp7429 {
tmp7427 := GHd(__e, V3261)

tmp7428 := GConsP(__e, tmp7427)

var ifres7414 Obj

if True == tmp7428 {
tmp7424 := GHd(__e, V3261)

tmp7425 := GHd(__e, tmp7424)

tmp7426 := GEq(__e, symforeign, tmp7425)

var ifres7415 Obj

if True == tmp7426 {
tmp7421 := GHd(__e, V3261)

tmp7422 := GTl(__e, tmp7421)

tmp7423 := GConsP(__e, tmp7422)

var ifres7416 Obj

if True == tmp7423 {
tmp7417 := GHd(__e, V3261)

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
return True
} else {
return False
}


panic("unreachable")
})

tmp7430 := Call2(__e, ns2_1set, symshen_4foreign_2, tmp7411)


_ = tmp7430

tmp7431 := MakeNative1(func(__e *ControlFlow, V3264 Obj) Obj {
tmp7436 := GConsP(__e, V3264)

var ifres7433 Obj

if True == tmp7436 {
tmp7434 := GTl(__e, V3264)

tmp7435 := GEq(__e, Nil, tmp7434)

ifres7433 = tmp7435


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

tmp7437 := Call2(__e, ns2_1set, symshen_4zero_1place_2, tmp7431)


_ = tmp7437

tmp7438 := MakeNative1(func(__e *ControlFlow, V3265 Obj) Obj {
tmp7441 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V3265)

if True == tmp7441 {
tmp7439 := PrimGuard1(__e, symstr, PrimStr, V3265)

return TailCall1(__e, SymFn(symshen_4internal_1to_1shen_2), tmp7439)


} else {
return False
}


panic("unreachable")
})

tmp7442 := Call2(__e, ns2_1set, symshen_4shen_1call_2, tmp7438)


_ = tmp7442

tmp7443 := MakeNative1(func(__e *ControlFlow, V3270 Obj) Obj {
tmp7467 := GConsP(__e, V3270)

var ifres7457 Obj

if True == tmp7467 {
tmp7465 := GHd(__e, V3270)

tmp7466 := GEq(__e, symprotect, tmp7465)

var ifres7458 Obj

if True == tmp7466 {
tmp7463 := GTl(__e, V3270)

tmp7464 := GConsP(__e, tmp7463)

var ifres7459 Obj

if True == tmp7464 {
tmp7460 := GTl(__e, V3270)

tmp7461 := GTl(__e, tmp7460)

tmp7462 := GEq(__e, Nil, tmp7461)

ifres7459 = tmp7462


} else {
ifres7459 = False


}

ifres7458 = ifres7459


} else {
ifres7458 = False


}

ifres7457 = ifres7458


} else {
ifres7457 = False


}

if True == ifres7457 {
return False
} else {
tmp7455 := GConsP(__e, V3270)

var ifres7445 Obj

if True == tmp7455 {
tmp7453 := GHd(__e, V3270)

tmp7454 := GEq(__e, symforeign, tmp7453)

var ifres7446 Obj

if True == tmp7454 {
tmp7451 := GTl(__e, V3270)

tmp7452 := GConsP(__e, tmp7451)

var ifres7447 Obj

if True == tmp7452 {
tmp7448 := GTl(__e, V3270)

tmp7449 := GTl(__e, tmp7448)

tmp7450 := GEq(__e, Nil, tmp7449)

ifres7447 = tmp7450


} else {
ifres7447 = False


}

ifres7446 = ifres7447


} else {
ifres7446 = False


}

ifres7445 = ifres7446


} else {
ifres7445 = False


}

if True == ifres7445 {
return False
} else {
return GConsP(__e, V3270)
}


}


panic("unreachable")
})

tmp7468 := Call2(__e, ns2_1set, symshen_4application_2, tmp7443)


_ = tmp7468

tmp7469 := MakeNative2(func(__e *ControlFlow, V3275, V3276 Obj) Obj {
tmp7475 := GEq(__e, MakeInteger(-1), V3276)

if True == tmp7475 {
tmp7473 := Call1(__e, SymFn(symshen_4lowercase_1symbol_2), V3275)


if True == tmp7473 {
tmp7470 := Call1(__e, SymFn(symexternal), symshen)


tmp7471 := Call2(__e, SymFn(symelement_2), V3275, tmp7470)


return GNot(__e, tmp7471)


} else {
return False
}


} else {
return False
}


panic("unreachable")
})

tmp7476 := Call2(__e, ns2_1set, symshen_4undefined_1f_2, tmp7469)


_ = tmp7476

tmp7477 := MakeNative1(func(__e *ControlFlow, V3277 Obj) Obj {
tmp7480 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V3277)

if True == tmp7480 {
tmp7478 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V3277)

return GNot(__e, tmp7478)


} else {
return False
}


panic("unreachable")
})

tmp7481 := Call2(__e, ns2_1set, symshen_4lowercase_1symbol_2, tmp7477)


_ = tmp7481

tmp7482 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V3278 := __self1
_ = V3278
tmp7508 := GConsP(__e, V3278)

var ifres7501 Obj

if True == tmp7508 {
tmp7506 := GTl(__e, V3278)

tmp7507 := GConsP(__e, tmp7506)

var ifres7502 Obj

if True == tmp7507 {
tmp7503 := GTl(__e, V3278)

tmp7504 := GTl(__e, tmp7503)

tmp7505 := GEq(__e, Nil, tmp7504)

ifres7502 = tmp7505


} else {
ifres7502 = False


}

ifres7501 = ifres7502


} else {
ifres7501 = False


}

if True == ifres7501 {
return V3278
} else {
tmp7499 := GConsP(__e, V3278)

var ifres7492 Obj

if True == tmp7499 {
tmp7497 := GTl(__e, V3278)

tmp7498 := GConsP(__e, tmp7497)

var ifres7493 Obj

if True == tmp7498 {
tmp7494 := GTl(__e, V3278)

tmp7495 := GTl(__e, tmp7494)

tmp7496 := GConsP(__e, tmp7495)

ifres7493 = tmp7496


} else {
ifres7493 = False


}

ifres7492 = ifres7493


} else {
ifres7492 = False


}

if True == ifres7492 {
tmp7483 := GHd(__e, V3278)

tmp7484 := GTl(__e, V3278)

tmp7485 := GHd(__e, tmp7484)

tmp7486 := GCons(__e, tmp7485, Nil)

tmp7487 := GCons(__e, tmp7483, tmp7486)

tmp7488 := GTl(__e, V3278)

tmp7489 := GTl(__e, tmp7488)

tmp7490 := GCons(__e, tmp7487, tmp7489)

if SymFn(symshen_4simple_1curry) == __self {
__self1 = tmp7490
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4simple_1curry), tmp7490)


} else {
return V3278
}


}


panic("unreachable")
})
return __self
})()

tmp7509 := Call2(__e, ns2_1set, symshen_4simple_1curry, tmp7482)


_ = tmp7509

tmp7510 := MakeNative1(func(__e *ControlFlow, V3279 Obj) Obj {
return TailCall1(__e, SymFn(symfn), V3279)
panic("unreachable")
})

tmp7511 := Call2(__e, ns2_1set, symfunction, tmp7510)


_ = tmp7511

tmp7512 := MakeNative1(func(__e *ControlFlow, V3280 Obj) Obj {
tmp7519 := Call1(__e, SymFn(symarity), V3280)


tmp7520 := GEq(__e, tmp7519, MakeInteger(0))

if True == tmp7520 {
return TailCall0(__e, V3280)
} else {
tmp7513 := MakeNative0(func(__e *ControlFlow) Obj {
tmp7514 := GValue(__e, sym_dproperty_1vector_d)

return TailCall3(__e, SymFn(symget), V3280, symshen_4lambda_1form, tmp7514)


panic("unreachable")
})

tmp7515 := MakeNative1(func(__e *ControlFlow, Z3281 Obj) Obj {
tmp7516 := Call3(__e, SymFn(symshen_4app), V3280, MakeString(" is undefined\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("fn: "), tmp7516))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("fn: "), tmp7516)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp7513, tmp7515)


}


panic("unreachable")
})

tmp7521 := Call2(__e, ns2_1set, symfn, tmp7512)


_ = tmp7521

tmp7522 := MakeNative1(func(__e *ControlFlow, V3284 Obj) Obj {
tmp7546 := GConsP(__e, V3284)

var ifres7536 Obj

if True == tmp7546 {
tmp7544 := GHd(__e, V3284)

tmp7545 := GEq(__e, symfn, tmp7544)

var ifres7537 Obj

if True == tmp7545 {
tmp7542 := GTl(__e, V3284)

tmp7543 := GConsP(__e, tmp7542)

var ifres7538 Obj

if True == tmp7543 {
tmp7539 := GTl(__e, V3284)

tmp7540 := GTl(__e, tmp7539)

tmp7541 := GEq(__e, Nil, tmp7540)

ifres7538 = tmp7541


} else {
ifres7538 = False


}

ifres7537 = ifres7538


} else {
ifres7537 = False


}

ifres7536 = ifres7537


} else {
ifres7536 = False


}

if True == ifres7536 {
return True
} else {
tmp7534 := GConsP(__e, V3284)

var ifres7524 Obj

if True == tmp7534 {
tmp7532 := GHd(__e, V3284)

tmp7533 := GEq(__e, symfunction, tmp7532)

var ifres7525 Obj

if True == tmp7533 {
tmp7530 := GTl(__e, V3284)

tmp7531 := GConsP(__e, tmp7530)

var ifres7526 Obj

if True == tmp7531 {
tmp7527 := GTl(__e, V3284)

tmp7528 := GTl(__e, tmp7527)

tmp7529 := GEq(__e, Nil, tmp7528)

ifres7526 = tmp7529


} else {
ifres7526 = False


}

ifres7525 = ifres7526


} else {
ifres7525 = False


}

ifres7524 = ifres7525


} else {
ifres7524 = False


}

if True == ifres7524 {
return True
} else {
return False
}


}


panic("unreachable")
})

tmp7547 := Call2(__e, ns2_1set, symshen_4fn_1call_2, tmp7522)


_ = tmp7547

tmp7548 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V3285 := __self1
_ = V3285
tmp7582 := GConsP(__e, V3285)

var ifres7572 Obj

if True == tmp7582 {
tmp7580 := GHd(__e, V3285)

tmp7581 := GEq(__e, symfunction, tmp7580)

var ifres7573 Obj

if True == tmp7581 {
tmp7578 := GTl(__e, V3285)

tmp7579 := GConsP(__e, tmp7578)

var ifres7574 Obj

if True == tmp7579 {
tmp7575 := GTl(__e, V3285)

tmp7576 := GTl(__e, tmp7575)

tmp7577 := GEq(__e, Nil, tmp7576)

ifres7574 = tmp7577


} else {
ifres7574 = False


}

ifres7573 = ifres7574


} else {
ifres7573 = False


}

ifres7572 = ifres7573


} else {
ifres7572 = False


}

if True == ifres7572 {
tmp7549 := GTl(__e, V3285)

tmp7550 := GCons(__e, symfn, tmp7549)

if SymFn(symshen_4fn_1call) == __self {
__self1 = tmp7550
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4fn_1call), tmp7550)


} else {
tmp7570 := GConsP(__e, V3285)

var ifres7560 Obj

if True == tmp7570 {
tmp7568 := GHd(__e, V3285)

tmp7569 := GEq(__e, symfn, tmp7568)

var ifres7561 Obj

if True == tmp7569 {
tmp7566 := GTl(__e, V3285)

tmp7567 := GConsP(__e, tmp7566)

var ifres7562 Obj

if True == tmp7567 {
tmp7563 := GTl(__e, V3285)

tmp7564 := GTl(__e, tmp7563)

tmp7565 := GEq(__e, Nil, tmp7564)

ifres7562 = tmp7565


} else {
ifres7562 = False


}

ifres7561 = ifres7562


} else {
ifres7561 = False


}

ifres7560 = ifres7561


} else {
ifres7560 = False


}

if True == ifres7560 {
tmp7551 := GTl(__e, V3285)

tmp7552 := GHd(__e, tmp7551)

tmp7553 := Call1(__e, SymFn(symarity), tmp7552)


let__5144 := tmp7553
_ = let__5144

tmp7558 := GEq(__e, let__5144, MakeInteger(-1))

if True == tmp7558 {
return V3285
} else {
tmp7556 := GEq(__e, let__5144, MakeInteger(0))

if True == tmp7556 {
return GTl(__e, V3285)
} else {
tmp7554 := GTl(__e, V3285)

return TailCall2(__e, SymFn(symshen_4lambda_1function), tmp7554, let__5144)


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

tmp7583 := Call2(__e, ns2_1set, symshen_4fn_1call, tmp7548)


_ = tmp7583

tmp7584 := MakeNative3(func(__e *ControlFlow, V3287, V3288, V3289 Obj) Obj {
let__5145 := GGt(__e, V3288, V3289)
_ = let__5145

var ifres7591 Obj

if True == let__5145 {
tmp7597 := Call0(__e, SymFn(symshen_4loading_2))


var ifres7592 Obj

if True == tmp7597 {
tmp7593 := GCons(__e, sym_1, Nil)

tmp7594 := GCons(__e, sym_7, tmp7593)

tmp7595 := Call2(__e, SymFn(symelement_2), V3287, tmp7594)


ifres7592 = GNot(__e, tmp7595)


} else {
ifres7592 = False


}

ifres7591 = ifres7592


} else {
ifres7591 = False


}

var ifres7586 Obj

if True == ifres7591 {
tmp7587 := Call3(__e, SymFn(symshen_4app), V3287, MakeString("\n"), symshen_4a)


tmp7588 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("partial application of "), tmp7587)

tmp7589 := Call0(__e, SymFn(symstoutput))


tmp7590 := Call2(__e, SymFn(sympr), tmp7588, tmp7589)


ifres7586 = tmp7590


} else {
ifres7586 = symshen_4skip


}

let__5146 := ifres7586
_ = let__5146

return let__5145


panic("unreachable")
})

tmp7598 := Call2(__e, ns2_1set, symshen_4partial_1application_d_2, tmp7584)


_ = tmp7598

tmp7599 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dloading_2_d)
panic("unreachable")
})

tmp7600 := Call2(__e, ns2_1set, symshen_4loading_2, tmp7599)


_ = tmp7600

tmp7601 := MakeNative3(func(__e *ControlFlow, V3296, V3297, V3298 Obj) Obj {
tmp7616 := GEq(__e, MakeInteger(-1), V3297)

if True == tmp7616 {
return False
} else {
let__5147 := GLt(__e, V3297, V3298)
_ = let__5147

var ifres7613 Obj

if True == let__5147 {
tmp7614 := Call0(__e, SymFn(symshen_4loading_2))


ifres7613 = tmp7614


} else {
ifres7613 = False


}

var ifres7603 Obj

if True == ifres7613 {
tmp7605 := GEq(__e, V3298, MakeInteger(1))

var ifres7604 Obj

if True == tmp7605 {
ifres7604 = MakeString("")


} else {
ifres7604 = MakeString("s")


}

tmp7606 := Call3(__e, SymFn(symshen_4app), ifres7604, MakeString("\n"), symshen_4a)


tmp7608 := Call3(__e, SymFn(symshen_4app), V3298, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" argument"), tmp7606), symshen_4a)


tmp7610 := Call3(__e, SymFn(symshen_4app), V3296, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" might not like "), tmp7608), symshen_4a)


tmp7611 := Call0(__e, SymFn(symstoutput))


tmp7612 := Call2(__e, SymFn(sympr), tmp7610, tmp7611)


ifres7603 = tmp7612


} else {
ifres7603 = symshen_4skip


}

let__5148 := ifres7603
_ = let__5148

return let__5147


}


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4overapplication_2, tmp7601)




panic("unreachable")
})

