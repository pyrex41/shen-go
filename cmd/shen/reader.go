package main

import . "github.com/pyrex41/shen-go/kl"

var ReaderMain = MakeNative(func(__e *ControlFlow) {
tmp5453 := MakeNative(func(__e *ControlFlow) {
V2556 := __e.Get(1)
_ = V2556
tmp5454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1file_1as_1bytelist) {
return PrimReadFileAsByteList(V2556)
}
__typedArg0 := V2556
return Call(__e, PrimFunc(symread_1file_1as_1bytelist), __typedArg0)
})()

let__5030 := tmp5454
_ = let__5030

tmp5455 := MakeNative(func(__e *ControlFlow) {
tmp5456 := MakeNative(func(__e *ControlFlow) {
Z2559 := __e.Get(1)
_ = Z2559
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2559)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp5456, let__5030)
return


}, 0)

tmp5457 := MakeNative(func(__e *ControlFlow) {
Z2560 := __e.Get(1)
_ = Z2560
tmp5458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dresidue_d)
}
__typedArg0 := symshen_4_dresidue_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4reader_1error), tmp5458)
return


}, 1)

tmp5459 := Call(__e, try_1catch, tmp5455, tmp5457)


let__5031 := tmp5459
_ = let__5031

tmp5460 := Call(__e, PrimFunc(symshen_4process_1sexprs), let__5031)


let__5032 := tmp5460
_ = let__5032

__e.Return(let__5032)
return


}, 1)

tmp5461 := Call(__e, ns2_1set, symread_1file, tmp5453)


_ = tmp5461

tmp5462 := MakeNative(func(__e *ControlFlow) {
V2562 := __e.Get(1)
_ = V2562
tmp5463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmaximum_1print_1sequence_1size_d)
}
__typedArg0 := sym_dmaximum_1print_1sequence_1size_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp5464 := Call(__e, PrimFunc(symshen_4reader_1error_1message), tmp5463, MakeInteger(0), V2562)


tmp5466 := Call(__e, PrimFunc(symshen_4proc_1nl), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("reader error near here: "))
__typedS1, __typedOK1 := TypedString(tmp5464)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("reader error near here: ")
__typedArg1 := tmp5464
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp5466)
}
__typedArg0 := tmp5466
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

tmp5467 := Call(__e, ns2_1set, symshen_4reader_1error, tmp5462)


_ = tmp5467

tmp5468 := MakeNative(func(__e *ControlFlow) {
V2570 := __e.Get(1)
_ = V2570
V2571 := __e.Get(2)
_ = V2571
V2572 := __e.Get(3)
_ = V2572
tmp5479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2572)
}
__typedArg0 := Nil
__typedArg1 := V2572
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5479 {
__e.Return(MakeString(""))
return
} else {
tmp5477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V2570, V2571)
}
__typedArg0 := V2570
__typedArg1 := V2571
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5477 {
__e.Return(MakeString(""))
return
} else {
tmp5475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2572)
}
__typedArg0 := V2572
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5475 {
tmp5469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2572)
}
__typedArg0 := V2572
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(tmp5469)
}
__typedArg0 := tmp5469
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp5471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V2571)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V2571
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()

tmp5472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2572)
}
__typedArg0 := V2572
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5473 := Call(__e, PrimFunc(symshen_4reader_1error_1message), V2570, tmp5471, tmp5472)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp5470)
__typedS1, __typedOK1 := TypedString(tmp5473)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp5470
__typedArg1 := tmp5473
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4reader_1error_1message)
return
}


}


}


}, 3)

tmp5480 := Call(__e, ns2_1set, symshen_4reader_1error_1message, tmp5468)


_ = tmp5480

tmp5481 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dit_d)
}
__typedArg0 := symshen_4_dit_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp5482 := Call(__e, ns2_1set, symit, tmp5481)


_ = tmp5482

tmp5483 := MakeNative(func(__e *ControlFlow) {
V2573 := __e.Get(1)
_ = V2573
tmp5484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symopen) {
return PrimOpenStream(V2573, symin)
}
__typedArg0 := V2573
__typedArg1 := symin
return Call(__e, PrimFunc(symopen), __typedArg0, __typedArg1)
})()

let__5033 := tmp5484
_ = let__5033

tmp5485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(let__5033)
}
__typedArg0 := let__5033
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

let__5034 := tmp5485
_ = let__5034

tmp5486 := Call(__e, PrimFunc(symshen_4read_1file_1as_1bytelist_1help), let__5033, let__5034, Nil)


let__5035 := tmp5486
_ = let__5035

tmp5487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symclose) {
return PrimCloseStream(let__5033)
}
__typedArg0 := let__5033
return Call(__e, PrimFunc(symclose), __typedArg0)
})()

let__5036 := tmp5487
_ = let__5036

__e.TailApply(PrimFunc(symreverse), let__5035)
return


}, 1)

tmp5488 := Call(__e, ns2_1set, symread_1file_1as_1bytelist, tmp5483)


_ = tmp5488

tmp5489 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__selftop:
V2578 := __self1
_ = V2578
V2579 := __self2
_ = V2579
V2580 := __self3
_ = V2580
tmp5493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(-1), V2579)
}
__typedArg0 := MakeInteger(-1)
__typedArg1 := V2579
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5493 {
__e.Return(V2580)
return
} else {
tmp5490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(V2578)
}
__typedArg0 := V2578
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

tmp5491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2579, V2580)
}
__typedArg0 := V2579
__typedArg1 := V2580
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4read_1file_1as_1bytelist_1help) == __self {
__self1, __self2, __self3 = V2578, tmp5490, tmp5491
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4read_1file_1as_1bytelist_1help), V2578, tmp5490, tmp5491)
return


}


}, 3)

tmp5494 := Call(__e, ns2_1set, symshen_4read_1file_1as_1bytelist_1help, tmp5489)


_ = tmp5494

tmp5495 := MakeNative(func(__e *ControlFlow) {
V2581 := __e.Get(1)
_ = V2581
tmp5496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symopen) {
return PrimOpenStream(V2581, symin)
}
__typedArg0 := V2581
__typedArg1 := symin
return Call(__e, PrimFunc(symopen), __typedArg0, __typedArg1)
})()

let__5037 := tmp5496
_ = let__5037

tmp5497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(let__5037)
}
__typedArg0 := let__5037
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4rfas_1h), let__5037, tmp5497, MakeString(""))
return


}, 1)

tmp5498 := Call(__e, ns2_1set, symread_1file_1as_1string, tmp5495)


_ = tmp5498

tmp5499 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__selftop:
V2583 := __self1
_ = V2583
V2584 := __self2
_ = V2584
V2585 := __self3
_ = V2585
tmp5505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(-1), V2584)
}
__typedArg0 := MakeInteger(-1)
__typedArg1 := V2584
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5505 {
tmp5500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symclose) {
return PrimCloseStream(V2583)
}
__typedArg0 := V2583
return Call(__e, PrimFunc(symclose), __typedArg0)
})()

_ = tmp5500

__e.Return(V2585)
return


} else {
tmp5501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(V2583)
}
__typedArg0 := V2583
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

tmp5502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(V2584)
}
__typedArg0 := V2584
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

if PrimFunc(symshen_4rfas_1h) == __self {
__self1, __self2, __self3 = V2583, tmp5501, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(V2585)
__typedS1, __typedOK1 := TypedString(tmp5502)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V2585
__typedArg1 := tmp5502
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4rfas_1h), V2583, tmp5501, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(V2585)
__typedS1, __typedOK1 := TypedString(tmp5502)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V2585
__typedArg1 := tmp5502
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}


}, 3)

tmp5506 := Call(__e, ns2_1set, symshen_4rfas_1h, tmp5499)


_ = tmp5506

tmp5507 := MakeNative(func(__e *ControlFlow) {
V2586 := __e.Get(1)
_ = V2586
tmp5508 := Call(__e, PrimFunc(symread), V2586)


__e.TailApply(PrimFunc(symeval_1kl), tmp5508)
return


}, 1)

tmp5509 := Call(__e, ns2_1set, syminput, tmp5507)


_ = tmp5509

tmp5510 := MakeNative(func(__e *ControlFlow) {
V2587 := __e.Get(1)
_ = V2587
V2588 := __e.Get(2)
_ = V2588
tmp5511 := Call(__e, PrimFunc(symshen_4monotype), V2587)


let__5038 := tmp5511
_ = let__5038

tmp5512 := Call(__e, PrimFunc(symread), V2588)


let__5039 := tmp5512
_ = let__5039

tmp5518 := Call(__e, PrimFunc(symshen_4typecheck), let__5039, V2587)


tmp5519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, tmp5518)
}
__typedArg0 := False
__typedArg1 := tmp5518
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5519 {
tmp5513 := Call(__e, PrimFunc(symshen_4app), V2587, MakeString("\n"), symshen_4r)


tmp5515 := Call(__e, PrimFunc(symshen_4app), let__5039, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" is not of type "))
__typedS1, __typedOK1 := TypedString(tmp5513)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" is not of type ")
__typedArg1 := tmp5513
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4r)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("type error: "))
__typedS1, __typedOK1 := TypedString(tmp5515)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("type error: ")
__typedArg1 := tmp5515
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("type error: "))
__typedS1, __typedOK1 := TypedString(tmp5515)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("type error: ")
__typedArg1 := tmp5515
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
__e.TailApply(PrimFunc(symeval_1kl), let__5039)
return
}


}, 2)

tmp5520 := Call(__e, ns2_1set, symshen_4input_1h_7, tmp5510)


_ = tmp5520

tmp5521 := MakeNative(func(__e *ControlFlow) {
V2591 := __e.Get(1)
_ = V2591
tmp5528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2591)
}
__typedArg0 := V2591
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5528 {
tmp5522 := MakeNative(func(__e *ControlFlow) {
Z2592 := __e.Get(1)
_ = Z2592
__e.TailApply(PrimFunc(symshen_4monotype), Z2592)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp5522, V2591)
return


} else {
tmp5526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V2591)
}
__typedArg0 := V2591
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp5526 {
tmp5523 := Call(__e, PrimFunc(symshen_4app), V2591, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("input+ expects a monotype: not "))
__typedS1, __typedOK1 := TypedString(tmp5523)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("input+ expects a monotype: not ")
__typedArg1 := tmp5523
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("input+ expects a monotype: not "))
__typedS1, __typedOK1 := TypedString(tmp5523)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("input+ expects a monotype: not ")
__typedArg1 := tmp5523
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
__e.Return(V2591)
return
}


}


}, 1)

tmp5529 := Call(__e, ns2_1set, symshen_4monotype, tmp5521)


_ = tmp5529

tmp5530 := MakeNative(func(__e *ControlFlow) {
V2593 := __e.Get(1)
_ = V2593
tmp5531 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2593)


tmp5532 := MakeNative(func(__e *ControlFlow) {
Z2594 := __e.Get(1)
_ = Z2594
__e.TailApply(PrimFunc(symshen_4return_2), Z2594)
return
}, 1)

__e.TailApply(PrimFunc(symshen_4read_1loop), V2593, tmp5531, Nil, tmp5532)
return


}, 1)

tmp5533 := Call(__e, ns2_1set, symlineread, tmp5530)


_ = tmp5533

tmp5534 := MakeNative(func(__e *ControlFlow) {
V2595 := __e.Get(1)
_ = V2595
tmp5535 := Call(__e, PrimFunc(symshen_4str_1_6bytes), V2595)


let__5040 := tmp5535
_ = let__5040

tmp5536 := MakeNative(func(__e *ControlFlow) {
Z2598 := __e.Get(1)
_ = Z2598
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2598)
return
}, 1)

tmp5537 := Call(__e, PrimFunc(symcompile), tmp5536, let__5040)


let__5041 := tmp5537
_ = let__5041

tmp5538 := Call(__e, PrimFunc(symshen_4process_1sexprs), let__5041)


let__5042 := tmp5538
_ = let__5042

__e.Return(let__5042)
return


}, 1)

tmp5539 := Call(__e, ns2_1set, symread_1from_1string, tmp5534)


_ = tmp5539

tmp5540 := MakeNative(func(__e *ControlFlow) {
V2600 := __e.Get(1)
_ = V2600
tmp5541 := Call(__e, PrimFunc(symshen_4str_1_6bytes), V2600)


let__5043 := tmp5541
_ = let__5043

tmp5542 := MakeNative(func(__e *ControlFlow) {
Z2603 := __e.Get(1)
_ = Z2603
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2603)
return
}, 1)

tmp5543 := Call(__e, PrimFunc(symcompile), tmp5542, let__5043)


let__5044 := tmp5543
_ = let__5044

__e.Return(let__5044)
return


}, 1)

tmp5544 := Call(__e, ns2_1set, symread_1from_1string_1unprocessed, tmp5540)


_ = tmp5544

tmp5545 := MakeNative(func(__e *ControlFlow) {
V2604 := __e.Get(1)
_ = V2604
tmp5553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V2604)
}
__typedArg0 := MakeString("")
__typedArg1 := V2604
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5553 {
__e.Return(Nil)
return
} else {
tmp5551 := Call(__e, PrimFunc(symshen_4_7string_2), V2604)


if True == tmp5551 {
tmp5546 := Call(__e, PrimFunc(symhdstr), V2604)


tmp5547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp5546)
}
__typedArg0 := tmp5546
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp5549 := Call(__e, PrimFunc(symshen_4str_1_6bytes), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2604)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V2604
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5547, tmp5549)
}
__typedArg0 := tmp5547
__typedArg1 := tmp5549
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4str_1_6bytes)
return
}


}


}, 1)

tmp5554 := Call(__e, ns2_1set, symshen_4str_1_6bytes, tmp5545)


_ = tmp5554

tmp5555 := MakeNative(func(__e *ControlFlow) {
V2605 := __e.Get(1)
_ = V2605
tmp5556 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2605)


tmp5557 := MakeNative(func(__e *ControlFlow) {
Z2606 := __e.Get(1)
_ = Z2606
__e.TailApply(PrimFunc(symshen_4whitespace_2), Z2606)
return
}, 1)

tmp5558 := Call(__e, PrimFunc(symshen_4read_1loop), V2605, tmp5556, Nil, tmp5557)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5558)
}
__typedArg0 := tmp5558
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


}, 1)

tmp5559 := Call(__e, ns2_1set, symread, tmp5555)


_ = tmp5559

tmp5560 := MakeNative(func(__e *ControlFlow) {
V2607 := __e.Get(1)
_ = V2607
tmp5563 := Call(__e, PrimFunc(symchar_1stream_2), V2607)


if True == tmp5563 {
tmp5561 := Call(__e, PrimFunc(symread_1unit_1string), V2607)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp5561)
}
__typedArg0 := tmp5561
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(V2607)
}
__typedArg0 := V2607
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})())
return
}


}, 1)

tmp5564 := Call(__e, ns2_1set, symshen_4my_1read_1byte, tmp5560)


_ = tmp5564

tmp5565 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__selftop:
V2612 := __self1
_ = V2612
V2613 := __self2
_ = V2613
V2614 := __self3
_ = V2614
V2615 := __self4
_ = V2615
tmp5588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(94), V2613)
}
__typedArg0 := MakeInteger(94)
__typedArg1 := V2613
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5588 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("read aborted"))
}
__typedArg0 := MakeString("read aborted")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp5586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(-1), V2613)
}
__typedArg0 := MakeInteger(-1)
__typedArg1 := V2613
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5586 {
tmp5569 := Call(__e, PrimFunc(symempty_2), V2614)


if True == tmp5569 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("error: empty stream"))
}
__typedArg0 := MakeString("error: empty stream")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp5566 := MakeNative(func(__e *ControlFlow) {
Z2616 := __e.Get(1)
_ = Z2616
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2616)
return
}, 1)

tmp5567 := Call(__e, PrimFunc(symcompile), tmp5566, V2614)


__e.TailApply(PrimFunc(symshen_4process_1sexprs), tmp5567)
return


}


} else {
tmp5584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V2613)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V2613
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5584 {
tmp5570 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2612)


if PrimFunc(symshen_4read_1loop) == __self {
__self1, __self2, __self3, __self4 = V2612, tmp5570, V2614, V2615
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4read_1loop), V2612, tmp5570, V2614, V2615)
return


} else {
tmp5582 := Call(__e, V2615, V2613)


if True == tmp5582 {
tmp5571 := Call(__e, PrimFunc(symshen_4try_1parse), V2614)


let__5045 := tmp5571
_ = let__5045

tmp5577 := Call(__e, PrimFunc(symshen_4nothing_1doing_2), let__5045)


if True == tmp5577 {
tmp5572 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2612)


tmp5573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2613, Nil)
}
__typedArg0 := V2613
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5574 := Call(__e, PrimFunc(symappend), V2614, tmp5573)


if PrimFunc(symshen_4read_1loop) == __self {
__self1, __self2, __self3, __self4 = V2612, tmp5572, tmp5574, V2615
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4read_1loop), V2612, tmp5572, tmp5574, V2615)
return


} else {
tmp5575 := Call(__e, PrimFunc(symshen_4record_1it), V2614)


_ = tmp5575

__e.TailApply(PrimFunc(symshen_4process_1sexprs), let__5045)
return


}


} else {
tmp5578 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2612)


tmp5579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2613, Nil)
}
__typedArg0 := V2613
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5580 := Call(__e, PrimFunc(symappend), V2614, tmp5579)


if PrimFunc(symshen_4read_1loop) == __self {
__self1, __self2, __self3, __self4 = V2612, tmp5578, tmp5580, V2615
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4read_1loop), V2612, tmp5578, tmp5580, V2615)
return


}


}


}


}


}, 4)

tmp5589 := Call(__e, ns2_1set, symshen_4read_1loop, tmp5565)


_ = tmp5589

tmp5590 := MakeNative(func(__e *ControlFlow) {
V2618 := __e.Get(1)
_ = V2618
tmp5591 := MakeNative(func(__e *ControlFlow) {
tmp5592 := MakeNative(func(__e *ControlFlow) {
Z2619 := __e.Get(1)
_ = Z2619
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2619)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp5592, V2618)
return


}, 0)

tmp5593 := MakeNative(func(__e *ControlFlow) {
Z2620 := __e.Get(1)
_ = Z2620
__e.Return(symshen_4i_1failed_b)
return
}, 1)

__e.TailApply(try_1catch, tmp5591, tmp5593)
return


}, 1)

tmp5594 := Call(__e, ns2_1set, symshen_4try_1parse, tmp5590)


_ = tmp5594

tmp5595 := MakeNative(func(__e *ControlFlow) {
V2623 := __e.Get(1)
_ = V2623
tmp5599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4i_1failed_b, V2623)
}
__typedArg0 := symshen_4i_1failed_b
__typedArg1 := V2623
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5599 {
__e.Return(True)
return
} else {
tmp5597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2623)
}
__typedArg0 := Nil
__typedArg1 := V2623
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5597 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp5600 := Call(__e, ns2_1set, symshen_4nothing_1doing_2, tmp5595)


_ = tmp5600

tmp5601 := MakeNative(func(__e *ControlFlow) {
V2624 := __e.Get(1)
_ = V2624
tmp5602 := Call(__e, PrimFunc(symshen_4bytes_1_6string), V2624)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dit_d, tmp5602)
}
__typedArg0 := symshen_4_dit_d
__typedArg1 := tmp5602
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp5603 := Call(__e, ns2_1set, symshen_4record_1it, tmp5601)


_ = tmp5603

tmp5604 := MakeNative(func(__e *ControlFlow) {
V2625 := __e.Get(1)
_ = V2625
tmp5612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2625)
}
__typedArg0 := Nil
__typedArg1 := V2625
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5612 {
__e.Return(MakeString(""))
return
} else {
tmp5610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2625)
}
__typedArg0 := V2625
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5610 {
tmp5605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2625)
}
__typedArg0 := V2625
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(tmp5605)
}
__typedArg0 := tmp5605
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp5607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2625)
}
__typedArg0 := V2625
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5608 := Call(__e, PrimFunc(symshen_4bytes_1_6string), tmp5607)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp5606)
__typedS1, __typedOK1 := TypedString(tmp5608)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp5606
__typedArg1 := tmp5608
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4bytes_1_6string)
return
}


}


}, 1)

tmp5613 := Call(__e, ns2_1set, symshen_4bytes_1_6string, tmp5604)


_ = tmp5613

tmp5614 := MakeNative(func(__e *ControlFlow) {
V2626 := __e.Get(1)
_ = V2626
tmp5615 := Call(__e, PrimFunc(symshen_4unpackage_emacroexpand), V2626)


let__5046 := tmp5615
_ = let__5046

tmp5616 := Call(__e, PrimFunc(symshen_4find_1arities), let__5046)


let__5047 := tmp5616
_ = let__5047

tmp5617 := Call(__e, PrimFunc(symshen_4find_1types), let__5046)


let__5048 := tmp5617
_ = let__5048

tmp5618 := MakeNative(func(__e *ControlFlow) {
Z2630 := __e.Get(1)
_ = Z2630
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2630, let__5048)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp5618, let__5046)
return


}, 1)

tmp5619 := Call(__e, ns2_1set, symshen_4process_1sexprs, tmp5614)


_ = tmp5619

tmp5620 := MakeNative(func(__e *ControlFlow) {
V2631 := __e.Get(1)
_ = V2631
tmp5642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2631)
}
__typedArg0 := V2631
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5633 Obj

if True == tmp5642 {
tmp5640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2631)
}
__typedArg0 := V2631
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5640)
}
__typedArg0 := tmp5640
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5635 Obj

if True == tmp5641 {
tmp5637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2631)
}
__typedArg0 := V2631
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp5639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp5637, tmp5638)
}
__typedArg0 := tmp5637
__typedArg1 := tmp5638
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5636 Obj

if True == tmp5639 {
ifres5636 = True


} else {
ifres5636 = False


}

ifres5635 = ifres5636


} else {
ifres5635 = False


}

var ifres5634 Obj

if True == ifres5635 {
ifres5634 = True


} else {
ifres5634 = False


}

ifres5633 = ifres5634


} else {
ifres5633 = False


}

if True == ifres5633 {
tmp5621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2631)
}
__typedArg0 := V2631
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5621)
}
__typedArg0 := tmp5621
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2631)
}
__typedArg0 := V2631
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5623)
}
__typedArg0 := tmp5623
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5625 := Call(__e, PrimFunc(symshen_4find_1types), tmp5624)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5622, tmp5625)
}
__typedArg0 := tmp5622
__typedArg1 := tmp5625
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp5631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2631)
}
__typedArg0 := V2631
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5631 {
tmp5626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2631)
}
__typedArg0 := V2631
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5627 := Call(__e, PrimFunc(symshen_4find_1types), tmp5626)


tmp5628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2631)
}
__typedArg0 := V2631
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5629 := Call(__e, PrimFunc(symshen_4find_1types), tmp5628)


__e.TailApply(PrimFunc(symappend), tmp5627, tmp5629)
return


} else {
__e.Return(Nil)
return
}


}


}, 1)

tmp5643 := Call(__e, ns2_1set, symshen_4find_1types, tmp5620)


_ = tmp5643

tmp5644 := MakeNative(func(__e *ControlFlow) {
V2634 := __e.Get(1)
_ = V2634
tmp5693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2634)
}
__typedArg0 := V2634
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5674 Obj

if True == tmp5693 {
tmp5691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2634)
}
__typedArg0 := V2634
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp5691)
}
__typedArg0 := symdefine
__typedArg1 := tmp5691
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5676 Obj

if True == tmp5692 {
tmp5689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2634)
}
__typedArg0 := V2634
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5689)
}
__typedArg0 := tmp5689
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5678 Obj

if True == tmp5690 {
tmp5686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2634)
}
__typedArg0 := V2634
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5686)
}
__typedArg0 := tmp5686
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5687)
}
__typedArg0 := tmp5687
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5680 Obj

if True == tmp5688 {
tmp5682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2634)
}
__typedArg0 := V2634
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5682)
}
__typedArg0 := tmp5682
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5683)
}
__typedArg0 := tmp5683
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, tmp5684)
}
__typedArg0 := sym_i
__typedArg1 := tmp5684
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5681 Obj

if True == tmp5685 {
ifres5681 = True


} else {
ifres5681 = False


}

ifres5680 = ifres5681


} else {
ifres5680 = False


}

var ifres5679 Obj

if True == ifres5680 {
ifres5679 = True


} else {
ifres5679 = False


}

ifres5678 = ifres5679


} else {
ifres5678 = False


}

var ifres5677 Obj

if True == ifres5678 {
ifres5677 = True


} else {
ifres5677 = False


}

ifres5676 = ifres5677


} else {
ifres5676 = False


}

var ifres5675 Obj

if True == ifres5676 {
ifres5675 = True


} else {
ifres5675 = False


}

ifres5674 = ifres5675


} else {
ifres5674 = False


}

if True == ifres5674 {
tmp5645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2634)
}
__typedArg0 := V2634
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5645)
}
__typedArg0 := tmp5645
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2634)
}
__typedArg0 := V2634
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5647)
}
__typedArg0 := tmp5647
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2634)
}
__typedArg0 := V2634
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5649)
}
__typedArg0 := tmp5649
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5650)
}
__typedArg0 := tmp5650
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5652 := Call(__e, PrimFunc(symshen_4find_1arity), tmp5648, MakeInteger(1), tmp5651)


__e.TailApply(PrimFunc(symshen_4store_1arity), tmp5646, tmp5652)
return


} else {
tmp5672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2634)
}
__typedArg0 := V2634
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5664 Obj

if True == tmp5672 {
tmp5670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2634)
}
__typedArg0 := V2634
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp5670)
}
__typedArg0 := symdefine
__typedArg1 := tmp5670
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5666 Obj

if True == tmp5671 {
tmp5668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2634)
}
__typedArg0 := V2634
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5668)
}
__typedArg0 := tmp5668
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5667 Obj

if True == tmp5669 {
ifres5667 = True


} else {
ifres5667 = False


}

ifres5666 = ifres5667


} else {
ifres5666 = False


}

var ifres5665 Obj

if True == ifres5666 {
ifres5665 = True


} else {
ifres5665 = False


}

ifres5664 = ifres5665


} else {
ifres5664 = False


}

if True == ifres5664 {
tmp5653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2634)
}
__typedArg0 := V2634
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5653)
}
__typedArg0 := tmp5653
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2634)
}
__typedArg0 := V2634
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5655)
}
__typedArg0 := tmp5655
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2634)
}
__typedArg0 := V2634
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5657)
}
__typedArg0 := tmp5657
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5659 := Call(__e, PrimFunc(symshen_4find_1arity), tmp5656, MakeInteger(0), tmp5658)


__e.TailApply(PrimFunc(symshen_4store_1arity), tmp5654, tmp5659)
return


} else {
tmp5662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2634)
}
__typedArg0 := V2634
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5662 {
tmp5660 := MakeNative(func(__e *ControlFlow) {
Z2635 := __e.Get(1)
_ = Z2635
__e.TailApply(PrimFunc(symshen_4find_1arities), Z2635)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp5660, V2634)
return


} else {
__e.Return(symshen_4skip)
return
}


}


}


}, 1)

tmp5694 := Call(__e, ns2_1set, symshen_4find_1arities, tmp5644)


_ = tmp5694

tmp5695 := MakeNative(func(__e *ControlFlow) {
V2636 := __e.Get(1)
_ = V2636
V2637 := __e.Get(2)
_ = V2637
tmp5696 := Call(__e, PrimFunc(symarity), V2636)


let__5049 := tmp5696
_ = let__5049

tmp5707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__5049, MakeInteger(-1))
}
__typedArg0 := let__5049
__typedArg1 := MakeInteger(-1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5707 {
__e.TailApply(PrimFunc(symshen_4execute_1store_1arity), V2636, V2637)
return
} else {
tmp5705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__5049, V2637)
}
__typedArg0 := let__5049
__typedArg1 := V2637
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5705 {
__e.Return(symshen_4skip)
return
} else {
tmp5703 := Call(__e, PrimFunc(symshen_4sysfunc_2), V2636)


if True == tmp5703 {
tmp5697 := Call(__e, PrimFunc(symshen_4app), V2636, MakeString(" is a system function\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp5697)
}
__typedArg0 := tmp5697
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
tmp5698 := Call(__e, PrimFunc(symshen_4app), V2636, MakeString(" may cause errors\n"), symshen_4a)


tmp5699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("changing the arity of "))
__typedS1, __typedOK1 := TypedString(tmp5698)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("changing the arity of ")
__typedArg1 := tmp5698
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp5700 := Call(__e, PrimFunc(symstoutput))


tmp5701 := Call(__e, PrimFunc(sympr), tmp5699, tmp5700)


_ = tmp5701

__e.TailApply(PrimFunc(symshen_4execute_1store_1arity), V2636, V2637)
return


}


}


}


}, 2)

tmp5708 := Call(__e, ns2_1set, symshen_4store_1arity, tmp5695)


_ = tmp5708

tmp5709 := MakeNative(func(__e *ControlFlow) {
V2639 := __e.Get(1)
_ = V2639
V2640 := __e.Get(2)
_ = V2640
tmp5714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V2640)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V2640
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5714 {
tmp5710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V2639, symarity, MakeInteger(0), tmp5710)
return


} else {
tmp5711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp5712 := Call(__e, PrimFunc(symput), V2639, symarity, V2640, tmp5711)


_ = tmp5712

__e.TailApply(PrimFunc(symshen_4update_1lambdatable), V2639, V2640)
return


}


}, 2)

tmp5715 := Call(__e, ns2_1set, symshen_4execute_1store_1arity, tmp5709)


_ = tmp5715

tmp5716 := MakeNative(func(__e *ControlFlow) {
V2641 := __e.Get(1)
_ = V2641
V2642 := __e.Get(2)
_ = V2642
tmp5717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2641, Nil)
}
__typedArg0 := V2641
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5718 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp5717, V2642)


tmp5719 := Call(__e, PrimFunc(symeval_1kl), tmp5718)


let__5050 := tmp5719
_ = let__5050

