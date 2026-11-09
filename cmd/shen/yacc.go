package main

import . "github.com/pyrex41/shen-go/kl"

var YaccMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp16685 := MakeNative2(func(__e *ControlFlow, V112, V113 Obj) Obj {
tmp16686 := Call1(__e, V112, V113)


let__16547 := tmp16686
_ = let__16547

tmp16693 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16547)


if True == tmp16693 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("parse failure\n"))
} else {
tmp16691 := Call1(__e, SymFn(symshen_4partial_1parse_1failure_2), let__16547)


if True == tmp16691 {
tmp16687 := Call1(__e, SymFn(symshen_4in_1_6), let__16547)


tmp16688 := PrimGuard2(__e, symset, PrimSet, symshen_4_dresidue_d, tmp16687)

_ = tmp16688

tmp16689 := GValue(__e, symshen_4_dresidue_d)

return TailCall1(__e, SymFn(symshen_4raise_1syntax_1error), tmp16689)


} else {
return TailCall1(__e, SymFn(symshen_4_5_1out), let__16547)
}


}


panic("unreachable")
})

tmp16694 := Call2(__e, ns2_1set, symcompile, tmp16685)


_ = tmp16694

tmp16695 := MakeNative1(func(__e *ControlFlow, V115 Obj) Obj {
tmp16696 := GValue(__e, sym_dmaximum_1print_1sequence_1size_d)

tmp16697 := Call3(__e, SymFn(symshen_4syntax_1error_1message), tmp16696, MakeInteger(0), V115)


tmp16699 := Call1(__e, SymFn(symshen_4proc_1nl), PrimGuard2(__e, symcn, PrimStringConcat, MakeString("syntax error here: "), tmp16697))


return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp16699)


panic("unreachable")
})

tmp16700 := Call2(__e, ns2_1set, symshen_4raise_1syntax_1error, tmp16695)


_ = tmp16700

tmp16701 := MakeNative3(func(__e *ControlFlow, V123, V124, V125 Obj) Obj {
tmp16712 := GEq(__e, Nil, V125)

if True == tmp16712 {
return MakeString("\n")
} else {
tmp16710 := GEq(__e, V123, V124)

if True == tmp16710 {
return MakeString("...etc \n")
} else {
tmp16708 := GConsP(__e, V125)

if True == tmp16708 {
tmp16702 := GHd(__e, V125)

tmp16703 := Call3(__e, SymFn(symshen_4app), tmp16702, MakeString(" "), symshen_4s)


tmp16704 := GAdd(__e, V124, MakeInteger(1))

tmp16705 := GTl(__e, V125)

tmp16706 := Call3(__e, SymFn(symshen_4syntax_1error_1message), V123, tmp16704, tmp16705)


return PrimGuard2(__e, symcn, PrimStringConcat, tmp16703, tmp16706)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4syntax_1error_1message)
}


}


}


panic("unreachable")
})

tmp16713 := Call2(__e, ns2_1set, symshen_4syntax_1error_1message, tmp16701)


_ = tmp16713

tmp16714 := MakeNative1(func(__e *ControlFlow, V126 Obj) Obj {
tmp16715 := Call0(__e, SymFn(symfail))


return GEq(__e, V126, tmp16715)


panic("unreachable")
})

tmp16716 := Call2(__e, ns2_1set, symshen_4parse_1failure_2, tmp16714)


_ = tmp16716

tmp16717 := MakeNative1(func(__e *ControlFlow, V127 Obj) Obj {
tmp16718 := Call1(__e, SymFn(symshen_4in_1_6), V127)


return GConsP(__e, tmp16718)


panic("unreachable")
})

tmp16719 := Call2(__e, ns2_1set, symshen_4partial_1parse_1failure_2, tmp16717)


_ = tmp16719

tmp16720 := MakeNative1(func(__e *ControlFlow, V130 Obj) Obj {
tmp16731 := GConsP(__e, V130)

var ifres16724 Obj

if True == tmp16731 {
tmp16729 := GTl(__e, V130)

tmp16730 := GConsP(__e, tmp16729)

var ifres16725 Obj

if True == tmp16730 {
tmp16726 := GTl(__e, V130)

tmp16727 := GTl(__e, tmp16726)

tmp16728 := GEq(__e, Nil, tmp16727)

ifres16725 = tmp16728


} else {
ifres16725 = False


}

ifres16724 = ifres16725


} else {
ifres16724 = False


}

if True == ifres16724 {
tmp16721 := GTl(__e, V130)

return GHd(__e, tmp16721)


} else {
tmp16722 := Call3(__e, SymFn(symshen_4app), V130, MakeString(" is not a YACC stream\n"), symshen_4s)


return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp16722)


}


panic("unreachable")
})

tmp16732 := Call2(__e, ns2_1set, symshen_4objectcode, tmp16720)


_ = tmp16732

tmp16733 := MakeNative1(func(__e *ControlFlow, V131 Obj) Obj {
tmp16734 := MakeNative1(func(__e *ControlFlow, Z132 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5yacc_6), Z132)
panic("unreachable")
})

return TailCall2(__e, SymFn(symcompile), tmp16734, V131)


panic("unreachable")
})

tmp16735 := Call2(__e, ns2_1set, symshen_4yacc_1_6shen, tmp16733)


_ = tmp16735

tmp16736 := MakeNative1(func(__e *ControlFlow, V133 Obj) Obj {
tmp16763 := GConsP(__e, V133)

var ifres16737 Obj

if True == tmp16763 {
tmp16738 := Call1(__e, SymFn(symhead), V133)


let__16549 := tmp16738
_ = let__16549

tmp16739 := Call1(__e, SymFn(symtail), V133)


let__16550 := tmp16739
_ = let__16550

tmp16740 := Call1(__e, SymFn(symshen_4_5yaccsig_6), let__16550)


let__16551 := tmp16740
_ = let__16551

tmp16761 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16551)


var ifres16741 Obj

if True == tmp16761 {
tmp16742 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16741 = tmp16742


} else {
tmp16743 := Call1(__e, SymFn(symshen_4_5_1out), let__16551)


let__16552 := tmp16743
_ = let__16552

tmp16744 := Call1(__e, SymFn(symshen_4in_1_6), let__16551)


let__16553 := tmp16744
_ = let__16553

tmp16745 := Call1(__e, SymFn(symshen_4_5c_1rules_6), let__16553)


let__16554 := tmp16745
_ = let__16554

tmp16760 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16554)


var ifres16746 Obj

if True == tmp16760 {
tmp16747 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16746 = tmp16747


} else {
tmp16748 := Call1(__e, SymFn(symshen_4_5_1out), let__16554)


let__16555 := tmp16748
_ = let__16555

tmp16749 := Call1(__e, SymFn(symshen_4in_1_6), let__16554)


let__16556 := tmp16749
_ = let__16556

tmp16750 := Call1(__e, SymFn(symgensym), symS)


let__16557 := tmp16750
_ = let__16557

tmp16751 := GCons(__e, let__16549, Nil)

tmp16752 := GCons(__e, symdefine, tmp16751)

tmp16753 := Call3(__e, SymFn(symshen_4c_1rules_1_6shen), let__16552, let__16557, let__16555)


tmp16754 := GCons(__e, tmp16753, Nil)

tmp16755 := GCons(__e, sym_1_6, tmp16754)

tmp16756 := GCons(__e, let__16557, tmp16755)

tmp16757 := Call2(__e, SymFn(symappend), let__16552, tmp16756)


tmp16758 := Call2(__e, SymFn(symappend), tmp16752, tmp16757)


let__16558 := tmp16758
_ = let__16558

tmp16759 := Call2(__e, SymFn(symshen_4comb), let__16556, let__16558)


ifres16746 = tmp16759


}

ifres16741 = ifres16746


}

ifres16737 = ifres16741


} else {
tmp16762 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16737 = tmp16762


}

let__16548 := ifres16737
_ = let__16548

tmp16765 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16548)


if True == tmp16765 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16548
}


panic("unreachable")
})

tmp16766 := Call2(__e, ns2_1set, symshen_4_5yacc_6, tmp16736)


_ = tmp16766

