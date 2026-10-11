package main

import . "github.com/pyrex41/shen-go/kl"

var LoadMain = MakeNative(func(__e *ControlFlow) {
tmp9461 := MakeNative(func(__e *ControlFlow) {
V1261 := __e.Get(1)
_ = V1261
tmp9462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtc_d)
}
__typedArg0 := symshen_4_dtc_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__9443 := tmp9462
_ = let__9443

tmp9463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symget_1time) {
return PrimGetTime(symrun)
}
__typedArg0 := symrun
return Call(__e, PrimFunc(symget_1time), __typedArg0)
})()

let__9445 := tmp9463
_ = let__9445

tmp9464 := Call(__e, PrimFunc(symread_1file), V1261)


tmp9465 := Call(__e, PrimFunc(symshen_4load_1help), let__9443, tmp9464)


let__9446 := tmp9465
_ = let__9446

tmp9466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symget_1time) {
return PrimGetTime(symrun)
}
__typedArg0 := symrun
return Call(__e, PrimFunc(symget_1time), __typedArg0)
})()

let__9447 := tmp9466
_ = let__9447

let__9448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(let__9447)
__typedN1, __typedOK1 := TypedFloat64(let__9445)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := let__9447
__typedArg1 := let__9445
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})()
_ = let__9448

tmp9468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(let__9448)
}
__typedArg0 := let__9448
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp9470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\nrun time: "))
__typedS1, __typedOK1 := TypedString(tmp9468)
__typedS2, __typedOK2 := TypedString(MakeString(" secs\n"))
if __typedOK0 && __typedOK1 && __typedOK2 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + (__typedS1 + __typedS2)))
}}
__typedArg0 := MakeString("\nrun time: ")
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp9468)
__typedS1, __typedOK1 := TypedString(MakeString(" secs\n"))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp9468
__typedArg1 := MakeString(" secs\n")
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp9471 := Call(__e, PrimFunc(symstoutput))


tmp9472 := Call(__e, PrimFunc(sympr), tmp9470, tmp9471)


let__9449 := tmp9472
_ = let__9449

let__9444 := let__9446
_ = let__9444

var ifres9473 Obj

if True == let__9443 {
tmp9474 := Call(__e, PrimFunc(syminferences))


tmp9475 := Call(__e, PrimFunc(symshen_4app), tmp9474, MakeString(" inferences\n"), symshen_4a)


tmp9476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\ntypechecked in "))
__typedS1, __typedOK1 := TypedString(tmp9475)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\ntypechecked in ")
__typedArg1 := tmp9475
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp9477 := Call(__e, PrimFunc(symstoutput))


tmp9478 := Call(__e, PrimFunc(sympr), tmp9476, tmp9477)


ifres9473 = tmp9478


} else {
ifres9473 = symshen_4skip


}

let__9450 := ifres9473
_ = let__9450

__e.Return(symloaded)
return


}, 1)

tmp9479 := Call(__e, ns2_1set, symload, tmp9461)


_ = tmp9479

tmp9480 := MakeNative(func(__e *ControlFlow) {
V1272 := __e.Get(1)
_ = V1272
V1273 := __e.Get(2)
_ = V1273
tmp9482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, V1272)
}
__typedArg0 := False
__typedArg1 := V1272
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9482 {
__e.TailApply(PrimFunc(symshen_4eval_1and_1print), V1273)
return
} else {
__e.TailApply(PrimFunc(symshen_4check_1eval_1and_1print), V1273)
return
}


}, 2)

tmp9483 := Call(__e, ns2_1set, symshen_4load_1help, tmp9480)


_ = tmp9483

tmp9484 := MakeNative(func(__e *ControlFlow) {
V1274 := __e.Get(1)
_ = V1274
tmp9485 := MakeNative(func(__e *ControlFlow) {
Z1275 := __e.Get(1)
_ = Z1275
tmp9486 := Call(__e, PrimFunc(symshen_4shen_1_6kl), Z1275)


tmp9487 := Call(__e, PrimFunc(symeval_1kl), tmp9486)


tmp9488 := Call(__e, PrimFunc(symshen_4app), tmp9487, MakeString("\n"), symshen_4s)


tmp9489 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp9488, tmp9489)
return


}, 1)

__e.TailApply(PrimFunc(symshen_4for_1each), tmp9485, V1274)
return


}, 1)

tmp9490 := Call(__e, ns2_1set, symshen_4eval_1and_1print, tmp9484)


_ = tmp9490

