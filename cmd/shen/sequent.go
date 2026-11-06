package main

import . "github.com/pyrex41/shen-go/kl"

var SequentMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp11613 := MakeNative1(func(__e *ControlFlow, V3418 Obj) Obj {
tmp11628 := GConsP(__e, V3418)

var ifres11614 Obj

if True == tmp11628 {
tmp11615 := Call1(__e, SymFn(symhead), V3418)


let__11383 := tmp11615
_ = let__11383

tmp11616 := Call1(__e, SymFn(symtail), V3418)


let__11384 := tmp11616
_ = let__11384

tmp11617 := Call1(__e, SymFn(symshen_4_5datatype_1rules_6), let__11384)


let__11385 := tmp11617
_ = let__11385

tmp11626 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11385)


var ifres11618 Obj

if True == tmp11626 {
tmp11619 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11618 = tmp11619


} else {
tmp11620 := Call1(__e, SymFn(symshen_4_5_1out), let__11385)


let__11386 := tmp11620
_ = let__11386

tmp11621 := Call1(__e, SymFn(symshen_4in_1_6), let__11385)


let__11387 := tmp11621
_ = let__11387

tmp11622 := Call2(__e, SymFn(symshen_4rules_1_6prolog), let__11383, let__11386)


let__11388 := tmp11622
_ = let__11388

tmp11623 := Call1(__e, SymFn(symfn), let__11383)


tmp11624 := Call2(__e, SymFn(symshen_4remember_1datatype), let__11383, tmp11623)


tmp11625 := Call2(__e, SymFn(symshen_4comb), let__11387, tmp11624)


ifres11618 = tmp11625


}

ifres11614 = ifres11618


} else {
tmp11627 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11614 = tmp11627


}

let__11382 := ifres11614
_ = let__11382

tmp11630 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11382)


if True == tmp11630 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11382
}


panic("unreachable")
})

tmp11631 := Call2(__e, ns2_1set, symshen_4_5datatype_6, tmp11613)


_ = tmp11631

tmp11632 := MakeNative2(func(__e *ControlFlow, V3426, V3427 Obj) Obj {
tmp11633 := GValue(__e, symshen_4_ddatatypes_d)

tmp11634 := Call3(__e, SymFn(symshen_4assoc_1_6), V3426, V3427, tmp11633)


tmp11635 := PrimGuard2(__e, symset, PrimSet, symshen_4_ddatatypes_d, tmp11634)

_ = tmp11635

tmp11636 := GValue(__e, symshen_4_dalldatatypes_d)

tmp11637 := Call3(__e, SymFn(symshen_4assoc_1_6), V3426, V3427, tmp11636)


tmp11638 := PrimGuard2(__e, symset, PrimSet, symshen_4_dalldatatypes_d, tmp11637)

_ = tmp11638

return V3426


panic("unreachable")
})

tmp11639 := Call2(__e, ns2_1set, symshen_4remember_1datatype, tmp11632)


_ = tmp11639

tmp11640 := MakeNative1(func(__e *ControlFlow, V3428 Obj) Obj {
tmp11641 := Call1(__e, SymFn(symshen_4_5datatype_1rule_6), V3428)


let__11390 := tmp11641
_ = let__11390

tmp11654 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11390)


var ifres11642 Obj

if True == tmp11654 {
tmp11643 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11642 = tmp11643


} else {
tmp11644 := Call1(__e, SymFn(symshen_4_5_1out), let__11390)


let__11391 := tmp11644
_ = let__11391

tmp11645 := Call1(__e, SymFn(symshen_4in_1_6), let__11390)


let__11392 := tmp11645
_ = let__11392

tmp11646 := Call1(__e, SymFn(symshen_4_5datatype_1rules_6), let__11392)


let__11393 := tmp11646
_ = let__11393

tmp11653 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11393)


var ifres11647 Obj

if True == tmp11653 {
tmp11648 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11647 = tmp11648


} else {
tmp11649 := Call1(__e, SymFn(symshen_4_5_1out), let__11393)


let__11394 := tmp11649
_ = let__11394

tmp11650 := Call1(__e, SymFn(symshen_4in_1_6), let__11393)


let__11395 := tmp11650
_ = let__11395

tmp11651 := Call2(__e, SymFn(symappend), let__11391, let__11394)


tmp11652 := Call2(__e, SymFn(symshen_4comb), let__11395, tmp11651)


ifres11647 = tmp11652


}

ifres11642 = ifres11647


}

let__11389 := ifres11642
_ = let__11389

tmp11670 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11389)


if True == tmp11670 {
tmp11655 := Call1(__e, SymFn(sym_5_b_6), V3428)


let__11397 := tmp11655
_ = let__11397

tmp11666 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11397)


var ifres11656 Obj

if True == tmp11666 {
tmp11657 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11656 = tmp11657


} else {
tmp11658 := Call1(__e, SymFn(symshen_4_5_1out), let__11397)


let__11398 := tmp11658
_ = let__11398

tmp11659 := Call1(__e, SymFn(symshen_4in_1_6), let__11397)


let__11399 := tmp11659
_ = let__11399

tmp11664 := Call1(__e, SymFn(symempty_2), let__11398)


var ifres11660 Obj

if True == tmp11664 {
ifres11660 = Nil


} else {
tmp11661 := Call3(__e, SymFn(symshen_4app), let__11398, MakeString("\n ..."), symshen_4r)


tmp11663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("datatype syntax error here:\n "), tmp11661))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("datatype syntax error here:\n "), tmp11661)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()

ifres11660 = tmp11663


}

tmp11665 := Call2(__e, SymFn(symshen_4comb), let__11399, ifres11660)


ifres11656 = tmp11665


}

let__11396 := ifres11656
_ = let__11396

tmp11668 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11396)


if True == tmp11668 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11396
}


} else {
return let__11389
}


panic("unreachable")
})

tmp11671 := Call2(__e, ns2_1set, symshen_4_5datatype_1rules_6, tmp11640)


_ = tmp11671

tmp11672 := MakeNative1(func(__e *ControlFlow, V3440 Obj) Obj {
tmp11673 := Call1(__e, SymFn(symshen_4_5single_6), V3440)


let__11401 := tmp11673
_ = let__11401

tmp11679 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11401)


var ifres11674 Obj

if True == tmp11679 {
tmp11675 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11674 = tmp11675


} else {
tmp11676 := Call1(__e, SymFn(symshen_4_5_1out), let__11401)


let__11402 := tmp11676
_ = let__11402

tmp11677 := Call1(__e, SymFn(symshen_4in_1_6), let__11401)


let__11403 := tmp11677
_ = let__11403

tmp11678 := Call2(__e, SymFn(symshen_4comb), let__11403, let__11402)


ifres11674 = tmp11678


}

let__11400 := ifres11674
_ = let__11400

tmp11690 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11400)


if True == tmp11690 {
tmp11680 := Call1(__e, SymFn(symshen_4_5double_6), V3440)


let__11405 := tmp11680
_ = let__11405

tmp11686 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11405)


var ifres11681 Obj

if True == tmp11686 {
tmp11682 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11681 = tmp11682


} else {
tmp11683 := Call1(__e, SymFn(symshen_4_5_1out), let__11405)


let__11406 := tmp11683
_ = let__11406

tmp11684 := Call1(__e, SymFn(symshen_4in_1_6), let__11405)


let__11407 := tmp11684
_ = let__11407

tmp11685 := Call2(__e, SymFn(symshen_4comb), let__11407, let__11406)


ifres11681 = tmp11685


}

let__11404 := ifres11681
_ = let__11404

tmp11688 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11404)


if True == tmp11688 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11404
}


} else {
return let__11400
}


panic("unreachable")
})

tmp11691 := Call2(__e, ns2_1set, symshen_4_5datatype_1rule_6, tmp11672)


_ = tmp11691

tmp11692 := MakeNative1(func(__e *ControlFlow, V3449 Obj) Obj {
tmp11693 := Call1(__e, SymFn(symshen_4_5sides_6), V3449)


let__11409 := tmp11693
_ = let__11409

tmp11725 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11409)


var ifres11694 Obj

if True == tmp11725 {
tmp11695 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11694 = tmp11695


} else {
tmp11696 := Call1(__e, SymFn(symshen_4_5_1out), let__11409)


let__11410 := tmp11696
_ = let__11410

tmp11697 := Call1(__e, SymFn(symshen_4in_1_6), let__11409)


let__11411 := tmp11697
_ = let__11411

tmp11698 := Call1(__e, SymFn(symshen_4_5prems_6), let__11411)


let__11412 := tmp11698
_ = let__11412

tmp11724 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11412)


var ifres11699 Obj

if True == tmp11724 {
tmp11700 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11699 = tmp11700


} else {
tmp11701 := Call1(__e, SymFn(symshen_4_5_1out), let__11412)


let__11413 := tmp11701
_ = let__11413

tmp11702 := Call1(__e, SymFn(symshen_4in_1_6), let__11412)


let__11414 := tmp11702
_ = let__11414

tmp11703 := Call1(__e, SymFn(symshen_4_5sng_6), let__11414)


let__11415 := tmp11703
_ = let__11415

tmp11723 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11415)


var ifres11704 Obj

if True == tmp11723 {
tmp11705 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11704 = tmp11705


} else {
tmp11706 := Call1(__e, SymFn(symshen_4in_1_6), let__11415)


let__11416 := tmp11706
_ = let__11416

tmp11707 := Call1(__e, SymFn(symshen_4_5conc_6), let__11416)


let__11417 := tmp11707
_ = let__11417

tmp11722 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11417)


var ifres11708 Obj

if True == tmp11722 {
tmp11709 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11708 = tmp11709


} else {
tmp11710 := Call1(__e, SymFn(symshen_4_5_1out), let__11417)


let__11418 := tmp11710
_ = let__11418

tmp11711 := Call1(__e, SymFn(symshen_4in_1_6), let__11417)


let__11419 := tmp11711
_ = let__11419

tmp11712 := Call1(__e, SymFn(symshen_4_5sc_6), let__11419)


let__11420 := tmp11712
_ = let__11420

tmp11721 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11420)


var ifres11713 Obj

if True == tmp11721 {
tmp11714 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11713 = tmp11714


} else {
tmp11715 := Call1(__e, SymFn(symshen_4in_1_6), let__11420)


let__11421 := tmp11715
_ = let__11421

tmp11716 := GCons(__e, let__11418, Nil)

tmp11717 := GCons(__e, let__11413, tmp11716)

tmp11718 := GCons(__e, let__11410, tmp11717)

tmp11719 := GCons(__e, tmp11718, Nil)

tmp11720 := Call2(__e, SymFn(symshen_4comb), let__11421, tmp11719)


ifres11713 = tmp11720


}

ifres11708 = ifres11713


}

ifres11704 = ifres11708


}

ifres11699 = ifres11704


}

ifres11694 = ifres11699


}

let__11408 := ifres11694
_ = let__11408

tmp11727 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11408)


if True == tmp11727 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11408
}


panic("unreachable")
})

tmp11728 := Call2(__e, ns2_1set, symshen_4_5single_6, tmp11692)


_ = tmp11728

