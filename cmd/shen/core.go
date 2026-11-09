package main

import . "github.com/pyrex41/shen-go/kl"

var CoreMain = MakeNative(func(__e *ControlFlow) {
tmp2750 := MakeNative(func(__e *ControlFlow) {
V540 := __e.Get(1)
_ = V540
tmp2751 := Call(__e, PrimFunc(symshen_4shen_1_6kl_1h), V540)


let__2598 := tmp2751
_ = let__2598

__e.TailApply(PrimFunc(symshen_4record_1and_1evaluate), let__2598)
return


}, 1)

tmp2752 := Call(__e, ns2_1set, symshen_4shen_1_6kl, tmp2750)


_ = tmp2752

tmp2753 := MakeNative(func(__e *ControlFlow) {
V542 := __e.Get(1)
_ = V542
tmp2802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V542)
}
__typedArg0 := V542
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2776 Obj

if True == tmp2802 {
tmp2800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V542)
}
__typedArg0 := V542
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp2800)
}
__typedArg0 := symdefun
__typedArg1 := tmp2800
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2778 Obj

if True == tmp2801 {
tmp2798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V542)
}
__typedArg0 := V542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2798)
}
__typedArg0 := tmp2798
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2780 Obj

if True == tmp2799 {
tmp2795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V542)
}
__typedArg0 := V542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2795)
}
__typedArg0 := tmp2795
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2796)
}
__typedArg0 := tmp2796
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2782 Obj

if True == tmp2797 {
tmp2791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V542)
}
__typedArg0 := V542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2791)
}
__typedArg0 := tmp2791
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2792)
}
__typedArg0 := tmp2792
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2793)
}
__typedArg0 := tmp2793
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2784 Obj

if True == tmp2794 {
tmp2786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V542)
}
__typedArg0 := V542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2786)
}
__typedArg0 := tmp2786
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2787)
}
__typedArg0 := tmp2787
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2788)
}
__typedArg0 := tmp2788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2789)
}
__typedArg0 := Nil
__typedArg1 := tmp2789
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2785 Obj

if True == tmp2790 {
ifres2785 = True


} else {
ifres2785 = False


}

ifres2784 = ifres2785


} else {
ifres2784 = False


}

var ifres2783 Obj

if True == ifres2784 {
ifres2783 = True


} else {
ifres2783 = False


}

ifres2782 = ifres2783


} else {
ifres2782 = False


}

var ifres2781 Obj

if True == ifres2782 {
ifres2781 = True


} else {
ifres2781 = False


}

ifres2780 = ifres2781


} else {
ifres2780 = False


}

var ifres2779 Obj

if True == ifres2780 {
ifres2779 = True


} else {
ifres2779 = False


}

ifres2778 = ifres2779


} else {
ifres2778 = False


}

var ifres2777 Obj

if True == ifres2778 {
ifres2777 = True


} else {
ifres2777 = False


}

ifres2776 = ifres2777


} else {
ifres2776 = False


}

if True == ifres2776 {
tmp2759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V542)
}
__typedArg0 := V542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2759)
}
__typedArg0 := tmp2759
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2761 := Call(__e, PrimFunc(symshen_4sysfunc_2), tmp2760)


var ifres2754 Obj

if True == tmp2761 {
tmp2755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V542)
}
__typedArg0 := V542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2755)
}
__typedArg0 := tmp2755
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2757 := Call(__e, PrimFunc(symshen_4app), tmp2756, MakeString(" is not a legitimate function name\n"), symshen_4a)


tmp2758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp2757)
}
__typedArg0 := tmp2757
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres2754 = tmp2758


} else {
ifres2754 = symshen_4skip


}

let__2599 := ifres2754
_ = let__2599

tmp2762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V542)
}
__typedArg0 := V542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2762)
}
__typedArg0 := tmp2762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V542)
}
__typedArg0 := V542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2764)
}
__typedArg0 := tmp2764
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2765)
}
__typedArg0 := tmp2765
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2767 := Call(__e, PrimFunc(symlength), tmp2766)


tmp2768 := Call(__e, PrimFunc(symshen_4store_1arity), tmp2763, tmp2767)


let__2600 := tmp2768
_ = let__2600

tmp2769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V542)
}
__typedArg0 := V542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2769)
}
__typedArg0 := tmp2769
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2771 := Call(__e, PrimFunc(symshen_4record_1kl), tmp2770, V542)


let__2601 := tmp2771
_ = let__2601

tmp2772 := Call(__e, PrimFunc(symeval_1kl), V542)


let__2602 := tmp2772
_ = let__2602

tmp2773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V542)
}
__typedArg0 := V542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2773)
}
__typedArg0 := tmp2773
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4fn_1print), tmp2774)
return


} else {
__e.Return(V542)
return
}


}, 1)

tmp2803 := Call(__e, ns2_1set, symshen_4record_1and_1evaluate, tmp2753)


_ = tmp2803

tmp2804 := MakeNative(func(__e *ControlFlow) {
V547 := __e.Get(1)
_ = V547
tmp2905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2897 Obj

if True == tmp2905 {
tmp2903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp2903)
}
__typedArg0 := symdefine
__typedArg1 := tmp2903
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2899 Obj

if True == tmp2904 {
tmp2901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2901)
}
__typedArg0 := tmp2901
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2900 Obj

if True == tmp2902 {
ifres2900 = True


} else {
ifres2900 = False


}

ifres2899 = ifres2900


} else {
ifres2899 = False


}

var ifres2898 Obj

if True == ifres2899 {
ifres2898 = True


} else {
ifres2898 = False


}

ifres2897 = ifres2898


} else {
ifres2897 = False


}

if True == ifres2897 {
tmp2805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2805)
}
__typedArg0 := tmp2805
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2807)
}
__typedArg0 := tmp2807
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4shendef_1_6kldef), tmp2806, tmp2808)
return


} else {
tmp2895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2869 Obj

if True == tmp2895 {
tmp2893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp2893)
}
__typedArg0 := symdefun
__typedArg1 := tmp2893
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2871 Obj

if True == tmp2894 {
tmp2891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2891)
}
__typedArg0 := tmp2891
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2873 Obj

if True == tmp2892 {
tmp2888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2888)
}
__typedArg0 := tmp2888
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2889)
}
__typedArg0 := tmp2889
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2875 Obj

if True == tmp2890 {
tmp2884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2884)
}
__typedArg0 := tmp2884
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2885)
}
__typedArg0 := tmp2885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2886)
}
__typedArg0 := tmp2886
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2877 Obj

if True == tmp2887 {
tmp2879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2879)
}
__typedArg0 := tmp2879
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2880)
}
__typedArg0 := tmp2880
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2881)
}
__typedArg0 := tmp2881
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2882)
}
__typedArg0 := Nil
__typedArg1 := tmp2882
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2878 Obj

if True == tmp2883 {
ifres2878 = True


} else {
ifres2878 = False


}

ifres2877 = ifres2878


} else {
ifres2877 = False


}

var ifres2876 Obj

if True == ifres2877 {
ifres2876 = True


} else {
ifres2876 = False


}

ifres2875 = ifres2876


} else {
ifres2875 = False


}

var ifres2874 Obj

if True == ifres2875 {
ifres2874 = True


} else {
ifres2874 = False


}

ifres2873 = ifres2874


} else {
ifres2873 = False


}

var ifres2872 Obj

if True == ifres2873 {
ifres2872 = True


} else {
ifres2872 = False


}

ifres2871 = ifres2872


} else {
ifres2871 = False


}

var ifres2870 Obj

if True == ifres2871 {
ifres2870 = True


} else {
ifres2870 = False


}

ifres2869 = ifres2870


} else {
ifres2869 = False


}

if True == ifres2869 {
__e.Return(V547)
return
} else {
tmp2867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2848 Obj

if True == tmp2867 {
tmp2865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtype, tmp2865)
}
__typedArg0 := symtype
__typedArg1 := tmp2865
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2850 Obj

if True == tmp2866 {
tmp2863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2863)
}
__typedArg0 := tmp2863
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2852 Obj

if True == tmp2864 {
tmp2860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2860)
}
__typedArg0 := tmp2860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2861)
}
__typedArg0 := tmp2861
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2854 Obj

if True == tmp2862 {
tmp2856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2856)
}
__typedArg0 := tmp2856
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2857)
}
__typedArg0 := tmp2857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2858)
}
__typedArg0 := Nil
__typedArg1 := tmp2858
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2855 Obj

if True == tmp2859 {
ifres2855 = True


} else {
ifres2855 = False


}

ifres2854 = ifres2855


} else {
ifres2854 = False


}

var ifres2853 Obj

if True == ifres2854 {
ifres2853 = True


} else {
ifres2853 = False


}

ifres2852 = ifres2853


} else {
ifres2852 = False


}

var ifres2851 Obj

if True == ifres2852 {
ifres2851 = True


} else {
ifres2851 = False


}

ifres2850 = ifres2851


} else {
ifres2850 = False


}

var ifres2849 Obj

if True == ifres2850 {
ifres2849 = True


} else {
ifres2849 = False


}

ifres2848 = ifres2849


} else {
ifres2848 = False


}

if True == ifres2848 {
tmp2809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2809)
}
__typedArg0 := tmp2809
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2811)
}
__typedArg0 := tmp2811
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2812)
}
__typedArg0 := tmp2812
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2814 := Call(__e, PrimFunc(symshen_4rcons__form), tmp2813)


tmp2815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2814, Nil)
}
__typedArg0 := tmp2814
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2810, tmp2815)
}
__typedArg0 := tmp2810
__typedArg1 := tmp2815
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp2816)
}
__typedArg0 := symtype
__typedArg1 := tmp2816
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp2846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2827 Obj

if True == tmp2846 {
tmp2844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput_7, tmp2844)
}
__typedArg0 := syminput_7
__typedArg1 := tmp2844
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2829 Obj

if True == tmp2845 {
tmp2842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2842)
}
__typedArg0 := tmp2842
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2831 Obj

if True == tmp2843 {
tmp2839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2839)
}
__typedArg0 := tmp2839
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2840)
}
__typedArg0 := tmp2840
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2833 Obj

if True == tmp2841 {
tmp2835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2835)
}
__typedArg0 := tmp2835
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2836)
}
__typedArg0 := tmp2836
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2837)
}
__typedArg0 := Nil
__typedArg1 := tmp2837
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2834 Obj

if True == tmp2838 {
ifres2834 = True


} else {
ifres2834 = False


}

ifres2833 = ifres2834


} else {
ifres2833 = False


}

var ifres2832 Obj

if True == ifres2833 {
ifres2832 = True


} else {
ifres2832 = False


}

ifres2831 = ifres2832


} else {
ifres2831 = False


}

var ifres2830 Obj

if True == ifres2831 {
ifres2830 = True


} else {
ifres2830 = False


}

ifres2829 = ifres2830


} else {
ifres2829 = False


}

var ifres2828 Obj

if True == ifres2829 {
ifres2828 = True


} else {
ifres2828 = False


}

ifres2827 = ifres2828


} else {
ifres2827 = False


}

if True == ifres2827 {
tmp2817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2817)
}
__typedArg0 := tmp2817
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2819 := Call(__e, PrimFunc(symshen_4rcons__form), tmp2818)


tmp2820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2820)
}
__typedArg0 := tmp2820
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2819, tmp2821)
}
__typedArg0 := tmp2819
__typedArg1 := tmp2821
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput_7, tmp2822)
}
__typedArg0 := syminput_7
__typedArg1 := tmp2822
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp2825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2825 {
tmp2823 := MakeNative(func(__e *ControlFlow) {
Z548 := __e.Get(1)
_ = Z548
__e.TailApply(PrimFunc(symshen_4shen_1_6kl_1h), Z548)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp2823, V547)
return


} else {
__e.Return(V547)
return
}


}


}


}


}


}, 1)

tmp2906 := Call(__e, ns2_1set, symshen_4shen_1_6kl_1h, tmp2804)


_ = tmp2906

tmp2907 := MakeNative(func(__e *ControlFlow) {
V549 := __e.Get(1)
_ = V549
V550 := __e.Get(2)
_ = V550
tmp2908 := MakeNative(func(__e *ControlFlow) {
Z551 := __e.Get(1)
_ = Z551
__e.TailApply(PrimFunc(symshen_4_5define_6), Z551)
return
}, 1)

tmp2909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V549, V550)
}
__typedArg0 := V549
__typedArg1 := V550
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symcompile), tmp2908, tmp2909)
return


}, 2)

tmp2910 := Call(__e, ns2_1set, symshen_4shendef_1_6kldef, tmp2907)


_ = tmp2910

tmp2911 := MakeNative(func(__e *ControlFlow) {
V552 := __e.Get(1)
_ = V552
tmp2912 := Call(__e, PrimFunc(symshen_4_5name_6), V552)


let__2604 := tmp2912
_ = let__2604

tmp2938 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2604)


var ifres2913 Obj

if True == tmp2938 {
tmp2914 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2913 = tmp2914


} else {
tmp2915 := Call(__e, PrimFunc(symshen_4_5_1out), let__2604)


let__2605 := tmp2915
_ = let__2605

tmp2916 := Call(__e, PrimFunc(symshen_4in_1_6), let__2604)


let__2606 := tmp2916
_ = let__2606

tmp2937 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2606, sym_i)


var ifres2917 Obj

if True == tmp2937 {
tmp2918 := Call(__e, PrimFunc(symtail), let__2606)


let__2607 := tmp2918
_ = let__2607

tmp2919 := Call(__e, PrimFunc(symshen_4_5signature_6), let__2607)


let__2608 := tmp2919
_ = let__2608

tmp2935 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2608)


var ifres2920 Obj

if True == tmp2935 {
tmp2921 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2920 = tmp2921


} else {
tmp2922 := Call(__e, PrimFunc(symshen_4in_1_6), let__2608)


let__2609 := tmp2922
_ = let__2609

tmp2934 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2609, sym_j)


var ifres2923 Obj

if True == tmp2934 {
tmp2924 := Call(__e, PrimFunc(symtail), let__2609)


let__2610 := tmp2924
_ = let__2610

tmp2925 := Call(__e, PrimFunc(symshen_4_5rules_6), let__2610)


let__2611 := tmp2925
_ = let__2611

tmp2932 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2611)


var ifres2926 Obj

if True == tmp2932 {
tmp2927 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2926 = tmp2927


} else {
tmp2928 := Call(__e, PrimFunc(symshen_4_5_1out), let__2611)


let__2612 := tmp2928
_ = let__2612

tmp2929 := Call(__e, PrimFunc(symshen_4in_1_6), let__2611)


let__2613 := tmp2929
_ = let__2613

tmp2930 := Call(__e, PrimFunc(symshen_4shendef_1_6kldef_1h), let__2605, let__2612)


tmp2931 := Call(__e, PrimFunc(symshen_4comb), let__2613, tmp2930)


ifres2926 = tmp2931


}

ifres2923 = ifres2926


} else {
tmp2933 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2923 = tmp2933


}

ifres2920 = ifres2923


}

ifres2917 = ifres2920


} else {
tmp2936 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2917 = tmp2936


}

ifres2913 = ifres2917


}

let__2603 := ifres2913
_ = let__2603

tmp2956 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2603)


if True == tmp2956 {
tmp2939 := Call(__e, PrimFunc(symshen_4_5name_6), V552)


let__2615 := tmp2939
_ = let__2615

tmp2952 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2615)


var ifres2940 Obj

if True == tmp2952 {
tmp2941 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2940 = tmp2941


} else {
tmp2942 := Call(__e, PrimFunc(symshen_4_5_1out), let__2615)


let__2616 := tmp2942
_ = let__2616

tmp2943 := Call(__e, PrimFunc(symshen_4in_1_6), let__2615)


let__2617 := tmp2943
_ = let__2617

tmp2944 := Call(__e, PrimFunc(symshen_4_5rules_6), let__2617)


let__2618 := tmp2944
_ = let__2618

tmp2951 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2618)


var ifres2945 Obj

if True == tmp2951 {
tmp2946 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2945 = tmp2946


} else {
tmp2947 := Call(__e, PrimFunc(symshen_4_5_1out), let__2618)


let__2619 := tmp2947
_ = let__2619

tmp2948 := Call(__e, PrimFunc(symshen_4in_1_6), let__2618)


let__2620 := tmp2948
_ = let__2620

tmp2949 := Call(__e, PrimFunc(symshen_4shendef_1_6kldef_1h), let__2616, let__2619)


tmp2950 := Call(__e, PrimFunc(symshen_4comb), let__2620, tmp2949)


ifres2945 = tmp2950


}

ifres2940 = ifres2945


}

let__2614 := ifres2940
_ = let__2614

tmp2954 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2614)


if True == tmp2954 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2614)
return
}


} else {
__e.Return(let__2603)
return
}


}, 1)

tmp2957 := Call(__e, ns2_1set, symshen_4_5define_6, tmp2911)


_ = tmp2957

tmp2958 := MakeNative(func(__e *ControlFlow) {
V571 := __e.Get(1)
_ = V571
V572 := __e.Get(2)
_ = V572
tmp2959 := MakeNative(func(__e *ControlFlow) {
Z574 := __e.Get(1)
_ = Z574
__e.TailApply(PrimFunc(symfst), Z574)
return
}, 1)

tmp2960 := Call(__e, PrimFunc(symmap), tmp2959, V572)


let__2621 := tmp2960
_ = let__2621

tmp2961 := Call(__e, PrimFunc(symshen_4arity_1chk), V571, let__2621)


let__2622 := tmp2961
_ = let__2622

tmp2962 := MakeNative(func(__e *ControlFlow) {
Z577 := __e.Get(1)
_ = Z577
__e.TailApply(PrimFunc(symshen_4free_1var_1chk), V571, Z577)
return
}, 1)

tmp2963 := Call(__e, PrimFunc(symmap), tmp2962, V572)


let__2623 := tmp2963
_ = let__2623

tmp2964 := Call(__e, PrimFunc(symshen_4unprotect), V572)


let__2624 := tmp2964
_ = let__2624

tmp2965 := Call(__e, PrimFunc(symshen_4compile_1to_1kl), V571, let__2624, let__2622)


tmp2966 := Call(__e, PrimFunc(symshen_4factorise_1code), tmp2965)


let__2625 := tmp2966
_ = let__2625

__e.Return(let__2625)
return


}, 2)

tmp2967 := Call(__e, ns2_1set, symshen_4shendef_1_6kldef_1h, tmp2958)


_ = tmp2967

tmp2968 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V580 := __self1
_ = V580
tmp2994 := Call(__e, PrimFunc(symtuple_2), V580)


if True == tmp2994 {
tmp2969 := Call(__e, PrimFunc(symfst), V580)


tmp2970 := Call(__e, PrimFunc(symshen_4unprotect), tmp2969)


tmp2971 := Call(__e, PrimFunc(symsnd), V580)


tmp2972 := Call(__e, PrimFunc(symshen_4unprotect), tmp2971)


__e.TailApply(PrimFunc(sym_8p), tmp2970, tmp2972)
return


} else {
tmp2992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V580)
}
__typedArg0 := V580
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2979 Obj

if True == tmp2992 {
tmp2990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V580)
}
__typedArg0 := V580
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symprotect, tmp2990)
}
__typedArg0 := symprotect
__typedArg1 := tmp2990
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2981 Obj

if True == tmp2991 {
tmp2988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V580)
}
__typedArg0 := V580
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2988)
}
__typedArg0 := tmp2988
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2983 Obj

if True == tmp2989 {
tmp2985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V580)
}
__typedArg0 := V580
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2985)
}
__typedArg0 := tmp2985
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2986)
}
__typedArg0 := Nil
__typedArg1 := tmp2986
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2984 Obj

if True == tmp2987 {
ifres2984 = True


} else {
ifres2984 = False


}

ifres2983 = ifres2984


} else {
ifres2983 = False


}

var ifres2982 Obj

if True == ifres2983 {
ifres2982 = True


} else {
ifres2982 = False


}

ifres2981 = ifres2982


} else {
ifres2981 = False


}

var ifres2980 Obj

if True == ifres2981 {
ifres2980 = True


} else {
ifres2980 = False


}

ifres2979 = ifres2980


} else {
ifres2979 = False


}

if True == ifres2979 {
tmp2973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V580)
}
__typedArg0 := V580
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2973)
}
__typedArg0 := tmp2973
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

if PrimFunc(symshen_4unprotect) == __self {
__self1 = tmp2974
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4unprotect), tmp2974)
return


} else {
tmp2977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V580)
}
__typedArg0 := V580
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2977 {
tmp2975 := MakeNative(func(__e *ControlFlow) {
Z581 := __e.Get(1)
_ = Z581
__e.TailApply(PrimFunc(symshen_4unprotect), Z581)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp2975, V580)
return


} else {
__e.Return(V580)
return
}


}


}


}, 1)

tmp2995 := Call(__e, ns2_1set, symshen_4unprotect, tmp2968)


_ = tmp2995

tmp2996 := MakeNative(func(__e *ControlFlow) {
V582 := __e.Get(1)
_ = V582
tmp3010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V582)
}
__typedArg0 := V582
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2997 Obj

if True == tmp3010 {
tmp2998 := Call(__e, PrimFunc(symhead), V582)


let__2627 := tmp2998
_ = let__2627

tmp2999 := Call(__e, PrimFunc(symtail), V582)


let__2628 := tmp2999
_ = let__2628

tmp3007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(let__2627)
}
__typedArg0 := let__2627
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

var ifres3003 Obj

if True == tmp3007 {
tmp3005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(let__2627)
}
__typedArg0 := let__2627
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

tmp3006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp3005)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp3005
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres3004 Obj

if True == tmp3006 {
ifres3004 = True


} else {
ifres3004 = False


}

ifres3003 = ifres3004


} else {
ifres3003 = False


}

var ifres3000 Obj

if True == ifres3003 {
ifres3000 = let__2627


} else {
tmp3001 := Call(__e, PrimFunc(symshen_4app), let__2627, MakeString(" is not a legitimate function name.\n"), symshen_4a)


tmp3002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp3001)
}
__typedArg0 := tmp3001
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres3000 = tmp3002


}

tmp3008 := Call(__e, PrimFunc(symshen_4comb), let__2628, ifres3000)


ifres2997 = tmp3008


} else {
tmp3009 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2997 = tmp3009


}

let__2626 := ifres2997
_ = let__2626

tmp3012 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2626)


if True == tmp3012 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2626)
return
}


}, 1)

tmp3013 := Call(__e, ns2_1set, symshen_4_5name_6, tmp2996)


_ = tmp3013

