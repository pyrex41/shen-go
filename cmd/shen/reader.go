package main

import . "github.com/pyrex41/shen-go/kl"

var ReaderMain = MakeNative(func(__e *ControlFlow) {
tmp5433 := MakeNative(func(__e *ControlFlow) {
V2560 := __e.Get(1)
_ = V2560
tmp5434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1file_1as_1bytelist) {
return PrimReadFileAsByteList(V2560)
}
__typedArg0 := V2560
return Call(__e, PrimFunc(symread_1file_1as_1bytelist), __typedArg0)
})()

let__5010 := tmp5434
_ = let__5010

tmp5435 := MakeNative(func(__e *ControlFlow) {
tmp5436 := MakeNative(func(__e *ControlFlow) {
Z2563 := __e.Get(1)
_ = Z2563
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2563)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp5436, let__5010)
return


}, 0)

tmp5437 := MakeNative(func(__e *ControlFlow) {
Z2564 := __e.Get(1)
_ = Z2564
tmp5438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dresidue_d)
}
__typedArg0 := symshen_4_dresidue_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4reader_1error), tmp5438)
return


}, 1)

tmp5439 := Call(__e, try_1catch, tmp5435, tmp5437)


let__5011 := tmp5439
_ = let__5011

tmp5440 := Call(__e, PrimFunc(symshen_4process_1sexprs), let__5011)


let__5012 := tmp5440
_ = let__5012

__e.Return(let__5012)
return


}, 1)

tmp5441 := Call(__e, ns2_1set, symread_1file, tmp5433)


_ = tmp5441

tmp5442 := MakeNative(func(__e *ControlFlow) {
V2566 := __e.Get(1)
_ = V2566
tmp5443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dmaximum_1print_1sequence_1size_d)
}
__typedArg0 := sym_dmaximum_1print_1sequence_1size_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp5444 := Call(__e, PrimFunc(symshen_4reader_1error_1message), tmp5443, MakeInteger(0), V2566)


tmp5446 := Call(__e, PrimFunc(symshen_4proc_1nl), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("reader error near here: "))
__typedS1, __typedOK1 := TypedString(tmp5444)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("reader error near here: ")
__typedArg1 := tmp5444
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp5446)
}
__typedArg0 := tmp5446
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

tmp5447 := Call(__e, ns2_1set, symshen_4reader_1error, tmp5442)


_ = tmp5447

tmp5448 := MakeNative(func(__e *ControlFlow) {
V2574 := __e.Get(1)
_ = V2574
V2575 := __e.Get(2)
_ = V2575
V2576 := __e.Get(3)
_ = V2576
tmp5459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2576)
}
__typedArg0 := Nil
__typedArg1 := V2576
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5459 {
__e.Return(MakeString(""))
return
} else {
tmp5457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V2574, V2575)
}
__typedArg0 := V2574
__typedArg1 := V2575
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5457 {
__e.Return(MakeString(""))
return
} else {
tmp5455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2576)
}
__typedArg0 := V2576
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5455 {
tmp5449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2576)
}
__typedArg0 := V2576
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(tmp5449)
}
__typedArg0 := tmp5449
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp5451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V2575)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V2575
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()

tmp5452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2576)
}
__typedArg0 := V2576
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5453 := Call(__e, PrimFunc(symshen_4reader_1error_1message), V2574, tmp5451, tmp5452)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp5450)
__typedS1, __typedOK1 := TypedString(tmp5453)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp5450
__typedArg1 := tmp5453
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

tmp5460 := Call(__e, ns2_1set, symshen_4reader_1error_1message, tmp5448)


_ = tmp5460

tmp5461 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dit_d)
}
__typedArg0 := symshen_4_dit_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp5462 := Call(__e, ns2_1set, symit, tmp5461)


_ = tmp5462

tmp5463 := MakeNative(func(__e *ControlFlow) {
V2577 := __e.Get(1)
_ = V2577
tmp5464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symopen) {
return PrimOpenStream(V2577, symin)
}
__typedArg0 := V2577
__typedArg1 := symin
return Call(__e, PrimFunc(symopen), __typedArg0, __typedArg1)
})()

let__5013 := tmp5464
_ = let__5013

tmp5465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(let__5013)
}
__typedArg0 := let__5013
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

let__5014 := tmp5465
_ = let__5014

tmp5466 := Call(__e, PrimFunc(symshen_4read_1file_1as_1bytelist_1help), let__5013, let__5014, Nil)


let__5015 := tmp5466
_ = let__5015

tmp5467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symclose) {
return PrimCloseStream(let__5013)
}
__typedArg0 := let__5013
return Call(__e, PrimFunc(symclose), __typedArg0)
})()

let__5016 := tmp5467
_ = let__5016

__e.TailApply(PrimFunc(symreverse), let__5015)
return


}, 1)

tmp5468 := Call(__e, ns2_1set, symread_1file_1as_1bytelist, tmp5463)


_ = tmp5468

tmp5469 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__selftop:
V2582 := __self1
_ = V2582
V2583 := __self2
_ = V2583
V2584 := __self3
_ = V2584
tmp5473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(-1), V2583)
}
__typedArg0 := MakeInteger(-1)
__typedArg1 := V2583
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5473 {
__e.Return(V2584)
return
} else {
tmp5470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(V2582)
}
__typedArg0 := V2582
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

tmp5471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2583, V2584)
}
__typedArg0 := V2583
__typedArg1 := V2584
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4read_1file_1as_1bytelist_1help) == __self {
__self1, __self2, __self3 = V2582, tmp5470, tmp5471
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4read_1file_1as_1bytelist_1help), V2582, tmp5470, tmp5471)
return


}


}, 3)

tmp5474 := Call(__e, ns2_1set, symshen_4read_1file_1as_1bytelist_1help, tmp5469)


_ = tmp5474

tmp5475 := MakeNative(func(__e *ControlFlow) {
V2585 := __e.Get(1)
_ = V2585
tmp5476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symopen) {
return PrimOpenStream(V2585, symin)
}
__typedArg0 := V2585
__typedArg1 := symin
return Call(__e, PrimFunc(symopen), __typedArg0, __typedArg1)
})()

let__5017 := tmp5476
_ = let__5017

tmp5477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(let__5017)
}
__typedArg0 := let__5017
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4rfas_1h), let__5017, tmp5477, MakeString(""))
return


}, 1)

tmp5478 := Call(__e, ns2_1set, symread_1file_1as_1string, tmp5475)


_ = tmp5478

tmp5479 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__selftop:
V2587 := __self1
_ = V2587
V2588 := __self2
_ = V2588
V2589 := __self3
_ = V2589
tmp5485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(-1), V2588)
}
__typedArg0 := MakeInteger(-1)
__typedArg1 := V2588
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5485 {
tmp5480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symclose) {
return PrimCloseStream(V2587)
}
__typedArg0 := V2587
return Call(__e, PrimFunc(symclose), __typedArg0)
})()

_ = tmp5480

__e.Return(V2589)
return


} else {
tmp5481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(V2587)
}
__typedArg0 := V2587
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

tmp5482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(V2588)
}
__typedArg0 := V2588
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

if PrimFunc(symshen_4rfas_1h) == __self {
__self1, __self2, __self3 = V2587, tmp5481, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(V2589)
__typedS1, __typedOK1 := TypedString(tmp5482)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V2589
__typedArg1 := tmp5482
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4rfas_1h), V2587, tmp5481, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(V2589)
__typedS1, __typedOK1 := TypedString(tmp5482)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := V2589
__typedArg1 := tmp5482
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}


}, 3)

tmp5486 := Call(__e, ns2_1set, symshen_4rfas_1h, tmp5479)


_ = tmp5486

tmp5487 := MakeNative(func(__e *ControlFlow) {
V2590 := __e.Get(1)
_ = V2590
tmp5488 := Call(__e, PrimFunc(symread), V2590)


__e.TailApply(PrimFunc(symeval_1kl), tmp5488)
return


}, 1)

tmp5489 := Call(__e, ns2_1set, syminput, tmp5487)


_ = tmp5489

tmp5490 := MakeNative(func(__e *ControlFlow) {
V2591 := __e.Get(1)
_ = V2591
V2592 := __e.Get(2)
_ = V2592
tmp5491 := Call(__e, PrimFunc(symshen_4monotype), V2591)


let__5018 := tmp5491
_ = let__5018

tmp5492 := Call(__e, PrimFunc(symread), V2592)


let__5019 := tmp5492
_ = let__5019

tmp5498 := Call(__e, PrimFunc(symshen_4typecheck), let__5019, V2591)


tmp5499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, tmp5498)
}
__typedArg0 := False
__typedArg1 := tmp5498
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5499 {
tmp5493 := Call(__e, PrimFunc(symshen_4app), V2591, MakeString("\n"), symshen_4r)


tmp5495 := Call(__e, PrimFunc(symshen_4app), let__5019, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" is not of type "))
__typedS1, __typedOK1 := TypedString(tmp5493)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" is not of type ")
__typedArg1 := tmp5493
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4r)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("type error: "))
__typedS1, __typedOK1 := TypedString(tmp5495)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("type error: ")
__typedArg1 := tmp5495
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("type error: "))
__typedS1, __typedOK1 := TypedString(tmp5495)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("type error: ")
__typedArg1 := tmp5495
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
__e.TailApply(PrimFunc(symeval_1kl), let__5019)
return
}


}, 2)

tmp5500 := Call(__e, ns2_1set, symshen_4input_1h_7, tmp5490)


_ = tmp5500

tmp5501 := MakeNative(func(__e *ControlFlow) {
V2595 := __e.Get(1)
_ = V2595
tmp5508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2595)
}
__typedArg0 := V2595
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5508 {
tmp5502 := MakeNative(func(__e *ControlFlow) {
Z2596 := __e.Get(1)
_ = Z2596
__e.TailApply(PrimFunc(symshen_4monotype), Z2596)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp5502, V2595)
return


} else {
tmp5506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V2595)
}
__typedArg0 := V2595
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp5506 {
tmp5503 := Call(__e, PrimFunc(symshen_4app), V2595, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("input+ expects a monotype: not "))
__typedS1, __typedOK1 := TypedString(tmp5503)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("input+ expects a monotype: not ")
__typedArg1 := tmp5503
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("input+ expects a monotype: not "))
__typedS1, __typedOK1 := TypedString(tmp5503)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("input+ expects a monotype: not ")
__typedArg1 := tmp5503
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
__e.Return(V2595)
return
}


}


}, 1)

tmp5509 := Call(__e, ns2_1set, symshen_4monotype, tmp5501)


_ = tmp5509

tmp5510 := MakeNative(func(__e *ControlFlow) {
V2597 := __e.Get(1)
_ = V2597
tmp5511 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2597)


tmp5512 := MakeNative(func(__e *ControlFlow) {
Z2598 := __e.Get(1)
_ = Z2598
__e.TailApply(PrimFunc(symshen_4return_2), Z2598)
return
}, 1)

__e.TailApply(PrimFunc(symshen_4read_1loop), V2597, tmp5511, Nil, tmp5512)
return


}, 1)

tmp5513 := Call(__e, ns2_1set, symlineread, tmp5510)


_ = tmp5513

tmp5514 := MakeNative(func(__e *ControlFlow) {
V2599 := __e.Get(1)
_ = V2599
tmp5515 := Call(__e, PrimFunc(symshen_4str_1_6bytes), V2599)


let__5020 := tmp5515
_ = let__5020

tmp5516 := MakeNative(func(__e *ControlFlow) {
Z2602 := __e.Get(1)
_ = Z2602
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2602)
return
}, 1)

tmp5517 := Call(__e, PrimFunc(symcompile), tmp5516, let__5020)


let__5021 := tmp5517
_ = let__5021

tmp5518 := Call(__e, PrimFunc(symshen_4process_1sexprs), let__5021)


let__5022 := tmp5518
_ = let__5022

__e.Return(let__5022)
return


}, 1)

tmp5519 := Call(__e, ns2_1set, symread_1from_1string, tmp5514)


_ = tmp5519

tmp5520 := MakeNative(func(__e *ControlFlow) {
V2604 := __e.Get(1)
_ = V2604
tmp5521 := Call(__e, PrimFunc(symshen_4str_1_6bytes), V2604)


let__5023 := tmp5521
_ = let__5023

tmp5522 := MakeNative(func(__e *ControlFlow) {
Z2607 := __e.Get(1)
_ = Z2607
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2607)
return
}, 1)

tmp5523 := Call(__e, PrimFunc(symcompile), tmp5522, let__5023)


let__5024 := tmp5523
_ = let__5024

__e.Return(let__5024)
return


}, 1)

tmp5524 := Call(__e, ns2_1set, symread_1from_1string_1unprocessed, tmp5520)


_ = tmp5524

tmp5525 := MakeNative(func(__e *ControlFlow) {
V2608 := __e.Get(1)
_ = V2608
tmp5533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V2608)
}
__typedArg0 := MakeString("")
__typedArg1 := V2608
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5533 {
__e.Return(Nil)
return
} else {
tmp5531 := Call(__e, PrimFunc(symshen_4_7string_2), V2608)


if True == tmp5531 {
tmp5526 := Call(__e, PrimFunc(symhdstr), V2608)


tmp5527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp5526)
}
__typedArg0 := tmp5526
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})()

tmp5529 := Call(__e, PrimFunc(symshen_4str_1_6bytes), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V2608)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V2608
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5527, tmp5529)
}
__typedArg0 := tmp5527
__typedArg1 := tmp5529
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4str_1_6bytes)
return
}


}


}, 1)

tmp5534 := Call(__e, ns2_1set, symshen_4str_1_6bytes, tmp5525)


_ = tmp5534

tmp5535 := MakeNative(func(__e *ControlFlow) {
V2609 := __e.Get(1)
_ = V2609
tmp5536 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2609)


tmp5537 := MakeNative(func(__e *ControlFlow) {
Z2610 := __e.Get(1)
_ = Z2610
__e.TailApply(PrimFunc(symshen_4whitespace_2), Z2610)
return
}, 1)

tmp5538 := Call(__e, PrimFunc(symshen_4read_1loop), V2609, tmp5536, Nil, tmp5537)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5538)
}
__typedArg0 := tmp5538
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


}, 1)

tmp5539 := Call(__e, ns2_1set, symread, tmp5535)


_ = tmp5539

tmp5540 := MakeNative(func(__e *ControlFlow) {
V2611 := __e.Get(1)
_ = V2611
tmp5543 := Call(__e, PrimFunc(symshen_4char_1stinput_2), V2611)


if True == tmp5543 {
tmp5541 := Call(__e, PrimFunc(symshen_4read_1unit_1string), V2611)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_1_6n) {
return PrimStringToNumber(tmp5541)
}
__typedArg0 := tmp5541
return Call(__e, PrimFunc(symstring_1_6n), __typedArg0)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(V2611)
}
__typedArg0 := V2611
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})())
return
}


}, 1)

tmp5544 := Call(__e, ns2_1set, symshen_4my_1read_1byte, tmp5540)


_ = tmp5544

tmp5545 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__selftop:
V2616 := __self1
_ = V2616
V2617 := __self2
_ = V2617
V2618 := __self3
_ = V2618
V2619 := __self4
_ = V2619
tmp5568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(94), V2617)
}
__typedArg0 := MakeInteger(94)
__typedArg1 := V2617
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5568 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("read aborted"))
}
__typedArg0 := MakeString("read aborted")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp5566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(-1), V2617)
}
__typedArg0 := MakeInteger(-1)
__typedArg1 := V2617
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5566 {
tmp5549 := Call(__e, PrimFunc(symempty_2), V2618)


if True == tmp5549 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("error: empty stream"))
}
__typedArg0 := MakeString("error: empty stream")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp5546 := MakeNative(func(__e *ControlFlow) {
Z2620 := __e.Get(1)
_ = Z2620
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2620)
return
}, 1)

tmp5547 := Call(__e, PrimFunc(symcompile), tmp5546, V2618)


__e.TailApply(PrimFunc(symshen_4process_1sexprs), tmp5547)
return


}


} else {
tmp5564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V2617)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V2617
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5564 {
tmp5550 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2616)


if PrimFunc(symshen_4read_1loop) == __self {
__self1, __self2, __self3, __self4 = V2616, tmp5550, V2618, V2619
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4read_1loop), V2616, tmp5550, V2618, V2619)
return


} else {
tmp5562 := Call(__e, V2619, V2617)


if True == tmp5562 {
tmp5551 := Call(__e, PrimFunc(symshen_4try_1parse), V2618)


let__5025 := tmp5551
_ = let__5025

tmp5557 := Call(__e, PrimFunc(symshen_4nothing_1doing_2), let__5025)


if True == tmp5557 {
tmp5552 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2616)


tmp5553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2617, Nil)
}
__typedArg0 := V2617
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5554 := Call(__e, PrimFunc(symappend), V2618, tmp5553)


if PrimFunc(symshen_4read_1loop) == __self {
__self1, __self2, __self3, __self4 = V2616, tmp5552, tmp5554, V2619
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4read_1loop), V2616, tmp5552, tmp5554, V2619)
return


} else {
tmp5555 := Call(__e, PrimFunc(symshen_4record_1it), V2618)


_ = tmp5555

__e.TailApply(PrimFunc(symshen_4process_1sexprs), let__5025)
return


}


} else {
tmp5558 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2616)


tmp5559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2617, Nil)
}
__typedArg0 := V2617
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5560 := Call(__e, PrimFunc(symappend), V2618, tmp5559)


if PrimFunc(symshen_4read_1loop) == __self {
__self1, __self2, __self3, __self4 = V2616, tmp5558, tmp5560, V2619
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4read_1loop), V2616, tmp5558, tmp5560, V2619)
return


}


}


}


}


}, 4)

tmp5569 := Call(__e, ns2_1set, symshen_4read_1loop, tmp5545)


_ = tmp5569

tmp5570 := MakeNative(func(__e *ControlFlow) {
V2622 := __e.Get(1)
_ = V2622
tmp5571 := MakeNative(func(__e *ControlFlow) {
tmp5572 := MakeNative(func(__e *ControlFlow) {
Z2623 := __e.Get(1)
_ = Z2623
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2623)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp5572, V2622)
return


}, 0)

tmp5573 := MakeNative(func(__e *ControlFlow) {
Z2624 := __e.Get(1)
_ = Z2624
__e.Return(symshen_4i_1failed_b)
return
}, 1)

__e.TailApply(try_1catch, tmp5571, tmp5573)
return


}, 1)

tmp5574 := Call(__e, ns2_1set, symshen_4try_1parse, tmp5570)


_ = tmp5574

tmp5575 := MakeNative(func(__e *ControlFlow) {
V2627 := __e.Get(1)
_ = V2627
tmp5579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4i_1failed_b, V2627)
}
__typedArg0 := symshen_4i_1failed_b
__typedArg1 := V2627
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5579 {
__e.Return(True)
return
} else {
tmp5577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2627)
}
__typedArg0 := Nil
__typedArg1 := V2627
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5577 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp5580 := Call(__e, ns2_1set, symshen_4nothing_1doing_2, tmp5575)


_ = tmp5580

tmp5581 := MakeNative(func(__e *ControlFlow) {
V2628 := __e.Get(1)
_ = V2628
tmp5582 := Call(__e, PrimFunc(symshen_4bytes_1_6string), V2628)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dit_d, tmp5582)
}
__typedArg0 := symshen_4_dit_d
__typedArg1 := tmp5582
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp5583 := Call(__e, ns2_1set, symshen_4record_1it, tmp5581)


_ = tmp5583

tmp5584 := MakeNative(func(__e *ControlFlow) {
V2629 := __e.Get(1)
_ = V2629
tmp5592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2629)
}
__typedArg0 := Nil
__typedArg1 := V2629
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5592 {
__e.Return(MakeString(""))
return
} else {
tmp5590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2629)
}
__typedArg0 := V2629
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5590 {
tmp5585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2629)
}
__typedArg0 := V2629
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(tmp5585)
}
__typedArg0 := tmp5585
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp5587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2629)
}
__typedArg0 := V2629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5588 := Call(__e, PrimFunc(symshen_4bytes_1_6string), tmp5587)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp5586)
__typedS1, __typedOK1 := TypedString(tmp5588)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp5586
__typedArg1 := tmp5588
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4bytes_1_6string)
return
}


}


}, 1)

tmp5593 := Call(__e, ns2_1set, symshen_4bytes_1_6string, tmp5584)


_ = tmp5593

tmp5594 := MakeNative(func(__e *ControlFlow) {
V2630 := __e.Get(1)
_ = V2630
tmp5595 := Call(__e, PrimFunc(symshen_4unpackage_emacroexpand), V2630)


let__5026 := tmp5595
_ = let__5026

tmp5596 := Call(__e, PrimFunc(symshen_4find_1arities), let__5026)


let__5027 := tmp5596
_ = let__5027

tmp5597 := Call(__e, PrimFunc(symshen_4find_1types), let__5026)


let__5028 := tmp5597
_ = let__5028

tmp5598 := MakeNative(func(__e *ControlFlow) {
Z2634 := __e.Get(1)
_ = Z2634
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2634, let__5028)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp5598, let__5026)
return


}, 1)

tmp5599 := Call(__e, ns2_1set, symshen_4process_1sexprs, tmp5594)


_ = tmp5599

tmp5600 := MakeNative(func(__e *ControlFlow) {
V2635 := __e.Get(1)
_ = V2635
tmp5622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2635)
}
__typedArg0 := V2635
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5613 Obj

if True == tmp5622 {
tmp5620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2635)
}
__typedArg0 := V2635
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5620)
}
__typedArg0 := tmp5620
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5615 Obj

if True == tmp5621 {
tmp5617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2635)
}
__typedArg0 := V2635
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp5619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp5617, tmp5618)
}
__typedArg0 := tmp5617
__typedArg1 := tmp5618
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
tmp5601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2635)
}
__typedArg0 := V2635
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5601)
}
__typedArg0 := tmp5601
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2635)
}
__typedArg0 := V2635
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5603)
}
__typedArg0 := tmp5603
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5605 := Call(__e, PrimFunc(symshen_4find_1types), tmp5604)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5602, tmp5605)
}
__typedArg0 := tmp5602
__typedArg1 := tmp5605
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp5611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2635)
}
__typedArg0 := V2635
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5611 {
tmp5606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2635)
}
__typedArg0 := V2635
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5607 := Call(__e, PrimFunc(symshen_4find_1types), tmp5606)


tmp5608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2635)
}
__typedArg0 := V2635
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5609 := Call(__e, PrimFunc(symshen_4find_1types), tmp5608)


__e.TailApply(PrimFunc(symappend), tmp5607, tmp5609)
return


} else {
__e.Return(Nil)
return
}


}


}, 1)

tmp5623 := Call(__e, ns2_1set, symshen_4find_1types, tmp5600)


_ = tmp5623

tmp5624 := MakeNative(func(__e *ControlFlow) {
V2638 := __e.Get(1)
_ = V2638
tmp5673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2638)
}
__typedArg0 := V2638
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5654 Obj

if True == tmp5673 {
tmp5671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2638)
}
__typedArg0 := V2638
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp5671)
}
__typedArg0 := symdefine
__typedArg1 := tmp5671
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5656 Obj

if True == tmp5672 {
tmp5669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2638)
}
__typedArg0 := V2638
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5669)
}
__typedArg0 := tmp5669
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5658 Obj

if True == tmp5670 {
tmp5666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2638)
}
__typedArg0 := V2638
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5666)
}
__typedArg0 := tmp5666
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5667)
}
__typedArg0 := tmp5667
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5660 Obj

if True == tmp5668 {
tmp5662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2638)
}
__typedArg0 := V2638
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5662)
}
__typedArg0 := tmp5662
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5663)
}
__typedArg0 := tmp5663
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, tmp5664)
}
__typedArg0 := sym_i
__typedArg1 := tmp5664
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
tmp5625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2638)
}
__typedArg0 := V2638
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5625)
}
__typedArg0 := tmp5625
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2638)
}
__typedArg0 := V2638
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5627)
}
__typedArg0 := tmp5627
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2638)
}
__typedArg0 := V2638
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5629)
}
__typedArg0 := tmp5629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5630)
}
__typedArg0 := tmp5630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5632 := Call(__e, PrimFunc(symshen_4find_1arity), tmp5628, MakeInteger(1), tmp5631)


__e.TailApply(PrimFunc(symshen_4store_1arity), tmp5626, tmp5632)
return


} else {
tmp5652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2638)
}
__typedArg0 := V2638
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5644 Obj

if True == tmp5652 {
tmp5650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2638)
}
__typedArg0 := V2638
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp5650)
}
__typedArg0 := symdefine
__typedArg1 := tmp5650
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5646 Obj

if True == tmp5651 {
tmp5648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2638)
}
__typedArg0 := V2638
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5648)
}
__typedArg0 := tmp5648
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

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
tmp5633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2638)
}
__typedArg0 := V2638
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5633)
}
__typedArg0 := tmp5633
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2638)
}
__typedArg0 := V2638
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5635)
}
__typedArg0 := tmp5635
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2638)
}
__typedArg0 := V2638
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp5637)
}
__typedArg0 := tmp5637
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5639 := Call(__e, PrimFunc(symshen_4find_1arity), tmp5636, MakeInteger(0), tmp5638)


__e.TailApply(PrimFunc(symshen_4store_1arity), tmp5634, tmp5639)
return


} else {
tmp5642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2638)
}
__typedArg0 := V2638
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5642 {
tmp5640 := MakeNative(func(__e *ControlFlow) {
Z2639 := __e.Get(1)
_ = Z2639
__e.TailApply(PrimFunc(symshen_4find_1arities), Z2639)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp5640, V2638)
return


} else {
__e.Return(symshen_4skip)
return
}


}


}


}, 1)

tmp5674 := Call(__e, ns2_1set, symshen_4find_1arities, tmp5624)


_ = tmp5674

tmp5675 := MakeNative(func(__e *ControlFlow) {
V2640 := __e.Get(1)
_ = V2640
V2641 := __e.Get(2)
_ = V2641
tmp5676 := Call(__e, PrimFunc(symarity), V2640)


let__5029 := tmp5676
_ = let__5029

tmp5687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__5029, MakeInteger(-1))
}
__typedArg0 := let__5029
__typedArg1 := MakeInteger(-1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5687 {
__e.TailApply(PrimFunc(symshen_4execute_1store_1arity), V2640, V2641)
return
} else {
tmp5685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__5029, V2641)
}
__typedArg0 := let__5029
__typedArg1 := V2641
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5685 {
__e.Return(symshen_4skip)
return
} else {
tmp5683 := Call(__e, PrimFunc(symshen_4sysfunc_2), V2640)


if True == tmp5683 {
tmp5677 := Call(__e, PrimFunc(symshen_4app), V2640, MakeString(" is a system function\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp5677)
}
__typedArg0 := tmp5677
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
tmp5678 := Call(__e, PrimFunc(symshen_4app), V2640, MakeString(" may cause errors\n"), symshen_4a)


tmp5679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("changing the arity of "))
__typedS1, __typedOK1 := TypedString(tmp5678)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("changing the arity of ")
__typedArg1 := tmp5678
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp5680 := Call(__e, PrimFunc(symstoutput))


tmp5681 := Call(__e, PrimFunc(sympr), tmp5679, tmp5680)


_ = tmp5681

__e.TailApply(PrimFunc(symshen_4execute_1store_1arity), V2640, V2641)
return


}


}


}


}, 2)