tmp16767 := MakeNative1(func(__e *ControlFlow, V145 Obj) Obj {
tmp16823 := GConsP(__e, V145)

var ifres16768 Obj

if True == tmp16823 {
tmp16769 := Call1(__e, SymFn(symhead), V145)


let__16560 := tmp16769
_ = let__16560

tmp16770 := Call1(__e, SymFn(symtail), V145)


let__16561 := tmp16770
_ = let__16561

tmp16821 := Call1(__e, SymFn(symshen_4ccons_2), let__16561)


var ifres16771 Obj

if True == tmp16821 {
tmp16772 := Call1(__e, SymFn(symhead), let__16561)


let__16562 := tmp16772
_ = let__16562

tmp16773 := Call1(__e, SymFn(symtail), let__16561)


let__16563 := tmp16773
_ = let__16563

tmp16819 := Call2(__e, SymFn(symshen_4hds_a_2), let__16562, symlist)


var ifres16774 Obj

if True == tmp16819 {
tmp16775 := Call1(__e, SymFn(symtail), let__16562)


let__16564 := tmp16775
_ = let__16564

tmp16817 := GConsP(__e, let__16564)

var ifres16776 Obj

if True == tmp16817 {
tmp16777 := Call1(__e, SymFn(symhead), let__16564)


let__16565 := tmp16777
_ = let__16565

tmp16778 := Call1(__e, SymFn(symtail), let__16564)


let__16566 := tmp16778
_ = let__16566

tmp16779 := Call1(__e, SymFn(sym_5end_6), let__16566)


let__16567 := tmp16779
_ = let__16567

tmp16815 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16567)


var ifres16780 Obj

if True == tmp16815 {
tmp16781 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16780 = tmp16781


} else {
tmp16782 := Call1(__e, SymFn(symshen_4in_1_6), let__16567)


let__16568 := tmp16782
_ = let__16568

tmp16814 := Call2(__e, SymFn(symshen_4hds_a_2), let__16563, sym_a_a_6)


var ifres16783 Obj

if True == tmp16814 {
tmp16784 := Call1(__e, SymFn(symtail), let__16563)


let__16569 := tmp16784
_ = let__16569

tmp16812 := GConsP(__e, let__16569)

var ifres16785 Obj

if True == tmp16812 {
tmp16786 := Call1(__e, SymFn(symhead), let__16569)


let__16570 := tmp16786
_ = let__16570

tmp16787 := Call1(__e, SymFn(symtail), let__16569)


let__16571 := tmp16787
_ = let__16571

tmp16810 := GConsP(__e, let__16571)

var ifres16788 Obj

if True == tmp16810 {
tmp16789 := Call1(__e, SymFn(symhead), let__16571)


let__16572 := tmp16789
_ = let__16572

tmp16790 := Call1(__e, SymFn(symtail), let__16571)


let__16573 := tmp16790
_ = let__16573

tmp16808 := GEq(__e, sym_i, let__16560)

var ifres16806 Obj

if True == tmp16808 {
tmp16807 := GEq(__e, sym_j, let__16572)

ifres16806 = tmp16807


} else {
ifres16806 = False


}

var ifres16791 Obj

if True == ifres16806 {
tmp16792 := GCons(__e, let__16565, Nil)

tmp16793 := GCons(__e, symlist, tmp16792)

tmp16794 := GCons(__e, let__16565, Nil)

tmp16795 := GCons(__e, symlist, tmp16794)

tmp16796 := GCons(__e, let__16570, Nil)

tmp16797 := GCons(__e, tmp16795, tmp16796)

tmp16798 := GCons(__e, symstr, tmp16797)

tmp16799 := GCons(__e, sym_j, Nil)

tmp16800 := GCons(__e, tmp16798, tmp16799)

tmp16801 := GCons(__e, sym_1_1_6, tmp16800)

tmp16802 := GCons(__e, tmp16793, tmp16801)

tmp16803 := GCons(__e, sym_i, tmp16802)

tmp16804 := Call2(__e, SymFn(symshen_4comb), let__16573, tmp16803)


ifres16791 = tmp16804


} else {
tmp16805 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16791 = tmp16805


}

ifres16788 = ifres16791


} else {
tmp16809 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16788 = tmp16809


}

ifres16785 = ifres16788


} else {
tmp16811 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16785 = tmp16811


}

ifres16783 = ifres16785


} else {
tmp16813 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16783 = tmp16813


}

ifres16780 = ifres16783


}

ifres16776 = ifres16780


} else {
tmp16816 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16776 = tmp16816


}

ifres16774 = ifres16776


} else {
tmp16818 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16774 = tmp16818


}

ifres16771 = ifres16774


} else {
tmp16820 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16771 = tmp16820


}

ifres16768 = ifres16771


} else {
tmp16822 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16768 = tmp16822


}

let__16559 := ifres16768
_ = let__16559

tmp16833 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16559)


if True == tmp16833 {
tmp16824 := Call1(__e, SymFn(sym_5e_6), V145)


let__16575 := tmp16824
_ = let__16575

tmp16829 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16575)


var ifres16825 Obj

if True == tmp16829 {
tmp16826 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16825 = tmp16826


} else {
tmp16827 := Call1(__e, SymFn(symshen_4in_1_6), let__16575)


let__16576 := tmp16827
_ = let__16576

tmp16828 := Call2(__e, SymFn(symshen_4comb), let__16576, Nil)


ifres16825 = tmp16828


}

let__16574 := ifres16825
_ = let__16574

tmp16831 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16574)


if True == tmp16831 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16574
}


} else {
return let__16559
}


panic("unreachable")
})

tmp16834 := Call2(__e, ns2_1set, symshen_4_5yaccsig_6, tmp16767)


_ = tmp16834

tmp16835 := MakeNative1(func(__e *ControlFlow, V164 Obj) Obj {
tmp16836 := Call1(__e, SymFn(symshen_4_5c_1rule_6), V164)


let__16578 := tmp16836
_ = let__16578

tmp16849 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16578)


var ifres16837 Obj

if True == tmp16849 {
tmp16838 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16837 = tmp16838


} else {
tmp16839 := Call1(__e, SymFn(symshen_4_5_1out), let__16578)


let__16579 := tmp16839
_ = let__16579

tmp16840 := Call1(__e, SymFn(symshen_4in_1_6), let__16578)


let__16580 := tmp16840
_ = let__16580

tmp16841 := Call1(__e, SymFn(symshen_4_5c_1rules_6), let__16580)


let__16581 := tmp16841
_ = let__16581

tmp16848 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16581)


var ifres16842 Obj

if True == tmp16848 {
tmp16843 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16842 = tmp16843


} else {
tmp16844 := Call1(__e, SymFn(symshen_4_5_1out), let__16581)


let__16582 := tmp16844
_ = let__16582

tmp16845 := Call1(__e, SymFn(symshen_4in_1_6), let__16581)


let__16583 := tmp16845
_ = let__16583

tmp16846 := GCons(__e, let__16579, let__16582)

tmp16847 := Call2(__e, SymFn(symshen_4comb), let__16583, tmp16846)


ifres16842 = tmp16847


}

ifres16837 = ifres16842


}

let__16577 := ifres16837
_ = let__16577

tmp16865 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16577)


if True == tmp16865 {
tmp16850 := Call1(__e, SymFn(sym_5_b_6), V164)


let__16585 := tmp16850
_ = let__16585

tmp16861 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16585)


var ifres16851 Obj

if True == tmp16861 {
tmp16852 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16851 = tmp16852


} else {
tmp16853 := Call1(__e, SymFn(symshen_4_5_1out), let__16585)


let__16586 := tmp16853
_ = let__16586

tmp16854 := Call1(__e, SymFn(symshen_4in_1_6), let__16585)


let__16587 := tmp16854
_ = let__16587

tmp16859 := Call1(__e, SymFn(symempty_2), let__16586)


var ifres16855 Obj

if True == tmp16859 {
ifres16855 = Nil


} else {
tmp16856 := Call3(__e, SymFn(symshen_4app), let__16586, MakeString("\n ..."), symshen_4r)


tmp16858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("YACC syntax error here:\n "), tmp16856))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("YACC syntax error here:\n "), tmp16856)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()

ifres16855 = tmp16858


}

tmp16860 := Call2(__e, SymFn(symshen_4comb), let__16587, ifres16855)


ifres16851 = tmp16860


}

let__16584 := ifres16851
_ = let__16584

tmp16863 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16584)


if True == tmp16863 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16584
}


} else {
return let__16577
}


panic("unreachable")
})

tmp16866 := Call2(__e, ns2_1set, symshen_4_5c_1rules_6, tmp16835)


_ = tmp16866

tmp16867 := MakeNative1(func(__e *ControlFlow, V176 Obj) Obj {
tmp16868 := Call1(__e, SymFn(symshen_4_5syntax_6), V176)


let__16589 := tmp16868
_ = let__16589

tmp16887 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16589)


var ifres16869 Obj

if True == tmp16887 {
tmp16870 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16869 = tmp16870


} else {
tmp16871 := Call1(__e, SymFn(symshen_4_5_1out), let__16589)


let__16590 := tmp16871
_ = let__16590

tmp16872 := Call1(__e, SymFn(symshen_4in_1_6), let__16589)


let__16591 := tmp16872
_ = let__16591

tmp16873 := Call1(__e, SymFn(symshen_4_5semantics_6), let__16591)


let__16592 := tmp16873
_ = let__16592

tmp16886 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16592)


var ifres16874 Obj

if True == tmp16886 {
tmp16875 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16874 = tmp16875


} else {
tmp16876 := Call1(__e, SymFn(symshen_4_5_1out), let__16592)


let__16593 := tmp16876
_ = let__16593

tmp16877 := Call1(__e, SymFn(symshen_4in_1_6), let__16592)


let__16594 := tmp16877
_ = let__16594

tmp16878 := Call1(__e, SymFn(symshen_4_5sc_6), let__16594)


let__16595 := tmp16878
_ = let__16595

tmp16885 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16595)


var ifres16879 Obj

if True == tmp16885 {
tmp16880 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16879 = tmp16880


} else {
tmp16881 := Call1(__e, SymFn(symshen_4in_1_6), let__16595)


let__16596 := tmp16881
_ = let__16596

tmp16882 := GCons(__e, let__16593, Nil)

tmp16883 := GCons(__e, let__16590, tmp16882)

tmp16884 := Call2(__e, SymFn(symshen_4comb), let__16596, tmp16883)


ifres16879 = tmp16884


}

ifres16874 = ifres16879


}

ifres16869 = ifres16874


}

let__16588 := ifres16869
_ = let__16588

tmp16906 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16588)


if True == tmp16906 {
tmp16888 := Call1(__e, SymFn(symshen_4_5syntax_6), V176)


let__16598 := tmp16888
_ = let__16598

tmp16902 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16598)


var ifres16889 Obj

if True == tmp16902 {
tmp16890 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16889 = tmp16890


} else {
tmp16891 := Call1(__e, SymFn(symshen_4_5_1out), let__16598)


let__16599 := tmp16891
_ = let__16599

tmp16892 := Call1(__e, SymFn(symshen_4in_1_6), let__16598)


let__16600 := tmp16892
_ = let__16600

tmp16893 := Call1(__e, SymFn(symshen_4_5sc_6), let__16600)


let__16601 := tmp16893
_ = let__16601

tmp16901 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16601)


var ifres16894 Obj

if True == tmp16901 {
tmp16895 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16894 = tmp16895


} else {
tmp16896 := Call1(__e, SymFn(symshen_4in_1_6), let__16601)


let__16602 := tmp16896
_ = let__16602

tmp16897 := Call1(__e, SymFn(symshen_4autocomplete), let__16599)


tmp16898 := GCons(__e, tmp16897, Nil)

tmp16899 := GCons(__e, let__16599, tmp16898)

tmp16900 := Call2(__e, SymFn(symshen_4comb), let__16602, tmp16899)


ifres16894 = tmp16900


}

ifres16889 = ifres16894


}

let__16597 := ifres16889
_ = let__16597

tmp16904 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16597)


if True == tmp16904 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16597
}


} else {
return let__16588
}


panic("unreachable")
})

tmp16907 := Call2(__e, ns2_1set, symshen_4_5c_1rule_6, tmp16867)


_ = tmp16907

