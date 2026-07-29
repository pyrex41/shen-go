package main

import . "github.com/pyrex41/shen-go/kl"

var DeclarationsMain = MakeNative(func(__e *ControlFlow) {
tmp7588 := PrimSet(symshen_4_dhistory_d, Nil)

_ = tmp7588

tmp7589 := PrimSet(symshen_4_dtc_d, False)

_ = tmp7589

tmp7590 := Call(__e, PrimFunc(symvector), MakeNumber(20000))


tmp7591 := PrimSet(sym_dproperty_1vector_d, tmp7590)

_ = tmp7591

tmp7592 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4macros), X)
return
}, 1)

tmp7593 := PrimCons(symshen_4macros, tmp7592)

tmp7594 := PrimCons(tmp7593, Nil)

tmp7595 := PrimSet(sym_dmacros_d, tmp7594)

_ = tmp7595

tmp7596 := PrimSet(symshen_4_dgensym_d, MakeNumber(0))

_ = tmp7596

tmp7597 := PrimSet(symshen_4_dtracking_d, Nil)

_ = tmp7597

tmp7598 := PrimSet(symshen_4_dprofiled_d, Nil)

_ = tmp7598

tmp7599 := PrimSet(sym_dhome_1directory_d, MakeString(""))

_ = tmp7599

tmp7600 := PrimCons(symtype, Nil)

tmp7601 := PrimCons(symshen_4input_1h_7, tmp7600)

tmp7602 := PrimCons(symopen, tmp7601)

tmp7603 := PrimCons(symset, tmp7602)

tmp7604 := PrimCons(symwhere, tmp7603)

tmp7605 := PrimCons(symlet, tmp7604)

tmp7606 := PrimCons(symlambda, tmp7605)

tmp7607 := PrimCons(symcons, tmp7606)

tmp7608 := PrimCons(sym_8v, tmp7607)

tmp7609 := PrimCons(sym_8s, tmp7608)

tmp7610 := PrimCons(sym_8p, tmp7609)

tmp7611 := PrimSet(symshen_4_dspecial_d, tmp7610)

_ = tmp7611

tmp7612 := PrimSet(symshen_4_dextraspecial_d, Nil)

_ = tmp7612

tmp7613 := PrimSet(symshen_4_dspy_d, False)

_ = tmp7613

tmp7614 := PrimSet(symshen_4_ddatatypes_d, Nil)

_ = tmp7614

tmp7615 := PrimSet(symshen_4_dalldatatypes_d, Nil)

_ = tmp7615

tmp7616 := PrimSet(symshen_4_dshen_1type_1theory_1enabled_2_d, True)

_ = tmp7616

tmp7617 := PrimSet(symshen_4_dpackage_d, symnull)

_ = tmp7617

tmp7618 := PrimSet(symshen_4_dsynonyms_d, Nil)

_ = tmp7618

tmp7619 := PrimSet(symshen_4_dsystem_d, Nil)

_ = tmp7619

tmp7620 := PrimSet(symshen_4_dsigf_d, Nil)

_ = tmp7620

tmp7621 := PrimSet(symshen_4_doccurs_d, True)

_ = tmp7621

tmp7622 := PrimSet(symshen_4_dfactorise_2_d, False)

_ = tmp7622

tmp7623 := PrimSet(symshen_4_dmaxinferences_d, MakeNumber(1000000))

_ = tmp7623

tmp7624 := PrimSet(sym_dmaximum_1print_1sequence_1size_d, MakeNumber(20))

_ = tmp7624

tmp7625 := PrimSet(symshen_4_dcall_d, MakeNumber(0))

_ = tmp7625

tmp7626 := PrimSet(symshen_4_dinfs_d, MakeNumber(0))

_ = tmp7626

tmp7627 := PrimSet(sym_dhush_d, False)

_ = tmp7627

tmp7628 := PrimSet(symshen_4_doptimise_d, False)

_ = tmp7628

tmp7629 := PrimSet(sym_dversion_d, MakeString("41.2"))

_ = tmp7629

tmp7630 := PrimSet(symshen_4_dnames_d, Nil)

_ = tmp7630

tmp7631 := PrimSet(symshen_4_dstep_d, False)

_ = tmp7631

tmp7632 := PrimSet(symshen_4_dit_d, MakeString(""))

_ = tmp7632

tmp7633 := PrimSet(symshen_4_dresidue_d, Nil)

_ = tmp7633

tmp7634 := PrimSet(symshen_4_dprolog_1memory_d, MakeNumber(1000))

_ = tmp7634

tmp7635 := PrimSet(symshen_4_dloading_2_d, False)

_ = tmp7635

tmp7636 := PrimSet(symshen_4_duserdefs_d, Nil)

_ = tmp7636

tmp7637 := MakeNative(func(__e *ControlFlow) {
tmp7638 := MakeNative(func(__e *ControlFlow) {
Z5718 := __e.Get(1)
_ = Z5718
__e.TailApply(PrimFunc(symshen_4typename), Z5718)
return
}, 1)

tmp7639 := PrimValue(symshen_4_dalldatatypes_d)

__e.TailApply(PrimFunc(symmap), tmp7638, tmp7639)
return


}, 0)

tmp7640 := Call(__e, ns2_1set, symdatatypes, tmp7637)


_ = tmp7640

tmp7641 := MakeNative(func(__e *ControlFlow) {
tmp7642 := MakeNative(func(__e *ControlFlow) {
Z5719 := __e.Get(1)
_ = Z5719
__e.TailApply(PrimFunc(symshen_4typename), Z5719)
return
}, 1)

tmp7643 := PrimValue(symshen_4_ddatatypes_d)

__e.TailApply(PrimFunc(symmap), tmp7642, tmp7643)
return


}, 0)

tmp7644 := Call(__e, ns2_1set, symincluded, tmp7641)


_ = tmp7644

tmp7645 := MakeNative(func(__e *ControlFlow) {
V5722 := __e.Get(1)
_ = V5722
tmp7650 := PrimIsPair(V5722)

if True == tmp7650 {
tmp7646 := PrimHead(V5722)

tmp7647 := PrimStr(tmp7646)

tmp7648 := Call(__e, PrimFunc(symshen_4typename_1h), tmp7647)


__e.Return(PrimIntern(tmp7648))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.typename")))
return
}


}, 1)

tmp7651 := Call(__e, ns2_1set, symshen_4typename, tmp7645)


_ = tmp7651

tmp7652 := MakeNative(func(__e *ControlFlow) {
V5723 := __e.Get(1)
_ = V5723
tmp7659 := PrimEqual(MakeString("#type"), V5723)

if True == tmp7659 {
__e.Return(MakeString(""))
return
} else {
tmp7657 := Call(__e, PrimFunc(symshen_4_7string_2), V5723)


if True == tmp7657 {
tmp7653 := Call(__e, PrimFunc(symhdstr), V5723)


tmp7654 := PrimTailString(V5723)

tmp7655 := Call(__e, PrimFunc(symshen_4typename_1h), tmp7654)


__e.Return(PrimStringConcat(tmp7653, tmp7655))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.typename-h")))
return
}


}


}, 1)

tmp7660 := Call(__e, ns2_1set, symshen_4typename_1h, tmp7652)


_ = tmp7660

tmp7661 := MakeNative(func(__e *ControlFlow) {
V5724 := __e.Get(1)
_ = V5724
tmp7665 := PrimLessThan(V5724, MakeNumber(0))

if True == tmp7665 {
__e.Return(PrimValue(symshen_4_dprolog_1memory_d))
return
} else {
tmp7663 := PrimIsInteger(V5724)

if True == tmp7663 {
__e.Return(PrimSet(symshen_4_dprolog_1memory_d, V5724))
return
} else {
__e.Return(PrimSimpleError(MakeString("prolog memory expects an integer value\n")))
return
}


}


}, 1)

tmp7666 := Call(__e, ns2_1set, symprolog_1memory, tmp7661)