tmp11729 := MakeNative1(func(__e *ControlFlow, V3464 Obj) Obj {
tmp11730 := Call1(__e, SymFn(symshen_4_5sides_6), V3464)


let__11423 := tmp11730
_ = let__11423

tmp11761 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11423)


var ifres11731 Obj

if True == tmp11761 {
tmp11732 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11731 = tmp11732


} else {
tmp11733 := Call1(__e, SymFn(symshen_4_5_1out), let__11423)


let__11424 := tmp11733
_ = let__11424

tmp11734 := Call1(__e, SymFn(symshen_4in_1_6), let__11423)


let__11425 := tmp11734
_ = let__11425

tmp11735 := Call1(__e, SymFn(symshen_4_5formulae_6), let__11425)


let__11426 := tmp11735
_ = let__11426

tmp11760 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11426)


var ifres11736 Obj

if True == tmp11760 {
tmp11737 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11736 = tmp11737


} else {
tmp11738 := Call1(__e, SymFn(symshen_4_5_1out), let__11426)


let__11427 := tmp11738
_ = let__11427

tmp11739 := Call1(__e, SymFn(symshen_4in_1_6), let__11426)


let__11428 := tmp11739
_ = let__11428

tmp11740 := Call1(__e, SymFn(symshen_4_5dbl_6), let__11428)


let__11429 := tmp11740
_ = let__11429

tmp11759 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11429)


var ifres11741 Obj

if True == tmp11759 {
tmp11742 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11741 = tmp11742


} else {
tmp11743 := Call1(__e, SymFn(symshen_4in_1_6), let__11429)


let__11430 := tmp11743
_ = let__11430

tmp11744 := Call1(__e, SymFn(symshen_4_5formula_6), let__11430)


let__11431 := tmp11744
_ = let__11431

tmp11758 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11431)


var ifres11745 Obj

if True == tmp11758 {
tmp11746 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11745 = tmp11746


} else {
tmp11747 := Call1(__e, SymFn(symshen_4_5_1out), let__11431)


let__11432 := tmp11747
_ = let__11432

tmp11748 := Call1(__e, SymFn(symshen_4in_1_6), let__11431)


let__11433 := tmp11748
_ = let__11433

tmp11749 := Call1(__e, SymFn(symshen_4_5sc_6), let__11433)


let__11434 := tmp11749
_ = let__11434

tmp11757 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11434)


var ifres11750 Obj

if True == tmp11757 {
tmp11751 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11750 = tmp11751


} else {
tmp11752 := Call1(__e, SymFn(symshen_4in_1_6), let__11434)


let__11435 := tmp11752
_ = let__11435

tmp11753 := GCons(__e, let__11432, Nil)

tmp11754 := GCons(__e, Nil, tmp11753)

tmp11755 := Call3(__e, SymFn(symshen_4lr_1rule), let__11424, let__11427, tmp11754)


tmp11756 := Call2(__e, SymFn(symshen_4comb), let__11435, tmp11755)


ifres11750 = tmp11756


}

ifres11745 = ifres11750


}

ifres11741 = ifres11745


}

ifres11736 = ifres11741


}

ifres11731 = ifres11736


}

let__11422 := ifres11731
_ = let__11422

tmp11763 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11422)


if True == tmp11763 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11422
}


panic("unreachable")
})

tmp11764 := Call2(__e, ns2_1set, symshen_4_5double_6, tmp11729)


_ = tmp11764

tmp11765 := MakeNative1(func(__e *ControlFlow, V3479 Obj) Obj {
tmp11766 := Call1(__e, SymFn(symshen_4_5formula_6), V3479)


let__11437 := tmp11766
_ = let__11437

tmp11786 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11437)


var ifres11767 Obj

if True == tmp11786 {
tmp11768 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11767 = tmp11768


} else {
tmp11769 := Call1(__e, SymFn(symshen_4_5_1out), let__11437)


let__11438 := tmp11769
_ = let__11438

tmp11770 := Call1(__e, SymFn(symshen_4in_1_6), let__11437)


let__11439 := tmp11770
_ = let__11439

tmp11771 := Call1(__e, SymFn(symshen_4_5sc_6), let__11439)


let__11440 := tmp11771
_ = let__11440

tmp11785 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11440)


var ifres11772 Obj

if True == tmp11785 {
tmp11773 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11772 = tmp11773


} else {
tmp11774 := Call1(__e, SymFn(symshen_4in_1_6), let__11440)


let__11441 := tmp11774
_ = let__11441

tmp11775 := Call1(__e, SymFn(symshen_4_5formulae_6), let__11441)


let__11442 := tmp11775
_ = let__11442

tmp11784 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11442)


var ifres11776 Obj

if True == tmp11784 {
tmp11777 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11776 = tmp11777


} else {
tmp11778 := Call1(__e, SymFn(symshen_4_5_1out), let__11442)


let__11443 := tmp11778
_ = let__11443

tmp11779 := Call1(__e, SymFn(symshen_4in_1_6), let__11442)


let__11444 := tmp11779
_ = let__11444

tmp11780 := GCons(__e, let__11438, Nil)

tmp11781 := GCons(__e, Nil, tmp11780)

tmp11782 := GCons(__e, tmp11781, let__11443)

tmp11783 := Call2(__e, SymFn(symshen_4comb), let__11444, tmp11782)


ifres11776 = tmp11783


}

ifres11772 = ifres11776


}

ifres11767 = ifres11772


}

let__11436 := ifres11767
_ = let__11436

tmp11805 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11436)


if True == tmp11805 {
tmp11787 := Call1(__e, SymFn(symshen_4_5formula_6), V3479)


let__11446 := tmp11787
_ = let__11446

tmp11801 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11446)


var ifres11788 Obj

if True == tmp11801 {
tmp11789 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11788 = tmp11789


} else {
tmp11790 := Call1(__e, SymFn(symshen_4_5_1out), let__11446)


let__11447 := tmp11790
_ = let__11447

tmp11791 := Call1(__e, SymFn(symshen_4in_1_6), let__11446)


let__11448 := tmp11791
_ = let__11448

tmp11792 := Call1(__e, SymFn(symshen_4_5sc_6), let__11448)


let__11449 := tmp11792
_ = let__11449

tmp11800 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11449)


var ifres11793 Obj

if True == tmp11800 {
tmp11794 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11793 = tmp11794


} else {
tmp11795 := Call1(__e, SymFn(symshen_4in_1_6), let__11449)


let__11450 := tmp11795
_ = let__11450

tmp11796 := GCons(__e, let__11447, Nil)

tmp11797 := GCons(__e, Nil, tmp11796)

tmp11798 := GCons(__e, tmp11797, Nil)

tmp11799 := Call2(__e, SymFn(symshen_4comb), let__11450, tmp11798)


ifres11793 = tmp11799


}

ifres11788 = ifres11793


}

let__11445 := ifres11788
_ = let__11445

tmp11803 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11445)


if True == tmp11803 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11445
}


} else {
return let__11436
}


panic("unreachable")
})

tmp11806 := Call2(__e, ns2_1set, symshen_4_5formulae_6, tmp11765)


_ = tmp11806

tmp11807 := MakeNative1(func(__e *ControlFlow, V3495 Obj) Obj {
tmp11808 := Call1(__e, SymFn(symshen_4_5ass_6), V3495)


let__11452 := tmp11808
_ = let__11452

tmp11826 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11452)


var ifres11809 Obj

if True == tmp11826 {
tmp11810 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11809 = tmp11810


} else {
tmp11811 := Call1(__e, SymFn(symshen_4_5_1out), let__11452)


let__11453 := tmp11811
_ = let__11453

tmp11812 := Call1(__e, SymFn(symshen_4in_1_6), let__11452)


let__11454 := tmp11812
_ = let__11454

tmp11825 := Call2(__e, SymFn(symshen_4hds_a_2), let__11454, sym_6_6)


var ifres11813 Obj

if True == tmp11825 {
tmp11814 := Call1(__e, SymFn(symtail), let__11454)


let__11455 := tmp11814
_ = let__11455

tmp11815 := Call1(__e, SymFn(symshen_4_5formula_6), let__11455)


let__11456 := tmp11815
_ = let__11456

tmp11823 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11456)


var ifres11816 Obj

if True == tmp11823 {
tmp11817 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11816 = tmp11817


} else {
tmp11818 := Call1(__e, SymFn(symshen_4_5_1out), let__11456)


let__11457 := tmp11818
_ = let__11457

tmp11819 := Call1(__e, SymFn(symshen_4in_1_6), let__11456)


let__11458 := tmp11819
_ = let__11458

tmp11820 := GCons(__e, let__11457, Nil)

tmp11821 := GCons(__e, let__11453, tmp11820)

tmp11822 := Call2(__e, SymFn(symshen_4comb), let__11458, tmp11821)


ifres11816 = tmp11822


}

ifres11813 = ifres11816


} else {
tmp11824 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11813 = tmp11824


}

ifres11809 = ifres11813


}

let__11451 := ifres11809
_ = let__11451

tmp11839 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11451)


if True == tmp11839 {
tmp11827 := Call1(__e, SymFn(symshen_4_5formula_6), V3495)


let__11460 := tmp11827
_ = let__11460

tmp11835 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11460)


var ifres11828 Obj

if True == tmp11835 {
tmp11829 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11828 = tmp11829


} else {
tmp11830 := Call1(__e, SymFn(symshen_4_5_1out), let__11460)


let__11461 := tmp11830
_ = let__11461

tmp11831 := Call1(__e, SymFn(symshen_4in_1_6), let__11460)


let__11462 := tmp11831
_ = let__11462

tmp11832 := GCons(__e, let__11461, Nil)

tmp11833 := GCons(__e, Nil, tmp11832)

tmp11834 := Call2(__e, SymFn(symshen_4comb), let__11462, tmp11833)


ifres11828 = tmp11834


}

let__11459 := ifres11828
_ = let__11459

tmp11837 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11459)


if True == tmp11837 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11459
}


} else {
return let__11451
}


panic("unreachable")
})

tmp11840 := Call2(__e, ns2_1set, symshen_4_5conc_6, tmp11807)


_ = tmp11840

tmp11841 := MakeNative1(func(__e *ControlFlow, V3508 Obj) Obj {
tmp11842 := Call1(__e, SymFn(symshen_4_5prem_6), V3508)


let__11464 := tmp11842
_ = let__11464

tmp11860 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11464)


var ifres11843 Obj

if True == tmp11860 {
tmp11844 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11843 = tmp11844


} else {
tmp11845 := Call1(__e, SymFn(symshen_4_5_1out), let__11464)


let__11465 := tmp11845
_ = let__11465

tmp11846 := Call1(__e, SymFn(symshen_4in_1_6), let__11464)


let__11466 := tmp11846
_ = let__11466

tmp11847 := Call1(__e, SymFn(symshen_4_5sc_6), let__11466)


let__11467 := tmp11847
_ = let__11467

tmp11859 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11467)


var ifres11848 Obj

if True == tmp11859 {
tmp11849 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11848 = tmp11849


} else {
tmp11850 := Call1(__e, SymFn(symshen_4in_1_6), let__11467)


let__11468 := tmp11850
_ = let__11468

tmp11851 := Call1(__e, SymFn(symshen_4_5prems_6), let__11468)


let__11469 := tmp11851
_ = let__11469

tmp11858 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11469)


var ifres11852 Obj

if True == tmp11858 {
tmp11853 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11852 = tmp11853


} else {
tmp11854 := Call1(__e, SymFn(symshen_4_5_1out), let__11469)


let__11470 := tmp11854
_ = let__11470

tmp11855 := Call1(__e, SymFn(symshen_4in_1_6), let__11469)


let__11471 := tmp11855
_ = let__11471

tmp11856 := GCons(__e, let__11465, let__11470)

tmp11857 := Call2(__e, SymFn(symshen_4comb), let__11471, tmp11856)


ifres11852 = tmp11857


}

ifres11848 = ifres11852


}

ifres11843 = ifres11848


}

let__11463 := ifres11843
_ = let__11463

tmp11870 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11463)


if True == tmp11870 {
tmp11861 := Call1(__e, SymFn(sym_5e_6), V3508)


let__11473 := tmp11861
_ = let__11473

tmp11866 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11473)


var ifres11862 Obj

if True == tmp11866 {
tmp11863 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11862 = tmp11863


} else {
tmp11864 := Call1(__e, SymFn(symshen_4in_1_6), let__11473)


let__11474 := tmp11864
_ = let__11474

tmp11865 := Call2(__e, SymFn(symshen_4comb), let__11474, Nil)


ifres11862 = tmp11865


}

let__11472 := ifres11862
_ = let__11472

tmp11868 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11472)


if True == tmp11868 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11472
}


} else {
return let__11463
}


panic("unreachable")
})

tmp11871 := Call2(__e, ns2_1set, symshen_4_5prems_6, tmp11841)


_ = tmp11871

