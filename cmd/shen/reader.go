package main

import . "github.com/pyrex41/shen-go/kl"

var ReaderMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp5433 := MakeNative1(func(__e *ControlFlow, V2556 Obj) Obj {
tmp5434 := PrimGuard1(__e, symread_1file_1as_1bytelist, PrimReadFileAsByteList, V2556)

let__5010 := tmp5434
_ = let__5010

tmp5435 := MakeNative0(func(__e *ControlFlow) Obj {
tmp5436 := MakeNative1(func(__e *ControlFlow, Z2559 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5s_1exprs_6), Z2559)
panic("unreachable")
})

return TailCall2(__e, SymFn(symcompile), tmp5436, let__5010)


panic("unreachable")
})

tmp5437 := MakeNative1(func(__e *ControlFlow, Z2560 Obj) Obj {
tmp5438 := GValue(__e, symshen_4_dresidue_d)

return TailCall1(__e, SymFn(symshen_4reader_1error), tmp5438)


panic("unreachable")
})

tmp5439 := Call2(__e, try_1catch, tmp5435, tmp5437)


let__5011 := tmp5439
_ = let__5011

tmp5440 := Call1(__e, SymFn(symshen_4process_1sexprs), let__5011)


let__5012 := tmp5440
_ = let__5012

return let__5012


panic("unreachable")
})

tmp5441 := Call2(__e, ns2_1set, symread_1file, tmp5433)


_ = tmp5441

tmp5442 := MakeNative1(func(__e *ControlFlow, V2562 Obj) Obj {
tmp5443 := GValue(__e, sym_dmaximum_1print_1sequence_1size_d)

tmp5444 := Call3(__e, SymFn(symshen_4reader_1error_1message), tmp5443, MakeInteger(0), V2562)


tmp5446 := Call1(__e, SymFn(symshen_4proc_1nl), PrimGuard2(__e, symcn, PrimStringConcat, MakeString("reader error near here: "), tmp5444))


return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp5446)


panic("unreachable")
})

tmp5447 := Call2(__e, ns2_1set, symshen_4reader_1error, tmp5442)


_ = tmp5447

tmp5448 := MakeNative3(func(__e *ControlFlow, V2570, V2571, V2572 Obj) Obj {
tmp5459 := GEq(__e, Nil, V2572)

if True == tmp5459 {
return MakeString("")
} else {
tmp5457 := GEq(__e, V2570, V2571)

if True == tmp5457 {
return MakeString("")
} else {
tmp5455 := GConsP(__e, V2572)

if True == tmp5455 {
tmp5449 := GHd(__e, V2572)

tmp5450 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, tmp5449)

tmp5451 := GAdd(__e, V2571, MakeInteger(1))

tmp5452 := GTl(__e, V2572)

tmp5453 := Call3(__e, SymFn(symshen_4reader_1error_1message), V2570, tmp5451, tmp5452)


return PrimGuard2(__e, symcn, PrimStringConcat, tmp5450, tmp5453)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4reader_1error_1message)
}


}


}


panic("unreachable")
})

tmp5460 := Call2(__e, ns2_1set, symshen_4reader_1error_1message, tmp5448)


_ = tmp5460

tmp5461 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dit_d)
panic("unreachable")
})

tmp5462 := Call2(__e, ns2_1set, symit, tmp5461)


_ = tmp5462

tmp5463 := MakeNative1(func(__e *ControlFlow, V2573 Obj) Obj {
tmp5464 := PrimGuard2(__e, symopen, PrimOpenStream, V2573, symin)

let__5013 := tmp5464
_ = let__5013

tmp5465 := PrimGuard1(__e, symread_1byte, PrimReadByte, let__5013)

let__5014 := tmp5465
_ = let__5014

tmp5466 := Call3(__e, SymFn(symshen_4read_1file_1as_1bytelist_1help), let__5013, let__5014, Nil)


let__5015 := tmp5466
_ = let__5015

tmp5467 := PrimGuard1(__e, symclose, PrimCloseStream, let__5013)

let__5016 := tmp5467
_ = let__5016

return TailCall1(__e, SymFn(symreverse), let__5015)


panic("unreachable")
})

tmp5468 := Call2(__e, ns2_1set, symread_1file_1as_1bytelist, tmp5463)


_ = tmp5468

tmp5469 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V2578 := __self1
_ = V2578
V2579 := __self2
_ = V2579
V2580 := __self3
_ = V2580
tmp5473 := GEq(__e, MakeInteger(-1), V2579)

if True == tmp5473 {
return V2580
} else {
tmp5470 := PrimGuard1(__e, symread_1byte, PrimReadByte, V2578)

tmp5471 := GCons(__e, V2579, V2580)

if SymFn(symshen_4read_1file_1as_1bytelist_1help) == __self {
__self1, __self2, __self3 = V2578, tmp5470, tmp5471
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4read_1file_1as_1bytelist_1help), V2578, tmp5470, tmp5471)


}


panic("unreachable")
})
return __self
})()

tmp5474 := Call2(__e, ns2_1set, symshen_4read_1file_1as_1bytelist_1help, tmp5469)


_ = tmp5474

tmp5475 := MakeNative1(func(__e *ControlFlow, V2581 Obj) Obj {
tmp5476 := PrimGuard2(__e, symopen, PrimOpenStream, V2581, symin)

let__5017 := tmp5476
_ = let__5017

tmp5477 := PrimGuard1(__e, symread_1byte, PrimReadByte, let__5017)

return TailCall3(__e, SymFn(symshen_4rfas_1h), let__5017, tmp5477, MakeString(""))


panic("unreachable")
})

tmp5478 := Call2(__e, ns2_1set, symread_1file_1as_1string, tmp5475)


_ = tmp5478

tmp5479 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V2583 := __self1
_ = V2583
V2584 := __self2
_ = V2584
V2585 := __self3
_ = V2585
tmp5485 := GEq(__e, MakeInteger(-1), V2584)

if True == tmp5485 {
tmp5480 := PrimGuard1(__e, symclose, PrimCloseStream, V2583)

_ = tmp5480

return V2585


} else {
tmp5481 := PrimGuard1(__e, symread_1byte, PrimReadByte, V2583)

tmp5482 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, V2584)

if SymFn(symshen_4rfas_1h) == __self {
__self1, __self2, __self3 = V2583, tmp5481, PrimGuard2(__e, symcn, PrimStringConcat, V2585, tmp5482)
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4rfas_1h), V2583, tmp5481, PrimGuard2(__e, symcn, PrimStringConcat, V2585, tmp5482))


}


panic("unreachable")
})
return __self
})()

tmp5486 := Call2(__e, ns2_1set, symshen_4rfas_1h, tmp5479)


_ = tmp5486

tmp5487 := MakeNative1(func(__e *ControlFlow, V2586 Obj) Obj {
tmp5488 := Call1(__e, SymFn(symread), V2586)


return TailCall1(__e, SymFn(symeval_1kl), tmp5488)


panic("unreachable")
})

tmp5489 := Call2(__e, ns2_1set, syminput, tmp5487)


_ = tmp5489

tmp5490 := MakeNative2(func(__e *ControlFlow, V2587, V2588 Obj) Obj {
tmp5491 := Call1(__e, SymFn(symshen_4monotype), V2587)


let__5018 := tmp5491
_ = let__5018

tmp5492 := Call1(__e, SymFn(symread), V2588)


let__5019 := tmp5492
_ = let__5019

tmp5498 := Call2(__e, SymFn(symshen_4typecheck), let__5019, V2587)


tmp5499 := GEq(__e, False, tmp5498)

if True == tmp5499 {
tmp5493 := Call3(__e, SymFn(symshen_4app), V2587, MakeString("\n"), symshen_4r)


tmp5495 := Call3(__e, SymFn(symshen_4app), let__5019, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" is not of type "), tmp5493), symshen_4r)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("type error: "), tmp5495))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("type error: "), tmp5495)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


} else {
return TailCall1(__e, SymFn(symeval_1kl), let__5019)
}


panic("unreachable")
})

tmp5500 := Call2(__e, ns2_1set, symshen_4input_1h_7, tmp5490)


_ = tmp5500

tmp5501 := MakeNative1(func(__e *ControlFlow, V2591 Obj) Obj {
tmp5508 := GConsP(__e, V2591)

if True == tmp5508 {
tmp5502 := MakeNative1(func(__e *ControlFlow, Z2592 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4monotype), Z2592)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp5502, V2591)


} else {
tmp5506 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V2591)

if True == tmp5506 {
tmp5503 := Call3(__e, SymFn(symshen_4app), V2591, MakeString("\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("input+ expects a monotype: not "), tmp5503))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("input+ expects a monotype: not "), tmp5503)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


} else {
return V2591
}


}


panic("unreachable")
})

tmp5509 := Call2(__e, ns2_1set, symshen_4monotype, tmp5501)


_ = tmp5509

tmp5510 := MakeNative1(func(__e *ControlFlow, V2593 Obj) Obj {
tmp5511 := Call1(__e, SymFn(symshen_4my_1read_1byte), V2593)


tmp5512 := MakeNative1(func(__e *ControlFlow, Z2594 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4return_2), Z2594)
panic("unreachable")
})

return TailCall4(__e, SymFn(symshen_4read_1loop), V2593, tmp5511, Nil, tmp5512)


panic("unreachable")
})

tmp5513 := Call2(__e, ns2_1set, symlineread, tmp5510)


_ = tmp5513

tmp5514 := MakeNative1(func(__e *ControlFlow, V2595 Obj) Obj {
tmp5515 := Call1(__e, SymFn(symshen_4str_1_6bytes), V2595)


let__5020 := tmp5515
_ = let__5020

tmp5516 := MakeNative1(func(__e *ControlFlow, Z2598 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5s_1exprs_6), Z2598)
panic("unreachable")
})

tmp5517 := Call2(__e, SymFn(symcompile), tmp5516, let__5020)


let__5021 := tmp5517
_ = let__5021

tmp5518 := Call1(__e, SymFn(symshen_4process_1sexprs), let__5021)


let__5022 := tmp5518
_ = let__5022

return let__5022


panic("unreachable")
})

tmp5519 := Call2(__e, ns2_1set, symread_1from_1string, tmp5514)


_ = tmp5519

tmp5520 := MakeNative1(func(__e *ControlFlow, V2600 Obj) Obj {
tmp5521 := Call1(__e, SymFn(symshen_4str_1_6bytes), V2600)


let__5023 := tmp5521
_ = let__5023

tmp5522 := MakeNative1(func(__e *ControlFlow, Z2603 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5s_1exprs_6), Z2603)
panic("unreachable")
})

tmp5523 := Call2(__e, SymFn(symcompile), tmp5522, let__5023)


let__5024 := tmp5523
_ = let__5024

return let__5024


panic("unreachable")
})

tmp5524 := Call2(__e, ns2_1set, symread_1from_1string_1unprocessed, tmp5520)


_ = tmp5524

tmp5525 := MakeNative1(func(__e *ControlFlow, V2604 Obj) Obj {
tmp5533 := GEq(__e, MakeString(""), V2604)

if True == tmp5533 {
return Nil
} else {
tmp5531 := Call1(__e, SymFn(symshen_4_7string_2), V2604)


if True == tmp5531 {
tmp5526 := Call1(__e, SymFn(symhdstr), V2604)


tmp5527 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp5526)

tmp5529 := Call1(__e, SymFn(symshen_4str_1_6bytes), PrimGuard1(__e, symtlstr, PrimTailString, V2604))


return GCons(__e, tmp5527, tmp5529)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4str_1_6bytes)
}


}


panic("unreachable")
})

tmp5534 := Call2(__e, ns2_1set, symshen_4str_1_6bytes, tmp5525)


_ = tmp5534

tmp5535 := MakeNative1(func(__e *ControlFlow, V2605 Obj) Obj {
tmp5536 := Call1(__e, SymFn(symshen_4my_1read_1byte), V2605)


tmp5537 := MakeNative1(func(__e *ControlFlow, Z2606 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4whitespace_2), Z2606)
panic("unreachable")
})

tmp5538 := Call4(__e, SymFn(symshen_4read_1loop), V2605, tmp5536, Nil, tmp5537)


return GHd(__e, tmp5538)


panic("unreachable")
})

tmp5539 := Call2(__e, ns2_1set, symread, tmp5535)


_ = tmp5539

tmp5540 := MakeNative1(func(__e *ControlFlow, V2607 Obj) Obj {
tmp5543 := Call1(__e, SymFn(symchar_1stream_2), V2607)


if True == tmp5543 {
tmp5541 := Call1(__e, SymFn(symread_1unit_1string), V2607)


return PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp5541)


} else {
return PrimGuard1(__e, symread_1byte, PrimReadByte, V2607)
}


panic("unreachable")
})

tmp5544 := Call2(__e, ns2_1set, symshen_4my_1read_1byte, tmp5540)


_ = tmp5544

tmp5545 := (func() Obj {
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
tmp5568 := GEq(__e, MakeInteger(94), V2613)

if True == tmp5568 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("read aborted"))
} else {
tmp5566 := GEq(__e, MakeInteger(-1), V2613)

if True == tmp5566 {
tmp5549 := Call1(__e, SymFn(symempty_2), V2614)


if True == tmp5549 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("error: empty stream"))
} else {
tmp5546 := MakeNative1(func(__e *ControlFlow, Z2616 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5s_1exprs_6), Z2616)
panic("unreachable")
})

tmp5547 := Call2(__e, SymFn(symcompile), tmp5546, V2614)


return TailCall1(__e, SymFn(symshen_4process_1sexprs), tmp5547)


}


} else {
tmp5564 := GEq(__e, MakeInteger(0), V2613)

if True == tmp5564 {
tmp5550 := Call1(__e, SymFn(symshen_4my_1read_1byte), V2612)


if SymFn(symshen_4read_1loop) == __self {
__self1, __self2, __self3, __self4 = V2612, tmp5550, V2614, V2615
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4read_1loop), V2612, tmp5550, V2614, V2615)


} else {
tmp5562 := Call1(__e, V2615, V2613)


if True == tmp5562 {
tmp5551 := Call1(__e, SymFn(symshen_4try_1parse), V2614)


let__5025 := tmp5551
_ = let__5025

tmp5557 := Call1(__e, SymFn(symshen_4nothing_1doing_2), let__5025)


if True == tmp5557 {
tmp5552 := Call1(__e, SymFn(symshen_4my_1read_1byte), V2612)


tmp5553 := GCons(__e, V2613, Nil)

tmp5554 := Call2(__e, SymFn(symappend), V2614, tmp5553)


if SymFn(symshen_4read_1loop) == __self {
__self1, __self2, __self3, __self4 = V2612, tmp5552, tmp5554, V2615
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4read_1loop), V2612, tmp5552, tmp5554, V2615)


} else {
tmp5555 := Call1(__e, SymFn(symshen_4record_1it), V2614)


_ = tmp5555

return TailCall1(__e, SymFn(symshen_4process_1sexprs), let__5025)


}


} else {
tmp5558 := Call1(__e, SymFn(symshen_4my_1read_1byte), V2612)


tmp5559 := GCons(__e, V2613, Nil)

tmp5560 := Call2(__e, SymFn(symappend), V2614, tmp5559)


if SymFn(symshen_4read_1loop) == __self {
__self1, __self2, __self3, __self4 = V2612, tmp5558, tmp5560, V2615
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4read_1loop), V2612, tmp5558, tmp5560, V2615)


}


}


}


}


panic("unreachable")
})
return __self
})()

tmp5569 := Call2(__e, ns2_1set, symshen_4read_1loop, tmp5545)


_ = tmp5569

tmp5570 := MakeNative1(func(__e *ControlFlow, V2618 Obj) Obj {
tmp5571 := MakeNative0(func(__e *ControlFlow) Obj {
tmp5572 := MakeNative1(func(__e *ControlFlow, Z2619 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5s_1exprs_6), Z2619)
panic("unreachable")
})

return TailCall2(__e, SymFn(symcompile), tmp5572, V2618)


panic("unreachable")
})

tmp5573 := MakeNative1(func(__e *ControlFlow, Z2620 Obj) Obj {
return symshen_4i_1failed_b
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp5571, tmp5573)


panic("unreachable")
})

tmp5574 := Call2(__e, ns2_1set, symshen_4try_1parse, tmp5570)


_ = tmp5574

tmp5575 := MakeNative1(func(__e *ControlFlow, V2623 Obj) Obj {
tmp5579 := GEq(__e, symshen_4i_1failed_b, V2623)

if True == tmp5579 {
return True
} else {
tmp5577 := GEq(__e, Nil, V2623)

if True == tmp5577 {
return True
} else {
return False
}


}


panic("unreachable")
})

tmp5580 := Call2(__e, ns2_1set, symshen_4nothing_1doing_2, tmp5575)


_ = tmp5580

tmp5581 := MakeNative1(func(__e *ControlFlow, V2624 Obj) Obj {
tmp5582 := Call1(__e, SymFn(symshen_4bytes_1_6string), V2624)


return PrimGuard2(__e, symset, PrimSet, symshen_4_dit_d, tmp5582)


panic("unreachable")
})

tmp5583 := Call2(__e, ns2_1set, symshen_4record_1it, tmp5581)


_ = tmp5583

tmp5584 := MakeNative1(func(__e *ControlFlow, V2625 Obj) Obj {
tmp5592 := GEq(__e, Nil, V2625)

if True == tmp5592 {
return MakeString("")
} else {
tmp5590 := GConsP(__e, V2625)

if True == tmp5590 {
tmp5585 := GHd(__e, V2625)

tmp5586 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, tmp5585)

tmp5587 := GTl(__e, V2625)

tmp5588 := Call1(__e, SymFn(symshen_4bytes_1_6string), tmp5587)


return PrimGuard2(__e, symcn, PrimStringConcat, tmp5586, tmp5588)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4bytes_1_6string)
}


}


panic("unreachable")
})

tmp5593 := Call2(__e, ns2_1set, symshen_4bytes_1_6string, tmp5584)


_ = tmp5593

tmp5594 := MakeNative1(func(__e *ControlFlow, V2626 Obj) Obj {
tmp5595 := Call1(__e, SymFn(symshen_4unpackage_emacroexpand), V2626)


let__5026 := tmp5595
_ = let__5026

tmp5596 := Call1(__e, SymFn(symshen_4find_1arities), let__5026)


let__5027 := tmp5596
_ = let__5027

tmp5597 := Call1(__e, SymFn(symshen_4find_1types), let__5026)


let__5028 := tmp5597
_ = let__5028

tmp5598 := MakeNative1(func(__e *ControlFlow, Z2630 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4process_1applications), Z2630, let__5028)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp5598, let__5026)


panic("unreachable")
})

tmp5599 := Call2(__e, ns2_1set, symshen_4process_1sexprs, tmp5594)


_ = tmp5599

tmp5600 := MakeNative1(func(__e *ControlFlow, V2631 Obj) Obj {
tmp5622 := GConsP(__e, V2631)

var ifres5613 Obj

if True == tmp5622 {
tmp5620 := GTl(__e, V2631)

tmp5621 := GConsP(__e, tmp5620)

var ifres5615 Obj

if True == tmp5621 {
tmp5617 := GHd(__e, V2631)

tmp5618 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp5619 := GEq(__e, tmp5617, tmp5618)

var ifres5616 Obj

if True == tmp5619 {
ifres5616 = True


} else {
ifres5616 = False


}

ifres5615 = ifres5616


} else {
ifres5615 = False


}

var ifres5614 Obj

if True == ifres5615 {
ifres5614 = True


} else {
ifres5614 = False


}

ifres5613 = ifres5614


} else {
ifres5613 = False


}

if True == ifres5613 {
tmp5601 := GTl(__e, V2631)

tmp5602 := GHd(__e, tmp5601)

tmp5603 := GTl(__e, V2631)

tmp5604 := GTl(__e, tmp5603)

tmp5605 := Call1(__e, SymFn(symshen_4find_1types), tmp5604)


return GCons(__e, tmp5602, tmp5605)


} else {
tmp5611 := GConsP(__e, V2631)

if True == tmp5611 {
tmp5606 := GHd(__e, V2631)

tmp5607 := Call1(__e, SymFn(symshen_4find_1types), tmp5606)


tmp5608 := GTl(__e, V2631)

tmp5609 := Call1(__e, SymFn(symshen_4find_1types), tmp5608)


return TailCall2(__e, SymFn(symappend), tmp5607, tmp5609)


} else {
return Nil
}


}


panic("unreachable")
})

tmp5623 := Call2(__e, ns2_1set, symshen_4find_1types, tmp5600)


_ = tmp5623

tmp5624 := MakeNative1(func(__e *ControlFlow, V2634 Obj) Obj {
tmp5673 := GConsP(__e, V2634)

var ifres5654 Obj

if True == tmp5673 {
tmp5671 := GHd(__e, V2634)

tmp5672 := GEq(__e, symdefine, tmp5671)

var ifres5656 Obj

if True == tmp5672 {
tmp5669 := GTl(__e, V2634)

tmp5670 := GConsP(__e, tmp5669)

var ifres5658 Obj

if True == tmp5670 {
tmp5666 := GTl(__e, V2634)

tmp5667 := GTl(__e, tmp5666)

tmp5668 := GConsP(__e, tmp5667)

var ifres5660 Obj

if True == tmp5668 {
tmp5662 := GTl(__e, V2634)

tmp5663 := GTl(__e, tmp5662)

tmp5664 := GHd(__e, tmp5663)

tmp5665 := GEq(__e, sym_i, tmp5664)

var ifres5661 Obj

if True == tmp5665 {
ifres5661 = True


} else {
ifres5661 = False


}

ifres5660 = ifres5661


} else {
ifres5660 = False


}

var ifres5659 Obj

if True == ifres5660 {
ifres5659 = True


} else {
ifres5659 = False


}

ifres5658 = ifres5659


} else {
ifres5658 = False


}

var ifres5657 Obj

if True == ifres5658 {
ifres5657 = True


} else {
ifres5657 = False


}

ifres5656 = ifres5657


} else {
ifres5656 = False


}

var ifres5655 Obj

if True == ifres5656 {
ifres5655 = True


} else {
ifres5655 = False


}

ifres5654 = ifres5655


} else {
ifres5654 = False


}

if True == ifres5654 {
tmp5625 := GTl(__e, V2634)

tmp5626 := GHd(__e, tmp5625)

tmp5627 := GTl(__e, V2634)

tmp5628 := GHd(__e, tmp5627)

tmp5629 := GTl(__e, V2634)

tmp5630 := GTl(__e, tmp5629)

tmp5631 := GTl(__e, tmp5630)

tmp5632 := Call3(__e, SymFn(symshen_4find_1arity), tmp5628, MakeInteger(1), tmp5631)


return TailCall2(__e, SymFn(symshen_4store_1arity), tmp5626, tmp5632)


} else {
tmp5652 := GConsP(__e, V2634)

var ifres5644 Obj

if True == tmp5652 {
tmp5650 := GHd(__e, V2634)

tmp5651 := GEq(__e, symdefine, tmp5650)

var ifres5646 Obj

if True == tmp5651 {
tmp5648 := GTl(__e, V2634)

tmp5649 := GConsP(__e, tmp5648)

var ifres5647 Obj

if True == tmp5649 {
ifres5647 = True


} else {
ifres5647 = False


}

ifres5646 = ifres5647


} else {
ifres5646 = False


}

var ifres5645 Obj

if True == ifres5646 {
ifres5645 = True


} else {
ifres5645 = False


}

ifres5644 = ifres5645


} else {
ifres5644 = False


}

if True == ifres5644 {
tmp5633 := GTl(__e, V2634)

tmp5634 := GHd(__e, tmp5633)

tmp5635 := GTl(__e, V2634)

tmp5636 := GHd(__e, tmp5635)

tmp5637 := GTl(__e, V2634)

tmp5638 := GTl(__e, tmp5637)

tmp5639 := Call3(__e, SymFn(symshen_4find_1arity), tmp5636, MakeInteger(0), tmp5638)


return TailCall2(__e, SymFn(symshen_4store_1arity), tmp5634, tmp5639)


} else {
tmp5642 := GConsP(__e, V2634)

if True == tmp5642 {
tmp5640 := MakeNative1(func(__e *ControlFlow, Z2635 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4find_1arities), Z2635)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp5640, V2634)


} else {
return symshen_4skip
}


}


}


panic("unreachable")
})

tmp5674 := Call2(__e, ns2_1set, symshen_4find_1arities, tmp5624)


_ = tmp5674

tmp5675 := MakeNative2(func(__e *ControlFlow, V2636, V2637 Obj) Obj {
tmp5676 := Call1(__e, SymFn(symarity), V2636)


let__5029 := tmp5676
_ = let__5029

tmp5687 := GEq(__e, let__5029, MakeInteger(-1))

if True == tmp5687 {
return TailCall2(__e, SymFn(symshen_4execute_1store_1arity), V2636, V2637)
} else {
tmp5685 := GEq(__e, let__5029, V2637)

if True == tmp5685 {
return symshen_4skip
} else {
tmp5683 := Call1(__e, SymFn(symshen_4sysfunc_2), V2636)


if True == tmp5683 {
tmp5677 := Call3(__e, SymFn(symshen_4app), V2636, MakeString(" is a system function\n"), symshen_4a)


return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp5677)


} else {
tmp5678 := Call3(__e, SymFn(symshen_4app), V2636, MakeString(" may cause errors\n"), symshen_4a)


tmp5679 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("changing the arity of "), tmp5678)

tmp5680 := Call0(__e, SymFn(symstoutput))


tmp5681 := Call2(__e, SymFn(sympr), tmp5679, tmp5680)


_ = tmp5681

return TailCall2(__e, SymFn(symshen_4execute_1store_1arity), V2636, V2637)


}


}


}


panic("unreachable")
})

tmp5688 := Call2(__e, ns2_1set, symshen_4store_1arity, tmp5675)


_ = tmp5688

tmp5689 := MakeNative2(func(__e *ControlFlow, V2639, V2640 Obj) Obj {
tmp5694 := GEq(__e, MakeInteger(0), V2640)

if True == tmp5694 {
tmp5690 := GValue(__e, sym_dproperty_1vector_d)

return TailCall4(__e, SymFn(symput), V2639, symarity, MakeInteger(0), tmp5690)


} else {
tmp5691 := GValue(__e, sym_dproperty_1vector_d)

tmp5692 := Call4(__e, SymFn(symput), V2639, symarity, V2640, tmp5691)


_ = tmp5692

return TailCall2(__e, SymFn(symshen_4update_1lambdatable), V2639, V2640)


}


panic("unreachable")
})

tmp5695 := Call2(__e, ns2_1set, symshen_4execute_1store_1arity, tmp5689)


_ = tmp5695

tmp5696 := MakeNative2(func(__e *ControlFlow, V2641, V2642 Obj) Obj {
tmp5697 := GCons(__e, V2641, Nil)

tmp5698 := Call2(__e, SymFn(symshen_4lambda_1function), tmp5697, V2642)


tmp5699 := Call1(__e, SymFn(symeval_1kl), tmp5698)


let__5030 := tmp5699
_ = let__5030

tmp5700 := GCons(__e, V2641, let__5030)

tmp5701 := Call1(__e, SymFn(symshen_4set_1lambda_1form_1entry), tmp5700)


let__5031 := tmp5701
_ = let__5031

return V2642


panic("unreachable")
})

tmp5702 := Call2(__e, ns2_1set, symshen_4update_1lambdatable, tmp5696)


_ = tmp5702

tmp5703 := MakeNative2(func(__e *ControlFlow, V2647, V2648 Obj) Obj {
tmp5719 := GEq(__e, MakeInteger(0), V2648)

if True == tmp5719 {
return symshen_4skip
} else {
tmp5717 := GEq(__e, MakeInteger(1), V2648)

if True == tmp5717 {
tmp5704 := Call1(__e, SymFn(symgensym), symY)


let__5032 := tmp5704
_ = let__5032

tmp5705 := GCons(__e, let__5032, Nil)

tmp5706 := Call2(__e, SymFn(symappend), V2647, tmp5705)


tmp5707 := GCons(__e, tmp5706, Nil)

tmp5708 := GCons(__e, let__5032, tmp5707)

return GCons(__e, symlambda, tmp5708)


} else {
tmp5709 := Call1(__e, SymFn(symgensym), symY)


let__5033 := tmp5709
_ = let__5033

tmp5710 := GCons(__e, let__5033, Nil)

tmp5711 := Call2(__e, SymFn(symappend), V2647, tmp5710)


tmp5713 := Call2(__e, SymFn(symshen_4lambda_1function), tmp5711, GSub(__e, V2648, MakeInteger(1)))


tmp5714 := GCons(__e, tmp5713, Nil)

tmp5715 := GCons(__e, let__5033, tmp5714)

return GCons(__e, symlambda, tmp5715)


}


}


panic("unreachable")
})

