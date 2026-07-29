package main

import . "github.com/pyrex41/shen-go/kl"

var ReaderMain = MakeNative(func(__e *ControlFlow) {
tmp4747 := MakeNative(func(__e *ControlFlow) {
V2182 := __e.Get(1)
_ = V2182
tmp4748 := MakeNative(func(__e *ControlFlow) {
W2183 := __e.Get(1)
_ = W2183
tmp4749 := MakeNative(func(__e *ControlFlow) {
W2184 := __e.Get(1)
_ = W2184
tmp4750 := MakeNative(func(__e *ControlFlow) {
W2187 := __e.Get(1)
_ = W2187
__e.Return(W2187)
return
}, 1)

tmp4751 := Call(__e, PrimFunc(symshen_4process_1sexprs), W2184)


__e.TailApply(tmp4750, tmp4751)
return


}, 1)

tmp4752 := MakeNative(func(__e *ControlFlow) {
tmp4753 := MakeNative(func(__e *ControlFlow) {
Z2185 := __e.Get(1)
_ = Z2185
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2185)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp4753, W2183)
return


}, 0)

tmp4754 := MakeNative(func(__e *ControlFlow) {
Z2186 := __e.Get(1)
_ = Z2186
tmp4755 := PrimValue(symshen_4_dresidue_d)

__e.TailApply(PrimFunc(symshen_4reader_1error), tmp4755)
return


}, 1)

tmp4756 := Call(__e, try_1catch, tmp4752, tmp4754)


__e.TailApply(tmp4749, tmp4756)
return


}, 1)

tmp4757 := PrimReadFileAsByteList(V2182)

__e.TailApply(tmp4748, tmp4757)
return


}, 1)

tmp4758 := Call(__e, ns2_1set, symread_1file, tmp4747)


_ = tmp4758

tmp4759 := MakeNative(func(__e *ControlFlow) {
V2188 := __e.Get(1)
_ = V2188
tmp4760 := PrimValue(sym_dmaximum_1print_1sequence_1size_d)

tmp4761 := Call(__e, PrimFunc(symshen_4reader_1error_1message), tmp4760, MakeNumber(0), V2188)


tmp4762 := PrimStringConcat(MakeString("reader error near here: "), tmp4761)

tmp4763 := Call(__e, PrimFunc(symshen_4proc_1nl), tmp4762)


__e.Return(PrimSimpleError(tmp4763))
return


}, 1)

tmp4764 := Call(__e, ns2_1set, symshen_4reader_1error, tmp4759)


_ = tmp4764

tmp4765 := MakeNative(func(__e *ControlFlow) {
V2196 := __e.Get(1)
_ = V2196
V2197 := __e.Get(2)
_ = V2197
V2198 := __e.Get(3)
_ = V2198
tmp4776 := PrimEqual(Nil, V2198)

if True == tmp4776 {
__e.Return(MakeString(""))
return
} else {
tmp4774 := PrimEqual(V2196, V2197)

if True == tmp4774 {
__e.Return(MakeString(""))
return
} else {
tmp4772 := PrimIsPair(V2198)

if True == tmp4772 {
tmp4766 := PrimHead(V2198)

tmp4767 := PrimNumberToString(tmp4766)

tmp4768 := PrimNumberAdd(V2197, MakeNumber(1))

tmp4769 := PrimTail(V2198)

tmp4770 := Call(__e, PrimFunc(symshen_4reader_1error_1message), V2196, tmp4768, tmp4769)


__e.Return(PrimStringConcat(tmp4767, tmp4770))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.reader-error-message")))
return
}


}


}


}, 3)

tmp4777 := Call(__e, ns2_1set, symshen_4reader_1error_1message, tmp4765)


_ = tmp4777

tmp4778 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dit_d))
return
}, 0)

tmp4779 := Call(__e, ns2_1set, symit, tmp4778)


_ = tmp4779

tmp4780 := MakeNative(func(__e *ControlFlow) {
V2199 := __e.Get(1)
_ = V2199
tmp4781 := MakeNative(func(__e *ControlFlow) {
W2200 := __e.Get(1)
_ = W2200
tmp4782 := MakeNative(func(__e *ControlFlow) {
W2201 := __e.Get(1)
_ = W2201
tmp4783 := MakeNative(func(__e *ControlFlow) {
W2202 := __e.Get(1)
_ = W2202
tmp4784 := MakeNative(func(__e *ControlFlow) {
W2203 := __e.Get(1)
_ = W2203
__e.TailApply(PrimFunc(symreverse), W2202)
return
}, 1)

tmp4785 := PrimCloseStream(W2200)

__e.TailApply(tmp4784, tmp4785)
return


}, 1)

tmp4786 := Call(__e, PrimFunc(symshen_4read_1file_1as_1bytelist_1help), W2200, W2201, Nil)


__e.TailApply(tmp4783, tmp4786)
return


}, 1)

tmp4787 := PrimReadByte(W2200)

__e.TailApply(tmp4782, tmp4787)
return


}, 1)

tmp4788 := PrimOpenStream(V2199, symin)

__e.TailApply(tmp4781, tmp4788)
return


}, 1)

tmp4789 := Call(__e, ns2_1set, symread_1file_1as_1bytelist, tmp4780)


_ = tmp4789

tmp4790 := MakeNative(func(__e *ControlFlow) {
V2204 := __e.Get(1)
_ = V2204
V2205 := __e.Get(2)
_ = V2205
V2206 := __e.Get(3)
_ = V2206
tmp4794 := PrimEqual(MakeNumber(-1), V2205)

if True == tmp4794 {
__e.Return(V2206)
return
} else {
tmp4791 := PrimReadByte(V2204)

tmp4792 := PrimCons(V2205, V2206)

__e.TailApply(PrimFunc(symshen_4read_1file_1as_1bytelist_1help), V2204, tmp4791, tmp4792)
return


}


}, 3)

tmp4795 := Call(__e, ns2_1set, symshen_4read_1file_1as_1bytelist_1help, tmp4790)


_ = tmp4795

tmp4796 := MakeNative(func(__e *ControlFlow) {
V2207 := __e.Get(1)
_ = V2207
tmp4797 := MakeNative(func(__e *ControlFlow) {
W2208 := __e.Get(1)
_ = W2208
tmp4798 := PrimReadByte(W2208)

__e.TailApply(PrimFunc(symshen_4rfas_1h), W2208, tmp4798, MakeString(""))
return


}, 1)

tmp4799 := PrimOpenStream(V2207, symin)

__e.TailApply(tmp4797, tmp4799)
return


}, 1)

tmp4800 := Call(__e, ns2_1set, symread_1file_1as_1string, tmp4796)


_ = tmp4800

tmp4801 := MakeNative(func(__e *ControlFlow) {
V2209 := __e.Get(1)
_ = V2209
V2210 := __e.Get(2)
_ = V2210
V2211 := __e.Get(3)
_ = V2211
tmp4807 := PrimEqual(MakeNumber(-1), V2210)

if True == tmp4807 {
tmp4802 := PrimCloseStream(V2209)

_ = tmp4802

__e.Return(V2211)
return


} else {
tmp4803 := PrimReadByte(V2209)

tmp4804 := PrimNumberToString(V2210)

tmp4805 := PrimStringConcat(V2211, tmp4804)

__e.TailApply(PrimFunc(symshen_4rfas_1h), V2209, tmp4803, tmp4805)
return


}


}, 3)

tmp4808 := Call(__e, ns2_1set, symshen_4rfas_1h, tmp4801)


_ = tmp4808

tmp4809 := MakeNative(func(__e *ControlFlow) {
V2212 := __e.Get(1)
_ = V2212
tmp4810 := Call(__e, PrimFunc(symread), V2212)


__e.TailApply(PrimFunc(symeval_1kl), tmp4810)
return


}, 1)

tmp4811 := Call(__e, ns2_1set, syminput, tmp4809)


_ = tmp4811

tmp4812 := MakeNative(func(__e *ControlFlow) {
V2213 := __e.Get(1)
_ = V2213
V2214 := __e.Get(2)
_ = V2214
tmp4813 := MakeNative(func(__e *ControlFlow) {
W2215 := __e.Get(1)
_ = W2215
tmp4814 := MakeNative(func(__e *ControlFlow) {
W2216 := __e.Get(1)
_ = W2216
tmp4820 := Call(__e, PrimFunc(symshen_4typecheck), W2216, V2213)


tmp4821 := PrimEqual(False, tmp4820)

if True == tmp4821 {
tmp4815 := Call(__e, PrimFunc(symshen_4app), V2213, MakeString("\n"), symshen_4r)


tmp4816 := PrimStringConcat(MakeString(" is not of type "), tmp4815)

tmp4817 := Call(__e, PrimFunc(symshen_4app), W2216, tmp4816, symshen_4r)


tmp4818 := PrimStringConcat(MakeString("type error: "), tmp4817)

__e.Return(PrimSimpleError(tmp4818))
return


} else {
__e.TailApply(PrimFunc(symeval_1kl), W2216)
return
}


}, 1)

tmp4822 := Call(__e, PrimFunc(symread), V2214)


__e.TailApply(tmp4814, tmp4822)
return


}, 1)

tmp4823 := Call(__e, PrimFunc(symshen_4monotype), V2213)


__e.TailApply(tmp4813, tmp4823)
return


}, 2)

tmp4824 := Call(__e, ns2_1set, symshen_4input_1h_7, tmp4812)


_ = tmp4824

tmp4825 := MakeNative(func(__e *ControlFlow) {
V2217 := __e.Get(1)
_ = V2217
tmp4832 := PrimIsPair(V2217)

if True == tmp4832 {
tmp4826 := MakeNative(func(__e *ControlFlow) {
Z2218 := __e.Get(1)
_ = Z2218
__e.TailApply(PrimFunc(symshen_4monotype), Z2218)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp4826, V2217)
return


} else {
tmp4830 := PrimIsVariable(V2217)

if True == tmp4830 {
tmp4827 := Call(__e, PrimFunc(symshen_4app), V2217, MakeString("\n"), symshen_4a)


tmp4828 := PrimStringConcat(MakeString("input+ expects a monotype: not "), tmp4827)

__e.Return(PrimSimpleError(tmp4828))
return


} else {
__e.Return(V2217)
return
}


}


}, 1)

tmp4833 := Call(__e, ns2_1set, symshen_4monotype, tmp4825)


_ = tmp4833

tmp4834 := MakeNative(func(__e *ControlFlow) {
V2219 := __e.Get(1)
_ = V2219
tmp4835 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2219)


tmp4836 := MakeNative(func(__e *ControlFlow) {
Z2220 := __e.Get(1)
_ = Z2220
__e.TailApply(PrimFunc(symshen_4return_2), Z2220)
return
}, 1)

__e.TailApply(PrimFunc(symshen_4read_1loop), V2219, tmp4835, Nil, tmp4836)
return


}, 1)

tmp4837 := Call(__e, ns2_1set, symlineread, tmp4834)


_ = tmp4837

tmp4838 := MakeNative(func(__e *ControlFlow) {
V2221 := __e.Get(1)
_ = V2221
tmp4839 := MakeNative(func(__e *ControlFlow) {
W2222 := __e.Get(1)
_ = W2222
tmp4840 := MakeNative(func(__e *ControlFlow) {
W2223 := __e.Get(1)
_ = W2223
tmp4841 := MakeNative(func(__e *ControlFlow) {
W2225 := __e.Get(1)
_ = W2225
__e.Return(W2225)
return
}, 1)

tmp4842 := Call(__e, PrimFunc(symshen_4process_1sexprs), W2223)


__e.TailApply(tmp4841, tmp4842)
return


}, 1)

tmp4843 := MakeNative(func(__e *ControlFlow) {
Z2224 := __e.Get(1)
_ = Z2224
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2224)
return
}, 1)

tmp4844 := Call(__e, PrimFunc(symcompile), tmp4843, W2222)


__e.TailApply(tmp4840, tmp4844)
return


}, 1)

tmp4845 := Call(__e, PrimFunc(symshen_4str_1_6bytes), V2221)


__e.TailApply(tmp4839, tmp4845)
return


}, 1)

tmp4846 := Call(__e, ns2_1set, symread_1from_1string, tmp4838)


_ = tmp4846

tmp4847 := MakeNative(func(__e *ControlFlow) {
V2226 := __e.Get(1)
_ = V2226
tmp4848 := MakeNative(func(__e *ControlFlow) {
W2227 := __e.Get(1)
_ = W2227
tmp4849 := MakeNative(func(__e *ControlFlow) {
W2228 := __e.Get(1)
_ = W2228
__e.Return(W2228)
return
}, 1)

tmp4850 := MakeNative(func(__e *ControlFlow) {
Z2229 := __e.Get(1)
_ = Z2229
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2229)
return
}, 1)

tmp4851 := Call(__e, PrimFunc(symcompile), tmp4850, W2227)


__e.TailApply(tmp4849, tmp4851)
return


}, 1)

tmp4852 := Call(__e, PrimFunc(symshen_4str_1_6bytes), V2226)


__e.TailApply(tmp4848, tmp4852)
return


}, 1)

tmp4853 := Call(__e, ns2_1set, symread_1from_1string_1unprocessed, tmp4847)


_ = tmp4853

tmp4854 := MakeNative(func(__e *ControlFlow) {
V2230 := __e.Get(1)
_ = V2230
tmp4862 := PrimEqual(MakeString(""), V2230)

if True == tmp4862 {
__e.Return(Nil)
return
} else {
tmp4860 := Call(__e, PrimFunc(symshen_4_7string_2), V2230)


if True == tmp4860 {
tmp4855 := Call(__e, PrimFunc(symhdstr), V2230)


tmp4856 := PrimStringToNumber(tmp4855)

tmp4857 := PrimTailString(V2230)

tmp4858 := Call(__e, PrimFunc(symshen_4str_1_6bytes), tmp4857)


__e.Return(PrimCons(tmp4856, tmp4858))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.str->bytes")))
return
}


}


}, 1)

tmp4863 := Call(__e, ns2_1set, symshen_4str_1_6bytes, tmp4854)


_ = tmp4863

tmp4864 := MakeNative(func(__e *ControlFlow) {
V2231 := __e.Get(1)
_ = V2231
tmp4865 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2231)


tmp4866 := MakeNative(func(__e *ControlFlow) {
Z2232 := __e.Get(1)
_ = Z2232
__e.TailApply(PrimFunc(symshen_4whitespace_2), Z2232)
return
}, 1)

tmp4867 := Call(__e, PrimFunc(symshen_4read_1loop), V2231, tmp4865, Nil, tmp4866)


__e.Return(PrimHead(tmp4867))
return


}, 1)

tmp4868 := Call(__e, ns2_1set, symread, tmp4864)


_ = tmp4868

tmp4869 := MakeNative(func(__e *ControlFlow) {
V2233 := __e.Get(1)
_ = V2233
tmp4872 := Call(__e, PrimFunc(symshen_4char_1stinput_2), V2233)


if True == tmp4872 {
tmp4870 := Call(__e, PrimFunc(symshen_4read_1unit_1string), V2233)


__e.Return(PrimStringToNumber(tmp4870))
return


} else {
__e.Return(PrimReadByte(V2233))
return
}


}, 1)

tmp4873 := Call(__e, ns2_1set, symshen_4my_1read_1byte, tmp4869)


_ = tmp4873

tmp4874 := MakeNative(func(__e *ControlFlow) {
V2238 := __e.Get(1)
_ = V2238
V2239 := __e.Get(2)
_ = V2239
V2240 := __e.Get(3)
_ = V2240
V2241 := __e.Get(4)
_ = V2241
tmp4897 := PrimEqual(MakeNumber(94), V2239)

if True == tmp4897 {
__e.Return(PrimSimpleError(MakeString("read aborted")))
return
} else {
tmp4895 := PrimEqual(MakeNumber(-1), V2239)

if True == tmp4895 {
tmp4877 := Call(__e, PrimFunc(symempty_2), V2240)


if True == tmp4877 {
__e.Return(PrimSimpleError(MakeString("error: empty stream")))
return
} else {
tmp4875 := MakeNative(func(__e *ControlFlow) {
Z2242 := __e.Get(1)
_ = Z2242
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2242)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp4875, V2240)
return


}


} else {
tmp4893 := PrimEqual(MakeNumber(0), V2239)

if True == tmp4893 {
tmp4878 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2238)


__e.TailApply(PrimFunc(symshen_4read_1loop), V2238, tmp4878, V2240, V2241)
return


} else {
tmp4891 := Call(__e, V2241, V2239)


if True == tmp4891 {
tmp4879 := MakeNative(func(__e *ControlFlow) {
W2243 := __e.Get(1)
_ = W2243
tmp4885 := Call(__e, PrimFunc(symshen_4nothing_1doing_2), W2243)


if True == tmp4885 {
tmp4880 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2238)


tmp4881 := PrimCons(V2239, Nil)

tmp4882 := Call(__e, PrimFunc(symappend), V2240, tmp4881)


__e.TailApply(PrimFunc(symshen_4read_1loop), V2238, tmp4880, tmp4882, V2241)
return


} else {
tmp4883 := Call(__e, PrimFunc(symshen_4record_1it), V2240)


_ = tmp4883

__e.Return(W2243)
return


}


}, 1)

tmp4886 := Call(__e, PrimFunc(symshen_4try_1parse), V2240)


__e.TailApply(tmp4879, tmp4886)
return


} else {
tmp4887 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2238)


tmp4888 := PrimCons(V2239, Nil)

tmp4889 := Call(__e, PrimFunc(symappend), V2240, tmp4888)


__e.TailApply(PrimFunc(symshen_4read_1loop), V2238, tmp4887, tmp4889, V2241)
return


}


}


}


}


}, 4)

tmp4898 := Call(__e, ns2_1set, symshen_4read_1loop, tmp4874)


_ = tmp4898

tmp4899 := MakeNative(func(__e *ControlFlow) {
V2244 := __e.Get(1)
_ = V2244
tmp4900 := MakeNative(func(__e *ControlFlow) {
W2245 := __e.Get(1)
_ = W2245
tmp4902 := Call(__e, PrimFunc(symshen_4nothing_1doing_2), W2245)


if True == tmp4902 {
__e.Return(symshen_4i_1failed_b)
return
} else {
__e.TailApply(PrimFunc(symshen_4process_1sexprs), W2245)
return
}


}, 1)

tmp4903 := MakeNative(func(__e *ControlFlow) {
tmp4904 := MakeNative(func(__e *ControlFlow) {
Z2246 := __e.Get(1)
_ = Z2246
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2246)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp4904, V2244)
return


}, 0)

tmp4905 := MakeNative(func(__e *ControlFlow) {
Z2247 := __e.Get(1)
_ = Z2247
__e.Return(symshen_4i_1failed_b)
return
}, 1)

tmp4906 := Call(__e, try_1catch, tmp4903, tmp4905)


__e.TailApply(tmp4900, tmp4906)
return


}, 1)

tmp4907 := Call(__e, ns2_1set, symshen_4try_1parse, tmp4899)


_ = tmp4907

tmp4908 := MakeNative(func(__e *ControlFlow) {
V2250 := __e.Get(1)
_ = V2250
tmp4912 := PrimEqual(symshen_4i_1failed_b, V2250)

if True == tmp4912 {
__e.Return(True)
return
} else {
tmp4910 := PrimEqual(Nil, V2250)

if True == tmp4910 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp4913 := Call(__e, ns2_1set, symshen_4nothing_1doing_2, tmp4908)


_ = tmp4913

tmp4914 := MakeNative(func(__e *ControlFlow) {
V2251 := __e.Get(1)
_ = V2251
tmp4915 := Call(__e, PrimFunc(symshen_4bytes_1_6string), V2251)


__e.Return(PrimSet(symshen_4_dit_d, tmp4915))
return


}, 1)

tmp4916 := Call(__e, ns2_1set, symshen_4record_1it, tmp4914)


_ = tmp4916

tmp4917 := MakeNative(func(__e *ControlFlow) {
V2252 := __e.Get(1)
_ = V2252
tmp4925 := PrimEqual(Nil, V2252)

if True == tmp4925 {
__e.Return(MakeString(""))
return
} else {
tmp4923 := PrimIsPair(V2252)

if True == tmp4923 {
tmp4918 := PrimHead(V2252)

tmp4919 := PrimNumberToString(tmp4918)

tmp4920 := PrimTail(V2252)

tmp4921 := Call(__e, PrimFunc(symshen_4bytes_1_6string), tmp4920)


__e.Return(PrimStringConcat(tmp4919, tmp4921))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.bytes->string")))
return
}


}


}, 1)

tmp4926 := Call(__e, ns2_1set, symshen_4bytes_1_6string, tmp4917)


_ = tmp4926

tmp4927 := MakeNative(func(__e *ControlFlow) {
V2253 := __e.Get(1)
_ = V2253
tmp4928 := MakeNative(func(__e *ControlFlow) {
W2254 := __e.Get(1)
_ = W2254
tmp4929 := MakeNative(func(__e *ControlFlow) {
W2255 := __e.Get(1)
_ = W2255
tmp4930 := MakeNative(func(__e *ControlFlow) {
W2256 := __e.Get(1)
_ = W2256
tmp4931 := MakeNative(func(__e *ControlFlow) {
Z2257 := __e.Get(1)
_ = Z2257
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2257, W2256)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp4931, W2254)
return


}, 1)

tmp4932 := Call(__e, PrimFunc(symshen_4find_1types), W2254)


__e.TailApply(tmp4930, tmp4932)
return


}, 1)

tmp4933 := Call(__e, PrimFunc(symshen_4find_1arities), W2254)


__e.TailApply(tmp4929, tmp4933)
return


}, 1)

tmp4934 := Call(__e, PrimFunc(symshen_4unpackage_emacroexpand), V2253)


__e.TailApply(tmp4928, tmp4934)
return


}, 1)

tmp4935 := Call(__e, ns2_1set, symshen_4process_1sexprs, tmp4927)


_ = tmp4935

tmp4936 := MakeNative(func(__e *ControlFlow) {
V2258 := __e.Get(1)
_ = V2258
tmp4958 := PrimIsPair(V2258)

var ifres4949 Obj

if True == tmp4958 {
tmp4956 := PrimTail(V2258)

tmp4957 := PrimIsPair(tmp4956)

var ifres4951 Obj

if True == tmp4957 {
tmp4953 := PrimHead(V2258)

tmp4954 := PrimIntern(MakeString(":"))

tmp4955 := PrimEqual(tmp4953, tmp4954)

var ifres4952 Obj

if True == tmp4955 {
ifres4952 = True


} else {
ifres4952 = False


}

ifres4951 = ifres4952


} else {
ifres4951 = False


}

var ifres4950 Obj

if True == ifres4951 {
ifres4950 = True


} else {
ifres4950 = False


}

ifres4949 = ifres4950


} else {
ifres4949 = False


}

if True == ifres4949 {
tmp4937 := PrimTail(V2258)

tmp4938 := PrimHead(tmp4937)

tmp4939 := PrimTail(V2258)

tmp4940 := PrimTail(tmp4939)

tmp4941 := Call(__e, PrimFunc(symshen_4find_1types), tmp4940)


__e.Return(PrimCons(tmp4938, tmp4941))
return


} else {
tmp4947 := PrimIsPair(V2258)

if True == tmp4947 {
tmp4942 := PrimHead(V2258)

tmp4943 := Call(__e, PrimFunc(symshen_4find_1types), tmp4942)


tmp4944 := PrimTail(V2258)

tmp4945 := Call(__e, PrimFunc(symshen_4find_1types), tmp4944)


__e.TailApply(PrimFunc(symappend), tmp4943, tmp4945)
return


} else {
__e.Return(Nil)
return
}


}


}, 1)

tmp4959 := Call(__e, ns2_1set, symshen_4find_1types, tmp4936)


_ = tmp4959

tmp4960 := MakeNative(func(__e *ControlFlow) {
V2261 := __e.Get(1)
_ = V2261
tmp5009 := PrimIsPair(V2261)

var ifres4990 Obj

if True == tmp5009 {
tmp5007 := PrimHead(V2261)

tmp5008 := PrimEqual(symdefine, tmp5007)

var ifres4992 Obj

if True == tmp5008 {
tmp5005 := PrimTail(V2261)

tmp5006 := PrimIsPair(tmp5005)

var ifres4994 Obj

if True == tmp5006 {
tmp5002 := PrimTail(V2261)

tmp5003 := PrimTail(tmp5002)

tmp5004 := PrimIsPair(tmp5003)

var ifres4996 Obj

if True == tmp5004 {
tmp4998 := PrimTail(V2261)

tmp4999 := PrimTail(tmp4998)

tmp5000 := PrimHead(tmp4999)

tmp5001 := PrimEqual(sym_i, tmp5000)

var ifres4997 Obj

if True == tmp5001 {
ifres4997 = True


} else {
ifres4997 = False


}

ifres4996 = ifres4997


} else {
ifres4996 = False


}

var ifres4995 Obj

if True == ifres4996 {
ifres4995 = True


} else {
ifres4995 = False


}

ifres4994 = ifres4995


} else {
ifres4994 = False


}

var ifres4993 Obj

if True == ifres4994 {
ifres4993 = True


} else {
ifres4993 = False


}

ifres4992 = ifres4993


} else {
ifres4992 = False


}

var ifres4991 Obj

if True == ifres4992 {
ifres4991 = True


} else {
ifres4991 = False


}

ifres4990 = ifres4991


} else {
ifres4990 = False


}

if True == ifres4990 {
tmp4961 := PrimTail(V2261)

tmp4962 := PrimHead(tmp4961)

tmp4963 := PrimTail(V2261)

tmp4964 := PrimHead(tmp4963)

tmp4965 := PrimTail(V2261)

tmp4966 := PrimTail(tmp4965)

tmp4967 := PrimTail(tmp4966)

tmp4968 := Call(__e, PrimFunc(symshen_4find_1arity), tmp4964, MakeNumber(1), tmp4967)


__e.TailApply(PrimFunc(symshen_4store_1arity), tmp4962, tmp4968)
return


} else {
tmp4988 := PrimIsPair(V2261)

var ifres4980 Obj

if True == tmp4988 {
tmp4986 := PrimHead(V2261)

tmp4987 := PrimEqual(symdefine, tmp4986)

var ifres4982 Obj

if True == tmp4987 {
tmp4984 := PrimTail(V2261)

tmp4985 := PrimIsPair(tmp4984)

var ifres4983 Obj

if True == tmp4985 {
ifres4983 = True


} else {
ifres4983 = False


}

ifres4982 = ifres4983


} else {
ifres4982 = False


}

var ifres4981 Obj

if True == ifres4982 {
ifres4981 = True


} else {
ifres4981 = False


}

ifres4980 = ifres4981


} else {
ifres4980 = False


}

if True == ifres4980 {
tmp4969 := PrimTail(V2261)

tmp4970 := PrimHead(tmp4969)

tmp4971 := PrimTail(V2261)

tmp4972 := PrimHead(tmp4971)

tmp4973 := PrimTail(V2261)

tmp4974 := PrimTail(tmp4973)

tmp4975 := Call(__e, PrimFunc(symshen_4find_1arity), tmp4972, MakeNumber(0), tmp4974)


__e.TailApply(PrimFunc(symshen_4store_1arity), tmp4970, tmp4975)
return


} else {
tmp4978 := PrimIsPair(V2261)

if True == tmp4978 {
tmp4976 := MakeNative(func(__e *ControlFlow) {
Z2262 := __e.Get(1)
_ = Z2262
__e.TailApply(PrimFunc(symshen_4find_1arities), Z2262)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp4976, V2261)
return


} else {
__e.Return(symshen_4skip)
return
}


}


}


}, 1)

tmp5010 := Call(__e, ns2_1set, symshen_4find_1arities, tmp4960)


_ = tmp5010

tmp5011 := MakeNative(func(__e *ControlFlow) {
V2263 := __e.Get(1)
_ = V2263
V2264 := __e.Get(2)
_ = V2264
tmp5012 := MakeNative(func(__e *ControlFlow) {
W2265 := __e.Get(1)
_ = W2265
tmp5023 := PrimEqual(W2265, MakeNumber(-1))

if True == tmp5023 {
__e.TailApply(PrimFunc(symshen_4execute_1store_1arity), V2263, V2264)
return
} else {
tmp5021 := PrimEqual(W2265, V2264)

if True == tmp5021 {
__e.Return(symshen_4skip)
return
} else {
tmp5019 := Call(__e, PrimFunc(symshen_4sysfunc_2), V2263)


if True == tmp5019 {
tmp5013 := Call(__e, PrimFunc(symshen_4app), V2263, MakeString(" is a system function\n"), symshen_4a)


__e.Return(PrimSimpleError(tmp5013))
return


} else {
tmp5014 := Call(__e, PrimFunc(symshen_4app), V2263, MakeString(" may cause errors\n"), symshen_4a)


tmp5015 := PrimStringConcat(MakeString("changing the arity of "), tmp5014)

tmp5016 := Call(__e, PrimFunc(symstoutput))


tmp5017 := Call(__e, PrimFunc(sympr), tmp5015, tmp5016)


_ = tmp5017

__e.TailApply(PrimFunc(symshen_4execute_1store_1arity), V2263, V2264)
return


}


}


}


}, 1)

tmp5024 := Call(__e, PrimFunc(symarity), V2263)


__e.TailApply(tmp5012, tmp5024)
return


}, 2)

tmp5025 := Call(__e, ns2_1set, symshen_4store_1arity, tmp5011)


_ = tmp5025

tmp5026 := MakeNative(func(__e *ControlFlow) {
V2266 := __e.Get(1)
_ = V2266
V2267 := __e.Get(2)
_ = V2267
tmp5031 := PrimEqual(MakeNumber(0), V2267)

if True == tmp5031 {
tmp5027 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V2266, symarity, MakeNumber(0), tmp5027)
return


} else {
tmp5028 := PrimValue(sym_dproperty_1vector_d)

tmp5029 := Call(__e, PrimFunc(symput), V2266, symarity, V2267, tmp5028)


_ = tmp5029

__e.TailApply(PrimFunc(symshen_4update_1lambdatable), V2266, V2267)
return


}


}, 2)

tmp5032 := Call(__e, ns2_1set, symshen_4execute_1store_1arity, tmp5026)


_ = tmp5032

tmp5033 := MakeNative(func(__e *ControlFlow) {
V2268 := __e.Get(1)
_ = V2268
V2269 := __e.Get(2)
_ = V2269
tmp5034 := MakeNative(func(__e *ControlFlow) {
W2270 := __e.Get(1)
_ = W2270
tmp5035 := MakeNative(func(__e *ControlFlow) {
W2271 := __e.Get(1)
_ = W2271
tmp5036 := MakeNative(func(__e *ControlFlow) {
W2272 := __e.Get(1)
_ = W2272
tmp5037 := MakeNative(func(__e *ControlFlow) {
W2273 := __e.Get(1)
_ = W2273
__e.Return(W2273)
return
}, 1)

tmp5038 := PrimSet(symshen_4_dlambdatable_d, W2272)

__e.TailApply(tmp5037, tmp5038)
return


}, 1)

tmp5039 := Call(__e, PrimFunc(symshen_4assoc_1_6), V2268, W2271, W2270)


__e.TailApply(tmp5036, tmp5039)
return


}, 1)

tmp5040 := PrimCons(V2268, Nil)

tmp5041 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp5040, V2269)


tmp5042 := Call(__e, PrimFunc(symeval_1kl), tmp5041)


__e.TailApply(tmp5035, tmp5042)
return


}, 1)

tmp5043 := PrimValue(symshen_4_dlambdatable_d)

__e.TailApply(tmp5034, tmp5043)
return


}, 2)

tmp5044 := Call(__e, ns2_1set, symshen_4update_1lambdatable, tmp5033)


_ = tmp5044

tmp5045 := MakeNative(func(__e *ControlFlow) {
V2276 := __e.Get(1)
_ = V2276
V2277 := __e.Get(2)
_ = V2277
tmp5063 := PrimEqual(MakeNumber(0), V2277)

if True == tmp5063 {
__e.Return(symshen_4skip)
return
} else {
tmp5061 := PrimEqual(MakeNumber(1), V2277)

if True == tmp5061 {
tmp5046 := MakeNative(func(__e *ControlFlow) {
W2278 := __e.Get(1)
_ = W2278
tmp5047 := PrimCons(W2278, Nil)

tmp5048 := Call(__e, PrimFunc(symappend), V2276, tmp5047)


tmp5049 := PrimCons(tmp5048, Nil)

tmp5050 := PrimCons(W2278, tmp5049)

__e.Return(PrimCons(symlambda, tmp5050))
return


}, 1)

tmp5051 := Call(__e, PrimFunc(symgensym), symY)


__e.TailApply(tmp5046, tmp5051)
return


} else {
tmp5052 := MakeNative(func(__e *ControlFlow) {
W2279 := __e.Get(1)
_ = W2279
tmp5053 := PrimCons(W2279, Nil)

tmp5054 := Call(__e, PrimFunc(symappend), V2276, tmp5053)


tmp5055 := PrimNumberSubtract(V2277, MakeNumber(1))

tmp5056 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp5054, tmp5055)


tmp5057 := PrimCons(tmp5056, Nil)

tmp5058 := PrimCons(W2279, tmp5057)

__e.Return(PrimCons(symlambda, tmp5058))
return


}, 1)

tmp5059 := Call(__e, PrimFunc(symgensym), symY)


__e.TailApply(tmp5052, tmp5059)
return


}


}


}, 2)

tmp5064 := Call(__e, ns2_1set, symshen_4lambda_1function, tmp5045)


_ = tmp5064

tmp5065 := MakeNative(func(__e *ControlFlow) {
V2289 := __e.Get(1)
_ = V2289
V2290 := __e.Get(2)
_ = V2290
V2291 := __e.Get(3)
_ = V2291
tmp5088 := PrimEqual(Nil, V2291)

if True == tmp5088 {
tmp5066 := PrimCons(V2289, V2290)

__e.Return(PrimCons(tmp5066, Nil))
return


} else {
tmp5086 := PrimIsPair(V2291)

var ifres5077 Obj

if True == tmp5086 {
tmp5084 := PrimHead(V2291)

tmp5085 := PrimIsPair(tmp5084)

var ifres5079 Obj

if True == tmp5085 {
tmp5081 := PrimHead(V2291)

tmp5082 := PrimHead(tmp5081)

tmp5083 := PrimEqual(V2289, tmp5082)

var ifres5080 Obj

if True == tmp5083 {
ifres5080 = True


} else {
ifres5080 = False


}

ifres5079 = ifres5080


} else {
ifres5079 = False


}

var ifres5078 Obj

if True == ifres5079 {
ifres5078 = True


} else {
ifres5078 = False


}

ifres5077 = ifres5078


} else {
ifres5077 = False


}

if True == ifres5077 {
tmp5067 := PrimHead(V2291)

tmp5068 := PrimHead(tmp5067)

tmp5069 := PrimCons(tmp5068, V2290)

tmp5070 := PrimTail(V2291)

__e.Return(PrimCons(tmp5069, tmp5070))
return


} else {
tmp5075 := PrimIsPair(V2291)

if True == tmp5075 {
tmp5071 := PrimHead(V2291)

tmp5072 := PrimTail(V2291)

tmp5073 := Call(__e, PrimFunc(symshen_4assoc_1_6), V2289, V2290, tmp5072)


__e.Return(PrimCons(tmp5071, tmp5073))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.assoc->")))
return
}


}


}


}, 3)

tmp5089 := Call(__e, ns2_1set, symshen_4assoc_1_6, tmp5065)


_ = tmp5089

tmp5090 := MakeNative(func(__e *ControlFlow) {
V2306 := __e.Get(1)
_ = V2306
V2307 := __e.Get(2)
_ = V2307
V2308 := __e.Get(3)
_ = V2308
tmp5137 := PrimEqual(MakeNumber(0), V2307)

var ifres5130 Obj

if True == tmp5137 {
tmp5136 := PrimIsPair(V2308)

var ifres5132 Obj

if True == tmp5136 {
tmp5134 := PrimHead(V2308)

tmp5135 := PrimEqual(tmp5134, sym_1_6)

var ifres5133 Obj

if True == tmp5135 {
ifres5133 = True


} else {
ifres5133 = False


}

ifres5132 = ifres5133


} else {
ifres5132 = False


}

var ifres5131 Obj

if True == ifres5132 {
ifres5131 = True


} else {
ifres5131 = False


}

ifres5130 = ifres5131


} else {
ifres5130 = False


}

if True == ifres5130 {
__e.Return(MakeNumber(0))
return
} else {
tmp5128 := PrimEqual(MakeNumber(0), V2307)

var ifres5121 Obj

if True == tmp5128 {
tmp5127 := PrimIsPair(V2308)

var ifres5123 Obj

if True == tmp5127 {
tmp5125 := PrimHead(V2308)

tmp5126 := PrimEqual(tmp5125, sym_5_1)

var ifres5124 Obj

if True == tmp5126 {
ifres5124 = True


} else {
ifres5124 = False


}

ifres5123 = ifres5124


} else {
ifres5123 = False


}

var ifres5122 Obj

if True == ifres5123 {
ifres5122 = True


} else {
ifres5122 = False


}

ifres5121 = ifres5122


} else {
ifres5121 = False


}

if True == ifres5121 {
__e.Return(MakeNumber(0))
return
} else {
tmp5119 := PrimEqual(MakeNumber(0), V2307)

var ifres5116 Obj

if True == tmp5119 {
tmp5118 := PrimIsPair(V2308)

var ifres5117 Obj

if True == tmp5118 {
ifres5117 = True


} else {
ifres5117 = False


}

ifres5116 = ifres5117


} else {
ifres5116 = False


}

if True == ifres5116 {
tmp5091 := PrimTail(V2308)

tmp5092 := Call(__e, PrimFunc(symshen_4find_1arity), V2306, MakeNumber(0), tmp5091)


__e.Return(PrimNumberAdd(MakeNumber(1), tmp5092))
return


} else {
tmp5114 := PrimEqual(MakeNumber(1), V2307)

var ifres5107 Obj

if True == tmp5114 {
tmp5113 := PrimIsPair(V2308)

var ifres5109 Obj

if True == tmp5113 {
tmp5111 := PrimHead(V2308)

tmp5112 := PrimEqual(sym_j, tmp5111)

var ifres5110 Obj

if True == tmp5112 {
ifres5110 = True


} else {
ifres5110 = False


}

ifres5109 = ifres5110


} else {
ifres5109 = False


}

var ifres5108 Obj

if True == ifres5109 {
ifres5108 = True


} else {
ifres5108 = False


}

ifres5107 = ifres5108


} else {
ifres5107 = False


}

if True == ifres5107 {
tmp5093 := PrimTail(V2308)

__e.TailApply(PrimFunc(symshen_4find_1arity), V2306, MakeNumber(0), tmp5093)
return


} else {
tmp5105 := PrimEqual(MakeNumber(1), V2307)

var ifres5102 Obj

if True == tmp5105 {
tmp5104 := PrimIsPair(V2308)

var ifres5103 Obj

if True == tmp5104 {
ifres5103 = True


} else {
ifres5103 = False


}

ifres5102 = ifres5103


} else {
ifres5102 = False


}

if True == ifres5102 {
tmp5094 := PrimTail(V2308)

__e.TailApply(PrimFunc(symshen_4find_1arity), V2306, MakeNumber(1), tmp5094)
return


} else {
tmp5100 := PrimEqual(MakeNumber(1), V2307)

if True == tmp5100 {
tmp5095 := Call(__e, PrimFunc(symshen_4app), V2306, MakeString(" definition: missing }\n"), symshen_4a)


tmp5096 := PrimStringConcat(MakeString("syntax error in "), tmp5095)

__e.Return(PrimSimpleError(tmp5096))
return


} else {
tmp5097 := Call(__e, PrimFunc(symshen_4app), V2306, MakeString(" definition: missing -> or <-\n"), symshen_4a)


tmp5098 := PrimStringConcat(MakeString("syntax error in "), tmp5097)

__e.Return(PrimSimpleError(tmp5098))
return


}


}


}


}


}


}


}, 3)

tmp5138 := Call(__e, ns2_1set, symshen_4find_1arity, tmp5090)


_ = tmp5138

tmp5139 := MakeNative(func(__e *ControlFlow) {
V2309 := __e.Get(1)
_ = V2309
tmp5140 := MakeNative(func(__e *ControlFlow) {
W2310 := __e.Get(1)
_ = W2310
tmp5385 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2310)


if True == tmp5385 {
tmp5141 := MakeNative(func(__e *ControlFlow) {
W2321 := __e.Get(1)
_ = W2321
tmp5353 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2321)


if True == tmp5353 {
tmp5142 := MakeNative(func(__e *ControlFlow) {
W2332 := __e.Get(1)
_ = W2332
tmp5335 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2332)


if True == tmp5335 {
tmp5143 := MakeNative(func(__e *ControlFlow) {
W2338 := __e.Get(1)
_ = W2338
tmp5317 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2338)


if True == tmp5317 {
tmp5144 := MakeNative(func(__e *ControlFlow) {
W2344 := __e.Get(1)
_ = W2344
tmp5299 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2344)


if True == tmp5299 {
tmp5145 := MakeNative(func(__e *ControlFlow) {
W2350 := __e.Get(1)
_ = W2350
tmp5280 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2350)


if True == tmp5280 {
tmp5146 := MakeNative(func(__e *ControlFlow) {
W2356 := __e.Get(1)
_ = W2356
tmp5255 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2356)


if True == tmp5255 {
tmp5147 := MakeNative(func(__e *ControlFlow) {
W2364 := __e.Get(1)
_ = W2364
tmp5236 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2364)


if True == tmp5236 {
tmp5148 := MakeNative(func(__e *ControlFlow) {
W2370 := __e.Get(1)
_ = W2370
tmp5217 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2370)


if True == tmp5217 {
tmp5149 := MakeNative(func(__e *ControlFlow) {
W2376 := __e.Get(1)
_ = W2376
tmp5200 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2376)


if True == tmp5200 {
tmp5150 := MakeNative(func(__e *ControlFlow) {
W2382 := __e.Get(1)
_ = W2382
tmp5180 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2382)


if True == tmp5180 {
tmp5151 := MakeNative(func(__e *ControlFlow) {
W2389 := __e.Get(1)
_ = W2389
tmp5163 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2389)


if True == tmp5163 {
tmp5152 := MakeNative(func(__e *ControlFlow) {
W2395 := __e.Get(1)
_ = W2395
tmp5154 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2395)


if True == tmp5154 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2395)
return
}


}, 1)

tmp5155 := MakeNative(func(__e *ControlFlow) {
W2396 := __e.Get(1)
_ = W2396
tmp5159 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2396)


if True == tmp5159 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5156 := MakeNative(func(__e *ControlFlow) {
W2397 := __e.Get(1)
_ = W2397
__e.TailApply(PrimFunc(symshen_4comb), W2397, Nil)
return
}, 1)

tmp5157 := Call(__e, PrimFunc(symshen_4in_1_6), W2396)


__e.TailApply(tmp5156, tmp5157)
return


}


}, 1)

tmp5160 := Call(__e, PrimFunc(sym_5e_6), V2309)


tmp5161 := Call(__e, tmp5155, tmp5160)


__e.TailApply(tmp5152, tmp5161)
return


} else {
__e.Return(W2389)
return
}


}, 1)

tmp5164 := MakeNative(func(__e *ControlFlow) {
W2390 := __e.Get(1)
_ = W2390
tmp5176 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2390)


if True == tmp5176 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5165 := MakeNative(func(__e *ControlFlow) {
W2391 := __e.Get(1)
_ = W2391
tmp5166 := MakeNative(func(__e *ControlFlow) {
W2392 := __e.Get(1)
_ = W2392
tmp5172 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2392)


if True == tmp5172 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5167 := MakeNative(func(__e *ControlFlow) {
W2393 := __e.Get(1)
_ = W2393
tmp5168 := MakeNative(func(__e *ControlFlow) {
W2394 := __e.Get(1)
_ = W2394
__e.TailApply(PrimFunc(symshen_4comb), W2394, W2393)
return
}, 1)

tmp5169 := Call(__e, PrimFunc(symshen_4in_1_6), W2392)


__e.TailApply(tmp5168, tmp5169)
return


}, 1)

tmp5170 := Call(__e, PrimFunc(symshen_4_5_1out), W2392)


__e.TailApply(tmp5167, tmp5170)
return


}


}, 1)

tmp5173 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2391)


__e.TailApply(tmp5166, tmp5173)
return


}, 1)

tmp5174 := Call(__e, PrimFunc(symshen_4in_1_6), W2390)


__e.TailApply(tmp5165, tmp5174)
return


}


}, 1)

tmp5177 := Call(__e, PrimFunc(symshen_4_5whitespaces_6), V2309)


tmp5178 := Call(__e, tmp5164, tmp5177)


__e.TailApply(tmp5151, tmp5178)
return


} else {
__e.Return(W2382)
return
}


}, 1)

tmp5181 := MakeNative(func(__e *ControlFlow) {
W2383 := __e.Get(1)
_ = W2383
tmp5196 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2383)


if True == tmp5196 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5182 := MakeNative(func(__e *ControlFlow) {
W2384 := __e.Get(1)
_ = W2384
tmp5183 := MakeNative(func(__e *ControlFlow) {
W2385 := __e.Get(1)
_ = W2385
tmp5184 := MakeNative(func(__e *ControlFlow) {
W2386 := __e.Get(1)
_ = W2386
tmp5191 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2386)


if True == tmp5191 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5185 := MakeNative(func(__e *ControlFlow) {
W2387 := __e.Get(1)
_ = W2387
tmp5186 := MakeNative(func(__e *ControlFlow) {
W2388 := __e.Get(1)
_ = W2388
tmp5187 := PrimCons(W2384, W2387)

__e.TailApply(PrimFunc(symshen_4comb), W2388, tmp5187)
return


}, 1)

tmp5188 := Call(__e, PrimFunc(symshen_4in_1_6), W2386)


__e.TailApply(tmp5186, tmp5188)
return


}, 1)

tmp5189 := Call(__e, PrimFunc(symshen_4_5_1out), W2386)


__e.TailApply(tmp5185, tmp5189)
return


}


}, 1)

tmp5192 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2385)


__e.TailApply(tmp5184, tmp5192)
return


}, 1)

tmp5193 := Call(__e, PrimFunc(symshen_4in_1_6), W2383)


__e.TailApply(tmp5183, tmp5193)
return


}, 1)

tmp5194 := Call(__e, PrimFunc(symshen_4_5_1out), W2383)


__e.TailApply(tmp5182, tmp5194)
return


}


}, 1)

tmp5197 := Call(__e, PrimFunc(symshen_4_5atom_6), V2309)


tmp5198 := Call(__e, tmp5181, tmp5197)


__e.TailApply(tmp5150, tmp5198)
return


} else {
__e.Return(W2376)
return
}


}, 1)

tmp5201 := MakeNative(func(__e *ControlFlow) {
W2377 := __e.Get(1)
_ = W2377
tmp5213 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2377)


if True == tmp5213 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5202 := MakeNative(func(__e *ControlFlow) {
W2378 := __e.Get(1)
_ = W2378
tmp5203 := MakeNative(func(__e *ControlFlow) {
W2379 := __e.Get(1)
_ = W2379
tmp5209 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2379)


if True == tmp5209 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5204 := MakeNative(func(__e *ControlFlow) {
W2380 := __e.Get(1)
_ = W2380
tmp5205 := MakeNative(func(__e *ControlFlow) {
W2381 := __e.Get(1)
_ = W2381
__e.TailApply(PrimFunc(symshen_4comb), W2381, W2380)
return
}, 1)

tmp5206 := Call(__e, PrimFunc(symshen_4in_1_6), W2379)


__e.TailApply(tmp5205, tmp5206)
return


}, 1)

tmp5207 := Call(__e, PrimFunc(symshen_4_5_1out), W2379)


__e.TailApply(tmp5204, tmp5207)
return


}


}, 1)

tmp5210 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2378)


__e.TailApply(tmp5203, tmp5210)
return


}, 1)

tmp5211 := Call(__e, PrimFunc(symshen_4in_1_6), W2377)


__e.TailApply(tmp5202, tmp5211)
return


}


}, 1)