tmp16908 := MakeNative1(func(__e *ControlFlow, V192 Obj) Obj {
tmp16934 := GConsP(__e, V192)

var ifres16928 Obj

if True == tmp16934 {
tmp16932 := GTl(__e, V192)

tmp16933 := GEq(__e, Nil, tmp16932)

var ifres16929 Obj

if True == tmp16933 {
tmp16930 := GHd(__e, V192)

tmp16931 := Call1(__e, SymFn(symshen_4non_1terminal_2), tmp16930)


ifres16929 = tmp16931


} else {
ifres16929 = False


}

ifres16928 = ifres16929


} else {
ifres16928 = False


}

if True == ifres16928 {
return GHd(__e, V192)
} else {
tmp16926 := GConsP(__e, V192)

var ifres16923 Obj

if True == tmp16926 {
tmp16924 := GHd(__e, V192)

tmp16925 := Call1(__e, SymFn(symshen_4non_1terminal_2), tmp16924)


ifres16923 = tmp16925


} else {
ifres16923 = False


}

if True == ifres16923 {
tmp16909 := GHd(__e, V192)

tmp16910 := GTl(__e, V192)

tmp16911 := Call1(__e, SymFn(symshen_4autocomplete), tmp16910)


tmp16912 := GCons(__e, tmp16911, Nil)

tmp16913 := GCons(__e, tmp16909, tmp16912)

return GCons(__e, symappend, tmp16913)


} else {
tmp16921 := GConsP(__e, V192)

if True == tmp16921 {
tmp16914 := GHd(__e, V192)

tmp16915 := Call1(__e, SymFn(symshen_4autocomplete), tmp16914)


tmp16916 := GTl(__e, V192)

tmp16917 := Call1(__e, SymFn(symshen_4autocomplete), tmp16916)


tmp16918 := GCons(__e, tmp16917, Nil)

tmp16919 := GCons(__e, tmp16915, tmp16918)

return GCons(__e, symcons, tmp16919)


} else {
return V192
}


}


}


panic("unreachable")
})

tmp16935 := Call2(__e, ns2_1set, symshen_4autocomplete, tmp16908)


_ = tmp16935

tmp16936 := MakeNative1(func(__e *ControlFlow, V193 Obj) Obj {
tmp16940 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V193)

if True == tmp16940 {
tmp16937 := Call1(__e, SymFn(symexplode), V193)


let__16603 := tmp16937
_ = let__16603

tmp16938 := MakeNative1(func(__e *ControlFlow, Z195 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5non_1terminal_2_6), Z195)
panic("unreachable")
})

return TailCall2(__e, SymFn(symcompile), tmp16938, let__16603)


} else {
return False
}


panic("unreachable")
})

tmp16941 := Call2(__e, ns2_1set, symshen_4non_1terminal_2, tmp16936)


_ = tmp16941

tmp16942 := MakeNative1(func(__e *ControlFlow, V196 Obj) Obj {
tmp16943 := Call1(__e, SymFn(symshen_4_5packagenames_6), V196)


let__16605 := tmp16943
_ = let__16605

tmp16953 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16605)


var ifres16944 Obj

if True == tmp16953 {
tmp16945 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16944 = tmp16945


} else {
tmp16946 := Call1(__e, SymFn(symshen_4in_1_6), let__16605)


let__16606 := tmp16946
_ = let__16606

tmp16947 := Call1(__e, SymFn(symshen_4_5non_1terminal_1name_6), let__16606)


let__16607 := tmp16947
_ = let__16607

tmp16952 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16607)


var ifres16948 Obj

if True == tmp16952 {
tmp16949 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16948 = tmp16949


} else {
tmp16950 := Call1(__e, SymFn(symshen_4in_1_6), let__16607)


let__16608 := tmp16950
_ = let__16608

tmp16951 := Call2(__e, SymFn(symshen_4comb), let__16608, True)


ifres16948 = tmp16951


}

ifres16944 = ifres16948


}

let__16604 := ifres16944
_ = let__16604

tmp16971 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16604)


if True == tmp16971 {
tmp16954 := Call1(__e, SymFn(symshen_4_5non_1terminal_1name_6), V196)


let__16610 := tmp16954
_ = let__16610

tmp16959 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16610)


var ifres16955 Obj

if True == tmp16959 {
tmp16956 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16955 = tmp16956


} else {
tmp16957 := Call1(__e, SymFn(symshen_4in_1_6), let__16610)


let__16611 := tmp16957
_ = let__16611

tmp16958 := Call2(__e, SymFn(symshen_4comb), let__16611, True)


ifres16955 = tmp16958


}

let__16609 := ifres16955
_ = let__16609

tmp16969 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16609)


if True == tmp16969 {
tmp16960 := Call1(__e, SymFn(sym_5_b_6), V196)


let__16613 := tmp16960
_ = let__16613

tmp16965 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16613)


var ifres16961 Obj

if True == tmp16965 {
tmp16962 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16961 = tmp16962


} else {
tmp16963 := Call1(__e, SymFn(symshen_4in_1_6), let__16613)


let__16614 := tmp16963
_ = let__16614

tmp16964 := Call2(__e, SymFn(symshen_4comb), let__16614, False)


ifres16961 = tmp16964


}

let__16612 := ifres16961
_ = let__16612

tmp16967 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16612)


if True == tmp16967 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16612
}


} else {
return let__16609
}


} else {
return let__16604
}


panic("unreachable")
})

tmp16972 := Call2(__e, ns2_1set, symshen_4_5non_1terminal_2_6, tmp16942)


_ = tmp16972

tmp16973 := MakeNative1(func(__e *ControlFlow, V208 Obj) Obj {
tmp16974 := Call1(__e, SymFn(symshen_4_5packagename_6), V208)


let__16616 := tmp16974
_ = let__16616

tmp16988 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16616)


var ifres16975 Obj

if True == tmp16988 {
tmp16976 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16975 = tmp16976


} else {
tmp16977 := Call1(__e, SymFn(symshen_4in_1_6), let__16616)


let__16617 := tmp16977
_ = let__16617

tmp16987 := Call2(__e, SymFn(symshen_4hds_a_2), let__16617, MakeString("."))


var ifres16978 Obj

if True == tmp16987 {
tmp16979 := Call1(__e, SymFn(symtail), let__16617)


let__16618 := tmp16979
_ = let__16618

tmp16980 := Call1(__e, SymFn(symshen_4_5packagenames_6), let__16618)


let__16619 := tmp16980
_ = let__16619

tmp16985 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16619)


var ifres16981 Obj

if True == tmp16985 {
tmp16982 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16981 = tmp16982


} else {
tmp16983 := Call1(__e, SymFn(symshen_4in_1_6), let__16619)


let__16620 := tmp16983
_ = let__16620

tmp16984 := Call2(__e, SymFn(symshen_4comb), let__16620, symshen_4skip)


ifres16981 = tmp16984


}

ifres16978 = ifres16981


} else {
tmp16986 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16978 = tmp16986


}

ifres16975 = ifres16978


}

let__16615 := ifres16975
_ = let__16615

tmp17002 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16615)


if True == tmp17002 {
tmp16989 := Call1(__e, SymFn(symshen_4_5packagename_6), V208)


let__16622 := tmp16989
_ = let__16622

tmp16998 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16622)


var ifres16990 Obj

if True == tmp16998 {
tmp16991 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16990 = tmp16991


} else {
tmp16992 := Call1(__e, SymFn(symshen_4in_1_6), let__16622)


let__16623 := tmp16992
_ = let__16623

tmp16997 := Call2(__e, SymFn(symshen_4hds_a_2), let__16623, MakeString("."))


var ifres16993 Obj

if True == tmp16997 {
tmp16994 := Call1(__e, SymFn(symtail), let__16623)


let__16624 := tmp16994
_ = let__16624

tmp16995 := Call2(__e, SymFn(symshen_4comb), let__16624, symshen_4skip)


ifres16993 = tmp16995


} else {
tmp16996 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres16993 = tmp16996


}

ifres16990 = ifres16993


}

let__16621 := ifres16990
_ = let__16621

tmp17000 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16621)


if True == tmp17000 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16621
}


} else {
return let__16615
}


panic("unreachable")
})

tmp17003 := Call2(__e, ns2_1set, symshen_4_5packagenames_6, tmp16973)


_ = tmp17003

tmp17004 := MakeNative1(func(__e *ControlFlow, V219 Obj) Obj {
tmp17005 := Call1(__e, SymFn(symshen_4_5packagechar_6), V219)


let__16626 := tmp17005
_ = let__16626

tmp17015 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16626)


var ifres17006 Obj

if True == tmp17015 {
tmp17007 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17006 = tmp17007


} else {
tmp17008 := Call1(__e, SymFn(symshen_4in_1_6), let__16626)


let__16627 := tmp17008
_ = let__16627

tmp17009 := Call1(__e, SymFn(symshen_4_5packagename_6), let__16627)


let__16628 := tmp17009
_ = let__16628

tmp17014 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16628)


var ifres17010 Obj

if True == tmp17014 {
tmp17011 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17010 = tmp17011


} else {
tmp17012 := Call1(__e, SymFn(symshen_4in_1_6), let__16628)


let__16629 := tmp17012
_ = let__16629

tmp17013 := Call2(__e, SymFn(symshen_4comb), let__16629, symshen_4skip)


ifres17010 = tmp17013


}

ifres17006 = ifres17010


}

let__16625 := ifres17006
_ = let__16625

tmp17025 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16625)


if True == tmp17025 {
tmp17016 := Call1(__e, SymFn(sym_5e_6), V219)


let__16631 := tmp17016
_ = let__16631

tmp17021 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16631)


var ifres17017 Obj

if True == tmp17021 {
tmp17018 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17017 = tmp17018


} else {
tmp17019 := Call1(__e, SymFn(symshen_4in_1_6), let__16631)


let__16632 := tmp17019
_ = let__16632

tmp17020 := Call2(__e, SymFn(symshen_4comb), let__16632, symshen_4skip)


ifres17017 = tmp17020


}

let__16630 := ifres17017
_ = let__16630

tmp17023 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16630)


if True == tmp17023 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16630
}


} else {
return let__16625
}


panic("unreachable")
})

tmp17026 := Call2(__e, ns2_1set, symshen_4_5packagename_6, tmp17004)


_ = tmp17026

tmp17027 := MakeNative1(func(__e *ControlFlow, V228 Obj) Obj {
tmp17037 := GConsP(__e, V228)

var ifres17028 Obj

if True == tmp17037 {
tmp17029 := Call1(__e, SymFn(symhead), V228)


let__16634 := tmp17029
_ = let__16634

tmp17030 := Call1(__e, SymFn(symtail), V228)


let__16635 := tmp17030
_ = let__16635

tmp17034 := GEq(__e, let__16634, MakeString("."))

tmp17035 := GNot(__e, tmp17034)

var ifres17031 Obj

if True == tmp17035 {
tmp17032 := Call2(__e, SymFn(symshen_4comb), let__16635, symshen_4skip)


ifres17031 = tmp17032


} else {
tmp17033 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17031 = tmp17033


}

ifres17028 = ifres17031


} else {
tmp17036 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17028 = tmp17036


}

let__16633 := ifres17028
_ = let__16633

tmp17039 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16633)


if True == tmp17039 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16633
}