tmp5688 := Call(__e, ns2_1set, symshen_4store_1arity, tmp5675)


_ = tmp5688

tmp5689 := MakeNative(func(__e *ControlFlow) {
V2643 := __e.Get(1)
_ = V2643
V2644 := __e.Get(2)
_ = V2644
tmp5694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V2644)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V2644
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5694 {
tmp5690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V2643, symarity, MakeInteger(0), tmp5690)
return


} else {
tmp5691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp5692 := Call(__e, PrimFunc(symput), V2643, symarity, V2644, tmp5691)


_ = tmp5692

__e.TailApply(PrimFunc(symshen_4update_1lambdatable), V2643, V2644)
return


}


}, 2)

tmp5695 := Call(__e, ns2_1set, symshen_4execute_1store_1arity, tmp5689)


_ = tmp5695

tmp5696 := MakeNative(func(__e *ControlFlow) {
V2645 := __e.Get(1)
_ = V2645
V2646 := __e.Get(2)
_ = V2646
tmp5697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2645, Nil)
}
__typedArg0 := V2645
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5698 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp5697, V2646)


tmp5699 := Call(__e, PrimFunc(symeval_1kl), tmp5698)


let__5030 := tmp5699
_ = let__5030

tmp5700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2645, let__5030)
}
__typedArg0 := V2645
__typedArg1 := let__5030
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5701 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp5700)


let__5031 := tmp5701
_ = let__5031

__e.Return(V2646)
return


}, 2)

tmp5702 := Call(__e, ns2_1set, symshen_4update_1lambdatable, tmp5696)


_ = tmp5702

tmp5703 := MakeNative(func(__e *ControlFlow) {
V2651 := __e.Get(1)
_ = V2651
V2652 := __e.Get(2)
_ = V2652
tmp5719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V2652)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V2652
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5719 {
__e.Return(symshen_4skip)
return
} else {
tmp5717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(1), V2652)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := V2652
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5717 {
tmp5704 := Call(__e, PrimFunc(symgensym), symY)


let__5032 := tmp5704
_ = let__5032

tmp5705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__5032, Nil)
}
__typedArg0 := let__5032
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5706 := Call(__e, PrimFunc(symappend), V2651, tmp5705)


tmp5707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5706, Nil)
}
__typedArg0 := tmp5706
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__5032, tmp5707)
}
__typedArg0 := let__5032
__typedArg1 := tmp5707
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp5708)
}
__typedArg0 := symlambda
__typedArg1 := tmp5708
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp5709 := Call(__e, PrimFunc(symgensym), symY)


let__5033 := tmp5709
_ = let__5033

tmp5710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__5033, Nil)
}
__typedArg0 := let__5033
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5711 := Call(__e, PrimFunc(symappend), V2651, tmp5710)


tmp5713 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp5711, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V2652)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V2652
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


tmp5714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5713, Nil)
}
__typedArg0 := tmp5713
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__5033, tmp5714)
}
__typedArg0 := let__5033
__typedArg1 := tmp5714
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp5715)
}
__typedArg0 := symlambda
__typedArg1 := tmp5715
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}


}, 2)

tmp5720 := Call(__e, ns2_1set, symshen_4lambda_1function, tmp5703)


_ = tmp5720

tmp5721 := MakeNative(func(__e *ControlFlow) {
V2664 := __e.Get(1)
_ = V2664
V2665 := __e.Get(2)
_ = V2665
V2666 := __e.Get(3)
_ = V2666
tmp5744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2666)
}
__typedArg0 := Nil
__typedArg1 := V2666
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5744 {
tmp5722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2664, V2665)
}
__typedArg0 := V2664
__typedArg1 := V2665
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5722, Nil)
}
__typedArg0 := tmp5722
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp5742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2666)
}
__typedArg0 := V2666
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5733 Obj

if True == tmp5742 {
tmp5740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2666)
}
__typedArg0 := V2666
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp5740)
}
__typedArg0 := tmp5740
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5735 Obj

if True == tmp5741 {
tmp5737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2666)
}
__typedArg0 := V2666
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5737)
}
__typedArg0 := tmp5737
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V2664, tmp5738)
}
__typedArg0 := V2664
__typedArg1 := tmp5738
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
tmp5723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2666)
}
__typedArg0 := V2666
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp5723)
}
__typedArg0 := tmp5723
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5724, V2665)
}
__typedArg0 := tmp5724
__typedArg1 := V2665
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2666)
}
__typedArg0 := V2666
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5725, tmp5726)
}
__typedArg0 := tmp5725
__typedArg1 := tmp5726
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp5731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2666)
}
__typedArg0 := V2666
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5731 {
tmp5727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2666)
}
__typedArg0 := V2666
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2666)
}
__typedArg0 := V2666
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5729 := Call(__e, PrimFunc(symshen_4assoc_1_6), V2664, V2665, tmp5728)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5727, tmp5729)
}
__typedArg0 := tmp5727
__typedArg1 := tmp5729
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

tmp5745 := Call(__e, ns2_1set, symshen_4assoc_1_6, tmp5721)


_ = tmp5745

tmp5746 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__selftop:
V2681 := __self1
_ = V2681
V2682 := __self2
_ = V2682
V2683 := __self3
_ = V2683
tmp5793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V2682)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V2682
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5786 Obj

if True == tmp5793 {
tmp5792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2683)
}
__typedArg0 := V2683
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5788 Obj

if True == tmp5792 {
tmp5790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2683)
}
__typedArg0 := V2683
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp5790, sym_1_6)
}
__typedArg0 := tmp5790
__typedArg1 := sym_1_6
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
__e.Return(MakeInteger(0))
return
} else {
tmp5784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V2682)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V2682
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5777 Obj

if True == tmp5784 {
tmp5783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2683)
}
__typedArg0 := V2683
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5779 Obj

if True == tmp5783 {
tmp5781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2683)
}
__typedArg0 := V2683
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp5781, sym_5_1)
}
__typedArg0 := tmp5781
__typedArg1 := sym_5_1
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
__e.Return(MakeInteger(0))
return
} else {
tmp5775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V2682)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V2682
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5772 Obj

if True == tmp5775 {
tmp5774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2683)
}
__typedArg0 := V2683
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

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
tmp5747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2683)
}
__typedArg0 := V2683
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5748 := Call(__e, PrimFunc(symshen_4find_1arity), V2681, MakeInteger(0), tmp5747)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp5748)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp5748
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


} else {
tmp5770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(1), V2682)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := V2682
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5763 Obj

if True == tmp5770 {
tmp5769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2683)
}
__typedArg0 := V2683
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres5765 Obj

if True == tmp5769 {
tmp5767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2683)
}
__typedArg0 := V2683
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, tmp5767)
}
__typedArg0 := sym_j
__typedArg1 := tmp5767
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
tmp5749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2683)
}
__typedArg0 := V2683
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4find_1arity) == __self {
__self1, __self2, __self3 = V2681, MakeInteger(0), tmp5749
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4find_1arity), V2681, MakeInteger(0), tmp5749)
return


} else {
tmp5761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(1), V2682)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := V2682
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres5758 Obj

if True == tmp5761 {
tmp5760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2683)
}
__typedArg0 := V2683
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

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
tmp5750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2683)
}
__typedArg0 := V2683
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4find_1arity) == __self {
__self1, __self2, __self3 = V2681, MakeInteger(1), tmp5750
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4find_1arity), V2681, MakeInteger(1), tmp5750)
return


} else {
tmp5756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(1), V2682)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := V2682
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5756 {
tmp5751 := Call(__e, PrimFunc(symshen_4app), V2681, MakeString(" definition: missing }\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("syntax error in "))
__typedS1, __typedOK1 := TypedString(tmp5751)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("syntax error in ")
__typedArg1 := tmp5751
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("syntax error in "))
__typedS1, __typedOK1 := TypedString(tmp5751)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("syntax error in ")
__typedArg1 := tmp5751
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
tmp5753 := Call(__e, PrimFunc(symshen_4app), V2681, MakeString(" definition: missing -> or <-\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("syntax error in "))
__typedS1, __typedOK1 := TypedString(tmp5753)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("syntax error in ")
__typedArg1 := tmp5753
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("syntax error in "))
__typedS1, __typedOK1 := TypedString(tmp5753)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("syntax error in ")
__typedArg1 := tmp5753
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

tmp5794 := Call(__e, ns2_1set, symshen_4find_1arity, tmp5746)


_ = tmp5794

tmp5795 := MakeNative(func(__e *ControlFlow) {
V2684 := __e.Get(1)
_ = V2684
tmp5796 := Call(__e, PrimFunc(symshen_4_5lsb_6), V2684)


let__5035 := tmp5796
_ = let__5035

tmp5820 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5035)


var ifres5797 Obj

if True == tmp5820 {
tmp5798 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5797 = tmp5798


} else {
tmp5799 := Call(__e, PrimFunc(symshen_4in_1_6), let__5035)


let__5036 := tmp5799
_ = let__5036

tmp5800 := Call(__e, PrimFunc(symshen_4_5s_1exprs1_6), let__5036)


let__5037 := tmp5800
_ = let__5037

tmp5819 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5037)


var ifres5801 Obj

if True == tmp5819 {
tmp5802 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5801 = tmp5802


} else {
tmp5803 := Call(__e, PrimFunc(symshen_4_5_1out), let__5037)


let__5038 := tmp5803
_ = let__5038

tmp5804 := Call(__e, PrimFunc(symshen_4in_1_6), let__5037)


let__5039 := tmp5804
_ = let__5039

tmp5805 := Call(__e, PrimFunc(symshen_4_5rsb_6), let__5039)


let__5040 := tmp5805
_ = let__5040

tmp5818 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5040)


var ifres5806 Obj

if True == tmp5818 {
tmp5807 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5806 = tmp5807


} else {
tmp5808 := Call(__e, PrimFunc(symshen_4in_1_6), let__5040)


let__5041 := tmp5808
_ = let__5041

tmp5809 := Call(__e, PrimFunc(symshen_4_5s_1exprs2_6), let__5041)


let__5042 := tmp5809
_ = let__5042

tmp5817 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5042)


var ifres5810 Obj

if True == tmp5817 {
tmp5811 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5810 = tmp5811


} else {
tmp5812 := Call(__e, PrimFunc(symshen_4_5_1out), let__5042)


let__5043 := tmp5812
_ = let__5043

tmp5813 := Call(__e, PrimFunc(symshen_4in_1_6), let__5042)


let__5044 := tmp5813
_ = let__5044

tmp5814 := Call(__e, PrimFunc(symshen_4cons_1form), let__5038)


tmp5815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5814, let__5043)
}
__typedArg0 := tmp5814
__typedArg1 := let__5043
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5816 := Call(__e, PrimFunc(symshen_4comb), let__5044, tmp5815)


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

tmp6014 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5034)


if True == tmp6014 {
tmp5821 := Call(__e, PrimFunc(symshen_4_5lrb_6), V2684)


let__5046 := tmp5821
_ = let__5046

tmp5844 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5046)


var ifres5822 Obj

if True == tmp5844 {
tmp5823 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5822 = tmp5823


} else {
tmp5824 := Call(__e, PrimFunc(symshen_4in_1_6), let__5046)


let__5047 := tmp5824
_ = let__5047

tmp5825 := Call(__e, PrimFunc(symshen_4_5s_1exprs1_6), let__5047)


let__5048 := tmp5825
_ = let__5048

tmp5843 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5048)


var ifres5826 Obj

if True == tmp5843 {
tmp5827 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5826 = tmp5827


} else {
tmp5828 := Call(__e, PrimFunc(symshen_4_5_1out), let__5048)


let__5049 := tmp5828
_ = let__5049

tmp5829 := Call(__e, PrimFunc(symshen_4in_1_6), let__5048)


let__5050 := tmp5829
_ = let__5050

tmp5830 := Call(__e, PrimFunc(symshen_4_5rrb_6), let__5050)


let__5051 := tmp5830
_ = let__5051

tmp5842 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5051)


var ifres5831 Obj

if True == tmp5842 {
tmp5832 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5831 = tmp5832


} else {
tmp5833 := Call(__e, PrimFunc(symshen_4in_1_6), let__5051)


let__5052 := tmp5833
_ = let__5052

tmp5834 := Call(__e, PrimFunc(symshen_4_5s_1exprs2_6), let__5052)


let__5053 := tmp5834
_ = let__5053

tmp5841 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5053)


var ifres5835 Obj

if True == tmp5841 {
tmp5836 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5835 = tmp5836


} else {
tmp5837 := Call(__e, PrimFunc(symshen_4_5_1out), let__5053)


let__5054 := tmp5837
_ = let__5054

tmp5838 := Call(__e, PrimFunc(symshen_4in_1_6), let__5053)


let__5055 := tmp5838
_ = let__5055

tmp5839 := Call(__e, PrimFunc(symshen_4add_1sexpr), let__5049, let__5054)


tmp5840 := Call(__e, PrimFunc(symshen_4comb), let__5055, tmp5839)


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

tmp6012 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5045)


if True == tmp6012 {
tmp5845 := Call(__e, PrimFunc(symshen_4_5lcurly_6), V2684)


let__5057 := tmp5845
_ = let__5057

tmp5857 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5057)


var ifres5846 Obj

if True == tmp5857 {
tmp5847 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5846 = tmp5847


} else {
tmp5848 := Call(__e, PrimFunc(symshen_4in_1_6), let__5057)


let__5058 := tmp5848
_ = let__5058

tmp5849 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__5058)


let__5059 := tmp5849
_ = let__5059

tmp5856 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5059)


var ifres5850 Obj

if True == tmp5856 {
tmp5851 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5850 = tmp5851


} else {
tmp5852 := Call(__e, PrimFunc(symshen_4_5_1out), let__5059)


let__5060 := tmp5852
_ = let__5060

tmp5853 := Call(__e, PrimFunc(symshen_4in_1_6), let__5059)


let__5061 := tmp5853
_ = let__5061

tmp5854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, let__5060)
}
__typedArg0 := sym_i
__typedArg1 := let__5060
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5855 := Call(__e, PrimFunc(symshen_4comb), let__5061, tmp5854)


ifres5850 = tmp5855


}

ifres5846 = ifres5850


}

let__5056 := ifres5846
_ = let__5056

tmp6010 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5056)


if True == tmp6010 {
tmp5858 := Call(__e, PrimFunc(symshen_4_5rcurly_6), V2684)


let__5063 := tmp5858
_ = let__5063

tmp5870 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5063)


var ifres5859 Obj

if True == tmp5870 {
tmp5860 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5859 = tmp5860


} else {
tmp5861 := Call(__e, PrimFunc(symshen_4in_1_6), let__5063)


let__5064 := tmp5861
_ = let__5064

tmp5862 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__5064)


let__5065 := tmp5862
_ = let__5065

tmp5869 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5065)


var ifres5863 Obj

if True == tmp5869 {
tmp5864 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5863 = tmp5864


} else {
tmp5865 := Call(__e, PrimFunc(symshen_4_5_1out), let__5065)


let__5066 := tmp5865
_ = let__5066

tmp5866 := Call(__e, PrimFunc(symshen_4in_1_6), let__5065)


let__5067 := tmp5866
_ = let__5067

tmp5867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, let__5066)
}
__typedArg0 := sym_j
__typedArg1 := let__5066
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5868 := Call(__e, PrimFunc(symshen_4comb), let__5067, tmp5867)


ifres5863 = tmp5868


}

ifres5859 = ifres5863


}

let__5062 := ifres5859
_ = let__5062

tmp6008 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5062)


if True == tmp6008 {
tmp5871 := Call(__e, PrimFunc(symshen_4_5bar_6), V2684)


let__5069 := tmp5871
_ = let__5069

tmp5883 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5069)


var ifres5872 Obj

if True == tmp5883 {
tmp5873 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5872 = tmp5873


} else {
tmp5874 := Call(__e, PrimFunc(symshen_4in_1_6), let__5069)


let__5070 := tmp5874
_ = let__5070

tmp5875 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__5070)


let__5071 := tmp5875
_ = let__5071

tmp5882 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5071)


var ifres5876 Obj

if True == tmp5882 {
tmp5877 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5876 = tmp5877


} else {
tmp5878 := Call(__e, PrimFunc(symshen_4_5_1out), let__5071)


let__5072 := tmp5878
_ = let__5072

tmp5879 := Call(__e, PrimFunc(symshen_4in_1_6), let__5071)


let__5073 := tmp5879
_ = let__5073

tmp5880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbar_b, let__5072)
}
__typedArg0 := symbar_b
__typedArg1 := let__5072
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5881 := Call(__e, PrimFunc(symshen_4comb), let__5073, tmp5880)


ifres5876 = tmp5881


}

ifres5872 = ifres5876


}

let__5068 := ifres5872
_ = let__5068

tmp6006 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5068)


if True == tmp6006 {
tmp5884 := Call(__e, PrimFunc(symshen_4_5semicolon_6), V2684)


let__5075 := tmp5884
_ = let__5075

tmp5897 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5075)


var ifres5885 Obj

if True == tmp5897 {
tmp5886 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5885 = tmp5886


} else {
tmp5887 := Call(__e, PrimFunc(symshen_4in_1_6), let__5075)


let__5076 := tmp5887
_ = let__5076

tmp5888 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__5076)


let__5077 := tmp5888
_ = let__5077

tmp5896 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5077)


var ifres5889 Obj

if True == tmp5896 {
tmp5890 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5889 = tmp5890


} else {
tmp5891 := Call(__e, PrimFunc(symshen_4_5_1out), let__5077)


let__5078 := tmp5891
_ = let__5078

tmp5892 := Call(__e, PrimFunc(symshen_4in_1_6), let__5077)


let__5079 := tmp5892
_ = let__5079

tmp5893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp5894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5893, let__5078)
}
__typedArg0 := tmp5893
__typedArg1 := let__5078
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5895 := Call(__e, PrimFunc(symshen_4comb), let__5079, tmp5894)


ifres5889 = tmp5895


}

ifres5885 = ifres5889


}

let__5074 := ifres5885
_ = let__5074

tmp6004 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5074)


if True == tmp6004 {
tmp5898 := Call(__e, PrimFunc(symshen_4_5colon_6), V2684)


let__5081 := tmp5898
_ = let__5081

tmp5916 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5081)


var ifres5899 Obj

if True == tmp5916 {
tmp5900 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5899 = tmp5900


} else {
tmp5901 := Call(__e, PrimFunc(symshen_4in_1_6), let__5081)


let__5082 := tmp5901
_ = let__5082

tmp5902 := Call(__e, PrimFunc(symshen_4_5equal_6), let__5082)


let__5083 := tmp5902
_ = let__5083

tmp5915 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5083)


var ifres5903 Obj

if True == tmp5915 {
tmp5904 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5903 = tmp5904


} else {
tmp5905 := Call(__e, PrimFunc(symshen_4in_1_6), let__5083)


let__5084 := tmp5905
_ = let__5084

tmp5906 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__5084)


let__5085 := tmp5906
_ = let__5085

tmp5914 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5085)


var ifres5907 Obj

if True == tmp5914 {
tmp5908 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5907 = tmp5908


} else {
tmp5909 := Call(__e, PrimFunc(symshen_4_5_1out), let__5085)


let__5086 := tmp5909
_ = let__5086

tmp5910 := Call(__e, PrimFunc(symshen_4in_1_6), let__5085)


let__5087 := tmp5910
_ = let__5087

tmp5911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":="))
}
__typedArg0 := MakeString(":=")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp5912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5911, let__5086)
}
__typedArg0 := tmp5911
__typedArg1 := let__5086
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5913 := Call(__e, PrimFunc(symshen_4comb), let__5087, tmp5912)


ifres5907 = tmp5913


}

ifres5903 = ifres5907


}

ifres5899 = ifres5903


}

let__5080 := ifres5899
_ = let__5080

tmp6002 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5080)


if True == tmp6002 {
tmp5917 := Call(__e, PrimFunc(symshen_4_5colon_6), V2684)


let__5089 := tmp5917
_ = let__5089

tmp5930 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5089)


var ifres5918 Obj

if True == tmp5930 {
tmp5919 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5918 = tmp5919


} else {
tmp5920 := Call(__e, PrimFunc(symshen_4in_1_6), let__5089)


let__5090 := tmp5920
_ = let__5090

tmp5921 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__5090)


let__5091 := tmp5921
_ = let__5091

tmp5929 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5091)


var ifres5922 Obj

if True == tmp5929 {
tmp5923 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5922 = tmp5923


} else {
tmp5924 := Call(__e, PrimFunc(symshen_4_5_1out), let__5091)


let__5092 := tmp5924
_ = let__5092

tmp5925 := Call(__e, PrimFunc(symshen_4in_1_6), let__5091)


let__5093 := tmp5925
_ = let__5093

tmp5926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp5927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5926, let__5092)
}
__typedArg0 := tmp5926
__typedArg1 := let__5092
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5928 := Call(__e, PrimFunc(symshen_4comb), let__5093, tmp5927)


ifres5922 = tmp5928


}

ifres5918 = ifres5922


}

let__5088 := ifres5918
_ = let__5088

tmp6000 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5088)


if True == tmp6000 {
tmp5931 := Call(__e, PrimFunc(symshen_4_5comma_6), V2684)


let__5095 := tmp5931
_ = let__5095

tmp5944 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5095)


var ifres5932 Obj

if True == tmp5944 {
tmp5933 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5932 = tmp5933


} else {
tmp5934 := Call(__e, PrimFunc(symshen_4in_1_6), let__5095)


let__5096 := tmp5934
_ = let__5096

tmp5935 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__5096)


let__5097 := tmp5935
_ = let__5097

tmp5943 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5097)


var ifres5936 Obj

if True == tmp5943 {
tmp5937 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5936 = tmp5937


} else {
tmp5938 := Call(__e, PrimFunc(symshen_4_5_1out), let__5097)


let__5098 := tmp5938
_ = let__5098

tmp5939 := Call(__e, PrimFunc(symshen_4in_1_6), let__5097)


let__5099 := tmp5939
_ = let__5099

tmp5940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(","))
}
__typedArg0 := MakeString(",")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp5941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5940, let__5098)
}
__typedArg0 := tmp5940
__typedArg1 := let__5098
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5942 := Call(__e, PrimFunc(symshen_4comb), let__5099, tmp5941)


ifres5936 = tmp5942


}

ifres5932 = ifres5936


}

let__5094 := ifres5932
_ = let__5094

tmp5998 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5094)


if True == tmp5998 {
tmp5945 := Call(__e, PrimFunc(symshen_4_5comment_6), V2684)


let__5101 := tmp5945
_ = let__5101

tmp5956 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5101)


var ifres5946 Obj

if True == tmp5956 {
tmp5947 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5946 = tmp5947


} else {
tmp5948 := Call(__e, PrimFunc(symshen_4in_1_6), let__5101)


let__5102 := tmp5948
_ = let__5102

tmp5949 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__5102)


let__5103 := tmp5949
_ = let__5103

tmp5955 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5103)


var ifres5950 Obj

if True == tmp5955 {
tmp5951 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5950 = tmp5951


} else {
tmp5952 := Call(__e, PrimFunc(symshen_4_5_1out), let__5103)


let__5104 := tmp5952
_ = let__5104

tmp5953 := Call(__e, PrimFunc(symshen_4in_1_6), let__5103)


let__5105 := tmp5953
_ = let__5105

tmp5954 := Call(__e, PrimFunc(symshen_4comb), let__5105, let__5104)


ifres5950 = tmp5954


}

ifres5946 = ifres5950


}

let__5100 := ifres5946
_ = let__5100

tmp5996 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5100)


if True == tmp5996 {
tmp5957 := Call(__e, PrimFunc(symshen_4_5atom_6), V2684)


let__5107 := tmp5957
_ = let__5107

tmp5970 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5107)


var ifres5958 Obj

if True == tmp5970 {
tmp5959 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5958 = tmp5959


} else {
tmp5960 := Call(__e, PrimFunc(symshen_4_5_1out), let__5107)


let__5108 := tmp5960
_ = let__5108

tmp5961 := Call(__e, PrimFunc(symshen_4in_1_6), let__5107)


let__5109 := tmp5961
_ = let__5109

tmp5962 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__5109)


let__5110 := tmp5962
_ = let__5110

tmp5969 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5110)


var ifres5963 Obj

if True == tmp5969 {
tmp5964 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5963 = tmp5964


} else {
tmp5965 := Call(__e, PrimFunc(symshen_4_5_1out), let__5110)


let__5111 := tmp5965
_ = let__5111

tmp5966 := Call(__e, PrimFunc(symshen_4in_1_6), let__5110)


let__5112 := tmp5966
_ = let__5112

tmp5967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__5108, let__5111)
}
__typedArg0 := let__5108
__typedArg1 := let__5111
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5968 := Call(__e, PrimFunc(symshen_4comb), let__5112, tmp5967)


ifres5963 = tmp5968


}

ifres5958 = ifres5963


}

let__5106 := ifres5958
_ = let__5106

tmp5994 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5106)


if True == tmp5994 {
tmp5971 := Call(__e, PrimFunc(symshen_4_5whitespaces_6), V2684)


let__5114 := tmp5971
_ = let__5114

tmp5982 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5114)


var ifres5972 Obj

if True == tmp5982 {
tmp5973 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5972 = tmp5973


} else {
tmp5974 := Call(__e, PrimFunc(symshen_4in_1_6), let__5114)


let__5115 := tmp5974
_ = let__5115

tmp5975 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), let__5115)


let__5116 := tmp5975
_ = let__5116

tmp5981 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5116)


var ifres5976 Obj

if True == tmp5981 {
tmp5977 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5976 = tmp5977


} else {
tmp5978 := Call(__e, PrimFunc(symshen_4_5_1out), let__5116)


let__5117 := tmp5978
_ = let__5117

tmp5979 := Call(__e, PrimFunc(symshen_4in_1_6), let__5116)


let__5118 := tmp5979
_ = let__5118

tmp5980 := Call(__e, PrimFunc(symshen_4comb), let__5118, let__5117)


ifres5976 = tmp5980


}

ifres5972 = ifres5976


}

let__5113 := ifres5972
_ = let__5113

tmp5992 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5113)


if True == tmp5992 {
tmp5983 := Call(__e, PrimFunc(sym_5e_6), V2684)


let__5120 := tmp5983
_ = let__5120

tmp5988 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5120)


var ifres5984 Obj

