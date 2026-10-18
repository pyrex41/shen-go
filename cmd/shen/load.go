package main

import . "github.com/pyrex41/shen-go/kl"

var LoadMain = MakeNative(func(__e *ControlFlow) {
tmp9468 := MakeNative(func(__e *ControlFlow) {
V1262 := __e.Get(1)
_ = V1262
tmp9469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtc_d)
}
__typedArg0 := symshen_4_dtc_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__9450 := tmp9469
_ = let__9450

tmp9470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symget_1time) {
return PrimGetTime(symrun)
}
__typedArg0 := symrun
return Call(__e, PrimFunc(symget_1time), __typedArg0)
})()

let__9452 := tmp9470
_ = let__9452

tmp9471 := Call(__e, PrimFunc(symread_1file), V1262)


tmp9472 := Call(__e, PrimFunc(symshen_4load_1help), let__9450, tmp9471)


let__9453 := tmp9472
_ = let__9453

tmp9473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symget_1time) {
return PrimGetTime(symrun)
}
__typedArg0 := symrun
return Call(__e, PrimFunc(symget_1time), __typedArg0)
})()

let__9454 := tmp9473
_ = let__9454

let__9455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(let__9454)
__typedN1, __typedOK1 := TypedFloat64(let__9452)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := let__9454
__typedArg1 := let__9452
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})()
_ = let__9455

tmp9475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(let__9455)
}
__typedArg0 := let__9455
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp9477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\nrun time: "))
__typedS1, __typedOK1 := TypedString(tmp9475)
__typedS2, __typedOK2 := TypedString(MakeString(" secs\n"))
if __typedOK0 && __typedOK1 && __typedOK2 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + (__typedS1 + __typedS2)))
}}
__typedArg0 := MakeString("\nrun time: ")
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp9475)
__typedS1, __typedOK1 := TypedString(MakeString(" secs\n"))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp9475
__typedArg1 := MakeString(" secs\n")
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp9478 := Call(__e, PrimFunc(symstoutput))


tmp9479 := Call(__e, PrimFunc(sympr), tmp9477, tmp9478)


let__9456 := tmp9479
_ = let__9456

let__9451 := let__9453
_ = let__9451

var ifres9480 Obj

if True == let__9450 {
tmp9481 := Call(__e, PrimFunc(syminferences))


tmp9482 := Call(__e, PrimFunc(symshen_4app), tmp9481, MakeString(" inferences\n"), symshen_4a)


tmp9483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\ntypechecked in "))
__typedS1, __typedOK1 := TypedString(tmp9482)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\ntypechecked in ")
__typedArg1 := tmp9482
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp9484 := Call(__e, PrimFunc(symstoutput))


tmp9485 := Call(__e, PrimFunc(sympr), tmp9483, tmp9484)


ifres9480 = tmp9485


} else {
ifres9480 = symshen_4skip


}

let__9457 := ifres9480
_ = let__9457

__e.Return(symloaded)
return


}, 1)

tmp9486 := Call(__e, ns2_1set, symload, tmp9468)


_ = tmp9486

tmp9487 := MakeNative(func(__e *ControlFlow) {
V1273 := __e.Get(1)
_ = V1273
V1274 := __e.Get(2)
_ = V1274
tmp9489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, V1273)
}
__typedArg0 := False
__typedArg1 := V1273
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9489 {
__e.TailApply(PrimFunc(symshen_4eval_1and_1print), V1274)
return
} else {
__e.TailApply(PrimFunc(symshen_4check_1eval_1and_1print), V1274)
return
}


}, 2)

tmp9490 := Call(__e, ns2_1set, symshen_4load_1help, tmp9487)


_ = tmp9490

tmp9491 := MakeNative(func(__e *ControlFlow) {
V1275 := __e.Get(1)
_ = V1275
tmp9492 := MakeNative(func(__e *ControlFlow) {
Z1276 := __e.Get(1)
_ = Z1276
tmp9493 := Call(__e, PrimFunc(symshen_4shen_1_6kl), Z1276)


tmp9494 := Call(__e, PrimFunc(symeval_1kl), tmp9493)


tmp9495 := Call(__e, PrimFunc(symshen_4app), tmp9494, MakeString("\n"), symshen_4s)


tmp9496 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp9495, tmp9496)
return


}, 1)

__e.TailApply(PrimFunc(symshen_4for_1each), tmp9492, V1275)
return


}, 1)

tmp9497 := Call(__e, ns2_1set, symshen_4eval_1and_1print, tmp9491)


_ = tmp9497

