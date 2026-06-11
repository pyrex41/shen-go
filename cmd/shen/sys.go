package main

import . "github.com/tiancaiamao/shen-go/kl"

var SysMain = MakeNative(func(__e *ControlFlow) {
tmp75 := MakeNative(func(__e *ControlFlow) {
V3767 := __e.Get(1)
_ = V3767
__e.TailApply(V3767)
return
}, 1)

tmp76 := Call(__e, ns2_1set, symthaw, tmp75)


_ = tmp76

tmp77 := MakeNative(func(__e *ControlFlow) {
V3768 := __e.Get(1)
_ = V3768
tmp78 := Call(__e, PrimFunc(symmacroexpand), V3768)


tmp79 := Call(__e, PrimFunc(symshen_4find_1types), V3768)


tmp80 := Call(__e, PrimFunc(symshen_4process_1applications), tmp78, tmp79)


tmp81 := Call(__e, PrimFunc(symshen_4shen_1_6kl), tmp80)


__e.TailApply(PrimFunc(symeval_1kl), tmp81)
return


}, 1)

tmp82 := Call(__e, ns2_1set, symeval, tmp77)


_ = tmp82

tmp83 := MakeNative(func(__e *ControlFlow) {
V3769 := __e.Get(1)
_ = V3769
tmp90 := PrimEqual(symnull, V3769)

if True == tmp90 {
__e.Return(Nil)
return
} else {
tmp84 := MakeNative(func(__e *ControlFlow) {
tmp85 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V3769, symshen_4external_1symbols, tmp85)
return


}, 0)

tmp86 := MakeNative(func(__e *ControlFlow) {
Z3770 := __e.Get(1)
_ = Z3770
tmp87 := Call(__e, PrimFunc(symshen_4app), V3769, MakeString(" does not exist.\n;"), symshen_4a)


tmp88 := PrimStringConcat(MakeString("package "), tmp87)

__e.Return(PrimSimpleError(tmp88))
return


}, 1)

__e.TailApply(try_1catch, tmp84, tmp86)
return


}


}, 1)

tmp91 := Call(__e, ns2_1set, symexternal, tmp83)


_ = tmp91

tmp92 := MakeNative(func(__e *ControlFlow) {
V3771 := __e.Get(1)
_ = V3771
tmp99 := PrimEqual(symnull, V3771)

if True == tmp99 {
__e.Return(Nil)
return
} else {
tmp93 := MakeNative(func(__e *ControlFlow) {
tmp94 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V3771, symshen_4internal_1symbols, tmp94)
return


}, 0)

tmp95 := MakeNative(func(__e *ControlFlow) {
Z3772 := __e.Get(1)
_ = Z3772
tmp96 := Call(__e, PrimFunc(symshen_4app), V3771, MakeString(" does not exist.\n;"), symshen_4a)


tmp97 := PrimStringConcat(MakeString("package "), tmp96)

__e.Return(PrimSimpleError(tmp97))
return


}, 1)

__e.TailApply(try_1catch, tmp93, tmp95)
return


}


}, 1)

tmp100 := Call(__e, ns2_1set, syminternal, tmp92)


_ = tmp100

tmp101 := MakeNative(func(__e *ControlFlow) {
V3773 := __e.Get(1)
_ = V3773
V3774 := __e.Get(2)
_ = V3774
tmp103 := Call(__e, V3773, V3774)


if True == tmp103 {
__e.TailApply(PrimFunc(symfail))
return
} else {
__e.Return(V3774)
return
}


}, 2)

tmp104 := Call(__e, ns2_1set, symfail_1if, tmp101)


_ = tmp104

tmp105 := MakeNative(func(__e *ControlFlow) {
V3775 := __e.Get(1)
_ = V3775
V3776 := __e.Get(2)
_ = V3776
__e.Return(PrimStringConcat(V3775, V3776))
return
}, 2)

tmp106 := Call(__e, ns2_1set, sym_8s, tmp105)


_ = tmp106

tmp107 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dtc_d))
return
}, 0)

tmp108 := Call(__e, ns2_1set, symtc_2, tmp107)


_ = tmp108

tmp109 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_doccurs_d))
return
}, 0)

tmp110 := Call(__e, ns2_1set, symoccurs_2, tmp109)


_ = tmp110

tmp111 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dfactorise_2_d))
return
}, 0)

tmp112 := Call(__e, ns2_1set, symfactorise_2, tmp111)


_ = tmp112

tmp113 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dtracking_d))
return
}, 0)

tmp114 := Call(__e, ns2_1set, symtracked, tmp113)


_ = tmp114

tmp115 := MakeNative(func(__e *ControlFlow) {
V3777 := __e.Get(1)
_ = V3777
tmp116 := MakeNative(func(__e *ControlFlow) {
tmp117 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V3777, symshen_4source, tmp117)
return


}, 0)

tmp118 := MakeNative(func(__e *ControlFlow) {
Z3778 := __e.Get(1)
_ = Z3778
tmp119 := Call(__e, PrimFunc(symshen_4app), V3777, MakeString(" not found.\n"), symshen_4a)


__e.Return(PrimSimpleError(tmp119))
return


}, 1)

__e.TailApply(try_1catch, tmp116, tmp118)
return


}, 1)

tmp120 := Call(__e, ns2_1set, symps, tmp115)


_ = tmp120

tmp121 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dstinput_d))
return
}, 0)

tmp122 := Call(__e, ns2_1set, symstinput, tmp121)


_ = tmp122

tmp123 := MakeNative(func(__e *ControlFlow) {
V3779 := __e.Get(1)
_ = V3779
tmp124 := MakeNative(func(__e *ControlFlow) {
W3780 := __e.Get(1)
_ = W3780
tmp125 := MakeNative(func(__e *ControlFlow) {
W3781 := __e.Get(1)
_ = W3781
tmp126 := MakeNative(func(__e *ControlFlow) {
W3782 := __e.Get(1)
_ = W3782
__e.Return(W3782)
return
}, 1)

tmp130 := PrimEqual(V3779, MakeNumber(0))

var ifres127 Obj

if True == tmp130 {
ifres127 = W3781


} else {
tmp128 := Call(__e, PrimFunc(symfail))


tmp129 := Call(__e, PrimFunc(symshen_4fillvector), W3781, MakeNumber(1), V3779, tmp128)


ifres127 = tmp129


}

__e.TailApply(tmp126, ifres127)
return


}, 1)

tmp131 := PrimVectorSet(W3780, MakeNumber(0), V3779)

__e.TailApply(tmp125, tmp131)
return


}, 1)

tmp132 := PrimNumberAdd(V3779, MakeNumber(1))

tmp133 := PrimAbsvector(tmp132)

__e.TailApply(tmp124, tmp133)
return


}, 1)

tmp134 := Call(__e, ns2_1set, symvector, tmp123)


_ = tmp134

tmp135 := MakeNative(func(__e *ControlFlow) {
V3784 := __e.Get(1)
_ = V3784
V3785 := __e.Get(2)
_ = V3785
V3786 := __e.Get(3)
_ = V3786
V3787 := __e.Get(4)
_ = V3787
tmp139 := PrimEqual(V3785, V3786)

if True == tmp139 {
__e.Return(PrimVectorSet(V3784, V3786, V3787))
return
} else {
tmp136 := PrimVectorSet(V3784, V3785, V3787)

tmp137 := PrimNumberAdd(MakeNumber(1), V3785)

__e.TailApply(PrimFunc(symshen_4fillvector), tmp136, tmp137, V3786, V3787)
return


}


}, 4)

tmp140 := Call(__e, ns2_1set, symshen_4fillvector, tmp135)


_ = tmp140