tmp11872 := MakeNative1(func(__e *ControlFlow, V3521 Obj) Obj {
tmp11877 := Call2(__e, SymFn(symshen_4hds_a_2), V3521, sym_b)


var ifres11873 Obj

if True == tmp11877 {
tmp11874 := Call1(__e, SymFn(symtail), V3521)


let__11476 := tmp11874
_ = let__11476

tmp11875 := Call2(__e, SymFn(symshen_4comb), let__11476, sym_b)


ifres11873 = tmp11875


} else {
tmp11876 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11873 = tmp11876


}

let__11475 := ifres11873
_ = let__11475

tmp11911 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11475)


if True == tmp11911 {
tmp11878 := Call1(__e, SymFn(symshen_4_5ass_6), V3521)


let__11478 := tmp11878
_ = let__11478

tmp11896 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11478)


var ifres11879 Obj

if True == tmp11896 {
tmp11880 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11879 = tmp11880


} else {
tmp11881 := Call1(__e, SymFn(symshen_4_5_1out), let__11478)


let__11479 := tmp11881
_ = let__11479

tmp11882 := Call1(__e, SymFn(symshen_4in_1_6), let__11478)


let__11480 := tmp11882
_ = let__11480

tmp11895 := Call2(__e, SymFn(symshen_4hds_a_2), let__11480, sym_6_6)


var ifres11883 Obj

if True == tmp11895 {
tmp11884 := Call1(__e, SymFn(symtail), let__11480)


let__11481 := tmp11884
_ = let__11481

tmp11885 := Call1(__e, SymFn(symshen_4_5formula_6), let__11481)


let__11482 := tmp11885
_ = let__11482

tmp11893 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11482)


var ifres11886 Obj

if True == tmp11893 {
tmp11887 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11886 = tmp11887


} else {
tmp11888 := Call1(__e, SymFn(symshen_4_5_1out), let__11482)


let__11483 := tmp11888
_ = let__11483

tmp11889 := Call1(__e, SymFn(symshen_4in_1_6), let__11482)


let__11484 := tmp11889
_ = let__11484

tmp11890 := GCons(__e, let__11483, Nil)

tmp11891 := GCons(__e, let__11479, tmp11890)

tmp11892 := Call2(__e, SymFn(symshen_4comb), let__11484, tmp11891)


ifres11886 = tmp11892


}

ifres11883 = ifres11886


} else {
tmp11894 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11883 = tmp11894


}

ifres11879 = ifres11883


}

let__11477 := ifres11879
_ = let__11477

tmp11909 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11477)


if True == tmp11909 {
tmp11897 := Call1(__e, SymFn(symshen_4_5formula_6), V3521)


let__11486 := tmp11897
_ = let__11486

tmp11905 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11486)


var ifres11898 Obj

if True == tmp11905 {
tmp11899 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11898 = tmp11899


} else {
tmp11900 := Call1(__e, SymFn(symshen_4_5_1out), let__11486)


let__11487 := tmp11900
_ = let__11487

tmp11901 := Call1(__e, SymFn(symshen_4in_1_6), let__11486)


let__11488 := tmp11901
_ = let__11488

tmp11902 := GCons(__e, let__11487, Nil)

tmp11903 := GCons(__e, Nil, tmp11902)

tmp11904 := Call2(__e, SymFn(symshen_4comb), let__11488, tmp11903)


ifres11898 = tmp11904


}

let__11485 := ifres11898
_ = let__11485

tmp11907 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11485)


if True == tmp11907 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11485
}


} else {
return let__11477
}


} else {
return let__11475
}


panic("unreachable")
})

tmp11912 := Call2(__e, ns2_1set, symshen_4_5prem_6, tmp11872)


_ = tmp11912

tmp11913 := MakeNative1(func(__e *ControlFlow, V3536 Obj) Obj {
tmp11914 := Call1(__e, SymFn(symshen_4_5formula_6), V3536)


let__11490 := tmp11914
_ = let__11490

tmp11932 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11490)


var ifres11915 Obj

if True == tmp11932 {
tmp11916 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11915 = tmp11916


} else {
tmp11917 := Call1(__e, SymFn(symshen_4_5_1out), let__11490)


let__11491 := tmp11917
_ = let__11491

tmp11918 := Call1(__e, SymFn(symshen_4in_1_6), let__11490)


let__11492 := tmp11918
_ = let__11492

tmp11919 := Call1(__e, SymFn(symshen_4_5iscomma_6), let__11492)


let__11493 := tmp11919
_ = let__11493

tmp11931 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11493)


var ifres11920 Obj

if True == tmp11931 {
tmp11921 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11920 = tmp11921


} else {
tmp11922 := Call1(__e, SymFn(symshen_4in_1_6), let__11493)


let__11494 := tmp11922
_ = let__11494

tmp11923 := Call1(__e, SymFn(symshen_4_5ass_6), let__11494)


let__11495 := tmp11923
_ = let__11495

tmp11930 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11495)


var ifres11924 Obj

if True == tmp11930 {
tmp11925 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11924 = tmp11925


} else {
tmp11926 := Call1(__e, SymFn(symshen_4_5_1out), let__11495)


let__11496 := tmp11926
_ = let__11496

tmp11927 := Call1(__e, SymFn(symshen_4in_1_6), let__11495)


let__11497 := tmp11927
_ = let__11497

tmp11928 := GCons(__e, let__11491, let__11496)

tmp11929 := Call2(__e, SymFn(symshen_4comb), let__11497, tmp11928)


ifres11924 = tmp11929


}

ifres11920 = ifres11924


}

ifres11915 = ifres11920


}

let__11489 := ifres11915
_ = let__11489

tmp11952 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11489)


if True == tmp11952 {
tmp11933 := Call1(__e, SymFn(symshen_4_5formula_6), V3536)


let__11499 := tmp11933
_ = let__11499

tmp11940 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11499)


var ifres11934 Obj

if True == tmp11940 {
tmp11935 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11934 = tmp11935


} else {
tmp11936 := Call1(__e, SymFn(symshen_4_5_1out), let__11499)


let__11500 := tmp11936
_ = let__11500

tmp11937 := Call1(__e, SymFn(symshen_4in_1_6), let__11499)


let__11501 := tmp11937
_ = let__11501

tmp11938 := GCons(__e, let__11500, Nil)

tmp11939 := Call2(__e, SymFn(symshen_4comb), let__11501, tmp11938)


ifres11934 = tmp11939


}

let__11498 := ifres11934
_ = let__11498

tmp11950 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11498)


if True == tmp11950 {
tmp11941 := Call1(__e, SymFn(sym_5e_6), V3536)


let__11503 := tmp11941
_ = let__11503

tmp11946 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11503)


var ifres11942 Obj

if True == tmp11946 {
tmp11943 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11942 = tmp11943


} else {
tmp11944 := Call1(__e, SymFn(symshen_4in_1_6), let__11503)


let__11504 := tmp11944
_ = let__11504

tmp11945 := Call2(__e, SymFn(symshen_4comb), let__11504, Nil)


ifres11942 = tmp11945


}

let__11502 := ifres11942
_ = let__11502

tmp11948 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11502)


if True == tmp11948 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11502
}


} else {
return let__11498
}


} else {
return let__11489
}


panic("unreachable")
})

tmp11953 := Call2(__e, ns2_1set, symshen_4_5ass_6, tmp11913)


_ = tmp11953

tmp11954 := MakeNative1(func(__e *ControlFlow, V3553 Obj) Obj {
tmp11964 := GConsP(__e, V3553)

var ifres11955 Obj

if True == tmp11964 {
tmp11956 := Call1(__e, SymFn(symhead), V3553)


let__11506 := tmp11956
_ = let__11506

tmp11957 := Call1(__e, SymFn(symtail), V3553)


let__11507 := tmp11957
_ = let__11507

tmp11961 := PrimGuard1(__e, symintern, PrimIntern, MakeString(","))

tmp11962 := GEq(__e, let__11506, tmp11961)

var ifres11958 Obj

if True == tmp11962 {
tmp11959 := Call2(__e, SymFn(symshen_4comb), let__11507, symshen_4skip)


ifres11958 = tmp11959


} else {
tmp11960 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11958 = tmp11960


}

ifres11955 = ifres11958


} else {
tmp11963 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11955 = tmp11963


}

let__11505 := ifres11955
_ = let__11505

tmp11966 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11505)


if True == tmp11966 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11505
}


panic("unreachable")
})

tmp11967 := Call2(__e, ns2_1set, symshen_4_5iscomma_6, tmp11954)


_ = tmp11967

tmp11968 := MakeNative1(func(__e *ControlFlow, V3557 Obj) Obj {
tmp11969 := Call1(__e, SymFn(symshen_4_5expr_6), V3557)


let__11509 := tmp11969
_ = let__11509

tmp11992 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11509)


var ifres11970 Obj

if True == tmp11992 {
tmp11971 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11970 = tmp11971


} else {
tmp11972 := Call1(__e, SymFn(symshen_4_5_1out), let__11509)


let__11510 := tmp11972
_ = let__11510

tmp11973 := Call1(__e, SymFn(symshen_4in_1_6), let__11509)


let__11511 := tmp11973
_ = let__11511

tmp11974 := Call1(__e, SymFn(symshen_4_5iscolon_6), let__11511)


let__11512 := tmp11974
_ = let__11512

tmp11991 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11512)


var ifres11975 Obj

if True == tmp11991 {
tmp11976 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11975 = tmp11976


} else {
tmp11977 := Call1(__e, SymFn(symshen_4in_1_6), let__11512)


let__11513 := tmp11977
_ = let__11513

tmp11978 := Call1(__e, SymFn(symshen_4_5type_6), let__11513)


let__11514 := tmp11978
_ = let__11514

tmp11990 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11514)


var ifres11979 Obj

if True == tmp11990 {
tmp11980 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11979 = tmp11980


} else {
tmp11981 := Call1(__e, SymFn(symshen_4_5_1out), let__11514)


let__11515 := tmp11981
_ = let__11515

tmp11982 := Call1(__e, SymFn(symshen_4in_1_6), let__11514)


let__11516 := tmp11982
_ = let__11516

tmp11983 := Call1(__e, SymFn(symshen_4curry), let__11510)


tmp11984 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp11985 := Call1(__e, SymFn(symshen_4rectify_1type), let__11515)


tmp11986 := GCons(__e, tmp11985, Nil)

tmp11987 := GCons(__e, tmp11984, tmp11986)

tmp11988 := GCons(__e, tmp11983, tmp11987)

tmp11989 := Call2(__e, SymFn(symshen_4comb), let__11516, tmp11988)


ifres11979 = tmp11989


}

ifres11975 = ifres11979


}

ifres11970 = ifres11975


}

let__11508 := ifres11970
_ = let__11508

tmp12003 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11508)


if True == tmp12003 {
tmp11993 := Call1(__e, SymFn(symshen_4_5expr_6), V3557)


let__11518 := tmp11993
_ = let__11518

tmp11999 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11518)


var ifres11994 Obj

if True == tmp11999 {
tmp11995 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11994 = tmp11995


} else {
tmp11996 := Call1(__e, SymFn(symshen_4_5_1out), let__11518)


let__11519 := tmp11996
_ = let__11519

tmp11997 := Call1(__e, SymFn(symshen_4in_1_6), let__11518)


let__11520 := tmp11997
_ = let__11520

tmp11998 := Call2(__e, SymFn(symshen_4comb), let__11520, let__11519)


ifres11994 = tmp11998


}

let__11517 := ifres11994
_ = let__11517

tmp12001 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11517)


if True == tmp12001 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11517
}


} else {
return let__11508
}


panic("unreachable")
})

tmp12004 := Call2(__e, ns2_1set, symshen_4_5formula_6, tmp11968)


_ = tmp12004

tmp12005 := MakeNative1(func(__e *ControlFlow, V3571 Obj) Obj {
tmp12015 := GConsP(__e, V3571)

var ifres12006 Obj

if True == tmp12015 {
tmp12007 := Call1(__e, SymFn(symhead), V3571)


let__11522 := tmp12007
_ = let__11522

tmp12008 := Call1(__e, SymFn(symtail), V3571)


let__11523 := tmp12008
_ = let__11523

tmp12012 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp12013 := GEq(__e, let__11522, tmp12012)

var ifres12009 Obj

if True == tmp12013 {
tmp12010 := Call2(__e, SymFn(symshen_4comb), let__11523, symshen_4skip)


ifres12009 = tmp12010


} else {
tmp12011 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12009 = tmp12011


}

ifres12006 = ifres12009


} else {
tmp12014 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12006 = tmp12014


}

let__11521 := ifres12006
_ = let__11521

tmp12017 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11521)


if True == tmp12017 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11521
}


panic("unreachable")
})

tmp12018 := Call2(__e, ns2_1set, symshen_4_5iscolon_6, tmp12005)


_ = tmp12018

tmp12019 := MakeNative1(func(__e *ControlFlow, V3575 Obj) Obj {
tmp12020 := Call1(__e, SymFn(symshen_4_5side_6), V3575)


let__11525 := tmp12020
_ = let__11525

tmp12033 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11525)


var ifres12021 Obj

if True == tmp12033 {
tmp12022 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12021 = tmp12022


} else {
tmp12023 := Call1(__e, SymFn(symshen_4_5_1out), let__11525)


let__11526 := tmp12023
_ = let__11526

tmp12024 := Call1(__e, SymFn(symshen_4in_1_6), let__11525)


let__11527 := tmp12024
_ = let__11527

tmp12025 := Call1(__e, SymFn(symshen_4_5sides_6), let__11527)


let__11528 := tmp12025
_ = let__11528

tmp12032 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11528)


var ifres12026 Obj

if True == tmp12032 {
tmp12027 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12026 = tmp12027


} else {
tmp12028 := Call1(__e, SymFn(symshen_4_5_1out), let__11528)


let__11529 := tmp12028
_ = let__11529

tmp12029 := Call1(__e, SymFn(symshen_4in_1_6), let__11528)


let__11530 := tmp12029
_ = let__11530

tmp12030 := GCons(__e, let__11526, let__11529)

tmp12031 := Call2(__e, SymFn(symshen_4comb), let__11530, tmp12030)


ifres12026 = tmp12031


}

ifres12021 = ifres12026


}

