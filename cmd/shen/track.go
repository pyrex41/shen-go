package main

import . "github.com/tiancaiamao/shen-go/kl"

var TrackMain = MakeNative(func(__e *ControlFlow) {
tmp13613 := MakeNative(func(__e *ControlFlow) {
V5370 := __e.Get(1)
_ = V5370
tmp13614 := Call(__e, PrimFunc(symshen_4app), V5370, MakeString(";\n"), symshen_4a)


tmp13615 := PrimStringConcat(MakeString("partial function "), tmp13614)

tmp13616 := Call(__e, PrimFunc(symstoutput))


tmp13617 := Call(__e, PrimFunc(sympr), tmp13615, tmp13616)


_ = tmp13617

tmp13626 := Call(__e, PrimFunc(symshen_4tracked_2), V5370)


tmp13627 := PrimNot(tmp13626)

var ifres13621 Obj

if True == tmp13627 {
tmp13623 := Call(__e, PrimFunc(symshen_4app), V5370, MakeString("? "), symshen_4a)


tmp13624 := PrimStringConcat(MakeString("track "), tmp13623)

tmp13625 := Call(__e, PrimFunc(symy_1or_1n_2), tmp13624)


var ifres13622 Obj

if True == tmp13625 {
ifres13622 = True


} else {
ifres13622 = False


}

ifres13621 = ifres13622


} else {
ifres13621 = False


}

var ifres13618 Obj

if True == ifres13621 {
tmp13619 := Call(__e, PrimFunc(symps), V5370)


tmp13620 := Call(__e, PrimFunc(symshen_4track_1function), tmp13619)


ifres13618 = tmp13620


} else {
ifres13618 = symshen_4ok


}

_ = ifres13618

__e.Return(PrimSimpleError(MakeString("aborted")))
return


}, 1)

tmp13628 := Call(__e, ns2_1set, symshen_4f_1error, tmp13613)


_ = tmp13628

tmp13629 := MakeNative(func(__e *ControlFlow) {
V5371 := __e.Get(1)
_ = V5371
tmp13630 := PrimValue(symshen_4_dtracking_d)

__e.TailApply(PrimFunc(symelement_2), V5371, tmp13630)
return


}, 1)

tmp13631 := Call(__e, ns2_1set, symshen_4tracked_2, tmp13629)


_ = tmp13631

tmp13632 := MakeNative(func(__e *ControlFlow) {
V5372 := __e.Get(1)
_ = V5372
tmp13633 := MakeNative(func(__e *ControlFlow) {
W5373 := __e.Get(1)
_ = W5373
__e.TailApply(PrimFunc(symshen_4track_1function), W5373)
return
}, 1)

tmp13634 := Call(__e, PrimFunc(symps), V5372)


__e.TailApply(tmp13633, tmp13634)
return


}, 1)

tmp13635 := Call(__e, ns2_1set, symtrack, tmp13632)


_ = tmp13635

tmp13636 := MakeNative(func(__e *ControlFlow) {
V5376 := __e.Get(1)
_ = V5376
tmp13693 := PrimIsPair(V5376)

var ifres13667 Obj

if True == tmp13693 {
tmp13691 := PrimHead(V5376)

tmp13692 := PrimEqual(symdefun, tmp13691)

var ifres13669 Obj

if True == tmp13692 {
tmp13689 := PrimTail(V5376)

tmp13690 := PrimIsPair(tmp13689)

var ifres13671 Obj

if True == tmp13690 {
tmp13686 := PrimTail(V5376)

tmp13687 := PrimTail(tmp13686)

tmp13688 := PrimIsPair(tmp13687)

var ifres13673 Obj

if True == tmp13688 {
tmp13682 := PrimTail(V5376)

tmp13683 := PrimTail(tmp13682)

tmp13684 := PrimTail(tmp13683)

tmp13685 := PrimIsPair(tmp13684)

var ifres13675 Obj

if True == tmp13685 {
tmp13677 := PrimTail(V5376)

tmp13678 := PrimTail(tmp13677)

tmp13679 := PrimTail(tmp13678)

tmp13680 := PrimTail(tmp13679)

tmp13681 := PrimEqual(Nil, tmp13680)

var ifres13676 Obj

if True == tmp13681 {
ifres13676 = True


} else {
ifres13676 = False


}

ifres13675 = ifres13676


} else {
ifres13675 = False


}

var ifres13674 Obj

if True == ifres13675 {
ifres13674 = True


} else {
ifres13674 = False


}

ifres13673 = ifres13674


} else {
ifres13673 = False


}

var ifres13672 Obj

if True == ifres13673 {
ifres13672 = True


} else {
ifres13672 = False


}

ifres13671 = ifres13672


} else {
ifres13671 = False


}

var ifres13670 Obj

if True == ifres13671 {
ifres13670 = True


} else {
ifres13670 = False


}

ifres13669 = ifres13670


} else {
ifres13669 = False


}

var ifres13668 Obj

if True == ifres13669 {
ifres13668 = True


} else {
ifres13668 = False


}

ifres13667 = ifres13668


} else {
ifres13667 = False


}

if True == ifres13667 {
tmp13637 := MakeNative(func(__e *ControlFlow) {
W5377 := __e.Get(1)
_ = W5377
tmp13638 := MakeNative(func(__e *ControlFlow) {
W5378 := __e.Get(1)
_ = W5378
tmp13639 := MakeNative(func(__e *ControlFlow) {
W5379 := __e.Get(1)
_ = W5379
tmp13640 := PrimTail(V5376)

__e.Return(PrimHead(tmp13640))
return


}, 1)

tmp13641 := PrimTail(V5376)

tmp13642 := PrimHead(tmp13641)

tmp13643 := PrimValue(symshen_4_dtracking_d)

tmp13644 := Call(__e, PrimFunc(symadjoin), tmp13642, tmp13643)


tmp13645 := PrimSet(symshen_4_dtracking_d, tmp13644)

__e.TailApply(tmp13639, tmp13645)
return


}, 1)

tmp13646 := Call(__e, PrimFunc(symeval_1kl), W5377)


__e.TailApply(tmp13638, tmp13646)
return


}, 1)

tmp13647 := PrimTail(V5376)

tmp13648 := PrimHead(tmp13647)

tmp13649 := PrimTail(V5376)

tmp13650 := PrimTail(tmp13649)

tmp13651 := PrimHead(tmp13650)

tmp13652 := PrimTail(V5376)

tmp13653 := PrimHead(tmp13652)

tmp13654 := PrimTail(V5376)

tmp13655 := PrimTail(tmp13654)

tmp13656 := PrimHead(tmp13655)

tmp13657 := PrimTail(V5376)

tmp13658 := PrimTail(tmp13657)

tmp13659 := PrimTail(tmp13658)

tmp13660 := PrimHead(tmp13659)

tmp13661 := Call(__e, PrimFunc(symshen_4insert_1tracking_1code), tmp13653, tmp13656, tmp13660)


tmp13662 := PrimCons(tmp13661, Nil)

tmp13663 := PrimCons(tmp13651, tmp13662)

tmp13664 := PrimCons(tmp13648, tmp13663)

tmp13665 := PrimCons(symdefun, tmp13664)

__e.TailApply(tmp13637, tmp13665)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.track-function")))
return
}


}, 1)

tmp13694 := Call(__e, ns2_1set, symshen_4track_1function, tmp13636)


_ = tmp13694

tmp13695 := MakeNative(func(__e *ControlFlow) {
V5380 := __e.Get(1)
_ = V5380
V5381 := __e.Get(2)
_ = V5381
V5382 := __e.Get(3)
_ = V5382
tmp13696 := PrimCons(symshen_4_dcall_d, Nil)

tmp13697 := PrimCons(symvalue, tmp13696)

tmp13698 := PrimCons(MakeNumber(1), Nil)

tmp13699 := PrimCons(tmp13697, tmp13698)

tmp13700 := PrimCons(sym_7, tmp13699)

tmp13701 := PrimCons(tmp13700, Nil)

tmp13702 := PrimCons(symshen_4_dcall_d, tmp13701)

tmp13703 := PrimCons(symset, tmp13702)

tmp13704 := PrimCons(symshen_4_dcall_d, Nil)

tmp13705 := PrimCons(symvalue, tmp13704)

tmp13706 := Call(__e, PrimFunc(symshen_4prolog_1track), V5382, V5381)


tmp13707 := Call(__e, PrimFunc(symshen_4cons_1form), tmp13706)


tmp13708 := PrimCons(tmp13707, Nil)

tmp13709 := PrimCons(V5380, tmp13708)

tmp13710 := PrimCons(tmp13705, tmp13709)

tmp13711 := PrimCons(symshen_4input_1track, tmp13710)

tmp13712 := PrimCons(symshen_4terpri_1or_1read_1char, Nil)

tmp13713 := PrimCons(symshen_4_dcall_d, Nil)

tmp13714 := PrimCons(symvalue, tmp13713)

tmp13715 := PrimCons(symResult, Nil)

tmp13716 := PrimCons(V5380, tmp13715)

tmp13717 := PrimCons(tmp13714, tmp13716)

tmp13718 := PrimCons(symshen_4output_1track, tmp13717)

tmp13719 := PrimCons(symshen_4_dcall_d, Nil)

tmp13720 := PrimCons(symvalue, tmp13719)

tmp13721 := PrimCons(MakeNumber(1), Nil)

tmp13722 := PrimCons(tmp13720, tmp13721)

tmp13723 := PrimCons(sym_1, tmp13722)

tmp13724 := PrimCons(tmp13723, Nil)

tmp13725 := PrimCons(symshen_4_dcall_d, tmp13724)

tmp13726 := PrimCons(symset, tmp13725)

tmp13727 := PrimCons(symshen_4terpri_1or_1read_1char, Nil)

tmp13728 := PrimCons(symResult, Nil)

tmp13729 := PrimCons(tmp13727, tmp13728)

tmp13730 := PrimCons(symdo, tmp13729)

tmp13731 := PrimCons(tmp13730, Nil)

tmp13732 := PrimCons(tmp13726, tmp13731)

tmp13733 := PrimCons(symdo, tmp13732)

tmp13734 := PrimCons(tmp13733, Nil)

tmp13735 := PrimCons(tmp13718, tmp13734)

tmp13736 := PrimCons(symdo, tmp13735)

tmp13737 := PrimCons(tmp13736, Nil)

tmp13738 := PrimCons(V5382, tmp13737)

tmp13739 := PrimCons(symResult, tmp13738)

tmp13740 := PrimCons(symlet, tmp13739)

tmp13741 := PrimCons(tmp13740, Nil)

tmp13742 := PrimCons(tmp13712, tmp13741)

tmp13743 := PrimCons(symdo, tmp13742)

tmp13744 := PrimCons(tmp13743, Nil)

tmp13745 := PrimCons(tmp13711, tmp13744)

tmp13746 := PrimCons(symdo, tmp13745)

tmp13747 := PrimCons(tmp13746, Nil)

tmp13748 := PrimCons(tmp13703, tmp13747)

__e.Return(PrimCons(symdo, tmp13748))
return


}, 3)

tmp13749 := Call(__e, ns2_1set, symshen_4insert_1tracking_1code, tmp13695)


_ = tmp13749

tmp13750 := MakeNative(func(__e *ControlFlow) {
V5383 := __e.Get(1)
_ = V5383
V5384 := __e.Get(2)
_ = V5384
tmp13753 := Call(__e, PrimFunc(symoccurrences), symshen_4incinfs, V5383)


tmp13754 := PrimEqual(tmp13753, MakeNumber(0))

if True == tmp13754 {
__e.Return(V5384)
return
} else {
tmp13751 := Call(__e, PrimFunc(symshen_4vector_1parameter), V5384)


__e.TailApply(PrimFunc(symshen_4vector_1dereference), V5384, tmp13751)
return


}


}, 2)

tmp13755 := Call(__e, ns2_1set, symshen_4prolog_1track, tmp13750)


_ = tmp13755

tmp13756 := MakeNative(func(__e *ControlFlow) {
V5387 := __e.Get(1)
_ = V5387
tmp13785 := PrimEqual(Nil, V5387)

if True == tmp13785 {
__e.Return(Nil)
return
} else {
tmp13783 := PrimIsPair(V5387)

var ifres13761 Obj

if True == tmp13783 {
tmp13781 := PrimTail(V5387)

tmp13782 := PrimIsPair(tmp13781)

var ifres13763 Obj

if True == tmp13782 {
tmp13778 := PrimTail(V5387)

tmp13779 := PrimTail(tmp13778)

tmp13780 := PrimIsPair(tmp13779)

var ifres13765 Obj

if True == tmp13780 {
tmp13774 := PrimTail(V5387)

tmp13775 := PrimTail(tmp13774)

tmp13776 := PrimTail(tmp13775)

tmp13777 := PrimIsPair(tmp13776)

var ifres13767 Obj

if True == tmp13777 {
tmp13769 := PrimTail(V5387)

tmp13770 := PrimTail(tmp13769)

tmp13771 := PrimTail(tmp13770)

tmp13772 := PrimTail(tmp13771)

tmp13773 := PrimEqual(Nil, tmp13772)

var ifres13768 Obj

if True == tmp13773 {
ifres13768 = True


} else {
ifres13768 = False


}

ifres13767 = ifres13768


} else {
ifres13767 = False


}

var ifres13766 Obj

if True == ifres13767 {
ifres13766 = True


} else {
ifres13766 = False


}

ifres13765 = ifres13766


} else {
ifres13765 = False


}

var ifres13764 Obj

if True == ifres13765 {
ifres13764 = True


} else {
ifres13764 = False


}

ifres13763 = ifres13764


} else {
ifres13763 = False


}

var ifres13762 Obj

if True == ifres13763 {
ifres13762 = True


} else {
ifres13762 = False


}

ifres13761 = ifres13762


} else {
ifres13761 = False


}

if True == ifres13761 {
__e.Return(PrimHead(V5387))
return
} else {
tmp13759 := PrimIsPair(V5387)

if True == tmp13759 {
tmp13757 := PrimTail(V5387)

__e.TailApply(PrimFunc(symshen_4vector_1parameter), tmp13757)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.vector-parameter")))
return
}


}


}


}, 1)

tmp13786 := Call(__e, ns2_1set, symshen_4vector_1parameter, tmp13756)


_ = tmp13786

tmp13787 := MakeNative(func(__e *ControlFlow) {
V5390 := __e.Get(1)
_ = V5390
V5391 := __e.Get(2)
_ = V5391
tmp13821 := PrimEqual(Nil, V5391)

if True == tmp13821 {
__e.Return(V5390)
return
} else {
tmp13819 := PrimIsPair(V5390)

var ifres13797 Obj

if True == tmp13819 {
tmp13817 := PrimTail(V5390)

tmp13818 := PrimIsPair(tmp13817)

var ifres13799 Obj

if True == tmp13818 {
tmp13814 := PrimTail(V5390)

tmp13815 := PrimTail(tmp13814)

tmp13816 := PrimIsPair(tmp13815)

var ifres13801 Obj

if True == tmp13816 {
tmp13810 := PrimTail(V5390)

tmp13811 := PrimTail(tmp13810)

tmp13812 := PrimTail(tmp13811)

tmp13813 := PrimIsPair(tmp13812)

var ifres13803 Obj

if True == tmp13813 {
tmp13805 := PrimTail(V5390)

tmp13806 := PrimTail(tmp13805)

tmp13807 := PrimTail(tmp13806)

tmp13808 := PrimTail(tmp13807)

tmp13809 := PrimEqual(Nil, tmp13808)

var ifres13804 Obj

if True == tmp13809 {
ifres13804 = True


} else {
ifres13804 = False


}

ifres13803 = ifres13804


} else {
ifres13803 = False


}

var ifres13802 Obj

if True == ifres13803 {
ifres13802 = True


} else {
ifres13802 = False


}

ifres13801 = ifres13802


} else {
ifres13801 = False


}

var ifres13800 Obj

if True == ifres13801 {
ifres13800 = True


} else {
ifres13800 = False


}

ifres13799 = ifres13800


} else {
ifres13799 = False


}

var ifres13798 Obj

if True == ifres13799 {
ifres13798 = True


} else {
ifres13798 = False


}

ifres13797 = ifres13798


} else {
ifres13797 = False


}

if True == ifres13797 {
__e.Return(V5390)
return
} else {
tmp13795 := PrimIsPair(V5390)

if True == tmp13795 {
tmp13788 := PrimHead(V5390)

tmp13789 := PrimCons(V5391, Nil)

tmp13790 := PrimCons(tmp13788, tmp13789)

tmp13791 := PrimCons(symshen_4deref, tmp13790)

tmp13792 := PrimTail(V5390)

tmp13793 := Call(__e, PrimFunc(symshen_4vector_1dereference), tmp13792, V5391)


__e.Return(PrimCons(tmp13791, tmp13793))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.vector-dereference")))
return
}


}


}


}, 2)

tmp13822 := Call(__e, ns2_1set, symshen_4vector_1dereference, tmp13787)


_ = tmp13822

tmp13823 := MakeNative(func(__e *ControlFlow) {
V5394 := __e.Get(1)
_ = V5394
tmp13827 := PrimEqual(sym_7, V5394)

if True == tmp13827 {
__e.Return(PrimSet(symshen_4_dstep_d, True))
return
} else {
tmp13825 := PrimEqual(sym_1, V5394)

if True == tmp13825 {
__e.Return(PrimSet(symshen_4_dstep_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("step expects a + or a -.\n")))
return
}


}


}, 1)

tmp13828 := Call(__e, ns2_1set, symstep, tmp13823)


_ = tmp13828

tmp13829 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dstep_d))
return
}, 0)

