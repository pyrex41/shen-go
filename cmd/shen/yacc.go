package main

import . "github.com/tiancaiamao/shen-go/kl"

var YaccMain = MakeNative(func(__e *ControlFlow) {
tmp15660 := MakeNative(func(__e *ControlFlow) {
V112 := __e.Get(1)
_ = V112
V113 := __e.Get(2)
_ = V113
tmp15661 := MakeNative(func(__e *ControlFlow) {
W114 := __e.Get(1)
_ = W114
tmp15668 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W114)


if True == tmp15668 {
__e.Return(PrimSimpleError(MakeString("parse failure\n")))
return
} else {
tmp15666 := Call(__e, PrimFunc(symshen_4partial_1parse_1failure_2), W114)


if True == tmp15666 {
tmp15662 := Call(__e, PrimFunc(symshen_4in_1_6), W114)


tmp15663 := PrimSet(symshen_4_dresidue_d, tmp15662)

_ = tmp15663

tmp15664 := PrimValue(symshen_4_dresidue_d)

__e.TailApply(PrimFunc(symshen_4raise_1syntax_1error), tmp15664)
return


} else {
__e.TailApply(PrimFunc(symshen_4_5_1out), W114)
return
}


}


}, 1)

tmp15669 := Call(__e, V112, V113)


__e.TailApply(tmp15661, tmp15669)
return


}, 2)

tmp15670 := Call(__e, ns2_1set, symcompile, tmp15660)


_ = tmp15670

tmp15671 := MakeNative(func(__e *ControlFlow) {
V115 := __e.Get(1)
_ = V115
tmp15672 := PrimValue(sym_dmaximum_1print_1sequence_1size_d)

tmp15673 := Call(__e, PrimFunc(symshen_4syntax_1error_1message), tmp15672, MakeNumber(0), V115)


tmp15674 := PrimStringConcat(MakeString("syntax error here: "), tmp15673)

tmp15675 := Call(__e, PrimFunc(symshen_4proc_1nl), tmp15674)


__e.Return(PrimSimpleError(tmp15675))
return


}, 1)

tmp15676 := Call(__e, ns2_1set, symshen_4raise_1syntax_1error, tmp15671)


_ = tmp15676

tmp15677 := MakeNative(func(__e *ControlFlow) {
V123 := __e.Get(1)
_ = V123
V124 := __e.Get(2)
_ = V124
V125 := __e.Get(3)
_ = V125
tmp15688 := PrimEqual(Nil, V125)

if True == tmp15688 {
__e.Return(MakeString("\n"))
return
} else {
tmp15686 := PrimEqual(V123, V124)

if True == tmp15686 {
__e.Return(MakeString("...etc \n"))
return
} else {
tmp15684 := PrimIsPair(V125)

if True == tmp15684 {
tmp15678 := PrimHead(V125)

tmp15679 := Call(__e, PrimFunc(symshen_4app), tmp15678, MakeString(" "), symshen_4s)


tmp15680 := PrimNumberAdd(V124, MakeNumber(1))

tmp15681 := PrimTail(V125)

tmp15682 := Call(__e, PrimFunc(symshen_4syntax_1error_1message), V123, tmp15680, tmp15681)


__e.Return(PrimStringConcat(tmp15679, tmp15682))
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4syntax_1error_1message)
return
}


}


}


}, 3)

tmp15689 := Call(__e, ns2_1set, symshen_4syntax_1error_1message, tmp15677)


_ = tmp15689

tmp15690 := MakeNative(func(__e *ControlFlow) {
V126 := __e.Get(1)
_ = V126
tmp15691 := Call(__e, PrimFunc(symfail))


__e.Return(PrimEqual(V126, tmp15691))
return


}, 1)

tmp15692 := Call(__e, ns2_1set, symshen_4parse_1failure_2, tmp15690)


_ = tmp15692

tmp15693 := MakeNative(func(__e *ControlFlow) {
V127 := __e.Get(1)
_ = V127
tmp15694 := Call(__e, PrimFunc(symshen_4in_1_6), V127)


__e.Return(PrimIsPair(tmp15694))
return


}, 1)

tmp15695 := Call(__e, ns2_1set, symshen_4partial_1parse_1failure_2, tmp15693)


_ = tmp15695

tmp15696 := MakeNative(func(__e *ControlFlow) {
V130 := __e.Get(1)
_ = V130
tmp15709 := PrimIsPair(V130)

var ifres15700 Obj

if True == tmp15709 {
tmp15707 := PrimTail(V130)

tmp15708 := PrimIsPair(tmp15707)

var ifres15702 Obj

if True == tmp15708 {
tmp15704 := PrimTail(V130)

tmp15705 := PrimTail(tmp15704)

tmp15706 := PrimEqual(Nil, tmp15705)

var ifres15703 Obj

if True == tmp15706 {
ifres15703 = True


} else {
ifres15703 = False


}

ifres15702 = ifres15703


} else {
ifres15702 = False


}

var ifres15701 Obj

if True == ifres15702 {
ifres15701 = True


} else {
ifres15701 = False


}

ifres15700 = ifres15701


} else {
ifres15700 = False


}

if True == ifres15700 {
tmp15697 := PrimTail(V130)

__e.Return(PrimHead(tmp15697))
return


} else {
tmp15698 := Call(__e, PrimFunc(symshen_4app), V130, MakeString(" is not a YACC stream\n"), symshen_4s)


__e.Return(PrimSimpleError(tmp15698))
return


}


}, 1)

tmp15710 := Call(__e, ns2_1set, symshen_4objectcode, tmp15696)


_ = tmp15710

tmp15711 := MakeNative(func(__e *ControlFlow) {
V131 := __e.Get(1)
_ = V131
tmp15712 := MakeNative(func(__e *ControlFlow) {
Z132 := __e.Get(1)
_ = Z132
__e.TailApply(PrimFunc(symshen_4_5yacc_6), Z132)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp15712, V131)
return


}, 1)

tmp15713 := Call(__e, ns2_1set, symshen_4yacc_1_6shen, tmp15711)


_ = tmp15713

tmp15714 := MakeNative(func(__e *ControlFlow) {
V133 := __e.Get(1)
_ = V133
tmp15715 := MakeNative(func(__e *ControlFlow) {
W134 := __e.Get(1)
_ = W134
tmp15717 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W134)


if True == tmp15717 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W134)
return
}


}, 1)

tmp15753 := PrimIsPair(V133)

var ifres15718 Obj

if True == tmp15753 {
tmp15719 := MakeNative(func(__e *ControlFlow) {
W135 := __e.Get(1)
_ = W135
tmp15720 := MakeNative(func(__e *ControlFlow) {
W136 := __e.Get(1)
_ = W136
tmp15721 := MakeNative(func(__e *ControlFlow) {
W137 := __e.Get(1)
_ = W137
tmp15747 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W137)


if True == tmp15747 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15722 := MakeNative(func(__e *ControlFlow) {
W138 := __e.Get(1)
_ = W138
tmp15723 := MakeNative(func(__e *ControlFlow) {
W139 := __e.Get(1)
_ = W139
tmp15724 := MakeNative(func(__e *ControlFlow) {
W140 := __e.Get(1)
_ = W140
tmp15742 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W140)


if True == tmp15742 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15725 := MakeNative(func(__e *ControlFlow) {
W141 := __e.Get(1)
_ = W141
tmp15726 := MakeNative(func(__e *ControlFlow) {
W142 := __e.Get(1)
_ = W142
tmp15727 := MakeNative(func(__e *ControlFlow) {
W143 := __e.Get(1)
_ = W143
tmp15728 := MakeNative(func(__e *ControlFlow) {
W144 := __e.Get(1)
_ = W144
__e.Return(W144)
return
}, 1)

tmp15729 := PrimCons(W135, Nil)

tmp15730 := PrimCons(symdefine, tmp15729)

tmp15731 := Call(__e, PrimFunc(symshen_4c_1rules_1_6shen), W138, W143, W141)


tmp15732 := PrimCons(tmp15731, Nil)

tmp15733 := PrimCons(sym_1_6, tmp15732)

tmp15734 := PrimCons(W143, tmp15733)

tmp15735 := Call(__e, PrimFunc(symappend), W138, tmp15734)


tmp15736 := Call(__e, PrimFunc(symappend), tmp15730, tmp15735)


__e.TailApply(tmp15728, tmp15736)
return


}, 1)

tmp15737 := Call(__e, PrimFunc(symgensym), symS)


tmp15738 := Call(__e, tmp15727, tmp15737)


__e.TailApply(PrimFunc(symshen_4comb), W142, tmp15738)
return


}, 1)

tmp15739 := Call(__e, PrimFunc(symshen_4in_1_6), W140)


__e.TailApply(tmp15726, tmp15739)
return


}, 1)

tmp15740 := Call(__e, PrimFunc(symshen_4_5_1out), W140)


__e.TailApply(tmp15725, tmp15740)
return


}


}, 1)

tmp15743 := Call(__e, PrimFunc(symshen_4_5c_1rules_6), W139)


__e.TailApply(tmp15724, tmp15743)
return


}, 1)

tmp15744 := Call(__e, PrimFunc(symshen_4in_1_6), W137)


__e.TailApply(tmp15723, tmp15744)
return


}, 1)

tmp15745 := Call(__e, PrimFunc(symshen_4_5_1out), W137)


__e.TailApply(tmp15722, tmp15745)
return


}


}, 1)

tmp15748 := Call(__e, PrimFunc(symshen_4_5yaccsig_6), W136)


__e.TailApply(tmp15721, tmp15748)
return


}, 1)

tmp15749 := Call(__e, PrimFunc(symtail), V133)


__e.TailApply(tmp15720, tmp15749)
return


}, 1)

tmp15750 := Call(__e, PrimFunc(symhead), V133)


tmp15751 := Call(__e, tmp15719, tmp15750)


ifres15718 = tmp15751


} else {
tmp15752 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres15718 = tmp15752


}

__e.TailApply(tmp15715, ifres15718)
return


}, 1)

tmp15754 := Call(__e, ns2_1set, symshen_4_5yacc_6, tmp15714)


_ = tmp15754

tmp15755 := MakeNative(func(__e *ControlFlow) {
V145 := __e.Get(1)
_ = V145
tmp15756 := MakeNative(func(__e *ControlFlow) {
W146 := __e.Get(1)
_ = W146
tmp15768 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W146)


if True == tmp15768 {
tmp15757 := MakeNative(func(__e *ControlFlow) {
W161 := __e.Get(1)
_ = W161
tmp15759 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W161)


if True == tmp15759 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W161)
return
}


}, 1)

tmp15760 := MakeNative(func(__e *ControlFlow) {
W162 := __e.Get(1)
_ = W162
tmp15764 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W162)


if True == tmp15764 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15761 := MakeNative(func(__e *ControlFlow) {
W163 := __e.Get(1)
_ = W163
__e.TailApply(PrimFunc(symshen_4comb), W163, Nil)
return
}, 1)

tmp15762 := Call(__e, PrimFunc(symshen_4in_1_6), W162)


__e.TailApply(tmp15761, tmp15762)
return


}


}, 1)

tmp15765 := Call(__e, PrimFunc(sym_5e_6), V145)


tmp15766 := Call(__e, tmp15760, tmp15765)


__e.TailApply(tmp15757, tmp15766)
return


} else {
__e.Return(W146)
return
}


}, 1)

tmp15831 := PrimIsPair(V145)

var ifres15769 Obj

