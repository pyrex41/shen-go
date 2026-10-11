package main

import . "github.com/pyrex41/shen-go/kl"

var TrackMain = MakeNative(func(__e *ControlFlow) {
tmp13528 := MakeNative(func(__e *ControlFlow) {
V5864 := __e.Get(1)
_ = V5864
tmp13529 := Call(__e, PrimFunc(symshen_4app), V5864, MakeString(";\n"), symshen_4a)


tmp13530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("partial function "))
__typedS1, __typedOK1 := TypedString(tmp13529)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("partial function ")
__typedArg1 := tmp13529
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp13531 := Call(__e, PrimFunc(symstoutput))


tmp13532 := Call(__e, PrimFunc(sympr), tmp13530, tmp13531)


_ = tmp13532

tmp13541 := Call(__e, PrimFunc(symshen_4tracked_2), V5864)


tmp13542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp13541)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp13541
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres13536 Obj

if True == tmp13542 {
tmp13538 := Call(__e, PrimFunc(symshen_4app), V5864, MakeString("? "), symshen_4a)


tmp13540 := Call(__e, PrimFunc(symy_1or_1n_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("track "))
__typedS1, __typedOK1 := TypedString(tmp13538)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("track ")
__typedArg1 := tmp13538
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())


var ifres13537 Obj

if True == tmp13540 {
ifres13537 = True


} else {
ifres13537 = False


}

ifres13536 = ifres13537


} else {
ifres13536 = False


}

var ifres13533 Obj

if True == ifres13536 {
tmp13534 := Call(__e, PrimFunc(symps), V5864)


tmp13535 := Call(__e, PrimFunc(symshen_4track_1function), tmp13534)


ifres13533 = tmp13535


} else {
ifres13533 = symshen_4ok


}

_ = ifres13533

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("aborted"))
}
__typedArg0 := MakeString("aborted")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

tmp13543 := Call(__e, ns2_1set, symshen_4f_1error, tmp13528)


_ = tmp13543

tmp13544 := MakeNative(func(__e *ControlFlow) {
V5865 := __e.Get(1)
_ = V5865
tmp13545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtracking_d)
}
__typedArg0 := symshen_4_dtracking_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symelement_2), V5865, tmp13545)
return


}, 1)

tmp13546 := Call(__e, ns2_1set, symshen_4tracked_2, tmp13544)


_ = tmp13546

tmp13547 := MakeNative(func(__e *ControlFlow) {
V5866 := __e.Get(1)
_ = V5866
tmp13548 := Call(__e, PrimFunc(symps), V5866)


let__13517 := tmp13548
_ = let__13517

__e.TailApply(PrimFunc(symshen_4track_1function), let__13517)
return


}, 1)

tmp13549 := Call(__e, ns2_1set, symtrack, tmp13547)


_ = tmp13549

tmp13550 := MakeNative(func(__e *ControlFlow) {
V5870 := __e.Get(1)
_ = V5870
tmp13604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5870)
}
__typedArg0 := V5870
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13578 Obj

if True == tmp13604 {
tmp13602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5870)
}
__typedArg0 := V5870
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp13602)
}
__typedArg0 := symdefun
__typedArg1 := tmp13602
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13580 Obj

if True == tmp13603 {
tmp13600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5870)
}
__typedArg0 := V5870
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13600)
}
__typedArg0 := tmp13600
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13582 Obj

if True == tmp13601 {
tmp13597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5870)
}
__typedArg0 := V5870
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13597)
}
__typedArg0 := tmp13597
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13598)
}
__typedArg0 := tmp13598
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13584 Obj

if True == tmp13599 {
tmp13593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5870)
}
__typedArg0 := V5870
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13593)
}
__typedArg0 := tmp13593
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13594)
}
__typedArg0 := tmp13594
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13595)
}
__typedArg0 := tmp13595
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13586 Obj

if True == tmp13596 {
tmp13588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5870)
}
__typedArg0 := V5870
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13588)
}
__typedArg0 := tmp13588
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13589)
}
__typedArg0 := tmp13589
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13590)
}
__typedArg0 := tmp13590
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13591)
}
__typedArg0 := Nil
__typedArg1 := tmp13591
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13587 Obj

if True == tmp13592 {
ifres13587 = True


} else {
ifres13587 = False


}

ifres13586 = ifres13587


} else {
ifres13586 = False


}

var ifres13585 Obj

if True == ifres13586 {
ifres13585 = True


} else {
ifres13585 = False


}

ifres13584 = ifres13585


} else {
ifres13584 = False


}

var ifres13583 Obj

if True == ifres13584 {
ifres13583 = True


} else {
ifres13583 = False


}

ifres13582 = ifres13583


} else {
ifres13582 = False


}

var ifres13581 Obj

if True == ifres13582 {
ifres13581 = True


} else {
ifres13581 = False


}

ifres13580 = ifres13581


} else {
ifres13580 = False


}

var ifres13579 Obj

if True == ifres13580 {
ifres13579 = True


} else {
ifres13579 = False


}

ifres13578 = ifres13579


} else {
ifres13578 = False


}

