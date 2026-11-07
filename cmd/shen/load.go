package main

import . "github.com/pyrex41/shen-go/kl"

var LoadMain = MakeNative(func(__e *ControlFlow) {
tmp9488 := MakeNative(func(__e *ControlFlow) {
V1262 := __e.Get(1)
_ = V1262
tmp9489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtc_d)
}
__typedArg0 := symshen_4_dtc_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__9470 := tmp9489
_ = let__9470

tmp9490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symget_1time) {
return PrimGetTime(symrun)
}
__typedArg0 := symrun
return Call(__e, PrimFunc(symget_1time), __typedArg0)
})()

let__9472 := tmp9490
_ = let__9472

tmp9491 := Call(__e, PrimFunc(symread_1file), V1262)


tmp9492 := Call(__e, PrimFunc(symshen_4load_1help), let__9470, tmp9491)


let__9473 := tmp9492
_ = let__9473

tmp9493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symget_1time) {
return PrimGetTime(symrun)
}
__typedArg0 := symrun
return Call(__e, PrimFunc(symget_1time), __typedArg0)
})()

let__9474 := tmp9493
_ = let__9474

let__9475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(let__9474)
__typedN1, __typedOK1 := TypedFloat64(let__9472)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := let__9474
__typedArg1 := let__9472
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})()
_ = let__9475

tmp9495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(let__9475)
}
__typedArg0 := let__9475
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp9497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\nrun time: "))
__typedS1, __typedOK1 := TypedString(tmp9495)
__typedS2, __typedOK2 := TypedString(MakeString(" secs\n"))
if __typedOK0 && __typedOK1 && __typedOK2 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + (__typedS1 + __typedS2)))
}}
__typedArg0 := MakeString("\nrun time: ")
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(tmp9495)
__typedS1, __typedOK1 := TypedString(MakeString(" secs\n"))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := tmp9495
__typedArg1 := MakeString(" secs\n")
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp9498 := Call(__e, PrimFunc(symstoutput))


tmp9499 := Call(__e, PrimFunc(sympr), tmp9497, tmp9498)


let__9476 := tmp9499
_ = let__9476

let__9471 := let__9473
_ = let__9471

var ifres9500 Obj

if True == let__9470 {
tmp9501 := Call(__e, PrimFunc(syminferences))


tmp9502 := Call(__e, PrimFunc(symshen_4app), tmp9501, MakeString(" inferences\n"), symshen_4a)


tmp9503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\ntypechecked in "))
__typedS1, __typedOK1 := TypedString(tmp9502)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\ntypechecked in ")
__typedArg1 := tmp9502
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp9504 := Call(__e, PrimFunc(symstoutput))


tmp9505 := Call(__e, PrimFunc(sympr), tmp9503, tmp9504)


ifres9500 = tmp9505


} else {
ifres9500 = symshen_4skip


}

let__9477 := ifres9500
_ = let__9477

__e.Return(symloaded)
return


}, 1)

tmp9506 := Call(__e, ns2_1set, symload, tmp9488)


_ = tmp9506

tmp9507 := MakeNative(func(__e *ControlFlow) {
V1273 := __e.Get(1)
_ = V1273
V1274 := __e.Get(2)
_ = V1274
tmp9509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, V1273)
}
__typedArg0 := False
__typedArg1 := V1273
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9509 {
__e.TailApply(PrimFunc(symshen_4eval_1and_1print), V1274)
return
} else {
__e.TailApply(PrimFunc(symshen_4check_1eval_1and_1print), V1274)
return
}


}, 2)

tmp9510 := Call(__e, ns2_1set, symshen_4load_1help, tmp9507)


_ = tmp9510

tmp9511 := MakeNative(func(__e *ControlFlow) {
V1275 := __e.Get(1)
_ = V1275
tmp9512 := MakeNative(func(__e *ControlFlow) {
Z1276 := __e.Get(1)
_ = Z1276
tmp9513 := Call(__e, PrimFunc(symshen_4shen_1_6kl), Z1276)


tmp9514 := Call(__e, PrimFunc(symeval_1kl), tmp9513)


tmp9515 := Call(__e, PrimFunc(symshen_4app), tmp9514, MakeString("\n"), symshen_4s)


tmp9516 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp9515, tmp9516)
return


}, 1)

__e.TailApply(PrimFunc(symshen_4for_1each), tmp9512, V1275)
return


}, 1)

tmp9517 := Call(__e, ns2_1set, symshen_4eval_1and_1print, tmp9511)