tmp9498 := MakeNative(func(__e *ControlFlow) {
V1277 := __e.Get(1)
_ = V1277
tmp9499 := MakeNative(func(__e *ControlFlow) {
Z1279 := __e.Get(1)
_ = Z1279
__e.TailApply(PrimFunc(symshen_4typetable), Z1279)
return
}, 1)

tmp9500 := Call(__e, PrimFunc(symmapcan), tmp9499, V1277)


let__9458 := tmp9500
_ = let__9458

tmp9501 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4assumetypes), let__9458)
return
}, 0)

tmp9502 := MakeNative(func(__e *ControlFlow) {
Z1281 := __e.Get(1)
_ = Z1281
__e.TailApply(PrimFunc(symshen_4unwind_1types), Z1281, let__9458)
return
}, 1)

tmp9503 := Call(__e, try_1catch, tmp9501, tmp9502)


let__9459 := tmp9503
_ = let__9459

tmp9504 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4work_1through), V1277)
return
}, 0)

tmp9505 := MakeNative(func(__e *ControlFlow) {
Z1282 := __e.Get(1)
_ = Z1282
__e.TailApply(PrimFunc(symshen_4unwind_1types), Z1282, let__9458)
return
}, 1)

__e.TailApply(try_1catch, tmp9504, tmp9505)
return


}, 1)

tmp9506 := Call(__e, ns2_1set, symshen_4check_1eval_1and_1print, tmp9498)


_ = tmp9506

tmp9507 := MakeNative(func(__e *ControlFlow) {
V1287 := __e.Get(1)
_ = V1287
tmp9552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1287)
}
__typedArg0 := V1287
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9533 Obj

if True == tmp9552 {
tmp9550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1287)
}
__typedArg0 := V1287
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp9550)
}
__typedArg0 := symdefine
__typedArg1 := tmp9550
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9535 Obj

if True == tmp9551 {
tmp9548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1287)
}
__typedArg0 := V1287
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9548)
}
__typedArg0 := tmp9548
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9537 Obj

if True == tmp9549 {
tmp9545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1287)
}
__typedArg0 := V1287
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9545)
}
__typedArg0 := tmp9545
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9546)
}
__typedArg0 := tmp9546
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9539 Obj

if True == tmp9547 {
tmp9541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1287)
}
__typedArg0 := V1287
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9541)
}
__typedArg0 := tmp9541
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9542)
}
__typedArg0 := tmp9542
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, tmp9543)
}
__typedArg0 := sym_i
__typedArg1 := tmp9543
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9540 Obj

if True == tmp9544 {
ifres9540 = True


} else {
ifres9540 = False


}

ifres9539 = ifres9540


} else {
ifres9539 = False


}

var ifres9538 Obj

if True == ifres9539 {
ifres9538 = True


} else {
ifres9538 = False


}

ifres9537 = ifres9538


} else {
ifres9537 = False


}

var ifres9536 Obj

if True == ifres9537 {
ifres9536 = True


} else {
ifres9536 = False


}

ifres9535 = ifres9536


} else {
ifres9535 = False


}

var ifres9534 Obj

if True == ifres9535 {
ifres9534 = True


} else {
ifres9534 = False


}

ifres9533 = ifres9534


} else {
ifres9533 = False


}

if True == ifres9533 {
tmp9508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1287)
}
__typedArg0 := V1287
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9508)
}
__typedArg0 := tmp9508
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1287)
}
__typedArg0 := V1287
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9510)
}
__typedArg0 := tmp9510
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1287)
}
__typedArg0 := V1287
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9512)
}
__typedArg0 := tmp9512
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9513)
}
__typedArg0 := tmp9513
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9515 := Call(__e, PrimFunc(symshen_4type_1F), tmp9511, tmp9514)


tmp9516 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp9515)


tmp9517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9516, Nil)
}
__typedArg0 := tmp9516
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9509, tmp9517)
}
__typedArg0 := tmp9509
__typedArg1 := tmp9517
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1287)
}
__typedArg0 := V1287
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9523 Obj

if True == tmp9531 {
tmp9529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1287)
}
__typedArg0 := V1287
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp9529)
}
__typedArg0 := symdefine
__typedArg1 := tmp9529
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9525 Obj

if True == tmp9530 {
tmp9527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1287)
}
__typedArg0 := V1287
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9527)
}
__typedArg0 := tmp9527
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9526 Obj

if True == tmp9528 {
ifres9526 = True


} else {
ifres9526 = False


}

ifres9525 = ifres9526


} else {
ifres9525 = False


}

var ifres9524 Obj

if True == ifres9525 {
ifres9524 = True


} else {
ifres9524 = False


}

ifres9523 = ifres9524


} else {
ifres9523 = False


}

