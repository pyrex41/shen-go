package main

import . "github.com/pyrex41/shen-go/kl"

var SequentMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp11633 := MakeNative1(func(__e *ControlFlow, V3418 Obj) Obj {
tmp11648 := GConsP(__e, V3418)

var ifres11634 Obj

if True == tmp11648 {
tmp11635 := Call1(__e, SymFn(symhead), V3418)


let__11403 := tmp11635
_ = let__11403

tmp11636 := Call1(__e, SymFn(symtail), V3418)


let__11404 := tmp11636
_ = let__11404

tmp11637 := Call1(__e, SymFn(symshen_4_5datatype_1rules_6), let__11404)


let__11405 := tmp11637
_ = let__11405

tmp11646 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11405)


var ifres11638 Obj

if True == tmp11646 {
tmp11639 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11638 = tmp11639


} else {
tmp11640 := Call1(__e, SymFn(symshen_4_5_1out), let__11405)


let__11406 := tmp11640
_ = let__11406

tmp11641 := Call1(__e, SymFn(symshen_4in_1_6), let__11405)


let__11407 := tmp11641
_ = let__11407

tmp11642 := Call2(__e, SymFn(symshen_4rules_1_6prolog), let__11403, let__11406)


let__11408 := tmp11642
_ = let__11408

tmp11643 := Call1(__e, SymFn(symfn), let__11403)


tmp11644 := Call2(__e, SymFn(symshen_4remember_1datatype), let__11403, tmp11643)


tmp11645 := Call2(__e, SymFn(symshen_4comb), let__11407, tmp11644)


ifres11638 = tmp11645


}

ifres11634 = ifres11638


} else {
tmp11647 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11634 = tmp11647


}

let__11402 := ifres11634
_ = let__11402

tmp11650 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11402)


if True == tmp11650 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11402
}


panic("unreachable")
})

tmp11651 := Call2(__e, ns2_1set, symshen_4_5datatype_6, tmp11633)


_ = tmp11651

tmp11652 := MakeNative2(func(__e *ControlFlow, V3426, V3427 Obj) Obj {
tmp11653 := GValue(__e, symshen_4_ddatatypes_d)

tmp11654 := Call3(__e, SymFn(symshen_4assoc_1_6), V3426, V3427, tmp11653)


tmp11655 := PrimGuard2(__e, symset, PrimSet, symshen_4_ddatatypes_d, tmp11654)

_ = tmp11655

tmp11656 := GValue(__e, symshen_4_dalldatatypes_d)

tmp11657 := Call3(__e, SymFn(symshen_4assoc_1_6), V3426, V3427, tmp11656)


tmp11658 := PrimGuard2(__e, symset, PrimSet, symshen_4_dalldatatypes_d, tmp11657)

_ = tmp11658

return V3426


panic("unreachable")
})

tmp11659 := Call2(__e, ns2_1set, symshen_4remember_1datatype, tmp11652)


_ = tmp11659

tmp11660 := MakeNative1(func(__e *ControlFlow, V3428 Obj) Obj {
tmp11661 := Call1(__e, SymFn(symshen_4_5datatype_1rule_6), V3428)


let__11410 := tmp11661
_ = let__11410

tmp11674 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11410)


var ifres11662 Obj

if True == tmp11674 {
tmp11663 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11662 = tmp11663


} else {
tmp11664 := Call1(__e, SymFn(symshen_4_5_1out), let__11410)


let__11411 := tmp11664
_ = let__11411

tmp11665 := Call1(__e, SymFn(symshen_4in_1_6), let__11410)


let__11412 := tmp11665
_ = let__11412

tmp11666 := Call1(__e, SymFn(symshen_4_5datatype_1rules_6), let__11412)


let__11413 := tmp11666
_ = let__11413

tmp11673 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11413)


var ifres11667 Obj

if True == tmp11673 {
tmp11668 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11667 = tmp11668


} else {
tmp11669 := Call1(__e, SymFn(symshen_4_5_1out), let__11413)


let__11414 := tmp11669
_ = let__11414

tmp11670 := Call1(__e, SymFn(symshen_4in_1_6), let__11413)


let__11415 := tmp11670
_ = let__11415

tmp11671 := Call2(__e, SymFn(symappend), let__11411, let__11414)


tmp11672 := Call2(__e, SymFn(symshen_4comb), let__11415, tmp11671)


ifres11667 = tmp11672


}

ifres11662 = ifres11667


}

let__11409 := ifres11662
_ = let__11409

tmp11690 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11409)


if True == tmp11690 {
tmp11675 := Call1(__e, SymFn(sym_5_b_6), V3428)


let__11417 := tmp11675
_ = let__11417

tmp11686 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11417)


var ifres11676 Obj

if True == tmp11686 {
tmp11677 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11676 = tmp11677


} else {
tmp11678 := Call1(__e, SymFn(symshen_4_5_1out), let__11417)


let__11418 := tmp11678
_ = let__11418

tmp11679 := Call1(__e, SymFn(symshen_4in_1_6), let__11417)


let__11419 := tmp11679
_ = let__11419

tmp11684 := Call1(__e, SymFn(symempty_2), let__11418)


var ifres11680 Obj

if True == tmp11684 {
ifres11680 = Nil


} else {
tmp11681 := Call3(__e, SymFn(symshen_4app), let__11418, MakeString("\n ..."), symshen_4r)


tmp11683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("datatype syntax error here:\n "), tmp11681))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("datatype syntax error here:\n "), tmp11681)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()

ifres11680 = tmp11683


}

tmp11685 := Call2(__e, SymFn(symshen_4comb), let__11419, ifres11680)


ifres11676 = tmp11685


}

let__11416 := ifres11676
_ = let__11416

tmp11688 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11416)


if True == tmp11688 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11416
}


} else {
return let__11409
}


panic("unreachable")
})

tmp11691 := Call2(__e, ns2_1set, symshen_4_5datatype_1rules_6, tmp11660)


_ = tmp11691

tmp11692 := MakeNative1(func(__e *ControlFlow, V3440 Obj) Obj {
tmp11693 := Call1(__e, SymFn(symshen_4_5single_6), V3440)


let__11421 := tmp11693
_ = let__11421

tmp11699 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11421)


var ifres11694 Obj

if True == tmp11699 {
tmp11695 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11694 = tmp11695


} else {
tmp11696 := Call1(__e, SymFn(symshen_4_5_1out), let__11421)


let__11422 := tmp11696
_ = let__11422

tmp11697 := Call1(__e, SymFn(symshen_4in_1_6), let__11421)


let__11423 := tmp11697
_ = let__11423

tmp11698 := Call2(__e, SymFn(symshen_4comb), let__11423, let__11422)


ifres11694 = tmp11698


}

let__11420 := ifres11694
_ = let__11420

tmp11710 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11420)


if True == tmp11710 {
tmp11700 := Call1(__e, SymFn(symshen_4_5double_6), V3440)


let__11425 := tmp11700
_ = let__11425

tmp11706 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11425)


var ifres11701 Obj

if True == tmp11706 {
tmp11702 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11701 = tmp11702


} else {
tmp11703 := Call1(__e, SymFn(symshen_4_5_1out), let__11425)


let__11426 := tmp11703
_ = let__11426

tmp11704 := Call1(__e, SymFn(symshen_4in_1_6), let__11425)


let__11427 := tmp11704
_ = let__11427

tmp11705 := Call2(__e, SymFn(symshen_4comb), let__11427, let__11426)


ifres11701 = tmp11705


}

let__11424 := ifres11701
_ = let__11424

tmp11708 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11424)


if True == tmp11708 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11424
}


} else {
return let__11420
}


panic("unreachable")
})

tmp11711 := Call2(__e, ns2_1set, symshen_4_5datatype_1rule_6, tmp11692)


_ = tmp11711

tmp11712 := MakeNative1(func(__e *ControlFlow, V3449 Obj) Obj {
tmp11713 := Call1(__e, SymFn(symshen_4_5sides_6), V3449)


let__11429 := tmp11713
_ = let__11429

tmp11745 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11429)


var ifres11714 Obj

if True == tmp11745 {
tmp11715 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11714 = tmp11715


} else {
tmp11716 := Call1(__e, SymFn(symshen_4_5_1out), let__11429)


let__11430 := tmp11716
_ = let__11430

tmp11717 := Call1(__e, SymFn(symshen_4in_1_6), let__11429)


let__11431 := tmp11717
_ = let__11431

tmp11718 := Call1(__e, SymFn(symshen_4_5prems_6), let__11431)


let__11432 := tmp11718
_ = let__11432

tmp11744 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11432)


var ifres11719 Obj

if True == tmp11744 {
tmp11720 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11719 = tmp11720


} else {
tmp11721 := Call1(__e, SymFn(symshen_4_5_1out), let__11432)


let__11433 := tmp11721
_ = let__11433

tmp11722 := Call1(__e, SymFn(symshen_4in_1_6), let__11432)


let__11434 := tmp11722
_ = let__11434

tmp11723 := Call1(__e, SymFn(symshen_4_5sng_6), let__11434)


let__11435 := tmp11723
_ = let__11435

tmp11743 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11435)


var ifres11724 Obj

if True == tmp11743 {
tmp11725 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11724 = tmp11725


} else {
tmp11726 := Call1(__e, SymFn(symshen_4in_1_6), let__11435)


let__11436 := tmp11726
_ = let__11436

tmp11727 := Call1(__e, SymFn(symshen_4_5conc_6), let__11436)


let__11437 := tmp11727
_ = let__11437

tmp11742 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11437)


var ifres11728 Obj

if True == tmp11742 {
tmp11729 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11728 = tmp11729


} else {
tmp11730 := Call1(__e, SymFn(symshen_4_5_1out), let__11437)


let__11438 := tmp11730
_ = let__11438

tmp11731 := Call1(__e, SymFn(symshen_4in_1_6), let__11437)


let__11439 := tmp11731
_ = let__11439

tmp11732 := Call1(__e, SymFn(symshen_4_5sc_6), let__11439)


let__11440 := tmp11732
_ = let__11440

tmp11741 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11440)


var ifres11733 Obj

if True == tmp11741 {
tmp11734 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11733 = tmp11734


} else {
tmp11735 := Call1(__e, SymFn(symshen_4in_1_6), let__11440)


let__11441 := tmp11735
_ = let__11441

tmp11736 := GCons(__e, let__11438, Nil)

tmp11737 := GCons(__e, let__11433, tmp11736)

tmp11738 := GCons(__e, let__11430, tmp11737)

tmp11739 := GCons(__e, tmp11738, Nil)

tmp11740 := Call2(__e, SymFn(symshen_4comb), let__11441, tmp11739)


ifres11733 = tmp11740


}

ifres11728 = ifres11733


}

ifres11724 = ifres11728


}

ifres11719 = ifres11724


}

ifres11714 = ifres11719


}

let__11428 := ifres11714
_ = let__11428

tmp11747 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11428)


if True == tmp11747 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11428
}


panic("unreachable")
})

tmp11748 := Call2(__e, ns2_1set, symshen_4_5single_6, tmp11712)


_ = tmp11748