tmp13830 := Call(__e, ns2_1set, symstep_2, tmp13829)


_ = tmp13830

tmp13831 := MakeNative(func(__e *ControlFlow) {
V5397 := __e.Get(1)
_ = V5397
tmp13835 := PrimEqual(sym_7, V5397)

if True == tmp13835 {
__e.Return(PrimSet(symshen_4_dspy_d, True))
return
} else {
tmp13833 := PrimEqual(sym_1, V5397)

if True == tmp13833 {
__e.Return(PrimSet(symshen_4_dspy_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("spy expects a + or a -.\n")))
return
}


}


}, 1)

tmp13836 := Call(__e, ns2_1set, symspy, tmp13831)


_ = tmp13836

tmp13837 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dspy_d))
return
}, 0)

tmp13838 := Call(__e, ns2_1set, symspy_2, tmp13837)


_ = tmp13838

tmp13839 := MakeNative(func(__e *ControlFlow) {
tmp13843 := PrimValue(symshen_4_dstep_d)

if True == tmp13843 {
tmp13840 := PrimValue(sym_dstinput_d)

tmp13841 := PrimReadByte(tmp13840)

__e.TailApply(PrimFunc(symshen_4check_1byte), tmp13841)
return


} else {
__e.TailApply(PrimFunc(symnl), MakeNumber(1))
return
}


}, 0)