tmp5214 := Call(__e, PrimFunc(symshen_4_5comment_6), V2309)


tmp5215 := Call(__e, tmp5201, tmp5214)


__e.TailApply(tmp5149, tmp5215)
return


} else {
__e.Return(W2370)
return
}


}, 1)

tmp5218 := MakeNative(func(__e *ControlFlow) {
W2371 := __e.Get(1)
_ = W2371
tmp5232 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2371)


if True == tmp5232 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5219 := MakeNative(func(__e *ControlFlow) {
W2372 := __e.Get(1)
_ = W2372
tmp5220 := MakeNative(func(__e *ControlFlow) {
W2373 := __e.Get(1)
_ = W2373
tmp5228 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2373)


if True == tmp5228 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5221 := MakeNative(func(__e *ControlFlow) {
W2374 := __e.Get(1)
_ = W2374
tmp5222 := MakeNative(func(__e *ControlFlow) {
W2375 := __e.Get(1)
_ = W2375
tmp5223 := PrimIntern(MakeString(","))

tmp5224 := PrimCons(tmp5223, W2374)

__e.TailApply(PrimFunc(symshen_4comb), W2375, tmp5224)
return


}, 1)

tmp5225 := Call(__e, PrimFunc(symshen_4in_1_6), W2373)


__e.TailApply(tmp5222, tmp5225)
return


}, 1)

tmp5226 := Call(__e, PrimFunc(symshen_4_5_1out), W2373)


__e.TailApply(tmp5221, tmp5226)
return


}


}, 1)

tmp5229 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2372)


__e.TailApply(tmp5220, tmp5229)
return


}, 1)

tmp5230 := Call(__e, PrimFunc(symshen_4in_1_6), W2371)


__e.TailApply(tmp5219, tmp5230)
return


}


}, 1)

tmp5233 := Call(__e, PrimFunc(symshen_4_5comma_6), V2309)


tmp5234 := Call(__e, tmp5218, tmp5233)


__e.TailApply(tmp5148, tmp5234)
return


} else {
__e.Return(W2364)
return
}


}, 1)

tmp5237 := MakeNative(func(__e *ControlFlow) {
W2365 := __e.Get(1)
_ = W2365
tmp5251 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2365)


if True == tmp5251 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5238 := MakeNative(func(__e *ControlFlow) {
W2366 := __e.Get(1)
_ = W2366
tmp5239 := MakeNative(func(__e *ControlFlow) {
W2367 := __e.Get(1)
_ = W2367
tmp5247 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2367)


if True == tmp5247 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5240 := MakeNative(func(__e *ControlFlow) {
W2368 := __e.Get(1)
_ = W2368
tmp5241 := MakeNative(func(__e *ControlFlow) {
W2369 := __e.Get(1)
_ = W2369
tmp5242 := PrimIntern(MakeString(":"))

tmp5243 := PrimCons(tmp5242, W2368)

__e.TailApply(PrimFunc(symshen_4comb), W2369, tmp5243)
return


}, 1)

tmp5244 := Call(__e, PrimFunc(symshen_4in_1_6), W2367)


__e.TailApply(tmp5241, tmp5244)
return


}, 1)

tmp5245 := Call(__e, PrimFunc(symshen_4_5_1out), W2367)


__e.TailApply(tmp5240, tmp5245)
return


}


}, 1)

tmp5248 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2366)


__e.TailApply(tmp5239, tmp5248)
return


}, 1)

tmp5249 := Call(__e, PrimFunc(symshen_4in_1_6), W2365)


__e.TailApply(tmp5238, tmp5249)
return


}


}, 1)

tmp5252 := Call(__e, PrimFunc(symshen_4_5colon_6), V2309)


tmp5253 := Call(__e, tmp5237, tmp5252)


__e.TailApply(tmp5147, tmp5253)
return


} else {
__e.Return(W2356)
return
}


}, 1)

tmp5256 := MakeNative(func(__e *ControlFlow) {
W2357 := __e.Get(1)
_ = W2357
tmp5276 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2357)


if True == tmp5276 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5257 := MakeNative(func(__e *ControlFlow) {
W2358 := __e.Get(1)
_ = W2358
tmp5258 := MakeNative(func(__e *ControlFlow) {
W2359 := __e.Get(1)
_ = W2359
tmp5272 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2359)


if True == tmp5272 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5259 := MakeNative(func(__e *ControlFlow) {
W2360 := __e.Get(1)
_ = W2360
tmp5260 := MakeNative(func(__e *ControlFlow) {
W2361 := __e.Get(1)
_ = W2361
tmp5268 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2361)


if True == tmp5268 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5261 := MakeNative(func(__e *ControlFlow) {
W2362 := __e.Get(1)
_ = W2362
tmp5262 := MakeNative(func(__e *ControlFlow) {
W2363 := __e.Get(1)
_ = W2363
tmp5263 := PrimIntern(MakeString(":="))

tmp5264 := PrimCons(tmp5263, W2362)

__e.TailApply(PrimFunc(symshen_4comb), W2363, tmp5264)
return


}, 1)

tmp5265 := Call(__e, PrimFunc(symshen_4in_1_6), W2361)


__e.TailApply(tmp5262, tmp5265)
return


}, 1)

tmp5266 := Call(__e, PrimFunc(symshen_4_5_1out), W2361)


__e.TailApply(tmp5261, tmp5266)
return


}


}, 1)

tmp5269 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2360)


__e.TailApply(tmp5260, tmp5269)
return


}, 1)

tmp5270 := Call(__e, PrimFunc(symshen_4in_1_6), W2359)


__e.TailApply(tmp5259, tmp5270)
return


}


}, 1)

tmp5273 := Call(__e, PrimFunc(symshen_4_5equal_6), W2358)


__e.TailApply(tmp5258, tmp5273)
return


}, 1)

tmp5274 := Call(__e, PrimFunc(symshen_4in_1_6), W2357)


__e.TailApply(tmp5257, tmp5274)
return


}


}, 1)

tmp5277 := Call(__e, PrimFunc(symshen_4_5colon_6), V2309)


tmp5278 := Call(__e, tmp5256, tmp5277)


__e.TailApply(tmp5146, tmp5278)
return


} else {
__e.Return(W2350)
return
}


}, 1)

tmp5281 := MakeNative(func(__e *ControlFlow) {
W2351 := __e.Get(1)
_ = W2351
tmp5295 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2351)


if True == tmp5295 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5282 := MakeNative(func(__e *ControlFlow) {
W2352 := __e.Get(1)
_ = W2352
tmp5283 := MakeNative(func(__e *ControlFlow) {
W2353 := __e.Get(1)
_ = W2353
tmp5291 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2353)


if True == tmp5291 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5284 := MakeNative(func(__e *ControlFlow) {
W2354 := __e.Get(1)
_ = W2354
tmp5285 := MakeNative(func(__e *ControlFlow) {
W2355 := __e.Get(1)
_ = W2355
tmp5286 := PrimIntern(MakeString(";"))

tmp5287 := PrimCons(tmp5286, W2354)

__e.TailApply(PrimFunc(symshen_4comb), W2355, tmp5287)
return


}, 1)

tmp5288 := Call(__e, PrimFunc(symshen_4in_1_6), W2353)


__e.TailApply(tmp5285, tmp5288)
return


}, 1)

tmp5289 := Call(__e, PrimFunc(symshen_4_5_1out), W2353)


__e.TailApply(tmp5284, tmp5289)
return


}


}, 1)

tmp5292 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2352)


__e.TailApply(tmp5283, tmp5292)
return


}, 1)

tmp5293 := Call(__e, PrimFunc(symshen_4in_1_6), W2351)


__e.TailApply(tmp5282, tmp5293)
return


}


}, 1)

tmp5296 := Call(__e, PrimFunc(symshen_4_5semicolon_6), V2309)


tmp5297 := Call(__e, tmp5281, tmp5296)


__e.TailApply(tmp5145, tmp5297)
return


} else {
__e.Return(W2344)
return
}


}, 1)

tmp5300 := MakeNative(func(__e *ControlFlow) {
W2345 := __e.Get(1)
_ = W2345
tmp5313 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2345)


if True == tmp5313 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5301 := MakeNative(func(__e *ControlFlow) {
W2346 := __e.Get(1)
_ = W2346
tmp5302 := MakeNative(func(__e *ControlFlow) {
W2347 := __e.Get(1)
_ = W2347
tmp5309 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2347)


if True == tmp5309 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5303 := MakeNative(func(__e *ControlFlow) {
W2348 := __e.Get(1)
_ = W2348
tmp5304 := MakeNative(func(__e *ControlFlow) {
W2349 := __e.Get(1)
_ = W2349
tmp5305 := PrimCons(symbar_b, W2348)

__e.TailApply(PrimFunc(symshen_4comb), W2349, tmp5305)
return


}, 1)

tmp5306 := Call(__e, PrimFunc(symshen_4in_1_6), W2347)


__e.TailApply(tmp5304, tmp5306)
return


}, 1)

tmp5307 := Call(__e, PrimFunc(symshen_4_5_1out), W2347)


__e.TailApply(tmp5303, tmp5307)
return


}


}, 1)

tmp5310 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2346)


__e.TailApply(tmp5302, tmp5310)
return


}, 1)

tmp5311 := Call(__e, PrimFunc(symshen_4in_1_6), W2345)


__e.TailApply(tmp5301, tmp5311)
return


}


}, 1)

tmp5314 := Call(__e, PrimFunc(symshen_4_5bar_6), V2309)


tmp5315 := Call(__e, tmp5300, tmp5314)


__e.TailApply(tmp5144, tmp5315)
return


} else {
__e.Return(W2338)
return
}


}, 1)

tmp5318 := MakeNative(func(__e *ControlFlow) {
W2339 := __e.Get(1)
_ = W2339
tmp5331 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2339)


if True == tmp5331 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5319 := MakeNative(func(__e *ControlFlow) {
W2340 := __e.Get(1)
_ = W2340
tmp5320 := MakeNative(func(__e *ControlFlow) {
W2341 := __e.Get(1)
_ = W2341
tmp5327 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2341)


if True == tmp5327 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5321 := MakeNative(func(__e *ControlFlow) {
W2342 := __e.Get(1)
_ = W2342
tmp5322 := MakeNative(func(__e *ControlFlow) {
W2343 := __e.Get(1)
_ = W2343
tmp5323 := PrimCons(sym_j, W2342)

__e.TailApply(PrimFunc(symshen_4comb), W2343, tmp5323)
return


}, 1)

tmp5324 := Call(__e, PrimFunc(symshen_4in_1_6), W2341)


__e.TailApply(tmp5322, tmp5324)
return


}, 1)

tmp5325 := Call(__e, PrimFunc(symshen_4_5_1out), W2341)


__e.TailApply(tmp5321, tmp5325)
return


}


}, 1)

tmp5328 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2340)


__e.TailApply(tmp5320, tmp5328)
return


}, 1)

tmp5329 := Call(__e, PrimFunc(symshen_4in_1_6), W2339)


__e.TailApply(tmp5319, tmp5329)
return


}


}, 1)

tmp5332 := Call(__e, PrimFunc(symshen_4_5rcurly_6), V2309)


tmp5333 := Call(__e, tmp5318, tmp5332)


__e.TailApply(tmp5143, tmp5333)
return


} else {
__e.Return(W2332)
return
}


}, 1)

tmp5336 := MakeNative(func(__e *ControlFlow) {
W2333 := __e.Get(1)
_ = W2333
tmp5349 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2333)


if True == tmp5349 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5337 := MakeNative(func(__e *ControlFlow) {
W2334 := __e.Get(1)
_ = W2334
tmp5338 := MakeNative(func(__e *ControlFlow) {
W2335 := __e.Get(1)
_ = W2335
tmp5345 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2335)


if True == tmp5345 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5339 := MakeNative(func(__e *ControlFlow) {
W2336 := __e.Get(1)
_ = W2336
tmp5340 := MakeNative(func(__e *ControlFlow) {
W2337 := __e.Get(1)
_ = W2337
tmp5341 := PrimCons(sym_i, W2336)

__e.TailApply(PrimFunc(symshen_4comb), W2337, tmp5341)
return


}, 1)

tmp5342 := Call(__e, PrimFunc(symshen_4in_1_6), W2335)


__e.TailApply(tmp5340, tmp5342)
return


}, 1)

tmp5343 := Call(__e, PrimFunc(symshen_4_5_1out), W2335)


__e.TailApply(tmp5339, tmp5343)
return


}


}, 1)

tmp5346 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2334)


__e.TailApply(tmp5338, tmp5346)
return


}, 1)

tmp5347 := Call(__e, PrimFunc(symshen_4in_1_6), W2333)


__e.TailApply(tmp5337, tmp5347)
return


}


}, 1)

tmp5350 := Call(__e, PrimFunc(symshen_4_5lcurly_6), V2309)


tmp5351 := Call(__e, tmp5336, tmp5350)


__e.TailApply(tmp5142, tmp5351)
return


} else {
__e.Return(W2321)
return
}


}, 1)

tmp5354 := MakeNative(func(__e *ControlFlow) {
W2322 := __e.Get(1)
_ = W2322
tmp5381 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2322)


if True == tmp5381 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5355 := MakeNative(func(__e *ControlFlow) {
W2323 := __e.Get(1)
_ = W2323
tmp5356 := MakeNative(func(__e *ControlFlow) {
W2324 := __e.Get(1)
_ = W2324
tmp5377 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2324)


if True == tmp5377 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5357 := MakeNative(func(__e *ControlFlow) {
W2325 := __e.Get(1)
_ = W2325
tmp5358 := MakeNative(func(__e *ControlFlow) {
W2326 := __e.Get(1)
_ = W2326
tmp5359 := MakeNative(func(__e *ControlFlow) {
W2327 := __e.Get(1)
_ = W2327
tmp5372 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2327)


if True == tmp5372 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5360 := MakeNative(func(__e *ControlFlow) {
W2328 := __e.Get(1)
_ = W2328
tmp5361 := MakeNative(func(__e *ControlFlow) {
W2329 := __e.Get(1)
_ = W2329
tmp5368 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2329)


if True == tmp5368 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5362 := MakeNative(func(__e *ControlFlow) {
W2330 := __e.Get(1)
_ = W2330
tmp5363 := MakeNative(func(__e *ControlFlow) {
W2331 := __e.Get(1)
_ = W2331
tmp5364 := Call(__e, PrimFunc(symshen_4add_1sexpr), W2325, W2330)


__e.TailApply(PrimFunc(symshen_4comb), W2331, tmp5364)
return


}, 1)

tmp5365 := Call(__e, PrimFunc(symshen_4in_1_6), W2329)


__e.TailApply(tmp5363, tmp5365)
return


}, 1)

tmp5366 := Call(__e, PrimFunc(symshen_4_5_1out), W2329)


__e.TailApply(tmp5362, tmp5366)
return


}


}, 1)

tmp5369 := Call(__e, PrimFunc(symshen_4_5s_1exprs2_6), W2328)


__e.TailApply(tmp5361, tmp5369)
return


}, 1)

tmp5370 := Call(__e, PrimFunc(symshen_4in_1_6), W2327)


__e.TailApply(tmp5360, tmp5370)
return


}


}, 1)

tmp5373 := Call(__e, PrimFunc(symshen_4_5rrb_6), W2326)


__e.TailApply(tmp5359, tmp5373)
return


}, 1)

tmp5374 := Call(__e, PrimFunc(symshen_4in_1_6), W2324)


__e.TailApply(tmp5358, tmp5374)
return


}, 1)

tmp5375 := Call(__e, PrimFunc(symshen_4_5_1out), W2324)


__e.TailApply(tmp5357, tmp5375)
return


}


}, 1)

tmp5378 := Call(__e, PrimFunc(symshen_4_5s_1exprs1_6), W2323)


__e.TailApply(tmp5356, tmp5378)
return


}, 1)

tmp5379 := Call(__e, PrimFunc(symshen_4in_1_6), W2322)


__e.TailApply(tmp5355, tmp5379)
return


}


}, 1)

tmp5382 := Call(__e, PrimFunc(symshen_4_5lrb_6), V2309)


tmp5383 := Call(__e, tmp5354, tmp5382)


__e.TailApply(tmp5141, tmp5383)
return


} else {
__e.Return(W2310)
return
}


}, 1)

tmp5386 := MakeNative(func(__e *ControlFlow) {
W2311 := __e.Get(1)
_ = W2311
tmp5414 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2311)


if True == tmp5414 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5387 := MakeNative(func(__e *ControlFlow) {
W2312 := __e.Get(1)
_ = W2312
tmp5388 := MakeNative(func(__e *ControlFlow) {
W2313 := __e.Get(1)
_ = W2313
tmp5410 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2313)


if True == tmp5410 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5389 := MakeNative(func(__e *ControlFlow) {
W2314 := __e.Get(1)
_ = W2314
tmp5390 := MakeNative(func(__e *ControlFlow) {
W2315 := __e.Get(1)
_ = W2315
tmp5391 := MakeNative(func(__e *ControlFlow) {
W2316 := __e.Get(1)
_ = W2316
tmp5405 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2316)


if True == tmp5405 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5392 := MakeNative(func(__e *ControlFlow) {
W2317 := __e.Get(1)
_ = W2317
tmp5393 := MakeNative(func(__e *ControlFlow) {
W2318 := __e.Get(1)
_ = W2318
tmp5401 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2318)


if True == tmp5401 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5394 := MakeNative(func(__e *ControlFlow) {
W2319 := __e.Get(1)
_ = W2319
tmp5395 := MakeNative(func(__e *ControlFlow) {
W2320 := __e.Get(1)
_ = W2320
tmp5396 := Call(__e, PrimFunc(symshen_4cons_1form), W2314)


tmp5397 := PrimCons(tmp5396, W2319)

__e.TailApply(PrimFunc(symshen_4comb), W2320, tmp5397)
return


}, 1)

tmp5398 := Call(__e, PrimFunc(symshen_4in_1_6), W2318)


__e.TailApply(tmp5395, tmp5398)
return


}, 1)

tmp5399 := Call(__e, PrimFunc(symshen_4_5_1out), W2318)


__e.TailApply(tmp5394, tmp5399)
return


}


}, 1)

tmp5402 := Call(__e, PrimFunc(symshen_4_5s_1exprs2_6), W2317)


__e.TailApply(tmp5393, tmp5402)
return


}, 1)

tmp5403 := Call(__e, PrimFunc(symshen_4in_1_6), W2316)


__e.TailApply(tmp5392, tmp5403)
return


}


}, 1)

tmp5406 := Call(__e, PrimFunc(symshen_4_5rsb_6), W2315)


__e.TailApply(tmp5391, tmp5406)
return


}, 1)

tmp5407 := Call(__e, PrimFunc(symshen_4in_1_6), W2313)


__e.TailApply(tmp5390, tmp5407)
return


}, 1)

tmp5408 := Call(__e, PrimFunc(symshen_4_5_1out), W2313)


__e.TailApply(tmp5389, tmp5408)
return


}


}, 1)