tmp9491 := MakeNative(func(__e *ControlFlow) {
V1276 := __e.Get(1)
_ = V1276
tmp9492 := MakeNative(func(__e *ControlFlow) {
Z1278 := __e.Get(1)
_ = Z1278
__e.TailApply(PrimFunc(symshen_4typetable), Z1278)
return
}, 1)

tmp9493 := Call(__e, PrimFunc(symmapcan), tmp9492, V1276)


let__9451 := tmp9493
_ = let__9451

tmp9494 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4assumetypes), let__9451)
return
}, 0)

tmp9495 := MakeNative(func(__e *ControlFlow) {
Z1280 := __e.Get(1)
_ = Z1280
__e.TailApply(PrimFunc(symshen_4unwind_1types), Z1280, let__9451)
return
}, 1)

tmp9496 := Call(__e, try_1catch, tmp9494, tmp9495)


let__9452 := tmp9496
_ = let__9452

tmp9497 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4work_1through), V1276)
return
}, 0)

tmp9498 := MakeNative(func(__e *ControlFlow) {
Z1281 := __e.Get(1)
_ = Z1281
__e.TailApply(PrimFunc(symshen_4unwind_1types), Z1281, let__9451)
return
}, 1)

__e.TailApply(try_1catch, tmp9497, tmp9498)
return


}, 1)

tmp9499 := Call(__e, ns2_1set, symshen_4check_1eval_1and_1print, tmp9491)


_ = tmp9499

tmp9500 := MakeNative(func(__e *ControlFlow) {
V1286 := __e.Get(1)
_ = V1286
tmp9545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1286)
}
__typedArg0 := V1286
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9526 Obj

if True == tmp9545 {
tmp9543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1286)
}
__typedArg0 := V1286
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp9543)
}
__typedArg0 := symdefine
__typedArg1 := tmp9543
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9528 Obj

if True == tmp9544 {
tmp9541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1286)
}
__typedArg0 := V1286
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9541)
}
__typedArg0 := tmp9541
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9530 Obj

if True == tmp9542 {
tmp9538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1286)
}
__typedArg0 := V1286
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9538)
}
__typedArg0 := tmp9538
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9539)
}
__typedArg0 := tmp9539
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9532 Obj

if True == tmp9540 {
tmp9534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1286)
}
__typedArg0 := V1286
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9534)
}
__typedArg0 := tmp9534
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9535)
}
__typedArg0 := tmp9535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, tmp9536)
}
__typedArg0 := sym_i
__typedArg1 := tmp9536
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9533 Obj

if True == tmp9537 {
ifres9533 = True


} else {
ifres9533 = False


}

ifres9532 = ifres9533


} else {
ifres9532 = False


}

var ifres9531 Obj

if True == ifres9532 {
ifres9531 = True


} else {
ifres9531 = False


}

ifres9530 = ifres9531


} else {
ifres9530 = False


}

var ifres9529 Obj

if True == ifres9530 {
ifres9529 = True


} else {
ifres9529 = False


}

ifres9528 = ifres9529


} else {
ifres9528 = False


}

var ifres9527 Obj

if True == ifres9528 {
ifres9527 = True


} else {
ifres9527 = False


}

ifres9526 = ifres9527


} else {
ifres9526 = False


}

if True == ifres9526 {
tmp9501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1286)
}
__typedArg0 := V1286
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9501)
}
__typedArg0 := tmp9501
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1286)
}
__typedArg0 := V1286
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9503)
}
__typedArg0 := tmp9503
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1286)
}
__typedArg0 := V1286
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9505)
}
__typedArg0 := tmp9505
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9506)
}
__typedArg0 := tmp9506
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9508 := Call(__e, PrimFunc(symshen_4type_1F), tmp9504, tmp9507)


tmp9509 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp9508)


tmp9510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9509, Nil)
}
__typedArg0 := tmp9509
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9502, tmp9510)
}
__typedArg0 := tmp9502
__typedArg1 := tmp9510
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1286)
}
__typedArg0 := V1286
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9516 Obj

if True == tmp9524 {
tmp9522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1286)
}
__typedArg0 := V1286
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp9522)
}
__typedArg0 := symdefine
__typedArg1 := tmp9522
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9518 Obj

if True == tmp9523 {
tmp9520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1286)
}
__typedArg0 := V1286
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9520)
}
__typedArg0 := tmp9520
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9519 Obj

if True == tmp9521 {
ifres9519 = True


} else {
ifres9519 = False


}

ifres9518 = ifres9519


} else {
ifres9518 = False


}

var ifres9517 Obj

if True == ifres9518 {
ifres9517 = True


} else {
ifres9517 = False


}

ifres9516 = ifres9517


} else {
ifres9516 = False


}