if True == tmp15831 {
tmp15770 := MakeNative(func(__e *ControlFlow) {
W147 := __e.Get(1)
_ = W147
tmp15771 := MakeNative(func(__e *ControlFlow) {
W148 := __e.Get(1)
_ = W148
tmp15826 := Call(__e, PrimFunc(symshen_4ccons_2), W148)


if True == tmp15826 {
tmp15772 := MakeNative(func(__e *ControlFlow) {
W149 := __e.Get(1)
_ = W149
tmp15773 := MakeNative(func(__e *ControlFlow) {
W150 := __e.Get(1)
_ = W150
tmp15822 := Call(__e, PrimFunc(symshen_4hds_a_2), W149, symlist)


if True == tmp15822 {
tmp15774 := MakeNative(func(__e *ControlFlow) {
W151 := __e.Get(1)
_ = W151
tmp15819 := PrimIsPair(W151)

if True == tmp15819 {
tmp15775 := MakeNative(func(__e *ControlFlow) {
W152 := __e.Get(1)
_ = W152
tmp15776 := MakeNative(func(__e *ControlFlow) {
W153 := __e.Get(1)
_ = W153
tmp15777 := MakeNative(func(__e *ControlFlow) {
W154 := __e.Get(1)
_ = W154
tmp15814 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W154)


if True == tmp15814 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15778 := MakeNative(func(__e *ControlFlow) {
W155 := __e.Get(1)
_ = W155
tmp15811 := Call(__e, PrimFunc(symshen_4hds_a_2), W150, sym_a_a_6)


if True == tmp15811 {
tmp15779 := MakeNative(func(__e *ControlFlow) {
W156 := __e.Get(1)
_ = W156
tmp15808 := PrimIsPair(W156)

if True == tmp15808 {
tmp15780 := MakeNative(func(__e *ControlFlow) {
W157 := __e.Get(1)
_ = W157
tmp15781 := MakeNative(func(__e *ControlFlow) {
W158 := __e.Get(1)
_ = W158
tmp15804 := PrimIsPair(W158)

if True == tmp15804 {
tmp15782 := MakeNative(func(__e *ControlFlow) {
W159 := __e.Get(1)
_ = W159
tmp15783 := MakeNative(func(__e *ControlFlow) {
W160 := __e.Get(1)
_ = W160
tmp15800 := PrimEqual(sym_i, W147)

var ifres15797 Obj

if True == tmp15800 {
tmp15799 := PrimEqual(sym_j, W159)

var ifres15798 Obj

if True == tmp15799 {
ifres15798 = True


} else {
ifres15798 = False


}

ifres15797 = ifres15798


} else {
ifres15797 = False


}

if True == ifres15797 {
tmp15784 := PrimCons(W152, Nil)

tmp15785 := PrimCons(symlist, tmp15784)

tmp15786 := PrimCons(W152, Nil)

tmp15787 := PrimCons(symlist, tmp15786)

tmp15788 := PrimCons(W157, Nil)

tmp15789 := PrimCons(tmp15787, tmp15788)

tmp15790 := PrimCons(symstr, tmp15789)

tmp15791 := PrimCons(sym_j, Nil)

tmp15792 := PrimCons(tmp15790, tmp15791)

tmp15793 := PrimCons(sym_1_1_6, tmp15792)

tmp15794 := PrimCons(tmp15785, tmp15793)

tmp15795 := PrimCons(sym_i, tmp15794)

__e.TailApply(PrimFunc(symshen_4comb), W160, tmp15795)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp15801 := Call(__e, PrimFunc(symtail), W158)


__e.TailApply(tmp15783, tmp15801)
return


}, 1)

tmp15802 := Call(__e, PrimFunc(symhead), W158)


__e.TailApply(tmp15782, tmp15802)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp15805 := Call(__e, PrimFunc(symtail), W156)


__e.TailApply(tmp15781, tmp15805)
return


}, 1)

tmp15806 := Call(__e, PrimFunc(symhead), W156)


__e.TailApply(tmp15780, tmp15806)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp15809 := Call(__e, PrimFunc(symtail), W150)


__e.TailApply(tmp15779, tmp15809)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp15812 := Call(__e, PrimFunc(symshen_4in_1_6), W154)


__e.TailApply(tmp15778, tmp15812)
return


}


}, 1)

tmp15815 := Call(__e, PrimFunc(sym_5end_6), W153)


__e.TailApply(tmp15777, tmp15815)
return


}, 1)

tmp15816 := Call(__e, PrimFunc(symtail), W151)


__e.TailApply(tmp15776, tmp15816)
return


}, 1)

tmp15817 := Call(__e, PrimFunc(symhead), W151)


__e.TailApply(tmp15775, tmp15817)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp15820 := Call(__e, PrimFunc(symtail), W149)


__e.TailApply(tmp15774, tmp15820)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp15823 := Call(__e, PrimFunc(symtail), W148)


__e.TailApply(tmp15773, tmp15823)
return


}, 1)

tmp15824 := Call(__e, PrimFunc(symhead), W148)


__e.TailApply(tmp15772, tmp15824)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp15827 := Call(__e, PrimFunc(symtail), V145)


__e.TailApply(tmp15771, tmp15827)
return


}, 1)

tmp15828 := Call(__e, PrimFunc(symhead), V145)


tmp15829 := Call(__e, tmp15770, tmp15828)


ifres15769 = tmp15829


} else {
tmp15830 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres15769 = tmp15830


}

__e.TailApply(tmp15756, ifres15769)
return


}, 1)

tmp15832 := Call(__e, ns2_1set, symshen_4_5yaccsig_6, tmp15755)


_ = tmp15832

tmp15833 := MakeNative(func(__e *ControlFlow) {
V164 := __e.Get(1)
_ = V164
tmp15834 := MakeNative(func(__e *ControlFlow) {
W165 := __e.Get(1)
_ = W165
tmp15853 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W165)


if True == tmp15853 {
tmp15835 := MakeNative(func(__e *ControlFlow) {
W172 := __e.Get(1)
_ = W172
tmp15837 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W172)


if True == tmp15837 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W172)
return
}


}, 1)

tmp15838 := MakeNative(func(__e *ControlFlow) {
W173 := __e.Get(1)
_ = W173
tmp15849 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W173)


if True == tmp15849 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15839 := MakeNative(func(__e *ControlFlow) {
W174 := __e.Get(1)
_ = W174
tmp15840 := MakeNative(func(__e *ControlFlow) {
W175 := __e.Get(1)
_ = W175
tmp15845 := Call(__e, PrimFunc(symempty_2), W174)


var ifres15841 Obj

if True == tmp15845 {
ifres15841 = Nil


} else {
tmp15842 := Call(__e, PrimFunc(symshen_4app), W174, MakeString("\n ..."), symshen_4r)


tmp15843 := PrimStringConcat(MakeString("YACC syntax error here:\n "), tmp15842)

tmp15844 := PrimSimpleError(tmp15843)

ifres15841 = tmp15844


}

__e.TailApply(PrimFunc(symshen_4comb), W175, ifres15841)
return


}, 1)

tmp15846 := Call(__e, PrimFunc(symshen_4in_1_6), W173)


__e.TailApply(tmp15840, tmp15846)
return


}, 1)

tmp15847 := Call(__e, PrimFunc(symshen_4_5_1out), W173)


__e.TailApply(tmp15839, tmp15847)
return


}


}, 1)

tmp15850 := Call(__e, PrimFunc(sym_5_b_6), V164)


tmp15851 := Call(__e, tmp15838, tmp15850)


__e.TailApply(tmp15835, tmp15851)
return


} else {
__e.Return(W165)
return
}


}, 1)

tmp15854 := MakeNative(func(__e *ControlFlow) {
W166 := __e.Get(1)
_ = W166
tmp15869 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W166)


if True == tmp15869 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15855 := MakeNative(func(__e *ControlFlow) {
W167 := __e.Get(1)
_ = W167
tmp15856 := MakeNative(func(__e *ControlFlow) {
W168 := __e.Get(1)
_ = W168
tmp15857 := MakeNative(func(__e *ControlFlow) {
W169 := __e.Get(1)
_ = W169
tmp15864 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W169)


if True == tmp15864 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15858 := MakeNative(func(__e *ControlFlow) {
W170 := __e.Get(1)
_ = W170
tmp15859 := MakeNative(func(__e *ControlFlow) {
W171 := __e.Get(1)
_ = W171
tmp15860 := PrimCons(W167, W170)

__e.TailApply(PrimFunc(symshen_4comb), W171, tmp15860)
return


}, 1)

tmp15861 := Call(__e, PrimFunc(symshen_4in_1_6), W169)


__e.TailApply(tmp15859, tmp15861)
return


}, 1)

tmp15862 := Call(__e, PrimFunc(symshen_4_5_1out), W169)


__e.TailApply(tmp15858, tmp15862)
return


}


}, 1)

tmp15865 := Call(__e, PrimFunc(symshen_4_5c_1rules_6), W168)


__e.TailApply(tmp15857, tmp15865)
return


}, 1)

tmp15866 := Call(__e, PrimFunc(symshen_4in_1_6), W166)


__e.TailApply(tmp15856, tmp15866)
return


}, 1)

tmp15867 := Call(__e, PrimFunc(symshen_4_5_1out), W166)


__e.TailApply(tmp15855, tmp15867)
return


}


}, 1)

tmp15870 := Call(__e, PrimFunc(symshen_4_5c_1rule_6), V164)


tmp15871 := Call(__e, tmp15854, tmp15870)


__e.TailApply(tmp15834, tmp15871)
return


}, 1)

tmp15872 := Call(__e, ns2_1set, symshen_4_5c_1rules_6, tmp15833)


_ = tmp15872

tmp15873 := MakeNative(func(__e *ControlFlow) {
V176 := __e.Get(1)
_ = V176
tmp15874 := MakeNative(func(__e *ControlFlow) {
W177 := __e.Get(1)
_ = W177
tmp15897 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W177)


if True == tmp15897 {
tmp15875 := MakeNative(func(__e *ControlFlow) {
W186 := __e.Get(1)
_ = W186
tmp15877 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W186)


if True == tmp15877 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W186)
return
}


}, 1)

tmp15878 := MakeNative(func(__e *ControlFlow) {
W187 := __e.Get(1)
_ = W187
tmp15893 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W187)


if True == tmp15893 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15879 := MakeNative(func(__e *ControlFlow) {
W188 := __e.Get(1)
_ = W188
tmp15880 := MakeNative(func(__e *ControlFlow) {
W189 := __e.Get(1)
_ = W189
tmp15881 := MakeNative(func(__e *ControlFlow) {
W190 := __e.Get(1)
_ = W190
tmp15888 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W190)


if True == tmp15888 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15882 := MakeNative(func(__e *ControlFlow) {
W191 := __e.Get(1)
_ = W191
tmp15883 := Call(__e, PrimFunc(symshen_4autocomplete), W188)


tmp15884 := PrimCons(tmp15883, Nil)

tmp15885 := PrimCons(W188, tmp15884)

__e.TailApply(PrimFunc(symshen_4comb), W191, tmp15885)
return


}, 1)

tmp15886 := Call(__e, PrimFunc(symshen_4in_1_6), W190)


__e.TailApply(tmp15882, tmp15886)
return


}


}, 1)

tmp15889 := Call(__e, PrimFunc(symshen_4_5sc_6), W189)


__e.TailApply(tmp15881, tmp15889)
return


}, 1)

tmp15890 := Call(__e, PrimFunc(symshen_4in_1_6), W187)


__e.TailApply(tmp15880, tmp15890)
return


}, 1)

tmp15891 := Call(__e, PrimFunc(symshen_4_5_1out), W187)


__e.TailApply(tmp15879, tmp15891)
return


}


}, 1)

tmp15894 := Call(__e, PrimFunc(symshen_4_5syntax_6), V176)


tmp15895 := Call(__e, tmp15878, tmp15894)


__e.TailApply(tmp15875, tmp15895)
return


} else {
__e.Return(W177)
return
}


}, 1)

tmp15898 := MakeNative(func(__e *ControlFlow) {
W178 := __e.Get(1)
_ = W178
tmp15920 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W178)


if True == tmp15920 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15899 := MakeNative(func(__e *ControlFlow) {
W179 := __e.Get(1)
_ = W179
tmp15900 := MakeNative(func(__e *ControlFlow) {
W180 := __e.Get(1)
_ = W180
tmp15901 := MakeNative(func(__e *ControlFlow) {
W181 := __e.Get(1)
_ = W181
tmp15915 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W181)


if True == tmp15915 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15902 := MakeNative(func(__e *ControlFlow) {
W182 := __e.Get(1)
_ = W182
tmp15903 := MakeNative(func(__e *ControlFlow) {
W183 := __e.Get(1)
_ = W183
tmp15904 := MakeNative(func(__e *ControlFlow) {
W184 := __e.Get(1)
_ = W184
tmp15910 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W184)


if True == tmp15910 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15905 := MakeNative(func(__e *ControlFlow) {
W185 := __e.Get(1)
_ = W185
tmp15906 := PrimCons(W182, Nil)

tmp15907 := PrimCons(W179, tmp15906)

__e.TailApply(PrimFunc(symshen_4comb), W185, tmp15907)
return


}, 1)

tmp15908 := Call(__e, PrimFunc(symshen_4in_1_6), W184)


__e.TailApply(tmp15905, tmp15908)
return


}


}, 1)

tmp15911 := Call(__e, PrimFunc(symshen_4_5sc_6), W183)


__e.TailApply(tmp15904, tmp15911)
return


}, 1)

tmp15912 := Call(__e, PrimFunc(symshen_4in_1_6), W181)


__e.TailApply(tmp15903, tmp15912)
return


}, 1)

tmp15913 := Call(__e, PrimFunc(symshen_4_5_1out), W181)


__e.TailApply(tmp15902, tmp15913)
return


}


}, 1)

tmp15916 := Call(__e, PrimFunc(symshen_4_5semantics_6), W180)


__e.TailApply(tmp15901, tmp15916)
return


}, 1)

tmp15917 := Call(__e, PrimFunc(symshen_4in_1_6), W178)


__e.TailApply(tmp15900, tmp15917)
return


}, 1)

tmp15918 := Call(__e, PrimFunc(symshen_4_5_1out), W178)


