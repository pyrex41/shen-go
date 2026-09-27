package main

import . "github.com/pyrex41/shen-go/kl"

var CoreMain = MakeNative(func(__e *ControlFlow) {
tmp2771 := MakeNative(func(__e *ControlFlow) {
V528 := __e.Get(1)
_ = V528
tmp2772 := Call(__e, PrimFunc(symshen_4shen_1_6kl_1h), V528)


let__2610 := tmp2772
_ = let__2610

__e.TailApply(PrimFunc(symshen_4record_1and_1evaluate), let__2610)
return


}, 1)

tmp2773 := Call(__e, ns2_1set, symshen_4shen_1_6kl, tmp2771)


_ = tmp2773

tmp2774 := MakeNative(func(__e *ControlFlow) {
V530 := __e.Get(1)
_ = V530
tmp2823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V530)
}
__typedArg0 := V530
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2797 Obj

if True == tmp2823 {
tmp2821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V530)
}
__typedArg0 := V530
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp2821)
}
__typedArg0 := symdefun
__typedArg1 := tmp2821
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2799 Obj

if True == tmp2822 {
tmp2819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V530)
}
__typedArg0 := V530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2819)
}
__typedArg0 := tmp2819
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2801 Obj

if True == tmp2820 {
tmp2816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V530)
}
__typedArg0 := V530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2816)
}
__typedArg0 := tmp2816
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2817)
}
__typedArg0 := tmp2817
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2803 Obj

if True == tmp2818 {
tmp2812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V530)
}
__typedArg0 := V530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2812)
}
__typedArg0 := tmp2812
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2813)
}
__typedArg0 := tmp2813
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2814)
}
__typedArg0 := tmp2814
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2805 Obj

if True == tmp2815 {
tmp2807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V530)
}
__typedArg0 := V530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2807)
}
__typedArg0 := tmp2807
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2808)
}
__typedArg0 := tmp2808
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2809)
}
__typedArg0 := tmp2809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2810)
}
__typedArg0 := Nil
__typedArg1 := tmp2810
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2806 Obj

if True == tmp2811 {
ifres2806 = True


} else {
ifres2806 = False


}

ifres2805 = ifres2806


} else {
ifres2805 = False


}

var ifres2804 Obj

if True == ifres2805 {
ifres2804 = True


} else {
ifres2804 = False


}

ifres2803 = ifres2804


} else {
ifres2803 = False


}

var ifres2802 Obj

if True == ifres2803 {
ifres2802 = True


} else {
ifres2802 = False


}

ifres2801 = ifres2802


} else {
ifres2801 = False


}

var ifres2800 Obj

if True == ifres2801 {
ifres2800 = True


} else {
ifres2800 = False


}

ifres2799 = ifres2800


} else {
ifres2799 = False


}

var ifres2798 Obj

if True == ifres2799 {
ifres2798 = True


} else {
ifres2798 = False


}

ifres2797 = ifres2798


} else {
ifres2797 = False


}

if True == ifres2797 {
tmp2780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V530)
}
__typedArg0 := V530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2780)
}
__typedArg0 := tmp2780
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2782 := Call(__e, PrimFunc(symshen_4sysfunc_2), tmp2781)


var ifres2775 Obj

if True == tmp2782 {
tmp2776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V530)
}
__typedArg0 := V530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2776)
}
__typedArg0 := tmp2776
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2778 := Call(__e, PrimFunc(symshen_4app), tmp2777, MakeString(" is not a legitimate function name\n"), symshen_4a)


tmp2779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp2778)
}
__typedArg0 := tmp2778
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres2775 = tmp2779


} else {
ifres2775 = symshen_4skip


}

let__2611 := ifres2775
_ = let__2611

tmp2783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V530)
}
__typedArg0 := V530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2783)
}
__typedArg0 := tmp2783
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V530)
}
__typedArg0 := V530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2785)
}
__typedArg0 := tmp2785
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2786)
}
__typedArg0 := tmp2786
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2788 := Call(__e, PrimFunc(symlength), tmp2787)


tmp2789 := Call(__e, PrimFunc(symshen_4store_1arity), tmp2784, tmp2788)


let__2612 := tmp2789
_ = let__2612

tmp2790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V530)
}
__typedArg0 := V530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2790)
}
__typedArg0 := tmp2790
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2792 := Call(__e, PrimFunc(symshen_4record_1kl), tmp2791, V530)


let__2613 := tmp2792
_ = let__2613

tmp2793 := Call(__e, PrimFunc(symeval_1kl), V530)


let__2614 := tmp2793
_ = let__2614

tmp2794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V530)
}
__typedArg0 := V530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2794)
}
__typedArg0 := tmp2794
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4fn_1print), tmp2795)
return


} else {
__e.Return(V530)
return
}


}, 1)

tmp2824 := Call(__e, ns2_1set, symshen_4record_1and_1evaluate, tmp2774)


_ = tmp2824

tmp2825 := MakeNative(func(__e *ControlFlow) {
V535 := __e.Get(1)
_ = V535
tmp2926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2918 Obj

if True == tmp2926 {
tmp2924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp2924)
}
__typedArg0 := symdefine
__typedArg1 := tmp2924
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2920 Obj

if True == tmp2925 {
tmp2922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2922)
}
__typedArg0 := tmp2922
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2921 Obj

if True == tmp2923 {
ifres2921 = True


} else {
ifres2921 = False


}

ifres2920 = ifres2921


} else {
ifres2920 = False


}

var ifres2919 Obj

if True == ifres2920 {
ifres2919 = True


} else {
ifres2919 = False


}

ifres2918 = ifres2919


} else {
ifres2918 = False


}

if True == ifres2918 {
tmp2826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2826)
}
__typedArg0 := tmp2826
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2828)
}
__typedArg0 := tmp2828
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4shendef_1_6kldef), tmp2827, tmp2829)
return


} else {
tmp2916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2890 Obj

if True == tmp2916 {
tmp2914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp2914)
}
__typedArg0 := symdefun
__typedArg1 := tmp2914
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2892 Obj

if True == tmp2915 {
tmp2912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2912)
}
__typedArg0 := tmp2912
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2894 Obj

if True == tmp2913 {
tmp2909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2909)
}
__typedArg0 := tmp2909
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2910)
}
__typedArg0 := tmp2910
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2896 Obj

if True == tmp2911 {
tmp2905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2905)
}
__typedArg0 := tmp2905
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2906)
}
__typedArg0 := tmp2906
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2907)
}
__typedArg0 := tmp2907
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2898 Obj

if True == tmp2908 {
tmp2900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2900)
}
__typedArg0 := tmp2900
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2901)
}
__typedArg0 := tmp2901
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2902)
}
__typedArg0 := tmp2902
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2903)
}
__typedArg0 := Nil
__typedArg1 := tmp2903
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2899 Obj

if True == tmp2904 {
ifres2899 = True


} else {
ifres2899 = False


}

ifres2898 = ifres2899


} else {
ifres2898 = False


}

var ifres2897 Obj

if True == ifres2898 {
ifres2897 = True


} else {
ifres2897 = False


}

ifres2896 = ifres2897


} else {
ifres2896 = False


}

var ifres2895 Obj

if True == ifres2896 {
ifres2895 = True


} else {
ifres2895 = False


}

ifres2894 = ifres2895


} else {
ifres2894 = False


}

var ifres2893 Obj

if True == ifres2894 {
ifres2893 = True


} else {
ifres2893 = False


}

ifres2892 = ifres2893


} else {
ifres2892 = False


}

var ifres2891 Obj

if True == ifres2892 {
ifres2891 = True


} else {
ifres2891 = False


}

ifres2890 = ifres2891


} else {
ifres2890 = False


}

if True == ifres2890 {
__e.Return(V535)
return
} else {
tmp2888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2869 Obj

if True == tmp2888 {
tmp2886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtype, tmp2886)
}
__typedArg0 := symtype
__typedArg1 := tmp2886
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2871 Obj

if True == tmp2887 {
tmp2884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2884)
}
__typedArg0 := tmp2884
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2873 Obj

if True == tmp2885 {
tmp2881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2882)
}
__typedArg0 := tmp2882
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2875 Obj

if True == tmp2883 {
tmp2877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2877)
}
__typedArg0 := tmp2877
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2878)
}
__typedArg0 := tmp2878
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2879)
}
__typedArg0 := Nil
__typedArg1 := tmp2879
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2876 Obj

if True == tmp2880 {
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
tmp2830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2830)
}
__typedArg0 := tmp2830
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2832)
}
__typedArg0 := tmp2832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2833)
}
__typedArg0 := tmp2833
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2835 := Call(__e, PrimFunc(symshen_4rcons__form), tmp2834)


tmp2836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2835, Nil)
}
__typedArg0 := tmp2835
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2831, tmp2836)
}
__typedArg0 := tmp2831
__typedArg1 := tmp2836
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp2837)
}
__typedArg0 := symtype
__typedArg1 := tmp2837
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp2867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2848 Obj

if True == tmp2867 {
tmp2865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput_7, tmp2865)
}
__typedArg0 := syminput_7
__typedArg1 := tmp2865
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2850 Obj

if True == tmp2866 {
tmp2863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
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
return PrimTail(V535)
}
__typedArg0 := V535
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
return PrimTail(V535)
}
__typedArg0 := V535
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
tmp2838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2838)
}
__typedArg0 := tmp2838
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2840 := Call(__e, PrimFunc(symshen_4rcons__form), tmp2839)


tmp2841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2841)
}
__typedArg0 := tmp2841
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2840, tmp2842)
}
__typedArg0 := tmp2840
__typedArg1 := tmp2842
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput_7, tmp2843)
}
__typedArg0 := syminput_7
__typedArg1 := tmp2843
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp2846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2846 {
tmp2844 := MakeNative(func(__e *ControlFlow) {
Z536 := __e.Get(1)
_ = Z536
__e.TailApply(PrimFunc(symshen_4shen_1_6kl_1h), Z536)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp2844, V535)
return


} else {
__e.Return(V535)
return
}


}


}


}


}


}, 1)

tmp2927 := Call(__e, ns2_1set, symshen_4shen_1_6kl_1h, tmp2825)


_ = tmp2927

tmp2928 := MakeNative(func(__e *ControlFlow) {
V537 := __e.Get(1)
_ = V537
V538 := __e.Get(2)
_ = V538
tmp2929 := MakeNative(func(__e *ControlFlow) {
Z539 := __e.Get(1)
_ = Z539
__e.TailApply(PrimFunc(symshen_4_5define_6), Z539)
return
}, 1)

tmp2930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V537, V538)
}
__typedArg0 := V537
__typedArg1 := V538
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symcompile), tmp2929, tmp2930)
return


}, 2)

tmp2931 := Call(__e, ns2_1set, symshen_4shendef_1_6kldef, tmp2928)


_ = tmp2931

tmp2932 := MakeNative(func(__e *ControlFlow) {
V540 := __e.Get(1)
_ = V540
tmp2933 := Call(__e, PrimFunc(symshen_4_5name_6), V540)


let__2616 := tmp2933
_ = let__2616

tmp2959 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2616)


var ifres2934 Obj

if True == tmp2959 {
tmp2935 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2934 = tmp2935


} else {
tmp2936 := Call(__e, PrimFunc(symshen_4_5_1out), let__2616)


let__2617 := tmp2936
_ = let__2617

tmp2937 := Call(__e, PrimFunc(symshen_4in_1_6), let__2616)


let__2618 := tmp2937
_ = let__2618

tmp2958 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2618, sym_i)


var ifres2938 Obj

if True == tmp2958 {
tmp2939 := Call(__e, PrimFunc(symtail), let__2618)


let__2619 := tmp2939
_ = let__2619

tmp2940 := Call(__e, PrimFunc(symshen_4_5signature_6), let__2619)


let__2620 := tmp2940
_ = let__2620

tmp2956 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2620)


var ifres2941 Obj

if True == tmp2956 {
tmp2942 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2941 = tmp2942


} else {
tmp2943 := Call(__e, PrimFunc(symshen_4in_1_6), let__2620)


let__2621 := tmp2943
_ = let__2621

tmp2955 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2621, sym_j)


var ifres2944 Obj

if True == tmp2955 {
tmp2945 := Call(__e, PrimFunc(symtail), let__2621)


let__2622 := tmp2945
_ = let__2622

tmp2946 := Call(__e, PrimFunc(symshen_4_5rules_6), let__2622)


let__2623 := tmp2946
_ = let__2623

tmp2953 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2623)


var ifres2947 Obj

if True == tmp2953 {
tmp2948 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2947 = tmp2948


} else {
tmp2949 := Call(__e, PrimFunc(symshen_4_5_1out), let__2623)


let__2624 := tmp2949
_ = let__2624

tmp2950 := Call(__e, PrimFunc(symshen_4in_1_6), let__2623)


let__2625 := tmp2950
_ = let__2625

tmp2951 := Call(__e, PrimFunc(symshen_4shendef_1_6kldef_1h), let__2617, let__2624)


tmp2952 := Call(__e, PrimFunc(symshen_4comb), let__2625, tmp2951)


ifres2947 = tmp2952


}

ifres2944 = ifres2947


} else {
tmp2954 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2944 = tmp2954


}

ifres2941 = ifres2944


}

ifres2938 = ifres2941


} else {
tmp2957 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2938 = tmp2957


}

ifres2934 = ifres2938


}

let__2615 := ifres2934
_ = let__2615

tmp2977 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2615)


if True == tmp2977 {
tmp2960 := Call(__e, PrimFunc(symshen_4_5name_6), V540)


let__2627 := tmp2960
_ = let__2627

tmp2973 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2627)


var ifres2961 Obj

if True == tmp2973 {
tmp2962 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2961 = tmp2962


} else {
tmp2963 := Call(__e, PrimFunc(symshen_4_5_1out), let__2627)


let__2628 := tmp2963
_ = let__2628

tmp2964 := Call(__e, PrimFunc(symshen_4in_1_6), let__2627)


let__2629 := tmp2964
_ = let__2629

tmp2965 := Call(__e, PrimFunc(symshen_4_5rules_6), let__2629)


let__2630 := tmp2965
_ = let__2630

tmp2972 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2630)


var ifres2966 Obj

if True == tmp2972 {
tmp2967 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2966 = tmp2967


} else {
tmp2968 := Call(__e, PrimFunc(symshen_4_5_1out), let__2630)


let__2631 := tmp2968
_ = let__2631

tmp2969 := Call(__e, PrimFunc(symshen_4in_1_6), let__2630)


let__2632 := tmp2969
_ = let__2632

tmp2970 := Call(__e, PrimFunc(symshen_4shendef_1_6kldef_1h), let__2628, let__2631)


tmp2971 := Call(__e, PrimFunc(symshen_4comb), let__2632, tmp2970)


ifres2966 = tmp2971


}

ifres2961 = ifres2966


}

let__2626 := ifres2961
_ = let__2626

tmp2975 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2626)


if True == tmp2975 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2626)
return
}


} else {
__e.Return(let__2615)
return
}


}, 1)

tmp2978 := Call(__e, ns2_1set, symshen_4_5define_6, tmp2932)


_ = tmp2978

tmp2979 := MakeNative(func(__e *ControlFlow) {
V559 := __e.Get(1)
_ = V559
V560 := __e.Get(2)
_ = V560
tmp2980 := MakeNative(func(__e *ControlFlow) {
Z562 := __e.Get(1)
_ = Z562
__e.TailApply(PrimFunc(symfst), Z562)
return
}, 1)

tmp2981 := Call(__e, PrimFunc(symmap), tmp2980, V560)


let__2633 := tmp2981
_ = let__2633

tmp2982 := Call(__e, PrimFunc(symshen_4arity_1chk), V559, let__2633)


let__2634 := tmp2982
_ = let__2634

tmp2983 := MakeNative(func(__e *ControlFlow) {
Z565 := __e.Get(1)
_ = Z565
__e.TailApply(PrimFunc(symshen_4free_1var_1chk), V559, Z565)
return
}, 1)

tmp2984 := Call(__e, PrimFunc(symmap), tmp2983, V560)


let__2635 := tmp2984
_ = let__2635

tmp2985 := Call(__e, PrimFunc(symshen_4unprotect), V560)


let__2636 := tmp2985
_ = let__2636

tmp2986 := Call(__e, PrimFunc(symshen_4compile_1to_1kl), V559, let__2636, let__2634)


tmp2987 := Call(__e, PrimFunc(symshen_4factorise_1code), tmp2986)


let__2637 := tmp2987
_ = let__2637

__e.Return(let__2637)
return


}, 2)

tmp2988 := Call(__e, ns2_1set, symshen_4shendef_1_6kldef_1h, tmp2979)


_ = tmp2988

tmp2989 := MakeNative(func(__e *ControlFlow) {
V568 := __e.Get(1)
_ = V568
tmp3015 := Call(__e, PrimFunc(symtuple_2), V568)


if True == tmp3015 {
tmp2990 := Call(__e, PrimFunc(symfst), V568)


tmp2991 := Call(__e, PrimFunc(symshen_4unprotect), tmp2990)


tmp2992 := Call(__e, PrimFunc(symsnd), V568)


tmp2993 := Call(__e, PrimFunc(symshen_4unprotect), tmp2992)


__e.TailApply(PrimFunc(sym_8p), tmp2991, tmp2993)
return


} else {
tmp3013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V568)
}
__typedArg0 := V568
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3000 Obj

if True == tmp3013 {
tmp3011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V568)
}
__typedArg0 := V568
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symprotect, tmp3011)
}
__typedArg0 := symprotect
__typedArg1 := tmp3011
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3002 Obj

if True == tmp3012 {
tmp3009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V568)
}
__typedArg0 := V568
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3009)
}
__typedArg0 := tmp3009
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3004 Obj

if True == tmp3010 {
tmp3006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V568)
}
__typedArg0 := V568
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3006)
}
__typedArg0 := tmp3006
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3007)
}
__typedArg0 := Nil
__typedArg1 := tmp3007
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3005 Obj

if True == tmp3008 {
ifres3005 = True


} else {
ifres3005 = False


}

ifres3004 = ifres3005


} else {
ifres3004 = False


}

var ifres3003 Obj

if True == ifres3004 {
ifres3003 = True


} else {
ifres3003 = False


}

ifres3002 = ifres3003


} else {
ifres3002 = False


}

var ifres3001 Obj

if True == ifres3002 {
ifres3001 = True


} else {
ifres3001 = False


}

ifres3000 = ifres3001


} else {
ifres3000 = False


}

if True == ifres3000 {
tmp2994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V568)
}
__typedArg0 := V568
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2994)
}
__typedArg0 := tmp2994
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4unprotect), tmp2995)
return


} else {
tmp2998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V568)
}
__typedArg0 := V568
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2998 {
tmp2996 := MakeNative(func(__e *ControlFlow) {
Z569 := __e.Get(1)
_ = Z569
__e.TailApply(PrimFunc(symshen_4unprotect), Z569)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp2996, V568)
return


} else {
__e.Return(V568)
return
}


}


}


}, 1)

tmp3016 := Call(__e, ns2_1set, symshen_4unprotect, tmp2989)


_ = tmp3016

tmp3017 := MakeNative(func(__e *ControlFlow) {
V570 := __e.Get(1)
_ = V570
tmp3031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V570)
}
__typedArg0 := V570
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3018 Obj

if True == tmp3031 {
tmp3019 := Call(__e, PrimFunc(symhead), V570)


let__2639 := tmp3019
_ = let__2639

tmp3020 := Call(__e, PrimFunc(symtail), V570)


let__2640 := tmp3020
_ = let__2640

tmp3028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(let__2639)
}
__typedArg0 := let__2639
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

var ifres3024 Obj

if True == tmp3028 {
tmp3026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(let__2639)
}
__typedArg0 := let__2639
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

tmp3027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp3026)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp3026
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres3025 Obj

if True == tmp3027 {
ifres3025 = True


} else {
ifres3025 = False


}

ifres3024 = ifres3025


} else {
ifres3024 = False


}

var ifres3021 Obj

if True == ifres3024 {
ifres3021 = let__2639


} else {
tmp3022 := Call(__e, PrimFunc(symshen_4app), let__2639, MakeString(" is not a legitimate function name.\n"), symshen_4a)


tmp3023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp3022)
}
__typedArg0 := tmp3022
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres3021 = tmp3023


}

tmp3029 := Call(__e, PrimFunc(symshen_4comb), let__2640, ifres3021)


ifres3018 = tmp3029


} else {
tmp3030 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3018 = tmp3030


}

let__2638 := ifres3018
_ = let__2638