if True == ifres13578 {
tmp13551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5870)
}
__typedArg0 := V5870
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13551)
}
__typedArg0 := tmp13551
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5870)
}
__typedArg0 := V5870
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13553)
}
__typedArg0 := tmp13553
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13554)
}
__typedArg0 := tmp13554
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5870)
}
__typedArg0 := V5870
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13556)
}
__typedArg0 := tmp13556
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5870)
}
__typedArg0 := V5870
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13558)
}
__typedArg0 := tmp13558
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13559)
}
__typedArg0 := tmp13559
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5870)
}
__typedArg0 := V5870
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13561)
}
__typedArg0 := tmp13561
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13562)
}
__typedArg0 := tmp13562
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13563)
}
__typedArg0 := tmp13563
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13565 := Call(__e, PrimFunc(symshen_4insert_1tracking_1code), tmp13557, tmp13560, tmp13564)


tmp13566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13565, Nil)
}
__typedArg0 := tmp13565
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13555, tmp13566)
}
__typedArg0 := tmp13555
__typedArg1 := tmp13566
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13552, tmp13567)
}
__typedArg0 := tmp13552
__typedArg1 := tmp13567
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp13568)
}
__typedArg0 := symdefun
__typedArg1 := tmp13568
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__13518 := tmp13569
_ = let__13518

tmp13570 := Call(__e, PrimFunc(symeval_1kl), let__13518)


let__13519 := tmp13570
_ = let__13519

tmp13571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5870)
}
__typedArg0 := V5870
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13571)
}
__typedArg0 := tmp13571
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtracking_d)
}
__typedArg0 := symshen_4_dtracking_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp13574 := Call(__e, PrimFunc(symadjoin), tmp13572, tmp13573)


tmp13575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dtracking_d, tmp13574)
}
__typedArg0 := symshen_4_dtracking_d
__typedArg1 := tmp13574
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

let__13520 := tmp13575
_ = let__13520

tmp13576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5870)
}
__typedArg0 := V5870
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13576)
}
__typedArg0 := tmp13576
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.track-function"))
}
__typedArg0 := MakeString("implementation error in shen.track-function")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp13605 := Call(__e, ns2_1set, symshen_4track_1function, tmp13550)


_ = tmp13605

tmp13606 := MakeNative(func(__e *ControlFlow) {
V5874 := __e.Get(1)
_ = V5874
V5875 := __e.Get(2)
_ = V5875
V5876 := __e.Get(3)
_ = V5876
tmp13607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, Nil)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp13607)
}
__typedArg0 := symvalue
__typedArg1 := tmp13607
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), Nil)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13608, tmp13609)
}
__typedArg0 := tmp13608
__typedArg1 := tmp13609
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp13610)
}
__typedArg0 := sym_7
__typedArg1 := tmp13610
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13611, Nil)
}
__typedArg0 := tmp13611
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, tmp13612)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := tmp13612
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp13613)
}
__typedArg0 := symset
__typedArg1 := tmp13613
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, Nil)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp13615)
}
__typedArg0 := symvalue
__typedArg1 := tmp13615
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13617 := Call(__e, PrimFunc(symshen_4prolog_1track), V5876, V5875)


tmp13618 := Call(__e, PrimFunc(symshen_4cons_1form), tmp13617)


tmp13619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13618, Nil)
}
__typedArg0 := tmp13618
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5874, tmp13619)
}
__typedArg0 := V5874
__typedArg1 := tmp13619
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13616, tmp13620)
}
__typedArg0 := tmp13616
__typedArg1 := tmp13620
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4input_1track, tmp13621)
}
__typedArg0 := symshen_4input_1track
__typedArg1 := tmp13621
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4terpri_1or_1read_1char, Nil)
}
__typedArg0 := symshen_4terpri_1or_1read_1char
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, Nil)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp13624)
}
__typedArg0 := symvalue
__typedArg1 := tmp13624
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5874, tmp13626)
}
__typedArg0 := V5874
__typedArg1 := tmp13626
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13625, tmp13627)
}
__typedArg0 := tmp13625
__typedArg1 := tmp13627
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4output_1track, tmp13628)
}
__typedArg0 := symshen_4output_1track
__typedArg1 := tmp13628
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, Nil)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvalue, tmp13630)
}
__typedArg0 := symvalue
__typedArg1 := tmp13630
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), Nil)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13631, tmp13632)
}
__typedArg0 := tmp13631
__typedArg1 := tmp13632
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13633)
}
__typedArg0 := sym_1
__typedArg1 := tmp13633
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13634, Nil)
}
__typedArg0 := tmp13634
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_dcall_d, tmp13635)
}
__typedArg0 := symshen_4_dcall_d
__typedArg1 := tmp13635
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symset, tmp13636)
}
__typedArg0 := symset
__typedArg1 := tmp13636
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4terpri_1or_1read_1char, Nil)
}
__typedArg0 := symshen_4terpri_1or_1read_1char
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13638, tmp13639)
}
__typedArg0 := tmp13638
__typedArg1 := tmp13639
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp13640)
}
__typedArg0 := symdo
__typedArg1 := tmp13640
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13641, Nil)
}
__typedArg0 := tmp13641
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13637, tmp13642)
}
__typedArg0 := tmp13637
__typedArg1 := tmp13642
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp13643)
}
__typedArg0 := symdo
__typedArg1 := tmp13643
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13644, Nil)
}
__typedArg0 := tmp13644
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13629, tmp13645)
}
__typedArg0 := tmp13629
__typedArg1 := tmp13645
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp13646)
}
__typedArg0 := symdo
__typedArg1 := tmp13646
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13647, Nil)
}
__typedArg0 := tmp13647
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5876, tmp13648)
}
__typedArg0 := V5876
__typedArg1 := tmp13648
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, tmp13649)
}
__typedArg0 := symResult
__typedArg1 := tmp13649
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp13650)
}
__typedArg0 := symlet
__typedArg1 := tmp13650
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13651, Nil)
}
__typedArg0 := tmp13651
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13623, tmp13652)
}
__typedArg0 := tmp13623
__typedArg1 := tmp13652
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp13653)
}
__typedArg0 := symdo
__typedArg1 := tmp13653
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13654, Nil)
}
__typedArg0 := tmp13654
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13622, tmp13655)
}
__typedArg0 := tmp13622
__typedArg1 := tmp13655
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp13656)
}
__typedArg0 := symdo
__typedArg1 := tmp13656
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13657, Nil)
}
__typedArg0 := tmp13657
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13614, tmp13658)
}
__typedArg0 := tmp13614
__typedArg1 := tmp13658
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp13659)
}
__typedArg0 := symdo
__typedArg1 := tmp13659
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 3)