__e.TailApply(tmp15899, tmp15918)
return


}


}, 1)

tmp15921 := Call(__e, PrimFunc(symshen_4_5syntax_6), V176)


tmp15922 := Call(__e, tmp15898, tmp15921)


__e.TailApply(tmp15874, tmp15922)
return


}, 1)

tmp15923 := Call(__e, ns2_1set, symshen_4_5c_1rule_6, tmp15873)


_ = tmp15923

tmp15924 := MakeNative(func(__e *ControlFlow) {
V192 := __e.Get(1)
_ = V192
tmp15953 := PrimIsPair(V192)

var ifres15945 Obj

if True == tmp15953 {
tmp15951 := PrimTail(V192)

tmp15952 := PrimEqual(Nil, tmp15951)

var ifres15947 Obj

if True == tmp15952 {
tmp15949 := PrimHead(V192)

tmp15950 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp15949)


var ifres15948 Obj

if True == tmp15950 {
ifres15948 = True


} else {
ifres15948 = False


}

ifres15947 = ifres15948


} else {
ifres15947 = False


}

var ifres15946 Obj

if True == ifres15947 {
ifres15946 = True


} else {
ifres15946 = False


}

ifres15945 = ifres15946


} else {
ifres15945 = False


}

if True == ifres15945 {
__e.Return(PrimHead(V192))
return
} else {
tmp15943 := PrimIsPair(V192)

var ifres15939 Obj

if True == tmp15943 {
tmp15941 := PrimHead(V192)

tmp15942 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp15941)


var ifres15940 Obj

if True == tmp15942 {
ifres15940 = True


} else {
ifres15940 = False


}

ifres15939 = ifres15940


} else {
ifres15939 = False


}

if True == ifres15939 {
tmp15925 := PrimHead(V192)

tmp15926 := PrimTail(V192)

tmp15927 := Call(__e, PrimFunc(symshen_4autocomplete), tmp15926)


tmp15928 := PrimCons(tmp15927, Nil)

tmp15929 := PrimCons(tmp15925, tmp15928)

__e.Return(PrimCons(symappend, tmp15929))
return


} else {
tmp15937 := PrimIsPair(V192)

if True == tmp15937 {
tmp15930 := PrimHead(V192)

tmp15931 := Call(__e, PrimFunc(symshen_4autocomplete), tmp15930)


tmp15932 := PrimTail(V192)

tmp15933 := Call(__e, PrimFunc(symshen_4autocomplete), tmp15932)


tmp15934 := PrimCons(tmp15933, Nil)

tmp15935 := PrimCons(tmp15931, tmp15934)

__e.Return(PrimCons(symcons, tmp15935))
return


} else {
__e.Return(V192)
return
}


}


}


}, 1)

tmp15954 := Call(__e, ns2_1set, symshen_4autocomplete, tmp15924)


_ = tmp15954

tmp15955 := MakeNative(func(__e *ControlFlow) {
V193 := __e.Get(1)
_ = V193
tmp15962 := PrimIsSymbol(V193)

if True == tmp15962 {
tmp15957 := MakeNative(func(__e *ControlFlow) {
W194 := __e.Get(1)
_ = W194
tmp15958 := MakeNative(func(__e *ControlFlow) {
Z195 := __e.Get(1)
_ = Z195
__e.TailApply(PrimFunc(symshen_4_5non_1terminal_2_6), Z195)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp15958, W194)
return


}, 1)

tmp15959 := Call(__e, PrimFunc(symexplode), V193)


tmp15960 := Call(__e, tmp15957, tmp15959)


if True == tmp15960 {
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

tmp15963 := Call(__e, ns2_1set, symshen_4non_1terminal_2, tmp15955)


_ = tmp15963

tmp15964 := MakeNative(func(__e *ControlFlow) {
V196 := __e.Get(1)
_ = V196
tmp15965 := MakeNative(func(__e *ControlFlow) {
W197 := __e.Get(1)
_ = W197
tmp15987 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W197)


if True == tmp15987 {
tmp15966 := MakeNative(func(__e *ControlFlow) {
W202 := __e.Get(1)
_ = W202
tmp15978 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W202)


if True == tmp15978 {
tmp15967 := MakeNative(func(__e *ControlFlow) {
W205 := __e.Get(1)
_ = W205
tmp15969 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W205)


if True == tmp15969 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W205)
return
}


}, 1)

tmp15970 := MakeNative(func(__e *ControlFlow) {
W206 := __e.Get(1)
_ = W206
tmp15974 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W206)


if True == tmp15974 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15971 := MakeNative(func(__e *ControlFlow) {
W207 := __e.Get(1)
_ = W207
__e.TailApply(PrimFunc(symshen_4comb), W207, False)
return
}, 1)

tmp15972 := Call(__e, PrimFunc(symshen_4in_1_6), W206)


__e.TailApply(tmp15971, tmp15972)
return


}


}, 1)

tmp15975 := Call(__e, PrimFunc(sym_5_b_6), V196)


tmp15976 := Call(__e, tmp15970, tmp15975)


__e.TailApply(tmp15967, tmp15976)
return


} else {
__e.Return(W202)
return
}


}, 1)

tmp15979 := MakeNative(func(__e *ControlFlow) {
W203 := __e.Get(1)
_ = W203
tmp15983 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W203)


if True == tmp15983 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15980 := MakeNative(func(__e *ControlFlow) {
W204 := __e.Get(1)
_ = W204
__e.TailApply(PrimFunc(symshen_4comb), W204, True)
return
}, 1)

tmp15981 := Call(__e, PrimFunc(symshen_4in_1_6), W203)


__e.TailApply(tmp15980, tmp15981)
return


}


}, 1)

tmp15984 := Call(__e, PrimFunc(symshen_4_5non_1terminal_1name_6), V196)


tmp15985 := Call(__e, tmp15979, tmp15984)


__e.TailApply(tmp15966, tmp15985)
return


} else {
__e.Return(W197)
return
}


}, 1)

tmp15988 := MakeNative(func(__e *ControlFlow) {
W198 := __e.Get(1)
_ = W198
tmp15998 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W198)


if True == tmp15998 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15989 := MakeNative(func(__e *ControlFlow) {
W199 := __e.Get(1)
_ = W199
tmp15990 := MakeNative(func(__e *ControlFlow) {
W200 := __e.Get(1)
_ = W200
tmp15994 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W200)


if True == tmp15994 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp15991 := MakeNative(func(__e *ControlFlow) {
W201 := __e.Get(1)
_ = W201
__e.TailApply(PrimFunc(symshen_4comb), W201, True)
return
}, 1)

tmp15992 := Call(__e, PrimFunc(symshen_4in_1_6), W200)


__e.TailApply(tmp15991, tmp15992)
return


}


}, 1)

tmp15995 := Call(__e, PrimFunc(symshen_4_5non_1terminal_1name_6), W199)


__e.TailApply(tmp15990, tmp15995)
return


}, 1)

tmp15996 := Call(__e, PrimFunc(symshen_4in_1_6), W198)


__e.TailApply(tmp15989, tmp15996)
return


}


}, 1)

tmp15999 := Call(__e, PrimFunc(symshen_4_5packagenames_6), V196)


tmp16000 := Call(__e, tmp15988, tmp15999)


__e.TailApply(tmp15965, tmp16000)
return


}, 1)

tmp16001 := Call(__e, ns2_1set, symshen_4_5non_1terminal_2_6, tmp15964)


_ = tmp16001

tmp16002 := MakeNative(func(__e *ControlFlow) {
V208 := __e.Get(1)
_ = V208
tmp16003 := MakeNative(func(__e *ControlFlow) {
W209 := __e.Get(1)
_ = W209
tmp16019 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W209)


if True == tmp16019 {
tmp16004 := MakeNative(func(__e *ControlFlow) {
W215 := __e.Get(1)
_ = W215
tmp16006 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W215)


if True == tmp16006 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W215)
return
}


}, 1)

tmp16007 := MakeNative(func(__e *ControlFlow) {
W216 := __e.Get(1)
_ = W216
tmp16015 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W216)


if True == tmp16015 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16008 := MakeNative(func(__e *ControlFlow) {
W217 := __e.Get(1)
_ = W217
tmp16012 := Call(__e, PrimFunc(symshen_4hds_a_2), W217, MakeString("."))


if True == tmp16012 {
tmp16009 := MakeNative(func(__e *ControlFlow) {
W218 := __e.Get(1)
_ = W218
__e.TailApply(PrimFunc(symshen_4comb), W218, symshen_4skip)
return
}, 1)

tmp16010 := Call(__e, PrimFunc(symtail), W217)


__e.TailApply(tmp16009, tmp16010)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16013 := Call(__e, PrimFunc(symshen_4in_1_6), W216)


__e.TailApply(tmp16008, tmp16013)
return


}


}, 1)

tmp16016 := Call(__e, PrimFunc(symshen_4_5packagename_6), V208)


tmp16017 := Call(__e, tmp16007, tmp16016)


__e.TailApply(tmp16004, tmp16017)
return


} else {
__e.Return(W209)
return
}


}, 1)

tmp16020 := MakeNative(func(__e *ControlFlow) {
W210 := __e.Get(1)
_ = W210
tmp16034 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W210)


if True == tmp16034 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16021 := MakeNative(func(__e *ControlFlow) {
W211 := __e.Get(1)
_ = W211
tmp16031 := Call(__e, PrimFunc(symshen_4hds_a_2), W211, MakeString("."))


if True == tmp16031 {
tmp16022 := MakeNative(func(__e *ControlFlow) {
W212 := __e.Get(1)
_ = W212
tmp16023 := MakeNative(func(__e *ControlFlow) {
W213 := __e.Get(1)
_ = W213
tmp16027 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W213)


if True == tmp16027 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16024 := MakeNative(func(__e *ControlFlow) {
W214 := __e.Get(1)
_ = W214
__e.TailApply(PrimFunc(symshen_4comb), W214, symshen_4skip)
return
}, 1)

tmp16025 := Call(__e, PrimFunc(symshen_4in_1_6), W213)


__e.TailApply(tmp16024, tmp16025)
return


}


}, 1)

tmp16028 := Call(__e, PrimFunc(symshen_4_5packagenames_6), W212)


__e.TailApply(tmp16023, tmp16028)
return


}, 1)

tmp16029 := Call(__e, PrimFunc(symtail), W211)


__e.TailApply(tmp16022, tmp16029)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16032 := Call(__e, PrimFunc(symshen_4in_1_6), W210)


__e.TailApply(tmp16021, tmp16032)
return


}


}, 1)

tmp16035 := Call(__e, PrimFunc(symshen_4_5packagename_6), V208)


tmp16036 := Call(__e, tmp16020, tmp16035)


__e.TailApply(tmp16003, tmp16036)
return


}, 1)

tmp16037 := Call(__e, ns2_1set, symshen_4_5packagenames_6, tmp16002)


_ = tmp16037

tmp16038 := MakeNative(func(__e *ControlFlow) {
V219 := __e.Get(1)
_ = V219
tmp16039 := MakeNative(func(__e *ControlFlow) {
W220 := __e.Get(1)
_ = W220
tmp16051 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W220)


if True == tmp16051 {
tmp16040 := MakeNative(func(__e *ControlFlow) {
W225 := __e.Get(1)
_ = W225
tmp16042 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W225)


if True == tmp16042 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W225)
return
}


}, 1)

tmp16043 := MakeNative(func(__e *ControlFlow) {
W226 := __e.Get(1)
_ = W226
tmp16047 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W226)


if True == tmp16047 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16044 := MakeNative(func(__e *ControlFlow) {
W227 := __e.Get(1)
_ = W227
__e.TailApply(PrimFunc(symshen_4comb), W227, symshen_4skip)
return
}, 1)

tmp16045 := Call(__e, PrimFunc(symshen_4in_1_6), W226)


__e.TailApply(tmp16044, tmp16045)
return


}


}, 1)

tmp16048 := Call(__e, PrimFunc(sym_5e_6), V219)


tmp16049 := Call(__e, tmp16043, tmp16048)


__e.TailApply(tmp16040, tmp16049)
return


} else {
__e.Return(W220)
return
}


}, 1)

tmp16052 := MakeNative(func(__e *ControlFlow) {
W221 := __e.Get(1)
_ = W221
tmp16062 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W221)


if True == tmp16062 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16053 := MakeNative(func(__e *ControlFlow) {
W222 := __e.Get(1)
_ = W222
tmp16054 := MakeNative(func(__e *ControlFlow) {
W223 := __e.Get(1)
_ = W223
tmp16058 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W223)


if True == tmp16058 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16055 := MakeNative(func(__e *ControlFlow) {
W224 := __e.Get(1)
_ = W224
__e.TailApply(PrimFunc(symshen_4comb), W224, symshen_4skip)
return
}, 1)

tmp16056 := Call(__e, PrimFunc(symshen_4in_1_6), W223)


__e.TailApply(tmp16055, tmp16056)
return


}


}, 1)

tmp16059 := Call(__e, PrimFunc(symshen_4_5packagename_6), W222)


__e.TailApply(tmp16054, tmp16059)
return


}, 1)