_ = tmp9517

tmp9518 := MakeNative(func(__e *ControlFlow) {
V1277 := __e.Get(1)
_ = V1277
tmp9519 := MakeNative(func(__e *ControlFlow) {
Z1279 := __e.Get(1)
_ = Z1279
__e.TailApply(PrimFunc(symshen_4typetable), Z1279)
return
}, 1)

tmp9520 := Call(__e, PrimFunc(symmapcan), tmp9519, V1277)


let__9478 := tmp9520
_ = let__9478

tmp9521 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4assumetypes), let__9478)
return
}, 0)

tmp9522 := MakeNative(func(__e *ControlFlow) {
Z1281 := __e.Get(1)
_ = Z1281
__e.TailApply(PrimFunc(symshen_4unwind_1types), Z1281, let__9478)
return
}, 1)

tmp9523 := Call(__e, try_1catch, tmp9521, tmp9522)


let__9479 := tmp9523
_ = let__9479

tmp9524 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4work_1through), V1277)
return
}, 0)

tmp9525 := MakeNative(func(__e *ControlFlow) {
Z1282 := __e.Get(1)
_ = Z1282
__e.TailApply(PrimFunc(symshen_4unwind_1types), Z1282, let__9478)
return
}, 1)

__e.TailApply(try_1catch, tmp9524, tmp9525)
return


}, 1)

tmp9526 := Call(__e, ns2_1set, symshen_4check_1eval_1and_1print, tmp9518)


_ = tmp9526

tmp9527 := MakeNative(func(__e *ControlFlow) {
V1287 := __e.Get(1)
_ = V1287
tmp9572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1287)
}
__typedArg0 := V1287
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9553 Obj

if True == tmp9572 {
tmp9570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1287)
}
__typedArg0 := V1287
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp9570)
}
__typedArg0 := symdefine
__typedArg1 := tmp9570
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9555 Obj

if True == tmp9571 {
tmp9568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1287)
}
__typedArg0 := V1287
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9568)
}
__typedArg0 := tmp9568
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9557 Obj

if True == tmp9569 {
tmp9565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1287)
}
__typedArg0 := V1287
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9565)
}
__typedArg0 := tmp9565
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9566)
}
__typedArg0 := tmp9566
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9559 Obj

if True == tmp9567 {
tmp9561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1287)
}
__typedArg0 := V1287
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9561)
}
__typedArg0 := tmp9561
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9562)
}
__typedArg0 := tmp9562
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_i, tmp9563)
}
__typedArg0 := sym_i
__typedArg1 := tmp9563
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9560 Obj

if True == tmp9564 {
ifres9560 = True


} else {
ifres9560 = False


}

ifres9559 = ifres9560


} else {
ifres9559 = False


}

var ifres9558 Obj

if True == ifres9559 {
ifres9558 = True


} else {
ifres9558 = False


}

ifres9557 = ifres9558


} else {
ifres9557 = False


}

var ifres9556 Obj

if True == ifres9557 {
ifres9556 = True


} else {
ifres9556 = False


}

ifres9555 = ifres9556


} else {
ifres9555 = False


}

var ifres9554 Obj

if True == ifres9555 {
ifres9554 = True


} else {
ifres9554 = False


}

ifres9553 = ifres9554


} else {
ifres9553 = False


}

if True == ifres9553 {
tmp9528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1287)
}
__typedArg0 := V1287
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9528)
}
__typedArg0 := tmp9528
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1287)
}
__typedArg0 := V1287
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9530)
}
__typedArg0 := tmp9530
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1287)
}
__typedArg0 := V1287
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9532)
}
__typedArg0 := tmp9532
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9533)
}
__typedArg0 := tmp9533
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9535 := Call(__e, PrimFunc(symshen_4type_1F), tmp9531, tmp9534)


tmp9536 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp9535)


tmp9537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9536, Nil)
}
__typedArg0 := tmp9536
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9529, tmp9537)
}
__typedArg0 := tmp9529
__typedArg1 := tmp9537
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1287)
}
__typedArg0 := V1287
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9543 Obj

if True == tmp9551 {
tmp9549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1287)
}
__typedArg0 := V1287
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp9549)
}
__typedArg0 := symdefine
__typedArg1 := tmp9549
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9545 Obj

if True == tmp9550 {
tmp9547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1287)
}
__typedArg0 := V1287
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9547)
}
__typedArg0 := tmp9547
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9546 Obj