tmp13660 := Call(__e, ns2_1set, symshen_4insert_1tracking_1code, tmp13606)


_ = tmp13660

tmp13661 := MakeNative(func(__e *ControlFlow) {
V5877 := __e.Get(1)
_ = V5877
V5878 := __e.Get(2)
_ = V5878
tmp13664 := Call(__e, PrimFunc(symoccurrences), symshen_4incinfs, V5877)


tmp13665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp13664, MakeInteger(0))
}
__typedArg0 := tmp13664
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13665 {
__e.Return(V5878)
return
} else {
tmp13662 := Call(__e, PrimFunc(symshen_4vector_1parameter), V5878)


__e.TailApply(PrimFunc(symshen_4vector_1dereference), V5878, tmp13662)
return


}


}, 2)

tmp13666 := Call(__e, ns2_1set, symshen_4prolog_1track, tmp13661)


_ = tmp13666

tmp13667 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V5881 := __self1
_ = V5881
tmp13696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5881)
}
__typedArg0 := Nil
__typedArg1 := V5881
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13696 {
__e.Return(Nil)
return
} else {
tmp13694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5881)
}
__typedArg0 := V5881
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13672 Obj

if True == tmp13694 {
tmp13692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5881)
}
__typedArg0 := V5881
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13692)
}
__typedArg0 := tmp13692
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13674 Obj

if True == tmp13693 {
tmp13689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5881)
}
__typedArg0 := V5881
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13689)
}
__typedArg0 := tmp13689
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13690)
}
__typedArg0 := tmp13690
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13676 Obj

if True == tmp13691 {
tmp13685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5881)
}
__typedArg0 := V5881
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13685)
}
__typedArg0 := tmp13685
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13686)
}
__typedArg0 := tmp13686
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13687)
}
__typedArg0 := tmp13687
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13678 Obj

if True == tmp13688 {
tmp13680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5881)
}
__typedArg0 := V5881
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13680)
}
__typedArg0 := tmp13680
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13681)
}
__typedArg0 := tmp13681
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13682)
}
__typedArg0 := tmp13682
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13683)
}
__typedArg0 := Nil
__typedArg1 := tmp13683
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13679 Obj

if True == tmp13684 {
ifres13679 = True


} else {
ifres13679 = False


}

ifres13678 = ifres13679


} else {
ifres13678 = False


}

var ifres13677 Obj

if True == ifres13678 {
ifres13677 = True


} else {
ifres13677 = False


}

ifres13676 = ifres13677


} else {
ifres13676 = False


}

var ifres13675 Obj

if True == ifres13676 {
ifres13675 = True


} else {
ifres13675 = False


}

ifres13674 = ifres13675


} else {
ifres13674 = False


}

var ifres13673 Obj

if True == ifres13674 {
ifres13673 = True


} else {
ifres13673 = False


}

ifres13672 = ifres13673


} else {
ifres13672 = False


}

if True == ifres13672 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5881)
}
__typedArg0 := V5881
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp13670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5881)
}
__typedArg0 := V5881
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13670 {
tmp13668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5881)
}
__typedArg0 := V5881
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4vector_1parameter) == __self {
__self1 = tmp13668
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4vector_1parameter), tmp13668)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4vector_1parameter)
return
}


}


}


}, 1)

tmp13697 := Call(__e, ns2_1set, symshen_4vector_1parameter, tmp13667)


_ = tmp13697

tmp13698 := MakeNative(func(__e *ControlFlow) {
V5884 := __e.Get(1)
_ = V5884
V5885 := __e.Get(2)
_ = V5885
tmp13732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5885)
}
__typedArg0 := Nil
__typedArg1 := V5885
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13732 {
__e.Return(V5884)
return
} else {
tmp13730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5884)
}
__typedArg0 := V5884
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13708 Obj