tmp3033 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2638)


if True == tmp3033 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2638)
return
}


}, 1)

tmp3034 := Call(__e, ns2_1set, symshen_4_5name_6, tmp3017)


_ = tmp3034

tmp3035 := MakeNative(func(__e *ControlFlow) {
V574 := __e.Get(1)
_ = V574
tmp3054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V574)
}
__typedArg0 := V574
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3036 Obj

if True == tmp3054 {
tmp3037 := Call(__e, PrimFunc(symhead), V574)


let__2642 := tmp3037
_ = let__2642

tmp3038 := Call(__e, PrimFunc(symtail), V574)


let__2643 := tmp3038
_ = let__2643

tmp3039 := Call(__e, PrimFunc(symshen_4_5signature_6), let__2643)


let__2644 := tmp3039
_ = let__2644

tmp3052 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2644)


var ifres3040 Obj

if True == tmp3052 {
tmp3041 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3040 = tmp3041


} else {
tmp3042 := Call(__e, PrimFunc(symshen_4_5_1out), let__2644)


let__2645 := tmp3042
_ = let__2645

tmp3043 := Call(__e, PrimFunc(symshen_4in_1_6), let__2644)


let__2646 := tmp3043
_ = let__2646

tmp3048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, Nil)
}
__typedArg0 := sym_j
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, tmp3048)
}
__typedArg0 := sym_i
__typedArg1 := tmp3048
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3050 := Call(__e, PrimFunc(symelement_2), let__2642, tmp3049)


tmp3051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp3050)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp3050
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres3044 Obj

if True == tmp3051 {
tmp3045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2642, let__2645)
}
__typedArg0 := let__2642
__typedArg1 := let__2645
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3046 := Call(__e, PrimFunc(symshen_4comb), let__2646, tmp3045)


ifres3044 = tmp3046


} else {
tmp3047 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3044 = tmp3047


}

ifres3040 = ifres3044


}

ifres3036 = ifres3040


} else {
tmp3053 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3036 = tmp3053


}

let__2641 := ifres3036
_ = let__2641

tmp3064 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2641)


if True == tmp3064 {
tmp3055 := Call(__e, PrimFunc(sym_5e_6), V574)


let__2648 := tmp3055
_ = let__2648

tmp3060 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2648)


var ifres3056 Obj

if True == tmp3060 {
tmp3057 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3056 = tmp3057


} else {
tmp3058 := Call(__e, PrimFunc(symshen_4in_1_6), let__2648)


let__2649 := tmp3058
_ = let__2649

tmp3059 := Call(__e, PrimFunc(symshen_4comb), let__2649, Nil)


ifres3056 = tmp3059


}

let__2647 := ifres3056
_ = let__2647

tmp3062 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2647)


if True == tmp3062 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2647)
return
}


} else {
__e.Return(let__2641)
return
}


}, 1)

tmp3065 := Call(__e, ns2_1set, symshen_4_5signature_6, tmp3035)


_ = tmp3065

tmp3066 := MakeNative(func(__e *ControlFlow) {
V584 := __e.Get(1)
_ = V584
tmp3067 := Call(__e, PrimFunc(symshen_4_5rule_6), V584)


let__2651 := tmp3067
_ = let__2651

tmp3081 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2651)


var ifres3068 Obj

if True == tmp3081 {
tmp3069 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3068 = tmp3069


} else {
tmp3070 := Call(__e, PrimFunc(symshen_4_5_1out), let__2651)


let__2652 := tmp3070
_ = let__2652

tmp3071 := Call(__e, PrimFunc(symshen_4in_1_6), let__2651)


let__2653 := tmp3071
_ = let__2653

tmp3072 := Call(__e, PrimFunc(symshen_4_5rules_6), let__2653)


let__2654 := tmp3072
_ = let__2654

tmp3080 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2654)


var ifres3073 Obj

if True == tmp3080 {
tmp3074 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3073 = tmp3074


} else {
tmp3075 := Call(__e, PrimFunc(symshen_4_5_1out), let__2654)


let__2655 := tmp3075
_ = let__2655

tmp3076 := Call(__e, PrimFunc(symshen_4in_1_6), let__2654)


let__2656 := tmp3076
_ = let__2656

tmp3077 := Call(__e, PrimFunc(symshen_4linearise), let__2652)


tmp3078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3077, let__2655)
}
__typedArg0 := tmp3077
__typedArg1 := let__2655
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3079 := Call(__e, PrimFunc(symshen_4comb), let__2656, tmp3078)


ifres3073 = tmp3079


}

ifres3068 = ifres3073


}

let__2650 := ifres3068
_ = let__2650

tmp3097 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2650)


if True == tmp3097 {
tmp3082 := Call(__e, PrimFunc(sym_5_b_6), V584)


let__2658 := tmp3082
_ = let__2658

tmp3093 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2658)


var ifres3083 Obj

if True == tmp3093 {
tmp3084 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3083 = tmp3084


} else {
tmp3085 := Call(__e, PrimFunc(symshen_4_5_1out), let__2658)


let__2659 := tmp3085
_ = let__2659

tmp3086 := Call(__e, PrimFunc(symshen_4in_1_6), let__2658)


let__2660 := tmp3086
_ = let__2660

tmp3091 := Call(__e, PrimFunc(symempty_2), let__2659)


var ifres3087 Obj

if True == tmp3091 {
ifres3087 = Nil


} else {
tmp3088 := Call(__e, PrimFunc(symshen_4app), let__2659, MakeString("\n ..."), symshen_4r)


tmp3090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Shen syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp3088)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Shen syntax error here:\n ")
__typedArg1 := tmp3088
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Shen syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp3088)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Shen syntax error here:\n ")
__typedArg1 := tmp3088
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres3087 = tmp3090


}

tmp3092 := Call(__e, PrimFunc(symshen_4comb), let__2660, ifres3087)


ifres3083 = tmp3092


}

let__2657 := ifres3083
_ = let__2657

tmp3095 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2657)


if True == tmp3095 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2657)
return
}


} else {
__e.Return(let__2650)
return
}


}, 1)

tmp3098 := Call(__e, ns2_1set, symshen_4_5rules_6, tmp3066)


_ = tmp3098

tmp3099 := MakeNative(func(__e *ControlFlow) {
V598 := __e.Get(1)
_ = V598
tmp3104 := Call(__e, PrimFunc(symtuple_2), V598)


if True == tmp3104 {
tmp3100 := Call(__e, PrimFunc(symfst), V598)


tmp3101 := Call(__e, PrimFunc(symfst), V598)


tmp3102 := Call(__e, PrimFunc(symsnd), V598)


__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp3100, tmp3101, Nil, tmp3102)
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

tmp3105 := Call(__e, ns2_1set, symshen_4linearise, tmp3099)


_ = tmp3105

tmp3106 := MakeNative(func(__e *ControlFlow) {
V611 := __e.Get(1)
_ = V611
V612 := __e.Get(2)
_ = V612
V613 := __e.Get(3)
_ = V613
V614 := __e.Get(4)
_ = V614
tmp3143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V611)
}
__typedArg0 := Nil
__typedArg1 := V611
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp3143 {
__e.TailApply(PrimFunc(sym_8p), V612, V614)
return
} else {
tmp3141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3137 Obj

if True == tmp3141 {
tmp3139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3139)
}
__typedArg0 := tmp3139
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3138 Obj

if True == tmp3140 {
ifres3138 = True


} else {
ifres3138 = False


}

ifres3137 = ifres3138


} else {
ifres3137 = False


}

if True == ifres3137 {
tmp3107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3109 := Call(__e, PrimFunc(symappend), tmp3107, tmp3108)


__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp3109, V612, V613, V614)
return


} else {
tmp3135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3131 Obj

if True == tmp3135 {
tmp3133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3134 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp3133)
}
__typedArg0 := tmp3133
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres3132 Obj

if True == tmp3134 {
ifres3132 = True


} else {
ifres3132 = False


}

ifres3131 = ifres3132


} else {
ifres3131 = False


}

if True == ifres3131 {
tmp3125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3126 := Call(__e, PrimFunc(symelement_2), tmp3125, V613)


if True == tmp3126 {
tmp3110 := Call(__e, PrimFunc(symgensym), symV)


let__2661 := tmp3110
_ = let__2661

tmp3111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3113 := Call(__e, PrimFunc(symshen_4rep_1X), tmp3112, let__2661, V612)


tmp3114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3114, Nil)
}
__typedArg0 := tmp3114
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2661, tmp3115)
}
__typedArg0 := let__2661
__typedArg1 := tmp3115
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp3116)
}
__typedArg0 := sym_a
__typedArg1 := tmp3116
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V614, Nil)
}
__typedArg0 := V614
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3117, tmp3118)
}
__typedArg0 := tmp3117
__typedArg1 := tmp3118
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp3119)
}
__typedArg0 := symwhere
__typedArg1 := tmp3119
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp3111, tmp3113, V613, tmp3120)
return


} else {
tmp3121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3122, V613)
}
__typedArg0 := tmp3122
__typedArg1 := V613
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp3121, V612, tmp3123, V614)
return


}


} else {
tmp3129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3129 {
tmp3127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp3127, V612, V613, V614)
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

tmp3144 := Call(__e, ns2_1set, symshen_4linearise_1h, tmp3106)


_ = tmp3144

tmp3145 := MakeNative(func(__e *ControlFlow) {
V616 := __e.Get(1)
_ = V616
tmp3146 := Call(__e, PrimFunc(symshen_4_5patterns_6), V616)


let__2663 := tmp3146
_ = let__2663

tmp3174 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2663)


var ifres3147 Obj

if True == tmp3174 {
tmp3148 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3147 = tmp3148


} else {
tmp3149 := Call(__e, PrimFunc(symshen_4_5_1out), let__2663)


let__2664 := tmp3149
_ = let__2664

tmp3150 := Call(__e, PrimFunc(symshen_4in_1_6), let__2663)


let__2665 := tmp3150
_ = let__2665

tmp3173 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2665, sym_1_6)


var ifres3151 Obj

if True == tmp3173 {
tmp3152 := Call(__e, PrimFunc(symtail), let__2665)


let__2666 := tmp3152
_ = let__2666

tmp3171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__2666)
}
__typedArg0 := let__2666
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3153 Obj

if True == tmp3171 {
tmp3154 := Call(__e, PrimFunc(symhead), let__2666)


let__2667 := tmp3154
_ = let__2667

tmp3155 := Call(__e, PrimFunc(symtail), let__2666)


let__2668 := tmp3155
_ = let__2668

tmp3169 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2668, symwhere)


var ifres3156 Obj

if True == tmp3169 {
tmp3157 := Call(__e, PrimFunc(symtail), let__2668)


let__2669 := tmp3157
_ = let__2669

tmp3167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__2669)
}
__typedArg0 := let__2669
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3158 Obj

if True == tmp3167 {
tmp3159 := Call(__e, PrimFunc(symhead), let__2669)


let__2670 := tmp3159
_ = let__2670

tmp3160 := Call(__e, PrimFunc(symtail), let__2669)


let__2671 := tmp3160
_ = let__2671

tmp3161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2667, Nil)
}
__typedArg0 := let__2667
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2670, tmp3161)
}
__typedArg0 := let__2670
__typedArg1 := tmp3161
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp3162)
}
__typedArg0 := symwhere
__typedArg1 := tmp3162
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3164 := Call(__e, PrimFunc(sym_8p), let__2664, tmp3163)


tmp3165 := Call(__e, PrimFunc(symshen_4comb), let__2671, tmp3164)


ifres3158 = tmp3165


} else {
tmp3166 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3158 = tmp3166


}

ifres3156 = ifres3158


} else {
tmp3168 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3156 = tmp3168


}

ifres3153 = ifres3156


} else {
tmp3170 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3153 = tmp3170


}

ifres3151 = ifres3153


} else {
tmp3172 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3151 = tmp3172


}

ifres3147 = ifres3151


}

let__2662 := ifres3147
_ = let__2662

tmp3249 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2662)


if True == tmp3249 {
tmp3175 := Call(__e, PrimFunc(symshen_4_5patterns_6), V616)


let__2673 := tmp3175
_ = let__2673

tmp3191 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2673)


var ifres3176 Obj

if True == tmp3191 {
tmp3177 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3176 = tmp3177


} else {
tmp3178 := Call(__e, PrimFunc(symshen_4_5_1out), let__2673)


let__2674 := tmp3178
_ = let__2674

tmp3179 := Call(__e, PrimFunc(symshen_4in_1_6), let__2673)


let__2675 := tmp3179
_ = let__2675

tmp3190 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2675, sym_1_6)


var ifres3180 Obj

if True == tmp3190 {
tmp3181 := Call(__e, PrimFunc(symtail), let__2675)


let__2676 := tmp3181
_ = let__2676

tmp3188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__2676)
}
__typedArg0 := let__2676
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3182 Obj

if True == tmp3188 {
tmp3183 := Call(__e, PrimFunc(symhead), let__2676)


let__2677 := tmp3183
_ = let__2677

tmp3184 := Call(__e, PrimFunc(symtail), let__2676)


let__2678 := tmp3184
_ = let__2678

tmp3185 := Call(__e, PrimFunc(sym_8p), let__2674, let__2677)


tmp3186 := Call(__e, PrimFunc(symshen_4comb), let__2678, tmp3185)


ifres3182 = tmp3186


} else {
tmp3187 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3182 = tmp3187


}

ifres3180 = ifres3182


} else {
tmp3189 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3180 = tmp3189


}

ifres3176 = ifres3180


}

let__2672 := ifres3176
_ = let__2672

tmp3247 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2672)


if True == tmp3247 {
tmp3192 := Call(__e, PrimFunc(symshen_4_5patterns_6), V616)


let__2680 := tmp3192
_ = let__2680

tmp3222 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2680)


var ifres3193 Obj

if True == tmp3222 {
tmp3194 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3193 = tmp3194


} else {
tmp3195 := Call(__e, PrimFunc(symshen_4_5_1out), let__2680)


let__2681 := tmp3195
_ = let__2681

tmp3196 := Call(__e, PrimFunc(symshen_4in_1_6), let__2680)


let__2682 := tmp3196
_ = let__2682

tmp3221 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2682, sym_5_1)


var ifres3197 Obj

if True == tmp3221 {
tmp3198 := Call(__e, PrimFunc(symtail), let__2682)


let__2683 := tmp3198
_ = let__2683

tmp3219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__2683)
}
__typedArg0 := let__2683
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3199 Obj

if True == tmp3219 {
tmp3200 := Call(__e, PrimFunc(symhead), let__2683)


let__2684 := tmp3200
_ = let__2684

tmp3201 := Call(__e, PrimFunc(symtail), let__2683)


let__2685 := tmp3201
_ = let__2685

tmp3217 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2685, symwhere)


var ifres3202 Obj

if True == tmp3217 {
tmp3203 := Call(__e, PrimFunc(symtail), let__2685)


let__2686 := tmp3203
_ = let__2686

tmp3215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__2686)
}
__typedArg0 := let__2686
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3204 Obj

if True == tmp3215 {
tmp3205 := Call(__e, PrimFunc(symhead), let__2686)


let__2687 := tmp3205
_ = let__2687

tmp3206 := Call(__e, PrimFunc(symtail), let__2686)


let__2688 := tmp3206
_ = let__2688

tmp3207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2684, Nil)
}
__typedArg0 := let__2684
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4choicepoint_b, tmp3207)
}
__typedArg0 := symshen_4choicepoint_b
__typedArg1 := tmp3207
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3208, Nil)
}
__typedArg0 := tmp3208
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2687, tmp3209)
}
__typedArg0 := let__2687
__typedArg1 := tmp3209
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp3210)
}
__typedArg0 := symwhere
__typedArg1 := tmp3210
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3212 := Call(__e, PrimFunc(sym_8p), let__2681, tmp3211)


tmp3213 := Call(__e, PrimFunc(symshen_4comb), let__2688, tmp3212)


ifres3204 = tmp3213


} else {
tmp3214 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3204 = tmp3214


}

ifres3202 = ifres3204


} else {
tmp3216 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3202 = tmp3216


}

ifres3199 = ifres3202


} else {
tmp3218 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3199 = tmp3218


}

ifres3197 = ifres3199


} else {
tmp3220 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3197 = tmp3220


}

ifres3193 = ifres3197


}

let__2679 := ifres3193
_ = let__2679

tmp3245 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2679)


if True == tmp3245 {
tmp3223 := Call(__e, PrimFunc(symshen_4_5patterns_6), V616)


let__2690 := tmp3223
_ = let__2690

tmp3241 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2690)


var ifres3224 Obj

if True == tmp3241 {
tmp3225 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3224 = tmp3225


} else {
tmp3226 := Call(__e, PrimFunc(symshen_4_5_1out), let__2690)


let__2691 := tmp3226
_ = let__2691

tmp3227 := Call(__e, PrimFunc(symshen_4in_1_6), let__2690)


let__2692 := tmp3227
_ = let__2692

tmp3240 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2692, sym_5_1)


var ifres3228 Obj

if True == tmp3240 {
tmp3229 := Call(__e, PrimFunc(symtail), let__2692)


let__2693 := tmp3229
_ = let__2693

tmp3238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__2693)
}
__typedArg0 := let__2693
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3230 Obj

if True == tmp3238 {
tmp3231 := Call(__e, PrimFunc(symhead), let__2693)


let__2694 := tmp3231
_ = let__2694

tmp3232 := Call(__e, PrimFunc(symtail), let__2693)


let__2695 := tmp3232
_ = let__2695

tmp3233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2694, Nil)
}
__typedArg0 := let__2694
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4choicepoint_b, tmp3233)
}
__typedArg0 := symshen_4choicepoint_b
__typedArg1 := tmp3233
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3235 := Call(__e, PrimFunc(sym_8p), let__2691, tmp3234)


tmp3236 := Call(__e, PrimFunc(symshen_4comb), let__2695, tmp3235)


ifres3230 = tmp3236


} else {
tmp3237 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3230 = tmp3237


}

ifres3228 = ifres3230


} else {
tmp3239 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3228 = tmp3239


}

ifres3224 = ifres3228


}

let__2689 := ifres3224
_ = let__2689

tmp3243 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2689)


if True == tmp3243 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2689)
return
}


} else {
__e.Return(let__2679)
return
}


} else {
__e.Return(let__2672)
return
}


} else {
__e.Return(let__2662)
return
}


}, 1)

tmp3250 := Call(__e, ns2_1set, symshen_4_5rule_6, tmp3145)


_ = tmp3250

tmp3251 := MakeNative(func(__e *ControlFlow) {
V651 := __e.Get(1)
_ = V651
tmp3252 := Call(__e, PrimFunc(symshen_4_5pattern_6), V651)


let__2697 := tmp3252
_ = let__2697

tmp3265 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2697)


var ifres3253 Obj

if True == tmp3265 {
tmp3254 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3253 = tmp3254


} else {
tmp3255 := Call(__e, PrimFunc(symshen_4_5_1out), let__2697)


let__2698 := tmp3255
_ = let__2698

tmp3256 := Call(__e, PrimFunc(symshen_4in_1_6), let__2697)


let__2699 := tmp3256
_ = let__2699

tmp3257 := Call(__e, PrimFunc(symshen_4_5patterns_6), let__2699)


let__2700 := tmp3257
_ = let__2700

tmp3264 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2700)


var ifres3258 Obj

if True == tmp3264 {
tmp3259 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3258 = tmp3259


} else {
tmp3260 := Call(__e, PrimFunc(symshen_4_5_1out), let__2700)


let__2701 := tmp3260
_ = let__2701

tmp3261 := Call(__e, PrimFunc(symshen_4in_1_6), let__2700)


let__2702 := tmp3261
_ = let__2702

tmp3262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2698, let__2701)
}
__typedArg0 := let__2698
__typedArg1 := let__2701
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3263 := Call(__e, PrimFunc(symshen_4comb), let__2702, tmp3262)


ifres3258 = tmp3263


}

ifres3253 = ifres3258


}

let__2696 := ifres3253
_ = let__2696

tmp3275 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2696)


if True == tmp3275 {
tmp3266 := Call(__e, PrimFunc(sym_5e_6), V651)


let__2704 := tmp3266
_ = let__2704

tmp3271 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2704)


var ifres3267 Obj

if True == tmp3271 {
tmp3268 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3267 = tmp3268


} else {
tmp3269 := Call(__e, PrimFunc(symshen_4in_1_6), let__2704)


let__2705 := tmp3269
_ = let__2705

tmp3270 := Call(__e, PrimFunc(symshen_4comb), let__2705, Nil)


ifres3267 = tmp3270


}

