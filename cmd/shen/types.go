package main

import . "github.com/tiancaiamao/shen-go/kl"

var TypesMain = MakeNative(func(__e *ControlFlow) {
tmp16838 := MakeNative(func(__e *ControlFlow) {
V5907 := __e.Get(1)
_ = V5907
V5908 := __e.Get(2)
_ = V5908
tmp16839 := MakeNative(func(__e *ControlFlow) {
W5909 := __e.Get(1)
_ = W5909
tmp16840 := MakeNative(func(__e *ControlFlow) {
W5910 := __e.Get(1)
_ = W5910
tmp16841 := MakeNative(func(__e *ControlFlow) {
W5915 := __e.Get(1)
_ = W5915
tmp16842 := MakeNative(func(__e *ControlFlow) {
W5916 := __e.Get(1)
_ = W5916
__e.Return(V5907)
return
}, 1)

tmp16843 := PrimValue(symshen_4_dsigf_d)

tmp16844 := Call(__e, PrimFunc(symshen_4assoc_1_6), V5907, W5915, tmp16843)


tmp16845 := PrimSet(symshen_4_dsigf_d, tmp16844)

__e.TailApply(tmp16842, tmp16845)
return


}, 1)

tmp16846 := Call(__e, PrimFunc(symshen_4prolog_1abstraction), V5908)


tmp16847 := Call(__e, PrimFunc(symeval_1kl), tmp16846)


__e.TailApply(tmp16841, tmp16847)
return


}, 1)

tmp16848 := MakeNative(func(__e *ControlFlow) {
Z5911 := __e.Get(1)
_ = Z5911
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5912 := __e.Get(1)
_ = Z5912
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5913 := __e.Get(1)
_ = Z5913
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5914 := __e.Get(1)
_ = Z5914
tmp16849 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16849

tmp16850 := Call(__e, PrimFunc(symshen_4deref), V5907, Z5911)


tmp16851 := Call(__e, PrimFunc(symreceive), tmp16850)


tmp16852 := Call(__e, PrimFunc(symshen_4deref), W5909, Z5911)


tmp16853 := Call(__e, PrimFunc(symreceive), tmp16852)


__e.TailApply(PrimFunc(symshen_4variancy), tmp16851, tmp16853, Z5911, Z5912, Z5913, Z5914)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp16854 := Call(__e, PrimFunc(symshen_4prolog_1vector))


tmp16855 := Call(__e, tmp16848, tmp16854)


tmp16856 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp16857 := Call(__e, PrimFunc(sym_8v), MakeNumber(0), tmp16856)


tmp16858 := Call(__e, PrimFunc(sym_8v), True, tmp16857)


tmp16859 := Call(__e, tmp16855, tmp16858)


tmp16860 := Call(__e, tmp16859, MakeNumber(0))


tmp16861 := MakeNative(func(__e *ControlFlow) {
__e.Return(True)
return
}, 0)

tmp16862 := Call(__e, tmp16860, tmp16861)


__e.TailApply(tmp16840, tmp16862)
return


}, 1)

tmp16863 := Call(__e, PrimFunc(symshen_4rectify_1type), V5908)


__e.TailApply(tmp16839, tmp16863)
return


}, 2)

tmp16864 := Call(__e, ns2_1set, symdeclare, tmp16838)


_ = tmp16864

tmp16865 := MakeNative(func(__e *ControlFlow) {
V5917 := __e.Get(1)
_ = V5917
V5918 := __e.Get(2)
_ = V5918
V5919 := __e.Get(3)
_ = V5919
V5920 := __e.Get(4)
_ = V5920
V5921 := __e.Get(5)
_ = V5921
V5922 := __e.Get(6)
_ = V5922
tmp16872 := Call(__e, PrimFunc(symshen_4unlocked_2), V5920)


if True == tmp16872 {
tmp16866 := MakeNative(func(__e *ControlFlow) {
W5923 := __e.Get(1)
_ = W5923
tmp16867 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16867

tmp16868 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4variants_2), V5917, W5923, V5918, V5919, V5920, V5921, V5922)
return
}, 0)

tmp16869 := Call(__e, PrimFunc(symshen_4system_1S_1h), V5917, W5923, Nil, V5919, V5920, V5921, tmp16868)


__e.TailApply(PrimFunc(symshen_4gc), V5919, tmp16869)
return


}, 1)

tmp16870 := Call(__e, PrimFunc(symshen_4newpv), V5919)


__e.TailApply(tmp16866, tmp16870)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp16873 := Call(__e, ns2_1set, symshen_4variancy, tmp16865)


_ = tmp16873

tmp16874 := MakeNative(func(__e *ControlFlow) {
V5924 := __e.Get(1)
_ = V5924
V5925 := __e.Get(2)
_ = V5925
V5926 := __e.Get(3)
_ = V5926
V5927 := __e.Get(4)
_ = V5927
V5928 := __e.Get(5)
_ = V5928
V5929 := __e.Get(6)
_ = V5929
V5930 := __e.Get(7)
_ = V5930
tmp16875 := MakeNative(func(__e *ControlFlow) {
W5931 := __e.Get(1)
_ = W5931
tmp16876 := MakeNative(func(__e *ControlFlow) {
W5932 := __e.Get(1)
_ = W5932
tmp16900 := PrimEqual(W5932, False)

if True == tmp16900 {
tmp16877 := MakeNative(func(__e *ControlFlow) {
W5935 := __e.Get(1)
_ = W5935
tmp16894 := PrimEqual(W5935, False)

if True == tmp16894 {
tmp16878 := MakeNative(func(__e *ControlFlow) {
W5936 := __e.Get(1)
_ = W5936
tmp16880 := PrimEqual(W5936, False)

if True == tmp16880 {
__e.TailApply(PrimFunc(symshen_4unlock), V5928, W5931)
return
} else {
__e.Return(W5936)
return
}


}, 1)

tmp16892 := Call(__e, PrimFunc(symshen_4unlocked_2), V5928)


var ifres16881 Obj

if True == tmp16892 {
tmp16882 := MakeNative(func(__e *ControlFlow) {
W5937 := __e.Get(1)
_ = W5937
tmp16883 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16883

tmp16884 := Call(__e, PrimFunc(symshen_4deref), V5924, V5927)


tmp16885 := Call(__e, PrimFunc(symshen_4app), tmp16884, MakeString(" may create errors\n"), symshen_4a)


tmp16886 := PrimStringConcat(MakeString("warning: changing the type of "), tmp16885)

tmp16887 := Call(__e, PrimFunc(symstoutput))


tmp16888 := Call(__e, PrimFunc(sympr), tmp16886, tmp16887)


tmp16889 := Call(__e, PrimFunc(symis), W5937, tmp16888, V5927, V5928, W5931, V5930)


__e.TailApply(PrimFunc(symshen_4gc), V5927, tmp16889)
return


}, 1)

tmp16890 := Call(__e, PrimFunc(symshen_4newpv), V5927)


tmp16891 := Call(__e, tmp16882, tmp16890)


ifres16881 = tmp16891


} else {
ifres16881 = False


}

__e.TailApply(tmp16878, ifres16881)
return


} else {
__e.Return(W5935)
return
}


}, 1)

tmp16898 := Call(__e, PrimFunc(symshen_4unlocked_2), V5928)


var ifres16895 Obj

if True == tmp16898 {
tmp16896 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16896

tmp16897 := Call(__e, PrimFunc(symis_b), V5925, V5926, V5927, V5928, W5931, V5930)


ifres16895 = tmp16897


} else {
ifres16895 = False


}

__e.TailApply(tmp16877, ifres16895)
return


} else {
__e.Return(W5932)
return
}


}, 1)

tmp16912 := Call(__e, PrimFunc(symshen_4unlocked_2), V5928)


var ifres16901 Obj

if True == tmp16912 {
tmp16902 := MakeNative(func(__e *ControlFlow) {
W5933 := __e.Get(1)
_ = W5933
tmp16903 := MakeNative(func(__e *ControlFlow) {
W5934 := __e.Get(1)
_ = W5934
tmp16907 := PrimEqual(W5933, symsymbol)

if True == tmp16907 {
__e.TailApply(PrimFunc(symthaw), W5934)
return
} else {
tmp16905 := Call(__e, PrimFunc(symshen_4pvar_2), W5933)


if True == tmp16905 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5933, symsymbol, V5927, W5934)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp16908 := MakeNative(func(__e *ControlFlow) {
tmp16909 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp16909

__e.TailApply(PrimFunc(symshen_4cut), V5927, V5928, W5931, V5930)
return


}, 0)

__e.TailApply(tmp16903, tmp16908)
return


}, 1)

tmp16910 := Call(__e, PrimFunc(symshen_4lazyderef), V5925, V5927)


tmp16911 := Call(__e, tmp16902, tmp16910)


ifres16901 = tmp16911


} else {
ifres16901 = False


}

__e.TailApply(tmp16876, ifres16901)
return


}, 1)

tmp16913 := PrimNumberAdd(V5929, MakeNumber(1))

__e.TailApply(tmp16875, tmp16913)
return


}, 7)

tmp16914 := Call(__e, ns2_1set, symshen_4variants_2, tmp16874)


_ = tmp16914

tmp16915 := MakeNative(func(__e *ControlFlow) {
V5938 := __e.Get(1)
_ = V5938
tmp16916 := MakeNative(func(__e *ControlFlow) {
W5939 := __e.Get(1)
_ = W5939
tmp16917 := MakeNative(func(__e *ControlFlow) {
W5940 := __e.Get(1)
_ = W5940
tmp16918 := MakeNative(func(__e *ControlFlow) {
W5941 := __e.Get(1)
_ = W5941
tmp16919 := MakeNative(func(__e *ControlFlow) {
W5942 := __e.Get(1)
_ = W5942
tmp16920 := MakeNative(func(__e *ControlFlow) {
W5943 := __e.Get(1)
_ = W5943
tmp16921 := MakeNative(func(__e *ControlFlow) {
W5944 := __e.Get(1)
_ = W5944
tmp16922 := Call(__e, PrimFunc(symshen_4rcons__form), V5938)


tmp16923 := PrimCons(W5942, Nil)

tmp16924 := PrimCons(W5941, tmp16923)

tmp16925 := PrimCons(W5940, tmp16924)

tmp16926 := PrimCons(W5939, tmp16925)

tmp16927 := PrimCons(tmp16922, tmp16926)

tmp16928 := PrimCons(W5943, tmp16927)

tmp16929 := PrimCons(symis_b, tmp16928)

tmp16930 := Call(__e, PrimFunc(symshen_4stpart), W5944, tmp16929, W5939)


tmp16931 := PrimCons(tmp16930, Nil)

tmp16932 := PrimCons(W5942, tmp16931)

tmp16933 := PrimCons(symlambda, tmp16932)

tmp16934 := PrimCons(tmp16933, Nil)

tmp16935 := PrimCons(W5941, tmp16934)

tmp16936 := PrimCons(symlambda, tmp16935)

tmp16937 := PrimCons(tmp16936, Nil)

tmp16938 := PrimCons(W5940, tmp16937)

tmp16939 := PrimCons(symlambda, tmp16938)

tmp16940 := PrimCons(tmp16939, Nil)

tmp16941 := PrimCons(W5939, tmp16940)

tmp16942 := PrimCons(symlambda, tmp16941)

tmp16943 := PrimCons(tmp16942, Nil)

tmp16944 := PrimCons(W5943, tmp16943)

__e.Return(PrimCons(symlambda, tmp16944))
return


}, 1)

tmp16945 := Call(__e, PrimFunc(symshen_4extract_1vars), V5938)


__e.TailApply(tmp16921, tmp16945)
return


}, 1)

tmp16946 := Call(__e, PrimFunc(symgensym), symV)


__e.TailApply(tmp16920, tmp16946)
return


}, 1)

tmp16947 := Call(__e, PrimFunc(symgensym), symC)


__e.TailApply(tmp16919, tmp16947)
return


}, 1)

tmp16948 := Call(__e, PrimFunc(symgensym), symKey)


__e.TailApply(tmp16918, tmp16948)
return


}, 1)

tmp16949 := Call(__e, PrimFunc(symgensym), symL)


__e.TailApply(tmp16917, tmp16949)
return


}, 1)

tmp16950 := Call(__e, PrimFunc(symgensym), symB)


__e.TailApply(tmp16916, tmp16950)
return


}, 1)

tmp16951 := Call(__e, ns2_1set, symshen_4prolog_1abstraction, tmp16915)


_ = tmp16951

tmp16952 := MakeNative(func(__e *ControlFlow) {
V5945 := __e.Get(1)
_ = V5945
tmp16953 := MakeNative(func(__e *ControlFlow) {
W5946 := __e.Get(1)
_ = W5946
__e.TailApply(W5946, V5945)
return
}, 1)

tmp16954 := PrimValue(symshen_4_ddemodulation_1function_d)

__e.TailApply(tmp16953, tmp16954)
return


}, 1)

__e.TailApply(ns2_1set, symshen_4demod, tmp16952)
return




}, 0)