tmp13844 := Call(__e, ns2_1set, symshen_4terpri_1or_1read_1char, tmp13839)


_ = tmp13844

tmp13845 := MakeNative(func(__e *ControlFlow) {
V5400 := __e.Get(1)
_ = V5400
tmp13847 := PrimEqual(MakeNumber(94), V5400)

if True == tmp13847 {
__e.Return(PrimSimpleError(MakeString("aborted")))
return
} else {
__e.Return(True)
return
}


}, 1)

tmp13848 := Call(__e, ns2_1set, symshen_4check_1byte, tmp13845)


_ = tmp13848

tmp13849 := MakeNative(func(__e *ControlFlow) {
V5401 := __e.Get(1)
_ = V5401
V5402 := __e.Get(2)
_ = V5402
V5403 := __e.Get(3)
_ = V5403
tmp13850 := Call(__e, PrimFunc(symshen_4spaces), V5401)


tmp13851 := Call(__e, PrimFunc(symshen_4spaces), V5401)


tmp13852 := Call(__e, PrimFunc(symshen_4app), tmp13851, MakeString(""), symshen_4a)


tmp13853 := PrimStringConcat(MakeString(" \n"), tmp13852)

tmp13854 := Call(__e, PrimFunc(symshen_4app), V5402, tmp13853, symshen_4a)


tmp13855 := PrimStringConcat(MakeString("> Inputs to "), tmp13854)

tmp13856 := Call(__e, PrimFunc(symshen_4app), V5401, tmp13855, symshen_4a)


tmp13857 := PrimStringConcat(MakeString("<"), tmp13856)

tmp13858 := Call(__e, PrimFunc(symshen_4app), tmp13850, tmp13857, symshen_4a)


tmp13859 := PrimStringConcat(MakeString("\n"), tmp13858)

tmp13860 := Call(__e, PrimFunc(symstoutput))


tmp13861 := Call(__e, PrimFunc(sympr), tmp13859, tmp13860)


_ = tmp13861

__e.TailApply(PrimFunc(symshen_4recursively_1print), V5403)
return


}, 3)

