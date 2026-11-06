package main

import . "github.com/pyrex41/shen-go/kl"

var YaccMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp16665 := MakeNative2(func(__e *ControlFlow, V112, V113 Obj) Obj {
tmp16666 := Call1(__e, V112, V113)


let__16527 := tmp16666
_ = let__16527

tmp16673 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16527)


if True == tmp16673 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("parse failure\n"))
} else {
tmp16671 := Call1(__e, SymFn(symshen_4partial_1parse_1failure_2), let__16527)


if True == tmp16671 {
tmp16667 := Call1(__e, SymFn(symshen_4in_1_6), let__16527)


tmp16668 := PrimGuard2(__e, symset, PrimSet, symshen_4_dresidue_d, tmp16667)

_ = tmp16668

tmp16669 := GValue(__e, symshen_4_dresidue_d)

return TailCall1(__e, SymFn(symshen_4raise_1syntax_1error), tmp16669)


} else {
return TailCall1(__e, SymFn(symshen_4_5_1out), let__16527)
}


}


panic("unreachable")
})

tmp16674 := Call2(__e, ns2_1set, symcompile, tmp16665)


_ = tmp16674

tmp16675 := MakeNative1(func(__e *ControlFlow, V115 Obj) Obj {
tmp16676 := GValue(__e, sym_dmaximum_1print_1sequence_1size_d)

tmp16677 := Call3(__e, SymFn(symshen_4syntax_1error_1message), tmp16676, MakeInteger(0), V115)


tmp16679 := Call1(__e, SymFn(symshen_4proc_1nl), PrimGuard2(__e, symcn, PrimStringConcat, MakeString("syntax error here: "), tmp16677))


return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp16679)


panic("unreachable")
})

tmp16680 := Call2(__e, ns2_1set, symshen_4raise_1syntax_1error, tmp16675)


_ = tmp16680

tmp16681 := MakeNative3(func(__e *ControlFlow, V123, V124, V125 Obj) Obj {
tmp16692 := GEq(__e, Nil, V125)

if True == tmp16692 {
return MakeString("\n")
} else {
tmp16690 := GEq(__e, V123, V124)

if True == tmp16690 {
return MakeString("...etc \n")
} else {
tmp16688 := GConsP(__e, V125)

if True == tmp16688 {
tmp16682 := GHd(__e, V125)

tmp16683 := Call3(__e, SymFn(symshen_4app), tmp16682, MakeString(" "), symshen_4s)


tmp16684 := GAdd(__e, V124, MakeInteger(1))

tmp16685 := GTl(__e, V125)

tmp16686 := Call3(__e, SymFn(symshen_4syntax_1error_1message), V123, tmp16684, tmp16685)


return PrimGuard2(__e, symcn, PrimStringConcat, tmp16683, tmp16686)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4syntax_1error_1message)
}


}


}


panic("unreachable")
})

tmp16693 := Call2(__e, ns2_1set, symshen_4syntax_1error_1message, tmp16681)


_ = tmp16693

tmp16694 := MakeNative1(func(__e *ControlFlow, V126 Obj) Obj {
tmp16695 := Call0(__e, SymFn(symfail))


return GEq(__e, V126, tmp16695)


panic("unreachable")
})

tmp16696 := Call2(__e, ns2_1set, symshen_4parse_1failure_2, tmp16694)


_ = tmp16696

tmp16697 := MakeNative1(func(__e *ControlFlow, V127 Obj) Obj {
tmp16698 := Call1(__e, SymFn(symshen_4in_1_6), V127)


return GConsP(__e, tmp16698)


panic("unreachable")
})

tmp16699 := Call2(__e, ns2_1set, symshen_4partial_1parse_1failure_2, tmp16697)


_ = tmp16699

tmp16700 := MakeNative1(func(__e *ControlFlow, V130 Obj) Obj {
tmp16711 := GConsP(__e, V130)

var ifres16704 Obj

if True == tmp16711 {
tmp16709 := GTl(__e, V130)

tmp16710 := GConsP(__e, tmp16709)

var ifres16705 Obj

if True == tmp16710 {
tmp16706 := GTl(__e, V130)

tmp16707 := GTl(__e, tmp16706)

tmp16708 := GEq(__e, Nil, tmp16707)

ifres16705 = tmp16708


} else {
ifres16705 = False


}

ifres16704 = ifres16705


} else {
ifres16704 = False


}

if True == ifres16704 {
tmp16701 := GTl(__e, V130)

return GHd(__e, tmp16701)


} else {
tmp16702 := Call3(__e, SymFn(symshen_4app), V130, MakeString(" is not a YACC stream\n"), symshen_4s)


return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp16702)


}


panic("unreachable")
})

tmp16712 := Call2(__e, ns2_1set, symshen_4objectcode, tmp16700)


_ = tmp16712

tmp16713 := MakeNative1(func(__e *ControlFlow, V131 Obj) Obj {
tmp16714 := MakeNative1(func(__e *ControlFlow, Z132 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5yacc_6), Z132)
panic("unreachable")
})

return TailCall2(__e, SymFn(symcompile), tmp16714, V131)


panic("unreachable")
})

tmp16715 := Call2(__e, ns2_1set, symshen_4yacc_1_6shen, tmp16713)


_ = tmp16715

tmp16716 := MakeNative1(func(__e *ControlFlow, V133 Obj) Obj {
tmp16743 := GConsP(__e, V133)

var ifres16717 Obj

if True == tmp16743 {
tmp16718 := Call1(__e, SymFn(symhead), V133)


let__16529 := tmp16718
_ = let__16529

tmp16719 := Call1(__e, SymFn(symtail), V133)


let__16530 := tmp16719
_ = let__16530

tmp16720 := Call1(__e, SymFn(symshen_4_5yaccsig_6), let__16530)


let__16531 := tmp16720
_ = let__16531

tmp16741 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16531)


var ifres16721 Obj

if True == tmp16741 {
tmp16722 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16721 = tmp16722


} else {
tmp16723 := Call1(__e, SymFn(symshen_4_5_1out), let__16531)


let__16532 := tmp16723
_ = let__16532

tmp16724 := Call1(__e, SymFn(symshen_4in_1_6), let__16531)


let__16533 := tmp16724
_ = let__16533

tmp16725 := Call1(__e, SymFn(symshen_4_5c_1rules_6), let__16533)


let__16534 := tmp16725
_ = let__16534

tmp16740 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16534)


var ifres16726 Obj

if True == tmp16740 {
tmp16727 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16726 = tmp16727


} else {
tmp16728 := Call1(__e, SymFn(symshen_4_5_1out), let__16534)


let__16535 := tmp16728
_ = let__16535

tmp16729 := Call1(__e, SymFn(symshen_4in_1_6), let__16534)


let__16536 := tmp16729
_ = let__16536

tmp16730 := Call1(__e, SymFn(symgensym), symS)


let__16537 := tmp16730
_ = let__16537

tmp16731 := GCons(__e, let__16529, Nil)

tmp16732 := GCons(__e, symdefine, tmp16731)

tmp16733 := Call3(__e, SymFn(symshen_4c_1rules_1_6shen), let__16532, let__16537, let__16535)


tmp16734 := GCons(__e, tmp16733, Nil)

tmp16735 := GCons(__e, sym_1_6, tmp16734)

tmp16736 := GCons(__e, let__16537, tmp16735)

tmp16737 := Call2(__e, SymFn(symappend), let__16532, tmp16736)


tmp16738 := Call2(__e, SymFn(symappend), tmp16732, tmp16737)


let__16538 := tmp16738
_ = let__16538

tmp16739 := Call2(__e, SymFn(symshen_4comb), let__16536, let__16538)


ifres16726 = tmp16739


}

ifres16721 = ifres16726


}

ifres16717 = ifres16721


} else {
tmp16742 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16717 = tmp16742


}

let__16528 := ifres16717
_ = let__16528

tmp16745 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16528)


if True == tmp16745 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16528
}


panic("unreachable")
})

tmp16746 := Call2(__e, ns2_1set, symshen_4_5yacc_6, tmp16716)


_ = tmp16746

tmp16747 := MakeNative1(func(__e *ControlFlow, V145 Obj) Obj {
tmp16803 := GConsP(__e, V145)

var ifres16748 Obj

if True == tmp16803 {
tmp16749 := Call1(__e, SymFn(symhead), V145)


let__16540 := tmp16749
_ = let__16540

tmp16750 := Call1(__e, SymFn(symtail), V145)


let__16541 := tmp16750
_ = let__16541

tmp16801 := Call1(__e, SymFn(symshen_4ccons_2), let__16541)


var ifres16751 Obj

if True == tmp16801 {
tmp16752 := Call1(__e, SymFn(symhead), let__16541)


let__16542 := tmp16752
_ = let__16542

tmp16753 := Call1(__e, SymFn(symtail), let__16541)


let__16543 := tmp16753
_ = let__16543

tmp16799 := Call2(__e, SymFn(symshen_4hds_a_2), let__16542, symlist)


var ifres16754 Obj

if True == tmp16799 {
tmp16755 := Call1(__e, SymFn(symtail), let__16542)


let__16544 := tmp16755
_ = let__16544

tmp16797 := GConsP(__e, let__16544)

var ifres16756 Obj

if True == tmp16797 {
tmp16757 := Call1(__e, SymFn(symhead), let__16544)


let__16545 := tmp16757
_ = let__16545

tmp16758 := Call1(__e, SymFn(symtail), let__16544)


let__16546 := tmp16758
_ = let__16546

tmp16759 := Call1(__e, SymFn(sym_5end_6), let__16546)


let__16547 := tmp16759
_ = let__16547

tmp16795 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16547)


var ifres16760 Obj

if True == tmp16795 {
tmp16761 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16760 = tmp16761


} else {
tmp16762 := Call1(__e, SymFn(symshen_4in_1_6), let__16547)


let__16548 := tmp16762
_ = let__16548

tmp16794 := Call2(__e, SymFn(symshen_4hds_a_2), let__16543, sym_a_a_6)


var ifres16763 Obj

if True == tmp16794 {
tmp16764 := Call1(__e, SymFn(symtail), let__16543)


let__16549 := tmp16764
_ = let__16549

tmp16792 := GConsP(__e, let__16549)

var ifres16765 Obj

if True == tmp16792 {
tmp16766 := Call1(__e, SymFn(symhead), let__16549)


let__16550 := tmp16766
_ = let__16550

tmp16767 := Call1(__e, SymFn(symtail), let__16549)


let__16551 := tmp16767
_ = let__16551

tmp16790 := GConsP(__e, let__16551)

var ifres16768 Obj

if True == tmp16790 {
tmp16769 := Call1(__e, SymFn(symhead), let__16551)


let__16552 := tmp16769
_ = let__16552

tmp16770 := Call1(__e, SymFn(symtail), let__16551)


let__16553 := tmp16770
_ = let__16553

tmp16788 := GEq(__e, sym_i, let__16540)

var ifres16786 Obj

if True == tmp16788 {
tmp16787 := GEq(__e, sym_j, let__16552)

ifres16786 = tmp16787


} else {
ifres16786 = False


}

var ifres16771 Obj

if True == ifres16786 {
tmp16772 := GCons(__e, let__16545, Nil)

tmp16773 := GCons(__e, symlist, tmp16772)

tmp16774 := GCons(__e, let__16545, Nil)

tmp16775 := GCons(__e, symlist, tmp16774)

tmp16776 := GCons(__e, let__16550, Nil)

tmp16777 := GCons(__e, tmp16775, tmp16776)

tmp16778 := GCons(__e, symstr, tmp16777)

tmp16779 := GCons(__e, sym_j, Nil)

tmp16780 := GCons(__e, tmp16778, tmp16779)

tmp16781 := GCons(__e, sym_1_1_6, tmp16780)

tmp16782 := GCons(__e, tmp16773, tmp16781)

tmp16783 := GCons(__e, sym_i, tmp16782)

tmp16784 := Call2(__e, SymFn(symshen_4comb), let__16553, tmp16783)


ifres16771 = tmp16784


} else {
tmp16785 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16771 = tmp16785


}

ifres16768 = ifres16771


} else {
tmp16789 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16768 = tmp16789


}

ifres16765 = ifres16768


} else {
tmp16791 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16765 = tmp16791


}

ifres16763 = ifres16765


} else {
tmp16793 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16763 = tmp16793


}

ifres16760 = ifres16763


}

ifres16756 = ifres16760


} else {
tmp16796 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16756 = tmp16796


}

ifres16754 = ifres16756


} else {
tmp16798 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16754 = tmp16798


}

ifres16751 = ifres16754


} else {
tmp16800 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16751 = tmp16800


}

ifres16748 = ifres16751


} else {
tmp16802 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16748 = tmp16802


}

let__16539 := ifres16748
_ = let__16539

tmp16813 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16539)


