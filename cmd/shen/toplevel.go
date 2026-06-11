package main

import . "github.com/tiancaiamao/shen-go/kl"

var TopLevelMain = MakeNative(func(__e *ControlFlow) {
tmp6752 := MakeNative(func(__e *ControlFlow) {
tmp6753 := Call(__e, PrimFunc(symshen_4credits))


_ = tmp6753

__e.TailApply(PrimFunc(symshen_4loop))
return


}, 0)

tmp6754 := Call(__e, ns2_1set, symshen_4repl, tmp6752)


_ = tmp6754

tmp6755 := MakeNative(func(__e *ControlFlow) {
tmp6756 := Call(__e, PrimFunc(symshen_4initialise__environment))


_ = tmp6756

tmp6757 := Call(__e, PrimFunc(symshen_4prompt))


_ = tmp6757

tmp6758 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4read_1evaluate_1print))
return
}, 0)

tmp6759 := MakeNative(func(__e *ControlFlow) {
Z5667 := __e.Get(1)
_ = Z5667
__e.TailApply(PrimFunc(symshen_4toplevel_1display_1exception), Z5667)
return
}, 1)

tmp6760 := Call(__e, try_1catch, tmp6758, tmp6759)


_ = tmp6760

__e.TailApply(PrimFunc(symshen_4loop))
return


}, 0)

tmp6761 := Call(__e, ns2_1set, symshen_4loop, tmp6755)


_ = tmp6761

tmp6762 := MakeNative(func(__e *ControlFlow) {
V5668 := __e.Get(1)
_ = V5668
tmp6763 := PrimErrorToString(V5668)

tmp6764 := Call(__e, PrimFunc(symstoutput))


tmp6765 := Call(__e, PrimFunc(sympr), tmp6763, tmp6764)


_ = tmp6765

__e.TailApply(PrimFunc(symnl), MakeNumber(0))
return


}, 1)

tmp6766 := Call(__e, ns2_1set, symshen_4toplevel_1display_1exception, tmp6762)


_ = tmp6766

tmp6767 := MakeNative(func(__e *ControlFlow) {
tmp6768 := Call(__e, PrimFunc(symstoutput))


tmp6769 := Call(__e, PrimFunc(sympr), MakeString("\nShen, www.shenlanguage.org, copyright (C) 2010-2024, Mark Tarver\n"), tmp6768)


_ = tmp6769

tmp6770 := PrimValue(sym_dversion_d)

tmp6771 := PrimValue(sym_dlanguage_d)

tmp6772 := PrimValue(sym_dimplementation_d)

tmp6773 := PrimValue(sym_drelease_d)

tmp6774 := Call(__e, PrimFunc(symshen_4app), tmp6773, MakeString("\n"), symshen_4a)


tmp6775 := PrimStringConcat(MakeString(" "), tmp6774)

tmp6776 := Call(__e, PrimFunc(symshen_4app), tmp6772, tmp6775, symshen_4a)


tmp6777 := PrimStringConcat(MakeString(", platform: "), tmp6776)

tmp6778 := Call(__e, PrimFunc(symshen_4app), tmp6771, tmp6777, symshen_4a)


tmp6779 := PrimStringConcat(MakeString(", language: "), tmp6778)

tmp6780 := Call(__e, PrimFunc(symshen_4app), tmp6770, tmp6779, symshen_4a)


tmp6781 := PrimStringConcat(MakeString("version: S"), tmp6780)

tmp6782 := Call(__e, PrimFunc(symstoutput))


tmp6783 := Call(__e, PrimFunc(sympr), tmp6781, tmp6782)


_ = tmp6783

tmp6784 := PrimValue(sym_dport_d)

tmp6785 := PrimValue(sym_dporters_d)

tmp6786 := Call(__e, PrimFunc(symshen_4app), tmp6785, MakeString("\n\n"), symshen_4a)


tmp6787 := PrimStringConcat(MakeString(", ported by "), tmp6786)

tmp6788 := Call(__e, PrimFunc(symshen_4app), tmp6784, tmp6787, symshen_4a)


tmp6789 := PrimStringConcat(MakeString("port "), tmp6788)

tmp6790 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp6789, tmp6790)
return


}, 0)

tmp6791 := Call(__e, ns2_1set, symshen_4credits, tmp6767)


_ = tmp6791

