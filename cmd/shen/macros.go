package main

import . "github.com/pyrex41/shen-go/kl"

var MacrosMain = MakeNative(func(__e *ControlFlow) {
tmp8796 := MakeNative(func(__e *ControlFlow) {
V5783 := __e.Get(1)
_ = V5783
tmp8797 := MakeNative(func(__e *ControlFlow) {
W5784 := __e.Get(1)
_ = W5784
__e.TailApply(PrimFunc(symshen_4macroexpand_1h), V5783, W5784, W5784)
return
}, 1)

tmp8798 := MakeNative(func(__e *ControlFlow) {
Z5785 := __e.Get(1)
_ = Z5785
__e.Return(PrimTail(Z5785))
return
}, 1)

tmp8799 := PrimValue(sym_dmacros_d)

tmp8800 := Call(__e, PrimFunc(symmap), tmp8798, tmp8799)


__e.TailApply(tmp8797, tmp8800)
return


}, 1)

tmp8801 := Call(__e, ns2_1set, symmacroexpand, tmp8796)


_ = tmp8801

tmp8802 := MakeNative(func(__e *ControlFlow) {
V5794 := __e.Get(1)
_ = V5794
V5795 := __e.Get(2)
_ = V5795
V5796 := __e.Get(3)
_ = V5796
tmp8812 := PrimEqual(Nil, V5795)

if True == tmp8812 {
__e.Return(V5794)
return
} else {
tmp8810 := PrimIsPair(V5795)

if True == tmp8810 {
tmp8803 := MakeNative(func(__e *ControlFlow) {
W5797 := __e.Get(1)
_ = W5797
tmp8806 := PrimEqual(V5794, W5797)

if True == tmp8806 {
tmp8804 := PrimTail(V5795)

__e.TailApply(PrimFunc(symshen_4macroexpand_1h), V5794, tmp8804, V5796)
return


} else {
__e.TailApply(PrimFunc(symshen_4macroexpand_1h), W5797, V5796, V5796)
return
}


}, 1)

tmp8807 := PrimHead(V5795)

tmp8808 := Call(__e, PrimFunc(symshen_4walk), tmp8807, V5794)


__e.TailApply(tmp8803, tmp8808)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.macroexpand-h")))
return
}


}


}, 3)

tmp8813 := Call(__e, ns2_1set, symshen_4macroexpand_1h, tmp8802)


_ = tmp8813

tmp8814 := MakeNative(func(__e *ControlFlow) {
V5798 := __e.Get(1)
_ = V5798
V5799 := __e.Get(2)
_ = V5799
tmp8818 := PrimIsPair(V5799)

if True == tmp8818 {
tmp8815 := MakeNative(func(__e *ControlFlow) {
Z5800 := __e.Get(1)
_ = Z5800
__e.TailApply(PrimFunc(symshen_4walk), V5798, Z5800)
return
}, 1)

tmp8816 := Call(__e, PrimFunc(symmap), tmp8815, V5799)


__e.TailApply(V5798, tmp8816)
return


} else {
__e.TailApply(V5798, V5799)
return
}


}, 2)

tmp8819 := Call(__e, ns2_1set, symshen_4walk, tmp8814)


_ = tmp8819

tmp8820 := MakeNative(func(__e *ControlFlow) {
V5801 := __e.Get(1)
_ = V5801
tmp8821 := MakeNative(func(__e *ControlFlow) {
GoTo5802 := __e.Get(1)
_ = GoTo5802
tmp9124 := PrimIsPair(V5801)

if True == tmp9124 {
tmp8822 := MakeNative(func(__e *ControlFlow) {
Select5807 := __e.Get(1)
_ = Select5807
tmp8823 := MakeNative(func(__e *ControlFlow) {
Select5808 := __e.Get(1)
_ = Select5808
tmp9120 := PrimEqual(symdefmacro, Select5807)

var ifres9117 Obj

if True == tmp9120 {
tmp9119 := PrimIsPair(Select5808)

var ifres9118 Obj

if True == tmp9119 {
ifres9118 = True


} else {
ifres9118 = False


}

ifres9117 = ifres9118


} else {
ifres9117 = False


}

if True == ifres9117 {
tmp8824 := PrimHead(Select5808)

tmp8825 := PrimTail(Select5808)

__e.TailApply(PrimFunc(symshen_4process_1def), tmp8824, tmp8825)
return


} else {
tmp9115 := PrimEqual(symdefcc, Select5807)

if True == tmp9115 {
__e.TailApply(PrimFunc(symshen_4yacc_1_6shen), Select5808)
return
} else {
tmp9113 := PrimEqual(symu_b, Select5807)

var ifres9106 Obj

if True == tmp9113 {
tmp9112 := PrimIsPair(Select5808)

var ifres9108 Obj

if True == tmp9112 {
tmp9110 := PrimTail(Select5808)

tmp9111 := PrimEqual(Nil, tmp9110)

var ifres9109 Obj

if True == tmp9111 {
ifres9109 = True


} else {
ifres9109 = False


}

ifres9108 = ifres9109


} else {
ifres9108 = False


}

var ifres9107 Obj

if True == ifres9108 {
ifres9107 = True


} else {
ifres9107 = False


}

ifres9106 = ifres9107


} else {
ifres9106 = False


}

if True == ifres9106 {
tmp8826 := PrimHead(Select5808)

tmp8827 := Call(__e, PrimFunc(symshen_4make_1uppercase), tmp8826)


tmp8828 := PrimCons(tmp8827, Nil)

__e.Return(PrimCons(symprotect, tmp8828))
return


} else {
tmp9104 := PrimEqual(symerror, Select5807)

var ifres9101 Obj

if True == tmp9104 {
tmp9103 := PrimIsPair(Select5808)

var ifres9102 Obj

if True == tmp9103 {
ifres9102 = True


} else {
ifres9102 = False


}

ifres9101 = ifres9102


} else {
ifres9101 = False


}

if True == ifres9101 {
tmp8829 := PrimHead(Select5808)

tmp8830 := PrimTail(Select5808)

tmp8831 := Call(__e, PrimFunc(symshen_4mkstr), tmp8829, tmp8830)


tmp8832 := PrimCons(tmp8831, Nil)

__e.Return(PrimCons(symsimple_1error, tmp8832))
return


} else {
tmp9099 := PrimEqual(symoutput, Select5807)

var ifres9096 Obj

if True == tmp9099 {
tmp9098 := PrimIsPair(Select5808)

var ifres9097 Obj

if True == tmp9098 {
ifres9097 = True


} else {
ifres9097 = False


}

ifres9096 = ifres9097


} else {
ifres9096 = False


}

if True == ifres9096 {
tmp8833 := PrimHead(Select5808)

tmp8834 := PrimTail(Select5808)

tmp8835 := Call(__e, PrimFunc(symshen_4mkstr), tmp8833, tmp8834)


tmp8836 := PrimCons(symstoutput, Nil)

tmp8837 := PrimCons(tmp8836, Nil)

tmp8838 := PrimCons(tmp8835, tmp8837)

__e.Return(PrimCons(sympr, tmp8838))
return


} else {
tmp9094 := PrimEqual(sympr, Select5807)

var ifres9087 Obj

if True == tmp9094 {
tmp9093 := PrimIsPair(Select5808)

var ifres9089 Obj

if True == tmp9093 {
tmp9091 := PrimTail(Select5808)

tmp9092 := PrimEqual(Nil, tmp9091)

var ifres9090 Obj

if True == tmp9092 {
ifres9090 = True


} else {
ifres9090 = False


}

ifres9089 = ifres9090


} else {
ifres9089 = False


}

var ifres9088 Obj

if True == ifres9089 {
ifres9088 = True


} else {
ifres9088 = False


}

ifres9087 = ifres9088


} else {
ifres9087 = False


}

if True == ifres9087 {
tmp8839 := PrimHead(Select5808)

tmp8840 := PrimCons(symstoutput, Nil)

tmp8841 := PrimCons(tmp8840, Nil)

tmp8842 := PrimCons(tmp8839, tmp8841)

__e.Return(PrimCons(sympr, tmp8842))
return


} else {
tmp9085 := PrimEqual(symmake_1string, Select5807)

var ifres9082 Obj

if True == tmp9085 {
tmp9084 := PrimIsPair(Select5808)

var ifres9083 Obj

if True == tmp9084 {
ifres9083 = True


} else {
ifres9083 = False


}

ifres9082 = ifres9083


} else {
ifres9082 = False


}

if True == ifres9082 {
tmp8843 := PrimHead(Select5808)

tmp8844 := PrimTail(Select5808)

__e.TailApply(PrimFunc(symshen_4mkstr), tmp8843, tmp8844)
return


} else {
tmp9080 := PrimEqual(symlineread, Select5807)

var ifres9077 Obj

if True == tmp9080 {
tmp9079 := PrimEqual(Nil, Select5808)

var ifres9078 Obj

if True == tmp9079 {
ifres9078 = True


} else {
ifres9078 = False


}

ifres9077 = ifres9078


} else {
ifres9077 = False


}

if True == ifres9077 {
tmp8845 := PrimCons(symstinput, Nil)

tmp8846 := PrimCons(tmp8845, Nil)

__e.Return(PrimCons(symlineread, tmp8846))
return


} else {
tmp9075 := PrimEqual(syminput, Select5807)

var ifres9072 Obj

if True == tmp9075 {
tmp9074 := PrimEqual(Nil, Select5808)

var ifres9073 Obj

if True == tmp9074 {
ifres9073 = True


} else {
ifres9073 = False


}

ifres9072 = ifres9073


} else {
ifres9072 = False


}

if True == ifres9072 {
tmp8847 := PrimCons(symstinput, Nil)

tmp8848 := PrimCons(tmp8847, Nil)

__e.Return(PrimCons(syminput, tmp8848))
return


} else {
tmp9070 := PrimEqual(symread, Select5807)

var ifres9067 Obj

if True == tmp9070 {
tmp9069 := PrimEqual(Nil, Select5808)

var ifres9068 Obj

if True == tmp9069 {
ifres9068 = True


} else {
ifres9068 = False


}

ifres9067 = ifres9068


} else {
ifres9067 = False


}

if True == ifres9067 {
tmp8849 := PrimCons(symstinput, Nil)

tmp8850 := PrimCons(tmp8849, Nil)

__e.Return(PrimCons(symread, tmp8850))
return


} else {
tmp9065 := PrimEqual(syminput_7, Select5807)

var ifres9062 Obj

if True == tmp9065 {
tmp9064 := PrimIsPair(Select5808)

var ifres9063 Obj

if True == tmp9064 {
ifres9063 = True


} else {
ifres9063 = False


}

ifres9062 = ifres9063


} else {
ifres9062 = False


}

if True == ifres9062 {
__e.TailApply(PrimFunc(symshen_4process_1input_7), V5801)
return
} else {
tmp9060 := PrimEqual(symread_1byte, Select5807)

var ifres9057 Obj

if True == tmp9060 {
tmp9059 := PrimEqual(Nil, Select5808)

var ifres9058 Obj

if True == tmp9059 {
ifres9058 = True


} else {
ifres9058 = False


}

ifres9057 = ifres9058


} else {
ifres9057 = False


}

if True == ifres9057 {
__e.TailApply(PrimFunc(symshen_4process_1read_1byte))
return
} else {
tmp9055 := PrimEqual(symprolog_2, Select5807)

if True == tmp9055 {
__e.TailApply(PrimFunc(symshen_4call_1prolog), Select5808)
return
} else {
tmp9053 := PrimEqual(symdefprolog, Select5807)

var ifres9050 Obj

if True == tmp9053 {
tmp9052 := PrimIsPair(Select5808)

var ifres9051 Obj

if True == tmp9052 {
ifres9051 = True


} else {
ifres9051 = False


}

ifres9050 = ifres9051


} else {
ifres9050 = False


}

if True == ifres9050 {
tmp8851 := PrimHead(Select5808)

tmp8852 := PrimTail(Select5808)

__e.TailApply(PrimFunc(symshen_4compile_1prolog), tmp8851, tmp8852)
return


} else {
tmp9048 := PrimEqual(symdatatype, Select5807)

var ifres9045 Obj

if True == tmp9048 {
tmp9047 := PrimIsPair(Select5808)

var ifres9046 Obj

if True == tmp9047 {
ifres9046 = True


} else {
ifres9046 = False


}

ifres9045 = ifres9046


} else {
ifres9045 = False


}

if True == ifres9045 {
tmp8853 := PrimHead(Select5808)

tmp8854 := PrimTail(Select5808)

__e.TailApply(PrimFunc(symshen_4process_1datatype), tmp8853, tmp8854)
return


} else {
tmp9043 := PrimEqual(sym_8s, Select5807)

if True == tmp9043 {
__e.TailApply(PrimFunc(symshen_4process_1_8s), V5801)
return
} else {
tmp9041 := PrimEqual(symsynonyms, Select5807)

if True == tmp9041 {
__e.TailApply(PrimFunc(symshen_4process_1synonyms), Select5808)
return
} else {
tmp9039 := PrimEqual(symnl, Select5807)

var ifres9036 Obj

if True == tmp9039 {
tmp9038 := PrimEqual(Nil, Select5808)

var ifres9037 Obj

if True == tmp9038 {
ifres9037 = True


} else {
ifres9037 = False


}

ifres9036 = ifres9037


} else {
ifres9036 = False


}

if True == ifres9036 {
tmp8855 := PrimCons(MakeNumber(1), Nil)

__e.Return(PrimCons(symnl, tmp8855))
return


} else {
tmp9034 := PrimEqual(symlet, Select5807)

if True == tmp9034 {
__e.TailApply(PrimFunc(symshen_4process_1let), V5801)
return
} else {
tmp9032 := PrimEqual(sym_c_4, Select5807)

if True == tmp9032 {
__e.TailApply(PrimFunc(symshen_4process_1lambda), V5801)
return
} else {
tmp9030 := PrimEqual(symcases, Select5807)

if True == tmp9030 {
__e.TailApply(PrimFunc(symshen_4process_1cases), V5801)
return
} else {
tmp9028 := PrimEqual(symtime, Select5807)

var ifres9021 Obj

if True == tmp9028 {
tmp9027 := PrimIsPair(Select5808)

var ifres9023 Obj

if True == tmp9027 {
tmp9025 := PrimTail(Select5808)

tmp9026 := PrimEqual(Nil, tmp9025)

var ifres9024 Obj

if True == tmp9026 {
ifres9024 = True


} else {
ifres9024 = False


}

ifres9023 = ifres9024


} else {
ifres9023 = False


}

var ifres9022 Obj

if True == ifres9023 {
ifres9022 = True


} else {
ifres9022 = False


}

ifres9021 = ifres9022


} else {
ifres9021 = False


}

if True == ifres9021 {
tmp8856 := PrimHead(Select5808)

__e.TailApply(PrimFunc(symshen_4process_1time), tmp8856)
return


} else {
tmp9019 := PrimEqual(symput, Select5807)

var ifres9001 Obj

if True == tmp9019 {
tmp9018 := PrimIsPair(Select5808)

var ifres9003 Obj

if True == tmp9018 {
tmp9016 := PrimTail(Select5808)

tmp9017 := PrimIsPair(tmp9016)

var ifres9005 Obj

if True == tmp9017 {
tmp9013 := PrimTail(Select5808)

tmp9014 := PrimTail(tmp9013)

tmp9015 := PrimIsPair(tmp9014)

var ifres9007 Obj

if True == tmp9015 {
tmp9009 := PrimTail(Select5808)

tmp9010 := PrimTail(tmp9009)

tmp9011 := PrimTail(tmp9010)

tmp9012 := PrimEqual(Nil, tmp9011)

var ifres9008 Obj

if True == tmp9012 {
ifres9008 = True


} else {
ifres9008 = False


}

ifres9007 = ifres9008


} else {
ifres9007 = False


}

var ifres9006 Obj

if True == ifres9007 {
ifres9006 = True


} else {
ifres9006 = False


}

ifres9005 = ifres9006


} else {
ifres9005 = False


}

var ifres9004 Obj

if True == ifres9005 {
ifres9004 = True


} else {
ifres9004 = False


}

ifres9003 = ifres9004


} else {
ifres9003 = False


}

var ifres9002 Obj

if True == ifres9003 {
ifres9002 = True


} else {
ifres9002 = False


}

ifres9001 = ifres9002


} else {
ifres9001 = False


}

if True == ifres9001 {
tmp8857 := PrimHead(Select5808)

tmp8858 := PrimTail(Select5808)

tmp8859 := PrimHead(tmp8858)

tmp8860 := PrimTail(Select5808)

tmp8861 := PrimTail(tmp8860)

tmp8862 := PrimHead(tmp8861)

tmp8863 := PrimCons(sym_dproperty_1vector_d, Nil)

tmp8864 := PrimCons(symvalue, tmp8863)

tmp8865 := PrimCons(tmp8864, Nil)

tmp8866 := PrimCons(tmp8862, tmp8865)

tmp8867 := PrimCons(tmp8859, tmp8866)

tmp8868 := PrimCons(tmp8857, tmp8867)

__e.Return(PrimCons(symput, tmp8868))
return


} else {
tmp8999 := PrimEqual(symget, Select5807)

var ifres8987 Obj

if True == tmp8999 {
tmp8998 := PrimIsPair(Select5808)

var ifres8989 Obj

if True == tmp8998 {
tmp8996 := PrimTail(Select5808)

tmp8997 := PrimIsPair(tmp8996)

var ifres8991 Obj

if True == tmp8997 {
tmp8993 := PrimTail(Select5808)

tmp8994 := PrimTail(tmp8993)

tmp8995 := PrimEqual(Nil, tmp8994)

var ifres8992 Obj

if True == tmp8995 {
ifres8992 = True


} else {
ifres8992 = False


}

ifres8991 = ifres8992


} else {
ifres8991 = False


}

var ifres8990 Obj

if True == ifres8991 {
ifres8990 = True


} else {
ifres8990 = False


}

ifres8989 = ifres8990


} else {
ifres8989 = False


}

var ifres8988 Obj

if True == ifres8989 {
ifres8988 = True


} else {
ifres8988 = False


}

ifres8987 = ifres8988


} else {
ifres8987 = False


}

if True == ifres8987 {
tmp8869 := PrimHead(Select5808)

tmp8870 := PrimTail(Select5808)

tmp8871 := PrimHead(tmp8870)

tmp8872 := PrimCons(sym_dproperty_1vector_d, Nil)

tmp8873 := PrimCons(symvalue, tmp8872)

tmp8874 := PrimCons(tmp8873, Nil)

tmp8875 := PrimCons(tmp8871, tmp8874)

tmp8876 := PrimCons(tmp8869, tmp8875)

__e.Return(PrimCons(symget, tmp8876))
return


} else {
tmp8985 := PrimEqual(symunput, Select5807)

var ifres8973 Obj

if True == tmp8985 {
tmp8984 := PrimIsPair(Select5808)

var ifres8975 Obj

if True == tmp8984 {
tmp8982 := PrimTail(Select5808)

tmp8983 := PrimIsPair(tmp8982)

var ifres8977 Obj

if True == tmp8983 {
tmp8979 := PrimTail(Select5808)

tmp8980 := PrimTail(tmp8979)

tmp8981 := PrimEqual(Nil, tmp8980)

var ifres8978 Obj

if True == tmp8981 {
ifres8978 = True


} else {
ifres8978 = False


}

ifres8977 = ifres8978


} else {
ifres8977 = False


}

var ifres8976 Obj

if True == ifres8977 {
ifres8976 = True


} else {
ifres8976 = False


}

ifres8975 = ifres8976


} else {
ifres8975 = False


}

var ifres8974 Obj

if True == ifres8975 {
ifres8974 = True


} else {
ifres8974 = False


}

ifres8973 = ifres8974


} else {
ifres8973 = False


}

if True == ifres8973 {
tmp8877 := PrimHead(Select5808)

tmp8878 := PrimTail(Select5808)

tmp8879 := PrimHead(tmp8878)

tmp8880 := PrimCons(sym_dproperty_1vector_d, Nil)

tmp8881 := PrimCons(symvalue, tmp8880)

tmp8882 := PrimCons(tmp8881, Nil)

tmp8883 := PrimCons(tmp8879, tmp8882)

tmp8884 := PrimCons(tmp8877, tmp8883)

__e.Return(PrimCons(symunput, tmp8884))
return


} else {
tmp8971 := PrimEqual(symshen_4_8c, Select5807)

var ifres8964 Obj

if True == tmp8971 {
tmp8970 := PrimIsPair(Select5808)

var ifres8966 Obj

if True == tmp8970 {
tmp8968 := PrimTail(Select5808)

tmp8969 := PrimEqual(Nil, tmp8968)

var ifres8967 Obj

if True == tmp8969 {
ifres8967 = True


} else {
ifres8967 = False


}

ifres8966 = ifres8967


} else {
ifres8966 = False


}

var ifres8965 Obj

if True == ifres8966 {
ifres8965 = True


} else {
ifres8965 = False


}

ifres8964 = ifres8965


} else {
ifres8964 = False


}

if True == ifres8964 {
tmp8885 := PrimHead(Select5808)

__e.TailApply(PrimFunc(symshen_4rcons__form), tmp8885)
return


} else {
tmp8886 := MakeNative(func(__e *ControlFlow) {
GoTo5803 := __e.Get(1)
_ = GoTo5803
tmp8933 := PrimEqual(symshen_4_8ch, Select5807)

if True == tmp8933 {
tmp8931 := PrimIsPair(Select5808)

if True == tmp8931 {
tmp8887 := MakeNative(func(__e *ControlFlow) {
Select5805 := __e.Get(1)
_ = Select5805
tmp8888 := MakeNative(func(__e *ControlFlow) {
Select5806 := __e.Get(1)
_ = Select5806
tmp8927 := PrimIsPair(Select5805)

var ifres8903 Obj

if True == tmp8927 {
tmp8925 := PrimTail(Select5805)

tmp8926 := PrimIsPair(tmp8925)

var ifres8905 Obj

if True == tmp8926 {
tmp8922 := PrimTail(Select5805)

tmp8923 := PrimTail(tmp8922)

tmp8924 := PrimIsPair(tmp8923)

var ifres8907 Obj

if True == tmp8924 {
tmp8918 := PrimTail(Select5805)

tmp8919 := PrimTail(tmp8918)

tmp8920 := PrimTail(tmp8919)

tmp8921 := PrimEqual(Nil, tmp8920)

var ifres8909 Obj

if True == tmp8921 {
tmp8917 := PrimEqual(Nil, Select5806)

var ifres8911 Obj

if True == tmp8917 {
tmp8913 := PrimTail(Select5805)

tmp8914 := PrimHead(tmp8913)

tmp8915 := PrimIntern(MakeString(":"))

tmp8916 := PrimEqual(tmp8914, tmp8915)

var ifres8912 Obj

if True == tmp8916 {
ifres8912 = True


} else {
ifres8912 = False


}

ifres8911 = ifres8912


} else {
ifres8911 = False


}

var ifres8910 Obj

if True == ifres8911 {
ifres8910 = True


} else {
ifres8910 = False


}

ifres8909 = ifres8910


} else {
ifres8909 = False


}

var ifres8908 Obj

if True == ifres8909 {
ifres8908 = True


} else {
ifres8908 = False


}

ifres8907 = ifres8908


} else {
ifres8907 = False


}

var ifres8906 Obj

if True == ifres8907 {
ifres8906 = True


} else {
ifres8906 = False


}

ifres8905 = ifres8906


} else {
ifres8905 = False


}

var ifres8904 Obj

if True == ifres8905 {
ifres8904 = True


} else {
ifres8904 = False


}

ifres8903 = ifres8904


} else {
ifres8903 = False


}

if True == ifres8903 {
tmp8889 := PrimHead(Select5805)

tmp8890 := PrimTail(Select5805)

tmp8891 := PrimHead(tmp8890)

tmp8892 := PrimTail(Select5805)

tmp8893 := PrimTail(tmp8892)

tmp8894 := PrimCons(sym_7, tmp8893)

tmp8895 := PrimCons(tmp8894, Nil)

tmp8896 := PrimCons(tmp8891, tmp8895)

tmp8897 := PrimCons(tmp8889, tmp8896)

tmp8898 := PrimCons(tmp8897, Nil)

tmp8899 := PrimCons(sym_1, tmp8898)

__e.TailApply(PrimFunc(symshen_4cons_1form_1respect_1modes), tmp8899)
return


} else {
tmp8901 := PrimEqual(Nil, Select5806)

if True == tmp8901 {
__e.TailApply(PrimFunc(symshen_4cons_1form_1respect_1modes), Select5805)
return
} else {
__e.TailApply(PrimFunc(symthaw), GoTo5803)
return
}


}


}, 1)

tmp8928 := PrimTail(Select5808)

__e.TailApply(tmp8888, tmp8928)
return


}, 1)

tmp8929 := PrimHead(Select5808)

__e.TailApply(tmp8887, tmp8929)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5803)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5803)
return
}


}, 1)

tmp8934 := MakeNative(func(__e *ControlFlow) {
tmp8962 := PrimIsPair(Select5808)

var ifres8942 Obj

if True == tmp8962 {
tmp8960 := PrimTail(Select5808)

tmp8961 := PrimIsPair(tmp8960)

var ifres8944 Obj

if True == tmp8961 {
tmp8957 := PrimTail(Select5808)

tmp8958 := PrimTail(tmp8957)

tmp8959 := PrimIsPair(tmp8958)

var ifres8946 Obj

if True == tmp8959 {
tmp8948 := PrimCons(symdo, Nil)

tmp8949 := PrimCons(sym_d, tmp8948)

tmp8950 := PrimCons(sym_7, tmp8949)

tmp8951 := PrimCons(symor, tmp8950)

tmp8952 := PrimCons(symand, tmp8951)

tmp8953 := PrimCons(symappend, tmp8952)

tmp8954 := PrimCons(sym_8v, tmp8953)

tmp8955 := PrimCons(sym_8p, tmp8954)

tmp8956 := Call(__e, PrimFunc(symelement_2), Select5807, tmp8955)


var ifres8947 Obj

if True == tmp8956 {
ifres8947 = True


} else {
ifres8947 = False


}

ifres8946 = ifres8947


} else {
ifres8946 = False


}

var ifres8945 Obj

if True == ifres8946 {
ifres8945 = True


} else {
ifres8945 = False


}

ifres8944 = ifres8945


} else {
ifres8944 = False


}

var ifres8943 Obj

if True == ifres8944 {
ifres8943 = True


} else {
ifres8943 = False


}

ifres8942 = ifres8943


} else {
ifres8942 = False


}

if True == ifres8942 {
tmp8935 := PrimHead(Select5808)

tmp8936 := PrimTail(Select5808)

tmp8937 := PrimCons(Select5807, tmp8936)

tmp8938 := Call(__e, PrimFunc(symshen_4process_1assoc), tmp8937)


tmp8939 := PrimCons(tmp8938, Nil)

tmp8940 := PrimCons(tmp8935, tmp8939)

__e.Return(PrimCons(Select5807, tmp8940))
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5802)
return
}


}, 0)

__e.TailApply(tmp8886, tmp8934)
return


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}, 1)

tmp9121 := PrimTail(V5801)

__e.TailApply(tmp8823, tmp9121)
return


}, 1)

tmp9122 := PrimHead(V5801)

__e.TailApply(tmp8822, tmp9122)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5802)
return
}


}, 1)

tmp9125 := MakeNative(func(__e *ControlFlow) {
__e.Return(V5801)
return
}, 0)

__e.TailApply(tmp8821, tmp9125)
return


}, 1)

tmp9126 := Call(__e, ns2_1set, symshen_4macros, tmp8820)


_ = tmp9126

tmp9127 := MakeNative(func(__e *ControlFlow) {
V5809 := __e.Get(1)
_ = V5809
tmp9128 := MakeNative(func(__e *ControlFlow) {
GoTo5810 := __e.Get(1)
_ = GoTo5810
tmp9155 := PrimIsPair(V5809)

if True == tmp9155 {
tmp9129 := MakeNative(func(__e *ControlFlow) {
Select5815 := __e.Get(1)
_ = Select5815
tmp9151 := PrimHead(V5809)

tmp9152 := PrimEqual(syminput_7, tmp9151)

if True == tmp9152 {
tmp9149 := PrimIsPair(Select5815)

if True == tmp9149 {
tmp9130 := MakeNative(func(__e *ControlFlow) {
Select5813 := __e.Get(1)
_ = Select5813
tmp9131 := MakeNative(func(__e *ControlFlow) {
Select5814 := __e.Get(1)
_ = Select5814
tmp9145 := PrimEqual(Nil, Select5814)

if True == tmp9145 {
tmp9132 := Call(__e, PrimFunc(symshen_4rcons__form), Select5813)


tmp9133 := PrimCons(symstinput, Nil)

tmp9134 := PrimCons(tmp9133, Nil)

tmp9135 := PrimCons(tmp9132, tmp9134)

__e.Return(PrimCons(symshen_4input_1h_7, tmp9135))
return


} else {
tmp9143 := PrimIsPair(Select5814)

var ifres9139 Obj

if True == tmp9143 {
tmp9141 := PrimTail(Select5814)

tmp9142 := PrimEqual(Nil, tmp9141)

var ifres9140 Obj

if True == tmp9142 {
ifres9140 = True


} else {
ifres9140 = False


}

ifres9139 = ifres9140


} else {
ifres9139 = False


}

if True == ifres9139 {
tmp9136 := Call(__e, PrimFunc(symshen_4rcons__form), Select5813)


tmp9137 := PrimCons(tmp9136, Select5814)

__e.Return(PrimCons(symshen_4input_1h_7, tmp9137))
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5810)
return
}


}


}, 1)

tmp9146 := PrimTail(Select5815)

__e.TailApply(tmp9131, tmp9146)
return


}, 1)

tmp9147 := PrimHead(Select5815)

__e.TailApply(tmp9130, tmp9147)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5810)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5810)
return
}


}, 1)

tmp9153 := PrimTail(V5809)

__e.TailApply(tmp9129, tmp9153)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5810)
return
}


}, 1)

tmp9156 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimSimpleError(MakeString("partial function shen.process-input+")))
return
}, 0)

__e.TailApply(tmp9128, tmp9156)
return


}, 1)

tmp9157 := Call(__e, ns2_1set, symshen_4process_1input_7, tmp9127)


_ = tmp9157

tmp9158 := MakeNative(func(__e *ControlFlow) {
V5816 := __e.Get(1)
_ = V5816
tmp9159 := MakeNative(func(__e *ControlFlow) {
GoTo5817 := __e.Get(1)
_ = GoTo5817
tmp9193 := PrimIsPair(V5816)

if True == tmp9193 {
tmp9160 := MakeNative(func(__e *ControlFlow) {
Select5818 := __e.Get(1)
_ = Select5818
tmp9161 := MakeNative(func(__e *ControlFlow) {
Select5819 := __e.Get(1)
_ = Select5819
tmp9189 := PrimEqual(sym_7, Select5818)

var ifres9182 Obj

if True == tmp9189 {
tmp9188 := PrimIsPair(Select5819)

var ifres9184 Obj

if True == tmp9188 {
tmp9186 := PrimTail(Select5819)

tmp9187 := PrimEqual(Nil, tmp9186)

var ifres9185 Obj

if True == tmp9187 {
ifres9185 = True


} else {
ifres9185 = False


}

ifres9184 = ifres9185


} else {
ifres9184 = False


}

var ifres9183 Obj

if True == ifres9184 {
ifres9183 = True


} else {
ifres9183 = False


}

ifres9182 = ifres9183


} else {
ifres9182 = False


}

if True == ifres9182 {
tmp9162 := PrimHead(Select5819)

tmp9163 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), tmp9162)


tmp9164 := PrimCons(tmp9163, Nil)

__e.Return(PrimCons(sym_7, tmp9164))
return


} else {
tmp9180 := PrimEqual(sym_1, Select5818)

var ifres9173 Obj

if True == tmp9180 {
tmp9179 := PrimIsPair(Select5819)

var ifres9175 Obj

if True == tmp9179 {
tmp9177 := PrimTail(Select5819)

tmp9178 := PrimEqual(Nil, tmp9177)

var ifres9176 Obj

if True == tmp9178 {
ifres9176 = True


} else {
ifres9176 = False


}

ifres9175 = ifres9176


} else {
ifres9175 = False


}

var ifres9174 Obj

if True == ifres9175 {
ifres9174 = True


} else {
ifres9174 = False


}

ifres9173 = ifres9174


} else {
ifres9173 = False


}

if True == ifres9173 {
tmp9165 := PrimHead(Select5819)

tmp9166 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), tmp9165)


tmp9167 := PrimCons(tmp9166, Nil)

__e.Return(PrimCons(sym_1, tmp9167))
return


} else {
tmp9168 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), Select5818)


tmp9169 := Call(__e, PrimFunc(symshen_4cons_1form_1respect_1modes), Select5819)


tmp9170 := PrimCons(tmp9169, Nil)

tmp9171 := PrimCons(tmp9168, tmp9170)

__e.Return(PrimCons(symcons, tmp9171))
return


}


}


}, 1)

tmp9190 := PrimTail(V5816)

__e.TailApply(tmp9161, tmp9190)
return


}, 1)

tmp9191 := PrimHead(V5816)

__e.TailApply(tmp9160, tmp9191)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5817)
return
}


}, 1)

tmp9194 := MakeNative(func(__e *ControlFlow) {
__e.Return(V5816)
return
}, 0)

__e.TailApply(tmp9159, tmp9194)
return


}, 1)

tmp9195 := Call(__e, ns2_1set, symshen_4cons_1form_1respect_1modes, tmp9158)


_ = tmp9195

tmp9196 := MakeNative(func(__e *ControlFlow) {
V5820 := __e.Get(1)
_ = V5820
V5821 := __e.Get(2)
_ = V5821
tmp9197 := MakeNative(func(__e *ControlFlow) {
W5822 := __e.Get(1)
_ = W5822
tmp9198 := MakeNative(func(__e *ControlFlow) {
W5823 := __e.Get(1)
_ = W5823
tmp9199 := MakeNative(func(__e *ControlFlow) {
W5824 := __e.Get(1)
_ = W5824
__e.Return(V5820)
return
}, 1)

tmp9200 := Call(__e, PrimFunc(symfn), V5820)


tmp9201 := Call(__e, PrimFunc(symshen_4record_1macro), V5820, tmp9200)


__e.TailApply(tmp9199, tmp9201)
return


}, 1)

tmp9202 := Call(__e, PrimFunc(symappend), V5821, W5822)


tmp9203 := PrimCons(V5820, tmp9202)

tmp9204 := PrimCons(symdefine, tmp9203)

tmp9205 := Call(__e, PrimFunc(symeval), tmp9204)


__e.TailApply(tmp9198, tmp9205)
return


}, 1)

tmp9206 := PrimCons(symX, Nil)

tmp9207 := PrimCons(sym_1_6, tmp9206)

tmp9208 := PrimCons(symX, tmp9207)

__e.TailApply(tmp9197, tmp9208)
return


}, 2)

tmp9209 := Call(__e, ns2_1set, symshen_4process_1def, tmp9196)


_ = tmp9209

tmp9210 := MakeNative(func(__e *ControlFlow) {
V5825 := __e.Get(1)
_ = V5825
tmp9250 := PrimIsPair(V5825)

var ifres9224 Obj

if True == tmp9250 {
tmp9248 := PrimHead(V5825)

tmp9249 := PrimEqual(symlet, tmp9248)

var ifres9226 Obj

if True == tmp9249 {
tmp9246 := PrimTail(V5825)

tmp9247 := PrimIsPair(tmp9246)

var ifres9228 Obj

if True == tmp9247 {
tmp9243 := PrimTail(V5825)

tmp9244 := PrimTail(tmp9243)

tmp9245 := PrimIsPair(tmp9244)

var ifres9230 Obj

if True == tmp9245 {
tmp9239 := PrimTail(V5825)

tmp9240 := PrimTail(tmp9239)

tmp9241 := PrimTail(tmp9240)

tmp9242 := PrimIsPair(tmp9241)

var ifres9232 Obj

if True == tmp9242 {
tmp9234 := PrimTail(V5825)

tmp9235 := PrimTail(tmp9234)

tmp9236 := PrimTail(tmp9235)

tmp9237 := PrimTail(tmp9236)

tmp9238 := PrimIsPair(tmp9237)

var ifres9233 Obj

if True == tmp9238 {
ifres9233 = True


} else {
ifres9233 = False


}

ifres9232 = ifres9233


} else {
ifres9232 = False


}

var ifres9231 Obj

if True == ifres9232 {
ifres9231 = True


} else {
ifres9231 = False


}

ifres9230 = ifres9231


} else {
ifres9230 = False


}

var ifres9229 Obj

if True == ifres9230 {
ifres9229 = True


} else {
ifres9229 = False


}

ifres9228 = ifres9229


} else {
ifres9228 = False


}

var ifres9227 Obj

if True == ifres9228 {
ifres9227 = True


} else {
ifres9227 = False


}

ifres9226 = ifres9227


} else {
ifres9226 = False


}

var ifres9225 Obj

if True == ifres9226 {
ifres9225 = True


} else {
ifres9225 = False


}

ifres9224 = ifres9225


} else {
ifres9224 = False


}

if True == ifres9224 {
tmp9211 := PrimTail(V5825)

tmp9212 := PrimHead(tmp9211)

tmp9213 := PrimTail(V5825)

tmp9214 := PrimTail(tmp9213)

tmp9215 := PrimHead(tmp9214)

tmp9216 := PrimTail(V5825)

tmp9217 := PrimTail(tmp9216)

tmp9218 := PrimTail(tmp9217)

tmp9219 := PrimCons(symlet, tmp9218)

tmp9220 := PrimCons(tmp9219, Nil)

tmp9221 := PrimCons(tmp9215, tmp9220)

tmp9222 := PrimCons(tmp9212, tmp9221)

__e.Return(PrimCons(symlet, tmp9222))
return


} else {
__e.Return(V5825)
return
}


}, 1)

tmp9251 := Call(__e, ns2_1set, symshen_4process_1let, tmp9210)


_ = tmp9251

tmp9252 := MakeNative(func(__e *ControlFlow) {
V5826 := __e.Get(1)
_ = V5826
tmp9253 := MakeNative(func(__e *ControlFlow) {
GoTo5828 := __e.Get(1)
_ = GoTo5828
tmp9288 := PrimIsPair(V5826)

if True == tmp9288 {
tmp9254 := MakeNative(func(__e *ControlFlow) {
Select5835 := __e.Get(1)
_ = Select5835
tmp9284 := PrimHead(V5826)

tmp9285 := PrimEqual(sym_8s, tmp9284)

if True == tmp9285 {
tmp9282 := PrimIsPair(Select5835)

if True == tmp9282 {
tmp9255 := MakeNative(func(__e *ControlFlow) {
Select5833 := __e.Get(1)
_ = Select5833
tmp9256 := MakeNative(func(__e *ControlFlow) {
Select5834 := __e.Get(1)
_ = Select5834
tmp9278 := PrimIsPair(Select5834)

if True == tmp9278 {
tmp9257 := MakeNative(func(__e *ControlFlow) {
Select5832 := __e.Get(1)
_ = Select5832
tmp9275 := PrimIsPair(Select5832)

if True == tmp9275 {
tmp9258 := PrimCons(sym_8s, Select5834)

tmp9259 := Call(__e, PrimFunc(symshen_4process_1_8s), tmp9258)


tmp9260 := PrimCons(tmp9259, Nil)

tmp9261 := PrimCons(Select5833, tmp9260)

__e.Return(PrimCons(sym_8s, tmp9261))
return


} else {
tmp9273 := PrimEqual(Nil, Select5832)

var ifres9270 Obj

if True == tmp9273 {
tmp9272 := PrimIsString(Select5833)

var ifres9271 Obj

if True == tmp9272 {
ifres9271 = True


} else {
ifres9271 = False


}

ifres9270 = ifres9271


} else {
ifres9270 = False


}

if True == ifres9270 {
tmp9262 := MakeNative(func(__e *ControlFlow) {
W5827 := __e.Get(1)
_ = W5827
tmp9266 := Call(__e, PrimFunc(symlength), W5827)


tmp9267 := PrimGreatThan(tmp9266, MakeNumber(1))

if True == tmp9267 {
tmp9263 := Call(__e, PrimFunc(symappend), W5827, Select5834)


tmp9264 := PrimCons(sym_8s, tmp9263)

__e.TailApply(PrimFunc(symshen_4process_1_8s), tmp9264)
return


} else {
__e.Return(V5826)
return
}


}, 1)

tmp9268 := Call(__e, PrimFunc(symexplode), Select5833)


__e.TailApply(tmp9262, tmp9268)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5828)
return
}


}


}, 1)

tmp9276 := PrimTail(Select5834)

__e.TailApply(tmp9257, tmp9276)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5828)
return
}


}, 1)

tmp9279 := PrimTail(Select5835)

__e.TailApply(tmp9256, tmp9279)
return


}, 1)

tmp9280 := PrimHead(Select5835)

__e.TailApply(tmp9255, tmp9280)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5828)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5828)
return
}


}, 1)

tmp9286 := PrimTail(V5826)

__e.TailApply(tmp9254, tmp9286)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5828)
return
}


}, 1)

tmp9289 := MakeNative(func(__e *ControlFlow) {
__e.Return(V5826)
return
}, 0)

__e.TailApply(tmp9253, tmp9289)
return


}, 1)

tmp9290 := Call(__e, ns2_1set, symshen_4process_1_8s, tmp9252)


_ = tmp9290

tmp9291 := MakeNative(func(__e *ControlFlow) {
V5836 := __e.Get(1)
_ = V5836
V5837 := __e.Get(2)
_ = V5837
tmp9292 := MakeNative(func(__e *ControlFlow) {
W5838 := __e.Get(1)
_ = W5838
tmp9293 := MakeNative(func(__e *ControlFlow) {
W5839 := __e.Get(1)
_ = W5839
__e.Return(W5838)
return
}, 1)

tmp9294 := MakeNative(func(__e *ControlFlow) {
Z5840 := __e.Get(1)
_ = Z5840
__e.TailApply(PrimFunc(symshen_4_5datatype_6), Z5840)
return
}, 1)

tmp9295 := PrimCons(W5838, V5837)

tmp9296 := Call(__e, PrimFunc(symcompile), tmp9294, tmp9295)


__e.TailApply(tmp9293, tmp9296)
return


}, 1)

tmp9297 := Call(__e, PrimFunc(symshen_4intern_1type), V5836)


__e.TailApply(tmp9292, tmp9297)
return


}, 2)

tmp9298 := Call(__e, ns2_1set, symshen_4process_1datatype, tmp9291)


_ = tmp9298

tmp9299 := MakeNative(func(__e *ControlFlow) {
V5841 := __e.Get(1)
_ = V5841
tmp9300 := PrimStr(V5841)

tmp9301 := PrimStringConcat(tmp9300, MakeString("#type"))

__e.Return(PrimIntern(tmp9301))
return


}, 1)

tmp9302 := Call(__e, ns2_1set, symshen_4intern_1type, tmp9299)


_ = tmp9302

tmp9303 := MakeNative(func(__e *ControlFlow) {
V5842 := __e.Get(1)
_ = V5842
tmp9304 := PrimValue(symshen_4_dsynonyms_d)

tmp9305 := Call(__e, PrimFunc(symappend), V5842, tmp9304)


tmp9306 := PrimSet(symshen_4_dsynonyms_d, tmp9305)

__e.TailApply(PrimFunc(symshen_4synonyms_1h), tmp9306)
return


}, 1)

tmp9307 := Call(__e, ns2_1set, symshen_4process_1synonyms, tmp9303)


_ = tmp9307

tmp9308 := MakeNative(func(__e *ControlFlow) {
V5843 := __e.Get(1)
_ = V5843
tmp9309 := MakeNative(func(__e *ControlFlow) {
W5844 := __e.Get(1)
_ = W5844
tmp9310 := MakeNative(func(__e *ControlFlow) {
W5846 := __e.Get(1)
_ = W5846
__e.Return(symsynonyms)
return
}, 1)

tmp9311 := Call(__e, PrimFunc(symshen_4compile_1synonyms), W5844)


tmp9312 := PrimCons(symshen_4demod, tmp9311)

tmp9313 := PrimCons(symdefine, tmp9312)

tmp9314 := Call(__e, PrimFunc(symeval), tmp9313)


__e.TailApply(tmp9310, tmp9314)
return


}, 1)

tmp9315 := MakeNative(func(__e *ControlFlow) {
Z5845 := __e.Get(1)
_ = Z5845
__e.TailApply(PrimFunc(symshen_4curry_1type), Z5845)
return
}, 1)

tmp9316 := Call(__e, PrimFunc(symmap), tmp9315, V5843)


__e.TailApply(tmp9309, tmp9316)
return


}, 1)

tmp9317 := Call(__e, ns2_1set, symshen_4synonyms_1h, tmp9308)


_ = tmp9317

tmp9318 := MakeNative(func(__e *ControlFlow) {
V5849 := __e.Get(1)
_ = V5849
tmp9340 := PrimEqual(Nil, V5849)

if True == tmp9340 {
tmp9319 := MakeNative(func(__e *ControlFlow) {
W5850 := __e.Get(1)
_ = W5850
tmp9320 := PrimCons(W5850, Nil)

tmp9321 := PrimCons(sym_1_6, tmp9320)

__e.Return(PrimCons(W5850, tmp9321))
return


}, 1)

tmp9322 := Call(__e, PrimFunc(symgensym), symX)


__e.TailApply(tmp9319, tmp9322)
return


} else {
tmp9338 := PrimIsPair(V5849)

var ifres9334 Obj

if True == tmp9338 {
tmp9336 := PrimTail(V5849)

tmp9337 := PrimIsPair(tmp9336)

var ifres9335 Obj

if True == tmp9337 {
ifres9335 = True


} else {
ifres9335 = False


}

ifres9334 = ifres9335


} else {
ifres9334 = False


}

if True == ifres9334 {
tmp9323 := PrimHead(V5849)

tmp9324 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9323)


tmp9325 := PrimTail(V5849)

tmp9326 := PrimHead(tmp9325)

tmp9327 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9326)


tmp9328 := PrimTail(V5849)

tmp9329 := PrimTail(tmp9328)

tmp9330 := Call(__e, PrimFunc(symshen_4compile_1synonyms), tmp9329)


tmp9331 := PrimCons(tmp9327, tmp9330)

tmp9332 := PrimCons(sym_1_6, tmp9331)

__e.Return(PrimCons(tmp9324, tmp9332))
return


} else {
__e.Return(PrimSimpleError(MakeString("synonyms requires an even number of arguments\n")))
return
}


}


}, 1)

tmp9341 := Call(__e, ns2_1set, symshen_4compile_1synonyms, tmp9318)


_ = tmp9341

tmp9342 := MakeNative(func(__e *ControlFlow) {
V5851 := __e.Get(1)
_ = V5851
tmp9343 := MakeNative(func(__e *ControlFlow) {
GoTo5852 := __e.Get(1)
_ = GoTo5852
tmp9371 := PrimIsPair(V5851)

if True == tmp9371 {
tmp9344 := MakeNative(func(__e *ControlFlow) {
Select5859 := __e.Get(1)
_ = Select5859
tmp9367 := PrimHead(V5851)

tmp9368 := PrimEqual(sym_c_4, tmp9367)

if True == tmp9368 {
tmp9365 := PrimIsPair(Select5859)

if True == tmp9365 {
tmp9345 := MakeNative(func(__e *ControlFlow) {
Select5857 := __e.Get(1)
_ = Select5857
tmp9346 := MakeNative(func(__e *ControlFlow) {
Select5858 := __e.Get(1)
_ = Select5858
tmp9361 := PrimIsPair(Select5858)

if True == tmp9361 {
tmp9347 := MakeNative(func(__e *ControlFlow) {
Select5856 := __e.Get(1)
_ = Select5856
tmp9358 := PrimIsPair(Select5856)

if True == tmp9358 {
tmp9348 := PrimCons(sym_c_4, Select5858)

tmp9349 := Call(__e, PrimFunc(symshen_4process_1lambda), tmp9348)


tmp9350 := PrimCons(tmp9349, Nil)

tmp9351 := PrimCons(Select5857, tmp9350)

__e.Return(PrimCons(symlambda, tmp9351))
return


} else {
tmp9356 := PrimEqual(Nil, Select5856)

if True == tmp9356 {
tmp9354 := PrimIsVariable(Select5857)

if True == tmp9354 {
__e.Return(PrimCons(symlambda, Select5859))
return
} else {
tmp9352 := Call(__e, PrimFunc(symshen_4app), Select5857, MakeString(" is not a variable\n"), symshen_4s)


__e.Return(PrimSimpleError(tmp9352))
return


}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5852)
return
}


}


}, 1)

tmp9359 := PrimTail(Select5858)

__e.TailApply(tmp9347, tmp9359)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5852)
return
}


}, 1)

tmp9362 := PrimTail(Select5859)

__e.TailApply(tmp9346, tmp9362)
return


}, 1)

tmp9363 := PrimHead(Select5859)

__e.TailApply(tmp9345, tmp9363)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5852)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5852)
return
}


}, 1)

tmp9369 := PrimTail(V5851)

__e.TailApply(tmp9344, tmp9369)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5852)
return
}


}, 1)

tmp9372 := MakeNative(func(__e *ControlFlow) {
__e.Return(V5851)
return
}, 0)

__e.TailApply(tmp9343, tmp9372)
return


}, 1)

tmp9373 := Call(__e, ns2_1set, symshen_4process_1lambda, tmp9342)


_ = tmp9373

tmp9374 := MakeNative(func(__e *ControlFlow) {
V5862 := __e.Get(1)
_ = V5862
tmp9375 := MakeNative(func(__e *ControlFlow) {
GoTo5863 := __e.Get(1)
_ = GoTo5863
tmp9415 := PrimIsPair(V5862)

if True == tmp9415 {
tmp9376 := MakeNative(func(__e *ControlFlow) {
Select5871 := __e.Get(1)
_ = Select5871
tmp9411 := PrimHead(V5862)

tmp9412 := PrimEqual(symcases, tmp9411)

if True == tmp9412 {
tmp9409 := PrimIsPair(Select5871)

if True == tmp9409 {
tmp9377 := MakeNative(func(__e *ControlFlow) {
Select5869 := __e.Get(1)
_ = Select5869
tmp9378 := MakeNative(func(__e *ControlFlow) {
Select5870 := __e.Get(1)
_ = Select5870
tmp9405 := PrimEqual(True, Select5869)

var ifres9402 Obj

if True == tmp9405 {
tmp9404 := PrimIsPair(Select5870)

var ifres9403 Obj

if True == tmp9404 {
ifres9403 = True


} else {
ifres9403 = False


}

ifres9402 = ifres9403


} else {
ifres9402 = False


}

if True == ifres9402 {
__e.Return(PrimHead(Select5870))
return
} else {
tmp9379 := MakeNative(func(__e *ControlFlow) {
GoTo5866 := __e.Get(1)
_ = GoTo5866
tmp9397 := PrimIsPair(Select5870)

if True == tmp9397 {
tmp9380 := MakeNative(func(__e *ControlFlow) {
Select5867 := __e.Get(1)
_ = Select5867
tmp9381 := MakeNative(func(__e *ControlFlow) {
Select5868 := __e.Get(1)
_ = Select5868
tmp9393 := PrimEqual(Nil, Select5868)

if True == tmp9393 {
tmp9382 := PrimCons(MakeString("error: cases exhausted"), Nil)

tmp9383 := PrimCons(symsimple_1error, tmp9382)

tmp9384 := PrimCons(tmp9383, Nil)

tmp9385 := PrimCons(Select5867, tmp9384)

tmp9386 := PrimCons(Select5869, tmp9385)

__e.Return(PrimCons(symif, tmp9386))
return


} else {
tmp9387 := PrimCons(symcases, Select5868)

tmp9388 := Call(__e, PrimFunc(symshen_4process_1cases), tmp9387)


tmp9389 := PrimCons(tmp9388, Nil)

tmp9390 := PrimCons(Select5867, tmp9389)

tmp9391 := PrimCons(Select5869, tmp9390)

__e.Return(PrimCons(symif, tmp9391))
return


}


}, 1)

tmp9394 := PrimTail(Select5870)

__e.TailApply(tmp9381, tmp9394)
return


}, 1)

tmp9395 := PrimHead(Select5870)

__e.TailApply(tmp9380, tmp9395)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5866)
return
}


}, 1)

tmp9398 := MakeNative(func(__e *ControlFlow) {
tmp9400 := PrimEqual(Nil, Select5870)

if True == tmp9400 {
__e.Return(PrimSimpleError(MakeString("error: odd number of case elements\n")))
return
} else {
__e.TailApply(PrimFunc(symthaw), GoTo5863)
return
}


}, 0)

__e.TailApply(tmp9379, tmp9398)
return


}


}, 1)

tmp9406 := PrimTail(Select5871)

__e.TailApply(tmp9378, tmp9406)
return


}, 1)

tmp9407 := PrimHead(Select5871)

__e.TailApply(tmp9377, tmp9407)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5863)
return
}


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5863)
return
}


}, 1)

tmp9413 := PrimTail(V5862)

__e.TailApply(tmp9376, tmp9413)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5863)
return
}


}, 1)