tmp141 := MakeNative(func(__e *ControlFlow) {
V3788 := __e.Get(1)
_ = V3788
tmp153 := PrimIsVector(V3788)

if True == tmp153 {
tmp143 := MakeNative(func(__e *ControlFlow) {
W3789 := __e.Get(1)
_ = W3789
tmp147 := PrimIsNumber(W3789)

if True == tmp147 {
tmp145 := PrimGreatEqual(W3789, MakeNumber(0))

if True == tmp145 {
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

tmp148 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimVectorGet(V3788, MakeNumber(0)))
return
}, 0)

tmp149 := MakeNative(func(__e *ControlFlow) {
Z3790 := __e.Get(1)
_ = Z3790
__e.Return(MakeNumber(-1))
return
}, 1)

tmp150 := Call(__e, try_1catch, tmp148, tmp149)


tmp151 := Call(__e, tmp143, tmp150)


if True == tmp151 {
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

tmp154 := Call(__e, ns2_1set, symvector_2, tmp141)


_ = tmp154

tmp155 := MakeNative(func(__e *ControlFlow) {
V3791 := __e.Get(1)
_ = V3791
V3792 := __e.Get(2)
_ = V3792
V3793 := __e.Get(3)
_ = V3793
tmp157 := PrimEqual(V3792, MakeNumber(0))

if True == tmp157 {
__e.Return(PrimSimpleError(MakeString("cannot access 0th element of a vector\n")))
return
} else {
__e.Return(PrimVectorSet(V3791, V3792, V3793))
return
}


}, 3)

tmp158 := Call(__e, ns2_1set, symvector_1_6, tmp155)


_ = tmp158

tmp159 := MakeNative(func(__e *ControlFlow) {
V3794 := __e.Get(1)
_ = V3794
V3795 := __e.Get(2)
_ = V3795
tmp166 := PrimEqual(V3795, MakeNumber(0))

if True == tmp166 {
__e.Return(PrimSimpleError(MakeString("cannot access 0th element of a vector\n")))
return
} else {
tmp160 := MakeNative(func(__e *ControlFlow) {
W3796 := __e.Get(1)
_ = W3796
tmp162 := Call(__e, PrimFunc(symfail))


tmp163 := PrimEqual(W3796, tmp162)

if True == tmp163 {
__e.Return(PrimSimpleError(MakeString("vector element not found\n")))
return
} else {
__e.Return(W3796)
return
}


}, 1)

tmp164 := PrimVectorGet(V3794, V3795)

__e.TailApply(tmp160, tmp164)
return


}


}, 2)

tmp167 := Call(__e, ns2_1set, sym_5_1vector, tmp159)


_ = tmp167

tmp168 := MakeNative(func(__e *ControlFlow) {
V3797 := __e.Get(1)
_ = V3797
tmp172 := PrimIsInteger(V3797)

if True == tmp172 {
tmp170 := PrimGreatEqual(V3797, MakeNumber(0))

if True == tmp170 {
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

tmp173 := Call(__e, ns2_1set, symshen_4posint_2, tmp168)


_ = tmp173

tmp174 := MakeNative(func(__e *ControlFlow) {
V3798 := __e.Get(1)
_ = V3798
__e.Return(PrimVectorGet(V3798, MakeNumber(0)))
return
}, 1)

tmp175 := Call(__e, ns2_1set, symlimit, tmp174)


_ = tmp175

tmp176 := MakeNative(func(__e *ControlFlow) {
V3799 := __e.Get(1)
_ = V3799
tmp207 := Call(__e, PrimFunc(symboolean_2), V3799)


var ifres192 Obj

if True == tmp207 {
ifres192 = True


} else {
tmp206 := PrimIsNumber(V3799)

var ifres194 Obj

if True == tmp206 {
ifres194 = True


} else {
tmp205 := PrimIsString(V3799)

var ifres196 Obj

if True == tmp205 {
ifres196 = True


} else {
tmp204 := PrimIsPair(V3799)

var ifres198 Obj

if True == tmp204 {
ifres198 = True


} else {
tmp203 := Call(__e, PrimFunc(symempty_2), V3799)


var ifres200 Obj

if True == tmp203 {
ifres200 = True


} else {
tmp202 := Call(__e, PrimFunc(symvector_2), V3799)


var ifres201 Obj

if True == tmp202 {
ifres201 = True


} else {
ifres201 = False


}

ifres200 = ifres201


}

var ifres199 Obj

if True == ifres200 {
ifres199 = True


} else {
ifres199 = False


}

ifres198 = ifres199


}

var ifres197 Obj

if True == ifres198 {
ifres197 = True


} else {
ifres197 = False


}

ifres196 = ifres197


}

var ifres195 Obj

if True == ifres196 {
ifres195 = True


} else {
ifres195 = False


}

ifres194 = ifres195


}

var ifres193 Obj

if True == ifres194 {
ifres193 = True


} else {
ifres193 = False


}

ifres192 = ifres193


}

if True == ifres192 {
__e.Return(False)
return
} else {
tmp182 := PrimIntern(MakeString(":"))

tmp183 := PrimIntern(MakeString(";"))

tmp184 := PrimIntern(MakeString(","))

tmp185 := PrimCons(tmp184, Nil)

tmp186 := PrimCons(tmp183, tmp185)

tmp187 := PrimCons(tmp182, tmp186)

tmp188 := PrimCons(sym_j, tmp187)

tmp189 := PrimCons(sym_i, tmp188)

tmp190 := Call(__e, PrimFunc(symelement_2), V3799, tmp189)


if True == tmp190 {
__e.Return(True)
return
} else {
tmp177 := MakeNative(func(__e *ControlFlow) {
tmp178 := MakeNative(func(__e *ControlFlow) {
W3800 := __e.Get(1)
_ = W3800
__e.TailApply(PrimFunc(symshen_4analyse_1symbol_2), W3800)
return
}, 1)

tmp179 := PrimStr(V3799)

__e.TailApply(tmp178, tmp179)
return


}, 0)

tmp180 := MakeNative(func(__e *ControlFlow) {
Z3801 := __e.Get(1)
_ = Z3801
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp177, tmp180)
return


}


}


}, 1)

tmp208 := Call(__e, ns2_1set, symsymbol_2, tmp176)


_ = tmp208

tmp209 := MakeNative(func(__e *ControlFlow) {
V3804 := __e.Get(1)
_ = V3804
tmp218 := Call(__e, PrimFunc(symshen_4_7string_2), V3804)


if True == tmp218 {
tmp214 := Call(__e, PrimFunc(symhdstr), V3804)


tmp215 := PrimStringToNumber(tmp214)

tmp216 := Call(__e, PrimFunc(symshen_4alpha_2), tmp215)


if True == tmp216 {
tmp211 := PrimTailString(V3804)

tmp212 := Call(__e, PrimFunc(symshen_4alphanums_2), tmp211)


if True == tmp212 {
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
__e.Return(PrimSimpleError(MakeString("implementation error in shen.analyse-symbol?")))
return
}


}, 1)

tmp219 := Call(__e, ns2_1set, symshen_4analyse_1symbol_2, tmp209)


_ = tmp219

tmp220 := MakeNative(func(__e *ControlFlow) {
V3807 := __e.Get(1)
_ = V3807
tmp235 := PrimEqual(MakeString(""), V3807)

if True == tmp235 {
__e.Return(True)
return
} else {
tmp233 := Call(__e, PrimFunc(symshen_4_7string_2), V3807)


if True == tmp233 {
tmp221 := MakeNative(func(__e *ControlFlow) {
W3808 := __e.Get(1)
_ = W3808
tmp229 := Call(__e, PrimFunc(symshen_4alpha_2), W3808)


var ifres226 Obj

if True == tmp229 {
ifres226 = True


} else {
tmp228 := Call(__e, PrimFunc(symshen_4digit_2), W3808)


var ifres227 Obj

if True == tmp228 {
ifres227 = True


} else {
ifres227 = False


}

ifres226 = ifres227


}

if True == ifres226 {
tmp223 := PrimTailString(V3807)

tmp224 := Call(__e, PrimFunc(symshen_4alphanums_2), tmp223)


if True == tmp224 {
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

tmp230 := Call(__e, PrimFunc(symhdstr), V3807)


tmp231 := PrimStringToNumber(tmp230)

__e.TailApply(tmp221, tmp231)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.alphanums?")))
return
}


}


}, 1)

tmp236 := Call(__e, ns2_1set, symshen_4alphanums_2, tmp220)


_ = tmp236

tmp237 := MakeNative(func(__e *ControlFlow) {
V3809 := __e.Get(1)
_ = V3809
tmp249 := Call(__e, PrimFunc(symboolean_2), V3809)


var ifres243 Obj

if True == tmp249 {
ifres243 = True


} else {
tmp248 := PrimIsNumber(V3809)

var ifres245 Obj

if True == tmp248 {
ifres245 = True


} else {
tmp247 := PrimIsString(V3809)

var ifres246 Obj

if True == tmp247 {
ifres246 = True


} else {
ifres246 = False


}

ifres245 = ifres246


}

var ifres244 Obj

if True == ifres245 {
ifres244 = True


} else {
ifres244 = False


}

ifres243 = ifres244


}

if True == ifres243 {
__e.Return(False)
return
} else {
tmp238 := MakeNative(func(__e *ControlFlow) {
tmp239 := MakeNative(func(__e *ControlFlow) {
W3810 := __e.Get(1)
_ = W3810
__e.TailApply(PrimFunc(symshen_4analyse_1variable_2), W3810)
return
}, 1)

tmp240 := PrimStr(V3809)

__e.TailApply(tmp239, tmp240)
return


}, 0)

tmp241 := MakeNative(func(__e *ControlFlow) {
Z3811 := __e.Get(1)
_ = Z3811
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp238, tmp241)
return


}


}, 1)

tmp250 := Call(__e, ns2_1set, symvariable_2, tmp237)


_ = tmp250

tmp251 := MakeNative(func(__e *ControlFlow) {
V3814 := __e.Get(1)
_ = V3814
tmp260 := Call(__e, PrimFunc(symshen_4_7string_2), V3814)


if True == tmp260 {
tmp256 := Call(__e, PrimFunc(symhdstr), V3814)


tmp257 := PrimStringToNumber(tmp256)

tmp258 := Call(__e, PrimFunc(symshen_4uppercase_2), tmp257)


if True == tmp258 {
tmp253 := PrimTailString(V3814)

tmp254 := Call(__e, PrimFunc(symshen_4alphanums_2), tmp253)


if True == tmp254 {
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
__e.Return(PrimSimpleError(MakeString("implementation error in shen.analyse-variable?")))
return
}


}, 1)

tmp261 := Call(__e, ns2_1set, symshen_4analyse_1variable_2, tmp251)


_ = tmp261

tmp262 := MakeNative(func(__e *ControlFlow) {
V3815 := __e.Get(1)
_ = V3815
tmp263 := PrimValue(symshen_4_dgensym_d)

tmp264 := PrimNumberAdd(MakeNumber(1), tmp263)

tmp265 := PrimSet(symshen_4_dgensym_d, tmp264)

__e.TailApply(PrimFunc(symconcat), V3815, tmp265)
return


}, 1)

tmp266 := Call(__e, ns2_1set, symgensym, tmp262)


_ = tmp266

tmp267 := MakeNative(func(__e *ControlFlow) {
V3816 := __e.Get(1)
_ = V3816
V3817 := __e.Get(2)
_ = V3817
tmp268 := PrimStr(V3816)

tmp269 := PrimStr(V3817)

tmp270 := PrimStringConcat(tmp268, tmp269)

__e.Return(PrimIntern(tmp270))
return


}, 2)

tmp271 := Call(__e, ns2_1set, symconcat, tmp267)


_ = tmp271

tmp272 := MakeNative(func(__e *ControlFlow) {
V3818 := __e.Get(1)
_ = V3818
V3819 := __e.Get(2)
_ = V3819
tmp273 := MakeNative(func(__e *ControlFlow) {
W3820 := __e.Get(1)
_ = W3820
tmp274 := MakeNative(func(__e *ControlFlow) {
W3821 := __e.Get(1)
_ = W3821
tmp275 := MakeNative(func(__e *ControlFlow) {
W3822 := __e.Get(1)
_ = W3822
tmp276 := MakeNative(func(__e *ControlFlow) {
W3823 := __e.Get(1)
_ = W3823
__e.Return(W3820)
return
}, 1)

tmp277 := PrimVectorSet(W3820, MakeNumber(2), V3819)

__e.TailApply(tmp276, tmp277)
return


}, 1)

tmp278 := PrimVectorSet(W3820, MakeNumber(1), V3818)

__e.TailApply(tmp275, tmp278)
return


}, 1)

tmp279 := PrimVectorSet(W3820, MakeNumber(0), symshen_4tuple)

__e.TailApply(tmp274, tmp279)
return


}, 1)

tmp280 := PrimAbsvector(MakeNumber(3))

__e.TailApply(tmp273, tmp280)
return


}, 2)

tmp281 := Call(__e, ns2_1set, sym_8p, tmp272)


_ = tmp281

tmp282 := MakeNative(func(__e *ControlFlow) {
V3824 := __e.Get(1)
_ = V3824
__e.Return(PrimVectorGet(V3824, MakeNumber(1)))
return
}, 1)

tmp283 := Call(__e, ns2_1set, symfst, tmp282)


_ = tmp283

tmp284 := MakeNative(func(__e *ControlFlow) {
V3825 := __e.Get(1)
_ = V3825
__e.Return(PrimVectorGet(V3825, MakeNumber(2)))
return
}, 1)

tmp285 := Call(__e, ns2_1set, symsnd, tmp284)


_ = tmp285

tmp286 := MakeNative(func(__e *ControlFlow) {
V3826 := __e.Get(1)
_ = V3826
tmp293 := PrimIsVector(V3826)

if True == tmp293 {
tmp288 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimVectorGet(V3826, MakeNumber(0)))
return
}, 0)

tmp289 := MakeNative(func(__e *ControlFlow) {
Z3827 := __e.Get(1)
_ = Z3827
__e.Return(symshen_4not_1tuple)
return
}, 1)

tmp290 := Call(__e, try_1catch, tmp288, tmp289)


tmp291 := PrimEqual(symshen_4tuple, tmp290)

if True == tmp291 {
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

tmp294 := Call(__e, ns2_1set, symtuple_2, tmp286)


_ = tmp294

tmp295 := MakeNative(func(__e *ControlFlow) {
V3832 := __e.Get(1)
_ = V3832
V3833 := __e.Get(2)
_ = V3833
tmp302 := PrimEqual(Nil, V3832)

if True == tmp302 {
__e.Return(V3833)
return
} else {
tmp300 := PrimIsPair(V3832)

if True == tmp300 {
tmp296 := PrimHead(V3832)

tmp297 := PrimTail(V3832)

tmp298 := Call(__e, PrimFunc(symappend), tmp297, V3833)


__e.Return(PrimCons(tmp296, tmp298))
return


} else {
__e.Return(PrimSimpleError(MakeString("attempt to append a non-list")))
return
}


}


}, 2)

tmp303 := Call(__e, ns2_1set, symappend, tmp295)


_ = tmp303

tmp304 := MakeNative(func(__e *ControlFlow) {
V3834 := __e.Get(1)
_ = V3834
V3835 := __e.Get(2)
_ = V3835
tmp305 := MakeNative(func(__e *ControlFlow) {
W3836 := __e.Get(1)
_ = W3836
tmp306 := MakeNative(func(__e *ControlFlow) {
W3837 := __e.Get(1)
_ = W3837
tmp307 := MakeNative(func(__e *ControlFlow) {
W3838 := __e.Get(1)
_ = W3838
tmp309 := PrimEqual(W3836, MakeNumber(0))

if True == tmp309 {
__e.Return(W3838)
return
} else {
__e.TailApply(PrimFunc(symshen_4_8v_1help), V3835, MakeNumber(1), W3836, W3838)
return
}


}, 1)

tmp310 := Call(__e, PrimFunc(symvector_1_6), W3837, MakeNumber(1), V3834)


__e.TailApply(tmp307, tmp310)
return


}, 1)

tmp311 := PrimNumberAdd(W3836, MakeNumber(1))

tmp312 := Call(__e, PrimFunc(symvector), tmp311)


__e.TailApply(tmp306, tmp312)
return


}, 1)

tmp313 := Call(__e, PrimFunc(symlimit), V3835)


__e.TailApply(tmp305, tmp313)
return


}, 2)

tmp314 := Call(__e, ns2_1set, sym_8v, tmp304)


_ = tmp314

tmp315 := MakeNative(func(__e *ControlFlow) {
V3840 := __e.Get(1)
_ = V3840
V3841 := __e.Get(2)
_ = V3841
V3842 := __e.Get(3)
_ = V3842
V3843 := __e.Get(4)
_ = V3843
tmp321 := PrimEqual(V3841, V3842)

if True == tmp321 {
tmp316 := PrimNumberAdd(V3842, MakeNumber(1))

__e.TailApply(PrimFunc(symshen_4copyfromvector), V3840, V3843, V3842, tmp316)
return


} else {
tmp317 := PrimNumberAdd(V3841, MakeNumber(1))

tmp318 := PrimNumberAdd(V3841, MakeNumber(1))

tmp319 := Call(__e, PrimFunc(symshen_4copyfromvector), V3840, V3843, V3841, tmp318)


__e.TailApply(PrimFunc(symshen_4_8v_1help), V3840, tmp317, V3842, tmp319)
return


}


}, 4)

tmp322 := Call(__e, ns2_1set, symshen_4_8v_1help, tmp315)


_ = tmp322

tmp323 := MakeNative(func(__e *ControlFlow) {
V3844 := __e.Get(1)
_ = V3844
V3845 := __e.Get(2)
_ = V3845
V3846 := __e.Get(3)
_ = V3846
V3847 := __e.Get(4)
_ = V3847
tmp324 := MakeNative(func(__e *ControlFlow) {
tmp325 := Call(__e, PrimFunc(sym_5_1vector), V3844, V3846)


__e.TailApply(PrimFunc(symvector_1_6), V3845, V3847, tmp325)
return


}, 0)

tmp326 := MakeNative(func(__e *ControlFlow) {
Z3848 := __e.Get(1)
_ = Z3848
__e.Return(V3845)
return
}, 1)

__e.TailApply(try_1catch, tmp324, tmp326)
return


}, 4)

tmp327 := Call(__e, ns2_1set, symshen_4copyfromvector, tmp323)


_ = tmp327

tmp328 := MakeNative(func(__e *ControlFlow) {
V3849 := __e.Get(1)
_ = V3849
tmp329 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(sym_5_1vector), V3849, MakeNumber(1))
return
}, 0)

tmp330 := MakeNative(func(__e *ControlFlow) {
Z3850 := __e.Get(1)
_ = Z3850
__e.Return(PrimSimpleError(MakeString("hdv needs a non-empty vector as an argument\n")))
return
}, 1)

__e.TailApply(try_1catch, tmp329, tmp330)
return


}, 1)

tmp331 := Call(__e, ns2_1set, symhdv, tmp328)


_ = tmp331

tmp332 := MakeNative(func(__e *ControlFlow) {
V3851 := __e.Get(1)
_ = V3851
tmp333 := MakeNative(func(__e *ControlFlow) {
W3852 := __e.Get(1)
_ = W3852
tmp342 := PrimEqual(W3852, MakeNumber(0))

if True == tmp342 {
__e.Return(PrimSimpleError(MakeString("cannot take the tail of the empty vector\n")))
return
} else {
tmp340 := PrimEqual(W3852, MakeNumber(1))

if True == tmp340 {
__e.TailApply(PrimFunc(symvector), MakeNumber(0))
return
} else {
tmp334 := MakeNative(func(__e *ControlFlow) {
W3853 := __e.Get(1)
_ = W3853
tmp335 := PrimNumberSubtract(W3852, MakeNumber(1))

tmp336 := Call(__e, PrimFunc(symvector), tmp335)


__e.TailApply(PrimFunc(symshen_4tlv_1help), V3851, MakeNumber(2), W3852, tmp336)
return


}, 1)

tmp337 := PrimNumberSubtract(W3852, MakeNumber(1))

tmp338 := Call(__e, PrimFunc(symvector), tmp337)


__e.TailApply(tmp334, tmp338)
return


}


}


}, 1)

tmp343 := Call(__e, PrimFunc(symlimit), V3851)


__e.TailApply(tmp333, tmp343)
return


}, 1)

tmp344 := Call(__e, ns2_1set, symtlv, tmp332)


_ = tmp344

tmp345 := MakeNative(func(__e *ControlFlow) {
V3855 := __e.Get(1)
_ = V3855
V3856 := __e.Get(2)
_ = V3856
V3857 := __e.Get(3)
_ = V3857
V3858 := __e.Get(4)
_ = V3858
tmp351 := PrimEqual(V3856, V3857)

if True == tmp351 {
tmp346 := PrimNumberSubtract(V3857, MakeNumber(1))

__e.TailApply(PrimFunc(symshen_4copyfromvector), V3855, V3858, V3857, tmp346)
return


} else {
tmp347 := PrimNumberAdd(V3856, MakeNumber(1))

tmp348 := PrimNumberSubtract(V3856, MakeNumber(1))

tmp349 := Call(__e, PrimFunc(symshen_4copyfromvector), V3855, V3858, V3856, tmp348)


__e.TailApply(PrimFunc(symshen_4tlv_1help), V3855, tmp347, V3857, tmp349)
return


}


}, 4)

tmp352 := Call(__e, ns2_1set, symshen_4tlv_1help, tmp345)


_ = tmp352

tmp353 := MakeNative(func(__e *ControlFlow) {
V3870 := __e.Get(1)
_ = V3870
V3871 := __e.Get(2)
_ = V3871
tmp369 := PrimEqual(Nil, V3871)

if True == tmp369 {
__e.Return(Nil)
return
} else {
tmp367 := PrimIsPair(V3871)

var ifres358 Obj

if True == tmp367 {
tmp365 := PrimHead(V3871)

tmp366 := PrimIsPair(tmp365)

var ifres360 Obj

if True == tmp366 {
tmp362 := PrimHead(V3871)

tmp363 := PrimHead(tmp362)

tmp364 := PrimEqual(V3870, tmp363)

var ifres361 Obj

if True == tmp364 {
ifres361 = True


} else {
ifres361 = False


}

ifres360 = ifres361


} else {
ifres360 = False


}

var ifres359 Obj

if True == ifres360 {
ifres359 = True


} else {
ifres359 = False


}

ifres358 = ifres359


} else {
ifres358 = False


}

if True == ifres358 {
__e.Return(PrimHead(V3871))
return
} else {
tmp356 := PrimIsPair(V3871)

if True == tmp356 {
tmp354 := PrimTail(V3871)

__e.TailApply(PrimFunc(symassoc), V3870, tmp354)
return


} else {
__e.Return(PrimSimpleError(MakeString("attempt to search a non-list with assoc\n")))
return
}


}


}


}, 2)

tmp370 := Call(__e, ns2_1set, symassoc, tmp353)


_ = tmp370

tmp371 := MakeNative(func(__e *ControlFlow) {
V3875 := __e.Get(1)
_ = V3875
V3876 := __e.Get(2)
_ = V3876
V3877 := __e.Get(3)
_ = V3877
tmp394 := PrimEqual(Nil, V3877)

if True == tmp394 {
tmp372 := PrimCons(V3875, V3876)

__e.Return(PrimCons(tmp372, Nil))
return


} else {
tmp392 := PrimIsPair(V3877)

var ifres383 Obj

if True == tmp392 {
tmp390 := PrimHead(V3877)

tmp391 := PrimIsPair(tmp390)

var ifres385 Obj

if True == tmp391 {
tmp387 := PrimHead(V3877)

tmp388 := PrimHead(tmp387)

tmp389 := PrimEqual(V3875, tmp388)

var ifres386 Obj

if True == tmp389 {
ifres386 = True


} else {
ifres386 = False


}

ifres385 = ifres386


} else {
ifres385 = False


}

var ifres384 Obj

if True == ifres385 {
ifres384 = True


} else {
ifres384 = False


}

ifres383 = ifres384


} else {
ifres383 = False


}

if True == ifres383 {
tmp373 := PrimHead(V3877)

tmp374 := PrimHead(tmp373)

tmp375 := PrimCons(tmp374, V3876)

tmp376 := PrimTail(V3877)

__e.Return(PrimCons(tmp375, tmp376))
return


} else {
tmp381 := PrimIsPair(V3877)

if True == tmp381 {
tmp377 := PrimHead(V3877)

tmp378 := PrimTail(V3877)

tmp379 := Call(__e, PrimFunc(symshen_4assoc_1set), V3875, V3876, tmp378)


__e.Return(PrimCons(tmp377, tmp379))
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4assoc_1set)
return
}


}


}


}, 3)

tmp395 := Call(__e, ns2_1set, symshen_4assoc_1set, tmp371)


_ = tmp395

tmp396 := MakeNative(func(__e *ControlFlow) {
V3881 := __e.Get(1)
_ = V3881
V3882 := __e.Get(2)
_ = V3882
tmp414 := PrimEqual(Nil, V3882)

if True == tmp414 {
__e.Return(Nil)
return
} else {
tmp412 := PrimIsPair(V3882)

var ifres403 Obj

if True == tmp412 {
tmp410 := PrimHead(V3882)

tmp411 := PrimIsPair(tmp410)

var ifres405 Obj

if True == tmp411 {
tmp407 := PrimHead(V3882)

tmp408 := PrimHead(tmp407)

tmp409 := PrimEqual(V3881, tmp408)

var ifres406 Obj

if True == tmp409 {
ifres406 = True


} else {
ifres406 = False


}

ifres405 = ifres406


} else {
ifres405 = False


}

var ifres404 Obj

if True == ifres405 {
ifres404 = True


} else {
ifres404 = False


}

ifres403 = ifres404


} else {
ifres403 = False


}

if True == ifres403 {
__e.Return(PrimTail(V3882))
return
} else {
tmp401 := PrimIsPair(V3882)

if True == tmp401 {
tmp397 := PrimHead(V3882)

tmp398 := PrimTail(V3882)

tmp399 := Call(__e, PrimFunc(symshen_4assoc_1rm), V3881, tmp398)


__e.Return(PrimCons(tmp397, tmp399))
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4assoc_1rm)
return
}


}


}


}, 2)

tmp415 := Call(__e, ns2_1set, symshen_4assoc_1rm, tmp396)


_ = tmp415

tmp416 := MakeNative(func(__e *ControlFlow) {
V3885 := __e.Get(1)
_ = V3885
tmp420 := PrimEqual(True, V3885)

if True == tmp420 {
__e.Return(True)
return
} else {
tmp418 := PrimEqual(False, V3885)

if True == tmp418 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp421 := Call(__e, ns2_1set, symboolean_2, tmp416)


_ = tmp421

tmp422 := MakeNative(func(__e *ControlFlow) {
V3886 := __e.Get(1)
_ = V3886
tmp427 := PrimEqual(MakeNumber(0), V3886)

if True == tmp427 {
__e.Return(MakeNumber(0))
return
} else {
tmp423 := Call(__e, PrimFunc(symstoutput))


tmp424 := Call(__e, PrimFunc(sympr), MakeString("\n"), tmp423)


_ = tmp424

tmp425 := PrimNumberSubtract(V3886, MakeNumber(1))

__e.TailApply(PrimFunc(symnl), tmp425)
return


}


}, 1)

tmp428 := Call(__e, ns2_1set, symnl, tmp422)


_ = tmp428

tmp429 := MakeNative(func(__e *ControlFlow) {
V3893 := __e.Get(1)
_ = V3893
V3894 := __e.Get(2)
_ = V3894
tmp440 := PrimEqual(Nil, V3893)

if True == tmp440 {
__e.Return(Nil)
return
} else {
tmp438 := PrimIsPair(V3893)

if True == tmp438 {
tmp435 := PrimHead(V3893)

tmp436 := Call(__e, PrimFunc(symelement_2), tmp435, V3894)


if True == tmp436 {
tmp430 := PrimTail(V3893)

__e.TailApply(PrimFunc(symdifference), tmp430, V3894)
return


} else {
tmp431 := PrimHead(V3893)

tmp432 := PrimTail(V3893)

tmp433 := Call(__e, PrimFunc(symdifference), tmp432, V3894)


__e.Return(PrimCons(tmp431, tmp433))
return


}


} else {
__e.Return(PrimSimpleError(MakeString("attempt to find the difference with a non-list\n")))
return
}


}


}, 2)

tmp441 := Call(__e, ns2_1set, symdifference, tmp429)


_ = tmp441

tmp442 := MakeNative(func(__e *ControlFlow) {
V3895 := __e.Get(1)
_ = V3895
V3896 := __e.Get(2)
_ = V3896
__e.Return(V3896)
return
}, 2)

tmp443 := Call(__e, ns2_1set, symdo, tmp442)


_ = tmp443

tmp444 := MakeNative(func(__e *ControlFlow) {
V3908 := __e.Get(1)
_ = V3908
V3909 := __e.Get(2)
_ = V3909
tmp455 := PrimEqual(Nil, V3909)

if True == tmp455 {
__e.Return(False)
return
} else {
tmp453 := PrimIsPair(V3909)

var ifres449 Obj

if True == tmp453 {
tmp451 := PrimHead(V3909)

tmp452 := PrimEqual(V3908, tmp451)

var ifres450 Obj

if True == tmp452 {
ifres450 = True


} else {
ifres450 = False


}

ifres449 = ifres450


} else {
ifres449 = False


}

if True == ifres449 {
__e.Return(True)
return
} else {
tmp447 := PrimIsPair(V3909)

if True == tmp447 {
tmp445 := PrimTail(V3909)

__e.TailApply(PrimFunc(symelement_2), V3908, tmp445)
return


} else {
__e.Return(PrimSimpleError(MakeString("attempt to find an element in a non-list\n")))
return
}


}


}


}, 2)

tmp456 := Call(__e, ns2_1set, symelement_2, tmp444)


_ = tmp456

tmp457 := MakeNative(func(__e *ControlFlow) {
V3912 := __e.Get(1)
_ = V3912
tmp459 := PrimEqual(Nil, V3912)

if True == tmp459 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp460 := Call(__e, ns2_1set, symempty_2, tmp457)


_ = tmp460

tmp461 := MakeNative(func(__e *ControlFlow) {
V3913 := __e.Get(1)
_ = V3913
V3914 := __e.Get(2)
_ = V3914
tmp462 := Call(__e, V3913, V3914)


__e.TailApply(PrimFunc(symshen_4fix_1help), V3913, V3914, tmp462)
return


}, 2)

tmp463 := Call(__e, ns2_1set, symfix, tmp461)


_ = tmp463

tmp464 := MakeNative(func(__e *ControlFlow) {
V3920 := __e.Get(1)
_ = V3920
V3921 := __e.Get(2)
_ = V3921
V3922 := __e.Get(3)
_ = V3922
tmp467 := PrimEqual(V3921, V3922)

if True == tmp467 {
__e.Return(V3922)
return
} else {
tmp465 := Call(__e, V3920, V3922)


__e.TailApply(PrimFunc(symshen_4fix_1help), V3920, V3922, tmp465)
return


}


}, 3)

tmp468 := Call(__e, ns2_1set, symshen_4fix_1help, tmp464)


_ = tmp468

tmp469 := MakeNative(func(__e *ControlFlow) {
V3923 := __e.Get(1)
_ = V3923
V3924 := __e.Get(2)
_ = V3924
V3925 := __e.Get(3)
_ = V3925
V3926 := __e.Get(4)
_ = V3926
tmp470 := MakeNative(func(__e *ControlFlow) {
W3927 := __e.Get(1)
_ = W3927
tmp471 := MakeNative(func(__e *ControlFlow) {
W3929 := __e.Get(1)
_ = W3929
tmp472 := MakeNative(func(__e *ControlFlow) {
W3930 := __e.Get(1)
_ = W3930
__e.Return(V3925)
return
}, 1)

tmp473 := Call(__e, PrimFunc(symshen_4dict_1_6), V3926, V3923, W3929)


__e.TailApply(tmp472, tmp473)
return


}, 1)

tmp474 := Call(__e, PrimFunc(symshen_4assoc_1set), V3924, V3925, W3927)


__e.TailApply(tmp471, tmp474)
return


}, 1)

tmp475 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4_5_1dict), V3926, V3923)
return
}, 0)

tmp476 := MakeNative(func(__e *ControlFlow) {
Z3928 := __e.Get(1)
_ = Z3928
__e.Return(Nil)
return
}, 1)

tmp477 := Call(__e, try_1catch, tmp475, tmp476)


__e.TailApply(tmp470, tmp477)
return


}, 4)

tmp478 := Call(__e, ns2_1set, symput, tmp469)


_ = tmp478

tmp479 := MakeNative(func(__e *ControlFlow) {
V3931 := __e.Get(1)
_ = V3931
V3932 := __e.Get(2)
_ = V3932
V3933 := __e.Get(3)
_ = V3933
tmp480 := MakeNative(func(__e *ControlFlow) {
W3934 := __e.Get(1)
_ = W3934
tmp481 := MakeNative(func(__e *ControlFlow) {
W3936 := __e.Get(1)
_ = W3936
tmp482 := MakeNative(func(__e *ControlFlow) {
W3937 := __e.Get(1)
_ = W3937
__e.Return(V3931)
return
}, 1)

tmp483 := Call(__e, PrimFunc(symshen_4dict_1_6), V3933, V3931, W3936)


__e.TailApply(tmp482, tmp483)
return


}, 1)

tmp484 := Call(__e, PrimFunc(symshen_4assoc_1rm), V3932, W3934)


__e.TailApply(tmp481, tmp484)
return


}, 1)

tmp485 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4_5_1dict), V3933, V3931)
return
}, 0)

tmp486 := MakeNative(func(__e *ControlFlow) {
Z3935 := __e.Get(1)
_ = Z3935
__e.Return(Nil)
return
}, 1)

tmp487 := Call(__e, try_1catch, tmp485, tmp486)


__e.TailApply(tmp480, tmp487)
return


}, 3)

tmp488 := Call(__e, ns2_1set, symunput, tmp479)


_ = tmp488

tmp489 := MakeNative(func(__e *ControlFlow) {
V3938 := __e.Get(1)
_ = V3938
V3939 := __e.Get(2)
_ = V3939
V3940 := __e.Get(3)
_ = V3940
tmp490 := MakeNative(func(__e *ControlFlow) {
W3941 := __e.Get(1)
_ = W3941
tmp491 := MakeNative(func(__e *ControlFlow) {
W3943 := __e.Get(1)
_ = W3943
tmp497 := Call(__e, PrimFunc(symempty_2), W3943)


if True == tmp497 {
tmp492 := Call(__e, PrimFunc(symshen_4app), V3938, MakeString("\n"), symshen_4s)


tmp493 := PrimStringConcat(MakeString(" not found for "), tmp492)

tmp494 := Call(__e, PrimFunc(symshen_4app), V3939, tmp493, symshen_4s)


tmp495 := PrimStringConcat(MakeString("attribute "), tmp494)

__e.Return(PrimSimpleError(tmp495))
return


} else {
__e.Return(PrimTail(W3943))
return
}


}, 1)

tmp498 := Call(__e, PrimFunc(symassoc), V3939, W3941)


__e.TailApply(tmp491, tmp498)
return


}, 1)

tmp499 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4_5_1dict), V3940, V3938)
return
}, 0)

tmp500 := MakeNative(func(__e *ControlFlow) {
Z3942 := __e.Get(1)
_ = Z3942
tmp501 := Call(__e, PrimFunc(symshen_4app), V3939, MakeString("\n"), symshen_4s)


tmp502 := PrimStringConcat(MakeString(" has no attributes: "), tmp501)

tmp503 := Call(__e, PrimFunc(symshen_4app), V3938, tmp502, symshen_4a)


__e.Return(PrimSimpleError(tmp503))
return


}, 1)

tmp504 := Call(__e, try_1catch, tmp499, tmp500)


__e.TailApply(tmp490, tmp504)
return


}, 3)

tmp505 := Call(__e, ns2_1set, symget, tmp489)


_ = tmp505

tmp506 := MakeNative(func(__e *ControlFlow) {
V3944 := __e.Get(1)
_ = V3944
V3945 := __e.Get(2)
_ = V3945
tmp507 := MakeNative(func(__e *ControlFlow) {
W3946 := __e.Get(1)
_ = W3946
tmp509 := PrimEqual(W3946, MakeNumber(0))

if True == tmp509 {
__e.Return(MakeNumber(1))
return
} else {
__e.Return(W3946)
return
}


}, 1)

tmp510 := Call(__e, PrimFunc(symshen_4hashkey), V3944)


tmp511 := Call(__e, PrimFunc(symshen_4mod), tmp510, V3945)


__e.TailApply(tmp507, tmp511)
return


}, 2)

tmp512 := Call(__e, ns2_1set, symhash, tmp506)


_ = tmp512

tmp513 := MakeNative(func(__e *ControlFlow) {
V3947 := __e.Get(1)
_ = V3947
tmp514 := MakeNative(func(__e *ControlFlow) {
W3948 := __e.Get(1)
_ = W3948
__e.TailApply(PrimFunc(symshen_4prodbutzero), W3948, MakeNumber(1))
return
}, 1)

tmp515 := MakeNative(func(__e *ControlFlow) {
Z3949 := __e.Get(1)
_ = Z3949
__e.Return(PrimStringToNumber(Z3949))
return
}, 1)

tmp516 := Call(__e, PrimFunc(symexplode), V3947)


tmp517 := Call(__e, PrimFunc(symmap), tmp515, tmp516)


__e.TailApply(tmp514, tmp517)
return


}, 1)

tmp518 := Call(__e, ns2_1set, symshen_4hashkey, tmp513)


_ = tmp518

tmp519 := MakeNative(func(__e *ControlFlow) {
V3950 := __e.Get(1)
_ = V3950
V3951 := __e.Get(2)
_ = V3951
tmp538 := PrimEqual(Nil, V3950)

if True == tmp538 {
__e.Return(V3951)
return
} else {
tmp536 := PrimIsPair(V3950)

var ifres532 Obj

if True == tmp536 {
tmp534 := PrimHead(V3950)

tmp535 := PrimEqual(MakeNumber(0), tmp534)

var ifres533 Obj

if True == tmp535 {
ifres533 = True


} else {
ifres533 = False


}

ifres532 = ifres533


} else {
ifres532 = False


}

if True == ifres532 {
tmp520 := PrimTail(V3950)

__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp520, V3951)
return


} else {
tmp530 := PrimIsPair(V3950)

if True == tmp530 {
tmp528 := PrimGreatThan(V3951, MakeNumber(10000000000))

if True == tmp528 {
tmp521 := PrimTail(V3950)

tmp522 := PrimHead(V3950)

tmp523 := PrimNumberAdd(V3951, tmp522)

__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp521, tmp523)
return


} else {
tmp524 := PrimTail(V3950)

tmp525 := PrimHead(V3950)

tmp526 := PrimNumberMultiply(V3951, tmp525)

__e.TailApply(PrimFunc(symshen_4prodbutzero), tmp524, tmp526)
return


}


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4prodbutzero)
return
}


}


}


}, 2)

tmp539 := Call(__e, ns2_1set, symshen_4prodbutzero, tmp519)


_ = tmp539

tmp540 := MakeNative(func(__e *ControlFlow) {
V3952 := __e.Get(1)
_ = V3952
V3953 := __e.Get(2)
_ = V3953
tmp541 := PrimCons(V3953, Nil)

tmp542 := Call(__e, PrimFunc(symshen_4multiples), V3952, tmp541)


__e.TailApply(PrimFunc(symshen_4modh), V3952, tmp542)
return


}, 2)

tmp543 := Call(__e, ns2_1set, symshen_4mod, tmp540)


_ = tmp543

tmp544 := MakeNative(func(__e *ControlFlow) {
V3958 := __e.Get(1)
_ = V3958
V3959 := __e.Get(2)
_ = V3959
tmp555 := PrimIsPair(V3959)

var ifres551 Obj

if True == tmp555 {
tmp553 := PrimHead(V3959)

tmp554 := PrimGreatThan(tmp553, V3958)

var ifres552 Obj

if True == tmp554 {
ifres552 = True


} else {
ifres552 = False


}

ifres551 = ifres552


} else {
ifres551 = False


}

if True == ifres551 {
__e.Return(PrimTail(V3959))
return
} else {
tmp549 := PrimIsPair(V3959)

if True == tmp549 {
tmp545 := PrimHead(V3959)

tmp546 := PrimNumberMultiply(MakeNumber(2), tmp545)

tmp547 := PrimCons(tmp546, V3959)

__e.TailApply(PrimFunc(symshen_4multiples), V3958, tmp547)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.multiples")))
return
}


}


}, 2)

tmp556 := Call(__e, ns2_1set, symshen_4multiples, tmp544)


_ = tmp556

tmp557 := MakeNative(func(__e *ControlFlow) {
V3966 := __e.Get(1)
_ = V3966
V3967 := __e.Get(2)
_ = V3967
tmp575 := PrimEqual(MakeNumber(0), V3966)

if True == tmp575 {
__e.Return(MakeNumber(0))
return
} else {
tmp573 := PrimEqual(Nil, V3967)

if True == tmp573 {
__e.Return(V3966)
return
} else {
tmp571 := PrimIsPair(V3967)

var ifres567 Obj

if True == tmp571 {
tmp569 := PrimHead(V3967)

tmp570 := PrimGreatThan(tmp569, V3966)

var ifres568 Obj

if True == tmp570 {
ifres568 = True


} else {
ifres568 = False


}

ifres567 = ifres568


} else {
ifres567 = False


}

if True == ifres567 {
tmp560 := PrimTail(V3967)

tmp561 := Call(__e, PrimFunc(symempty_2), tmp560)


if True == tmp561 {
__e.Return(V3966)
return
} else {
tmp558 := PrimTail(V3967)

__e.TailApply(PrimFunc(symshen_4modh), V3966, tmp558)
return


}


} else {
tmp565 := PrimIsPair(V3967)

if True == tmp565 {
tmp562 := PrimHead(V3967)

tmp563 := PrimNumberSubtract(V3966, tmp562)

__e.TailApply(PrimFunc(symshen_4modh), tmp563, V3967)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.modh")))
return
}


}


}


}


}, 2)

tmp576 := Call(__e, ns2_1set, symshen_4modh, tmp557)


_ = tmp576

tmp577 := MakeNative(func(__e *ControlFlow) {
V3970 := __e.Get(1)
_ = V3970
tmp584 := PrimEqual(Nil, V3970)

if True == tmp584 {
__e.Return(MakeNumber(0))
return
} else {
tmp582 := PrimIsPair(V3970)

if True == tmp582 {
tmp578 := PrimHead(V3970)

tmp579 := PrimTail(V3970)

tmp580 := Call(__e, PrimFunc(symsum), tmp579)


__e.Return(PrimNumberAdd(tmp578, tmp580))
return


} else {
__e.Return(PrimSimpleError(MakeString("attempt to sum a non-list\n")))
return
}


}


}, 1)

tmp585 := Call(__e, ns2_1set, symsum, tmp577)


_ = tmp585

tmp586 := MakeNative(func(__e *ControlFlow) {
V3975 := __e.Get(1)
_ = V3975
tmp588 := PrimIsPair(V3975)

if True == tmp588 {
__e.Return(PrimHead(V3975))
return
} else {
__e.Return(PrimSimpleError(MakeString("head expects a non-empty list\n")))
return
}


}, 1)

tmp589 := Call(__e, ns2_1set, symhead, tmp586)


_ = tmp589

tmp590 := MakeNative(func(__e *ControlFlow) {
V3980 := __e.Get(1)
_ = V3980
tmp592 := PrimIsPair(V3980)

if True == tmp592 {
__e.Return(PrimTail(V3980))
return
} else {
__e.Return(PrimSimpleError(MakeString("tail expects a non-empty list\n")))
return
}


}, 1)

tmp593 := Call(__e, ns2_1set, symtail, tmp590)


_ = tmp593

tmp594 := MakeNative(func(__e *ControlFlow) {
V3981 := __e.Get(1)
_ = V3981
__e.Return(PrimPos(V3981, MakeNumber(0)))
return
}, 1)

tmp595 := Call(__e, ns2_1set, symhdstr, tmp594)


_ = tmp595

tmp596 := MakeNative(func(__e *ControlFlow) {
V3988 := __e.Get(1)
_ = V3988
V3989 := __e.Get(2)
_ = V3989
tmp607 := PrimEqual(Nil, V3988)

if True == tmp607 {
__e.Return(Nil)
return
} else {
tmp605 := PrimIsPair(V3988)

if True == tmp605 {
tmp602 := PrimHead(V3988)

tmp603 := Call(__e, PrimFunc(symelement_2), tmp602, V3989)


if True == tmp603 {
tmp597 := PrimHead(V3988)

tmp598 := PrimTail(V3988)

tmp599 := Call(__e, PrimFunc(symintersection), tmp598, V3989)


__e.Return(PrimCons(tmp597, tmp599))
return


} else {
tmp600 := PrimTail(V3988)

__e.TailApply(PrimFunc(symintersection), tmp600, V3989)
return


}


} else {
__e.Return(PrimSimpleError(MakeString("attempt to find the intersection with a non-list\n")))
return
}


}


}, 2)

tmp608 := Call(__e, ns2_1set, symintersection, tmp596)


_ = tmp608

tmp609 := MakeNative(func(__e *ControlFlow) {
V3990 := __e.Get(1)
_ = V3990
__e.TailApply(PrimFunc(symshen_4reverse_1help), V3990, Nil)
return
}, 1)

tmp610 := Call(__e, ns2_1set, symreverse, tmp609)


_ = tmp610

tmp611 := MakeNative(func(__e *ControlFlow) {
V3995 := __e.Get(1)
_ = V3995
V3996 := __e.Get(2)
_ = V3996
tmp618 := PrimEqual(Nil, V3995)

if True == tmp618 {
__e.Return(V3996)
return
} else {
tmp616 := PrimIsPair(V3995)

if True == tmp616 {
tmp612 := PrimTail(V3995)

tmp613 := PrimHead(V3995)

tmp614 := PrimCons(tmp613, V3996)

__e.TailApply(PrimFunc(symshen_4reverse_1help), tmp612, tmp614)
return


} else {
__e.Return(PrimSimpleError(MakeString("attempt to reverse a non-list\n")))
return
}


}


}, 2)

tmp619 := Call(__e, ns2_1set, symshen_4reverse_1help, tmp611)


_ = tmp619

tmp620 := MakeNative(func(__e *ControlFlow) {
V4001 := __e.Get(1)
_ = V4001
V4002 := __e.Get(2)
_ = V4002
tmp631 := PrimEqual(Nil, V4001)

if True == tmp631 {
__e.Return(V4002)
return
} else {
tmp629 := PrimIsPair(V4001)

if True == tmp629 {
tmp626 := PrimHead(V4001)

tmp627 := Call(__e, PrimFunc(symelement_2), tmp626, V4002)


if True == tmp627 {
tmp621 := PrimTail(V4001)

__e.TailApply(PrimFunc(symunion), tmp621, V4002)
return


} else {
tmp622 := PrimHead(V4001)

tmp623 := PrimTail(V4001)

tmp624 := Call(__e, PrimFunc(symunion), tmp623, V4002)


__e.Return(PrimCons(tmp622, tmp624))
return


}


} else {
__e.Return(PrimSimpleError(MakeString("attempt to find the union with a non-list\n")))
return
}


}


}, 2)

tmp632 := Call(__e, ns2_1set, symunion, tmp620)


_ = tmp632

tmp633 := MakeNative(func(__e *ControlFlow) {
V4003 := __e.Get(1)
_ = V4003
tmp634 := MakeNative(func(__e *ControlFlow) {
W4004 := __e.Get(1)
_ = W4004
tmp635 := MakeNative(func(__e *ControlFlow) {
W4005 := __e.Get(1)
_ = W4005
tmp636 := MakeNative(func(__e *ControlFlow) {
W4006 := __e.Get(1)
_ = W4006
tmp642 := PrimEqual(MakeString("y"), W4006)

if True == tmp642 {
__e.Return(True)
return
} else {
tmp640 := PrimEqual(MakeString("n"), W4006)

if True == tmp640 {
__e.Return(False)
return
} else {
tmp637 := Call(__e, PrimFunc(symstoutput))


tmp638 := Call(__e, PrimFunc(sympr), MakeString("please answer y or n\n"), tmp637)


_ = tmp638

__e.TailApply(PrimFunc(symy_1or_1n_2), V4003)
return


}


}


}, 1)

tmp643 := Call(__e, PrimFunc(symstinput))


tmp644 := Call(__e, PrimFunc(symread), tmp643)


tmp645 := Call(__e, PrimFunc(symshen_4app), tmp644, MakeString(""), symshen_4s)


__e.TailApply(tmp636, tmp645)
return


}, 1)

tmp646 := Call(__e, PrimFunc(symstoutput))


tmp647 := Call(__e, PrimFunc(sympr), MakeString(" (y/n) "), tmp646)


__e.TailApply(tmp635, tmp647)
return


}, 1)

tmp648 := Call(__e, PrimFunc(symshen_4proc_1nl), V4003)


tmp649 := Call(__e, PrimFunc(symstoutput))


tmp650 := Call(__e, PrimFunc(sympr), tmp648, tmp649)


__e.TailApply(tmp634, tmp650)
return


}, 1)

tmp651 := Call(__e, ns2_1set, symy_1or_1n_2, tmp633)


_ = tmp651

tmp652 := MakeNative(func(__e *ControlFlow) {
V4007 := __e.Get(1)
_ = V4007
if True == V4007 {
__e.Return(False)
return
} else {
__e.Return(True)
return
}
}, 1)

tmp654 := Call(__e, ns2_1set, symnot, tmp652)


_ = tmp654

tmp655 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimSimpleError(MakeString("")))
return
}, 0)

tmp656 := Call(__e, ns2_1set, symabort, tmp655)


_ = tmp656

tmp657 := MakeNative(func(__e *ControlFlow) {
V4013 := __e.Get(1)
_ = V4013
V4014 := __e.Get(2)
_ = V4014
V4015 := __e.Get(3)
_ = V4015
tmp665 := PrimEqual(V4014, V4015)

if True == tmp665 {
__e.Return(V4013)
return
} else {
tmp663 := PrimIsPair(V4015)

if True == tmp663 {
tmp658 := PrimHead(V4015)

tmp659 := Call(__e, PrimFunc(symsubst), V4013, V4014, tmp658)


tmp660 := PrimTail(V4015)

tmp661 := Call(__e, PrimFunc(symsubst), V4013, V4014, tmp660)


__e.Return(PrimCons(tmp659, tmp661))
return


} else {
__e.Return(V4015)
return
}


}


}, 3)

tmp666 := Call(__e, ns2_1set, symsubst, tmp657)


_ = tmp666

tmp667 := MakeNative(func(__e *ControlFlow) {
V4016 := __e.Get(1)
_ = V4016
tmp668 := Call(__e, PrimFunc(symshen_4app), V4016, MakeString(""), symshen_4a)


__e.TailApply(PrimFunc(symshen_4explode_1h), tmp668)
return


}, 1)

tmp669 := Call(__e, ns2_1set, symexplode, tmp667)


_ = tmp669

tmp670 := MakeNative(func(__e *ControlFlow) {
V4019 := __e.Get(1)
_ = V4019
tmp677 := PrimEqual(MakeString(""), V4019)

if True == tmp677 {
__e.Return(Nil)
return
} else {
tmp675 := Call(__e, PrimFunc(symshen_4_7string_2), V4019)


if True == tmp675 {
tmp671 := Call(__e, PrimFunc(symhdstr), V4019)


tmp672 := PrimTailString(V4019)

tmp673 := Call(__e, PrimFunc(symshen_4explode_1h), tmp672)


__e.Return(PrimCons(tmp671, tmp673))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in explode-h")))
return
}


}


}, 1)

tmp678 := Call(__e, ns2_1set, symshen_4explode_1h, tmp670)


_ = tmp678

tmp679 := MakeNative(func(__e *ControlFlow) {
V4020 := __e.Get(1)
_ = V4020
tmp682 := PrimEqual(V4020, MakeString(""))

var ifres680 Obj

if True == tmp682 {
ifres680 = MakeString("")


} else {
tmp681 := Call(__e, PrimFunc(symshen_4app), V4020, MakeString("/"), symshen_4a)


ifres680 = tmp681


}

__e.Return(PrimSet(sym_dhome_1directory_d, ifres680))
return


}, 1)

tmp683 := Call(__e, ns2_1set, symcd, tmp679)


_ = tmp683

tmp684 := MakeNative(func(__e *ControlFlow) {
V4021 := __e.Get(1)
_ = V4021
V4022 := __e.Get(2)
_ = V4022
tmp692 := PrimEqual(Nil, V4022)

if True == tmp692 {
__e.Return(True)
return
} else {
tmp690 := PrimIsPair(V4022)

if True == tmp690 {
tmp685 := MakeNative(func(__e *ControlFlow) {
W4023 := __e.Get(1)
_ = W4023
tmp686 := PrimTail(V4022)

__e.TailApply(PrimFunc(symshen_4for_1each), V4021, tmp686)
return


}, 1)

tmp687 := PrimHead(V4022)

tmp688 := Call(__e, V4021, tmp687)


__e.TailApply(tmp685, tmp688)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4for_1each)
return
}


}


}, 2)

tmp693 := Call(__e, ns2_1set, symshen_4for_1each, tmp684)


_ = tmp693

tmp694 := MakeNative(func(__e *ControlFlow) {
V4024 := __e.Get(1)
_ = V4024
V4025 := __e.Get(2)
_ = V4025
__e.TailApply(PrimFunc(symshen_4map_1h), V4024, V4025, Nil)
return
}, 2)

tmp695 := Call(__e, ns2_1set, symmap, tmp694)


_ = tmp695

tmp696 := MakeNative(func(__e *ControlFlow) {
V4026 := __e.Get(1)
_ = V4026
V4027 := __e.Get(2)
_ = V4027
V4028 := __e.Get(3)
_ = V4028
tmp704 := PrimEqual(Nil, V4027)

if True == tmp704 {
__e.TailApply(PrimFunc(symreverse), V4028)
return
} else {
tmp702 := PrimIsPair(V4027)

if True == tmp702 {
tmp697 := PrimTail(V4027)

tmp698 := PrimHead(V4027)

tmp699 := Call(__e, V4026, tmp698)


tmp700 := PrimCons(tmp699, V4028)

__e.TailApply(PrimFunc(symshen_4map_1h), V4026, tmp697, tmp700)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4map_1h)
return
}


}


}, 3)