if True == ifres9523 {
tmp9518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1287)
}
__typedArg0 := V1287
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9518)
}
__typedArg0 := tmp9518
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9520 := Call(__e, PrimFunc(symshen_4app), tmp9519, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing { in "))
__typedS1, __typedOK1 := TypedString(tmp9520)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing { in ")
__typedArg1 := tmp9520
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing { in "))
__typedS1, __typedOK1 := TypedString(tmp9520)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing { in ")
__typedArg1 := tmp9520
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


} else {
__e.Return(Nil)
return
}


}


}, 1)

tmp9553 := Call(__e, ns2_1set, symshen_4typetable, tmp9507)


_ = tmp9553

tmp9554 := MakeNative(func(__e *ControlFlow) {
V1294 := __e.Get(1)
_ = V1294
V1295 := __e.Get(2)
_ = V1295
tmp9567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1295)
}
__typedArg0 := V1295
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9563 Obj

if True == tmp9567 {
tmp9565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1295)
}
__typedArg0 := V1295
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, tmp9565)
}
__typedArg0 := sym_j
__typedArg1 := tmp9565
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9564 Obj

if True == tmp9566 {
ifres9564 = True


} else {
ifres9564 = False


}

ifres9563 = ifres9564


} else {
ifres9563 = False


}

if True == ifres9563 {
__e.Return(Nil)
return
} else {
tmp9561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1295)
}
__typedArg0 := V1295
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9561 {
tmp9555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1295)
}
__typedArg0 := V1295
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1295)
}
__typedArg0 := V1295
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9557 := Call(__e, PrimFunc(symshen_4type_1F), V1294, tmp9556)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9555, tmp9557)
}
__typedArg0 := tmp9555
__typedArg1 := tmp9557
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9558 := Call(__e, PrimFunc(symshen_4app), V1294, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing } in "))
__typedS1, __typedOK1 := TypedString(tmp9558)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing } in ")
__typedArg1 := tmp9558
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing } in "))
__typedS1, __typedOK1 := TypedString(tmp9558)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing } in ")
__typedArg1 := tmp9558
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}


}, 2)

tmp9568 := Call(__e, ns2_1set, symshen_4type_1F, tmp9554)


_ = tmp9568

tmp9569 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V1298 := __self1
_ = V1298
tmp9583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1298)
}
__typedArg0 := Nil
__typedArg1 := V1298
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9583 {
__e.Return(Nil)
return
} else {
tmp9581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1298)
}
__typedArg0 := V1298
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9577 Obj

if True == tmp9581 {
tmp9579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1298)
}
__typedArg0 := V1298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9579)
}
__typedArg0 := tmp9579
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9578 Obj

if True == tmp9580 {
ifres9578 = True


} else {
ifres9578 = False


}

ifres9577 = ifres9578


} else {
ifres9577 = False


}

if True == ifres9577 {
tmp9570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1298)
}
__typedArg0 := V1298
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1298)
}
__typedArg0 := V1298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9571)
}
__typedArg0 := tmp9571
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9573 := Call(__e, PrimFunc(symdeclare), tmp9570, tmp9572)


_ = tmp9573

tmp9574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1298)
}
__typedArg0 := V1298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9574)
}
__typedArg0 := tmp9574
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4assumetypes) == __self {
__self1 = tmp9575
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4assumetypes), tmp9575)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.assumetype"))
}
__typedArg0 := MakeString("implementation error in shen.assumetype")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp9584 := Call(__e, ns2_1set, symshen_4assumetypes, tmp9569)


_ = tmp9584

tmp9585 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V1303 := __self1
_ = V1303
V1304 := __self2
_ = V1304
tmp9596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1304)
}
__typedArg0 := V1304
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9592 Obj

if True == tmp9596 {
tmp9594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1304)
}
__typedArg0 := V1304
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9594)
}
__typedArg0 := tmp9594
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9593 Obj

if True == tmp9595 {
ifres9593 = True


} else {
ifres9593 = False


}

ifres9592 = ifres9593


} else {
ifres9592 = False


}

if True == ifres9592 {
tmp9586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1304)
}
__typedArg0 := V1304
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9587 := Call(__e, PrimFunc(symdestroy), tmp9586)


_ = tmp9587

tmp9588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1304)
}
__typedArg0 := V1304
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9588)
}
__typedArg0 := tmp9588
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4unwind_1types) == __self {
__self1, __self2 = V1303, tmp9589
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4unwind_1types), V1303, tmp9589)
return


} else {
tmp9590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symerror_1to_1string) {
return PrimErrorToString(V1303)
}
__typedArg0 := V1303
return Call(__e, PrimFunc(symerror_1to_1string), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp9590)
}
__typedArg0 := tmp9590
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 2)

tmp9597 := Call(__e, ns2_1set, symshen_4unwind_1types, tmp9585)


_ = tmp9597