_ = tmp7666

tmp7667 := MakeNative(func(__e *ControlFlow) {
V5727 := __e.Get(1)
_ = V5727
tmp7683 := PrimEqual(Nil, V5727)

if True == tmp7683 {
__e.Return(Nil)
return
} else {
tmp7681 := PrimIsPair(V5727)

var ifres7677 Obj

if True == tmp7681 {
tmp7679 := PrimTail(V5727)

tmp7680 := PrimIsPair(tmp7679)

var ifres7678 Obj

if True == tmp7680 {
ifres7678 = True


} else {
ifres7678 = False


}

ifres7677 = ifres7678


} else {
ifres7677 = False


}

if True == ifres7677 {
tmp7668 := MakeNative(func(__e *ControlFlow) {
W5728 := __e.Get(1)
_ = W5728
tmp7669 := PrimTail(V5727)

tmp7670 := PrimTail(tmp7669)

__e.TailApply(PrimFunc(symshen_4initialise_1arity_1table), tmp7670)
return


}, 1)

tmp7671 := PrimHead(V5727)

tmp7672 := PrimTail(V5727)

tmp7673 := PrimHead(tmp7672)

tmp7674 := PrimValue(sym_dproperty_1vector_d)

tmp7675 := Call(__e, PrimFunc(symput), tmp7671, symarity, tmp7673, tmp7674)


__e.TailApply(tmp7668, tmp7675)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.initialise-arity-table")))
return
}


}


}, 1)

tmp7684 := Call(__e, ns2_1set, symshen_4initialise_1lambda_1tables, tmp7667)


_ = tmp7684

tmp7685 := MakeNative(func(__e *ControlFlow) {
V5729 := __e.Get(1)
_ = V5729
tmp7686 := MakeNative(func(__e *ControlFlow) {
tmp7687 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V5729, symarity, tmp7687)
return


}, 0)

tmp7688 := MakeNative(func(__e *ControlFlow) {
Z5730 := __e.Get(1)
_ = Z5730
__e.Return(MakeNumber(-1))
return
}, 1)

__e.TailApply(try_1catch, tmp7686, tmp7688)
return


}, 1)

tmp7689 := Call(__e, ns2_1set, symarity, tmp7685)


_ = tmp7689

tmp7690 := MakeNative(func(__e *ControlFlow) {
V5733 := __e.Get(1)
_ = V5733
tmp7706 := PrimEqual(Nil, V5733)

if True == tmp7706 {
__e.Return(Nil)
return
} else {
tmp7704 := PrimIsPair(V5733)

var ifres7700 Obj

if True == tmp7704 {
tmp7702 := PrimTail(V5733)

tmp7703 := PrimIsPair(tmp7702)

var ifres7701 Obj

if True == tmp7703 {
ifres7701 = True


} else {
ifres7701 = False


}

ifres7700 = ifres7701


} else {
ifres7700 = False


}

if True == ifres7700 {
tmp7691 := MakeNative(func(__e *ControlFlow) {
W5734 := __e.Get(1)
_ = W5734
tmp7692 := PrimTail(V5733)

tmp7693 := PrimTail(tmp7692)

__e.TailApply(PrimFunc(symshen_4initialise_1arity_1table), tmp7693)
return


}, 1)

tmp7694 := PrimHead(V5733)

tmp7695 := PrimTail(V5733)

tmp7696 := PrimHead(tmp7695)

tmp7697 := PrimValue(sym_dproperty_1vector_d)

tmp7698 := Call(__e, PrimFunc(symput), tmp7694, symarity, tmp7696, tmp7697)


__e.TailApply(tmp7691, tmp7698)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.initialise-arity-table")))
return
}


}


}, 1)

tmp7707 := Call(__e, ns2_1set, symshen_4initialise_1arity_1table, tmp7690)


_ = tmp7707

tmp7708 := PrimCons(MakeNumber(2), Nil)

tmp7709 := PrimCons(sym_8s, tmp7708)

tmp7710 := PrimCons(MakeNumber(2), tmp7709)

tmp7711 := PrimCons(sym_8v, tmp7710)

tmp7712 := PrimCons(MakeNumber(2), tmp7711)

tmp7713 := PrimCons(sym_8p, tmp7712)

tmp7714 := PrimCons(MakeNumber(1), tmp7713)

tmp7715 := PrimCons(sym_5_b_6, tmp7714)

tmp7716 := PrimCons(MakeNumber(1), tmp7715)

tmp7717 := PrimCons(sym_5end_6, tmp7716)

tmp7718 := PrimCons(MakeNumber(1), tmp7717)

tmp7719 := PrimCons(sym_5e_6, tmp7718)

tmp7720 := PrimCons(MakeNumber(2), tmp7719)

tmp7721 := PrimCons(sym_a_a, tmp7720)

tmp7722 := PrimCons(MakeNumber(2), tmp7721)

tmp7723 := PrimCons(sym_1, tmp7722)

tmp7724 := PrimCons(MakeNumber(2), tmp7723)

tmp7725 := PrimCons(sym_c, tmp7724)

tmp7726 := PrimCons(MakeNumber(2), tmp7725)

tmp7727 := PrimCons(sym_d, tmp7726)

tmp7728 := PrimCons(MakeNumber(2), tmp7727)

tmp7729 := PrimCons(sym_7, tmp7728)

tmp7730 := PrimCons(MakeNumber(1), tmp7729)

tmp7731 := PrimCons(symy_1or_1n_2, tmp7730)

tmp7732 := PrimCons(MakeNumber(2), tmp7731)

tmp7733 := PrimCons(symwrite_1to_1file, tmp7732)

tmp7734 := PrimCons(MakeNumber(2), tmp7733)

tmp7735 := PrimCons(symwrite_1byte, tmp7734)

tmp7736 := PrimCons(MakeNumber(5), tmp7735)

tmp7737 := PrimCons(symwhen, tmp7736)

tmp7738 := PrimCons(MakeNumber(0), tmp7737)

tmp7739 := PrimCons(symversion, tmp7738)

tmp7740 := PrimCons(MakeNumber(5), tmp7739)

tmp7741 := PrimCons(symvar_2, tmp7740)

tmp7742 := PrimCons(MakeNumber(1), tmp7741)

tmp7743 := PrimCons(symvariable_2, tmp7742)

tmp7744 := PrimCons(MakeNumber(1), tmp7743)

tmp7745 := PrimCons(symvalue, tmp7744)

tmp7746 := PrimCons(MakeNumber(3), tmp7745)

tmp7747 := PrimCons(symvector_1_6, tmp7746)

tmp7748 := PrimCons(MakeNumber(1), tmp7747)

tmp7749 := PrimCons(symvector_2, tmp7748)

tmp7750 := PrimCons(MakeNumber(1), tmp7749)

tmp7751 := PrimCons(symvector, tmp7750)

tmp7752 := PrimCons(MakeNumber(0), tmp7751)

tmp7753 := PrimCons(symuserdefs, tmp7752)

tmp7754 := PrimCons(MakeNumber(2), tmp7753)

tmp7755 := PrimCons(symupdate_1lambda_1table, tmp7754)

tmp7756 := PrimCons(MakeNumber(1), tmp7755)

tmp7757 := PrimCons(symundefmacro, tmp7756)

tmp7758 := PrimCons(MakeNumber(1), tmp7757)

tmp7759 := PrimCons(symuntrack, tmp7758)

tmp7760 := PrimCons(MakeNumber(2), tmp7759)

tmp7761 := PrimCons(symunion, tmp7760)

tmp7762 := PrimCons(MakeNumber(1), tmp7761)

tmp7763 := PrimCons(symunprofile, tmp7762)

tmp7764 := PrimCons(MakeNumber(3), tmp7763)

tmp7765 := PrimCons(symunput, tmp7764)

tmp7766 := PrimCons(MakeNumber(1), tmp7765)