tmp5411 := Call(__e, PrimFunc(symshen_4_5s_1exprs1_6), W2312)


__e.TailApply(tmp5388, tmp5411)
return


}, 1)

tmp5412 := Call(__e, PrimFunc(symshen_4in_1_6), W2311)


__e.TailApply(tmp5387, tmp5412)
return


}


}, 1)

tmp5415 := Call(__e, PrimFunc(symshen_4_5lsb_6), V2309)


tmp5416 := Call(__e, tmp5386, tmp5415)


__e.TailApply(tmp5140, tmp5416)
return


}, 1)

tmp5417 := Call(__e, ns2_1set, symshen_4_5s_1exprs_6, tmp5139)


_ = tmp5417

tmp5418 := MakeNative(func(__e *ControlFlow) {
V2398 := __e.Get(1)
_ = V2398
V2399 := __e.Get(2)
_ = V2399
tmp5436 := PrimIsPair(V2398)

var ifres5423 Obj

if True == tmp5436 {
tmp5434 := PrimHead(V2398)

tmp5435 := PrimEqual(sym_3, tmp5434)

var ifres5425 Obj

if True == tmp5435 {
tmp5432 := PrimTail(V2398)

tmp5433 := PrimIsPair(tmp5432)

var ifres5427 Obj

if True == tmp5433 {
tmp5429 := PrimTail(V2398)

tmp5430 := PrimTail(tmp5429)

tmp5431 := PrimEqual(Nil, tmp5430)

var ifres5428 Obj

if True == tmp5431 {
ifres5428 = True


} else {
ifres5428 = False


}

ifres5427 = ifres5428


} else {
ifres5427 = False


}

var ifres5426 Obj

if True == ifres5427 {
ifres5426 = True


} else {
ifres5426 = False


}

ifres5425 = ifres5426


} else {
ifres5425 = False


}

var ifres5424 Obj

if True == ifres5425 {
ifres5424 = True


} else {
ifres5424 = False


}

ifres5423 = ifres5424


} else {
ifres5423 = False


}

if True == ifres5423 {
tmp5419 := PrimTail(V2398)

tmp5420 := PrimHead(tmp5419)

tmp5421 := Call(__e, PrimFunc(symexplode), tmp5420)


__e.TailApply(PrimFunc(symappend), tmp5421, V2399)
return


} else {
__e.Return(PrimCons(V2398, V2399))
return
}


}, 2)

tmp5437 := Call(__e, ns2_1set, symshen_4add_1sexpr, tmp5418)


_ = tmp5437

tmp5438 := MakeNative(func(__e *ControlFlow) {
V2400 := __e.Get(1)
_ = V2400
tmp5439 := MakeNative(func(__e *ControlFlow) {
W2401 := __e.Get(1)
_ = W2401
tmp5441 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2401)


if True == tmp5441 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2401)
return
}


}, 1)

tmp5447 := Call(__e, PrimFunc(symshen_4hds_a_2), V2400, MakeNumber(91))


var ifres5442 Obj

if True == tmp5447 {
tmp5443 := MakeNative(func(__e *ControlFlow) {
W2402 := __e.Get(1)
_ = W2402
__e.TailApply(PrimFunc(symshen_4comb), W2402, symshen_4skip)
return
}, 1)

tmp5444 := Call(__e, PrimFunc(symtail), V2400)


tmp5445 := Call(__e, tmp5443, tmp5444)


ifres5442 = tmp5445


} else {
tmp5446 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5442 = tmp5446


}

__e.TailApply(tmp5439, ifres5442)
return


}, 1)

tmp5448 := Call(__e, ns2_1set, symshen_4_5lsb_6, tmp5438)


_ = tmp5448

tmp5449 := MakeNative(func(__e *ControlFlow) {
V2403 := __e.Get(1)
_ = V2403
tmp5450 := MakeNative(func(__e *ControlFlow) {
W2404 := __e.Get(1)
_ = W2404
tmp5452 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2404)


if True == tmp5452 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2404)
return
}


}, 1)

tmp5458 := Call(__e, PrimFunc(symshen_4hds_a_2), V2403, MakeNumber(93))


var ifres5453 Obj

if True == tmp5458 {
tmp5454 := MakeNative(func(__e *ControlFlow) {
W2405 := __e.Get(1)
_ = W2405
__e.TailApply(PrimFunc(symshen_4comb), W2405, symshen_4skip)
return
}, 1)

tmp5455 := Call(__e, PrimFunc(symtail), V2403)


tmp5456 := Call(__e, tmp5454, tmp5455)


ifres5453 = tmp5456


} else {
tmp5457 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5453 = tmp5457


}

__e.TailApply(tmp5450, ifres5453)
return


}, 1)

tmp5459 := Call(__e, ns2_1set, symshen_4_5rsb_6, tmp5449)


_ = tmp5459

tmp5460 := MakeNative(func(__e *ControlFlow) {
V2406 := __e.Get(1)
_ = V2406
tmp5461 := MakeNative(func(__e *ControlFlow) {
W2407 := __e.Get(1)
_ = W2407
tmp5463 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2407)


if True == tmp5463 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2407)
return
}


}, 1)

tmp5464 := MakeNative(func(__e *ControlFlow) {
W2408 := __e.Get(1)
_ = W2408
tmp5470 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2408)


if True == tmp5470 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5465 := MakeNative(func(__e *ControlFlow) {
W2409 := __e.Get(1)
_ = W2409
tmp5466 := MakeNative(func(__e *ControlFlow) {
W2410 := __e.Get(1)
_ = W2410
__e.TailApply(PrimFunc(symshen_4comb), W2410, W2409)
return
}, 1)

tmp5467 := Call(__e, PrimFunc(symshen_4in_1_6), W2408)


__e.TailApply(tmp5466, tmp5467)
return


}, 1)

tmp5468 := Call(__e, PrimFunc(symshen_4_5_1out), W2408)


__e.TailApply(tmp5465, tmp5468)
return


}


}, 1)

tmp5471 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), V2406)


tmp5472 := Call(__e, tmp5464, tmp5471)


__e.TailApply(tmp5461, tmp5472)
return


}, 1)

tmp5473 := Call(__e, ns2_1set, symshen_4_5s_1exprs1_6, tmp5460)


_ = tmp5473

tmp5474 := MakeNative(func(__e *ControlFlow) {
V2411 := __e.Get(1)
_ = V2411
tmp5475 := MakeNative(func(__e *ControlFlow) {
W2412 := __e.Get(1)
_ = W2412
tmp5477 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2412)


if True == tmp5477 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2412)
return
}


}, 1)

tmp5478 := MakeNative(func(__e *ControlFlow) {
W2413 := __e.Get(1)
_ = W2413
tmp5484 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2413)


if True == tmp5484 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5479 := MakeNative(func(__e *ControlFlow) {
W2414 := __e.Get(1)
_ = W2414
tmp5480 := MakeNative(func(__e *ControlFlow) {
W2415 := __e.Get(1)
_ = W2415
__e.TailApply(PrimFunc(symshen_4comb), W2415, W2414)
return
}, 1)

tmp5481 := Call(__e, PrimFunc(symshen_4in_1_6), W2413)


__e.TailApply(tmp5480, tmp5481)
return


}, 1)

tmp5482 := Call(__e, PrimFunc(symshen_4_5_1out), W2413)


__e.TailApply(tmp5479, tmp5482)
return


}


}, 1)

tmp5485 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), V2411)


tmp5486 := Call(__e, tmp5478, tmp5485)


__e.TailApply(tmp5475, tmp5486)
return


}, 1)

tmp5487 := Call(__e, ns2_1set, symshen_4_5s_1exprs2_6, tmp5474)


_ = tmp5487

tmp5488 := MakeNative(func(__e *ControlFlow) {
V2417 := __e.Get(1)
_ = V2417
tmp5545 := PrimEqual(Nil, V2417)

if True == tmp5545 {
__e.Return(Nil)
return
} else {
tmp5543 := PrimIsPair(V2417)

var ifres5523 Obj

if True == tmp5543 {
tmp5541 := PrimTail(V2417)

tmp5542 := PrimIsPair(tmp5541)

var ifres5525 Obj

if True == tmp5542 {
tmp5538 := PrimTail(V2417)

tmp5539 := PrimTail(tmp5538)

tmp5540 := PrimIsPair(tmp5539)

var ifres5527 Obj

if True == tmp5540 {
tmp5534 := PrimTail(V2417)

tmp5535 := PrimTail(tmp5534)

tmp5536 := PrimTail(tmp5535)

tmp5537 := PrimEqual(Nil, tmp5536)

var ifres5529 Obj

if True == tmp5537 {
tmp5531 := PrimTail(V2417)

tmp5532 := PrimHead(tmp5531)

tmp5533 := PrimEqual(tmp5532, symbar_b)

var ifres5530 Obj

if True == tmp5533 {
ifres5530 = True


} else {
ifres5530 = False


}

ifres5529 = ifres5530


} else {
ifres5529 = False


}

var ifres5528 Obj

if True == ifres5529 {
ifres5528 = True


} else {
ifres5528 = False


}

ifres5527 = ifres5528


} else {
ifres5527 = False


}

var ifres5526 Obj

if True == ifres5527 {
ifres5526 = True


} else {
ifres5526 = False


}

ifres5525 = ifres5526


} else {
ifres5525 = False


}

var ifres5524 Obj

if True == ifres5525 {
ifres5524 = True


} else {
ifres5524 = False


}

ifres5523 = ifres5524


} else {
ifres5523 = False


}

if True == ifres5523 {
tmp5489 := PrimHead(V2417)

tmp5490 := PrimTail(V2417)

tmp5491 := PrimTail(tmp5490)

tmp5492 := PrimCons(tmp5489, tmp5491)

__e.Return(PrimCons(symcons, tmp5492))
return


} else {
tmp5521 := PrimIsPair(V2417)

var ifres5501 Obj

if True == tmp5521 {
tmp5519 := PrimTail(V2417)

tmp5520 := PrimIsPair(tmp5519)

var ifres5503 Obj

if True == tmp5520 {
tmp5516 := PrimTail(V2417)

tmp5517 := PrimTail(tmp5516)

tmp5518 := PrimIsPair(tmp5517)

var ifres5505 Obj

if True == tmp5518 {
tmp5512 := PrimTail(V2417)

tmp5513 := PrimTail(tmp5512)

tmp5514 := PrimTail(tmp5513)

tmp5515 := PrimIsPair(tmp5514)

var ifres5507 Obj

if True == tmp5515 {
tmp5509 := PrimTail(V2417)

tmp5510 := PrimHead(tmp5509)

tmp5511 := PrimEqual(tmp5510, symbar_b)

var ifres5508 Obj

if True == tmp5511 {
ifres5508 = True


} else {
ifres5508 = False


}

ifres5507 = ifres5508


} else {
ifres5507 = False


}

var ifres5506 Obj

if True == ifres5507 {
ifres5506 = True


} else {
ifres5506 = False


}

ifres5505 = ifres5506


} else {
ifres5505 = False


}

var ifres5504 Obj

if True == ifres5505 {
ifres5504 = True


} else {
ifres5504 = False


}

ifres5503 = ifres5504


} else {
ifres5503 = False


}

var ifres5502 Obj

if True == ifres5503 {
ifres5502 = True


} else {
ifres5502 = False


}

ifres5501 = ifres5502


} else {
ifres5501 = False


}

if True == ifres5501 {
__e.Return(PrimSimpleError(MakeString("misapplication of |\n")))
return
} else {
tmp5499 := PrimIsPair(V2417)

if True == tmp5499 {
tmp5493 := PrimHead(V2417)

tmp5494 := PrimTail(V2417)

tmp5495 := Call(__e, PrimFunc(symshen_4cons_1form), tmp5494)


tmp5496 := PrimCons(tmp5495, Nil)

tmp5497 := PrimCons(tmp5493, tmp5496)

__e.Return(PrimCons(symcons, tmp5497))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.cons-form")))
return
}


}


}


}


}, 1)

tmp5546 := Call(__e, ns2_1set, symshen_4cons_1form, tmp5488)


_ = tmp5546

tmp5547 := MakeNative(func(__e *ControlFlow) {
V2418 := __e.Get(1)
_ = V2418
tmp5548 := MakeNative(func(__e *ControlFlow) {
W2419 := __e.Get(1)
_ = W2419
tmp5550 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2419)


if True == tmp5550 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2419)
return
}


}, 1)

tmp5556 := Call(__e, PrimFunc(symshen_4hds_a_2), V2418, MakeNumber(40))


var ifres5551 Obj

if True == tmp5556 {
tmp5552 := MakeNative(func(__e *ControlFlow) {
W2420 := __e.Get(1)
_ = W2420
__e.TailApply(PrimFunc(symshen_4comb), W2420, symshen_4skip)
return
}, 1)

tmp5553 := Call(__e, PrimFunc(symtail), V2418)


tmp5554 := Call(__e, tmp5552, tmp5553)


ifres5551 = tmp5554


} else {
tmp5555 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5551 = tmp5555


}

__e.TailApply(tmp5548, ifres5551)
return


}, 1)

tmp5557 := Call(__e, ns2_1set, symshen_4_5lrb_6, tmp5547)


_ = tmp5557

tmp5558 := MakeNative(func(__e *ControlFlow) {
V2421 := __e.Get(1)
_ = V2421
tmp5559 := MakeNative(func(__e *ControlFlow) {
W2422 := __e.Get(1)
_ = W2422
tmp5561 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2422)


if True == tmp5561 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2422)
return
}


}, 1)

tmp5567 := Call(__e, PrimFunc(symshen_4hds_a_2), V2421, MakeNumber(41))


var ifres5562 Obj

if True == tmp5567 {
tmp5563 := MakeNative(func(__e *ControlFlow) {
W2423 := __e.Get(1)
_ = W2423
__e.TailApply(PrimFunc(symshen_4comb), W2423, symshen_4skip)
return
}, 1)

tmp5564 := Call(__e, PrimFunc(symtail), V2421)


tmp5565 := Call(__e, tmp5563, tmp5564)


ifres5562 = tmp5565


} else {
tmp5566 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5562 = tmp5566


}

__e.TailApply(tmp5559, ifres5562)
return


}, 1)

tmp5568 := Call(__e, ns2_1set, symshen_4_5rrb_6, tmp5558)


_ = tmp5568

tmp5569 := MakeNative(func(__e *ControlFlow) {
V2424 := __e.Get(1)
_ = V2424
tmp5570 := MakeNative(func(__e *ControlFlow) {
W2425 := __e.Get(1)
_ = W2425
tmp5572 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2425)


if True == tmp5572 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2425)
return
}


}, 1)

tmp5578 := Call(__e, PrimFunc(symshen_4hds_a_2), V2424, MakeNumber(123))


var ifres5573 Obj

if True == tmp5578 {
tmp5574 := MakeNative(func(__e *ControlFlow) {
W2426 := __e.Get(1)
_ = W2426
__e.TailApply(PrimFunc(symshen_4comb), W2426, symshen_4skip)
return
}, 1)

tmp5575 := Call(__e, PrimFunc(symtail), V2424)


tmp5576 := Call(__e, tmp5574, tmp5575)


ifres5573 = tmp5576


} else {
tmp5577 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5573 = tmp5577


}

__e.TailApply(tmp5570, ifres5573)
return


}, 1)

tmp5579 := Call(__e, ns2_1set, symshen_4_5lcurly_6, tmp5569)


_ = tmp5579

tmp5580 := MakeNative(func(__e *ControlFlow) {
V2427 := __e.Get(1)
_ = V2427
tmp5581 := MakeNative(func(__e *ControlFlow) {
W2428 := __e.Get(1)
_ = W2428
tmp5583 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2428)


if True == tmp5583 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2428)
return
}


}, 1)

tmp5589 := Call(__e, PrimFunc(symshen_4hds_a_2), V2427, MakeNumber(125))


var ifres5584 Obj

if True == tmp5589 {
tmp5585 := MakeNative(func(__e *ControlFlow) {
W2429 := __e.Get(1)
_ = W2429
__e.TailApply(PrimFunc(symshen_4comb), W2429, symshen_4skip)
return
}, 1)

tmp5586 := Call(__e, PrimFunc(symtail), V2427)


tmp5587 := Call(__e, tmp5585, tmp5586)


ifres5584 = tmp5587


} else {
tmp5588 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5584 = tmp5588


}

__e.TailApply(tmp5581, ifres5584)
return


}, 1)

tmp5590 := Call(__e, ns2_1set, symshen_4_5rcurly_6, tmp5580)


_ = tmp5590

tmp5591 := MakeNative(func(__e *ControlFlow) {
V2430 := __e.Get(1)
_ = V2430
tmp5592 := MakeNative(func(__e *ControlFlow) {
W2431 := __e.Get(1)
_ = W2431
tmp5594 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2431)


if True == tmp5594 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2431)
return
}


}, 1)

tmp5600 := Call(__e, PrimFunc(symshen_4hds_a_2), V2430, MakeNumber(124))


var ifres5595 Obj

if True == tmp5600 {
tmp5596 := MakeNative(func(__e *ControlFlow) {
W2432 := __e.Get(1)
_ = W2432
__e.TailApply(PrimFunc(symshen_4comb), W2432, symshen_4skip)
return
}, 1)

tmp5597 := Call(__e, PrimFunc(symtail), V2430)


tmp5598 := Call(__e, tmp5596, tmp5597)


ifres5595 = tmp5598


} else {
tmp5599 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5595 = tmp5599


}

__e.TailApply(tmp5592, ifres5595)
return


}, 1)

tmp5601 := Call(__e, ns2_1set, symshen_4_5bar_6, tmp5591)


_ = tmp5601

tmp5602 := MakeNative(func(__e *ControlFlow) {
V2433 := __e.Get(1)
_ = V2433
tmp5603 := MakeNative(func(__e *ControlFlow) {
W2434 := __e.Get(1)
_ = W2434
tmp5605 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2434)


if True == tmp5605 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2434)
return
}


}, 1)

tmp5611 := Call(__e, PrimFunc(symshen_4hds_a_2), V2433, MakeNumber(59))


var ifres5606 Obj

if True == tmp5611 {
tmp5607 := MakeNative(func(__e *ControlFlow) {
W2435 := __e.Get(1)
_ = W2435
__e.TailApply(PrimFunc(symshen_4comb), W2435, symshen_4skip)
return
}, 1)

tmp5608 := Call(__e, PrimFunc(symtail), V2433)


tmp5609 := Call(__e, tmp5607, tmp5608)


ifres5606 = tmp5609


} else {
tmp5610 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5606 = tmp5610


}

__e.TailApply(tmp5603, ifres5606)
return


}, 1)

tmp5612 := Call(__e, ns2_1set, symshen_4_5semicolon_6, tmp5602)


_ = tmp5612

tmp5613 := MakeNative(func(__e *ControlFlow) {
V2436 := __e.Get(1)
_ = V2436
tmp5614 := MakeNative(func(__e *ControlFlow) {
W2437 := __e.Get(1)
_ = W2437
tmp5616 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2437)


if True == tmp5616 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2437)
return
}


}, 1)

tmp5622 := Call(__e, PrimFunc(symshen_4hds_a_2), V2436, MakeNumber(58))


var ifres5617 Obj

if True == tmp5622 {
tmp5618 := MakeNative(func(__e *ControlFlow) {
W2438 := __e.Get(1)
_ = W2438
__e.TailApply(PrimFunc(symshen_4comb), W2438, symshen_4skip)
return
}, 1)

tmp5619 := Call(__e, PrimFunc(symtail), V2436)


tmp5620 := Call(__e, tmp5618, tmp5619)


ifres5617 = tmp5620


} else {
tmp5621 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5617 = tmp5621


}

__e.TailApply(tmp5614, ifres5617)
return


}, 1)

tmp5623 := Call(__e, ns2_1set, symshen_4_5colon_6, tmp5613)


_ = tmp5623

tmp5624 := MakeNative(func(__e *ControlFlow) {
V2439 := __e.Get(1)
_ = V2439
tmp5625 := MakeNative(func(__e *ControlFlow) {
W2440 := __e.Get(1)
_ = W2440
tmp5627 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2440)


if True == tmp5627 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2440)
return
}


}, 1)

tmp5633 := Call(__e, PrimFunc(symshen_4hds_a_2), V2439, MakeNumber(44))


var ifres5628 Obj

if True == tmp5633 {
tmp5629 := MakeNative(func(__e *ControlFlow) {
W2441 := __e.Get(1)
_ = W2441
__e.TailApply(PrimFunc(symshen_4comb), W2441, symshen_4skip)
return
}, 1)

tmp5630 := Call(__e, PrimFunc(symtail), V2439)


tmp5631 := Call(__e, tmp5629, tmp5630)


ifres5628 = tmp5631


} else {
tmp5632 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5628 = tmp5632


}

__e.TailApply(tmp5625, ifres5628)
return


}, 1)

tmp5634 := Call(__e, ns2_1set, symshen_4_5comma_6, tmp5624)


_ = tmp5634

tmp5635 := MakeNative(func(__e *ControlFlow) {
V2442 := __e.Get(1)
_ = V2442
tmp5636 := MakeNative(func(__e *ControlFlow) {
W2443 := __e.Get(1)
_ = W2443
tmp5638 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2443)


if True == tmp5638 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2443)
return
}


}, 1)

tmp5644 := Call(__e, PrimFunc(symshen_4hds_a_2), V2442, MakeNumber(61))


var ifres5639 Obj

if True == tmp5644 {
tmp5640 := MakeNative(func(__e *ControlFlow) {
W2444 := __e.Get(1)
_ = W2444
__e.TailApply(PrimFunc(symshen_4comb), W2444, symshen_4skip)
return
}, 1)

tmp5641 := Call(__e, PrimFunc(symtail), V2442)


tmp5642 := Call(__e, tmp5640, tmp5641)


ifres5639 = tmp5642


} else {
tmp5643 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5639 = tmp5643


}

__e.TailApply(tmp5636, ifres5639)
return


}, 1)

tmp5645 := Call(__e, ns2_1set, symshen_4_5equal_6, tmp5635)


_ = tmp5645

tmp5646 := MakeNative(func(__e *ControlFlow) {
V2445 := __e.Get(1)
_ = V2445
tmp5647 := MakeNative(func(__e *ControlFlow) {
W2446 := __e.Get(1)
_ = W2446
tmp5659 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2446)


if True == tmp5659 {
tmp5648 := MakeNative(func(__e *ControlFlow) {
W2449 := __e.Get(1)
_ = W2449
tmp5650 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2449)


if True == tmp5650 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2449)
return
}


}, 1)

tmp5651 := MakeNative(func(__e *ControlFlow) {
W2450 := __e.Get(1)
_ = W2450
tmp5655 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2450)


if True == tmp5655 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5652 := MakeNative(func(__e *ControlFlow) {
W2451 := __e.Get(1)
_ = W2451
__e.TailApply(PrimFunc(symshen_4comb), W2451, symshen_4skip)
return
}, 1)

tmp5653 := Call(__e, PrimFunc(symshen_4in_1_6), W2450)


__e.TailApply(tmp5652, tmp5653)
return


}


}, 1)

tmp5656 := Call(__e, PrimFunc(symshen_4_5multiline_6), V2445)


tmp5657 := Call(__e, tmp5651, tmp5656)


__e.TailApply(tmp5648, tmp5657)
return


} else {
__e.Return(W2446)
return
}


}, 1)

tmp5660 := MakeNative(func(__e *ControlFlow) {
W2447 := __e.Get(1)
_ = W2447
tmp5664 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2447)


if True == tmp5664 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5661 := MakeNative(func(__e *ControlFlow) {
W2448 := __e.Get(1)
_ = W2448
__e.TailApply(PrimFunc(symshen_4comb), W2448, symshen_4skip)
return
}, 1)

tmp5662 := Call(__e, PrimFunc(symshen_4in_1_6), W2447)


__e.TailApply(tmp5661, tmp5662)
return


}


}, 1)

tmp5665 := Call(__e, PrimFunc(symshen_4_5singleline_6), V2445)


tmp5666 := Call(__e, tmp5660, tmp5665)


__e.TailApply(tmp5647, tmp5666)
return


}, 1)

tmp5667 := Call(__e, ns2_1set, symshen_4_5comment_6, tmp5646)


_ = tmp5667

tmp5668 := MakeNative(func(__e *ControlFlow) {
V2452 := __e.Get(1)
_ = V2452
tmp5669 := MakeNative(func(__e *ControlFlow) {
W2453 := __e.Get(1)
_ = W2453
tmp5671 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2453)


if True == tmp5671 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2453)
return
}


}, 1)

tmp5672 := MakeNative(func(__e *ControlFlow) {
W2454 := __e.Get(1)
_ = W2454
tmp5694 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2454)


if True == tmp5694 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5673 := MakeNative(func(__e *ControlFlow) {
W2455 := __e.Get(1)
_ = W2455
tmp5674 := MakeNative(func(__e *ControlFlow) {
W2456 := __e.Get(1)
_ = W2456
tmp5690 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2456)


if True == tmp5690 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5675 := MakeNative(func(__e *ControlFlow) {
W2457 := __e.Get(1)
_ = W2457
tmp5676 := MakeNative(func(__e *ControlFlow) {
W2458 := __e.Get(1)
_ = W2458
tmp5686 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2458)


if True == tmp5686 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5677 := MakeNative(func(__e *ControlFlow) {
W2459 := __e.Get(1)
_ = W2459
tmp5678 := MakeNative(func(__e *ControlFlow) {
W2460 := __e.Get(1)
_ = W2460
tmp5682 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2460)


if True == tmp5682 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5679 := MakeNative(func(__e *ControlFlow) {
W2461 := __e.Get(1)
_ = W2461
__e.TailApply(PrimFunc(symshen_4comb), W2461, symshen_4skip)
return
}, 1)

tmp5680 := Call(__e, PrimFunc(symshen_4in_1_6), W2460)


__e.TailApply(tmp5679, tmp5680)
return


}


}, 1)

tmp5683 := Call(__e, PrimFunc(symshen_4_5returns_6), W2459)


__e.TailApply(tmp5678, tmp5683)
return


}, 1)

tmp5684 := Call(__e, PrimFunc(symshen_4in_1_6), W2458)


__e.TailApply(tmp5677, tmp5684)
return


}


}, 1)

tmp5687 := Call(__e, PrimFunc(symshen_4_5shortnatters_6), W2457)


__e.TailApply(tmp5676, tmp5687)
return


}, 1)

tmp5688 := Call(__e, PrimFunc(symshen_4in_1_6), W2456)


__e.TailApply(tmp5675, tmp5688)
return


}


}, 1)

tmp5691 := Call(__e, PrimFunc(symshen_4_5backslash_6), W2455)


__e.TailApply(tmp5674, tmp5691)
return


}, 1)

tmp5692 := Call(__e, PrimFunc(symshen_4in_1_6), W2454)


__e.TailApply(tmp5673, tmp5692)
return


}


}, 1)

tmp5695 := Call(__e, PrimFunc(symshen_4_5backslash_6), V2452)


tmp5696 := Call(__e, tmp5672, tmp5695)


__e.TailApply(tmp5669, tmp5696)
return


}, 1)

tmp5697 := Call(__e, ns2_1set, symshen_4_5singleline_6, tmp5668)


_ = tmp5697

tmp5698 := MakeNative(func(__e *ControlFlow) {
V2462 := __e.Get(1)
_ = V2462
tmp5699 := MakeNative(func(__e *ControlFlow) {
W2463 := __e.Get(1)
_ = W2463
tmp5701 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2463)


if True == tmp5701 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2463)
return
}


}, 1)

tmp5707 := Call(__e, PrimFunc(symshen_4hds_a_2), V2462, MakeNumber(92))


var ifres5702 Obj

if True == tmp5707 {
tmp5703 := MakeNative(func(__e *ControlFlow) {
W2464 := __e.Get(1)
_ = W2464
__e.TailApply(PrimFunc(symshen_4comb), W2464, symshen_4skip)
return
}, 1)

tmp5704 := Call(__e, PrimFunc(symtail), V2462)


tmp5705 := Call(__e, tmp5703, tmp5704)


ifres5702 = tmp5705


} else {
tmp5706 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5702 = tmp5706


}

__e.TailApply(tmp5699, ifres5702)
return


}, 1)

tmp5708 := Call(__e, ns2_1set, symshen_4_5backslash_6, tmp5698)


_ = tmp5708

tmp5709 := MakeNative(func(__e *ControlFlow) {
V2465 := __e.Get(1)
_ = V2465
tmp5710 := MakeNative(func(__e *ControlFlow) {
W2466 := __e.Get(1)
_ = W2466
tmp5722 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2466)


if True == tmp5722 {
tmp5711 := MakeNative(func(__e *ControlFlow) {
W2471 := __e.Get(1)
_ = W2471
tmp5713 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2471)


if True == tmp5713 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2471)
return
}


}, 1)

tmp5714 := MakeNative(func(__e *ControlFlow) {
W2472 := __e.Get(1)
_ = W2472
tmp5718 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2472)


if True == tmp5718 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5715 := MakeNative(func(__e *ControlFlow) {
W2473 := __e.Get(1)
_ = W2473
__e.TailApply(PrimFunc(symshen_4comb), W2473, symshen_4skip)
return
}, 1)

tmp5716 := Call(__e, PrimFunc(symshen_4in_1_6), W2472)


__e.TailApply(tmp5715, tmp5716)
return


}


}, 1)

tmp5719 := Call(__e, PrimFunc(sym_5e_6), V2465)


tmp5720 := Call(__e, tmp5714, tmp5719)


__e.TailApply(tmp5711, tmp5720)
return


} else {
__e.Return(W2466)
return
}


}, 1)

tmp5723 := MakeNative(func(__e *ControlFlow) {
W2467 := __e.Get(1)
_ = W2467
tmp5733 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2467)


if True == tmp5733 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5724 := MakeNative(func(__e *ControlFlow) {
W2468 := __e.Get(1)
_ = W2468
tmp5725 := MakeNative(func(__e *ControlFlow) {
W2469 := __e.Get(1)
_ = W2469
tmp5729 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2469)


if True == tmp5729 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5726 := MakeNative(func(__e *ControlFlow) {
W2470 := __e.Get(1)
_ = W2470
__e.TailApply(PrimFunc(symshen_4comb), W2470, symshen_4skip)
return
}, 1)

tmp5727 := Call(__e, PrimFunc(symshen_4in_1_6), W2469)


__e.TailApply(tmp5726, tmp5727)
return


}


}, 1)

tmp5730 := Call(__e, PrimFunc(symshen_4_5shortnatters_6), W2468)


__e.TailApply(tmp5725, tmp5730)
return


}, 1)

tmp5731 := Call(__e, PrimFunc(symshen_4in_1_6), W2467)


__e.TailApply(tmp5724, tmp5731)
return


}


}, 1)

tmp5734 := Call(__e, PrimFunc(symshen_4_5shortnatter_6), V2465)


tmp5735 := Call(__e, tmp5723, tmp5734)


__e.TailApply(tmp5710, tmp5735)
return


}, 1)

tmp5736 := Call(__e, ns2_1set, symshen_4_5shortnatters_6, tmp5709)


_ = tmp5736

tmp5737 := MakeNative(func(__e *ControlFlow) {
V2474 := __e.Get(1)
_ = V2474
tmp5738 := MakeNative(func(__e *ControlFlow) {
W2475 := __e.Get(1)
_ = W2475
tmp5740 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2475)


if True == tmp5740 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2475)
return
}


}, 1)

tmp5751 := PrimIsPair(V2474)

var ifres5741 Obj

if True == tmp5751 {
tmp5742 := MakeNative(func(__e *ControlFlow) {
W2476 := __e.Get(1)
_ = W2476
tmp5743 := MakeNative(func(__e *ControlFlow) {
W2477 := __e.Get(1)
_ = W2477
tmp5745 := Call(__e, PrimFunc(symshen_4return_2), W2476)


tmp5746 := PrimNot(tmp5745)

if True == tmp5746 {
__e.TailApply(PrimFunc(symshen_4comb), W2477, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp5747 := Call(__e, PrimFunc(symtail), V2474)


__e.TailApply(tmp5743, tmp5747)
return


}, 1)

tmp5748 := Call(__e, PrimFunc(symhead), V2474)


tmp5749 := Call(__e, tmp5742, tmp5748)


ifres5741 = tmp5749


} else {
tmp5750 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5741 = tmp5750


}

__e.TailApply(tmp5738, ifres5741)
return


}, 1)

tmp5752 := Call(__e, ns2_1set, symshen_4_5shortnatter_6, tmp5737)


_ = tmp5752

tmp5753 := MakeNative(func(__e *ControlFlow) {
V2478 := __e.Get(1)
_ = V2478
tmp5754 := MakeNative(func(__e *ControlFlow) {
W2479 := __e.Get(1)
_ = W2479
tmp5766 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2479)


if True == tmp5766 {
tmp5755 := MakeNative(func(__e *ControlFlow) {
W2484 := __e.Get(1)
_ = W2484
tmp5757 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2484)


if True == tmp5757 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2484)
return
}


}, 1)

tmp5758 := MakeNative(func(__e *ControlFlow) {
W2485 := __e.Get(1)
_ = W2485
tmp5762 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2485)


if True == tmp5762 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5759 := MakeNative(func(__e *ControlFlow) {
W2486 := __e.Get(1)
_ = W2486
__e.TailApply(PrimFunc(symshen_4comb), W2486, symshen_4skip)
return
}, 1)

tmp5760 := Call(__e, PrimFunc(symshen_4in_1_6), W2485)


__e.TailApply(tmp5759, tmp5760)
return


}


}, 1)

tmp5763 := Call(__e, PrimFunc(symshen_4_5return_6), V2478)


tmp5764 := Call(__e, tmp5758, tmp5763)


__e.TailApply(tmp5755, tmp5764)
return


} else {
__e.Return(W2479)
return
}


}, 1)