tmp11749 := MakeNative1(func(__e *ControlFlow, V3464 Obj) Obj {
tmp11750 := Call1(__e, SymFn(symshen_4_5sides_6), V3464)


let__11443 := tmp11750
_ = let__11443

tmp11781 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11443)


var ifres11751 Obj

if True == tmp11781 {
tmp11752 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11751 = tmp11752


} else {
tmp11753 := Call1(__e, SymFn(symshen_4_5_1out), let__11443)


let__11444 := tmp11753
_ = let__11444

tmp11754 := Call1(__e, SymFn(symshen_4in_1_6), let__11443)


let__11445 := tmp11754
_ = let__11445

tmp11755 := Call1(__e, SymFn(symshen_4_5formulae_6), let__11445)


let__11446 := tmp11755
_ = let__11446

tmp11780 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11446)


var ifres11756 Obj

if True == tmp11780 {
tmp11757 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11756 = tmp11757


} else {
tmp11758 := Call1(__e, SymFn(symshen_4_5_1out), let__11446)


let__11447 := tmp11758
_ = let__11447

tmp11759 := Call1(__e, SymFn(symshen_4in_1_6), let__11446)


let__11448 := tmp11759
_ = let__11448

tmp11760 := Call1(__e, SymFn(symshen_4_5dbl_6), let__11448)


let__11449 := tmp11760
_ = let__11449

tmp11779 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11449)


var ifres11761 Obj

if True == tmp11779 {
tmp11762 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11761 = tmp11762


} else {
tmp11763 := Call1(__e, SymFn(symshen_4in_1_6), let__11449)


let__11450 := tmp11763
_ = let__11450

tmp11764 := Call1(__e, SymFn(symshen_4_5formula_6), let__11450)


let__11451 := tmp11764
_ = let__11451

tmp11778 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11451)


var ifres11765 Obj

if True == tmp11778 {
tmp11766 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11765 = tmp11766


} else {
tmp11767 := Call1(__e, SymFn(symshen_4_5_1out), let__11451)


let__11452 := tmp11767
_ = let__11452

tmp11768 := Call1(__e, SymFn(symshen_4in_1_6), let__11451)


let__11453 := tmp11768
_ = let__11453

tmp11769 := Call1(__e, SymFn(symshen_4_5sc_6), let__11453)


let__11454 := tmp11769
_ = let__11454

tmp11777 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11454)


var ifres11770 Obj

if True == tmp11777 {
tmp11771 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11770 = tmp11771


} else {
tmp11772 := Call1(__e, SymFn(symshen_4in_1_6), let__11454)


let__11455 := tmp11772
_ = let__11455

tmp11773 := GCons(__e, let__11452, Nil)

tmp11774 := GCons(__e, Nil, tmp11773)

tmp11775 := Call3(__e, SymFn(symshen_4lr_1rule), let__11444, let__11447, tmp11774)


tmp11776 := Call2(__e, SymFn(symshen_4comb), let__11455, tmp11775)


ifres11770 = tmp11776


}

ifres11765 = ifres11770


}

ifres11761 = ifres11765


}

ifres11756 = ifres11761


}

ifres11751 = ifres11756


}

let__11442 := ifres11751
_ = let__11442

tmp11783 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11442)


if True == tmp11783 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11442
}


panic("unreachable")
})

tmp11784 := Call2(__e, ns2_1set, symshen_4_5double_6, tmp11749)


_ = tmp11784

tmp11785 := MakeNative1(func(__e *ControlFlow, V3479 Obj) Obj {
tmp11786 := Call1(__e, SymFn(symshen_4_5formula_6), V3479)


let__11457 := tmp11786
_ = let__11457

tmp11806 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11457)


var ifres11787 Obj

if True == tmp11806 {
tmp11788 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11787 = tmp11788


} else {
tmp11789 := Call1(__e, SymFn(symshen_4_5_1out), let__11457)


let__11458 := tmp11789
_ = let__11458

tmp11790 := Call1(__e, SymFn(symshen_4in_1_6), let__11457)


let__11459 := tmp11790
_ = let__11459

tmp11791 := Call1(__e, SymFn(symshen_4_5sc_6), let__11459)


let__11460 := tmp11791
_ = let__11460

tmp11805 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11460)


var ifres11792 Obj

if True == tmp11805 {
tmp11793 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11792 = tmp11793


} else {
tmp11794 := Call1(__e, SymFn(symshen_4in_1_6), let__11460)


let__11461 := tmp11794
_ = let__11461

tmp11795 := Call1(__e, SymFn(symshen_4_5formulae_6), let__11461)


let__11462 := tmp11795
_ = let__11462

tmp11804 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11462)


var ifres11796 Obj

if True == tmp11804 {
tmp11797 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11796 = tmp11797


} else {
tmp11798 := Call1(__e, SymFn(symshen_4_5_1out), let__11462)


let__11463 := tmp11798
_ = let__11463

tmp11799 := Call1(__e, SymFn(symshen_4in_1_6), let__11462)


let__11464 := tmp11799
_ = let__11464

tmp11800 := GCons(__e, let__11458, Nil)

tmp11801 := GCons(__e, Nil, tmp11800)

tmp11802 := GCons(__e, tmp11801, let__11463)

tmp11803 := Call2(__e, SymFn(symshen_4comb), let__11464, tmp11802)


ifres11796 = tmp11803


}

ifres11792 = ifres11796


}

ifres11787 = ifres11792


}

let__11456 := ifres11787
_ = let__11456

tmp11825 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11456)


if True == tmp11825 {
tmp11807 := Call1(__e, SymFn(symshen_4_5formula_6), V3479)


let__11466 := tmp11807
_ = let__11466

tmp11821 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11466)


var ifres11808 Obj

if True == tmp11821 {
tmp11809 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11808 = tmp11809


} else {
tmp11810 := Call1(__e, SymFn(symshen_4_5_1out), let__11466)


let__11467 := tmp11810
_ = let__11467

tmp11811 := Call1(__e, SymFn(symshen_4in_1_6), let__11466)


let__11468 := tmp11811
_ = let__11468

tmp11812 := Call1(__e, SymFn(symshen_4_5sc_6), let__11468)


let__11469 := tmp11812
_ = let__11469

tmp11820 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11469)


var ifres11813 Obj

if True == tmp11820 {
tmp11814 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11813 = tmp11814


} else {
tmp11815 := Call1(__e, SymFn(symshen_4in_1_6), let__11469)


let__11470 := tmp11815
_ = let__11470

tmp11816 := GCons(__e, let__11467, Nil)

tmp11817 := GCons(__e, Nil, tmp11816)

tmp11818 := GCons(__e, tmp11817, Nil)

tmp11819 := Call2(__e, SymFn(symshen_4comb), let__11470, tmp11818)


ifres11813 = tmp11819


}

ifres11808 = ifres11813


}

let__11465 := ifres11808
_ = let__11465

tmp11823 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11465)


if True == tmp11823 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11465
}


} else {
return let__11456
}


panic("unreachable")
})

tmp11826 := Call2(__e, ns2_1set, symshen_4_5formulae_6, tmp11785)


_ = tmp11826

tmp11827 := MakeNative1(func(__e *ControlFlow, V3495 Obj) Obj {
tmp11828 := Call1(__e, SymFn(symshen_4_5ass_6), V3495)


let__11472 := tmp11828
_ = let__11472

tmp11846 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11472)


var ifres11829 Obj

if True == tmp11846 {
tmp11830 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11829 = tmp11830


} else {
tmp11831 := Call1(__e, SymFn(symshen_4_5_1out), let__11472)


let__11473 := tmp11831
_ = let__11473

tmp11832 := Call1(__e, SymFn(symshen_4in_1_6), let__11472)


let__11474 := tmp11832
_ = let__11474

tmp11845 := Call2(__e, SymFn(symshen_4hds_a_2), let__11474, sym_6_6)


var ifres11833 Obj

if True == tmp11845 {
tmp11834 := Call1(__e, SymFn(symtail), let__11474)


let__11475 := tmp11834
_ = let__11475

tmp11835 := Call1(__e, SymFn(symshen_4_5formula_6), let__11475)


let__11476 := tmp11835
_ = let__11476

tmp11843 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11476)


var ifres11836 Obj

if True == tmp11843 {
tmp11837 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11836 = tmp11837


} else {
tmp11838 := Call1(__e, SymFn(symshen_4_5_1out), let__11476)


let__11477 := tmp11838
_ = let__11477

tmp11839 := Call1(__e, SymFn(symshen_4in_1_6), let__11476)


let__11478 := tmp11839
_ = let__11478

tmp11840 := GCons(__e, let__11477, Nil)

tmp11841 := GCons(__e, let__11473, tmp11840)

tmp11842 := Call2(__e, SymFn(symshen_4comb), let__11478, tmp11841)


ifres11836 = tmp11842


}

ifres11833 = ifres11836


} else {
tmp11844 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11833 = tmp11844


}

ifres11829 = ifres11833


}

let__11471 := ifres11829
_ = let__11471

tmp11859 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11471)


if True == tmp11859 {
tmp11847 := Call1(__e, SymFn(symshen_4_5formula_6), V3495)


let__11480 := tmp11847
_ = let__11480

tmp11855 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11480)


var ifres11848 Obj

if True == tmp11855 {
tmp11849 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11848 = tmp11849


} else {
tmp11850 := Call1(__e, SymFn(symshen_4_5_1out), let__11480)


let__11481 := tmp11850
_ = let__11481

tmp11851 := Call1(__e, SymFn(symshen_4in_1_6), let__11480)


let__11482 := tmp11851
_ = let__11482

tmp11852 := GCons(__e, let__11481, Nil)

tmp11853 := GCons(__e, Nil, tmp11852)

tmp11854 := Call2(__e, SymFn(symshen_4comb), let__11482, tmp11853)


ifres11848 = tmp11854


}

let__11479 := ifres11848
_ = let__11479

tmp11857 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11479)


if True == tmp11857 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11479
}


} else {
return let__11471
}


panic("unreachable")
})

tmp11860 := Call2(__e, ns2_1set, symshen_4_5conc_6, tmp11827)


_ = tmp11860

tmp11861 := MakeNative1(func(__e *ControlFlow, V3508 Obj) Obj {
tmp11862 := Call1(__e, SymFn(symshen_4_5prem_6), V3508)


let__11484 := tmp11862
_ = let__11484

tmp11880 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11484)


var ifres11863 Obj

if True == tmp11880 {
tmp11864 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11863 = tmp11864


} else {
tmp11865 := Call1(__e, SymFn(symshen_4_5_1out), let__11484)


let__11485 := tmp11865
_ = let__11485

tmp11866 := Call1(__e, SymFn(symshen_4in_1_6), let__11484)


let__11486 := tmp11866
_ = let__11486

tmp11867 := Call1(__e, SymFn(symshen_4_5sc_6), let__11486)


let__11487 := tmp11867
_ = let__11487

tmp11879 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11487)


var ifres11868 Obj

if True == tmp11879 {
tmp11869 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11868 = tmp11869


} else {
tmp11870 := Call1(__e, SymFn(symshen_4in_1_6), let__11487)


let__11488 := tmp11870
_ = let__11488

tmp11871 := Call1(__e, SymFn(symshen_4_5prems_6), let__11488)


let__11489 := tmp11871
_ = let__11489

tmp11878 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11489)


var ifres11872 Obj

if True == tmp11878 {
tmp11873 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11872 = tmp11873


} else {
tmp11874 := Call1(__e, SymFn(symshen_4_5_1out), let__11489)


let__11490 := tmp11874
_ = let__11490

tmp11875 := Call1(__e, SymFn(symshen_4in_1_6), let__11489)


let__11491 := tmp11875
_ = let__11491

tmp11876 := GCons(__e, let__11485, let__11490)

tmp11877 := Call2(__e, SymFn(symshen_4comb), let__11491, tmp11876)


ifres11872 = tmp11877


}

ifres11868 = ifres11872


}

ifres11863 = ifres11868


}

let__11483 := ifres11863
_ = let__11483

tmp11890 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11483)


if True == tmp11890 {
tmp11881 := Call1(__e, SymFn(sym_5e_6), V3508)


let__11493 := tmp11881
_ = let__11493

tmp11886 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11493)


var ifres11882 Obj

if True == tmp11886 {
tmp11883 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11882 = tmp11883


} else {
tmp11884 := Call1(__e, SymFn(symshen_4in_1_6), let__11493)


let__11494 := tmp11884
_ = let__11494

tmp11885 := Call2(__e, SymFn(symshen_4comb), let__11494, Nil)


ifres11882 = tmp11885


}

let__11492 := ifres11882
_ = let__11492

tmp11888 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11492)


if True == tmp11888 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11492
}


} else {
return let__11483
}


panic("unreachable")
})

tmp11891 := Call2(__e, ns2_1set, symshen_4_5prems_6, tmp11861)


_ = tmp11891