if True == ifres9516 {
tmp9511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1286)
}
__typedArg0 := V1286
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9511)
}
__typedArg0 := tmp9511
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9513 := Call(__e, PrimFunc(symshen_4app), tmp9512, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing { in "))
__typedS1, __typedOK1 := TypedString(tmp9513)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing { in ")
__typedArg1 := tmp9513
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing { in "))
__typedS1, __typedOK1 := TypedString(tmp9513)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing { in ")
__typedArg1 := tmp9513
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

tmp9546 := Call(__e, ns2_1set, symshen_4typetable, tmp9500)


_ = tmp9546

tmp9547 := MakeNative(func(__e *ControlFlow) {
V1293 := __e.Get(1)
_ = V1293
V1294 := __e.Get(2)
_ = V1294
tmp9560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1294)
}
__typedArg0 := V1294
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9556 Obj

if True == tmp9560 {
tmp9558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1294)
}
__typedArg0 := V1294
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, tmp9558)
}
__typedArg0 := sym_j
__typedArg1 := tmp9558
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9557 Obj

if True == tmp9559 {
ifres9557 = True


} else {
ifres9557 = False


}

ifres9556 = ifres9557


} else {
ifres9556 = False


}

if True == ifres9556 {
__e.Return(Nil)
return
} else {
tmp9554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1294)
}
__typedArg0 := V1294
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9554 {
tmp9548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1294)
}
__typedArg0 := V1294
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1294)
}
__typedArg0 := V1294
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9550 := Call(__e, PrimFunc(symshen_4type_1F), V1293, tmp9549)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9548, tmp9550)
}
__typedArg0 := tmp9548
__typedArg1 := tmp9550
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9551 := Call(__e, PrimFunc(symshen_4app), V1293, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing } in "))
__typedS1, __typedOK1 := TypedString(tmp9551)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing } in ")
__typedArg1 := tmp9551
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing } in "))
__typedS1, __typedOK1 := TypedString(tmp9551)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing } in ")
__typedArg1 := tmp9551
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}


}, 2)

tmp9561 := Call(__e, ns2_1set, symshen_4type_1F, tmp9547)


_ = tmp9561

tmp9562 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V1297 := __self1
_ = V1297
tmp9576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1297)
}
__typedArg0 := Nil
__typedArg1 := V1297
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9576 {
__e.Return(Nil)
return
} else {
tmp9574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1297)
}
__typedArg0 := V1297
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9570 Obj

if True == tmp9574 {
tmp9572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1297)
}
__typedArg0 := V1297
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9572)
}
__typedArg0 := tmp9572
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9571 Obj

if True == tmp9573 {
ifres9571 = True


} else {
ifres9571 = False


}

ifres9570 = ifres9571


} else {
ifres9570 = False


}

if True == ifres9570 {
tmp9563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1297)
}
__typedArg0 := V1297
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1297)
}
__typedArg0 := V1297
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9564)
}
__typedArg0 := tmp9564
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9566 := Call(__e, PrimFunc(symdeclare), tmp9563, tmp9565)


_ = tmp9566

tmp9567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1297)
}
__typedArg0 := V1297
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9567)
}
__typedArg0 := tmp9567
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4assumetypes) == __self {
__self1 = tmp9568
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4assumetypes), tmp9568)
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

tmp9577 := Call(__e, ns2_1set, symshen_4assumetypes, tmp9562)


_ = tmp9577

tmp9578 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V1302 := __self1
_ = V1302
V1303 := __self2
_ = V1303
tmp9589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1303)
}
__typedArg0 := V1303
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9585 Obj

if True == tmp9589 {
tmp9587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1303)
}
__typedArg0 := V1303
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9587)
}
__typedArg0 := tmp9587
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9586 Obj

if True == tmp9588 {
ifres9586 = True


} else {
ifres9586 = False


}

ifres9585 = ifres9586


} else {
ifres9585 = False


}

if True == ifres9585 {
tmp9579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1303)
}
__typedArg0 := V1303
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9580 := Call(__e, PrimFunc(symdestroy), tmp9579)


_ = tmp9580

tmp9581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1303)
}
__typedArg0 := V1303
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9581)
}
__typedArg0 := tmp9581
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4unwind_1types) == __self {
__self1, __self2 = V1302, tmp9582
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4unwind_1types), V1302, tmp9582)
return


} else {
tmp9583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symerror_1to_1string) {
return PrimErrorToString(V1302)
}
__typedArg0 := V1302
return Call(__e, PrimFunc(symerror_1to_1string), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp9583)
}
__typedArg0 := tmp9583
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 2)

tmp9590 := Call(__e, ns2_1set, symshen_4unwind_1types, tmp9578)


_ = tmp9590

tmp9591 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V1306 := __self1
_ = V1306
tmp9637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1306)
}
__typedArg0 := Nil
__typedArg1 := V1306
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9637 {
__e.Return(Nil)
return
} else {
tmp9635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1306)
}
__typedArg0 := V1306
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9620 Obj