tmp5767 := MakeNative(func(__e *ControlFlow) {
W2480 := __e.Get(1)
_ = W2480
tmp5777 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2480)


if True == tmp5777 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5768 := MakeNative(func(__e *ControlFlow) {
W2481 := __e.Get(1)
_ = W2481
tmp5769 := MakeNative(func(__e *ControlFlow) {
W2482 := __e.Get(1)
_ = W2482
tmp5773 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2482)


if True == tmp5773 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5770 := MakeNative(func(__e *ControlFlow) {
W2483 := __e.Get(1)
_ = W2483
__e.TailApply(PrimFunc(symshen_4comb), W2483, symshen_4skip)
return
}, 1)

tmp5771 := Call(__e, PrimFunc(symshen_4in_1_6), W2482)


__e.TailApply(tmp5770, tmp5771)
return


}


}, 1)

tmp5774 := Call(__e, PrimFunc(symshen_4_5returns_6), W2481)


__e.TailApply(tmp5769, tmp5774)
return


}, 1)

tmp5775 := Call(__e, PrimFunc(symshen_4in_1_6), W2480)


__e.TailApply(tmp5768, tmp5775)
return


}


}, 1)

tmp5778 := Call(__e, PrimFunc(symshen_4_5return_6), V2478)


tmp5779 := Call(__e, tmp5767, tmp5778)


__e.TailApply(tmp5754, tmp5779)
return


}, 1)

tmp5780 := Call(__e, ns2_1set, symshen_4_5returns_6, tmp5753)


_ = tmp5780

tmp5781 := MakeNative(func(__e *ControlFlow) {
V2487 := __e.Get(1)
_ = V2487
tmp5782 := MakeNative(func(__e *ControlFlow) {
W2488 := __e.Get(1)
_ = W2488
tmp5784 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2488)


if True == tmp5784 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2488)
return
}


}, 1)

tmp5794 := PrimIsPair(V2487)

var ifres5785 Obj

if True == tmp5794 {
tmp5786 := MakeNative(func(__e *ControlFlow) {
W2489 := __e.Get(1)
_ = W2489
tmp5787 := MakeNative(func(__e *ControlFlow) {
W2490 := __e.Get(1)
_ = W2490
tmp5789 := Call(__e, PrimFunc(symshen_4return_2), W2489)


if True == tmp5789 {
__e.TailApply(PrimFunc(symshen_4comb), W2490, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp5790 := Call(__e, PrimFunc(symtail), V2487)


__e.TailApply(tmp5787, tmp5790)
return


}, 1)

tmp5791 := Call(__e, PrimFunc(symhead), V2487)


tmp5792 := Call(__e, tmp5786, tmp5791)


ifres5785 = tmp5792


} else {
tmp5793 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5785 = tmp5793


}

__e.TailApply(tmp5782, ifres5785)
return


}, 1)

tmp5795 := Call(__e, ns2_1set, symshen_4_5return_6, tmp5781)


_ = tmp5795

tmp5796 := MakeNative(func(__e *ControlFlow) {
V2491 := __e.Get(1)
_ = V2491
tmp5797 := PrimCons(MakeNumber(13), Nil)

tmp5798 := PrimCons(MakeNumber(10), tmp5797)

tmp5799 := PrimCons(MakeNumber(9), tmp5798)

__e.TailApply(PrimFunc(symelement_2), V2491, tmp5799)
return


}, 1)

tmp5800 := Call(__e, ns2_1set, symshen_4return_2, tmp5796)


_ = tmp5800

tmp5801 := MakeNative(func(__e *ControlFlow) {
V2492 := __e.Get(1)
_ = V2492
tmp5802 := MakeNative(func(__e *ControlFlow) {
W2493 := __e.Get(1)
_ = W2493
tmp5804 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2493)


if True == tmp5804 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2493)
return
}


}, 1)

tmp5805 := MakeNative(func(__e *ControlFlow) {
W2494 := __e.Get(1)
_ = W2494
tmp5821 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2494)


if True == tmp5821 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5806 := MakeNative(func(__e *ControlFlow) {
W2495 := __e.Get(1)
_ = W2495
tmp5807 := MakeNative(func(__e *ControlFlow) {
W2496 := __e.Get(1)
_ = W2496
tmp5817 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2496)


if True == tmp5817 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5808 := MakeNative(func(__e *ControlFlow) {
W2497 := __e.Get(1)
_ = W2497
tmp5809 := MakeNative(func(__e *ControlFlow) {
W2498 := __e.Get(1)
_ = W2498
tmp5813 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2498)


if True == tmp5813 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5810 := MakeNative(func(__e *ControlFlow) {
W2499 := __e.Get(1)
_ = W2499
__e.TailApply(PrimFunc(symshen_4comb), W2499, symshen_4skip)
return
}, 1)

tmp5811 := Call(__e, PrimFunc(symshen_4in_1_6), W2498)


__e.TailApply(tmp5810, tmp5811)
return


}


}, 1)

tmp5814 := Call(__e, PrimFunc(symshen_4_5longnatter_6), W2497)


__e.TailApply(tmp5809, tmp5814)
return


}, 1)

tmp5815 := Call(__e, PrimFunc(symshen_4in_1_6), W2496)


__e.TailApply(tmp5808, tmp5815)
return


}


}, 1)

tmp5818 := Call(__e, PrimFunc(symshen_4_5times_6), W2495)


__e.TailApply(tmp5807, tmp5818)
return


}, 1)

tmp5819 := Call(__e, PrimFunc(symshen_4in_1_6), W2494)


__e.TailApply(tmp5806, tmp5819)
return


}


}, 1)

tmp5822 := Call(__e, PrimFunc(symshen_4_5backslash_6), V2492)


tmp5823 := Call(__e, tmp5805, tmp5822)


__e.TailApply(tmp5802, tmp5823)
return


}, 1)

tmp5824 := Call(__e, ns2_1set, symshen_4_5multiline_6, tmp5801)


_ = tmp5824

tmp5825 := MakeNative(func(__e *ControlFlow) {
V2500 := __e.Get(1)
_ = V2500
tmp5826 := MakeNative(func(__e *ControlFlow) {
W2501 := __e.Get(1)
_ = W2501
tmp5828 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2501)


if True == tmp5828 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2501)
return
}


}, 1)

tmp5834 := Call(__e, PrimFunc(symshen_4hds_a_2), V2500, MakeNumber(42))


var ifres5829 Obj

if True == tmp5834 {
tmp5830 := MakeNative(func(__e *ControlFlow) {
W2502 := __e.Get(1)
_ = W2502
__e.TailApply(PrimFunc(symshen_4comb), W2502, symshen_4skip)
return
}, 1)

tmp5831 := Call(__e, PrimFunc(symtail), V2500)


tmp5832 := Call(__e, tmp5830, tmp5831)


ifres5829 = tmp5832


} else {
tmp5833 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5829 = tmp5833


}

__e.TailApply(tmp5826, ifres5829)
return


}, 1)

tmp5835 := Call(__e, ns2_1set, symshen_4_5times_6, tmp5825)


_ = tmp5835

tmp5836 := MakeNative(func(__e *ControlFlow) {
V2503 := __e.Get(1)
_ = V2503
tmp5837 := MakeNative(func(__e *ControlFlow) {
W2504 := __e.Get(1)
_ = W2504
tmp5870 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2504)


if True == tmp5870 {
tmp5838 := MakeNative(func(__e *ControlFlow) {
W2509 := __e.Get(1)
_ = W2509
tmp5855 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2509)


if True == tmp5855 {
tmp5839 := MakeNative(func(__e *ControlFlow) {
W2514 := __e.Get(1)
_ = W2514
tmp5841 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2514)


if True == tmp5841 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2514)
return
}


}, 1)

tmp5853 := PrimIsPair(V2503)

var ifres5842 Obj

if True == tmp5853 {
tmp5843 := MakeNative(func(__e *ControlFlow) {
W2515 := __e.Get(1)
_ = W2515
tmp5844 := MakeNative(func(__e *ControlFlow) {
W2516 := __e.Get(1)
_ = W2516
tmp5848 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2516)


if True == tmp5848 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5845 := MakeNative(func(__e *ControlFlow) {
W2517 := __e.Get(1)
_ = W2517
__e.TailApply(PrimFunc(symshen_4comb), W2517, symshen_4skip)
return
}, 1)

tmp5846 := Call(__e, PrimFunc(symshen_4in_1_6), W2516)


__e.TailApply(tmp5845, tmp5846)
return


}


}, 1)

tmp5849 := Call(__e, PrimFunc(symshen_4_5longnatter_6), W2515)


__e.TailApply(tmp5844, tmp5849)
return


}, 1)

tmp5850 := Call(__e, PrimFunc(symtail), V2503)


tmp5851 := Call(__e, tmp5843, tmp5850)


ifres5842 = tmp5851


} else {
tmp5852 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5842 = tmp5852


}

__e.TailApply(tmp5839, ifres5842)
return


} else {
__e.Return(W2509)
return
}


}, 1)

tmp5856 := MakeNative(func(__e *ControlFlow) {
W2510 := __e.Get(1)
_ = W2510
tmp5866 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2510)


if True == tmp5866 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5857 := MakeNative(func(__e *ControlFlow) {
W2511 := __e.Get(1)
_ = W2511
tmp5858 := MakeNative(func(__e *ControlFlow) {
W2512 := __e.Get(1)
_ = W2512
tmp5862 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2512)


if True == tmp5862 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5859 := MakeNative(func(__e *ControlFlow) {
W2513 := __e.Get(1)
_ = W2513
__e.TailApply(PrimFunc(symshen_4comb), W2513, symshen_4skip)
return
}, 1)

tmp5860 := Call(__e, PrimFunc(symshen_4in_1_6), W2512)


__e.TailApply(tmp5859, tmp5860)
return


}


}, 1)

tmp5863 := Call(__e, PrimFunc(symshen_4_5backslash_6), W2511)


__e.TailApply(tmp5858, tmp5863)
return


}, 1)

tmp5864 := Call(__e, PrimFunc(symshen_4in_1_6), W2510)


__e.TailApply(tmp5857, tmp5864)
return


}


}, 1)

tmp5867 := Call(__e, PrimFunc(symshen_4_5times_6), V2503)


tmp5868 := Call(__e, tmp5856, tmp5867)


__e.TailApply(tmp5838, tmp5868)
return


} else {
__e.Return(W2504)
return
}


}, 1)

tmp5871 := MakeNative(func(__e *ControlFlow) {
W2505 := __e.Get(1)
_ = W2505
tmp5881 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2505)


if True == tmp5881 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5872 := MakeNative(func(__e *ControlFlow) {
W2506 := __e.Get(1)
_ = W2506
tmp5873 := MakeNative(func(__e *ControlFlow) {
W2507 := __e.Get(1)
_ = W2507
tmp5877 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2507)


if True == tmp5877 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5874 := MakeNative(func(__e *ControlFlow) {
W2508 := __e.Get(1)
_ = W2508
__e.TailApply(PrimFunc(symshen_4comb), W2508, symshen_4skip)
return
}, 1)

tmp5875 := Call(__e, PrimFunc(symshen_4in_1_6), W2507)


__e.TailApply(tmp5874, tmp5875)
return


}


}, 1)

tmp5878 := Call(__e, PrimFunc(symshen_4_5longnatter_6), W2506)


__e.TailApply(tmp5873, tmp5878)
return


}, 1)

tmp5879 := Call(__e, PrimFunc(symshen_4in_1_6), W2505)


__e.TailApply(tmp5872, tmp5879)
return


}


}, 1)

tmp5882 := Call(__e, PrimFunc(symshen_4_5comment_6), V2503)


tmp5883 := Call(__e, tmp5871, tmp5882)


__e.TailApply(tmp5837, tmp5883)
return


}, 1)

tmp5884 := Call(__e, ns2_1set, symshen_4_5longnatter_6, tmp5836)


_ = tmp5884

tmp5885 := MakeNative(func(__e *ControlFlow) {
V2518 := __e.Get(1)
_ = V2518
tmp5886 := MakeNative(func(__e *ControlFlow) {
W2519 := __e.Get(1)
_ = W2519
tmp5917 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2519)


if True == tmp5917 {
tmp5887 := MakeNative(func(__e *ControlFlow) {
W2523 := __e.Get(1)
_ = W2523
tmp5906 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2523)


if True == tmp5906 {
tmp5888 := MakeNative(func(__e *ControlFlow) {
W2527 := __e.Get(1)
_ = W2527
tmp5890 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2527)


if True == tmp5890 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2527)
return
}


}, 1)

tmp5891 := MakeNative(func(__e *ControlFlow) {
W2528 := __e.Get(1)
_ = W2528
tmp5902 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2528)


if True == tmp5902 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5892 := MakeNative(func(__e *ControlFlow) {
W2529 := __e.Get(1)
_ = W2529
tmp5893 := MakeNative(func(__e *ControlFlow) {
W2530 := __e.Get(1)
_ = W2530
tmp5898 := PrimEqual(W2529, MakeString("<>"))

var ifres5894 Obj

if True == tmp5898 {
tmp5895 := PrimCons(MakeNumber(0), Nil)

tmp5896 := PrimCons(symvector, tmp5895)

ifres5894 = tmp5896


} else {
tmp5897 := PrimIntern(W2529)

ifres5894 = tmp5897


}

__e.TailApply(PrimFunc(symshen_4comb), W2530, ifres5894)
return


}, 1)

tmp5899 := Call(__e, PrimFunc(symshen_4in_1_6), W2528)


__e.TailApply(tmp5893, tmp5899)
return


}, 1)

tmp5900 := Call(__e, PrimFunc(symshen_4_5_1out), W2528)


__e.TailApply(tmp5892, tmp5900)
return


}


}, 1)

tmp5903 := Call(__e, PrimFunc(symshen_4_5sym_6), V2518)


tmp5904 := Call(__e, tmp5891, tmp5903)


__e.TailApply(tmp5888, tmp5904)
return


} else {
__e.Return(W2523)
return
}


}, 1)

tmp5907 := MakeNative(func(__e *ControlFlow) {
W2524 := __e.Get(1)
_ = W2524
tmp5913 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2524)


if True == tmp5913 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5908 := MakeNative(func(__e *ControlFlow) {
W2525 := __e.Get(1)
_ = W2525
tmp5909 := MakeNative(func(__e *ControlFlow) {
W2526 := __e.Get(1)
_ = W2526
__e.TailApply(PrimFunc(symshen_4comb), W2526, W2525)
return
}, 1)

tmp5910 := Call(__e, PrimFunc(symshen_4in_1_6), W2524)


__e.TailApply(tmp5909, tmp5910)
return


}, 1)

tmp5911 := Call(__e, PrimFunc(symshen_4_5_1out), W2524)


__e.TailApply(tmp5908, tmp5911)
return


}


}, 1)

tmp5914 := Call(__e, PrimFunc(symshen_4_5number_6), V2518)


tmp5915 := Call(__e, tmp5907, tmp5914)


__e.TailApply(tmp5887, tmp5915)
return


} else {
__e.Return(W2519)
return
}


}, 1)

tmp5918 := MakeNative(func(__e *ControlFlow) {
W2520 := __e.Get(1)
_ = W2520
tmp5924 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2520)


if True == tmp5924 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5919 := MakeNative(func(__e *ControlFlow) {
W2521 := __e.Get(1)
_ = W2521
tmp5920 := MakeNative(func(__e *ControlFlow) {
W2522 := __e.Get(1)
_ = W2522
__e.TailApply(PrimFunc(symshen_4comb), W2522, W2521)
return
}, 1)

tmp5921 := Call(__e, PrimFunc(symshen_4in_1_6), W2520)


__e.TailApply(tmp5920, tmp5921)
return


}, 1)

tmp5922 := Call(__e, PrimFunc(symshen_4_5_1out), W2520)


__e.TailApply(tmp5919, tmp5922)
return


}


}, 1)

tmp5925 := Call(__e, PrimFunc(symshen_4_5str_6), V2518)


tmp5926 := Call(__e, tmp5918, tmp5925)


__e.TailApply(tmp5886, tmp5926)
return


}, 1)

tmp5927 := Call(__e, ns2_1set, symshen_4_5atom_6, tmp5885)


_ = tmp5927

tmp5928 := MakeNative(func(__e *ControlFlow) {
V2531 := __e.Get(1)
_ = V2531
tmp5929 := MakeNative(func(__e *ControlFlow) {
W2532 := __e.Get(1)
_ = W2532
tmp5931 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2532)


if True == tmp5931 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2532)
return
}


}, 1)

tmp5932 := MakeNative(func(__e *ControlFlow) {
W2533 := __e.Get(1)
_ = W2533
tmp5947 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2533)


if True == tmp5947 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5933 := MakeNative(func(__e *ControlFlow) {
W2534 := __e.Get(1)
_ = W2534
tmp5934 := MakeNative(func(__e *ControlFlow) {
W2535 := __e.Get(1)
_ = W2535
tmp5935 := MakeNative(func(__e *ControlFlow) {
W2536 := __e.Get(1)
_ = W2536
tmp5942 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2536)


if True == tmp5942 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5936 := MakeNative(func(__e *ControlFlow) {
W2537 := __e.Get(1)
_ = W2537
tmp5937 := MakeNative(func(__e *ControlFlow) {
W2538 := __e.Get(1)
_ = W2538
tmp5938 := PrimStringConcat(W2534, W2537)

__e.TailApply(PrimFunc(symshen_4comb), W2538, tmp5938)
return


}, 1)

tmp5939 := Call(__e, PrimFunc(symshen_4in_1_6), W2536)


__e.TailApply(tmp5937, tmp5939)
return


}, 1)

tmp5940 := Call(__e, PrimFunc(symshen_4_5_1out), W2536)


__e.TailApply(tmp5936, tmp5940)
return


}


}, 1)

tmp5943 := Call(__e, PrimFunc(symshen_4_5alphanums_6), W2535)


__e.TailApply(tmp5935, tmp5943)
return


}, 1)

tmp5944 := Call(__e, PrimFunc(symshen_4in_1_6), W2533)


__e.TailApply(tmp5934, tmp5944)
return


}, 1)

tmp5945 := Call(__e, PrimFunc(symshen_4_5_1out), W2533)


__e.TailApply(tmp5933, tmp5945)
return


}


}, 1)

tmp5948 := Call(__e, PrimFunc(symshen_4_5alpha_6), V2531)


tmp5949 := Call(__e, tmp5932, tmp5948)


__e.TailApply(tmp5929, tmp5949)
return


}, 1)

tmp5950 := Call(__e, ns2_1set, symshen_4_5sym_6, tmp5928)


_ = tmp5950

tmp5951 := MakeNative(func(__e *ControlFlow) {
V2539 := __e.Get(1)
_ = V2539
tmp5952 := MakeNative(func(__e *ControlFlow) {
W2540 := __e.Get(1)
_ = W2540
tmp5954 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2540)


if True == tmp5954 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2540)
return
}


}, 1)

tmp5965 := PrimIsPair(V2539)

var ifres5955 Obj

if True == tmp5965 {
tmp5956 := MakeNative(func(__e *ControlFlow) {
W2541 := __e.Get(1)
_ = W2541
tmp5957 := MakeNative(func(__e *ControlFlow) {
W2542 := __e.Get(1)
_ = W2542
tmp5960 := Call(__e, PrimFunc(symshen_4alpha_2), W2541)


if True == tmp5960 {
tmp5958 := PrimNumberToString(W2541)

__e.TailApply(PrimFunc(symshen_4comb), W2542, tmp5958)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp5961 := Call(__e, PrimFunc(symtail), V2539)


__e.TailApply(tmp5957, tmp5961)
return


}, 1)

tmp5962 := Call(__e, PrimFunc(symhead), V2539)


tmp5963 := Call(__e, tmp5956, tmp5962)


ifres5955 = tmp5963


} else {
tmp5964 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5955 = tmp5964


}

__e.TailApply(tmp5952, ifres5955)
return


}, 1)

tmp5966 := Call(__e, ns2_1set, symshen_4_5alpha_6, tmp5951)


_ = tmp5966

tmp5967 := MakeNative(func(__e *ControlFlow) {
V2543 := __e.Get(1)
_ = V2543
tmp5974 := Call(__e, PrimFunc(symshen_4lowercase_2), V2543)


if True == tmp5974 {
__e.Return(True)
return
} else {
tmp5972 := Call(__e, PrimFunc(symshen_4uppercase_2), V2543)


var ifres5969 Obj

if True == tmp5972 {
ifres5969 = True


} else {
tmp5971 := Call(__e, PrimFunc(symshen_4misc_2), V2543)


var ifres5970 Obj

if True == tmp5971 {
ifres5970 = True


} else {
ifres5970 = False


}

ifres5969 = ifres5970


}

if True == ifres5969 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp5975 := Call(__e, ns2_1set, symshen_4alpha_2, tmp5967)


_ = tmp5975

tmp5976 := MakeNative(func(__e *ControlFlow) {
V2544 := __e.Get(1)
_ = V2544
tmp5980 := PrimGreatEqual(V2544, MakeNumber(97))

if True == tmp5980 {
tmp5978 := PrimLessEqual(V2544, MakeNumber(122))

if True == tmp5978 {
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

tmp5981 := Call(__e, ns2_1set, symshen_4lowercase_2, tmp5976)


_ = tmp5981

tmp5982 := MakeNative(func(__e *ControlFlow) {
V2545 := __e.Get(1)
_ = V2545
tmp5986 := PrimGreatEqual(V2545, MakeNumber(65))

if True == tmp5986 {
tmp5984 := PrimLessEqual(V2545, MakeNumber(90))

if True == tmp5984 {
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

tmp5987 := Call(__e, ns2_1set, symshen_4uppercase_2, tmp5982)


_ = tmp5987

tmp5988 := MakeNative(func(__e *ControlFlow) {
V2546 := __e.Get(1)
_ = V2546
tmp5989 := PrimCons(MakeNumber(96), Nil)

tmp5990 := PrimCons(MakeNumber(35), tmp5989)

tmp5991 := PrimCons(MakeNumber(39), tmp5990)

tmp5992 := PrimCons(MakeNumber(37), tmp5991)

tmp5993 := PrimCons(MakeNumber(38), tmp5992)

tmp5994 := PrimCons(MakeNumber(60), tmp5993)

tmp5995 := PrimCons(MakeNumber(62), tmp5994)

tmp5996 := PrimCons(MakeNumber(46), tmp5995)

tmp5997 := PrimCons(MakeNumber(126), tmp5996)

tmp5998 := PrimCons(MakeNumber(64), tmp5997)

tmp5999 := PrimCons(MakeNumber(33), tmp5998)

tmp6000 := PrimCons(MakeNumber(36), tmp5999)

tmp6001 := PrimCons(MakeNumber(63), tmp6000)

tmp6002 := PrimCons(MakeNumber(95), tmp6001)

tmp6003 := PrimCons(MakeNumber(43), tmp6002)

tmp6004 := PrimCons(MakeNumber(47), tmp6003)

tmp6005 := PrimCons(MakeNumber(42), tmp6004)

tmp6006 := PrimCons(MakeNumber(45), tmp6005)

tmp6007 := PrimCons(MakeNumber(61), tmp6006)

__e.TailApply(PrimFunc(symelement_2), V2546, tmp6007)
return


}, 1)

tmp6008 := Call(__e, ns2_1set, symshen_4misc_2, tmp5988)


_ = tmp6008

tmp6009 := MakeNative(func(__e *ControlFlow) {
V2547 := __e.Get(1)
_ = V2547
tmp6010 := MakeNative(func(__e *ControlFlow) {
W2548 := __e.Get(1)
_ = W2548
tmp6022 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2548)


if True == tmp6022 {
tmp6011 := MakeNative(func(__e *ControlFlow) {
W2555 := __e.Get(1)
_ = W2555
tmp6013 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2555)


if True == tmp6013 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2555)
return
}


}, 1)

tmp6014 := MakeNative(func(__e *ControlFlow) {
W2556 := __e.Get(1)
_ = W2556
tmp6018 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2556)


if True == tmp6018 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6015 := MakeNative(func(__e *ControlFlow) {
W2557 := __e.Get(1)
_ = W2557
__e.TailApply(PrimFunc(symshen_4comb), W2557, MakeString(""))
return
}, 1)

tmp6016 := Call(__e, PrimFunc(symshen_4in_1_6), W2556)


__e.TailApply(tmp6015, tmp6016)
return


}


}, 1)

tmp6019 := Call(__e, PrimFunc(sym_5e_6), V2547)


tmp6020 := Call(__e, tmp6014, tmp6019)


__e.TailApply(tmp6011, tmp6020)
return


} else {
__e.Return(W2548)
return
}


}, 1)

tmp6023 := MakeNative(func(__e *ControlFlow) {
W2549 := __e.Get(1)
_ = W2549
tmp6038 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2549)


if True == tmp6038 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6024 := MakeNative(func(__e *ControlFlow) {
W2550 := __e.Get(1)
_ = W2550
tmp6025 := MakeNative(func(__e *ControlFlow) {
W2551 := __e.Get(1)
_ = W2551
tmp6026 := MakeNative(func(__e *ControlFlow) {
W2552 := __e.Get(1)
_ = W2552
tmp6033 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2552)


if True == tmp6033 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6027 := MakeNative(func(__e *ControlFlow) {
W2553 := __e.Get(1)
_ = W2553
tmp6028 := MakeNative(func(__e *ControlFlow) {
W2554 := __e.Get(1)
_ = W2554
tmp6029 := PrimStringConcat(W2550, W2553)

__e.TailApply(PrimFunc(symshen_4comb), W2554, tmp6029)
return


}, 1)

tmp6030 := Call(__e, PrimFunc(symshen_4in_1_6), W2552)


__e.TailApply(tmp6028, tmp6030)
return


}, 1)

tmp6031 := Call(__e, PrimFunc(symshen_4_5_1out), W2552)


__e.TailApply(tmp6027, tmp6031)
return


}


}, 1)

tmp6034 := Call(__e, PrimFunc(symshen_4_5alphanums_6), W2551)


__e.TailApply(tmp6026, tmp6034)
return


}, 1)

tmp6035 := Call(__e, PrimFunc(symshen_4in_1_6), W2549)


__e.TailApply(tmp6025, tmp6035)
return


}, 1)

tmp6036 := Call(__e, PrimFunc(symshen_4_5_1out), W2549)


__e.TailApply(tmp6024, tmp6036)
return


}


}, 1)

tmp6039 := Call(__e, PrimFunc(symshen_4_5alphanum_6), V2547)


tmp6040 := Call(__e, tmp6023, tmp6039)


__e.TailApply(tmp6010, tmp6040)
return


}, 1)

tmp6041 := Call(__e, ns2_1set, symshen_4_5alphanums_6, tmp6009)


_ = tmp6041

tmp6042 := MakeNative(func(__e *ControlFlow) {
V2558 := __e.Get(1)
_ = V2558
tmp6043 := MakeNative(func(__e *ControlFlow) {
W2559 := __e.Get(1)
_ = W2559
tmp6057 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2559)


if True == tmp6057 {
tmp6044 := MakeNative(func(__e *ControlFlow) {
W2563 := __e.Get(1)
_ = W2563
tmp6046 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2563)


if True == tmp6046 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2563)
return
}


}, 1)

tmp6047 := MakeNative(func(__e *ControlFlow) {
W2564 := __e.Get(1)
_ = W2564
tmp6053 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2564)


if True == tmp6053 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6048 := MakeNative(func(__e *ControlFlow) {
W2565 := __e.Get(1)
_ = W2565
tmp6049 := MakeNative(func(__e *ControlFlow) {
W2566 := __e.Get(1)
_ = W2566
__e.TailApply(PrimFunc(symshen_4comb), W2566, W2565)
return
}, 1)

tmp6050 := Call(__e, PrimFunc(symshen_4in_1_6), W2564)


__e.TailApply(tmp6049, tmp6050)
return


}, 1)

tmp6051 := Call(__e, PrimFunc(symshen_4_5_1out), W2564)


__e.TailApply(tmp6048, tmp6051)
return


}


}, 1)

tmp6054 := Call(__e, PrimFunc(symshen_4_5numeral_6), V2558)


tmp6055 := Call(__e, tmp6047, tmp6054)


__e.TailApply(tmp6044, tmp6055)
return


} else {
__e.Return(W2559)
return
}


}, 1)

tmp6058 := MakeNative(func(__e *ControlFlow) {
W2560 := __e.Get(1)
_ = W2560
tmp6064 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2560)


if True == tmp6064 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6059 := MakeNative(func(__e *ControlFlow) {
W2561 := __e.Get(1)
_ = W2561
tmp6060 := MakeNative(func(__e *ControlFlow) {
W2562 := __e.Get(1)
_ = W2562
__e.TailApply(PrimFunc(symshen_4comb), W2562, W2561)
return
}, 1)

tmp6061 := Call(__e, PrimFunc(symshen_4in_1_6), W2560)


__e.TailApply(tmp6060, tmp6061)
return


}, 1)

tmp6062 := Call(__e, PrimFunc(symshen_4_5_1out), W2560)


__e.TailApply(tmp6059, tmp6062)
return


}


}, 1)

tmp6065 := Call(__e, PrimFunc(symshen_4_5alpha_6), V2558)


tmp6066 := Call(__e, tmp6058, tmp6065)


__e.TailApply(tmp6043, tmp6066)
return


}, 1)

tmp6067 := Call(__e, ns2_1set, symshen_4_5alphanum_6, tmp6042)


_ = tmp6067

tmp6068 := MakeNative(func(__e *ControlFlow) {
V2567 := __e.Get(1)
_ = V2567
tmp6069 := MakeNative(func(__e *ControlFlow) {
W2568 := __e.Get(1)
_ = W2568
tmp6071 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2568)


if True == tmp6071 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2568)
return
}


}, 1)

tmp6082 := PrimIsPair(V2567)

var ifres6072 Obj

if True == tmp6082 {
tmp6073 := MakeNative(func(__e *ControlFlow) {
W2569 := __e.Get(1)
_ = W2569
tmp6074 := MakeNative(func(__e *ControlFlow) {
W2570 := __e.Get(1)
_ = W2570
tmp6077 := Call(__e, PrimFunc(symshen_4digit_2), W2569)


if True == tmp6077 {
tmp6075 := PrimNumberToString(W2569)

__e.TailApply(PrimFunc(symshen_4comb), W2570, tmp6075)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp6078 := Call(__e, PrimFunc(symtail), V2567)


__e.TailApply(tmp6074, tmp6078)
return


}, 1)

tmp6079 := Call(__e, PrimFunc(symhead), V2567)


tmp6080 := Call(__e, tmp6073, tmp6079)


ifres6072 = tmp6080


} else {
tmp6081 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6072 = tmp6081


}

__e.TailApply(tmp6069, ifres6072)
return


}, 1)

tmp6083 := Call(__e, ns2_1set, symshen_4_5numeral_6, tmp6068)


_ = tmp6083

tmp6084 := MakeNative(func(__e *ControlFlow) {
V2571 := __e.Get(1)
_ = V2571
tmp6088 := PrimGreatEqual(V2571, MakeNumber(48))

if True == tmp6088 {
tmp6086 := PrimLessEqual(V2571, MakeNumber(57))

if True == tmp6086 {
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

tmp6089 := Call(__e, ns2_1set, symshen_4digit_2, tmp6084)


_ = tmp6089

tmp6090 := MakeNative(func(__e *ControlFlow) {
V2572 := __e.Get(1)
_ = V2572
tmp6091 := MakeNative(func(__e *ControlFlow) {
W2573 := __e.Get(1)
_ = W2573
tmp6093 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2573)


if True == tmp6093 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2573)
return
}


}, 1)

tmp6094 := MakeNative(func(__e *ControlFlow) {
W2574 := __e.Get(1)
_ = W2574
tmp6112 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2574)


if True == tmp6112 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6095 := MakeNative(func(__e *ControlFlow) {
W2575 := __e.Get(1)
_ = W2575
tmp6096 := MakeNative(func(__e *ControlFlow) {
W2576 := __e.Get(1)
_ = W2576
tmp6108 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2576)


if True == tmp6108 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6097 := MakeNative(func(__e *ControlFlow) {
W2577 := __e.Get(1)
_ = W2577
tmp6098 := MakeNative(func(__e *ControlFlow) {
W2578 := __e.Get(1)
_ = W2578
tmp6099 := MakeNative(func(__e *ControlFlow) {
W2579 := __e.Get(1)
_ = W2579
tmp6103 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2579)


if True == tmp6103 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6100 := MakeNative(func(__e *ControlFlow) {
W2580 := __e.Get(1)
_ = W2580
__e.TailApply(PrimFunc(symshen_4comb), W2580, W2577)
return
}, 1)

tmp6101 := Call(__e, PrimFunc(symshen_4in_1_6), W2579)


__e.TailApply(tmp6100, tmp6101)
return


}


}, 1)

tmp6104 := Call(__e, PrimFunc(symshen_4_5dbq_6), W2578)


__e.TailApply(tmp6099, tmp6104)
return


}, 1)

tmp6105 := Call(__e, PrimFunc(symshen_4in_1_6), W2576)


__e.TailApply(tmp6098, tmp6105)
return


}, 1)

tmp6106 := Call(__e, PrimFunc(symshen_4_5_1out), W2576)


__e.TailApply(tmp6097, tmp6106)
return


}


}, 1)

tmp6109 := Call(__e, PrimFunc(symshen_4_5strcontents_6), W2575)


__e.TailApply(tmp6096, tmp6109)
return


}, 1)

tmp6110 := Call(__e, PrimFunc(symshen_4in_1_6), W2574)


__e.TailApply(tmp6095, tmp6110)
return


}


}, 1)