tmp9598 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V1307 := __self1
_ = V1307
tmp9644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1307)
}
__typedArg0 := Nil
__typedArg1 := V1307
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9644 {
__e.Return(Nil)
return
} else {
tmp9642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9627 Obj

if True == tmp9642 {
tmp9640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9640)
}
__typedArg0 := tmp9640
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9629 Obj

if True == tmp9641 {
tmp9637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9637)
}
__typedArg0 := tmp9637
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9638)
}
__typedArg0 := tmp9638
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9631 Obj

if True == tmp9639 {
tmp9633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9633)
}
__typedArg0 := tmp9633
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp9636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp9634, tmp9635)
}
__typedArg0 := tmp9634
__typedArg1 := tmp9635
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9632 Obj

if True == tmp9636 {
ifres9632 = True


} else {
ifres9632 = False


}

ifres9631 = ifres9632


} else {
ifres9631 = False


}

var ifres9630 Obj

if True == ifres9631 {
ifres9630 = True


} else {
ifres9630 = False


}

ifres9629 = ifres9630


} else {
ifres9629 = False


}

var ifres9628 Obj

if True == ifres9629 {
ifres9628 = True


} else {
ifres9628 = False


}

ifres9627 = ifres9628


} else {
ifres9627 = False


}

if True == ifres9627 {
tmp9599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9600)
}
__typedArg0 := tmp9600
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9601)
}
__typedArg0 := tmp9601
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9603 := Call(__e, PrimFunc(symshen_4typecheck), tmp9599, tmp9602)


let__9460 := tmp9603
_ = let__9460

tmp9617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__9460, False)
}
__typedArg0 := let__9460
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9617 {
__e.TailApply(PrimFunc(symshen_4type_1error))
return
} else {
tmp9604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9605 := Call(__e, PrimFunc(symshen_4shen_1_6kl), tmp9604)


tmp9606 := Call(__e, PrimFunc(symeval_1kl), tmp9605)


let__9461 := tmp9606
_ = let__9461

tmp9607 := Call(__e, PrimFunc(symshen_4pretty_1type), let__9460)


tmp9608 := Call(__e, PrimFunc(symshen_4app), tmp9607, MakeString("\n"), symshen_4r)


tmp9610 := Call(__e, PrimFunc(symshen_4app), let__9461, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" : "))
__typedS1, __typedOK1 := TypedString(tmp9608)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" : ")
__typedArg1 := tmp9608
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4s)


tmp9611 := Call(__e, PrimFunc(symstoutput))


tmp9612 := Call(__e, PrimFunc(sympr), tmp9610, tmp9611)


let__9462 := tmp9612
_ = let__9462

tmp9613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9613)
}
__typedArg0 := tmp9613
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9614)
}
__typedArg0 := tmp9614
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4work_1through) == __self {
__self1 = tmp9615
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4work_1through), tmp9615)
return


}


} else {
tmp9625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9625 {
tmp9618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp9620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp9620)
}
__typedArg0 := symA
__typedArg1 := tmp9620
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9619, tmp9621)
}
__typedArg0 := tmp9619
__typedArg1 := tmp9621
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9618, tmp9622)
}
__typedArg0 := tmp9618
__typedArg1 := tmp9622
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4work_1through) == __self {
__self1 = tmp9623
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4work_1through), tmp9623)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.work-through"))
}
__typedArg0 := MakeString("implementation error in shen.work-through")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 1)

tmp9645 := Call(__e, ns2_1set, symshen_4work_1through, tmp9598)


_ = tmp9645

tmp9646 := MakeNative(func(__e *ControlFlow) {
V1312 := __e.Get(1)
_ = V1312
tmp9788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9662 Obj

if True == tmp9788 {
tmp9786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9786)
}
__typedArg0 := tmp9786
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9664 Obj

if True == tmp9787 {
tmp9783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9783)
}
__typedArg0 := tmp9783
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp9784)
}
__typedArg0 := symlist
__typedArg1 := tmp9784
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9666 Obj

if True == tmp9785 {
tmp9780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9780)
}
__typedArg0 := tmp9780
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9781)
}
__typedArg0 := tmp9781
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9668 Obj

if True == tmp9782 {
tmp9776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9776)
}
__typedArg0 := tmp9776
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9777)
}
__typedArg0 := tmp9777
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9778)
}
__typedArg0 := Nil
__typedArg1 := tmp9778
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9670 Obj

if True == tmp9779 {
tmp9774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9774)
}
__typedArg0 := tmp9774
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9672 Obj

if True == tmp9775 {
tmp9771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9771)
}
__typedArg0 := tmp9771
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1_1_6, tmp9772)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp9772
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9674 Obj

if True == tmp9773 {
tmp9768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9768)
}
__typedArg0 := tmp9768
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9769)
}
__typedArg0 := tmp9769
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9676 Obj