if True == tmp16813 {
tmp16804 := Call1(__e, SymFn(sym_5e_6), V145)


let__16555 := tmp16804
_ = let__16555

tmp16809 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16555)


var ifres16805 Obj

if True == tmp16809 {
tmp16806 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16805 = tmp16806


} else {
tmp16807 := Call1(__e, SymFn(symshen_4in_1_6), let__16555)


let__16556 := tmp16807
_ = let__16556

tmp16808 := Call2(__e, SymFn(symshen_4comb), let__16556, Nil)


ifres16805 = tmp16808


}

let__16554 := ifres16805
_ = let__16554

tmp16811 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16554)


if True == tmp16811 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16554
}


} else {
return let__16539
}


panic("unreachable")
})

tmp16814 := Call2(__e, ns2_1set, symshen_4_5yaccsig_6, tmp16747)


_ = tmp16814

tmp16815 := MakeNative1(func(__e *ControlFlow, V164 Obj) Obj {
tmp16816 := Call1(__e, SymFn(symshen_4_5c_1rule_6), V164)


let__16558 := tmp16816
_ = let__16558

tmp16829 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16558)


var ifres16817 Obj

if True == tmp16829 {
tmp16818 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16817 = tmp16818


} else {
tmp16819 := Call1(__e, SymFn(symshen_4_5_1out), let__16558)


let__16559 := tmp16819
_ = let__16559

tmp16820 := Call1(__e, SymFn(symshen_4in_1_6), let__16558)


let__16560 := tmp16820
_ = let__16560

tmp16821 := Call1(__e, SymFn(symshen_4_5c_1rules_6), let__16560)


let__16561 := tmp16821
_ = let__16561

tmp16828 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16561)


var ifres16822 Obj

if True == tmp16828 {
tmp16823 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16822 = tmp16823


} else {
tmp16824 := Call1(__e, SymFn(symshen_4_5_1out), let__16561)


let__16562 := tmp16824
_ = let__16562

tmp16825 := Call1(__e, SymFn(symshen_4in_1_6), let__16561)


let__16563 := tmp16825
_ = let__16563

tmp16826 := GCons(__e, let__16559, let__16562)

tmp16827 := Call2(__e, SymFn(symshen_4comb), let__16563, tmp16826)


ifres16822 = tmp16827


}

ifres16817 = ifres16822


}

let__16557 := ifres16817
_ = let__16557

tmp16845 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16557)


if True == tmp16845 {
tmp16830 := Call1(__e, SymFn(sym_5_b_6), V164)


let__16565 := tmp16830
_ = let__16565

tmp16841 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16565)


var ifres16831 Obj

if True == tmp16841 {
tmp16832 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16831 = tmp16832


} else {
tmp16833 := Call1(__e, SymFn(symshen_4_5_1out), let__16565)


let__16566 := tmp16833
_ = let__16566

tmp16834 := Call1(__e, SymFn(symshen_4in_1_6), let__16565)


let__16567 := tmp16834
_ = let__16567

tmp16839 := Call1(__e, SymFn(symempty_2), let__16566)


var ifres16835 Obj

if True == tmp16839 {
ifres16835 = Nil


} else {
tmp16836 := Call3(__e, SymFn(symshen_4app), let__16566, MakeString("\n ..."), symshen_4r)


tmp16838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("YACC syntax error here:\n "), tmp16836))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("YACC syntax error here:\n "), tmp16836)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()

ifres16835 = tmp16838


}

tmp16840 := Call2(__e, SymFn(symshen_4comb), let__16567, ifres16835)


ifres16831 = tmp16840


}

let__16564 := ifres16831
_ = let__16564

tmp16843 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16564)


if True == tmp16843 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16564
}


} else {
return let__16557
}


panic("unreachable")
})

tmp16846 := Call2(__e, ns2_1set, symshen_4_5c_1rules_6, tmp16815)


_ = tmp16846

tmp16847 := MakeNative1(func(__e *ControlFlow, V176 Obj) Obj {
tmp16848 := Call1(__e, SymFn(symshen_4_5syntax_6), V176)


let__16569 := tmp16848
_ = let__16569

tmp16867 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16569)


var ifres16849 Obj

if True == tmp16867 {
tmp16850 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16849 = tmp16850


} else {
tmp16851 := Call1(__e, SymFn(symshen_4_5_1out), let__16569)


let__16570 := tmp16851
_ = let__16570

tmp16852 := Call1(__e, SymFn(symshen_4in_1_6), let__16569)


let__16571 := tmp16852
_ = let__16571

tmp16853 := Call1(__e, SymFn(symshen_4_5semantics_6), let__16571)


let__16572 := tmp16853
_ = let__16572

tmp16866 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16572)


var ifres16854 Obj

if True == tmp16866 {
tmp16855 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16854 = tmp16855


} else {
tmp16856 := Call1(__e, SymFn(symshen_4_5_1out), let__16572)


let__16573 := tmp16856
_ = let__16573

tmp16857 := Call1(__e, SymFn(symshen_4in_1_6), let__16572)


let__16574 := tmp16857
_ = let__16574

tmp16858 := Call1(__e, SymFn(symshen_4_5sc_6), let__16574)


let__16575 := tmp16858
_ = let__16575

tmp16865 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16575)


var ifres16859 Obj

if True == tmp16865 {
tmp16860 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16859 = tmp16860


} else {
tmp16861 := Call1(__e, SymFn(symshen_4in_1_6), let__16575)


let__16576 := tmp16861
_ = let__16576

tmp16862 := GCons(__e, let__16573, Nil)

tmp16863 := GCons(__e, let__16570, tmp16862)

tmp16864 := Call2(__e, SymFn(symshen_4comb), let__16576, tmp16863)


ifres16859 = tmp16864


}

ifres16854 = ifres16859


}

ifres16849 = ifres16854


}

let__16568 := ifres16849
_ = let__16568

tmp16886 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16568)


if True == tmp16886 {
tmp16868 := Call1(__e, SymFn(symshen_4_5syntax_6), V176)


let__16578 := tmp16868
_ = let__16578

tmp16882 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16578)


var ifres16869 Obj

if True == tmp16882 {
tmp16870 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16869 = tmp16870


} else {
tmp16871 := Call1(__e, SymFn(symshen_4_5_1out), let__16578)


let__16579 := tmp16871
_ = let__16579

tmp16872 := Call1(__e, SymFn(symshen_4in_1_6), let__16578)


let__16580 := tmp16872
_ = let__16580

tmp16873 := Call1(__e, SymFn(symshen_4_5sc_6), let__16580)


let__16581 := tmp16873
_ = let__16581

tmp16881 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16581)


var ifres16874 Obj

if True == tmp16881 {
tmp16875 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16874 = tmp16875


} else {
tmp16876 := Call1(__e, SymFn(symshen_4in_1_6), let__16581)


let__16582 := tmp16876
_ = let__16582

tmp16877 := Call1(__e, SymFn(symshen_4autocomplete), let__16579)


tmp16878 := GCons(__e, tmp16877, Nil)

tmp16879 := GCons(__e, let__16579, tmp16878)

tmp16880 := Call2(__e, SymFn(symshen_4comb), let__16582, tmp16879)


ifres16874 = tmp16880


}

ifres16869 = ifres16874


}

let__16577 := ifres16869
_ = let__16577

tmp16884 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16577)


if True == tmp16884 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16577
}


} else {
return let__16568
}


panic("unreachable")
})

tmp16887 := Call2(__e, ns2_1set, symshen_4_5c_1rule_6, tmp16847)


_ = tmp16887