tmp6113 := Call(__e, PrimFunc(symshen_4_5dbq_6), V2572)


tmp6114 := Call(__e, tmp6094, tmp6113)


__e.TailApply(tmp6091, tmp6114)
return


}, 1)

tmp6115 := Call(__e, ns2_1set, symshen_4_5str_6, tmp6090)


_ = tmp6115

tmp6116 := MakeNative(func(__e *ControlFlow) {
V2581 := __e.Get(1)
_ = V2581
tmp6117 := MakeNative(func(__e *ControlFlow) {
W2582 := __e.Get(1)
_ = W2582
tmp6119 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2582)


if True == tmp6119 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2582)
return
}


}, 1)

tmp6125 := Call(__e, PrimFunc(symshen_4hds_a_2), V2581, MakeNumber(34))


var ifres6120 Obj

if True == tmp6125 {
tmp6121 := MakeNative(func(__e *ControlFlow) {
W2583 := __e.Get(1)
_ = W2583
__e.TailApply(PrimFunc(symshen_4comb), W2583, symshen_4skip)
return
}, 1)

tmp6122 := Call(__e, PrimFunc(symtail), V2581)


tmp6123 := Call(__e, tmp6121, tmp6122)


ifres6120 = tmp6123


} else {
tmp6124 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6120 = tmp6124


}

__e.TailApply(tmp6117, ifres6120)
return


}, 1)

tmp6126 := Call(__e, ns2_1set, symshen_4_5dbq_6, tmp6116)


_ = tmp6126

tmp6127 := MakeNative(func(__e *ControlFlow) {
V2584 := __e.Get(1)
_ = V2584
tmp6128 := MakeNative(func(__e *ControlFlow) {
W2585 := __e.Get(1)
_ = W2585
tmp6140 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2585)


if True == tmp6140 {
tmp6129 := MakeNative(func(__e *ControlFlow) {
W2592 := __e.Get(1)
_ = W2592
tmp6131 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2592)


if True == tmp6131 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2592)
return
}


}, 1)

tmp6132 := MakeNative(func(__e *ControlFlow) {
W2593 := __e.Get(1)
_ = W2593
tmp6136 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2593)


if True == tmp6136 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6133 := MakeNative(func(__e *ControlFlow) {
W2594 := __e.Get(1)
_ = W2594
__e.TailApply(PrimFunc(symshen_4comb), W2594, MakeString(""))
return
}, 1)

tmp6134 := Call(__e, PrimFunc(symshen_4in_1_6), W2593)


__e.TailApply(tmp6133, tmp6134)
return


}


}, 1)

tmp6137 := Call(__e, PrimFunc(sym_5e_6), V2584)


tmp6138 := Call(__e, tmp6132, tmp6137)


__e.TailApply(tmp6129, tmp6138)
return


} else {
__e.Return(W2585)
return
}


}, 1)

tmp6141 := MakeNative(func(__e *ControlFlow) {
W2586 := __e.Get(1)
_ = W2586
tmp6156 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2586)


if True == tmp6156 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6142 := MakeNative(func(__e *ControlFlow) {
W2587 := __e.Get(1)
_ = W2587
tmp6143 := MakeNative(func(__e *ControlFlow) {
W2588 := __e.Get(1)
_ = W2588
tmp6144 := MakeNative(func(__e *ControlFlow) {
W2589 := __e.Get(1)
_ = W2589
tmp6151 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2589)


if True == tmp6151 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6145 := MakeNative(func(__e *ControlFlow) {
W2590 := __e.Get(1)
_ = W2590
tmp6146 := MakeNative(func(__e *ControlFlow) {
W2591 := __e.Get(1)
_ = W2591
tmp6147 := PrimStringConcat(W2587, W2590)

__e.TailApply(PrimFunc(symshen_4comb), W2591, tmp6147)
return


}, 1)

tmp6148 := Call(__e, PrimFunc(symshen_4in_1_6), W2589)


__e.TailApply(tmp6146, tmp6148)
return


}, 1)

tmp6149 := Call(__e, PrimFunc(symshen_4_5_1out), W2589)


__e.TailApply(tmp6145, tmp6149)
return


}


}, 1)

tmp6152 := Call(__e, PrimFunc(symshen_4_5strcontents_6), W2588)


__e.TailApply(tmp6144, tmp6152)
return


}, 1)

tmp6153 := Call(__e, PrimFunc(symshen_4in_1_6), W2586)


__e.TailApply(tmp6143, tmp6153)
return


}, 1)

tmp6154 := Call(__e, PrimFunc(symshen_4_5_1out), W2586)


__e.TailApply(tmp6142, tmp6154)
return


}


}, 1)

tmp6157 := Call(__e, PrimFunc(symshen_4_5strc_6), V2584)


tmp6158 := Call(__e, tmp6141, tmp6157)


__e.TailApply(tmp6128, tmp6158)
return


}, 1)

tmp6159 := Call(__e, ns2_1set, symshen_4_5strcontents_6, tmp6127)


_ = tmp6159

tmp6160 := MakeNative(func(__e *ControlFlow) {
V2595 := __e.Get(1)
_ = V2595
tmp6161 := MakeNative(func(__e *ControlFlow) {
W2596 := __e.Get(1)
_ = W2596
tmp6175 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2596)


if True == tmp6175 {
tmp6162 := MakeNative(func(__e *ControlFlow) {
W2600 := __e.Get(1)
_ = W2600
tmp6164 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2600)


if True == tmp6164 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2600)
return
}


}, 1)

tmp6165 := MakeNative(func(__e *ControlFlow) {
W2601 := __e.Get(1)
_ = W2601
tmp6171 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2601)


if True == tmp6171 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6166 := MakeNative(func(__e *ControlFlow) {
W2602 := __e.Get(1)
_ = W2602
tmp6167 := MakeNative(func(__e *ControlFlow) {
W2603 := __e.Get(1)
_ = W2603
__e.TailApply(PrimFunc(symshen_4comb), W2603, W2602)
return
}, 1)

tmp6168 := Call(__e, PrimFunc(symshen_4in_1_6), W2601)


__e.TailApply(tmp6167, tmp6168)
return


}, 1)

tmp6169 := Call(__e, PrimFunc(symshen_4_5_1out), W2601)


__e.TailApply(tmp6166, tmp6169)
return


}


}, 1)

tmp6172 := Call(__e, PrimFunc(symshen_4_5notdbq_6), V2595)


tmp6173 := Call(__e, tmp6165, tmp6172)


__e.TailApply(tmp6162, tmp6173)
return


} else {
__e.Return(W2596)
return
}


}, 1)

tmp6176 := MakeNative(func(__e *ControlFlow) {
W2597 := __e.Get(1)
_ = W2597
tmp6182 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2597)


if True == tmp6182 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6177 := MakeNative(func(__e *ControlFlow) {
W2598 := __e.Get(1)
_ = W2598
tmp6178 := MakeNative(func(__e *ControlFlow) {
W2599 := __e.Get(1)
_ = W2599
__e.TailApply(PrimFunc(symshen_4comb), W2599, W2598)
return
}, 1)

tmp6179 := Call(__e, PrimFunc(symshen_4in_1_6), W2597)


__e.TailApply(tmp6178, tmp6179)
return


}, 1)

tmp6180 := Call(__e, PrimFunc(symshen_4_5_1out), W2597)


__e.TailApply(tmp6177, tmp6180)
return


}


}, 1)

tmp6183 := Call(__e, PrimFunc(symshen_4_5control_6), V2595)


tmp6184 := Call(__e, tmp6176, tmp6183)


__e.TailApply(tmp6161, tmp6184)
return


}, 1)

tmp6185 := Call(__e, ns2_1set, symshen_4_5strc_6, tmp6160)


_ = tmp6185

tmp6186 := MakeNative(func(__e *ControlFlow) {
V2604 := __e.Get(1)
_ = V2604
tmp6187 := MakeNative(func(__e *ControlFlow) {
W2605 := __e.Get(1)
_ = W2605
tmp6189 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2605)


if True == tmp6189 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2605)
return
}


}, 1)

tmp6190 := MakeNative(func(__e *ControlFlow) {
W2606 := __e.Get(1)
_ = W2606
tmp6215 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2606)


if True == tmp6215 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6191 := MakeNative(func(__e *ControlFlow) {
W2607 := __e.Get(1)
_ = W2607
tmp6192 := MakeNative(func(__e *ControlFlow) {
W2608 := __e.Get(1)
_ = W2608
tmp6211 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2608)


if True == tmp6211 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6193 := MakeNative(func(__e *ControlFlow) {
W2609 := __e.Get(1)
_ = W2609
tmp6194 := MakeNative(func(__e *ControlFlow) {
W2610 := __e.Get(1)
_ = W2610
tmp6207 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2610)


if True == tmp6207 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6195 := MakeNative(func(__e *ControlFlow) {
W2611 := __e.Get(1)
_ = W2611
tmp6196 := MakeNative(func(__e *ControlFlow) {
W2612 := __e.Get(1)
_ = W2612
tmp6197 := MakeNative(func(__e *ControlFlow) {
W2613 := __e.Get(1)
_ = W2613
tmp6202 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2613)


if True == tmp6202 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6198 := MakeNative(func(__e *ControlFlow) {
W2614 := __e.Get(1)
_ = W2614
tmp6199 := PrimNumberToString(W2611)

__e.TailApply(PrimFunc(symshen_4comb), W2614, tmp6199)
return


}, 1)

tmp6200 := Call(__e, PrimFunc(symshen_4in_1_6), W2613)


__e.TailApply(tmp6198, tmp6200)
return


}


}, 1)

tmp6203 := Call(__e, PrimFunc(symshen_4_5semicolon_6), W2612)


__e.TailApply(tmp6197, tmp6203)
return


}, 1)

tmp6204 := Call(__e, PrimFunc(symshen_4in_1_6), W2610)


__e.TailApply(tmp6196, tmp6204)
return


}, 1)

tmp6205 := Call(__e, PrimFunc(symshen_4_5_1out), W2610)


__e.TailApply(tmp6195, tmp6205)
return


}


}, 1)

tmp6208 := Call(__e, PrimFunc(symshen_4_5integer_6), W2609)


__e.TailApply(tmp6194, tmp6208)
return


}, 1)

tmp6209 := Call(__e, PrimFunc(symshen_4in_1_6), W2608)


__e.TailApply(tmp6193, tmp6209)
return


}


}, 1)

tmp6212 := Call(__e, PrimFunc(symshen_4_5hash_6), W2607)


__e.TailApply(tmp6192, tmp6212)
return


}, 1)

tmp6213 := Call(__e, PrimFunc(symshen_4in_1_6), W2606)


__e.TailApply(tmp6191, tmp6213)
return


}


}, 1)

tmp6216 := Call(__e, PrimFunc(symshen_4_5lowC_6), V2604)


tmp6217 := Call(__e, tmp6190, tmp6216)


__e.TailApply(tmp6187, tmp6217)
return


}, 1)

tmp6218 := Call(__e, ns2_1set, symshen_4_5control_6, tmp6186)


_ = tmp6218

tmp6219 := MakeNative(func(__e *ControlFlow) {
V2615 := __e.Get(1)
_ = V2615
tmp6220 := MakeNative(func(__e *ControlFlow) {
W2616 := __e.Get(1)
_ = W2616
tmp6222 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2616)


if True == tmp6222 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2616)
return
}


}, 1)

tmp6234 := PrimIsPair(V2615)

var ifres6223 Obj

if True == tmp6234 {
tmp6224 := MakeNative(func(__e *ControlFlow) {
W2617 := __e.Get(1)
_ = W2617
tmp6225 := MakeNative(func(__e *ControlFlow) {
W2618 := __e.Get(1)
_ = W2618
tmp6228 := PrimEqual(W2617, MakeNumber(34))

tmp6229 := PrimNot(tmp6228)

if True == tmp6229 {
tmp6226 := PrimNumberToString(W2617)

__e.TailApply(PrimFunc(symshen_4comb), W2618, tmp6226)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp6230 := Call(__e, PrimFunc(symtail), V2615)


__e.TailApply(tmp6225, tmp6230)
return


}, 1)

tmp6231 := Call(__e, PrimFunc(symhead), V2615)


tmp6232 := Call(__e, tmp6224, tmp6231)


ifres6223 = tmp6232


} else {
tmp6233 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6223 = tmp6233


}

__e.TailApply(tmp6220, ifres6223)
return


}, 1)

tmp6235 := Call(__e, ns2_1set, symshen_4_5notdbq_6, tmp6219)


_ = tmp6235

tmp6236 := MakeNative(func(__e *ControlFlow) {
V2619 := __e.Get(1)
_ = V2619
tmp6237 := MakeNative(func(__e *ControlFlow) {
W2620 := __e.Get(1)
_ = W2620
tmp6239 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2620)


if True == tmp6239 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2620)
return
}


}, 1)

tmp6245 := Call(__e, PrimFunc(symshen_4hds_a_2), V2619, MakeNumber(99))


var ifres6240 Obj

if True == tmp6245 {
tmp6241 := MakeNative(func(__e *ControlFlow) {
W2621 := __e.Get(1)
_ = W2621
__e.TailApply(PrimFunc(symshen_4comb), W2621, symshen_4skip)
return
}, 1)

tmp6242 := Call(__e, PrimFunc(symtail), V2619)


tmp6243 := Call(__e, tmp6241, tmp6242)


ifres6240 = tmp6243


} else {
tmp6244 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6240 = tmp6244


}

__e.TailApply(tmp6237, ifres6240)
return


}, 1)

tmp6246 := Call(__e, ns2_1set, symshen_4_5lowC_6, tmp6236)


_ = tmp6246

tmp6247 := MakeNative(func(__e *ControlFlow) {
V2622 := __e.Get(1)
_ = V2622
tmp6248 := MakeNative(func(__e *ControlFlow) {
W2623 := __e.Get(1)
_ = W2623
tmp6250 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2623)


if True == tmp6250 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2623)
return
}


}, 1)

tmp6256 := Call(__e, PrimFunc(symshen_4hds_a_2), V2622, MakeNumber(35))


var ifres6251 Obj

if True == tmp6256 {
tmp6252 := MakeNative(func(__e *ControlFlow) {
W2624 := __e.Get(1)
_ = W2624
__e.TailApply(PrimFunc(symshen_4comb), W2624, symshen_4skip)
return
}, 1)

tmp6253 := Call(__e, PrimFunc(symtail), V2622)


tmp6254 := Call(__e, tmp6252, tmp6253)


ifres6251 = tmp6254


} else {
tmp6255 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6251 = tmp6255


}

__e.TailApply(tmp6248, ifres6251)
return


}, 1)

tmp6257 := Call(__e, ns2_1set, symshen_4_5hash_6, tmp6247)


_ = tmp6257

tmp6258 := MakeNative(func(__e *ControlFlow) {
V2625 := __e.Get(1)
_ = V2625
tmp6259 := MakeNative(func(__e *ControlFlow) {
W2626 := __e.Get(1)
_ = W2626
tmp6315 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2626)


if True == tmp6315 {
tmp6260 := MakeNative(func(__e *ControlFlow) {
W2632 := __e.Get(1)
_ = W2632
tmp6298 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2632)


if True == tmp6298 {
tmp6261 := MakeNative(func(__e *ControlFlow) {
W2638 := __e.Get(1)
_ = W2638
tmp6287 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2638)


if True == tmp6287 {
tmp6262 := MakeNative(func(__e *ControlFlow) {
W2642 := __e.Get(1)
_ = W2642
tmp6276 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2642)


if True == tmp6276 {
tmp6263 := MakeNative(func(__e *ControlFlow) {
W2646 := __e.Get(1)
_ = W2646
tmp6265 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2646)


if True == tmp6265 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2646)
return
}


}, 1)

tmp6266 := MakeNative(func(__e *ControlFlow) {
W2647 := __e.Get(1)
_ = W2647
tmp6272 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2647)


if True == tmp6272 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6267 := MakeNative(func(__e *ControlFlow) {
W2648 := __e.Get(1)
_ = W2648
tmp6268 := MakeNative(func(__e *ControlFlow) {
W2649 := __e.Get(1)
_ = W2649
__e.TailApply(PrimFunc(symshen_4comb), W2649, W2648)
return
}, 1)

tmp6269 := Call(__e, PrimFunc(symshen_4in_1_6), W2647)


__e.TailApply(tmp6268, tmp6269)
return


}, 1)

tmp6270 := Call(__e, PrimFunc(symshen_4_5_1out), W2647)


__e.TailApply(tmp6267, tmp6270)
return


}


}, 1)

tmp6273 := Call(__e, PrimFunc(symshen_4_5integer_6), V2625)


tmp6274 := Call(__e, tmp6266, tmp6273)


__e.TailApply(tmp6263, tmp6274)
return


} else {
__e.Return(W2642)
return
}


}, 1)

tmp6277 := MakeNative(func(__e *ControlFlow) {
W2643 := __e.Get(1)
_ = W2643
tmp6283 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2643)


if True == tmp6283 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6278 := MakeNative(func(__e *ControlFlow) {
W2644 := __e.Get(1)
_ = W2644
tmp6279 := MakeNative(func(__e *ControlFlow) {
W2645 := __e.Get(1)
_ = W2645
__e.TailApply(PrimFunc(symshen_4comb), W2645, W2644)
return
}, 1)

tmp6280 := Call(__e, PrimFunc(symshen_4in_1_6), W2643)


__e.TailApply(tmp6279, tmp6280)
return


}, 1)

tmp6281 := Call(__e, PrimFunc(symshen_4_5_1out), W2643)


__e.TailApply(tmp6278, tmp6281)
return


}


}, 1)

tmp6284 := Call(__e, PrimFunc(symshen_4_5float_6), V2625)


tmp6285 := Call(__e, tmp6277, tmp6284)


__e.TailApply(tmp6262, tmp6285)
return


} else {
__e.Return(W2638)
return
}


}, 1)

tmp6288 := MakeNative(func(__e *ControlFlow) {
W2639 := __e.Get(1)
_ = W2639
tmp6294 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2639)


if True == tmp6294 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6289 := MakeNative(func(__e *ControlFlow) {
W2640 := __e.Get(1)
_ = W2640
tmp6290 := MakeNative(func(__e *ControlFlow) {
W2641 := __e.Get(1)
_ = W2641
__e.TailApply(PrimFunc(symshen_4comb), W2641, W2640)
return
}, 1)

tmp6291 := Call(__e, PrimFunc(symshen_4in_1_6), W2639)


__e.TailApply(tmp6290, tmp6291)
return


}, 1)

tmp6292 := Call(__e, PrimFunc(symshen_4_5_1out), W2639)


__e.TailApply(tmp6289, tmp6292)
return


}


}, 1)

tmp6295 := Call(__e, PrimFunc(symshen_4_5e_1number_6), V2625)


tmp6296 := Call(__e, tmp6288, tmp6295)


__e.TailApply(tmp6261, tmp6296)
return


} else {
__e.Return(W2632)
return
}


}, 1)

tmp6299 := MakeNative(func(__e *ControlFlow) {
W2633 := __e.Get(1)
_ = W2633
tmp6311 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2633)


if True == tmp6311 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6300 := MakeNative(func(__e *ControlFlow) {
W2634 := __e.Get(1)
_ = W2634
tmp6301 := MakeNative(func(__e *ControlFlow) {
W2635 := __e.Get(1)
_ = W2635
tmp6307 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2635)


if True == tmp6307 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6302 := MakeNative(func(__e *ControlFlow) {
W2636 := __e.Get(1)
_ = W2636
tmp6303 := MakeNative(func(__e *ControlFlow) {
W2637 := __e.Get(1)
_ = W2637
__e.TailApply(PrimFunc(symshen_4comb), W2637, W2636)
return
}, 1)

tmp6304 := Call(__e, PrimFunc(symshen_4in_1_6), W2635)


__e.TailApply(tmp6303, tmp6304)
return


}, 1)

tmp6305 := Call(__e, PrimFunc(symshen_4_5_1out), W2635)


__e.TailApply(tmp6302, tmp6305)
return


}


}, 1)

tmp6308 := Call(__e, PrimFunc(symshen_4_5number_6), W2634)


__e.TailApply(tmp6301, tmp6308)
return


}, 1)

tmp6309 := Call(__e, PrimFunc(symshen_4in_1_6), W2633)


__e.TailApply(tmp6300, tmp6309)
return


}


}, 1)

tmp6312 := Call(__e, PrimFunc(symshen_4_5plus_6), V2625)


tmp6313 := Call(__e, tmp6299, tmp6312)


__e.TailApply(tmp6260, tmp6313)
return


} else {
__e.Return(W2626)
return
}


}, 1)

tmp6316 := MakeNative(func(__e *ControlFlow) {
W2627 := __e.Get(1)
_ = W2627
tmp6329 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2627)


if True == tmp6329 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6317 := MakeNative(func(__e *ControlFlow) {
W2628 := __e.Get(1)
_ = W2628
tmp6318 := MakeNative(func(__e *ControlFlow) {
W2629 := __e.Get(1)
_ = W2629
tmp6325 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2629)


if True == tmp6325 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6319 := MakeNative(func(__e *ControlFlow) {
W2630 := __e.Get(1)
_ = W2630
tmp6320 := MakeNative(func(__e *ControlFlow) {
W2631 := __e.Get(1)
_ = W2631
tmp6321 := PrimNumberSubtract(MakeNumber(0), W2630)

__e.TailApply(PrimFunc(symshen_4comb), W2631, tmp6321)
return


}, 1)

tmp6322 := Call(__e, PrimFunc(symshen_4in_1_6), W2629)


__e.TailApply(tmp6320, tmp6322)
return


}, 1)

tmp6323 := Call(__e, PrimFunc(symshen_4_5_1out), W2629)


__e.TailApply(tmp6319, tmp6323)
return


}


}, 1)

tmp6326 := Call(__e, PrimFunc(symshen_4_5number_6), W2628)


__e.TailApply(tmp6318, tmp6326)
return


}, 1)

tmp6327 := Call(__e, PrimFunc(symshen_4in_1_6), W2627)


__e.TailApply(tmp6317, tmp6327)
return


}


}, 1)

tmp6330 := Call(__e, PrimFunc(symshen_4_5minus_6), V2625)


tmp6331 := Call(__e, tmp6316, tmp6330)


__e.TailApply(tmp6259, tmp6331)
return


}, 1)

tmp6332 := Call(__e, ns2_1set, symshen_4_5number_6, tmp6258)


_ = tmp6332

tmp6333 := MakeNative(func(__e *ControlFlow) {
V2650 := __e.Get(1)
_ = V2650
tmp6334 := MakeNative(func(__e *ControlFlow) {
W2651 := __e.Get(1)
_ = W2651
tmp6336 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2651)


if True == tmp6336 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2651)
return
}


}, 1)

tmp6342 := Call(__e, PrimFunc(symshen_4hds_a_2), V2650, MakeNumber(45))


var ifres6337 Obj

if True == tmp6342 {
tmp6338 := MakeNative(func(__e *ControlFlow) {
W2652 := __e.Get(1)
_ = W2652
__e.TailApply(PrimFunc(symshen_4comb), W2652, symshen_4skip)
return
}, 1)

tmp6339 := Call(__e, PrimFunc(symtail), V2650)


tmp6340 := Call(__e, tmp6338, tmp6339)


ifres6337 = tmp6340


} else {
tmp6341 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6337 = tmp6341


}

__e.TailApply(tmp6334, ifres6337)
return


}, 1)

tmp6343 := Call(__e, ns2_1set, symshen_4_5minus_6, tmp6333)


_ = tmp6343

tmp6344 := MakeNative(func(__e *ControlFlow) {
V2653 := __e.Get(1)
_ = V2653
tmp6345 := MakeNative(func(__e *ControlFlow) {
W2654 := __e.Get(1)
_ = W2654
tmp6347 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2654)


if True == tmp6347 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2654)
return
}


}, 1)

tmp6353 := Call(__e, PrimFunc(symshen_4hds_a_2), V2653, MakeNumber(43))


var ifres6348 Obj

if True == tmp6353 {
tmp6349 := MakeNative(func(__e *ControlFlow) {
W2655 := __e.Get(1)
_ = W2655
__e.TailApply(PrimFunc(symshen_4comb), W2655, symshen_4skip)
return
}, 1)

tmp6350 := Call(__e, PrimFunc(symtail), V2653)


tmp6351 := Call(__e, tmp6349, tmp6350)


ifres6348 = tmp6351


} else {
tmp6352 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6348 = tmp6352


}

__e.TailApply(tmp6345, ifres6348)
return


}, 1)

tmp6354 := Call(__e, ns2_1set, symshen_4_5plus_6, tmp6344)


_ = tmp6354

tmp6355 := MakeNative(func(__e *ControlFlow) {
V2656 := __e.Get(1)
_ = V2656
tmp6356 := MakeNative(func(__e *ControlFlow) {
W2657 := __e.Get(1)
_ = W2657
tmp6358 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2657)


if True == tmp6358 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2657)
return
}


}, 1)

tmp6359 := MakeNative(func(__e *ControlFlow) {
W2658 := __e.Get(1)
_ = W2658
tmp6366 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2658)


if True == tmp6366 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6360 := MakeNative(func(__e *ControlFlow) {
W2659 := __e.Get(1)
_ = W2659
tmp6361 := MakeNative(func(__e *ControlFlow) {
W2660 := __e.Get(1)
_ = W2660
tmp6362 := Call(__e, PrimFunc(symshen_4compute_1integer), W2659)


__e.TailApply(PrimFunc(symshen_4comb), W2660, tmp6362)
return


}, 1)

tmp6363 := Call(__e, PrimFunc(symshen_4in_1_6), W2658)


__e.TailApply(tmp6361, tmp6363)
return


}, 1)

tmp6364 := Call(__e, PrimFunc(symshen_4_5_1out), W2658)


__e.TailApply(tmp6360, tmp6364)
return


}


}, 1)

tmp6367 := Call(__e, PrimFunc(symshen_4_5digits_6), V2656)


tmp6368 := Call(__e, tmp6359, tmp6367)


__e.TailApply(tmp6356, tmp6368)
return


}, 1)

tmp6369 := Call(__e, ns2_1set, symshen_4_5integer_6, tmp6355)


_ = tmp6369

tmp6370 := MakeNative(func(__e *ControlFlow) {
V2661 := __e.Get(1)
_ = V2661
tmp6371 := MakeNative(func(__e *ControlFlow) {
W2662 := __e.Get(1)
_ = W2662
tmp6386 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2662)


if True == tmp6386 {
tmp6372 := MakeNative(func(__e *ControlFlow) {
W2669 := __e.Get(1)
_ = W2669
tmp6374 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2669)


if True == tmp6374 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2669)
return
}


}, 1)

tmp6375 := MakeNative(func(__e *ControlFlow) {
W2670 := __e.Get(1)
_ = W2670
tmp6382 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2670)


if True == tmp6382 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6376 := MakeNative(func(__e *ControlFlow) {
W2671 := __e.Get(1)
_ = W2671
tmp6377 := MakeNative(func(__e *ControlFlow) {
W2672 := __e.Get(1)
_ = W2672
tmp6378 := PrimCons(W2671, Nil)

__e.TailApply(PrimFunc(symshen_4comb), W2672, tmp6378)
return


}, 1)

tmp6379 := Call(__e, PrimFunc(symshen_4in_1_6), W2670)


__e.TailApply(tmp6377, tmp6379)
return


}, 1)

tmp6380 := Call(__e, PrimFunc(symshen_4_5_1out), W2670)


__e.TailApply(tmp6376, tmp6380)
return


}


}, 1)

tmp6383 := Call(__e, PrimFunc(symshen_4_5digit_6), V2661)


tmp6384 := Call(__e, tmp6375, tmp6383)


__e.TailApply(tmp6372, tmp6384)
return


} else {
__e.Return(W2662)
return
}


}, 1)

tmp6387 := MakeNative(func(__e *ControlFlow) {
W2663 := __e.Get(1)
_ = W2663
tmp6402 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2663)


if True == tmp6402 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6388 := MakeNative(func(__e *ControlFlow) {
W2664 := __e.Get(1)
_ = W2664
tmp6389 := MakeNative(func(__e *ControlFlow) {
W2665 := __e.Get(1)
_ = W2665
tmp6390 := MakeNative(func(__e *ControlFlow) {
W2666 := __e.Get(1)
_ = W2666
tmp6397 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2666)


if True == tmp6397 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6391 := MakeNative(func(__e *ControlFlow) {
W2667 := __e.Get(1)
_ = W2667
tmp6392 := MakeNative(func(__e *ControlFlow) {
W2668 := __e.Get(1)
_ = W2668
tmp6393 := PrimCons(W2664, W2667)

__e.TailApply(PrimFunc(symshen_4comb), W2668, tmp6393)
return


}, 1)

tmp6394 := Call(__e, PrimFunc(symshen_4in_1_6), W2666)


__e.TailApply(tmp6392, tmp6394)
return


}, 1)

tmp6395 := Call(__e, PrimFunc(symshen_4_5_1out), W2666)


__e.TailApply(tmp6391, tmp6395)
return


}


}, 1)

tmp6398 := Call(__e, PrimFunc(symshen_4_5digits_6), W2665)


__e.TailApply(tmp6390, tmp6398)
return


}, 1)

tmp6399 := Call(__e, PrimFunc(symshen_4in_1_6), W2663)


__e.TailApply(tmp6389, tmp6399)
return


}, 1)

tmp6400 := Call(__e, PrimFunc(symshen_4_5_1out), W2663)


__e.TailApply(tmp6388, tmp6400)
return


}


}, 1)

tmp6403 := Call(__e, PrimFunc(symshen_4_5digit_6), V2661)


tmp6404 := Call(__e, tmp6387, tmp6403)


__e.TailApply(tmp6371, tmp6404)
return


}, 1)

tmp6405 := Call(__e, ns2_1set, symshen_4_5digits_6, tmp6370)


_ = tmp6405

tmp6406 := MakeNative(func(__e *ControlFlow) {
V2673 := __e.Get(1)
_ = V2673
tmp6407 := MakeNative(func(__e *ControlFlow) {
W2674 := __e.Get(1)
_ = W2674
tmp6409 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2674)


if True == tmp6409 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2674)
return
}


}, 1)

tmp6420 := PrimIsPair(V2673)

var ifres6410 Obj

if True == tmp6420 {
tmp6411 := MakeNative(func(__e *ControlFlow) {
W2675 := __e.Get(1)
_ = W2675
tmp6412 := MakeNative(func(__e *ControlFlow) {
W2676 := __e.Get(1)
_ = W2676
tmp6415 := Call(__e, PrimFunc(symshen_4digit_2), W2675)


if True == tmp6415 {
tmp6413 := Call(__e, PrimFunc(symshen_4byte_1_6digit), W2675)


__e.TailApply(PrimFunc(symshen_4comb), W2676, tmp6413)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp6416 := Call(__e, PrimFunc(symtail), V2673)


__e.TailApply(tmp6412, tmp6416)
return


}, 1)

tmp6417 := Call(__e, PrimFunc(symhead), V2673)


tmp6418 := Call(__e, tmp6411, tmp6417)


ifres6410 = tmp6418


} else {
tmp6419 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6410 = tmp6419


}

__e.TailApply(tmp6407, ifres6410)
return


}, 1)

tmp6421 := Call(__e, ns2_1set, symshen_4_5digit_6, tmp6406)


_ = tmp6421

tmp6422 := MakeNative(func(__e *ControlFlow) {
V2677 := __e.Get(1)
_ = V2677
__e.Return(PrimNumberSubtract(V2677, MakeNumber(48)))
return
}, 1)

tmp6423 := Call(__e, ns2_1set, symshen_4byte_1_6digit, tmp6422)


_ = tmp6423

tmp6424 := MakeNative(func(__e *ControlFlow) {
V2678 := __e.Get(1)
_ = V2678
tmp6425 := Call(__e, PrimFunc(symreverse), V2678)


__e.TailApply(PrimFunc(symshen_4compute_1integer_1h), tmp6425, MakeNumber(0))
return


}, 1)

tmp6426 := Call(__e, ns2_1set, symshen_4compute_1integer, tmp6424)


_ = tmp6426

tmp6427 := MakeNative(func(__e *ControlFlow) {
V2681 := __e.Get(1)
_ = V2681
V2682 := __e.Get(2)
_ = V2682
tmp6437 := PrimEqual(Nil, V2681)

if True == tmp6437 {
__e.Return(MakeNumber(0))
return
} else {
tmp6435 := PrimIsPair(V2681)

if True == tmp6435 {
tmp6428 := Call(__e, PrimFunc(symshen_4expt), MakeNumber(10), V2682)


tmp6429 := PrimHead(V2681)

tmp6430 := PrimNumberMultiply(tmp6428, tmp6429)

tmp6431 := PrimTail(V2681)

tmp6432 := PrimNumberAdd(V2682, MakeNumber(1))

tmp6433 := Call(__e, PrimFunc(symshen_4compute_1integer_1h), tmp6431, tmp6432)


__e.Return(PrimNumberAdd(tmp6430, tmp6433))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.compute-integer-h")))
return
}


}


}, 2)

tmp6438 := Call(__e, ns2_1set, symshen_4compute_1integer_1h, tmp6427)


_ = tmp6438

tmp6439 := MakeNative(func(__e *ControlFlow) {
V2685 := __e.Get(1)
_ = V2685
V2686 := __e.Get(2)
_ = V2686
tmp6447 := PrimEqual(MakeNumber(0), V2686)

if True == tmp6447 {
__e.Return(MakeNumber(1))
return
} else {
tmp6445 := PrimGreatThan(V2686, MakeNumber(0))

if True == tmp6445 {
tmp6440 := PrimNumberSubtract(V2686, MakeNumber(1))

tmp6441 := Call(__e, PrimFunc(symshen_4expt), V2685, tmp6440)


__e.Return(PrimNumberMultiply(V2685, tmp6441))
return


} else {
tmp6442 := PrimNumberAdd(V2686, MakeNumber(1))

tmp6443 := Call(__e, PrimFunc(symshen_4expt), V2685, tmp6442)


__e.Return(PrimNumberDivide(tmp6443, V2685))
return


}


}


}, 2)

tmp6448 := Call(__e, ns2_1set, symshen_4expt, tmp6439)


_ = tmp6448

tmp6449 := MakeNative(func(__e *ControlFlow) {
V2687 := __e.Get(1)
_ = V2687
tmp6450 := MakeNative(func(__e *ControlFlow) {
W2688 := __e.Get(1)
_ = W2688
tmp6470 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2688)


if True == tmp6470 {
tmp6451 := MakeNative(func(__e *ControlFlow) {
W2697 := __e.Get(1)
_ = W2697
tmp6453 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2697)


if True == tmp6453 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2697)
return
}


}, 1)

tmp6454 := MakeNative(func(__e *ControlFlow) {
W2698 := __e.Get(1)
_ = W2698
tmp6466 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2698)


if True == tmp6466 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6455 := MakeNative(func(__e *ControlFlow) {
W2699 := __e.Get(1)
_ = W2699
tmp6456 := MakeNative(func(__e *ControlFlow) {
W2700 := __e.Get(1)
_ = W2700
tmp6462 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2700)


if True == tmp6462 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6457 := MakeNative(func(__e *ControlFlow) {
W2701 := __e.Get(1)
_ = W2701
tmp6458 := MakeNative(func(__e *ControlFlow) {
W2702 := __e.Get(1)
_ = W2702
__e.TailApply(PrimFunc(symshen_4comb), W2702, W2701)
return
}, 1)

tmp6459 := Call(__e, PrimFunc(symshen_4in_1_6), W2700)


__e.TailApply(tmp6458, tmp6459)
return


}, 1)

tmp6460 := Call(__e, PrimFunc(symshen_4_5_1out), W2700)


__e.TailApply(tmp6457, tmp6460)
return


}


}, 1)

tmp6463 := Call(__e, PrimFunc(symshen_4_5fraction_6), W2699)


__e.TailApply(tmp6456, tmp6463)
return


}, 1)

tmp6464 := Call(__e, PrimFunc(symshen_4in_1_6), W2698)


__e.TailApply(tmp6455, tmp6464)
return


}


}, 1)