if True == tmp9770 {
tmp9764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9764)
}
__typedArg0 := tmp9764
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9765)
}
__typedArg0 := tmp9765
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9766)
}
__typedArg0 := tmp9766
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9678 Obj

if True == tmp9767 {
tmp9759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9759)
}
__typedArg0 := tmp9759
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9760)
}
__typedArg0 := tmp9760
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9761)
}
__typedArg0 := tmp9761
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symstr, tmp9762)
}
__typedArg0 := symstr
__typedArg1 := tmp9762
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9680 Obj

if True == tmp9763 {
tmp9754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9754)
}
__typedArg0 := tmp9754
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9755)
}
__typedArg0 := tmp9755
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9756)
}
__typedArg0 := tmp9756
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9757)
}
__typedArg0 := tmp9757
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9682 Obj

if True == tmp9758 {
tmp9748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9748)
}
__typedArg0 := tmp9748
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9749)
}
__typedArg0 := tmp9749
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9750)
}
__typedArg0 := tmp9750
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9751)
}
__typedArg0 := tmp9751
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9752)
}
__typedArg0 := tmp9752
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9684 Obj

if True == tmp9753 {
tmp9741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9741)
}
__typedArg0 := tmp9741
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9742)
}
__typedArg0 := tmp9742
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9743)
}
__typedArg0 := tmp9743
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9744)
}
__typedArg0 := tmp9744
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9745)
}
__typedArg0 := tmp9745
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp9746)
}
__typedArg0 := symlist
__typedArg1 := tmp9746
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9686 Obj

if True == tmp9747 {
tmp9734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9734)
}
__typedArg0 := tmp9734
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9735)
}
__typedArg0 := tmp9735
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9736)
}
__typedArg0 := tmp9736
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9737)
}
__typedArg0 := tmp9737
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9738)
}
__typedArg0 := tmp9738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9739)
}
__typedArg0 := tmp9739
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9688 Obj

if True == tmp9740 {
tmp9726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9726)
}
__typedArg0 := tmp9726
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9727)
}
__typedArg0 := tmp9727
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9728)
}
__typedArg0 := tmp9728
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9729)
}
__typedArg0 := tmp9729
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9730)
}
__typedArg0 := tmp9730
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9731)
}
__typedArg0 := tmp9731
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9732)
}
__typedArg0 := Nil
__typedArg1 := tmp9732
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9690 Obj

if True == tmp9733 {
tmp9720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9720)
}
__typedArg0 := tmp9720
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9721)
}
__typedArg0 := tmp9721
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9722)
}
__typedArg0 := tmp9722
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9723)
}
__typedArg0 := tmp9723
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9724)
}
__typedArg0 := tmp9724
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9692 Obj

if True == tmp9725 {
tmp9713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9713)
}
__typedArg0 := tmp9713
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9714)
}
__typedArg0 := tmp9714
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9715)
}
__typedArg0 := tmp9715
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9716)
}
__typedArg0 := tmp9716
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9717)
}
__typedArg0 := tmp9717
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9718)
}
__typedArg0 := Nil
__typedArg1 := tmp9718
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9694 Obj

if True == tmp9719 {
tmp9709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9709)
}
__typedArg0 := tmp9709
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9710)
}
__typedArg0 := tmp9710
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9711)
}
__typedArg0 := Nil
__typedArg1 := tmp9711
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9696 Obj

if True == tmp9712 {
tmp9698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9698)
}
__typedArg0 := tmp9698
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9699)
}
__typedArg0 := tmp9699
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9701)
}
__typedArg0 := tmp9701
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9702)
}
__typedArg0 := tmp9702
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9703)
}
__typedArg0 := tmp9703
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9704)
}
__typedArg0 := tmp9704
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9705)
}
__typedArg0 := tmp9705
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9706)
}
__typedArg0 := tmp9706
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp9700, tmp9707)
}
__typedArg0 := tmp9700
__typedArg1 := tmp9707
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9697 Obj

if True == tmp9708 {
ifres9697 = True


} else {
ifres9697 = False


}

ifres9696 = ifres9697


} else {
ifres9696 = False


}

var ifres9695 Obj

if True == ifres9696 {
ifres9695 = True


} else {
ifres9695 = False


}

ifres9694 = ifres9695


} else {
ifres9694 = False


}

var ifres9693 Obj

if True == ifres9694 {
ifres9693 = True


} else {
ifres9693 = False


}

ifres9692 = ifres9693


} else {
ifres9692 = False


}

var ifres9691 Obj

if True == ifres9692 {
ifres9691 = True


} else {
ifres9691 = False


}

ifres9690 = ifres9691


} else {
ifres9690 = False


}