tmp16888 := MakeNative1(func(__e *ControlFlow, V192 Obj) Obj {
tmp16914 := GConsP(__e, V192)

var ifres16908 Obj

if True == tmp16914 {
tmp16912 := GTl(__e, V192)

tmp16913 := GEq(__e, Nil, tmp16912)

var ifres16909 Obj

if True == tmp16913 {
tmp16910 := GHd(__e, V192)

tmp16911 := Call1(__e, SymFn(symshen_4non_1terminal_2), tmp16910)


ifres16909 = tmp16911


} else {
ifres16909 = False


}

ifres16908 = ifres16909


} else {
ifres16908 = False


}

if True == ifres16908 {
return GHd(__e, V192)
} else {
tmp16906 := GConsP(__e, V192)

var ifres16903 Obj

if True == tmp16906 {
tmp16904 := GHd(__e, V192)

tmp16905 := Call1(__e, SymFn(symshen_4non_1terminal_2), tmp16904)


ifres16903 = tmp16905


} else {
ifres16903 = False


}

if True == ifres16903 {
tmp16889 := GHd(__e, V192)

tmp16890 := GTl(__e, V192)

tmp16891 := Call1(__e, SymFn(symshen_4autocomplete), tmp16890)


tmp16892 := GCons(__e, tmp16891, Nil)

tmp16893 := GCons(__e, tmp16889, tmp16892)

return GCons(__e, symappend, tmp16893)


} else {
tmp16901 := GConsP(__e, V192)

if True == tmp16901 {
tmp16894 := GHd(__e, V192)

tmp16895 := Call1(__e, SymFn(symshen_4autocomplete), tmp16894)


tmp16896 := GTl(__e, V192)

tmp16897 := Call1(__e, SymFn(symshen_4autocomplete), tmp16896)


tmp16898 := GCons(__e, tmp16897, Nil)

tmp16899 := GCons(__e, tmp16895, tmp16898)

return GCons(__e, symcons, tmp16899)


} else {
return V192
}


}


}


panic("unreachable")
})

tmp16915 := Call2(__e, ns2_1set, symshen_4autocomplete, tmp16888)


_ = tmp16915

tmp16916 := MakeNative1(func(__e *ControlFlow, V193 Obj) Obj {
tmp16920 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V193)

if True == tmp16920 {
tmp16917 := Call1(__e, SymFn(symexplode), V193)


let__16583 := tmp16917
_ = let__16583

tmp16918 := MakeNative1(func(__e *ControlFlow, Z195 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5non_1terminal_2_6), Z195)
panic("unreachable")
})

return TailCall2(__e, SymFn(symcompile), tmp16918, let__16583)


} else {
return False
}


panic("unreachable")
})

tmp16921 := Call2(__e, ns2_1set, symshen_4non_1terminal_2, tmp16916)


_ = tmp16921

tmp16922 := MakeNative1(func(__e *ControlFlow, V196 Obj) Obj {
tmp16923 := Call1(__e, SymFn(symshen_4_5packagenames_6), V196)


let__16585 := tmp16923
_ = let__16585

tmp16933 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16585)


var ifres16924 Obj

if True == tmp16933 {
tmp16925 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16924 = tmp16925


} else {
tmp16926 := Call1(__e, SymFn(symshen_4in_1_6), let__16585)


let__16586 := tmp16926
_ = let__16586

tmp16927 := Call1(__e, SymFn(symshen_4_5non_1terminal_1name_6), let__16586)


let__16587 := tmp16927
_ = let__16587

tmp16932 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16587)


var ifres16928 Obj

if True == tmp16932 {
tmp16929 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16928 = tmp16929


} else {
tmp16930 := Call1(__e, SymFn(symshen_4in_1_6), let__16587)


let__16588 := tmp16930
_ = let__16588

tmp16931 := Call2(__e, SymFn(symshen_4comb), let__16588, True)


ifres16928 = tmp16931


}

ifres16924 = ifres16928


}

let__16584 := ifres16924
_ = let__16584

tmp16951 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16584)


if True == tmp16951 {
tmp16934 := Call1(__e, SymFn(symshen_4_5non_1terminal_1name_6), V196)


let__16590 := tmp16934
_ = let__16590

tmp16939 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16590)


var ifres16935 Obj

if True == tmp16939 {
tmp16936 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16935 = tmp16936


} else {
tmp16937 := Call1(__e, SymFn(symshen_4in_1_6), let__16590)


let__16591 := tmp16937
_ = let__16591

tmp16938 := Call2(__e, SymFn(symshen_4comb), let__16591, True)


ifres16935 = tmp16938


}

let__16589 := ifres16935
_ = let__16589

tmp16949 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16589)


if True == tmp16949 {
tmp16940 := Call1(__e, SymFn(sym_5_b_6), V196)


let__16593 := tmp16940
_ = let__16593

tmp16945 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16593)


var ifres16941 Obj

if True == tmp16945 {
tmp16942 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16941 = tmp16942


} else {
tmp16943 := Call1(__e, SymFn(symshen_4in_1_6), let__16593)


let__16594 := tmp16943
_ = let__16594

tmp16944 := Call2(__e, SymFn(symshen_4comb), let__16594, False)


ifres16941 = tmp16944


}

let__16592 := ifres16941
_ = let__16592

tmp16947 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16592)


if True == tmp16947 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16592
}


} else {
return let__16589
}


} else {
return let__16584
}


panic("unreachable")
})

tmp16952 := Call2(__e, ns2_1set, symshen_4_5non_1terminal_2_6, tmp16922)


_ = tmp16952

tmp16953 := MakeNative1(func(__e *ControlFlow, V208 Obj) Obj {
tmp16954 := Call1(__e, SymFn(symshen_4_5packagename_6), V208)


let__16596 := tmp16954
_ = let__16596

tmp16968 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16596)


var ifres16955 Obj

if True == tmp16968 {
tmp16956 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16955 = tmp16956


} else {
tmp16957 := Call1(__e, SymFn(symshen_4in_1_6), let__16596)


let__16597 := tmp16957
_ = let__16597

tmp16967 := Call2(__e, SymFn(symshen_4hds_a_2), let__16597, MakeString("."))


var ifres16958 Obj

if True == tmp16967 {
tmp16959 := Call1(__e, SymFn(symtail), let__16597)


let__16598 := tmp16959
_ = let__16598

tmp16960 := Call1(__e, SymFn(symshen_4_5packagenames_6), let__16598)


let__16599 := tmp16960
_ = let__16599

tmp16965 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16599)


var ifres16961 Obj

if True == tmp16965 {
tmp16962 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16961 = tmp16962


} else {
tmp16963 := Call1(__e, SymFn(symshen_4in_1_6), let__16599)


let__16600 := tmp16963
_ = let__16600

tmp16964 := Call2(__e, SymFn(symshen_4comb), let__16600, symshen_4skip)


ifres16961 = tmp16964


}

ifres16958 = ifres16961


} else {
tmp16966 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16958 = tmp16966


}

ifres16955 = ifres16958


}

let__16595 := ifres16955
_ = let__16595

tmp16982 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16595)


if True == tmp16982 {
tmp16969 := Call1(__e, SymFn(symshen_4_5packagename_6), V208)


let__16602 := tmp16969
_ = let__16602

tmp16978 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16602)


var ifres16970 Obj

if True == tmp16978 {
tmp16971 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16970 = tmp16971


} else {
tmp16972 := Call1(__e, SymFn(symshen_4in_1_6), let__16602)


let__16603 := tmp16972
_ = let__16603

tmp16977 := Call2(__e, SymFn(symshen_4hds_a_2), let__16603, MakeString("."))


var ifres16973 Obj

if True == tmp16977 {
tmp16974 := Call1(__e, SymFn(symtail), let__16603)


let__16604 := tmp16974
_ = let__16604

tmp16975 := Call2(__e, SymFn(symshen_4comb), let__16604, symshen_4skip)


ifres16973 = tmp16975


} else {
tmp16976 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16973 = tmp16976


}

ifres16970 = ifres16973


}

let__16601 := ifres16970
_ = let__16601

tmp16980 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16601)


if True == tmp16980 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16601
}


} else {
return let__16595
}


panic("unreachable")
})

tmp16983 := Call2(__e, ns2_1set, symshen_4_5packagenames_6, tmp16953)


_ = tmp16983

tmp16984 := MakeNative1(func(__e *ControlFlow, V219 Obj) Obj {
tmp16985 := Call1(__e, SymFn(symshen_4_5packagechar_6), V219)


let__16606 := tmp16985
_ = let__16606

tmp16995 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16606)


var ifres16986 Obj

if True == tmp16995 {
tmp16987 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16986 = tmp16987


} else {
tmp16988 := Call1(__e, SymFn(symshen_4in_1_6), let__16606)


let__16607 := tmp16988
_ = let__16607

tmp16989 := Call1(__e, SymFn(symshen_4_5packagename_6), let__16607)


let__16608 := tmp16989
_ = let__16608

tmp16994 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16608)


var ifres16990 Obj

if True == tmp16994 {
tmp16991 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16990 = tmp16991


} else {
tmp16992 := Call1(__e, SymFn(symshen_4in_1_6), let__16608)


let__16609 := tmp16992
_ = let__16609

tmp16993 := Call2(__e, SymFn(symshen_4comb), let__16609, symshen_4skip)


ifres16990 = tmp16993


}

ifres16986 = ifres16990


}

let__16605 := ifres16986
_ = let__16605

tmp17005 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16605)


if True == tmp17005 {
tmp16996 := Call1(__e, SymFn(sym_5e_6), V219)


let__16611 := tmp16996
_ = let__16611

tmp17001 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16611)


var ifres16997 Obj

if True == tmp17001 {
tmp16998 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16997 = tmp16998


} else {
tmp16999 := Call1(__e, SymFn(symshen_4in_1_6), let__16611)


let__16612 := tmp16999
_ = let__16612

tmp17000 := Call2(__e, SymFn(symshen_4comb), let__16612, symshen_4skip)


ifres16997 = tmp17000


}

let__16610 := ifres16997
_ = let__16610

tmp17003 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16610)


if True == tmp17003 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16610
}


} else {
return let__16605
}


panic("unreachable")
})

tmp17006 := Call2(__e, ns2_1set, symshen_4_5packagename_6, tmp16984)


_ = tmp17006

tmp17007 := MakeNative1(func(__e *ControlFlow, V228 Obj) Obj {
tmp17017 := GConsP(__e, V228)

var ifres17008 Obj

if True == tmp17017 {
tmp17009 := Call1(__e, SymFn(symhead), V228)


let__16614 := tmp17009
_ = let__16614

tmp17010 := Call1(__e, SymFn(symtail), V228)


let__16615 := tmp17010
_ = let__16615

tmp17014 := GEq(__e, let__16614, MakeString("."))

tmp17015 := GNot(__e, tmp17014)

var ifres17011 Obj

if True == tmp17015 {
tmp17012 := Call2(__e, SymFn(symshen_4comb), let__16615, symshen_4skip)


ifres17011 = tmp17012


} else {
tmp17013 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17011 = tmp17013


}

ifres17008 = ifres17011


} else {
tmp17016 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17008 = tmp17016


}

let__16613 := ifres17008
_ = let__16613

tmp17019 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16613)


if True == tmp17019 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16613
}