tmp6467 := Call(__e, PrimFunc(symshen_4_5stop_6), V2687)


tmp6468 := Call(__e, tmp6454, tmp6467)


__e.TailApply(tmp6451, tmp6468)
return


} else {
__e.Return(W2688)
return
}


}, 1)

tmp6471 := MakeNative(func(__e *ControlFlow) {
W2689 := __e.Get(1)
_ = W2689
tmp6492 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2689)


if True == tmp6492 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6472 := MakeNative(func(__e *ControlFlow) {
W2690 := __e.Get(1)
_ = W2690
tmp6473 := MakeNative(func(__e *ControlFlow) {
W2691 := __e.Get(1)
_ = W2691
tmp6474 := MakeNative(func(__e *ControlFlow) {
W2692 := __e.Get(1)
_ = W2692
tmp6487 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2692)


if True == tmp6487 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6475 := MakeNative(func(__e *ControlFlow) {
W2693 := __e.Get(1)
_ = W2693
tmp6476 := MakeNative(func(__e *ControlFlow) {
W2694 := __e.Get(1)
_ = W2694
tmp6483 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2694)


if True == tmp6483 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6477 := MakeNative(func(__e *ControlFlow) {
W2695 := __e.Get(1)
_ = W2695
tmp6478 := MakeNative(func(__e *ControlFlow) {
W2696 := __e.Get(1)
_ = W2696
tmp6479 := PrimNumberAdd(W2690, W2695)

__e.TailApply(PrimFunc(symshen_4comb), W2696, tmp6479)
return


}, 1)

tmp6480 := Call(__e, PrimFunc(symshen_4in_1_6), W2694)


__e.TailApply(tmp6478, tmp6480)
return


}, 1)

tmp6481 := Call(__e, PrimFunc(symshen_4_5_1out), W2694)


__e.TailApply(tmp6477, tmp6481)
return


}


}, 1)

tmp6484 := Call(__e, PrimFunc(symshen_4_5fraction_6), W2693)


__e.TailApply(tmp6476, tmp6484)
return


}, 1)

tmp6485 := Call(__e, PrimFunc(symshen_4in_1_6), W2692)


__e.TailApply(tmp6475, tmp6485)
return


}


}, 1)

tmp6488 := Call(__e, PrimFunc(symshen_4_5stop_6), W2691)


__e.TailApply(tmp6474, tmp6488)
return


}, 1)

tmp6489 := Call(__e, PrimFunc(symshen_4in_1_6), W2689)


__e.TailApply(tmp6473, tmp6489)
return


}, 1)

tmp6490 := Call(__e, PrimFunc(symshen_4_5_1out), W2689)


__e.TailApply(tmp6472, tmp6490)
return


}


}, 1)

tmp6493 := Call(__e, PrimFunc(symshen_4_5integer_6), V2687)


tmp6494 := Call(__e, tmp6471, tmp6493)


__e.TailApply(tmp6450, tmp6494)
return


}, 1)

tmp6495 := Call(__e, ns2_1set, symshen_4_5float_6, tmp6449)


_ = tmp6495

tmp6496 := MakeNative(func(__e *ControlFlow) {
V2703 := __e.Get(1)
_ = V2703
tmp6497 := MakeNative(func(__e *ControlFlow) {
W2704 := __e.Get(1)
_ = W2704
tmp6499 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2704)


if True == tmp6499 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2704)
return
}


}, 1)

tmp6505 := Call(__e, PrimFunc(symshen_4hds_a_2), V2703, MakeNumber(46))


var ifres6500 Obj

if True == tmp6505 {
tmp6501 := MakeNative(func(__e *ControlFlow) {
W2705 := __e.Get(1)
_ = W2705
__e.TailApply(PrimFunc(symshen_4comb), W2705, symshen_4skip)
return
}, 1)

tmp6502 := Call(__e, PrimFunc(symtail), V2703)


tmp6503 := Call(__e, tmp6501, tmp6502)


ifres6500 = tmp6503


} else {
tmp6504 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6500 = tmp6504


}

__e.TailApply(tmp6497, ifres6500)
return


}, 1)

tmp6506 := Call(__e, ns2_1set, symshen_4_5stop_6, tmp6496)


_ = tmp6506

tmp6507 := MakeNative(func(__e *ControlFlow) {
V2706 := __e.Get(1)
_ = V2706
tmp6508 := MakeNative(func(__e *ControlFlow) {
W2707 := __e.Get(1)
_ = W2707
tmp6510 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2707)


if True == tmp6510 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2707)
return
}


}, 1)

tmp6511 := MakeNative(func(__e *ControlFlow) {
W2708 := __e.Get(1)
_ = W2708
tmp6518 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2708)


if True == tmp6518 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6512 := MakeNative(func(__e *ControlFlow) {
W2709 := __e.Get(1)
_ = W2709
tmp6513 := MakeNative(func(__e *ControlFlow) {
W2710 := __e.Get(1)
_ = W2710
tmp6514 := Call(__e, PrimFunc(symshen_4compute_1fraction), W2709)


__e.TailApply(PrimFunc(symshen_4comb), W2710, tmp6514)
return


}, 1)

tmp6515 := Call(__e, PrimFunc(symshen_4in_1_6), W2708)


__e.TailApply(tmp6513, tmp6515)
return


}, 1)

tmp6516 := Call(__e, PrimFunc(symshen_4_5_1out), W2708)


__e.TailApply(tmp6512, tmp6516)
return


}


}, 1)

tmp6519 := Call(__e, PrimFunc(symshen_4_5digits_6), V2706)


tmp6520 := Call(__e, tmp6511, tmp6519)


__e.TailApply(tmp6508, tmp6520)
return


}, 1)

tmp6521 := Call(__e, ns2_1set, symshen_4_5fraction_6, tmp6507)


_ = tmp6521

tmp6522 := MakeNative(func(__e *ControlFlow) {
V2711 := __e.Get(1)
_ = V2711
__e.TailApply(PrimFunc(symshen_4compute_1fraction_1h), V2711, MakeNumber(-1))
return
}, 1)

tmp6523 := Call(__e, ns2_1set, symshen_4compute_1fraction, tmp6522)


_ = tmp6523

tmp6524 := MakeNative(func(__e *ControlFlow) {
V2714 := __e.Get(1)
_ = V2714
V2715 := __e.Get(2)
_ = V2715
tmp6534 := PrimEqual(Nil, V2714)

if True == tmp6534 {
__e.Return(MakeNumber(0))
return
} else {
tmp6532 := PrimIsPair(V2714)

if True == tmp6532 {
tmp6525 := Call(__e, PrimFunc(symshen_4expt), MakeNumber(10), V2715)


tmp6526 := PrimHead(V2714)

tmp6527 := PrimNumberMultiply(tmp6525, tmp6526)

tmp6528 := PrimTail(V2714)

tmp6529 := PrimNumberSubtract(V2715, MakeNumber(1))

tmp6530 := Call(__e, PrimFunc(symshen_4compute_1fraction_1h), tmp6528, tmp6529)


__e.Return(PrimNumberAdd(tmp6527, tmp6530))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.compute-fraction-h")))
return
}


}


}, 2)

tmp6535 := Call(__e, ns2_1set, symshen_4compute_1fraction_1h, tmp6524)


_ = tmp6535

tmp6536 := MakeNative(func(__e *ControlFlow) {
V2716 := __e.Get(1)
_ = V2716
tmp6537 := MakeNative(func(__e *ControlFlow) {
W2717 := __e.Get(1)
_ = W2717
tmp6566 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2717)


if True == tmp6566 {
tmp6538 := MakeNative(func(__e *ControlFlow) {
W2726 := __e.Get(1)
_ = W2726
tmp6540 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2726)


if True == tmp6540 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2726)
return
}


}, 1)

tmp6541 := MakeNative(func(__e *ControlFlow) {
W2727 := __e.Get(1)
_ = W2727
tmp6562 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2727)


if True == tmp6562 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6542 := MakeNative(func(__e *ControlFlow) {
W2728 := __e.Get(1)
_ = W2728
tmp6543 := MakeNative(func(__e *ControlFlow) {
W2729 := __e.Get(1)
_ = W2729
tmp6544 := MakeNative(func(__e *ControlFlow) {
W2730 := __e.Get(1)
_ = W2730
tmp6557 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2730)


if True == tmp6557 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6545 := MakeNative(func(__e *ControlFlow) {
W2731 := __e.Get(1)
_ = W2731
tmp6546 := MakeNative(func(__e *ControlFlow) {
W2732 := __e.Get(1)
_ = W2732
tmp6553 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2732)


if True == tmp6553 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6547 := MakeNative(func(__e *ControlFlow) {
W2733 := __e.Get(1)
_ = W2733
tmp6548 := MakeNative(func(__e *ControlFlow) {
W2734 := __e.Get(1)
_ = W2734
tmp6549 := Call(__e, PrimFunc(symshen_4compute_1E), W2728, W2733)


__e.TailApply(PrimFunc(symshen_4comb), W2734, tmp6549)
return


}, 1)

tmp6550 := Call(__e, PrimFunc(symshen_4in_1_6), W2732)


__e.TailApply(tmp6548, tmp6550)
return


}, 1)

tmp6551 := Call(__e, PrimFunc(symshen_4_5_1out), W2732)


__e.TailApply(tmp6547, tmp6551)
return


}


}, 1)

tmp6554 := Call(__e, PrimFunc(symshen_4_5log10_6), W2731)


__e.TailApply(tmp6546, tmp6554)
return


}, 1)

tmp6555 := Call(__e, PrimFunc(symshen_4in_1_6), W2730)


__e.TailApply(tmp6545, tmp6555)
return


}


}, 1)

tmp6558 := Call(__e, PrimFunc(symshen_4_5lowE_6), W2729)


__e.TailApply(tmp6544, tmp6558)
return


}, 1)

tmp6559 := Call(__e, PrimFunc(symshen_4in_1_6), W2727)


__e.TailApply(tmp6543, tmp6559)
return


}, 1)

tmp6560 := Call(__e, PrimFunc(symshen_4_5_1out), W2727)


__e.TailApply(tmp6542, tmp6560)
return


}


}, 1)

tmp6563 := Call(__e, PrimFunc(symshen_4_5integer_6), V2716)


tmp6564 := Call(__e, tmp6541, tmp6563)


__e.TailApply(tmp6538, tmp6564)
return


} else {
__e.Return(W2717)
return
}


}, 1)

tmp6567 := MakeNative(func(__e *ControlFlow) {
W2718 := __e.Get(1)
_ = W2718
tmp6588 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2718)


if True == tmp6588 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6568 := MakeNative(func(__e *ControlFlow) {
W2719 := __e.Get(1)
_ = W2719
tmp6569 := MakeNative(func(__e *ControlFlow) {
W2720 := __e.Get(1)
_ = W2720
tmp6570 := MakeNative(func(__e *ControlFlow) {
W2721 := __e.Get(1)
_ = W2721
tmp6583 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2721)


if True == tmp6583 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6571 := MakeNative(func(__e *ControlFlow) {
W2722 := __e.Get(1)
_ = W2722
tmp6572 := MakeNative(func(__e *ControlFlow) {
W2723 := __e.Get(1)
_ = W2723
tmp6579 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2723)


if True == tmp6579 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6573 := MakeNative(func(__e *ControlFlow) {
W2724 := __e.Get(1)
_ = W2724
tmp6574 := MakeNative(func(__e *ControlFlow) {
W2725 := __e.Get(1)
_ = W2725
tmp6575 := Call(__e, PrimFunc(symshen_4compute_1E), W2719, W2724)


__e.TailApply(PrimFunc(symshen_4comb), W2725, tmp6575)
return


}, 1)

tmp6576 := Call(__e, PrimFunc(symshen_4in_1_6), W2723)


__e.TailApply(tmp6574, tmp6576)
return


}, 1)

tmp6577 := Call(__e, PrimFunc(symshen_4_5_1out), W2723)


__e.TailApply(tmp6573, tmp6577)
return


}


}, 1)

tmp6580 := Call(__e, PrimFunc(symshen_4_5log10_6), W2722)


__e.TailApply(tmp6572, tmp6580)
return


}, 1)

tmp6581 := Call(__e, PrimFunc(symshen_4in_1_6), W2721)


__e.TailApply(tmp6571, tmp6581)
return


}


}, 1)

tmp6584 := Call(__e, PrimFunc(symshen_4_5lowE_6), W2720)


__e.TailApply(tmp6570, tmp6584)
return


}, 1)

tmp6585 := Call(__e, PrimFunc(symshen_4in_1_6), W2718)


__e.TailApply(tmp6569, tmp6585)
return


}, 1)

tmp6586 := Call(__e, PrimFunc(symshen_4_5_1out), W2718)


__e.TailApply(tmp6568, tmp6586)
return


}


}, 1)

tmp6589 := Call(__e, PrimFunc(symshen_4_5float_6), V2716)


tmp6590 := Call(__e, tmp6567, tmp6589)


__e.TailApply(tmp6537, tmp6590)
return


}, 1)

tmp6591 := Call(__e, ns2_1set, symshen_4_5e_1number_6, tmp6536)


_ = tmp6591

tmp6592 := MakeNative(func(__e *ControlFlow) {
V2735 := __e.Get(1)
_ = V2735
tmp6593 := MakeNative(func(__e *ControlFlow) {
W2736 := __e.Get(1)
_ = W2736
tmp6626 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2736)


if True == tmp6626 {
tmp6594 := MakeNative(func(__e *ControlFlow) {
W2742 := __e.Get(1)
_ = W2742
tmp6608 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2742)


if True == tmp6608 {
tmp6595 := MakeNative(func(__e *ControlFlow) {
W2748 := __e.Get(1)
_ = W2748
tmp6597 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2748)


if True == tmp6597 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2748)
return
}


}, 1)

tmp6598 := MakeNative(func(__e *ControlFlow) {
W2749 := __e.Get(1)
_ = W2749
tmp6604 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2749)


if True == tmp6604 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6599 := MakeNative(func(__e *ControlFlow) {
W2750 := __e.Get(1)
_ = W2750
tmp6600 := MakeNative(func(__e *ControlFlow) {
W2751 := __e.Get(1)
_ = W2751
__e.TailApply(PrimFunc(symshen_4comb), W2751, W2750)
return
}, 1)

tmp6601 := Call(__e, PrimFunc(symshen_4in_1_6), W2749)


__e.TailApply(tmp6600, tmp6601)
return


}, 1)

tmp6602 := Call(__e, PrimFunc(symshen_4_5_1out), W2749)


__e.TailApply(tmp6599, tmp6602)
return


}


}, 1)

tmp6605 := Call(__e, PrimFunc(symshen_4_5integer_6), V2735)


tmp6606 := Call(__e, tmp6598, tmp6605)


__e.TailApply(tmp6595, tmp6606)
return


} else {
__e.Return(W2742)
return
}


}, 1)

tmp6609 := MakeNative(func(__e *ControlFlow) {
W2743 := __e.Get(1)
_ = W2743
tmp6622 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2743)


if True == tmp6622 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6610 := MakeNative(func(__e *ControlFlow) {
W2744 := __e.Get(1)
_ = W2744
tmp6611 := MakeNative(func(__e *ControlFlow) {
W2745 := __e.Get(1)
_ = W2745
tmp6618 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2745)


if True == tmp6618 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6612 := MakeNative(func(__e *ControlFlow) {
W2746 := __e.Get(1)
_ = W2746
tmp6613 := MakeNative(func(__e *ControlFlow) {
W2747 := __e.Get(1)
_ = W2747
tmp6614 := PrimNumberSubtract(MakeNumber(0), W2746)

__e.TailApply(PrimFunc(symshen_4comb), W2747, tmp6614)
return


}, 1)

tmp6615 := Call(__e, PrimFunc(symshen_4in_1_6), W2745)


__e.TailApply(tmp6613, tmp6615)
return


}, 1)

tmp6616 := Call(__e, PrimFunc(symshen_4_5_1out), W2745)


__e.TailApply(tmp6612, tmp6616)
return


}


}, 1)

tmp6619 := Call(__e, PrimFunc(symshen_4_5log10_6), W2744)


__e.TailApply(tmp6611, tmp6619)
return


}, 1)

tmp6620 := Call(__e, PrimFunc(symshen_4in_1_6), W2743)


__e.TailApply(tmp6610, tmp6620)
return


}


}, 1)

tmp6623 := Call(__e, PrimFunc(symshen_4_5minus_6), V2735)


tmp6624 := Call(__e, tmp6609, tmp6623)


__e.TailApply(tmp6594, tmp6624)
return


} else {
__e.Return(W2736)
return
}


}, 1)

tmp6627 := MakeNative(func(__e *ControlFlow) {
W2737 := __e.Get(1)
_ = W2737
tmp6639 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2737)


if True == tmp6639 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6628 := MakeNative(func(__e *ControlFlow) {
W2738 := __e.Get(1)
_ = W2738
tmp6629 := MakeNative(func(__e *ControlFlow) {
W2739 := __e.Get(1)
_ = W2739
tmp6635 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2739)


if True == tmp6635 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6630 := MakeNative(func(__e *ControlFlow) {
W2740 := __e.Get(1)
_ = W2740
tmp6631 := MakeNative(func(__e *ControlFlow) {
W2741 := __e.Get(1)
_ = W2741
__e.TailApply(PrimFunc(symshen_4comb), W2741, W2740)
return
}, 1)

tmp6632 := Call(__e, PrimFunc(symshen_4in_1_6), W2739)


__e.TailApply(tmp6631, tmp6632)
return


}, 1)

tmp6633 := Call(__e, PrimFunc(symshen_4_5_1out), W2739)


__e.TailApply(tmp6630, tmp6633)
return


}


}, 1)

tmp6636 := Call(__e, PrimFunc(symshen_4_5log10_6), W2738)


__e.TailApply(tmp6629, tmp6636)
return


}, 1)

tmp6637 := Call(__e, PrimFunc(symshen_4in_1_6), W2737)


__e.TailApply(tmp6628, tmp6637)
return


}


}, 1)

tmp6640 := Call(__e, PrimFunc(symshen_4_5plus_6), V2735)


tmp6641 := Call(__e, tmp6627, tmp6640)


__e.TailApply(tmp6593, tmp6641)
return


}, 1)

tmp6642 := Call(__e, ns2_1set, symshen_4_5log10_6, tmp6592)


_ = tmp6642

tmp6643 := MakeNative(func(__e *ControlFlow) {
V2752 := __e.Get(1)
_ = V2752
tmp6644 := MakeNative(func(__e *ControlFlow) {
W2753 := __e.Get(1)
_ = W2753
tmp6646 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2753)


if True == tmp6646 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2753)
return
}


}, 1)

tmp6652 := Call(__e, PrimFunc(symshen_4hds_a_2), V2752, MakeNumber(101))


var ifres6647 Obj

if True == tmp6652 {
tmp6648 := MakeNative(func(__e *ControlFlow) {
W2754 := __e.Get(1)
_ = W2754
__e.TailApply(PrimFunc(symshen_4comb), W2754, symshen_4skip)
return
}, 1)

tmp6649 := Call(__e, PrimFunc(symtail), V2752)


tmp6650 := Call(__e, tmp6648, tmp6649)


ifres6647 = tmp6650


} else {
tmp6651 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6647 = tmp6651


}

__e.TailApply(tmp6644, ifres6647)
return


}, 1)

tmp6653 := Call(__e, ns2_1set, symshen_4_5lowE_6, tmp6643)


_ = tmp6653

tmp6654 := MakeNative(func(__e *ControlFlow) {
V2755 := __e.Get(1)
_ = V2755
V2756 := __e.Get(2)
_ = V2756
tmp6655 := Call(__e, PrimFunc(symshen_4expt), MakeNumber(10), V2756)


__e.Return(PrimNumberMultiply(V2755, tmp6655))
return


}, 2)

tmp6656 := Call(__e, ns2_1set, symshen_4compute_1E, tmp6654)


_ = tmp6656

tmp6657 := MakeNative(func(__e *ControlFlow) {
V2757 := __e.Get(1)
_ = V2757
tmp6658 := MakeNative(func(__e *ControlFlow) {
W2758 := __e.Get(1)
_ = W2758
tmp6670 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2758)


if True == tmp6670 {
tmp6659 := MakeNative(func(__e *ControlFlow) {
W2763 := __e.Get(1)
_ = W2763
tmp6661 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2763)


if True == tmp6661 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2763)
return
}


}, 1)

tmp6662 := MakeNative(func(__e *ControlFlow) {
W2764 := __e.Get(1)
_ = W2764
tmp6666 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2764)


if True == tmp6666 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6663 := MakeNative(func(__e *ControlFlow) {
W2765 := __e.Get(1)
_ = W2765
__e.TailApply(PrimFunc(symshen_4comb), W2765, symshen_4skip)
return
}, 1)

tmp6664 := Call(__e, PrimFunc(symshen_4in_1_6), W2764)


__e.TailApply(tmp6663, tmp6664)
return


}


}, 1)

tmp6667 := Call(__e, PrimFunc(symshen_4_5whitespace_6), V2757)


tmp6668 := Call(__e, tmp6662, tmp6667)


__e.TailApply(tmp6659, tmp6668)
return


} else {
__e.Return(W2758)
return
}


}, 1)

tmp6671 := MakeNative(func(__e *ControlFlow) {
W2759 := __e.Get(1)
_ = W2759
tmp6681 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2759)


if True == tmp6681 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6672 := MakeNative(func(__e *ControlFlow) {
W2760 := __e.Get(1)
_ = W2760
tmp6673 := MakeNative(func(__e *ControlFlow) {
W2761 := __e.Get(1)
_ = W2761
tmp6677 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2761)


if True == tmp6677 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6674 := MakeNative(func(__e *ControlFlow) {
W2762 := __e.Get(1)
_ = W2762
__e.TailApply(PrimFunc(symshen_4comb), W2762, symshen_4skip)
return
}, 1)

tmp6675 := Call(__e, PrimFunc(symshen_4in_1_6), W2761)


__e.TailApply(tmp6674, tmp6675)
return


}


}, 1)

tmp6678 := Call(__e, PrimFunc(symshen_4_5whitespaces_6), W2760)


__e.TailApply(tmp6673, tmp6678)
return


}, 1)

tmp6679 := Call(__e, PrimFunc(symshen_4in_1_6), W2759)


__e.TailApply(tmp6672, tmp6679)
return


}


}, 1)

tmp6682 := Call(__e, PrimFunc(symshen_4_5whitespace_6), V2757)


tmp6683 := Call(__e, tmp6671, tmp6682)


__e.TailApply(tmp6658, tmp6683)
return


}, 1)

tmp6684 := Call(__e, ns2_1set, symshen_4_5whitespaces_6, tmp6657)


_ = tmp6684

tmp6685 := MakeNative(func(__e *ControlFlow) {
V2766 := __e.Get(1)
_ = V2766
tmp6686 := MakeNative(func(__e *ControlFlow) {
W2767 := __e.Get(1)
_ = W2767
tmp6688 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2767)


if True == tmp6688 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2767)
return
}


}, 1)

tmp6698 := PrimIsPair(V2766)

var ifres6689 Obj

if True == tmp6698 {
tmp6690 := MakeNative(func(__e *ControlFlow) {
W2768 := __e.Get(1)
_ = W2768
tmp6691 := MakeNative(func(__e *ControlFlow) {
W2769 := __e.Get(1)
_ = W2769
tmp6693 := Call(__e, PrimFunc(symshen_4whitespace_2), W2768)


if True == tmp6693 {
__e.TailApply(PrimFunc(symshen_4comb), W2769, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp6694 := Call(__e, PrimFunc(symtail), V2766)


__e.TailApply(tmp6691, tmp6694)
return


}, 1)

tmp6695 := Call(__e, PrimFunc(symhead), V2766)


tmp6696 := Call(__e, tmp6690, tmp6695)


ifres6689 = tmp6696


} else {
tmp6697 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6689 = tmp6697


}

__e.TailApply(tmp6686, ifres6689)
return


}, 1)

tmp6699 := Call(__e, ns2_1set, symshen_4_5whitespace_6, tmp6685)


_ = tmp6699

tmp6700 := MakeNative(func(__e *ControlFlow) {
V2772 := __e.Get(1)
_ = V2772
tmp6708 := PrimEqual(MakeNumber(32), V2772)

if True == tmp6708 {
__e.Return(True)
return
} else {
tmp6706 := PrimEqual(MakeNumber(13), V2772)

if True == tmp6706 {
__e.Return(True)
return
} else {
tmp6704 := PrimEqual(MakeNumber(10), V2772)

if True == tmp6704 {
__e.Return(True)
return
} else {
tmp6702 := PrimEqual(MakeNumber(9), V2772)

if True == tmp6702 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}


}


}, 1)

tmp6709 := Call(__e, ns2_1set, symshen_4whitespace_2, tmp6700)


_ = tmp6709

tmp6710 := MakeNative(func(__e *ControlFlow) {
V2773 := __e.Get(1)
_ = V2773
tmp6733 := PrimEqual(Nil, V2773)

if True == tmp6733 {
__e.Return(Nil)
return
} else {
tmp6731 := PrimIsPair(V2773)

var ifres6727 Obj

if True == tmp6731 {
tmp6729 := PrimHead(V2773)

tmp6730 := Call(__e, PrimFunc(symshen_4packaged_2), tmp6729)


var ifres6728 Obj

if True == tmp6730 {
ifres6728 = True


} else {
ifres6728 = False


}

ifres6727 = ifres6728


} else {
ifres6727 = False


}

if True == ifres6727 {
tmp6711 := PrimHead(V2773)

tmp6712 := Call(__e, PrimFunc(symshen_4unpackage), tmp6711)


tmp6713 := PrimTail(V2773)

tmp6714 := Call(__e, PrimFunc(symappend), tmp6712, tmp6713)


__e.TailApply(PrimFunc(symshen_4unpackage_emacroexpand), tmp6714)
return


} else {
tmp6725 := PrimIsPair(V2773)

if True == tmp6725 {
tmp6715 := MakeNative(func(__e *ControlFlow) {
W2774 := __e.Get(1)
_ = W2774
tmp6721 := Call(__e, PrimFunc(symshen_4packaged_2), W2774)


if True == tmp6721 {
tmp6716 := PrimTail(V2773)

tmp6717 := PrimCons(W2774, tmp6716)

__e.TailApply(PrimFunc(symshen_4unpackage_emacroexpand), tmp6717)
return


} else {
tmp6718 := PrimTail(V2773)

tmp6719 := Call(__e, PrimFunc(symshen_4unpackage_emacroexpand), tmp6718)


__e.Return(PrimCons(W2774, tmp6719))
return


}


}, 1)

tmp6722 := PrimHead(V2773)

tmp6723 := Call(__e, PrimFunc(symmacroexpand), tmp6722)


__e.TailApply(tmp6715, tmp6723)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.unpackage&macroexpand")))
return
}


}


}


}, 1)

tmp6734 := Call(__e, ns2_1set, symshen_4unpackage_emacroexpand, tmp6710)


_ = tmp6734

tmp6735 := MakeNative(func(__e *ControlFlow) {
V2777 := __e.Get(1)
_ = V2777
tmp6750 := PrimIsPair(V2777)

var ifres6737 Obj

if True == tmp6750 {
tmp6748 := PrimHead(V2777)

tmp6749 := PrimEqual(sympackage, tmp6748)

var ifres6739 Obj

if True == tmp6749 {
tmp6746 := PrimTail(V2777)

tmp6747 := PrimIsPair(tmp6746)

var ifres6741 Obj

if True == tmp6747 {
tmp6743 := PrimTail(V2777)

tmp6744 := PrimTail(tmp6743)

tmp6745 := PrimIsPair(tmp6744)

var ifres6742 Obj

if True == tmp6745 {
ifres6742 = True


} else {
ifres6742 = False


}

ifres6741 = ifres6742


} else {
ifres6741 = False


}

var ifres6740 Obj

if True == ifres6741 {
ifres6740 = True


} else {
ifres6740 = False


}

ifres6739 = ifres6740


} else {
ifres6739 = False


}

var ifres6738 Obj

if True == ifres6739 {
ifres6738 = True


} else {
ifres6738 = False


}

ifres6737 = ifres6738


} else {
ifres6737 = False


}

if True == ifres6737 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp6751 := Call(__e, ns2_1set, symshen_4packaged_2, tmp6735)


_ = tmp6751

tmp6752 := MakeNative(func(__e *ControlFlow) {
V2780 := __e.Get(1)
_ = V2780
tmp6813 := PrimIsPair(V2780)

var ifres6795 Obj

if True == tmp6813 {
tmp6811 := PrimHead(V2780)

tmp6812 := PrimEqual(sympackage, tmp6811)

var ifres6797 Obj

if True == tmp6812 {
tmp6809 := PrimTail(V2780)

tmp6810 := PrimIsPair(tmp6809)

var ifres6799 Obj

if True == tmp6810 {
tmp6806 := PrimTail(V2780)

tmp6807 := PrimHead(tmp6806)

tmp6808 := PrimEqual(symnull, tmp6807)

var ifres6801 Obj

if True == tmp6808 {
tmp6803 := PrimTail(V2780)

tmp6804 := PrimTail(tmp6803)

tmp6805 := PrimIsPair(tmp6804)

var ifres6802 Obj

if True == tmp6805 {
ifres6802 = True


} else {
ifres6802 = False


}

ifres6801 = ifres6802


} else {
ifres6801 = False


}

var ifres6800 Obj

if True == ifres6801 {
ifres6800 = True


} else {
ifres6800 = False


}

ifres6799 = ifres6800


} else {
ifres6799 = False


}

var ifres6798 Obj

if True == ifres6799 {
ifres6798 = True


} else {
ifres6798 = False


}

ifres6797 = ifres6798


} else {
ifres6797 = False


}

var ifres6796 Obj

if True == ifres6797 {
ifres6796 = True


} else {
ifres6796 = False


}

ifres6795 = ifres6796


} else {
ifres6795 = False


}

if True == ifres6795 {
tmp6753 := PrimTail(V2780)

tmp6754 := PrimTail(tmp6753)

__e.Return(PrimTail(tmp6754))
return


} else {
tmp6793 := PrimIsPair(V2780)

var ifres6780 Obj

if True == tmp6793 {
tmp6791 := PrimHead(V2780)

tmp6792 := PrimEqual(sympackage, tmp6791)

var ifres6782 Obj

if True == tmp6792 {
tmp6789 := PrimTail(V2780)

tmp6790 := PrimIsPair(tmp6789)

var ifres6784 Obj

if True == tmp6790 {
tmp6786 := PrimTail(V2780)

tmp6787 := PrimTail(tmp6786)

tmp6788 := PrimIsPair(tmp6787)

var ifres6785 Obj

if True == tmp6788 {
ifres6785 = True


} else {
ifres6785 = False


}

ifres6784 = ifres6785


} else {
ifres6784 = False


}

var ifres6783 Obj

if True == ifres6784 {
ifres6783 = True


} else {
ifres6783 = False


}

ifres6782 = ifres6783


} else {
ifres6782 = False


}

var ifres6781 Obj

if True == ifres6782 {
ifres6781 = True


} else {
ifres6781 = False


}

ifres6780 = ifres6781


} else {
ifres6780 = False


}

if True == ifres6780 {
tmp6755 := MakeNative(func(__e *ControlFlow) {
W2781 := __e.Get(1)
_ = W2781
tmp6756 := MakeNative(func(__e *ControlFlow) {
W2782 := __e.Get(1)
_ = W2782
tmp6757 := MakeNative(func(__e *ControlFlow) {
W2783 := __e.Get(1)
_ = W2783
tmp6758 := MakeNative(func(__e *ControlFlow) {
W2784 := __e.Get(1)
_ = W2784
__e.Return(W2782)
return
}, 1)

tmp6759 := PrimTail(V2780)

tmp6760 := PrimHead(tmp6759)

tmp6761 := PrimTail(V2780)

tmp6762 := PrimTail(tmp6761)

tmp6763 := PrimTail(tmp6762)

tmp6764 := Call(__e, PrimFunc(symshen_4record_1internal), tmp6760, W2781, tmp6763)


__e.TailApply(tmp6758, tmp6764)
return


}, 1)

tmp6765 := PrimTail(V2780)

tmp6766 := PrimHead(tmp6765)

tmp6767 := Call(__e, PrimFunc(symshen_4record_1external), tmp6766, W2781)


__e.TailApply(tmp6757, tmp6767)
return


}, 1)

tmp6768 := PrimTail(V2780)

tmp6769 := PrimHead(tmp6768)

tmp6770 := PrimStr(tmp6769)

tmp6771 := PrimTail(V2780)

tmp6772 := PrimTail(tmp6771)

tmp6773 := PrimTail(tmp6772)

tmp6774 := Call(__e, PrimFunc(symshen_4package_1symbols), tmp6770, W2781, tmp6773)


__e.TailApply(tmp6756, tmp6774)
return


}, 1)

tmp6775 := PrimTail(V2780)

tmp6776 := PrimTail(tmp6775)

tmp6777 := PrimHead(tmp6776)

tmp6778 := Call(__e, PrimFunc(symeval), tmp6777)


__e.TailApply(tmp6755, tmp6778)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.unpackage")))
return
}


}


}, 1)

tmp6814 := Call(__e, ns2_1set, symshen_4unpackage, tmp6752)


_ = tmp6814

tmp6815 := MakeNative(func(__e *ControlFlow) {
V2785 := __e.Get(1)
_ = V2785
V2786 := __e.Get(2)
_ = V2786
V2787 := __e.Get(3)
_ = V2787
tmp6816 := MakeNative(func(__e *ControlFlow) {
W2788 := __e.Get(1)
_ = W2788
tmp6817 := MakeNative(func(__e *ControlFlow) {
W2790 := __e.Get(1)
_ = W2790
tmp6818 := Call(__e, PrimFunc(symunion), W2790, W2788)


tmp6819 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V2785, symshen_4internal_1symbols, tmp6818, tmp6819)
return


}, 1)

tmp6820 := PrimStr(V2785)

tmp6821 := Call(__e, PrimFunc(symshen_4internal_1symbols), tmp6820, V2786, V2787)


__e.TailApply(tmp6817, tmp6821)
return


}, 1)

tmp6822 := MakeNative(func(__e *ControlFlow) {
tmp6823 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V2785, symshen_4internal_1symbols, tmp6823)
return


}, 0)

tmp6824 := MakeNative(func(__e *ControlFlow) {
Z2789 := __e.Get(1)
_ = Z2789
__e.Return(Nil)
return
}, 1)

tmp6825 := Call(__e, try_1catch, tmp6822, tmp6824)


__e.TailApply(tmp6816, tmp6825)
return


}, 3)

tmp6826 := Call(__e, ns2_1set, symshen_4record_1internal, tmp6815)


_ = tmp6826

tmp6827 := MakeNative(func(__e *ControlFlow) {
V2797 := __e.Get(1)
_ = V2797
V2798 := __e.Get(2)
_ = V2798
V2799 := __e.Get(3)
_ = V2799
tmp6836 := PrimIsPair(V2799)

if True == tmp6836 {
tmp6828 := PrimHead(V2799)

tmp6829 := Call(__e, PrimFunc(symshen_4internal_1symbols), V2797, V2798, tmp6828)


tmp6830 := PrimTail(V2799)

tmp6831 := Call(__e, PrimFunc(symshen_4internal_1symbols), V2797, V2798, tmp6830)


__e.TailApply(PrimFunc(symunion), tmp6829, tmp6831)
return


} else {
tmp6834 := Call(__e, PrimFunc(symshen_4internal_2), V2799, V2797, V2798)


if True == tmp6834 {
tmp6832 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V2797, V2799)


__e.Return(PrimCons(tmp6832, Nil))
return


} else {
__e.Return(Nil)
return
}


}


}, 3)

tmp6837 := Call(__e, ns2_1set, symshen_4internal_1symbols, tmp6827)


_ = tmp6837

tmp6838 := MakeNative(func(__e *ControlFlow) {
V2800 := __e.Get(1)
_ = V2800
V2801 := __e.Get(2)
_ = V2801
tmp6839 := MakeNative(func(__e *ControlFlow) {
W2802 := __e.Get(1)
_ = W2802
tmp6840 := Call(__e, PrimFunc(symunion), V2801, W2802)


tmp6841 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V2800, symshen_4external_1symbols, tmp6840, tmp6841)
return


}, 1)

tmp6842 := MakeNative(func(__e *ControlFlow) {
tmp6843 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V2800, symshen_4external_1symbols, tmp6843)
return


}, 0)

tmp6844 := MakeNative(func(__e *ControlFlow) {
Z2803 := __e.Get(1)
_ = Z2803
__e.Return(Nil)
return
}, 1)

tmp6845 := Call(__e, try_1catch, tmp6842, tmp6844)


__e.TailApply(tmp6839, tmp6845)
return


}, 2)