if True == tmp9635 {
tmp9633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1306)
}
__typedArg0 := V1306
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9633)
}
__typedArg0 := tmp9633
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9622 Obj

if True == tmp9634 {
tmp9630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1306)
}
__typedArg0 := V1306
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9630)
}
__typedArg0 := tmp9630
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9631)
}
__typedArg0 := tmp9631
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9624 Obj

if True == tmp9632 {
tmp9626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1306)
}
__typedArg0 := V1306
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9626)
}
__typedArg0 := tmp9626
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp9629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp9627, tmp9628)
}
__typedArg0 := tmp9627
__typedArg1 := tmp9628
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9625 Obj

if True == tmp9629 {
ifres9625 = True


} else {
ifres9625 = False


}

ifres9624 = ifres9625


} else {
ifres9624 = False


}

var ifres9623 Obj

if True == ifres9624 {
ifres9623 = True


} else {
ifres9623 = False


}

ifres9622 = ifres9623


} else {
ifres9622 = False


}

var ifres9621 Obj

if True == ifres9622 {
ifres9621 = True


} else {
ifres9621 = False


}

ifres9620 = ifres9621


} else {
ifres9620 = False


}

if True == ifres9620 {
tmp9592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1306)
}
__typedArg0 := V1306
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1306)
}
__typedArg0 := V1306
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9593)
}
__typedArg0 := tmp9593
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9594)
}
__typedArg0 := tmp9594
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9596 := Call(__e, PrimFunc(symshen_4typecheck), tmp9592, tmp9595)


let__9453 := tmp9596
_ = let__9453

tmp9610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__9453, False)
}
__typedArg0 := let__9453
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9610 {
__e.TailApply(PrimFunc(symshen_4type_1error))
return
} else {
tmp9597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1306)
}
__typedArg0 := V1306
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9598 := Call(__e, PrimFunc(symshen_4shen_1_6kl), tmp9597)


tmp9599 := Call(__e, PrimFunc(symeval_1kl), tmp9598)


let__9454 := tmp9599
_ = let__9454

tmp9600 := Call(__e, PrimFunc(symshen_4pretty_1type), let__9453)


tmp9601 := Call(__e, PrimFunc(symshen_4app), tmp9600, MakeString("\n"), symshen_4r)


tmp9603 := Call(__e, PrimFunc(symshen_4app), let__9454, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" : "))
__typedS1, __typedOK1 := TypedString(tmp9601)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" : ")
__typedArg1 := tmp9601
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4s)


tmp9604 := Call(__e, PrimFunc(symstoutput))


tmp9605 := Call(__e, PrimFunc(sympr), tmp9603, tmp9604)


let__9455 := tmp9605
_ = let__9455

tmp9606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1306)
}
__typedArg0 := V1306
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9606)
}
__typedArg0 := tmp9606
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9607)
}
__typedArg0 := tmp9607
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4work_1through) == __self {
__self1 = tmp9608
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4work_1through), tmp9608)
return


}


} else {
tmp9618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1306)
}
__typedArg0 := V1306
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9618 {
tmp9611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1306)
}
__typedArg0 := V1306
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp9613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1306)
}
__typedArg0 := V1306
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp9613)
}
__typedArg0 := symA
__typedArg1 := tmp9613
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9612, tmp9614)
}
__typedArg0 := tmp9612
__typedArg1 := tmp9614
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9611, tmp9615)
}
__typedArg0 := tmp9611
__typedArg1 := tmp9615
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4work_1through) == __self {
__self1 = tmp9616
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4work_1through), tmp9616)
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

tmp9638 := Call(__e, ns2_1set, symshen_4work_1through, tmp9591)


_ = tmp9638

tmp9639 := MakeNative(func(__e *ControlFlow) {
V1311 := __e.Get(1)
_ = V1311
tmp9781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1311)
}
__typedArg0 := V1311
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9655 Obj

if True == tmp9781 {
tmp9779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1311)
}
__typedArg0 := V1311
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9779)
}
__typedArg0 := tmp9779
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9657 Obj

if True == tmp9780 {
tmp9776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1311)
}
__typedArg0 := V1311
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9776)
}
__typedArg0 := tmp9776
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp9777)
}
__typedArg0 := symlist
__typedArg1 := tmp9777
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9659 Obj

if True == tmp9778 {
tmp9773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1311)
}
__typedArg0 := V1311
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9773)
}
__typedArg0 := tmp9773
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9774)
}
__typedArg0 := tmp9774
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9661 Obj

if True == tmp9775 {
tmp9769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1311)
}
__typedArg0 := V1311
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9769)
}
__typedArg0 := tmp9769
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9770)
}
__typedArg0 := tmp9770
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9771)
}
__typedArg0 := Nil
__typedArg1 := tmp9771
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9663 Obj