panic("unreachable")
})

tmp17020 := Call2(__e, ns2_1set, symshen_4_5packagechar_6, tmp17007)


_ = tmp17020

tmp17021 := MakeNative1(func(__e *ControlFlow, V232 Obj) Obj {
tmp17039 := Call2(__e, SymFn(symshen_4hds_a_2), V232, MakeString("<"))


var ifres17022 Obj

if True == tmp17039 {
tmp17023 := Call1(__e, SymFn(symtail), V232)


let__16617 := tmp17023
_ = let__16617

tmp17024 := Call1(__e, SymFn(sym_5_b_6), let__16617)


let__16618 := tmp17024
_ = let__16618

tmp17037 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16618)


var ifres17025 Obj

if True == tmp17037 {
tmp17026 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17025 = tmp17026


} else {
tmp17027 := Call1(__e, SymFn(symshen_4_5_1out), let__16618)


let__16619 := tmp17027
_ = let__16619

tmp17028 := Call1(__e, SymFn(symshen_4in_1_6), let__16618)


let__16620 := tmp17028
_ = let__16620

tmp17032 := Call1(__e, SymFn(symreverse), let__16619)


let__16621 := tmp17032
_ = let__16621

tmp17036 := GConsP(__e, let__16621)

var ifres17033 Obj

if True == tmp17036 {
tmp17034 := GHd(__e, let__16621)

tmp17035 := GEq(__e, tmp17034, MakeString(">"))

ifres17033 = tmp17035


} else {
ifres17033 = False


}

var ifres17029 Obj

if True == ifres17033 {
tmp17030 := Call2(__e, SymFn(symshen_4comb), let__16620, symshen_4skip)


ifres17029 = tmp17030


} else {
tmp17031 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17029 = tmp17031


}

ifres17025 = ifres17029


}

ifres17022 = ifres17025


} else {
tmp17038 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17022 = tmp17038


}

let__16616 := ifres17022
_ = let__16616

tmp17041 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16616)


if True == tmp17041 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16616
}


panic("unreachable")
})

tmp17042 := Call2(__e, ns2_1set, symshen_4_5non_1terminal_1name_6, tmp17021)


_ = tmp17042

tmp17043 := MakeNative1(func(__e *ControlFlow, V239 Obj) Obj {
tmp17044 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

return GEq(__e, V239, tmp17044)


panic("unreachable")
})

tmp17045 := Call2(__e, ns2_1set, symshen_4semicolon_2, tmp17043)


_ = tmp17045

tmp17046 := MakeNative1(func(__e *ControlFlow, V240 Obj) Obj {
tmp17055 := GConsP(__e, V240)

var ifres17047 Obj

if True == tmp17055 {
tmp17048 := Call1(__e, SymFn(symhead), V240)


let__16623 := tmp17048
_ = let__16623

tmp17049 := Call1(__e, SymFn(symtail), V240)


let__16624 := tmp17049
_ = let__16624

tmp17053 := Call1(__e, SymFn(symshen_4colon_1equal_2), let__16623)


var ifres17050 Obj

if True == tmp17053 {
tmp17051 := Call2(__e, SymFn(symshen_4comb), let__16624, symshen_4skip)


ifres17050 = tmp17051


} else {
tmp17052 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17050 = tmp17052


}

ifres17047 = ifres17050


} else {
tmp17054 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17047 = tmp17054


}

let__16622 := ifres17047
_ = let__16622

tmp17057 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16622)


if True == tmp17057 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16622
}


panic("unreachable")
})

tmp17058 := Call2(__e, ns2_1set, symshen_4_5colon_1equal_6, tmp17046)


_ = tmp17058

tmp17059 := MakeNative1(func(__e *ControlFlow, V244 Obj) Obj {
tmp17060 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":="))

return GEq(__e, tmp17060, V244)


panic("unreachable")
})

tmp17061 := Call2(__e, ns2_1set, symshen_4colon_1equal_2, tmp17059)


_ = tmp17061

tmp17062 := MakeNative1(func(__e *ControlFlow, V245 Obj) Obj {
tmp17063 := Call1(__e, SymFn(symshen_4_5syntax_1item_6), V245)


let__16626 := tmp17063
_ = let__16626

tmp17076 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16626)


var ifres17064 Obj

if True == tmp17076 {
tmp17065 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17064 = tmp17065


} else {
tmp17066 := Call1(__e, SymFn(symshen_4_5_1out), let__16626)


let__16627 := tmp17066
_ = let__16627

tmp17067 := Call1(__e, SymFn(symshen_4in_1_6), let__16626)


let__16628 := tmp17067
_ = let__16628

tmp17068 := Call1(__e, SymFn(symshen_4_5syntax_6), let__16628)


let__16629 := tmp17068
_ = let__16629

tmp17075 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16629)


var ifres17069 Obj

if True == tmp17075 {
tmp17070 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17069 = tmp17070


} else {
tmp17071 := Call1(__e, SymFn(symshen_4_5_1out), let__16629)


let__16630 := tmp17071
_ = let__16630

tmp17072 := Call1(__e, SymFn(symshen_4in_1_6), let__16629)


let__16631 := tmp17072
_ = let__16631

tmp17073 := GCons(__e, let__16627, let__16630)

tmp17074 := Call2(__e, SymFn(symshen_4comb), let__16631, tmp17073)


ifres17069 = tmp17074


}

ifres17064 = ifres17069


}

let__16625 := ifres17064
_ = let__16625

tmp17088 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16625)


if True == tmp17088 {
tmp17077 := Call1(__e, SymFn(symshen_4_5syntax_1item_6), V245)


let__16633 := tmp17077
_ = let__16633

tmp17084 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16633)


var ifres17078 Obj

if True == tmp17084 {
tmp17079 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17078 = tmp17079


} else {
tmp17080 := Call1(__e, SymFn(symshen_4_5_1out), let__16633)


let__16634 := tmp17080
_ = let__16634

tmp17081 := Call1(__e, SymFn(symshen_4in_1_6), let__16633)


let__16635 := tmp17081
_ = let__16635

tmp17082 := GCons(__e, let__16634, Nil)

tmp17083 := Call2(__e, SymFn(symshen_4comb), let__16635, tmp17082)


ifres17078 = tmp17083


}

let__16632 := ifres17078
_ = let__16632

tmp17086 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16632)


if True == tmp17086 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16632
}


} else {
return let__16625
}


panic("unreachable")
})

tmp17089 := Call2(__e, ns2_1set, symshen_4_5syntax_6, tmp17062)


_ = tmp17089

tmp17090 := MakeNative1(func(__e *ControlFlow, V257 Obj) Obj {
tmp17099 := GConsP(__e, V257)

var ifres17091 Obj

if True == tmp17099 {
tmp17092 := Call1(__e, SymFn(symhead), V257)


let__16637 := tmp17092
_ = let__16637

tmp17093 := Call1(__e, SymFn(symtail), V257)


let__16638 := tmp17093
_ = let__16638

tmp17097 := Call1(__e, SymFn(symshen_4syntax_1item_2), let__16637)


var ifres17094 Obj

if True == tmp17097 {
tmp17095 := Call2(__e, SymFn(symshen_4comb), let__16638, let__16637)


ifres17094 = tmp17095


} else {
tmp17096 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17094 = tmp17096


}

ifres17091 = ifres17094


} else {
tmp17098 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17091 = tmp17098


}

let__16636 := ifres17091
_ = let__16636

tmp17101 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16636)


if True == tmp17101 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16636
}


panic("unreachable")
})

tmp17102 := Call2(__e, ns2_1set, symshen_4_5syntax_1item_6, tmp17090)


_ = tmp17102

tmp17103 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V263 := __self1
_ = V263
tmp17133 := Call1(__e, SymFn(symshen_4colon_1equal_2), V263)


if True == tmp17133 {
return False
} else {
tmp17131 := Call1(__e, SymFn(symshen_4semicolon_2), V263)


if True == tmp17131 {
return False
} else {
tmp17129 := Call1(__e, SymFn(symatom_2), V263)


if True == tmp17129 {
return True
} else {
tmp17127 := GConsP(__e, V263)

var ifres17112 Obj

if True == tmp17127 {
tmp17125 := GHd(__e, V263)

tmp17126 := GEq(__e, symcons, tmp17125)

var ifres17113 Obj

if True == tmp17126 {
tmp17123 := GTl(__e, V263)

tmp17124 := GConsP(__e, tmp17123)

var ifres17114 Obj

if True == tmp17124 {
tmp17120 := GTl(__e, V263)

tmp17121 := GTl(__e, tmp17120)

tmp17122 := GConsP(__e, tmp17121)

var ifres17115 Obj

if True == tmp17122 {
tmp17116 := GTl(__e, V263)

tmp17117 := GTl(__e, tmp17116)

tmp17118 := GTl(__e, tmp17117)

tmp17119 := GEq(__e, Nil, tmp17118)

ifres17115 = tmp17119


} else {
ifres17115 = False


}

ifres17114 = ifres17115


} else {
ifres17114 = False


}

ifres17113 = ifres17114


} else {
ifres17113 = False


}

ifres17112 = ifres17113


} else {
ifres17112 = False


}

if True == ifres17112 {
tmp17108 := GTl(__e, V263)

tmp17109 := GHd(__e, tmp17108)

tmp17110 := Call1(__e, SymFn(symshen_4syntax_1item_2), tmp17109)


if True == tmp17110 {
tmp17104 := GTl(__e, V263)

tmp17105 := GTl(__e, tmp17104)

tmp17106 := GHd(__e, tmp17105)

if SymFn(symshen_4syntax_1item_2) == __self {
__self1 = tmp17106
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4syntax_1item_2), tmp17106)


} else {
return False
}


} else {
return False
}


}


}


}


panic("unreachable")
})
return __self
})()

tmp17134 := Call2(__e, ns2_1set, symshen_4syntax_1item_2, tmp17103)


_ = tmp17134

