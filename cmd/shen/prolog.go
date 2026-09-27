package main

import . "github.com/pyrex41/shen-go/kl"

var PrologMain = MakeNative(func(__e *ControlFlow) {
tmp10580 := MakeNative(func(__e *ControlFlow) {
V1245 := __e.Get(1)
_ = V1245
__e.TailApply(PrimFunc(symshen_4assert_d), V1245, symshen_4top)
return
}, 1)

tmp10581 := Call(__e, ns2_1set, symasserta, tmp10580)


_ = tmp10581

tmp10582 := MakeNative(func(__e *ControlFlow) {
V1246 := __e.Get(1)
_ = V1246
__e.TailApply(PrimFunc(symshen_4assert_d), V1246, symshen_4bottom)
return
}, 1)

tmp10583 := Call(__e, ns2_1set, symassertz, tmp10582)


_ = tmp10583

tmp10584 := MakeNative(func(__e *ControlFlow) {
V1247 := __e.Get(1)
_ = V1247
V1248 := __e.Get(2)
_ = V1248
tmp10611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1247)
}
__typedArg0 := V1247
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10602 Obj

if True == tmp10611 {
tmp10609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1247)
}
__typedArg0 := V1247
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10609)
}
__typedArg0 := tmp10609
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10604 Obj

if True == tmp10610 {
tmp10606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1247)
}
__typedArg0 := V1247
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10606)
}
__typedArg0 := tmp10606
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_5_1_1, tmp10607)
}
__typedArg0 := sym_5_1_1
__typedArg1 := tmp10607
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10605 Obj

if True == tmp10608 {
ifres10605 = True


} else {
ifres10605 = False


}

ifres10604 = ifres10605


} else {
ifres10604 = False


}

var ifres10603 Obj

if True == ifres10604 {
ifres10603 = True


} else {
ifres10603 = False


}

ifres10602 = ifres10603


} else {
ifres10602 = False


}

if True == ifres10602 {
tmp10585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1247)
}
__typedArg0 := V1247
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10586 := Call(__e, PrimFunc(symshen_4predicate), tmp10585)


let__10362 := tmp10586
_ = let__10362

tmp10587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1247)
}
__typedArg0 := V1247
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10588 := Call(__e, PrimFunc(symshen_4terms), tmp10587)


let__10363 := tmp10588
_ = let__10363

tmp10589 := Call(__e, PrimFunc(symlength), let__10363)


let__10364 := tmp10589
_ = let__10364

tmp10590 := Call(__e, PrimFunc(symshen_4parameters), let__10364)


let__10365 := tmp10590
_ = let__10365

tmp10591 := Call(__e, PrimFunc(symarity), let__10362)


let__10366 := tmp10591
_ = let__10366

tmp10597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10366, MakeInteger(-1))
}
__typedArg0 := let__10366
__typedArg1 := MakeInteger(-1)
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10592 Obj

if True == tmp10597 {
tmp10593 := Call(__e, PrimFunc(symshen_4create_1skeleton), let__10362, let__10365)


tmp10594 := Call(__e, PrimFunc(symeval), tmp10593)


_ = tmp10594

tmp10595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp10596 := Call(__e, PrimFunc(symput), let__10362, symshen_4dynamic, Nil, tmp10595)


ifres10592 = tmp10596


} else {
ifres10592 = symshen_4skip


}

let__10367 := ifres10592
_ = let__10367

tmp10598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1247)
}
__typedArg0 := V1247
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10598)
}
__typedArg0 := tmp10598
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10600 := Call(__e, PrimFunc(symshen_4insert_1info), let__10362, let__10363, tmp10599, V1247, V1248)


let__10368 := tmp10600
_ = let__10368

__e.Return(let__10362)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.assert*"))
}
__typedArg0 := MakeString("partial function shen.assert*")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 2)

tmp10612 := Call(__e, ns2_1set, symshen_4assert_d, tmp10584)


_ = tmp10612

tmp10613 := MakeNative(func(__e *ControlFlow) {
V1258 := __e.Get(1)
_ = V1258
tmp10615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1258)
}
__typedArg0 := V1258
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10615 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1258)
}
__typedArg0 := V1258
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
__e.Return(V1258)
return
}


}, 1)

tmp10616 := Call(__e, ns2_1set, symshen_4predicate, tmp10613)


_ = tmp10616

tmp10617 := MakeNative(func(__e *ControlFlow) {
V1263 := __e.Get(1)
_ = V1263
tmp10619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1263)
}
__typedArg0 := V1263
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10619 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1263)
}
__typedArg0 := V1263
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return
} else {
__e.Return(Nil)
return
}


}, 1)

tmp10620 := Call(__e, ns2_1set, symshen_4terms, tmp10617)


_ = tmp10620

tmp10621 := MakeNative(func(__e *ControlFlow) {
V1264 := __e.Get(1)
_ = V1264
V1265 := __e.Get(2)
_ = V1265
tmp10622 := Call(__e, PrimFunc(symshen_4dynamic_1default), V1264, V1265)


tmp10623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1264, tmp10622)
}
__typedArg0 := V1264
__typedArg1 := tmp10622
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp10623)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp10623
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp10624 := Call(__e, ns2_1set, symshen_4create_1skeleton, tmp10621)


_ = tmp10624

tmp10625 := MakeNative(func(__e *ControlFlow) {
V1266 := __e.Get(1)
_ = V1266
V1267 := __e.Get(2)
_ = V1267
tmp10626 := Call(__e, PrimFunc(symshen_4cons_1form), V1267)


tmp10627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4dynamic, Nil)
}
__typedArg0 := symshen_4dynamic
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1266, tmp10627)
}
__typedArg0 := V1266
__typedArg1 := tmp10627
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symget, tmp10628)
}
__typedArg0 := symget
__typedArg1 := tmp10628
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10629, Nil)
}
__typedArg0 := tmp10629
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10626, tmp10630)
}
__typedArg0 := tmp10626
__typedArg1 := tmp10630
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4call_1dynamic, tmp10631)
}
__typedArg0 := symshen_4call_1dynamic
__typedArg1 := tmp10631
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp10634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10633, Nil)
}
__typedArg0 := tmp10633
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10632, tmp10634)
}
__typedArg0 := tmp10632
__typedArg1 := tmp10634
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, tmp10635)
}
__typedArg0 := sym_5_1_1
__typedArg1 := tmp10635
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symappend), V1267, tmp10636)
return


}, 2)

tmp10637 := Call(__e, ns2_1set, symshen_4dynamic_1default, tmp10625)


_ = tmp10637

tmp10638 := MakeNative(func(__e *ControlFlow) {
V1268 := __e.Get(1)
_ = V1268
V1269 := __e.Get(2)
_ = V1269
V1270 := __e.Get(3)
_ = V1270
V1271 := __e.Get(4)
_ = V1271
V1272 := __e.Get(5)
_ = V1272
tmp10639 := Call(__e, PrimFunc(symgensym), symshen_4g)


let__10369 := tmp10639
_ = let__10369

tmp10640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10369, Nil)
}
__typedArg0 := let__10369
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp10640)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp10640
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, V1270)
}
__typedArg0 := sym_5_1_1
__typedArg1 := V1270
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10643 := Call(__e, PrimFunc(symappend), V1269, tmp10642)


tmp10644 := Call(__e, PrimFunc(symappend), tmp10641, tmp10643)


tmp10645 := Call(__e, PrimFunc(symeval), tmp10644)


let__10370 := tmp10645
_ = let__10370

tmp10646 := Call(__e, PrimFunc(symfn), let__10369)


tmp10647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10369, V1271)
}
__typedArg0 := let__10369
__typedArg1 := V1271
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10646, tmp10647)
}
__typedArg0 := tmp10646
__typedArg1 := tmp10647
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__10371 := tmp10648
_ = let__10371

tmp10649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp10650 := Call(__e, PrimFunc(symget), V1268, symshen_4dynamic, tmp10649)


let__10372 := tmp10650
_ = let__10372

tmp10655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1272, symshen_4top)
}
__typedArg0 := V1272
__typedArg1 := symshen_4top
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10651 Obj

if True == tmp10655 {
tmp10652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10371, let__10372)
}
__typedArg0 := let__10371
__typedArg1 := let__10372
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

ifres10651 = tmp10652


} else {
tmp10653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10371, Nil)
}
__typedArg0 := let__10371
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10654 := Call(__e, PrimFunc(symappend), let__10372, tmp10653)


ifres10651 = tmp10654


}

let__10373 := ifres10651
_ = let__10373

tmp10656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V1268, symshen_4dynamic, let__10373, tmp10656)
return


}, 5)

tmp10657 := Call(__e, ns2_1set, symshen_4insert_1info, tmp10638)


_ = tmp10657

tmp10658 := MakeNative(func(__e *ControlFlow) {
tmp10659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dnames_d)
}
__typedArg0 := symshen_4_dnames_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__10374 := tmp10659
_ = let__10374

tmp10665 := Call(__e, PrimFunc(symempty_2), let__10374)


var ifres10660 Obj

if True == tmp10665 {
tmp10661 := Call(__e, PrimFunc(symgensym), symshen_4g)


ifres10660 = tmp10661


} else {
tmp10662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__10374)
}
__typedArg0 := let__10374
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dnames_d, tmp10662)
}
__typedArg0 := symshen_4_dnames_d
__typedArg1 := tmp10662
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp10663

tmp10664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__10374)
}
__typedArg0 := let__10374
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

ifres10660 = tmp10664


}

let__10375 := ifres10660
_ = let__10375

__e.Return(let__10375)
return


}, 0)

tmp10666 := Call(__e, ns2_1set, symshen_4newname, tmp10658)


_ = tmp10666

tmp10667 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
__self6 := __e.Get(6)
__selftop:
V1280 := __self1
_ = V1280
V1281 := __self2
_ = V1281
V1282 := __self3
_ = V1282
V1283 := __self4
_ = V1283
V1284 := __self5
_ = V1284
V1285 := __self6
_ = V1285
tmp10679 := Call(__e, PrimFunc(symshen_4unlocked_2), V1283)


var ifres10668 Obj

if True == tmp10679 {
tmp10669 := Call(__e, PrimFunc(symshen_4lazyderef), V1281, V1282)


let__10377 := tmp10669
_ = let__10377

tmp10678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__10377)
}
__typedArg0 := let__10377
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10670 Obj

if True == tmp10678 {
tmp10671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__10377)
}
__typedArg0 := let__10377
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10672 := Call(__e, PrimFunc(symshen_4lazyderef), tmp10671, V1282)


let__10378 := tmp10672
_ = let__10378

tmp10677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__10378)
}
__typedArg0 := let__10378
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10673 Obj

if True == tmp10677 {
tmp10674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(let__10378)
}
__typedArg0 := let__10378
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

let__10379 := tmp10674
_ = let__10379

tmp10675 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp10675

tmp10676 := Call(__e, PrimFunc(symshen_4callrec), let__10379, V1280, V1282, V1283, V1284, V1285)


ifres10673 = tmp10676


} else {
ifres10673 = False


}

ifres10670 = ifres10673


} else {
ifres10670 = False


}

ifres10668 = ifres10670


} else {
ifres10668 = False


}

let__10376 := ifres10668
_ = let__10376

tmp10688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10376, False)
}
__typedArg0 := let__10376
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10688 {
tmp10686 := Call(__e, PrimFunc(symshen_4unlocked_2), V1283)


if True == tmp10686 {
tmp10680 := Call(__e, PrimFunc(symshen_4lazyderef), V1281, V1282)


let__10380 := tmp10680
_ = let__10380

tmp10684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__10380)
}
__typedArg0 := let__10380
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10684 {
tmp10681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(let__10380)
}
__typedArg0 := let__10380
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

let__10381 := tmp10681
_ = let__10381

tmp10682 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp10682

if PrimFunc(symshen_4call_1dynamic) == __self {
__self1, __self2, __self3, __self4, __self5, __self6 = V1280, let__10381, V1282, V1283, V1284, V1285
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4call_1dynamic), V1280, let__10381, V1282, V1283, V1284, V1285)
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
__e.Return(let__10376)
return
}


}, 6)

tmp10689 := Call(__e, ns2_1set, symshen_4call_1dynamic, tmp10667)


_ = tmp10689

tmp10690 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
__self6 := __e.Get(6)
__selftop:
V1292 := __self1
_ = V1292
V1293 := __self2
_ = V1293
V1294 := __self3
_ = V1294
V1295 := __self4
_ = V1295
V1296 := __self5
_ = V1296
V1297 := __self6
_ = V1297
tmp10700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1293)
}
__typedArg0 := Nil
__typedArg1 := V1293
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10700 {
tmp10691 := Call(__e, V1292, V1294)


tmp10692 := Call(__e, tmp10691, V1295)


tmp10693 := Call(__e, tmp10692, V1296)


__e.TailApply(tmp10693, V1297)
return


} else {
tmp10698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1293)
}
__typedArg0 := V1293
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10698 {
tmp10694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1293)
}
__typedArg0 := V1293
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10695 := Call(__e, V1292, tmp10694)


tmp10696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1293)
}
__typedArg0 := V1293
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4callrec) == __self {
__self1, __self2, __self3, __self4, __self5, __self6 = tmp10695, tmp10696, V1294, V1295, V1296, V1297
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4callrec), tmp10695, tmp10696, V1294, V1295, V1296, V1297)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.callrec"))
}
__typedArg0 := MakeString("partial function shen.callrec")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 6)

tmp10701 := Call(__e, ns2_1set, symshen_4callrec, tmp10690)


_ = tmp10701

tmp10702 := MakeNative(func(__e *ControlFlow) {
V1298 := __e.Get(1)
_ = V1298
tmp10719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1298)
}
__typedArg0 := V1298
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10710 Obj

if True == tmp10719 {
tmp10717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1298)
}
__typedArg0 := V1298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10717)
}
__typedArg0 := tmp10717
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10712 Obj

if True == tmp10718 {
tmp10714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1298)
}
__typedArg0 := V1298
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10714)
}
__typedArg0 := tmp10714
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_5_1_1, tmp10715)
}
__typedArg0 := sym_5_1_1
__typedArg1 := tmp10715
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10713 Obj

if True == tmp10716 {
ifres10713 = True


} else {
ifres10713 = False


}

ifres10712 = ifres10713


} else {
ifres10712 = False


}

var ifres10711 Obj

if True == ifres10712 {
ifres10711 = True


} else {
ifres10711 = False


}

ifres10710 = ifres10711


} else {
ifres10710 = False


}

if True == ifres10710 {
tmp10703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1298)
}
__typedArg0 := V1298
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10704 := Call(__e, PrimFunc(symshen_4predicate), tmp10703)


let__10382 := tmp10704
_ = let__10382

tmp10705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp10706 := Call(__e, PrimFunc(symget), let__10382, symshen_4dynamic, tmp10705)


let__10383 := tmp10706
_ = let__10383

tmp10707 := Call(__e, PrimFunc(symshen_4retract_1clause), V1298, let__10383)


tmp10708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), let__10382, symshen_4dynamic, tmp10707, tmp10708)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function retract"))
}
__typedArg0 := MakeString("partial function retract")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp10720 := Call(__e, ns2_1set, symretract, tmp10702)


_ = tmp10720

tmp10721 := MakeNative(func(__e *ControlFlow) {
V1306 := __e.Get(1)
_ = V1306
V1307 := __e.Get(2)
_ = V1307
tmp10751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1307)
}
__typedArg0 := Nil
__typedArg1 := V1307
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10751 {
__e.Return(Nil)
return
} else {
tmp10749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10734 Obj

if True == tmp10749 {
tmp10747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10747)
}
__typedArg0 := tmp10747
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10736 Obj

if True == tmp10748 {
tmp10744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10744)
}
__typedArg0 := tmp10744
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10745)
}
__typedArg0 := tmp10745
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10738 Obj

if True == tmp10746 {
tmp10740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10740)
}
__typedArg0 := tmp10740
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10741)
}
__typedArg0 := tmp10741
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1306, tmp10742)
}
__typedArg0 := V1306
__typedArg1 := tmp10742
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10739 Obj

if True == tmp10743 {
ifres10739 = True


} else {
ifres10739 = False


}

ifres10738 = ifres10739


} else {
ifres10738 = False


}

var ifres10737 Obj

if True == ifres10738 {
ifres10737 = True


} else {
ifres10737 = False


}

ifres10736 = ifres10737


} else {
ifres10736 = False


}

var ifres10735 Obj

if True == ifres10736 {
ifres10735 = True


} else {
ifres10735 = False


}

ifres10734 = ifres10735


} else {
ifres10734 = False


}

if True == ifres10734 {
tmp10722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10722)
}
__typedArg0 := tmp10722
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10723)
}
__typedArg0 := tmp10723
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dnames_d)
}
__typedArg0 := symshen_4_dnames_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp10726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10724, tmp10725)
}
__typedArg0 := tmp10724
__typedArg1 := tmp10725
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dnames_d, tmp10726)
}
__typedArg0 := symshen_4_dnames_d
__typedArg1 := tmp10726
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp10727

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symtl), __typedArg0)
})())
return


} else {
tmp10732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp10732 {
tmp10728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1307)
}
__typedArg0 := V1307
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10730 := Call(__e, PrimFunc(symshen_4retract_1clause), V1306, tmp10729)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10728, tmp10730)
}
__typedArg0 := tmp10728
__typedArg1 := tmp10730
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.retract-clause"))
}
__typedArg0 := MakeString("partial function shen.retract-clause")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 2)

tmp10752 := Call(__e, ns2_1set, symshen_4retract_1clause, tmp10721)


_ = tmp10752

tmp10753 := MakeNative(func(__e *ControlFlow) {
V1308 := __e.Get(1)
_ = V1308
V1309 := __e.Get(2)
_ = V1309
tmp10754 := MakeNative(func(__e *ControlFlow) {
Z1310 := __e.Get(1)
_ = Z1310
__e.TailApply(PrimFunc(symshen_4_5defprolog_6), Z1310)
return
}, 1)

tmp10755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1308, V1309)
}
__typedArg0 := V1308
__typedArg1 := V1309
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symcompile), tmp10754, tmp10755)
return


}, 2)

tmp10756 := Call(__e, ns2_1set, symshen_4compile_1prolog, tmp10753)


_ = tmp10756

tmp10757 := MakeNative(func(__e *ControlFlow) {
V1311 := __e.Get(1)
_ = V1311
tmp10773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1311)
}
__typedArg0 := V1311
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10758 Obj

if True == tmp10773 {
tmp10759 := Call(__e, PrimFunc(symhead), V1311)


let__10385 := tmp10759
_ = let__10385

tmp10760 := Call(__e, PrimFunc(symtail), V1311)


let__10386 := tmp10760
_ = let__10386

tmp10761 := Call(__e, PrimFunc(symshen_4_5clauses_6), let__10386)


let__10387 := tmp10761
_ = let__10387

tmp10771 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10387)


var ifres10762 Obj

if True == tmp10771 {
tmp10763 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10762 = tmp10763


} else {
tmp10764 := Call(__e, PrimFunc(symshen_4_5_1out), let__10387)


let__10388 := tmp10764
_ = let__10388

tmp10765 := Call(__e, PrimFunc(symshen_4in_1_6), let__10387)


let__10389 := tmp10765
_ = let__10389

tmp10766 := Call(__e, PrimFunc(symshen_4prolog_1arity_1check), let__10385, let__10388)


let__10390 := tmp10766
_ = let__10390

tmp10767 := MakeNative(func(__e *ControlFlow) {
Z1320 := __e.Get(1)
_ = Z1320
__e.TailApply(PrimFunc(symshen_4linearise_1clause), Z1320)
return
}, 1)

tmp10768 := Call(__e, PrimFunc(symmap), tmp10767, let__10388)


let__10391 := tmp10768
_ = let__10391

tmp10769 := Call(__e, PrimFunc(symshen_4horn_1clause_1procedure), let__10385, let__10391)


tmp10770 := Call(__e, PrimFunc(symshen_4comb), let__10389, tmp10769)


ifres10762 = tmp10770


}

ifres10758 = ifres10762


} else {
tmp10772 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10758 = tmp10772


}

let__10384 := ifres10758
_ = let__10384

tmp10775 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10384)


if True == tmp10775 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10384)
return
}


}, 1)

tmp10776 := Call(__e, ns2_1set, symshen_4_5defprolog_6, tmp10757)


_ = tmp10776

tmp10777 := MakeNative(func(__e *ControlFlow) {
V1323 := __e.Get(1)
_ = V1323
V1324 := __e.Get(2)
_ = V1324
tmp10821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10802 Obj

if True == tmp10821 {
tmp10819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10819)
}
__typedArg0 := tmp10819
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10804 Obj

if True == tmp10820 {
tmp10816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10816)
}
__typedArg0 := tmp10816
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10817)
}
__typedArg0 := tmp10817
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10806 Obj

if True == tmp10818 {
tmp10812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10812)
}
__typedArg0 := tmp10812
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10813)
}
__typedArg0 := tmp10813
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10814)
}
__typedArg0 := Nil
__typedArg1 := tmp10814
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10808 Obj

if True == tmp10815 {
tmp10810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10810)
}
__typedArg0 := Nil
__typedArg1 := tmp10810
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10809 Obj

if True == tmp10811 {
ifres10809 = True


} else {
ifres10809 = False


}

ifres10808 = ifres10809


} else {
ifres10808 = False


}

var ifres10807 Obj

if True == ifres10808 {
ifres10807 = True


} else {
ifres10807 = False


}

ifres10806 = ifres10807


} else {
ifres10806 = False


}

var ifres10805 Obj

if True == ifres10806 {
ifres10805 = True


} else {
ifres10805 = False


}

ifres10804 = ifres10805


} else {
ifres10804 = False


}

var ifres10803 Obj

if True == ifres10804 {
ifres10803 = True


} else {
ifres10803 = False


}

ifres10802 = ifres10803


} else {
ifres10802 = False


}

if True == ifres10802 {
tmp10778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10778)
}
__typedArg0 := tmp10778
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symlength), tmp10779)
return


} else {
tmp10800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10785 Obj

if True == tmp10800 {
tmp10798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10798)
}
__typedArg0 := tmp10798
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10787 Obj

if True == tmp10799 {
tmp10795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10795)
}
__typedArg0 := tmp10795
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10796)
}
__typedArg0 := tmp10796
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10789 Obj

if True == tmp10797 {
tmp10791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10791)
}
__typedArg0 := tmp10791
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10792)
}
__typedArg0 := tmp10792
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10793)
}
__typedArg0 := Nil
__typedArg1 := tmp10793
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10790 Obj

if True == tmp10794 {
ifres10790 = True


} else {
ifres10790 = False


}

ifres10789 = ifres10790


} else {
ifres10789 = False


}

var ifres10788 Obj

if True == ifres10789 {
ifres10788 = True


} else {
ifres10788 = False


}

ifres10787 = ifres10788


} else {
ifres10787 = False


}

var ifres10786 Obj

if True == ifres10787 {
ifres10786 = True


} else {
ifres10786 = False


}

ifres10785 = ifres10786


} else {
ifres10785 = False


}