tmp13862 := Call(__e, ns2_1set, symshen_4input_1track, tmp13849)


_ = tmp13862

tmp13863 := MakeNative(func(__e *ControlFlow) {
V5406 := __e.Get(1)
_ = V5406
tmp13873 := PrimEqual(Nil, V5406)

if True == tmp13873 {
tmp13864 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString(" ==>"), tmp13864)
return


} else {
tmp13871 := PrimIsPair(V5406)

if True == tmp13871 {
tmp13865 := PrimHead(V5406)

tmp13866 := Call(__e, PrimFunc(symprint), tmp13865)


_ = tmp13866

tmp13867 := Call(__e, PrimFunc(symstoutput))


tmp13868 := Call(__e, PrimFunc(sympr), MakeString(", "), tmp13867)


_ = tmp13868

tmp13869 := PrimTail(V5406)

__e.TailApply(PrimFunc(symshen_4recursively_1print), tmp13869)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.recursively-print")))
return
}


}


}, 1)

tmp13874 := Call(__e, ns2_1set, symshen_4recursively_1print, tmp13863)


_ = tmp13874

tmp13875 := MakeNative(func(__e *ControlFlow) {
V5407 := __e.Get(1)
_ = V5407
tmp13879 := PrimEqual(MakeNumber(0), V5407)

if True == tmp13879 {
__e.Return(MakeString(""))
return
} else {
tmp13876 := PrimNumberSubtract(V5407, MakeNumber(1))

tmp13877 := Call(__e, PrimFunc(symshen_4spaces), tmp13876)


__e.Return(PrimStringConcat(MakeString(" "), tmp13877))
return


}


}, 1)