tmp3014 := MakeNative(func(__e *ControlFlow) {
V586 := __e.Get(1)
_ = V586
tmp3033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V586)
}
__typedArg0 := V586
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3015 Obj

if True == tmp3033 {
tmp3016 := Call(__e, PrimFunc(symhead), V586)


let__2630 := tmp3016
_ = let__2630

tmp3017 := Call(__e, PrimFunc(symtail), V586)


let__2631 := tmp3017
_ = let__2631

tmp3018 := Call(__e, PrimFunc(symshen_4_5signature_6), let__2631)


let__2632 := tmp3018
_ = let__2632

tmp3031 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2632)


var ifres3019 Obj

if True == tmp3031 {
tmp3020 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3019 = tmp3020


} else {
tmp3021 := Call(__e, PrimFunc(symshen_4_5_1out), let__2632)


let__2633 := tmp3021
_ = let__2633

tmp3022 := Call(__e, PrimFunc(symshen_4in_1_6), let__2632)


let__2634 := tmp3022
_ = let__2634

tmp3027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, Nil)
}
__typedArg0 := sym_j
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, tmp3027)
}
__typedArg0 := sym_i
__typedArg1 := tmp3027
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3029 := Call(__e, PrimFunc(symelement_2), let__2630, tmp3028)


tmp3030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp3029)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp3029
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres3023 Obj

if True == tmp3030 {
tmp3024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2630, let__2633)
}
__typedArg0 := let__2630
__typedArg1 := let__2633
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3025 := Call(__e, PrimFunc(symshen_4comb), let__2634, tmp3024)


ifres3023 = tmp3025


} else {
tmp3026 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3023 = tmp3026


}

ifres3019 = ifres3023


}

ifres3015 = ifres3019


} else {
tmp3032 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3015 = tmp3032


}

let__2629 := ifres3015
_ = let__2629

tmp3043 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2629)


if True == tmp3043 {
tmp3034 := Call(__e, PrimFunc(sym_5e_6), V586)


let__2636 := tmp3034
_ = let__2636

tmp3039 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2636)


var ifres3035 Obj

if True == tmp3039 {
tmp3036 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3035 = tmp3036


} else {
tmp3037 := Call(__e, PrimFunc(symshen_4in_1_6), let__2636)


let__2637 := tmp3037
_ = let__2637

tmp3038 := Call(__e, PrimFunc(symshen_4comb), let__2637, Nil)


ifres3035 = tmp3038


}

let__2635 := ifres3035
_ = let__2635

tmp3041 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2635)


if True == tmp3041 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2635)
return
}


} else {
__e.Return(let__2629)
return
}


}, 1)

tmp3044 := Call(__e, ns2_1set, symshen_4_5signature_6, tmp3014)


_ = tmp3044

tmp3045 := MakeNative(func(__e *ControlFlow) {
V596 := __e.Get(1)
_ = V596
tmp3046 := Call(__e, PrimFunc(symshen_4_5rule_6), V596)


let__2639 := tmp3046
_ = let__2639

tmp3060 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2639)


var ifres3047 Obj

if True == tmp3060 {
tmp3048 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3047 = tmp3048


} else {
tmp3049 := Call(__e, PrimFunc(symshen_4_5_1out), let__2639)


let__2640 := tmp3049
_ = let__2640

tmp3050 := Call(__e, PrimFunc(symshen_4in_1_6), let__2639)


let__2641 := tmp3050
_ = let__2641

tmp3051 := Call(__e, PrimFunc(symshen_4_5rules_6), let__2641)


let__2642 := tmp3051
_ = let__2642

tmp3059 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2642)


var ifres3052 Obj

if True == tmp3059 {
tmp3053 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3052 = tmp3053


} else {
tmp3054 := Call(__e, PrimFunc(symshen_4_5_1out), let__2642)


let__2643 := tmp3054
_ = let__2643

tmp3055 := Call(__e, PrimFunc(symshen_4in_1_6), let__2642)


let__2644 := tmp3055
_ = let__2644

tmp3056 := Call(__e, PrimFunc(symshen_4linearise), let__2640)


tmp3057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3056, let__2643)
}
__typedArg0 := tmp3056
__typedArg1 := let__2643
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3058 := Call(__e, PrimFunc(symshen_4comb), let__2644, tmp3057)


ifres3052 = tmp3058


}

ifres3047 = ifres3052


}

let__2638 := ifres3047
_ = let__2638

tmp3076 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2638)


if True == tmp3076 {
tmp3061 := Call(__e, PrimFunc(sym_5_b_6), V596)


let__2646 := tmp3061
_ = let__2646

tmp3072 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2646)


var ifres3062 Obj

if True == tmp3072 {
tmp3063 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3062 = tmp3063


} else {
tmp3064 := Call(__e, PrimFunc(symshen_4_5_1out), let__2646)


let__2647 := tmp3064
_ = let__2647

tmp3065 := Call(__e, PrimFunc(symshen_4in_1_6), let__2646)


let__2648 := tmp3065
_ = let__2648

tmp3070 := Call(__e, PrimFunc(symempty_2), let__2647)


var ifres3066 Obj

if True == tmp3070 {
ifres3066 = Nil


} else {
tmp3067 := Call(__e, PrimFunc(symshen_4app), let__2647, MakeString("\n ..."), symshen_4r)


tmp3069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Shen syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp3067)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Shen syntax error here:\n ")
__typedArg1 := tmp3067
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Shen syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp3067)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Shen syntax error here:\n ")
__typedArg1 := tmp3067
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres3066 = tmp3069


}

tmp3071 := Call(__e, PrimFunc(symshen_4comb), let__2648, ifres3066)


ifres3062 = tmp3071


}

let__2645 := ifres3062
_ = let__2645

tmp3074 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2645)


if True == tmp3074 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2645)
return
}


} else {
__e.Return(let__2638)
return
}


}, 1)

tmp3077 := Call(__e, ns2_1set, symshen_4_5rules_6, tmp3045)


_ = tmp3077

tmp3078 := MakeNative(func(__e *ControlFlow) {
V610 := __e.Get(1)
_ = V610
tmp3083 := Call(__e, PrimFunc(symtuple_2), V610)


if True == tmp3083 {
tmp3079 := Call(__e, PrimFunc(symfst), V610)


tmp3080 := Call(__e, PrimFunc(symfst), V610)


tmp3081 := Call(__e, PrimFunc(symsnd), V610)


__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp3079, tmp3080, Nil, tmp3081)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.linearise"))
}
__typedArg0 := MakeString("implementation error in shen.linearise")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp3084 := Call(__e, ns2_1set, symshen_4linearise, tmp3078)


_ = tmp3084

tmp3085 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__selftop:
V623 := __self1
_ = V623
V624 := __self2
_ = V624
V625 := __self3
_ = V625
V626 := __self4
_ = V626
tmp3122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V623)
}
__typedArg0 := Nil
__typedArg1 := V623
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp3122 {
__e.TailApply(PrimFunc(sym_8p), V624, V626)
return
} else {
tmp3120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3116 Obj

if True == tmp3120 {
tmp3118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3118)
}
__typedArg0 := tmp3118
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3117 Obj

if True == tmp3119 {
ifres3117 = True


} else {
ifres3117 = False


}

ifres3116 = ifres3117


} else {
ifres3116 = False


}

if True == ifres3116 {
tmp3086 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3088 := Call(__e, PrimFunc(symappend), tmp3086, tmp3087)


if PrimFunc(symshen_4linearise_1h) == __self {
__self1, __self2, __self3, __self4 = tmp3088, V624, V625, V626
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp3088, V624, V625, V626)
return


} else {
tmp3114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3110 Obj

if True == tmp3114 {
tmp3112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp3112)
}
__typedArg0 := tmp3112
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres3111 Obj

if True == tmp3113 {
ifres3111 = True


} else {
ifres3111 = False


}

ifres3110 = ifres3111


} else {
ifres3110 = False


}

if True == ifres3110 {
tmp3104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3105 := Call(__e, PrimFunc(symelement_2), tmp3104, V625)


if True == tmp3105 {
tmp3089 := Call(__e, PrimFunc(symgensym), symV)


let__2649 := tmp3089
_ = let__2649

tmp3090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3092 := Call(__e, PrimFunc(symshen_4rep_1X), tmp3091, let__2649, V624)


tmp3093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3093, Nil)
}
__typedArg0 := tmp3093
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2649, tmp3094)
}
__typedArg0 := let__2649
__typedArg1 := tmp3094
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp3095)
}
__typedArg0 := sym_a
__typedArg1 := tmp3095
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V626, Nil)
}
__typedArg0 := V626
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3096, tmp3097)
}
__typedArg0 := tmp3096
__typedArg1 := tmp3097
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp3098)
}
__typedArg0 := symwhere
__typedArg1 := tmp3098
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4linearise_1h) == __self {
__self1, __self2, __self3, __self4 = tmp3090, tmp3092, V625, tmp3099
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp3090, tmp3092, V625, tmp3099)
return


} else {
tmp3100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3101, V625)
}
__typedArg0 := tmp3101
__typedArg1 := V625
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4linearise_1h) == __self {
__self1, __self2, __self3, __self4 = tmp3100, V624, tmp3102, V626
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp3100, V624, tmp3102, V626)
return


}


} else {
tmp3108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3108 {
tmp3106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4linearise_1h) == __self {
__self1, __self2, __self3, __self4 = tmp3106, V624, V625, V626
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp3106, V624, V625, V626)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.linearise-h"))
}
__typedArg0 := MakeString("implementation error in shen.linearise-h")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}


}, 4)

tmp3123 := Call(__e, ns2_1set, symshen_4linearise_1h, tmp3085)


_ = tmp3123

tmp3124 := MakeNative(func(__e *ControlFlow) {
V628 := __e.Get(1)
_ = V628
tmp3125 := Call(__e, PrimFunc(symshen_4_5patterns_6), V628)


let__2651 := tmp3125
_ = let__2651

tmp3153 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2651)


var ifres3126 Obj

if True == tmp3153 {
tmp3127 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3126 = tmp3127


} else {
tmp3128 := Call(__e, PrimFunc(symshen_4_5_1out), let__2651)


let__2652 := tmp3128
_ = let__2652

tmp3129 := Call(__e, PrimFunc(symshen_4in_1_6), let__2651)


let__2653 := tmp3129
_ = let__2653

tmp3152 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2653, sym_1_6)


var ifres3130 Obj

if True == tmp3152 {
tmp3131 := Call(__e, PrimFunc(symtail), let__2653)


let__2654 := tmp3131
_ = let__2654

tmp3150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__2654)
}
__typedArg0 := let__2654
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3132 Obj

if True == tmp3150 {
tmp3133 := Call(__e, PrimFunc(symhead), let__2654)


let__2655 := tmp3133
_ = let__2655

tmp3134 := Call(__e, PrimFunc(symtail), let__2654)


let__2656 := tmp3134
_ = let__2656

tmp3148 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2656, symwhere)


var ifres3135 Obj

if True == tmp3148 {
tmp3136 := Call(__e, PrimFunc(symtail), let__2656)


let__2657 := tmp3136
_ = let__2657

tmp3146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__2657)
}
__typedArg0 := let__2657
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3137 Obj

if True == tmp3146 {
tmp3138 := Call(__e, PrimFunc(symhead), let__2657)


let__2658 := tmp3138
_ = let__2658

tmp3139 := Call(__e, PrimFunc(symtail), let__2657)


let__2659 := tmp3139
_ = let__2659

tmp3140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2655, Nil)
}
__typedArg0 := let__2655
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2658, tmp3140)
}
__typedArg0 := let__2658
__typedArg1 := tmp3140
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp3141)
}
__typedArg0 := symwhere
__typedArg1 := tmp3141
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3143 := Call(__e, PrimFunc(sym_8p), let__2652, tmp3142)


tmp3144 := Call(__e, PrimFunc(symshen_4comb), let__2659, tmp3143)


ifres3137 = tmp3144


} else {
tmp3145 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3137 = tmp3145


}

ifres3135 = ifres3137


} else {
tmp3147 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3135 = tmp3147


}

ifres3132 = ifres3135


} else {
tmp3149 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3132 = tmp3149


}

ifres3130 = ifres3132


} else {
tmp3151 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3130 = tmp3151


}

ifres3126 = ifres3130


}

let__2650 := ifres3126
_ = let__2650

tmp3228 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2650)


if True == tmp3228 {
tmp3154 := Call(__e, PrimFunc(symshen_4_5patterns_6), V628)


let__2661 := tmp3154
_ = let__2661

tmp3170 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2661)


var ifres3155 Obj

if True == tmp3170 {
tmp3156 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3155 = tmp3156


} else {
tmp3157 := Call(__e, PrimFunc(symshen_4_5_1out), let__2661)


let__2662 := tmp3157
_ = let__2662

tmp3158 := Call(__e, PrimFunc(symshen_4in_1_6), let__2661)


let__2663 := tmp3158
_ = let__2663

tmp3169 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2663, sym_1_6)


var ifres3159 Obj

if True == tmp3169 {
tmp3160 := Call(__e, PrimFunc(symtail), let__2663)


let__2664 := tmp3160
_ = let__2664

tmp3167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__2664)
}
__typedArg0 := let__2664
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3161 Obj

if True == tmp3167 {
tmp3162 := Call(__e, PrimFunc(symhead), let__2664)


let__2665 := tmp3162
_ = let__2665

tmp3163 := Call(__e, PrimFunc(symtail), let__2664)


let__2666 := tmp3163
_ = let__2666

tmp3164 := Call(__e, PrimFunc(sym_8p), let__2662, let__2665)


tmp3165 := Call(__e, PrimFunc(symshen_4comb), let__2666, tmp3164)


ifres3161 = tmp3165


} else {
tmp3166 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3161 = tmp3166


}

ifres3159 = ifres3161


} else {
tmp3168 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3159 = tmp3168


}

ifres3155 = ifres3159


}

let__2660 := ifres3155
_ = let__2660

tmp3226 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2660)


if True == tmp3226 {
tmp3171 := Call(__e, PrimFunc(symshen_4_5patterns_6), V628)


let__2668 := tmp3171
_ = let__2668

tmp3201 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2668)


var ifres3172 Obj

if True == tmp3201 {
tmp3173 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3172 = tmp3173


} else {
tmp3174 := Call(__e, PrimFunc(symshen_4_5_1out), let__2668)


let__2669 := tmp3174
_ = let__2669

tmp3175 := Call(__e, PrimFunc(symshen_4in_1_6), let__2668)


let__2670 := tmp3175
_ = let__2670

tmp3200 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2670, sym_5_1)


var ifres3176 Obj

if True == tmp3200 {
tmp3177 := Call(__e, PrimFunc(symtail), let__2670)


let__2671 := tmp3177
_ = let__2671

tmp3198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__2671)
}
__typedArg0 := let__2671
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3178 Obj

if True == tmp3198 {
tmp3179 := Call(__e, PrimFunc(symhead), let__2671)


let__2672 := tmp3179
_ = let__2672

tmp3180 := Call(__e, PrimFunc(symtail), let__2671)


let__2673 := tmp3180
_ = let__2673

tmp3196 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2673, symwhere)


var ifres3181 Obj

if True == tmp3196 {
tmp3182 := Call(__e, PrimFunc(symtail), let__2673)


let__2674 := tmp3182
_ = let__2674

tmp3194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__2674)
}
__typedArg0 := let__2674
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3183 Obj

if True == tmp3194 {
tmp3184 := Call(__e, PrimFunc(symhead), let__2674)


let__2675 := tmp3184
_ = let__2675

tmp3185 := Call(__e, PrimFunc(symtail), let__2674)


let__2676 := tmp3185
_ = let__2676

tmp3186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2672, Nil)
}
__typedArg0 := let__2672
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4choicepoint_b, tmp3186)
}
__typedArg0 := symshen_4choicepoint_b
__typedArg1 := tmp3186
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3187, Nil)
}
__typedArg0 := tmp3187
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2675, tmp3188)
}
__typedArg0 := let__2675
__typedArg1 := tmp3188
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp3189)
}
__typedArg0 := symwhere
__typedArg1 := tmp3189
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3191 := Call(__e, PrimFunc(sym_8p), let__2669, tmp3190)


tmp3192 := Call(__e, PrimFunc(symshen_4comb), let__2676, tmp3191)


ifres3183 = tmp3192


} else {
tmp3193 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3183 = tmp3193


}

ifres3181 = ifres3183


} else {
tmp3195 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3181 = tmp3195


}

ifres3178 = ifres3181


} else {
tmp3197 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3178 = tmp3197


}

ifres3176 = ifres3178


} else {
tmp3199 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3176 = tmp3199


}

ifres3172 = ifres3176


}

let__2667 := ifres3172
_ = let__2667

tmp3224 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2667)


if True == tmp3224 {
tmp3202 := Call(__e, PrimFunc(symshen_4_5patterns_6), V628)


let__2678 := tmp3202
_ = let__2678

tmp3220 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2678)


var ifres3203 Obj

if True == tmp3220 {
tmp3204 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3203 = tmp3204


} else {
tmp3205 := Call(__e, PrimFunc(symshen_4_5_1out), let__2678)


let__2679 := tmp3205
_ = let__2679

tmp3206 := Call(__e, PrimFunc(symshen_4in_1_6), let__2678)


let__2680 := tmp3206
_ = let__2680

tmp3219 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2680, sym_5_1)


var ifres3207 Obj

if True == tmp3219 {
tmp3208 := Call(__e, PrimFunc(symtail), let__2680)


let__2681 := tmp3208
_ = let__2681

tmp3217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__2681)
}
__typedArg0 := let__2681
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3209 Obj

if True == tmp3217 {
tmp3210 := Call(__e, PrimFunc(symhead), let__2681)


let__2682 := tmp3210
_ = let__2682

tmp3211 := Call(__e, PrimFunc(symtail), let__2681)


let__2683 := tmp3211
_ = let__2683

tmp3212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2682, Nil)
}
__typedArg0 := let__2682
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4choicepoint_b, tmp3212)
}
__typedArg0 := symshen_4choicepoint_b
__typedArg1 := tmp3212
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3214 := Call(__e, PrimFunc(sym_8p), let__2679, tmp3213)


tmp3215 := Call(__e, PrimFunc(symshen_4comb), let__2683, tmp3214)


ifres3209 = tmp3215


} else {
tmp3216 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3209 = tmp3216


}

ifres3207 = ifres3209


} else {
tmp3218 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3207 = tmp3218


}

ifres3203 = ifres3207


}

let__2677 := ifres3203
_ = let__2677

tmp3222 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2677)


if True == tmp3222 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2677)
return
}


} else {
__e.Return(let__2667)
return
}


} else {
__e.Return(let__2660)
return
}


} else {
__e.Return(let__2650)
return
}


}, 1)

tmp3229 := Call(__e, ns2_1set, symshen_4_5rule_6, tmp3124)


_ = tmp3229

tmp3230 := MakeNative(func(__e *ControlFlow) {
V663 := __e.Get(1)
_ = V663
tmp3231 := Call(__e, PrimFunc(symshen_4_5pattern_6), V663)


let__2685 := tmp3231
_ = let__2685

tmp3244 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2685)


var ifres3232 Obj

if True == tmp3244 {
tmp3233 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3232 = tmp3233


} else {
tmp3234 := Call(__e, PrimFunc(symshen_4_5_1out), let__2685)


let__2686 := tmp3234
_ = let__2686

tmp3235 := Call(__e, PrimFunc(symshen_4in_1_6), let__2685)


let__2687 := tmp3235
_ = let__2687

tmp3236 := Call(__e, PrimFunc(symshen_4_5patterns_6), let__2687)


let__2688 := tmp3236
_ = let__2688

tmp3243 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2688)


var ifres3237 Obj

if True == tmp3243 {
tmp3238 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3237 = tmp3238


} else {
tmp3239 := Call(__e, PrimFunc(symshen_4_5_1out), let__2688)


let__2689 := tmp3239
_ = let__2689

tmp3240 := Call(__e, PrimFunc(symshen_4in_1_6), let__2688)


let__2690 := tmp3240
_ = let__2690

tmp3241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2686, let__2689)
}
__typedArg0 := let__2686
__typedArg1 := let__2689
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3242 := Call(__e, PrimFunc(symshen_4comb), let__2690, tmp3241)


ifres3237 = tmp3242


}

ifres3232 = ifres3237


}

let__2684 := ifres3232
_ = let__2684

tmp3254 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2684)


if True == tmp3254 {
tmp3245 := Call(__e, PrimFunc(sym_5e_6), V663)


let__2692 := tmp3245
_ = let__2692

tmp3250 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2692)


var ifres3246 Obj

if True == tmp3250 {
tmp3247 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3246 = tmp3247


} else {
tmp3248 := Call(__e, PrimFunc(symshen_4in_1_6), let__2692)


let__2693 := tmp3248
_ = let__2693

tmp3249 := Call(__e, PrimFunc(symshen_4comb), let__2693, Nil)


ifres3246 = tmp3249


}

let__2691 := ifres3246
_ = let__2691

tmp3252 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2691)


if True == tmp3252 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2691)
return
}


} else {
__e.Return(let__2684)
return
}


}, 1)

tmp3255 := Call(__e, ns2_1set, symshen_4_5patterns_6, tmp3230)


_ = tmp3255

tmp3256 := MakeNative(func(__e *ControlFlow) {
V674 := __e.Get(1)
_ = V674
tmp3277 := Call(__e, PrimFunc(symshen_4ccons_2), V674)


var ifres3257 Obj

if True == tmp3277 {
tmp3258 := Call(__e, PrimFunc(symhead), V674)


let__2695 := tmp3258
_ = let__2695

tmp3259 := Call(__e, PrimFunc(symtail), V674)


let__2696 := tmp3259
_ = let__2696

tmp3275 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2695, symvector)


var ifres3260 Obj

if True == tmp3275 {
tmp3261 := Call(__e, PrimFunc(symtail), let__2695)


let__2697 := tmp3261
_ = let__2697

tmp3273 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2697, MakeInteger(0))


var ifres3262 Obj

if True == tmp3273 {
tmp3263 := Call(__e, PrimFunc(symtail), let__2697)


let__2698 := tmp3263
_ = let__2698

tmp3264 := Call(__e, PrimFunc(sym_5end_6), let__2698)


let__2699 := tmp3264
_ = let__2699

tmp3271 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2699)


var ifres3265 Obj

if True == tmp3271 {
tmp3266 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3265 = tmp3266


} else {
tmp3267 := Call(__e, PrimFunc(symshen_4in_1_6), let__2699)


let__2700 := tmp3267
_ = let__2700

tmp3268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), Nil)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp3268)
}
__typedArg0 := symvector
__typedArg1 := tmp3268
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3270 := Call(__e, PrimFunc(symshen_4comb), let__2696, tmp3269)


