package main

import . "github.com/pyrex41/shen-go/kl"

var SequentMain = MakeNative(func(__e *ControlFlow) {
tmp12788 := MakeNative(func(__e *ControlFlow) {
V3033 := __e.Get(1)
_ = V3033
tmp12803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3033)
}
__typedArg0 := V3033
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres12789 Obj

if True == tmp12803 {
tmp12790 := Call(__e, PrimFunc(symhead), V3033)


W303512558 := tmp12790
_ = W303512558

tmp12791 := Call(__e, PrimFunc(symtail), V3033)


W303612559 := tmp12791
_ = W303612559

tmp12792 := Call(__e, PrimFunc(symshen_4_5datatype_1rules_6), W303612559)


W303712560 := tmp12792
_ = W303712560

tmp12801 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W303712560)


var ifres12793 Obj

if True == tmp12801 {
tmp12794 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12793 = tmp12794


} else {
tmp12795 := Call(__e, PrimFunc(symshen_4_5_1out), W303712560)


W303812561 := tmp12795
_ = W303812561

tmp12796 := Call(__e, PrimFunc(symshen_4in_1_6), W303712560)


W303912562 := tmp12796
_ = W303912562

tmp12797 := Call(__e, PrimFunc(symshen_4rules_1_6prolog), W303512558, W303812561)


W304012563 := tmp12797
_ = W304012563

tmp12798 := Call(__e, PrimFunc(symfn), W303512558)


tmp12799 := Call(__e, PrimFunc(symshen_4remember_1datatype), W303512558, tmp12798)


tmp12800 := Call(__e, PrimFunc(symshen_4comb), W303912562, tmp12799)


ifres12793 = tmp12800


}

ifres12789 = ifres12793


} else {
tmp12802 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12789 = tmp12802


}

W303412557 := ifres12789
_ = W303412557

tmp12805 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W303412557)


if True == tmp12805 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W303412557)
return
}


}, 1)

tmp12806 := Call(__e, ns2_1set, symshen_4_5datatype_6, tmp12788)


_ = tmp12806

tmp12807 := MakeNative(func(__e *ControlFlow) {
V3041 := __e.Get(1)
_ = V3041
V3042 := __e.Get(2)
_ = V3042
tmp12808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp12809 := Call(__e, PrimFunc(symshen_4assoc_1_6), V3041, V3042, tmp12808)


tmp12810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, tmp12809)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := tmp12809
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp12810

tmp12811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dalldatatypes_d)
}
__typedArg0 := symshen_4_dalldatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp12812 := Call(__e, PrimFunc(symshen_4assoc_1_6), V3041, V3042, tmp12811)


tmp12813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dalldatatypes_d, tmp12812)
}
__typedArg0 := symshen_4_dalldatatypes_d
__typedArg1 := tmp12812
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp12813

__e.Return(V3041)
return


}, 2)

tmp12814 := Call(__e, ns2_1set, symshen_4remember_1datatype, tmp12807)


_ = tmp12814

tmp12815 := MakeNative(func(__e *ControlFlow) {
V3043 := __e.Get(1)
_ = V3043
tmp12816 := Call(__e, PrimFunc(symshen_4_5datatype_1rule_6), V3043)


W304512565 := tmp12816
_ = W304512565

tmp12829 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W304512565)


var ifres12817 Obj

if True == tmp12829 {
tmp12818 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12817 = tmp12818


} else {
tmp12819 := Call(__e, PrimFunc(symshen_4_5_1out), W304512565)


W304612566 := tmp12819
_ = W304612566

tmp12820 := Call(__e, PrimFunc(symshen_4in_1_6), W304512565)


W304712567 := tmp12820
_ = W304712567

tmp12821 := Call(__e, PrimFunc(symshen_4_5datatype_1rules_6), W304712567)


W304812568 := tmp12821
_ = W304812568

tmp12828 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W304812568)


var ifres12822 Obj

if True == tmp12828 {
tmp12823 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12822 = tmp12823


} else {
tmp12824 := Call(__e, PrimFunc(symshen_4_5_1out), W304812568)


W304912569 := tmp12824
_ = W304912569

tmp12825 := Call(__e, PrimFunc(symshen_4in_1_6), W304812568)


W305012570 := tmp12825
_ = W305012570

tmp12826 := Call(__e, PrimFunc(symappend), W304612566, W304912569)


tmp12827 := Call(__e, PrimFunc(symshen_4comb), W305012570, tmp12826)


ifres12822 = tmp12827


}

ifres12817 = ifres12822


}

W304412564 := ifres12817
_ = W304412564

tmp12845 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W304412564)


if True == tmp12845 {
tmp12830 := Call(__e, PrimFunc(sym_5_b_6), V3043)


W305212572 := tmp12830
_ = W305212572

tmp12841 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W305212572)


var ifres12831 Obj

if True == tmp12841 {
tmp12832 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12831 = tmp12832


} else {
tmp12833 := Call(__e, PrimFunc(symshen_4_5_1out), W305212572)


W305312573 := tmp12833
_ = W305312573

tmp12834 := Call(__e, PrimFunc(symshen_4in_1_6), W305212572)


W305412574 := tmp12834
_ = W305412574

tmp12839 := Call(__e, PrimFunc(symempty_2), W305312573)


var ifres12835 Obj

if True == tmp12839 {
ifres12835 = Nil


} else {
tmp12836 := Call(__e, PrimFunc(symshen_4app), W305312573, MakeString("\n ..."), symshen_4r)


tmp12838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("datatype syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp12836)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("datatype syntax error here:\n ")
__typedArg1 := tmp12836
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("datatype syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp12836)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("datatype syntax error here:\n ")
__typedArg1 := tmp12836
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres12835 = tmp12838


}

tmp12840 := Call(__e, PrimFunc(symshen_4comb), W305412574, ifres12835)


ifres12831 = tmp12840


}

W305112571 := ifres12831
_ = W305112571

tmp12843 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W305112571)


if True == tmp12843 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W305112571)
return
}


} else {
__e.Return(W304412564)
return
}


}, 1)

tmp12846 := Call(__e, ns2_1set, symshen_4_5datatype_1rules_6, tmp12815)


_ = tmp12846

tmp12847 := MakeNative(func(__e *ControlFlow) {
V3055 := __e.Get(1)
_ = V3055
tmp12848 := Call(__e, PrimFunc(symshen_4_5single_6), V3055)


W305712576 := tmp12848
_ = W305712576

tmp12854 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W305712576)


var ifres12849 Obj

if True == tmp12854 {
tmp12850 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12849 = tmp12850


} else {
tmp12851 := Call(__e, PrimFunc(symshen_4_5_1out), W305712576)


W305812577 := tmp12851
_ = W305812577

tmp12852 := Call(__e, PrimFunc(symshen_4in_1_6), W305712576)


W305912578 := tmp12852
_ = W305912578

tmp12853 := Call(__e, PrimFunc(symshen_4comb), W305912578, W305812577)


ifres12849 = tmp12853


}

W305612575 := ifres12849
_ = W305612575

tmp12865 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W305612575)


if True == tmp12865 {
tmp12855 := Call(__e, PrimFunc(symshen_4_5double_6), V3055)


W306112580 := tmp12855
_ = W306112580

tmp12861 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W306112580)


var ifres12856 Obj

if True == tmp12861 {
tmp12857 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12856 = tmp12857


} else {
tmp12858 := Call(__e, PrimFunc(symshen_4_5_1out), W306112580)


W306212581 := tmp12858
_ = W306212581

tmp12859 := Call(__e, PrimFunc(symshen_4in_1_6), W306112580)


W306312582 := tmp12859
_ = W306312582

tmp12860 := Call(__e, PrimFunc(symshen_4comb), W306312582, W306212581)


ifres12856 = tmp12860


}

W306012579 := ifres12856
_ = W306012579

tmp12863 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W306012579)


if True == tmp12863 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W306012579)
return
}


} else {
__e.Return(W305612575)
return
}


}, 1)

tmp12866 := Call(__e, ns2_1set, symshen_4_5datatype_1rule_6, tmp12847)


_ = tmp12866

tmp12867 := MakeNative(func(__e *ControlFlow) {
V3064 := __e.Get(1)
_ = V3064
tmp12868 := Call(__e, PrimFunc(symshen_4_5sides_6), V3064)


W306612584 := tmp12868
_ = W306612584

tmp12900 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W306612584)


var ifres12869 Obj

if True == tmp12900 {
tmp12870 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12869 = tmp12870


} else {
tmp12871 := Call(__e, PrimFunc(symshen_4_5_1out), W306612584)


W306712585 := tmp12871
_ = W306712585

tmp12872 := Call(__e, PrimFunc(symshen_4in_1_6), W306612584)


W306812586 := tmp12872
_ = W306812586

tmp12873 := Call(__e, PrimFunc(symshen_4_5prems_6), W306812586)


W306912587 := tmp12873
_ = W306912587

tmp12899 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W306912587)


var ifres12874 Obj

if True == tmp12899 {
tmp12875 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12874 = tmp12875


} else {
tmp12876 := Call(__e, PrimFunc(symshen_4_5_1out), W306912587)


W307012588 := tmp12876
_ = W307012588

tmp12877 := Call(__e, PrimFunc(symshen_4in_1_6), W306912587)


W307112589 := tmp12877
_ = W307112589

tmp12878 := Call(__e, PrimFunc(symshen_4_5sng_6), W307112589)


W307212590 := tmp12878
_ = W307212590

tmp12898 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W307212590)


var ifres12879 Obj

if True == tmp12898 {
tmp12880 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12879 = tmp12880


} else {
tmp12881 := Call(__e, PrimFunc(symshen_4in_1_6), W307212590)


W307312591 := tmp12881
_ = W307312591

tmp12882 := Call(__e, PrimFunc(symshen_4_5conc_6), W307312591)


W307412592 := tmp12882
_ = W307412592

tmp12897 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W307412592)


var ifres12883 Obj

if True == tmp12897 {
tmp12884 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12883 = tmp12884


} else {
tmp12885 := Call(__e, PrimFunc(symshen_4_5_1out), W307412592)


W307512593 := tmp12885
_ = W307512593

tmp12886 := Call(__e, PrimFunc(symshen_4in_1_6), W307412592)


W307612594 := tmp12886
_ = W307612594

tmp12887 := Call(__e, PrimFunc(symshen_4_5sc_6), W307612594)


W307712595 := tmp12887
_ = W307712595

tmp12896 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W307712595)


var ifres12888 Obj

if True == tmp12896 {
tmp12889 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12888 = tmp12889


} else {
tmp12890 := Call(__e, PrimFunc(symshen_4in_1_6), W307712595)


W307812596 := tmp12890
_ = W307812596

tmp12891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W307512593, Nil)
}
__typedArg0 := W307512593
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W307012588, tmp12891)
}
__typedArg0 := W307012588
__typedArg1 := tmp12891
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W306712585, tmp12892)
}
__typedArg0 := W306712585
__typedArg1 := tmp12892
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12893, Nil)
}
__typedArg0 := tmp12893
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12895 := Call(__e, PrimFunc(symshen_4comb), W307812596, tmp12894)


ifres12888 = tmp12895


}

ifres12883 = ifres12888


}

ifres12879 = ifres12883


}

ifres12874 = ifres12879


}

ifres12869 = ifres12874


}

W306512583 := ifres12869
_ = W306512583

tmp12902 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W306512583)


if True == tmp12902 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W306512583)
return
}


}, 1)

tmp12903 := Call(__e, ns2_1set, symshen_4_5single_6, tmp12867)


_ = tmp12903

tmp12904 := MakeNative(func(__e *ControlFlow) {
V3079 := __e.Get(1)
_ = V3079
tmp12905 := Call(__e, PrimFunc(symshen_4_5sides_6), V3079)


W308112598 := tmp12905
_ = W308112598

tmp12936 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W308112598)


var ifres12906 Obj

if True == tmp12936 {
tmp12907 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12906 = tmp12907


} else {
tmp12908 := Call(__e, PrimFunc(symshen_4_5_1out), W308112598)


W308212599 := tmp12908
_ = W308212599

tmp12909 := Call(__e, PrimFunc(symshen_4in_1_6), W308112598)


W308312600 := tmp12909
_ = W308312600

tmp12910 := Call(__e, PrimFunc(symshen_4_5formulae_6), W308312600)


W308412601 := tmp12910
_ = W308412601

tmp12935 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W308412601)


var ifres12911 Obj

if True == tmp12935 {
tmp12912 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12911 = tmp12912


} else {
tmp12913 := Call(__e, PrimFunc(symshen_4_5_1out), W308412601)


W308512602 := tmp12913
_ = W308512602

tmp12914 := Call(__e, PrimFunc(symshen_4in_1_6), W308412601)


W308612603 := tmp12914
_ = W308612603

tmp12915 := Call(__e, PrimFunc(symshen_4_5dbl_6), W308612603)


W308712604 := tmp12915
_ = W308712604

tmp12934 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W308712604)


var ifres12916 Obj

if True == tmp12934 {
tmp12917 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12916 = tmp12917


} else {
tmp12918 := Call(__e, PrimFunc(symshen_4in_1_6), W308712604)


W308812605 := tmp12918
_ = W308812605

tmp12919 := Call(__e, PrimFunc(symshen_4_5formula_6), W308812605)


W308912606 := tmp12919
_ = W308912606

tmp12933 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W308912606)


var ifres12920 Obj

if True == tmp12933 {
tmp12921 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12920 = tmp12921


} else {
tmp12922 := Call(__e, PrimFunc(symshen_4_5_1out), W308912606)


W309012607 := tmp12922
_ = W309012607

tmp12923 := Call(__e, PrimFunc(symshen_4in_1_6), W308912606)


W309112608 := tmp12923
_ = W309112608

tmp12924 := Call(__e, PrimFunc(symshen_4_5sc_6), W309112608)


W309212609 := tmp12924
_ = W309212609

tmp12932 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W309212609)


var ifres12925 Obj

if True == tmp12932 {
tmp12926 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12925 = tmp12926


} else {
tmp12927 := Call(__e, PrimFunc(symshen_4in_1_6), W309212609)


W309312610 := tmp12927
_ = W309312610

tmp12928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W309012607, Nil)
}
__typedArg0 := W309012607
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp12928)
}
__typedArg0 := Nil
__typedArg1 := tmp12928
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12930 := Call(__e, PrimFunc(symshen_4lr_1rule), W308212599, W308512602, tmp12929)