let__11524 := ifres12021
_ = let__11524

tmp12043 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11524)


if True == tmp12043 {
tmp12034 := Call1(__e, SymFn(sym_5e_6), V3575)


let__11532 := tmp12034
_ = let__11532

tmp12039 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11532)


var ifres12035 Obj

if True == tmp12039 {
tmp12036 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12035 = tmp12036


} else {
tmp12037 := Call1(__e, SymFn(symshen_4in_1_6), let__11532)


let__11533 := tmp12037
_ = let__11533

tmp12038 := Call2(__e, SymFn(symshen_4comb), let__11533, Nil)


ifres12035 = tmp12038


}

let__11531 := ifres12035
_ = let__11531

tmp12041 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11531)


if True == tmp12041 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11531
}


} else {
return let__11524
}


panic("unreachable")
})

tmp12044 := Call2(__e, ns2_1set, symshen_4_5sides_6, tmp12019)


_ = tmp12044

tmp12045 := MakeNative1(func(__e *ControlFlow, V3586 Obj) Obj {
tmp12057 := Call2(__e, SymFn(symshen_4hds_a_2), V3586, symif)


var ifres12046 Obj

if True == tmp12057 {
tmp12047 := Call1(__e, SymFn(symtail), V3586)


let__11535 := tmp12047
_ = let__11535

tmp12055 := GConsP(__e, let__11535)

var ifres12048 Obj

if True == tmp12055 {
tmp12049 := Call1(__e, SymFn(symhead), let__11535)


let__11536 := tmp12049
_ = let__11536

tmp12050 := Call1(__e, SymFn(symtail), let__11535)


let__11537 := tmp12050
_ = let__11537

tmp12051 := GCons(__e, let__11536, Nil)

tmp12052 := GCons(__e, symif, tmp12051)

tmp12053 := Call2(__e, SymFn(symshen_4comb), let__11537, tmp12052)


ifres12048 = tmp12053


} else {
tmp12054 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12048 = tmp12054


}

ifres12046 = ifres12048


} else {
tmp12056 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12046 = tmp12056


}

let__11534 := ifres12046
_ = let__11534

tmp12113 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11534)


if True == tmp12113 {
tmp12075 := Call2(__e, SymFn(symshen_4hds_a_2), V3586, symlet)


var ifres12058 Obj

if True == tmp12075 {
tmp12059 := Call1(__e, SymFn(symtail), V3586)


let__11539 := tmp12059
_ = let__11539

tmp12073 := GConsP(__e, let__11539)

var ifres12060 Obj

if True == tmp12073 {
tmp12061 := Call1(__e, SymFn(symhead), let__11539)


let__11540 := tmp12061
_ = let__11540

tmp12062 := Call1(__e, SymFn(symtail), let__11539)


let__11541 := tmp12062
_ = let__11541

tmp12071 := GConsP(__e, let__11541)

var ifres12063 Obj

if True == tmp12071 {
tmp12064 := Call1(__e, SymFn(symhead), let__11541)


let__11542 := tmp12064
_ = let__11542

tmp12065 := Call1(__e, SymFn(symtail), let__11541)


let__11543 := tmp12065
_ = let__11543

tmp12066 := GCons(__e, let__11542, Nil)

tmp12067 := GCons(__e, let__11540, tmp12066)

tmp12068 := GCons(__e, symlet, tmp12067)

tmp12069 := Call2(__e, SymFn(symshen_4comb), let__11543, tmp12068)


ifres12063 = tmp12069


} else {
tmp12070 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12063 = tmp12070


}

ifres12060 = ifres12063


} else {
tmp12072 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12060 = tmp12072


}

ifres12058 = ifres12060


} else {
tmp12074 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12058 = tmp12074


}

let__11538 := ifres12058
_ = let__11538

tmp12111 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11538)


if True == tmp12111 {
tmp12090 := Call2(__e, SymFn(symshen_4hds_a_2), V3586, symctxt)


var ifres12076 Obj

if True == tmp12090 {
tmp12077 := Call1(__e, SymFn(symtail), V3586)


let__11545 := tmp12077
_ = let__11545

tmp12088 := GConsP(__e, let__11545)

var ifres12078 Obj

if True == tmp12088 {
tmp12079 := Call1(__e, SymFn(symhead), let__11545)


let__11546 := tmp12079
_ = let__11546

tmp12080 := Call1(__e, SymFn(symtail), let__11545)


let__11547 := tmp12080
_ = let__11547

tmp12086 := PrimGuard1(__e, symvariable_2, PrimIsVariable, let__11546)

var ifres12081 Obj

if True == tmp12086 {
tmp12082 := GCons(__e, let__11546, Nil)

tmp12083 := GCons(__e, symctxt, tmp12082)

tmp12084 := Call2(__e, SymFn(symshen_4comb), let__11547, tmp12083)


ifres12081 = tmp12084


} else {
tmp12085 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12081 = tmp12085


}

ifres12078 = ifres12081


} else {
tmp12087 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12078 = tmp12087


}

ifres12076 = ifres12078


} else {
tmp12089 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12076 = tmp12089


}

let__11544 := ifres12076
_ = let__11544

tmp12109 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11544)


if True == tmp12109 {
tmp12105 := Call2(__e, SymFn(symshen_4hds_a_2), V3586, symsqts)


var ifres12091 Obj

if True == tmp12105 {
tmp12092 := Call1(__e, SymFn(symtail), V3586)


let__11549 := tmp12092
_ = let__11549

tmp12103 := GConsP(__e, let__11549)

var ifres12093 Obj

if True == tmp12103 {
tmp12094 := Call1(__e, SymFn(symhead), let__11549)


let__11550 := tmp12094
_ = let__11550

tmp12095 := Call1(__e, SymFn(symtail), let__11549)


let__11551 := tmp12095
_ = let__11551

tmp12101 := PrimGuard1(__e, symvariable_2, PrimIsVariable, let__11550)

var ifres12096 Obj

if True == tmp12101 {
tmp12097 := GCons(__e, let__11550, Nil)

tmp12098 := GCons(__e, symsqts, tmp12097)

tmp12099 := Call2(__e, SymFn(symshen_4comb), let__11551, tmp12098)


ifres12096 = tmp12099


} else {
tmp12100 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12096 = tmp12100


}

ifres12093 = ifres12096


} else {
tmp12102 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12093 = tmp12102


}

ifres12091 = ifres12093


} else {
tmp12104 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12091 = tmp12104


}

let__11548 := ifres12091
_ = let__11548

tmp12107 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11548)


if True == tmp12107 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11548
}


} else {
return let__11544
}


} else {
return let__11538
}


} else {
return let__11534
}


panic("unreachable")
})

tmp12114 := Call2(__e, ns2_1set, symshen_4_5side_6, tmp12045)


_ = tmp12114

tmp12115 := MakeNative3(func(__e *ControlFlow, V3611, V3612, V3613 Obj) Obj {
tmp12142 := GConsP(__e, V3613)

var ifres12132 Obj

if True == tmp12142 {
tmp12140 := GHd(__e, V3613)

tmp12141 := GEq(__e, Nil, tmp12140)

var ifres12133 Obj

if True == tmp12141 {
tmp12138 := GTl(__e, V3613)

tmp12139 := GConsP(__e, tmp12138)

var ifres12134 Obj

if True == tmp12139 {
tmp12135 := GTl(__e, V3613)

tmp12136 := GTl(__e, tmp12135)

tmp12137 := GEq(__e, Nil, tmp12136)

ifres12134 = tmp12137


} else {
ifres12134 = False


}

ifres12133 = ifres12134


} else {
ifres12133 = False


}

ifres12132 = ifres12133


} else {
ifres12132 = False


}

if True == ifres12132 {
tmp12116 := Call1(__e, SymFn(symgensym), symP)


let__11552 := tmp12116
_ = let__11552

tmp12117 := GTl(__e, V3613)

tmp12118 := GCons(__e, let__11552, Nil)

tmp12119 := GCons(__e, tmp12117, tmp12118)

let__11553 := tmp12119
_ = let__11553

tmp12120 := Call1(__e, SymFn(symshen_4coll_1formulae), V3612)


tmp12121 := GCons(__e, let__11552, Nil)

tmp12122 := GCons(__e, tmp12120, tmp12121)

let__11554 := tmp12122
_ = let__11554

tmp12123 := GCons(__e, let__11554, Nil)

tmp12124 := GCons(__e, let__11553, Nil)

tmp12125 := GCons(__e, tmp12123, tmp12124)

tmp12126 := GCons(__e, V3611, tmp12125)

let__11555 := tmp12126
_ = let__11555

tmp12127 := GCons(__e, V3613, Nil)

tmp12128 := GCons(__e, V3612, tmp12127)

tmp12129 := GCons(__e, V3611, tmp12128)

let__11556 := tmp12129
_ = let__11556

tmp12130 := GCons(__e, let__11555, Nil)

return GCons(__e, let__11556, tmp12130)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.lr-rule"))
}


panic("unreachable")
})

tmp12143 := Call2(__e, ns2_1set, symshen_4lr_1rule, tmp12115)


_ = tmp12143

tmp12144 := MakeNative1(func(__e *ControlFlow, V3621 Obj) Obj {
tmp12169 := GEq(__e, Nil, V3621)

if True == tmp12169 {
return Nil
} else {
tmp12167 := GConsP(__e, V3621)

var ifres12151 Obj

if True == tmp12167 {
tmp12165 := GHd(__e, V3621)

tmp12166 := GConsP(__e, tmp12165)

var ifres12152 Obj

if True == tmp12166 {
tmp12162 := GHd(__e, V3621)

tmp12163 := GHd(__e, tmp12162)

tmp12164 := GEq(__e, Nil, tmp12163)

var ifres12153 Obj

if True == tmp12164 {
tmp12159 := GHd(__e, V3621)

tmp12160 := GTl(__e, tmp12159)

tmp12161 := GConsP(__e, tmp12160)

var ifres12154 Obj

if True == tmp12161 {
tmp12155 := GHd(__e, V3621)

tmp12156 := GTl(__e, tmp12155)

tmp12157 := GTl(__e, tmp12156)

tmp12158 := GEq(__e, Nil, tmp12157)

ifres12154 = tmp12158


} else {
ifres12154 = False


}

ifres12153 = ifres12154


} else {
ifres12153 = False


}

ifres12152 = ifres12153


} else {
ifres12152 = False


}

ifres12151 = ifres12152


} else {
ifres12151 = False


}

if True == ifres12151 {
tmp12145 := GHd(__e, V3621)

tmp12146 := GTl(__e, tmp12145)

tmp12147 := GHd(__e, tmp12146)

tmp12148 := GTl(__e, V3621)

tmp12149 := Call1(__e, SymFn(symshen_4coll_1formulae), tmp12148)


return GCons(__e, tmp12147, tmp12149)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.coll-formulae"))
}


}


panic("unreachable")
})

tmp12170 := Call2(__e, ns2_1set, symshen_4coll_1formulae, tmp12144)


_ = tmp12170

tmp12171 := MakeNative1(func(__e *ControlFlow, V3622 Obj) Obj {
tmp12182 := GConsP(__e, V3622)

var ifres12172 Obj

if True == tmp12182 {
tmp12173 := Call1(__e, SymFn(symhead), V3622)


let__11558 := tmp12173
_ = let__11558

tmp12174 := Call1(__e, SymFn(symtail), V3622)


let__11559 := tmp12174
_ = let__11559

tmp12179 := Call1(__e, SymFn(symshen_4key_1in_1sequent_1calculus_2), let__11558)


tmp12180 := GNot(__e, tmp12179)

var ifres12175 Obj

if True == tmp12180 {
tmp12176 := Call1(__e, SymFn(symmacroexpand), let__11558)


tmp12177 := Call2(__e, SymFn(symshen_4comb), let__11559, tmp12176)


ifres12175 = tmp12177


} else {
tmp12178 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12175 = tmp12178


}

ifres12172 = ifres12175


} else {
tmp12181 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12172 = tmp12181


}

let__11557 := ifres12172
_ = let__11557

tmp12184 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11557)


if True == tmp12184 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11557
}


