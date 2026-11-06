package main

import . "github.com/pyrex41/shen-go/kl"

var LoadMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp9468 := MakeNative1(func(__e *ControlFlow, V1262 Obj) Obj {
tmp9469 := GValue(__e, symshen_4_dtc_d)

let__9450 := tmp9469
_ = let__9450

tmp9470 := PrimGuard1(__e, symget_1time, PrimGetTime, symrun)

let__9452 := tmp9470
_ = let__9452

tmp9471 := Call1(__e, SymFn(symread_1file), V1262)


tmp9472 := Call2(__e, SymFn(symshen_4load_1help), let__9450, tmp9471)


let__9453 := tmp9472
_ = let__9453

tmp9473 := PrimGuard1(__e, symget_1time, PrimGetTime, symrun)

let__9454 := tmp9473
_ = let__9454

let__9455 := GSub(__e, let__9454, let__9452)
_ = let__9455

tmp9475 := PrimGuard1(__e, symstr, PrimStr, let__9455)

tmp9477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\nrun time: "))
__typedS1, __typedOK1 := TypedString(tmp9475)
__typedS2, __typedOK2 := TypedString(MakeString(" secs\n"))
if __typedOK0 && __typedOK1 && __typedOK2 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + (__typedS1 + __typedS2)))
}}
__typedArg0 := MakeString("\nrun time: ")
__typedArg1 := PrimGuard2(__e, symcn, PrimStringConcat, tmp9475, MakeString(" secs\n"))
return Call2(__e, SymFn(symcn), __typedArg0, __typedArg1)
})()

tmp9478 := Call0(__e, SymFn(symstoutput))


tmp9479 := Call2(__e, SymFn(sympr), tmp9477, tmp9478)


let__9456 := tmp9479
_ = let__9456

let__9451 := let__9453
_ = let__9451

var ifres9480 Obj

if True == let__9450 {
tmp9481 := Call0(__e, SymFn(syminferences))


tmp9482 := Call3(__e, SymFn(symshen_4app), tmp9481, MakeString(" inferences\n"), symshen_4a)


tmp9483 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("\ntypechecked in "), tmp9482)

tmp9484 := Call0(__e, SymFn(symstoutput))


tmp9485 := Call2(__e, SymFn(sympr), tmp9483, tmp9484)


ifres9480 = tmp9485


} else {
ifres9480 = symshen_4skip


}

let__9457 := ifres9480
_ = let__9457

return symloaded


panic("unreachable")
})

tmp9486 := Call2(__e, ns2_1set, symload, tmp9468)


_ = tmp9486

tmp9487 := MakeNative2(func(__e *ControlFlow, V1273, V1274 Obj) Obj {
tmp9489 := GEq(__e, False, V1273)

if True == tmp9489 {
return TailCall1(__e, SymFn(symshen_4eval_1and_1print), V1274)
} else {
return TailCall1(__e, SymFn(symshen_4check_1eval_1and_1print), V1274)
}


panic("unreachable")
})

tmp9490 := Call2(__e, ns2_1set, symshen_4load_1help, tmp9487)


_ = tmp9490

tmp9491 := MakeNative1(func(__e *ControlFlow, V1275 Obj) Obj {
tmp9492 := MakeNative1(func(__e *ControlFlow, Z1276 Obj) Obj {
tmp9493 := Call1(__e, SymFn(symshen_4shen_1_6kl), Z1276)


tmp9494 := Call1(__e, SymFn(symeval_1kl), tmp9493)


tmp9495 := Call3(__e, SymFn(symshen_4app), tmp9494, MakeString("\n"), symshen_4s)


tmp9496 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp9495, tmp9496)


panic("unreachable")
})

return TailCall2(__e, SymFn(symshen_4for_1each), tmp9492, V1275)


panic("unreachable")
})

tmp9497 := Call2(__e, ns2_1set, symshen_4eval_1and_1print, tmp9491)


_ = tmp9497

tmp9498 := MakeNative1(func(__e *ControlFlow, V1277 Obj) Obj {
tmp9499 := MakeNative1(func(__e *ControlFlow, Z1279 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4typetable), Z1279)
panic("unreachable")
})