tmp12931 := Call(__e, PrimFunc(symshen_4comb), W309312610, tmp12930)


ifres12925 = tmp12931


}

ifres12920 = ifres12925


}

ifres12916 = ifres12920


}

ifres12911 = ifres12916


}

ifres12906 = ifres12911


}

W308012597 := ifres12906
_ = W308012597

tmp12938 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W308012597)


if True == tmp12938 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W308012597)
return
}


}, 1)

tmp12939 := Call(__e, ns2_1set, symshen_4_5double_6, tmp12904)


_ = tmp12939

tmp12940 := MakeNative(func(__e *ControlFlow) {
V3094 := __e.Get(1)
_ = V3094
tmp12941 := Call(__e, PrimFunc(symshen_4_5formula_6), V3094)


W309612612 := tmp12941
_ = W309612612

tmp12961 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W309612612)


var ifres12942 Obj

if True == tmp12961 {
tmp12943 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12942 = tmp12943


} else {
tmp12944 := Call(__e, PrimFunc(symshen_4_5_1out), W309612612)


W309712613 := tmp12944
_ = W309712613

tmp12945 := Call(__e, PrimFunc(symshen_4in_1_6), W309612612)


W309812614 := tmp12945
_ = W309812614

tmp12946 := Call(__e, PrimFunc(symshen_4_5sc_6), W309812614)


W309912615 := tmp12946
_ = W309912615

tmp12960 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W309912615)


var ifres12947 Obj

if True == tmp12960 {
tmp12948 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12947 = tmp12948


} else {
tmp12949 := Call(__e, PrimFunc(symshen_4in_1_6), W309912615)


W310012616 := tmp12949
_ = W310012616

tmp12950 := Call(__e, PrimFunc(symshen_4_5formulae_6), W310012616)


W310112617 := tmp12950
_ = W310112617

tmp12959 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W310112617)


var ifres12951 Obj

if True == tmp12959 {
tmp12952 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12951 = tmp12952


} else {
tmp12953 := Call(__e, PrimFunc(symshen_4_5_1out), W310112617)


W310212618 := tmp12953
_ = W310212618

tmp12954 := Call(__e, PrimFunc(symshen_4in_1_6), W310112617)


W310312619 := tmp12954
_ = W310312619

tmp12955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W309712613, Nil)
}
__typedArg0 := W309712613
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp12955)
}
__typedArg0 := Nil
__typedArg1 := tmp12955
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12956, W310212618)
}
__typedArg0 := tmp12956
__typedArg1 := W310212618
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12958 := Call(__e, PrimFunc(symshen_4comb), W310312619, tmp12957)


ifres12951 = tmp12958


}

ifres12947 = ifres12951


}

ifres12942 = ifres12947


}

W309512611 := ifres12942
_ = W309512611

tmp12980 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W309512611)


if True == tmp12980 {
tmp12962 := Call(__e, PrimFunc(symshen_4_5formula_6), V3094)


W310512621 := tmp12962
_ = W310512621

tmp12976 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W310512621)


var ifres12963 Obj

if True == tmp12976 {
tmp12964 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12963 = tmp12964


} else {
tmp12965 := Call(__e, PrimFunc(symshen_4_5_1out), W310512621)


W310612622 := tmp12965
_ = W310612622

tmp12966 := Call(__e, PrimFunc(symshen_4in_1_6), W310512621)


W310712623 := tmp12966
_ = W310712623

tmp12967 := Call(__e, PrimFunc(symshen_4_5sc_6), W310712623)


W310812624 := tmp12967
_ = W310812624

tmp12975 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W310812624)


var ifres12968 Obj

if True == tmp12975 {
tmp12969 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12968 = tmp12969


} else {
tmp12970 := Call(__e, PrimFunc(symshen_4in_1_6), W310812624)


W310912625 := tmp12970
_ = W310912625

tmp12971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W310612622, Nil)
}
__typedArg0 := W310612622
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp12971)
}
__typedArg0 := Nil
__typedArg1 := tmp12971
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp12972, Nil)
}
__typedArg0 := tmp12972
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12974 := Call(__e, PrimFunc(symshen_4comb), W310912625, tmp12973)


ifres12968 = tmp12974


}

ifres12963 = ifres12968


}

W310412620 := ifres12963
_ = W310412620

tmp12978 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W310412620)


if True == tmp12978 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W310412620)
return
}


} else {
__e.Return(W309512611)
return
}


}, 1)

tmp12981 := Call(__e, ns2_1set, symshen_4_5formulae_6, tmp12940)


_ = tmp12981

tmp12982 := MakeNative(func(__e *ControlFlow) {
V3110 := __e.Get(1)
_ = V3110
tmp12983 := Call(__e, PrimFunc(symshen_4_5ass_6), V3110)


W311212627 := tmp12983
_ = W311212627

tmp13001 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W311212627)


var ifres12984 Obj

if True == tmp13001 {
tmp12985 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12984 = tmp12985


} else {
tmp12986 := Call(__e, PrimFunc(symshen_4_5_1out), W311212627)


W311312628 := tmp12986
_ = W311312628

tmp12987 := Call(__e, PrimFunc(symshen_4in_1_6), W311212627)


W311412629 := tmp12987
_ = W311412629

tmp13000 := Call(__e, PrimFunc(symshen_4hds_a_2), W311412629, sym_6_6)


var ifres12988 Obj

if True == tmp13000 {
tmp12989 := Call(__e, PrimFunc(symtail), W311412629)


W311512630 := tmp12989
_ = W311512630

tmp12990 := Call(__e, PrimFunc(symshen_4_5formula_6), W311512630)


W311612631 := tmp12990
_ = W311612631

tmp12998 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W311612631)


var ifres12991 Obj

if True == tmp12998 {
tmp12992 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12991 = tmp12992


} else {
tmp12993 := Call(__e, PrimFunc(symshen_4_5_1out), W311612631)


W311712632 := tmp12993
_ = W311712632

tmp12994 := Call(__e, PrimFunc(symshen_4in_1_6), W311612631)


W311812633 := tmp12994
_ = W311812633

tmp12995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W311712632, Nil)
}
__typedArg0 := W311712632
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W311312628, tmp12995)
}
__typedArg0 := W311312628
__typedArg1 := tmp12995
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp12997 := Call(__e, PrimFunc(symshen_4comb), W311812633, tmp12996)


ifres12991 = tmp12997


}

ifres12988 = ifres12991


} else {
tmp12999 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12988 = tmp12999


}

ifres12984 = ifres12988


}

W311112626 := ifres12984
_ = W311112626

tmp13014 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W311112626)


if True == tmp13014 {
tmp13002 := Call(__e, PrimFunc(symshen_4_5formula_6), V3110)


W312012635 := tmp13002
_ = W312012635

tmp13010 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W312012635)


var ifres13003 Obj

if True == tmp13010 {
tmp13004 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13003 = tmp13004


} else {
tmp13005 := Call(__e, PrimFunc(symshen_4_5_1out), W312012635)


W312112636 := tmp13005
_ = W312112636

tmp13006 := Call(__e, PrimFunc(symshen_4in_1_6), W312012635)


W312212637 := tmp13006
_ = W312212637

tmp13007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W312112636, Nil)
}
__typedArg0 := W312112636
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp13007)
}
__typedArg0 := Nil
__typedArg1 := tmp13007
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13009 := Call(__e, PrimFunc(symshen_4comb), W312212637, tmp13008)


ifres13003 = tmp13009


}

W311912634 := ifres13003
_ = W311912634

tmp13012 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W311912634)


if True == tmp13012 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W311912634)
return
}


} else {
__e.Return(W311112626)
return
}


}, 1)

tmp13015 := Call(__e, ns2_1set, symshen_4_5conc_6, tmp12982)


_ = tmp13015

tmp13016 := MakeNative(func(__e *ControlFlow) {
V3123 := __e.Get(1)
_ = V3123
tmp13017 := Call(__e, PrimFunc(symshen_4_5prem_6), V3123)


W312512639 := tmp13017
_ = W312512639

tmp13035 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W312512639)


var ifres13018 Obj

if True == tmp13035 {
tmp13019 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13018 = tmp13019


} else {
tmp13020 := Call(__e, PrimFunc(symshen_4_5_1out), W312512639)


W312612640 := tmp13020
_ = W312612640

tmp13021 := Call(__e, PrimFunc(symshen_4in_1_6), W312512639)


W312712641 := tmp13021
_ = W312712641

tmp13022 := Call(__e, PrimFunc(symshen_4_5sc_6), W312712641)


W312812642 := tmp13022
_ = W312812642

tmp13034 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W312812642)


var ifres13023 Obj

if True == tmp13034 {
tmp13024 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13023 = tmp13024


} else {
tmp13025 := Call(__e, PrimFunc(symshen_4in_1_6), W312812642)


W312912643 := tmp13025
_ = W312912643

tmp13026 := Call(__e, PrimFunc(symshen_4_5prems_6), W312912643)


W313012644 := tmp13026
_ = W313012644

tmp13033 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W313012644)


var ifres13027 Obj

if True == tmp13033 {
tmp13028 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13027 = tmp13028


} else {
tmp13029 := Call(__e, PrimFunc(symshen_4_5_1out), W313012644)


W313112645 := tmp13029
_ = W313112645

tmp13030 := Call(__e, PrimFunc(symshen_4in_1_6), W313012644)


W313212646 := tmp13030
_ = W313212646

tmp13031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W312612640, W313112645)
}
__typedArg0 := W312612640
__typedArg1 := W313112645
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13032 := Call(__e, PrimFunc(symshen_4comb), W313212646, tmp13031)


ifres13027 = tmp13032


}

ifres13023 = ifres13027


}

ifres13018 = ifres13023


}

W312412638 := ifres13018
_ = W312412638

tmp13045 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W312412638)


if True == tmp13045 {
tmp13036 := Call(__e, PrimFunc(sym_5e_6), V3123)


W313412648 := tmp13036
_ = W313412648

tmp13041 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W313412648)


var ifres13037 Obj

if True == tmp13041 {
tmp13038 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13037 = tmp13038


} else {
tmp13039 := Call(__e, PrimFunc(symshen_4in_1_6), W313412648)


W313512649 := tmp13039
_ = W313512649

tmp13040 := Call(__e, PrimFunc(symshen_4comb), W313512649, Nil)


ifres13037 = tmp13040


}

W313312647 := ifres13037
_ = W313312647

tmp13043 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W313312647)


if True == tmp13043 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W313312647)
return
}


} else {
__e.Return(W312412638)
return
}


}, 1)

tmp13046 := Call(__e, ns2_1set, symshen_4_5prems_6, tmp13016)


_ = tmp13046

tmp13047 := MakeNative(func(__e *ControlFlow) {
V3136 := __e.Get(1)
_ = V3136
tmp13052 := Call(__e, PrimFunc(symshen_4hds_a_2), V3136, sym_b)


var ifres13048 Obj

if True == tmp13052 {
tmp13049 := Call(__e, PrimFunc(symtail), V3136)


W313812651 := tmp13049
_ = W313812651

tmp13050 := Call(__e, PrimFunc(symshen_4comb), W313812651, sym_b)


ifres13048 = tmp13050


} else {
tmp13051 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13048 = tmp13051


}

W313712650 := ifres13048
_ = W313712650

tmp13086 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W313712650)


if True == tmp13086 {
tmp13053 := Call(__e, PrimFunc(symshen_4_5ass_6), V3136)


W314012653 := tmp13053
_ = W314012653

tmp13071 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W314012653)


var ifres13054 Obj

if True == tmp13071 {
tmp13055 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13054 = tmp13055


} else {
tmp13056 := Call(__e, PrimFunc(symshen_4_5_1out), W314012653)


W314112654 := tmp13056
_ = W314112654

tmp13057 := Call(__e, PrimFunc(symshen_4in_1_6), W314012653)


W314212655 := tmp13057
_ = W314212655

tmp13070 := Call(__e, PrimFunc(symshen_4hds_a_2), W314212655, sym_6_6)


var ifres13058 Obj

if True == tmp13070 {
tmp13059 := Call(__e, PrimFunc(symtail), W314212655)


W314312656 := tmp13059
_ = W314312656

tmp13060 := Call(__e, PrimFunc(symshen_4_5formula_6), W314312656)


W314412657 := tmp13060
_ = W314412657

tmp13068 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W314412657)


var ifres13061 Obj

if True == tmp13068 {
tmp13062 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13061 = tmp13062


} else {
tmp13063 := Call(__e, PrimFunc(symshen_4_5_1out), W314412657)


W314512658 := tmp13063
_ = W314512658

tmp13064 := Call(__e, PrimFunc(symshen_4in_1_6), W314412657)


W314612659 := tmp13064
_ = W314612659

tmp13065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W314512658, Nil)
}
__typedArg0 := W314512658
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W314112654, tmp13065)
}
__typedArg0 := W314112654
__typedArg1 := tmp13065
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13067 := Call(__e, PrimFunc(symshen_4comb), W314612659, tmp13066)


ifres13061 = tmp13067


}