tmp6792 := MakeNative(func(__e *ControlFlow) {
tmp6793 := PrimSet(symshen_4_dcall_d, MakeNumber(0))

_ = tmp6793

__e.Return(PrimSet(symshen_4_dinfs_d, MakeNumber(0)))
return


}, 0)

tmp6794 := Call(__e, ns2_1set, symshen_4initialise__environment, tmp6792)


_ = tmp6794

tmp6795 := MakeNative(func(__e *ControlFlow) {
tmp6807 := PrimValue(symshen_4_dtc_d)

if True == tmp6807 {
tmp6796 := PrimValue(symshen_4_dhistory_d)

tmp6797 := Call(__e, PrimFunc(symlength), tmp6796)


tmp6798 := Call(__e, PrimFunc(symshen_4app), tmp6797, MakeString("+) "), symshen_4a)


tmp6799 := PrimStringConcat(MakeString("\n("), tmp6798)

tmp6800 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp6799, tmp6800)
return


} else {
tmp6801 := PrimValue(symshen_4_dhistory_d)

tmp6802 := Call(__e, PrimFunc(symlength), tmp6801)


tmp6803 := Call(__e, PrimFunc(symshen_4app), tmp6802, MakeString("-) "), symshen_4a)


tmp6804 := PrimStringConcat(MakeString("\n("), tmp6803)

tmp6805 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp6804, tmp6805)
return


}


}, 0)

tmp6808 := Call(__e, ns2_1set, symshen_4prompt, tmp6795)


_ = tmp6808

tmp6809 := MakeNative(func(__e *ControlFlow) {
tmp6810 := MakeNative(func(__e *ControlFlow) {
W5669 := __e.Get(1)
_ = W5669
tmp6811 := MakeNative(func(__e *ControlFlow) {
W5670 := __e.Get(1)
_ = W5670
tmp6812 := MakeNative(func(__e *ControlFlow) {
W5671 := __e.Get(1)
_ = W5671
tmp6813 := PrimValue(symshen_4_dtc_d)

__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), W5670, W5671, tmp6813)
return


}, 1)

tmp6814 := Call(__e, PrimFunc(symshen_4update_1history))


__e.TailApply(tmp6812, tmp6814)
return


}, 1)

tmp6815 := Call(__e, PrimFunc(symstinput))


tmp6816 := Call(__e, PrimFunc(symlineread), tmp6815)


tmp6817 := Call(__e, PrimFunc(symshen_4package_1user_1input), W5669, tmp6816)


__e.TailApply(tmp6811, tmp6817)
return


}, 1)

tmp6818 := PrimValue(symshen_4_dpackage_d)

__e.TailApply(tmp6810, tmp6818)
return


}, 0)

tmp6819 := Call(__e, ns2_1set, symshen_4read_1evaluate_1print, tmp6809)


_ = tmp6819