if True == ifres10785 {
tmp10780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10780)
}
__typedArg0 := tmp10780
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10782 := Call(__e, PrimFunc(symlength), tmp10781)


tmp10783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1324)
}
__typedArg0 := V1324
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4pac_1h), V1323, tmp10782, tmp10783)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.prolog-arity-check"))
}
__typedArg0 := MakeString("partial function shen.prolog-arity-check")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp10822 := Call(__e, ns2_1set, symshen_4prolog_1arity_1check, tmp10777)


_ = tmp10822

tmp10823 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__selftop:
V1329 := __self1
_ = V1329
V1330 := __self2
_ = V1330
V1331 := __self3
_ = V1331
tmp10839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1331)
}
__typedArg0 := Nil
__typedArg1 := V1331
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10839 {
__e.Return(V1330)
return
} else {
tmp10837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1331)
}
__typedArg0 := V1331
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10833 Obj

if True == tmp10837 {
tmp10835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1331)
}
__typedArg0 := V1331
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10835)
}
__typedArg0 := tmp10835
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10834 Obj

if True == tmp10836 {
ifres10834 = True


} else {
ifres10834 = False


}

ifres10833 = ifres10834


} else {
ifres10833 = False


}

if True == ifres10833 {
tmp10828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1331)
}
__typedArg0 := V1331
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10828)
}
__typedArg0 := tmp10828
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10830 := Call(__e, PrimFunc(symlength), tmp10829)


tmp10831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1330, tmp10830)
}
__typedArg0 := V1330
__typedArg1 := tmp10830
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp10831 {
tmp10824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1331)
}
__typedArg0 := V1331
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4pac_1h) == __self {
__self1, __self2, __self3 = V1329, V1330, tmp10824
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4pac_1h), V1329, V1330, tmp10824)
return


} else {
tmp10825 := Call(__e, PrimFunc(symshen_4app), V1329, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("arity error in prolog procedure "))
__typedS1, __typedOK1 := TypedString(tmp10825)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("arity error in prolog procedure ")
__typedArg1 := tmp10825
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("arity error in prolog procedure "))
__typedS1, __typedOK1 := TypedString(tmp10825)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("arity error in prolog procedure ")
__typedArg1 := tmp10825
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.pac-h"))
}
__typedArg0 := MakeString("partial function shen.pac-h")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 3)

tmp10840 := Call(__e, ns2_1set, symshen_4pac_1h, tmp10823)


_ = tmp10840

tmp10841 := MakeNative(func(__e *ControlFlow) {
V1332 := __e.Get(1)
_ = V1332
tmp10842 := Call(__e, PrimFunc(symshen_4_5clause_6), V1332)


let__10393 := tmp10842
_ = let__10393

tmp10855 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10393)


var ifres10843 Obj

if True == tmp10855 {
tmp10844 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10843 = tmp10844


} else {
tmp10845 := Call(__e, PrimFunc(symshen_4_5_1out), let__10393)


let__10394 := tmp10845
_ = let__10394

tmp10846 := Call(__e, PrimFunc(symshen_4in_1_6), let__10393)


let__10395 := tmp10846
_ = let__10395

tmp10847 := Call(__e, PrimFunc(symshen_4_5clauses_6), let__10395)


let__10396 := tmp10847
_ = let__10396

tmp10854 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10396)


var ifres10848 Obj

if True == tmp10854 {
tmp10849 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10848 = tmp10849


} else {
tmp10850 := Call(__e, PrimFunc(symshen_4_5_1out), let__10396)


let__10397 := tmp10850
_ = let__10397

tmp10851 := Call(__e, PrimFunc(symshen_4in_1_6), let__10396)


let__10398 := tmp10851
_ = let__10398

tmp10852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10394, let__10397)
}
__typedArg0 := let__10394
__typedArg1 := let__10397
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10853 := Call(__e, PrimFunc(symshen_4comb), let__10398, tmp10852)


ifres10848 = tmp10853


}

ifres10843 = ifres10848


}

let__10392 := ifres10843
_ = let__10392

tmp10871 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10392)


if True == tmp10871 {
tmp10856 := Call(__e, PrimFunc(sym_5_b_6), V1332)


let__10400 := tmp10856
_ = let__10400

tmp10867 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10400)


var ifres10857 Obj

if True == tmp10867 {
tmp10858 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10857 = tmp10858


} else {
tmp10859 := Call(__e, PrimFunc(symshen_4_5_1out), let__10400)


let__10401 := tmp10859
_ = let__10401

tmp10860 := Call(__e, PrimFunc(symshen_4in_1_6), let__10400)


let__10402 := tmp10860
_ = let__10402

tmp10865 := Call(__e, PrimFunc(symempty_2), let__10401)


var ifres10861 Obj

if True == tmp10865 {
ifres10861 = Nil


} else {
tmp10862 := Call(__e, PrimFunc(symshen_4app), let__10401, MakeString("\n ..."), symshen_4r)


tmp10864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Prolog syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp10862)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Prolog syntax error here:\n ")
__typedArg1 := tmp10862
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Prolog syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp10862)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Prolog syntax error here:\n ")
__typedArg1 := tmp10862
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres10861 = tmp10864


}

tmp10866 := Call(__e, PrimFunc(symshen_4comb), let__10402, ifres10861)


ifres10857 = tmp10866


}

let__10399 := ifres10857
_ = let__10399

tmp10869 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10399)


if True == tmp10869 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10399)
return
}


} else {
__e.Return(let__10392)
return
}


}, 1)

tmp10872 := Call(__e, ns2_1set, symshen_4_5clauses_6, tmp10841)


_ = tmp10872

tmp10873 := MakeNative(func(__e *ControlFlow) {
V1344 := __e.Get(1)
_ = V1344
tmp10889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1344)
}
__typedArg0 := V1344
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10880 Obj

if True == tmp10889 {
tmp10887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1344)
}
__typedArg0 := V1344
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10887)
}
__typedArg0 := tmp10887
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10882 Obj

if True == tmp10888 {
tmp10884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1344)
}
__typedArg0 := V1344
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10884)
}
__typedArg0 := tmp10884
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10885)
}
__typedArg0 := Nil
__typedArg1 := tmp10885
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10883 Obj

if True == tmp10886 {
ifres10883 = True


} else {
ifres10883 = False


}

ifres10882 = ifres10883


} else {
ifres10882 = False


}

var ifres10881 Obj

if True == ifres10882 {
ifres10881 = True


} else {
ifres10881 = False


}

ifres10880 = ifres10881


} else {
ifres10880 = False


}

if True == ifres10880 {
tmp10874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1344)
}
__typedArg0 := V1344
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1344)
}
__typedArg0 := V1344
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10875)
}
__typedArg0 := tmp10875
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10877 := Call(__e, PrimFunc(sym_8p), tmp10874, tmp10876)


tmp10878 := Call(__e, PrimFunc(symshen_4linearise), tmp10877)


__e.TailApply(PrimFunc(symshen_4lch), tmp10878)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.linearise-clause"))
}
__typedArg0 := MakeString("partial function shen.linearise-clause")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp10890 := Call(__e, ns2_1set, symshen_4linearise_1clause, tmp10873)


_ = tmp10890

tmp10891 := MakeNative(func(__e *ControlFlow) {
V1345 := __e.Get(1)
_ = V1345
tmp10897 := Call(__e, PrimFunc(symtuple_2), V1345)


if True == tmp10897 {
tmp10892 := Call(__e, PrimFunc(symfst), V1345)


tmp10893 := Call(__e, PrimFunc(symsnd), V1345)


tmp10894 := Call(__e, PrimFunc(symshen_4lchh), tmp10893)


tmp10895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10894, Nil)
}
__typedArg0 := tmp10894
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10892, tmp10895)
}
__typedArg0 := tmp10892
__typedArg1 := tmp10895
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.lch"))
}
__typedArg0 := MakeString("partial function shen.lch")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp10898 := Call(__e, ns2_1set, symshen_4lch, tmp10891)


_ = tmp10898

tmp10899 := MakeNative(func(__e *ControlFlow) {
V1346 := __e.Get(1)
_ = V1346
tmp10962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1346)
}
__typedArg0 := V1346
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10911 Obj

if True == tmp10962 {
tmp10960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1346)
}
__typedArg0 := V1346
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp10960)
}
__typedArg0 := symwhere
__typedArg1 := tmp10960
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10913 Obj

if True == tmp10961 {
tmp10958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1346)
}
__typedArg0 := V1346
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10958)
}
__typedArg0 := tmp10958
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10915 Obj

if True == tmp10959 {
tmp10955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1346)
}
__typedArg0 := V1346
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10955)
}
__typedArg0 := tmp10955
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10956)
}
__typedArg0 := tmp10956
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10917 Obj

if True == tmp10957 {
tmp10951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1346)
}
__typedArg0 := V1346
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10951)
}
__typedArg0 := tmp10951
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10952)
}
__typedArg0 := tmp10952
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_a, tmp10953)
}
__typedArg0 := sym_a
__typedArg1 := tmp10953
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10919 Obj

if True == tmp10954 {
tmp10947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1346)
}
__typedArg0 := V1346
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10947)
}
__typedArg0 := tmp10947
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10948)
}
__typedArg0 := tmp10948
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10949)
}
__typedArg0 := tmp10949
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10921 Obj

if True == tmp10950 {
tmp10942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1346)
}
__typedArg0 := V1346
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10942)
}
__typedArg0 := tmp10942
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10943)
}
__typedArg0 := tmp10943
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10944)
}
__typedArg0 := tmp10944
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10945)
}
__typedArg0 := tmp10945
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10923 Obj

if True == tmp10946 {
tmp10936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1346)
}
__typedArg0 := V1346
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10936)
}
__typedArg0 := tmp10936
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10937)
}
__typedArg0 := tmp10937
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10938)
}
__typedArg0 := tmp10938
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10939)
}
__typedArg0 := tmp10939
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10940)
}
__typedArg0 := Nil
__typedArg1 := tmp10940
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10925 Obj

if True == tmp10941 {
tmp10933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1346)
}
__typedArg0 := V1346
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10933)
}
__typedArg0 := tmp10933
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp10934)
}
__typedArg0 := tmp10934
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres10927 Obj

if True == tmp10935 {
tmp10929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1346)
}
__typedArg0 := V1346
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10929)
}
__typedArg0 := tmp10929
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10930)
}
__typedArg0 := tmp10930
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp10931)
}
__typedArg0 := Nil
__typedArg1 := tmp10931
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres10928 Obj

if True == tmp10932 {
ifres10928 = True


} else {
ifres10928 = False


}

ifres10927 = ifres10928


} else {
ifres10927 = False


}

var ifres10926 Obj

if True == ifres10927 {
ifres10926 = True


} else {
ifres10926 = False


}

ifres10925 = ifres10926


} else {
ifres10925 = False


}

var ifres10924 Obj

if True == ifres10925 {
ifres10924 = True


} else {
ifres10924 = False


}

ifres10923 = ifres10924


} else {
ifres10923 = False


}

var ifres10922 Obj

if True == ifres10923 {
ifres10922 = True


} else {
ifres10922 = False


}

ifres10921 = ifres10922


} else {
ifres10921 = False


}

var ifres10920 Obj

if True == ifres10921 {
ifres10920 = True


} else {
ifres10920 = False


}

ifres10919 = ifres10920


} else {
ifres10919 = False


}

var ifres10918 Obj

if True == ifres10919 {
ifres10918 = True


} else {
ifres10918 = False


}

ifres10917 = ifres10918


} else {
ifres10917 = False


}

var ifres10916 Obj

if True == ifres10917 {
ifres10916 = True


} else {
ifres10916 = False


}

ifres10915 = ifres10916


} else {
ifres10915 = False


}

var ifres10914 Obj

if True == ifres10915 {
ifres10914 = True


} else {
ifres10914 = False


}

ifres10913 = ifres10914


} else {
ifres10913 = False


}

var ifres10912 Obj

if True == ifres10913 {
ifres10912 = True


} else {
ifres10912 = False


}

ifres10911 = ifres10912


} else {
ifres10911 = False


}

if True == ifres10911 {
tmp10901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_doccurs_d)
}
__typedArg0 := symshen_4_doccurs_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

var ifres10900 Obj

if True == tmp10901 {
ifres10900 = symis_b


} else {
ifres10900 = symis


}

tmp10902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1346)
}
__typedArg0 := V1346
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10902)
}
__typedArg0 := tmp10902
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10903)
}
__typedArg0 := tmp10903
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(ifres10900, tmp10904)
}
__typedArg0 := ifres10900
__typedArg1 := tmp10904
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1346)
}
__typedArg0 := V1346
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp10906)
}
__typedArg0 := tmp10906
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp10908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp10907)
}
__typedArg0 := tmp10907
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp10909 := Call(__e, PrimFunc(symshen_4lchh), tmp10908)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp10905, tmp10909)
}
__typedArg0 := tmp10905
__typedArg1 := tmp10909
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V1346)
return
}


}, 1)

tmp10963 := Call(__e, ns2_1set, symshen_4lchh, tmp10899)


_ = tmp10963

tmp10964 := MakeNative(func(__e *ControlFlow) {
V1347 := __e.Get(1)
_ = V1347
tmp10965 := Call(__e, PrimFunc(symshen_4_5head_6), V1347)


let__10404 := tmp10965
_ = let__10404

tmp10988 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10404)


var ifres10966 Obj

if True == tmp10988 {
tmp10967 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10966 = tmp10967


} else {
tmp10968 := Call(__e, PrimFunc(symshen_4_5_1out), let__10404)


let__10405 := tmp10968
_ = let__10405

tmp10969 := Call(__e, PrimFunc(symshen_4in_1_6), let__10404)


let__10406 := tmp10969
_ = let__10406

tmp10987 := Call(__e, PrimFunc(symshen_4hds_a_2), let__10406, sym_5_1_1)


var ifres10970 Obj

if True == tmp10987 {
tmp10971 := Call(__e, PrimFunc(symtail), let__10406)


let__10407 := tmp10971
_ = let__10407

tmp10972 := Call(__e, PrimFunc(symshen_4_5body_6), let__10407)


let__10408 := tmp10972
_ = let__10408

tmp10985 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10408)


var ifres10973 Obj

if True == tmp10985 {
tmp10974 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10973 = tmp10974


} else {
tmp10975 := Call(__e, PrimFunc(symshen_4_5_1out), let__10408)


let__10409 := tmp10975
_ = let__10409

tmp10976 := Call(__e, PrimFunc(symshen_4in_1_6), let__10408)


let__10410 := tmp10976
_ = let__10410

tmp10977 := Call(__e, PrimFunc(symshen_4_5sc_6), let__10410)


let__10411 := tmp10977
_ = let__10411

tmp10984 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10411)


var ifres10978 Obj

if True == tmp10984 {
tmp10979 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10978 = tmp10979


} else {
tmp10980 := Call(__e, PrimFunc(symshen_4in_1_6), let__10411)


let__10412 := tmp10980
_ = let__10412

tmp10981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10409, Nil)
}
__typedArg0 := let__10409
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10405, tmp10981)
}
__typedArg0 := let__10405
__typedArg1 := tmp10981
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp10983 := Call(__e, PrimFunc(symshen_4comb), let__10412, tmp10982)


ifres10978 = tmp10983


}

ifres10973 = ifres10978


}

ifres10970 = ifres10973


} else {
tmp10986 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10970 = tmp10986


}

ifres10966 = ifres10970


}

let__10403 := ifres10966
_ = let__10403

tmp10990 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10403)


if True == tmp10990 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10403)
return
}


}, 1)

tmp10991 := Call(__e, ns2_1set, symshen_4_5clause_6, tmp10964)


_ = tmp10991

tmp10992 := MakeNative(func(__e *ControlFlow) {
V1358 := __e.Get(1)
_ = V1358
tmp10993 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1358)


let__10414 := tmp10993
_ = let__10414

tmp11006 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10414)


var ifres10994 Obj

if True == tmp11006 {
tmp10995 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10994 = tmp10995


} else {
tmp10996 := Call(__e, PrimFunc(symshen_4_5_1out), let__10414)


let__10415 := tmp10996
_ = let__10415

tmp10997 := Call(__e, PrimFunc(symshen_4in_1_6), let__10414)


let__10416 := tmp10997
_ = let__10416

tmp10998 := Call(__e, PrimFunc(symshen_4_5head_6), let__10416)


let__10417 := tmp10998
_ = let__10417

tmp11005 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10417)


var ifres10999 Obj

if True == tmp11005 {
tmp11000 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10999 = tmp11000


} else {
tmp11001 := Call(__e, PrimFunc(symshen_4_5_1out), let__10417)


let__10418 := tmp11001
_ = let__10418

tmp11002 := Call(__e, PrimFunc(symshen_4in_1_6), let__10417)


let__10419 := tmp11002
_ = let__10419

tmp11003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10415, let__10418)
}
__typedArg0 := let__10415
__typedArg1 := let__10418
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11004 := Call(__e, PrimFunc(symshen_4comb), let__10419, tmp11003)


ifres10999 = tmp11004


}

ifres10994 = ifres10999


}

let__10413 := ifres10994
_ = let__10413

tmp11016 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10413)


if True == tmp11016 {
tmp11007 := Call(__e, PrimFunc(sym_5e_6), V1358)


let__10421 := tmp11007
_ = let__10421

tmp11012 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10421)


var ifres11008 Obj

if True == tmp11012 {
tmp11009 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11008 = tmp11009


} else {
tmp11010 := Call(__e, PrimFunc(symshen_4in_1_6), let__10421)


let__10422 := tmp11010
_ = let__10422

tmp11011 := Call(__e, PrimFunc(symshen_4comb), let__10422, Nil)


ifres11008 = tmp11011


}

let__10420 := ifres11008
_ = let__10420

tmp11014 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10420)


if True == tmp11014 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10420)
return
}


} else {
__e.Return(let__10413)
return
}


}, 1)

tmp11017 := Call(__e, ns2_1set, symshen_4_5head_6, tmp10992)


_ = tmp11017

tmp11018 := MakeNative(func(__e *ControlFlow) {
V1369 := __e.Get(1)
_ = V1369
tmp11031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1369)
}
__typedArg0 := V1369
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11019 Obj

if True == tmp11031 {
tmp11020 := Call(__e, PrimFunc(symhead), V1369)


let__10424 := tmp11020
_ = let__10424

tmp11021 := Call(__e, PrimFunc(symtail), V1369)


let__10425 := tmp11021
_ = let__10425

tmp11029 := Call(__e, PrimFunc(symatom_2), let__10424)


var ifres11025 Obj

if True == tmp11029 {
tmp11027 := Call(__e, PrimFunc(symshen_4prolog_1keyword_2), let__10424)


tmp11028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp11027)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp11027
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres11026 Obj

if True == tmp11028 {
ifres11026 = True


} else {
ifres11026 = False


}

ifres11025 = ifres11026


} else {
ifres11025 = False


}

var ifres11022 Obj

if True == ifres11025 {
tmp11023 := Call(__e, PrimFunc(symshen_4comb), let__10425, let__10424)


ifres11022 = tmp11023


} else {
tmp11024 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11022 = tmp11024


}

ifres11019 = ifres11022


} else {
tmp11030 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11019 = tmp11030


}

let__10423 := ifres11019
_ = let__10423

tmp11185 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10423)


if True == tmp11185 {
tmp11041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1369)
}
__typedArg0 := V1369
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11032 Obj

if True == tmp11041 {
tmp11033 := Call(__e, PrimFunc(symhead), V1369)


let__10427 := tmp11033
_ = let__10427

tmp11034 := Call(__e, PrimFunc(symtail), V1369)


let__10428 := tmp11034
_ = let__10428

tmp11038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp11039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10427, tmp11038)
}
__typedArg0 := let__10427
__typedArg1 := tmp11038
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11035 Obj

if True == tmp11039 {
tmp11036 := Call(__e, PrimFunc(symshen_4comb), let__10428, let__10427)


ifres11035 = tmp11036


} else {
tmp11037 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11035 = tmp11037


}

ifres11032 = ifres11035


} else {
tmp11040 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11032 = tmp11040


}

let__10426 := ifres11032
_ = let__10426

tmp11183 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10426)


if True == tmp11183 {
tmp11071 := Call(__e, PrimFunc(symshen_4ccons_2), V1369)


var ifres11042 Obj

if True == tmp11071 {
tmp11043 := Call(__e, PrimFunc(symhead), V1369)


let__10430 := tmp11043
_ = let__10430

tmp11044 := Call(__e, PrimFunc(symtail), V1369)


let__10431 := tmp11044
_ = let__10431

tmp11069 := Call(__e, PrimFunc(symshen_4hds_a_2), let__10430, symcons)


var ifres11045 Obj

if True == tmp11069 {
tmp11046 := Call(__e, PrimFunc(symtail), let__10430)


let__10432 := tmp11046
_ = let__10432

tmp11047 := Call(__e, PrimFunc(symshen_4_5hterm1_6), let__10432)


let__10433 := tmp11047
_ = let__10433

tmp11067 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10433)


var ifres11048 Obj

if True == tmp11067 {
tmp11049 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11048 = tmp11049


} else {
tmp11050 := Call(__e, PrimFunc(symshen_4_5_1out), let__10433)


let__10434 := tmp11050
_ = let__10434

tmp11051 := Call(__e, PrimFunc(symshen_4in_1_6), let__10433)


let__10435 := tmp11051
_ = let__10435

tmp11052 := Call(__e, PrimFunc(symshen_4_5hterm2_6), let__10435)


let__10436 := tmp11052
_ = let__10436

tmp11066 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10436)


var ifres11053 Obj

if True == tmp11066 {
tmp11054 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11053 = tmp11054


} else {
tmp11055 := Call(__e, PrimFunc(symshen_4_5_1out), let__10436)


let__10437 := tmp11055
_ = let__10437

tmp11056 := Call(__e, PrimFunc(symshen_4in_1_6), let__10436)


let__10438 := tmp11056
_ = let__10438

tmp11057 := Call(__e, PrimFunc(sym_5end_6), let__10438)


let__10439 := tmp11057
_ = let__10439

tmp11065 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10439)


var ifres11058 Obj

if True == tmp11065 {
tmp11059 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11058 = tmp11059


} else {
tmp11060 := Call(__e, PrimFunc(symshen_4in_1_6), let__10439)


let__10440 := tmp11060
_ = let__10440

tmp11061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10437, Nil)
}
__typedArg0 := let__10437
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10434, tmp11061)
}
__typedArg0 := let__10434
__typedArg1 := tmp11061
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp11062)
}
__typedArg0 := symcons
__typedArg1 := tmp11062
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11064 := Call(__e, PrimFunc(symshen_4comb), let__10431, tmp11063)


ifres11058 = tmp11064


}

ifres11053 = ifres11058


}

ifres11048 = ifres11053


}

ifres11045 = ifres11048


} else {
tmp11068 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11045 = tmp11068


}

ifres11042 = ifres11045


} else {
tmp11070 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11042 = tmp11070


}

let__10429 := ifres11042
_ = let__10429

tmp11181 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10429)