tmp13880 := Call(__e, ns2_1set, symshen_4spaces, tmp13875)


_ = tmp13880

tmp13881 := MakeNative(func(__e *ControlFlow) {
V5408 := __e.Get(1)
_ = V5408
V5409 := __e.Get(2)
_ = V5409
V5410 := __e.Get(3)
_ = V5410
tmp13882 := Call(__e, PrimFunc(symshen_4spaces), V5408)


tmp13883 := Call(__e, PrimFunc(symshen_4spaces), V5408)


tmp13884 := Call(__e, PrimFunc(symshen_4app), V5410, MakeString(""), symshen_4s)


tmp13885 := PrimStringConcat(MakeString("==> "), tmp13884)

tmp13886 := Call(__e, PrimFunc(symshen_4app), tmp13883, tmp13885, symshen_4a)


tmp13887 := PrimStringConcat(MakeString(" \n"), tmp13886)

tmp13888 := Call(__e, PrimFunc(symshen_4app), V5409, tmp13887, symshen_4a)


tmp13889 := PrimStringConcat(MakeString("> Output from "), tmp13888)

tmp13890 := Call(__e, PrimFunc(symshen_4app), V5408, tmp13889, symshen_4a)


tmp13891 := PrimStringConcat(MakeString("<"), tmp13890)

tmp13892 := Call(__e, PrimFunc(symshen_4app), tmp13882, tmp13891, symshen_4a)


tmp13893 := PrimStringConcat(MakeString("\n"), tmp13892)

tmp13894 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp13893, tmp13894)
return


}, 3)

tmp13895 := Call(__e, ns2_1set, symshen_4output_1track, tmp13881)