tmp5720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2641, let__5050)
}
__typedArg0 := V2641
__typedArg1 := let__5050
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5721 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp5720)


let__5051 := tmp5721
_ = let__5051

__e.Return(V2642)
return


}, 2)

tmp5722 := Call(__e, ns2_1set, symshen_4update_1lambdatable, tmp5716)


_ = tmp5722

tmp5723 := MakeNative(func(__e *ControlFlow) {
V2647 := __e.Get(1)
_ = V2647
V2648 := __e.Get(2)
_ = V2648
tmp5739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V2648)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V2648
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5739 {
__e.Return(symshen_4skip)
return
} else {
tmp5737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(1), V2648)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := V2648
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5737 {
tmp5724 := Call(__e, PrimFunc(symgensym), symY)


let__5052 := tmp5724
_ = let__5052

tmp5725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__5052, Nil)
}
__typedArg0 := let__5052
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5726 := Call(__e, PrimFunc(symappend), V2647, tmp5725)


tmp5727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5726, Nil)
}
__typedArg0 := tmp5726
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__5052, tmp5727)
}
__typedArg0 := let__5052
__typedArg1 := tmp5727
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp5728)
}
__typedArg0 := symlambda
__typedArg1 := tmp5728
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp5729 := Call(__e, PrimFunc(symgensym), symY)


let__5053 := tmp5729
_ = let__5053

tmp5730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__5053, Nil)
}
__typedArg0 := let__5053
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5731 := Call(__e, PrimFunc(symappend), V2647, tmp5730)


tmp5733 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp5731, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V2648)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V2648
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


tmp5734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5733, Nil)
}
__typedArg0 := tmp5733
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__5053, tmp5734)
}
__typedArg0 := let__5053
__typedArg1 := tmp5734
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp5735)
}
__typedArg0 := symlambda
__typedArg1 := tmp5735
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}


}, 2)

tmp5740 := Call(__e, ns2_1set, symshen_4lambda_1function, tmp5723)


_ = tmp5740

tmp5741 := MakeNative(func(__e *ControlFlow) {
V2660 := __e.Get(1)
_ = V2660
V2661 := __e.Get(2)
_ = V2661
V2662 := __e.Get(3)
_ = V2662
tmp5764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2662)
}
__typedArg0 := Nil
__typedArg1 := V2662
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5764 {
tmp5742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2660, V2661)
}
__typedArg0 := V2660
__typedArg1 := V2661
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5742, Nil)
}
__typedArg0 := tmp5742
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp5762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2662)
}
__typedArg0 := V2662
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5753 Obj

if True == tmp5762 {
tmp5760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2662)
}
__typedArg0 := V2662
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5760)
}
__typedArg0 := tmp5760
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5755 Obj

if True == tmp5761 {
tmp5757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2662)
}
__typedArg0 := V2662
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5757)
}
__typedArg0 := tmp5757
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V2660, tmp5758)
}
__typedArg0 := V2660
__typedArg1 := tmp5758
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5756 Obj

if True == tmp5759 {
ifres5756 = True


} else {
ifres5756 = False


}

ifres5755 = ifres5756


} else {
ifres5755 = False


}

var ifres5754 Obj

if True == ifres5755 {
ifres5754 = True


} else {
ifres5754 = False


}

ifres5753 = ifres5754


} else {
ifres5753 = False


}

if True == ifres5753 {
tmp5743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2662)
}
__typedArg0 := V2662
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5743)
}
__typedArg0 := tmp5743
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5744, V2661)
}
__typedArg0 := tmp5744
__typedArg1 := V2661
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2662)
}
__typedArg0 := V2662
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5745, tmp5746)
}
__typedArg0 := tmp5745
__typedArg1 := tmp5746
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp5751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2662)
}
__typedArg0 := V2662
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5751 {
tmp5747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2662)
}
__typedArg0 := V2662
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2662)
}
__typedArg0 := V2662
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5749 := Call(__e, PrimFunc(symshen_4assoc_1_6), V2660, V2661, tmp5748)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5747, tmp5749)
}
__typedArg0 := tmp5747
__typedArg1 := tmp5749
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.assoc->"))
}
__typedArg0 := MakeString("implementation error in shen.assoc->")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 3)

tmp5765 := Call(__e, ns2_1set, symshen_4assoc_1_6, tmp5741)


_ = tmp5765

tmp5766 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__selftop:
V2677 := __self1
_ = V2677
V2678 := __self2
_ = V2678
V2679 := __self3
_ = V2679
tmp5813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V2678)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V2678
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5806 Obj

if True == tmp5813 {
tmp5812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2679)
}
__typedArg0 := V2679
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5808 Obj

if True == tmp5812 {
tmp5810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2679)
}
__typedArg0 := V2679
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp5810, sym_1_6)
}
__typedArg0 := tmp5810
__typedArg1 := sym_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5809 Obj

if True == tmp5811 {
ifres5809 = True


} else {
ifres5809 = False


}

ifres5808 = ifres5809


} else {
ifres5808 = False


}

var ifres5807 Obj

if True == ifres5808 {
ifres5807 = True


} else {
ifres5807 = False


}

ifres5806 = ifres5807


} else {
ifres5806 = False


}

if True == ifres5806 {
__e.Return(MakeInteger(0))
return
} else {
tmp5804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V2678)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V2678
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5797 Obj

if True == tmp5804 {
tmp5803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2679)
}
__typedArg0 := V2679
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5799 Obj

if True == tmp5803 {
tmp5801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2679)
}
__typedArg0 := V2679
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp5801, sym_5_1)
}
__typedArg0 := tmp5801
__typedArg1 := sym_5_1
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5800 Obj

if True == tmp5802 {
ifres5800 = True


} else {
ifres5800 = False


}

ifres5799 = ifres5800


} else {
ifres5799 = False


}

var ifres5798 Obj

if True == ifres5799 {
ifres5798 = True


} else {
ifres5798 = False


}

ifres5797 = ifres5798


} else {
ifres5797 = False


}

if True == ifres5797 {
__e.Return(MakeInteger(0))
return
} else {
tmp5795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V2678)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V2678
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5792 Obj

if True == tmp5795 {
tmp5794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2679)
}
__typedArg0 := V2679
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5793 Obj

if True == tmp5794 {
ifres5793 = True


} else {
ifres5793 = False


}

ifres5792 = ifres5793


} else {
ifres5792 = False


}

if True == ifres5792 {
tmp5767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2679)
}
__typedArg0 := V2679
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5768 := Call(__e, PrimFunc(symshen_4find_1arity), V2677, MakeInteger(0), tmp5767)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp5768)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp5768
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


} else {
tmp5790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(1), V2678)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := V2678
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5783 Obj

if True == tmp5790 {
tmp5789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2679)
}
__typedArg0 := V2679
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5785 Obj

if True == tmp5789 {
tmp5787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2679)
}
__typedArg0 := V2679
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, tmp5787)
}
__typedArg0 := sym_j
__typedArg1 := tmp5787
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5786 Obj

if True == tmp5788 {
ifres5786 = True


} else {
ifres5786 = False


}

ifres5785 = ifres5786


} else {
ifres5785 = False


}

var ifres5784 Obj

if True == ifres5785 {
ifres5784 = True


} else {
ifres5784 = False


}

ifres5783 = ifres5784


} else {
ifres5783 = False


}

if True == ifres5783 {
tmp5769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2679)
}
__typedArg0 := V2679
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4find_1arity) == __self {
__self1, __self2, __self3 = V2677, MakeInteger(0), tmp5769
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4find_1arity), V2677, MakeInteger(0), tmp5769)
return


} else {
tmp5781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(1), V2678)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := V2678
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5778 Obj

if True == tmp5781 {
tmp5780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2679)
}
__typedArg0 := V2679
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5779 Obj

if True == tmp5780 {
ifres5779 = True


} else {
ifres5779 = False


}

ifres5778 = ifres5779


} else {
ifres5778 = False


}

if True == ifres5778 {
tmp5770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2679)
}
__typedArg0 := V2679
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4find_1arity) == __self {
__self1, __self2, __self3 = V2677, MakeInteger(1), tmp5770
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4find_1arity), V2677, MakeInteger(1), tmp5770)
return


} else {
tmp5776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(1), V2678)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := V2678
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5776 {
tmp5771 := Call(__e, PrimFunc(symshen_4app), V2677, MakeString(" definition: missing }\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("syntax error in "))
__typedS1, __typedOK1 := TypedString(tmp5771)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("syntax error in ")
__typedArg1 := tmp5771
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("syntax error in "))
__typedS1, __typedOK1 := TypedString(tmp5771)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("syntax error in ")
__typedArg1 := tmp5771
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
tmp5773 := Call(__e, PrimFunc(symshen_4app), V2677, MakeString(" definition: missing -> or <-\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("syntax error in "))
__typedS1, __typedOK1 := TypedString(tmp5773)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("syntax error in ")
__typedArg1 := tmp5773
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("syntax error in "))
__typedS1, __typedOK1 := TypedString(tmp5773)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("syntax error in ")
__typedArg1 := tmp5773
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}


}


}


}


}


}, 3)

tmp5814 := Call(__e, ns2_1set, symshen_4find_1arity, tmp5766)


_ = tmp5814

tmp5815 := MakeNative(func(__e *ControlFlow) {
V2680 := __e.Get(1)
_ = V2680
tmp5816 := Call(__e, PrimFunc(symshen_4_5lsb_6), V2680)


let__5055 := tmp5816
_ = let__5055

tmp5840 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5055)


var ifres5817 Obj

if True == tmp5840 {
tmp5818 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5817 = tmp5818


} else {
tmp5819 := Call(__e, PrimFunc(symshen_4in_1_6), let__5055)


let__5056 := tmp5819
_ = let__5056

tmp5820 := Call(__e, PrimFunc(symshen_4_5s_1exprs1_6), let__5056)


let__5057 := tmp5820
_ = let__5057

tmp5839 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5057)


var ifres5821 Obj

if True == tmp5839 {
tmp5822 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5821 = tmp5822


} else {
tmp5823 := Call(__e, PrimFunc(symshen_4_5_1out), let__5057)


let__5058 := tmp5823
_ = let__5058

tmp5824 := Call(__e, PrimFunc(symshen_4in_1_6), let__5057)


let__5059 := tmp5824
_ = let__5059

tmp5825 := Call(__e, PrimFunc(symshen_4_5rsb_6), let__5059)


let__5060 := tmp5825
_ = let__5060

tmp5838 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5060)


var ifres5826 Obj

if True == tmp5838 {
tmp5827 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5826 = tmp5827


} else {
tmp5828 := Call(__e, PrimFunc(symshen_4in_1_6), let__5060)


let__5061 := tmp5828
_ = let__5061

tmp5829 := Call(__e, PrimFunc(symshen_4_5s_1exprs2_6), let__5061)


let__5062 := tmp5829
_ = let__5062

tmp5837 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5062)


var ifres5830 Obj

if True == tmp5837 {
tmp5831 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5830 = tmp5831


} else {
tmp5832 := Call(__e, PrimFunc(symshen_4_5_1out), let__5062)


let__5063 := tmp5832
_ = let__5063

tmp5833 := Call(__e, PrimFunc(symshen_4in_1_6), let__5062)


let__5064 := tmp5833
_ = let__5064

tmp5834 := Call(__e, PrimFunc(symshen_4cons_1form), let__5058)


tmp5835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5834, let__5063)
}
__typedArg0 := tmp5834
__typedArg1 := let__5063
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5836 := Call(__e, PrimFunc(symshen_4comb), let__5064, tmp5835)


ifres5830 = tmp5836


}

ifres5826 = ifres5830


}

ifres5821 = ifres5826


}

ifres5817 = ifres5821


}

let__5054 := ifres5817
_ = let__5054

tmp6034 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5054)


if True == tmp6034 {
tmp5841 := Call(__e, PrimFunc(symshen_4_5lrb_6), V2680)


let__5066 := tmp5841
_ = let__5066

tmp5864 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5066)


var ifres5842 Obj

if True == tmp5864 {
tmp5843 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5842 = tmp5843


} else {
tmp5844 := Call(__e, PrimFunc(symshen_4in_1_6), let__5066)


let__5067 := tmp5844
_ = let__5067

tmp5845 := Call(__e, PrimFunc(symshen_4_5s_1exprs1_6), let__5067)


let__5068 := tmp5845
_ = let__5068

tmp5863 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5068)


var ifres5846 Obj

if True == tmp5863 {
tmp5847 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5846 = tmp5847


} else {
tmp5848 := Call(__e, PrimFunc(symshen_4_5_1out), let__5068)


let__5069 := tmp5848
_ = let__5069

tmp5849 := Call(__e, PrimFunc(symshen_4in_1_6), let__5068)


let__5070 := tmp5849
_ = let__5070

tmp5850 := Call(__e, PrimFunc(symshen_4_5rrb_6), let__5070)


let__5071 := tmp5850
_ = let__5071

tmp5862 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5071)


var ifres5851 Obj

if True == tmp5862 {
tmp5852 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5851 = tmp5852


} else {
tmp5853 := Call(__e, PrimFunc(symshen_4in_1_6), let__5071)


let__5072 := tmp5853
_ = let__5072

tmp5854 := Call(__e, PrimFunc(symshen_4_5s_1exprs2_6), let__5072)


let__5073 := tmp5854
_ = let__5073

tmp5861 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5073)


var ifres5855 Obj

if True == tmp5861 {
tmp5856 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5855 = tmp5856


} else {
tmp5857 := Call(__e, PrimFunc(symshen_4_5_1out), let__5073)


let__5074 := tmp5857
_ = let__5074

tmp5858 := Call(__e, PrimFunc(symshen_4in_1_6), let__5073)


let__5075 := tmp5858
_ = let__5075

tmp5859 := Call(__e, PrimFunc(symshen_4add_1sexpr), let__5069, let__5074)


tmp5860 := Call(__e, PrimFunc(symshen_4comb), let__5075, tmp5859)


ifres5855 = tmp5860


}

ifres5851 = ifres5855


}

ifres5846 = ifres5851


}

ifres5842 = ifres5846


}

let__5065 := ifres5842
_ = let__5065

tmp6032 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5065)


if True == tmp6032 {
tmp5865 := Call(__e, PrimFunc(symshen_4_5lcurly_6), V2680)


let__5077 := tmp5865
_ = let__5077

tmp5877 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5077)


var ifres5866 Obj

if True == tmp5877 {
tmp5867 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5866 = tmp5867


} else {
tmp5868 := Call(__e, PrimFunc(symshen_4in_1_6), let__5077)


let__5078 := tmp5868
_ = let__5078

tmp5869 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__5078)


let__5079 := tmp5869
_ = let__5079

tmp5876 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5079)


var ifres5870 Obj

if True == tmp5876 {
tmp5871 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5870 = tmp5871


} else {
tmp5872 := Call(__e, PrimFunc(symshen_4_5_1out), let__5079)


let__5080 := tmp5872
_ = let__5080

tmp5873 := Call(__e, PrimFunc(symshen_4in_1_6), let__5079)


let__5081 := tmp5873
_ = let__5081

tmp5874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, let__5080)
}
__typedArg0 := sym_i
__typedArg1 := let__5080
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5875 := Call(__e, PrimFunc(symshen_4comb), let__5081, tmp5874)


ifres5870 = tmp5875


}

ifres5866 = ifres5870


}

let__5076 := ifres5866
_ = let__5076

tmp6030 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5076)


if True == tmp6030 {
tmp5878 := Call(__e, PrimFunc(symshen_4_5rcurly_6), V2680)


let__5083 := tmp5878
_ = let__5083

tmp5890 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5083)


var ifres5879 Obj

if True == tmp5890 {
tmp5880 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5879 = tmp5880


} else {
tmp5881 := Call(__e, PrimFunc(symshen_4in_1_6), let__5083)


let__5084 := tmp5881
_ = let__5084

tmp5882 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__5084)


let__5085 := tmp5882
_ = let__5085

tmp5889 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5085)


var ifres5883 Obj

if True == tmp5889 {
tmp5884 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5883 = tmp5884


} else {
tmp5885 := Call(__e, PrimFunc(symshen_4_5_1out), let__5085)


let__5086 := tmp5885
_ = let__5086

tmp5886 := Call(__e, PrimFunc(symshen_4in_1_6), let__5085)


let__5087 := tmp5886
_ = let__5087

tmp5887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, let__5086)
}
__typedArg0 := sym_j
__typedArg1 := let__5086
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5888 := Call(__e, PrimFunc(symshen_4comb), let__5087, tmp5887)


ifres5883 = tmp5888


}

ifres5879 = ifres5883


}

let__5082 := ifres5879
_ = let__5082

tmp6028 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5082)


if True == tmp6028 {
tmp5891 := Call(__e, PrimFunc(symshen_4_5bar_6), V2680)


let__5089 := tmp5891
_ = let__5089

tmp5903 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5089)


var ifres5892 Obj

if True == tmp5903 {
tmp5893 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5892 = tmp5893


} else {
tmp5894 := Call(__e, PrimFunc(symshen_4in_1_6), let__5089)


let__5090 := tmp5894
_ = let__5090

tmp5895 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__5090)


let__5091 := tmp5895
_ = let__5091

tmp5902 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5091)


var ifres5896 Obj

if True == tmp5902 {
tmp5897 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5896 = tmp5897


} else {
tmp5898 := Call(__e, PrimFunc(symshen_4_5_1out), let__5091)


let__5092 := tmp5898
_ = let__5092

tmp5899 := Call(__e, PrimFunc(symshen_4in_1_6), let__5091)


let__5093 := tmp5899
_ = let__5093

tmp5900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbar_b, let__5092)
}
__typedArg0 := symbar_b
__typedArg1 := let__5092
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5901 := Call(__e, PrimFunc(symshen_4comb), let__5093, tmp5900)


ifres5896 = tmp5901


}

ifres5892 = ifres5896


}

let__5088 := ifres5892
_ = let__5088

tmp6026 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5088)


if True == tmp6026 {
tmp5904 := Call(__e, PrimFunc(symshen_4_5semicolon_6), V2680)


let__5095 := tmp5904
_ = let__5095

tmp5917 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5095)


var ifres5905 Obj

if True == tmp5917 {
tmp5906 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5905 = tmp5906


} else {
tmp5907 := Call(__e, PrimFunc(symshen_4in_1_6), let__5095)


let__5096 := tmp5907
_ = let__5096

tmp5908 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__5096)


let__5097 := tmp5908
_ = let__5097

tmp5916 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5097)


var ifres5909 Obj

if True == tmp5916 {
tmp5910 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5909 = tmp5910


} else {
tmp5911 := Call(__e, PrimFunc(symshen_4_5_1out), let__5097)


let__5098 := tmp5911
_ = let__5098

tmp5912 := Call(__e, PrimFunc(symshen_4in_1_6), let__5097)


let__5099 := tmp5912
_ = let__5099

tmp5913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp5914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5913, let__5098)
}
__typedArg0 := tmp5913
__typedArg1 := let__5098
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5915 := Call(__e, PrimFunc(symshen_4comb), let__5099, tmp5914)


ifres5909 = tmp5915


}

ifres5905 = ifres5909


}

let__5094 := ifres5905
_ = let__5094

tmp6024 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5094)


if True == tmp6024 {
tmp5918 := Call(__e, PrimFunc(symshen_4_5colon_6), V2680)


let__5101 := tmp5918
_ = let__5101

tmp5936 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5101)


var ifres5919 Obj

if True == tmp5936 {
tmp5920 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5919 = tmp5920


} else {
tmp5921 := Call(__e, PrimFunc(symshen_4in_1_6), let__5101)


let__5102 := tmp5921
_ = let__5102

tmp5922 := Call(__e, PrimFunc(symshen_4_5equal_6), let__5102)


let__5103 := tmp5922
_ = let__5103

tmp5935 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5103)


var ifres5923 Obj

if True == tmp5935 {
tmp5924 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5923 = tmp5924


} else {
tmp5925 := Call(__e, PrimFunc(symshen_4in_1_6), let__5103)


let__5104 := tmp5925
_ = let__5104

tmp5926 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__5104)


let__5105 := tmp5926
_ = let__5105

tmp5934 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5105)


var ifres5927 Obj

if True == tmp5934 {
tmp5928 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5927 = tmp5928


} else {
tmp5929 := Call(__e, PrimFunc(symshen_4_5_1out), let__5105)


let__5106 := tmp5929
_ = let__5106

tmp5930 := Call(__e, PrimFunc(symshen_4in_1_6), let__5105)


let__5107 := tmp5930
_ = let__5107

tmp5931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":="))
}
__typedArg0 := MakeString(":=")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp5932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5931, let__5106)
}
__typedArg0 := tmp5931
__typedArg1 := let__5106
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5933 := Call(__e, PrimFunc(symshen_4comb), let__5107, tmp5932)


ifres5927 = tmp5933


}

ifres5923 = ifres5927


}

ifres5919 = ifres5923


}

let__5100 := ifres5919
_ = let__5100

tmp6022 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5100)


if True == tmp6022 {
tmp5937 := Call(__e, PrimFunc(symshen_4_5colon_6), V2680)


let__5109 := tmp5937
_ = let__5109

tmp5950 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5109)


var ifres5938 Obj

if True == tmp5950 {
tmp5939 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5938 = tmp5939


} else {
tmp5940 := Call(__e, PrimFunc(symshen_4in_1_6), let__5109)


let__5110 := tmp5940
_ = let__5110

tmp5941 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__5110)


let__5111 := tmp5941
_ = let__5111

tmp5949 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5111)


var ifres5942 Obj

if True == tmp5949 {
tmp5943 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5942 = tmp5943


} else {
tmp5944 := Call(__e, PrimFunc(symshen_4_5_1out), let__5111)


let__5112 := tmp5944
_ = let__5112

tmp5945 := Call(__e, PrimFunc(symshen_4in_1_6), let__5111)


let__5113 := tmp5945
_ = let__5113

tmp5946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp5947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5946, let__5112)
}
__typedArg0 := tmp5946
__typedArg1 := let__5112
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5948 := Call(__e, PrimFunc(symshen_4comb), let__5113, tmp5947)


ifres5942 = tmp5948


}

ifres5938 = ifres5942


}

let__5108 := ifres5938
_ = let__5108

tmp6020 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5108)


if True == tmp6020 {
tmp5951 := Call(__e, PrimFunc(symshen_4_5comma_6), V2680)


let__5115 := tmp5951
_ = let__5115

tmp5964 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5115)


var ifres5952 Obj

if True == tmp5964 {
tmp5953 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5952 = tmp5953


} else {
tmp5954 := Call(__e, PrimFunc(symshen_4in_1_6), let__5115)


let__5116 := tmp5954
_ = let__5116

tmp5955 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__5116)


let__5117 := tmp5955
_ = let__5117

tmp5963 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5117)


var ifres5956 Obj

if True == tmp5963 {
tmp5957 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5956 = tmp5957


} else {
tmp5958 := Call(__e, PrimFunc(symshen_4_5_1out), let__5117)


let__5118 := tmp5958
_ = let__5118

tmp5959 := Call(__e, PrimFunc(symshen_4in_1_6), let__5117)


let__5119 := tmp5959
_ = let__5119

tmp5960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(","))
}
__typedArg0 := MakeString(",")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp5961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5960, let__5118)
}
__typedArg0 := tmp5960
__typedArg1 := let__5118
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5962 := Call(__e, PrimFunc(symshen_4comb), let__5119, tmp5961)


ifres5956 = tmp5962


}

ifres5952 = ifres5956


}

let__5114 := ifres5952
_ = let__5114

tmp6018 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5114)


if True == tmp6018 {
tmp5965 := Call(__e, PrimFunc(symshen_4_5comment_6), V2680)


let__5121 := tmp5965
_ = let__5121

tmp5976 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5121)


var ifres5966 Obj

if True == tmp5976 {
tmp5967 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5966 = tmp5967


} else {
tmp5968 := Call(__e, PrimFunc(symshen_4in_1_6), let__5121)


let__5122 := tmp5968
_ = let__5122

tmp5969 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__5122)


let__5123 := tmp5969
_ = let__5123

tmp5975 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5123)


var ifres5970 Obj

if True == tmp5975 {
tmp5971 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5970 = tmp5971


} else {
tmp5972 := Call(__e, PrimFunc(symshen_4_5_1out), let__5123)


let__5124 := tmp5972
_ = let__5124

tmp5973 := Call(__e, PrimFunc(symshen_4in_1_6), let__5123)


let__5125 := tmp5973
_ = let__5125

tmp5974 := Call(__e, PrimFunc(symshen_4comb), let__5125, let__5124)


ifres5970 = tmp5974


}

ifres5966 = ifres5970


}

let__5120 := ifres5966
_ = let__5120

tmp6016 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5120)


if True == tmp6016 {
tmp5977 := Call(__e, PrimFunc(symshen_4_5atom_6), V2680)


let__5127 := tmp5977
_ = let__5127

tmp5990 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5127)


var ifres5978 Obj

if True == tmp5990 {
tmp5979 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5978 = tmp5979


} else {
tmp5980 := Call(__e, PrimFunc(symshen_4_5_1out), let__5127)


let__5128 := tmp5980
_ = let__5128

tmp5981 := Call(__e, PrimFunc(symshen_4in_1_6), let__5127)


let__5129 := tmp5981
_ = let__5129

tmp5982 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__5129)


let__5130 := tmp5982
_ = let__5130

tmp5989 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5130)


var ifres5983 Obj

if True == tmp5989 {
tmp5984 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5983 = tmp5984


} else {
tmp5985 := Call(__e, PrimFunc(symshen_4_5_1out), let__5130)


let__5131 := tmp5985
_ = let__5131

tmp5986 := Call(__e, PrimFunc(symshen_4in_1_6), let__5130)


let__5132 := tmp5986
_ = let__5132

tmp5987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__5128, let__5131)
}
__typedArg0 := let__5128
__typedArg1 := let__5131
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5988 := Call(__e, PrimFunc(symshen_4comb), let__5132, tmp5987)


ifres5983 = tmp5988


}

ifres5978 = ifres5983


}

let__5126 := ifres5978
_ = let__5126

tmp6014 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5126)


if True == tmp6014 {
tmp5991 := Call(__e, PrimFunc(symshen_4_5whitespaces_6), V2680)


let__5134 := tmp5991
_ = let__5134

tmp6002 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5134)


var ifres5992 Obj

if True == tmp6002 {
tmp5993 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5992 = tmp5993


} else {
tmp5994 := Call(__e, PrimFunc(symshen_4in_1_6), let__5134)


let__5135 := tmp5994
_ = let__5135

tmp5995 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__5135)


let__5136 := tmp5995
_ = let__5136

tmp6001 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5136)


var ifres5996 Obj

if True == tmp6001 {
tmp5997 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5996 = tmp5997


} else {
tmp5998 := Call(__e, PrimFunc(symshen_4_5_1out), let__5136)


let__5137 := tmp5998
_ = let__5137

tmp5999 := Call(__e, PrimFunc(symshen_4in_1_6), let__5136)


let__5138 := tmp5999
_ = let__5138

tmp6000 := Call(__e, PrimFunc(symshen_4comb), let__5138, let__5137)


ifres5996 = tmp6000


}

ifres5992 = ifres5996


}

let__5133 := ifres5992
_ = let__5133

tmp6012 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5133)


if True == tmp6012 {
tmp6003 := Call(__e, PrimFunc(sym_5e_6), V2680)


let__5140 := tmp6003
_ = let__5140

tmp6008 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5140)


var ifres6004 Obj

if True == tmp6008 {
tmp6005 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6004 = tmp6005


} else {
tmp6006 := Call(__e, PrimFunc(symshen_4in_1_6), let__5140)


let__5141 := tmp6006
_ = let__5141

tmp6007 := Call(__e, PrimFunc(symshen_4comb), let__5141, Nil)


ifres6004 = tmp6007


}

let__5139 := ifres6004
_ = let__5139

tmp6010 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5139)


if True == tmp6010 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5139)
return
}


} else {
__e.Return(let__5133)
return
}


} else {
__e.Return(let__5126)
return
}


} else {
__e.Return(let__5120)
return
}


} else {
__e.Return(let__5114)
return
}


} else {
__e.Return(let__5108)
return
}


} else {
__e.Return(let__5100)
return
}


} else {
__e.Return(let__5094)
return
}


} else {
__e.Return(let__5088)
return
}


} else {
__e.Return(let__5082)
return
}


} else {
__e.Return(let__5076)
return
}


} else {
__e.Return(let__5065)
return
}


} else {
__e.Return(let__5054)
return
}


}, 1)

tmp6035 := Call(__e, ns2_1set, symshen_4_5s_1exprs_6, tmp5815)


_ = tmp6035

tmp6036 := MakeNative(func(__e *ControlFlow) {
V2769 := __e.Get(1)
_ = V2769
V2770 := __e.Get(2)
_ = V2770
tmp6054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2769)
}
__typedArg0 := V2769
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6041 Obj

if True == tmp6054 {
tmp6052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2769)
}
__typedArg0 := V2769
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_3, tmp6052)
}
__typedArg0 := sym_3
__typedArg1 := tmp6052
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres6043 Obj

if True == tmp6053 {
tmp6050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2769)
}
__typedArg0 := V2769
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp6050)
}
__typedArg0 := tmp6050
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6045 Obj

if True == tmp6051 {
tmp6047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2769)
}
__typedArg0 := V2769
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6047)
}
__typedArg0 := tmp6047
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp6048)
}
__typedArg0 := Nil
__typedArg1 := tmp6048
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres6046 Obj

if True == tmp6049 {
ifres6046 = True


} else {
ifres6046 = False


}

ifres6045 = ifres6046


} else {
ifres6045 = False


}

var ifres6044 Obj

if True == ifres6045 {
ifres6044 = True


} else {
ifres6044 = False


}

ifres6043 = ifres6044


} else {
ifres6043 = False


}

var ifres6042 Obj

if True == ifres6043 {
ifres6042 = True


} else {
ifres6042 = False


}

ifres6041 = ifres6042


} else {
ifres6041 = False


}

