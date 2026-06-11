package main

import . "github.com/tiancaiamao/shen-go/kl"

var SequentMain = MakeNative(func(__e *ControlFlow) {
tmp10582 := MakeNative(func(__e *ControlFlow) {
V3365 := __e.Get(1)
_ = V3365
tmp10583 := MakeNative(func(__e *ControlFlow) {
W3366 := __e.Get(1)
_ = W3366
tmp10585 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3366)


if True == tmp10585 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3366)
return
}


}, 1)

tmp10605 := PrimIsPair(V3365)

var ifres10586 Obj

if True == tmp10605 {
tmp10587 := MakeNative(func(__e *ControlFlow) {
W3367 := __e.Get(1)
_ = W3367
tmp10588 := MakeNative(func(__e *ControlFlow) {
W3368 := __e.Get(1)
_ = W3368
tmp10589 := MakeNative(func(__e *ControlFlow) {
W3369 := __e.Get(1)
_ = W3369
tmp10599 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3369)


if True == tmp10599 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10590 := MakeNative(func(__e *ControlFlow) {
W3370 := __e.Get(1)
_ = W3370
tmp10591 := MakeNative(func(__e *ControlFlow) {
W3371 := __e.Get(1)
_ = W3371
tmp10592 := MakeNative(func(__e *ControlFlow) {
W3372 := __e.Get(1)
_ = W3372
tmp10593 := Call(__e, PrimFunc(symfn), W3367)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), W3367, tmp10593)
return


}, 1)

tmp10594 := Call(__e, PrimFunc(symshen_4rules_1_6prolog), W3367, W3370)


tmp10595 := Call(__e, tmp10592, tmp10594)


__e.TailApply(PrimFunc(symshen_4comb), W3371, tmp10595)
return


}, 1)

tmp10596 := Call(__e, PrimFunc(symshen_4in_1_6), W3369)


__e.TailApply(tmp10591, tmp10596)
return


}, 1)

tmp10597 := Call(__e, PrimFunc(symshen_4_5_1out), W3369)


__e.TailApply(tmp10590, tmp10597)
return


}


}, 1)

tmp10600 := Call(__e, PrimFunc(symshen_4_5datatype_1rules_6), W3368)


__e.TailApply(tmp10589, tmp10600)
return


}, 1)

tmp10601 := Call(__e, PrimFunc(symtail), V3365)


__e.TailApply(tmp10588, tmp10601)
return


}, 1)

tmp10602 := Call(__e, PrimFunc(symhead), V3365)


tmp10603 := Call(__e, tmp10587, tmp10602)


ifres10586 = tmp10603


} else {
tmp10604 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10586 = tmp10604


}

__e.TailApply(tmp10583, ifres10586)
return


}, 1)

tmp10606 := Call(__e, ns2_1set, symshen_4_5datatype_6, tmp10582)


_ = tmp10606

tmp10607 := MakeNative(func(__e *ControlFlow) {
V3373 := __e.Get(1)
_ = V3373
V3374 := __e.Get(2)
_ = V3374
tmp10608 := PrimValue(symshen_4_ddatatypes_d)

tmp10609 := Call(__e, PrimFunc(symshen_4assoc_1_6), V3373, V3374, tmp10608)


tmp10610 := PrimSet(symshen_4_ddatatypes_d, tmp10609)

_ = tmp10610

tmp10611 := PrimValue(symshen_4_dalldatatypes_d)

tmp10612 := Call(__e, PrimFunc(symshen_4assoc_1_6), V3373, V3374, tmp10611)


tmp10613 := PrimSet(symshen_4_dalldatatypes_d, tmp10612)

_ = tmp10613

__e.Return(V3373)
return


}, 2)

tmp10614 := Call(__e, ns2_1set, symshen_4remember_1datatype, tmp10607)


_ = tmp10614

tmp10615 := MakeNative(func(__e *ControlFlow) {
V3375 := __e.Get(1)
_ = V3375
tmp10616 := MakeNative(func(__e *ControlFlow) {
W3376 := __e.Get(1)
_ = W3376
tmp10635 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3376)


if True == tmp10635 {
tmp10617 := MakeNative(func(__e *ControlFlow) {
W3383 := __e.Get(1)
_ = W3383
tmp10619 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3383)


if True == tmp10619 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3383)
return
}


}, 1)

tmp10620 := MakeNative(func(__e *ControlFlow) {
W3384 := __e.Get(1)
_ = W3384
tmp10631 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3384)


if True == tmp10631 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10621 := MakeNative(func(__e *ControlFlow) {
W3385 := __e.Get(1)
_ = W3385
tmp10622 := MakeNative(func(__e *ControlFlow) {
W3386 := __e.Get(1)
_ = W3386
tmp10627 := Call(__e, PrimFunc(symempty_2), W3385)


var ifres10623 Obj

if True == tmp10627 {
ifres10623 = Nil


} else {
tmp10624 := Call(__e, PrimFunc(symshen_4app), W3385, MakeString("\n ..."), symshen_4r)


tmp10625 := PrimStringConcat(MakeString("datatype syntax error here:\n "), tmp10624)

tmp10626 := PrimSimpleError(tmp10625)

ifres10623 = tmp10626


}

__e.TailApply(PrimFunc(symshen_4comb), W3386, ifres10623)
return


}, 1)

tmp10628 := Call(__e, PrimFunc(symshen_4in_1_6), W3384)


__e.TailApply(tmp10622, tmp10628)
return


}, 1)

tmp10629 := Call(__e, PrimFunc(symshen_4_5_1out), W3384)


__e.TailApply(tmp10621, tmp10629)
return


}


}, 1)

tmp10632 := Call(__e, PrimFunc(sym_5_b_6), V3375)


tmp10633 := Call(__e, tmp10620, tmp10632)


__e.TailApply(tmp10617, tmp10633)
return


} else {
__e.Return(W3376)
return
}


}, 1)

tmp10636 := MakeNative(func(__e *ControlFlow) {
W3377 := __e.Get(1)
_ = W3377
tmp10651 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3377)


if True == tmp10651 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10637 := MakeNative(func(__e *ControlFlow) {
W3378 := __e.Get(1)
_ = W3378
tmp10638 := MakeNative(func(__e *ControlFlow) {
W3379 := __e.Get(1)
_ = W3379
tmp10639 := MakeNative(func(__e *ControlFlow) {
W3380 := __e.Get(1)
_ = W3380
tmp10646 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3380)


if True == tmp10646 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10640 := MakeNative(func(__e *ControlFlow) {
W3381 := __e.Get(1)
_ = W3381
tmp10641 := MakeNative(func(__e *ControlFlow) {
W3382 := __e.Get(1)
_ = W3382
tmp10642 := Call(__e, PrimFunc(symappend), W3378, W3381)


__e.TailApply(PrimFunc(symshen_4comb), W3382, tmp10642)
return


}, 1)

tmp10643 := Call(__e, PrimFunc(symshen_4in_1_6), W3380)


__e.TailApply(tmp10641, tmp10643)
return


}, 1)

tmp10644 := Call(__e, PrimFunc(symshen_4_5_1out), W3380)


__e.TailApply(tmp10640, tmp10644)
return


}


}, 1)

tmp10647 := Call(__e, PrimFunc(symshen_4_5datatype_1rules_6), W3379)


__e.TailApply(tmp10639, tmp10647)
return


}, 1)

tmp10648 := Call(__e, PrimFunc(symshen_4in_1_6), W3377)


__e.TailApply(tmp10638, tmp10648)
return


}, 1)

tmp10649 := Call(__e, PrimFunc(symshen_4_5_1out), W3377)


__e.TailApply(tmp10637, tmp10649)
return


}


}, 1)

tmp10652 := Call(__e, PrimFunc(symshen_4_5datatype_1rule_6), V3375)


tmp10653 := Call(__e, tmp10636, tmp10652)


__e.TailApply(tmp10616, tmp10653)
return


}, 1)

tmp10654 := Call(__e, ns2_1set, symshen_4_5datatype_1rules_6, tmp10615)


_ = tmp10654

tmp10655 := MakeNative(func(__e *ControlFlow) {
V3387 := __e.Get(1)
_ = V3387
tmp10656 := MakeNative(func(__e *ControlFlow) {
W3388 := __e.Get(1)
_ = W3388
tmp10670 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3388)


if True == tmp10670 {
tmp10657 := MakeNative(func(__e *ControlFlow) {
W3392 := __e.Get(1)
_ = W3392
tmp10659 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3392)


if True == tmp10659 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3392)
return
}


}, 1)

tmp10660 := MakeNative(func(__e *ControlFlow) {
W3393 := __e.Get(1)
_ = W3393
tmp10666 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3393)


if True == tmp10666 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10661 := MakeNative(func(__e *ControlFlow) {
W3394 := __e.Get(1)
_ = W3394
tmp10662 := MakeNative(func(__e *ControlFlow) {
W3395 := __e.Get(1)
_ = W3395
__e.TailApply(PrimFunc(symshen_4comb), W3395, W3394)
return
}, 1)

tmp10663 := Call(__e, PrimFunc(symshen_4in_1_6), W3393)


__e.TailApply(tmp10662, tmp10663)
return


}, 1)

tmp10664 := Call(__e, PrimFunc(symshen_4_5_1out), W3393)


__e.TailApply(tmp10661, tmp10664)
return


}


}, 1)

tmp10667 := Call(__e, PrimFunc(symshen_4_5double_6), V3387)


tmp10668 := Call(__e, tmp10660, tmp10667)


__e.TailApply(tmp10657, tmp10668)
return


} else {
__e.Return(W3388)
return
}


}, 1)

tmp10671 := MakeNative(func(__e *ControlFlow) {
W3389 := __e.Get(1)
_ = W3389
tmp10677 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3389)


if True == tmp10677 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10672 := MakeNative(func(__e *ControlFlow) {
W3390 := __e.Get(1)
_ = W3390
tmp10673 := MakeNative(func(__e *ControlFlow) {
W3391 := __e.Get(1)
_ = W3391
__e.TailApply(PrimFunc(symshen_4comb), W3391, W3390)
return
}, 1)

tmp10674 := Call(__e, PrimFunc(symshen_4in_1_6), W3389)


__e.TailApply(tmp10673, tmp10674)
return


}, 1)

tmp10675 := Call(__e, PrimFunc(symshen_4_5_1out), W3389)


__e.TailApply(tmp10672, tmp10675)
return


}


}, 1)

tmp10678 := Call(__e, PrimFunc(symshen_4_5single_6), V3387)


tmp10679 := Call(__e, tmp10671, tmp10678)


__e.TailApply(tmp10656, tmp10679)
return


}, 1)

tmp10680 := Call(__e, ns2_1set, symshen_4_5datatype_1rule_6, tmp10655)


_ = tmp10680

tmp10681 := MakeNative(func(__e *ControlFlow) {
V3396 := __e.Get(1)
_ = V3396
tmp10682 := MakeNative(func(__e *ControlFlow) {
W3397 := __e.Get(1)
_ = W3397
tmp10684 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3397)


if True == tmp10684 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3397)
return
}


}, 1)

tmp10685 := MakeNative(func(__e *ControlFlow) {
W3398 := __e.Get(1)
_ = W3398
tmp10723 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3398)


if True == tmp10723 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10686 := MakeNative(func(__e *ControlFlow) {
W3399 := __e.Get(1)
_ = W3399
tmp10687 := MakeNative(func(__e *ControlFlow) {
W3400 := __e.Get(1)
_ = W3400
tmp10688 := MakeNative(func(__e *ControlFlow) {
W3401 := __e.Get(1)
_ = W3401
tmp10718 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3401)


if True == tmp10718 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10689 := MakeNative(func(__e *ControlFlow) {
W3402 := __e.Get(1)
_ = W3402
tmp10690 := MakeNative(func(__e *ControlFlow) {
W3403 := __e.Get(1)
_ = W3403
tmp10691 := MakeNative(func(__e *ControlFlow) {
W3404 := __e.Get(1)
_ = W3404
tmp10713 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3404)


if True == tmp10713 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10692 := MakeNative(func(__e *ControlFlow) {
W3405 := __e.Get(1)
_ = W3405
tmp10693 := MakeNative(func(__e *ControlFlow) {
W3406 := __e.Get(1)
_ = W3406
tmp10709 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3406)


if True == tmp10709 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10694 := MakeNative(func(__e *ControlFlow) {
W3407 := __e.Get(1)
_ = W3407
tmp10695 := MakeNative(func(__e *ControlFlow) {
W3408 := __e.Get(1)
_ = W3408
tmp10696 := MakeNative(func(__e *ControlFlow) {
W3409 := __e.Get(1)
_ = W3409
tmp10704 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3409)


if True == tmp10704 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10697 := MakeNative(func(__e *ControlFlow) {
W3410 := __e.Get(1)
_ = W3410
tmp10698 := PrimCons(W3407, Nil)

tmp10699 := PrimCons(W3402, tmp10698)

tmp10700 := PrimCons(W3399, tmp10699)

tmp10701 := PrimCons(tmp10700, Nil)

__e.TailApply(PrimFunc(symshen_4comb), W3410, tmp10701)
return


}, 1)

tmp10702 := Call(__e, PrimFunc(symshen_4in_1_6), W3409)


__e.TailApply(tmp10697, tmp10702)
return


}


}, 1)

tmp10705 := Call(__e, PrimFunc(symshen_4_5sc_6), W3408)


__e.TailApply(tmp10696, tmp10705)
return


}, 1)

tmp10706 := Call(__e, PrimFunc(symshen_4in_1_6), W3406)


__e.TailApply(tmp10695, tmp10706)
return


}, 1)

tmp10707 := Call(__e, PrimFunc(symshen_4_5_1out), W3406)


__e.TailApply(tmp10694, tmp10707)
return


}


}, 1)

tmp10710 := Call(__e, PrimFunc(symshen_4_5conc_6), W3405)


__e.TailApply(tmp10693, tmp10710)
return


}, 1)

tmp10711 := Call(__e, PrimFunc(symshen_4in_1_6), W3404)


__e.TailApply(tmp10692, tmp10711)
return


}


}, 1)

tmp10714 := Call(__e, PrimFunc(symshen_4_5sng_6), W3403)


__e.TailApply(tmp10691, tmp10714)
return


}, 1)

tmp10715 := Call(__e, PrimFunc(symshen_4in_1_6), W3401)


__e.TailApply(tmp10690, tmp10715)
return


}, 1)

tmp10716 := Call(__e, PrimFunc(symshen_4_5_1out), W3401)


__e.TailApply(tmp10689, tmp10716)
return


}


}, 1)

tmp10719 := Call(__e, PrimFunc(symshen_4_5prems_6), W3400)


__e.TailApply(tmp10688, tmp10719)
return


}, 1)

tmp10720 := Call(__e, PrimFunc(symshen_4in_1_6), W3398)


__e.TailApply(tmp10687, tmp10720)
return


}, 1)

tmp10721 := Call(__e, PrimFunc(symshen_4_5_1out), W3398)


__e.TailApply(tmp10686, tmp10721)
return


}


}, 1)

tmp10724 := Call(__e, PrimFunc(symshen_4_5sides_6), V3396)


tmp10725 := Call(__e, tmp10685, tmp10724)


__e.TailApply(tmp10682, tmp10725)
return


}, 1)

tmp10726 := Call(__e, ns2_1set, symshen_4_5single_6, tmp10681)


_ = tmp10726

tmp10727 := MakeNative(func(__e *ControlFlow) {
V3411 := __e.Get(1)
_ = V3411
tmp10728 := MakeNative(func(__e *ControlFlow) {
W3412 := __e.Get(1)
_ = W3412
tmp10730 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3412)


if True == tmp10730 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3412)
return
}


}, 1)