let__2703 := ifres3267
_ = let__2703

tmp3273 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2703)


if True == tmp3273 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2703)
return
}


} else {
__e.Return(let__2696)
return
}


}, 1)

tmp3276 := Call(__e, ns2_1set, symshen_4_5patterns_6, tmp3251)


_ = tmp3276

tmp3277 := MakeNative(func(__e *ControlFlow) {
V662 := __e.Get(1)
_ = V662
tmp3309 := Call(__e, PrimFunc(symshen_4ccons_2), V662)


var ifres3278 Obj

if True == tmp3309 {
tmp3279 := Call(__e, PrimFunc(symhead), V662)


let__2707 := tmp3279
_ = let__2707

tmp3280 := Call(__e, PrimFunc(symtail), V662)


let__2708 := tmp3280
_ = let__2708

tmp3281 := Call(__e, PrimFunc(symshen_4_5constructor_6), let__2707)


let__2709 := tmp3281
_ = let__2709

tmp3307 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2709)


var ifres3282 Obj

if True == tmp3307 {
tmp3283 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3282 = tmp3283


} else {
tmp3284 := Call(__e, PrimFunc(symshen_4_5_1out), let__2709)


let__2710 := tmp3284
_ = let__2710

tmp3285 := Call(__e, PrimFunc(symshen_4in_1_6), let__2709)


let__2711 := tmp3285
_ = let__2711

tmp3286 := Call(__e, PrimFunc(symshen_4_5pattern1_6), let__2711)


let__2712 := tmp3286
_ = let__2712

tmp3306 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2712)


var ifres3287 Obj

if True == tmp3306 {
tmp3288 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3287 = tmp3288


} else {
tmp3289 := Call(__e, PrimFunc(symshen_4_5_1out), let__2712)


let__2713 := tmp3289
_ = let__2713

tmp3290 := Call(__e, PrimFunc(symshen_4in_1_6), let__2712)


let__2714 := tmp3290
_ = let__2714

tmp3291 := Call(__e, PrimFunc(symshen_4_5pattern2_6), let__2714)


let__2715 := tmp3291
_ = let__2715

tmp3305 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2715)


var ifres3292 Obj

if True == tmp3305 {
tmp3293 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3292 = tmp3293


} else {
tmp3294 := Call(__e, PrimFunc(symshen_4_5_1out), let__2715)


let__2716 := tmp3294
_ = let__2716

tmp3295 := Call(__e, PrimFunc(symshen_4in_1_6), let__2715)


let__2717 := tmp3295
_ = let__2717

tmp3296 := Call(__e, PrimFunc(sym_5end_6), let__2717)


let__2718 := tmp3296
_ = let__2718

tmp3304 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2718)


var ifres3297 Obj

if True == tmp3304 {
tmp3298 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3297 = tmp3298


} else {
tmp3299 := Call(__e, PrimFunc(symshen_4in_1_6), let__2718)


let__2719 := tmp3299
_ = let__2719

tmp3300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2716, Nil)
}
__typedArg0 := let__2716
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2713, tmp3300)
}
__typedArg0 := let__2713
__typedArg1 := tmp3300
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2710, tmp3301)
}
__typedArg0 := let__2710
__typedArg1 := tmp3301
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3303 := Call(__e, PrimFunc(symshen_4comb), let__2708, tmp3302)


ifres3297 = tmp3303


}

ifres3292 = ifres3297


}

ifres3287 = ifres3292


}

ifres3282 = ifres3287


}

ifres3278 = ifres3282


} else {
tmp3308 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3278 = tmp3308


}

let__2706 := ifres3278
_ = let__2706

tmp3355 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2706)


if True == tmp3355 {
tmp3330 := Call(__e, PrimFunc(symshen_4ccons_2), V662)


var ifres3310 Obj

if True == tmp3330 {
tmp3311 := Call(__e, PrimFunc(symhead), V662)


let__2721 := tmp3311
_ = let__2721

tmp3312 := Call(__e, PrimFunc(symtail), V662)


let__2722 := tmp3312
_ = let__2722

tmp3328 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2721, symvector)


var ifres3313 Obj

if True == tmp3328 {
tmp3314 := Call(__e, PrimFunc(symtail), let__2721)


let__2723 := tmp3314
_ = let__2723

tmp3326 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2723, MakeNumber(0))


var ifres3315 Obj

if True == tmp3326 {
tmp3316 := Call(__e, PrimFunc(symtail), let__2723)


let__2724 := tmp3316
_ = let__2724

tmp3317 := Call(__e, PrimFunc(sym_5end_6), let__2724)


let__2725 := tmp3317
_ = let__2725

tmp3324 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2725)


var ifres3318 Obj

if True == tmp3324 {
tmp3319 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3318 = tmp3319


} else {
tmp3320 := Call(__e, PrimFunc(symshen_4in_1_6), let__2725)


let__2726 := tmp3320
_ = let__2726

tmp3321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), Nil)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp3321)
}
__typedArg0 := symvector
__typedArg1 := tmp3321
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3323 := Call(__e, PrimFunc(symshen_4comb), let__2722, tmp3322)


ifres3318 = tmp3323


}

ifres3315 = ifres3318


} else {
tmp3325 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3315 = tmp3325


}

ifres3313 = ifres3315


} else {
tmp3327 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3313 = tmp3327


}

ifres3310 = ifres3313


} else {
tmp3329 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3310 = tmp3329


}

let__2720 := ifres3310
_ = let__2720

tmp3353 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2720)


if True == tmp3353 {
tmp3340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V662)
}
__typedArg0 := V662
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3331 Obj

if True == tmp3340 {
tmp3332 := Call(__e, PrimFunc(symhead), V662)


let__2728 := tmp3332
_ = let__2728

tmp3333 := Call(__e, PrimFunc(symtail), V662)


let__2729 := tmp3333
_ = let__2729

tmp3338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__2728)
}
__typedArg0 := let__2728
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3334 Obj

if True == tmp3338 {
tmp3335 := Call(__e, PrimFunc(symshen_4constructor_1error), let__2728)


tmp3336 := Call(__e, PrimFunc(symshen_4comb), let__2729, tmp3335)


ifres3334 = tmp3336


} else {
tmp3337 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3334 = tmp3337


}

ifres3331 = ifres3334


} else {
tmp3339 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3331 = tmp3339


}

let__2727 := ifres3331
_ = let__2727

tmp3351 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2727)


if True == tmp3351 {
tmp3341 := Call(__e, PrimFunc(symshen_4_5simple_1pattern_6), V662)


let__2731 := tmp3341
_ = let__2731

tmp3347 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2731)


var ifres3342 Obj

if True == tmp3347 {
tmp3343 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3342 = tmp3343


} else {
tmp3344 := Call(__e, PrimFunc(symshen_4_5_1out), let__2731)


let__2732 := tmp3344
_ = let__2732

tmp3345 := Call(__e, PrimFunc(symshen_4in_1_6), let__2731)


let__2733 := tmp3345
_ = let__2733

tmp3346 := Call(__e, PrimFunc(symshen_4comb), let__2733, let__2732)


ifres3342 = tmp3346


}

let__2730 := ifres3342
_ = let__2730

tmp3349 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2730)


if True == tmp3349 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2730)
return
}


} else {
__e.Return(let__2727)
return
}


} else {
__e.Return(let__2720)
return
}


} else {
__e.Return(let__2706)
return
}


}, 1)

tmp3356 := Call(__e, ns2_1set, symshen_4_5pattern_6, tmp3277)


_ = tmp3356

tmp3357 := MakeNative(func(__e *ControlFlow) {
V691 := __e.Get(1)
_ = V691
tmp3366 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V691)
}
__typedArg0 := V691
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3358 Obj

if True == tmp3366 {
tmp3359 := Call(__e, PrimFunc(symhead), V691)


let__2735 := tmp3359
_ = let__2735

tmp3360 := Call(__e, PrimFunc(symtail), V691)


let__2736 := tmp3360
_ = let__2736

tmp3364 := Call(__e, PrimFunc(symshen_4constructor_2), let__2735)


var ifres3361 Obj

if True == tmp3364 {
tmp3362 := Call(__e, PrimFunc(symshen_4comb), let__2736, let__2735)


ifres3361 = tmp3362


} else {
tmp3363 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3361 = tmp3363


}

ifres3358 = ifres3361


} else {
tmp3365 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3358 = tmp3365


}

let__2734 := ifres3358
_ = let__2734

tmp3368 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2734)


if True == tmp3368 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2734)
return
}


}, 1)

tmp3369 := Call(__e, ns2_1set, symshen_4_5constructor_6, tmp3357)


_ = tmp3369

tmp3370 := MakeNative(func(__e *ControlFlow) {
V695 := __e.Get(1)
_ = V695
tmp3371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, Nil)
}
__typedArg0 := sym_8v
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp3371)
}
__typedArg0 := sym_8s
__typedArg1 := tmp3371
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp3372)
}
__typedArg0 := sym_8p
__typedArg1 := tmp3372
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp3373)
}
__typedArg0 := symcons
__typedArg1 := tmp3373
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symelement_2), V695, tmp3374)
return


}, 1)

tmp3375 := Call(__e, ns2_1set, symshen_4constructor_2, tmp3370)


_ = tmp3375

tmp3376 := MakeNative(func(__e *ControlFlow) {
V696 := __e.Get(1)
_ = V696
tmp3377 := Call(__e, PrimFunc(symshen_4app), V696, MakeString(" is not a legitimate constructor\n"), symshen_4r)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp3377)
}
__typedArg0 := tmp3377
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

tmp3378 := Call(__e, ns2_1set, symshen_4constructor_1error, tmp3376)


_ = tmp3378

tmp3379 := MakeNative(func(__e *ControlFlow) {
V697 := __e.Get(1)
_ = V697
tmp3389 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V697)
}
__typedArg0 := V697
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3380 Obj

if True == tmp3389 {
tmp3381 := Call(__e, PrimFunc(symhead), V697)


let__2738 := tmp3381
_ = let__2738

tmp3382 := Call(__e, PrimFunc(symtail), V697)


let__2739 := tmp3382
_ = let__2739

tmp3387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__2738, sym__)
}
__typedArg0 := let__2738
__typedArg1 := sym__
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3383 Obj

if True == tmp3387 {
tmp3384 := Call(__e, PrimFunc(symgensym), symY)


tmp3385 := Call(__e, PrimFunc(symshen_4comb), let__2739, tmp3384)


ifres3383 = tmp3385


} else {
tmp3386 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3383 = tmp3386


}

ifres3380 = ifres3383


} else {
tmp3388 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3380 = tmp3388


}

let__2737 := ifres3380
_ = let__2737

tmp3405 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2737)


if True == tmp3405 {
tmp3401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V697)
}
__typedArg0 := V697
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3390 Obj

if True == tmp3401 {
tmp3391 := Call(__e, PrimFunc(symhead), V697)


let__2741 := tmp3391
_ = let__2741

tmp3392 := Call(__e, PrimFunc(symtail), V697)


let__2742 := tmp3392
_ = let__2742

tmp3396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1, Nil)
}
__typedArg0 := sym_5_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp3396)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp3396
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3398 := Call(__e, PrimFunc(symelement_2), let__2741, tmp3397)


tmp3399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp3398)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp3398
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres3393 Obj

if True == tmp3399 {
tmp3394 := Call(__e, PrimFunc(symshen_4comb), let__2742, let__2741)


ifres3393 = tmp3394


} else {
tmp3395 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3393 = tmp3395


}

ifres3390 = ifres3393


} else {
tmp3400 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3390 = tmp3400


}

let__2740 := ifres3390
_ = let__2740

tmp3403 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2740)


if True == tmp3403 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2740)
return
}


} else {
__e.Return(let__2737)
return
}


}, 1)

tmp3406 := Call(__e, ns2_1set, symshen_4_5simple_1pattern_6, tmp3379)


_ = tmp3406

tmp3407 := MakeNative(func(__e *ControlFlow) {
V704 := __e.Get(1)
_ = V704
tmp3408 := Call(__e, PrimFunc(symshen_4_5pattern_6), V704)


let__2744 := tmp3408
_ = let__2744

tmp3414 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2744)


var ifres3409 Obj

if True == tmp3414 {
tmp3410 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3409 = tmp3410


} else {
tmp3411 := Call(__e, PrimFunc(symshen_4_5_1out), let__2744)


let__2745 := tmp3411
_ = let__2745

tmp3412 := Call(__e, PrimFunc(symshen_4in_1_6), let__2744)


let__2746 := tmp3412
_ = let__2746

tmp3413 := Call(__e, PrimFunc(symshen_4comb), let__2746, let__2745)


ifres3409 = tmp3413


}

let__2743 := ifres3409
_ = let__2743

tmp3416 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2743)


if True == tmp3416 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2743)
return
}


}, 1)

tmp3417 := Call(__e, ns2_1set, symshen_4_5pattern1_6, tmp3407)


_ = tmp3417

tmp3418 := MakeNative(func(__e *ControlFlow) {
V709 := __e.Get(1)
_ = V709
tmp3419 := Call(__e, PrimFunc(symshen_4_5pattern_6), V709)


let__2748 := tmp3419
_ = let__2748

tmp3425 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2748)


var ifres3420 Obj

if True == tmp3425 {
tmp3421 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3420 = tmp3421


} else {
tmp3422 := Call(__e, PrimFunc(symshen_4_5_1out), let__2748)


let__2749 := tmp3422
_ = let__2749

tmp3423 := Call(__e, PrimFunc(symshen_4in_1_6), let__2748)


let__2750 := tmp3423
_ = let__2750

tmp3424 := Call(__e, PrimFunc(symshen_4comb), let__2750, let__2749)


ifres3420 = tmp3424


}

let__2747 := ifres3420
_ = let__2747

tmp3427 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2747)


if True == tmp3427 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2747)
return
}


}, 1)

tmp3428 := Call(__e, ns2_1set, symshen_4_5pattern2_6, tmp3418)


_ = tmp3428

tmp3429 := MakeNative(func(__e *ControlFlow) {
V714 := __e.Get(1)
_ = V714
tmp3430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(MakeNumber(2))
}
__typedArg0 := MakeNumber(2)
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

let__2751 := tmp3430
_ = let__2751

tmp3431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__2751, MakeNumber(0), symshen_4printF)
}
__typedArg0 := let__2751
__typedArg1 := MakeNumber(0)
__typedArg2 := symshen_4printF
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__2752 := tmp3431
_ = let__2752

tmp3432 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V714)
}
__typedArg0 := V714
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp3433 := Call(__e, PrimFunc(sym_8s), tmp3432, MakeString(")"))


tmp3434 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp3433)


tmp3435 := Call(__e, PrimFunc(sym_8s), MakeString("n"), tmp3434)


tmp3436 := Call(__e, PrimFunc(sym_8s), MakeString("f"), tmp3435)


tmp3437 := Call(__e, PrimFunc(sym_8s), MakeString("("), tmp3436)


tmp3438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__2752, MakeNumber(1), tmp3437)
}
__typedArg0 := let__2752
__typedArg1 := MakeNumber(1)
__typedArg2 := tmp3437
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__2753 := tmp3438
_ = let__2753

__e.Return(let__2753)
return


}, 1)

tmp3439 := Call(__e, ns2_1set, symshen_4fn_1print, tmp3429)


_ = tmp3439

tmp3440 := MakeNative(func(__e *ControlFlow) {
V718 := __e.Get(1)
_ = V718
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_5_1address) {
return PrimVectorGet(V718, MakeNumber(1))
}
__typedArg0 := V718
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_5_1address), __typedArg0, __typedArg1)
})())
return
}, 1)

tmp3441 := Call(__e, ns2_1set, symshen_4printF, tmp3440)


_ = tmp3441

tmp3442 := MakeNative(func(__e *ControlFlow) {
V723 := __e.Get(1)
_ = V723
V724 := __e.Get(2)
_ = V724
tmp3466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V724)
}
__typedArg0 := V724
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3462 Obj

if True == tmp3466 {
tmp3464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V724)
}
__typedArg0 := V724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3464)
}
__typedArg0 := Nil
__typedArg1 := tmp3464
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3463 Obj

if True == tmp3465 {
ifres3463 = True


} else {
ifres3463 = False


}

ifres3462 = ifres3463


} else {
ifres3462 = False


}

if True == ifres3462 {
tmp3443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V724)
}
__typedArg0 := V724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symlength), tmp3443)
return


} else {
tmp3460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V724)
}
__typedArg0 := V724
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3448 Obj

if True == tmp3460 {
tmp3458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V724)
}
__typedArg0 := V724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3458)
}
__typedArg0 := tmp3458
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3450 Obj

if True == tmp3459 {
tmp3452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V724)
}
__typedArg0 := V724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3453 := Call(__e, PrimFunc(symlength), tmp3452)


tmp3454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V724)
}
__typedArg0 := V724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3454)
}
__typedArg0 := tmp3454
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3456 := Call(__e, PrimFunc(symlength), tmp3455)


tmp3457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp3453, tmp3456)
}
__typedArg0 := tmp3453
__typedArg1 := tmp3456
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3451 Obj

if True == tmp3457 {
ifres3451 = True


} else {
ifres3451 = False


}

ifres3450 = ifres3451


} else {
ifres3450 = False


}

var ifres3449 Obj

if True == ifres3450 {
ifres3449 = True


} else {
ifres3449 = False


}

ifres3448 = ifres3449


} else {
ifres3448 = False


}

if True == ifres3448 {
tmp3444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V724)
}
__typedArg0 := V724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4arity_1chk), V723, tmp3444)
return


} else {
tmp3445 := Call(__e, PrimFunc(symshen_4app), V723, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("arity error in "))
__typedS1, __typedOK1 := TypedString(tmp3445)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("arity error in ")
__typedArg1 := tmp3445
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("arity error in "))
__typedS1, __typedOK1 := TypedString(tmp3445)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("arity error in ")
__typedArg1 := tmp3445
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}


}, 2)

tmp3467 := Call(__e, ns2_1set, symshen_4arity_1chk, tmp3442)


_ = tmp3467

tmp3468 := MakeNative(func(__e *ControlFlow) {
V725 := __e.Get(1)
_ = V725
V726 := __e.Get(2)
_ = V726
tmp3474 := Call(__e, PrimFunc(symtuple_2), V726)


if True == tmp3474 {
tmp3469 := Call(__e, PrimFunc(symfst), V726)


tmp3470 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp3469)


tmp3471 := Call(__e, PrimFunc(symsnd), V726)


tmp3472 := Call(__e, PrimFunc(symshen_4find_1free_1vars), tmp3470, tmp3471)


__e.TailApply(PrimFunc(symshen_4free_1variable_1error_1message), V725, tmp3472)
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.free-var-chk"))
}
__typedArg0 := MakeString("partial function shen.free-var-chk")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 2)

tmp3475 := Call(__e, ns2_1set, symshen_4free_1var_1chk, tmp3468)


_ = tmp3475

tmp3476 := MakeNative(func(__e *ControlFlow) {
V727 := __e.Get(1)
_ = V727
V728 := __e.Get(2)
_ = V728
tmp3488 := Call(__e, PrimFunc(symempty_2), V728)


if True == tmp3488 {
__e.Return(symshen_4skip)
return
} else {
tmp3477 := Call(__e, PrimFunc(symshen_4app), V727, MakeString(":"), symshen_4a)


tmp3478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("free variables in "))
__typedS1, __typedOK1 := TypedString(tmp3477)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("free variables in ")
__typedArg1 := tmp3477
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp3479 := Call(__e, PrimFunc(symstoutput))


tmp3480 := Call(__e, PrimFunc(sympr), tmp3478, tmp3479)


_ = tmp3480

tmp3481 := MakeNative(func(__e *ControlFlow) {
Z729 := __e.Get(1)
_ = Z729
tmp3482 := Call(__e, PrimFunc(symshen_4app), Z729, MakeString(""), symshen_4a)


tmp3483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" "))
__typedS1, __typedOK1 := TypedString(tmp3482)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" ")
__typedArg1 := tmp3482
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp3484 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp3483, tmp3484)
return


}, 1)

tmp3485 := Call(__e, PrimFunc(symmap), tmp3481, V728)


_ = tmp3485

tmp3486 := Call(__e, PrimFunc(symnl), MakeNumber(1))


_ = tmp3486

__e.TailApply(PrimFunc(symabort))
return


}


}, 2)

tmp3489 := Call(__e, ns2_1set, symshen_4free_1variable_1error_1message, tmp3476)