if True == tmp13730 {
tmp13728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5884)
}
__typedArg0 := V5884
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13728)
}
__typedArg0 := tmp13728
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13710 Obj

if True == tmp13729 {
tmp13725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5884)
}
__typedArg0 := V5884
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13725)
}
__typedArg0 := tmp13725
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13726)
}
__typedArg0 := tmp13726
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13712 Obj

if True == tmp13727 {
tmp13721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5884)
}
__typedArg0 := V5884
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13721)
}
__typedArg0 := tmp13721
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13722)
}
__typedArg0 := tmp13722
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13723)
}
__typedArg0 := tmp13723
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13714 Obj

if True == tmp13724 {
tmp13716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5884)
}
__typedArg0 := V5884
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13716)
}
__typedArg0 := tmp13716
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13717)
}
__typedArg0 := tmp13717
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13718)
}
__typedArg0 := tmp13718
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13719)
}
__typedArg0 := Nil
__typedArg1 := tmp13719
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13715 Obj

if True == tmp13720 {
ifres13715 = True


} else {
ifres13715 = False


}

ifres13714 = ifres13715


} else {
ifres13714 = False


}

var ifres13713 Obj

if True == ifres13714 {
ifres13713 = True


} else {
ifres13713 = False


}

ifres13712 = ifres13713


} else {
ifres13712 = False


}

var ifres13711 Obj

if True == ifres13712 {
ifres13711 = True


} else {
ifres13711 = False


}

ifres13710 = ifres13711


} else {
ifres13710 = False


}

var ifres13709 Obj

if True == ifres13710 {
ifres13709 = True


} else {
ifres13709 = False


}

ifres13708 = ifres13709


} else {
ifres13708 = False


}

if True == ifres13708 {
__e.Return(V5884)
return
} else {
tmp13706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5884)
}
__typedArg0 := V5884
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13706 {
tmp13699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5884)
}
__typedArg0 := V5884
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5885, Nil)
}
__typedArg0 := V5885
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13699, tmp13700)
}
__typedArg0 := tmp13699
__typedArg1 := tmp13700
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4deref, tmp13701)
}
__typedArg0 := symshen_4deref
__typedArg1 := tmp13701
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5884)
}
__typedArg0 := V5884
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13704 := Call(__e, PrimFunc(symshen_4vector_1dereference), tmp13703, V5885)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13702, tmp13704)
}
__typedArg0 := tmp13702
__typedArg1 := tmp13704
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4vector_1dereference)
return
}


}


}


}, 2)

tmp13733 := Call(__e, ns2_1set, symshen_4vector_1dereference, tmp13698)


_ = tmp13733

tmp13734 := MakeNative(func(__e *ControlFlow) {
V5888 := __e.Get(1)
_ = V5888
tmp13738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V5888)
}
__typedArg0 := sym_7
__typedArg1 := V5888
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13738 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dstep_d, True)
}
__typedArg0 := symshen_4_dstep_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
tmp13736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V5888)
}
__typedArg0 := sym_1
__typedArg1 := V5888
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13736 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dstep_d, False)
}
__typedArg0 := symshen_4_dstep_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("step expects a + or a -.\n"))
}
__typedArg0 := MakeString("step expects a + or a -.\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp13739 := Call(__e, ns2_1set, symstep, tmp13734)


_ = tmp13739

tmp13740 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dstep_d)
}
__typedArg0 := symshen_4_dstep_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp13741 := Call(__e, ns2_1set, symshen_4step_2, tmp13740)


_ = tmp13741

tmp13742 := MakeNative(func(__e *ControlFlow) {
V5891 := __e.Get(1)
_ = V5891
tmp13746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V5891)
}
__typedArg0 := sym_7
__typedArg1 := V5891
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13746 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspy_d, True)
}
__typedArg0 := symshen_4_dspy_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
tmp13744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V5891)
}
__typedArg0 := sym_1
__typedArg1 := V5891
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13744 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dspy_d, False)
}
__typedArg0 := symshen_4_dspy_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("spy expects a + or a -.\n"))
}
__typedArg0 := MakeString("spy expects a + or a -.\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp13747 := Call(__e, ns2_1set, symspy, tmp13742)


_ = tmp13747

tmp13748 := MakeNative(func(__e *ControlFlow) {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dspy_d)
}
__typedArg0 := symshen_4_dspy_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})())
return
}, 0)

tmp13749 := Call(__e, ns2_1set, symshen_4spy_2, tmp13748)


_ = tmp13749

tmp13750 := MakeNative(func(__e *ControlFlow) {
tmp13754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dstep_d)
}
__typedArg0 := symshen_4_dstep_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

if True == tmp13754 {
tmp13751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dstinput_d)
}
__typedArg0 := sym_dstinput_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp13752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symread_1byte) {
return PrimReadByte(tmp13751)
}
__typedArg0 := tmp13751
return Call(__e, PrimFunc(symread_1byte), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4check_1byte), tmp13752)
return


} else {
__e.TailApply(PrimFunc(symnl), MakeInteger(1))
return
}


}, 0)

tmp13755 := Call(__e, ns2_1set, symshen_4terpri_1or_1read_1char, tmp13750)


_ = tmp13755