if True == tmp9772 {
tmp9767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1311)
}
__typedArg0 := V1311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9767)
}
__typedArg0 := tmp9767
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9665 Obj

if True == tmp9768 {
tmp9764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1311)
}
__typedArg0 := V1311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9764)
}
__typedArg0 := tmp9764
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1_1_6, tmp9765)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp9765
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9667 Obj

if True == tmp9766 {
tmp9761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1311)
}
__typedArg0 := V1311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9761)
}
__typedArg0 := tmp9761
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9762)
}
__typedArg0 := tmp9762
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9669 Obj

if True == tmp9763 {
tmp9757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1311)
}
__typedArg0 := V1311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9757)
}
__typedArg0 := tmp9757
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9758)
}
__typedArg0 := tmp9758
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9759)
}
__typedArg0 := tmp9759
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9671 Obj

if True == tmp9760 {
tmp9752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1311)
}
__typedArg0 := V1311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9752)
}
__typedArg0 := tmp9752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9753)
}
__typedArg0 := tmp9753
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9754)
}
__typedArg0 := tmp9754
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symstr, tmp9755)
}
__typedArg0 := symstr
__typedArg1 := tmp9755
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9673 Obj

if True == tmp9756 {
tmp9747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1311)
}
__typedArg0 := V1311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9747)
}
__typedArg0 := tmp9747
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9748)
}
__typedArg0 := tmp9748
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9749)
}
__typedArg0 := tmp9749
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9750)
}
__typedArg0 := tmp9750
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9675 Obj

if True == tmp9751 {
tmp9741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1311)
}
__typedArg0 := V1311
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9745)
}
__typedArg0 := tmp9745
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9677 Obj

if True == tmp9746 {
tmp9734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1311)
}
__typedArg0 := V1311
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9738)
}
__typedArg0 := tmp9738
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp9739)
}
__typedArg0 := symlist
__typedArg1 := tmp9739
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9679 Obj

if True == tmp9740 {
tmp9727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1311)
}
__typedArg0 := V1311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9727)
}
__typedArg0 := tmp9727
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9728)
}
__typedArg0 := tmp9728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9729)
}
__typedArg0 := tmp9729
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9730)
}
__typedArg0 := tmp9730
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9731)
}
__typedArg0 := tmp9731
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9732)
}
__typedArg0 := tmp9732
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9681 Obj

if True == tmp9733 {
tmp9719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1311)
}
__typedArg0 := V1311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9719)
}
__typedArg0 := tmp9719
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9720)
}
__typedArg0 := tmp9720
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9721)
}
__typedArg0 := tmp9721
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9722)
}
__typedArg0 := tmp9722
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9723)
}
__typedArg0 := tmp9723
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9724)
}
__typedArg0 := tmp9724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9725)
}
__typedArg0 := Nil
__typedArg1 := tmp9725
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9683 Obj

if True == tmp9726 {
tmp9713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1311)
}
__typedArg0 := V1311
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9717)
}
__typedArg0 := tmp9717
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9685 Obj

if True == tmp9718 {
tmp9706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1311)
}
__typedArg0 := V1311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9706)
}
__typedArg0 := tmp9706
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9707)
}
__typedArg0 := tmp9707
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9708)
}
__typedArg0 := tmp9708
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

var ifres9687 Obj

if True == tmp9712 {
tmp9702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1311)
}
__typedArg0 := V1311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9702)
}
__typedArg0 := tmp9702
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9703)
}
__typedArg0 := tmp9703
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9704)
}
__typedArg0 := Nil
__typedArg1 := tmp9704
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9689 Obj

if True == tmp9705 {
tmp9691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1311)
}
__typedArg0 := V1311
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9691)
}
__typedArg0 := tmp9691
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9692)
}
__typedArg0 := tmp9692
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1311)
}
__typedArg0 := V1311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9694)
}
__typedArg0 := tmp9694
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9695)
}
__typedArg0 := tmp9695
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9696)
}
__typedArg0 := tmp9696
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9697)
}
__typedArg0 := tmp9697
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp9693, tmp9700)
}
__typedArg0 := tmp9693
__typedArg1 := tmp9700
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9690 Obj

if True == tmp9701 {
ifres9690 = True


} else {
ifres9690 = False


}

ifres9689 = ifres9690


} else {
ifres9689 = False


}

var ifres9688 Obj

if True == ifres9689 {
ifres9688 = True


} else {
ifres9688 = False


}

ifres9687 = ifres9688


} else {
ifres9687 = False


}

var ifres9686 Obj

if True == ifres9687 {
ifres9686 = True


} else {
ifres9686 = False


}