tmp5720 := Call2(__e, ns2_1set, symshen_4lambda_1function, tmp5703)


_ = tmp5720

tmp5721 := MakeNative3(func(__e *ControlFlow, V2660, V2661, V2662 Obj) Obj {
tmp5744 := GEq(__e, Nil, V2662)

if True == tmp5744 {
tmp5722 := GCons(__e, V2660, V2661)

return GCons(__e, tmp5722, Nil)


} else {
tmp5742 := GConsP(__e, V2662)

var ifres5733 Obj

if True == tmp5742 {
tmp5740 := GHd(__e, V2662)

tmp5741 := GConsP(__e, tmp5740)

var ifres5735 Obj

if True == tmp5741 {
tmp5737 := GHd(__e, V2662)

tmp5738 := GHd(__e, tmp5737)

tmp5739 := GEq(__e, V2660, tmp5738)

var ifres5736 Obj

if True == tmp5739 {
ifres5736 = True


} else {
ifres5736 = False


}

ifres5735 = ifres5736


} else {
ifres5735 = False


}

var ifres5734 Obj

if True == ifres5735 {
ifres5734 = True


} else {
ifres5734 = False


}

ifres5733 = ifres5734


} else {
ifres5733 = False


}

if True == ifres5733 {
tmp5723 := GHd(__e, V2662)

tmp5724 := GHd(__e, tmp5723)

tmp5725 := GCons(__e, tmp5724, V2661)

tmp5726 := GTl(__e, V2662)

return GCons(__e, tmp5725, tmp5726)


} else {
tmp5731 := GConsP(__e, V2662)

if True == tmp5731 {
tmp5727 := GHd(__e, V2662)

tmp5728 := GTl(__e, V2662)

tmp5729 := Call3(__e, SymFn(symshen_4assoc_1_6), V2660, V2661, tmp5728)


return GCons(__e, tmp5727, tmp5729)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.assoc->"))
}


}


}


panic("unreachable")
})

tmp5745 := Call2(__e, ns2_1set, symshen_4assoc_1_6, tmp5721)


_ = tmp5745

tmp5746 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V2677 := __self1
_ = V2677
V2678 := __self2
_ = V2678
V2679 := __self3
_ = V2679
tmp5793 := GEq(__e, MakeInteger(0), V2678)

var ifres5786 Obj

if True == tmp5793 {
tmp5792 := GConsP(__e, V2679)

var ifres5788 Obj

if True == tmp5792 {
tmp5790 := GHd(__e, V2679)

tmp5791 := GEq(__e, tmp5790, sym_1_6)

var ifres5789 Obj

if True == tmp5791 {
ifres5789 = True


} else {
ifres5789 = False


}

ifres5788 = ifres5789


} else {
ifres5788 = False


}

var ifres5787 Obj

if True == ifres5788 {
ifres5787 = True


} else {
ifres5787 = False


}

ifres5786 = ifres5787


} else {
ifres5786 = False


}

if True == ifres5786 {
return MakeInteger(0)
} else {
tmp5784 := GEq(__e, MakeInteger(0), V2678)

var ifres5777 Obj

if True == tmp5784 {
tmp5783 := GConsP(__e, V2679)

var ifres5779 Obj

if True == tmp5783 {
tmp5781 := GHd(__e, V2679)

tmp5782 := GEq(__e, tmp5781, sym_5_1)

var ifres5780 Obj

if True == tmp5782 {
ifres5780 = True


} else {
ifres5780 = False


}

ifres5779 = ifres5780


} else {
ifres5779 = False


}

var ifres5778 Obj

if True == ifres5779 {
ifres5778 = True


} else {
ifres5778 = False


}

ifres5777 = ifres5778


} else {
ifres5777 = False


}

if True == ifres5777 {
return MakeInteger(0)
} else {
tmp5775 := GEq(__e, MakeInteger(0), V2678)

var ifres5772 Obj

if True == tmp5775 {
tmp5774 := GConsP(__e, V2679)

var ifres5773 Obj

if True == tmp5774 {
ifres5773 = True


} else {
ifres5773 = False


}

ifres5772 = ifres5773


} else {
ifres5772 = False


}

if True == ifres5772 {
tmp5747 := GTl(__e, V2679)

tmp5748 := Call3(__e, SymFn(symshen_4find_1arity), V2677, MakeInteger(0), tmp5747)


return GAdd(__e, MakeInteger(1), tmp5748)


} else {
tmp5770 := GEq(__e, MakeInteger(1), V2678)

var ifres5763 Obj

if True == tmp5770 {
tmp5769 := GConsP(__e, V2679)

var ifres5765 Obj

if True == tmp5769 {
tmp5767 := GHd(__e, V2679)

tmp5768 := GEq(__e, sym_j, tmp5767)

var ifres5766 Obj

if True == tmp5768 {
ifres5766 = True


} else {
ifres5766 = False


}

ifres5765 = ifres5766


} else {
ifres5765 = False


}

var ifres5764 Obj

if True == ifres5765 {
ifres5764 = True


} else {
ifres5764 = False


}

ifres5763 = ifres5764


} else {
ifres5763 = False


}

if True == ifres5763 {
tmp5749 := GTl(__e, V2679)

if SymFn(symshen_4find_1arity) == __self {
__self1, __self2, __self3 = V2677, MakeInteger(0), tmp5749
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4find_1arity), V2677, MakeInteger(0), tmp5749)


} else {
tmp5761 := GEq(__e, MakeInteger(1), V2678)

var ifres5758 Obj

if True == tmp5761 {
tmp5760 := GConsP(__e, V2679)

var ifres5759 Obj

if True == tmp5760 {
ifres5759 = True


} else {
ifres5759 = False


}

ifres5758 = ifres5759


} else {
ifres5758 = False


}

if True == ifres5758 {
tmp5750 := GTl(__e, V2679)

if SymFn(symshen_4find_1arity) == __self {
__self1, __self2, __self3 = V2677, MakeInteger(1), tmp5750
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4find_1arity), V2677, MakeInteger(1), tmp5750)


} else {
tmp5756 := GEq(__e, MakeInteger(1), V2678)

if True == tmp5756 {
tmp5751 := Call3(__e, SymFn(symshen_4app), V2677, MakeString(" definition: missing }\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("syntax error in "), tmp5751))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("syntax error in "), tmp5751)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


} else {
tmp5753 := Call3(__e, SymFn(symshen_4app), V2677, MakeString(" definition: missing -> or <-\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("syntax error in "), tmp5753))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("syntax error in "), tmp5753)
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

tmp5794 := Call2(__e, ns2_1set, symshen_4find_1arity, tmp5746)


_ = tmp5794

tmp5795 := MakeNative1(func(__e *ControlFlow, V2680 Obj) Obj {
tmp5796 := Call1(__e, SymFn(symshen_4_5lsb_6), V2680)


let__5035 := tmp5796
_ = let__5035

tmp5820 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5035)


var ifres5797 Obj

if True == tmp5820 {
tmp5798 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5797 = tmp5798


} else {
tmp5799 := Call1(__e, SymFn(symshen_4in_1_6), let__5035)


let__5036 := tmp5799
_ = let__5036

tmp5800 := Call1(__e, SymFn(symshen_4_5s_1exprs1_6), let__5036)


let__5037 := tmp5800
_ = let__5037

tmp5819 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5037)


var ifres5801 Obj

if True == tmp5819 {
tmp5802 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5801 = tmp5802


} else {
tmp5803 := Call1(__e, SymFn(symshen_4_5_1out), let__5037)


let__5038 := tmp5803
_ = let__5038

tmp5804 := Call1(__e, SymFn(symshen_4in_1_6), let__5037)


let__5039 := tmp5804
_ = let__5039

tmp5805 := Call1(__e, SymFn(symshen_4_5rsb_6), let__5039)


let__5040 := tmp5805
_ = let__5040

tmp5818 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5040)


var ifres5806 Obj

if True == tmp5818 {
tmp5807 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5806 = tmp5807


} else {
tmp5808 := Call1(__e, SymFn(symshen_4in_1_6), let__5040)


let__5041 := tmp5808
_ = let__5041

tmp5809 := Call1(__e, SymFn(symshen_4_5s_1exprs2_6), let__5041)


let__5042 := tmp5809
_ = let__5042

tmp5817 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5042)


var ifres5810 Obj

if True == tmp5817 {
tmp5811 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5810 = tmp5811


} else {
tmp5812 := Call1(__e, SymFn(symshen_4_5_1out), let__5042)


let__5043 := tmp5812
_ = let__5043

tmp5813 := Call1(__e, SymFn(symshen_4in_1_6), let__5042)


let__5044 := tmp5813
_ = let__5044

tmp5814 := Call1(__e, SymFn(symshen_4cons_1form), let__5038)


tmp5815 := GCons(__e, tmp5814, let__5043)

tmp5816 := Call2(__e, SymFn(symshen_4comb), let__5044, tmp5815)


ifres5810 = tmp5816


}

ifres5806 = ifres5810


}

ifres5801 = ifres5806


}

ifres5797 = ifres5801


}

let__5034 := ifres5797
_ = let__5034

tmp6014 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5034)


if True == tmp6014 {
tmp5821 := Call1(__e, SymFn(symshen_4_5lrb_6), V2680)


let__5046 := tmp5821
_ = let__5046

tmp5844 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5046)


var ifres5822 Obj

if True == tmp5844 {
tmp5823 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5822 = tmp5823


} else {
tmp5824 := Call1(__e, SymFn(symshen_4in_1_6), let__5046)


let__5047 := tmp5824
_ = let__5047

tmp5825 := Call1(__e, SymFn(symshen_4_5s_1exprs1_6), let__5047)


let__5048 := tmp5825
_ = let__5048

tmp5843 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5048)


var ifres5826 Obj

if True == tmp5843 {
tmp5827 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5826 = tmp5827


} else {
tmp5828 := Call1(__e, SymFn(symshen_4_5_1out), let__5048)


let__5049 := tmp5828
_ = let__5049

tmp5829 := Call1(__e, SymFn(symshen_4in_1_6), let__5048)


let__5050 := tmp5829
_ = let__5050

tmp5830 := Call1(__e, SymFn(symshen_4_5rrb_6), let__5050)


let__5051 := tmp5830
_ = let__5051

tmp5842 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5051)


var ifres5831 Obj

if True == tmp5842 {
tmp5832 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5831 = tmp5832


} else {
tmp5833 := Call1(__e, SymFn(symshen_4in_1_6), let__5051)


let__5052 := tmp5833
_ = let__5052

tmp5834 := Call1(__e, SymFn(symshen_4_5s_1exprs2_6), let__5052)


let__5053 := tmp5834
_ = let__5053

tmp5841 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5053)


var ifres5835 Obj

if True == tmp5841 {
tmp5836 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5835 = tmp5836


} else {
tmp5837 := Call1(__e, SymFn(symshen_4_5_1out), let__5053)


let__5054 := tmp5837
_ = let__5054

tmp5838 := Call1(__e, SymFn(symshen_4in_1_6), let__5053)


let__5055 := tmp5838
_ = let__5055

tmp5839 := Call2(__e, SymFn(symshen_4add_1sexpr), let__5049, let__5054)


tmp5840 := Call2(__e, SymFn(symshen_4comb), let__5055, tmp5839)


ifres5835 = tmp5840


}

ifres5831 = ifres5835


}

ifres5826 = ifres5831


}

ifres5822 = ifres5826


}

let__5045 := ifres5822
_ = let__5045

tmp6012 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5045)


if True == tmp6012 {
tmp5845 := Call1(__e, SymFn(symshen_4_5lcurly_6), V2680)


let__5057 := tmp5845
_ = let__5057

tmp5857 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5057)


var ifres5846 Obj

if True == tmp5857 {
tmp5847 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5846 = tmp5847


} else {
tmp5848 := Call1(__e, SymFn(symshen_4in_1_6), let__5057)


let__5058 := tmp5848
_ = let__5058

tmp5849 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__5058)


let__5059 := tmp5849
_ = let__5059

tmp5856 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5059)


var ifres5850 Obj

if True == tmp5856 {
tmp5851 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5850 = tmp5851


} else {
tmp5852 := Call1(__e, SymFn(symshen_4_5_1out), let__5059)


let__5060 := tmp5852
_ = let__5060

tmp5853 := Call1(__e, SymFn(symshen_4in_1_6), let__5059)


let__5061 := tmp5853
_ = let__5061

tmp5854 := GCons(__e, sym_i, let__5060)

tmp5855 := Call2(__e, SymFn(symshen_4comb), let__5061, tmp5854)


ifres5850 = tmp5855


}

ifres5846 = ifres5850


}

let__5056 := ifres5846
_ = let__5056

tmp6010 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5056)


if True == tmp6010 {
tmp5858 := Call1(__e, SymFn(symshen_4_5rcurly_6), V2680)


let__5063 := tmp5858
_ = let__5063

tmp5870 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5063)


var ifres5859 Obj

if True == tmp5870 {
tmp5860 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5859 = tmp5860


} else {
tmp5861 := Call1(__e, SymFn(symshen_4in_1_6), let__5063)


let__5064 := tmp5861
_ = let__5064

tmp5862 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__5064)


let__5065 := tmp5862
_ = let__5065

tmp5869 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5065)


var ifres5863 Obj

if True == tmp5869 {
tmp5864 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5863 = tmp5864


} else {
tmp5865 := Call1(__e, SymFn(symshen_4_5_1out), let__5065)


let__5066 := tmp5865
_ = let__5066

tmp5866 := Call1(__e, SymFn(symshen_4in_1_6), let__5065)


let__5067 := tmp5866
_ = let__5067

tmp5867 := GCons(__e, sym_j, let__5066)

tmp5868 := Call2(__e, SymFn(symshen_4comb), let__5067, tmp5867)


ifres5863 = tmp5868


}

ifres5859 = ifres5863


}

let__5062 := ifres5859
_ = let__5062

tmp6008 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5062)


if True == tmp6008 {
tmp5871 := Call1(__e, SymFn(symshen_4_5bar_6), V2680)


let__5069 := tmp5871
_ = let__5069

tmp5883 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5069)


var ifres5872 Obj

if True == tmp5883 {
tmp5873 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5872 = tmp5873


} else {
tmp5874 := Call1(__e, SymFn(symshen_4in_1_6), let__5069)


let__5070 := tmp5874
_ = let__5070

tmp5875 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__5070)


let__5071 := tmp5875
_ = let__5071

tmp5882 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5071)


var ifres5876 Obj

if True == tmp5882 {
tmp5877 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5876 = tmp5877


} else {
tmp5878 := Call1(__e, SymFn(symshen_4_5_1out), let__5071)


let__5072 := tmp5878
_ = let__5072

tmp5879 := Call1(__e, SymFn(symshen_4in_1_6), let__5071)


let__5073 := tmp5879
_ = let__5073

tmp5880 := GCons(__e, symbar_b, let__5072)

tmp5881 := Call2(__e, SymFn(symshen_4comb), let__5073, tmp5880)


ifres5876 = tmp5881


}

ifres5872 = ifres5876


}

let__5068 := ifres5872
_ = let__5068

tmp6006 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5068)


if True == tmp6006 {
tmp5884 := Call1(__e, SymFn(symshen_4_5semicolon_6), V2680)


let__5075 := tmp5884
_ = let__5075

tmp5897 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5075)


var ifres5885 Obj

if True == tmp5897 {
tmp5886 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5885 = tmp5886


} else {
tmp5887 := Call1(__e, SymFn(symshen_4in_1_6), let__5075)


let__5076 := tmp5887
_ = let__5076

tmp5888 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__5076)


let__5077 := tmp5888
_ = let__5077

tmp5896 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5077)


var ifres5889 Obj

if True == tmp5896 {
tmp5890 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5889 = tmp5890


} else {
tmp5891 := Call1(__e, SymFn(symshen_4_5_1out), let__5077)


let__5078 := tmp5891
_ = let__5078

tmp5892 := Call1(__e, SymFn(symshen_4in_1_6), let__5077)


let__5079 := tmp5892
_ = let__5079

tmp5893 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp5894 := GCons(__e, tmp5893, let__5078)

tmp5895 := Call2(__e, SymFn(symshen_4comb), let__5079, tmp5894)


ifres5889 = tmp5895


}

ifres5885 = ifres5889


}

let__5074 := ifres5885
_ = let__5074

tmp6004 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5074)


if True == tmp6004 {
tmp5898 := Call1(__e, SymFn(symshen_4_5colon_6), V2680)


let__5081 := tmp5898
_ = let__5081

tmp5916 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5081)


var ifres5899 Obj

if True == tmp5916 {
tmp5900 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5899 = tmp5900


} else {
tmp5901 := Call1(__e, SymFn(symshen_4in_1_6), let__5081)


let__5082 := tmp5901
_ = let__5082

tmp5902 := Call1(__e, SymFn(symshen_4_5equal_6), let__5082)


let__5083 := tmp5902
_ = let__5083

tmp5915 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5083)


var ifres5903 Obj

if True == tmp5915 {
tmp5904 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5903 = tmp5904


} else {
tmp5905 := Call1(__e, SymFn(symshen_4in_1_6), let__5083)


let__5084 := tmp5905
_ = let__5084

tmp5906 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__5084)


let__5085 := tmp5906
_ = let__5085

tmp5914 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5085)


var ifres5907 Obj

if True == tmp5914 {
tmp5908 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5907 = tmp5908


} else {
tmp5909 := Call1(__e, SymFn(symshen_4_5_1out), let__5085)


let__5086 := tmp5909
_ = let__5086

tmp5910 := Call1(__e, SymFn(symshen_4in_1_6), let__5085)


let__5087 := tmp5910
_ = let__5087

tmp5911 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":="))

tmp5912 := GCons(__e, tmp5911, let__5086)

tmp5913 := Call2(__e, SymFn(symshen_4comb), let__5087, tmp5912)


ifres5907 = tmp5913


}

ifres5903 = ifres5907


}

ifres5899 = ifres5903


}

let__5080 := ifres5899
_ = let__5080

tmp6002 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5080)


if True == tmp6002 {
tmp5917 := Call1(__e, SymFn(symshen_4_5colon_6), V2680)


let__5089 := tmp5917
_ = let__5089

tmp5930 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5089)


var ifres5918 Obj

if True == tmp5930 {
tmp5919 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5918 = tmp5919


} else {
tmp5920 := Call1(__e, SymFn(symshen_4in_1_6), let__5089)


let__5090 := tmp5920
_ = let__5090

tmp5921 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__5090)


let__5091 := tmp5921
_ = let__5091

tmp5929 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5091)


var ifres5922 Obj

if True == tmp5929 {
tmp5923 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5922 = tmp5923


} else {
tmp5924 := Call1(__e, SymFn(symshen_4_5_1out), let__5091)


let__5092 := tmp5924
_ = let__5092

tmp5925 := Call1(__e, SymFn(symshen_4in_1_6), let__5091)


let__5093 := tmp5925
_ = let__5093

tmp5926 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp5927 := GCons(__e, tmp5926, let__5092)

tmp5928 := Call2(__e, SymFn(symshen_4comb), let__5093, tmp5927)


ifres5922 = tmp5928


}

ifres5918 = ifres5922


}

let__5088 := ifres5918
_ = let__5088

tmp6000 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5088)


if True == tmp6000 {
tmp5931 := Call1(__e, SymFn(symshen_4_5comma_6), V2680)


let__5095 := tmp5931
_ = let__5095

tmp5944 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5095)


var ifres5932 Obj

if True == tmp5944 {
tmp5933 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5932 = tmp5933


} else {
tmp5934 := Call1(__e, SymFn(symshen_4in_1_6), let__5095)


let__5096 := tmp5934
_ = let__5096

tmp5935 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__5096)


let__5097 := tmp5935
_ = let__5097

tmp5943 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5097)


var ifres5936 Obj

if True == tmp5943 {
tmp5937 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5936 = tmp5937


} else {
tmp5938 := Call1(__e, SymFn(symshen_4_5_1out), let__5097)


let__5098 := tmp5938
_ = let__5098

tmp5939 := Call1(__e, SymFn(symshen_4in_1_6), let__5097)


let__5099 := tmp5939
_ = let__5099

tmp5940 := PrimGuard1(__e, symintern, PrimIntern, MakeString(","))

tmp5941 := GCons(__e, tmp5940, let__5098)

tmp5942 := Call2(__e, SymFn(symshen_4comb), let__5099, tmp5941)


ifres5936 = tmp5942


}

ifres5932 = ifres5936


}

let__5094 := ifres5932
_ = let__5094

tmp5998 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5094)


if True == tmp5998 {
tmp5945 := Call1(__e, SymFn(symshen_4_5comment_6), V2680)


let__5101 := tmp5945
_ = let__5101

tmp5956 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5101)


var ifres5946 Obj

if True == tmp5956 {
tmp5947 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5946 = tmp5947


} else {
tmp5948 := Call1(__e, SymFn(symshen_4in_1_6), let__5101)


let__5102 := tmp5948
_ = let__5102

tmp5949 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__5102)


let__5103 := tmp5949
_ = let__5103

tmp5955 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5103)


var ifres5950 Obj

if True == tmp5955 {
tmp5951 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5950 = tmp5951


} else {
tmp5952 := Call1(__e, SymFn(symshen_4_5_1out), let__5103)


let__5104 := tmp5952
_ = let__5104

tmp5953 := Call1(__e, SymFn(symshen_4in_1_6), let__5103)


let__5105 := tmp5953
_ = let__5105

tmp5954 := Call2(__e, SymFn(symshen_4comb), let__5105, let__5104)


ifres5950 = tmp5954


}

ifres5946 = ifres5950


}

let__5100 := ifres5946
_ = let__5100

tmp5996 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5100)


if True == tmp5996 {
tmp5957 := Call1(__e, SymFn(symshen_4_5atom_6), V2680)


let__5107 := tmp5957
_ = let__5107

tmp5970 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5107)


var ifres5958 Obj

if True == tmp5970 {
tmp5959 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5958 = tmp5959


} else {
tmp5960 := Call1(__e, SymFn(symshen_4_5_1out), let__5107)


let__5108 := tmp5960
_ = let__5108

tmp5961 := Call1(__e, SymFn(symshen_4in_1_6), let__5107)


let__5109 := tmp5961
_ = let__5109

tmp5962 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__5109)


let__5110 := tmp5962
_ = let__5110

tmp5969 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5110)


var ifres5963 Obj

if True == tmp5969 {
tmp5964 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5963 = tmp5964


} else {
tmp5965 := Call1(__e, SymFn(symshen_4_5_1out), let__5110)


let__5111 := tmp5965
_ = let__5111

tmp5966 := Call1(__e, SymFn(symshen_4in_1_6), let__5110)


let__5112 := tmp5966
_ = let__5112

tmp5967 := GCons(__e, let__5108, let__5111)

tmp5968 := Call2(__e, SymFn(symshen_4comb), let__5112, tmp5967)


ifres5963 = tmp5968


}

ifres5958 = ifres5963


}

let__5106 := ifres5958
_ = let__5106

tmp5994 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5106)


if True == tmp5994 {
tmp5971 := Call1(__e, SymFn(symshen_4_5whitespaces_6), V2680)


let__5114 := tmp5971
_ = let__5114

tmp5982 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5114)


var ifres5972 Obj

if True == tmp5982 {
tmp5973 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5972 = tmp5973


} else {
tmp5974 := Call1(__e, SymFn(symshen_4in_1_6), let__5114)


let__5115 := tmp5974
_ = let__5115

tmp5975 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), let__5115)


let__5116 := tmp5975
_ = let__5116

tmp5981 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5116)


var ifres5976 Obj

if True == tmp5981 {
tmp5977 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5976 = tmp5977


} else {
tmp5978 := Call1(__e, SymFn(symshen_4_5_1out), let__5116)


let__5117 := tmp5978
_ = let__5117

tmp5979 := Call1(__e, SymFn(symshen_4in_1_6), let__5116)


let__5118 := tmp5979
_ = let__5118

tmp5980 := Call2(__e, SymFn(symshen_4comb), let__5118, let__5117)


ifres5976 = tmp5980


}

ifres5972 = ifres5976


}

let__5113 := ifres5972
_ = let__5113

tmp5992 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5113)


if True == tmp5992 {
tmp5983 := Call1(__e, SymFn(sym_5e_6), V2680)


let__5120 := tmp5983
_ = let__5120

tmp5988 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5120)


var ifres5984 Obj

if True == tmp5988 {
tmp5985 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres5984 = tmp5985


} else {
tmp5986 := Call1(__e, SymFn(symshen_4in_1_6), let__5120)


let__5121 := tmp5986
_ = let__5121

tmp5987 := Call2(__e, SymFn(symshen_4comb), let__5121, Nil)


ifres5984 = tmp5987


}

let__5119 := ifres5984
_ = let__5119

tmp5990 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5119)


if True == tmp5990 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5119
}


} else {
return let__5113
}


} else {
return let__5106
}


} else {
return let__5100
}


} else {
return let__5094
}


} else {
return let__5088
}


} else {
return let__5080
}


} else {
return let__5074
}


} else {
return let__5068
}


} else {
return let__5062
}


} else {
return let__5056
}


} else {
return let__5045
}


} else {
return let__5034
}


panic("unreachable")
})

tmp6015 := Call2(__e, ns2_1set, symshen_4_5s_1exprs_6, tmp5795)


_ = tmp6015

tmp6016 := MakeNative2(func(__e *ControlFlow, V2769, V2770 Obj) Obj {
tmp6034 := GConsP(__e, V2769)

var ifres6021 Obj

if True == tmp6034 {
tmp6032 := GHd(__e, V2769)

tmp6033 := GEq(__e, sym_3, tmp6032)

var ifres6023 Obj

if True == tmp6033 {
tmp6030 := GTl(__e, V2769)

tmp6031 := GConsP(__e, tmp6030)

var ifres6025 Obj

if True == tmp6031 {
tmp6027 := GTl(__e, V2769)

tmp6028 := GTl(__e, tmp6027)

tmp6029 := GEq(__e, Nil, tmp6028)

var ifres6026 Obj

if True == tmp6029 {
ifres6026 = True


} else {
ifres6026 = False


}

ifres6025 = ifres6026


} else {
ifres6025 = False


}

var ifres6024 Obj

if True == ifres6025 {
ifres6024 = True


} else {
ifres6024 = False


}

ifres6023 = ifres6024


} else {
ifres6023 = False


}

var ifres6022 Obj

if True == ifres6023 {
ifres6022 = True


} else {
ifres6022 = False


}

ifres6021 = ifres6022


} else {
ifres6021 = False


}

if True == ifres6021 {
tmp6017 := GTl(__e, V2769)

tmp6018 := GHd(__e, tmp6017)

tmp6019 := Call1(__e, SymFn(symexplode), tmp6018)


return TailCall2(__e, SymFn(symappend), tmp6019, V2770)


} else {
return GCons(__e, V2769, V2770)
}


panic("unreachable")
})

tmp6035 := Call2(__e, ns2_1set, symshen_4add_1sexpr, tmp6016)


_ = tmp6035

tmp6036 := MakeNative1(func(__e *ControlFlow, V2771 Obj) Obj {
tmp6041 := Call2(__e, SymFn(symshen_4hds_a_2), V2771, MakeInteger(91))


var ifres6037 Obj

if True == tmp6041 {
tmp6038 := Call1(__e, SymFn(symtail), V2771)


let__5123 := tmp6038
_ = let__5123

tmp6039 := Call2(__e, SymFn(symshen_4comb), let__5123, symshen_4skip)


ifres6037 = tmp6039


} else {
tmp6040 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6037 = tmp6040


}

let__5122 := ifres6037
_ = let__5122

tmp6043 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5122)


if True == tmp6043 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5122
}


panic("unreachable")
})

tmp6044 := Call2(__e, ns2_1set, symshen_4_5lsb_6, tmp6036)


_ = tmp6044