if True == tmp11181 {
tmp11094 := Call(__e, PrimFunc(symshen_4ccons_2), V1369)


var ifres11072 Obj

if True == tmp11094 {
tmp11073 := Call(__e, PrimFunc(symhead), V1369)


let__10442 := tmp11073
_ = let__10442

tmp11074 := Call(__e, PrimFunc(symtail), V1369)


let__10443 := tmp11074
_ = let__10443

tmp11092 := Call(__e, PrimFunc(symshen_4hds_a_2), let__10442, sym_7)


var ifres11075 Obj

if True == tmp11092 {
tmp11076 := Call(__e, PrimFunc(symtail), let__10442)


let__10444 := tmp11076
_ = let__10444

tmp11077 := Call(__e, PrimFunc(symshen_4_5hterm_6), let__10444)


let__10445 := tmp11077
_ = let__10445

tmp11090 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10445)


var ifres11078 Obj

if True == tmp11090 {
tmp11079 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11078 = tmp11079


} else {
tmp11080 := Call(__e, PrimFunc(symshen_4_5_1out), let__10445)


let__10446 := tmp11080
_ = let__10446

tmp11081 := Call(__e, PrimFunc(symshen_4in_1_6), let__10445)


let__10447 := tmp11081
_ = let__10447

tmp11082 := Call(__e, PrimFunc(sym_5end_6), let__10447)


let__10448 := tmp11082
_ = let__10448

tmp11089 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10448)


var ifres11083 Obj

if True == tmp11089 {
tmp11084 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11083 = tmp11084


} else {
tmp11085 := Call(__e, PrimFunc(symshen_4in_1_6), let__10448)


let__10449 := tmp11085
_ = let__10449

tmp11086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10446, Nil)
}
__typedArg0 := let__10446
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_7m, tmp11086)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp11086
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11088 := Call(__e, PrimFunc(symshen_4comb), let__10443, tmp11087)


ifres11083 = tmp11088


}

ifres11078 = ifres11083


}

ifres11075 = ifres11078


} else {
tmp11091 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11075 = tmp11091


}

ifres11072 = ifres11075


} else {
tmp11093 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11072 = tmp11093


}

let__10441 := ifres11072
_ = let__10441

tmp11179 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10441)


if True == tmp11179 {
tmp11117 := Call(__e, PrimFunc(symshen_4ccons_2), V1369)


var ifres11095 Obj

if True == tmp11117 {
tmp11096 := Call(__e, PrimFunc(symhead), V1369)


let__10451 := tmp11096
_ = let__10451

tmp11097 := Call(__e, PrimFunc(symtail), V1369)


let__10452 := tmp11097
_ = let__10452

tmp11115 := Call(__e, PrimFunc(symshen_4hds_a_2), let__10451, sym_1)


var ifres11098 Obj

if True == tmp11115 {
tmp11099 := Call(__e, PrimFunc(symtail), let__10451)


let__10453 := tmp11099
_ = let__10453

tmp11100 := Call(__e, PrimFunc(symshen_4_5hterm_6), let__10453)


let__10454 := tmp11100
_ = let__10454

tmp11113 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10454)


var ifres11101 Obj

if True == tmp11113 {
tmp11102 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11101 = tmp11102


} else {
tmp11103 := Call(__e, PrimFunc(symshen_4_5_1out), let__10454)


let__10455 := tmp11103
_ = let__10455

tmp11104 := Call(__e, PrimFunc(symshen_4in_1_6), let__10454)


let__10456 := tmp11104
_ = let__10456

tmp11105 := Call(__e, PrimFunc(sym_5end_6), let__10456)


let__10457 := tmp11105
_ = let__10457

tmp11112 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10457)


var ifres11106 Obj

if True == tmp11112 {
tmp11107 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11106 = tmp11107


} else {
tmp11108 := Call(__e, PrimFunc(symshen_4in_1_6), let__10457)


let__10458 := tmp11108
_ = let__10458

tmp11109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10455, Nil)
}
__typedArg0 := let__10455
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_1m, tmp11109)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp11109
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11111 := Call(__e, PrimFunc(symshen_4comb), let__10452, tmp11110)


ifres11106 = tmp11111


}

ifres11101 = ifres11106


}

ifres11098 = ifres11101


} else {
tmp11114 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11098 = tmp11114


}

ifres11095 = ifres11098


} else {
tmp11116 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11095 = tmp11116


}

let__10450 := ifres11095
_ = let__10450

tmp11177 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10450)


if True == tmp11177 {
tmp11144 := Call(__e, PrimFunc(symshen_4ccons_2), V1369)


var ifres11118 Obj

if True == tmp11144 {
tmp11119 := Call(__e, PrimFunc(symhead), V1369)


let__10460 := tmp11119
_ = let__10460

tmp11120 := Call(__e, PrimFunc(symtail), V1369)


let__10461 := tmp11120
_ = let__10461

tmp11142 := Call(__e, PrimFunc(symshen_4hds_a_2), let__10460, symmode)


var ifres11121 Obj

if True == tmp11142 {
tmp11122 := Call(__e, PrimFunc(symtail), let__10460)


let__10462 := tmp11122
_ = let__10462

tmp11123 := Call(__e, PrimFunc(symshen_4_5hterm_6), let__10462)


let__10463 := tmp11123
_ = let__10463

tmp11140 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10463)


var ifres11124 Obj

if True == tmp11140 {
tmp11125 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11124 = tmp11125


} else {
tmp11126 := Call(__e, PrimFunc(symshen_4_5_1out), let__10463)


let__10464 := tmp11126
_ = let__10464

tmp11127 := Call(__e, PrimFunc(symshen_4in_1_6), let__10463)


let__10465 := tmp11127
_ = let__10465

tmp11139 := Call(__e, PrimFunc(symshen_4hds_a_2), let__10465, sym_7)


var ifres11128 Obj

if True == tmp11139 {
tmp11129 := Call(__e, PrimFunc(symtail), let__10465)


let__10466 := tmp11129
_ = let__10466

tmp11130 := Call(__e, PrimFunc(sym_5end_6), let__10466)


let__10467 := tmp11130
_ = let__10467

tmp11137 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10467)


var ifres11131 Obj

if True == tmp11137 {
tmp11132 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11131 = tmp11132


} else {
tmp11133 := Call(__e, PrimFunc(symshen_4in_1_6), let__10467)


let__10468 := tmp11133
_ = let__10468

tmp11134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10464, Nil)
}
__typedArg0 := let__10464
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_7m, tmp11134)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp11134
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11136 := Call(__e, PrimFunc(symshen_4comb), let__10461, tmp11135)


ifres11131 = tmp11136


}

ifres11128 = ifres11131


} else {
tmp11138 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11128 = tmp11138


}

ifres11124 = ifres11128


}

ifres11121 = ifres11124


} else {
tmp11141 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11121 = tmp11141


}

ifres11118 = ifres11121


} else {
tmp11143 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11118 = tmp11143


}

let__10459 := ifres11118
_ = let__10459

tmp11175 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10459)


if True == tmp11175 {
tmp11171 := Call(__e, PrimFunc(symshen_4ccons_2), V1369)


var ifres11145 Obj

if True == tmp11171 {
tmp11146 := Call(__e, PrimFunc(symhead), V1369)


let__10470 := tmp11146
_ = let__10470

tmp11147 := Call(__e, PrimFunc(symtail), V1369)


let__10471 := tmp11147
_ = let__10471

tmp11169 := Call(__e, PrimFunc(symshen_4hds_a_2), let__10470, symmode)


var ifres11148 Obj

if True == tmp11169 {
tmp11149 := Call(__e, PrimFunc(symtail), let__10470)


let__10472 := tmp11149
_ = let__10472

tmp11150 := Call(__e, PrimFunc(symshen_4_5hterm_6), let__10472)


let__10473 := tmp11150
_ = let__10473

tmp11167 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10473)


var ifres11151 Obj

if True == tmp11167 {
tmp11152 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11151 = tmp11152


} else {
tmp11153 := Call(__e, PrimFunc(symshen_4_5_1out), let__10473)


let__10474 := tmp11153
_ = let__10474

tmp11154 := Call(__e, PrimFunc(symshen_4in_1_6), let__10473)


let__10475 := tmp11154
_ = let__10475

tmp11166 := Call(__e, PrimFunc(symshen_4hds_a_2), let__10475, sym_1)


var ifres11155 Obj

if True == tmp11166 {
tmp11156 := Call(__e, PrimFunc(symtail), let__10475)


let__10476 := tmp11156
_ = let__10476

tmp11157 := Call(__e, PrimFunc(sym_5end_6), let__10476)


let__10477 := tmp11157
_ = let__10477

tmp11164 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10477)


var ifres11158 Obj

if True == tmp11164 {
tmp11159 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11158 = tmp11159


} else {
tmp11160 := Call(__e, PrimFunc(symshen_4in_1_6), let__10477)


let__10478 := tmp11160
_ = let__10478

tmp11161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10474, Nil)
}
__typedArg0 := let__10474
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_1m, tmp11161)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp11161
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11163 := Call(__e, PrimFunc(symshen_4comb), let__10471, tmp11162)


ifres11158 = tmp11163


}

ifres11155 = ifres11158


} else {
tmp11165 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11155 = tmp11165


}

ifres11151 = ifres11155


}

ifres11148 = ifres11151


} else {
tmp11168 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11148 = tmp11168


}

ifres11145 = ifres11148


} else {
tmp11170 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11145 = tmp11170


}

let__10469 := ifres11145
_ = let__10469

tmp11173 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10469)


if True == tmp11173 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10469)
return
}


} else {
__e.Return(let__10459)
return
}


} else {
__e.Return(let__10450)
return
}


} else {
__e.Return(let__10441)
return
}


} else {
__e.Return(let__10429)
return
}


} else {
__e.Return(let__10426)
return
}


} else {
__e.Return(let__10423)
return
}


}, 1)

tmp11186 := Call(__e, ns2_1set, symshen_4_5hterm_6, tmp11018)


_ = tmp11186

tmp11187 := MakeNative(func(__e *ControlFlow) {
V1426 := __e.Get(1)
_ = V1426
tmp11188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp11189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11188, tmp11189)
}
__typedArg0 := tmp11188
__typedArg1 := tmp11189
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symelement_2), V1426, tmp11190)
return


}, 1)

tmp11191 := Call(__e, ns2_1set, symshen_4prolog_1keyword_2, tmp11187)


_ = tmp11191

tmp11192 := MakeNative(func(__e *ControlFlow) {
V1427 := __e.Get(1)
_ = V1427
tmp11205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V1427)
}
__typedArg0 := V1427
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp11205 {
__e.Return(True)
return
} else {
tmp11203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V1427)
}
__typedArg0 := V1427
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})()

var ifres11194 Obj

if True == tmp11203 {
ifres11194 = True


} else {
tmp11202 := Call(__e, PrimFunc(symboolean_2), V1427)


var ifres11196 Obj

if True == tmp11202 {
ifres11196 = True


} else {
tmp11201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnumber_2) {
return PrimIsNumber(V1427)
}
__typedArg0 := V1427
return Call(__e, PrimFunc(symnumber_2), __typedArg0)
})()

var ifres11198 Obj

if True == tmp11201 {
ifres11198 = True


} else {
tmp11200 := Call(__e, PrimFunc(symempty_2), V1427)


var ifres11199 Obj

if True == tmp11200 {
ifres11199 = True


} else {
ifres11199 = False


}

ifres11198 = ifres11199


}

var ifres11197 Obj

if True == ifres11198 {
ifres11197 = True


} else {
ifres11197 = False


}

ifres11196 = ifres11197


}

var ifres11195 Obj

if True == ifres11196 {
ifres11195 = True


} else {
ifres11195 = False


}

ifres11194 = ifres11195


}

if True == ifres11194 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp11206 := Call(__e, ns2_1set, symatom_2, tmp11192)


_ = tmp11206

tmp11207 := MakeNative(func(__e *ControlFlow) {
V1428 := __e.Get(1)
_ = V1428
tmp11208 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1428)


let__10480 := tmp11208
_ = let__10480

tmp11214 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10480)


var ifres11209 Obj

if True == tmp11214 {
tmp11210 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11209 = tmp11210


} else {
tmp11211 := Call(__e, PrimFunc(symshen_4_5_1out), let__10480)


let__10481 := tmp11211
_ = let__10481

tmp11212 := Call(__e, PrimFunc(symshen_4in_1_6), let__10480)


let__10482 := tmp11212
_ = let__10482

tmp11213 := Call(__e, PrimFunc(symshen_4comb), let__10482, let__10481)


ifres11209 = tmp11213


}

let__10479 := ifres11209
_ = let__10479

tmp11216 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10479)


if True == tmp11216 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10479)
return
}


}, 1)

tmp11217 := Call(__e, ns2_1set, symshen_4_5hterm1_6, tmp11207)


_ = tmp11217

tmp11218 := MakeNative(func(__e *ControlFlow) {
V1433 := __e.Get(1)
_ = V1433
tmp11219 := Call(__e, PrimFunc(symshen_4_5hterm_6), V1433)


let__10484 := tmp11219
_ = let__10484

tmp11225 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10484)


var ifres11220 Obj

if True == tmp11225 {
tmp11221 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11220 = tmp11221


} else {
tmp11222 := Call(__e, PrimFunc(symshen_4_5_1out), let__10484)


let__10485 := tmp11222
_ = let__10485

tmp11223 := Call(__e, PrimFunc(symshen_4in_1_6), let__10484)


let__10486 := tmp11223
_ = let__10486

tmp11224 := Call(__e, PrimFunc(symshen_4comb), let__10486, let__10485)


ifres11220 = tmp11224


}

let__10483 := ifres11220
_ = let__10483

tmp11227 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10483)


if True == tmp11227 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10483)
return
}


}, 1)

tmp11228 := Call(__e, ns2_1set, symshen_4_5hterm2_6, tmp11218)


_ = tmp11228

tmp11229 := MakeNative(func(__e *ControlFlow) {
V1438 := __e.Get(1)
_ = V1438
tmp11230 := Call(__e, PrimFunc(symshen_4_5literal_6), V1438)


let__10488 := tmp11230
_ = let__10488

tmp11243 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10488)


var ifres11231 Obj

if True == tmp11243 {
tmp11232 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11231 = tmp11232


} else {
tmp11233 := Call(__e, PrimFunc(symshen_4_5_1out), let__10488)


let__10489 := tmp11233
_ = let__10489

tmp11234 := Call(__e, PrimFunc(symshen_4in_1_6), let__10488)


let__10490 := tmp11234
_ = let__10490

tmp11235 := Call(__e, PrimFunc(symshen_4_5body_6), let__10490)


let__10491 := tmp11235
_ = let__10491

tmp11242 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10491)


var ifres11236 Obj

if True == tmp11242 {
tmp11237 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11236 = tmp11237


} else {
tmp11238 := Call(__e, PrimFunc(symshen_4_5_1out), let__10491)


let__10492 := tmp11238
_ = let__10492

tmp11239 := Call(__e, PrimFunc(symshen_4in_1_6), let__10491)


let__10493 := tmp11239
_ = let__10493

tmp11240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10489, let__10492)
}
__typedArg0 := let__10489
__typedArg1 := let__10492
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11241 := Call(__e, PrimFunc(symshen_4comb), let__10493, tmp11240)


ifres11236 = tmp11241


}

ifres11231 = ifres11236


}

let__10487 := ifres11231
_ = let__10487

tmp11253 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10487)


if True == tmp11253 {
tmp11244 := Call(__e, PrimFunc(sym_5e_6), V1438)


let__10495 := tmp11244
_ = let__10495

tmp11249 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10495)


var ifres11245 Obj

if True == tmp11249 {
tmp11246 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11245 = tmp11246


} else {
tmp11247 := Call(__e, PrimFunc(symshen_4in_1_6), let__10495)


let__10496 := tmp11247
_ = let__10496

tmp11248 := Call(__e, PrimFunc(symshen_4comb), let__10496, Nil)


ifres11245 = tmp11248


}

let__10494 := ifres11245
_ = let__10494

tmp11251 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10494)


if True == tmp11251 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10494)
return
}


} else {
__e.Return(let__10487)
return
}


}, 1)

tmp11254 := Call(__e, ns2_1set, symshen_4_5body_6, tmp11229)


_ = tmp11254

tmp11255 := MakeNative(func(__e *ControlFlow) {
V1449 := __e.Get(1)
_ = V1449
tmp11260 := Call(__e, PrimFunc(symshen_4hds_a_2), V1449, sym_b)


var ifres11256 Obj

if True == tmp11260 {
tmp11257 := Call(__e, PrimFunc(symtail), V1449)


let__10498 := tmp11257
_ = let__10498

tmp11258 := Call(__e, PrimFunc(symshen_4comb), let__10498, sym_b)


ifres11256 = tmp11258


} else {
tmp11259 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11256 = tmp11259


}

let__10497 := ifres11256
_ = let__10497

tmp11281 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10497)


if True == tmp11281 {
tmp11277 := Call(__e, PrimFunc(symshen_4ccons_2), V1449)


var ifres11261 Obj

if True == tmp11277 {
tmp11262 := Call(__e, PrimFunc(symhead), V1449)


let__10500 := tmp11262
_ = let__10500

tmp11263 := Call(__e, PrimFunc(symtail), V1449)


let__10501 := tmp11263
_ = let__10501

tmp11264 := Call(__e, PrimFunc(symshen_4_5bterms_6), let__10500)


let__10502 := tmp11264
_ = let__10502

tmp11275 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10502)


var ifres11265 Obj

if True == tmp11275 {
tmp11266 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11265 = tmp11266


} else {
tmp11267 := Call(__e, PrimFunc(symshen_4_5_1out), let__10502)


let__10503 := tmp11267
_ = let__10503

tmp11268 := Call(__e, PrimFunc(symshen_4in_1_6), let__10502)


let__10504 := tmp11268
_ = let__10504

tmp11269 := Call(__e, PrimFunc(sym_5end_6), let__10504)


let__10505 := tmp11269
_ = let__10505

tmp11274 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10505)


var ifres11270 Obj

if True == tmp11274 {
tmp11271 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11270 = tmp11271


} else {
tmp11272 := Call(__e, PrimFunc(symshen_4in_1_6), let__10505)


let__10506 := tmp11272
_ = let__10506

tmp11273 := Call(__e, PrimFunc(symshen_4comb), let__10501, let__10503)


ifres11270 = tmp11273


}

ifres11265 = ifres11270


}

ifres11261 = ifres11265


} else {
tmp11276 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11261 = tmp11276


}

let__10499 := ifres11261
_ = let__10499

tmp11279 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10499)


if True == tmp11279 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10499)
return
}


} else {
__e.Return(let__10497)
return
}


}, 1)

tmp11282 := Call(__e, ns2_1set, symshen_4_5literal_6, tmp11255)


_ = tmp11282

tmp11283 := MakeNative(func(__e *ControlFlow) {
V1460 := __e.Get(1)
_ = V1460
tmp11284 := Call(__e, PrimFunc(symshen_4_5bterm_6), V1460)


let__10508 := tmp11284
_ = let__10508

tmp11297 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10508)


var ifres11285 Obj

if True == tmp11297 {
tmp11286 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11285 = tmp11286


} else {
tmp11287 := Call(__e, PrimFunc(symshen_4_5_1out), let__10508)


let__10509 := tmp11287
_ = let__10509

tmp11288 := Call(__e, PrimFunc(symshen_4in_1_6), let__10508)


let__10510 := tmp11288
_ = let__10510

tmp11289 := Call(__e, PrimFunc(symshen_4_5bterms_6), let__10510)


let__10511 := tmp11289
_ = let__10511

tmp11296 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10511)


var ifres11290 Obj

if True == tmp11296 {
tmp11291 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11290 = tmp11291


} else {
tmp11292 := Call(__e, PrimFunc(symshen_4_5_1out), let__10511)


let__10512 := tmp11292
_ = let__10512

tmp11293 := Call(__e, PrimFunc(symshen_4in_1_6), let__10511)


let__10513 := tmp11293
_ = let__10513

tmp11294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10509, let__10512)
}
__typedArg0 := let__10509
__typedArg1 := let__10512
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11295 := Call(__e, PrimFunc(symshen_4comb), let__10513, tmp11294)


ifres11290 = tmp11295


}

ifres11285 = ifres11290


}

let__10507 := ifres11285
_ = let__10507

tmp11307 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10507)


if True == tmp11307 {
tmp11298 := Call(__e, PrimFunc(sym_5e_6), V1460)


let__10515 := tmp11298
_ = let__10515

tmp11303 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10515)


var ifres11299 Obj

if True == tmp11303 {
tmp11300 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11299 = tmp11300


} else {
tmp11301 := Call(__e, PrimFunc(symshen_4in_1_6), let__10515)


let__10516 := tmp11301
_ = let__10516

tmp11302 := Call(__e, PrimFunc(symshen_4comb), let__10516, Nil)


ifres11299 = tmp11302


}

let__10514 := ifres11299
_ = let__10514

tmp11305 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10514)


if True == tmp11305 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10514)
return
}


} else {
__e.Return(let__10507)
return
}


}, 1)

tmp11308 := Call(__e, ns2_1set, symshen_4_5bterms_6, tmp11283)


_ = tmp11308

tmp11309 := MakeNative(func(__e *ControlFlow) {
V1471 := __e.Get(1)
_ = V1471
tmp11310 := Call(__e, PrimFunc(symshen_4_5wildcard_6), V1471)


let__10518 := tmp11310
_ = let__10518

tmp11316 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10518)


var ifres11311 Obj

if True == tmp11316 {
tmp11312 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11311 = tmp11312


} else {
tmp11313 := Call(__e, PrimFunc(symshen_4_5_1out), let__10518)


let__10519 := tmp11313
_ = let__10519

tmp11314 := Call(__e, PrimFunc(symshen_4in_1_6), let__10518)


let__10520 := tmp11314
_ = let__10520

tmp11315 := Call(__e, PrimFunc(symshen_4comb), let__10520, let__10519)


ifres11311 = tmp11315


}

let__10517 := ifres11311
_ = let__10517

tmp11348 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10517)


if True == tmp11348 {
tmp11325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1471)
}
__typedArg0 := V1471
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11317 Obj

if True == tmp11325 {
tmp11318 := Call(__e, PrimFunc(symhead), V1471)


let__10522 := tmp11318
_ = let__10522

tmp11319 := Call(__e, PrimFunc(symtail), V1471)


let__10523 := tmp11319
_ = let__10523

tmp11323 := Call(__e, PrimFunc(symatom_2), let__10522)


var ifres11320 Obj

if True == tmp11323 {
tmp11321 := Call(__e, PrimFunc(symshen_4comb), let__10523, let__10522)


ifres11320 = tmp11321


} else {
tmp11322 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11320 = tmp11322


}

ifres11317 = ifres11320


} else {
tmp11324 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11317 = tmp11324


}

let__10521 := ifres11317
_ = let__10521

tmp11346 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10521)