tmp9500 := Call2(__e, SymFn(symmapcan), tmp9499, V1277)


let__9458 := tmp9500
_ = let__9458

tmp9501 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, SymFn(symshen_4assumetypes), let__9458)
panic("unreachable")
})

tmp9502 := MakeNative1(func(__e *ControlFlow, Z1281 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4unwind_1types), Z1281, let__9458)
panic("unreachable")
})

tmp9503 := Call2(__e, try_1catch, tmp9501, tmp9502)


let__9459 := tmp9503
_ = let__9459

tmp9504 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, SymFn(symshen_4work_1through), V1277)
panic("unreachable")
})

tmp9505 := MakeNative1(func(__e *ControlFlow, Z1282 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4unwind_1types), Z1282, let__9458)
panic("unreachable")
})

return TailCall2(__e, try_1catch, tmp9504, tmp9505)


panic("unreachable")
})

tmp9506 := Call2(__e, ns2_1set, symshen_4check_1eval_1and_1print, tmp9498)


_ = tmp9506

tmp9507 := MakeNative1(func(__e *ControlFlow, V1287 Obj) Obj {
tmp9552 := GConsP(__e, V1287)

var ifres9533 Obj

if True == tmp9552 {
tmp9550 := GHd(__e, V1287)

tmp9551 := GEq(__e, symdefine, tmp9550)

var ifres9535 Obj

if True == tmp9551 {
tmp9548 := GTl(__e, V1287)

tmp9549 := GConsP(__e, tmp9548)

var ifres9537 Obj

if True == tmp9549 {
tmp9545 := GTl(__e, V1287)

tmp9546 := GTl(__e, tmp9545)

tmp9547 := GConsP(__e, tmp9546)

var ifres9539 Obj

if True == tmp9547 {
tmp9541 := GTl(__e, V1287)

tmp9542 := GTl(__e, tmp9541)

tmp9543 := GHd(__e, tmp9542)

tmp9544 := GEq(__e, sym_i, tmp9543)

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
tmp9508 := GTl(__e, V1287)

tmp9509 := GHd(__e, tmp9508)

tmp9510 := GTl(__e, V1287)

tmp9511 := GHd(__e, tmp9510)

tmp9512 := GTl(__e, V1287)

tmp9513 := GTl(__e, tmp9512)

tmp9514 := GTl(__e, tmp9513)

tmp9515 := Call2(__e, SymFn(symshen_4type_1F), tmp9511, tmp9514)


tmp9516 := Call1(__e, SymFn(symshen_4rectify_1type), tmp9515)


tmp9517 := GCons(__e, tmp9516, Nil)

return GCons(__e, tmp9509, tmp9517)


} else {
tmp9531 := GConsP(__e, V1287)

var ifres9523 Obj

if True == tmp9531 {
tmp9529 := GHd(__e, V1287)

tmp9530 := GEq(__e, symdefine, tmp9529)

var ifres9525 Obj

if True == tmp9530 {
tmp9527 := GTl(__e, V1287)

tmp9528 := GConsP(__e, tmp9527)

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
tmp9518 := GTl(__e, V1287)

tmp9519 := GHd(__e, tmp9518)

tmp9520 := Call3(__e, SymFn(symshen_4app), tmp9519, MakeString("\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("missing { in "), tmp9520))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("missing { in "), tmp9520)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


} else {
return Nil
}


}


panic("unreachable")
})

tmp9553 := Call2(__e, ns2_1set, symshen_4typetable, tmp9507)


_ = tmp9553

tmp9554 := MakeNative2(func(__e *ControlFlow, V1294, V1295 Obj) Obj {
tmp9567 := GConsP(__e, V1295)

var ifres9563 Obj

if True == tmp9567 {
tmp9565 := GHd(__e, V1295)

tmp9566 := GEq(__e, sym_j, tmp9565)

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
return Nil
} else {
tmp9561 := GConsP(__e, V1295)

if True == tmp9561 {
tmp9555 := GHd(__e, V1295)

tmp9556 := GTl(__e, V1295)

tmp9557 := Call2(__e, SymFn(symshen_4type_1F), V1294, tmp9556)


return GCons(__e, tmp9555, tmp9557)


} else {
tmp9558 := Call3(__e, SymFn(symshen_4app), V1294, MakeString("\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("missing } in "), tmp9558))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("missing } in "), tmp9558)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


}


panic("unreachable")
})

tmp9568 := Call2(__e, ns2_1set, symshen_4type_1F, tmp9554)


_ = tmp9568

tmp9569 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V1298 := __self1
_ = V1298
tmp9583 := GEq(__e, Nil, V1298)

if True == tmp9583 {
return Nil
} else {
tmp9581 := GConsP(__e, V1298)

var ifres9577 Obj

if True == tmp9581 {
tmp9579 := GTl(__e, V1298)

tmp9580 := GConsP(__e, tmp9579)

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
tmp9570 := GHd(__e, V1298)

tmp9571 := GTl(__e, V1298)

tmp9572 := GHd(__e, tmp9571)

tmp9573 := Call2(__e, SymFn(symdeclare), tmp9570, tmp9572)


_ = tmp9573

tmp9574 := GTl(__e, V1298)

tmp9575 := GTl(__e, tmp9574)

if SymFn(symshen_4assumetypes) == __self {
__self1 = tmp9575
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4assumetypes), tmp9575)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.assumetype"))
}


}


panic("unreachable")
})
return __self
})()

tmp9584 := Call2(__e, ns2_1set, symshen_4assumetypes, tmp9569)


_ = tmp9584

tmp9585 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V1303 := __self1
_ = V1303
V1304 := __self2
_ = V1304
tmp9596 := GConsP(__e, V1304)

var ifres9592 Obj

if True == tmp9596 {
tmp9594 := GTl(__e, V1304)

tmp9595 := GConsP(__e, tmp9594)

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
tmp9586 := GHd(__e, V1304)

tmp9587 := Call1(__e, SymFn(symdestroy), tmp9586)


_ = tmp9587

tmp9588 := GTl(__e, V1304)

tmp9589 := GTl(__e, tmp9588)

if SymFn(symshen_4unwind_1types) == __self {
__self1, __self2 = V1303, tmp9589
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4unwind_1types), V1303, tmp9589)


} else {
tmp9590 := PrimGuard1(__e, symerror_1to_1string, PrimErrorToString, V1303)

return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp9590)


}


panic("unreachable")
})
return __self
})()