ifres3265 = tmp3270


}

ifres3262 = ifres3265


} else {
tmp3272 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3262 = tmp3272


}

ifres3260 = ifres3262


} else {
tmp3274 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3260 = tmp3274


}

ifres3257 = ifres3260


} else {
tmp3276 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3257 = tmp3276


}

let__2694 := ifres3257
_ = let__2694

tmp3300 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2694)


if True == tmp3300 {
tmp3287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V674)
}
__typedArg0 := V674
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3278 Obj

if True == tmp3287 {
tmp3279 := Call(__e, PrimFunc(symhead), V674)


let__2702 := tmp3279
_ = let__2702

tmp3280 := Call(__e, PrimFunc(symtail), V674)


let__2703 := tmp3280
_ = let__2703

tmp3285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__2702)
}
__typedArg0 := let__2702
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3281 Obj

if True == tmp3285 {
tmp3282 := Call(__e, PrimFunc(symshen_4compound_1pattern), let__2702)


tmp3283 := Call(__e, PrimFunc(symshen_4comb), let__2703, tmp3282)


ifres3281 = tmp3283


} else {
tmp3284 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3281 = tmp3284


}

ifres3278 = ifres3281


} else {
tmp3286 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3278 = tmp3286


}

let__2701 := ifres3278
_ = let__2701

tmp3298 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2701)


if True == tmp3298 {
tmp3288 := Call(__e, PrimFunc(symshen_4_5simple_1pattern_6), V674)


let__2705 := tmp3288
_ = let__2705

tmp3294 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2705)


var ifres3289 Obj

if True == tmp3294 {
tmp3290 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3289 = tmp3290


} else {
tmp3291 := Call(__e, PrimFunc(symshen_4_5_1out), let__2705)


let__2706 := tmp3291
_ = let__2706

tmp3292 := Call(__e, PrimFunc(symshen_4in_1_6), let__2705)


let__2707 := tmp3292
_ = let__2707

tmp3293 := Call(__e, PrimFunc(symshen_4comb), let__2707, let__2706)


ifres3289 = tmp3293


}

let__2704 := ifres3289
_ = let__2704

tmp3296 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2704)


if True == tmp3296 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2704)
return
}


} else {
__e.Return(let__2701)
return
}


} else {
__e.Return(let__2694)
return
}


}, 1)

tmp3301 := Call(__e, ns2_1set, symshen_4_5pattern_6, tmp3256)


_ = tmp3301

tmp3302 := MakeNative(func(__e *ControlFlow) {
V689 := __e.Get(1)
_ = V689
tmp3311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V689)
}
__typedArg0 := V689
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3303 Obj

if True == tmp3311 {
tmp3304 := Call(__e, PrimFunc(symhead), V689)


let__2709 := tmp3304
_ = let__2709

tmp3305 := Call(__e, PrimFunc(symtail), V689)


let__2710 := tmp3305
_ = let__2710

tmp3309 := Call(__e, PrimFunc(symshen_4constructor_2), let__2709)


var ifres3306 Obj

if True == tmp3309 {
tmp3307 := Call(__e, PrimFunc(symshen_4comb), let__2710, let__2709)


ifres3306 = tmp3307


} else {
tmp3308 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3306 = tmp3308


}

ifres3303 = ifres3306


} else {
tmp3310 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3303 = tmp3310


}

let__2708 := ifres3303
_ = let__2708

tmp3313 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2708)


if True == tmp3313 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2708)
return
}


}, 1)

tmp3314 := Call(__e, ns2_1set, symshen_4_5constructor_6, tmp3302)


_ = tmp3314

tmp3315 := MakeNative(func(__e *ControlFlow) {
V693 := __e.Get(1)
_ = V693
tmp3316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, Nil)
}
__typedArg0 := sym_8v
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp3316)
}
__typedArg0 := sym_8s
__typedArg1 := tmp3316
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp3317)
}
__typedArg0 := sym_8p
__typedArg1 := tmp3317
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp3318)
}
__typedArg0 := symcons
__typedArg1 := tmp3318
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symelement_2), V693, tmp3319)
return


}, 1)

tmp3320 := Call(__e, ns2_1set, symshen_4constructor_2, tmp3315)


_ = tmp3320

tmp3321 := MakeNative(func(__e *ControlFlow) {
V694 := __e.Get(1)
_ = V694
tmp3322 := Call(__e, PrimFunc(symshen_4app), V694, MakeString(" is not a legitimate constructor\n"), symshen_4r)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp3322)
}
__typedArg0 := tmp3322
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

tmp3323 := Call(__e, ns2_1set, symshen_4constructor_1error, tmp3321)


_ = tmp3323

tmp3324 := MakeNative(func(__e *ControlFlow) {
V695 := __e.Get(1)
_ = V695
V696 := __e.Get(2)
_ = V696
tmp3325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dcustom_1pattern_1compiler_d)
}
__typedArg0 := symshen_4_dcustom_1pattern_1compiler_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__2711 := tmp3325
_ = let__2711

tmp3328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__2711, False)
}
__typedArg0 := let__2711
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp3328 {
__e.TailApply(PrimFunc(symthaw), V696)
return
} else {
tmp3326 := Call(__e, let__2711, V695)


__e.TailApply(tmp3326, V696)
return


}


}, 2)

tmp3329 := Call(__e, ns2_1set, symshen_4custom_1pattern_1compiler, tmp3324)


_ = tmp3329

tmp3330 := MakeNative(func(__e *ControlFlow) {
V698 := __e.Get(1)
_ = V698
tmp3331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dcustom_1pattern_1reducer_d)
}
__typedArg0 := symshen_4_dcustom_1pattern_1reducer_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__2712 := tmp3331
_ = let__2712

tmp3333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__2712, False)
}
__typedArg0 := let__2712
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp3333 {
__e.TailApply(PrimFunc(symfail))
return
} else {
__e.TailApply(let__2712, V698)
return
}


}, 1)

tmp3334 := Call(__e, ns2_1set, symshen_4custom_1pattern_1reducer, tmp3330)


_ = tmp3334

tmp3335 := MakeNative(func(__e *ControlFlow) {
V700 := __e.Get(1)
_ = V700
tmp3336 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4compound_1pattern_1h), V700)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4custom_1pattern_1compiler), V700, tmp3336)
return


}, 1)

tmp3337 := Call(__e, ns2_1set, symshen_4compound_1pattern, tmp3335)


_ = tmp3337

tmp3338 := MakeNative(func(__e *ControlFlow) {
V701 := __e.Get(1)
_ = V701
tmp3369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V701)
}
__typedArg0 := V701
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3350 Obj

if True == tmp3369 {
tmp3367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V701)
}
__typedArg0 := V701
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3367)
}
__typedArg0 := tmp3367
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3352 Obj

if True == tmp3368 {
tmp3364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V701)
}
__typedArg0 := V701
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3364)
}
__typedArg0 := tmp3364
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3365)
}
__typedArg0 := tmp3365
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3354 Obj

if True == tmp3366 {
tmp3360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V701)
}
__typedArg0 := V701
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3360)
}
__typedArg0 := tmp3360
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3361)
}
__typedArg0 := tmp3361
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3362)
}
__typedArg0 := Nil
__typedArg1 := tmp3362
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3356 Obj

if True == tmp3363 {
tmp3358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V701)
}
__typedArg0 := V701
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3359 := Call(__e, PrimFunc(symshen_4constructor_2), tmp3358)


var ifres3357 Obj

if True == tmp3359 {
ifres3357 = True


} else {
ifres3357 = False


}

ifres3356 = ifres3357


} else {
ifres3356 = False


}

var ifres3355 Obj

if True == ifres3356 {
ifres3355 = True


} else {
ifres3355 = False


}

ifres3354 = ifres3355


} else {
ifres3354 = False


}

var ifres3353 Obj

if True == ifres3354 {
ifres3353 = True


} else {
ifres3353 = False


}

ifres3352 = ifres3353


} else {
ifres3352 = False


}

var ifres3351 Obj

if True == ifres3352 {
ifres3351 = True


} else {
ifres3351 = False


}

ifres3350 = ifres3351


} else {
ifres3350 = False


}

if True == ifres3350 {
tmp3339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V701)
}
__typedArg0 := V701
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V701)
}
__typedArg0 := V701
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3340)
}
__typedArg0 := tmp3340
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3342 := Call(__e, PrimFunc(symshen_4compile_1pattern_1fragment), tmp3341)


tmp3343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V701)
}
__typedArg0 := V701
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3343)
}
__typedArg0 := tmp3343
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3344)
}
__typedArg0 := tmp3344
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3346 := Call(__e, PrimFunc(symshen_4compile_1pattern_1fragment), tmp3345)


tmp3347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3346, Nil)
}
__typedArg0 := tmp3346
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3342, tmp3347)
}
__typedArg0 := tmp3342
__typedArg1 := tmp3347
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3339, tmp3348)
}
__typedArg0 := tmp3339
__typedArg1 := tmp3348
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4constructor_1error), V701)
return
}


}, 1)

tmp3370 := Call(__e, ns2_1set, symshen_4compound_1pattern_1h, tmp3338)


_ = tmp3370

tmp3371 := MakeNative(func(__e *ControlFlow) {
V702 := __e.Get(1)
_ = V702
tmp3400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V702)
}
__typedArg0 := V702
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3382 Obj

if True == tmp3400 {
tmp3398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V702)
}
__typedArg0 := V702
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symvector, tmp3398)
}
__typedArg0 := symvector
__typedArg1 := tmp3398
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3384 Obj

if True == tmp3399 {
tmp3396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V702)
}
__typedArg0 := V702
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3396)
}
__typedArg0 := tmp3396
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3386 Obj

if True == tmp3397 {
tmp3393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V702)
}
__typedArg0 := V702
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3393)
}
__typedArg0 := tmp3393
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), tmp3394)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp3394
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3388 Obj

if True == tmp3395 {
tmp3390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V702)
}
__typedArg0 := V702
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3390)
}
__typedArg0 := tmp3390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3391)
}
__typedArg0 := Nil
__typedArg1 := tmp3391
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3389 Obj

if True == tmp3392 {
ifres3389 = True


} else {
ifres3389 = False


}

ifres3388 = ifres3389


} else {
ifres3388 = False


}

var ifres3387 Obj

if True == ifres3388 {
ifres3387 = True


} else {
ifres3387 = False


}

ifres3386 = ifres3387


} else {
ifres3386 = False


}

var ifres3385 Obj

if True == ifres3386 {
ifres3385 = True


} else {
ifres3385 = False


}

ifres3384 = ifres3385


} else {
ifres3384 = False


}

var ifres3383 Obj

if True == ifres3384 {
ifres3383 = True


} else {
ifres3383 = False


}

ifres3382 = ifres3383


} else {
ifres3382 = False


}

if True == ifres3382 {
__e.Return(V702)
return
} else {
tmp3380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V702)
}
__typedArg0 := V702
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3380 {
__e.TailApply(PrimFunc(symshen_4compound_1pattern), V702)
return
} else {
tmp3378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V702, sym__)
}
__typedArg0 := V702
__typedArg1 := sym__
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp3378 {
__e.TailApply(PrimFunc(symgensym), symY)
return
} else {
tmp3373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1, Nil)
}
__typedArg0 := sym_5_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp3373)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp3373
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3375 := Call(__e, PrimFunc(symelement_2), V702, tmp3374)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp3375)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp3375
return Call(__e, PrimFunc(symnot), __typedArg0)
})() {
__e.Return(V702)
return
} else {
__e.TailApply(PrimFunc(symshen_4constructor_1error), V702)
return
}


}


}


}


}, 1)

tmp3401 := Call(__e, ns2_1set, symshen_4compile_1pattern_1fragment, tmp3371)


_ = tmp3401

tmp3402 := MakeNative(func(__e *ControlFlow) {
V707 := __e.Get(1)
_ = V707
tmp3428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V707)
}
__typedArg0 := V707
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3404 Obj

if True == tmp3428 {
tmp3426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V707)
}
__typedArg0 := V707
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8p, tmp3426)
}
__typedArg0 := sym_8p
__typedArg1 := tmp3426
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3406 Obj

if True == tmp3427 {
tmp3424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V707)
}
__typedArg0 := V707
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3424)
}
__typedArg0 := tmp3424
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3408 Obj

if True == tmp3425 {
tmp3421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V707)
}
__typedArg0 := V707
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3421)
}
__typedArg0 := tmp3421
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4custom_1pattern, tmp3422)
}
__typedArg0 := symshen_4custom_1pattern
__typedArg1 := tmp3422
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3410 Obj

if True == tmp3423 {
tmp3418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V707)
}
__typedArg0 := V707
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3418)
}
__typedArg0 := tmp3418
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3419)
}
__typedArg0 := tmp3419
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3412 Obj

if True == tmp3420 {
tmp3414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V707)
}
__typedArg0 := V707
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3414)
}
__typedArg0 := tmp3414
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3415)
}
__typedArg0 := tmp3415
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3416)
}
__typedArg0 := Nil
__typedArg1 := tmp3416
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3413 Obj

if True == tmp3417 {
ifres3413 = True


} else {
ifres3413 = False


}

ifres3412 = ifres3413


} else {
ifres3412 = False


}

var ifres3411 Obj

if True == ifres3412 {
ifres3411 = True


} else {
ifres3411 = False


}

ifres3410 = ifres3411


} else {
ifres3410 = False


}

var ifres3409 Obj

if True == ifres3410 {
ifres3409 = True


} else {
ifres3409 = False


}

ifres3408 = ifres3409


} else {
ifres3408 = False


}

var ifres3407 Obj

if True == ifres3408 {
ifres3407 = True


} else {
ifres3407 = False


}

ifres3406 = ifres3407


} else {
ifres3406 = False


}

var ifres3405 Obj

if True == ifres3406 {
ifres3405 = True


} else {
ifres3405 = False


}

ifres3404 = ifres3405


} else {
ifres3404 = False


}

if True == ifres3404 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp3429 := Call(__e, ns2_1set, symshen_4custom_1pattern_2, tmp3402)


_ = tmp3429

tmp3430 := MakeNative(func(__e *ControlFlow) {
V710 := __e.Get(1)
_ = V710
tmp3458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V710)
}
__typedArg0 := V710
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3434 Obj

if True == tmp3458 {
tmp3456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V710)
}
__typedArg0 := V710
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8p, tmp3456)
}
__typedArg0 := sym_8p
__typedArg1 := tmp3456
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3436 Obj

if True == tmp3457 {
tmp3454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V710)
}
__typedArg0 := V710
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3454)
}
__typedArg0 := tmp3454
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3438 Obj

if True == tmp3455 {
tmp3451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V710)
}
__typedArg0 := V710
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3451)
}
__typedArg0 := tmp3451
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4custom_1pattern, tmp3452)
}
__typedArg0 := symshen_4custom_1pattern
__typedArg1 := tmp3452
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3440 Obj

if True == tmp3453 {
tmp3448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V710)
}
__typedArg0 := V710
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3448)
}
__typedArg0 := tmp3448
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3449)
}
__typedArg0 := tmp3449
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3442 Obj

if True == tmp3450 {
tmp3444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V710)
}
__typedArg0 := V710
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3444)
}
__typedArg0 := tmp3444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3445)
}
__typedArg0 := tmp3445
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3446)
}
__typedArg0 := Nil
__typedArg1 := tmp3446
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3443 Obj

if True == tmp3447 {
ifres3443 = True


} else {
ifres3443 = False


}

ifres3442 = ifres3443


} else {
ifres3442 = False


}

var ifres3441 Obj

if True == ifres3442 {
ifres3441 = True


} else {
ifres3441 = False


}

ifres3440 = ifres3441


} else {
ifres3440 = False


}

var ifres3439 Obj

if True == ifres3440 {
ifres3439 = True


} else {
ifres3439 = False


}

ifres3438 = ifres3439


} else {
ifres3438 = False


}

var ifres3437 Obj

if True == ifres3438 {
ifres3437 = True


} else {
ifres3437 = False


}

ifres3436 = ifres3437


} else {
ifres3436 = False


}

var ifres3435 Obj

if True == ifres3436 {
ifres3435 = True


} else {
ifres3435 = False


}

ifres3434 = ifres3435


} else {
ifres3434 = False


}

if True == ifres3434 {
tmp3431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V710)
}
__typedArg0 := V710
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3431)
}
__typedArg0 := tmp3431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3432)
}
__typedArg0 := tmp3432
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.custom-pattern-body"))
}
__typedArg0 := MakeString("implementation error in shen.custom-pattern-body")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 1)

tmp3459 := Call(__e, ns2_1set, symshen_4custom_1pattern_1body, tmp3430)


_ = tmp3459

tmp3460 := MakeNative(func(__e *ControlFlow) {
V711 := __e.Get(1)
_ = V711
tmp3470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V711)
}
__typedArg0 := V711
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3461 Obj

if True == tmp3470 {
tmp3462 := Call(__e, PrimFunc(symhead), V711)


let__2714 := tmp3462
_ = let__2714

tmp3463 := Call(__e, PrimFunc(symtail), V711)


let__2715 := tmp3463
_ = let__2715

tmp3468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__2714, sym__)
}
__typedArg0 := let__2714
__typedArg1 := sym__
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3464 Obj

if True == tmp3468 {
tmp3465 := Call(__e, PrimFunc(symgensym), symY)


tmp3466 := Call(__e, PrimFunc(symshen_4comb), let__2715, tmp3465)


ifres3464 = tmp3466


} else {
tmp3467 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3464 = tmp3467


}

ifres3461 = ifres3464


} else {
tmp3469 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3461 = tmp3469


}

let__2713 := ifres3461
_ = let__2713

tmp3486 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2713)


if True == tmp3486 {
tmp3482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V711)
}
__typedArg0 := V711
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3471 Obj

if True == tmp3482 {
tmp3472 := Call(__e, PrimFunc(symhead), V711)


let__2717 := tmp3472
_ = let__2717

tmp3473 := Call(__e, PrimFunc(symtail), V711)


let__2718 := tmp3473
_ = let__2718

tmp3477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1, Nil)
}
__typedArg0 := sym_5_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp3477)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp3477
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3479 := Call(__e, PrimFunc(symelement_2), let__2717, tmp3478)


tmp3480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp3479)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp3479
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres3474 Obj

if True == tmp3480 {
tmp3475 := Call(__e, PrimFunc(symshen_4comb), let__2718, let__2717)


ifres3474 = tmp3475


} else {
tmp3476 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3474 = tmp3476


}

ifres3471 = ifres3474


} else {
tmp3481 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3471 = tmp3481


}

let__2716 := ifres3471
_ = let__2716

tmp3484 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2716)


if True == tmp3484 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2716)
return
}


} else {
__e.Return(let__2713)
return
}


}, 1)

tmp3487 := Call(__e, ns2_1set, symshen_4_5simple_1pattern_6, tmp3460)


_ = tmp3487

tmp3488 := MakeNative(func(__e *ControlFlow) {
V718 := __e.Get(1)
_ = V718
tmp3489 := Call(__e, PrimFunc(symshen_4_5pattern_6), V718)


let__2720 := tmp3489
_ = let__2720

tmp3495 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2720)


var ifres3490 Obj

if True == tmp3495 {
tmp3491 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3490 = tmp3491


} else {
tmp3492 := Call(__e, PrimFunc(symshen_4_5_1out), let__2720)


let__2721 := tmp3492
_ = let__2721

tmp3493 := Call(__e, PrimFunc(symshen_4in_1_6), let__2720)


let__2722 := tmp3493
_ = let__2722

tmp3494 := Call(__e, PrimFunc(symshen_4comb), let__2722, let__2721)


ifres3490 = tmp3494


}

let__2719 := ifres3490
_ = let__2719

tmp3497 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2719)


if True == tmp3497 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2719)
return
}


}, 1)

tmp3498 := Call(__e, ns2_1set, symshen_4_5pattern1_6, tmp3488)


_ = tmp3498

tmp3499 := MakeNative(func(__e *ControlFlow) {
V723 := __e.Get(1)
_ = V723
tmp3500 := Call(__e, PrimFunc(symshen_4_5pattern_6), V723)


let__2724 := tmp3500
_ = let__2724

tmp3506 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2724)


var ifres3501 Obj

if True == tmp3506 {
tmp3502 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3501 = tmp3502


} else {
tmp3503 := Call(__e, PrimFunc(symshen_4_5_1out), let__2724)


let__2725 := tmp3503
_ = let__2725

tmp3504 := Call(__e, PrimFunc(symshen_4in_1_6), let__2724)


let__2726 := tmp3504
_ = let__2726

tmp3505 := Call(__e, PrimFunc(symshen_4comb), let__2726, let__2725)


ifres3501 = tmp3505


}

let__2723 := ifres3501
_ = let__2723

tmp3508 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2723)


if True == tmp3508 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2723)
return
}


}, 1)

tmp3509 := Call(__e, ns2_1set, symshen_4_5pattern2_6, tmp3499)


_ = tmp3509

tmp3510 := MakeNative(func(__e *ControlFlow) {
V728 := __e.Get(1)
_ = V728
tmp3511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(MakeInteger(2))
}
__typedArg0 := MakeInteger(2)
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

let__2727 := tmp3511
_ = let__2727

tmp3512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__2727, MakeInteger(0), symshen_4printF)
}
__typedArg0 := let__2727
__typedArg1 := MakeInteger(0)
__typedArg2 := symshen_4printF
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__2728 := tmp3512
_ = let__2728

tmp3513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V728)
}
__typedArg0 := V728
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp3514 := Call(__e, PrimFunc(sym_8s), tmp3513, MakeString(")"))


tmp3515 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp3514)


tmp3516 := Call(__e, PrimFunc(sym_8s), MakeString("n"), tmp3515)


tmp3517 := Call(__e, PrimFunc(sym_8s), MakeString("f"), tmp3516)


tmp3518 := Call(__e, PrimFunc(sym_8s), MakeString("("), tmp3517)


tmp3519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__2728, MakeInteger(1), tmp3518)
}
__typedArg0 := let__2728
__typedArg1 := MakeInteger(1)
__typedArg2 := tmp3518
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__2729 := tmp3519
_ = let__2729

__e.Return(let__2729)
return


}, 1)

tmp3520 := Call(__e, ns2_1set, symshen_4fn_1print, tmp3510)


_ = tmp3520

tmp3521 := MakeNative(func(__e *ControlFlow) {
V732 := __e.Get(1)
_ = V732
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V732, MakeInteger(1))
}
__typedArg0 := V732
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp3522 := Call(__e, ns2_1set, symshen_4printF, tmp3521)