if True == tmp5988 {
tmp5985 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5984 = tmp5985


} else {
tmp5986 := Call(__e, PrimFunc(symshen_4in_1_6), let__5120)


let__5121 := tmp5986
_ = let__5121

tmp5987 := Call(__e, PrimFunc(symshen_4comb), let__5121, Nil)


ifres5984 = tmp5987


}

let__5119 := ifres5984
_ = let__5119

tmp5990 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5119)


if True == tmp5990 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5119)
return
}


} else {
__e.Return(let__5113)
return
}


} else {
__e.Return(let__5106)
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
__e.Return(let__5080)
return
}


} else {
__e.Return(let__5074)
return
}


} else {
__e.Return(let__5068)
return
}


} else {
__e.Return(let__5062)
return
}


} else {
__e.Return(let__5056)
return
}


} else {
__e.Return(let__5045)
return
}


} else {
__e.Return(let__5034)
return
}


}, 1)

tmp6015 := Call(__e, ns2_1set, symshen_4_5s_1exprs_6, tmp5795)


_ = tmp6015

tmp6016 := MakeNative(func(__e *ControlFlow) {
V2773 := __e.Get(1)
_ = V2773
V2774 := __e.Get(2)
_ = V2774
tmp6034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2773)
}
__typedArg0 := V2773
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6021 Obj

if True == tmp6034 {
tmp6032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2773)
}
__typedArg0 := V2773
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_3, tmp6032)
}
__typedArg0 := sym_3
__typedArg1 := tmp6032
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres6023 Obj

if True == tmp6033 {
tmp6030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2773)
}
__typedArg0 := V2773
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp6030)
}
__typedArg0 := tmp6030
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6025 Obj

if True == tmp6031 {
tmp6027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2773)
}
__typedArg0 := V2773
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6027)
}
__typedArg0 := tmp6027
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp6028)
}
__typedArg0 := Nil
__typedArg1 := tmp6028
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
tmp6017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2773)
}
__typedArg0 := V2773
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp6017)
}
__typedArg0 := tmp6017
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6019 := Call(__e, PrimFunc(symexplode), tmp6018)


__e.TailApply(PrimFunc(symappend), tmp6019, V2774)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V2773, V2774)
}
__typedArg0 := V2773
__typedArg1 := V2774
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
}


}, 2)

tmp6035 := Call(__e, ns2_1set, symshen_4add_1sexpr, tmp6016)


_ = tmp6035

tmp6036 := MakeNative(func(__e *ControlFlow) {
V2775 := __e.Get(1)
_ = V2775
tmp6041 := Call(__e, PrimFunc(symshen_4hds_a_2), V2775, MakeInteger(91))


var ifres6037 Obj

if True == tmp6041 {
tmp6038 := Call(__e, PrimFunc(symtail), V2775)


let__5123 := tmp6038
_ = let__5123

tmp6039 := Call(__e, PrimFunc(symshen_4comb), let__5123, symshen_4skip)


ifres6037 = tmp6039


} else {
tmp6040 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6037 = tmp6040


}

let__5122 := ifres6037
_ = let__5122

tmp6043 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5122)


if True == tmp6043 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5122)
return
}


}, 1)

tmp6044 := Call(__e, ns2_1set, symshen_4_5lsb_6, tmp6036)


_ = tmp6044

tmp6045 := MakeNative(func(__e *ControlFlow) {
V2778 := __e.Get(1)
_ = V2778
tmp6050 := Call(__e, PrimFunc(symshen_4hds_a_2), V2778, MakeInteger(93))


var ifres6046 Obj

if True == tmp6050 {
tmp6047 := Call(__e, PrimFunc(symtail), V2778)


let__5125 := tmp6047
_ = let__5125

tmp6048 := Call(__e, PrimFunc(symshen_4comb), let__5125, symshen_4skip)


ifres6046 = tmp6048


} else {
tmp6049 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6046 = tmp6049


}

let__5124 := ifres6046
_ = let__5124

tmp6052 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5124)


if True == tmp6052 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5124)
return
}


}, 1)

tmp6053 := Call(__e, ns2_1set, symshen_4_5rsb_6, tmp6045)


_ = tmp6053

tmp6054 := MakeNative(func(__e *ControlFlow) {
V2781 := __e.Get(1)
_ = V2781
tmp6055 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), V2781)


let__5127 := tmp6055
_ = let__5127

tmp6061 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5127)


var ifres6056 Obj

if True == tmp6061 {
tmp6057 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6056 = tmp6057


} else {
tmp6058 := Call(__e, PrimFunc(symshen_4_5_1out), let__5127)


let__5128 := tmp6058
_ = let__5128

tmp6059 := Call(__e, PrimFunc(symshen_4in_1_6), let__5127)


let__5129 := tmp6059
_ = let__5129

tmp6060 := Call(__e, PrimFunc(symshen_4comb), let__5129, let__5128)


ifres6056 = tmp6060


}

let__5126 := ifres6056
_ = let__5126

tmp6063 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5126)


if True == tmp6063 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5126)
return
}


}, 1)

tmp6064 := Call(__e, ns2_1set, symshen_4_5s_1exprs1_6, tmp6054)


_ = tmp6064

tmp6065 := MakeNative(func(__e *ControlFlow) {
V2786 := __e.Get(1)
_ = V2786
tmp6066 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), V2786)


let__5131 := tmp6066
_ = let__5131

tmp6072 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5131)


var ifres6067 Obj

if True == tmp6072 {
tmp6068 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6067 = tmp6068


} else {
tmp6069 := Call(__e, PrimFunc(symshen_4_5_1out), let__5131)


let__5132 := tmp6069
_ = let__5132

tmp6070 := Call(__e, PrimFunc(symshen_4in_1_6), let__5131)


let__5133 := tmp6070
_ = let__5133

tmp6071 := Call(__e, PrimFunc(symshen_4comb), let__5133, let__5132)


ifres6067 = tmp6071


}

let__5130 := ifres6067
_ = let__5130

tmp6074 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5130)


if True == tmp6074 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5130)
return
}


}, 1)

tmp6075 := Call(__e, ns2_1set, symshen_4_5s_1exprs2_6, tmp6065)


_ = tmp6075

tmp6076 := MakeNative(func(__e *ControlFlow) {
V2792 := __e.Get(1)
_ = V2792
tmp6133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V2792)
}
__typedArg0 := Nil
__typedArg1 := V2792
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp6133 {
__e.Return(Nil)
return
} else {
tmp6131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2792)
}
__typedArg0 := V2792
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6111 Obj

if True == tmp6131 {
tmp6129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2792)
}
__typedArg0 := V2792
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp6129)
}
__typedArg0 := tmp6129
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6113 Obj

if True == tmp6130 {
tmp6126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2792)
}
__typedArg0 := V2792
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6126)
}
__typedArg0 := tmp6126
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp6127)
}
__typedArg0 := tmp6127
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6115 Obj

if True == tmp6128 {
tmp6122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2792)
}
__typedArg0 := V2792
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6122)
}
__typedArg0 := tmp6122
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6123)
}
__typedArg0 := tmp6123
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp6124)
}
__typedArg0 := Nil
__typedArg1 := tmp6124
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres6117 Obj

if True == tmp6125 {
tmp6119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2792)
}
__typedArg0 := V2792
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp6119)
}
__typedArg0 := tmp6119
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp6120, symbar_b)
}
__typedArg0 := tmp6120
__typedArg1 := symbar_b
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
tmp6077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2792)
}
__typedArg0 := V2792
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2792)
}
__typedArg0 := V2792
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6078)
}
__typedArg0 := tmp6078
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp6077, tmp6079)
}
__typedArg0 := tmp6077
__typedArg1 := tmp6079
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp6080)
}
__typedArg0 := symcons
__typedArg1 := tmp6080
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp6109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2792)
}
__typedArg0 := V2792
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6089 Obj

if True == tmp6109 {
tmp6107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2792)
}
__typedArg0 := V2792
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp6107)
}
__typedArg0 := tmp6107
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6091 Obj

if True == tmp6108 {
tmp6104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2792)
}
__typedArg0 := V2792
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6104)
}
__typedArg0 := tmp6104
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp6105)
}
__typedArg0 := tmp6105
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6093 Obj

if True == tmp6106 {
tmp6100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2792)
}
__typedArg0 := V2792
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6100)
}
__typedArg0 := tmp6100
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp6101)
}
__typedArg0 := tmp6101
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp6102)
}
__typedArg0 := tmp6102
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6095 Obj

if True == tmp6103 {
tmp6097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2792)
}
__typedArg0 := V2792
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp6097)
}
__typedArg0 := tmp6097
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp6098, symbar_b)
}
__typedArg0 := tmp6098
__typedArg1 := symbar_b
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("misapplication of |\n"))
}
__typedArg0 := MakeString("misapplication of |\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
tmp6087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2792)
}
__typedArg0 := V2792
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp6087 {
tmp6081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V2792)
}
__typedArg0 := V2792
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V2792)
}
__typedArg0 := V2792
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6083 := Call(__e, PrimFunc(symshen_4cons_1form), tmp6082)


tmp6084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp6083, Nil)
}
__typedArg0 := tmp6083
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp6081, tmp6084)
}
__typedArg0 := tmp6081
__typedArg1 := tmp6084
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp6085)
}
__typedArg0 := symcons
__typedArg1 := tmp6085
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

tmp6134 := Call(__e, ns2_1set, symshen_4cons_1form, tmp6076)


_ = tmp6134

tmp6135 := MakeNative(func(__e *ControlFlow) {
V2793 := __e.Get(1)
_ = V2793
tmp6140 := Call(__e, PrimFunc(symshen_4hds_a_2), V2793, MakeInteger(40))


var ifres6136 Obj

if True == tmp6140 {
tmp6137 := Call(__e, PrimFunc(symtail), V2793)


let__5135 := tmp6137
_ = let__5135

tmp6138 := Call(__e, PrimFunc(symshen_4comb), let__5135, symshen_4skip)


ifres6136 = tmp6138


} else {
tmp6139 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6136 = tmp6139


}

let__5134 := ifres6136
_ = let__5134

tmp6142 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5134)


if True == tmp6142 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5134)
return
}


}, 1)

tmp6143 := Call(__e, ns2_1set, symshen_4_5lrb_6, tmp6135)


_ = tmp6143

tmp6144 := MakeNative(func(__e *ControlFlow) {
V2796 := __e.Get(1)
_ = V2796
tmp6149 := Call(__e, PrimFunc(symshen_4hds_a_2), V2796, MakeInteger(41))


var ifres6145 Obj

if True == tmp6149 {
tmp6146 := Call(__e, PrimFunc(symtail), V2796)


let__5137 := tmp6146
_ = let__5137

tmp6147 := Call(__e, PrimFunc(symshen_4comb), let__5137, symshen_4skip)


ifres6145 = tmp6147


} else {
tmp6148 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6145 = tmp6148


}

let__5136 := ifres6145
_ = let__5136

tmp6151 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5136)


if True == tmp6151 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5136)
return
}


}, 1)

tmp6152 := Call(__e, ns2_1set, symshen_4_5rrb_6, tmp6144)


_ = tmp6152

tmp6153 := MakeNative(func(__e *ControlFlow) {
V2799 := __e.Get(1)
_ = V2799
tmp6158 := Call(__e, PrimFunc(symshen_4hds_a_2), V2799, MakeInteger(123))


var ifres6154 Obj

if True == tmp6158 {
tmp6155 := Call(__e, PrimFunc(symtail), V2799)


let__5139 := tmp6155
_ = let__5139

tmp6156 := Call(__e, PrimFunc(symshen_4comb), let__5139, symshen_4skip)


ifres6154 = tmp6156


} else {
tmp6157 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6154 = tmp6157


}

let__5138 := ifres6154
_ = let__5138

tmp6160 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5138)


if True == tmp6160 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5138)
return
}


}, 1)

tmp6161 := Call(__e, ns2_1set, symshen_4_5lcurly_6, tmp6153)


_ = tmp6161

tmp6162 := MakeNative(func(__e *ControlFlow) {
V2802 := __e.Get(1)
_ = V2802
tmp6167 := Call(__e, PrimFunc(symshen_4hds_a_2), V2802, MakeInteger(125))


var ifres6163 Obj

if True == tmp6167 {
tmp6164 := Call(__e, PrimFunc(symtail), V2802)


let__5141 := tmp6164
_ = let__5141

tmp6165 := Call(__e, PrimFunc(symshen_4comb), let__5141, symshen_4skip)


ifres6163 = tmp6165


} else {
tmp6166 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6163 = tmp6166


}

let__5140 := ifres6163
_ = let__5140

tmp6169 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5140)


if True == tmp6169 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5140)
return
}


}, 1)

tmp6170 := Call(__e, ns2_1set, symshen_4_5rcurly_6, tmp6162)


_ = tmp6170

tmp6171 := MakeNative(func(__e *ControlFlow) {
V2805 := __e.Get(1)
_ = V2805
tmp6176 := Call(__e, PrimFunc(symshen_4hds_a_2), V2805, MakeInteger(124))


var ifres6172 Obj

if True == tmp6176 {
tmp6173 := Call(__e, PrimFunc(symtail), V2805)


let__5143 := tmp6173
_ = let__5143

tmp6174 := Call(__e, PrimFunc(symshen_4comb), let__5143, symshen_4skip)


ifres6172 = tmp6174


} else {
tmp6175 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6172 = tmp6175


}

let__5142 := ifres6172
_ = let__5142

tmp6178 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5142)


if True == tmp6178 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5142)
return
}


}, 1)

tmp6179 := Call(__e, ns2_1set, symshen_4_5bar_6, tmp6171)


_ = tmp6179

tmp6180 := MakeNative(func(__e *ControlFlow) {
V2808 := __e.Get(1)
_ = V2808
tmp6185 := Call(__e, PrimFunc(symshen_4hds_a_2), V2808, MakeInteger(59))


var ifres6181 Obj

if True == tmp6185 {
tmp6182 := Call(__e, PrimFunc(symtail), V2808)


let__5145 := tmp6182
_ = let__5145

tmp6183 := Call(__e, PrimFunc(symshen_4comb), let__5145, symshen_4skip)


ifres6181 = tmp6183


} else {
tmp6184 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6181 = tmp6184


}

let__5144 := ifres6181
_ = let__5144

tmp6187 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5144)


if True == tmp6187 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5144)
return
}


}, 1)

tmp6188 := Call(__e, ns2_1set, symshen_4_5semicolon_6, tmp6180)


_ = tmp6188

tmp6189 := MakeNative(func(__e *ControlFlow) {
V2811 := __e.Get(1)
_ = V2811
tmp6194 := Call(__e, PrimFunc(symshen_4hds_a_2), V2811, MakeInteger(58))


var ifres6190 Obj

if True == tmp6194 {
tmp6191 := Call(__e, PrimFunc(symtail), V2811)


let__5147 := tmp6191
_ = let__5147

tmp6192 := Call(__e, PrimFunc(symshen_4comb), let__5147, symshen_4skip)


ifres6190 = tmp6192


} else {
tmp6193 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6190 = tmp6193


}

let__5146 := ifres6190
_ = let__5146

tmp6196 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5146)


if True == tmp6196 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5146)
return
}


}, 1)

tmp6197 := Call(__e, ns2_1set, symshen_4_5colon_6, tmp6189)


_ = tmp6197

tmp6198 := MakeNative(func(__e *ControlFlow) {
V2814 := __e.Get(1)
_ = V2814
tmp6203 := Call(__e, PrimFunc(symshen_4hds_a_2), V2814, MakeInteger(44))


var ifres6199 Obj

if True == tmp6203 {
tmp6200 := Call(__e, PrimFunc(symtail), V2814)


let__5149 := tmp6200
_ = let__5149

tmp6201 := Call(__e, PrimFunc(symshen_4comb), let__5149, symshen_4skip)


ifres6199 = tmp6201


} else {
tmp6202 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6199 = tmp6202


}

let__5148 := ifres6199
_ = let__5148

tmp6205 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5148)


if True == tmp6205 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5148)
return
}


}, 1)

tmp6206 := Call(__e, ns2_1set, symshen_4_5comma_6, tmp6198)


_ = tmp6206

tmp6207 := MakeNative(func(__e *ControlFlow) {
V2817 := __e.Get(1)
_ = V2817
tmp6212 := Call(__e, PrimFunc(symshen_4hds_a_2), V2817, MakeInteger(61))


var ifres6208 Obj

if True == tmp6212 {
tmp6209 := Call(__e, PrimFunc(symtail), V2817)


let__5151 := tmp6209
_ = let__5151

tmp6210 := Call(__e, PrimFunc(symshen_4comb), let__5151, symshen_4skip)


ifres6208 = tmp6210


} else {
tmp6211 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6208 = tmp6211


}

let__5150 := ifres6208
_ = let__5150

tmp6214 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5150)


if True == tmp6214 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5150)
return
}


}, 1)

tmp6215 := Call(__e, ns2_1set, symshen_4_5equal_6, tmp6207)


_ = tmp6215

tmp6216 := MakeNative(func(__e *ControlFlow) {
V2820 := __e.Get(1)
_ = V2820
tmp6217 := Call(__e, PrimFunc(symshen_4_5singleline_6), V2820)


let__5153 := tmp6217
_ = let__5153

tmp6222 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5153)


var ifres6218 Obj

if True == tmp6222 {
tmp6219 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6218 = tmp6219


} else {
tmp6220 := Call(__e, PrimFunc(symshen_4in_1_6), let__5153)


let__5154 := tmp6220
_ = let__5154

tmp6221 := Call(__e, PrimFunc(symshen_4comb), let__5154, symshen_4skip)


ifres6218 = tmp6221


}

let__5152 := ifres6218
_ = let__5152

tmp6232 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5152)


if True == tmp6232 {
tmp6223 := Call(__e, PrimFunc(symshen_4_5multiline_6), V2820)


let__5156 := tmp6223
_ = let__5156

tmp6228 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5156)


var ifres6224 Obj

if True == tmp6228 {
tmp6225 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6224 = tmp6225


} else {
tmp6226 := Call(__e, PrimFunc(symshen_4in_1_6), let__5156)


let__5157 := tmp6226
_ = let__5157

tmp6227 := Call(__e, PrimFunc(symshen_4comb), let__5157, symshen_4skip)


ifres6224 = tmp6227


}

let__5155 := ifres6224
_ = let__5155

tmp6230 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5155)


if True == tmp6230 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5155)
return
}


} else {
__e.Return(let__5152)
return
}


}, 1)

tmp6233 := Call(__e, ns2_1set, symshen_4_5comment_6, tmp6216)


_ = tmp6233

tmp6234 := MakeNative(func(__e *ControlFlow) {
V2827 := __e.Get(1)
_ = V2827
tmp6235 := Call(__e, PrimFunc(symshen_4_5backslash_6), V2827)


let__5159 := tmp6235
_ = let__5159

tmp6255 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5159)


var ifres6236 Obj

if True == tmp6255 {
tmp6237 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6236 = tmp6237


} else {
tmp6238 := Call(__e, PrimFunc(symshen_4in_1_6), let__5159)


let__5160 := tmp6238
_ = let__5160

tmp6239 := Call(__e, PrimFunc(symshen_4_5backslash_6), let__5160)


let__5161 := tmp6239
_ = let__5161

tmp6254 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5161)


var ifres6240 Obj

if True == tmp6254 {
tmp6241 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6240 = tmp6241


} else {
tmp6242 := Call(__e, PrimFunc(symshen_4in_1_6), let__5161)


let__5162 := tmp6242
_ = let__5162

tmp6243 := Call(__e, PrimFunc(symshen_4_5shortnatters_6), let__5162)


let__5163 := tmp6243
_ = let__5163

tmp6253 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5163)


var ifres6244 Obj

if True == tmp6253 {
tmp6245 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6244 = tmp6245


} else {
tmp6246 := Call(__e, PrimFunc(symshen_4in_1_6), let__5163)


let__5164 := tmp6246
_ = let__5164

tmp6247 := Call(__e, PrimFunc(symshen_4_5returns_6), let__5164)


let__5165 := tmp6247
_ = let__5165

tmp6252 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5165)


var ifres6248 Obj

if True == tmp6252 {
tmp6249 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6248 = tmp6249


} else {
tmp6250 := Call(__e, PrimFunc(symshen_4in_1_6), let__5165)


let__5166 := tmp6250
_ = let__5166

tmp6251 := Call(__e, PrimFunc(symshen_4comb), let__5166, symshen_4skip)


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

tmp6257 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5158)


if True == tmp6257 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5158)
return
}


}, 1)

tmp6258 := Call(__e, ns2_1set, symshen_4_5singleline_6, tmp6234)


_ = tmp6258

tmp6259 := MakeNative(func(__e *ControlFlow) {
V2837 := __e.Get(1)
_ = V2837
tmp6264 := Call(__e, PrimFunc(symshen_4hds_a_2), V2837, MakeInteger(92))


var ifres6260 Obj

if True == tmp6264 {
tmp6261 := Call(__e, PrimFunc(symtail), V2837)


let__5168 := tmp6261
_ = let__5168

tmp6262 := Call(__e, PrimFunc(symshen_4comb), let__5168, symshen_4skip)


ifres6260 = tmp6262


} else {
tmp6263 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6260 = tmp6263


}

let__5167 := ifres6260
_ = let__5167

tmp6266 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5167)


if True == tmp6266 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5167)
return
}


}, 1)

tmp6267 := Call(__e, ns2_1set, symshen_4_5backslash_6, tmp6259)


_ = tmp6267

tmp6268 := MakeNative(func(__e *ControlFlow) {
V2840 := __e.Get(1)
_ = V2840
tmp6269 := Call(__e, PrimFunc(symshen_4_5shortnatter_6), V2840)


let__5170 := tmp6269
_ = let__5170

tmp6279 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5170)


var ifres6270 Obj

if True == tmp6279 {
tmp6271 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6270 = tmp6271


} else {
tmp6272 := Call(__e, PrimFunc(symshen_4in_1_6), let__5170)


let__5171 := tmp6272
_ = let__5171

tmp6273 := Call(__e, PrimFunc(symshen_4_5shortnatters_6), let__5171)


let__5172 := tmp6273
_ = let__5172

tmp6278 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5172)


var ifres6274 Obj

if True == tmp6278 {
tmp6275 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6274 = tmp6275


} else {
tmp6276 := Call(__e, PrimFunc(symshen_4in_1_6), let__5172)


let__5173 := tmp6276
_ = let__5173

tmp6277 := Call(__e, PrimFunc(symshen_4comb), let__5173, symshen_4skip)


ifres6274 = tmp6277


}

ifres6270 = ifres6274


}

let__5169 := ifres6270
_ = let__5169

tmp6289 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5169)


if True == tmp6289 {
tmp6280 := Call(__e, PrimFunc(sym_5e_6), V2840)


let__5175 := tmp6280
_ = let__5175

tmp6285 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5175)


var ifres6281 Obj

if True == tmp6285 {
tmp6282 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6281 = tmp6282


} else {
tmp6283 := Call(__e, PrimFunc(symshen_4in_1_6), let__5175)


let__5176 := tmp6283
_ = let__5176

tmp6284 := Call(__e, PrimFunc(symshen_4comb), let__5176, symshen_4skip)


ifres6281 = tmp6284


}

let__5174 := ifres6281
_ = let__5174

tmp6287 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5174)


if True == tmp6287 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5174)
return
}


} else {
__e.Return(let__5169)
return
}


}, 1)

tmp6290 := Call(__e, ns2_1set, symshen_4_5shortnatters_6, tmp6268)


_ = tmp6290

tmp6291 := MakeNative(func(__e *ControlFlow) {
V2849 := __e.Get(1)
_ = V2849
tmp6301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2849)
}
__typedArg0 := V2849
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6292 Obj

if True == tmp6301 {
tmp6293 := Call(__e, PrimFunc(symhead), V2849)


let__5178 := tmp6293
_ = let__5178

tmp6294 := Call(__e, PrimFunc(symtail), V2849)


let__5179 := tmp6294
_ = let__5179

tmp6298 := Call(__e, PrimFunc(symshen_4return_2), let__5178)


tmp6299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp6298)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp6298
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres6295 Obj

if True == tmp6299 {
tmp6296 := Call(__e, PrimFunc(symshen_4comb), let__5179, symshen_4skip)


ifres6295 = tmp6296


} else {
tmp6297 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6295 = tmp6297


}

ifres6292 = ifres6295


} else {
tmp6300 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6292 = tmp6300


}

let__5177 := ifres6292
_ = let__5177

tmp6303 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5177)


if True == tmp6303 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5177)
return
}


}, 1)

tmp6304 := Call(__e, ns2_1set, symshen_4_5shortnatter_6, tmp6291)


_ = tmp6304

tmp6305 := MakeNative(func(__e *ControlFlow) {
V2853 := __e.Get(1)
_ = V2853
tmp6306 := Call(__e, PrimFunc(symshen_4_5return_6), V2853)


let__5181 := tmp6306
_ = let__5181

tmp6316 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5181)


var ifres6307 Obj

if True == tmp6316 {
tmp6308 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6307 = tmp6308


} else {
tmp6309 := Call(__e, PrimFunc(symshen_4in_1_6), let__5181)


let__5182 := tmp6309
_ = let__5182

tmp6310 := Call(__e, PrimFunc(symshen_4_5returns_6), let__5182)


let__5183 := tmp6310
_ = let__5183

tmp6315 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5183)


var ifres6311 Obj

if True == tmp6315 {
tmp6312 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6311 = tmp6312


} else {
tmp6313 := Call(__e, PrimFunc(symshen_4in_1_6), let__5183)


let__5184 := tmp6313
_ = let__5184

tmp6314 := Call(__e, PrimFunc(symshen_4comb), let__5184, symshen_4skip)


ifres6311 = tmp6314


}

ifres6307 = ifres6311


}

let__5180 := ifres6307
_ = let__5180

tmp6326 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5180)


if True == tmp6326 {
tmp6317 := Call(__e, PrimFunc(symshen_4_5return_6), V2853)


let__5186 := tmp6317
_ = let__5186

tmp6322 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5186)


var ifres6318 Obj

if True == tmp6322 {
tmp6319 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6318 = tmp6319


} else {
tmp6320 := Call(__e, PrimFunc(symshen_4in_1_6), let__5186)


let__5187 := tmp6320
_ = let__5187

tmp6321 := Call(__e, PrimFunc(symshen_4comb), let__5187, symshen_4skip)


ifres6318 = tmp6321


}

let__5185 := ifres6318
_ = let__5185

tmp6324 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5185)


if True == tmp6324 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5185)
return
}


} else {
__e.Return(let__5180)
return
}


}, 1)

tmp6327 := Call(__e, ns2_1set, symshen_4_5returns_6, tmp6305)


_ = tmp6327

tmp6328 := MakeNative(func(__e *ControlFlow) {
V2862 := __e.Get(1)
_ = V2862
tmp6337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2862)
}
__typedArg0 := V2862
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6329 Obj