ifres13058 = ifres13061


} else {
tmp13069 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13058 = tmp13069


}

ifres13054 = ifres13058


}

W313912652 := ifres13054
_ = W313912652

tmp13084 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W313912652)


if True == tmp13084 {
tmp13072 := Call(__e, PrimFunc(symshen_4_5formula_6), V3136)


W314812661 := tmp13072
_ = W314812661

tmp13080 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W314812661)


var ifres13073 Obj

if True == tmp13080 {
tmp13074 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13073 = tmp13074


} else {
tmp13075 := Call(__e, PrimFunc(symshen_4_5_1out), W314812661)


W314912662 := tmp13075
_ = W314912662

tmp13076 := Call(__e, PrimFunc(symshen_4in_1_6), W314812661)


W315012663 := tmp13076
_ = W315012663

tmp13077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W314912662, Nil)
}
__typedArg0 := W314912662
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(Nil, tmp13077)
}
__typedArg0 := Nil
__typedArg1 := tmp13077
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13079 := Call(__e, PrimFunc(symshen_4comb), W315012663, tmp13078)


ifres13073 = tmp13079


}

W314712660 := ifres13073
_ = W314712660

tmp13082 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W314712660)


if True == tmp13082 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W314712660)
return
}


} else {
__e.Return(W313912652)
return
}


} else {
__e.Return(W313712650)
return
}


}, 1)

tmp13087 := Call(__e, ns2_1set, symshen_4_5prem_6, tmp13047)


_ = tmp13087

tmp13088 := MakeNative(func(__e *ControlFlow) {
V3151 := __e.Get(1)
_ = V3151
tmp13089 := Call(__e, PrimFunc(symshen_4_5formula_6), V3151)


W315312665 := tmp13089
_ = W315312665

tmp13107 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W315312665)


var ifres13090 Obj

if True == tmp13107 {
tmp13091 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13090 = tmp13091


} else {
tmp13092 := Call(__e, PrimFunc(symshen_4_5_1out), W315312665)


W315412666 := tmp13092
_ = W315412666

tmp13093 := Call(__e, PrimFunc(symshen_4in_1_6), W315312665)


W315512667 := tmp13093
_ = W315512667

tmp13094 := Call(__e, PrimFunc(symshen_4_5iscomma_6), W315512667)


W315612668 := tmp13094
_ = W315612668

tmp13106 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W315612668)


var ifres13095 Obj

if True == tmp13106 {
tmp13096 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13095 = tmp13096


} else {
tmp13097 := Call(__e, PrimFunc(symshen_4in_1_6), W315612668)


W315712669 := tmp13097
_ = W315712669

tmp13098 := Call(__e, PrimFunc(symshen_4_5ass_6), W315712669)


W315812670 := tmp13098
_ = W315812670

tmp13105 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W315812670)


var ifres13099 Obj

if True == tmp13105 {
tmp13100 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13099 = tmp13100


} else {
tmp13101 := Call(__e, PrimFunc(symshen_4_5_1out), W315812670)


W315912671 := tmp13101
_ = W315912671

tmp13102 := Call(__e, PrimFunc(symshen_4in_1_6), W315812670)


W316012672 := tmp13102
_ = W316012672

tmp13103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W315412666, W315912671)
}
__typedArg0 := W315412666
__typedArg1 := W315912671
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13104 := Call(__e, PrimFunc(symshen_4comb), W316012672, tmp13103)


ifres13099 = tmp13104


}

ifres13095 = ifres13099


}

ifres13090 = ifres13095


}

W315212664 := ifres13090
_ = W315212664

tmp13127 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W315212664)


if True == tmp13127 {
tmp13108 := Call(__e, PrimFunc(symshen_4_5formula_6), V3151)


W316212674 := tmp13108
_ = W316212674

tmp13115 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W316212674)


var ifres13109 Obj

if True == tmp13115 {
tmp13110 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13109 = tmp13110


} else {
tmp13111 := Call(__e, PrimFunc(symshen_4_5_1out), W316212674)


W316312675 := tmp13111
_ = W316312675

tmp13112 := Call(__e, PrimFunc(symshen_4in_1_6), W316212674)


W316412676 := tmp13112
_ = W316412676

tmp13113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W316312675, Nil)
}
__typedArg0 := W316312675
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13114 := Call(__e, PrimFunc(symshen_4comb), W316412676, tmp13113)


ifres13109 = tmp13114


}

W316112673 := ifres13109
_ = W316112673

tmp13125 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W316112673)


if True == tmp13125 {
tmp13116 := Call(__e, PrimFunc(sym_5e_6), V3151)


W316612678 := tmp13116
_ = W316612678

tmp13121 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W316612678)


var ifres13117 Obj

if True == tmp13121 {
tmp13118 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13117 = tmp13118


} else {
tmp13119 := Call(__e, PrimFunc(symshen_4in_1_6), W316612678)


W316712679 := tmp13119
_ = W316712679

tmp13120 := Call(__e, PrimFunc(symshen_4comb), W316712679, Nil)


ifres13117 = tmp13120


}

W316512677 := ifres13117
_ = W316512677

tmp13123 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W316512677)


if True == tmp13123 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W316512677)
return
}


} else {
__e.Return(W316112673)
return
}


} else {
__e.Return(W315212664)
return
}


}, 1)

tmp13128 := Call(__e, ns2_1set, symshen_4_5ass_6, tmp13088)


_ = tmp13128

tmp13129 := MakeNative(func(__e *ControlFlow) {
V3168 := __e.Get(1)
_ = V3168
tmp13139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3168)
}
__typedArg0 := V3168
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13130 Obj

if True == tmp13139 {
tmp13131 := Call(__e, PrimFunc(symhead), V3168)


W317012681 := tmp13131
_ = W317012681

tmp13132 := Call(__e, PrimFunc(symtail), V3168)


W317112682 := tmp13132
_ = W317112682

tmp13136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(","))
}
__typedArg0 := MakeString(",")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W317012681, tmp13136)
}
__typedArg0 := W317012681
__typedArg1 := tmp13136
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13133 Obj

if True == tmp13137 {
tmp13134 := Call(__e, PrimFunc(symshen_4comb), W317112682, symshen_4skip)


ifres13133 = tmp13134


} else {
tmp13135 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13133 = tmp13135


}

ifres13130 = ifres13133


} else {
tmp13138 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13130 = tmp13138


}

W316912680 := ifres13130
_ = W316912680

tmp13141 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W316912680)


if True == tmp13141 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W316912680)
return
}


}, 1)

tmp13142 := Call(__e, ns2_1set, symshen_4_5iscomma_6, tmp13129)


_ = tmp13142

tmp13143 := MakeNative(func(__e *ControlFlow) {
V3172 := __e.Get(1)
_ = V3172
tmp13144 := Call(__e, PrimFunc(symshen_4_5expr_6), V3172)


W317412684 := tmp13144
_ = W317412684

tmp13167 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W317412684)


var ifres13145 Obj

if True == tmp13167 {
tmp13146 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13145 = tmp13146


} else {
tmp13147 := Call(__e, PrimFunc(symshen_4_5_1out), W317412684)


W317512685 := tmp13147
_ = W317512685

tmp13148 := Call(__e, PrimFunc(symshen_4in_1_6), W317412684)


W317612686 := tmp13148
_ = W317612686

tmp13149 := Call(__e, PrimFunc(symshen_4_5iscolon_6), W317612686)


W317712687 := tmp13149
_ = W317712687

tmp13166 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W317712687)


var ifres13150 Obj

if True == tmp13166 {
tmp13151 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13150 = tmp13151


} else {
tmp13152 := Call(__e, PrimFunc(symshen_4in_1_6), W317712687)


W317812688 := tmp13152
_ = W317812688

tmp13153 := Call(__e, PrimFunc(symshen_4_5type_6), W317812688)


W317912689 := tmp13153
_ = W317912689

tmp13165 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W317912689)


var ifres13154 Obj

if True == tmp13165 {
tmp13155 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13154 = tmp13155


} else {
tmp13156 := Call(__e, PrimFunc(symshen_4_5_1out), W317912689)


W318012690 := tmp13156
_ = W318012690

tmp13157 := Call(__e, PrimFunc(symshen_4in_1_6), W317912689)


W318112691 := tmp13157
_ = W318112691

tmp13158 := Call(__e, PrimFunc(symshen_4curry), W317512685)


tmp13159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13160 := Call(__e, PrimFunc(symshen_4rectify_1type), W318012690)


tmp13161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13160, Nil)
}
__typedArg0 := tmp13160
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13159, tmp13161)
}
__typedArg0 := tmp13159
__typedArg1 := tmp13161
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13158, tmp13162)
}
__typedArg0 := tmp13158
__typedArg1 := tmp13162
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13164 := Call(__e, PrimFunc(symshen_4comb), W318112691, tmp13163)


ifres13154 = tmp13164


}

ifres13150 = ifres13154


}

ifres13145 = ifres13150


}

W317312683 := ifres13145
_ = W317312683

tmp13178 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W317312683)


if True == tmp13178 {
tmp13168 := Call(__e, PrimFunc(symshen_4_5expr_6), V3172)


W318312693 := tmp13168
_ = W318312693

tmp13174 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W318312693)


var ifres13169 Obj

if True == tmp13174 {
tmp13170 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13169 = tmp13170


} else {
tmp13171 := Call(__e, PrimFunc(symshen_4_5_1out), W318312693)


W318412694 := tmp13171
_ = W318412694

tmp13172 := Call(__e, PrimFunc(symshen_4in_1_6), W318312693)


W318512695 := tmp13172
_ = W318512695

tmp13173 := Call(__e, PrimFunc(symshen_4comb), W318512695, W318412694)


ifres13169 = tmp13173


}

W318212692 := ifres13169
_ = W318212692

tmp13176 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W318212692)


if True == tmp13176 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W318212692)
return
}


} else {
__e.Return(W317312683)
return
}


}, 1)

tmp13179 := Call(__e, ns2_1set, symshen_4_5formula_6, tmp13143)


_ = tmp13179

tmp13180 := MakeNative(func(__e *ControlFlow) {
V3186 := __e.Get(1)
_ = V3186
tmp13190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3186)
}
__typedArg0 := V3186
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13181 Obj

if True == tmp13190 {
tmp13182 := Call(__e, PrimFunc(symhead), V3186)


W318812697 := tmp13182
_ = W318812697

tmp13183 := Call(__e, PrimFunc(symtail), V3186)


W318912698 := tmp13183
_ = W318912698

tmp13187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W318812697, tmp13187)
}
__typedArg0 := W318812697
__typedArg1 := tmp13187
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13184 Obj

if True == tmp13188 {
tmp13185 := Call(__e, PrimFunc(symshen_4comb), W318912698, symshen_4skip)


ifres13184 = tmp13185


} else {
tmp13186 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13184 = tmp13186


}

ifres13181 = ifres13184


} else {
tmp13189 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13181 = tmp13189


}

W318712696 := ifres13181
_ = W318712696

tmp13192 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W318712696)


if True == tmp13192 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W318712696)
return
}


}, 1)

tmp13193 := Call(__e, ns2_1set, symshen_4_5iscolon_6, tmp13180)


_ = tmp13193

tmp13194 := MakeNative(func(__e *ControlFlow) {
V3190 := __e.Get(1)
_ = V3190
tmp13195 := Call(__e, PrimFunc(symshen_4_5side_6), V3190)


W319212700 := tmp13195
_ = W319212700

tmp13208 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W319212700)


var ifres13196 Obj

if True == tmp13208 {
tmp13197 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13196 = tmp13197


} else {
tmp13198 := Call(__e, PrimFunc(symshen_4_5_1out), W319212700)


W319312701 := tmp13198
_ = W319312701

tmp13199 := Call(__e, PrimFunc(symshen_4in_1_6), W319212700)


W319412702 := tmp13199
_ = W319412702

tmp13200 := Call(__e, PrimFunc(symshen_4_5sides_6), W319412702)


W319512703 := tmp13200
_ = W319512703

tmp13207 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W319512703)


var ifres13201 Obj

if True == tmp13207 {
tmp13202 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13201 = tmp13202


} else {
tmp13203 := Call(__e, PrimFunc(symshen_4_5_1out), W319512703)


W319612704 := tmp13203
_ = W319612704

tmp13204 := Call(__e, PrimFunc(symshen_4in_1_6), W319512703)


W319712705 := tmp13204
_ = W319712705

tmp13205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W319312701, W319612704)
}
__typedArg0 := W319312701
__typedArg1 := W319612704
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13206 := Call(__e, PrimFunc(symshen_4comb), W319712705, tmp13205)


ifres13201 = tmp13206


}

ifres13196 = ifres13201


}

W319112699 := ifres13196
_ = W319112699

tmp13218 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W319112699)


if True == tmp13218 {
tmp13209 := Call(__e, PrimFunc(sym_5e_6), V3190)


W319912707 := tmp13209
_ = W319912707

tmp13214 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W319912707)


var ifres13210 Obj

if True == tmp13214 {
tmp13211 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13210 = tmp13211


} else {
tmp13212 := Call(__e, PrimFunc(symshen_4in_1_6), W319912707)


W320012708 := tmp13212
_ = W320012708

tmp13213 := Call(__e, PrimFunc(symshen_4comb), W320012708, Nil)


ifres13210 = tmp13213


}

W319812706 := ifres13210
_ = W319812706

tmp13216 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W319812706)


if True == tmp13216 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W319812706)
return
}


} else {
__e.Return(W319112699)
return
}


}, 1)

tmp13219 := Call(__e, ns2_1set, symshen_4_5sides_6, tmp13194)


_ = tmp13219