var ifres9689 Obj

if True == ifres9690 {
ifres9689 = True


} else {
ifres9689 = False


}

ifres9688 = ifres9689


} else {
ifres9688 = False


}

var ifres9687 Obj

if True == ifres9688 {
ifres9687 = True


} else {
ifres9687 = False


}

ifres9686 = ifres9687


} else {
ifres9686 = False


}

var ifres9685 Obj

if True == ifres9686 {
ifres9685 = True


} else {
ifres9685 = False


}

ifres9684 = ifres9685


} else {
ifres9684 = False


}

var ifres9683 Obj

if True == ifres9684 {
ifres9683 = True


} else {
ifres9683 = False


}

ifres9682 = ifres9683


} else {
ifres9682 = False


}

var ifres9681 Obj

if True == ifres9682 {
ifres9681 = True


} else {
ifres9681 = False


}

ifres9680 = ifres9681


} else {
ifres9680 = False


}

var ifres9679 Obj

if True == ifres9680 {
ifres9679 = True


} else {
ifres9679 = False


}

ifres9678 = ifres9679


} else {
ifres9678 = False


}

var ifres9677 Obj

if True == ifres9678 {
ifres9677 = True


} else {
ifres9677 = False


}

ifres9676 = ifres9677


} else {
ifres9676 = False


}

var ifres9675 Obj

if True == ifres9676 {
ifres9675 = True


} else {
ifres9675 = False


}

ifres9674 = ifres9675


} else {
ifres9674 = False


}

var ifres9673 Obj

if True == ifres9674 {
ifres9673 = True


} else {
ifres9673 = False


}

ifres9672 = ifres9673


} else {
ifres9672 = False


}

var ifres9671 Obj

if True == ifres9672 {
ifres9671 = True


} else {
ifres9671 = False


}

ifres9670 = ifres9671


} else {
ifres9670 = False


}

var ifres9669 Obj

if True == ifres9670 {
ifres9669 = True


} else {
ifres9669 = False


}

ifres9668 = ifres9669


} else {
ifres9668 = False


}

var ifres9667 Obj

if True == ifres9668 {
ifres9667 = True


} else {
ifres9667 = False


}

ifres9666 = ifres9667


} else {
ifres9666 = False


}

var ifres9665 Obj

if True == ifres9666 {
ifres9665 = True


} else {
ifres9665 = False


}

ifres9664 = ifres9665


} else {
ifres9664 = False


}

var ifres9663 Obj

if True == ifres9664 {
ifres9663 = True


} else {
ifres9663 = False


}

ifres9662 = ifres9663


} else {
ifres9662 = False


}

if True == ifres9662 {
tmp9647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9647)
}
__typedArg0 := tmp9647
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9648)
}
__typedArg0 := tmp9648
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9649)
}
__typedArg0 := tmp9649
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9650)
}
__typedArg0 := tmp9650
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9652)
}
__typedArg0 := tmp9652
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9653)
}
__typedArg0 := tmp9653
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9654)
}
__typedArg0 := tmp9654
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9655)
}
__typedArg0 := tmp9655
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a_6, tmp9656)
}
__typedArg0 := sym_a_a_6
__typedArg1 := tmp9656
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9651, tmp9657)
}
__typedArg0 := tmp9651
__typedArg1 := tmp9657
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9660 {
tmp9658 := MakeNative(func(__e *ControlFlow) {
Z1313 := __e.Get(1)
_ = Z1313
__e.TailApply(PrimFunc(symshen_4pretty_1type), Z1313)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp9658, V1312)
return


} else {
__e.Return(V1312)
return
}


}


}, 1)

tmp9789 := Call(__e, ns2_1set, symshen_4pretty_1type, tmp9646)


_ = tmp9789

tmp9790 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("type error\n"))
}
__typedArg0 := MakeString("type error\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}, 0)

tmp9791 := Call(__e, ns2_1set, symshen_4type_1error, tmp9790)


_ = tmp9791

tmp9792 := MakeNative(func(__e *ControlFlow) {
V1314 := __e.Get(1)
_ = V1314
tmp9793 := Call(__e, PrimFunc(symshen_4klfile), V1314)


let__9463 := tmp9793
_ = let__9463

tmp9794 := Call(__e, PrimFunc(symread_1file), V1314)


let__9464 := tmp9794
_ = let__9464

tmp9795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symopen) {
return PrimOpenStream(let__9463, symout)
}
__typedArg0 := let__9463
__typedArg1 := symout
return Call(__e, PrimFunc(symopen), __typedArg0, __typedArg1)
})()

let__9465 := tmp9795
_ = let__9465