tmp11892 := MakeNative1(func(__e *ControlFlow, V3521 Obj) Obj {
tmp11897 := Call2(__e, SymFn(symshen_4hds_a_2), V3521, sym_b)


var ifres11893 Obj

if True == tmp11897 {
tmp11894 := Call1(__e, SymFn(symtail), V3521)


let__11496 := tmp11894
_ = let__11496

tmp11895 := Call2(__e, SymFn(symshen_4comb), let__11496, sym_b)


ifres11893 = tmp11895


} else {
tmp11896 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11893 = tmp11896


}

let__11495 := ifres11893
_ = let__11495

tmp11931 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11495)


if True == tmp11931 {
tmp11898 := Call1(__e, SymFn(symshen_4_5ass_6), V3521)


let__11498 := tmp11898
_ = let__11498

tmp11916 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11498)


var ifres11899 Obj

if True == tmp11916 {
tmp11900 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11899 = tmp11900


} else {
tmp11901 := Call1(__e, SymFn(symshen_4_5_1out), let__11498)


let__11499 := tmp11901
_ = let__11499

tmp11902 := Call1(__e, SymFn(symshen_4in_1_6), let__11498)


let__11500 := tmp11902
_ = let__11500

tmp11915 := Call2(__e, SymFn(symshen_4hds_a_2), let__11500, sym_6_6)


var ifres11903 Obj

if True == tmp11915 {
tmp11904 := Call1(__e, SymFn(symtail), let__11500)


let__11501 := tmp11904
_ = let__11501

tmp11905 := Call1(__e, SymFn(symshen_4_5formula_6), let__11501)


let__11502 := tmp11905
_ = let__11502

tmp11913 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11502)


var ifres11906 Obj

if True == tmp11913 {
tmp11907 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11906 = tmp11907


} else {
tmp11908 := Call1(__e, SymFn(symshen_4_5_1out), let__11502)


let__11503 := tmp11908
_ = let__11503

tmp11909 := Call1(__e, SymFn(symshen_4in_1_6), let__11502)


let__11504 := tmp11909
_ = let__11504

tmp11910 := GCons(__e, let__11503, Nil)

tmp11911 := GCons(__e, let__11499, tmp11910)

tmp11912 := Call2(__e, SymFn(symshen_4comb), let__11504, tmp11911)


ifres11906 = tmp11912


}

ifres11903 = ifres11906


} else {
tmp11914 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11903 = tmp11914


}

ifres11899 = ifres11903


}

let__11497 := ifres11899
_ = let__11497

tmp11929 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11497)


if True == tmp11929 {
tmp11917 := Call1(__e, SymFn(symshen_4_5formula_6), V3521)


let__11506 := tmp11917
_ = let__11506

tmp11925 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11506)


var ifres11918 Obj

if True == tmp11925 {
tmp11919 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11918 = tmp11919


} else {
tmp11920 := Call1(__e, SymFn(symshen_4_5_1out), let__11506)


let__11507 := tmp11920
_ = let__11507

tmp11921 := Call1(__e, SymFn(symshen_4in_1_6), let__11506)


let__11508 := tmp11921
_ = let__11508

tmp11922 := GCons(__e, let__11507, Nil)

tmp11923 := GCons(__e, Nil, tmp11922)

tmp11924 := Call2(__e, SymFn(symshen_4comb), let__11508, tmp11923)


ifres11918 = tmp11924


}

let__11505 := ifres11918
_ = let__11505

tmp11927 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11505)


if True == tmp11927 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11505
}


} else {
return let__11497
}


} else {
return let__11495
}


panic("unreachable")
})

tmp11932 := Call2(__e, ns2_1set, symshen_4_5prem_6, tmp11892)


_ = tmp11932

tmp11933 := MakeNative1(func(__e *ControlFlow, V3536 Obj) Obj {
tmp11934 := Call1(__e, SymFn(symshen_4_5formula_6), V3536)


let__11510 := tmp11934
_ = let__11510

tmp11952 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11510)


var ifres11935 Obj

if True == tmp11952 {
tmp11936 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11935 = tmp11936


} else {
tmp11937 := Call1(__e, SymFn(symshen_4_5_1out), let__11510)


let__11511 := tmp11937
_ = let__11511

tmp11938 := Call1(__e, SymFn(symshen_4in_1_6), let__11510)


let__11512 := tmp11938
_ = let__11512

tmp11939 := Call1(__e, SymFn(symshen_4_5iscomma_6), let__11512)


let__11513 := tmp11939
_ = let__11513

tmp11951 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11513)


var ifres11940 Obj

if True == tmp11951 {
tmp11941 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11940 = tmp11941


} else {
tmp11942 := Call1(__e, SymFn(symshen_4in_1_6), let__11513)


let__11514 := tmp11942
_ = let__11514

tmp11943 := Call1(__e, SymFn(symshen_4_5ass_6), let__11514)


let__11515 := tmp11943
_ = let__11515

tmp11950 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11515)


var ifres11944 Obj

if True == tmp11950 {
tmp11945 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11944 = tmp11945


} else {
tmp11946 := Call1(__e, SymFn(symshen_4_5_1out), let__11515)


let__11516 := tmp11946
_ = let__11516

tmp11947 := Call1(__e, SymFn(symshen_4in_1_6), let__11515)


let__11517 := tmp11947
_ = let__11517

tmp11948 := GCons(__e, let__11511, let__11516)

tmp11949 := Call2(__e, SymFn(symshen_4comb), let__11517, tmp11948)


ifres11944 = tmp11949


}

ifres11940 = ifres11944


}

ifres11935 = ifres11940


}

let__11509 := ifres11935
_ = let__11509

tmp11972 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11509)


if True == tmp11972 {
tmp11953 := Call1(__e, SymFn(symshen_4_5formula_6), V3536)


let__11519 := tmp11953
_ = let__11519

tmp11960 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11519)


var ifres11954 Obj

if True == tmp11960 {
tmp11955 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11954 = tmp11955


} else {
tmp11956 := Call1(__e, SymFn(symshen_4_5_1out), let__11519)


let__11520 := tmp11956
_ = let__11520

tmp11957 := Call1(__e, SymFn(symshen_4in_1_6), let__11519)


let__11521 := tmp11957
_ = let__11521

tmp11958 := GCons(__e, let__11520, Nil)

tmp11959 := Call2(__e, SymFn(symshen_4comb), let__11521, tmp11958)


ifres11954 = tmp11959


}

let__11518 := ifres11954
_ = let__11518

tmp11970 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11518)


if True == tmp11970 {
tmp11961 := Call1(__e, SymFn(sym_5e_6), V3536)


let__11523 := tmp11961
_ = let__11523

tmp11966 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11523)


var ifres11962 Obj

if True == tmp11966 {
tmp11963 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11962 = tmp11963


} else {
tmp11964 := Call1(__e, SymFn(symshen_4in_1_6), let__11523)


let__11524 := tmp11964
_ = let__11524

tmp11965 := Call2(__e, SymFn(symshen_4comb), let__11524, Nil)


ifres11962 = tmp11965


}

let__11522 := ifres11962
_ = let__11522

tmp11968 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11522)


if True == tmp11968 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11522
}


} else {
return let__11518
}


} else {
return let__11509
}


panic("unreachable")
})

tmp11973 := Call2(__e, ns2_1set, symshen_4_5ass_6, tmp11933)


_ = tmp11973

tmp11974 := MakeNative1(func(__e *ControlFlow, V3553 Obj) Obj {
tmp11984 := GConsP(__e, V3553)

var ifres11975 Obj

if True == tmp11984 {
tmp11976 := Call1(__e, SymFn(symhead), V3553)


let__11526 := tmp11976
_ = let__11526

tmp11977 := Call1(__e, SymFn(symtail), V3553)


let__11527 := tmp11977
_ = let__11527

tmp11981 := PrimGuard1(__e, symintern, PrimIntern, MakeString(","))

tmp11982 := GEq(__e, let__11526, tmp11981)

var ifres11978 Obj

if True == tmp11982 {
tmp11979 := Call2(__e, SymFn(symshen_4comb), let__11527, symshen_4skip)


ifres11978 = tmp11979


} else {
tmp11980 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11978 = tmp11980


}

ifres11975 = ifres11978


} else {
tmp11983 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11975 = tmp11983


}

let__11525 := ifres11975
_ = let__11525

tmp11986 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11525)


if True == tmp11986 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11525
}


panic("unreachable")
})

tmp11987 := Call2(__e, ns2_1set, symshen_4_5iscomma_6, tmp11974)


_ = tmp11987

tmp11988 := MakeNative1(func(__e *ControlFlow, V3557 Obj) Obj {
tmp11989 := Call1(__e, SymFn(symshen_4_5expr_6), V3557)


let__11529 := tmp11989
_ = let__11529

tmp12012 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11529)


var ifres11990 Obj

if True == tmp12012 {
tmp11991 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11990 = tmp11991


} else {
tmp11992 := Call1(__e, SymFn(symshen_4_5_1out), let__11529)


let__11530 := tmp11992
_ = let__11530

tmp11993 := Call1(__e, SymFn(symshen_4in_1_6), let__11529)


let__11531 := tmp11993
_ = let__11531

tmp11994 := Call1(__e, SymFn(symshen_4_5iscolon_6), let__11531)


let__11532 := tmp11994
_ = let__11532

tmp12011 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11532)


var ifres11995 Obj

if True == tmp12011 {
tmp11996 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11995 = tmp11996


} else {
tmp11997 := Call1(__e, SymFn(symshen_4in_1_6), let__11532)


let__11533 := tmp11997
_ = let__11533

tmp11998 := Call1(__e, SymFn(symshen_4_5type_6), let__11533)


let__11534 := tmp11998
_ = let__11534

tmp12010 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11534)


var ifres11999 Obj

if True == tmp12010 {
tmp12000 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres11999 = tmp12000


} else {
tmp12001 := Call1(__e, SymFn(symshen_4_5_1out), let__11534)


let__11535 := tmp12001
_ = let__11535

tmp12002 := Call1(__e, SymFn(symshen_4in_1_6), let__11534)


let__11536 := tmp12002
_ = let__11536

tmp12003 := Call1(__e, SymFn(symshen_4curry), let__11530)


tmp12004 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp12005 := Call1(__e, SymFn(symshen_4rectify_1type), let__11535)


tmp12006 := GCons(__e, tmp12005, Nil)

tmp12007 := GCons(__e, tmp12004, tmp12006)

tmp12008 := GCons(__e, tmp12003, tmp12007)

tmp12009 := Call2(__e, SymFn(symshen_4comb), let__11536, tmp12008)


ifres11999 = tmp12009


}

ifres11995 = ifres11999


}

ifres11990 = ifres11995


}

let__11528 := ifres11990
_ = let__11528

tmp12023 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11528)


if True == tmp12023 {
tmp12013 := Call1(__e, SymFn(symshen_4_5expr_6), V3557)


let__11538 := tmp12013
_ = let__11538

tmp12019 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11538)


var ifres12014 Obj

if True == tmp12019 {
tmp12015 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12014 = tmp12015


} else {
tmp12016 := Call1(__e, SymFn(symshen_4_5_1out), let__11538)


let__11539 := tmp12016
_ = let__11539

tmp12017 := Call1(__e, SymFn(symshen_4in_1_6), let__11538)


let__11540 := tmp12017
_ = let__11540

tmp12018 := Call2(__e, SymFn(symshen_4comb), let__11540, let__11539)


ifres12014 = tmp12018


}

let__11537 := ifres12014
_ = let__11537

tmp12021 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11537)


if True == tmp12021 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11537
}


} else {
return let__11528
}


panic("unreachable")
})

tmp12024 := Call2(__e, ns2_1set, symshen_4_5formula_6, tmp11988)


_ = tmp12024

tmp12025 := MakeNative1(func(__e *ControlFlow, V3571 Obj) Obj {
tmp12035 := GConsP(__e, V3571)

var ifres12026 Obj

if True == tmp12035 {
tmp12027 := Call1(__e, SymFn(symhead), V3571)


let__11542 := tmp12027
_ = let__11542

tmp12028 := Call1(__e, SymFn(symtail), V3571)


let__11543 := tmp12028
_ = let__11543

tmp12032 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp12033 := GEq(__e, let__11542, tmp12032)

var ifres12029 Obj

if True == tmp12033 {
tmp12030 := Call2(__e, SymFn(symshen_4comb), let__11543, symshen_4skip)


ifres12029 = tmp12030


} else {
tmp12031 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12029 = tmp12031


}

ifres12026 = ifres12029


} else {
tmp12034 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12026 = tmp12034


}

let__11541 := ifres12026
_ = let__11541

tmp12037 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11541)


if True == tmp12037 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11541
}


panic("unreachable")
})

tmp12038 := Call2(__e, ns2_1set, symshen_4_5iscolon_6, tmp12025)


_ = tmp12038