ifres9685 = ifres9686


} else {
ifres9685 = False


}

var ifres9684 Obj

if True == ifres9685 {
ifres9684 = True


} else {
ifres9684 = False


}

ifres9683 = ifres9684


} else {
ifres9683 = False


}

var ifres9682 Obj

if True == ifres9683 {
ifres9682 = True


} else {
ifres9682 = False


}

ifres9681 = ifres9682


} else {
ifres9681 = False


}

var ifres9680 Obj

if True == ifres9681 {
ifres9680 = True


} else {
ifres9680 = False


}

ifres9679 = ifres9680


} else {
ifres9679 = False


}

var ifres9678 Obj

if True == ifres9679 {
ifres9678 = True


} else {
ifres9678 = False


}

ifres9677 = ifres9678


} else {
ifres9677 = False


}

var ifres9676 Obj

if True == ifres9677 {
ifres9676 = True


} else {
ifres9676 = False


}

ifres9675 = ifres9676


} else {
ifres9675 = False


}

var ifres9674 Obj

if True == ifres9675 {
ifres9674 = True


} else {
ifres9674 = False


}

ifres9673 = ifres9674


} else {
ifres9673 = False


}

var ifres9672 Obj

if True == ifres9673 {
ifres9672 = True


} else {
ifres9672 = False


}

ifres9671 = ifres9672


} else {
ifres9671 = False


}

var ifres9670 Obj

if True == ifres9671 {
ifres9670 = True


} else {
ifres9670 = False


}

ifres9669 = ifres9670


} else {
ifres9669 = False


}

var ifres9668 Obj

if True == ifres9669 {
ifres9668 = True


} else {
ifres9668 = False


}

ifres9667 = ifres9668


} else {
ifres9667 = False


}

var ifres9666 Obj

if True == ifres9667 {
ifres9666 = True


} else {
ifres9666 = False


}

ifres9665 = ifres9666


} else {
ifres9665 = False


}

var ifres9664 Obj

if True == ifres9665 {
ifres9664 = True


} else {
ifres9664 = False


}

ifres9663 = ifres9664


} else {
ifres9663 = False


}

var ifres9662 Obj

if True == ifres9663 {
ifres9662 = True


} else {
ifres9662 = False


}

ifres9661 = ifres9662


} else {
ifres9661 = False


}

var ifres9660 Obj

if True == ifres9661 {
ifres9660 = True


} else {
ifres9660 = False


}

ifres9659 = ifres9660


} else {
ifres9659 = False


}

var ifres9658 Obj

if True == ifres9659 {
ifres9658 = True


} else {
ifres9658 = False


}

ifres9657 = ifres9658


} else {
ifres9657 = False


}

var ifres9656 Obj

if True == ifres9657 {
ifres9656 = True


} else {
ifres9656 = False


}

ifres9655 = ifres9656


} else {
ifres9655 = False


}

if True == ifres9655 {
tmp9640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1311)
}
__typedArg0 := V1311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9640)
}
__typedArg0 := tmp9640
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9641)
}
__typedArg0 := tmp9641
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9642)
}
__typedArg0 := tmp9642
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9643)
}
__typedArg0 := tmp9643
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1311)
}
__typedArg0 := V1311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9645)
}
__typedArg0 := tmp9645
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9646)
}
__typedArg0 := tmp9646
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9647)
}
__typedArg0 := tmp9647
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9648)
}
__typedArg0 := tmp9648
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a_6, tmp9649)
}
__typedArg0 := sym_a_a_6
__typedArg1 := tmp9649
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9644, tmp9650)
}
__typedArg0 := tmp9644
__typedArg1 := tmp9650
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1311)
}
__typedArg0 := V1311
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9653 {
tmp9651 := MakeNative(func(__e *ControlFlow) {
Z1312 := __e.Get(1)
_ = Z1312
__e.TailApply(PrimFunc(symshen_4pretty_1type), Z1312)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp9651, V1311)
return


} else {
__e.Return(V1311)
return
}


}


}, 1)

tmp9782 := Call(__e, ns2_1set, symshen_4pretty_1type, tmp9639)


_ = tmp9782

tmp9783 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("type error\n"))
}
__typedArg0 := MakeString("type error\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}, 0)

tmp9784 := Call(__e, ns2_1set, symshen_4type_1error, tmp9783)


_ = tmp9784

tmp9785 := MakeNative(func(__e *ControlFlow) {
V1313 := __e.Get(1)
_ = V1313
tmp9786 := Call(__e, PrimFunc(symshen_4klfile), V1313)


let__9456 := tmp9786
_ = let__9456

tmp9787 := Call(__e, PrimFunc(symread_1file), V1313)


let__9457 := tmp9787
_ = let__9457

tmp9788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symopen) {
return PrimOpenStream(let__9456, symout)
}
__typedArg0 := let__9456
__typedArg1 := symout
return Call(__e, PrimFunc(symopen), __typedArg0, __typedArg1)
})()