tmp9416 := MakeNative(func(__e *ControlFlow) {
__e.Return(V5862)
return
}, 0)

__e.TailApply(tmp9375, tmp9416)
return


}, 1)

tmp9417 := Call(__e, ns2_1set, symshen_4process_1cases, tmp9374)


_ = tmp9417

tmp9418 := MakeNative(func(__e *ControlFlow) {
V5872 := __e.Get(1)
_ = V5872
tmp9419 := PrimCons(symrun, Nil)

tmp9420 := PrimCons(symget_1time, tmp9419)

tmp9421 := PrimCons(symrun, Nil)

tmp9422 := PrimCons(symget_1time, tmp9421)

tmp9423 := PrimCons(symStart, Nil)

tmp9424 := PrimCons(symFinish, tmp9423)

tmp9425 := PrimCons(sym_1, tmp9424)

tmp9426 := PrimCons(symTime, Nil)

tmp9427 := PrimCons(symstr, tmp9426)

tmp9428 := PrimCons(MakeString(" secs\n"), Nil)

tmp9429 := PrimCons(tmp9427, tmp9428)

tmp9430 := PrimCons(symcn, tmp9429)

tmp9431 := PrimCons(tmp9430, Nil)

tmp9432 := PrimCons(MakeString("\nrun time: "), tmp9431)

tmp9433 := PrimCons(symcn, tmp9432)

tmp9434 := PrimCons(symstoutput, Nil)

tmp9435 := PrimCons(tmp9434, Nil)

tmp9436 := PrimCons(tmp9433, tmp9435)

tmp9437 := PrimCons(sympr, tmp9436)

tmp9438 := PrimCons(symResult, Nil)

tmp9439 := PrimCons(tmp9437, tmp9438)

tmp9440 := PrimCons(symMessage, tmp9439)

tmp9441 := PrimCons(tmp9425, tmp9440)

tmp9442 := PrimCons(symTime, tmp9441)

tmp9443 := PrimCons(tmp9422, tmp9442)

tmp9444 := PrimCons(symFinish, tmp9443)

tmp9445 := PrimCons(V5872, tmp9444)

tmp9446 := PrimCons(symResult, tmp9445)

tmp9447 := PrimCons(tmp9420, tmp9446)

tmp9448 := PrimCons(symStart, tmp9447)

__e.Return(PrimCons(symlet, tmp9448))
return


}, 1)