if True == tmp9548 {
ifres9546 = True


} else {
ifres9546 = False


}

ifres9545 = ifres9546


} else {
ifres9545 = False


}

var ifres9544 Obj

if True == ifres9545 {
ifres9544 = True


} else {
ifres9544 = False


}

ifres9543 = ifres9544


} else {
ifres9543 = False


}

if True == ifres9543 {
tmp9538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1287)
}
__typedArg0 := V1287
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9538)
}
__typedArg0 := tmp9538
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9540 := Call(__e, PrimFunc(symshen_4app), tmp9539, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing { in "))
__typedS1, __typedOK1 := TypedString(tmp9540)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing { in ")
__typedArg1 := tmp9540
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing { in "))
__typedS1, __typedOK1 := TypedString(tmp9540)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing { in ")
__typedArg1 := tmp9540
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

tmp9573 := Call(__e, ns2_1set, symshen_4typetable, tmp9527)


_ = tmp9573

tmp9574 := MakeNative(func(__e *ControlFlow) {
V1294 := __e.Get(1)
_ = V1294
V1295 := __e.Get(2)
_ = V1295
tmp9587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1295)
}
__typedArg0 := V1295
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9583 Obj

if True == tmp9587 {
tmp9585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1295)
}
__typedArg0 := V1295
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_j, tmp9585)
}
__typedArg0 := sym_j
__typedArg1 := tmp9585
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9584 Obj

if True == tmp9586 {
ifres9584 = True


} else {
ifres9584 = False


}

ifres9583 = ifres9584


} else {
ifres9583 = False


}

if True == ifres9583 {
__e.Return(Nil)
return
} else {
tmp9581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1295)
}
__typedArg0 := V1295
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9581 {
tmp9575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1295)
}
__typedArg0 := V1295
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1295)
}
__typedArg0 := V1295
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9577 := Call(__e, PrimFunc(symshen_4type_1F), V1294, tmp9576)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9575, tmp9577)
}
__typedArg0 := tmp9575
__typedArg1 := tmp9577
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9578 := Call(__e, PrimFunc(symshen_4app), V1294, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing } in "))
__typedS1, __typedOK1 := TypedString(tmp9578)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing } in ")
__typedArg1 := tmp9578
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("missing } in "))
__typedS1, __typedOK1 := TypedString(tmp9578)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("missing } in ")
__typedArg1 := tmp9578
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}


}, 2)

tmp9588 := Call(__e, ns2_1set, symshen_4type_1F, tmp9574)


_ = tmp9588

tmp9589 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V1298 := __self1
_ = V1298
tmp9603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1298)
}
__typedArg0 := Nil
__typedArg1 := V1298
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9603 {
__e.Return(Nil)
return
} else {
tmp9601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1298)
}
__typedArg0 := V1298
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9597 Obj

if True == tmp9601 {
tmp9599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1298)
}
__typedArg0 := V1298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9599)
}
__typedArg0 := tmp9599
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9598 Obj

if True == tmp9600 {
ifres9598 = True


} else {
ifres9598 = False


}

ifres9597 = ifres9598


} else {
ifres9597 = False


}

if True == ifres9597 {
tmp9590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1298)
}
__typedArg0 := V1298
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1298)
}
__typedArg0 := V1298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9591)
}
__typedArg0 := tmp9591
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9593 := Call(__e, PrimFunc(symdeclare), tmp9590, tmp9592)


_ = tmp9593

tmp9594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1298)
}
__typedArg0 := V1298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9594)
}
__typedArg0 := tmp9594
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4assumetypes) == __self {
__self1 = tmp9595
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4assumetypes), tmp9595)
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

tmp9604 := Call(__e, ns2_1set, symshen_4assumetypes, tmp9589)


_ = tmp9604

tmp9605 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V1303 := __self1
_ = V1303
V1304 := __self2
_ = V1304
tmp9616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1304)
}
__typedArg0 := V1304
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9612 Obj

if True == tmp9616 {
tmp9614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1304)
}
__typedArg0 := V1304
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9614)
}
__typedArg0 := tmp9614
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9613 Obj

if True == tmp9615 {
ifres9613 = True


} else {
ifres9613 = False


}

ifres9612 = ifres9613


} else {
ifres9612 = False


}

if True == ifres9612 {
tmp9606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1304)
}
__typedArg0 := V1304
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9607 := Call(__e, PrimFunc(symdestroy), tmp9606)


_ = tmp9607