if True == tmp6337 {
tmp6330 := Call(__e, PrimFunc(symhead), V2862)


let__5189 := tmp6330
_ = let__5189

tmp6331 := Call(__e, PrimFunc(symtail), V2862)


let__5190 := tmp6331
_ = let__5190

tmp6335 := Call(__e, PrimFunc(symshen_4return_2), let__5189)


var ifres6332 Obj

if True == tmp6335 {
tmp6333 := Call(__e, PrimFunc(symshen_4comb), let__5190, symshen_4skip)


ifres6332 = tmp6333


} else {
tmp6334 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6332 = tmp6334


}

ifres6329 = ifres6332


} else {
tmp6336 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6329 = tmp6336


}

let__5188 := ifres6329
_ = let__5188

tmp6339 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5188)


if True == tmp6339 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5188)
return
}


}, 1)

tmp6340 := Call(__e, ns2_1set, symshen_4_5return_6, tmp6328)


_ = tmp6340

tmp6341 := MakeNative(func(__e *ControlFlow) {
V2866 := __e.Get(1)
_ = V2866
tmp6342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(13), Nil)
}
__typedArg0 := MakeInteger(13)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(10), tmp6342)
}
__typedArg0 := MakeInteger(10)
__typedArg1 := tmp6342
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(9), tmp6343)
}
__typedArg0 := MakeInteger(9)
__typedArg1 := tmp6343
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symelement_2), V2866, tmp6344)
return


}, 1)

tmp6345 := Call(__e, ns2_1set, symshen_4return_2, tmp6341)


_ = tmp6345

tmp6346 := MakeNative(func(__e *ControlFlow) {
V2867 := __e.Get(1)
_ = V2867
tmp6347 := Call(__e, PrimFunc(symshen_4_5backslash_6), V2867)


let__5192 := tmp6347
_ = let__5192

tmp6362 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5192)


var ifres6348 Obj

if True == tmp6362 {
tmp6349 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6348 = tmp6349


} else {
tmp6350 := Call(__e, PrimFunc(symshen_4in_1_6), let__5192)


let__5193 := tmp6350
_ = let__5193

tmp6351 := Call(__e, PrimFunc(symshen_4_5times_6), let__5193)


let__5194 := tmp6351
_ = let__5194

tmp6361 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5194)


var ifres6352 Obj

if True == tmp6361 {
tmp6353 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6352 = tmp6353


} else {
tmp6354 := Call(__e, PrimFunc(symshen_4in_1_6), let__5194)


let__5195 := tmp6354
_ = let__5195

tmp6355 := Call(__e, PrimFunc(symshen_4_5longnatter_6), let__5195)


let__5196 := tmp6355
_ = let__5196

tmp6360 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5196)


var ifres6356 Obj

if True == tmp6360 {
tmp6357 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6356 = tmp6357


} else {
tmp6358 := Call(__e, PrimFunc(symshen_4in_1_6), let__5196)


let__5197 := tmp6358
_ = let__5197

tmp6359 := Call(__e, PrimFunc(symshen_4comb), let__5197, symshen_4skip)


ifres6356 = tmp6359


}

ifres6352 = ifres6356


}

ifres6348 = ifres6352


}

let__5191 := ifres6348
_ = let__5191

tmp6364 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5191)


if True == tmp6364 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5191)
return
}


}, 1)

tmp6365 := Call(__e, ns2_1set, symshen_4_5multiline_6, tmp6346)


_ = tmp6365

tmp6366 := MakeNative(func(__e *ControlFlow) {
V2875 := __e.Get(1)
_ = V2875
tmp6371 := Call(__e, PrimFunc(symshen_4hds_a_2), V2875, MakeInteger(42))


var ifres6367 Obj

if True == tmp6371 {
tmp6368 := Call(__e, PrimFunc(symtail), V2875)


let__5199 := tmp6368
_ = let__5199

tmp6369 := Call(__e, PrimFunc(symshen_4comb), let__5199, symshen_4skip)


ifres6367 = tmp6369


} else {
tmp6370 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6367 = tmp6370


}

let__5198 := ifres6367
_ = let__5198

tmp6373 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5198)


if True == tmp6373 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5198)
return
}


}, 1)

tmp6374 := Call(__e, ns2_1set, symshen_4_5times_6, tmp6366)


_ = tmp6374

tmp6375 := MakeNative(func(__e *ControlFlow) {
V2878 := __e.Get(1)
_ = V2878
tmp6376 := Call(__e, PrimFunc(symshen_4_5comment_6), V2878)


let__5201 := tmp6376
_ = let__5201

tmp6386 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5201)


var ifres6377 Obj

if True == tmp6386 {
tmp6378 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6377 = tmp6378


} else {
tmp6379 := Call(__e, PrimFunc(symshen_4in_1_6), let__5201)


let__5202 := tmp6379
_ = let__5202

tmp6380 := Call(__e, PrimFunc(symshen_4_5longnatter_6), let__5202)


let__5203 := tmp6380
_ = let__5203

tmp6385 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5203)


var ifres6381 Obj

if True == tmp6385 {
tmp6382 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6381 = tmp6382


} else {
tmp6383 := Call(__e, PrimFunc(symshen_4in_1_6), let__5203)


let__5204 := tmp6383
_ = let__5204

tmp6384 := Call(__e, PrimFunc(symshen_4comb), let__5204, symshen_4skip)


ifres6381 = tmp6384


}

ifres6377 = ifres6381


}

let__5200 := ifres6377
_ = let__5200

tmp6413 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5200)


if True == tmp6413 {
tmp6387 := Call(__e, PrimFunc(symshen_4_5times_6), V2878)


let__5206 := tmp6387
_ = let__5206

tmp6397 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5206)


var ifres6388 Obj

if True == tmp6397 {
tmp6389 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6388 = tmp6389


} else {
tmp6390 := Call(__e, PrimFunc(symshen_4in_1_6), let__5206)


let__5207 := tmp6390
_ = let__5207

tmp6391 := Call(__e, PrimFunc(symshen_4_5backslash_6), let__5207)


let__5208 := tmp6391
_ = let__5208

tmp6396 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5208)


var ifres6392 Obj

if True == tmp6396 {
tmp6393 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6392 = tmp6393


} else {
tmp6394 := Call(__e, PrimFunc(symshen_4in_1_6), let__5208)


let__5209 := tmp6394
_ = let__5209

tmp6395 := Call(__e, PrimFunc(symshen_4comb), let__5209, symshen_4skip)


ifres6392 = tmp6395


}

ifres6388 = ifres6392


}

let__5205 := ifres6388
_ = let__5205

tmp6411 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5205)


if True == tmp6411 {
tmp6407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2878)
}
__typedArg0 := V2878
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6398 Obj

if True == tmp6407 {
tmp6399 := Call(__e, PrimFunc(symtail), V2878)


let__5211 := tmp6399
_ = let__5211

tmp6400 := Call(__e, PrimFunc(symshen_4_5longnatter_6), let__5211)


let__5212 := tmp6400
_ = let__5212

tmp6405 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5212)


var ifres6401 Obj

if True == tmp6405 {
tmp6402 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6401 = tmp6402


} else {
tmp6403 := Call(__e, PrimFunc(symshen_4in_1_6), let__5212)


let__5213 := tmp6403
_ = let__5213

tmp6404 := Call(__e, PrimFunc(symshen_4comb), let__5213, symshen_4skip)


ifres6401 = tmp6404


}

ifres6398 = ifres6401


} else {
tmp6406 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6398 = tmp6406


}

let__5210 := ifres6398
_ = let__5210

tmp6409 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5210)


if True == tmp6409 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5210)
return
}


} else {
__e.Return(let__5205)
return
}


} else {
__e.Return(let__5200)
return
}


}, 1)

tmp6414 := Call(__e, ns2_1set, symshen_4_5longnatter_6, tmp6375)


_ = tmp6414

tmp6415 := MakeNative(func(__e *ControlFlow) {
V2893 := __e.Get(1)
_ = V2893
tmp6416 := Call(__e, PrimFunc(symshen_4_5str_6), V2893)


let__5215 := tmp6416
_ = let__5215

tmp6422 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5215)


var ifres6417 Obj

if True == tmp6422 {
tmp6418 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6417 = tmp6418


} else {
tmp6419 := Call(__e, PrimFunc(symshen_4_5_1out), let__5215)


let__5216 := tmp6419
_ = let__5216

tmp6420 := Call(__e, PrimFunc(symshen_4in_1_6), let__5215)


let__5217 := tmp6420
_ = let__5217

tmp6421 := Call(__e, PrimFunc(symshen_4comb), let__5217, let__5216)


ifres6417 = tmp6421


}

let__5214 := ifres6417
_ = let__5214

tmp6447 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5214)


if True == tmp6447 {
tmp6423 := Call(__e, PrimFunc(symshen_4_5number_6), V2893)


let__5219 := tmp6423
_ = let__5219

tmp6429 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5219)


var ifres6424 Obj

if True == tmp6429 {
tmp6425 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6424 = tmp6425


} else {
tmp6426 := Call(__e, PrimFunc(symshen_4_5_1out), let__5219)


let__5220 := tmp6426
_ = let__5220

tmp6427 := Call(__e, PrimFunc(symshen_4in_1_6), let__5219)


let__5221 := tmp6427
_ = let__5221

tmp6428 := Call(__e, PrimFunc(symshen_4comb), let__5221, let__5220)


ifres6424 = tmp6428


}

let__5218 := ifres6424
_ = let__5218

tmp6445 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5218)


if True == tmp6445 {
tmp6430 := Call(__e, PrimFunc(symshen_4_5sym_6), V2893)


let__5223 := tmp6430
_ = let__5223

tmp6441 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5223)


var ifres6431 Obj

if True == tmp6441 {
tmp6432 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6431 = tmp6432


} else {
tmp6433 := Call(__e, PrimFunc(symshen_4_5_1out), let__5223)


let__5224 := tmp6433
_ = let__5224

tmp6434 := Call(__e, PrimFunc(symshen_4in_1_6), let__5223)


let__5225 := tmp6434
_ = let__5225

tmp6439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__5224, MakeString("<>"))
}
__typedArg0 := let__5224
__typedArg1 := MakeString("<>")
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres6435 Obj

if True == tmp6439 {
tmp6436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), Nil)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp6436)
}
__typedArg0 := symvector
__typedArg1 := tmp6436
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

ifres6435 = tmp6437


} else {
tmp6438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(let__5224)
}
__typedArg0 := let__5224
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

ifres6435 = tmp6438


}

tmp6440 := Call(__e, PrimFunc(symshen_4comb), let__5225, ifres6435)


ifres6431 = tmp6440


}

let__5222 := ifres6431
_ = let__5222

tmp6443 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5222)


if True == tmp6443 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5222)
return
}


} else {
__e.Return(let__5218)
return
}


} else {
__e.Return(let__5214)
return
}


}, 1)

tmp6448 := Call(__e, ns2_1set, symshen_4_5atom_6, tmp6415)


_ = tmp6448

tmp6449 := MakeNative(func(__e *ControlFlow) {
V2906 := __e.Get(1)
_ = V2906
tmp6450 := Call(__e, PrimFunc(symshen_4_5alpha_6), V2906)


let__5227 := tmp6450
_ = let__5227

tmp6463 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5227)


var ifres6451 Obj

if True == tmp6463 {
tmp6452 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6451 = tmp6452


} else {
tmp6453 := Call(__e, PrimFunc(symshen_4_5_1out), let__5227)


let__5228 := tmp6453
_ = let__5228

tmp6454 := Call(__e, PrimFunc(symshen_4in_1_6), let__5227)


let__5229 := tmp6454
_ = let__5229

tmp6455 := Call(__e, PrimFunc(symshen_4_5alphanums_6), let__5229)


let__5230 := tmp6455
_ = let__5230

tmp6462 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5230)


var ifres6456 Obj

if True == tmp6462 {
tmp6457 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6456 = tmp6457


} else {
tmp6458 := Call(__e, PrimFunc(symshen_4_5_1out), let__5230)


let__5231 := tmp6458
_ = let__5231

tmp6459 := Call(__e, PrimFunc(symshen_4in_1_6), let__5230)


let__5232 := tmp6459
_ = let__5232

tmp6461 := Call(__e, PrimFunc(symshen_4comb), let__5232, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(let__5228)
__typedS1, __typedOK1 := TypedString(let__5231)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := let__5228
__typedArg1 := let__5231
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())


ifres6456 = tmp6461


}

ifres6451 = ifres6456


}

let__5226 := ifres6451
_ = let__5226

tmp6465 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5226)


if True == tmp6465 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5226)
return
}


}, 1)

tmp6466 := Call(__e, ns2_1set, symshen_4_5sym_6, tmp6449)


_ = tmp6466

tmp6467 := MakeNative(func(__e *ControlFlow) {
V2914 := __e.Get(1)
_ = V2914
tmp6477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2914)
}
__typedArg0 := V2914
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6468 Obj

if True == tmp6477 {
tmp6469 := Call(__e, PrimFunc(symhead), V2914)


let__5234 := tmp6469
_ = let__5234

tmp6470 := Call(__e, PrimFunc(symtail), V2914)


let__5235 := tmp6470
_ = let__5235

tmp6475 := Call(__e, PrimFunc(symshen_4alpha_2), let__5234)


var ifres6471 Obj

if True == tmp6475 {
tmp6472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(let__5234)
}
__typedArg0 := let__5234
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp6473 := Call(__e, PrimFunc(symshen_4comb), let__5235, tmp6472)


ifres6471 = tmp6473


} else {
tmp6474 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6471 = tmp6474


}

ifres6468 = ifres6471


} else {
tmp6476 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6468 = tmp6476


}

let__5233 := ifres6468
_ = let__5233

tmp6479 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5233)


if True == tmp6479 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5233)
return
}


}, 1)

tmp6480 := Call(__e, ns2_1set, symshen_4_5alpha_6, tmp6467)


_ = tmp6480

tmp6481 := MakeNative(func(__e *ControlFlow) {
V2918 := __e.Get(1)
_ = V2918
tmp6488 := Call(__e, PrimFunc(symshen_4lowercase_2), V2918)


if True == tmp6488 {
__e.Return(True)
return
} else {
tmp6486 := Call(__e, PrimFunc(symshen_4uppercase_2), V2918)


var ifres6483 Obj

if True == tmp6486 {
ifres6483 = True


} else {
tmp6485 := Call(__e, PrimFunc(symshen_4misc_2), V2918)


var ifres6484 Obj

if True == tmp6485 {
ifres6484 = True


} else {
ifres6484 = False


}

ifres6483 = ifres6484


}

if True == ifres6483 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp6489 := Call(__e, ns2_1set, symshen_4alpha_2, tmp6481)


_ = tmp6489

tmp6490 := MakeNative(func(__e *ControlFlow) {
V2919 := __e.Get(1)
_ = V2919
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6_a) {
__typedN0, __typedOK0 := TypedFloat64(V2919)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(97))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6_a) {
return TypedMaterializeBoolean((__typedN0 >= __typedN1))
}}
__typedArg0 := V2919
__typedArg1 := MakeInteger(97)
return Call(__e, PrimFunc(sym_6_a), __typedArg0, __typedArg1)
})() {
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_a) {
__typedN0, __typedOK0 := TypedFloat64(V2919)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(122))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5_a) {
return TypedMaterializeBoolean((__typedN0 <= __typedN1))
}}
__typedArg0 := V2919
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

tmp6495 := Call(__e, ns2_1set, symshen_4lowercase_2, tmp6490)


_ = tmp6495

tmp6496 := MakeNative(func(__e *ControlFlow) {
V2920 := __e.Get(1)
_ = V2920
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6_a) {
__typedN0, __typedOK0 := TypedFloat64(V2920)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(65))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6_a) {
return TypedMaterializeBoolean((__typedN0 >= __typedN1))
}}
__typedArg0 := V2920
__typedArg1 := MakeInteger(65)
return Call(__e, PrimFunc(sym_6_a), __typedArg0, __typedArg1)
})() {
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_a) {
__typedN0, __typedOK0 := TypedFloat64(V2920)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(90))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5_a) {
return TypedMaterializeBoolean((__typedN0 <= __typedN1))
}}
__typedArg0 := V2920
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

tmp6501 := Call(__e, ns2_1set, symshen_4uppercase_2, tmp6496)


_ = tmp6501

tmp6502 := MakeNative(func(__e *ControlFlow) {
V2921 := __e.Get(1)
_ = V2921
tmp6503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(96), Nil)
}
__typedArg0 := MakeInteger(96)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(35), tmp6503)
}
__typedArg0 := MakeInteger(35)
__typedArg1 := tmp6503
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(39), tmp6504)
}
__typedArg0 := MakeInteger(39)
__typedArg1 := tmp6504
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(37), tmp6505)
}
__typedArg0 := MakeInteger(37)
__typedArg1 := tmp6505
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(38), tmp6506)
}
__typedArg0 := MakeInteger(38)
__typedArg1 := tmp6506
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(60), tmp6507)
}
__typedArg0 := MakeInteger(60)
__typedArg1 := tmp6507
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(62), tmp6508)
}
__typedArg0 := MakeInteger(62)
__typedArg1 := tmp6508
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(46), tmp6509)
}
__typedArg0 := MakeInteger(46)
__typedArg1 := tmp6509
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(126), tmp6510)
}
__typedArg0 := MakeInteger(126)
__typedArg1 := tmp6510
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(64), tmp6511)
}
__typedArg0 := MakeInteger(64)
__typedArg1 := tmp6511
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(33), tmp6512)
}
__typedArg0 := MakeInteger(33)
__typedArg1 := tmp6512
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(36), tmp6513)
}
__typedArg0 := MakeInteger(36)
__typedArg1 := tmp6513
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(63), tmp6514)
}
__typedArg0 := MakeInteger(63)
__typedArg1 := tmp6514
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(95), tmp6515)
}
__typedArg0 := MakeInteger(95)
__typedArg1 := tmp6515
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(43), tmp6516)
}
__typedArg0 := MakeInteger(43)
__typedArg1 := tmp6516
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(47), tmp6517)
}
__typedArg0 := MakeInteger(47)
__typedArg1 := tmp6517
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(42), tmp6518)
}
__typedArg0 := MakeInteger(42)
__typedArg1 := tmp6518
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(45), tmp6519)
}
__typedArg0 := MakeInteger(45)
__typedArg1 := tmp6519
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(61), tmp6520)
}
__typedArg0 := MakeInteger(61)
__typedArg1 := tmp6520
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symelement_2), V2921, tmp6521)
return


}, 1)

tmp6522 := Call(__e, ns2_1set, symshen_4misc_2, tmp6502)


_ = tmp6522

tmp6523 := MakeNative(func(__e *ControlFlow) {
V2922 := __e.Get(1)
_ = V2922
tmp6524 := Call(__e, PrimFunc(symshen_4_5alphanum_6), V2922)


let__5237 := tmp6524
_ = let__5237

tmp6537 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5237)


var ifres6525 Obj

if True == tmp6537 {
tmp6526 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6525 = tmp6526


} else {
tmp6527 := Call(__e, PrimFunc(symshen_4_5_1out), let__5237)


let__5238 := tmp6527
_ = let__5238

tmp6528 := Call(__e, PrimFunc(symshen_4in_1_6), let__5237)


let__5239 := tmp6528
_ = let__5239

tmp6529 := Call(__e, PrimFunc(symshen_4_5alphanums_6), let__5239)


let__5240 := tmp6529
_ = let__5240

tmp6536 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5240)


var ifres6530 Obj

if True == tmp6536 {
tmp6531 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6530 = tmp6531


} else {
tmp6532 := Call(__e, PrimFunc(symshen_4_5_1out), let__5240)


let__5241 := tmp6532
_ = let__5241

tmp6533 := Call(__e, PrimFunc(symshen_4in_1_6), let__5240)


let__5242 := tmp6533
_ = let__5242

tmp6535 := Call(__e, PrimFunc(symshen_4comb), let__5242, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(let__5238)
__typedS1, __typedOK1 := TypedString(let__5241)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := let__5238
__typedArg1 := let__5241
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())


ifres6530 = tmp6535


}

ifres6525 = ifres6530


}

let__5236 := ifres6525
_ = let__5236

tmp6547 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5236)


if True == tmp6547 {
tmp6538 := Call(__e, PrimFunc(sym_5e_6), V2922)


let__5244 := tmp6538
_ = let__5244

tmp6543 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5244)


var ifres6539 Obj

if True == tmp6543 {
tmp6540 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6539 = tmp6540


} else {
tmp6541 := Call(__e, PrimFunc(symshen_4in_1_6), let__5244)


let__5245 := tmp6541
_ = let__5245

tmp6542 := Call(__e, PrimFunc(symshen_4comb), let__5245, MakeString(""))


ifres6539 = tmp6542


}

let__5243 := ifres6539
_ = let__5243

tmp6545 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5243)


if True == tmp6545 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5243)
return
}


} else {
__e.Return(let__5236)
return
}


}, 1)

tmp6548 := Call(__e, ns2_1set, symshen_4_5alphanums_6, tmp6523)


_ = tmp6548

tmp6549 := MakeNative(func(__e *ControlFlow) {
V2933 := __e.Get(1)
_ = V2933
tmp6550 := Call(__e, PrimFunc(symshen_4_5alpha_6), V2933)


let__5247 := tmp6550
_ = let__5247

tmp6556 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5247)


var ifres6551 Obj

if True == tmp6556 {
tmp6552 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6551 = tmp6552


} else {
tmp6553 := Call(__e, PrimFunc(symshen_4_5_1out), let__5247)


let__5248 := tmp6553
_ = let__5248

tmp6554 := Call(__e, PrimFunc(symshen_4in_1_6), let__5247)


let__5249 := tmp6554
_ = let__5249

tmp6555 := Call(__e, PrimFunc(symshen_4comb), let__5249, let__5248)


ifres6551 = tmp6555


}

let__5246 := ifres6551
_ = let__5246

tmp6567 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5246)


if True == tmp6567 {
tmp6557 := Call(__e, PrimFunc(symshen_4_5numeral_6), V2933)


let__5251 := tmp6557
_ = let__5251

tmp6563 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5251)


var ifres6558 Obj

if True == tmp6563 {
tmp6559 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6558 = tmp6559


} else {
tmp6560 := Call(__e, PrimFunc(symshen_4_5_1out), let__5251)


let__5252 := tmp6560
_ = let__5252

tmp6561 := Call(__e, PrimFunc(symshen_4in_1_6), let__5251)


let__5253 := tmp6561
_ = let__5253

tmp6562 := Call(__e, PrimFunc(symshen_4comb), let__5253, let__5252)


ifres6558 = tmp6562


}

let__5250 := ifres6558
_ = let__5250

tmp6565 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5250)


if True == tmp6565 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5250)
return
}


} else {
__e.Return(let__5246)
return
}


}, 1)

tmp6568 := Call(__e, ns2_1set, symshen_4_5alphanum_6, tmp6549)


_ = tmp6568

tmp6569 := MakeNative(func(__e *ControlFlow) {
V2942 := __e.Get(1)
_ = V2942
tmp6579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2942)
}
__typedArg0 := V2942
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6570 Obj

if True == tmp6579 {
tmp6571 := Call(__e, PrimFunc(symhead), V2942)


let__5255 := tmp6571
_ = let__5255

tmp6572 := Call(__e, PrimFunc(symtail), V2942)


let__5256 := tmp6572
_ = let__5256

tmp6577 := Call(__e, PrimFunc(symshen_4digit_2), let__5255)


var ifres6573 Obj

if True == tmp6577 {
tmp6574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(let__5255)
}
__typedArg0 := let__5255
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp6575 := Call(__e, PrimFunc(symshen_4comb), let__5256, tmp6574)


ifres6573 = tmp6575


} else {
tmp6576 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6573 = tmp6576


}

ifres6570 = ifres6573


} else {
tmp6578 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6570 = tmp6578


}

let__5254 := ifres6570
_ = let__5254

tmp6581 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5254)


if True == tmp6581 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5254)
return
}


}, 1)

tmp6582 := Call(__e, ns2_1set, symshen_4_5numeral_6, tmp6569)


_ = tmp6582

tmp6583 := MakeNative(func(__e *ControlFlow) {
V2946 := __e.Get(1)
_ = V2946
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6_a) {
__typedN0, __typedOK0 := TypedFloat64(V2946)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(48))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6_a) {
return TypedMaterializeBoolean((__typedN0 >= __typedN1))
}}
__typedArg0 := V2946
__typedArg1 := MakeInteger(48)
return Call(__e, PrimFunc(sym_6_a), __typedArg0, __typedArg1)
})() {
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_a) {
__typedN0, __typedOK0 := TypedFloat64(V2946)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(57))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5_a) {
return TypedMaterializeBoolean((__typedN0 <= __typedN1))
}}
__typedArg0 := V2946
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

tmp6588 := Call(__e, ns2_1set, symshen_4digit_2, tmp6583)


_ = tmp6588

tmp6589 := MakeNative(func(__e *ControlFlow) {
V2947 := __e.Get(1)
_ = V2947
tmp6590 := Call(__e, PrimFunc(symshen_4_5dbq_6), V2947)


let__5258 := tmp6590
_ = let__5258

tmp6606 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5258)


var ifres6591 Obj

if True == tmp6606 {
tmp6592 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6591 = tmp6592


} else {
tmp6593 := Call(__e, PrimFunc(symshen_4in_1_6), let__5258)


let__5259 := tmp6593
_ = let__5259

tmp6594 := Call(__e, PrimFunc(symshen_4_5strcontents_6), let__5259)


let__5260 := tmp6594
_ = let__5260

tmp6605 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5260)


var ifres6595 Obj

if True == tmp6605 {
tmp6596 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6595 = tmp6596


} else {
tmp6597 := Call(__e, PrimFunc(symshen_4_5_1out), let__5260)


let__5261 := tmp6597
_ = let__5261

tmp6598 := Call(__e, PrimFunc(symshen_4in_1_6), let__5260)


let__5262 := tmp6598
_ = let__5262

tmp6599 := Call(__e, PrimFunc(symshen_4_5dbq_6), let__5262)


let__5263 := tmp6599
_ = let__5263

tmp6604 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5263)


var ifres6600 Obj

if True == tmp6604 {
tmp6601 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6600 = tmp6601


} else {
tmp6602 := Call(__e, PrimFunc(symshen_4in_1_6), let__5263)


let__5264 := tmp6602
_ = let__5264

tmp6603 := Call(__e, PrimFunc(symshen_4comb), let__5264, let__5261)


ifres6600 = tmp6603


}

ifres6595 = ifres6600


}

ifres6591 = ifres6595


}

let__5257 := ifres6591
_ = let__5257

tmp6608 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5257)


if True == tmp6608 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5257)
return
}


}, 1)

tmp6609 := Call(__e, ns2_1set, symshen_4_5str_6, tmp6589)


_ = tmp6609