tmp13220 := MakeNative(func(__e *ControlFlow) {
V3201 := __e.Get(1)
_ = V3201
tmp13232 := Call(__e, PrimFunc(symshen_4hds_a_2), V3201, symif)


var ifres13221 Obj

if True == tmp13232 {
tmp13222 := Call(__e, PrimFunc(symtail), V3201)


W320312710 := tmp13222
_ = W320312710

tmp13230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W320312710)
}
__typedArg0 := W320312710
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13223 Obj

if True == tmp13230 {
tmp13224 := Call(__e, PrimFunc(symhead), W320312710)


W320412711 := tmp13224
_ = W320412711

tmp13225 := Call(__e, PrimFunc(symtail), W320312710)


W320512712 := tmp13225
_ = W320512712

tmp13226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W320412711, Nil)
}
__typedArg0 := W320412711
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp13226)
}
__typedArg0 := symif
__typedArg1 := tmp13226
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13228 := Call(__e, PrimFunc(symshen_4comb), W320512712, tmp13227)


ifres13223 = tmp13228


} else {
tmp13229 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13223 = tmp13229


}

ifres13221 = ifres13223


} else {
tmp13231 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13221 = tmp13231


}

W320212709 := ifres13221
_ = W320212709

tmp13288 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W320212709)


if True == tmp13288 {
tmp13250 := Call(__e, PrimFunc(symshen_4hds_a_2), V3201, symlet)


var ifres13233 Obj

if True == tmp13250 {
tmp13234 := Call(__e, PrimFunc(symtail), V3201)


W320712714 := tmp13234
_ = W320712714

tmp13248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W320712714)
}
__typedArg0 := W320712714
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13235 Obj

if True == tmp13248 {
tmp13236 := Call(__e, PrimFunc(symhead), W320712714)


W320812715 := tmp13236
_ = W320812715

tmp13237 := Call(__e, PrimFunc(symtail), W320712714)


W320912716 := tmp13237
_ = W320912716

tmp13246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W320912716)
}
__typedArg0 := W320912716
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13238 Obj

if True == tmp13246 {
tmp13239 := Call(__e, PrimFunc(symhead), W320912716)


W321012717 := tmp13239
_ = W321012717

tmp13240 := Call(__e, PrimFunc(symtail), W320912716)


W321112718 := tmp13240
_ = W321112718

tmp13241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W321012717, Nil)
}
__typedArg0 := W321012717
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W320812715, tmp13241)
}
__typedArg0 := W320812715
__typedArg1 := tmp13241
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp13242)
}
__typedArg0 := symlet
__typedArg1 := tmp13242
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13244 := Call(__e, PrimFunc(symshen_4comb), W321112718, tmp13243)


ifres13238 = tmp13244


} else {
tmp13245 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13238 = tmp13245


}

ifres13235 = ifres13238


} else {
tmp13247 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13235 = tmp13247


}

ifres13233 = ifres13235


} else {
tmp13249 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13233 = tmp13249


}

W320612713 := ifres13233
_ = W320612713

tmp13286 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W320612713)


if True == tmp13286 {
tmp13265 := Call(__e, PrimFunc(symshen_4hds_a_2), V3201, symctxt)


var ifres13251 Obj

if True == tmp13265 {
tmp13252 := Call(__e, PrimFunc(symtail), V3201)


W321312720 := tmp13252
_ = W321312720

tmp13263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W321312720)
}
__typedArg0 := W321312720
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13253 Obj

if True == tmp13263 {
tmp13254 := Call(__e, PrimFunc(symhead), W321312720)


W321412721 := tmp13254
_ = W321412721

tmp13255 := Call(__e, PrimFunc(symtail), W321312720)


W321512722 := tmp13255
_ = W321512722

tmp13261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(W321412721)
}
__typedArg0 := W321412721
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres13256 Obj

if True == tmp13261 {
tmp13257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W321412721, Nil)
}
__typedArg0 := W321412721
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symctxt, tmp13257)
}
__typedArg0 := symctxt
__typedArg1 := tmp13257
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13259 := Call(__e, PrimFunc(symshen_4comb), W321512722, tmp13258)


ifres13256 = tmp13259


} else {
tmp13260 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13256 = tmp13260


}

ifres13253 = ifres13256


} else {
tmp13262 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13253 = tmp13262


}

ifres13251 = ifres13253


} else {
tmp13264 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13251 = tmp13264


}

W321212719 := ifres13251
_ = W321212719

tmp13284 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W321212719)


if True == tmp13284 {
tmp13280 := Call(__e, PrimFunc(symshen_4hds_a_2), V3201, symsqts)


var ifres13266 Obj

if True == tmp13280 {
tmp13267 := Call(__e, PrimFunc(symtail), V3201)


W321712724 := tmp13267
_ = W321712724

tmp13278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W321712724)
}
__typedArg0 := W321712724
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13268 Obj

if True == tmp13278 {
tmp13269 := Call(__e, PrimFunc(symhead), W321712724)


W321812725 := tmp13269
_ = W321812725

tmp13270 := Call(__e, PrimFunc(symtail), W321712724)


W321912726 := tmp13270
_ = W321912726

tmp13276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(W321812725)
}
__typedArg0 := W321812725
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres13271 Obj

if True == tmp13276 {
tmp13272 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W321812725, Nil)
}
__typedArg0 := W321812725
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsqts, tmp13272)
}
__typedArg0 := symsqts
__typedArg1 := tmp13272
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13274 := Call(__e, PrimFunc(symshen_4comb), W321912726, tmp13273)


ifres13271 = tmp13274


} else {
tmp13275 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13271 = tmp13275


}

ifres13268 = ifres13271


} else {
tmp13277 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13268 = tmp13277


}

ifres13266 = ifres13268


} else {
tmp13279 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13266 = tmp13279


}

W321612723 := ifres13266
_ = W321612723

tmp13282 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W321612723)


if True == tmp13282 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W321612723)
return
}


} else {
__e.Return(W321212719)
return
}


} else {
__e.Return(W320612713)
return
}


} else {
__e.Return(W320212709)
return
}


}, 1)

tmp13289 := Call(__e, ns2_1set, symshen_4_5side_6, tmp13220)


_ = tmp13289

tmp13290 := MakeNative(func(__e *ControlFlow) {
V3226 := __e.Get(1)
_ = V3226
V3227 := __e.Get(2)
_ = V3227
V3228 := __e.Get(3)
_ = V3228
tmp13320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3228)
}
__typedArg0 := V3228
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13307 Obj

if True == tmp13320 {
tmp13318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3228)
}
__typedArg0 := V3228
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13318)
}
__typedArg0 := Nil
__typedArg1 := tmp13318
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13309 Obj

if True == tmp13319 {
tmp13316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3228)
}
__typedArg0 := V3228
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13316)
}
__typedArg0 := tmp13316
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13311 Obj

if True == tmp13317 {
tmp13313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3228)
}
__typedArg0 := V3228
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13313)
}
__typedArg0 := tmp13313
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13314)
}
__typedArg0 := Nil
__typedArg1 := tmp13314
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13312 Obj

if True == tmp13315 {
ifres13312 = True


} else {
ifres13312 = False


}

ifres13311 = ifres13312


} else {
ifres13311 = False


}

var ifres13310 Obj

if True == ifres13311 {
ifres13310 = True


} else {
ifres13310 = False


}

ifres13309 = ifres13310


} else {
ifres13309 = False


}

var ifres13308 Obj

if True == ifres13309 {
ifres13308 = True


} else {
ifres13308 = False


}

ifres13307 = ifres13308


} else {
ifres13307 = False


}

if True == ifres13307 {
tmp13291 := Call(__e, PrimFunc(symgensym), symP)


W322912727 := tmp13291
_ = W322912727

tmp13292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3228)
}
__typedArg0 := V3228
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W322912727, Nil)
}
__typedArg0 := W322912727
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13292, tmp13293)
}
__typedArg0 := tmp13292
__typedArg1 := tmp13293
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

W323012728 := tmp13294
_ = W323012728

tmp13295 := Call(__e, PrimFunc(symshen_4coll_1formulae), V3227)


tmp13296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W322912727, Nil)
}
__typedArg0 := W322912727
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13295, tmp13296)
}
__typedArg0 := tmp13295
__typedArg1 := tmp13296
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

W323112729 := tmp13297
_ = W323112729

tmp13298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W323112729, Nil)
}
__typedArg0 := W323112729
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W323012728, Nil)
}
__typedArg0 := W323012728
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13298, tmp13299)
}
__typedArg0 := tmp13298
__typedArg1 := tmp13299
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3226, tmp13300)
}
__typedArg0 := V3226
__typedArg1 := tmp13300
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

W323212730 := tmp13301
_ = W323212730

tmp13302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3228, Nil)
}
__typedArg0 := V3228
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3227, tmp13302)
}
__typedArg0 := V3227
__typedArg1 := tmp13302
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3226, tmp13303)
}
__typedArg0 := V3226
__typedArg1 := tmp13303
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

W323312731 := tmp13304
_ = W323312731

tmp13305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W323212730, Nil)
}
__typedArg0 := W323212730
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W323312731, tmp13305)
}
__typedArg0 := W323312731
__typedArg1 := tmp13305
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.lr-rule"))
}
__typedArg0 := MakeString("implementation error in shen.lr-rule")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 3)

tmp13321 := Call(__e, ns2_1set, symshen_4lr_1rule, tmp13290)


_ = tmp13321

tmp13322 := MakeNative(func(__e *ControlFlow) {
V3236 := __e.Get(1)
_ = V3236
tmp13351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3236)
}
__typedArg0 := Nil
__typedArg1 := V3236
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13351 {
__e.Return(Nil)
return
} else {
tmp13349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3236)
}
__typedArg0 := V3236
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13329 Obj

if True == tmp13349 {
tmp13347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3236)
}
__typedArg0 := V3236
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13347)
}
__typedArg0 := tmp13347
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13331 Obj

if True == tmp13348 {
tmp13344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3236)
}
__typedArg0 := V3236
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13344)
}
__typedArg0 := tmp13344
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13345)
}
__typedArg0 := Nil
__typedArg1 := tmp13345
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13333 Obj

if True == tmp13346 {
tmp13341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3236)
}
__typedArg0 := V3236
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13341)
}
__typedArg0 := tmp13341
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13342)
}
__typedArg0 := tmp13342
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13335 Obj

if True == tmp13343 {
tmp13337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3236)
}
__typedArg0 := V3236
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13337)
}
__typedArg0 := tmp13337
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13338)
}
__typedArg0 := tmp13338
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13339)
}
__typedArg0 := Nil
__typedArg1 := tmp13339
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13336 Obj

if True == tmp13340 {
ifres13336 = True


} else {
ifres13336 = False


}

ifres13335 = ifres13336


} else {
ifres13335 = False


}

var ifres13334 Obj

if True == ifres13335 {
ifres13334 = True


} else {
ifres13334 = False


}

ifres13333 = ifres13334


} else {
ifres13333 = False


}

var ifres13332 Obj

if True == ifres13333 {
ifres13332 = True


} else {
ifres13332 = False


}

ifres13331 = ifres13332


} else {
ifres13331 = False


}

var ifres13330 Obj

if True == ifres13331 {
ifres13330 = True


} else {
ifres13330 = False


}

ifres13329 = ifres13330


} else {
ifres13329 = False


}

if True == ifres13329 {
tmp13323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3236)
}
__typedArg0 := V3236
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13323)
}
__typedArg0 := tmp13323
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13324)
}
__typedArg0 := tmp13324
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3236)
}
__typedArg0 := V3236
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13327 := Call(__e, PrimFunc(symshen_4coll_1formulae), tmp13326)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13325, tmp13327)
}
__typedArg0 := tmp13325
__typedArg1 := tmp13327
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.coll-formulae"))
}
__typedArg0 := MakeString("implementation error in shen.coll-formulae")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp13352 := Call(__e, ns2_1set, symshen_4coll_1formulae, tmp13322)


_ = tmp13352

tmp13353 := MakeNative(func(__e *ControlFlow) {
V3237 := __e.Get(1)
_ = V3237
tmp13364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3237)
}
__typedArg0 := V3237
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13354 Obj

if True == tmp13364 {
tmp13355 := Call(__e, PrimFunc(symhead), V3237)


W323912733 := tmp13355
_ = W323912733

tmp13356 := Call(__e, PrimFunc(symtail), V3237)


W324012734 := tmp13356
_ = W324012734

tmp13361 := Call(__e, PrimFunc(symshen_4key_1in_1sequent_1calculus_2), W323912733)


tmp13362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp13361)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp13361
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres13357 Obj

if True == tmp13362 {
tmp13358 := Call(__e, PrimFunc(symmacroexpand), W323912733)


tmp13359 := Call(__e, PrimFunc(symshen_4comb), W324012734, tmp13358)


ifres13357 = tmp13359


} else {
tmp13360 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13357 = tmp13360


}

ifres13354 = ifres13357


} else {
tmp13363 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13354 = tmp13363


}

W323812732 := ifres13354
_ = W323812732

tmp13366 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W323812732)


if True == tmp13366 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W323812732)
return
}


}, 1)

tmp13367 := Call(__e, ns2_1set, symshen_4_5expr_6, tmp13353)


_ = tmp13367

tmp13368 := MakeNative(func(__e *ControlFlow) {
V3241 := __e.Get(1)
_ = V3241
tmp13375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(","))
}
__typedArg0 := MakeString(",")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(":"))
}
__typedArg0 := MakeString(":")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13377, tmp13378)
}
__typedArg0 := tmp13377
__typedArg1 := tmp13378
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13376, tmp13379)
}
__typedArg0 := tmp13376
__typedArg1 := tmp13379
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13375, tmp13380)
}
__typedArg0 := tmp13375
__typedArg1 := tmp13380
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_6_6, tmp13381)
}
__typedArg0 := sym_6_6
__typedArg1 := tmp13381
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13383 := Call(__e, PrimFunc(symelement_2), V3241, tmp13382)