if True == tmp11346 {
tmp11342 := Call(__e, PrimFunc(symshen_4ccons_2), V1471)


var ifres11326 Obj

if True == tmp11342 {
tmp11327 := Call(__e, PrimFunc(symhead), V1471)


let__10525 := tmp11327
_ = let__10525

tmp11328 := Call(__e, PrimFunc(symtail), V1471)


let__10526 := tmp11328
_ = let__10526

tmp11329 := Call(__e, PrimFunc(symshen_4_5bterms_6), let__10525)


let__10527 := tmp11329
_ = let__10527

tmp11340 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10527)


var ifres11330 Obj

if True == tmp11340 {
tmp11331 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11330 = tmp11331


} else {
tmp11332 := Call(__e, PrimFunc(symshen_4_5_1out), let__10527)


let__10528 := tmp11332
_ = let__10528

tmp11333 := Call(__e, PrimFunc(symshen_4in_1_6), let__10527)


let__10529 := tmp11333
_ = let__10529

tmp11334 := Call(__e, PrimFunc(sym_5end_6), let__10529)


let__10530 := tmp11334
_ = let__10530

tmp11339 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10530)


var ifres11335 Obj

if True == tmp11339 {
tmp11336 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11335 = tmp11336


} else {
tmp11337 := Call(__e, PrimFunc(symshen_4in_1_6), let__10530)


let__10531 := tmp11337
_ = let__10531

tmp11338 := Call(__e, PrimFunc(symshen_4comb), let__10526, let__10528)


ifres11335 = tmp11338


}

ifres11330 = ifres11335


}

ifres11326 = ifres11330


} else {
tmp11341 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11326 = tmp11341


}

let__10524 := ifres11326
_ = let__10524

tmp11344 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10524)


if True == tmp11344 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10524)
return
}


} else {
__e.Return(let__10521)
return
}


} else {
__e.Return(let__10517)
return
}


}, 1)

tmp11349 := Call(__e, ns2_1set, symshen_4_5bterm_6, tmp11309)


_ = tmp11349

tmp11350 := MakeNative(func(__e *ControlFlow) {
V1487 := __e.Get(1)
_ = V1487
tmp11360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1487)
}
__typedArg0 := V1487
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11351 Obj

if True == tmp11360 {
tmp11352 := Call(__e, PrimFunc(symhead), V1487)


let__10533 := tmp11352
_ = let__10533

tmp11353 := Call(__e, PrimFunc(symtail), V1487)


let__10534 := tmp11353
_ = let__10534

tmp11358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10533, sym__)
}
__typedArg0 := let__10533
__typedArg1 := sym__
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11354 Obj

if True == tmp11358 {
tmp11355 := Call(__e, PrimFunc(symgensym), symY)


tmp11356 := Call(__e, PrimFunc(symshen_4comb), let__10534, tmp11355)


ifres11354 = tmp11356


} else {
tmp11357 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11354 = tmp11357


}

ifres11351 = ifres11354


} else {
tmp11359 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11351 = tmp11359


}

let__10532 := ifres11351
_ = let__10532

tmp11362 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10532)


if True == tmp11362 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10532)
return
}


}, 1)

tmp11363 := Call(__e, ns2_1set, symshen_4_5wildcard_6, tmp11350)


_ = tmp11363

tmp11364 := MakeNative(func(__e *ControlFlow) {
V1491 := __e.Get(1)
_ = V1491
tmp11373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1491)
}
__typedArg0 := V1491
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11365 Obj

if True == tmp11373 {
tmp11366 := Call(__e, PrimFunc(symhead), V1491)


let__10536 := tmp11366
_ = let__10536

tmp11367 := Call(__e, PrimFunc(symtail), V1491)


let__10537 := tmp11367
_ = let__10537

tmp11371 := Call(__e, PrimFunc(symshen_4semicolon_2), let__10536)


var ifres11368 Obj

if True == tmp11371 {
tmp11369 := Call(__e, PrimFunc(symshen_4comb), let__10537, let__10536)


ifres11368 = tmp11369


} else {
tmp11370 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11368 = tmp11370


}

ifres11365 = ifres11368


} else {
tmp11372 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11365 = tmp11372


}

let__10535 := ifres11365
_ = let__10535

tmp11375 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__10535)


if True == tmp11375 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__10535)
return
}


}, 1)

tmp11376 := Call(__e, ns2_1set, symshen_4_5sc_6, tmp11364)


_ = tmp11376

tmp11377 := MakeNative(func(__e *ControlFlow) {
V1495 := __e.Get(1)
_ = V1495
V1496 := __e.Get(2)
_ = V1496
tmp11378 := Call(__e, PrimFunc(symgensym), symB)


let__10538 := tmp11378
_ = let__10538

tmp11379 := Call(__e, PrimFunc(symgensym), symL)


let__10539 := tmp11379
_ = let__10539

tmp11380 := Call(__e, PrimFunc(symgensym), symK)


let__10540 := tmp11380
_ = let__10540

tmp11381 := Call(__e, PrimFunc(symgensym), symC)


let__10541 := tmp11381
_ = let__10541

tmp11382 := Call(__e, PrimFunc(symshen_4prolog_1parameters), V1496)


let__10542 := tmp11382
_ = let__10542

tmp11383 := Call(__e, PrimFunc(symshen_4hascut_2), V1496)


let__10543 := tmp11383
_ = let__10543

tmp11384 := Call(__e, PrimFunc(symshen_4prolog_1fbody), V1496, let__10542, let__10538, let__10539, let__10540, let__10541, let__10543)


let__10544 := tmp11384
_ = let__10544

var ifres11385 Obj

if True == let__10543 {
tmp11386 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(1), Nil)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10540, tmp11386)
}
__typedArg0 := let__10540
__typedArg1 := tmp11386
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_7, tmp11387)
}
__typedArg0 := sym_7
__typedArg1 := tmp11387
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10544, Nil)
}
__typedArg0 := let__10544
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11388, tmp11389)
}
__typedArg0 := tmp11388
__typedArg1 := tmp11389
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10540, tmp11390)
}
__typedArg0 := let__10540
__typedArg1 := tmp11390
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11391)
}
__typedArg0 := symlet
__typedArg1 := tmp11391
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

ifres11385 = tmp11392


} else {
ifres11385 = let__10544


}

let__10545 := ifres11385
_ = let__10545

tmp11393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, Nil)
}
__typedArg0 := sym_1_6
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10541, tmp11393)
}
__typedArg0 := let__10541
__typedArg1 := tmp11393
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10540, tmp11394)
}
__typedArg0 := let__10540
__typedArg1 := tmp11394
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10539, tmp11395)
}
__typedArg0 := let__10539
__typedArg1 := tmp11395
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10538, tmp11396)
}
__typedArg0 := let__10538
__typedArg1 := tmp11396
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10545, Nil)
}
__typedArg0 := let__10545
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11399 := Call(__e, PrimFunc(symappend), tmp11397, tmp11398)


tmp11400 := Call(__e, PrimFunc(symappend), let__10542, tmp11399)


tmp11401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1495, tmp11400)
}
__typedArg0 := V1495
__typedArg1 := tmp11400
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefine, tmp11401)
}
__typedArg0 := symdefine
__typedArg1 := tmp11401
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__10546 := tmp11402
_ = let__10546

__e.Return(let__10546)
return


}, 2)

tmp11403 := Call(__e, ns2_1set, symshen_4horn_1clause_1procedure, tmp11377)


_ = tmp11403

tmp11404 := MakeNative(func(__e *ControlFlow) {
V1508 := __e.Get(1)
_ = V1508
tmp11414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_b, V1508)
}
__typedArg0 := sym_b
__typedArg1 := V1508
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11414 {
__e.Return(True)
return
} else {
tmp11412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1508)
}
__typedArg0 := V1508
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11412 {
tmp11409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1508)
}
__typedArg0 := V1508
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11410 := Call(__e, PrimFunc(symshen_4hascut_2), tmp11409)


if True == tmp11410 {
__e.Return(True)
return
} else {
tmp11406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1508)
}
__typedArg0 := V1508
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11407 := Call(__e, PrimFunc(symshen_4hascut_2), tmp11406)


if True == tmp11407 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


} else {
__e.Return(False)
return
}


}


}, 1)

tmp11415 := Call(__e, ns2_1set, symshen_4hascut_2, tmp11404)


_ = tmp11415

tmp11416 := MakeNative(func(__e *ControlFlow) {
V1513 := __e.Get(1)
_ = V1513
tmp11425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1513)
}
__typedArg0 := V1513
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11421 Obj

if True == tmp11425 {
tmp11423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1513)
}
__typedArg0 := V1513
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11423)
}
__typedArg0 := tmp11423
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11422 Obj

if True == tmp11424 {
ifres11422 = True


} else {
ifres11422 = False


}

ifres11421 = ifres11422


} else {
ifres11421 = False


}

if True == ifres11421 {
tmp11417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1513)
}
__typedArg0 := V1513
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11417)
}
__typedArg0 := tmp11417
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11419 := Call(__e, PrimFunc(symlength), tmp11418)


__e.TailApply(PrimFunc(symshen_4parameters), tmp11419)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.prolog-parameters"))
}
__typedArg0 := MakeString("partial function shen.prolog-parameters")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp11426 := Call(__e, ns2_1set, symshen_4prolog_1parameters, tmp11416)


_ = tmp11426

tmp11427 := MakeNative(func(__e *ControlFlow) {
V1534 := __e.Get(1)
_ = V1534
V1535 := __e.Get(2)
_ = V1535
V1536 := __e.Get(3)
_ = V1536
V1537 := __e.Get(4)
_ = V1537
V1538 := __e.Get(5)
_ = V1538
V1539 := __e.Get(6)
_ = V1539
V1540 := __e.Get(7)
_ = V1540
tmp11517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1534)
}
__typedArg0 := Nil
__typedArg1 := V1534
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11514 Obj

if True == tmp11517 {
tmp11516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, V1540)
}
__typedArg0 := True
__typedArg1 := V1540
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11515 Obj

if True == tmp11516 {
ifres11515 = True


} else {
ifres11515 = False


}

ifres11514 = ifres11515


} else {
ifres11514 = False


}

if True == ifres11514 {
tmp11428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1538, Nil)
}
__typedArg0 := V1538
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1537, tmp11428)
}
__typedArg0 := V1537
__typedArg1 := tmp11428
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4unlock, tmp11429)
}
__typedArg0 := symshen_4unlock
__typedArg1 := tmp11429
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11490 Obj

if True == tmp11512 {
tmp11510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11510)
}
__typedArg0 := tmp11510
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11492 Obj

if True == tmp11511 {
tmp11507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11507)
}
__typedArg0 := tmp11507
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11508)
}
__typedArg0 := tmp11508
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11494 Obj

if True == tmp11509 {
tmp11503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11503)
}
__typedArg0 := tmp11503
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11504)
}
__typedArg0 := tmp11504
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11505)
}
__typedArg0 := Nil
__typedArg1 := tmp11505
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11496 Obj

if True == tmp11506 {
tmp11501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11501)
}
__typedArg0 := Nil
__typedArg1 := tmp11501
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11498 Obj

if True == tmp11502 {
tmp11500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(False, V1540)
}
__typedArg0 := False
__typedArg1 := V1540
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11499 Obj

if True == tmp11500 {
ifres11499 = True


} else {
ifres11499 = False


}

ifres11498 = ifres11499


} else {
ifres11498 = False


}

var ifres11497 Obj

if True == ifres11498 {
ifres11497 = True


} else {
ifres11497 = False


}

ifres11496 = ifres11497


} else {
ifres11496 = False


}

var ifres11495 Obj

if True == ifres11496 {
ifres11495 = True


} else {
ifres11495 = False


}

ifres11494 = ifres11495


} else {
ifres11494 = False


}

var ifres11493 Obj

if True == ifres11494 {
ifres11493 = True


} else {
ifres11493 = False


}

ifres11492 = ifres11493


} else {
ifres11492 = False


}

var ifres11491 Obj

if True == ifres11492 {
ifres11491 = True


} else {
ifres11491 = False


}

ifres11490 = ifres11491


} else {
ifres11490 = False


}

if True == ifres11490 {
tmp11430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11430)
}
__typedArg0 := tmp11430
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11432)
}
__typedArg0 := tmp11432
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11433)
}
__typedArg0 := tmp11433
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11435 := Call(__e, PrimFunc(symshen_4continue), tmp11431, tmp11434, V1536, V1537, V1538, V1539)


let__10547 := tmp11435
_ = let__10547

tmp11436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1537, Nil)
}
__typedArg0 := V1537
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4unlocked_2, tmp11436)
}
__typedArg0 := symshen_4unlocked_2
__typedArg1 := tmp11436
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11438)
}
__typedArg0 := tmp11438
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11440 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11439, V1535, V1536, let__10547)


tmp11441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11440, tmp11441)
}
__typedArg0 := tmp11440
__typedArg1 := tmp11441
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11437, tmp11442)
}
__typedArg0 := tmp11437
__typedArg1 := tmp11442
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11443)
}
__typedArg0 := symif
__typedArg1 := tmp11443
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11473 Obj

if True == tmp11488 {
tmp11486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11486)
}
__typedArg0 := tmp11486
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11475 Obj

if True == tmp11487 {
tmp11483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11483)
}
__typedArg0 := tmp11483
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11484)
}
__typedArg0 := tmp11484
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11477 Obj

if True == tmp11485 {
tmp11479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11479)
}
__typedArg0 := tmp11479
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11480)
}
__typedArg0 := tmp11480
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11481)
}
__typedArg0 := Nil
__typedArg1 := tmp11481
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11478 Obj

if True == tmp11482 {
ifres11478 = True


} else {
ifres11478 = False


}

ifres11477 = ifres11478


} else {
ifres11477 = False


}

var ifres11476 Obj

if True == ifres11477 {
ifres11476 = True


} else {
ifres11476 = False


}

ifres11475 = ifres11476


} else {
ifres11475 = False


}

var ifres11474 Obj

if True == ifres11475 {
ifres11474 = True


} else {
ifres11474 = False


}

ifres11473 = ifres11474


} else {
ifres11473 = False


}

if True == ifres11473 {
tmp11444 := Call(__e, PrimFunc(symgensym), symC)


let__10548 := tmp11444
_ = let__10548

tmp11445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11445)
}
__typedArg0 := tmp11445
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11447)
}
__typedArg0 := tmp11447
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11448)
}
__typedArg0 := tmp11448
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11450 := Call(__e, PrimFunc(symshen_4continue), tmp11446, tmp11449, V1536, V1537, V1538, V1539)


let__10549 := tmp11450
_ = let__10549

tmp11451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1537, Nil)
}
__typedArg0 := V1537
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4unlocked_2, tmp11451)
}
__typedArg0 := symshen_4unlocked_2
__typedArg1 := tmp11451
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11453)
}
__typedArg0 := tmp11453
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11455 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11454, V1535, V1536, let__10549)


tmp11456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11455, tmp11456)
}
__typedArg0 := tmp11455
__typedArg1 := tmp11456
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11452, tmp11457)
}
__typedArg0 := tmp11452
__typedArg1 := tmp11457
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11458)
}
__typedArg0 := symif
__typedArg1 := tmp11458
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10548, tmp11460)
}
__typedArg0 := let__10548
__typedArg1 := tmp11460
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp11461)
}
__typedArg0 := sym_a
__typedArg1 := tmp11461
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1534)
}
__typedArg0 := V1534
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11464 := Call(__e, PrimFunc(symshen_4prolog_1fbody), tmp11463, V1535, V1536, V1537, V1538, V1539, V1540)


tmp11465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10548, Nil)
}
__typedArg0 := let__10548
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11464, tmp11465)
}
__typedArg0 := tmp11464
__typedArg1 := tmp11465
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11462, tmp11466)
}
__typedArg0 := tmp11462
__typedArg1 := tmp11466
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp11467)
}
__typedArg0 := symif
__typedArg1 := tmp11467
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11468, Nil)
}
__typedArg0 := tmp11468
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11459, tmp11469)
}
__typedArg0 := tmp11459
__typedArg1 := tmp11469
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10548, tmp11470)
}
__typedArg0 := let__10548
__typedArg1 := tmp11470
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp11471)
}
__typedArg0 := symlet
__typedArg1 := tmp11471
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.prolog-fbody"))
}
__typedArg0 := MakeString("implementation error in shen.prolog-fbody")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 7)

tmp11518 := Call(__e, ns2_1set, symshen_4prolog_1fbody, tmp11427)


_ = tmp11518

tmp11519 := MakeNative(func(__e *ControlFlow) {
V1544 := __e.Get(1)
_ = V1544
V1545 := __e.Get(2)
_ = V1545
tmp11524 := Call(__e, PrimFunc(symshen_4locked_2), V1544)


var ifres11521 Obj

if True == tmp11524 {
tmp11523 := Call(__e, PrimFunc(symshen_4fits_2), V1545, V1544)


var ifres11522 Obj

if True == tmp11523 {
ifres11522 = True


} else {
ifres11522 = False


}

ifres11521 = ifres11522


} else {
ifres11521 = False


}

if True == ifres11521 {
__e.TailApply(PrimFunc(symshen_4openlock), V1544)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp11525 := Call(__e, ns2_1set, symshen_4unlock, tmp11519)


_ = tmp11525

tmp11526 := MakeNative(func(__e *ControlFlow) {
V1546 := __e.Get(1)
_ = V1546
tmp11527 := Call(__e, PrimFunc(symshen_4unlocked_2), V1546)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp11527)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp11527
return Call(__e, PrimFunc(symnot), __typedArg0)
})())
return


}, 1)

tmp11528 := Call(__e, ns2_1set, symshen_4locked_2, tmp11526)


_ = tmp11528

tmp11529 := MakeNative(func(__e *ControlFlow) {
V1547 := __e.Get(1)
_ = V1547
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1547, MakeInteger(1))
}
__typedArg0 := V1547
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp11530 := Call(__e, ns2_1set, symshen_4unlocked_2, tmp11529)


_ = tmp11530

tmp11531 := MakeNative(func(__e *ControlFlow) {
V1548 := __e.Get(1)
_ = V1548
tmp11532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V1548, MakeInteger(1), True)
}
__typedArg0 := V1548
__typedArg1 := MakeInteger(1)
__typedArg2 := True
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

_ = tmp11532

__e.Return(False)
return


}, 1)

tmp11533 := Call(__e, ns2_1set, symshen_4openlock, tmp11531)


_ = tmp11533

tmp11534 := MakeNative(func(__e *ControlFlow) {
V1549 := __e.Get(1)
_ = V1549
V1550 := __e.Get(2)
_ = V1550
tmp11535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1550, MakeInteger(2))
}
__typedArg0 := V1550
__typedArg1 := MakeInteger(2)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1549, tmp11535)
}
__typedArg0 := V1549
__typedArg1 := tmp11535
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return


}, 2)

tmp11536 := Call(__e, ns2_1set, symshen_4fits_2, tmp11534)


_ = tmp11536

tmp11537 := MakeNative(func(__e *ControlFlow) {
V1553 := __e.Get(1)
_ = V1553
V1554 := __e.Get(2)
_ = V1554
V1555 := __e.Get(3)
_ = V1555
V1556 := __e.Get(4)
_ = V1556
tmp11538 := Call(__e, PrimFunc(symthaw), V1556)


let__10550 := tmp11538
_ = let__10550

tmp11543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10550, False)
}
__typedArg0 := let__10550
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11540 Obj

if True == tmp11543 {
tmp11542 := Call(__e, PrimFunc(symshen_4unlocked_2), V1554)


var ifres11541 Obj

if True == tmp11542 {
ifres11541 = True


} else {
ifres11541 = False


}

ifres11540 = ifres11541


} else {
ifres11540 = False


}

if True == ifres11540 {
__e.TailApply(PrimFunc(symshen_4lock), V1555, V1554)
return
} else {
__e.Return(let__10550)
return
}


}, 4)

tmp11544 := Call(__e, ns2_1set, symshen_4cut, tmp11537)


_ = tmp11544

tmp11545 := MakeNative(func(__e *ControlFlow) {
V1558 := __e.Get(1)
_ = V1558
V1559 := __e.Get(2)
_ = V1559
tmp11546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V1559, MakeInteger(1), False)
}
__typedArg0 := V1559
__typedArg1 := MakeInteger(1)
__typedArg2 := False
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__10551 := tmp11546
_ = let__10551

tmp11547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V1559, MakeInteger(2), V1558)
}
__typedArg0 := V1559
__typedArg1 := MakeInteger(2)
__typedArg2 := V1558
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__10552 := tmp11547
_ = let__10552

__e.Return(False)
return


}, 2)

tmp11548 := Call(__e, ns2_1set, symshen_4lock, tmp11545)


_ = tmp11548

tmp11549 := MakeNative(func(__e *ControlFlow) {
V1562 := __e.Get(1)
_ = V1562
V1563 := __e.Get(2)
_ = V1563
V1564 := __e.Get(3)
_ = V1564
V1565 := __e.Get(4)
_ = V1565
V1566 := __e.Get(5)
_ = V1566
V1567 := __e.Get(6)
_ = V1567
tmp11550 := Call(__e, PrimFunc(symshen_4extract_1vars), V1562)


let__10553 := tmp11550
_ = let__10553

tmp11551 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), V1563)


let__10554 := tmp11551
_ = let__10554

tmp11552 := Call(__e, PrimFunc(symdifference), let__10554, let__10553)


let__10555 := tmp11552
_ = let__10555

tmp11553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4incinfs, Nil)
}
__typedArg0 := symshen_4incinfs
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11554 := Call(__e, PrimFunc(symshen_4compile_1body), V1563, V1564, V1565, V1566, V1567)


tmp11555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11554, Nil)
}
__typedArg0 := tmp11554
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11553, tmp11555)
}
__typedArg0 := tmp11553
__typedArg1 := tmp11555
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdo, tmp11556)
}
__typedArg0 := symdo
__typedArg1 := tmp11556
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__10556 := tmp11557
_ = let__10556

__e.TailApply(PrimFunc(symshen_4stpart), let__10555, let__10556, V1564)
return


}, 6)

tmp11558 := Call(__e, ns2_1set, symshen_4continue, tmp11549)


_ = tmp11558

tmp11559 := MakeNative(func(__e *ControlFlow) {
V1574 := __e.Get(1)
_ = V1574
tmp11594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1574)
}
__typedArg0 := V1574
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11575 Obj

if True == tmp11594 {
tmp11592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1574)
}
__typedArg0 := V1574
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp11592)
}
__typedArg0 := symlambda
__typedArg1 := tmp11592
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11577 Obj

if True == tmp11593 {
tmp11590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1574)
}
__typedArg0 := V1574
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11590)
}
__typedArg0 := tmp11590
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11579 Obj

if True == tmp11591 {
tmp11587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1574)
}
__typedArg0 := V1574
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11587)
}
__typedArg0 := tmp11587
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11588)
}
__typedArg0 := tmp11588
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11581 Obj

if True == tmp11589 {
tmp11583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1574)
}
__typedArg0 := V1574
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11583)
}
__typedArg0 := tmp11583
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11584)
}
__typedArg0 := tmp11584
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11585)
}
__typedArg0 := Nil
__typedArg1 := tmp11585
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11582 Obj

if True == tmp11586 {
ifres11582 = True


} else {
ifres11582 = False


}

ifres11581 = ifres11582


} else {
ifres11581 = False


}

var ifres11580 Obj

if True == ifres11581 {
ifres11580 = True


} else {
ifres11580 = False


}