tmp12039 := MakeNative1(func(__e *ControlFlow, V3575 Obj) Obj {
tmp12040 := Call1(__e, SymFn(symshen_4_5side_6), V3575)


let__11545 := tmp12040
_ = let__11545

tmp12053 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11545)


var ifres12041 Obj

if True == tmp12053 {
tmp12042 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12041 = tmp12042


} else {
tmp12043 := Call1(__e, SymFn(symshen_4_5_1out), let__11545)


let__11546 := tmp12043
_ = let__11546

tmp12044 := Call1(__e, SymFn(symshen_4in_1_6), let__11545)


let__11547 := tmp12044
_ = let__11547

tmp12045 := Call1(__e, SymFn(symshen_4_5sides_6), let__11547)


let__11548 := tmp12045
_ = let__11548

tmp12052 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11548)


var ifres12046 Obj

if True == tmp12052 {
tmp12047 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12046 = tmp12047


} else {
tmp12048 := Call1(__e, SymFn(symshen_4_5_1out), let__11548)


let__11549 := tmp12048
_ = let__11549

tmp12049 := Call1(__e, SymFn(symshen_4in_1_6), let__11548)


let__11550 := tmp12049
_ = let__11550

tmp12050 := GCons(__e, let__11546, let__11549)

tmp12051 := Call2(__e, SymFn(symshen_4comb), let__11550, tmp12050)


ifres12046 = tmp12051


}

ifres12041 = ifres12046


}

let__11544 := ifres12041
_ = let__11544

tmp12063 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11544)


if True == tmp12063 {
tmp12054 := Call1(__e, SymFn(sym_5e_6), V3575)


let__11552 := tmp12054
_ = let__11552

tmp12059 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11552)


var ifres12055 Obj

if True == tmp12059 {
tmp12056 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12055 = tmp12056


} else {
tmp12057 := Call1(__e, SymFn(symshen_4in_1_6), let__11552)


let__11553 := tmp12057
_ = let__11553

tmp12058 := Call2(__e, SymFn(symshen_4comb), let__11553, Nil)


ifres12055 = tmp12058


}

let__11551 := ifres12055
_ = let__11551

tmp12061 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11551)


if True == tmp12061 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11551
}


} else {
return let__11544
}


panic("unreachable")
})

tmp12064 := Call2(__e, ns2_1set, symshen_4_5sides_6, tmp12039)


_ = tmp12064

tmp12065 := MakeNative1(func(__e *ControlFlow, V3586 Obj) Obj {
tmp12077 := Call2(__e, SymFn(symshen_4hds_a_2), V3586, symif)


var ifres12066 Obj

if True == tmp12077 {
tmp12067 := Call1(__e, SymFn(symtail), V3586)


let__11555 := tmp12067
_ = let__11555

tmp12075 := GConsP(__e, let__11555)

var ifres12068 Obj

if True == tmp12075 {
tmp12069 := Call1(__e, SymFn(symhead), let__11555)


let__11556 := tmp12069
_ = let__11556

tmp12070 := Call1(__e, SymFn(symtail), let__11555)


let__11557 := tmp12070
_ = let__11557

tmp12071 := GCons(__e, let__11556, Nil)

tmp12072 := GCons(__e, symif, tmp12071)

tmp12073 := Call2(__e, SymFn(symshen_4comb), let__11557, tmp12072)


ifres12068 = tmp12073


} else {
tmp12074 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12068 = tmp12074


}

ifres12066 = ifres12068


} else {
tmp12076 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12066 = tmp12076


}

let__11554 := ifres12066
_ = let__11554

tmp12133 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11554)


if True == tmp12133 {
tmp12095 := Call2(__e, SymFn(symshen_4hds_a_2), V3586, symlet)


var ifres12078 Obj

if True == tmp12095 {
tmp12079 := Call1(__e, SymFn(symtail), V3586)


let__11559 := tmp12079
_ = let__11559

tmp12093 := GConsP(__e, let__11559)

var ifres12080 Obj

if True == tmp12093 {
tmp12081 := Call1(__e, SymFn(symhead), let__11559)


let__11560 := tmp12081
_ = let__11560

tmp12082 := Call1(__e, SymFn(symtail), let__11559)


let__11561 := tmp12082
_ = let__11561

tmp12091 := GConsP(__e, let__11561)

var ifres12083 Obj

if True == tmp12091 {
tmp12084 := Call1(__e, SymFn(symhead), let__11561)


let__11562 := tmp12084
_ = let__11562

tmp12085 := Call1(__e, SymFn(symtail), let__11561)


let__11563 := tmp12085
_ = let__11563

tmp12086 := GCons(__e, let__11562, Nil)

tmp12087 := GCons(__e, let__11560, tmp12086)

tmp12088 := GCons(__e, symlet, tmp12087)

tmp12089 := Call2(__e, SymFn(symshen_4comb), let__11563, tmp12088)


ifres12083 = tmp12089


} else {
tmp12090 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12083 = tmp12090


}

ifres12080 = ifres12083


} else {
tmp12092 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12080 = tmp12092


}

ifres12078 = ifres12080


} else {
tmp12094 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12078 = tmp12094


}

let__11558 := ifres12078
_ = let__11558

tmp12131 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11558)


if True == tmp12131 {
tmp12110 := Call2(__e, SymFn(symshen_4hds_a_2), V3586, symctxt)


var ifres12096 Obj

if True == tmp12110 {
tmp12097 := Call1(__e, SymFn(symtail), V3586)


let__11565 := tmp12097
_ = let__11565

tmp12108 := GConsP(__e, let__11565)

var ifres12098 Obj

if True == tmp12108 {
tmp12099 := Call1(__e, SymFn(symhead), let__11565)


let__11566 := tmp12099
_ = let__11566

tmp12100 := Call1(__e, SymFn(symtail), let__11565)


let__11567 := tmp12100
_ = let__11567

tmp12106 := PrimGuard1(__e, symvariable_2, PrimIsVariable, let__11566)

var ifres12101 Obj

if True == tmp12106 {
tmp12102 := GCons(__e, let__11566, Nil)

tmp12103 := GCons(__e, symctxt, tmp12102)

tmp12104 := Call2(__e, SymFn(symshen_4comb), let__11567, tmp12103)


ifres12101 = tmp12104


} else {
tmp12105 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12101 = tmp12105


}

ifres12098 = ifres12101


} else {
tmp12107 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12098 = tmp12107


}

ifres12096 = ifres12098


} else {
tmp12109 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12096 = tmp12109


}

let__11564 := ifres12096
_ = let__11564

tmp12129 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11564)


if True == tmp12129 {
tmp12125 := Call2(__e, SymFn(symshen_4hds_a_2), V3586, symsqts)


var ifres12111 Obj

if True == tmp12125 {
tmp12112 := Call1(__e, SymFn(symtail), V3586)


let__11569 := tmp12112
_ = let__11569

tmp12123 := GConsP(__e, let__11569)

var ifres12113 Obj

if True == tmp12123 {
tmp12114 := Call1(__e, SymFn(symhead), let__11569)


let__11570 := tmp12114
_ = let__11570

tmp12115 := Call1(__e, SymFn(symtail), let__11569)


let__11571 := tmp12115
_ = let__11571

tmp12121 := PrimGuard1(__e, symvariable_2, PrimIsVariable, let__11570)

var ifres12116 Obj

if True == tmp12121 {
tmp12117 := GCons(__e, let__11570, Nil)

tmp12118 := GCons(__e, symsqts, tmp12117)

tmp12119 := Call2(__e, SymFn(symshen_4comb), let__11571, tmp12118)


ifres12116 = tmp12119


} else {
tmp12120 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12116 = tmp12120


}

ifres12113 = ifres12116


} else {
tmp12122 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12113 = tmp12122


}

ifres12111 = ifres12113


} else {
tmp12124 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12111 = tmp12124


}

let__11568 := ifres12111
_ = let__11568

tmp12127 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11568)


if True == tmp12127 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11568
}


} else {
return let__11564
}


} else {
return let__11558
}


} else {
return let__11554
}


panic("unreachable")
})

tmp12134 := Call2(__e, ns2_1set, symshen_4_5side_6, tmp12065)


_ = tmp12134

tmp12135 := MakeNative3(func(__e *ControlFlow, V3611, V3612, V3613 Obj) Obj {
tmp12162 := GConsP(__e, V3613)

var ifres12152 Obj

if True == tmp12162 {
tmp12160 := GHd(__e, V3613)

tmp12161 := GEq(__e, Nil, tmp12160)

var ifres12153 Obj

if True == tmp12161 {
tmp12158 := GTl(__e, V3613)

tmp12159 := GConsP(__e, tmp12158)

var ifres12154 Obj

if True == tmp12159 {
tmp12155 := GTl(__e, V3613)

tmp12156 := GTl(__e, tmp12155)

tmp12157 := GEq(__e, Nil, tmp12156)

ifres12154 = tmp12157


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

if True == ifres12152 {
tmp12136 := Call1(__e, SymFn(symgensym), symP)


let__11572 := tmp12136
_ = let__11572

tmp12137 := GTl(__e, V3613)

tmp12138 := GCons(__e, let__11572, Nil)

tmp12139 := GCons(__e, tmp12137, tmp12138)

let__11573 := tmp12139
_ = let__11573

tmp12140 := Call1(__e, SymFn(symshen_4coll_1formulae), V3612)


tmp12141 := GCons(__e, let__11572, Nil)

tmp12142 := GCons(__e, tmp12140, tmp12141)

let__11574 := tmp12142
_ = let__11574

tmp12143 := GCons(__e, let__11574, Nil)

tmp12144 := GCons(__e, let__11573, Nil)

tmp12145 := GCons(__e, tmp12143, tmp12144)

tmp12146 := GCons(__e, V3611, tmp12145)

let__11575 := tmp12146
_ = let__11575

tmp12147 := GCons(__e, V3613, Nil)

tmp12148 := GCons(__e, V3612, tmp12147)

tmp12149 := GCons(__e, V3611, tmp12148)

let__11576 := tmp12149
_ = let__11576

tmp12150 := GCons(__e, let__11575, Nil)

return GCons(__e, let__11576, tmp12150)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.lr-rule"))
}


panic("unreachable")
})

tmp12163 := Call2(__e, ns2_1set, symshen_4lr_1rule, tmp12135)


_ = tmp12163

tmp12164 := MakeNative1(func(__e *ControlFlow, V3621 Obj) Obj {
tmp12189 := GEq(__e, Nil, V3621)

if True == tmp12189 {
return Nil
} else {
tmp12187 := GConsP(__e, V3621)

var ifres12171 Obj

if True == tmp12187 {
tmp12185 := GHd(__e, V3621)

tmp12186 := GConsP(__e, tmp12185)

var ifres12172 Obj

if True == tmp12186 {
tmp12182 := GHd(__e, V3621)

tmp12183 := GHd(__e, tmp12182)

tmp12184 := GEq(__e, Nil, tmp12183)

var ifres12173 Obj

if True == tmp12184 {
tmp12179 := GHd(__e, V3621)

tmp12180 := GTl(__e, tmp12179)

tmp12181 := GConsP(__e, tmp12180)

var ifres12174 Obj

if True == tmp12181 {
tmp12175 := GHd(__e, V3621)

tmp12176 := GTl(__e, tmp12175)

tmp12177 := GTl(__e, tmp12176)

tmp12178 := GEq(__e, Nil, tmp12177)

ifres12174 = tmp12178


} else {
ifres12174 = False


}

ifres12173 = ifres12174


} else {
ifres12173 = False


}

ifres12172 = ifres12173


} else {
ifres12172 = False


}

ifres12171 = ifres12172


} else {
ifres12171 = False


}

if True == ifres12171 {
tmp12165 := GHd(__e, V3621)

tmp12166 := GTl(__e, tmp12165)

tmp12167 := GHd(__e, tmp12166)

tmp12168 := GTl(__e, V3621)

tmp12169 := Call1(__e, SymFn(symshen_4coll_1formulae), tmp12168)


return GCons(__e, tmp12167, tmp12169)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.coll-formulae"))
}


}


panic("unreachable")
})

tmp12190 := Call2(__e, ns2_1set, symshen_4coll_1formulae, tmp12164)


_ = tmp12190

tmp12191 := MakeNative1(func(__e *ControlFlow, V3622 Obj) Obj {
tmp12202 := GConsP(__e, V3622)

var ifres12192 Obj

if True == tmp12202 {
tmp12193 := Call1(__e, SymFn(symhead), V3622)


let__11578 := tmp12193
_ = let__11578

tmp12194 := Call1(__e, SymFn(symtail), V3622)


let__11579 := tmp12194
_ = let__11579

tmp12199 := Call1(__e, SymFn(symshen_4key_1in_1sequent_1calculus_2), let__11578)


tmp12200 := GNot(__e, tmp12199)

var ifres12195 Obj

if True == tmp12200 {
tmp12196 := Call1(__e, SymFn(symmacroexpand), let__11578)


tmp12197 := Call2(__e, SymFn(symshen_4comb), let__11579, tmp12196)


ifres12195 = tmp12197


} else {
tmp12198 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12195 = tmp12198


}

ifres12192 = ifres12195


} else {
tmp12201 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12192 = tmp12201


}

let__11577 := ifres12192
_ = let__11577

tmp12204 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11577)


if True == tmp12204 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11577
}


panic("unreachable")
})

tmp12205 := Call2(__e, ns2_1set, symshen_4_5expr_6, tmp12191)


_ = tmp12205

tmp12206 := MakeNative1(func(__e *ControlFlow, V3626 Obj) Obj {
tmp12210 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp12211 := PrimGuard1(__e, symintern, PrimIntern, MakeString(","))

tmp12212 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp12213 := GCons(__e, sym_5_1_1, Nil)

tmp12214 := GCons(__e, tmp12212, tmp12213)

tmp12215 := GCons(__e, tmp12211, tmp12214)

tmp12216 := GCons(__e, tmp12210, tmp12215)

tmp12217 := GCons(__e, sym_6_6, tmp12216)

tmp12218 := Call2(__e, SymFn(symelement_2), V3626, tmp12217)


if True == tmp12218 {
return True
} else {
tmp12208 := Call1(__e, SymFn(symshen_4sng_2), V3626)


if True == tmp12208 {
return True
} else {
return TailCall1(__e, SymFn(symshen_4dbl_2), V3626)
}


}


panic("unreachable")
})

tmp12219 := Call2(__e, ns2_1set, symshen_4key_1in_1sequent_1calculus_2, tmp12206)


_ = tmp12219

tmp12220 := MakeNative1(func(__e *ControlFlow, V3627 Obj) Obj {
tmp12221 := Call1(__e, SymFn(symshen_4_5expr_6), V3627)


let__11581 := tmp12221
_ = let__11581

tmp12227 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11581)


var ifres12222 Obj

if True == tmp12227 {
tmp12223 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12222 = tmp12223


} else {
tmp12224 := Call1(__e, SymFn(symshen_4_5_1out), let__11581)


let__11582 := tmp12224
_ = let__11582

tmp12225 := Call1(__e, SymFn(symshen_4in_1_6), let__11581)


let__11583 := tmp12225
_ = let__11583

tmp12226 := Call2(__e, SymFn(symshen_4comb), let__11583, let__11582)


ifres12222 = tmp12226


}

let__11580 := ifres12222
_ = let__11580

tmp12229 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11580)


if True == tmp12229 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11580
}