let__9458 := tmp9788
_ = let__9458

tmp9789 := MakeNative(func(__e *ControlFlow) {
Z1318 := __e.Get(1)
_ = Z1318
tmp9790 := Call(__e, PrimFunc(symshen_4shen_1_6kl_1h), Z1318)


__e.TailApply(PrimFunc(symshen_4partial), tmp9790)
return


}, 1)

tmp9791 := Call(__e, PrimFunc(symmap), tmp9789, let__9457)


let__9459 := tmp9791
_ = let__9459

tmp9792 := Call(__e, PrimFunc(symshen_4write_1kl), let__9459, let__9458)


let__9460 := tmp9792
_ = let__9460

__e.Return(let__9456)
return


}, 1)

tmp9793 := Call(__e, ns2_1set, symbootstrap, tmp9785)


_ = tmp9793

tmp9794 := MakeNative(func(__e *ControlFlow) {
V1320 := __e.Get(1)
_ = V1320
tmp9817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1320)
}
__typedArg0 := V1320
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9804 Obj

if True == tmp9817 {
tmp9815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1320)
}
__typedArg0 := V1320
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4f_1error, tmp9815)
}
__typedArg0 := symshen_4f_1error
__typedArg1 := tmp9815
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9806 Obj

if True == tmp9816 {
tmp9813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1320)
}
__typedArg0 := V1320
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9813)
}
__typedArg0 := tmp9813
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9808 Obj

if True == tmp9814 {
tmp9810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1320)
}
__typedArg0 := V1320
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9810)
}
__typedArg0 := tmp9810
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9811)
}
__typedArg0 := Nil
__typedArg1 := tmp9811
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9809 Obj

if True == tmp9812 {
ifres9809 = True


} else {
ifres9809 = False


}

ifres9808 = ifres9809


} else {
ifres9808 = False


}

var ifres9807 Obj

if True == ifres9808 {
ifres9807 = True


} else {
ifres9807 = False


}

ifres9806 = ifres9807


} else {
ifres9806 = False


}

var ifres9805 Obj

if True == ifres9806 {
ifres9805 = True


} else {
ifres9805 = False


}

ifres9804 = ifres9805


} else {
ifres9804 = False


}

if True == ifres9804 {
tmp9795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1320)
}
__typedArg0 := V1320
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9795)
}
__typedArg0 := tmp9795
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(tmp9796)
}
__typedArg0 := tmp9796
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp9799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("partial function "))
__typedS1, __typedOK1 := TypedString(tmp9797)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("partial function ")
__typedArg1 := tmp9797
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), Nil)
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("partial function "))
__typedS1, __typedOK1 := TypedString(tmp9797)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("partial function ")
__typedArg1 := tmp9797
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp9799)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp9799
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1320)
}
__typedArg0 := V1320
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9802 {
tmp9800 := MakeNative(func(__e *ControlFlow) {
Z1321 := __e.Get(1)
_ = Z1321
__e.TailApply(PrimFunc(symshen_4partial), Z1321)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp9800, V1320)
return


} else {
__e.Return(V1320)
return
}


}


}, 1)

tmp9818 := Call(__e, ns2_1set, symshen_4partial, tmp9794)


_ = tmp9818

tmp9819 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V1324 := __self1
_ = V1324
V1325 := __self2
_ = V1325
tmp9833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1324)
}
__typedArg0 := Nil
__typedArg1 := V1324
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9833 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symclose) {
return PrimCloseStream(V1325)
}
__typedArg0 := V1325
return Call(__e, PrimFunc(symclose), __typedArg0)
})())
return
} else {
tmp9831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9827 Obj

if True == tmp9831 {
tmp9829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9829)
}
__typedArg0 := tmp9829
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9828 Obj

if True == tmp9830 {
ifres9828 = True


} else {
ifres9828 = False


}

ifres9827 = ifres9828


} else {
ifres9827 = False


}

if True == ifres9827 {
tmp9820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9822 := Call(__e, PrimFunc(symshen_4write_1kl_1h), tmp9821, V1325)


_ = tmp9822

if PrimFunc(symshen_4write_1kl) == __self {
__self1, __self2 = tmp9820, V1325
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4write_1kl), tmp9820, V1325)
return


} else {
tmp9825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9825 {
tmp9823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4write_1kl) == __self {
__self1, __self2 = tmp9823, V1325
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4write_1kl), tmp9823, V1325)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4write_1kl)
return
}


}


}


}, 2)

