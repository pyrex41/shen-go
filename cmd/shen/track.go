package main

import . "github.com/tiancaiamao/shen-go/kl"

var TrackMain = MakeNative(func(__e *ControlFlow) {
tmp11917 := MakeNative(func(__e *ControlFlow) {
V5790 := __e.Get(1)
_ = V5790
tmp11918 := Call(__e, PrimFunc(symshen_4app), V5790, MakeString(";\n"), symshen_4a)


tmp11919 := PrimStringConcat(MakeString("partial function "), tmp11918)

tmp11920 := Call(__e, PrimFunc(symstoutput))


tmp11921 := Call(__e, PrimFunc(sympr), tmp11919, tmp11920)


_ = tmp11921

tmp11930 := Call(__e, PrimFunc(symshen_4tracked_2), V5790)


tmp11931 := PrimNot(tmp11930)

var ifres11925 Obj

if True == tmp11931 {
tmp11927 := Call(__e, PrimFunc(symshen_4app), V5790, MakeString("? "), symshen_4a)


tmp11928 := PrimStringConcat(MakeString("track "), tmp11927)

tmp11929 := Call(__e, PrimFunc(symy_1or_1n_2), tmp11928)


var ifres11926 Obj

if True == tmp11929 {
ifres11926 = True


} else {
ifres11926 = False


}

ifres11925 = ifres11926


} else {
ifres11925 = False


}

var ifres11922 Obj

if True == ifres11925 {
tmp11923 := Call(__e, PrimFunc(symps), V5790)


tmp11924 := Call(__e, PrimFunc(symshen_4track_1function), tmp11923)


ifres11922 = tmp11924


} else {
ifres11922 = symshen_4ok


}

_ = ifres11922

__e.Return(PrimSimpleError(MakeString("aborted")))
return


}, 1)

tmp11932 := Call(__e, ns2_1set, symshen_4f_1error, tmp11917)


_ = tmp11932

tmp11933 := MakeNative(func(__e *ControlFlow) {
V5791 := __e.Get(1)
_ = V5791
tmp11934 := PrimValue(symshen_4_dtracking_d)

__e.TailApply(PrimFunc(symelement_2), V5791, tmp11934)
return


}, 1)

tmp11935 := Call(__e, ns2_1set, symshen_4tracked_2, tmp11933)


_ = tmp11935

tmp11936 := MakeNative(func(__e *ControlFlow) {
V5792 := __e.Get(1)
_ = V5792
tmp11937 := MakeNative(func(__e *ControlFlow) {
W5793 := __e.Get(1)
_ = W5793
__e.TailApply(PrimFunc(symshen_4track_1function), W5793)
return
}, 1)

tmp11938 := Call(__e, PrimFunc(symps), V5792)


__e.TailApply(tmp11937, tmp11938)
return


}, 1)

tmp11939 := Call(__e, ns2_1set, symtrack, tmp11936)


_ = tmp11939

tmp11940 := MakeNative(func(__e *ControlFlow) {
V5796 := __e.Get(1)
_ = V5796
tmp11997 := PrimIsPair(V5796)

var ifres11971 Obj

if True == tmp11997 {
tmp11995 := PrimHead(V5796)

tmp11996 := PrimEqual(symdefun, tmp11995)

var ifres11973 Obj

if True == tmp11996 {
tmp11993 := PrimTail(V5796)

tmp11994 := PrimIsPair(tmp11993)

var ifres11975 Obj

if True == tmp11994 {
tmp11990 := PrimTail(V5796)

tmp11991 := PrimTail(tmp11990)

tmp11992 := PrimIsPair(tmp11991)

var ifres11977 Obj

if True == tmp11992 {
tmp11986 := PrimTail(V5796)

tmp11987 := PrimTail(tmp11986)

tmp11988 := PrimTail(tmp11987)

tmp11989 := PrimIsPair(tmp11988)

var ifres11979 Obj

if True == tmp11989 {
tmp11981 := PrimTail(V5796)

tmp11982 := PrimTail(tmp11981)

tmp11983 := PrimTail(tmp11982)

tmp11984 := PrimTail(tmp11983)

tmp11985 := PrimEqual(Nil, tmp11984)

var ifres11980 Obj

if True == tmp11985 {
ifres11980 = True


} else {
ifres11980 = False


}

ifres11979 = ifres11980


} else {
ifres11979 = False


}

var ifres11978 Obj

if True == ifres11979 {
ifres11978 = True


} else {
ifres11978 = False


}

ifres11977 = ifres11978


} else {
ifres11977 = False


}

var ifres11976 Obj

if True == ifres11977 {
ifres11976 = True


} else {
ifres11976 = False


}

ifres11975 = ifres11976


} else {
ifres11975 = False


}

var ifres11974 Obj

if True == ifres11975 {
ifres11974 = True


} else {
ifres11974 = False


}

ifres11973 = ifres11974


} else {
ifres11973 = False


}

var ifres11972 Obj

if True == ifres11973 {
ifres11972 = True


} else {
ifres11972 = False


}

ifres11971 = ifres11972


} else {
ifres11971 = False


}

if True == ifres11971 {
tmp11941 := MakeNative(func(__e *ControlFlow) {
W5797 := __e.Get(1)
_ = W5797
tmp11942 := MakeNative(func(__e *ControlFlow) {
W5798 := __e.Get(1)
_ = W5798
tmp11943 := MakeNative(func(__e *ControlFlow) {
W5799 := __e.Get(1)
_ = W5799
tmp11944 := PrimTail(V5796)

__e.Return(PrimHead(tmp11944))
return


}, 1)

tmp11945 := PrimTail(V5796)

tmp11946 := PrimHead(tmp11945)

tmp11947 := PrimValue(symshen_4_dtracking_d)

tmp11948 := Call(__e, PrimFunc(symadjoin), tmp11946, tmp11947)


tmp11949 := PrimSet(symshen_4_dtracking_d, tmp11948)

__e.TailApply(tmp11943, tmp11949)
return


}, 1)

tmp11950 := Call(__e, PrimFunc(symeval_1kl), W5797)


__e.TailApply(tmp11942, tmp11950)
return


}, 1)

tmp11951 := PrimTail(V5796)

tmp11952 := PrimHead(tmp11951)

tmp11953 := PrimTail(V5796)

tmp11954 := PrimTail(tmp11953)

tmp11955 := PrimHead(tmp11954)

tmp11956 := PrimTail(V5796)

tmp11957 := PrimHead(tmp11956)

tmp11958 := PrimTail(V5796)

tmp11959 := PrimTail(tmp11958)

tmp11960 := PrimHead(tmp11959)

tmp11961 := PrimTail(V5796)

tmp11962 := PrimTail(tmp11961)

tmp11963 := PrimTail(tmp11962)

tmp11964 := PrimHead(tmp11963)

tmp11965 := Call(__e, PrimFunc(symshen_4insert_1tracking_1code), tmp11957, tmp11960, tmp11964)


tmp11966 := PrimCons(tmp11965, Nil)

tmp11967 := PrimCons(tmp11955, tmp11966)

tmp11968 := PrimCons(tmp11952, tmp11967)

tmp11969 := PrimCons(symdefun, tmp11968)

__e.TailApply(tmp11941, tmp11969)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.track-function")))
return
}


}, 1)

tmp11998 := Call(__e, ns2_1set, symshen_4track_1function, tmp11940)


_ = tmp11998

tmp11999 := MakeNative(func(__e *ControlFlow) {
V5800 := __e.Get(1)
_ = V5800
V5801 := __e.Get(2)
_ = V5801
V5802 := __e.Get(3)
_ = V5802
tmp12000 := PrimCons(symshen_4_dcall_d, Nil)

tmp12001 := PrimCons(symvalue, tmp12000)

tmp12002 := PrimCons(MakeNumber(1), Nil)

tmp12003 := PrimCons(tmp12001, tmp12002)

tmp12004 := PrimCons(sym_7, tmp12003)

tmp12005 := PrimCons(tmp12004, Nil)

tmp12006 := PrimCons(symshen_4_dcall_d, tmp12005)

tmp12007 := PrimCons(symset, tmp12006)

tmp12008 := PrimCons(symshen_4_dcall_d, Nil)

tmp12009 := PrimCons(symvalue, tmp12008)

tmp12010 := Call(__e, PrimFunc(symshen_4prolog_1track), V5802, V5801)


tmp12011 := Call(__e, PrimFunc(symshen_4cons_1form), tmp12010)


tmp12012 := PrimCons(tmp12011, Nil)

tmp12013 := PrimCons(V5800, tmp12012)

tmp12014 := PrimCons(tmp12009, tmp12013)

tmp12015 := PrimCons(symshen_4input_1track, tmp12014)

tmp12016 := PrimCons(symshen_4terpri_1or_1read_1char, Nil)

tmp12017 := PrimCons(symshen_4_dcall_d, Nil)

tmp12018 := PrimCons(symvalue, tmp12017)

tmp12019 := PrimCons(symResult, Nil)

tmp12020 := PrimCons(V5800, tmp12019)

tmp12021 := PrimCons(tmp12018, tmp12020)

tmp12022 := PrimCons(symshen_4output_1track, tmp12021)

tmp12023 := PrimCons(symshen_4_dcall_d, Nil)

tmp12024 := PrimCons(symvalue, tmp12023)

tmp12025 := PrimCons(MakeNumber(1), Nil)

tmp12026 := PrimCons(tmp12024, tmp12025)

tmp12027 := PrimCons(sym_1, tmp12026)

tmp12028 := PrimCons(tmp12027, Nil)

tmp12029 := PrimCons(symshen_4_dcall_d, tmp12028)

tmp12030 := PrimCons(symset, tmp12029)

tmp12031 := PrimCons(symshen_4terpri_1or_1read_1char, Nil)

tmp12032 := PrimCons(symResult, Nil)

tmp12033 := PrimCons(tmp12031, tmp12032)

tmp12034 := PrimCons(symdo, tmp12033)

tmp12035 := PrimCons(tmp12034, Nil)

tmp12036 := PrimCons(tmp12030, tmp12035)

tmp12037 := PrimCons(symdo, tmp12036)

tmp12038 := PrimCons(tmp12037, Nil)

tmp12039 := PrimCons(tmp12022, tmp12038)

tmp12040 := PrimCons(symdo, tmp12039)

tmp12041 := PrimCons(tmp12040, Nil)

tmp12042 := PrimCons(V5802, tmp12041)

tmp12043 := PrimCons(symResult, tmp12042)

tmp12044 := PrimCons(symlet, tmp12043)

tmp12045 := PrimCons(tmp12044, Nil)

tmp12046 := PrimCons(tmp12016, tmp12045)

tmp12047 := PrimCons(symdo, tmp12046)

tmp12048 := PrimCons(tmp12047, Nil)

tmp12049 := PrimCons(tmp12015, tmp12048)

tmp12050 := PrimCons(symdo, tmp12049)

tmp12051 := PrimCons(tmp12050, Nil)

tmp12052 := PrimCons(tmp12007, tmp12051)

__e.Return(PrimCons(symdo, tmp12052))
return


}, 3)

tmp12053 := Call(__e, ns2_1set, symshen_4insert_1tracking_1code, tmp11999)


_ = tmp12053

tmp12054 := MakeNative(func(__e *ControlFlow) {
V5803 := __e.Get(1)
_ = V5803
V5804 := __e.Get(2)
_ = V5804
tmp12057 := Call(__e, PrimFunc(symoccurrences), symshen_4incinfs, V5803)


tmp12058 := PrimEqual(tmp12057, MakeNumber(0))

if True == tmp12058 {
__e.Return(V5804)
return
} else {
tmp12055 := Call(__e, PrimFunc(symshen_4vector_1parameter), V5804)


__e.TailApply(PrimFunc(symshen_4vector_1dereference), V5804, tmp12055)
return


}


}, 2)

tmp12059 := Call(__e, ns2_1set, symshen_4prolog_1track, tmp12054)


_ = tmp12059

tmp12060 := MakeNative(func(__e *ControlFlow) {
V5807 := __e.Get(1)
_ = V5807
tmp12089 := PrimEqual(Nil, V5807)

if True == tmp12089 {
__e.Return(Nil)
return
} else {
tmp12087 := PrimIsPair(V5807)

var ifres12065 Obj

if True == tmp12087 {
tmp12085 := PrimTail(V5807)

tmp12086 := PrimIsPair(tmp12085)

var ifres12067 Obj

if True == tmp12086 {
tmp12082 := PrimTail(V5807)

tmp12083 := PrimTail(tmp12082)

tmp12084 := PrimIsPair(tmp12083)

var ifres12069 Obj

if True == tmp12084 {
tmp12078 := PrimTail(V5807)

tmp12079 := PrimTail(tmp12078)

tmp12080 := PrimTail(tmp12079)

tmp12081 := PrimIsPair(tmp12080)

var ifres12071 Obj

if True == tmp12081 {
tmp12073 := PrimTail(V5807)

tmp12074 := PrimTail(tmp12073)

tmp12075 := PrimTail(tmp12074)

tmp12076 := PrimTail(tmp12075)

tmp12077 := PrimEqual(Nil, tmp12076)

var ifres12072 Obj

if True == tmp12077 {
ifres12072 = True


} else {
ifres12072 = False


}

ifres12071 = ifres12072


} else {
ifres12071 = False


}

var ifres12070 Obj

if True == ifres12071 {
ifres12070 = True


} else {
ifres12070 = False


}

ifres12069 = ifres12070


} else {
ifres12069 = False


}

var ifres12068 Obj

if True == ifres12069 {
ifres12068 = True


} else {
ifres12068 = False


}

ifres12067 = ifres12068


} else {
ifres12067 = False


}

var ifres12066 Obj

if True == ifres12067 {
ifres12066 = True


} else {
ifres12066 = False


}

ifres12065 = ifres12066


} else {
ifres12065 = False


}

if True == ifres12065 {
__e.Return(PrimHead(V5807))
return
} else {
tmp12063 := PrimIsPair(V5807)

if True == tmp12063 {
tmp12061 := PrimTail(V5807)

__e.TailApply(PrimFunc(symshen_4vector_1parameter), tmp12061)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4vector_1parameter)
return
}


}


}


}, 1)

tmp12090 := Call(__e, ns2_1set, symshen_4vector_1parameter, tmp12060)


_ = tmp12090

tmp12091 := MakeNative(func(__e *ControlFlow) {
V5810 := __e.Get(1)
_ = V5810
V5811 := __e.Get(2)
_ = V5811
tmp12125 := PrimEqual(Nil, V5811)

if True == tmp12125 {
__e.Return(V5810)
return
} else {
tmp12123 := PrimIsPair(V5810)

var ifres12101 Obj

if True == tmp12123 {
tmp12121 := PrimTail(V5810)

tmp12122 := PrimIsPair(tmp12121)

var ifres12103 Obj

if True == tmp12122 {
tmp12118 := PrimTail(V5810)

tmp12119 := PrimTail(tmp12118)

tmp12120 := PrimIsPair(tmp12119)

var ifres12105 Obj

if True == tmp12120 {
tmp12114 := PrimTail(V5810)

tmp12115 := PrimTail(tmp12114)

tmp12116 := PrimTail(tmp12115)

tmp12117 := PrimIsPair(tmp12116)

var ifres12107 Obj

if True == tmp12117 {
tmp12109 := PrimTail(V5810)

tmp12110 := PrimTail(tmp12109)

tmp12111 := PrimTail(tmp12110)

tmp12112 := PrimTail(tmp12111)

tmp12113 := PrimEqual(Nil, tmp12112)

var ifres12108 Obj

if True == tmp12113 {
ifres12108 = True


} else {
ifres12108 = False


}

ifres12107 = ifres12108


} else {
ifres12107 = False


}

var ifres12106 Obj

if True == ifres12107 {
ifres12106 = True


} else {
ifres12106 = False


}

ifres12105 = ifres12106


} else {
ifres12105 = False


}

var ifres12104 Obj

if True == ifres12105 {
ifres12104 = True


} else {
ifres12104 = False


}

ifres12103 = ifres12104


} else {
ifres12103 = False


}

var ifres12102 Obj

if True == ifres12103 {
ifres12102 = True


} else {
ifres12102 = False


}

ifres12101 = ifres12102


} else {
ifres12101 = False


}

if True == ifres12101 {
__e.Return(V5810)
return
} else {
tmp12099 := PrimIsPair(V5810)

if True == tmp12099 {
tmp12092 := PrimHead(V5810)

tmp12093 := PrimCons(V5811, Nil)

tmp12094 := PrimCons(tmp12092, tmp12093)

tmp12095 := PrimCons(symshen_4deref, tmp12094)

tmp12096 := PrimTail(V5810)

tmp12097 := Call(__e, PrimFunc(symshen_4vector_1dereference), tmp12096, V5811)


__e.Return(PrimCons(tmp12095, tmp12097))
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4vector_1dereference)
return
}


}


}


}, 2)