_ = tmp3522

tmp3523 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V737 := __self1
_ = V737
V738 := __self2
_ = V738
tmp3547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3543 Obj

if True == tmp3547 {
tmp3545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3545)
}
__typedArg0 := Nil
__typedArg1 := tmp3545
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3544 Obj

if True == tmp3546 {
ifres3544 = True


} else {
ifres3544 = False


}

ifres3543 = ifres3544


} else {
ifres3543 = False


}

if True == ifres3543 {
tmp3524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symlength), tmp3524)
return


} else {
tmp3541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3529 Obj

if True == tmp3541 {
tmp3539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3539)
}
__typedArg0 := tmp3539
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3531 Obj

if True == tmp3540 {
tmp3533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3534 := Call(__e, PrimFunc(symlength), tmp3533)


tmp3535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3535)
}
__typedArg0 := tmp3535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3537 := Call(__e, PrimFunc(symlength), tmp3536)


tmp3538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp3534, tmp3537)
}
__typedArg0 := tmp3534
__typedArg1 := tmp3537
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3532 Obj

if True == tmp3538 {
ifres3532 = True


} else {
ifres3532 = False


}

ifres3531 = ifres3532


} else {
ifres3531 = False


}

var ifres3530 Obj

if True == ifres3531 {
ifres3530 = True


} else {
ifres3530 = False


}

ifres3529 = ifres3530


} else {
ifres3529 = False


}

if True == ifres3529 {
tmp3525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4arity_1chk) == __self {
__self1, __self2 = V737, tmp3525
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4arity_1chk), V737, tmp3525)
return


} else {
tmp3526 := Call(__e, PrimFunc(symshen_4app), V737, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("arity error in "))
__typedS1, __typedOK1 := TypedString(tmp3526)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("arity error in ")
__typedArg1 := tmp3526
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("arity error in "))
__typedS1, __typedOK1 := TypedString(tmp3526)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("arity error in ")
__typedArg1 := tmp3526
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}


}, 2)

tmp3548 := Call(__e, ns2_1set, symshen_4arity_1chk, tmp3523)


_ = tmp3548

tmp3549 := MakeNative(func(__e *ControlFlow) {
V739 := __e.Get(1)
_ = V739
V740 := __e.Get(2)
_ = V740
tmp3555 := Call(__e, PrimFunc(symtuple_2), V740)


if True == tmp3555 {
tmp3550 := Call(__e, PrimFunc(symfst), V740)


tmp3551 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp3550)


tmp3552 := Call(__e, PrimFunc(symsnd), V740)


tmp3553 := Call(__e, PrimFunc(symshen_4find_1free_1vars), tmp3551, tmp3552)


__e.TailApply(PrimFunc(symshen_4free_1variable_1error_1message), V739, tmp3553)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4free_1var_1chk)
return
}


}, 2)

tmp3556 := Call(__e, ns2_1set, symshen_4free_1var_1chk, tmp3549)


_ = tmp3556

tmp3557 := MakeNative(func(__e *ControlFlow) {
V741 := __e.Get(1)
_ = V741
V742 := __e.Get(2)
_ = V742
tmp3569 := Call(__e, PrimFunc(symempty_2), V742)


if True == tmp3569 {
__e.Return(symshen_4skip)
return
} else {
tmp3558 := Call(__e, PrimFunc(symshen_4app), V741, MakeString(":"), symshen_4a)


tmp3559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("free variables in "))
__typedS1, __typedOK1 := TypedString(tmp3558)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("free variables in ")
__typedArg1 := tmp3558
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp3560 := Call(__e, PrimFunc(symstoutput))


tmp3561 := Call(__e, PrimFunc(sympr), tmp3559, tmp3560)


_ = tmp3561

tmp3562 := MakeNative(func(__e *ControlFlow) {
Z743 := __e.Get(1)
_ = Z743
tmp3563 := Call(__e, PrimFunc(symshen_4app), Z743, MakeString(""), symshen_4a)


tmp3564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" "))
__typedS1, __typedOK1 := TypedString(tmp3563)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" ")
__typedArg1 := tmp3563
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp3565 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp3564, tmp3565)
return


}, 1)

tmp3566 := Call(__e, PrimFunc(symshen_4for_1each), tmp3562, V742)


_ = tmp3566

tmp3567 := Call(__e, PrimFunc(symnl), MakeInteger(1))


_ = tmp3567

__e.TailApply(PrimFunc(symabort))
return


}


}, 2)

tmp3570 := Call(__e, ns2_1set, symshen_4free_1variable_1error_1message, tmp3557)


_ = tmp3570

tmp3571 := MakeNative(func(__e *ControlFlow) {
V746 := __e.Get(1)
_ = V746
tmp3579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V746)
}
__typedArg0 := V746
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp3579 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V746, Nil)
}
__typedArg0 := V746
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
} else {
tmp3577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V746)
}
__typedArg0 := V746
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3577 {
tmp3572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V746)
}
__typedArg0 := V746
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3573 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp3572)


tmp3574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V746)
}
__typedArg0 := V746
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3575 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp3574)


__e.TailApply(PrimFunc(symunion), tmp3573, tmp3575)
return


} else {
__e.Return(Nil)
return
}


}


}, 1)

tmp3580 := Call(__e, ns2_1set, symshen_4extract_1vars, tmp3571)


_ = tmp3580

tmp3581 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V751 := __self1
_ = V751
V752 := __self2
_ = V752
tmp3671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3658 Obj

if True == tmp3671 {
tmp3669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symprotect, tmp3669)
}
__typedArg0 := symprotect
__typedArg1 := tmp3669
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3660 Obj

if True == tmp3670 {
tmp3667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3667)
}
__typedArg0 := tmp3667
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3662 Obj

if True == tmp3668 {
tmp3664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3664)
}
__typedArg0 := tmp3664
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3665)
}
__typedArg0 := Nil
__typedArg1 := tmp3665
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3663 Obj

if True == tmp3666 {
ifres3663 = True


} else {
ifres3663 = False


}

ifres3662 = ifres3663


} else {
ifres3662 = False


}

var ifres3661 Obj

if True == ifres3662 {
ifres3661 = True


} else {
ifres3661 = False


}

ifres3660 = ifres3661


} else {
ifres3660 = False


}

var ifres3659 Obj

if True == ifres3660 {
ifres3659 = True


} else {
ifres3659 = False


}

ifres3658 = ifres3659


} else {
ifres3658 = False


}

if True == ifres3658 {
__e.Return(Nil)
return
} else {
tmp3656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3630 Obj

if True == tmp3656 {
tmp3654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp3654)
}
__typedArg0 := symlet
__typedArg1 := tmp3654
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3632 Obj

if True == tmp3655 {
tmp3652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3652)
}
__typedArg0 := tmp3652
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3634 Obj

if True == tmp3653 {
tmp3649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3649)
}
__typedArg0 := tmp3649
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3650)
}
__typedArg0 := tmp3650
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3636 Obj

if True == tmp3651 {
tmp3645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3645)
}
__typedArg0 := tmp3645
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3646)
}
__typedArg0 := tmp3646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3647)
}
__typedArg0 := tmp3647
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3638 Obj

if True == tmp3648 {
tmp3640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3640)
}
__typedArg0 := tmp3640
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3641)
}
__typedArg0 := tmp3641
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3642)
}
__typedArg0 := tmp3642
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3643)
}
__typedArg0 := Nil
__typedArg1 := tmp3643
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3639 Obj

if True == tmp3644 {
ifres3639 = True


} else {
ifres3639 = False


}

ifres3638 = ifres3639


} else {
ifres3638 = False


}

var ifres3637 Obj

if True == ifres3638 {
ifres3637 = True


} else {
ifres3637 = False


}

ifres3636 = ifres3637


} else {
ifres3636 = False


}

var ifres3635 Obj

if True == ifres3636 {
ifres3635 = True


} else {
ifres3635 = False


}

ifres3634 = ifres3635


} else {
ifres3634 = False


}

var ifres3633 Obj

if True == ifres3634 {
ifres3633 = True


} else {
ifres3633 = False


}

ifres3632 = ifres3633


} else {
ifres3632 = False


}

var ifres3631 Obj

if True == ifres3632 {
ifres3631 = True


} else {
ifres3631 = False


}

ifres3630 = ifres3631


} else {
ifres3630 = False


}

if True == ifres3630 {
tmp3582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3582)
}
__typedArg0 := tmp3582
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3583)
}
__typedArg0 := tmp3583
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3585 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V751, tmp3584)


tmp3586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3586)
}
__typedArg0 := tmp3586
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3587, V751)
}
__typedArg0 := tmp3587
__typedArg1 := V751
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3589)
}
__typedArg0 := tmp3589
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3590)
}
__typedArg0 := tmp3590
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3591)
}
__typedArg0 := tmp3591
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3593 := Call(__e, PrimFunc(symshen_4find_1free_1vars), tmp3588, tmp3592)


__e.TailApply(PrimFunc(symunion), tmp3585, tmp3593)
return


} else {
tmp3628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3609 Obj

if True == tmp3628 {
tmp3626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp3626)
}
__typedArg0 := symlambda
__typedArg1 := tmp3626
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3611 Obj

if True == tmp3627 {
tmp3624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3624)
}
__typedArg0 := tmp3624
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3613 Obj

if True == tmp3625 {
tmp3621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3621)
}
__typedArg0 := tmp3621
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3622)
}
__typedArg0 := tmp3622
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3615 Obj

if True == tmp3623 {
tmp3617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3617)
}
__typedArg0 := tmp3617
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3618)
}
__typedArg0 := tmp3618
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3619)
}
__typedArg0 := Nil
__typedArg1 := tmp3619
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3616 Obj

if True == tmp3620 {
ifres3616 = True


} else {
ifres3616 = False


}

ifres3615 = ifres3616


} else {
ifres3615 = False


}

var ifres3614 Obj

if True == ifres3615 {
ifres3614 = True


} else {
ifres3614 = False


}

ifres3613 = ifres3614


} else {
ifres3613 = False


}

var ifres3612 Obj

if True == ifres3613 {
ifres3612 = True


} else {
ifres3612 = False


}

ifres3611 = ifres3612


} else {
ifres3611 = False


}

var ifres3610 Obj

if True == ifres3611 {
ifres3610 = True


} else {
ifres3610 = False


}

ifres3609 = ifres3610


} else {
ifres3609 = False


}

if True == ifres3609 {
tmp3594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3594)
}
__typedArg0 := tmp3594
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3595, V751)
}
__typedArg0 := tmp3595
__typedArg1 := V751
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3597)
}
__typedArg0 := tmp3597
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3598)
}
__typedArg0 := tmp3598
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

if PrimFunc(symshen_4find_1free_1vars) == __self {
__self1, __self2 = tmp3596, tmp3599
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4find_1free_1vars), tmp3596, tmp3599)
return


} else {
tmp3607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3607 {
tmp3600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3601 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V751, tmp3600)


tmp3602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3603 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V751, tmp3602)


__e.TailApply(PrimFunc(symunion), tmp3601, tmp3603)
return


} else {
tmp3605 := Call(__e, PrimFunc(symshen_4free_1variable_2), V752, V751)


if True == tmp3605 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V752, Nil)
}
__typedArg0 := V752
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
} else {
__e.Return(Nil)
return
}


}


}


}


}


}, 2)

tmp3672 := Call(__e, ns2_1set, symshen_4find_1free_1vars, tmp3581)


_ = tmp3672

tmp3673 := MakeNative(func(__e *ControlFlow) {
V753 := __e.Get(1)
_ = V753
V754 := __e.Get(2)
_ = V754
tmp3678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V753)
}
__typedArg0 := V753
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp3678 {
tmp3675 := Call(__e, PrimFunc(symelement_2), V753, V754)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp3675)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp3675
return Call(__e, PrimFunc(symnot), __typedArg0)
})() {
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


}, 2)

tmp3679 := Call(__e, ns2_1set, symshen_4free_1variable_2, tmp3673)


_ = tmp3679

tmp3680 := MakeNative(func(__e *ControlFlow) {
V755 := __e.Get(1)
_ = V755
V756 := __e.Get(2)
_ = V756
tmp3681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_duserdefs_d)
}
__typedArg0 := symshen_4_duserdefs_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp3682 := Call(__e, PrimFunc(symadjoin), V755, tmp3681)


tmp3683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_duserdefs_d, tmp3682)
}
__typedArg0 := symshen_4_duserdefs_d
__typedArg1 := tmp3682
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp3683

tmp3684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V755, symshen_4source, V756, tmp3684)
return


}, 2)

tmp3685 := Call(__e, ns2_1set, symshen_4record_1kl, tmp3680)


_ = tmp3685

tmp3686 := MakeNative(func(__e *ControlFlow) {
V757 := __e.Get(1)
_ = V757
V758 := __e.Get(2)
_ = V758
V759 := __e.Get(3)
_ = V759
tmp3687 := Call(__e, PrimFunc(symshen_4parameters), V759)


let__2730 := tmp3687
_ = let__2730

tmp3688 := Call(__e, PrimFunc(symshen_4kl_1body), V758, let__2730)


tmp3689 := Call(__e, PrimFunc(symshen_4scan_1body), V757, tmp3688)


let__2731 := tmp3689
_ = let__2731

tmp3690 := Call(__e, PrimFunc(symshen_4cond_1form), let__2731)


tmp3691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3690, Nil)
}
__typedArg0 := tmp3690
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2730, tmp3691)
}
__typedArg0 := let__2730
__typedArg1 := tmp3691
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V757, tmp3692)
}
__typedArg0 := V757
__typedArg1 := tmp3692
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp3693)
}
__typedArg0 := symdefun
__typedArg1 := tmp3693
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__2732 := tmp3694
_ = let__2732

__e.Return(let__2732)
return


}, 3)

tmp3695 := Call(__e, ns2_1set, symshen_4compile_1to_1kl, tmp3686)


_ = tmp3695

tmp3696 := MakeNative(func(__e *ControlFlow) {
V763 := __e.Get(1)
_ = V763
tmp3701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V763)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V763
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp3701 {
__e.Return(Nil)
return
} else {
tmp3697 := Call(__e, PrimFunc(symgensym), symV)


tmp3699 := Call(__e, PrimFunc(symshen_4parameters), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V763)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V763
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3697, tmp3699)
}
__typedArg0 := tmp3697
__typedArg1 := tmp3699
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp3702 := Call(__e, ns2_1set, symshen_4parameters, tmp3696)


_ = tmp3702

tmp3703 := MakeNative(func(__e *ControlFlow) {
V766 := __e.Get(1)
_ = V766
tmp3727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V766)
}
__typedArg0 := V766
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3707 Obj

if True == tmp3727 {
tmp3725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V766)
}
__typedArg0 := V766
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3725)
}
__typedArg0 := tmp3725
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3709 Obj

if True == tmp3726 {
tmp3722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V766)
}
__typedArg0 := V766
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3722)
}
__typedArg0 := tmp3722
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, tmp3723)
}
__typedArg0 := True
__typedArg1 := tmp3723
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3711 Obj

if True == tmp3724 {
tmp3719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V766)
}
__typedArg0 := V766
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3719)
}
__typedArg0 := tmp3719
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3720)
}
__typedArg0 := tmp3720
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3713 Obj

if True == tmp3721 {
tmp3715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V766)
}
__typedArg0 := V766
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3715)
}
__typedArg0 := tmp3715
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3716)
}
__typedArg0 := tmp3716
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3717)
}
__typedArg0 := Nil
__typedArg1 := tmp3717
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3714 Obj

if True == tmp3718 {
ifres3714 = True


} else {
ifres3714 = False


}

ifres3713 = ifres3714


} else {
ifres3713 = False


}

var ifres3712 Obj

if True == ifres3713 {
ifres3712 = True


} else {
ifres3712 = False


}

ifres3711 = ifres3712


} else {
ifres3711 = False


}

var ifres3710 Obj

if True == ifres3711 {
ifres3710 = True


} else {
ifres3710 = False


}

ifres3709 = ifres3710


} else {
ifres3709 = False


}

var ifres3708 Obj

if True == ifres3709 {
ifres3708 = True


} else {
ifres3708 = False


}

ifres3707 = ifres3708


} else {
ifres3707 = False


}

if True == ifres3707 {
tmp3704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V766)
}
__typedArg0 := V766
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3704)
}
__typedArg0 := tmp3704
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3705)
}
__typedArg0 := tmp3705
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcond, V766)
}
__typedArg0 := symcond
__typedArg1 := V766
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
}


}, 1)

tmp3728 := Call(__e, ns2_1set, symshen_4cond_1form, tmp3703)


_ = tmp3728

tmp3729 := MakeNative(func(__e *ControlFlow) {
V775 := __e.Get(1)
_ = V775
V776 := __e.Get(2)
_ = V776
tmp3773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V776)
}
__typedArg0 := Nil
__typedArg1 := V776
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp3773 {
tmp3730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V775, Nil)
}
__typedArg0 := V775
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4f_1error, tmp3730)
}
__typedArg0 := symshen_4f_1error
__typedArg1 := tmp3730
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3731, Nil)
}
__typedArg0 := tmp3731
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp3732)
}
__typedArg0 := True
__typedArg1 := tmp3732
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3733, Nil)
}
__typedArg0 := tmp3733
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp3771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3767 Obj

if True == tmp3771 {
tmp3769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3770 := Call(__e, PrimFunc(symshen_4choicepoint_2), tmp3769)


var ifres3768 Obj

if True == tmp3770 {
ifres3768 = True


} else {
ifres3768 = False


}

ifres3767 = ifres3768


} else {
ifres3767 = False


}

if True == ifres3767 {
tmp3734 := Call(__e, PrimFunc(symgensym), symFreeze)


tmp3735 := Call(__e, PrimFunc(symgensym), symResult)


tmp3736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4choicepoint), V775, tmp3734, tmp3735, tmp3736, tmp3737)
return


} else {
tmp3765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3745 Obj

if True == tmp3765 {
tmp3763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3763)
}
__typedArg0 := tmp3763
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3747 Obj

if True == tmp3764 {
tmp3760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3760)
}
__typedArg0 := tmp3760
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, tmp3761)
}
__typedArg0 := True
__typedArg1 := tmp3761
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3749 Obj

if True == tmp3762 {
tmp3757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3757)
}
__typedArg0 := tmp3757
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3758)
}
__typedArg0 := tmp3758
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3751 Obj

if True == tmp3759 {
tmp3753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3753)
}
__typedArg0 := tmp3753
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3754)
}
__typedArg0 := tmp3754
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3755)
}
__typedArg0 := Nil
__typedArg1 := tmp3755
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3752 Obj

if True == tmp3756 {
ifres3752 = True


} else {
ifres3752 = False


}

ifres3751 = ifres3752


} else {
ifres3751 = False


}

var ifres3750 Obj

if True == ifres3751 {
ifres3750 = True


} else {
ifres3750 = False


}

ifres3749 = ifres3750


} else {
ifres3749 = False


}

var ifres3748 Obj

if True == ifres3749 {
ifres3748 = True


} else {
ifres3748 = False


}

ifres3747 = ifres3748


} else {
ifres3747 = False


}

var ifres3746 Obj

if True == ifres3747 {
ifres3746 = True


} else {
ifres3746 = False


}

ifres3745 = ifres3746


} else {
ifres3745 = False


}

if True == ifres3745 {
tmp3738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3738, Nil)
}
__typedArg0 := tmp3738
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp3743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3743 {
tmp3739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3741 := Call(__e, PrimFunc(symshen_4scan_1body), V775, tmp3740)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3739, tmp3741)
}
__typedArg0 := tmp3739
__typedArg1 := tmp3741
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.scan-body"))
}
__typedArg0 := MakeString("implementation error in shen.scan-body")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}


}, 2)

tmp3774 := Call(__e, ns2_1set, symshen_4scan_1body, tmp3729)


_ = tmp3774

tmp3775 := MakeNative(func(__e *ControlFlow) {
V783 := __e.Get(1)
_ = V783
tmp3810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V783)
}
__typedArg0 := V783
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3777 Obj

if True == tmp3810 {
tmp3808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V783)
}
__typedArg0 := V783
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3808)
}
__typedArg0 := tmp3808
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3779 Obj

if True == tmp3809 {
tmp3805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V783)
}
__typedArg0 := V783
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3805)
}
__typedArg0 := tmp3805
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3806)
}
__typedArg0 := tmp3806
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3781 Obj

if True == tmp3807 {
tmp3801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V783)
}
__typedArg0 := V783
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3801)
}
__typedArg0 := tmp3801
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3802)
}
__typedArg0 := tmp3802
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4choicepoint_b, tmp3803)
}
__typedArg0 := symshen_4choicepoint_b
__typedArg1 := tmp3803
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3783 Obj

if True == tmp3804 {
tmp3797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V783)
}
__typedArg0 := V783
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3797)
}
__typedArg0 := tmp3797
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3798)
}
__typedArg0 := tmp3798
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3799)
}
__typedArg0 := tmp3799
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3785 Obj

if True == tmp3800 {
tmp3792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V783)
}
__typedArg0 := V783
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3792)
}
__typedArg0 := tmp3792
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3793)
}
__typedArg0 := tmp3793
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3794)
}
__typedArg0 := tmp3794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3795)
}
__typedArg0 := Nil
__typedArg1 := tmp3795
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3787 Obj

if True == tmp3796 {
tmp3789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V783)
}
__typedArg0 := V783
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3789)
}
__typedArg0 := tmp3789
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3790)
}
__typedArg0 := Nil
__typedArg1 := tmp3790
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3788 Obj

if True == tmp3791 {
ifres3788 = True


} else {
ifres3788 = False


}

ifres3787 = ifres3788


} else {
ifres3787 = False


}

var ifres3786 Obj

if True == ifres3787 {
ifres3786 = True


} else {
ifres3786 = False


}

ifres3785 = ifres3786


} else {
ifres3785 = False


}

var ifres3784 Obj

if True == ifres3785 {
ifres3784 = True


} else {
ifres3784 = False


}

ifres3783 = ifres3784


} else {
ifres3783 = False


}

var ifres3782 Obj

if True == ifres3783 {
ifres3782 = True


} else {
ifres3782 = False


}

ifres3781 = ifres3782


} else {
ifres3781 = False


}

var ifres3780 Obj

if True == ifres3781 {
ifres3780 = True


} else {
ifres3780 = False


}

ifres3779 = ifres3780


} else {
ifres3779 = False


}

var ifres3778 Obj

if True == ifres3779 {
ifres3778 = True


} else {
ifres3778 = False


}