if True == ifres6041 {
tmp6037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2769)
}
__typedArg0 := V2769
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp6037)
}
__typedArg0 := tmp6037
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6039 := Call(__e, PrimFunc(symexplode), tmp6038)


__e.TailApply(PrimFunc(symappend), tmp6039, V2770)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2769, V2770)
}
__typedArg0 := V2769
__typedArg1 := V2770
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
}


}, 2)

tmp6055 := Call(__e, ns2_1set, symshen_4add_1sexpr, tmp6036)


_ = tmp6055

tmp6056 := MakeNative(func(__e *ControlFlow) {
V2771 := __e.Get(1)
_ = V2771
tmp6061 := Call(__e, PrimFunc(symshen_4hds_a_2), V2771, MakeInteger(91))


var ifres6057 Obj

if True == tmp6061 {
tmp6058 := Call(__e, PrimFunc(symtail), V2771)


let__5143 := tmp6058
_ = let__5143

tmp6059 := Call(__e, PrimFunc(symshen_4comb), let__5143, symshen_4skip)


ifres6057 = tmp6059


} else {
tmp6060 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6057 = tmp6060


}

let__5142 := ifres6057
_ = let__5142

tmp6063 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5142)


if True == tmp6063 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5142)
return
}


}, 1)

tmp6064 := Call(__e, ns2_1set, symshen_4_5lsb_6, tmp6056)


_ = tmp6064

tmp6065 := MakeNative(func(__e *ControlFlow) {
V2774 := __e.Get(1)
_ = V2774
tmp6070 := Call(__e, PrimFunc(symshen_4hds_a_2), V2774, MakeInteger(93))


var ifres6066 Obj

if True == tmp6070 {
tmp6067 := Call(__e, PrimFunc(symtail), V2774)


let__5145 := tmp6067
_ = let__5145

tmp6068 := Call(__e, PrimFunc(symshen_4comb), let__5145, symshen_4skip)


ifres6066 = tmp6068


} else {
tmp6069 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6066 = tmp6069


}

let__5144 := ifres6066
_ = let__5144

tmp6072 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5144)


if True == tmp6072 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5144)
return
}


}, 1)

tmp6073 := Call(__e, ns2_1set, symshen_4_5rsb_6, tmp6065)


_ = tmp6073

tmp6074 := MakeNative(func(__e *ControlFlow) {
V2777 := __e.Get(1)
_ = V2777
tmp6075 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), V2777)


let__5147 := tmp6075
_ = let__5147

tmp6081 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5147)


var ifres6076 Obj

if True == tmp6081 {
tmp6077 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6076 = tmp6077


} else {
tmp6078 := Call(__e, PrimFunc(symshen_4_5_1out), let__5147)


let__5148 := tmp6078
_ = let__5148

tmp6079 := Call(__e, PrimFunc(symshen_4in_1_6), let__5147)


let__5149 := tmp6079
_ = let__5149

tmp6080 := Call(__e, PrimFunc(symshen_4comb), let__5149, let__5148)


ifres6076 = tmp6080


}

let__5146 := ifres6076
_ = let__5146

tmp6083 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5146)


if True == tmp6083 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5146)
return
}


}, 1)

tmp6084 := Call(__e, ns2_1set, symshen_4_5s_1exprs1_6, tmp6074)


_ = tmp6084

tmp6085 := MakeNative(func(__e *ControlFlow) {
V2782 := __e.Get(1)
_ = V2782
tmp6086 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), V2782)


let__5151 := tmp6086
_ = let__5151

tmp6092 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5151)


var ifres6087 Obj

if True == tmp6092 {
tmp6088 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6087 = tmp6088


} else {
tmp6089 := Call(__e, PrimFunc(symshen_4_5_1out), let__5151)


let__5152 := tmp6089
_ = let__5152

tmp6090 := Call(__e, PrimFunc(symshen_4in_1_6), let__5151)


let__5153 := tmp6090
_ = let__5153

tmp6091 := Call(__e, PrimFunc(symshen_4comb), let__5153, let__5152)


ifres6087 = tmp6091


}

let__5150 := ifres6087
_ = let__5150

tmp6094 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5150)


if True == tmp6094 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5150)
return
}


}, 1)

tmp6095 := Call(__e, ns2_1set, symshen_4_5s_1exprs2_6, tmp6085)


_ = tmp6095

tmp6096 := MakeNative(func(__e *ControlFlow) {
V2788 := __e.Get(1)
_ = V2788
tmp6153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2788)
}
__typedArg0 := Nil
__typedArg1 := V2788
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp6153 {
__e.Return(Nil)
return
} else {
tmp6151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2788)
}
__typedArg0 := V2788
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6131 Obj

if True == tmp6151 {
tmp6149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2788)
}
__typedArg0 := V2788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp6149)
}
__typedArg0 := tmp6149
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6133 Obj

if True == tmp6150 {
tmp6146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2788)
}
__typedArg0 := V2788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6146)
}
__typedArg0 := tmp6146
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp6147)
}
__typedArg0 := tmp6147
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6135 Obj

if True == tmp6148 {
tmp6142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2788)
}
__typedArg0 := V2788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6142)
}
__typedArg0 := tmp6142
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6143)
}
__typedArg0 := tmp6143
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp6144)
}
__typedArg0 := Nil
__typedArg1 := tmp6144
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres6137 Obj

if True == tmp6145 {
tmp6139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2788)
}
__typedArg0 := V2788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp6139)
}
__typedArg0 := tmp6139
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp6140, symbar_b)
}
__typedArg0 := tmp6140
__typedArg1 := symbar_b
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres6138 Obj

if True == tmp6141 {
ifres6138 = True


} else {
ifres6138 = False


}

ifres6137 = ifres6138


} else {
ifres6137 = False


}

var ifres6136 Obj

if True == ifres6137 {
ifres6136 = True


} else {
ifres6136 = False


}

ifres6135 = ifres6136


} else {
ifres6135 = False


}

var ifres6134 Obj

if True == ifres6135 {
ifres6134 = True


} else {
ifres6134 = False


}

ifres6133 = ifres6134


} else {
ifres6133 = False


}

var ifres6132 Obj

if True == ifres6133 {
ifres6132 = True


} else {
ifres6132 = False


}

ifres6131 = ifres6132


} else {
ifres6131 = False


}

if True == ifres6131 {
tmp6097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2788)
}
__typedArg0 := V2788
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2788)
}
__typedArg0 := V2788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6098)
}
__typedArg0 := tmp6098
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp6097, tmp6099)
}
__typedArg0 := tmp6097
__typedArg1 := tmp6099
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp6100)
}
__typedArg0 := symcons
__typedArg1 := tmp6100
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp6129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2788)
}
__typedArg0 := V2788
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6109 Obj

if True == tmp6129 {
tmp6127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2788)
}
__typedArg0 := V2788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp6127)
}
__typedArg0 := tmp6127
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6111 Obj

if True == tmp6128 {
tmp6124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2788)
}
__typedArg0 := V2788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6124)
}
__typedArg0 := tmp6124
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp6125)
}
__typedArg0 := tmp6125
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6113 Obj

if True == tmp6126 {
tmp6120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2788)
}
__typedArg0 := V2788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6120)
}
__typedArg0 := tmp6120
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6121)
}
__typedArg0 := tmp6121
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp6122)
}
__typedArg0 := tmp6122
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6115 Obj

if True == tmp6123 {
tmp6117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2788)
}
__typedArg0 := V2788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp6117)
}
__typedArg0 := tmp6117
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp6118, symbar_b)
}
__typedArg0 := tmp6118
__typedArg1 := symbar_b
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres6116 Obj

if True == tmp6119 {
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

var ifres6110 Obj

if True == ifres6111 {
ifres6110 = True


} else {
ifres6110 = False


}

ifres6109 = ifres6110


} else {
ifres6109 = False


}

if True == ifres6109 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("misapplication of |\n"))
}
__typedArg0 := MakeString("misapplication of |\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp6107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2788)
}
__typedArg0 := V2788
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp6107 {
tmp6101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2788)
}
__typedArg0 := V2788
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2788)
}
__typedArg0 := V2788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6103 := Call(__e, PrimFunc(symshen_4cons_1form), tmp6102)


tmp6104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp6103, Nil)
}
__typedArg0 := tmp6103
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp6101, tmp6104)
}
__typedArg0 := tmp6101
__typedArg1 := tmp6104
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp6105)
}
__typedArg0 := symcons
__typedArg1 := tmp6105
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4cons_1form)
return
}


}


}


}


}, 1)

tmp6154 := Call(__e, ns2_1set, symshen_4cons_1form, tmp6096)


_ = tmp6154

tmp6155 := MakeNative(func(__e *ControlFlow) {
V2789 := __e.Get(1)
_ = V2789
tmp6160 := Call(__e, PrimFunc(symshen_4hds_a_2), V2789, MakeInteger(40))


var ifres6156 Obj

if True == tmp6160 {
tmp6157 := Call(__e, PrimFunc(symtail), V2789)


let__5155 := tmp6157
_ = let__5155

tmp6158 := Call(__e, PrimFunc(symshen_4comb), let__5155, symshen_4skip)


ifres6156 = tmp6158


} else {
tmp6159 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6156 = tmp6159


}

let__5154 := ifres6156
_ = let__5154

tmp6162 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5154)


if True == tmp6162 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5154)
return
}


}, 1)

tmp6163 := Call(__e, ns2_1set, symshen_4_5lrb_6, tmp6155)


_ = tmp6163

tmp6164 := MakeNative(func(__e *ControlFlow) {
V2792 := __e.Get(1)
_ = V2792
tmp6169 := Call(__e, PrimFunc(symshen_4hds_a_2), V2792, MakeInteger(41))


var ifres6165 Obj

if True == tmp6169 {
tmp6166 := Call(__e, PrimFunc(symtail), V2792)


let__5157 := tmp6166
_ = let__5157

tmp6167 := Call(__e, PrimFunc(symshen_4comb), let__5157, symshen_4skip)


ifres6165 = tmp6167


} else {
tmp6168 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6165 = tmp6168


}

let__5156 := ifres6165
_ = let__5156

tmp6171 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5156)


if True == tmp6171 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5156)
return
}


}, 1)

tmp6172 := Call(__e, ns2_1set, symshen_4_5rrb_6, tmp6164)


_ = tmp6172

tmp6173 := MakeNative(func(__e *ControlFlow) {
V2795 := __e.Get(1)
_ = V2795
tmp6178 := Call(__e, PrimFunc(symshen_4hds_a_2), V2795, MakeInteger(123))


var ifres6174 Obj

if True == tmp6178 {
tmp6175 := Call(__e, PrimFunc(symtail), V2795)


let__5159 := tmp6175
_ = let__5159

tmp6176 := Call(__e, PrimFunc(symshen_4comb), let__5159, symshen_4skip)


ifres6174 = tmp6176


} else {
tmp6177 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6174 = tmp6177


}

let__5158 := ifres6174
_ = let__5158

tmp6180 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5158)


if True == tmp6180 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5158)
return
}


}, 1)

tmp6181 := Call(__e, ns2_1set, symshen_4_5lcurly_6, tmp6173)


_ = tmp6181

tmp6182 := MakeNative(func(__e *ControlFlow) {
V2798 := __e.Get(1)
_ = V2798
tmp6187 := Call(__e, PrimFunc(symshen_4hds_a_2), V2798, MakeInteger(125))


var ifres6183 Obj

if True == tmp6187 {
tmp6184 := Call(__e, PrimFunc(symtail), V2798)


let__5161 := tmp6184
_ = let__5161

tmp6185 := Call(__e, PrimFunc(symshen_4comb), let__5161, symshen_4skip)


ifres6183 = tmp6185


} else {
tmp6186 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6183 = tmp6186


}

let__5160 := ifres6183
_ = let__5160

tmp6189 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5160)


if True == tmp6189 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5160)
return
}


}, 1)

tmp6190 := Call(__e, ns2_1set, symshen_4_5rcurly_6, tmp6182)


_ = tmp6190

tmp6191 := MakeNative(func(__e *ControlFlow) {
V2801 := __e.Get(1)
_ = V2801
tmp6196 := Call(__e, PrimFunc(symshen_4hds_a_2), V2801, MakeInteger(124))


var ifres6192 Obj

if True == tmp6196 {
tmp6193 := Call(__e, PrimFunc(symtail), V2801)


let__5163 := tmp6193
_ = let__5163

tmp6194 := Call(__e, PrimFunc(symshen_4comb), let__5163, symshen_4skip)


ifres6192 = tmp6194


} else {
tmp6195 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6192 = tmp6195


}

let__5162 := ifres6192
_ = let__5162

tmp6198 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5162)


if True == tmp6198 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5162)
return
}


}, 1)

tmp6199 := Call(__e, ns2_1set, symshen_4_5bar_6, tmp6191)


_ = tmp6199

tmp6200 := MakeNative(func(__e *ControlFlow) {
V2804 := __e.Get(1)
_ = V2804
tmp6205 := Call(__e, PrimFunc(symshen_4hds_a_2), V2804, MakeInteger(59))


var ifres6201 Obj

if True == tmp6205 {
tmp6202 := Call(__e, PrimFunc(symtail), V2804)


let__5165 := tmp6202
_ = let__5165

tmp6203 := Call(__e, PrimFunc(symshen_4comb), let__5165, symshen_4skip)


ifres6201 = tmp6203


} else {
tmp6204 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6201 = tmp6204


}

let__5164 := ifres6201
_ = let__5164

tmp6207 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5164)


if True == tmp6207 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5164)
return
}


}, 1)

tmp6208 := Call(__e, ns2_1set, symshen_4_5semicolon_6, tmp6200)


_ = tmp6208

tmp6209 := MakeNative(func(__e *ControlFlow) {
V2807 := __e.Get(1)
_ = V2807
tmp6214 := Call(__e, PrimFunc(symshen_4hds_a_2), V2807, MakeInteger(58))


var ifres6210 Obj

if True == tmp6214 {
tmp6211 := Call(__e, PrimFunc(symtail), V2807)


let__5167 := tmp6211
_ = let__5167

tmp6212 := Call(__e, PrimFunc(symshen_4comb), let__5167, symshen_4skip)


ifres6210 = tmp6212


} else {
tmp6213 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6210 = tmp6213


}

let__5166 := ifres6210
_ = let__5166

tmp6216 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5166)


if True == tmp6216 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5166)
return
}


}, 1)

tmp6217 := Call(__e, ns2_1set, symshen_4_5colon_6, tmp6209)


_ = tmp6217

tmp6218 := MakeNative(func(__e *ControlFlow) {
V2810 := __e.Get(1)
_ = V2810
tmp6223 := Call(__e, PrimFunc(symshen_4hds_a_2), V2810, MakeInteger(44))


var ifres6219 Obj

if True == tmp6223 {
tmp6220 := Call(__e, PrimFunc(symtail), V2810)


let__5169 := tmp6220
_ = let__5169

tmp6221 := Call(__e, PrimFunc(symshen_4comb), let__5169, symshen_4skip)


ifres6219 = tmp6221


} else {
tmp6222 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6219 = tmp6222


}

let__5168 := ifres6219
_ = let__5168

tmp6225 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5168)


if True == tmp6225 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5168)
return
}


}, 1)

tmp6226 := Call(__e, ns2_1set, symshen_4_5comma_6, tmp6218)


_ = tmp6226

tmp6227 := MakeNative(func(__e *ControlFlow) {
V2813 := __e.Get(1)
_ = V2813
tmp6232 := Call(__e, PrimFunc(symshen_4hds_a_2), V2813, MakeInteger(61))


var ifres6228 Obj

if True == tmp6232 {
tmp6229 := Call(__e, PrimFunc(symtail), V2813)


let__5171 := tmp6229
_ = let__5171

tmp6230 := Call(__e, PrimFunc(symshen_4comb), let__5171, symshen_4skip)


ifres6228 = tmp6230


} else {
tmp6231 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6228 = tmp6231


}

let__5170 := ifres6228
_ = let__5170

tmp6234 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5170)


if True == tmp6234 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5170)
return
}


}, 1)

tmp6235 := Call(__e, ns2_1set, symshen_4_5equal_6, tmp6227)


_ = tmp6235

tmp6236 := MakeNative(func(__e *ControlFlow) {
V2816 := __e.Get(1)
_ = V2816
tmp6237 := Call(__e, PrimFunc(symshen_4_5singleline_6), V2816)


let__5173 := tmp6237
_ = let__5173

tmp6242 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5173)


var ifres6238 Obj

if True == tmp6242 {
tmp6239 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6238 = tmp6239


} else {
tmp6240 := Call(__e, PrimFunc(symshen_4in_1_6), let__5173)


let__5174 := tmp6240
_ = let__5174

tmp6241 := Call(__e, PrimFunc(symshen_4comb), let__5174, symshen_4skip)


ifres6238 = tmp6241


}

let__5172 := ifres6238
_ = let__5172

tmp6252 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5172)


if True == tmp6252 {
tmp6243 := Call(__e, PrimFunc(symshen_4_5multiline_6), V2816)


let__5176 := tmp6243
_ = let__5176

tmp6248 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5176)


var ifres6244 Obj

if True == tmp6248 {
tmp6245 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6244 = tmp6245


} else {
tmp6246 := Call(__e, PrimFunc(symshen_4in_1_6), let__5176)


let__5177 := tmp6246
_ = let__5177

tmp6247 := Call(__e, PrimFunc(symshen_4comb), let__5177, symshen_4skip)


ifres6244 = tmp6247


}

let__5175 := ifres6244
_ = let__5175

tmp6250 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5175)


if True == tmp6250 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5175)
return
}


} else {
__e.Return(let__5172)
return
}


}, 1)

tmp6253 := Call(__e, ns2_1set, symshen_4_5comment_6, tmp6236)


_ = tmp6253

tmp6254 := MakeNative(func(__e *ControlFlow) {
V2823 := __e.Get(1)
_ = V2823
tmp6255 := Call(__e, PrimFunc(symshen_4_5backslash_6), V2823)


let__5179 := tmp6255
_ = let__5179

tmp6275 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5179)


var ifres6256 Obj

if True == tmp6275 {
tmp6257 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6256 = tmp6257


} else {
tmp6258 := Call(__e, PrimFunc(symshen_4in_1_6), let__5179)


let__5180 := tmp6258
_ = let__5180

tmp6259 := Call(__e, PrimFunc(symshen_4_5backslash_6), let__5180)


let__5181 := tmp6259
_ = let__5181

tmp6274 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5181)


var ifres6260 Obj

if True == tmp6274 {
tmp6261 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6260 = tmp6261


} else {
tmp6262 := Call(__e, PrimFunc(symshen_4in_1_6), let__5181)


let__5182 := tmp6262
_ = let__5182

tmp6263 := Call(__e, PrimFunc(symshen_4_5shortnatters_6), let__5182)


let__5183 := tmp6263
_ = let__5183

tmp6273 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5183)


var ifres6264 Obj

if True == tmp6273 {
tmp6265 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6264 = tmp6265


} else {
tmp6266 := Call(__e, PrimFunc(symshen_4in_1_6), let__5183)


let__5184 := tmp6266
_ = let__5184

tmp6267 := Call(__e, PrimFunc(symshen_4_5returns_6), let__5184)


let__5185 := tmp6267
_ = let__5185

tmp6272 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5185)


var ifres6268 Obj

if True == tmp6272 {
tmp6269 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6268 = tmp6269


} else {
tmp6270 := Call(__e, PrimFunc(symshen_4in_1_6), let__5185)


let__5186 := tmp6270
_ = let__5186

tmp6271 := Call(__e, PrimFunc(symshen_4comb), let__5186, symshen_4skip)


ifres6268 = tmp6271


}

ifres6264 = ifres6268


}

ifres6260 = ifres6264


}

ifres6256 = ifres6260


}

let__5178 := ifres6256
_ = let__5178

tmp6277 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5178)


if True == tmp6277 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5178)
return
}


}, 1)

tmp6278 := Call(__e, ns2_1set, symshen_4_5singleline_6, tmp6254)


_ = tmp6278

tmp6279 := MakeNative(func(__e *ControlFlow) {
V2833 := __e.Get(1)
_ = V2833
tmp6284 := Call(__e, PrimFunc(symshen_4hds_a_2), V2833, MakeInteger(92))


var ifres6280 Obj

if True == tmp6284 {
tmp6281 := Call(__e, PrimFunc(symtail), V2833)


let__5188 := tmp6281
_ = let__5188

tmp6282 := Call(__e, PrimFunc(symshen_4comb), let__5188, symshen_4skip)


ifres6280 = tmp6282


} else {
tmp6283 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6280 = tmp6283


}

let__5187 := ifres6280
_ = let__5187

tmp6286 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5187)


if True == tmp6286 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5187)
return
}


}, 1)

tmp6287 := Call(__e, ns2_1set, symshen_4_5backslash_6, tmp6279)


_ = tmp6287

tmp6288 := MakeNative(func(__e *ControlFlow) {
V2836 := __e.Get(1)
_ = V2836
tmp6289 := Call(__e, PrimFunc(symshen_4_5shortnatter_6), V2836)


let__5190 := tmp6289
_ = let__5190

tmp6299 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5190)


var ifres6290 Obj

if True == tmp6299 {
tmp6291 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6290 = tmp6291


} else {
tmp6292 := Call(__e, PrimFunc(symshen_4in_1_6), let__5190)


let__5191 := tmp6292
_ = let__5191

tmp6293 := Call(__e, PrimFunc(symshen_4_5shortnatters_6), let__5191)


let__5192 := tmp6293
_ = let__5192

tmp6298 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5192)


var ifres6294 Obj

if True == tmp6298 {
tmp6295 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6294 = tmp6295


} else {
tmp6296 := Call(__e, PrimFunc(symshen_4in_1_6), let__5192)


let__5193 := tmp6296
_ = let__5193

tmp6297 := Call(__e, PrimFunc(symshen_4comb), let__5193, symshen_4skip)


ifres6294 = tmp6297


}

ifres6290 = ifres6294


}

let__5189 := ifres6290
_ = let__5189

tmp6309 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5189)


if True == tmp6309 {
tmp6300 := Call(__e, PrimFunc(sym_5e_6), V2836)


let__5195 := tmp6300
_ = let__5195

tmp6305 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5195)


var ifres6301 Obj

if True == tmp6305 {
tmp6302 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6301 = tmp6302


} else {
tmp6303 := Call(__e, PrimFunc(symshen_4in_1_6), let__5195)


let__5196 := tmp6303
_ = let__5196

tmp6304 := Call(__e, PrimFunc(symshen_4comb), let__5196, symshen_4skip)


ifres6301 = tmp6304


}

let__5194 := ifres6301
_ = let__5194

tmp6307 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5194)


if True == tmp6307 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5194)
return
}


} else {
__e.Return(let__5189)
return
}


}, 1)

tmp6310 := Call(__e, ns2_1set, symshen_4_5shortnatters_6, tmp6288)


_ = tmp6310

tmp6311 := MakeNative(func(__e *ControlFlow) {
V2845 := __e.Get(1)
_ = V2845
tmp6321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2845)
}
__typedArg0 := V2845
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6312 Obj

if True == tmp6321 {
tmp6313 := Call(__e, PrimFunc(symhead), V2845)


let__5198 := tmp6313
_ = let__5198

tmp6314 := Call(__e, PrimFunc(symtail), V2845)


let__5199 := tmp6314
_ = let__5199

tmp6318 := Call(__e, PrimFunc(symshen_4return_2), let__5198)


tmp6319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp6318)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp6318
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres6315 Obj

if True == tmp6319 {
tmp6316 := Call(__e, PrimFunc(symshen_4comb), let__5199, symshen_4skip)


ifres6315 = tmp6316


} else {
tmp6317 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6315 = tmp6317


}

ifres6312 = ifres6315


} else {
tmp6320 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6312 = tmp6320


}

let__5197 := ifres6312
_ = let__5197

tmp6323 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5197)


if True == tmp6323 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5197)
return
}


}, 1)

tmp6324 := Call(__e, ns2_1set, symshen_4_5shortnatter_6, tmp6311)


_ = tmp6324

tmp6325 := MakeNative(func(__e *ControlFlow) {
V2849 := __e.Get(1)
_ = V2849
tmp6326 := Call(__e, PrimFunc(symshen_4_5return_6), V2849)


let__5201 := tmp6326
_ = let__5201

tmp6336 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5201)


var ifres6327 Obj

if True == tmp6336 {
tmp6328 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6327 = tmp6328


} else {
tmp6329 := Call(__e, PrimFunc(symshen_4in_1_6), let__5201)


let__5202 := tmp6329
_ = let__5202

tmp6330 := Call(__e, PrimFunc(symshen_4_5returns_6), let__5202)


let__5203 := tmp6330
_ = let__5203

tmp6335 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5203)


var ifres6331 Obj

if True == tmp6335 {
tmp6332 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6331 = tmp6332


} else {
tmp6333 := Call(__e, PrimFunc(symshen_4in_1_6), let__5203)


let__5204 := tmp6333
_ = let__5204

tmp6334 := Call(__e, PrimFunc(symshen_4comb), let__5204, symshen_4skip)


ifres6331 = tmp6334


}

ifres6327 = ifres6331


}

let__5200 := ifres6327
_ = let__5200

tmp6346 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5200)


if True == tmp6346 {
tmp6337 := Call(__e, PrimFunc(symshen_4_5return_6), V2849)


let__5206 := tmp6337
_ = let__5206

tmp6342 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5206)


var ifres6338 Obj

if True == tmp6342 {
tmp6339 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6338 = tmp6339


} else {
tmp6340 := Call(__e, PrimFunc(symshen_4in_1_6), let__5206)


let__5207 := tmp6340
_ = let__5207

tmp6341 := Call(__e, PrimFunc(symshen_4comb), let__5207, symshen_4skip)


ifres6338 = tmp6341


}

let__5205 := ifres6338
_ = let__5205

tmp6344 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5205)


if True == tmp6344 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5205)
return
}


} else {
__e.Return(let__5200)
return
}


}, 1)

tmp6347 := Call(__e, ns2_1set, symshen_4_5returns_6, tmp6325)


_ = tmp6347

tmp6348 := MakeNative(func(__e *ControlFlow) {
V2858 := __e.Get(1)
_ = V2858
tmp6357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2858)
}
__typedArg0 := V2858
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6349 Obj

if True == tmp6357 {
tmp6350 := Call(__e, PrimFunc(symhead), V2858)


let__5209 := tmp6350
_ = let__5209

tmp6351 := Call(__e, PrimFunc(symtail), V2858)


let__5210 := tmp6351
_ = let__5210

tmp6355 := Call(__e, PrimFunc(symshen_4return_2), let__5209)


var ifres6352 Obj

if True == tmp6355 {
tmp6353 := Call(__e, PrimFunc(symshen_4comb), let__5210, symshen_4skip)


ifres6352 = tmp6353


} else {
tmp6354 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6352 = tmp6354


}

ifres6349 = ifres6352


} else {
tmp6356 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6349 = tmp6356


}

let__5208 := ifres6349
_ = let__5208

tmp6359 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5208)


if True == tmp6359 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5208)
return
}


}, 1)

tmp6360 := Call(__e, ns2_1set, symshen_4_5return_6, tmp6348)


_ = tmp6360

tmp6361 := MakeNative(func(__e *ControlFlow) {
V2862 := __e.Get(1)
_ = V2862
tmp6362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(13), Nil)
}
__typedArg0 := MakeInteger(13)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(10), tmp6362)
}
__typedArg0 := MakeInteger(10)
__typedArg1 := tmp6362
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(9), tmp6363)
}
__typedArg0 := MakeInteger(9)
__typedArg1 := tmp6363
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symelement_2), V2862, tmp6364)
return


}, 1)

tmp6365 := Call(__e, ns2_1set, symshen_4return_2, tmp6361)


_ = tmp6365

tmp6366 := MakeNative(func(__e *ControlFlow) {
V2863 := __e.Get(1)
_ = V2863
tmp6367 := Call(__e, PrimFunc(symshen_4_5backslash_6), V2863)


let__5212 := tmp6367
_ = let__5212

tmp6382 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5212)


var ifres6368 Obj

if True == tmp6382 {
tmp6369 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6368 = tmp6369


} else {
tmp6370 := Call(__e, PrimFunc(symshen_4in_1_6), let__5212)


let__5213 := tmp6370
_ = let__5213

tmp6371 := Call(__e, PrimFunc(symshen_4_5times_6), let__5213)


let__5214 := tmp6371
_ = let__5214

tmp6381 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5214)


var ifres6372 Obj

if True == tmp6381 {
tmp6373 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6372 = tmp6373


} else {
tmp6374 := Call(__e, PrimFunc(symshen_4in_1_6), let__5214)


let__5215 := tmp6374
_ = let__5215

tmp6375 := Call(__e, PrimFunc(symshen_4_5longnatter_6), let__5215)


let__5216 := tmp6375
_ = let__5216

tmp6380 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5216)


var ifres6376 Obj

if True == tmp6380 {
tmp6377 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6376 = tmp6377


} else {
tmp6378 := Call(__e, PrimFunc(symshen_4in_1_6), let__5216)


let__5217 := tmp6378
_ = let__5217

tmp6379 := Call(__e, PrimFunc(symshen_4comb), let__5217, symshen_4skip)


ifres6376 = tmp6379


}

ifres6372 = ifres6376


}

ifres6368 = ifres6372


}

let__5211 := ifres6368
_ = let__5211

tmp6384 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5211)


if True == tmp6384 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5211)
return
}


}, 1)

tmp6385 := Call(__e, ns2_1set, symshen_4_5multiline_6, tmp6366)


_ = tmp6385

tmp6386 := MakeNative(func(__e *ControlFlow) {
V2871 := __e.Get(1)
_ = V2871
tmp6391 := Call(__e, PrimFunc(symshen_4hds_a_2), V2871, MakeInteger(42))


var ifres6387 Obj

if True == tmp6391 {
tmp6388 := Call(__e, PrimFunc(symtail), V2871)


let__5219 := tmp6388
_ = let__5219

tmp6389 := Call(__e, PrimFunc(symshen_4comb), let__5219, symshen_4skip)


ifres6387 = tmp6389


} else {
tmp6390 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6387 = tmp6390


}

let__5218 := ifres6387
_ = let__5218

tmp6393 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5218)


if True == tmp6393 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5218)
return
}


}, 1)