if True == tmp13383 {
__e.Return(True)
return
} else {
tmp13373 := Call(__e, PrimFunc(symshen_4sng_2), V3241)


var ifres13370 Obj

if True == tmp13373 {
ifres13370 = True


} else {
tmp13372 := Call(__e, PrimFunc(symshen_4dbl_2), V3241)


var ifres13371 Obj

if True == tmp13372 {
ifres13371 = True


} else {
ifres13371 = False


}

ifres13370 = ifres13371


}

if True == ifres13370 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp13384 := Call(__e, ns2_1set, symshen_4key_1in_1sequent_1calculus_2, tmp13368)


_ = tmp13384

tmp13385 := MakeNative(func(__e *ControlFlow) {
V3242 := __e.Get(1)
_ = V3242
tmp13386 := Call(__e, PrimFunc(symshen_4_5expr_6), V3242)


W324412736 := tmp13386
_ = W324412736

tmp13392 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W324412736)


var ifres13387 Obj

if True == tmp13392 {
tmp13388 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13387 = tmp13388


} else {
tmp13389 := Call(__e, PrimFunc(symshen_4_5_1out), W324412736)


W324512737 := tmp13389
_ = W324512737

tmp13390 := Call(__e, PrimFunc(symshen_4in_1_6), W324412736)


W324612738 := tmp13390
_ = W324612738

tmp13391 := Call(__e, PrimFunc(symshen_4comb), W324612738, W324512737)


ifres13387 = tmp13391


}

W324312735 := ifres13387
_ = W324312735

tmp13394 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W324312735)


if True == tmp13394 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W324312735)
return
}


}, 1)

tmp13395 := Call(__e, ns2_1set, symshen_4_5type_6, tmp13385)


_ = tmp13395

tmp13396 := MakeNative(func(__e *ControlFlow) {
V3247 := __e.Get(1)
_ = V3247
tmp13405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3247)
}
__typedArg0 := V3247
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13397 Obj

if True == tmp13405 {
tmp13398 := Call(__e, PrimFunc(symhead), V3247)


W324912740 := tmp13398
_ = W324912740

tmp13399 := Call(__e, PrimFunc(symtail), V3247)


W325012741 := tmp13399
_ = W325012741

tmp13403 := Call(__e, PrimFunc(symshen_4dbl_2), W324912740)


var ifres13400 Obj

if True == tmp13403 {
tmp13401 := Call(__e, PrimFunc(symshen_4comb), W325012741, W324912740)


ifres13400 = tmp13401


} else {
tmp13402 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13400 = tmp13402


}

ifres13397 = ifres13400


} else {
tmp13404 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13397 = tmp13404


}

W324812739 := ifres13397
_ = W324812739

tmp13407 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W324812739)


if True == tmp13407 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W324812739)
return
}


}, 1)

tmp13408 := Call(__e, ns2_1set, symshen_4_5dbl_6, tmp13396)


_ = tmp13408

tmp13409 := MakeNative(func(__e *ControlFlow) {
V3251 := __e.Get(1)
_ = V3251
tmp13418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3251)
}
__typedArg0 := V3251
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13410 Obj

if True == tmp13418 {
tmp13411 := Call(__e, PrimFunc(symhead), V3251)


W325312743 := tmp13411
_ = W325312743

tmp13412 := Call(__e, PrimFunc(symtail), V3251)


W325412744 := tmp13412
_ = W325412744

tmp13416 := Call(__e, PrimFunc(symshen_4sng_2), W325312743)


var ifres13413 Obj

if True == tmp13416 {
tmp13414 := Call(__e, PrimFunc(symshen_4comb), W325412744, W325312743)


ifres13413 = tmp13414


} else {
tmp13415 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13413 = tmp13415


}

ifres13410 = ifres13413


} else {
tmp13417 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres13410 = tmp13417


}

W325212742 := ifres13410
_ = W325212742

tmp13420 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W325212742)


if True == tmp13420 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W325212742)
return
}


}, 1)

tmp13421 := Call(__e, ns2_1set, symshen_4_5sng_6, tmp13409)


_ = tmp13421

tmp13422 := MakeNative(func(__e *ControlFlow) {
V3255 := __e.Get(1)
_ = V3255
tmp13427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V3255)
}
__typedArg0 := V3255
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp13427 {
tmp13424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3255)
}
__typedArg0 := V3255
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp13425 := Call(__e, PrimFunc(symshen_4sng_1h_2), tmp13424)


if True == tmp13425 {
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

tmp13428 := Call(__e, ns2_1set, symshen_4sng_2, tmp13422)


_ = tmp13428

tmp13429 := MakeNative(func(__e *ControlFlow) {
V3258 := __e.Get(1)
_ = V3258
tmp13438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("___"), V3258)
}
__typedArg0 := MakeString("___")
__typedArg1 := V3258
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13438 {
__e.Return(True)
return
} else {
tmp13436 := Call(__e, PrimFunc(symshen_4_7string_2), V3258)


var ifres13432 Obj

if True == tmp13436 {
tmp13434 := Call(__e, PrimFunc(symhdstr), V3258)


tmp13435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("_"), tmp13434)
}
__typedArg0 := MakeString("_")
__typedArg1 := tmp13434
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13433 Obj

if True == tmp13435 {
ifres13433 = True


} else {
ifres13433 = False


}

ifres13432 = ifres13433


} else {
ifres13432 = False


}

if True == ifres13432 {
__e.TailApply(PrimFunc(symshen_4sng_1h_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3258)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3258
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp13439 := Call(__e, ns2_1set, symshen_4sng_1h_2, tmp13429)


_ = tmp13439

tmp13440 := MakeNative(func(__e *ControlFlow) {
V3259 := __e.Get(1)
_ = V3259
tmp13445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(V3259)
}
__typedArg0 := V3259
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

if True == tmp13445 {
tmp13442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V3259)
}
__typedArg0 := V3259
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp13443 := Call(__e, PrimFunc(symshen_4dbl_1h_2), tmp13442)


if True == tmp13443 {
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

tmp13446 := Call(__e, ns2_1set, symshen_4dbl_2, tmp13440)


_ = tmp13446

tmp13447 := MakeNative(func(__e *ControlFlow) {
V3262 := __e.Get(1)
_ = V3262
tmp13456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("==="), V3262)
}
__typedArg0 := MakeString("===")
__typedArg1 := V3262
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13456 {
__e.Return(True)
return
} else {
tmp13454 := Call(__e, PrimFunc(symshen_4_7string_2), V3262)


var ifres13450 Obj

if True == tmp13454 {
tmp13452 := Call(__e, PrimFunc(symhdstr), V3262)


tmp13453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString("="), tmp13452)
}
__typedArg0 := MakeString("=")
__typedArg1 := tmp13452
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13451 Obj

if True == tmp13453 {
ifres13451 = True


} else {
ifres13451 = False


}

ifres13450 = ifres13451


} else {
ifres13450 = False


}

if True == ifres13450 {
__e.TailApply(PrimFunc(symshen_4dbl_1h_2), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtlstr) {
__typedS0, __typedOK0 := TypedString(V3262)
if __typedOK0 && HasCanonicalPrimitiveBinding(symtlstr) {
return TypedMaterializeString(TypedStringTailValue(__typedS0))
}}
__typedArg0 := V3262
return Call(__e, PrimFunc(symtlstr), __typedArg0)
})())
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp13457 := Call(__e, ns2_1set, symshen_4dbl_1h_2, tmp13447)


_ = tmp13457

tmp13458 := MakeNative(func(__e *ControlFlow) {
V3263 := __e.Get(1)
_ = V3263
V3264 := __e.Get(2)
_ = V3264
tmp13459 := MakeNative(func(__e *ControlFlow) {
Z3266 := __e.Get(1)
_ = Z3266
__e.TailApply(PrimFunc(symshen_4rule_1_6clause), Z3266)
return
}, 1)

tmp13460 := Call(__e, PrimFunc(symmapcan), tmp13459, V3264)


W326512745 := tmp13460
_ = W326512745

tmp13461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3263, W326512745)
}
__typedArg0 := V3263
__typedArg1 := W326512745
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp13461)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp13461
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

W326712746 := tmp13462
_ = W326712746

__e.TailApply(PrimFunc(symeval), W326712746)
return


}, 2)

tmp13463 := Call(__e, ns2_1set, symshen_4rules_1_6prolog, tmp13458)


_ = tmp13463

tmp13464 := MakeNative(func(__e *ControlFlow) {
V3268 := __e.Get(1)
_ = V3268
tmp13524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13488 Obj

if True == tmp13524 {
tmp13522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13522)
}
__typedArg0 := tmp13522
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13490 Obj

if True == tmp13523 {
tmp13519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13519)
}
__typedArg0 := tmp13519
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13520)
}
__typedArg0 := tmp13520
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13492 Obj

if True == tmp13521 {
tmp13515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13515)
}
__typedArg0 := tmp13515
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13516)
}
__typedArg0 := tmp13516
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13517)
}
__typedArg0 := tmp13517
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13494 Obj

if True == tmp13518 {
tmp13510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13510)
}
__typedArg0 := tmp13510
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13511)
}
__typedArg0 := tmp13511
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13512)
}
__typedArg0 := tmp13512
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13513)
}
__typedArg0 := tmp13513
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13496 Obj

if True == tmp13514 {
tmp13504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13504)
}
__typedArg0 := tmp13504
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13505)
}
__typedArg0 := tmp13505
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13506)
}
__typedArg0 := tmp13506
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13507)
}
__typedArg0 := tmp13507
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13508)
}
__typedArg0 := Nil
__typedArg1 := tmp13508
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13498 Obj

if True == tmp13509 {
tmp13500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13500)
}
__typedArg0 := tmp13500
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13501)
}
__typedArg0 := tmp13501
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13502)
}
__typedArg0 := Nil
__typedArg1 := tmp13502
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13499 Obj

if True == tmp13503 {
ifres13499 = True


} else {
ifres13499 = False


}

ifres13498 = ifres13499


} else {
ifres13498 = False


}

var ifres13497 Obj

if True == ifres13498 {
ifres13497 = True


} else {
ifres13497 = False


}

ifres13496 = ifres13497


} else {
ifres13496 = False


}

var ifres13495 Obj

if True == ifres13496 {
ifres13495 = True


} else {
ifres13495 = False


}

ifres13494 = ifres13495


} else {
ifres13494 = False


}

var ifres13493 Obj

if True == ifres13494 {
ifres13493 = True


} else {
ifres13493 = False


}

ifres13492 = ifres13493


} else {
ifres13492 = False


}

var ifres13491 Obj

if True == ifres13492 {
ifres13491 = True


} else {
ifres13491 = False


}

ifres13490 = ifres13491


} else {
ifres13490 = False


}

var ifres13489 Obj

if True == ifres13490 {
ifres13489 = True


} else {
ifres13489 = False


}

ifres13488 = ifres13489


} else {
ifres13488 = False


}

if True == ifres13488 {
tmp13465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13465)
}
__typedArg0 := tmp13465
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13466)
}
__typedArg0 := tmp13466
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13467)
}
__typedArg0 := tmp13467
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13468)
}
__typedArg0 := tmp13468
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13470 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp13469)


W326912747 := tmp13470
_ = W326912747

tmp13471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13471)
}
__typedArg0 := tmp13471
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13472)
}
__typedArg0 := tmp13472
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13473)
}
__typedArg0 := tmp13473
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13474)
}
__typedArg0 := tmp13474
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13476 := Call(__e, PrimFunc(symshen_4rule_1_6head), tmp13475)


tmp13477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13479)
}
__typedArg0 := tmp13479
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3268)
}
__typedArg0 := V3268
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13481)
}
__typedArg0 := tmp13481
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13482)
}
__typedArg0 := tmp13482
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13483)
}
__typedArg0 := tmp13483
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13485 := Call(__e, PrimFunc(symshen_4rule_1_6body), W326912747, symAssumptions, tmp13478, tmp13480, tmp13484)


tmp13486 := Call(__e, PrimFunc(symappend), tmp13477, tmp13485)


__e.TailApply(PrimFunc(symappend), tmp13476, tmp13486)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.rule->clause"))
}
__typedArg0 := MakeString("partial function shen.rule->clause")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp13525 := Call(__e, ns2_1set, symshen_4rule_1_6clause, tmp13464)


_ = tmp13525

tmp13526 := MakeNative(func(__e *ControlFlow) {
V3270 := __e.Get(1)
_ = V3270
tmp13527 := Call(__e, PrimFunc(symshen_4macro_1_8ch), V3270)


tmp13528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symAssumptions, Nil)
}
__typedArg0 := symAssumptions
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13527, tmp13528)
}
__typedArg0 := tmp13527
__typedArg1 := tmp13528
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp13529 := Call(__e, ns2_1set, symshen_4rule_1_6head, tmp13526)


_ = tmp13529

tmp13530 := MakeNative(func(__e *ControlFlow) {
V3271 := __e.Get(1)
_ = V3271
tmp13531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3271, Nil)
}
__typedArg0 := V3271
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_8ch, tmp13531)
}
__typedArg0 := symshen_4_8ch
__typedArg1 := tmp13531
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp13532 := Call(__e, ns2_1set, symshen_4macro_1_8ch, tmp13530)


_ = tmp13532

tmp13533 := MakeNative(func(__e *ControlFlow) {
V3272 := __e.Get(1)
_ = V3272
tmp13534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3272, Nil)
}
__typedArg0 := V3272
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4_8c, tmp13534)
}
__typedArg0 := symshen_4_8c
__typedArg1 := tmp13534
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp13535 := Call(__e, ns2_1set, symshen_4macro_1_8c, tmp13533)


_ = tmp13535