tmp17135 := MakeNative1(func(__e *ControlFlow, V264 Obj) Obj {
tmp17136 := Call1(__e, SymFn(symshen_4_5colon_1equal_6), V264)


let__16640 := tmp17136
_ = let__16640

tmp17162 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16640)


var ifres17137 Obj

if True == tmp17162 {
tmp17138 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17137 = tmp17138


} else {
tmp17139 := Call1(__e, SymFn(symshen_4in_1_6), let__16640)


let__16641 := tmp17139
_ = let__16641

tmp17161 := GConsP(__e, let__16641)

var ifres17140 Obj

if True == tmp17161 {
tmp17141 := Call1(__e, SymFn(symhead), let__16641)


let__16642 := tmp17141
_ = let__16642

tmp17142 := Call1(__e, SymFn(symtail), let__16641)


let__16643 := tmp17142
_ = let__16643

tmp17159 := Call2(__e, SymFn(symshen_4hds_a_2), let__16643, symwhere)


var ifres17143 Obj

if True == tmp17159 {
tmp17144 := Call1(__e, SymFn(symtail), let__16643)


let__16644 := tmp17144
_ = let__16644

tmp17157 := GConsP(__e, let__16644)

var ifres17145 Obj

if True == tmp17157 {
tmp17146 := Call1(__e, SymFn(symhead), let__16644)


let__16645 := tmp17146
_ = let__16645

tmp17147 := Call1(__e, SymFn(symtail), let__16644)


let__16646 := tmp17147
_ = let__16646

tmp17154 := Call1(__e, SymFn(symshen_4semicolon_2), let__16642)


tmp17155 := GNot(__e, tmp17154)

var ifres17148 Obj

if True == tmp17155 {
tmp17149 := GCons(__e, let__16642, Nil)

tmp17150 := GCons(__e, let__16645, tmp17149)

tmp17151 := GCons(__e, symwhere, tmp17150)

tmp17152 := Call2(__e, SymFn(symshen_4comb), let__16646, tmp17151)


ifres17148 = tmp17152


} else {
tmp17153 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17148 = tmp17153


}

ifres17145 = ifres17148


} else {
tmp17156 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17145 = tmp17156


}

ifres17143 = ifres17145


} else {
tmp17158 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17143 = tmp17158


}

ifres17140 = ifres17143


} else {
tmp17160 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17140 = tmp17160


}

ifres17137 = ifres17140


}

let__16639 := ifres17137
_ = let__16639

tmp17181 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16639)


if True == tmp17181 {
tmp17163 := Call1(__e, SymFn(symshen_4_5colon_1equal_6), V264)


let__16648 := tmp17163
_ = let__16648

tmp17177 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16648)


var ifres17164 Obj

if True == tmp17177 {
tmp17165 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17164 = tmp17165


} else {
tmp17166 := Call1(__e, SymFn(symshen_4in_1_6), let__16648)


let__16649 := tmp17166
_ = let__16649

tmp17176 := GConsP(__e, let__16649)

var ifres17167 Obj

if True == tmp17176 {
tmp17168 := Call1(__e, SymFn(symhead), let__16649)


let__16650 := tmp17168
_ = let__16650

tmp17169 := Call1(__e, SymFn(symtail), let__16649)


let__16651 := tmp17169
_ = let__16651

tmp17173 := Call1(__e, SymFn(symshen_4semicolon_2), let__16650)


tmp17174 := GNot(__e, tmp17173)

var ifres17170 Obj

if True == tmp17174 {
tmp17171 := Call2(__e, SymFn(symshen_4comb), let__16651, let__16650)


ifres17170 = tmp17171


} else {
tmp17172 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17170 = tmp17172


}

ifres17167 = ifres17170


} else {
tmp17175 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17167 = tmp17175


}

ifres17164 = ifres17167


}

let__16647 := ifres17164
_ = let__16647

tmp17179 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16647)


if True == tmp17179 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16647
}


} else {
return let__16639
}


panic("unreachable")
})

tmp17182 := Call2(__e, ns2_1set, symshen_4_5semantics_6, tmp17135)


_ = tmp17182

tmp17183 := MakeNative3(func(__e *ControlFlow, V286, V287, V288 Obj) Obj {
tmp17191 := GEq(__e, Nil, V288)

if True == tmp17191 {
return GCons(__e, symshen_4parse_1failure, Nil)
} else {
tmp17189 := GConsP(__e, V288)

if True == tmp17189 {
tmp17184 := GHd(__e, V288)

tmp17185 := Call3(__e, SymFn(symshen_4c_1rule_1_6shen), V286, tmp17184, V287)


tmp17186 := GTl(__e, V288)

tmp17187 := Call3(__e, SymFn(symshen_4c_1rules_1_6shen), V286, V287, tmp17186)


return TailCall2(__e, SymFn(symshen_4combine_1c_1code), tmp17185, tmp17187)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.c-rules->shen\n"))
}


}


panic("unreachable")
})

tmp17192 := Call2(__e, ns2_1set, symshen_4c_1rules_1_6shen, tmp17183)


_ = tmp17192

tmp17193 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall0(__e, SymFn(symfail))
panic("unreachable")
})

tmp17194 := Call2(__e, ns2_1set, symshen_4parse_1failure, tmp17193)


_ = tmp17194

tmp17195 := MakeNative2(func(__e *ControlFlow, V289, V290 Obj) Obj {
tmp17196 := GCons(__e, symResult, Nil)

tmp17197 := GCons(__e, symshen_4parse_1failure_2, tmp17196)

tmp17198 := GCons(__e, symResult, Nil)

tmp17199 := GCons(__e, V290, tmp17198)

tmp17200 := GCons(__e, tmp17197, tmp17199)

tmp17201 := GCons(__e, symif, tmp17200)

tmp17202 := GCons(__e, tmp17201, Nil)

tmp17203 := GCons(__e, V289, tmp17202)

tmp17204 := GCons(__e, symResult, tmp17203)

return GCons(__e, symlet, tmp17204)


panic("unreachable")
})

tmp17205 := Call2(__e, ns2_1set, symshen_4combine_1c_1code, tmp17195)


_ = tmp17205

tmp17206 := MakeNative3(func(__e *ControlFlow, V297, V298, V299 Obj) Obj {
tmp17218 := GConsP(__e, V298)

var ifres17211 Obj

if True == tmp17218 {
tmp17216 := GTl(__e, V298)

tmp17217 := GConsP(__e, tmp17216)

var ifres17212 Obj

if True == tmp17217 {
tmp17213 := GTl(__e, V298)

tmp17214 := GTl(__e, tmp17213)

tmp17215 := GEq(__e, Nil, tmp17214)

ifres17212 = tmp17215


} else {
ifres17212 = False


}

ifres17211 = ifres17212


} else {
ifres17211 = False


}

if True == ifres17211 {
tmp17207 := GHd(__e, V298)

tmp17208 := GTl(__e, V298)

tmp17209 := GHd(__e, tmp17208)

return TailCall4(__e, SymFn(symshen_4yacc_1syntax), V297, V299, tmp17207, tmp17209)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.c-rule->shen\n"))
}


panic("unreachable")
})

tmp17219 := Call2(__e, ns2_1set, symshen_4c_1rule_1_6shen, tmp17206)


_ = tmp17219

tmp17220 := MakeNative4(func(__e *ControlFlow, V308, V309, V310, V311 Obj) Obj {
tmp17279 := GEq(__e, Nil, V310)

var ifres17262 Obj

if True == tmp17279 {
tmp17278 := GConsP(__e, V311)

var ifres17263 Obj

if True == tmp17278 {
tmp17276 := GHd(__e, V311)

tmp17277 := GEq(__e, symwhere, tmp17276)

var ifres17264 Obj

if True == tmp17277 {
tmp17274 := GTl(__e, V311)

tmp17275 := GConsP(__e, tmp17274)

var ifres17265 Obj

if True == tmp17275 {
tmp17271 := GTl(__e, V311)

tmp17272 := GTl(__e, tmp17271)

tmp17273 := GConsP(__e, tmp17272)

var ifres17266 Obj

if True == tmp17273 {
tmp17267 := GTl(__e, V311)

tmp17268 := GTl(__e, tmp17267)

tmp17269 := GTl(__e, tmp17268)

tmp17270 := GEq(__e, Nil, tmp17269)

ifres17266 = tmp17270


} else {
ifres17266 = False


}

ifres17265 = ifres17266


} else {
ifres17265 = False


}

ifres17264 = ifres17265


} else {
ifres17264 = False


}

ifres17263 = ifres17264


} else {
ifres17263 = False


}

ifres17262 = ifres17263


} else {
ifres17262 = False


}

if True == ifres17262 {
tmp17221 := GTl(__e, V311)

tmp17222 := GHd(__e, tmp17221)

tmp17223 := Call1(__e, SymFn(symshen_4process_1yacc_1semantics), tmp17222)


tmp17224 := GTl(__e, V311)

tmp17225 := GTl(__e, tmp17224)

tmp17226 := GHd(__e, tmp17225)

tmp17227 := Call4(__e, SymFn(symshen_4yacc_1syntax), V308, V309, Nil, tmp17226)


tmp17228 := GCons(__e, symshen_4parse_1failure, Nil)

tmp17229 := GCons(__e, tmp17228, Nil)

tmp17230 := GCons(__e, tmp17227, tmp17229)

tmp17231 := GCons(__e, tmp17223, tmp17230)

return GCons(__e, symif, tmp17231)


} else {
tmp17260 := GEq(__e, Nil, V310)

if True == tmp17260 {
return TailCall3(__e, SymFn(symshen_4yacc_1semantics), V308, V309, V311)
} else {
tmp17258 := GConsP(__e, V310)

if True == tmp17258 {
tmp17255 := GHd(__e, V310)

tmp17256 := Call1(__e, SymFn(symshen_4non_1terminal_2), tmp17255)


if True == tmp17256 {
tmp17232 := GHd(__e, V310)

tmp17233 := GTl(__e, V310)

return TailCall5(__e, SymFn(symshen_4non_1terminalcode), V308, V309, tmp17232, tmp17233, V311)


} else {
tmp17252 := GHd(__e, V310)

tmp17253 := PrimGuard1(__e, symvariable_2, PrimIsVariable, tmp17252)

if True == tmp17253 {
tmp17234 := GHd(__e, V310)

tmp17235 := GTl(__e, V310)

return TailCall5(__e, SymFn(symshen_4variablecode), V308, V309, tmp17234, tmp17235, V311)


} else {
tmp17249 := GHd(__e, V310)

tmp17250 := GEq(__e, sym__, tmp17249)

if True == tmp17250 {
tmp17236 := GHd(__e, V310)

tmp17237 := GTl(__e, V310)

return TailCall5(__e, SymFn(symshen_4wildcardcode), V308, V309, tmp17236, tmp17237, V311)


} else {
tmp17246 := GHd(__e, V310)

tmp17247 := Call1(__e, SymFn(symatom_2), tmp17246)


if True == tmp17247 {
tmp17238 := GHd(__e, V310)

tmp17239 := GTl(__e, V310)

return TailCall5(__e, SymFn(symshen_4terminalcode), V308, V309, tmp17238, tmp17239, V311)


} else {
tmp17243 := GHd(__e, V310)

tmp17244 := GConsP(__e, tmp17243)

if True == tmp17244 {
tmp17240 := GHd(__e, V310)

tmp17241 := GTl(__e, V310)

return TailCall5(__e, SymFn(symshen_4conscode), V308, V309, tmp17240, tmp17241, V311)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.yacc-syntax\n"))
}


}


}


}


}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.yacc-syntax\n"))
}


}


}


panic("unreachable")
})