_ = tmp3489

tmp3490 := MakeNative(func(__e *ControlFlow) {
V732 := __e.Get(1)
_ = V732
tmp3498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V732)
}
__typedArg0 := V732
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp3498 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V732, Nil)
}
__typedArg0 := V732
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
} else {
tmp3496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V732)
}
__typedArg0 := V732
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3496 {
tmp3491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V732)
}
__typedArg0 := V732
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3492 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp3491)


tmp3493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V732)
}
__typedArg0 := V732
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3494 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp3493)


__e.TailApply(PrimFunc(symunion), tmp3492, tmp3494)
return


} else {
__e.Return(Nil)
return
}


}


}, 1)

tmp3499 := Call(__e, ns2_1set, symshen_4extract_1vars, tmp3490)


_ = tmp3499

tmp3500 := MakeNative(func(__e *ControlFlow) {
V737 := __e.Get(1)
_ = V737
V738 := __e.Get(2)
_ = V738
tmp3590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3577 Obj

if True == tmp3590 {
tmp3588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symprotect, tmp3588)
}
__typedArg0 := symprotect
__typedArg1 := tmp3588
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3579 Obj

if True == tmp3589 {
tmp3586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3586)
}
__typedArg0 := tmp3586
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3581 Obj

if True == tmp3587 {
tmp3583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3583)
}
__typedArg0 := tmp3583
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3584)
}
__typedArg0 := Nil
__typedArg1 := tmp3584
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3582 Obj

if True == tmp3585 {
ifres3582 = True


} else {
ifres3582 = False


}

ifres3581 = ifres3582


} else {
ifres3581 = False


}

var ifres3580 Obj

if True == ifres3581 {
ifres3580 = True


} else {
ifres3580 = False


}

ifres3579 = ifres3580


} else {
ifres3579 = False


}

var ifres3578 Obj

if True == ifres3579 {
ifres3578 = True


} else {
ifres3578 = False


}

ifres3577 = ifres3578


} else {
ifres3577 = False


}

if True == ifres3577 {
__e.Return(Nil)
return
} else {
tmp3575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3549 Obj

if True == tmp3575 {
tmp3573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp3573)
}
__typedArg0 := symlet
__typedArg1 := tmp3573
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3551 Obj

if True == tmp3574 {
tmp3571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3571)
}
__typedArg0 := tmp3571
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3553 Obj

if True == tmp3572 {
tmp3568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3568)
}
__typedArg0 := tmp3568
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3569)
}
__typedArg0 := tmp3569
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3555 Obj

if True == tmp3570 {
tmp3564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3564)
}
__typedArg0 := tmp3564
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3565)
}
__typedArg0 := tmp3565
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3566)
}
__typedArg0 := tmp3566
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3557 Obj

if True == tmp3567 {
tmp3559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3559)
}
__typedArg0 := tmp3559
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3560)
}
__typedArg0 := tmp3560
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3561)
}
__typedArg0 := tmp3561
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3562)
}
__typedArg0 := Nil
__typedArg1 := tmp3562
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3558 Obj

if True == tmp3563 {
ifres3558 = True


} else {
ifres3558 = False


}

ifres3557 = ifres3558


} else {
ifres3557 = False


}

var ifres3556 Obj

if True == ifres3557 {
ifres3556 = True


} else {
ifres3556 = False


}

ifres3555 = ifres3556


} else {
ifres3555 = False


}

var ifres3554 Obj

if True == ifres3555 {
ifres3554 = True


} else {
ifres3554 = False


}

ifres3553 = ifres3554


} else {
ifres3553 = False


}

var ifres3552 Obj

if True == ifres3553 {
ifres3552 = True


} else {
ifres3552 = False


}

ifres3551 = ifres3552


} else {
ifres3551 = False


}

var ifres3550 Obj

if True == ifres3551 {
ifres3550 = True


} else {
ifres3550 = False


}

ifres3549 = ifres3550


} else {
ifres3549 = False


}

if True == ifres3549 {
tmp3501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3501)
}
__typedArg0 := tmp3501
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3502)
}
__typedArg0 := tmp3502
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3504 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V737, tmp3503)


tmp3505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3505)
}
__typedArg0 := tmp3505
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3506, V737)
}
__typedArg0 := tmp3506
__typedArg1 := V737
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3508)
}
__typedArg0 := tmp3508
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3509)
}
__typedArg0 := tmp3509
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3510)
}
__typedArg0 := tmp3510
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3512 := Call(__e, PrimFunc(symshen_4find_1free_1vars), tmp3507, tmp3511)


__e.TailApply(PrimFunc(symunion), tmp3504, tmp3512)
return


} else {
tmp3547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3528 Obj

if True == tmp3547 {
tmp3545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp3545)
}
__typedArg0 := symlambda
__typedArg1 := tmp3545
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3530 Obj

if True == tmp3546 {
tmp3543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3543)
}
__typedArg0 := tmp3543
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3532 Obj

if True == tmp3544 {
tmp3540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3540)
}
__typedArg0 := tmp3540
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3541)
}
__typedArg0 := tmp3541
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3534 Obj

if True == tmp3542 {
tmp3536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3536)
}
__typedArg0 := tmp3536
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3537)
}
__typedArg0 := tmp3537
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3538)
}
__typedArg0 := Nil
__typedArg1 := tmp3538
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3535 Obj

if True == tmp3539 {
ifres3535 = True


} else {
ifres3535 = False


}

ifres3534 = ifres3535


} else {
ifres3534 = False


}

var ifres3533 Obj

if True == ifres3534 {
ifres3533 = True


} else {
ifres3533 = False


}

ifres3532 = ifres3533


} else {
ifres3532 = False


}

var ifres3531 Obj

if True == ifres3532 {
ifres3531 = True


} else {
ifres3531 = False


}

ifres3530 = ifres3531


} else {
ifres3530 = False


}

var ifres3529 Obj

if True == ifres3530 {
ifres3529 = True


} else {
ifres3529 = False


}

ifres3528 = ifres3529


} else {
ifres3528 = False


}

if True == ifres3528 {
tmp3513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3513)
}
__typedArg0 := tmp3513
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3514, V737)
}
__typedArg0 := tmp3514
__typedArg1 := V737
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3516)
}
__typedArg0 := tmp3516
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3517)
}
__typedArg0 := tmp3517
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4find_1free_1vars), tmp3515, tmp3518)
return


} else {
tmp3526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3526 {
tmp3519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3520 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V737, tmp3519)


tmp3521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3522 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V737, tmp3521)


__e.TailApply(PrimFunc(symunion), tmp3520, tmp3522)
return


} else {
tmp3524 := Call(__e, PrimFunc(symshen_4free_1variable_2), V738, V737)


if True == tmp3524 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V738, Nil)
}
__typedArg0 := V738
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

tmp3591 := Call(__e, ns2_1set, symshen_4find_1free_1vars, tmp3500)


_ = tmp3591

tmp3592 := MakeNative(func(__e *ControlFlow) {
V739 := __e.Get(1)
_ = V739
V740 := __e.Get(2)
_ = V740
tmp3597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V739)
}
__typedArg0 := V739
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp3597 {
tmp3594 := Call(__e, PrimFunc(symelement_2), V739, V740)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp3594)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp3594
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

tmp3598 := Call(__e, ns2_1set, symshen_4free_1variable_2, tmp3592)


_ = tmp3598

tmp3599 := MakeNative(func(__e *ControlFlow) {
V741 := __e.Get(1)
_ = V741
V742 := __e.Get(2)
_ = V742
tmp3600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_duserdefs_d)
}
__typedArg0 := symshen_4_duserdefs_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp3601 := Call(__e, PrimFunc(symadjoin), V741, tmp3600)


tmp3602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_duserdefs_d, tmp3601)
}
__typedArg0 := symshen_4_duserdefs_d
__typedArg1 := tmp3601
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp3602

tmp3603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V741, symshen_4source, V742, tmp3603)
return


}, 2)

tmp3604 := Call(__e, ns2_1set, symshen_4record_1kl, tmp3599)


_ = tmp3604

tmp3605 := MakeNative(func(__e *ControlFlow) {
V743 := __e.Get(1)
_ = V743
V744 := __e.Get(2)
_ = V744
V745 := __e.Get(3)
_ = V745
tmp3606 := Call(__e, PrimFunc(symshen_4parameters), V745)


let__2754 := tmp3606
_ = let__2754

tmp3607 := Call(__e, PrimFunc(symshen_4kl_1body), V744, let__2754)


tmp3608 := Call(__e, PrimFunc(symshen_4scan_1body), V743, tmp3607)


let__2755 := tmp3608
_ = let__2755

tmp3609 := Call(__e, PrimFunc(symshen_4cond_1form), let__2755)


tmp3610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3609, Nil)
}
__typedArg0 := tmp3609
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2754, tmp3610)
}
__typedArg0 := let__2754
__typedArg1 := tmp3610
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V743, tmp3611)
}
__typedArg0 := V743
__typedArg1 := tmp3611
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp3612)
}
__typedArg0 := symdefun
__typedArg1 := tmp3612
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__2756 := tmp3613
_ = let__2756

__e.Return(let__2756)
return


}, 3)

tmp3614 := Call(__e, ns2_1set, symshen_4compile_1to_1kl, tmp3605)


_ = tmp3614

tmp3615 := MakeNative(func(__e *ControlFlow) {
V749 := __e.Get(1)
_ = V749
tmp3620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V749)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V749
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp3620 {
__e.Return(Nil)
return
} else {
tmp3616 := Call(__e, PrimFunc(symgensym), symV)


tmp3618 := Call(__e, PrimFunc(symshen_4parameters), (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_1) {
__typedN0, __typedOK0 := TypedFloat64(V749)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_1) {
return TypedMaterializeNumber((__typedN0 - __typedN1))
}}
__typedArg0 := V749
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_1), __typedArg0, __typedArg1)
})())


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3616, tmp3618)
}
__typedArg0 := tmp3616
__typedArg1 := tmp3618
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp3621 := Call(__e, ns2_1set, symshen_4parameters, tmp3615)


_ = tmp3621

tmp3622 := MakeNative(func(__e *ControlFlow) {
V752 := __e.Get(1)
_ = V752
tmp3646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3626 Obj

if True == tmp3646 {
tmp3644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3644)
}
__typedArg0 := tmp3644
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3628 Obj

if True == tmp3645 {
tmp3641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3641)
}
__typedArg0 := tmp3641
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, tmp3642)
}
__typedArg0 := True
__typedArg1 := tmp3642
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3630 Obj

if True == tmp3643 {
tmp3638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3638)
}
__typedArg0 := tmp3638
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3639)
}
__typedArg0 := tmp3639
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3632 Obj

if True == tmp3640 {
tmp3634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3634)
}
__typedArg0 := tmp3634
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3635)
}
__typedArg0 := tmp3635
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3636)
}
__typedArg0 := Nil
__typedArg1 := tmp3636
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3633 Obj

if True == tmp3637 {
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

var ifres3629 Obj

if True == ifres3630 {
ifres3629 = True


} else {
ifres3629 = False


}

ifres3628 = ifres3629


} else {
ifres3628 = False


}

var ifres3627 Obj

if True == ifres3628 {
ifres3627 = True


} else {
ifres3627 = False


}

ifres3626 = ifres3627


} else {
ifres3626 = False


}

if True == ifres3626 {
tmp3623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3623)
}
__typedArg0 := tmp3623
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3624)
}
__typedArg0 := tmp3624
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcond, V752)
}
__typedArg0 := symcond
__typedArg1 := V752
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return
}


}, 1)

tmp3647 := Call(__e, ns2_1set, symshen_4cond_1form, tmp3622)


_ = tmp3647

tmp3648 := MakeNative(func(__e *ControlFlow) {
V761 := __e.Get(1)
_ = V761
V762 := __e.Get(2)
_ = V762
tmp3692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V762)
}
__typedArg0 := Nil
__typedArg1 := V762
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp3692 {
tmp3649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V761, Nil)
}
__typedArg0 := V761
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4f_1error, tmp3649)
}
__typedArg0 := symshen_4f_1error
__typedArg1 := tmp3649
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3650, Nil)
}
__typedArg0 := tmp3650
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp3651)
}
__typedArg0 := True
__typedArg1 := tmp3651
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3652, Nil)
}
__typedArg0 := tmp3652
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp3690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3686 Obj

if True == tmp3690 {
tmp3688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3689 := Call(__e, PrimFunc(symshen_4choicepoint_2), tmp3688)


var ifres3687 Obj

if True == tmp3689 {
ifres3687 = True


} else {
ifres3687 = False


}

ifres3686 = ifres3687


} else {
ifres3686 = False


}

if True == ifres3686 {
tmp3653 := Call(__e, PrimFunc(symgensym), symFreeze)


tmp3654 := Call(__e, PrimFunc(symgensym), symResult)


tmp3655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4choicepoint), V761, tmp3653, tmp3654, tmp3655, tmp3656)
return


} else {
tmp3684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3664 Obj

if True == tmp3684 {
tmp3682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3682)
}
__typedArg0 := tmp3682
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3666 Obj

if True == tmp3683 {
tmp3679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3679)
}
__typedArg0 := tmp3679
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, tmp3680)
}
__typedArg0 := True
__typedArg1 := tmp3680
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3668 Obj

if True == tmp3681 {
tmp3676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3676)
}
__typedArg0 := tmp3676
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3677)
}
__typedArg0 := tmp3677
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3670 Obj

if True == tmp3678 {
tmp3672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3672)
}
__typedArg0 := tmp3672
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3673)
}
__typedArg0 := tmp3673
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3674)
}
__typedArg0 := Nil
__typedArg1 := tmp3674
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3671 Obj

if True == tmp3675 {
ifres3671 = True


} else {
ifres3671 = False


}

ifres3670 = ifres3671


} else {
ifres3670 = False


}

var ifres3669 Obj

if True == ifres3670 {
ifres3669 = True


} else {
ifres3669 = False


}

ifres3668 = ifres3669


} else {
ifres3668 = False


}

var ifres3667 Obj

if True == ifres3668 {
ifres3667 = True


} else {
ifres3667 = False


}

ifres3666 = ifres3667


} else {
ifres3666 = False


}

var ifres3665 Obj

if True == ifres3666 {
ifres3665 = True


} else {
ifres3665 = False


}

ifres3664 = ifres3665


} else {
ifres3664 = False


}

if True == ifres3664 {
tmp3657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3657, Nil)
}
__typedArg0 := tmp3657
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp3662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3662 {
tmp3658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3660 := Call(__e, PrimFunc(symshen_4scan_1body), V761, tmp3659)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3658, tmp3660)
}
__typedArg0 := tmp3658
__typedArg1 := tmp3660
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

tmp3693 := Call(__e, ns2_1set, symshen_4scan_1body, tmp3648)


_ = tmp3693

tmp3694 := MakeNative(func(__e *ControlFlow) {
V769 := __e.Get(1)
_ = V769
tmp3729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V769)
}
__typedArg0 := V769
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3696 Obj

if True == tmp3729 {
tmp3727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V769)
}
__typedArg0 := V769
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3727)
}
__typedArg0 := tmp3727
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3698 Obj

if True == tmp3728 {
tmp3724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V769)
}
__typedArg0 := V769
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3724)
}
__typedArg0 := tmp3724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3725)
}
__typedArg0 := tmp3725
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3700 Obj

if True == tmp3726 {
tmp3720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V769)
}
__typedArg0 := V769
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3720)
}
__typedArg0 := tmp3720
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3721)
}
__typedArg0 := tmp3721
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4choicepoint_b, tmp3722)
}
__typedArg0 := symshen_4choicepoint_b
__typedArg1 := tmp3722
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3702 Obj

if True == tmp3723 {
tmp3716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V769)
}
__typedArg0 := V769
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3716)
}
__typedArg0 := tmp3716
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3717)
}
__typedArg0 := tmp3717
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3718)
}
__typedArg0 := tmp3718
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3704 Obj

if True == tmp3719 {
tmp3711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V769)
}
__typedArg0 := V769
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3711)
}
__typedArg0 := tmp3711
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3712)
}
__typedArg0 := tmp3712
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3713)
}
__typedArg0 := tmp3713
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3714)
}
__typedArg0 := Nil
__typedArg1 := tmp3714
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3706 Obj

if True == tmp3715 {
tmp3708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V769)
}
__typedArg0 := V769
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3708)
}
__typedArg0 := tmp3708
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3709)
}
__typedArg0 := Nil
__typedArg1 := tmp3709
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3707 Obj

if True == tmp3710 {
ifres3707 = True


} else {
ifres3707 = False


}

ifres3706 = ifres3707


} else {
ifres3706 = False


}

var ifres3705 Obj

if True == ifres3706 {
ifres3705 = True


} else {
ifres3705 = False


}

ifres3704 = ifres3705


} else {
ifres3704 = False


}

var ifres3703 Obj

if True == ifres3704 {
ifres3703 = True


} else {
ifres3703 = False


}

ifres3702 = ifres3703


} else {
ifres3702 = False


}

var ifres3701 Obj

if True == ifres3702 {
ifres3701 = True


} else {
ifres3701 = False


}

ifres3700 = ifres3701


} else {
ifres3700 = False


}

var ifres3699 Obj

if True == ifres3700 {
ifres3699 = True


} else {
ifres3699 = False


}

ifres3698 = ifres3699


} else {
ifres3698 = False


}

var ifres3697 Obj

if True == ifres3698 {
ifres3697 = True


} else {
ifres3697 = False


}

ifres3696 = ifres3697


} else {
ifres3696 = False


}

if True == ifres3696 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp3730 := Call(__e, ns2_1set, symshen_4choicepoint_2, tmp3694)


_ = tmp3730

tmp3731 := MakeNative(func(__e *ControlFlow) {
V785 := __e.Get(1)
_ = V785
V786 := __e.Get(2)
_ = V786
V787 := __e.Get(3)
_ = V787
V788 := __e.Get(4)
_ = V788
V789 := __e.Get(5)
_ = V789
tmp3923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3845 Obj

if True == tmp3923 {
tmp3921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3921)
}
__typedArg0 := tmp3921
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3847 Obj

if True == tmp3922 {
tmp3918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3918)
}
__typedArg0 := tmp3918
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3919)
}
__typedArg0 := tmp3919
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3849 Obj

if True == tmp3920 {
tmp3914 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3914)
}
__typedArg0 := tmp3914
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3915)
}
__typedArg0 := tmp3915
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3916)
}
__typedArg0 := tmp3916
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3851 Obj

if True == tmp3917 {
tmp3909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3909)
}
__typedArg0 := tmp3909
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3910)
}
__typedArg0 := tmp3910
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3911)
}
__typedArg0 := tmp3911
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3912)
}
__typedArg0 := tmp3912
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3853 Obj

if True == tmp3913 {
tmp3903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3903)
}
__typedArg0 := tmp3903
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3904)
}
__typedArg0 := tmp3904
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3905)
}
__typedArg0 := tmp3905
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3906)
}
__typedArg0 := tmp3906
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail_1if, tmp3907)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp3907
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3855 Obj

if True == tmp3908 {
tmp3897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3897)
}
__typedArg0 := tmp3897
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3898)
}
__typedArg0 := tmp3898
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3899)
}
__typedArg0 := tmp3899
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3900)
}
__typedArg0 := tmp3900
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3901)
}
__typedArg0 := tmp3901
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3857 Obj

if True == tmp3902 {
tmp3890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3890)
}
__typedArg0 := tmp3890
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3891)
}
__typedArg0 := tmp3891
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3892)
}
__typedArg0 := tmp3892
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3893)
}
__typedArg0 := tmp3893
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3894)
}
__typedArg0 := tmp3894
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3895)
}
__typedArg0 := tmp3895
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3859 Obj

if True == tmp3896 {
tmp3882 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3882)
}
__typedArg0 := tmp3882
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3883)
}
__typedArg0 := tmp3883
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3884)
}
__typedArg0 := tmp3884
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3885)
}
__typedArg0 := tmp3885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3886)
}
__typedArg0 := tmp3886
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3887)
}
__typedArg0 := tmp3887
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3888)
}
__typedArg0 := Nil
__typedArg1 := tmp3888
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3861 Obj

if True == tmp3889 {
tmp3877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3877)
}
__typedArg0 := tmp3877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3878)
}
__typedArg0 := tmp3878
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3879)
}
__typedArg0 := tmp3879
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3881 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3880)
}
__typedArg0 := Nil
__typedArg1 := tmp3880
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3863 Obj