tmp9597 := Call2(__e, ns2_1set, symshen_4unwind_1types, tmp9585)


_ = tmp9597

tmp9598 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V1307 := __self1
_ = V1307
tmp9644 := GEq(__e, Nil, V1307)

if True == tmp9644 {
return Nil
} else {
tmp9642 := GConsP(__e, V1307)

var ifres9627 Obj

if True == tmp9642 {
tmp9640 := GTl(__e, V1307)

tmp9641 := GConsP(__e, tmp9640)

var ifres9629 Obj

if True == tmp9641 {
tmp9637 := GTl(__e, V1307)

tmp9638 := GTl(__e, tmp9637)

tmp9639 := GConsP(__e, tmp9638)

var ifres9631 Obj

if True == tmp9639 {
tmp9633 := GTl(__e, V1307)

tmp9634 := GHd(__e, tmp9633)

tmp9635 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp9636 := GEq(__e, tmp9634, tmp9635)

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
tmp9599 := GHd(__e, V1307)

tmp9600 := GTl(__e, V1307)

tmp9601 := GTl(__e, tmp9600)

tmp9602 := GHd(__e, tmp9601)

tmp9603 := Call2(__e, SymFn(symshen_4typecheck), tmp9599, tmp9602)


let__9460 := tmp9603
_ = let__9460

tmp9617 := GEq(__e, let__9460, False)

if True == tmp9617 {
return TailCall0(__e, SymFn(symshen_4type_1error))
} else {
tmp9604 := GHd(__e, V1307)

tmp9605 := Call1(__e, SymFn(symshen_4shen_1_6kl), tmp9604)


tmp9606 := Call1(__e, SymFn(symeval_1kl), tmp9605)


let__9461 := tmp9606
_ = let__9461

tmp9607 := Call1(__e, SymFn(symshen_4pretty_1type), let__9460)


tmp9608 := Call3(__e, SymFn(symshen_4app), tmp9607, MakeString("\n"), symshen_4r)


tmp9610 := Call3(__e, SymFn(symshen_4app), let__9461, PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" : "), tmp9608), symshen_4s)


tmp9611 := Call0(__e, SymFn(symstoutput))


tmp9612 := Call2(__e, SymFn(sympr), tmp9610, tmp9611)


let__9462 := tmp9612
_ = let__9462

tmp9613 := GTl(__e, V1307)

tmp9614 := GTl(__e, tmp9613)

tmp9615 := GTl(__e, tmp9614)

if SymFn(symshen_4work_1through) == __self {
__self1 = tmp9615
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4work_1through), tmp9615)


}


} else {
tmp9625 := GConsP(__e, V1307)

if True == tmp9625 {
tmp9618 := GHd(__e, V1307)

tmp9619 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp9620 := GTl(__e, V1307)

tmp9621 := GCons(__e, symA, tmp9620)

tmp9622 := GCons(__e, tmp9619, tmp9621)

tmp9623 := GCons(__e, tmp9618, tmp9622)

if SymFn(symshen_4work_1through) == __self {
__self1 = tmp9623
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4work_1through), tmp9623)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.work-through"))
}


}


}