tmp6610 := MakeNative(func(__e *ControlFlow) {
V2956 := __e.Get(1)
_ = V2956
tmp6615 := Call(__e, PrimFunc(symshen_4hds_a_2), V2956, MakeInteger(34))


var ifres6611 Obj

if True == tmp6615 {
tmp6612 := Call(__e, PrimFunc(symtail), V2956)


let__5266 := tmp6612
_ = let__5266

tmp6613 := Call(__e, PrimFunc(symshen_4comb), let__5266, symshen_4skip)


ifres6611 = tmp6613


} else {
tmp6614 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6611 = tmp6614


}

let__5265 := ifres6611
_ = let__5265

tmp6617 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5265)


if True == tmp6617 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5265)
return
}


}, 1)

tmp6618 := Call(__e, ns2_1set, symshen_4_5dbq_6, tmp6610)


_ = tmp6618

tmp6619 := MakeNative(func(__e *ControlFlow) {
V2959 := __e.Get(1)
_ = V2959
tmp6620 := Call(__e, PrimFunc(symshen_4_5strc_6), V2959)


let__5268 := tmp6620
_ = let__5268

tmp6633 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5268)


var ifres6621 Obj

if True == tmp6633 {
tmp6622 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6621 = tmp6622


} else {
tmp6623 := Call(__e, PrimFunc(symshen_4_5_1out), let__5268)


let__5269 := tmp6623
_ = let__5269

tmp6624 := Call(__e, PrimFunc(symshen_4in_1_6), let__5268)


let__5270 := tmp6624
_ = let__5270

tmp6625 := Call(__e, PrimFunc(symshen_4_5strcontents_6), let__5270)


let__5271 := tmp6625
_ = let__5271

tmp6632 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5271)


var ifres6626 Obj

if True == tmp6632 {
tmp6627 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6626 = tmp6627


} else {
tmp6628 := Call(__e, PrimFunc(symshen_4_5_1out), let__5271)


let__5272 := tmp6628
_ = let__5272

tmp6629 := Call(__e, PrimFunc(symshen_4in_1_6), let__5271)


let__5273 := tmp6629
_ = let__5273

tmp6631 := Call(__e, PrimFunc(symshen_4comb), let__5273, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(let__5269)
__typedS1, __typedOK1 := TypedString(let__5272)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := let__5269
__typedArg1 := let__5272
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())


ifres6626 = tmp6631


}

ifres6621 = ifres6626


}

let__5267 := ifres6621
_ = let__5267

tmp6643 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5267)


if True == tmp6643 {
tmp6634 := Call(__e, PrimFunc(sym_5e_6), V2959)


let__5275 := tmp6634
_ = let__5275

tmp6639 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5275)


var ifres6635 Obj

if True == tmp6639 {
tmp6636 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6635 = tmp6636


} else {
tmp6637 := Call(__e, PrimFunc(symshen_4in_1_6), let__5275)


let__5276 := tmp6637
_ = let__5276

tmp6638 := Call(__e, PrimFunc(symshen_4comb), let__5276, MakeString(""))


ifres6635 = tmp6638


}

let__5274 := ifres6635
_ = let__5274

tmp6641 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5274)


if True == tmp6641 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5274)
return
}


} else {
__e.Return(let__5267)
return
}


}, 1)

tmp6644 := Call(__e, ns2_1set, symshen_4_5strcontents_6, tmp6619)


_ = tmp6644

tmp6645 := MakeNative(func(__e *ControlFlow) {
V2970 := __e.Get(1)
_ = V2970
tmp6646 := Call(__e, PrimFunc(symshen_4_5control_6), V2970)


let__5278 := tmp6646
_ = let__5278

tmp6652 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5278)


var ifres6647 Obj

if True == tmp6652 {
tmp6648 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6647 = tmp6648


} else {
tmp6649 := Call(__e, PrimFunc(symshen_4_5_1out), let__5278)


let__5279 := tmp6649
_ = let__5279

tmp6650 := Call(__e, PrimFunc(symshen_4in_1_6), let__5278)


let__5280 := tmp6650
_ = let__5280

tmp6651 := Call(__e, PrimFunc(symshen_4comb), let__5280, let__5279)


ifres6647 = tmp6651


}

let__5277 := ifres6647
_ = let__5277

tmp6663 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5277)


if True == tmp6663 {
tmp6653 := Call(__e, PrimFunc(symshen_4_5notdbq_6), V2970)


let__5282 := tmp6653
_ = let__5282

tmp6659 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5282)


var ifres6654 Obj

if True == tmp6659 {
tmp6655 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6654 = tmp6655


} else {
tmp6656 := Call(__e, PrimFunc(symshen_4_5_1out), let__5282)


let__5283 := tmp6656
_ = let__5283

tmp6657 := Call(__e, PrimFunc(symshen_4in_1_6), let__5282)


let__5284 := tmp6657
_ = let__5284

tmp6658 := Call(__e, PrimFunc(symshen_4comb), let__5284, let__5283)


ifres6654 = tmp6658


}

let__5281 := ifres6654
_ = let__5281

tmp6661 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5281)


if True == tmp6661 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5281)
return
}


} else {
__e.Return(let__5277)
return
}


}, 1)

tmp6664 := Call(__e, ns2_1set, symshen_4_5strc_6, tmp6645)


_ = tmp6664

tmp6665 := MakeNative(func(__e *ControlFlow) {
V2979 := __e.Get(1)
_ = V2979
tmp6666 := Call(__e, PrimFunc(symshen_4_5lowC_6), V2979)


let__5286 := tmp6666
_ = let__5286

tmp6688 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5286)


var ifres6667 Obj

if True == tmp6688 {
tmp6668 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6667 = tmp6668


} else {
tmp6669 := Call(__e, PrimFunc(symshen_4in_1_6), let__5286)


let__5287 := tmp6669
_ = let__5287

tmp6670 := Call(__e, PrimFunc(symshen_4_5hash_6), let__5287)


let__5288 := tmp6670
_ = let__5288

tmp6687 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5288)


var ifres6671 Obj

if True == tmp6687 {
tmp6672 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6671 = tmp6672


} else {
tmp6673 := Call(__e, PrimFunc(symshen_4in_1_6), let__5288)


let__5289 := tmp6673
_ = let__5289

tmp6674 := Call(__e, PrimFunc(symshen_4_5integer_6), let__5289)


let__5290 := tmp6674
_ = let__5290

tmp6686 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5290)


var ifres6675 Obj

if True == tmp6686 {
tmp6676 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6675 = tmp6676


} else {
tmp6677 := Call(__e, PrimFunc(symshen_4_5_1out), let__5290)


let__5291 := tmp6677
_ = let__5291

tmp6678 := Call(__e, PrimFunc(symshen_4in_1_6), let__5290)


let__5292 := tmp6678
_ = let__5292

tmp6679 := Call(__e, PrimFunc(symshen_4_5semicolon_6), let__5292)


let__5293 := tmp6679
_ = let__5293

tmp6685 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5293)


var ifres6680 Obj

if True == tmp6685 {
tmp6681 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6680 = tmp6681


} else {
tmp6682 := Call(__e, PrimFunc(symshen_4in_1_6), let__5293)


let__5294 := tmp6682
_ = let__5294

tmp6683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(let__5291)
}
__typedArg0 := let__5291
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp6684 := Call(__e, PrimFunc(symshen_4comb), let__5294, tmp6683)


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

tmp6690 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5285)


if True == tmp6690 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5285)
return
}


}, 1)

tmp6691 := Call(__e, ns2_1set, symshen_4_5control_6, tmp6665)


_ = tmp6691

tmp6692 := MakeNative(func(__e *ControlFlow) {
V2990 := __e.Get(1)
_ = V2990
tmp6703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V2990)
}
__typedArg0 := V2990
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6693 Obj

if True == tmp6703 {
tmp6694 := Call(__e, PrimFunc(symhead), V2990)


let__5296 := tmp6694
_ = let__5296

tmp6695 := Call(__e, PrimFunc(symtail), V2990)


let__5297 := tmp6695
_ = let__5297

tmp6700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__5296, MakeInteger(34))
}
__typedArg0 := let__5296
__typedArg1 := MakeInteger(34)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp6701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp6700)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp6700
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres6696 Obj

if True == tmp6701 {
tmp6697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symn_1_6string) {
return PrimNumberToString(let__5296)
}
__typedArg0 := let__5296
return Call(__e, PrimFunc(symn_1_6string), __typedArg0)
})()

tmp6698 := Call(__e, PrimFunc(symshen_4comb), let__5297, tmp6697)


ifres6696 = tmp6698


} else {
tmp6699 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6696 = tmp6699


}

ifres6693 = ifres6696


} else {
tmp6702 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6693 = tmp6702


}

let__5295 := ifres6693
_ = let__5295

tmp6705 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5295)


if True == tmp6705 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5295)
return
}


}, 1)

tmp6706 := Call(__e, ns2_1set, symshen_4_5notdbq_6, tmp6692)


_ = tmp6706

tmp6707 := MakeNative(func(__e *ControlFlow) {
V2994 := __e.Get(1)
_ = V2994
tmp6712 := Call(__e, PrimFunc(symshen_4hds_a_2), V2994, MakeInteger(99))


var ifres6708 Obj

if True == tmp6712 {
tmp6709 := Call(__e, PrimFunc(symtail), V2994)


let__5299 := tmp6709
_ = let__5299

tmp6710 := Call(__e, PrimFunc(symshen_4comb), let__5299, symshen_4skip)


ifres6708 = tmp6710


} else {
tmp6711 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6708 = tmp6711


}

let__5298 := ifres6708
_ = let__5298

tmp6714 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5298)


if True == tmp6714 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5298)
return
}


}, 1)

tmp6715 := Call(__e, ns2_1set, symshen_4_5lowC_6, tmp6707)


_ = tmp6715

tmp6716 := MakeNative(func(__e *ControlFlow) {
V2997 := __e.Get(1)
_ = V2997
tmp6721 := Call(__e, PrimFunc(symshen_4hds_a_2), V2997, MakeInteger(35))


var ifres6717 Obj

if True == tmp6721 {
tmp6718 := Call(__e, PrimFunc(symtail), V2997)


let__5301 := tmp6718
_ = let__5301

tmp6719 := Call(__e, PrimFunc(symshen_4comb), let__5301, symshen_4skip)


ifres6717 = tmp6719


} else {
tmp6720 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6717 = tmp6720


}

let__5300 := ifres6717
_ = let__5300

tmp6723 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5300)


if True == tmp6723 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5300)
return
}


}, 1)

tmp6724 := Call(__e, ns2_1set, symshen_4_5hash_6, tmp6716)


_ = tmp6724

tmp6725 := MakeNative(func(__e *ControlFlow) {
V3000 := __e.Get(1)
_ = V3000
tmp6726 := Call(__e, PrimFunc(symshen_4_5minus_6), V3000)


let__5303 := tmp6726
_ = let__5303

tmp6738 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5303)


var ifres6727 Obj

if True == tmp6738 {
tmp6728 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6727 = tmp6728


} else {
tmp6729 := Call(__e, PrimFunc(symshen_4in_1_6), let__5303)


let__5304 := tmp6729
_ = let__5304

tmp6730 := Call(__e, PrimFunc(symshen_4_5number_6), let__5304)


let__5305 := tmp6730
_ = let__5305

tmp6737 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5305)


var ifres6731 Obj

if True == tmp6737 {
tmp6732 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6731 = tmp6732


} else {
tmp6733 := Call(__e, PrimFunc(symshen_4_5_1out), let__5305)


let__5306 := tmp6733
_ = let__5306

tmp6734 := Call(__e, PrimFunc(symshen_4in_1_6), let__5305)


let__5307 := tmp6734
_ = let__5307

tmp6736 := Call(__e, PrimFunc(symshen_4comb), let__5307, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(0))
__typedN1, __typedOK1 := TypedFloat64(let__5306)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := MakeInteger(0)
__typedArg1 := let__5306
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


ifres6731 = tmp6736


}

ifres6727 = ifres6731


}

let__5302 := ifres6727
_ = let__5302

tmp6781 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5302)


if True == tmp6781 {
tmp6739 := Call(__e, PrimFunc(symshen_4_5plus_6), V3000)


let__5309 := tmp6739
_ = let__5309

tmp6750 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5309)


var ifres6740 Obj

if True == tmp6750 {
tmp6741 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6740 = tmp6741


} else {
tmp6742 := Call(__e, PrimFunc(symshen_4in_1_6), let__5309)


let__5310 := tmp6742
_ = let__5310

tmp6743 := Call(__e, PrimFunc(symshen_4_5number_6), let__5310)


let__5311 := tmp6743
_ = let__5311

tmp6749 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5311)


var ifres6744 Obj

if True == tmp6749 {
tmp6745 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6744 = tmp6745


} else {
tmp6746 := Call(__e, PrimFunc(symshen_4_5_1out), let__5311)


let__5312 := tmp6746
_ = let__5312

tmp6747 := Call(__e, PrimFunc(symshen_4in_1_6), let__5311)


let__5313 := tmp6747
_ = let__5313

tmp6748 := Call(__e, PrimFunc(symshen_4comb), let__5313, let__5312)


ifres6744 = tmp6748


}

ifres6740 = ifres6744


}

let__5308 := ifres6740
_ = let__5308

tmp6779 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5308)


if True == tmp6779 {
tmp6751 := Call(__e, PrimFunc(symshen_4_5e_1number_6), V3000)


let__5315 := tmp6751
_ = let__5315

tmp6757 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5315)


var ifres6752 Obj

if True == tmp6757 {
tmp6753 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6752 = tmp6753


} else {
tmp6754 := Call(__e, PrimFunc(symshen_4_5_1out), let__5315)


let__5316 := tmp6754
_ = let__5316

tmp6755 := Call(__e, PrimFunc(symshen_4in_1_6), let__5315)


let__5317 := tmp6755
_ = let__5317

tmp6756 := Call(__e, PrimFunc(symshen_4comb), let__5317, let__5316)


ifres6752 = tmp6756


}

let__5314 := ifres6752
_ = let__5314

tmp6777 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5314)


if True == tmp6777 {
tmp6758 := Call(__e, PrimFunc(symshen_4_5float_6), V3000)


let__5319 := tmp6758
_ = let__5319

tmp6764 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5319)


var ifres6759 Obj

if True == tmp6764 {
tmp6760 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6759 = tmp6760


} else {
tmp6761 := Call(__e, PrimFunc(symshen_4_5_1out), let__5319)


let__5320 := tmp6761
_ = let__5320

tmp6762 := Call(__e, PrimFunc(symshen_4in_1_6), let__5319)


let__5321 := tmp6762
_ = let__5321

tmp6763 := Call(__e, PrimFunc(symshen_4comb), let__5321, let__5320)


ifres6759 = tmp6763


}

let__5318 := ifres6759
_ = let__5318

tmp6775 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5318)


if True == tmp6775 {
tmp6765 := Call(__e, PrimFunc(symshen_4_5integer_6), V3000)


let__5323 := tmp6765
_ = let__5323

tmp6771 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5323)


var ifres6766 Obj

if True == tmp6771 {
tmp6767 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6766 = tmp6767


} else {
tmp6768 := Call(__e, PrimFunc(symshen_4_5_1out), let__5323)


let__5324 := tmp6768
_ = let__5324

tmp6769 := Call(__e, PrimFunc(symshen_4in_1_6), let__5323)


let__5325 := tmp6769
_ = let__5325

tmp6770 := Call(__e, PrimFunc(symshen_4comb), let__5325, let__5324)


ifres6766 = tmp6770


}

let__5322 := ifres6766
_ = let__5322

tmp6773 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5322)


if True == tmp6773 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5322)
return
}


} else {
__e.Return(let__5318)
return
}


} else {
__e.Return(let__5314)
return
}


} else {
__e.Return(let__5308)
return
}


} else {
__e.Return(let__5302)
return
}


}, 1)

tmp6782 := Call(__e, ns2_1set, symshen_4_5number_6, tmp6725)


_ = tmp6782

tmp6783 := MakeNative(func(__e *ControlFlow) {
V3025 := __e.Get(1)
_ = V3025
tmp6788 := Call(__e, PrimFunc(symshen_4hds_a_2), V3025, MakeInteger(45))


var ifres6784 Obj

if True == tmp6788 {
tmp6785 := Call(__e, PrimFunc(symtail), V3025)


let__5327 := tmp6785
_ = let__5327

tmp6786 := Call(__e, PrimFunc(symshen_4comb), let__5327, symshen_4skip)


ifres6784 = tmp6786


} else {
tmp6787 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6784 = tmp6787


}

let__5326 := ifres6784
_ = let__5326

tmp6790 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5326)


if True == tmp6790 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5326)
return
}


}, 1)

tmp6791 := Call(__e, ns2_1set, symshen_4_5minus_6, tmp6783)


_ = tmp6791

tmp6792 := MakeNative(func(__e *ControlFlow) {
V3028 := __e.Get(1)
_ = V3028
tmp6797 := Call(__e, PrimFunc(symshen_4hds_a_2), V3028, MakeInteger(43))


var ifres6793 Obj

if True == tmp6797 {
tmp6794 := Call(__e, PrimFunc(symtail), V3028)


let__5329 := tmp6794
_ = let__5329

tmp6795 := Call(__e, PrimFunc(symshen_4comb), let__5329, symshen_4skip)


ifres6793 = tmp6795


} else {
tmp6796 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6793 = tmp6796


}

let__5328 := ifres6793
_ = let__5328

tmp6799 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5328)


if True == tmp6799 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5328)
return
}


}, 1)

tmp6800 := Call(__e, ns2_1set, symshen_4_5plus_6, tmp6792)


_ = tmp6800

tmp6801 := MakeNative(func(__e *ControlFlow) {
V3031 := __e.Get(1)
_ = V3031
tmp6802 := Call(__e, PrimFunc(symshen_4_5digits_6), V3031)


let__5331 := tmp6802
_ = let__5331

tmp6809 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5331)


var ifres6803 Obj

if True == tmp6809 {
tmp6804 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6803 = tmp6804


} else {
tmp6805 := Call(__e, PrimFunc(symshen_4_5_1out), let__5331)


let__5332 := tmp6805
_ = let__5332

tmp6806 := Call(__e, PrimFunc(symshen_4in_1_6), let__5331)


let__5333 := tmp6806
_ = let__5333

tmp6807 := Call(__e, PrimFunc(symshen_4compute_1integer), let__5332)


tmp6808 := Call(__e, PrimFunc(symshen_4comb), let__5333, tmp6807)


ifres6803 = tmp6808


}

let__5330 := ifres6803
_ = let__5330

tmp6811 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5330)


if True == tmp6811 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5330)
return
}


}, 1)

tmp6812 := Call(__e, ns2_1set, symshen_4_5integer_6, tmp6801)


_ = tmp6812

tmp6813 := MakeNative(func(__e *ControlFlow) {
V3036 := __e.Get(1)
_ = V3036
tmp6814 := Call(__e, PrimFunc(symshen_4_5digit_6), V3036)


let__5335 := tmp6814
_ = let__5335

tmp6827 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5335)


var ifres6815 Obj

if True == tmp6827 {
tmp6816 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6815 = tmp6816


} else {
tmp6817 := Call(__e, PrimFunc(symshen_4_5_1out), let__5335)


let__5336 := tmp6817
_ = let__5336

tmp6818 := Call(__e, PrimFunc(symshen_4in_1_6), let__5335)


let__5337 := tmp6818
_ = let__5337

tmp6819 := Call(__e, PrimFunc(symshen_4_5digits_6), let__5337)


let__5338 := tmp6819
_ = let__5338

tmp6826 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5338)


var ifres6820 Obj

if True == tmp6826 {
tmp6821 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6820 = tmp6821


} else {
tmp6822 := Call(__e, PrimFunc(symshen_4_5_1out), let__5338)


let__5339 := tmp6822
_ = let__5339

tmp6823 := Call(__e, PrimFunc(symshen_4in_1_6), let__5338)


let__5340 := tmp6823
_ = let__5340

tmp6824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__5336, let__5339)
}
__typedArg0 := let__5336
__typedArg1 := let__5339
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6825 := Call(__e, PrimFunc(symshen_4comb), let__5340, tmp6824)


ifres6820 = tmp6825


}

ifres6815 = ifres6820


}

let__5334 := ifres6815
_ = let__5334

tmp6839 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5334)


if True == tmp6839 {
tmp6828 := Call(__e, PrimFunc(symshen_4_5digit_6), V3036)


let__5342 := tmp6828
_ = let__5342

tmp6835 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5342)


var ifres6829 Obj

if True == tmp6835 {
tmp6830 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6829 = tmp6830


} else {
tmp6831 := Call(__e, PrimFunc(symshen_4_5_1out), let__5342)


let__5343 := tmp6831
_ = let__5343

tmp6832 := Call(__e, PrimFunc(symshen_4in_1_6), let__5342)


let__5344 := tmp6832
_ = let__5344

tmp6833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__5343, Nil)
}
__typedArg0 := let__5343
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp6834 := Call(__e, PrimFunc(symshen_4comb), let__5344, tmp6833)


ifres6829 = tmp6834


}

let__5341 := ifres6829
_ = let__5341

tmp6837 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5341)


if True == tmp6837 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5341)
return
}


} else {
__e.Return(let__5334)
return
}


}, 1)

tmp6840 := Call(__e, ns2_1set, symshen_4_5digits_6, tmp6813)


_ = tmp6840

tmp6841 := MakeNative(func(__e *ControlFlow) {
V3048 := __e.Get(1)
_ = V3048
tmp6851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3048)
}
__typedArg0 := V3048
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres6842 Obj

if True == tmp6851 {
tmp6843 := Call(__e, PrimFunc(symhead), V3048)


let__5346 := tmp6843
_ = let__5346

tmp6844 := Call(__e, PrimFunc(symtail), V3048)


let__5347 := tmp6844
_ = let__5347

tmp6849 := Call(__e, PrimFunc(symshen_4digit_2), let__5346)


var ifres6845 Obj

if True == tmp6849 {
tmp6846 := Call(__e, PrimFunc(symshen_4byte_1_6digit), let__5346)


tmp6847 := Call(__e, PrimFunc(symshen_4comb), let__5347, tmp6846)


ifres6845 = tmp6847


} else {
tmp6848 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6845 = tmp6848


}

ifres6842 = ifres6845


} else {
tmp6850 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6842 = tmp6850


}

let__5345 := ifres6842
_ = let__5345

tmp6853 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5345)


if True == tmp6853 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5345)
return
}


}, 1)

tmp6854 := Call(__e, ns2_1set, symshen_4_5digit_6, tmp6841)


_ = tmp6854

tmp6855 := MakeNative(func(__e *ControlFlow) {
V3052 := __e.Get(1)
_ = V3052
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V3052)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(48))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V3052
__typedArg1 := MakeInteger(48)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp6856 := Call(__e, ns2_1set, symshen_4byte_1_6digit, tmp6855)


_ = tmp6856

tmp6857 := MakeNative(func(__e *ControlFlow) {
V3053 := __e.Get(1)
_ = V3053
tmp6858 := Call(__e, PrimFunc(symreverse), V3053)


__e.TailApply(PrimFunc(symshen_4compute_1integer_1h), tmp6858, MakeInteger(0))
return


}, 1)

tmp6859 := Call(__e, ns2_1set, symshen_4compute_1integer, tmp6857)


_ = tmp6859

tmp6860 := MakeNative(func(__e *ControlFlow) {
V3056 := __e.Get(1)
_ = V3056
V3057 := __e.Get(2)
_ = V3057
tmp6870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3056)
}
__typedArg0 := Nil
__typedArg1 := V3056
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp6870 {
__e.Return(MakeInteger(0))
return
} else {
tmp6868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3056)
}
__typedArg0 := V3056
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp6868 {
tmp6861 := Call(__e, PrimFunc(symshen_4expt), MakeInteger(10), V3057)


tmp6862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3056)
}
__typedArg0 := V3056
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(tmp6861)
__typedN1, __typedOK1 := TypedFloat64(tmp6862)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := tmp6861
__typedArg1 := tmp6862
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})()

tmp6864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3056)
}
__typedArg0 := V3056
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6866 := Call(__e, PrimFunc(symshen_4compute_1integer_1h), tmp6864, (func() Obj {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(tmp6863)
__typedN1, __typedOK1 := TypedFloat64(tmp6866)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := tmp6863
__typedArg1 := tmp6866
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4compute_1integer_1h)
return
}


}


}, 2)

tmp6871 := Call(__e, ns2_1set, symshen_4compute_1integer_1h, tmp6860)


_ = tmp6871

tmp6872 := MakeNative(func(__e *ControlFlow) {
V3060 := __e.Get(1)
_ = V3060
V3061 := __e.Get(2)
_ = V3061
tmp6880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V3061)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V3061
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp6880 {
__e.Return(MakeInteger(1))
return
} else {
if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(V3061)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(0))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := V3061
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
tmp6874 := Call(__e, PrimFunc(symshen_4expt), V3060, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V3061)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V3061
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(V3060)
__typedN1, __typedOK1 := TypedFloat64(tmp6874)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := V3060
__typedArg1 := tmp6874
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})())
return


} else {
tmp6876 := Call(__e, PrimFunc(symshen_4expt), V3060, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V3061)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V3061
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_c) {
__typedN0, __typedOK0 := TypedFloat64(tmp6876)
__typedN1, __typedOK1 := TypedFloat64(V3060)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_c) {
return TypedMaterializeNumber(TypedDivideValue(__typedN0, __typedN1))
}}
__typedArg0 := tmp6876
__typedArg1 := V3060
return Call(__e, PrimFunc(sym_c), __typedArg0, __typedArg1)
})())
return


}


}


}, 2)

tmp6881 := Call(__e, ns2_1set, symshen_4expt, tmp6872)


_ = tmp6881

tmp6882 := MakeNative(func(__e *ControlFlow) {
V3062 := __e.Get(1)
_ = V3062
tmp6883 := Call(__e, PrimFunc(symshen_4_5integer_6), V3062)


let__5349 := tmp6883
_ = let__5349

tmp6901 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5349)


var ifres6884 Obj

if True == tmp6901 {
tmp6885 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6884 = tmp6885


} else {
tmp6886 := Call(__e, PrimFunc(symshen_4_5_1out), let__5349)


let__5350 := tmp6886
_ = let__5350

tmp6887 := Call(__e, PrimFunc(symshen_4in_1_6), let__5349)


let__5351 := tmp6887
_ = let__5351

tmp6888 := Call(__e, PrimFunc(symshen_4_5stop_6), let__5351)


let__5352 := tmp6888
_ = let__5352

tmp6900 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5352)


var ifres6889 Obj

if True == tmp6900 {
tmp6890 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6889 = tmp6890


} else {
tmp6891 := Call(__e, PrimFunc(symshen_4in_1_6), let__5352)


let__5353 := tmp6891
_ = let__5353

tmp6892 := Call(__e, PrimFunc(symshen_4_5fraction_6), let__5353)


let__5354 := tmp6892
_ = let__5354

tmp6899 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5354)


var ifres6893 Obj

if True == tmp6899 {
tmp6894 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6893 = tmp6894


} else {
tmp6895 := Call(__e, PrimFunc(symshen_4_5_1out), let__5354)


let__5355 := tmp6895
_ = let__5355

tmp6896 := Call(__e, PrimFunc(symshen_4in_1_6), let__5354)


let__5356 := tmp6896
_ = let__5356

tmp6898 := Call(__e, PrimFunc(symshen_4comb), let__5356, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(let__5350)
__typedN1, __typedOK1 := TypedFloat64(let__5355)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := let__5350
__typedArg1 := let__5355
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())


ifres6893 = tmp6898


}

ifres6889 = ifres6893


}