tmp10731 := MakeNative(func(__e *ControlFlow) {
W3413 := __e.Get(1)
_ = W3413
tmp10768 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3413)


if True == tmp10768 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10732 := MakeNative(func(__e *ControlFlow) {
W3414 := __e.Get(1)
_ = W3414
tmp10733 := MakeNative(func(__e *ControlFlow) {
W3415 := __e.Get(1)
_ = W3415
tmp10734 := MakeNative(func(__e *ControlFlow) {
W3416 := __e.Get(1)
_ = W3416
tmp10763 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3416)


if True == tmp10763 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10735 := MakeNative(func(__e *ControlFlow) {
W3417 := __e.Get(1)
_ = W3417
tmp10736 := MakeNative(func(__e *ControlFlow) {
W3418 := __e.Get(1)
_ = W3418
tmp10737 := MakeNative(func(__e *ControlFlow) {
W3419 := __e.Get(1)
_ = W3419
tmp10758 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3419)


if True == tmp10758 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10738 := MakeNative(func(__e *ControlFlow) {
W3420 := __e.Get(1)
_ = W3420
tmp10739 := MakeNative(func(__e *ControlFlow) {
W3421 := __e.Get(1)
_ = W3421
tmp10754 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3421)


if True == tmp10754 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10740 := MakeNative(func(__e *ControlFlow) {
W3422 := __e.Get(1)
_ = W3422
tmp10741 := MakeNative(func(__e *ControlFlow) {
W3423 := __e.Get(1)
_ = W3423
tmp10742 := MakeNative(func(__e *ControlFlow) {
W3424 := __e.Get(1)
_ = W3424
tmp10749 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3424)


if True == tmp10749 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10743 := MakeNative(func(__e *ControlFlow) {
W3425 := __e.Get(1)
_ = W3425
tmp10744 := PrimCons(W3422, Nil)

tmp10745 := PrimCons(Nil, tmp10744)

tmp10746 := Call(__e, PrimFunc(symshen_4lr_1rule), W3414, W3417, tmp10745)


__e.TailApply(PrimFunc(symshen_4comb), W3425, tmp10746)
return


}, 1)

tmp10747 := Call(__e, PrimFunc(symshen_4in_1_6), W3424)


__e.TailApply(tmp10743, tmp10747)
return


}


}, 1)

tmp10750 := Call(__e, PrimFunc(symshen_4_5sc_6), W3423)


__e.TailApply(tmp10742, tmp10750)
return


}, 1)

tmp10751 := Call(__e, PrimFunc(symshen_4in_1_6), W3421)


__e.TailApply(tmp10741, tmp10751)
return


}, 1)

tmp10752 := Call(__e, PrimFunc(symshen_4_5_1out), W3421)


__e.TailApply(tmp10740, tmp10752)
return


}


}, 1)

tmp10755 := Call(__e, PrimFunc(symshen_4_5formula_6), W3420)


__e.TailApply(tmp10739, tmp10755)
return


}, 1)

tmp10756 := Call(__e, PrimFunc(symshen_4in_1_6), W3419)


__e.TailApply(tmp10738, tmp10756)
return


}


}, 1)

tmp10759 := Call(__e, PrimFunc(symshen_4_5dbl_6), W3418)


__e.TailApply(tmp10737, tmp10759)
return


}, 1)

tmp10760 := Call(__e, PrimFunc(symshen_4in_1_6), W3416)


__e.TailApply(tmp10736, tmp10760)
return


}, 1)

tmp10761 := Call(__e, PrimFunc(symshen_4_5_1out), W3416)


__e.TailApply(tmp10735, tmp10761)
return


}


}, 1)

tmp10764 := Call(__e, PrimFunc(symshen_4_5formulae_6), W3415)


__e.TailApply(tmp10734, tmp10764)
return


}, 1)

tmp10765 := Call(__e, PrimFunc(symshen_4in_1_6), W3413)


__e.TailApply(tmp10733, tmp10765)
return


}, 1)

tmp10766 := Call(__e, PrimFunc(symshen_4_5_1out), W3413)


__e.TailApply(tmp10732, tmp10766)
return


}


}, 1)

tmp10769 := Call(__e, PrimFunc(symshen_4_5sides_6), V3411)


tmp10770 := Call(__e, tmp10731, tmp10769)


__e.TailApply(tmp10728, tmp10770)
return


}, 1)

tmp10771 := Call(__e, ns2_1set, symshen_4_5double_6, tmp10727)


_ = tmp10771

tmp10772 := MakeNative(func(__e *ControlFlow) {
V3426 := __e.Get(1)
_ = V3426
tmp10773 := MakeNative(func(__e *ControlFlow) {
W3427 := __e.Get(1)
_ = W3427
tmp10796 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3427)


if True == tmp10796 {
tmp10774 := MakeNative(func(__e *ControlFlow) {
W3436 := __e.Get(1)
_ = W3436
tmp10776 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3436)


if True == tmp10776 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3436)
return
}


}, 1)

tmp10777 := MakeNative(func(__e *ControlFlow) {
W3437 := __e.Get(1)
_ = W3437
tmp10792 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3437)


if True == tmp10792 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10778 := MakeNative(func(__e *ControlFlow) {
W3438 := __e.Get(1)
_ = W3438
tmp10779 := MakeNative(func(__e *ControlFlow) {
W3439 := __e.Get(1)
_ = W3439
tmp10780 := MakeNative(func(__e *ControlFlow) {
W3440 := __e.Get(1)
_ = W3440
tmp10787 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3440)


if True == tmp10787 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10781 := MakeNative(func(__e *ControlFlow) {
W3441 := __e.Get(1)
_ = W3441
tmp10782 := PrimCons(W3438, Nil)

tmp10783 := PrimCons(Nil, tmp10782)

tmp10784 := PrimCons(tmp10783, Nil)

__e.TailApply(PrimFunc(symshen_4comb), W3441, tmp10784)
return


}, 1)

tmp10785 := Call(__e, PrimFunc(symshen_4in_1_6), W3440)


__e.TailApply(tmp10781, tmp10785)
return


}


}, 1)

tmp10788 := Call(__e, PrimFunc(symshen_4_5sc_6), W3439)


__e.TailApply(tmp10780, tmp10788)
return


}, 1)

tmp10789 := Call(__e, PrimFunc(symshen_4in_1_6), W3437)


__e.TailApply(tmp10779, tmp10789)
return


}, 1)

tmp10790 := Call(__e, PrimFunc(symshen_4_5_1out), W3437)


__e.TailApply(tmp10778, tmp10790)
return


}


}, 1)

tmp10793 := Call(__e, PrimFunc(symshen_4_5formula_6), V3426)


tmp10794 := Call(__e, tmp10777, tmp10793)


__e.TailApply(tmp10774, tmp10794)
return


} else {
__e.Return(W3427)
return
}


}, 1)

tmp10797 := MakeNative(func(__e *ControlFlow) {
W3428 := __e.Get(1)
_ = W3428
tmp10820 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3428)


if True == tmp10820 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10798 := MakeNative(func(__e *ControlFlow) {
W3429 := __e.Get(1)
_ = W3429
tmp10799 := MakeNative(func(__e *ControlFlow) {
W3430 := __e.Get(1)
_ = W3430
tmp10800 := MakeNative(func(__e *ControlFlow) {
W3431 := __e.Get(1)
_ = W3431
tmp10815 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3431)


if True == tmp10815 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10801 := MakeNative(func(__e *ControlFlow) {
W3432 := __e.Get(1)
_ = W3432
tmp10802 := MakeNative(func(__e *ControlFlow) {
W3433 := __e.Get(1)
_ = W3433
tmp10811 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3433)


if True == tmp10811 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10803 := MakeNative(func(__e *ControlFlow) {
W3434 := __e.Get(1)
_ = W3434
tmp10804 := MakeNative(func(__e *ControlFlow) {
W3435 := __e.Get(1)
_ = W3435
tmp10805 := PrimCons(W3429, Nil)

tmp10806 := PrimCons(Nil, tmp10805)

tmp10807 := PrimCons(tmp10806, W3434)

__e.TailApply(PrimFunc(symshen_4comb), W3435, tmp10807)
return


}, 1)

tmp10808 := Call(__e, PrimFunc(symshen_4in_1_6), W3433)


__e.TailApply(tmp10804, tmp10808)
return


}, 1)

tmp10809 := Call(__e, PrimFunc(symshen_4_5_1out), W3433)


__e.TailApply(tmp10803, tmp10809)
return


}


}, 1)

tmp10812 := Call(__e, PrimFunc(symshen_4_5formulae_6), W3432)


__e.TailApply(tmp10802, tmp10812)
return


}, 1)

tmp10813 := Call(__e, PrimFunc(symshen_4in_1_6), W3431)


__e.TailApply(tmp10801, tmp10813)
return


}


}, 1)

tmp10816 := Call(__e, PrimFunc(symshen_4_5sc_6), W3430)


__e.TailApply(tmp10800, tmp10816)
return


}, 1)

tmp10817 := Call(__e, PrimFunc(symshen_4in_1_6), W3428)


__e.TailApply(tmp10799, tmp10817)
return


}, 1)

tmp10818 := Call(__e, PrimFunc(symshen_4_5_1out), W3428)


__e.TailApply(tmp10798, tmp10818)
return


}


}, 1)

tmp10821 := Call(__e, PrimFunc(symshen_4_5formula_6), V3426)


tmp10822 := Call(__e, tmp10797, tmp10821)


__e.TailApply(tmp10773, tmp10822)
return


}, 1)

tmp10823 := Call(__e, ns2_1set, symshen_4_5formulae_6, tmp10772)


_ = tmp10823

tmp10824 := MakeNative(func(__e *ControlFlow) {
V3442 := __e.Get(1)
_ = V3442
tmp10825 := MakeNative(func(__e *ControlFlow) {
W3443 := __e.Get(1)
_ = W3443
tmp10841 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3443)


if True == tmp10841 {
tmp10826 := MakeNative(func(__e *ControlFlow) {
W3451 := __e.Get(1)
_ = W3451
tmp10828 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3451)


if True == tmp10828 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3451)
return
}


}, 1)

tmp10829 := MakeNative(func(__e *ControlFlow) {
W3452 := __e.Get(1)
_ = W3452
tmp10837 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3452)


if True == tmp10837 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10830 := MakeNative(func(__e *ControlFlow) {
W3453 := __e.Get(1)
_ = W3453
tmp10831 := MakeNative(func(__e *ControlFlow) {
W3454 := __e.Get(1)
_ = W3454
tmp10832 := PrimCons(W3453, Nil)

tmp10833 := PrimCons(Nil, tmp10832)

__e.TailApply(PrimFunc(symshen_4comb), W3454, tmp10833)
return


}, 1)

tmp10834 := Call(__e, PrimFunc(symshen_4in_1_6), W3452)


__e.TailApply(tmp10831, tmp10834)
return


}, 1)

tmp10835 := Call(__e, PrimFunc(symshen_4_5_1out), W3452)


__e.TailApply(tmp10830, tmp10835)
return


}


}, 1)

tmp10838 := Call(__e, PrimFunc(symshen_4_5formula_6), V3442)


tmp10839 := Call(__e, tmp10829, tmp10838)


__e.TailApply(tmp10826, tmp10839)
return


} else {
__e.Return(W3443)
return
}


}, 1)

tmp10842 := MakeNative(func(__e *ControlFlow) {
W3444 := __e.Get(1)
_ = W3444
tmp10862 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3444)


if True == tmp10862 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10843 := MakeNative(func(__e *ControlFlow) {
W3445 := __e.Get(1)
_ = W3445
tmp10844 := MakeNative(func(__e *ControlFlow) {
W3446 := __e.Get(1)
_ = W3446
tmp10858 := Call(__e, PrimFunc(symshen_4hds_a_2), W3446, sym_6_6)


if True == tmp10858 {
tmp10845 := MakeNative(func(__e *ControlFlow) {
W3447 := __e.Get(1)
_ = W3447
tmp10846 := MakeNative(func(__e *ControlFlow) {
W3448 := __e.Get(1)
_ = W3448
tmp10854 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3448)


if True == tmp10854 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10847 := MakeNative(func(__e *ControlFlow) {
W3449 := __e.Get(1)
_ = W3449
tmp10848 := MakeNative(func(__e *ControlFlow) {
W3450 := __e.Get(1)
_ = W3450
tmp10849 := PrimCons(W3449, Nil)

tmp10850 := PrimCons(W3445, tmp10849)

__e.TailApply(PrimFunc(symshen_4comb), W3450, tmp10850)
return


}, 1)

tmp10851 := Call(__e, PrimFunc(symshen_4in_1_6), W3448)


__e.TailApply(tmp10848, tmp10851)
return


}, 1)

tmp10852 := Call(__e, PrimFunc(symshen_4_5_1out), W3448)


__e.TailApply(tmp10847, tmp10852)
return


}


}, 1)

tmp10855 := Call(__e, PrimFunc(symshen_4_5formula_6), W3447)


__e.TailApply(tmp10846, tmp10855)
return


}, 1)

tmp10856 := Call(__e, PrimFunc(symtail), W3446)


__e.TailApply(tmp10845, tmp10856)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10859 := Call(__e, PrimFunc(symshen_4in_1_6), W3444)


__e.TailApply(tmp10844, tmp10859)
return


}, 1)

tmp10860 := Call(__e, PrimFunc(symshen_4_5_1out), W3444)


__e.TailApply(tmp10843, tmp10860)
return


}


}, 1)

tmp10863 := Call(__e, PrimFunc(symshen_4_5ass_6), V3442)


tmp10864 := Call(__e, tmp10842, tmp10863)


__e.TailApply(tmp10825, tmp10864)
return


}, 1)

tmp10865 := Call(__e, ns2_1set, symshen_4_5conc_6, tmp10824)


_ = tmp10865

tmp10866 := MakeNative(func(__e *ControlFlow) {
V3455 := __e.Get(1)
_ = V3455
tmp10867 := MakeNative(func(__e *ControlFlow) {
W3456 := __e.Get(1)
_ = W3456
tmp10879 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3456)


if True == tmp10879 {
tmp10868 := MakeNative(func(__e *ControlFlow) {
W3465 := __e.Get(1)
_ = W3465
tmp10870 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3465)


if True == tmp10870 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3465)
return
}


}, 1)

tmp10871 := MakeNative(func(__e *ControlFlow) {
W3466 := __e.Get(1)
_ = W3466
tmp10875 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3466)


if True == tmp10875 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10872 := MakeNative(func(__e *ControlFlow) {
W3467 := __e.Get(1)
_ = W3467
__e.TailApply(PrimFunc(symshen_4comb), W3467, Nil)
return
}, 1)

tmp10873 := Call(__e, PrimFunc(symshen_4in_1_6), W3466)


__e.TailApply(tmp10872, tmp10873)
return


}


}, 1)

tmp10876 := Call(__e, PrimFunc(sym_5e_6), V3455)


tmp10877 := Call(__e, tmp10871, tmp10876)


__e.TailApply(tmp10868, tmp10877)
return


} else {
__e.Return(W3456)
return
}


}, 1)

tmp10880 := MakeNative(func(__e *ControlFlow) {
W3457 := __e.Get(1)
_ = W3457
tmp10901 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3457)


if True == tmp10901 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10881 := MakeNative(func(__e *ControlFlow) {
W3458 := __e.Get(1)
_ = W3458
tmp10882 := MakeNative(func(__e *ControlFlow) {
W3459 := __e.Get(1)
_ = W3459
tmp10883 := MakeNative(func(__e *ControlFlow) {
W3460 := __e.Get(1)
_ = W3460
tmp10896 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3460)


if True == tmp10896 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10884 := MakeNative(func(__e *ControlFlow) {
W3461 := __e.Get(1)
_ = W3461
tmp10885 := MakeNative(func(__e *ControlFlow) {
W3462 := __e.Get(1)
_ = W3462
tmp10892 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3462)


if True == tmp10892 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10886 := MakeNative(func(__e *ControlFlow) {
W3463 := __e.Get(1)
_ = W3463
tmp10887 := MakeNative(func(__e *ControlFlow) {
W3464 := __e.Get(1)
_ = W3464
tmp10888 := PrimCons(W3458, W3463)

__e.TailApply(PrimFunc(symshen_4comb), W3464, tmp10888)
return


}, 1)

tmp10889 := Call(__e, PrimFunc(symshen_4in_1_6), W3462)


__e.TailApply(tmp10887, tmp10889)
return


}, 1)

tmp10890 := Call(__e, PrimFunc(symshen_4_5_1out), W3462)


__e.TailApply(tmp10886, tmp10890)
return


}


}, 1)

tmp10893 := Call(__e, PrimFunc(symshen_4_5prems_6), W3461)


__e.TailApply(tmp10885, tmp10893)
return


}, 1)

tmp10894 := Call(__e, PrimFunc(symshen_4in_1_6), W3460)


__e.TailApply(tmp10884, tmp10894)
return


}


}, 1)

tmp10897 := Call(__e, PrimFunc(symshen_4_5sc_6), W3459)


__e.TailApply(tmp10883, tmp10897)
return


}, 1)