panic("unreachable")
})

tmp17040 := Call2(__e, ns2_1set, symshen_4_5packagechar_6, tmp17027)


_ = tmp17040

tmp17041 := MakeNative1(func(__e *ControlFlow, V232 Obj) Obj {
tmp17059 := Call2(__e, SymFn(symshen_4hds_a_2), V232, MakeString("<"))


var ifres17042 Obj

if True == tmp17059 {
tmp17043 := Call1(__e, SymFn(symtail), V232)


let__16637 := tmp17043
_ = let__16637

tmp17044 := Call1(__e, SymFn(sym_5_b_6), let__16637)


let__16638 := tmp17044
_ = let__16638

tmp17057 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16638)


var ifres17045 Obj

if True == tmp17057 {
tmp17046 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17045 = tmp17046


} else {
tmp17047 := Call1(__e, SymFn(symshen_4_5_1out), let__16638)


let__16639 := tmp17047
_ = let__16639

tmp17048 := Call1(__e, SymFn(symshen_4in_1_6), let__16638)


let__16640 := tmp17048
_ = let__16640

tmp17052 := Call1(__e, SymFn(symreverse), let__16639)


let__16641 := tmp17052
_ = let__16641

tmp17056 := GConsP(__e, let__16641)

var ifres17053 Obj

if True == tmp17056 {
tmp17054 := GHd(__e, let__16641)

tmp17055 := GEq(__e, tmp17054, MakeString(">"))

ifres17053 = tmp17055


} else {
ifres17053 = False


}

var ifres17049 Obj

if True == ifres17053 {
tmp17050 := Call2(__e, SymFn(symshen_4comb), let__16640, symshen_4skip)


ifres17049 = tmp17050


} else {
tmp17051 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17049 = tmp17051


}

ifres17045 = ifres17049


}

ifres17042 = ifres17045


} else {
tmp17058 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17042 = tmp17058


}

let__16636 := ifres17042
_ = let__16636

tmp17061 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16636)


if True == tmp17061 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16636
}


panic("unreachable")
})

tmp17062 := Call2(__e, ns2_1set, symshen_4_5non_1terminal_1name_6, tmp17041)


_ = tmp17062

tmp17063 := MakeNative1(func(__e *ControlFlow, V239 Obj) Obj {
tmp17064 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

return GEq(__e, V239, tmp17064)


panic("unreachable")
})

tmp17065 := Call2(__e, ns2_1set, symshen_4semicolon_2, tmp17063)


_ = tmp17065

tmp17066 := MakeNative1(func(__e *ControlFlow, V240 Obj) Obj {
tmp17075 := GConsP(__e, V240)

var ifres17067 Obj

if True == tmp17075 {
tmp17068 := Call1(__e, SymFn(symhead), V240)


let__16643 := tmp17068
_ = let__16643

tmp17069 := Call1(__e, SymFn(symtail), V240)


let__16644 := tmp17069
_ = let__16644

tmp17073 := Call1(__e, SymFn(symshen_4colon_1equal_2), let__16643)


var ifres17070 Obj

if True == tmp17073 {
tmp17071 := Call2(__e, SymFn(symshen_4comb), let__16644, symshen_4skip)


ifres17070 = tmp17071


} else {
tmp17072 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17070 = tmp17072


}

ifres17067 = ifres17070


} else {
tmp17074 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17067 = tmp17074


}

let__16642 := ifres17067
_ = let__16642

tmp17077 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16642)


if True == tmp17077 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16642
}


panic("unreachable")
})

tmp17078 := Call2(__e, ns2_1set, symshen_4_5colon_1equal_6, tmp17066)


_ = tmp17078

tmp17079 := MakeNative1(func(__e *ControlFlow, V244 Obj) Obj {
tmp17080 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":="))

return GEq(__e, tmp17080, V244)


panic("unreachable")
})

tmp17081 := Call2(__e, ns2_1set, symshen_4colon_1equal_2, tmp17079)


_ = tmp17081

tmp17082 := MakeNative1(func(__e *ControlFlow, V245 Obj) Obj {
tmp17083 := Call1(__e, SymFn(symshen_4_5syntax_1item_6), V245)


let__16646 := tmp17083
_ = let__16646

tmp17096 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16646)


var ifres17084 Obj

if True == tmp17096 {
tmp17085 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17084 = tmp17085


} else {
tmp17086 := Call1(__e, SymFn(symshen_4_5_1out), let__16646)


let__16647 := tmp17086
_ = let__16647

tmp17087 := Call1(__e, SymFn(symshen_4in_1_6), let__16646)


let__16648 := tmp17087
_ = let__16648

tmp17088 := Call1(__e, SymFn(symshen_4_5syntax_6), let__16648)


let__16649 := tmp17088
_ = let__16649

tmp17095 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16649)


var ifres17089 Obj

if True == tmp17095 {
tmp17090 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17089 = tmp17090


} else {
tmp17091 := Call1(__e, SymFn(symshen_4_5_1out), let__16649)


let__16650 := tmp17091
_ = let__16650

tmp17092 := Call1(__e, SymFn(symshen_4in_1_6), let__16649)


let__16651 := tmp17092
_ = let__16651

tmp17093 := GCons(__e, let__16647, let__16650)

tmp17094 := Call2(__e, SymFn(symshen_4comb), let__16651, tmp17093)


ifres17089 = tmp17094


}

ifres17084 = ifres17089


}

let__16645 := ifres17084
_ = let__16645

tmp17108 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16645)


if True == tmp17108 {
tmp17097 := Call1(__e, SymFn(symshen_4_5syntax_1item_6), V245)


let__16653 := tmp17097
_ = let__16653

tmp17104 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16653)


var ifres17098 Obj

if True == tmp17104 {
tmp17099 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17098 = tmp17099


} else {
tmp17100 := Call1(__e, SymFn(symshen_4_5_1out), let__16653)


let__16654 := tmp17100
_ = let__16654

tmp17101 := Call1(__e, SymFn(symshen_4in_1_6), let__16653)


let__16655 := tmp17101
_ = let__16655

tmp17102 := GCons(__e, let__16654, Nil)

tmp17103 := Call2(__e, SymFn(symshen_4comb), let__16655, tmp17102)


ifres17098 = tmp17103


}

let__16652 := ifres17098
_ = let__16652

tmp17106 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16652)


if True == tmp17106 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16652
}


} else {
return let__16645
}


panic("unreachable")
})

tmp17109 := Call2(__e, ns2_1set, symshen_4_5syntax_6, tmp17082)


_ = tmp17109

tmp17110 := MakeNative1(func(__e *ControlFlow, V257 Obj) Obj {
tmp17119 := GConsP(__e, V257)

var ifres17111 Obj

if True == tmp17119 {
tmp17112 := Call1(__e, SymFn(symhead), V257)


let__16657 := tmp17112
_ = let__16657

tmp17113 := Call1(__e, SymFn(symtail), V257)


let__16658 := tmp17113
_ = let__16658

tmp17117 := Call1(__e, SymFn(symshen_4syntax_1item_2), let__16657)


var ifres17114 Obj

if True == tmp17117 {
tmp17115 := Call2(__e, SymFn(symshen_4comb), let__16658, let__16657)


ifres17114 = tmp17115


} else {
tmp17116 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17114 = tmp17116


}

ifres17111 = ifres17114


} else {
tmp17118 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17111 = tmp17118


}

let__16656 := ifres17111
_ = let__16656

tmp17121 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16656)


if True == tmp17121 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16656
}


panic("unreachable")
})

tmp17122 := Call2(__e, ns2_1set, symshen_4_5syntax_1item_6, tmp17110)


_ = tmp17122

tmp17123 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V263 := __self1
_ = V263
tmp17153 := Call1(__e, SymFn(symshen_4colon_1equal_2), V263)