tmp7767 := PrimCons(symundefmacro, tmp7766)

tmp7768 := PrimCons(MakeNumber(1), tmp7767)

tmp7769 := PrimCons(symunabsolute, tmp7768)

tmp7770 := PrimCons(MakeNumber(5), tmp7769)

tmp7771 := PrimCons(symreturn, tmp7770)

tmp7772 := PrimCons(MakeNumber(2), tmp7771)

tmp7773 := PrimCons(symtype, tmp7772)

tmp7774 := PrimCons(MakeNumber(1), tmp7773)

tmp7775 := PrimCons(symtuple_2, tmp7774)

tmp7776 := PrimCons(MakeNumber(2), tmp7775)

tmp7777 := PrimCons(symtrap_1error, tmp7776)

tmp7778 := PrimCons(MakeNumber(0), tmp7777)

tmp7779 := PrimCons(symtracked, tmp7778)

tmp7780 := PrimCons(MakeNumber(1), tmp7779)

tmp7781 := PrimCons(symtrack, tmp7780)

tmp7782 := PrimCons(MakeNumber(1), tmp7781)

tmp7783 := PrimCons(symtlstr, tmp7782)

tmp7784 := PrimCons(MakeNumber(1), tmp7783)

tmp7785 := PrimCons(symthaw, tmp7784)

tmp7786 := PrimCons(MakeNumber(0), tmp7785)

tmp7787 := PrimCons(symtc_2, tmp7786)

tmp7788 := PrimCons(MakeNumber(1), tmp7787)

tmp7789 := PrimCons(symtc, tmp7788)

tmp7790 := PrimCons(MakeNumber(1), tmp7789)

tmp7791 := PrimCons(symtl, tmp7790)

tmp7792 := PrimCons(MakeNumber(1), tmp7791)

tmp7793 := PrimCons(symtail, tmp7792)

tmp7794 := PrimCons(MakeNumber(1), tmp7793)

tmp7795 := PrimCons(symsystemf, tmp7794)

tmp7796 := PrimCons(MakeNumber(1), tmp7795)

tmp7797 := PrimCons(symsymbol_2, tmp7796)

tmp7798 := PrimCons(MakeNumber(1), tmp7797)

tmp7799 := PrimCons(symsum, tmp7798)

tmp7800 := PrimCons(MakeNumber(3), tmp7799)

tmp7801 := PrimCons(symsubst, tmp7800)

tmp7802 := PrimCons(MakeNumber(1), tmp7801)

tmp7803 := PrimCons(symstring_2, tmp7802)

tmp7804 := PrimCons(MakeNumber(1), tmp7803)

tmp7805 := PrimCons(symstring_1_6symbol, tmp7804)

tmp7806 := PrimCons(MakeNumber(1), tmp7805)

tmp7807 := PrimCons(symstring_1_6n, tmp7806)

tmp7808 := PrimCons(MakeNumber(1), tmp7807)

tmp7809 := PrimCons(symstr, tmp7808)

tmp7810 := PrimCons(MakeNumber(0), tmp7809)

tmp7811 := PrimCons(symstoutput, tmp7810)

tmp7812 := PrimCons(MakeNumber(0), tmp7811)

tmp7813 := PrimCons(symstinput, tmp7812)

tmp7814 := PrimCons(MakeNumber(0), tmp7813)

tmp7815 := PrimCons(symstep_2, tmp7814)

tmp7816 := PrimCons(MakeNumber(1), tmp7815)

tmp7817 := PrimCons(symstep, tmp7816)

tmp7818 := PrimCons(MakeNumber(0), tmp7817)

tmp7819 := PrimCons(symspy_2, tmp7818)

tmp7820 := PrimCons(MakeNumber(1), tmp7819)

tmp7821 := PrimCons(symspy, tmp7820)

tmp7822 := PrimCons(MakeNumber(2), tmp7821)

tmp7823 := PrimCons(symspecialise, tmp7822)

tmp7824 := PrimCons(MakeNumber(1), tmp7823)

tmp7825 := PrimCons(symsnd, tmp7824)

tmp7826 := PrimCons(MakeNumber(1), tmp7825)

tmp7827 := PrimCons(symsimple_1error, tmp7826)

tmp7828 := PrimCons(MakeNumber(2), tmp7827)

tmp7829 := PrimCons(symset, tmp7828)

tmp7830 := PrimCons(MakeNumber(1), tmp7829)

tmp7831 := PrimCons(symreverse, tmp7830)

tmp7832 := PrimCons(MakeNumber(2), tmp7831)

tmp7833 := PrimCons(symremove, tmp7832)

tmp7834 := PrimCons(MakeNumber(0), tmp7833)

tmp7835 := PrimCons(symrelease, tmp7834)

tmp7836 := PrimCons(MakeNumber(1), tmp7835)

tmp7837 := PrimCons(symreceive, tmp7836)

tmp7838 := PrimCons(MakeNumber(1), tmp7837)

tmp7839 := PrimCons(symshen_4read_1unit_1string, tmp7838)

tmp7840 := PrimCons(MakeNumber(1), tmp7839)

tmp7841 := PrimCons(symread_1from_1string_1unprocessed, tmp7840)

tmp7842 := PrimCons(MakeNumber(1), tmp7841)

tmp7843 := PrimCons(symread_1from_1string, tmp7842)

tmp7844 := PrimCons(MakeNumber(1), tmp7843)

tmp7845 := PrimCons(symread_1byte, tmp7844)

tmp7846 := PrimCons(MakeNumber(1), tmp7845)

tmp7847 := PrimCons(symread, tmp7846)

tmp7848 := PrimCons(MakeNumber(1), tmp7847)

tmp7849 := PrimCons(symread_1file, tmp7848)

tmp7850 := PrimCons(MakeNumber(1), tmp7849)

tmp7851 := PrimCons(symread_1file_1as_1bytelist, tmp7850)

tmp7852 := PrimCons(MakeNumber(1), tmp7851)

tmp7853 := PrimCons(symread_1file_1as_1string, tmp7852)

tmp7854 := PrimCons(MakeNumber(4), tmp7853)

tmp7855 := PrimCons(symput, tmp7854)

tmp7856 := PrimCons(MakeNumber(1), tmp7855)

tmp7857 := PrimCons(symprotect, tmp7856)

tmp7858 := PrimCons(MakeNumber(1), tmp7857)

tmp7859 := PrimCons(sympreclude_1all_1but, tmp7858)

tmp7860 := PrimCons(MakeNumber(1), tmp7859)

tmp7861 := PrimCons(sympreclude, tmp7860)

tmp7862 := PrimCons(MakeNumber(1), tmp7861)

tmp7863 := PrimCons(symps, tmp7862)

tmp7864 := PrimCons(MakeNumber(2), tmp7863)

tmp7865 := PrimCons(sympr, tmp7864)

tmp7866 := PrimCons(MakeNumber(1), tmp7865)

tmp7867 := PrimCons(symprofile_1results, tmp7866)

tmp7868 := PrimCons(MakeNumber(1), tmp7867)

tmp7869 := PrimCons(symprolog_1memory, tmp7868)

tmp7870 := PrimCons(MakeNumber(1), tmp7869)

tmp7871 := PrimCons(symshen_4printF, tmp7870)

tmp7872 := PrimCons(MakeNumber(1), tmp7871)

tmp7873 := PrimCons(symshen_4print_1freshterm, tmp7872)

tmp7874 := PrimCons(MakeNumber(1), tmp7873)

tmp7875 := PrimCons(symshen_4print_1prolog_1vector, tmp7874)

tmp7876 := PrimCons(MakeNumber(1), tmp7875)

tmp7877 := PrimCons(symprofile, tmp7876)

tmp7878 := PrimCons(MakeNumber(1), tmp7877)

tmp7879 := PrimCons(symprint, tmp7878)

tmp7880 := PrimCons(MakeNumber(1), tmp7879)