tmp9449 := Call(__e, ns2_1set, symshen_4process_1time, tmp9418)


_ = tmp9449

tmp9450 := MakeNative(func(__e *ControlFlow) {
V5873 := __e.Get(1)
_ = V5873
tmp9476 := PrimIsPair(V5873)

var ifres9461 Obj

if True == tmp9476 {
tmp9474 := PrimTail(V5873)

tmp9475 := PrimIsPair(tmp9474)

var ifres9463 Obj

if True == tmp9475 {
tmp9471 := PrimTail(V5873)

tmp9472 := PrimTail(tmp9471)

tmp9473 := PrimIsPair(tmp9472)

var ifres9465 Obj

if True == tmp9473 {
tmp9467 := PrimTail(V5873)

tmp9468 := PrimTail(tmp9467)

tmp9469 := PrimTail(tmp9468)

tmp9470 := PrimIsPair(tmp9469)

var ifres9466 Obj

if True == tmp9470 {
ifres9466 = True


} else {
ifres9466 = False


}

ifres9465 = ifres9466


} else {
ifres9465 = False


}

var ifres9464 Obj

if True == ifres9465 {
ifres9464 = True


} else {
ifres9464 = False


}

ifres9463 = ifres9464


} else {
ifres9463 = False


}

var ifres9462 Obj

if True == ifres9463 {
ifres9462 = True


} else {
ifres9462 = False


}

ifres9461 = ifres9462


} else {
ifres9461 = False


}

if True == ifres9461 {
tmp9451 := PrimHead(V5873)

tmp9452 := PrimTail(V5873)

tmp9453 := PrimHead(tmp9452)

tmp9454 := PrimHead(V5873)

tmp9455 := PrimTail(V5873)

tmp9456 := PrimTail(tmp9455)

tmp9457 := PrimCons(tmp9454, tmp9456)

tmp9458 := PrimCons(tmp9457, Nil)

tmp9459 := PrimCons(tmp9453, tmp9458)

__e.Return(PrimCons(tmp9451, tmp9459))
return


} else {
__e.Return(V5873)
return
}


}, 1)