ifres3777 = ifres3778


} else {
ifres3777 = False


}

if True == ifres3777 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp3811 := Call(__e, ns2_1set, symshen_4choicepoint_2, tmp3775)


_ = tmp3811

tmp3812 := MakeNative(func(__e *ControlFlow) {
V799 := __e.Get(1)
_ = V799
V800 := __e.Get(2)
_ = V800
V801 := __e.Get(3)
_ = V801
V802 := __e.Get(4)
_ = V802
V803 := __e.Get(5)
_ = V803
tmp4004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3926 Obj

if True == tmp4004 {
tmp4002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4002)
}
__typedArg0 := tmp4002
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3928 Obj

if True == tmp4003 {
tmp3999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3999)
}
__typedArg0 := tmp3999
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4000)
}
__typedArg0 := tmp4000
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3930 Obj

if True == tmp4001 {
tmp3995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3995)
}
__typedArg0 := tmp3995
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3996)
}
__typedArg0 := tmp3996
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3997)
}
__typedArg0 := tmp3997
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3932 Obj

if True == tmp3998 {
tmp3990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3990)
}
__typedArg0 := tmp3990
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3991)
}
__typedArg0 := tmp3991
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3992)
}
__typedArg0 := tmp3992
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3993)
}
__typedArg0 := tmp3993
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3934 Obj

if True == tmp3994 {
tmp3984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3984)
}
__typedArg0 := tmp3984
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3985)
}
__typedArg0 := tmp3985
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3986)
}
__typedArg0 := tmp3986
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3987)
}
__typedArg0 := tmp3987
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail_1if, tmp3988)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp3988
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3936 Obj

if True == tmp3989 {
tmp3978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3978)
}
__typedArg0 := tmp3978
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3979)
}
__typedArg0 := tmp3979
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3980)
}
__typedArg0 := tmp3980
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3981)
}
__typedArg0 := tmp3981
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3982)
}
__typedArg0 := tmp3982
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3938 Obj

if True == tmp3983 {
tmp3971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3971)
}
__typedArg0 := tmp3971
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3972)
}
__typedArg0 := tmp3972
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3973)
}
__typedArg0 := tmp3973
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3974)
}
__typedArg0 := tmp3974
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3975)
}
__typedArg0 := tmp3975
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3976)
}
__typedArg0 := tmp3976
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3940 Obj

if True == tmp3977 {
tmp3963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3963)
}
__typedArg0 := tmp3963
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3964)
}
__typedArg0 := tmp3964
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3965)
}
__typedArg0 := tmp3965
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3966)
}
__typedArg0 := tmp3966
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3967)
}
__typedArg0 := tmp3967
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3968)
}
__typedArg0 := tmp3968
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3969)
}
__typedArg0 := Nil
__typedArg1 := tmp3969
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3942 Obj

if True == tmp3970 {
tmp3958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3958)
}
__typedArg0 := tmp3958
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3959)
}
__typedArg0 := tmp3959
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3960)
}
__typedArg0 := tmp3960
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3961)
}
__typedArg0 := Nil
__typedArg1 := tmp3961
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3944 Obj

if True == tmp3962 {
tmp3955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3955)
}
__typedArg0 := tmp3955
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3956)
}
__typedArg0 := Nil
__typedArg1 := tmp3956
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3946 Obj

if True == tmp3957 {
tmp3948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3948)
}
__typedArg0 := tmp3948
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3949)
}
__typedArg0 := tmp3949
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3950)
}
__typedArg0 := tmp3950
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3951)
}
__typedArg0 := tmp3951
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3952)
}
__typedArg0 := tmp3952
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V799, tmp3953)
}
__typedArg0 := V799
__typedArg1 := tmp3953
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3947 Obj

if True == tmp3954 {
ifres3947 = True


} else {
ifres3947 = False


}

ifres3946 = ifres3947


} else {
ifres3946 = False


}

var ifres3945 Obj

if True == ifres3946 {
ifres3945 = True


} else {
ifres3945 = False


}

ifres3944 = ifres3945


} else {
ifres3944 = False


}

var ifres3943 Obj

if True == ifres3944 {
ifres3943 = True


} else {
ifres3943 = False


}

ifres3942 = ifres3943


} else {
ifres3942 = False


}

var ifres3941 Obj

if True == ifres3942 {
ifres3941 = True


} else {
ifres3941 = False


}

ifres3940 = ifres3941


} else {
ifres3940 = False


}

var ifres3939 Obj

if True == ifres3940 {
ifres3939 = True


} else {
ifres3939 = False


}

ifres3938 = ifres3939


} else {
ifres3938 = False


}

var ifres3937 Obj

if True == ifres3938 {
ifres3937 = True


} else {
ifres3937 = False


}

ifres3936 = ifres3937


} else {
ifres3936 = False


}

var ifres3935 Obj

if True == ifres3936 {
ifres3935 = True


} else {
ifres3935 = False


}

ifres3934 = ifres3935


} else {
ifres3934 = False


}

var ifres3933 Obj

if True == ifres3934 {
ifres3933 = True


} else {
ifres3933 = False


}

ifres3932 = ifres3933


} else {
ifres3932 = False


}

var ifres3931 Obj

if True == ifres3932 {
ifres3931 = True


} else {
ifres3931 = False


}

ifres3930 = ifres3931


} else {
ifres3930 = False


}

var ifres3929 Obj

if True == ifres3930 {
ifres3929 = True


} else {
ifres3929 = False


}

ifres3928 = ifres3929


} else {
ifres3928 = False


}

var ifres3927 Obj

if True == ifres3928 {
ifres3927 = True


} else {
ifres3927 = False


}

ifres3926 = ifres3927


} else {
ifres3926 = False


}

if True == ifres3926 {
tmp3813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3813)
}
__typedArg0 := tmp3813
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3814)
}
__typedArg0 := tmp3814
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3815)
}
__typedArg0 := tmp3815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3816)
}
__typedArg0 := tmp3816
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3817)
}
__typedArg0 := tmp3817
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3819 := Call(__e, PrimFunc(symshen_4scan_1body), tmp3818, V803)


tmp3820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcond, tmp3819)
}
__typedArg0 := symcond
__typedArg1 := tmp3819
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3820, Nil)
}
__typedArg0 := tmp3820
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp3821)
}
__typedArg0 := symfreeze
__typedArg1 := tmp3821
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3824)
}
__typedArg0 := tmp3824
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3825)
}
__typedArg0 := tmp3825
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3826)
}
__typedArg0 := tmp3826
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3827)
}
__typedArg0 := tmp3827
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3828)
}
__typedArg0 := tmp3828
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3829)
}
__typedArg0 := tmp3829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3831)
}
__typedArg0 := tmp3831
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3832)
}
__typedArg0 := tmp3832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3833)
}
__typedArg0 := tmp3833
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3834)
}
__typedArg0 := tmp3834
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3835)
}
__typedArg0 := tmp3835
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V801, Nil)
}
__typedArg0 := V801
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3836, tmp3837)
}
__typedArg0 := tmp3836
__typedArg1 := tmp3837
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V800, Nil)
}
__typedArg0 := V800
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp3839)
}
__typedArg0 := symthaw
__typedArg1 := tmp3839
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V801, Nil)
}
__typedArg0 := V801
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3840, tmp3841)
}
__typedArg0 := tmp3840
__typedArg1 := tmp3841
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3838, tmp3842)
}
__typedArg0 := tmp3838
__typedArg1 := tmp3842
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp3843)
}
__typedArg0 := symif
__typedArg1 := tmp3843
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3844, Nil)
}
__typedArg0 := tmp3844
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3830, tmp3845)
}
__typedArg0 := tmp3830
__typedArg1 := tmp3845
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V801, tmp3846)
}
__typedArg0 := V801
__typedArg1 := tmp3846
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp3847)
}
__typedArg0 := symlet
__typedArg1 := tmp3847
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V800, Nil)
}
__typedArg0 := V800
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp3849)
}
__typedArg0 := symthaw
__typedArg1 := tmp3849
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3850, Nil)
}
__typedArg0 := tmp3850
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3848, tmp3851)
}
__typedArg0 := tmp3848
__typedArg1 := tmp3851
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3823, tmp3852)
}
__typedArg0 := tmp3823
__typedArg1 := tmp3852
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp3853)
}
__typedArg0 := symif
__typedArg1 := tmp3853
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3854, Nil)
}
__typedArg0 := tmp3854
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3822, tmp3855)
}
__typedArg0 := tmp3822
__typedArg1 := tmp3855
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V800, tmp3856)
}
__typedArg0 := V800
__typedArg1 := tmp3856
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp3857)
}
__typedArg0 := symlet
__typedArg1 := tmp3857
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3858, Nil)
}
__typedArg0 := tmp3858
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp3859)
}
__typedArg0 := True
__typedArg1 := tmp3859
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3860, Nil)
}
__typedArg0 := tmp3860
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp3924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3897 Obj

if True == tmp3924 {
tmp3922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3922)
}
__typedArg0 := tmp3922
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3899 Obj

if True == tmp3923 {
tmp3919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3919)
}
__typedArg0 := tmp3919
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3920)
}
__typedArg0 := tmp3920
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3901 Obj

if True == tmp3921 {
tmp3915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3915)
}
__typedArg0 := tmp3915
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3916)
}
__typedArg0 := tmp3916
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3917)
}
__typedArg0 := tmp3917
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3903 Obj

if True == tmp3918 {
tmp3910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3910)
}
__typedArg0 := tmp3910
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3911)
}
__typedArg0 := tmp3911
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3912)
}
__typedArg0 := tmp3912
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3913)
}
__typedArg0 := Nil
__typedArg1 := tmp3913
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3905 Obj

if True == tmp3914 {
tmp3907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3907)
}
__typedArg0 := tmp3907
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3908)
}
__typedArg0 := Nil
__typedArg1 := tmp3908
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3906 Obj

if True == tmp3909 {
ifres3906 = True


} else {
ifres3906 = False


}

ifres3905 = ifres3906


} else {
ifres3905 = False


}

var ifres3904 Obj

if True == ifres3905 {
ifres3904 = True


} else {
ifres3904 = False


}

ifres3903 = ifres3904


} else {
ifres3903 = False


}

var ifres3902 Obj

if True == ifres3903 {
ifres3902 = True


} else {
ifres3902 = False


}

ifres3901 = ifres3902


} else {
ifres3901 = False


}

var ifres3900 Obj

if True == ifres3901 {
ifres3900 = True


} else {
ifres3900 = False


}

ifres3899 = ifres3900


} else {
ifres3899 = False


}

var ifres3898 Obj

if True == ifres3899 {
ifres3898 = True


} else {
ifres3898 = False


}

ifres3897 = ifres3898


} else {
ifres3897 = False


}

if True == ifres3897 {
tmp3861 := Call(__e, PrimFunc(symshen_4scan_1body), V799, V803)


tmp3862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcond, tmp3861)
}
__typedArg0 := symcond
__typedArg1 := tmp3861
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3862, Nil)
}
__typedArg0 := tmp3862
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp3863)
}
__typedArg0 := symfreeze
__typedArg1 := tmp3863
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3866)
}
__typedArg0 := tmp3866
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3867)
}
__typedArg0 := tmp3867
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3868)
}
__typedArg0 := tmp3868
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, Nil)
}
__typedArg0 := symfail
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3870, Nil)
}
__typedArg0 := tmp3870
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V801, tmp3871)
}
__typedArg0 := V801
__typedArg1 := tmp3871
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp3872)
}
__typedArg0 := sym_a
__typedArg1 := tmp3872
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V800, Nil)
}
__typedArg0 := V800
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp3874)
}
__typedArg0 := symthaw
__typedArg1 := tmp3874
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V801, Nil)
}
__typedArg0 := V801
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3875, tmp3876)
}
__typedArg0 := tmp3875
__typedArg1 := tmp3876
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3873, tmp3877)
}
__typedArg0 := tmp3873
__typedArg1 := tmp3877
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp3878)
}
__typedArg0 := symif
__typedArg1 := tmp3878
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3879, Nil)
}
__typedArg0 := tmp3879
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3869, tmp3880)
}
__typedArg0 := tmp3869
__typedArg1 := tmp3880
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V801, tmp3881)
}
__typedArg0 := V801
__typedArg1 := tmp3881
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp3882)
}
__typedArg0 := symlet
__typedArg1 := tmp3882
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V800, Nil)
}
__typedArg0 := V800
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp3884)
}
__typedArg0 := symthaw
__typedArg1 := tmp3884
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3885, Nil)
}
__typedArg0 := tmp3885
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3883, tmp3886)
}
__typedArg0 := tmp3883
__typedArg1 := tmp3886
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3865, tmp3887)
}
__typedArg0 := tmp3865
__typedArg1 := tmp3887
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp3888)
}
__typedArg0 := symif
__typedArg1 := tmp3888
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3889, Nil)
}
__typedArg0 := tmp3889
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3864, tmp3890)
}
__typedArg0 := tmp3864
__typedArg1 := tmp3890
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V800, tmp3891)
}
__typedArg0 := V800
__typedArg1 := tmp3891
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp3892)
}
__typedArg0 := symlet
__typedArg1 := tmp3892
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3893, Nil)
}
__typedArg0 := tmp3893
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp3894)
}
__typedArg0 := True
__typedArg1 := tmp3894
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3895, Nil)
}
__typedArg0 := tmp3895
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.choicepoint"))
}
__typedArg0 := MakeString("implementation error in shen.choicepoint")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 5)

tmp4005 := Call(__e, ns2_1set, symshen_4choicepoint, tmp3812)


_ = tmp4005

tmp4006 := MakeNative(func(__e *ControlFlow) {
V805 := __e.Get(1)
_ = V805
V806 := __e.Get(2)
_ = V806
V807 := __e.Get(3)
_ = V807
tmp4019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V805, V807)
}
__typedArg0 := V805
__typedArg1 := V807
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4019 {
__e.Return(V806)
return
} else {
tmp4017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V807)
}
__typedArg0 := V807
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp4017 {
tmp4007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V807)
}
__typedArg0 := V807
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4008 := Call(__e, PrimFunc(symshen_4rep_1X), V805, V806, tmp4007)


let__2733 := tmp4008
_ = let__2733

tmp4014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V807)
}
__typedArg0 := V807
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__2733, tmp4014)
}
__typedArg0 := let__2733
__typedArg1 := tmp4014
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4015 {
tmp4009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V807)
}
__typedArg0 := V807
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V807)
}
__typedArg0 := V807
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4011 := Call(__e, PrimFunc(symshen_4rep_1X), V805, V806, tmp4010)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4009, tmp4011)
}
__typedArg0 := tmp4009
__typedArg1 := tmp4011
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp4012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V807)
}
__typedArg0 := V807
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2733, tmp4012)
}
__typedArg0 := let__2733
__typedArg1 := tmp4012
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
__e.Return(V807)
return
}


}


}, 3)

tmp4020 := Call(__e, ns2_1set, symshen_4rep_1X, tmp4006)


_ = tmp4020

tmp4021 := MakeNative(func(__e *ControlFlow) {
V809 := __e.Get(1)
_ = V809
tmp4100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4081 Obj

if True == tmp4100 {
tmp4098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp4098)
}
__typedArg0 := symlambda
__typedArg1 := tmp4098
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4083 Obj

if True == tmp4099 {
tmp4096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4096)
}
__typedArg0 := tmp4096
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4085 Obj

if True == tmp4097 {
tmp4093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4093)
}
__typedArg0 := tmp4093
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4094)
}
__typedArg0 := tmp4094
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4087 Obj

if True == tmp4095 {
tmp4089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4089)
}
__typedArg0 := tmp4089
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4090)
}
__typedArg0 := tmp4090
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4091)
}
__typedArg0 := Nil
__typedArg1 := tmp4091
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4088 Obj

if True == tmp4092 {
ifres4088 = True


} else {
ifres4088 = False


}

ifres4087 = ifres4088


} else {
ifres4087 = False


}

var ifres4086 Obj

if True == ifres4087 {
ifres4086 = True


} else {
ifres4086 = False


}

ifres4085 = ifres4086


} else {
ifres4085 = False


}

var ifres4084 Obj

if True == ifres4085 {
ifres4084 = True


} else {
ifres4084 = False


}

ifres4083 = ifres4084


} else {
ifres4083 = False


}

var ifres4082 Obj

if True == ifres4083 {
ifres4082 = True


} else {
ifres4082 = False


}

ifres4081 = ifres4082


} else {
ifres4081 = False


}

if True == ifres4081 {
tmp4022 := Call(__e, PrimFunc(symgensym), symZ)


let__2734 := tmp4022
_ = let__2734

tmp4023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4023)
}
__typedArg0 := tmp4023
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4025)
}
__typedArg0 := tmp4025
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4026)
}
__typedArg0 := tmp4026
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4028 := Call(__e, PrimFunc(symshen_4beta), tmp4024, let__2734, tmp4027)


tmp4029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4028, Nil)
}
__typedArg0 := tmp4028
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2734, tmp4029)
}
__typedArg0 := let__2734
__typedArg1 := tmp4029
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp4030)
}
__typedArg0 := symlambda
__typedArg1 := tmp4030
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__2735 := tmp4031
_ = let__2735

tmp4032 := MakeNative(func(__e *ControlFlow) {
Z812 := __e.Get(1)
_ = Z812
__e.TailApply(PrimFunc(symshen_4alpha_1convert), Z812)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp4032, let__2735)
return


} else {
tmp4079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4053 Obj

if True == tmp4079 {
tmp4077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp4077)
}
__typedArg0 := symlet
__typedArg1 := tmp4077
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4055 Obj

if True == tmp4078 {
tmp4075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4075)
}
__typedArg0 := tmp4075
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4057 Obj

if True == tmp4076 {
tmp4072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4072)
}
__typedArg0 := tmp4072
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4073)
}
__typedArg0 := tmp4073
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4059 Obj

if True == tmp4074 {
tmp4068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4068)
}
__typedArg0 := tmp4068
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4069)
}
__typedArg0 := tmp4069
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4070)
}
__typedArg0 := tmp4070
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4061 Obj

if True == tmp4071 {
tmp4063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4063)
}
__typedArg0 := tmp4063
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4064)
}
__typedArg0 := tmp4064
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4065)
}
__typedArg0 := tmp4065
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4066)
}
__typedArg0 := Nil
__typedArg1 := tmp4066
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4062 Obj

if True == tmp4067 {
ifres4062 = True


} else {
ifres4062 = False


}

ifres4061 = ifres4062


} else {
ifres4061 = False


}

var ifres4060 Obj

if True == ifres4061 {
ifres4060 = True


} else {
ifres4060 = False


}

ifres4059 = ifres4060


} else {
ifres4059 = False


}

var ifres4058 Obj

if True == ifres4059 {
ifres4058 = True


} else {
ifres4058 = False


}

ifres4057 = ifres4058


} else {
ifres4057 = False


}

var ifres4056 Obj

if True == ifres4057 {
ifres4056 = True


} else {
ifres4056 = False


}

ifres4055 = ifres4056


} else {
ifres4055 = False


}

var ifres4054 Obj

if True == ifres4055 {
ifres4054 = True


} else {
ifres4054 = False


}

ifres4053 = ifres4054


} else {
ifres4053 = False


}

if True == ifres4053 {
tmp4033 := Call(__e, PrimFunc(symgensym), symW)


let__2736 := tmp4033
_ = let__2736

tmp4034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4034)
}
__typedArg0 := tmp4034
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4035)
}
__typedArg0 := tmp4035
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4037)
}
__typedArg0 := tmp4037
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4039)
}
__typedArg0 := tmp4039
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4041 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4040)
}
__typedArg0 := tmp4040
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4041)
}
__typedArg0 := tmp4041
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4043 := Call(__e, PrimFunc(symshen_4beta), tmp4038, let__2736, tmp4042)


tmp4044 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4043, Nil)
}
__typedArg0 := tmp4043
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4036, tmp4044)
}
__typedArg0 := tmp4036
__typedArg1 := tmp4044
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2736, tmp4045)
}
__typedArg0 := let__2736
__typedArg1 := tmp4045
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp4046)
}
__typedArg0 := symlet
__typedArg1 := tmp4046
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__2737 := tmp4047
_ = let__2737

tmp4048 := MakeNative(func(__e *ControlFlow) {
Z815 := __e.Get(1)
_ = Z815
__e.TailApply(PrimFunc(symshen_4alpha_1convert), Z815)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp4048, let__2737)
return


} else {
tmp4051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp4051 {
tmp4049 := MakeNative(func(__e *ControlFlow) {
Z816 := __e.Get(1)
_ = Z816
__e.TailApply(PrimFunc(symshen_4alpha_1convert), Z816)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp4049, V809)
return


} else {
__e.Return(V809)
return
}


}


}


}, 1)

tmp4101 := Call(__e, ns2_1set, symshen_4alpha_1convert, tmp4021)


_ = tmp4101

tmp4102 := MakeNative(func(__e *ControlFlow) {
V817 := __e.Get(1)
_ = V817
V818 := __e.Get(2)
_ = V818
tmp4103 := MakeNative(func(__e *ControlFlow) {
Z819 := __e.Get(1)
_ = Z819
tmp4104 := Call(__e, PrimFunc(symfst), Z819)


tmp4105 := Call(__e, PrimFunc(symsnd), Z819)


tmp4106 := Call(__e, PrimFunc(symshen_4alpha_1convert), tmp4105)


__e.TailApply(PrimFunc(symshen_4triple_1stack), Nil, tmp4104, V818, tmp4106)
return


}, 1)

__e.TailApply(PrimFunc(symmap), tmp4103, V817)
return


}, 2)

tmp4107 := Call(__e, ns2_1set, symshen_4kl_1body, tmp4102)


_ = tmp4107

tmp4108 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__self4 := __e.Get(4)
__selftop:
V828 := __self1
_ = V828
V829 := __self2
_ = V829
V830 := __self3
_ = V830
V831 := __self4
_ = V831
tmp4257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V829)
}
__typedArg0 := Nil
__typedArg1 := V829
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4232 Obj

if True == tmp4257 {
tmp4256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V830)
}
__typedArg0 := Nil
__typedArg1 := V830
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4234 Obj

if True == tmp4256 {
tmp4255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V831)
}
__typedArg0 := V831
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4236 Obj

if True == tmp4255 {
tmp4253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V831)
}
__typedArg0 := V831
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp4253)
}
__typedArg0 := symwhere
__typedArg1 := tmp4253
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4238 Obj

if True == tmp4254 {
tmp4251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V831)
}
__typedArg0 := V831
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4251)
}
__typedArg0 := tmp4251
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4240 Obj