tmp10898 := Call(__e, PrimFunc(symshen_4in_1_6), W3457)


__e.TailApply(tmp10882, tmp10898)
return


}, 1)

tmp10899 := Call(__e, PrimFunc(symshen_4_5_1out), W3457)


__e.TailApply(tmp10881, tmp10899)
return


}


}, 1)

tmp10902 := Call(__e, PrimFunc(symshen_4_5prem_6), V3455)


tmp10903 := Call(__e, tmp10880, tmp10902)


__e.TailApply(tmp10867, tmp10903)
return


}, 1)

tmp10904 := Call(__e, ns2_1set, symshen_4_5prems_6, tmp10866)


_ = tmp10904

tmp10905 := MakeNative(func(__e *ControlFlow) {
V3468 := __e.Get(1)
_ = V3468
tmp10906 := MakeNative(func(__e *ControlFlow) {
W3469 := __e.Get(1)
_ = W3469
tmp10948 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3469)


if True == tmp10948 {
tmp10907 := MakeNative(func(__e *ControlFlow) {
W3471 := __e.Get(1)
_ = W3471
tmp10923 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3471)


if True == tmp10923 {
tmp10908 := MakeNative(func(__e *ControlFlow) {
W3479 := __e.Get(1)
_ = W3479
tmp10910 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3479)


if True == tmp10910 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3479)
return
}


}, 1)

tmp10911 := MakeNative(func(__e *ControlFlow) {
W3480 := __e.Get(1)
_ = W3480
tmp10919 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3480)


if True == tmp10919 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10912 := MakeNative(func(__e *ControlFlow) {
W3481 := __e.Get(1)
_ = W3481
tmp10913 := MakeNative(func(__e *ControlFlow) {
W3482 := __e.Get(1)
_ = W3482
tmp10914 := PrimCons(W3481, Nil)

tmp10915 := PrimCons(Nil, tmp10914)

__e.TailApply(PrimFunc(symshen_4comb), W3482, tmp10915)
return


}, 1)

tmp10916 := Call(__e, PrimFunc(symshen_4in_1_6), W3480)


__e.TailApply(tmp10913, tmp10916)
return


}, 1)

tmp10917 := Call(__e, PrimFunc(symshen_4_5_1out), W3480)


__e.TailApply(tmp10912, tmp10917)
return


}


}, 1)

tmp10920 := Call(__e, PrimFunc(symshen_4_5formula_6), V3468)


tmp10921 := Call(__e, tmp10911, tmp10920)


__e.TailApply(tmp10908, tmp10921)
return


} else {
__e.Return(W3471)
return
}


}, 1)

tmp10924 := MakeNative(func(__e *ControlFlow) {
W3472 := __e.Get(1)
_ = W3472
tmp10944 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3472)


if True == tmp10944 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10925 := MakeNative(func(__e *ControlFlow) {
W3473 := __e.Get(1)
_ = W3473
tmp10926 := MakeNative(func(__e *ControlFlow) {
W3474 := __e.Get(1)
_ = W3474
tmp10940 := Call(__e, PrimFunc(symshen_4hds_a_2), W3474, sym_6_6)


if True == tmp10940 {
tmp10927 := MakeNative(func(__e *ControlFlow) {
W3475 := __e.Get(1)
_ = W3475
tmp10928 := MakeNative(func(__e *ControlFlow) {
W3476 := __e.Get(1)
_ = W3476
tmp10936 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3476)


if True == tmp10936 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10929 := MakeNative(func(__e *ControlFlow) {
W3477 := __e.Get(1)
_ = W3477
tmp10930 := MakeNative(func(__e *ControlFlow) {
W3478 := __e.Get(1)
_ = W3478
tmp10931 := PrimCons(W3477, Nil)

tmp10932 := PrimCons(W3473, tmp10931)

__e.TailApply(PrimFunc(symshen_4comb), W3478, tmp10932)
return


}, 1)

tmp10933 := Call(__e, PrimFunc(symshen_4in_1_6), W3476)


__e.TailApply(tmp10930, tmp10933)
return


}, 1)

tmp10934 := Call(__e, PrimFunc(symshen_4_5_1out), W3476)


__e.TailApply(tmp10929, tmp10934)
return


}


}, 1)

tmp10937 := Call(__e, PrimFunc(symshen_4_5formula_6), W3475)


__e.TailApply(tmp10928, tmp10937)
return


}, 1)

tmp10938 := Call(__e, PrimFunc(symtail), W3474)


__e.TailApply(tmp10927, tmp10938)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp10941 := Call(__e, PrimFunc(symshen_4in_1_6), W3472)


__e.TailApply(tmp10926, tmp10941)
return


}, 1)

tmp10942 := Call(__e, PrimFunc(symshen_4_5_1out), W3472)


__e.TailApply(tmp10925, tmp10942)
return


}


}, 1)

tmp10945 := Call(__e, PrimFunc(symshen_4_5ass_6), V3468)


tmp10946 := Call(__e, tmp10924, tmp10945)


__e.TailApply(tmp10907, tmp10946)
return


} else {
__e.Return(W3469)
return
}


}, 1)

tmp10954 := Call(__e, PrimFunc(symshen_4hds_a_2), V3468, sym_b)


var ifres10949 Obj

if True == tmp10954 {
tmp10950 := MakeNative(func(__e *ControlFlow) {
W3470 := __e.Get(1)
_ = W3470
__e.TailApply(PrimFunc(symshen_4comb), W3470, sym_b)
return
}, 1)

tmp10951 := Call(__e, PrimFunc(symtail), V3468)


tmp10952 := Call(__e, tmp10950, tmp10951)


ifres10949 = tmp10952


} else {
tmp10953 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres10949 = tmp10953


}

__e.TailApply(tmp10906, ifres10949)
return


}, 1)

tmp10955 := Call(__e, ns2_1set, symshen_4_5prem_6, tmp10905)


_ = tmp10955

tmp10956 := MakeNative(func(__e *ControlFlow) {
V3483 := __e.Get(1)
_ = V3483
tmp10957 := MakeNative(func(__e *ControlFlow) {
W3484 := __e.Get(1)
_ = W3484
tmp10982 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3484)


if True == tmp10982 {
tmp10958 := MakeNative(func(__e *ControlFlow) {
W3493 := __e.Get(1)
_ = W3493
tmp10970 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3493)


if True == tmp10970 {
tmp10959 := MakeNative(func(__e *ControlFlow) {
W3497 := __e.Get(1)
_ = W3497
tmp10961 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3497)


if True == tmp10961 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3497)
return
}


}, 1)

tmp10962 := MakeNative(func(__e *ControlFlow) {
W3498 := __e.Get(1)
_ = W3498
tmp10966 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3498)


if True == tmp10966 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10963 := MakeNative(func(__e *ControlFlow) {
W3499 := __e.Get(1)
_ = W3499
__e.TailApply(PrimFunc(symshen_4comb), W3499, Nil)
return
}, 1)

tmp10964 := Call(__e, PrimFunc(symshen_4in_1_6), W3498)


__e.TailApply(tmp10963, tmp10964)
return


}


}, 1)

tmp10967 := Call(__e, PrimFunc(sym_5e_6), V3483)


tmp10968 := Call(__e, tmp10962, tmp10967)


__e.TailApply(tmp10959, tmp10968)
return


} else {
__e.Return(W3493)
return
}


}, 1)

tmp10971 := MakeNative(func(__e *ControlFlow) {
W3494 := __e.Get(1)
_ = W3494
tmp10978 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3494)


if True == tmp10978 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10972 := MakeNative(func(__e *ControlFlow) {
W3495 := __e.Get(1)
_ = W3495
tmp10973 := MakeNative(func(__e *ControlFlow) {
W3496 := __e.Get(1)
_ = W3496
tmp10974 := PrimCons(W3495, Nil)

__e.TailApply(PrimFunc(symshen_4comb), W3496, tmp10974)
return


}, 1)

tmp10975 := Call(__e, PrimFunc(symshen_4in_1_6), W3494)


__e.TailApply(tmp10973, tmp10975)
return


}, 1)

tmp10976 := Call(__e, PrimFunc(symshen_4_5_1out), W3494)


__e.TailApply(tmp10972, tmp10976)
return


}


}, 1)

tmp10979 := Call(__e, PrimFunc(symshen_4_5formula_6), V3483)


tmp10980 := Call(__e, tmp10971, tmp10979)


__e.TailApply(tmp10958, tmp10980)
return


} else {
__e.Return(W3484)
return
}


}, 1)

tmp10983 := MakeNative(func(__e *ControlFlow) {
W3485 := __e.Get(1)
_ = W3485
tmp11004 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3485)


if True == tmp11004 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10984 := MakeNative(func(__e *ControlFlow) {
W3486 := __e.Get(1)
_ = W3486
tmp10985 := MakeNative(func(__e *ControlFlow) {
W3487 := __e.Get(1)
_ = W3487
tmp10986 := MakeNative(func(__e *ControlFlow) {
W3488 := __e.Get(1)
_ = W3488
tmp10999 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3488)


if True == tmp10999 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10987 := MakeNative(func(__e *ControlFlow) {
W3489 := __e.Get(1)
_ = W3489
tmp10988 := MakeNative(func(__e *ControlFlow) {
W3490 := __e.Get(1)
_ = W3490
tmp10995 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3490)


if True == tmp10995 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp10989 := MakeNative(func(__e *ControlFlow) {
W3491 := __e.Get(1)
_ = W3491
tmp10990 := MakeNative(func(__e *ControlFlow) {
W3492 := __e.Get(1)
_ = W3492
tmp10991 := PrimCons(W3486, W3491)

__e.TailApply(PrimFunc(symshen_4comb), W3492, tmp10991)
return


}, 1)

tmp10992 := Call(__e, PrimFunc(symshen_4in_1_6), W3490)


__e.TailApply(tmp10990, tmp10992)
return


}, 1)

tmp10993 := Call(__e, PrimFunc(symshen_4_5_1out), W3490)


__e.TailApply(tmp10989, tmp10993)
return


}


}, 1)

tmp10996 := Call(__e, PrimFunc(symshen_4_5ass_6), W3489)


__e.TailApply(tmp10988, tmp10996)
return


}, 1)

tmp10997 := Call(__e, PrimFunc(symshen_4in_1_6), W3488)


__e.TailApply(tmp10987, tmp10997)
return


}


}, 1)

tmp11000 := Call(__e, PrimFunc(symshen_4_5iscomma_6), W3487)


__e.TailApply(tmp10986, tmp11000)
return


}, 1)

tmp11001 := Call(__e, PrimFunc(symshen_4in_1_6), W3485)


__e.TailApply(tmp10985, tmp11001)
return


}, 1)

tmp11002 := Call(__e, PrimFunc(symshen_4_5_1out), W3485)


__e.TailApply(tmp10984, tmp11002)
return


}


}, 1)

tmp11005 := Call(__e, PrimFunc(symshen_4_5formula_6), V3483)


tmp11006 := Call(__e, tmp10983, tmp11005)


__e.TailApply(tmp10957, tmp11006)
return


}, 1)

tmp11007 := Call(__e, ns2_1set, symshen_4_5ass_6, tmp10956)


_ = tmp11007

tmp11008 := MakeNative(func(__e *ControlFlow) {
V3500 := __e.Get(1)
_ = V3500
tmp11009 := MakeNative(func(__e *ControlFlow) {
W3501 := __e.Get(1)
_ = W3501
tmp11011 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3501)


if True == tmp11011 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3501)
return
}


}, 1)

tmp11022 := PrimIsPair(V3500)

var ifres11012 Obj

if True == tmp11022 {
tmp11013 := MakeNative(func(__e *ControlFlow) {
W3502 := __e.Get(1)
_ = W3502
tmp11014 := MakeNative(func(__e *ControlFlow) {
W3503 := __e.Get(1)
_ = W3503
tmp11016 := PrimIntern(MakeString(","))

tmp11017 := PrimEqual(W3502, tmp11016)

if True == tmp11017 {
__e.TailApply(PrimFunc(symshen_4comb), W3503, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp11018 := Call(__e, PrimFunc(symtail), V3500)


__e.TailApply(tmp11014, tmp11018)
return


}, 1)

tmp11019 := Call(__e, PrimFunc(symhead), V3500)


tmp11020 := Call(__e, tmp11013, tmp11019)


ifres11012 = tmp11020


} else {
tmp11021 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11012 = tmp11021


}

__e.TailApply(tmp11009, ifres11012)
return


}, 1)

tmp11023 := Call(__e, ns2_1set, symshen_4_5iscomma_6, tmp11008)


_ = tmp11023

tmp11024 := MakeNative(func(__e *ControlFlow) {
V3504 := __e.Get(1)
_ = V3504
tmp11025 := MakeNative(func(__e *ControlFlow) {
W3505 := __e.Get(1)
_ = W3505
tmp11039 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3505)


if True == tmp11039 {
tmp11026 := MakeNative(func(__e *ControlFlow) {
W3514 := __e.Get(1)
_ = W3514
tmp11028 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3514)


if True == tmp11028 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3514)
return
}


}, 1)

tmp11029 := MakeNative(func(__e *ControlFlow) {
W3515 := __e.Get(1)
_ = W3515
tmp11035 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3515)


if True == tmp11035 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11030 := MakeNative(func(__e *ControlFlow) {
W3516 := __e.Get(1)
_ = W3516
tmp11031 := MakeNative(func(__e *ControlFlow) {
W3517 := __e.Get(1)
_ = W3517
__e.TailApply(PrimFunc(symshen_4comb), W3517, W3516)
return
}, 1)

tmp11032 := Call(__e, PrimFunc(symshen_4in_1_6), W3515)


__e.TailApply(tmp11031, tmp11032)
return


}, 1)

tmp11033 := Call(__e, PrimFunc(symshen_4_5_1out), W3515)


__e.TailApply(tmp11030, tmp11033)
return


}


}, 1)

tmp11036 := Call(__e, PrimFunc(symshen_4_5expr_6), V3504)


tmp11037 := Call(__e, tmp11029, tmp11036)


__e.TailApply(tmp11026, tmp11037)
return


} else {
__e.Return(W3505)
return
}


}, 1)