panic("unreachable")
})

tmp12185 := Call2(__e, ns2_1set, symshen_4_5expr_6, tmp12171)


_ = tmp12185

tmp12186 := MakeNative1(func(__e *ControlFlow, V3626 Obj) Obj {
tmp12190 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp12191 := PrimGuard1(__e, symintern, PrimIntern, MakeString(","))

tmp12192 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp12193 := GCons(__e, sym_5_1_1, Nil)

tmp12194 := GCons(__e, tmp12192, tmp12193)

tmp12195 := GCons(__e, tmp12191, tmp12194)

tmp12196 := GCons(__e, tmp12190, tmp12195)

tmp12197 := GCons(__e, sym_6_6, tmp12196)

tmp12198 := Call2(__e, SymFn(symelement_2), V3626, tmp12197)


if True == tmp12198 {
return True
} else {
tmp12188 := Call1(__e, SymFn(symshen_4sng_2), V3626)


if True == tmp12188 {
return True
} else {
return TailCall1(__e, SymFn(symshen_4dbl_2), V3626)
}


}


panic("unreachable")
})

tmp12199 := Call2(__e, ns2_1set, symshen_4key_1in_1sequent_1calculus_2, tmp12186)


_ = tmp12199

tmp12200 := MakeNative1(func(__e *ControlFlow, V3627 Obj) Obj {
tmp12201 := Call1(__e, SymFn(symshen_4_5expr_6), V3627)


let__11561 := tmp12201
_ = let__11561

tmp12207 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11561)


var ifres12202 Obj

if True == tmp12207 {
tmp12203 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12202 = tmp12203


} else {
tmp12204 := Call1(__e, SymFn(symshen_4_5_1out), let__11561)


let__11562 := tmp12204
_ = let__11562

tmp12205 := Call1(__e, SymFn(symshen_4in_1_6), let__11561)


let__11563 := tmp12205
_ = let__11563

tmp12206 := Call2(__e, SymFn(symshen_4comb), let__11563, let__11562)


ifres12202 = tmp12206


}

let__11560 := ifres12202
_ = let__11560

tmp12209 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11560)


if True == tmp12209 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11560
}


panic("unreachable")
})

tmp12210 := Call2(__e, ns2_1set, symshen_4_5type_6, tmp12200)


_ = tmp12210

tmp12211 := MakeNative1(func(__e *ControlFlow, V3632 Obj) Obj {
tmp12220 := GConsP(__e, V3632)

var ifres12212 Obj

if True == tmp12220 {
tmp12213 := Call1(__e, SymFn(symhead), V3632)


let__11565 := tmp12213
_ = let__11565

tmp12214 := Call1(__e, SymFn(symtail), V3632)


let__11566 := tmp12214
_ = let__11566

tmp12218 := Call1(__e, SymFn(symshen_4dbl_2), let__11565)


var ifres12215 Obj

if True == tmp12218 {
tmp12216 := Call2(__e, SymFn(symshen_4comb), let__11566, let__11565)


ifres12215 = tmp12216


} else {
tmp12217 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12215 = tmp12217


}

ifres12212 = ifres12215


} else {
tmp12219 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12212 = tmp12219


}

let__11564 := ifres12212
_ = let__11564

tmp12222 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11564)


if True == tmp12222 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11564
}


panic("unreachable")
})

tmp12223 := Call2(__e, ns2_1set, symshen_4_5dbl_6, tmp12211)


_ = tmp12223

tmp12224 := MakeNative1(func(__e *ControlFlow, V3636 Obj) Obj {
tmp12233 := GConsP(__e, V3636)

var ifres12225 Obj

if True == tmp12233 {
tmp12226 := Call1(__e, SymFn(symhead), V3636)


let__11568 := tmp12226
_ = let__11568

tmp12227 := Call1(__e, SymFn(symtail), V3636)


let__11569 := tmp12227
_ = let__11569

tmp12231 := Call1(__e, SymFn(symshen_4sng_2), let__11568)


var ifres12228 Obj

if True == tmp12231 {
tmp12229 := Call2(__e, SymFn(symshen_4comb), let__11569, let__11568)


ifres12228 = tmp12229


} else {
tmp12230 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12228 = tmp12230


}

ifres12225 = ifres12228


} else {
tmp12232 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12225 = tmp12232


}

let__11567 := ifres12225
_ = let__11567

tmp12235 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11567)


if True == tmp12235 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11567
}


panic("unreachable")
})

tmp12236 := Call2(__e, ns2_1set, symshen_4_5sng_6, tmp12224)


_ = tmp12236

tmp12237 := MakeNative1(func(__e *ControlFlow, V3640 Obj) Obj {
tmp12240 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V3640)

if True == tmp12240 {
tmp12238 := PrimGuard1(__e, symstr, PrimStr, V3640)

return TailCall1(__e, SymFn(symshen_4sng_1h_2), tmp12238)


} else {
return False
}


panic("unreachable")
})

tmp12241 := Call2(__e, ns2_1set, symshen_4sng_2, tmp12237)


_ = tmp12241

tmp12242 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V3643 := __self1
_ = V3643
tmp12250 := GEq(__e, MakeString("___"), V3643)

if True == tmp12250 {
return True
} else {
tmp12248 := Call1(__e, SymFn(symshen_4_7string_2), V3643)


var ifres12245 Obj

if True == tmp12248 {
tmp12246 := Call1(__e, SymFn(symhdstr), V3643)


tmp12247 := GEq(__e, MakeString("_"), tmp12246)

ifres12245 = tmp12247


} else {
ifres12245 = False


}

if True == ifres12245 {
if SymFn(symshen_4sng_1h_2) == __self {
__self1 = PrimGuard1(__e, symtlstr, PrimTailString, V3643)
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4sng_1h_2), PrimGuard1(__e, symtlstr, PrimTailString, V3643))


} else {
return False
}


}


panic("unreachable")
})
return __self
})()

tmp12251 := Call2(__e, ns2_1set, symshen_4sng_1h_2, tmp12242)


_ = tmp12251

tmp12252 := MakeNative1(func(__e *ControlFlow, V3644 Obj) Obj {
tmp12255 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V3644)

if True == tmp12255 {
tmp12253 := PrimGuard1(__e, symstr, PrimStr, V3644)

return TailCall1(__e, SymFn(symshen_4dbl_1h_2), tmp12253)


} else {
return False
}


panic("unreachable")
})

tmp12256 := Call2(__e, ns2_1set, symshen_4dbl_2, tmp12252)


_ = tmp12256

tmp12257 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V3647 := __self1
_ = V3647
tmp12265 := GEq(__e, MakeString("==="), V3647)

if True == tmp12265 {
return True
} else {
tmp12263 := Call1(__e, SymFn(symshen_4_7string_2), V3647)


var ifres12260 Obj

if True == tmp12263 {
tmp12261 := Call1(__e, SymFn(symhdstr), V3647)


tmp12262 := GEq(__e, MakeString("="), tmp12261)

ifres12260 = tmp12262


} else {
ifres12260 = False


}

if True == ifres12260 {
if SymFn(symshen_4dbl_1h_2) == __self {
__self1 = PrimGuard1(__e, symtlstr, PrimTailString, V3647)
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4dbl_1h_2), PrimGuard1(__e, symtlstr, PrimTailString, V3647))


} else {
return False
}


}


panic("unreachable")
})
return __self
})()

tmp12266 := Call2(__e, ns2_1set, symshen_4dbl_1h_2, tmp12257)


_ = tmp12266

tmp12267 := MakeNative2(func(__e *ControlFlow, V3648, V3649 Obj) Obj {
tmp12268 := MakeNative1(func(__e *ControlFlow, Z3651 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4rule_1_6clause), Z3651)
panic("unreachable")
})

tmp12269 := Call2(__e, SymFn(symmapcan), tmp12268, V3649)


let__11570 := tmp12269
_ = let__11570

tmp12270 := GCons(__e, V3648, let__11570)

tmp12271 := GCons(__e, symdefprolog, tmp12270)

let__11571 := tmp12271
_ = let__11571

return TailCall1(__e, SymFn(symeval), let__11571)


panic("unreachable")
})

tmp12272 := Call2(__e, ns2_1set, symshen_4rules_1_6prolog, tmp12267)


_ = tmp12272

tmp12273 := MakeNative1(func(__e *ControlFlow, V3653 Obj) Obj {
tmp12327 := GConsP(__e, V3653)

var ifres12297 Obj

if True == tmp12327 {
tmp12325 := GTl(__e, V3653)

tmp12326 := GConsP(__e, tmp12325)

var ifres12298 Obj

if True == tmp12326 {
tmp12322 := GTl(__e, V3653)

tmp12323 := GTl(__e, tmp12322)

tmp12324 := GConsP(__e, tmp12323)

var ifres12299 Obj

if True == tmp12324 {
tmp12318 := GTl(__e, V3653)

tmp12319 := GTl(__e, tmp12318)

tmp12320 := GHd(__e, tmp12319)

tmp12321 := GConsP(__e, tmp12320)

var ifres12300 Obj

if True == tmp12321 {
tmp12313 := GTl(__e, V3653)

tmp12314 := GTl(__e, tmp12313)

tmp12315 := GHd(__e, tmp12314)

tmp12316 := GTl(__e, tmp12315)

tmp12317 := GConsP(__e, tmp12316)

var ifres12301 Obj

if True == tmp12317 {
tmp12307 := GTl(__e, V3653)

tmp12308 := GTl(__e, tmp12307)

tmp12309 := GHd(__e, tmp12308)

tmp12310 := GTl(__e, tmp12309)

tmp12311 := GTl(__e, tmp12310)

tmp12312 := GEq(__e, Nil, tmp12311)

var ifres12302 Obj

if True == tmp12312 {
tmp12303 := GTl(__e, V3653)

tmp12304 := GTl(__e, tmp12303)

tmp12305 := GTl(__e, tmp12304)

tmp12306 := GEq(__e, Nil, tmp12305)

ifres12302 = tmp12306


} else {
ifres12302 = False


}

ifres12301 = ifres12302


} else {
ifres12301 = False


}

ifres12300 = ifres12301


} else {
ifres12300 = False


}

ifres12299 = ifres12300


} else {
ifres12299 = False


}

ifres12298 = ifres12299


} else {
ifres12298 = False


}

ifres12297 = ifres12298


} else {
ifres12297 = False


}

if True == ifres12297 {
tmp12274 := GTl(__e, V3653)

tmp12275 := GTl(__e, tmp12274)

tmp12276 := GHd(__e, tmp12275)

tmp12277 := GTl(__e, tmp12276)

tmp12278 := GHd(__e, tmp12277)

tmp12279 := Call1(__e, SymFn(symshen_4extract_1vars), tmp12278)


let__11572 := tmp12279
_ = let__11572

tmp12280 := GTl(__e, V3653)

tmp12281 := GTl(__e, tmp12280)

tmp12282 := GHd(__e, tmp12281)

tmp12283 := GTl(__e, tmp12282)

tmp12284 := GHd(__e, tmp12283)

tmp12285 := Call1(__e, SymFn(symshen_4rule_1_6head), tmp12284)


tmp12286 := GCons(__e, sym_5_1_1, Nil)

tmp12287 := GHd(__e, V3653)

tmp12288 := GTl(__e, V3653)

tmp12289 := GHd(__e, tmp12288)

tmp12290 := GTl(__e, V3653)

tmp12291 := GTl(__e, tmp12290)

tmp12292 := GHd(__e, tmp12291)

tmp12293 := GHd(__e, tmp12292)

tmp12294 := Call5(__e, SymFn(symshen_4rule_1_6body), let__11572, symAssumptions, tmp12287, tmp12289, tmp12293)


tmp12295 := Call2(__e, SymFn(symappend), tmp12286, tmp12294)


return TailCall2(__e, SymFn(symappend), tmp12285, tmp12295)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4rule_1_6clause)
}


panic("unreachable")
})

tmp12328 := Call2(__e, ns2_1set, symshen_4rule_1_6clause, tmp12273)


_ = tmp12328

tmp12329 := MakeNative1(func(__e *ControlFlow, V3655 Obj) Obj {
tmp12330 := Call1(__e, SymFn(symshen_4macro_1_8ch), V3655)


tmp12331 := GCons(__e, symAssumptions, Nil)

return GCons(__e, tmp12330, tmp12331)


panic("unreachable")
})