tmp6820 := MakeNative(func(__e *ControlFlow) {
V5672 := __e.Get(1)
_ = V5672
V5673 := __e.Get(2)
_ = V5673
tmp6827 := PrimEqual(symnull, V5672)

if True == tmp6827 {
__e.Return(V5673)
return
} else {
tmp6821 := MakeNative(func(__e *ControlFlow) {
W5674 := __e.Get(1)
_ = W5674
tmp6822 := MakeNative(func(__e *ControlFlow) {
W5675 := __e.Get(1)
_ = W5675
tmp6823 := MakeNative(func(__e *ControlFlow) {
Z5676 := __e.Get(1)
_ = Z5676
__e.TailApply(PrimFunc(symshen_4pui_1h), W5674, W5675, Z5676)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp6823, V5673)
return


}, 1)

tmp6824 := Call(__e, PrimFunc(symexternal), V5672)


__e.TailApply(tmp6822, tmp6824)
return


}, 1)

tmp6825 := PrimStr(V5672)

__e.TailApply(tmp6821, tmp6825)
return


}


}, 2)

tmp6828 := Call(__e, ns2_1set, symshen_4package_1user_1input, tmp6820)


_ = tmp6828

tmp6829 := MakeNative(func(__e *ControlFlow) {
V5681 := __e.Get(1)
_ = V5681
V5682 := __e.Get(2)
_ = V5682
V5683 := __e.Get(3)
_ = V5683
tmp6870 := PrimIsPair(V5683)

var ifres6857 Obj

if True == tmp6870 {
tmp6868 := PrimHead(V5683)

tmp6869 := PrimEqual(symfn, tmp6868)

var ifres6859 Obj

if True == tmp6869 {
tmp6866 := PrimTail(V5683)

tmp6867 := PrimIsPair(tmp6866)

var ifres6861 Obj

if True == tmp6867 {
tmp6863 := PrimTail(V5683)

tmp6864 := PrimTail(tmp6863)

tmp6865 := PrimEqual(Nil, tmp6864)

var ifres6862 Obj

if True == tmp6865 {
ifres6862 = True


} else {
ifres6862 = False


}

ifres6861 = ifres6862


} else {
ifres6861 = False


}

var ifres6860 Obj

if True == ifres6861 {
ifres6860 = True


} else {
ifres6860 = False


}

ifres6859 = ifres6860


} else {
ifres6859 = False


}

var ifres6858 Obj

if True == ifres6859 {
ifres6858 = True


} else {
ifres6858 = False


}

ifres6857 = ifres6858


} else {
ifres6857 = False


}

if True == ifres6857 {
tmp6835 := PrimTail(V5683)

tmp6836 := PrimHead(tmp6835)

tmp6837 := Call(__e, PrimFunc(symshen_4internal_2), tmp6836, V5681, V5682)


if True == tmp6837 {
tmp6830 := PrimTail(V5683)

tmp6831 := PrimHead(tmp6830)

tmp6832 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V5681, tmp6831)


tmp6833 := PrimCons(tmp6832, Nil)

__e.Return(PrimCons(symfn, tmp6833))
return


} else {
__e.Return(V5683)
return
}


} else {
tmp6855 := PrimIsPair(V5683)

if True == tmp6855 {
tmp6852 := PrimHead(V5683)

tmp6853 := Call(__e, PrimFunc(symshen_4internal_2), tmp6852, V5681, V5682)


if True == tmp6853 {
tmp6838 := PrimHead(V5683)

tmp6839 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V5681, tmp6838)


tmp6840 := MakeNative(func(__e *ControlFlow) {
Z5684 := __e.Get(1)
_ = Z5684
__e.TailApply(PrimFunc(symshen_4pui_1h), V5681, V5682, Z5684)
return
}, 1)

tmp6841 := PrimTail(V5683)

tmp6842 := Call(__e, PrimFunc(symmap), tmp6840, tmp6841)


__e.Return(PrimCons(tmp6839, tmp6842))
return


} else {
tmp6849 := PrimHead(V5683)

tmp6850 := PrimIsPair(tmp6849)

if True == tmp6850 {
tmp6843 := MakeNative(func(__e *ControlFlow) {
Z5685 := __e.Get(1)
_ = Z5685
__e.TailApply(PrimFunc(symshen_4pui_1h), V5681, V5682, Z5685)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp6843, V5683)
return


} else {
tmp6844 := PrimHead(V5683)

tmp6845 := MakeNative(func(__e *ControlFlow) {
Z5686 := __e.Get(1)
_ = Z5686
__e.TailApply(PrimFunc(symshen_4pui_1h), V5681, V5682, Z5686)
return
}, 1)

tmp6846 := PrimTail(V5683)

tmp6847 := Call(__e, PrimFunc(symmap), tmp6845, tmp6846)


__e.Return(PrimCons(tmp6844, tmp6847))
return


}


}


} else {
__e.Return(V5683)
return
}


}


}, 3)

tmp6871 := Call(__e, ns2_1set, symshen_4pui_1h, tmp6829)


_ = tmp6871

tmp6872 := MakeNative(func(__e *ControlFlow) {
tmp6873 := Call(__e, PrimFunc(symit))


tmp6874 := Call(__e, PrimFunc(symshen_4trim_1it), tmp6873)


tmp6875 := PrimValue(symshen_4_dhistory_d)

tmp6876 := PrimCons(tmp6874, tmp6875)

__e.Return(PrimSet(symshen_4_dhistory_d, tmp6876))
return


}, 0)

tmp6877 := Call(__e, ns2_1set, symshen_4update_1history, tmp6872)


_ = tmp6877