tmp9608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1304)
}
__typedArg0 := V1304
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9608)
}
__typedArg0 := tmp9608
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4unwind_1types) == __self {
__self1, __self2 = V1303, tmp9609
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4unwind_1types), V1303, tmp9609)
return


} else {
tmp9610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symerror_1to_1string) {
return PrimErrorToString(V1303)
}
__typedArg0 := V1303
return Call(__e, PrimFunc(symerror_1to_1string), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp9610)
}
__typedArg0 := tmp9610
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 2)

tmp9617 := Call(__e, ns2_1set, symshen_4unwind_1types, tmp9605)


_ = tmp9617

tmp9618 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V1307 := __self1
_ = V1307
tmp9664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1307)
}
__typedArg0 := Nil
__typedArg1 := V1307
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9664 {
__e.Return(Nil)
return
} else {
tmp9662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9647 Obj

if True == tmp9662 {
tmp9660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9660)
}
__typedArg0 := tmp9660
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9649 Obj

if True == tmp9661 {
tmp9657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9657)
}
__typedArg0 := tmp9657
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9658)
}
__typedArg0 := tmp9658
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9651 Obj

if True == tmp9659 {
tmp9653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1307)
}
__typedArg0 := V1307
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp9656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp9654, tmp9655)
}
__typedArg0 := tmp9654
__typedArg1 := tmp9655
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9652 Obj

if True == tmp9656 {
ifres9652 = True


} else {
ifres9652 = False


}

ifres9651 = ifres9652


} else {
ifres9651 = False


}

var ifres9650 Obj

if True == ifres9651 {
ifres9650 = True


} else {
ifres9650 = False


}

ifres9649 = ifres9650


} else {
ifres9649 = False


}

var ifres9648 Obj

if True == ifres9649 {
ifres9648 = True


} else {
ifres9648 = False


}

ifres9647 = ifres9648


} else {
ifres9647 = False


}

if True == ifres9647 {
tmp9619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9620)
}
__typedArg0 := tmp9620
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9621)
}
__typedArg0 := tmp9621
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9623 := Call(__e, PrimFunc(symshen_4typecheck), tmp9619, tmp9622)


let__9480 := tmp9623
_ = let__9480

tmp9637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__9480, False)
}
__typedArg0 := let__9480
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9637 {
__e.TailApply(PrimFunc(symshen_4type_1error))
return
} else {
tmp9624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9625 := Call(__e, PrimFunc(symshen_4shen_1_6kl), tmp9624)


tmp9626 := Call(__e, PrimFunc(symeval_1kl), tmp9625)


let__9481 := tmp9626
_ = let__9481

tmp9627 := Call(__e, PrimFunc(symshen_4pretty_1type), let__9480)


tmp9628 := Call(__e, PrimFunc(symshen_4app), tmp9627, MakeString("\n"), symshen_4r)


tmp9630 := Call(__e, PrimFunc(symshen_4app), let__9481, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" : "))
__typedS1, __typedOK1 := TypedString(tmp9628)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" : ")
__typedArg1 := tmp9628
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4s)


tmp9631 := Call(__e, PrimFunc(symstoutput))


tmp9632 := Call(__e, PrimFunc(sympr), tmp9630, tmp9631)


let__9482 := tmp9632
_ = let__9482

tmp9633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9633)
}
__typedArg0 := tmp9633
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9634)
}
__typedArg0 := tmp9634
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4work_1through) == __self {
__self1 = tmp9635
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4work_1through), tmp9635)
return


}


} else {
tmp9645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9645 {
tmp9638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp9640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symA, tmp9640)
}
__typedArg0 := symA
__typedArg1 := tmp9640
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9639, tmp9641)
}
__typedArg0 := tmp9639
__typedArg1 := tmp9641
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp9643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9638, tmp9642)
}
__typedArg0 := tmp9638
__typedArg1 := tmp9642
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4work_1through) == __self {
__self1 = tmp9643
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4work_1through), tmp9643)
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

tmp9665 := Call(__e, ns2_1set, symshen_4work_1through, tmp9618)


_ = tmp9665

tmp9666 := MakeNative(func(__e *ControlFlow) {
V1312 := __e.Get(1)
_ = V1312
tmp9808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9682 Obj

if True == tmp9808 {
tmp9806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9806)
}
__typedArg0 := tmp9806
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9684 Obj

if True == tmp9807 {
tmp9803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9803)
}
__typedArg0 := tmp9803
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp9804)
}
__typedArg0 := symlist
__typedArg1 := tmp9804
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9686 Obj