tmp6394 := Call(__e, ns2_1set, symshen_4_5times_6, tmp6386)


_ = tmp6394

tmp6395 := MakeNative(func(__e *ControlFlow) {
V2874 := __e.Get(1)
_ = V2874
tmp6396 := Call(__e, PrimFunc(symshen_4_5comment_6), V2874)


let__5221 := tmp6396
_ = let__5221

tmp6406 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5221)


var ifres6397 Obj

if True == tmp6406 {
tmp6398 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6397 = tmp6398


} else {
tmp6399 := Call(__e, PrimFunc(symshen_4in_1_6), let__5221)


let__5222 := tmp6399
_ = let__5222

tmp6400 := Call(__e, PrimFunc(symshen_4_5longnatter_6), let__5222)


let__5223 := tmp6400
_ = let__5223

tmp6405 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5223)


var ifres6401 Obj

if True == tmp6405 {
tmp6402 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6401 = tmp6402


} else {
tmp6403 := Call(__e, PrimFunc(symshen_4in_1_6), let__5223)


let__5224 := tmp6403
_ = let__5224

tmp6404 := Call(__e, PrimFunc(symshen_4comb), let__5224, symshen_4skip)


ifres6401 = tmp6404


}

ifres6397 = ifres6401


}

let__5220 := ifres6397
_ = let__5220

tmp6433 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5220)


if True == tmp6433 {
tmp6407 := Call(__e, PrimFunc(symshen_4_5times_6), V2874)


let__5226 := tmp6407
_ = let__5226

tmp6417 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5226)


var ifres6408 Obj

if True == tmp6417 {
tmp6409 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6408 = tmp6409


} else {
tmp6410 := Call(__e, PrimFunc(symshen_4in_1_6), let__5226)


let__5227 := tmp6410
_ = let__5227

tmp6411 := Call(__e, PrimFunc(symshen_4_5backslash_6), let__5227)


let__5228 := tmp6411
_ = let__5228

tmp6416 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5228)


var ifres6412 Obj

if True == tmp6416 {
tmp6413 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6412 = tmp6413


} else {
tmp6414 := Call(__e, PrimFunc(symshen_4in_1_6), let__5228)


let__5229 := tmp6414
_ = let__5229

tmp6415 := Call(__e, PrimFunc(symshen_4comb), let__5229, symshen_4skip)


ifres6412 = tmp6415


}

ifres6408 = ifres6412


}

let__5225 := ifres6408
_ = let__5225

tmp6431 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5225)


if True == tmp6431 {
tmp6427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2874)
}
__typedArg0 := V2874
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6418 Obj

if True == tmp6427 {
tmp6419 := Call(__e, PrimFunc(symtail), V2874)


let__5231 := tmp6419
_ = let__5231

tmp6420 := Call(__e, PrimFunc(symshen_4_5longnatter_6), let__5231)


let__5232 := tmp6420
_ = let__5232

tmp6425 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5232)


var ifres6421 Obj

if True == tmp6425 {
tmp6422 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6421 = tmp6422


} else {
tmp6423 := Call(__e, PrimFunc(symshen_4in_1_6), let__5232)


let__5233 := tmp6423
_ = let__5233

tmp6424 := Call(__e, PrimFunc(symshen_4comb), let__5233, symshen_4skip)


ifres6421 = tmp6424


}

ifres6418 = ifres6421


} else {
tmp6426 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6418 = tmp6426


}

let__5230 := ifres6418
_ = let__5230

tmp6429 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5230)


if True == tmp6429 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5230)
return
}


} else {
__e.Return(let__5225)
return
}


} else {
__e.Return(let__5220)
return
}


}, 1)

tmp6434 := Call(__e, ns2_1set, symshen_4_5longnatter_6, tmp6395)


_ = tmp6434

tmp6435 := MakeNative(func(__e *ControlFlow) {
V2889 := __e.Get(1)
_ = V2889
tmp6436 := Call(__e, PrimFunc(symshen_4_5str_6), V2889)


let__5235 := tmp6436
_ = let__5235

tmp6442 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5235)


var ifres6437 Obj

if True == tmp6442 {
tmp6438 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6437 = tmp6438


} else {
tmp6439 := Call(__e, PrimFunc(symshen_4_5_1out), let__5235)


let__5236 := tmp6439
_ = let__5236

tmp6440 := Call(__e, PrimFunc(symshen_4in_1_6), let__5235)


let__5237 := tmp6440
_ = let__5237

tmp6441 := Call(__e, PrimFunc(symshen_4comb), let__5237, let__5236)


ifres6437 = tmp6441


}

let__5234 := ifres6437
_ = let__5234

tmp6467 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5234)


if True == tmp6467 {
tmp6443 := Call(__e, PrimFunc(symshen_4_5number_6), V2889)


let__5239 := tmp6443
_ = let__5239

tmp6449 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5239)


var ifres6444 Obj

if True == tmp6449 {
tmp6445 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6444 = tmp6445


} else {
tmp6446 := Call(__e, PrimFunc(symshen_4_5_1out), let__5239)


let__5240 := tmp6446
_ = let__5240

tmp6447 := Call(__e, PrimFunc(symshen_4in_1_6), let__5239)


let__5241 := tmp6447
_ = let__5241

tmp6448 := Call(__e, PrimFunc(symshen_4comb), let__5241, let__5240)


ifres6444 = tmp6448


}

let__5238 := ifres6444
_ = let__5238

tmp6465 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5238)


if True == tmp6465 {
tmp6450 := Call(__e, PrimFunc(symshen_4_5sym_6), V2889)


let__5243 := tmp6450
_ = let__5243

tmp6461 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5243)


var ifres6451 Obj

if True == tmp6461 {
tmp6452 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6451 = tmp6452


} else {
tmp6453 := Call(__e, PrimFunc(symshen_4_5_1out), let__5243)


let__5244 := tmp6453
_ = let__5244

tmp6454 := Call(__e, PrimFunc(symshen_4in_1_6), let__5243)


let__5245 := tmp6454
_ = let__5245

tmp6459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__5244, MakeString("<>"))
}
__typedArg0 := let__5244
__typedArg1 := MakeString("<>")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres6455 Obj

if True == tmp6459 {
tmp6456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), Nil)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp6456)
}
__typedArg0 := symvector
__typedArg1 := tmp6456
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

ifres6455 = tmp6457


} else {
tmp6458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(let__5244)
}
__typedArg0 := let__5244
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

ifres6455 = tmp6458


}

tmp6460 := Call(__e, PrimFunc(symshen_4comb), let__5245, ifres6455)


ifres6451 = tmp6460


}

let__5242 := ifres6451
_ = let__5242

tmp6463 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5242)


if True == tmp6463 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5242)
return
}


} else {
__e.Return(let__5238)
return
}


} else {
__e.Return(let__5234)
return
}


}, 1)

tmp6468 := Call(__e, ns2_1set, symshen_4_5atom_6, tmp6435)


_ = tmp6468

tmp6469 := MakeNative(func(__e *ControlFlow) {
V2902 := __e.Get(1)
_ = V2902
tmp6470 := Call(__e, PrimFunc(symshen_4_5alpha_6), V2902)


let__5247 := tmp6470
_ = let__5247

tmp6483 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5247)


var ifres6471 Obj

if True == tmp6483 {
tmp6472 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6471 = tmp6472


} else {
tmp6473 := Call(__e, PrimFunc(symshen_4_5_1out), let__5247)


let__5248 := tmp6473
_ = let__5248

tmp6474 := Call(__e, PrimFunc(symshen_4in_1_6), let__5247)


let__5249 := tmp6474
_ = let__5249

tmp6475 := Call(__e, PrimFunc(symshen_4_5alphanums_6), let__5249)


let__5250 := tmp6475
_ = let__5250

tmp6482 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5250)


var ifres6476 Obj

if True == tmp6482 {
tmp6477 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6476 = tmp6477


} else {
tmp6478 := Call(__e, PrimFunc(symshen_4_5_1out), let__5250)


let__5251 := tmp6478
_ = let__5251

tmp6479 := Call(__e, PrimFunc(symshen_4in_1_6), let__5250)


let__5252 := tmp6479
_ = let__5252

tmp6481 := Call(__e, PrimFunc(symshen_4comb), let__5252, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(let__5248)
__typedS1, __typedOK1 := TypedString(let__5251)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := let__5248
__typedArg1 := let__5251
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())


ifres6476 = tmp6481


}

ifres6471 = ifres6476


}

let__5246 := ifres6471
_ = let__5246

tmp6485 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5246)


if True == tmp6485 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5246)
return
}


}, 1)

tmp6486 := Call(__e, ns2_1set, symshen_4_5sym_6, tmp6469)


_ = tmp6486

tmp6487 := MakeNative(func(__e *ControlFlow) {
V2910 := __e.Get(1)
_ = V2910
tmp6497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2910)
}
__typedArg0 := V2910
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6488 Obj

if True == tmp6497 {
tmp6489 := Call(__e, PrimFunc(symhead), V2910)


let__5254 := tmp6489
_ = let__5254

tmp6490 := Call(__e, PrimFunc(symtail), V2910)


let__5255 := tmp6490
_ = let__5255

tmp6495 := Call(__e, PrimFunc(symshen_4alpha_2), let__5254)


var ifres6491 Obj

if True == tmp6495 {
tmp6492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(let__5254)
}
__typedArg0 := let__5254
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp6493 := Call(__e, PrimFunc(symshen_4comb), let__5255, tmp6492)


ifres6491 = tmp6493


} else {
tmp6494 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6491 = tmp6494


}

ifres6488 = ifres6491


} else {
tmp6496 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6488 = tmp6496


}

let__5253 := ifres6488
_ = let__5253

tmp6499 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5253)


if True == tmp6499 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5253)
return
}


}, 1)

tmp6500 := Call(__e, ns2_1set, symshen_4_5alpha_6, tmp6487)


_ = tmp6500

tmp6501 := MakeNative(func(__e *ControlFlow) {
V2914 := __e.Get(1)
_ = V2914
tmp6508 := Call(__e, PrimFunc(symshen_4lowercase_2), V2914)


if True == tmp6508 {
__e.Return(True)
return
} else {
tmp6506 := Call(__e, PrimFunc(symshen_4uppercase_2), V2914)


var ifres6503 Obj

if True == tmp6506 {
ifres6503 = True


} else {
tmp6505 := Call(__e, PrimFunc(symshen_4misc_2), V2914)


var ifres6504 Obj

if True == tmp6505 {
ifres6504 = True


} else {
ifres6504 = False


}

ifres6503 = ifres6504


}

if True == ifres6503 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp6509 := Call(__e, ns2_1set, symshen_4alpha_2, tmp6501)


_ = tmp6509

tmp6510 := MakeNative(func(__e *ControlFlow) {
V2915 := __e.Get(1)
_ = V2915
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6_a) {
__typedN0, __typedOK0 := TypedFloat64(V2915)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(97))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6_a) {
return TypedMaterializeBoolean((__typedN0 >= __typedN1))
}}
__typedArg0 := V2915
__typedArg1 := MakeInteger(97)
return Call(__e, PrimFunc(sym_6_a), __typedArg0, __typedArg1)
})() {
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_a) {
__typedN0, __typedOK0 := TypedFloat64(V2915)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(122))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5_a) {
return TypedMaterializeBoolean((__typedN0 <= __typedN1))
}}
__typedArg0 := V2915
__typedArg1 := MakeInteger(122)
return Call(__e, PrimFunc(sym_5_a), __typedArg0, __typedArg1)
})() {
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

tmp6515 := Call(__e, ns2_1set, symshen_4lowercase_2, tmp6510)


_ = tmp6515

tmp6516 := MakeNative(func(__e *ControlFlow) {
V2916 := __e.Get(1)
_ = V2916
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6_a) {
__typedN0, __typedOK0 := TypedFloat64(V2916)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(65))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6_a) {
return TypedMaterializeBoolean((__typedN0 >= __typedN1))
}}
__typedArg0 := V2916
__typedArg1 := MakeInteger(65)
return Call(__e, PrimFunc(sym_6_a), __typedArg0, __typedArg1)
})() {
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_a) {
__typedN0, __typedOK0 := TypedFloat64(V2916)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(90))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5_a) {
return TypedMaterializeBoolean((__typedN0 <= __typedN1))
}}
__typedArg0 := V2916
__typedArg1 := MakeInteger(90)
return Call(__e, PrimFunc(sym_5_a), __typedArg0, __typedArg1)
})() {
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

tmp6521 := Call(__e, ns2_1set, symshen_4uppercase_2, tmp6516)


_ = tmp6521

tmp6522 := MakeNative(func(__e *ControlFlow) {
V2917 := __e.Get(1)
_ = V2917
tmp6523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(96), Nil)
}
__typedArg0 := MakeInteger(96)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(35), tmp6523)
}
__typedArg0 := MakeInteger(35)
__typedArg1 := tmp6523
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(39), tmp6524)
}
__typedArg0 := MakeInteger(39)
__typedArg1 := tmp6524
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(37), tmp6525)
}
__typedArg0 := MakeInteger(37)
__typedArg1 := tmp6525
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(38), tmp6526)
}
__typedArg0 := MakeInteger(38)
__typedArg1 := tmp6526
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(60), tmp6527)
}
__typedArg0 := MakeInteger(60)
__typedArg1 := tmp6527
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(62), tmp6528)
}
__typedArg0 := MakeInteger(62)
__typedArg1 := tmp6528
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(46), tmp6529)
}
__typedArg0 := MakeInteger(46)
__typedArg1 := tmp6529
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(126), tmp6530)
}
__typedArg0 := MakeInteger(126)
__typedArg1 := tmp6530
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(64), tmp6531)
}
__typedArg0 := MakeInteger(64)
__typedArg1 := tmp6531
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(33), tmp6532)
}
__typedArg0 := MakeInteger(33)
__typedArg1 := tmp6532
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(36), tmp6533)
}
__typedArg0 := MakeInteger(36)
__typedArg1 := tmp6533
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(63), tmp6534)
}
__typedArg0 := MakeInteger(63)
__typedArg1 := tmp6534
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(95), tmp6535)
}
__typedArg0 := MakeInteger(95)
__typedArg1 := tmp6535
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(43), tmp6536)
}
__typedArg0 := MakeInteger(43)
__typedArg1 := tmp6536
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(47), tmp6537)
}
__typedArg0 := MakeInteger(47)
__typedArg1 := tmp6537
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(42), tmp6538)
}
__typedArg0 := MakeInteger(42)
__typedArg1 := tmp6538
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(45), tmp6539)
}
__typedArg0 := MakeInteger(45)
__typedArg1 := tmp6539
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(61), tmp6540)
}
__typedArg0 := MakeInteger(61)
__typedArg1 := tmp6540
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symelement_2), V2917, tmp6541)
return


}, 1)

tmp6542 := Call(__e, ns2_1set, symshen_4misc_2, tmp6522)


_ = tmp6542

tmp6543 := MakeNative(func(__e *ControlFlow) {
V2918 := __e.Get(1)
_ = V2918
tmp6544 := Call(__e, PrimFunc(symshen_4_5alphanum_6), V2918)


let__5257 := tmp6544
_ = let__5257

tmp6557 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5257)


var ifres6545 Obj

if True == tmp6557 {
tmp6546 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6545 = tmp6546


} else {
tmp6547 := Call(__e, PrimFunc(symshen_4_5_1out), let__5257)


let__5258 := tmp6547
_ = let__5258

tmp6548 := Call(__e, PrimFunc(symshen_4in_1_6), let__5257)


let__5259 := tmp6548
_ = let__5259

tmp6549 := Call(__e, PrimFunc(symshen_4_5alphanums_6), let__5259)


let__5260 := tmp6549
_ = let__5260

tmp6556 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5260)


var ifres6550 Obj

if True == tmp6556 {
tmp6551 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6550 = tmp6551


} else {
tmp6552 := Call(__e, PrimFunc(symshen_4_5_1out), let__5260)


let__5261 := tmp6552
_ = let__5261

tmp6553 := Call(__e, PrimFunc(symshen_4in_1_6), let__5260)


let__5262 := tmp6553
_ = let__5262

tmp6555 := Call(__e, PrimFunc(symshen_4comb), let__5262, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(let__5258)
__typedS1, __typedOK1 := TypedString(let__5261)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := let__5258
__typedArg1 := let__5261
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())


ifres6550 = tmp6555


}

ifres6545 = ifres6550


}

let__5256 := ifres6545
_ = let__5256

tmp6567 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5256)


if True == tmp6567 {
tmp6558 := Call(__e, PrimFunc(sym_5e_6), V2918)


let__5264 := tmp6558
_ = let__5264

tmp6563 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5264)


var ifres6559 Obj

if True == tmp6563 {
tmp6560 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6559 = tmp6560


} else {
tmp6561 := Call(__e, PrimFunc(symshen_4in_1_6), let__5264)


let__5265 := tmp6561
_ = let__5265

tmp6562 := Call(__e, PrimFunc(symshen_4comb), let__5265, MakeString(""))


ifres6559 = tmp6562


}

let__5263 := ifres6559
_ = let__5263

tmp6565 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5263)


if True == tmp6565 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5263)
return
}


} else {
__e.Return(let__5256)
return
}


}, 1)

tmp6568 := Call(__e, ns2_1set, symshen_4_5alphanums_6, tmp6543)


_ = tmp6568

tmp6569 := MakeNative(func(__e *ControlFlow) {
V2929 := __e.Get(1)
_ = V2929
tmp6570 := Call(__e, PrimFunc(symshen_4_5alpha_6), V2929)


let__5267 := tmp6570
_ = let__5267

tmp6576 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5267)


var ifres6571 Obj

if True == tmp6576 {
tmp6572 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6571 = tmp6572


} else {
tmp6573 := Call(__e, PrimFunc(symshen_4_5_1out), let__5267)


let__5268 := tmp6573
_ = let__5268

tmp6574 := Call(__e, PrimFunc(symshen_4in_1_6), let__5267)


let__5269 := tmp6574
_ = let__5269

tmp6575 := Call(__e, PrimFunc(symshen_4comb), let__5269, let__5268)


ifres6571 = tmp6575


}

let__5266 := ifres6571
_ = let__5266

tmp6587 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5266)


if True == tmp6587 {
tmp6577 := Call(__e, PrimFunc(symshen_4_5numeral_6), V2929)


let__5271 := tmp6577
_ = let__5271

tmp6583 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5271)


var ifres6578 Obj

if True == tmp6583 {
tmp6579 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6578 = tmp6579


} else {
tmp6580 := Call(__e, PrimFunc(symshen_4_5_1out), let__5271)


let__5272 := tmp6580
_ = let__5272

tmp6581 := Call(__e, PrimFunc(symshen_4in_1_6), let__5271)


let__5273 := tmp6581
_ = let__5273

tmp6582 := Call(__e, PrimFunc(symshen_4comb), let__5273, let__5272)


ifres6578 = tmp6582


}

let__5270 := ifres6578
_ = let__5270

tmp6585 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5270)


if True == tmp6585 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5270)
return
}


} else {
__e.Return(let__5266)
return
}


}, 1)

tmp6588 := Call(__e, ns2_1set, symshen_4_5alphanum_6, tmp6569)


_ = tmp6588

tmp6589 := MakeNative(func(__e *ControlFlow) {
V2938 := __e.Get(1)
_ = V2938
tmp6599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2938)
}
__typedArg0 := V2938
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6590 Obj

if True == tmp6599 {
tmp6591 := Call(__e, PrimFunc(symhead), V2938)


let__5275 := tmp6591
_ = let__5275

tmp6592 := Call(__e, PrimFunc(symtail), V2938)


let__5276 := tmp6592
_ = let__5276

tmp6597 := Call(__e, PrimFunc(symshen_4digit_2), let__5275)


var ifres6593 Obj

if True == tmp6597 {
tmp6594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(let__5275)
}
__typedArg0 := let__5275
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp6595 := Call(__e, PrimFunc(symshen_4comb), let__5276, tmp6594)


ifres6593 = tmp6595


} else {
tmp6596 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6593 = tmp6596


}

ifres6590 = ifres6593


} else {
tmp6598 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6590 = tmp6598


}

let__5274 := ifres6590
_ = let__5274

tmp6601 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5274)


if True == tmp6601 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5274)
return
}


}, 1)

tmp6602 := Call(__e, ns2_1set, symshen_4_5numeral_6, tmp6589)


_ = tmp6602

tmp6603 := MakeNative(func(__e *ControlFlow) {
V2942 := __e.Get(1)
_ = V2942
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6_a) {
__typedN0, __typedOK0 := TypedFloat64(V2942)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(48))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6_a) {
return TypedMaterializeBoolean((__typedN0 >= __typedN1))
}}
__typedArg0 := V2942
__typedArg1 := MakeInteger(48)
return Call(__e, PrimFunc(sym_6_a), __typedArg0, __typedArg1)
})() {
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_a) {
__typedN0, __typedOK0 := TypedFloat64(V2942)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(57))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5_a) {
return TypedMaterializeBoolean((__typedN0 <= __typedN1))
}}
__typedArg0 := V2942
__typedArg1 := MakeInteger(57)
return Call(__e, PrimFunc(sym_5_a), __typedArg0, __typedArg1)
})() {
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

tmp6608 := Call(__e, ns2_1set, symshen_4digit_2, tmp6603)


_ = tmp6608

tmp6609 := MakeNative(func(__e *ControlFlow) {
V2943 := __e.Get(1)
_ = V2943
tmp6610 := Call(__e, PrimFunc(symshen_4_5dbq_6), V2943)


let__5278 := tmp6610
_ = let__5278

tmp6626 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5278)


var ifres6611 Obj

if True == tmp6626 {
tmp6612 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6611 = tmp6612


} else {
tmp6613 := Call(__e, PrimFunc(symshen_4in_1_6), let__5278)


let__5279 := tmp6613
_ = let__5279

tmp6614 := Call(__e, PrimFunc(symshen_4_5strcontents_6), let__5279)


let__5280 := tmp6614
_ = let__5280

tmp6625 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5280)


var ifres6615 Obj

if True == tmp6625 {
tmp6616 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6615 = tmp6616


} else {
tmp6617 := Call(__e, PrimFunc(symshen_4_5_1out), let__5280)


let__5281 := tmp6617
_ = let__5281

tmp6618 := Call(__e, PrimFunc(symshen_4in_1_6), let__5280)


let__5282 := tmp6618
_ = let__5282

tmp6619 := Call(__e, PrimFunc(symshen_4_5dbq_6), let__5282)


let__5283 := tmp6619
_ = let__5283

tmp6624 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5283)


var ifres6620 Obj

if True == tmp6624 {
tmp6621 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6620 = tmp6621


} else {
tmp6622 := Call(__e, PrimFunc(symshen_4in_1_6), let__5283)


let__5284 := tmp6622
_ = let__5284

tmp6623 := Call(__e, PrimFunc(symshen_4comb), let__5284, let__5281)


ifres6620 = tmp6623


}

ifres6615 = ifres6620


}

ifres6611 = ifres6615


}

let__5277 := ifres6611
_ = let__5277

tmp6628 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5277)


if True == tmp6628 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5277)
return
}


}, 1)

tmp6629 := Call(__e, ns2_1set, symshen_4_5str_6, tmp6609)


_ = tmp6629

tmp6630 := MakeNative(func(__e *ControlFlow) {
V2952 := __e.Get(1)
_ = V2952
tmp6635 := Call(__e, PrimFunc(symshen_4hds_a_2), V2952, MakeInteger(34))


var ifres6631 Obj

if True == tmp6635 {
tmp6632 := Call(__e, PrimFunc(symtail), V2952)


let__5286 := tmp6632
_ = let__5286

tmp6633 := Call(__e, PrimFunc(symshen_4comb), let__5286, symshen_4skip)


ifres6631 = tmp6633


} else {
tmp6634 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6631 = tmp6634


}

let__5285 := ifres6631
_ = let__5285

tmp6637 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5285)


if True == tmp6637 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5285)
return
}


}, 1)

tmp6638 := Call(__e, ns2_1set, symshen_4_5dbq_6, tmp6630)


_ = tmp6638

tmp6639 := MakeNative(func(__e *ControlFlow) {
V2955 := __e.Get(1)
_ = V2955
tmp6640 := Call(__e, PrimFunc(symshen_4_5strc_6), V2955)


let__5288 := tmp6640
_ = let__5288

tmp6653 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5288)


var ifres6641 Obj

if True == tmp6653 {
tmp6642 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6641 = tmp6642


} else {
tmp6643 := Call(__e, PrimFunc(symshen_4_5_1out), let__5288)


let__5289 := tmp6643
_ = let__5289

tmp6644 := Call(__e, PrimFunc(symshen_4in_1_6), let__5288)


let__5290 := tmp6644
_ = let__5290

tmp6645 := Call(__e, PrimFunc(symshen_4_5strcontents_6), let__5290)


let__5291 := tmp6645
_ = let__5291

tmp6652 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5291)


var ifres6646 Obj

if True == tmp6652 {
tmp6647 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6646 = tmp6647


} else {
tmp6648 := Call(__e, PrimFunc(symshen_4_5_1out), let__5291)


let__5292 := tmp6648
_ = let__5292

tmp6649 := Call(__e, PrimFunc(symshen_4in_1_6), let__5291)


let__5293 := tmp6649
_ = let__5293

tmp6651 := Call(__e, PrimFunc(symshen_4comb), let__5293, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(let__5289)
__typedS1, __typedOK1 := TypedString(let__5292)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := let__5289
__typedArg1 := let__5292
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())


ifres6646 = tmp6651


}

ifres6641 = ifres6646


}

let__5287 := ifres6641
_ = let__5287

tmp6663 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5287)


if True == tmp6663 {
tmp6654 := Call(__e, PrimFunc(sym_5e_6), V2955)


let__5295 := tmp6654
_ = let__5295

tmp6659 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5295)


var ifres6655 Obj

if True == tmp6659 {
tmp6656 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6655 = tmp6656


} else {
tmp6657 := Call(__e, PrimFunc(symshen_4in_1_6), let__5295)


let__5296 := tmp6657
_ = let__5296

tmp6658 := Call(__e, PrimFunc(symshen_4comb), let__5296, MakeString(""))


ifres6655 = tmp6658


}

let__5294 := ifres6655
_ = let__5294

tmp6661 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5294)


if True == tmp6661 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5294)
return
}


} else {
__e.Return(let__5287)
return
}


}, 1)

tmp6664 := Call(__e, ns2_1set, symshen_4_5strcontents_6, tmp6639)


_ = tmp6664

tmp6665 := MakeNative(func(__e *ControlFlow) {
V2966 := __e.Get(1)
_ = V2966
tmp6666 := Call(__e, PrimFunc(symshen_4_5control_6), V2966)


let__5298 := tmp6666
_ = let__5298

tmp6672 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5298)


var ifres6667 Obj

if True == tmp6672 {
tmp6668 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6667 = tmp6668


} else {
tmp6669 := Call(__e, PrimFunc(symshen_4_5_1out), let__5298)


let__5299 := tmp6669
_ = let__5299

tmp6670 := Call(__e, PrimFunc(symshen_4in_1_6), let__5298)


let__5300 := tmp6670
_ = let__5300

tmp6671 := Call(__e, PrimFunc(symshen_4comb), let__5300, let__5299)


ifres6667 = tmp6671


}

let__5297 := ifres6667
_ = let__5297

tmp6683 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5297)


if True == tmp6683 {
tmp6673 := Call(__e, PrimFunc(symshen_4_5notdbq_6), V2966)


let__5302 := tmp6673
_ = let__5302

tmp6679 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5302)


var ifres6674 Obj

if True == tmp6679 {
tmp6675 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6674 = tmp6675


} else {
tmp6676 := Call(__e, PrimFunc(symshen_4_5_1out), let__5302)


let__5303 := tmp6676
_ = let__5303

tmp6677 := Call(__e, PrimFunc(symshen_4in_1_6), let__5302)


let__5304 := tmp6677
_ = let__5304

tmp6678 := Call(__e, PrimFunc(symshen_4comb), let__5304, let__5303)


ifres6674 = tmp6678


}

let__5301 := ifres6674
_ = let__5301

tmp6681 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5301)


if True == tmp6681 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5301)
return
}


} else {
__e.Return(let__5297)
return
}


}, 1)

tmp6684 := Call(__e, ns2_1set, symshen_4_5strc_6, tmp6665)


_ = tmp6684

tmp6685 := MakeNative(func(__e *ControlFlow) {
V2975 := __e.Get(1)
_ = V2975
tmp6686 := Call(__e, PrimFunc(symshen_4_5lowC_6), V2975)


let__5306 := tmp6686
_ = let__5306

tmp6708 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5306)


var ifres6687 Obj

if True == tmp6708 {
tmp6688 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6687 = tmp6688


} else {
tmp6689 := Call(__e, PrimFunc(symshen_4in_1_6), let__5306)


let__5307 := tmp6689
_ = let__5307

tmp6690 := Call(__e, PrimFunc(symshen_4_5hash_6), let__5307)


let__5308 := tmp6690
_ = let__5308

tmp6707 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5308)


var ifres6691 Obj

if True == tmp6707 {
tmp6692 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6691 = tmp6692


} else {
tmp6693 := Call(__e, PrimFunc(symshen_4in_1_6), let__5308)


let__5309 := tmp6693
_ = let__5309

tmp6694 := Call(__e, PrimFunc(symshen_4_5integer_6), let__5309)


let__5310 := tmp6694
_ = let__5310

tmp6706 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5310)


var ifres6695 Obj

if True == tmp6706 {
tmp6696 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6695 = tmp6696


} else {
tmp6697 := Call(__e, PrimFunc(symshen_4_5_1out), let__5310)


let__5311 := tmp6697
_ = let__5311

tmp6698 := Call(__e, PrimFunc(symshen_4in_1_6), let__5310)


let__5312 := tmp6698
_ = let__5312

tmp6699 := Call(__e, PrimFunc(symshen_4_5semicolon_6), let__5312)


let__5313 := tmp6699
_ = let__5313

tmp6705 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5313)


var ifres6700 Obj

if True == tmp6705 {
tmp6701 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6700 = tmp6701


} else {
tmp6702 := Call(__e, PrimFunc(symshen_4in_1_6), let__5313)


let__5314 := tmp6702
_ = let__5314

tmp6703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(let__5311)
}
__typedArg0 := let__5311
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp6704 := Call(__e, PrimFunc(symshen_4comb), let__5314, tmp6703)


ifres6700 = tmp6704


}