ifres11579 = ifres11580


} else {
ifres11579 = False


}

var ifres11578 Obj

if True == ifres11579 {
ifres11578 = True


} else {
ifres11578 = False


}

ifres11577 = ifres11578


} else {
ifres11577 = False


}

var ifres11576 Obj

if True == ifres11577 {
ifres11576 = True


} else {
ifres11576 = False


}

ifres11575 = ifres11576


} else {
ifres11575 = False


}

if True == ifres11575 {
tmp11560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1574)
}
__typedArg0 := V1574
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11560)
}
__typedArg0 := tmp11560
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1574)
}
__typedArg0 := V1574
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11562)
}
__typedArg0 := tmp11562
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11563)
}
__typedArg0 := tmp11563
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11565 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp11564)


__e.TailApply(PrimFunc(symremove), tmp11561, tmp11565)
return


} else {
tmp11573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1574)
}
__typedArg0 := V1574
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11573 {
tmp11566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1574)
}
__typedArg0 := V1574
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11567 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp11566)


tmp11568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1574)
}
__typedArg0 := V1574
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11569 := Call(__e, PrimFunc(symshen_4extract_1free_1vars), tmp11568)


__e.TailApply(PrimFunc(symunion), tmp11567, tmp11569)
return


} else {
tmp11571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V1574)
}
__typedArg0 := V1574
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp11571 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1574, Nil)
}
__typedArg0 := V1574
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
} else {
__e.Return(Nil)
return
}


}


}


}, 1)

tmp11595 := Call(__e, ns2_1set, symshen_4extract_1free_1vars, tmp11559)


_ = tmp11595

tmp11596 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
__selftop:
V1591 := __self1
_ = V1591
V1592 := __self2
_ = V1592
V1593 := __self3
_ = V1593
V1594 := __self4
_ = V1594
V1595 := __self5
_ = V1595
tmp11630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1591)
}
__typedArg0 := Nil
__typedArg1 := V1591
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11630 {
tmp11597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1595, Nil)
}
__typedArg0 := V1595
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp11597)
}
__typedArg0 := symthaw
__typedArg1 := tmp11597
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1591)
}
__typedArg0 := V1591
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11624 Obj

if True == tmp11628 {
tmp11626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1591)
}
__typedArg0 := V1591
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_b, tmp11626)
}
__typedArg0 := sym_b
__typedArg1 := tmp11626
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11625 Obj

if True == tmp11627 {
ifres11625 = True


} else {
ifres11625 = False


}

ifres11624 = ifres11625


} else {
ifres11624 = False


}

if True == ifres11624 {
tmp11598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4cut, Nil)
}
__typedArg0 := symshen_4cut
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1591)
}
__typedArg0 := V1591
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11598, tmp11599)
}
__typedArg0 := tmp11598
__typedArg1 := tmp11599
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4compile_1body) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp11600, V1592, V1593, V1594, V1595
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4compile_1body), tmp11600, V1592, V1593, V1594, V1595)
return


} else {
tmp11622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1591)
}
__typedArg0 := V1591
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11618 Obj

if True == tmp11622 {
tmp11620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1591)
}
__typedArg0 := V1591
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11620)
}
__typedArg0 := Nil
__typedArg1 := tmp11620
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11619 Obj

if True == tmp11621 {
ifres11619 = True


} else {
ifres11619 = False


}

ifres11618 = ifres11619


} else {
ifres11618 = False


}

if True == ifres11618 {
tmp11601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1591)
}
__typedArg0 := V1591
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11602 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11601, V1592)


tmp11603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1595, Nil)
}
__typedArg0 := V1595
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1594, tmp11603)
}
__typedArg0 := V1594
__typedArg1 := tmp11603
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1593, tmp11604)
}
__typedArg0 := V1593
__typedArg1 := tmp11604
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1592, tmp11605)
}
__typedArg0 := V1592
__typedArg1 := tmp11605
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symappend), tmp11602, tmp11606)
return


} else {
tmp11616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1591)
}
__typedArg0 := V1591
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11616 {
tmp11607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1591)
}
__typedArg0 := V1591
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11608 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11607, V1592)


let__10557 := tmp11608
_ = let__10557

tmp11609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1591)
}
__typedArg0 := V1591
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11610 := Call(__e, PrimFunc(symshen_4freeze_1literals), tmp11609, V1592, V1593, V1594, V1595)


tmp11611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11610, Nil)
}
__typedArg0 := tmp11610
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1594, tmp11611)
}
__typedArg0 := V1594
__typedArg1 := tmp11611
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1593, tmp11612)
}
__typedArg0 := V1593
__typedArg1 := tmp11612
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1592, tmp11613)
}
__typedArg0 := V1592
__typedArg1 := tmp11613
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symappend), let__10557, tmp11614)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.compile-fbody"))
}
__typedArg0 := MakeString("implementation error in shen.compile-fbody")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}


}, 5)

tmp11631 := Call(__e, ns2_1set, symshen_4compile_1body, tmp11596)


_ = tmp11631

tmp11632 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
__selftop:
V1613 := __self1
_ = V1613
V1614 := __self2
_ = V1614
V1615 := __self3
_ = V1615
V1616 := __self4
_ = V1616
V1617 := __self5
_ = V1617
tmp11655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1613)
}
__typedArg0 := Nil
__typedArg1 := V1613
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11655 {
__e.Return(V1617)
return
} else {
tmp11653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1613)
}
__typedArg0 := V1613
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11649 Obj

if True == tmp11653 {
tmp11651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1613)
}
__typedArg0 := V1613
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_b, tmp11651)
}
__typedArg0 := sym_b
__typedArg1 := tmp11651
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11650 Obj

if True == tmp11652 {
ifres11650 = True


} else {
ifres11650 = False


}

ifres11649 = ifres11650


} else {
ifres11649 = False


}

if True == ifres11649 {
tmp11633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4cut, Nil)
}
__typedArg0 := symshen_4cut
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1613)
}
__typedArg0 := V1613
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11633, tmp11634)
}
__typedArg0 := tmp11633
__typedArg1 := tmp11634
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4freeze_1literals) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp11635, V1614, V1615, V1616, V1617
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4freeze_1literals), tmp11635, V1614, V1615, V1616, V1617)
return


} else {
tmp11647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1613)
}
__typedArg0 := V1613
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11647 {
tmp11636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1613)
}
__typedArg0 := V1613
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11637 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11636, V1614)


let__10558 := tmp11637
_ = let__10558

tmp11638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1613)
}
__typedArg0 := V1613
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11639 := Call(__e, PrimFunc(symshen_4freeze_1literals), tmp11638, V1614, V1615, V1616, V1617)


tmp11640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11639, Nil)
}
__typedArg0 := tmp11639
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1616, tmp11640)
}
__typedArg0 := V1616
__typedArg1 := tmp11640
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1615, tmp11641)
}
__typedArg0 := V1615
__typedArg1 := tmp11641
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1614, tmp11642)
}
__typedArg0 := V1614
__typedArg1 := tmp11642
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11644 := Call(__e, PrimFunc(symappend), let__10558, tmp11643)


tmp11645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11644, Nil)
}
__typedArg0 := tmp11644
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp11645)
}
__typedArg0 := symfreeze
__typedArg1 := tmp11645
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.freeze-literals"))
}
__typedArg0 := MakeString("implementation error in shen.freeze-literals")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 5)

tmp11656 := Call(__e, ns2_1set, symshen_4freeze_1literals, tmp11632)


_ = tmp11656

tmp11657 := MakeNative(func(__e *ControlFlow) {
V1623 := __e.Get(1)
_ = V1623
V1624 := __e.Get(2)
_ = V1624
tmp11672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1623)
}
__typedArg0 := V1623
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11668 Obj

if True == tmp11672 {
tmp11670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1623)
}
__typedArg0 := V1623
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfork, tmp11670)
}
__typedArg0 := symfork
__typedArg1 := tmp11670
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11669 Obj

if True == tmp11671 {
ifres11669 = True


} else {
ifres11669 = False


}

ifres11668 = ifres11669


} else {
ifres11668 = False


}

if True == ifres11668 {
tmp11658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1623)
}
__typedArg0 := V1623
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11659 := Call(__e, PrimFunc(symshen_4deref_1forked_1literals), tmp11658, V1624)


tmp11660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11659, Nil)
}
__typedArg0 := tmp11659
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfork, tmp11660)
}
__typedArg0 := symfork
__typedArg1 := tmp11660
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1623)
}
__typedArg0 := V1623
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11666 {
tmp11661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1623)
}
__typedArg0 := V1623
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11662 := MakeNative(func(__e *ControlFlow) {
Z1625 := __e.Get(1)
_ = Z1625
__e.TailApply(PrimFunc(symshen_4function_1calls), Z1625, V1624)
return
}, 1)

tmp11663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1623)
}
__typedArg0 := V1623
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11664 := Call(__e, PrimFunc(symmap), tmp11662, tmp11663)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11661, tmp11664)
}
__typedArg0 := tmp11661
__typedArg1 := tmp11664
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.deref-calls"))
}
__typedArg0 := MakeString("implementation error in shen.deref-calls")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp11673 := Call(__e, ns2_1set, symshen_4deref_1calls, tmp11657)


_ = tmp11673

tmp11674 := MakeNative(func(__e *ControlFlow) {
V1632 := __e.Get(1)
_ = V1632
V1633 := __e.Get(2)
_ = V1633
tmp11684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1632)
}
__typedArg0 := Nil
__typedArg1 := V1632
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp11684 {
__e.Return(Nil)
return
} else {
tmp11682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1632)
}
__typedArg0 := V1632
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11682 {
tmp11675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1632)
}
__typedArg0 := V1632
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11676 := Call(__e, PrimFunc(symshen_4deref_1calls), tmp11675, V1633)


tmp11677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1632)
}
__typedArg0 := V1632
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11678 := Call(__e, PrimFunc(symshen_4deref_1forked_1literals), tmp11677, V1633)


tmp11679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11678, Nil)
}
__typedArg0 := tmp11678
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11676, tmp11679)
}
__typedArg0 := tmp11676
__typedArg1 := tmp11679
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp11680)
}
__typedArg0 := symcons
__typedArg1 := tmp11680
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("fork requires a list of literals\n"))
}
__typedArg0 := MakeString("fork requires a list of literals\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp11685 := Call(__e, ns2_1set, symshen_4deref_1forked_1literals, tmp11674)


_ = tmp11685

tmp11686 := MakeNative(func(__e *ControlFlow) {
V1636 := __e.Get(1)
_ = V1636
V1637 := __e.Get(2)
_ = V1637
tmp11718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1636)
}
__typedArg0 := V1636
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11699 Obj

if True == tmp11718 {
tmp11716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1636)
}
__typedArg0 := V1636
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp11716)
}
__typedArg0 := symcons
__typedArg1 := tmp11716
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11701 Obj

if True == tmp11717 {
tmp11714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1636)
}
__typedArg0 := V1636
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11714)
}
__typedArg0 := tmp11714
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11703 Obj

if True == tmp11715 {
tmp11711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1636)
}
__typedArg0 := V1636
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11711)
}
__typedArg0 := tmp11711
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11712)
}
__typedArg0 := tmp11712
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11705 Obj

if True == tmp11713 {
tmp11707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1636)
}
__typedArg0 := V1636
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11707)
}
__typedArg0 := tmp11707
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11708)
}
__typedArg0 := tmp11708
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11709)
}
__typedArg0 := Nil
__typedArg1 := tmp11709
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11706 Obj

if True == tmp11710 {
ifres11706 = True


} else {
ifres11706 = False


}

ifres11705 = ifres11706


} else {
ifres11705 = False


}

var ifres11704 Obj

if True == ifres11705 {
ifres11704 = True


} else {
ifres11704 = False


}

ifres11703 = ifres11704


} else {
ifres11703 = False


}

var ifres11702 Obj

if True == ifres11703 {
ifres11702 = True


} else {
ifres11702 = False


}

ifres11701 = ifres11702


} else {
ifres11701 = False


}

var ifres11700 Obj

if True == ifres11701 {
ifres11700 = True


} else {
ifres11700 = False


}

ifres11699 = ifres11700


} else {
ifres11699 = False


}

if True == ifres11699 {
tmp11687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1636)
}
__typedArg0 := V1636
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11687)
}
__typedArg0 := tmp11687
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11689 := Call(__e, PrimFunc(symshen_4function_1calls), tmp11688, V1637)


tmp11690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1636)
}
__typedArg0 := V1636
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11690)
}
__typedArg0 := tmp11690
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11691)
}
__typedArg0 := tmp11691
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11693 := Call(__e, PrimFunc(symshen_4function_1calls), tmp11692, V1637)


tmp11694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11693, Nil)
}
__typedArg0 := tmp11693
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11689, tmp11694)
}
__typedArg0 := tmp11689
__typedArg1 := tmp11694
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp11695)
}
__typedArg0 := symcons
__typedArg1 := tmp11695
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1636)
}
__typedArg0 := V1636
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11697 {
__e.TailApply(PrimFunc(symshen_4deref_1terms), V1636, V1637, Nil)
return
} else {
__e.Return(V1636)
return
}


}


}, 2)

tmp11719 := Call(__e, ns2_1set, symshen_4function_1calls, tmp11686)


_ = tmp11719

tmp11720 := MakeNative(func(__e *ControlFlow) {
V1646 := __e.Get(1)
_ = V1646
V1647 := __e.Get(2)
_ = V1647
V1648 := __e.Get(3)
_ = V1648
tmp11814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11801 Obj

if True == tmp11814 {
tmp11812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), tmp11812)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp11812
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11803 Obj

if True == tmp11813 {
tmp11810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11810)
}
__typedArg0 := tmp11810
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11805 Obj

if True == tmp11811 {
tmp11807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11807)
}
__typedArg0 := tmp11807
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11808)
}
__typedArg0 := Nil
__typedArg1 := tmp11808
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11806 Obj

if True == tmp11809 {
ifres11806 = True


} else {
ifres11806 = False


}

ifres11805 = ifres11806


} else {
ifres11805 = False


}

var ifres11804 Obj

if True == ifres11805 {
ifres11804 = True


} else {
ifres11804 = False


}

ifres11803 = ifres11804


} else {
ifres11803 = False


}

var ifres11802 Obj

if True == ifres11803 {
ifres11802 = True


} else {
ifres11802 = False


}

ifres11801 = ifres11802


} else {
ifres11801 = False


}

if True == ifres11801 {
tmp11727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11727)
}
__typedArg0 := tmp11727
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp11728)
}
__typedArg0 := tmp11728
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp11729 {
tmp11721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11721)
}
__typedArg0 := tmp11721
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp11722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11722)
}
__typedArg0 := tmp11722
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11724 := Call(__e, PrimFunc(symshen_4app), tmp11723, MakeString("\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attempt to optimise a non-variable "))
__typedS1, __typedOK1 := TypedString(tmp11724)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attempt to optimise a non-variable ")
__typedArg1 := tmp11724
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attempt to optimise a non-variable "))
__typedS1, __typedOK1 := TypedString(tmp11724)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attempt to optimise a non-variable ")
__typedArg1 := tmp11724
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


} else {
tmp11799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11786 Obj

if True == tmp11799 {
tmp11797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(1), tmp11797)
}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp11797
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11788 Obj

if True == tmp11798 {
tmp11795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11795)
}
__typedArg0 := tmp11795
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11790 Obj

if True == tmp11796 {
tmp11792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11792)
}
__typedArg0 := tmp11792
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11793)
}
__typedArg0 := Nil
__typedArg1 := tmp11793
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11791 Obj

if True == tmp11794 {
ifres11791 = True


} else {
ifres11791 = False


}

ifres11790 = ifres11791


} else {
ifres11790 = False


}

var ifres11789 Obj

if True == ifres11790 {
ifres11789 = True


} else {
ifres11789 = False


}

ifres11788 = ifres11789


} else {
ifres11788 = False


}

var ifres11787 Obj

if True == ifres11788 {
ifres11787 = True


} else {
ifres11787 = False


}

ifres11786 = ifres11787


} else {
ifres11786 = False


}

if True == ifres11786 {
tmp11739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11739)
}
__typedArg0 := tmp11739
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp11740)
}
__typedArg0 := tmp11740
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp11741 {
tmp11730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11730)
}
__typedArg0 := tmp11730
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1647, Nil)
}
__typedArg0 := V1647
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11731, tmp11732)
}
__typedArg0 := tmp11731
__typedArg1 := tmp11732
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp11733)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp11733
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11734)
}
__typedArg0 := tmp11734
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11736 := Call(__e, PrimFunc(symshen_4app), tmp11735, MakeString("\n"), symshen_4s)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attempt to optimise a non-variable "))
__typedS1, __typedOK1 := TypedString(tmp11736)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attempt to optimise a non-variable ")
__typedArg1 := tmp11736
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("attempt to optimise a non-variable "))
__typedS1, __typedOK1 := TypedString(tmp11736)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("attempt to optimise a non-variable ")
__typedArg1 := tmp11736
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


} else {
tmp11783 := Call(__e, PrimFunc(symelement_2), V1646, V1648)


tmp11784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp11783)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp11783
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres11780 Obj

if True == tmp11784 {
tmp11782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres11781 Obj

if True == tmp11782 {
ifres11781 = True


} else {
ifres11781 = False


}

ifres11780 = ifres11781


} else {
ifres11780 = False


}

if True == ifres11780 {
tmp11742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1647, Nil)
}
__typedArg0 := V1647
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1646, tmp11742)
}
__typedArg0 := V1646
__typedArg1 := tmp11742
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4deref, tmp11743)
}
__typedArg0 := symshen_4deref
__typedArg1 := tmp11743
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11759 Obj

if True == tmp11778 {
tmp11776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp11776)
}
__typedArg0 := symlambda
__typedArg1 := tmp11776
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11761 Obj

if True == tmp11777 {
tmp11774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11774)
}
__typedArg0 := tmp11774
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11763 Obj

if True == tmp11775 {
tmp11771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11771)
}
__typedArg0 := tmp11771
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11772)
}
__typedArg0 := tmp11772
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11765 Obj

if True == tmp11773 {
tmp11767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11767)
}
__typedArg0 := tmp11767
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11768)
}
__typedArg0 := tmp11768
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11769)
}
__typedArg0 := Nil
__typedArg1 := tmp11769
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11766 Obj

if True == tmp11770 {
ifres11766 = True


} else {
ifres11766 = False


}

ifres11765 = ifres11766


} else {
ifres11765 = False


}

var ifres11764 Obj

if True == ifres11765 {
ifres11764 = True


} else {
ifres11764 = False


}

ifres11763 = ifres11764


} else {
ifres11763 = False


}

var ifres11762 Obj

if True == ifres11763 {
ifres11762 = True


} else {
ifres11762 = False


}

ifres11761 = ifres11762


} else {
ifres11761 = False


}

var ifres11760 Obj

if True == ifres11761 {
ifres11760 = True


} else {
ifres11760 = False


}

ifres11759 = ifres11760


} else {
ifres11759 = False


}

if True == ifres11759 {
tmp11744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11744)
}
__typedArg0 := tmp11744
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11746)
}
__typedArg0 := tmp11746
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11747)
}
__typedArg0 := tmp11747
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11749)
}
__typedArg0 := tmp11749
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11750, V1648)
}
__typedArg0 := tmp11750
__typedArg1 := V1648
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11752 := Call(__e, PrimFunc(symshen_4deref_1terms), tmp11748, V1647, tmp11751)


tmp11753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11752, Nil)
}
__typedArg0 := tmp11752
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11745, tmp11753)
}
__typedArg0 := tmp11745
__typedArg1 := tmp11753
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp11754)
}
__typedArg0 := symlambda
__typedArg1 := tmp11754
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp11757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1646)
}
__typedArg0 := V1646
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp11757 {
tmp11755 := MakeNative(func(__e *ControlFlow) {
Z1649 := __e.Get(1)
_ = Z1649
__e.TailApply(PrimFunc(symshen_4deref_1terms), Z1649, V1647, V1648)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp11755, V1646)
return


} else {
__e.Return(V1646)
return
}


}


}


}


}


}, 3)

tmp11815 := Call(__e, ns2_1set, symshen_4deref_1terms, tmp11720)


_ = tmp11815

tmp11816 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
__selftop:
V1667 := __self1
_ = V1667
V1668 := __self2
_ = V1668
V1669 := __self3
_ = V1669
V1670 := __self4
_ = V1670
V1671 := __self5
_ = V1671
tmp11992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1668)
}
__typedArg0 := Nil
__typedArg1 := V1668
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11989 Obj

if True == tmp11992 {
tmp11991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1669)
}
__typedArg0 := Nil
__typedArg1 := V1669
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11990 Obj

if True == tmp11991 {
ifres11990 = True


} else {
ifres11990 = False


}

ifres11989 = ifres11990


} else {
ifres11989 = False


}

if True == ifres11989 {
__e.Return(V1671)
return
} else {
tmp11987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11967 Obj

if True == tmp11987 {
tmp11985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11985)
}
__typedArg0 := tmp11985
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11969 Obj

if True == tmp11986 {
tmp11982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11982)
}
__typedArg0 := tmp11982
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, tmp11983)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp11983
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11971 Obj

if True == tmp11984 {
tmp11979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11979)
}
__typedArg0 := tmp11979
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11980)
}
__typedArg0 := tmp11980
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11973 Obj

if True == tmp11981 {
tmp11975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11975)
}
__typedArg0 := tmp11975
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11976)
}
__typedArg0 := tmp11976
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11977)
}
__typedArg0 := Nil
__typedArg1 := tmp11977
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11974 Obj

if True == tmp11978 {
ifres11974 = True


} else {
ifres11974 = False


}

ifres11973 = ifres11974


} else {
ifres11973 = False


}

var ifres11972 Obj

if True == ifres11973 {
ifres11972 = True


} else {
ifres11972 = False


}

ifres11971 = ifres11972


} else {
ifres11971 = False


}

var ifres11970 Obj

if True == ifres11971 {
ifres11970 = True


} else {
ifres11970 = False


}

ifres11969 = ifres11970


} else {
ifres11969 = False


}

var ifres11968 Obj

if True == ifres11969 {
ifres11968 = True


} else {
ifres11968 = False


}

ifres11967 = ifres11968


} else {
ifres11967 = False


}

if True == ifres11967 {
tmp11817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11817)
}
__typedArg0 := tmp11817
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11818)
}
__typedArg0 := tmp11818
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1667, tmp11820)
}
__typedArg0 := V1667
__typedArg1 := tmp11820
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11819, tmp11821)
}
__typedArg0 := tmp11819
__typedArg1 := tmp11821
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_7m, tmp11822)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp11822
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = V1667, tmp11823, V1669, V1670, V1671
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4compile_1head), V1667, tmp11823, V1669, V1670, V1671)
return


} else {
tmp11965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11945 Obj

if True == tmp11965 {
tmp11963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11963)
}
__typedArg0 := tmp11963
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11947 Obj

if True == tmp11964 {
tmp11960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11960)
}
__typedArg0 := tmp11960
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_1m, tmp11961)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp11961
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11949 Obj