tmp7881 := PrimCons(sympreclude_1all_1but, tmp7880)

tmp7882 := PrimCons(MakeNumber(2), tmp7881)

tmp7883 := PrimCons(sympos, tmp7882)

tmp7884 := PrimCons(MakeNumber(0), tmp7883)

tmp7885 := PrimCons(symporters, tmp7884)

tmp7886 := PrimCons(MakeNumber(0), tmp7885)

tmp7887 := PrimCons(symport, tmp7886)

tmp7888 := PrimCons(MakeNumber(1), tmp7887)

tmp7889 := PrimCons(sympackage_2, tmp7888)

tmp7890 := PrimCons(MakeNumber(3), tmp7889)

tmp7891 := PrimCons(sympackage, tmp7890)

tmp7892 := PrimCons(MakeNumber(0), tmp7891)

tmp7893 := PrimCons(symos, tmp7892)

tmp7894 := PrimCons(MakeNumber(2), tmp7893)

tmp7895 := PrimCons(symor, tmp7894)

tmp7896 := PrimCons(MakeNumber(0), tmp7895)

tmp7897 := PrimCons(symoptimise_2, tmp7896)

tmp7898 := PrimCons(MakeNumber(1), tmp7897)

tmp7899 := PrimCons(symoptimise, tmp7898)

tmp7900 := PrimCons(MakeNumber(2), tmp7899)

tmp7901 := PrimCons(symopen, tmp7900)

tmp7902 := PrimCons(MakeNumber(1), tmp7901)

tmp7903 := PrimCons(symoccurs_1check, tmp7902)

tmp7904 := PrimCons(MakeNumber(0), tmp7903)

tmp7905 := PrimCons(symoccurs_2, tmp7904)

tmp7906 := PrimCons(MakeNumber(2), tmp7905)

tmp7907 := PrimCons(symoccurrences, tmp7906)

tmp7908 := PrimCons(MakeNumber(1), tmp7907)

tmp7909 := PrimCons(symoccurs_1check, tmp7908)

tmp7910 := PrimCons(MakeNumber(1), tmp7909)

tmp7911 := PrimCons(symnumber_2, tmp7910)

tmp7912 := PrimCons(MakeNumber(1), tmp7911)

tmp7913 := PrimCons(symn_1_6string, tmp7912)

tmp7914 := PrimCons(MakeNumber(2), tmp7913)

tmp7915 := PrimCons(symnth, tmp7914)

tmp7916 := PrimCons(MakeNumber(1), tmp7915)

tmp7917 := PrimCons(symnot, tmp7916)

tmp7918 := PrimCons(MakeNumber(1), tmp7917)

tmp7919 := PrimCons(symnl, tmp7918)

tmp7920 := PrimCons(MakeNumber(1), tmp7919)

tmp7921 := PrimCons(symmaxinferences, tmp7920)

tmp7922 := PrimCons(MakeNumber(2), tmp7921)

tmp7923 := PrimCons(symmapcan, tmp7922)

tmp7924 := PrimCons(MakeNumber(2), tmp7923)

tmp7925 := PrimCons(symmap, tmp7924)

tmp7926 := PrimCons(MakeNumber(1), tmp7925)

tmp7927 := PrimCons(symmacroexpand, tmp7926)

tmp7928 := PrimCons(MakeNumber(1), tmp7927)

tmp7929 := PrimCons(symvector, tmp7928)

tmp7930 := PrimCons(MakeNumber(2), tmp7929)

tmp7931 := PrimCons(sym_5_a, tmp7930)

tmp7932 := PrimCons(MakeNumber(2), tmp7931)

tmp7933 := PrimCons(sym_5, tmp7932)

tmp7934 := PrimCons(MakeNumber(1), tmp7933)

tmp7935 := PrimCons(symload, tmp7934)

tmp7936 := PrimCons(MakeNumber(1), tmp7935)

tmp7937 := PrimCons(symlist, tmp7936)

tmp7938 := PrimCons(MakeNumber(1), tmp7937)

tmp7939 := PrimCons(symlineread, tmp7938)

tmp7940 := PrimCons(MakeNumber(1), tmp7939)

tmp7941 := PrimCons(symlimit, tmp7940)

tmp7942 := PrimCons(MakeNumber(1), tmp7941)

tmp7943 := PrimCons(symlength, tmp7942)

tmp7944 := PrimCons(MakeNumber(0), tmp7943)

tmp7945 := PrimCons(symlanguage, tmp7944)

tmp7946 := PrimCons(MakeNumber(6), tmp7945)

tmp7947 := PrimCons(symis_b, tmp7946)

tmp7948 := PrimCons(MakeNumber(6), tmp7947)

tmp7949 := PrimCons(symis, tmp7948)

tmp7950 := PrimCons(MakeNumber(0), tmp7949)

tmp7951 := PrimCons(symit, tmp7950)

tmp7952 := PrimCons(MakeNumber(1), tmp7951)

tmp7953 := PrimCons(syminternal, tmp7952)

tmp7954 := PrimCons(MakeNumber(2), tmp7953)

tmp7955 := PrimCons(symintersection, tmp7954)

tmp7956 := PrimCons(MakeNumber(1), tmp7955)

tmp7957 := PrimCons(syminclude_1all_1but, tmp7956)

tmp7958 := PrimCons(MakeNumber(0), tmp7957)

tmp7959 := PrimCons(symimplementation, tmp7958)

tmp7960 := PrimCons(MakeNumber(2), tmp7959)

tmp7961 := PrimCons(syminput_7, tmp7960)

tmp7962 := PrimCons(MakeNumber(1), tmp7961)

tmp7963 := PrimCons(syminput, tmp7962)

tmp7964 := PrimCons(MakeNumber(0), tmp7963)

tmp7965 := PrimCons(syminferences, tmp7964)

tmp7966 := PrimCons(MakeNumber(1), tmp7965)

tmp7967 := PrimCons(symintern, tmp7966)

tmp7968 := PrimCons(MakeNumber(1), tmp7967)

tmp7969 := PrimCons(syminternal, tmp7968)

tmp7970 := PrimCons(MakeNumber(1), tmp7969)

tmp7971 := PrimCons(syminteger_2, tmp7970)

tmp7972 := PrimCons(MakeNumber(1), tmp7971)

tmp7973 := PrimCons(symin_1package, tmp7972)

tmp7974 := PrimCons(MakeNumber(0), tmp7973)

tmp7975 := PrimCons(symincluded, tmp7974)

tmp7976 := PrimCons(MakeNumber(1), tmp7975)

tmp7977 := PrimCons(syminclude, tmp7976)

tmp7978 := PrimCons(MakeNumber(3), tmp7977)

tmp7979 := PrimCons(symif, tmp7978)

tmp7980 := PrimCons(MakeNumber(1), tmp7979)

tmp7981 := PrimCons(symhush, tmp7980)

tmp7982 := PrimCons(MakeNumber(0), tmp7981)

tmp7983 := PrimCons(symhush_2, tmp7982)

tmp7984 := PrimCons(MakeNumber(1), tmp7983)

tmp7985 := PrimCons(symhead, tmp7984)

tmp7986 := PrimCons(MakeNumber(1), tmp7985)

tmp7987 := PrimCons(symhdstr, tmp7986)

tmp7988 := PrimCons(MakeNumber(1), tmp7987)

tmp7989 := PrimCons(symhdv, tmp7988)

tmp7990 := PrimCons(MakeNumber(1), tmp7989)

tmp7991 := PrimCons(symhd, tmp7990)

tmp7992 := PrimCons(MakeNumber(2), tmp7991)

tmp7993 := PrimCons(symhash, tmp7992)

tmp7994 := PrimCons(MakeNumber(2), tmp7993)

tmp7995 := PrimCons(sym_a, tmp7994)

tmp7996 := PrimCons(MakeNumber(2), tmp7995)

tmp7997 := PrimCons(sym_6_a, tmp7996)