panic("unreachable")
})

tmp12230 := Call2(__e, ns2_1set, symshen_4_5type_6, tmp12220)


_ = tmp12230

tmp12231 := MakeNative1(func(__e *ControlFlow, V3632 Obj) Obj {
tmp12240 := GConsP(__e, V3632)

var ifres12232 Obj

if True == tmp12240 {
tmp12233 := Call1(__e, SymFn(symhead), V3632)


let__11585 := tmp12233
_ = let__11585

tmp12234 := Call1(__e, SymFn(symtail), V3632)


let__11586 := tmp12234
_ = let__11586

tmp12238 := Call1(__e, SymFn(symshen_4dbl_2), let__11585)


var ifres12235 Obj

if True == tmp12238 {
tmp12236 := Call2(__e, SymFn(symshen_4comb), let__11586, let__11585)


ifres12235 = tmp12236


} else {
tmp12237 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12235 = tmp12237


}

ifres12232 = ifres12235


} else {
tmp12239 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12232 = tmp12239


}

let__11584 := ifres12232
_ = let__11584

tmp12242 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11584)


if True == tmp12242 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11584
}


panic("unreachable")
})

tmp12243 := Call2(__e, ns2_1set, symshen_4_5dbl_6, tmp12231)


_ = tmp12243

tmp12244 := MakeNative1(func(__e *ControlFlow, V3636 Obj) Obj {
tmp12253 := GConsP(__e, V3636)

var ifres12245 Obj

if True == tmp12253 {
tmp12246 := Call1(__e, SymFn(symhead), V3636)


let__11588 := tmp12246
_ = let__11588

tmp12247 := Call1(__e, SymFn(symtail), V3636)


let__11589 := tmp12247
_ = let__11589

tmp12251 := Call1(__e, SymFn(symshen_4sng_2), let__11588)


var ifres12248 Obj

if True == tmp12251 {
tmp12249 := Call2(__e, SymFn(symshen_4comb), let__11589, let__11588)


ifres12248 = tmp12249


} else {
tmp12250 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12248 = tmp12250


}

ifres12245 = ifres12248


} else {
tmp12252 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres12245 = tmp12252


}

let__11587 := ifres12245
_ = let__11587

tmp12255 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__11587)


if True == tmp12255 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__11587
}


panic("unreachable")
})

tmp12256 := Call2(__e, ns2_1set, symshen_4_5sng_6, tmp12244)


_ = tmp12256

tmp12257 := MakeNative1(func(__e *ControlFlow, V3640 Obj) Obj {
tmp12260 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V3640)

if True == tmp12260 {
tmp12258 := PrimGuard1(__e, symstr, PrimStr, V3640)

return TailCall1(__e, SymFn(symshen_4sng_1h_2), tmp12258)


} else {
return False
}


panic("unreachable")
})

tmp12261 := Call2(__e, ns2_1set, symshen_4sng_2, tmp12257)


_ = tmp12261

tmp12262 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V3643 := __self1
_ = V3643
tmp12270 := GEq(__e, MakeString("___"), V3643)

if True == tmp12270 {
return True
} else {
tmp12268 := Call1(__e, SymFn(symshen_4_7string_2), V3643)


var ifres12265 Obj

if True == tmp12268 {
tmp12266 := Call1(__e, SymFn(symhdstr), V3643)


tmp12267 := GEq(__e, MakeString("_"), tmp12266)

ifres12265 = tmp12267


} else {
ifres12265 = False


}

if True == ifres12265 {
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

tmp12271 := Call2(__e, ns2_1set, symshen_4sng_1h_2, tmp12262)


_ = tmp12271

tmp12272 := MakeNative1(func(__e *ControlFlow, V3644 Obj) Obj {
tmp12275 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V3644)

if True == tmp12275 {
tmp12273 := PrimGuard1(__e, symstr, PrimStr, V3644)

return TailCall1(__e, SymFn(symshen_4dbl_1h_2), tmp12273)


} else {
return False
}


panic("unreachable")
})

tmp12276 := Call2(__e, ns2_1set, symshen_4dbl_2, tmp12272)


_ = tmp12276

tmp12277 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V3647 := __self1
_ = V3647
tmp12285 := GEq(__e, MakeString("==="), V3647)

if True == tmp12285 {
return True
} else {
tmp12283 := Call1(__e, SymFn(symshen_4_7string_2), V3647)


var ifres12280 Obj

if True == tmp12283 {
tmp12281 := Call1(__e, SymFn(symhdstr), V3647)


tmp12282 := GEq(__e, MakeString("="), tmp12281)

ifres12280 = tmp12282


} else {
ifres12280 = False


}

if True == ifres12280 {
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

tmp12286 := Call2(__e, ns2_1set, symshen_4dbl_1h_2, tmp12277)


_ = tmp12286

tmp12287 := MakeNative2(func(__e *ControlFlow, V3648, V3649 Obj) Obj {
tmp12288 := MakeNative1(func(__e *ControlFlow, Z3651 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4rule_1_6clause), Z3651)
panic("unreachable")
})

tmp12289 := Call2(__e, SymFn(symmapcan), tmp12288, V3649)


let__11590 := tmp12289
_ = let__11590

tmp12290 := GCons(__e, V3648, let__11590)

tmp12291 := GCons(__e, symdefprolog, tmp12290)

let__11591 := tmp12291
_ = let__11591

return TailCall1(__e, SymFn(symeval), let__11591)


panic("unreachable")
})

tmp12292 := Call2(__e, ns2_1set, symshen_4rules_1_6prolog, tmp12287)


_ = tmp12292

tmp12293 := MakeNative1(func(__e *ControlFlow, V3653 Obj) Obj {
tmp12347 := GConsP(__e, V3653)

var ifres12317 Obj

if True == tmp12347 {
tmp12345 := GTl(__e, V3653)

tmp12346 := GConsP(__e, tmp12345)

var ifres12318 Obj

if True == tmp12346 {
tmp12342 := GTl(__e, V3653)

tmp12343 := GTl(__e, tmp12342)

tmp12344 := GConsP(__e, tmp12343)

var ifres12319 Obj

if True == tmp12344 {
tmp12338 := GTl(__e, V3653)

tmp12339 := GTl(__e, tmp12338)

tmp12340 := GHd(__e, tmp12339)

tmp12341 := GConsP(__e, tmp12340)

var ifres12320 Obj

if True == tmp12341 {
tmp12333 := GTl(__e, V3653)

tmp12334 := GTl(__e, tmp12333)

tmp12335 := GHd(__e, tmp12334)

tmp12336 := GTl(__e, tmp12335)

tmp12337 := GConsP(__e, tmp12336)

var ifres12321 Obj

if True == tmp12337 {
tmp12327 := GTl(__e, V3653)

tmp12328 := GTl(__e, tmp12327)

tmp12329 := GHd(__e, tmp12328)

tmp12330 := GTl(__e, tmp12329)

tmp12331 := GTl(__e, tmp12330)

tmp12332 := GEq(__e, Nil, tmp12331)

var ifres12322 Obj

if True == tmp12332 {
tmp12323 := GTl(__e, V3653)

tmp12324 := GTl(__e, tmp12323)

tmp12325 := GTl(__e, tmp12324)

tmp12326 := GEq(__e, Nil, tmp12325)

ifres12322 = tmp12326


} else {
ifres12322 = False


}

ifres12321 = ifres12322


} else {
ifres12321 = False


}

ifres12320 = ifres12321


} else {
ifres12320 = False


}

ifres12319 = ifres12320


} else {
ifres12319 = False


}

ifres12318 = ifres12319


} else {
ifres12318 = False


}

ifres12317 = ifres12318


} else {
ifres12317 = False


}

if True == ifres12317 {
tmp12294 := GTl(__e, V3653)

tmp12295 := GTl(__e, tmp12294)

tmp12296 := GHd(__e, tmp12295)

tmp12297 := GTl(__e, tmp12296)

tmp12298 := GHd(__e, tmp12297)

tmp12299 := Call1(__e, SymFn(symshen_4extract_1vars), tmp12298)


let__11592 := tmp12299
_ = let__11592

tmp12300 := GTl(__e, V3653)

tmp12301 := GTl(__e, tmp12300)

tmp12302 := GHd(__e, tmp12301)

tmp12303 := GTl(__e, tmp12302)

tmp12304 := GHd(__e, tmp12303)

tmp12305 := Call1(__e, SymFn(symshen_4rule_1_6head), tmp12304)


tmp12306 := GCons(__e, sym_5_1_1, Nil)

tmp12307 := GHd(__e, V3653)

tmp12308 := GTl(__e, V3653)

tmp12309 := GHd(__e, tmp12308)

tmp12310 := GTl(__e, V3653)

tmp12311 := GTl(__e, tmp12310)

tmp12312 := GHd(__e, tmp12311)

tmp12313 := GHd(__e, tmp12312)

tmp12314 := Call5(__e, SymFn(symshen_4rule_1_6body), let__11592, symAssumptions, tmp12307, tmp12309, tmp12313)


tmp12315 := Call2(__e, SymFn(symappend), tmp12306, tmp12314)


return TailCall2(__e, SymFn(symappend), tmp12305, tmp12315)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4rule_1_6clause)
}


panic("unreachable")
})

tmp12348 := Call2(__e, ns2_1set, symshen_4rule_1_6clause, tmp12293)


_ = tmp12348

tmp12349 := MakeNative1(func(__e *ControlFlow, V3655 Obj) Obj {
tmp12350 := Call1(__e, SymFn(symshen_4macro_1_8ch), V3655)


tmp12351 := GCons(__e, symAssumptions, Nil)

return GCons(__e, tmp12350, tmp12351)


panic("unreachable")
})