tmp17280 := Call2(__e, ns2_1set, symshen_4yacc_1syntax, tmp17220)


_ = tmp17280

tmp17281 := MakeNative5(func(__e *ControlFlow, V312, V313, V314, V315, V316 Obj) Obj {
tmp17282 := Call2(__e, SymFn(symconcat), symParse, V314)


let__16652 := tmp17282
_ = let__16652

tmp17283 := Call2(__e, SymFn(symconcat), symAction, V314)


let__16653 := tmp17283
_ = let__16653

tmp17284 := Call2(__e, SymFn(symconcat), symRemainder, V314)


let__16654 := tmp17284
_ = let__16654

tmp17285 := GCons(__e, V313, Nil)

tmp17286 := GCons(__e, V314, tmp17285)

tmp17287 := GCons(__e, let__16652, Nil)

tmp17288 := GCons(__e, symshen_4parse_1failure_2, tmp17287)

tmp17289 := GCons(__e, symshen_4parse_1failure, Nil)

tmp17290 := GCons(__e, let__16652, Nil)

tmp17291 := GCons(__e, symshen_4in_1_6, tmp17290)

tmp17292 := Call4(__e, SymFn(symshen_4yacc_1syntax), V312, let__16654, V315, V316)


tmp17293 := GCons(__e, tmp17292, Nil)

tmp17294 := GCons(__e, tmp17291, tmp17293)

tmp17295 := GCons(__e, let__16654, tmp17294)

tmp17296 := GCons(__e, symlet, tmp17295)

let__16655 := tmp17296
_ = let__16655

tmp17306 := Call2(__e, SymFn(symshen_4occurs_1check_2), V314, V316)


var ifres17304 Obj

if True == tmp17306 {
ifres17304 = True


} else {
tmp17305 := Call2(__e, SymFn(symshen_4occurs_1check_2), let__16653, V316)


ifres17304 = tmp17305


}

var ifres17297 Obj

if True == ifres17304 {
tmp17298 := GCons(__e, let__16652, Nil)

tmp17299 := GCons(__e, symshen_4_5_1out, tmp17298)

tmp17300 := GCons(__e, let__16655, Nil)

tmp17301 := GCons(__e, tmp17299, tmp17300)

tmp17302 := GCons(__e, let__16653, tmp17301)

tmp17303 := GCons(__e, symlet, tmp17302)

ifres17297 = tmp17303


} else {
ifres17297 = let__16655


}

tmp17307 := GCons(__e, ifres17297, Nil)

tmp17308 := GCons(__e, tmp17289, tmp17307)

tmp17309 := GCons(__e, tmp17288, tmp17308)

tmp17310 := GCons(__e, symif, tmp17309)

tmp17311 := GCons(__e, tmp17310, Nil)

tmp17312 := GCons(__e, tmp17286, tmp17311)

tmp17313 := GCons(__e, let__16652, tmp17312)

return GCons(__e, symlet, tmp17313)


panic("unreachable")
})

tmp17314 := Call2(__e, ns2_1set, symshen_4non_1terminalcode, tmp17281)


_ = tmp17314

tmp17315 := MakeNative5(func(__e *ControlFlow, V321, V322, V323, V324, V325 Obj) Obj {
tmp17316 := Call1(__e, SymFn(symgensym), symRemainder)


let__16656 := tmp17316
_ = let__16656

tmp17317 := GCons(__e, V322, Nil)

tmp17318 := GCons(__e, symcons_2, tmp17317)

tmp17319 := GCons(__e, V322, Nil)

tmp17320 := GCons(__e, symtail, tmp17319)

tmp17321 := Call4(__e, SymFn(symshen_4yacc_1syntax), V321, let__16656, V324, V325)


tmp17322 := GCons(__e, tmp17321, Nil)

tmp17323 := GCons(__e, tmp17320, tmp17322)

tmp17324 := GCons(__e, let__16656, tmp17323)

tmp17325 := GCons(__e, symlet, tmp17324)

let__16657 := tmp17325
_ = let__16657

tmp17333 := Call2(__e, SymFn(symshen_4occurs_1check_2), V323, V325)


var ifres17326 Obj

if True == tmp17333 {
tmp17327 := GCons(__e, V322, Nil)

tmp17328 := GCons(__e, symhead, tmp17327)

tmp17329 := GCons(__e, let__16657, Nil)

tmp17330 := GCons(__e, tmp17328, tmp17329)

tmp17331 := GCons(__e, V323, tmp17330)

tmp17332 := GCons(__e, symlet, tmp17331)

ifres17326 = tmp17332


} else {
ifres17326 = let__16657


}

tmp17334 := GCons(__e, symshen_4parse_1failure, Nil)

tmp17335 := GCons(__e, tmp17334, Nil)

tmp17336 := GCons(__e, ifres17326, tmp17335)

tmp17337 := GCons(__e, tmp17318, tmp17336)

return GCons(__e, symif, tmp17337)


panic("unreachable")
})

tmp17338 := Call2(__e, ns2_1set, symshen_4variablecode, tmp17315)


_ = tmp17338

tmp17339 := MakeNative5(func(__e *ControlFlow, V328, V329, V330, V331, V332 Obj) Obj {
tmp17340 := Call1(__e, SymFn(symgensym), symRemainder)


let__16658 := tmp17340
_ = let__16658

tmp17341 := GCons(__e, V329, Nil)

tmp17342 := GCons(__e, symcons_2, tmp17341)

tmp17343 := GCons(__e, V329, Nil)

tmp17344 := GCons(__e, symtail, tmp17343)

tmp17345 := Call4(__e, SymFn(symshen_4yacc_1syntax), V328, let__16658, V331, V332)


tmp17346 := GCons(__e, tmp17345, Nil)

tmp17347 := GCons(__e, tmp17344, tmp17346)

tmp17348 := GCons(__e, let__16658, tmp17347)

tmp17349 := GCons(__e, symlet, tmp17348)

tmp17350 := GCons(__e, symshen_4parse_1failure, Nil)

tmp17351 := GCons(__e, tmp17350, Nil)

tmp17352 := GCons(__e, tmp17349, tmp17351)

tmp17353 := GCons(__e, tmp17342, tmp17352)

return GCons(__e, symif, tmp17353)


panic("unreachable")
})

tmp17354 := Call2(__e, ns2_1set, symshen_4wildcardcode, tmp17339)


_ = tmp17354

tmp17355 := MakeNative5(func(__e *ControlFlow, V334, V335, V336, V337, V338 Obj) Obj {
tmp17356 := Call1(__e, SymFn(symgensym), symRemainder)


let__16659 := tmp17356
_ = let__16659

tmp17357 := GCons(__e, V336, Nil)

tmp17358 := GCons(__e, V335, tmp17357)

tmp17359 := GCons(__e, symshen_4hds_a_2, tmp17358)

tmp17360 := GCons(__e, V335, Nil)

tmp17361 := GCons(__e, symtail, tmp17360)

tmp17362 := Call4(__e, SymFn(symshen_4yacc_1syntax), V334, let__16659, V337, V338)


tmp17363 := GCons(__e, tmp17362, Nil)

tmp17364 := GCons(__e, tmp17361, tmp17363)

tmp17365 := GCons(__e, let__16659, tmp17364)

tmp17366 := GCons(__e, symlet, tmp17365)

tmp17367 := GCons(__e, symshen_4parse_1failure, Nil)

tmp17368 := GCons(__e, tmp17367, Nil)

tmp17369 := GCons(__e, tmp17366, tmp17368)

tmp17370 := GCons(__e, tmp17359, tmp17369)

return GCons(__e, symif, tmp17370)


panic("unreachable")
})

tmp17371 := Call2(__e, ns2_1set, symshen_4terminalcode, tmp17355)


_ = tmp17371

tmp17372 := MakeNative2(func(__e *ControlFlow, V347, V348 Obj) Obj {
tmp17377 := GConsP(__e, V347)

var ifres17374 Obj

if True == tmp17377 {
tmp17375 := GHd(__e, V347)

tmp17376 := GEq(__e, tmp17375, V348)

ifres17374 = tmp17376


} else {
ifres17374 = False


}

if True == ifres17374 {
return True
} else {
return False
}


panic("unreachable")
})

tmp17378 := Call2(__e, ns2_1set, symshen_4hds_a_2, tmp17372)


_ = tmp17378

tmp17379 := MakeNative5(func(__e *ControlFlow, V349, V350, V351, V352, V353 Obj) Obj {
tmp17380 := Call1(__e, SymFn(symgensym), symRemainder)


let__16660 := tmp17380
_ = let__16660

tmp17381 := Call1(__e, SymFn(symgensym), symHd)


let__16661 := tmp17381
_ = let__16661

tmp17382 := Call1(__e, SymFn(symgensym), symTl)


let__16662 := tmp17382
_ = let__16662

tmp17383 := GCons(__e, V350, Nil)

tmp17384 := GCons(__e, symshen_4ccons_2, tmp17383)

tmp17385 := GCons(__e, V350, Nil)

tmp17386 := GCons(__e, symhead, tmp17385)

tmp17387 := GCons(__e, V350, Nil)

tmp17388 := GCons(__e, symtail, tmp17387)

tmp17389 := Call1(__e, SymFn(symshen_4decons), V351)


tmp17390 := GCons(__e, sym_5end_6, Nil)

tmp17391 := Call2(__e, SymFn(symappend), tmp17389, tmp17390)


tmp17392 := Call4(__e, SymFn(symshen_4yacc_1syntax), V349, let__16662, V352, V353)


tmp17393 := GCons(__e, tmp17392, Nil)

tmp17394 := GCons(__e, symshen_4processed, tmp17393)

tmp17395 := Call4(__e, SymFn(symshen_4yacc_1syntax), V349, let__16661, tmp17391, tmp17394)


tmp17396 := GCons(__e, tmp17395, Nil)

tmp17397 := GCons(__e, tmp17388, tmp17396)

tmp17398 := GCons(__e, let__16662, tmp17397)

tmp17399 := GCons(__e, tmp17386, tmp17398)

tmp17400 := GCons(__e, let__16661, tmp17399)

tmp17401 := GCons(__e, symlet, tmp17400)

tmp17402 := GCons(__e, symshen_4parse_1failure, Nil)

tmp17403 := GCons(__e, tmp17402, Nil)

tmp17404 := GCons(__e, tmp17401, tmp17403)

tmp17405 := GCons(__e, tmp17384, tmp17404)

return GCons(__e, symif, tmp17405)


panic("unreachable")
})