tmp6878 := MakeNative(func(__e *ControlFlow) {
V5687 := __e.Get(1)
_ = V5687
tmp6886 := Call(__e, PrimFunc(symshen_4_7string_2), V5687)


var ifres6881 Obj

if True == tmp6886 {
tmp6883 := Call(__e, PrimFunc(symhdstr), V5687)


tmp6884 := PrimStringToNumber(tmp6883)

tmp6885 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp6884)


var ifres6882 Obj

if True == tmp6885 {
ifres6882 = True


} else {
ifres6882 = False


}

ifres6881 = ifres6882


} else {
ifres6881 = False


}

if True == ifres6881 {
tmp6879 := PrimTailString(V5687)

__e.TailApply(PrimFunc(symshen_4trim_1it), tmp6879)
return


} else {
__e.Return(V5687)
return
}


}, 1)

tmp6887 := Call(__e, ns2_1set, symshen_4trim_1it, tmp6878)


_ = tmp6887

tmp6888 := MakeNative(func(__e *ControlFlow) {
V5706 := __e.Get(1)
_ = V5706
V5707 := __e.Get(2)
_ = V5707
V5708 := __e.Get(3)
_ = V5708
tmp7018 := PrimIsPair(V5706)

var ifres6987 Obj

if True == tmp7018 {
tmp7016 := PrimTail(V5706)

tmp7017 := PrimEqual(Nil, tmp7016)

var ifres6989 Obj

if True == tmp7017 {
tmp7015 := PrimIsPair(V5707)

var ifres6991 Obj

if True == tmp7015 {
tmp7013 := PrimHead(V5707)

tmp7014 := Call(__e, PrimFunc(symshen_4_7string_2), tmp7013)


var ifres6993 Obj

if True == tmp7014 {
tmp7010 := PrimHead(V5707)

tmp7011 := Call(__e, PrimFunc(symhdstr), tmp7010)


tmp7012 := PrimEqual(MakeString("!"), tmp7011)

var ifres6995 Obj

if True == tmp7012 {
tmp7007 := PrimHead(V5707)

tmp7008 := PrimTailString(tmp7007)

tmp7009 := Call(__e, PrimFunc(symshen_4_7string_2), tmp7008)


var ifres6997 Obj

if True == tmp7009 {
tmp7003 := PrimHead(V5707)

tmp7004 := PrimTailString(tmp7003)

tmp7005 := Call(__e, PrimFunc(symhdstr), tmp7004)


tmp7006 := PrimEqual(MakeString("!"), tmp7005)

var ifres6999 Obj

if True == tmp7006 {
tmp7001 := PrimTail(V5707)

tmp7002 := PrimIsPair(tmp7001)

var ifres7000 Obj

if True == tmp7002 {
ifres7000 = True


} else {
ifres7000 = False


}

ifres6999 = ifres7000


} else {
ifres6999 = False


}

var ifres6998 Obj

if True == ifres6999 {
ifres6998 = True


} else {
ifres6998 = False


}

ifres6997 = ifres6998


} else {
ifres6997 = False


}

var ifres6996 Obj

if True == ifres6997 {
ifres6996 = True


} else {
ifres6996 = False


}

ifres6995 = ifres6996


} else {
ifres6995 = False


}

var ifres6994 Obj

if True == ifres6995 {
ifres6994 = True


} else {
ifres6994 = False


}

ifres6993 = ifres6994


} else {
ifres6993 = False


}

var ifres6992 Obj

if True == ifres6993 {
ifres6992 = True


} else {
ifres6992 = False


}

ifres6991 = ifres6992


} else {
ifres6991 = False


}

var ifres6990 Obj

if True == ifres6991 {
ifres6990 = True


} else {
ifres6990 = False


}

ifres6989 = ifres6990


} else {
ifres6989 = False


}

var ifres6988 Obj

if True == ifres6989 {
ifres6988 = True


} else {
ifres6988 = False


}

ifres6987 = ifres6988


} else {
ifres6987 = False


}

if True == ifres6987 {
tmp6889 := MakeNative(func(__e *ControlFlow) {
W5709 := __e.Get(1)
_ = W5709
tmp6890 := MakeNative(func(__e *ControlFlow) {
W5710 := __e.Get(1)
_ = W5710
tmp6891 := MakeNative(func(__e *ControlFlow) {
W5711 := __e.Get(1)
_ = W5711
__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), W5709, W5710, V5708)
return
}, 1)

tmp6892 := PrimTail(V5707)

tmp6893 := PrimHead(tmp6892)

tmp6894 := Call(__e, PrimFunc(symshen_4app), tmp6893, MakeString("\n"), symshen_4a)


tmp6895 := Call(__e, PrimFunc(symstoutput))


tmp6896 := Call(__e, PrimFunc(sympr), tmp6894, tmp6895)


__e.TailApply(tmp6891, tmp6896)
return


}, 1)

tmp6897 := PrimTail(V5707)

tmp6898 := PrimHead(tmp6897)

tmp6899 := PrimTail(V5707)

tmp6900 := PrimCons(tmp6898, tmp6899)

tmp6901 := PrimSet(symshen_4_dhistory_d, tmp6900)

__e.TailApply(tmp6890, tmp6901)
return


}, 1)

tmp6902 := PrimTail(V5707)

tmp6903 := PrimHead(tmp6902)

tmp6904 := Call(__e, PrimFunc(symread_1from_1string), tmp6903)


__e.TailApply(tmp6889, tmp6904)
return


} else {
tmp6985 := PrimIsPair(V5706)

var ifres6969 Obj

if True == tmp6985 {
tmp6983 := PrimTail(V5706)

tmp6984 := PrimEqual(Nil, tmp6983)

var ifres6971 Obj

if True == tmp6984 {
tmp6982 := PrimIsPair(V5707)

var ifres6973 Obj

if True == tmp6982 {
tmp6980 := PrimHead(V5707)

tmp6981 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6980)


var ifres6975 Obj

if True == tmp6981 {
tmp6977 := PrimHead(V5707)

tmp6978 := Call(__e, PrimFunc(symhdstr), tmp6977)


tmp6979 := PrimEqual(MakeString("!"), tmp6978)

var ifres6976 Obj

if True == tmp6979 {
ifres6976 = True


} else {
ifres6976 = False


}

ifres6975 = ifres6976


} else {
ifres6975 = False


}

var ifres6974 Obj

if True == ifres6975 {
ifres6974 = True


} else {
ifres6974 = False


}

ifres6973 = ifres6974


} else {
ifres6973 = False


}

var ifres6972 Obj

if True == ifres6973 {
ifres6972 = True


} else {
ifres6972 = False


}

ifres6971 = ifres6972


} else {
ifres6971 = False


}

var ifres6970 Obj

if True == ifres6971 {
ifres6970 = True


} else {
ifres6970 = False


}

ifres6969 = ifres6970


} else {
ifres6969 = False


}