if True == tmp9805 {
tmp9800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9800)
}
__typedArg0 := tmp9800
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9801)
}
__typedArg0 := tmp9801
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9688 Obj

if True == tmp9802 {
tmp9796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9796)
}
__typedArg0 := tmp9796
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9797)
}
__typedArg0 := tmp9797
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9798)
}
__typedArg0 := Nil
__typedArg1 := tmp9798
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9690 Obj

if True == tmp9799 {
tmp9794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9794)
}
__typedArg0 := tmp9794
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9692 Obj

if True == tmp9795 {
tmp9791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9791)
}
__typedArg0 := tmp9791
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1_1_6, tmp9792)
}
__typedArg0 := sym_1_1_6
__typedArg1 := tmp9792
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9694 Obj

if True == tmp9793 {
tmp9788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9788)
}
__typedArg0 := tmp9788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9789)
}
__typedArg0 := tmp9789
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9696 Obj

if True == tmp9790 {
tmp9784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9784)
}
__typedArg0 := tmp9784
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9785)
}
__typedArg0 := tmp9785
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9786)
}
__typedArg0 := tmp9786
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9698 Obj

if True == tmp9787 {
tmp9779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9779)
}
__typedArg0 := tmp9779
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9780)
}
__typedArg0 := tmp9780
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9781)
}
__typedArg0 := tmp9781
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symstr, tmp9782)
}
__typedArg0 := symstr
__typedArg1 := tmp9782
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9700 Obj

if True == tmp9783 {
tmp9774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9774)
}
__typedArg0 := tmp9774
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9775)
}
__typedArg0 := tmp9775
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9777)
}
__typedArg0 := tmp9777
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9702 Obj

if True == tmp9778 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9769)
}
__typedArg0 := tmp9769
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9770)
}
__typedArg0 := tmp9770
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9772)
}
__typedArg0 := tmp9772
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9704 Obj

if True == tmp9773 {
tmp9761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9762)
}
__typedArg0 := tmp9762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9763)
}
__typedArg0 := tmp9763
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9765)
}
__typedArg0 := tmp9765
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlist, tmp9766)
}
__typedArg0 := symlist
__typedArg1 := tmp9766
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9706 Obj

if True == tmp9767 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9757)
}
__typedArg0 := tmp9757
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9758)
}
__typedArg0 := tmp9758
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9759)
}
__typedArg0 := tmp9759
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9708 Obj

if True == tmp9760 {
tmp9746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9746)
}
__typedArg0 := tmp9746
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9747)
}
__typedArg0 := tmp9747
return Call(__e, PrimFunc(symhd), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9751)
}
__typedArg0 := tmp9751
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9752)
}
__typedArg0 := Nil
__typedArg1 := tmp9752
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9710 Obj

if True == tmp9753 {
tmp9740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9740)
}
__typedArg0 := tmp9740
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9741)
}
__typedArg0 := tmp9741
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9742)
}
__typedArg0 := tmp9742
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9743)
}
__typedArg0 := tmp9743
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9744)
}
__typedArg0 := tmp9744
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9712 Obj

if True == tmp9745 {
tmp9733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9733)
}
__typedArg0 := tmp9733
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9734)
}
__typedArg0 := tmp9734
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9735)
}
__typedArg0 := tmp9735
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9736)
}
__typedArg0 := tmp9736
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9737)
}
__typedArg0 := tmp9737
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9738)
}
__typedArg0 := Nil
__typedArg1 := tmp9738
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9714 Obj

if True == tmp9739 {
tmp9729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9729)
}
__typedArg0 := tmp9729
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9730)
}
__typedArg0 := tmp9730
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9731)
}
__typedArg0 := Nil
__typedArg1 := tmp9731
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9716 Obj

if True == tmp9732 {
tmp9718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9718)
}
__typedArg0 := tmp9718
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9719)
}
__typedArg0 := tmp9719
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9724)
}
__typedArg0 := tmp9724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9725)
}
__typedArg0 := tmp9725
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9726)
}
__typedArg0 := tmp9726
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp9720, tmp9727)
}
__typedArg0 := tmp9720
__typedArg1 := tmp9727
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9717 Obj

if True == tmp9728 {
ifres9717 = True


} else {
ifres9717 = False


}

ifres9716 = ifres9717


} else {
ifres9716 = False


}

var ifres9715 Obj

if True == ifres9716 {
ifres9715 = True


} else {
ifres9715 = False


}