tmp6045 := MakeNative1(func(__e *ControlFlow, V2774 Obj) Obj {
tmp6050 := Call2(__e, SymFn(symshen_4hds_a_2), V2774, MakeInteger(93))


var ifres6046 Obj

if True == tmp6050 {
tmp6047 := Call1(__e, SymFn(symtail), V2774)


let__5125 := tmp6047
_ = let__5125

tmp6048 := Call2(__e, SymFn(symshen_4comb), let__5125, symshen_4skip)


ifres6046 = tmp6048


} else {
tmp6049 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6046 = tmp6049


}

let__5124 := ifres6046
_ = let__5124

tmp6052 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5124)


if True == tmp6052 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5124
}


panic("unreachable")
})

tmp6053 := Call2(__e, ns2_1set, symshen_4_5rsb_6, tmp6045)


_ = tmp6053

tmp6054 := MakeNative1(func(__e *ControlFlow, V2777 Obj) Obj {
tmp6055 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), V2777)


let__5127 := tmp6055
_ = let__5127

tmp6061 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5127)


var ifres6056 Obj

if True == tmp6061 {
tmp6057 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6056 = tmp6057


} else {
tmp6058 := Call1(__e, SymFn(symshen_4_5_1out), let__5127)


let__5128 := tmp6058
_ = let__5128

tmp6059 := Call1(__e, SymFn(symshen_4in_1_6), let__5127)


let__5129 := tmp6059
_ = let__5129

tmp6060 := Call2(__e, SymFn(symshen_4comb), let__5129, let__5128)


ifres6056 = tmp6060


}

let__5126 := ifres6056
_ = let__5126

tmp6063 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5126)


if True == tmp6063 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5126
}


panic("unreachable")
})

tmp6064 := Call2(__e, ns2_1set, symshen_4_5s_1exprs1_6, tmp6054)


_ = tmp6064

tmp6065 := MakeNative1(func(__e *ControlFlow, V2782 Obj) Obj {
tmp6066 := Call1(__e, SymFn(symshen_4_5s_1exprs_6), V2782)


let__5131 := tmp6066
_ = let__5131

tmp6072 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5131)


var ifres6067 Obj

if True == tmp6072 {
tmp6068 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6067 = tmp6068


} else {
tmp6069 := Call1(__e, SymFn(symshen_4_5_1out), let__5131)


let__5132 := tmp6069
_ = let__5132

tmp6070 := Call1(__e, SymFn(symshen_4in_1_6), let__5131)


let__5133 := tmp6070
_ = let__5133

tmp6071 := Call2(__e, SymFn(symshen_4comb), let__5133, let__5132)


ifres6067 = tmp6071


}

let__5130 := ifres6067
_ = let__5130

tmp6074 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5130)


if True == tmp6074 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5130
}


panic("unreachable")
})

tmp6075 := Call2(__e, ns2_1set, symshen_4_5s_1exprs2_6, tmp6065)


_ = tmp6075

tmp6076 := MakeNative1(func(__e *ControlFlow, V2788 Obj) Obj {
tmp6133 := GEq(__e, Nil, V2788)

if True == tmp6133 {
return Nil
} else {
tmp6131 := GConsP(__e, V2788)

var ifres6111 Obj

if True == tmp6131 {
tmp6129 := GTl(__e, V2788)

tmp6130 := GConsP(__e, tmp6129)

var ifres6113 Obj

if True == tmp6130 {
tmp6126 := GTl(__e, V2788)

tmp6127 := GTl(__e, tmp6126)

tmp6128 := GConsP(__e, tmp6127)

var ifres6115 Obj

if True == tmp6128 {
tmp6122 := GTl(__e, V2788)

tmp6123 := GTl(__e, tmp6122)

tmp6124 := GTl(__e, tmp6123)

tmp6125 := GEq(__e, Nil, tmp6124)

var ifres6117 Obj

if True == tmp6125 {
tmp6119 := GTl(__e, V2788)

tmp6120 := GHd(__e, tmp6119)

tmp6121 := GEq(__e, tmp6120, symbar_b)

var ifres6118 Obj

if True == tmp6121 {
ifres6118 = True


} else {
ifres6118 = False


}

ifres6117 = ifres6118


} else {
ifres6117 = False


}

var ifres6116 Obj

if True == ifres6117 {
ifres6116 = True


} else {
ifres6116 = False


}

ifres6115 = ifres6116


} else {
ifres6115 = False


}

var ifres6114 Obj

if True == ifres6115 {
ifres6114 = True


} else {
ifres6114 = False


}

ifres6113 = ifres6114


} else {
ifres6113 = False


}

var ifres6112 Obj

if True == ifres6113 {
ifres6112 = True


} else {
ifres6112 = False


}

ifres6111 = ifres6112


} else {
ifres6111 = False


}

if True == ifres6111 {
tmp6077 := GHd(__e, V2788)

tmp6078 := GTl(__e, V2788)

tmp6079 := GTl(__e, tmp6078)

tmp6080 := GCons(__e, tmp6077, tmp6079)

return GCons(__e, symcons, tmp6080)


} else {
tmp6109 := GConsP(__e, V2788)

var ifres6089 Obj

if True == tmp6109 {
tmp6107 := GTl(__e, V2788)

tmp6108 := GConsP(__e, tmp6107)

var ifres6091 Obj

if True == tmp6108 {
tmp6104 := GTl(__e, V2788)

tmp6105 := GTl(__e, tmp6104)

tmp6106 := GConsP(__e, tmp6105)

var ifres6093 Obj

if True == tmp6106 {
tmp6100 := GTl(__e, V2788)

tmp6101 := GTl(__e, tmp6100)

tmp6102 := GTl(__e, tmp6101)

tmp6103 := GConsP(__e, tmp6102)

var ifres6095 Obj

if True == tmp6103 {
tmp6097 := GTl(__e, V2788)

tmp6098 := GHd(__e, tmp6097)

tmp6099 := GEq(__e, tmp6098, symbar_b)

var ifres6096 Obj

if True == tmp6099 {
ifres6096 = True


} else {
ifres6096 = False


}

ifres6095 = ifres6096


} else {
ifres6095 = False


}

var ifres6094 Obj

if True == ifres6095 {
ifres6094 = True


} else {
ifres6094 = False


}

ifres6093 = ifres6094


} else {
ifres6093 = False


}

var ifres6092 Obj

if True == ifres6093 {
ifres6092 = True


} else {
ifres6092 = False


}

ifres6091 = ifres6092


} else {
ifres6091 = False


}

var ifres6090 Obj

if True == ifres6091 {
ifres6090 = True


} else {
ifres6090 = False


}

ifres6089 = ifres6090


} else {
ifres6089 = False


}

if True == ifres6089 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("misapplication of |\n"))
} else {
tmp6087 := GConsP(__e, V2788)

if True == tmp6087 {
tmp6081 := GHd(__e, V2788)

tmp6082 := GTl(__e, V2788)

tmp6083 := Call1(__e, SymFn(symshen_4cons_1form), tmp6082)


tmp6084 := GCons(__e, tmp6083, Nil)

tmp6085 := GCons(__e, tmp6081, tmp6084)

return GCons(__e, symcons, tmp6085)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4cons_1form)
}


}


}


}


panic("unreachable")
})

tmp6134 := Call2(__e, ns2_1set, symshen_4cons_1form, tmp6076)


_ = tmp6134

tmp6135 := MakeNative1(func(__e *ControlFlow, V2789 Obj) Obj {
tmp6140 := Call2(__e, SymFn(symshen_4hds_a_2), V2789, MakeInteger(40))


var ifres6136 Obj

if True == tmp6140 {
tmp6137 := Call1(__e, SymFn(symtail), V2789)


let__5135 := tmp6137
_ = let__5135

tmp6138 := Call2(__e, SymFn(symshen_4comb), let__5135, symshen_4skip)


ifres6136 = tmp6138


} else {
tmp6139 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6136 = tmp6139


}

let__5134 := ifres6136
_ = let__5134

tmp6142 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5134)


if True == tmp6142 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5134
}


panic("unreachable")
})

tmp6143 := Call2(__e, ns2_1set, symshen_4_5lrb_6, tmp6135)


_ = tmp6143

tmp6144 := MakeNative1(func(__e *ControlFlow, V2792 Obj) Obj {
tmp6149 := Call2(__e, SymFn(symshen_4hds_a_2), V2792, MakeInteger(41))


var ifres6145 Obj

if True == tmp6149 {
tmp6146 := Call1(__e, SymFn(symtail), V2792)


let__5137 := tmp6146
_ = let__5137

tmp6147 := Call2(__e, SymFn(symshen_4comb), let__5137, symshen_4skip)


ifres6145 = tmp6147


} else {
tmp6148 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6145 = tmp6148


}

let__5136 := ifres6145
_ = let__5136

tmp6151 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5136)


if True == tmp6151 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5136
}


panic("unreachable")
})

tmp6152 := Call2(__e, ns2_1set, symshen_4_5rrb_6, tmp6144)


_ = tmp6152

tmp6153 := MakeNative1(func(__e *ControlFlow, V2795 Obj) Obj {
tmp6158 := Call2(__e, SymFn(symshen_4hds_a_2), V2795, MakeInteger(123))


var ifres6154 Obj

if True == tmp6158 {
tmp6155 := Call1(__e, SymFn(symtail), V2795)


let__5139 := tmp6155
_ = let__5139

tmp6156 := Call2(__e, SymFn(symshen_4comb), let__5139, symshen_4skip)


ifres6154 = tmp6156


} else {
tmp6157 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6154 = tmp6157


}

let__5138 := ifres6154
_ = let__5138

tmp6160 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5138)


if True == tmp6160 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5138
}


panic("unreachable")
})

tmp6161 := Call2(__e, ns2_1set, symshen_4_5lcurly_6, tmp6153)


_ = tmp6161

tmp6162 := MakeNative1(func(__e *ControlFlow, V2798 Obj) Obj {
tmp6167 := Call2(__e, SymFn(symshen_4hds_a_2), V2798, MakeInteger(125))


var ifres6163 Obj

if True == tmp6167 {
tmp6164 := Call1(__e, SymFn(symtail), V2798)


let__5141 := tmp6164
_ = let__5141

tmp6165 := Call2(__e, SymFn(symshen_4comb), let__5141, symshen_4skip)


ifres6163 = tmp6165


} else {
tmp6166 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6163 = tmp6166


}

let__5140 := ifres6163
_ = let__5140

tmp6169 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5140)


if True == tmp6169 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5140
}


panic("unreachable")
})

tmp6170 := Call2(__e, ns2_1set, symshen_4_5rcurly_6, tmp6162)


_ = tmp6170

tmp6171 := MakeNative1(func(__e *ControlFlow, V2801 Obj) Obj {
tmp6176 := Call2(__e, SymFn(symshen_4hds_a_2), V2801, MakeInteger(124))


var ifres6172 Obj

if True == tmp6176 {
tmp6173 := Call1(__e, SymFn(symtail), V2801)


let__5143 := tmp6173
_ = let__5143

tmp6174 := Call2(__e, SymFn(symshen_4comb), let__5143, symshen_4skip)


ifres6172 = tmp6174


} else {
tmp6175 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6172 = tmp6175


}

let__5142 := ifres6172
_ = let__5142

tmp6178 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5142)


if True == tmp6178 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5142
}


panic("unreachable")
})

tmp6179 := Call2(__e, ns2_1set, symshen_4_5bar_6, tmp6171)


_ = tmp6179

tmp6180 := MakeNative1(func(__e *ControlFlow, V2804 Obj) Obj {
tmp6185 := Call2(__e, SymFn(symshen_4hds_a_2), V2804, MakeInteger(59))


var ifres6181 Obj

if True == tmp6185 {
tmp6182 := Call1(__e, SymFn(symtail), V2804)


let__5145 := tmp6182
_ = let__5145

tmp6183 := Call2(__e, SymFn(symshen_4comb), let__5145, symshen_4skip)


ifres6181 = tmp6183


} else {
tmp6184 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6181 = tmp6184


}

let__5144 := ifres6181
_ = let__5144

tmp6187 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5144)


if True == tmp6187 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5144
}


panic("unreachable")
})

tmp6188 := Call2(__e, ns2_1set, symshen_4_5semicolon_6, tmp6180)


_ = tmp6188

tmp6189 := MakeNative1(func(__e *ControlFlow, V2807 Obj) Obj {
tmp6194 := Call2(__e, SymFn(symshen_4hds_a_2), V2807, MakeInteger(58))


var ifres6190 Obj

if True == tmp6194 {
tmp6191 := Call1(__e, SymFn(symtail), V2807)


let__5147 := tmp6191
_ = let__5147

tmp6192 := Call2(__e, SymFn(symshen_4comb), let__5147, symshen_4skip)


ifres6190 = tmp6192


} else {
tmp6193 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6190 = tmp6193


}

let__5146 := ifres6190
_ = let__5146

tmp6196 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5146)


if True == tmp6196 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5146
}


panic("unreachable")
})

tmp6197 := Call2(__e, ns2_1set, symshen_4_5colon_6, tmp6189)


_ = tmp6197

tmp6198 := MakeNative1(func(__e *ControlFlow, V2810 Obj) Obj {
tmp6203 := Call2(__e, SymFn(symshen_4hds_a_2), V2810, MakeInteger(44))


var ifres6199 Obj

if True == tmp6203 {
tmp6200 := Call1(__e, SymFn(symtail), V2810)


let__5149 := tmp6200
_ = let__5149

tmp6201 := Call2(__e, SymFn(symshen_4comb), let__5149, symshen_4skip)


ifres6199 = tmp6201


} else {
tmp6202 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6199 = tmp6202


}

let__5148 := ifres6199
_ = let__5148

tmp6205 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5148)


if True == tmp6205 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5148
}


panic("unreachable")
})

tmp6206 := Call2(__e, ns2_1set, symshen_4_5comma_6, tmp6198)


_ = tmp6206

tmp6207 := MakeNative1(func(__e *ControlFlow, V2813 Obj) Obj {
tmp6212 := Call2(__e, SymFn(symshen_4hds_a_2), V2813, MakeInteger(61))


var ifres6208 Obj

if True == tmp6212 {
tmp6209 := Call1(__e, SymFn(symtail), V2813)


let__5151 := tmp6209
_ = let__5151

tmp6210 := Call2(__e, SymFn(symshen_4comb), let__5151, symshen_4skip)


ifres6208 = tmp6210


} else {
tmp6211 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6208 = tmp6211


}

let__5150 := ifres6208
_ = let__5150

tmp6214 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5150)


if True == tmp6214 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5150
}


panic("unreachable")
})

tmp6215 := Call2(__e, ns2_1set, symshen_4_5equal_6, tmp6207)


_ = tmp6215

tmp6216 := MakeNative1(func(__e *ControlFlow, V2816 Obj) Obj {
tmp6217 := Call1(__e, SymFn(symshen_4_5singleline_6), V2816)


let__5153 := tmp6217
_ = let__5153

tmp6222 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5153)


var ifres6218 Obj

if True == tmp6222 {
tmp6219 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6218 = tmp6219


} else {
tmp6220 := Call1(__e, SymFn(symshen_4in_1_6), let__5153)


let__5154 := tmp6220
_ = let__5154

tmp6221 := Call2(__e, SymFn(symshen_4comb), let__5154, symshen_4skip)


ifres6218 = tmp6221


}

let__5152 := ifres6218
_ = let__5152

tmp6232 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5152)


if True == tmp6232 {
tmp6223 := Call1(__e, SymFn(symshen_4_5multiline_6), V2816)


let__5156 := tmp6223
_ = let__5156

tmp6228 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5156)


var ifres6224 Obj

if True == tmp6228 {
tmp6225 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6224 = tmp6225


} else {
tmp6226 := Call1(__e, SymFn(symshen_4in_1_6), let__5156)


let__5157 := tmp6226
_ = let__5157

tmp6227 := Call2(__e, SymFn(symshen_4comb), let__5157, symshen_4skip)


ifres6224 = tmp6227


}

let__5155 := ifres6224
_ = let__5155

tmp6230 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5155)


if True == tmp6230 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5155
}


} else {
return let__5152
}


panic("unreachable")
})

tmp6233 := Call2(__e, ns2_1set, symshen_4_5comment_6, tmp6216)


_ = tmp6233

tmp6234 := MakeNative1(func(__e *ControlFlow, V2823 Obj) Obj {
tmp6235 := Call1(__e, SymFn(symshen_4_5backslash_6), V2823)


let__5159 := tmp6235
_ = let__5159

tmp6255 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5159)


var ifres6236 Obj

if True == tmp6255 {
tmp6237 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6236 = tmp6237


} else {
tmp6238 := Call1(__e, SymFn(symshen_4in_1_6), let__5159)


let__5160 := tmp6238
_ = let__5160

tmp6239 := Call1(__e, SymFn(symshen_4_5backslash_6), let__5160)


let__5161 := tmp6239
_ = let__5161

tmp6254 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5161)


var ifres6240 Obj

if True == tmp6254 {
tmp6241 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6240 = tmp6241


} else {
tmp6242 := Call1(__e, SymFn(symshen_4in_1_6), let__5161)


let__5162 := tmp6242
_ = let__5162

tmp6243 := Call1(__e, SymFn(symshen_4_5shortnatters_6), let__5162)


let__5163 := tmp6243
_ = let__5163

tmp6253 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5163)


var ifres6244 Obj

if True == tmp6253 {
tmp6245 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6244 = tmp6245


} else {
tmp6246 := Call1(__e, SymFn(symshen_4in_1_6), let__5163)


let__5164 := tmp6246
_ = let__5164

tmp6247 := Call1(__e, SymFn(symshen_4_5returns_6), let__5164)


let__5165 := tmp6247
_ = let__5165

tmp6252 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5165)


var ifres6248 Obj

if True == tmp6252 {
tmp6249 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6248 = tmp6249


} else {
tmp6250 := Call1(__e, SymFn(symshen_4in_1_6), let__5165)


let__5166 := tmp6250
_ = let__5166

tmp6251 := Call2(__e, SymFn(symshen_4comb), let__5166, symshen_4skip)


ifres6248 = tmp6251


}

ifres6244 = ifres6248


}

ifres6240 = ifres6244


}

ifres6236 = ifres6240


}

let__5158 := ifres6236
_ = let__5158

tmp6257 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5158)


if True == tmp6257 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5158
}


panic("unreachable")
})

tmp6258 := Call2(__e, ns2_1set, symshen_4_5singleline_6, tmp6234)


_ = tmp6258

tmp6259 := MakeNative1(func(__e *ControlFlow, V2833 Obj) Obj {
tmp6264 := Call2(__e, SymFn(symshen_4hds_a_2), V2833, MakeInteger(92))


var ifres6260 Obj

if True == tmp6264 {
tmp6261 := Call1(__e, SymFn(symtail), V2833)


let__5168 := tmp6261
_ = let__5168

tmp6262 := Call2(__e, SymFn(symshen_4comb), let__5168, symshen_4skip)


ifres6260 = tmp6262


} else {
tmp6263 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6260 = tmp6263


}

let__5167 := ifres6260
_ = let__5167

tmp6266 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5167)


if True == tmp6266 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5167
}


panic("unreachable")
})

tmp6267 := Call2(__e, ns2_1set, symshen_4_5backslash_6, tmp6259)


_ = tmp6267

tmp6268 := MakeNative1(func(__e *ControlFlow, V2836 Obj) Obj {
tmp6269 := Call1(__e, SymFn(symshen_4_5shortnatter_6), V2836)


let__5170 := tmp6269
_ = let__5170

tmp6279 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5170)


var ifres6270 Obj

if True == tmp6279 {
tmp6271 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6270 = tmp6271


} else {
tmp6272 := Call1(__e, SymFn(symshen_4in_1_6), let__5170)


let__5171 := tmp6272
_ = let__5171

tmp6273 := Call1(__e, SymFn(symshen_4_5shortnatters_6), let__5171)


let__5172 := tmp6273
_ = let__5172

tmp6278 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5172)


var ifres6274 Obj

if True == tmp6278 {
tmp6275 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6274 = tmp6275


} else {
tmp6276 := Call1(__e, SymFn(symshen_4in_1_6), let__5172)


let__5173 := tmp6276
_ = let__5173

tmp6277 := Call2(__e, SymFn(symshen_4comb), let__5173, symshen_4skip)


ifres6274 = tmp6277


}

ifres6270 = ifres6274


}

let__5169 := ifres6270
_ = let__5169

tmp6289 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5169)


if True == tmp6289 {
tmp6280 := Call1(__e, SymFn(sym_5e_6), V2836)


let__5175 := tmp6280
_ = let__5175

tmp6285 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5175)


var ifres6281 Obj

if True == tmp6285 {
tmp6282 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6281 = tmp6282


} else {
tmp6283 := Call1(__e, SymFn(symshen_4in_1_6), let__5175)


let__5176 := tmp6283
_ = let__5176

tmp6284 := Call2(__e, SymFn(symshen_4comb), let__5176, symshen_4skip)


ifres6281 = tmp6284


}

let__5174 := ifres6281
_ = let__5174

tmp6287 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5174)


if True == tmp6287 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5174
}


} else {
return let__5169
}


panic("unreachable")
})

tmp6290 := Call2(__e, ns2_1set, symshen_4_5shortnatters_6, tmp6268)


_ = tmp6290

tmp6291 := MakeNative1(func(__e *ControlFlow, V2845 Obj) Obj {
tmp6301 := GConsP(__e, V2845)

var ifres6292 Obj

if True == tmp6301 {
tmp6293 := Call1(__e, SymFn(symhead), V2845)


let__5178 := tmp6293
_ = let__5178

tmp6294 := Call1(__e, SymFn(symtail), V2845)


let__5179 := tmp6294
_ = let__5179

tmp6298 := Call1(__e, SymFn(symshen_4return_2), let__5178)


tmp6299 := GNot(__e, tmp6298)

var ifres6295 Obj

if True == tmp6299 {
tmp6296 := Call2(__e, SymFn(symshen_4comb), let__5179, symshen_4skip)


ifres6295 = tmp6296


} else {
tmp6297 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6295 = tmp6297


}

ifres6292 = ifres6295


} else {
tmp6300 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6292 = tmp6300


}

let__5177 := ifres6292
_ = let__5177

tmp6303 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5177)


if True == tmp6303 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5177
}


panic("unreachable")
})

tmp6304 := Call2(__e, ns2_1set, symshen_4_5shortnatter_6, tmp6291)


_ = tmp6304

tmp6305 := MakeNative1(func(__e *ControlFlow, V2849 Obj) Obj {
tmp6306 := Call1(__e, SymFn(symshen_4_5return_6), V2849)


let__5181 := tmp6306
_ = let__5181

tmp6316 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5181)


var ifres6307 Obj

if True == tmp6316 {
tmp6308 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6307 = tmp6308


} else {
tmp6309 := Call1(__e, SymFn(symshen_4in_1_6), let__5181)


let__5182 := tmp6309
_ = let__5182

tmp6310 := Call1(__e, SymFn(symshen_4_5returns_6), let__5182)


let__5183 := tmp6310
_ = let__5183

tmp6315 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5183)


var ifres6311 Obj

if True == tmp6315 {
tmp6312 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6311 = tmp6312


} else {
tmp6313 := Call1(__e, SymFn(symshen_4in_1_6), let__5183)


let__5184 := tmp6313
_ = let__5184

tmp6314 := Call2(__e, SymFn(symshen_4comb), let__5184, symshen_4skip)


ifres6311 = tmp6314


}

ifres6307 = ifres6311


}

let__5180 := ifres6307
_ = let__5180

tmp6326 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5180)


if True == tmp6326 {
tmp6317 := Call1(__e, SymFn(symshen_4_5return_6), V2849)


let__5186 := tmp6317
_ = let__5186

tmp6322 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5186)


var ifres6318 Obj

if True == tmp6322 {
tmp6319 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6318 = tmp6319


} else {
tmp6320 := Call1(__e, SymFn(symshen_4in_1_6), let__5186)


let__5187 := tmp6320
_ = let__5187

tmp6321 := Call2(__e, SymFn(symshen_4comb), let__5187, symshen_4skip)


ifres6318 = tmp6321


}

let__5185 := ifres6318
_ = let__5185

tmp6324 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5185)


if True == tmp6324 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5185
}


} else {
return let__5180
}


panic("unreachable")
})

tmp6327 := Call2(__e, ns2_1set, symshen_4_5returns_6, tmp6305)


_ = tmp6327

tmp6328 := MakeNative1(func(__e *ControlFlow, V2858 Obj) Obj {
tmp6337 := GConsP(__e, V2858)

var ifres6329 Obj

if True == tmp6337 {
tmp6330 := Call1(__e, SymFn(symhead), V2858)


let__5189 := tmp6330
_ = let__5189

tmp6331 := Call1(__e, SymFn(symtail), V2858)


let__5190 := tmp6331
_ = let__5190

tmp6335 := Call1(__e, SymFn(symshen_4return_2), let__5189)


var ifres6332 Obj

if True == tmp6335 {
tmp6333 := Call2(__e, SymFn(symshen_4comb), let__5190, symshen_4skip)


ifres6332 = tmp6333


} else {
tmp6334 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6332 = tmp6334


}

ifres6329 = ifres6332


} else {
tmp6336 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6329 = tmp6336


}

let__5188 := ifres6329
_ = let__5188

tmp6339 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5188)


if True == tmp6339 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5188
}


panic("unreachable")
})

tmp6340 := Call2(__e, ns2_1set, symshen_4_5return_6, tmp6328)


_ = tmp6340

tmp6341 := MakeNative1(func(__e *ControlFlow, V2862 Obj) Obj {
tmp6342 := GCons(__e, MakeInteger(13), Nil)

tmp6343 := GCons(__e, MakeInteger(10), tmp6342)

tmp6344 := GCons(__e, MakeInteger(9), tmp6343)

return TailCall2(__e, SymFn(symelement_2), V2862, tmp6344)


panic("unreachable")
})

tmp6345 := Call2(__e, ns2_1set, symshen_4return_2, tmp6341)


_ = tmp6345

tmp6346 := MakeNative1(func(__e *ControlFlow, V2863 Obj) Obj {
tmp6347 := Call1(__e, SymFn(symshen_4_5backslash_6), V2863)


let__5192 := tmp6347
_ = let__5192

tmp6362 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5192)


var ifres6348 Obj

if True == tmp6362 {
tmp6349 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6348 = tmp6349


} else {
tmp6350 := Call1(__e, SymFn(symshen_4in_1_6), let__5192)


let__5193 := tmp6350
_ = let__5193

tmp6351 := Call1(__e, SymFn(symshen_4_5times_6), let__5193)


let__5194 := tmp6351
_ = let__5194

tmp6361 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5194)


var ifres6352 Obj

if True == tmp6361 {
tmp6353 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6352 = tmp6353


} else {
tmp6354 := Call1(__e, SymFn(symshen_4in_1_6), let__5194)


let__5195 := tmp6354
_ = let__5195

tmp6355 := Call1(__e, SymFn(symshen_4_5longnatter_6), let__5195)


let__5196 := tmp6355
_ = let__5196

tmp6360 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5196)


var ifres6356 Obj

if True == tmp6360 {
tmp6357 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6356 = tmp6357


} else {
tmp6358 := Call1(__e, SymFn(symshen_4in_1_6), let__5196)


let__5197 := tmp6358
_ = let__5197

tmp6359 := Call2(__e, SymFn(symshen_4comb), let__5197, symshen_4skip)


ifres6356 = tmp6359


}

ifres6352 = ifres6356


}

ifres6348 = ifres6352


}

let__5191 := ifres6348
_ = let__5191

tmp6364 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5191)


if True == tmp6364 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5191
}


panic("unreachable")
})

tmp6365 := Call2(__e, ns2_1set, symshen_4_5multiline_6, tmp6346)


_ = tmp6365

tmp6366 := MakeNative1(func(__e *ControlFlow, V2871 Obj) Obj {
tmp6371 := Call2(__e, SymFn(symshen_4hds_a_2), V2871, MakeInteger(42))


var ifres6367 Obj

if True == tmp6371 {
tmp6368 := Call1(__e, SymFn(symtail), V2871)


let__5199 := tmp6368
_ = let__5199

tmp6369 := Call2(__e, SymFn(symshen_4comb), let__5199, symshen_4skip)


ifres6367 = tmp6369


} else {
tmp6370 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6367 = tmp6370


}

let__5198 := ifres6367
_ = let__5198

tmp6373 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5198)


if True == tmp6373 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5198
}


panic("unreachable")
})

tmp6374 := Call2(__e, ns2_1set, symshen_4_5times_6, tmp6366)


_ = tmp6374