if True == ifres6969 {
tmp6905 := MakeNative(func(__e *ControlFlow) {
W5712 := __e.Get(1)
_ = W5712
tmp6906 := MakeNative(func(__e *ControlFlow) {
W5713 := __e.Get(1)
_ = W5713
tmp6907 := MakeNative(func(__e *ControlFlow) {
W5714 := __e.Get(1)
_ = W5714
tmp6908 := MakeNative(func(__e *ControlFlow) {
W5715 := __e.Get(1)
_ = W5715
tmp6909 := MakeNative(func(__e *ControlFlow) {
W5716 := __e.Get(1)
_ = W5716
__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), W5715, W5716, V5708)
return
}, 1)

tmp6910 := PrimTail(V5707)

tmp6911 := PrimCons(W5713, tmp6910)

tmp6912 := PrimSet(symshen_4_dhistory_d, tmp6911)

__e.TailApply(tmp6909, tmp6912)
return


}, 1)

tmp6913 := Call(__e, PrimFunc(symread_1from_1string), W5713)


__e.TailApply(tmp6908, tmp6913)
return


}, 1)

tmp6914 := Call(__e, PrimFunc(symshen_4app), W5713, MakeString("\n"), symshen_4a)


tmp6915 := Call(__e, PrimFunc(symstoutput))


tmp6916 := Call(__e, PrimFunc(sympr), tmp6914, tmp6915)


__e.TailApply(tmp6907, tmp6916)
return


}, 1)

tmp6917 := PrimHead(V5707)

tmp6918 := PrimTailString(tmp6917)

tmp6919 := PrimTail(V5707)

tmp6920 := Call(__e, PrimFunc(symshen_4use_1history), W5712, tmp6918, tmp6919)


__e.TailApply(tmp6906, tmp6920)
return


}, 1)

tmp6926 := PrimHead(V5707)

tmp6927 := PrimTailString(tmp6926)

tmp6928 := PrimEqual(tmp6927, MakeString(""))

var ifres6921 Obj

if True == tmp6928 {
ifres6921 = Nil


} else {
tmp6922 := PrimHead(V5707)

tmp6923 := PrimTailString(tmp6922)

tmp6924 := Call(__e, PrimFunc(symread_1from_1string), tmp6923)


tmp6925 := PrimHead(tmp6924)

ifres6921 = tmp6925


}