tmp7998 := PrimCons(MakeNumber(2), tmp7997)

tmp7999 := PrimCons(sym_6, tmp7998)

tmp8000 := PrimCons(MakeNumber(2), tmp7999)

tmp8001 := PrimCons(sym_5_1vector, tmp8000)

tmp8002 := PrimCons(MakeNumber(2), tmp8001)

tmp8003 := PrimCons(sym_5_1address, tmp8002)

tmp8004 := PrimCons(MakeNumber(3), tmp8003)

tmp8005 := PrimCons(symaddress_1_6, tmp8004)

tmp8006 := PrimCons(MakeNumber(1), tmp8005)

tmp8007 := PrimCons(symget_1time, tmp8006)

tmp8008 := PrimCons(MakeNumber(3), tmp8007)

tmp8009 := PrimCons(symget, tmp8008)

tmp8010 := PrimCons(MakeNumber(1), tmp8009)

tmp8011 := PrimCons(symgensym, tmp8010)

tmp8012 := PrimCons(MakeNumber(1), tmp8011)

tmp8013 := PrimCons(symfunction, tmp8012)

tmp8014 := PrimCons(MakeNumber(1), tmp8013)

tmp8015 := PrimCons(symfn, tmp8014)

tmp8016 := PrimCons(MakeNumber(1), tmp8015)

tmp8017 := PrimCons(symfst, tmp8016)

tmp8018 := PrimCons(MakeNumber(0), tmp8017)

tmp8019 := PrimCons(symfresh, tmp8018)

tmp8020 := PrimCons(MakeNumber(1), tmp8019)

tmp8021 := PrimCons(symfreeze, tmp8020)

tmp8022 := PrimCons(MakeNumber(5), tmp8021)

tmp8023 := PrimCons(symfork, tmp8022)

tmp8024 := PrimCons(MakeNumber(1), tmp8023)

tmp8025 := PrimCons(symforeign, tmp8024)

tmp8026 := PrimCons(MakeNumber(7), tmp8025)

tmp8027 := PrimCons(symfindall, tmp8026)

tmp8028 := PrimCons(MakeNumber(2), tmp8027)

tmp8029 := PrimCons(symfix, tmp8028)

tmp8030 := PrimCons(MakeNumber(0), tmp8029)

tmp8031 := PrimCons(symfail, tmp8030)

tmp8032 := PrimCons(MakeNumber(2), tmp8031)

tmp8033 := PrimCons(symfail_1if, tmp8032)

tmp8034 := PrimCons(MakeNumber(0), tmp8033)

tmp8035 := PrimCons(symfactorise_2, tmp8034)

tmp8036 := PrimCons(MakeNumber(1), tmp8035)

tmp8037 := PrimCons(symfactorise, tmp8036)

tmp8038 := PrimCons(MakeNumber(1), tmp8037)

tmp8039 := PrimCons(symexternal, tmp8038)

tmp8040 := PrimCons(MakeNumber(1), tmp8039)

tmp8041 := PrimCons(symexplode, tmp8040)

tmp8042 := PrimCons(MakeNumber(1), tmp8041)

tmp8043 := PrimCons(symeval_1kl, tmp8042)

tmp8044 := PrimCons(MakeNumber(1), tmp8043)

tmp8045 := PrimCons(symeval, tmp8044)

tmp8046 := PrimCons(MakeNumber(1), tmp8045)

tmp8047 := PrimCons(symerror_1to_1string, tmp8046)

tmp8048 := PrimCons(MakeNumber(1), tmp8047)

tmp8049 := PrimCons(symexternal, tmp8048)

tmp8050 := PrimCons(MakeNumber(1), tmp8049)

tmp8051 := PrimCons(symenable_1type_1theory, tmp8050)

tmp8052 := PrimCons(MakeNumber(1), tmp8051)

tmp8053 := PrimCons(symempty_2, tmp8052)

tmp8054 := PrimCons(MakeNumber(2), tmp8053)

tmp8055 := PrimCons(symelement_2, tmp8054)

tmp8056 := PrimCons(MakeNumber(2), tmp8055)

tmp8057 := PrimCons(symdo, tmp8056)

tmp8058 := PrimCons(MakeNumber(2), tmp8057)

tmp8059 := PrimCons(symdifference, tmp8058)

tmp8060 := PrimCons(MakeNumber(1), tmp8059)

tmp8061 := PrimCons(symdestroy, tmp8060)

tmp8062 := PrimCons(MakeNumber(2), tmp8061)

tmp8063 := PrimCons(symdeclare, tmp8062)

tmp8064 := PrimCons(MakeNumber(0), tmp8063)

tmp8065 := PrimCons(symdatatypes, tmp8064)

tmp8066 := PrimCons(MakeNumber(1), tmp8065)

tmp8067 := PrimCons(symclose, tmp8066)

tmp8068 := PrimCons(MakeNumber(2), tmp8067)

tmp8069 := PrimCons(symcn, tmp8068)

tmp8070 := PrimCons(MakeNumber(1), tmp8069)

tmp8071 := PrimCons(symcons_2, tmp8070)

tmp8072 := PrimCons(MakeNumber(2), tmp8071)

tmp8073 := PrimCons(symcons, tmp8072)

tmp8074 := PrimCons(MakeNumber(2), tmp8073)

tmp8075 := PrimCons(symconcat, tmp8074)

tmp8076 := PrimCons(MakeNumber(2), tmp8075)

tmp8077 := PrimCons(symcompile, tmp8076)

tmp8078 := PrimCons(MakeNumber(1), tmp8077)

tmp8079 := PrimCons(symcd, tmp8078)

tmp8080 := PrimCons(MakeNumber(5), tmp8079)

tmp8081 := PrimCons(symcall, tmp8080)

tmp8082 := PrimCons(MakeNumber(6), tmp8081)

tmp8083 := PrimCons(symbind, tmp8082)

tmp8084 := PrimCons(MakeNumber(1), tmp8083)

tmp8085 := PrimCons(symbound_2, tmp8084)

tmp8086 := PrimCons(MakeNumber(1), tmp8085)

tmp8087 := PrimCons(symbootstrap, tmp8086)

tmp8088 := PrimCons(MakeNumber(1), tmp8087)

tmp8089 := PrimCons(symboolean_2, tmp8088)

tmp8090 := PrimCons(MakeNumber(1), tmp8089)

tmp8091 := PrimCons(symatom_2, tmp8090)

tmp8092 := PrimCons(MakeNumber(2), tmp8091)

tmp8093 := PrimCons(symassoc, tmp8092)

tmp8094 := PrimCons(MakeNumber(1), tmp8093)

tmp8095 := PrimCons(symarity, tmp8094)

tmp8096 := PrimCons(MakeNumber(2), tmp8095)

tmp8097 := PrimCons(symappend, tmp8096)

tmp8098 := PrimCons(MakeNumber(2), tmp8097)

tmp8099 := PrimCons(symand, tmp8098)

tmp8100 := PrimCons(MakeNumber(2), tmp8099)

tmp8101 := PrimCons(symadjoin, tmp8100)

tmp8102 := PrimCons(MakeNumber(3), tmp8101)

tmp8103 := PrimCons(symaddress_1_6, tmp8102)

tmp8104 := PrimCons(MakeNumber(1), tmp8103)

tmp8105 := PrimCons(symabsvector, tmp8104)

tmp8106 := PrimCons(MakeNumber(1), tmp8105)

tmp8107 := PrimCons(symabsvector_2, tmp8106)

tmp8108 := PrimCons(MakeNumber(1), tmp8107)

tmp8109 := PrimCons(symabsolute, tmp8108)

tmp8110 := PrimCons(MakeNumber(0), tmp8109)

tmp8111 := PrimCons(symabort, tmp8110)

tmp8112 := Call(__e, PrimFunc(symshen_4initialise_1arity_1table), tmp8111)


_ = tmp8112