ifres6695 = ifres6700


}

ifres6691 = ifres6695


}

ifres6687 = ifres6691


}

let__5305 := ifres6687
_ = let__5305

tmp6710 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5305)


if True == tmp6710 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5305)
return
}


}, 1)

tmp6711 := Call(__e, ns2_1set, symshen_4_5control_6, tmp6685)


_ = tmp6711

tmp6712 := MakeNative(func(__e *ControlFlow) {
V2986 := __e.Get(1)
_ = V2986
tmp6723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2986)
}
__typedArg0 := V2986
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6713 Obj

if True == tmp6723 {
tmp6714 := Call(__e, PrimFunc(symhead), V2986)


let__5316 := tmp6714
_ = let__5316

tmp6715 := Call(__e, PrimFunc(symtail), V2986)


let__5317 := tmp6715
_ = let__5317

tmp6720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__5316, MakeInteger(34))
}
__typedArg0 := let__5316
__typedArg1 := MakeInteger(34)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp6721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp6720)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp6720
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres6716 Obj

if True == tmp6721 {
tmp6717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(let__5316)
}
__typedArg0 := let__5316
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp6718 := Call(__e, PrimFunc(symshen_4comb), let__5317, tmp6717)


ifres6716 = tmp6718


} else {
tmp6719 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6716 = tmp6719


}

ifres6713 = ifres6716


} else {
tmp6722 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6713 = tmp6722


}

let__5315 := ifres6713
_ = let__5315

tmp6725 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5315)


if True == tmp6725 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5315)
return
}


}, 1)

tmp6726 := Call(__e, ns2_1set, symshen_4_5notdbq_6, tmp6712)


_ = tmp6726

tmp6727 := MakeNative(func(__e *ControlFlow) {
V2990 := __e.Get(1)
_ = V2990
tmp6732 := Call(__e, PrimFunc(symshen_4hds_a_2), V2990, MakeInteger(99))


var ifres6728 Obj

if True == tmp6732 {
tmp6729 := Call(__e, PrimFunc(symtail), V2990)


let__5319 := tmp6729
_ = let__5319

tmp6730 := Call(__e, PrimFunc(symshen_4comb), let__5319, symshen_4skip)


ifres6728 = tmp6730


} else {
tmp6731 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6728 = tmp6731


}

let__5318 := ifres6728
_ = let__5318

tmp6734 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5318)


if True == tmp6734 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5318)
return
}


}, 1)

tmp6735 := Call(__e, ns2_1set, symshen_4_5lowC_6, tmp6727)


_ = tmp6735

tmp6736 := MakeNative(func(__e *ControlFlow) {
V2993 := __e.Get(1)
_ = V2993
tmp6741 := Call(__e, PrimFunc(symshen_4hds_a_2), V2993, MakeInteger(35))


var ifres6737 Obj

if True == tmp6741 {
tmp6738 := Call(__e, PrimFunc(symtail), V2993)


let__5321 := tmp6738
_ = let__5321

tmp6739 := Call(__e, PrimFunc(symshen_4comb), let__5321, symshen_4skip)


ifres6737 = tmp6739


} else {
tmp6740 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6737 = tmp6740


}

let__5320 := ifres6737
_ = let__5320

tmp6743 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5320)


if True == tmp6743 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5320)
return
}


}, 1)

tmp6744 := Call(__e, ns2_1set, symshen_4_5hash_6, tmp6736)


_ = tmp6744

tmp6745 := MakeNative(func(__e *ControlFlow) {
V2996 := __e.Get(1)
_ = V2996
tmp6746 := Call(__e, PrimFunc(symshen_4_5minus_6), V2996)


let__5323 := tmp6746
_ = let__5323

tmp6758 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5323)


var ifres6747 Obj

if True == tmp6758 {
tmp6748 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6747 = tmp6748


} else {
tmp6749 := Call(__e, PrimFunc(symshen_4in_1_6), let__5323)


let__5324 := tmp6749
_ = let__5324

tmp6750 := Call(__e, PrimFunc(symshen_4_5number_6), let__5324)


let__5325 := tmp6750
_ = let__5325

tmp6757 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5325)


var ifres6751 Obj

if True == tmp6757 {
tmp6752 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6751 = tmp6752


} else {
tmp6753 := Call(__e, PrimFunc(symshen_4_5_1out), let__5325)


let__5326 := tmp6753
_ = let__5326

tmp6754 := Call(__e, PrimFunc(symshen_4in_1_6), let__5325)


let__5327 := tmp6754
_ = let__5327

tmp6756 := Call(__e, PrimFunc(symshen_4comb), let__5327, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(0))
__typedN1, __typedOK1 := TypedFloat64(let__5326)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := MakeInteger(0)
__typedArg1 := let__5326
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


ifres6751 = tmp6756


}

ifres6747 = ifres6751


}

let__5322 := ifres6747
_ = let__5322

tmp6801 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5322)


if True == tmp6801 {
tmp6759 := Call(__e, PrimFunc(symshen_4_5plus_6), V2996)


let__5329 := tmp6759
_ = let__5329

tmp6770 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5329)


var ifres6760 Obj

if True == tmp6770 {
tmp6761 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6760 = tmp6761


} else {
tmp6762 := Call(__e, PrimFunc(symshen_4in_1_6), let__5329)


let__5330 := tmp6762
_ = let__5330

tmp6763 := Call(__e, PrimFunc(symshen_4_5number_6), let__5330)


let__5331 := tmp6763
_ = let__5331

tmp6769 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5331)


var ifres6764 Obj

if True == tmp6769 {
tmp6765 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6764 = tmp6765


} else {
tmp6766 := Call(__e, PrimFunc(symshen_4_5_1out), let__5331)


let__5332 := tmp6766
_ = let__5332

tmp6767 := Call(__e, PrimFunc(symshen_4in_1_6), let__5331)


let__5333 := tmp6767
_ = let__5333

tmp6768 := Call(__e, PrimFunc(symshen_4comb), let__5333, let__5332)


ifres6764 = tmp6768


}

ifres6760 = ifres6764


}

let__5328 := ifres6760
_ = let__5328

tmp6799 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5328)


if True == tmp6799 {
tmp6771 := Call(__e, PrimFunc(symshen_4_5e_1number_6), V2996)


let__5335 := tmp6771
_ = let__5335

tmp6777 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5335)


var ifres6772 Obj

if True == tmp6777 {
tmp6773 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6772 = tmp6773


} else {
tmp6774 := Call(__e, PrimFunc(symshen_4_5_1out), let__5335)


let__5336 := tmp6774
_ = let__5336

tmp6775 := Call(__e, PrimFunc(symshen_4in_1_6), let__5335)


let__5337 := tmp6775
_ = let__5337

tmp6776 := Call(__e, PrimFunc(symshen_4comb), let__5337, let__5336)


ifres6772 = tmp6776


}

let__5334 := ifres6772
_ = let__5334

tmp6797 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5334)


if True == tmp6797 {
tmp6778 := Call(__e, PrimFunc(symshen_4_5float_6), V2996)


let__5339 := tmp6778
_ = let__5339

tmp6784 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5339)


var ifres6779 Obj

if True == tmp6784 {
tmp6780 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6779 = tmp6780


} else {
tmp6781 := Call(__e, PrimFunc(symshen_4_5_1out), let__5339)


let__5340 := tmp6781
_ = let__5340

tmp6782 := Call(__e, PrimFunc(symshen_4in_1_6), let__5339)


let__5341 := tmp6782
_ = let__5341

tmp6783 := Call(__e, PrimFunc(symshen_4comb), let__5341, let__5340)


ifres6779 = tmp6783


}

let__5338 := ifres6779
_ = let__5338

tmp6795 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5338)


if True == tmp6795 {
tmp6785 := Call(__e, PrimFunc(symshen_4_5integer_6), V2996)


let__5343 := tmp6785
_ = let__5343

tmp6791 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5343)


var ifres6786 Obj

if True == tmp6791 {
tmp6787 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6786 = tmp6787


} else {
tmp6788 := Call(__e, PrimFunc(symshen_4_5_1out), let__5343)


let__5344 := tmp6788
_ = let__5344

tmp6789 := Call(__e, PrimFunc(symshen_4in_1_6), let__5343)


let__5345 := tmp6789
_ = let__5345

tmp6790 := Call(__e, PrimFunc(symshen_4comb), let__5345, let__5344)


ifres6786 = tmp6790


}

let__5342 := ifres6786
_ = let__5342

tmp6793 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5342)


if True == tmp6793 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5342)
return
}


} else {
__e.Return(let__5338)
return
}


} else {
__e.Return(let__5334)
return
}


} else {
__e.Return(let__5328)
return
}


} else {
__e.Return(let__5322)
return
}


}, 1)

tmp6802 := Call(__e, ns2_1set, symshen_4_5number_6, tmp6745)


_ = tmp6802

tmp6803 := MakeNative(func(__e *ControlFlow) {
V3021 := __e.Get(1)
_ = V3021
tmp6808 := Call(__e, PrimFunc(symshen_4hds_a_2), V3021, MakeInteger(45))


var ifres6804 Obj

if True == tmp6808 {
tmp6805 := Call(__e, PrimFunc(symtail), V3021)


let__5347 := tmp6805
_ = let__5347

tmp6806 := Call(__e, PrimFunc(symshen_4comb), let__5347, symshen_4skip)


ifres6804 = tmp6806


} else {
tmp6807 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6804 = tmp6807


}

let__5346 := ifres6804
_ = let__5346

tmp6810 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5346)


if True == tmp6810 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5346)
return
}


}, 1)

tmp6811 := Call(__e, ns2_1set, symshen_4_5minus_6, tmp6803)


_ = tmp6811

tmp6812 := MakeNative(func(__e *ControlFlow) {
V3024 := __e.Get(1)
_ = V3024
tmp6817 := Call(__e, PrimFunc(symshen_4hds_a_2), V3024, MakeInteger(43))


var ifres6813 Obj

if True == tmp6817 {
tmp6814 := Call(__e, PrimFunc(symtail), V3024)


let__5349 := tmp6814
_ = let__5349

tmp6815 := Call(__e, PrimFunc(symshen_4comb), let__5349, symshen_4skip)


ifres6813 = tmp6815


} else {
tmp6816 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6813 = tmp6816


}

let__5348 := ifres6813
_ = let__5348

tmp6819 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5348)


if True == tmp6819 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5348)
return
}


}, 1)

tmp6820 := Call(__e, ns2_1set, symshen_4_5plus_6, tmp6812)


_ = tmp6820

tmp6821 := MakeNative(func(__e *ControlFlow) {
V3027 := __e.Get(1)
_ = V3027
tmp6822 := Call(__e, PrimFunc(symshen_4_5digits_6), V3027)


let__5351 := tmp6822
_ = let__5351

tmp6829 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5351)


var ifres6823 Obj

if True == tmp6829 {
tmp6824 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6823 = tmp6824


} else {
tmp6825 := Call(__e, PrimFunc(symshen_4_5_1out), let__5351)


let__5352 := tmp6825
_ = let__5352

tmp6826 := Call(__e, PrimFunc(symshen_4in_1_6), let__5351)


let__5353 := tmp6826
_ = let__5353

tmp6827 := Call(__e, PrimFunc(symshen_4compute_1integer), let__5352)


tmp6828 := Call(__e, PrimFunc(symshen_4comb), let__5353, tmp6827)


ifres6823 = tmp6828


}

let__5350 := ifres6823
_ = let__5350

tmp6831 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5350)


if True == tmp6831 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5350)
return
}


}, 1)

tmp6832 := Call(__e, ns2_1set, symshen_4_5integer_6, tmp6821)


_ = tmp6832

tmp6833 := MakeNative(func(__e *ControlFlow) {
V3032 := __e.Get(1)
_ = V3032
tmp6834 := Call(__e, PrimFunc(symshen_4_5digit_6), V3032)


let__5355 := tmp6834
_ = let__5355

tmp6847 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5355)


var ifres6835 Obj

if True == tmp6847 {
tmp6836 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6835 = tmp6836


} else {
tmp6837 := Call(__e, PrimFunc(symshen_4_5_1out), let__5355)


let__5356 := tmp6837
_ = let__5356

tmp6838 := Call(__e, PrimFunc(symshen_4in_1_6), let__5355)


let__5357 := tmp6838
_ = let__5357

tmp6839 := Call(__e, PrimFunc(symshen_4_5digits_6), let__5357)


let__5358 := tmp6839
_ = let__5358

tmp6846 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5358)


var ifres6840 Obj

if True == tmp6846 {
tmp6841 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6840 = tmp6841


} else {
tmp6842 := Call(__e, PrimFunc(symshen_4_5_1out), let__5358)


let__5359 := tmp6842
_ = let__5359

tmp6843 := Call(__e, PrimFunc(symshen_4in_1_6), let__5358)


let__5360 := tmp6843
_ = let__5360

tmp6844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__5356, let__5359)
}
__typedArg0 := let__5356
__typedArg1 := let__5359
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6845 := Call(__e, PrimFunc(symshen_4comb), let__5360, tmp6844)


ifres6840 = tmp6845


}

ifres6835 = ifres6840


}

let__5354 := ifres6835
_ = let__5354

tmp6859 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5354)


if True == tmp6859 {
tmp6848 := Call(__e, PrimFunc(symshen_4_5digit_6), V3032)


let__5362 := tmp6848
_ = let__5362

tmp6855 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5362)


var ifres6849 Obj

if True == tmp6855 {
tmp6850 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6849 = tmp6850


} else {
tmp6851 := Call(__e, PrimFunc(symshen_4_5_1out), let__5362)


let__5363 := tmp6851
_ = let__5363

tmp6852 := Call(__e, PrimFunc(symshen_4in_1_6), let__5362)


let__5364 := tmp6852
_ = let__5364

tmp6853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__5363, Nil)
}
__typedArg0 := let__5363
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6854 := Call(__e, PrimFunc(symshen_4comb), let__5364, tmp6853)


ifres6849 = tmp6854


}

let__5361 := ifres6849
_ = let__5361

tmp6857 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5361)


if True == tmp6857 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5361)
return
}


} else {
__e.Return(let__5354)
return
}


}, 1)

tmp6860 := Call(__e, ns2_1set, symshen_4_5digits_6, tmp6833)


_ = tmp6860

tmp6861 := MakeNative(func(__e *ControlFlow) {
V3044 := __e.Get(1)
_ = V3044
tmp6871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3044)
}
__typedArg0 := V3044
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6862 Obj

if True == tmp6871 {
tmp6863 := Call(__e, PrimFunc(symhead), V3044)


let__5366 := tmp6863
_ = let__5366

tmp6864 := Call(__e, PrimFunc(symtail), V3044)


let__5367 := tmp6864
_ = let__5367

tmp6869 := Call(__e, PrimFunc(symshen_4digit_2), let__5366)


var ifres6865 Obj

if True == tmp6869 {
tmp6866 := Call(__e, PrimFunc(symshen_4byte_1_6digit), let__5366)


tmp6867 := Call(__e, PrimFunc(symshen_4comb), let__5367, tmp6866)


ifres6865 = tmp6867


} else {
tmp6868 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6865 = tmp6868


}

ifres6862 = ifres6865


} else {
tmp6870 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6862 = tmp6870


}

let__5365 := ifres6862
_ = let__5365

tmp6873 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5365)


if True == tmp6873 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5365)
return
}


}, 1)

tmp6874 := Call(__e, ns2_1set, symshen_4_5digit_6, tmp6861)


_ = tmp6874

tmp6875 := MakeNative(func(__e *ControlFlow) {
V3048 := __e.Get(1)
_ = V3048
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V3048)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(48))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V3048
__typedArg1 := MakeInteger(48)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp6876 := Call(__e, ns2_1set, symshen_4byte_1_6digit, tmp6875)


_ = tmp6876

tmp6877 := MakeNative(func(__e *ControlFlow) {
V3049 := __e.Get(1)
_ = V3049
tmp6878 := Call(__e, PrimFunc(symreverse), V3049)


__e.TailApply(PrimFunc(symshen_4compute_1integer_1h), tmp6878, MakeInteger(0))
return


}, 1)

tmp6879 := Call(__e, ns2_1set, symshen_4compute_1integer, tmp6877)


_ = tmp6879

tmp6880 := MakeNative(func(__e *ControlFlow) {
V3052 := __e.Get(1)
_ = V3052
V3053 := __e.Get(2)
_ = V3053
tmp6890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3052)
}
__typedArg0 := Nil
__typedArg1 := V3052
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp6890 {
__e.Return(MakeInteger(0))
return
} else {
tmp6888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3052)
}
__typedArg0 := V3052
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp6888 {
tmp6881 := Call(__e, PrimFunc(symshen_4expt), MakeInteger(10), V3053)


tmp6882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3052)
}
__typedArg0 := V3052
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(tmp6881)
__typedN1, __typedOK1 := TypedFloat64(tmp6882)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := tmp6881
__typedArg1 := tmp6882
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})()

tmp6884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3052)
}
__typedArg0 := V3052
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6886 := Call(__e, PrimFunc(symshen_4compute_1integer_1h), tmp6884, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V3053)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V3053
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(tmp6883)
__typedN1, __typedOK1 := TypedFloat64(tmp6886)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := tmp6883
__typedArg1 := tmp6886
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4compute_1integer_1h)
return
}


}


}, 2)

tmp6891 := Call(__e, ns2_1set, symshen_4compute_1integer_1h, tmp6880)


_ = tmp6891

tmp6892 := MakeNative(func(__e *ControlFlow) {
V3056 := __e.Get(1)
_ = V3056
V3057 := __e.Get(2)
_ = V3057
tmp6900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V3057)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V3057
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp6900 {
__e.Return(MakeInteger(1))
return
} else {
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(V3057)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(0))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := V3057
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
tmp6894 := Call(__e, PrimFunc(symshen_4expt), V3056, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V3057)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V3057
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(V3056)
__typedN1, __typedOK1 := TypedFloat64(tmp6894)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := V3056
__typedArg1 := tmp6894
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})())
return


} else {
tmp6896 := Call(__e, PrimFunc(symshen_4expt), V3056, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V3057)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V3057
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_c) {
__typedN0, __typedOK0 := TypedFloat64(tmp6896)
__typedN1, __typedOK1 := TypedFloat64(V3056)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_c) {
return TypedMaterializeNumber(TypedDivideValue(__typedN0, __typedN1))
}}
__typedArg0 := tmp6896
__typedArg1 := V3056
return Call(__e, PrimFunc(sym_c), __typedArg0, __typedArg1)
})())
return


}


}


}, 2)

tmp6901 := Call(__e, ns2_1set, symshen_4expt, tmp6892)


_ = tmp6901

tmp6902 := MakeNative(func(__e *ControlFlow) {
V3058 := __e.Get(1)
_ = V3058
tmp6903 := Call(__e, PrimFunc(symshen_4_5integer_6), V3058)


let__5369 := tmp6903
_ = let__5369

tmp6921 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5369)


var ifres6904 Obj

if True == tmp6921 {
tmp6905 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6904 = tmp6905


} else {
tmp6906 := Call(__e, PrimFunc(symshen_4_5_1out), let__5369)


let__5370 := tmp6906
_ = let__5370

tmp6907 := Call(__e, PrimFunc(symshen_4in_1_6), let__5369)


let__5371 := tmp6907
_ = let__5371

tmp6908 := Call(__e, PrimFunc(symshen_4_5stop_6), let__5371)


let__5372 := tmp6908
_ = let__5372

tmp6920 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5372)


var ifres6909 Obj

if True == tmp6920 {
tmp6910 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6909 = tmp6910


} else {
tmp6911 := Call(__e, PrimFunc(symshen_4in_1_6), let__5372)


let__5373 := tmp6911
_ = let__5373

tmp6912 := Call(__e, PrimFunc(symshen_4_5fraction_6), let__5373)


let__5374 := tmp6912
_ = let__5374

tmp6919 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5374)


var ifres6913 Obj

if True == tmp6919 {
tmp6914 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6913 = tmp6914


} else {
tmp6915 := Call(__e, PrimFunc(symshen_4_5_1out), let__5374)


let__5375 := tmp6915
_ = let__5375

tmp6916 := Call(__e, PrimFunc(symshen_4in_1_6), let__5374)


let__5376 := tmp6916
_ = let__5376

tmp6918 := Call(__e, PrimFunc(symshen_4comb), let__5376, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(let__5370)
__typedN1, __typedOK1 := TypedFloat64(let__5375)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := let__5370
__typedArg1 := let__5375
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())


ifres6913 = tmp6918


}

ifres6909 = ifres6913


}

ifres6904 = ifres6909


}

let__5368 := ifres6904
_ = let__5368

tmp6937 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5368)


if True == tmp6937 {
tmp6922 := Call(__e, PrimFunc(symshen_4_5stop_6), V3058)


let__5378 := tmp6922
_ = let__5378

tmp6933 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5378)


var ifres6923 Obj

if True == tmp6933 {
tmp6924 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6923 = tmp6924


} else {
tmp6925 := Call(__e, PrimFunc(symshen_4in_1_6), let__5378)


let__5379 := tmp6925
_ = let__5379

tmp6926 := Call(__e, PrimFunc(symshen_4_5fraction_6), let__5379)


let__5380 := tmp6926
_ = let__5380

tmp6932 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5380)


var ifres6927 Obj

if True == tmp6932 {
tmp6928 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6927 = tmp6928


} else {
tmp6929 := Call(__e, PrimFunc(symshen_4_5_1out), let__5380)


let__5381 := tmp6929
_ = let__5381

tmp6930 := Call(__e, PrimFunc(symshen_4in_1_6), let__5380)


let__5382 := tmp6930
_ = let__5382

tmp6931 := Call(__e, PrimFunc(symshen_4comb), let__5382, let__5381)


ifres6927 = tmp6931


}

ifres6923 = ifres6927


}

let__5377 := ifres6923
_ = let__5377

tmp6935 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5377)


if True == tmp6935 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5377)
return
}


} else {
__e.Return(let__5368)
return
}


}, 1)

tmp6938 := Call(__e, ns2_1set, symshen_4_5float_6, tmp6902)


_ = tmp6938

tmp6939 := MakeNative(func(__e *ControlFlow) {
V3074 := __e.Get(1)
_ = V3074
tmp6944 := Call(__e, PrimFunc(symshen_4hds_a_2), V3074, MakeInteger(46))


var ifres6940 Obj

if True == tmp6944 {
tmp6941 := Call(__e, PrimFunc(symtail), V3074)


let__5384 := tmp6941
_ = let__5384

tmp6942 := Call(__e, PrimFunc(symshen_4comb), let__5384, symshen_4skip)


ifres6940 = tmp6942


} else {
tmp6943 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6940 = tmp6943


}

let__5383 := ifres6940
_ = let__5383

tmp6946 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5383)


if True == tmp6946 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5383)
return
}


}, 1)

tmp6947 := Call(__e, ns2_1set, symshen_4_5stop_6, tmp6939)


_ = tmp6947

tmp6948 := MakeNative(func(__e *ControlFlow) {
V3077 := __e.Get(1)
_ = V3077
tmp6949 := Call(__e, PrimFunc(symshen_4_5digits_6), V3077)


let__5386 := tmp6949
_ = let__5386

tmp6956 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5386)


var ifres6950 Obj

if True == tmp6956 {
tmp6951 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6950 = tmp6951


} else {
tmp6952 := Call(__e, PrimFunc(symshen_4_5_1out), let__5386)


let__5387 := tmp6952
_ = let__5387

tmp6953 := Call(__e, PrimFunc(symshen_4in_1_6), let__5386)


let__5388 := tmp6953
_ = let__5388

tmp6954 := Call(__e, PrimFunc(symshen_4compute_1fraction), let__5387)


tmp6955 := Call(__e, PrimFunc(symshen_4comb), let__5388, tmp6954)


ifres6950 = tmp6955


}

let__5385 := ifres6950
_ = let__5385

tmp6958 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5385)


if True == tmp6958 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5385)
return
}


}, 1)

tmp6959 := Call(__e, ns2_1set, symshen_4_5fraction_6, tmp6948)


_ = tmp6959

tmp6960 := MakeNative(func(__e *ControlFlow) {
V3082 := __e.Get(1)
_ = V3082
__e.TailApply(PrimFunc(symshen_4compute_1fraction_1h), V3082, MakeInteger(-1))
return
}, 1)

tmp6961 := Call(__e, ns2_1set, symshen_4compute_1fraction, tmp6960)


_ = tmp6961

tmp6962 := MakeNative(func(__e *ControlFlow) {
V3085 := __e.Get(1)
_ = V3085
V3086 := __e.Get(2)
_ = V3086
tmp6972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3085)
}
__typedArg0 := Nil
__typedArg1 := V3085
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp6972 {
__e.Return(MakeInteger(0))
return
} else {
tmp6970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3085)
}
__typedArg0 := V3085
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp6970 {
tmp6963 := Call(__e, PrimFunc(symshen_4expt), MakeInteger(10), V3086)


tmp6964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3085)
}
__typedArg0 := V3085
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(tmp6963)
__typedN1, __typedOK1 := TypedFloat64(tmp6964)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := tmp6963
__typedArg1 := tmp6964
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})()

tmp6966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3085)
}
__typedArg0 := V3085
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6968 := Call(__e, PrimFunc(symshen_4compute_1fraction_1h), tmp6966, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V3086)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V3086
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(tmp6965)
__typedN1, __typedOK1 := TypedFloat64(tmp6968)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := tmp6965
__typedArg1 := tmp6968
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4compute_1fraction_1h)
return
}


}


}, 2)

tmp6973 := Call(__e, ns2_1set, symshen_4compute_1fraction_1h, tmp6962)


_ = tmp6973

tmp6974 := MakeNative(func(__e *ControlFlow) {
V3087 := __e.Get(1)
_ = V3087
tmp6975 := Call(__e, PrimFunc(symshen_4_5float_6), V3087)


let__5390 := tmp6975
_ = let__5390

tmp6993 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5390)


var ifres6976 Obj

if True == tmp6993 {
tmp6977 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6976 = tmp6977


} else {
tmp6978 := Call(__e, PrimFunc(symshen_4_5_1out), let__5390)


let__5391 := tmp6978
_ = let__5391

tmp6979 := Call(__e, PrimFunc(symshen_4in_1_6), let__5390)


let__5392 := tmp6979
_ = let__5392

tmp6980 := Call(__e, PrimFunc(symshen_4_5lowE_6), let__5392)


let__5393 := tmp6980
_ = let__5393

tmp6992 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5393)


var ifres6981 Obj

if True == tmp6992 {
tmp6982 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6981 = tmp6982


} else {
tmp6983 := Call(__e, PrimFunc(symshen_4in_1_6), let__5393)


let__5394 := tmp6983
_ = let__5394

tmp6984 := Call(__e, PrimFunc(symshen_4_5log10_6), let__5394)


let__5395 := tmp6984
_ = let__5395

tmp6991 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5395)


var ifres6985 Obj

if True == tmp6991 {
tmp6986 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6985 = tmp6986


} else {
tmp6987 := Call(__e, PrimFunc(symshen_4_5_1out), let__5395)


let__5396 := tmp6987
_ = let__5396

tmp6988 := Call(__e, PrimFunc(symshen_4in_1_6), let__5395)


let__5397 := tmp6988
_ = let__5397

tmp6989 := Call(__e, PrimFunc(symshen_4compute_1E), let__5391, let__5396)


tmp6990 := Call(__e, PrimFunc(symshen_4comb), let__5397, tmp6989)


ifres6985 = tmp6990


}

ifres6981 = ifres6985


}

ifres6976 = ifres6981


}

let__5389 := ifres6976
_ = let__5389

tmp7016 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5389)


if True == tmp7016 {
tmp6994 := Call(__e, PrimFunc(symshen_4_5integer_6), V3087)


let__5399 := tmp6994
_ = let__5399

tmp7012 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5399)


var ifres6995 Obj

if True == tmp7012 {
tmp6996 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6995 = tmp6996


} else {
tmp6997 := Call(__e, PrimFunc(symshen_4_5_1out), let__5399)


let__5400 := tmp6997
_ = let__5400

tmp6998 := Call(__e, PrimFunc(symshen_4in_1_6), let__5399)


let__5401 := tmp6998
_ = let__5401

tmp6999 := Call(__e, PrimFunc(symshen_4_5lowE_6), let__5401)


let__5402 := tmp6999
_ = let__5402

tmp7011 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5402)


var ifres7000 Obj

if True == tmp7011 {
tmp7001 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres7000 = tmp7001


} else {
tmp7002 := Call(__e, PrimFunc(symshen_4in_1_6), let__5402)


let__5403 := tmp7002
_ = let__5403

tmp7003 := Call(__e, PrimFunc(symshen_4_5log10_6), let__5403)


let__5404 := tmp7003
_ = let__5404

tmp7010 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5404)


var ifres7004 Obj

if True == tmp7010 {
tmp7005 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres7004 = tmp7005


} else {
tmp7006 := Call(__e, PrimFunc(symshen_4_5_1out), let__5404)


let__5405 := tmp7006
_ = let__5405

tmp7007 := Call(__e, PrimFunc(symshen_4in_1_6), let__5404)


let__5406 := tmp7007
_ = let__5406

tmp7008 := Call(__e, PrimFunc(symshen_4compute_1E), let__5400, let__5405)


tmp7009 := Call(__e, PrimFunc(symshen_4comb), let__5406, tmp7008)


ifres7004 = tmp7009


}