tmp12332 := Call2(__e, ns2_1set, symshen_4rule_1_6head, tmp12329)


_ = tmp12332

tmp12333 := MakeNative1(func(__e *ControlFlow, V3656 Obj) Obj {
tmp12334 := GCons(__e, V3656, Nil)

return GCons(__e, symshen_4_8ch, tmp12334)


panic("unreachable")
})

tmp12335 := Call2(__e, ns2_1set, symshen_4macro_1_8ch, tmp12333)


_ = tmp12335

tmp12336 := MakeNative1(func(__e *ControlFlow, V3657 Obj) Obj {
tmp12337 := GCons(__e, V3657, Nil)

return GCons(__e, symshen_4_8c, tmp12337)


panic("unreachable")
})

tmp12338 := Call2(__e, ns2_1set, symshen_4macro_1_8c, tmp12336)


_ = tmp12338

tmp12339 := MakeNative5(func(__e *ControlFlow, V3658, V3659, V3660, V3661, V3662 Obj) Obj {
tmp12367 := GEq(__e, Nil, V3662)

if True == tmp12367 {
return TailCall5(__e, SymFn(symshen_4side_1conditions_1_6goals), Nil, V3658, V3659, V3660, V3661)
} else {
tmp12365 := GEq(__e, Nil, V3661)

var ifres12360 Obj

if True == tmp12365 {
tmp12364 := GConsP(__e, V3662)

var ifres12361 Obj

if True == tmp12364 {
tmp12362 := GTl(__e, V3662)

tmp12363 := GEq(__e, Nil, tmp12362)

ifres12361 = tmp12363


} else {
ifres12361 = False


}

ifres12360 = ifres12361


} else {
ifres12360 = False


}

if True == ifres12360 {
tmp12340 := GHd(__e, V3662)

tmp12341 := Call2(__e, SymFn(symshen_4passive_1variables), tmp12340, V3658)


let__11573 := tmp12341
_ = let__11573

tmp12342 := GHd(__e, V3662)

tmp12343 := Call2(__e, SymFn(symshen_4remove_1bystanders), V3658, tmp12342)


let__11574 := tmp12343
_ = let__11574

tmp12344 := GHd(__e, V3662)

tmp12345 := Call4(__e, SymFn(symshen_4specialise_1member), tmp12344, V3659, let__11574, let__11573)


tmp12346 := Call5(__e, SymFn(symshen_4side_1conditions_1_6goals), Nil, V3658, V3659, V3660, Nil)


return GCons(__e, tmp12345, tmp12346)


} else {
tmp12358 := GConsP(__e, V3662)

if True == tmp12358 {
tmp12347 := Call1(__e, SymFn(symgensym), symNewAssumptions)


let__11575 := tmp12347
_ = let__11575

tmp12348 := GHd(__e, V3662)

tmp12349 := Call2(__e, SymFn(symshen_4passive_1variables), tmp12348, V3658)


let__11576 := tmp12349
_ = let__11576

tmp12350 := GHd(__e, V3662)

tmp12351 := Call2(__e, SymFn(symshen_4remove_1bystanders), V3658, tmp12350)


let__11577 := tmp12351
_ = let__11577

tmp12352 := GHd(__e, V3662)

tmp12353 := Call5(__e, SymFn(symshen_4specialise_1consume), tmp12352, V3659, let__11577, let__11576, let__11575)


tmp12354 := Call2(__e, SymFn(symappend), V3658, let__11576)


tmp12355 := GTl(__e, V3662)

tmp12356 := Call5(__e, SymFn(symshen_4rule_1_6body), tmp12354, let__11575, V3660, V3661, tmp12355)


return GCons(__e, tmp12353, tmp12356)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4rule_1_6body)
}


}


}


panic("unreachable")
})

tmp12368 := Call2(__e, ns2_1set, symshen_4rule_1_6body, tmp12339)


_ = tmp12368

tmp12369 := MakeNative4(func(__e *ControlFlow, V3668, V3669, V3670, V3671 Obj) Obj {
tmp12370 := Call1(__e, SymFn(symgensym), symshen_4member)


let__11578 := tmp12370
_ = let__11578

tmp12371 := Call4(__e, SymFn(symshen_4member_1clause), let__11578, V3668, V3670, V3671)


let__11579 := tmp12371
_ = let__11579

tmp12372 := Call2(__e, SymFn(symappend), V3670, V3671)


tmp12373 := GCons(__e, V3669, tmp12372)

return GCons(__e, let__11578, tmp12373)


panic("unreachable")
})

tmp12374 := Call2(__e, ns2_1set, symshen_4specialise_1member, tmp12369)


_ = tmp12374

tmp12375 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V3676 := __self1
_ = V3676
V3677 := __self2
_ = V3677
tmp12388 := GEq(__e, Nil, V3676)

if True == tmp12388 {
return Nil
} else {
tmp12386 := GConsP(__e, V3676)

var ifres12383 Obj

if True == tmp12386 {
tmp12384 := GHd(__e, V3676)

tmp12385 := Call2(__e, SymFn(symshen_4occurs_1check_2), tmp12384, V3677)


ifres12383 = tmp12385


} else {
ifres12383 = False


}

if True == ifres12383 {
tmp12376 := GHd(__e, V3676)

tmp12377 := GTl(__e, V3676)

tmp12378 := Call2(__e, SymFn(symshen_4remove_1bystanders), tmp12377, V3677)


return GCons(__e, tmp12376, tmp12378)


} else {
tmp12381 := GConsP(__e, V3676)

if True == tmp12381 {
tmp12379 := GTl(__e, V3676)

if SymFn(symshen_4remove_1bystanders) == __self {
__self1, __self2 = tmp12379, V3677
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4remove_1bystanders), tmp12379, V3677)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4remove_1bystanders)
}


}


}


panic("unreachable")
})
return __self
})()

tmp12389 := Call2(__e, ns2_1set, symshen_4remove_1bystanders, tmp12375)


_ = tmp12389

tmp12390 := MakeNative4(func(__e *ControlFlow, V3678, V3679, V3680, V3681 Obj) Obj {
tmp12391 := Call1(__e, SymFn(symlength), V3681)


tmp12392 := Call1(__e, SymFn(symshen_4nvars), tmp12391)


let__11580 := tmp12392
_ = let__11580

tmp12393 := Call1(__e, SymFn(symshen_4macro_1_8ch), V3679)


tmp12394 := GCons(__e, sym__, Nil)

tmp12395 := GCons(__e, tmp12393, tmp12394)

tmp12396 := GCons(__e, symcons, tmp12395)

tmp12397 := GCons(__e, tmp12396, Nil)

tmp12398 := GCons(__e, sym_1, tmp12397)

tmp12399 := GCons(__e, tmp12398, Nil)

tmp12400 := GCons(__e, sym_5_1_1, Nil)

tmp12401 := Call2(__e, SymFn(symshen_4passive_1bind), V3681, let__11580)


tmp12402 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp12403 := GCons(__e, tmp12402, Nil)

tmp12404 := Call2(__e, SymFn(symappend), tmp12401, tmp12403)


tmp12405 := Call2(__e, SymFn(symappend), tmp12400, tmp12404)


tmp12406 := Call2(__e, SymFn(symappend), let__11580, tmp12405)


tmp12407 := Call2(__e, SymFn(symappend), V3680, tmp12406)


tmp12408 := Call2(__e, SymFn(symappend), tmp12399, tmp12407)


let__11581 := tmp12408
_ = let__11581

tmp12409 := Call1(__e, SymFn(symgensym), symHypotheses)


let__11583 := tmp12409
_ = let__11583

tmp12410 := Call2(__e, SymFn(symappend), V3680, V3681)


let__11584 := tmp12410
_ = let__11584

tmp12411 := GCons(__e, let__11583, Nil)

tmp12412 := GCons(__e, sym__, tmp12411)

tmp12413 := GCons(__e, symcons, tmp12412)

tmp12414 := GCons(__e, tmp12413, Nil)

tmp12415 := GCons(__e, sym_1, tmp12414)

tmp12416 := GCons(__e, tmp12415, Nil)

tmp12417 := Call2(__e, SymFn(symappend), tmp12416, let__11584)


let__11585 := tmp12417
_ = let__11585

tmp12418 := GCons(__e, let__11583, let__11584)

tmp12419 := GCons(__e, V3678, tmp12418)

tmp12420 := GCons(__e, tmp12419, Nil)

let__11586 := tmp12420
_ = let__11586

tmp12421 := GCons(__e, sym_5_1_1, Nil)

tmp12422 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp12423 := GCons(__e, tmp12422, Nil)

tmp12424 := Call2(__e, SymFn(symappend), let__11586, tmp12423)


tmp12425 := Call2(__e, SymFn(symappend), tmp12421, tmp12424)


tmp12426 := Call2(__e, SymFn(symappend), let__11585, tmp12425)


let__11582 := tmp12426
_ = let__11582

tmp12427 := Call2(__e, SymFn(symappend), let__11581, let__11582)


tmp12428 := GCons(__e, V3678, tmp12427)

tmp12429 := GCons(__e, symdefprolog, tmp12428)

let__11587 := tmp12429
_ = let__11587

return TailCall1(__e, SymFn(symeval), let__11587)


panic("unreachable")
})

tmp12430 := Call2(__e, ns2_1set, symshen_4member_1clause, tmp12390)


_ = tmp12430

tmp12431 := MakeNative1(func(__e *ControlFlow, V3690 Obj) Obj {
tmp12436 := GEq(__e, MakeInteger(0), V3690)

if True == tmp12436 {
return Nil
} else {
tmp12432 := Call1(__e, SymFn(symgensym), symNewV)


tmp12434 := Call1(__e, SymFn(symshen_4nvars), GSub(__e, V3690, MakeInteger(1)))


return GCons(__e, tmp12432, tmp12434)


}


panic("unreachable")
})

tmp12437 := Call2(__e, ns2_1set, symshen_4nvars, tmp12431)


_ = tmp12437

tmp12438 := MakeNative2(func(__e *ControlFlow, V3691, V3692 Obj) Obj {
tmp12454 := GEq(__e, Nil, V3691)

var ifres12452 Obj

if True == tmp12454 {
tmp12453 := GEq(__e, Nil, V3692)

ifres12452 = tmp12453


} else {
ifres12452 = False


}

if True == ifres12452 {
return Nil
} else {
tmp12450 := GConsP(__e, V3691)

var ifres12448 Obj

if True == tmp12450 {
tmp12449 := GConsP(__e, V3692)

ifres12448 = tmp12449


} else {
ifres12448 = False


}

if True == ifres12448 {
tmp12439 := GHd(__e, V3692)

tmp12440 := GHd(__e, V3691)

tmp12441 := GCons(__e, tmp12440, Nil)

tmp12442 := GCons(__e, tmp12439, tmp12441)

tmp12443 := GCons(__e, symbind, tmp12442)

tmp12444 := GTl(__e, V3691)

tmp12445 := GTl(__e, V3692)

tmp12446 := Call2(__e, SymFn(symshen_4passive_1bind), tmp12444, tmp12445)


return GCons(__e, tmp12443, tmp12446)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4passive_1bind)
}


}


panic("unreachable")
})

tmp12455 := Call2(__e, ns2_1set, symshen_4passive_1bind, tmp12438)


_ = tmp12455

tmp12456 := MakeNative5(func(__e *ControlFlow, V3693, V3694, V3695, V3696, V3697 Obj) Obj {
tmp12457 := Call1(__e, SymFn(symgensym), symshen_4consume)


let__11588 := tmp12457
_ = let__11588

tmp12458 := Call5(__e, SymFn(symshen_4consume_1clause), let__11588, V3693, V3695, V3696, V3697)


let__11589 := tmp12458
_ = let__11589

tmp12459 := Call2(__e, SymFn(symappend), V3695, V3696)


tmp12460 := GCons(__e, V3697, tmp12459)

tmp12461 := GCons(__e, V3694, tmp12460)

return GCons(__e, let__11588, tmp12461)


panic("unreachable")
})

tmp12462 := Call2(__e, ns2_1set, symshen_4specialise_1consume, tmp12456)


_ = tmp12462