if True == tmp3881 {
tmp3874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3874)
}
__typedArg0 := tmp3874
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3875)
}
__typedArg0 := Nil
__typedArg1 := tmp3875
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3865 Obj

if True == tmp3876 {
tmp3867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3867)
}
__typedArg0 := tmp3867
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3868)
}
__typedArg0 := tmp3868
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3869)
}
__typedArg0 := tmp3869
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3870)
}
__typedArg0 := tmp3870
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3871)
}
__typedArg0 := tmp3871
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V785, tmp3872)
}
__typedArg0 := V785
__typedArg1 := tmp3872
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3866 Obj

if True == tmp3873 {
ifres3866 = True


} else {
ifres3866 = False


}

ifres3865 = ifres3866


} else {
ifres3865 = False


}

var ifres3864 Obj

if True == ifres3865 {
ifres3864 = True


} else {
ifres3864 = False


}

ifres3863 = ifres3864


} else {
ifres3863 = False


}

var ifres3862 Obj

if True == ifres3863 {
ifres3862 = True


} else {
ifres3862 = False


}

ifres3861 = ifres3862


} else {
ifres3861 = False


}

var ifres3860 Obj

if True == ifres3861 {
ifres3860 = True


} else {
ifres3860 = False


}

ifres3859 = ifres3860


} else {
ifres3859 = False


}

var ifres3858 Obj

if True == ifres3859 {
ifres3858 = True


} else {
ifres3858 = False


}

ifres3857 = ifres3858


} else {
ifres3857 = False


}

var ifres3856 Obj

if True == ifres3857 {
ifres3856 = True


} else {
ifres3856 = False


}

ifres3855 = ifres3856


} else {
ifres3855 = False


}

var ifres3854 Obj

if True == ifres3855 {
ifres3854 = True


} else {
ifres3854 = False


}

ifres3853 = ifres3854


} else {
ifres3853 = False


}

var ifres3852 Obj

if True == ifres3853 {
ifres3852 = True


} else {
ifres3852 = False


}

ifres3851 = ifres3852


} else {
ifres3851 = False


}

var ifres3850 Obj

if True == ifres3851 {
ifres3850 = True


} else {
ifres3850 = False


}

ifres3849 = ifres3850


} else {
ifres3849 = False


}

var ifres3848 Obj

if True == ifres3849 {
ifres3848 = True


} else {
ifres3848 = False


}

ifres3847 = ifres3848


} else {
ifres3847 = False


}

var ifres3846 Obj

if True == ifres3847 {
ifres3846 = True


} else {
ifres3846 = False


}

ifres3845 = ifres3846


} else {
ifres3845 = False


}

if True == ifres3845 {
tmp3732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3732)
}
__typedArg0 := tmp3732
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3733)
}
__typedArg0 := tmp3733
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3734)
}
__typedArg0 := tmp3734
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3735)
}
__typedArg0 := tmp3735
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3736)
}
__typedArg0 := tmp3736
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3738 := Call(__e, PrimFunc(symshen_4scan_1body), tmp3737, V789)


tmp3739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcond, tmp3738)
}
__typedArg0 := symcond
__typedArg1 := tmp3738
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3739, Nil)
}
__typedArg0 := tmp3739
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp3740)
}
__typedArg0 := symfreeze
__typedArg1 := tmp3740
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3743)
}
__typedArg0 := tmp3743
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3744)
}
__typedArg0 := tmp3744
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3745)
}
__typedArg0 := tmp3745
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3746)
}
__typedArg0 := tmp3746
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3747)
}
__typedArg0 := tmp3747
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3748)
}
__typedArg0 := tmp3748
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3750)
}
__typedArg0 := tmp3750
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3751)
}
__typedArg0 := tmp3751
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3752)
}
__typedArg0 := tmp3752
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3754)
}
__typedArg0 := tmp3754
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V787, Nil)
}
__typedArg0 := V787
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3755, tmp3756)
}
__typedArg0 := tmp3755
__typedArg1 := tmp3756
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V786, Nil)
}
__typedArg0 := V786
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp3758)
}
__typedArg0 := symthaw
__typedArg1 := tmp3758
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V787, Nil)
}
__typedArg0 := V787
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3759, tmp3760)
}
__typedArg0 := tmp3759
__typedArg1 := tmp3760
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3757, tmp3761)
}
__typedArg0 := tmp3757
__typedArg1 := tmp3761
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp3762)
}
__typedArg0 := symif
__typedArg1 := tmp3762
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3764 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3763, Nil)
}
__typedArg0 := tmp3763
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3765 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3749, tmp3764)
}
__typedArg0 := tmp3749
__typedArg1 := tmp3764
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3766 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V787, tmp3765)
}
__typedArg0 := V787
__typedArg1 := tmp3765
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp3766)
}
__typedArg0 := symlet
__typedArg1 := tmp3766
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V786, Nil)
}
__typedArg0 := V786
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp3768)
}
__typedArg0 := symthaw
__typedArg1 := tmp3768
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3769, Nil)
}
__typedArg0 := tmp3769
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3767, tmp3770)
}
__typedArg0 := tmp3767
__typedArg1 := tmp3770
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3742, tmp3771)
}
__typedArg0 := tmp3742
__typedArg1 := tmp3771
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp3772)
}
__typedArg0 := symif
__typedArg1 := tmp3772
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3773, Nil)
}
__typedArg0 := tmp3773
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3741, tmp3774)
}
__typedArg0 := tmp3741
__typedArg1 := tmp3774
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V786, tmp3775)
}
__typedArg0 := V786
__typedArg1 := tmp3775
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp3776)
}
__typedArg0 := symlet
__typedArg1 := tmp3776
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3777, Nil)
}
__typedArg0 := tmp3777
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp3778)
}
__typedArg0 := True
__typedArg1 := tmp3778
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3779, Nil)
}
__typedArg0 := tmp3779
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp3843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3816 Obj

if True == tmp3843 {
tmp3841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3841)
}
__typedArg0 := tmp3841
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3818 Obj

if True == tmp3842 {
tmp3838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3838)
}
__typedArg0 := tmp3838
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3839)
}
__typedArg0 := tmp3839
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3820 Obj

if True == tmp3840 {
tmp3834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3834)
}
__typedArg0 := tmp3834
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3835)
}
__typedArg0 := tmp3835
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3836)
}
__typedArg0 := tmp3836
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3822 Obj

if True == tmp3837 {
tmp3829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
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
return PrimTail(tmp3830)
}
__typedArg0 := tmp3830
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3831)
}
__typedArg0 := tmp3831
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3832)
}
__typedArg0 := Nil
__typedArg1 := tmp3832
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3824 Obj

if True == tmp3833 {
tmp3826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3826)
}
__typedArg0 := tmp3826
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3827)
}
__typedArg0 := Nil
__typedArg1 := tmp3827
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3825 Obj

if True == tmp3828 {
ifres3825 = True


} else {
ifres3825 = False


}

ifres3824 = ifres3825


} else {
ifres3824 = False


}

var ifres3823 Obj

if True == ifres3824 {
ifres3823 = True


} else {
ifres3823 = False


}

ifres3822 = ifres3823


} else {
ifres3822 = False


}

var ifres3821 Obj

if True == ifres3822 {
ifres3821 = True


} else {
ifres3821 = False


}

ifres3820 = ifres3821


} else {
ifres3820 = False


}

var ifres3819 Obj

if True == ifres3820 {
ifres3819 = True


} else {
ifres3819 = False


}

ifres3818 = ifres3819


} else {
ifres3818 = False


}

var ifres3817 Obj

if True == ifres3818 {
ifres3817 = True


} else {
ifres3817 = False


}

ifres3816 = ifres3817


} else {
ifres3816 = False


}

if True == ifres3816 {
tmp3780 := Call(__e, PrimFunc(symshen_4scan_1body), V785, V789)


tmp3781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcond, tmp3780)
}
__typedArg0 := symcond
__typedArg1 := tmp3780
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3781, Nil)
}
__typedArg0 := tmp3781
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp3782)
}
__typedArg0 := symfreeze
__typedArg1 := tmp3782
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3785)
}
__typedArg0 := tmp3785
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3786)
}
__typedArg0 := tmp3786
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3787)
}
__typedArg0 := tmp3787
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, Nil)
}
__typedArg0 := symfail
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3789, Nil)
}
__typedArg0 := tmp3789
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V787, tmp3790)
}
__typedArg0 := V787
__typedArg1 := tmp3790
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp3791)
}
__typedArg0 := sym_a
__typedArg1 := tmp3791
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V786, Nil)
}
__typedArg0 := V786
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp3793)
}
__typedArg0 := symthaw
__typedArg1 := tmp3793
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V787, Nil)
}
__typedArg0 := V787
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3794, tmp3795)
}
__typedArg0 := tmp3794
__typedArg1 := tmp3795
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3792, tmp3796)
}
__typedArg0 := tmp3792
__typedArg1 := tmp3796
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp3797)
}
__typedArg0 := symif
__typedArg1 := tmp3797
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3798, Nil)
}
__typedArg0 := tmp3798
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3788, tmp3799)
}
__typedArg0 := tmp3788
__typedArg1 := tmp3799
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V787, tmp3800)
}
__typedArg0 := V787
__typedArg1 := tmp3800
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp3801)
}
__typedArg0 := symlet
__typedArg1 := tmp3801
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V786, Nil)
}
__typedArg0 := V786
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp3803)
}
__typedArg0 := symthaw
__typedArg1 := tmp3803
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3804, Nil)
}
__typedArg0 := tmp3804
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3802, tmp3805)
}
__typedArg0 := tmp3802
__typedArg1 := tmp3805
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3784, tmp3806)
}
__typedArg0 := tmp3784
__typedArg1 := tmp3806
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp3807)
}
__typedArg0 := symif
__typedArg1 := tmp3807
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3808, Nil)
}
__typedArg0 := tmp3808
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3783, tmp3809)
}
__typedArg0 := tmp3783
__typedArg1 := tmp3809
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V786, tmp3810)
}
__typedArg0 := V786
__typedArg1 := tmp3810
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp3811)
}
__typedArg0 := symlet
__typedArg1 := tmp3811
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3812, Nil)
}
__typedArg0 := tmp3812
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp3813)
}
__typedArg0 := True
__typedArg1 := tmp3813
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3814, Nil)
}
__typedArg0 := tmp3814
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

tmp3924 := Call(__e, ns2_1set, symshen_4choicepoint, tmp3731)


_ = tmp3924

tmp3925 := MakeNative(func(__e *ControlFlow) {
V791 := __e.Get(1)
_ = V791
V792 := __e.Get(2)
_ = V792
V793 := __e.Get(3)
_ = V793
tmp3938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V791, V793)
}
__typedArg0 := V791
__typedArg1 := V793
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp3938 {
__e.Return(V792)
return
} else {
tmp3936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V793)
}
__typedArg0 := V793
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3936 {
tmp3926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V793)
}
__typedArg0 := V793
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3927 := Call(__e, PrimFunc(symshen_4rep_1X), V791, V792, tmp3926)


let__2757 := tmp3927
_ = let__2757

tmp3933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V793)
}
__typedArg0 := V793
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__2757, tmp3933)
}
__typedArg0 := let__2757
__typedArg1 := tmp3933
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp3934 {
tmp3928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V793)
}
__typedArg0 := V793
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V793)
}
__typedArg0 := V793
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3930 := Call(__e, PrimFunc(symshen_4rep_1X), V791, V792, tmp3929)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3928, tmp3930)
}
__typedArg0 := tmp3928
__typedArg1 := tmp3930
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp3931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V793)
}
__typedArg0 := V793
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2757, tmp3931)
}
__typedArg0 := let__2757
__typedArg1 := tmp3931
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


} else {
__e.Return(V793)
return
}


}


}, 3)

tmp3939 := Call(__e, ns2_1set, symshen_4rep_1X, tmp3925)


_ = tmp3939

tmp3940 := MakeNative(func(__e *ControlFlow) {
V795 := __e.Get(1)
_ = V795
V796 := __e.Get(2)
_ = V796
tmp3941 := MakeNative(func(__e *ControlFlow) {
Z797 := __e.Get(1)
_ = Z797
tmp3942 := Call(__e, PrimFunc(symfst), Z797)


tmp3943 := Call(__e, PrimFunc(symsnd), Z797)


tmp3944 := Call(__e, PrimFunc(symshen_4alpha_1convert), tmp3943)


__e.TailApply(PrimFunc(symshen_4triple_1stack), Nil, tmp3942, V796, tmp3944)
return


}, 1)

__e.TailApply(PrimFunc(symmap), tmp3941, V795)
return


}, 2)

tmp3945 := Call(__e, ns2_1set, symshen_4kl_1body, tmp3940)


_ = tmp3945

tmp3946 := MakeNative(func(__e *ControlFlow) {
V798 := __e.Get(1)
_ = V798
tmp4025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4006 Obj

if True == tmp4025 {
tmp4023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp4023)
}
__typedArg0 := symlambda
__typedArg1 := tmp4023
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4008 Obj

if True == tmp4024 {
tmp4021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4021)
}
__typedArg0 := tmp4021
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4010 Obj

if True == tmp4022 {
tmp4018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4018)
}
__typedArg0 := tmp4018
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4019)
}
__typedArg0 := tmp4019
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4012 Obj

if True == tmp4020 {
tmp4014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4014)
}
__typedArg0 := tmp4014
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4015)
}
__typedArg0 := tmp4015
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4016)
}
__typedArg0 := Nil
__typedArg1 := tmp4016
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4013 Obj

if True == tmp4017 {
ifres4013 = True


} else {
ifres4013 = False


}

ifres4012 = ifres4013


} else {
ifres4012 = False


}

var ifres4011 Obj

if True == ifres4012 {
ifres4011 = True


} else {
ifres4011 = False


}

ifres4010 = ifres4011


} else {
ifres4010 = False


}

var ifres4009 Obj

if True == ifres4010 {
ifres4009 = True


} else {
ifres4009 = False


}

ifres4008 = ifres4009


} else {
ifres4008 = False


}

var ifres4007 Obj

if True == ifres4008 {
ifres4007 = True


} else {
ifres4007 = False


}

ifres4006 = ifres4007


} else {
ifres4006 = False


}

if True == ifres4006 {
tmp3947 := Call(__e, PrimFunc(symgensym), symZ)


let__2758 := tmp3947
_ = let__2758

tmp3948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V798)
}
__typedArg0 := V798
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
return PrimTail(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3950)
}
__typedArg0 := tmp3950
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3951)
}
__typedArg0 := tmp3951
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3953 := Call(__e, PrimFunc(symshen_4beta), tmp3949, let__2758, tmp3952)


tmp3954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3953, Nil)
}
__typedArg0 := tmp3953
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2758, tmp3954)
}
__typedArg0 := let__2758
__typedArg1 := tmp3954
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp3955)
}
__typedArg0 := symlambda
__typedArg1 := tmp3955
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__2759 := tmp3956
_ = let__2759

tmp3957 := MakeNative(func(__e *ControlFlow) {
Z801 := __e.Get(1)
_ = Z801
__e.TailApply(PrimFunc(symshen_4alpha_1convert), Z801)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp3957, let__2759)
return


} else {
tmp4004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3978 Obj

if True == tmp4004 {
tmp4002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp4002)
}
__typedArg0 := symlet
__typedArg1 := tmp4002
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3980 Obj

if True == tmp4003 {
tmp4000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4000)
}
__typedArg0 := tmp4000
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3982 Obj

if True == tmp4001 {
tmp3997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3997)
}
__typedArg0 := tmp3997
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3998)
}
__typedArg0 := tmp3998
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3984 Obj

if True == tmp3999 {
tmp3993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3993)
}
__typedArg0 := tmp3993
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3994)
}
__typedArg0 := tmp3994
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3995)
}
__typedArg0 := tmp3995
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3986 Obj

if True == tmp3996 {
tmp3988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3988)
}
__typedArg0 := tmp3988
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3989)
}
__typedArg0 := tmp3989
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3990)
}
__typedArg0 := tmp3990
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3991)
}
__typedArg0 := Nil
__typedArg1 := tmp3991
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3987 Obj

if True == tmp3992 {
ifres3987 = True


} else {
ifres3987 = False


}

ifres3986 = ifres3987


} else {
ifres3986 = False


}

var ifres3985 Obj

if True == ifres3986 {
ifres3985 = True


} else {
ifres3985 = False


}

ifres3984 = ifres3985


} else {
ifres3984 = False


}

var ifres3983 Obj

if True == ifres3984 {
ifres3983 = True


} else {
ifres3983 = False


}

ifres3982 = ifres3983


} else {
ifres3982 = False


}

var ifres3981 Obj

if True == ifres3982 {
ifres3981 = True


} else {
ifres3981 = False


}

ifres3980 = ifres3981


} else {
ifres3980 = False


}

var ifres3979 Obj

if True == ifres3980 {
ifres3979 = True


} else {
ifres3979 = False


}

ifres3978 = ifres3979


} else {
ifres3978 = False


}

if True == ifres3978 {
tmp3958 := Call(__e, PrimFunc(symgensym), symW)


let__2760 := tmp3958
_ = let__2760

tmp3959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3959)
}
__typedArg0 := tmp3959
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3960)
}
__typedArg0 := tmp3960
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3962)
}
__typedArg0 := tmp3962
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3964)
}
__typedArg0 := tmp3964
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3965)
}
__typedArg0 := tmp3965
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3966)
}
__typedArg0 := tmp3966
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3968 := Call(__e, PrimFunc(symshen_4beta), tmp3963, let__2760, tmp3967)


tmp3969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3968, Nil)
}
__typedArg0 := tmp3968
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3961, tmp3969)
}
__typedArg0 := tmp3961
__typedArg1 := tmp3969
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2760, tmp3970)
}
__typedArg0 := let__2760
__typedArg1 := tmp3970
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp3971)
}
__typedArg0 := symlet
__typedArg1 := tmp3971
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__2761 := tmp3972
_ = let__2761

tmp3973 := MakeNative(func(__e *ControlFlow) {
Z804 := __e.Get(1)
_ = Z804
__e.TailApply(PrimFunc(symshen_4alpha_1convert), Z804)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp3973, let__2761)
return


} else {
tmp3976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3976 {
tmp3974 := MakeNative(func(__e *ControlFlow) {
Z805 := __e.Get(1)
_ = Z805
__e.TailApply(PrimFunc(symshen_4alpha_1convert), Z805)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp3974, V798)
return


} else {
__e.Return(V798)
return
}


}


}


}, 1)

tmp4026 := Call(__e, ns2_1set, symshen_4alpha_1convert, tmp3946)


_ = tmp4026

tmp4027 := MakeNative(func(__e *ControlFlow) {
V814 := __e.Get(1)
_ = V814
V815 := __e.Get(2)
_ = V815
V816 := __e.Get(3)
_ = V816
V817 := __e.Get(4)
_ = V817
tmp4157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V815)
}
__typedArg0 := Nil
__typedArg1 := V815
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4132 Obj

if True == tmp4157 {
tmp4156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V816)
}
__typedArg0 := Nil
__typedArg1 := V816
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4134 Obj

if True == tmp4156 {
tmp4155 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V817)
}
__typedArg0 := V817
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4136 Obj

if True == tmp4155 {
tmp4153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V817)
}
__typedArg0 := V817
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4154 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp4153)
}
__typedArg0 := symwhere
__typedArg1 := tmp4153
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4138 Obj

if True == tmp4154 {
tmp4151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V817)
}
__typedArg0 := V817
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4151)
}
__typedArg0 := tmp4151
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4140 Obj

if True == tmp4152 {
tmp4148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V817)
}
__typedArg0 := V817
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4148)
}
__typedArg0 := tmp4148
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4149)
}
__typedArg0 := tmp4149
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4142 Obj

if True == tmp4150 {
tmp4144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V817)
}
__typedArg0 := V817
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4145 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4144)
}
__typedArg0 := tmp4144
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4145)
}
__typedArg0 := tmp4145
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4146)
}
__typedArg0 := Nil
__typedArg1 := tmp4146
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4143 Obj

if True == tmp4147 {
ifres4143 = True


} else {
ifres4143 = False


}

ifres4142 = ifres4143


} else {
ifres4142 = False


}

var ifres4141 Obj

if True == ifres4142 {
ifres4141 = True


} else {
ifres4141 = False


}

ifres4140 = ifres4141


} else {
ifres4140 = False


}

var ifres4139 Obj

if True == ifres4140 {
ifres4139 = True


} else {
ifres4139 = False


}

ifres4138 = ifres4139


} else {
ifres4138 = False


}