_ = tmp13895

tmp13896 := MakeNative(func(__e *ControlFlow) {
V5411 := __e.Get(1)
_ = V5411
tmp13897 := PrimValue(symshen_4_dtracking_d)

tmp13898 := Call(__e, PrimFunc(symremove), V5411, tmp13897)


tmp13899 := PrimSet(symshen_4_dtracking_d, tmp13898)

_ = tmp13899

tmp13900 := MakeNative(func(__e *ControlFlow) {
tmp13901 := Call(__e, PrimFunc(symps), V5411)


__e.TailApply(PrimFunc(symeval), tmp13901)
return


}, 0)

tmp13902 := MakeNative(func(__e *ControlFlow) {
Z5412 := __e.Get(1)
_ = Z5412
__e.Return(V5411)
return
}, 1)

tmp13903 := Call(__e, try_1catch, tmp13900, tmp13902)


_ = tmp13903

__e.Return(V5411)
return


}, 1)

tmp13904 := Call(__e, ns2_1set, symuntrack, tmp13896)


_ = tmp13904

tmp13905 := MakeNative(func(__e *ControlFlow) {
V5413 := __e.Get(1)
_ = V5413
V5414 := __e.Get(2)
_ = V5414
__e.TailApply(PrimFunc(symshen_4remove_1h), V5413, V5414, Nil)
return
}, 2)

tmp13906 := Call(__e, ns2_1set, symremove, tmp13905)


_ = tmp13906

tmp13907 := MakeNative(func(__e *ControlFlow) {
V5424 := __e.Get(1)
_ = V5424
V5425 := __e.Get(2)
_ = V5425
V5426 := __e.Get(3)
_ = V5426
tmp13922 := PrimEqual(Nil, V5425)

if True == tmp13922 {
__e.TailApply(PrimFunc(symreverse), V5426)
return
} else {
tmp13920 := PrimIsPair(V5425)

var ifres13916 Obj

if True == tmp13920 {
tmp13918 := PrimHead(V5425)

tmp13919 := PrimEqual(V5424, tmp13918)

var ifres13917 Obj

if True == tmp13919 {
ifres13917 = True


} else {
ifres13917 = False


}

ifres13916 = ifres13917


} else {
ifres13916 = False


}

if True == ifres13916 {
tmp13908 := PrimHead(V5425)

tmp13909 := PrimTail(V5425)

__e.TailApply(PrimFunc(symshen_4remove_1h), tmp13908, tmp13909, V5426)
return


} else {
tmp13914 := PrimIsPair(V5425)

if True == tmp13914 {
tmp13910 := PrimTail(V5425)

tmp13911 := PrimHead(V5425)

tmp13912 := PrimCons(tmp13911, V5426)

__e.TailApply(PrimFunc(symshen_4remove_1h), V5424, tmp13910, tmp13912)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.remove-h")))
return
}


}


}


}, 3)

tmp13923 := Call(__e, ns2_1set, symshen_4remove_1h, tmp13907)


_ = tmp13923

tmp13924 := MakeNative(func(__e *ControlFlow) {
V5427 := __e.Get(1)
_ = V5427
tmp13925 := PrimValue(symshen_4_dprofiled_d)

tmp13926 := PrimCons(V5427, tmp13925)

tmp13927 := PrimSet(symshen_4_dprofiled_d, tmp13926)

_ = tmp13927

tmp13928 := Call(__e, PrimFunc(symps), V5427)


__e.TailApply(PrimFunc(symshen_4profile_1help), tmp13928)
return


}, 1)

tmp13929 := Call(__e, ns2_1set, symprofile, tmp13924)


_ = tmp13929