ifres6884 = ifres6889


}

let__5348 := ifres6884
_ = let__5348

tmp6917 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5348)


if True == tmp6917 {
tmp6902 := Call(__e, PrimFunc(symshen_4_5stop_6), V3062)


let__5358 := tmp6902
_ = let__5358

tmp6913 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5358)


var ifres6903 Obj

if True == tmp6913 {
tmp6904 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6903 = tmp6904


} else {
tmp6905 := Call(__e, PrimFunc(symshen_4in_1_6), let__5358)


let__5359 := tmp6905
_ = let__5359

tmp6906 := Call(__e, PrimFunc(symshen_4_5fraction_6), let__5359)


let__5360 := tmp6906
_ = let__5360

tmp6912 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5360)


var ifres6907 Obj

if True == tmp6912 {
tmp6908 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6907 = tmp6908


} else {
tmp6909 := Call(__e, PrimFunc(symshen_4_5_1out), let__5360)


let__5361 := tmp6909
_ = let__5361

tmp6910 := Call(__e, PrimFunc(symshen_4in_1_6), let__5360)


let__5362 := tmp6910
_ = let__5362

tmp6911 := Call(__e, PrimFunc(symshen_4comb), let__5362, let__5361)


ifres6907 = tmp6911


}

ifres6903 = ifres6907


}

let__5357 := ifres6903
_ = let__5357

tmp6915 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5357)


if True == tmp6915 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5357)
return
}


} else {
__e.Return(let__5348)
return
}


}, 1)

tmp6918 := Call(__e, ns2_1set, symshen_4_5float_6, tmp6882)


_ = tmp6918

tmp6919 := MakeNative(func(__e *ControlFlow) {
V3078 := __e.Get(1)
_ = V3078
tmp6924 := Call(__e, PrimFunc(symshen_4hds_a_2), V3078, MakeInteger(46))


var ifres6920 Obj

if True == tmp6924 {
tmp6921 := Call(__e, PrimFunc(symtail), V3078)


let__5364 := tmp6921
_ = let__5364

tmp6922 := Call(__e, PrimFunc(symshen_4comb), let__5364, symshen_4skip)


ifres6920 = tmp6922


} else {
tmp6923 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6920 = tmp6923


}

let__5363 := ifres6920
_ = let__5363

tmp6926 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5363)


if True == tmp6926 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5363)
return
}


}, 1)

tmp6927 := Call(__e, ns2_1set, symshen_4_5stop_6, tmp6919)


_ = tmp6927

tmp6928 := MakeNative(func(__e *ControlFlow) {
V3081 := __e.Get(1)
_ = V3081
tmp6929 := Call(__e, PrimFunc(symshen_4_5digits_6), V3081)


let__5366 := tmp6929
_ = let__5366

tmp6936 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5366)


var ifres6930 Obj

if True == tmp6936 {
tmp6931 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6930 = tmp6931


} else {
tmp6932 := Call(__e, PrimFunc(symshen_4_5_1out), let__5366)


let__5367 := tmp6932
_ = let__5367

tmp6933 := Call(__e, PrimFunc(symshen_4in_1_6), let__5366)


let__5368 := tmp6933
_ = let__5368

tmp6934 := Call(__e, PrimFunc(symshen_4compute_1fraction), let__5367)


tmp6935 := Call(__e, PrimFunc(symshen_4comb), let__5368, tmp6934)


ifres6930 = tmp6935


}

let__5365 := ifres6930
_ = let__5365

tmp6938 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5365)


if True == tmp6938 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5365)
return
}


}, 1)

tmp6939 := Call(__e, ns2_1set, symshen_4_5fraction_6, tmp6928)


_ = tmp6939

tmp6940 := MakeNative(func(__e *ControlFlow) {
V3086 := __e.Get(1)
_ = V3086
__e.TailApply(PrimFunc(symshen_4compute_1fraction_1h), V3086, MakeInteger(-1))
return
}, 1)

tmp6941 := Call(__e, ns2_1set, symshen_4compute_1fraction, tmp6940)


_ = tmp6941

tmp6942 := MakeNative(func(__e *ControlFlow) {
V3089 := __e.Get(1)
_ = V3089
V3090 := __e.Get(2)
_ = V3090
tmp6952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3089)
}
__typedArg0 := Nil
__typedArg1 := V3089
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp6952 {
__e.Return(MakeInteger(0))
return
} else {
tmp6950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3089)
}
__typedArg0 := V3089
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp6950 {
tmp6943 := Call(__e, PrimFunc(symshen_4expt), MakeInteger(10), V3090)


tmp6944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3089)
}
__typedArg0 := V3089
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp6945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(tmp6943)
__typedN1, __typedOK1 := TypedFloat64(tmp6944)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := tmp6943
__typedArg1 := tmp6944
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})()

tmp6946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3089)
}
__typedArg0 := V3089
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp6948 := Call(__e, PrimFunc(symshen_4compute_1fraction_1h), tmp6946, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V3090)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V3090
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(tmp6945)
__typedN1, __typedOK1 := TypedFloat64(tmp6948)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := tmp6945
__typedArg1 := tmp6948
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4compute_1fraction_1h)
return
}


}


}, 2)

tmp6953 := Call(__e, ns2_1set, symshen_4compute_1fraction_1h, tmp6942)


_ = tmp6953

tmp6954 := MakeNative(func(__e *ControlFlow) {
V3091 := __e.Get(1)
_ = V3091
tmp6955 := Call(__e, PrimFunc(symshen_4_5float_6), V3091)


let__5370 := tmp6955
_ = let__5370

tmp6973 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5370)


var ifres6956 Obj

if True == tmp6973 {
tmp6957 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6956 = tmp6957


} else {
tmp6958 := Call(__e, PrimFunc(symshen_4_5_1out), let__5370)


let__5371 := tmp6958
_ = let__5371

tmp6959 := Call(__e, PrimFunc(symshen_4in_1_6), let__5370)


let__5372 := tmp6959
_ = let__5372

tmp6960 := Call(__e, PrimFunc(symshen_4_5lowE_6), let__5372)


let__5373 := tmp6960
_ = let__5373

tmp6972 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5373)


var ifres6961 Obj

if True == tmp6972 {
tmp6962 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6961 = tmp6962


} else {
tmp6963 := Call(__e, PrimFunc(symshen_4in_1_6), let__5373)


let__5374 := tmp6963
_ = let__5374

tmp6964 := Call(__e, PrimFunc(symshen_4_5log10_6), let__5374)


let__5375 := tmp6964
_ = let__5375

tmp6971 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5375)


var ifres6965 Obj

if True == tmp6971 {
tmp6966 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6965 = tmp6966


} else {
tmp6967 := Call(__e, PrimFunc(symshen_4_5_1out), let__5375)


let__5376 := tmp6967
_ = let__5376

tmp6968 := Call(__e, PrimFunc(symshen_4in_1_6), let__5375)


let__5377 := tmp6968
_ = let__5377

tmp6969 := Call(__e, PrimFunc(symshen_4compute_1E), let__5371, let__5376)


tmp6970 := Call(__e, PrimFunc(symshen_4comb), let__5377, tmp6969)


ifres6965 = tmp6970


}

ifres6961 = ifres6965


}

ifres6956 = ifres6961


}

let__5369 := ifres6956
_ = let__5369

tmp6996 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5369)


if True == tmp6996 {
tmp6974 := Call(__e, PrimFunc(symshen_4_5integer_6), V3091)


let__5379 := tmp6974
_ = let__5379

tmp6992 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5379)


var ifres6975 Obj

if True == tmp6992 {
tmp6976 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6975 = tmp6976


} else {
tmp6977 := Call(__e, PrimFunc(symshen_4_5_1out), let__5379)


let__5380 := tmp6977
_ = let__5380

tmp6978 := Call(__e, PrimFunc(symshen_4in_1_6), let__5379)


let__5381 := tmp6978
_ = let__5381

tmp6979 := Call(__e, PrimFunc(symshen_4_5lowE_6), let__5381)


let__5382 := tmp6979
_ = let__5382

tmp6991 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5382)


var ifres6980 Obj

if True == tmp6991 {
tmp6981 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6980 = tmp6981


} else {
tmp6982 := Call(__e, PrimFunc(symshen_4in_1_6), let__5382)


let__5383 := tmp6982
_ = let__5383

tmp6983 := Call(__e, PrimFunc(symshen_4_5log10_6), let__5383)


let__5384 := tmp6983
_ = let__5384

tmp6990 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5384)


var ifres6984 Obj

if True == tmp6990 {
tmp6985 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6984 = tmp6985


} else {
tmp6986 := Call(__e, PrimFunc(symshen_4_5_1out), let__5384)


let__5385 := tmp6986
_ = let__5385

tmp6987 := Call(__e, PrimFunc(symshen_4in_1_6), let__5384)


let__5386 := tmp6987
_ = let__5386

tmp6988 := Call(__e, PrimFunc(symshen_4compute_1E), let__5380, let__5385)


tmp6989 := Call(__e, PrimFunc(symshen_4comb), let__5386, tmp6988)


ifres6984 = tmp6989


}

ifres6980 = ifres6984


}

ifres6975 = ifres6980


}

let__5378 := ifres6975
_ = let__5378

tmp6994 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5378)


if True == tmp6994 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5378)
return
}


} else {
__e.Return(let__5369)
return
}


}, 1)

tmp6997 := Call(__e, ns2_1set, symshen_4_5e_1number_6, tmp6954)


_ = tmp6997

tmp6998 := MakeNative(func(__e *ControlFlow) {
V3110 := __e.Get(1)
_ = V3110
tmp6999 := Call(__e, PrimFunc(symshen_4_5plus_6), V3110)


let__5388 := tmp6999
_ = let__5388

tmp7010 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5388)


var ifres7000 Obj

if True == tmp7010 {
tmp7001 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres7000 = tmp7001


} else {
tmp7002 := Call(__e, PrimFunc(symshen_4in_1_6), let__5388)


let__5389 := tmp7002
_ = let__5389

tmp7003 := Call(__e, PrimFunc(symshen_4_5log10_6), let__5389)


let__5390 := tmp7003
_ = let__5390

tmp7009 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5390)


var ifres7004 Obj

if True == tmp7009 {
tmp7005 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres7004 = tmp7005


} else {
tmp7006 := Call(__e, PrimFunc(symshen_4_5_1out), let__5390)


let__5391 := tmp7006
_ = let__5391

tmp7007 := Call(__e, PrimFunc(symshen_4in_1_6), let__5390)


let__5392 := tmp7007
_ = let__5392

tmp7008 := Call(__e, PrimFunc(symshen_4comb), let__5392, let__5391)


ifres7004 = tmp7008


}

ifres7000 = ifres7004


}

let__5387 := ifres7000
_ = let__5387

tmp7036 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5387)


if True == tmp7036 {
tmp7011 := Call(__e, PrimFunc(symshen_4_5minus_6), V3110)


let__5394 := tmp7011
_ = let__5394

tmp7023 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5394)


var ifres7012 Obj

if True == tmp7023 {
tmp7013 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres7012 = tmp7013


} else {
tmp7014 := Call(__e, PrimFunc(symshen_4in_1_6), let__5394)


let__5395 := tmp7014
_ = let__5395

tmp7015 := Call(__e, PrimFunc(symshen_4_5log10_6), let__5395)


let__5396 := tmp7015
_ = let__5396

tmp7022 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5396)


var ifres7016 Obj

if True == tmp7022 {
tmp7017 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres7016 = tmp7017


} else {
tmp7018 := Call(__e, PrimFunc(symshen_4_5_1out), let__5396)


let__5397 := tmp7018
_ = let__5397

tmp7019 := Call(__e, PrimFunc(symshen_4in_1_6), let__5396)


let__5398 := tmp7019
_ = let__5398

tmp7021 := Call(__e, PrimFunc(symshen_4comb), let__5398, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(0))
__typedN1, __typedOK1 := TypedFloat64(let__5397)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := MakeInteger(0)
__typedArg1 := let__5397
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


ifres7016 = tmp7021


}

ifres7012 = ifres7016


}

let__5393 := ifres7012
_ = let__5393

tmp7034 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5393)


if True == tmp7034 {
tmp7024 := Call(__e, PrimFunc(symshen_4_5integer_6), V3110)


let__5400 := tmp7024
_ = let__5400

tmp7030 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5400)


var ifres7025 Obj

if True == tmp7030 {
tmp7026 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres7025 = tmp7026


} else {
tmp7027 := Call(__e, PrimFunc(symshen_4_5_1out), let__5400)


let__5401 := tmp7027
_ = let__5401

tmp7028 := Call(__e, PrimFunc(symshen_4in_1_6), let__5400)


let__5402 := tmp7028
_ = let__5402

tmp7029 := Call(__e, PrimFunc(symshen_4comb), let__5402, let__5401)


ifres7025 = tmp7029


}

let__5399 := ifres7025
_ = let__5399

tmp7032 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5399)


if True == tmp7032 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5399)
return
}


} else {
__e.Return(let__5393)
return
}


} else {
__e.Return(let__5387)
return
}


}, 1)

tmp7037 := Call(__e, ns2_1set, symshen_4_5log10_6, tmp6998)


_ = tmp7037

tmp7038 := MakeNative(func(__e *ControlFlow) {
V3127 := __e.Get(1)
_ = V3127
tmp7043 := Call(__e, PrimFunc(symshen_4hds_a_2), V3127, MakeInteger(101))


var ifres7039 Obj

if True == tmp7043 {
tmp7040 := Call(__e, PrimFunc(symtail), V3127)


let__5404 := tmp7040
_ = let__5404

tmp7041 := Call(__e, PrimFunc(symshen_4comb), let__5404, symshen_4skip)


ifres7039 = tmp7041


} else {
tmp7042 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres7039 = tmp7042


}

let__5403 := ifres7039
_ = let__5403

tmp7045 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5403)


if True == tmp7045 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5403)
return
}


}, 1)

tmp7046 := Call(__e, ns2_1set, symshen_4_5lowE_6, tmp7038)


_ = tmp7046

tmp7047 := MakeNative(func(__e *ControlFlow) {
V3130 := __e.Get(1)
_ = V3130
V3131 := __e.Get(2)
_ = V3131
tmp7048 := Call(__e, PrimFunc(symshen_4expt), MakeInteger(10), V3131)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_d) {
__typedN0, __typedOK0 := TypedFloat64(V3130)
__typedN1, __typedOK1 := TypedFloat64(tmp7048)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_d) {
return TypedMaterializeNumber((__typedN0 * __typedN1))
}}
__typedArg0 := V3130
__typedArg1 := tmp7048
return Call(__e, PrimFunc(sym_d), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp7049 := Call(__e, ns2_1set, symshen_4compute_1E, tmp7047)


_ = tmp7049

tmp7050 := MakeNative(func(__e *ControlFlow) {
V3132 := __e.Get(1)
_ = V3132
tmp7051 := Call(__e, PrimFunc(symshen_4_5whitespace_6), V3132)


let__5406 := tmp7051
_ = let__5406

tmp7061 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5406)


var ifres7052 Obj

if True == tmp7061 {
tmp7053 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres7052 = tmp7053


} else {
tmp7054 := Call(__e, PrimFunc(symshen_4in_1_6), let__5406)


let__5407 := tmp7054
_ = let__5407

tmp7055 := Call(__e, PrimFunc(symshen_4_5whitespaces_6), let__5407)


let__5408 := tmp7055
_ = let__5408

tmp7060 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5408)


var ifres7056 Obj

if True == tmp7060 {
tmp7057 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres7056 = tmp7057


} else {
tmp7058 := Call(__e, PrimFunc(symshen_4in_1_6), let__5408)


let__5409 := tmp7058
_ = let__5409

tmp7059 := Call(__e, PrimFunc(symshen_4comb), let__5409, symshen_4skip)


ifres7056 = tmp7059


}

ifres7052 = ifres7056


}

let__5405 := ifres7052
_ = let__5405

tmp7071 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5405)


if True == tmp7071 {
tmp7062 := Call(__e, PrimFunc(symshen_4_5whitespace_6), V3132)


let__5411 := tmp7062
_ = let__5411

tmp7067 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5411)


var ifres7063 Obj

if True == tmp7067 {
tmp7064 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres7063 = tmp7064


} else {
tmp7065 := Call(__e, PrimFunc(symshen_4in_1_6), let__5411)


let__5412 := tmp7065
_ = let__5412

tmp7066 := Call(__e, PrimFunc(symshen_4comb), let__5412, symshen_4skip)


ifres7063 = tmp7066


}

let__5410 := ifres7063
_ = let__5410

tmp7069 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5410)


if True == tmp7069 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5410)
return
}


} else {
__e.Return(let__5405)
return
}


}, 1)

tmp7072 := Call(__e, ns2_1set, symshen_4_5whitespaces_6, tmp7050)


_ = tmp7072

tmp7073 := MakeNative(func(__e *ControlFlow) {
V3141 := __e.Get(1)
_ = V3141
tmp7082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3141)
}
__typedArg0 := V3141
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7074 Obj

if True == tmp7082 {
tmp7075 := Call(__e, PrimFunc(symhead), V3141)


let__5414 := tmp7075
_ = let__5414

tmp7076 := Call(__e, PrimFunc(symtail), V3141)


let__5415 := tmp7076
_ = let__5415

tmp7080 := Call(__e, PrimFunc(symshen_4whitespace_2), let__5414)


var ifres7077 Obj

if True == tmp7080 {
tmp7078 := Call(__e, PrimFunc(symshen_4comb), let__5415, symshen_4skip)


ifres7077 = tmp7078


} else {
tmp7079 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres7077 = tmp7079


}

ifres7074 = ifres7077


} else {
tmp7081 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres7074 = tmp7081


}

let__5413 := ifres7074
_ = let__5413

tmp7084 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__5413)


if True == tmp7084 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__5413)
return
}


}, 1)

tmp7085 := Call(__e, ns2_1set, symshen_4_5whitespace_6, tmp7073)


_ = tmp7085

tmp7086 := MakeNative(func(__e *ControlFlow) {
V3147 := __e.Get(1)
_ = V3147
tmp7094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(32), V3147)
}
__typedArg0 := MakeInteger(32)
__typedArg1 := V3147
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7094 {
__e.Return(True)
return
} else {
tmp7092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(13), V3147)
}
__typedArg0 := MakeInteger(13)
__typedArg1 := V3147
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7092 {
__e.Return(True)
return
} else {
tmp7090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(10), V3147)
}
__typedArg0 := MakeInteger(10)
__typedArg1 := V3147
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7090 {
__e.Return(True)
return
} else {
tmp7088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(9), V3147)
}
__typedArg0 := MakeInteger(9)
__typedArg1 := V3147
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7088 {
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

tmp7095 := Call(__e, ns2_1set, symshen_4whitespace_2, tmp7086)


_ = tmp7095

tmp7096 := MakeNative(func(__e *ControlFlow) {
V3150 := __e.Get(1)
_ = V3150
V3151 := __e.Get(2)
_ = V3151
tmp7111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V3150)
}
__typedArg0 := V3150
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

var ifres7103 Obj

if True == tmp7111 {
tmp7109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V3150)
}
__typedArg0 := V3150
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

tmp7110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7109)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7109
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres7105 Obj

if True == tmp7110 {
tmp7107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3150, sympackage)
}
__typedArg0 := V3150
__typedArg1 := sympackage
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

tmp7108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7107)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7107
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

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
tmp7097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsource_1form_1handlers_d)
}
__typedArg0 := symshen_4_dsource_1form_1handlers_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp7098 := Call(__e, PrimFunc(symshen_4update_1assoc), V3150, V3151, tmp7097)


tmp7099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsource_1form_1handlers_d, tmp7098)
}
__typedArg0 := symshen_4_dsource_1form_1handlers_d
__typedArg1 := tmp7098
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7099

__e.Return(V3150)
return


} else {
tmp7100 := Call(__e, PrimFunc(symshen_4app), V3150, MakeString("\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("invalid source-form head: "))
__typedS1, __typedOK1 := TypedString(tmp7100)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("invalid source-form head: ")
__typedArg1 := tmp7100
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("invalid source-form head: "))
__typedS1, __typedOK1 := TypedString(tmp7100)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("invalid source-form head: ")
__typedArg1 := tmp7100
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 2)

tmp7112 := Call(__e, ns2_1set, symshen_4register_1source_1form, tmp7096)


_ = tmp7112

tmp7113 := MakeNative(func(__e *ControlFlow) {
V3152 := __e.Get(1)
_ = V3152
tmp7114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsource_1form_1handlers_d)
}
__typedArg0 := symshen_4_dsource_1form_1handlers_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp7115 := Call(__e, PrimFunc(symassoc), V3152, tmp7114)


tmp7116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsource_1form_1handlers_d)
}
__typedArg0 := symshen_4_dsource_1form_1handlers_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp7117 := Call(__e, PrimFunc(symremove), tmp7115, tmp7116)


tmp7118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsource_1form_1handlers_d, tmp7117)
}
__typedArg0 := symshen_4_dsource_1form_1handlers_d
__typedArg1 := tmp7117
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp7118

__e.Return(V3152)
return


}, 1)

tmp7119 := Call(__e, ns2_1set, symshen_4unregister_1source_1form, tmp7113)


_ = tmp7119

tmp7120 := MakeNative(func(__e *ControlFlow) {
V3159 := __e.Get(1)
_ = V3159
V3160 := __e.Get(2)
_ = V3160
tmp7123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3159)
}
__typedArg0 := V3159
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7123 {
tmp7121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3159)
}
__typedArg0 := V3159
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symassoc), tmp7121, V3160)
return


} else {
__e.Return(Nil)
return
}


}, 2)

tmp7124 := Call(__e, ns2_1set, symshen_4source_1form_1handler, tmp7120)


_ = tmp7124

tmp7125 := MakeNative(func(__e *ControlFlow) {
V3163 := __e.Get(1)
_ = V3163
V3164 := __e.Get(2)
_ = V3164
tmp7133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3164)
}
__typedArg0 := V3164
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7133 {
tmp7126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3164)
}
__typedArg0 := V3164
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7127 := Call(__e, tmp7126, V3163)


let__5416 := tmp7127
_ = let__5416

tmp7131 := Call(__e, PrimFunc(symshen_4source_1form_1list_2), let__5416)


if True == tmp7131 {
__e.Return(let__5416)
return
} else {
tmp7128 := Call(__e, PrimFunc(symshen_4app), V3163, MakeString("\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("source-form handler must return a list of forms: "))
__typedS1, __typedOK1 := TypedString(tmp7128)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("source-form handler must return a list of forms: ")
__typedArg1 := tmp7128
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("source-form handler must return a list of forms: "))
__typedS1, __typedOK1 := TypedString(tmp7128)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("source-form handler must return a list of forms: ")
__typedArg1 := tmp7128
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

tmp7134 := Call(__e, ns2_1set, symshen_4source_1form_1expansion, tmp7125)


_ = tmp7134

tmp7135 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V3170 := __self1
_ = V3170
tmp7140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3170)
}
__typedArg0 := Nil
__typedArg1 := V3170
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7140 {
__e.Return(True)
return
} else {
tmp7138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3170)
}
__typedArg0 := V3170
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7138 {
tmp7136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3170)
}
__typedArg0 := V3170
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4source_1form_1list_2) == __self {
__self1 = tmp7136
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4source_1form_1list_2), tmp7136)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp7141 := Call(__e, ns2_1set, symshen_4source_1form_1list_2, tmp7135)


_ = tmp7141

tmp7142 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V3171 := __self1
_ = V3171
tmp7178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3171)
}
__typedArg0 := Nil
__typedArg1 := V3171
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7178 {
__e.Return(Nil)
return
} else {
tmp7176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3171)
}
__typedArg0 := V3171
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7172 Obj

if True == tmp7176 {
tmp7174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3171)
}
__typedArg0 := V3171
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7175 := Call(__e, PrimFunc(symshen_4packaged_2), tmp7174)


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
tmp7143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3171)
}
__typedArg0 := V3171
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7144 := Call(__e, PrimFunc(symshen_4unpackage), tmp7143)


tmp7145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3171)
}
__typedArg0 := V3171
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7146 := Call(__e, PrimFunc(symappend), tmp7144, tmp7145)


if PrimFunc(symshen_4unpackage_emacroexpand) == __self {
__self1 = tmp7146
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4unpackage_emacroexpand), tmp7146)
return


} else {
tmp7170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3171)
}
__typedArg0 := V3171
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7170 {
tmp7147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3171)
}
__typedArg0 := V3171
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsource_1form_1handlers_d)
}
__typedArg0 := symshen_4_dsource_1form_1handlers_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp7149 := Call(__e, PrimFunc(symshen_4source_1form_1handler), tmp7147, tmp7148)


let__5417 := tmp7149
_ = let__5417

tmp7168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__5417)
}
__typedArg0 := let__5417
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7168 {
tmp7150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3171)
}
__typedArg0 := V3171
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7151 := Call(__e, PrimFunc(symshen_4source_1form_1expansion), tmp7150, let__5417)


tmp7152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3171)
}
__typedArg0 := V3171
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7153 := Call(__e, PrimFunc(symappend), tmp7151, tmp7152)


if PrimFunc(symshen_4unpackage_emacroexpand) == __self {
__self1 = tmp7153
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4unpackage_emacroexpand), tmp7153)
return


} else {
tmp7154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3171)
}
__typedArg0 := V3171
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7155 := Call(__e, PrimFunc(symmacroexpand), tmp7154)


let__5418 := tmp7155
_ = let__5418

tmp7166 := Call(__e, PrimFunc(symshen_4packaged_2), let__5418)


var ifres7161 Obj

if True == tmp7166 {
ifres7161 = True


} else {
tmp7163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dsource_1form_1handlers_d)
}
__typedArg0 := symshen_4_dsource_1form_1handlers_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp7164 := Call(__e, PrimFunc(symshen_4source_1form_1handler), let__5418, tmp7163)


tmp7165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7164)
}
__typedArg0 := tmp7164
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7162 Obj

if True == tmp7165 {
ifres7162 = True


} else {
ifres7162 = False


}

ifres7161 = ifres7162


}

if True == ifres7161 {
tmp7156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3171)
}
__typedArg0 := V3171
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__5418, tmp7156)
}
__typedArg0 := let__5418
__typedArg1 := tmp7156
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4unpackage_emacroexpand) == __self {
__self1 = tmp7157
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4unpackage_emacroexpand), tmp7157)
return


} else {
tmp7158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3171)
}
__typedArg0 := V3171
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7159 := Call(__e, PrimFunc(symshen_4unpackage_emacroexpand), tmp7158)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__5418, tmp7159)
}
__typedArg0 := let__5418
__typedArg1 := tmp7159
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