if True == tmp17153 {
return False
} else {
tmp17151 := Call1(__e, SymFn(symshen_4semicolon_2), V263)


if True == tmp17151 {
return False
} else {
tmp17149 := Call1(__e, SymFn(symatom_2), V263)


if True == tmp17149 {
return True
} else {
tmp17147 := GConsP(__e, V263)

var ifres17132 Obj

if True == tmp17147 {
tmp17145 := GHd(__e, V263)

tmp17146 := GEq(__e, symcons, tmp17145)

var ifres17133 Obj

if True == tmp17146 {
tmp17143 := GTl(__e, V263)

tmp17144 := GConsP(__e, tmp17143)

var ifres17134 Obj

if True == tmp17144 {
tmp17140 := GTl(__e, V263)

tmp17141 := GTl(__e, tmp17140)

tmp17142 := GConsP(__e, tmp17141)

var ifres17135 Obj

if True == tmp17142 {
tmp17136 := GTl(__e, V263)

tmp17137 := GTl(__e, tmp17136)

tmp17138 := GTl(__e, tmp17137)

tmp17139 := GEq(__e, Nil, tmp17138)

ifres17135 = tmp17139


} else {
ifres17135 = False


}

ifres17134 = ifres17135


} else {
ifres17134 = False


}

ifres17133 = ifres17134


} else {
ifres17133 = False


}

ifres17132 = ifres17133


} else {
ifres17132 = False


}

if True == ifres17132 {
tmp17128 := GTl(__e, V263)

tmp17129 := GHd(__e, tmp17128)

tmp17130 := Call1(__e, SymFn(symshen_4syntax_1item_2), tmp17129)


if True == tmp17130 {
tmp17124 := GTl(__e, V263)

tmp17125 := GTl(__e, tmp17124)

tmp17126 := GHd(__e, tmp17125)

if SymFn(symshen_4syntax_1item_2) == __self {
__self1 = tmp17126
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4syntax_1item_2), tmp17126)


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

tmp17154 := Call2(__e, ns2_1set, symshen_4syntax_1item_2, tmp17123)


_ = tmp17154

tmp17155 := MakeNative1(func(__e *ControlFlow, V264 Obj) Obj {
tmp17156 := Call1(__e, SymFn(symshen_4_5colon_1equal_6), V264)


let__16660 := tmp17156
_ = let__16660

tmp17182 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16660)


var ifres17157 Obj

if True == tmp17182 {
tmp17158 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17157 = tmp17158


} else {
tmp17159 := Call1(__e, SymFn(symshen_4in_1_6), let__16660)


let__16661 := tmp17159
_ = let__16661

tmp17181 := GConsP(__e, let__16661)

var ifres17160 Obj

if True == tmp17181 {
tmp17161 := Call1(__e, SymFn(symhead), let__16661)


let__16662 := tmp17161
_ = let__16662

tmp17162 := Call1(__e, SymFn(symtail), let__16661)


let__16663 := tmp17162
_ = let__16663

tmp17179 := Call2(__e, SymFn(symshen_4hds_a_2), let__16663, symwhere)


var ifres17163 Obj

if True == tmp17179 {
tmp17164 := Call1(__e, SymFn(symtail), let__16663)


let__16664 := tmp17164
_ = let__16664

tmp17177 := GConsP(__e, let__16664)

var ifres17165 Obj

if True == tmp17177 {
tmp17166 := Call1(__e, SymFn(symhead), let__16664)


let__16665 := tmp17166
_ = let__16665

tmp17167 := Call1(__e, SymFn(symtail), let__16664)


let__16666 := tmp17167
_ = let__16666

tmp17174 := Call1(__e, SymFn(symshen_4semicolon_2), let__16662)


tmp17175 := GNot(__e, tmp17174)

var ifres17168 Obj

if True == tmp17175 {
tmp17169 := GCons(__e, let__16662, Nil)

tmp17170 := GCons(__e, let__16665, tmp17169)

tmp17171 := GCons(__e, symwhere, tmp17170)

tmp17172 := Call2(__e, SymFn(symshen_4comb), let__16666, tmp17171)


ifres17168 = tmp17172


} else {
tmp17173 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17168 = tmp17173


}

ifres17165 = ifres17168


} else {
tmp17176 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17165 = tmp17176


}

ifres17163 = ifres17165


} else {
tmp17178 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17163 = tmp17178


}

ifres17160 = ifres17163


} else {
tmp17180 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17160 = tmp17180


}

ifres17157 = ifres17160


}

let__16659 := ifres17157
_ = let__16659

tmp17201 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16659)


if True == tmp17201 {
tmp17183 := Call1(__e, SymFn(symshen_4_5colon_1equal_6), V264)


let__16668 := tmp17183
_ = let__16668

tmp17197 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16668)


var ifres17184 Obj

if True == tmp17197 {
tmp17185 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17184 = tmp17185


} else {
tmp17186 := Call1(__e, SymFn(symshen_4in_1_6), let__16668)


let__16669 := tmp17186
_ = let__16669

tmp17196 := GConsP(__e, let__16669)

var ifres17187 Obj

if True == tmp17196 {
tmp17188 := Call1(__e, SymFn(symhead), let__16669)


let__16670 := tmp17188
_ = let__16670

tmp17189 := Call1(__e, SymFn(symtail), let__16669)


let__16671 := tmp17189
_ = let__16671

tmp17193 := Call1(__e, SymFn(symshen_4semicolon_2), let__16670)


tmp17194 := GNot(__e, tmp17193)

var ifres17190 Obj

if True == tmp17194 {
tmp17191 := Call2(__e, SymFn(symshen_4comb), let__16671, let__16670)


ifres17190 = tmp17191


} else {
tmp17192 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17190 = tmp17192


}

ifres17187 = ifres17190


} else {
tmp17195 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17187 = tmp17195


}

ifres17184 = ifres17187


}

let__16667 := ifres17184
_ = let__16667

tmp17199 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__16667)


if True == tmp17199 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__16667
}


} else {
return let__16659
}


panic("unreachable")
})

tmp17202 := Call2(__e, ns2_1set, symshen_4_5semantics_6, tmp17155)


_ = tmp17202

tmp17203 := MakeNative3(func(__e *ControlFlow, V286, V287, V288 Obj) Obj {
tmp17211 := GEq(__e, Nil, V288)

if True == tmp17211 {
return GCons(__e, symshen_4parse_1failure, Nil)
} else {
tmp17209 := GConsP(__e, V288)

if True == tmp17209 {
tmp17204 := GHd(__e, V288)

tmp17205 := Call3(__e, SymFn(symshen_4c_1rule_1_6shen), V286, tmp17204, V287)


tmp17206 := GTl(__e, V288)

tmp17207 := Call3(__e, SymFn(symshen_4c_1rules_1_6shen), V286, V287, tmp17206)


return TailCall2(__e, SymFn(symshen_4combine_1c_1code), tmp17205, tmp17207)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.c-rules->shen\n"))
}


}


panic("unreachable")
})

tmp17212 := Call2(__e, ns2_1set, symshen_4c_1rules_1_6shen, tmp17203)


_ = tmp17212

tmp17213 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall0(__e, SymFn(symfail))
panic("unreachable")
})

tmp17214 := Call2(__e, ns2_1set, symshen_4parse_1failure, tmp17213)


_ = tmp17214

tmp17215 := MakeNative2(func(__e *ControlFlow, V289, V290 Obj) Obj {
tmp17216 := GCons(__e, symResult, Nil)

tmp17217 := GCons(__e, symshen_4parse_1failure_2, tmp17216)

tmp17218 := GCons(__e, symResult, Nil)

tmp17219 := GCons(__e, V290, tmp17218)

tmp17220 := GCons(__e, tmp17217, tmp17219)

tmp17221 := GCons(__e, symif, tmp17220)

tmp17222 := GCons(__e, tmp17221, Nil)

tmp17223 := GCons(__e, V289, tmp17222)

tmp17224 := GCons(__e, symResult, tmp17223)

return GCons(__e, symlet, tmp17224)


panic("unreachable")
})

tmp17225 := Call2(__e, ns2_1set, symshen_4combine_1c_1code, tmp17215)


_ = tmp17225

tmp17226 := MakeNative3(func(__e *ControlFlow, V297, V298, V299 Obj) Obj {
tmp17238 := GConsP(__e, V298)

var ifres17231 Obj

if True == tmp17238 {
tmp17236 := GTl(__e, V298)

tmp17237 := GConsP(__e, tmp17236)

var ifres17232 Obj

if True == tmp17237 {
tmp17233 := GTl(__e, V298)

tmp17234 := GTl(__e, tmp17233)

tmp17235 := GEq(__e, Nil, tmp17234)

ifres17232 = tmp17235


} else {
ifres17232 = False


}

ifres17231 = ifres17232


} else {
ifres17231 = False


}

if True == ifres17231 {
tmp17227 := GHd(__e, V298)

tmp17228 := GTl(__e, V298)

tmp17229 := GHd(__e, tmp17228)

return TailCall4(__e, SymFn(symshen_4yacc_1syntax), V297, V299, tmp17227, tmp17229)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.c-rule->shen\n"))
}


panic("unreachable")
})

tmp17239 := Call2(__e, ns2_1set, symshen_4c_1rule_1_6shen, tmp17226)


_ = tmp17239