tmp13930 := MakeNative(func(__e *ControlFlow) {
V5430 := __e.Get(1)
_ = V5430
tmp14000 := PrimIsPair(V5430)

var ifres13974 Obj

if True == tmp14000 {
tmp13998 := PrimHead(V5430)

tmp13999 := PrimEqual(symdefun, tmp13998)

var ifres13976 Obj

if True == tmp13999 {
tmp13996 := PrimTail(V5430)

tmp13997 := PrimIsPair(tmp13996)

var ifres13978 Obj

if True == tmp13997 {
tmp13993 := PrimTail(V5430)

tmp13994 := PrimTail(tmp13993)

tmp13995 := PrimIsPair(tmp13994)

var ifres13980 Obj

if True == tmp13995 {
tmp13989 := PrimTail(V5430)

tmp13990 := PrimTail(tmp13989)

tmp13991 := PrimTail(tmp13990)

tmp13992 := PrimIsPair(tmp13991)

var ifres13982 Obj

if True == tmp13992 {
tmp13984 := PrimTail(V5430)

tmp13985 := PrimTail(tmp13984)

tmp13986 := PrimTail(tmp13985)

tmp13987 := PrimTail(tmp13986)

tmp13988 := PrimEqual(Nil, tmp13987)

var ifres13983 Obj

if True == tmp13988 {
ifres13983 = True


} else {
ifres13983 = False


}

ifres13982 = ifres13983


} else {
ifres13982 = False


}

var ifres13981 Obj

if True == ifres13982 {
ifres13981 = True


} else {
ifres13981 = False


}

ifres13980 = ifres13981


} else {
ifres13980 = False


}

var ifres13979 Obj

if True == ifres13980 {
ifres13979 = True


} else {
ifres13979 = False


}

ifres13978 = ifres13979


} else {
ifres13978 = False


}

var ifres13977 Obj

if True == ifres13978 {
ifres13977 = True


} else {
ifres13977 = False


}

ifres13976 = ifres13977


} else {
ifres13976 = False


}

var ifres13975 Obj

if True == ifres13976 {
ifres13975 = True


} else {
ifres13975 = False


}

ifres13974 = ifres13975


} else {
ifres13974 = False


}

if True == ifres13974 {
tmp13931 := MakeNative(func(__e *ControlFlow) {
W5431 := __e.Get(1)
_ = W5431
tmp13932 := MakeNative(func(__e *ControlFlow) {
W5432 := __e.Get(1)
_ = W5432
tmp13933 := MakeNative(func(__e *ControlFlow) {
W5433 := __e.Get(1)
_ = W5433
tmp13934 := MakeNative(func(__e *ControlFlow) {
W5434 := __e.Get(1)
_ = W5434
tmp13935 := MakeNative(func(__e *ControlFlow) {
W5435 := __e.Get(1)
_ = W5435
tmp13936 := PrimTail(V5430)

__e.Return(PrimHead(tmp13936))
return


}, 1)

tmp13937 := Call(__e, PrimFunc(symeval_1kl), W5433)


__e.TailApply(tmp13935, tmp13937)
return


}, 1)

tmp13938 := Call(__e, PrimFunc(symeval_1kl), W5432)


__e.TailApply(tmp13934, tmp13938)
return


}, 1)

tmp13939 := PrimTail(V5430)

tmp13940 := PrimTail(tmp13939)

tmp13941 := PrimHead(tmp13940)

tmp13942 := PrimTail(V5430)

tmp13943 := PrimHead(tmp13942)

tmp13944 := PrimTail(V5430)

tmp13945 := PrimTail(tmp13944)

tmp13946 := PrimTail(tmp13945)

tmp13947 := PrimHead(tmp13946)

tmp13948 := Call(__e, PrimFunc(symsubst), W5431, tmp13943, tmp13947)


tmp13949 := PrimCons(tmp13948, Nil)

tmp13950 := PrimCons(tmp13941, tmp13949)

tmp13951 := PrimCons(W5431, tmp13950)

tmp13952 := PrimCons(symdefun, tmp13951)

__e.TailApply(tmp13933, tmp13952)
return


}, 1)

tmp13953 := PrimTail(V5430)

tmp13954 := PrimHead(tmp13953)

tmp13955 := PrimTail(V5430)

tmp13956 := PrimTail(tmp13955)

tmp13957 := PrimHead(tmp13956)

tmp13958 := PrimTail(V5430)

tmp13959 := PrimHead(tmp13958)

tmp13960 := PrimTail(V5430)

tmp13961 := PrimTail(tmp13960)

tmp13962 := PrimHead(tmp13961)

tmp13963 := PrimTail(V5430)

tmp13964 := PrimTail(tmp13963)

tmp13965 := PrimHead(tmp13964)

tmp13966 := PrimCons(W5431, tmp13965)

tmp13967 := Call(__e, PrimFunc(symshen_4profile_1func), tmp13959, tmp13962, tmp13966)


tmp13968 := PrimCons(tmp13967, Nil)

tmp13969 := PrimCons(tmp13957, tmp13968)

tmp13970 := PrimCons(tmp13954, tmp13969)

tmp13971 := PrimCons(symdefun, tmp13970)

__e.TailApply(tmp13932, tmp13971)
return


}, 1)

tmp13972 := Call(__e, PrimFunc(symgensym), symshen_4f)


__e.TailApply(tmp13931, tmp13972)
return


} else {
__e.Return(PrimSimpleError(MakeString("Cannot profile.\n")))
return
}


}, 1)

tmp14001 := Call(__e, ns2_1set, symshen_4profile_1help, tmp13930)


_ = tmp14001

tmp14002 := MakeNative(func(__e *ControlFlow) {
V5436 := __e.Get(1)
_ = V5436
tmp14003 := PrimValue(symshen_4_dprofiled_d)

tmp14004 := Call(__e, PrimFunc(symremove), V5436, tmp14003)


tmp14005 := PrimSet(symshen_4_dprofiled_d, tmp14004)

_ = tmp14005

tmp14006 := MakeNative(func(__e *ControlFlow) {
tmp14007 := Call(__e, PrimFunc(symps), V5436)


__e.TailApply(PrimFunc(symeval), tmp14007)
return


}, 0)

tmp14008 := MakeNative(func(__e *ControlFlow) {
Z5437 := __e.Get(1)
_ = Z5437
__e.Return(V5436)
return
}, 1)

__e.TailApply(try_1catch, tmp14006, tmp14008)
return


}, 1)