tmp8113 := MakeNative(func(__e *ControlFlow) {
V5735 := __e.Get(1)
_ = V5735
tmp8114 := MakeNative(func(__e *ControlFlow) {
W5736 := __e.Get(1)
_ = W5736
tmp8115 := MakeNative(func(__e *ControlFlow) {
W5737 := __e.Get(1)
_ = W5737
__e.Return(V5735)
return
}, 1)

tmp8116 := Call(__e, PrimFunc(symadjoin), V5735, W5736)


tmp8117 := PrimValue(sym_dproperty_1vector_d)

tmp8118 := Call(__e, PrimFunc(symput), symshen, symshen_4external_1symbols, tmp8116, tmp8117)


__e.TailApply(tmp8115, tmp8118)
return


}, 1)

tmp8119 := PrimValue(sym_dproperty_1vector_d)

tmp8120 := Call(__e, PrimFunc(symget), symshen, symshen_4external_1symbols, tmp8119)


__e.TailApply(tmp8114, tmp8120)
return


}, 1)

tmp8121 := Call(__e, ns2_1set, symsystemf, tmp8113)


_ = tmp8121

tmp8122 := MakeNative(func(__e *ControlFlow) {
V5738 := __e.Get(1)
_ = V5738
V5739 := __e.Get(2)
_ = V5739
tmp8124 := Call(__e, PrimFunc(symelement_2), V5738, V5739)


if True == tmp8124 {
__e.Return(V5739)
return
} else {
__e.Return(PrimCons(V5738, V5739))
return
}


}, 2)

tmp8125 := Call(__e, ns2_1set, symadjoin, tmp8122)


_ = tmp8125

tmp8126 := PrimIntern(MakeString(":"))

tmp8127 := PrimIntern(MakeString(";"))

tmp8128 := PrimIntern(MakeString(":="))

tmp8129 := PrimIntern(MakeString(","))

tmp8130 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp8131 := PrimIntern(MakeString("bar!"))

tmp8132 := PrimCons(symabort, Nil)

tmp8133 := PrimCons(symabsolute, tmp8132)

tmp8134 := PrimCons(symabsvector, tmp8133)

tmp8135 := PrimCons(symabsvector_2, tmp8134)

tmp8136 := PrimCons(symaddress_1_6, tmp8135)

tmp8137 := PrimCons(sym_5_1address, tmp8136)

tmp8138 := PrimCons(symadjoin, tmp8137)

tmp8139 := PrimCons(symand, tmp8138)

tmp8140 := PrimCons(symappend, tmp8139)

tmp8141 := PrimCons(symarity, tmp8140)

tmp8142 := PrimCons(symassoc, tmp8141)

tmp8143 := PrimCons(symassertz, tmp8142)

tmp8144 := PrimCons(symasserta, tmp8143)

tmp8145 := PrimCons(symatom_2, tmp8144)

tmp8146 := PrimCons(symstep_2, tmp8145)

tmp8147 := PrimCons(symspy_2, tmp8146)

tmp8148 := PrimCons(tmp8131, tmp8147)

tmp8149 := PrimCons(symbootstrap, tmp8148)

tmp8150 := PrimCons(symboolean, tmp8149)

tmp8151 := PrimCons(symboolean_2, tmp8150)

tmp8152 := PrimCons(symbound_2, tmp8151)

tmp8153 := PrimCons(symbind, tmp8152)

tmp8154 := PrimCons(symclose, tmp8153)

tmp8155 := PrimCons(symcall, tmp8154)

tmp8156 := PrimCons(symcases, tmp8155)

tmp8157 := PrimCons(symcd, tmp8156)

tmp8158 := PrimCons(symcompile, tmp8157)

tmp8159 := PrimCons(symconcat, tmp8158)

tmp8160 := PrimCons(symcond, tmp8159)

tmp8161 := PrimCons(symcons, tmp8160)

tmp8162 := PrimCons(symcons_2, tmp8161)

tmp8163 := PrimCons(symcn, tmp8162)

tmp8164 := PrimCons(symctxt, tmp8163)

tmp8165 := PrimCons(symdatatypes, tmp8164)

tmp8166 := PrimCons(symdatatype, tmp8165)

tmp8167 := PrimCons(symdeclare, tmp8166)

tmp8168 := PrimCons(symdefprolog, tmp8167)

tmp8169 := PrimCons(symdefcc, tmp8168)

tmp8170 := PrimCons(symdefmacro, tmp8169)

tmp8171 := PrimCons(symdefine, tmp8170)

tmp8172 := PrimCons(symdefun, tmp8171)

tmp8173 := PrimCons(symdestroy, tmp8172)

tmp8174 := PrimCons(symdifference, tmp8173)

tmp8175 := PrimCons(symdo, tmp8174)

tmp8176 := PrimCons(symelement_2, tmp8175)

tmp8177 := PrimCons(symempty_2, tmp8176)

tmp8178 := PrimCons(symerror, tmp8177)

tmp8179 := PrimCons(symerror_1to_1string, tmp8178)

tmp8180 := PrimCons(symeval, tmp8179)

tmp8181 := PrimCons(symeval_1kl, tmp8180)

tmp8182 := PrimCons(symexception, tmp8181)

tmp8183 := PrimCons(symexternal, tmp8182)

tmp8184 := PrimCons(symexplode, tmp8183)

tmp8185 := PrimCons(symenable_1type_1theory, tmp8184)

tmp8186 := PrimCons(False, tmp8185)

tmp8187 := PrimCons(symfindall, tmp8186)

tmp8188 := PrimCons(symfactorise_2, tmp8187)

tmp8189 := PrimCons(symfactorise, tmp8188)

tmp8190 := PrimCons(symfail_1if, tmp8189)

tmp8191 := PrimCons(symfail, tmp8190)

tmp8192 := PrimCons(symfile, tmp8191)

tmp8193 := PrimCons(symfix, tmp8192)

tmp8194 := PrimCons(symforeign, tmp8193)

tmp8195 := PrimCons(symfork, tmp8194)

tmp8196 := PrimCons(symfresh, tmp8195)

tmp8197 := PrimCons(symfreeze, tmp8196)

tmp8198 := PrimCons(symfst, tmp8197)

tmp8199 := PrimCons(symfunction, tmp8198)

tmp8200 := PrimCons(symfn, tmp8199)

tmp8201 := PrimCons(symgensym, tmp8200)

tmp8202 := PrimCons(symget_1time, tmp8201)

tmp8203 := PrimCons(symget, tmp8202)

tmp8204 := PrimCons(symhash, tmp8203)

tmp8205 := PrimCons(symhdstr, tmp8204)

tmp8206 := PrimCons(symhdv, tmp8205)

tmp8207 := PrimCons(symhd, tmp8206)

tmp8208 := PrimCons(symhead, tmp8207)

tmp8209 := PrimCons(symhush_2, tmp8208)

tmp8210 := PrimCons(symhush_2, tmp8209)

tmp8211 := PrimCons(symif, tmp8210)

tmp8212 := PrimCons(symimplementation, tmp8211)

tmp8213 := PrimCons(syminternal, tmp8212)

tmp8214 := PrimCons(symin_1package, tmp8213)

tmp8215 := PrimCons(symin, tmp8214)

tmp8216 := PrimCons(symis_b, tmp8215)

tmp8217 := PrimCons(symis, tmp8216)

tmp8218 := PrimCons(symit, tmp8217)

tmp8219 := PrimCons(syminclude_1all_1but, tmp8218)

tmp8220 := PrimCons(syminclude, tmp8219)

tmp8221 := PrimCons(symincluded, tmp8220)

tmp8222 := PrimCons(syminput_7, tmp8221)

tmp8223 := PrimCons(syminput, tmp8222)

tmp8224 := PrimCons(syminteger_2, tmp8223)

tmp8225 := PrimCons(symintern, tmp8224)

tmp8226 := PrimCons(syminferences, tmp8225)