tmp9796 := MakeNative(func(__e *ControlFlow) {
Z1319 := __e.Get(1)
_ = Z1319
tmp9797 := Call(__e, PrimFunc(symshen_4shen_1_6kl_1h), Z1319)


__e.TailApply(PrimFunc(symshen_4partial), tmp9797)
return


}, 1)

tmp9798 := Call(__e, PrimFunc(symmap), tmp9796, let__9464)


let__9466 := tmp9798
_ = let__9466

tmp9799 := Call(__e, PrimFunc(symshen_4write_1kl), let__9466, let__9465)


let__9467 := tmp9799
_ = let__9467

__e.Return(let__9463)
return


}, 1)

tmp9800 := Call(__e, ns2_1set, symbootstrap, tmp9792)


_ = tmp9800

tmp9801 := MakeNative(func(__e *ControlFlow) {
V1321 := __e.Get(1)
_ = V1321
tmp9824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1321)
}
__typedArg0 := V1321
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9811 Obj

if True == tmp9824 {
tmp9822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1321)
}
__typedArg0 := V1321
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4f_1error, tmp9822)
}
__typedArg0 := symshen_4f_1error
__typedArg1 := tmp9822
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9813 Obj

if True == tmp9823 {
tmp9820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1321)
}
__typedArg0 := V1321
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9820)
}
__typedArg0 := tmp9820
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9815 Obj

if True == tmp9821 {
tmp9817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1321)
}
__typedArg0 := V1321
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9817)
}
__typedArg0 := tmp9817
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9818)
}
__typedArg0 := Nil
__typedArg1 := tmp9818
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9816 Obj

if True == tmp9819 {
ifres9816 = True


} else {
ifres9816 = False


}

ifres9815 = ifres9816


} else {
ifres9815 = False


}

var ifres9814 Obj

if True == ifres9815 {
ifres9814 = True


} else {
ifres9814 = False


}

ifres9813 = ifres9814


} else {
ifres9813 = False


}

var ifres9812 Obj

if True == ifres9813 {
ifres9812 = True


} else {
ifres9812 = False


}

ifres9811 = ifres9812


} else {
ifres9811 = False


}

if True == ifres9811 {
tmp9802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1321)
}
__typedArg0 := V1321
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9802)
}
__typedArg0 := tmp9802
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(tmp9803)
}
__typedArg0 := tmp9803
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp9806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("partial function "))
__typedS1, __typedOK1 := TypedString(tmp9804)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("partial function ")
__typedArg1 := tmp9804
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), Nil)
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("partial function "))
__typedS1, __typedOK1 := TypedString(tmp9804)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("partial function ")
__typedArg1 := tmp9804
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp9806)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp9806
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1321)
}
__typedArg0 := V1321
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9809 {
tmp9807 := MakeNative(func(__e *ControlFlow) {
Z1322 := __e.Get(1)
_ = Z1322
__e.TailApply(PrimFunc(symshen_4partial), Z1322)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp9807, V1321)
return


} else {
__e.Return(V1321)
return
}


}


}, 1)

tmp9825 := Call(__e, ns2_1set, symshen_4partial, tmp9801)


_ = tmp9825

tmp9826 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V1325 := __self1
_ = V1325
V1326 := __self2
_ = V1326
tmp9840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1325)
}
__typedArg0 := Nil
__typedArg1 := V1325
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9840 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symclose) {
return PrimCloseStream(V1326)
}
__typedArg0 := V1326
return Call(__e, PrimFunc(symclose), __typedArg0)
})())
return
} else {
tmp9838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1325)
}
__typedArg0 := V1325
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9834 Obj

if True == tmp9838 {
tmp9836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1325)
}
__typedArg0 := V1325
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9836)
}
__typedArg0 := tmp9836
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9835 Obj

if True == tmp9837 {
ifres9835 = True


} else {
ifres9835 = False


}

ifres9834 = ifres9835


} else {
ifres9834 = False


}

if True == ifres9834 {
tmp9827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1325)
}
__typedArg0 := V1325
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1325)
}
__typedArg0 := V1325
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9829 := Call(__e, PrimFunc(symshen_4write_1kl_1h), tmp9828, V1326)


_ = tmp9829

if PrimFunc(symshen_4write_1kl) == __self {
__self1, __self2 = tmp9827, V1326
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4write_1kl), tmp9827, V1326)
return


} else {
tmp9832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1325)
}
__typedArg0 := V1325
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9832 {
tmp9830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1325)
}
__typedArg0 := V1325
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4write_1kl) == __self {
__self1, __self2 = tmp9830, V1326
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4write_1kl), tmp9830, V1326)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4write_1kl)
return
}


}


}


}, 2)

tmp9841 := Call(__e, ns2_1set, symshen_4write_1kl, tmp9826)


_ = tmp9841