var ifres4137 Obj

if True == ifres4138 {
ifres4137 = True


} else {
ifres4137 = False


}

ifres4136 = ifres4137


} else {
ifres4136 = False


}

var ifres4135 Obj

if True == ifres4136 {
ifres4135 = True


} else {
ifres4135 = False


}

ifres4134 = ifres4135


} else {
ifres4134 = False


}

var ifres4133 Obj

if True == ifres4134 {
ifres4133 = True


} else {
ifres4133 = False


}

ifres4132 = ifres4133


} else {
ifres4132 = False


}

if True == ifres4132 {
tmp4028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V817)
}
__typedArg0 := V817
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4028)
}
__typedArg0 := tmp4028
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4029, V814)
}
__typedArg0 := tmp4029
__typedArg1 := V814
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4031 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V817)
}
__typedArg0 := V817
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4032 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4031)
}
__typedArg0 := tmp4031
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4033 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4032)
}
__typedArg0 := tmp4032
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp4030, Nil, Nil, tmp4033)
return


} else {
tmp4130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V815)
}
__typedArg0 := Nil
__typedArg1 := V815
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4127 Obj

if True == tmp4130 {
tmp4129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V816)
}
__typedArg0 := Nil
__typedArg1 := V816
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4128 Obj

if True == tmp4129 {
ifres4128 = True


} else {
ifres4128 = False


}

ifres4127 = ifres4128


} else {
ifres4127 = False


}

if True == ifres4127 {
tmp4034 := Call(__e, PrimFunc(symreverse), V814)


tmp4035 := Call(__e, PrimFunc(symshen_4rectify_1test), tmp4034)


tmp4036 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V817, Nil)
}
__typedArg0 := V817
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4035, tmp4036)
}
__typedArg0 := tmp4035
__typedArg1 := tmp4036
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp4125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4118 Obj

if True == tmp4125 {
tmp4124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V816)
}
__typedArg0 := V816
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4120 Obj

if True == tmp4124 {
tmp4122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp4122)
}
__typedArg0 := tmp4122
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres4121 Obj

if True == tmp4123 {
ifres4121 = True


} else {
ifres4121 = False


}

ifres4120 = ifres4121


} else {
ifres4120 = False


}

var ifres4119 Obj

if True == ifres4120 {
ifres4119 = True


} else {
ifres4119 = False


}

ifres4118 = ifres4119


} else {
ifres4118 = False


}

if True == ifres4118 {
tmp4037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V816)
}
__typedArg0 := V816
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4039 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V816)
}
__typedArg0 := V816
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4041 := Call(__e, PrimFunc(symshen_4beta), tmp4039, tmp4040, V817)


__e.TailApply(PrimFunc(symshen_4triple_1stack), V814, tmp4037, tmp4038, tmp4041)
return


} else {
tmp4116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4091 Obj

if True == tmp4116 {
tmp4114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4114)
}
__typedArg0 := tmp4114
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4093 Obj

if True == tmp4115 {
tmp4111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4111)
}
__typedArg0 := tmp4111
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4113 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4112)
}
__typedArg0 := tmp4112
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4095 Obj

if True == tmp4113 {
tmp4107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4107)
}
__typedArg0 := tmp4107
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4108)
}
__typedArg0 := tmp4108
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4109)
}
__typedArg0 := tmp4109
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4097 Obj

if True == tmp4110 {
tmp4102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4102)
}
__typedArg0 := tmp4102
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4103)
}
__typedArg0 := tmp4103
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4104)
}
__typedArg0 := tmp4104
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4105)
}
__typedArg0 := Nil
__typedArg1 := tmp4105
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4099 Obj

if True == tmp4106 {
tmp4101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V816)
}
__typedArg0 := V816
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4100 Obj

if True == tmp4101 {
ifres4100 = True


} else {
ifres4100 = False


}

ifres4099 = ifres4100


} else {
ifres4099 = False


}

var ifres4098 Obj

if True == ifres4099 {
ifres4098 = True


} else {
ifres4098 = False


}

ifres4097 = ifres4098


} else {
ifres4097 = False


}

var ifres4096 Obj

if True == ifres4097 {
ifres4096 = True


} else {
ifres4096 = False


}

ifres4095 = ifres4096


} else {
ifres4095 = False


}

var ifres4094 Obj

if True == ifres4095 {
ifres4094 = True


} else {
ifres4094 = False


}

ifres4093 = ifres4094


} else {
ifres4093 = False


}

var ifres4092 Obj

if True == ifres4093 {
ifres4092 = True


} else {
ifres4092 = False


}

ifres4091 = ifres4092


} else {
ifres4091 = False


}

if True == ifres4091 {
tmp4042 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4042)
}
__typedArg0 := tmp4042
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4044 := Call(__e, PrimFunc(symshen_4op_1test), tmp4043)


tmp4045 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V816)
}
__typedArg0 := V816
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4045, Nil)
}
__typedArg0 := tmp4045
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4044, tmp4046)
}
__typedArg0 := tmp4044
__typedArg1 := tmp4046
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4047, V814)
}
__typedArg0 := tmp4047
__typedArg1 := V814
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4049)
}
__typedArg0 := tmp4049
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4050)
}
__typedArg0 := tmp4050
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4052)
}
__typedArg0 := tmp4052
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4053)
}
__typedArg0 := tmp4053
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4054)
}
__typedArg0 := tmp4054
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4055, tmp4056)
}
__typedArg0 := tmp4055
__typedArg1 := tmp4056
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4051, tmp4057)
}
__typedArg0 := tmp4051
__typedArg1 := tmp4057
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4059)
}
__typedArg0 := tmp4059
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4061 := Call(__e, PrimFunc(symshen_4op1), tmp4060)


tmp4062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V816)
}
__typedArg0 := V816
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4063 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4062, Nil)
}
__typedArg0 := tmp4062
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4061, tmp4063)
}
__typedArg0 := tmp4061
__typedArg1 := tmp4063
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4065)
}
__typedArg0 := tmp4065
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4067 := Call(__e, PrimFunc(symshen_4op2), tmp4066)


tmp4068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V816)
}
__typedArg0 := V816
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4068, Nil)
}
__typedArg0 := tmp4068
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4067, tmp4069)
}
__typedArg0 := tmp4067
__typedArg1 := tmp4069
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V816)
}
__typedArg0 := V816
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4072 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4070, tmp4071)
}
__typedArg0 := tmp4070
__typedArg1 := tmp4071
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4064, tmp4072)
}
__typedArg0 := tmp4064
__typedArg1 := tmp4072
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V816)
}
__typedArg0 := V816
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4076 := Call(__e, PrimFunc(symshen_4beta), tmp4074, tmp4075, V817)


__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp4048, tmp4058, tmp4073, tmp4076)
return


} else {
tmp4089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4086 Obj

if True == tmp4089 {
tmp4088 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V816)
}
__typedArg0 := V816
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4087 Obj

if True == tmp4088 {
ifres4087 = True


} else {
ifres4087 = False


}

ifres4086 = ifres4087


} else {
ifres4086 = False


}

if True == ifres4086 {
tmp4077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V816)
}
__typedArg0 := V816
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4078, Nil)
}
__typedArg0 := tmp4078
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4077, tmp4079)
}
__typedArg0 := tmp4077
__typedArg1 := tmp4079
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp4080)
}
__typedArg0 := sym_a
__typedArg1 := tmp4080
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4081, V814)
}
__typedArg0 := tmp4081
__typedArg1 := V814
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V816)
}
__typedArg0 := V816
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp4082, tmp4083, tmp4084, V817)
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


}, 4)

tmp4158 := Call(__e, ns2_1set, symshen_4triple_1stack, tmp4027)


_ = tmp4158

tmp4159 := MakeNative(func(__e *ControlFlow) {
V820 := __e.Get(1)
_ = V820
tmp4178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V820)
}
__typedArg0 := Nil
__typedArg1 := V820
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4178 {
__e.Return(True)
return
} else {
tmp4176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V820)
}
__typedArg0 := V820
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4172 Obj

if True == tmp4176 {
tmp4174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V820)
}
__typedArg0 := V820
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4174)
}
__typedArg0 := Nil
__typedArg1 := tmp4174
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4173 Obj

if True == tmp4175 {
ifres4173 = True


} else {
ifres4173 = False


}

ifres4172 = ifres4173


} else {
ifres4172 = False


}

if True == ifres4172 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V820)
}
__typedArg0 := V820
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp4170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V820)
}
__typedArg0 := V820
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4166 Obj

if True == tmp4170 {
tmp4168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V820)
}
__typedArg0 := V820
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4168)
}
__typedArg0 := tmp4168
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4167 Obj

if True == tmp4169 {
ifres4167 = True


} else {
ifres4167 = False


}

ifres4166 = ifres4167


} else {
ifres4166 = False


}

if True == ifres4166 {
tmp4160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V820)
}
__typedArg0 := V820
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V820)
}
__typedArg0 := V820
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4162 := Call(__e, PrimFunc(symshen_4rectify_1test), tmp4161)


tmp4163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4162, Nil)
}
__typedArg0 := tmp4162
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4160, tmp4163)
}
__typedArg0 := tmp4160
__typedArg1 := tmp4163
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp4164)
}
__typedArg0 := symand
__typedArg1 := tmp4164
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

tmp4179 := Call(__e, ns2_1set, symshen_4rectify_1test, tmp4159)


_ = tmp4179

tmp4180 := MakeNative(func(__e *ControlFlow) {
V830 := __e.Get(1)
_ = V830
V831 := __e.Get(2)
_ = V831
V832 := __e.Get(3)
_ = V832
tmp4257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V830, V832)
}
__typedArg0 := V830
__typedArg1 := V832
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4257 {
__e.Return(V831)
return
} else {
tmp4255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4231 Obj

if True == tmp4255 {
tmp4253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp4253)
}
__typedArg0 := symlambda
__typedArg1 := tmp4253
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4233 Obj

if True == tmp4254 {
tmp4251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4251)
}
__typedArg0 := tmp4251
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4235 Obj

if True == tmp4252 {
tmp4248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
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

var ifres4237 Obj

if True == tmp4250 {
tmp4244 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
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

var ifres4239 Obj

if True == tmp4247 {
tmp4241 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4241)
}
__typedArg0 := tmp4241
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4243 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V830, tmp4242)
}
__typedArg0 := V830
__typedArg1 := tmp4242
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4240 Obj

if True == tmp4243 {
ifres4240 = True


} else {
ifres4240 = False


}

ifres4239 = ifres4240


} else {
ifres4239 = False


}

var ifres4238 Obj

if True == ifres4239 {
ifres4238 = True


} else {
ifres4238 = False


}

ifres4237 = ifres4238


} else {
ifres4237 = False


}

var ifres4236 Obj

if True == ifres4237 {
ifres4236 = True


} else {
ifres4236 = False


}

ifres4235 = ifres4236


} else {
ifres4235 = False


}

var ifres4234 Obj

if True == ifres4235 {
ifres4234 = True


} else {
ifres4234 = False


}

ifres4233 = ifres4234


} else {
ifres4233 = False


}

var ifres4232 Obj

if True == ifres4233 {
ifres4232 = True


} else {
ifres4232 = False


}

ifres4231 = ifres4232


} else {
ifres4231 = False


}

if True == ifres4231 {
__e.Return(V832)
return
} else {
tmp4229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4198 Obj

if True == tmp4229 {
tmp4227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp4227)
}
__typedArg0 := symlet
__typedArg1 := tmp4227
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4200 Obj

if True == tmp4228 {
tmp4225 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4226 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4225)
}
__typedArg0 := tmp4225
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4202 Obj

if True == tmp4226 {
tmp4222 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4222)
}
__typedArg0 := tmp4222
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4223)
}
__typedArg0 := tmp4223
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4204 Obj

if True == tmp4224 {
tmp4218 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4218)
}
__typedArg0 := tmp4218
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4219)
}
__typedArg0 := tmp4219
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4220)
}
__typedArg0 := tmp4220
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4206 Obj

if True == tmp4221 {
tmp4213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4214 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4213)
}
__typedArg0 := tmp4213
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4215 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4214)
}
__typedArg0 := tmp4214
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4216 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4215)
}
__typedArg0 := tmp4215
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4216)
}
__typedArg0 := Nil
__typedArg1 := tmp4216
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4208 Obj

if True == tmp4217 {
tmp4210 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4210)
}
__typedArg0 := tmp4210
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V830, tmp4211)
}
__typedArg0 := V830
__typedArg1 := tmp4211
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4209 Obj

if True == tmp4212 {
ifres4209 = True


} else {
ifres4209 = False


}

ifres4208 = ifres4209


} else {
ifres4208 = False


}

var ifres4207 Obj

if True == ifres4208 {
ifres4207 = True


} else {
ifres4207 = False


}

ifres4206 = ifres4207


} else {
ifres4206 = False


}

var ifres4205 Obj

if True == ifres4206 {
ifres4205 = True


} else {
ifres4205 = False


}

ifres4204 = ifres4205


} else {
ifres4204 = False


}

var ifres4203 Obj

if True == ifres4204 {
ifres4203 = True


} else {
ifres4203 = False


}

ifres4202 = ifres4203


} else {
ifres4202 = False


}

var ifres4201 Obj

if True == ifres4202 {
ifres4201 = True


} else {
ifres4201 = False


}

ifres4200 = ifres4201


} else {
ifres4200 = False


}

var ifres4199 Obj

if True == ifres4200 {
ifres4199 = True


} else {
ifres4199 = False


}

ifres4198 = ifres4199


} else {
ifres4198 = False


}

if True == ifres4198 {
tmp4181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4181)
}
__typedArg0 := tmp4181
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4183)
}
__typedArg0 := tmp4183
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4185)
}
__typedArg0 := tmp4185
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4186)
}
__typedArg0 := tmp4186
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4188 := Call(__e, PrimFunc(symshen_4beta), tmp4184, V831, tmp4187)


tmp4189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4189)
}
__typedArg0 := tmp4189
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4191 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4190)
}
__typedArg0 := tmp4190
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4192 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4188, tmp4191)
}
__typedArg0 := tmp4188
__typedArg1 := tmp4191
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4182, tmp4192)
}
__typedArg0 := tmp4182
__typedArg1 := tmp4192
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp4193)
}
__typedArg0 := symlet
__typedArg1 := tmp4193
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp4196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp4196 {
tmp4194 := MakeNative(func(__e *ControlFlow) {
Z833 := __e.Get(1)
_ = Z833
__e.TailApply(PrimFunc(symshen_4beta), V830, V831, Z833)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp4194, V832)
return


} else {
__e.Return(V832)
return
}


}


}


}


}, 3)

tmp4258 := Call(__e, ns2_1set, symshen_4beta, tmp4180)


_ = tmp4258

tmp4259 := MakeNative(func(__e *ControlFlow) {
V836 := __e.Get(1)
_ = V836
tmp4267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, V836)
}
__typedArg0 := symcons
__typedArg1 := V836
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4267 {
__e.Return(symhd)
return
} else {
tmp4265 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8s, V836)
}
__typedArg0 := sym_8s
__typedArg1 := V836
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4265 {
__e.Return(symhdstr)
return
} else {
tmp4263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8p, V836)
}
__typedArg0 := sym_8p
__typedArg1 := V836
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4263 {
__e.Return(symfst)
return
} else {
tmp4261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8v, V836)
}
__typedArg0 := sym_8v
__typedArg1 := V836
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4261 {
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

tmp4268 := Call(__e, ns2_1set, symshen_4op1, tmp4259)


_ = tmp4268

tmp4269 := MakeNative(func(__e *ControlFlow) {
V839 := __e.Get(1)
_ = V839
tmp4277 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, V839)
}
__typedArg0 := symcons
__typedArg1 := V839
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4277 {
__e.Return(symtl)
return
} else {
tmp4275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8s, V839)
}
__typedArg0 := sym_8s
__typedArg1 := V839
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4275 {
__e.Return(symtlstr)
return
} else {
tmp4273 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8p, V839)
}
__typedArg0 := sym_8p
__typedArg1 := V839
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4273 {
__e.Return(symsnd)
return
} else {
tmp4271 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8v, V839)
}
__typedArg0 := sym_8v
__typedArg1 := V839
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4271 {
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

tmp4278 := Call(__e, ns2_1set, symshen_4op2, tmp4269)


_ = tmp4278

tmp4279 := MakeNative(func(__e *ControlFlow) {
V842 := __e.Get(1)
_ = V842
tmp4287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, V842)
}
__typedArg0 := symcons
__typedArg1 := V842
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4287 {
__e.Return(symcons_2)
return
} else {
tmp4285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8s, V842)
}
__typedArg0 := sym_8s
__typedArg1 := V842
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4285 {
__e.Return(symshen_4_7string_2)
return
} else {
tmp4283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8p, V842)
}
__typedArg0 := sym_8p
__typedArg1 := V842
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4283 {
__e.Return(symtuple_2)
return
} else {
tmp4281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8v, V842)
}
__typedArg0 := sym_8v
__typedArg1 := V842
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4281 {
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

tmp4288 := Call(__e, ns2_1set, symshen_4op_1test, tmp4279)


_ = tmp4288

tmp4289 := MakeNative(func(__e *ControlFlow) {
V843 := __e.Get(1)
_ = V843
tmp4291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V843)
}
__typedArg0 := MakeString("")
__typedArg1 := V843
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4291 {
__e.Return(False)
return
} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstring_2) {
return PrimIsString(V843)
}
__typedArg0 := V843
return Call(__e, PrimFunc(symstring_2), __typedArg0)
})())
return
}


}, 1)

tmp4292 := Call(__e, ns2_1set, symshen_4_7string_2, tmp4289)


_ = tmp4292

tmp4293 := MakeNative(func(__e *ControlFlow) {
V844 := __e.Get(1)
_ = V844
tmp4295 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp4296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V844, tmp4295)
}
__typedArg0 := V844
__typedArg1 := tmp4295
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4296 {
__e.Return(False)
return
} else {
__e.TailApply(PrimFunc(symvector_2), V844)
return
}


}, 1)

tmp4297 := Call(__e, ns2_1set, symshen_4_7vector_2, tmp4293)


_ = tmp4297

tmp4298 := MakeNative(func(__e *ControlFlow) {
V847 := __e.Get(1)
_ = V847
tmp4302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V847)
}
__typedArg0 := sym_7
__typedArg1 := V847
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4302 {
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
tmp4300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V847)
}
__typedArg0 := sym_1
__typedArg1 := V847
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4300 {
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

tmp4303 := Call(__e, ns2_1set, symfactorise, tmp4298)


_ = tmp4303

tmp4304 := MakeNative(func(__e *ControlFlow) {
V848 := __e.Get(1)
_ = V848
tmp4306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dfactorise_2_d)
}
__typedArg0 := symshen_4_dfactorise_2_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

if True == tmp4306 {
__e.TailApply(PrimFunc(symshen_4factor), V848)
return
} else {
__e.Return(V848)
return
}


}, 1)

tmp4307 := Call(__e, ns2_1set, symshen_4factorise_1code, tmp4304)


_ = tmp4307

tmp4308 := MakeNative(func(__e *ControlFlow) {
V849 := __e.Get(1)
_ = V849
tmp4365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V849)
}
__typedArg0 := V849
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4324 Obj

if True == tmp4365 {
tmp4363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V849)
}
__typedArg0 := V849
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4364 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp4363)
}
__typedArg0 := symdefun
__typedArg1 := tmp4363
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4326 Obj

if True == tmp4364 {
tmp4361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V849)
}
__typedArg0 := V849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4362 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4361)
}
__typedArg0 := tmp4361
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4328 Obj

if True == tmp4362 {
tmp4358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V849)
}
__typedArg0 := V849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4358)
}
__typedArg0 := tmp4358
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4359)
}
__typedArg0 := tmp4359
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4330 Obj

if True == tmp4360 {
tmp4354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V849)
}
__typedArg0 := V849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4354)
}
__typedArg0 := tmp4354
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4355)
}
__typedArg0 := tmp4355
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4356)
}
__typedArg0 := tmp4356
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4332 Obj

if True == tmp4357 {
tmp4349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V849)
}
__typedArg0 := V849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4349)
}
__typedArg0 := tmp4349
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4350)
}
__typedArg0 := tmp4350
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4351)
}
__typedArg0 := tmp4351
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4352)
}
__typedArg0 := tmp4352
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4334 Obj

