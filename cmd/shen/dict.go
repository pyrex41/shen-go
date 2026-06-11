package main

import . "github.com/tiancaiamao/shen-go/kl"

var DictMain = MakeNative(func(__e *ControlFlow) {
tmp16955 := MakeNative(func(__e *ControlFlow) {
V4162 := __e.Get(1)
_ = V4162
tmp16971 := PrimLessThan(V4162, MakeNumber(1))

if True == tmp16971 {
tmp16956 := Call(__e, PrimFunc(symshen_4app), V4162, MakeString(""), symshen_4s)


tmp16957 := PrimStringConcat(MakeString("invalid initial dict size: "), tmp16956)

__e.Return(PrimSimpleError(tmp16957))
return


} else {
tmp16958 := MakeNative(func(__e *ControlFlow) {
W4163 := __e.Get(1)
_ = W4163
tmp16959 := MakeNative(func(__e *ControlFlow) {
W4164 := __e.Get(1)
_ = W4164
tmp16960 := MakeNative(func(__e *ControlFlow) {
W4165 := __e.Get(1)
_ = W4165
tmp16961 := MakeNative(func(__e *ControlFlow) {
W4166 := __e.Get(1)
_ = W4166
tmp16962 := MakeNative(func(__e *ControlFlow) {
W4167 := __e.Get(1)
_ = W4167
__e.Return(W4163)
return
}, 1)

tmp16963 := PrimNumberAdd(MakeNumber(2), V4162)

tmp16964 := Call(__e, PrimFunc(symshen_4fillvector), W4163, MakeNumber(3), tmp16963, Nil)


__e.TailApply(tmp16962, tmp16964)
return


}, 1)

tmp16965 := PrimVectorSet(W4163, MakeNumber(2), MakeNumber(0))

__e.TailApply(tmp16961, tmp16965)
return


}, 1)

tmp16966 := PrimVectorSet(W4163, MakeNumber(1), V4162)

__e.TailApply(tmp16960, tmp16966)
return


}, 1)

tmp16967 := PrimVectorSet(W4163, MakeNumber(0), symshen_4dictionary)

__e.TailApply(tmp16959, tmp16967)
return


}, 1)

tmp16968 := PrimNumberAdd(MakeNumber(3), V4162)

tmp16969 := PrimAbsvector(tmp16968)

__e.TailApply(tmp16958, tmp16969)
return


}


}, 1)

tmp16972 := Call(__e, ns2_1set, symshen_4dict, tmp16955)


_ = tmp16972

tmp16973 := MakeNative(func(__e *ControlFlow) {
V4168 := __e.Get(1)
_ = V4168
tmp16980 := PrimIsVector(V4168)

if True == tmp16980 {
tmp16975 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimVectorGet(V4168, MakeNumber(0)))
return
}, 0)

tmp16976 := MakeNative(func(__e *ControlFlow) {
Z4169 := __e.Get(1)
_ = Z4169
__e.Return(symshen_4not_1dictionary)
return
}, 1)

tmp16977 := Call(__e, try_1catch, tmp16975, tmp16976)


tmp16978 := PrimEqual(tmp16977, symshen_4dictionary)

if True == tmp16978 {
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

tmp16981 := Call(__e, ns2_1set, symshen_4dict_2, tmp16973)


_ = tmp16981

tmp16982 := MakeNative(func(__e *ControlFlow) {
V4170 := __e.Get(1)
_ = V4170
__e.Return(PrimVectorGet(V4170, MakeNumber(1)))
return
}, 1)

tmp16983 := Call(__e, ns2_1set, symshen_4dict_1capacity, tmp16982)


_ = tmp16983

tmp16984 := MakeNative(func(__e *ControlFlow) {
V4171 := __e.Get(1)
_ = V4171
__e.Return(PrimVectorGet(V4171, MakeNumber(2)))
return
}, 1)

tmp16985 := Call(__e, ns2_1set, symshen_4dict_1count, tmp16984)


_ = tmp16985

tmp16986 := MakeNative(func(__e *ControlFlow) {
V4172 := __e.Get(1)
_ = V4172
V4173 := __e.Get(2)
_ = V4173
__e.Return(PrimVectorSet(V4172, MakeNumber(2), V4173))
return
}, 2)

tmp16987 := Call(__e, ns2_1set, symshen_4dict_1count_1_6, tmp16986)


_ = tmp16987

tmp16988 := MakeNative(func(__e *ControlFlow) {
V4174 := __e.Get(1)
_ = V4174
V4175 := __e.Get(2)
_ = V4175
tmp16989 := PrimNumberAdd(MakeNumber(3), V4175)

__e.Return(PrimVectorGet(V4174, tmp16989))
return


}, 2)

tmp16990 := Call(__e, ns2_1set, symshen_4_5_1dict_1bucket, tmp16988)


_ = tmp16990

tmp16991 := MakeNative(func(__e *ControlFlow) {
V4176 := __e.Get(1)
_ = V4176
V4177 := __e.Get(2)
_ = V4177
V4178 := __e.Get(3)
_ = V4178
tmp16992 := PrimNumberAdd(MakeNumber(3), V4177)

__e.Return(PrimVectorSet(V4176, tmp16992, V4178))
return


}, 3)

tmp16993 := Call(__e, ns2_1set, symshen_4dict_1bucket_1_6, tmp16991)


_ = tmp16993

tmp16994 := MakeNative(func(__e *ControlFlow) {
V4179 := __e.Get(1)
_ = V4179
V4180 := __e.Get(2)
_ = V4180
V4181 := __e.Get(3)
_ = V4181
tmp16995 := MakeNative(func(__e *ControlFlow) {
W4182 := __e.Get(1)
_ = W4182
tmp16996 := Call(__e, PrimFunc(symshen_4dict_1count), V4179)


tmp16997 := PrimNumberAdd(W4182, tmp16996)

__e.TailApply(PrimFunc(symshen_4dict_1count_1_6), V4179, tmp16997)
return


}, 1)

tmp16998 := Call(__e, PrimFunc(symlength), V4181)


tmp16999 := Call(__e, PrimFunc(symlength), V4180)


tmp17000 := PrimNumberSubtract(tmp16998, tmp16999)

__e.TailApply(tmp16995, tmp17000)
return


}, 3)

tmp17001 := Call(__e, ns2_1set, symshen_4dict_1update_1count, tmp16994)


_ = tmp17001

tmp17002 := MakeNative(func(__e *ControlFlow) {
V4183 := __e.Get(1)
_ = V4183
V4184 := __e.Get(2)
_ = V4184
V4185 := __e.Get(3)
_ = V4185
tmp17003 := MakeNative(func(__e *ControlFlow) {
W4186 := __e.Get(1)
_ = W4186
tmp17004 := MakeNative(func(__e *ControlFlow) {
W4187 := __e.Get(1)
_ = W4187
tmp17005 := MakeNative(func(__e *ControlFlow) {
W4188 := __e.Get(1)
_ = W4188
tmp17006 := MakeNative(func(__e *ControlFlow) {
W4189 := __e.Get(1)
_ = W4189
tmp17007 := MakeNative(func(__e *ControlFlow) {
W4190 := __e.Get(1)
_ = W4190
__e.Return(V4185)
return
}, 1)

tmp17008 := Call(__e, PrimFunc(symshen_4dict_1update_1count), V4183, W4187, W4188)


__e.TailApply(tmp17007, tmp17008)
return


}, 1)

tmp17009 := Call(__e, PrimFunc(symshen_4dict_1bucket_1_6), V4183, W4186, W4188)


__e.TailApply(tmp17006, tmp17009)
return


}, 1)

tmp17010 := Call(__e, PrimFunc(symshen_4assoc_1set), V4184, V4185, W4187)


__e.TailApply(tmp17005, tmp17010)
return


}, 1)

tmp17011 := Call(__e, PrimFunc(symshen_4_5_1dict_1bucket), V4183, W4186)


__e.TailApply(tmp17004, tmp17011)
return


}, 1)

tmp17012 := Call(__e, PrimFunc(symshen_4dict_1capacity), V4183)


tmp17013 := Call(__e, PrimFunc(symhash), V4184, tmp17012)


__e.TailApply(tmp17003, tmp17013)
return


}, 3)

tmp17014 := Call(__e, ns2_1set, symshen_4dict_1_6, tmp17002)


_ = tmp17014

tmp17015 := MakeNative(func(__e *ControlFlow) {
V4191 := __e.Get(1)
_ = V4191
V4192 := __e.Get(2)
_ = V4192
tmp17016 := MakeNative(func(__e *ControlFlow) {
W4193 := __e.Get(1)
_ = W4193
tmp17017 := MakeNative(func(__e *ControlFlow) {
W4194 := __e.Get(1)
_ = W4194
tmp17018 := MakeNative(func(__e *ControlFlow) {
W4195 := __e.Get(1)
_ = W4195
tmp17022 := Call(__e, PrimFunc(symempty_2), W4195)


if True == tmp17022 {
tmp17019 := Call(__e, PrimFunc(symshen_4app), V4192, MakeString(" not found in dict\n"), symshen_4a)


tmp17020 := PrimStringConcat(MakeString("value "), tmp17019)

__e.Return(PrimSimpleError(tmp17020))
return


} else {
__e.Return(PrimTail(W4195))
return
}


}, 1)

tmp17023 := Call(__e, PrimFunc(symassoc), V4192, W4194)


__e.TailApply(tmp17018, tmp17023)
return


}, 1)

tmp17024 := Call(__e, PrimFunc(symshen_4_5_1dict_1bucket), V4191, W4193)


__e.TailApply(tmp17017, tmp17024)
return


}, 1)

tmp17025 := Call(__e, PrimFunc(symshen_4dict_1capacity), V4191)


tmp17026 := Call(__e, PrimFunc(symhash), V4192, tmp17025)


__e.TailApply(tmp17016, tmp17026)
return


}, 2)

tmp17027 := Call(__e, ns2_1set, symshen_4_5_1dict, tmp17015)


_ = tmp17027

tmp17028 := MakeNative(func(__e *ControlFlow) {
V4196 := __e.Get(1)
_ = V4196
V4197 := __e.Get(2)
_ = V4197
tmp17029 := MakeNative(func(__e *ControlFlow) {
W4198 := __e.Get(1)
_ = W4198
tmp17030 := MakeNative(func(__e *ControlFlow) {
W4199 := __e.Get(1)
_ = W4199
tmp17031 := MakeNative(func(__e *ControlFlow) {
W4200 := __e.Get(1)
_ = W4200
tmp17032 := MakeNative(func(__e *ControlFlow) {
W4201 := __e.Get(1)
_ = W4201
tmp17033 := MakeNative(func(__e *ControlFlow) {
W4202 := __e.Get(1)
_ = W4202
__e.Return(V4197)
return
}, 1)

tmp17034 := Call(__e, PrimFunc(symshen_4dict_1update_1count), V4196, W4199, W4200)


__e.TailApply(tmp17033, tmp17034)
return


}, 1)

tmp17035 := Call(__e, PrimFunc(symshen_4dict_1bucket_1_6), V4196, W4198, W4200)


__e.TailApply(tmp17032, tmp17035)
return


}, 1)

tmp17036 := Call(__e, PrimFunc(symshen_4assoc_1rm), V4197, W4199)


__e.TailApply(tmp17031, tmp17036)
return


}, 1)

tmp17037 := Call(__e, PrimFunc(symshen_4_5_1dict_1bucket), V4196, W4198)


__e.TailApply(tmp17030, tmp17037)
return


}, 1)

tmp17038 := Call(__e, PrimFunc(symshen_4dict_1capacity), V4196)


tmp17039 := Call(__e, PrimFunc(symhash), V4197, tmp17038)


__e.TailApply(tmp17029, tmp17039)
return


}, 2)

tmp17040 := Call(__e, ns2_1set, symshen_4dict_1rm, tmp17028)


_ = tmp17040

tmp17041 := MakeNative(func(__e *ControlFlow) {
V4203 := __e.Get(1)
_ = V4203
V4204 := __e.Get(2)
_ = V4204
V4205 := __e.Get(3)
_ = V4205
tmp17042 := MakeNative(func(__e *ControlFlow) {
W4206 := __e.Get(1)
_ = W4206
__e.TailApply(PrimFunc(symshen_4dict_1fold_1h), V4203, V4204, V4205, MakeNumber(0), W4206)
return
}, 1)

tmp17043 := Call(__e, PrimFunc(symshen_4dict_1capacity), V4204)


__e.TailApply(tmp17042, tmp17043)
return


}, 3)

tmp17044 := Call(__e, ns2_1set, symshen_4dict_1fold, tmp17041)


_ = tmp17044

tmp17045 := MakeNative(func(__e *ControlFlow) {
V4208 := __e.Get(1)
_ = V4208
V4209 := __e.Get(2)
_ = V4209
V4210 := __e.Get(3)
_ = V4210
V4211 := __e.Get(4)
_ = V4211
V4212 := __e.Get(5)
_ = V4212
tmp17052 := PrimEqual(V4211, V4212)

if True == tmp17052 {
__e.Return(V4210)
return
} else {
tmp17046 := MakeNative(func(__e *ControlFlow) {
W4213 := __e.Get(1)
_ = W4213
tmp17047 := MakeNative(func(__e *ControlFlow) {
W4214 := __e.Get(1)
_ = W4214
tmp17048 := PrimNumberAdd(MakeNumber(1), V4211)

__e.TailApply(PrimFunc(symshen_4dict_1fold_1h), V4208, V4209, W4214, tmp17048, V4212)
return


}, 1)

tmp17049 := Call(__e, PrimFunc(symshen_4bucket_1fold), V4208, W4213, V4210)


__e.TailApply(tmp17047, tmp17049)
return


}, 1)

tmp17050 := Call(__e, PrimFunc(symshen_4_5_1dict_1bucket), V4209, V4211)


__e.TailApply(tmp17046, tmp17050)
return


}


}, 5)

tmp17053 := Call(__e, ns2_1set, symshen_4dict_1fold_1h, tmp17045)


_ = tmp17053

tmp17054 := MakeNative(func(__e *ControlFlow) {
V4215 := __e.Get(1)
_ = V4215
V4216 := __e.Get(2)
_ = V4216
V4217 := __e.Get(3)
_ = V4217
tmp17070 := PrimEqual(Nil, V4216)

if True == tmp17070 {
__e.Return(V4217)
return
} else {
tmp17068 := PrimIsPair(V4216)

var ifres17064 Obj

if True == tmp17068 {
tmp17066 := PrimHead(V4216)

tmp17067 := PrimIsPair(tmp17066)

var ifres17065 Obj

if True == tmp17067 {
ifres17065 = True


} else {
ifres17065 = False


}

ifres17064 = ifres17065


} else {
ifres17064 = False


}

if True == ifres17064 {
tmp17055 := PrimHead(V4216)

tmp17056 := PrimHead(tmp17055)

tmp17057 := Call(__e, V4215, tmp17056)


tmp17058 := PrimHead(V4216)

tmp17059 := PrimTail(tmp17058)

tmp17060 := Call(__e, tmp17057, tmp17059)


tmp17061 := PrimTail(V4216)

tmp17062 := Call(__e, PrimFunc(symshen_4bucket_1fold), V4215, tmp17061, V4217)


__e.TailApply(tmp17060, tmp17062)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4bucket_1fold)
return
}


}


}, 3)

tmp17071 := Call(__e, ns2_1set, symshen_4bucket_1fold, tmp17054)


_ = tmp17071

tmp17072 := MakeNative(func(__e *ControlFlow) {
V4218 := __e.Get(1)
_ = V4218
tmp17073 := MakeNative(func(__e *ControlFlow) {
Z4219 := __e.Get(1)
_ = Z4219
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4220 := __e.Get(1)
_ = Z4220
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4221 := __e.Get(1)
_ = Z4221
__e.Return(PrimCons(Z4219, Z4221))
return
}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(PrimFunc(symshen_4dict_1fold), tmp17073, V4218, Nil)
return


}, 1)

tmp17074 := Call(__e, ns2_1set, symshen_4dict_1keys, tmp17072)


_ = tmp17074

tmp17075 := MakeNative(func(__e *ControlFlow) {
V4222 := __e.Get(1)
_ = V4222
tmp17076 := MakeNative(func(__e *ControlFlow) {
Z4223 := __e.Get(1)
_ = Z4223
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4224 := __e.Get(1)
_ = Z4224
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4225 := __e.Get(1)
_ = Z4225
__e.Return(PrimCons(Z4224, Z4225))
return
}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(PrimFunc(symshen_4dict_1fold), tmp17076, V4222, Nil)
return


}, 1)

__e.TailApply(ns2_1set, symshen_4dict_1values, tmp17075)
return




}, 0)