tmp13536 := MakeNative(func(__e *ControlFlow) {
V3273 := __e.Get(1)
_ = V3273
V3274 := __e.Get(2)
_ = V3274
V3275 := __e.Get(3)
_ = V3275
V3276 := __e.Get(4)
_ = V3276
V3277 := __e.Get(5)
_ = V3277
tmp13566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3277)
}
__typedArg0 := Nil
__typedArg1 := V3277
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13566 {
__e.TailApply(PrimFunc(symshen_4side_1conditions_1_6goals), Nil, V3273, V3274, V3275, V3276)
return
} else {
tmp13564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3276)
}
__typedArg0 := Nil
__typedArg1 := V3276
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13557 Obj

if True == tmp13564 {
tmp13563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3277)
}
__typedArg0 := V3277
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13559 Obj

if True == tmp13563 {
tmp13561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3277)
}
__typedArg0 := V3277
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13561)
}
__typedArg0 := Nil
__typedArg1 := tmp13561
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13560 Obj

if True == tmp13562 {
ifres13560 = True


} else {
ifres13560 = False


}

ifres13559 = ifres13560


} else {
ifres13559 = False


}

var ifres13558 Obj

if True == ifres13559 {
ifres13558 = True


} else {
ifres13558 = False


}

ifres13557 = ifres13558


} else {
ifres13557 = False


}

if True == ifres13557 {
tmp13537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3277)
}
__typedArg0 := V3277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13538 := Call(__e, PrimFunc(symshen_4passive_1variables), tmp13537, V3273)


W327812748 := tmp13538
_ = W327812748

tmp13539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3277)
}
__typedArg0 := V3277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13540 := Call(__e, PrimFunc(symshen_4remove_1bystanders), V3273, tmp13539)


W327912749 := tmp13540
_ = W327912749

tmp13541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3277)
}
__typedArg0 := V3277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13542 := Call(__e, PrimFunc(symshen_4specialise_1member), tmp13541, V3274, W327912749, W327812748)


tmp13543 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), Nil, V3273, V3274, V3275, Nil)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13542, tmp13543)
}
__typedArg0 := tmp13542
__typedArg1 := tmp13543
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3277)
}
__typedArg0 := V3277
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13555 {
tmp13544 := Call(__e, PrimFunc(symgensym), symNewAssumptions)


W328012750 := tmp13544
_ = W328012750

tmp13545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3277)
}
__typedArg0 := V3277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13546 := Call(__e, PrimFunc(symshen_4passive_1variables), tmp13545, V3273)


W328112751 := tmp13546
_ = W328112751

tmp13547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3277)
}
__typedArg0 := V3277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13548 := Call(__e, PrimFunc(symshen_4remove_1bystanders), V3273, tmp13547)


W328212752 := tmp13548
_ = W328212752

tmp13549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3277)
}
__typedArg0 := V3277
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13550 := Call(__e, PrimFunc(symshen_4specialise_1consume), tmp13549, V3274, W328212752, W328112751, W328012750)


tmp13551 := Call(__e, PrimFunc(symappend), V3273, W328112751)


tmp13552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3277)
}
__typedArg0 := V3277
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13553 := Call(__e, PrimFunc(symshen_4rule_1_6body), tmp13551, W328012750, V3275, V3276, tmp13552)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13550, tmp13553)
}
__typedArg0 := tmp13550
__typedArg1 := tmp13553
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.rule->body"))
}
__typedArg0 := MakeString("partial function shen.rule->body")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 5)

tmp13567 := Call(__e, ns2_1set, symshen_4rule_1_6body, tmp13536)


_ = tmp13567

tmp13568 := MakeNative(func(__e *ControlFlow) {
V3283 := __e.Get(1)
_ = V3283
V3284 := __e.Get(2)
_ = V3284
V3285 := __e.Get(3)
_ = V3285
V3286 := __e.Get(4)
_ = V3286
tmp13569 := Call(__e, PrimFunc(symgensym), symshen_4member)


W328712753 := tmp13569
_ = W328712753

tmp13570 := Call(__e, PrimFunc(symshen_4member_1clause), W328712753, V3283, V3285, V3286)


W328812754 := tmp13570
_ = W328812754

tmp13571 := Call(__e, PrimFunc(symappend), V3285, V3286)


tmp13572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3284, tmp13571)
}
__typedArg0 := V3284
__typedArg1 := tmp13571
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W328712753, tmp13572)
}
__typedArg0 := W328712753
__typedArg1 := tmp13572
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 4)

tmp13573 := Call(__e, ns2_1set, symshen_4specialise_1member, tmp13568)


_ = tmp13573

tmp13574 := MakeNative(func(__e *ControlFlow) {
V3291 := __e.Get(1)
_ = V3291
V3292 := __e.Get(2)
_ = V3292
tmp13588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3291)
}
__typedArg0 := Nil
__typedArg1 := V3291
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13588 {
__e.Return(Nil)
return
} else {
tmp13586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3291)
}
__typedArg0 := V3291
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13582 Obj

if True == tmp13586 {
tmp13584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3291)
}
__typedArg0 := V3291
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13585 := Call(__e, PrimFunc(symshen_4occurs_1check_2), tmp13584, V3292)


var ifres13583 Obj

if True == tmp13585 {
ifres13583 = True


} else {
ifres13583 = False


}

ifres13582 = ifres13583


} else {
ifres13582 = False


}

if True == ifres13582 {
tmp13575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3291)
}
__typedArg0 := V3291
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3291)
}
__typedArg0 := V3291
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13577 := Call(__e, PrimFunc(symshen_4remove_1bystanders), tmp13576, V3292)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13575, tmp13577)
}
__typedArg0 := tmp13575
__typedArg1 := tmp13577
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3291)
}
__typedArg0 := V3291
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13580 {
tmp13578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3291)
}
__typedArg0 := V3291
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4remove_1bystanders), tmp13578, V3292)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.remove-bystanders"))
}
__typedArg0 := MakeString("partial function shen.remove-bystanders")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 2)

tmp13589 := Call(__e, ns2_1set, symshen_4remove_1bystanders, tmp13574)


_ = tmp13589

tmp13590 := MakeNative(func(__e *ControlFlow) {
V3293 := __e.Get(1)
_ = V3293
V3294 := __e.Get(2)
_ = V3294
V3295 := __e.Get(3)
_ = V3295
V3296 := __e.Get(4)
_ = V3296
tmp13591 := Call(__e, PrimFunc(symlength), V3296)


tmp13592 := Call(__e, PrimFunc(symshen_4nvars), tmp13591)


W329712755 := tmp13592
_ = W329712755

tmp13593 := Call(__e, PrimFunc(symshen_4macro_1_8ch), V3294)


tmp13594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym__, Nil)
}
__typedArg0 := sym__
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13593, tmp13594)
}
__typedArg0 := tmp13593
__typedArg1 := tmp13594
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13595)
}
__typedArg0 := symcons
__typedArg1 := tmp13595
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13596, Nil)
}
__typedArg0 := tmp13596
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13597)
}
__typedArg0 := sym_1
__typedArg1 := tmp13597
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13598, Nil)
}
__typedArg0 := tmp13598
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13601 := Call(__e, PrimFunc(symshen_4passive_1bind), V3296, W329712755)


tmp13602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13602, Nil)
}
__typedArg0 := tmp13602
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13604 := Call(__e, PrimFunc(symappend), tmp13601, tmp13603)


tmp13605 := Call(__e, PrimFunc(symappend), tmp13600, tmp13604)


tmp13606 := Call(__e, PrimFunc(symappend), W329712755, tmp13605)


tmp13607 := Call(__e, PrimFunc(symappend), V3295, tmp13606)


tmp13608 := Call(__e, PrimFunc(symappend), tmp13599, tmp13607)


W329812756 := tmp13608
_ = W329812756

tmp13609 := Call(__e, PrimFunc(symgensym), symHypotheses)


W330012758 := tmp13609
_ = W330012758

tmp13610 := Call(__e, PrimFunc(symappend), V3295, V3296)


W330112759 := tmp13610
_ = W330112759

tmp13611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W330012758, Nil)
}
__typedArg0 := W330012758
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym__, tmp13611)
}
__typedArg0 := sym__
__typedArg1 := tmp13611
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13612)
}
__typedArg0 := symcons
__typedArg1 := tmp13612
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13613, Nil)
}
__typedArg0 := tmp13613
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13614)
}
__typedArg0 := sym_1
__typedArg1 := tmp13614
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13615, Nil)
}
__typedArg0 := tmp13615
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13617 := Call(__e, PrimFunc(symappend), tmp13616, W330112759)


W330212760 := tmp13617
_ = W330212760

tmp13618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W330012758, W330112759)
}
__typedArg0 := W330012758
__typedArg1 := W330112759
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3293, tmp13618)
}
__typedArg0 := V3293
__typedArg1 := tmp13618
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13619, Nil)
}
__typedArg0 := tmp13619
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

W330312761 := tmp13620
_ = W330312761

tmp13621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13622, Nil)
}
__typedArg0 := tmp13622
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13624 := Call(__e, PrimFunc(symappend), W330312761, tmp13623)


tmp13625 := Call(__e, PrimFunc(symappend), tmp13621, tmp13624)


tmp13626 := Call(__e, PrimFunc(symappend), W330212760, tmp13625)


W329912757 := tmp13626
_ = W329912757

tmp13627 := Call(__e, PrimFunc(symappend), W329812756, W329912757)


tmp13628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3293, tmp13627)
}
__typedArg0 := V3293
__typedArg1 := tmp13627
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp13628)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp13628
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

W330412762 := tmp13629
_ = W330412762

__e.TailApply(PrimFunc(symeval), W330412762)
return


}, 4)

tmp13630 := Call(__e, ns2_1set, symshen_4member_1clause, tmp13590)


_ = tmp13630

tmp13631 := MakeNative(func(__e *ControlFlow) {
V3305 := __e.Get(1)
_ = V3305
tmp13636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V3305)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V3305
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13636 {
__e.Return(Nil)
return
} else {
tmp13632 := Call(__e, PrimFunc(symgensym), symNewV)


tmp13634 := Call(__e, PrimFunc(symshen_4nvars), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V3305)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V3305
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13632, tmp13634)
}
__typedArg0 := tmp13632
__typedArg1 := tmp13634
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp13637 := Call(__e, ns2_1set, symshen_4nvars, tmp13631)


_ = tmp13637

tmp13638 := MakeNative(func(__e *ControlFlow) {
V3306 := __e.Get(1)
_ = V3306
V3307 := __e.Get(2)
_ = V3307
tmp13656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3306)
}
__typedArg0 := Nil
__typedArg1 := V3306
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13653 Obj

if True == tmp13656 {
tmp13655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3307)
}
__typedArg0 := Nil
__typedArg1 := V3307
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13654 Obj

if True == tmp13655 {
ifres13654 = True


} else {
ifres13654 = False


}

ifres13653 = ifres13654


} else {
ifres13653 = False


}

if True == ifres13653 {
__e.Return(Nil)
return
} else {
tmp13651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3306)
}
__typedArg0 := V3306
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13648 Obj

if True == tmp13651 {
tmp13650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3307)
}
__typedArg0 := V3307
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13649 Obj

if True == tmp13650 {
ifres13649 = True


} else {
ifres13649 = False


}

ifres13648 = ifres13649


} else {
ifres13648 = False


}

if True == ifres13648 {
tmp13639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3307)
}
__typedArg0 := V3307
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3306)
}
__typedArg0 := V3306
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13640, Nil)
}
__typedArg0 := tmp13640
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13639, tmp13641)
}
__typedArg0 := tmp13639
__typedArg1 := tmp13641
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13642)
}
__typedArg0 := symbind
__typedArg1 := tmp13642
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3306)
}
__typedArg0 := V3306
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3307)
}
__typedArg0 := V3307
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13646 := Call(__e, PrimFunc(symshen_4passive_1bind), tmp13644, tmp13645)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13643, tmp13646)
}
__typedArg0 := tmp13643
__typedArg1 := tmp13646
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.passive-bind"))
}
__typedArg0 := MakeString("partial function shen.passive-bind")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp13657 := Call(__e, ns2_1set, symshen_4passive_1bind, tmp13638)


_ = tmp13657

tmp13658 := MakeNative(func(__e *ControlFlow) {
V3308 := __e.Get(1)
_ = V3308
V3309 := __e.Get(2)
_ = V3309
V3310 := __e.Get(3)
_ = V3310
V3311 := __e.Get(4)
_ = V3311
V3312 := __e.Get(5)
_ = V3312
tmp13659 := Call(__e, PrimFunc(symgensym), symshen_4consume)


W331312763 := tmp13659
_ = W331312763

tmp13660 := Call(__e, PrimFunc(symshen_4consume_1clause), W331312763, V3308, V3310, V3311, V3312)


W331412764 := tmp13660
_ = W331412764

tmp13661 := Call(__e, PrimFunc(symappend), V3310, V3311)


tmp13662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3312, tmp13661)
}
__typedArg0 := V3312
__typedArg1 := tmp13661
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3309, tmp13662)
}
__typedArg0 := V3309
__typedArg1 := tmp13662
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W331312763, tmp13663)
}
__typedArg0 := W331312763
__typedArg1 := tmp13663
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 5)

tmp13664 := Call(__e, ns2_1set, symshen_4specialise_1consume, tmp13658)


_ = tmp13664