__e.TailApply(tmp6905, ifres6921)
return


} else {
tmp6967 := PrimIsPair(V5706)

var ifres6951 Obj

if True == tmp6967 {
tmp6965 := PrimTail(V5706)

tmp6966 := PrimEqual(Nil, tmp6965)

var ifres6953 Obj

if True == tmp6966 {
tmp6964 := PrimIsPair(V5707)

var ifres6955 Obj

if True == tmp6964 {
tmp6962 := PrimHead(V5707)

tmp6963 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6962)


var ifres6957 Obj

if True == tmp6963 {
tmp6959 := PrimHead(V5707)

tmp6960 := Call(__e, PrimFunc(symhdstr), tmp6959)


tmp6961 := PrimEqual(MakeString("%"), tmp6960)

var ifres6958 Obj

if True == tmp6961 {
ifres6958 = True


} else {
ifres6958 = False


}

ifres6957 = ifres6958


} else {
ifres6957 = False


}

var ifres6956 Obj

if True == ifres6957 {
ifres6956 = True


} else {
ifres6956 = False


}

ifres6955 = ifres6956


} else {
ifres6955 = False


}

var ifres6954 Obj

if True == ifres6955 {
ifres6954 = True


} else {
ifres6954 = False


}

ifres6953 = ifres6954


} else {
ifres6953 = False


}

var ifres6952 Obj

if True == ifres6953 {
ifres6952 = True


} else {
ifres6952 = False


}

ifres6951 = ifres6952


} else {
ifres6951 = False


}

if True == ifres6951 {
tmp6929 := MakeNative(func(__e *ControlFlow) {
W5717 := __e.Get(1)
_ = W5717
tmp6930 := MakeNative(func(__e *ControlFlow) {
W5718 := __e.Get(1)
_ = W5718
tmp6931 := MakeNative(func(__e *ControlFlow) {
W5719 := __e.Get(1)
_ = W5719
__e.TailApply(PrimFunc(symabort))
return
}, 1)

tmp6932 := PrimTail(V5707)

tmp6933 := PrimSet(symshen_4_dhistory_d, tmp6932)

__e.TailApply(tmp6931, tmp6933)
return


}, 1)

tmp6934 := PrimHead(V5707)

tmp6935 := PrimTailString(tmp6934)

tmp6936 := PrimTail(V5707)

tmp6937 := Call(__e, PrimFunc(symshen_4peek_1history), W5717, tmp6935, tmp6936)


__e.TailApply(tmp6930, tmp6937)
return


}, 1)

tmp6943 := PrimHead(V5707)

tmp6944 := PrimTailString(tmp6943)

tmp6945 := PrimEqual(tmp6944, MakeString(""))

var ifres6938 Obj

if True == tmp6945 {
ifres6938 = Nil


} else {
tmp6939 := PrimHead(V5707)

tmp6940 := PrimTailString(tmp6939)

tmp6941 := Call(__e, PrimFunc(symread_1from_1string), tmp6940)


tmp6942 := PrimHead(tmp6941)

ifres6938 = tmp6942


}

__e.TailApply(tmp6929, ifres6938)
return


} else {
tmp6949 := PrimEqual(True, V5708)

if True == tmp6949 {
__e.TailApply(PrimFunc(symshen_4check_1eval_1and_1print), V5706)
return
} else {
tmp6947 := PrimEqual(False, V5708)

if True == tmp6947 {
__e.TailApply(PrimFunc(symshen_4eval_1and_1print), V5706)
return
} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.evaluate-lineread")))
return
}


}


}


}


}


}, 3)

tmp7019 := Call(__e, ns2_1set, symshen_4evaluate_1lineread, tmp6888)


_ = tmp7019

tmp7020 := MakeNative(func(__e *ControlFlow) {
V5720 := __e.Get(1)
_ = V5720
V5721 := __e.Get(2)
_ = V5721
V5722 := __e.Get(3)
_ = V5722
tmp7026 := PrimIsInteger(V5720)

if True == tmp7026 {
tmp7021 := PrimNumberAdd(MakeNumber(1), V5720)

tmp7022 := Call(__e, PrimFunc(symreverse), V5722)


__e.TailApply(PrimFunc(symnth), tmp7021, tmp7022)
return


} else {
tmp7024 := PrimIsSymbol(V5720)

if True == tmp7024 {
__e.TailApply(PrimFunc(symshen_4string_1match), V5721, V5722)
return
} else {
__e.Return(PrimSimpleError(MakeString("! expects a number or a symbol\n")))
return
}


}


}, 3)