tmp12352 := Call2(__e, ns2_1set, symshen_4rule_1_6head, tmp12349)


_ = tmp12352

tmp12353 := MakeNative1(func(__e *ControlFlow, V3656 Obj) Obj {
tmp12354 := GCons(__e, V3656, Nil)

return GCons(__e, symshen_4_8ch, tmp12354)


panic("unreachable")
})

tmp12355 := Call2(__e, ns2_1set, symshen_4macro_1_8ch, tmp12353)


_ = tmp12355

tmp12356 := MakeNative1(func(__e *ControlFlow, V3657 Obj) Obj {
tmp12357 := GCons(__e, V3657, Nil)

return GCons(__e, symshen_4_8c, tmp12357)


panic("unreachable")
})

tmp12358 := Call2(__e, ns2_1set, symshen_4macro_1_8c, tmp12356)


_ = tmp12358

tmp12359 := MakeNative5(func(__e *ControlFlow, V3658, V3659, V3660, V3661, V3662 Obj) Obj {
tmp12387 := GEq(__e, Nil, V3662)

if True == tmp12387 {
return TailCall5(__e, SymFn(symshen_4side_1conditions_1_6goals), Nil, V3658, V3659, V3660, V3661)
} else {
tmp12385 := GEq(__e, Nil, V3661)

var ifres12380 Obj

if True == tmp12385 {
tmp12384 := GConsP(__e, V3662)

var ifres12381 Obj

if True == tmp12384 {
tmp12382 := GTl(__e, V3662)

tmp12383 := GEq(__e, Nil, tmp12382)

ifres12381 = tmp12383


} else {
ifres12381 = False


}

ifres12380 = ifres12381


} else {
ifres12380 = False


}

if True == ifres12380 {
tmp12360 := GHd(__e, V3662)

tmp12361 := Call2(__e, SymFn(symshen_4passive_1variables), tmp12360, V3658)


let__11593 := tmp12361
_ = let__11593

tmp12362 := GHd(__e, V3662)

tmp12363 := Call2(__e, SymFn(symshen_4remove_1bystanders), V3658, tmp12362)


let__11594 := tmp12363
_ = let__11594

tmp12364 := GHd(__e, V3662)

tmp12365 := Call4(__e, SymFn(symshen_4specialise_1member), tmp12364, V3659, let__11594, let__11593)


tmp12366 := Call5(__e, SymFn(symshen_4side_1conditions_1_6goals), Nil, V3658, V3659, V3660, Nil)


return GCons(__e, tmp12365, tmp12366)


} else {
tmp12378 := GConsP(__e, V3662)

if True == tmp12378 {
tmp12367 := Call1(__e, SymFn(symgensym), symNewAssumptions)


let__11595 := tmp12367
_ = let__11595

tmp12368 := GHd(__e, V3662)

tmp12369 := Call2(__e, SymFn(symshen_4passive_1variables), tmp12368, V3658)


let__11596 := tmp12369
_ = let__11596

tmp12370 := GHd(__e, V3662)

tmp12371 := Call2(__e, SymFn(symshen_4remove_1bystanders), V3658, tmp12370)


let__11597 := tmp12371
_ = let__11597

tmp12372 := GHd(__e, V3662)

tmp12373 := Call5(__e, SymFn(symshen_4specialise_1consume), tmp12372, V3659, let__11597, let__11596, let__11595)


tmp12374 := Call2(__e, SymFn(symappend), V3658, let__11596)


tmp12375 := GTl(__e, V3662)

tmp12376 := Call5(__e, SymFn(symshen_4rule_1_6body), tmp12374, let__11595, V3660, V3661, tmp12375)


return GCons(__e, tmp12373, tmp12376)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4rule_1_6body)
}


}


}


panic("unreachable")
})

tmp12388 := Call2(__e, ns2_1set, symshen_4rule_1_6body, tmp12359)


_ = tmp12388

tmp12389 := MakeNative4(func(__e *ControlFlow, V3668, V3669, V3670, V3671 Obj) Obj {
tmp12390 := Call1(__e, SymFn(symgensym), symshen_4member)


let__11598 := tmp12390
_ = let__11598

tmp12391 := Call4(__e, SymFn(symshen_4member_1clause), let__11598, V3668, V3670, V3671)


let__11599 := tmp12391
_ = let__11599

tmp12392 := Call2(__e, SymFn(symappend), V3670, V3671)


tmp12393 := GCons(__e, V3669, tmp12392)

return GCons(__e, let__11598, tmp12393)


panic("unreachable")
})

tmp12394 := Call2(__e, ns2_1set, symshen_4specialise_1member, tmp12389)


_ = tmp12394

tmp12395 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V3676 := __self1
_ = V3676
V3677 := __self2
_ = V3677
tmp12408 := GEq(__e, Nil, V3676)

if True == tmp12408 {
return Nil
} else {
tmp12406 := GConsP(__e, V3676)

var ifres12403 Obj

if True == tmp12406 {
tmp12404 := GHd(__e, V3676)

tmp12405 := Call2(__e, SymFn(symshen_4occurs_1check_2), tmp12404, V3677)


ifres12403 = tmp12405


} else {
ifres12403 = False


}

if True == ifres12403 {
tmp12396 := GHd(__e, V3676)

tmp12397 := GTl(__e, V3676)

tmp12398 := Call2(__e, SymFn(symshen_4remove_1bystanders), tmp12397, V3677)


return GCons(__e, tmp12396, tmp12398)


} else {
tmp12401 := GConsP(__e, V3676)

if True == tmp12401 {
tmp12399 := GTl(__e, V3676)

if SymFn(symshen_4remove_1bystanders) == __self {
__self1, __self2 = tmp12399, V3677
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4remove_1bystanders), tmp12399, V3677)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4remove_1bystanders)
}


}


}


panic("unreachable")
})
return __self
})()

tmp12409 := Call2(__e, ns2_1set, symshen_4remove_1bystanders, tmp12395)


_ = tmp12409

tmp12410 := MakeNative4(func(__e *ControlFlow, V3678, V3679, V3680, V3681 Obj) Obj {
tmp12411 := Call1(__e, SymFn(symlength), V3681)


tmp12412 := Call1(__e, SymFn(symshen_4nvars), tmp12411)


let__11600 := tmp12412
_ = let__11600

tmp12413 := Call1(__e, SymFn(symshen_4macro_1_8ch), V3679)


tmp12414 := GCons(__e, sym__, Nil)

tmp12415 := GCons(__e, tmp12413, tmp12414)

tmp12416 := GCons(__e, symcons, tmp12415)

tmp12417 := GCons(__e, tmp12416, Nil)

tmp12418 := GCons(__e, sym_1, tmp12417)

tmp12419 := GCons(__e, tmp12418, Nil)

tmp12420 := GCons(__e, sym_5_1_1, Nil)

tmp12421 := Call2(__e, SymFn(symshen_4passive_1bind), V3681, let__11600)


tmp12422 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp12423 := GCons(__e, tmp12422, Nil)

tmp12424 := Call2(__e, SymFn(symappend), tmp12421, tmp12423)


tmp12425 := Call2(__e, SymFn(symappend), tmp12420, tmp12424)


tmp12426 := Call2(__e, SymFn(symappend), let__11600, tmp12425)


tmp12427 := Call2(__e, SymFn(symappend), V3680, tmp12426)


tmp12428 := Call2(__e, SymFn(symappend), tmp12419, tmp12427)


let__11601 := tmp12428
_ = let__11601

tmp12429 := Call1(__e, SymFn(symgensym), symHypotheses)


let__11603 := tmp12429
_ = let__11603

tmp12430 := Call2(__e, SymFn(symappend), V3680, V3681)


let__11604 := tmp12430
_ = let__11604

tmp12431 := GCons(__e, let__11603, Nil)

tmp12432 := GCons(__e, sym__, tmp12431)

tmp12433 := GCons(__e, symcons, tmp12432)

tmp12434 := GCons(__e, tmp12433, Nil)

tmp12435 := GCons(__e, sym_1, tmp12434)

tmp12436 := GCons(__e, tmp12435, Nil)

tmp12437 := Call2(__e, SymFn(symappend), tmp12436, let__11604)


let__11605 := tmp12437
_ = let__11605

tmp12438 := GCons(__e, let__11603, let__11604)

tmp12439 := GCons(__e, V3678, tmp12438)

tmp12440 := GCons(__e, tmp12439, Nil)

let__11606 := tmp12440
_ = let__11606

tmp12441 := GCons(__e, sym_5_1_1, Nil)

tmp12442 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp12443 := GCons(__e, tmp12442, Nil)

tmp12444 := Call2(__e, SymFn(symappend), let__11606, tmp12443)


tmp12445 := Call2(__e, SymFn(symappend), tmp12441, tmp12444)


tmp12446 := Call2(__e, SymFn(symappend), let__11605, tmp12445)


let__11602 := tmp12446
_ = let__11602

tmp12447 := Call2(__e, SymFn(symappend), let__11601, let__11602)


tmp12448 := GCons(__e, V3678, tmp12447)

tmp12449 := GCons(__e, symdefprolog, tmp12448)

let__11607 := tmp12449
_ = let__11607

return TailCall1(__e, SymFn(symeval), let__11607)


panic("unreachable")
})

tmp12450 := Call2(__e, ns2_1set, symshen_4member_1clause, tmp12410)


_ = tmp12450

tmp12451 := MakeNative1(func(__e *ControlFlow, V3690 Obj) Obj {
tmp12456 := GEq(__e, MakeInteger(0), V3690)

if True == tmp12456 {
return Nil
} else {
tmp12452 := Call1(__e, SymFn(symgensym), symNewV)


tmp12454 := Call1(__e, SymFn(symshen_4nvars), GSub(__e, V3690, MakeInteger(1)))


return GCons(__e, tmp12452, tmp12454)


}


panic("unreachable")
})

tmp12457 := Call2(__e, ns2_1set, symshen_4nvars, tmp12451)


_ = tmp12457

tmp12458 := MakeNative2(func(__e *ControlFlow, V3691, V3692 Obj) Obj {
tmp12474 := GEq(__e, Nil, V3691)

var ifres12472 Obj

if True == tmp12474 {
tmp12473 := GEq(__e, Nil, V3692)

ifres12472 = tmp12473


} else {
ifres12472 = False


}

if True == ifres12472 {
return Nil
} else {
tmp12470 := GConsP(__e, V3691)

var ifres12468 Obj

if True == tmp12470 {
tmp12469 := GConsP(__e, V3692)

ifres12468 = tmp12469


} else {
ifres12468 = False


}

if True == ifres12468 {
tmp12459 := GHd(__e, V3692)

tmp12460 := GHd(__e, V3691)

tmp12461 := GCons(__e, tmp12460, Nil)

tmp12462 := GCons(__e, tmp12459, tmp12461)

tmp12463 := GCons(__e, symbind, tmp12462)

tmp12464 := GTl(__e, V3691)

tmp12465 := GTl(__e, V3692)

tmp12466 := Call2(__e, SymFn(symshen_4passive_1bind), tmp12464, tmp12465)


return GCons(__e, tmp12463, tmp12466)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4passive_1bind)
}


}


panic("unreachable")
})

tmp12475 := Call2(__e, ns2_1set, symshen_4passive_1bind, tmp12458)


_ = tmp12475

tmp12476 := MakeNative5(func(__e *ControlFlow, V3693, V3694, V3695, V3696, V3697 Obj) Obj {
tmp12477 := Call1(__e, SymFn(symgensym), symshen_4consume)


let__11608 := tmp12477
_ = let__11608

tmp12478 := Call5(__e, SymFn(symshen_4consume_1clause), let__11608, V3693, V3695, V3696, V3697)


let__11609 := tmp12478
_ = let__11609

tmp12479 := Call2(__e, SymFn(symappend), V3695, V3696)


tmp12480 := GCons(__e, V3697, tmp12479)

tmp12481 := GCons(__e, V3694, tmp12480)

return GCons(__e, let__11608, tmp12481)


panic("unreachable")
})

tmp12482 := Call2(__e, ns2_1set, symshen_4specialise_1consume, tmp12476)


_ = tmp12482