tmp11040 := MakeNative(func(__e *ControlFlow) {
W3506 := __e.Get(1)
_ = W3506
tmp11066 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3506)


if True == tmp11066 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11041 := MakeNative(func(__e *ControlFlow) {
W3507 := __e.Get(1)
_ = W3507
tmp11042 := MakeNative(func(__e *ControlFlow) {
W3508 := __e.Get(1)
_ = W3508
tmp11043 := MakeNative(func(__e *ControlFlow) {
W3509 := __e.Get(1)
_ = W3509
tmp11061 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3509)


if True == tmp11061 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11044 := MakeNative(func(__e *ControlFlow) {
W3510 := __e.Get(1)
_ = W3510
tmp11045 := MakeNative(func(__e *ControlFlow) {
W3511 := __e.Get(1)
_ = W3511
tmp11057 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3511)


if True == tmp11057 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11046 := MakeNative(func(__e *ControlFlow) {
W3512 := __e.Get(1)
_ = W3512
tmp11047 := MakeNative(func(__e *ControlFlow) {
W3513 := __e.Get(1)
_ = W3513
tmp11048 := Call(__e, PrimFunc(symshen_4curry), W3507)


tmp11049 := PrimIntern(MakeString(":"))

tmp11050 := Call(__e, PrimFunc(symshen_4rectify_1type), W3512)


tmp11051 := PrimCons(tmp11050, Nil)

tmp11052 := PrimCons(tmp11049, tmp11051)

tmp11053 := PrimCons(tmp11048, tmp11052)

__e.TailApply(PrimFunc(symshen_4comb), W3513, tmp11053)
return


}, 1)

tmp11054 := Call(__e, PrimFunc(symshen_4in_1_6), W3511)


__e.TailApply(tmp11047, tmp11054)
return


}, 1)

tmp11055 := Call(__e, PrimFunc(symshen_4_5_1out), W3511)


__e.TailApply(tmp11046, tmp11055)
return


}


}, 1)

tmp11058 := Call(__e, PrimFunc(symshen_4_5type_6), W3510)


__e.TailApply(tmp11045, tmp11058)
return


}, 1)

tmp11059 := Call(__e, PrimFunc(symshen_4in_1_6), W3509)


__e.TailApply(tmp11044, tmp11059)
return


}


}, 1)

tmp11062 := Call(__e, PrimFunc(symshen_4_5iscolon_6), W3508)


__e.TailApply(tmp11043, tmp11062)
return


}, 1)

tmp11063 := Call(__e, PrimFunc(symshen_4in_1_6), W3506)


__e.TailApply(tmp11042, tmp11063)
return


}, 1)

tmp11064 := Call(__e, PrimFunc(symshen_4_5_1out), W3506)


__e.TailApply(tmp11041, tmp11064)
return


}


}, 1)

tmp11067 := Call(__e, PrimFunc(symshen_4_5expr_6), V3504)


tmp11068 := Call(__e, tmp11040, tmp11067)


__e.TailApply(tmp11025, tmp11068)
return


}, 1)

tmp11069 := Call(__e, ns2_1set, symshen_4_5formula_6, tmp11024)


_ = tmp11069

tmp11070 := MakeNative(func(__e *ControlFlow) {
V3518 := __e.Get(1)
_ = V3518
tmp11071 := MakeNative(func(__e *ControlFlow) {
W3519 := __e.Get(1)
_ = W3519
tmp11073 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3519)


if True == tmp11073 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3519)
return
}


}, 1)

tmp11084 := PrimIsPair(V3518)

var ifres11074 Obj

if True == tmp11084 {
tmp11075 := MakeNative(func(__e *ControlFlow) {
W3520 := __e.Get(1)
_ = W3520
tmp11076 := MakeNative(func(__e *ControlFlow) {
W3521 := __e.Get(1)
_ = W3521
tmp11078 := PrimIntern(MakeString(":"))

tmp11079 := PrimEqual(W3520, tmp11078)

if True == tmp11079 {
__e.TailApply(PrimFunc(symshen_4comb), W3521, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp11080 := Call(__e, PrimFunc(symtail), V3518)


__e.TailApply(tmp11076, tmp11080)
return


}, 1)

tmp11081 := Call(__e, PrimFunc(symhead), V3518)


tmp11082 := Call(__e, tmp11075, tmp11081)


ifres11074 = tmp11082


} else {
tmp11083 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11074 = tmp11083


}

__e.TailApply(tmp11071, ifres11074)
return


}, 1)

tmp11085 := Call(__e, ns2_1set, symshen_4_5iscolon_6, tmp11070)


_ = tmp11085

tmp11086 := MakeNative(func(__e *ControlFlow) {
V3522 := __e.Get(1)
_ = V3522
tmp11087 := MakeNative(func(__e *ControlFlow) {
W3523 := __e.Get(1)
_ = W3523
tmp11099 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3523)


if True == tmp11099 {
tmp11088 := MakeNative(func(__e *ControlFlow) {
W3530 := __e.Get(1)
_ = W3530
tmp11090 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3530)


if True == tmp11090 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3530)
return
}


}, 1)

tmp11091 := MakeNative(func(__e *ControlFlow) {
W3531 := __e.Get(1)
_ = W3531
tmp11095 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3531)


if True == tmp11095 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11092 := MakeNative(func(__e *ControlFlow) {
W3532 := __e.Get(1)
_ = W3532
__e.TailApply(PrimFunc(symshen_4comb), W3532, Nil)
return
}, 1)

tmp11093 := Call(__e, PrimFunc(symshen_4in_1_6), W3531)


__e.TailApply(tmp11092, tmp11093)
return


}


}, 1)

tmp11096 := Call(__e, PrimFunc(sym_5e_6), V3522)


tmp11097 := Call(__e, tmp11091, tmp11096)


__e.TailApply(tmp11088, tmp11097)
return


} else {
__e.Return(W3523)
return
}


}, 1)

tmp11100 := MakeNative(func(__e *ControlFlow) {
W3524 := __e.Get(1)
_ = W3524
tmp11115 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3524)


if True == tmp11115 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11101 := MakeNative(func(__e *ControlFlow) {
W3525 := __e.Get(1)
_ = W3525
tmp11102 := MakeNative(func(__e *ControlFlow) {
W3526 := __e.Get(1)
_ = W3526
tmp11103 := MakeNative(func(__e *ControlFlow) {
W3527 := __e.Get(1)
_ = W3527
tmp11110 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3527)


if True == tmp11110 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11104 := MakeNative(func(__e *ControlFlow) {
W3528 := __e.Get(1)
_ = W3528
tmp11105 := MakeNative(func(__e *ControlFlow) {
W3529 := __e.Get(1)
_ = W3529
tmp11106 := PrimCons(W3525, W3528)

__e.TailApply(PrimFunc(symshen_4comb), W3529, tmp11106)
return


}, 1)

tmp11107 := Call(__e, PrimFunc(symshen_4in_1_6), W3527)


__e.TailApply(tmp11105, tmp11107)
return


}, 1)

tmp11108 := Call(__e, PrimFunc(symshen_4_5_1out), W3527)


__e.TailApply(tmp11104, tmp11108)
return


}


}, 1)

tmp11111 := Call(__e, PrimFunc(symshen_4_5sides_6), W3526)


__e.TailApply(tmp11103, tmp11111)
return


}, 1)

tmp11112 := Call(__e, PrimFunc(symshen_4in_1_6), W3524)


__e.TailApply(tmp11102, tmp11112)
return


}, 1)

tmp11113 := Call(__e, PrimFunc(symshen_4_5_1out), W3524)


__e.TailApply(tmp11101, tmp11113)
return


}


}, 1)

tmp11116 := Call(__e, PrimFunc(symshen_4_5side_6), V3522)


tmp11117 := Call(__e, tmp11100, tmp11116)


__e.TailApply(tmp11087, tmp11117)
return


}, 1)

tmp11118 := Call(__e, ns2_1set, symshen_4_5sides_6, tmp11086)


_ = tmp11118