tmp13756 := MakeNative(func(__e *ControlFlow) {
V5894 := __e.Get(1)
_ = V5894
tmp13758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(94), V5894)
}
__typedArg0 := MakeInteger(94)
__typedArg1 := V5894
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13758 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("aborted"))
}
__typedArg0 := MakeString("aborted")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
} else {
__e.Return(True)
return
}


}, 1)

tmp13759 := Call(__e, ns2_1set, symshen_4check_1byte, tmp13756)


_ = tmp13759

tmp13760 := MakeNative(func(__e *ControlFlow) {
V5895 := __e.Get(1)
_ = V5895
V5896 := __e.Get(2)
_ = V5896
V5897 := __e.Get(3)
_ = V5897
tmp13761 := Call(__e, PrimFunc(symshen_4spaces), V5895)


tmp13762 := Call(__e, PrimFunc(symshen_4spaces), V5895)


tmp13763 := Call(__e, PrimFunc(symshen_4app), tmp13762, MakeString(""), symshen_4a)


tmp13765 := Call(__e, PrimFunc(symshen_4app), V5896, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" \n"))
__typedS1, __typedOK1 := TypedString(tmp13763)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" \n")
__typedArg1 := tmp13763
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp13767 := Call(__e, PrimFunc(symshen_4app), V5895, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("> Inputs to "))
__typedS1, __typedOK1 := TypedString(tmp13765)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("> Inputs to ")
__typedArg1 := tmp13765
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp13769 := Call(__e, PrimFunc(symshen_4app), tmp13761, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("<"))
__typedS1, __typedOK1 := TypedString(tmp13767)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("<")
__typedArg1 := tmp13767
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp13770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\n"))
__typedS1, __typedOK1 := TypedString(tmp13769)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\n")
__typedArg1 := tmp13769
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp13771 := Call(__e, PrimFunc(symstoutput))


tmp13772 := Call(__e, PrimFunc(sympr), tmp13770, tmp13771)


_ = tmp13772

__e.TailApply(PrimFunc(symshen_4recursively_1print), V5897)
return


}, 3)

tmp13773 := Call(__e, ns2_1set, symshen_4input_1track, tmp13760)


_ = tmp13773

tmp13774 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V5900 := __self1
_ = V5900
tmp13784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5900)
}
__typedArg0 := Nil
__typedArg1 := V5900
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13784 {
tmp13775 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString(" ==>"), tmp13775)
return


} else {
tmp13782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5900)
}
__typedArg0 := V5900
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13782 {
tmp13776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5900)
}
__typedArg0 := V5900
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13777 := Call(__e, PrimFunc(symprint), tmp13776)


_ = tmp13777

tmp13778 := Call(__e, PrimFunc(symstoutput))


tmp13779 := Call(__e, PrimFunc(sympr), MakeString(", "), tmp13778)


_ = tmp13779

tmp13780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5900)
}
__typedArg0 := V5900
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4recursively_1print) == __self {
__self1 = tmp13780
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4recursively_1print), tmp13780)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.recursively-print"))
}
__typedArg0 := MakeString("implementation error in shen.recursively-print")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp13785 := Call(__e, ns2_1set, symshen_4recursively_1print, tmp13774)


_ = tmp13785

tmp13786 := MakeNative(func(__e *ControlFlow) {
V5901 := __e.Get(1)
_ = V5901
tmp13790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V5901)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V5901
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13790 {
__e.Return(MakeString(""))
return
} else {
tmp13788 := Call(__e, PrimFunc(symshen_4spaces), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V5901)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V5901
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" "))
__typedS1, __typedOK1 := TypedString(tmp13788)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" ")
__typedArg1 := tmp13788
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp13791 := Call(__e, ns2_1set, symshen_4spaces, tmp13786)


_ = tmp13791

tmp13792 := MakeNative(func(__e *ControlFlow) {
V5902 := __e.Get(1)
_ = V5902
V5903 := __e.Get(2)
_ = V5903
V5904 := __e.Get(3)
_ = V5904
tmp13793 := Call(__e, PrimFunc(symshen_4spaces), V5902)


tmp13794 := Call(__e, PrimFunc(symshen_4spaces), V5902)


tmp13795 := Call(__e, PrimFunc(symshen_4app), V5904, MakeString(""), symshen_4s)


tmp13797 := Call(__e, PrimFunc(symshen_4app), tmp13794, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("==> "))
__typedS1, __typedOK1 := TypedString(tmp13795)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("==> ")
__typedArg1 := tmp13795
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp13799 := Call(__e, PrimFunc(symshen_4app), V5903, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" \n"))
__typedS1, __typedOK1 := TypedString(tmp13797)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" \n")
__typedArg1 := tmp13797
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp13801 := Call(__e, PrimFunc(symshen_4app), V5902, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("> Output from "))
__typedS1, __typedOK1 := TypedString(tmp13799)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("> Output from ")
__typedArg1 := tmp13799
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp13803 := Call(__e, PrimFunc(symshen_4app), tmp13793, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("<"))
__typedS1, __typedOK1 := TypedString(tmp13801)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("<")
__typedArg1 := tmp13801
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})(), symshen_4a)


tmp13804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("\n"))
__typedS1, __typedOK1 := TypedString(tmp13803)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("\n")
__typedArg1 := tmp13803
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp13805 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp13804, tmp13805)
return


}, 3)