tmp12126 := Call(__e, ns2_1set, symshen_4vector_1dereference, tmp12091)


_ = tmp12126

tmp12127 := MakeNative(func(__e *ControlFlow) {
V5814 := __e.Get(1)
_ = V5814
tmp12131 := PrimEqual(sym_7, V5814)

if True == tmp12131 {
__e.Return(PrimSet(symshen_4_dstep_d, True))
return
} else {
tmp12129 := PrimEqual(sym_1, V5814)

if True == tmp12129 {
__e.Return(PrimSet(symshen_4_dstep_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("step expects a + or a -.\n")))
return
}


}


}, 1)

tmp12132 := Call(__e, ns2_1set, symstep, tmp12127)


_ = tmp12132

tmp12133 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dstep_d))
return
}, 0)

tmp12134 := Call(__e, ns2_1set, symshen_4step_2, tmp12133)


_ = tmp12134

tmp12135 := MakeNative(func(__e *ControlFlow) {
V5817 := __e.Get(1)
_ = V5817
tmp12139 := PrimEqual(sym_7, V5817)

if True == tmp12139 {
__e.Return(PrimSet(symshen_4_dspy_d, True))
return
} else {
tmp12137 := PrimEqual(sym_1, V5817)

if True == tmp12137 {
__e.Return(PrimSet(symshen_4_dspy_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("spy expects a + or a -.\n")))
return
}


}


}, 1)