ifres9714 = ifres9715


} else {
ifres9714 = False


}

var ifres9713 Obj

if True == ifres9714 {
ifres9713 = True


} else {
ifres9713 = False


}

ifres9712 = ifres9713


} else {
ifres9712 = False


}

var ifres9711 Obj

if True == ifres9712 {
ifres9711 = True


} else {
ifres9711 = False


}

ifres9710 = ifres9711


} else {
ifres9710 = False


}

var ifres9709 Obj

if True == ifres9710 {
ifres9709 = True


} else {
ifres9709 = False


}

ifres9708 = ifres9709


} else {
ifres9708 = False


}

var ifres9707 Obj

if True == ifres9708 {
ifres9707 = True


} else {
ifres9707 = False


}

ifres9706 = ifres9707


} else {
ifres9706 = False


}

var ifres9705 Obj

if True == ifres9706 {
ifres9705 = True


} else {
ifres9705 = False


}

ifres9704 = ifres9705


} else {
ifres9704 = False


}

var ifres9703 Obj

if True == ifres9704 {
ifres9703 = True


} else {
ifres9703 = False


}

ifres9702 = ifres9703


} else {
ifres9702 = False


}

var ifres9701 Obj

if True == ifres9702 {
ifres9701 = True


} else {
ifres9701 = False


}

ifres9700 = ifres9701


} else {
ifres9700 = False


}

var ifres9699 Obj

if True == ifres9700 {
ifres9699 = True


} else {
ifres9699 = False


}

ifres9698 = ifres9699


} else {
ifres9698 = False


}

var ifres9697 Obj

if True == ifres9698 {
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

if True == ifres9682 {
tmp9667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9667)
}
__typedArg0 := tmp9667
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9668)
}
__typedArg0 := tmp9668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9669)
}
__typedArg0 := tmp9669
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9670)
}
__typedArg0 := tmp9670
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9672)
}
__typedArg0 := tmp9672
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9673)
}
__typedArg0 := tmp9673
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9674)
}
__typedArg0 := tmp9674
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9675)
}
__typedArg0 := tmp9675
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a_a_6, tmp9676)
}
__typedArg0 := sym_a_a_6
__typedArg1 := tmp9676
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp9671, tmp9677)
}
__typedArg0 := tmp9671
__typedArg1 := tmp9677
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1312)
}
__typedArg0 := V1312
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9680 {
tmp9678 := MakeNative(func(__e *ControlFlow) {
Z1313 := __e.Get(1)
_ = Z1313
__e.TailApply(PrimFunc(symshen_4pretty_1type), Z1313)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp9678, V1312)
return


} else {
__e.Return(V1312)
return
}


}


}, 1)

tmp9809 := Call(__e, ns2_1set, symshen_4pretty_1type, tmp9666)


_ = tmp9809

tmp9810 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("type error\n"))
}
__typedArg0 := MakeString("type error\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}, 0)

tmp9811 := Call(__e, ns2_1set, symshen_4type_1error, tmp9810)


_ = tmp9811

tmp9812 := MakeNative(func(__e *ControlFlow) {
V1314 := __e.Get(1)
_ = V1314
tmp9813 := Call(__e, PrimFunc(symshen_4klfile), V1314)


let__9483 := tmp9813
_ = let__9483

tmp9814 := Call(__e, PrimFunc(symread_1file), V1314)


let__9484 := tmp9814
_ = let__9484

tmp9815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symopen) {
return PrimOpenStream(let__9483, symout)
}
__typedArg0 := let__9483
__typedArg1 := symout
return Call(__e, PrimFunc(symopen), __typedArg0, __typedArg1)
})()

let__9485 := tmp9815
_ = let__9485

tmp9816 := MakeNative(func(__e *ControlFlow) {
Z1319 := __e.Get(1)
_ = Z1319
tmp9817 := Call(__e, PrimFunc(symshen_4shen_1_6kl_1h), Z1319)


__e.TailApply(PrimFunc(symshen_4partial), tmp9817)
return


}, 1)

tmp9818 := Call(__e, PrimFunc(symmap), tmp9816, let__9484)


let__9486 := tmp9818
_ = let__9486

tmp9819 := Call(__e, PrimFunc(symshen_4write_1kl), let__9486, let__9485)


let__9487 := tmp9819
_ = let__9487

__e.Return(let__9483)
return


}, 1)

tmp9820 := Call(__e, ns2_1set, symbootstrap, tmp9812)