tmp9477 := Call(__e, ns2_1set, symshen_4process_1assoc, tmp9450)


_ = tmp9477

tmp9478 := MakeNative(func(__e *ControlFlow) {
V5874 := __e.Get(1)
_ = V5874
tmp9479 := PrimStr(V5874)

tmp9480 := Call(__e, PrimFunc(symshen_4mu_1h), tmp9479)


__e.Return(PrimIntern(tmp9480))
return


}, 1)

tmp9481 := Call(__e, ns2_1set, symshen_4make_1uppercase, tmp9478)


_ = tmp9481

tmp9482 := MakeNative(func(__e *ControlFlow) {
V5875 := __e.Get(1)
_ = V5875
tmp9501 := PrimEqual(MakeString(""), V5875)

if True == tmp9501 {
__e.Return(MakeString(""))
return
} else {
tmp9499 := Call(__e, PrimFunc(symshen_4_7string_2), V5875)


if True == tmp9499 {
tmp9483 := MakeNative(func(__e *ControlFlow) {
W5876 := __e.Get(1)
_ = W5876
tmp9484 := MakeNative(func(__e *ControlFlow) {
W5877 := __e.Get(1)
_ = W5877
tmp9485 := MakeNative(func(__e *ControlFlow) {
W5878 := __e.Get(1)
_ = W5878
tmp9486 := PrimTailString(V5875)

tmp9487 := Call(__e, PrimFunc(symshen_4mu_1h), tmp9486)


__e.TailApply(PrimFunc(sym_8s), W5878, tmp9487)
return


}, 1)

tmp9494 := PrimGreatEqual(W5876, MakeNumber(97))

var ifres9491 Obj

if True == tmp9494 {
tmp9493 := PrimLessEqual(W5876, MakeNumber(122))

var ifres9492 Obj

if True == tmp9493 {
ifres9492 = True


} else {
ifres9492 = False


}

ifres9491 = ifres9492


} else {
ifres9491 = False


}

var ifres9488 Obj

if True == ifres9491 {
tmp9489 := PrimNumberToString(W5877)

ifres9488 = tmp9489


} else {
tmp9490 := Call(__e, PrimFunc(symhdstr), V5875)


ifres9488 = tmp9490


}

__e.TailApply(tmp9485, ifres9488)
return


}, 1)

tmp9495 := PrimNumberSubtract(W5876, MakeNumber(32))

__e.TailApply(tmp9484, tmp9495)
return


}, 1)

tmp9496 := Call(__e, PrimFunc(symhdstr), V5875)


tmp9497 := PrimStringToNumber(tmp9496)

__e.TailApply(tmp9483, tmp9497)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.mu-h")))
return
}


}


}, 1)