if True == tmp11962 {
tmp11957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11957)
}
__typedArg0 := tmp11957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11958)
}
__typedArg0 := tmp11958
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11951 Obj

if True == tmp11959 {
tmp11953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11953)
}
__typedArg0 := tmp11953
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11954)
}
__typedArg0 := tmp11954
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11955)
}
__typedArg0 := Nil
__typedArg1 := tmp11955
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11952 Obj

if True == tmp11956 {
ifres11952 = True


} else {
ifres11952 = False


}

ifres11951 = ifres11952


} else {
ifres11951 = False


}

var ifres11950 Obj

if True == ifres11951 {
ifres11950 = True


} else {
ifres11950 = False


}

ifres11949 = ifres11950


} else {
ifres11949 = False


}

var ifres11948 Obj

if True == ifres11949 {
ifres11948 = True


} else {
ifres11948 = False


}

ifres11947 = ifres11948


} else {
ifres11947 = False


}

var ifres11946 Obj

if True == ifres11947 {
ifres11946 = True


} else {
ifres11946 = False


}

ifres11945 = ifres11946


} else {
ifres11945 = False


}

if True == ifres11945 {
tmp11824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11824)
}
__typedArg0 := tmp11824
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11825)
}
__typedArg0 := tmp11825
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1667, tmp11827)
}
__typedArg0 := V1667
__typedArg1 := tmp11827
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp11826, tmp11828)
}
__typedArg0 := tmp11826
__typedArg1 := tmp11828
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp11830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_1m, tmp11829)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp11829
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = V1667, tmp11830, V1669, V1670, V1671
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4compile_1head), V1667, tmp11830, V1669, V1670, V1671)
return


} else {
tmp11943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11939 Obj

if True == tmp11943 {
tmp11941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_1m, tmp11941)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp11941
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11940 Obj

if True == tmp11942 {
ifres11940 = True


} else {
ifres11940 = False


}

ifres11939 = ifres11940


} else {
ifres11939 = False


}

if True == ifres11939 {
tmp11831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = symshen_4_1m, tmp11831, V1669, V1670, V1671
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp11831, V1669, V1670, V1671)
return


} else {
tmp11937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11933 Obj

if True == tmp11937 {
tmp11935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, tmp11935)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp11935
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11934 Obj

if True == tmp11936 {
ifres11934 = True


} else {
ifres11934 = False


}

ifres11933 = ifres11934


} else {
ifres11933 = False


}

if True == ifres11933 {
tmp11832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = symshen_4_7m, tmp11832, V1669, V1670, V1671
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp11832, V1669, V1670, V1671)
return


} else {
tmp11931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11924 Obj

if True == tmp11931 {
tmp11930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1669)
}
__typedArg0 := V1669
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11926 Obj

if True == tmp11930 {
tmp11928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11929 := Call(__e, PrimFunc(symshen_4wildcard_2), tmp11928)


var ifres11927 Obj

if True == tmp11929 {
ifres11927 = True


} else {
ifres11927 = False


}

ifres11926 = ifres11927


} else {
ifres11926 = False


}

var ifres11925 Obj

if True == ifres11926 {
ifres11925 = True


} else {
ifres11925 = False


}

ifres11924 = ifres11925


} else {
ifres11924 = False


}

if True == ifres11924 {
tmp11833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1669)
}
__typedArg0 := V1669
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4compile_1head) == __self {
__self1, __self2, __self3, __self4, __self5 = V1667, tmp11833, tmp11834, V1670, V1671
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4compile_1head), V1667, tmp11833, tmp11834, V1670, V1671)
return


} else {
tmp11922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11918 Obj

if True == tmp11922 {
tmp11920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp11920)
}
__typedArg0 := tmp11920
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres11919 Obj

if True == tmp11921 {
ifres11919 = True


} else {
ifres11919 = False


}

ifres11918 = ifres11919


} else {
ifres11918 = False


}

if True == ifres11918 {
__e.TailApply(PrimFunc(symshen_4variable_1case), V1667, V1668, V1669, V1670, V1671)
return
} else {
tmp11916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_1m, V1667)
}
__typedArg0 := symshen_4_1m
__typedArg1 := V1667
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11909 Obj

if True == tmp11916 {
tmp11915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11911 Obj

if True == tmp11915 {
tmp11913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11914 := Call(__e, PrimFunc(symatom_2), tmp11913)


var ifres11912 Obj

if True == tmp11914 {
ifres11912 = True


} else {
ifres11912 = False


}

ifres11911 = ifres11912


} else {
ifres11911 = False


}

var ifres11910 Obj

if True == ifres11911 {
ifres11910 = True


} else {
ifres11910 = False


}

ifres11909 = ifres11910


} else {
ifres11909 = False


}

if True == ifres11909 {
__e.TailApply(PrimFunc(symshen_4atom_1case_1minus), V1668, V1669, V1670, V1671)
return
} else {
tmp11907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_1m, V1667)
}
__typedArg0 := symshen_4_1m
__typedArg1 := V1667
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11877 Obj

if True == tmp11907 {
tmp11906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11879 Obj

if True == tmp11906 {
tmp11904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11904)
}
__typedArg0 := tmp11904
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11881 Obj

if True == tmp11905 {
tmp11901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11901)
}
__typedArg0 := tmp11901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp11902)
}
__typedArg0 := symcons
__typedArg1 := tmp11902
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11883 Obj

if True == tmp11903 {
tmp11898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11898)
}
__typedArg0 := tmp11898
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11899)
}
__typedArg0 := tmp11899
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11885 Obj

if True == tmp11900 {
tmp11894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11894)
}
__typedArg0 := tmp11894
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11895)
}
__typedArg0 := tmp11895
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11896)
}
__typedArg0 := tmp11896
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11887 Obj

if True == tmp11897 {
tmp11889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11889)
}
__typedArg0 := tmp11889
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11890)
}
__typedArg0 := tmp11890
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11891)
}
__typedArg0 := tmp11891
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11892)
}
__typedArg0 := Nil
__typedArg1 := tmp11892
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11888 Obj

if True == tmp11893 {
ifres11888 = True


} else {
ifres11888 = False


}

ifres11887 = ifres11888


} else {
ifres11887 = False


}

var ifres11886 Obj

if True == ifres11887 {
ifres11886 = True


} else {
ifres11886 = False


}

ifres11885 = ifres11886


} else {
ifres11885 = False


}

var ifres11884 Obj

if True == ifres11885 {
ifres11884 = True


} else {
ifres11884 = False


}

ifres11883 = ifres11884


} else {
ifres11883 = False


}

var ifres11882 Obj

if True == ifres11883 {
ifres11882 = True


} else {
ifres11882 = False


}

ifres11881 = ifres11882


} else {
ifres11881 = False


}

var ifres11880 Obj

if True == ifres11881 {
ifres11880 = True


} else {
ifres11880 = False


}

ifres11879 = ifres11880


} else {
ifres11879 = False


}

var ifres11878 Obj

if True == ifres11879 {
ifres11878 = True


} else {
ifres11878 = False


}

ifres11877 = ifres11878


} else {
ifres11877 = False


}

if True == ifres11877 {
__e.TailApply(PrimFunc(symshen_4cons_1case_1minus), V1668, V1669, V1670, V1671)
return
} else {
tmp11875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, V1667)
}
__typedArg0 := symshen_4_7m
__typedArg1 := V1667
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11868 Obj

if True == tmp11875 {
tmp11874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11870 Obj

if True == tmp11874 {
tmp11872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11873 := Call(__e, PrimFunc(symatom_2), tmp11872)


var ifres11871 Obj

if True == tmp11873 {
ifres11871 = True


} else {
ifres11871 = False


}

ifres11870 = ifres11871


} else {
ifres11870 = False


}

var ifres11869 Obj

if True == ifres11870 {
ifres11869 = True


} else {
ifres11869 = False


}

ifres11868 = ifres11869


} else {
ifres11868 = False


}

if True == ifres11868 {
__e.TailApply(PrimFunc(symshen_4atom_1case_1plus), V1668, V1669, V1670, V1671)
return
} else {
tmp11866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, V1667)
}
__typedArg0 := symshen_4_7m
__typedArg1 := V1667
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11836 Obj

if True == tmp11866 {
tmp11865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11838 Obj

if True == tmp11865 {
tmp11863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11863)
}
__typedArg0 := tmp11863
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11840 Obj

if True == tmp11864 {
tmp11860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp11860)
}
__typedArg0 := tmp11860
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp11861)
}
__typedArg0 := symcons
__typedArg1 := tmp11861
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11842 Obj

if True == tmp11862 {
tmp11857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11857)
}
__typedArg0 := tmp11857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11858)
}
__typedArg0 := tmp11858
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11844 Obj

if True == tmp11859 {
tmp11853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11853)
}
__typedArg0 := tmp11853
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11854)
}
__typedArg0 := tmp11854
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp11855)
}
__typedArg0 := tmp11855
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres11846 Obj

if True == tmp11856 {
tmp11848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1668)
}
__typedArg0 := V1668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11848)
}
__typedArg0 := tmp11848
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11849)
}
__typedArg0 := tmp11849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp11850)
}
__typedArg0 := tmp11850
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp11851)
}
__typedArg0 := Nil
__typedArg1 := tmp11851
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres11847 Obj

if True == tmp11852 {
ifres11847 = True


} else {
ifres11847 = False


}

ifres11846 = ifres11847


} else {
ifres11846 = False


}

var ifres11845 Obj

if True == ifres11846 {
ifres11845 = True


} else {
ifres11845 = False


}

ifres11844 = ifres11845


} else {
ifres11844 = False


}

var ifres11843 Obj

if True == ifres11844 {
ifres11843 = True


} else {
ifres11843 = False


}

ifres11842 = ifres11843


} else {
ifres11842 = False


}

var ifres11841 Obj

if True == ifres11842 {
ifres11841 = True


} else {
ifres11841 = False


}

ifres11840 = ifres11841


} else {
ifres11840 = False


}

var ifres11839 Obj

if True == ifres11840 {
ifres11839 = True


} else {
ifres11839 = False


}

ifres11838 = ifres11839


} else {
ifres11838 = False


}

var ifres11837 Obj

if True == ifres11838 {
ifres11837 = True


} else {
ifres11837 = False


}

ifres11836 = ifres11837


} else {
ifres11836 = False


}

if True == ifres11836 {
__e.TailApply(PrimFunc(symshen_4cons_1case_1plus), V1668, V1669, V1670, V1671)
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.compile-head"))
}
__typedArg0 := MakeString("implementation error in shen.compile-head")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
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


}


}, 5)

tmp11993 := Call(__e, ns2_1set, symshen_4compile_1head, tmp11816)


_ = tmp11993

tmp11994 := MakeNative(func(__e *ControlFlow) {
V1682 := __e.Get(1)
_ = V1682
V1683 := __e.Get(2)
_ = V1683
V1684 := __e.Get(3)
_ = V1684
V1685 := __e.Get(4)
_ = V1685
V1686 := __e.Get(5)
_ = V1686
tmp12015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1683)
}
__typedArg0 := V1683
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12012 Obj

if True == tmp12015 {
tmp12014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1684)
}
__typedArg0 := V1684
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12013 Obj

if True == tmp12014 {
ifres12013 = True


} else {
ifres12013 = False


}

ifres12012 = ifres12013


} else {
ifres12012 = False


}

if True == ifres12012 {
tmp12009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1684)
}
__typedArg0 := V1684
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp12009)
}
__typedArg0 := tmp12009
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp12010 {
tmp11995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1683)
}
__typedArg0 := V1683
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1684)
}
__typedArg0 := V1684
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp11997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1684)
}
__typedArg0 := V1684
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1683)
}
__typedArg0 := V1683
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp11999 := Call(__e, PrimFunc(symsubst), tmp11997, tmp11998, V1686)


__e.TailApply(PrimFunc(symshen_4compile_1head), V1682, tmp11995, tmp11996, V1685, tmp11999)
return


} else {
tmp12000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1683)
}
__typedArg0 := V1683
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1684)
}
__typedArg0 := V1684
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1683)
}
__typedArg0 := V1683
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1684)
}
__typedArg0 := V1684
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12004 := Call(__e, PrimFunc(symshen_4compile_1head), V1682, tmp12002, tmp12003, V1685, V1686)


tmp12005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12004, Nil)
}
__typedArg0 := tmp12004
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12001, tmp12005)
}
__typedArg0 := tmp12001
__typedArg1 := tmp12005
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12000, tmp12006)
}
__typedArg0 := tmp12000
__typedArg1 := tmp12006
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp12007)
}
__typedArg0 := symlet
__typedArg1 := tmp12007
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.variable-case"))
}
__typedArg0 := MakeString("implementation error in shen.variable-case")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 5)

tmp12016 := Call(__e, ns2_1set, symshen_4variable_1case, tmp11994)


_ = tmp12016

tmp12017 := MakeNative(func(__e *ControlFlow) {
V1695 := __e.Get(1)
_ = V1695
V1696 := __e.Get(2)
_ = V1696
V1697 := __e.Get(3)
_ = V1697
V1698 := __e.Get(4)
_ = V1698
tmp12041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1695)
}
__typedArg0 := V1695
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12038 Obj

if True == tmp12041 {
tmp12040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1696)
}
__typedArg0 := V1696
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12039 Obj

if True == tmp12040 {
ifres12039 = True


} else {
ifres12039 = False


}

ifres12038 = ifres12039


} else {
ifres12038 = False


}

if True == ifres12038 {
tmp12018 := Call(__e, PrimFunc(symgensym), symTm)


let__10559 := tmp12018
_ = let__10559

tmp12019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1696)
}
__typedArg0 := V1696
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1697, Nil)
}
__typedArg0 := V1697
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12019, tmp12020)
}
__typedArg0 := tmp12019
__typedArg1 := tmp12020
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp12021)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp12021
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1695)
}
__typedArg0 := V1695
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12023, Nil)
}
__typedArg0 := tmp12023
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10559, tmp12024)
}
__typedArg0 := let__10559
__typedArg1 := tmp12024
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp12025)
}
__typedArg0 := sym_a
__typedArg1 := tmp12025
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1695)
}
__typedArg0 := V1695
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1696)
}
__typedArg0 := V1696
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12029 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp12027, tmp12028, V1697, V1698)


tmp12030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12029, tmp12030)
}
__typedArg0 := tmp12029
__typedArg1 := tmp12030
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12026, tmp12031)
}
__typedArg0 := tmp12026
__typedArg1 := tmp12031
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp12032)
}
__typedArg0 := symif
__typedArg1 := tmp12032
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12033, Nil)
}
__typedArg0 := tmp12033
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12022, tmp12034)
}
__typedArg0 := tmp12022
__typedArg1 := tmp12034
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10559, tmp12035)
}
__typedArg0 := let__10559
__typedArg1 := tmp12035
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp12036)
}
__typedArg0 := symlet
__typedArg1 := tmp12036
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.atom-case-minus"))
}
__typedArg0 := MakeString("implementation error in shen.atom-case-minus")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 4)

tmp12042 := Call(__e, ns2_1set, symshen_4atom_1case_1minus, tmp12017)


_ = tmp12042

tmp12043 := MakeNative(func(__e *ControlFlow) {
V1708 := __e.Get(1)
_ = V1708
V1709 := __e.Get(2)
_ = V1709
V1710 := __e.Get(3)
_ = V1710
V1711 := __e.Get(4)
_ = V1711
tmp12107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1708)
}
__typedArg0 := V1708
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12077 Obj

if True == tmp12107 {
tmp12105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1708)
}
__typedArg0 := V1708
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp12105)
}
__typedArg0 := tmp12105
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12079 Obj

if True == tmp12106 {
tmp12102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1708)
}
__typedArg0 := V1708
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp12102)
}
__typedArg0 := tmp12102
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp12103)
}
__typedArg0 := symcons
__typedArg1 := tmp12103
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12081 Obj

if True == tmp12104 {
tmp12099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1708)
}
__typedArg0 := V1708
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12099)
}
__typedArg0 := tmp12099
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp12100)
}
__typedArg0 := tmp12100
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12083 Obj

if True == tmp12101 {
tmp12095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1708)
}
__typedArg0 := V1708
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12095)
}
__typedArg0 := tmp12095
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12096)
}
__typedArg0 := tmp12096
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp12097)
}
__typedArg0 := tmp12097
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12085 Obj

if True == tmp12098 {
tmp12090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1708)
}
__typedArg0 := V1708
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12090)
}
__typedArg0 := tmp12090
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12091)
}
__typedArg0 := tmp12091
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12092)
}
__typedArg0 := tmp12092
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp12093)
}
__typedArg0 := Nil
__typedArg1 := tmp12093
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12087 Obj

if True == tmp12094 {
tmp12089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1709)
}
__typedArg0 := V1709
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12088 Obj

if True == tmp12089 {
ifres12088 = True


} else {
ifres12088 = False


}

ifres12087 = ifres12088


} else {
ifres12087 = False


}

var ifres12086 Obj

if True == ifres12087 {
ifres12086 = True


} else {
ifres12086 = False


}

ifres12085 = ifres12086


} else {
ifres12085 = False


}

var ifres12084 Obj

if True == ifres12085 {
ifres12084 = True


} else {
ifres12084 = False


}

ifres12083 = ifres12084


} else {
ifres12083 = False


}

var ifres12082 Obj

if True == ifres12083 {
ifres12082 = True


} else {
ifres12082 = False


}

ifres12081 = ifres12082


} else {
ifres12081 = False


}

var ifres12080 Obj

if True == ifres12081 {
ifres12080 = True


} else {
ifres12080 = False


}

ifres12079 = ifres12080


} else {
ifres12079 = False


}

var ifres12078 Obj

if True == ifres12079 {
ifres12078 = True


} else {
ifres12078 = False


}

ifres12077 = ifres12078


} else {
ifres12077 = False


}

if True == ifres12077 {
tmp12044 := Call(__e, PrimFunc(symgensym), symTm)


let__10560 := tmp12044
_ = let__10560

tmp12045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1709)
}
__typedArg0 := V1709
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1710, Nil)
}
__typedArg0 := V1710
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12045, tmp12046)
}
__typedArg0 := tmp12045
__typedArg1 := tmp12046
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp12047)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp12047
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10560, Nil)
}
__typedArg0 := let__10560
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp12049)
}
__typedArg0 := symcons_2
__typedArg1 := tmp12049
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1708)
}
__typedArg0 := V1708
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12051)
}
__typedArg0 := tmp12051
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp12052)
}
__typedArg0 := tmp12052
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1708)
}
__typedArg0 := V1708
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12054)
}
__typedArg0 := tmp12054
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12055)
}
__typedArg0 := tmp12055
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp12056)
}
__typedArg0 := tmp12056
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1708)
}
__typedArg0 := V1708
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12057, tmp12058)
}
__typedArg0 := tmp12057
__typedArg1 := tmp12058
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12053, tmp12059)
}
__typedArg0 := tmp12053
__typedArg1 := tmp12059
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10560, Nil)
}
__typedArg0 := let__10560
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhd, tmp12061)
}
__typedArg0 := symhd
__typedArg1 := tmp12061
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10560, Nil)
}
__typedArg0 := let__10560
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtl, tmp12063)
}
__typedArg0 := symtl
__typedArg1 := tmp12063
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1709)
}
__typedArg0 := V1709
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12064, tmp12065)
}
__typedArg0 := tmp12064
__typedArg1 := tmp12065
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12062, tmp12066)
}
__typedArg0 := tmp12062
__typedArg1 := tmp12066
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12068 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_1m, tmp12060, tmp12067, V1710, V1711)


tmp12069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12068, tmp12069)
}
__typedArg0 := tmp12068
__typedArg1 := tmp12069
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12050, tmp12070)
}
__typedArg0 := tmp12050
__typedArg1 := tmp12070
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp12071)
}
__typedArg0 := symif
__typedArg1 := tmp12071
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12072, Nil)
}
__typedArg0 := tmp12072
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12048, tmp12073)
}
__typedArg0 := tmp12048
__typedArg1 := tmp12073
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10560, tmp12074)
}
__typedArg0 := let__10560
__typedArg1 := tmp12074
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp12075)
}
__typedArg0 := symlet
__typedArg1 := tmp12075
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.cons-case-minus"))
}
__typedArg0 := MakeString("implementation error in shen.cons-case-minus")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 4)

tmp12108 := Call(__e, ns2_1set, symshen_4cons_1case_1minus, tmp12043)


_ = tmp12108

tmp12109 := MakeNative(func(__e *ControlFlow) {
V1721 := __e.Get(1)
_ = V1721
V1722 := __e.Get(2)
_ = V1722
V1723 := __e.Get(3)
_ = V1723
V1724 := __e.Get(4)
_ = V1724
tmp12153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1721)
}
__typedArg0 := V1721
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12150 Obj

if True == tmp12153 {
tmp12152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1722)
}
__typedArg0 := V1722
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12151 Obj

if True == tmp12152 {
ifres12151 = True


} else {
ifres12151 = False


}

ifres12150 = ifres12151


} else {
ifres12150 = False


}

if True == ifres12150 {
tmp12110 := Call(__e, PrimFunc(symgensym), symTm)


let__10561 := tmp12110
_ = let__10561

tmp12111 := Call(__e, PrimFunc(symgensym), symGoTo)


let__10562 := tmp12111
_ = let__10562

tmp12112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1722)
}
__typedArg0 := V1722
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1723, Nil)
}
__typedArg0 := V1723
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12112, tmp12113)
}
__typedArg0 := tmp12112
__typedArg1 := tmp12113
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp12114)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp12114
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1721)
}
__typedArg0 := V1721
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1722)
}
__typedArg0 := V1722
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12118 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp12116, tmp12117, V1723, V1724)


tmp12119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12118, Nil)
}
__typedArg0 := tmp12118
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp12119)
}
__typedArg0 := symfreeze
__typedArg1 := tmp12119
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1721)
}
__typedArg0 := V1721
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12121, Nil)
}
__typedArg0 := tmp12121
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10561, tmp12122)
}
__typedArg0 := let__10561
__typedArg1 := tmp12122
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp12123)
}
__typedArg0 := sym_a
__typedArg1 := tmp12123
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10562, Nil)
}
__typedArg0 := let__10562
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp12125)
}
__typedArg0 := symthaw
__typedArg1 := tmp12125
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10561, Nil)
}
__typedArg0 := let__10561
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4pvar_2, tmp12127)
}
__typedArg0 := symshen_4pvar_2
__typedArg1 := tmp12127
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1721)
}
__typedArg0 := V1721
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12130 := Call(__e, PrimFunc(symshen_4demode), tmp12129)