_ = tmp9820

tmp9821 := MakeNative(func(__e *ControlFlow) {
V1321 := __e.Get(1)
_ = V1321
tmp9844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1321)
}
__typedArg0 := V1321
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9831 Obj

if True == tmp9844 {
tmp9842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1321)
}
__typedArg0 := V1321
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4f_1error, tmp9842)
}
__typedArg0 := symshen_4f_1error
__typedArg1 := tmp9842
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9833 Obj

if True == tmp9843 {
tmp9840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1321)
}
__typedArg0 := V1321
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9840)
}
__typedArg0 := tmp9840
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9835 Obj

if True == tmp9841 {
tmp9837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1321)
}
__typedArg0 := V1321
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9837)
}
__typedArg0 := tmp9837
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9838)
}
__typedArg0 := Nil
__typedArg1 := tmp9838
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9836 Obj

if True == tmp9839 {
ifres9836 = True


} else {
ifres9836 = False


}

ifres9835 = ifres9836


} else {
ifres9835 = False


}

var ifres9834 Obj

if True == ifres9835 {
ifres9834 = True


} else {
ifres9834 = False


}

ifres9833 = ifres9834


} else {
ifres9833 = False


}

var ifres9832 Obj

if True == ifres9833 {
ifres9832 = True


} else {
ifres9832 = False


}

ifres9831 = ifres9832


} else {
ifres9831 = False


}

if True == ifres9831 {
tmp9822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1321)
}
__typedArg0 := V1321
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9822)
}
__typedArg0 := tmp9822
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(tmp9823)
}
__typedArg0 := tmp9823
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp9826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("partial function "))
__typedS1, __typedOK1 := TypedString(tmp9824)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("partial function ")
__typedArg1 := tmp9824
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), Nil)
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("partial function "))
__typedS1, __typedOK1 := TypedString(tmp9824)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("partial function ")
__typedArg1 := tmp9824
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp9826)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp9826
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp9829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1321)
}
__typedArg0 := V1321
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9829 {
tmp9827 := MakeNative(func(__e *ControlFlow) {
Z1322 := __e.Get(1)
_ = Z1322
__e.TailApply(PrimFunc(symshen_4partial), Z1322)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp9827, V1321)
return


} else {
__e.Return(V1321)
return
}


}


}, 1)

tmp9845 := Call(__e, ns2_1set, symshen_4partial, tmp9821)


_ = tmp9845

tmp9846 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V1325 := __self1
_ = V1325
V1326 := __self2
_ = V1326
tmp9860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1325)
}
__typedArg0 := Nil
__typedArg1 := V1325
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9860 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symclose) {
return PrimCloseStream(V1326)
}
__typedArg0 := V1326
return Call(__e, PrimFunc(symclose), __typedArg0)
})())
return
} else {
tmp9858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1325)
}
__typedArg0 := V1325
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9854 Obj

if True == tmp9858 {
tmp9856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1325)
}
__typedArg0 := V1325
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9856)
}
__typedArg0 := tmp9856
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9855 Obj

if True == tmp9857 {
ifres9855 = True


} else {
ifres9855 = False


}

ifres9854 = ifres9855


} else {
ifres9854 = False


}

if True == ifres9854 {
tmp9847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1325)
}
__typedArg0 := V1325
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1325)
}
__typedArg0 := V1325
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9849 := Call(__e, PrimFunc(symshen_4write_1kl_1h), tmp9848, V1326)


_ = tmp9849

if PrimFunc(symshen_4write_1kl) == __self {
__self1, __self2 = tmp9847, V1326
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4write_1kl), tmp9847, V1326)
return


} else {
tmp9852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1325)
}
__typedArg0 := V1325
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp9852 {
tmp9850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1325)
}
__typedArg0 := V1325
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4write_1kl) == __self {
__self1, __self2 = tmp9850, V1326
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4write_1kl), tmp9850, V1326)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4write_1kl)
return
}


}


}


}, 2)

tmp9861 := Call(__e, ns2_1set, symshen_4write_1kl, tmp9846)


_ = tmp9861

tmp9862 := MakeNative(func(__e *ControlFlow) {
V1329 := __e.Get(1)
_ = V1329
V1330 := __e.Get(2)
_ = V1330
tmp9902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1329)
}
__typedArg0 := V1329
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9865 Obj

if True == tmp9902 {
tmp9900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1329)
}
__typedArg0 := V1329
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp9900)
}
__typedArg0 := symdefun
__typedArg1 := tmp9900
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9867 Obj