ifres7000 = ifres7004


}

ifres6995 = ifres7000


}

let__5398 := ifres6995
_ = let__5398

tmp7014 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5398)


if True == tmp7014 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5398)
return
}


} else {
__e.Return(let__5389)
return
}


}, 1)

tmp7017 := Call(__e, ns2_1set, symshen_4_5e_1number_6, tmp6974)


_ = tmp7017

tmp7018 := MakeNative(func(__e *ControlFlow) {
V3106 := __e.Get(1)
_ = V3106
tmp7019 := Call(__e, PrimFunc(symshen_4_5plus_6), V3106)


let__5408 := tmp7019
_ = let__5408

tmp7030 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5408)


var ifres7020 Obj

if True == tmp7030 {
tmp7021 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres7020 = tmp7021


} else {
tmp7022 := Call(__e, PrimFunc(symshen_4in_1_6), let__5408)


let__5409 := tmp7022
_ = let__5409

tmp7023 := Call(__e, PrimFunc(symshen_4_5log10_6), let__5409)


let__5410 := tmp7023
_ = let__5410

tmp7029 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5410)


var ifres7024 Obj

if True == tmp7029 {
tmp7025 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres7024 = tmp7025


} else {
tmp7026 := Call(__e, PrimFunc(symshen_4_5_1out), let__5410)


let__5411 := tmp7026
_ = let__5411

tmp7027 := Call(__e, PrimFunc(symshen_4in_1_6), let__5410)


let__5412 := tmp7027
_ = let__5412

tmp7028 := Call(__e, PrimFunc(symshen_4comb), let__5412, let__5411)


ifres7024 = tmp7028


}

ifres7020 = ifres7024


}

let__5407 := ifres7020
_ = let__5407

tmp7056 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5407)


if True == tmp7056 {
tmp7031 := Call(__e, PrimFunc(symshen_4_5minus_6), V3106)


let__5414 := tmp7031
_ = let__5414

tmp7043 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5414)


var ifres7032 Obj

if True == tmp7043 {
tmp7033 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres7032 = tmp7033


} else {
tmp7034 := Call(__e, PrimFunc(symshen_4in_1_6), let__5414)


let__5415 := tmp7034
_ = let__5415

tmp7035 := Call(__e, PrimFunc(symshen_4_5log10_6), let__5415)


let__5416 := tmp7035
_ = let__5416

tmp7042 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5416)


var ifres7036 Obj

if True == tmp7042 {
tmp7037 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres7036 = tmp7037


} else {
tmp7038 := Call(__e, PrimFunc(symshen_4_5_1out), let__5416)


let__5417 := tmp7038
_ = let__5417

tmp7039 := Call(__e, PrimFunc(symshen_4in_1_6), let__5416)


let__5418 := tmp7039
_ = let__5418

tmp7041 := Call(__e, PrimFunc(symshen_4comb), let__5418, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(0))
__typedN1, __typedOK1 := TypedFloat64(let__5417)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := MakeInteger(0)
__typedArg1 := let__5417
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


ifres7036 = tmp7041


}

ifres7032 = ifres7036


}

let__5413 := ifres7032
_ = let__5413

tmp7054 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5413)


if True == tmp7054 {
tmp7044 := Call(__e, PrimFunc(symshen_4_5integer_6), V3106)


let__5420 := tmp7044
_ = let__5420

tmp7050 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5420)


var ifres7045 Obj

if True == tmp7050 {
tmp7046 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres7045 = tmp7046


} else {
tmp7047 := Call(__e, PrimFunc(symshen_4_5_1out), let__5420)


let__5421 := tmp7047
_ = let__5421

tmp7048 := Call(__e, PrimFunc(symshen_4in_1_6), let__5420)


let__5422 := tmp7048
_ = let__5422

tmp7049 := Call(__e, PrimFunc(symshen_4comb), let__5422, let__5421)


ifres7045 = tmp7049


}

let__5419 := ifres7045
_ = let__5419

tmp7052 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5419)


if True == tmp7052 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5419)
return
}


} else {
__e.Return(let__5413)
return
}


} else {
__e.Return(let__5407)
return
}


}, 1)

tmp7057 := Call(__e, ns2_1set, symshen_4_5log10_6, tmp7018)


_ = tmp7057

tmp7058 := MakeNative(func(__e *ControlFlow) {
V3123 := __e.Get(1)
_ = V3123
tmp7063 := Call(__e, PrimFunc(symshen_4hds_a_2), V3123, MakeInteger(101))


var ifres7059 Obj

if True == tmp7063 {
tmp7060 := Call(__e, PrimFunc(symtail), V3123)


let__5424 := tmp7060
_ = let__5424

tmp7061 := Call(__e, PrimFunc(symshen_4comb), let__5424, symshen_4skip)


ifres7059 = tmp7061


} else {
tmp7062 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres7059 = tmp7062


}

let__5423 := ifres7059
_ = let__5423

tmp7065 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5423)


if True == tmp7065 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5423)
return
}


}, 1)

tmp7066 := Call(__e, ns2_1set, symshen_4_5lowE_6, tmp7058)


_ = tmp7066

tmp7067 := MakeNative(func(__e *ControlFlow) {
V3126 := __e.Get(1)
_ = V3126
V3127 := __e.Get(2)
_ = V3127
tmp7068 := Call(__e, PrimFunc(symshen_4expt), MakeInteger(10), V3127)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(V3126)
__typedN1, __typedOK1 := TypedFloat64(tmp7068)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := V3126
__typedArg1 := tmp7068
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp7069 := Call(__e, ns2_1set, symshen_4compute_1E, tmp7067)


_ = tmp7069

tmp7070 := MakeNative(func(__e *ControlFlow) {
V3128 := __e.Get(1)
_ = V3128
tmp7071 := Call(__e, PrimFunc(symshen_4_5whitespace_6), V3128)


let__5426 := tmp7071
_ = let__5426

tmp7081 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5426)


var ifres7072 Obj

if True == tmp7081 {
tmp7073 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres7072 = tmp7073


} else {
tmp7074 := Call(__e, PrimFunc(symshen_4in_1_6), let__5426)


let__5427 := tmp7074
_ = let__5427

tmp7075 := Call(__e, PrimFunc(symshen_4_5whitespaces_6), let__5427)


let__5428 := tmp7075
_ = let__5428

tmp7080 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5428)


var ifres7076 Obj

if True == tmp7080 {
tmp7077 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres7076 = tmp7077


} else {
tmp7078 := Call(__e, PrimFunc(symshen_4in_1_6), let__5428)


let__5429 := tmp7078
_ = let__5429

tmp7079 := Call(__e, PrimFunc(symshen_4comb), let__5429, symshen_4skip)


ifres7076 = tmp7079


}

ifres7072 = ifres7076


}

let__5425 := ifres7072
_ = let__5425

tmp7091 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5425)


if True == tmp7091 {
tmp7082 := Call(__e, PrimFunc(symshen_4_5whitespace_6), V3128)


let__5431 := tmp7082
_ = let__5431

tmp7087 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5431)


var ifres7083 Obj

if True == tmp7087 {
tmp7084 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres7083 = tmp7084


} else {
tmp7085 := Call(__e, PrimFunc(symshen_4in_1_6), let__5431)


let__5432 := tmp7085
_ = let__5432

tmp7086 := Call(__e, PrimFunc(symshen_4comb), let__5432, symshen_4skip)


ifres7083 = tmp7086


}

let__5430 := ifres7083
_ = let__5430

tmp7089 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5430)


if True == tmp7089 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5430)
return
}


} else {
__e.Return(let__5425)
return
}


}, 1)

tmp7092 := Call(__e, ns2_1set, symshen_4_5whitespaces_6, tmp7070)


_ = tmp7092

tmp7093 := MakeNative(func(__e *ControlFlow) {
V3137 := __e.Get(1)
_ = V3137
tmp7102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3137)
}
__typedArg0 := V3137
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7094 Obj

if True == tmp7102 {
tmp7095 := Call(__e, PrimFunc(symhead), V3137)


let__5434 := tmp7095
_ = let__5434

tmp7096 := Call(__e, PrimFunc(symtail), V3137)


let__5435 := tmp7096
_ = let__5435

tmp7100 := Call(__e, PrimFunc(symshen_4whitespace_2), let__5434)


var ifres7097 Obj

if True == tmp7100 {
tmp7098 := Call(__e, PrimFunc(symshen_4comb), let__5435, symshen_4skip)


ifres7097 = tmp7098


} else {
tmp7099 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres7097 = tmp7099


}

ifres7094 = ifres7097


} else {
tmp7101 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres7094 = tmp7101


}

let__5433 := ifres7094
_ = let__5433

tmp7104 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5433)


if True == tmp7104 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5433)
return
}


}, 1)

tmp7105 := Call(__e, ns2_1set, symshen_4_5whitespace_6, tmp7093)


_ = tmp7105

tmp7106 := MakeNative(func(__e *ControlFlow) {
V3143 := __e.Get(1)
_ = V3143
tmp7114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(32), V3143)
}
__typedArg0 := MakeInteger(32)
__typedArg1 := V3143
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7114 {
__e.Return(True)
return
} else {
tmp7112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(13), V3143)
}
__typedArg0 := MakeInteger(13)
__typedArg1 := V3143
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7112 {
__e.Return(True)
return
} else {
tmp7110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(10), V3143)
}
__typedArg0 := MakeInteger(10)
__typedArg1 := V3143
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7110 {
__e.Return(True)
return
} else {
tmp7108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(9), V3143)
}
__typedArg0 := MakeInteger(9)
__typedArg1 := V3143
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7108 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}


}


}, 1)

tmp7115 := Call(__e, ns2_1set, symshen_4whitespace_2, tmp7106)


_ = tmp7115

tmp7116 := MakeNative(func(__e *ControlFlow) {
V3146 := __e.Get(1)
_ = V3146
V3147 := __e.Get(2)
_ = V3147
tmp7131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V3146)
}
__typedArg0 := V3146
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

var ifres7123 Obj

if True == tmp7131 {
tmp7129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V3146)
}
__typedArg0 := V3146
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

tmp7130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7129)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7129
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres7125 Obj

if True == tmp7130 {
tmp7127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3146, sympackage)
}
__typedArg0 := V3146
__typedArg1 := sympackage
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp7128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7127)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7127
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres7126 Obj

if True == tmp7128 {
ifres7126 = True


} else {
ifres7126 = False


}

ifres7125 = ifres7126


} else {
ifres7125 = False


}

var ifres7124 Obj

if True == ifres7125 {
ifres7124 = True


} else {
ifres7124 = False


}

ifres7123 = ifres7124


} else {
ifres7123 = False


}

if True == ifres7123 {
tmp7117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsource_1form_1handlers_d)
}
__typedArg0 := symshen_4_dsource_1form_1handlers_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp7118 := Call(__e, PrimFunc(symshen_4update_1assoc), V3146, V3147, tmp7117)


tmp7119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsource_1form_1handlers_d, tmp7118)
}
__typedArg0 := symshen_4_dsource_1form_1handlers_d
__typedArg1 := tmp7118
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7119

__e.Return(V3146)
return


} else {
tmp7120 := Call(__e, PrimFunc(symshen_4app), V3146, MakeString("\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("invalid source-form head: "))
__typedS1, __typedOK1 := TypedString(tmp7120)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("invalid source-form head: ")
__typedArg1 := tmp7120
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("invalid source-form head: "))
__typedS1, __typedOK1 := TypedString(tmp7120)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("invalid source-form head: ")
__typedArg1 := tmp7120
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 2)

tmp7132 := Call(__e, ns2_1set, symshen_4register_1source_1form, tmp7116)


_ = tmp7132

tmp7133 := MakeNative(func(__e *ControlFlow) {
V3148 := __e.Get(1)
_ = V3148
tmp7134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsource_1form_1handlers_d)
}
__typedArg0 := symshen_4_dsource_1form_1handlers_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp7135 := Call(__e, PrimFunc(symassoc), V3148, tmp7134)


tmp7136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsource_1form_1handlers_d)
}
__typedArg0 := symshen_4_dsource_1form_1handlers_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp7137 := Call(__e, PrimFunc(symremove), tmp7135, tmp7136)


tmp7138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsource_1form_1handlers_d, tmp7137)
}
__typedArg0 := symshen_4_dsource_1form_1handlers_d
__typedArg1 := tmp7137
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7138

__e.Return(V3148)
return


}, 1)

tmp7139 := Call(__e, ns2_1set, symshen_4unregister_1source_1form, tmp7133)


_ = tmp7139

tmp7140 := MakeNative(func(__e *ControlFlow) {
V3155 := __e.Get(1)
_ = V3155
V3156 := __e.Get(2)
_ = V3156
tmp7143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3155)
}
__typedArg0 := V3155
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7143 {
tmp7141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3155)
}
__typedArg0 := V3155
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symassoc), tmp7141, V3156)
return


} else {
__e.Return(Nil)
return
}


}, 2)

tmp7144 := Call(__e, ns2_1set, symshen_4source_1form_1handler, tmp7140)


_ = tmp7144

tmp7145 := MakeNative(func(__e *ControlFlow) {
V3159 := __e.Get(1)
_ = V3159
V3160 := __e.Get(2)
_ = V3160
tmp7153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3160)
}
__typedArg0 := V3160
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7153 {
tmp7146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3160)
}
__typedArg0 := V3160
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7147 := Call(__e, tmp7146, V3159)


let__5436 := tmp7147
_ = let__5436

tmp7151 := Call(__e, PrimFunc(symshen_4source_1form_1list_2), let__5436)


if True == tmp7151 {
__e.Return(let__5436)
return
} else {
tmp7148 := Call(__e, PrimFunc(symshen_4app), V3159, MakeString("\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("source-form handler must return a list of forms: "))
__typedS1, __typedOK1 := TypedString(tmp7148)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("source-form handler must return a list of forms: ")
__typedArg1 := tmp7148
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("source-form handler must return a list of forms: "))
__typedS1, __typedOK1 := TypedString(tmp7148)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("source-form handler must return a list of forms: ")
__typedArg1 := tmp7148
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4source_1form_1expansion)
return
}


}, 2)

tmp7154 := Call(__e, ns2_1set, symshen_4source_1form_1expansion, tmp7145)


_ = tmp7154

tmp7155 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V3166 := __self1
_ = V3166
tmp7160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3166)
}
__typedArg0 := Nil
__typedArg1 := V3166
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7160 {
__e.Return(True)
return
} else {
tmp7158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3166)
}
__typedArg0 := V3166
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7158 {
tmp7156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3166)
}
__typedArg0 := V3166
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4source_1form_1list_2) == __self {
__self1 = tmp7156
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4source_1form_1list_2), tmp7156)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp7161 := Call(__e, ns2_1set, symshen_4source_1form_1list_2, tmp7155)


_ = tmp7161

tmp7162 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V3167 := __self1
_ = V3167
tmp7198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3167)
}
__typedArg0 := Nil
__typedArg1 := V3167
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7198 {
__e.Return(Nil)
return
} else {
tmp7196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3167)
}
__typedArg0 := V3167
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7192 Obj

if True == tmp7196 {
tmp7194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3167)
}
__typedArg0 := V3167
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7195 := Call(__e, PrimFunc(symshen_4packaged_2), tmp7194)


var ifres7193 Obj

if True == tmp7195 {
ifres7193 = True


} else {
ifres7193 = False


}

ifres7192 = ifres7193


} else {
ifres7192 = False


}

if True == ifres7192 {
tmp7163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3167)
}
__typedArg0 := V3167
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7164 := Call(__e, PrimFunc(symshen_4unpackage), tmp7163)


tmp7165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3167)
}
__typedArg0 := V3167
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7166 := Call(__e, PrimFunc(symappend), tmp7164, tmp7165)


if PrimFunc(symshen_4unpackage_emacroexpand) == __self {
__self1 = tmp7166
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4unpackage_emacroexpand), tmp7166)
return


} else {
tmp7190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3167)
}
__typedArg0 := V3167
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7190 {
tmp7167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3167)
}
__typedArg0 := V3167
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsource_1form_1handlers_d)
}
__typedArg0 := symshen_4_dsource_1form_1handlers_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp7169 := Call(__e, PrimFunc(symshen_4source_1form_1handler), tmp7167, tmp7168)


let__5437 := tmp7169
_ = let__5437

tmp7188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__5437)
}
__typedArg0 := let__5437
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7188 {
tmp7170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3167)
}
__typedArg0 := V3167
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7171 := Call(__e, PrimFunc(symshen_4source_1form_1expansion), tmp7170, let__5437)


tmp7172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3167)
}
__typedArg0 := V3167
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7173 := Call(__e, PrimFunc(symappend), tmp7171, tmp7172)


if PrimFunc(symshen_4unpackage_emacroexpand) == __self {
__self1 = tmp7173
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4unpackage_emacroexpand), tmp7173)
return


} else {
tmp7174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3167)
}
__typedArg0 := V3167
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7175 := Call(__e, PrimFunc(symmacroexpand), tmp7174)


let__5438 := tmp7175
_ = let__5438

tmp7186 := Call(__e, PrimFunc(symshen_4packaged_2), let__5438)


var ifres7181 Obj

if True == tmp7186 {
ifres7181 = True


} else {
tmp7183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsource_1form_1handlers_d)
}
__typedArg0 := symshen_4_dsource_1form_1handlers_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp7184 := Call(__e, PrimFunc(symshen_4source_1form_1handler), let__5438, tmp7183)


tmp7185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7184)
}
__typedArg0 := tmp7184
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7182 Obj

if True == tmp7185 {
ifres7182 = True


} else {
ifres7182 = False


}

ifres7181 = ifres7182


}

if True == ifres7181 {
tmp7176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3167)
}
__typedArg0 := V3167
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__5438, tmp7176)
}
__typedArg0 := let__5438
__typedArg1 := tmp7176
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4unpackage_emacroexpand) == __self {
__self1 = tmp7177
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4unpackage_emacroexpand), tmp7177)
return


} else {
tmp7178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3167)
}
__typedArg0 := V3167
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7179 := Call(__e, PrimFunc(symshen_4unpackage_emacroexpand), tmp7178)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__5438, tmp7179)
}
__typedArg0 := let__5438
__typedArg1 := tmp7179
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4unpackage_emacroexpand)
return
}


}


}


}, 1)

tmp7199 := Call(__e, ns2_1set, symshen_4unpackage_emacroexpand, tmp7162)


_ = tmp7199

tmp7200 := MakeNative(func(__e *ControlFlow) {
V3172 := __e.Get(1)
_ = V3172
tmp7215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3172)
}
__typedArg0 := V3172
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7202 Obj

if True == tmp7215 {
tmp7213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3172)
}
__typedArg0 := V3172
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sympackage, tmp7213)
}
__typedArg0 := sympackage
__typedArg1 := tmp7213
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7204 Obj

if True == tmp7214 {
tmp7211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3172)
}
__typedArg0 := V3172
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7211)
}
__typedArg0 := tmp7211
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7206 Obj

if True == tmp7212 {
tmp7208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3172)
}
__typedArg0 := V3172
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7208)
}
__typedArg0 := tmp7208
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7209)
}
__typedArg0 := tmp7209
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7207 Obj

if True == tmp7210 {
ifres7207 = True


} else {
ifres7207 = False


}

ifres7206 = ifres7207


} else {
ifres7206 = False


}

var ifres7205 Obj

if True == ifres7206 {
ifres7205 = True


} else {
ifres7205 = False


}

ifres7204 = ifres7205


} else {
ifres7204 = False


}

var ifres7203 Obj

if True == ifres7204 {
ifres7203 = True


} else {
ifres7203 = False


}

ifres7202 = ifres7203


} else {
ifres7202 = False


}

if True == ifres7202 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp7216 := Call(__e, ns2_1set, symshen_4packaged_2, tmp7200)


_ = tmp7216

tmp7217 := MakeNative(func(__e *ControlFlow) {
V3175 := __e.Get(1)
_ = V3175
tmp7274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3175)
}
__typedArg0 := V3175
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7256 Obj

if True == tmp7274 {
tmp7272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3175)
}
__typedArg0 := V3175
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sympackage, tmp7272)
}
__typedArg0 := sympackage
__typedArg1 := tmp7272
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7258 Obj

if True == tmp7273 {
tmp7270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3175)
}
__typedArg0 := V3175
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7270)
}
__typedArg0 := tmp7270
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7260 Obj

if True == tmp7271 {
tmp7267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3175)
}
__typedArg0 := V3175
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7267)
}
__typedArg0 := tmp7267
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnull, tmp7268)
}
__typedArg0 := symnull
__typedArg1 := tmp7268
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7262 Obj

if True == tmp7269 {
tmp7264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3175)
}
__typedArg0 := V3175
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7264)
}
__typedArg0 := tmp7264
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7265)
}
__typedArg0 := tmp7265
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7263 Obj

if True == tmp7266 {
ifres7263 = True


} else {
ifres7263 = False


}

ifres7262 = ifres7263


} else {
ifres7262 = False


}

var ifres7261 Obj

if True == ifres7262 {
ifres7261 = True


} else {
ifres7261 = False


}

ifres7260 = ifres7261


} else {
ifres7260 = False


}

var ifres7259 Obj

if True == ifres7260 {
ifres7259 = True


} else {
ifres7259 = False


}

ifres7258 = ifres7259


} else {
ifres7258 = False


}

var ifres7257 Obj

if True == ifres7258 {
ifres7257 = True


} else {
ifres7257 = False


}

ifres7256 = ifres7257


} else {
ifres7256 = False


}

if True == ifres7256 {
tmp7218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3175)
}
__typedArg0 := V3175
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7218)
}
__typedArg0 := tmp7218
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7219)
}
__typedArg0 := tmp7219
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return


} else {
tmp7254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3175)
}
__typedArg0 := V3175
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7241 Obj

if True == tmp7254 {
tmp7252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3175)
}
__typedArg0 := V3175
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sympackage, tmp7252)
}
__typedArg0 := sympackage
__typedArg1 := tmp7252
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7243 Obj

if True == tmp7253 {
tmp7250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3175)
}
__typedArg0 := V3175
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7250)
}
__typedArg0 := tmp7250
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7245 Obj

if True == tmp7251 {
tmp7247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3175)
}
__typedArg0 := V3175
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7247)
}
__typedArg0 := tmp7247
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7248)
}
__typedArg0 := tmp7248
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7246 Obj

if True == tmp7249 {
ifres7246 = True


} else {
ifres7246 = False


}

ifres7245 = ifres7246


} else {
ifres7245 = False


}

var ifres7244 Obj

if True == ifres7245 {
ifres7244 = True


} else {
ifres7244 = False


}

ifres7243 = ifres7244


} else {
ifres7243 = False


}

var ifres7242 Obj

if True == ifres7243 {
ifres7242 = True


} else {
ifres7242 = False


}

ifres7241 = ifres7242


} else {
ifres7241 = False


}

if True == ifres7241 {
tmp7220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3175)
}
__typedArg0 := V3175
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7220)
}
__typedArg0 := tmp7220
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7221)
}
__typedArg0 := tmp7221
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7223 := Call(__e, PrimFunc(symeval), tmp7222)


let__5439 := tmp7223
_ = let__5439

tmp7224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3175)
}
__typedArg0 := V3175
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7224)
}
__typedArg0 := tmp7224
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(tmp7225)
}
__typedArg0 := tmp7225
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp7227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3175)
}
__typedArg0 := V3175
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7227)
}
__typedArg0 := tmp7227
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7228)
}
__typedArg0 := tmp7228
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7230 := Call(__e, PrimFunc(symshen_4package_1symbols), tmp7226, let__5439, tmp7229)


let__5440 := tmp7230
_ = let__5440

tmp7231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3175)
}
__typedArg0 := V3175
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7231)
}
__typedArg0 := tmp7231
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7233 := Call(__e, PrimFunc(symshen_4record_1external), tmp7232, let__5439)


let__5441 := tmp7233
_ = let__5441

tmp7234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3175)
}
__typedArg0 := V3175
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7234)
}
__typedArg0 := tmp7234
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3175)
}
__typedArg0 := V3175
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7236)
}
__typedArg0 := tmp7236
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7237)
}
__typedArg0 := tmp7237
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7239 := Call(__e, PrimFunc(symshen_4record_1internal), tmp7235, let__5439, tmp7238)


let__5442 := tmp7239
_ = let__5442

__e.Return(let__5440)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4unpackage)
return
}


}


}, 1)

tmp7275 := Call(__e, ns2_1set, symshen_4unpackage, tmp7217)


_ = tmp7275

tmp7276 := MakeNative(func(__e *ControlFlow) {
V3180 := __e.Get(1)
_ = V3180
V3181 := __e.Get(2)
_ = V3181
V3182 := __e.Get(3)
_ = V3182
tmp7277 := MakeNative(func(__e *ControlFlow) {
tmp7278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V3180, symshen_4internal_1symbols, tmp7278)
return


}, 0)

tmp7279 := MakeNative(func(__e *ControlFlow) {
Z3184 := __e.Get(1)
_ = Z3184
__e.Return(Nil)
return
}, 1)

tmp7280 := Call(__e, try_1catch, tmp7277, tmp7279)


let__5443 := tmp7280
_ = let__5443

tmp7281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3180)
}
__typedArg0 := V3180
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp7282 := Call(__e, PrimFunc(symshen_4internal_1symbols), tmp7281, V3181, V3182)


let__5444 := tmp7282
_ = let__5444

tmp7283 := Call(__e, PrimFunc(symunion), let__5444, let__5443)


tmp7284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V3180, symshen_4internal_1symbols, tmp7283, tmp7284)
return


}, 3)

tmp7285 := Call(__e, ns2_1set, symshen_4record_1internal, tmp7276)


_ = tmp7285

tmp7286 := MakeNative(func(__e *ControlFlow) {
V3192 := __e.Get(1)
_ = V3192
V3193 := __e.Get(2)
_ = V3193
V3194 := __e.Get(3)
_ = V3194
tmp7295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3194)
}
__typedArg0 := V3194
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7295 {
tmp7287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3194)
}
__typedArg0 := V3194
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7288 := Call(__e, PrimFunc(symshen_4internal_1symbols), V3192, V3193, tmp7287)


tmp7289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3194)
}
__typedArg0 := V3194
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7290 := Call(__e, PrimFunc(symshen_4internal_1symbols), V3192, V3193, tmp7289)


__e.TailApply(PrimFunc(symunion), tmp7288, tmp7290)
return


} else {
tmp7293 := Call(__e, PrimFunc(symshen_4internal_2), V3194, V3192, V3193)


if True == tmp7293 {
tmp7291 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V3192, V3194)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7291, Nil)
}
__typedArg0 := tmp7291
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(Nil)
return
}


}


}, 3)

tmp7296 := Call(__e, ns2_1set, symshen_4internal_1symbols, tmp7286)


_ = tmp7296

tmp7297 := MakeNative(func(__e *ControlFlow) {
V3195 := __e.Get(1)
_ = V3195
V3196 := __e.Get(2)
_ = V3196
tmp7298 := MakeNative(func(__e *ControlFlow) {
tmp7299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V3195, symshen_4external_1symbols, tmp7299)
return


}, 0)

tmp7300 := MakeNative(func(__e *ControlFlow) {
Z3198 := __e.Get(1)
_ = Z3198
__e.Return(Nil)
return
}, 1)

tmp7301 := Call(__e, try_1catch, tmp7298, tmp7300)


let__5445 := tmp7301
_ = let__5445

tmp7302 := Call(__e, PrimFunc(symunion), V3196, let__5445)


tmp7303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V3195, symshen_4external_1symbols, tmp7302, tmp7303)
return


}, 2)

tmp7304 := Call(__e, ns2_1set, symshen_4record_1external, tmp7297)


_ = tmp7304

tmp7305 := MakeNative(func(__e *ControlFlow) {
V3203 := __e.Get(1)
_ = V3203
V3204 := __e.Get(2)
_ = V3204
V3205 := __e.Get(3)
_ = V3205
tmp7310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3205)
}
__typedArg0 := V3205
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7310 {
tmp7306 := MakeNative(func(__e *ControlFlow) {
Z3206 := __e.Get(1)
_ = Z3206
__e.TailApply(PrimFunc(symshen_4package_1symbols), V3203, V3204, Z3206)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp7306, V3205)
return


} else {
tmp7308 := Call(__e, PrimFunc(symshen_4internal_2), V3205, V3203, V3204)


if True == tmp7308 {
__e.TailApply(PrimFunc(symshen_4intern_1in_1package), V3203, V3205)
return
} else {
__e.Return(V3205)
return
}


}


}, 3)

tmp7311 := Call(__e, ns2_1set, symshen_4package_1symbols, tmp7305)


_ = tmp7311

tmp7312 := MakeNative(func(__e *ControlFlow) {
V3207 := __e.Get(1)
_ = V3207
V3208 := __e.Get(2)
_ = V3208
tmp7313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3208)
}
__typedArg0 := V3208
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp7314 := Call(__e, PrimFunc(sym_8s), MakeString("."), tmp7313)


tmp7315 := Call(__e, PrimFunc(sym_8s), V3207, tmp7314)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(tmp7315)
}
__typedArg0 := tmp7315
return Call(__e, PrimFunc(symintern), __typedArg0)
})())
return


}, 2)

tmp7316 := Call(__e, ns2_1set, symshen_4intern_1in_1package, tmp7312)


_ = tmp7316

tmp7317 := MakeNative(func(__e *ControlFlow) {
V3209 := __e.Get(1)
_ = V3209
V3210 := __e.Get(2)
_ = V3210
V3211 := __e.Get(3)
_ = V3211
tmp7347 := Call(__e, PrimFunc(symelement_2), V3209, V3211)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7347)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7347
return Call(__e, PrimFunc(symnot), __typedArg0)
})() {
tmp7344 := Call(__e, PrimFunc(symshen_4sng_2), V3209)


tmp7345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7344)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7344
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres7319 Obj