tmp9502 := Call(__e, ns2_1set, symshen_4mu_1h, tmp9482)


_ = tmp9502

tmp9503 := MakeNative(func(__e *ControlFlow) {
V5879 := __e.Get(1)
_ = V5879
V5880 := __e.Get(2)
_ = V5880
tmp9504 := PrimValue(sym_dmacros_d)

tmp9505 := Call(__e, PrimFunc(symshen_4update_1assoc), V5879, V5880, tmp9504)


__e.Return(PrimSet(sym_dmacros_d, tmp9505))
return


}, 2)

tmp9506 := Call(__e, ns2_1set, symshen_4record_1macro, tmp9503)


_ = tmp9506

tmp9507 := MakeNative(func(__e *ControlFlow) {
V5890 := __e.Get(1)
_ = V5890
V5891 := __e.Get(2)
_ = V5891
V5892 := __e.Get(3)
_ = V5892
tmp9527 := PrimEqual(Nil, V5892)

if True == tmp9527 {
tmp9508 := PrimCons(V5890, V5891)

__e.Return(PrimCons(tmp9508, Nil))
return


} else {
tmp9509 := MakeNative(func(__e *ControlFlow) {
GoTo5893 := __e.Get(1)
_ = GoTo5893
tmp9524 := PrimIsPair(V5892)

if True == tmp9524 {
tmp9510 := MakeNative(func(__e *ControlFlow) {
Select5894 := __e.Get(1)
_ = Select5894
tmp9511 := MakeNative(func(__e *ControlFlow) {
Select5895 := __e.Get(1)
_ = Select5895
tmp9520 := PrimIsPair(Select5894)

var ifres9516 Obj

if True == tmp9520 {
tmp9518 := PrimHead(Select5894)

tmp9519 := PrimEqual(V5890, tmp9518)

var ifres9517 Obj

if True == tmp9519 {
ifres9517 = True


} else {
ifres9517 = False


}

ifres9516 = ifres9517


} else {
ifres9516 = False


}

if True == ifres9516 {
tmp9512 := PrimHead(Select5894)

tmp9513 := PrimCons(tmp9512, V5891)

__e.Return(PrimCons(tmp9513, Select5895))
return


} else {
tmp9514 := Call(__e, PrimFunc(symshen_4update_1assoc), V5890, V5891, Select5895)


__e.Return(PrimCons(Select5894, tmp9514))
return


}


}, 1)

tmp9521 := PrimTail(V5892)

__e.TailApply(tmp9511, tmp9521)
return


}, 1)

tmp9522 := PrimHead(V5892)

__e.TailApply(tmp9510, tmp9522)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5893)
return
}


}, 1)

tmp9525 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.update-assoc")))
return
}, 0)