tmp11119 := MakeNative(func(__e *ControlFlow) {
V3533 := __e.Get(1)
_ = V3533
tmp11120 := MakeNative(func(__e *ControlFlow) {
W3534 := __e.Get(1)
_ = W3534
tmp11165 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3534)


if True == tmp11165 {
tmp11121 := MakeNative(func(__e *ControlFlow) {
W3538 := __e.Get(1)
_ = W3538
tmp11142 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3538)


if True == tmp11142 {
tmp11122 := MakeNative(func(__e *ControlFlow) {
W3544 := __e.Get(1)
_ = W3544
tmp11124 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3544)


if True == tmp11124 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3544)
return
}


}, 1)

tmp11140 := Call(__e, PrimFunc(symshen_4hds_a_2), V3533, symctxt)


var ifres11125 Obj

if True == tmp11140 {
tmp11126 := MakeNative(func(__e *ControlFlow) {
W3545 := __e.Get(1)
_ = W3545
tmp11136 := PrimIsPair(W3545)

if True == tmp11136 {
tmp11127 := MakeNative(func(__e *ControlFlow) {
W3546 := __e.Get(1)
_ = W3546
tmp11128 := MakeNative(func(__e *ControlFlow) {
W3547 := __e.Get(1)
_ = W3547
tmp11132 := PrimIsVariable(W3546)

if True == tmp11132 {
tmp11129 := PrimCons(W3546, Nil)

tmp11130 := PrimCons(symctxt, tmp11129)

__e.TailApply(PrimFunc(symshen_4comb), W3547, tmp11130)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp11133 := Call(__e, PrimFunc(symtail), W3545)


__e.TailApply(tmp11128, tmp11133)
return


}, 1)

tmp11134 := Call(__e, PrimFunc(symhead), W3545)


__e.TailApply(tmp11127, tmp11134)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp11137 := Call(__e, PrimFunc(symtail), V3533)


tmp11138 := Call(__e, tmp11126, tmp11137)


ifres11125 = tmp11138


} else {
tmp11139 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11125 = tmp11139


}

__e.TailApply(tmp11122, ifres11125)
return


} else {
__e.Return(W3538)
return
}


}, 1)

tmp11163 := Call(__e, PrimFunc(symshen_4hds_a_2), V3533, symlet)


var ifres11143 Obj

if True == tmp11163 {
tmp11144 := MakeNative(func(__e *ControlFlow) {
W3539 := __e.Get(1)
_ = W3539
tmp11159 := PrimIsPair(W3539)

if True == tmp11159 {
tmp11145 := MakeNative(func(__e *ControlFlow) {
W3540 := __e.Get(1)
_ = W3540
tmp11146 := MakeNative(func(__e *ControlFlow) {
W3541 := __e.Get(1)
_ = W3541
tmp11155 := PrimIsPair(W3541)

if True == tmp11155 {
tmp11147 := MakeNative(func(__e *ControlFlow) {
W3542 := __e.Get(1)
_ = W3542
tmp11148 := MakeNative(func(__e *ControlFlow) {
W3543 := __e.Get(1)
_ = W3543
tmp11149 := PrimCons(W3542, Nil)

tmp11150 := PrimCons(W3540, tmp11149)

tmp11151 := PrimCons(symlet, tmp11150)

__e.TailApply(PrimFunc(symshen_4comb), W3543, tmp11151)
return


}, 1)

tmp11152 := Call(__e, PrimFunc(symtail), W3541)


__e.TailApply(tmp11148, tmp11152)
return


}, 1)

tmp11153 := Call(__e, PrimFunc(symhead), W3541)


__e.TailApply(tmp11147, tmp11153)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp11156 := Call(__e, PrimFunc(symtail), W3539)


__e.TailApply(tmp11146, tmp11156)
return


}, 1)

tmp11157 := Call(__e, PrimFunc(symhead), W3539)


__e.TailApply(tmp11145, tmp11157)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp11160 := Call(__e, PrimFunc(symtail), V3533)


tmp11161 := Call(__e, tmp11144, tmp11160)


ifres11143 = tmp11161


} else {
tmp11162 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11143 = tmp11162


}

__e.TailApply(tmp11121, ifres11143)
return


} else {
__e.Return(W3534)
return
}


}, 1)

tmp11179 := Call(__e, PrimFunc(symshen_4hds_a_2), V3533, symif)


var ifres11166 Obj

if True == tmp11179 {
tmp11167 := MakeNative(func(__e *ControlFlow) {
W3535 := __e.Get(1)
_ = W3535
tmp11175 := PrimIsPair(W3535)

if True == tmp11175 {
tmp11168 := MakeNative(func(__e *ControlFlow) {
W3536 := __e.Get(1)
_ = W3536
tmp11169 := MakeNative(func(__e *ControlFlow) {
W3537 := __e.Get(1)
_ = W3537
tmp11170 := PrimCons(W3536, Nil)

tmp11171 := PrimCons(symif, tmp11170)

__e.TailApply(PrimFunc(symshen_4comb), W3537, tmp11171)
return


}, 1)

tmp11172 := Call(__e, PrimFunc(symtail), W3535)


__e.TailApply(tmp11169, tmp11172)
return


}, 1)

tmp11173 := Call(__e, PrimFunc(symhead), W3535)


__e.TailApply(tmp11168, tmp11173)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp11176 := Call(__e, PrimFunc(symtail), V3533)


tmp11177 := Call(__e, tmp11167, tmp11176)


ifres11166 = tmp11177


} else {
tmp11178 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11166 = tmp11178


}

__e.TailApply(tmp11120, ifres11166)
return


}, 1)

tmp11180 := Call(__e, ns2_1set, symshen_4_5side_6, tmp11119)


_ = tmp11180

tmp11181 := MakeNative(func(__e *ControlFlow) {
V3554 := __e.Get(1)
_ = V3554
V3555 := __e.Get(2)
_ = V3555
V3556 := __e.Get(3)
_ = V3556
tmp11216 := PrimIsPair(V3556)

var ifres11203 Obj

if True == tmp11216 {
tmp11214 := PrimHead(V3556)

tmp11215 := PrimEqual(Nil, tmp11214)

var ifres11205 Obj

if True == tmp11215 {
tmp11212 := PrimTail(V3556)

tmp11213 := PrimIsPair(tmp11212)

var ifres11207 Obj

if True == tmp11213 {
tmp11209 := PrimTail(V3556)

tmp11210 := PrimTail(tmp11209)

tmp11211 := PrimEqual(Nil, tmp11210)

var ifres11208 Obj

if True == tmp11211 {
ifres11208 = True


} else {
ifres11208 = False


}

ifres11207 = ifres11208


} else {
ifres11207 = False


}

var ifres11206 Obj

if True == ifres11207 {
ifres11206 = True


} else {
ifres11206 = False


}

ifres11205 = ifres11206


} else {
ifres11205 = False


}

var ifres11204 Obj

if True == ifres11205 {
ifres11204 = True


} else {
ifres11204 = False


}

ifres11203 = ifres11204


} else {
ifres11203 = False


}

if True == ifres11203 {
tmp11182 := MakeNative(func(__e *ControlFlow) {
W3557 := __e.Get(1)
_ = W3557
tmp11183 := MakeNative(func(__e *ControlFlow) {
W3558 := __e.Get(1)
_ = W3558
tmp11184 := MakeNative(func(__e *ControlFlow) {
W3559 := __e.Get(1)
_ = W3559
tmp11185 := MakeNative(func(__e *ControlFlow) {
W3560 := __e.Get(1)
_ = W3560
tmp11186 := MakeNative(func(__e *ControlFlow) {
W3561 := __e.Get(1)
_ = W3561
tmp11187 := PrimCons(W3560, Nil)

__e.Return(PrimCons(W3561, tmp11187))
return


}, 1)

tmp11188 := PrimCons(V3556, Nil)

tmp11189 := PrimCons(V3555, tmp11188)

tmp11190 := PrimCons(V3554, tmp11189)

__e.TailApply(tmp11186, tmp11190)
return


}, 1)

tmp11191 := PrimCons(W3559, Nil)

tmp11192 := PrimCons(W3558, Nil)

tmp11193 := PrimCons(tmp11191, tmp11192)

tmp11194 := PrimCons(V3554, tmp11193)

__e.TailApply(tmp11185, tmp11194)
return


}, 1)

tmp11195 := Call(__e, PrimFunc(symshen_4coll_1formulae), V3555)


tmp11196 := PrimCons(W3557, Nil)

tmp11197 := PrimCons(tmp11195, tmp11196)

__e.TailApply(tmp11184, tmp11197)
return


}, 1)

tmp11198 := PrimTail(V3556)

tmp11199 := PrimCons(W3557, Nil)

tmp11200 := PrimCons(tmp11198, tmp11199)

__e.TailApply(tmp11183, tmp11200)
return


}, 1)

tmp11201 := Call(__e, PrimFunc(symgensym), symP)


__e.TailApply(tmp11182, tmp11201)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.lr-rule")))
return
}


}, 3)

tmp11217 := Call(__e, ns2_1set, symshen_4lr_1rule, tmp11181)


_ = tmp11217

tmp11218 := MakeNative(func(__e *ControlFlow) {
V3564 := __e.Get(1)
_ = V3564
tmp11247 := PrimEqual(Nil, V3564)

if True == tmp11247 {
__e.Return(Nil)
return
} else {
tmp11245 := PrimIsPair(V3564)

var ifres11225 Obj

if True == tmp11245 {
tmp11243 := PrimHead(V3564)

tmp11244 := PrimIsPair(tmp11243)

var ifres11227 Obj

if True == tmp11244 {
tmp11240 := PrimHead(V3564)

tmp11241 := PrimHead(tmp11240)

tmp11242 := PrimEqual(Nil, tmp11241)

var ifres11229 Obj

if True == tmp11242 {
tmp11237 := PrimHead(V3564)

tmp11238 := PrimTail(tmp11237)

tmp11239 := PrimIsPair(tmp11238)

var ifres11231 Obj

if True == tmp11239 {
tmp11233 := PrimHead(V3564)

tmp11234 := PrimTail(tmp11233)

tmp11235 := PrimTail(tmp11234)

tmp11236 := PrimEqual(Nil, tmp11235)

var ifres11232 Obj

if True == tmp11236 {
ifres11232 = True


} else {
ifres11232 = False


}

ifres11231 = ifres11232


} else {
ifres11231 = False


}

var ifres11230 Obj

if True == ifres11231 {
ifres11230 = True


} else {
ifres11230 = False


}

ifres11229 = ifres11230


} else {
ifres11229 = False


}

var ifres11228 Obj

if True == ifres11229 {
ifres11228 = True


} else {
ifres11228 = False


}

ifres11227 = ifres11228


} else {
ifres11227 = False


}

var ifres11226 Obj

if True == ifres11227 {
ifres11226 = True


} else {
ifres11226 = False


}

ifres11225 = ifres11226


} else {
ifres11225 = False


}

if True == ifres11225 {
tmp11219 := PrimHead(V3564)

tmp11220 := PrimTail(tmp11219)

tmp11221 := PrimHead(tmp11220)

tmp11222 := PrimTail(V3564)

tmp11223 := Call(__e, PrimFunc(symshen_4coll_1formulae), tmp11222)


__e.Return(PrimCons(tmp11221, tmp11223))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.coll-formulae")))
return
}


}


}, 1)

tmp11248 := Call(__e, ns2_1set, symshen_4coll_1formulae, tmp11218)


_ = tmp11248

tmp11249 := MakeNative(func(__e *ControlFlow) {
V3565 := __e.Get(1)
_ = V3565
tmp11250 := MakeNative(func(__e *ControlFlow) {
W3566 := __e.Get(1)
_ = W3566
tmp11252 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3566)


if True == tmp11252 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3566)
return
}


}, 1)

tmp11264 := PrimIsPair(V3565)

var ifres11253 Obj

if True == tmp11264 {
tmp11254 := MakeNative(func(__e *ControlFlow) {
W3567 := __e.Get(1)
_ = W3567
tmp11255 := MakeNative(func(__e *ControlFlow) {
W3568 := __e.Get(1)
_ = W3568
tmp11258 := Call(__e, PrimFunc(symshen_4key_1in_1sequent_1calculus_2), W3567)


tmp11259 := PrimNot(tmp11258)

if True == tmp11259 {
tmp11256 := Call(__e, PrimFunc(symmacroexpand), W3567)


__e.TailApply(PrimFunc(symshen_4comb), W3568, tmp11256)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp11260 := Call(__e, PrimFunc(symtail), V3565)


__e.TailApply(tmp11255, tmp11260)
return


}, 1)

tmp11261 := Call(__e, PrimFunc(symhead), V3565)


tmp11262 := Call(__e, tmp11254, tmp11261)


ifres11253 = tmp11262


} else {
tmp11263 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11253 = tmp11263


}

__e.TailApply(tmp11250, ifres11253)
return


}, 1)

tmp11265 := Call(__e, ns2_1set, symshen_4_5expr_6, tmp11249)


_ = tmp11265

tmp11266 := MakeNative(func(__e *ControlFlow) {
V3569 := __e.Get(1)
_ = V3569
tmp11273 := PrimIntern(MakeString(";"))

tmp11274 := PrimIntern(MakeString(","))

tmp11275 := PrimIntern(MakeString(":"))

tmp11276 := PrimCons(sym_5_1_1, Nil)

tmp11277 := PrimCons(tmp11275, tmp11276)

tmp11278 := PrimCons(tmp11274, tmp11277)

tmp11279 := PrimCons(tmp11273, tmp11278)

tmp11280 := PrimCons(sym_6_6, tmp11279)

tmp11281 := Call(__e, PrimFunc(symelement_2), V3569, tmp11280)


if True == tmp11281 {
__e.Return(True)
return
} else {
tmp11271 := Call(__e, PrimFunc(symshen_4sng_2), V3569)


var ifres11268 Obj

if True == tmp11271 {
ifres11268 = True


} else {
tmp11270 := Call(__e, PrimFunc(symshen_4dbl_2), V3569)


var ifres11269 Obj

if True == tmp11270 {
ifres11269 = True


} else {
ifres11269 = False


}

ifres11268 = ifres11269


}

if True == ifres11268 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp11282 := Call(__e, ns2_1set, symshen_4key_1in_1sequent_1calculus_2, tmp11266)


_ = tmp11282

tmp11283 := MakeNative(func(__e *ControlFlow) {
V3570 := __e.Get(1)
_ = V3570
tmp11284 := MakeNative(func(__e *ControlFlow) {
W3571 := __e.Get(1)
_ = W3571
tmp11286 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3571)


if True == tmp11286 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3571)
return
}


}, 1)

tmp11287 := MakeNative(func(__e *ControlFlow) {
W3572 := __e.Get(1)
_ = W3572
tmp11293 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3572)


if True == tmp11293 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11288 := MakeNative(func(__e *ControlFlow) {
W3573 := __e.Get(1)
_ = W3573
tmp11289 := MakeNative(func(__e *ControlFlow) {
W3574 := __e.Get(1)
_ = W3574
__e.TailApply(PrimFunc(symshen_4comb), W3574, W3573)
return
}, 1)

tmp11290 := Call(__e, PrimFunc(symshen_4in_1_6), W3572)


__e.TailApply(tmp11289, tmp11290)
return


}, 1)

tmp11291 := Call(__e, PrimFunc(symshen_4_5_1out), W3572)


__e.TailApply(tmp11288, tmp11291)
return


}


}, 1)

tmp11294 := Call(__e, PrimFunc(symshen_4_5expr_6), V3570)


tmp11295 := Call(__e, tmp11287, tmp11294)


__e.TailApply(tmp11284, tmp11295)
return


}, 1)

tmp11296 := Call(__e, ns2_1set, symshen_4_5type_6, tmp11283)


_ = tmp11296

tmp11297 := MakeNative(func(__e *ControlFlow) {
V3575 := __e.Get(1)
_ = V3575
tmp11298 := MakeNative(func(__e *ControlFlow) {
W3576 := __e.Get(1)
_ = W3576
tmp11300 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3576)


if True == tmp11300 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3576)
return
}


}, 1)

tmp11310 := PrimIsPair(V3575)

var ifres11301 Obj

if True == tmp11310 {
tmp11302 := MakeNative(func(__e *ControlFlow) {
W3577 := __e.Get(1)
_ = W3577
tmp11303 := MakeNative(func(__e *ControlFlow) {
W3578 := __e.Get(1)
_ = W3578
tmp11305 := Call(__e, PrimFunc(symshen_4dbl_2), W3577)


if True == tmp11305 {
__e.TailApply(PrimFunc(symshen_4comb), W3578, W3577)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp11306 := Call(__e, PrimFunc(symtail), V3575)


__e.TailApply(tmp11303, tmp11306)
return


}, 1)

tmp11307 := Call(__e, PrimFunc(symhead), V3575)


tmp11308 := Call(__e, tmp11302, tmp11307)


ifres11301 = tmp11308


} else {
tmp11309 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11301 = tmp11309


}

__e.TailApply(tmp11298, ifres11301)
return


}, 1)

tmp11311 := Call(__e, ns2_1set, symshen_4_5dbl_6, tmp11297)


_ = tmp11311

tmp11312 := MakeNative(func(__e *ControlFlow) {
V3579 := __e.Get(1)
_ = V3579
tmp11313 := MakeNative(func(__e *ControlFlow) {
W3580 := __e.Get(1)
_ = W3580
tmp11315 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3580)


if True == tmp11315 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3580)
return
}


}, 1)

tmp11325 := PrimIsPair(V3579)

var ifres11316 Obj

if True == tmp11325 {
tmp11317 := MakeNative(func(__e *ControlFlow) {
W3581 := __e.Get(1)
_ = W3581
tmp11318 := MakeNative(func(__e *ControlFlow) {
W3582 := __e.Get(1)
_ = W3582
tmp11320 := Call(__e, PrimFunc(symshen_4sng_2), W3581)


if True == tmp11320 {
__e.TailApply(PrimFunc(symshen_4comb), W3582, W3581)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp11321 := Call(__e, PrimFunc(symtail), V3579)


__e.TailApply(tmp11318, tmp11321)
return


}, 1)

tmp11322 := Call(__e, PrimFunc(symhead), V3579)


tmp11323 := Call(__e, tmp11317, tmp11322)


ifres11316 = tmp11323


} else {
tmp11324 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11316 = tmp11324


}

__e.TailApply(tmp11313, ifres11316)
return


}, 1)

tmp11326 := Call(__e, ns2_1set, symshen_4_5sng_6, tmp11312)


_ = tmp11326

tmp11327 := MakeNative(func(__e *ControlFlow) {
V3583 := __e.Get(1)
_ = V3583
tmp11332 := PrimIsSymbol(V3583)

if True == tmp11332 {
tmp11329 := PrimStr(V3583)

tmp11330 := Call(__e, PrimFunc(symshen_4sng_1h_2), tmp11329)


if True == tmp11330 {
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

tmp11333 := Call(__e, ns2_1set, symshen_4sng_2, tmp11327)


_ = tmp11333

tmp11334 := MakeNative(func(__e *ControlFlow) {
V3586 := __e.Get(1)
_ = V3586
tmp11343 := PrimEqual(MakeString("___"), V3586)

if True == tmp11343 {
__e.Return(True)
return
} else {
tmp11341 := Call(__e, PrimFunc(symshen_4_7string_2), V3586)


var ifres11337 Obj

if True == tmp11341 {
tmp11339 := Call(__e, PrimFunc(symhdstr), V3586)


tmp11340 := PrimEqual(MakeString("_"), tmp11339)

var ifres11338 Obj

if True == tmp11340 {
ifres11338 = True


} else {
ifres11338 = False


}

ifres11337 = ifres11338


} else {
ifres11337 = False


}

if True == ifres11337 {
tmp11335 := PrimTailString(V3586)

__e.TailApply(PrimFunc(symshen_4sng_1h_2), tmp11335)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp11344 := Call(__e, ns2_1set, symshen_4sng_1h_2, tmp11334)


_ = tmp11344

tmp11345 := MakeNative(func(__e *ControlFlow) {
V3587 := __e.Get(1)
_ = V3587
tmp11350 := PrimIsSymbol(V3587)

if True == tmp11350 {
tmp11347 := PrimStr(V3587)

tmp11348 := Call(__e, PrimFunc(symshen_4dbl_1h_2), tmp11347)


if True == tmp11348 {
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

tmp11351 := Call(__e, ns2_1set, symshen_4dbl_2, tmp11345)


_ = tmp11351

tmp11352 := MakeNative(func(__e *ControlFlow) {
V3590 := __e.Get(1)
_ = V3590
tmp11361 := PrimEqual(MakeString("==="), V3590)

if True == tmp11361 {
__e.Return(True)
return
} else {
tmp11359 := Call(__e, PrimFunc(symshen_4_7string_2), V3590)


var ifres11355 Obj

if True == tmp11359 {
tmp11357 := Call(__e, PrimFunc(symhdstr), V3590)


tmp11358 := PrimEqual(MakeString("="), tmp11357)

var ifres11356 Obj

if True == tmp11358 {
ifres11356 = True


} else {
ifres11356 = False


}

ifres11355 = ifres11356


} else {
ifres11355 = False


}

if True == ifres11355 {
tmp11353 := PrimTailString(V3590)

__e.TailApply(PrimFunc(symshen_4dbl_1h_2), tmp11353)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp11362 := Call(__e, ns2_1set, symshen_4dbl_1h_2, tmp11352)


_ = tmp11362

tmp11363 := MakeNative(func(__e *ControlFlow) {
V3591 := __e.Get(1)
_ = V3591
V3592 := __e.Get(2)
_ = V3592
tmp11364 := MakeNative(func(__e *ControlFlow) {
W3593 := __e.Get(1)
_ = W3593
tmp11365 := MakeNative(func(__e *ControlFlow) {
W3595 := __e.Get(1)
_ = W3595
__e.TailApply(PrimFunc(symeval), W3595)
return
}, 1)

tmp11366 := PrimCons(V3591, W3593)

tmp11367 := PrimCons(symdefprolog, tmp11366)

__e.TailApply(tmp11365, tmp11367)
return


}, 1)

tmp11368 := MakeNative(func(__e *ControlFlow) {
Z3594 := __e.Get(1)
_ = Z3594
__e.TailApply(PrimFunc(symshen_4rule_1_6clause), Z3594)
return
}, 1)

tmp11369 := Call(__e, PrimFunc(symmapcan), tmp11368, V3592)


__e.TailApply(tmp11364, tmp11369)
return


}, 2)

tmp11370 := Call(__e, ns2_1set, symshen_4rules_1_6prolog, tmp11363)


_ = tmp11370

tmp11371 := MakeNative(func(__e *ControlFlow) {
V3596 := __e.Get(1)
_ = V3596
tmp11432 := PrimIsPair(V3596)

var ifres11396 Obj

if True == tmp11432 {
tmp11430 := PrimTail(V3596)

tmp11431 := PrimIsPair(tmp11430)

var ifres11398 Obj

if True == tmp11431 {
tmp11427 := PrimTail(V3596)

tmp11428 := PrimTail(tmp11427)

tmp11429 := PrimIsPair(tmp11428)

var ifres11400 Obj

if True == tmp11429 {
tmp11423 := PrimTail(V3596)

tmp11424 := PrimTail(tmp11423)

tmp11425 := PrimHead(tmp11424)

tmp11426 := PrimIsPair(tmp11425)

var ifres11402 Obj

if True == tmp11426 {
tmp11418 := PrimTail(V3596)

tmp11419 := PrimTail(tmp11418)

tmp11420 := PrimHead(tmp11419)

tmp11421 := PrimTail(tmp11420)

tmp11422 := PrimIsPair(tmp11421)

var ifres11404 Obj

if True == tmp11422 {
tmp11412 := PrimTail(V3596)

tmp11413 := PrimTail(tmp11412)

tmp11414 := PrimHead(tmp11413)

tmp11415 := PrimTail(tmp11414)

tmp11416 := PrimTail(tmp11415)

tmp11417 := PrimEqual(Nil, tmp11416)

var ifres11406 Obj

if True == tmp11417 {
tmp11408 := PrimTail(V3596)

tmp11409 := PrimTail(tmp11408)

tmp11410 := PrimTail(tmp11409)

tmp11411 := PrimEqual(Nil, tmp11410)

var ifres11407 Obj

if True == tmp11411 {
ifres11407 = True


} else {
ifres11407 = False


}

ifres11406 = ifres11407


} else {
ifres11406 = False


}

var ifres11405 Obj

if True == ifres11406 {
ifres11405 = True


} else {
ifres11405 = False


}

ifres11404 = ifres11405


} else {
ifres11404 = False


}

var ifres11403 Obj

if True == ifres11404 {
ifres11403 = True


} else {
ifres11403 = False


}

ifres11402 = ifres11403


} else {
ifres11402 = False


}

var ifres11401 Obj

if True == ifres11402 {
ifres11401 = True


} else {
ifres11401 = False


}

ifres11400 = ifres11401


} else {
ifres11400 = False


}

var ifres11399 Obj

if True == ifres11400 {
ifres11399 = True


} else {
ifres11399 = False


}

ifres11398 = ifres11399


} else {
ifres11398 = False


}

var ifres11397 Obj

if True == ifres11398 {
ifres11397 = True


} else {
ifres11397 = False


}

ifres11396 = ifres11397


} else {
ifres11396 = False


}

if True == ifres11396 {
tmp11372 := MakeNative(func(__e *ControlFlow) {
W3597 := __e.Get(1)
_ = W3597
tmp11373 := PrimTail(V3596)

tmp11374 := PrimTail(tmp11373)

tmp11375 := PrimHead(tmp11374)

tmp11376 := PrimTail(tmp11375)

tmp11377 := PrimHead(tmp11376)

tmp11378 := Call(__e, PrimFunc(symshen_4rule_1_6head), tmp11377)


tmp11379 := PrimCons(sym_5_1_1, Nil)

tmp11380 := PrimHead(V3596)

tmp11381 := PrimTail(V3596)

tmp11382 := PrimHead(tmp11381)

tmp11383 := PrimTail(V3596)

tmp11384 := PrimTail(tmp11383)

tmp11385 := PrimHead(tmp11384)

tmp11386 := PrimHead(tmp11385)

tmp11387 := Call(__e, PrimFunc(symshen_4rule_1_6body), W3597, symAssumptions, tmp11380, tmp11382, tmp11386)


tmp11388 := Call(__e, PrimFunc(symappend), tmp11379, tmp11387)


__e.TailApply(PrimFunc(symappend), tmp11378, tmp11388)
return


}, 1)

tmp11389 := PrimTail(V3596)

tmp11390 := PrimTail(tmp11389)

tmp11391 := PrimHead(tmp11390)

tmp11392 := PrimTail(tmp11391)

tmp11393 := PrimHead(tmp11392)

tmp11394 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp11393)


__e.TailApply(tmp11372, tmp11394)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4rule_1_6clause)
return
}


}, 1)

tmp11433 := Call(__e, ns2_1set, symshen_4rule_1_6clause, tmp11371)


_ = tmp11433

tmp11434 := MakeNative(func(__e *ControlFlow) {
V3598 := __e.Get(1)
_ = V3598
tmp11435 := Call(__e, PrimFunc(symshen_4macro_1_8ch), V3598)


tmp11436 := PrimCons(symAssumptions, Nil)

__e.Return(PrimCons(tmp11435, tmp11436))
return


}, 1)

tmp11437 := Call(__e, ns2_1set, symshen_4rule_1_6head, tmp11434)


_ = tmp11437

tmp11438 := MakeNative(func(__e *ControlFlow) {
V3599 := __e.Get(1)
_ = V3599
tmp11439 := PrimCons(V3599, Nil)

__e.Return(PrimCons(symshen_4_8ch, tmp11439))
return


}, 1)

tmp11440 := Call(__e, ns2_1set, symshen_4macro_1_8ch, tmp11438)


_ = tmp11440

tmp11441 := MakeNative(func(__e *ControlFlow) {
V3600 := __e.Get(1)
_ = V3600
tmp11442 := PrimCons(V3600, Nil)

__e.Return(PrimCons(symshen_4_8c, tmp11442))
return


}, 1)

tmp11443 := Call(__e, ns2_1set, symshen_4macro_1_8c, tmp11441)


_ = tmp11443

tmp11444 := MakeNative(func(__e *ControlFlow) {
V3601 := __e.Get(1)
_ = V3601
V3602 := __e.Get(2)
_ = V3602
V3603 := __e.Get(3)
_ = V3603
V3604 := __e.Get(4)
_ = V3604
V3605 := __e.Get(5)
_ = V3605
tmp11479 := PrimEqual(Nil, V3605)

if True == tmp11479 {
__e.TailApply(PrimFunc(symshen_4side_1conditions_1_6goals), Nil, V3601, V3602, V3603, V3604)
return
} else {
tmp11477 := PrimEqual(Nil, V3604)

var ifres11470 Obj

if True == tmp11477 {
tmp11476 := PrimIsPair(V3605)

var ifres11472 Obj

if True == tmp11476 {
tmp11474 := PrimTail(V3605)

tmp11475 := PrimEqual(Nil, tmp11474)

var ifres11473 Obj

if True == tmp11475 {
ifres11473 = True


} else {
ifres11473 = False


}

ifres11472 = ifres11473


} else {
ifres11472 = False


}

var ifres11471 Obj

if True == ifres11472 {
ifres11471 = True


} else {
ifres11471 = False


}

ifres11470 = ifres11471


} else {
ifres11470 = False


}

if True == ifres11470 {
tmp11445 := MakeNative(func(__e *ControlFlow) {
W3606 := __e.Get(1)
_ = W3606
tmp11446 := MakeNative(func(__e *ControlFlow) {
W3607 := __e.Get(1)
_ = W3607
tmp11447 := PrimHead(V3605)

tmp11448 := Call(__e, PrimFunc(symshen_4specialise_1member), tmp11447, V3602, W3607, W3606)


tmp11449 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), Nil, V3601, V3602, V3603, Nil)


__e.Return(PrimCons(tmp11448, tmp11449))
return


}, 1)

tmp11450 := PrimHead(V3605)

tmp11451 := Call(__e, PrimFunc(symshen_4remove_1bystanders), V3601, tmp11450)


__e.TailApply(tmp11446, tmp11451)
return


}, 1)

tmp11452 := PrimHead(V3605)

tmp11453 := Call(__e, PrimFunc(symshen_4passive_1variables), tmp11452, V3601)


__e.TailApply(tmp11445, tmp11453)
return


} else {
tmp11468 := PrimIsPair(V3605)

if True == tmp11468 {
tmp11454 := MakeNative(func(__e *ControlFlow) {
W3608 := __e.Get(1)
_ = W3608
tmp11455 := MakeNative(func(__e *ControlFlow) {
W3609 := __e.Get(1)
_ = W3609
tmp11456 := MakeNative(func(__e *ControlFlow) {
W3610 := __e.Get(1)
_ = W3610
tmp11457 := PrimHead(V3605)

tmp11458 := Call(__e, PrimFunc(symshen_4specialise_1consume), tmp11457, V3602, W3610, W3609, W3608)


tmp11459 := Call(__e, PrimFunc(symappend), V3601, W3609)


tmp11460 := PrimTail(V3605)

tmp11461 := Call(__e, PrimFunc(symshen_4rule_1_6body), tmp11459, W3608, V3603, V3604, tmp11460)


__e.Return(PrimCons(tmp11458, tmp11461))
return


}, 1)

tmp11462 := PrimHead(V3605)

tmp11463 := Call(__e, PrimFunc(symshen_4remove_1bystanders), V3601, tmp11462)


__e.TailApply(tmp11456, tmp11463)
return


}, 1)

tmp11464 := PrimHead(V3605)

tmp11465 := Call(__e, PrimFunc(symshen_4passive_1variables), tmp11464, V3601)


__e.TailApply(tmp11455, tmp11465)
return


}, 1)

tmp11466 := Call(__e, PrimFunc(symgensym), symNewAssumptions)


__e.TailApply(tmp11454, tmp11466)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4rule_1_6body)
return
}


}


}


}, 5)

tmp11480 := Call(__e, ns2_1set, symshen_4rule_1_6body, tmp11444)


_ = tmp11480

tmp11481 := MakeNative(func(__e *ControlFlow) {
V3611 := __e.Get(1)
_ = V3611
V3612 := __e.Get(2)
_ = V3612
V3613 := __e.Get(3)
_ = V3613
V3614 := __e.Get(4)
_ = V3614
tmp11482 := MakeNative(func(__e *ControlFlow) {
W3615 := __e.Get(1)
_ = W3615
tmp11483 := MakeNative(func(__e *ControlFlow) {
W3616 := __e.Get(1)
_ = W3616
tmp11484 := Call(__e, PrimFunc(symappend), V3613, V3614)


tmp11485 := PrimCons(V3612, tmp11484)

__e.Return(PrimCons(W3615, tmp11485))
return


}, 1)

tmp11486 := Call(__e, PrimFunc(symshen_4member_1clause), W3615, V3611, V3613, V3614)


__e.TailApply(tmp11483, tmp11486)
return


}, 1)

tmp11487 := Call(__e, PrimFunc(symgensym), symshen_4member)


__e.TailApply(tmp11482, tmp11487)
return


}, 4)

tmp11488 := Call(__e, ns2_1set, symshen_4specialise_1member, tmp11481)


_ = tmp11488

tmp11489 := MakeNative(func(__e *ControlFlow) {
V3619 := __e.Get(1)
_ = V3619
V3620 := __e.Get(2)
_ = V3620
tmp11503 := PrimEqual(Nil, V3619)

if True == tmp11503 {
__e.Return(Nil)
return
} else {
tmp11501 := PrimIsPair(V3619)

var ifres11497 Obj

if True == tmp11501 {
tmp11499 := PrimHead(V3619)

tmp11500 := Call(__e, PrimFunc(symshen_4occurs_1check_2), tmp11499, V3620)


var ifres11498 Obj

if True == tmp11500 {
ifres11498 = True


} else {
ifres11498 = False


}

ifres11497 = ifres11498


} else {
ifres11497 = False


}

if True == ifres11497 {
tmp11490 := PrimHead(V3619)

tmp11491 := PrimTail(V3619)

tmp11492 := Call(__e, PrimFunc(symshen_4remove_1bystanders), tmp11491, V3620)


__e.Return(PrimCons(tmp11490, tmp11492))
return


} else {
tmp11495 := PrimIsPair(V3619)

if True == tmp11495 {
tmp11493 := PrimTail(V3619)

__e.TailApply(PrimFunc(symshen_4remove_1bystanders), tmp11493, V3620)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4remove_1bystanders)
return
}


}


}


}, 2)

tmp11504 := Call(__e, ns2_1set, symshen_4remove_1bystanders, tmp11489)


_ = tmp11504

tmp11505 := MakeNative(func(__e *ControlFlow) {
V3621 := __e.Get(1)
_ = V3621
V3622 := __e.Get(2)
_ = V3622
V3623 := __e.Get(3)
_ = V3623
V3624 := __e.Get(4)
_ = V3624
tmp11506 := MakeNative(func(__e *ControlFlow) {
W3625 := __e.Get(1)
_ = W3625
tmp11507 := MakeNative(func(__e *ControlFlow) {
W3626 := __e.Get(1)
_ = W3626
tmp11508 := MakeNative(func(__e *ControlFlow) {
W3627 := __e.Get(1)
_ = W3627
tmp11509 := MakeNative(func(__e *ControlFlow) {
W3632 := __e.Get(1)
_ = W3632
__e.TailApply(PrimFunc(symeval), W3632)
return
}, 1)

tmp11510 := Call(__e, PrimFunc(symappend), W3626, W3627)


tmp11511 := PrimCons(V3621, tmp11510)

tmp11512 := PrimCons(symdefprolog, tmp11511)

__e.TailApply(tmp11509, tmp11512)
return


}, 1)

tmp11513 := MakeNative(func(__e *ControlFlow) {
W3628 := __e.Get(1)
_ = W3628
tmp11514 := MakeNative(func(__e *ControlFlow) {
W3629 := __e.Get(1)
_ = W3629
tmp11515 := MakeNative(func(__e *ControlFlow) {
W3630 := __e.Get(1)
_ = W3630
tmp11516 := MakeNative(func(__e *ControlFlow) {
W3631 := __e.Get(1)
_ = W3631
tmp11517 := PrimCons(sym_5_1_1, Nil)

tmp11518 := PrimIntern(MakeString(";"))

tmp11519 := PrimCons(tmp11518, Nil)

tmp11520 := Call(__e, PrimFunc(symappend), W3631, tmp11519)


tmp11521 := Call(__e, PrimFunc(symappend), tmp11517, tmp11520)


__e.TailApply(PrimFunc(symappend), W3630, tmp11521)
return


}, 1)

tmp11522 := PrimCons(W3628, W3629)

tmp11523 := PrimCons(V3621, tmp11522)

tmp11524 := PrimCons(tmp11523, Nil)

__e.TailApply(tmp11516, tmp11524)
return


}, 1)

tmp11525 := PrimCons(W3628, Nil)

tmp11526 := PrimCons(sym__, tmp11525)

tmp11527 := PrimCons(symcons, tmp11526)

tmp11528 := PrimCons(tmp11527, Nil)

tmp11529 := PrimCons(sym_1, tmp11528)

tmp11530 := PrimCons(tmp11529, Nil)

tmp11531 := Call(__e, PrimFunc(symappend), tmp11530, W3629)


__e.TailApply(tmp11515, tmp11531)
return


}, 1)

tmp11532 := Call(__e, PrimFunc(symappend), V3623, V3624)


__e.TailApply(tmp11514, tmp11532)
return


}, 1)

tmp11533 := Call(__e, PrimFunc(symgensym), symHypotheses)


tmp11534 := Call(__e, tmp11513, tmp11533)


__e.TailApply(tmp11508, tmp11534)
return


}, 1)

tmp11535 := Call(__e, PrimFunc(symshen_4macro_1_8ch), V3622)


tmp11536 := PrimCons(sym__, Nil)

tmp11537 := PrimCons(tmp11535, tmp11536)

tmp11538 := PrimCons(symcons, tmp11537)

tmp11539 := PrimCons(tmp11538, Nil)

tmp11540 := PrimCons(sym_1, tmp11539)

tmp11541 := PrimCons(tmp11540, Nil)

tmp11542 := PrimCons(sym_5_1_1, Nil)

tmp11543 := Call(__e, PrimFunc(symshen_4passive_1bind), V3624, W3625)


tmp11544 := PrimIntern(MakeString(";"))

tmp11545 := PrimCons(tmp11544, Nil)

tmp11546 := Call(__e, PrimFunc(symappend), tmp11543, tmp11545)


tmp11547 := Call(__e, PrimFunc(symappend), tmp11542, tmp11546)


tmp11548 := Call(__e, PrimFunc(symappend), W3625, tmp11547)


tmp11549 := Call(__e, PrimFunc(symappend), V3623, tmp11548)


tmp11550 := Call(__e, PrimFunc(symappend), tmp11541, tmp11549)


__e.TailApply(tmp11507, tmp11550)
return


}, 1)

tmp11551 := Call(__e, PrimFunc(symlength), V3624)


tmp11552 := Call(__e, PrimFunc(symshen_4nvars), tmp11551)


__e.TailApply(tmp11506, tmp11552)
return


}, 4)

tmp11553 := Call(__e, ns2_1set, symshen_4member_1clause, tmp11505)


_ = tmp11553

tmp11554 := MakeNative(func(__e *ControlFlow) {
V3633 := __e.Get(1)
_ = V3633
tmp11559 := PrimEqual(MakeNumber(0), V3633)

if True == tmp11559 {
__e.Return(Nil)
return
} else {
tmp11555 := Call(__e, PrimFunc(symgensym), symNewV)


tmp11556 := PrimNumberSubtract(V3633, MakeNumber(1))

tmp11557 := Call(__e, PrimFunc(symshen_4nvars), tmp11556)


__e.Return(PrimCons(tmp11555, tmp11557))
return


}


}, 1)

tmp11560 := Call(__e, ns2_1set, symshen_4nvars, tmp11554)


_ = tmp11560

tmp11561 := MakeNative(func(__e *ControlFlow) {
V3634 := __e.Get(1)
_ = V3634
V3635 := __e.Get(2)
_ = V3635
tmp11579 := PrimEqual(Nil, V3634)

var ifres11576 Obj

if True == tmp11579 {
tmp11578 := PrimEqual(Nil, V3635)

var ifres11577 Obj

if True == tmp11578 {
ifres11577 = True


} else {
ifres11577 = False


}

ifres11576 = ifres11577


} else {
ifres11576 = False


}

if True == ifres11576 {
__e.Return(Nil)
return
} else {
tmp11574 := PrimIsPair(V3634)

var ifres11571 Obj

if True == tmp11574 {
tmp11573 := PrimIsPair(V3635)

var ifres11572 Obj

if True == tmp11573 {
ifres11572 = True


} else {
ifres11572 = False


}

ifres11571 = ifres11572


} else {
ifres11571 = False


}

if True == ifres11571 {
tmp11562 := PrimHead(V3635)

tmp11563 := PrimHead(V3634)

tmp11564 := PrimCons(tmp11563, Nil)

tmp11565 := PrimCons(tmp11562, tmp11564)

tmp11566 := PrimCons(symbind, tmp11565)

tmp11567 := PrimTail(V3634)

tmp11568 := PrimTail(V3635)

tmp11569 := Call(__e, PrimFunc(symshen_4passive_1bind), tmp11567, tmp11568)


__e.Return(PrimCons(tmp11566, tmp11569))
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4passive_1bind)
return
}


}


}, 2)

tmp11580 := Call(__e, ns2_1set, symshen_4passive_1bind, tmp11561)


_ = tmp11580

tmp11581 := MakeNative(func(__e *ControlFlow) {
V3636 := __e.Get(1)
_ = V3636
V3637 := __e.Get(2)
_ = V3637
V3638 := __e.Get(3)
_ = V3638
V3639 := __e.Get(4)
_ = V3639
V3640 := __e.Get(5)
_ = V3640
tmp11582 := MakeNative(func(__e *ControlFlow) {
W3641 := __e.Get(1)
_ = W3641
tmp11583 := MakeNative(func(__e *ControlFlow) {
W3642 := __e.Get(1)
_ = W3642
tmp11584 := Call(__e, PrimFunc(symappend), V3638, V3639)


tmp11585 := PrimCons(V3640, tmp11584)

tmp11586 := PrimCons(V3637, tmp11585)

__e.Return(PrimCons(W3641, tmp11586))
return


}, 1)

tmp11587 := Call(__e, PrimFunc(symshen_4consume_1clause), W3641, V3636, V3638, V3639, V3640)


__e.TailApply(tmp11583, tmp11587)
return


}, 1)

tmp11588 := Call(__e, PrimFunc(symgensym), symshen_4consume)


__e.TailApply(tmp11582, tmp11588)
return


}, 5)

tmp11589 := Call(__e, ns2_1set, symshen_4specialise_1consume, tmp11581)


_ = tmp11589

tmp11590 := MakeNative(func(__e *ControlFlow) {
V3643 := __e.Get(1)
_ = V3643
V3644 := __e.Get(2)
_ = V3644
V3645 := __e.Get(3)
_ = V3645
V3646 := __e.Get(4)
_ = V3646
V3647 := __e.Get(5)
_ = V3647
tmp11591 := MakeNative(func(__e *ControlFlow) {
W3648 := __e.Get(1)
_ = W3648
tmp11592 := MakeNative(func(__e *ControlFlow) {
W3649 := __e.Get(1)
_ = W3649
tmp11593 := MakeNative(func(__e *ControlFlow) {
W3650 := __e.Get(1)
_ = W3650
tmp11594 := MakeNative(func(__e *ControlFlow) {
W3651 := __e.Get(1)
_ = W3651
tmp11595 := MakeNative(func(__e *ControlFlow) {
W3657 := __e.Get(1)
_ = W3657
__e.TailApply(PrimFunc(symeval), W3657)
return
}, 1)

tmp11596 := Call(__e, PrimFunc(symappend), W3650, W3651)


tmp11597 := PrimCons(V3643, tmp11596)

tmp11598 := PrimCons(symdefprolog, tmp11597)

__e.TailApply(tmp11595, tmp11598)
return


}, 1)

tmp11599 := MakeNative(func(__e *ControlFlow) {
W3652 := __e.Get(1)
_ = W3652
tmp11600 := MakeNative(func(__e *ControlFlow) {
W3653 := __e.Get(1)
_ = W3653
tmp11601 := MakeNative(func(__e *ControlFlow) {
W3654 := __e.Get(1)
_ = W3654
tmp11602 := MakeNative(func(__e *ControlFlow) {
W3655 := __e.Get(1)
_ = W3655
tmp11603 := MakeNative(func(__e *ControlFlow) {
W3656 := __e.Get(1)
_ = W3656
tmp11604 := PrimCons(sym_5_1_1, Nil)

tmp11605 := PrimIntern(MakeString(";"))

tmp11606 := PrimCons(tmp11605, Nil)

tmp11607 := Call(__e, PrimFunc(symappend), W3656, tmp11606)


tmp11608 := Call(__e, PrimFunc(symappend), tmp11604, tmp11607)


__e.TailApply(PrimFunc(symappend), W3655, tmp11608)
return


}, 1)

tmp11609 := PrimCons(W3649, Nil)

tmp11610 := PrimCons(W3654, tmp11609)

tmp11611 := PrimCons(symbind, tmp11610)

tmp11612 := PrimCons(V3647, W3653)

tmp11613 := PrimCons(W3652, tmp11612)

tmp11614 := PrimCons(V3643, tmp11613)

tmp11615 := PrimCons(tmp11614, Nil)

tmp11616 := PrimCons(tmp11611, tmp11615)

__e.TailApply(tmp11603, tmp11616)
return


}, 1)

tmp11617 := PrimCons(W3652, Nil)

tmp11618 := PrimCons(W3649, tmp11617)

tmp11619 := PrimCons(symcons, tmp11618)

tmp11620 := PrimCons(tmp11619, Nil)

tmp11621 := PrimCons(sym_1, tmp11620)

tmp11622 := PrimCons(V3647, Nil)

tmp11623 := PrimCons(W3654, tmp11622)

tmp11624 := PrimCons(symcons, tmp11623)

tmp11625 := PrimCons(tmp11624, W3653)

tmp11626 := PrimCons(tmp11621, tmp11625)

__e.TailApply(tmp11602, tmp11626)
return


}, 1)

tmp11627 := Call(__e, PrimFunc(symgensym), symAssumptions)


__e.TailApply(tmp11601, tmp11627)
return


}, 1)

tmp11628 := Call(__e, PrimFunc(symappend), V3645, V3646)


__e.TailApply(tmp11600, tmp11628)
return


}, 1)

tmp11629 := Call(__e, PrimFunc(symgensym), symHypotheses)


tmp11630 := Call(__e, tmp11599, tmp11629)


__e.TailApply(tmp11594, tmp11630)
return


}, 1)

tmp11631 := Call(__e, PrimFunc(symshen_4macro_1_8ch), V3644)


tmp11632 := PrimCons(W3649, Nil)

tmp11633 := PrimCons(tmp11631, tmp11632)

tmp11634 := PrimCons(symcons, tmp11633)

tmp11635 := PrimCons(tmp11634, Nil)

tmp11636 := PrimCons(sym_1, tmp11635)

tmp11637 := PrimCons(sym_5_1_1, Nil)

tmp11638 := Call(__e, PrimFunc(symshen_4passive_1bind), V3646, W3648)


tmp11639 := PrimCons(W3649, Nil)

tmp11640 := PrimCons(V3647, tmp11639)

tmp11641 := PrimCons(symbind, tmp11640)

tmp11642 := PrimIntern(MakeString(";"))

tmp11643 := PrimCons(tmp11642, Nil)

tmp11644 := PrimCons(tmp11641, tmp11643)

tmp11645 := Call(__e, PrimFunc(symappend), tmp11638, tmp11644)


tmp11646 := Call(__e, PrimFunc(symappend), tmp11637, tmp11645)


tmp11647 := Call(__e, PrimFunc(symappend), W3648, tmp11646)


tmp11648 := Call(__e, PrimFunc(symappend), V3645, tmp11647)


tmp11649 := PrimCons(V3647, tmp11648)

tmp11650 := PrimCons(tmp11636, tmp11649)

__e.TailApply(tmp11593, tmp11650)
return


}, 1)

tmp11651 := Call(__e, PrimFunc(symgensym), symAssumption)


__e.TailApply(tmp11592, tmp11651)
return


}, 1)

tmp11652 := Call(__e, PrimFunc(symlength), V3646)


tmp11653 := Call(__e, PrimFunc(symshen_4nvars), tmp11652)


__e.TailApply(tmp11591, tmp11653)
return


}, 5)

tmp11654 := Call(__e, ns2_1set, symshen_4consume_1clause, tmp11590)


_ = tmp11654

tmp11655 := MakeNative(func(__e *ControlFlow) {
V3658 := __e.Get(1)
_ = V3658
V3659 := __e.Get(2)
_ = V3659
tmp11656 := Call(__e, PrimFunc(symshen_4extract_1vars), V3658)


__e.TailApply(PrimFunc(symdifference), tmp11656, V3659)
return


}, 2)

tmp11657 := Call(__e, ns2_1set, symshen_4passive_1variables, tmp11655)


_ = tmp11657

tmp11658 := MakeNative(func(__e *ControlFlow) {
V3662 := __e.Get(1)
_ = V3662
V3663 := __e.Get(2)
_ = V3663
V3664 := __e.Get(3)
_ = V3664
V3665 := __e.Get(4)
_ = V3665
V3666 := __e.Get(5)
_ = V3666
tmp11786 := PrimEqual(Nil, V3665)

if True == tmp11786 {
__e.TailApply(PrimFunc(symshen_4premises_1_6goals), V3662, V3664, V3666)
return
} else {
tmp11784 := PrimIsPair(V3665)

var ifres11764 Obj

if True == tmp11784 {
tmp11782 := PrimHead(V3665)

tmp11783 := PrimIsPair(tmp11782)

var ifres11766 Obj

if True == tmp11783 {
tmp11779 := PrimHead(V3665)

tmp11780 := PrimHead(tmp11779)

tmp11781 := PrimEqual(symif, tmp11780)

var ifres11768 Obj

if True == tmp11781 {
tmp11776 := PrimHead(V3665)

tmp11777 := PrimTail(tmp11776)

tmp11778 := PrimIsPair(tmp11777)

var ifres11770 Obj

if True == tmp11778 {
tmp11772 := PrimHead(V3665)

tmp11773 := PrimTail(tmp11772)

tmp11774 := PrimTail(tmp11773)

tmp11775 := PrimEqual(Nil, tmp11774)

var ifres11771 Obj

if True == tmp11775 {
ifres11771 = True


} else {
ifres11771 = False


}

ifres11770 = ifres11771


} else {
ifres11770 = False


}

var ifres11769 Obj

if True == ifres11770 {
ifres11769 = True


} else {
ifres11769 = False


}

ifres11768 = ifres11769


} else {
ifres11768 = False


}

var ifres11767 Obj

if True == ifres11768 {
ifres11767 = True


} else {
ifres11767 = False


}

ifres11766 = ifres11767


} else {
ifres11766 = False


}

var ifres11765 Obj

if True == ifres11766 {
ifres11765 = True


} else {
ifres11765 = False


}

ifres11764 = ifres11765


} else {
ifres11764 = False


}

if True == ifres11764 {
tmp11659 := PrimHead(V3665)

tmp11660 := PrimTail(tmp11659)

tmp11661 := PrimCons(symwhen, tmp11660)

tmp11662 := PrimTail(V3665)

tmp11663 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), V3662, V3663, V3664, tmp11662, V3666)


__e.Return(PrimCons(tmp11661, tmp11663))
return


} else {
tmp11762 := PrimIsPair(V3665)

var ifres11735 Obj

if True == tmp11762 {
tmp11760 := PrimHead(V3665)

tmp11761 := PrimIsPair(tmp11760)

var ifres11737 Obj

if True == tmp11761 {
tmp11757 := PrimHead(V3665)

tmp11758 := PrimHead(tmp11757)

tmp11759 := PrimEqual(symlet, tmp11758)

var ifres11739 Obj

if True == tmp11759 {
tmp11754 := PrimHead(V3665)

tmp11755 := PrimTail(tmp11754)

tmp11756 := PrimIsPair(tmp11755)

var ifres11741 Obj

if True == tmp11756 {
tmp11750 := PrimHead(V3665)

tmp11751 := PrimTail(tmp11750)

tmp11752 := PrimTail(tmp11751)

tmp11753 := PrimIsPair(tmp11752)

var ifres11743 Obj

if True == tmp11753 {
tmp11745 := PrimHead(V3665)

tmp11746 := PrimTail(tmp11745)

tmp11747 := PrimTail(tmp11746)

tmp11748 := PrimTail(tmp11747)

tmp11749 := PrimEqual(Nil, tmp11748)

var ifres11744 Obj

if True == tmp11749 {
ifres11744 = True


} else {
ifres11744 = False


}

ifres11743 = ifres11744


} else {
ifres11743 = False


}

var ifres11742 Obj

if True == ifres11743 {
ifres11742 = True


} else {
ifres11742 = False


}

ifres11741 = ifres11742


} else {
ifres11741 = False


}

var ifres11740 Obj

if True == ifres11741 {
ifres11740 = True


} else {
ifres11740 = False


}

ifres11739 = ifres11740


} else {
ifres11739 = False


}

var ifres11738 Obj

if True == ifres11739 {
ifres11738 = True


} else {
ifres11738 = False


}

ifres11737 = ifres11738


} else {
ifres11737 = False


}

var ifres11736 Obj

if True == ifres11737 {
ifres11736 = True


} else {
ifres11736 = False


}

ifres11735 = ifres11736


} else {
ifres11735 = False


}

if True == ifres11735 {
tmp11679 := PrimHead(V3665)

tmp11680 := PrimTail(tmp11679)

tmp11681 := PrimHead(tmp11680)

tmp11682 := Call(__e, PrimFunc(symelement_2), tmp11681, V3663)


if True == tmp11682 {
tmp11664 := PrimHead(V3665)

tmp11665 := PrimTail(tmp11664)

tmp11666 := PrimCons(symis_b, tmp11665)

tmp11667 := PrimTail(V3665)

tmp11668 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), V3662, V3663, V3664, tmp11667, V3666)


__e.Return(PrimCons(tmp11666, tmp11668))
return


} else {
tmp11669 := PrimHead(V3665)

tmp11670 := PrimTail(tmp11669)

tmp11671 := PrimCons(symbind, tmp11670)

tmp11672 := PrimHead(V3665)

tmp11673 := PrimTail(tmp11672)

tmp11674 := PrimHead(tmp11673)

tmp11675 := PrimCons(tmp11674, V3663)

tmp11676 := PrimTail(V3665)

tmp11677 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), V3662, tmp11675, V3664, tmp11676, V3666)


__e.Return(PrimCons(tmp11671, tmp11677))
return


}


} else {
tmp11733 := PrimIsPair(V3665)

var ifres11713 Obj

if True == tmp11733 {
tmp11731 := PrimHead(V3665)

tmp11732 := PrimIsPair(tmp11731)

var ifres11715 Obj

if True == tmp11732 {
tmp11728 := PrimHead(V3665)

tmp11729 := PrimHead(tmp11728)

tmp11730 := PrimEqual(symctxt, tmp11729)

var ifres11717 Obj

if True == tmp11730 {
tmp11725 := PrimHead(V3665)

tmp11726 := PrimTail(tmp11725)

tmp11727 := PrimIsPair(tmp11726)

var ifres11719 Obj

if True == tmp11727 {
tmp11721 := PrimHead(V3665)

tmp11722 := PrimTail(tmp11721)

tmp11723 := PrimTail(tmp11722)

tmp11724 := PrimEqual(Nil, tmp11723)

var ifres11720 Obj

if True == tmp11724 {
ifres11720 = True


} else {
ifres11720 = False


}

ifres11719 = ifres11720


} else {
ifres11719 = False


}

var ifres11718 Obj

if True == ifres11719 {
ifres11718 = True


} else {
ifres11718 = False


}

ifres11717 = ifres11718


} else {
ifres11717 = False


}

var ifres11716 Obj

if True == ifres11717 {
ifres11716 = True


} else {
ifres11716 = False


}

ifres11715 = ifres11716


} else {
ifres11715 = False


}

var ifres11714 Obj

if True == ifres11715 {
ifres11714 = True


} else {
ifres11714 = False


}

ifres11713 = ifres11714


} else {
ifres11713 = False


}

if True == ifres11713 {
tmp11708 := PrimHead(V3665)

tmp11709 := PrimTail(tmp11708)

tmp11710 := PrimHead(tmp11709)

tmp11711 := Call(__e, PrimFunc(symelement_2), tmp11710, V3663)


if True == tmp11711 {
tmp11683 := PrimHead(V3665)

tmp11684 := PrimTail(tmp11683)

tmp11685 := PrimHead(tmp11684)

tmp11686 := PrimCons(tmp11685, V3662)

tmp11687 := PrimTail(V3665)

__e.TailApply(PrimFunc(symshen_4side_1conditions_1_6goals), tmp11686, V3663, V3664, tmp11687, V3666)
return


} else {
tmp11688 := PrimHead(V3665)

tmp11689 := PrimTail(tmp11688)

tmp11690 := PrimHead(tmp11689)

tmp11691 := PrimCons(V3664, Nil)

tmp11692 := PrimCons(tmp11690, tmp11691)

tmp11693 := PrimCons(symbind, tmp11692)

tmp11694 := PrimHead(V3665)

tmp11695 := PrimTail(tmp11694)

tmp11696 := PrimHead(tmp11695)

tmp11697 := PrimCons(tmp11696, V3662)

tmp11698 := PrimHead(V3665)

tmp11699 := PrimTail(tmp11698)

tmp11700 := PrimHead(tmp11699)

tmp11701 := PrimCons(tmp11700, V3663)

tmp11702 := PrimHead(V3665)

tmp11703 := PrimTail(tmp11702)

tmp11704 := PrimHead(tmp11703)

tmp11705 := PrimTail(V3665)

tmp11706 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), tmp11697, tmp11701, tmp11704, tmp11705, V3666)


__e.Return(PrimCons(tmp11693, tmp11706))
return


}


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4side_1conditions_1_6goals)
return
}


}


}


}


}, 5)

tmp11787 := Call(__e, ns2_1set, symshen_4side_1conditions_1_6goals, tmp11658)


_ = tmp11787

tmp11788 := MakeNative(func(__e *ControlFlow) {
V3671 := __e.Get(1)
_ = V3671
V3672 := __e.Get(2)
_ = V3672
V3673 := __e.Get(3)
_ = V3673
tmp11838 := PrimEqual(Nil, V3673)

if True == tmp11838 {
tmp11789 := PrimIntern(MakeString(";"))

__e.Return(PrimCons(tmp11789, Nil))
return


} else {
tmp11836 := PrimIsPair(V3673)

var ifres11832 Obj

if True == tmp11836 {
tmp11834 := PrimHead(V3673)

tmp11835 := PrimEqual(sym_b, tmp11834)

var ifres11833 Obj

if True == tmp11835 {
ifres11833 = True


} else {
ifres11833 = False


}

ifres11832 = ifres11833


} else {
ifres11832 = False


}

if True == ifres11832 {
tmp11790 := PrimTail(V3673)

tmp11791 := Call(__e, PrimFunc(symshen_4premises_1_6goals), V3671, V3672, tmp11790)


__e.Return(PrimCons(sym_b, tmp11791))
return


} else {
tmp11830 := PrimIsPair(V3673)

var ifres11826 Obj

if True == tmp11830 {
tmp11828 := PrimHead(V3673)

tmp11829 := PrimEqual(symfail, tmp11828)

var ifres11827 Obj

if True == tmp11829 {
ifres11827 = True


} else {
ifres11827 = False


}

ifres11826 = ifres11827


} else {
ifres11826 = False


}

if True == ifres11826 {
tmp11792 := PrimCons(False, Nil)

tmp11793 := PrimCons(symwhen, tmp11792)

tmp11794 := PrimTail(V3673)

tmp11795 := Call(__e, PrimFunc(symshen_4premises_1_6goals), V3671, V3672, tmp11794)


__e.Return(PrimCons(tmp11793, tmp11795))
return


} else {
tmp11824 := PrimIsPair(V3673)

var ifres11809 Obj

if True == tmp11824 {
tmp11822 := PrimHead(V3673)

tmp11823 := PrimIsPair(tmp11822)

var ifres11811 Obj

if True == tmp11823 {
tmp11819 := PrimHead(V3673)

tmp11820 := PrimTail(tmp11819)

tmp11821 := PrimIsPair(tmp11820)

var ifres11813 Obj

if True == tmp11821 {
tmp11815 := PrimHead(V3673)

tmp11816 := PrimTail(tmp11815)

tmp11817 := PrimTail(tmp11816)

tmp11818 := PrimEqual(Nil, tmp11817)

var ifres11814 Obj

if True == tmp11818 {
ifres11814 = True


} else {
ifres11814 = False


}

ifres11813 = ifres11814


} else {
ifres11813 = False


}

var ifres11812 Obj

if True == ifres11813 {
ifres11812 = True


} else {
ifres11812 = False


}

ifres11811 = ifres11812


} else {
ifres11811 = False


}

var ifres11810 Obj

if True == ifres11811 {
ifres11810 = True


} else {
ifres11810 = False


}

ifres11809 = ifres11810


} else {
ifres11809 = False


}

if True == ifres11809 {
tmp11796 := PrimHead(V3673)

tmp11797 := PrimTail(tmp11796)

tmp11798 := PrimHead(tmp11797)

tmp11799 := Call(__e, PrimFunc(symshen_4macro_1_8c), tmp11798)


tmp11800 := PrimHead(V3673)

tmp11801 := PrimHead(tmp11800)

tmp11802 := Call(__e, PrimFunc(symshen_4construct_1context), V3671, tmp11801, V3672)


tmp11803 := PrimCons(tmp11802, Nil)

tmp11804 := PrimCons(tmp11799, tmp11803)

tmp11805 := PrimCons(symshen_4system_1S, tmp11804)

tmp11806 := PrimTail(V3673)

tmp11807 := Call(__e, PrimFunc(symshen_4premises_1_6goals), V3671, V3672, tmp11806)


__e.Return(PrimCons(tmp11805, tmp11807))
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4premises_1_6goals)
return
}


}


}


}


}, 3)

tmp11839 := Call(__e, ns2_1set, symshen_4premises_1_6goals, tmp11788)


_ = tmp11839

tmp11840 := MakeNative(func(__e *ControlFlow) {
V3677 := __e.Get(1)
_ = V3677
V3678 := __e.Get(2)
_ = V3678
V3679 := __e.Get(3)
_ = V3679
tmp11860 := PrimEqual(Nil, V3678)

if True == tmp11860 {
__e.Return(V3679)
return
} else {
tmp11858 := PrimIsPair(V3678)

var ifres11850 Obj

if True == tmp11858 {
tmp11856 := PrimTail(V3678)

tmp11857 := PrimEqual(Nil, tmp11856)

var ifres11852 Obj

if True == tmp11857 {
tmp11854 := PrimHead(V3678)

tmp11855 := Call(__e, PrimFunc(symelement_2), tmp11854, V3677)


var ifres11853 Obj

if True == tmp11855 {
ifres11853 = True


} else {
ifres11853 = False


}

ifres11852 = ifres11853


} else {
ifres11852 = False


}

var ifres11851 Obj

if True == ifres11852 {
ifres11851 = True


} else {
ifres11851 = False


}

ifres11850 = ifres11851


} else {
ifres11850 = False


}

if True == ifres11850 {
__e.Return(PrimHead(V3678))
return
} else {
tmp11848 := PrimIsPair(V3678)

if True == tmp11848 {
tmp11841 := PrimHead(V3678)

tmp11842 := Call(__e, PrimFunc(symshen_4macro_1_8c), tmp11841)


tmp11843 := PrimTail(V3678)

tmp11844 := Call(__e, PrimFunc(symshen_4construct_1context), V3677, tmp11843, V3679)


tmp11845 := PrimCons(tmp11844, Nil)

tmp11846 := PrimCons(tmp11842, tmp11845)

__e.Return(PrimCons(symcons, tmp11846))
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4construct_1context)
return
}


}


}


}, 3)

tmp11861 := Call(__e, ns2_1set, symshen_4construct_1context, tmp11840)


_ = tmp11861

tmp11862 := MakeNative(func(__e *ControlFlow) {
V3680 := __e.Get(1)
_ = V3680
tmp11863 := MakeNative(func(__e *ControlFlow) {
W3681 := __e.Get(1)
_ = W3681
tmp11864 := MakeNative(func(__e *ControlFlow) {
W3683 := __e.Get(1)
_ = W3683
tmp11865 := MakeNative(func(__e *ControlFlow) {
W3684 := __e.Get(1)
_ = W3684
tmp11866 := MakeNative(func(__e *ControlFlow) {
W3685 := __e.Get(1)
_ = W3685
__e.TailApply(PrimFunc(symshen_4show_1datatypes), W3685)
return
}, 1)

tmp11867 := PrimSet(symshen_4_ddatatypes_d, W3684)

__e.TailApply(tmp11866, tmp11867)
return


}, 1)

tmp11868 := Call(__e, PrimFunc(symshen_4remove_1datatypes), W3681, W3683)


__e.TailApply(tmp11865, tmp11868)
return


}, 1)

tmp11869 := PrimValue(symshen_4_ddatatypes_d)

__e.TailApply(tmp11864, tmp11869)
return


}, 1)

tmp11870 := MakeNative(func(__e *ControlFlow) {
Z3682 := __e.Get(1)
_ = Z3682
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3682)
return
}, 1)

tmp11871 := Call(__e, PrimFunc(symmap), tmp11870, V3680)


__e.TailApply(tmp11863, tmp11871)
return


}, 1)

tmp11872 := Call(__e, ns2_1set, sympreclude, tmp11862)


_ = tmp11872

tmp11873 := MakeNative(func(__e *ControlFlow) {
V3690 := __e.Get(1)
_ = V3690
V3691 := __e.Get(2)
_ = V3691
tmp11880 := PrimEqual(Nil, V3690)

if True == tmp11880 {
__e.Return(V3691)
return
} else {
tmp11878 := PrimIsPair(V3690)

if True == tmp11878 {
tmp11874 := PrimTail(V3690)

tmp11875 := PrimHead(V3690)

tmp11876 := Call(__e, PrimFunc(symshen_4unassoc), tmp11875, V3691)


__e.TailApply(PrimFunc(symshen_4remove_1datatypes), tmp11874, tmp11876)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.remove-datatypes")))
return
}


}


}, 2)

tmp11881 := Call(__e, ns2_1set, symshen_4remove_1datatypes, tmp11873)


_ = tmp11881

tmp11882 := MakeNative(func(__e *ControlFlow) {
V3692 := __e.Get(1)
_ = V3692
tmp11883 := MakeNative(func(__e *ControlFlow) {
Z3693 := __e.Get(1)
_ = Z3693
__e.Return(PrimHead(Z3693))
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp11883, V3692)
return


}, 1)

tmp11884 := Call(__e, ns2_1set, symshen_4show_1datatypes, tmp11882)


_ = tmp11884

tmp11885 := MakeNative(func(__e *ControlFlow) {
V3694 := __e.Get(1)
_ = V3694
tmp11886 := MakeNative(func(__e *ControlFlow) {
W3695 := __e.Get(1)
_ = W3695
tmp11887 := MakeNative(func(__e *ControlFlow) {
W3697 := __e.Get(1)
_ = W3697
tmp11888 := MakeNative(func(__e *ControlFlow) {
W3699 := __e.Get(1)
_ = W3699
__e.TailApply(PrimFunc(symshen_4show_1datatypes), W3699)
return
}, 1)

tmp11889 := PrimValue(symshen_4_ddatatypes_d)

__e.TailApply(tmp11888, tmp11889)
return


}, 1)

tmp11890 := MakeNative(func(__e *ControlFlow) {
Z3698 := __e.Get(1)
_ = Z3698
tmp11891 := Call(__e, PrimFunc(symfn), Z3698)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), Z3698, tmp11891)
return


}, 1)

tmp11892 := Call(__e, PrimFunc(symmap), tmp11890, W3695)


__e.TailApply(tmp11887, tmp11892)
return


}, 1)

tmp11893 := MakeNative(func(__e *ControlFlow) {
Z3696 := __e.Get(1)
_ = Z3696
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3696)
return
}, 1)

tmp11894 := Call(__e, PrimFunc(symmap), tmp11893, V3694)


__e.TailApply(tmp11886, tmp11894)
return


}, 1)

tmp11895 := Call(__e, ns2_1set, syminclude, tmp11885)


_ = tmp11895

tmp11896 := MakeNative(func(__e *ControlFlow) {
V3700 := __e.Get(1)
_ = V3700
tmp11897 := MakeNative(func(__e *ControlFlow) {
W3701 := __e.Get(1)
_ = W3701
tmp11898 := MakeNative(func(__e *ControlFlow) {
W3702 := __e.Get(1)
_ = W3702
tmp11899 := MakeNative(func(__e *ControlFlow) {
W3704 := __e.Get(1)
_ = W3704
tmp11900 := PrimValue(symshen_4_ddatatypes_d)

__e.TailApply(PrimFunc(symshen_4show_1datatypes), tmp11900)
return


}, 1)

tmp11901 := MakeNative(func(__e *ControlFlow) {
Z3705 := __e.Get(1)
_ = Z3705
tmp11902 := Call(__e, PrimFunc(symfn), Z3705)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), Z3705, tmp11902)
return


}, 1)

tmp11903 := Call(__e, PrimFunc(symmap), tmp11901, W3702)


__e.TailApply(tmp11899, tmp11903)
return


}, 1)

tmp11904 := MakeNative(func(__e *ControlFlow) {
Z3703 := __e.Get(1)
_ = Z3703
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3703)
return
}, 1)

tmp11905 := Call(__e, PrimFunc(symmap), tmp11904, V3700)


__e.TailApply(tmp11898, tmp11905)
return


}, 1)

tmp11906 := PrimSet(symshen_4_ddatatypes_d, Nil)

__e.TailApply(tmp11897, tmp11906)
return


}, 1)