tmp12140 := Call(__e, ns2_1set, symspy, tmp12135)


_ = tmp12140

tmp12141 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dspy_d))
return
}, 0)

tmp12142 := Call(__e, ns2_1set, symshen_4spy_2, tmp12141)


_ = tmp12142

tmp12143 := MakeNative(func(__e *ControlFlow) {
tmp12147 := PrimValue(symshen_4_dstep_d)

if True == tmp12147 {
tmp12144 := PrimValue(sym_dstinput_d)

tmp12145 := PrimReadByte(tmp12144)

__e.TailApply(PrimFunc(symshen_4check_1byte), tmp12145)
return


} else {
__e.TailApply(PrimFunc(symnl), MakeNumber(1))
return
}


}, 0)

tmp12148 := Call(__e, ns2_1set, symshen_4terpri_1or_1read_1char, tmp12143)


_ = tmp12148

tmp12149 := MakeNative(func(__e *ControlFlow) {
V5820 := __e.Get(1)
_ = V5820
tmp12151 := PrimEqual(MakeNumber(94), V5820)

if True == tmp12151 {
__e.Return(PrimSimpleError(MakeString("aborted")))
return
} else {
__e.Return(True)
return
}


}, 1)

tmp12152 := Call(__e, ns2_1set, symshen_4check_1byte, tmp12149)