tmp7179 := Call(__e, ns2_1set, symshen_4unpackage_emacroexpand, tmp7142)


_ = tmp7179

tmp7180 := MakeNative(func(__e *ControlFlow) {
V3176 := __e.Get(1)
_ = V3176
tmp7195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3176)
}
__typedArg0 := V3176
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7182 Obj

if True == tmp7195 {
tmp7193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3176)
}
__typedArg0 := V3176
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sympackage, tmp7193)
}
__typedArg0 := sympackage
__typedArg1 := tmp7193
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7184 Obj

if True == tmp7194 {
tmp7191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3176)
}
__typedArg0 := V3176
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7191)
}
__typedArg0 := tmp7191
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7186 Obj

if True == tmp7192 {
tmp7188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3176)
}
__typedArg0 := V3176
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7188)
}
__typedArg0 := tmp7188
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7189)
}
__typedArg0 := tmp7189
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

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
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp7196 := Call(__e, ns2_1set, symshen_4packaged_2, tmp7180)


_ = tmp7196

tmp7197 := MakeNative(func(__e *ControlFlow) {
V3179 := __e.Get(1)
_ = V3179
tmp7254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3179)
}
__typedArg0 := V3179
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7236 Obj

if True == tmp7254 {
tmp7252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3179)
}
__typedArg0 := V3179
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

var ifres7238 Obj

if True == tmp7253 {
tmp7250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3179)
}
__typedArg0 := V3179
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7250)
}
__typedArg0 := tmp7250
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7240 Obj

if True == tmp7251 {
tmp7247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3179)
}
__typedArg0 := V3179
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7247)
}
__typedArg0 := tmp7247
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symnull, tmp7248)
}
__typedArg0 := symnull
__typedArg1 := tmp7248
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7242 Obj

if True == tmp7249 {
tmp7244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3179)
}
__typedArg0 := V3179
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7244)
}
__typedArg0 := tmp7244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7245)
}
__typedArg0 := tmp7245
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

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
tmp7198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3179)
}
__typedArg0 := V3179
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7198)
}
__typedArg0 := tmp7198
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7199)
}
__typedArg0 := tmp7199
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return


} else {
tmp7234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3179)
}
__typedArg0 := V3179
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7221 Obj

if True == tmp7234 {
tmp7232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3179)
}
__typedArg0 := V3179
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sympackage, tmp7232)
}
__typedArg0 := sympackage
__typedArg1 := tmp7232
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7223 Obj

if True == tmp7233 {
tmp7230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3179)
}
__typedArg0 := V3179
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7230)
}
__typedArg0 := tmp7230
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7225 Obj

if True == tmp7231 {
tmp7227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3179)
}
__typedArg0 := V3179
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7228)
}
__typedArg0 := tmp7228
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

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
tmp7200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3179)
}
__typedArg0 := V3179
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7200)
}
__typedArg0 := tmp7200
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7201)
}
__typedArg0 := tmp7201
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7203 := Call(__e, PrimFunc(symeval), tmp7202)


let__5419 := tmp7203
_ = let__5419

tmp7204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3179)
}
__typedArg0 := V3179
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7204)
}
__typedArg0 := tmp7204
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(tmp7205)
}
__typedArg0 := tmp7205
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp7207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3179)
}
__typedArg0 := V3179
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7207)
}
__typedArg0 := tmp7207
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7208)
}
__typedArg0 := tmp7208
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7210 := Call(__e, PrimFunc(symshen_4package_1symbols), tmp7206, let__5419, tmp7209)


let__5420 := tmp7210
_ = let__5420

tmp7211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3179)
}
__typedArg0 := V3179
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7211)
}
__typedArg0 := tmp7211
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7213 := Call(__e, PrimFunc(symshen_4record_1external), tmp7212, let__5419)


let__5421 := tmp7213
_ = let__5421

tmp7214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3179)
}
__typedArg0 := V3179
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7214)
}
__typedArg0 := tmp7214
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3179)
}
__typedArg0 := V3179
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7216)
}
__typedArg0 := tmp7216
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7217)
}
__typedArg0 := tmp7217
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7219 := Call(__e, PrimFunc(symshen_4record_1internal), tmp7215, let__5419, tmp7218)


let__5422 := tmp7219
_ = let__5422

__e.Return(let__5420)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4unpackage)
return
}


}


}, 1)

tmp7255 := Call(__e, ns2_1set, symshen_4unpackage, tmp7197)


_ = tmp7255

tmp7256 := MakeNative(func(__e *ControlFlow) {
V3184 := __e.Get(1)
_ = V3184
V3185 := __e.Get(2)
_ = V3185
V3186 := __e.Get(3)
_ = V3186
tmp7257 := MakeNative(func(__e *ControlFlow) {
tmp7258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V3184, symshen_4internal_1symbols, tmp7258)
return


}, 0)

tmp7259 := MakeNative(func(__e *ControlFlow) {
Z3188 := __e.Get(1)
_ = Z3188
__e.Return(Nil)
return
}, 1)

tmp7260 := Call(__e, try_1catch, tmp7257, tmp7259)


let__5423 := tmp7260
_ = let__5423

tmp7261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3184)
}
__typedArg0 := V3184
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp7262 := Call(__e, PrimFunc(symshen_4internal_1symbols), tmp7261, V3185, V3186)


let__5424 := tmp7262
_ = let__5424

tmp7263 := Call(__e, PrimFunc(symunion), let__5424, let__5423)


tmp7264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V3184, symshen_4internal_1symbols, tmp7263, tmp7264)
return


}, 3)

tmp7265 := Call(__e, ns2_1set, symshen_4record_1internal, tmp7256)


_ = tmp7265

tmp7266 := MakeNative(func(__e *ControlFlow) {
V3196 := __e.Get(1)
_ = V3196
V3197 := __e.Get(2)
_ = V3197
V3198 := __e.Get(3)
_ = V3198
tmp7275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3198)
}
__typedArg0 := V3198
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7275 {
tmp7267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3198)
}
__typedArg0 := V3198
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7268 := Call(__e, PrimFunc(symshen_4internal_1symbols), V3196, V3197, tmp7267)


tmp7269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3198)
}
__typedArg0 := V3198
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7270 := Call(__e, PrimFunc(symshen_4internal_1symbols), V3196, V3197, tmp7269)


__e.TailApply(PrimFunc(symunion), tmp7268, tmp7270)
return


} else {
tmp7273 := Call(__e, PrimFunc(symshen_4internal_2), V3198, V3196, V3197)


if True == tmp7273 {
tmp7271 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V3196, V3198)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7271, Nil)
}
__typedArg0 := tmp7271
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

tmp7276 := Call(__e, ns2_1set, symshen_4internal_1symbols, tmp7266)


_ = tmp7276

tmp7277 := MakeNative(func(__e *ControlFlow) {
V3199 := __e.Get(1)
_ = V3199
V3200 := __e.Get(2)
_ = V3200
tmp7278 := MakeNative(func(__e *ControlFlow) {
tmp7279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V3199, symshen_4external_1symbols, tmp7279)
return


}, 0)

tmp7280 := MakeNative(func(__e *ControlFlow) {
Z3202 := __e.Get(1)
_ = Z3202
__e.Return(Nil)
return
}, 1)

tmp7281 := Call(__e, try_1catch, tmp7278, tmp7280)


let__5425 := tmp7281
_ = let__5425

tmp7282 := Call(__e, PrimFunc(symunion), V3200, let__5425)


tmp7283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V3199, symshen_4external_1symbols, tmp7282, tmp7283)
return


}, 2)

tmp7284 := Call(__e, ns2_1set, symshen_4record_1external, tmp7277)


_ = tmp7284

tmp7285 := MakeNative(func(__e *ControlFlow) {
V3207 := __e.Get(1)
_ = V3207
V3208 := __e.Get(2)
_ = V3208
V3209 := __e.Get(3)
_ = V3209
tmp7290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3209)
}
__typedArg0 := V3209
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7290 {
tmp7286 := MakeNative(func(__e *ControlFlow) {
Z3210 := __e.Get(1)
_ = Z3210
__e.TailApply(PrimFunc(symshen_4package_1symbols), V3207, V3208, Z3210)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp7286, V3209)
return


} else {
tmp7288 := Call(__e, PrimFunc(symshen_4internal_2), V3209, V3207, V3208)


if True == tmp7288 {
__e.TailApply(PrimFunc(symshen_4intern_1in_1package), V3207, V3209)
return
} else {
__e.Return(V3209)
return
}


}


}, 3)

tmp7291 := Call(__e, ns2_1set, symshen_4package_1symbols, tmp7285)


_ = tmp7291

tmp7292 := MakeNative(func(__e *ControlFlow) {
V3211 := __e.Get(1)
_ = V3211
V3212 := __e.Get(2)
_ = V3212
tmp7293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3212)
}
__typedArg0 := V3212
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp7294 := Call(__e, PrimFunc(sym_8s), MakeString("."), tmp7293)


tmp7295 := Call(__e, PrimFunc(sym_8s), V3211, tmp7294)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(tmp7295)
}
__typedArg0 := tmp7295
return Call(__e, PrimFunc(symintern), __typedArg0)
})())
return


}, 2)

tmp7296 := Call(__e, ns2_1set, symshen_4intern_1in_1package, tmp7292)


_ = tmp7296

tmp7297 := MakeNative(func(__e *ControlFlow) {
V3213 := __e.Get(1)
_ = V3213
V3214 := __e.Get(2)
_ = V3214
V3215 := __e.Get(3)
_ = V3215
tmp7327 := Call(__e, PrimFunc(symelement_2), V3213, V3215)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7327)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7327
return Call(__e, PrimFunc(symnot), __typedArg0)
})() {
tmp7324 := Call(__e, PrimFunc(symshen_4sng_2), V3213)


tmp7325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7324)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7324
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres7299 Obj

if True == tmp7325 {
tmp7322 := Call(__e, PrimFunc(symshen_4dbl_2), V3213)


tmp7323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7322)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7322
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres7301 Obj

if True == tmp7323 {
tmp7321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V3213)
}
__typedArg0 := V3213
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

var ifres7303 Obj

if True == tmp7321 {
tmp7319 := Call(__e, PrimFunc(symshen_4sysfunc_2), V3213)


tmp7320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7319)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7319
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres7305 Obj

if True == tmp7320 {
tmp7317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V3213)
}
__typedArg0 := V3213
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

tmp7318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7317)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7317
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres7307 Obj

if True == tmp7318 {
tmp7314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3213)
}
__typedArg0 := V3213
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp7315 := Call(__e, PrimFunc(symshen_4internal_1to_1shen_2), tmp7314)


tmp7316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7315)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7315
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres7309 Obj

if True == tmp7316 {
tmp7311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3213)
}
__typedArg0 := V3213
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp7312 := Call(__e, PrimFunc(symshen_4internal_1to_1P_2), V3214, tmp7311)


tmp7313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7312)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7312
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

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

tmp7329 := Call(__e, ns2_1set, symshen_4internal_2, tmp7297)


_ = tmp7329

tmp7330 := MakeNative(func(__e *ControlFlow) {
V3220 := __e.Get(1)
_ = V3220
tmp7384 := Call(__e, PrimFunc(symshen_4_7string_2), V3220)


var ifres7332 Obj

if True == tmp7384 {
tmp7382 := Call(__e, PrimFunc(symhdstr), V3220)


tmp7383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("s"), tmp7382)
}
__typedArg0 := MakeString("s")
__typedArg1 := tmp7382
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7334 Obj

if True == tmp7383 {
tmp7381 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3220)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3220
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres7336 Obj

if True == tmp7381 {
tmp7378 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3220)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3220
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp7379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("h"), tmp7378)
}
__typedArg0 := MakeString("h")
__typedArg1 := tmp7378
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7338 Obj

if True == tmp7379 {
tmp7376 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3220)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3220)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3220
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres7340 Obj

if True == tmp7376 {
tmp7372 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3220)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3220)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3220
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp7373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("e"), tmp7372)
}
__typedArg0 := MakeString("e")
__typedArg1 := tmp7372
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7342 Obj

if True == tmp7373 {
tmp7369 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3220)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(TypedStringTailValue(__typedS0))))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3220)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3220)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3220
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres7344 Obj

if True == tmp7369 {
tmp7364 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3220)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(TypedStringTailValue(__typedS0))))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3220)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3220)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3220
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp7365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("n"), tmp7364)
}
__typedArg0 := MakeString("n")
__typedArg1 := tmp7364
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7346 Obj

if True == tmp7365 {
tmp7360 := Call(__e, PrimFunc(symshen_4_7string_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3220)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(TypedStringTailValue(TypedStringTailValue(__typedS0)))))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3220)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(TypedStringTailValue(__typedS0))))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3220)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3220)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3220
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


var ifres7348 Obj

if True == tmp7360 {
tmp7354 := Call(__e, PrimFunc(symhdstr), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3220)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(TypedStringTailValue(TypedStringTailValue(__typedS0)))))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3220)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(TypedStringTailValue(__typedS0))))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3220)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(TypedStringTailValue(__typedS0)))
}}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3220)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3220
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


tmp7355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("."), tmp7354)
}
__typedArg0 := MakeString(".")
__typedArg1 := tmp7354
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp7385 := Call(__e, ns2_1set, symshen_4internal_1to_1shen_2, tmp7330)


_ = tmp7385

tmp7386 := MakeNative(func(__e *ControlFlow) {
V3221 := __e.Get(1)
_ = V3221
tmp7387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp7388 := Call(__e, PrimFunc(symget), symshen, symshen_4external_1symbols, tmp7387)


__e.TailApply(PrimFunc(symelement_2), V3221, tmp7388)
return


}, 1)

tmp7389 := Call(__e, ns2_1set, symshen_4sysfunc_2, tmp7386)


_ = tmp7389

tmp7390 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V3229 := __self1
_ = V3229
V3230 := __self2
_ = V3230
tmp7411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V3229)
}
__typedArg0 := MakeString("")
__typedArg1 := V3229
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7404 Obj

if True == tmp7411 {
tmp7410 := Call(__e, PrimFunc(symshen_4_7string_2), V3230)


var ifres7406 Obj

if True == tmp7410 {
tmp7408 := Call(__e, PrimFunc(symhdstr), V3230)


tmp7409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("."), tmp7408)
}
__typedArg0 := MakeString(".")
__typedArg1 := tmp7408
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
__e.Return(True)
return
} else {
tmp7402 := Call(__e, PrimFunc(symshen_4_7string_2), V3229)


var ifres7394 Obj

if True == tmp7402 {
tmp7401 := Call(__e, PrimFunc(symshen_4_7string_2), V3230)


var ifres7396 Obj

if True == tmp7401 {
tmp7398 := Call(__e, PrimFunc(symhdstr), V3229)


tmp7399 := Call(__e, PrimFunc(symhdstr), V3230)


tmp7400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp7398, tmp7399)
}
__typedArg0 := tmp7398
__typedArg1 := tmp7399
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
tmp7391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3229)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3229
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()

if PrimFunc(symshen_4internal_1to_1P_2) == __self {
__self1, __self2 = tmp7391, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3230)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3230
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})()
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4internal_1to_1P_2), tmp7391, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3230)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3230
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())
return


} else {
__e.Return(False)
return
}


}


}, 2)

tmp7412 := Call(__e, ns2_1set, symshen_4internal_1to_1P_2, tmp7390)


_ = tmp7412

tmp7413 := MakeNative(func(__e *ControlFlow) {
V3233 := __e.Get(1)
_ = V3233
V3234 := __e.Get(2)
_ = V3234
tmp7434 := Call(__e, PrimFunc(symelement_2), V3233, V3234)


if True == tmp7434 {
__e.Return(V3233)
return
} else {
tmp7432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3233)
}
__typedArg0 := V3233
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7428 Obj

if True == tmp7432 {
tmp7430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3233)
}
__typedArg0 := V3233
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcond, tmp7430)
}
__typedArg0 := symcond
__typedArg1 := tmp7430
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
tmp7414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3233)
}
__typedArg0 := V3233
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7415 := Call(__e, PrimFunc(symshen_4process_1cond_1clauses), tmp7414, V3234)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcond, tmp7415)
}
__typedArg0 := symcond
__typedArg1 := tmp7415
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3233)
}
__typedArg0 := V3233
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7422 Obj

if True == tmp7426 {
tmp7424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3233)
}
__typedArg0 := V3233
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7425 := Call(__e, PrimFunc(symshen_4non_1application_2), tmp7424)


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
tmp7416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3233)
}
__typedArg0 := V3233
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4special_1case), tmp7416, V3233, V3234)
return


} else {
tmp7420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3233)
}
__typedArg0 := V3233
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7420 {
tmp7417 := MakeNative(func(__e *ControlFlow) {
Z3235 := __e.Get(1)
_ = Z3235
__e.TailApply(PrimFunc(symshen_4process_1applications), Z3235, V3234)
return
}, 1)

tmp7418 := Call(__e, PrimFunc(symmap), tmp7417, V3233)


__e.TailApply(PrimFunc(symshen_4process_1application), tmp7418, V3234)
return


} else {
__e.Return(V3233)
return
}


}


}


}


}, 2)

tmp7435 := Call(__e, ns2_1set, symshen_4process_1applications, tmp7413)


_ = tmp7435

tmp7436 := MakeNative(func(__e *ControlFlow) {
V3238 := __e.Get(1)
_ = V3238
tmp7446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, V3238)
}
__typedArg0 := symdefine
__typedArg1 := V3238
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7446 {
__e.Return(True)
return
} else {
tmp7444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, V3238)
}
__typedArg0 := symdefun
__typedArg1 := V3238
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7444 {
__e.Return(True)
return
} else {
tmp7442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsynonyms, V3238)
}
__typedArg0 := symsynonyms
__typedArg1 := V3238
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7442 {
__e.Return(True)
return
} else {
tmp7440 := Call(__e, PrimFunc(symshen_4special_2), V3238)


if True == tmp7440 {
__e.Return(True)
return
} else {
tmp7438 := Call(__e, PrimFunc(symshen_4extraspecial_2), V3238)


if True == tmp7438 {
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

tmp7447 := Call(__e, ns2_1set, symshen_4non_1application_2, tmp7436)


_ = tmp7447

tmp7448 := MakeNative(func(__e *ControlFlow) {
V3243 := __e.Get(1)
_ = V3243
V3244 := __e.Get(2)
_ = V3244
V3245 := __e.Get(3)
_ = V3245
tmp7690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, V3243)
}
__typedArg0 := symlambda
__typedArg1 := V3243
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7668 Obj

if True == tmp7690 {
tmp7689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7670 Obj

if True == tmp7689 {
tmp7687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp7687)
}
__typedArg0 := symlambda
__typedArg1 := tmp7687
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7672 Obj

if True == tmp7688 {
tmp7685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7685)
}
__typedArg0 := tmp7685
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7674 Obj

if True == tmp7686 {
tmp7682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7682)
}
__typedArg0 := tmp7682
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7683)
}
__typedArg0 := tmp7683
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7676 Obj

if True == tmp7684 {
tmp7678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7679)
}
__typedArg0 := tmp7679
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7680)
}
__typedArg0 := Nil
__typedArg1 := tmp7680
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
tmp7449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7449)
}
__typedArg0 := tmp7449
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7451)
}
__typedArg0 := tmp7451
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7452)
}
__typedArg0 := tmp7452
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7454 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7453, V3245)


tmp7455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7454, Nil)
}
__typedArg0 := tmp7454
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7450, tmp7455)
}
__typedArg0 := tmp7450
__typedArg1 := tmp7455
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp7456)
}
__typedArg0 := symlambda
__typedArg1 := tmp7456
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, V3243)
}
__typedArg0 := symlet
__typedArg1 := V3243
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7637 Obj

if True == tmp7666 {
tmp7665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7639 Obj

if True == tmp7665 {
tmp7663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp7663)
}
__typedArg0 := symlet
__typedArg1 := tmp7663
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7641 Obj

if True == tmp7664 {
tmp7661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7661)
}
__typedArg0 := tmp7661
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7643 Obj

if True == tmp7662 {
tmp7658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7658)
}
__typedArg0 := tmp7658
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7659)
}
__typedArg0 := tmp7659
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7645 Obj

if True == tmp7660 {
tmp7654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7654)
}
__typedArg0 := tmp7654
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7655)
}
__typedArg0 := tmp7655
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7656)
}
__typedArg0 := tmp7656
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7647 Obj

if True == tmp7657 {
tmp7649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7649)
}
__typedArg0 := tmp7649
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7650)
}
__typedArg0 := tmp7650
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7651)
}
__typedArg0 := tmp7651
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7652)
}
__typedArg0 := Nil
__typedArg1 := tmp7652
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
tmp7457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7457)
}
__typedArg0 := tmp7457
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7459)
}
__typedArg0 := tmp7459
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7460)
}
__typedArg0 := tmp7460
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7462 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7461, V3245)


tmp7463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7463)
}
__typedArg0 := tmp7463
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7464)
}
__typedArg0 := tmp7464
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7465)
}
__typedArg0 := tmp7465
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7467 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7466, V3245)


tmp7468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7467, Nil)
}
__typedArg0 := tmp7467
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7462, tmp7468)
}
__typedArg0 := tmp7462
__typedArg1 := tmp7468
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7458, tmp7469)
}
__typedArg0 := tmp7458
__typedArg1 := tmp7469
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp7470)
}
__typedArg0 := symlet
__typedArg1 := tmp7470
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, V3243)
}
__typedArg0 := symdefun
__typedArg1 := V3243
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7606 Obj

if True == tmp7635 {
tmp7634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7608 Obj

if True == tmp7634 {
tmp7632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp7632)
}
__typedArg0 := symdefun
__typedArg1 := tmp7632
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7610 Obj

if True == tmp7633 {
tmp7630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7630)
}
__typedArg0 := tmp7630
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7612 Obj

if True == tmp7631 {
tmp7627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7627)
}
__typedArg0 := tmp7627
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7628)
}
__typedArg0 := tmp7628
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7614 Obj

if True == tmp7629 {
tmp7623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7623)
}
__typedArg0 := tmp7623
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7624)
}
__typedArg0 := tmp7624
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7625)
}
__typedArg0 := tmp7625
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7616 Obj

if True == tmp7626 {
tmp7618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7618)
}
__typedArg0 := tmp7618
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7619)
}
__typedArg0 := tmp7619
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7620)
}
__typedArg0 := tmp7620
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7621)
}
__typedArg0 := Nil
__typedArg1 := tmp7621
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
__e.Return(V3244)
return
} else {
tmp7604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, V3243)
}
__typedArg0 := symdefine
__typedArg1 := V3243
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7582 Obj

if True == tmp7604 {
tmp7603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7584 Obj

if True == tmp7603 {
tmp7601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp7601)
}
__typedArg0 := symdefine
__typedArg1 := tmp7601
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7586 Obj

if True == tmp7602 {
tmp7599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7599)
}
__typedArg0 := tmp7599
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7588 Obj

if True == tmp7600 {
tmp7596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7596)
}
__typedArg0 := tmp7596
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7597)
}
__typedArg0 := tmp7597
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7590 Obj

if True == tmp7598 {
tmp7592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7592)
}
__typedArg0 := tmp7592
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7593)
}
__typedArg0 := tmp7593
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, tmp7594)
}
__typedArg0 := sym_i
__typedArg1 := tmp7594
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
tmp7471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7471)
}
__typedArg0 := tmp7471
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7473)
}
__typedArg0 := tmp7473
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7475)
}
__typedArg0 := tmp7475
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7476)
}
__typedArg0 := tmp7476
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7478 := Call(__e, PrimFunc(symshen_4process_1after_1type), tmp7474, tmp7477, V3245)


tmp7479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, tmp7478)
}
__typedArg0 := sym_i
__typedArg1 := tmp7478
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7472, tmp7479)
}
__typedArg0 := tmp7472
__typedArg1 := tmp7479
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp7480)
}
__typedArg0 := symdefine
__typedArg1 := tmp7480
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, V3243)
}
__typedArg0 := symdefine
__typedArg1 := V3243
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7569 Obj

if True == tmp7580 {
tmp7579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7571 Obj

if True == tmp7579 {
tmp7577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp7577)
}
__typedArg0 := symdefine
__typedArg1 := tmp7577
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7573 Obj

if True == tmp7578 {
tmp7575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7575)
}
__typedArg0 := tmp7575
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

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
tmp7481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7481)
}
__typedArg0 := tmp7481
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7483 := MakeNative(func(__e *ControlFlow) {
Z3246 := __e.Get(1)
_ = Z3246
__e.TailApply(PrimFunc(symshen_4process_1applications), Z3246, V3245)
return
}, 1)

tmp7484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7484)
}
__typedArg0 := tmp7484
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7486 := Call(__e, PrimFunc(symmap), tmp7483, tmp7485)


tmp7487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7482, tmp7486)
}
__typedArg0 := tmp7482
__typedArg1 := tmp7486
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp7487)
}
__typedArg0 := symdefine
__typedArg1 := tmp7487
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsynonyms, V3243)
}
__typedArg0 := symsynonyms
__typedArg1 := V3243
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7567 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsynonyms, V3244)
}
__typedArg0 := symsynonyms
__typedArg1 := V3244
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
} else {
tmp7565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtype, V3243)
}
__typedArg0 := symtype
__typedArg1 := V3243
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7543 Obj

if True == tmp7565 {
tmp7564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7545 Obj

if True == tmp7564 {
tmp7562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtype, tmp7562)
}
__typedArg0 := symtype
__typedArg1 := tmp7562
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7547 Obj

if True == tmp7563 {
tmp7560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7560)
}
__typedArg0 := tmp7560
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7549 Obj

if True == tmp7561 {
tmp7557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7557)
}
__typedArg0 := tmp7557
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7558)
}
__typedArg0 := tmp7558
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7551 Obj

if True == tmp7559 {
tmp7553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7554)
}
__typedArg0 := tmp7554
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7555)
}
__typedArg0 := Nil
__typedArg1 := tmp7555
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
tmp7488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7488)
}
__typedArg0 := tmp7488
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7490 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7489, V3245)


tmp7491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7491)
}
__typedArg0 := tmp7491
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7490, tmp7492)
}
__typedArg0 := tmp7490
__typedArg1 := tmp7492
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp7493)
}
__typedArg0 := symtype
__typedArg1 := tmp7493
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput_7, V3243)
}
__typedArg0 := syminput_7
__typedArg1 := V3243
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7519 Obj

if True == tmp7541 {
tmp7540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7521 Obj

if True == tmp7540 {
tmp7538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput_7, tmp7538)
}
__typedArg0 := syminput_7
__typedArg1 := tmp7538
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7523 Obj

if True == tmp7539 {
tmp7536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7536)
}
__typedArg0 := tmp7536
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7525 Obj