panic("unreachable")
})
return __self
})()

tmp9645 := Call2(__e, ns2_1set, symshen_4work_1through, tmp9598)


_ = tmp9645

tmp9646 := MakeNative1(func(__e *ControlFlow, V1312 Obj) Obj {
tmp9788 := GConsP(__e, V1312)

var ifres9662 Obj

if True == tmp9788 {
tmp9786 := GHd(__e, V1312)

tmp9787 := GConsP(__e, tmp9786)

var ifres9664 Obj

if True == tmp9787 {
tmp9783 := GHd(__e, V1312)

tmp9784 := GHd(__e, tmp9783)

tmp9785 := GEq(__e, symlist, tmp9784)

var ifres9666 Obj

if True == tmp9785 {
tmp9780 := GHd(__e, V1312)

tmp9781 := GTl(__e, tmp9780)

tmp9782 := GConsP(__e, tmp9781)

var ifres9668 Obj

if True == tmp9782 {
tmp9776 := GHd(__e, V1312)

tmp9777 := GTl(__e, tmp9776)

tmp9778 := GTl(__e, tmp9777)

tmp9779 := GEq(__e, Nil, tmp9778)

var ifres9670 Obj

if True == tmp9779 {
tmp9774 := GTl(__e, V1312)

tmp9775 := GConsP(__e, tmp9774)

var ifres9672 Obj

if True == tmp9775 {
tmp9771 := GTl(__e, V1312)

tmp9772 := GHd(__e, tmp9771)

tmp9773 := GEq(__e, sym_1_1_6, tmp9772)

var ifres9674 Obj

if True == tmp9773 {
tmp9768 := GTl(__e, V1312)

tmp9769 := GTl(__e, tmp9768)

tmp9770 := GConsP(__e, tmp9769)

var ifres9676 Obj

if True == tmp9770 {
tmp9764 := GTl(__e, V1312)

tmp9765 := GTl(__e, tmp9764)

tmp9766 := GHd(__e, tmp9765)

tmp9767 := GConsP(__e, tmp9766)

var ifres9678 Obj

if True == tmp9767 {
tmp9759 := GTl(__e, V1312)

tmp9760 := GTl(__e, tmp9759)

tmp9761 := GHd(__e, tmp9760)

tmp9762 := GHd(__e, tmp9761)

tmp9763 := GEq(__e, symstr, tmp9762)

var ifres9680 Obj

if True == tmp9763 {
tmp9754 := GTl(__e, V1312)

tmp9755 := GTl(__e, tmp9754)

tmp9756 := GHd(__e, tmp9755)

tmp9757 := GTl(__e, tmp9756)

tmp9758 := GConsP(__e, tmp9757)

var ifres9682 Obj

if True == tmp9758 {
tmp9748 := GTl(__e, V1312)

tmp9749 := GTl(__e, tmp9748)

tmp9750 := GHd(__e, tmp9749)

tmp9751 := GTl(__e, tmp9750)

tmp9752 := GHd(__e, tmp9751)

tmp9753 := GConsP(__e, tmp9752)

var ifres9684 Obj

if True == tmp9753 {
tmp9741 := GTl(__e, V1312)

tmp9742 := GTl(__e, tmp9741)

tmp9743 := GHd(__e, tmp9742)

tmp9744 := GTl(__e, tmp9743)

tmp9745 := GHd(__e, tmp9744)

tmp9746 := GHd(__e, tmp9745)

tmp9747 := GEq(__e, symlist, tmp9746)

var ifres9686 Obj

if True == tmp9747 {
tmp9734 := GTl(__e, V1312)

tmp9735 := GTl(__e, tmp9734)

tmp9736 := GHd(__e, tmp9735)

tmp9737 := GTl(__e, tmp9736)

tmp9738 := GHd(__e, tmp9737)

tmp9739 := GTl(__e, tmp9738)

tmp9740 := GConsP(__e, tmp9739)

var ifres9688 Obj

if True == tmp9740 {
tmp9726 := GTl(__e, V1312)

tmp9727 := GTl(__e, tmp9726)

tmp9728 := GHd(__e, tmp9727)

tmp9729 := GTl(__e, tmp9728)

tmp9730 := GHd(__e, tmp9729)

tmp9731 := GTl(__e, tmp9730)

tmp9732 := GTl(__e, tmp9731)

tmp9733 := GEq(__e, Nil, tmp9732)

var ifres9690 Obj

if True == tmp9733 {
tmp9720 := GTl(__e, V1312)

tmp9721 := GTl(__e, tmp9720)

tmp9722 := GHd(__e, tmp9721)

tmp9723 := GTl(__e, tmp9722)

tmp9724 := GTl(__e, tmp9723)

tmp9725 := GConsP(__e, tmp9724)

var ifres9692 Obj

if True == tmp9725 {
tmp9713 := GTl(__e, V1312)

tmp9714 := GTl(__e, tmp9713)

tmp9715 := GHd(__e, tmp9714)

tmp9716 := GTl(__e, tmp9715)

tmp9717 := GTl(__e, tmp9716)

tmp9718 := GTl(__e, tmp9717)

tmp9719 := GEq(__e, Nil, tmp9718)

var ifres9694 Obj

if True == tmp9719 {
tmp9709 := GTl(__e, V1312)

tmp9710 := GTl(__e, tmp9709)

tmp9711 := GTl(__e, tmp9710)

tmp9712 := GEq(__e, Nil, tmp9711)

var ifres9696 Obj

if True == tmp9712 {
tmp9698 := GHd(__e, V1312)

tmp9699 := GTl(__e, tmp9698)

tmp9700 := GHd(__e, tmp9699)

tmp9701 := GTl(__e, V1312)

tmp9702 := GTl(__e, tmp9701)

tmp9703 := GHd(__e, tmp9702)

tmp9704 := GTl(__e, tmp9703)

tmp9705 := GHd(__e, tmp9704)

tmp9706 := GTl(__e, tmp9705)

tmp9707 := GHd(__e, tmp9706)

tmp9708 := GEq(__e, tmp9700, tmp9707)

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
tmp9647 := GTl(__e, V1312)

tmp9648 := GTl(__e, tmp9647)

tmp9649 := GHd(__e, tmp9648)

tmp9650 := GTl(__e, tmp9649)

tmp9651 := GHd(__e, tmp9650)

tmp9652 := GTl(__e, V1312)

tmp9653 := GTl(__e, tmp9652)

tmp9654 := GHd(__e, tmp9653)

tmp9655 := GTl(__e, tmp9654)

tmp9656 := GTl(__e, tmp9655)

tmp9657 := GCons(__e, sym_a_a_6, tmp9656)

return GCons(__e, tmp9651, tmp9657)


} else {
tmp9660 := GConsP(__e, V1312)

if True == tmp9660 {
tmp9658 := MakeNative1(func(__e *ControlFlow, Z1313 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4pretty_1type), Z1313)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp9658, V1312)


} else {
return V1312
}


}


panic("unreachable")
})

tmp9789 := Call2(__e, ns2_1set, symshen_4pretty_1type, tmp9646)


_ = tmp9789

tmp9790 := MakeNative0(func(__e *ControlFlow) Obj {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("type error\n"))
panic("unreachable")
})

tmp9791 := Call2(__e, ns2_1set, symshen_4type_1error, tmp9790)


_ = tmp9791

tmp9792 := MakeNative1(func(__e *ControlFlow, V1314 Obj) Obj {
tmp9793 := Call1(__e, SymFn(symshen_4klfile), V1314)


let__9463 := tmp9793
_ = let__9463

tmp9794 := Call1(__e, SymFn(symread_1file), V1314)


let__9464 := tmp9794
_ = let__9464

tmp9795 := PrimGuard2(__e, symopen, PrimOpenStream, let__9463, symout)

let__9465 := tmp9795
_ = let__9465

tmp9796 := MakeNative1(func(__e *ControlFlow, Z1319 Obj) Obj {
tmp9797 := Call1(__e, SymFn(symshen_4shen_1_6kl_1h), Z1319)


return TailCall1(__e, SymFn(symshen_4partial), tmp9797)


panic("unreachable")
})

tmp9798 := Call2(__e, SymFn(symmap), tmp9796, let__9464)


let__9466 := tmp9798
_ = let__9466

tmp9799 := Call2(__e, SymFn(symshen_4write_1kl), let__9466, let__9465)


let__9467 := tmp9799
_ = let__9467

return let__9463


panic("unreachable")
})

tmp9800 := Call2(__e, ns2_1set, symbootstrap, tmp9792)


_ = tmp9800

tmp9801 := MakeNative1(func(__e *ControlFlow, V1321 Obj) Obj {
tmp9824 := GConsP(__e, V1321)

var ifres9811 Obj

if True == tmp9824 {
tmp9822 := GHd(__e, V1321)

tmp9823 := GEq(__e, symshen_4f_1error, tmp9822)

var ifres9813 Obj

if True == tmp9823 {
tmp9820 := GTl(__e, V1321)

tmp9821 := GConsP(__e, tmp9820)

var ifres9815 Obj

if True == tmp9821 {
tmp9817 := GTl(__e, V1321)

tmp9818 := GTl(__e, tmp9817)

tmp9819 := GEq(__e, Nil, tmp9818)

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
tmp9802 := GTl(__e, V1321)

tmp9803 := GHd(__e, tmp9802)

tmp9804 := PrimGuard1(__e, symstr, PrimStr, tmp9803)

tmp9806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("partial function "), tmp9804), Nil)
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("partial function "), tmp9804)
__typedArg1 := Nil
return Call2(__e, SymFn(symcons), __typedArg0, __typedArg1)
})()

return GCons(__e, symsimple_1error, tmp9806)


} else {
tmp9809 := GConsP(__e, V1321)

if True == tmp9809 {
tmp9807 := MakeNative1(func(__e *ControlFlow, Z1322 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4partial), Z1322)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp9807, V1321)


} else {
return V1321
}


}


panic("unreachable")
})

tmp9825 := Call2(__e, ns2_1set, symshen_4partial, tmp9801)


_ = tmp9825

tmp9826 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V1325 := __self1
_ = V1325
V1326 := __self2
_ = V1326
tmp9840 := GEq(__e, Nil, V1325)

if True == tmp9840 {
return PrimGuard1(__e, symclose, PrimCloseStream, V1326)
} else {
tmp9838 := GConsP(__e, V1325)

var ifres9834 Obj

if True == tmp9838 {
tmp9836 := GHd(__e, V1325)

tmp9837 := GConsP(__e, tmp9836)

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
tmp9827 := GTl(__e, V1325)

tmp9828 := GHd(__e, V1325)

tmp9829 := Call2(__e, SymFn(symshen_4write_1kl_1h), tmp9828, V1326)


_ = tmp9829

if SymFn(symshen_4write_1kl) == __self {
__self1, __self2 = tmp9827, V1326
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4write_1kl), tmp9827, V1326)


} else {
tmp9832 := GConsP(__e, V1325)

if True == tmp9832 {
tmp9830 := GTl(__e, V1325)

if SymFn(symshen_4write_1kl) == __self {
__self1, __self2 = tmp9830, V1326
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4write_1kl), tmp9830, V1326)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4write_1kl)
}


}


}


panic("unreachable")
})
return __self
})()

tmp9841 := Call2(__e, ns2_1set, symshen_4write_1kl, tmp9826)


_ = tmp9841

tmp9842 := MakeNative2(func(__e *ControlFlow, V1329, V1330 Obj) Obj {
tmp9882 := GConsP(__e, V1329)

var ifres9845 Obj

if True == tmp9882 {
tmp9880 := GHd(__e, V1329)

tmp9881 := GEq(__e, symdefun, tmp9880)

var ifres9847 Obj

if True == tmp9881 {
tmp9878 := GTl(__e, V1329)

tmp9879 := GConsP(__e, tmp9878)

var ifres9849 Obj

if True == tmp9879 {
tmp9875 := GTl(__e, V1329)

tmp9876 := GHd(__e, tmp9875)

tmp9877 := GEq(__e, symfail, tmp9876)

var ifres9851 Obj

if True == tmp9877 {
tmp9872 := GTl(__e, V1329)

tmp9873 := GTl(__e, tmp9872)

tmp9874 := GConsP(__e, tmp9873)

var ifres9853 Obj

if True == tmp9874 {
tmp9868 := GTl(__e, V1329)

tmp9869 := GTl(__e, tmp9868)

tmp9870 := GHd(__e, tmp9869)

tmp9871 := GEq(__e, Nil, tmp9870)

var ifres9855 Obj

if True == tmp9871 {
tmp9864 := GTl(__e, V1329)

tmp9865 := GTl(__e, tmp9864)

tmp9866 := GTl(__e, tmp9865)

tmp9867 := GConsP(__e, tmp9866)

var ifres9857 Obj

if True == tmp9867 {
tmp9859 := GTl(__e, V1329)

tmp9860 := GTl(__e, tmp9859)

tmp9861 := GTl(__e, tmp9860)

tmp9862 := GTl(__e, tmp9861)

tmp9863 := GEq(__e, Nil, tmp9862)

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
return TailCall2(__e, SymFn(sympr), MakeString("(defun fail () shen.fail!)"), V1330)
} else {
tmp9843 := Call3(__e, SymFn(symshen_4app), V1329, MakeString("\n\n"), symshen_4r)


return TailCall2(__e, SymFn(sympr), tmp9843, V1330)


}


panic("unreachable")
})

tmp9883 := Call2(__e, ns2_1set, symshen_4write_1kl_1h, tmp9842)


_ = tmp9883

tmp9884 := MakeNative1(func(__e *ControlFlow, V1331 Obj) Obj {
tmp9893 := GEq(__e, MakeString(""), V1331)

if True == tmp9893 {
return MakeString(".kl")
} else {
tmp9891 := GEq(__e, MakeString(".shen"), V1331)

if True == tmp9891 {
return MakeString(".kl")
} else {
tmp9889 := Call1(__e, SymFn(symshen_4_7string_2), V1331)


if True == tmp9889 {
tmp9885 := Call1(__e, SymFn(symhdstr), V1331)


tmp9887 := Call1(__e, SymFn(symshen_4klfile), PrimGuard1(__e, symtlstr, PrimTailString, V1331))


return TailCall2(__e, SymFn(sym_8s), tmp9885, tmp9887)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4klfile)
}


}


}


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4klfile, tmp9884)




panic("unreachable")
})