tmp705 := Call(__e, ns2_1set, symshen_4map_1h, tmp696)


_ = tmp705

tmp706 := MakeNative(func(__e *ControlFlow) {
V4029 := __e.Get(1)
_ = V4029
__e.TailApply(PrimFunc(symshen_4length_1h), V4029, MakeNumber(0))
return
}, 1)

tmp707 := Call(__e, ns2_1set, symlength, tmp706)


_ = tmp707

tmp708 := MakeNative(func(__e *ControlFlow) {
V4034 := __e.Get(1)
_ = V4034
V4035 := __e.Get(2)
_ = V4035
tmp712 := PrimEqual(Nil, V4034)

if True == tmp712 {
__e.Return(V4035)
return
} else {
tmp709 := PrimTail(V4034)

tmp710 := PrimNumberAdd(V4035, MakeNumber(1))

__e.TailApply(PrimFunc(symshen_4length_1h), tmp709, tmp710)
return


}


}, 2)

tmp713 := Call(__e, ns2_1set, symshen_4length_1h, tmp708)


_ = tmp713

tmp714 := MakeNative(func(__e *ControlFlow) {
V4041 := __e.Get(1)
_ = V4041
V4042 := __e.Get(2)
_ = V4042
tmp722 := PrimEqual(V4041, V4042)

if True == tmp722 {
__e.Return(MakeNumber(1))
return
} else {
tmp720 := PrimIsPair(V4042)

if True == tmp720 {
tmp715 := PrimHead(V4042)

tmp716 := Call(__e, PrimFunc(symoccurrences), V4041, tmp715)


tmp717 := PrimTail(V4042)

tmp718 := Call(__e, PrimFunc(symoccurrences), V4041, tmp717)


__e.Return(PrimNumberAdd(tmp716, tmp718))
return


} else {
__e.Return(MakeNumber(0))
return
}


}


}, 2)