_ = tmp12152

tmp12153 := MakeNative(func(__e *ControlFlow) {
V5821 := __e.Get(1)
_ = V5821
V5822 := __e.Get(2)
_ = V5822
V5823 := __e.Get(3)
_ = V5823
tmp12154 := Call(__e, PrimFunc(symshen_4spaces), V5821)


tmp12155 := Call(__e, PrimFunc(symshen_4spaces), V5821)


tmp12156 := Call(__e, PrimFunc(symshen_4app), tmp12155, MakeString(""), symshen_4a)


tmp12157 := PrimStringConcat(MakeString(" \n"), tmp12156)

tmp12158 := Call(__e, PrimFunc(symshen_4app), V5822, tmp12157, symshen_4a)


tmp12159 := PrimStringConcat(MakeString("> Inputs to "), tmp12158)

tmp12160 := Call(__e, PrimFunc(symshen_4app), V5821, tmp12159, symshen_4a)


tmp12161 := PrimStringConcat(MakeString("<"), tmp12160)

tmp12162 := Call(__e, PrimFunc(symshen_4app), tmp12154, tmp12161, symshen_4a)


tmp12163 := PrimStringConcat(MakeString("\n"), tmp12162)

tmp12164 := Call(__e, PrimFunc(symstoutput))


tmp12165 := Call(__e, PrimFunc(sympr), tmp12163, tmp12164)


_ = tmp12165

__e.TailApply(PrimFunc(symshen_4recursively_1print), V5823)
return


}, 3)