tmp13806 := Call(__e, ns2_1set, symshen_4output_1track, tmp13792)


_ = tmp13806

tmp13807 := MakeNative(func(__e *ControlFlow) {
V5905 := __e.Get(1)
_ = V5905
tmp13808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dtracking_d)
}
__typedArg0 := symshen_4_dtracking_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp13809 := Call(__e, PrimFunc(symremove), V5905, tmp13808)


tmp13810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dtracking_d, tmp13809)
}
__typedArg0 := symshen_4_dtracking_d
__typedArg1 := tmp13809
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp13810

tmp13811 := MakeNative(func(__e *ControlFlow) {
tmp13812 := Call(__e, PrimFunc(symps), V5905)


__e.TailApply(PrimFunc(symeval), tmp13812)
return


}, 0)

tmp13813 := MakeNative(func(__e *ControlFlow) {
Z5906 := __e.Get(1)
_ = Z5906
__e.Return(V5905)
return
}, 1)

tmp13814 := Call(__e, try_1catch, tmp13811, tmp13813)


_ = tmp13814

__e.Return(V5905)
return


}, 1)

tmp13815 := Call(__e, ns2_1set, symuntrack, tmp13807)


_ = tmp13815

tmp13816 := MakeNative(func(__e *ControlFlow) {
V5907 := __e.Get(1)
_ = V5907
V5908 := __e.Get(2)
_ = V5908
__e.TailApply(PrimFunc(symshen_4remove_1h), V5907, V5908, Nil)
return
}, 2)

tmp13817 := Call(__e, ns2_1set, symremove, tmp13816)


_ = tmp13817

tmp13818 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__selftop:
V5918 := __self1
_ = V5918
V5919 := __self2
_ = V5919
V5920 := __self3
_ = V5920
tmp13833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V5919)
}
__typedArg0 := Nil
__typedArg1 := V5919
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13833 {
__e.TailApply(PrimFunc(symreverse), V5920)
return
} else {
tmp13831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5919)
}
__typedArg0 := V5919
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13827 Obj

if True == tmp13831 {
tmp13829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5919)
}
__typedArg0 := V5919
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V5918, tmp13829)
}
__typedArg0 := V5918
__typedArg1 := tmp13829
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13828 Obj

if True == tmp13830 {
ifres13828 = True


} else {
ifres13828 = False


}

ifres13827 = ifres13828


} else {
ifres13827 = False


}

if True == ifres13827 {
tmp13819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5919)
}
__typedArg0 := V5919
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5919)
}
__typedArg0 := V5919
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4remove_1h) == __self {
__self1, __self2, __self3 = tmp13819, tmp13820, V5920
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4remove_1h), tmp13819, tmp13820, V5920)
return


} else {
tmp13825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5919)
}
__typedArg0 := V5919
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13825 {
tmp13821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5919)
}
__typedArg0 := V5919
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5919)
}
__typedArg0 := V5919
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13822, V5920)
}
__typedArg0 := tmp13822
__typedArg1 := V5920
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4remove_1h) == __self {
__self1, __self2, __self3 = V5918, tmp13821, tmp13823
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4remove_1h), V5918, tmp13821, tmp13823)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.remove-h"))
}
__typedArg0 := MakeString("implementation error in shen.remove-h")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 3)

tmp13834 := Call(__e, ns2_1set, symshen_4remove_1h, tmp13818)


_ = tmp13834

tmp13835 := MakeNative(func(__e *ControlFlow) {
V5921 := __e.Get(1)
_ = V5921
tmp13836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dprofiled_d)
}
__typedArg0 := symshen_4_dprofiled_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp13837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5921, tmp13836)
}
__typedArg0 := V5921
__typedArg1 := tmp13836
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dprofiled_d, tmp13837)
}
__typedArg0 := symshen_4_dprofiled_d
__typedArg1 := tmp13837
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp13838

tmp13839 := Call(__e, PrimFunc(symps), V5921)


__e.TailApply(PrimFunc(symshen_4profile_1help), tmp13839)
return


}, 1)

tmp13840 := Call(__e, ns2_1set, symprofile, tmp13835)


_ = tmp13840

tmp13841 := MakeNative(func(__e *ControlFlow) {
V5924 := __e.Get(1)
_ = V5924
tmp13906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V5924)
}
__typedArg0 := V5924
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13880 Obj

if True == tmp13906 {
tmp13904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V5924)
}
__typedArg0 := V5924
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp13904)
}
__typedArg0 := symdefun
__typedArg1 := tmp13904
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13882 Obj

if True == tmp13905 {
tmp13902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5924)
}
__typedArg0 := V5924
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13902)
}
__typedArg0 := tmp13902
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13884 Obj

if True == tmp13903 {
tmp13899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5924)
}
__typedArg0 := V5924
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13899)
}
__typedArg0 := tmp13899
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13900)
}
__typedArg0 := tmp13900
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13886 Obj

if True == tmp13901 {
tmp13895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5924)
}
__typedArg0 := V5924
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13895)
}
__typedArg0 := tmp13895
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13896)
}
__typedArg0 := tmp13896
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13897)
}
__typedArg0 := tmp13897
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13888 Obj