tmp12463 := MakeNative5(func(__e *ControlFlow, V3700, V3701, V3702, V3703, V3704 Obj) Obj {
tmp12464 := Call1(__e, SymFn(symlength), V3703)


tmp12465 := Call1(__e, SymFn(symshen_4nvars), tmp12464)


let__11590 := tmp12465
_ = let__11590

tmp12466 := Call1(__e, SymFn(symgensym), symAssumption)


let__11591 := tmp12466
_ = let__11591

tmp12467 := Call1(__e, SymFn(symshen_4macro_1_8ch), V3701)


tmp12468 := GCons(__e, let__11591, Nil)

tmp12469 := GCons(__e, tmp12467, tmp12468)

tmp12470 := GCons(__e, symcons, tmp12469)

tmp12471 := GCons(__e, tmp12470, Nil)

tmp12472 := GCons(__e, sym_1, tmp12471)

tmp12473 := GCons(__e, sym_5_1_1, Nil)

tmp12474 := Call2(__e, SymFn(symshen_4passive_1bind), V3703, let__11590)


tmp12475 := GCons(__e, let__11591, Nil)

tmp12476 := GCons(__e, V3704, tmp12475)

tmp12477 := GCons(__e, symbind, tmp12476)

tmp12478 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp12479 := GCons(__e, tmp12478, Nil)

tmp12480 := GCons(__e, tmp12477, tmp12479)

tmp12481 := Call2(__e, SymFn(symappend), tmp12474, tmp12480)


tmp12482 := Call2(__e, SymFn(symappend), tmp12473, tmp12481)


tmp12483 := Call2(__e, SymFn(symappend), let__11590, tmp12482)


tmp12484 := Call2(__e, SymFn(symappend), V3702, tmp12483)


tmp12485 := GCons(__e, V3704, tmp12484)

tmp12486 := GCons(__e, tmp12472, tmp12485)

let__11592 := tmp12486
_ = let__11592

tmp12487 := Call1(__e, SymFn(symgensym), symHypotheses)


let__11594 := tmp12487
_ = let__11594

tmp12488 := Call2(__e, SymFn(symappend), V3702, V3703)


let__11595 := tmp12488
_ = let__11595

tmp12489 := Call1(__e, SymFn(symgensym), symAssumptions)


let__11596 := tmp12489
_ = let__11596

tmp12490 := GCons(__e, let__11594, Nil)

tmp12491 := GCons(__e, let__11591, tmp12490)

tmp12492 := GCons(__e, symcons, tmp12491)

tmp12493 := GCons(__e, tmp12492, Nil)

tmp12494 := GCons(__e, sym_1, tmp12493)

tmp12495 := GCons(__e, V3704, Nil)

tmp12496 := GCons(__e, let__11596, tmp12495)

tmp12497 := GCons(__e, symcons, tmp12496)

tmp12498 := GCons(__e, tmp12497, let__11595)

tmp12499 := GCons(__e, tmp12494, tmp12498)

let__11597 := tmp12499
_ = let__11597

tmp12500 := GCons(__e, let__11591, Nil)

tmp12501 := GCons(__e, let__11596, tmp12500)

tmp12502 := GCons(__e, symbind, tmp12501)

tmp12503 := GCons(__e, V3704, let__11595)

tmp12504 := GCons(__e, let__11594, tmp12503)

tmp12505 := GCons(__e, V3700, tmp12504)

tmp12506 := GCons(__e, tmp12505, Nil)

tmp12507 := GCons(__e, tmp12502, tmp12506)

let__11598 := tmp12507
_ = let__11598

tmp12508 := GCons(__e, sym_5_1_1, Nil)

tmp12509 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp12510 := GCons(__e, tmp12509, Nil)

tmp12511 := Call2(__e, SymFn(symappend), let__11598, tmp12510)


tmp12512 := Call2(__e, SymFn(symappend), tmp12508, tmp12511)


tmp12513 := Call2(__e, SymFn(symappend), let__11597, tmp12512)


let__11593 := tmp12513
_ = let__11593

tmp12514 := Call2(__e, SymFn(symappend), let__11592, let__11593)


tmp12515 := GCons(__e, V3700, tmp12514)

tmp12516 := GCons(__e, symdefprolog, tmp12515)

let__11599 := tmp12516
_ = let__11599

return TailCall1(__e, SymFn(symeval), let__11599)


panic("unreachable")
})

tmp12517 := Call2(__e, ns2_1set, symshen_4consume_1clause, tmp12463)


_ = tmp12517

tmp12518 := MakeNative2(func(__e *ControlFlow, V3715, V3716 Obj) Obj {
tmp12519 := Call1(__e, SymFn(symshen_4extract_1vars), V3715)


return TailCall2(__e, SymFn(symdifference), tmp12519, V3716)


panic("unreachable")
})

tmp12520 := Call2(__e, ns2_1set, symshen_4passive_1variables, tmp12518)


_ = tmp12520

tmp12521 := (func() Obj {
var __self Obj
__self = MakeNative5(func(__e *ControlFlow, __self1, __self2, __self3, __self4, __self5 Obj) Obj {
__selftop:
V3721 := __self1
_ = V3721
V3722 := __self2
_ = V3722
V3723 := __self3
_ = V3723
V3724 := __self4
_ = V3724
V3725 := __self5
_ = V3725
tmp12655 := GEq(__e, Nil, V3724)

if True == tmp12655 {
return TailCall3(__e, SymFn(symshen_4premises_1_6goals), V3721, V3723, V3725)
} else {
tmp12653 := GConsP(__e, V3724)

var ifres12637 Obj

if True == tmp12653 {
tmp12651 := GHd(__e, V3724)

tmp12652 := GConsP(__e, tmp12651)

var ifres12638 Obj

if True == tmp12652 {
tmp12648 := GHd(__e, V3724)

tmp12649 := GHd(__e, tmp12648)

tmp12650 := GEq(__e, symif, tmp12649)

var ifres12639 Obj

if True == tmp12650 {
tmp12645 := GHd(__e, V3724)

tmp12646 := GTl(__e, tmp12645)

tmp12647 := GConsP(__e, tmp12646)

var ifres12640 Obj

if True == tmp12647 {
tmp12641 := GHd(__e, V3724)

tmp12642 := GTl(__e, tmp12641)

tmp12643 := GTl(__e, tmp12642)

tmp12644 := GEq(__e, Nil, tmp12643)

ifres12640 = tmp12644


} else {
ifres12640 = False


}

ifres12639 = ifres12640


} else {
ifres12639 = False


}

ifres12638 = ifres12639


} else {
ifres12638 = False


}

ifres12637 = ifres12638


} else {
ifres12637 = False


}

if True == ifres12637 {
tmp12522 := GHd(__e, V3724)

tmp12523 := GTl(__e, tmp12522)

tmp12524 := GCons(__e, symwhen, tmp12523)

tmp12525 := GTl(__e, V3724)

tmp12526 := Call5(__e, SymFn(symshen_4side_1conditions_1_6goals), V3721, V3722, V3723, tmp12525, V3725)


return GCons(__e, tmp12524, tmp12526)


} else {
tmp12635 := GConsP(__e, V3724)

var ifres12613 Obj

if True == tmp12635 {
tmp12633 := GHd(__e, V3724)

tmp12634 := GConsP(__e, tmp12633)

var ifres12614 Obj

if True == tmp12634 {
tmp12630 := GHd(__e, V3724)

tmp12631 := GHd(__e, tmp12630)

tmp12632 := GEq(__e, symlet, tmp12631)

var ifres12615 Obj

if True == tmp12632 {
tmp12627 := GHd(__e, V3724)

tmp12628 := GTl(__e, tmp12627)

tmp12629 := GConsP(__e, tmp12628)

var ifres12616 Obj

if True == tmp12629 {
tmp12623 := GHd(__e, V3724)

tmp12624 := GTl(__e, tmp12623)

tmp12625 := GTl(__e, tmp12624)

tmp12626 := GConsP(__e, tmp12625)

var ifres12617 Obj

if True == tmp12626 {
tmp12618 := GHd(__e, V3724)

tmp12619 := GTl(__e, tmp12618)

tmp12620 := GTl(__e, tmp12619)

tmp12621 := GTl(__e, tmp12620)

tmp12622 := GEq(__e, Nil, tmp12621)

ifres12617 = tmp12622


} else {
ifres12617 = False


}

ifres12616 = ifres12617


} else {
ifres12616 = False


}

ifres12615 = ifres12616


} else {
ifres12615 = False


}

ifres12614 = ifres12615


} else {
ifres12614 = False


}

ifres12613 = ifres12614


} else {
ifres12613 = False


}

if True == ifres12613 {
tmp12542 := GHd(__e, V3724)

tmp12543 := GTl(__e, tmp12542)

tmp12544 := GHd(__e, tmp12543)

tmp12545 := Call2(__e, SymFn(symelement_2), tmp12544, V3722)


if True == tmp12545 {
tmp12527 := GHd(__e, V3724)

tmp12528 := GTl(__e, tmp12527)

tmp12529 := GCons(__e, symis_b, tmp12528)

tmp12530 := GTl(__e, V3724)

tmp12531 := Call5(__e, SymFn(symshen_4side_1conditions_1_6goals), V3721, V3722, V3723, tmp12530, V3725)


return GCons(__e, tmp12529, tmp12531)


} else {
tmp12532 := GHd(__e, V3724)

tmp12533 := GTl(__e, tmp12532)

tmp12534 := GCons(__e, symbind, tmp12533)

tmp12535 := GHd(__e, V3724)

tmp12536 := GTl(__e, tmp12535)

tmp12537 := GHd(__e, tmp12536)

tmp12538 := GCons(__e, tmp12537, V3722)

tmp12539 := GTl(__e, V3724)

tmp12540 := Call5(__e, SymFn(symshen_4side_1conditions_1_6goals), V3721, tmp12538, V3723, tmp12539, V3725)


return GCons(__e, tmp12534, tmp12540)


}


} else {
tmp12611 := GConsP(__e, V3724)

var ifres12595 Obj

if True == tmp12611 {
tmp12609 := GHd(__e, V3724)

tmp12610 := GConsP(__e, tmp12609)

var ifres12596 Obj

if True == tmp12610 {
tmp12606 := GHd(__e, V3724)

tmp12607 := GHd(__e, tmp12606)

tmp12608 := GEq(__e, symctxt, tmp12607)

var ifres12597 Obj

if True == tmp12608 {
tmp12603 := GHd(__e, V3724)

tmp12604 := GTl(__e, tmp12603)

tmp12605 := GConsP(__e, tmp12604)

var ifres12598 Obj

if True == tmp12605 {
tmp12599 := GHd(__e, V3724)

tmp12600 := GTl(__e, tmp12599)

tmp12601 := GTl(__e, tmp12600)

tmp12602 := GEq(__e, Nil, tmp12601)

ifres12598 = tmp12602


} else {
ifres12598 = False


}

ifres12597 = ifres12598


} else {
ifres12597 = False


}

ifres12596 = ifres12597


} else {
ifres12596 = False


}

ifres12595 = ifres12596


} else {
ifres12595 = False


}

if True == ifres12595 {
tmp12571 := GHd(__e, V3724)

tmp12572 := GTl(__e, tmp12571)

tmp12573 := GHd(__e, tmp12572)

tmp12574 := Call2(__e, SymFn(symelement_2), tmp12573, V3722)


if True == tmp12574 {
tmp12546 := GHd(__e, V3724)

tmp12547 := GTl(__e, tmp12546)

tmp12548 := GHd(__e, tmp12547)

tmp12549 := GCons(__e, tmp12548, V3721)

tmp12550 := GTl(__e, V3724)

if SymFn(symshen_4side_1conditions_1_6goals) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp12549, V3722, V3723, tmp12550, V3725
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4side_1conditions_1_6goals), tmp12549, V3722, V3723, tmp12550, V3725)


} else {
tmp12551 := GHd(__e, V3724)

tmp12552 := GTl(__e, tmp12551)

tmp12553 := GHd(__e, tmp12552)

tmp12554 := GCons(__e, V3723, Nil)

tmp12555 := GCons(__e, tmp12553, tmp12554)

tmp12556 := GCons(__e, symbind, tmp12555)

tmp12557 := GHd(__e, V3724)

tmp12558 := GTl(__e, tmp12557)

tmp12559 := GHd(__e, tmp12558)

tmp12560 := GCons(__e, tmp12559, V3721)

tmp12561 := GHd(__e, V3724)

tmp12562 := GTl(__e, tmp12561)

tmp12563 := GHd(__e, tmp12562)

tmp12564 := GCons(__e, tmp12563, V3722)

tmp12565 := GHd(__e, V3724)

tmp12566 := GTl(__e, tmp12565)

tmp12567 := GHd(__e, tmp12566)

tmp12568 := GTl(__e, V3724)

tmp12569 := Call5(__e, SymFn(symshen_4side_1conditions_1_6goals), tmp12560, tmp12564, tmp12567, tmp12568, V3725)


return GCons(__e, tmp12556, tmp12569)


}


} else {
tmp12593 := GConsP(__e, V3724)

var ifres12577 Obj

if True == tmp12593 {
tmp12591 := GHd(__e, V3724)

tmp12592 := GConsP(__e, tmp12591)

var ifres12578 Obj

if True == tmp12592 {
tmp12588 := GHd(__e, V3724)

tmp12589 := GHd(__e, tmp12588)

tmp12590 := GEq(__e, symsqts, tmp12589)

var ifres12579 Obj

if True == tmp12590 {
tmp12585 := GHd(__e, V3724)

tmp12586 := GTl(__e, tmp12585)

tmp12587 := GConsP(__e, tmp12586)

var ifres12580 Obj

if True == tmp12587 {
tmp12581 := GHd(__e, V3724)

tmp12582 := GTl(__e, tmp12581)

tmp12583 := GTl(__e, tmp12582)

tmp12584 := GEq(__e, Nil, tmp12583)

ifres12580 = tmp12584


} else {
ifres12580 = False


}

ifres12579 = ifres12580


} else {
ifres12579 = False


}

ifres12578 = ifres12579


} else {
ifres12578 = False


}

ifres12577 = ifres12578


} else {
ifres12577 = False


}

if True == ifres12577 {
tmp12575 := GTl(__e, V3724)

if SymFn(symshen_4side_1conditions_1_6goals) == __self {
__self1, __self2, __self3, __self4, __self5 = V3721, V3722, V3723, tmp12575, V3725
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4side_1conditions_1_6goals), V3721, V3722, V3723, tmp12575, V3725)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4side_1conditions_1_6goals)
}


}


}


}


}