tmp12166 := Call(__e, ns2_1set, symshen_4input_1track, tmp12153)


_ = tmp12166

tmp12167 := MakeNative(func(__e *ControlFlow) {
V5826 := __e.Get(1)
_ = V5826
tmp12177 := PrimEqual(Nil, V5826)

if True == tmp12177 {
tmp12168 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString(" ==>"), tmp12168)
return


} else {
tmp12175 := PrimIsPair(V5826)

if True == tmp12175 {
tmp12169 := PrimHead(V5826)

tmp12170 := Call(__e, PrimFunc(symprint), tmp12169)


_ = tmp12170

tmp12171 := Call(__e, PrimFunc(symstoutput))


tmp12172 := Call(__e, PrimFunc(sympr), MakeString(", "), tmp12171)


_ = tmp12172

tmp12173 := PrimTail(V5826)

__e.TailApply(PrimFunc(symshen_4recursively_1print), tmp12173)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.recursively-print")))
return
}


}


}, 1)

tmp12178 := Call(__e, ns2_1set, symshen_4recursively_1print, tmp12167)


_ = tmp12178

tmp12179 := MakeNative(func(__e *ControlFlow) {
V5827 := __e.Get(1)
_ = V5827
tmp12183 := PrimEqual(MakeNumber(0), V5827)

if True == tmp12183 {
__e.Return(MakeString(""))
return
} else {
tmp12180 := PrimNumberSubtract(V5827, MakeNumber(1))

tmp12181 := Call(__e, PrimFunc(symshen_4spaces), tmp12180)


__e.Return(PrimStringConcat(MakeString(" "), tmp12181))
return


}


}, 1)

tmp12184 := Call(__e, ns2_1set, symshen_4spaces, tmp12179)


_ = tmp12184

tmp12185 := MakeNative(func(__e *ControlFlow) {
V5828 := __e.Get(1)
_ = V5828
V5829 := __e.Get(2)
_ = V5829
V5830 := __e.Get(3)
_ = V5830
tmp12186 := Call(__e, PrimFunc(symshen_4spaces), V5828)


tmp12187 := Call(__e, PrimFunc(symshen_4spaces), V5828)


tmp12188 := Call(__e, PrimFunc(symshen_4app), V5830, MakeString(""), symshen_4s)


tmp12189 := PrimStringConcat(MakeString("==> "), tmp12188)

tmp12190 := Call(__e, PrimFunc(symshen_4app), tmp12187, tmp12189, symshen_4a)


tmp12191 := PrimStringConcat(MakeString(" \n"), tmp12190)

tmp12192 := Call(__e, PrimFunc(symshen_4app), V5829, tmp12191, symshen_4a)


tmp12193 := PrimStringConcat(MakeString("> Output from "), tmp12192)

tmp12194 := Call(__e, PrimFunc(symshen_4app), V5828, tmp12193, symshen_4a)


tmp12195 := PrimStringConcat(MakeString("<"), tmp12194)

tmp12196 := Call(__e, PrimFunc(symshen_4app), tmp12186, tmp12195, symshen_4a)


tmp12197 := PrimStringConcat(MakeString("\n"), tmp12196)

tmp12198 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp12197, tmp12198)
return


}, 3)

tmp12199 := Call(__e, ns2_1set, symshen_4output_1track, tmp12185)