if True == tmp4252 {
tmp4248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V831)
}
__typedArg0 := V831
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4249 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4248)
}
__typedArg0 := tmp4248
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4250 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4249)
}
__typedArg0 := tmp4249
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4242 Obj

if True == tmp4250 {
tmp4244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V831)
}
__typedArg0 := V831
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4245 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4244)
}
__typedArg0 := tmp4244
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4245)
}
__typedArg0 := tmp4245
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4247 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4246)
}
__typedArg0 := Nil
__typedArg1 := tmp4246
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4243 Obj

if True == tmp4247 {
ifres4243 = True


} else {
ifres4243 = False


}

ifres4242 = ifres4243


} else {
ifres4242 = False


}

var ifres4241 Obj

if True == ifres4242 {
ifres4241 = True


} else {
ifres4241 = False


}

ifres4240 = ifres4241


} else {
ifres4240 = False


}

var ifres4239 Obj

if True == ifres4240 {
ifres4239 = True


} else {
ifres4239 = False


}

ifres4238 = ifres4239


} else {
ifres4238 = False


}

var ifres4237 Obj

if True == ifres4238 {
ifres4237 = True


} else {
ifres4237 = False


}

ifres4236 = ifres4237


} else {
ifres4236 = False


}

var ifres4235 Obj

if True == ifres4236 {
ifres4235 = True


} else {
ifres4235 = False


}

ifres4234 = ifres4235


} else {
ifres4234 = False


}

var ifres4233 Obj

if True == ifres4234 {
ifres4233 = True


} else {
ifres4233 = False


}

ifres4232 = ifres4233


} else {
ifres4232 = False


}

if True == ifres4232 {
tmp4109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V831)
}
__typedArg0 := V831
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4109)
}
__typedArg0 := tmp4109
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4110, V828)
}
__typedArg0 := tmp4110
__typedArg1 := V828
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V831)
}
__typedArg0 := V831
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4112)
}
__typedArg0 := tmp4112
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4113)
}
__typedArg0 := tmp4113
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

if PrimFunc(symshen_4triple_1stack) == __self {
__self1, __self2, __self3, __self4 = tmp4111, Nil, Nil, tmp4114
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp4111, Nil, Nil, tmp4114)
return


} else {
tmp4230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V829)
}
__typedArg0 := Nil
__typedArg1 := V829
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4227 Obj

if True == tmp4230 {
tmp4229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V830)
}
__typedArg0 := Nil
__typedArg1 := V830
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4228 Obj

if True == tmp4229 {
ifres4228 = True


} else {
ifres4228 = False


}

ifres4227 = ifres4228


} else {
ifres4227 = False


}

if True == ifres4227 {
tmp4115 := Call(__e, PrimFunc(symreverse), V828)


tmp4116 := Call(__e, PrimFunc(symshen_4rectify_1test), tmp4115)


tmp4117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V831, Nil)
}
__typedArg0 := V831
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4116, tmp4117)
}
__typedArg0 := tmp4116
__typedArg1 := tmp4117
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp4225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4218 Obj

if True == tmp4225 {
tmp4224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4220 Obj

if True == tmp4224 {
tmp4222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp4222)
}
__typedArg0 := tmp4222
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres4221 Obj

if True == tmp4223 {
ifres4221 = True


} else {
ifres4221 = False


}

ifres4220 = ifres4221


} else {
ifres4220 = False


}

var ifres4219 Obj

if True == ifres4220 {
ifres4219 = True


} else {
ifres4219 = False


}

ifres4218 = ifres4219


} else {
ifres4218 = False


}

if True == ifres4218 {
tmp4118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4122 := Call(__e, PrimFunc(symshen_4beta), tmp4120, tmp4121, V831)


if PrimFunc(symshen_4triple_1stack) == __self {
__self1, __self2, __self3, __self4 = V828, tmp4118, tmp4119, tmp4122
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4triple_1stack), V828, tmp4118, tmp4119, tmp4122)
return


} else {
tmp4216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4209 Obj

if True == tmp4216 {
tmp4215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4211 Obj

if True == tmp4215 {
tmp4213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4214 := Call(__e, PrimFunc(symshen_4custom_1pattern_2), tmp4213)


var ifres4212 Obj

if True == tmp4214 {
ifres4212 = True


} else {
ifres4212 = False


}

ifres4211 = ifres4212


} else {
ifres4211 = False


}

var ifres4210 Obj

if True == ifres4211 {
ifres4210 = True


} else {
ifres4210 = False


}

ifres4209 = ifres4210


} else {
ifres4209 = False


}

if True == ifres4209 {
tmp4123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4124 := Call(__e, PrimFunc(symshen_4custom_1pattern_1body), tmp4123)


tmp4125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4custom_1pattern_1triple_1stack), V828, tmp4124, tmp4125, tmp4126, tmp4127, V831)
return


} else {
tmp4207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4177 Obj

if True == tmp4207 {
tmp4205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4205)
}
__typedArg0 := tmp4205
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4179 Obj

if True == tmp4206 {
tmp4202 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4202)
}
__typedArg0 := tmp4202
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4203)
}
__typedArg0 := tmp4203
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4181 Obj

if True == tmp4204 {
tmp4198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4198)
}
__typedArg0 := tmp4198
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4199)
}
__typedArg0 := tmp4199
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4200)
}
__typedArg0 := tmp4200
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4183 Obj

if True == tmp4201 {
tmp4193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4193)
}
__typedArg0 := tmp4193
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4194)
}
__typedArg0 := tmp4194
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4195)
}
__typedArg0 := tmp4195
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4196)
}
__typedArg0 := Nil
__typedArg1 := tmp4196
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4185 Obj

if True == tmp4197 {
tmp4192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4187 Obj

if True == tmp4192 {
tmp4189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4189)
}
__typedArg0 := tmp4189
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4191 := Call(__e, PrimFunc(symshen_4constructor_2), tmp4190)


var ifres4188 Obj

if True == tmp4191 {
ifres4188 = True


} else {
ifres4188 = False


}

ifres4187 = ifres4188


} else {
ifres4187 = False


}

var ifres4186 Obj

if True == ifres4187 {
ifres4186 = True


} else {
ifres4186 = False


}

ifres4185 = ifres4186


} else {
ifres4185 = False


}

var ifres4184 Obj

if True == ifres4185 {
ifres4184 = True


} else {
ifres4184 = False


}

ifres4183 = ifres4184


} else {
ifres4183 = False


}

var ifres4182 Obj

if True == ifres4183 {
ifres4182 = True


} else {
ifres4182 = False


}

ifres4181 = ifres4182


} else {
ifres4181 = False


}

var ifres4180 Obj

if True == ifres4181 {
ifres4180 = True


} else {
ifres4180 = False


}

ifres4179 = ifres4180


} else {
ifres4179 = False


}

var ifres4178 Obj

if True == ifres4179 {
ifres4178 = True


} else {
ifres4178 = False


}

ifres4177 = ifres4178


} else {
ifres4177 = False


}

if True == ifres4177 {
tmp4128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4128)
}
__typedArg0 := tmp4128
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4130 := Call(__e, PrimFunc(symshen_4op_1test), tmp4129)


tmp4131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4131, Nil)
}
__typedArg0 := tmp4131
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4130, tmp4132)
}
__typedArg0 := tmp4130
__typedArg1 := tmp4132
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4133, V828)
}
__typedArg0 := tmp4133
__typedArg1 := V828
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4135)
}
__typedArg0 := tmp4135
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4136)
}
__typedArg0 := tmp4136
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4138)
}
__typedArg0 := tmp4138
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4139)
}
__typedArg0 := tmp4139
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4140)
}
__typedArg0 := tmp4140
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4141, tmp4142)
}
__typedArg0 := tmp4141
__typedArg1 := tmp4142
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4137, tmp4143)
}
__typedArg0 := tmp4137
__typedArg1 := tmp4143
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4145)
}
__typedArg0 := tmp4145
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4147 := Call(__e, PrimFunc(symshen_4op1), tmp4146)


tmp4148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4148, Nil)
}
__typedArg0 := tmp4148
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4147, tmp4149)
}
__typedArg0 := tmp4147
__typedArg1 := tmp4149
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4151)
}
__typedArg0 := tmp4151
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4153 := Call(__e, PrimFunc(symshen_4op2), tmp4152)


tmp4154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4154, Nil)
}
__typedArg0 := tmp4154
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4153, tmp4155)
}
__typedArg0 := tmp4153
__typedArg1 := tmp4155
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4156, tmp4157)
}
__typedArg0 := tmp4156
__typedArg1 := tmp4157
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4150, tmp4158)
}
__typedArg0 := tmp4150
__typedArg1 := tmp4158
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4162 := Call(__e, PrimFunc(symshen_4beta), tmp4160, tmp4161, V831)


if PrimFunc(symshen_4triple_1stack) == __self {
__self1, __self2, __self3, __self4 = tmp4134, tmp4144, tmp4159, tmp4162
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp4134, tmp4144, tmp4159, tmp4162)
return


} else {
tmp4175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4172 Obj

if True == tmp4175 {
tmp4174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4173 Obj

if True == tmp4174 {
ifres4173 = True


} else {
ifres4173 = False


}

ifres4172 = ifres4173


} else {
ifres4172 = False


}

if True == ifres4172 {
tmp4163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4164, Nil)
}
__typedArg0 := tmp4164
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4163, tmp4165)
}
__typedArg0 := tmp4163
__typedArg1 := tmp4165
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp4166)
}
__typedArg0 := sym_a
__typedArg1 := tmp4166
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4167, V828)
}
__typedArg0 := tmp4167
__typedArg1 := V828
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4triple_1stack) == __self {
__self1, __self2, __self3, __self4 = tmp4168, tmp4169, tmp4170, V831
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp4168, tmp4169, tmp4170, V831)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.triple-stack"))
}
__typedArg0 := MakeString("implementation error in shen.triple-stack")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}


}


}


}, 4)

tmp4258 := Call(__e, ns2_1set, symshen_4triple_1stack, tmp4108)


_ = tmp4258

tmp4259 := MakeNative(func(__e *ControlFlow) {
V832 := __e.Get(1)
_ = V832
V833 := __e.Get(2)
_ = V833
V834 := __e.Get(3)
_ = V834
V835 := __e.Get(4)
_ = V835
V836 := __e.Get(5)
_ = V836
V837 := __e.Get(6)
_ = V837
tmp4260 := Call(__e, PrimFunc(sym_8p), V833, V835)


tmp4261 := Call(__e, PrimFunc(symshen_4custom_1pattern_1reducer), tmp4260)


let__2738 := tmp4261
_ = let__2738

tmp4274 := Call(__e, PrimFunc(symtuple_2), let__2738)


if True == tmp4274 {
tmp4262 := Call(__e, PrimFunc(symfst), let__2738)


let__2739 := tmp4262
_ = let__2739

tmp4263 := Call(__e, PrimFunc(symsnd), let__2738)


let__2740 := tmp4263
_ = let__2740

tmp4264 := Call(__e, PrimFunc(symreverse), let__2739)


tmp4265 := Call(__e, PrimFunc(symappend), tmp4264, V832)


tmp4266 := MakeNative(func(__e *ControlFlow) {
Z841 := __e.Get(1)
_ = Z841
__e.TailApply(PrimFunc(symfst), Z841)
return
}, 1)

tmp4267 := Call(__e, PrimFunc(symmap), tmp4266, let__2740)


tmp4268 := Call(__e, PrimFunc(symappend), tmp4267, V834)


tmp4269 := MakeNative(func(__e *ControlFlow) {
Z842 := __e.Get(1)
_ = Z842
__e.TailApply(PrimFunc(symsnd), Z842)
return
}, 1)

tmp4270 := Call(__e, PrimFunc(symmap), tmp4269, let__2740)


tmp4271 := Call(__e, PrimFunc(symappend), tmp4270, V836)


tmp4272 := Call(__e, PrimFunc(symshen_4beta), V833, V835, V837)


__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp4265, tmp4268, tmp4271, tmp4272)
return


} else {
__e.TailApply(PrimFunc(symshen_4constructor_1error), V833)
return
}


}, 6)

tmp4275 := Call(__e, ns2_1set, symshen_4custom_1pattern_1triple_1stack, tmp4259)


_ = tmp4275

tmp4276 := MakeNative(func(__e *ControlFlow) {
V845 := __e.Get(1)
_ = V845
tmp4295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V845)
}
__typedArg0 := Nil
__typedArg1 := V845
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4295 {
__e.Return(True)
return
} else {
tmp4293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V845)
}
__typedArg0 := V845
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4289 Obj

if True == tmp4293 {
tmp4291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V845)
}
__typedArg0 := V845
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4291)
}
__typedArg0 := Nil
__typedArg1 := tmp4291
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4290 Obj

if True == tmp4292 {
ifres4290 = True


} else {
ifres4290 = False


}

ifres4289 = ifres4290


} else {
ifres4289 = False


}

if True == ifres4289 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V845)
}
__typedArg0 := V845
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp4287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V845)
}
__typedArg0 := V845
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4283 Obj

if True == tmp4287 {
tmp4285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V845)
}
__typedArg0 := V845
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4285)
}
__typedArg0 := tmp4285
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4284 Obj

if True == tmp4286 {
ifres4284 = True


} else {
ifres4284 = False


}

ifres4283 = ifres4284


} else {
ifres4283 = False


}

if True == ifres4283 {
tmp4277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V845)
}
__typedArg0 := V845
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V845)
}
__typedArg0 := V845
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4279 := Call(__e, PrimFunc(symshen_4rectify_1test), tmp4278)


tmp4280 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4279, Nil)
}
__typedArg0 := tmp4279
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4277, tmp4280)
}
__typedArg0 := tmp4277
__typedArg1 := tmp4280
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp4281)
}
__typedArg0 := symand
__typedArg1 := tmp4281
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.rectify-test"))
}
__typedArg0 := MakeString("implementation error in shen.rectify-test")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 1)

tmp4296 := Call(__e, ns2_1set, symshen_4rectify_1test, tmp4276)


_ = tmp4296

tmp4297 := MakeNative(func(__e *ControlFlow) {
V855 := __e.Get(1)
_ = V855
V856 := __e.Get(2)
_ = V856
V857 := __e.Get(3)
_ = V857
tmp4374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V855, V857)
}
__typedArg0 := V855
__typedArg1 := V857
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4374 {
__e.Return(V856)
return
} else {
tmp4372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4348 Obj

if True == tmp4372 {
tmp4370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp4370)
}
__typedArg0 := symlambda
__typedArg1 := tmp4370
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4350 Obj

if True == tmp4371 {
tmp4368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4368)
}
__typedArg0 := tmp4368
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4352 Obj

if True == tmp4369 {
tmp4365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4365)
}
__typedArg0 := tmp4365
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4366)
}
__typedArg0 := tmp4366
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4354 Obj

if True == tmp4367 {
tmp4361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4361)
}
__typedArg0 := tmp4361
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4362)
}
__typedArg0 := tmp4362
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4363)
}
__typedArg0 := Nil
__typedArg1 := tmp4363
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4356 Obj

if True == tmp4364 {
tmp4358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4358)
}
__typedArg0 := tmp4358
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V855, tmp4359)
}
__typedArg0 := V855
__typedArg1 := tmp4359
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4357 Obj

if True == tmp4360 {
ifres4357 = True


} else {
ifres4357 = False


}

ifres4356 = ifres4357


} else {
ifres4356 = False


}

var ifres4355 Obj

if True == ifres4356 {
ifres4355 = True


} else {
ifres4355 = False


}

ifres4354 = ifres4355


} else {
ifres4354 = False


}

var ifres4353 Obj

if True == ifres4354 {
ifres4353 = True


} else {
ifres4353 = False


}

ifres4352 = ifres4353


} else {
ifres4352 = False


}

var ifres4351 Obj

if True == ifres4352 {
ifres4351 = True


} else {
ifres4351 = False


}

ifres4350 = ifres4351


} else {
ifres4350 = False


}

var ifres4349 Obj

if True == ifres4350 {
ifres4349 = True


} else {
ifres4349 = False


}

ifres4348 = ifres4349


} else {
ifres4348 = False


}

if True == ifres4348 {
__e.Return(V857)
return
} else {
tmp4346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4315 Obj

if True == tmp4346 {
tmp4344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp4344)
}
__typedArg0 := symlet
__typedArg1 := tmp4344
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4317 Obj

if True == tmp4345 {
tmp4342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4342)
}
__typedArg0 := tmp4342
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4319 Obj

if True == tmp4343 {
tmp4339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4339)
}
__typedArg0 := tmp4339
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4340)
}
__typedArg0 := tmp4340
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4321 Obj

if True == tmp4341 {
tmp4335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4335)
}
__typedArg0 := tmp4335
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4336)
}
__typedArg0 := tmp4336
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4337)
}
__typedArg0 := tmp4337
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4323 Obj

if True == tmp4338 {
tmp4330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4330)
}
__typedArg0 := tmp4330
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4332 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4331)
}
__typedArg0 := tmp4331
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4333 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4332)
}
__typedArg0 := tmp4332
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4333)
}
__typedArg0 := Nil
__typedArg1 := tmp4333
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4325 Obj

if True == tmp4334 {
tmp4327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4327)
}
__typedArg0 := tmp4327
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V855, tmp4328)
}
__typedArg0 := V855
__typedArg1 := tmp4328
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4326 Obj

if True == tmp4329 {
ifres4326 = True


} else {
ifres4326 = False


}

ifres4325 = ifres4326


} else {
ifres4325 = False


}

var ifres4324 Obj

if True == ifres4325 {
ifres4324 = True


} else {
ifres4324 = False


}

ifres4323 = ifres4324


} else {
ifres4323 = False


}

var ifres4322 Obj

if True == ifres4323 {
ifres4322 = True


} else {
ifres4322 = False


}

ifres4321 = ifres4322


} else {
ifres4321 = False


}

var ifres4320 Obj

if True == ifres4321 {
ifres4320 = True


} else {
ifres4320 = False


}

ifres4319 = ifres4320


} else {
ifres4319 = False


}

var ifres4318 Obj

if True == ifres4319 {
ifres4318 = True


} else {
ifres4318 = False


}

ifres4317 = ifres4318


} else {
ifres4317 = False


}

var ifres4316 Obj

if True == ifres4317 {
ifres4316 = True


} else {
ifres4316 = False


}

ifres4315 = ifres4316


} else {
ifres4315 = False


}

if True == ifres4315 {
tmp4298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4298)
}
__typedArg0 := tmp4298
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4300)
}
__typedArg0 := tmp4300
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4302)
}
__typedArg0 := tmp4302
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4303)
}
__typedArg0 := tmp4303
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4305 := Call(__e, PrimFunc(symshen_4beta), tmp4301, V856, tmp4304)


tmp4306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4306)
}
__typedArg0 := tmp4306
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4307)
}
__typedArg0 := tmp4307
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4305, tmp4308)
}
__typedArg0 := tmp4305
__typedArg1 := tmp4308
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4299, tmp4309)
}
__typedArg0 := tmp4299
__typedArg1 := tmp4309
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp4310)
}
__typedArg0 := symlet
__typedArg1 := tmp4310
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp4313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp4313 {
tmp4311 := MakeNative(func(__e *ControlFlow) {
Z858 := __e.Get(1)
_ = Z858
__e.TailApply(PrimFunc(symshen_4beta), V855, V856, Z858)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp4311, V857)
return


} else {
__e.Return(V857)
return
}


}


}


}


}, 3)

tmp4375 := Call(__e, ns2_1set, symshen_4beta, tmp4297)


_ = tmp4375

tmp4376 := MakeNative(func(__e *ControlFlow) {
V861 := __e.Get(1)
_ = V861
tmp4384 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, V861)
}
__typedArg0 := symcons
__typedArg1 := V861
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4384 {
__e.Return(symhd)
return
} else {
tmp4382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8s, V861)
}
__typedArg0 := sym_8s
__typedArg1 := V861
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4382 {
__e.Return(symhdstr)
return
} else {
tmp4380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8p, V861)
}
__typedArg0 := sym_8p
__typedArg1 := V861
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4380 {
__e.Return(symfst)
return
} else {
tmp4378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8v, V861)
}
__typedArg0 := sym_8v
__typedArg1 := V861
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4378 {
__e.Return(symhdv)
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.op1"))
}
__typedArg0 := MakeString("implementation error in shen.op1")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}


}, 1)

tmp4385 := Call(__e, ns2_1set, symshen_4op1, tmp4376)


_ = tmp4385

tmp4386 := MakeNative(func(__e *ControlFlow) {
V864 := __e.Get(1)
_ = V864
tmp4394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, V864)
}
__typedArg0 := symcons
__typedArg1 := V864
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4394 {
__e.Return(symtl)
return
} else {
tmp4392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8s, V864)
}
__typedArg0 := sym_8s
__typedArg1 := V864
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4392 {
__e.Return(symtlstr)
return
} else {
tmp4390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8p, V864)
}
__typedArg0 := sym_8p
__typedArg1 := V864
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4390 {
__e.Return(symsnd)
return
} else {
tmp4388 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8v, V864)
}
__typedArg0 := sym_8v
__typedArg1 := V864
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4388 {
__e.Return(symtlv)
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.op2"))
}
__typedArg0 := MakeString("implementation error in shen.op2")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}


}, 1)

tmp4395 := Call(__e, ns2_1set, symshen_4op2, tmp4386)


_ = tmp4395

tmp4396 := MakeNative(func(__e *ControlFlow) {
V867 := __e.Get(1)
_ = V867
tmp4404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, V867)
}
__typedArg0 := symcons
__typedArg1 := V867
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4404 {
__e.Return(symcons_2)
return
} else {
tmp4402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8s, V867)
}
__typedArg0 := sym_8s
__typedArg1 := V867
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4402 {
__e.Return(symshen_4_7string_2)
return
} else {
tmp4400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8p, V867)
}
__typedArg0 := sym_8p
__typedArg1 := V867
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4400 {
__e.Return(symtuple_2)
return
} else {
tmp4398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8v, V867)
}
__typedArg0 := sym_8v
__typedArg1 := V867
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4398 {
__e.Return(symshen_4_7vector_2)
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("implementation error in shen.op-test"))
}
__typedArg0 := MakeString("implementation error in shen.op-test")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}


}, 1)

tmp4405 := Call(__e, ns2_1set, symshen_4op_1test, tmp4396)


_ = tmp4405

tmp4406 := MakeNative(func(__e *ControlFlow) {
V868 := __e.Get(1)
_ = V868
tmp4408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V868)
}
__typedArg0 := MakeString("")
__typedArg1 := V868
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4408 {
__e.Return(False)
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V868)
}
__typedArg0 := V868
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})())
return
}


}, 1)