tmp7027 := Call(__e, ns2_1set, symshen_4use_1history, tmp7020)


_ = tmp7027

tmp7028 := MakeNative(func(__e *ControlFlow) {
V5723 := __e.Get(1)
_ = V5723
V5724 := __e.Get(2)
_ = V5724
V5725 := __e.Get(3)
_ = V5725
tmp7042 := PrimIsInteger(V5723)

if True == tmp7042 {
tmp7029 := PrimNumberAdd(MakeNumber(1), V5723)

tmp7030 := Call(__e, PrimFunc(symreverse), V5725)


tmp7031 := Call(__e, PrimFunc(symnth), tmp7029, tmp7030)


tmp7032 := Call(__e, PrimFunc(symshen_4app), tmp7031, MakeString(""), symshen_4a)


tmp7033 := PrimStringConcat(MakeString("\n"), tmp7032)

tmp7034 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp7033, tmp7034)
return


} else {
tmp7040 := PrimEqual(V5724, MakeString(""))

var ifres7037 Obj

if True == tmp7040 {
ifres7037 = True


} else {
tmp7039 := PrimIsSymbol(V5723)

var ifres7038 Obj

if True == tmp7039 {
ifres7038 = True


} else {
ifres7038 = False


}

ifres7037 = ifres7038


}

if True == ifres7037 {
tmp7035 := Call(__e, PrimFunc(symreverse), V5725)


__e.TailApply(PrimFunc(symshen_4recursive_1string_1match), MakeNumber(0), V5724, tmp7035)
return


} else {
__e.Return(PrimSimpleError(MakeString("% expects a number or a symbol\n")))
return
}


}


}, 3)

tmp7043 := Call(__e, ns2_1set, symshen_4peek_1history, tmp7028)


_ = tmp7043

tmp7044 := MakeNative(func(__e *ControlFlow) {
V5735 := __e.Get(1)
_ = V5735
V5736 := __e.Get(2)
_ = V5736
tmp7055 := PrimEqual(Nil, V5736)

if True == tmp7055 {
__e.Return(PrimSimpleError(MakeString("\ninput not found")))
return
} else {
tmp7053 := PrimIsPair(V5736)

var ifres7049 Obj

if True == tmp7053 {
tmp7051 := PrimHead(V5736)

tmp7052 := Call(__e, PrimFunc(symshen_4string_1prefix_2), V5735, tmp7051)


var ifres7050 Obj

if True == tmp7052 {
ifres7050 = True


} else {
ifres7050 = False


}

ifres7049 = ifres7050


} else {
ifres7049 = False


}

if True == ifres7049 {
__e.Return(PrimHead(V5736))
return
} else {
tmp7047 := PrimIsPair(V5736)

if True == tmp7047 {
tmp7045 := PrimTail(V5736)

__e.TailApply(PrimFunc(symshen_4string_1match), V5735, tmp7045)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.string-match")))
return
}


}


}


}, 2)

tmp7056 := Call(__e, ns2_1set, symshen_4string_1match, tmp7044)


_ = tmp7056