panic("unreachable")
})
return __self
})()

tmp12656 := Call2(__e, ns2_1set, symshen_4side_1conditions_1_6goals, tmp12521)


_ = tmp12656

tmp12657 := MakeNative3(func(__e *ControlFlow, V3730, V3731, V3732 Obj) Obj {
tmp12702 := GEq(__e, Nil, V3732)

if True == tmp12702 {
tmp12658 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

return GCons(__e, tmp12658, Nil)


} else {
tmp12700 := GConsP(__e, V3732)

var ifres12697 Obj

if True == tmp12700 {
tmp12698 := GHd(__e, V3732)

tmp12699 := GEq(__e, sym_b, tmp12698)

ifres12697 = tmp12699


} else {
ifres12697 = False


}

if True == ifres12697 {
tmp12659 := GTl(__e, V3732)

tmp12660 := Call3(__e, SymFn(symshen_4premises_1_6goals), V3730, V3731, tmp12659)


return GCons(__e, sym_b, tmp12660)


} else {
tmp12695 := GConsP(__e, V3732)

var ifres12692 Obj

if True == tmp12695 {
tmp12693 := GHd(__e, V3732)

tmp12694 := GEq(__e, symfail, tmp12693)

ifres12692 = tmp12694


} else {
ifres12692 = False


}

if True == ifres12692 {
tmp12661 := GCons(__e, False, Nil)

tmp12662 := GCons(__e, symwhen, tmp12661)

tmp12663 := GTl(__e, V3732)

tmp12664 := Call3(__e, SymFn(symshen_4premises_1_6goals), V3730, V3731, tmp12663)


return GCons(__e, tmp12662, tmp12664)


} else {
tmp12690 := GConsP(__e, V3732)

var ifres12678 Obj

if True == tmp12690 {
tmp12688 := GHd(__e, V3732)

tmp12689 := GConsP(__e, tmp12688)

var ifres12679 Obj

if True == tmp12689 {
tmp12685 := GHd(__e, V3732)

tmp12686 := GTl(__e, tmp12685)

tmp12687 := GConsP(__e, tmp12686)

var ifres12680 Obj

if True == tmp12687 {
tmp12681 := GHd(__e, V3732)

tmp12682 := GTl(__e, tmp12681)

tmp12683 := GTl(__e, tmp12682)

tmp12684 := GEq(__e, Nil, tmp12683)

ifres12680 = tmp12684


} else {
ifres12680 = False


}

ifres12679 = ifres12680


} else {
ifres12679 = False


}

ifres12678 = ifres12679


} else {
ifres12678 = False


}

if True == ifres12678 {
tmp12665 := GHd(__e, V3732)

tmp12666 := GTl(__e, tmp12665)

tmp12667 := GHd(__e, tmp12666)

tmp12668 := Call1(__e, SymFn(symshen_4macro_1_8c), tmp12667)


tmp12669 := GHd(__e, V3732)

tmp12670 := GHd(__e, tmp12669)

tmp12671 := Call3(__e, SymFn(symshen_4construct_1context), V3730, tmp12670, V3731)


tmp12672 := GCons(__e, tmp12671, Nil)

tmp12673 := GCons(__e, tmp12668, tmp12672)

tmp12674 := GCons(__e, symshen_4system_1S, tmp12673)

tmp12675 := GTl(__e, V3732)

tmp12676 := Call3(__e, SymFn(symshen_4premises_1_6goals), V3730, V3731, tmp12675)


return GCons(__e, tmp12674, tmp12676)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4premises_1_6goals)
}


}


}


}


panic("unreachable")
})

tmp12703 := Call2(__e, ns2_1set, symshen_4premises_1_6goals, tmp12657)


_ = tmp12703

tmp12704 := MakeNative3(func(__e *ControlFlow, V3736, V3737, V3738 Obj) Obj {
tmp12722 := GEq(__e, Nil, V3737)

if True == tmp12722 {
return V3738
} else {
tmp12720 := GConsP(__e, V3737)

var ifres12714 Obj

if True == tmp12720 {
tmp12718 := GTl(__e, V3737)

tmp12719 := GEq(__e, Nil, tmp12718)

var ifres12715 Obj

if True == tmp12719 {
tmp12716 := GHd(__e, V3737)

tmp12717 := Call2(__e, SymFn(symelement_2), tmp12716, V3736)


ifres12715 = tmp12717


} else {
ifres12715 = False


}

ifres12714 = ifres12715


} else {
ifres12714 = False


}

if True == ifres12714 {
return GHd(__e, V3737)
} else {
tmp12712 := GConsP(__e, V3737)

if True == tmp12712 {
tmp12705 := GHd(__e, V3737)

tmp12706 := Call1(__e, SymFn(symshen_4macro_1_8c), tmp12705)


tmp12707 := GTl(__e, V3737)

tmp12708 := Call3(__e, SymFn(symshen_4construct_1context), V3736, tmp12707, V3738)


tmp12709 := GCons(__e, tmp12708, Nil)

tmp12710 := GCons(__e, tmp12706, tmp12709)

return GCons(__e, symcons, tmp12710)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4construct_1context)
}


}


}


panic("unreachable")
})

tmp12723 := Call2(__e, ns2_1set, symshen_4construct_1context, tmp12704)


_ = tmp12723

tmp12724 := MakeNative1(func(__e *ControlFlow, V3739 Obj) Obj {
tmp12725 := MakeNative1(func(__e *ControlFlow, Z3741 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4intern_1type), Z3741)
panic("unreachable")
})

tmp12726 := Call2(__e, SymFn(symmap), tmp12725, V3739)


let__11600 := tmp12726
_ = let__11600

tmp12727 := GValue(__e, symshen_4_ddatatypes_d)

let__11601 := tmp12727
_ = let__11601

tmp12728 := Call2(__e, SymFn(symshen_4remove_1datatypes), let__11600, let__11601)


let__11602 := tmp12728
_ = let__11602

tmp12729 := PrimGuard2(__e, symset, PrimSet, symshen_4_ddatatypes_d, let__11602)

let__11603 := tmp12729
_ = let__11603

return TailCall1(__e, SymFn(symshen_4show_1datatypes), let__11603)


panic("unreachable")
})

tmp12730 := Call2(__e, ns2_1set, sympreclude, tmp12724)


_ = tmp12730

tmp12731 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V3749 := __self1
_ = V3749
V3750 := __self2
_ = V3750
tmp12738 := GEq(__e, Nil, V3749)

if True == tmp12738 {
return V3750
} else {
tmp12736 := GConsP(__e, V3749)

if True == tmp12736 {
tmp12732 := GTl(__e, V3749)

tmp12733 := GHd(__e, V3749)

tmp12734 := Call2(__e, SymFn(symshen_4unassoc), tmp12733, V3750)


if SymFn(symshen_4remove_1datatypes) == __self {
__self1, __self2 = tmp12732, tmp12734
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4remove_1datatypes), tmp12732, tmp12734)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.remove-datatypes"))
}


}


panic("unreachable")
})
return __self
})()

tmp12739 := Call2(__e, ns2_1set, symshen_4remove_1datatypes, tmp12731)


_ = tmp12739

tmp12740 := MakeNative1(func(__e *ControlFlow, V3751 Obj) Obj {
tmp12741 := MakeNative1(func(__e *ControlFlow, Z3752 Obj) Obj {
return GHd(__e, Z3752)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp12741, V3751)


panic("unreachable")
})

tmp12742 := Call2(__e, ns2_1set, symshen_4show_1datatypes, tmp12740)


_ = tmp12742

tmp12743 := MakeNative1(func(__e *ControlFlow, V3753 Obj) Obj {
tmp12744 := MakeNative1(func(__e *ControlFlow, Z3755 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4intern_1type), Z3755)
panic("unreachable")
})

tmp12745 := Call2(__e, SymFn(symmap), tmp12744, V3753)


let__11604 := tmp12745
_ = let__11604

tmp12746 := MakeNative1(func(__e *ControlFlow, Z3757 Obj) Obj {
tmp12747 := Call1(__e, SymFn(symfn), Z3757)


return TailCall2(__e, SymFn(symshen_4remember_1datatype), Z3757, tmp12747)


panic("unreachable")
})

tmp12748 := Call2(__e, SymFn(symmap), tmp12746, let__11604)


let__11605 := tmp12748
_ = let__11605

tmp12749 := GValue(__e, symshen_4_ddatatypes_d)

let__11606 := tmp12749
_ = let__11606

return TailCall1(__e, SymFn(symshen_4show_1datatypes), let__11606)


panic("unreachable")
})

tmp12750 := Call2(__e, ns2_1set, syminclude, tmp12743)


_ = tmp12750

tmp12751 := MakeNative1(func(__e *ControlFlow, V3759 Obj) Obj {
tmp12752 := PrimGuard2(__e, symset, PrimSet, symshen_4_ddatatypes_d, Nil)

let__11607 := tmp12752
_ = let__11607

tmp12753 := MakeNative1(func(__e *ControlFlow, Z3762 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4intern_1type), Z3762)
panic("unreachable")
})

tmp12754 := Call2(__e, SymFn(symmap), tmp12753, V3759)


let__11608 := tmp12754
_ = let__11608

tmp12755 := MakeNative1(func(__e *ControlFlow, Z3764 Obj) Obj {
tmp12756 := Call1(__e, SymFn(symfn), Z3764)


return TailCall2(__e, SymFn(symshen_4remember_1datatype), Z3764, tmp12756)


panic("unreachable")
})

tmp12757 := Call2(__e, SymFn(symmap), tmp12755, let__11608)


let__11609 := tmp12757
_ = let__11609

tmp12758 := GValue(__e, symshen_4_ddatatypes_d)

return TailCall1(__e, SymFn(symshen_4show_1datatypes), tmp12758)


panic("unreachable")
})

tmp12759 := Call2(__e, ns2_1set, sympreclude_1all_1but, tmp12751)


_ = tmp12759

tmp12760 := MakeNative1(func(__e *ControlFlow, V3765 Obj) Obj {
tmp12761 := MakeNative1(func(__e *ControlFlow, Z3767 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4intern_1type), Z3767)
panic("unreachable")
})

tmp12762 := Call2(__e, SymFn(symmap), tmp12761, V3765)


let__11610 := tmp12762
_ = let__11610

tmp12763 := GValue(__e, symshen_4_dalldatatypes_d)

let__11611 := tmp12763
_ = let__11611

tmp12764 := Call2(__e, SymFn(symshen_4remove_1datatypes), let__11610, let__11611)


tmp12765 := PrimGuard2(__e, symset, PrimSet, symshen_4_ddatatypes_d, tmp12764)

let__11612 := tmp12765
_ = let__11612

return TailCall1(__e, SymFn(symshen_4show_1datatypes), let__11612)


panic("unreachable")
})

return TailCall2(__e, ns2_1set, syminclude_1all_1but, tmp12760)




panic("unreachable")
})