if True == tmp4353 {
tmp4343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V849)
}
__typedArg0 := V849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4343)
}
__typedArg0 := tmp4343
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4345 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4344)
}
__typedArg0 := tmp4344
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4345)
}
__typedArg0 := tmp4345
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4346)
}
__typedArg0 := tmp4346
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcond, tmp4347)
}
__typedArg0 := symcond
__typedArg1 := tmp4347
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4336 Obj

if True == tmp4348 {
tmp4338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V849)
}
__typedArg0 := V849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4338)
}
__typedArg0 := tmp4338
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4340)
}
__typedArg0 := tmp4340
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4341)
}
__typedArg0 := Nil
__typedArg1 := tmp4341
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4337 Obj

if True == tmp4342 {
ifres4337 = True


} else {
ifres4337 = False


}

ifres4336 = ifres4337


} else {
ifres4336 = False


}

var ifres4335 Obj

if True == ifres4336 {
ifres4335 = True


} else {
ifres4335 = False


}

ifres4334 = ifres4335


} else {
ifres4334 = False


}

var ifres4333 Obj

if True == ifres4334 {
ifres4333 = True


} else {
ifres4333 = False


}

ifres4332 = ifres4333


} else {
ifres4332 = False


}

var ifres4331 Obj

if True == ifres4332 {
ifres4331 = True


} else {
ifres4331 = False


}

ifres4330 = ifres4331


} else {
ifres4330 = False


}

var ifres4329 Obj

if True == ifres4330 {
ifres4329 = True


} else {
ifres4329 = False


}

ifres4328 = ifres4329


} else {
ifres4328 = False


}

var ifres4327 Obj

if True == ifres4328 {
ifres4327 = True


} else {
ifres4327 = False


}

ifres4326 = ifres4327


} else {
ifres4326 = False


}

var ifres4325 Obj

if True == ifres4326 {
ifres4325 = True


} else {
ifres4325 = False


}

ifres4324 = ifres4325


} else {
ifres4324 = False


}

if True == ifres4324 {
tmp4309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V849)
}
__typedArg0 := V849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4309)
}
__typedArg0 := tmp4309
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V849)
}
__typedArg0 := V849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4311)
}
__typedArg0 := tmp4311
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4312)
}
__typedArg0 := tmp4312
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V849)
}
__typedArg0 := V849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4314)
}
__typedArg0 := tmp4314
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4315)
}
__typedArg0 := tmp4315
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4317 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4316)
}
__typedArg0 := tmp4316
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4317)
}
__typedArg0 := tmp4317
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4319 := Call(__e, PrimFunc(symshen_4factor_1recognisors), tmp4318)


tmp4320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4319, Nil)
}
__typedArg0 := tmp4319
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4313, tmp4320)
}
__typedArg0 := tmp4313
__typedArg1 := tmp4320
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4310, tmp4321)
}
__typedArg0 := tmp4310
__typedArg1 := tmp4321
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp4322)
}
__typedArg0 := symdefun
__typedArg1 := tmp4322
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V849)
return
}


}, 1)

tmp4366 := Call(__e, ns2_1set, symshen_4factor, tmp4308)


_ = tmp4366

tmp4367 := MakeNative(func(__e *ControlFlow) {
V852 := __e.Get(1)
_ = V852
tmp4517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4497 Obj

if True == tmp4517 {
tmp4515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4515)
}
__typedArg0 := tmp4515
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4499 Obj

if True == tmp4516 {
tmp4512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4512)
}
__typedArg0 := tmp4512
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, tmp4513)
}
__typedArg0 := True
__typedArg1 := tmp4513
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4501 Obj

if True == tmp4514 {
tmp4509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4509)
}
__typedArg0 := tmp4509
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4510)
}
__typedArg0 := tmp4510
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4503 Obj

if True == tmp4511 {
tmp4505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4505)
}
__typedArg0 := tmp4505
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4506)
}
__typedArg0 := tmp4506
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4507)
}
__typedArg0 := Nil
__typedArg1 := tmp4507
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4504 Obj

if True == tmp4508 {
ifres4504 = True


} else {
ifres4504 = False


}

ifres4503 = ifres4504


} else {
ifres4503 = False


}

var ifres4502 Obj

if True == ifres4503 {
ifres4502 = True


} else {
ifres4502 = False


}

ifres4501 = ifres4502


} else {
ifres4501 = False


}

var ifres4500 Obj

if True == ifres4501 {
ifres4500 = True


} else {
ifres4500 = False


}

ifres4499 = ifres4500


} else {
ifres4499 = False


}

var ifres4498 Obj

if True == ifres4499 {
ifres4498 = True


} else {
ifres4498 = False


}

ifres4497 = ifres4498


} else {
ifres4497 = False


}

if True == ifres4497 {
tmp4368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4368)
}
__typedArg0 := tmp4368
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4369)
}
__typedArg0 := tmp4369
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp4495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4448 Obj

if True == tmp4495 {
tmp4493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4493)
}
__typedArg0 := tmp4493
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4450 Obj

if True == tmp4494 {
tmp4490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4490)
}
__typedArg0 := tmp4490
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4491)
}
__typedArg0 := tmp4491
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4452 Obj

if True == tmp4492 {
tmp4486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4486)
}
__typedArg0 := tmp4486
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symand, tmp4488)
}
__typedArg0 := symand
__typedArg1 := tmp4488
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4454 Obj

if True == tmp4489 {
tmp4482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4482)
}
__typedArg0 := tmp4482
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4483)
}
__typedArg0 := tmp4483
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4484)
}
__typedArg0 := tmp4484
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4456 Obj

if True == tmp4485 {
tmp4477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4477)
}
__typedArg0 := tmp4477
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4478)
}
__typedArg0 := tmp4478
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4479)
}
__typedArg0 := tmp4479
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4480)
}
__typedArg0 := tmp4480
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4458 Obj

if True == tmp4481 {
tmp4471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4471)
}
__typedArg0 := tmp4471
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4472)
}
__typedArg0 := tmp4472
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4473)
}
__typedArg0 := tmp4473
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4474)
}
__typedArg0 := tmp4474
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4475)
}
__typedArg0 := Nil
__typedArg1 := tmp4475
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4460 Obj

if True == tmp4476 {
tmp4468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4468)
}
__typedArg0 := tmp4468
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4469)
}
__typedArg0 := tmp4469
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4462 Obj

if True == tmp4470 {
tmp4464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4464)
}
__typedArg0 := tmp4464
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4465)
}
__typedArg0 := tmp4465
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4466)
}
__typedArg0 := Nil
__typedArg1 := tmp4466
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4463 Obj

if True == tmp4467 {
ifres4463 = True


} else {
ifres4463 = False


}

ifres4462 = ifres4463


} else {
ifres4462 = False


}

var ifres4461 Obj

if True == ifres4462 {
ifres4461 = True


} else {
ifres4461 = False


}

ifres4460 = ifres4461


} else {
ifres4460 = False


}

var ifres4459 Obj

if True == ifres4460 {
ifres4459 = True


} else {
ifres4459 = False


}

ifres4458 = ifres4459


} else {
ifres4458 = False


}

var ifres4457 Obj

if True == ifres4458 {
ifres4457 = True


} else {
ifres4457 = False


}

ifres4456 = ifres4457


} else {
ifres4456 = False


}

var ifres4455 Obj

if True == ifres4456 {
ifres4455 = True


} else {
ifres4455 = False


}

ifres4454 = ifres4455


} else {
ifres4454 = False


}

var ifres4453 Obj

if True == ifres4454 {
ifres4453 = True


} else {
ifres4453 = False


}

ifres4452 = ifres4453


} else {
ifres4452 = False


}

var ifres4451 Obj

if True == ifres4452 {
ifres4451 = True


} else {
ifres4451 = False


}

ifres4450 = ifres4451


} else {
ifres4450 = False


}

var ifres4449 Obj

if True == ifres4450 {
ifres4449 = True


} else {
ifres4449 = False


}

ifres4448 = ifres4449


} else {
ifres4448 = False


}

if True == ifres4448 {
tmp4370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4370)
}
__typedArg0 := tmp4370
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4372 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4371)
}
__typedArg0 := tmp4371
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4372)
}
__typedArg0 := tmp4372
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4374 := Call(__e, PrimFunc(symshen_4pivot_1on), tmp4373, V852, Nil)


let__2762 := tmp4374
_ = let__2762

tmp4375 := Call(__e, PrimFunc(symfst), let__2762)


let__2763 := tmp4375
_ = let__2763

tmp4419 := Call(__e, PrimFunc(symshen_4bad_1pivot_2), let__2763)


if True == tmp4419 {
tmp4376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, Nil)
}
__typedArg0 := symif
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4377)
}
__typedArg0 := tmp4377
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4379)
}
__typedArg0 := tmp4379
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4380)
}
__typedArg0 := tmp4380
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4382 := Call(__e, PrimFunc(symshen_4recursively_1factor_1selectors), tmp4378, tmp4381)


tmp4383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4384 := Call(__e, PrimFunc(symshen_4factor_1recognisors), tmp4383)


tmp4385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4384, Nil)
}
__typedArg0 := tmp4384
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4386 := Call(__e, PrimFunc(symappend), tmp4382, tmp4385)


__e.TailApply(PrimFunc(symappend), tmp4376, tmp4386)
return


} else {
tmp4387 := Call(__e, PrimFunc(symsnd), let__2762)


let__2764 := tmp4387
_ = let__2764

tmp4388 := Call(__e, PrimFunc(symshen_4factor_1recognisors), let__2764)


let__2765 := tmp4388
_ = let__2765

tmp4389 := Call(__e, PrimFunc(symgensym), symGoTo)


let__2766 := tmp4389
_ = let__2766

tmp4390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2766, Nil)
}
__typedArg0 := let__2766
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp4390)
}
__typedArg0 := symthaw
__typedArg1 := tmp4390
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4391, Nil)
}
__typedArg0 := tmp4391
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp4392)
}
__typedArg0 := True
__typedArg1 := tmp4392
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4393, let__2763)
}
__typedArg0 := tmp4393
__typedArg1 := let__2763
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4395 := Call(__e, PrimFunc(symreverse), tmp4394)


let__2767 := tmp4395
_ = let__2767

tmp4396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2765, Nil)
}
__typedArg0 := let__2765
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp4396)
}
__typedArg0 := symfreeze
__typedArg1 := tmp4396
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4398)
}
__typedArg0 := tmp4398
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4399)
}
__typedArg0 := tmp4399
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4400)
}
__typedArg0 := tmp4400
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4402)
}
__typedArg0 := tmp4402
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4403)
}
__typedArg0 := tmp4403
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4404)
}
__typedArg0 := tmp4404
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4406 := Call(__e, PrimFunc(symshen_4factor_1recognisors), let__2767)


tmp4407 := Call(__e, PrimFunc(symshen_4factor_1selectors), tmp4405, tmp4406)


tmp4408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2766, Nil)
}
__typedArg0 := let__2766
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp4408)
}
__typedArg0 := symthaw
__typedArg1 := tmp4408
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4409, Nil)
}
__typedArg0 := tmp4409
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4407, tmp4410)
}
__typedArg0 := tmp4407
__typedArg1 := tmp4410
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4401, tmp4411)
}
__typedArg0 := tmp4401
__typedArg1 := tmp4411
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp4412)
}
__typedArg0 := symif
__typedArg1 := tmp4412
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4413, Nil)
}
__typedArg0 := tmp4413
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4397, tmp4414)
}
__typedArg0 := tmp4397
__typedArg1 := tmp4414
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2766, tmp4415)
}
__typedArg0 := let__2766
__typedArg1 := tmp4415
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp4416)
}
__typedArg0 := symlet
__typedArg1 := tmp4416
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__2768 := tmp4417
_ = let__2768

__e.TailApply(PrimFunc(symshen_4remove_1indirection), let__2768)
return


}


} else {
tmp4446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4431 Obj

if True == tmp4446 {
tmp4444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4444)
}
__typedArg0 := tmp4444
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4433 Obj

if True == tmp4445 {
tmp4441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4441)
}
__typedArg0 := tmp4441
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4442)
}
__typedArg0 := tmp4442
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4435 Obj

if True == tmp4443 {
tmp4437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4437)
}
__typedArg0 := tmp4437
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4438)
}
__typedArg0 := tmp4438
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4439)
}
__typedArg0 := Nil
__typedArg1 := tmp4439
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4436 Obj

if True == tmp4440 {
ifres4436 = True


} else {
ifres4436 = False


}

ifres4435 = ifres4436


} else {
ifres4435 = False


}

var ifres4434 Obj

if True == ifres4435 {
ifres4434 = True


} else {
ifres4434 = False


}

ifres4433 = ifres4434


} else {
ifres4433 = False


}

var ifres4432 Obj

if True == ifres4433 {
ifres4432 = True


} else {
ifres4432 = False


}

ifres4431 = ifres4432


} else {
ifres4431 = False


}

if True == ifres4431 {
tmp4420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4420)
}
__typedArg0 := tmp4420
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4422)
}
__typedArg0 := tmp4422
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4423)
}
__typedArg0 := tmp4423
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4426 := Call(__e, PrimFunc(symshen_4factor_1recognisors), tmp4425)


tmp4427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4426, Nil)
}
__typedArg0 := tmp4426
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4424, tmp4427)
}
__typedArg0 := tmp4424
__typedArg1 := tmp4427
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4421, tmp4428)
}
__typedArg0 := tmp4421
__typedArg1 := tmp4428
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp4429)
}
__typedArg0 := symif
__typedArg1 := tmp4429
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.factor-recognisors"))
}
__typedArg0 := MakeString("partial function shen.factor-recognisors")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}


}, 1)

tmp4518 := Call(__e, ns2_1set, symshen_4factor_1recognisors, tmp4367)


_ = tmp4518

tmp4519 := MakeNative(func(__e *ControlFlow) {
V860 := __e.Get(1)
_ = V860
V861 := __e.Get(2)
_ = V861
tmp4591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4565 Obj

if True == tmp4591 {
tmp4589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp4589)
}
__typedArg0 := symlet
__typedArg1 := tmp4589
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4567 Obj

if True == tmp4590 {
tmp4587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4587)
}
__typedArg0 := tmp4587
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4569 Obj

if True == tmp4588 {
tmp4584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4584)
}
__typedArg0 := tmp4584
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4585)
}
__typedArg0 := tmp4585
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4571 Obj

if True == tmp4586 {
tmp4580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4580)
}
__typedArg0 := tmp4580
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4581)
}
__typedArg0 := tmp4581
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4582)
}
__typedArg0 := tmp4582
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4573 Obj

if True == tmp4583 {
tmp4575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4575)
}
__typedArg0 := tmp4575
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4576)
}
__typedArg0 := tmp4576
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4577)
}
__typedArg0 := tmp4577
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4578)
}
__typedArg0 := Nil
__typedArg1 := tmp4578
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4574 Obj

if True == tmp4579 {
ifres4574 = True


} else {
ifres4574 = False


}

ifres4573 = ifres4574


} else {
ifres4573 = False


}

var ifres4572 Obj

if True == ifres4573 {
ifres4572 = True


} else {
ifres4572 = False


}

ifres4571 = ifres4572


} else {
ifres4571 = False


}

var ifres4570 Obj

if True == ifres4571 {
ifres4570 = True


} else {
ifres4570 = False


}

ifres4569 = ifres4570


} else {
ifres4569 = False


}

var ifres4568 Obj

if True == ifres4569 {
ifres4568 = True


} else {
ifres4568 = False


}

ifres4567 = ifres4568


} else {
ifres4567 = False


}

var ifres4566 Obj

if True == ifres4567 {
ifres4566 = True


} else {
ifres4566 = False


}

ifres4565 = ifres4566


} else {
ifres4565 = False


}

if True == ifres4565 {
tmp4520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4520)
}
__typedArg0 := tmp4520
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4522)
}
__typedArg0 := tmp4522
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4523)
}
__typedArg0 := tmp4523
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4525)
}
__typedArg0 := tmp4525
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4526)
}
__typedArg0 := tmp4526
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4527)
}
__typedArg0 := tmp4527
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4529 := Call(__e, PrimFunc(symshen_4recursively_1factor_1selectors), tmp4528, V861)


__e.TailApply(PrimFunc(symshen_4restore_1local), tmp4521, tmp4524, tmp4529)
return


} else {
tmp4563 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4544 Obj

if True == tmp4563 {
tmp4561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symand, tmp4561)
}
__typedArg0 := symand
__typedArg1 := tmp4561
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4546 Obj

if True == tmp4562 {
tmp4559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4559)
}
__typedArg0 := tmp4559
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4548 Obj

if True == tmp4560 {
tmp4556 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4556)
}
__typedArg0 := tmp4556
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4557)
}
__typedArg0 := tmp4557
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4550 Obj

if True == tmp4558 {
tmp4552 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4553 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4552)
}
__typedArg0 := tmp4552
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4554 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4553)
}
__typedArg0 := tmp4553
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4554)
}
__typedArg0 := Nil
__typedArg1 := tmp4554
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4551 Obj

if True == tmp4555 {
ifres4551 = True


} else {
ifres4551 = False


}

ifres4550 = ifres4551


} else {
ifres4550 = False


}

var ifres4549 Obj

if True == ifres4550 {
ifres4549 = True


} else {
ifres4549 = False


}

ifres4548 = ifres4549


} else {
ifres4548 = False


}

var ifres4547 Obj

if True == ifres4548 {
ifres4547 = True


} else {
ifres4547 = False


}

ifres4546 = ifres4547


} else {
ifres4546 = False


}

var ifres4545 Obj

if True == ifres4546 {
ifres4545 = True


} else {
ifres4545 = False


}

ifres4544 = ifres4545


} else {
ifres4544 = False


}

if True == ifres4544 {
tmp4530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4530)
}
__typedArg0 := tmp4530
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4532)
}
__typedArg0 := tmp4532
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4534)
}
__typedArg0 := tmp4534
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4536 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4535)
}
__typedArg0 := tmp4535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4537 := Call(__e, PrimFunc(symshen_4factor_1selectors), tmp4533, tmp4536)


tmp4538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4538)
}
__typedArg0 := tmp4538
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4540 := Call(__e, PrimFunc(symshen_4factor_1selectors), tmp4539, V861)


tmp4541 := Call(__e, PrimFunc(symshen_4recursively_1factor_1selectors), tmp4537, tmp4540)


__e.TailApply(PrimFunc(symshen_4restore_1P), tmp4531, tmp4541)
return


} else {
tmp4542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V861, Nil)
}
__typedArg0 := V861
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V860, tmp4542)
}
__typedArg0 := V860
__typedArg1 := tmp4542
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}


}, 2)

tmp4592 := Call(__e, ns2_1set, symshen_4recursively_1factor_1selectors, tmp4519)


_ = tmp4592

tmp4593 := MakeNative(func(__e *ControlFlow) {
V862 := __e.Get(1)
_ = V862
V863 := __e.Get(2)
_ = V863
tmp4609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V863)
}
__typedArg0 := V863
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4600 Obj

if True == tmp4609 {
tmp4607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V863)
}
__typedArg0 := V863
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4607)
}
__typedArg0 := tmp4607
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4602 Obj

if True == tmp4608 {
tmp4604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V863)
}
__typedArg0 := V863
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4604)
}
__typedArg0 := tmp4604
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4605)
}
__typedArg0 := Nil
__typedArg1 := tmp4605
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4603 Obj

if True == tmp4606 {
ifres4603 = True


} else {
ifres4603 = False


}

ifres4602 = ifres4603


} else {
ifres4602 = False


}

var ifres4601 Obj

if True == ifres4602 {
ifres4601 = True


} else {
ifres4601 = False


}

ifres4600 = ifres4601


} else {
ifres4600 = False


}

if True == ifres4600 {
tmp4594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V863)
}
__typedArg0 := V863
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4594, Nil)
}
__typedArg0 := tmp4594
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V862, tmp4595)
}
__typedArg0 := V862
__typedArg1 := tmp4595
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp4596)
}
__typedArg0 := symand
__typedArg1 := tmp4596
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V863)
}
__typedArg0 := V863
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4597, tmp4598)
}
__typedArg0 := tmp4597
__typedArg1 := tmp4598
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.restore-P"))
}
__typedArg0 := MakeString("partial function shen.restore-P")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 2)

tmp4610 := Call(__e, ns2_1set, symshen_4restore_1P, tmp4593)


_ = tmp4610

tmp4611 := MakeNative(func(__e *ControlFlow) {
V864 := __e.Get(1)
_ = V864
V865 := __e.Get(2)
_ = V865
V866 := __e.Get(3)
_ = V866
tmp4628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V866)
}
__typedArg0 := V866
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4619 Obj