if True == tmp9901 {
tmp9898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1329)
}
__typedArg0 := V1329
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9898)
}
__typedArg0 := tmp9898
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9869 Obj

if True == tmp9899 {
tmp9895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1329)
}
__typedArg0 := V1329
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9895)
}
__typedArg0 := tmp9895
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail, tmp9896)
}
__typedArg0 := symfail
__typedArg1 := tmp9896
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9871 Obj

if True == tmp9897 {
tmp9892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1329)
}
__typedArg0 := V1329
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9892)
}
__typedArg0 := tmp9892
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9893)
}
__typedArg0 := tmp9893
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9873 Obj

if True == tmp9894 {
tmp9888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1329)
}
__typedArg0 := V1329
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9888)
}
__typedArg0 := tmp9888
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp9889)
}
__typedArg0 := tmp9889
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp9891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9890)
}
__typedArg0 := Nil
__typedArg1 := tmp9890
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9875 Obj

if True == tmp9891 {
tmp9884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1329)
}
__typedArg0 := V1329
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9884)
}
__typedArg0 := tmp9884
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9885)
}
__typedArg0 := tmp9885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp9886)
}
__typedArg0 := tmp9886
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres9877 Obj

if True == tmp9887 {
tmp9879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1329)
}
__typedArg0 := V1329
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9879)
}
__typedArg0 := tmp9879
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9880)
}
__typedArg0 := tmp9880
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp9881)
}
__typedArg0 := tmp9881
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp9883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp9882)
}
__typedArg0 := Nil
__typedArg1 := tmp9882
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres9878 Obj

if True == tmp9883 {
ifres9878 = True


} else {
ifres9878 = False


}

ifres9877 = ifres9878


} else {
ifres9877 = False


}

var ifres9876 Obj

if True == ifres9877 {
ifres9876 = True


} else {
ifres9876 = False


}

ifres9875 = ifres9876


} else {
ifres9875 = False


}

var ifres9874 Obj

if True == ifres9875 {
ifres9874 = True


} else {
ifres9874 = False


}

ifres9873 = ifres9874


} else {
ifres9873 = False


}

var ifres9872 Obj

if True == ifres9873 {
ifres9872 = True


} else {
ifres9872 = False


}

ifres9871 = ifres9872


} else {
ifres9871 = False


}

var ifres9870 Obj

if True == ifres9871 {
ifres9870 = True


} else {
ifres9870 = False


}

ifres9869 = ifres9870


} else {
ifres9869 = False


}

var ifres9868 Obj

if True == ifres9869 {
ifres9868 = True


} else {
ifres9868 = False


}

ifres9867 = ifres9868


} else {
ifres9867 = False


}

var ifres9866 Obj

if True == ifres9867 {
ifres9866 = True


} else {
ifres9866 = False


}

ifres9865 = ifres9866


} else {
ifres9865 = False


}

if True == ifres9865 {
__e.TailApply(PrimFunc(sympr), MakeString("(defun fail () shen.fail!)"), V1330)
return
} else {
tmp9863 := Call(__e, PrimFunc(symshen_4app), V1329, MakeString("\n\n"), symshen_4r)


__e.TailApply(PrimFunc(sympr), tmp9863, V1330)
return


}


}, 2)

tmp9903 := Call(__e, ns2_1set, symshen_4write_1kl_1h, tmp9862)


_ = tmp9903

tmp9904 := MakeNative(func(__e *ControlFlow) {
V1331 := __e.Get(1)
_ = V1331
tmp9913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V1331)
}
__typedArg0 := MakeString("")
__typedArg1 := V1331
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9913 {
__e.Return(MakeString(".kl"))
return
} else {
tmp9911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(".shen"), V1331)
}
__typedArg0 := MakeString(".shen")
__typedArg1 := V1331
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp9911 {
__e.Return(MakeString(".kl"))
return
} else {
tmp9909 := Call(__e, PrimFunc(symshen_4_7string_2), V1331)


if True == tmp9909 {
tmp9905 := Call(__e, PrimFunc(symhdstr), V1331)


tmp9907 := Call(__e, PrimFunc(symshen_4klfile), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V1331)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V1331
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())


__e.TailApply(PrimFunc(sym_8s), tmp9905, tmp9907)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4klfile)
return
}


}


}


}, 1)

__e.TailApply(ns2_1set, symshen_4klfile, tmp9904)
return




}, 0)