tmp723 := Call(__e, ns2_1set, symoccurrences, tmp714)


_ = tmp723

tmp724 := MakeNative(func(__e *ControlFlow) {
V4047 := __e.Get(1)
_ = V4047
V4048 := __e.Get(2)
_ = V4048
tmp737 := PrimEqual(MakeNumber(1), V4047)

var ifres734 Obj

if True == tmp737 {
tmp736 := PrimIsPair(V4048)

var ifres735 Obj

if True == tmp736 {
ifres735 = True


} else {
ifres735 = False


}

ifres734 = ifres735


} else {
ifres734 = False


}

if True == ifres734 {
__e.Return(PrimHead(V4048))
return
} else {
tmp732 := PrimIsPair(V4048)

if True == tmp732 {
tmp725 := PrimNumberSubtract(V4047, MakeNumber(1))

tmp726 := PrimTail(V4048)

__e.TailApply(PrimFunc(symnth), tmp725, tmp726)
return


} else {
tmp727 := Call(__e, PrimFunc(symshen_4app), V4048, MakeString("\n"), symshen_4a)


tmp728 := PrimStringConcat(MakeString(", "), tmp727)

tmp729 := Call(__e, PrimFunc(symshen_4app), V4047, tmp728, symshen_4a)


tmp730 := PrimStringConcat(MakeString("nth applied to "), tmp729)

__e.Return(PrimSimpleError(tmp730))
return


}


}


}, 2)