tmp17406 := Call2(__e, ns2_1set, symshen_4conscode, tmp17379)


_ = tmp17406

tmp17407 := MakeNative1(func(__e *ControlFlow, V367 Obj) Obj {
tmp17417 := GConsP(__e, V367)

var ifres17414 Obj

if True == tmp17417 {
tmp17415 := GHd(__e, V367)

tmp17416 := GConsP(__e, tmp17415)

ifres17414 = tmp17416


} else {
ifres17414 = False


}

if True == ifres17414 {
return True
} else {
tmp17412 := GConsP(__e, V367)

var ifres17409 Obj

if True == tmp17412 {
tmp17410 := GHd(__e, V367)

tmp17411 := GEq(__e, Nil, tmp17410)

ifres17409 = tmp17411


} else {
ifres17409 = False


}

if True == ifres17409 {
return True
} else {
return False
}


}


panic("unreachable")
})

tmp17418 := Call2(__e, ns2_1set, symshen_4ccons_2, tmp17407)


_ = tmp17418

tmp17419 := MakeNative1(func(__e *ControlFlow, V368 Obj) Obj {
tmp17442 := GConsP(__e, V368)

var ifres17427 Obj

if True == tmp17442 {
tmp17440 := GHd(__e, V368)

tmp17441 := GEq(__e, symcons, tmp17440)

var ifres17428 Obj

if True == tmp17441 {
tmp17438 := GTl(__e, V368)

tmp17439 := GConsP(__e, tmp17438)

var ifres17429 Obj

if True == tmp17439 {
tmp17435 := GTl(__e, V368)

tmp17436 := GTl(__e, tmp17435)

tmp17437 := GConsP(__e, tmp17436)

var ifres17430 Obj

if True == tmp17437 {
tmp17431 := GTl(__e, V368)

tmp17432 := GTl(__e, tmp17431)

tmp17433 := GTl(__e, tmp17432)

tmp17434 := GEq(__e, Nil, tmp17433)

ifres17430 = tmp17434


} else {
ifres17430 = False


}

ifres17429 = ifres17430


} else {
ifres17429 = False


}

ifres17428 = ifres17429


} else {
ifres17428 = False


}

ifres17427 = ifres17428


} else {
ifres17427 = False


}

if True == ifres17427 {
tmp17420 := GTl(__e, V368)

tmp17421 := GHd(__e, tmp17420)

tmp17422 := GTl(__e, V368)

tmp17423 := GTl(__e, tmp17422)

tmp17424 := GHd(__e, tmp17423)

tmp17425 := Call1(__e, SymFn(symshen_4decons), tmp17424)


return GCons(__e, tmp17421, tmp17425)


} else {
return V368
}


panic("unreachable")
})

tmp17443 := Call2(__e, ns2_1set, symshen_4decons, tmp17419)


_ = tmp17443

tmp17444 := MakeNative2(func(__e *ControlFlow, V369, V370 Obj) Obj {
tmp17445 := GCons(__e, V370, Nil)

return GCons(__e, V369, tmp17445)


panic("unreachable")
})

tmp17446 := Call2(__e, ns2_1set, symshen_4comb, tmp17444)


_ = tmp17446

tmp17447 := MakeNative3(func(__e *ControlFlow, V375, V376, V377 Obj) Obj {
tmp17464 := GConsP(__e, V377)

var ifres17454 Obj

if True == tmp17464 {
tmp17462 := GHd(__e, V377)

tmp17463 := GEq(__e, symshen_4processed, tmp17462)

var ifres17455 Obj

if True == tmp17463 {
tmp17460 := GTl(__e, V377)

tmp17461 := GConsP(__e, tmp17460)

var ifres17456 Obj

if True == tmp17461 {
tmp17457 := GTl(__e, V377)

tmp17458 := GTl(__e, tmp17457)

tmp17459 := GEq(__e, Nil, tmp17458)

ifres17456 = tmp17459


} else {
ifres17456 = False


}

ifres17455 = ifres17456


} else {
ifres17455 = False


}

ifres17454 = ifres17455


} else {
ifres17454 = False


}

if True == ifres17454 {
tmp17448 := GTl(__e, V377)

return GHd(__e, tmp17448)


} else {
tmp17449 := Call1(__e, SymFn(symshen_4process_1yacc_1semantics), V377)


let__16663 := tmp17449
_ = let__16663

tmp17450 := Call2(__e, SymFn(symshen_4use_1type_1info), V375, let__16663)


let__16664 := tmp17450
_ = let__16664

tmp17451 := GCons(__e, let__16664, Nil)

tmp17452 := GCons(__e, V376, tmp17451)

return GCons(__e, symshen_4comb, tmp17452)


}


panic("unreachable")
})

tmp17465 := Call2(__e, ns2_1set, symshen_4yacc_1semantics, tmp17447)


_ = tmp17465