if True == tmp4628 {
tmp4626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V866)
}
__typedArg0 := V866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4626)
}
__typedArg0 := tmp4626
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4621 Obj

if True == tmp4627 {
tmp4623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V866)
}
__typedArg0 := V866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4623)
}
__typedArg0 := tmp4623
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4624)
}
__typedArg0 := Nil
__typedArg1 := tmp4624
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4622 Obj

if True == tmp4625 {
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

if True == ifres4619 {
tmp4612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V866)
}
__typedArg0 := V866
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4612, Nil)
}
__typedArg0 := tmp4612
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V865, tmp4613)
}
__typedArg0 := V865
__typedArg1 := tmp4613
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V864, tmp4614)
}
__typedArg0 := V864
__typedArg1 := tmp4614
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp4615)
}
__typedArg0 := symlet
__typedArg1 := tmp4615
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V866)
}
__typedArg0 := V866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4616, tmp4617)
}
__typedArg0 := tmp4616
__typedArg1 := tmp4617
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.restore-local"))
}
__typedArg0 := MakeString("partial function shen.restore-local")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}, 3)

tmp4629 := Call(__e, ns2_1set, symshen_4restore_1local, tmp4611)


_ = tmp4629

tmp4630 := MakeNative(func(__e *ControlFlow) {
V871 := __e.Get(1)
_ = V871
tmp4636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V871)
}
__typedArg0 := V871
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4632 Obj

if True == tmp4636 {
tmp4634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V871)
}
__typedArg0 := V871
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4634)
}
__typedArg0 := Nil
__typedArg1 := tmp4634
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4633 Obj

if True == tmp4635 {
ifres4633 = True


} else {
ifres4633 = False


}

ifres4632 = ifres4633


} else {
ifres4632 = False


}

if True == ifres4632 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp4637 := Call(__e, ns2_1set, symshen_4bad_1pivot_2, tmp4630)


_ = tmp4637

tmp4638 := MakeNative(func(__e *ControlFlow) {
V872 := __e.Get(1)
_ = V872
tmp4753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4653 Obj

if True == tmp4753 {
tmp4751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp4751)
}
__typedArg0 := symlet
__typedArg1 := tmp4751
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4655 Obj

if True == tmp4752 {
tmp4749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4749)
}
__typedArg0 := tmp4749
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4657 Obj

if True == tmp4750 {
tmp4746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4746)
}
__typedArg0 := tmp4746
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4747)
}
__typedArg0 := tmp4747
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4659 Obj

if True == tmp4748 {
tmp4742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4742)
}
__typedArg0 := tmp4742
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4743)
}
__typedArg0 := tmp4743
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4744)
}
__typedArg0 := tmp4744
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4661 Obj

if True == tmp4745 {
tmp4737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4737)
}
__typedArg0 := tmp4737
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4738)
}
__typedArg0 := tmp4738
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4739)
}
__typedArg0 := tmp4739
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfreeze, tmp4740)
}
__typedArg0 := symfreeze
__typedArg1 := tmp4740
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4663 Obj

if True == tmp4741 {
tmp4732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4732)
}
__typedArg0 := tmp4732
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4733)
}
__typedArg0 := tmp4733
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4734)
}
__typedArg0 := tmp4734
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4735)
}
__typedArg0 := tmp4735
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4665 Obj

if True == tmp4736 {
tmp4726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4726)
}
__typedArg0 := tmp4726
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4727)
}
__typedArg0 := tmp4727
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4728)
}
__typedArg0 := tmp4728
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4729)
}
__typedArg0 := tmp4729
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4730)
}
__typedArg0 := tmp4730
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4667 Obj

if True == tmp4731 {
tmp4719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4719)
}
__typedArg0 := tmp4719
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4720)
}
__typedArg0 := tmp4720
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4721)
}
__typedArg0 := tmp4721
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4722)
}
__typedArg0 := tmp4722
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4723)
}
__typedArg0 := tmp4723
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symthaw, tmp4724)
}
__typedArg0 := symthaw
__typedArg1 := tmp4724
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4669 Obj

if True == tmp4725 {
tmp4712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4712)
}
__typedArg0 := tmp4712
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4713)
}
__typedArg0 := tmp4713
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4714)
}
__typedArg0 := tmp4714
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4715)
}
__typedArg0 := tmp4715
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4716)
}
__typedArg0 := tmp4716
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4717)
}
__typedArg0 := tmp4717
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4671 Obj

if True == tmp4718 {
tmp4704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4705)
}
__typedArg0 := tmp4705
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4706)
}
__typedArg0 := tmp4706
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4707)
}
__typedArg0 := tmp4707
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4708)
}
__typedArg0 := tmp4708
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4709)
}
__typedArg0 := tmp4709
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4710)
}
__typedArg0 := Nil
__typedArg1 := tmp4710
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4673 Obj

if True == tmp4711 {
tmp4698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4698)
}
__typedArg0 := tmp4698
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4699)
}
__typedArg0 := tmp4699
return Call(__e, PrimFunc(symhd), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4702)
}
__typedArg0 := Nil
__typedArg1 := tmp4702
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4675 Obj

if True == tmp4703 {
tmp4694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4694)
}
__typedArg0 := tmp4694
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4696)
}
__typedArg0 := tmp4696
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4677 Obj

if True == tmp4697 {
tmp4689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4689)
}
__typedArg0 := tmp4689
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4690)
}
__typedArg0 := tmp4690
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4691)
}
__typedArg0 := tmp4691
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4692)
}
__typedArg0 := Nil
__typedArg1 := tmp4692
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4679 Obj

if True == tmp4693 {
tmp4681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4681)
}
__typedArg0 := tmp4681
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4682)
}
__typedArg0 := tmp4682
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4683)
}
__typedArg0 := tmp4683
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4684)
}
__typedArg0 := tmp4684
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4685)
}
__typedArg0 := tmp4685
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4686)
}
__typedArg0 := tmp4686
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(tmp4687)
}
__typedArg0 := tmp4687
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

var ifres4680 Obj

if True == tmp4688 {
ifres4680 = True


} else {
ifres4680 = False


}

ifres4679 = ifres4680


} else {
ifres4679 = False


}

var ifres4678 Obj

if True == ifres4679 {
ifres4678 = True


} else {
ifres4678 = False


}

ifres4677 = ifres4678


} else {
ifres4677 = False


}

var ifres4676 Obj

if True == ifres4677 {
ifres4676 = True


} else {
ifres4676 = False


}

ifres4675 = ifres4676


} else {
ifres4675 = False


}

var ifres4674 Obj

if True == ifres4675 {
ifres4674 = True


} else {
ifres4674 = False


}

ifres4673 = ifres4674


} else {
ifres4673 = False


}

var ifres4672 Obj

if True == ifres4673 {
ifres4672 = True


} else {
ifres4672 = False


}

ifres4671 = ifres4672


} else {
ifres4671 = False


}

var ifres4670 Obj

if True == ifres4671 {
ifres4670 = True


} else {
ifres4670 = False


}

ifres4669 = ifres4670


} else {
ifres4669 = False


}

var ifres4668 Obj

if True == ifres4669 {
ifres4668 = True


} else {
ifres4668 = False


}

ifres4667 = ifres4668


} else {
ifres4667 = False


}

var ifres4666 Obj

if True == ifres4667 {
ifres4666 = True


} else {
ifres4666 = False


}

ifres4665 = ifres4666


} else {
ifres4665 = False


}

var ifres4664 Obj

if True == ifres4665 {
ifres4664 = True


} else {
ifres4664 = False


}

ifres4663 = ifres4664


} else {
ifres4663 = False


}

var ifres4662 Obj

if True == ifres4663 {
ifres4662 = True


} else {
ifres4662 = False


}

ifres4661 = ifres4662


} else {
ifres4661 = False


}

var ifres4660 Obj

if True == ifres4661 {
ifres4660 = True


} else {
ifres4660 = False


}

ifres4659 = ifres4660


} else {
ifres4659 = False


}

var ifres4658 Obj

if True == ifres4659 {
ifres4658 = True


} else {
ifres4658 = False


}

ifres4657 = ifres4658


} else {
ifres4657 = False


}

var ifres4656 Obj

if True == ifres4657 {
ifres4656 = True


} else {
ifres4656 = False


}

ifres4655 = ifres4656


} else {
ifres4655 = False


}

var ifres4654 Obj

if True == ifres4655 {
ifres4654 = True


} else {
ifres4654 = False


}

ifres4653 = ifres4654


} else {
ifres4653 = False


}

if True == ifres4653 {
tmp4639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4639)
}
__typedArg0 := tmp4639
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
return PrimTail(tmp4641)
}
__typedArg0 := tmp4641
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4643 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4642)
}
__typedArg0 := tmp4642
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4644 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4643)
}
__typedArg0 := tmp4643
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4644)
}
__typedArg0 := tmp4644
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4646)
}
__typedArg0 := tmp4646
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4648)
}
__typedArg0 := tmp4648
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4649)
}
__typedArg0 := tmp4649
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4650)
}
__typedArg0 := tmp4650
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symsubst), tmp4645, tmp4647, tmp4651)
return


} else {
__e.Return(V872)
return
}


}, 1)

tmp4754 := Call(__e, ns2_1set, symshen_4remove_1indirection, tmp4638)


_ = tmp4754

tmp4755 := MakeNative(func(__e *ControlFlow) {
V875 := __e.Get(1)
_ = V875
V876 := __e.Get(2)
_ = V876
V877 := __e.Get(3)
_ = V877
tmp4854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4800 Obj

if True == tmp4854 {
tmp4852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4852)
}
__typedArg0 := tmp4852
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4802 Obj

if True == tmp4853 {
tmp4849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
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

var ifres4804 Obj

if True == tmp4851 {
tmp4845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4845)
}
__typedArg0 := tmp4845
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4846)
}
__typedArg0 := tmp4846
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symand, tmp4847)
}
__typedArg0 := symand
__typedArg1 := tmp4847
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4806 Obj

if True == tmp4848 {
tmp4841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4841)
}
__typedArg0 := tmp4841
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4842)
}
__typedArg0 := tmp4842
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4843)
}
__typedArg0 := tmp4843
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4808 Obj

if True == tmp4844 {
tmp4836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4836)
}
__typedArg0 := tmp4836
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4837)
}
__typedArg0 := tmp4837
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4838)
}
__typedArg0 := tmp4838
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4839)
}
__typedArg0 := tmp4839
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4810 Obj

if True == tmp4840 {
tmp4830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4830)
}
__typedArg0 := tmp4830
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4831)
}
__typedArg0 := tmp4831
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4833)
}
__typedArg0 := tmp4833
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4834)
}
__typedArg0 := Nil
__typedArg1 := tmp4834
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4812 Obj

if True == tmp4835 {
tmp4827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4827)
}
__typedArg0 := tmp4827
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4828)
}
__typedArg0 := tmp4828
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4814 Obj

if True == tmp4829 {
tmp4823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4823)
}
__typedArg0 := tmp4823
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4825)
}
__typedArg0 := Nil
__typedArg1 := tmp4825
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4816 Obj

if True == tmp4826 {
tmp4818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4818)
}
__typedArg0 := tmp4818
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4819)
}
__typedArg0 := tmp4819
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4820)
}
__typedArg0 := tmp4820
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V875, tmp4821)
}
__typedArg0 := V875
__typedArg1 := tmp4821
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4817 Obj

if True == tmp4822 {
ifres4817 = True


} else {
ifres4817 = False


}

ifres4816 = ifres4817


} else {
ifres4816 = False


}

var ifres4815 Obj

if True == ifres4816 {
ifres4815 = True


} else {
ifres4815 = False


}

ifres4814 = ifres4815


} else {
ifres4814 = False


}

var ifres4813 Obj

if True == ifres4814 {
ifres4813 = True


} else {
ifres4813 = False


}

ifres4812 = ifres4813


} else {
ifres4812 = False


}

var ifres4811 Obj

if True == ifres4812 {
ifres4811 = True


} else {
ifres4811 = False


}

ifres4810 = ifres4811


} else {
ifres4810 = False


}

var ifres4809 Obj

if True == ifres4810 {
ifres4809 = True


} else {
ifres4809 = False


}

ifres4808 = ifres4809


} else {
ifres4808 = False


}

var ifres4807 Obj

if True == ifres4808 {
ifres4807 = True


} else {
ifres4807 = False


}

ifres4806 = ifres4807


} else {
ifres4806 = False


}

var ifres4805 Obj

if True == ifres4806 {
ifres4805 = True


} else {
ifres4805 = False


}

ifres4804 = ifres4805


} else {
ifres4804 = False


}

var ifres4803 Obj

if True == ifres4804 {
ifres4803 = True


} else {
ifres4803 = False


}

ifres4802 = ifres4803


} else {
ifres4802 = False


}

var ifres4801 Obj

if True == ifres4802 {
ifres4801 = True


} else {
ifres4801 = False


}

ifres4800 = ifres4801


} else {
ifres4800 = False


}

if True == ifres4800 {
tmp4756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4756)
}
__typedArg0 := tmp4756
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4757)
}
__typedArg0 := tmp4757
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4758)
}
__typedArg0 := tmp4758
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4762 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4761)
}
__typedArg0 := tmp4761
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4763 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4762)
}
__typedArg0 := tmp4762
return Call(__e, PrimFunc(symtl), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4766)
}
__typedArg0 := tmp4766
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4765, tmp4767)
}
__typedArg0 := tmp4765
__typedArg1 := tmp4767
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4768, V877)
}
__typedArg0 := tmp4768
__typedArg1 := V877
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4pivot_1on), tmp4759, tmp4760, tmp4769)
return


} else {
tmp4798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4778 Obj

if True == tmp4798 {
tmp4796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4796)
}
__typedArg0 := tmp4796
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4780 Obj

if True == tmp4797 {
tmp4793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4793)
}
__typedArg0 := tmp4793
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4794)
}
__typedArg0 := tmp4794
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4782 Obj

if True == tmp4795 {
tmp4789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4789)
}
__typedArg0 := tmp4789
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4790)
}
__typedArg0 := tmp4790
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4791)
}
__typedArg0 := Nil
__typedArg1 := tmp4791
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4784 Obj

if True == tmp4792 {
tmp4786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4786)
}
__typedArg0 := tmp4786
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V875, tmp4787)
}
__typedArg0 := V875
__typedArg1 := tmp4787
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4785 Obj

if True == tmp4788 {
ifres4785 = True


} else {
ifres4785 = False


}

ifres4784 = ifres4785


} else {
ifres4784 = False


}

var ifres4783 Obj

if True == ifres4784 {
ifres4783 = True


} else {
ifres4783 = False


}

ifres4782 = ifres4783


} else {
ifres4782 = False


}

var ifres4781 Obj

if True == ifres4782 {
ifres4781 = True


} else {
ifres4781 = False


}

ifres4780 = ifres4781


} else {
ifres4780 = False


}

var ifres4779 Obj

if True == ifres4780 {
ifres4779 = True


} else {
ifres4779 = False


}

ifres4778 = ifres4779


} else {
ifres4778 = False


}

if True == ifres4778 {
tmp4770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4770)
}
__typedArg0 := tmp4770
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4773)
}
__typedArg0 := tmp4773
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp4774)
}
__typedArg0 := True
__typedArg1 := tmp4774
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4775, V877)
}
__typedArg0 := tmp4775
__typedArg1 := V877
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4pivot_1on), tmp4771, tmp4772, tmp4776)
return


} else {
__e.TailApply(PrimFunc(sym_8p), V877, V876)
return
}


}


}, 3)

tmp4855 := Call(__e, ns2_1set, symshen_4pivot_1on, tmp4755)


_ = tmp4855

tmp4856 := MakeNative(func(__e *ControlFlow) {
V880 := __e.Get(1)
_ = V880
V881 := __e.Get(2)
_ = V881
tmp4879 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V880)
}
__typedArg0 := V880
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4870 Obj

if True == tmp4879 {
tmp4877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V880)
}
__typedArg0 := V880
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4877)
}
__typedArg0 := tmp4877
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4872 Obj

if True == tmp4878 {
tmp4874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V880)
}
__typedArg0 := V880
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4874)
}
__typedArg0 := tmp4874
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4875)
}
__typedArg0 := Nil
__typedArg1 := tmp4875
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4873 Obj

if True == tmp4876 {
ifres4873 = True


} else {
ifres4873 = False


}

ifres4872 = ifres4873


} else {
ifres4872 = False


}

var ifres4871 Obj

if True == ifres4872 {
ifres4871 = True


} else {
ifres4871 = False


}

ifres4870 = ifres4871


} else {
ifres4870 = False


}

if True == ifres4870 {
tmp4857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V880)
}
__typedArg0 := V880
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4858 := Call(__e, PrimFunc(symshen_4op), tmp4857)


let__2769 := tmp4858
_ = let__2769

tmp4868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4skip, let__2769)
}
__typedArg0 := symshen_4skip
__typedArg1 := let__2769
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4868 {
__e.Return(V881)
return
} else {
tmp4859 := Call(__e, PrimFunc(symshen_4op1), let__2769)


tmp4860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V880)
}
__typedArg0 := V880
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4859, tmp4860)
}
__typedArg0 := tmp4859
__typedArg1 := tmp4860
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4862 := Call(__e, PrimFunc(symshen_4op2), let__2769)


tmp4863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V880)
}
__typedArg0 := V880
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4862, tmp4863)
}
__typedArg0 := tmp4862
__typedArg1 := tmp4863
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4864, Nil)
}
__typedArg0 := tmp4864
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4861, tmp4865)
}
__typedArg0 := tmp4861
__typedArg1 := tmp4865
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4factor_1selectors_1h), tmp4866, V881)
return


}


} else {
__e.Return(V881)
return
}


}, 2)

tmp4880 := Call(__e, ns2_1set, symshen_4factor_1selectors, tmp4856)


_ = tmp4880

tmp4881 := MakeNative(func(__e *ControlFlow) {
V885 := __e.Get(1)
_ = V885
tmp4889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons_2, V885)
}
__typedArg0 := symcons_2
__typedArg1 := V885
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4889 {
__e.Return(symcons)
return
} else {
tmp4887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7string_2, V885)
}
__typedArg0 := symshen_4_7string_2
__typedArg1 := V885
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4887 {
__e.Return(sym_8s)
return
} else {
tmp4885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7vector_2, V885)
}
__typedArg0 := symshen_4_7vector_2
__typedArg1 := V885
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4885 {
__e.Return(sym_8v)
return
} else {
tmp4883 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtuple_2, V885)
}
__typedArg0 := symtuple_2
__typedArg1 := V885
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4883 {
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

tmp4890 := Call(__e, ns2_1set, symshen_4op, tmp4881)


_ = tmp4890

tmp4891 := MakeNative(func(__e *ControlFlow) {
V886 := __e.Get(1)
_ = V886
V887 := __e.Get(2)
_ = V887
tmp4909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V886)
}
__typedArg0 := Nil
__typedArg1 := V886
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4909 {
__e.Return(V887)
return
} else {
tmp4907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V886)
}
__typedArg0 := V886
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp4907 {
tmp4903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V886)
}
__typedArg0 := V886
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4904 := Call(__e, PrimFunc(symoccurrences), tmp4903, V887)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(tmp4904)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := tmp4904
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
tmp4892 := Call(__e, PrimFunc(symgensym), symSelect)


let__2770 := tmp4892
_ = let__2770

tmp4893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V886)
}
__typedArg0 := V886
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V886)
}
__typedArg0 := V886
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V886)
}
__typedArg0 := V886
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4896 := Call(__e, PrimFunc(symsubst), let__2770, tmp4895, V887)


tmp4897 := Call(__e, PrimFunc(symshen_4factor_1selectors_1h), tmp4894, tmp4896)


tmp4898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4897, Nil)
}
__typedArg0 := tmp4897
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4893, tmp4898)
}
__typedArg0 := tmp4893
__typedArg1 := tmp4898
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2770, tmp4899)
}
__typedArg0 := let__2770
__typedArg1 := tmp4899
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp4900)
}
__typedArg0 := symlet
__typedArg1 := tmp4900
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp4901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V886)
}
__typedArg0 := V886
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4factor_1selectors_1h), tmp4901, V887)
return


}


} else {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(MakeString("partial function shen.factor-selectors-h"))
}
__typedArg0 := MakeString("partial function shen.factor-selectors-h")
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return
}


}


}, 2)

__e.TailApply(ns2_1set, symshen_4factor_1selectors_1h, tmp4891)
return




}, 0)