tmp16060 := Call(__e, PrimFunc(symshen_4in_1_6), W221)


__e.TailApply(tmp16053, tmp16060)
return


}


}, 1)

tmp16063 := Call(__e, PrimFunc(symshen_4_5packagechar_6), V219)


tmp16064 := Call(__e, tmp16052, tmp16063)


__e.TailApply(tmp16039, tmp16064)
return


}, 1)

tmp16065 := Call(__e, ns2_1set, symshen_4_5packagename_6, tmp16038)


_ = tmp16065

tmp16066 := MakeNative(func(__e *ControlFlow) {
V228 := __e.Get(1)
_ = V228
tmp16067 := MakeNative(func(__e *ControlFlow) {
W229 := __e.Get(1)
_ = W229
tmp16069 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W229)


if True == tmp16069 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W229)
return
}


}, 1)

tmp16080 := PrimIsPair(V228)

var ifres16070 Obj

if True == tmp16080 {
tmp16071 := MakeNative(func(__e *ControlFlow) {
W230 := __e.Get(1)
_ = W230
tmp16072 := MakeNative(func(__e *ControlFlow) {
W231 := __e.Get(1)
_ = W231
tmp16074 := PrimEqual(W230, MakeString("."))

tmp16075 := PrimNot(tmp16074)

if True == tmp16075 {
__e.TailApply(PrimFunc(symshen_4comb), W231, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16076 := Call(__e, PrimFunc(symtail), V228)


__e.TailApply(tmp16072, tmp16076)
return


}, 1)

tmp16077 := Call(__e, PrimFunc(symhead), V228)


tmp16078 := Call(__e, tmp16071, tmp16077)


ifres16070 = tmp16078


} else {
tmp16079 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16070 = tmp16079


}

__e.TailApply(tmp16067, ifres16070)
return


}, 1)

tmp16081 := Call(__e, ns2_1set, symshen_4_5packagechar_6, tmp16066)


_ = tmp16081

tmp16082 := MakeNative(func(__e *ControlFlow) {
V232 := __e.Get(1)
_ = V232
tmp16083 := MakeNative(func(__e *ControlFlow) {
W233 := __e.Get(1)
_ = W233
tmp16085 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W233)


if True == tmp16085 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W233)
return
}


}, 1)

tmp16108 := Call(__e, PrimFunc(symshen_4hds_a_2), V232, MakeString("<"))


var ifres16086 Obj

if True == tmp16108 {
tmp16087 := MakeNative(func(__e *ControlFlow) {
W234 := __e.Get(1)
_ = W234
tmp16088 := MakeNative(func(__e *ControlFlow) {
W235 := __e.Get(1)
_ = W235
tmp16103 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W235)


if True == tmp16103 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16089 := MakeNative(func(__e *ControlFlow) {
W236 := __e.Get(1)
_ = W236
tmp16090 := MakeNative(func(__e *ControlFlow) {
W237 := __e.Get(1)
_ = W237
tmp16092 := MakeNative(func(__e *ControlFlow) {
W238 := __e.Get(1)
_ = W238
tmp16097 := PrimIsPair(W238)

if True == tmp16097 {
tmp16094 := PrimHead(W238)

tmp16095 := PrimEqual(tmp16094, MakeString(">"))

if True == tmp16095 {
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

tmp16098 := Call(__e, PrimFunc(symreverse), W236)


tmp16099 := Call(__e, tmp16092, tmp16098)


if True == tmp16099 {
__e.TailApply(PrimFunc(symshen_4comb), W237, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16100 := Call(__e, PrimFunc(symshen_4in_1_6), W235)


__e.TailApply(tmp16090, tmp16100)
return


}, 1)

tmp16101 := Call(__e, PrimFunc(symshen_4_5_1out), W235)


__e.TailApply(tmp16089, tmp16101)
return


}


}, 1)

tmp16104 := Call(__e, PrimFunc(sym_5_b_6), W234)


__e.TailApply(tmp16088, tmp16104)
return


}, 1)

tmp16105 := Call(__e, PrimFunc(symtail), V232)


tmp16106 := Call(__e, tmp16087, tmp16105)


ifres16086 = tmp16106


} else {
tmp16107 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16086 = tmp16107


}

__e.TailApply(tmp16083, ifres16086)
return


}, 1)

tmp16109 := Call(__e, ns2_1set, symshen_4_5non_1terminal_1name_6, tmp16082)


_ = tmp16109

tmp16110 := MakeNative(func(__e *ControlFlow) {
V239 := __e.Get(1)
_ = V239
tmp16111 := PrimIntern(MakeString(";"))

__e.Return(PrimEqual(V239, tmp16111))
return


}, 1)

tmp16112 := Call(__e, ns2_1set, symshen_4semicolon_2, tmp16110)


_ = tmp16112

tmp16113 := MakeNative(func(__e *ControlFlow) {
V240 := __e.Get(1)
_ = V240
tmp16114 := MakeNative(func(__e *ControlFlow) {
W241 := __e.Get(1)
_ = W241
tmp16116 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W241)


if True == tmp16116 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W241)
return
}


}, 1)

tmp16126 := PrimIsPair(V240)

var ifres16117 Obj

if True == tmp16126 {
tmp16118 := MakeNative(func(__e *ControlFlow) {
W242 := __e.Get(1)
_ = W242
tmp16119 := MakeNative(func(__e *ControlFlow) {
W243 := __e.Get(1)
_ = W243
tmp16121 := Call(__e, PrimFunc(symshen_4colon_1equal_2), W242)


if True == tmp16121 {
__e.TailApply(PrimFunc(symshen_4comb), W243, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16122 := Call(__e, PrimFunc(symtail), V240)


__e.TailApply(tmp16119, tmp16122)
return


}, 1)

tmp16123 := Call(__e, PrimFunc(symhead), V240)


tmp16124 := Call(__e, tmp16118, tmp16123)


ifres16117 = tmp16124


} else {
tmp16125 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16117 = tmp16125


}

__e.TailApply(tmp16114, ifres16117)
return


}, 1)

tmp16127 := Call(__e, ns2_1set, symshen_4_5colon_1equal_6, tmp16113)


_ = tmp16127

tmp16128 := MakeNative(func(__e *ControlFlow) {
V244 := __e.Get(1)
_ = V244
tmp16129 := PrimIntern(MakeString(":="))

__e.Return(PrimEqual(tmp16129, V244))
return


}, 1)

tmp16130 := Call(__e, ns2_1set, symshen_4colon_1equal_2, tmp16128)


_ = tmp16130

tmp16131 := MakeNative(func(__e *ControlFlow) {
V245 := __e.Get(1)
_ = V245
tmp16132 := MakeNative(func(__e *ControlFlow) {
W246 := __e.Get(1)
_ = W246
tmp16147 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W246)


if True == tmp16147 {
tmp16133 := MakeNative(func(__e *ControlFlow) {
W253 := __e.Get(1)
_ = W253
tmp16135 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W253)


if True == tmp16135 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W253)
return
}


}, 1)

tmp16136 := MakeNative(func(__e *ControlFlow) {
W254 := __e.Get(1)
_ = W254
tmp16143 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W254)


if True == tmp16143 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16137 := MakeNative(func(__e *ControlFlow) {
W255 := __e.Get(1)
_ = W255
tmp16138 := MakeNative(func(__e *ControlFlow) {
W256 := __e.Get(1)
_ = W256
tmp16139 := PrimCons(W255, Nil)

__e.TailApply(PrimFunc(symshen_4comb), W256, tmp16139)
return


}, 1)

tmp16140 := Call(__e, PrimFunc(symshen_4in_1_6), W254)


__e.TailApply(tmp16138, tmp16140)
return


}, 1)

tmp16141 := Call(__e, PrimFunc(symshen_4_5_1out), W254)


__e.TailApply(tmp16137, tmp16141)
return


}


}, 1)

tmp16144 := Call(__e, PrimFunc(symshen_4_5syntax_1item_6), V245)


tmp16145 := Call(__e, tmp16136, tmp16144)


__e.TailApply(tmp16133, tmp16145)
return


} else {
__e.Return(W246)
return
}


}, 1)

tmp16148 := MakeNative(func(__e *ControlFlow) {
W247 := __e.Get(1)
_ = W247
tmp16163 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W247)


if True == tmp16163 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16149 := MakeNative(func(__e *ControlFlow) {
W248 := __e.Get(1)
_ = W248
tmp16150 := MakeNative(func(__e *ControlFlow) {
W249 := __e.Get(1)
_ = W249
tmp16151 := MakeNative(func(__e *ControlFlow) {
W250 := __e.Get(1)
_ = W250
tmp16158 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W250)


if True == tmp16158 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16152 := MakeNative(func(__e *ControlFlow) {
W251 := __e.Get(1)
_ = W251
tmp16153 := MakeNative(func(__e *ControlFlow) {
W252 := __e.Get(1)
_ = W252
tmp16154 := PrimCons(W248, W251)

__e.TailApply(PrimFunc(symshen_4comb), W252, tmp16154)
return


}, 1)

tmp16155 := Call(__e, PrimFunc(symshen_4in_1_6), W250)


__e.TailApply(tmp16153, tmp16155)
return


}, 1)

tmp16156 := Call(__e, PrimFunc(symshen_4_5_1out), W250)


__e.TailApply(tmp16152, tmp16156)
return


}


}, 1)

tmp16159 := Call(__e, PrimFunc(symshen_4_5syntax_6), W249)


__e.TailApply(tmp16151, tmp16159)
return


}, 1)

tmp16160 := Call(__e, PrimFunc(symshen_4in_1_6), W247)


__e.TailApply(tmp16150, tmp16160)
return


}, 1)

tmp16161 := Call(__e, PrimFunc(symshen_4_5_1out), W247)


__e.TailApply(tmp16149, tmp16161)
return


}


}, 1)

tmp16164 := Call(__e, PrimFunc(symshen_4_5syntax_1item_6), V245)


tmp16165 := Call(__e, tmp16148, tmp16164)


__e.TailApply(tmp16132, tmp16165)
return


}, 1)

tmp16166 := Call(__e, ns2_1set, symshen_4_5syntax_6, tmp16131)


_ = tmp16166

tmp16167 := MakeNative(func(__e *ControlFlow) {
V257 := __e.Get(1)
_ = V257
tmp16168 := MakeNative(func(__e *ControlFlow) {
W258 := __e.Get(1)
_ = W258
tmp16170 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W258)


if True == tmp16170 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W258)
return
}


}, 1)

tmp16180 := PrimIsPair(V257)

var ifres16171 Obj