tmp6375 := MakeNative1(func(__e *ControlFlow, V2874 Obj) Obj {
tmp6376 := Call1(__e, SymFn(symshen_4_5comment_6), V2874)


let__5201 := tmp6376
_ = let__5201

tmp6386 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5201)


var ifres6377 Obj

if True == tmp6386 {
tmp6378 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6377 = tmp6378


} else {
tmp6379 := Call1(__e, SymFn(symshen_4in_1_6), let__5201)


let__5202 := tmp6379
_ = let__5202

tmp6380 := Call1(__e, SymFn(symshen_4_5longnatter_6), let__5202)


let__5203 := tmp6380
_ = let__5203

tmp6385 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5203)


var ifres6381 Obj

if True == tmp6385 {
tmp6382 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6381 = tmp6382


} else {
tmp6383 := Call1(__e, SymFn(symshen_4in_1_6), let__5203)


let__5204 := tmp6383
_ = let__5204

tmp6384 := Call2(__e, SymFn(symshen_4comb), let__5204, symshen_4skip)


ifres6381 = tmp6384


}

ifres6377 = ifres6381


}

let__5200 := ifres6377
_ = let__5200

tmp6413 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5200)


if True == tmp6413 {
tmp6387 := Call1(__e, SymFn(symshen_4_5times_6), V2874)


let__5206 := tmp6387
_ = let__5206

tmp6397 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5206)


var ifres6388 Obj

if True == tmp6397 {
tmp6389 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6388 = tmp6389


} else {
tmp6390 := Call1(__e, SymFn(symshen_4in_1_6), let__5206)


let__5207 := tmp6390
_ = let__5207

tmp6391 := Call1(__e, SymFn(symshen_4_5backslash_6), let__5207)


let__5208 := tmp6391
_ = let__5208

tmp6396 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5208)


var ifres6392 Obj

if True == tmp6396 {
tmp6393 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6392 = tmp6393


} else {
tmp6394 := Call1(__e, SymFn(symshen_4in_1_6), let__5208)


let__5209 := tmp6394
_ = let__5209

tmp6395 := Call2(__e, SymFn(symshen_4comb), let__5209, symshen_4skip)


ifres6392 = tmp6395


}

ifres6388 = ifres6392


}

let__5205 := ifres6388
_ = let__5205

tmp6411 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5205)


if True == tmp6411 {
tmp6407 := GConsP(__e, V2874)

var ifres6398 Obj

if True == tmp6407 {
tmp6399 := Call1(__e, SymFn(symtail), V2874)


let__5211 := tmp6399
_ = let__5211

tmp6400 := Call1(__e, SymFn(symshen_4_5longnatter_6), let__5211)


let__5212 := tmp6400
_ = let__5212

tmp6405 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5212)


var ifres6401 Obj

if True == tmp6405 {
tmp6402 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6401 = tmp6402


} else {
tmp6403 := Call1(__e, SymFn(symshen_4in_1_6), let__5212)


let__5213 := tmp6403
_ = let__5213

tmp6404 := Call2(__e, SymFn(symshen_4comb), let__5213, symshen_4skip)


ifres6401 = tmp6404


}

ifres6398 = ifres6401


} else {
tmp6406 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6398 = tmp6406


}

let__5210 := ifres6398
_ = let__5210

tmp6409 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5210)


if True == tmp6409 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5210
}


} else {
return let__5205
}


} else {
return let__5200
}


panic("unreachable")
})

tmp6414 := Call2(__e, ns2_1set, symshen_4_5longnatter_6, tmp6375)


_ = tmp6414

tmp6415 := MakeNative1(func(__e *ControlFlow, V2889 Obj) Obj {
tmp6416 := Call1(__e, SymFn(symshen_4_5str_6), V2889)


let__5215 := tmp6416
_ = let__5215

tmp6422 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5215)


var ifres6417 Obj

if True == tmp6422 {
tmp6418 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6417 = tmp6418


} else {
tmp6419 := Call1(__e, SymFn(symshen_4_5_1out), let__5215)


let__5216 := tmp6419
_ = let__5216

tmp6420 := Call1(__e, SymFn(symshen_4in_1_6), let__5215)


let__5217 := tmp6420
_ = let__5217

tmp6421 := Call2(__e, SymFn(symshen_4comb), let__5217, let__5216)


ifres6417 = tmp6421


}

let__5214 := ifres6417
_ = let__5214

tmp6447 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5214)


if True == tmp6447 {
tmp6423 := Call1(__e, SymFn(symshen_4_5number_6), V2889)


let__5219 := tmp6423
_ = let__5219

tmp6429 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5219)


var ifres6424 Obj

if True == tmp6429 {
tmp6425 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6424 = tmp6425


} else {
tmp6426 := Call1(__e, SymFn(symshen_4_5_1out), let__5219)


let__5220 := tmp6426
_ = let__5220

tmp6427 := Call1(__e, SymFn(symshen_4in_1_6), let__5219)


let__5221 := tmp6427
_ = let__5221

tmp6428 := Call2(__e, SymFn(symshen_4comb), let__5221, let__5220)


ifres6424 = tmp6428


}

let__5218 := ifres6424
_ = let__5218

tmp6445 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5218)


if True == tmp6445 {
tmp6430 := Call1(__e, SymFn(symshen_4_5sym_6), V2889)


let__5223 := tmp6430
_ = let__5223

tmp6441 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5223)


var ifres6431 Obj

if True == tmp6441 {
tmp6432 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6431 = tmp6432


} else {
tmp6433 := Call1(__e, SymFn(symshen_4_5_1out), let__5223)


let__5224 := tmp6433
_ = let__5224

tmp6434 := Call1(__e, SymFn(symshen_4in_1_6), let__5223)


let__5225 := tmp6434
_ = let__5225

tmp6439 := GEq(__e, let__5224, MakeString("<>"))

var ifres6435 Obj

if True == tmp6439 {
tmp6436 := GCons(__e, MakeInteger(0), Nil)

tmp6437 := GCons(__e, symvector, tmp6436)

ifres6435 = tmp6437


} else {
tmp6438 := PrimGuard1(__e, symintern, PrimIntern, let__5224)

ifres6435 = tmp6438


}

tmp6440 := Call2(__e, SymFn(symshen_4comb), let__5225, ifres6435)


ifres6431 = tmp6440


}

let__5222 := ifres6431
_ = let__5222

tmp6443 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5222)


if True == tmp6443 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5222
}


} else {
return let__5218
}


} else {
return let__5214
}


panic("unreachable")
})

tmp6448 := Call2(__e, ns2_1set, symshen_4_5atom_6, tmp6415)


_ = tmp6448

tmp6449 := MakeNative1(func(__e *ControlFlow, V2902 Obj) Obj {
tmp6450 := Call1(__e, SymFn(symshen_4_5alpha_6), V2902)


let__5227 := tmp6450
_ = let__5227

tmp6463 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5227)


var ifres6451 Obj

if True == tmp6463 {
tmp6452 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6451 = tmp6452


} else {
tmp6453 := Call1(__e, SymFn(symshen_4_5_1out), let__5227)


let__5228 := tmp6453
_ = let__5228

tmp6454 := Call1(__e, SymFn(symshen_4in_1_6), let__5227)


let__5229 := tmp6454
_ = let__5229

tmp6455 := Call1(__e, SymFn(symshen_4_5alphanums_6), let__5229)


let__5230 := tmp6455
_ = let__5230

tmp6462 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5230)


var ifres6456 Obj

if True == tmp6462 {
tmp6457 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6456 = tmp6457


} else {
tmp6458 := Call1(__e, SymFn(symshen_4_5_1out), let__5230)


let__5231 := tmp6458
_ = let__5231

tmp6459 := Call1(__e, SymFn(symshen_4in_1_6), let__5230)


let__5232 := tmp6459
_ = let__5232

tmp6461 := Call2(__e, SymFn(symshen_4comb), let__5232, PrimGuard2(__e, symcn, PrimStringConcat, let__5228, let__5231))


ifres6456 = tmp6461


}

ifres6451 = ifres6456


}

let__5226 := ifres6451
_ = let__5226

tmp6465 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5226)


if True == tmp6465 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5226
}


panic("unreachable")
})

tmp6466 := Call2(__e, ns2_1set, symshen_4_5sym_6, tmp6449)


_ = tmp6466

tmp6467 := MakeNative1(func(__e *ControlFlow, V2910 Obj) Obj {
tmp6477 := GConsP(__e, V2910)

var ifres6468 Obj

if True == tmp6477 {
tmp6469 := Call1(__e, SymFn(symhead), V2910)


let__5234 := tmp6469
_ = let__5234

tmp6470 := Call1(__e, SymFn(symtail), V2910)


let__5235 := tmp6470
_ = let__5235

tmp6475 := Call1(__e, SymFn(symshen_4alpha_2), let__5234)


var ifres6471 Obj

if True == tmp6475 {
tmp6472 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, let__5234)

tmp6473 := Call2(__e, SymFn(symshen_4comb), let__5235, tmp6472)


ifres6471 = tmp6473


} else {
tmp6474 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6471 = tmp6474


}

ifres6468 = ifres6471


} else {
tmp6476 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6468 = tmp6476


}

let__5233 := ifres6468
_ = let__5233

tmp6479 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5233)


if True == tmp6479 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5233
}


panic("unreachable")
})

tmp6480 := Call2(__e, ns2_1set, symshen_4_5alpha_6, tmp6467)


_ = tmp6480

tmp6481 := MakeNative1(func(__e *ControlFlow, V2914 Obj) Obj {
tmp6488 := Call1(__e, SymFn(symshen_4lowercase_2), V2914)


if True == tmp6488 {
return True
} else {
tmp6486 := Call1(__e, SymFn(symshen_4uppercase_2), V2914)


var ifres6483 Obj

if True == tmp6486 {
ifres6483 = True


} else {
tmp6485 := Call1(__e, SymFn(symshen_4misc_2), V2914)


var ifres6484 Obj

if True == tmp6485 {
ifres6484 = True


} else {
ifres6484 = False


}

ifres6483 = ifres6484


}

if True == ifres6483 {
return True
} else {
return False
}


}


panic("unreachable")
})

tmp6489 := Call2(__e, ns2_1set, symshen_4alpha_2, tmp6481)


_ = tmp6489

tmp6490 := MakeNative1(func(__e *ControlFlow, V2915 Obj) Obj {
if True == GGe(__e, V2915, MakeInteger(97)) {
if True == GLe(__e, V2915, MakeInteger(122)) {
return True
} else {
return False
}


} else {
return False
}


panic("unreachable")
})

tmp6495 := Call2(__e, ns2_1set, symshen_4lowercase_2, tmp6490)


_ = tmp6495

tmp6496 := MakeNative1(func(__e *ControlFlow, V2916 Obj) Obj {
if True == GGe(__e, V2916, MakeInteger(65)) {
if True == GLe(__e, V2916, MakeInteger(90)) {
return True
} else {
return False
}


} else {
return False
}


panic("unreachable")
})

tmp6501 := Call2(__e, ns2_1set, symshen_4uppercase_2, tmp6496)


_ = tmp6501

tmp6502 := MakeNative1(func(__e *ControlFlow, V2917 Obj) Obj {
tmp6503 := GCons(__e, MakeInteger(96), Nil)

tmp6504 := GCons(__e, MakeInteger(35), tmp6503)

tmp6505 := GCons(__e, MakeInteger(39), tmp6504)

tmp6506 := GCons(__e, MakeInteger(37), tmp6505)

tmp6507 := GCons(__e, MakeInteger(38), tmp6506)

tmp6508 := GCons(__e, MakeInteger(60), tmp6507)

tmp6509 := GCons(__e, MakeInteger(62), tmp6508)

tmp6510 := GCons(__e, MakeInteger(46), tmp6509)

tmp6511 := GCons(__e, MakeInteger(126), tmp6510)

tmp6512 := GCons(__e, MakeInteger(64), tmp6511)

tmp6513 := GCons(__e, MakeInteger(33), tmp6512)

tmp6514 := GCons(__e, MakeInteger(36), tmp6513)

tmp6515 := GCons(__e, MakeInteger(63), tmp6514)

tmp6516 := GCons(__e, MakeInteger(95), tmp6515)

tmp6517 := GCons(__e, MakeInteger(43), tmp6516)

tmp6518 := GCons(__e, MakeInteger(47), tmp6517)

tmp6519 := GCons(__e, MakeInteger(42), tmp6518)

tmp6520 := GCons(__e, MakeInteger(45), tmp6519)

tmp6521 := GCons(__e, MakeInteger(61), tmp6520)

return TailCall2(__e, SymFn(symelement_2), V2917, tmp6521)


panic("unreachable")
})

tmp6522 := Call2(__e, ns2_1set, symshen_4misc_2, tmp6502)


_ = tmp6522

tmp6523 := MakeNative1(func(__e *ControlFlow, V2918 Obj) Obj {
tmp6524 := Call1(__e, SymFn(symshen_4_5alphanum_6), V2918)


let__5237 := tmp6524
_ = let__5237

tmp6537 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5237)


var ifres6525 Obj

if True == tmp6537 {
tmp6526 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6525 = tmp6526


} else {
tmp6527 := Call1(__e, SymFn(symshen_4_5_1out), let__5237)


let__5238 := tmp6527
_ = let__5238

tmp6528 := Call1(__e, SymFn(symshen_4in_1_6), let__5237)


let__5239 := tmp6528
_ = let__5239

tmp6529 := Call1(__e, SymFn(symshen_4_5alphanums_6), let__5239)


let__5240 := tmp6529
_ = let__5240

tmp6536 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5240)


var ifres6530 Obj

if True == tmp6536 {
tmp6531 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6530 = tmp6531


} else {
tmp6532 := Call1(__e, SymFn(symshen_4_5_1out), let__5240)


let__5241 := tmp6532
_ = let__5241

tmp6533 := Call1(__e, SymFn(symshen_4in_1_6), let__5240)


let__5242 := tmp6533
_ = let__5242

tmp6535 := Call2(__e, SymFn(symshen_4comb), let__5242, PrimGuard2(__e, symcn, PrimStringConcat, let__5238, let__5241))


ifres6530 = tmp6535


}

ifres6525 = ifres6530


}

let__5236 := ifres6525
_ = let__5236

tmp6547 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5236)


if True == tmp6547 {
tmp6538 := Call1(__e, SymFn(sym_5e_6), V2918)


let__5244 := tmp6538
_ = let__5244

tmp6543 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5244)


var ifres6539 Obj

if True == tmp6543 {
tmp6540 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6539 = tmp6540


} else {
tmp6541 := Call1(__e, SymFn(symshen_4in_1_6), let__5244)


let__5245 := tmp6541
_ = let__5245

tmp6542 := Call2(__e, SymFn(symshen_4comb), let__5245, MakeString(""))


ifres6539 = tmp6542


}

let__5243 := ifres6539
_ = let__5243

tmp6545 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5243)


if True == tmp6545 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5243
}


} else {
return let__5236
}


panic("unreachable")
})

tmp6548 := Call2(__e, ns2_1set, symshen_4_5alphanums_6, tmp6523)


_ = tmp6548

tmp6549 := MakeNative1(func(__e *ControlFlow, V2929 Obj) Obj {
tmp6550 := Call1(__e, SymFn(symshen_4_5alpha_6), V2929)


let__5247 := tmp6550
_ = let__5247

tmp6556 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5247)


var ifres6551 Obj

if True == tmp6556 {
tmp6552 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6551 = tmp6552


} else {
tmp6553 := Call1(__e, SymFn(symshen_4_5_1out), let__5247)


let__5248 := tmp6553
_ = let__5248

tmp6554 := Call1(__e, SymFn(symshen_4in_1_6), let__5247)


let__5249 := tmp6554
_ = let__5249

tmp6555 := Call2(__e, SymFn(symshen_4comb), let__5249, let__5248)


ifres6551 = tmp6555


}

let__5246 := ifres6551
_ = let__5246

tmp6567 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5246)


if True == tmp6567 {
tmp6557 := Call1(__e, SymFn(symshen_4_5numeral_6), V2929)


let__5251 := tmp6557
_ = let__5251

tmp6563 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5251)


var ifres6558 Obj

if True == tmp6563 {
tmp6559 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6558 = tmp6559


} else {
tmp6560 := Call1(__e, SymFn(symshen_4_5_1out), let__5251)


let__5252 := tmp6560
_ = let__5252

tmp6561 := Call1(__e, SymFn(symshen_4in_1_6), let__5251)


let__5253 := tmp6561
_ = let__5253

tmp6562 := Call2(__e, SymFn(symshen_4comb), let__5253, let__5252)


ifres6558 = tmp6562


}

let__5250 := ifres6558
_ = let__5250

tmp6565 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5250)


if True == tmp6565 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5250
}


} else {
return let__5246
}


panic("unreachable")
})

tmp6568 := Call2(__e, ns2_1set, symshen_4_5alphanum_6, tmp6549)


_ = tmp6568

tmp6569 := MakeNative1(func(__e *ControlFlow, V2938 Obj) Obj {
tmp6579 := GConsP(__e, V2938)

var ifres6570 Obj

if True == tmp6579 {
tmp6571 := Call1(__e, SymFn(symhead), V2938)


let__5255 := tmp6571
_ = let__5255

tmp6572 := Call1(__e, SymFn(symtail), V2938)


let__5256 := tmp6572
_ = let__5256

tmp6577 := Call1(__e, SymFn(symshen_4digit_2), let__5255)


var ifres6573 Obj

if True == tmp6577 {
tmp6574 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, let__5255)

tmp6575 := Call2(__e, SymFn(symshen_4comb), let__5256, tmp6574)


ifres6573 = tmp6575


} else {
tmp6576 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6573 = tmp6576


}

ifres6570 = ifres6573


} else {
tmp6578 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6570 = tmp6578


}

let__5254 := ifres6570
_ = let__5254

tmp6581 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5254)


if True == tmp6581 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5254
}


panic("unreachable")
})

tmp6582 := Call2(__e, ns2_1set, symshen_4_5numeral_6, tmp6569)


_ = tmp6582

tmp6583 := MakeNative1(func(__e *ControlFlow, V2942 Obj) Obj {
if True == GGe(__e, V2942, MakeInteger(48)) {
if True == GLe(__e, V2942, MakeInteger(57)) {
return True
} else {
return False
}


} else {
return False
}


panic("unreachable")
})

tmp6588 := Call2(__e, ns2_1set, symshen_4digit_2, tmp6583)


_ = tmp6588

tmp6589 := MakeNative1(func(__e *ControlFlow, V2943 Obj) Obj {
tmp6590 := Call1(__e, SymFn(symshen_4_5dbq_6), V2943)


let__5258 := tmp6590
_ = let__5258

tmp6606 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5258)


var ifres6591 Obj

if True == tmp6606 {
tmp6592 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6591 = tmp6592


} else {
tmp6593 := Call1(__e, SymFn(symshen_4in_1_6), let__5258)


let__5259 := tmp6593
_ = let__5259

tmp6594 := Call1(__e, SymFn(symshen_4_5strcontents_6), let__5259)


let__5260 := tmp6594
_ = let__5260

tmp6605 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5260)


var ifres6595 Obj

if True == tmp6605 {
tmp6596 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6595 = tmp6596


} else {
tmp6597 := Call1(__e, SymFn(symshen_4_5_1out), let__5260)


let__5261 := tmp6597
_ = let__5261

tmp6598 := Call1(__e, SymFn(symshen_4in_1_6), let__5260)


let__5262 := tmp6598
_ = let__5262

tmp6599 := Call1(__e, SymFn(symshen_4_5dbq_6), let__5262)


let__5263 := tmp6599
_ = let__5263

tmp6604 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5263)


var ifres6600 Obj

if True == tmp6604 {
tmp6601 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6600 = tmp6601


} else {
tmp6602 := Call1(__e, SymFn(symshen_4in_1_6), let__5263)


let__5264 := tmp6602
_ = let__5264

tmp6603 := Call2(__e, SymFn(symshen_4comb), let__5264, let__5261)


ifres6600 = tmp6603


}

ifres6595 = ifres6600


}

ifres6591 = ifres6595


}

let__5257 := ifres6591
_ = let__5257

tmp6608 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5257)


if True == tmp6608 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5257
}


panic("unreachable")
})

tmp6609 := Call2(__e, ns2_1set, symshen_4_5str_6, tmp6589)


_ = tmp6609

tmp6610 := MakeNative1(func(__e *ControlFlow, V2952 Obj) Obj {
tmp6615 := Call2(__e, SymFn(symshen_4hds_a_2), V2952, MakeInteger(34))


var ifres6611 Obj

if True == tmp6615 {
tmp6612 := Call1(__e, SymFn(symtail), V2952)


let__5266 := tmp6612
_ = let__5266

tmp6613 := Call2(__e, SymFn(symshen_4comb), let__5266, symshen_4skip)


ifres6611 = tmp6613


} else {
tmp6614 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6611 = tmp6614


}

let__5265 := ifres6611
_ = let__5265

tmp6617 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5265)


if True == tmp6617 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5265
}


panic("unreachable")
})

tmp6618 := Call2(__e, ns2_1set, symshen_4_5dbq_6, tmp6610)


_ = tmp6618

tmp6619 := MakeNative1(func(__e *ControlFlow, V2955 Obj) Obj {
tmp6620 := Call1(__e, SymFn(symshen_4_5strc_6), V2955)


let__5268 := tmp6620
_ = let__5268

tmp6633 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5268)


var ifres6621 Obj

if True == tmp6633 {
tmp6622 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6621 = tmp6622


} else {
tmp6623 := Call1(__e, SymFn(symshen_4_5_1out), let__5268)


let__5269 := tmp6623
_ = let__5269

tmp6624 := Call1(__e, SymFn(symshen_4in_1_6), let__5268)


let__5270 := tmp6624
_ = let__5270

tmp6625 := Call1(__e, SymFn(symshen_4_5strcontents_6), let__5270)


let__5271 := tmp6625
_ = let__5271

tmp6632 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5271)


var ifres6626 Obj

if True == tmp6632 {
tmp6627 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6626 = tmp6627


} else {
tmp6628 := Call1(__e, SymFn(symshen_4_5_1out), let__5271)


let__5272 := tmp6628
_ = let__5272

tmp6629 := Call1(__e, SymFn(symshen_4in_1_6), let__5271)


let__5273 := tmp6629
_ = let__5273

tmp6631 := Call2(__e, SymFn(symshen_4comb), let__5273, PrimGuard2(__e, symcn, PrimStringConcat, let__5269, let__5272))


ifres6626 = tmp6631


}

ifres6621 = ifres6626


}

let__5267 := ifres6621
_ = let__5267

tmp6643 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5267)


if True == tmp6643 {
tmp6634 := Call1(__e, SymFn(sym_5e_6), V2955)


let__5275 := tmp6634
_ = let__5275

tmp6639 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5275)


var ifres6635 Obj

if True == tmp6639 {
tmp6636 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6635 = tmp6636


} else {
tmp6637 := Call1(__e, SymFn(symshen_4in_1_6), let__5275)


let__5276 := tmp6637
_ = let__5276

tmp6638 := Call2(__e, SymFn(symshen_4comb), let__5276, MakeString(""))


ifres6635 = tmp6638


}

let__5274 := ifres6635
_ = let__5274

tmp6641 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5274)


if True == tmp6641 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5274
}


} else {
return let__5267
}


panic("unreachable")
})

tmp6644 := Call2(__e, ns2_1set, symshen_4_5strcontents_6, tmp6619)


_ = tmp6644

tmp6645 := MakeNative1(func(__e *ControlFlow, V2966 Obj) Obj {
tmp6646 := Call1(__e, SymFn(symshen_4_5control_6), V2966)


let__5278 := tmp6646
_ = let__5278

tmp6652 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5278)


var ifres6647 Obj

if True == tmp6652 {
tmp6648 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6647 = tmp6648


} else {
tmp6649 := Call1(__e, SymFn(symshen_4_5_1out), let__5278)


let__5279 := tmp6649
_ = let__5279

tmp6650 := Call1(__e, SymFn(symshen_4in_1_6), let__5278)


let__5280 := tmp6650
_ = let__5280

tmp6651 := Call2(__e, SymFn(symshen_4comb), let__5280, let__5279)


ifres6647 = tmp6651


}

let__5277 := ifres6647
_ = let__5277

tmp6663 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5277)


if True == tmp6663 {
tmp6653 := Call1(__e, SymFn(symshen_4_5notdbq_6), V2966)


let__5282 := tmp6653
_ = let__5282

tmp6659 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5282)


var ifres6654 Obj

if True == tmp6659 {
tmp6655 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6654 = tmp6655


} else {
tmp6656 := Call1(__e, SymFn(symshen_4_5_1out), let__5282)


let__5283 := tmp6656
_ = let__5283

tmp6657 := Call1(__e, SymFn(symshen_4in_1_6), let__5282)


let__5284 := tmp6657
_ = let__5284

tmp6658 := Call2(__e, SymFn(symshen_4comb), let__5284, let__5283)


ifres6654 = tmp6658


}

let__5281 := ifres6654
_ = let__5281

tmp6661 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5281)


if True == tmp6661 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5281
}


} else {
return let__5277
}


panic("unreachable")
})

tmp6664 := Call2(__e, ns2_1set, symshen_4_5strc_6, tmp6645)


_ = tmp6664

tmp6665 := MakeNative1(func(__e *ControlFlow, V2975 Obj) Obj {
tmp6666 := Call1(__e, SymFn(symshen_4_5lowC_6), V2975)


let__5286 := tmp6666
_ = let__5286

tmp6688 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5286)


var ifres6667 Obj

if True == tmp6688 {
tmp6668 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6667 = tmp6668


} else {
tmp6669 := Call1(__e, SymFn(symshen_4in_1_6), let__5286)


let__5287 := tmp6669
_ = let__5287

tmp6670 := Call1(__e, SymFn(symshen_4_5hash_6), let__5287)


let__5288 := tmp6670
_ = let__5288

tmp6687 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5288)


var ifres6671 Obj

if True == tmp6687 {
tmp6672 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6671 = tmp6672


} else {
tmp6673 := Call1(__e, SymFn(symshen_4in_1_6), let__5288)


let__5289 := tmp6673
_ = let__5289

tmp6674 := Call1(__e, SymFn(symshen_4_5integer_6), let__5289)


let__5290 := tmp6674
_ = let__5290

tmp6686 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5290)


var ifres6675 Obj

if True == tmp6686 {
tmp6676 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6675 = tmp6676


} else {
tmp6677 := Call1(__e, SymFn(symshen_4_5_1out), let__5290)


let__5291 := tmp6677
_ = let__5291

tmp6678 := Call1(__e, SymFn(symshen_4in_1_6), let__5290)


let__5292 := tmp6678
_ = let__5292

tmp6679 := Call1(__e, SymFn(symshen_4_5semicolon_6), let__5292)


let__5293 := tmp6679
_ = let__5293

tmp6685 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5293)


var ifres6680 Obj

if True == tmp6685 {
tmp6681 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6680 = tmp6681


} else {
tmp6682 := Call1(__e, SymFn(symshen_4in_1_6), let__5293)


let__5294 := tmp6682
_ = let__5294

tmp6683 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, let__5291)

tmp6684 := Call2(__e, SymFn(symshen_4comb), let__5294, tmp6683)


ifres6680 = tmp6684


}

ifres6675 = ifres6680


}

ifres6671 = ifres6675


}

ifres6667 = ifres6671


}

let__5285 := ifres6667
_ = let__5285

tmp6690 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5285)


if True == tmp6690 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5285
}


panic("unreachable")
})

tmp6691 := Call2(__e, ns2_1set, symshen_4_5control_6, tmp6665)


_ = tmp6691

tmp6692 := MakeNative1(func(__e *ControlFlow, V2986 Obj) Obj {
tmp6703 := GConsP(__e, V2986)

var ifres6693 Obj

if True == tmp6703 {
tmp6694 := Call1(__e, SymFn(symhead), V2986)


let__5296 := tmp6694
_ = let__5296

tmp6695 := Call1(__e, SymFn(symtail), V2986)


let__5297 := tmp6695
_ = let__5297

tmp6700 := GEq(__e, let__5296, MakeInteger(34))

tmp6701 := GNot(__e, tmp6700)

var ifres6696 Obj

if True == tmp6701 {
tmp6697 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, let__5296)

tmp6698 := Call2(__e, SymFn(symshen_4comb), let__5297, tmp6697)


ifres6696 = tmp6698


} else {
tmp6699 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6696 = tmp6699


}

ifres6693 = ifres6696


} else {
tmp6702 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6693 = tmp6702


}

let__5295 := ifres6693
_ = let__5295

tmp6705 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5295)