if True == tmp7345 {
tmp7342 := Call(__e, PrimFunc(symshen_4dbl_2), V3209)


tmp7343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7342)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7342
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres7321 Obj

if True == tmp7343 {
tmp7341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V3209)
}
__typedArg0 := V3209
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

var ifres7323 Obj

if True == tmp7341 {
tmp7339 := Call(__e, PrimFunc(symshen_4sysfunc_2), V3209)


tmp7340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7339)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7339
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres7325 Obj

if True == tmp7340 {
tmp7337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V3209)
}
__typedArg0 := V3209
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

tmp7338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7337)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7337
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres7327 Obj

if True == tmp7338 {
tmp7334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3209)
}
__typedArg0 := V3209
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp7335 := Call(__e, PrimFunc(symshen_4internal_1to_1shen_2), tmp7334)


tmp7336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7335)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7335
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres7329 Obj

if True == tmp7336 {
tmp7331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3209)
}
__typedArg0 := V3209
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp7332 := Call(__e, PrimFunc(symshen_4internal_1to_1P_2), V3210, tmp7331)


tmp7333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7332)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7332
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres7330 Obj

if True == tmp7333 {
ifres7330 = True


} else {
ifres7330 = False


}

ifres7329 = ifres7330


} else {
ifres7329 = False


}

var ifres7328 Obj

if True == ifres7329 {
ifres7328 = True


} else {
ifres7328 = False


}

ifres7327 = ifres7328


} else {
ifres7327 = False


}

var ifres7326 Obj

if True == ifres7327 {
ifres7326 = True


} else {
ifres7326 = False


}

ifres7325 = ifres7326


} else {
ifres7325 = False


}

var ifres7324 Obj

if True == ifres7325 {
ifres7324 = True


} else {
ifres7324 = False


}

ifres7323 = ifres7324


} else {
ifres7323 = False


}

var ifres7322 Obj

if True == ifres7323 {
ifres7322 = True


} else {
ifres7322 = False


}

ifres7321 = ifres7322


} else {
ifres7321 = False


}

var ifres7320 Obj

if True == ifres7321 {
ifres7320 = True


} else {
ifres7320 = False


}

ifres7319 = ifres7320


} else {
ifres7319 = False


}

if True == ifres7319 {
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


}, 3)

tmp7349 := Call(__e, ns2_1set, symshen_4internal_2, tmp7317)


_ = tmp7349

tmp7350 := MakeNative(func(__e *ControlFlow) {
V3216 := __e.Get(1)
_ = V3216
tmp7404 := Call(__e, PrimFunc(symshen_4_7string_2), V3216)


var ifres7352 Obj

if True == tmp7404 {
tmp7402 := Call(__e, PrimFunc(symhdstr), V3216)


tmp7403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("s"), tmp7402)
}
__typedArg0 := MakeString("s")
__typedArg1 := tmp7402
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7354 Obj

if True == tmp7403 {
tmp7401 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3216
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres7356 Obj

if True == tmp7401 {
tmp7398 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3216
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp7399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("h"), tmp7398)
}
__typedArg0 := MakeString("h")
__typedArg1 := tmp7398
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7358 Obj

if True == tmp7399 {
tmp7396 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3216
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres7360 Obj

if True == tmp7396 {
tmp7392 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3216
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp7393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("e"), tmp7392)
}
__typedArg0 := MakeString("e")
__typedArg1 := tmp7392
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7362 Obj

if True == tmp7393 {
tmp7389 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
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
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3216
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres7364 Obj

if True == tmp7389 {
tmp7384 := Call(__e, PrimFunc(symhdstr), (func() Obj {
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
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3216
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp7385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("n"), tmp7384)
}
__typedArg0 := MakeString("n")
__typedArg1 := tmp7384
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7366 Obj

if True == tmp7385 {
tmp7380 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
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
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3216
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres7368 Obj

if True == tmp7380 {
tmp7374 := Call(__e, PrimFunc(symhdstr), (func() Obj {
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
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3216)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3216
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp7375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("."), tmp7374)
}
__typedArg0 := MakeString(".")
__typedArg1 := tmp7374
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7369 Obj

if True == tmp7375 {
ifres7369 = True


} else {
ifres7369 = False


}

ifres7368 = ifres7369


} else {
ifres7368 = False


}

var ifres7367 Obj

if True == ifres7368 {
ifres7367 = True


} else {
ifres7367 = False


}

ifres7366 = ifres7367


} else {
ifres7366 = False


}

var ifres7365 Obj

if True == ifres7366 {
ifres7365 = True


} else {
ifres7365 = False


}

ifres7364 = ifres7365


} else {
ifres7364 = False


}

var ifres7363 Obj

if True == ifres7364 {
ifres7363 = True


} else {
ifres7363 = False


}

ifres7362 = ifres7363


} else {
ifres7362 = False


}

var ifres7361 Obj

if True == ifres7362 {
ifres7361 = True


} else {
ifres7361 = False


}

ifres7360 = ifres7361


} else {
ifres7360 = False


}

var ifres7359 Obj

if True == ifres7360 {
ifres7359 = True


} else {
ifres7359 = False


}

ifres7358 = ifres7359


} else {
ifres7358 = False


}

var ifres7357 Obj

if True == ifres7358 {
ifres7357 = True


} else {
ifres7357 = False


}

ifres7356 = ifres7357


} else {
ifres7356 = False


}

var ifres7355 Obj

if True == ifres7356 {
ifres7355 = True


} else {
ifres7355 = False


}

ifres7354 = ifres7355


} else {
ifres7354 = False


}

var ifres7353 Obj

if True == ifres7354 {
ifres7353 = True


} else {
ifres7353 = False


}

ifres7352 = ifres7353


} else {
ifres7352 = False


}

if True == ifres7352 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp7405 := Call(__e, ns2_1set, symshen_4internal_1to_1shen_2, tmp7350)


_ = tmp7405

tmp7406 := MakeNative(func(__e *ControlFlow) {
V3217 := __e.Get(1)
_ = V3217
tmp7407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp7408 := Call(__e, PrimFunc(symget), symshen, symshen_4external_1symbols, tmp7407)


__e.TailApply(PrimFunc(symelement_2), V3217, tmp7408)
return


}, 1)

tmp7409 := Call(__e, ns2_1set, symshen_4sysfunc_2, tmp7406)


_ = tmp7409

tmp7410 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V3225 := __self1
_ = V3225
V3226 := __self2
_ = V3226
tmp7431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V3225)
}
__typedArg0 := MakeString("")
__typedArg1 := V3225
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7424 Obj

if True == tmp7431 {
tmp7430 := Call(__e, PrimFunc(symshen_4_7string_2), V3226)


var ifres7426 Obj

if True == tmp7430 {
tmp7428 := Call(__e, PrimFunc(symhdstr), V3226)


tmp7429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("."), tmp7428)
}
__typedArg0 := MakeString(".")
__typedArg1 := tmp7428
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7427 Obj

if True == tmp7429 {
ifres7427 = True


} else {
ifres7427 = False


}

ifres7426 = ifres7427


} else {
ifres7426 = False


}

var ifres7425 Obj

if True == ifres7426 {
ifres7425 = True


} else {
ifres7425 = False


}

ifres7424 = ifres7425


} else {
ifres7424 = False


}

if True == ifres7424 {
__e.Return(True)
return
} else {
tmp7422 := Call(__e, PrimFunc(symshen_4_7string_2), V3225)


var ifres7414 Obj

if True == tmp7422 {
tmp7421 := Call(__e, PrimFunc(symshen_4_7string_2), V3226)


var ifres7416 Obj

if True == tmp7421 {
tmp7418 := Call(__e, PrimFunc(symhdstr), V3225)


tmp7419 := Call(__e, PrimFunc(symhdstr), V3226)


tmp7420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp7418, tmp7419)
}
__typedArg0 := tmp7418
__typedArg1 := tmp7419
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7417 Obj

if True == tmp7420 {
ifres7417 = True


} else {
ifres7417 = False


}

ifres7416 = ifres7417


} else {
ifres7416 = False


}

var ifres7415 Obj

if True == ifres7416 {
ifres7415 = True


} else {
ifres7415 = False


}

ifres7414 = ifres7415


} else {
ifres7414 = False


}

if True == ifres7414 {
tmp7411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3225)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3225
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

if PrimFunc(symshen_4internal_1to_1P_2) == __self {
__self1, __self2 = tmp7411, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3226)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3226
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4internal_1to_1P_2), tmp7411, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3226)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3226
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())
return


} else {
__e.Return(False)
return
}


}


}, 2)

tmp7432 := Call(__e, ns2_1set, symshen_4internal_1to_1P_2, tmp7410)


_ = tmp7432

tmp7433 := MakeNative(func(__e *ControlFlow) {
V3229 := __e.Get(1)
_ = V3229
V3230 := __e.Get(2)
_ = V3230
tmp7454 := Call(__e, PrimFunc(symelement_2), V3229, V3230)


if True == tmp7454 {
__e.Return(V3229)
return
} else {
tmp7452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3229)
}
__typedArg0 := V3229
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7448 Obj

if True == tmp7452 {
tmp7450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3229)
}
__typedArg0 := V3229
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcond, tmp7450)
}
__typedArg0 := symcond
__typedArg1 := tmp7450
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7449 Obj

if True == tmp7451 {
ifres7449 = True


} else {
ifres7449 = False


}

ifres7448 = ifres7449


} else {
ifres7448 = False


}

if True == ifres7448 {
tmp7434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3229)
}
__typedArg0 := V3229
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7435 := Call(__e, PrimFunc(symshen_4process_1cond_1clauses), tmp7434, V3230)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcond, tmp7435)
}
__typedArg0 := symcond
__typedArg1 := tmp7435
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3229)
}
__typedArg0 := V3229
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7442 Obj

if True == tmp7446 {
tmp7444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3229)
}
__typedArg0 := V3229
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7445 := Call(__e, PrimFunc(symshen_4non_1application_2), tmp7444)


var ifres7443 Obj

if True == tmp7445 {
ifres7443 = True


} else {
ifres7443 = False


}

ifres7442 = ifres7443


} else {
ifres7442 = False


}

if True == ifres7442 {
tmp7436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3229)
}
__typedArg0 := V3229
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4special_1case), tmp7436, V3229, V3230)
return


} else {
tmp7440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3229)
}
__typedArg0 := V3229
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7440 {
tmp7437 := MakeNative(func(__e *ControlFlow) {
Z3231 := __e.Get(1)
_ = Z3231
__e.TailApply(PrimFunc(symshen_4process_1applications), Z3231, V3230)
return
}, 1)

tmp7438 := Call(__e, PrimFunc(symmap), tmp7437, V3229)


__e.TailApply(PrimFunc(symshen_4process_1application), tmp7438, V3230)
return


} else {
__e.Return(V3229)
return
}


}


}


}


}, 2)

tmp7455 := Call(__e, ns2_1set, symshen_4process_1applications, tmp7433)


_ = tmp7455

tmp7456 := MakeNative(func(__e *ControlFlow) {
V3234 := __e.Get(1)
_ = V3234
tmp7466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, V3234)
}
__typedArg0 := symdefine
__typedArg1 := V3234
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7466 {
__e.Return(True)
return
} else {
tmp7464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, V3234)
}
__typedArg0 := symdefun
__typedArg1 := V3234
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7464 {
__e.Return(True)
return
} else {
tmp7462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsynonyms, V3234)
}
__typedArg0 := symsynonyms
__typedArg1 := V3234
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7462 {
__e.Return(True)
return
} else {
tmp7460 := Call(__e, PrimFunc(symshen_4special_2), V3234)


if True == tmp7460 {
__e.Return(True)
return
} else {
tmp7458 := Call(__e, PrimFunc(symshen_4extraspecial_2), V3234)


if True == tmp7458 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}


}


}


}, 1)

tmp7467 := Call(__e, ns2_1set, symshen_4non_1application_2, tmp7456)


_ = tmp7467

tmp7468 := MakeNative(func(__e *ControlFlow) {
V3239 := __e.Get(1)
_ = V3239
V3240 := __e.Get(2)
_ = V3240
V3241 := __e.Get(3)
_ = V3241
tmp7710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, V3239)
}
__typedArg0 := symlambda
__typedArg1 := V3239
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7688 Obj

if True == tmp7710 {
tmp7709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7690 Obj

if True == tmp7709 {
tmp7707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp7707)
}
__typedArg0 := symlambda
__typedArg1 := tmp7707
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7692 Obj

if True == tmp7708 {
tmp7705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7705)
}
__typedArg0 := tmp7705
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7694 Obj

if True == tmp7706 {
tmp7702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7702)
}
__typedArg0 := tmp7702
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7703)
}
__typedArg0 := tmp7703
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7696 Obj

if True == tmp7704 {
tmp7698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7698)
}
__typedArg0 := tmp7698
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7699)
}
__typedArg0 := tmp7699
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7700)
}
__typedArg0 := Nil
__typedArg1 := tmp7700
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7697 Obj

if True == tmp7701 {
ifres7697 = True


} else {
ifres7697 = False


}

ifres7696 = ifres7697


} else {
ifres7696 = False


}

var ifres7695 Obj

if True == ifres7696 {
ifres7695 = True


} else {
ifres7695 = False


}

ifres7694 = ifres7695


} else {
ifres7694 = False


}

var ifres7693 Obj

if True == ifres7694 {
ifres7693 = True


} else {
ifres7693 = False


}

ifres7692 = ifres7693


} else {
ifres7692 = False


}

var ifres7691 Obj

if True == ifres7692 {
ifres7691 = True


} else {
ifres7691 = False


}

ifres7690 = ifres7691


} else {
ifres7690 = False


}

var ifres7689 Obj

if True == ifres7690 {
ifres7689 = True


} else {
ifres7689 = False


}

ifres7688 = ifres7689


} else {
ifres7688 = False


}

if True == ifres7688 {
tmp7469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7469)
}
__typedArg0 := tmp7469
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7471)
}
__typedArg0 := tmp7471
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7472)
}
__typedArg0 := tmp7472
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7474 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7473, V3241)


tmp7475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7474, Nil)
}
__typedArg0 := tmp7474
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7470, tmp7475)
}
__typedArg0 := tmp7470
__typedArg1 := tmp7475
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp7476)
}
__typedArg0 := symlambda
__typedArg1 := tmp7476
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, V3239)
}
__typedArg0 := symlet
__typedArg1 := V3239
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7657 Obj

if True == tmp7686 {
tmp7685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7659 Obj

if True == tmp7685 {
tmp7683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp7683)
}
__typedArg0 := symlet
__typedArg1 := tmp7683
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7661 Obj

if True == tmp7684 {
tmp7681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7681)
}
__typedArg0 := tmp7681
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7663 Obj

if True == tmp7682 {
tmp7678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7678)
}
__typedArg0 := tmp7678
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7679)
}
__typedArg0 := tmp7679
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7665 Obj

if True == tmp7680 {
tmp7674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7674)
}
__typedArg0 := tmp7674
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7675)
}
__typedArg0 := tmp7675
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7676)
}
__typedArg0 := tmp7676
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7667 Obj

if True == tmp7677 {
tmp7669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7669)
}
__typedArg0 := tmp7669
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7670)
}
__typedArg0 := tmp7670
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7671)
}
__typedArg0 := tmp7671
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7672)
}
__typedArg0 := Nil
__typedArg1 := tmp7672
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7668 Obj

if True == tmp7673 {
ifres7668 = True


} else {
ifres7668 = False


}

ifres7667 = ifres7668


} else {
ifres7667 = False


}

var ifres7666 Obj

if True == ifres7667 {
ifres7666 = True


} else {
ifres7666 = False


}

ifres7665 = ifres7666


} else {
ifres7665 = False


}

var ifres7664 Obj

if True == ifres7665 {
ifres7664 = True


} else {
ifres7664 = False


}

ifres7663 = ifres7664


} else {
ifres7663 = False


}

var ifres7662 Obj

if True == ifres7663 {
ifres7662 = True


} else {
ifres7662 = False


}

ifres7661 = ifres7662


} else {
ifres7661 = False


}

var ifres7660 Obj

if True == ifres7661 {
ifres7660 = True


} else {
ifres7660 = False


}

ifres7659 = ifres7660


} else {
ifres7659 = False


}

var ifres7658 Obj

if True == ifres7659 {
ifres7658 = True


} else {
ifres7658 = False


}

ifres7657 = ifres7658


} else {
ifres7657 = False


}

if True == ifres7657 {
tmp7477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7477)
}
__typedArg0 := tmp7477
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7479)
}
__typedArg0 := tmp7479
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7480)
}
__typedArg0 := tmp7480
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7482 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7481, V3241)


tmp7483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7483)
}
__typedArg0 := tmp7483
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7484)
}
__typedArg0 := tmp7484
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7485)
}
__typedArg0 := tmp7485
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7487 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7486, V3241)


tmp7488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7487, Nil)
}
__typedArg0 := tmp7487
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7482, tmp7488)
}
__typedArg0 := tmp7482
__typedArg1 := tmp7488
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7478, tmp7489)
}
__typedArg0 := tmp7478
__typedArg1 := tmp7489
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp7490)
}
__typedArg0 := symlet
__typedArg1 := tmp7490
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, V3239)
}
__typedArg0 := symdefun
__typedArg1 := V3239
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7626 Obj

if True == tmp7655 {
tmp7654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7628 Obj

if True == tmp7654 {
tmp7652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp7652)
}
__typedArg0 := symdefun
__typedArg1 := tmp7652
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7630 Obj

if True == tmp7653 {
tmp7650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7650)
}
__typedArg0 := tmp7650
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7632 Obj

if True == tmp7651 {
tmp7647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7647)
}
__typedArg0 := tmp7647
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7648)
}
__typedArg0 := tmp7648
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7634 Obj

if True == tmp7649 {
tmp7643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7643)
}
__typedArg0 := tmp7643
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7644)
}
__typedArg0 := tmp7644
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7645)
}
__typedArg0 := tmp7645
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7636 Obj

if True == tmp7646 {
tmp7638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7638)
}
__typedArg0 := tmp7638
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7639)
}
__typedArg0 := tmp7639
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7640)
}
__typedArg0 := tmp7640
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7641)
}
__typedArg0 := Nil
__typedArg1 := tmp7641
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7637 Obj

if True == tmp7642 {
ifres7637 = True


} else {
ifres7637 = False


}

ifres7636 = ifres7637


} else {
ifres7636 = False


}

var ifres7635 Obj

if True == ifres7636 {
ifres7635 = True


} else {
ifres7635 = False


}

ifres7634 = ifres7635


} else {
ifres7634 = False


}

var ifres7633 Obj

if True == ifres7634 {
ifres7633 = True


} else {
ifres7633 = False


}

ifres7632 = ifres7633


} else {
ifres7632 = False


}

var ifres7631 Obj

if True == ifres7632 {
ifres7631 = True


} else {
ifres7631 = False


}

ifres7630 = ifres7631


} else {
ifres7630 = False


}

var ifres7629 Obj

if True == ifres7630 {
ifres7629 = True


} else {
ifres7629 = False


}

ifres7628 = ifres7629


} else {
ifres7628 = False


}

var ifres7627 Obj

if True == ifres7628 {
ifres7627 = True


} else {
ifres7627 = False


}

ifres7626 = ifres7627


} else {
ifres7626 = False


}

if True == ifres7626 {
__e.Return(V3240)
return
} else {
tmp7624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, V3239)
}
__typedArg0 := symdefine
__typedArg1 := V3239
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7602 Obj

if True == tmp7624 {
tmp7623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7604 Obj

if True == tmp7623 {
tmp7621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp7621)
}
__typedArg0 := symdefine
__typedArg1 := tmp7621
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7606 Obj

if True == tmp7622 {
tmp7619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7619)
}
__typedArg0 := tmp7619
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7608 Obj

if True == tmp7620 {
tmp7616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7616)
}
__typedArg0 := tmp7616
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7617)
}
__typedArg0 := tmp7617
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7610 Obj

if True == tmp7618 {
tmp7612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7612)
}
__typedArg0 := tmp7612
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7613)
}
__typedArg0 := tmp7613
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, tmp7614)
}
__typedArg0 := sym_i
__typedArg1 := tmp7614
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7611 Obj

if True == tmp7615 {
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

var ifres7605 Obj

if True == ifres7606 {
ifres7605 = True


} else {
ifres7605 = False


}

ifres7604 = ifres7605


} else {
ifres7604 = False


}

var ifres7603 Obj

if True == ifres7604 {
ifres7603 = True


} else {
ifres7603 = False


}

ifres7602 = ifres7603


} else {
ifres7602 = False


}

if True == ifres7602 {
tmp7491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7491)
}
__typedArg0 := tmp7491
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7493)
}
__typedArg0 := tmp7493
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7495)
}
__typedArg0 := tmp7495
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7496)
}
__typedArg0 := tmp7496
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7498 := Call(__e, PrimFunc(symshen_4process_1after_1type), tmp7494, tmp7497, V3241)


tmp7499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, tmp7498)
}
__typedArg0 := sym_i
__typedArg1 := tmp7498
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7492, tmp7499)
}
__typedArg0 := tmp7492
__typedArg1 := tmp7499
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp7500)
}
__typedArg0 := symdefine
__typedArg1 := tmp7500
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, V3239)
}
__typedArg0 := symdefine
__typedArg1 := V3239
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7589 Obj

if True == tmp7600 {
tmp7599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7591 Obj

if True == tmp7599 {
tmp7597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp7597)
}
__typedArg0 := symdefine
__typedArg1 := tmp7597
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7593 Obj

if True == tmp7598 {
tmp7595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7595)
}
__typedArg0 := tmp7595
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7594 Obj

if True == tmp7596 {
ifres7594 = True


} else {
ifres7594 = False


}

ifres7593 = ifres7594


} else {
ifres7593 = False


}

var ifres7592 Obj

if True == ifres7593 {
ifres7592 = True


} else {
ifres7592 = False


}

ifres7591 = ifres7592


} else {
ifres7591 = False


}

var ifres7590 Obj

if True == ifres7591 {
ifres7590 = True


} else {
ifres7590 = False


}

ifres7589 = ifres7590


} else {
ifres7589 = False


}

if True == ifres7589 {
tmp7501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7501)
}
__typedArg0 := tmp7501
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7503 := MakeNative(func(__e *ControlFlow) {
Z3242 := __e.Get(1)
_ = Z3242
__e.TailApply(PrimFunc(symshen_4process_1applications), Z3242, V3241)
return
}, 1)

tmp7504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7504)
}
__typedArg0 := tmp7504
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7506 := Call(__e, PrimFunc(symmap), tmp7503, tmp7505)


tmp7507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7502, tmp7506)
}
__typedArg0 := tmp7502
__typedArg1 := tmp7506
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp7507)
}
__typedArg0 := symdefine
__typedArg1 := tmp7507
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsynonyms, V3239)
}
__typedArg0 := symsynonyms
__typedArg1 := V3239
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7587 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsynonyms, V3240)
}
__typedArg0 := symsynonyms
__typedArg1 := V3240
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
} else {
tmp7585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtype, V3239)
}
__typedArg0 := symtype
__typedArg1 := V3239
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7563 Obj

if True == tmp7585 {
tmp7584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7565 Obj

if True == tmp7584 {
tmp7582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtype, tmp7582)
}
__typedArg0 := symtype
__typedArg1 := tmp7582
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7567 Obj

if True == tmp7583 {
tmp7580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7580)
}
__typedArg0 := tmp7580
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7569 Obj

if True == tmp7581 {
tmp7577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7577)
}
__typedArg0 := tmp7577
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7578)
}
__typedArg0 := tmp7578
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7571 Obj

if True == tmp7579 {
tmp7573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7573)
}
__typedArg0 := tmp7573
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7574)
}
__typedArg0 := tmp7574
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7575)
}
__typedArg0 := Nil
__typedArg1 := tmp7575
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7572 Obj

if True == tmp7576 {
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

var ifres7568 Obj

if True == ifres7569 {
ifres7568 = True


} else {
ifres7568 = False


}

ifres7567 = ifres7568


} else {
ifres7567 = False


}

var ifres7566 Obj

if True == ifres7567 {
ifres7566 = True


} else {
ifres7566 = False


}

ifres7565 = ifres7566


} else {
ifres7565 = False


}

var ifres7564 Obj

if True == ifres7565 {
ifres7564 = True


} else {
ifres7564 = False


}

ifres7563 = ifres7564


} else {
ifres7563 = False


}

if True == ifres7563 {
tmp7508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7508)
}
__typedArg0 := tmp7508
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7510 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7509, V3241)


tmp7511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7511)
}
__typedArg0 := tmp7511
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7510, tmp7512)
}
__typedArg0 := tmp7510
__typedArg1 := tmp7512
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp7513)
}
__typedArg0 := symtype
__typedArg1 := tmp7513
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput_7, V3239)
}
__typedArg0 := syminput_7
__typedArg1 := V3239
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7539 Obj

if True == tmp7561 {
tmp7560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7541 Obj

if True == tmp7560 {
tmp7558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput_7, tmp7558)
}
__typedArg0 := syminput_7
__typedArg1 := tmp7558
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7543 Obj

if True == tmp7559 {
tmp7556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7556)
}
__typedArg0 := tmp7556
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7545 Obj

if True == tmp7557 {
tmp7553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7553)
}
__typedArg0 := tmp7553
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7554)
}
__typedArg0 := tmp7554
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7547 Obj

if True == tmp7555 {
tmp7549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7549)
}
__typedArg0 := tmp7549
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7550)
}
__typedArg0 := tmp7550
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7551)
}
__typedArg0 := Nil
__typedArg1 := tmp7551
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7548 Obj

if True == tmp7552 {
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

var ifres7542 Obj

if True == ifres7543 {
ifres7542 = True


} else {
ifres7542 = False


}

ifres7541 = ifres7542


} else {
ifres7541 = False


}

var ifres7540 Obj

if True == ifres7541 {
ifres7540 = True


} else {
ifres7540 = False


}

ifres7539 = ifres7540


} else {
ifres7539 = False


}

if True == ifres7539 {
tmp7514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7514)
}
__typedArg0 := tmp7514
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7516)
}
__typedArg0 := tmp7516
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7517)
}
__typedArg0 := tmp7517
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7519 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7518, V3241)


tmp7520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7519, Nil)
}
__typedArg0 := tmp7519
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7515, tmp7520)
}
__typedArg0 := tmp7515
__typedArg1 := tmp7520
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput_7, tmp7521)
}
__typedArg0 := syminput_7
__typedArg1 := tmp7521
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7533 Obj

if True == tmp7537 {
tmp7535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7536 := Call(__e, PrimFunc(symshen_4special_2), tmp7535)


var ifres7534 Obj

if True == tmp7536 {
ifres7534 = True


} else {
ifres7534 = False


}

ifres7533 = ifres7534


} else {
ifres7533 = False


}

if True == ifres7533 {
tmp7522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7523 := MakeNative(func(__e *ControlFlow) {
Z3243 := __e.Get(1)
_ = Z3243
__e.TailApply(PrimFunc(symshen_4process_1applications), Z3243, V3241)
return
}, 1)

tmp7524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7525 := Call(__e, PrimFunc(symmap), tmp7523, tmp7524)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7522, tmp7525)
}
__typedArg0 := tmp7522
__typedArg1 := tmp7525
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7527 Obj

if True == tmp7531 {
tmp7529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3240)
}
__typedArg0 := V3240
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7530 := Call(__e, PrimFunc(symshen_4extraspecial_2), tmp7529)


var ifres7528 Obj

if True == tmp7530 {
ifres7528 = True


} else {
ifres7528 = False


}

ifres7527 = ifres7528


} else {
ifres7527 = False


}

if True == ifres7527 {
__e.Return(V3240)
return
} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4special_1case)
return
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


}, 3)

tmp7711 := Call(__e, ns2_1set, symshen_4special_1case, tmp7468)


_ = tmp7711

tmp7712 := MakeNative(func(__e *ControlFlow) {
V3246 := __e.Get(1)
_ = V3246
V3247 := __e.Get(2)
_ = V3247
tmp7742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3246)
}
__typedArg0 := Nil
__typedArg1 := V3246
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7742 {
__e.Return(Nil)
return
} else {
tmp7740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3246)
}
__typedArg0 := V3246
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7725 Obj

if True == tmp7740 {
tmp7738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3246)
}
__typedArg0 := V3246
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7738)
}
__typedArg0 := tmp7738
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7727 Obj

if True == tmp7739 {
tmp7735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3246)
}
__typedArg0 := V3246
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7735)
}
__typedArg0 := tmp7735
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7736)
}
__typedArg0 := tmp7736
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7729 Obj

if True == tmp7737 {
tmp7731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3246)
}
__typedArg0 := V3246
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7731)
}
__typedArg0 := tmp7731
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7732)
}
__typedArg0 := tmp7732
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7733)
}
__typedArg0 := Nil
__typedArg1 := tmp7733
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7730 Obj

if True == tmp7734 {
ifres7730 = True


} else {
ifres7730 = False


}

ifres7729 = ifres7730


} else {
ifres7729 = False


}

var ifres7728 Obj

if True == ifres7729 {
ifres7728 = True


} else {
ifres7728 = False


}

ifres7727 = ifres7728


} else {
ifres7727 = False


}

var ifres7726 Obj

if True == ifres7727 {
ifres7726 = True


} else {
ifres7726 = False


}

ifres7725 = ifres7726


} else {
ifres7725 = False


}

if True == ifres7725 {
tmp7713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3246)
}
__typedArg0 := V3246
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7713)
}
__typedArg0 := tmp7713
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7715 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7714, V3247)


tmp7716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3246)
}
__typedArg0 := V3246
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7716)
}
__typedArg0 := tmp7716
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7717)
}
__typedArg0 := tmp7717
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7719 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7718, V3247)