tmp11907 := Call(__e, ns2_1set, sympreclude_1all_1but, tmp11896)


_ = tmp11907

tmp11908 := MakeNative(func(__e *ControlFlow) {
V3706 := __e.Get(1)
_ = V3706
tmp11909 := MakeNative(func(__e *ControlFlow) {
W3707 := __e.Get(1)
_ = W3707
tmp11910 := MakeNative(func(__e *ControlFlow) {
W3709 := __e.Get(1)
_ = W3709
tmp11911 := MakeNative(func(__e *ControlFlow) {
W3710 := __e.Get(1)
_ = W3710
__e.TailApply(PrimFunc(symshen_4show_1datatypes), W3710)
return
}, 1)

tmp11912 := Call(__e, PrimFunc(symshen_4remove_1datatypes), W3707, W3709)


tmp11913 := PrimSet(symshen_4_ddatatypes_d, tmp11912)

__e.TailApply(tmp11911, tmp11913)
return


}, 1)

tmp11914 := PrimValue(symshen_4_dalldatatypes_d)

__e.TailApply(tmp11910, tmp11914)
return


}, 1)

tmp11915 := MakeNative(func(__e *ControlFlow) {
Z3708 := __e.Get(1)
_ = Z3708
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3708)
return
}, 1)

tmp11916 := Call(__e, PrimFunc(symmap), tmp11915, V3706)


__e.TailApply(tmp11909, tmp11916)
return


}, 1)

__e.TailApply(ns2_1set, syminclude_1all_1but, tmp11908)
return




}, 0)