__e.TailApply(tmp9509, tmp9525)
return


}


}, 3)

tmp9528 := Call(__e, ns2_1set, symshen_4update_1assoc, tmp9507)


_ = tmp9528

tmp9529 := MakeNative(func(__e *ControlFlow) {
tmp9537 := Call(__e, PrimFunc(symstinput))


tmp9538 := Call(__e, PrimFunc(symshen_4char_1stinput_2), tmp9537)


if True == tmp9538 {
tmp9530 := PrimCons(symstinput, Nil)

tmp9531 := PrimCons(tmp9530, Nil)

tmp9532 := PrimCons(symshen_4read_1unit_1string, tmp9531)

tmp9533 := PrimCons(tmp9532, Nil)

__e.Return(PrimCons(symstring_1_6n, tmp9533))
return


} else {
tmp9534 := PrimCons(symstinput, Nil)

tmp9535 := PrimCons(tmp9534, Nil)

__e.Return(PrimCons(symread_1byte, tmp9535))
return


}


}, 0)

tmp9539 := Call(__e, ns2_1set, symshen_4process_1read_1byte, tmp9529)


_ = tmp9539

tmp9540 := MakeNative(func(__e *ControlFlow) {
V5896 := __e.Get(1)
_ = V5896
tmp9541 := MakeNative(func(__e *ControlFlow) {
W5897 := __e.Get(1)
_ = W5897
tmp9542 := MakeNative(func(__e *ControlFlow) {
W5898 := __e.Get(1)
_ = W5898
tmp9543 := MakeNative(func(__e *ControlFlow) {
W5899 := __e.Get(1)
_ = W5899
tmp9544 := MakeNative(func(__e *ControlFlow) {
W5900 := __e.Get(1)
_ = W5900
tmp9545 := MakeNative(func(__e *ControlFlow) {
W5901 := __e.Get(1)
_ = W5901
tmp9546 := MakeNative(func(__e *ControlFlow) {
W5903 := __e.Get(1)
_ = W5903
tmp9547 := MakeNative(func(__e *ControlFlow) {
W5904 := __e.Get(1)
_ = W5904
tmp9548 := MakeNative(func(__e *ControlFlow) {
W5905 := __e.Get(1)
_ = W5905
tmp9549 := MakeNative(func(__e *ControlFlow) {
W5906 := __e.Get(1)
_ = W5906
tmp9550 := MakeNative(func(__e *ControlFlow) {
W5907 := __e.Get(1)
_ = W5907
tmp9551 := MakeNative(func(__e *ControlFlow) {
W5908 := __e.Get(1)
_ = W5908
tmp9552 := PrimCons(W5900, Nil)

tmp9553 := PrimCons(W5899, tmp9552)

tmp9554 := PrimCons(W5898, tmp9553)

tmp9555 := PrimCons(W5897, tmp9554)

__e.Return(PrimCons(W5908, tmp9555))
return


}, 1)

tmp9556 := Call(__e, PrimFunc(symshen_4continue), W5903, W5901, W5904, W5905, W5906, W5907)


tmp9557 := PrimCons(tmp9556, Nil)

tmp9558 := PrimCons(W5907, tmp9557)

tmp9559 := PrimCons(symlambda, tmp9558)

tmp9560 := PrimCons(tmp9559, Nil)

tmp9561 := PrimCons(W5906, tmp9560)

tmp9562 := PrimCons(symlambda, tmp9561)

tmp9563 := PrimCons(tmp9562, Nil)

tmp9564 := PrimCons(W5905, tmp9563)

tmp9565 := PrimCons(symlambda, tmp9564)

tmp9566 := PrimCons(tmp9565, Nil)

tmp9567 := PrimCons(W5904, tmp9566)

tmp9568 := PrimCons(symlambda, tmp9567)

__e.TailApply(tmp9551, tmp9568)
return


}, 1)

tmp9569 := Call(__e, PrimFunc(symgensym), symC)


__e.TailApply(tmp9550, tmp9569)
return


}, 1)

tmp9570 := Call(__e, PrimFunc(symgensym), symK)


__e.TailApply(tmp9549, tmp9570)
return


}, 1)

tmp9571 := Call(__e, PrimFunc(symgensym), symL)


__e.TailApply(tmp9548, tmp9571)
return


}, 1)

tmp9572 := Call(__e, PrimFunc(symgensym), symV)


__e.TailApply(tmp9547, tmp9572)
return


}, 1)

tmp9573 := Call(__e, PrimFunc(symshen_4received), V5896)


__e.TailApply(tmp9546, tmp9573)
return


}, 1)

tmp9574 := MakeNative(func(__e *ControlFlow) {
Z5902 := __e.Get(1)
_ = Z5902
__e.TailApply(PrimFunc(symshen_4_5body_6), Z5902)
return
}, 1)

tmp9575 := Call(__e, PrimFunc(symcompile), tmp9574, V5896)


__e.TailApply(tmp9545, tmp9575)
return


}, 1)

tmp9576 := PrimCons(True, Nil)

tmp9577 := PrimCons(symfreeze, tmp9576)

__e.TailApply(tmp9544, tmp9577)
return


}, 1)

__e.TailApply(tmp9543, MakeNumber(0))
return


}, 1)

tmp9578 := PrimCons(MakeNumber(0), Nil)

tmp9579 := PrimCons(symvector, tmp9578)

tmp9580 := PrimCons(tmp9579, Nil)

tmp9581 := PrimCons(MakeNumber(0), tmp9580)

tmp9582 := PrimCons(True, tmp9581)

tmp9583 := PrimCons(sym_8v, tmp9582)

__e.TailApply(tmp9542, tmp9583)
return


}, 1)

tmp9584 := PrimCons(symshen_4prolog_1vector, Nil)

__e.TailApply(tmp9541, tmp9584)
return


}, 1)