tmp13665 := MakeNative(func(__e *ControlFlow) {
V3315 := __e.Get(1)
_ = V3315
V3316 := __e.Get(2)
_ = V3316
V3317 := __e.Get(3)
_ = V3317
V3318 := __e.Get(4)
_ = V3318
V3319 := __e.Get(5)
_ = V3319
tmp13666 := Call(__e, PrimFunc(symlength), V3318)


tmp13667 := Call(__e, PrimFunc(symshen_4nvars), tmp13666)


W332012765 := tmp13667
_ = W332012765

tmp13668 := Call(__e, PrimFunc(symgensym), symAssumption)


W332112766 := tmp13668
_ = W332112766

tmp13669 := Call(__e, PrimFunc(symshen_4macro_1_8ch), V3316)


tmp13670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W332112766, Nil)
}
__typedArg0 := W332112766
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13669, tmp13670)
}
__typedArg0 := tmp13669
__typedArg1 := tmp13670
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13671)
}
__typedArg0 := symcons
__typedArg1 := tmp13671
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13672, Nil)
}
__typedArg0 := tmp13672
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13673)
}
__typedArg0 := sym_1
__typedArg1 := tmp13673
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13676 := Call(__e, PrimFunc(symshen_4passive_1bind), V3318, W332012765)


tmp13677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W332112766, Nil)
}
__typedArg0 := W332112766
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3319, tmp13677)
}
__typedArg0 := V3319
__typedArg1 := tmp13677
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13678)
}
__typedArg0 := symbind
__typedArg1 := tmp13678
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13680, Nil)
}
__typedArg0 := tmp13680
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13679, tmp13681)
}
__typedArg0 := tmp13679
__typedArg1 := tmp13681
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13683 := Call(__e, PrimFunc(symappend), tmp13676, tmp13682)


tmp13684 := Call(__e, PrimFunc(symappend), tmp13675, tmp13683)


tmp13685 := Call(__e, PrimFunc(symappend), W332012765, tmp13684)


tmp13686 := Call(__e, PrimFunc(symappend), V3317, tmp13685)


tmp13687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3319, tmp13686)
}
__typedArg0 := V3319
__typedArg1 := tmp13686
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13674, tmp13687)
}
__typedArg0 := tmp13674
__typedArg1 := tmp13687
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

W332212767 := tmp13688
_ = W332212767

tmp13689 := Call(__e, PrimFunc(symgensym), symHypotheses)


W332412769 := tmp13689
_ = W332412769

tmp13690 := Call(__e, PrimFunc(symappend), V3317, V3318)


W332512770 := tmp13690
_ = W332512770

tmp13691 := Call(__e, PrimFunc(symgensym), symAssumptions)


W332612771 := tmp13691
_ = W332612771

tmp13692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W332412769, Nil)
}
__typedArg0 := W332412769
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W332112766, tmp13692)
}
__typedArg0 := W332112766
__typedArg1 := tmp13692
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13693)
}
__typedArg0 := symcons
__typedArg1 := tmp13693
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13694, Nil)
}
__typedArg0 := tmp13694
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1, tmp13695)
}
__typedArg0 := sym_1
__typedArg1 := tmp13695
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3319, Nil)
}
__typedArg0 := V3319
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W332612771, tmp13697)
}
__typedArg0 := W332612771
__typedArg1 := tmp13697
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13698)
}
__typedArg0 := symcons
__typedArg1 := tmp13698
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13699, W332512770)
}
__typedArg0 := tmp13699
__typedArg1 := W332512770
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13696, tmp13700)
}
__typedArg0 := tmp13696
__typedArg1 := tmp13700
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

W332712772 := tmp13701
_ = W332712772

tmp13702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W332112766, Nil)
}
__typedArg0 := W332112766
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W332612771, tmp13702)
}
__typedArg0 := W332612771
__typedArg1 := tmp13702
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13703)
}
__typedArg0 := symbind
__typedArg1 := tmp13703
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3319, W332512770)
}
__typedArg0 := V3319
__typedArg1 := W332512770
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W332412769, tmp13705)
}
__typedArg0 := W332412769
__typedArg1 := tmp13705
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3315, tmp13706)
}
__typedArg0 := V3315
__typedArg1 := tmp13706
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13707, Nil)
}
__typedArg0 := tmp13707
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13704, tmp13708)
}
__typedArg0 := tmp13704
__typedArg1 := tmp13708
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

W332812773 := tmp13709
_ = W332812773

tmp13710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1_1, Nil)
}
__typedArg0 := sym_5_1_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

tmp13712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13711, Nil)
}
__typedArg0 := tmp13711
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13713 := Call(__e, PrimFunc(symappend), W332812773, tmp13712)


tmp13714 := Call(__e, PrimFunc(symappend), tmp13710, tmp13713)


tmp13715 := Call(__e, PrimFunc(symappend), W332712772, tmp13714)


W332312768 := tmp13715
_ = W332312768

tmp13716 := Call(__e, PrimFunc(symappend), W332212767, W332312768)


tmp13717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3315, tmp13716)
}
__typedArg0 := V3315
__typedArg1 := tmp13716
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefprolog, tmp13717)
}
__typedArg0 := symdefprolog
__typedArg1 := tmp13717
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

W332912774 := tmp13718
_ = W332912774

__e.TailApply(PrimFunc(symeval), W332912774)
return


}, 5)

tmp13719 := Call(__e, ns2_1set, symshen_4consume_1clause, tmp13665)


_ = tmp13719

tmp13720 := MakeNative(func(__e *ControlFlow) {
V3330 := __e.Get(1)
_ = V3330
V3331 := __e.Get(2)
_ = V3331
tmp13721 := Call(__e, PrimFunc(symshen_4extract_1vars), V3330)


__e.TailApply(PrimFunc(symdifference), tmp13721, V3331)
return


}, 2)

tmp13722 := Call(__e, ns2_1set, symshen_4passive_1variables, tmp13720)


_ = tmp13722

tmp13723 := MakeNative(func(__e *ControlFlow) {
V3336 := __e.Get(1)
_ = V3336
V3337 := __e.Get(2)
_ = V3337
V3338 := __e.Get(3)
_ = V3338
V3339 := __e.Get(4)
_ = V3339
V3340 := __e.Get(5)
_ = V3340
tmp13874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3339)
}
__typedArg0 := Nil
__typedArg1 := V3339
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13874 {
__e.TailApply(PrimFunc(symshen_4premises_1_6goals), V3336, V3338, V3340)
return
} else {
tmp13872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13852 Obj

if True == tmp13872 {
tmp13870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13870)
}
__typedArg0 := tmp13870
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13854 Obj

if True == tmp13871 {
tmp13867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13867)
}
__typedArg0 := tmp13867
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symif, tmp13868)
}
__typedArg0 := symif
__typedArg1 := tmp13868
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13856 Obj

if True == tmp13869 {
tmp13864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13864)
}
__typedArg0 := tmp13864
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13865)
}
__typedArg0 := tmp13865
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13858 Obj

if True == tmp13866 {
tmp13860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13860)
}
__typedArg0 := tmp13860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13861)
}
__typedArg0 := tmp13861
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13862)
}
__typedArg0 := Nil
__typedArg1 := tmp13862
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13859 Obj

if True == tmp13863 {
ifres13859 = True


} else {
ifres13859 = False


}

ifres13858 = ifres13859


} else {
ifres13858 = False


}

var ifres13857 Obj

if True == ifres13858 {
ifres13857 = True


} else {
ifres13857 = False


}

ifres13856 = ifres13857


} else {
ifres13856 = False


}

var ifres13855 Obj

if True == ifres13856 {
ifres13855 = True


} else {
ifres13855 = False


}

ifres13854 = ifres13855


} else {
ifres13854 = False


}

var ifres13853 Obj

if True == ifres13854 {
ifres13853 = True


} else {
ifres13853 = False


}

ifres13852 = ifres13853


} else {
ifres13852 = False


}

if True == ifres13852 {
tmp13724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13724)
}
__typedArg0 := tmp13724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhen, tmp13725)
}
__typedArg0 := symwhen
__typedArg1 := tmp13725
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13728 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), V3336, V3337, V3338, tmp13727, V3340)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13726, tmp13728)
}
__typedArg0 := tmp13726
__typedArg1 := tmp13728
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13823 Obj

if True == tmp13850 {
tmp13848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13848)
}
__typedArg0 := tmp13848
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13825 Obj

if True == tmp13849 {
tmp13845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13845)
}
__typedArg0 := tmp13845
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp13846)
}
__typedArg0 := symlet
__typedArg1 := tmp13846
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13827 Obj

if True == tmp13847 {
tmp13842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13842)
}
__typedArg0 := tmp13842
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13843)
}
__typedArg0 := tmp13843
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13829 Obj

if True == tmp13844 {
tmp13838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13838)
}
__typedArg0 := tmp13838
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13839)
}
__typedArg0 := tmp13839
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13840)
}
__typedArg0 := tmp13840
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13831 Obj

if True == tmp13841 {
tmp13833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13833)
}
__typedArg0 := tmp13833
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13834)
}
__typedArg0 := tmp13834
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13835)
}
__typedArg0 := tmp13835
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13836)
}
__typedArg0 := Nil
__typedArg1 := tmp13836
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13832 Obj

if True == tmp13837 {
ifres13832 = True


} else {
ifres13832 = False


}

ifres13831 = ifres13832


} else {
ifres13831 = False


}

var ifres13830 Obj

if True == ifres13831 {
ifres13830 = True


} else {
ifres13830 = False


}

ifres13829 = ifres13830


} else {
ifres13829 = False


}

var ifres13828 Obj

if True == ifres13829 {
ifres13828 = True


} else {
ifres13828 = False


}

ifres13827 = ifres13828


} else {
ifres13827 = False


}

var ifres13826 Obj

if True == ifres13827 {
ifres13826 = True


} else {
ifres13826 = False


}

ifres13825 = ifres13826


} else {
ifres13825 = False


}

var ifres13824 Obj

if True == ifres13825 {
ifres13824 = True


} else {
ifres13824 = False


}

ifres13823 = ifres13824


} else {
ifres13823 = False


}

if True == ifres13823 {
tmp13744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13744)
}
__typedArg0 := tmp13744
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13745)
}
__typedArg0 := tmp13745
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13747 := Call(__e, PrimFunc(symelement_2), tmp13746, V3337)


if True == tmp13747 {
tmp13729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13729)
}
__typedArg0 := tmp13729
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symis_b, tmp13730)
}
__typedArg0 := symis_b
__typedArg1 := tmp13730
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13733 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), V3336, V3337, V3338, tmp13732, V3340)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13731, tmp13733)
}
__typedArg0 := tmp13731
__typedArg1 := tmp13733
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13734)
}
__typedArg0 := tmp13734
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13735)
}
__typedArg0 := symbind
__typedArg1 := tmp13735
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13737)
}
__typedArg0 := tmp13737
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13738)
}
__typedArg0 := tmp13738
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13739, V3337)
}
__typedArg0 := tmp13739
__typedArg1 := V3337
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13742 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), V3336, tmp13740, V3338, tmp13741, V3340)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13736, tmp13742)
}
__typedArg0 := tmp13736
__typedArg1 := tmp13742
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
tmp13821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13801 Obj

if True == tmp13821 {
tmp13819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13819)
}
__typedArg0 := tmp13819
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13803 Obj

if True == tmp13820 {
tmp13816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13816)
}
__typedArg0 := tmp13816
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symctxt, tmp13817)
}
__typedArg0 := symctxt
__typedArg1 := tmp13817
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13805 Obj

if True == tmp13818 {
tmp13813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13813)
}
__typedArg0 := tmp13813
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13814)
}
__typedArg0 := tmp13814
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13807 Obj

if True == tmp13815 {
tmp13809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13809)
}
__typedArg0 := tmp13809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13810)
}
__typedArg0 := tmp13810
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13811)
}
__typedArg0 := Nil
__typedArg1 := tmp13811
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13808 Obj

if True == tmp13812 {
ifres13808 = True


} else {
ifres13808 = False


}

ifres13807 = ifres13808


} else {
ifres13807 = False


}

var ifres13806 Obj

if True == ifres13807 {
ifres13806 = True


} else {
ifres13806 = False


}

ifres13805 = ifres13806


} else {
ifres13805 = False


}

var ifres13804 Obj

if True == ifres13805 {
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

if True == ifres13801 {
tmp13773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13773)
}
__typedArg0 := tmp13773
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13774)
}
__typedArg0 := tmp13774
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13776 := Call(__e, PrimFunc(symelement_2), tmp13775, V3337)


if True == tmp13776 {
tmp13748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13748)
}
__typedArg0 := tmp13748
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13749)
}
__typedArg0 := tmp13749
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13750, V3336)
}
__typedArg0 := tmp13750
__typedArg1 := V3336
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4side_1conditions_1_6goals), tmp13751, V3337, V3338, tmp13752, V3340)
return


} else {
tmp13753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13753)
}
__typedArg0 := tmp13753
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13754)
}
__typedArg0 := tmp13754
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V3338, Nil)
}
__typedArg0 := V3338
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13755, tmp13756)
}
__typedArg0 := tmp13755
__typedArg1 := tmp13756
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symbind, tmp13757)
}
__typedArg0 := symbind
__typedArg1 := tmp13757
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13759)
}
__typedArg0 := tmp13759
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13760)
}
__typedArg0 := tmp13760
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13761, V3336)
}
__typedArg0 := tmp13761
__typedArg1 := V3336
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13763)
}
__typedArg0 := tmp13763
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13764)
}
__typedArg0 := tmp13764
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13765, V3337)
}
__typedArg0 := tmp13765
__typedArg1 := V3337
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13767)
}
__typedArg0 := tmp13767
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13768)
}
__typedArg0 := tmp13768
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13771 := Call(__e, PrimFunc(symshen_4side_1conditions_1_6goals), tmp13762, tmp13766, tmp13769, tmp13770, V3340)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13758, tmp13771)
}
__typedArg0 := tmp13758
__typedArg1 := tmp13771
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
tmp13799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13779 Obj

if True == tmp13799 {
tmp13797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13797)
}
__typedArg0 := tmp13797
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13781 Obj