_ = tmp12199

tmp12200 := MakeNative(func(__e *ControlFlow) {
V5831 := __e.Get(1)
_ = V5831
tmp12201 := PrimValue(symshen_4_dtracking_d)

tmp12202 := Call(__e, PrimFunc(symremove), V5831, tmp12201)


tmp12203 := PrimSet(symshen_4_dtracking_d, tmp12202)

_ = tmp12203

tmp12204 := MakeNative(func(__e *ControlFlow) {
tmp12205 := Call(__e, PrimFunc(symps), V5831)


__e.TailApply(PrimFunc(symeval), tmp12205)
return


}, 0)

tmp12206 := MakeNative(func(__e *ControlFlow) {
Z5832 := __e.Get(1)
_ = Z5832
__e.Return(V5831)
return
}, 1)

tmp12207 := Call(__e, try_1catch, tmp12204, tmp12206)


_ = tmp12207

__e.Return(V5831)
return


}, 1)

tmp12208 := Call(__e, ns2_1set, symuntrack, tmp12200)


_ = tmp12208

tmp12209 := MakeNative(func(__e *ControlFlow) {
V5833 := __e.Get(1)
_ = V5833
V5834 := __e.Get(2)
_ = V5834
__e.TailApply(PrimFunc(symshen_4remove_1h), V5833, V5834, Nil)
return
}, 2)

tmp12210 := Call(__e, ns2_1set, symremove, tmp12209)


_ = tmp12210

tmp12211 := MakeNative(func(__e *ControlFlow) {
V5844 := __e.Get(1)
_ = V5844
V5845 := __e.Get(2)
_ = V5845
V5846 := __e.Get(3)
_ = V5846
tmp12226 := PrimEqual(Nil, V5845)

if True == tmp12226 {
__e.TailApply(PrimFunc(symreverse), V5846)
return
} else {
tmp12224 := PrimIsPair(V5845)

var ifres12220 Obj

if True == tmp12224 {
tmp12222 := PrimHead(V5845)

tmp12223 := PrimEqual(V5844, tmp12222)

var ifres12221 Obj

if True == tmp12223 {
ifres12221 = True


} else {
ifres12221 = False


}

ifres12220 = ifres12221


} else {
ifres12220 = False


}

if True == ifres12220 {
tmp12212 := PrimHead(V5845)

tmp12213 := PrimTail(V5845)

__e.TailApply(PrimFunc(symshen_4remove_1h), tmp12212, tmp12213, V5846)
return


} else {
tmp12218 := PrimIsPair(V5845)

if True == tmp12218 {
tmp12214 := PrimTail(V5845)

tmp12215 := PrimHead(V5845)

tmp12216 := PrimCons(tmp12215, V5846)

__e.TailApply(PrimFunc(symshen_4remove_1h), V5844, tmp12214, tmp12216)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.remove-h")))
return
}


}


}


}, 3)

tmp12227 := Call(__e, ns2_1set, symshen_4remove_1h, tmp12211)


_ = tmp12227

tmp12228 := MakeNative(func(__e *ControlFlow) {
V5847 := __e.Get(1)
_ = V5847
tmp12229 := PrimValue(symshen_4_dprofiled_d)

tmp12230 := PrimCons(V5847, tmp12229)

tmp12231 := PrimSet(symshen_4_dprofiled_d, tmp12230)

_ = tmp12231

tmp12232 := Call(__e, PrimFunc(symps), V5847)


__e.TailApply(PrimFunc(symshen_4profile_1help), tmp12232)
return


}, 1)

tmp12233 := Call(__e, ns2_1set, symprofile, tmp12228)


_ = tmp12233