tmp6846 := Call(__e, ns2_1set, symshen_4record_1external, tmp6838)


_ = tmp6846

tmp6847 := MakeNative(func(__e *ControlFlow) {
V2808 := __e.Get(1)
_ = V2808
V2809 := __e.Get(2)
_ = V2809
V2810 := __e.Get(3)
_ = V2810
tmp6852 := PrimIsPair(V2810)

if True == tmp6852 {
tmp6848 := MakeNative(func(__e *ControlFlow) {
Z2811 := __e.Get(1)
_ = Z2811
__e.TailApply(PrimFunc(symshen_4package_1symbols), V2808, V2809, Z2811)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp6848, V2810)
return


} else {
tmp6850 := Call(__e, PrimFunc(symshen_4internal_2), V2810, V2808, V2809)


if True == tmp6850 {
__e.TailApply(PrimFunc(symshen_4intern_1in_1package), V2808, V2810)
return
} else {
__e.Return(V2810)
return
}


}


}, 3)

tmp6853 := Call(__e, ns2_1set, symshen_4package_1symbols, tmp6847)


_ = tmp6853

tmp6854 := MakeNative(func(__e *ControlFlow) {
V2812 := __e.Get(1)
_ = V2812
V2813 := __e.Get(2)
_ = V2813
tmp6855 := PrimStr(V2813)

tmp6856 := Call(__e, PrimFunc(sym_8s), MakeString("."), tmp6855)


tmp6857 := Call(__e, PrimFunc(sym_8s), V2812, tmp6856)


__e.Return(PrimIntern(tmp6857))
return


}, 2)

tmp6858 := Call(__e, ns2_1set, symshen_4intern_1in_1package, tmp6854)


_ = tmp6858

tmp6859 := MakeNative(func(__e *ControlFlow) {
V2814 := __e.Get(1)
_ = V2814
V2815 := __e.Get(2)
_ = V2815
V2816 := __e.Get(3)
_ = V2816
tmp6889 := Call(__e, PrimFunc(symelement_2), V2814, V2816)


tmp6890 := PrimNot(tmp6889)

if True == tmp6890 {
tmp6886 := Call(__e, PrimFunc(symshen_4sng_2), V2814)


tmp6887 := PrimNot(tmp6886)

var ifres6861 Obj

if True == tmp6887 {
tmp6884 := Call(__e, PrimFunc(symshen_4dbl_2), V2814)


tmp6885 := PrimNot(tmp6884)

var ifres6863 Obj

if True == tmp6885 {
tmp6883 := PrimIsSymbol(V2814)

var ifres6865 Obj

if True == tmp6883 {
tmp6881 := Call(__e, PrimFunc(symshen_4sysfunc_2), V2814)


tmp6882 := PrimNot(tmp6881)

var ifres6867 Obj

if True == tmp6882 {
tmp6879 := PrimIsVariable(V2814)

tmp6880 := PrimNot(tmp6879)

var ifres6869 Obj

if True == tmp6880 {
tmp6876 := PrimStr(V2814)

tmp6877 := Call(__e, PrimFunc(symshen_4internal_1to_1shen_2), tmp6876)


tmp6878 := PrimNot(tmp6877)

var ifres6871 Obj

if True == tmp6878 {
tmp6873 := PrimStr(V2814)

tmp6874 := Call(__e, PrimFunc(symshen_4internal_1to_1P_2), V2815, tmp6873)


tmp6875 := PrimNot(tmp6874)

var ifres6872 Obj

if True == tmp6875 {
ifres6872 = True


} else {
ifres6872 = False


}

ifres6871 = ifres6872


} else {
ifres6871 = False


}

var ifres6870 Obj

if True == ifres6871 {
ifres6870 = True


} else {
ifres6870 = False


}

ifres6869 = ifres6870


} else {
ifres6869 = False


}

var ifres6868 Obj

if True == ifres6869 {
ifres6868 = True


} else {
ifres6868 = False


}

ifres6867 = ifres6868


} else {
ifres6867 = False


}

var ifres6866 Obj

if True == ifres6867 {
ifres6866 = True


} else {
ifres6866 = False


}

ifres6865 = ifres6866


} else {
ifres6865 = False


}

var ifres6864 Obj

if True == ifres6865 {
ifres6864 = True


} else {
ifres6864 = False


}

ifres6863 = ifres6864


} else {
ifres6863 = False


}

var ifres6862 Obj

if True == ifres6863 {
ifres6862 = True


} else {
ifres6862 = False


}

ifres6861 = ifres6862


} else {
ifres6861 = False


}

if True == ifres6861 {
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


}, 3)

tmp6891 := Call(__e, ns2_1set, symshen_4internal_2, tmp6859)


_ = tmp6891

tmp6892 := MakeNative(func(__e *ControlFlow) {
V2821 := __e.Get(1)
_ = V2821
tmp6946 := Call(__e, PrimFunc(symshen_4_7string_2), V2821)


var ifres6894 Obj

if True == tmp6946 {
tmp6944 := Call(__e, PrimFunc(symhdstr), V2821)


tmp6945 := PrimEqual(MakeString("s"), tmp6944)

var ifres6896 Obj

if True == tmp6945 {
tmp6942 := PrimTailString(V2821)

tmp6943 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6942)


var ifres6898 Obj

if True == tmp6943 {
tmp6939 := PrimTailString(V2821)

tmp6940 := Call(__e, PrimFunc(symhdstr), tmp6939)


tmp6941 := PrimEqual(MakeString("h"), tmp6940)

var ifres6900 Obj

if True == tmp6941 {
tmp6936 := PrimTailString(V2821)

tmp6937 := PrimTailString(tmp6936)

tmp6938 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6937)


var ifres6902 Obj

if True == tmp6938 {
tmp6932 := PrimTailString(V2821)

tmp6933 := PrimTailString(tmp6932)

tmp6934 := Call(__e, PrimFunc(symhdstr), tmp6933)


tmp6935 := PrimEqual(MakeString("e"), tmp6934)

var ifres6904 Obj

if True == tmp6935 {
tmp6928 := PrimTailString(V2821)

tmp6929 := PrimTailString(tmp6928)

tmp6930 := PrimTailString(tmp6929)

tmp6931 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6930)


var ifres6906 Obj

if True == tmp6931 {
tmp6923 := PrimTailString(V2821)

tmp6924 := PrimTailString(tmp6923)

tmp6925 := PrimTailString(tmp6924)

tmp6926 := Call(__e, PrimFunc(symhdstr), tmp6925)


tmp6927 := PrimEqual(MakeString("n"), tmp6926)

var ifres6908 Obj

if True == tmp6927 {
tmp6918 := PrimTailString(V2821)

tmp6919 := PrimTailString(tmp6918)

tmp6920 := PrimTailString(tmp6919)

tmp6921 := PrimTailString(tmp6920)

tmp6922 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6921)


var ifres6910 Obj

if True == tmp6922 {
tmp6912 := PrimTailString(V2821)

tmp6913 := PrimTailString(tmp6912)

tmp6914 := PrimTailString(tmp6913)

tmp6915 := PrimTailString(tmp6914)

tmp6916 := Call(__e, PrimFunc(symhdstr), tmp6915)


tmp6917 := PrimEqual(MakeString("."), tmp6916)

var ifres6911 Obj

if True == tmp6917 {
ifres6911 = True


} else {
ifres6911 = False


}

ifres6910 = ifres6911


} else {
ifres6910 = False


}

var ifres6909 Obj

if True == ifres6910 {
ifres6909 = True


} else {
ifres6909 = False


}

ifres6908 = ifres6909


} else {
ifres6908 = False


}

var ifres6907 Obj

if True == ifres6908 {
ifres6907 = True


} else {
ifres6907 = False


}

ifres6906 = ifres6907


} else {
ifres6906 = False


}

var ifres6905 Obj

if True == ifres6906 {
ifres6905 = True


} else {
ifres6905 = False


}

ifres6904 = ifres6905


} else {
ifres6904 = False


}

var ifres6903 Obj

if True == ifres6904 {
ifres6903 = True


} else {
ifres6903 = False


}

ifres6902 = ifres6903


} else {
ifres6902 = False


}

var ifres6901 Obj

if True == ifres6902 {
ifres6901 = True


} else {
ifres6901 = False


}

ifres6900 = ifres6901


} else {
ifres6900 = False


}

var ifres6899 Obj

if True == ifres6900 {
ifres6899 = True


} else {
ifres6899 = False


}

ifres6898 = ifres6899


} else {
ifres6898 = False


}

var ifres6897 Obj

if True == ifres6898 {
ifres6897 = True


} else {
ifres6897 = False


}

ifres6896 = ifres6897


} else {
ifres6896 = False


}

var ifres6895 Obj

if True == ifres6896 {
ifres6895 = True


} else {
ifres6895 = False


}

ifres6894 = ifres6895


} else {
ifres6894 = False


}

if True == ifres6894 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp6947 := Call(__e, ns2_1set, symshen_4internal_1to_1shen_2, tmp6892)


_ = tmp6947

tmp6948 := MakeNative(func(__e *ControlFlow) {
V2822 := __e.Get(1)
_ = V2822
tmp6949 := PrimValue(sym_dproperty_1vector_d)

tmp6950 := Call(__e, PrimFunc(symget), symshen, symshen_4external_1symbols, tmp6949)


__e.TailApply(PrimFunc(symelement_2), V2822, tmp6950)
return


}, 1)

tmp6951 := Call(__e, ns2_1set, symshen_4sysfunc_2, tmp6948)


_ = tmp6951

tmp6952 := MakeNative(func(__e *ControlFlow) {
V2830 := __e.Get(1)
_ = V2830
V2831 := __e.Get(2)
_ = V2831
tmp6973 := PrimEqual(MakeString(""), V2830)

var ifres6966 Obj

if True == tmp6973 {
tmp6972 := Call(__e, PrimFunc(symshen_4_7string_2), V2831)


var ifres6968 Obj

if True == tmp6972 {
tmp6970 := Call(__e, PrimFunc(symhdstr), V2831)


tmp6971 := PrimEqual(MakeString("."), tmp6970)

var ifres6969 Obj

if True == tmp6971 {
ifres6969 = True


} else {
ifres6969 = False


}

ifres6968 = ifres6969


} else {
ifres6968 = False


}

var ifres6967 Obj

if True == ifres6968 {
ifres6967 = True


} else {
ifres6967 = False


}

ifres6966 = ifres6967


} else {
ifres6966 = False


}

if True == ifres6966 {
__e.Return(True)
return
} else {
tmp6964 := Call(__e, PrimFunc(symshen_4_7string_2), V2830)


var ifres6956 Obj

if True == tmp6964 {
tmp6963 := Call(__e, PrimFunc(symshen_4_7string_2), V2831)


var ifres6958 Obj

if True == tmp6963 {
tmp6960 := Call(__e, PrimFunc(symhdstr), V2830)


tmp6961 := Call(__e, PrimFunc(symhdstr), V2831)


tmp6962 := PrimEqual(tmp6960, tmp6961)

var ifres6959 Obj

if True == tmp6962 {
ifres6959 = True


} else {
ifres6959 = False


}

ifres6958 = ifres6959


} else {
ifres6958 = False


}

var ifres6957 Obj

if True == ifres6958 {
ifres6957 = True


} else {
ifres6957 = False


}

ifres6956 = ifres6957


} else {
ifres6956 = False


}

if True == ifres6956 {
tmp6953 := PrimTailString(V2830)

tmp6954 := PrimTailString(V2831)

__e.TailApply(PrimFunc(symshen_4internal_1to_1P_2), tmp6953, tmp6954)
return


} else {
__e.Return(False)
return
}


}


}, 2)

tmp6974 := Call(__e, ns2_1set, symshen_4internal_1to_1P_2, tmp6952)


_ = tmp6974

tmp6975 := MakeNative(func(__e *ControlFlow) {
V2834 := __e.Get(1)
_ = V2834
V2835 := __e.Get(2)
_ = V2835
tmp6988 := Call(__e, PrimFunc(symelement_2), V2834, V2835)


if True == tmp6988 {
__e.Return(V2834)
return
} else {
tmp6986 := PrimIsPair(V2834)

var ifres6982 Obj

if True == tmp6986 {
tmp6984 := PrimHead(V2834)

tmp6985 := Call(__e, PrimFunc(symshen_4non_1application_2), tmp6984)


var ifres6983 Obj

if True == tmp6985 {
ifres6983 = True


} else {
ifres6983 = False


}

ifres6982 = ifres6983


} else {
ifres6982 = False


}

if True == ifres6982 {
tmp6976 := PrimHead(V2834)

__e.TailApply(PrimFunc(symshen_4special_1case), tmp6976, V2834, V2835)
return


} else {
tmp6980 := PrimIsPair(V2834)

if True == tmp6980 {
tmp6977 := MakeNative(func(__e *ControlFlow) {
Z2836 := __e.Get(1)
_ = Z2836
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2836, V2835)
return
}, 1)

tmp6978 := Call(__e, PrimFunc(symmap), tmp6977, V2834)


__e.TailApply(PrimFunc(symshen_4process_1application), tmp6978, V2835)
return


} else {
__e.Return(V2834)
return
}


}


}


}, 2)

tmp6989 := Call(__e, ns2_1set, symshen_4process_1applications, tmp6975)


_ = tmp6989