tmp738 := Call(__e, ns2_1set, symnth, tmp724)


_ = tmp738

tmp739 := MakeNative(func(__e *ControlFlow) {
V4049 := __e.Get(1)
_ = V4049
tmp746 := PrimIsNumber(V4049)

if True == tmp746 {
tmp741 := MakeNative(func(__e *ControlFlow) {
W4050 := __e.Get(1)
_ = W4050
tmp742 := Call(__e, PrimFunc(symshen_4magless), W4050, MakeNumber(1))


__e.TailApply(PrimFunc(symshen_4integer_1test_2), W4050, tmp742)
return


}, 1)

tmp743 := Call(__e, PrimFunc(symshen_4abs), V4049)


tmp744 := Call(__e, tmp741, tmp743)


if True == tmp744 {
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

tmp747 := Call(__e, ns2_1set, syminteger_2, tmp739)


_ = tmp747

tmp748 := MakeNative(func(__e *ControlFlow) {
V4051 := __e.Get(1)
_ = V4051
tmp750 := PrimGreatThan(V4051, MakeNumber(0))

if True == tmp750 {
__e.Return(V4051)
return
} else {
__e.Return(PrimNumberSubtract(MakeNumber(0), V4051))
return
}


}, 1)

tmp751 := Call(__e, ns2_1set, symshen_4abs, tmp748)


_ = tmp751

tmp752 := MakeNative(func(__e *ControlFlow) {
V4052 := __e.Get(1)
_ = V4052
V4053 := __e.Get(2)
_ = V4053
tmp753 := MakeNative(func(__e *ControlFlow) {
W4054 := __e.Get(1)
_ = W4054
tmp755 := PrimGreatThan(W4054, V4052)

if True == tmp755 {
__e.Return(V4053)
return
} else {
__e.TailApply(PrimFunc(symshen_4magless), V4052, W4054)
return
}


}, 1)

tmp756 := PrimNumberMultiply(V4053, MakeNumber(2))

__e.TailApply(tmp753, tmp756)
return


}, 2)

tmp757 := Call(__e, ns2_1set, symshen_4magless, tmp752)


_ = tmp757

tmp758 := MakeNative(func(__e *ControlFlow) {
V4058 := __e.Get(1)
_ = V4058
V4059 := __e.Get(2)
_ = V4059
tmp766 := PrimEqual(MakeNumber(0), V4058)

if True == tmp766 {
__e.Return(True)
return
} else {
tmp764 := PrimGreatThan(MakeNumber(1), V4058)

if True == tmp764 {
__e.Return(False)
return
} else {
tmp759 := MakeNative(func(__e *ControlFlow) {
W4060 := __e.Get(1)
_ = W4060
tmp761 := PrimGreatThan(MakeNumber(0), W4060)

if True == tmp761 {
__e.Return(PrimIsInteger(V4058))
return
} else {
__e.TailApply(PrimFunc(symshen_4integer_1test_2), W4060, V4059)
return
}


}, 1)

tmp762 := PrimNumberSubtract(V4058, V4059)

__e.TailApply(tmp759, tmp762)
return


}


}


}, 2)

tmp767 := Call(__e, ns2_1set, symshen_4integer_1test_2, tmp758)


_ = tmp767

tmp768 := MakeNative(func(__e *ControlFlow) {
V4067 := __e.Get(1)
_ = V4067
V4068 := __e.Get(2)
_ = V4068
tmp776 := PrimEqual(Nil, V4068)

if True == tmp776 {
__e.Return(Nil)
return
} else {
tmp774 := PrimIsPair(V4068)

if True == tmp774 {
tmp769 := PrimHead(V4068)

tmp770 := Call(__e, V4067, tmp769)


tmp771 := PrimTail(V4068)

tmp772 := Call(__e, PrimFunc(symmapcan), V4067, tmp771)


__e.TailApply(PrimFunc(symappend), tmp770, tmp772)
return


} else {
__e.Return(PrimSimpleError(MakeString("attempt to mapcan over a non-list\n")))
return
}


}


}, 2)

tmp777 := Call(__e, ns2_1set, symmapcan, tmp768)


_ = tmp777

tmp778 := MakeNative(func(__e *ControlFlow) {
V4074 := __e.Get(1)
_ = V4074
V4075 := __e.Get(2)
_ = V4075
tmp780 := PrimEqual(V4074, V4075)

if True == tmp780 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 2)

tmp781 := Call(__e, ns2_1set, sym_a_a, tmp778)


_ = tmp781

tmp782 := MakeNative(func(__e *ControlFlow) {
V4076 := __e.Get(1)
_ = V4076
tmp792 := PrimIsSymbol(V4076)

if True == tmp792 {
tmp784 := MakeNative(func(__e *ControlFlow) {
W4077 := __e.Get(1)
_ = W4077
tmp786 := PrimEqual(W4077, symshen_4this_1symbol_1is_1unbound)

if True == tmp786 {
__e.Return(False)
return
} else {
__e.Return(True)
return
}


}, 1)

tmp787 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(V4076))
return
}, 0)

tmp788 := MakeNative(func(__e *ControlFlow) {
Z4078 := __e.Get(1)
_ = Z4078
__e.Return(symshen_4this_1symbol_1is_1unbound)
return
}, 1)

tmp789 := Call(__e, try_1catch, tmp787, tmp788)


tmp790 := Call(__e, tmp784, tmp789)


if True == tmp790 {
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

tmp793 := Call(__e, ns2_1set, symbound_2, tmp782)


_ = tmp793

tmp794 := MakeNative(func(__e *ControlFlow) {
V4079 := __e.Get(1)
_ = V4079
tmp800 := PrimEqual(MakeString(""), V4079)

if True == tmp800 {
__e.Return(Nil)
return
} else {
tmp795 := PrimPos(V4079, MakeNumber(0))

tmp796 := PrimStringToNumber(tmp795)

tmp797 := PrimTailString(V4079)

tmp798 := Call(__e, PrimFunc(symshen_4string_1_6bytes), tmp797)


__e.Return(PrimCons(tmp796, tmp798))
return


}


}, 1)

tmp801 := Call(__e, ns2_1set, symshen_4string_1_6bytes, tmp794)


_ = tmp801

tmp802 := MakeNative(func(__e *ControlFlow) {
V4080 := __e.Get(1)
_ = V4080
tmp806 := PrimLessThan(V4080, MakeNumber(0))

if True == tmp806 {
__e.Return(PrimValue(symshen_4_dmaxinferences_d))
return
} else {
tmp804 := PrimIsInteger(V4080)

if True == tmp804 {
__e.Return(PrimSet(symshen_4_dmaxinferences_d, V4080))
return
} else {
__e.Return(PrimSimpleError(MakeString("maxinferences expects an integer value\n")))
return
}


}


}, 1)

tmp807 := Call(__e, ns2_1set, symmaxinferences, tmp802)


_ = tmp807

tmp808 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dinfs_d))
return
}, 0)

tmp809 := Call(__e, ns2_1set, syminferences, tmp808)


_ = tmp809

tmp810 := MakeNative(func(__e *ControlFlow) {
V4081 := __e.Get(1)
_ = V4081
__e.Return(V4081)
return
}, 1)

tmp811 := Call(__e, ns2_1set, symprotect, tmp810)


_ = tmp811

tmp812 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dsterror_d))
return
}, 0)