tmp8227 := PrimCons(symintersection, tmp8226)

tmp8228 := PrimCons(symis, tmp8227)

tmp8229 := PrimCons(symlanguage, tmp8228)

tmp8230 := PrimCons(symlambda, tmp8229)

tmp8231 := PrimCons(symlazy, tmp8230)

tmp8232 := PrimCons(symlet, tmp8231)

tmp8233 := PrimCons(symlength, tmp8232)

tmp8234 := PrimCons(symlimit, tmp8233)

tmp8235 := PrimCons(symlineread, tmp8234)

tmp8236 := PrimCons(symlist, tmp8235)

tmp8237 := PrimCons(symloaded, tmp8236)

tmp8238 := PrimCons(symload, tmp8237)

tmp8239 := PrimCons(symmake_1string, tmp8238)

tmp8240 := PrimCons(symmap, tmp8239)

tmp8241 := PrimCons(symmapcan, tmp8240)

tmp8242 := PrimCons(symmaxinferences, tmp8241)

tmp8243 := PrimCons(symmacroexpand, tmp8242)

tmp8244 := PrimCons(symmode, tmp8243)

tmp8245 := PrimCons(symnl, tmp8244)

tmp8246 := PrimCons(symnot, tmp8245)

tmp8247 := PrimCons(symnth, tmp8246)

tmp8248 := PrimCons(symnull, tmp8247)

tmp8249 := PrimCons(symnumber, tmp8248)

tmp8250 := PrimCons(symnumber_2, tmp8249)

tmp8251 := PrimCons(symn_1_6string, tmp8250)

tmp8252 := PrimCons(symoccurs_2, tmp8251)

tmp8253 := PrimCons(symoccurs_1check, tmp8252)

tmp8254 := PrimCons(symoccurrences, tmp8253)

tmp8255 := PrimCons(symopen, tmp8254)

tmp8256 := PrimCons(symoptimise_2, tmp8255)

tmp8257 := PrimCons(symoptimise, tmp8256)

tmp8258 := PrimCons(symor, tmp8257)

tmp8259 := PrimCons(symos, tmp8258)

tmp8260 := PrimCons(symout, tmp8259)

tmp8261 := PrimCons(symoutput, tmp8260)

tmp8262 := PrimCons(sympackage, tmp8261)

tmp8263 := PrimCons(symport, tmp8262)

tmp8264 := PrimCons(symporters, tmp8263)

tmp8265 := PrimCons(sympos, tmp8264)

tmp8266 := PrimCons(sympr, tmp8265)

tmp8267 := PrimCons(symprint, tmp8266)

tmp8268 := PrimCons(symprolog_1memory, tmp8267)

tmp8269 := PrimCons(symprofile, tmp8268)

tmp8270 := PrimCons(symprofile_1results, tmp8269)

tmp8271 := PrimCons(symprotect, tmp8270)

tmp8272 := PrimCons(symprolog_2, tmp8271)

tmp8273 := PrimCons(symps, tmp8272)

tmp8274 := PrimCons(sympreclude_1all_1but, tmp8273)

tmp8275 := PrimCons(sympreclude, tmp8274)

tmp8276 := PrimCons(symput, tmp8275)

tmp8277 := PrimCons(sympackage_2, tmp8276)

tmp8278 := PrimCons(symread_1from_1string_1unprocessed, tmp8277)

tmp8279 := PrimCons(symread_1from_1string, tmp8278)

tmp8280 := PrimCons(symread_1byte, tmp8279)

tmp8281 := PrimCons(symread_1file_1as_1string, tmp8280)

tmp8282 := PrimCons(symread_1file_1as_1bytelist, tmp8281)

tmp8283 := PrimCons(symread_1file, tmp8282)

tmp8284 := PrimCons(symreceive, tmp8283)

tmp8285 := PrimCons(symread, tmp8284)

tmp8286 := PrimCons(symrelease, tmp8285)

tmp8287 := PrimCons(symremove, tmp8286)

tmp8288 := PrimCons(symretract, tmp8287)

tmp8289 := PrimCons(symreverse, tmp8288)

tmp8290 := PrimCons(symrun, tmp8289)

tmp8291 := PrimCons(symstr, tmp8290)

tmp8292 := PrimCons(symsave, tmp8291)

tmp8293 := PrimCons(symset, tmp8292)

tmp8294 := PrimCons(symsimple_1error, tmp8293)

tmp8295 := PrimCons(symsnd, tmp8294)

tmp8296 := PrimCons(symspecialise, tmp8295)

tmp8297 := PrimCons(symspy, tmp8296)

tmp8298 := PrimCons(symsqts, tmp8297)

tmp8299 := PrimCons(symstep, tmp8298)

tmp8300 := PrimCons(symstoutput, tmp8299)

tmp8301 := PrimCons(symstinput, tmp8300)

tmp8302 := PrimCons(symstring, tmp8301)

tmp8303 := PrimCons(symstream, tmp8302)

tmp8304 := PrimCons(symstring_1_6n, tmp8303)

tmp8305 := PrimCons(symstring_2, tmp8304)

tmp8306 := PrimCons(symsubst, tmp8305)

tmp8307 := PrimCons(symsum, tmp8306)

tmp8308 := PrimCons(symstring_1_6symbol, tmp8307)

tmp8309 := PrimCons(symsymbol_2, tmp8308)

tmp8310 := PrimCons(symsymbol, tmp8309)

tmp8311 := PrimCons(symsystem_1S_2, tmp8310)

tmp8312 := PrimCons(symsynonyms, tmp8311)

tmp8313 := PrimCons(symsystemf, tmp8312)

tmp8314 := PrimCons(symtail, tmp8313)

tmp8315 := PrimCons(symtlv, tmp8314)

tmp8316 := PrimCons(symtlstr, tmp8315)

tmp8317 := PrimCons(symtl, tmp8316)

tmp8318 := PrimCons(symtc, tmp8317)

tmp8319 := PrimCons(symtc_2, tmp8318)

tmp8320 := PrimCons(symthaw, tmp8319)

tmp8321 := PrimCons(symtime, tmp8320)

tmp8322 := PrimCons(symtrack, tmp8321)

tmp8323 := PrimCons(symtracked, tmp8322)

tmp8324 := PrimCons(symtrap_1error, tmp8323)

tmp8325 := PrimCons(True, tmp8324)

tmp8326 := PrimCons(symtuple_2, tmp8325)

tmp8327 := PrimCons(symtype, tmp8326)

tmp8328 := PrimCons(symreturn, tmp8327)

tmp8329 := PrimCons(symunabsolute, tmp8328)

tmp8330 := PrimCons(symundefmacro, tmp8329)

tmp8331 := PrimCons(symunprofile, tmp8330)

tmp8332 := PrimCons(symunput, tmp8331)

tmp8333 := PrimCons(symunion, tmp8332)

tmp8334 := PrimCons(symunix, tmp8333)

tmp8335 := PrimCons(symunit, tmp8334)

tmp8336 := PrimCons(symuntrack, tmp8335)

tmp8337 := PrimCons(symunspecialise, tmp8336)

tmp8338 := PrimCons(symupdate_1lambda_1table, tmp8337)

tmp8339 := PrimCons(symu_b, tmp8338)

tmp8340 := PrimCons(symuserdefs, tmp8339)

tmp8341 := PrimCons(symvector_2, tmp8340)

tmp8342 := PrimCons(symvector, tmp8341)

tmp8343 := PrimCons(sym_5_1vector, tmp8342)

tmp8344 := PrimCons(symvector_1_6, tmp8343)

tmp8345 := PrimCons(symvalue, tmp8344)

tmp8346 := PrimCons(symvar_2, tmp8345)

tmp8347 := PrimCons(symvariable_2, tmp8346)

tmp8348 := PrimCons(symverified, tmp8347)

tmp8349 := PrimCons(symversion, tmp8348)

tmp8350 := PrimCons(symwhen, tmp8349)