tmp9842 := MakeNative(func(__e *ControlFlow) {
V1329 := __e.Get(1)
_ = V1329
V1330 := __e.Get(2)
_ = V1330
tmp9882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1329)
}
__typedArg0 := V1329
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9845 Obj

if True == tmp9882 {
tmp9880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1329)
}
__typedArg0 := V1329
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp9880)
}
__typedArg0 := symdefun
__typedArg1 := tmp9880
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9847 Obj

if True == tmp9881 {
tmp9878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1329)
}
__typedArg0 := V1329
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9878)
}
__typedArg0 := tmp9878
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9849 Obj

if True == tmp9879 {
tmp9875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1329)
}
__typedArg0 := V1329
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9875)
}
__typedArg0 := tmp9875
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail, tmp9876)
}
__typedArg0 := symfail
__typedArg1 := tmp9876
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9851 Obj

if True == tmp9877 {
tmp9872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1329)
}
__typedArg0 := V1329
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9872)
}
__typedArg0 := tmp9872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9873)
}
__typedArg0 := tmp9873
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9853 Obj

if True == tmp9874 {
tmp9868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1329)
}
__typedArg0 := V1329
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9868)
}
__typedArg0 := tmp9868
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9869)
}
__typedArg0 := tmp9869
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9870)
}
__typedArg0 := Nil
__typedArg1 := tmp9870
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9855 Obj

if True == tmp9871 {
tmp9864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1329)
}
__typedArg0 := V1329
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9864)
}
__typedArg0 := tmp9864
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9865)
}
__typedArg0 := tmp9865
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9866)
}
__typedArg0 := tmp9866
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9857 Obj

if True == tmp9867 {
tmp9859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1329)
}
__typedArg0 := V1329
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9859)
}
__typedArg0 := tmp9859
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9860)
}
__typedArg0 := tmp9860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9861)
}
__typedArg0 := tmp9861
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9862)
}
__typedArg0 := Nil
__typedArg1 := tmp9862
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9858 Obj

if True == tmp9863 {
ifres9858 = True


} else {
ifres9858 = False


}

ifres9857 = ifres9858


} else {
ifres9857 = False


}

var ifres9856 Obj

if True == ifres9857 {
ifres9856 = True


} else {
ifres9856 = False


}

ifres9855 = ifres9856


} else {
ifres9855 = False


}

var ifres9854 Obj

if True == ifres9855 {
ifres9854 = True


} else {
ifres9854 = False


}

ifres9853 = ifres9854


} else {
ifres9853 = False


}

var ifres9852 Obj

if True == ifres9853 {
ifres9852 = True


} else {
ifres9852 = False


}

ifres9851 = ifres9852


} else {
ifres9851 = False


}

var ifres9850 Obj

if True == ifres9851 {
ifres9850 = True


} else {
ifres9850 = False


}

ifres9849 = ifres9850


} else {
ifres9849 = False


}

var ifres9848 Obj

if True == ifres9849 {
ifres9848 = True


} else {
ifres9848 = False


}

ifres9847 = ifres9848


} else {
ifres9847 = False


}

var ifres9846 Obj

if True == ifres9847 {
ifres9846 = True


} else {
ifres9846 = False


}

ifres9845 = ifres9846


} else {
ifres9845 = False


}

if True == ifres9845 {
__e.TailApply(PrimFunc(sympr), MakeString("(defun fail () shen.fail!)"), V1330)
return
} else {
tmp9843 := Call(__e, PrimFunc(symshen_4app), V1329, MakeString("\n\n"), symshen_4r)


__e.TailApply(PrimFunc(sympr), tmp9843, V1330)
return


}


}, 2)

tmp9883 := Call(__e, ns2_1set, symshen_4write_1kl_1h, tmp9842)


_ = tmp9883

tmp9884 := MakeNative(func(__e *ControlFlow) {
V1331 := __e.Get(1)
_ = V1331
tmp9893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V1331)
}
__typedArg0 := MakeString("")
__typedArg1 := V1331
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9893 {
__e.Return(MakeString(".kl"))
return
} else {
tmp9891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(".shen"), V1331)
}
__typedArg0 := MakeString(".shen")
__typedArg1 := V1331
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9891 {
__e.Return(MakeString(".kl"))
return
} else {
tmp9889 := Call(__e, PrimFunc(symshen_4_7string_2), V1331)


if True == tmp9889 {
tmp9885 := Call(__e, PrimFunc(symhdstr), V1331)


tmp9887 := Call(__e, PrimFunc(symshen_4klfile), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V1331)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V1331
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.TailApply(PrimFunc(sym_8s), tmp9885, tmp9887)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4klfile)
return
}


}


}


}, 1)

__e.TailApply(ns2_1set, symshen_4klfile, tmp9884)
return




}, 0)