if True == tmp16180 {
tmp16172 := MakeNative(func(__e *ControlFlow) {
W259 := __e.Get(1)
_ = W259
tmp16173 := MakeNative(func(__e *ControlFlow) {
W260 := __e.Get(1)
_ = W260
tmp16175 := Call(__e, PrimFunc(symshen_4syntax_1item_2), W259)


if True == tmp16175 {
__e.TailApply(PrimFunc(symshen_4comb), W260, W259)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16176 := Call(__e, PrimFunc(symtail), V257)


__e.TailApply(tmp16173, tmp16176)
return


}, 1)

tmp16177 := Call(__e, PrimFunc(symhead), V257)


tmp16178 := Call(__e, tmp16172, tmp16177)


ifres16171 = tmp16178


} else {
tmp16179 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres16171 = tmp16179


}

__e.TailApply(tmp16168, ifres16171)
return


}, 1)

tmp16181 := Call(__e, ns2_1set, symshen_4_5syntax_1item_6, tmp16167)


_ = tmp16181

tmp16182 := MakeNative(func(__e *ControlFlow) {
V263 := __e.Get(1)
_ = V263
tmp16218 := Call(__e, PrimFunc(symshen_4colon_1equal_2), V263)


if True == tmp16218 {
__e.Return(False)
return
} else {
tmp16216 := Call(__e, PrimFunc(symshen_4semicolon_2), V263)


if True == tmp16216 {
__e.Return(False)
return
} else {
tmp16214 := Call(__e, PrimFunc(symatom_2), V263)


if True == tmp16214 {
__e.Return(True)
return
} else {
tmp16212 := PrimIsPair(V263)

var ifres16193 Obj

if True == tmp16212 {
tmp16210 := PrimHead(V263)

tmp16211 := PrimEqual(symcons, tmp16210)

var ifres16195 Obj

if True == tmp16211 {
tmp16208 := PrimTail(V263)

tmp16209 := PrimIsPair(tmp16208)

var ifres16197 Obj

if True == tmp16209 {
tmp16205 := PrimTail(V263)

tmp16206 := PrimTail(tmp16205)

tmp16207 := PrimIsPair(tmp16206)

var ifres16199 Obj

if True == tmp16207 {
tmp16201 := PrimTail(V263)

tmp16202 := PrimTail(tmp16201)

tmp16203 := PrimTail(tmp16202)

tmp16204 := PrimEqual(Nil, tmp16203)

var ifres16200 Obj

if True == tmp16204 {
ifres16200 = True


} else {
ifres16200 = False


}

ifres16199 = ifres16200


} else {
ifres16199 = False


}

var ifres16198 Obj

if True == ifres16199 {
ifres16198 = True


} else {
ifres16198 = False


}

ifres16197 = ifres16198


} else {
ifres16197 = False


}

var ifres16196 Obj

if True == ifres16197 {
ifres16196 = True


} else {
ifres16196 = False


}

ifres16195 = ifres16196


} else {
ifres16195 = False


}

var ifres16194 Obj

if True == ifres16195 {
ifres16194 = True


} else {
ifres16194 = False


}

ifres16193 = ifres16194


} else {
ifres16193 = False


}

if True == ifres16193 {
tmp16189 := PrimTail(V263)

tmp16190 := PrimHead(tmp16189)

tmp16191 := Call(__e, PrimFunc(symshen_4syntax_1item_2), tmp16190)


if True == tmp16191 {
tmp16184 := PrimTail(V263)

tmp16185 := PrimTail(tmp16184)

tmp16186 := PrimHead(tmp16185)

tmp16187 := Call(__e, PrimFunc(symshen_4syntax_1item_2), tmp16186)


if True == tmp16187 {
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


} else {
__e.Return(False)
return
}


}


}


}


}, 1)

tmp16219 := Call(__e, ns2_1set, symshen_4syntax_1item_2, tmp16182)


_ = tmp16219

tmp16220 := MakeNative(func(__e *ControlFlow) {
V264 := __e.Get(1)
_ = V264
tmp16221 := MakeNative(func(__e *ControlFlow) {
W265 := __e.Get(1)
_ = W265
tmp16242 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W265)


if True == tmp16242 {
tmp16222 := MakeNative(func(__e *ControlFlow) {
W273 := __e.Get(1)
_ = W273
tmp16224 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W273)


if True == tmp16224 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W273)
return
}


}, 1)

tmp16225 := MakeNative(func(__e *ControlFlow) {
W274 := __e.Get(1)
_ = W274
tmp16238 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W274)


if True == tmp16238 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16226 := MakeNative(func(__e *ControlFlow) {
W275 := __e.Get(1)
_ = W275
tmp16235 := PrimIsPair(W275)

if True == tmp16235 {
tmp16227 := MakeNative(func(__e *ControlFlow) {
W276 := __e.Get(1)
_ = W276
tmp16228 := MakeNative(func(__e *ControlFlow) {
W277 := __e.Get(1)
_ = W277
tmp16230 := Call(__e, PrimFunc(symshen_4semicolon_2), W276)


tmp16231 := PrimNot(tmp16230)

if True == tmp16231 {
__e.TailApply(PrimFunc(symshen_4comb), W277, W276)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16232 := Call(__e, PrimFunc(symtail), W275)


__e.TailApply(tmp16228, tmp16232)
return


}, 1)

tmp16233 := Call(__e, PrimFunc(symhead), W275)


__e.TailApply(tmp16227, tmp16233)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16236 := Call(__e, PrimFunc(symshen_4in_1_6), W274)


__e.TailApply(tmp16226, tmp16236)
return


}


}, 1)

tmp16239 := Call(__e, PrimFunc(symshen_4_5colon_1equal_6), V264)


tmp16240 := Call(__e, tmp16225, tmp16239)


__e.TailApply(tmp16222, tmp16240)
return


} else {
__e.Return(W265)
return
}


}, 1)

tmp16243 := MakeNative(func(__e *ControlFlow) {
W266 := __e.Get(1)
_ = W266
tmp16269 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W266)


if True == tmp16269 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp16244 := MakeNative(func(__e *ControlFlow) {
W267 := __e.Get(1)
_ = W267
tmp16266 := PrimIsPair(W267)

if True == tmp16266 {
tmp16245 := MakeNative(func(__e *ControlFlow) {
W268 := __e.Get(1)
_ = W268
tmp16246 := MakeNative(func(__e *ControlFlow) {
W269 := __e.Get(1)
_ = W269
tmp16262 := Call(__e, PrimFunc(symshen_4hds_a_2), W269, symwhere)


if True == tmp16262 {
tmp16247 := MakeNative(func(__e *ControlFlow) {
W270 := __e.Get(1)
_ = W270
tmp16259 := PrimIsPair(W270)

if True == tmp16259 {
tmp16248 := MakeNative(func(__e *ControlFlow) {
W271 := __e.Get(1)
_ = W271
tmp16249 := MakeNative(func(__e *ControlFlow) {
W272 := __e.Get(1)
_ = W272
tmp16254 := Call(__e, PrimFunc(symshen_4semicolon_2), W268)


tmp16255 := PrimNot(tmp16254)

if True == tmp16255 {
tmp16250 := PrimCons(W268, Nil)

tmp16251 := PrimCons(W271, tmp16250)

tmp16252 := PrimCons(symwhere, tmp16251)

__e.TailApply(PrimFunc(symshen_4comb), W272, tmp16252)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16256 := Call(__e, PrimFunc(symtail), W270)


__e.TailApply(tmp16249, tmp16256)
return


}, 1)

tmp16257 := Call(__e, PrimFunc(symhead), W270)


__e.TailApply(tmp16248, tmp16257)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16260 := Call(__e, PrimFunc(symtail), W269)


__e.TailApply(tmp16247, tmp16260)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16263 := Call(__e, PrimFunc(symtail), W267)


__e.TailApply(tmp16246, tmp16263)
return


}, 1)

tmp16264 := Call(__e, PrimFunc(symhead), W267)


__e.TailApply(tmp16245, tmp16264)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp16267 := Call(__e, PrimFunc(symshen_4in_1_6), W266)


__e.TailApply(tmp16244, tmp16267)
return


}


}, 1)

tmp16270 := Call(__e, PrimFunc(symshen_4_5colon_1equal_6), V264)


tmp16271 := Call(__e, tmp16243, tmp16270)


__e.TailApply(tmp16221, tmp16271)
return


}, 1)

tmp16272 := Call(__e, ns2_1set, symshen_4_5semantics_6, tmp16220)


_ = tmp16272

tmp16273 := MakeNative(func(__e *ControlFlow) {
V286 := __e.Get(1)
_ = V286
V287 := __e.Get(2)
_ = V287
V288 := __e.Get(3)
_ = V288
tmp16281 := PrimEqual(Nil, V288)

if True == tmp16281 {
__e.Return(PrimCons(symshen_4parse_1failure, Nil))
return
} else {
tmp16279 := PrimIsPair(V288)

if True == tmp16279 {
tmp16274 := PrimHead(V288)

tmp16275 := Call(__e, PrimFunc(symshen_4c_1rule_1_6shen), V286, tmp16274, V287)


tmp16276 := PrimTail(V288)

tmp16277 := Call(__e, PrimFunc(symshen_4c_1rules_1_6shen), V286, V287, tmp16276)


__e.TailApply(PrimFunc(symshen_4combine_1c_1code), tmp16275, tmp16277)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.c-rules->shen\n")))
return
}


}


}, 3)

tmp16282 := Call(__e, ns2_1set, symshen_4c_1rules_1_6shen, tmp16273)


_ = tmp16282

tmp16283 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symfail))
return
}, 0)

tmp16284 := Call(__e, ns2_1set, symshen_4parse_1failure, tmp16283)


_ = tmp16284

tmp16285 := MakeNative(func(__e *ControlFlow) {
V289 := __e.Get(1)
_ = V289
V290 := __e.Get(2)
_ = V290
tmp16286 := PrimCons(symResult, Nil)

tmp16287 := PrimCons(symshen_4parse_1failure_2, tmp16286)

tmp16288 := PrimCons(symResult, Nil)

tmp16289 := PrimCons(V290, tmp16288)

tmp16290 := PrimCons(tmp16287, tmp16289)

tmp16291 := PrimCons(symif, tmp16290)

tmp16292 := PrimCons(tmp16291, Nil)

tmp16293 := PrimCons(V289, tmp16292)

tmp16294 := PrimCons(symResult, tmp16293)

__e.Return(PrimCons(symlet, tmp16294))
return


}, 2)

tmp16295 := Call(__e, ns2_1set, symshen_4combine_1c_1code, tmp16285)


_ = tmp16295

tmp16296 := MakeNative(func(__e *ControlFlow) {
V297 := __e.Get(1)
_ = V297
V298 := __e.Get(2)
_ = V298
V299 := __e.Get(3)
_ = V299
tmp16310 := PrimIsPair(V298)

var ifres16301 Obj

if True == tmp16310 {
tmp16308 := PrimTail(V298)

tmp16309 := PrimIsPair(tmp16308)

var ifres16303 Obj

if True == tmp16309 {
tmp16305 := PrimTail(V298)

tmp16306 := PrimTail(tmp16305)

tmp16307 := PrimEqual(Nil, tmp16306)

var ifres16304 Obj

if True == tmp16307 {
ifres16304 = True


} else {
ifres16304 = False


}

ifres16303 = ifres16304


} else {
ifres16303 = False


}

var ifres16302 Obj

if True == ifres16303 {
ifres16302 = True


} else {
ifres16302 = False


}

ifres16301 = ifres16302


} else {
ifres16301 = False


}

if True == ifres16301 {
tmp16297 := PrimHead(V298)

tmp16298 := PrimTail(V298)

tmp16299 := PrimHead(tmp16298)

__e.TailApply(PrimFunc(symshen_4yacc_1syntax), V297, V299, tmp16297, tmp16299)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.c-rule->shen\n")))
return
}


}, 3)

tmp16311 := Call(__e, ns2_1set, symshen_4c_1rule_1_6shen, tmp16296)


_ = tmp16311

tmp16312 := MakeNative(func(__e *ControlFlow) {
V308 := __e.Get(1)
_ = V308
V309 := __e.Get(2)
_ = V309
V310 := __e.Get(3)
_ = V310
V311 := __e.Get(4)
_ = V311
tmp16376 := PrimEqual(Nil, V310)

var ifres16354 Obj

if True == tmp16376 {
tmp16375 := PrimIsPair(V311)

var ifres16356 Obj

if True == tmp16375 {
tmp16373 := PrimHead(V311)

tmp16374 := PrimEqual(symwhere, tmp16373)

var ifres16358 Obj

if True == tmp16374 {
tmp16371 := PrimTail(V311)

tmp16372 := PrimIsPair(tmp16371)

var ifres16360 Obj

if True == tmp16372 {
tmp16368 := PrimTail(V311)

tmp16369 := PrimTail(tmp16368)

tmp16370 := PrimIsPair(tmp16369)

var ifres16362 Obj

if True == tmp16370 {
tmp16364 := PrimTail(V311)

tmp16365 := PrimTail(tmp16364)

tmp16366 := PrimTail(tmp16365)

tmp16367 := PrimEqual(Nil, tmp16366)

var ifres16363 Obj

if True == tmp16367 {
ifres16363 = True


} else {
ifres16363 = False


}

ifres16362 = ifres16363


} else {
ifres16362 = False


}

var ifres16361 Obj

if True == ifres16362 {
ifres16361 = True


} else {
ifres16361 = False


}

ifres16360 = ifres16361


} else {
ifres16360 = False


}

var ifres16359 Obj

if True == ifres16360 {
ifres16359 = True


} else {
ifres16359 = False


}

ifres16358 = ifres16359


} else {
ifres16358 = False


}

var ifres16357 Obj

if True == ifres16358 {
ifres16357 = True


} else {
ifres16357 = False


}

ifres16356 = ifres16357


} else {
ifres16356 = False


}

var ifres16355 Obj

if True == ifres16356 {
ifres16355 = True


} else {
ifres16355 = False


}

ifres16354 = ifres16355


} else {
ifres16354 = False


}

if True == ifres16354 {
tmp16313 := PrimTail(V311)

tmp16314 := PrimHead(tmp16313)

tmp16315 := Call(__e, PrimFunc(symshen_4process_1yacc_1semantics), tmp16314)


tmp16316 := PrimTail(V311)

tmp16317 := PrimTail(tmp16316)

tmp16318 := PrimHead(tmp16317)

tmp16319 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V308, V309, Nil, tmp16318)


tmp16320 := PrimCons(symshen_4parse_1failure, Nil)

tmp16321 := PrimCons(tmp16320, Nil)

tmp16322 := PrimCons(tmp16319, tmp16321)

tmp16323 := PrimCons(tmp16315, tmp16322)

__e.Return(PrimCons(symif, tmp16323))
return


} else {
tmp16352 := PrimEqual(Nil, V310)

if True == tmp16352 {
__e.TailApply(PrimFunc(symshen_4yacc_1semantics), V308, V309, V311)
return
} else {
tmp16350 := PrimIsPair(V310)

if True == tmp16350 {
tmp16347 := PrimHead(V310)

tmp16348 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp16347)


if True == tmp16348 {
tmp16324 := PrimHead(V310)

tmp16325 := PrimTail(V310)

__e.TailApply(PrimFunc(symshen_4non_1terminalcode), V308, V309, tmp16324, tmp16325, V311)
return


} else {
tmp16344 := PrimHead(V310)

tmp16345 := PrimIsVariable(tmp16344)

if True == tmp16345 {
tmp16326 := PrimHead(V310)

tmp16327 := PrimTail(V310)

__e.TailApply(PrimFunc(symshen_4variablecode), V308, V309, tmp16326, tmp16327, V311)
return


} else {
tmp16341 := PrimHead(V310)

tmp16342 := PrimEqual(sym__, tmp16341)

if True == tmp16342 {
tmp16328 := PrimHead(V310)

tmp16329 := PrimTail(V310)

__e.TailApply(PrimFunc(symshen_4wildcardcode), V308, V309, tmp16328, tmp16329, V311)
return


} else {
tmp16338 := PrimHead(V310)

tmp16339 := Call(__e, PrimFunc(symatom_2), tmp16338)


if True == tmp16339 {
tmp16330 := PrimHead(V310)

tmp16331 := PrimTail(V310)

__e.TailApply(PrimFunc(symshen_4terminalcode), V308, V309, tmp16330, tmp16331, V311)
return


} else {
tmp16335 := PrimHead(V310)

tmp16336 := PrimIsPair(tmp16335)

if True == tmp16336 {
tmp16332 := PrimHead(V310)

tmp16333 := PrimTail(V310)

__e.TailApply(PrimFunc(symshen_4conscode), V308, V309, tmp16332, tmp16333, V311)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.yacc-syntax\n")))
return
}


}


}


}


}


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.yacc-syntax\n")))
return
}


}


}


}, 4)

tmp16377 := Call(__e, ns2_1set, symshen_4yacc_1syntax, tmp16312)


_ = tmp16377

tmp16378 := MakeNative(func(__e *ControlFlow) {
V312 := __e.Get(1)
_ = V312
V313 := __e.Get(2)
_ = V313
V314 := __e.Get(3)
_ = V314
V315 := __e.Get(4)
_ = V315
V316 := __e.Get(5)
_ = V316
tmp16379 := MakeNative(func(__e *ControlFlow) {
W317 := __e.Get(1)
_ = W317
tmp16380 := MakeNative(func(__e *ControlFlow) {
W318 := __e.Get(1)
_ = W318
tmp16381 := MakeNative(func(__e *ControlFlow) {
W319 := __e.Get(1)
_ = W319
tmp16382 := PrimCons(V313, Nil)

tmp16383 := PrimCons(V314, tmp16382)

tmp16384 := PrimCons(W317, Nil)

tmp16385 := PrimCons(symshen_4parse_1failure_2, tmp16384)

tmp16386 := PrimCons(symshen_4parse_1failure, Nil)

tmp16387 := MakeNative(func(__e *ControlFlow) {
W320 := __e.Get(1)
_ = W320
tmp16397 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V314, V316)


var ifres16394 Obj

if True == tmp16397 {
ifres16394 = True


} else {
tmp16396 := Call(__e, PrimFunc(symshen_4occurs_1check_2), W318, V316)


var ifres16395 Obj

if True == tmp16396 {
ifres16395 = True


} else {
ifres16395 = False


}

ifres16394 = ifres16395


}

if True == ifres16394 {
tmp16388 := PrimCons(W317, Nil)

tmp16389 := PrimCons(symshen_4_5_1out, tmp16388)

tmp16390 := PrimCons(W320, Nil)

tmp16391 := PrimCons(tmp16389, tmp16390)

tmp16392 := PrimCons(W318, tmp16391)

__e.Return(PrimCons(symlet, tmp16392))
return


} else {
__e.Return(W320)
return
}


}, 1)

tmp16398 := PrimCons(W317, Nil)

tmp16399 := PrimCons(symshen_4in_1_6, tmp16398)

tmp16400 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V312, W319, V315, V316)


tmp16401 := PrimCons(tmp16400, Nil)

tmp16402 := PrimCons(tmp16399, tmp16401)

tmp16403 := PrimCons(W319, tmp16402)

tmp16404 := PrimCons(symlet, tmp16403)

tmp16405 := Call(__e, tmp16387, tmp16404)


tmp16406 := PrimCons(tmp16405, Nil)

tmp16407 := PrimCons(tmp16386, tmp16406)

tmp16408 := PrimCons(tmp16385, tmp16407)

tmp16409 := PrimCons(symif, tmp16408)

tmp16410 := PrimCons(tmp16409, Nil)

tmp16411 := PrimCons(tmp16383, tmp16410)

tmp16412 := PrimCons(W317, tmp16411)

__e.Return(PrimCons(symlet, tmp16412))
return


}, 1)

tmp16413 := Call(__e, PrimFunc(symconcat), symRemainder, V314)


__e.TailApply(tmp16381, tmp16413)
return


}, 1)

tmp16414 := Call(__e, PrimFunc(symconcat), symAction, V314)


__e.TailApply(tmp16380, tmp16414)
return


}, 1)

tmp16415 := Call(__e, PrimFunc(symconcat), symParse, V314)


__e.TailApply(tmp16379, tmp16415)
return


}, 5)

tmp16416 := Call(__e, ns2_1set, symshen_4non_1terminalcode, tmp16378)


_ = tmp16416

tmp16417 := MakeNative(func(__e *ControlFlow) {
V321 := __e.Get(1)
_ = V321
V322 := __e.Get(2)
_ = V322
V323 := __e.Get(3)
_ = V323
V324 := __e.Get(4)
_ = V324
V325 := __e.Get(5)
_ = V325
tmp16418 := MakeNative(func(__e *ControlFlow) {
W326 := __e.Get(1)
_ = W326
tmp16419 := PrimCons(V322, Nil)

tmp16420 := PrimCons(symcons_2, tmp16419)

tmp16421 := MakeNative(func(__e *ControlFlow) {
W327 := __e.Get(1)
_ = W327
tmp16428 := Call(__e, PrimFunc(symshen_4occurs_1check_2), V323, V325)


if True == tmp16428 {
tmp16422 := PrimCons(V322, Nil)

tmp16423 := PrimCons(symhead, tmp16422)

tmp16424 := PrimCons(W327, Nil)

tmp16425 := PrimCons(tmp16423, tmp16424)

tmp16426 := PrimCons(V323, tmp16425)

__e.Return(PrimCons(symlet, tmp16426))
return


} else {
__e.Return(W327)
return
}


}, 1)

tmp16429 := PrimCons(V322, Nil)

tmp16430 := PrimCons(symtail, tmp16429)

tmp16431 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V321, W326, V324, V325)


tmp16432 := PrimCons(tmp16431, Nil)

tmp16433 := PrimCons(tmp16430, tmp16432)

tmp16434 := PrimCons(W326, tmp16433)

tmp16435 := PrimCons(symlet, tmp16434)

tmp16436 := Call(__e, tmp16421, tmp16435)


tmp16437 := PrimCons(symshen_4parse_1failure, Nil)

tmp16438 := PrimCons(tmp16437, Nil)

tmp16439 := PrimCons(tmp16436, tmp16438)

tmp16440 := PrimCons(tmp16420, tmp16439)

__e.Return(PrimCons(symif, tmp16440))
return


}, 1)

tmp16441 := Call(__e, PrimFunc(symgensym), symRemainder)


__e.TailApply(tmp16418, tmp16441)
return


}, 5)

tmp16442 := Call(__e, ns2_1set, symshen_4variablecode, tmp16417)


_ = tmp16442

tmp16443 := MakeNative(func(__e *ControlFlow) {
V328 := __e.Get(1)
_ = V328
V329 := __e.Get(2)
_ = V329
V330 := __e.Get(3)
_ = V330
V331 := __e.Get(4)
_ = V331
V332 := __e.Get(5)
_ = V332
tmp16444 := MakeNative(func(__e *ControlFlow) {
W333 := __e.Get(1)
_ = W333
tmp16445 := PrimCons(V329, Nil)

tmp16446 := PrimCons(symcons_2, tmp16445)

tmp16447 := PrimCons(V329, Nil)

tmp16448 := PrimCons(symtail, tmp16447)

tmp16449 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V328, W333, V331, V332)


tmp16450 := PrimCons(tmp16449, Nil)

tmp16451 := PrimCons(tmp16448, tmp16450)

tmp16452 := PrimCons(W333, tmp16451)

tmp16453 := PrimCons(symlet, tmp16452)

tmp16454 := PrimCons(symshen_4parse_1failure, Nil)

tmp16455 := PrimCons(tmp16454, Nil)

tmp16456 := PrimCons(tmp16453, tmp16455)

tmp16457 := PrimCons(tmp16446, tmp16456)

__e.Return(PrimCons(symif, tmp16457))
return


}, 1)

tmp16458 := Call(__e, PrimFunc(symgensym), symRemainder)


__e.TailApply(tmp16444, tmp16458)
return


}, 5)

tmp16459 := Call(__e, ns2_1set, symshen_4wildcardcode, tmp16443)


_ = tmp16459

tmp16460 := MakeNative(func(__e *ControlFlow) {
V334 := __e.Get(1)
_ = V334
V335 := __e.Get(2)
_ = V335
V336 := __e.Get(3)
_ = V336
V337 := __e.Get(4)
_ = V337
V338 := __e.Get(5)
_ = V338
tmp16461 := MakeNative(func(__e *ControlFlow) {
W339 := __e.Get(1)
_ = W339
tmp16462 := PrimCons(V336, Nil)

tmp16463 := PrimCons(V335, tmp16462)

tmp16464 := PrimCons(symshen_4hds_a_2, tmp16463)

tmp16465 := PrimCons(V335, Nil)

tmp16466 := PrimCons(symtail, tmp16465)

tmp16467 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V334, W339, V337, V338)


tmp16468 := PrimCons(tmp16467, Nil)

tmp16469 := PrimCons(tmp16466, tmp16468)

tmp16470 := PrimCons(W339, tmp16469)

tmp16471 := PrimCons(symlet, tmp16470)

tmp16472 := PrimCons(symshen_4parse_1failure, Nil)

tmp16473 := PrimCons(tmp16472, Nil)

tmp16474 := PrimCons(tmp16471, tmp16473)

tmp16475 := PrimCons(tmp16464, tmp16474)

__e.Return(PrimCons(symif, tmp16475))
return


}, 1)

tmp16476 := Call(__e, PrimFunc(symgensym), symRemainder)


__e.TailApply(tmp16461, tmp16476)
return


}, 5)

tmp16477 := Call(__e, ns2_1set, symshen_4terminalcode, tmp16460)


_ = tmp16477

tmp16478 := MakeNative(func(__e *ControlFlow) {
V347 := __e.Get(1)
_ = V347
V348 := __e.Get(2)
_ = V348
tmp16484 := PrimIsPair(V347)

var ifres16480 Obj

if True == tmp16484 {
tmp16482 := PrimHead(V347)

tmp16483 := PrimEqual(tmp16482, V348)

var ifres16481 Obj

if True == tmp16483 {
ifres16481 = True


} else {
ifres16481 = False


}

ifres16480 = ifres16481


} else {
ifres16480 = False


}

if True == ifres16480 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp16485 := Call(__e, ns2_1set, symshen_4hds_a_2, tmp16478)


_ = tmp16485

tmp16486 := MakeNative(func(__e *ControlFlow) {
V349 := __e.Get(1)
_ = V349
V350 := __e.Get(2)
_ = V350
V351 := __e.Get(3)
_ = V351
V352 := __e.Get(4)
_ = V352
V353 := __e.Get(5)
_ = V353
tmp16487 := MakeNative(func(__e *ControlFlow) {
W354 := __e.Get(1)
_ = W354
tmp16488 := MakeNative(func(__e *ControlFlow) {
W355 := __e.Get(1)
_ = W355
tmp16489 := MakeNative(func(__e *ControlFlow) {
W356 := __e.Get(1)
_ = W356
tmp16490 := PrimCons(V350, Nil)

tmp16491 := PrimCons(symshen_4ccons_2, tmp16490)

tmp16492 := PrimCons(V350, Nil)

tmp16493 := PrimCons(symhead, tmp16492)

tmp16494 := PrimCons(V350, Nil)

tmp16495 := PrimCons(symtail, tmp16494)

tmp16496 := Call(__e, PrimFunc(symshen_4decons), V351)


tmp16497 := PrimCons(sym_5end_6, Nil)

tmp16498 := Call(__e, PrimFunc(symappend), tmp16496, tmp16497)


tmp16499 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V349, W356, V352, V353)


tmp16500 := PrimCons(tmp16499, Nil)

tmp16501 := PrimCons(symshen_4processed, tmp16500)

tmp16502 := Call(__e, PrimFunc(symshen_4yacc_1syntax), V349, W355, tmp16498, tmp16501)


tmp16503 := PrimCons(tmp16502, Nil)

tmp16504 := PrimCons(tmp16495, tmp16503)

tmp16505 := PrimCons(W356, tmp16504)

tmp16506 := PrimCons(tmp16493, tmp16505)

tmp16507 := PrimCons(W355, tmp16506)

tmp16508 := PrimCons(symlet, tmp16507)

tmp16509 := PrimCons(symshen_4parse_1failure, Nil)

tmp16510 := PrimCons(tmp16509, Nil)

tmp16511 := PrimCons(tmp16508, tmp16510)

tmp16512 := PrimCons(tmp16491, tmp16511)

__e.Return(PrimCons(symif, tmp16512))
return


}, 1)

tmp16513 := Call(__e, PrimFunc(symgensym), symTl)


__e.TailApply(tmp16489, tmp16513)
return


}, 1)

tmp16514 := Call(__e, PrimFunc(symgensym), symHd)


__e.TailApply(tmp16488, tmp16514)
return


}, 1)

tmp16515 := Call(__e, PrimFunc(symgensym), symRemainder)


__e.TailApply(tmp16487, tmp16515)
return


}, 5)

tmp16516 := Call(__e, ns2_1set, symshen_4conscode, tmp16486)


_ = tmp16516

tmp16517 := MakeNative(func(__e *ControlFlow) {
V365 := __e.Get(1)
_ = V365
tmp16523 := PrimIsPair(V365)

var ifres16519 Obj

if True == tmp16523 {
tmp16521 := PrimHead(V365)

tmp16522 := PrimIsPair(tmp16521)

var ifres16520 Obj

if True == tmp16522 {
ifres16520 = True


} else {
ifres16520 = False


}

ifres16519 = ifres16520


} else {
ifres16519 = False


}

if True == ifres16519 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp16524 := Call(__e, ns2_1set, symshen_4ccons_2, tmp16517)


_ = tmp16524

tmp16525 := MakeNative(func(__e *ControlFlow) {
V366 := __e.Get(1)
_ = V366
tmp16552 := PrimIsPair(V366)

var ifres16533 Obj

if True == tmp16552 {
tmp16550 := PrimHead(V366)

tmp16551 := PrimEqual(symcons, tmp16550)

var ifres16535 Obj

if True == tmp16551 {
tmp16548 := PrimTail(V366)

tmp16549 := PrimIsPair(tmp16548)

var ifres16537 Obj

if True == tmp16549 {
tmp16545 := PrimTail(V366)

tmp16546 := PrimTail(tmp16545)

tmp16547 := PrimIsPair(tmp16546)

var ifres16539 Obj

if True == tmp16547 {
tmp16541 := PrimTail(V366)

tmp16542 := PrimTail(tmp16541)

tmp16543 := PrimTail(tmp16542)

tmp16544 := PrimEqual(Nil, tmp16543)

var ifres16540 Obj

if True == tmp16544 {
ifres16540 = True


} else {
ifres16540 = False


}

ifres16539 = ifres16540


} else {
ifres16539 = False


}

var ifres16538 Obj

if True == ifres16539 {
ifres16538 = True


} else {
ifres16538 = False


}

ifres16537 = ifres16538


} else {
ifres16537 = False


}

var ifres16536 Obj

if True == ifres16537 {
ifres16536 = True


} else {
ifres16536 = False


}

ifres16535 = ifres16536


} else {
ifres16535 = False


}

var ifres16534 Obj

if True == ifres16535 {
ifres16534 = True


} else {
ifres16534 = False


}

ifres16533 = ifres16534


} else {
ifres16533 = False


}

if True == ifres16533 {
tmp16526 := PrimTail(V366)

tmp16527 := PrimHead(tmp16526)

tmp16528 := PrimTail(V366)

tmp16529 := PrimTail(tmp16528)

tmp16530 := PrimHead(tmp16529)

tmp16531 := Call(__e, PrimFunc(symshen_4decons), tmp16530)


__e.Return(PrimCons(tmp16527, tmp16531))
return


} else {
__e.Return(V366)
return
}


}, 1)

tmp16553 := Call(__e, ns2_1set, symshen_4decons, tmp16525)


_ = tmp16553

tmp16554 := MakeNative(func(__e *ControlFlow) {
V367 := __e.Get(1)
_ = V367
V368 := __e.Get(2)
_ = V368
tmp16555 := PrimCons(V368, Nil)

__e.Return(PrimCons(V367, tmp16555))
return


}, 2)

tmp16556 := Call(__e, ns2_1set, symshen_4comb, tmp16554)


_ = tmp16556

tmp16557 := MakeNative(func(__e *ControlFlow) {
V373 := __e.Get(1)
_ = V373
V374 := __e.Get(2)
_ = V374
V375 := __e.Get(3)
_ = V375
tmp16579 := PrimIsPair(V375)

var ifres16566 Obj

if True == tmp16579 {
tmp16577 := PrimHead(V375)

tmp16578 := PrimEqual(symshen_4processed, tmp16577)

var ifres16568 Obj

if True == tmp16578 {
tmp16575 := PrimTail(V375)

tmp16576 := PrimIsPair(tmp16575)

var ifres16570 Obj

if True == tmp16576 {
tmp16572 := PrimTail(V375)

tmp16573 := PrimTail(tmp16572)

tmp16574 := PrimEqual(Nil, tmp16573)

var ifres16571 Obj

if True == tmp16574 {
ifres16571 = True


} else {
ifres16571 = False


}

ifres16570 = ifres16571


} else {
ifres16570 = False


}

var ifres16569 Obj

if True == ifres16570 {
ifres16569 = True


} else {
ifres16569 = False


}

ifres16568 = ifres16569


} else {
ifres16568 = False


}

var ifres16567 Obj

if True == ifres16568 {
ifres16567 = True


} else {
ifres16567 = False


}

ifres16566 = ifres16567


} else {
ifres16566 = False


}

if True == ifres16566 {
tmp16558 := PrimTail(V375)

__e.Return(PrimHead(tmp16558))
return


} else {
tmp16559 := MakeNative(func(__e *ControlFlow) {
W376 := __e.Get(1)
_ = W376
tmp16560 := MakeNative(func(__e *ControlFlow) {
W377 := __e.Get(1)
_ = W377
tmp16561 := PrimCons(W377, Nil)

tmp16562 := PrimCons(V374, tmp16561)

__e.Return(PrimCons(symshen_4comb, tmp16562))
return


}, 1)

tmp16563 := Call(__e, PrimFunc(symshen_4use_1type_1info), V373, W376)


__e.TailApply(tmp16560, tmp16563)
return


}, 1)

tmp16564 := Call(__e, PrimFunc(symshen_4process_1yacc_1semantics), V375)


__e.TailApply(tmp16559, tmp16564)
return


}


}, 3)

tmp16580 := Call(__e, ns2_1set, symshen_4yacc_1semantics, tmp16557)


_ = tmp16580

tmp16581 := MakeNative(func(__e *ControlFlow) {
V381 := __e.Get(1)
_ = V381
V382 := __e.Get(2)
_ = V382
tmp16769 := PrimIsPair(V381)

var ifres16590 Obj

if True == tmp16769 {
tmp16767 := PrimHead(V381)

tmp16768 := PrimEqual(sym_i, tmp16767)

var ifres16592 Obj

if True == tmp16768 {
tmp16765 := PrimTail(V381)

tmp16766 := PrimIsPair(tmp16765)

var ifres16594 Obj

if True == tmp16766 {
tmp16762 := PrimTail(V381)

tmp16763 := PrimHead(tmp16762)

tmp16764 := PrimIsPair(tmp16763)

var ifres16596 Obj

if True == tmp16764 {
tmp16758 := PrimTail(V381)

tmp16759 := PrimHead(tmp16758)

tmp16760 := PrimHead(tmp16759)

tmp16761 := PrimEqual(symlist, tmp16760)

var ifres16598 Obj

if True == tmp16761 {
tmp16754 := PrimTail(V381)

tmp16755 := PrimHead(tmp16754)

tmp16756 := PrimTail(tmp16755)

tmp16757 := PrimIsPair(tmp16756)

var ifres16600 Obj

if True == tmp16757 {
tmp16749 := PrimTail(V381)

tmp16750 := PrimHead(tmp16749)

tmp16751 := PrimTail(tmp16750)

tmp16752 := PrimTail(tmp16751)

tmp16753 := PrimEqual(Nil, tmp16752)

var ifres16602 Obj

if True == tmp16753 {
tmp16746 := PrimTail(V381)

tmp16747 := PrimTail(tmp16746)

tmp16748 := PrimIsPair(tmp16747)

var ifres16604 Obj

if True == tmp16748 {
tmp16742 := PrimTail(V381)

tmp16743 := PrimTail(tmp16742)

tmp16744 := PrimHead(tmp16743)

tmp16745 := PrimEqual(sym_1_1_6, tmp16744)

var ifres16606 Obj

if True == tmp16745 {
tmp16738 := PrimTail(V381)

tmp16739 := PrimTail(tmp16738)

tmp16740 := PrimTail(tmp16739)

tmp16741 := PrimIsPair(tmp16740)

var ifres16608 Obj

if True == tmp16741 {
tmp16733 := PrimTail(V381)

tmp16734 := PrimTail(tmp16733)

tmp16735 := PrimTail(tmp16734)

tmp16736 := PrimHead(tmp16735)

tmp16737 := PrimIsPair(tmp16736)

var ifres16610 Obj

if True == tmp16737 {
tmp16727 := PrimTail(V381)

tmp16728 := PrimTail(tmp16727)

tmp16729 := PrimTail(tmp16728)

tmp16730 := PrimHead(tmp16729)

tmp16731 := PrimHead(tmp16730)

tmp16732 := PrimEqual(symstr, tmp16731)

var ifres16612 Obj

if True == tmp16732 {
tmp16721 := PrimTail(V381)

tmp16722 := PrimTail(tmp16721)

tmp16723 := PrimTail(tmp16722)

tmp16724 := PrimHead(tmp16723)

tmp16725 := PrimTail(tmp16724)

tmp16726 := PrimIsPair(tmp16725)

var ifres16614 Obj

if True == tmp16726 {
tmp16714 := PrimTail(V381)

tmp16715 := PrimTail(tmp16714)

tmp16716 := PrimTail(tmp16715)

tmp16717 := PrimHead(tmp16716)

tmp16718 := PrimTail(tmp16717)

tmp16719 := PrimHead(tmp16718)

tmp16720 := PrimIsPair(tmp16719)

var ifres16616 Obj

if True == tmp16720 {
tmp16706 := PrimTail(V381)

tmp16707 := PrimTail(tmp16706)

tmp16708 := PrimTail(tmp16707)

tmp16709 := PrimHead(tmp16708)

tmp16710 := PrimTail(tmp16709)

tmp16711 := PrimHead(tmp16710)

tmp16712 := PrimHead(tmp16711)

tmp16713 := PrimEqual(symlist, tmp16712)

var ifres16618 Obj

if True == tmp16713 {
tmp16698 := PrimTail(V381)

tmp16699 := PrimTail(tmp16698)

tmp16700 := PrimTail(tmp16699)

tmp16701 := PrimHead(tmp16700)

tmp16702 := PrimTail(tmp16701)

tmp16703 := PrimHead(tmp16702)

tmp16704 := PrimTail(tmp16703)

tmp16705 := PrimIsPair(tmp16704)

var ifres16620 Obj

if True == tmp16705 {
tmp16689 := PrimTail(V381)

tmp16690 := PrimTail(tmp16689)

tmp16691 := PrimTail(tmp16690)

tmp16692 := PrimHead(tmp16691)

tmp16693 := PrimTail(tmp16692)

tmp16694 := PrimHead(tmp16693)

tmp16695 := PrimTail(tmp16694)

tmp16696 := PrimTail(tmp16695)

tmp16697 := PrimEqual(Nil, tmp16696)

var ifres16622 Obj

if True == tmp16697 {
tmp16682 := PrimTail(V381)

tmp16683 := PrimTail(tmp16682)

tmp16684 := PrimTail(tmp16683)

tmp16685 := PrimHead(tmp16684)

tmp16686 := PrimTail(tmp16685)

tmp16687 := PrimTail(tmp16686)

tmp16688 := PrimIsPair(tmp16687)

var ifres16624 Obj

if True == tmp16688 {
tmp16674 := PrimTail(V381)

tmp16675 := PrimTail(tmp16674)

tmp16676 := PrimTail(tmp16675)

tmp16677 := PrimHead(tmp16676)

tmp16678 := PrimTail(tmp16677)

tmp16679 := PrimTail(tmp16678)

tmp16680 := PrimTail(tmp16679)

tmp16681 := PrimEqual(Nil, tmp16680)

var ifres16626 Obj

if True == tmp16681 {
tmp16669 := PrimTail(V381)

tmp16670 := PrimTail(tmp16669)

tmp16671 := PrimTail(tmp16670)

tmp16672 := PrimTail(tmp16671)

tmp16673 := PrimIsPair(tmp16672)

var ifres16628 Obj

if True == tmp16673 {
tmp16663 := PrimTail(V381)

tmp16664 := PrimTail(tmp16663)

tmp16665 := PrimTail(tmp16664)

tmp16666 := PrimTail(tmp16665)

tmp16667 := PrimHead(tmp16666)

tmp16668 := PrimEqual(sym_j, tmp16667)

var ifres16630 Obj

if True == tmp16668 {
tmp16657 := PrimTail(V381)

tmp16658 := PrimTail(tmp16657)

tmp16659 := PrimTail(tmp16658)

tmp16660 := PrimTail(tmp16659)

tmp16661 := PrimTail(tmp16660)

tmp16662 := PrimEqual(Nil, tmp16661)

var ifres16632 Obj

if True == tmp16662 {
tmp16644 := PrimTail(V381)

tmp16645 := PrimHead(tmp16644)

tmp16646 := PrimTail(tmp16645)

tmp16647 := PrimHead(tmp16646)

tmp16648 := PrimTail(V381)

tmp16649 := PrimTail(tmp16648)

tmp16650 := PrimTail(tmp16649)

tmp16651 := PrimHead(tmp16650)

tmp16652 := PrimTail(tmp16651)

tmp16653 := PrimHead(tmp16652)

tmp16654 := PrimTail(tmp16653)

tmp16655 := PrimHead(tmp16654)

tmp16656 := PrimEqual(tmp16647, tmp16655)

var ifres16634 Obj

if True == tmp16656 {
tmp16636 := PrimTail(V381)

tmp16637 := PrimTail(tmp16636)

tmp16638 := PrimTail(tmp16637)

tmp16639 := PrimHead(tmp16638)

tmp16640 := PrimTail(tmp16639)

tmp16641 := PrimTail(tmp16640)

tmp16642 := PrimHead(tmp16641)

tmp16643 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp16642)


var ifres16635 Obj

if True == tmp16643 {
ifres16635 = True


} else {
ifres16635 = False


}

ifres16634 = ifres16635


} else {
ifres16634 = False


}

var ifres16633 Obj

if True == ifres16634 {
ifres16633 = True


} else {
ifres16633 = False


}

ifres16632 = ifres16633


} else {
ifres16632 = False


}

var ifres16631 Obj

if True == ifres16632 {
ifres16631 = True


} else {
ifres16631 = False


}

ifres16630 = ifres16631


} else {
ifres16630 = False


}

var ifres16629 Obj

if True == ifres16630 {
ifres16629 = True


} else {
ifres16629 = False


}

ifres16628 = ifres16629


} else {
ifres16628 = False


}

var ifres16627 Obj

if True == ifres16628 {
ifres16627 = True


} else {
ifres16627 = False


}

ifres16626 = ifres16627


} else {
ifres16626 = False


}

var ifres16625 Obj

if True == ifres16626 {
ifres16625 = True


} else {
ifres16625 = False


}

ifres16624 = ifres16625


} else {
ifres16624 = False


}

var ifres16623 Obj

if True == ifres16624 {
ifres16623 = True


} else {
ifres16623 = False


}

ifres16622 = ifres16623


} else {
ifres16622 = False


}

var ifres16621 Obj

if True == ifres16622 {
ifres16621 = True


} else {
ifres16621 = False


}

ifres16620 = ifres16621


} else {
ifres16620 = False


}

var ifres16619 Obj

if True == ifres16620 {
ifres16619 = True


} else {
ifres16619 = False


}

ifres16618 = ifres16619


} else {
ifres16618 = False


}

var ifres16617 Obj

if True == ifres16618 {
ifres16617 = True


} else {
ifres16617 = False


}

ifres16616 = ifres16617


} else {
ifres16616 = False


}

var ifres16615 Obj

if True == ifres16616 {
ifres16615 = True


} else {
ifres16615 = False


}

ifres16614 = ifres16615


} else {
ifres16614 = False


}

var ifres16613 Obj

if True == ifres16614 {
ifres16613 = True


} else {
ifres16613 = False


}

ifres16612 = ifres16613


} else {
ifres16612 = False


}

var ifres16611 Obj

if True == ifres16612 {
ifres16611 = True


} else {
ifres16611 = False


}

ifres16610 = ifres16611


} else {
ifres16610 = False


}

var ifres16609 Obj

if True == ifres16610 {
ifres16609 = True


} else {
ifres16609 = False


}

ifres16608 = ifres16609


} else {
ifres16608 = False


}

var ifres16607 Obj

if True == ifres16608 {
ifres16607 = True


} else {
ifres16607 = False


}

ifres16606 = ifres16607


} else {
ifres16606 = False


}

var ifres16605 Obj

if True == ifres16606 {
ifres16605 = True


} else {
ifres16605 = False


}

ifres16604 = ifres16605


} else {
ifres16604 = False


}

var ifres16603 Obj

if True == ifres16604 {
ifres16603 = True


} else {
ifres16603 = False


}

ifres16602 = ifres16603


} else {
ifres16602 = False


}

var ifres16601 Obj

if True == ifres16602 {
ifres16601 = True


} else {
ifres16601 = False


}

ifres16600 = ifres16601


} else {
ifres16600 = False


}

var ifres16599 Obj

if True == ifres16600 {
ifres16599 = True


} else {
ifres16599 = False


}

ifres16598 = ifres16599


} else {
ifres16598 = False


}

var ifres16597 Obj

if True == ifres16598 {
ifres16597 = True


} else {
ifres16597 = False


}

ifres16596 = ifres16597


} else {
ifres16596 = False


}

var ifres16595 Obj

if True == ifres16596 {
ifres16595 = True


} else {
ifres16595 = False


}

ifres16594 = ifres16595


} else {
ifres16594 = False


}

var ifres16593 Obj

if True == ifres16594 {
ifres16593 = True


} else {
ifres16593 = False


}

ifres16592 = ifres16593


} else {
ifres16592 = False


}

var ifres16591 Obj

if True == ifres16592 {
ifres16591 = True


} else {
ifres16591 = False


}

ifres16590 = ifres16591


} else {
ifres16590 = False


}

if True == ifres16590 {
tmp16582 := PrimTail(V381)

tmp16583 := PrimTail(tmp16582)

tmp16584 := PrimTail(tmp16583)

tmp16585 := PrimHead(tmp16584)

tmp16586 := PrimTail(tmp16585)

tmp16587 := PrimTail(tmp16586)

tmp16588 := PrimCons(V382, tmp16587)

__e.Return(PrimCons(symtype, tmp16588))
return


} else {
__e.Return(V382)
return
}


}, 2)

tmp16770 := Call(__e, ns2_1set, symshen_4use_1type_1info, tmp16581)


_ = tmp16770

tmp16771 := MakeNative(func(__e *ControlFlow) {
V385 := __e.Get(1)
_ = V385
tmp16781 := PrimIsVariable(V385)

if True == tmp16781 {
__e.Return(False)
return
} else {
tmp16779 := PrimIsPair(V385)

if True == tmp16779 {
tmp16776 := PrimHead(V385)

tmp16777 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp16776)


if True == tmp16777 {
tmp16773 := PrimTail(V385)

tmp16774 := Call(__e, PrimFunc(symshen_4monomorphic_2), tmp16773)


if True == tmp16774 {
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


} else {
__e.Return(True)
return
}


}


}, 1)

tmp16782 := Call(__e, ns2_1set, symshen_4monomorphic_2, tmp16771)


_ = tmp16782

tmp16783 := MakeNative(func(__e *ControlFlow) {
V386 := __e.Get(1)
_ = V386
tmp16809 := PrimIsPair(V386)

var ifres16791 Obj

if True == tmp16809 {
tmp16807 := PrimHead(V386)

tmp16808 := PrimEqual(symprotect, tmp16807)

var ifres16793 Obj

if True == tmp16808 {
tmp16805 := PrimTail(V386)

tmp16806 := PrimIsPair(tmp16805)

var ifres16795 Obj

if True == tmp16806 {
tmp16802 := PrimTail(V386)

tmp16803 := PrimTail(tmp16802)

tmp16804 := PrimEqual(Nil, tmp16803)

var ifres16797 Obj

if True == tmp16804 {
tmp16799 := PrimTail(V386)

tmp16800 := PrimHead(tmp16799)

tmp16801 := Call(__e, PrimFunc(symshen_4non_1terminal_2), tmp16800)


var ifres16798 Obj

if True == tmp16801 {
ifres16798 = True


} else {
ifres16798 = False


}

ifres16797 = ifres16798


} else {
ifres16797 = False


}

var ifres16796 Obj

if True == ifres16797 {
ifres16796 = True


} else {
ifres16796 = False


}

ifres16795 = ifres16796


} else {
ifres16795 = False


}

var ifres16794 Obj

if True == ifres16795 {
ifres16794 = True


} else {
ifres16794 = False


}

ifres16793 = ifres16794


} else {
ifres16793 = False


}

var ifres16792 Obj

if True == ifres16793 {
ifres16792 = True


} else {
ifres16792 = False


}

ifres16791 = ifres16792


} else {
ifres16791 = False


}

if True == ifres16791 {
tmp16784 := PrimTail(V386)

__e.Return(PrimHead(tmp16784))
return


} else {
tmp16789 := PrimIsPair(V386)

if True == tmp16789 {
tmp16785 := MakeNative(func(__e *ControlFlow) {
Z387 := __e.Get(1)
_ = Z387
__e.TailApply(PrimFunc(symshen_4process_1yacc_1semantics), Z387)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp16785, V386)
return


} else {
tmp16787 := Call(__e, PrimFunc(symshen_4non_1terminal_2), V386)


if True == tmp16787 {
__e.TailApply(PrimFunc(symconcat), symAction, V386)
return
} else {
__e.Return(V386)
return
}


}


}


}, 1)

tmp16810 := Call(__e, ns2_1set, symshen_4process_1yacc_1semantics, tmp16783)


_ = tmp16810

tmp16811 := MakeNative(func(__e *ControlFlow) {
V390 := __e.Get(1)
_ = V390
tmp16824 := PrimIsPair(V390)

var ifres16815 Obj

if True == tmp16824 {
tmp16822 := PrimTail(V390)

tmp16823 := PrimIsPair(tmp16822)

var ifres16817 Obj

if True == tmp16823 {
tmp16819 := PrimTail(V390)

tmp16820 := PrimTail(tmp16819)

tmp16821 := PrimEqual(Nil, tmp16820)

var ifres16818 Obj

if True == tmp16821 {
ifres16818 = True


} else {
ifres16818 = False


}

ifres16817 = ifres16818


} else {
ifres16817 = False


}

var ifres16816 Obj

if True == ifres16817 {
ifres16816 = True


} else {
ifres16816 = False


}

ifres16815 = ifres16816


} else {
ifres16815 = False


}

if True == ifres16815 {
tmp16812 := PrimTail(V390)

__e.Return(PrimHead(tmp16812))
return


} else {
tmp16813 := PrimTail(V390)

__e.Return(PrimHead(tmp16813))
return


}


}, 1)

tmp16825 := Call(__e, ns2_1set, symshen_4_5_1out, tmp16811)


_ = tmp16825

tmp16826 := MakeNative(func(__e *ControlFlow) {
V391 := __e.Get(1)
_ = V391
__e.Return(PrimHead(V391))
return
}, 1)

tmp16827 := Call(__e, ns2_1set, symshen_4in_1_6, tmp16826)


_ = tmp16827

tmp16828 := MakeNative(func(__e *ControlFlow) {
V392 := __e.Get(1)
_ = V392
tmp16829 := PrimCons(V392, Nil)

__e.Return(PrimCons(Nil, tmp16829))
return


}, 1)

tmp16830 := Call(__e, ns2_1set, sym_5_b_6, tmp16828)


_ = tmp16830

tmp16831 := MakeNative(func(__e *ControlFlow) {
V393 := __e.Get(1)
_ = V393
tmp16832 := PrimCons(Nil, Nil)

__e.Return(PrimCons(V393, tmp16832))
return


}, 1)

tmp16833 := Call(__e, ns2_1set, sym_5e_6, tmp16831)


_ = tmp16833

tmp16834 := MakeNative(func(__e *ControlFlow) {
V396 := __e.Get(1)
_ = V396
tmp16837 := PrimEqual(Nil, V396)

if True == tmp16837 {
tmp16835 := PrimCons(Nil, Nil)

__e.Return(PrimCons(Nil, tmp16835))
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

__e.TailApply(ns2_1set, sym_5end_6, tmp16834)
return




}, 0)