tmp14009 := Call(__e, ns2_1set, symunprofile, tmp14002)


_ = tmp14009

tmp14010 := MakeNative(func(__e *ControlFlow) {
V5438 := __e.Get(1)
_ = V5438
tmp14011 := PrimValue(symshen_4_dprofiled_d)

__e.TailApply(PrimFunc(symelement_2), V5438, tmp14011)
return


}, 1)

tmp14012 := Call(__e, ns2_1set, symshen_4profiled_2, tmp14010)


_ = tmp14012

tmp14013 := MakeNative(func(__e *ControlFlow) {
V5439 := __e.Get(1)
_ = V5439
V5440 := __e.Get(2)
_ = V5440
V5441 := __e.Get(3)
_ = V5441
tmp14014 := PrimCons(symrun, Nil)

tmp14015 := PrimCons(symget_1time, tmp14014)

tmp14016 := PrimCons(symrun, Nil)

tmp14017 := PrimCons(symget_1time, tmp14016)

tmp14018 := PrimCons(symStart, Nil)

tmp14019 := PrimCons(tmp14017, tmp14018)

tmp14020 := PrimCons(sym_1, tmp14019)

tmp14021 := PrimCons(V5439, Nil)

tmp14022 := PrimCons(symshen_4get_1profile, tmp14021)

tmp14023 := PrimCons(symFinish, Nil)

tmp14024 := PrimCons(tmp14022, tmp14023)

tmp14025 := PrimCons(sym_7, tmp14024)

tmp14026 := PrimCons(tmp14025, Nil)

tmp14027 := PrimCons(V5439, tmp14026)

tmp14028 := PrimCons(symshen_4put_1profile, tmp14027)

tmp14029 := PrimCons(symResult, Nil)

tmp14030 := PrimCons(tmp14028, tmp14029)

tmp14031 := PrimCons(symRecord, tmp14030)

tmp14032 := PrimCons(symlet, tmp14031)

tmp14033 := PrimCons(tmp14032, Nil)

tmp14034 := PrimCons(tmp14020, tmp14033)

tmp14035 := PrimCons(symFinish, tmp14034)

tmp14036 := PrimCons(symlet, tmp14035)

tmp14037 := PrimCons(tmp14036, Nil)

tmp14038 := PrimCons(V5441, tmp14037)

tmp14039 := PrimCons(symResult, tmp14038)

tmp14040 := PrimCons(symlet, tmp14039)

tmp14041 := PrimCons(tmp14040, Nil)

tmp14042 := PrimCons(tmp14015, tmp14041)

tmp14043 := PrimCons(symStart, tmp14042)

__e.Return(PrimCons(symlet, tmp14043))
return


}, 3)

tmp14044 := Call(__e, ns2_1set, symshen_4profile_1func, tmp14013)


_ = tmp14044

tmp14045 := MakeNative(func(__e *ControlFlow) {
V5442 := __e.Get(1)
_ = V5442
tmp14046 := MakeNative(func(__e *ControlFlow) {
W5443 := __e.Get(1)
_ = W5443
tmp14047 := MakeNative(func(__e *ControlFlow) {
W5444 := __e.Get(1)
_ = W5444
__e.TailApply(PrimFunc(sym_8p), V5442, W5443)
return
}, 1)

tmp14048 := Call(__e, PrimFunc(symshen_4put_1profile), V5442, MakeNumber(0))


__e.TailApply(tmp14047, tmp14048)
return


}, 1)

tmp14049 := Call(__e, PrimFunc(symshen_4get_1profile), V5442)


__e.TailApply(tmp14046, tmp14049)
return


}, 1)

tmp14050 := Call(__e, ns2_1set, symprofile_1results, tmp14045)


_ = tmp14050

tmp14051 := MakeNative(func(__e *ControlFlow) {
V5445 := __e.Get(1)
_ = V5445
tmp14052 := MakeNative(func(__e *ControlFlow) {
tmp14053 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V5445, symprofile, tmp14053)
return


}, 0)

tmp14054 := MakeNative(func(__e *ControlFlow) {
Z5446 := __e.Get(1)
_ = Z5446
__e.Return(MakeNumber(0))
return
}, 1)

__e.TailApply(try_1catch, tmp14052, tmp14054)
return


}, 1)

tmp14055 := Call(__e, ns2_1set, symshen_4get_1profile, tmp14051)


_ = tmp14055

tmp14056 := MakeNative(func(__e *ControlFlow) {
V5447 := __e.Get(1)
_ = V5447
V5448 := __e.Get(2)
_ = V5448
tmp14057 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V5447, symprofile, V5448, tmp14057)
return


}, 2)

__e.TailApply(ns2_1set, symshen_4put_1profile, tmp14056)
return




}, 0)