if True == tmp6705 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5295
}


panic("unreachable")
})

tmp6706 := Call2(__e, ns2_1set, symshen_4_5notdbq_6, tmp6692)


_ = tmp6706

tmp6707 := MakeNative1(func(__e *ControlFlow, V2990 Obj) Obj {
tmp6712 := Call2(__e, SymFn(symshen_4hds_a_2), V2990, MakeInteger(99))


var ifres6708 Obj

if True == tmp6712 {
tmp6709 := Call1(__e, SymFn(symtail), V2990)


let__5299 := tmp6709
_ = let__5299

tmp6710 := Call2(__e, SymFn(symshen_4comb), let__5299, symshen_4skip)


ifres6708 = tmp6710


} else {
tmp6711 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6708 = tmp6711


}

let__5298 := ifres6708
_ = let__5298

tmp6714 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5298)


if True == tmp6714 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5298
}


panic("unreachable")
})

tmp6715 := Call2(__e, ns2_1set, symshen_4_5lowC_6, tmp6707)


_ = tmp6715

tmp6716 := MakeNative1(func(__e *ControlFlow, V2993 Obj) Obj {
tmp6721 := Call2(__e, SymFn(symshen_4hds_a_2), V2993, MakeInteger(35))


var ifres6717 Obj

if True == tmp6721 {
tmp6718 := Call1(__e, SymFn(symtail), V2993)


let__5301 := tmp6718
_ = let__5301

tmp6719 := Call2(__e, SymFn(symshen_4comb), let__5301, symshen_4skip)


ifres6717 = tmp6719


} else {
tmp6720 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6717 = tmp6720


}

let__5300 := ifres6717
_ = let__5300

tmp6723 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5300)


if True == tmp6723 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5300
}


panic("unreachable")
})

tmp6724 := Call2(__e, ns2_1set, symshen_4_5hash_6, tmp6716)


_ = tmp6724

tmp6725 := MakeNative1(func(__e *ControlFlow, V2996 Obj) Obj {
tmp6726 := Call1(__e, SymFn(symshen_4_5minus_6), V2996)


let__5303 := tmp6726
_ = let__5303

tmp6738 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5303)


var ifres6727 Obj

if True == tmp6738 {
tmp6728 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6727 = tmp6728


} else {
tmp6729 := Call1(__e, SymFn(symshen_4in_1_6), let__5303)


let__5304 := tmp6729
_ = let__5304

tmp6730 := Call1(__e, SymFn(symshen_4_5number_6), let__5304)


let__5305 := tmp6730
_ = let__5305

tmp6737 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5305)


var ifres6731 Obj

if True == tmp6737 {
tmp6732 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6731 = tmp6732


} else {
tmp6733 := Call1(__e, SymFn(symshen_4_5_1out), let__5305)


let__5306 := tmp6733
_ = let__5306

tmp6734 := Call1(__e, SymFn(symshen_4in_1_6), let__5305)


let__5307 := tmp6734
_ = let__5307

tmp6736 := Call2(__e, SymFn(symshen_4comb), let__5307, GSub(__e, MakeInteger(0), let__5306))


ifres6731 = tmp6736


}

ifres6727 = ifres6731


}

let__5302 := ifres6727
_ = let__5302

tmp6781 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5302)


if True == tmp6781 {
tmp6739 := Call1(__e, SymFn(symshen_4_5plus_6), V2996)


let__5309 := tmp6739
_ = let__5309

tmp6750 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5309)


var ifres6740 Obj

if True == tmp6750 {
tmp6741 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6740 = tmp6741


} else {
tmp6742 := Call1(__e, SymFn(symshen_4in_1_6), let__5309)


let__5310 := tmp6742
_ = let__5310

tmp6743 := Call1(__e, SymFn(symshen_4_5number_6), let__5310)


let__5311 := tmp6743
_ = let__5311

tmp6749 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5311)


var ifres6744 Obj

if True == tmp6749 {
tmp6745 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6744 = tmp6745


} else {
tmp6746 := Call1(__e, SymFn(symshen_4_5_1out), let__5311)


let__5312 := tmp6746
_ = let__5312

tmp6747 := Call1(__e, SymFn(symshen_4in_1_6), let__5311)


let__5313 := tmp6747
_ = let__5313

tmp6748 := Call2(__e, SymFn(symshen_4comb), let__5313, let__5312)


ifres6744 = tmp6748


}

ifres6740 = ifres6744


}

let__5308 := ifres6740
_ = let__5308

tmp6779 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5308)


if True == tmp6779 {
tmp6751 := Call1(__e, SymFn(symshen_4_5e_1number_6), V2996)


let__5315 := tmp6751
_ = let__5315

tmp6757 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5315)


var ifres6752 Obj

if True == tmp6757 {
tmp6753 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6752 = tmp6753


} else {
tmp6754 := Call1(__e, SymFn(symshen_4_5_1out), let__5315)


let__5316 := tmp6754
_ = let__5316

tmp6755 := Call1(__e, SymFn(symshen_4in_1_6), let__5315)


let__5317 := tmp6755
_ = let__5317

tmp6756 := Call2(__e, SymFn(symshen_4comb), let__5317, let__5316)


ifres6752 = tmp6756


}

let__5314 := ifres6752
_ = let__5314

tmp6777 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5314)


if True == tmp6777 {
tmp6758 := Call1(__e, SymFn(symshen_4_5float_6), V2996)


let__5319 := tmp6758
_ = let__5319

tmp6764 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5319)


var ifres6759 Obj

if True == tmp6764 {
tmp6760 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6759 = tmp6760


} else {
tmp6761 := Call1(__e, SymFn(symshen_4_5_1out), let__5319)


let__5320 := tmp6761
_ = let__5320

tmp6762 := Call1(__e, SymFn(symshen_4in_1_6), let__5319)


let__5321 := tmp6762
_ = let__5321

tmp6763 := Call2(__e, SymFn(symshen_4comb), let__5321, let__5320)


ifres6759 = tmp6763


}

let__5318 := ifres6759
_ = let__5318

tmp6775 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5318)


if True == tmp6775 {
tmp6765 := Call1(__e, SymFn(symshen_4_5integer_6), V2996)


let__5323 := tmp6765
_ = let__5323

tmp6771 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5323)


var ifres6766 Obj

if True == tmp6771 {
tmp6767 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6766 = tmp6767


} else {
tmp6768 := Call1(__e, SymFn(symshen_4_5_1out), let__5323)


let__5324 := tmp6768
_ = let__5324

tmp6769 := Call1(__e, SymFn(symshen_4in_1_6), let__5323)


let__5325 := tmp6769
_ = let__5325

tmp6770 := Call2(__e, SymFn(symshen_4comb), let__5325, let__5324)


ifres6766 = tmp6770


}

let__5322 := ifres6766
_ = let__5322

tmp6773 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5322)


if True == tmp6773 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5322
}


} else {
return let__5318
}


} else {
return let__5314
}


} else {
return let__5308
}


} else {
return let__5302
}


panic("unreachable")
})

tmp6782 := Call2(__e, ns2_1set, symshen_4_5number_6, tmp6725)


_ = tmp6782

tmp6783 := MakeNative1(func(__e *ControlFlow, V3021 Obj) Obj {
tmp6788 := Call2(__e, SymFn(symshen_4hds_a_2), V3021, MakeInteger(45))


var ifres6784 Obj

if True == tmp6788 {
tmp6785 := Call1(__e, SymFn(symtail), V3021)


let__5327 := tmp6785
_ = let__5327

tmp6786 := Call2(__e, SymFn(symshen_4comb), let__5327, symshen_4skip)


ifres6784 = tmp6786


} else {
tmp6787 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6784 = tmp6787


}

let__5326 := ifres6784
_ = let__5326

tmp6790 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5326)


if True == tmp6790 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5326
}


panic("unreachable")
})

tmp6791 := Call2(__e, ns2_1set, symshen_4_5minus_6, tmp6783)


_ = tmp6791

tmp6792 := MakeNative1(func(__e *ControlFlow, V3024 Obj) Obj {
tmp6797 := Call2(__e, SymFn(symshen_4hds_a_2), V3024, MakeInteger(43))


var ifres6793 Obj

if True == tmp6797 {
tmp6794 := Call1(__e, SymFn(symtail), V3024)


let__5329 := tmp6794
_ = let__5329

tmp6795 := Call2(__e, SymFn(symshen_4comb), let__5329, symshen_4skip)


ifres6793 = tmp6795


} else {
tmp6796 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6793 = tmp6796


}

let__5328 := ifres6793
_ = let__5328

tmp6799 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5328)


if True == tmp6799 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5328
}


panic("unreachable")
})

tmp6800 := Call2(__e, ns2_1set, symshen_4_5plus_6, tmp6792)


_ = tmp6800

tmp6801 := MakeNative1(func(__e *ControlFlow, V3027 Obj) Obj {
tmp6802 := Call1(__e, SymFn(symshen_4_5digits_6), V3027)


let__5331 := tmp6802
_ = let__5331

tmp6809 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5331)


var ifres6803 Obj

if True == tmp6809 {
tmp6804 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6803 = tmp6804


} else {
tmp6805 := Call1(__e, SymFn(symshen_4_5_1out), let__5331)


let__5332 := tmp6805
_ = let__5332

tmp6806 := Call1(__e, SymFn(symshen_4in_1_6), let__5331)


let__5333 := tmp6806
_ = let__5333

tmp6807 := Call1(__e, SymFn(symshen_4compute_1integer), let__5332)


tmp6808 := Call2(__e, SymFn(symshen_4comb), let__5333, tmp6807)


ifres6803 = tmp6808


}

let__5330 := ifres6803
_ = let__5330

tmp6811 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5330)


if True == tmp6811 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5330
}


panic("unreachable")
})

tmp6812 := Call2(__e, ns2_1set, symshen_4_5integer_6, tmp6801)


_ = tmp6812

tmp6813 := MakeNative1(func(__e *ControlFlow, V3032 Obj) Obj {
tmp6814 := Call1(__e, SymFn(symshen_4_5digit_6), V3032)


let__5335 := tmp6814
_ = let__5335

tmp6827 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5335)


var ifres6815 Obj

if True == tmp6827 {
tmp6816 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6815 = tmp6816


} else {
tmp6817 := Call1(__e, SymFn(symshen_4_5_1out), let__5335)


let__5336 := tmp6817
_ = let__5336

tmp6818 := Call1(__e, SymFn(symshen_4in_1_6), let__5335)


let__5337 := tmp6818
_ = let__5337

tmp6819 := Call1(__e, SymFn(symshen_4_5digits_6), let__5337)


let__5338 := tmp6819
_ = let__5338

tmp6826 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5338)


var ifres6820 Obj

if True == tmp6826 {
tmp6821 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6820 = tmp6821


} else {
tmp6822 := Call1(__e, SymFn(symshen_4_5_1out), let__5338)


let__5339 := tmp6822
_ = let__5339

tmp6823 := Call1(__e, SymFn(symshen_4in_1_6), let__5338)


let__5340 := tmp6823
_ = let__5340

tmp6824 := GCons(__e, let__5336, let__5339)

tmp6825 := Call2(__e, SymFn(symshen_4comb), let__5340, tmp6824)


ifres6820 = tmp6825


}

ifres6815 = ifres6820


}

let__5334 := ifres6815
_ = let__5334

tmp6839 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5334)


if True == tmp6839 {
tmp6828 := Call1(__e, SymFn(symshen_4_5digit_6), V3032)


let__5342 := tmp6828
_ = let__5342

tmp6835 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5342)


var ifres6829 Obj

if True == tmp6835 {
tmp6830 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6829 = tmp6830


} else {
tmp6831 := Call1(__e, SymFn(symshen_4_5_1out), let__5342)


let__5343 := tmp6831
_ = let__5343

tmp6832 := Call1(__e, SymFn(symshen_4in_1_6), let__5342)


let__5344 := tmp6832
_ = let__5344

tmp6833 := GCons(__e, let__5343, Nil)

tmp6834 := Call2(__e, SymFn(symshen_4comb), let__5344, tmp6833)


ifres6829 = tmp6834


}

let__5341 := ifres6829
_ = let__5341

tmp6837 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5341)


if True == tmp6837 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5341
}


} else {
return let__5334
}


panic("unreachable")
})

tmp6840 := Call2(__e, ns2_1set, symshen_4_5digits_6, tmp6813)


_ = tmp6840

tmp6841 := MakeNative1(func(__e *ControlFlow, V3044 Obj) Obj {
tmp6851 := GConsP(__e, V3044)

var ifres6842 Obj

if True == tmp6851 {
tmp6843 := Call1(__e, SymFn(symhead), V3044)


let__5346 := tmp6843
_ = let__5346

tmp6844 := Call1(__e, SymFn(symtail), V3044)


let__5347 := tmp6844
_ = let__5347

tmp6849 := Call1(__e, SymFn(symshen_4digit_2), let__5346)


var ifres6845 Obj

if True == tmp6849 {
tmp6846 := Call1(__e, SymFn(symshen_4byte_1_6digit), let__5346)


tmp6847 := Call2(__e, SymFn(symshen_4comb), let__5347, tmp6846)


ifres6845 = tmp6847


} else {
tmp6848 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6845 = tmp6848


}

ifres6842 = ifres6845


} else {
tmp6850 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6842 = tmp6850


}

let__5345 := ifres6842
_ = let__5345

tmp6853 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5345)


if True == tmp6853 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5345
}


panic("unreachable")
})

tmp6854 := Call2(__e, ns2_1set, symshen_4_5digit_6, tmp6841)


_ = tmp6854

tmp6855 := MakeNative1(func(__e *ControlFlow, V3048 Obj) Obj {
return GSub(__e, V3048, MakeInteger(48))
panic("unreachable")
})

tmp6856 := Call2(__e, ns2_1set, symshen_4byte_1_6digit, tmp6855)


_ = tmp6856

tmp6857 := MakeNative1(func(__e *ControlFlow, V3049 Obj) Obj {
tmp6858 := Call1(__e, SymFn(symreverse), V3049)


return TailCall2(__e, SymFn(symshen_4compute_1integer_1h), tmp6858, MakeInteger(0))


panic("unreachable")
})

tmp6859 := Call2(__e, ns2_1set, symshen_4compute_1integer, tmp6857)


_ = tmp6859

tmp6860 := MakeNative2(func(__e *ControlFlow, V3052, V3053 Obj) Obj {
tmp6870 := GEq(__e, Nil, V3052)

if True == tmp6870 {
return MakeInteger(0)
} else {
tmp6868 := GConsP(__e, V3052)

if True == tmp6868 {
tmp6861 := Call2(__e, SymFn(symshen_4expt), MakeInteger(10), V3053)


tmp6862 := GHd(__e, V3052)

tmp6863 := GMul(__e, tmp6861, tmp6862)

tmp6864 := GTl(__e, V3052)

tmp6866 := Call2(__e, SymFn(symshen_4compute_1integer_1h), tmp6864, GAdd(__e, V3053, MakeInteger(1)))


return GAdd(__e, tmp6863, tmp6866)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4compute_1integer_1h)
}


}


panic("unreachable")
})

tmp6871 := Call2(__e, ns2_1set, symshen_4compute_1integer_1h, tmp6860)


_ = tmp6871

tmp6872 := MakeNative2(func(__e *ControlFlow, V3056, V3057 Obj) Obj {
tmp6880 := GEq(__e, MakeInteger(0), V3057)

if True == tmp6880 {
return MakeInteger(1)
} else {
if True == GGt(__e, V3057, MakeInteger(0)) {
tmp6874 := Call2(__e, SymFn(symshen_4expt), V3056, GSub(__e, V3057, MakeInteger(1)))


return GMul(__e, V3056, tmp6874)


} else {
tmp6876 := Call2(__e, SymFn(symshen_4expt), V3056, GAdd(__e, V3057, MakeInteger(1)))


return GDiv(__e, tmp6876, V3056)


}


}


panic("unreachable")
})

tmp6881 := Call2(__e, ns2_1set, symshen_4expt, tmp6872)


_ = tmp6881

tmp6882 := MakeNative1(func(__e *ControlFlow, V3058 Obj) Obj {
tmp6883 := Call1(__e, SymFn(symshen_4_5integer_6), V3058)


let__5349 := tmp6883
_ = let__5349

tmp6901 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5349)


var ifres6884 Obj

if True == tmp6901 {
tmp6885 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6884 = tmp6885


} else {
tmp6886 := Call1(__e, SymFn(symshen_4_5_1out), let__5349)


let__5350 := tmp6886
_ = let__5350

tmp6887 := Call1(__e, SymFn(symshen_4in_1_6), let__5349)


let__5351 := tmp6887
_ = let__5351

tmp6888 := Call1(__e, SymFn(symshen_4_5stop_6), let__5351)


let__5352 := tmp6888
_ = let__5352

tmp6900 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5352)


var ifres6889 Obj

if True == tmp6900 {
tmp6890 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6889 = tmp6890


} else {
tmp6891 := Call1(__e, SymFn(symshen_4in_1_6), let__5352)


let__5353 := tmp6891
_ = let__5353

tmp6892 := Call1(__e, SymFn(symshen_4_5fraction_6), let__5353)


let__5354 := tmp6892
_ = let__5354

tmp6899 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5354)


var ifres6893 Obj

if True == tmp6899 {
tmp6894 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6893 = tmp6894


} else {
tmp6895 := Call1(__e, SymFn(symshen_4_5_1out), let__5354)


let__5355 := tmp6895
_ = let__5355

tmp6896 := Call1(__e, SymFn(symshen_4in_1_6), let__5354)


let__5356 := tmp6896
_ = let__5356

tmp6898 := Call2(__e, SymFn(symshen_4comb), let__5356, GAdd(__e, let__5350, let__5355))


ifres6893 = tmp6898


}

ifres6889 = ifres6893


}

ifres6884 = ifres6889


}

let__5348 := ifres6884
_ = let__5348

tmp6917 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5348)


if True == tmp6917 {
tmp6902 := Call1(__e, SymFn(symshen_4_5stop_6), V3058)


let__5358 := tmp6902
_ = let__5358

tmp6913 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5358)


var ifres6903 Obj

if True == tmp6913 {
tmp6904 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6903 = tmp6904


} else {
tmp6905 := Call1(__e, SymFn(symshen_4in_1_6), let__5358)


let__5359 := tmp6905
_ = let__5359

tmp6906 := Call1(__e, SymFn(symshen_4_5fraction_6), let__5359)


let__5360 := tmp6906
_ = let__5360

tmp6912 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5360)


var ifres6907 Obj

if True == tmp6912 {
tmp6908 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6907 = tmp6908


} else {
tmp6909 := Call1(__e, SymFn(symshen_4_5_1out), let__5360)


let__5361 := tmp6909
_ = let__5361

tmp6910 := Call1(__e, SymFn(symshen_4in_1_6), let__5360)


let__5362 := tmp6910
_ = let__5362

tmp6911 := Call2(__e, SymFn(symshen_4comb), let__5362, let__5361)


ifres6907 = tmp6911


}

ifres6903 = ifres6907


}

let__5357 := ifres6903
_ = let__5357

tmp6915 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5357)


if True == tmp6915 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5357
}


} else {
return let__5348
}


panic("unreachable")
})

tmp6918 := Call2(__e, ns2_1set, symshen_4_5float_6, tmp6882)


_ = tmp6918

tmp6919 := MakeNative1(func(__e *ControlFlow, V3074 Obj) Obj {
tmp6924 := Call2(__e, SymFn(symshen_4hds_a_2), V3074, MakeInteger(46))


var ifres6920 Obj

if True == tmp6924 {
tmp6921 := Call1(__e, SymFn(symtail), V3074)


let__5364 := tmp6921
_ = let__5364

tmp6922 := Call2(__e, SymFn(symshen_4comb), let__5364, symshen_4skip)


ifres6920 = tmp6922


} else {
tmp6923 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6920 = tmp6923


}

let__5363 := ifres6920
_ = let__5363

tmp6926 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5363)


if True == tmp6926 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5363
}


panic("unreachable")
})

tmp6927 := Call2(__e, ns2_1set, symshen_4_5stop_6, tmp6919)


_ = tmp6927

tmp6928 := MakeNative1(func(__e *ControlFlow, V3077 Obj) Obj {
tmp6929 := Call1(__e, SymFn(symshen_4_5digits_6), V3077)


let__5366 := tmp6929
_ = let__5366

tmp6936 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5366)


var ifres6930 Obj

if True == tmp6936 {
tmp6931 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6930 = tmp6931


} else {
tmp6932 := Call1(__e, SymFn(symshen_4_5_1out), let__5366)


let__5367 := tmp6932
_ = let__5367

tmp6933 := Call1(__e, SymFn(symshen_4in_1_6), let__5366)


let__5368 := tmp6933
_ = let__5368

tmp6934 := Call1(__e, SymFn(symshen_4compute_1fraction), let__5367)


tmp6935 := Call2(__e, SymFn(symshen_4comb), let__5368, tmp6934)


ifres6930 = tmp6935


}

let__5365 := ifres6930
_ = let__5365

tmp6938 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5365)


if True == tmp6938 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5365
}


panic("unreachable")
})

tmp6939 := Call2(__e, ns2_1set, symshen_4_5fraction_6, tmp6928)


_ = tmp6939

tmp6940 := MakeNative1(func(__e *ControlFlow, V3082 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4compute_1fraction_1h), V3082, MakeInteger(-1))
panic("unreachable")
})

tmp6941 := Call2(__e, ns2_1set, symshen_4compute_1fraction, tmp6940)


_ = tmp6941

tmp6942 := MakeNative2(func(__e *ControlFlow, V3085, V3086 Obj) Obj {
tmp6952 := GEq(__e, Nil, V3085)

if True == tmp6952 {
return MakeInteger(0)
} else {
tmp6950 := GConsP(__e, V3085)

if True == tmp6950 {
tmp6943 := Call2(__e, SymFn(symshen_4expt), MakeInteger(10), V3086)


tmp6944 := GHd(__e, V3085)

tmp6945 := GMul(__e, tmp6943, tmp6944)

tmp6946 := GTl(__e, V3085)

tmp6948 := Call2(__e, SymFn(symshen_4compute_1fraction_1h), tmp6946, GSub(__e, V3086, MakeInteger(1)))


return GAdd(__e, tmp6945, tmp6948)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4compute_1fraction_1h)
}


}


panic("unreachable")
})

tmp6953 := Call2(__e, ns2_1set, symshen_4compute_1fraction_1h, tmp6942)


_ = tmp6953

tmp6954 := MakeNative1(func(__e *ControlFlow, V3087 Obj) Obj {
tmp6955 := Call1(__e, SymFn(symshen_4_5float_6), V3087)


let__5370 := tmp6955
_ = let__5370

tmp6973 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5370)


var ifres6956 Obj

if True == tmp6973 {
tmp6957 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6956 = tmp6957


} else {
tmp6958 := Call1(__e, SymFn(symshen_4_5_1out), let__5370)


let__5371 := tmp6958
_ = let__5371

tmp6959 := Call1(__e, SymFn(symshen_4in_1_6), let__5370)


let__5372 := tmp6959
_ = let__5372

tmp6960 := Call1(__e, SymFn(symshen_4_5lowE_6), let__5372)


let__5373 := tmp6960
_ = let__5373

tmp6972 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5373)


var ifres6961 Obj

if True == tmp6972 {
tmp6962 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6961 = tmp6962


} else {
tmp6963 := Call1(__e, SymFn(symshen_4in_1_6), let__5373)


let__5374 := tmp6963
_ = let__5374

tmp6964 := Call1(__e, SymFn(symshen_4_5log10_6), let__5374)


let__5375 := tmp6964
_ = let__5375

tmp6971 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5375)


var ifres6965 Obj

if True == tmp6971 {
tmp6966 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6965 = tmp6966


} else {
tmp6967 := Call1(__e, SymFn(symshen_4_5_1out), let__5375)


let__5376 := tmp6967
_ = let__5376

tmp6968 := Call1(__e, SymFn(symshen_4in_1_6), let__5375)


let__5377 := tmp6968
_ = let__5377

tmp6969 := Call2(__e, SymFn(symshen_4compute_1E), let__5371, let__5376)


tmp6970 := Call2(__e, SymFn(symshen_4comb), let__5377, tmp6969)


ifres6965 = tmp6970


}

ifres6961 = ifres6965


}

ifres6956 = ifres6961


}

let__5369 := ifres6956
_ = let__5369

tmp6996 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5369)


if True == tmp6996 {
tmp6974 := Call1(__e, SymFn(symshen_4_5integer_6), V3087)


let__5379 := tmp6974
_ = let__5379

tmp6992 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5379)


var ifres6975 Obj

if True == tmp6992 {
tmp6976 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6975 = tmp6976


} else {
tmp6977 := Call1(__e, SymFn(symshen_4_5_1out), let__5379)


let__5380 := tmp6977
_ = let__5380

tmp6978 := Call1(__e, SymFn(symshen_4in_1_6), let__5379)


let__5381 := tmp6978
_ = let__5381

tmp6979 := Call1(__e, SymFn(symshen_4_5lowE_6), let__5381)


let__5382 := tmp6979
_ = let__5382

tmp6991 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5382)


var ifres6980 Obj

if True == tmp6991 {
tmp6981 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6980 = tmp6981


} else {
tmp6982 := Call1(__e, SymFn(symshen_4in_1_6), let__5382)


let__5383 := tmp6982
_ = let__5383

tmp6983 := Call1(__e, SymFn(symshen_4_5log10_6), let__5383)


let__5384 := tmp6983
_ = let__5384

tmp6990 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5384)


var ifres6984 Obj

if True == tmp6990 {
tmp6985 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres6984 = tmp6985


} else {
tmp6986 := Call1(__e, SymFn(symshen_4_5_1out), let__5384)


let__5385 := tmp6986
_ = let__5385

tmp6987 := Call1(__e, SymFn(symshen_4in_1_6), let__5384)


let__5386 := tmp6987
_ = let__5386

tmp6988 := Call2(__e, SymFn(symshen_4compute_1E), let__5380, let__5385)


tmp6989 := Call2(__e, SymFn(symshen_4comb), let__5386, tmp6988)


ifres6984 = tmp6989


}

ifres6980 = ifres6984


}

ifres6975 = ifres6980


}

let__5378 := ifres6975
_ = let__5378

tmp6994 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5378)


if True == tmp6994 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5378
}


} else {
return let__5369
}


panic("unreachable")
})

tmp6997 := Call2(__e, ns2_1set, symshen_4_5e_1number_6, tmp6954)


_ = tmp6997

tmp6998 := MakeNative1(func(__e *ControlFlow, V3106 Obj) Obj {
tmp6999 := Call1(__e, SymFn(symshen_4_5plus_6), V3106)


let__5388 := tmp6999
_ = let__5388

tmp7010 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5388)


var ifres7000 Obj

if True == tmp7010 {
tmp7001 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres7000 = tmp7001


} else {
tmp7002 := Call1(__e, SymFn(symshen_4in_1_6), let__5388)


let__5389 := tmp7002
_ = let__5389

tmp7003 := Call1(__e, SymFn(symshen_4_5log10_6), let__5389)


let__5390 := tmp7003
_ = let__5390

tmp7009 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5390)


var ifres7004 Obj

if True == tmp7009 {
tmp7005 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres7004 = tmp7005


} else {
tmp7006 := Call1(__e, SymFn(symshen_4_5_1out), let__5390)


let__5391 := tmp7006
_ = let__5391

tmp7007 := Call1(__e, SymFn(symshen_4in_1_6), let__5390)


let__5392 := tmp7007
_ = let__5392

tmp7008 := Call2(__e, SymFn(symshen_4comb), let__5392, let__5391)


ifres7004 = tmp7008


}

ifres7000 = ifres7004


}

let__5387 := ifres7000
_ = let__5387

tmp7036 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5387)


if True == tmp7036 {
tmp7011 := Call1(__e, SymFn(symshen_4_5minus_6), V3106)


let__5394 := tmp7011
_ = let__5394

tmp7023 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5394)


var ifres7012 Obj

if True == tmp7023 {
tmp7013 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres7012 = tmp7013


} else {
tmp7014 := Call1(__e, SymFn(symshen_4in_1_6), let__5394)


let__5395 := tmp7014
_ = let__5395

tmp7015 := Call1(__e, SymFn(symshen_4_5log10_6), let__5395)


let__5396 := tmp7015
_ = let__5396

tmp7022 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5396)