tmp17466 := MakeNative2(func(__e *ControlFlow, V383, V384 Obj) Obj {
tmp17631 := GConsP(__e, V383)

var ifres17475 Obj

if True == tmp17631 {
tmp17629 := GHd(__e, V383)

tmp17630 := GEq(__e, sym_i, tmp17629)

var ifres17476 Obj

if True == tmp17630 {
tmp17627 := GTl(__e, V383)

tmp17628 := GConsP(__e, tmp17627)

var ifres17477 Obj

if True == tmp17628 {
tmp17624 := GTl(__e, V383)

tmp17625 := GHd(__e, tmp17624)

tmp17626 := GConsP(__e, tmp17625)

var ifres17478 Obj

if True == tmp17626 {
tmp17620 := GTl(__e, V383)

tmp17621 := GHd(__e, tmp17620)

tmp17622 := GHd(__e, tmp17621)

tmp17623 := GEq(__e, symlist, tmp17622)

var ifres17479 Obj

if True == tmp17623 {
tmp17616 := GTl(__e, V383)

tmp17617 := GHd(__e, tmp17616)

tmp17618 := GTl(__e, tmp17617)

tmp17619 := GConsP(__e, tmp17618)

var ifres17480 Obj

if True == tmp17619 {
tmp17611 := GTl(__e, V383)

tmp17612 := GHd(__e, tmp17611)

tmp17613 := GTl(__e, tmp17612)

tmp17614 := GTl(__e, tmp17613)

tmp17615 := GEq(__e, Nil, tmp17614)

var ifres17481 Obj

if True == tmp17615 {
tmp17608 := GTl(__e, V383)

tmp17609 := GTl(__e, tmp17608)

tmp17610 := GConsP(__e, tmp17609)

var ifres17482 Obj

if True == tmp17610 {
tmp17604 := GTl(__e, V383)

tmp17605 := GTl(__e, tmp17604)

tmp17606 := GHd(__e, tmp17605)

tmp17607 := GEq(__e, sym_1_1_6, tmp17606)

var ifres17483 Obj

if True == tmp17607 {
tmp17600 := GTl(__e, V383)

tmp17601 := GTl(__e, tmp17600)

tmp17602 := GTl(__e, tmp17601)

tmp17603 := GConsP(__e, tmp17602)

var ifres17484 Obj

if True == tmp17603 {
tmp17595 := GTl(__e, V383)

tmp17596 := GTl(__e, tmp17595)

tmp17597 := GTl(__e, tmp17596)

tmp17598 := GHd(__e, tmp17597)

tmp17599 := GConsP(__e, tmp17598)

var ifres17485 Obj

if True == tmp17599 {
tmp17589 := GTl(__e, V383)

tmp17590 := GTl(__e, tmp17589)

tmp17591 := GTl(__e, tmp17590)

tmp17592 := GHd(__e, tmp17591)

tmp17593 := GHd(__e, tmp17592)

tmp17594 := GEq(__e, symstr, tmp17593)

var ifres17486 Obj

if True == tmp17594 {
tmp17583 := GTl(__e, V383)

tmp17584 := GTl(__e, tmp17583)

tmp17585 := GTl(__e, tmp17584)

tmp17586 := GHd(__e, tmp17585)

tmp17587 := GTl(__e, tmp17586)

tmp17588 := GConsP(__e, tmp17587)

var ifres17487 Obj

if True == tmp17588 {
tmp17576 := GTl(__e, V383)

tmp17577 := GTl(__e, tmp17576)

tmp17578 := GTl(__e, tmp17577)

tmp17579 := GHd(__e, tmp17578)

tmp17580 := GTl(__e, tmp17579)

tmp17581 := GHd(__e, tmp17580)

tmp17582 := GConsP(__e, tmp17581)

var ifres17488 Obj

if True == tmp17582 {
tmp17568 := GTl(__e, V383)

tmp17569 := GTl(__e, tmp17568)

tmp17570 := GTl(__e, tmp17569)

tmp17571 := GHd(__e, tmp17570)

tmp17572 := GTl(__e, tmp17571)

tmp17573 := GHd(__e, tmp17572)

tmp17574 := GHd(__e, tmp17573)

tmp17575 := GEq(__e, symlist, tmp17574)

var ifres17489 Obj

if True == tmp17575 {
tmp17560 := GTl(__e, V383)

tmp17561 := GTl(__e, tmp17560)

tmp17562 := GTl(__e, tmp17561)

tmp17563 := GHd(__e, tmp17562)

tmp17564 := GTl(__e, tmp17563)

tmp17565 := GHd(__e, tmp17564)

tmp17566 := GTl(__e, tmp17565)

tmp17567 := GConsP(__e, tmp17566)

var ifres17490 Obj

if True == tmp17567 {
tmp17551 := GTl(__e, V383)

tmp17552 := GTl(__e, tmp17551)

tmp17553 := GTl(__e, tmp17552)

tmp17554 := GHd(__e, tmp17553)

tmp17555 := GTl(__e, tmp17554)

tmp17556 := GHd(__e, tmp17555)

tmp17557 := GTl(__e, tmp17556)

tmp17558 := GTl(__e, tmp17557)

tmp17559 := GEq(__e, Nil, tmp17558)

var ifres17491 Obj

if True == tmp17559 {
tmp17544 := GTl(__e, V383)

tmp17545 := GTl(__e, tmp17544)

tmp17546 := GTl(__e, tmp17545)

tmp17547 := GHd(__e, tmp17546)

tmp17548 := GTl(__e, tmp17547)

tmp17549 := GTl(__e, tmp17548)

tmp17550 := GConsP(__e, tmp17549)

var ifres17492 Obj

if True == tmp17550 {
tmp17536 := GTl(__e, V383)

tmp17537 := GTl(__e, tmp17536)

tmp17538 := GTl(__e, tmp17537)

tmp17539 := GHd(__e, tmp17538)

tmp17540 := GTl(__e, tmp17539)

tmp17541 := GTl(__e, tmp17540)

tmp17542 := GTl(__e, tmp17541)

tmp17543 := GEq(__e, Nil, tmp17542)

var ifres17493 Obj

if True == tmp17543 {
tmp17531 := GTl(__e, V383)

tmp17532 := GTl(__e, tmp17531)

tmp17533 := GTl(__e, tmp17532)

tmp17534 := GTl(__e, tmp17533)

tmp17535 := GConsP(__e, tmp17534)

var ifres17494 Obj

if True == tmp17535 {
tmp17525 := GTl(__e, V383)

tmp17526 := GTl(__e, tmp17525)

tmp17527 := GTl(__e, tmp17526)

tmp17528 := GTl(__e, tmp17527)

tmp17529 := GHd(__e, tmp17528)

tmp17530 := GEq(__e, sym_j, tmp17529)

var ifres17495 Obj

if True == tmp17530 {
tmp17519 := GTl(__e, V383)

tmp17520 := GTl(__e, tmp17519)

tmp17521 := GTl(__e, tmp17520)

tmp17522 := GTl(__e, tmp17521)

tmp17523 := GTl(__e, tmp17522)

tmp17524 := GEq(__e, Nil, tmp17523)

var ifres17496 Obj

if True == tmp17524 {
tmp17506 := GTl(__e, V383)

tmp17507 := GHd(__e, tmp17506)

tmp17508 := GTl(__e, tmp17507)

tmp17509 := GHd(__e, tmp17508)

tmp17510 := GTl(__e, V383)

tmp17511 := GTl(__e, tmp17510)

tmp17512 := GTl(__e, tmp17511)

tmp17513 := GHd(__e, tmp17512)

tmp17514 := GTl(__e, tmp17513)

tmp17515 := GHd(__e, tmp17514)

tmp17516 := GTl(__e, tmp17515)

tmp17517 := GHd(__e, tmp17516)

tmp17518 := GEq(__e, tmp17509, tmp17517)

var ifres17497 Obj

if True == tmp17518 {
tmp17498 := GTl(__e, V383)

tmp17499 := GTl(__e, tmp17498)

tmp17500 := GTl(__e, tmp17499)

tmp17501 := GHd(__e, tmp17500)

tmp17502 := GTl(__e, tmp17501)

tmp17503 := GTl(__e, tmp17502)

tmp17504 := GHd(__e, tmp17503)

tmp17505 := Call1(__e, SymFn(symshen_4monomorphic_2), tmp17504)


ifres17497 = tmp17505


} else {
ifres17497 = False


}

ifres17496 = ifres17497


} else {
ifres17496 = False


}

ifres17495 = ifres17496


} else {
ifres17495 = False


}

ifres17494 = ifres17495


} else {
ifres17494 = False


}

ifres17493 = ifres17494


} else {
ifres17493 = False


}

ifres17492 = ifres17493


} else {
ifres17492 = False


}

ifres17491 = ifres17492


} else {
ifres17491 = False


}

ifres17490 = ifres17491


} else {
ifres17490 = False


}

ifres17489 = ifres17490


} else {
ifres17489 = False


}

ifres17488 = ifres17489


} else {
ifres17488 = False


}

ifres17487 = ifres17488


} else {
ifres17487 = False


}

ifres17486 = ifres17487


} else {
ifres17486 = False


}

ifres17485 = ifres17486


} else {
ifres17485 = False


}

ifres17484 = ifres17485


} else {
ifres17484 = False


}

ifres17483 = ifres17484


} else {
ifres17483 = False


}

ifres17482 = ifres17483


} else {
ifres17482 = False


}

ifres17481 = ifres17482


} else {
ifres17481 = False


}

ifres17480 = ifres17481


} else {
ifres17480 = False


}

ifres17479 = ifres17480


} else {
ifres17479 = False


}

ifres17478 = ifres17479


} else {
ifres17478 = False


}

ifres17477 = ifres17478


} else {
ifres17477 = False


}

ifres17476 = ifres17477


} else {
ifres17476 = False


}

ifres17475 = ifres17476


} else {
ifres17475 = False


}

if True == ifres17475 {
tmp17467 := GTl(__e, V383)

tmp17468 := GTl(__e, tmp17467)

tmp17469 := GTl(__e, tmp17468)

tmp17470 := GHd(__e, tmp17469)

tmp17471 := GTl(__e, tmp17470)

tmp17472 := GTl(__e, tmp17471)

tmp17473 := GCons(__e, V384, tmp17472)

return GCons(__e, symtype, tmp17473)


} else {
return V384
}


panic("unreachable")
})

tmp17632 := Call2(__e, ns2_1set, symshen_4use_1type_1info, tmp17466)


_ = tmp17632

tmp17633 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V387 := __self1
_ = V387
tmp17641 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V387)

if True == tmp17641 {
return False
} else {
tmp17639 := GConsP(__e, V387)

if True == tmp17639 {
tmp17636 := GHd(__e, V387)

tmp17637 := Call1(__e, SymFn(symshen_4monomorphic_2), tmp17636)


if True == tmp17637 {
tmp17634 := GTl(__e, V387)

if SymFn(symshen_4monomorphic_2) == __self {
__self1 = tmp17634
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4monomorphic_2), tmp17634)


} else {
return False
}


} else {
return True
}


}


panic("unreachable")
})
return __self
})()

tmp17642 := Call2(__e, ns2_1set, symshen_4monomorphic_2, tmp17633)


_ = tmp17642

tmp17643 := MakeNative1(func(__e *ControlFlow, V388 Obj) Obj {
tmp17665 := GConsP(__e, V388)

var ifres17651 Obj

if True == tmp17665 {
tmp17663 := GHd(__e, V388)

tmp17664 := GEq(__e, symprotect, tmp17663)

var ifres17652 Obj

if True == tmp17664 {
tmp17661 := GTl(__e, V388)

tmp17662 := GConsP(__e, tmp17661)

var ifres17653 Obj

if True == tmp17662 {
tmp17658 := GTl(__e, V388)

tmp17659 := GTl(__e, tmp17658)

tmp17660 := GEq(__e, Nil, tmp17659)

var ifres17654 Obj

if True == tmp17660 {
tmp17655 := GTl(__e, V388)

tmp17656 := GHd(__e, tmp17655)

tmp17657 := Call1(__e, SymFn(symshen_4non_1terminal_2), tmp17656)


ifres17654 = tmp17657


} else {
ifres17654 = False


}

ifres17653 = ifres17654


} else {
ifres17653 = False


}

ifres17652 = ifres17653


} else {
ifres17652 = False


}

ifres17651 = ifres17652


} else {
ifres17651 = False


}

if True == ifres17651 {
tmp17644 := GTl(__e, V388)

return GHd(__e, tmp17644)


} else {
tmp17649 := GConsP(__e, V388)

if True == tmp17649 {
tmp17645 := MakeNative1(func(__e *ControlFlow, Z389 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4process_1yacc_1semantics), Z389)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp17645, V388)


} else {
tmp17647 := Call1(__e, SymFn(symshen_4non_1terminal_2), V388)


if True == tmp17647 {
return TailCall2(__e, SymFn(symconcat), symAction, V388)
} else {
return V388
}


}


}


panic("unreachable")
})

tmp17666 := Call2(__e, ns2_1set, symshen_4process_1yacc_1semantics, tmp17643)


_ = tmp17666

tmp17667 := MakeNative1(func(__e *ControlFlow, V392 Obj) Obj {
tmp17678 := GConsP(__e, V392)

var ifres17671 Obj

if True == tmp17678 {
tmp17676 := GTl(__e, V392)

tmp17677 := GConsP(__e, tmp17676)

var ifres17672 Obj

if True == tmp17677 {
tmp17673 := GTl(__e, V392)

tmp17674 := GTl(__e, tmp17673)

tmp17675 := GEq(__e, Nil, tmp17674)

ifres17672 = tmp17675


} else {
ifres17672 = False


}

ifres17671 = ifres17672


} else {
ifres17671 = False


}

if True == ifres17671 {
tmp17668 := GTl(__e, V392)

return GHd(__e, tmp17668)


} else {
tmp17669 := GTl(__e, V392)

return GHd(__e, tmp17669)


}


panic("unreachable")
})

tmp17679 := Call2(__e, ns2_1set, symshen_4_5_1out, tmp17667)


_ = tmp17679

tmp17680 := MakeNative1(func(__e *ControlFlow, V393 Obj) Obj {
return GHd(__e, V393)
panic("unreachable")
})

tmp17681 := Call2(__e, ns2_1set, symshen_4in_1_6, tmp17680)


_ = tmp17681

tmp17682 := MakeNative1(func(__e *ControlFlow, V394 Obj) Obj {
tmp17683 := GCons(__e, V394, Nil)

return GCons(__e, Nil, tmp17683)


panic("unreachable")
})

tmp17684 := Call2(__e, ns2_1set, sym_5_b_6, tmp17682)


_ = tmp17684

tmp17685 := MakeNative1(func(__e *ControlFlow, V395 Obj) Obj {
tmp17686 := GCons(__e, Nil, Nil)

return GCons(__e, V395, tmp17686)


panic("unreachable")
})

tmp17687 := Call2(__e, ns2_1set, sym_5e_6, tmp17685)


_ = tmp17687

tmp17688 := MakeNative1(func(__e *ControlFlow, V398 Obj) Obj {
tmp17691 := GEq(__e, Nil, V398)

if True == tmp17691 {
tmp17689 := GCons(__e, Nil, Nil)

return GCons(__e, Nil, tmp17689)


} else {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
}


panic("unreachable")
})

return TailCall2(__e, ns2_1set, sym_5end_6, tmp17688)




panic("unreachable")
})