tmp9585 := Call(__e, ns2_1set, symshen_4call_1prolog, tmp9540)


_ = tmp9585

tmp9586 := MakeNative(func(__e *ControlFlow) {
V5911 := __e.Get(1)
_ = V5911
tmp9587 := MakeNative(func(__e *ControlFlow) {
GoTo5912 := __e.Get(1)
_ = GoTo5912
tmp9604 := PrimIsPair(V5911)

if True == tmp9604 {
tmp9588 := MakeNative(func(__e *ControlFlow) {
Select5913 := __e.Get(1)
_ = Select5913
tmp9589 := MakeNative(func(__e *ControlFlow) {
Select5914 := __e.Get(1)
_ = Select5914
tmp9600 := PrimEqual(symreceive, Select5913)

var ifres9593 Obj

if True == tmp9600 {
tmp9599 := PrimIsPair(Select5914)

var ifres9595 Obj

if True == tmp9599 {
tmp9597 := PrimTail(Select5914)

tmp9598 := PrimEqual(Nil, tmp9597)

var ifres9596 Obj

if True == tmp9598 {
ifres9596 = True


} else {
ifres9596 = False


}

ifres9595 = ifres9596


} else {
ifres9595 = False


}

var ifres9594 Obj

if True == ifres9595 {
ifres9594 = True


} else {
ifres9594 = False


}

ifres9593 = ifres9594


} else {
ifres9593 = False


}

if True == ifres9593 {
__e.Return(Select5914)
return
} else {
tmp9590 := Call(__e, PrimFunc(symshen_4received), Select5913)


tmp9591 := Call(__e, PrimFunc(symshen_4received), Select5914)


__e.TailApply(PrimFunc(symunion), tmp9590, tmp9591)
return


}


}, 1)

tmp9601 := PrimTail(V5911)

__e.TailApply(tmp9589, tmp9601)
return


}, 1)

tmp9602 := PrimHead(V5911)

__e.TailApply(tmp9588, tmp9602)
return


} else {
__e.TailApply(PrimFunc(symthaw), GoTo5912)
return
}


}, 1)

tmp9605 := MakeNative(func(__e *ControlFlow) {
__e.Return(Nil)
return
}, 0)

__e.TailApply(tmp9587, tmp9605)
return


}, 1)