if True == tmp13898 {
tmp13890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5924)
}
__typedArg0 := V5924
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13890)
}
__typedArg0 := tmp13890
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13891)
}
__typedArg0 := tmp13891
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13892)
}
__typedArg0 := tmp13892
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13893)
}
__typedArg0 := Nil
__typedArg1 := tmp13893
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13889 Obj

if True == tmp13894 {
ifres13889 = True


} else {
ifres13889 = False


}

ifres13888 = ifres13889


} else {
ifres13888 = False


}

var ifres13887 Obj

if True == ifres13888 {
ifres13887 = True


} else {
ifres13887 = False


}

ifres13886 = ifres13887


} else {
ifres13886 = False


}

var ifres13885 Obj

if True == ifres13886 {
ifres13885 = True


} else {
ifres13885 = False


}

ifres13884 = ifres13885


} else {
ifres13884 = False


}

var ifres13883 Obj

if True == ifres13884 {
ifres13883 = True


} else {
ifres13883 = False


}

ifres13882 = ifres13883


} else {
ifres13882 = False


}

var ifres13881 Obj

if True == ifres13882 {
ifres13881 = True


} else {
ifres13881 = False


}

ifres13880 = ifres13881


} else {
ifres13880 = False


}

if True == ifres13880 {
tmp13842 := Call(__e, PrimFunc(symgensym), symshen_4f)


let__13521 := tmp13842
_ = let__13521

tmp13843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5924)
}
__typedArg0 := V5924
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13843)
}
__typedArg0 := tmp13843
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5924)
}
__typedArg0 := V5924
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13845)
}
__typedArg0 := tmp13845
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13846)
}
__typedArg0 := tmp13846
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5924)
}
__typedArg0 := V5924
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13848)
}
__typedArg0 := tmp13848
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5924)
}
__typedArg0 := V5924
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13850)
}
__typedArg0 := tmp13850
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13851)
}
__typedArg0 := tmp13851
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5924)
}
__typedArg0 := V5924
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13853)
}
__typedArg0 := tmp13853
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13854)
}
__typedArg0 := tmp13854
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__13521, tmp13855)
}
__typedArg0 := let__13521
__typedArg1 := tmp13855
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13857 := Call(__e, PrimFunc(symshen_4profile_1func), tmp13849, tmp13852, tmp13856)


tmp13858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13857, Nil)
}
__typedArg0 := tmp13857
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13847, tmp13858)
}
__typedArg0 := tmp13847
__typedArg1 := tmp13858
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13844, tmp13859)
}
__typedArg0 := tmp13844
__typedArg1 := tmp13859
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp13860)
}
__typedArg0 := symdefun
__typedArg1 := tmp13860
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__13522 := tmp13861
_ = let__13522

tmp13862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5924)
}
__typedArg0 := V5924
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13862)
}
__typedArg0 := tmp13862
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13863)
}
__typedArg0 := tmp13863
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5924)
}
__typedArg0 := V5924
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13865)
}
__typedArg0 := tmp13865
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5924)
}
__typedArg0 := V5924
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13867)
}
__typedArg0 := tmp13867
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13868)
}
__typedArg0 := tmp13868
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13869)
}
__typedArg0 := tmp13869
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13871 := Call(__e, PrimFunc(symsubst), let__13521, tmp13866, tmp13870)


tmp13872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13871, Nil)
}
__typedArg0 := tmp13871
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13864, tmp13872)
}
__typedArg0 := tmp13864
__typedArg1 := tmp13872
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__13521, tmp13873)
}
__typedArg0 := let__13521
__typedArg1 := tmp13873
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp13874)
}
__typedArg0 := symdefun
__typedArg1 := tmp13874
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__13523 := tmp13875
_ = let__13523

tmp13876 := Call(__e, PrimFunc(symeval_1kl), let__13522)


let__13524 := tmp13876
_ = let__13524

tmp13877 := Call(__e, PrimFunc(symeval_1kl), let__13523)


let__13525 := tmp13877
_ = let__13525

tmp13878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V5924)
}
__typedArg0 := V5924
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13878)
}
__typedArg0 := tmp13878
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("Cannot profile.\n"))
}
__typedArg0 := MakeString("Cannot profile.\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp13907 := Call(__e, ns2_1set, symshen_4profile_1help, tmp13841)


_ = tmp13907

tmp13908 := MakeNative(func(__e *ControlFlow) {
V5930 := __e.Get(1)
_ = V5930
tmp13909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dprofiled_d)
}
__typedArg0 := symshen_4_dprofiled_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp13910 := Call(__e, PrimFunc(symremove), V5930, tmp13909)


tmp13911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dprofiled_d, tmp13910)
}
__typedArg0 := symshen_4_dprofiled_d
__typedArg1 := tmp13910
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp13911

tmp13912 := MakeNative(func(__e *ControlFlow) {
tmp13913 := Call(__e, PrimFunc(symps), V5930)


__e.TailApply(PrimFunc(symeval), tmp13913)
return


}, 0)

tmp13914 := MakeNative(func(__e *ControlFlow) {
Z5931 := __e.Get(1)
_ = Z5931
__e.Return(V5930)
return
}, 1)

__e.TailApply(try_1catch, tmp13912, tmp13914)
return


}, 1)