tmp4409 := Call(__e, ns2_1set, symshen_4_7string_2, tmp4406)


_ = tmp4409

tmp4410 := MakeNative(func(__e *ControlFlow) {
V869 := __e.Get(1)
_ = V869
tmp4412 := Call(__e, PrimFunc(symvector), MakeInteger(0))


tmp4413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V869, tmp4412)
}
__typedArg0 := V869
__typedArg1 := tmp4412
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4413 {
__e.Return(False)
return
} else {
__e.TailApply(PrimFunc(symvector_2), V869)
return
}


}, 1)

tmp4414 := Call(__e, ns2_1set, symshen_4_7vector_2, tmp4410)


_ = tmp4414

tmp4415 := MakeNative(func(__e *ControlFlow) {
V872 := __e.Get(1)
_ = V872
tmp4419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V872)
}
__typedArg0 := sym_7
__typedArg1 := V872
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4419 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dfactorise_2_d, True)
}
__typedArg0 := symshen_4_dfactorise_2_d
__typedArg1 := True
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
tmp4417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V872)
}
__typedArg0 := sym_1
__typedArg1 := V872
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4417 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_dfactorise_2_d, False)
}
__typedArg0 := symshen_4_dfactorise_2_d
__typedArg1 := False
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})())
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("factorise expects a + or a -\n"))
}
__typedArg0 := MakeString("factorise expects a + or a -\n")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 1)

tmp4420 := Call(__e, ns2_1set, symfactorise, tmp4415)


_ = tmp4420

tmp4421 := MakeNative(func(__e *ControlFlow) {
V873 := __e.Get(1)
_ = V873
tmp4423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dfactorise_2_d)
}
__typedArg0 := symshen_4_dfactorise_2_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

if True == tmp4423 {
__e.TailApply(PrimFunc(symshen_4factor), V873)
return
} else {
__e.Return(V873)
return
}


}, 1)

tmp4424 := Call(__e, ns2_1set, symshen_4factorise_1code, tmp4421)


_ = tmp4424

tmp4425 := MakeNative(func(__e *ControlFlow) {
V874 := __e.Get(1)
_ = V874
tmp4482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V874)
}
__typedArg0 := V874
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4441 Obj

if True == tmp4482 {
tmp4480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V874)
}
__typedArg0 := V874
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp4480)
}
__typedArg0 := symdefun
__typedArg1 := tmp4480
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4443 Obj

if True == tmp4481 {
tmp4478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V874)
}
__typedArg0 := V874
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4478)
}
__typedArg0 := tmp4478
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4445 Obj

if True == tmp4479 {
tmp4475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V874)
}
__typedArg0 := V874
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4475)
}
__typedArg0 := tmp4475
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4476)
}
__typedArg0 := tmp4476
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4447 Obj

if True == tmp4477 {
tmp4471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V874)
}
__typedArg0 := V874
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4471)
}
__typedArg0 := tmp4471
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4472)
}
__typedArg0 := tmp4472
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4473)
}
__typedArg0 := tmp4473
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4449 Obj

if True == tmp4474 {
tmp4466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V874)
}
__typedArg0 := V874
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4466)
}
__typedArg0 := tmp4466
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4467)
}
__typedArg0 := tmp4467
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4468)
}
__typedArg0 := tmp4468
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4469)
}
__typedArg0 := tmp4469
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4451 Obj

if True == tmp4470 {
tmp4460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V874)
}
__typedArg0 := V874
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4460)
}
__typedArg0 := tmp4460
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4461)
}
__typedArg0 := tmp4461
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4462)
}
__typedArg0 := tmp4462
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4463)
}
__typedArg0 := tmp4463
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcond, tmp4464)
}
__typedArg0 := symcond
__typedArg1 := tmp4464
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4453 Obj

if True == tmp4465 {
tmp4455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V874)
}
__typedArg0 := V874
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4455)
}
__typedArg0 := tmp4455
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4456)
}
__typedArg0 := tmp4456
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4457)
}
__typedArg0 := tmp4457
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4458)
}
__typedArg0 := Nil
__typedArg1 := tmp4458
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4454 Obj

if True == tmp4459 {
ifres4454 = True


} else {
ifres4454 = False


}

ifres4453 = ifres4454


} else {
ifres4453 = False


}

var ifres4452 Obj

if True == ifres4453 {
ifres4452 = True


} else {
ifres4452 = False


}

ifres4451 = ifres4452


} else {
ifres4451 = False


}

var ifres4450 Obj

if True == ifres4451 {
ifres4450 = True


} else {
ifres4450 = False


}

ifres4449 = ifres4450


} else {
ifres4449 = False


}

var ifres4448 Obj

if True == ifres4449 {
ifres4448 = True


} else {
ifres4448 = False


}

ifres4447 = ifres4448


} else {
ifres4447 = False


}

var ifres4446 Obj

if True == ifres4447 {
ifres4446 = True


} else {
ifres4446 = False


}

ifres4445 = ifres4446


} else {
ifres4445 = False


}

var ifres4444 Obj

if True == ifres4445 {
ifres4444 = True


} else {
ifres4444 = False


}

ifres4443 = ifres4444


} else {
ifres4443 = False


}

var ifres4442 Obj

if True == ifres4443 {
ifres4442 = True


} else {
ifres4442 = False


}

ifres4441 = ifres4442


} else {
ifres4441 = False


}

if True == ifres4441 {
tmp4426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V874)
}
__typedArg0 := V874
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4426)
}
__typedArg0 := tmp4426
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V874)
}
__typedArg0 := V874
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4428)
}
__typedArg0 := tmp4428
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4429)
}
__typedArg0 := tmp4429
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V874)
}
__typedArg0 := V874
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4431)
}
__typedArg0 := tmp4431
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4432)
}
__typedArg0 := tmp4432
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4433)
}
__typedArg0 := tmp4433
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4434)
}
__typedArg0 := tmp4434
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4436 := Call(__e, PrimFunc(symshen_4factor_1recognisors), tmp4435)


tmp4437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4436, Nil)
}
__typedArg0 := tmp4436
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4430, tmp4437)
}
__typedArg0 := tmp4430
__typedArg1 := tmp4437
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4427, tmp4438)
}
__typedArg0 := tmp4427
__typedArg1 := tmp4438
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp4439)
}
__typedArg0 := symdefun
__typedArg1 := tmp4439
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V874)
return
}


}, 1)

tmp4483 := Call(__e, ns2_1set, symshen_4factor, tmp4425)


_ = tmp4483

tmp4484 := MakeNative(func(__e *ControlFlow) {
V877 := __e.Get(1)
_ = V877
tmp4637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4617 Obj

if True == tmp4637 {
tmp4635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4635)
}
__typedArg0 := tmp4635
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4619 Obj

if True == tmp4636 {
tmp4632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4632)
}
__typedArg0 := tmp4632
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, tmp4633)
}
__typedArg0 := True
__typedArg1 := tmp4633
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4621 Obj

if True == tmp4634 {
tmp4629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4629)
}
__typedArg0 := tmp4629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4630)
}
__typedArg0 := tmp4630
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4623 Obj

if True == tmp4631 {
tmp4625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4625)
}
__typedArg0 := tmp4625
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4626)
}
__typedArg0 := tmp4626
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4627)
}
__typedArg0 := Nil
__typedArg1 := tmp4627
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4624 Obj

if True == tmp4628 {
ifres4624 = True


} else {
ifres4624 = False


}

ifres4623 = ifres4624


} else {
ifres4623 = False


}

var ifres4622 Obj

if True == ifres4623 {
ifres4622 = True


} else {
ifres4622 = False


}

ifres4621 = ifres4622


} else {
ifres4621 = False


}

var ifres4620 Obj

if True == ifres4621 {
ifres4620 = True


} else {
ifres4620 = False


}

ifres4619 = ifres4620


} else {
ifres4619 = False


}

var ifres4618 Obj

if True == ifres4619 {
ifres4618 = True


} else {
ifres4618 = False


}

ifres4617 = ifres4618


} else {
ifres4617 = False


}

if True == ifres4617 {
tmp4485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4485)
}
__typedArg0 := tmp4485
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4486)
}
__typedArg0 := tmp4486
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp4615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4568 Obj

if True == tmp4615 {
tmp4613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4613)
}
__typedArg0 := tmp4613
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4570 Obj

if True == tmp4614 {
tmp4610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4610)
}
__typedArg0 := tmp4610
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4611)
}
__typedArg0 := tmp4611
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4572 Obj

if True == tmp4612 {
tmp4606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4606)
}
__typedArg0 := tmp4606
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4607)
}
__typedArg0 := tmp4607
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symand, tmp4608)
}
__typedArg0 := symand
__typedArg1 := tmp4608
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4574 Obj

if True == tmp4609 {
tmp4602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4602)
}
__typedArg0 := tmp4602
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4603)
}
__typedArg0 := tmp4603
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4604)
}
__typedArg0 := tmp4604
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4576 Obj

if True == tmp4605 {
tmp4597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4597)
}
__typedArg0 := tmp4597
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4598)
}
__typedArg0 := tmp4598
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4599)
}
__typedArg0 := tmp4599
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4600)
}
__typedArg0 := tmp4600
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4578 Obj

if True == tmp4601 {
tmp4591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4591)
}
__typedArg0 := tmp4591
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4592)
}
__typedArg0 := tmp4592
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4593)
}
__typedArg0 := tmp4593
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4594)
}
__typedArg0 := tmp4594
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4595)
}
__typedArg0 := Nil
__typedArg1 := tmp4595
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4580 Obj

if True == tmp4596 {
tmp4588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4588)
}
__typedArg0 := tmp4588
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4589)
}
__typedArg0 := tmp4589
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4582 Obj

if True == tmp4590 {
tmp4584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4584)
}
__typedArg0 := tmp4584
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4585)
}
__typedArg0 := tmp4585
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4586)
}
__typedArg0 := Nil
__typedArg1 := tmp4586
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4583 Obj

if True == tmp4587 {
ifres4583 = True


} else {
ifres4583 = False


}

ifres4582 = ifres4583


} else {
ifres4582 = False


}

var ifres4581 Obj

if True == ifres4582 {
ifres4581 = True


} else {
ifres4581 = False


}

ifres4580 = ifres4581


} else {
ifres4580 = False


}

var ifres4579 Obj

if True == ifres4580 {
ifres4579 = True


} else {
ifres4579 = False


}

ifres4578 = ifres4579


} else {
ifres4578 = False


}

var ifres4577 Obj

if True == ifres4578 {
ifres4577 = True


} else {
ifres4577 = False


}

ifres4576 = ifres4577


} else {
ifres4576 = False


}

var ifres4575 Obj

if True == ifres4576 {
ifres4575 = True


} else {
ifres4575 = False


}

ifres4574 = ifres4575


} else {
ifres4574 = False


}

var ifres4573 Obj

if True == ifres4574 {
ifres4573 = True


} else {
ifres4573 = False


}

ifres4572 = ifres4573


} else {
ifres4572 = False


}

var ifres4571 Obj

if True == ifres4572 {
ifres4571 = True


} else {
ifres4571 = False


}

ifres4570 = ifres4571


} else {
ifres4570 = False


}

var ifres4569 Obj

if True == ifres4570 {
ifres4569 = True


} else {
ifres4569 = False


}

ifres4568 = ifres4569


} else {
ifres4568 = False


}

if True == ifres4568 {
tmp4487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4487)
}
__typedArg0 := tmp4487
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4488)
}
__typedArg0 := tmp4488
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4489)
}
__typedArg0 := tmp4489
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4491 := Call(__e, PrimFunc(symshen_4pivot_1on), tmp4490, V877, Nil)


let__2741 := tmp4491
_ = let__2741

tmp4492 := Call(__e, PrimFunc(symfst), let__2741)


let__2742 := tmp4492
_ = let__2742

tmp4536 := Call(__e, PrimFunc(symshen_4bad_1pivot_2), let__2742)


if True == tmp4536 {
tmp4493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, Nil)
}
__typedArg0 := symif
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4494)
}
__typedArg0 := tmp4494
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4496)
}
__typedArg0 := tmp4496
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4497)
}
__typedArg0 := tmp4497
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4499 := Call(__e, PrimFunc(symshen_4recursively_1factor_1selectors), tmp4495, tmp4498)


tmp4500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4501 := Call(__e, PrimFunc(symshen_4factor_1recognisors), tmp4500)


tmp4502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4501, Nil)
}
__typedArg0 := tmp4501
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4503 := Call(__e, PrimFunc(symappend), tmp4499, tmp4502)


__e.TailApply(PrimFunc(symappend), tmp4493, tmp4503)
return


} else {
tmp4504 := Call(__e, PrimFunc(symsnd), let__2741)


let__2743 := tmp4504
_ = let__2743

tmp4505 := Call(__e, PrimFunc(symshen_4factor_1recognisors), let__2743)


let__2744 := tmp4505
_ = let__2744

tmp4506 := Call(__e, PrimFunc(symgensym), symGoTo)


let__2745 := tmp4506
_ = let__2745

tmp4507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2745, Nil)
}
__typedArg0 := let__2745
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp4507)
}
__typedArg0 := symthaw
__typedArg1 := tmp4507
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4508, Nil)
}
__typedArg0 := tmp4508
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp4509)
}
__typedArg0 := True
__typedArg1 := tmp4509
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4510, let__2742)
}
__typedArg0 := tmp4510
__typedArg1 := let__2742
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4512 := Call(__e, PrimFunc(symreverse), tmp4511)


let__2746 := tmp4512
_ = let__2746

tmp4513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2744, Nil)
}
__typedArg0 := let__2744
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp4513)
}
__typedArg0 := symfreeze
__typedArg1 := tmp4513
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4515)
}
__typedArg0 := tmp4515
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4516)
}
__typedArg0 := tmp4516
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4517)
}
__typedArg0 := tmp4517
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4519)
}
__typedArg0 := tmp4519
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4520)
}
__typedArg0 := tmp4520
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4521)
}
__typedArg0 := tmp4521
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4523 := Call(__e, PrimFunc(symshen_4factor_1recognisors), let__2746)


tmp4524 := Call(__e, PrimFunc(symshen_4factor_1selectors), tmp4522, tmp4523)


tmp4525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2745, Nil)
}
__typedArg0 := let__2745
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp4525)
}
__typedArg0 := symthaw
__typedArg1 := tmp4525
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4526, Nil)
}
__typedArg0 := tmp4526
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4524, tmp4527)
}
__typedArg0 := tmp4524
__typedArg1 := tmp4527
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4518, tmp4528)
}
__typedArg0 := tmp4518
__typedArg1 := tmp4528
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp4529)
}
__typedArg0 := symif
__typedArg1 := tmp4529
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4530, Nil)
}
__typedArg0 := tmp4530
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4514, tmp4531)
}
__typedArg0 := tmp4514
__typedArg1 := tmp4531
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2745, tmp4532)
}
__typedArg0 := let__2745
__typedArg1 := tmp4532
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp4533)
}
__typedArg0 := symlet
__typedArg1 := tmp4533
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__2747 := tmp4534
_ = let__2747

__e.TailApply(PrimFunc(symshen_4remove_1indirection), let__2747)
return


}


} else {
tmp4566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4551 Obj

if True == tmp4566 {
tmp4564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4564)
}
__typedArg0 := tmp4564
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4553 Obj

if True == tmp4565 {
tmp4561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4561)
}
__typedArg0 := tmp4561
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4562)
}
__typedArg0 := tmp4562
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4555 Obj

if True == tmp4563 {
tmp4557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4557)
}
__typedArg0 := tmp4557
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4558)
}
__typedArg0 := tmp4558
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4559)
}
__typedArg0 := Nil
__typedArg1 := tmp4559
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4556 Obj

if True == tmp4560 {
ifres4556 = True


} else {
ifres4556 = False


}

ifres4555 = ifres4556


} else {
ifres4555 = False


}

var ifres4554 Obj

if True == ifres4555 {
ifres4554 = True


} else {
ifres4554 = False


}

ifres4553 = ifres4554


} else {
ifres4553 = False


}

var ifres4552 Obj

if True == ifres4553 {
ifres4552 = True


} else {
ifres4552 = False


}

ifres4551 = ifres4552


} else {
ifres4551 = False


}

if True == ifres4551 {
tmp4537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4537)
}
__typedArg0 := tmp4537
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4539)
}
__typedArg0 := tmp4539
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4540)
}
__typedArg0 := tmp4540
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4543 := Call(__e, PrimFunc(symshen_4factor_1recognisors), tmp4542)


tmp4544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4543, Nil)
}
__typedArg0 := tmp4543
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4541, tmp4544)
}
__typedArg0 := tmp4541
__typedArg1 := tmp4544
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4538, tmp4545)
}
__typedArg0 := tmp4538
__typedArg1 := tmp4545
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp4546)
}
__typedArg0 := symif
__typedArg1 := tmp4546
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp4549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V877)
}
__typedArg0 := Nil
__typedArg1 := V877
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4549 {
tmp4547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeString("condition failure"), Nil)
}
__typedArg0 := MakeString("condition failure")
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symsimple_1error, tmp4547)
}
__typedArg0 := symsimple_1error
__typedArg1 := tmp4547
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4factor_1recognisors)
return
}


}


}


}


}, 1)

tmp4638 := Call(__e, ns2_1set, symshen_4factor_1recognisors, tmp4484)


_ = tmp4638

tmp4639 := MakeNative(func(__e *ControlFlow) {
V885 := __e.Get(1)
_ = V885
V886 := __e.Get(2)
_ = V886
tmp4711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4685 Obj

if True == tmp4711 {
tmp4709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp4709)
}
__typedArg0 := symlet
__typedArg1 := tmp4709
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4687 Obj

if True == tmp4710 {
tmp4707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4707)
}
__typedArg0 := tmp4707
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4689 Obj

if True == tmp4708 {
tmp4704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4704)
}
__typedArg0 := tmp4704
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4705)
}
__typedArg0 := tmp4705
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4691 Obj

if True == tmp4706 {
tmp4700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4700)
}
__typedArg0 := tmp4700
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4701)
}
__typedArg0 := tmp4701
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4702)
}
__typedArg0 := tmp4702
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4693 Obj

if True == tmp4703 {
tmp4695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4695)
}
__typedArg0 := tmp4695
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4696)
}
__typedArg0 := tmp4696
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4697)
}
__typedArg0 := tmp4697
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4698)
}
__typedArg0 := Nil
__typedArg1 := tmp4698
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4694 Obj

if True == tmp4699 {
ifres4694 = True


} else {
ifres4694 = False


}

ifres4693 = ifres4694


} else {
ifres4693 = False


}

var ifres4692 Obj

if True == ifres4693 {
ifres4692 = True


} else {
ifres4692 = False


}

ifres4691 = ifres4692


} else {
ifres4691 = False


}

var ifres4690 Obj

if True == ifres4691 {
ifres4690 = True


} else {
ifres4690 = False


}

ifres4689 = ifres4690


} else {
ifres4689 = False


}

var ifres4688 Obj

if True == ifres4689 {
ifres4688 = True


} else {
ifres4688 = False


}

ifres4687 = ifres4688


} else {
ifres4687 = False


}

var ifres4686 Obj

if True == ifres4687 {
ifres4686 = True


} else {
ifres4686 = False


}

ifres4685 = ifres4686


} else {
ifres4685 = False


}

if True == ifres4685 {
tmp4640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4640)
}
__typedArg0 := tmp4640
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4642)
}
__typedArg0 := tmp4642
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4643)
}
__typedArg0 := tmp4643
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4645)
}
__typedArg0 := tmp4645
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4646)
}
__typedArg0 := tmp4646
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4647)
}
__typedArg0 := tmp4647
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4649 := Call(__e, PrimFunc(symshen_4recursively_1factor_1selectors), tmp4648, V886)


__e.TailApply(PrimFunc(symshen_4restore_1local), tmp4641, tmp4644, tmp4649)
return


} else {
tmp4683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4664 Obj

if True == tmp4683 {
tmp4681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symand, tmp4681)
}
__typedArg0 := symand
__typedArg1 := tmp4681
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4666 Obj

if True == tmp4682 {
tmp4679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4679)
}
__typedArg0 := tmp4679
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4668 Obj

if True == tmp4680 {
tmp4676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4676)
}
__typedArg0 := tmp4676
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4677)
}
__typedArg0 := tmp4677
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4670 Obj

if True == tmp4678 {
tmp4672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4672)
}
__typedArg0 := tmp4672
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4673)
}
__typedArg0 := tmp4673
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4674)
}
__typedArg0 := Nil
__typedArg1 := tmp4674
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4671 Obj

if True == tmp4675 {
ifres4671 = True


} else {
ifres4671 = False


}

ifres4670 = ifres4671


} else {
ifres4670 = False


}

var ifres4669 Obj

if True == ifres4670 {
ifres4669 = True


} else {
ifres4669 = False


}

ifres4668 = ifres4669


} else {
ifres4668 = False


}

var ifres4667 Obj

if True == ifres4668 {
ifres4667 = True


} else {
ifres4667 = False


}

ifres4666 = ifres4667


} else {
ifres4666 = False


}

var ifres4665 Obj

if True == ifres4666 {
ifres4665 = True


} else {
ifres4665 = False


}

ifres4664 = ifres4665


} else {
ifres4664 = False


}

if True == ifres4664 {
tmp4650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4650)
}
__typedArg0 := tmp4650
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4652)
}
__typedArg0 := tmp4652
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4654)
}
__typedArg0 := tmp4654
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4655)
}
__typedArg0 := tmp4655
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4657 := Call(__e, PrimFunc(symshen_4factor_1selectors), tmp4653, tmp4656)


tmp4658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4658)
}
__typedArg0 := tmp4658
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4660 := Call(__e, PrimFunc(symshen_4factor_1selectors), tmp4659, V886)


tmp4661 := Call(__e, PrimFunc(symshen_4recursively_1factor_1selectors), tmp4657, tmp4660)


__e.TailApply(PrimFunc(symshen_4restore_1P), tmp4651, tmp4661)
return


} else {
tmp4662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V886, Nil)
}
__typedArg0 := V886
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V885, tmp4662)
}
__typedArg0 := V885
__typedArg1 := tmp4662
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}


}, 2)

tmp4712 := Call(__e, ns2_1set, symshen_4recursively_1factor_1selectors, tmp4639)


_ = tmp4712

tmp4713 := MakeNative(func(__e *ControlFlow) {
V887 := __e.Get(1)
_ = V887
V888 := __e.Get(2)
_ = V888
tmp4729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V888)
}
__typedArg0 := V888
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4720 Obj