tmp12483 := MakeNative5(func(__e *ControlFlow, V3700, V3701, V3702, V3703, V3704 Obj) Obj {
tmp12484 := Call1(__e, SymFn(symlength), V3703)


tmp12485 := Call1(__e, SymFn(symshen_4nvars), tmp12484)


let__11610 := tmp12485
_ = let__11610

tmp12486 := Call1(__e, SymFn(symgensym), symAssumption)


let__11611 := tmp12486
_ = let__11611

tmp12487 := Call1(__e, SymFn(symshen_4macro_1_8ch), V3701)


tmp12488 := GCons(__e, let__11611, Nil)

tmp12489 := GCons(__e, tmp12487, tmp12488)

tmp12490 := GCons(__e, symcons, tmp12489)

tmp12491 := GCons(__e, tmp12490, Nil)

tmp12492 := GCons(__e, sym_1, tmp12491)

tmp12493 := GCons(__e, sym_5_1_1, Nil)

tmp12494 := Call2(__e, SymFn(symshen_4passive_1bind), V3703, let__11610)


tmp12495 := GCons(__e, let__11611, Nil)

tmp12496 := GCons(__e, V3704, tmp12495)

tmp12497 := GCons(__e, symbind, tmp12496)

tmp12498 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp12499 := GCons(__e, tmp12498, Nil)

tmp12500 := GCons(__e, tmp12497, tmp12499)

tmp12501 := Call2(__e, SymFn(symappend), tmp12494, tmp12500)


tmp12502 := Call2(__e, SymFn(symappend), tmp12493, tmp12501)


tmp12503 := Call2(__e, SymFn(symappend), let__11610, tmp12502)


tmp12504 := Call2(__e, SymFn(symappend), V3702, tmp12503)


tmp12505 := GCons(__e, V3704, tmp12504)

tmp12506 := GCons(__e, tmp12492, tmp12505)

let__11612 := tmp12506
_ = let__11612

tmp12507 := Call1(__e, SymFn(symgensym), symHypotheses)


let__11614 := tmp12507
_ = let__11614

tmp12508 := Call2(__e, SymFn(symappend), V3702, V3703)


let__11615 := tmp12508
_ = let__11615

tmp12509 := Call1(__e, SymFn(symgensym), symAssumptions)


let__11616 := tmp12509
_ = let__11616

tmp12510 := GCons(__e, let__11614, Nil)

tmp12511 := GCons(__e, let__11611, tmp12510)

tmp12512 := GCons(__e, symcons, tmp12511)

tmp12513 := GCons(__e, tmp12512, Nil)

tmp12514 := GCons(__e, sym_1, tmp12513)

tmp12515 := GCons(__e, V3704, Nil)

tmp12516 := GCons(__e, let__11616, tmp12515)

tmp12517 := GCons(__e, symcons, tmp12516)

tmp12518 := GCons(__e, tmp12517, let__11615)

tmp12519 := GCons(__e, tmp12514, tmp12518)

let__11617 := tmp12519
_ = let__11617

tmp12520 := GCons(__e, let__11611, Nil)

tmp12521 := GCons(__e, let__11616, tmp12520)

tmp12522 := GCons(__e, symbind, tmp12521)

tmp12523 := GCons(__e, V3704, let__11615)

tmp12524 := GCons(__e, let__11614, tmp12523)

tmp12525 := GCons(__e, V3700, tmp12524)

tmp12526 := GCons(__e, tmp12525, Nil)

tmp12527 := GCons(__e, tmp12522, tmp12526)

let__11618 := tmp12527
_ = let__11618

tmp12528 := GCons(__e, sym_5_1_1, Nil)

tmp12529 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

tmp12530 := GCons(__e, tmp12529, Nil)

tmp12531 := Call2(__e, SymFn(symappend), let__11618, tmp12530)


tmp12532 := Call2(__e, SymFn(symappend), tmp12528, tmp12531)


tmp12533 := Call2(__e, SymFn(symappend), let__11617, tmp12532)


let__11613 := tmp12533
_ = let__11613

tmp12534 := Call2(__e, SymFn(symappend), let__11612, let__11613)


tmp12535 := GCons(__e, V3700, tmp12534)

tmp12536 := GCons(__e, symdefprolog, tmp12535)

let__11619 := tmp12536
_ = let__11619

return TailCall1(__e, SymFn(symeval), let__11619)


panic("unreachable")
})

tmp12537 := Call2(__e, ns2_1set, symshen_4consume_1clause, tmp12483)


_ = tmp12537

tmp12538 := MakeNative2(func(__e *ControlFlow, V3715, V3716 Obj) Obj {
tmp12539 := Call1(__e, SymFn(symshen_4extract_1vars), V3715)


return TailCall2(__e, SymFn(symdifference), tmp12539, V3716)


panic("unreachable")
})

tmp12540 := Call2(__e, ns2_1set, symshen_4passive_1variables, tmp12538)


_ = tmp12540