tmp9834 := Call(__e, ns2_1set, symshen_4write_1kl, tmp9819)


_ = tmp9834

tmp9835 := MakeNative(func(__e *ControlFlow) {
V1328 := __e.Get(1)
_ = V1328
V1329 := __e.Get(2)
_ = V1329
tmp9875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1328)
}
__typedArg0 := V1328
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9838 Obj

if True == tmp9875 {
tmp9873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1328)
}
__typedArg0 := V1328
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp9873)
}
__typedArg0 := symdefun
__typedArg1 := tmp9873
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9840 Obj

if True == tmp9874 {
tmp9871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1328)
}
__typedArg0 := V1328
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9871)
}
__typedArg0 := tmp9871
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9842 Obj

if True == tmp9872 {
tmp9868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1328)
}
__typedArg0 := V1328
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9868)
}
__typedArg0 := tmp9868
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail, tmp9869)
}
__typedArg0 := symfail
__typedArg1 := tmp9869
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9844 Obj

if True == tmp9870 {
tmp9865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1328)
}
__typedArg0 := V1328
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

var ifres9846 Obj

if True == tmp9867 {
tmp9861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1328)
}
__typedArg0 := V1328
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9862)
}
__typedArg0 := tmp9862
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9863)
}
__typedArg0 := Nil
__typedArg1 := tmp9863
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9848 Obj

if True == tmp9864 {
tmp9857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1328)
}
__typedArg0 := V1328
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9857)
}
__typedArg0 := tmp9857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9858)
}
__typedArg0 := tmp9858
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9859)
}
__typedArg0 := tmp9859
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9850 Obj

if True == tmp9860 {
tmp9852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1328)
}
__typedArg0 := V1328
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9852)
}
__typedArg0 := tmp9852
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9853)
}
__typedArg0 := tmp9853
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9854)
}
__typedArg0 := tmp9854
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9855)
}
__typedArg0 := Nil
__typedArg1 := tmp9855
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9851 Obj

if True == tmp9856 {
ifres9851 = True


} else {
ifres9851 = False


}

ifres9850 = ifres9851


} else {
ifres9850 = False


}

var ifres9849 Obj

if True == ifres9850 {
ifres9849 = True


} else {
ifres9849 = False


}

ifres9848 = ifres9849


} else {
ifres9848 = False


}

var ifres9847 Obj

if True == ifres9848 {
ifres9847 = True


} else {
ifres9847 = False


}

ifres9846 = ifres9847


} else {
ifres9846 = False


}

var ifres9845 Obj

if True == ifres9846 {
ifres9845 = True


} else {
ifres9845 = False


}

ifres9844 = ifres9845


} else {
ifres9844 = False


}

var ifres9843 Obj

if True == ifres9844 {
ifres9843 = True


} else {
ifres9843 = False


}

ifres9842 = ifres9843


} else {
ifres9842 = False


}

var ifres9841 Obj

if True == ifres9842 {
ifres9841 = True


} else {
ifres9841 = False


}

ifres9840 = ifres9841


} else {
ifres9840 = False


}

var ifres9839 Obj

if True == ifres9840 {
ifres9839 = True


} else {
ifres9839 = False


}

ifres9838 = ifres9839


} else {
ifres9838 = False


}

if True == ifres9838 {
__e.TailApply(PrimFunc(sympr), MakeString("(defun fail () shen.fail!)"), V1329)
return
} else {
tmp9836 := Call(__e, PrimFunc(symshen_4app), V1328, MakeString("\n\n"), symshen_4r)


__e.TailApply(PrimFunc(sympr), tmp9836, V1329)
return


}


}, 2)

tmp9876 := Call(__e, ns2_1set, symshen_4write_1kl_1h, tmp9835)


_ = tmp9876

tmp9877 := MakeNative(func(__e *ControlFlow) {
V1330 := __e.Get(1)
_ = V1330
tmp9886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V1330)
}
__typedArg0 := MakeString("")
__typedArg1 := V1330
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9886 {
__e.Return(MakeString(".kl"))
return
} else {
tmp9884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(".shen"), V1330)
}
__typedArg0 := MakeString(".shen")
__typedArg1 := V1330
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9884 {
__e.Return(MakeString(".kl"))
return
} else {
tmp9882 := Call(__e, PrimFunc(symshen_4_7string_2), V1330)


if True == tmp9882 {
tmp9878 := Call(__e, PrimFunc(symhdstr), V1330)


tmp9880 := Call(__e, PrimFunc(symshen_4klfile), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V1330)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V1330
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.TailApply(PrimFunc(sym_8s), tmp9878, tmp9880)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4klfile)
return
}


}


}


}, 1)

__e.TailApply(ns2_1set, symshen_4klfile, tmp9877)
return




}, 0)