tmp813 := Call(__e, ns2_1set, symsterror, tmp812)


_ = tmp813

tmp814 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dstoutput_d))
return
}, 0)

tmp815 := Call(__e, ns2_1set, symstoutput, tmp814)


_ = tmp815

tmp816 := MakeNative(func(__e *ControlFlow) {
V4082 := __e.Get(1)
_ = V4082
tmp817 := MakeNative(func(__e *ControlFlow) {
W4083 := __e.Get(1)
_ = W4083
tmp821 := PrimIsSymbol(W4083)

if True == tmp821 {
__e.Return(W4083)
return
} else {
tmp818 := Call(__e, PrimFunc(symshen_4app), V4082, MakeString(" to a symbol"), symshen_4s)


tmp819 := PrimStringConcat(MakeString("cannot intern "), tmp818)

__e.Return(PrimSimpleError(tmp819))
return


}


}, 1)

tmp822 := PrimIntern(V4082)

__e.TailApply(tmp817, tmp822)
return


}, 1)

tmp823 := Call(__e, ns2_1set, symstring_1_6symbol, tmp816)


_ = tmp823

tmp824 := MakeNative(func(__e *ControlFlow) {
V4086 := __e.Get(1)
_ = V4086
tmp828 := PrimEqual(sym_7, V4086)

if True == tmp828 {
__e.Return(PrimSet(symshen_4_doptimise_d, True))
return
} else {
tmp826 := PrimEqual(sym_1, V4086)

if True == tmp826 {
__e.Return(PrimSet(symshen_4_doptimise_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("optimise expects a + or a -.\n")))
return
}


}


}, 1)