tmp7720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7719, Nil)
}
__typedArg0 := tmp7719
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7715, tmp7720)
}
__typedArg0 := tmp7715
__typedArg1 := tmp7720
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3246)
}
__typedArg0 := V3246
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7723 := Call(__e, PrimFunc(symshen_4process_1cond_1clauses), tmp7722, V3247)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7721, tmp7723)
}
__typedArg0 := tmp7721
__typedArg1 := tmp7723
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4process_1cond_1clauses)
return
}


}


}, 2)

tmp7743 := Call(__e, ns2_1set, symshen_4process_1cond_1clauses, tmp7712)


_ = tmp7743

tmp7744 := MakeNative(func(__e *ControlFlow) {
V3250 := __e.Get(1)
_ = V3250
V3251 := __e.Get(2)
_ = V3251
V3252 := __e.Get(3)
_ = V3252
tmp7760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3251)
}
__typedArg0 := V3251
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7756 Obj

if True == tmp7760 {
tmp7758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3251)
}
__typedArg0 := V3251
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, tmp7758)
}
__typedArg0 := sym_j
__typedArg1 := tmp7758
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7757 Obj

if True == tmp7759 {
ifres7757 = True


} else {
ifres7757 = False


}

ifres7756 = ifres7757


} else {
ifres7756 = False


}

if True == ifres7756 {
tmp7745 := MakeNative(func(__e *ControlFlow) {
Z3253 := __e.Get(1)
_ = Z3253
__e.TailApply(PrimFunc(symshen_4process_1applications), Z3253, V3252)
return
}, 1)

tmp7746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3251)
}
__typedArg0 := V3251
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7747 := Call(__e, PrimFunc(symmap), tmp7745, tmp7746)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, tmp7747)
}
__typedArg0 := sym_j
__typedArg1 := tmp7747
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3251)
}
__typedArg0 := V3251
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7754 {
tmp7748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3251)
}
__typedArg0 := V3251
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3251)
}
__typedArg0 := V3251
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7750 := Call(__e, PrimFunc(symshen_4process_1after_1type), V3250, tmp7749, V3252)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7748, tmp7750)
}
__typedArg0 := tmp7748
__typedArg1 := tmp7750
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7751 := Call(__e, PrimFunc(symshen_4app), V3250, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing } in "))
__typedS1, __typedOK1 := TypedString(tmp7751)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing } in ")
__typedArg1 := tmp7751
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing } in "))
__typedS1, __typedOK1 := TypedString(tmp7751)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing } in ")
__typedArg1 := tmp7751
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}


}, 3)

tmp7761 := Call(__e, ns2_1set, symshen_4process_1after_1type, tmp7744)


_ = tmp7761

tmp7762 := MakeNative(func(__e *ControlFlow) {
V3254 := __e.Get(1)
_ = V3254
V3255 := __e.Get(2)
_ = V3255
tmp7805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3254)
}
__typedArg0 := V3254
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7805 {
tmp7763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3254)
}
__typedArg0 := V3254
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7764 := Call(__e, PrimFunc(symarity), tmp7763)


let__5446 := tmp7764
_ = let__5446

tmp7765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3254)
}
__typedArg0 := V3254
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7766 := Call(__e, PrimFunc(symlength), tmp7765)


let__5447 := tmp7766
_ = let__5447

tmp7803 := Call(__e, PrimFunc(symelement_2), V3254, V3255)


if True == tmp7803 {
__e.Return(V3254)
return
} else {
tmp7800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3254)
}
__typedArg0 := V3254
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7801 := Call(__e, PrimFunc(symshen_4shen_1call_2), tmp7800)


if True == tmp7801 {
__e.Return(V3254)
return
} else {
tmp7798 := Call(__e, PrimFunc(symshen_4foreign_2), V3254)


if True == tmp7798 {
__e.TailApply(PrimFunc(symshen_4unpack_1foreign), V3254)
return
} else {
tmp7796 := Call(__e, PrimFunc(symshen_4fn_1call_2), V3254)


if True == tmp7796 {
__e.TailApply(PrimFunc(symshen_4fn_1call), V3254)
return
} else {
tmp7794 := Call(__e, PrimFunc(symshen_4zero_1place_2), V3254)


if True == tmp7794 {
__e.Return(V3254)
return
} else {
tmp7791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3254)
}
__typedArg0 := V3254
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7792 := Call(__e, PrimFunc(symshen_4undefined_1f_2), tmp7791, let__5446)


if True == tmp7792 {
tmp7767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3254)
}
__typedArg0 := V3254
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7767, Nil)
}
__typedArg0 := tmp7767
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp7768)
}
__typedArg0 := symfn
__typedArg1 := tmp7768
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3254)
}
__typedArg0 := V3254
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7769, tmp7770)
}
__typedArg0 := tmp7769
__typedArg1 := tmp7770
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp7771)
return


} else {
tmp7788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3254)
}
__typedArg0 := V3254
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp7788)
}
__typedArg0 := tmp7788
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp7789 {
__e.TailApply(PrimFunc(symshen_4simple_1curry), V3254)
return
} else {
tmp7785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3254)
}
__typedArg0 := V3254
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7786 := Call(__e, PrimFunc(symshen_4application_2), tmp7785)


if True == tmp7786 {
__e.TailApply(PrimFunc(symshen_4simple_1curry), V3254)
return
} else {
tmp7782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3254)
}
__typedArg0 := V3254
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7783 := Call(__e, PrimFunc(symshen_4partial_1application_d_2), tmp7782, let__5446, let__5447)


if True == tmp7783 {
__e.TailApply(PrimFunc(symshen_4lambda_1function), V3254, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(let__5446)
__typedN1, __typedOK1 := TypedFloat64(let__5447)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := let__5446
__typedArg1 := let__5447
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())
return


} else {
tmp7779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3254)
}
__typedArg0 := V3254
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7780 := Call(__e, PrimFunc(symshen_4overapplication_2), tmp7779, let__5446, let__5447)


if True == tmp7780 {
tmp7773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3254)
}
__typedArg0 := V3254
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7773, Nil)
}
__typedArg0 := tmp7773
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp7774)
}
__typedArg0 := symfn
__typedArg1 := tmp7774
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3254)
}
__typedArg0 := V3254
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7775, tmp7776)
}
__typedArg0 := tmp7775
__typedArg1 := tmp7776
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp7777)
return


} else {
__e.Return(V3254)
return
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
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4process_1application)
return
}


}, 2)

tmp7806 := Call(__e, ns2_1set, symshen_4process_1application, tmp7762)


_ = tmp7806

tmp7807 := MakeNative(func(__e *ControlFlow) {
V3258 := __e.Get(1)
_ = V3258
tmp7833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3258)
}
__typedArg0 := V3258
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7813 Obj

if True == tmp7833 {
tmp7831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3258)
}
__typedArg0 := V3258
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7831)
}
__typedArg0 := tmp7831
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7815 Obj

if True == tmp7832 {
tmp7828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3258)
}
__typedArg0 := V3258
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7828)
}
__typedArg0 := tmp7828
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symforeign, tmp7829)
}
__typedArg0 := symforeign
__typedArg1 := tmp7829
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7817 Obj

if True == tmp7830 {
tmp7825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3258)
}
__typedArg0 := V3258
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7825)
}
__typedArg0 := tmp7825
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7826)
}
__typedArg0 := tmp7826
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7819 Obj

if True == tmp7827 {
tmp7821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3258)
}
__typedArg0 := V3258
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7821)
}
__typedArg0 := tmp7821
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7822)
}
__typedArg0 := tmp7822
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7823)
}
__typedArg0 := Nil
__typedArg1 := tmp7823
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7820 Obj

if True == tmp7824 {
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

var ifres7816 Obj

if True == ifres7817 {
ifres7816 = True


} else {
ifres7816 = False


}

ifres7815 = ifres7816


} else {
ifres7815 = False


}

var ifres7814 Obj

if True == ifres7815 {
ifres7814 = True


} else {
ifres7814 = False


}

ifres7813 = ifres7814


} else {
ifres7813 = False


}

if True == ifres7813 {
tmp7808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3258)
}
__typedArg0 := V3258
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7808)
}
__typedArg0 := tmp7808
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7809)
}
__typedArg0 := tmp7809
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3258)
}
__typedArg0 := V3258
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7810, tmp7811)
}
__typedArg0 := tmp7810
__typedArg1 := tmp7811
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4unpack_1foreign)
return
}


}, 1)

tmp7834 := Call(__e, ns2_1set, symshen_4unpack_1foreign, tmp7807)


_ = tmp7834

tmp7835 := MakeNative(func(__e *ControlFlow) {
V3261 := __e.Get(1)
_ = V3261
tmp7857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3261)
}
__typedArg0 := V3261
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7837 Obj

if True == tmp7857 {
tmp7855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3261)
}
__typedArg0 := V3261
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7855)
}
__typedArg0 := tmp7855
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7839 Obj

if True == tmp7856 {
tmp7852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3261)
}
__typedArg0 := V3261
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7852)
}
__typedArg0 := tmp7852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symforeign, tmp7853)
}
__typedArg0 := symforeign
__typedArg1 := tmp7853
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7841 Obj

if True == tmp7854 {
tmp7849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3261)
}
__typedArg0 := V3261
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7849)
}
__typedArg0 := tmp7849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7850)
}
__typedArg0 := tmp7850
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7843 Obj

if True == tmp7851 {
tmp7845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3261)
}
__typedArg0 := V3261
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7845)
}
__typedArg0 := tmp7845
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7846)
}
__typedArg0 := tmp7846
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7847)
}
__typedArg0 := Nil
__typedArg1 := tmp7847
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7844 Obj

if True == tmp7848 {
ifres7844 = True


} else {
ifres7844 = False


}

ifres7843 = ifres7844


} else {
ifres7843 = False


}

var ifres7842 Obj

if True == ifres7843 {
ifres7842 = True


} else {
ifres7842 = False


}

ifres7841 = ifres7842


} else {
ifres7841 = False


}

var ifres7840 Obj

if True == ifres7841 {
ifres7840 = True


} else {
ifres7840 = False


}

ifres7839 = ifres7840


} else {
ifres7839 = False


}

var ifres7838 Obj

if True == ifres7839 {
ifres7838 = True


} else {
ifres7838 = False


}

ifres7837 = ifres7838


} else {
ifres7837 = False


}

if True == ifres7837 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp7858 := Call(__e, ns2_1set, symshen_4foreign_2, tmp7835)


_ = tmp7858

tmp7859 := MakeNative(func(__e *ControlFlow) {
V3264 := __e.Get(1)
_ = V3264
tmp7865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3264)
}
__typedArg0 := V3264
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7861 Obj

if True == tmp7865 {
tmp7863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3264)
}
__typedArg0 := V3264
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7863)
}
__typedArg0 := Nil
__typedArg1 := tmp7863
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7862 Obj

if True == tmp7864 {
ifres7862 = True


} else {
ifres7862 = False


}

ifres7861 = ifres7862


} else {
ifres7861 = False


}

if True == ifres7861 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp7866 := Call(__e, ns2_1set, symshen_4zero_1place_2, tmp7859)


_ = tmp7866

tmp7867 := MakeNative(func(__e *ControlFlow) {
V3265 := __e.Get(1)
_ = V3265
tmp7872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V3265)
}
__typedArg0 := V3265
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp7872 {
tmp7869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3265)
}
__typedArg0 := V3265
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp7870 := Call(__e, PrimFunc(symshen_4internal_1to_1shen_2), tmp7869)


if True == tmp7870 {
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

tmp7873 := Call(__e, ns2_1set, symshen_4shen_1call_2, tmp7867)


_ = tmp7873

tmp7874 := MakeNative(func(__e *ControlFlow) {
V3270 := __e.Get(1)
_ = V3270
tmp7904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3270)
}
__typedArg0 := V3270
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7891 Obj

if True == tmp7904 {
tmp7902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3270)
}
__typedArg0 := V3270
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symprotect, tmp7902)
}
__typedArg0 := symprotect
__typedArg1 := tmp7902
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7893 Obj

if True == tmp7903 {
tmp7900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3270)
}
__typedArg0 := V3270
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7900)
}
__typedArg0 := tmp7900
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7895 Obj

if True == tmp7901 {
tmp7897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3270)
}
__typedArg0 := V3270
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7897)
}
__typedArg0 := tmp7897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7898)
}
__typedArg0 := Nil
__typedArg1 := tmp7898
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7896 Obj

if True == tmp7899 {
ifres7896 = True


} else {
ifres7896 = False


}

ifres7895 = ifres7896


} else {
ifres7895 = False


}

var ifres7894 Obj

if True == ifres7895 {
ifres7894 = True


} else {
ifres7894 = False


}

ifres7893 = ifres7894


} else {
ifres7893 = False


}

var ifres7892 Obj

if True == ifres7893 {
ifres7892 = True


} else {
ifres7892 = False


}

ifres7891 = ifres7892


} else {
ifres7891 = False


}

if True == ifres7891 {
__e.Return(False)
return
} else {
tmp7889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3270)
}
__typedArg0 := V3270
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7876 Obj

if True == tmp7889 {
tmp7887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3270)
}
__typedArg0 := V3270
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symforeign, tmp7887)
}
__typedArg0 := symforeign
__typedArg1 := tmp7887
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7878 Obj

if True == tmp7888 {
tmp7885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3270)
}
__typedArg0 := V3270
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7885)
}
__typedArg0 := tmp7885
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7880 Obj

if True == tmp7886 {
tmp7882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3270)
}
__typedArg0 := V3270
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7882)
}
__typedArg0 := tmp7882
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7883)
}
__typedArg0 := Nil
__typedArg1 := tmp7883
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7881 Obj

if True == tmp7884 {
ifres7881 = True


} else {
ifres7881 = False


}

ifres7880 = ifres7881


} else {
ifres7880 = False


}

var ifres7879 Obj

if True == ifres7880 {
ifres7879 = True


} else {
ifres7879 = False


}

ifres7878 = ifres7879


} else {
ifres7878 = False


}

var ifres7877 Obj

if True == ifres7878 {
ifres7877 = True


} else {
ifres7877 = False


}

ifres7876 = ifres7877


} else {
ifres7876 = False


}

if True == ifres7876 {
__e.Return(False)
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3270)
}
__typedArg0 := V3270
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})())
return
}


}


}, 1)

tmp7905 := Call(__e, ns2_1set, symshen_4application_2, tmp7874)


_ = tmp7905

tmp7906 := MakeNative(func(__e *ControlFlow) {
V3275 := __e.Get(1)
_ = V3275
V3276 := __e.Get(2)
_ = V3276
tmp7914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(-1), V3276)
}
__typedArg0 := MakeInteger(-1)
__typedArg1 := V3276
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7914 {
tmp7912 := Call(__e, PrimFunc(symshen_4lowercase_1symbol_2), V3275)


if True == tmp7912 {
tmp7908 := Call(__e, PrimFunc(symexternal), symshen)


tmp7909 := Call(__e, PrimFunc(symelement_2), V3275, tmp7908)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7909)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7909
return Call(__e, PrimFunc(symnot), __typedArg0)
})() {
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


} else {
__e.Return(False)
return
}


}, 2)

tmp7915 := Call(__e, ns2_1set, symshen_4undefined_1f_2, tmp7906)


_ = tmp7915

tmp7916 := MakeNative(func(__e *ControlFlow) {
V3277 := __e.Get(1)
_ = V3277
tmp7921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V3277)
}
__typedArg0 := V3277
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp7921 {
tmp7918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V3277)
}
__typedArg0 := V3277
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7918)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7918
return Call(__e, PrimFunc(symnot), __typedArg0)
})() {
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

tmp7922 := Call(__e, ns2_1set, symshen_4lowercase_1symbol_2, tmp7916)


_ = tmp7922

tmp7923 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V3278 := __self1
_ = V3278
tmp7953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3278)
}
__typedArg0 := V3278
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7944 Obj

if True == tmp7953 {
tmp7951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3278)
}
__typedArg0 := V3278
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7951)
}
__typedArg0 := tmp7951
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7946 Obj

if True == tmp7952 {
tmp7948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3278)
}
__typedArg0 := V3278
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7948)
}
__typedArg0 := tmp7948
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7949)
}
__typedArg0 := Nil
__typedArg1 := tmp7949
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7947 Obj

if True == tmp7950 {
ifres7947 = True


} else {
ifres7947 = False


}

ifres7946 = ifres7947


} else {
ifres7946 = False


}

var ifres7945 Obj

if True == ifres7946 {
ifres7945 = True


} else {
ifres7945 = False


}

ifres7944 = ifres7945


} else {
ifres7944 = False


}

if True == ifres7944 {
__e.Return(V3278)
return
} else {
tmp7942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3278)
}
__typedArg0 := V3278
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7933 Obj

if True == tmp7942 {
tmp7940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3278)
}
__typedArg0 := V3278
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7940)
}
__typedArg0 := tmp7940
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7935 Obj

if True == tmp7941 {
tmp7937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3278)
}
__typedArg0 := V3278
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7937)
}
__typedArg0 := tmp7937
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7938)
}
__typedArg0 := tmp7938
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7936 Obj

if True == tmp7939 {
ifres7936 = True


} else {
ifres7936 = False


}

ifres7935 = ifres7936


} else {
ifres7935 = False


}

var ifres7934 Obj

if True == ifres7935 {
ifres7934 = True


} else {
ifres7934 = False


}

ifres7933 = ifres7934


} else {
ifres7933 = False


}

if True == ifres7933 {
tmp7924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3278)
}
__typedArg0 := V3278
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3278)
}
__typedArg0 := V3278
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7925)
}
__typedArg0 := tmp7925
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7926, Nil)
}
__typedArg0 := tmp7926
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7924, tmp7927)
}
__typedArg0 := tmp7924
__typedArg1 := tmp7927
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3278)
}
__typedArg0 := V3278
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7929)
}
__typedArg0 := tmp7929
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7928, tmp7930)
}
__typedArg0 := tmp7928
__typedArg1 := tmp7930
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4simple_1curry) == __self {
__self1 = tmp7931
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp7931)
return


} else {
__e.Return(V3278)
return
}


}


}, 1)

tmp7954 := Call(__e, ns2_1set, symshen_4simple_1curry, tmp7923)


_ = tmp7954

tmp7955 := MakeNative(func(__e *ControlFlow) {
V3279 := __e.Get(1)
_ = V3279
__e.TailApply(PrimFunc(symfn), V3279)
return
}, 1)

tmp7956 := Call(__e, ns2_1set, symfunction, tmp7955)


_ = tmp7956

tmp7957 := MakeNative(func(__e *ControlFlow) {
V3280 := __e.Get(1)
_ = V3280
tmp7964 := Call(__e, PrimFunc(symarity), V3280)


tmp7965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp7964, MakeInteger(0))
}
__typedArg0 := tmp7964
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7965 {
__e.TailApply(V3280)
return
} else {
tmp7958 := MakeNative(func(__e *ControlFlow) {
tmp7959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V3280, symshen_4lambda_1form, tmp7959)
return


}, 0)

tmp7960 := MakeNative(func(__e *ControlFlow) {
Z3281 := __e.Get(1)
_ = Z3281
tmp7961 := Call(__e, PrimFunc(symshen_4app), V3280, MakeString(" is undefined\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("fn: "))
__typedS1, __typedOK1 := TypedString(tmp7961)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("fn: ")
__typedArg1 := tmp7961
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("fn: "))
__typedS1, __typedOK1 := TypedString(tmp7961)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("fn: ")
__typedArg1 := tmp7961
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

__e.TailApply(try_1catch, tmp7958, tmp7960)
return


}


}, 1)

tmp7966 := Call(__e, ns2_1set, symfn, tmp7957)


_ = tmp7966

tmp7967 := MakeNative(func(__e *ControlFlow) {
V3284 := __e.Get(1)
_ = V3284
tmp7997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3284)
}
__typedArg0 := V3284
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7984 Obj

if True == tmp7997 {
tmp7995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3284)
}
__typedArg0 := V3284
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfn, tmp7995)
}
__typedArg0 := symfn
__typedArg1 := tmp7995
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7986 Obj

if True == tmp7996 {
tmp7993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3284)
}
__typedArg0 := V3284
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7993)
}
__typedArg0 := tmp7993
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7988 Obj

if True == tmp7994 {
tmp7990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3284)
}
__typedArg0 := V3284
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7990)
}
__typedArg0 := tmp7990
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7991)
}
__typedArg0 := Nil
__typedArg1 := tmp7991
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7989 Obj

if True == tmp7992 {
ifres7989 = True


} else {
ifres7989 = False


}

ifres7988 = ifres7989


} else {
ifres7988 = False


}

var ifres7987 Obj

if True == ifres7988 {
ifres7987 = True


} else {
ifres7987 = False


}

ifres7986 = ifres7987


} else {
ifres7986 = False


}

var ifres7985 Obj

if True == ifres7986 {
ifres7985 = True


} else {
ifres7985 = False


}

ifres7984 = ifres7985


} else {
ifres7984 = False


}

if True == ifres7984 {
__e.Return(True)
return
} else {
tmp7982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3284)
}
__typedArg0 := V3284
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7969 Obj

if True == tmp7982 {
tmp7980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3284)
}
__typedArg0 := V3284
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfunction, tmp7980)
}
__typedArg0 := symfunction
__typedArg1 := tmp7980
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7971 Obj

if True == tmp7981 {
tmp7978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3284)
}
__typedArg0 := V3284
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7978)
}
__typedArg0 := tmp7978
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7973 Obj

if True == tmp7979 {
tmp7975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3284)
}
__typedArg0 := V3284
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7975)
}
__typedArg0 := tmp7975
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7976)
}
__typedArg0 := Nil
__typedArg1 := tmp7976
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7974 Obj

if True == tmp7977 {
ifres7974 = True


} else {
ifres7974 = False


}

ifres7973 = ifres7974


} else {
ifres7973 = False


}

var ifres7972 Obj

if True == ifres7973 {
ifres7972 = True


} else {
ifres7972 = False


}

ifres7971 = ifres7972


} else {
ifres7971 = False


}

var ifres7970 Obj

if True == ifres7971 {
ifres7970 = True


} else {
ifres7970 = False


}

ifres7969 = ifres7970


} else {
ifres7969 = False


}

if True == ifres7969 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp7998 := Call(__e, ns2_1set, symshen_4fn_1call_2, tmp7967)


_ = tmp7998

tmp7999 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V3285 := __self1
_ = V3285
tmp8039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3285)
}
__typedArg0 := V3285
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8026 Obj

if True == tmp8039 {
tmp8037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3285)
}
__typedArg0 := V3285
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfunction, tmp8037)
}
__typedArg0 := symfunction
__typedArg1 := tmp8037
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8028 Obj

if True == tmp8038 {
tmp8035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3285)
}
__typedArg0 := V3285
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8035)
}
__typedArg0 := tmp8035
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8030 Obj

if True == tmp8036 {
tmp8032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3285)
}
__typedArg0 := V3285
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp8032)
}
__typedArg0 := tmp8032
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8033)
}
__typedArg0 := Nil
__typedArg1 := tmp8033
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8031 Obj

if True == tmp8034 {
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

var ifres8027 Obj

if True == ifres8028 {
ifres8027 = True


} else {
ifres8027 = False


}

ifres8026 = ifres8027


} else {
ifres8026 = False


}

if True == ifres8026 {
tmp8000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3285)
}
__typedArg0 := V3285
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp8000)
}
__typedArg0 := symfn
__typedArg1 := tmp8000
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4fn_1call) == __self {
__self1 = tmp8001
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4fn_1call), tmp8001)
return


} else {
tmp8024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3285)
}
__typedArg0 := V3285
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8011 Obj

if True == tmp8024 {
tmp8022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3285)
}
__typedArg0 := V3285
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfn, tmp8022)
}
__typedArg0 := symfn
__typedArg1 := tmp8022
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8013 Obj

if True == tmp8023 {
tmp8020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3285)
}
__typedArg0 := V3285
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8020)
}
__typedArg0 := tmp8020
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8015 Obj

if True == tmp8021 {
tmp8017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3285)
}
__typedArg0 := V3285
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp8017)
}
__typedArg0 := tmp8017
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8018)
}
__typedArg0 := Nil
__typedArg1 := tmp8018
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8016 Obj

if True == tmp8019 {
ifres8016 = True


} else {
ifres8016 = False


}

ifres8015 = ifres8016


} else {
ifres8015 = False


}

var ifres8014 Obj

if True == ifres8015 {
ifres8014 = True


} else {
ifres8014 = False


}

ifres8013 = ifres8014


} else {
ifres8013 = False


}

var ifres8012 Obj

if True == ifres8013 {
ifres8012 = True


} else {
ifres8012 = False


}

ifres8011 = ifres8012


} else {
ifres8011 = False


}

if True == ifres8011 {
tmp8002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3285)
}
__typedArg0 := V3285
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp8002)
}
__typedArg0 := tmp8002
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8004 := Call(__e, PrimFunc(symarity), tmp8003)


let__5448 := tmp8004
_ = let__5448

tmp8009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__5448, MakeInteger(-1))
}
__typedArg0 := let__5448
__typedArg1 := MakeInteger(-1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8009 {
__e.Return(V3285)
return
} else {
tmp8007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__5448, MakeInteger(0))
}
__typedArg0 := let__5448
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8007 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3285)
}
__typedArg0 := V3285
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
} else {
tmp8005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3285)
}
__typedArg0 := V3285
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4lambda_1function), tmp8005, let__5448)
return


}


}


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4fn_1call)
return
}


}


}, 1)

tmp8040 := Call(__e, ns2_1set, symshen_4fn_1call, tmp7999)


_ = tmp8040

tmp8041 := MakeNative(func(__e *ControlFlow) {
V3287 := __e.Get(1)
_ = V3287
V3288 := __e.Get(2)
_ = V3288
V3289 := __e.Get(3)
_ = V3289
let__5449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(V3288)
__typedN1, __typedOK1 := TypedFloat64(V3289)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := V3288
__typedArg1 := V3289
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})()
_ = let__5449

var ifres8048 Obj

if True == let__5449 {
tmp8056 := Call(__e, PrimFunc(symshen_4loading_2))


var ifres8050 Obj

if True == tmp8056 {
tmp8052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, Nil)
}
__typedArg0 := sym_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp8052)
}
__typedArg0 := sym_7
__typedArg1 := tmp8052
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8054 := Call(__e, PrimFunc(symelement_2), V3287, tmp8053)


tmp8055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp8054)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp8054
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres8051 Obj

if True == tmp8055 {
ifres8051 = True


} else {
ifres8051 = False


}

ifres8050 = ifres8051


} else {
ifres8050 = False


}

var ifres8049 Obj

if True == ifres8050 {
ifres8049 = True


} else {
ifres8049 = False


}

ifres8048 = ifres8049


} else {
ifres8048 = False


}

var ifres8043 Obj

if True == ifres8048 {
tmp8044 := Call(__e, PrimFunc(symshen_4app), V3287, MakeString("\n"), symshen_4a)


tmp8045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("partial application of "))
__typedS1, __typedOK1 := TypedString(tmp8044)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("partial application of ")
__typedArg1 := tmp8044
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp8046 := Call(__e, PrimFunc(symstoutput))


tmp8047 := Call(__e, PrimFunc(sympr), tmp8045, tmp8046)


ifres8043 = tmp8047


} else {
ifres8043 = symshen_4skip


}

let__5450 := ifres8043
_ = let__5450

__e.Return(let__5449)
return


}, 3)

tmp8057 := Call(__e, ns2_1set, symshen_4partial_1application_d_2, tmp8041)


_ = tmp8057

tmp8058 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dloading_2_d)
}
__typedArg0 := symshen_4_dloading_2_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp8059 := Call(__e, ns2_1set, symshen_4loading_2, tmp8058)


_ = tmp8059

tmp8060 := MakeNative(func(__e *ControlFlow) {
V3296 := __e.Get(1)
_ = V3296
V3297 := __e.Get(2)
_ = V3297
V3298 := __e.Get(3)
_ = V3298
tmp8076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(-1), V3297)
}
__typedArg0 := MakeInteger(-1)
__typedArg1 := V3297
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8076 {
__e.Return(False)
return
} else {
let__5451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5) {
__typedN0, __typedOK0 := TypedFloat64(V3297)
__typedN1, __typedOK1 := TypedFloat64(V3298)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5) {
return TypedMaterializeBoolean((__typedN0 < __typedN1))
}}
__typedArg0 := V3297
__typedArg1 := V3298
return Call(__e, PrimFunc(sym_5), __typedArg0, __typedArg1)
})()
_ = let__5451

var ifres8072 Obj

if True == let__5451 {
tmp8074 := Call(__e, PrimFunc(symshen_4loading_2))


var ifres8073 Obj

if True == tmp8074 {
ifres8073 = True


} else {
ifres8073 = False


}

ifres8072 = ifres8073


} else {
ifres8072 = False


}

var ifres8062 Obj

if True == ifres8072 {
tmp8064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3298, MakeInteger(1))
}
__typedArg0 := V3298
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8063 Obj

if True == tmp8064 {
ifres8063 = MakeString("")


} else {
ifres8063 = MakeString("s")


}

tmp8065 := Call(__e, PrimFunc(symshen_4app), ifres8063, MakeString("\n"), symshen_4a)


tmp8067 := Call(__e, PrimFunc(symshen_4app), V3298, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" argument"))
__typedS1, __typedOK1 := TypedString(tmp8065)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" argument")
__typedArg1 := tmp8065
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp8069 := Call(__e, PrimFunc(symshen_4app), V3296, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" might not like "))
__typedS1, __typedOK1 := TypedString(tmp8067)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" might not like ")
__typedArg1 := tmp8067
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp8070 := Call(__e, PrimFunc(symstoutput))


tmp8071 := Call(__e, PrimFunc(sympr), tmp8069, tmp8070)


ifres8062 = tmp8071


} else {
ifres8062 = symshen_4skip


}

let__5452 := ifres8062
_ = let__5452

__e.Return(let__5451)
return


}


}, 3)

__e.TailApply(ns2_1set, symshen_4overapplication_2, tmp8060)
return




}, 0)