tmp12131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10562, Nil)
}
__typedArg0 := let__10562
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1723, tmp12131)
}
__typedArg0 := V1723
__typedArg1 := tmp12131
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12130, tmp12132)
}
__typedArg0 := tmp12130
__typedArg1 := tmp12132
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10561, tmp12133)
}
__typedArg0 := let__10561
__typedArg1 := tmp12133
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4bind_b, tmp12134)
}
__typedArg0 := symshen_4bind_b
__typedArg1 := tmp12134
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12135, tmp12136)
}
__typedArg0 := tmp12135
__typedArg1 := tmp12136
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12128, tmp12137)
}
__typedArg0 := tmp12128
__typedArg1 := tmp12137
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp12138)
}
__typedArg0 := symif
__typedArg1 := tmp12138
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12139, Nil)
}
__typedArg0 := tmp12139
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12126, tmp12140)
}
__typedArg0 := tmp12126
__typedArg1 := tmp12140
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12124, tmp12141)
}
__typedArg0 := tmp12124
__typedArg1 := tmp12141
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp12142)
}
__typedArg0 := symif
__typedArg1 := tmp12142
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12143, Nil)
}
__typedArg0 := tmp12143
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12120, tmp12144)
}
__typedArg0 := tmp12120
__typedArg1 := tmp12144
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10562, tmp12145)
}
__typedArg0 := let__10562
__typedArg1 := tmp12145
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12115, tmp12146)
}
__typedArg0 := tmp12115
__typedArg1 := tmp12146
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10561, tmp12147)
}
__typedArg0 := let__10561
__typedArg1 := tmp12147
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp12148)
}
__typedArg0 := symlet
__typedArg1 := tmp12148
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.atom-case-plus"))
}
__typedArg0 := MakeString("implementation error in shen.atom-case-plus")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 4)

tmp12154 := Call(__e, ns2_1set, symshen_4atom_1case_1plus, tmp12109)


_ = tmp12154

tmp12155 := MakeNative(func(__e *ControlFlow) {
V1735 := __e.Get(1)
_ = V1735
V1736 := __e.Get(2)
_ = V1736
V1737 := __e.Get(3)
_ = V1737
V1738 := __e.Get(4)
_ = V1738
tmp12246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1735)
}
__typedArg0 := V1735
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12216 Obj

if True == tmp12246 {
tmp12244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1735)
}
__typedArg0 := V1735
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp12244)
}
__typedArg0 := tmp12244
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12218 Obj

if True == tmp12245 {
tmp12241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1735)
}
__typedArg0 := V1735
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp12241)
}
__typedArg0 := tmp12241
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, tmp12242)
}
__typedArg0 := symcons
__typedArg1 := tmp12242
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12220 Obj

if True == tmp12243 {
tmp12238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1735)
}
__typedArg0 := V1735
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12238)
}
__typedArg0 := tmp12238
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp12239)
}
__typedArg0 := tmp12239
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12222 Obj

if True == tmp12240 {
tmp12234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1735)
}
__typedArg0 := V1735
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12234)
}
__typedArg0 := tmp12234
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12235)
}
__typedArg0 := tmp12235
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp12236)
}
__typedArg0 := tmp12236
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12224 Obj

if True == tmp12237 {
tmp12229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1735)
}
__typedArg0 := V1735
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12229)
}
__typedArg0 := tmp12229
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12230)
}
__typedArg0 := tmp12230
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12231)
}
__typedArg0 := tmp12231
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp12232)
}
__typedArg0 := Nil
__typedArg1 := tmp12232
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12226 Obj

if True == tmp12233 {
tmp12228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1736)
}
__typedArg0 := V1736
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12227 Obj

if True == tmp12228 {
ifres12227 = True


} else {
ifres12227 = False


}

ifres12226 = ifres12227


} else {
ifres12226 = False


}

var ifres12225 Obj

if True == ifres12226 {
ifres12225 = True


} else {
ifres12225 = False


}

ifres12224 = ifres12225


} else {
ifres12224 = False


}

var ifres12223 Obj

if True == ifres12224 {
ifres12223 = True


} else {
ifres12223 = False


}

ifres12222 = ifres12223


} else {
ifres12222 = False


}

var ifres12221 Obj

if True == ifres12222 {
ifres12221 = True


} else {
ifres12221 = False


}

ifres12220 = ifres12221


} else {
ifres12220 = False


}

var ifres12219 Obj

if True == ifres12220 {
ifres12219 = True


} else {
ifres12219 = False


}

ifres12218 = ifres12219


} else {
ifres12218 = False


}

var ifres12217 Obj

if True == ifres12218 {
ifres12217 = True


} else {
ifres12217 = False


}

ifres12216 = ifres12217


} else {
ifres12216 = False


}

if True == ifres12216 {
tmp12156 := Call(__e, PrimFunc(symgensym), symTm)


let__10563 := tmp12156
_ = let__10563

tmp12157 := Call(__e, PrimFunc(symgensym), symGoTo)


let__10564 := tmp12157
_ = let__10564

tmp12158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1735)
}
__typedArg0 := V1735
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12158)
}
__typedArg0 := tmp12158
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp12159)
}
__typedArg0 := tmp12159
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1735)
}
__typedArg0 := V1735
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12161)
}
__typedArg0 := tmp12161
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12162)
}
__typedArg0 := tmp12162
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp12163)
}
__typedArg0 := tmp12163
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12160, tmp12164)
}
__typedArg0 := tmp12160
__typedArg1 := tmp12164
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12166 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp12165)


let__10565 := tmp12166
_ = let__10565

tmp12167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1735)
}
__typedArg0 := V1735
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12168 := Call(__e, PrimFunc(symshen_4tame), tmp12167)


let__10566 := tmp12168
_ = let__10566

tmp12169 := Call(__e, PrimFunc(symshen_4extract_1vars), let__10566)


let__10567 := tmp12169
_ = let__10567

tmp12170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1736)
}
__typedArg0 := V1736
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1737, Nil)
}
__typedArg0 := V1737
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12170, tmp12171)
}
__typedArg0 := tmp12170
__typedArg1 := tmp12171
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4lazyderef, tmp12172)
}
__typedArg0 := symshen_4lazyderef
__typedArg1 := tmp12172
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1735)
}
__typedArg0 := V1735
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1736)
}
__typedArg0 := V1736
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12176 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp12174, tmp12175, V1737, V1738)


tmp12177 := Call(__e, PrimFunc(symshen_4goto), let__10565, tmp12176)


tmp12178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10563, Nil)
}
__typedArg0 := let__10563
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons_2, tmp12178)
}
__typedArg0 := symcons_2
__typedArg1 := tmp12178
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1735)
}
__typedArg0 := V1735
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12180)
}
__typedArg0 := tmp12180
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10563, Nil)
}
__typedArg0 := let__10563
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symhd, tmp12182)
}
__typedArg0 := symhd
__typedArg1 := tmp12182
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10563, Nil)
}
__typedArg0 := let__10563
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtl, tmp12184)
}
__typedArg0 := symtl
__typedArg1 := tmp12184
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12185, Nil)
}
__typedArg0 := tmp12185
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12183, tmp12186)
}
__typedArg0 := tmp12183
__typedArg1 := tmp12186
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12188 := Call(__e, PrimFunc(symshen_4invoke), let__10564, let__10565)


tmp12189 := Call(__e, PrimFunc(symshen_4compile_1head), symshen_4_7m, tmp12181, tmp12187, V1737, tmp12188)


tmp12190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10563, Nil)
}
__typedArg0 := let__10563
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4pvar_2, tmp12190)
}
__typedArg0 := symshen_4pvar_2
__typedArg1 := tmp12190
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12192 := Call(__e, PrimFunc(symshen_4demode), let__10566)


tmp12193 := Call(__e, PrimFunc(symshen_4invoke), let__10564, let__10565)


tmp12194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12193, Nil)
}
__typedArg0 := tmp12193
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp12194)
}
__typedArg0 := symfreeze
__typedArg1 := tmp12194
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12195, Nil)
}
__typedArg0 := tmp12195
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1737, tmp12196)
}
__typedArg0 := V1737
__typedArg1 := tmp12196
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12192, tmp12197)
}
__typedArg0 := tmp12192
__typedArg1 := tmp12197
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10563, tmp12198)
}
__typedArg0 := let__10563
__typedArg1 := tmp12198
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4bind_b, tmp12199)
}
__typedArg0 := symshen_4bind_b
__typedArg1 := tmp12199
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12201 := Call(__e, PrimFunc(symshen_4stpart), let__10567, tmp12200, V1737)


tmp12202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12201, tmp12202)
}
__typedArg0 := tmp12201
__typedArg1 := tmp12202
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12191, tmp12203)
}
__typedArg0 := tmp12191
__typedArg1 := tmp12203
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp12204)
}
__typedArg0 := symif
__typedArg1 := tmp12204
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12205, Nil)
}
__typedArg0 := tmp12205
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12189, tmp12206)
}
__typedArg0 := tmp12189
__typedArg1 := tmp12206
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12179, tmp12207)
}
__typedArg0 := tmp12179
__typedArg1 := tmp12207
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp12208)
}
__typedArg0 := symif
__typedArg1 := tmp12208
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12209, Nil)
}
__typedArg0 := tmp12209
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12177, tmp12210)
}
__typedArg0 := tmp12177
__typedArg1 := tmp12210
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10564, tmp12211)
}
__typedArg0 := let__10564
__typedArg1 := tmp12211
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12173, tmp12212)
}
__typedArg0 := tmp12173
__typedArg1 := tmp12212
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__10563, tmp12213)
}
__typedArg0 := let__10563
__typedArg1 := tmp12213
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp12214)
}
__typedArg0 := symlet
__typedArg1 := tmp12214
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.cons-case-plus"))
}
__typedArg0 := MakeString("implementation error in shen.cons-case-plus")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 4)

tmp12247 := Call(__e, ns2_1set, symshen_4cons_1case_1plus, tmp12155)


_ = tmp12247

tmp12248 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V1744 := __self1
_ = V1744
tmp12285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1744)
}
__typedArg0 := V1744
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12272 Obj

if True == tmp12285 {
tmp12283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1744)
}
__typedArg0 := V1744
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7m, tmp12283)
}
__typedArg0 := symshen_4_7m
__typedArg1 := tmp12283
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12274 Obj

if True == tmp12284 {
tmp12281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1744)
}
__typedArg0 := V1744
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp12281)
}
__typedArg0 := tmp12281
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12276 Obj

if True == tmp12282 {
tmp12278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1744)
}
__typedArg0 := V1744
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12279 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12278)
}
__typedArg0 := tmp12278
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp12279)
}
__typedArg0 := Nil
__typedArg1 := tmp12279
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12277 Obj

if True == tmp12280 {
ifres12277 = True


} else {
ifres12277 = False


}

ifres12276 = ifres12277


} else {
ifres12276 = False


}

var ifres12275 Obj

if True == ifres12276 {
ifres12275 = True


} else {
ifres12275 = False


}

ifres12274 = ifres12275


} else {
ifres12274 = False


}

var ifres12273 Obj

if True == ifres12274 {
ifres12273 = True


} else {
ifres12273 = False


}

ifres12272 = ifres12273


} else {
ifres12272 = False


}

if True == ifres12272 {
tmp12249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1744)
}
__typedArg0 := V1744
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp12249)
}
__typedArg0 := tmp12249
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

if PrimFunc(symshen_4demode) == __self {
__self1 = tmp12250
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4demode), tmp12250)
return


} else {
tmp12270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1744)
}
__typedArg0 := V1744
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12257 Obj

if True == tmp12270 {
tmp12268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1744)
}
__typedArg0 := V1744
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_1m, tmp12268)
}
__typedArg0 := symshen_4_1m
__typedArg1 := tmp12268
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12259 Obj

if True == tmp12269 {
tmp12266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1744)
}
__typedArg0 := V1744
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp12266)
}
__typedArg0 := tmp12266
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12261 Obj

if True == tmp12267 {
tmp12263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1744)
}
__typedArg0 := V1744
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp12263)
}
__typedArg0 := tmp12263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp12264)
}
__typedArg0 := Nil
__typedArg1 := tmp12264
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres12262 Obj

if True == tmp12265 {
ifres12262 = True


} else {
ifres12262 = False


}

ifres12261 = ifres12262


} else {
ifres12261 = False


}

var ifres12260 Obj

if True == ifres12261 {
ifres12260 = True


} else {
ifres12260 = False


}

ifres12259 = ifres12260


} else {
ifres12259 = False


}

var ifres12258 Obj

if True == ifres12259 {
ifres12258 = True


} else {
ifres12258 = False


}

ifres12257 = ifres12258


} else {
ifres12257 = False


}

if True == ifres12257 {
tmp12251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1744)
}
__typedArg0 := V1744
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp12251)
}
__typedArg0 := tmp12251
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

if PrimFunc(symshen_4demode) == __self {
__self1 = tmp12252
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4demode), tmp12252)
return


} else {
tmp12255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1744)
}
__typedArg0 := V1744
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12255 {
tmp12253 := MakeNative(func(__e *ControlFlow) {
Z1745 := __e.Get(1)
_ = Z1745
__e.TailApply(PrimFunc(symshen_4demode), Z1745)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp12253, V1744)
return


} else {
__e.Return(V1744)
return
}


}


}


}, 1)

tmp12286 := Call(__e, ns2_1set, symshen_4demode, tmp12248)


_ = tmp12286

tmp12287 := MakeNative(func(__e *ControlFlow) {
V1746 := __e.Get(1)
_ = V1746
tmp12292 := Call(__e, PrimFunc(symshen_4wildcard_2), V1746)


if True == tmp12292 {
__e.TailApply(PrimFunc(symgensym), symY)
return
} else {
tmp12290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1746)
}
__typedArg0 := V1746
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12290 {
tmp12288 := MakeNative(func(__e *ControlFlow) {
Z1747 := __e.Get(1)
_ = Z1747
__e.TailApply(PrimFunc(symshen_4tame), Z1747)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp12288, V1746)
return


} else {
__e.Return(V1746)
return
}


}


}, 1)

tmp12293 := Call(__e, ns2_1set, symshen_4tame, tmp12287)


_ = tmp12293

tmp12294 := MakeNative(func(__e *ControlFlow) {
V1748 := __e.Get(1)
_ = V1748
V1749 := __e.Get(2)
_ = V1749
tmp12297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1748)
}
__typedArg0 := Nil
__typedArg1 := V1748
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12297 {
tmp12295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1749, Nil)
}
__typedArg0 := V1749
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp12295)
}
__typedArg0 := symfreeze
__typedArg1 := tmp12295
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4goto_1h), V1748, V1749)
return
}


}, 2)

tmp12298 := Call(__e, ns2_1set, symshen_4goto, tmp12294)


_ = tmp12298

tmp12299 := MakeNative(func(__e *ControlFlow) {
V1750 := __e.Get(1)
_ = V1750
V1751 := __e.Get(2)
_ = V1751
tmp12308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1750)
}
__typedArg0 := Nil
__typedArg1 := V1750
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12308 {
__e.Return(V1751)
return
} else {
tmp12306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1750)
}
__typedArg0 := V1750
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12306 {
tmp12300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1750)
}
__typedArg0 := V1750
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1750)
}
__typedArg0 := V1750
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12302 := Call(__e, PrimFunc(symshen_4goto_1h), tmp12301, V1751)


tmp12303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12302, Nil)
}
__typedArg0 := tmp12302
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12300, tmp12303)
}
__typedArg0 := tmp12300
__typedArg1 := tmp12303
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp12304)
}
__typedArg0 := symlambda
__typedArg1 := tmp12304
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.goto-h"))
}
__typedArg0 := MakeString("partial function shen.goto-h")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp12309 := Call(__e, ns2_1set, symshen_4goto_1h, tmp12299)


_ = tmp12309

tmp12310 := MakeNative(func(__e *ControlFlow) {
V1752 := __e.Get(1)
_ = V1752
V1753 := __e.Get(2)
_ = V1753
tmp12313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1753)
}
__typedArg0 := Nil
__typedArg1 := V1753
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12313 {
tmp12311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1752, Nil)
}
__typedArg0 := V1752
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp12311)
}
__typedArg0 := symthaw
__typedArg1 := tmp12311
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1752, V1753)
}
__typedArg0 := V1752
__typedArg1 := V1753
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
}


}, 2)

tmp12314 := Call(__e, ns2_1set, symshen_4invoke, tmp12310)


_ = tmp12314

tmp12315 := MakeNative(func(__e *ControlFlow) {
V1754 := __e.Get(1)
_ = V1754
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1754, sym__)
}
__typedArg0 := V1754
__typedArg1 := sym__
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp12316 := Call(__e, ns2_1set, symshen_4wildcard_2, tmp12315)


_ = tmp12316

tmp12317 := MakeNative(func(__e *ControlFlow) {
V1755 := __e.Get(1)
_ = V1755
tmp12318 := MakeNative(func(__e *ControlFlow) {
tmp12323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector_2) {
return PrimIsVector(V1755)
}
__typedArg0 := V1755
return Call(__e, PrimFunc(symabsvector_2), __typedArg0)
})()

if True == tmp12323 {
tmp12320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1755, MakeInteger(0))
}
__typedArg0 := V1755
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp12321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp12320, symshen_4pvar)
}
__typedArg0 := tmp12320
__typedArg1 := symshen_4pvar
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12321 {
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


}, 0)

tmp12324 := MakeNative(func(__e *ControlFlow) {
Z1756 := __e.Get(1)
_ = Z1756
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp12318, tmp12324)
return


}, 1)

tmp12325 := Call(__e, ns2_1set, symshen_4pvar_2, tmp12317)


_ = tmp12325

tmp12326 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V1757 := __self1
_ = V1757
V1758 := __self2
_ = V1758
tmp12332 := Call(__e, PrimFunc(symshen_4pvar_2), V1757)


if True == tmp12332 {
tmp12327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1757, MakeInteger(1))
}
__typedArg0 := V1757
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp12328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1758, tmp12327)
}
__typedArg0 := V1758
__typedArg1 := tmp12327
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

let__10568 := tmp12328
_ = let__10568

tmp12330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10568, symshen_4_1null_1)
}
__typedArg0 := let__10568
__typedArg1 := symshen_4_1null_1
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12330 {
__e.Return(V1757)
return
} else {
if PrimFunc(symshen_4lazyderef) == __self {
__self1, __self2 = let__10568, V1758
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4lazyderef), let__10568, V1758)
return
}


} else {
__e.Return(V1757)
return
}


}, 2)

tmp12333 := Call(__e, ns2_1set, symshen_4lazyderef, tmp12326)


_ = tmp12333

tmp12334 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V1760 := __self1
_ = V1760
V1761 := __self2
_ = V1761
tmp12346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1760)
}
__typedArg0 := V1760
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12346 {
tmp12335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1760)
}
__typedArg0 := V1760
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12336 := Call(__e, PrimFunc(symshen_4deref), tmp12335, V1761)


tmp12337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1760)
}
__typedArg0 := V1760
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12338 := Call(__e, PrimFunc(symshen_4deref), tmp12337, V1761)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12336, tmp12338)
}
__typedArg0 := tmp12336
__typedArg1 := tmp12338
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp12344 := Call(__e, PrimFunc(symshen_4pvar_2), V1760)


if True == tmp12344 {
tmp12339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1760, MakeInteger(1))
}
__typedArg0 := V1760
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp12340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1761, tmp12339)
}
__typedArg0 := V1761
__typedArg1 := tmp12339
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

let__10569 := tmp12340
_ = let__10569

tmp12342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10569, symshen_4_1null_1)
}
__typedArg0 := let__10569
__typedArg1 := symshen_4_1null_1
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12342 {
__e.Return(V1760)
return
} else {
if PrimFunc(symshen_4deref) == __self {
__self1, __self2 = let__10569, V1761
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4deref), let__10569, V1761)
return
}


} else {
__e.Return(V1760)
return
}


}


}, 2)

tmp12347 := Call(__e, ns2_1set, symshen_4deref, tmp12334)


_ = tmp12347

tmp12348 := MakeNative(func(__e *ControlFlow) {
V1763 := __e.Get(1)
_ = V1763
V1764 := __e.Get(2)
_ = V1764
V1765 := __e.Get(3)
_ = V1765
V1766 := __e.Get(4)
_ = V1766
tmp12349 := Call(__e, PrimFunc(symshen_4bindv), V1763, V1764, V1765)


let__10570 := tmp12349
_ = let__10570

tmp12350 := Call(__e, PrimFunc(symthaw), V1766)


let__10571 := tmp12350
_ = let__10571

tmp12352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10571, False)
}
__typedArg0 := let__10571
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12352 {
__e.TailApply(PrimFunc(symshen_4unwind), V1763, V1765, let__10571)
return
} else {
__e.Return(let__10571)
return
}


}, 4)

tmp12353 := Call(__e, ns2_1set, symshen_4bind_b, tmp12348)


_ = tmp12353

tmp12354 := MakeNative(func(__e *ControlFlow) {
V1769 := __e.Get(1)
_ = V1769
V1770 := __e.Get(2)
_ = V1770
V1771 := __e.Get(3)
_ = V1771
tmp12355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1769, MakeInteger(1))
}
__typedArg0 := V1769
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V1771, tmp12355, V1770)
}
__typedArg0 := V1771
__typedArg1 := tmp12355
__typedArg2 := V1770
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return


}, 3)

tmp12356 := Call(__e, ns2_1set, symshen_4bindv, tmp12354)


_ = tmp12356

tmp12357 := MakeNative(func(__e *ControlFlow) {
V1772 := __e.Get(1)
_ = V1772
V1773 := __e.Get(2)
_ = V1773
V1774 := __e.Get(3)
_ = V1774
tmp12358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1772, MakeInteger(1))
}
__typedArg0 := V1772
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp12359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V1773, tmp12358, symshen_4_1null_1)
}
__typedArg0 := V1773
__typedArg1 := tmp12358
__typedArg2 := symshen_4_1null_1
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

_ = tmp12359

__e.Return(V1774)
return


}, 3)

tmp12360 := Call(__e, ns2_1set, symshen_4unwind, tmp12357)


_ = tmp12360

tmp12361 := MakeNative(func(__e *ControlFlow) {
V1783 := __e.Get(1)
_ = V1783
V1784 := __e.Get(2)
_ = V1784
V1785 := __e.Get(3)
_ = V1785
tmp12376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1783)
}
__typedArg0 := Nil
__typedArg1 := V1783
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12376 {
__e.Return(V1784)
return
} else {
tmp12374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1783)
}
__typedArg0 := V1783
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12374 {
tmp12362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1783)
}
__typedArg0 := V1783
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1785, Nil)
}
__typedArg0 := V1785
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4newpv, tmp12363)
}
__typedArg0 := symshen_4newpv
__typedArg1 := tmp12363
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1783)
}
__typedArg0 := V1783
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12366 := Call(__e, PrimFunc(symshen_4stpart), tmp12365, V1784, V1785)


tmp12367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12366, Nil)
}
__typedArg0 := tmp12366
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V1785, tmp12367)
}
__typedArg0 := V1785
__typedArg1 := tmp12367
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4gc, tmp12368)
}
__typedArg0 := symshen_4gc
__typedArg1 := tmp12368
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12369, Nil)
}
__typedArg0 := tmp12369
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12364, tmp12370)
}
__typedArg0 := tmp12364
__typedArg1 := tmp12370
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12362, tmp12371)
}
__typedArg0 := tmp12362
__typedArg1 := tmp12371
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp12372)
}
__typedArg0 := symlet
__typedArg1 := tmp12372
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.stpart"))
}
__typedArg0 := MakeString("implementation error in shen.stpart")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 3)

tmp12377 := Call(__e, ns2_1set, symshen_4stpart, tmp12361)