tmp829 := Call(__e, ns2_1set, symoptimise, tmp824)


_ = tmp829

tmp830 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dos_d))
return
}, 0)

tmp831 := Call(__e, ns2_1set, symos, tmp830)


_ = tmp831

tmp832 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dlanguage_d))
return
}, 0)

tmp833 := Call(__e, ns2_1set, symlanguage, tmp832)


_ = tmp833

tmp834 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dversion_d))
return
}, 0)

tmp835 := Call(__e, ns2_1set, symversion, tmp834)


_ = tmp835

tmp836 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dport_d))
return
}, 0)

tmp837 := Call(__e, ns2_1set, symport, tmp836)


_ = tmp837

tmp838 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dporters_d))
return
}, 0)

tmp839 := Call(__e, ns2_1set, symporters, tmp838)


_ = tmp839

tmp840 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dimplementation_d))
return
}, 0)

tmp841 := Call(__e, ns2_1set, symimplementation, tmp840)


_ = tmp841

tmp842 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_drelease_d))
return
}, 0)

tmp843 := Call(__e, ns2_1set, symrelease, tmp842)


_ = tmp843

tmp844 := MakeNative(func(__e *ControlFlow) {
V4087 := __e.Get(1)
_ = V4087
tmp849 := PrimEqual(symnull, V4087)

if True == tmp849 {
__e.Return(True)
return
} else {
tmp845 := MakeNative(func(__e *ControlFlow) {
tmp846 := Call(__e, PrimFunc(symexternal), V4087)


_ = tmp846

__e.Return(True)
return


}, 0)

tmp847 := MakeNative(func(__e *ControlFlow) {
Z4088 := __e.Get(1)
_ = Z4088
__e.Return(False)
return
}, 1)

__e.TailApply(try_1catch, tmp845, tmp847)
return


}


}, 1)

tmp850 := Call(__e, ns2_1set, sympackage_2, tmp844)


_ = tmp850

tmp851 := MakeNative(func(__e *ControlFlow) {
__e.Return(sym_4_4_4)
return
}, 0)

tmp852 := Call(__e, ns2_1set, symfail, tmp851)


_ = tmp852

tmp853 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_duserdefs_d))
return
}, 0)

tmp854 := Call(__e, ns2_1set, symuserdefs, tmp853)


_ = tmp854

tmp855 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_doptimise_d))
return
}, 0)

tmp856 := Call(__e, ns2_1set, symoptimise_2, tmp855)