if True == tmp7537 {
tmp7533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7533)
}
__typedArg0 := tmp7533
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7534)
}
__typedArg0 := tmp7534
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7527 Obj

if True == tmp7535 {
tmp7529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7529)
}
__typedArg0 := tmp7529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7530)
}
__typedArg0 := tmp7530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7531)
}
__typedArg0 := Nil
__typedArg1 := tmp7531
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
tmp7494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7494)
}
__typedArg0 := tmp7494
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7496)
}
__typedArg0 := tmp7496
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7497)
}
__typedArg0 := tmp7497
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7499 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7498, V3245)


tmp7500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7499, Nil)
}
__typedArg0 := tmp7499
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7495, tmp7500)
}
__typedArg0 := tmp7495
__typedArg1 := tmp7500
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput_7, tmp7501)
}
__typedArg0 := syminput_7
__typedArg1 := tmp7501
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7513 Obj

if True == tmp7517 {
tmp7515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7516 := Call(__e, PrimFunc(symshen_4special_2), tmp7515)


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
tmp7502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7503 := MakeNative(func(__e *ControlFlow) {
Z3247 := __e.Get(1)
_ = Z3247
__e.TailApply(PrimFunc(symshen_4process_1applications), Z3247, V3245)
return
}, 1)

tmp7504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7505 := Call(__e, PrimFunc(symmap), tmp7503, tmp7504)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7502, tmp7505)
}
__typedArg0 := tmp7502
__typedArg1 := tmp7505
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7507 Obj

if True == tmp7511 {
tmp7509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3244)
}
__typedArg0 := V3244
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7510 := Call(__e, PrimFunc(symshen_4extraspecial_2), tmp7509)


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
__e.Return(V3244)
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

tmp7691 := Call(__e, ns2_1set, symshen_4special_1case, tmp7448)


_ = tmp7691

tmp7692 := MakeNative(func(__e *ControlFlow) {
V3250 := __e.Get(1)
_ = V3250
V3251 := __e.Get(2)
_ = V3251
tmp7722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3250)
}
__typedArg0 := Nil
__typedArg1 := V3250
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7722 {
__e.Return(Nil)
return
} else {
tmp7720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3250)
}
__typedArg0 := V3250
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7705 Obj

if True == tmp7720 {
tmp7718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3250)
}
__typedArg0 := V3250
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7718)
}
__typedArg0 := tmp7718
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7707 Obj

if True == tmp7719 {
tmp7715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3250)
}
__typedArg0 := V3250
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7715)
}
__typedArg0 := tmp7715
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7716)
}
__typedArg0 := tmp7716
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7709 Obj

if True == tmp7717 {
tmp7711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3250)
}
__typedArg0 := V3250
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7711)
}
__typedArg0 := tmp7711
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7712)
}
__typedArg0 := tmp7712
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7713)
}
__typedArg0 := Nil
__typedArg1 := tmp7713
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
tmp7693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3250)
}
__typedArg0 := V3250
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7693)
}
__typedArg0 := tmp7693
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7695 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7694, V3251)


tmp7696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3250)
}
__typedArg0 := V3250
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7696)
}
__typedArg0 := tmp7696
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7697)
}
__typedArg0 := tmp7697
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7699 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7698, V3251)


tmp7700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7699, Nil)
}
__typedArg0 := tmp7699
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7695, tmp7700)
}
__typedArg0 := tmp7695
__typedArg1 := tmp7700
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3250)
}
__typedArg0 := V3250
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7703 := Call(__e, PrimFunc(symshen_4process_1cond_1clauses), tmp7702, V3251)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7701, tmp7703)
}
__typedArg0 := tmp7701
__typedArg1 := tmp7703
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4process_1cond_1clauses)
return
}


}


}, 2)

tmp7723 := Call(__e, ns2_1set, symshen_4process_1cond_1clauses, tmp7692)


_ = tmp7723

tmp7724 := MakeNative(func(__e *ControlFlow) {
V3254 := __e.Get(1)
_ = V3254
V3255 := __e.Get(2)
_ = V3255
V3256 := __e.Get(3)
_ = V3256
tmp7740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3255)
}
__typedArg0 := V3255
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7736 Obj

if True == tmp7740 {
tmp7738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3255)
}
__typedArg0 := V3255
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, tmp7738)
}
__typedArg0 := sym_j
__typedArg1 := tmp7738
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
tmp7725 := MakeNative(func(__e *ControlFlow) {
Z3257 := __e.Get(1)
_ = Z3257
__e.TailApply(PrimFunc(symshen_4process_1applications), Z3257, V3256)
return
}, 1)

tmp7726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3255)
}
__typedArg0 := V3255
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7727 := Call(__e, PrimFunc(symmap), tmp7725, tmp7726)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, tmp7727)
}
__typedArg0 := sym_j
__typedArg1 := tmp7727
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3255)
}
__typedArg0 := V3255
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7734 {
tmp7728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3255)
}
__typedArg0 := V3255
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3255)
}
__typedArg0 := V3255
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7730 := Call(__e, PrimFunc(symshen_4process_1after_1type), V3254, tmp7729, V3256)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7728, tmp7730)
}
__typedArg0 := tmp7728
__typedArg1 := tmp7730
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp7731 := Call(__e, PrimFunc(symshen_4app), V3254, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing } in "))
__typedS1, __typedOK1 := TypedString(tmp7731)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing } in ")
__typedArg1 := tmp7731
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing } in "))
__typedS1, __typedOK1 := TypedString(tmp7731)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing } in ")
__typedArg1 := tmp7731
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}


}, 3)

tmp7741 := Call(__e, ns2_1set, symshen_4process_1after_1type, tmp7724)


_ = tmp7741

tmp7742 := MakeNative(func(__e *ControlFlow) {
V3258 := __e.Get(1)
_ = V3258
V3259 := __e.Get(2)
_ = V3259
tmp7785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3258)
}
__typedArg0 := V3258
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp7785 {
tmp7743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3258)
}
__typedArg0 := V3258
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7744 := Call(__e, PrimFunc(symarity), tmp7743)


let__5426 := tmp7744
_ = let__5426

tmp7745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3258)
}
__typedArg0 := V3258
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7746 := Call(__e, PrimFunc(symlength), tmp7745)


let__5427 := tmp7746
_ = let__5427

tmp7783 := Call(__e, PrimFunc(symelement_2), V3258, V3259)


if True == tmp7783 {
__e.Return(V3258)
return
} else {
tmp7780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3258)
}
__typedArg0 := V3258
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7781 := Call(__e, PrimFunc(symshen_4shen_1call_2), tmp7780)


if True == tmp7781 {
__e.Return(V3258)
return
} else {
tmp7778 := Call(__e, PrimFunc(symshen_4foreign_2), V3258)


if True == tmp7778 {
__e.TailApply(PrimFunc(symshen_4unpack_1foreign), V3258)
return
} else {
tmp7776 := Call(__e, PrimFunc(symshen_4fn_1call_2), V3258)


if True == tmp7776 {
__e.TailApply(PrimFunc(symshen_4fn_1call), V3258)
return
} else {
tmp7774 := Call(__e, PrimFunc(symshen_4zero_1place_2), V3258)


if True == tmp7774 {
__e.Return(V3258)
return
} else {
tmp7771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3258)
}
__typedArg0 := V3258
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7772 := Call(__e, PrimFunc(symshen_4undefined_1f_2), tmp7771, let__5426)


if True == tmp7772 {
tmp7747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3258)
}
__typedArg0 := V3258
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7747, Nil)
}
__typedArg0 := tmp7747
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp7748)
}
__typedArg0 := symfn
__typedArg1 := tmp7748
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3258)
}
__typedArg0 := V3258
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7749, tmp7750)
}
__typedArg0 := tmp7749
__typedArg1 := tmp7750
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp7751)
return


} else {
tmp7768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3258)
}
__typedArg0 := V3258
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp7768)
}
__typedArg0 := tmp7768
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp7769 {
__e.TailApply(PrimFunc(symshen_4simple_1curry), V3258)
return
} else {
tmp7765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3258)
}
__typedArg0 := V3258
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7766 := Call(__e, PrimFunc(symshen_4application_2), tmp7765)


if True == tmp7766 {
__e.TailApply(PrimFunc(symshen_4simple_1curry), V3258)
return
} else {
tmp7762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3258)
}
__typedArg0 := V3258
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7763 := Call(__e, PrimFunc(symshen_4partial_1application_d_2), tmp7762, let__5426, let__5427)


if True == tmp7763 {
__e.TailApply(PrimFunc(symshen_4lambda_1function), V3258, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(let__5426)
__typedN1, __typedOK1 := TypedFloat64(let__5427)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := let__5426
__typedArg1 := let__5427
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())
return


} else {
tmp7759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3258)
}
__typedArg0 := V3258
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7760 := Call(__e, PrimFunc(symshen_4overapplication_2), tmp7759, let__5426, let__5427)


if True == tmp7760 {
tmp7753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3258)
}
__typedArg0 := V3258
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7753, Nil)
}
__typedArg0 := tmp7753
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp7754)
}
__typedArg0 := symfn
__typedArg1 := tmp7754
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3258)
}
__typedArg0 := V3258
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7755, tmp7756)
}
__typedArg0 := tmp7755
__typedArg1 := tmp7756
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp7757)
return


} else {
__e.Return(V3258)
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

tmp7786 := Call(__e, ns2_1set, symshen_4process_1application, tmp7742)


_ = tmp7786

tmp7787 := MakeNative(func(__e *ControlFlow) {
V3262 := __e.Get(1)
_ = V3262
tmp7813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3262)
}
__typedArg0 := V3262
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7793 Obj

if True == tmp7813 {
tmp7811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3262)
}
__typedArg0 := V3262
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7811)
}
__typedArg0 := tmp7811
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7795 Obj

if True == tmp7812 {
tmp7808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3262)
}
__typedArg0 := V3262
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7808)
}
__typedArg0 := tmp7808
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symforeign, tmp7809)
}
__typedArg0 := symforeign
__typedArg1 := tmp7809
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7797 Obj

if True == tmp7810 {
tmp7805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3262)
}
__typedArg0 := V3262
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7805)
}
__typedArg0 := tmp7805
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7806)
}
__typedArg0 := tmp7806
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7799 Obj

if True == tmp7807 {
tmp7801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3262)
}
__typedArg0 := V3262
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7801)
}
__typedArg0 := tmp7801
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7802)
}
__typedArg0 := tmp7802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7803)
}
__typedArg0 := Nil
__typedArg1 := tmp7803
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
tmp7788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3262)
}
__typedArg0 := V3262
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7788)
}
__typedArg0 := tmp7788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7789)
}
__typedArg0 := tmp7789
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3262)
}
__typedArg0 := V3262
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7790, tmp7791)
}
__typedArg0 := tmp7790
__typedArg1 := tmp7791
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4unpack_1foreign)
return
}


}, 1)

tmp7814 := Call(__e, ns2_1set, symshen_4unpack_1foreign, tmp7787)


_ = tmp7814

tmp7815 := MakeNative(func(__e *ControlFlow) {
V3265 := __e.Get(1)
_ = V3265
tmp7837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3265)
}
__typedArg0 := V3265
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7817 Obj

if True == tmp7837 {
tmp7835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3265)
}
__typedArg0 := V3265
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7835)
}
__typedArg0 := tmp7835
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7819 Obj

if True == tmp7836 {
tmp7832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3265)
}
__typedArg0 := V3265
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7832)
}
__typedArg0 := tmp7832
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symforeign, tmp7833)
}
__typedArg0 := symforeign
__typedArg1 := tmp7833
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7821 Obj

if True == tmp7834 {
tmp7829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3265)
}
__typedArg0 := V3265
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7829)
}
__typedArg0 := tmp7829
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7830)
}
__typedArg0 := tmp7830
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7823 Obj

if True == tmp7831 {
tmp7825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3265)
}
__typedArg0 := V3265
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7826)
}
__typedArg0 := tmp7826
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7827)
}
__typedArg0 := Nil
__typedArg1 := tmp7827
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp7838 := Call(__e, ns2_1set, symshen_4foreign_2, tmp7815)


_ = tmp7838

tmp7839 := MakeNative(func(__e *ControlFlow) {
V3268 := __e.Get(1)
_ = V3268
tmp7845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7841 Obj

if True == tmp7845 {
tmp7843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7843)
}
__typedArg0 := Nil
__typedArg1 := tmp7843
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp7846 := Call(__e, ns2_1set, symshen_4zero_1place_2, tmp7839)


_ = tmp7846

tmp7847 := MakeNative(func(__e *ControlFlow) {
V3269 := __e.Get(1)
_ = V3269
tmp7852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V3269)
}
__typedArg0 := V3269
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp7852 {
tmp7849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3269)
}
__typedArg0 := V3269
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp7850 := Call(__e, PrimFunc(symshen_4internal_1to_1shen_2), tmp7849)


if True == tmp7850 {
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

tmp7853 := Call(__e, ns2_1set, symshen_4shen_1call_2, tmp7847)


_ = tmp7853

tmp7854 := MakeNative(func(__e *ControlFlow) {
V3274 := __e.Get(1)
_ = V3274
tmp7884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3274)
}
__typedArg0 := V3274
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7871 Obj

if True == tmp7884 {
tmp7882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3274)
}
__typedArg0 := V3274
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symprotect, tmp7882)
}
__typedArg0 := symprotect
__typedArg1 := tmp7882
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7873 Obj

if True == tmp7883 {
tmp7880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3274)
}
__typedArg0 := V3274
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7880)
}
__typedArg0 := tmp7880
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7875 Obj

if True == tmp7881 {
tmp7877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3274)
}
__typedArg0 := V3274
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7877)
}
__typedArg0 := tmp7877
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7878)
}
__typedArg0 := Nil
__typedArg1 := tmp7878
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
__e.Return(False)
return
} else {
tmp7869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3274)
}
__typedArg0 := V3274
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7856 Obj

if True == tmp7869 {
tmp7867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3274)
}
__typedArg0 := V3274
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symforeign, tmp7867)
}
__typedArg0 := symforeign
__typedArg1 := tmp7867
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7858 Obj

if True == tmp7868 {
tmp7865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3274)
}
__typedArg0 := V3274
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7865)
}
__typedArg0 := tmp7865
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7860 Obj

if True == tmp7866 {
tmp7862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3274)
}
__typedArg0 := V3274
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7862)
}
__typedArg0 := tmp7862
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
__e.Return(False)
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3274)
}
__typedArg0 := V3274
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})())
return
}


}


}, 1)

tmp7885 := Call(__e, ns2_1set, symshen_4application_2, tmp7854)


_ = tmp7885

tmp7886 := MakeNative(func(__e *ControlFlow) {
V3279 := __e.Get(1)
_ = V3279
V3280 := __e.Get(2)
_ = V3280
tmp7894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(-1), V3280)
}
__typedArg0 := MakeInteger(-1)
__typedArg1 := V3280
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7894 {
tmp7892 := Call(__e, PrimFunc(symshen_4lowercase_1symbol_2), V3279)


if True == tmp7892 {
tmp7888 := Call(__e, PrimFunc(symexternal), symshen)


tmp7889 := Call(__e, PrimFunc(symelement_2), V3279, tmp7888)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7889)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7889
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

tmp7895 := Call(__e, ns2_1set, symshen_4undefined_1f_2, tmp7886)


_ = tmp7895

tmp7896 := MakeNative(func(__e *ControlFlow) {
V3281 := __e.Get(1)
_ = V3281
tmp7901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V3281)
}
__typedArg0 := V3281
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp7901 {
tmp7898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V3281)
}
__typedArg0 := V3281
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp7898)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp7898
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

tmp7902 := Call(__e, ns2_1set, symshen_4lowercase_1symbol_2, tmp7896)


_ = tmp7902

tmp7903 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V3282 := __self1
_ = V3282
tmp7933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3282)
}
__typedArg0 := V3282
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7924 Obj

if True == tmp7933 {
tmp7931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3282)
}
__typedArg0 := V3282
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7931)
}
__typedArg0 := tmp7931
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7926 Obj

if True == tmp7932 {
tmp7928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3282)
}
__typedArg0 := V3282
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7928)
}
__typedArg0 := tmp7928
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7929)
}
__typedArg0 := Nil
__typedArg1 := tmp7929
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
__e.Return(V3282)
return
} else {
tmp7922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3282)
}
__typedArg0 := V3282
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7913 Obj

if True == tmp7922 {
tmp7920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3282)
}
__typedArg0 := V3282
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7920)
}
__typedArg0 := tmp7920
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7915 Obj

if True == tmp7921 {
tmp7917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3282)
}
__typedArg0 := V3282
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7917)
}
__typedArg0 := tmp7917
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7918)
}
__typedArg0 := tmp7918
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

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
tmp7904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3282)
}
__typedArg0 := V3282
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3282)
}
__typedArg0 := V3282
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7905)
}
__typedArg0 := tmp7905
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7906, Nil)
}
__typedArg0 := tmp7906
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7904, tmp7907)
}
__typedArg0 := tmp7904
__typedArg1 := tmp7907
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp7909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3282)
}
__typedArg0 := V3282
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7909)
}
__typedArg0 := tmp7909
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp7908, tmp7910)
}
__typedArg0 := tmp7908
__typedArg1 := tmp7910
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4simple_1curry) == __self {
__self1 = tmp7911
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp7911)
return


} else {
__e.Return(V3282)
return
}


}


}, 1)

tmp7934 := Call(__e, ns2_1set, symshen_4simple_1curry, tmp7903)


_ = tmp7934

tmp7935 := MakeNative(func(__e *ControlFlow) {
V3283 := __e.Get(1)
_ = V3283
__e.TailApply(PrimFunc(symfn), V3283)
return
}, 1)

tmp7936 := Call(__e, ns2_1set, symfunction, tmp7935)


_ = tmp7936

tmp7937 := MakeNative(func(__e *ControlFlow) {
V3284 := __e.Get(1)
_ = V3284
tmp7944 := Call(__e, PrimFunc(symarity), V3284)


tmp7945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp7944, MakeInteger(0))
}
__typedArg0 := tmp7944
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7945 {
__e.TailApply(V3284)
return
} else {
tmp7938 := MakeNative(func(__e *ControlFlow) {
tmp7939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V3284, symshen_4lambda_1form, tmp7939)
return


}, 0)

tmp7940 := MakeNative(func(__e *ControlFlow) {
Z3285 := __e.Get(1)
_ = Z3285
tmp7941 := Call(__e, PrimFunc(symshen_4app), V3284, MakeString(" is undefined\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("fn: "))
__typedS1, __typedOK1 := TypedString(tmp7941)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("fn: ")
__typedArg1 := tmp7941
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("fn: "))
__typedS1, __typedOK1 := TypedString(tmp7941)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("fn: ")
__typedArg1 := tmp7941
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

__e.TailApply(try_1catch, tmp7938, tmp7940)
return


}


}, 1)

tmp7946 := Call(__e, ns2_1set, symfn, tmp7937)


_ = tmp7946

tmp7947 := MakeNative(func(__e *ControlFlow) {
V3288 := __e.Get(1)
_ = V3288
tmp7977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3288)
}
__typedArg0 := V3288
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7964 Obj

if True == tmp7977 {
tmp7975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3288)
}
__typedArg0 := V3288
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfn, tmp7975)
}
__typedArg0 := symfn
__typedArg1 := tmp7975
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7966 Obj

if True == tmp7976 {
tmp7973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3288)
}
__typedArg0 := V3288
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7973)
}
__typedArg0 := tmp7973
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7968 Obj

if True == tmp7974 {
tmp7970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3288)
}
__typedArg0 := V3288
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7970)
}
__typedArg0 := tmp7970
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7971)
}
__typedArg0 := Nil
__typedArg1 := tmp7971
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
__e.Return(True)
return
} else {
tmp7962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3288)
}
__typedArg0 := V3288
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7949 Obj

if True == tmp7962 {
tmp7960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3288)
}
__typedArg0 := V3288
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfunction, tmp7960)
}
__typedArg0 := symfunction
__typedArg1 := tmp7960
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7951 Obj

if True == tmp7961 {
tmp7958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3288)
}
__typedArg0 := V3288
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp7958)
}
__typedArg0 := tmp7958
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7953 Obj

if True == tmp7959 {
tmp7955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3288)
}
__typedArg0 := V3288
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7955)
}
__typedArg0 := tmp7955
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7956)
}
__typedArg0 := Nil
__typedArg1 := tmp7956
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp7978 := Call(__e, ns2_1set, symshen_4fn_1call_2, tmp7947)


_ = tmp7978

tmp7979 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V3289 := __self1
_ = V3289
tmp8019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3289)
}
__typedArg0 := V3289
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8006 Obj

if True == tmp8019 {
tmp8017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3289)
}
__typedArg0 := V3289
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfunction, tmp8017)
}
__typedArg0 := symfunction
__typedArg1 := tmp8017
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8008 Obj

if True == tmp8018 {
tmp8015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3289)
}
__typedArg0 := V3289
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8015)
}
__typedArg0 := tmp8015
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres8010 Obj

if True == tmp8016 {
tmp8012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3289)
}
__typedArg0 := V3289
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp8012)
}
__typedArg0 := tmp8012
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp8013)
}
__typedArg0 := Nil
__typedArg1 := tmp8013
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
tmp7980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3289)
}
__typedArg0 := V3289
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfn, tmp7980)
}
__typedArg0 := symfn
__typedArg1 := tmp7980
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4fn_1call) == __self {
__self1 = tmp7981
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4fn_1call), tmp7981)
return


} else {
tmp8004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3289)
}
__typedArg0 := V3289
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7991 Obj

if True == tmp8004 {
tmp8002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3289)
}
__typedArg0 := V3289
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp8003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfn, tmp8002)
}
__typedArg0 := symfn
__typedArg1 := tmp8002
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres7993 Obj

if True == tmp8003 {
tmp8000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3289)
}
__typedArg0 := V3289
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp8001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp8000)
}
__typedArg0 := tmp8000
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres7995 Obj

if True == tmp8001 {
tmp7997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3289)
}
__typedArg0 := V3289
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp7997)
}
__typedArg0 := tmp7997
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp7998)
}
__typedArg0 := Nil
__typedArg1 := tmp7998
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

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
tmp7982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3289)
}
__typedArg0 := V3289
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp7983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp7982)
}
__typedArg0 := tmp7982
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp7984 := Call(__e, PrimFunc(symarity), tmp7983)


let__5428 := tmp7984
_ = let__5428

tmp7989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__5428, MakeInteger(-1))
}
__typedArg0 := let__5428
__typedArg1 := MakeInteger(-1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7989 {
__e.Return(V3289)
return
} else {
tmp7987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__5428, MakeInteger(0))
}
__typedArg0 := let__5428
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp7987 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3289)
}
__typedArg0 := V3289
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
} else {
tmp7985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3289)
}
__typedArg0 := V3289
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4lambda_1function), tmp7985, let__5428)
return


}


}


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4fn_1call)
return
}


}


}, 1)

tmp8020 := Call(__e, ns2_1set, symshen_4fn_1call, tmp7979)


_ = tmp8020

tmp8021 := MakeNative(func(__e *ControlFlow) {
V3291 := __e.Get(1)
_ = V3291
V3292 := __e.Get(2)
_ = V3292
V3293 := __e.Get(3)
_ = V3293
let__5429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(V3292)
__typedN1, __typedOK1 := TypedFloat64(V3293)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := V3292
__typedArg1 := V3293
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})()
_ = let__5429

var ifres8028 Obj

if True == let__5429 {
tmp8036 := Call(__e, PrimFunc(symshen_4loading_2))


var ifres8030 Obj

if True == tmp8036 {
tmp8032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, Nil)
}
__typedArg0 := sym_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp8032)
}
__typedArg0 := sym_7
__typedArg1 := tmp8032
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp8034 := Call(__e, PrimFunc(symelement_2), V3291, tmp8033)


tmp8035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp8034)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp8034
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

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
tmp8024 := Call(__e, PrimFunc(symshen_4app), V3291, MakeString("\n"), symshen_4a)


tmp8025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("partial application of "))
__typedS1, __typedOK1 := TypedString(tmp8024)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("partial application of ")
__typedArg1 := tmp8024
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp8026 := Call(__e, PrimFunc(symstoutput))


tmp8027 := Call(__e, PrimFunc(sympr), tmp8025, tmp8026)


ifres8023 = tmp8027


} else {
ifres8023 = symshen_4skip


}

let__5430 := ifres8023
_ = let__5430

__e.Return(let__5429)
return


}, 3)

tmp8037 := Call(__e, ns2_1set, symshen_4partial_1application_d_2, tmp8021)


_ = tmp8037

tmp8038 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dloading_2_d)
}
__typedArg0 := symshen_4_dloading_2_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp8039 := Call(__e, ns2_1set, symshen_4loading_2, tmp8038)


_ = tmp8039

tmp8040 := MakeNative(func(__e *ControlFlow) {
V3300 := __e.Get(1)
_ = V3300
V3301 := __e.Get(2)
_ = V3301
V3302 := __e.Get(3)
_ = V3302
tmp8056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(-1), V3301)
}
__typedArg0 := MakeInteger(-1)
__typedArg1 := V3301
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp8056 {
__e.Return(False)
return
} else {
let__5431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5) {
__typedN0, __typedOK0 := TypedFloat64(V3301)
__typedN1, __typedOK1 := TypedFloat64(V3302)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_5) {
return TypedMaterializeBoolean((__typedN0 < __typedN1))
}}
__typedArg0 := V3301
__typedArg1 := V3302
return Call(__e, PrimFunc(sym_5), __typedArg0, __typedArg1)
})()
_ = let__5431

var ifres8052 Obj

if True == let__5431 {
tmp8054 := Call(__e, PrimFunc(symshen_4loading_2))


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
tmp8044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V3302, MakeInteger(1))
}
__typedArg0 := V3302
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres8043 Obj

if True == tmp8044 {
ifres8043 = MakeString("")


} else {
ifres8043 = MakeString("s")


}

tmp8045 := Call(__e, PrimFunc(symshen_4app), ifres8043, MakeString("\n"), symshen_4a)


tmp8047 := Call(__e, PrimFunc(symshen_4app), V3302, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" argument"))
__typedS1, __typedOK1 := TypedString(tmp8045)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" argument")
__typedArg1 := tmp8045
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp8049 := Call(__e, PrimFunc(symshen_4app), V3300, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" might not like "))
__typedS1, __typedOK1 := TypedString(tmp8047)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" might not like ")
__typedArg1 := tmp8047
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp8050 := Call(__e, PrimFunc(symstoutput))


tmp8051 := Call(__e, PrimFunc(sympr), tmp8049, tmp8050)


ifres8042 = tmp8051


} else {
ifres8042 = symshen_4skip


}

let__5432 := ifres8042
_ = let__5432

__e.Return(let__5431)
return


}


}, 3)

__e.TailApply(ns2_1set, symshen_4overapplication_2, tmp8040)
return




}, 0)