tmp6990 := MakeNative(func(__e *ControlFlow) {
V2839 := __e.Get(1)
_ = V2839
tmp7000 := PrimEqual(symdefine, V2839)

if True == tmp7000 {
__e.Return(True)
return
} else {
tmp6998 := PrimEqual(symdefun, V2839)

if True == tmp6998 {
__e.Return(True)
return
} else {
tmp6996 := PrimEqual(symsynonyms, V2839)

if True == tmp6996 {
__e.Return(True)
return
} else {
tmp6994 := Call(__e, PrimFunc(symshen_4special_2), V2839)


if True == tmp6994 {
__e.Return(True)
return
} else {
tmp6992 := Call(__e, PrimFunc(symshen_4extraspecial_2), V2839)


if True == tmp6992 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}


}


}


}, 1)

tmp7001 := Call(__e, ns2_1set, symshen_4non_1application_2, tmp6990)


_ = tmp7001

tmp7002 := MakeNative(func(__e *ControlFlow) {
V2844 := __e.Get(1)
_ = V2844
V2845 := __e.Get(2)
_ = V2845
V2846 := __e.Get(3)
_ = V2846
tmp7244 := PrimEqual(symlambda, V2844)

var ifres7222 Obj

if True == tmp7244 {
tmp7243 := PrimIsPair(V2845)

var ifres7224 Obj

if True == tmp7243 {
tmp7241 := PrimHead(V2845)

tmp7242 := PrimEqual(symlambda, tmp7241)

var ifres7226 Obj

if True == tmp7242 {
tmp7239 := PrimTail(V2845)

tmp7240 := PrimIsPair(tmp7239)

var ifres7228 Obj

if True == tmp7240 {
tmp7236 := PrimTail(V2845)

tmp7237 := PrimTail(tmp7236)

tmp7238 := PrimIsPair(tmp7237)

var ifres7230 Obj

if True == tmp7238 {
tmp7232 := PrimTail(V2845)

tmp7233 := PrimTail(tmp7232)

tmp7234 := PrimTail(tmp7233)

tmp7235 := PrimEqual(Nil, tmp7234)

var ifres7231 Obj

if True == tmp7235 {
ifres7231 = True


} else {
ifres7231 = False


}

ifres7230 = ifres7231


} else {
ifres7230 = False


}

var ifres7229 Obj

if True == ifres7230 {
ifres7229 = True


} else {
ifres7229 = False


}

ifres7228 = ifres7229


} else {
ifres7228 = False


}

var ifres7227 Obj

if True == ifres7228 {
ifres7227 = True


} else {
ifres7227 = False


}

ifres7226 = ifres7227


} else {
ifres7226 = False


}

var ifres7225 Obj

if True == ifres7226 {
ifres7225 = True


} else {
ifres7225 = False


}

ifres7224 = ifres7225


} else {
ifres7224 = False


}

var ifres7223 Obj

if True == ifres7224 {
ifres7223 = True


} else {
ifres7223 = False


}

ifres7222 = ifres7223


} else {
ifres7222 = False


}

if True == ifres7222 {
tmp7003 := PrimTail(V2845)

tmp7004 := PrimHead(tmp7003)

tmp7005 := PrimTail(V2845)

tmp7006 := PrimTail(tmp7005)

tmp7007 := PrimHead(tmp7006)

tmp7008 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7007, V2846)


tmp7009 := PrimCons(tmp7008, Nil)

tmp7010 := PrimCons(tmp7004, tmp7009)

__e.Return(PrimCons(symlambda, tmp7010))
return


} else {
tmp7220 := PrimEqual(symlet, V2844)

var ifres7191 Obj

if True == tmp7220 {
tmp7219 := PrimIsPair(V2845)

var ifres7193 Obj

if True == tmp7219 {
tmp7217 := PrimHead(V2845)

tmp7218 := PrimEqual(symlet, tmp7217)

var ifres7195 Obj

if True == tmp7218 {
tmp7215 := PrimTail(V2845)

tmp7216 := PrimIsPair(tmp7215)

var ifres7197 Obj

if True == tmp7216 {
tmp7212 := PrimTail(V2845)

tmp7213 := PrimTail(tmp7212)

tmp7214 := PrimIsPair(tmp7213)

var ifres7199 Obj

if True == tmp7214 {
tmp7208 := PrimTail(V2845)

tmp7209 := PrimTail(tmp7208)

tmp7210 := PrimTail(tmp7209)

tmp7211 := PrimIsPair(tmp7210)

var ifres7201 Obj

if True == tmp7211 {
tmp7203 := PrimTail(V2845)

tmp7204 := PrimTail(tmp7203)

tmp7205 := PrimTail(tmp7204)

tmp7206 := PrimTail(tmp7205)

tmp7207 := PrimEqual(Nil, tmp7206)

var ifres7202 Obj

if True == tmp7207 {
ifres7202 = True


} else {
ifres7202 = False


}

ifres7201 = ifres7202


} else {
ifres7201 = False


}

var ifres7200 Obj

if True == ifres7201 {
ifres7200 = True


} else {
ifres7200 = False


}

ifres7199 = ifres7200


} else {
ifres7199 = False


}

var ifres7198 Obj

if True == ifres7199 {
ifres7198 = True


} else {
ifres7198 = False


}

ifres7197 = ifres7198


} else {
ifres7197 = False


}

var ifres7196 Obj

if True == ifres7197 {
ifres7196 = True


} else {
ifres7196 = False


}

ifres7195 = ifres7196


} else {
ifres7195 = False


}

var ifres7194 Obj

if True == ifres7195 {
ifres7194 = True


} else {
ifres7194 = False


}

ifres7193 = ifres7194


} else {
ifres7193 = False


}

var ifres7192 Obj

if True == ifres7193 {
ifres7192 = True


} else {
ifres7192 = False


}

ifres7191 = ifres7192


} else {
ifres7191 = False


}

if True == ifres7191 {
tmp7011 := PrimTail(V2845)

tmp7012 := PrimHead(tmp7011)

tmp7013 := PrimTail(V2845)

tmp7014 := PrimTail(tmp7013)

tmp7015 := PrimHead(tmp7014)

tmp7016 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7015, V2846)


tmp7017 := PrimTail(V2845)

tmp7018 := PrimTail(tmp7017)

tmp7019 := PrimTail(tmp7018)

tmp7020 := PrimHead(tmp7019)

tmp7021 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7020, V2846)


tmp7022 := PrimCons(tmp7021, Nil)

tmp7023 := PrimCons(tmp7016, tmp7022)

tmp7024 := PrimCons(tmp7012, tmp7023)

__e.Return(PrimCons(symlet, tmp7024))
return


} else {
tmp7189 := PrimEqual(symdefun, V2844)

var ifres7160 Obj

if True == tmp7189 {
tmp7188 := PrimIsPair(V2845)

var ifres7162 Obj

if True == tmp7188 {
tmp7186 := PrimHead(V2845)

tmp7187 := PrimEqual(symdefun, tmp7186)

var ifres7164 Obj

if True == tmp7187 {
tmp7184 := PrimTail(V2845)

tmp7185 := PrimIsPair(tmp7184)

var ifres7166 Obj

if True == tmp7185 {
tmp7181 := PrimTail(V2845)

tmp7182 := PrimTail(tmp7181)

tmp7183 := PrimIsPair(tmp7182)

var ifres7168 Obj

if True == tmp7183 {
tmp7177 := PrimTail(V2845)

tmp7178 := PrimTail(tmp7177)

tmp7179 := PrimTail(tmp7178)

tmp7180 := PrimIsPair(tmp7179)

var ifres7170 Obj

if True == tmp7180 {
tmp7172 := PrimTail(V2845)

tmp7173 := PrimTail(tmp7172)

tmp7174 := PrimTail(tmp7173)

tmp7175 := PrimTail(tmp7174)

tmp7176 := PrimEqual(Nil, tmp7175)

var ifres7171 Obj

if True == tmp7176 {
ifres7171 = True


} else {
ifres7171 = False


}

ifres7170 = ifres7171


} else {
ifres7170 = False


}

var ifres7169 Obj

if True == ifres7170 {
ifres7169 = True


} else {
ifres7169 = False


}

ifres7168 = ifres7169


} else {
ifres7168 = False


}

var ifres7167 Obj

if True == ifres7168 {
ifres7167 = True


} else {
ifres7167 = False


}

ifres7166 = ifres7167


} else {
ifres7166 = False


}

var ifres7165 Obj

if True == ifres7166 {
ifres7165 = True


} else {
ifres7165 = False


}

ifres7164 = ifres7165


} else {
ifres7164 = False


}

var ifres7163 Obj

if True == ifres7164 {
ifres7163 = True


} else {
ifres7163 = False


}

ifres7162 = ifres7163


} else {
ifres7162 = False


}

var ifres7161 Obj

if True == ifres7162 {
ifres7161 = True


} else {
ifres7161 = False


}

ifres7160 = ifres7161


} else {
ifres7160 = False


}

if True == ifres7160 {
__e.Return(V2845)
return
} else {
tmp7158 := PrimEqual(symdefine, V2844)

var ifres7136 Obj

if True == tmp7158 {
tmp7157 := PrimIsPair(V2845)

var ifres7138 Obj

if True == tmp7157 {
tmp7155 := PrimHead(V2845)

tmp7156 := PrimEqual(symdefine, tmp7155)

var ifres7140 Obj

if True == tmp7156 {
tmp7153 := PrimTail(V2845)

tmp7154 := PrimIsPair(tmp7153)

var ifres7142 Obj

if True == tmp7154 {
tmp7150 := PrimTail(V2845)

tmp7151 := PrimTail(tmp7150)

tmp7152 := PrimIsPair(tmp7151)

var ifres7144 Obj

if True == tmp7152 {
tmp7146 := PrimTail(V2845)

tmp7147 := PrimTail(tmp7146)

tmp7148 := PrimHead(tmp7147)

tmp7149 := PrimEqual(sym_i, tmp7148)

var ifres7145 Obj

if True == tmp7149 {
ifres7145 = True


} else {
ifres7145 = False


}

ifres7144 = ifres7145


} else {
ifres7144 = False


}

var ifres7143 Obj

if True == ifres7144 {
ifres7143 = True


} else {
ifres7143 = False


}

ifres7142 = ifres7143


} else {
ifres7142 = False


}

var ifres7141 Obj

if True == ifres7142 {
ifres7141 = True


} else {
ifres7141 = False


}

ifres7140 = ifres7141


} else {
ifres7140 = False


}

var ifres7139 Obj

if True == ifres7140 {
ifres7139 = True


} else {
ifres7139 = False


}

ifres7138 = ifres7139


} else {
ifres7138 = False


}

var ifres7137 Obj

if True == ifres7138 {
ifres7137 = True


} else {
ifres7137 = False


}

ifres7136 = ifres7137


} else {
ifres7136 = False


}

if True == ifres7136 {
tmp7025 := PrimTail(V2845)

tmp7026 := PrimHead(tmp7025)

tmp7027 := PrimTail(V2845)

tmp7028 := PrimHead(tmp7027)

tmp7029 := PrimTail(V2845)

tmp7030 := PrimTail(tmp7029)

tmp7031 := PrimTail(tmp7030)

tmp7032 := Call(__e, PrimFunc(symshen_4process_1after_1type), tmp7028, tmp7031, V2846)


tmp7033 := PrimCons(sym_i, tmp7032)

tmp7034 := PrimCons(tmp7026, tmp7033)

__e.Return(PrimCons(symdefine, tmp7034))
return


} else {
tmp7134 := PrimEqual(symdefine, V2844)

var ifres7123 Obj

if True == tmp7134 {
tmp7133 := PrimIsPair(V2845)

var ifres7125 Obj

if True == tmp7133 {
tmp7131 := PrimHead(V2845)

tmp7132 := PrimEqual(symdefine, tmp7131)

var ifres7127 Obj

if True == tmp7132 {
tmp7129 := PrimTail(V2845)

tmp7130 := PrimIsPair(tmp7129)

var ifres7128 Obj

if True == tmp7130 {
ifres7128 = True


} else {
ifres7128 = False


}

ifres7127 = ifres7128


} else {
ifres7127 = False


}

var ifres7126 Obj

if True == ifres7127 {
ifres7126 = True


} else {
ifres7126 = False


}

ifres7125 = ifres7126


} else {
ifres7125 = False


}

var ifres7124 Obj

if True == ifres7125 {
ifres7124 = True


} else {
ifres7124 = False


}

ifres7123 = ifres7124


} else {
ifres7123 = False


}

if True == ifres7123 {
tmp7035 := PrimTail(V2845)

tmp7036 := PrimHead(tmp7035)

tmp7037 := MakeNative(func(__e *ControlFlow) {
Z2847 := __e.Get(1)
_ = Z2847
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2847, V2846)
return
}, 1)

tmp7038 := PrimTail(V2845)

tmp7039 := PrimTail(tmp7038)

tmp7040 := Call(__e, PrimFunc(symmap), tmp7037, tmp7039)


tmp7041 := PrimCons(tmp7036, tmp7040)

__e.Return(PrimCons(symdefine, tmp7041))
return


} else {
tmp7121 := PrimEqual(symsynonyms, V2844)

if True == tmp7121 {
__e.Return(PrimCons(symsynonyms, V2845))
return
} else {
tmp7119 := PrimEqual(symtype, V2844)

var ifres7097 Obj

if True == tmp7119 {
tmp7118 := PrimIsPair(V2845)

var ifres7099 Obj

if True == tmp7118 {
tmp7116 := PrimHead(V2845)

tmp7117 := PrimEqual(symtype, tmp7116)

var ifres7101 Obj

if True == tmp7117 {
tmp7114 := PrimTail(V2845)

tmp7115 := PrimIsPair(tmp7114)

var ifres7103 Obj

if True == tmp7115 {
tmp7111 := PrimTail(V2845)

tmp7112 := PrimTail(tmp7111)

tmp7113 := PrimIsPair(tmp7112)

var ifres7105 Obj

if True == tmp7113 {
tmp7107 := PrimTail(V2845)

tmp7108 := PrimTail(tmp7107)

tmp7109 := PrimTail(tmp7108)

tmp7110 := PrimEqual(Nil, tmp7109)

var ifres7106 Obj

if True == tmp7110 {
ifres7106 = True


} else {
ifres7106 = False


}

ifres7105 = ifres7106


} else {
ifres7105 = False


}

var ifres7104 Obj

if True == ifres7105 {
ifres7104 = True


} else {
ifres7104 = False


}

ifres7103 = ifres7104


} else {
ifres7103 = False


}

var ifres7102 Obj

if True == ifres7103 {
ifres7102 = True


} else {
ifres7102 = False


}

ifres7101 = ifres7102


} else {
ifres7101 = False


}

var ifres7100 Obj

if True == ifres7101 {
ifres7100 = True


} else {
ifres7100 = False


}

ifres7099 = ifres7100


} else {
ifres7099 = False


}

var ifres7098 Obj

if True == ifres7099 {
ifres7098 = True


} else {
ifres7098 = False


}

ifres7097 = ifres7098


} else {
ifres7097 = False


}

if True == ifres7097 {
tmp7042 := PrimTail(V2845)

tmp7043 := PrimHead(tmp7042)

tmp7044 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7043, V2846)


tmp7045 := PrimTail(V2845)

tmp7046 := PrimTail(tmp7045)

tmp7047 := PrimCons(tmp7044, tmp7046)

__e.Return(PrimCons(symtype, tmp7047))
return


} else {
tmp7095 := PrimEqual(syminput_7, V2844)

var ifres7073 Obj

if True == tmp7095 {
tmp7094 := PrimIsPair(V2845)

var ifres7075 Obj

if True == tmp7094 {
tmp7092 := PrimHead(V2845)

tmp7093 := PrimEqual(syminput_7, tmp7092)

var ifres7077 Obj

if True == tmp7093 {
tmp7090 := PrimTail(V2845)

tmp7091 := PrimIsPair(tmp7090)

var ifres7079 Obj

if True == tmp7091 {
tmp7087 := PrimTail(V2845)

tmp7088 := PrimTail(tmp7087)

tmp7089 := PrimIsPair(tmp7088)

var ifres7081 Obj

if True == tmp7089 {
tmp7083 := PrimTail(V2845)

tmp7084 := PrimTail(tmp7083)

tmp7085 := PrimTail(tmp7084)

tmp7086 := PrimEqual(Nil, tmp7085)

var ifres7082 Obj

if True == tmp7086 {
ifres7082 = True


} else {
ifres7082 = False


}

ifres7081 = ifres7082


} else {
ifres7081 = False


}

var ifres7080 Obj

if True == ifres7081 {
ifres7080 = True


} else {
ifres7080 = False


}

ifres7079 = ifres7080


} else {
ifres7079 = False


}

var ifres7078 Obj

if True == ifres7079 {
ifres7078 = True


} else {
ifres7078 = False


}

ifres7077 = ifres7078


} else {
ifres7077 = False


}

var ifres7076 Obj

if True == ifres7077 {
ifres7076 = True


} else {
ifres7076 = False


}

ifres7075 = ifres7076


} else {
ifres7075 = False


}

var ifres7074 Obj

if True == ifres7075 {
ifres7074 = True


} else {
ifres7074 = False


}

ifres7073 = ifres7074


} else {
ifres7073 = False


}

if True == ifres7073 {
tmp7048 := PrimTail(V2845)

tmp7049 := PrimHead(tmp7048)

tmp7050 := PrimTail(V2845)

tmp7051 := PrimTail(tmp7050)

tmp7052 := PrimHead(tmp7051)

tmp7053 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7052, V2846)


tmp7054 := PrimCons(tmp7053, Nil)

tmp7055 := PrimCons(tmp7049, tmp7054)

__e.Return(PrimCons(syminput_7, tmp7055))
return


} else {
tmp7071 := PrimIsPair(V2845)

var ifres7067 Obj

if True == tmp7071 {
tmp7069 := PrimHead(V2845)

tmp7070 := Call(__e, PrimFunc(symshen_4special_2), tmp7069)


var ifres7068 Obj

if True == tmp7070 {
ifres7068 = True


} else {
ifres7068 = False


}

ifres7067 = ifres7068


} else {
ifres7067 = False


}

if True == ifres7067 {
tmp7056 := PrimHead(V2845)

tmp7057 := MakeNative(func(__e *ControlFlow) {
Z2848 := __e.Get(1)
_ = Z2848
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2848, V2846)
return
}, 1)

tmp7058 := PrimTail(V2845)

tmp7059 := Call(__e, PrimFunc(symmap), tmp7057, tmp7058)


__e.Return(PrimCons(tmp7056, tmp7059))
return


} else {
tmp7065 := PrimIsPair(V2845)

var ifres7061 Obj

if True == tmp7065 {
tmp7063 := PrimHead(V2845)

tmp7064 := Call(__e, PrimFunc(symshen_4extraspecial_2), tmp7063)


var ifres7062 Obj

if True == tmp7064 {
ifres7062 = True


} else {
ifres7062 = False


}

ifres7061 = ifres7062


} else {
ifres7061 = False


}

if True == ifres7061 {
__e.Return(V2845)
return
} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.special-case")))
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


}, 3)

tmp7245 := Call(__e, ns2_1set, symshen_4special_1case, tmp7002)


_ = tmp7245

tmp7246 := MakeNative(func(__e *ControlFlow) {
V2851 := __e.Get(1)
_ = V2851
V2852 := __e.Get(2)
_ = V2852
V2853 := __e.Get(3)
_ = V2853
tmp7262 := PrimIsPair(V2852)

var ifres7258 Obj

if True == tmp7262 {
tmp7260 := PrimHead(V2852)

tmp7261 := PrimEqual(sym_j, tmp7260)

var ifres7259 Obj

if True == tmp7261 {
ifres7259 = True


} else {
ifres7259 = False


}

ifres7258 = ifres7259


} else {
ifres7258 = False


}

if True == ifres7258 {
tmp7247 := MakeNative(func(__e *ControlFlow) {
Z2854 := __e.Get(1)
_ = Z2854
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2854, V2853)
return
}, 1)

tmp7248 := PrimTail(V2852)

tmp7249 := Call(__e, PrimFunc(symmap), tmp7247, tmp7248)


__e.Return(PrimCons(sym_j, tmp7249))
return


} else {
tmp7256 := PrimIsPair(V2852)

if True == tmp7256 {
tmp7250 := PrimHead(V2852)

tmp7251 := PrimTail(V2852)

tmp7252 := Call(__e, PrimFunc(symshen_4process_1after_1type), V2851, tmp7251, V2853)


__e.Return(PrimCons(tmp7250, tmp7252))
return


} else {
tmp7253 := Call(__e, PrimFunc(symshen_4app), V2851, MakeString("\n"), symshen_4a)


tmp7254 := PrimStringConcat(MakeString("missing } in "), tmp7253)

__e.Return(PrimSimpleError(tmp7254))
return


}


}


}, 3)

tmp7263 := Call(__e, ns2_1set, symshen_4process_1after_1type, tmp7246)


_ = tmp7263

tmp7264 := MakeNative(func(__e *ControlFlow) {
V2855 := __e.Get(1)
_ = V2855
V2856 := __e.Get(2)
_ = V2856
tmp7309 := PrimIsPair(V2855)

if True == tmp7309 {
tmp7265 := MakeNative(func(__e *ControlFlow) {
W2857 := __e.Get(1)
_ = W2857
tmp7266 := MakeNative(func(__e *ControlFlow) {
W2858 := __e.Get(1)
_ = W2858
tmp7303 := Call(__e, PrimFunc(symelement_2), V2855, V2856)


if True == tmp7303 {
__e.Return(V2855)
return
} else {
tmp7300 := PrimHead(V2855)

tmp7301 := Call(__e, PrimFunc(symshen_4shen_1call_2), tmp7300)


if True == tmp7301 {
__e.Return(V2855)
return
} else {
tmp7298 := Call(__e, PrimFunc(symshen_4foreign_2), V2855)


if True == tmp7298 {
__e.TailApply(PrimFunc(symshen_4unpack_1foreign), V2855)
return
} else {
tmp7296 := Call(__e, PrimFunc(symshen_4fn_1call_2), V2855)


if True == tmp7296 {
__e.TailApply(PrimFunc(symshen_4fn_1call), V2855)
return
} else {
tmp7294 := Call(__e, PrimFunc(symshen_4zero_1place_2), V2855)


if True == tmp7294 {
__e.Return(V2855)
return
} else {
tmp7291 := PrimHead(V2855)

tmp7292 := Call(__e, PrimFunc(symshen_4undefined_1f_2), tmp7291, W2857)


if True == tmp7292 {
tmp7267 := PrimHead(V2855)

tmp7268 := PrimCons(tmp7267, Nil)

tmp7269 := PrimCons(symfn, tmp7268)

tmp7270 := PrimTail(V2855)

tmp7271 := PrimCons(tmp7269, tmp7270)

__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp7271)
return


} else {
tmp7288 := PrimHead(V2855)

tmp7289 := PrimIsVariable(tmp7288)

if True == tmp7289 {
__e.TailApply(PrimFunc(symshen_4simple_1curry), V2855)
return
} else {
tmp7285 := PrimHead(V2855)

tmp7286 := Call(__e, PrimFunc(symshen_4application_2), tmp7285)


if True == tmp7286 {
__e.TailApply(PrimFunc(symshen_4simple_1curry), V2855)
return
} else {
tmp7282 := PrimHead(V2855)

tmp7283 := Call(__e, PrimFunc(symshen_4partial_1application_d_2), tmp7282, W2857, W2858)


if True == tmp7283 {
tmp7272 := PrimNumberSubtract(W2857, W2858)

__e.TailApply(PrimFunc(symshen_4lambda_1function), V2855, tmp7272)
return


} else {
tmp7279 := PrimHead(V2855)

tmp7280 := Call(__e, PrimFunc(symshen_4overapplication_2), tmp7279, W2857, W2858)


if True == tmp7280 {
tmp7273 := PrimHead(V2855)

tmp7274 := PrimCons(tmp7273, Nil)

tmp7275 := PrimCons(symfn, tmp7274)

tmp7276 := PrimTail(V2855)

tmp7277 := PrimCons(tmp7275, tmp7276)

__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp7277)
return


} else {
__e.Return(V2855)
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


}, 1)

tmp7304 := PrimTail(V2855)

tmp7305 := Call(__e, PrimFunc(symlength), tmp7304)


__e.TailApply(tmp7266, tmp7305)
return


}, 1)

tmp7306 := PrimHead(V2855)

tmp7307 := Call(__e, PrimFunc(symarity), tmp7306)


__e.TailApply(tmp7265, tmp7307)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.process-application")))
return
}


}, 2)

tmp7310 := Call(__e, ns2_1set, symshen_4process_1application, tmp7264)


_ = tmp7310

tmp7311 := MakeNative(func(__e *ControlFlow) {
V2859 := __e.Get(1)
_ = V2859
tmp7337 := PrimIsPair(V2859)

var ifres7317 Obj

if True == tmp7337 {
tmp7335 := PrimHead(V2859)

tmp7336 := PrimIsPair(tmp7335)

var ifres7319 Obj

if True == tmp7336 {
tmp7332 := PrimHead(V2859)

tmp7333 := PrimHead(tmp7332)

tmp7334 := PrimEqual(symforeign, tmp7333)

var ifres7321 Obj

if True == tmp7334 {
tmp7329 := PrimHead(V2859)

tmp7330 := PrimTail(tmp7329)

tmp7331 := PrimIsPair(tmp7330)

var ifres7323 Obj

if True == tmp7331 {
tmp7325 := PrimHead(V2859)

tmp7326 := PrimTail(tmp7325)

tmp7327 := PrimTail(tmp7326)

tmp7328 := PrimEqual(Nil, tmp7327)

var ifres7324 Obj

if True == tmp7328 {
ifres7324 = True


} else {
ifres7324 = False


}

ifres7323 = ifres7324


} else {
ifres7323 = False


}

var ifres7322 Obj

if True == ifres7323 {
ifres7322 = True


} else {
ifres7322 = False


}

ifres7321 = ifres7322


} else {
ifres7321 = False


}

var ifres7320 Obj

if True == ifres7321 {
ifres7320 = True


} else {
ifres7320 = False


}

ifres7319 = ifres7320


} else {
ifres7319 = False


}

var ifres7318 Obj

if True == ifres7319 {
ifres7318 = True


} else {
ifres7318 = False


}

ifres7317 = ifres7318


} else {
ifres7317 = False


}

if True == ifres7317 {
tmp7312 := PrimHead(V2859)

tmp7313 := PrimTail(tmp7312)

tmp7314 := PrimHead(tmp7313)

tmp7315 := PrimTail(V2859)

__e.Return(PrimCons(tmp7314, tmp7315))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.unpack-foreign")))
return
}


}, 1)

tmp7338 := Call(__e, ns2_1set, symshen_4unpack_1foreign, tmp7311)


_ = tmp7338

tmp7339 := MakeNative(func(__e *ControlFlow) {
V2862 := __e.Get(1)
_ = V2862
tmp7361 := PrimIsPair(V2862)

var ifres7341 Obj

if True == tmp7361 {
tmp7359 := PrimHead(V2862)

tmp7360 := PrimIsPair(tmp7359)

var ifres7343 Obj

if True == tmp7360 {
tmp7356 := PrimHead(V2862)

tmp7357 := PrimHead(tmp7356)

tmp7358 := PrimEqual(symforeign, tmp7357)

var ifres7345 Obj

if True == tmp7358 {
tmp7353 := PrimHead(V2862)

tmp7354 := PrimTail(tmp7353)

tmp7355 := PrimIsPair(tmp7354)

var ifres7347 Obj

if True == tmp7355 {
tmp7349 := PrimHead(V2862)

tmp7350 := PrimTail(tmp7349)

tmp7351 := PrimTail(tmp7350)

tmp7352 := PrimEqual(Nil, tmp7351)

var ifres7348 Obj

if True == tmp7352 {
ifres7348 = True


} else {
ifres7348 = False


}

ifres7347 = ifres7348


} else {
ifres7347 = False


}

var ifres7346 Obj

if True == ifres7347 {
ifres7346 = True


} else {
ifres7346 = False


}

ifres7345 = ifres7346


} else {
ifres7345 = False


}

var ifres7344 Obj

if True == ifres7345 {
ifres7344 = True


} else {
ifres7344 = False


}

ifres7343 = ifres7344


} else {
ifres7343 = False


}

var ifres7342 Obj

if True == ifres7343 {
ifres7342 = True


} else {
ifres7342 = False


}

ifres7341 = ifres7342


} else {
ifres7341 = False


}

if True == ifres7341 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp7362 := Call(__e, ns2_1set, symshen_4foreign_2, tmp7339)


_ = tmp7362

tmp7363 := MakeNative(func(__e *ControlFlow) {
V2865 := __e.Get(1)
_ = V2865
tmp7369 := PrimIsPair(V2865)

var ifres7365 Obj

if True == tmp7369 {
tmp7367 := PrimTail(V2865)

tmp7368 := PrimEqual(Nil, tmp7367)

var ifres7366 Obj

if True == tmp7368 {
ifres7366 = True


} else {
ifres7366 = False


}

ifres7365 = ifres7366


} else {
ifres7365 = False


}

if True == ifres7365 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp7370 := Call(__e, ns2_1set, symshen_4zero_1place_2, tmp7363)


_ = tmp7370

tmp7371 := MakeNative(func(__e *ControlFlow) {
V2866 := __e.Get(1)
_ = V2866
tmp7376 := PrimIsSymbol(V2866)

if True == tmp7376 {
tmp7373 := PrimStr(V2866)

tmp7374 := Call(__e, PrimFunc(symshen_4internal_1to_1shen_2), tmp7373)


if True == tmp7374 {
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

tmp7377 := Call(__e, ns2_1set, symshen_4shen_1call_2, tmp7371)


_ = tmp7377

tmp7378 := MakeNative(func(__e *ControlFlow) {
V2871 := __e.Get(1)
_ = V2871
tmp7408 := PrimIsPair(V2871)

var ifres7395 Obj

if True == tmp7408 {
tmp7406 := PrimHead(V2871)

tmp7407 := PrimEqual(symprotect, tmp7406)

var ifres7397 Obj

if True == tmp7407 {
tmp7404 := PrimTail(V2871)

tmp7405 := PrimIsPair(tmp7404)

var ifres7399 Obj

if True == tmp7405 {
tmp7401 := PrimTail(V2871)

tmp7402 := PrimTail(tmp7401)

tmp7403 := PrimEqual(Nil, tmp7402)

var ifres7400 Obj

if True == tmp7403 {
ifres7400 = True


} else {
ifres7400 = False


}

ifres7399 = ifres7400


} else {
ifres7399 = False


}

var ifres7398 Obj

if True == ifres7399 {
ifres7398 = True


} else {
ifres7398 = False


}

ifres7397 = ifres7398


} else {
ifres7397 = False


}

var ifres7396 Obj

if True == ifres7397 {
ifres7396 = True


} else {
ifres7396 = False


}

ifres7395 = ifres7396


} else {
ifres7395 = False


}

if True == ifres7395 {
__e.Return(False)
return
} else {
tmp7393 := PrimIsPair(V2871)

var ifres7380 Obj

if True == tmp7393 {
tmp7391 := PrimHead(V2871)

tmp7392 := PrimEqual(symforeign, tmp7391)

var ifres7382 Obj

if True == tmp7392 {
tmp7389 := PrimTail(V2871)

tmp7390 := PrimIsPair(tmp7389)

var ifres7384 Obj

if True == tmp7390 {
tmp7386 := PrimTail(V2871)

tmp7387 := PrimTail(tmp7386)

tmp7388 := PrimEqual(Nil, tmp7387)

var ifres7385 Obj

if True == tmp7388 {
ifres7385 = True


} else {
ifres7385 = False


}

ifres7384 = ifres7385


} else {
ifres7384 = False


}

var ifres7383 Obj

if True == ifres7384 {
ifres7383 = True


} else {
ifres7383 = False


}

ifres7382 = ifres7383


} else {
ifres7382 = False


}

var ifres7381 Obj

if True == ifres7382 {
ifres7381 = True


} else {
ifres7381 = False


}

ifres7380 = ifres7381


} else {
ifres7380 = False


}

if True == ifres7380 {
__e.Return(False)
return
} else {
__e.Return(PrimIsPair(V2871))
return
}


}


}, 1)

tmp7409 := Call(__e, ns2_1set, symshen_4application_2, tmp7378)


_ = tmp7409

tmp7410 := MakeNative(func(__e *ControlFlow) {
V2876 := __e.Get(1)
_ = V2876
V2877 := __e.Get(2)
_ = V2877
tmp7418 := PrimEqual(MakeNumber(-1), V2877)

if True == tmp7418 {
tmp7416 := Call(__e, PrimFunc(symshen_4lowercase_1symbol_2), V2876)


if True == tmp7416 {
tmp7412 := Call(__e, PrimFunc(symexternal), symshen)


tmp7413 := Call(__e, PrimFunc(symelement_2), V2876, tmp7412)


tmp7414 := PrimNot(tmp7413)

if True == tmp7414 {
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


}, 2)

tmp7419 := Call(__e, ns2_1set, symshen_4undefined_1f_2, tmp7410)


_ = tmp7419

tmp7420 := MakeNative(func(__e *ControlFlow) {
V2878 := __e.Get(1)
_ = V2878
tmp7425 := PrimIsSymbol(V2878)

if True == tmp7425 {
tmp7422 := PrimIsVariable(V2878)

tmp7423 := PrimNot(tmp7422)

if True == tmp7423 {
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

tmp7426 := Call(__e, ns2_1set, symshen_4lowercase_1symbol_2, tmp7420)


_ = tmp7426

tmp7427 := MakeNative(func(__e *ControlFlow) {
V2879 := __e.Get(1)
_ = V2879
tmp7457 := PrimIsPair(V2879)

var ifres7448 Obj

if True == tmp7457 {
tmp7455 := PrimTail(V2879)

tmp7456 := PrimIsPair(tmp7455)

var ifres7450 Obj

if True == tmp7456 {
tmp7452 := PrimTail(V2879)

tmp7453 := PrimTail(tmp7452)

tmp7454 := PrimEqual(Nil, tmp7453)

var ifres7451 Obj

if True == tmp7454 {
ifres7451 = True


} else {
ifres7451 = False


}

ifres7450 = ifres7451


} else {
ifres7450 = False


}

var ifres7449 Obj

if True == ifres7450 {
ifres7449 = True


} else {
ifres7449 = False


}

ifres7448 = ifres7449


} else {
ifres7448 = False


}

if True == ifres7448 {
__e.Return(V2879)
return
} else {
tmp7446 := PrimIsPair(V2879)

var ifres7437 Obj

if True == tmp7446 {
tmp7444 := PrimTail(V2879)

tmp7445 := PrimIsPair(tmp7444)

var ifres7439 Obj

if True == tmp7445 {
tmp7441 := PrimTail(V2879)

tmp7442 := PrimTail(tmp7441)

tmp7443 := PrimIsPair(tmp7442)

var ifres7440 Obj

if True == tmp7443 {
ifres7440 = True


} else {
ifres7440 = False


}

ifres7439 = ifres7440


} else {
ifres7439 = False


}

var ifres7438 Obj

if True == ifres7439 {
ifres7438 = True


} else {
ifres7438 = False


}

ifres7437 = ifres7438


} else {
ifres7437 = False


}

if True == ifres7437 {
tmp7428 := PrimHead(V2879)

tmp7429 := PrimTail(V2879)

tmp7430 := PrimHead(tmp7429)

tmp7431 := PrimCons(tmp7430, Nil)

tmp7432 := PrimCons(tmp7428, tmp7431)

tmp7433 := PrimTail(V2879)

tmp7434 := PrimTail(tmp7433)

tmp7435 := PrimCons(tmp7432, tmp7434)

__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp7435)
return


} else {
__e.Return(V2879)
return
}


}


}, 1)

tmp7458 := Call(__e, ns2_1set, symshen_4simple_1curry, tmp7427)


_ = tmp7458

tmp7459 := MakeNative(func(__e *ControlFlow) {
V2880 := __e.Get(1)
_ = V2880
__e.TailApply(PrimFunc(symfn), V2880)
return
}, 1)

tmp7460 := Call(__e, ns2_1set, symfunction, tmp7459)


_ = tmp7460

tmp7461 := MakeNative(func(__e *ControlFlow) {
V2881 := __e.Get(1)
_ = V2881
tmp7470 := Call(__e, PrimFunc(symarity), V2881)


tmp7471 := PrimEqual(tmp7470, MakeNumber(0))

if True == tmp7471 {
__e.TailApply(V2881)
return
} else {
tmp7462 := MakeNative(func(__e *ControlFlow) {
W2882 := __e.Get(1)
_ = W2882
tmp7466 := Call(__e, PrimFunc(symempty_2), W2882)


if True == tmp7466 {
tmp7463 := Call(__e, PrimFunc(symshen_4app), V2881, MakeString(" is undefined\n"), symshen_4a)


tmp7464 := PrimStringConcat(MakeString("fn: "), tmp7463)

__e.Return(PrimSimpleError(tmp7464))
return


} else {
__e.Return(PrimTail(W2882))
return
}


}, 1)

tmp7467 := PrimValue(symshen_4_dlambdatable_d)

tmp7468 := Call(__e, PrimFunc(symassoc), V2881, tmp7467)


__e.TailApply(tmp7462, tmp7468)
return


}


}, 1)

tmp7472 := Call(__e, ns2_1set, symfn, tmp7461)


_ = tmp7472

tmp7473 := MakeNative(func(__e *ControlFlow) {
V2885 := __e.Get(1)
_ = V2885
tmp7503 := PrimIsPair(V2885)

var ifres7490 Obj

if True == tmp7503 {
tmp7501 := PrimHead(V2885)

tmp7502 := PrimEqual(symfn, tmp7501)

var ifres7492 Obj

if True == tmp7502 {
tmp7499 := PrimTail(V2885)

tmp7500 := PrimIsPair(tmp7499)

var ifres7494 Obj

if True == tmp7500 {
tmp7496 := PrimTail(V2885)

tmp7497 := PrimTail(tmp7496)

tmp7498 := PrimEqual(Nil, tmp7497)

var ifres7495 Obj

if True == tmp7498 {
ifres7495 = True


} else {
ifres7495 = False


}

ifres7494 = ifres7495


} else {
ifres7494 = False


}

var ifres7493 Obj

if True == ifres7494 {
ifres7493 = True


} else {
ifres7493 = False


}

ifres7492 = ifres7493


} else {
ifres7492 = False


}

var ifres7491 Obj

if True == ifres7492 {
ifres7491 = True


} else {
ifres7491 = False


}

ifres7490 = ifres7491


} else {
ifres7490 = False


}

if True == ifres7490 {
__e.Return(True)
return
} else {
tmp7488 := PrimIsPair(V2885)

var ifres7475 Obj

if True == tmp7488 {
tmp7486 := PrimHead(V2885)

tmp7487 := PrimEqual(symfunction, tmp7486)

var ifres7477 Obj

if True == tmp7487 {
tmp7484 := PrimTail(V2885)

tmp7485 := PrimIsPair(tmp7484)

var ifres7479 Obj

if True == tmp7485 {
tmp7481 := PrimTail(V2885)

tmp7482 := PrimTail(tmp7481)

tmp7483 := PrimEqual(Nil, tmp7482)

var ifres7480 Obj

if True == tmp7483 {
ifres7480 = True


} else {
ifres7480 = False


}

ifres7479 = ifres7480


} else {
ifres7479 = False


}

var ifres7478 Obj

if True == ifres7479 {
ifres7478 = True


} else {
ifres7478 = False


}

ifres7477 = ifres7478


} else {
ifres7477 = False


}

var ifres7476 Obj

if True == ifres7477 {
ifres7476 = True


} else {
ifres7476 = False


}

ifres7475 = ifres7476


} else {
ifres7475 = False


}

if True == ifres7475 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp7504 := Call(__e, ns2_1set, symshen_4fn_1call_2, tmp7473)


_ = tmp7504

tmp7505 := MakeNative(func(__e *ControlFlow) {
V2886 := __e.Get(1)
_ = V2886
tmp7546 := PrimIsPair(V2886)

var ifres7533 Obj

if True == tmp7546 {
tmp7544 := PrimHead(V2886)

tmp7545 := PrimEqual(symfunction, tmp7544)

var ifres7535 Obj

if True == tmp7545 {
tmp7542 := PrimTail(V2886)

tmp7543 := PrimIsPair(tmp7542)

var ifres7537 Obj

if True == tmp7543 {
tmp7539 := PrimTail(V2886)

tmp7540 := PrimTail(tmp7539)

tmp7541 := PrimEqual(Nil, tmp7540)

var ifres7538 Obj

if True == tmp7541 {
ifres7538 = True


} else {
ifres7538 = False


}

ifres7537 = ifres7538


} else {
ifres7537 = False


}

var ifres7536 Obj

if True == ifres7537 {
ifres7536 = True


} else {
ifres7536 = False


}

ifres7535 = ifres7536


} else {
ifres7535 = False


}

var ifres7534 Obj

if True == ifres7535 {
ifres7534 = True


} else {
ifres7534 = False


}

ifres7533 = ifres7534


} else {
ifres7533 = False


}

if True == ifres7533 {
tmp7506 := PrimTail(V2886)

tmp7507 := PrimCons(symfn, tmp7506)

__e.TailApply(PrimFunc(symshen_4fn_1call), tmp7507)
return


} else {
tmp7531 := PrimIsPair(V2886)

var ifres7518 Obj

if True == tmp7531 {
tmp7529 := PrimHead(V2886)

tmp7530 := PrimEqual(symfn, tmp7529)

var ifres7520 Obj

if True == tmp7530 {
tmp7527 := PrimTail(V2886)

tmp7528 := PrimIsPair(tmp7527)

var ifres7522 Obj

if True == tmp7528 {
tmp7524 := PrimTail(V2886)

tmp7525 := PrimTail(tmp7524)

tmp7526 := PrimEqual(Nil, tmp7525)

var ifres7523 Obj

if True == tmp7526 {
ifres7523 = True


} else {
ifres7523 = False


}

ifres7522 = ifres7523


} else {
ifres7522 = False


}

var ifres7521 Obj

if True == ifres7522 {
ifres7521 = True


} else {
ifres7521 = False


}

ifres7520 = ifres7521


} else {
ifres7520 = False


}

var ifres7519 Obj

if True == ifres7520 {
ifres7519 = True


} else {
ifres7519 = False


}

ifres7518 = ifres7519


} else {
ifres7518 = False


}

if True == ifres7518 {
tmp7508 := MakeNative(func(__e *ControlFlow) {
W2887 := __e.Get(1)
_ = W2887
tmp7513 := PrimEqual(W2887, MakeNumber(-1))

if True == tmp7513 {
__e.Return(V2886)
return
} else {
tmp7511 := PrimEqual(W2887, MakeNumber(0))

if True == tmp7511 {
__e.Return(PrimTail(V2886))
return
} else {
tmp7509 := PrimTail(V2886)

__e.TailApply(PrimFunc(symshen_4lambda_1function), tmp7509, W2887)
return


}


}


}, 1)

tmp7514 := PrimTail(V2886)

tmp7515 := PrimHead(tmp7514)

tmp7516 := Call(__e, PrimFunc(symarity), tmp7515)


__e.TailApply(tmp7508, tmp7516)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.fn-call")))
return
}


}


}, 1)

tmp7547 := Call(__e, ns2_1set, symshen_4fn_1call, tmp7505)


_ = tmp7547

tmp7548 := MakeNative(func(__e *ControlFlow) {
V2888 := __e.Get(1)
_ = V2888
V2889 := __e.Get(2)
_ = V2889
V2890 := __e.Get(3)
_ = V2890
tmp7549 := MakeNative(func(__e *ControlFlow) {
W2891 := __e.Get(1)
_ = W2891
tmp7550 := MakeNative(func(__e *ControlFlow) {
W2892 := __e.Get(1)
_ = W2892
__e.Return(W2891)
return
}, 1)

var ifres7556 Obj

if True == W2891 {
tmp7564 := Call(__e, PrimFunc(symshen_4loading_2))


var ifres7558 Obj

if True == tmp7564 {
tmp7560 := PrimCons(sym_1, Nil)

tmp7561 := PrimCons(sym_7, tmp7560)

tmp7562 := Call(__e, PrimFunc(symelement_2), V2888, tmp7561)


tmp7563 := PrimNot(tmp7562)

var ifres7559 Obj

if True == tmp7563 {
ifres7559 = True


} else {
ifres7559 = False


}

ifres7558 = ifres7559


} else {
ifres7558 = False


}

var ifres7557 Obj

if True == ifres7558 {
ifres7557 = True


} else {
ifres7557 = False


}

ifres7556 = ifres7557


} else {
ifres7556 = False


}

var ifres7551 Obj

if True == ifres7556 {
tmp7552 := Call(__e, PrimFunc(symshen_4app), V2888, MakeString("\n"), symshen_4a)


tmp7553 := PrimStringConcat(MakeString("partial application of "), tmp7552)

tmp7554 := Call(__e, PrimFunc(symstoutput))


tmp7555 := Call(__e, PrimFunc(sympr), tmp7553, tmp7554)


ifres7551 = tmp7555


} else {
ifres7551 = symshen_4skip


}

__e.TailApply(tmp7550, ifres7551)
return


}, 1)

tmp7565 := PrimGreatThan(V2889, V2890)

__e.TailApply(tmp7549, tmp7565)
return


}, 3)

tmp7566 := Call(__e, ns2_1set, symshen_4partial_1application_d_2, tmp7548)


_ = tmp7566

tmp7567 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dloading_2_d))
return
}, 0)

tmp7568 := Call(__e, ns2_1set, symshen_4loading_2, tmp7567)


_ = tmp7568

tmp7569 := MakeNative(func(__e *ControlFlow) {
V2897 := __e.Get(1)
_ = V2897
V2898 := __e.Get(2)
_ = V2898
V2899 := __e.Get(3)
_ = V2899
tmp7587 := PrimEqual(MakeNumber(-1), V2898)

if True == tmp7587 {
__e.Return(False)
return
} else {
tmp7570 := MakeNative(func(__e *ControlFlow) {
W2900 := __e.Get(1)
_ = W2900
tmp7571 := MakeNative(func(__e *ControlFlow) {
W2901 := __e.Get(1)
_ = W2901
__e.Return(W2900)
return
}, 1)

var ifres7582 Obj

if True == W2900 {
tmp7584 := Call(__e, PrimFunc(symshen_4loading_2))


var ifres7583 Obj

if True == tmp7584 {
ifres7583 = True


} else {
ifres7583 = False


}

ifres7582 = ifres7583


} else {
ifres7582 = False


}

var ifres7572 Obj

if True == ifres7582 {
tmp7574 := PrimEqual(V2899, MakeNumber(1))

var ifres7573 Obj

if True == tmp7574 {
ifres7573 = MakeString("")


} else {
ifres7573 = MakeString("s")


}

tmp7575 := Call(__e, PrimFunc(symshen_4app), ifres7573, MakeString("\n"), symshen_4a)


tmp7576 := PrimStringConcat(MakeString(" argument"), tmp7575)

tmp7577 := Call(__e, PrimFunc(symshen_4app), V2899, tmp7576, symshen_4a)


tmp7578 := PrimStringConcat(MakeString(" might not like "), tmp7577)

tmp7579 := Call(__e, PrimFunc(symshen_4app), V2897, tmp7578, symshen_4a)


tmp7580 := Call(__e, PrimFunc(symstoutput))


tmp7581 := Call(__e, PrimFunc(sympr), tmp7579, tmp7580)


ifres7572 = tmp7581


} else {
ifres7572 = symshen_4skip


}

__e.TailApply(tmp7571, ifres7572)
return


}, 1)

tmp7585 := PrimLessThan(V2898, V2899)

__e.TailApply(tmp7570, tmp7585)
return


}


}, 3)

__e.TailApply(ns2_1set, symshen_4overapplication_2, tmp7569)
return




}, 0)