tmp7057 := MakeNative(func(__e *ControlFlow) {
V5744 := __e.Get(1)
_ = V5744
V5745 := __e.Get(2)
_ = V5745
tmp7094 := PrimEqual(MakeString(""), V5744)

if True == tmp7094 {
__e.Return(True)
return
} else {
tmp7092 := Call(__e, PrimFunc(symshen_4_7string_2), V5744)


var ifres7087 Obj

if True == tmp7092 {
tmp7089 := Call(__e, PrimFunc(symhdstr), V5744)


tmp7090 := PrimStringToNumber(tmp7089)

tmp7091 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp7090)


var ifres7088 Obj

if True == tmp7091 {
ifres7088 = True


} else {
ifres7088 = False


}

ifres7087 = ifres7088


} else {
ifres7087 = False


}

if True == ifres7087 {
tmp7058 := PrimTailString(V5744)

__e.TailApply(PrimFunc(symshen_4string_1prefix_2), tmp7058, V5745)
return


} else {
tmp7085 := Call(__e, PrimFunc(symshen_4_7string_2), V5745)


var ifres7080 Obj

if True == tmp7085 {
tmp7082 := Call(__e, PrimFunc(symhdstr), V5745)


tmp7083 := PrimStringToNumber(tmp7082)

tmp7084 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp7083)


var ifres7081 Obj

if True == tmp7084 {
ifres7081 = True


} else {
ifres7081 = False


}

ifres7080 = ifres7081


} else {
ifres7080 = False


}

if True == ifres7080 {
tmp7059 := PrimTailString(V5745)

__e.TailApply(PrimFunc(symshen_4string_1prefix_2), V5744, tmp7059)
return


} else {
tmp7078 := Call(__e, PrimFunc(symshen_4_7string_2), V5745)


var ifres7074 Obj

if True == tmp7078 {
tmp7076 := Call(__e, PrimFunc(symhdstr), V5745)


tmp7077 := PrimEqual(MakeString("("), tmp7076)

var ifres7075 Obj

if True == tmp7077 {
ifres7075 = True


} else {
ifres7075 = False


}

ifres7074 = ifres7075


} else {
ifres7074 = False


}

if True == ifres7074 {
tmp7060 := PrimTailString(V5745)

__e.TailApply(PrimFunc(symshen_4string_1prefix_2), V5744, tmp7060)
return


} else {
tmp7072 := Call(__e, PrimFunc(symshen_4_7string_2), V5744)


var ifres7064 Obj

if True == tmp7072 {
tmp7071 := Call(__e, PrimFunc(symshen_4_7string_2), V5745)


var ifres7066 Obj

if True == tmp7071 {
tmp7068 := Call(__e, PrimFunc(symhdstr), V5744)


tmp7069 := Call(__e, PrimFunc(symhdstr), V5745)


tmp7070 := PrimEqual(tmp7068, tmp7069)

var ifres7067 Obj

if True == tmp7070 {
ifres7067 = True


} else {
ifres7067 = False


}

ifres7066 = ifres7067


} else {
ifres7066 = False


}

var ifres7065 Obj

if True == ifres7066 {
ifres7065 = True


} else {
ifres7065 = False


}

ifres7064 = ifres7065


} else {
ifres7064 = False


}

if True == ifres7064 {
tmp7061 := PrimTailString(V5744)

tmp7062 := PrimTailString(V5745)

__e.TailApply(PrimFunc(symshen_4string_1prefix_2), tmp7061, tmp7062)
return


} else {
__e.Return(False)
return
}


}


}


}


}


}, 2)

tmp7095 := Call(__e, ns2_1set, symshen_4string_1prefix_2, tmp7057)


_ = tmp7095

tmp7096 := MakeNative(func(__e *ControlFlow) {
V5756 := __e.Get(1)
_ = V5756
V5757 := __e.Get(2)
_ = V5757
V5758 := __e.Get(3)
_ = V5758
tmp7111 := PrimEqual(Nil, V5758)

if True == tmp7111 {
__e.Return(symshen_4skip)
return
} else {
tmp7109 := PrimIsPair(V5758)

if True == tmp7109 {
tmp7104 := PrimHead(V5758)

tmp7105 := Call(__e, PrimFunc(symshen_4string_1prefix_2), V5757, tmp7104)


var ifres7097 Obj

if True == tmp7105 {
tmp7098 := PrimHead(V5758)

tmp7099 := Call(__e, PrimFunc(symshen_4app), tmp7098, MakeString("\n"), symshen_4a)


tmp7100 := PrimStringConcat(MakeString(". "), tmp7099)

tmp7101 := Call(__e, PrimFunc(symshen_4app), V5756, tmp7100, symshen_4a)


tmp7102 := Call(__e, PrimFunc(symstoutput))


tmp7103 := Call(__e, PrimFunc(sympr), tmp7101, tmp7102)


ifres7097 = tmp7103


} else {
ifres7097 = symshen_4skip


}

_ = ifres7097

tmp7106 := PrimNumberAdd(V5756, MakeNumber(1))

tmp7107 := PrimTail(V5758)

__e.TailApply(PrimFunc(symshen_4recursive_1string_1match), tmp7106, V5757, tmp7107)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.recursive-string-match")))
return
}


}


}, 3)

__e.TailApply(ns2_1set, symshen_4recursive_1string_1match, tmp7096)
return




}, 0)