tmp12234 := MakeNative(func(__e *ControlFlow) {
V5850 := __e.Get(1)
_ = V5850
tmp12304 := PrimIsPair(V5850)

var ifres12278 Obj

if True == tmp12304 {
tmp12302 := PrimHead(V5850)

tmp12303 := PrimEqual(symdefun, tmp12302)

var ifres12280 Obj

if True == tmp12303 {
tmp12300 := PrimTail(V5850)

tmp12301 := PrimIsPair(tmp12300)

var ifres12282 Obj

if True == tmp12301 {
tmp12297 := PrimTail(V5850)

tmp12298 := PrimTail(tmp12297)

tmp12299 := PrimIsPair(tmp12298)

var ifres12284 Obj

if True == tmp12299 {
tmp12293 := PrimTail(V5850)

tmp12294 := PrimTail(tmp12293)

tmp12295 := PrimTail(tmp12294)

tmp12296 := PrimIsPair(tmp12295)

var ifres12286 Obj

if True == tmp12296 {
tmp12288 := PrimTail(V5850)

tmp12289 := PrimTail(tmp12288)

tmp12290 := PrimTail(tmp12289)

tmp12291 := PrimTail(tmp12290)

tmp12292 := PrimEqual(Nil, tmp12291)

var ifres12287 Obj

if True == tmp12292 {
ifres12287 = True


} else {
ifres12287 = False


}

ifres12286 = ifres12287


} else {
ifres12286 = False


}

var ifres12285 Obj

if True == ifres12286 {
ifres12285 = True


} else {
ifres12285 = False


}

ifres12284 = ifres12285


} else {
ifres12284 = False


}

var ifres12283 Obj

if True == ifres12284 {
ifres12283 = True


} else {
ifres12283 = False


}

ifres12282 = ifres12283


} else {
ifres12282 = False


}

var ifres12281 Obj

if True == ifres12282 {
ifres12281 = True


} else {
ifres12281 = False


}

ifres12280 = ifres12281


} else {
ifres12280 = False


}

var ifres12279 Obj

if True == ifres12280 {
ifres12279 = True


} else {
ifres12279 = False


}

ifres12278 = ifres12279


} else {
ifres12278 = False


}

if True == ifres12278 {
tmp12235 := MakeNative(func(__e *ControlFlow) {
W5851 := __e.Get(1)
_ = W5851
tmp12236 := MakeNative(func(__e *ControlFlow) {
W5852 := __e.Get(1)
_ = W5852
tmp12237 := MakeNative(func(__e *ControlFlow) {
W5853 := __e.Get(1)
_ = W5853
tmp12238 := MakeNative(func(__e *ControlFlow) {
W5854 := __e.Get(1)
_ = W5854
tmp12239 := MakeNative(func(__e *ControlFlow) {
W5855 := __e.Get(1)
_ = W5855
tmp12240 := PrimTail(V5850)

__e.Return(PrimHead(tmp12240))
return


}, 1)

tmp12241 := Call(__e, PrimFunc(symeval_1kl), W5853)


__e.TailApply(tmp12239, tmp12241)
return


}, 1)

tmp12242 := Call(__e, PrimFunc(symeval_1kl), W5852)


__e.TailApply(tmp12238, tmp12242)
return


}, 1)

tmp12243 := PrimTail(V5850)

tmp12244 := PrimTail(tmp12243)

tmp12245 := PrimHead(tmp12244)

tmp12246 := PrimTail(V5850)

tmp12247 := PrimHead(tmp12246)

tmp12248 := PrimTail(V5850)

tmp12249 := PrimTail(tmp12248)

tmp12250 := PrimTail(tmp12249)

tmp12251 := PrimHead(tmp12250)

tmp12252 := Call(__e, PrimFunc(symsubst), W5851, tmp12247, tmp12251)


tmp12253 := PrimCons(tmp12252, Nil)

tmp12254 := PrimCons(tmp12245, tmp12253)

tmp12255 := PrimCons(W5851, tmp12254)

tmp12256 := PrimCons(symdefun, tmp12255)

__e.TailApply(tmp12237, tmp12256)
return


}, 1)

tmp12257 := PrimTail(V5850)

tmp12258 := PrimHead(tmp12257)

tmp12259 := PrimTail(V5850)

tmp12260 := PrimTail(tmp12259)

tmp12261 := PrimHead(tmp12260)

tmp12262 := PrimTail(V5850)

tmp12263 := PrimHead(tmp12262)

tmp12264 := PrimTail(V5850)

tmp12265 := PrimTail(tmp12264)

tmp12266 := PrimHead(tmp12265)

tmp12267 := PrimTail(V5850)

tmp12268 := PrimTail(tmp12267)

tmp12269 := PrimHead(tmp12268)

tmp12270 := PrimCons(W5851, tmp12269)

tmp12271 := Call(__e, PrimFunc(symshen_4profile_1func), tmp12263, tmp12266, tmp12270)


tmp12272 := PrimCons(tmp12271, Nil)

tmp12273 := PrimCons(tmp12261, tmp12272)

tmp12274 := PrimCons(tmp12258, tmp12273)

tmp12275 := PrimCons(symdefun, tmp12274)

__e.TailApply(tmp12236, tmp12275)
return


}, 1)

tmp12276 := Call(__e, PrimFunc(symgensym), symshen_4f)


__e.TailApply(tmp12235, tmp12276)
return


} else {
__e.Return(PrimSimpleError(MakeString("Cannot profile.\n")))
return
}


}, 1)

tmp12305 := Call(__e, ns2_1set, symshen_4profile_1help, tmp12234)


_ = tmp12305

tmp12306 := MakeNative(func(__e *ControlFlow) {
V5856 := __e.Get(1)
_ = V5856
tmp12307 := PrimValue(symshen_4_dprofiled_d)

tmp12308 := Call(__e, PrimFunc(symremove), V5856, tmp12307)


tmp12309 := PrimSet(symshen_4_dprofiled_d, tmp12308)

_ = tmp12309

tmp12310 := MakeNative(func(__e *ControlFlow) {
tmp12311 := Call(__e, PrimFunc(symps), V5856)


__e.TailApply(PrimFunc(symeval), tmp12311)
return


}, 0)

tmp12312 := MakeNative(func(__e *ControlFlow) {
Z5857 := __e.Get(1)
_ = Z5857
__e.Return(V5856)
return
}, 1)

__e.TailApply(try_1catch, tmp12310, tmp12312)
return


}, 1)