tmp12541 := (func() Obj {
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
tmp12675 := GEq(__e, Nil, V3724)

if True == tmp12675 {
return TailCall3(__e, SymFn(symshen_4premises_1_6goals), V3721, V3723, V3725)
} else {
tmp12673 := GConsP(__e, V3724)

var ifres12657 Obj

if True == tmp12673 {
tmp12671 := GHd(__e, V3724)

tmp12672 := GConsP(__e, tmp12671)

var ifres12658 Obj

if True == tmp12672 {
tmp12668 := GHd(__e, V3724)

tmp12669 := GHd(__e, tmp12668)

tmp12670 := GEq(__e, symif, tmp12669)

var ifres12659 Obj

if True == tmp12670 {
tmp12665 := GHd(__e, V3724)

tmp12666 := GTl(__e, tmp12665)

tmp12667 := GConsP(__e, tmp12666)

var ifres12660 Obj

if True == tmp12667 {
tmp12661 := GHd(__e, V3724)

tmp12662 := GTl(__e, tmp12661)

tmp12663 := GTl(__e, tmp12662)

tmp12664 := GEq(__e, Nil, tmp12663)

ifres12660 = tmp12664


} else {
ifres12660 = False


}

ifres12659 = ifres12660


} else {
ifres12659 = False


}

ifres12658 = ifres12659


} else {
ifres12658 = False


}

ifres12657 = ifres12658


} else {
ifres12657 = False


}

if True == ifres12657 {
tmp12542 := GHd(__e, V3724)

tmp12543 := GTl(__e, tmp12542)

tmp12544 := GCons(__e, symwhen, tmp12543)

tmp12545 := GTl(__e, V3724)

tmp12546 := Call5(__e, SymFn(symshen_4side_1conditions_1_6goals), V3721, V3722, V3723, tmp12545, V3725)


return GCons(__e, tmp12544, tmp12546)


} else {
tmp12655 := GConsP(__e, V3724)

var ifres12633 Obj

if True == tmp12655 {
tmp12653 := GHd(__e, V3724)

tmp12654 := GConsP(__e, tmp12653)

var ifres12634 Obj

if True == tmp12654 {
tmp12650 := GHd(__e, V3724)

tmp12651 := GHd(__e, tmp12650)

tmp12652 := GEq(__e, symlet, tmp12651)

var ifres12635 Obj

if True == tmp12652 {
tmp12647 := GHd(__e, V3724)

tmp12648 := GTl(__e, tmp12647)

tmp12649 := GConsP(__e, tmp12648)

var ifres12636 Obj

if True == tmp12649 {
tmp12643 := GHd(__e, V3724)

tmp12644 := GTl(__e, tmp12643)

tmp12645 := GTl(__e, tmp12644)

tmp12646 := GConsP(__e, tmp12645)

var ifres12637 Obj

if True == tmp12646 {
tmp12638 := GHd(__e, V3724)

tmp12639 := GTl(__e, tmp12638)

tmp12640 := GTl(__e, tmp12639)

tmp12641 := GTl(__e, tmp12640)

tmp12642 := GEq(__e, Nil, tmp12641)

ifres12637 = tmp12642


} else {
ifres12637 = False


}

ifres12636 = ifres12637


} else {
ifres12636 = False


}

ifres12635 = ifres12636


} else {
ifres12635 = False


}

ifres12634 = ifres12635


} else {
ifres12634 = False


}

ifres12633 = ifres12634


} else {
ifres12633 = False


}

if True == ifres12633 {
tmp12562 := GHd(__e, V3724)

tmp12563 := GTl(__e, tmp12562)

tmp12564 := GHd(__e, tmp12563)

tmp12565 := Call2(__e, SymFn(symelement_2), tmp12564, V3722)


if True == tmp12565 {
tmp12547 := GHd(__e, V3724)

tmp12548 := GTl(__e, tmp12547)

tmp12549 := GCons(__e, symis_b, tmp12548)

tmp12550 := GTl(__e, V3724)

tmp12551 := Call5(__e, SymFn(symshen_4side_1conditions_1_6goals), V3721, V3722, V3723, tmp12550, V3725)


return GCons(__e, tmp12549, tmp12551)


} else {
tmp12552 := GHd(__e, V3724)

tmp12553 := GTl(__e, tmp12552)

tmp12554 := GCons(__e, symbind, tmp12553)

tmp12555 := GHd(__e, V3724)

tmp12556 := GTl(__e, tmp12555)

tmp12557 := GHd(__e, tmp12556)

tmp12558 := GCons(__e, tmp12557, V3722)

tmp12559 := GTl(__e, V3724)

tmp12560 := Call5(__e, SymFn(symshen_4side_1conditions_1_6goals), V3721, tmp12558, V3723, tmp12559, V3725)


return GCons(__e, tmp12554, tmp12560)


}


} else {
tmp12631 := GConsP(__e, V3724)

var ifres12615 Obj

if True == tmp12631 {
tmp12629 := GHd(__e, V3724)

tmp12630 := GConsP(__e, tmp12629)

var ifres12616 Obj

if True == tmp12630 {
tmp12626 := GHd(__e, V3724)

tmp12627 := GHd(__e, tmp12626)

tmp12628 := GEq(__e, symctxt, tmp12627)

var ifres12617 Obj

if True == tmp12628 {
tmp12623 := GHd(__e, V3724)

tmp12624 := GTl(__e, tmp12623)

tmp12625 := GConsP(__e, tmp12624)

var ifres12618 Obj

if True == tmp12625 {
tmp12619 := GHd(__e, V3724)

tmp12620 := GTl(__e, tmp12619)

tmp12621 := GTl(__e, tmp12620)

tmp12622 := GEq(__e, Nil, tmp12621)

ifres12618 = tmp12622


} else {
ifres12618 = False


}

ifres12617 = ifres12618


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

if True == ifres12615 {
tmp12591 := GHd(__e, V3724)

tmp12592 := GTl(__e, tmp12591)

tmp12593 := GHd(__e, tmp12592)

tmp12594 := Call2(__e, SymFn(symelement_2), tmp12593, V3722)


if True == tmp12594 {
tmp12566 := GHd(__e, V3724)

tmp12567 := GTl(__e, tmp12566)

tmp12568 := GHd(__e, tmp12567)

tmp12569 := GCons(__e, tmp12568, V3721)

tmp12570 := GTl(__e, V3724)

if SymFn(symshen_4side_1conditions_1_6goals) == __self {
__self1, __self2, __self3, __self4, __self5 = tmp12569, V3722, V3723, tmp12570, V3725
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4side_1conditions_1_6goals), tmp12569, V3722, V3723, tmp12570, V3725)


} else {
tmp12571 := GHd(__e, V3724)

tmp12572 := GTl(__e, tmp12571)

tmp12573 := GHd(__e, tmp12572)

tmp12574 := GCons(__e, V3723, Nil)

tmp12575 := GCons(__e, tmp12573, tmp12574)

tmp12576 := GCons(__e, symbind, tmp12575)

tmp12577 := GHd(__e, V3724)

tmp12578 := GTl(__e, tmp12577)

tmp12579 := GHd(__e, tmp12578)

tmp12580 := GCons(__e, tmp12579, V3721)

tmp12581 := GHd(__e, V3724)

tmp12582 := GTl(__e, tmp12581)

tmp12583 := GHd(__e, tmp12582)

tmp12584 := GCons(__e, tmp12583, V3722)

tmp12585 := GHd(__e, V3724)

tmp12586 := GTl(__e, tmp12585)

tmp12587 := GHd(__e, tmp12586)

tmp12588 := GTl(__e, V3724)

tmp12589 := Call5(__e, SymFn(symshen_4side_1conditions_1_6goals), tmp12580, tmp12584, tmp12587, tmp12588, V3725)


return GCons(__e, tmp12576, tmp12589)


}


} else {
tmp12613 := GConsP(__e, V3724)

var ifres12597 Obj

if True == tmp12613 {
tmp12611 := GHd(__e, V3724)

tmp12612 := GConsP(__e, tmp12611)

var ifres12598 Obj

if True == tmp12612 {
tmp12608 := GHd(__e, V3724)

tmp12609 := GHd(__e, tmp12608)

tmp12610 := GEq(__e, symsqts, tmp12609)

var ifres12599 Obj

if True == tmp12610 {
tmp12605 := GHd(__e, V3724)

tmp12606 := GTl(__e, tmp12605)

tmp12607 := GConsP(__e, tmp12606)

var ifres12600 Obj

if True == tmp12607 {
tmp12601 := GHd(__e, V3724)

tmp12602 := GTl(__e, tmp12601)

tmp12603 := GTl(__e, tmp12602)

tmp12604 := GEq(__e, Nil, tmp12603)

ifres12600 = tmp12604


} else {
ifres12600 = False


}

ifres12599 = ifres12600


} else {
ifres12599 = False


}

ifres12598 = ifres12599


} else {
ifres12598 = False


}

ifres12597 = ifres12598


} else {
ifres12597 = False


}

if True == ifres12597 {
tmp12595 := GTl(__e, V3724)

if SymFn(symshen_4side_1conditions_1_6goals) == __self {
__self1, __self2, __self3, __self4, __self5 = V3721, V3722, V3723, tmp12595, V3725
__e.Tick()
goto __selftop
}
return TailCall5(__e, SymFn(symshen_4side_1conditions_1_6goals), V3721, V3722, V3723, tmp12595, V3725)


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

tmp12676 := Call2(__e, ns2_1set, symshen_4side_1conditions_1_6goals, tmp12541)


_ = tmp12676

tmp12677 := MakeNative3(func(__e *ControlFlow, V3730, V3731, V3732 Obj) Obj {
tmp12722 := GEq(__e, Nil, V3732)

if True == tmp12722 {
tmp12678 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

return GCons(__e, tmp12678, Nil)


} else {
tmp12720 := GConsP(__e, V3732)

var ifres12717 Obj

if True == tmp12720 {
tmp12718 := GHd(__e, V3732)

tmp12719 := GEq(__e, sym_b, tmp12718)

ifres12717 = tmp12719


} else {
ifres12717 = False


}

if True == ifres12717 {
tmp12679 := GTl(__e, V3732)

tmp12680 := Call3(__e, SymFn(symshen_4premises_1_6goals), V3730, V3731, tmp12679)


return GCons(__e, sym_b, tmp12680)


} else {
tmp12715 := GConsP(__e, V3732)

var ifres12712 Obj

if True == tmp12715 {
tmp12713 := GHd(__e, V3732)

tmp12714 := GEq(__e, symfail, tmp12713)

ifres12712 = tmp12714


} else {
ifres12712 = False


}

if True == ifres12712 {
tmp12681 := GCons(__e, False, Nil)

tmp12682 := GCons(__e, symwhen, tmp12681)

tmp12683 := GTl(__e, V3732)

tmp12684 := Call3(__e, SymFn(symshen_4premises_1_6goals), V3730, V3731, tmp12683)


return GCons(__e, tmp12682, tmp12684)


} else {
tmp12710 := GConsP(__e, V3732)

var ifres12698 Obj

if True == tmp12710 {
tmp12708 := GHd(__e, V3732)

tmp12709 := GConsP(__e, tmp12708)

var ifres12699 Obj

if True == tmp12709 {
tmp12705 := GHd(__e, V3732)

tmp12706 := GTl(__e, tmp12705)

tmp12707 := GConsP(__e, tmp12706)

var ifres12700 Obj

if True == tmp12707 {
tmp12701 := GHd(__e, V3732)

tmp12702 := GTl(__e, tmp12701)

tmp12703 := GTl(__e, tmp12702)

tmp12704 := GEq(__e, Nil, tmp12703)

ifres12700 = tmp12704


} else {
ifres12700 = False


}

ifres12699 = ifres12700


} else {
ifres12699 = False


}

ifres12698 = ifres12699


} else {
ifres12698 = False


}

if True == ifres12698 {
tmp12685 := GHd(__e, V3732)

tmp12686 := GTl(__e, tmp12685)

tmp12687 := GHd(__e, tmp12686)

tmp12688 := Call1(__e, SymFn(symshen_4macro_1_8c), tmp12687)


tmp12689 := GHd(__e, V3732)

tmp12690 := GHd(__e, tmp12689)

tmp12691 := Call3(__e, SymFn(symshen_4construct_1context), V3730, tmp12690, V3731)


tmp12692 := GCons(__e, tmp12691, Nil)

tmp12693 := GCons(__e, tmp12688, tmp12692)

tmp12694 := GCons(__e, symshen_4system_1S, tmp12693)

tmp12695 := GTl(__e, V3732)

tmp12696 := Call3(__e, SymFn(symshen_4premises_1_6goals), V3730, V3731, tmp12695)


return GCons(__e, tmp12694, tmp12696)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4premises_1_6goals)
}


}


}


}


panic("unreachable")
})

tmp12723 := Call2(__e, ns2_1set, symshen_4premises_1_6goals, tmp12677)


_ = tmp12723

tmp12724 := MakeNative3(func(__e *ControlFlow, V3736, V3737, V3738 Obj) Obj {
tmp12742 := GEq(__e, Nil, V3737)

if True == tmp12742 {
return V3738
} else {
tmp12740 := GConsP(__e, V3737)

var ifres12734 Obj

if True == tmp12740 {
tmp12738 := GTl(__e, V3737)

tmp12739 := GEq(__e, Nil, tmp12738)

var ifres12735 Obj

if True == tmp12739 {
tmp12736 := GHd(__e, V3737)

tmp12737 := Call2(__e, SymFn(symelement_2), tmp12736, V3736)


ifres12735 = tmp12737


} else {
ifres12735 = False


}

ifres12734 = ifres12735


} else {
ifres12734 = False


}

if True == ifres12734 {
return GHd(__e, V3737)
} else {
tmp12732 := GConsP(__e, V3737)

if True == tmp12732 {
tmp12725 := GHd(__e, V3737)

tmp12726 := Call1(__e, SymFn(symshen_4macro_1_8c), tmp12725)


tmp12727 := GTl(__e, V3737)

tmp12728 := Call3(__e, SymFn(symshen_4construct_1context), V3736, tmp12727, V3738)


tmp12729 := GCons(__e, tmp12728, Nil)

tmp12730 := GCons(__e, tmp12726, tmp12729)

return GCons(__e, symcons, tmp12730)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4construct_1context)
}


}


}


panic("unreachable")
})

tmp12743 := Call2(__e, ns2_1set, symshen_4construct_1context, tmp12724)


_ = tmp12743

tmp12744 := MakeNative1(func(__e *ControlFlow, V3739 Obj) Obj {
tmp12745 := MakeNative1(func(__e *ControlFlow, Z3741 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4intern_1type), Z3741)
panic("unreachable")
})

tmp12746 := Call2(__e, SymFn(symmap), tmp12745, V3739)


let__11620 := tmp12746
_ = let__11620

tmp12747 := GValue(__e, symshen_4_ddatatypes_d)

let__11621 := tmp12747
_ = let__11621

tmp12748 := Call2(__e, SymFn(symshen_4remove_1datatypes), let__11620, let__11621)


let__11622 := tmp12748
_ = let__11622

tmp12749 := PrimGuard2(__e, symset, PrimSet, symshen_4_ddatatypes_d, let__11622)

let__11623 := tmp12749
_ = let__11623

return TailCall1(__e, SymFn(symshen_4show_1datatypes), let__11623)


panic("unreachable")
})

tmp12750 := Call2(__e, ns2_1set, sympreclude, tmp12744)


_ = tmp12750

tmp12751 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V3749 := __self1
_ = V3749
V3750 := __self2
_ = V3750
tmp12758 := GEq(__e, Nil, V3749)

if True == tmp12758 {
return V3750
} else {
tmp12756 := GConsP(__e, V3749)

if True == tmp12756 {
tmp12752 := GTl(__e, V3749)

tmp12753 := GHd(__e, V3749)

tmp12754 := Call2(__e, SymFn(symshen_4unassoc), tmp12753, V3750)


if SymFn(symshen_4remove_1datatypes) == __self {
__self1, __self2 = tmp12752, tmp12754
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4remove_1datatypes), tmp12752, tmp12754)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.remove-datatypes"))
}


}


panic("unreachable")
})
return __self
})()

tmp12759 := Call2(__e, ns2_1set, symshen_4remove_1datatypes, tmp12751)


_ = tmp12759

tmp12760 := MakeNative1(func(__e *ControlFlow, V3751 Obj) Obj {
tmp12761 := MakeNative1(func(__e *ControlFlow, Z3752 Obj) Obj {
return GHd(__e, Z3752)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp12761, V3751)


panic("unreachable")
})

tmp12762 := Call2(__e, ns2_1set, symshen_4show_1datatypes, tmp12760)


_ = tmp12762

tmp12763 := MakeNative1(func(__e *ControlFlow, V3753 Obj) Obj {
tmp12764 := MakeNative1(func(__e *ControlFlow, Z3755 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4intern_1type), Z3755)
panic("unreachable")
})

tmp12765 := Call2(__e, SymFn(symmap), tmp12764, V3753)


let__11624 := tmp12765
_ = let__11624

tmp12766 := MakeNative1(func(__e *ControlFlow, Z3757 Obj) Obj {
tmp12767 := Call1(__e, SymFn(symfn), Z3757)


return TailCall2(__e, SymFn(symshen_4remember_1datatype), Z3757, tmp12767)


panic("unreachable")
})

tmp12768 := Call2(__e, SymFn(symmap), tmp12766, let__11624)


let__11625 := tmp12768
_ = let__11625

tmp12769 := GValue(__e, symshen_4_ddatatypes_d)

let__11626 := tmp12769
_ = let__11626

return TailCall1(__e, SymFn(symshen_4show_1datatypes), let__11626)


panic("unreachable")
})

tmp12770 := Call2(__e, ns2_1set, syminclude, tmp12763)


_ = tmp12770

tmp12771 := MakeNative1(func(__e *ControlFlow, V3759 Obj) Obj {
tmp12772 := PrimGuard2(__e, symset, PrimSet, symshen_4_ddatatypes_d, Nil)

let__11627 := tmp12772
_ = let__11627

tmp12773 := MakeNative1(func(__e *ControlFlow, Z3762 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4intern_1type), Z3762)
panic("unreachable")
})

tmp12774 := Call2(__e, SymFn(symmap), tmp12773, V3759)


let__11628 := tmp12774
_ = let__11628

tmp12775 := MakeNative1(func(__e *ControlFlow, Z3764 Obj) Obj {
tmp12776 := Call1(__e, SymFn(symfn), Z3764)


return TailCall2(__e, SymFn(symshen_4remember_1datatype), Z3764, tmp12776)


panic("unreachable")
})

tmp12777 := Call2(__e, SymFn(symmap), tmp12775, let__11628)


let__11629 := tmp12777
_ = let__11629

tmp12778 := GValue(__e, symshen_4_ddatatypes_d)

return TailCall1(__e, SymFn(symshen_4show_1datatypes), tmp12778)


panic("unreachable")
})

tmp12779 := Call2(__e, ns2_1set, sympreclude_1all_1but, tmp12771)


_ = tmp12779

tmp12780 := MakeNative1(func(__e *ControlFlow, V3765 Obj) Obj {
tmp12781 := MakeNative1(func(__e *ControlFlow, Z3767 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4intern_1type), Z3767)
panic("unreachable")
})

tmp12782 := Call2(__e, SymFn(symmap), tmp12781, V3765)


let__11630 := tmp12782
_ = let__11630

tmp12783 := GValue(__e, symshen_4_dalldatatypes_d)

let__11631 := tmp12783
_ = let__11631

tmp12784 := Call2(__e, SymFn(symshen_4remove_1datatypes), let__11630, let__11631)


tmp12785 := PrimGuard2(__e, symset, PrimSet, symshen_4_ddatatypes_d, tmp12784)

let__11632 := tmp12785
_ = let__11632

return TailCall1(__e, SymFn(symshen_4show_1datatypes), let__11632)


panic("unreachable")
})

return TailCall2(__e, ns2_1set, syminclude_1all_1but, tmp12780)




panic("unreachable")
})