tmp17240 := MakeNative4(func(__e *ControlFlow, V308, V309, V310, V311 Obj) Obj {
tmp17299 := GEq(__e, Nil, V310)

var ifres17282 Obj

if True == tmp17299 {
tmp17298 := GConsP(__e, V311)

var ifres17283 Obj

if True == tmp17298 {
tmp17296 := GHd(__e, V311)

tmp17297 := GEq(__e, symwhere, tmp17296)

var ifres17284 Obj

if True == tmp17297 {
tmp17294 := GTl(__e, V311)

tmp17295 := GConsP(__e, tmp17294)

var ifres17285 Obj

if True == tmp17295 {
tmp17291 := GTl(__e, V311)

tmp17292 := GTl(__e, tmp17291)

tmp17293 := GConsP(__e, tmp17292)

var ifres17286 Obj

if True == tmp17293 {
tmp17287 := GTl(__e, V311)

tmp17288 := GTl(__e, tmp17287)

tmp17289 := GTl(__e, tmp17288)

tmp17290 := GEq(__e, Nil, tmp17289)

ifres17286 = tmp17290


} else {
ifres17286 = False


}

ifres17285 = ifres17286


} else {
ifres17285 = False


}

ifres17284 = ifres17285


} else {
ifres17284 = False


}

ifres17283 = ifres17284


} else {
ifres17283 = False


}

ifres17282 = ifres17283


} else {
ifres17282 = False


}

if True == ifres17282 {
tmp17241 := GTl(__e, V311)

tmp17242 := GHd(__e, tmp17241)

tmp17243 := Call1(__e, SymFn(symshen_4process_1yacc_1semantics), tmp17242)


tmp17244 := GTl(__e, V311)

tmp17245 := GTl(__e, tmp17244)

tmp17246 := GHd(__e, tmp17245)

tmp17247 := Call4(__e, SymFn(symshen_4yacc_1syntax), V308, V309, Nil, tmp17246)


tmp17248 := GCons(__e, symshen_4parse_1failure, Nil)

tmp17249 := GCons(__e, tmp17248, Nil)

tmp17250 := GCons(__e, tmp17247, tmp17249)

tmp17251 := GCons(__e, tmp17243, tmp17250)

return GCons(__e, symif, tmp17251)


} else {
tmp17280 := GEq(__e, Nil, V310)

if True == tmp17280 {
return TailCall3(__e, SymFn(symshen_4yacc_1semantics), V308, V309, V311)
} else {
tmp17278 := GConsP(__e, V310)

if True == tmp17278 {
tmp17275 := GHd(__e, V310)

tmp17276 := Call1(__e, SymFn(symshen_4non_1terminal_2), tmp17275)


if True == tmp17276 {
tmp17252 := GHd(__e, V310)

tmp17253 := GTl(__e, V310)

return TailCall5(__e, SymFn(symshen_4non_1terminalcode), V308, V309, tmp17252, tmp17253, V311)


} else {
tmp17272 := GHd(__e, V310)

tmp17273 := PrimGuard1(__e, symvariable_2, PrimIsVariable, tmp17272)

if True == tmp17273 {
tmp17254 := GHd(__e, V310)

tmp17255 := GTl(__e, V310)

return TailCall5(__e, SymFn(symshen_4variablecode), V308, V309, tmp17254, tmp17255, V311)


} else {
tmp17269 := GHd(__e, V310)

tmp17270 := GEq(__e, sym__, tmp17269)

if True == tmp17270 {
tmp17256 := GHd(__e, V310)

tmp17257 := GTl(__e, V310)

return TailCall5(__e, SymFn(symshen_4wildcardcode), V308, V309, tmp17256, tmp17257, V311)


} else {
tmp17266 := GHd(__e, V310)

tmp17267 := Call1(__e, SymFn(symatom_2), tmp17266)


if True == tmp17267 {
tmp17258 := GHd(__e, V310)

tmp17259 := GTl(__e, V310)

return TailCall5(__e, SymFn(symshen_4terminalcode), V308, V309, tmp17258, tmp17259, V311)


} else {
tmp17263 := GHd(__e, V310)

tmp17264 := GConsP(__e, tmp17263)

if True == tmp17264 {
tmp17260 := GHd(__e, V310)

tmp17261 := GTl(__e, V310)

return TailCall5(__e, SymFn(symshen_4conscode), V308, V309, tmp17260, tmp17261, V311)


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

tmp17300 := Call2(__e, ns2_1set, symshen_4yacc_1syntax, tmp17240)


_ = tmp17300

tmp17301 := MakeNative5(func(__e *ControlFlow, V312, V313, V314, V315, V316 Obj) Obj {
tmp17302 := Call2(__e, SymFn(symconcat), symParse, V314)


let__16672 := tmp17302
_ = let__16672

tmp17303 := Call2(__e, SymFn(symconcat), symAction, V314)


let__16673 := tmp17303
_ = let__16673

tmp17304 := Call2(__e, SymFn(symconcat), symRemainder, V314)


let__16674 := tmp17304
_ = let__16674

tmp17305 := GCons(__e, V313, Nil)

tmp17306 := GCons(__e, V314, tmp17305)

tmp17307 := GCons(__e, let__16672, Nil)

tmp17308 := GCons(__e, symshen_4parse_1failure_2, tmp17307)

tmp17309 := GCons(__e, symshen_4parse_1failure, Nil)

tmp17310 := GCons(__e, let__16672, Nil)

tmp17311 := GCons(__e, symshen_4in_1_6, tmp17310)

tmp17312 := Call4(__e, SymFn(symshen_4yacc_1syntax), V312, let__16674, V315, V316)


tmp17313 := GCons(__e, tmp17312, Nil)

tmp17314 := GCons(__e, tmp17311, tmp17313)

tmp17315 := GCons(__e, let__16674, tmp17314)

tmp17316 := GCons(__e, symlet, tmp17315)

let__16675 := tmp17316
_ = let__16675

tmp17326 := Call2(__e, SymFn(symshen_4occurs_1check_2), V314, V316)


var ifres17324 Obj

if True == tmp17326 {
ifres17324 = True


} else {
tmp17325 := Call2(__e, SymFn(symshen_4occurs_1check_2), let__16673, V316)


ifres17324 = tmp17325


}

var ifres17317 Obj

if True == ifres17324 {
tmp17318 := GCons(__e, let__16672, Nil)

tmp17319 := GCons(__e, symshen_4_5_1out, tmp17318)

tmp17320 := GCons(__e, let__16675, Nil)

tmp17321 := GCons(__e, tmp17319, tmp17320)

tmp17322 := GCons(__e, let__16673, tmp17321)

tmp17323 := GCons(__e, symlet, tmp17322)

ifres17317 = tmp17323


} else {
ifres17317 = let__16675


}

tmp17327 := GCons(__e, ifres17317, Nil)

tmp17328 := GCons(__e, tmp17309, tmp17327)

tmp17329 := GCons(__e, tmp17308, tmp17328)

tmp17330 := GCons(__e, symif, tmp17329)

tmp17331 := GCons(__e, tmp17330, Nil)

tmp17332 := GCons(__e, tmp17306, tmp17331)

tmp17333 := GCons(__e, let__16672, tmp17332)

return GCons(__e, symlet, tmp17333)


panic("unreachable")
})

tmp17334 := Call2(__e, ns2_1set, symshen_4non_1terminalcode, tmp17301)


_ = tmp17334

tmp17335 := MakeNative5(func(__e *ControlFlow, V321, V322, V323, V324, V325 Obj) Obj {
tmp17336 := Call1(__e, SymFn(symgensym), symRemainder)


let__16676 := tmp17336
_ = let__16676

tmp17337 := GCons(__e, V322, Nil)

tmp17338 := GCons(__e, symcons_2, tmp17337)

tmp17339 := GCons(__e, V322, Nil)

tmp17340 := GCons(__e, symtail, tmp17339)

tmp17341 := Call4(__e, SymFn(symshen_4yacc_1syntax), V321, let__16676, V324, V325)


tmp17342 := GCons(__e, tmp17341, Nil)

tmp17343 := GCons(__e, tmp17340, tmp17342)

tmp17344 := GCons(__e, let__16676, tmp17343)

tmp17345 := GCons(__e, symlet, tmp17344)

let__16677 := tmp17345
_ = let__16677

tmp17353 := Call2(__e, SymFn(symshen_4occurs_1check_2), V323, V325)


var ifres17346 Obj

if True == tmp17353 {
tmp17347 := GCons(__e, V322, Nil)

tmp17348 := GCons(__e, symhead, tmp17347)

tmp17349 := GCons(__e, let__16677, Nil)

tmp17350 := GCons(__e, tmp17348, tmp17349)

tmp17351 := GCons(__e, V323, tmp17350)

tmp17352 := GCons(__e, symlet, tmp17351)

ifres17346 = tmp17352


} else {
ifres17346 = let__16677


}

tmp17354 := GCons(__e, symshen_4parse_1failure, Nil)

tmp17355 := GCons(__e, tmp17354, Nil)

tmp17356 := GCons(__e, ifres17346, tmp17355)

tmp17357 := GCons(__e, tmp17338, tmp17356)

return GCons(__e, symif, tmp17357)


panic("unreachable")
})

tmp17358 := Call2(__e, ns2_1set, symshen_4variablecode, tmp17335)


_ = tmp17358

tmp17359 := MakeNative5(func(__e *ControlFlow, V328, V329, V330, V331, V332 Obj) Obj {
tmp17360 := Call1(__e, SymFn(symgensym), symRemainder)


let__16678 := tmp17360
_ = let__16678

tmp17361 := GCons(__e, V329, Nil)

tmp17362 := GCons(__e, symcons_2, tmp17361)

tmp17363 := GCons(__e, V329, Nil)

tmp17364 := GCons(__e, symtail, tmp17363)

tmp17365 := Call4(__e, SymFn(symshen_4yacc_1syntax), V328, let__16678, V331, V332)


tmp17366 := GCons(__e, tmp17365, Nil)

tmp17367 := GCons(__e, tmp17364, tmp17366)

tmp17368 := GCons(__e, let__16678, tmp17367)

tmp17369 := GCons(__e, symlet, tmp17368)

tmp17370 := GCons(__e, symshen_4parse_1failure, Nil)

tmp17371 := GCons(__e, tmp17370, Nil)

tmp17372 := GCons(__e, tmp17369, tmp17371)

tmp17373 := GCons(__e, tmp17362, tmp17372)

return GCons(__e, symif, tmp17373)


panic("unreachable")
})

tmp17374 := Call2(__e, ns2_1set, symshen_4wildcardcode, tmp17359)


_ = tmp17374

tmp17375 := MakeNative5(func(__e *ControlFlow, V334, V335, V336, V337, V338 Obj) Obj {
tmp17376 := Call1(__e, SymFn(symgensym), symRemainder)


let__16679 := tmp17376
_ = let__16679

tmp17377 := GCons(__e, V336, Nil)

tmp17378 := GCons(__e, V335, tmp17377)

tmp17379 := GCons(__e, symshen_4hds_a_2, tmp17378)

tmp17380 := GCons(__e, V335, Nil)

tmp17381 := GCons(__e, symtail, tmp17380)

tmp17382 := Call4(__e, SymFn(symshen_4yacc_1syntax), V334, let__16679, V337, V338)


tmp17383 := GCons(__e, tmp17382, Nil)

tmp17384 := GCons(__e, tmp17381, tmp17383)

tmp17385 := GCons(__e, let__16679, tmp17384)

tmp17386 := GCons(__e, symlet, tmp17385)

tmp17387 := GCons(__e, symshen_4parse_1failure, Nil)

tmp17388 := GCons(__e, tmp17387, Nil)

tmp17389 := GCons(__e, tmp17386, tmp17388)

tmp17390 := GCons(__e, tmp17379, tmp17389)

return GCons(__e, symif, tmp17390)


panic("unreachable")
})

tmp17391 := Call2(__e, ns2_1set, symshen_4terminalcode, tmp17375)


_ = tmp17391

tmp17392 := MakeNative2(func(__e *ControlFlow, V347, V348 Obj) Obj {
tmp17397 := GConsP(__e, V347)

var ifres17394 Obj

if True == tmp17397 {
tmp17395 := GHd(__e, V347)

tmp17396 := GEq(__e, tmp17395, V348)

ifres17394 = tmp17396


} else {
ifres17394 = False


}

if True == ifres17394 {
return True
} else {
return False
}


panic("unreachable")
})

tmp17398 := Call2(__e, ns2_1set, symshen_4hds_a_2, tmp17392)


_ = tmp17398

tmp17399 := MakeNative5(func(__e *ControlFlow, V349, V350, V351, V352, V353 Obj) Obj {
tmp17400 := Call1(__e, SymFn(symgensym), symRemainder)


let__16680 := tmp17400
_ = let__16680

tmp17401 := Call1(__e, SymFn(symgensym), symHd)


let__16681 := tmp17401
_ = let__16681

tmp17402 := Call1(__e, SymFn(symgensym), symTl)


let__16682 := tmp17402
_ = let__16682

tmp17403 := GCons(__e, V350, Nil)

tmp17404 := GCons(__e, symshen_4ccons_2, tmp17403)

tmp17405 := GCons(__e, V350, Nil)

tmp17406 := GCons(__e, symhead, tmp17405)

tmp17407 := GCons(__e, V350, Nil)

tmp17408 := GCons(__e, symtail, tmp17407)

tmp17409 := Call1(__e, SymFn(symshen_4decons), V351)


tmp17410 := GCons(__e, sym_5end_6, Nil)

tmp17411 := Call2(__e, SymFn(symappend), tmp17409, tmp17410)


tmp17412 := Call4(__e, SymFn(symshen_4yacc_1syntax), V349, let__16682, V352, V353)


tmp17413 := GCons(__e, tmp17412, Nil)

tmp17414 := GCons(__e, symshen_4processed, tmp17413)

tmp17415 := Call4(__e, SymFn(symshen_4yacc_1syntax), V349, let__16681, tmp17411, tmp17414)


tmp17416 := GCons(__e, tmp17415, Nil)

tmp17417 := GCons(__e, tmp17408, tmp17416)

tmp17418 := GCons(__e, let__16682, tmp17417)

tmp17419 := GCons(__e, tmp17406, tmp17418)

tmp17420 := GCons(__e, let__16681, tmp17419)

tmp17421 := GCons(__e, symlet, tmp17420)

tmp17422 := GCons(__e, symshen_4parse_1failure, Nil)

tmp17423 := GCons(__e, tmp17422, Nil)

tmp17424 := GCons(__e, tmp17421, tmp17423)

tmp17425 := GCons(__e, tmp17404, tmp17424)

return GCons(__e, symif, tmp17425)


panic("unreachable")
})

tmp17426 := Call2(__e, ns2_1set, symshen_4conscode, tmp17399)


_ = tmp17426

tmp17427 := MakeNative1(func(__e *ControlFlow, V367 Obj) Obj {
tmp17437 := GConsP(__e, V367)

var ifres17434 Obj

if True == tmp17437 {
tmp17435 := GHd(__e, V367)

tmp17436 := GConsP(__e, tmp17435)

ifres17434 = tmp17436


} else {
ifres17434 = False


}

if True == ifres17434 {
return True
} else {
tmp17432 := GConsP(__e, V367)

var ifres17429 Obj

if True == tmp17432 {
tmp17430 := GHd(__e, V367)

tmp17431 := GEq(__e, Nil, tmp17430)

ifres17429 = tmp17431


} else {
ifres17429 = False


}

if True == ifres17429 {
return True
} else {
return False
}


}


panic("unreachable")
})

tmp17438 := Call2(__e, ns2_1set, symshen_4ccons_2, tmp17427)


_ = tmp17438

tmp17439 := MakeNative1(func(__e *ControlFlow, V368 Obj) Obj {
tmp17462 := GConsP(__e, V368)

var ifres17447 Obj

if True == tmp17462 {
tmp17460 := GHd(__e, V368)

tmp17461 := GEq(__e, symcons, tmp17460)

var ifres17448 Obj

if True == tmp17461 {
tmp17458 := GTl(__e, V368)

tmp17459 := GConsP(__e, tmp17458)

var ifres17449 Obj

if True == tmp17459 {
tmp17455 := GTl(__e, V368)

tmp17456 := GTl(__e, tmp17455)

tmp17457 := GConsP(__e, tmp17456)

var ifres17450 Obj

if True == tmp17457 {
tmp17451 := GTl(__e, V368)

tmp17452 := GTl(__e, tmp17451)

tmp17453 := GTl(__e, tmp17452)

tmp17454 := GEq(__e, Nil, tmp17453)

ifres17450 = tmp17454


} else {
ifres17450 = False


}

ifres17449 = ifres17450


} else {
ifres17449 = False


}

ifres17448 = ifres17449


} else {
ifres17448 = False


}

ifres17447 = ifres17448


} else {
ifres17447 = False


}

if True == ifres17447 {
tmp17440 := GTl(__e, V368)

tmp17441 := GHd(__e, tmp17440)

tmp17442 := GTl(__e, V368)

tmp17443 := GTl(__e, tmp17442)

tmp17444 := GHd(__e, tmp17443)

tmp17445 := Call1(__e, SymFn(symshen_4decons), tmp17444)


return GCons(__e, tmp17441, tmp17445)


} else {
return V368
}


panic("unreachable")
})

tmp17463 := Call2(__e, ns2_1set, symshen_4decons, tmp17439)


_ = tmp17463

tmp17464 := MakeNative2(func(__e *ControlFlow, V369, V370 Obj) Obj {
tmp17465 := GCons(__e, V370, Nil)

return GCons(__e, V369, tmp17465)


panic("unreachable")
})

tmp17466 := Call2(__e, ns2_1set, symshen_4comb, tmp17464)


_ = tmp17466

tmp17467 := MakeNative3(func(__e *ControlFlow, V375, V376, V377 Obj) Obj {
tmp17484 := GConsP(__e, V377)

var ifres17474 Obj

if True == tmp17484 {
tmp17482 := GHd(__e, V377)

tmp17483 := GEq(__e, symshen_4processed, tmp17482)

var ifres17475 Obj

if True == tmp17483 {
tmp17480 := GTl(__e, V377)

tmp17481 := GConsP(__e, tmp17480)

var ifres17476 Obj

if True == tmp17481 {
tmp17477 := GTl(__e, V377)

tmp17478 := GTl(__e, tmp17477)

tmp17479 := GEq(__e, Nil, tmp17478)

ifres17476 = tmp17479


} else {
ifres17476 = False


}

ifres17475 = ifres17476


} else {
ifres17475 = False


}

ifres17474 = ifres17475


} else {
ifres17474 = False


}

if True == ifres17474 {
tmp17468 := GTl(__e, V377)

return GHd(__e, tmp17468)


} else {
tmp17469 := Call1(__e, SymFn(symshen_4process_1yacc_1semantics), V377)


let__16683 := tmp17469
_ = let__16683

tmp17470 := Call2(__e, SymFn(symshen_4use_1type_1info), V375, let__16683)


let__16684 := tmp17470
_ = let__16684

tmp17471 := GCons(__e, let__16684, Nil)

tmp17472 := GCons(__e, V376, tmp17471)

return GCons(__e, symshen_4comb, tmp17472)


}


panic("unreachable")
})

tmp17485 := Call2(__e, ns2_1set, symshen_4yacc_1semantics, tmp17467)


_ = tmp17485

tmp17486 := MakeNative2(func(__e *ControlFlow, V383, V384 Obj) Obj {
tmp17651 := GConsP(__e, V383)

var ifres17495 Obj

if True == tmp17651 {
tmp17649 := GHd(__e, V383)

tmp17650 := GEq(__e, sym_i, tmp17649)

var ifres17496 Obj

if True == tmp17650 {
tmp17647 := GTl(__e, V383)

tmp17648 := GConsP(__e, tmp17647)

var ifres17497 Obj

if True == tmp17648 {
tmp17644 := GTl(__e, V383)

tmp17645 := GHd(__e, tmp17644)

tmp17646 := GConsP(__e, tmp17645)

var ifres17498 Obj

if True == tmp17646 {
tmp17640 := GTl(__e, V383)

tmp17641 := GHd(__e, tmp17640)

tmp17642 := GHd(__e, tmp17641)

tmp17643 := GEq(__e, symlist, tmp17642)

var ifres17499 Obj

if True == tmp17643 {
tmp17636 := GTl(__e, V383)

tmp17637 := GHd(__e, tmp17636)

tmp17638 := GTl(__e, tmp17637)

tmp17639 := GConsP(__e, tmp17638)

var ifres17500 Obj

if True == tmp17639 {
tmp17631 := GTl(__e, V383)

tmp17632 := GHd(__e, tmp17631)

tmp17633 := GTl(__e, tmp17632)

tmp17634 := GTl(__e, tmp17633)

tmp17635 := GEq(__e, Nil, tmp17634)

var ifres17501 Obj

if True == tmp17635 {
tmp17628 := GTl(__e, V383)

tmp17629 := GTl(__e, tmp17628)

tmp17630 := GConsP(__e, tmp17629)

var ifres17502 Obj

if True == tmp17630 {
tmp17624 := GTl(__e, V383)

tmp17625 := GTl(__e, tmp17624)

tmp17626 := GHd(__e, tmp17625)

tmp17627 := GEq(__e, sym_1_1_6, tmp17626)

var ifres17503 Obj

if True == tmp17627 {
tmp17620 := GTl(__e, V383)

tmp17621 := GTl(__e, tmp17620)

tmp17622 := GTl(__e, tmp17621)

tmp17623 := GConsP(__e, tmp17622)

var ifres17504 Obj

if True == tmp17623 {
tmp17615 := GTl(__e, V383)

tmp17616 := GTl(__e, tmp17615)

tmp17617 := GTl(__e, tmp17616)

tmp17618 := GHd(__e, tmp17617)

tmp17619 := GConsP(__e, tmp17618)

var ifres17505 Obj

if True == tmp17619 {
tmp17609 := GTl(__e, V383)

tmp17610 := GTl(__e, tmp17609)

tmp17611 := GTl(__e, tmp17610)

tmp17612 := GHd(__e, tmp17611)

tmp17613 := GHd(__e, tmp17612)

tmp17614 := GEq(__e, symstr, tmp17613)

var ifres17506 Obj

if True == tmp17614 {
tmp17603 := GTl(__e, V383)

tmp17604 := GTl(__e, tmp17603)

tmp17605 := GTl(__e, tmp17604)

tmp17606 := GHd(__e, tmp17605)

tmp17607 := GTl(__e, tmp17606)

tmp17608 := GConsP(__e, tmp17607)

var ifres17507 Obj

if True == tmp17608 {
tmp17596 := GTl(__e, V383)

tmp17597 := GTl(__e, tmp17596)

tmp17598 := GTl(__e, tmp17597)

tmp17599 := GHd(__e, tmp17598)

tmp17600 := GTl(__e, tmp17599)

tmp17601 := GHd(__e, tmp17600)

tmp17602 := GConsP(__e, tmp17601)

var ifres17508 Obj

if True == tmp17602 {
tmp17588 := GTl(__e, V383)

tmp17589 := GTl(__e, tmp17588)

tmp17590 := GTl(__e, tmp17589)

tmp17591 := GHd(__e, tmp17590)

tmp17592 := GTl(__e, tmp17591)

tmp17593 := GHd(__e, tmp17592)

tmp17594 := GHd(__e, tmp17593)

tmp17595 := GEq(__e, symlist, tmp17594)

var ifres17509 Obj

if True == tmp17595 {
tmp17580 := GTl(__e, V383)

tmp17581 := GTl(__e, tmp17580)

tmp17582 := GTl(__e, tmp17581)

tmp17583 := GHd(__e, tmp17582)

tmp17584 := GTl(__e, tmp17583)

tmp17585 := GHd(__e, tmp17584)

tmp17586 := GTl(__e, tmp17585)

tmp17587 := GConsP(__e, tmp17586)

var ifres17510 Obj

if True == tmp17587 {
tmp17571 := GTl(__e, V383)

tmp17572 := GTl(__e, tmp17571)

tmp17573 := GTl(__e, tmp17572)

tmp17574 := GHd(__e, tmp17573)

tmp17575 := GTl(__e, tmp17574)

tmp17576 := GHd(__e, tmp17575)

tmp17577 := GTl(__e, tmp17576)

tmp17578 := GTl(__e, tmp17577)

tmp17579 := GEq(__e, Nil, tmp17578)

var ifres17511 Obj

if True == tmp17579 {
tmp17564 := GTl(__e, V383)

tmp17565 := GTl(__e, tmp17564)

tmp17566 := GTl(__e, tmp17565)

tmp17567 := GHd(__e, tmp17566)

tmp17568 := GTl(__e, tmp17567)

tmp17569 := GTl(__e, tmp17568)

tmp17570 := GConsP(__e, tmp17569)

var ifres17512 Obj

if True == tmp17570 {
tmp17556 := GTl(__e, V383)

tmp17557 := GTl(__e, tmp17556)

tmp17558 := GTl(__e, tmp17557)

tmp17559 := GHd(__e, tmp17558)

tmp17560 := GTl(__e, tmp17559)

tmp17561 := GTl(__e, tmp17560)

tmp17562 := GTl(__e, tmp17561)

tmp17563 := GEq(__e, Nil, tmp17562)

var ifres17513 Obj

if True == tmp17563 {
tmp17551 := GTl(__e, V383)

tmp17552 := GTl(__e, tmp17551)

tmp17553 := GTl(__e, tmp17552)

tmp17554 := GTl(__e, tmp17553)

tmp17555 := GConsP(__e, tmp17554)

var ifres17514 Obj

if True == tmp17555 {
tmp17545 := GTl(__e, V383)

tmp17546 := GTl(__e, tmp17545)

tmp17547 := GTl(__e, tmp17546)

tmp17548 := GTl(__e, tmp17547)

tmp17549 := GHd(__e, tmp17548)

tmp17550 := GEq(__e, sym_j, tmp17549)

var ifres17515 Obj

if True == tmp17550 {
tmp17539 := GTl(__e, V383)

tmp17540 := GTl(__e, tmp17539)

tmp17541 := GTl(__e, tmp17540)

tmp17542 := GTl(__e, tmp17541)

tmp17543 := GTl(__e, tmp17542)

tmp17544 := GEq(__e, Nil, tmp17543)

var ifres17516 Obj

if True == tmp17544 {
tmp17526 := GTl(__e, V383)

tmp17527 := GHd(__e, tmp17526)

tmp17528 := GTl(__e, tmp17527)

tmp17529 := GHd(__e, tmp17528)

tmp17530 := GTl(__e, V383)

tmp17531 := GTl(__e, tmp17530)

tmp17532 := GTl(__e, tmp17531)

tmp17533 := GHd(__e, tmp17532)

tmp17534 := GTl(__e, tmp17533)

tmp17535 := GHd(__e, tmp17534)

tmp17536 := GTl(__e, tmp17535)

tmp17537 := GHd(__e, tmp17536)

tmp17538 := GEq(__e, tmp17529, tmp17537)

var ifres17517 Obj

if True == tmp17538 {
tmp17518 := GTl(__e, V383)

tmp17519 := GTl(__e, tmp17518)

tmp17520 := GTl(__e, tmp17519)

tmp17521 := GHd(__e, tmp17520)

tmp17522 := GTl(__e, tmp17521)

tmp17523 := GTl(__e, tmp17522)

tmp17524 := GHd(__e, tmp17523)

tmp17525 := Call1(__e, SymFn(symshen_4monomorphic_2), tmp17524)


ifres17517 = tmp17525


} else {
ifres17517 = False


}

ifres17516 = ifres17517


} else {
ifres17516 = False


}

ifres17515 = ifres17516


} else {
ifres17515 = False


}

ifres17514 = ifres17515


} else {
ifres17514 = False


}

ifres17513 = ifres17514


} else {
ifres17513 = False


}

ifres17512 = ifres17513


} else {
ifres17512 = False


}

ifres17511 = ifres17512


} else {
ifres17511 = False


}

ifres17510 = ifres17511


} else {
ifres17510 = False


}

ifres17509 = ifres17510


} else {
ifres17509 = False


}

ifres17508 = ifres17509


} else {
ifres17508 = False


}

ifres17507 = ifres17508


} else {
ifres17507 = False


}

ifres17506 = ifres17507


} else {
ifres17506 = False


}

ifres17505 = ifres17506


} else {
ifres17505 = False


}

ifres17504 = ifres17505


} else {
ifres17504 = False


}

ifres17503 = ifres17504


} else {
ifres17503 = False


}

ifres17502 = ifres17503


} else {
ifres17502 = False


}

ifres17501 = ifres17502


} else {
ifres17501 = False


}

ifres17500 = ifres17501


} else {
ifres17500 = False


}

ifres17499 = ifres17500


} else {
ifres17499 = False


}

ifres17498 = ifres17499


} else {
ifres17498 = False


}

ifres17497 = ifres17498


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

if True == ifres17495 {
tmp17487 := GTl(__e, V383)

tmp17488 := GTl(__e, tmp17487)

tmp17489 := GTl(__e, tmp17488)

tmp17490 := GHd(__e, tmp17489)

tmp17491 := GTl(__e, tmp17490)

tmp17492 := GTl(__e, tmp17491)

tmp17493 := GCons(__e, V384, tmp17492)

return GCons(__e, symtype, tmp17493)


} else {
return V384
}


panic("unreachable")
})

tmp17652 := Call2(__e, ns2_1set, symshen_4use_1type_1info, tmp17486)


_ = tmp17652

tmp17653 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V387 := __self1
_ = V387
tmp17661 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V387)

if True == tmp17661 {
return False
} else {
tmp17659 := GConsP(__e, V387)

if True == tmp17659 {
tmp17656 := GHd(__e, V387)

tmp17657 := Call1(__e, SymFn(symshen_4monomorphic_2), tmp17656)


if True == tmp17657 {
tmp17654 := GTl(__e, V387)

if SymFn(symshen_4monomorphic_2) == __self {
__self1 = tmp17654
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4monomorphic_2), tmp17654)


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

tmp17662 := Call2(__e, ns2_1set, symshen_4monomorphic_2, tmp17653)


_ = tmp17662

tmp17663 := MakeNative1(func(__e *ControlFlow, V388 Obj) Obj {
tmp17685 := GConsP(__e, V388)

var ifres17671 Obj

if True == tmp17685 {
tmp17683 := GHd(__e, V388)

tmp17684 := GEq(__e, symprotect, tmp17683)

var ifres17672 Obj

if True == tmp17684 {
tmp17681 := GTl(__e, V388)

tmp17682 := GConsP(__e, tmp17681)

var ifres17673 Obj

if True == tmp17682 {
tmp17678 := GTl(__e, V388)

tmp17679 := GTl(__e, tmp17678)

tmp17680 := GEq(__e, Nil, tmp17679)

var ifres17674 Obj

if True == tmp17680 {
tmp17675 := GTl(__e, V388)

tmp17676 := GHd(__e, tmp17675)

tmp17677 := Call1(__e, SymFn(symshen_4non_1terminal_2), tmp17676)


ifres17674 = tmp17677


} else {
ifres17674 = False


}

ifres17673 = ifres17674


} else {
ifres17673 = False


}

ifres17672 = ifres17673


} else {
ifres17672 = False


}

ifres17671 = ifres17672


} else {
ifres17671 = False


}

if True == ifres17671 {
tmp17664 := GTl(__e, V388)

return GHd(__e, tmp17664)


} else {
tmp17669 := GConsP(__e, V388)

if True == tmp17669 {
tmp17665 := MakeNative1(func(__e *ControlFlow, Z389 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4process_1yacc_1semantics), Z389)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp17665, V388)


} else {
tmp17667 := Call1(__e, SymFn(symshen_4non_1terminal_2), V388)


if True == tmp17667 {
return TailCall2(__e, SymFn(symconcat), symAction, V388)
} else {
return V388
}


}


}


panic("unreachable")
})

tmp17686 := Call2(__e, ns2_1set, symshen_4process_1yacc_1semantics, tmp17663)


_ = tmp17686

tmp17687 := MakeNative1(func(__e *ControlFlow, V392 Obj) Obj {
tmp17698 := GConsP(__e, V392)

var ifres17691 Obj

if True == tmp17698 {
tmp17696 := GTl(__e, V392)

tmp17697 := GConsP(__e, tmp17696)

var ifres17692 Obj

if True == tmp17697 {
tmp17693 := GTl(__e, V392)

tmp17694 := GTl(__e, tmp17693)

tmp17695 := GEq(__e, Nil, tmp17694)

ifres17692 = tmp17695


} else {
ifres17692 = False


}

ifres17691 = ifres17692


} else {
ifres17691 = False


}

if True == ifres17691 {
tmp17688 := GTl(__e, V392)

return GHd(__e, tmp17688)


} else {
tmp17689 := GTl(__e, V392)

return GHd(__e, tmp17689)


}


panic("unreachable")
})

tmp17699 := Call2(__e, ns2_1set, symshen_4_5_1out, tmp17687)


_ = tmp17699

tmp17700 := MakeNative1(func(__e *ControlFlow, V393 Obj) Obj {
return GHd(__e, V393)
panic("unreachable")
})

tmp17701 := Call2(__e, ns2_1set, symshen_4in_1_6, tmp17700)


_ = tmp17701

tmp17702 := MakeNative1(func(__e *ControlFlow, V394 Obj) Obj {
tmp17703 := GCons(__e, V394, Nil)

return GCons(__e, Nil, tmp17703)


panic("unreachable")
})

tmp17704 := Call2(__e, ns2_1set, sym_5_b_6, tmp17702)


_ = tmp17704

tmp17705 := MakeNative1(func(__e *ControlFlow, V395 Obj) Obj {
tmp17706 := GCons(__e, Nil, Nil)

return GCons(__e, V395, tmp17706)


panic("unreachable")
})

tmp17707 := Call2(__e, ns2_1set, sym_5e_6, tmp17705)


_ = tmp17707

tmp17708 := MakeNative1(func(__e *ControlFlow, V398 Obj) Obj {
tmp17711 := GEq(__e, Nil, V398)

if True == tmp17711 {
tmp17709 := GCons(__e, Nil, Nil)

return GCons(__e, Nil, tmp17709)


} else {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
}


panic("unreachable")
})

return TailCall2(__e, ns2_1set, sym_5end_6, tmp17708)




panic("unreachable")
})