var ifres7016 Obj

if True == tmp7022 {
tmp7017 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres7016 = tmp7017


} else {
tmp7018 := Call1(__e, SymFn(symshen_4_5_1out), let__5396)


let__5397 := tmp7018
_ = let__5397

tmp7019 := Call1(__e, SymFn(symshen_4in_1_6), let__5396)


let__5398 := tmp7019
_ = let__5398

tmp7021 := Call2(__e, SymFn(symshen_4comb), let__5398, GSub(__e, MakeInteger(0), let__5397))


ifres7016 = tmp7021


}

ifres7012 = ifres7016


}

let__5393 := ifres7012
_ = let__5393

tmp7034 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5393)


if True == tmp7034 {
tmp7024 := Call1(__e, SymFn(symshen_4_5integer_6), V3106)


let__5400 := tmp7024
_ = let__5400

tmp7030 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5400)


var ifres7025 Obj

if True == tmp7030 {
tmp7026 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres7025 = tmp7026


} else {
tmp7027 := Call1(__e, SymFn(symshen_4_5_1out), let__5400)


let__5401 := tmp7027
_ = let__5401

tmp7028 := Call1(__e, SymFn(symshen_4in_1_6), let__5400)


let__5402 := tmp7028
_ = let__5402

tmp7029 := Call2(__e, SymFn(symshen_4comb), let__5402, let__5401)


ifres7025 = tmp7029


}

let__5399 := ifres7025
_ = let__5399

tmp7032 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5399)


if True == tmp7032 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5399
}


} else {
return let__5393
}


} else {
return let__5387
}


panic("unreachable")
})

tmp7037 := Call2(__e, ns2_1set, symshen_4_5log10_6, tmp6998)


_ = tmp7037

tmp7038 := MakeNative1(func(__e *ControlFlow, V3123 Obj) Obj {
tmp7043 := Call2(__e, SymFn(symshen_4hds_a_2), V3123, MakeInteger(101))


var ifres7039 Obj

if True == tmp7043 {
tmp7040 := Call1(__e, SymFn(symtail), V3123)


let__5404 := tmp7040
_ = let__5404

tmp7041 := Call2(__e, SymFn(symshen_4comb), let__5404, symshen_4skip)


ifres7039 = tmp7041


} else {
tmp7042 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres7039 = tmp7042


}

let__5403 := ifres7039
_ = let__5403

tmp7045 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5403)


if True == tmp7045 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5403
}


panic("unreachable")
})

tmp7046 := Call2(__e, ns2_1set, symshen_4_5lowE_6, tmp7038)


_ = tmp7046

tmp7047 := MakeNative2(func(__e *ControlFlow, V3126, V3127 Obj) Obj {
tmp7048 := Call2(__e, SymFn(symshen_4expt), MakeInteger(10), V3127)


return GMul(__e, V3126, tmp7048)


panic("unreachable")
})

tmp7049 := Call2(__e, ns2_1set, symshen_4compute_1E, tmp7047)


_ = tmp7049

tmp7050 := MakeNative1(func(__e *ControlFlow, V3128 Obj) Obj {
tmp7051 := Call1(__e, SymFn(symshen_4_5whitespace_6), V3128)


let__5406 := tmp7051
_ = let__5406

tmp7061 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5406)


var ifres7052 Obj

if True == tmp7061 {
tmp7053 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres7052 = tmp7053


} else {
tmp7054 := Call1(__e, SymFn(symshen_4in_1_6), let__5406)


let__5407 := tmp7054
_ = let__5407

tmp7055 := Call1(__e, SymFn(symshen_4_5whitespaces_6), let__5407)


let__5408 := tmp7055
_ = let__5408

tmp7060 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5408)


var ifres7056 Obj

if True == tmp7060 {
tmp7057 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres7056 = tmp7057


} else {
tmp7058 := Call1(__e, SymFn(symshen_4in_1_6), let__5408)


let__5409 := tmp7058
_ = let__5409

tmp7059 := Call2(__e, SymFn(symshen_4comb), let__5409, symshen_4skip)


ifres7056 = tmp7059


}

ifres7052 = ifres7056


}

let__5405 := ifres7052
_ = let__5405

tmp7071 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5405)


if True == tmp7071 {
tmp7062 := Call1(__e, SymFn(symshen_4_5whitespace_6), V3128)


let__5411 := tmp7062
_ = let__5411

tmp7067 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5411)


var ifres7063 Obj

if True == tmp7067 {
tmp7064 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres7063 = tmp7064


} else {
tmp7065 := Call1(__e, SymFn(symshen_4in_1_6), let__5411)


let__5412 := tmp7065
_ = let__5412

tmp7066 := Call2(__e, SymFn(symshen_4comb), let__5412, symshen_4skip)


ifres7063 = tmp7066


}

let__5410 := ifres7063
_ = let__5410

tmp7069 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5410)


if True == tmp7069 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5410
}


} else {
return let__5405
}


panic("unreachable")
})

tmp7072 := Call2(__e, ns2_1set, symshen_4_5whitespaces_6, tmp7050)


_ = tmp7072

tmp7073 := MakeNative1(func(__e *ControlFlow, V3137 Obj) Obj {
tmp7082 := GConsP(__e, V3137)

var ifres7074 Obj

if True == tmp7082 {
tmp7075 := Call1(__e, SymFn(symhead), V3137)


let__5414 := tmp7075
_ = let__5414

tmp7076 := Call1(__e, SymFn(symtail), V3137)


let__5415 := tmp7076
_ = let__5415

tmp7080 := Call1(__e, SymFn(symshen_4whitespace_2), let__5414)


var ifres7077 Obj

if True == tmp7080 {
tmp7078 := Call2(__e, SymFn(symshen_4comb), let__5415, symshen_4skip)


ifres7077 = tmp7078


} else {
tmp7079 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres7077 = tmp7079


}

ifres7074 = ifres7077


} else {
tmp7081 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres7074 = tmp7081


}

let__5413 := ifres7074
_ = let__5413

tmp7084 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__5413)


if True == tmp7084 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__5413
}


panic("unreachable")
})

tmp7085 := Call2(__e, ns2_1set, symshen_4_5whitespace_6, tmp7073)


_ = tmp7085

tmp7086 := MakeNative1(func(__e *ControlFlow, V3143 Obj) Obj {
tmp7094 := GEq(__e, MakeInteger(32), V3143)

if True == tmp7094 {
return True
} else {
tmp7092 := GEq(__e, MakeInteger(13), V3143)

if True == tmp7092 {
return True
} else {
tmp7090 := GEq(__e, MakeInteger(10), V3143)

if True == tmp7090 {
return True
} else {
tmp7088 := GEq(__e, MakeInteger(9), V3143)

if True == tmp7088 {
return True
} else {
return False
}


}


}


}


panic("unreachable")
})

tmp7095 := Call2(__e, ns2_1set, symshen_4whitespace_2, tmp7086)


_ = tmp7095

tmp7096 := MakeNative2(func(__e *ControlFlow, V3146, V3147 Obj) Obj {
tmp7111 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V3146)

var ifres7103 Obj

if True == tmp7111 {
tmp7109 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V3146)

tmp7110 := GNot(__e, tmp7109)

var ifres7105 Obj

if True == tmp7110 {
tmp7107 := GEq(__e, V3146, sympackage)

tmp7108 := GNot(__e, tmp7107)

var ifres7106 Obj

if True == tmp7108 {
ifres7106 = True


} else {
ifres7106 = False


}

ifres7105 = ifres7106


} else {
ifres7105 = False


}

var ifres7104 Obj

if True == ifres7105 {
ifres7104 = True


} else {
ifres7104 = False


}

ifres7103 = ifres7104


} else {
ifres7103 = False


}

if True == ifres7103 {
tmp7097 := GValue(__e, symshen_4_dsource_1form_1handlers_d)

tmp7098 := Call3(__e, SymFn(symshen_4update_1assoc), V3146, V3147, tmp7097)


tmp7099 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsource_1form_1handlers_d, tmp7098)

_ = tmp7099

return V3146


} else {
tmp7100 := Call3(__e, SymFn(symshen_4app), V3146, MakeString("\n"), symshen_4s)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("invalid source-form head: "), tmp7100))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("invalid source-form head: "), tmp7100)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


panic("unreachable")
})

tmp7112 := Call2(__e, ns2_1set, symshen_4register_1source_1form, tmp7096)


_ = tmp7112

tmp7113 := MakeNative1(func(__e *ControlFlow, V3148 Obj) Obj {
tmp7114 := GValue(__e, symshen_4_dsource_1form_1handlers_d)

tmp7115 := Call2(__e, SymFn(symassoc), V3148, tmp7114)


tmp7116 := GValue(__e, symshen_4_dsource_1form_1handlers_d)

tmp7117 := Call2(__e, SymFn(symremove), tmp7115, tmp7116)


tmp7118 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsource_1form_1handlers_d, tmp7117)

_ = tmp7118

return V3148


panic("unreachable")
})

tmp7119 := Call2(__e, ns2_1set, symshen_4unregister_1source_1form, tmp7113)


_ = tmp7119

tmp7120 := MakeNative2(func(__e *ControlFlow, V3155, V3156 Obj) Obj {
tmp7123 := GConsP(__e, V3155)

if True == tmp7123 {
tmp7121 := GHd(__e, V3155)

return TailCall2(__e, SymFn(symassoc), tmp7121, V3156)


} else {
return Nil
}


panic("unreachable")
})

tmp7124 := Call2(__e, ns2_1set, symshen_4source_1form_1handler, tmp7120)


_ = tmp7124

tmp7125 := MakeNative2(func(__e *ControlFlow, V3159, V3160 Obj) Obj {
tmp7133 := GConsP(__e, V3160)

if True == tmp7133 {
tmp7126 := GTl(__e, V3160)

tmp7127 := Call1(__e, tmp7126, V3159)


let__5416 := tmp7127
_ = let__5416

tmp7131 := Call1(__e, SymFn(symshen_4source_1form_1list_2), let__5416)


if True == tmp7131 {
return let__5416
} else {
tmp7128 := Call3(__e, SymFn(symshen_4app), V3159, MakeString("\n"), symshen_4s)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("source-form handler must return a list of forms: "), tmp7128))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("source-form handler must return a list of forms: "), tmp7128)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4source_1form_1expansion)
}


panic("unreachable")
})

tmp7134 := Call2(__e, ns2_1set, symshen_4source_1form_1expansion, tmp7125)


_ = tmp7134

tmp7135 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V3166 := __self1
_ = V3166
tmp7140 := GEq(__e, Nil, V3166)

if True == tmp7140 {
return True
} else {
tmp7138 := GConsP(__e, V3166)

if True == tmp7138 {
tmp7136 := GTl(__e, V3166)

if SymFn(symshen_4source_1form_1list_2) == __self {
__self1 = tmp7136
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4source_1form_1list_2), tmp7136)


} else {
return False
}


}


panic("unreachable")
})
return __self
})()

tmp7141 := Call2(__e, ns2_1set, symshen_4source_1form_1list_2, tmp7135)


_ = tmp7141

tmp7142 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V3167 := __self1
_ = V3167
tmp7178 := GEq(__e, Nil, V3167)

if True == tmp7178 {
return Nil
} else {
tmp7176 := GConsP(__e, V3167)

var ifres7172 Obj

if True == tmp7176 {
tmp7174 := GHd(__e, V3167)

tmp7175 := Call1(__e, SymFn(symshen_4packaged_2), tmp7174)


var ifres7173 Obj

if True == tmp7175 {
ifres7173 = True


} else {
ifres7173 = False


}

ifres7172 = ifres7173


} else {
ifres7172 = False


}

if True == ifres7172 {
tmp7143 := GHd(__e, V3167)

tmp7144 := Call1(__e, SymFn(symshen_4unpackage), tmp7143)


tmp7145 := GTl(__e, V3167)

tmp7146 := Call2(__e, SymFn(symappend), tmp7144, tmp7145)


if SymFn(symshen_4unpackage_emacroexpand) == __self {
__self1 = tmp7146
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4unpackage_emacroexpand), tmp7146)


} else {
tmp7170 := GConsP(__e, V3167)

if True == tmp7170 {
tmp7147 := GHd(__e, V3167)

tmp7148 := GValue(__e, symshen_4_dsource_1form_1handlers_d)

tmp7149 := Call2(__e, SymFn(symshen_4source_1form_1handler), tmp7147, tmp7148)


let__5417 := tmp7149
_ = let__5417

tmp7168 := GConsP(__e, let__5417)

if True == tmp7168 {
tmp7150 := GHd(__e, V3167)

tmp7151 := Call2(__e, SymFn(symshen_4source_1form_1expansion), tmp7150, let__5417)


tmp7152 := GTl(__e, V3167)

tmp7153 := Call2(__e, SymFn(symappend), tmp7151, tmp7152)


if SymFn(symshen_4unpackage_emacroexpand) == __self {
__self1 = tmp7153
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4unpackage_emacroexpand), tmp7153)


} else {
tmp7154 := GHd(__e, V3167)

tmp7155 := Call1(__e, SymFn(symmacroexpand), tmp7154)


let__5418 := tmp7155
_ = let__5418

tmp7166 := Call1(__e, SymFn(symshen_4packaged_2), let__5418)


var ifres7161 Obj

if True == tmp7166 {
ifres7161 = True


} else {
tmp7163 := GValue(__e, symshen_4_dsource_1form_1handlers_d)

tmp7164 := Call2(__e, SymFn(symshen_4source_1form_1handler), let__5418, tmp7163)


tmp7165 := GConsP(__e, tmp7164)

var ifres7162 Obj

if True == tmp7165 {
ifres7162 = True


} else {
ifres7162 = False


}

ifres7161 = ifres7162


}

if True == ifres7161 {
tmp7156 := GTl(__e, V3167)

tmp7157 := GCons(__e, let__5418, tmp7156)

if SymFn(symshen_4unpackage_emacroexpand) == __self {
__self1 = tmp7157
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4unpackage_emacroexpand), tmp7157)


} else {
tmp7158 := GTl(__e, V3167)

tmp7159 := Call1(__e, SymFn(symshen_4unpackage_emacroexpand), tmp7158)


return GCons(__e, let__5418, tmp7159)


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

tmp7179 := Call2(__e, ns2_1set, symshen_4unpackage_emacroexpand, tmp7142)


_ = tmp7179

tmp7180 := MakeNative1(func(__e *ControlFlow, V3172 Obj) Obj {
tmp7195 := GConsP(__e, V3172)

var ifres7182 Obj

if True == tmp7195 {
tmp7193 := GHd(__e, V3172)

tmp7194 := GEq(__e, sympackage, tmp7193)

var ifres7184 Obj

if True == tmp7194 {
tmp7191 := GTl(__e, V3172)

tmp7192 := GConsP(__e, tmp7191)

var ifres7186 Obj

if True == tmp7192 {
tmp7188 := GTl(__e, V3172)

tmp7189 := GTl(__e, tmp7188)

tmp7190 := GConsP(__e, tmp7189)

var ifres7187 Obj

if True == tmp7190 {
ifres7187 = True


} else {
ifres7187 = False


}

ifres7186 = ifres7187


} else {
ifres7186 = False


}

var ifres7185 Obj

if True == ifres7186 {
ifres7185 = True


} else {
ifres7185 = False


}

ifres7184 = ifres7185


} else {
ifres7184 = False


}

var ifres7183 Obj

if True == ifres7184 {
ifres7183 = True


} else {
ifres7183 = False


}

ifres7182 = ifres7183


} else {
ifres7182 = False


}

if True == ifres7182 {
return True
} else {
return False
}


panic("unreachable")
})

tmp7196 := Call2(__e, ns2_1set, symshen_4packaged_2, tmp7180)


_ = tmp7196

tmp7197 := MakeNative1(func(__e *ControlFlow, V3175 Obj) Obj {
tmp7254 := GConsP(__e, V3175)

var ifres7236 Obj

if True == tmp7254 {
tmp7252 := GHd(__e, V3175)

tmp7253 := GEq(__e, sympackage, tmp7252)

var ifres7238 Obj

if True == tmp7253 {
tmp7250 := GTl(__e, V3175)

tmp7251 := GConsP(__e, tmp7250)

var ifres7240 Obj

if True == tmp7251 {
tmp7247 := GTl(__e, V3175)

tmp7248 := GHd(__e, tmp7247)

tmp7249 := GEq(__e, symnull, tmp7248)

var ifres7242 Obj

if True == tmp7249 {
tmp7244 := GTl(__e, V3175)

tmp7245 := GTl(__e, tmp7244)

tmp7246 := GConsP(__e, tmp7245)

var ifres7243 Obj

if True == tmp7246 {
ifres7243 = True


} else {
ifres7243 = False


}

ifres7242 = ifres7243


} else {
ifres7242 = False


}

var ifres7241 Obj

if True == ifres7242 {
ifres7241 = True


} else {
ifres7241 = False


}

ifres7240 = ifres7241


} else {
ifres7240 = False


}

var ifres7239 Obj

if True == ifres7240 {
ifres7239 = True


} else {
ifres7239 = False


}

ifres7238 = ifres7239


} else {
ifres7238 = False


}

var ifres7237 Obj

if True == ifres7238 {
ifres7237 = True


} else {
ifres7237 = False


}

ifres7236 = ifres7237


} else {
ifres7236 = False


}

if True == ifres7236 {
tmp7198 := GTl(__e, V3175)

tmp7199 := GTl(__e, tmp7198)

return GTl(__e, tmp7199)


} else {
tmp7234 := GConsP(__e, V3175)

var ifres7221 Obj

if True == tmp7234 {
tmp7232 := GHd(__e, V3175)

tmp7233 := GEq(__e, sympackage, tmp7232)

var ifres7223 Obj

if True == tmp7233 {
tmp7230 := GTl(__e, V3175)

tmp7231 := GConsP(__e, tmp7230)

var ifres7225 Obj

if True == tmp7231 {
tmp7227 := GTl(__e, V3175)

tmp7228 := GTl(__e, tmp7227)

tmp7229 := GConsP(__e, tmp7228)

var ifres7226 Obj

if True == tmp7229 {
ifres7226 = True


} else {
ifres7226 = False


}

ifres7225 = ifres7226


} else {
ifres7225 = False


}

var ifres7224 Obj

if True == ifres7225 {
ifres7224 = True


} else {
ifres7224 = False


}

ifres7223 = ifres7224


} else {
ifres7223 = False


}

var ifres7222 Obj

if True == ifres7223 {
ifres7222 = True


} else {
ifres7222 = False


}

ifres7221 = ifres7222


} else {
ifres7221 = False


}

if True == ifres7221 {
tmp7200 := GTl(__e, V3175)

tmp7201 := GTl(__e, tmp7200)

tmp7202 := GHd(__e, tmp7201)

tmp7203 := Call1(__e, SymFn(symeval), tmp7202)


let__5419 := tmp7203
_ = let__5419

tmp7204 := GTl(__e, V3175)

tmp7205 := GHd(__e, tmp7204)

tmp7206 := PrimGuard1(__e, symstr, PrimStr, tmp7205)

tmp7207 := GTl(__e, V3175)

tmp7208 := GTl(__e, tmp7207)

tmp7209 := GTl(__e, tmp7208)

tmp7210 := Call3(__e, SymFn(symshen_4package_1symbols), tmp7206, let__5419, tmp7209)


let__5420 := tmp7210
_ = let__5420

tmp7211 := GTl(__e, V3175)

tmp7212 := GHd(__e, tmp7211)

tmp7213 := Call2(__e, SymFn(symshen_4record_1external), tmp7212, let__5419)


let__5421 := tmp7213
_ = let__5421

tmp7214 := GTl(__e, V3175)

tmp7215 := GHd(__e, tmp7214)

tmp7216 := GTl(__e, V3175)

tmp7217 := GTl(__e, tmp7216)

tmp7218 := GTl(__e, tmp7217)

tmp7219 := Call3(__e, SymFn(symshen_4record_1internal), tmp7215, let__5419, tmp7218)


let__5422 := tmp7219
_ = let__5422

return let__5420


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4unpackage)
}


}


panic("unreachable")
})

tmp7255 := Call2(__e, ns2_1set, symshen_4unpackage, tmp7197)


_ = tmp7255

tmp7256 := MakeNative3(func(__e *ControlFlow, V3180, V3181, V3182 Obj) Obj {
tmp7257 := MakeNative0(func(__e *ControlFlow) Obj {
tmp7258 := GValue(__e, sym_dproperty_1vector_d)

return TailCall3(__e, SymFn(symget), V3180, symshen_4internal_1symbols, tmp7258)


panic("unreachable")
})

tmp7259 := MakeNative1(func(__e *ControlFlow, Z3184 Obj) Obj {
return Nil
panic("unreachable")
})

tmp7260 := Call2(__e, try_1catch, tmp7257, tmp7259)


let__5423 := tmp7260
_ = let__5423

tmp7261 := PrimGuard1(__e, symstr, PrimStr, V3180)

tmp7262 := Call3(__e, SymFn(symshen_4internal_1symbols), tmp7261, V3181, V3182)


let__5424 := tmp7262
_ = let__5424

tmp7263 := Call2(__e, SymFn(symunion), let__5424, let__5423)


tmp7264 := GValue(__e, sym_dproperty_1vector_d)

return TailCall4(__e, SymFn(symput), V3180, symshen_4internal_1symbols, tmp7263, tmp7264)


panic("unreachable")
})

tmp7265 := Call2(__e, ns2_1set, symshen_4record_1internal, tmp7256)


_ = tmp7265

tmp7266 := MakeNative3(func(__e *ControlFlow, V3192, V3193, V3194 Obj) Obj {
tmp7275 := GConsP(__e, V3194)

if True == tmp7275 {
tmp7267 := GHd(__e, V3194)

tmp7268 := Call3(__e, SymFn(symshen_4internal_1symbols), V3192, V3193, tmp7267)


tmp7269 := GTl(__e, V3194)

tmp7270 := Call3(__e, SymFn(symshen_4internal_1symbols), V3192, V3193, tmp7269)


return TailCall2(__e, SymFn(symunion), tmp7268, tmp7270)


} else {
tmp7273 := Call3(__e, SymFn(symshen_4internal_2), V3194, V3192, V3193)


if True == tmp7273 {
tmp7271 := Call2(__e, SymFn(symshen_4intern_1in_1package), V3192, V3194)


return GCons(__e, tmp7271, Nil)


} else {
return Nil
}


}


panic("unreachable")
})

tmp7276 := Call2(__e, ns2_1set, symshen_4internal_1symbols, tmp7266)


_ = tmp7276

tmp7277 := MakeNative2(func(__e *ControlFlow, V3195, V3196 Obj) Obj {
tmp7278 := MakeNative0(func(__e *ControlFlow) Obj {
tmp7279 := GValue(__e, sym_dproperty_1vector_d)

return TailCall3(__e, SymFn(symget), V3195, symshen_4external_1symbols, tmp7279)


panic("unreachable")
})

tmp7280 := MakeNative1(func(__e *ControlFlow, Z3198 Obj) Obj {
return Nil
panic("unreachable")
})

tmp7281 := Call2(__e, try_1catch, tmp7278, tmp7280)


let__5425 := tmp7281
_ = let__5425

tmp7282 := Call2(__e, SymFn(symunion), V3196, let__5425)


tmp7283 := GValue(__e, sym_dproperty_1vector_d)

return TailCall4(__e, SymFn(symput), V3195, symshen_4external_1symbols, tmp7282, tmp7283)


panic("unreachable")
})

tmp7284 := Call2(__e, ns2_1set, symshen_4record_1external, tmp7277)


_ = tmp7284

tmp7285 := MakeNative3(func(__e *ControlFlow, V3203, V3204, V3205 Obj) Obj {
tmp7290 := GConsP(__e, V3205)

if True == tmp7290 {
tmp7286 := MakeNative1(func(__e *ControlFlow, Z3206 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4package_1symbols), V3203, V3204, Z3206)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp7286, V3205)


} else {
tmp7288 := Call3(__e, SymFn(symshen_4internal_2), V3205, V3203, V3204)


if True == tmp7288 {
return TailCall2(__e, SymFn(symshen_4intern_1in_1package), V3203, V3205)
} else {
return V3205
}


}


panic("unreachable")
})

tmp7291 := Call2(__e, ns2_1set, symshen_4package_1symbols, tmp7285)


_ = tmp7291

tmp7292 := MakeNative2(func(__e *ControlFlow, V3207, V3208 Obj) Obj {
tmp7293 := PrimGuard1(__e, symstr, PrimStr, V3208)

tmp7294 := Call2(__e, SymFn(sym_8s), MakeString("."), tmp7293)


tmp7295 := Call2(__e, SymFn(sym_8s), V3207, tmp7294)


return PrimGuard1(__e, symintern, PrimIntern, tmp7295)


panic("unreachable")
})

tmp7296 := Call2(__e, ns2_1set, symshen_4intern_1in_1package, tmp7292)


_ = tmp7296

tmp7297 := MakeNative3(func(__e *ControlFlow, V3209, V3210, V3211 Obj) Obj {
tmp7327 := Call2(__e, SymFn(symelement_2), V3209, V3211)


if True == GNot(__e, tmp7327) {
tmp7324 := Call1(__e, SymFn(symshen_4sng_2), V3209)


tmp7325 := GNot(__e, tmp7324)

var ifres7299 Obj

if True == tmp7325 {
tmp7322 := Call1(__e, SymFn(symshen_4dbl_2), V3209)


tmp7323 := GNot(__e, tmp7322)

var ifres7301 Obj

if True == tmp7323 {
tmp7321 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V3209)

var ifres7303 Obj

if True == tmp7321 {
tmp7319 := Call1(__e, SymFn(symshen_4sysfunc_2), V3209)


tmp7320 := GNot(__e, tmp7319)

var ifres7305 Obj

if True == tmp7320 {
tmp7317 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V3209)

tmp7318 := GNot(__e, tmp7317)

var ifres7307 Obj

if True == tmp7318 {
tmp7314 := PrimGuard1(__e, symstr, PrimStr, V3209)

tmp7315 := Call1(__e, SymFn(symshen_4internal_1to_1shen_2), tmp7314)


tmp7316 := GNot(__e, tmp7315)

var ifres7309 Obj

if True == tmp7316 {
tmp7311 := PrimGuard1(__e, symstr, PrimStr, V3209)

tmp7312 := Call2(__e, SymFn(symshen_4internal_1to_1P_2), V3210, tmp7311)


tmp7313 := GNot(__e, tmp7312)

var ifres7310 Obj

if True == tmp7313 {
ifres7310 = True


} else {
ifres7310 = False


}

ifres7309 = ifres7310


} else {
ifres7309 = False


}

var ifres7308 Obj

if True == ifres7309 {
ifres7308 = True


} else {
ifres7308 = False


}

ifres7307 = ifres7308


} else {
ifres7307 = False


}

var ifres7306 Obj

if True == ifres7307 {
ifres7306 = True


} else {
ifres7306 = False


}

ifres7305 = ifres7306


} else {
ifres7305 = False


}

var ifres7304 Obj

if True == ifres7305 {
ifres7304 = True


} else {
ifres7304 = False


}

ifres7303 = ifres7304


} else {
ifres7303 = False


}

var ifres7302 Obj

if True == ifres7303 {
ifres7302 = True


} else {
ifres7302 = False


}

ifres7301 = ifres7302


} else {
ifres7301 = False


}

var ifres7300 Obj

if True == ifres7301 {
ifres7300 = True


} else {
ifres7300 = False


}

ifres7299 = ifres7300


} else {
ifres7299 = False


}

if True == ifres7299 {
return True
} else {
return False
}


} else {
return False
}


panic("unreachable")
})

tmp7329 := Call2(__e, ns2_1set, symshen_4internal_2, tmp7297)


_ = tmp7329