_ = tmp856

tmp857 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dhush_d))
return
}, 0)

tmp858 := Call(__e, ns2_1set, symhush_2, tmp857)


_ = tmp858

tmp859 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dshen_1type_1theory_1enabled_2_d))
return
}, 0)

tmp860 := Call(__e, ns2_1set, symsystem_1S_2, tmp859)


_ = tmp860

tmp861 := MakeNative(func(__e *ControlFlow) {
V4091 := __e.Get(1)
_ = V4091
tmp865 := PrimEqual(sym_7, V4091)

if True == tmp865 {
__e.Return(PrimSet(symshen_4_dshen_1type_1theory_1enabled_2_d, True))
return
} else {
tmp863 := PrimEqual(sym_1, V4091)

if True == tmp863 {
__e.Return(PrimSet(symshen_4_dshen_1type_1theory_1enabled_2_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("enable-type-theory expects a + or a -\n")))
return
}


}


}, 1)

tmp866 := Call(__e, ns2_1set, symenable_1type_1theory, tmp861)


_ = tmp866

tmp867 := MakeNative(func(__e *ControlFlow) {
V4094 := __e.Get(1)
_ = V4094
tmp871 := PrimEqual(sym_7, V4094)

if True == tmp871 {
__e.Return(PrimSet(sym_dhush_d, True))
return
} else {
tmp869 := PrimEqual(sym_1, V4094)

if True == tmp869 {
__e.Return(PrimSet(sym_dhush_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("hush expects a + or a -\n")))
return
}


}


}, 1)

tmp872 := Call(__e, ns2_1set, symhush, tmp867)


_ = tmp872

tmp873 := MakeNative(func(__e *ControlFlow) {
V4097 := __e.Get(1)
_ = V4097
tmp877 := PrimEqual(sym_7, V4097)

if True == tmp877 {
__e.Return(PrimSet(symshen_4_dtc_d, True))
return
} else {
tmp875 := PrimEqual(sym_1, V4097)

if True == tmp875 {
__e.Return(PrimSet(symshen_4_dtc_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("tc expects a + or -")))
return
}


}


}, 1)

tmp878 := Call(__e, ns2_1set, symtc, tmp873)


_ = tmp878

tmp879 := MakeNative(func(__e *ControlFlow) {
V4098 := __e.Get(1)
_ = V4098
tmp880 := PrimValue(symshen_4_dsigf_d)

tmp881 := Call(__e, PrimFunc(symshen_4unassoc), V4098, tmp880)


tmp882 := PrimSet(symshen_4_dsigf_d, tmp881)

_ = tmp882

__e.Return(V4098)
return


}, 1)

tmp883 := Call(__e, ns2_1set, symdestroy, tmp879)


_ = tmp883

tmp884 := MakeNative(func(__e *ControlFlow) {
V4108 := __e.Get(1)
_ = V4108
V4109 := __e.Get(2)
_ = V4109
tmp902 := PrimEqual(Nil, V4109)

if True == tmp902 {
__e.Return(Nil)
return
} else {
tmp900 := PrimIsPair(V4109)

var ifres891 Obj

if True == tmp900 {
tmp898 := PrimHead(V4109)

tmp899 := PrimIsPair(tmp898)

var ifres893 Obj

if True == tmp899 {
tmp895 := PrimHead(V4109)

tmp896 := PrimHead(tmp895)

tmp897 := PrimEqual(V4108, tmp896)

var ifres894 Obj

if True == tmp897 {
ifres894 = True


} else {
ifres894 = False


}

ifres893 = ifres894


} else {
ifres893 = False


}

var ifres892 Obj

if True == ifres893 {
ifres892 = True


} else {
ifres892 = False


}

ifres891 = ifres892


} else {
ifres891 = False


}

if True == ifres891 {
__e.Return(PrimTail(V4109))
return
} else {
tmp889 := PrimIsPair(V4109)

if True == tmp889 {
tmp885 := PrimHead(V4109)

tmp886 := PrimTail(V4109)

tmp887 := Call(__e, PrimFunc(symshen_4unassoc), V4108, tmp886)


__e.Return(PrimCons(tmp885, tmp887))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.unassoc")))
return
}


}


}


}, 2)

tmp903 := Call(__e, ns2_1set, symshen_4unassoc, tmp884)


_ = tmp903

tmp904 := MakeNative(func(__e *ControlFlow) {
V4110 := __e.Get(1)
_ = V4110
tmp908 := Call(__e, PrimFunc(sympackage_2), V4110)


if True == tmp908 {
__e.Return(PrimSet(symshen_4_dpackage_d, V4110))
return
} else {
tmp905 := Call(__e, PrimFunc(symshen_4app), V4110, MakeString(" does not exist\n"), symshen_4a)


tmp906 := PrimStringConcat(MakeString("package "), tmp905)

__e.Return(PrimSimpleError(tmp906))
return


}


}, 1)

tmp909 := Call(__e, ns2_1set, symin_1package, tmp904)


_ = tmp909

tmp910 := MakeNative(func(__e *ControlFlow) {
V4111 := __e.Get(1)
_ = V4111
V4112 := __e.Get(2)
_ = V4112
tmp911 := MakeNative(func(__e *ControlFlow) {
W4113 := __e.Get(1)
_ = W4113
tmp912 := MakeNative(func(__e *ControlFlow) {
W4114 := __e.Get(1)
_ = W4114
tmp913 := MakeNative(func(__e *ControlFlow) {
W4115 := __e.Get(1)
_ = W4115
tmp914 := MakeNative(func(__e *ControlFlow) {
W4116 := __e.Get(1)
_ = W4116
__e.Return(V4112)
return
}, 1)

tmp915 := PrimCloseStream(W4113)

__e.TailApply(tmp914, tmp915)
return


}, 1)

tmp916 := Call(__e, PrimFunc(sympr), W4114, W4113)


__e.TailApply(tmp913, tmp916)
return


}, 1)

tmp919 := PrimIsString(V4112)

var ifres917 Obj

if True == tmp919 {
ifres917 = V4112


} else {
tmp918 := Call(__e, PrimFunc(symshen_4app), V4112, MakeString(""), symshen_4s)


ifres917 = tmp918


}

__e.TailApply(tmp912, ifres917)
return


}, 1)

tmp920 := PrimOpenStream(V4111, symout)

__e.TailApply(tmp911, tmp920)
return


}, 2)

tmp921 := Call(__e, ns2_1set, symwrite_1to_1file, tmp910)


_ = tmp921

tmp922 := MakeNative(func(__e *ControlFlow) {
tmp923 := Call(__e, PrimFunc(symgensym), symshen_4t)


__e.TailApply(PrimFunc(symshen_4freshterm), tmp923)
return


}, 0)

tmp924 := Call(__e, ns2_1set, symfresh, tmp922)


_ = tmp924

tmp925 := MakeNative(func(__e *ControlFlow) {
V4117 := __e.Get(1)
_ = V4117
V4118 := __e.Get(2)
_ = V4118
tmp926 := MakeNative(func(__e *ControlFlow) {
W4119 := __e.Get(1)
_ = W4119
tmp927 := MakeNative(func(__e *ControlFlow) {
W4120 := __e.Get(1)
_ = W4120
tmp928 := MakeNative(func(__e *ControlFlow) {
W4121 := __e.Get(1)
_ = W4121
__e.Return(V4117)
return
}, 1)

tmp929 := PrimCons(V4117, W4120)

tmp930 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp929)


__e.TailApply(tmp928, tmp930)
return


}, 1)

tmp931 := Call(__e, PrimFunc(symshen_4lambda_1entry), V4117)


__e.TailApply(tmp927, tmp931)
return


}, 1)

tmp932 := PrimValue(sym_dproperty_1vector_d)

tmp933 := Call(__e, PrimFunc(symput), V4117, symarity, V4118, tmp932)


__e.TailApply(tmp926, tmp933)
return


}, 2)

tmp934 := Call(__e, ns2_1set, symupdate_1lambda_1table, tmp925)


_ = tmp934

tmp935 := MakeNative(func(__e *ControlFlow) {
V4124 := __e.Get(1)
_ = V4124
V4125 := __e.Get(2)
_ = V4125
tmp959 := PrimEqual(MakeNumber(0), V4125)

if True == tmp959 {
tmp936 := PrimValue(symshen_4_dspecial_d)

tmp937 := Call(__e, PrimFunc(symremove), V4124, tmp936)


tmp938 := PrimSet(symshen_4_dspecial_d, tmp937)

_ = tmp938

tmp939 := PrimValue(symshen_4_dextraspecial_d)

tmp940 := Call(__e, PrimFunc(symremove), V4124, tmp939)


tmp941 := PrimSet(symshen_4_dextraspecial_d, tmp940)

_ = tmp941

__e.Return(V4124)
return


} else {
tmp957 := PrimEqual(MakeNumber(1), V4125)

if True == tmp957 {
tmp942 := PrimValue(symshen_4_dspecial_d)

tmp943 := Call(__e, PrimFunc(symadjoin), V4124, tmp942)


tmp944 := PrimSet(symshen_4_dspecial_d, tmp943)

_ = tmp944

tmp945 := PrimValue(symshen_4_dextraspecial_d)

tmp946 := Call(__e, PrimFunc(symremove), V4124, tmp945)


tmp947 := PrimSet(symshen_4_dextraspecial_d, tmp946)

_ = tmp947

__e.Return(V4124)
return


} else {
tmp955 := PrimEqual(MakeNumber(2), V4125)

if True == tmp955 {
tmp948 := PrimValue(symshen_4_dspecial_d)

tmp949 := Call(__e, PrimFunc(symremove), V4124, tmp948)


tmp950 := PrimSet(symshen_4_dspecial_d, tmp949)

_ = tmp950

tmp951 := PrimValue(symshen_4_dextraspecial_d)

tmp952 := Call(__e, PrimFunc(symadjoin), V4124, tmp951)


tmp953 := PrimSet(symshen_4_dextraspecial_d, tmp952)

_ = tmp953

__e.Return(V4124)
return


} else {
__e.Return(PrimSimpleError(MakeString("specialise requires values of 0, 1 or 2\n")))
return
}


}


}


}, 2)

tmp960 := Call(__e, ns2_1set, symspecialise, tmp935)


_ = tmp960

tmp961 := MakeNative(func(__e *ControlFlow) {
V4126 := __e.Get(1)
_ = V4126
tmp962 := PrimValue(sym_dabsolute_d)

tmp963 := PrimCons(V4126, tmp962)

__e.Return(PrimSet(sym_dabsolute_d, tmp963))
return


}, 1)

tmp964 := Call(__e, ns2_1set, symabsolute, tmp961)


_ = tmp964

tmp965 := MakeNative(func(__e *ControlFlow) {
V4127 := __e.Get(1)
_ = V4127
tmp966 := PrimValue(sym_dabsolute_d)

tmp967 := Call(__e, PrimFunc(symremove), V4127, tmp966)


__e.Return(PrimSet(sym_dabsolute_d, tmp967))
return


}, 1)

__e.TailApply(ns2_1set, symunabsolute, tmp965)
return




}, 0)