if True == tmp13798 {
tmp13794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13794)
}
__typedArg0 := tmp13794
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symsqts, tmp13795)
}
__typedArg0 := symsqts
__typedArg1 := tmp13795
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13783 Obj

if True == tmp13796 {
tmp13791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13791)
}
__typedArg0 := tmp13791
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13792)
}
__typedArg0 := tmp13792
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13785 Obj

if True == tmp13793 {
tmp13787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13787)
}
__typedArg0 := tmp13787
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13788)
}
__typedArg0 := tmp13788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13789)
}
__typedArg0 := Nil
__typedArg1 := tmp13789
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13786 Obj

if True == tmp13790 {
ifres13786 = True


} else {
ifres13786 = False


}

ifres13785 = ifres13786


} else {
ifres13785 = False


}

var ifres13784 Obj

if True == ifres13785 {
ifres13784 = True


} else {
ifres13784 = False


}

ifres13783 = ifres13784


} else {
ifres13783 = False


}

var ifres13782 Obj

if True == ifres13783 {
ifres13782 = True


} else {
ifres13782 = False


}

ifres13781 = ifres13782


} else {
ifres13781 = False


}

var ifres13780 Obj

if True == ifres13781 {
ifres13780 = True


} else {
ifres13780 = False


}

ifres13779 = ifres13780


} else {
ifres13779 = False


}

if True == ifres13779 {
tmp13777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3339)
}
__typedArg0 := V3339
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4side_1conditions_1_6goals), V3336, V3337, V3338, tmp13777, V3340)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.side-conditions->goals"))
}
__typedArg0 := MakeString("partial function shen.side-conditions->goals")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}


}


}, 5)

tmp13875 := Call(__e, ns2_1set, symshen_4side_1conditions_1_6goals, tmp13723)


_ = tmp13875

tmp13876 := MakeNative(func(__e *ControlFlow) {
V3345 := __e.Get(1)
_ = V3345
V3346 := __e.Get(2)
_ = V3346
V3347 := __e.Get(3)
_ = V3347
tmp13926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3347)
}
__typedArg0 := Nil
__typedArg1 := V3347
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13926 {
tmp13877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(MakeString(";"))
}
__typedArg0 := MakeString(";")
return Call(__e, PrimFunc(symintern), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13877, Nil)
}
__typedArg0 := tmp13877
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13920 Obj

if True == tmp13924 {
tmp13922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_b, tmp13922)
}
__typedArg0 := sym_b
__typedArg1 := tmp13922
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13921 Obj

if True == tmp13923 {
ifres13921 = True


} else {
ifres13921 = False


}

ifres13920 = ifres13921


} else {
ifres13920 = False


}

if True == ifres13920 {
tmp13878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13879 := Call(__e, PrimFunc(symshen_4premises_1_6goals), V3345, V3346, tmp13878)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_b, tmp13879)
}
__typedArg0 := sym_b
__typedArg1 := tmp13879
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13914 Obj

if True == tmp13918 {
tmp13916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail, tmp13916)
}
__typedArg0 := symfail
__typedArg1 := tmp13916
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13915 Obj

if True == tmp13917 {
ifres13915 = True


} else {
ifres13915 = False


}

ifres13914 = ifres13915


} else {
ifres13914 = False


}

if True == ifres13914 {
tmp13880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(False, Nil)
}
__typedArg0 := False
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhen, tmp13880)
}
__typedArg0 := symwhen
__typedArg1 := tmp13880
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13883 := Call(__e, PrimFunc(symshen_4premises_1_6goals), V3345, V3346, tmp13882)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13881, tmp13883)
}
__typedArg0 := tmp13881
__typedArg1 := tmp13883
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp13912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13897 Obj

if True == tmp13912 {
tmp13910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13910)
}
__typedArg0 := tmp13910
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13899 Obj

if True == tmp13911 {
tmp13907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13907)
}
__typedArg0 := tmp13907
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp13908)
}
__typedArg0 := tmp13908
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13901 Obj

if True == tmp13909 {
tmp13903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13903)
}
__typedArg0 := tmp13903
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13904)
}
__typedArg0 := tmp13904
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13905)
}
__typedArg0 := Nil
__typedArg1 := tmp13905
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13902 Obj

if True == tmp13906 {
ifres13902 = True


} else {
ifres13902 = False


}

ifres13901 = ifres13902


} else {
ifres13901 = False


}

var ifres13900 Obj

if True == ifres13901 {
ifres13900 = True


} else {
ifres13900 = False


}

ifres13899 = ifres13900


} else {
ifres13899 = False


}

var ifres13898 Obj

if True == ifres13899 {
ifres13898 = True


} else {
ifres13898 = False


}

ifres13897 = ifres13898


} else {
ifres13897 = False


}

if True == ifres13897 {
tmp13884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp13884)
}
__typedArg0 := tmp13884
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13885)
}
__typedArg0 := tmp13885
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13887 := Call(__e, PrimFunc(symshen_4macro_1_8c), tmp13886)


tmp13888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp13888)
}
__typedArg0 := tmp13888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13890 := Call(__e, PrimFunc(symshen_4construct_1context), V3345, tmp13889, V3346)


tmp13891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13890, Nil)
}
__typedArg0 := tmp13890
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13887, tmp13891)
}
__typedArg0 := tmp13887
__typedArg1 := tmp13891
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4system_1S, tmp13892)
}
__typedArg0 := symshen_4system_1S
__typedArg1 := tmp13892
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3347)
}
__typedArg0 := V3347
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13895 := Call(__e, PrimFunc(symshen_4premises_1_6goals), V3345, V3346, tmp13894)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13893, tmp13895)
}
__typedArg0 := tmp13893
__typedArg1 := tmp13895
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.premises->goals"))
}
__typedArg0 := MakeString("partial function shen.premises->goals")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}


}, 3)

tmp13927 := Call(__e, ns2_1set, symshen_4premises_1_6goals, tmp13876)


_ = tmp13927

tmp13928 := MakeNative(func(__e *ControlFlow) {
V3351 := __e.Get(1)
_ = V3351
V3352 := __e.Get(2)
_ = V3352
V3353 := __e.Get(3)
_ = V3353
tmp13948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3352)
}
__typedArg0 := Nil
__typedArg1 := V3352
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13948 {
__e.Return(V3353)
return
} else {
tmp13946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3352)
}
__typedArg0 := V3352
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres13938 Obj

if True == tmp13946 {
tmp13944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3352)
}
__typedArg0 := V3352
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp13944)
}
__typedArg0 := Nil
__typedArg1 := tmp13944
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres13940 Obj

if True == tmp13945 {
tmp13942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3352)
}
__typedArg0 := V3352
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13943 := Call(__e, PrimFunc(symelement_2), tmp13942, V3351)


var ifres13941 Obj

if True == tmp13943 {
ifres13941 = True


} else {
ifres13941 = False


}

ifres13940 = ifres13941


} else {
ifres13940 = False


}

var ifres13939 Obj

if True == ifres13940 {
ifres13939 = True


} else {
ifres13939 = False


}

ifres13938 = ifres13939


} else {
ifres13938 = False


}

if True == ifres13938 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3352)
}
__typedArg0 := V3352
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp13936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3352)
}
__typedArg0 := V3352
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13936 {
tmp13929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3352)
}
__typedArg0 := V3352
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13930 := Call(__e, PrimFunc(symshen_4macro_1_8c), tmp13929)


tmp13931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3352)
}
__typedArg0 := V3352
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13932 := Call(__e, PrimFunc(symshen_4construct_1context), V3351, tmp13931, V3353)


tmp13933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13932, Nil)
}
__typedArg0 := tmp13932
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp13934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp13930, tmp13933)
}
__typedArg0 := tmp13930
__typedArg1 := tmp13933
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp13934)
}
__typedArg0 := symcons
__typedArg1 := tmp13934
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.construct-context"))
}
__typedArg0 := MakeString("partial function shen.construct-context")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 3)

tmp13949 := Call(__e, ns2_1set, symshen_4construct_1context, tmp13928)


_ = tmp13949

tmp13950 := MakeNative(func(__e *ControlFlow) {
V3354 := __e.Get(1)
_ = V3354
tmp13951 := MakeNative(func(__e *ControlFlow) {
Z3356 := __e.Get(1)
_ = Z3356
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3356)
return
}, 1)

tmp13952 := Call(__e, PrimFunc(symmap), tmp13951, V3354)


W335512775 := tmp13952
_ = W335512775

tmp13953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

W335712776 := tmp13953
_ = W335712776

tmp13954 := Call(__e, PrimFunc(symshen_4remove_1datatypes), W335512775, W335712776)


W335812777 := tmp13954
_ = W335812777

tmp13955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, W335812777)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := W335812777
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

W335912778 := tmp13955
_ = W335912778

__e.TailApply(PrimFunc(symshen_4show_1datatypes), W335912778)
return


}, 1)

tmp13956 := Call(__e, ns2_1set, sympreclude, tmp13950)


_ = tmp13956

tmp13957 := MakeNative(func(__e *ControlFlow) {
V3364 := __e.Get(1)
_ = V3364
V3365 := __e.Get(2)
_ = V3365
tmp13964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V3364)
}
__typedArg0 := Nil
__typedArg1 := V3364
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp13964 {
__e.Return(V3365)
return
} else {
tmp13962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V3364)
}
__typedArg0 := V3364
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp13962 {
tmp13958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V3364)
}
__typedArg0 := V3364
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp13959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V3364)
}
__typedArg0 := V3364
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp13960 := Call(__e, PrimFunc(symshen_4unassoc), tmp13959, V3365)


__e.TailApply(PrimFunc(symshen_4remove_1datatypes), tmp13958, tmp13960)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.remove-datatypes"))
}
__typedArg0 := MakeString("implementation error in shen.remove-datatypes")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

tmp13965 := Call(__e, ns2_1set, symshen_4remove_1datatypes, tmp13957)


_ = tmp13965

tmp13966 := MakeNative(func(__e *ControlFlow) {
V3366 := __e.Get(1)
_ = V3366
tmp13967 := MakeNative(func(__e *ControlFlow) {
Z3367 := __e.Get(1)
_ = Z3367
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(Z3367)
}
__typedArg0 := Z3367
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp13967, V3366)
return


}, 1)

tmp13968 := Call(__e, ns2_1set, symshen_4show_1datatypes, tmp13966)


_ = tmp13968

tmp13969 := MakeNative(func(__e *ControlFlow) {
V3368 := __e.Get(1)
_ = V3368
tmp13970 := MakeNative(func(__e *ControlFlow) {
Z3370 := __e.Get(1)
_ = Z3370
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3370)
return
}, 1)

tmp13971 := Call(__e, PrimFunc(symmap), tmp13970, V3368)


W336912779 := tmp13971
_ = W336912779

tmp13972 := MakeNative(func(__e *ControlFlow) {
Z3372 := __e.Get(1)
_ = Z3372
tmp13973 := Call(__e, PrimFunc(symfn), Z3372)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), Z3372, tmp13973)
return


}, 1)

tmp13974 := Call(__e, PrimFunc(symmap), tmp13972, W336912779)


W337112780 := tmp13974
_ = W337112780

tmp13975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

W337312781 := tmp13975
_ = W337312781

__e.TailApply(PrimFunc(symshen_4show_1datatypes), W337312781)
return


}, 1)

tmp13976 := Call(__e, ns2_1set, syminclude, tmp13969)


_ = tmp13976

tmp13977 := MakeNative(func(__e *ControlFlow) {
V3374 := __e.Get(1)
_ = V3374
tmp13978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, Nil)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := Nil
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

W337512782 := tmp13978
_ = W337512782

tmp13979 := MakeNative(func(__e *ControlFlow) {
Z3377 := __e.Get(1)
_ = Z3377
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3377)
return
}, 1)

tmp13980 := Call(__e, PrimFunc(symmap), tmp13979, V3374)


W337612783 := tmp13980
_ = W337612783

tmp13981 := MakeNative(func(__e *ControlFlow) {
Z3379 := __e.Get(1)
_ = Z3379
tmp13982 := Call(__e, PrimFunc(symfn), Z3379)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), Z3379, tmp13982)
return


}, 1)

tmp13983 := Call(__e, PrimFunc(symmap), tmp13981, W337612783)


W337812784 := tmp13983
_ = W337812784

tmp13984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_ddatatypes_d)
}
__typedArg0 := symshen_4_ddatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4show_1datatypes), tmp13984)
return


}, 1)

tmp13985 := Call(__e, ns2_1set, sympreclude_1all_1but, tmp13977)


_ = tmp13985

tmp13986 := MakeNative(func(__e *ControlFlow) {
V3380 := __e.Get(1)
_ = V3380
tmp13987 := MakeNative(func(__e *ControlFlow) {
Z3382 := __e.Get(1)
_ = Z3382
__e.TailApply(PrimFunc(symshen_4intern_1type), Z3382)
return
}, 1)

tmp13988 := Call(__e, PrimFunc(symmap), tmp13987, V3380)


W338112785 := tmp13988
_ = W338112785

tmp13989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dalldatatypes_d)
}
__typedArg0 := symshen_4_dalldatatypes_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

W338312786 := tmp13989
_ = W338312786

tmp13990 := Call(__e, PrimFunc(symshen_4remove_1datatypes), W338112785, W338312786)


tmp13991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_ddatatypes_d, tmp13990)
}
__typedArg0 := symshen_4_ddatatypes_d
__typedArg1 := tmp13990
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

W338412787 := tmp13991
_ = W338412787

__e.TailApply(PrimFunc(symshen_4show_1datatypes), W338412787)
return


}, 1)

__e.TailApply(ns2_1set, syminclude_1all_1but, tmp13986)
return




}, 0)