tmp9606 := Call(__e, ns2_1set, symshen_4received, tmp9586)


_ = tmp9606

tmp9607 := MakeNative(func(__e *ControlFlow) {
tmp9608 := MakeNative(func(__e *ControlFlow) {
W5915 := __e.Get(1)
_ = W5915
tmp9609 := MakeNative(func(__e *ControlFlow) {
W5916 := __e.Get(1)
_ = W5916
tmp9610 := MakeNative(func(__e *ControlFlow) {
W5917 := __e.Get(1)
_ = W5917
__e.Return(W5917)
return
}, 1)

tmp9611 := PrimVectorSet(W5915, MakeNumber(1), MakeNumber(2))

__e.TailApply(tmp9610, tmp9611)
return


}, 1)

tmp9612 := PrimVectorSet(W5915, MakeNumber(0), symshen_4print_1prolog_1vector)

__e.TailApply(tmp9609, tmp9612)
return


}, 1)

tmp9613 := PrimValue(symshen_4_dprolog_1memory_d)

tmp9614 := PrimAbsvector(tmp9613)

__e.TailApply(tmp9608, tmp9614)
return


}, 0)

tmp9615 := Call(__e, ns2_1set, symshen_4prolog_1vector, tmp9607)


_ = tmp9615

tmp9616 := MakeNative(func(__e *ControlFlow) {
V5918 := __e.Get(1)
_ = V5918
__e.Return(V5918)
return
}, 1)

tmp9617 := Call(__e, ns2_1set, symreceive, tmp9616)


_ = tmp9617

tmp9618 := MakeNative(func(__e *ControlFlow) {
V5919 := __e.Get(1)
_ = V5919
tmp9626 := PrimIsPair(V5919)

if True == tmp9626 {
tmp9619 := PrimHead(V5919)

tmp9620 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9619)


tmp9621 := PrimTail(V5919)

tmp9622 := Call(__e, PrimFunc(symshen_4rcons__form), tmp9621)


tmp9623 := PrimCons(tmp9622, Nil)

tmp9624 := PrimCons(tmp9620, tmp9623)

__e.Return(PrimCons(symcons, tmp9624))
return


} else {
__e.Return(V5919)
return
}


}, 1)

tmp9627 := Call(__e, ns2_1set, symshen_4rcons__form, tmp9618)


_ = tmp9627

tmp9628 := MakeNative(func(__e *ControlFlow) {
V5920 := __e.Get(1)
_ = V5920
tmp9635 := PrimIsPair(V5920)

if True == tmp9635 {
tmp9629 := PrimHead(V5920)

tmp9630 := PrimTail(V5920)

tmp9631 := Call(__e, PrimFunc(symshen_4tuple_1up), tmp9630)


tmp9632 := PrimCons(tmp9631, Nil)

tmp9633 := PrimCons(tmp9629, tmp9632)

__e.Return(PrimCons(sym_8p, tmp9633))
return


} else {
__e.Return(V5920)
return
}


}, 1)

tmp9636 := Call(__e, ns2_1set, symshen_4tuple_1up, tmp9628)


_ = tmp9636

tmp9637 := MakeNative(func(__e *ControlFlow) {
V5921 := __e.Get(1)
_ = V5921
tmp9638 := PrimValue(sym_dmacros_d)

tmp9639 := Call(__e, PrimFunc(symassoc), V5921, tmp9638)


tmp9640 := PrimValue(sym_dmacros_d)

tmp9641 := Call(__e, PrimFunc(symremove), tmp9639, tmp9640)


tmp9642 := PrimSet(sym_dmacros_d, tmp9641)

_ = tmp9642

__e.Return(V5921)
return


}, 1)

__e.TailApply(ns2_1set, symundefmacro, tmp9637)
return




}, 0)