tmp12313 := Call(__e, ns2_1set, symunprofile, tmp12306)


_ = tmp12313

tmp12314 := MakeNative(func(__e *ControlFlow) {
V5858 := __e.Get(1)
_ = V5858
tmp12315 := PrimValue(symshen_4_dprofiled_d)

__e.TailApply(PrimFunc(symelement_2), V5858, tmp12315)
return


}, 1)

tmp12316 := Call(__e, ns2_1set, symshen_4profiled_2, tmp12314)


_ = tmp12316

tmp12317 := MakeNative(func(__e *ControlFlow) {
V5859 := __e.Get(1)
_ = V5859
V5860 := __e.Get(2)
_ = V5860
V5861 := __e.Get(3)
_ = V5861
tmp12318 := PrimCons(symrun, Nil)

tmp12319 := PrimCons(symget_1time, tmp12318)

tmp12320 := PrimCons(symrun, Nil)

tmp12321 := PrimCons(symget_1time, tmp12320)

tmp12322 := PrimCons(symStart, Nil)

tmp12323 := PrimCons(tmp12321, tmp12322)

tmp12324 := PrimCons(sym_1, tmp12323)

tmp12325 := PrimCons(V5859, Nil)

tmp12326 := PrimCons(symshen_4get_1profile, tmp12325)

tmp12327 := PrimCons(symFinish, Nil)

tmp12328 := PrimCons(tmp12326, tmp12327)

tmp12329 := PrimCons(sym_7, tmp12328)

tmp12330 := PrimCons(tmp12329, Nil)

tmp12331 := PrimCons(V5859, tmp12330)

tmp12332 := PrimCons(symshen_4put_1profile, tmp12331)

tmp12333 := PrimCons(symResult, Nil)

tmp12334 := PrimCons(tmp12332, tmp12333)

tmp12335 := PrimCons(symRecord, tmp12334)

tmp12336 := PrimCons(symlet, tmp12335)

tmp12337 := PrimCons(tmp12336, Nil)

tmp12338 := PrimCons(tmp12324, tmp12337)

tmp12339 := PrimCons(symFinish, tmp12338)

tmp12340 := PrimCons(symlet, tmp12339)

tmp12341 := PrimCons(tmp12340, Nil)

tmp12342 := PrimCons(V5861, tmp12341)

tmp12343 := PrimCons(symResult, tmp12342)

tmp12344 := PrimCons(symlet, tmp12343)

tmp12345 := PrimCons(tmp12344, Nil)

tmp12346 := PrimCons(tmp12319, tmp12345)

tmp12347 := PrimCons(symStart, tmp12346)

__e.Return(PrimCons(symlet, tmp12347))
return


}, 3)

tmp12348 := Call(__e, ns2_1set, symshen_4profile_1func, tmp12317)


_ = tmp12348

tmp12349 := MakeNative(func(__e *ControlFlow) {
V5862 := __e.Get(1)
_ = V5862
tmp12350 := MakeNative(func(__e *ControlFlow) {
W5863 := __e.Get(1)
_ = W5863
tmp12351 := MakeNative(func(__e *ControlFlow) {
W5864 := __e.Get(1)
_ = W5864
__e.TailApply(PrimFunc(sym_8p), V5862, W5863)
return
}, 1)

tmp12352 := Call(__e, PrimFunc(symshen_4put_1profile), V5862, MakeNumber(0))


__e.TailApply(tmp12351, tmp12352)
return


}, 1)

tmp12353 := Call(__e, PrimFunc(symshen_4get_1profile), V5862)


__e.TailApply(tmp12350, tmp12353)
return


}, 1)

tmp12354 := Call(__e, ns2_1set, symprofile_1results, tmp12349)


_ = tmp12354

tmp12355 := MakeNative(func(__e *ControlFlow) {
V5865 := __e.Get(1)
_ = V5865
tmp12356 := MakeNative(func(__e *ControlFlow) {
tmp12357 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V5865, symprofile, tmp12357)
return


}, 0)

tmp12358 := MakeNative(func(__e *ControlFlow) {
Z5866 := __e.Get(1)
_ = Z5866
__e.Return(MakeNumber(0))
return
}, 1)

__e.TailApply(try_1catch, tmp12356, tmp12358)
return


}, 1)

tmp12359 := Call(__e, ns2_1set, symshen_4get_1profile, tmp12355)


_ = tmp12359

tmp12360 := MakeNative(func(__e *ControlFlow) {
V5867 := __e.Get(1)
_ = V5867
V5868 := __e.Get(2)
_ = V5868
tmp12361 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V5867, symprofile, V5868, tmp12361)
return


}, 2)

__e.TailApply(ns2_1set, symshen_4put_1profile, tmp12360)
return




}, 0)