tmp13915 := Call(__e, ns2_1set, symunprofile, tmp13908)


_ = tmp13915

tmp13916 := MakeNative(func(__e *ControlFlow) {
V5932 := __e.Get(1)
_ = V5932
tmp13917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dprofiled_d)
}
__typedArg0 := symshen_4_dprofiled_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symelement_2), V5932, tmp13917)
return


}, 1)

tmp13918 := Call(__e, ns2_1set, symshen_4profiled_2, tmp13916)


_ = tmp13918

tmp13919 := MakeNative(func(__e *ControlFlow) {
V5933 := __e.Get(1)
_ = V5933
V5934 := __e.Get(2)
_ = V5934
V5935 := __e.Get(3)
_ = V5935
tmp13920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrun, Nil)
}
__typedArg0 := symrun
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp13920)
}
__typedArg0 := symget_1time
__typedArg1 := tmp13920
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symrun, Nil)
}
__typedArg0 := symrun
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget_1time, tmp13922)
}
__typedArg0 := symget_1time
__typedArg1 := tmp13922
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symStart, Nil)
}
__typedArg0 := symStart
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13923, tmp13924)
}
__typedArg0 := tmp13923
__typedArg1 := tmp13924
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13925)
}
__typedArg0 := sym_1
__typedArg1 := tmp13925
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5933, Nil)
}
__typedArg0 := V5933
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4get_1profile, tmp13927)
}
__typedArg0 := symshen_4get_1profile
__typedArg1 := tmp13927
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symFinish, Nil)
}
__typedArg0 := symFinish
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13928, tmp13929)
}
__typedArg0 := tmp13928
__typedArg1 := tmp13929
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp13930)
}
__typedArg0 := sym_7
__typedArg1 := tmp13930
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13931, Nil)
}
__typedArg0 := tmp13931
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5933, tmp13932)
}
__typedArg0 := V5933
__typedArg1 := tmp13932
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4put_1profile, tmp13933)
}
__typedArg0 := symshen_4put_1profile
__typedArg1 := tmp13933
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, Nil)
}
__typedArg0 := symResult
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13934, tmp13935)
}
__typedArg0 := tmp13934
__typedArg1 := tmp13935
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symRecord, tmp13936)
}
__typedArg0 := symRecord
__typedArg1 := tmp13936
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp13937)
}
__typedArg0 := symlet
__typedArg1 := tmp13937
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13938, Nil)
}
__typedArg0 := tmp13938
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13926, tmp13939)
}
__typedArg0 := tmp13926
__typedArg1 := tmp13939
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symFinish, tmp13940)
}
__typedArg0 := symFinish
__typedArg1 := tmp13940
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp13941)
}
__typedArg0 := symlet
__typedArg1 := tmp13941
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13942, Nil)
}
__typedArg0 := tmp13942
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V5935, tmp13943)
}
__typedArg0 := V5935
__typedArg1 := tmp13943
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symResult, tmp13944)
}
__typedArg0 := symResult
__typedArg1 := tmp13944
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp13945)
}
__typedArg0 := symlet
__typedArg1 := tmp13945
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13946, Nil)
}
__typedArg0 := tmp13946
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13921, tmp13947)
}
__typedArg0 := tmp13921
__typedArg1 := tmp13947
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symStart, tmp13948)
}
__typedArg0 := symStart
__typedArg1 := tmp13948
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp13949)
}
__typedArg0 := symlet
__typedArg1 := tmp13949
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 3)

tmp13950 := Call(__e, ns2_1set, symshen_4profile_1func, tmp13919)


_ = tmp13950

tmp13951 := MakeNative(func(__e *ControlFlow) {
V5936 := __e.Get(1)
_ = V5936
tmp13952 := Call(__e, PrimFunc(symshen_4get_1profile), V5936)


let__13526 := tmp13952
_ = let__13526

tmp13953 := Call(__e, PrimFunc(symshen_4put_1profile), V5936, MakeInteger(0))


let__13527 := tmp13953
_ = let__13527

__e.TailApply(PrimFunc(sym_8p), V5936, let__13526)
return


}, 1)

tmp13954 := Call(__e, ns2_1set, symprofile_1results, tmp13951)


_ = tmp13954

tmp13955 := MakeNative(func(__e *ControlFlow) {
V5939 := __e.Get(1)
_ = V5939
tmp13956 := MakeNative(func(__e *ControlFlow) {
tmp13957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symget), V5939, symprofile, tmp13957)
return


}, 0)

tmp13958 := MakeNative(func(__e *ControlFlow) {
Z5940 := __e.Get(1)
_ = Z5940
__e.Return(MakeInteger(0))
return
}, 1)

__e.TailApply(try_1catch, tmp13956, tmp13958)
return


}, 1)

tmp13959 := Call(__e, ns2_1set, symshen_4get_1profile, tmp13955)


_ = tmp13959

tmp13960 := MakeNative(func(__e *ControlFlow) {
V5941 := __e.Get(1)
_ = V5941
V5942 := __e.Get(2)
_ = V5942
tmp13961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V5941, symprofile, V5942, tmp13961)
return


}, 2)

__e.TailApply(ns2_1set, symshen_4put_1profile, tmp13960)
return




}, 0)