tmp8351 := PrimCons(symwhere, tmp8350)

tmp8352 := PrimCons(symwrite_1byte, tmp8351)

tmp8353 := PrimCons(symwrite_1to_1file, tmp8352)

tmp8354 := PrimCons(symy_1or_1n_2, tmp8353)

tmp8355 := PrimCons(tmp8130, tmp8354)

tmp8356 := PrimCons(sym_6_6, tmp8355)

tmp8357 := PrimCons(sym_5, tmp8356)

tmp8358 := PrimCons(sym_5_a, tmp8357)

tmp8359 := PrimCons(sym_7, tmp8358)

tmp8360 := PrimCons(sym_d, tmp8359)

tmp8361 := PrimCons(sym_c, tmp8360)

tmp8362 := PrimCons(sym_1, tmp8361)

tmp8363 := PrimCons(sym_3, tmp8362)

tmp8364 := PrimCons(sym_5end_6, tmp8363)

tmp8365 := PrimCons(sym_5_b_6, tmp8364)

tmp8366 := PrimCons(sym_c_4, tmp8365)

tmp8367 := PrimCons(sym_a_a_6, tmp8366)

tmp8368 := PrimCons(sym_6, tmp8367)

tmp8369 := PrimCons(sym_6_a, tmp8368)

tmp8370 := PrimCons(sym_a, tmp8369)

tmp8371 := PrimCons(sym_a_a, tmp8370)

tmp8372 := PrimCons(sym_5e_6, tmp8371)

tmp8373 := PrimCons(sym_1_6, tmp8372)

tmp8374 := PrimCons(sym_5_1, tmp8373)

tmp8375 := PrimCons(sym_dhush_d, tmp8374)

tmp8376 := PrimCons(sym_dporters_d, tmp8375)

tmp8377 := PrimCons(sym_dport_d, tmp8376)

tmp8378 := PrimCons(sym_8s, tmp8377)

tmp8379 := PrimCons(sym_8p, tmp8378)

tmp8380 := PrimCons(sym_8v, tmp8379)

tmp8381 := PrimCons(sym_dproperty_1vector_d, tmp8380)

tmp8382 := PrimCons(sym_drelease_d, tmp8381)

tmp8383 := PrimCons(sym_dos_d, tmp8382)

tmp8384 := PrimCons(sym_dmacros_d, tmp8383)

tmp8385 := PrimCons(sym_dmaximum_1print_1sequence_1size_d, tmp8384)

tmp8386 := PrimCons(sym_dversion_d, tmp8385)

tmp8387 := PrimCons(sym_dhome_1directory_d, tmp8386)

tmp8388 := PrimCons(sym_dstoutput_d, tmp8387)

tmp8389 := PrimCons(sym_dstinput_d, tmp8388)

tmp8390 := PrimCons(sym_dimplementation_d, tmp8389)

tmp8391 := PrimCons(sym_dlanguage_d, tmp8390)

tmp8392 := PrimCons(sym__, tmp8391)

tmp8393 := PrimCons(tmp8129, tmp8392)

tmp8394 := PrimCons(tmp8128, tmp8393)

tmp8395 := PrimCons(tmp8127, tmp8394)

tmp8396 := PrimCons(tmp8126, tmp8395)

tmp8397 := PrimCons(sym_e_e, tmp8396)

tmp8398 := PrimCons(sym_5_1_1, tmp8397)

tmp8399 := PrimCons(sym_1_1_6, tmp8398)

tmp8400 := PrimCons(sym_i, tmp8399)

tmp8401 := PrimCons(sym_j, tmp8400)

tmp8402 := PrimCons(sym_b, tmp8401)

tmp8403 := PrimValue(sym_dproperty_1vector_d)

tmp8404 := Call(__e, PrimFunc(symput), symshen, symshen_4external_1symbols, tmp8402, tmp8403)


_ = tmp8404

tmp8405 := MakeNative(func(__e *ControlFlow) {
V5740 := __e.Get(1)
_ = V5740
tmp8406 := MakeNative(func(__e *ControlFlow) {
W5741 := __e.Get(1)
_ = W5741
tmp8414 := PrimEqual(W5741, MakeNumber(-1))

var ifres8411 Obj

if True == tmp8414 {
ifres8411 = True


} else {
tmp8413 := PrimEqual(W5741, MakeNumber(0))

var ifres8412 Obj

if True == tmp8413 {
ifres8412 = True


} else {
ifres8412 = False


}

ifres8411 = ifres8412


}

if True == ifres8411 {
__e.Return(Nil)
return
} else {
tmp8407 := PrimCons(V5740, Nil)

tmp8408 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp8407, W5741)


tmp8409 := Call(__e, PrimFunc(symeval_1kl), tmp8408)


__e.Return(PrimCons(V5740, tmp8409))
return


}


}, 1)

tmp8415 := Call(__e, PrimFunc(symarity), V5740)


__e.TailApply(tmp8406, tmp8415)
return


}, 1)

tmp8416 := Call(__e, ns2_1set, symshen_4lambda_1entry, tmp8405)


_ = tmp8416

tmp8417 := MakeNative(func(__e *ControlFlow) {
V5742 := __e.Get(1)
_ = V5742
tmp8418 := MakeNative(func(__e *ControlFlow) {
W5743 := __e.Get(1)
_ = W5743
tmp8419 := MakeNative(func(__e *ControlFlow) {
Z5745 := __e.Get(1)
_ = Z5745
__e.TailApply(PrimFunc(symshen_4tuple), Z5745)
return
}, 1)

tmp8420 := PrimCons(symshen_4tuple, tmp8419)

tmp8421 := MakeNative(func(__e *ControlFlow) {
Z5746 := __e.Get(1)
_ = Z5746
__e.TailApply(PrimFunc(symshen_4pvar), Z5746)
return
}, 1)

tmp8422 := PrimCons(symshen_4pvar, tmp8421)

tmp8423 := MakeNative(func(__e *ControlFlow) {
Z5747 := __e.Get(1)
_ = Z5747
__e.TailApply(PrimFunc(symshen_4print_1prolog_1vector), Z5747)
return
}, 1)

tmp8424 := PrimCons(symshen_4print_1prolog_1vector, tmp8423)

tmp8425 := MakeNative(func(__e *ControlFlow) {
Z5748 := __e.Get(1)
_ = Z5748
__e.TailApply(PrimFunc(symshen_4print_1freshterm), Z5748)
return
}, 1)

tmp8426 := PrimCons(symshen_4print_1freshterm, tmp8425)

tmp8427 := MakeNative(func(__e *ControlFlow) {
Z5749 := __e.Get(1)
_ = Z5749
__e.TailApply(PrimFunc(symshen_4printF), Z5749)
return
}, 1)

tmp8428 := PrimCons(symshen_4printF, tmp8427)

tmp8429 := PrimCons(tmp8428, W5743)

tmp8430 := PrimCons(tmp8426, tmp8429)

tmp8431 := PrimCons(tmp8424, tmp8430)

tmp8432 := PrimCons(tmp8422, tmp8431)

tmp8433 := PrimCons(tmp8420, tmp8432)

__e.Return(PrimSet(symshen_4_dlambdatable_d, tmp8433))
return


}, 1)

tmp8434 := MakeNative(func(__e *ControlFlow) {
Z5744 := __e.Get(1)
_ = Z5744
__e.TailApply(PrimFunc(symshen_4lambda_1entry), Z5744)
return
}, 1)

tmp8435 := Call(__e, PrimFunc(symmap), tmp8434, V5742)


__e.TailApply(tmp8418, tmp8435)
return


}, 1)

tmp8436 := Call(__e, ns2_1set, symshen_4build_1lambda_1table, tmp8417)


_ = tmp8436

tmp8437 := Call(__e, PrimFunc(symexternal), symshen)


__e.TailApply(PrimFunc(symshen_4build_1lambda_1table), tmp8437)
return




}, 0)