tmp7330 := MakeNative1(func(__e *ControlFlow, V3216 Obj) Obj {
tmp7384 := Call1(__e, SymFn(symshen_4_7string_2), V3216)


var ifres7332 Obj

if True == tmp7384 {
tmp7382 := Call1(__e, SymFn(symhdstr), V3216)


tmp7383 := GEq(__e, MakeString("s"), tmp7382)

var ifres7334 Obj

if True == tmp7383 {
tmp7381 := Call1(__e, SymFn(symshen_4_7string_2), PrimGuard1(__e, symtlstr, PrimTailString, V3216))


var ifres7336 Obj

if True == tmp7381 {
tmp7378 := Call1(__e, SymFn(symhdstr), PrimGuard1(__e, symtlstr, PrimTailString, V3216))


tmp7379 := GEq(__e, MakeString("h"), tmp7378)

var ifres7338 Obj

if True == tmp7379 {
tmp7376 := Call1(__e, SymFn(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V3216)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})())


var ifres7340 Obj

if True == tmp7376 {
tmp7372 := Call1(__e, SymFn(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := PrimGuard1(__e, symtlstr, PrimTailString, V3216)
return Call1(__e, SymFn(symtlstr), __typedArg0)
})())


tmp7373 := GEq(__e, MakeString("e"), tmp7372)

var ifres7342 Obj

if True == tmp7373 {
tmp7369 := Call1(__e, SymFn(symshen_4_7string_2), (func() Obj {
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


var ifres7344 Obj

if True == tmp7369 {
tmp7364 := Call1(__e, SymFn(symhdstr), (func() Obj {
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


tmp7365 := GEq(__e, MakeString("n"), tmp7364)

var ifres7346 Obj

if True == tmp7365 {
tmp7360 := Call1(__e, SymFn(symshen_4_7string_2), (func() Obj {
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


var ifres7348 Obj

if True == tmp7360 {
tmp7354 := Call1(__e, SymFn(symhdstr), (func() Obj {
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


tmp7355 := GEq(__e, MakeString("."), tmp7354)

var ifres7349 Obj

if True == tmp7355 {
ifres7349 = True


} else {
ifres7349 = False


}

ifres7348 = ifres7349


} else {
ifres7348 = False


}

var ifres7347 Obj

if True == ifres7348 {
ifres7347 = True


} else {
ifres7347 = False


}

ifres7346 = ifres7347


} else {
ifres7346 = False


}

var ifres7345 Obj

if True == ifres7346 {
ifres7345 = True


} else {
ifres7345 = False


}

ifres7344 = ifres7345


} else {
ifres7344 = False


}

var ifres7343 Obj

if True == ifres7344 {
ifres7343 = True


} else {
ifres7343 = False


}

ifres7342 = ifres7343


} else {
ifres7342 = False


}

var ifres7341 Obj

if True == ifres7342 {
ifres7341 = True


} else {
ifres7341 = False


}

ifres7340 = ifres7341


} else {
ifres7340 = False


}

var ifres7339 Obj

if True == ifres7340 {
ifres7339 = True


} else {
ifres7339 = False


}

ifres7338 = ifres7339


} else {
ifres7338 = False


}

var ifres7337 Obj

if True == ifres7338 {
ifres7337 = True


} else {
ifres7337 = False


}

ifres7336 = ifres7337


} else {
ifres7336 = False


}

var ifres7335 Obj

if True == ifres7336 {
ifres7335 = True


} else {
ifres7335 = False


}

ifres7334 = ifres7335


} else {
ifres7334 = False


}

var ifres7333 Obj

if True == ifres7334 {
ifres7333 = True


} else {
ifres7333 = False


}

ifres7332 = ifres7333


} else {
ifres7332 = False


}

if True == ifres7332 {
return True
} else {
return False
}


panic("unreachable")
})

tmp7385 := Call2(__e, ns2_1set, symshen_4internal_1to_1shen_2, tmp7330)


_ = tmp7385

tmp7386 := MakeNative1(func(__e *ControlFlow, V3217 Obj) Obj {
tmp7387 := GValue(__e, sym_dproperty_1vector_d)

tmp7388 := Call3(__e, SymFn(symget), symshen, symshen_4external_1symbols, tmp7387)


return TailCall2(__e, SymFn(symelement_2), V3217, tmp7388)


panic("unreachable")
})

tmp7389 := Call2(__e, ns2_1set, symshen_4sysfunc_2, tmp7386)


_ = tmp7389

tmp7390 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V3225 := __self1
_ = V3225
V3226 := __self2
_ = V3226
tmp7411 := GEq(__e, MakeString(""), V3225)

var ifres7404 Obj

if True == tmp7411 {
tmp7410 := Call1(__e, SymFn(symshen_4_7string_2), V3226)


var ifres7406 Obj

if True == tmp7410 {
tmp7408 := Call1(__e, SymFn(symhdstr), V3226)


tmp7409 := GEq(__e, MakeString("."), tmp7408)

var ifres7407 Obj

if True == tmp7409 {
ifres7407 = True


} else {
ifres7407 = False


}

ifres7406 = ifres7407


} else {
ifres7406 = False


}

var ifres7405 Obj

if True == ifres7406 {
ifres7405 = True


} else {
ifres7405 = False


}

ifres7404 = ifres7405


} else {
ifres7404 = False


}

if True == ifres7404 {
return True
} else {
tmp7402 := Call1(__e, SymFn(symshen_4_7string_2), V3225)


var ifres7394 Obj

if True == tmp7402 {
tmp7401 := Call1(__e, SymFn(symshen_4_7string_2), V3226)


var ifres7396 Obj

if True == tmp7401 {
tmp7398 := Call1(__e, SymFn(symhdstr), V3225)


tmp7399 := Call1(__e, SymFn(symhdstr), V3226)


tmp7400 := GEq(__e, tmp7398, tmp7399)

var ifres7397 Obj

if True == tmp7400 {
ifres7397 = True


} else {
ifres7397 = False


}

ifres7396 = ifres7397


} else {
ifres7396 = False


}

var ifres7395 Obj

if True == ifres7396 {
ifres7395 = True


} else {
ifres7395 = False


}

ifres7394 = ifres7395


} else {
ifres7394 = False


}

if True == ifres7394 {
tmp7391 := PrimGuard1(__e, symtlstr, PrimTailString, V3225)

if SymFn(symshen_4internal_1to_1P_2) == __self {
__self1, __self2 = tmp7391, PrimGuard1(__e, symtlstr, PrimTailString, V3226)
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4internal_1to_1P_2), tmp7391, PrimGuard1(__e, symtlstr, PrimTailString, V3226))


} else {
return False
}


}


panic("unreachable")
})
return __self
})()

tmp7412 := Call2(__e, ns2_1set, symshen_4internal_1to_1P_2, tmp7390)


_ = tmp7412

tmp7413 := MakeNative2(func(__e *ControlFlow, V3229, V3230 Obj) Obj {
tmp7434 := Call2(__e, SymFn(symelement_2), V3229, V3230)


if True == tmp7434 {
return V3229
} else {
tmp7432 := GConsP(__e, V3229)

var ifres7428 Obj

if True == tmp7432 {
tmp7430 := GHd(__e, V3229)

tmp7431 := GEq(__e, symcond, tmp7430)

var ifres7429 Obj

if True == tmp7431 {
ifres7429 = True


} else {
ifres7429 = False


}

ifres7428 = ifres7429


} else {
ifres7428 = False


}

if True == ifres7428 {
tmp7414 := GTl(__e, V3229)

tmp7415 := Call2(__e, SymFn(symshen_4process_1cond_1clauses), tmp7414, V3230)


return GCons(__e, symcond, tmp7415)


} else {
tmp7426 := GConsP(__e, V3229)

var ifres7422 Obj

if True == tmp7426 {
tmp7424 := GHd(__e, V3229)

tmp7425 := Call1(__e, SymFn(symshen_4non_1application_2), tmp7424)


var ifres7423 Obj

if True == tmp7425 {
ifres7423 = True


} else {
ifres7423 = False


}

ifres7422 = ifres7423


} else {
ifres7422 = False


}

if True == ifres7422 {
tmp7416 := GHd(__e, V3229)

return TailCall3(__e, SymFn(symshen_4special_1case), tmp7416, V3229, V3230)


} else {
tmp7420 := GConsP(__e, V3229)

if True == tmp7420 {
tmp7417 := MakeNative1(func(__e *ControlFlow, Z3231 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4process_1applications), Z3231, V3230)
panic("unreachable")
})

tmp7418 := Call2(__e, SymFn(symmap), tmp7417, V3229)


return TailCall2(__e, SymFn(symshen_4process_1application), tmp7418, V3230)


} else {
return V3229
}


}


}


}


panic("unreachable")
})

tmp7435 := Call2(__e, ns2_1set, symshen_4process_1applications, tmp7413)


_ = tmp7435

tmp7436 := MakeNative1(func(__e *ControlFlow, V3234 Obj) Obj {
tmp7446 := GEq(__e, symdefine, V3234)

if True == tmp7446 {
return True
} else {
tmp7444 := GEq(__e, symdefun, V3234)

if True == tmp7444 {
return True
} else {
tmp7442 := GEq(__e, symsynonyms, V3234)

if True == tmp7442 {
return True
} else {
tmp7440 := Call1(__e, SymFn(symshen_4special_2), V3234)


if True == tmp7440 {
return True
} else {
tmp7438 := Call1(__e, SymFn(symshen_4extraspecial_2), V3234)


if True == tmp7438 {
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

tmp7447 := Call2(__e, ns2_1set, symshen_4non_1application_2, tmp7436)


_ = tmp7447

tmp7448 := MakeNative3(func(__e *ControlFlow, V3239, V3240, V3241 Obj) Obj {
tmp7690 := GEq(__e, symlambda, V3239)

var ifres7668 Obj

if True == tmp7690 {
tmp7689 := GConsP(__e, V3240)

var ifres7670 Obj

if True == tmp7689 {
tmp7687 := GHd(__e, V3240)

tmp7688 := GEq(__e, symlambda, tmp7687)

var ifres7672 Obj

if True == tmp7688 {
tmp7685 := GTl(__e, V3240)

tmp7686 := GConsP(__e, tmp7685)

var ifres7674 Obj

if True == tmp7686 {
tmp7682 := GTl(__e, V3240)

tmp7683 := GTl(__e, tmp7682)

tmp7684 := GConsP(__e, tmp7683)

var ifres7676 Obj

if True == tmp7684 {
tmp7678 := GTl(__e, V3240)

tmp7679 := GTl(__e, tmp7678)

tmp7680 := GTl(__e, tmp7679)

tmp7681 := GEq(__e, Nil, tmp7680)

var ifres7677 Obj

if True == tmp7681 {
ifres7677 = True


} else {
ifres7677 = False


}

ifres7676 = ifres7677


} else {
ifres7676 = False


}

var ifres7675 Obj

if True == ifres7676 {
ifres7675 = True


} else {
ifres7675 = False


}

ifres7674 = ifres7675


} else {
ifres7674 = False


}

var ifres7673 Obj

if True == ifres7674 {
ifres7673 = True


} else {
ifres7673 = False


}

ifres7672 = ifres7673


} else {
ifres7672 = False


}

var ifres7671 Obj

if True == ifres7672 {
ifres7671 = True


} else {
ifres7671 = False


}

ifres7670 = ifres7671


} else {
ifres7670 = False


}

var ifres7669 Obj

if True == ifres7670 {
ifres7669 = True


} else {
ifres7669 = False


}

ifres7668 = ifres7669


} else {
ifres7668 = False


}

if True == ifres7668 {
tmp7449 := GTl(__e, V3240)

tmp7450 := GHd(__e, tmp7449)

tmp7451 := GTl(__e, V3240)

tmp7452 := GTl(__e, tmp7451)

tmp7453 := GHd(__e, tmp7452)

tmp7454 := Call2(__e, SymFn(symshen_4process_1applications), tmp7453, V3241)


tmp7455 := GCons(__e, tmp7454, Nil)

tmp7456 := GCons(__e, tmp7450, tmp7455)

return GCons(__e, symlambda, tmp7456)


} else {
tmp7666 := GEq(__e, symlet, V3239)

var ifres7637 Obj

if True == tmp7666 {
tmp7665 := GConsP(__e, V3240)

var ifres7639 Obj

if True == tmp7665 {
tmp7663 := GHd(__e, V3240)

tmp7664 := GEq(__e, symlet, tmp7663)

var ifres7641 Obj

if True == tmp7664 {
tmp7661 := GTl(__e, V3240)

tmp7662 := GConsP(__e, tmp7661)

var ifres7643 Obj

if True == tmp7662 {
tmp7658 := GTl(__e, V3240)

tmp7659 := GTl(__e, tmp7658)

tmp7660 := GConsP(__e, tmp7659)

var ifres7645 Obj

if True == tmp7660 {
tmp7654 := GTl(__e, V3240)

tmp7655 := GTl(__e, tmp7654)

tmp7656 := GTl(__e, tmp7655)

tmp7657 := GConsP(__e, tmp7656)

var ifres7647 Obj

if True == tmp7657 {
tmp7649 := GTl(__e, V3240)

tmp7650 := GTl(__e, tmp7649)

tmp7651 := GTl(__e, tmp7650)

tmp7652 := GTl(__e, tmp7651)

tmp7653 := GEq(__e, Nil, tmp7652)

var ifres7648 Obj

if True == tmp7653 {
ifres7648 = True


} else {
ifres7648 = False


}

ifres7647 = ifres7648


} else {
ifres7647 = False


}

var ifres7646 Obj

if True == ifres7647 {
ifres7646 = True


} else {
ifres7646 = False


}

ifres7645 = ifres7646


} else {
ifres7645 = False


}

var ifres7644 Obj

if True == ifres7645 {
ifres7644 = True


} else {
ifres7644 = False


}

ifres7643 = ifres7644


} else {
ifres7643 = False


}

var ifres7642 Obj

if True == ifres7643 {
ifres7642 = True


} else {
ifres7642 = False


}

ifres7641 = ifres7642


} else {
ifres7641 = False


}

var ifres7640 Obj

if True == ifres7641 {
ifres7640 = True


} else {
ifres7640 = False


}

ifres7639 = ifres7640


} else {
ifres7639 = False


}

var ifres7638 Obj

if True == ifres7639 {
ifres7638 = True


} else {
ifres7638 = False


}

ifres7637 = ifres7638


} else {
ifres7637 = False


}

if True == ifres7637 {
tmp7457 := GTl(__e, V3240)

tmp7458 := GHd(__e, tmp7457)

tmp7459 := GTl(__e, V3240)

tmp7460 := GTl(__e, tmp7459)

tmp7461 := GHd(__e, tmp7460)

tmp7462 := Call2(__e, SymFn(symshen_4process_1applications), tmp7461, V3241)


tmp7463 := GTl(__e, V3240)

tmp7464 := GTl(__e, tmp7463)

tmp7465 := GTl(__e, tmp7464)

tmp7466 := GHd(__e, tmp7465)

tmp7467 := Call2(__e, SymFn(symshen_4process_1applications), tmp7466, V3241)


tmp7468 := GCons(__e, tmp7467, Nil)

tmp7469 := GCons(__e, tmp7462, tmp7468)

tmp7470 := GCons(__e, tmp7458, tmp7469)

return GCons(__e, symlet, tmp7470)


} else {
tmp7635 := GEq(__e, symdefun, V3239)

var ifres7606 Obj

if True == tmp7635 {
tmp7634 := GConsP(__e, V3240)

var ifres7608 Obj

if True == tmp7634 {
tmp7632 := GHd(__e, V3240)

tmp7633 := GEq(__e, symdefun, tmp7632)

var ifres7610 Obj

if True == tmp7633 {
tmp7630 := GTl(__e, V3240)

tmp7631 := GConsP(__e, tmp7630)

var ifres7612 Obj

if True == tmp7631 {
tmp7627 := GTl(__e, V3240)

tmp7628 := GTl(__e, tmp7627)

tmp7629 := GConsP(__e, tmp7628)

var ifres7614 Obj

if True == tmp7629 {
tmp7623 := GTl(__e, V3240)

tmp7624 := GTl(__e, tmp7623)

tmp7625 := GTl(__e, tmp7624)

tmp7626 := GConsP(__e, tmp7625)

var ifres7616 Obj

if True == tmp7626 {
tmp7618 := GTl(__e, V3240)

tmp7619 := GTl(__e, tmp7618)

tmp7620 := GTl(__e, tmp7619)

tmp7621 := GTl(__e, tmp7620)

tmp7622 := GEq(__e, Nil, tmp7621)

var ifres7617 Obj

if True == tmp7622 {
ifres7617 = True


} else {
ifres7617 = False


}

ifres7616 = ifres7617


} else {
ifres7616 = False


}

var ifres7615 Obj

if True == ifres7616 {
ifres7615 = True


} else {
ifres7615 = False


}

ifres7614 = ifres7615


} else {
ifres7614 = False


}

var ifres7613 Obj

if True == ifres7614 {
ifres7613 = True


} else {
ifres7613 = False


}

ifres7612 = ifres7613


} else {
ifres7612 = False


}

var ifres7611 Obj

if True == ifres7612 {
ifres7611 = True


} else {
ifres7611 = False


}

ifres7610 = ifres7611


} else {
ifres7610 = False


}

var ifres7609 Obj

if True == ifres7610 {
ifres7609 = True


} else {
ifres7609 = False


}

ifres7608 = ifres7609


} else {
ifres7608 = False


}

var ifres7607 Obj

if True == ifres7608 {
ifres7607 = True


} else {
ifres7607 = False


}

ifres7606 = ifres7607


} else {
ifres7606 = False


}

if True == ifres7606 {
return V3240
} else {
tmp7604 := GEq(__e, symdefine, V3239)

var ifres7582 Obj

if True == tmp7604 {
tmp7603 := GConsP(__e, V3240)

var ifres7584 Obj

if True == tmp7603 {
tmp7601 := GHd(__e, V3240)

tmp7602 := GEq(__e, symdefine, tmp7601)

var ifres7586 Obj

if True == tmp7602 {
tmp7599 := GTl(__e, V3240)

tmp7600 := GConsP(__e, tmp7599)

var ifres7588 Obj

if True == tmp7600 {
tmp7596 := GTl(__e, V3240)

tmp7597 := GTl(__e, tmp7596)

tmp7598 := GConsP(__e, tmp7597)

var ifres7590 Obj

if True == tmp7598 {
tmp7592 := GTl(__e, V3240)

tmp7593 := GTl(__e, tmp7592)

tmp7594 := GHd(__e, tmp7593)

tmp7595 := GEq(__e, sym_i, tmp7594)

var ifres7591 Obj

if True == tmp7595 {
ifres7591 = True


} else {
ifres7591 = False


}

ifres7590 = ifres7591


} else {
ifres7590 = False


}

var ifres7589 Obj

if True == ifres7590 {
ifres7589 = True


} else {
ifres7589 = False


}

ifres7588 = ifres7589


} else {
ifres7588 = False


}

var ifres7587 Obj

if True == ifres7588 {
ifres7587 = True


} else {
ifres7587 = False


}

ifres7586 = ifres7587


} else {
ifres7586 = False


}

var ifres7585 Obj

if True == ifres7586 {
ifres7585 = True


} else {
ifres7585 = False


}

ifres7584 = ifres7585


} else {
ifres7584 = False


}

var ifres7583 Obj

if True == ifres7584 {
ifres7583 = True


} else {
ifres7583 = False


}

ifres7582 = ifres7583


} else {
ifres7582 = False


}

if True == ifres7582 {
tmp7471 := GTl(__e, V3240)

tmp7472 := GHd(__e, tmp7471)

tmp7473 := GTl(__e, V3240)

tmp7474 := GHd(__e, tmp7473)

tmp7475 := GTl(__e, V3240)

tmp7476 := GTl(__e, tmp7475)

tmp7477 := GTl(__e, tmp7476)

tmp7478 := Call3(__e, SymFn(symshen_4process_1after_1type), tmp7474, tmp7477, V3241)


tmp7479 := GCons(__e, sym_i, tmp7478)

tmp7480 := GCons(__e, tmp7472, tmp7479)

return GCons(__e, symdefine, tmp7480)


} else {
tmp7580 := GEq(__e, symdefine, V3239)

var ifres7569 Obj

if True == tmp7580 {
tmp7579 := GConsP(__e, V3240)

var ifres7571 Obj

if True == tmp7579 {
tmp7577 := GHd(__e, V3240)

tmp7578 := GEq(__e, symdefine, tmp7577)

var ifres7573 Obj

if True == tmp7578 {
tmp7575 := GTl(__e, V3240)

tmp7576 := GConsP(__e, tmp7575)

var ifres7574 Obj

if True == tmp7576 {
ifres7574 = True


} else {
ifres7574 = False


}

ifres7573 = ifres7574


} else {
ifres7573 = False


}

var ifres7572 Obj

if True == ifres7573 {
ifres7572 = True


} else {
ifres7572 = False


}

ifres7571 = ifres7572


} else {
ifres7571 = False


}

var ifres7570 Obj

if True == ifres7571 {
ifres7570 = True


} else {
ifres7570 = False


}

ifres7569 = ifres7570


} else {
ifres7569 = False


}

if True == ifres7569 {
tmp7481 := GTl(__e, V3240)

tmp7482 := GHd(__e, tmp7481)

tmp7483 := MakeNative1(func(__e *ControlFlow, Z3242 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4process_1applications), Z3242, V3241)
panic("unreachable")
})

tmp7484 := GTl(__e, V3240)

tmp7485 := GTl(__e, tmp7484)

tmp7486 := Call2(__e, SymFn(symmap), tmp7483, tmp7485)


tmp7487 := GCons(__e, tmp7482, tmp7486)

return GCons(__e, symdefine, tmp7487)


} else {
tmp7567 := GEq(__e, symsynonyms, V3239)

if True == tmp7567 {
return GCons(__e, symsynonyms, V3240)
} else {
tmp7565 := GEq(__e, symtype, V3239)

var ifres7543 Obj

if True == tmp7565 {
tmp7564 := GConsP(__e, V3240)

var ifres7545 Obj

if True == tmp7564 {
tmp7562 := GHd(__e, V3240)

tmp7563 := GEq(__e, symtype, tmp7562)

var ifres7547 Obj

if True == tmp7563 {
tmp7560 := GTl(__e, V3240)

tmp7561 := GConsP(__e, tmp7560)

var ifres7549 Obj

if True == tmp7561 {
tmp7557 := GTl(__e, V3240)

tmp7558 := GTl(__e, tmp7557)

tmp7559 := GConsP(__e, tmp7558)

var ifres7551 Obj

if True == tmp7559 {
tmp7553 := GTl(__e, V3240)

tmp7554 := GTl(__e, tmp7553)

tmp7555 := GTl(__e, tmp7554)

tmp7556 := GEq(__e, Nil, tmp7555)

var ifres7552 Obj

if True == tmp7556 {
ifres7552 = True


} else {
ifres7552 = False


}

ifres7551 = ifres7552


} else {
ifres7551 = False


}

var ifres7550 Obj

if True == ifres7551 {
ifres7550 = True


} else {
ifres7550 = False


}

ifres7549 = ifres7550


} else {
ifres7549 = False


}

var ifres7548 Obj

if True == ifres7549 {
ifres7548 = True


} else {
ifres7548 = False


}

ifres7547 = ifres7548


} else {
ifres7547 = False


}

var ifres7546 Obj

if True == ifres7547 {
ifres7546 = True


} else {
ifres7546 = False


}

ifres7545 = ifres7546


} else {
ifres7545 = False


}

var ifres7544 Obj

if True == ifres7545 {
ifres7544 = True


} else {
ifres7544 = False


}

ifres7543 = ifres7544


} else {
ifres7543 = False


}

if True == ifres7543 {
tmp7488 := GTl(__e, V3240)

tmp7489 := GHd(__e, tmp7488)

tmp7490 := Call2(__e, SymFn(symshen_4process_1applications), tmp7489, V3241)


tmp7491 := GTl(__e, V3240)

tmp7492 := GTl(__e, tmp7491)

tmp7493 := GCons(__e, tmp7490, tmp7492)

return GCons(__e, symtype, tmp7493)


} else {
tmp7541 := GEq(__e, syminput_7, V3239)

var ifres7519 Obj

if True == tmp7541 {
tmp7540 := GConsP(__e, V3240)

var ifres7521 Obj

if True == tmp7540 {
tmp7538 := GHd(__e, V3240)

tmp7539 := GEq(__e, syminput_7, tmp7538)

var ifres7523 Obj

if True == tmp7539 {
tmp7536 := GTl(__e, V3240)

tmp7537 := GConsP(__e, tmp7536)

var ifres7525 Obj

if True == tmp7537 {
tmp7533 := GTl(__e, V3240)

tmp7534 := GTl(__e, tmp7533)

tmp7535 := GConsP(__e, tmp7534)

var ifres7527 Obj

if True == tmp7535 {
tmp7529 := GTl(__e, V3240)

tmp7530 := GTl(__e, tmp7529)

tmp7531 := GTl(__e, tmp7530)

tmp7532 := GEq(__e, Nil, tmp7531)

var ifres7528 Obj

if True == tmp7532 {
ifres7528 = True


} else {
ifres7528 = False


}

ifres7527 = ifres7528


} else {
ifres7527 = False


}

var ifres7526 Obj

if True == ifres7527 {
ifres7526 = True


} else {
ifres7526 = False


}

ifres7525 = ifres7526


} else {
ifres7525 = False


}

var ifres7524 Obj

if True == ifres7525 {
ifres7524 = True


} else {
ifres7524 = False


}

ifres7523 = ifres7524


} else {
ifres7523 = False


}

var ifres7522 Obj

if True == ifres7523 {
ifres7522 = True


} else {
ifres7522 = False


}

ifres7521 = ifres7522


} else {
ifres7521 = False


}

var ifres7520 Obj

if True == ifres7521 {
ifres7520 = True


} else {
ifres7520 = False


}

ifres7519 = ifres7520


} else {
ifres7519 = False


}

if True == ifres7519 {
tmp7494 := GTl(__e, V3240)

tmp7495 := GHd(__e, tmp7494)

tmp7496 := GTl(__e, V3240)

tmp7497 := GTl(__e, tmp7496)

tmp7498 := GHd(__e, tmp7497)

tmp7499 := Call2(__e, SymFn(symshen_4process_1applications), tmp7498, V3241)


tmp7500 := GCons(__e, tmp7499, Nil)

tmp7501 := GCons(__e, tmp7495, tmp7500)

return GCons(__e, syminput_7, tmp7501)


} else {
tmp7517 := GConsP(__e, V3240)

var ifres7513 Obj

if True == tmp7517 {
tmp7515 := GHd(__e, V3240)

tmp7516 := Call1(__e, SymFn(symshen_4special_2), tmp7515)


var ifres7514 Obj

if True == tmp7516 {
ifres7514 = True


} else {
ifres7514 = False


}

ifres7513 = ifres7514


} else {
ifres7513 = False


}

if True == ifres7513 {
tmp7502 := GHd(__e, V3240)

tmp7503 := MakeNative1(func(__e *ControlFlow, Z3243 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4process_1applications), Z3243, V3241)
panic("unreachable")
})

tmp7504 := GTl(__e, V3240)

tmp7505 := Call2(__e, SymFn(symmap), tmp7503, tmp7504)


return GCons(__e, tmp7502, tmp7505)


} else {
tmp7511 := GConsP(__e, V3240)

var ifres7507 Obj

if True == tmp7511 {
tmp7509 := GHd(__e, V3240)

tmp7510 := Call1(__e, SymFn(symshen_4extraspecial_2), tmp7509)


var ifres7508 Obj

if True == tmp7510 {
ifres7508 = True


} else {
ifres7508 = False


}

ifres7507 = ifres7508


} else {
ifres7507 = False


}

if True == ifres7507 {
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

tmp7691 := Call2(__e, ns2_1set, symshen_4special_1case, tmp7448)


_ = tmp7691

tmp7692 := MakeNative2(func(__e *ControlFlow, V3246, V3247 Obj) Obj {
tmp7722 := GEq(__e, Nil, V3246)

if True == tmp7722 {
return Nil
} else {
tmp7720 := GConsP(__e, V3246)

var ifres7705 Obj

if True == tmp7720 {
tmp7718 := GHd(__e, V3246)

tmp7719 := GConsP(__e, tmp7718)

var ifres7707 Obj

if True == tmp7719 {
tmp7715 := GHd(__e, V3246)

tmp7716 := GTl(__e, tmp7715)

tmp7717 := GConsP(__e, tmp7716)

var ifres7709 Obj

if True == tmp7717 {
tmp7711 := GHd(__e, V3246)

tmp7712 := GTl(__e, tmp7711)

tmp7713 := GTl(__e, tmp7712)

tmp7714 := GEq(__e, Nil, tmp7713)

var ifres7710 Obj

if True == tmp7714 {
ifres7710 = True


} else {
ifres7710 = False


}

ifres7709 = ifres7710


} else {
ifres7709 = False


}

var ifres7708 Obj

if True == ifres7709 {
ifres7708 = True


} else {
ifres7708 = False


}

ifres7707 = ifres7708


} else {
ifres7707 = False


}

var ifres7706 Obj

if True == ifres7707 {
ifres7706 = True


} else {
ifres7706 = False


}

ifres7705 = ifres7706


} else {
ifres7705 = False


}

if True == ifres7705 {
tmp7693 := GHd(__e, V3246)

tmp7694 := GHd(__e, tmp7693)

tmp7695 := Call2(__e, SymFn(symshen_4process_1applications), tmp7694, V3247)


tmp7696 := GHd(__e, V3246)

tmp7697 := GTl(__e, tmp7696)

tmp7698 := GHd(__e, tmp7697)

tmp7699 := Call2(__e, SymFn(symshen_4process_1applications), tmp7698, V3247)


tmp7700 := GCons(__e, tmp7699, Nil)

tmp7701 := GCons(__e, tmp7695, tmp7700)

tmp7702 := GTl(__e, V3246)

tmp7703 := Call2(__e, SymFn(symshen_4process_1cond_1clauses), tmp7702, V3247)


return GCons(__e, tmp7701, tmp7703)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4process_1cond_1clauses)
}


}


panic("unreachable")
})

tmp7723 := Call2(__e, ns2_1set, symshen_4process_1cond_1clauses, tmp7692)


_ = tmp7723

tmp7724 := MakeNative3(func(__e *ControlFlow, V3250, V3251, V3252 Obj) Obj {
tmp7740 := GConsP(__e, V3251)

var ifres7736 Obj

if True == tmp7740 {
tmp7738 := GHd(__e, V3251)

tmp7739 := GEq(__e, sym_j, tmp7738)

var ifres7737 Obj

if True == tmp7739 {
ifres7737 = True


} else {
ifres7737 = False


}

ifres7736 = ifres7737


} else {
ifres7736 = False


}

if True == ifres7736 {
tmp7725 := MakeNative1(func(__e *ControlFlow, Z3253 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4process_1applications), Z3253, V3252)
panic("unreachable")
})

tmp7726 := GTl(__e, V3251)

tmp7727 := Call2(__e, SymFn(symmap), tmp7725, tmp7726)


return GCons(__e, sym_j, tmp7727)


} else {
tmp7734 := GConsP(__e, V3251)

if True == tmp7734 {
tmp7728 := GHd(__e, V3251)

tmp7729 := GTl(__e, V3251)

tmp7730 := Call3(__e, SymFn(symshen_4process_1after_1type), V3250, tmp7729, V3252)


return GCons(__e, tmp7728, tmp7730)


} else {
tmp7731 := Call3(__e, SymFn(symshen_4app), V3250, MakeString("\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("missing } in "), tmp7731))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("missing } in "), tmp7731)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


}


panic("unreachable")
})

tmp7741 := Call2(__e, ns2_1set, symshen_4process_1after_1type, tmp7724)


_ = tmp7741

tmp7742 := MakeNative2(func(__e *ControlFlow, V3254, V3255 Obj) Obj {
tmp7785 := GConsP(__e, V3254)

if True == tmp7785 {
tmp7743 := GHd(__e, V3254)

tmp7744 := Call1(__e, SymFn(symarity), tmp7743)


let__5426 := tmp7744
_ = let__5426

tmp7745 := GTl(__e, V3254)

tmp7746 := Call1(__e, SymFn(symlength), tmp7745)


let__5427 := tmp7746
_ = let__5427

tmp7783 := Call2(__e, SymFn(symelement_2), V3254, V3255)


if True == tmp7783 {
return V3254
} else {
tmp7780 := GHd(__e, V3254)

tmp7781 := Call1(__e, SymFn(symshen_4shen_1call_2), tmp7780)


if True == tmp7781 {
return V3254
} else {
tmp7778 := Call1(__e, SymFn(symshen_4foreign_2), V3254)


if True == tmp7778 {
return TailCall1(__e, SymFn(symshen_4unpack_1foreign), V3254)
} else {
tmp7776 := Call1(__e, SymFn(symshen_4fn_1call_2), V3254)


if True == tmp7776 {
return TailCall1(__e, SymFn(symshen_4fn_1call), V3254)
} else {
tmp7774 := Call1(__e, SymFn(symshen_4zero_1place_2), V3254)


if True == tmp7774 {
return V3254
} else {
tmp7771 := GHd(__e, V3254)

tmp7772 := Call2(__e, SymFn(symshen_4undefined_1f_2), tmp7771, let__5426)


if True == tmp7772 {
tmp7747 := GHd(__e, V3254)

tmp7748 := GCons(__e, tmp7747, Nil)

tmp7749 := GCons(__e, symfn, tmp7748)

tmp7750 := GTl(__e, V3254)

tmp7751 := GCons(__e, tmp7749, tmp7750)

return TailCall1(__e, SymFn(symshen_4simple_1curry), tmp7751)


} else {
tmp7768 := GHd(__e, V3254)

tmp7769 := PrimGuard1(__e, symvariable_2, PrimIsVariable, tmp7768)

if True == tmp7769 {
return TailCall1(__e, SymFn(symshen_4simple_1curry), V3254)
} else {
tmp7765 := GHd(__e, V3254)

tmp7766 := Call1(__e, SymFn(symshen_4application_2), tmp7765)


if True == tmp7766 {
return TailCall1(__e, SymFn(symshen_4simple_1curry), V3254)
} else {
tmp7762 := GHd(__e, V3254)

tmp7763 := Call3(__e, SymFn(symshen_4partial_1application_d_2), tmp7762, let__5426, let__5427)


if True == tmp7763 {
return TailCall2(__e, SymFn(symshen_4lambda_1function), V3254, GSub(__e, let__5426, let__5427))


} else {
tmp7759 := GHd(__e, V3254)

tmp7760 := Call3(__e, SymFn(symshen_4overapplication_2), tmp7759, let__5426, let__5427)


if True == tmp7760 {
tmp7753 := GHd(__e, V3254)

tmp7754 := GCons(__e, tmp7753, Nil)

tmp7755 := GCons(__e, symfn, tmp7754)

tmp7756 := GTl(__e, V3254)

tmp7757 := GCons(__e, tmp7755, tmp7756)

return TailCall1(__e, SymFn(symshen_4simple_1curry), tmp7757)


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

tmp7786 := Call2(__e, ns2_1set, symshen_4process_1application, tmp7742)


_ = tmp7786

tmp7787 := MakeNative1(func(__e *ControlFlow, V3258 Obj) Obj {
tmp7813 := GConsP(__e, V3258)

var ifres7793 Obj

if True == tmp7813 {
tmp7811 := GHd(__e, V3258)

tmp7812 := GConsP(__e, tmp7811)

var ifres7795 Obj

if True == tmp7812 {
tmp7808 := GHd(__e, V3258)

tmp7809 := GHd(__e, tmp7808)

tmp7810 := GEq(__e, symforeign, tmp7809)

var ifres7797 Obj

if True == tmp7810 {
tmp7805 := GHd(__e, V3258)

tmp7806 := GTl(__e, tmp7805)

tmp7807 := GConsP(__e, tmp7806)

var ifres7799 Obj

if True == tmp7807 {
tmp7801 := GHd(__e, V3258)

tmp7802 := GTl(__e, tmp7801)

tmp7803 := GTl(__e, tmp7802)

tmp7804 := GEq(__e, Nil, tmp7803)

var ifres7800 Obj

if True == tmp7804 {
ifres7800 = True


} else {
ifres7800 = False


}

ifres7799 = ifres7800


} else {
ifres7799 = False


}

var ifres7798 Obj

if True == ifres7799 {
ifres7798 = True


} else {
ifres7798 = False


}

ifres7797 = ifres7798


} else {
ifres7797 = False


}

var ifres7796 Obj

if True == ifres7797 {
ifres7796 = True


} else {
ifres7796 = False


}

ifres7795 = ifres7796


} else {
ifres7795 = False


}

var ifres7794 Obj

if True == ifres7795 {
ifres7794 = True


} else {
ifres7794 = False


}

ifres7793 = ifres7794


} else {
ifres7793 = False


}

if True == ifres7793 {
tmp7788 := GHd(__e, V3258)

tmp7789 := GTl(__e, tmp7788)

tmp7790 := GHd(__e, tmp7789)

tmp7791 := GTl(__e, V3258)

return GCons(__e, tmp7790, tmp7791)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4unpack_1foreign)
}


panic("unreachable")
})

tmp7814 := Call2(__e, ns2_1set, symshen_4unpack_1foreign, tmp7787)


_ = tmp7814

tmp7815 := MakeNative1(func(__e *ControlFlow, V3261 Obj) Obj {
tmp7837 := GConsP(__e, V3261)

var ifres7817 Obj

if True == tmp7837 {
tmp7835 := GHd(__e, V3261)

tmp7836 := GConsP(__e, tmp7835)

var ifres7819 Obj

if True == tmp7836 {
tmp7832 := GHd(__e, V3261)

tmp7833 := GHd(__e, tmp7832)

tmp7834 := GEq(__e, symforeign, tmp7833)

var ifres7821 Obj

if True == tmp7834 {
tmp7829 := GHd(__e, V3261)

tmp7830 := GTl(__e, tmp7829)

tmp7831 := GConsP(__e, tmp7830)

var ifres7823 Obj

if True == tmp7831 {
tmp7825 := GHd(__e, V3261)

tmp7826 := GTl(__e, tmp7825)

tmp7827 := GTl(__e, tmp7826)

tmp7828 := GEq(__e, Nil, tmp7827)

var ifres7824 Obj

if True == tmp7828 {
ifres7824 = True


} else {
ifres7824 = False


}

ifres7823 = ifres7824


} else {
ifres7823 = False


}

var ifres7822 Obj

if True == ifres7823 {
ifres7822 = True


} else {
ifres7822 = False


}

ifres7821 = ifres7822


} else {
ifres7821 = False


}

var ifres7820 Obj

if True == ifres7821 {
ifres7820 = True


} else {
ifres7820 = False


}

ifres7819 = ifres7820


} else {
ifres7819 = False


}

var ifres7818 Obj

if True == ifres7819 {
ifres7818 = True


} else {
ifres7818 = False


}

ifres7817 = ifres7818


} else {
ifres7817 = False


}

if True == ifres7817 {
return True
} else {
return False
}


panic("unreachable")
})

tmp7838 := Call2(__e, ns2_1set, symshen_4foreign_2, tmp7815)


_ = tmp7838

tmp7839 := MakeNative1(func(__e *ControlFlow, V3264 Obj) Obj {
tmp7845 := GConsP(__e, V3264)

var ifres7841 Obj

if True == tmp7845 {
tmp7843 := GTl(__e, V3264)

tmp7844 := GEq(__e, Nil, tmp7843)

var ifres7842 Obj

if True == tmp7844 {
ifres7842 = True


} else {
ifres7842 = False


}

ifres7841 = ifres7842


} else {
ifres7841 = False


}

if True == ifres7841 {
return True
} else {
return False
}


panic("unreachable")
})

tmp7846 := Call2(__e, ns2_1set, symshen_4zero_1place_2, tmp7839)


_ = tmp7846

tmp7847 := MakeNative1(func(__e *ControlFlow, V3265 Obj) Obj {
tmp7852 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V3265)

if True == tmp7852 {
tmp7849 := PrimGuard1(__e, symstr, PrimStr, V3265)

tmp7850 := Call1(__e, SymFn(symshen_4internal_1to_1shen_2), tmp7849)


if True == tmp7850 {
return True
} else {
return False
}


} else {
return False
}


panic("unreachable")
})

tmp7853 := Call2(__e, ns2_1set, symshen_4shen_1call_2, tmp7847)


_ = tmp7853

tmp7854 := MakeNative1(func(__e *ControlFlow, V3270 Obj) Obj {
tmp7884 := GConsP(__e, V3270)

var ifres7871 Obj

if True == tmp7884 {
tmp7882 := GHd(__e, V3270)

tmp7883 := GEq(__e, symprotect, tmp7882)

var ifres7873 Obj

if True == tmp7883 {
tmp7880 := GTl(__e, V3270)

tmp7881 := GConsP(__e, tmp7880)

var ifres7875 Obj

if True == tmp7881 {
tmp7877 := GTl(__e, V3270)

tmp7878 := GTl(__e, tmp7877)

tmp7879 := GEq(__e, Nil, tmp7878)

var ifres7876 Obj

if True == tmp7879 {
ifres7876 = True


} else {
ifres7876 = False


}

ifres7875 = ifres7876


} else {
ifres7875 = False


}

var ifres7874 Obj

if True == ifres7875 {
ifres7874 = True


} else {
ifres7874 = False


}

ifres7873 = ifres7874


} else {
ifres7873 = False


}

var ifres7872 Obj

if True == ifres7873 {
ifres7872 = True


} else {
ifres7872 = False


}

ifres7871 = ifres7872


} else {
ifres7871 = False


}

if True == ifres7871 {
return False
} else {
tmp7869 := GConsP(__e, V3270)

var ifres7856 Obj

if True == tmp7869 {
tmp7867 := GHd(__e, V3270)

tmp7868 := GEq(__e, symforeign, tmp7867)

var ifres7858 Obj

if True == tmp7868 {
tmp7865 := GTl(__e, V3270)

tmp7866 := GConsP(__e, tmp7865)

var ifres7860 Obj

if True == tmp7866 {
tmp7862 := GTl(__e, V3270)

tmp7863 := GTl(__e, tmp7862)

tmp7864 := GEq(__e, Nil, tmp7863)

var ifres7861 Obj

if True == tmp7864 {
ifres7861 = True


} else {
ifres7861 = False


}

ifres7860 = ifres7861


} else {
ifres7860 = False


}

var ifres7859 Obj

if True == ifres7860 {
ifres7859 = True


} else {
ifres7859 = False


}

ifres7858 = ifres7859


} else {
ifres7858 = False


}

var ifres7857 Obj

if True == ifres7858 {
ifres7857 = True


} else {
ifres7857 = False


}

ifres7856 = ifres7857


} else {
ifres7856 = False


}

if True == ifres7856 {
return False
} else {
return GConsP(__e, V3270)
}


}


panic("unreachable")
})

tmp7885 := Call2(__e, ns2_1set, symshen_4application_2, tmp7854)


_ = tmp7885

tmp7886 := MakeNative2(func(__e *ControlFlow, V3275, V3276 Obj) Obj {
tmp7894 := GEq(__e, MakeInteger(-1), V3276)

if True == tmp7894 {
tmp7892 := Call1(__e, SymFn(symshen_4lowercase_1symbol_2), V3275)


if True == tmp7892 {
tmp7888 := Call1(__e, SymFn(symexternal), symshen)


tmp7889 := Call2(__e, SymFn(symelement_2), V3275, tmp7888)


if True == GNot(__e, tmp7889) {
return True
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

tmp7895 := Call2(__e, ns2_1set, symshen_4undefined_1f_2, tmp7886)


_ = tmp7895

tmp7896 := MakeNative1(func(__e *ControlFlow, V3277 Obj) Obj {
tmp7901 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V3277)

if True == tmp7901 {
tmp7898 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V3277)

if True == GNot(__e, tmp7898) {
return True
} else {
return False
}


} else {
return False
}


panic("unreachable")
})

tmp7902 := Call2(__e, ns2_1set, symshen_4lowercase_1symbol_2, tmp7896)


_ = tmp7902

tmp7903 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V3278 := __self1
_ = V3278
tmp7933 := GConsP(__e, V3278)

var ifres7924 Obj

if True == tmp7933 {
tmp7931 := GTl(__e, V3278)

tmp7932 := GConsP(__e, tmp7931)

var ifres7926 Obj

if True == tmp7932 {
tmp7928 := GTl(__e, V3278)

tmp7929 := GTl(__e, tmp7928)

tmp7930 := GEq(__e, Nil, tmp7929)

var ifres7927 Obj

if True == tmp7930 {
ifres7927 = True


} else {
ifres7927 = False


}

ifres7926 = ifres7927


} else {
ifres7926 = False


}

var ifres7925 Obj

if True == ifres7926 {
ifres7925 = True


} else {
ifres7925 = False


}

ifres7924 = ifres7925


} else {
ifres7924 = False


}

if True == ifres7924 {
return V3278
} else {
tmp7922 := GConsP(__e, V3278)

var ifres7913 Obj

if True == tmp7922 {
tmp7920 := GTl(__e, V3278)

tmp7921 := GConsP(__e, tmp7920)

var ifres7915 Obj

if True == tmp7921 {
tmp7917 := GTl(__e, V3278)

tmp7918 := GTl(__e, tmp7917)

tmp7919 := GConsP(__e, tmp7918)

var ifres7916 Obj

if True == tmp7919 {
ifres7916 = True


} else {
ifres7916 = False


}

ifres7915 = ifres7916


} else {
ifres7915 = False


}

var ifres7914 Obj

if True == ifres7915 {
ifres7914 = True


} else {
ifres7914 = False


}

ifres7913 = ifres7914


} else {
ifres7913 = False


}

if True == ifres7913 {
tmp7904 := GHd(__e, V3278)

tmp7905 := GTl(__e, V3278)

tmp7906 := GHd(__e, tmp7905)

tmp7907 := GCons(__e, tmp7906, Nil)

tmp7908 := GCons(__e, tmp7904, tmp7907)

tmp7909 := GTl(__e, V3278)

tmp7910 := GTl(__e, tmp7909)

tmp7911 := GCons(__e, tmp7908, tmp7910)

if SymFn(symshen_4simple_1curry) == __self {
__self1 = tmp7911
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4simple_1curry), tmp7911)


} else {
return V3278
}


}


panic("unreachable")
})
return __self
})()

tmp7934 := Call2(__e, ns2_1set, symshen_4simple_1curry, tmp7903)


_ = tmp7934

tmp7935 := MakeNative1(func(__e *ControlFlow, V3279 Obj) Obj {
return TailCall1(__e, SymFn(symfn), V3279)
panic("unreachable")
})

tmp7936 := Call2(__e, ns2_1set, symfunction, tmp7935)


_ = tmp7936

tmp7937 := MakeNative1(func(__e *ControlFlow, V3280 Obj) Obj {
tmp7944 := Call1(__e, SymFn(symarity), V3280)


tmp7945 := GEq(__e, tmp7944, MakeInteger(0))

if True == tmp7945 {
return TailCall0(__e, V3280)
} else {
tmp7938 := MakeNative0(func(__e *ControlFlow) Obj {
tmp7939 := GValue(__e, sym_dproperty_1vector_d)

return TailCall3(__e, SymFn(symget), V3280, symshen_4lambda_1form, tmp7939)


panic("unreachable")
})

tmp7940 := MakeNative1(func(__e *ControlFlow, Z3281 Obj) Obj {
tmp7941 := Call3(__e, SymFn(symshen_4app), V3280, MakeString(" is undefined\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("fn: "), tmp7941))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("fn: "), tmp7941)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp7938, tmp7940)


}


panic("unreachable")
})

tmp7946 := Call2(__e, ns2_1set, symfn, tmp7937)


_ = tmp7946

tmp7947 := MakeNative1(func(__e *ControlFlow, V3284 Obj) Obj {
tmp7977 := GConsP(__e, V3284)

var ifres7964 Obj

if True == tmp7977 {
tmp7975 := GHd(__e, V3284)

tmp7976 := GEq(__e, symfn, tmp7975)

var ifres7966 Obj

if True == tmp7976 {
tmp7973 := GTl(__e, V3284)

tmp7974 := GConsP(__e, tmp7973)

var ifres7968 Obj

if True == tmp7974 {
tmp7970 := GTl(__e, V3284)

tmp7971 := GTl(__e, tmp7970)

tmp7972 := GEq(__e, Nil, tmp7971)

var ifres7969 Obj

if True == tmp7972 {
ifres7969 = True


} else {
ifres7969 = False


}

ifres7968 = ifres7969


} else {
ifres7968 = False


}

var ifres7967 Obj

if True == ifres7968 {
ifres7967 = True


} else {
ifres7967 = False


}

ifres7966 = ifres7967


} else {
ifres7966 = False


}

var ifres7965 Obj

if True == ifres7966 {
ifres7965 = True


} else {
ifres7965 = False


}

ifres7964 = ifres7965


} else {
ifres7964 = False


}

if True == ifres7964 {
return True
} else {
tmp7962 := GConsP(__e, V3284)

var ifres7949 Obj

if True == tmp7962 {
tmp7960 := GHd(__e, V3284)

tmp7961 := GEq(__e, symfunction, tmp7960)

var ifres7951 Obj

if True == tmp7961 {
tmp7958 := GTl(__e, V3284)

tmp7959 := GConsP(__e, tmp7958)

var ifres7953 Obj

if True == tmp7959 {
tmp7955 := GTl(__e, V3284)

tmp7956 := GTl(__e, tmp7955)

tmp7957 := GEq(__e, Nil, tmp7956)

var ifres7954 Obj

if True == tmp7957 {
ifres7954 = True


} else {
ifres7954 = False


}

ifres7953 = ifres7954


} else {
ifres7953 = False


}

var ifres7952 Obj

if True == ifres7953 {
ifres7952 = True


} else {
ifres7952 = False


}

ifres7951 = ifres7952


} else {
ifres7951 = False


}

var ifres7950 Obj

if True == ifres7951 {
ifres7950 = True


} else {
ifres7950 = False


}

ifres7949 = ifres7950


} else {
ifres7949 = False


}

if True == ifres7949 {
return True
} else {
return False
}


}


panic("unreachable")
})

tmp7978 := Call2(__e, ns2_1set, symshen_4fn_1call_2, tmp7947)


_ = tmp7978

tmp7979 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V3285 := __self1
_ = V3285
tmp8019 := GConsP(__e, V3285)

var ifres8006 Obj

if True == tmp8019 {
tmp8017 := GHd(__e, V3285)

tmp8018 := GEq(__e, symfunction, tmp8017)

var ifres8008 Obj

if True == tmp8018 {
tmp8015 := GTl(__e, V3285)

tmp8016 := GConsP(__e, tmp8015)

var ifres8010 Obj

if True == tmp8016 {
tmp8012 := GTl(__e, V3285)

tmp8013 := GTl(__e, tmp8012)

tmp8014 := GEq(__e, Nil, tmp8013)

var ifres8011 Obj

if True == tmp8014 {
ifres8011 = True


} else {
ifres8011 = False


}

ifres8010 = ifres8011


} else {
ifres8010 = False


}

var ifres8009 Obj

if True == ifres8010 {
ifres8009 = True


} else {
ifres8009 = False


}

ifres8008 = ifres8009


} else {
ifres8008 = False


}

var ifres8007 Obj

if True == ifres8008 {
ifres8007 = True


} else {
ifres8007 = False


}

ifres8006 = ifres8007


} else {
ifres8006 = False


}

if True == ifres8006 {
tmp7980 := GTl(__e, V3285)

tmp7981 := GCons(__e, symfn, tmp7980)

if SymFn(symshen_4fn_1call) == __self {
__self1 = tmp7981
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4fn_1call), tmp7981)


} else {
tmp8004 := GConsP(__e, V3285)

var ifres7991 Obj

if True == tmp8004 {
tmp8002 := GHd(__e, V3285)

tmp8003 := GEq(__e, symfn, tmp8002)

var ifres7993 Obj

if True == tmp8003 {
tmp8000 := GTl(__e, V3285)

tmp8001 := GConsP(__e, tmp8000)

var ifres7995 Obj

if True == tmp8001 {
tmp7997 := GTl(__e, V3285)

tmp7998 := GTl(__e, tmp7997)

tmp7999 := GEq(__e, Nil, tmp7998)

var ifres7996 Obj

if True == tmp7999 {
ifres7996 = True


} else {
ifres7996 = False


}

ifres7995 = ifres7996


} else {
ifres7995 = False


}

var ifres7994 Obj

if True == ifres7995 {
ifres7994 = True


} else {
ifres7994 = False


}

ifres7993 = ifres7994


} else {
ifres7993 = False


}

var ifres7992 Obj

if True == ifres7993 {
ifres7992 = True


} else {
ifres7992 = False


}

ifres7991 = ifres7992


} else {
ifres7991 = False


}

if True == ifres7991 {
tmp7982 := GTl(__e, V3285)

tmp7983 := GHd(__e, tmp7982)

tmp7984 := Call1(__e, SymFn(symarity), tmp7983)


let__5428 := tmp7984
_ = let__5428

tmp7989 := GEq(__e, let__5428, MakeInteger(-1))

if True == tmp7989 {
return V3285
} else {
tmp7987 := GEq(__e, let__5428, MakeInteger(0))

if True == tmp7987 {
return GTl(__e, V3285)
} else {
tmp7985 := GTl(__e, V3285)

return TailCall2(__e, SymFn(symshen_4lambda_1function), tmp7985, let__5428)


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

tmp8020 := Call2(__e, ns2_1set, symshen_4fn_1call, tmp7979)


_ = tmp8020

tmp8021 := MakeNative3(func(__e *ControlFlow, V3287, V3288, V3289 Obj) Obj {
let__5429 := GGt(__e, V3288, V3289)
_ = let__5429

var ifres8028 Obj

if True == let__5429 {
tmp8036 := Call0(__e, SymFn(symshen_4loading_2))


var ifres8030 Obj

if True == tmp8036 {
tmp8032 := GCons(__e, sym_1, Nil)

tmp8033 := GCons(__e, sym_7, tmp8032)

tmp8034 := Call2(__e, SymFn(symelement_2), V3287, tmp8033)


tmp8035 := GNot(__e, tmp8034)

var ifres8031 Obj

if True == tmp8035 {
ifres8031 = True


} else {
ifres8031 = False


}

ifres8030 = ifres8031


} else {
ifres8030 = False


}

var ifres8029 Obj

if True == ifres8030 {
ifres8029 = True


} else {
ifres8029 = False


}

ifres8028 = ifres8029


} else {
ifres8028 = False


}

var ifres8023 Obj

if True == ifres8028 {
tmp8024 := Call3(__e, SymFn(symshen_4app), V3287, MakeString("\n"), symshen_4a)


tmp8025 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("partial application of "), tmp8024)

tmp8026 := Call0(__e, SymFn(symstoutput))


tmp8027 := Call2(__e, SymFn(sympr), tmp8025, tmp8026)


ifres8023 = tmp8027


} else {
ifres8023 = symshen_4skip


}

let__5430 := ifres8023
_ = let__5430

return let__5429


panic("unreachable")
})

tmp8037 := Call2(__e, ns2_1set, symshen_4partial_1application_d_2, tmp8021)


_ = tmp8037

tmp8038 := MakeNative0(func(__e *ControlFlow) Obj {
return GValue(__e, symshen_4_dloading_2_d)
panic("unreachable")
})

tmp8039 := Call2(__e, ns2_1set, symshen_4loading_2, tmp8038)


_ = tmp8039

tmp8040 := MakeNative3(func(__e *ControlFlow, V3296, V3297, V3298 Obj) Obj {
tmp8056 := GEq(__e, MakeInteger(-1), V3297)

if True == tmp8056 {
return False
} else {
let__5431 := GLt(__e, V3297, V3298)
_ = let__5431

var ifres8052 Obj

if True == let__5431 {
tmp8054 := Call0(__e, SymFn(symshen_4loading_2))


var ifres8053 Obj

if True == tmp8054 {
ifres8053 = True


} else {
ifres8053 = False


}

ifres8052 = ifres8053


} else {
ifres8052 = False


}

var ifres8042 Obj

if True == ifres8052 {
tmp8044 := GEq(__e, V3298, MakeInteger(1))

var ifres8043 Obj

if True == tmp8044 {
ifres8043 = MakeString("")


} else {
ifres8043 = MakeString("s")


}

tmp8045 := Call3(__e, SymFn(symshen_4app), ifres8043, MakeString("\n"), symshen_4a)


tmp8047 := Call3(__e, SymFn(symshen_4app), V3298, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" argument"), tmp8045), symshen_4a)


tmp8049 := Call3(__e, SymFn(symshen_4app), V3296, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" might not like "), tmp8047), symshen_4a)


tmp8050 := Call0(__e, SymFn(symstoutput))


tmp8051 := Call2(__e, SymFn(sympr), tmp8049, tmp8050)


ifres8042 = tmp8051


} else {
ifres8042 = symshen_4skip


}

let__5432 := ifres8042
_ = let__5432

return let__5431


}


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4overapplication_2, tmp8040)




panic("unreachable")
})