_ = tmp12377

tmp12378 := MakeNative(func(__e *ControlFlow) {
V1786 := __e.Get(1)
_ = V1786
V1787 := __e.Get(2)
_ = V1787
tmp12382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1787, False)
}
__typedArg0 := V1787
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12382 {
tmp12379 := Call(__e, PrimFunc(symshen_4ticket_1number), V1786)


let__10572 := tmp12379
_ = let__10572

tmp12380 := Call(__e, PrimFunc(symshen_4decrement_1ticket), let__10572, V1786)


_ = tmp12380

__e.Return(V1787)
return


} else {
__e.Return(V1787)
return
}


}, 2)

tmp12383 := Call(__e, ns2_1set, symshen_4gc, tmp12378)


_ = tmp12383

tmp12384 := MakeNative(func(__e *ControlFlow) {
V1789 := __e.Get(1)
_ = V1789
V1790 := __e.Get(2)
_ = V1790
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V1790, MakeInteger(1), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V1789)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V1789
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())
}
__typedArg0 := V1790
__typedArg1 := MakeInteger(1)
__typedArg2 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V1789)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V1789
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return


}, 2)

tmp12386 := Call(__e, ns2_1set, symshen_4decrement_1ticket, tmp12384)


_ = tmp12386

tmp12387 := MakeNative(func(__e *ControlFlow) {
V1791 := __e.Get(1)
_ = V1791
tmp12388 := Call(__e, PrimFunc(symshen_4ticket_1number), V1791)


let__10573 := tmp12388
_ = let__10573

tmp12389 := Call(__e, PrimFunc(symshen_4make_1prolog_1variable), let__10573)


let__10574 := tmp12389
_ = let__10574

tmp12390 := Call(__e, PrimFunc(symshen_4nextticket), V1791, let__10573)


let__10575 := tmp12390
_ = let__10575

__e.Return(let__10574)
return


}, 1)

tmp12391 := Call(__e, ns2_1set, symshen_4newpv, tmp12387)


_ = tmp12391

tmp12392 := MakeNative(func(__e *ControlFlow) {
V1795 := __e.Get(1)
_ = V1795
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1795, MakeInteger(1))
}
__typedArg0 := V1795
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp12393 := Call(__e, ns2_1set, symshen_4ticket_1number, tmp12392)


_ = tmp12393

tmp12394 := MakeNative(func(__e *ControlFlow) {
V1796 := __e.Get(1)
_ = V1796
V1797 := __e.Get(2)
_ = V1797
tmp12395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(V1796, V1797, symshen_4_1null_1)
}
__typedArg0 := V1796
__typedArg1 := V1797
__typedArg2 := symshen_4_1null_1
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__10576 := tmp12395
_ = let__10576

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__10576, MakeInteger(1), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V1797)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V1797
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := let__10576
__typedArg1 := MakeInteger(1)
__typedArg2 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(V1797)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := V1797
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return


}, 2)

tmp12397 := Call(__e, ns2_1set, symshen_4nextticket, tmp12394)


_ = tmp12397

tmp12398 := MakeNative(func(__e *ControlFlow) {
V1799 := __e.Get(1)
_ = V1799
tmp12399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(MakeInteger(2))
}
__typedArg0 := MakeInteger(2)
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

tmp12400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(tmp12399, MakeInteger(0), symshen_4pvar)
}
__typedArg0 := tmp12399
__typedArg1 := MakeInteger(0)
__typedArg2 := symshen_4pvar
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(tmp12400, MakeInteger(1), V1799)
}
__typedArg0 := tmp12400
__typedArg1 := MakeInteger(1)
__typedArg2 := V1799
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})())
return


}, 1)

tmp12401 := Call(__e, ns2_1set, symshen_4make_1prolog_1variable, tmp12398)


_ = tmp12401

tmp12402 := MakeNative(func(__e *ControlFlow) {
V1800 := __e.Get(1)
_ = V1800
tmp12403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V1800, MakeInteger(1))
}
__typedArg0 := V1800
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})()

tmp12404 := Call(__e, PrimFunc(symshen_4app), tmp12403, MakeString(""), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Var"))
__typedS1, __typedOK1 := TypedString(tmp12404)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Var")
__typedArg1 := tmp12404
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp12405 := Call(__e, ns2_1set, symshen_4pvar, tmp12402)


_ = tmp12405

tmp12406 := MakeNative(func(__e *ControlFlow) {
tmp12407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dinfs_d)
}
__typedArg0 := symshen_4_dinfs_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dinfs_d, (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp12407)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp12407
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})())
}
__typedArg0 := symshen_4_dinfs_d
__typedArg1 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_7) {
__typedN0, __typedOK0 := TypedFloat64(MakeNumber(1))
__typedN1, __typedOK1 := TypedFloat64(tmp12407)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_7) {
return TypedMaterializeNumber((__typedN0 + __typedN1))
}}
__typedArg0 := MakeInteger(1)
__typedArg1 := tmp12407
return Call(__e, PrimFunc(sym_7), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return


}, 0)

tmp12409 := Call(__e, ns2_1set, symshen_4incinfs, tmp12406)


_ = tmp12409

tmp12410 := MakeNative(func(__e *ControlFlow) {
V1801 := __e.Get(1)
_ = V1801
tmp12417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(syminteger_2) {
return PrimIsInteger(V1801)
}
__typedArg0 := V1801
return Call(__e, PrimFunc(syminteger_2), __typedArg0)
})()

var ifres12414 Obj

if True == tmp12417 {
tmp12416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(V1801)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(0))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := V1801
__typedArg1 := MakeInteger(0)
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})()

var ifres12415 Obj

if True == tmp12416 {
ifres12415 = True


} else {
ifres12415 = False


}

ifres12414 = ifres12415


} else {
ifres12414 = False


}

if True == ifres12414 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dsize_1prolog_1vector_d, V1801)
}
__typedArg0 := symshen_4_dsize_1prolog_1vector_d
__typedArg1 := V1801
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
tmp12411 := Call(__e, PrimFunc(symshen_4app), V1801, MakeString(""), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("prolog vector size: size should be a positive integer; not "))
__typedS1, __typedOK1 := TypedString(tmp12411)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("prolog vector size: size should be a positive integer; not ")
__typedArg1 := tmp12411
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("prolog vector size: size should be a positive integer; not "))
__typedS1, __typedOK1 := TypedString(tmp12411)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("prolog vector size: size should be a positive integer; not ")
__typedArg1 := tmp12411
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}, 1)

tmp12418 := Call(__e, ns2_1set, symshen_4prolog_1vector_1size, tmp12410)


_ = tmp12418

tmp12419 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__selftop:
V1813 := __self1
_ = V1813
V1814 := __self2
_ = V1814
V1815 := __self3
_ = V1815
V1816 := __self4
_ = V1816
tmp12449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1813, V1814)
}
__typedArg0 := V1813
__typedArg1 := V1814
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12449 {
__e.TailApply(PrimFunc(symthaw), V1816)
return
} else {
tmp12447 := Call(__e, PrimFunc(symshen_4pvar_2), V1813)


var ifres12442 Obj

if True == tmp12447 {
tmp12444 := Call(__e, PrimFunc(symshen_4deref), V1814, V1815)


tmp12445 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V1813, tmp12444)


tmp12446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp12445)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp12445
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres12443 Obj

if True == tmp12446 {
ifres12443 = True


} else {
ifres12443 = False


}

ifres12442 = ifres12443


} else {
ifres12442 = False


}

if True == ifres12442 {
__e.TailApply(PrimFunc(symshen_4bind_b), V1813, V1814, V1815, V1816)
return
} else {
tmp12440 := Call(__e, PrimFunc(symshen_4pvar_2), V1814)


var ifres12435 Obj

if True == tmp12440 {
tmp12437 := Call(__e, PrimFunc(symshen_4deref), V1813, V1815)


tmp12438 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V1814, tmp12437)


tmp12439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp12438)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp12438
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres12436 Obj

if True == tmp12439 {
ifres12436 = True


} else {
ifres12436 = False


}

ifres12435 = ifres12436


} else {
ifres12435 = False


}

if True == ifres12435 {
__e.TailApply(PrimFunc(symshen_4bind_b), V1814, V1813, V1815, V1816)
return
} else {
tmp12433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1813)
}
__typedArg0 := V1813
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12430 Obj

if True == tmp12433 {
tmp12432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1814)
}
__typedArg0 := V1814
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12431 Obj

if True == tmp12432 {
ifres12431 = True


} else {
ifres12431 = False


}

ifres12430 = ifres12431


} else {
ifres12430 = False


}

if True == ifres12430 {
tmp12420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1813)
}
__typedArg0 := V1813
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12421 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12420, V1815)


tmp12422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1814)
}
__typedArg0 := V1814
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12423 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12422, V1815)


tmp12424 := MakeNative(func(__e *ControlFlow) {
tmp12425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1813)
}
__typedArg0 := V1813
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12426 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12425, V1815)


tmp12427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1814)
}
__typedArg0 := V1814
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12428 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12427, V1815)


__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp12426, tmp12428, V1815, V1816)
return


}, 0)

if PrimFunc(symshen_4lzy_a_b) == __self {
__self1, __self2, __self3, __self4 = tmp12421, tmp12423, V1815, tmp12424
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp12421, tmp12423, V1815, tmp12424)
return


} else {
__e.Return(False)
return
}


}


}


}


}, 4)

tmp12450 := Call(__e, ns2_1set, symshen_4lzy_a_b, tmp12419)


_ = tmp12450

tmp12451 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__selftop:
V1828 := __self1
_ = V1828
V1829 := __self2
_ = V1829
V1830 := __self3
_ = V1830
V1831 := __self4
_ = V1831
tmp12471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1828, V1829)
}
__typedArg0 := V1828
__typedArg1 := V1829
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12471 {
__e.TailApply(PrimFunc(symthaw), V1831)
return
} else {
tmp12469 := Call(__e, PrimFunc(symshen_4pvar_2), V1828)


if True == tmp12469 {
__e.TailApply(PrimFunc(symshen_4bind_b), V1828, V1829, V1830, V1831)
return
} else {
tmp12467 := Call(__e, PrimFunc(symshen_4pvar_2), V1829)


if True == tmp12467 {
__e.TailApply(PrimFunc(symshen_4bind_b), V1829, V1828, V1830, V1831)
return
} else {
tmp12465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1828)
}
__typedArg0 := V1828
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12462 Obj

if True == tmp12465 {
tmp12464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1829)
}
__typedArg0 := V1829
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12463 Obj

if True == tmp12464 {
ifres12463 = True


} else {
ifres12463 = False


}

ifres12462 = ifres12463


} else {
ifres12462 = False


}

if True == ifres12462 {
tmp12452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1828)
}
__typedArg0 := V1828
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12453 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12452, V1830)


tmp12454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1829)
}
__typedArg0 := V1829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12455 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12454, V1830)


tmp12456 := MakeNative(func(__e *ControlFlow) {
tmp12457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1828)
}
__typedArg0 := V1828
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12458 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12457, V1830)


tmp12459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1829)
}
__typedArg0 := V1829
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12460 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12459, V1830)


__e.TailApply(PrimFunc(symshen_4lzy_a), tmp12458, tmp12460, V1830, V1831)
return


}, 0)

if PrimFunc(symshen_4lzy_a) == __self {
__self1, __self2, __self3, __self4 = tmp12453, tmp12455, V1830, tmp12456
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4lzy_a), tmp12453, tmp12455, V1830, tmp12456)
return


} else {
__e.Return(False)
return
}


}


}


}


}, 4)

tmp12472 := Call(__e, ns2_1set, symshen_4lzy_a, tmp12451)


_ = tmp12472

tmp12473 := MakeNative(func(__e *ControlFlow) {
V1837 := __e.Get(1)
_ = V1837
V1838 := __e.Get(2)
_ = V1838
tmp12483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V1837, V1838)
}
__typedArg0 := V1837
__typedArg1 := V1838
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12483 {
__e.Return(True)
return
} else {
tmp12481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1838)
}
__typedArg0 := V1838
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12481 {
tmp12478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1838)
}
__typedArg0 := V1838
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12479 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V1837, tmp12478)


if True == tmp12479 {
__e.Return(True)
return
} else {
tmp12475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1838)
}
__typedArg0 := V1838
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp12476 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V1837, tmp12475)


if True == tmp12476 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


} else {
__e.Return(False)
return
}


}


}, 2)

tmp12484 := Call(__e, ns2_1set, symshen_4occurs_1check_2, tmp12473)


_ = tmp12484

tmp12485 := MakeNative(func(__e *ControlFlow) {
V1839 := __e.Get(1)
_ = V1839
V1840 := __e.Get(2)
_ = V1840
V1841 := __e.Get(3)
_ = V1841
V1842 := __e.Get(4)
_ = V1842
V1843 := __e.Get(5)
_ = V1843
tmp12486 := Call(__e, V1839, V1840)


tmp12487 := Call(__e, tmp12486, V1841)


tmp12488 := Call(__e, tmp12487, V1842)


__e.TailApply(tmp12488, V1843)
return


}, 5)

tmp12489 := Call(__e, ns2_1set, symcall, tmp12485)


_ = tmp12489

tmp12490 := MakeNative(func(__e *ControlFlow) {
V1850 := __e.Get(1)
_ = V1850
V1851 := __e.Get(2)
_ = V1851
V1852 := __e.Get(3)
_ = V1852
V1853 := __e.Get(4)
_ = V1853
V1854 := __e.Get(5)
_ = V1854
__e.TailApply(PrimFunc(symshen_4deref), V1850, V1851)
return
}, 5)

tmp12491 := Call(__e, ns2_1set, symreturn, tmp12490)


_ = tmp12491

tmp12492 := MakeNative(func(__e *ControlFlow) {
V1861 := __e.Get(1)
_ = V1861
V1862 := __e.Get(2)
_ = V1862
V1863 := __e.Get(3)
_ = V1863
V1864 := __e.Get(4)
_ = V1864
V1865 := __e.Get(5)
_ = V1865
if True == V1861 {
__e.TailApply(PrimFunc(symthaw), V1865)
return
} else {
__e.Return(False)
return
}
}, 5)

tmp12494 := Call(__e, ns2_1set, symwhen, tmp12492)


_ = tmp12494

tmp12495 := MakeNative(func(__e *ControlFlow) {
V1866 := __e.Get(1)
_ = V1866
V1867 := __e.Get(2)
_ = V1867
V1868 := __e.Get(3)
_ = V1868
V1869 := __e.Get(4)
_ = V1869
V1870 := __e.Get(5)
_ = V1870
V1871 := __e.Get(6)
_ = V1871
tmp12496 := Call(__e, PrimFunc(symshen_4lazyderef), V1866, V1868)


tmp12497 := Call(__e, PrimFunc(symshen_4lazyderef), V1867, V1868)


__e.TailApply(PrimFunc(symshen_4lzy_a), tmp12496, tmp12497, V1868, V1871)
return


}, 6)

tmp12498 := Call(__e, ns2_1set, symis, tmp12495)


_ = tmp12498

tmp12499 := MakeNative(func(__e *ControlFlow) {
V1872 := __e.Get(1)
_ = V1872
V1873 := __e.Get(2)
_ = V1873
V1874 := __e.Get(3)
_ = V1874
V1875 := __e.Get(4)
_ = V1875
V1876 := __e.Get(5)
_ = V1876
V1877 := __e.Get(6)
_ = V1877
tmp12500 := Call(__e, PrimFunc(symshen_4lazyderef), V1872, V1874)


tmp12501 := Call(__e, PrimFunc(symshen_4lazyderef), V1873, V1874)


__e.TailApply(PrimFunc(symshen_4lzy_a_b), tmp12500, tmp12501, V1874, V1877)
return


}, 6)

tmp12502 := Call(__e, ns2_1set, symis_b, tmp12499)


_ = tmp12502

tmp12503 := MakeNative(func(__e *ControlFlow) {
V1882 := __e.Get(1)
_ = V1882
V1883 := __e.Get(2)
_ = V1883
V1884 := __e.Get(3)
_ = V1884
V1885 := __e.Get(4)
_ = V1885
V1886 := __e.Get(5)
_ = V1886
V1887 := __e.Get(6)
_ = V1887
__e.TailApply(PrimFunc(symshen_4bind_b), V1882, V1883, V1884, V1887)
return
}, 6)

tmp12504 := Call(__e, ns2_1set, symbind, tmp12503)


_ = tmp12504

tmp12505 := MakeNative(func(__e *ControlFlow) {
V1888 := __e.Get(1)
_ = V1888
V1889 := __e.Get(2)
_ = V1889
V1890 := __e.Get(3)
_ = V1890
V1891 := __e.Get(4)
_ = V1891
V1892 := __e.Get(5)
_ = V1892
tmp12507 := Call(__e, PrimFunc(symshen_4lazyderef), V1888, V1889)


tmp12508 := Call(__e, PrimFunc(symshen_4pvar_2), tmp12507)


if True == tmp12508 {
__e.TailApply(PrimFunc(symthaw), V1892)
return
} else {
__e.Return(False)
return
}


}, 5)

tmp12509 := Call(__e, ns2_1set, symvar_2, tmp12505)


_ = tmp12509

tmp12510 := MakeNative(func(__e *ControlFlow) {
V1895 := __e.Get(1)
_ = V1895
__e.Return(MakeString("|prolog vector|"))
return
}, 1)

tmp12511 := Call(__e, ns2_1set, symshen_4print_1prolog_1vector, tmp12510)


_ = tmp12511

tmp12512 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__self5 := __e.Get(5)
__selftop:
V1914 := __self1
_ = V1914
V1915 := __self2
_ = V1915
V1916 := __self3
_ = V1916
V1917 := __self4
_ = V1917
V1918 := __self5
_ = V1918
tmp12524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V1914)
}
__typedArg0 := Nil
__typedArg1 := V1914
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12524 {
__e.Return(False)
return
} else {
tmp12522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V1914)
}
__typedArg0 := V1914
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp12522 {
tmp12513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V1914)
}
__typedArg0 := V1914
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp12514 := Call(__e, tmp12513, V1915)


tmp12515 := Call(__e, tmp12514, V1916)


tmp12516 := Call(__e, tmp12515, V1917)


tmp12517 := Call(__e, tmp12516, V1918)


let__10577 := tmp12517
_ = let__10577

tmp12520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10577, False)
}
__typedArg0 := let__10577
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12520 {
tmp12518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V1914)
}
__typedArg0 := V1914
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symfork) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp12518, V1915, V1916, V1917, V1918
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symfork), tmp12518, V1915, V1916, V1917, V1918)
return


} else {
__e.Return(let__10577)
return
}


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("fork expects a list of literals\n"))
}
__typedArg0 := MakeString("fork expects a list of literals\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 5)

tmp12525 := Call(__e, ns2_1set, symfork, tmp12512)


_ = tmp12525

tmp12526 := MakeNative(func(__e *ControlFlow) {
V1920 := __e.Get(1)
_ = V1920
V1921 := __e.Get(2)
_ = V1921
V1922 := __e.Get(3)
_ = V1922
V1923 := __e.Get(4)
_ = V1923
V1924 := __e.Get(5)
_ = V1924
V1925 := __e.Get(6)
_ = V1925
V1926 := __e.Get(7)
_ = V1926
tmp12532 := Call(__e, PrimFunc(symshen_4unlocked_2), V1924)


if True == tmp12532 {
tmp12527 := Call(__e, PrimFunc(symshen_4newpv), V1923)


let__10578 := tmp12527
_ = let__10578

tmp12528 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12528

tmp12529 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4findall_1h), V1920, V1921, V1922, let__10578, V1923, V1924, V1925, V1926)
return
}, 0)

tmp12530 := Call(__e, PrimFunc(symis), let__10578, Nil, V1923, V1924, V1925, tmp12529)


__e.TailApply(PrimFunc(symshen_4gc), V1923, tmp12530)
return


} else {
__e.Return(False)
return
}


}, 7)

tmp12533 := Call(__e, ns2_1set, symfindall, tmp12526)


_ = tmp12533

tmp12534 := MakeNative(func(__e *ControlFlow) {
V1928 := __e.Get(1)
_ = V1928
V1929 := __e.Get(2)
_ = V1929
V1930 := __e.Get(3)
_ = V1930
V1931 := __e.Get(4)
_ = V1931
V1932 := __e.Get(5)
_ = V1932
V1933 := __e.Get(6)
_ = V1933
V1934 := __e.Get(7)
_ = V1934
V1935 := __e.Get(8)
_ = V1935
tmp12539 := Call(__e, PrimFunc(symshen_4unlocked_2), V1933)


var ifres12535 Obj

if True == tmp12539 {
tmp12536 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12536

tmp12537 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4overbind), V1928, V1931, V1932, V1933, V1934, V1935)
return
}, 0)

tmp12538 := Call(__e, PrimFunc(symcall), V1929, V1932, V1933, V1934, tmp12537)


ifres12535 = tmp12538


} else {
ifres12535 = False


}

let__10579 := ifres12535
_ = let__10579

tmp12544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__10579, False)
}
__typedArg0 := let__10579
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12544 {
tmp12542 := Call(__e, PrimFunc(symshen_4unlocked_2), V1933)


if True == tmp12542 {
tmp12540 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12540

__e.TailApply(PrimFunc(symis_b), V1930, V1931, V1932, V1933, V1934, V1935)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(let__10579)
return
}


}, 8)

tmp12545 := Call(__e, ns2_1set, symshen_4findall_1h, tmp12534)


_ = tmp12545

tmp12546 := MakeNative(func(__e *ControlFlow) {
V1943 := __e.Get(1)
_ = V1943
V1944 := __e.Get(2)
_ = V1944
V1945 := __e.Get(3)
_ = V1945
V1946 := __e.Get(4)
_ = V1946
V1947 := __e.Get(5)
_ = V1947
V1948 := __e.Get(6)
_ = V1948
tmp12547 := Call(__e, PrimFunc(symshen_4deref), V1943, V1945)


tmp12548 := Call(__e, PrimFunc(symshen_4lazyderef), V1944, V1945)


tmp12549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12547, tmp12548)
}
__typedArg0 := tmp12547
__typedArg1 := tmp12548
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12550 := Call(__e, PrimFunc(symshen_4bindv), V1944, tmp12549, V1945)


_ = tmp12550

__e.Return(False)
return


}, 6)

tmp12551 := Call(__e, ns2_1set, symshen_4overbind, tmp12546)


_ = tmp12551

tmp12552 := MakeNative(func(__e *ControlFlow) {
V1951 := __e.Get(1)
_ = V1951
tmp12556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V1951)
}
__typedArg0 := sym_7
__typedArg1 := V1951
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12556 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_doccurs_d, True)
}
__typedArg0 := symshen_4_doccurs_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
tmp12554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V1951)
}
__typedArg0 := sym_1
__typedArg1 := V1951
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp12554 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_doccurs_d, False)
}
__typedArg0 := symshen_4_doccurs_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("occurs-check expects a + or a -.\n"))
}
__typedArg0 := MakeString("occurs-check expects a + or a -.\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

__e.TailApply(ns2_1set, symoccurs_1check, tmp12552)
return




}, 0)