if True == tmp4729 {
tmp4727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V888)
}
__typedArg0 := V888
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4727)
}
__typedArg0 := tmp4727
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4722 Obj

if True == tmp4728 {
tmp4724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V888)
}
__typedArg0 := V888
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4724)
}
__typedArg0 := tmp4724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4725)
}
__typedArg0 := Nil
__typedArg1 := tmp4725
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4723 Obj

if True == tmp4726 {
ifres4723 = True


} else {
ifres4723 = False


}

ifres4722 = ifres4723


} else {
ifres4722 = False


}

var ifres4721 Obj

if True == ifres4722 {
ifres4721 = True


} else {
ifres4721 = False


}

ifres4720 = ifres4721


} else {
ifres4720 = False


}

if True == ifres4720 {
tmp4714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V888)
}
__typedArg0 := V888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4714, Nil)
}
__typedArg0 := tmp4714
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V887, tmp4715)
}
__typedArg0 := V887
__typedArg1 := tmp4715
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp4716)
}
__typedArg0 := symand
__typedArg1 := tmp4716
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V888)
}
__typedArg0 := V888
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4717, tmp4718)
}
__typedArg0 := tmp4717
__typedArg1 := tmp4718
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4restore_1P)
return
}


}, 2)

tmp4730 := Call(__e, ns2_1set, symshen_4restore_1P, tmp4713)


_ = tmp4730

tmp4731 := MakeNative(func(__e *ControlFlow) {
V889 := __e.Get(1)
_ = V889
V890 := __e.Get(2)
_ = V890
V891 := __e.Get(3)
_ = V891
tmp4748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V891)
}
__typedArg0 := V891
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4739 Obj

if True == tmp4748 {
tmp4746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V891)
}
__typedArg0 := V891
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4746)
}
__typedArg0 := tmp4746
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4741 Obj

if True == tmp4747 {
tmp4743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V891)
}
__typedArg0 := V891
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4743)
}
__typedArg0 := tmp4743
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4744)
}
__typedArg0 := Nil
__typedArg1 := tmp4744
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4742 Obj

if True == tmp4745 {
ifres4742 = True


} else {
ifres4742 = False


}

ifres4741 = ifres4742


} else {
ifres4741 = False


}

var ifres4740 Obj

if True == ifres4741 {
ifres4740 = True


} else {
ifres4740 = False


}

ifres4739 = ifres4740


} else {
ifres4739 = False


}

if True == ifres4739 {
tmp4732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V891)
}
__typedArg0 := V891
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4732, Nil)
}
__typedArg0 := tmp4732
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V890, tmp4733)
}
__typedArg0 := V890
__typedArg1 := tmp4733
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V889, tmp4734)
}
__typedArg0 := V889
__typedArg1 := tmp4734
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp4735)
}
__typedArg0 := symlet
__typedArg1 := tmp4735
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V891)
}
__typedArg0 := V891
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4736, tmp4737)
}
__typedArg0 := tmp4736
__typedArg1 := tmp4737
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4restore_1local)
return
}


}, 3)

tmp4749 := Call(__e, ns2_1set, symshen_4restore_1local, tmp4731)


_ = tmp4749

tmp4750 := MakeNative(func(__e *ControlFlow) {
V896 := __e.Get(1)
_ = V896
tmp4756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V896)
}
__typedArg0 := V896
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4752 Obj

if True == tmp4756 {
tmp4754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V896)
}
__typedArg0 := V896
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4754)
}
__typedArg0 := Nil
__typedArg1 := tmp4754
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4753 Obj

if True == tmp4755 {
ifres4753 = True


} else {
ifres4753 = False


}

ifres4752 = ifres4753


} else {
ifres4752 = False


}

if True == ifres4752 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp4757 := Call(__e, ns2_1set, symshen_4bad_1pivot_2, tmp4750)


_ = tmp4757

tmp4758 := MakeNative(func(__e *ControlFlow) {
V897 := __e.Get(1)
_ = V897
tmp4873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4773 Obj

if True == tmp4873 {
tmp4871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp4871)
}
__typedArg0 := symlet
__typedArg1 := tmp4871
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4775 Obj

if True == tmp4872 {
tmp4869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4869)
}
__typedArg0 := tmp4869
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4777 Obj

if True == tmp4870 {
tmp4866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4866)
}
__typedArg0 := tmp4866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4867)
}
__typedArg0 := tmp4867
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4779 Obj

if True == tmp4868 {
tmp4862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4862)
}
__typedArg0 := tmp4862
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4863)
}
__typedArg0 := tmp4863
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4864)
}
__typedArg0 := tmp4864
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4781 Obj

if True == tmp4865 {
tmp4857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4857)
}
__typedArg0 := tmp4857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4858)
}
__typedArg0 := tmp4858
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4859)
}
__typedArg0 := tmp4859
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfreeze, tmp4860)
}
__typedArg0 := symfreeze
__typedArg1 := tmp4860
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4783 Obj

if True == tmp4861 {
tmp4852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4852)
}
__typedArg0 := tmp4852
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4853)
}
__typedArg0 := tmp4853
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4854)
}
__typedArg0 := tmp4854
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4855)
}
__typedArg0 := tmp4855
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4785 Obj

if True == tmp4856 {
tmp4846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4846)
}
__typedArg0 := tmp4846
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4847)
}
__typedArg0 := tmp4847
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4848)
}
__typedArg0 := tmp4848
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4849)
}
__typedArg0 := tmp4849
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4850)
}
__typedArg0 := tmp4850
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4787 Obj

if True == tmp4851 {
tmp4839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4839)
}
__typedArg0 := tmp4839
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4840)
}
__typedArg0 := tmp4840
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4841)
}
__typedArg0 := tmp4841
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4842)
}
__typedArg0 := tmp4842
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4843)
}
__typedArg0 := tmp4843
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symthaw, tmp4844)
}
__typedArg0 := symthaw
__typedArg1 := tmp4844
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4789 Obj

if True == tmp4845 {
tmp4832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4832)
}
__typedArg0 := tmp4832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4833)
}
__typedArg0 := tmp4833
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4834)
}
__typedArg0 := tmp4834
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4835)
}
__typedArg0 := tmp4835
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4836)
}
__typedArg0 := tmp4836
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4837)
}
__typedArg0 := tmp4837
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4791 Obj

if True == tmp4838 {
tmp4824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4824)
}
__typedArg0 := tmp4824
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4825)
}
__typedArg0 := tmp4825
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4826)
}
__typedArg0 := tmp4826
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4827)
}
__typedArg0 := tmp4827
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4828)
}
__typedArg0 := tmp4828
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4829)
}
__typedArg0 := tmp4829
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4830)
}
__typedArg0 := Nil
__typedArg1 := tmp4830
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4793 Obj

if True == tmp4831 {
tmp4818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4818)
}
__typedArg0 := tmp4818
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4819)
}
__typedArg0 := tmp4819
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4820)
}
__typedArg0 := tmp4820
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4821)
}
__typedArg0 := tmp4821
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4822)
}
__typedArg0 := Nil
__typedArg1 := tmp4822
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4795 Obj

if True == tmp4823 {
tmp4814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4814)
}
__typedArg0 := tmp4814
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4815)
}
__typedArg0 := tmp4815
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4816)
}
__typedArg0 := tmp4816
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4797 Obj

if True == tmp4817 {
tmp4809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4809)
}
__typedArg0 := tmp4809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4810)
}
__typedArg0 := tmp4810
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4811)
}
__typedArg0 := tmp4811
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4812)
}
__typedArg0 := Nil
__typedArg1 := tmp4812
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4799 Obj

if True == tmp4813 {
tmp4801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4801)
}
__typedArg0 := tmp4801
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4802)
}
__typedArg0 := tmp4802
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4803)
}
__typedArg0 := tmp4803
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4804)
}
__typedArg0 := tmp4804
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4805)
}
__typedArg0 := tmp4805
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4806)
}
__typedArg0 := tmp4806
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(tmp4807)
}
__typedArg0 := tmp4807
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

var ifres4800 Obj

if True == tmp4808 {
ifres4800 = True


} else {
ifres4800 = False


}

ifres4799 = ifres4800


} else {
ifres4799 = False


}

var ifres4798 Obj

if True == ifres4799 {
ifres4798 = True


} else {
ifres4798 = False


}

ifres4797 = ifres4798


} else {
ifres4797 = False


}

var ifres4796 Obj

if True == ifres4797 {
ifres4796 = True


} else {
ifres4796 = False


}

ifres4795 = ifres4796


} else {
ifres4795 = False


}

var ifres4794 Obj

if True == ifres4795 {
ifres4794 = True


} else {
ifres4794 = False


}

ifres4793 = ifres4794


} else {
ifres4793 = False


}

var ifres4792 Obj

if True == ifres4793 {
ifres4792 = True


} else {
ifres4792 = False


}

ifres4791 = ifres4792


} else {
ifres4791 = False


}

var ifres4790 Obj

if True == ifres4791 {
ifres4790 = True


} else {
ifres4790 = False


}

ifres4789 = ifres4790


} else {
ifres4789 = False


}

var ifres4788 Obj

if True == ifres4789 {
ifres4788 = True


} else {
ifres4788 = False


}

ifres4787 = ifres4788


} else {
ifres4787 = False


}

var ifres4786 Obj

if True == ifres4787 {
ifres4786 = True


} else {
ifres4786 = False


}

ifres4785 = ifres4786


} else {
ifres4785 = False


}

var ifres4784 Obj

if True == ifres4785 {
ifres4784 = True


} else {
ifres4784 = False


}

ifres4783 = ifres4784


} else {
ifres4783 = False


}

var ifres4782 Obj

if True == ifres4783 {
ifres4782 = True


} else {
ifres4782 = False


}

ifres4781 = ifres4782


} else {
ifres4781 = False


}

var ifres4780 Obj

if True == ifres4781 {
ifres4780 = True


} else {
ifres4780 = False


}

ifres4779 = ifres4780


} else {
ifres4779 = False


}

var ifres4778 Obj

if True == ifres4779 {
ifres4778 = True


} else {
ifres4778 = False


}

ifres4777 = ifres4778


} else {
ifres4777 = False


}

var ifres4776 Obj

if True == ifres4777 {
ifres4776 = True


} else {
ifres4776 = False


}

ifres4775 = ifres4776


} else {
ifres4775 = False


}

var ifres4774 Obj

if True == ifres4775 {
ifres4774 = True


} else {
ifres4774 = False


}

ifres4773 = ifres4774


} else {
ifres4773 = False


}

if True == ifres4773 {
tmp4759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4759)
}
__typedArg0 := tmp4759
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4760)
}
__typedArg0 := tmp4760
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4761)
}
__typedArg0 := tmp4761
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4762)
}
__typedArg0 := tmp4762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4763)
}
__typedArg0 := tmp4763
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4764)
}
__typedArg0 := tmp4764
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4766)
}
__typedArg0 := tmp4766
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4768)
}
__typedArg0 := tmp4768
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4769)
}
__typedArg0 := tmp4769
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4770)
}
__typedArg0 := tmp4770
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symsubst), tmp4765, tmp4767, tmp4771)
return


} else {
__e.Return(V897)
return
}


}, 1)

tmp4874 := Call(__e, ns2_1set, symshen_4remove_1indirection, tmp4758)


_ = tmp4874

tmp4875 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__self3 := __e.Get(3)
__selftop:
V900 := __self1
_ = V900
V901 := __self2
_ = V901
V902 := __self3
_ = V902
tmp4974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4920 Obj

if True == tmp4974 {
tmp4972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4972)
}
__typedArg0 := tmp4972
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4922 Obj

if True == tmp4973 {
tmp4969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4969)
}
__typedArg0 := tmp4969
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4970)
}
__typedArg0 := tmp4970
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4924 Obj

if True == tmp4971 {
tmp4965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4965)
}
__typedArg0 := tmp4965
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4966)
}
__typedArg0 := tmp4966
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symand, tmp4967)
}
__typedArg0 := symand
__typedArg1 := tmp4967
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4926 Obj

if True == tmp4968 {
tmp4961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4961)
}
__typedArg0 := tmp4961
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4962)
}
__typedArg0 := tmp4962
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4963)
}
__typedArg0 := tmp4963
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4928 Obj

if True == tmp4964 {
tmp4956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4956)
}
__typedArg0 := tmp4956
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4957)
}
__typedArg0 := tmp4957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4958)
}
__typedArg0 := tmp4958
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4959)
}
__typedArg0 := tmp4959
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4930 Obj

if True == tmp4960 {
tmp4950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4950)
}
__typedArg0 := tmp4950
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4951)
}
__typedArg0 := tmp4951
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4952)
}
__typedArg0 := tmp4952
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4953)
}
__typedArg0 := tmp4953
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4954)
}
__typedArg0 := Nil
__typedArg1 := tmp4954
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4932 Obj

if True == tmp4955 {
tmp4947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4947)
}
__typedArg0 := tmp4947
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4948)
}
__typedArg0 := tmp4948
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4934 Obj

if True == tmp4949 {
tmp4943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4943)
}
__typedArg0 := tmp4943
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4944)
}
__typedArg0 := tmp4944
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4945)
}
__typedArg0 := Nil
__typedArg1 := tmp4945
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4936 Obj

if True == tmp4946 {
tmp4938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4938)
}
__typedArg0 := tmp4938
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4939)
}
__typedArg0 := tmp4939
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4940)
}
__typedArg0 := tmp4940
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V900, tmp4941)
}
__typedArg0 := V900
__typedArg1 := tmp4941
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4937 Obj

if True == tmp4942 {
ifres4937 = True


} else {
ifres4937 = False


}

ifres4936 = ifres4937


} else {
ifres4936 = False


}

var ifres4935 Obj

if True == ifres4936 {
ifres4935 = True


} else {
ifres4935 = False


}

ifres4934 = ifres4935


} else {
ifres4934 = False


}

var ifres4933 Obj

if True == ifres4934 {
ifres4933 = True


} else {
ifres4933 = False


}

ifres4932 = ifres4933


} else {
ifres4932 = False


}

var ifres4931 Obj

if True == ifres4932 {
ifres4931 = True


} else {
ifres4931 = False


}

ifres4930 = ifres4931


} else {
ifres4930 = False


}

var ifres4929 Obj

if True == ifres4930 {
ifres4929 = True


} else {
ifres4929 = False


}

ifres4928 = ifres4929


} else {
ifres4928 = False


}

var ifres4927 Obj

if True == ifres4928 {
ifres4927 = True


} else {
ifres4927 = False


}

ifres4926 = ifres4927


} else {
ifres4926 = False


}

var ifres4925 Obj

if True == ifres4926 {
ifres4925 = True


} else {
ifres4925 = False


}

ifres4924 = ifres4925


} else {
ifres4924 = False


}

var ifres4923 Obj

if True == ifres4924 {
ifres4923 = True


} else {
ifres4923 = False


}

ifres4922 = ifres4923


} else {
ifres4922 = False


}

var ifres4921 Obj

if True == ifres4922 {
ifres4921 = True


} else {
ifres4921 = False


}

ifres4920 = ifres4921


} else {
ifres4920 = False


}

if True == ifres4920 {
tmp4876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4876)
}
__typedArg0 := tmp4876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4877)
}
__typedArg0 := tmp4877
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4878)
}
__typedArg0 := tmp4878
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4881)
}
__typedArg0 := tmp4881
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4882)
}
__typedArg0 := tmp4882
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4883)
}
__typedArg0 := tmp4883
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4884)
}
__typedArg0 := tmp4884
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4886)
}
__typedArg0 := tmp4886
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4885, tmp4887)
}
__typedArg0 := tmp4885
__typedArg1 := tmp4887
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4888, V902)
}
__typedArg0 := tmp4888
__typedArg1 := V902
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4pivot_1on) == __self {
__self1, __self2, __self3 = tmp4879, tmp4880, tmp4889
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4pivot_1on), tmp4879, tmp4880, tmp4889)
return


} else {
tmp4918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4898 Obj

if True == tmp4918 {
tmp4916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4916)
}
__typedArg0 := tmp4916
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4900 Obj

if True == tmp4917 {
tmp4913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4913)
}
__typedArg0 := tmp4913
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4914)
}
__typedArg0 := tmp4914
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4902 Obj

if True == tmp4915 {
tmp4909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4909)
}
__typedArg0 := tmp4909
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4910)
}
__typedArg0 := tmp4910
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4911)
}
__typedArg0 := Nil
__typedArg1 := tmp4911
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4904 Obj

if True == tmp4912 {
tmp4906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4906)
}
__typedArg0 := tmp4906
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V900, tmp4907)
}
__typedArg0 := V900
__typedArg1 := tmp4907
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4905 Obj

if True == tmp4908 {
ifres4905 = True


} else {
ifres4905 = False


}

ifres4904 = ifres4905


} else {
ifres4904 = False


}

var ifres4903 Obj

if True == ifres4904 {
ifres4903 = True


} else {
ifres4903 = False


}

ifres4902 = ifres4903


} else {
ifres4902 = False


}

var ifres4901 Obj

if True == ifres4902 {
ifres4901 = True


} else {
ifres4901 = False


}

ifres4900 = ifres4901


} else {
ifres4900 = False


}

var ifres4899 Obj

if True == ifres4900 {
ifres4899 = True


} else {
ifres4899 = False


}

ifres4898 = ifres4899


} else {
ifres4898 = False


}

if True == ifres4898 {
tmp4890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4890)
}
__typedArg0 := tmp4890
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4893)
}
__typedArg0 := tmp4893
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp4894)
}
__typedArg0 := True
__typedArg1 := tmp4894
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4895, V902)
}
__typedArg0 := tmp4895
__typedArg1 := V902
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4pivot_1on) == __self {
__self1, __self2, __self3 = tmp4891, tmp4892, tmp4896
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4pivot_1on), tmp4891, tmp4892, tmp4896)
return


} else {
__e.TailApply(PrimFunc(sym_8p), V902, V901)
return
}


}


}, 3)

tmp4975 := Call(__e, ns2_1set, symshen_4pivot_1on, tmp4875)


_ = tmp4975

tmp4976 := MakeNative(func(__e *ControlFlow) {
V905 := __e.Get(1)
_ = V905
V906 := __e.Get(2)
_ = V906
tmp4999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V905)
}
__typedArg0 := V905
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4990 Obj

if True == tmp4999 {
tmp4997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V905)
}
__typedArg0 := V905
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4997)
}
__typedArg0 := tmp4997
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4992 Obj

if True == tmp4998 {
tmp4994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V905)
}
__typedArg0 := V905
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4994)
}
__typedArg0 := tmp4994
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4995)
}
__typedArg0 := Nil
__typedArg1 := tmp4995
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4993 Obj

if True == tmp4996 {
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
tmp4977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V905)
}
__typedArg0 := V905
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4978 := Call(__e, PrimFunc(symshen_4op), tmp4977)


let__2748 := tmp4978
_ = let__2748

tmp4988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4skip, let__2748)
}
__typedArg0 := symshen_4skip
__typedArg1 := let__2748
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4988 {
__e.Return(V906)
return
} else {
tmp4979 := Call(__e, PrimFunc(symshen_4op1), let__2748)


tmp4980 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V905)
}
__typedArg0 := V905
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4981 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4979, tmp4980)
}
__typedArg0 := tmp4979
__typedArg1 := tmp4980
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4982 := Call(__e, PrimFunc(symshen_4op2), let__2748)


tmp4983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V905)
}
__typedArg0 := V905
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4982, tmp4983)
}
__typedArg0 := tmp4982
__typedArg1 := tmp4983
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4984, Nil)
}
__typedArg0 := tmp4984
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4981, tmp4985)
}
__typedArg0 := tmp4981
__typedArg1 := tmp4985
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4factor_1selectors_1h), tmp4986, V906)
return


}


} else {
__e.Return(V906)
return
}


}, 2)

tmp5000 := Call(__e, ns2_1set, symshen_4factor_1selectors, tmp4976)


_ = tmp5000

tmp5001 := MakeNative(func(__e *ControlFlow) {
V910 := __e.Get(1)
_ = V910
tmp5009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons_2, V910)
}
__typedArg0 := symcons_2
__typedArg1 := V910
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5009 {
__e.Return(symcons)
return
} else {
tmp5007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7string_2, V910)
}
__typedArg0 := symshen_4_7string_2
__typedArg1 := V910
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5007 {
__e.Return(sym_8s)
return
} else {
tmp5005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7vector_2, V910)
}
__typedArg0 := symshen_4_7vector_2
__typedArg1 := V910
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5005 {
__e.Return(sym_8v)
return
} else {
tmp5003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtuple_2, V910)
}
__typedArg0 := symtuple_2
__typedArg1 := V910
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5003 {
__e.Return(sym_8p)
return
} else {
__e.Return(symshen_4skip)
return
}


}


}


}


}, 1)

tmp5010 := Call(__e, ns2_1set, symshen_4op, tmp5001)


_ = tmp5010

tmp5011 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V911 := __self1
_ = V911
V912 := __self2
_ = V912
tmp5029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V911)
}
__typedArg0 := Nil
__typedArg1 := V911
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5029 {
__e.Return(V912)
return
} else {
tmp5027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V911)
}
__typedArg0 := V911
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5027 {
tmp5023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V911)
}
__typedArg0 := V911
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5024 := Call(__e, PrimFunc(symoccurrences), tmp5023, V912)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(tmp5024)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := tmp5024
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
tmp5012 := Call(__e, PrimFunc(symgensym), symSelect)


let__2749 := tmp5012
_ = let__2749

tmp5013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V911)
}
__typedArg0 := V911
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V911)
}
__typedArg0 := V911
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp5015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V911)
}
__typedArg0 := V911
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5016 := Call(__e, PrimFunc(symsubst), let__2749, tmp5015, V912)


tmp5017 := Call(__e, PrimFunc(symshen_4factor_1selectors_1h), tmp5014, tmp5016)


tmp5018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5017, Nil)
}
__typedArg0 := tmp5017
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp5013, tmp5018)
}
__typedArg0 := tmp5013
__typedArg1 := tmp5018
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2749, tmp5019)
}
__typedArg0 := let__2749
__typedArg1 := tmp5019
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp5020)
}
__typedArg0 := symlet
__typedArg1 := tmp5020
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp5021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V911)
}
__typedArg0 := V911
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4factor_1selectors_1h) == __self {
__self1, __self2 = tmp5021, V912
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4factor_1selectors_1h), tmp5021, V912)
return


}


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4factor_1selectors_1h)
return
}


}


}, 2)

__e.TailApply(ns2_1set, symshen_4factor_1selectors_1h, tmp5011)
return




}, 0)

