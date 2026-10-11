package main

import . "github.com/pyrex41/shen-go/kl"

var CoreMain = MakeNative(func(__e *ControlFlow) {
tmp2733 := MakeNative(func(__e *ControlFlow) {
V540 := __e.Get(1)
_ = V540
tmp2734 := Call(__e, PrimFunc(symshen_4shen_1_6kl_1h), V540)


let__2581 := tmp2734
_ = let__2581

__e.TailApply(PrimFunc(symshen_4record_1and_1evaluate), let__2581)
return


}, 1)

tmp2735 := Call(__e, ns2_1set, symshen_4shen_1_6kl, tmp2733)


_ = tmp2735

tmp2736 := MakeNative(func(__e *ControlFlow) {
V542 := __e.Get(1)
_ = V542
tmp2785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V542)
}
__typedArg0 := V542
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2759 Obj

if True == tmp2785 {
tmp2783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V542)
}
__typedArg0 := V542
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp2783)
}
__typedArg0 := symdefun
__typedArg1 := tmp2783
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2761 Obj

if True == tmp2784 {
tmp2781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V542)
}
__typedArg0 := V542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2781)
}
__typedArg0 := tmp2781
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2763 Obj

if True == tmp2782 {
tmp2778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V542)
}
__typedArg0 := V542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2778)
}
__typedArg0 := tmp2778
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2779)
}
__typedArg0 := tmp2779
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2765 Obj

if True == tmp2780 {
tmp2774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V542)
}
__typedArg0 := V542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2774)
}
__typedArg0 := tmp2774
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2775)
}
__typedArg0 := tmp2775
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2776)
}
__typedArg0 := tmp2776
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2767 Obj

if True == tmp2777 {
tmp2769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V542)
}
__typedArg0 := V542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2769)
}
__typedArg0 := tmp2769
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2770)
}
__typedArg0 := tmp2770
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2771)
}
__typedArg0 := tmp2771
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2772)
}
__typedArg0 := Nil
__typedArg1 := tmp2772
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2768 Obj

if True == tmp2773 {
ifres2768 = True


} else {
ifres2768 = False


}

ifres2767 = ifres2768


} else {
ifres2767 = False


}

var ifres2766 Obj

if True == ifres2767 {
ifres2766 = True


} else {
ifres2766 = False


}

ifres2765 = ifres2766


} else {
ifres2765 = False


}

var ifres2764 Obj

if True == ifres2765 {
ifres2764 = True


} else {
ifres2764 = False


}

ifres2763 = ifres2764


} else {
ifres2763 = False


}

var ifres2762 Obj

if True == ifres2763 {
ifres2762 = True


} else {
ifres2762 = False


}

ifres2761 = ifres2762


} else {
ifres2761 = False


}

var ifres2760 Obj

if True == ifres2761 {
ifres2760 = True


} else {
ifres2760 = False


}

ifres2759 = ifres2760


} else {
ifres2759 = False


}

if True == ifres2759 {
tmp2742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V542)
}
__typedArg0 := V542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2742)
}
__typedArg0 := tmp2742
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2744 := Call(__e, PrimFunc(symshen_4sysfunc_2), tmp2743)


var ifres2737 Obj

if True == tmp2744 {
tmp2738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V542)
}
__typedArg0 := V542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2738)
}
__typedArg0 := tmp2738
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2740 := Call(__e, PrimFunc(symshen_4app), tmp2739, MakeString(" is not a legitimate function name\n"), symshen_4a)


tmp2741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp2740)
}
__typedArg0 := tmp2740
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres2737 = tmp2741


} else {
ifres2737 = symshen_4skip


}

let__2582 := ifres2737
_ = let__2582

tmp2745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V542)
}
__typedArg0 := V542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2745)
}
__typedArg0 := tmp2745
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V542)
}
__typedArg0 := V542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2747)
}
__typedArg0 := tmp2747
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2748)
}
__typedArg0 := tmp2748
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2750 := Call(__e, PrimFunc(symlength), tmp2749)


tmp2751 := Call(__e, PrimFunc(symshen_4store_1arity), tmp2746, tmp2750)


let__2583 := tmp2751
_ = let__2583

tmp2752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V542)
}
__typedArg0 := V542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2752)
}
__typedArg0 := tmp2752
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2754 := Call(__e, PrimFunc(symshen_4record_1kl), tmp2753, V542)


let__2584 := tmp2754
_ = let__2584

tmp2755 := Call(__e, PrimFunc(symeval_1kl), V542)


let__2585 := tmp2755
_ = let__2585

tmp2756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V542)
}
__typedArg0 := V542
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2756)
}
__typedArg0 := tmp2756
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4fn_1print), tmp2757)
return


} else {
__e.Return(V542)
return
}


}, 1)

tmp2786 := Call(__e, ns2_1set, symshen_4record_1and_1evaluate, tmp2736)


_ = tmp2786

tmp2787 := MakeNative(func(__e *ControlFlow) {
V547 := __e.Get(1)
_ = V547
tmp2888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2880 Obj

if True == tmp2888 {
tmp2886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp2886)
}
__typedArg0 := symdefine
__typedArg1 := tmp2886
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2882 Obj

if True == tmp2887 {
tmp2884 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2885 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2884)
}
__typedArg0 := tmp2884
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2883 Obj

if True == tmp2885 {
ifres2883 = True


} else {
ifres2883 = False


}

ifres2882 = ifres2883


} else {
ifres2882 = False


}

var ifres2881 Obj

if True == ifres2882 {
ifres2881 = True


} else {
ifres2881 = False


}

ifres2880 = ifres2881


} else {
ifres2880 = False


}

if True == ifres2880 {
tmp2788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2788)
}
__typedArg0 := tmp2788
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2790)
}
__typedArg0 := tmp2790
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4shendef_1_6kldef), tmp2789, tmp2791)
return


} else {
tmp2878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2852 Obj

if True == tmp2878 {
tmp2876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp2876)
}
__typedArg0 := symdefun
__typedArg1 := tmp2876
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2854 Obj

if True == tmp2877 {
tmp2874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2874)
}
__typedArg0 := tmp2874
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2856 Obj

if True == tmp2875 {
tmp2871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2871)
}
__typedArg0 := tmp2871
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2872)
}
__typedArg0 := tmp2872
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2858 Obj

if True == tmp2873 {
tmp2867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2867)
}
__typedArg0 := tmp2867
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2868)
}
__typedArg0 := tmp2868
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2869)
}
__typedArg0 := tmp2869
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2860 Obj

if True == tmp2870 {
tmp2862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2862)
}
__typedArg0 := tmp2862
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2863)
}
__typedArg0 := tmp2863
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2864)
}
__typedArg0 := tmp2864
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2865)
}
__typedArg0 := Nil
__typedArg1 := tmp2865
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2861 Obj

if True == tmp2866 {
ifres2861 = True


} else {
ifres2861 = False


}

ifres2860 = ifres2861


} else {
ifres2860 = False


}

var ifres2859 Obj

if True == ifres2860 {
ifres2859 = True


} else {
ifres2859 = False


}

ifres2858 = ifres2859


} else {
ifres2858 = False


}

var ifres2857 Obj

if True == ifres2858 {
ifres2857 = True


} else {
ifres2857 = False


}

ifres2856 = ifres2857


} else {
ifres2856 = False


}

var ifres2855 Obj

if True == ifres2856 {
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

if True == ifres2852 {
__e.Return(V547)
return
} else {
tmp2850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2831 Obj

if True == tmp2850 {
tmp2848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtype, tmp2848)
}
__typedArg0 := symtype
__typedArg1 := tmp2848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2833 Obj

if True == tmp2849 {
tmp2846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2846)
}
__typedArg0 := tmp2846
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2835 Obj

if True == tmp2847 {
tmp2843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2843)
}
__typedArg0 := tmp2843
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2844)
}
__typedArg0 := tmp2844
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2837 Obj

if True == tmp2845 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2840)
}
__typedArg0 := tmp2840
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2841)
}
__typedArg0 := Nil
__typedArg1 := tmp2841
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2838 Obj

if True == tmp2842 {
ifres2838 = True


} else {
ifres2838 = False


}

ifres2837 = ifres2838


} else {
ifres2837 = False


}

var ifres2836 Obj

if True == ifres2837 {
ifres2836 = True


} else {
ifres2836 = False


}

ifres2835 = ifres2836


} else {
ifres2835 = False


}

var ifres2834 Obj

if True == ifres2835 {
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

if True == ifres2831 {
tmp2792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2792)
}
__typedArg0 := tmp2792
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2794)
}
__typedArg0 := tmp2794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2795)
}
__typedArg0 := tmp2795
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2797 := Call(__e, PrimFunc(symshen_4rcons__form), tmp2796)


tmp2798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2797, Nil)
}
__typedArg0 := tmp2797
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2793, tmp2798)
}
__typedArg0 := tmp2793
__typedArg1 := tmp2798
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp2799)
}
__typedArg0 := symtype
__typedArg1 := tmp2799
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp2829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2810 Obj

if True == tmp2829 {
tmp2827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput_7, tmp2827)
}
__typedArg0 := syminput_7
__typedArg1 := tmp2827
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2812 Obj

if True == tmp2828 {
tmp2825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2825)
}
__typedArg0 := tmp2825
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2814 Obj

if True == tmp2826 {
tmp2822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2822)
}
__typedArg0 := tmp2822
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2823)
}
__typedArg0 := tmp2823
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2816 Obj

if True == tmp2824 {
tmp2818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2818)
}
__typedArg0 := tmp2818
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2819)
}
__typedArg0 := tmp2819
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2820)
}
__typedArg0 := Nil
__typedArg1 := tmp2820
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2817 Obj

if True == tmp2821 {
ifres2817 = True


} else {
ifres2817 = False


}

ifres2816 = ifres2817


} else {
ifres2816 = False


}

var ifres2815 Obj

if True == ifres2816 {
ifres2815 = True


} else {
ifres2815 = False


}

ifres2814 = ifres2815


} else {
ifres2814 = False


}

var ifres2813 Obj

if True == ifres2814 {
ifres2813 = True


} else {
ifres2813 = False


}

ifres2812 = ifres2813


} else {
ifres2812 = False


}

var ifres2811 Obj

if True == ifres2812 {
ifres2811 = True


} else {
ifres2811 = False


}

ifres2810 = ifres2811


} else {
ifres2810 = False


}

if True == ifres2810 {
tmp2800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2800)
}
__typedArg0 := tmp2800
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2802 := Call(__e, PrimFunc(symshen_4rcons__form), tmp2801)


tmp2803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2803)
}
__typedArg0 := tmp2803
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2802, tmp2804)
}
__typedArg0 := tmp2802
__typedArg1 := tmp2804
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput_7, tmp2805)
}
__typedArg0 := syminput_7
__typedArg1 := tmp2805
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp2808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V547)
}
__typedArg0 := V547
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2808 {
tmp2806 := MakeNative(func(__e *ControlFlow) {
Z548 := __e.Get(1)
_ = Z548
__e.TailApply(PrimFunc(symshen_4shen_1_6kl_1h), Z548)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp2806, V547)
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

tmp2889 := Call(__e, ns2_1set, symshen_4shen_1_6kl_1h, tmp2787)


_ = tmp2889

tmp2890 := MakeNative(func(__e *ControlFlow) {
V549 := __e.Get(1)
_ = V549
V550 := __e.Get(2)
_ = V550
tmp2891 := MakeNative(func(__e *ControlFlow) {
Z551 := __e.Get(1)
_ = Z551
__e.TailApply(PrimFunc(symshen_4_5define_6), Z551)
return
}, 1)

tmp2892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V549, V550)
}
__typedArg0 := V549
__typedArg1 := V550
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symcompile), tmp2891, tmp2892)
return


}, 2)

tmp2893 := Call(__e, ns2_1set, symshen_4shendef_1_6kldef, tmp2890)


_ = tmp2893

tmp2894 := MakeNative(func(__e *ControlFlow) {
V552 := __e.Get(1)
_ = V552
tmp2895 := Call(__e, PrimFunc(symshen_4_5name_6), V552)


let__2587 := tmp2895
_ = let__2587

tmp2921 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2587)


var ifres2896 Obj

if True == tmp2921 {
tmp2897 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2896 = tmp2897


} else {
tmp2898 := Call(__e, PrimFunc(symshen_4_5_1out), let__2587)


let__2588 := tmp2898
_ = let__2588

tmp2899 := Call(__e, PrimFunc(symshen_4in_1_6), let__2587)


let__2589 := tmp2899
_ = let__2589

tmp2920 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2589, sym_i)


var ifres2900 Obj

if True == tmp2920 {
tmp2901 := Call(__e, PrimFunc(symtail), let__2589)


let__2590 := tmp2901
_ = let__2590

tmp2902 := Call(__e, PrimFunc(symshen_4_5signature_6), let__2590)


let__2591 := tmp2902
_ = let__2591

tmp2918 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2591)


var ifres2903 Obj

if True == tmp2918 {
tmp2904 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2903 = tmp2904


} else {
tmp2905 := Call(__e, PrimFunc(symshen_4in_1_6), let__2591)


let__2592 := tmp2905
_ = let__2592

tmp2917 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2592, sym_j)


var ifres2906 Obj

if True == tmp2917 {
tmp2907 := Call(__e, PrimFunc(symtail), let__2592)


let__2593 := tmp2907
_ = let__2593

tmp2908 := Call(__e, PrimFunc(symshen_4_5rules_6), let__2593)


let__2594 := tmp2908
_ = let__2594

tmp2915 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2594)


var ifres2909 Obj

if True == tmp2915 {
tmp2910 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2909 = tmp2910


} else {
tmp2911 := Call(__e, PrimFunc(symshen_4_5_1out), let__2594)


let__2595 := tmp2911
_ = let__2595

tmp2912 := Call(__e, PrimFunc(symshen_4in_1_6), let__2594)


let__2596 := tmp2912
_ = let__2596

tmp2913 := Call(__e, PrimFunc(symshen_4shendef_1_6kldef_1h), let__2588, let__2595)


tmp2914 := Call(__e, PrimFunc(symshen_4comb), let__2596, tmp2913)


ifres2909 = tmp2914


}

ifres2906 = ifres2909


} else {
tmp2916 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2906 = tmp2916


}

ifres2903 = ifres2906


}

ifres2900 = ifres2903


} else {
tmp2919 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2900 = tmp2919


}

ifres2896 = ifres2900


}

let__2586 := ifres2896
_ = let__2586

tmp2939 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2586)


if True == tmp2939 {
tmp2922 := Call(__e, PrimFunc(symshen_4_5name_6), V552)


let__2598 := tmp2922
_ = let__2598

tmp2935 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2598)


var ifres2923 Obj

if True == tmp2935 {
tmp2924 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2923 = tmp2924


} else {
tmp2925 := Call(__e, PrimFunc(symshen_4_5_1out), let__2598)


let__2599 := tmp2925
_ = let__2599

tmp2926 := Call(__e, PrimFunc(symshen_4in_1_6), let__2598)


let__2600 := tmp2926
_ = let__2600

tmp2927 := Call(__e, PrimFunc(symshen_4_5rules_6), let__2600)


let__2601 := tmp2927
_ = let__2601

tmp2934 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2601)


var ifres2928 Obj

if True == tmp2934 {
tmp2929 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2928 = tmp2929


} else {
tmp2930 := Call(__e, PrimFunc(symshen_4_5_1out), let__2601)


let__2602 := tmp2930
_ = let__2602

tmp2931 := Call(__e, PrimFunc(symshen_4in_1_6), let__2601)


let__2603 := tmp2931
_ = let__2603

tmp2932 := Call(__e, PrimFunc(symshen_4shendef_1_6kldef_1h), let__2599, let__2602)


tmp2933 := Call(__e, PrimFunc(symshen_4comb), let__2603, tmp2932)


ifres2928 = tmp2933


}

ifres2923 = ifres2928


}

let__2597 := ifres2923
_ = let__2597

tmp2937 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2597)


if True == tmp2937 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2597)
return
}


} else {
__e.Return(let__2586)
return
}


}, 1)

tmp2940 := Call(__e, ns2_1set, symshen_4_5define_6, tmp2894)


_ = tmp2940

tmp2941 := MakeNative(func(__e *ControlFlow) {
V571 := __e.Get(1)
_ = V571
V572 := __e.Get(2)
_ = V572
tmp2942 := MakeNative(func(__e *ControlFlow) {
Z574 := __e.Get(1)
_ = Z574
__e.TailApply(PrimFunc(symfst), Z574)
return
}, 1)

tmp2943 := Call(__e, PrimFunc(symmap), tmp2942, V572)


let__2604 := tmp2943
_ = let__2604

tmp2944 := Call(__e, PrimFunc(symshen_4arity_1chk), V571, let__2604)


let__2605 := tmp2944
_ = let__2605

tmp2945 := MakeNative(func(__e *ControlFlow) {
Z577 := __e.Get(1)
_ = Z577
__e.TailApply(PrimFunc(symshen_4free_1var_1chk), V571, Z577)
return
}, 1)

tmp2946 := Call(__e, PrimFunc(symmap), tmp2945, V572)


let__2606 := tmp2946
_ = let__2606

tmp2947 := Call(__e, PrimFunc(symshen_4unprotect), V572)


let__2607 := tmp2947
_ = let__2607

tmp2948 := Call(__e, PrimFunc(symshen_4compile_1to_1kl), V571, let__2607, let__2605)


tmp2949 := Call(__e, PrimFunc(symshen_4factorise_1code), tmp2948)


let__2608 := tmp2949
_ = let__2608

__e.Return(let__2608)
return


}, 2)

tmp2950 := Call(__e, ns2_1set, symshen_4shendef_1_6kldef_1h, tmp2941)


_ = tmp2950

tmp2951 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__selftop:
V580 := __self1
_ = V580
tmp2977 := Call(__e, PrimFunc(symtuple_2), V580)


if True == tmp2977 {
tmp2952 := Call(__e, PrimFunc(symfst), V580)


tmp2953 := Call(__e, PrimFunc(symshen_4unprotect), tmp2952)


tmp2954 := Call(__e, PrimFunc(symsnd), V580)


tmp2955 := Call(__e, PrimFunc(symshen_4unprotect), tmp2954)


__e.TailApply(PrimFunc(sym_8p), tmp2953, tmp2955)
return


} else {
tmp2975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V580)
}
__typedArg0 := V580
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2962 Obj

if True == tmp2975 {
tmp2973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V580)
}
__typedArg0 := V580
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symprotect, tmp2973)
}
__typedArg0 := symprotect
__typedArg1 := tmp2973
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2964 Obj

if True == tmp2974 {
tmp2971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V580)
}
__typedArg0 := V580
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2971)
}
__typedArg0 := tmp2971
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2966 Obj

if True == tmp2972 {
tmp2968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V580)
}
__typedArg0 := V580
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2968)
}
__typedArg0 := tmp2968
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2969)
}
__typedArg0 := Nil
__typedArg1 := tmp2969
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2967 Obj

if True == tmp2970 {
ifres2967 = True


} else {
ifres2967 = False


}

ifres2966 = ifres2967


} else {
ifres2966 = False


}

var ifres2965 Obj

if True == ifres2966 {
ifres2965 = True


} else {
ifres2965 = False


}

ifres2964 = ifres2965


} else {
ifres2964 = False


}

var ifres2963 Obj

if True == ifres2964 {
ifres2963 = True


} else {
ifres2963 = False


}

ifres2962 = ifres2963


} else {
ifres2962 = False


}

if True == ifres2962 {
tmp2956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V580)
}
__typedArg0 := V580
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2956)
}
__typedArg0 := tmp2956
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

if PrimFunc(symshen_4unprotect) == __self {
__self1 = tmp2957
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4unprotect), tmp2957)
return


} else {
tmp2960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V580)
}
__typedArg0 := V580
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2960 {
tmp2958 := MakeNative(func(__e *ControlFlow) {
Z581 := __e.Get(1)
_ = Z581
__e.TailApply(PrimFunc(symshen_4unprotect), Z581)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp2958, V580)
return


} else {
__e.Return(V580)
return
}


}


}


}, 1)

tmp2978 := Call(__e, ns2_1set, symshen_4unprotect, tmp2951)


_ = tmp2978

tmp2979 := MakeNative(func(__e *ControlFlow) {
V582 := __e.Get(1)
_ = V582
tmp2993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V582)
}
__typedArg0 := V582
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2980 Obj

if True == tmp2993 {
tmp2981 := Call(__e, PrimFunc(symhead), V582)


let__2610 := tmp2981
_ = let__2610

tmp2982 := Call(__e, PrimFunc(symtail), V582)


let__2611 := tmp2982
_ = let__2611

tmp2990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(let__2610)
}
__typedArg0 := let__2610
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

var ifres2986 Obj

if True == tmp2990 {
tmp2988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(let__2610)
}
__typedArg0 := let__2610
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

tmp2989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp2988)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp2988
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres2987 Obj

if True == tmp2989 {
ifres2987 = True


} else {
ifres2987 = False


}

ifres2986 = ifres2987


} else {
ifres2986 = False


}

var ifres2983 Obj

if True == ifres2986 {
ifres2983 = let__2610


} else {
tmp2984 := Call(__e, PrimFunc(symshen_4app), let__2610, MakeString(" is not a legitimate function name.\n"), symshen_4a)


tmp2985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp2984)
}
__typedArg0 := tmp2984
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres2983 = tmp2985


}

tmp2991 := Call(__e, PrimFunc(symshen_4comb), let__2611, ifres2983)


ifres2980 = tmp2991


} else {
tmp2992 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2980 = tmp2992


}

let__2609 := ifres2980
_ = let__2609

tmp2995 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2609)


if True == tmp2995 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2609)
return
}


}, 1)

tmp2996 := Call(__e, ns2_1set, symshen_4_5name_6, tmp2979)


_ = tmp2996

tmp2997 := MakeNative(func(__e *ControlFlow) {
V586 := __e.Get(1)
_ = V586
tmp3016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V586)
}
__typedArg0 := V586
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2998 Obj

if True == tmp3016 {
tmp2999 := Call(__e, PrimFunc(symhead), V586)


let__2613 := tmp2999
_ = let__2613

tmp3000 := Call(__e, PrimFunc(symtail), V586)


let__2614 := tmp3000
_ = let__2614

tmp3001 := Call(__e, PrimFunc(symshen_4_5signature_6), let__2614)


let__2615 := tmp3001
_ = let__2615

tmp3014 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2615)


var ifres3002 Obj

if True == tmp3014 {
tmp3003 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3002 = tmp3003


} else {
tmp3004 := Call(__e, PrimFunc(symshen_4_5_1out), let__2615)


let__2616 := tmp3004
_ = let__2616

tmp3005 := Call(__e, PrimFunc(symshen_4in_1_6), let__2615)


let__2617 := tmp3005
_ = let__2617

tmp3010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, Nil)
}
__typedArg0 := sym_j
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, tmp3010)
}
__typedArg0 := sym_i
__typedArg1 := tmp3010
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3012 := Call(__e, PrimFunc(symelement_2), let__2613, tmp3011)


tmp3013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp3012)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp3012
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres3006 Obj

if True == tmp3013 {
tmp3007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2613, let__2616)
}
__typedArg0 := let__2613
__typedArg1 := let__2616
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3008 := Call(__e, PrimFunc(symshen_4comb), let__2617, tmp3007)


ifres3006 = tmp3008


} else {
tmp3009 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3006 = tmp3009


}

ifres3002 = ifres3006


}

ifres2998 = ifres3002


} else {
tmp3015 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2998 = tmp3015


}

let__2612 := ifres2998
_ = let__2612

tmp3026 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2612)


if True == tmp3026 {
tmp3017 := Call(__e, PrimFunc(sym_5e_6), V586)


let__2619 := tmp3017
_ = let__2619

tmp3022 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2619)


var ifres3018 Obj

if True == tmp3022 {
tmp3019 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3018 = tmp3019


} else {
tmp3020 := Call(__e, PrimFunc(symshen_4in_1_6), let__2619)


let__2620 := tmp3020
_ = let__2620

tmp3021 := Call(__e, PrimFunc(symshen_4comb), let__2620, Nil)


ifres3018 = tmp3021


}

let__2618 := ifres3018
_ = let__2618

tmp3024 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2618)


if True == tmp3024 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2618)
return
}


} else {
__e.Return(let__2612)
return
}


}, 1)

tmp3027 := Call(__e, ns2_1set, symshen_4_5signature_6, tmp2997)


_ = tmp3027

tmp3028 := MakeNative(func(__e *ControlFlow) {
V596 := __e.Get(1)
_ = V596
tmp3029 := Call(__e, PrimFunc(symshen_4_5rule_6), V596)


let__2622 := tmp3029
_ = let__2622

tmp3043 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2622)


var ifres3030 Obj

if True == tmp3043 {
tmp3031 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3030 = tmp3031


} else {
tmp3032 := Call(__e, PrimFunc(symshen_4_5_1out), let__2622)


let__2623 := tmp3032
_ = let__2623

tmp3033 := Call(__e, PrimFunc(symshen_4in_1_6), let__2622)


let__2624 := tmp3033
_ = let__2624

tmp3034 := Call(__e, PrimFunc(symshen_4_5rules_6), let__2624)


let__2625 := tmp3034
_ = let__2625

tmp3042 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2625)


var ifres3035 Obj

if True == tmp3042 {
tmp3036 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3035 = tmp3036


} else {
tmp3037 := Call(__e, PrimFunc(symshen_4_5_1out), let__2625)


let__2626 := tmp3037
_ = let__2626

tmp3038 := Call(__e, PrimFunc(symshen_4in_1_6), let__2625)


let__2627 := tmp3038
_ = let__2627

tmp3039 := Call(__e, PrimFunc(symshen_4linearise), let__2623)


tmp3040 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3039, let__2626)
}
__typedArg0 := tmp3039
__typedArg1 := let__2626
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3041 := Call(__e, PrimFunc(symshen_4comb), let__2627, tmp3040)


ifres3035 = tmp3041


}

ifres3030 = ifres3035


}

let__2621 := ifres3030
_ = let__2621

tmp3059 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2621)


if True == tmp3059 {
tmp3044 := Call(__e, PrimFunc(sym_5_b_6), V596)


let__2629 := tmp3044
_ = let__2629

tmp3055 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2629)


var ifres3045 Obj

if True == tmp3055 {
tmp3046 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3045 = tmp3046


} else {
tmp3047 := Call(__e, PrimFunc(symshen_4_5_1out), let__2629)


let__2630 := tmp3047
_ = let__2630

tmp3048 := Call(__e, PrimFunc(symshen_4in_1_6), let__2629)


let__2631 := tmp3048
_ = let__2631

tmp3053 := Call(__e, PrimFunc(symempty_2), let__2630)


var ifres3049 Obj

if True == tmp3053 {
ifres3049 = Nil


} else {
tmp3050 := Call(__e, PrimFunc(symshen_4app), let__2630, MakeString("\n ..."), symshen_4r)


tmp3052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Shen syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp3050)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Shen syntax error here:\n ")
__typedArg1 := tmp3050
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Shen syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp3050)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Shen syntax error here:\n ")
__typedArg1 := tmp3050
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres3049 = tmp3052


}

tmp3054 := Call(__e, PrimFunc(symshen_4comb), let__2631, ifres3049)


ifres3045 = tmp3054


}

let__2628 := ifres3045
_ = let__2628

tmp3057 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2628)


if True == tmp3057 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2628)
return
}


} else {
__e.Return(let__2621)
return
}


}, 1)

tmp3060 := Call(__e, ns2_1set, symshen_4_5rules_6, tmp3028)


_ = tmp3060

tmp3061 := MakeNative(func(__e *ControlFlow) {
V610 := __e.Get(1)
_ = V610
tmp3066 := Call(__e, PrimFunc(symtuple_2), V610)


if True == tmp3066 {
tmp3062 := Call(__e, PrimFunc(symfst), V610)


tmp3063 := Call(__e, PrimFunc(symfst), V610)


tmp3064 := Call(__e, PrimFunc(symsnd), V610)


__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp3062, tmp3063, Nil, tmp3064)
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

tmp3067 := Call(__e, ns2_1set, symshen_4linearise, tmp3061)


_ = tmp3067

tmp3068 := MakeNative(func(__e *ControlFlow) {
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
tmp3105 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V623)
}
__typedArg0 := Nil
__typedArg1 := V623
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp3105 {
__e.TailApply(PrimFunc(sym_8p), V624, V626)
return
} else {
tmp3103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3099 Obj

if True == tmp3103 {
tmp3101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3101)
}
__typedArg0 := tmp3101
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3100 Obj

if True == tmp3102 {
ifres3100 = True


} else {
ifres3100 = False


}

ifres3099 = ifres3100


} else {
ifres3099 = False


}

if True == ifres3099 {
tmp3069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3071 := Call(__e, PrimFunc(symappend), tmp3069, tmp3070)


if PrimFunc(symshen_4linearise_1h) == __self {
__self1, __self2, __self3, __self4 = tmp3071, V624, V625, V626
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp3071, V624, V625, V626)
return


} else {
tmp3097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3093 Obj

if True == tmp3097 {
tmp3095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp3095)
}
__typedArg0 := tmp3095
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres3094 Obj

if True == tmp3096 {
ifres3094 = True


} else {
ifres3094 = False


}

ifres3093 = ifres3094


} else {
ifres3093 = False


}

if True == ifres3093 {
tmp3087 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3088 := Call(__e, PrimFunc(symelement_2), tmp3087, V625)


if True == tmp3088 {
tmp3072 := Call(__e, PrimFunc(symgensym), symV)


let__2632 := tmp3072
_ = let__2632

tmp3073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3074 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3075 := Call(__e, PrimFunc(symshen_4rep_1X), tmp3074, let__2632, V624)


tmp3076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3076, Nil)
}
__typedArg0 := tmp3076
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2632, tmp3077)
}
__typedArg0 := let__2632
__typedArg1 := tmp3077
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp3078)
}
__typedArg0 := sym_a
__typedArg1 := tmp3078
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V626, Nil)
}
__typedArg0 := V626
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3079, tmp3080)
}
__typedArg0 := tmp3079
__typedArg1 := tmp3080
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp3081)
}
__typedArg0 := symwhere
__typedArg1 := tmp3081
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4linearise_1h) == __self {
__self1, __self2, __self3, __self4 = tmp3073, tmp3075, V625, tmp3082
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp3073, tmp3075, V625, tmp3082)
return


} else {
tmp3083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3084 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3085 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3084, V625)
}
__typedArg0 := tmp3084
__typedArg1 := V625
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4linearise_1h) == __self {
__self1, __self2, __self3, __self4 = tmp3083, V624, tmp3085, V626
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp3083, V624, tmp3085, V626)
return


}


} else {
tmp3091 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3091 {
tmp3089 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V623)
}
__typedArg0 := V623
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4linearise_1h) == __self {
__self1, __self2, __self3, __self4 = tmp3089, V624, V625, V626
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp3089, V624, V625, V626)
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

tmp3106 := Call(__e, ns2_1set, symshen_4linearise_1h, tmp3068)


_ = tmp3106

tmp3107 := MakeNative(func(__e *ControlFlow) {
V628 := __e.Get(1)
_ = V628
tmp3108 := Call(__e, PrimFunc(symshen_4_5patterns_6), V628)


let__2634 := tmp3108
_ = let__2634

tmp3136 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2634)


var ifres3109 Obj

if True == tmp3136 {
tmp3110 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3109 = tmp3110


} else {
tmp3111 := Call(__e, PrimFunc(symshen_4_5_1out), let__2634)


let__2635 := tmp3111
_ = let__2635

tmp3112 := Call(__e, PrimFunc(symshen_4in_1_6), let__2634)


let__2636 := tmp3112
_ = let__2636

tmp3135 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2636, sym_1_6)


var ifres3113 Obj

if True == tmp3135 {
tmp3114 := Call(__e, PrimFunc(symtail), let__2636)


let__2637 := tmp3114
_ = let__2637

tmp3133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__2637)
}
__typedArg0 := let__2637
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3115 Obj

if True == tmp3133 {
tmp3116 := Call(__e, PrimFunc(symhead), let__2637)


let__2638 := tmp3116
_ = let__2638

tmp3117 := Call(__e, PrimFunc(symtail), let__2637)


let__2639 := tmp3117
_ = let__2639

tmp3131 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2639, symwhere)


var ifres3118 Obj

if True == tmp3131 {
tmp3119 := Call(__e, PrimFunc(symtail), let__2639)


let__2640 := tmp3119
_ = let__2640

tmp3129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__2640)
}
__typedArg0 := let__2640
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3120 Obj

if True == tmp3129 {
tmp3121 := Call(__e, PrimFunc(symhead), let__2640)


let__2641 := tmp3121
_ = let__2641

tmp3122 := Call(__e, PrimFunc(symtail), let__2640)


let__2642 := tmp3122
_ = let__2642

tmp3123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2638, Nil)
}
__typedArg0 := let__2638
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2641, tmp3123)
}
__typedArg0 := let__2641
__typedArg1 := tmp3123
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp3124)
}
__typedArg0 := symwhere
__typedArg1 := tmp3124
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3126 := Call(__e, PrimFunc(sym_8p), let__2635, tmp3125)


tmp3127 := Call(__e, PrimFunc(symshen_4comb), let__2642, tmp3126)


ifres3120 = tmp3127


} else {
tmp3128 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3120 = tmp3128


}

ifres3118 = ifres3120


} else {
tmp3130 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3118 = tmp3130


}

ifres3115 = ifres3118


} else {
tmp3132 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3115 = tmp3132


}

ifres3113 = ifres3115


} else {
tmp3134 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3113 = tmp3134


}

ifres3109 = ifres3113


}

let__2633 := ifres3109
_ = let__2633

tmp3211 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2633)


if True == tmp3211 {
tmp3137 := Call(__e, PrimFunc(symshen_4_5patterns_6), V628)


let__2644 := tmp3137
_ = let__2644

tmp3153 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2644)


var ifres3138 Obj

if True == tmp3153 {
tmp3139 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3138 = tmp3139


} else {
tmp3140 := Call(__e, PrimFunc(symshen_4_5_1out), let__2644)


let__2645 := tmp3140
_ = let__2645

tmp3141 := Call(__e, PrimFunc(symshen_4in_1_6), let__2644)


let__2646 := tmp3141
_ = let__2646

tmp3152 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2646, sym_1_6)


var ifres3142 Obj

if True == tmp3152 {
tmp3143 := Call(__e, PrimFunc(symtail), let__2646)


let__2647 := tmp3143
_ = let__2647

tmp3150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__2647)
}
__typedArg0 := let__2647
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3144 Obj

if True == tmp3150 {
tmp3145 := Call(__e, PrimFunc(symhead), let__2647)


let__2648 := tmp3145
_ = let__2648

tmp3146 := Call(__e, PrimFunc(symtail), let__2647)


let__2649 := tmp3146
_ = let__2649

tmp3147 := Call(__e, PrimFunc(sym_8p), let__2645, let__2648)


tmp3148 := Call(__e, PrimFunc(symshen_4comb), let__2649, tmp3147)


ifres3144 = tmp3148


} else {
tmp3149 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3144 = tmp3149


}

ifres3142 = ifres3144


} else {
tmp3151 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3142 = tmp3151


}

ifres3138 = ifres3142


}

let__2643 := ifres3138
_ = let__2643

tmp3209 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2643)


if True == tmp3209 {
tmp3154 := Call(__e, PrimFunc(symshen_4_5patterns_6), V628)


let__2651 := tmp3154
_ = let__2651

tmp3184 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2651)


var ifres3155 Obj

if True == tmp3184 {
tmp3156 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3155 = tmp3156


} else {
tmp3157 := Call(__e, PrimFunc(symshen_4_5_1out), let__2651)


let__2652 := tmp3157
_ = let__2652

tmp3158 := Call(__e, PrimFunc(symshen_4in_1_6), let__2651)


let__2653 := tmp3158
_ = let__2653

tmp3183 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2653, sym_5_1)


var ifres3159 Obj

if True == tmp3183 {
tmp3160 := Call(__e, PrimFunc(symtail), let__2653)


let__2654 := tmp3160
_ = let__2654

tmp3181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__2654)
}
__typedArg0 := let__2654
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3161 Obj

if True == tmp3181 {
tmp3162 := Call(__e, PrimFunc(symhead), let__2654)


let__2655 := tmp3162
_ = let__2655

tmp3163 := Call(__e, PrimFunc(symtail), let__2654)


let__2656 := tmp3163
_ = let__2656

tmp3179 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2656, symwhere)


var ifres3164 Obj

if True == tmp3179 {
tmp3165 := Call(__e, PrimFunc(symtail), let__2656)


let__2657 := tmp3165
_ = let__2657

tmp3177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__2657)
}
__typedArg0 := let__2657
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3166 Obj

if True == tmp3177 {
tmp3167 := Call(__e, PrimFunc(symhead), let__2657)


let__2658 := tmp3167
_ = let__2658

tmp3168 := Call(__e, PrimFunc(symtail), let__2657)


let__2659 := tmp3168
_ = let__2659

tmp3169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2655, Nil)
}
__typedArg0 := let__2655
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4choicepoint_b, tmp3169)
}
__typedArg0 := symshen_4choicepoint_b
__typedArg1 := tmp3169
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3170, Nil)
}
__typedArg0 := tmp3170
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2658, tmp3171)
}
__typedArg0 := let__2658
__typedArg1 := tmp3171
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp3172)
}
__typedArg0 := symwhere
__typedArg1 := tmp3172
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3174 := Call(__e, PrimFunc(sym_8p), let__2652, tmp3173)


tmp3175 := Call(__e, PrimFunc(symshen_4comb), let__2659, tmp3174)


ifres3166 = tmp3175


} else {
tmp3176 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3166 = tmp3176


}

ifres3164 = ifres3166


} else {
tmp3178 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3164 = tmp3178


}

ifres3161 = ifres3164


} else {
tmp3180 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3161 = tmp3180


}

ifres3159 = ifres3161


} else {
tmp3182 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3159 = tmp3182


}

ifres3155 = ifres3159


}

let__2650 := ifres3155
_ = let__2650

tmp3207 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2650)


if True == tmp3207 {
tmp3185 := Call(__e, PrimFunc(symshen_4_5patterns_6), V628)


let__2661 := tmp3185
_ = let__2661

tmp3203 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2661)


var ifres3186 Obj

if True == tmp3203 {
tmp3187 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3186 = tmp3187


} else {
tmp3188 := Call(__e, PrimFunc(symshen_4_5_1out), let__2661)


let__2662 := tmp3188
_ = let__2662

tmp3189 := Call(__e, PrimFunc(symshen_4in_1_6), let__2661)


let__2663 := tmp3189
_ = let__2663

tmp3202 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2663, sym_5_1)


var ifres3190 Obj

if True == tmp3202 {
tmp3191 := Call(__e, PrimFunc(symtail), let__2663)


let__2664 := tmp3191
_ = let__2664

tmp3200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__2664)
}
__typedArg0 := let__2664
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3192 Obj

if True == tmp3200 {
tmp3193 := Call(__e, PrimFunc(symhead), let__2664)


let__2665 := tmp3193
_ = let__2665

tmp3194 := Call(__e, PrimFunc(symtail), let__2664)


let__2666 := tmp3194
_ = let__2666

tmp3195 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2665, Nil)
}
__typedArg0 := let__2665
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4choicepoint_b, tmp3195)
}
__typedArg0 := symshen_4choicepoint_b
__typedArg1 := tmp3195
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3197 := Call(__e, PrimFunc(sym_8p), let__2662, tmp3196)


tmp3198 := Call(__e, PrimFunc(symshen_4comb), let__2666, tmp3197)


ifres3192 = tmp3198


} else {
tmp3199 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3192 = tmp3199


}

ifres3190 = ifres3192


} else {
tmp3201 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3190 = tmp3201


}

ifres3186 = ifres3190


}

let__2660 := ifres3186
_ = let__2660

tmp3205 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2660)


if True == tmp3205 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2660)
return
}


} else {
__e.Return(let__2650)
return
}


} else {
__e.Return(let__2643)
return
}


} else {
__e.Return(let__2633)
return
}


}, 1)

tmp3212 := Call(__e, ns2_1set, symshen_4_5rule_6, tmp3107)


_ = tmp3212

tmp3213 := MakeNative(func(__e *ControlFlow) {
V663 := __e.Get(1)
_ = V663
tmp3214 := Call(__e, PrimFunc(symshen_4_5pattern_6), V663)


let__2668 := tmp3214
_ = let__2668

tmp3227 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2668)


var ifres3215 Obj

if True == tmp3227 {
tmp3216 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3215 = tmp3216


} else {
tmp3217 := Call(__e, PrimFunc(symshen_4_5_1out), let__2668)


let__2669 := tmp3217
_ = let__2669

tmp3218 := Call(__e, PrimFunc(symshen_4in_1_6), let__2668)


let__2670 := tmp3218
_ = let__2670

tmp3219 := Call(__e, PrimFunc(symshen_4_5patterns_6), let__2670)


let__2671 := tmp3219
_ = let__2671

tmp3226 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2671)


var ifres3220 Obj

if True == tmp3226 {
tmp3221 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3220 = tmp3221


} else {
tmp3222 := Call(__e, PrimFunc(symshen_4_5_1out), let__2671)


let__2672 := tmp3222
_ = let__2672

tmp3223 := Call(__e, PrimFunc(symshen_4in_1_6), let__2671)


let__2673 := tmp3223
_ = let__2673

tmp3224 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2669, let__2672)
}
__typedArg0 := let__2669
__typedArg1 := let__2672
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3225 := Call(__e, PrimFunc(symshen_4comb), let__2673, tmp3224)


ifres3220 = tmp3225


}

ifres3215 = ifres3220


}

let__2667 := ifres3215
_ = let__2667

tmp3237 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2667)


if True == tmp3237 {
tmp3228 := Call(__e, PrimFunc(sym_5e_6), V663)


let__2675 := tmp3228
_ = let__2675

tmp3233 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2675)


var ifres3229 Obj

if True == tmp3233 {
tmp3230 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3229 = tmp3230


} else {
tmp3231 := Call(__e, PrimFunc(symshen_4in_1_6), let__2675)


let__2676 := tmp3231
_ = let__2676

tmp3232 := Call(__e, PrimFunc(symshen_4comb), let__2676, Nil)


ifres3229 = tmp3232


}

let__2674 := ifres3229
_ = let__2674

tmp3235 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2674)


if True == tmp3235 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2674)
return
}


} else {
__e.Return(let__2667)
return
}


}, 1)

tmp3238 := Call(__e, ns2_1set, symshen_4_5patterns_6, tmp3213)


_ = tmp3238

tmp3239 := MakeNative(func(__e *ControlFlow) {
V674 := __e.Get(1)
_ = V674
tmp3260 := Call(__e, PrimFunc(symshen_4ccons_2), V674)


var ifres3240 Obj

if True == tmp3260 {
tmp3241 := Call(__e, PrimFunc(symhead), V674)


let__2678 := tmp3241
_ = let__2678

tmp3242 := Call(__e, PrimFunc(symtail), V674)


let__2679 := tmp3242
_ = let__2679

tmp3258 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2678, symvector)


var ifres3243 Obj

if True == tmp3258 {
tmp3244 := Call(__e, PrimFunc(symtail), let__2678)


let__2680 := tmp3244
_ = let__2680

tmp3256 := Call(__e, PrimFunc(symshen_4hds_a_2), let__2680, MakeInteger(0))


var ifres3245 Obj

if True == tmp3256 {
tmp3246 := Call(__e, PrimFunc(symtail), let__2680)


let__2681 := tmp3246
_ = let__2681

tmp3247 := Call(__e, PrimFunc(sym_5end_6), let__2681)


let__2682 := tmp3247
_ = let__2682

tmp3254 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2682)


var ifres3248 Obj

if True == tmp3254 {
tmp3249 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3248 = tmp3249


} else {
tmp3250 := Call(__e, PrimFunc(symshen_4in_1_6), let__2682)


let__2683 := tmp3250
_ = let__2683

tmp3251 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeInteger(0), Nil)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp3251)
}
__typedArg0 := symvector
__typedArg1 := tmp3251
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3253 := Call(__e, PrimFunc(symshen_4comb), let__2679, tmp3252)


ifres3248 = tmp3253


}

ifres3245 = ifres3248


} else {
tmp3255 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3245 = tmp3255


}

ifres3243 = ifres3245


} else {
tmp3257 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3243 = tmp3257


}

ifres3240 = ifres3243


} else {
tmp3259 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3240 = tmp3259


}

let__2677 := ifres3240
_ = let__2677

tmp3283 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2677)


if True == tmp3283 {
tmp3270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V674)
}
__typedArg0 := V674
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3261 Obj

if True == tmp3270 {
tmp3262 := Call(__e, PrimFunc(symhead), V674)


let__2685 := tmp3262
_ = let__2685

tmp3263 := Call(__e, PrimFunc(symtail), V674)


let__2686 := tmp3263
_ = let__2686

tmp3268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(let__2685)
}
__typedArg0 := let__2685
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3264 Obj

if True == tmp3268 {
tmp3265 := Call(__e, PrimFunc(symshen_4compound_1pattern), let__2685)


tmp3266 := Call(__e, PrimFunc(symshen_4comb), let__2686, tmp3265)


ifres3264 = tmp3266


} else {
tmp3267 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3264 = tmp3267


}

ifres3261 = ifres3264


} else {
tmp3269 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3261 = tmp3269


}

let__2684 := ifres3261
_ = let__2684

tmp3281 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2684)


if True == tmp3281 {
tmp3271 := Call(__e, PrimFunc(symshen_4_5simple_1pattern_6), V674)


let__2688 := tmp3271
_ = let__2688

tmp3277 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2688)


var ifres3272 Obj

if True == tmp3277 {
tmp3273 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3272 = tmp3273


} else {
tmp3274 := Call(__e, PrimFunc(symshen_4_5_1out), let__2688)


let__2689 := tmp3274
_ = let__2689

tmp3275 := Call(__e, PrimFunc(symshen_4in_1_6), let__2688)


let__2690 := tmp3275
_ = let__2690

tmp3276 := Call(__e, PrimFunc(symshen_4comb), let__2690, let__2689)


ifres3272 = tmp3276


}

let__2687 := ifres3272
_ = let__2687

tmp3279 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2687)


if True == tmp3279 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2687)
return
}


} else {
__e.Return(let__2684)
return
}


} else {
__e.Return(let__2677)
return
}


}, 1)

tmp3284 := Call(__e, ns2_1set, symshen_4_5pattern_6, tmp3239)


_ = tmp3284

tmp3285 := MakeNative(func(__e *ControlFlow) {
V689 := __e.Get(1)
_ = V689
tmp3294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V689)
}
__typedArg0 := V689
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3286 Obj

if True == tmp3294 {
tmp3287 := Call(__e, PrimFunc(symhead), V689)


let__2692 := tmp3287
_ = let__2692

tmp3288 := Call(__e, PrimFunc(symtail), V689)


let__2693 := tmp3288
_ = let__2693

tmp3292 := Call(__e, PrimFunc(symshen_4constructor_2), let__2692)


var ifres3289 Obj

if True == tmp3292 {
tmp3290 := Call(__e, PrimFunc(symshen_4comb), let__2693, let__2692)


ifres3289 = tmp3290


} else {
tmp3291 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3289 = tmp3291


}

ifres3286 = ifres3289


} else {
tmp3293 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3286 = tmp3293


}

let__2691 := ifres3286
_ = let__2691

tmp3296 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2691)


if True == tmp3296 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2691)
return
}


}, 1)

tmp3297 := Call(__e, ns2_1set, symshen_4_5constructor_6, tmp3285)


_ = tmp3297

tmp3298 := MakeNative(func(__e *ControlFlow) {
V693 := __e.Get(1)
_ = V693
tmp3299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, Nil)
}
__typedArg0 := sym_8v
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp3299)
}
__typedArg0 := sym_8s
__typedArg1 := tmp3299
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp3300)
}
__typedArg0 := sym_8p
__typedArg1 := tmp3300
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp3301)
}
__typedArg0 := symcons
__typedArg1 := tmp3301
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symelement_2), V693, tmp3302)
return


}, 1)

tmp3303 := Call(__e, ns2_1set, symshen_4constructor_2, tmp3298)


_ = tmp3303

tmp3304 := MakeNative(func(__e *ControlFlow) {
V694 := __e.Get(1)
_ = V694
tmp3305 := Call(__e, PrimFunc(symshen_4app), V694, MakeString(" is not a legitimate constructor\n"), symshen_4r)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp3305)
}
__typedArg0 := tmp3305
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

tmp3306 := Call(__e, ns2_1set, symshen_4constructor_1error, tmp3304)


_ = tmp3306

tmp3307 := MakeNative(func(__e *ControlFlow) {
V695 := __e.Get(1)
_ = V695
V696 := __e.Get(2)
_ = V696
tmp3308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dcustom_1pattern_1compiler_d)
}
__typedArg0 := symshen_4_dcustom_1pattern_1compiler_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__2694 := tmp3308
_ = let__2694

tmp3311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__2694, False)
}
__typedArg0 := let__2694
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp3311 {
__e.TailApply(PrimFunc(symthaw), V696)
return
} else {
tmp3309 := Call(__e, let__2694, V695)


__e.TailApply(tmp3309, V696)
return


}


}, 2)

tmp3312 := Call(__e, ns2_1set, symshen_4custom_1pattern_1compiler, tmp3307)


_ = tmp3312

tmp3313 := MakeNative(func(__e *ControlFlow) {
V698 := __e.Get(1)
_ = V698
tmp3314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dcustom_1pattern_1reducer_d)
}
__typedArg0 := symshen_4_dcustom_1pattern_1reducer_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

let__2695 := tmp3314
_ = let__2695

tmp3316 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__2695, False)
}
__typedArg0 := let__2695
__typedArg1 := False
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp3316 {
__e.TailApply(PrimFunc(symfail))
return
} else {
__e.TailApply(let__2695, V698)
return
}


}, 1)

tmp3317 := Call(__e, ns2_1set, symshen_4custom_1pattern_1reducer, tmp3313)


_ = tmp3317

tmp3318 := MakeNative(func(__e *ControlFlow) {
V700 := __e.Get(1)
_ = V700
tmp3319 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4compound_1pattern_1h), V700)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4custom_1pattern_1compiler), V700, tmp3319)
return


}, 1)

tmp3320 := Call(__e, ns2_1set, symshen_4compound_1pattern, tmp3318)


_ = tmp3320

tmp3321 := MakeNative(func(__e *ControlFlow) {
V701 := __e.Get(1)
_ = V701
tmp3352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V701)
}
__typedArg0 := V701
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3333 Obj

if True == tmp3352 {
tmp3350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V701)
}
__typedArg0 := V701
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3350)
}
__typedArg0 := tmp3350
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3335 Obj

if True == tmp3351 {
tmp3347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V701)
}
__typedArg0 := V701
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3347)
}
__typedArg0 := tmp3347
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3348)
}
__typedArg0 := tmp3348
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3337 Obj

if True == tmp3349 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3344)
}
__typedArg0 := tmp3344
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3345)
}
__typedArg0 := Nil
__typedArg1 := tmp3345
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3339 Obj

if True == tmp3346 {
tmp3341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V701)
}
__typedArg0 := V701
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3342 := Call(__e, PrimFunc(symshen_4constructor_2), tmp3341)


var ifres3340 Obj

if True == tmp3342 {
ifres3340 = True


} else {
ifres3340 = False


}

ifres3339 = ifres3340


} else {
ifres3339 = False


}

var ifres3338 Obj

if True == ifres3339 {
ifres3338 = True


} else {
ifres3338 = False


}

ifres3337 = ifres3338


} else {
ifres3337 = False


}

var ifres3336 Obj

if True == ifres3337 {
ifres3336 = True


} else {
ifres3336 = False


}

ifres3335 = ifres3336


} else {
ifres3335 = False


}

var ifres3334 Obj

if True == ifres3335 {
ifres3334 = True


} else {
ifres3334 = False


}

ifres3333 = ifres3334


} else {
ifres3333 = False


}

if True == ifres3333 {
tmp3322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V701)
}
__typedArg0 := V701
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V701)
}
__typedArg0 := V701
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3323)
}
__typedArg0 := tmp3323
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3325 := Call(__e, PrimFunc(symshen_4compile_1pattern_1fragment), tmp3324)


tmp3326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V701)
}
__typedArg0 := V701
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3326)
}
__typedArg0 := tmp3326
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3327)
}
__typedArg0 := tmp3327
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3329 := Call(__e, PrimFunc(symshen_4compile_1pattern_1fragment), tmp3328)


tmp3330 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3329, Nil)
}
__typedArg0 := tmp3329
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3331 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3325, tmp3330)
}
__typedArg0 := tmp3325
__typedArg1 := tmp3330
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3322, tmp3331)
}
__typedArg0 := tmp3322
__typedArg1 := tmp3331
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4constructor_1error), V701)
return
}


}, 1)

tmp3353 := Call(__e, ns2_1set, symshen_4compound_1pattern_1h, tmp3321)


_ = tmp3353

tmp3354 := MakeNative(func(__e *ControlFlow) {
V702 := __e.Get(1)
_ = V702
tmp3383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V702)
}
__typedArg0 := V702
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3365 Obj

if True == tmp3383 {
tmp3381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V702)
}
__typedArg0 := V702
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symvector, tmp3381)
}
__typedArg0 := symvector
__typedArg1 := tmp3381
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3367 Obj

if True == tmp3382 {
tmp3379 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V702)
}
__typedArg0 := V702
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3379)
}
__typedArg0 := tmp3379
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3369 Obj

if True == tmp3380 {
tmp3376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V702)
}
__typedArg0 := V702
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3376)
}
__typedArg0 := tmp3376
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), tmp3377)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := tmp3377
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3371 Obj

if True == tmp3378 {
tmp3373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V702)
}
__typedArg0 := V702
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3373)
}
__typedArg0 := tmp3373
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3374)
}
__typedArg0 := Nil
__typedArg1 := tmp3374
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3372 Obj

if True == tmp3375 {
ifres3372 = True


} else {
ifres3372 = False


}

ifres3371 = ifres3372


} else {
ifres3371 = False


}

var ifres3370 Obj

if True == ifres3371 {
ifres3370 = True


} else {
ifres3370 = False


}

ifres3369 = ifres3370


} else {
ifres3369 = False


}

var ifres3368 Obj

if True == ifres3369 {
ifres3368 = True


} else {
ifres3368 = False


}

ifres3367 = ifres3368


} else {
ifres3367 = False


}

var ifres3366 Obj

if True == ifres3367 {
ifres3366 = True


} else {
ifres3366 = False


}

ifres3365 = ifres3366


} else {
ifres3365 = False


}

if True == ifres3365 {
__e.Return(V702)
return
} else {
tmp3363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V702)
}
__typedArg0 := V702
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3363 {
__e.TailApply(PrimFunc(symshen_4compound_1pattern), V702)
return
} else {
tmp3361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V702, sym__)
}
__typedArg0 := V702
__typedArg1 := sym__
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp3361 {
__e.TailApply(PrimFunc(symgensym), symY)
return
} else {
tmp3356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1, Nil)
}
__typedArg0 := sym_5_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp3356)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp3356
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3358 := Call(__e, PrimFunc(symelement_2), V702, tmp3357)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp3358)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp3358
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

tmp3384 := Call(__e, ns2_1set, symshen_4compile_1pattern_1fragment, tmp3354)


_ = tmp3384

tmp3385 := MakeNative(func(__e *ControlFlow) {
V707 := __e.Get(1)
_ = V707
tmp3411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V707)
}
__typedArg0 := V707
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3387 Obj

if True == tmp3411 {
tmp3409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V707)
}
__typedArg0 := V707
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8p, tmp3409)
}
__typedArg0 := sym_8p
__typedArg1 := tmp3409
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3389 Obj

if True == tmp3410 {
tmp3407 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V707)
}
__typedArg0 := V707
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3408 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3407)
}
__typedArg0 := tmp3407
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3391 Obj

if True == tmp3408 {
tmp3404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V707)
}
__typedArg0 := V707
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3405 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3404)
}
__typedArg0 := tmp3404
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4custom_1pattern, tmp3405)
}
__typedArg0 := symshen_4custom_1pattern
__typedArg1 := tmp3405
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3393 Obj

if True == tmp3406 {
tmp3401 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V707)
}
__typedArg0 := V707
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3401)
}
__typedArg0 := tmp3401
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3402)
}
__typedArg0 := tmp3402
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3395 Obj

if True == tmp3403 {
tmp3397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V707)
}
__typedArg0 := V707
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3397)
}
__typedArg0 := tmp3397
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3398)
}
__typedArg0 := tmp3398
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3399)
}
__typedArg0 := Nil
__typedArg1 := tmp3399
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3396 Obj

if True == tmp3400 {
ifres3396 = True


} else {
ifres3396 = False


}

ifres3395 = ifres3396


} else {
ifres3395 = False


}

var ifres3394 Obj

if True == ifres3395 {
ifres3394 = True


} else {
ifres3394 = False


}

ifres3393 = ifres3394


} else {
ifres3393 = False


}

var ifres3392 Obj

if True == ifres3393 {
ifres3392 = True


} else {
ifres3392 = False


}

ifres3391 = ifres3392


} else {
ifres3391 = False


}

var ifres3390 Obj

if True == ifres3391 {
ifres3390 = True


} else {
ifres3390 = False


}

ifres3389 = ifres3390


} else {
ifres3389 = False


}

var ifres3388 Obj

if True == ifres3389 {
ifres3388 = True


} else {
ifres3388 = False


}

ifres3387 = ifres3388


} else {
ifres3387 = False


}

if True == ifres3387 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp3412 := Call(__e, ns2_1set, symshen_4custom_1pattern_2, tmp3385)


_ = tmp3412

tmp3413 := MakeNative(func(__e *ControlFlow) {
V710 := __e.Get(1)
_ = V710
tmp3441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V710)
}
__typedArg0 := V710
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3417 Obj

if True == tmp3441 {
tmp3439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V710)
}
__typedArg0 := V710
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8p, tmp3439)
}
__typedArg0 := sym_8p
__typedArg1 := tmp3439
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3419 Obj

if True == tmp3440 {
tmp3437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V710)
}
__typedArg0 := V710
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3437)
}
__typedArg0 := tmp3437
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3421 Obj

if True == tmp3438 {
tmp3434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V710)
}
__typedArg0 := V710
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3434)
}
__typedArg0 := tmp3434
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4custom_1pattern, tmp3435)
}
__typedArg0 := symshen_4custom_1pattern
__typedArg1 := tmp3435
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3423 Obj

if True == tmp3436 {
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

tmp3433 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3432)
}
__typedArg0 := tmp3432
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3425 Obj

if True == tmp3433 {
tmp3427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V710)
}
__typedArg0 := V710
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3427)
}
__typedArg0 := tmp3427
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3428)
}
__typedArg0 := tmp3428
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3429)
}
__typedArg0 := Nil
__typedArg1 := tmp3429
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3426 Obj

if True == tmp3430 {
ifres3426 = True


} else {
ifres3426 = False


}

ifres3425 = ifres3426


} else {
ifres3425 = False


}

var ifres3424 Obj

if True == ifres3425 {
ifres3424 = True


} else {
ifres3424 = False


}

ifres3423 = ifres3424


} else {
ifres3423 = False


}

var ifres3422 Obj

if True == ifres3423 {
ifres3422 = True


} else {
ifres3422 = False


}

ifres3421 = ifres3422


} else {
ifres3421 = False


}

var ifres3420 Obj

if True == ifres3421 {
ifres3420 = True


} else {
ifres3420 = False


}

ifres3419 = ifres3420


} else {
ifres3419 = False


}

var ifres3418 Obj

if True == ifres3419 {
ifres3418 = True


} else {
ifres3418 = False


}

ifres3417 = ifres3418


} else {
ifres3417 = False


}

if True == ifres3417 {
tmp3414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V710)
}
__typedArg0 := V710
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3414)
}
__typedArg0 := tmp3414
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3415)
}
__typedArg0 := tmp3415
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

tmp3442 := Call(__e, ns2_1set, symshen_4custom_1pattern_1body, tmp3413)


_ = tmp3442

tmp3443 := MakeNative(func(__e *ControlFlow) {
V711 := __e.Get(1)
_ = V711
tmp3453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V711)
}
__typedArg0 := V711
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3444 Obj

if True == tmp3453 {
tmp3445 := Call(__e, PrimFunc(symhead), V711)


let__2697 := tmp3445
_ = let__2697

tmp3446 := Call(__e, PrimFunc(symtail), V711)


let__2698 := tmp3446
_ = let__2698

tmp3451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__2697, sym__)
}
__typedArg0 := let__2697
__typedArg1 := sym__
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3447 Obj

if True == tmp3451 {
tmp3448 := Call(__e, PrimFunc(symgensym), symY)


tmp3449 := Call(__e, PrimFunc(symshen_4comb), let__2698, tmp3448)


ifres3447 = tmp3449


} else {
tmp3450 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3447 = tmp3450


}

ifres3444 = ifres3447


} else {
tmp3452 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3444 = tmp3452


}

let__2696 := ifres3444
_ = let__2696

tmp3469 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2696)


if True == tmp3469 {
tmp3465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V711)
}
__typedArg0 := V711
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3454 Obj

if True == tmp3465 {
tmp3455 := Call(__e, PrimFunc(symhead), V711)


let__2700 := tmp3455
_ = let__2700

tmp3456 := Call(__e, PrimFunc(symtail), V711)


let__2701 := tmp3456
_ = let__2701

tmp3460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1, Nil)
}
__typedArg0 := sym_5_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp3460)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp3460
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3462 := Call(__e, PrimFunc(symelement_2), let__2700, tmp3461)


tmp3463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp3462)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp3462
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres3457 Obj

if True == tmp3463 {
tmp3458 := Call(__e, PrimFunc(symshen_4comb), let__2701, let__2700)


ifres3457 = tmp3458


} else {
tmp3459 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3457 = tmp3459


}

ifres3454 = ifres3457


} else {
tmp3464 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3454 = tmp3464


}

let__2699 := ifres3454
_ = let__2699

tmp3467 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2699)


if True == tmp3467 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2699)
return
}


} else {
__e.Return(let__2696)
return
}


}, 1)

tmp3470 := Call(__e, ns2_1set, symshen_4_5simple_1pattern_6, tmp3443)


_ = tmp3470

tmp3471 := MakeNative(func(__e *ControlFlow) {
V718 := __e.Get(1)
_ = V718
tmp3472 := Call(__e, PrimFunc(symshen_4_5pattern_6), V718)


let__2703 := tmp3472
_ = let__2703

tmp3478 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2703)


var ifres3473 Obj

if True == tmp3478 {
tmp3474 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3473 = tmp3474


} else {
tmp3475 := Call(__e, PrimFunc(symshen_4_5_1out), let__2703)


let__2704 := tmp3475
_ = let__2704

tmp3476 := Call(__e, PrimFunc(symshen_4in_1_6), let__2703)


let__2705 := tmp3476
_ = let__2705

tmp3477 := Call(__e, PrimFunc(symshen_4comb), let__2705, let__2704)


ifres3473 = tmp3477


}

let__2702 := ifres3473
_ = let__2702

tmp3480 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2702)


if True == tmp3480 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2702)
return
}


}, 1)

tmp3481 := Call(__e, ns2_1set, symshen_4_5pattern1_6, tmp3471)


_ = tmp3481

tmp3482 := MakeNative(func(__e *ControlFlow) {
V723 := __e.Get(1)
_ = V723
tmp3483 := Call(__e, PrimFunc(symshen_4_5pattern_6), V723)


let__2707 := tmp3483
_ = let__2707

tmp3489 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2707)


var ifres3484 Obj

if True == tmp3489 {
tmp3485 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3484 = tmp3485


} else {
tmp3486 := Call(__e, PrimFunc(symshen_4_5_1out), let__2707)


let__2708 := tmp3486
_ = let__2708

tmp3487 := Call(__e, PrimFunc(symshen_4in_1_6), let__2707)


let__2709 := tmp3487
_ = let__2709

tmp3488 := Call(__e, PrimFunc(symshen_4comb), let__2709, let__2708)


ifres3484 = tmp3488


}

let__2706 := ifres3484
_ = let__2706

tmp3491 := Call(__e, PrimFunc(symshen_4parse_1failure_2), let__2706)


if True == tmp3491 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(let__2706)
return
}


}, 1)

tmp3492 := Call(__e, ns2_1set, symshen_4_5pattern2_6, tmp3482)


_ = tmp3492

tmp3493 := MakeNative(func(__e *ControlFlow) {
V728 := __e.Get(1)
_ = V728
tmp3494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(MakeInteger(2))
}
__typedArg0 := MakeInteger(2)
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

let__2710 := tmp3494
_ = let__2710

tmp3495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__2710, MakeInteger(0), symshen_4printF)
}
__typedArg0 := let__2710
__typedArg1 := MakeInteger(0)
__typedArg2 := symshen_4printF
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__2711 := tmp3495
_ = let__2711

tmp3496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V728)
}
__typedArg0 := V728
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp3497 := Call(__e, PrimFunc(sym_8s), tmp3496, MakeString(")"))


tmp3498 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp3497)


tmp3499 := Call(__e, PrimFunc(sym_8s), MakeString("n"), tmp3498)


tmp3500 := Call(__e, PrimFunc(sym_8s), MakeString("f"), tmp3499)


tmp3501 := Call(__e, PrimFunc(sym_8s), MakeString("("), tmp3500)


tmp3502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(let__2711, MakeInteger(1), tmp3501)
}
__typedArg0 := let__2711
__typedArg1 := MakeInteger(1)
__typedArg2 := tmp3501
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

let__2712 := tmp3502
_ = let__2712

__e.Return(let__2712)
return


}, 1)

tmp3503 := Call(__e, ns2_1set, symshen_4fn_1print, tmp3493)


_ = tmp3503

tmp3504 := MakeNative(func(__e *ControlFlow) {
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

tmp3505 := Call(__e, ns2_1set, symshen_4printF, tmp3504)


_ = tmp3505

tmp3506 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V737 := __self1
_ = V737
V738 := __self2
_ = V738
tmp3530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3526 Obj

if True == tmp3530 {
tmp3528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3528)
}
__typedArg0 := Nil
__typedArg1 := tmp3528
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3527 Obj

if True == tmp3529 {
ifres3527 = True


} else {
ifres3527 = False


}

ifres3526 = ifres3527


} else {
ifres3526 = False


}

if True == ifres3526 {
tmp3507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symlength), tmp3507)
return


} else {
tmp3524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3512 Obj

if True == tmp3524 {
tmp3522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3522)
}
__typedArg0 := tmp3522
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3514 Obj

if True == tmp3523 {
tmp3516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3517 := Call(__e, PrimFunc(symlength), tmp3516)


tmp3518 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3519 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3518)
}
__typedArg0 := tmp3518
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3520 := Call(__e, PrimFunc(symlength), tmp3519)


tmp3521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp3517, tmp3520)
}
__typedArg0 := tmp3517
__typedArg1 := tmp3520
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3515 Obj

if True == tmp3521 {
ifres3515 = True


} else {
ifres3515 = False


}

ifres3514 = ifres3515


} else {
ifres3514 = False


}

var ifres3513 Obj

if True == ifres3514 {
ifres3513 = True


} else {
ifres3513 = False


}

ifres3512 = ifres3513


} else {
ifres3512 = False


}

if True == ifres3512 {
tmp3508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4arity_1chk) == __self {
__self1, __self2 = V737, tmp3508
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4arity_1chk), V737, tmp3508)
return


} else {
tmp3509 := Call(__e, PrimFunc(symshen_4app), V737, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("arity error in "))
__typedS1, __typedOK1 := TypedString(tmp3509)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("arity error in ")
__typedArg1 := tmp3509
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("arity error in "))
__typedS1, __typedOK1 := TypedString(tmp3509)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("arity error in ")
__typedArg1 := tmp3509
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}


}, 2)

tmp3531 := Call(__e, ns2_1set, symshen_4arity_1chk, tmp3506)


_ = tmp3531

tmp3532 := MakeNative(func(__e *ControlFlow) {
V739 := __e.Get(1)
_ = V739
V740 := __e.Get(2)
_ = V740
tmp3538 := Call(__e, PrimFunc(symtuple_2), V740)


if True == tmp3538 {
tmp3533 := Call(__e, PrimFunc(symfst), V740)


tmp3534 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp3533)


tmp3535 := Call(__e, PrimFunc(symsnd), V740)


tmp3536 := Call(__e, PrimFunc(symshen_4find_1free_1vars), tmp3534, tmp3535)


__e.TailApply(PrimFunc(symshen_4free_1variable_1error_1message), V739, tmp3536)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4free_1var_1chk)
return
}


}, 2)

tmp3539 := Call(__e, ns2_1set, symshen_4free_1var_1chk, tmp3532)


_ = tmp3539

tmp3540 := MakeNative(func(__e *ControlFlow) {
V741 := __e.Get(1)
_ = V741
V742 := __e.Get(2)
_ = V742
tmp3552 := Call(__e, PrimFunc(symempty_2), V742)


if True == tmp3552 {
__e.Return(symshen_4skip)
return
} else {
tmp3541 := Call(__e, PrimFunc(symshen_4app), V741, MakeString(":"), symshen_4a)


tmp3542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("free variables in "))
__typedS1, __typedOK1 := TypedString(tmp3541)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("free variables in ")
__typedArg1 := tmp3541
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp3543 := Call(__e, PrimFunc(symstoutput))


tmp3544 := Call(__e, PrimFunc(sympr), tmp3542, tmp3543)


_ = tmp3544

tmp3545 := MakeNative(func(__e *ControlFlow) {
Z743 := __e.Get(1)
_ = Z743
tmp3546 := Call(__e, PrimFunc(symshen_4app), Z743, MakeString(""), symshen_4a)


tmp3547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" "))
__typedS1, __typedOK1 := TypedString(tmp3546)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" ")
__typedArg1 := tmp3546
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp3548 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp3547, tmp3548)
return


}, 1)

tmp3549 := Call(__e, PrimFunc(symshen_4for_1each), tmp3545, V742)


_ = tmp3549

tmp3550 := Call(__e, PrimFunc(symnl), MakeInteger(1))


_ = tmp3550

__e.TailApply(PrimFunc(symabort))
return


}


}, 2)

tmp3553 := Call(__e, ns2_1set, symshen_4free_1variable_1error_1message, tmp3540)


_ = tmp3553

tmp3554 := MakeNative(func(__e *ControlFlow) {
V746 := __e.Get(1)
_ = V746
tmp3562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V746)
}
__typedArg0 := V746
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp3562 {
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
tmp3560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V746)
}
__typedArg0 := V746
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3560 {
tmp3555 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V746)
}
__typedArg0 := V746
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3556 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp3555)


tmp3557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V746)
}
__typedArg0 := V746
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3558 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp3557)


__e.TailApply(PrimFunc(symunion), tmp3556, tmp3558)
return


} else {
__e.Return(Nil)
return
}


}


}, 1)

tmp3563 := Call(__e, ns2_1set, symshen_4extract_1vars, tmp3554)


_ = tmp3563

tmp3564 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V751 := __self1
_ = V751
V752 := __self2
_ = V752
tmp3654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3641 Obj

if True == tmp3654 {
tmp3652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symprotect, tmp3652)
}
__typedArg0 := symprotect
__typedArg1 := tmp3652
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3643 Obj

if True == tmp3653 {
tmp3650 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3650)
}
__typedArg0 := tmp3650
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3645 Obj

if True == tmp3651 {
tmp3647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3647)
}
__typedArg0 := tmp3647
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3648)
}
__typedArg0 := Nil
__typedArg1 := tmp3648
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3646 Obj

if True == tmp3649 {
ifres3646 = True


} else {
ifres3646 = False


}

ifres3645 = ifres3646


} else {
ifres3645 = False


}

var ifres3644 Obj

if True == ifres3645 {
ifres3644 = True


} else {
ifres3644 = False


}

ifres3643 = ifres3644


} else {
ifres3643 = False


}

var ifres3642 Obj

if True == ifres3643 {
ifres3642 = True


} else {
ifres3642 = False


}

ifres3641 = ifres3642


} else {
ifres3641 = False


}

if True == ifres3641 {
__e.Return(Nil)
return
} else {
tmp3639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3613 Obj

if True == tmp3639 {
tmp3637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp3637)
}
__typedArg0 := symlet
__typedArg1 := tmp3637
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3615 Obj

if True == tmp3638 {
tmp3635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3635)
}
__typedArg0 := tmp3635
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3617 Obj

if True == tmp3636 {
tmp3632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3632)
}
__typedArg0 := tmp3632
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3633)
}
__typedArg0 := tmp3633
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3619 Obj

if True == tmp3634 {
tmp3628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3628)
}
__typedArg0 := tmp3628
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3629)
}
__typedArg0 := tmp3629
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3630)
}
__typedArg0 := tmp3630
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3621 Obj

if True == tmp3631 {
tmp3623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3623)
}
__typedArg0 := tmp3623
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3624)
}
__typedArg0 := tmp3624
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3625)
}
__typedArg0 := tmp3625
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3626)
}
__typedArg0 := Nil
__typedArg1 := tmp3626
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3622 Obj

if True == tmp3627 {
ifres3622 = True


} else {
ifres3622 = False


}

ifres3621 = ifres3622


} else {
ifres3621 = False


}

var ifres3620 Obj

if True == ifres3621 {
ifres3620 = True


} else {
ifres3620 = False


}

ifres3619 = ifres3620


} else {
ifres3619 = False


}

var ifres3618 Obj

if True == ifres3619 {
ifres3618 = True


} else {
ifres3618 = False


}

ifres3617 = ifres3618


} else {
ifres3617 = False


}

var ifres3616 Obj

if True == ifres3617 {
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

if True == ifres3613 {
tmp3565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3566)
}
__typedArg0 := tmp3566
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3568 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V751, tmp3567)


tmp3569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3569)
}
__typedArg0 := tmp3569
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3570, V751)
}
__typedArg0 := tmp3570
__typedArg1 := V751
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3572)
}
__typedArg0 := tmp3572
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3573)
}
__typedArg0 := tmp3573
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3574)
}
__typedArg0 := tmp3574
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3576 := Call(__e, PrimFunc(symshen_4find_1free_1vars), tmp3571, tmp3575)


__e.TailApply(PrimFunc(symunion), tmp3568, tmp3576)
return


} else {
tmp3611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3592 Obj

if True == tmp3611 {
tmp3609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp3609)
}
__typedArg0 := symlambda
__typedArg1 := tmp3609
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3594 Obj

if True == tmp3610 {
tmp3607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3607)
}
__typedArg0 := tmp3607
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3596 Obj

if True == tmp3608 {
tmp3604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3604)
}
__typedArg0 := tmp3604
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3605)
}
__typedArg0 := tmp3605
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3598 Obj

if True == tmp3606 {
tmp3600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3600)
}
__typedArg0 := tmp3600
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3601)
}
__typedArg0 := tmp3601
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3602)
}
__typedArg0 := Nil
__typedArg1 := tmp3602
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3599 Obj

if True == tmp3603 {
ifres3599 = True


} else {
ifres3599 = False


}

ifres3598 = ifres3599


} else {
ifres3598 = False


}

var ifres3597 Obj

if True == ifres3598 {
ifres3597 = True


} else {
ifres3597 = False


}

ifres3596 = ifres3597


} else {
ifres3596 = False


}

var ifres3595 Obj

if True == ifres3596 {
ifres3595 = True


} else {
ifres3595 = False


}

ifres3594 = ifres3595


} else {
ifres3594 = False


}

var ifres3593 Obj

if True == ifres3594 {
ifres3593 = True


} else {
ifres3593 = False


}

ifres3592 = ifres3593


} else {
ifres3592 = False


}

if True == ifres3592 {
tmp3577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3577)
}
__typedArg0 := tmp3577
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3578, V751)
}
__typedArg0 := tmp3578
__typedArg1 := V751
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3580)
}
__typedArg0 := tmp3580
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3581)
}
__typedArg0 := tmp3581
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

if PrimFunc(symshen_4find_1free_1vars) == __self {
__self1, __self2 = tmp3579, tmp3582
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4find_1free_1vars), tmp3579, tmp3582)
return


} else {
tmp3590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3590 {
tmp3583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3584 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V751, tmp3583)


tmp3585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3586 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V751, tmp3585)


__e.TailApply(PrimFunc(symunion), tmp3584, tmp3586)
return


} else {
tmp3588 := Call(__e, PrimFunc(symshen_4free_1variable_2), V752, V751)


if True == tmp3588 {
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

tmp3655 := Call(__e, ns2_1set, symshen_4find_1free_1vars, tmp3564)


_ = tmp3655

tmp3656 := MakeNative(func(__e *ControlFlow) {
V753 := __e.Get(1)
_ = V753
V754 := __e.Get(2)
_ = V754
tmp3661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V753)
}
__typedArg0 := V753
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp3661 {
tmp3658 := Call(__e, PrimFunc(symelement_2), V753, V754)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp3658)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp3658
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

tmp3662 := Call(__e, ns2_1set, symshen_4free_1variable_2, tmp3656)


_ = tmp3662

tmp3663 := MakeNative(func(__e *ControlFlow) {
V755 := __e.Get(1)
_ = V755
V756 := __e.Get(2)
_ = V756
tmp3664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_duserdefs_d)
}
__typedArg0 := symshen_4_duserdefs_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp3665 := Call(__e, PrimFunc(symadjoin), V755, tmp3664)


tmp3666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_duserdefs_d, tmp3665)
}
__typedArg0 := symshen_4_duserdefs_d
__typedArg1 := tmp3665
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp3666

tmp3667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V755, symshen_4source, V756, tmp3667)
return


}, 2)

tmp3668 := Call(__e, ns2_1set, symshen_4record_1kl, tmp3663)


_ = tmp3668

tmp3669 := MakeNative(func(__e *ControlFlow) {
V757 := __e.Get(1)
_ = V757
V758 := __e.Get(2)
_ = V758
V759 := __e.Get(3)
_ = V759
tmp3670 := Call(__e, PrimFunc(symshen_4parameters), V759)


let__2713 := tmp3670
_ = let__2713

tmp3671 := Call(__e, PrimFunc(symshen_4kl_1body), V758, let__2713)


tmp3672 := Call(__e, PrimFunc(symshen_4scan_1body), V757, tmp3671)


let__2714 := tmp3672
_ = let__2714

tmp3673 := Call(__e, PrimFunc(symshen_4cond_1form), let__2714)


tmp3674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3673, Nil)
}
__typedArg0 := tmp3673
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2713, tmp3674)
}
__typedArg0 := let__2713
__typedArg1 := tmp3674
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V757, tmp3675)
}
__typedArg0 := V757
__typedArg1 := tmp3675
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp3676)
}
__typedArg0 := symdefun
__typedArg1 := tmp3676
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__2715 := tmp3677
_ = let__2715

__e.Return(let__2715)
return


}, 3)

tmp3678 := Call(__e, ns2_1set, symshen_4compile_1to_1kl, tmp3669)


_ = tmp3678

tmp3679 := MakeNative(func(__e *ControlFlow) {
V763 := __e.Get(1)
_ = V763
tmp3684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeInteger(0), V763)
}
__typedArg0 := MakeInteger(0)
__typedArg1 := V763
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp3684 {
__e.Return(Nil)
return
} else {
tmp3680 := Call(__e, PrimFunc(symgensym), symV)


tmp3682 := Call(__e, PrimFunc(symshen_4parameters), (func() Obj {
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
return PrimCons(tmp3680, tmp3682)
}
__typedArg0 := tmp3680
__typedArg1 := tmp3682
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp3685 := Call(__e, ns2_1set, symshen_4parameters, tmp3679)


_ = tmp3685

tmp3686 := MakeNative(func(__e *ControlFlow) {
V766 := __e.Get(1)
_ = V766
tmp3710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V766)
}
__typedArg0 := V766
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3690 Obj

if True == tmp3710 {
tmp3708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V766)
}
__typedArg0 := V766
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3708)
}
__typedArg0 := tmp3708
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3692 Obj

if True == tmp3709 {
tmp3705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V766)
}
__typedArg0 := V766
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3705)
}
__typedArg0 := tmp3705
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, tmp3706)
}
__typedArg0 := True
__typedArg1 := tmp3706
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3694 Obj

if True == tmp3707 {
tmp3702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V766)
}
__typedArg0 := V766
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3702)
}
__typedArg0 := tmp3702
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3703)
}
__typedArg0 := tmp3703
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3696 Obj

if True == tmp3704 {
tmp3698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V766)
}
__typedArg0 := V766
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3698)
}
__typedArg0 := tmp3698
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3699)
}
__typedArg0 := tmp3699
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3700)
}
__typedArg0 := Nil
__typedArg1 := tmp3700
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3697 Obj

if True == tmp3701 {
ifres3697 = True


} else {
ifres3697 = False


}

ifres3696 = ifres3697


} else {
ifres3696 = False


}

var ifres3695 Obj

if True == ifres3696 {
ifres3695 = True


} else {
ifres3695 = False


}

ifres3694 = ifres3695


} else {
ifres3694 = False


}

var ifres3693 Obj

if True == ifres3694 {
ifres3693 = True


} else {
ifres3693 = False


}

ifres3692 = ifres3693


} else {
ifres3692 = False


}

var ifres3691 Obj

if True == ifres3692 {
ifres3691 = True


} else {
ifres3691 = False


}

ifres3690 = ifres3691


} else {
ifres3690 = False


}

if True == ifres3690 {
tmp3687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V766)
}
__typedArg0 := V766
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3687)
}
__typedArg0 := tmp3687
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3688)
}
__typedArg0 := tmp3688
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

tmp3711 := Call(__e, ns2_1set, symshen_4cond_1form, tmp3686)


_ = tmp3711

tmp3712 := MakeNative(func(__e *ControlFlow) {
V775 := __e.Get(1)
_ = V775
V776 := __e.Get(2)
_ = V776
tmp3756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V776)
}
__typedArg0 := Nil
__typedArg1 := V776
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp3756 {
tmp3713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V775, Nil)
}
__typedArg0 := V775
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4f_1error, tmp3713)
}
__typedArg0 := symshen_4f_1error
__typedArg1 := tmp3713
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3714, Nil)
}
__typedArg0 := tmp3714
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp3715)
}
__typedArg0 := True
__typedArg1 := tmp3715
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3716, Nil)
}
__typedArg0 := tmp3716
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp3754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3750 Obj

if True == tmp3754 {
tmp3752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3753 := Call(__e, PrimFunc(symshen_4choicepoint_2), tmp3752)


var ifres3751 Obj

if True == tmp3753 {
ifres3751 = True


} else {
ifres3751 = False


}

ifres3750 = ifres3751


} else {
ifres3750 = False


}

if True == ifres3750 {
tmp3717 := Call(__e, PrimFunc(symgensym), symFreeze)


tmp3718 := Call(__e, PrimFunc(symgensym), symResult)


tmp3719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4choicepoint), V775, tmp3717, tmp3718, tmp3719, tmp3720)
return


} else {
tmp3748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3728 Obj

if True == tmp3748 {
tmp3746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3746)
}
__typedArg0 := tmp3746
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3730 Obj

if True == tmp3747 {
tmp3743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3743)
}
__typedArg0 := tmp3743
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, tmp3744)
}
__typedArg0 := True
__typedArg1 := tmp3744
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3732 Obj

if True == tmp3745 {
tmp3740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3740)
}
__typedArg0 := tmp3740
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3741)
}
__typedArg0 := tmp3741
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3734 Obj

if True == tmp3742 {
tmp3736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3736)
}
__typedArg0 := tmp3736
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3737)
}
__typedArg0 := tmp3737
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3738)
}
__typedArg0 := Nil
__typedArg1 := tmp3738
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3735 Obj

if True == tmp3739 {
ifres3735 = True


} else {
ifres3735 = False


}

ifres3734 = ifres3735


} else {
ifres3734 = False


}

var ifres3733 Obj

if True == ifres3734 {
ifres3733 = True


} else {
ifres3733 = False


}

ifres3732 = ifres3733


} else {
ifres3732 = False


}

var ifres3731 Obj

if True == ifres3732 {
ifres3731 = True


} else {
ifres3731 = False


}

ifres3730 = ifres3731


} else {
ifres3730 = False


}

var ifres3729 Obj

if True == ifres3730 {
ifres3729 = True


} else {
ifres3729 = False


}

ifres3728 = ifres3729


} else {
ifres3728 = False


}

if True == ifres3728 {
tmp3721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3721, Nil)
}
__typedArg0 := tmp3721
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp3726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3726 {
tmp3722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V776)
}
__typedArg0 := V776
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3724 := Call(__e, PrimFunc(symshen_4scan_1body), V775, tmp3723)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3722, tmp3724)
}
__typedArg0 := tmp3722
__typedArg1 := tmp3724
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

tmp3757 := Call(__e, ns2_1set, symshen_4scan_1body, tmp3712)


_ = tmp3757

tmp3758 := MakeNative(func(__e *ControlFlow) {
V783 := __e.Get(1)
_ = V783
tmp3793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V783)
}
__typedArg0 := V783
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3760 Obj

if True == tmp3793 {
tmp3791 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V783)
}
__typedArg0 := V783
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3792 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3791)
}
__typedArg0 := tmp3791
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3762 Obj

if True == tmp3792 {
tmp3788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V783)
}
__typedArg0 := V783
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3788)
}
__typedArg0 := tmp3788
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3789)
}
__typedArg0 := tmp3789
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3764 Obj

if True == tmp3790 {
tmp3784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V783)
}
__typedArg0 := V783
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3784)
}
__typedArg0 := tmp3784
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3785)
}
__typedArg0 := tmp3785
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4choicepoint_b, tmp3786)
}
__typedArg0 := symshen_4choicepoint_b
__typedArg1 := tmp3786
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3766 Obj

if True == tmp3787 {
tmp3780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V783)
}
__typedArg0 := V783
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3780)
}
__typedArg0 := tmp3780
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3781)
}
__typedArg0 := tmp3781
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3782)
}
__typedArg0 := tmp3782
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3768 Obj

if True == tmp3783 {
tmp3775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V783)
}
__typedArg0 := V783
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3775)
}
__typedArg0 := tmp3775
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3776)
}
__typedArg0 := tmp3776
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3777)
}
__typedArg0 := tmp3777
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3778)
}
__typedArg0 := Nil
__typedArg1 := tmp3778
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3770 Obj

if True == tmp3779 {
tmp3772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V783)
}
__typedArg0 := V783
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3772)
}
__typedArg0 := tmp3772
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3773)
}
__typedArg0 := Nil
__typedArg1 := tmp3773
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3771 Obj

if True == tmp3774 {
ifres3771 = True


} else {
ifres3771 = False


}

ifres3770 = ifres3771


} else {
ifres3770 = False


}

var ifres3769 Obj

if True == ifres3770 {
ifres3769 = True


} else {
ifres3769 = False


}

ifres3768 = ifres3769


} else {
ifres3768 = False


}

var ifres3767 Obj

if True == ifres3768 {
ifres3767 = True


} else {
ifres3767 = False


}

ifres3766 = ifres3767


} else {
ifres3766 = False


}

var ifres3765 Obj

if True == ifres3766 {
ifres3765 = True


} else {
ifres3765 = False


}

ifres3764 = ifres3765


} else {
ifres3764 = False


}

var ifres3763 Obj

if True == ifres3764 {
ifres3763 = True


} else {
ifres3763 = False


}

ifres3762 = ifres3763


} else {
ifres3762 = False


}

var ifres3761 Obj

if True == ifres3762 {
ifres3761 = True


} else {
ifres3761 = False


}

ifres3760 = ifres3761


} else {
ifres3760 = False


}

if True == ifres3760 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp3794 := Call(__e, ns2_1set, symshen_4choicepoint_2, tmp3758)


_ = tmp3794

tmp3795 := MakeNative(func(__e *ControlFlow) {
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
tmp3987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3909 Obj

if True == tmp3987 {
tmp3985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3985)
}
__typedArg0 := tmp3985
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3911 Obj

if True == tmp3986 {
tmp3982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3982)
}
__typedArg0 := tmp3982
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3983)
}
__typedArg0 := tmp3983
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3913 Obj

if True == tmp3984 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3980)
}
__typedArg0 := tmp3980
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3915 Obj

if True == tmp3981 {
tmp3973 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3975)
}
__typedArg0 := tmp3975
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3976)
}
__typedArg0 := tmp3976
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3917 Obj

if True == tmp3977 {
tmp3967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3967)
}
__typedArg0 := tmp3967
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3968)
}
__typedArg0 := tmp3968
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3969)
}
__typedArg0 := tmp3969
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3970)
}
__typedArg0 := tmp3970
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3972 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail_1if, tmp3971)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp3971
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3919 Obj

if True == tmp3972 {
tmp3961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3962 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3961)
}
__typedArg0 := tmp3961
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3962)
}
__typedArg0 := tmp3962
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3965)
}
__typedArg0 := tmp3965
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3921 Obj

if True == tmp3966 {
tmp3954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3955 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3954)
}
__typedArg0 := tmp3954
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3956 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3955)
}
__typedArg0 := tmp3955
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3956)
}
__typedArg0 := tmp3956
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3957)
}
__typedArg0 := tmp3957
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3959 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3958)
}
__typedArg0 := tmp3958
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3959)
}
__typedArg0 := tmp3959
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3923 Obj

if True == tmp3960 {
tmp3946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3946)
}
__typedArg0 := tmp3946
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3947)
}
__typedArg0 := tmp3947
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3950)
}
__typedArg0 := tmp3950
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3951)
}
__typedArg0 := tmp3951
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3952)
}
__typedArg0 := Nil
__typedArg1 := tmp3952
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3925 Obj

if True == tmp3953 {
tmp3941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3941)
}
__typedArg0 := tmp3941
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3942)
}
__typedArg0 := tmp3942
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3943)
}
__typedArg0 := tmp3943
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3944)
}
__typedArg0 := Nil
__typedArg1 := tmp3944
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3927 Obj

if True == tmp3945 {
tmp3938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3938)
}
__typedArg0 := tmp3938
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3939)
}
__typedArg0 := Nil
__typedArg1 := tmp3939
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3929 Obj

if True == tmp3940 {
tmp3931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3931)
}
__typedArg0 := tmp3931
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3932)
}
__typedArg0 := tmp3932
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3933)
}
__typedArg0 := tmp3933
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3934)
}
__typedArg0 := tmp3934
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3935)
}
__typedArg0 := tmp3935
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V799, tmp3936)
}
__typedArg0 := V799
__typedArg1 := tmp3936
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3930 Obj

if True == tmp3937 {
ifres3930 = True


} else {
ifres3930 = False


}

ifres3929 = ifres3930


} else {
ifres3929 = False


}

var ifres3928 Obj

if True == ifres3929 {
ifres3928 = True


} else {
ifres3928 = False


}

ifres3927 = ifres3928


} else {
ifres3927 = False


}

var ifres3926 Obj

if True == ifres3927 {
ifres3926 = True


} else {
ifres3926 = False


}

ifres3925 = ifres3926


} else {
ifres3925 = False


}

var ifres3924 Obj

if True == ifres3925 {
ifres3924 = True


} else {
ifres3924 = False


}

ifres3923 = ifres3924


} else {
ifres3923 = False


}

var ifres3922 Obj

if True == ifres3923 {
ifres3922 = True


} else {
ifres3922 = False


}

ifres3921 = ifres3922


} else {
ifres3921 = False


}

var ifres3920 Obj

if True == ifres3921 {
ifres3920 = True


} else {
ifres3920 = False


}

ifres3919 = ifres3920


} else {
ifres3919 = False


}

var ifres3918 Obj

if True == ifres3919 {
ifres3918 = True


} else {
ifres3918 = False


}

ifres3917 = ifres3918


} else {
ifres3917 = False


}

var ifres3916 Obj

if True == ifres3917 {
ifres3916 = True


} else {
ifres3916 = False


}

ifres3915 = ifres3916


} else {
ifres3915 = False


}

var ifres3914 Obj

if True == ifres3915 {
ifres3914 = True


} else {
ifres3914 = False


}

ifres3913 = ifres3914


} else {
ifres3913 = False


}

var ifres3912 Obj

if True == ifres3913 {
ifres3912 = True


} else {
ifres3912 = False


}

ifres3911 = ifres3912


} else {
ifres3911 = False


}

var ifres3910 Obj

if True == ifres3911 {
ifres3910 = True


} else {
ifres3910 = False


}

ifres3909 = ifres3910


} else {
ifres3909 = False


}

if True == ifres3909 {
tmp3796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3796)
}
__typedArg0 := tmp3796
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3797)
}
__typedArg0 := tmp3797
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3798)
}
__typedArg0 := tmp3798
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3799)
}
__typedArg0 := tmp3799
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3800)
}
__typedArg0 := tmp3800
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3802 := Call(__e, PrimFunc(symshen_4scan_1body), tmp3801, V803)


tmp3803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcond, tmp3802)
}
__typedArg0 := symcond
__typedArg1 := tmp3802
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3803, Nil)
}
__typedArg0 := tmp3803
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp3804)
}
__typedArg0 := symfreeze
__typedArg1 := tmp3804
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3807)
}
__typedArg0 := tmp3807
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3808)
}
__typedArg0 := tmp3808
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3809)
}
__typedArg0 := tmp3809
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3810)
}
__typedArg0 := tmp3810
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3811)
}
__typedArg0 := tmp3811
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3812)
}
__typedArg0 := tmp3812
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3814)
}
__typedArg0 := tmp3814
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3815)
}
__typedArg0 := tmp3815
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3816)
}
__typedArg0 := tmp3816
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3817)
}
__typedArg0 := tmp3817
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3818)
}
__typedArg0 := tmp3818
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V801, Nil)
}
__typedArg0 := V801
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3819, tmp3820)
}
__typedArg0 := tmp3819
__typedArg1 := tmp3820
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V800, Nil)
}
__typedArg0 := V800
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp3822)
}
__typedArg0 := symthaw
__typedArg1 := tmp3822
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V801, Nil)
}
__typedArg0 := V801
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3823, tmp3824)
}
__typedArg0 := tmp3823
__typedArg1 := tmp3824
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3821, tmp3825)
}
__typedArg0 := tmp3821
__typedArg1 := tmp3825
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp3826)
}
__typedArg0 := symif
__typedArg1 := tmp3826
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3827, Nil)
}
__typedArg0 := tmp3827
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3813, tmp3828)
}
__typedArg0 := tmp3813
__typedArg1 := tmp3828
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V801, tmp3829)
}
__typedArg0 := V801
__typedArg1 := tmp3829
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp3830)
}
__typedArg0 := symlet
__typedArg1 := tmp3830
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V800, Nil)
}
__typedArg0 := V800
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp3832)
}
__typedArg0 := symthaw
__typedArg1 := tmp3832
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3834 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3833, Nil)
}
__typedArg0 := tmp3833
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3835 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3831, tmp3834)
}
__typedArg0 := tmp3831
__typedArg1 := tmp3834
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3836 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3806, tmp3835)
}
__typedArg0 := tmp3806
__typedArg1 := tmp3835
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp3836)
}
__typedArg0 := symif
__typedArg1 := tmp3836
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3837, Nil)
}
__typedArg0 := tmp3837
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3805, tmp3838)
}
__typedArg0 := tmp3805
__typedArg1 := tmp3838
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V800, tmp3839)
}
__typedArg0 := V800
__typedArg1 := tmp3839
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp3840)
}
__typedArg0 := symlet
__typedArg1 := tmp3840
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3841, Nil)
}
__typedArg0 := tmp3841
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp3842)
}
__typedArg0 := True
__typedArg1 := tmp3842
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3843, Nil)
}
__typedArg0 := tmp3843
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp3907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3880 Obj

if True == tmp3907 {
tmp3905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3905)
}
__typedArg0 := tmp3905
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3882 Obj

if True == tmp3906 {
tmp3902 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3903 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3902)
}
__typedArg0 := tmp3902
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3904 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3903)
}
__typedArg0 := tmp3903
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3884 Obj

if True == tmp3904 {
tmp3898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3898)
}
__typedArg0 := tmp3898
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3899)
}
__typedArg0 := tmp3899
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3901 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3900)
}
__typedArg0 := tmp3900
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3886 Obj

if True == tmp3901 {
tmp3893 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3894 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3893)
}
__typedArg0 := tmp3893
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3895 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3894)
}
__typedArg0 := tmp3894
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3895)
}
__typedArg0 := tmp3895
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3896)
}
__typedArg0 := Nil
__typedArg1 := tmp3896
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3888 Obj

if True == tmp3897 {
tmp3890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3890)
}
__typedArg0 := tmp3890
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3891)
}
__typedArg0 := Nil
__typedArg1 := tmp3891
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3889 Obj

if True == tmp3892 {
ifres3889 = True


} else {
ifres3889 = False


}

ifres3888 = ifres3889


} else {
ifres3888 = False


}

var ifres3887 Obj

if True == ifres3888 {
ifres3887 = True


} else {
ifres3887 = False


}

ifres3886 = ifres3887


} else {
ifres3886 = False


}

var ifres3885 Obj

if True == ifres3886 {
ifres3885 = True


} else {
ifres3885 = False


}

ifres3884 = ifres3885


} else {
ifres3884 = False


}

var ifres3883 Obj

if True == ifres3884 {
ifres3883 = True


} else {
ifres3883 = False


}

ifres3882 = ifres3883


} else {
ifres3882 = False


}

var ifres3881 Obj

if True == ifres3882 {
ifres3881 = True


} else {
ifres3881 = False


}

ifres3880 = ifres3881


} else {
ifres3880 = False


}

if True == ifres3880 {
tmp3844 := Call(__e, PrimFunc(symshen_4scan_1body), V799, V803)


tmp3845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcond, tmp3844)
}
__typedArg0 := symcond
__typedArg1 := tmp3844
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3845, Nil)
}
__typedArg0 := tmp3845
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp3846)
}
__typedArg0 := symfreeze
__typedArg1 := tmp3846
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V802)
}
__typedArg0 := V802
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3849)
}
__typedArg0 := tmp3849
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3850)
}
__typedArg0 := tmp3850
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3851)
}
__typedArg0 := tmp3851
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, Nil)
}
__typedArg0 := symfail
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3853, Nil)
}
__typedArg0 := tmp3853
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V801, tmp3854)
}
__typedArg0 := V801
__typedArg1 := tmp3854
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp3855)
}
__typedArg0 := sym_a
__typedArg1 := tmp3855
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V800, Nil)
}
__typedArg0 := V800
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp3857)
}
__typedArg0 := symthaw
__typedArg1 := tmp3857
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V801, Nil)
}
__typedArg0 := V801
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3858, tmp3859)
}
__typedArg0 := tmp3858
__typedArg1 := tmp3859
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3856, tmp3860)
}
__typedArg0 := tmp3856
__typedArg1 := tmp3860
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp3861)
}
__typedArg0 := symif
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
return PrimCons(tmp3852, tmp3863)
}
__typedArg0 := tmp3852
__typedArg1 := tmp3863
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V801, tmp3864)
}
__typedArg0 := V801
__typedArg1 := tmp3864
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp3865)
}
__typedArg0 := symlet
__typedArg1 := tmp3865
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V800, Nil)
}
__typedArg0 := V800
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp3867)
}
__typedArg0 := symthaw
__typedArg1 := tmp3867
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3868, Nil)
}
__typedArg0 := tmp3868
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3866, tmp3869)
}
__typedArg0 := tmp3866
__typedArg1 := tmp3869
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3848, tmp3870)
}
__typedArg0 := tmp3848
__typedArg1 := tmp3870
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp3871)
}
__typedArg0 := symif
__typedArg1 := tmp3871
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3872, Nil)
}
__typedArg0 := tmp3872
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3847, tmp3873)
}
__typedArg0 := tmp3847
__typedArg1 := tmp3873
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V800, tmp3874)
}
__typedArg0 := V800
__typedArg1 := tmp3874
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp3875)
}
__typedArg0 := symlet
__typedArg1 := tmp3875
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3877 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3876, Nil)
}
__typedArg0 := tmp3876
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp3877)
}
__typedArg0 := True
__typedArg1 := tmp3877
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3878, Nil)
}
__typedArg0 := tmp3878
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

tmp3988 := Call(__e, ns2_1set, symshen_4choicepoint, tmp3795)


_ = tmp3988

tmp3989 := MakeNative(func(__e *ControlFlow) {
V805 := __e.Get(1)
_ = V805
V806 := __e.Get(2)
_ = V806
V807 := __e.Get(3)
_ = V807
tmp4002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V805, V807)
}
__typedArg0 := V805
__typedArg1 := V807
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4002 {
__e.Return(V806)
return
} else {
tmp4000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V807)
}
__typedArg0 := V807
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp4000 {
tmp3990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V807)
}
__typedArg0 := V807
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3991 := Call(__e, PrimFunc(symshen_4rep_1X), V805, V806, tmp3990)


let__2716 := tmp3991
_ = let__2716

tmp3997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V807)
}
__typedArg0 := V807
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(let__2716, tmp3997)
}
__typedArg0 := let__2716
__typedArg1 := tmp3997
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp3998 {
tmp3992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V807)
}
__typedArg0 := V807
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V807)
}
__typedArg0 := V807
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3994 := Call(__e, PrimFunc(symshen_4rep_1X), V805, V806, tmp3993)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3992, tmp3994)
}
__typedArg0 := tmp3992
__typedArg1 := tmp3994
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp3995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V807)
}
__typedArg0 := V807
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2716, tmp3995)
}
__typedArg0 := let__2716
__typedArg1 := tmp3995
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

tmp4003 := Call(__e, ns2_1set, symshen_4rep_1X, tmp3989)


_ = tmp4003

tmp4004 := MakeNative(func(__e *ControlFlow) {
V809 := __e.Get(1)
_ = V809
tmp4083 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4064 Obj

if True == tmp4083 {
tmp4081 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4082 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp4081)
}
__typedArg0 := symlambda
__typedArg1 := tmp4081
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4066 Obj

if True == tmp4082 {
tmp4079 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4079)
}
__typedArg0 := tmp4079
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4068 Obj

if True == tmp4080 {
tmp4076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4077 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4076)
}
__typedArg0 := tmp4076
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4078 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4077)
}
__typedArg0 := tmp4077
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4070 Obj

if True == tmp4078 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4073)
}
__typedArg0 := tmp4073
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4074)
}
__typedArg0 := Nil
__typedArg1 := tmp4074
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4071 Obj

if True == tmp4075 {
ifres4071 = True


} else {
ifres4071 = False


}

ifres4070 = ifres4071


} else {
ifres4070 = False


}

var ifres4069 Obj

if True == ifres4070 {
ifres4069 = True


} else {
ifres4069 = False


}

ifres4068 = ifres4069


} else {
ifres4068 = False


}

var ifres4067 Obj

if True == ifres4068 {
ifres4067 = True


} else {
ifres4067 = False


}

ifres4066 = ifres4067


} else {
ifres4066 = False


}

var ifres4065 Obj

if True == ifres4066 {
ifres4065 = True


} else {
ifres4065 = False


}

ifres4064 = ifres4065


} else {
ifres4064 = False


}

if True == ifres4064 {
tmp4005 := Call(__e, PrimFunc(symgensym), symZ)


let__2717 := tmp4005
_ = let__2717

tmp4006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4006)
}
__typedArg0 := tmp4006
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4008)
}
__typedArg0 := tmp4008
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4009)
}
__typedArg0 := tmp4009
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4011 := Call(__e, PrimFunc(symshen_4beta), tmp4007, let__2717, tmp4010)


tmp4012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4011, Nil)
}
__typedArg0 := tmp4011
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2717, tmp4012)
}
__typedArg0 := let__2717
__typedArg1 := tmp4012
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp4013)
}
__typedArg0 := symlambda
__typedArg1 := tmp4013
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__2718 := tmp4014
_ = let__2718

tmp4015 := MakeNative(func(__e *ControlFlow) {
Z812 := __e.Get(1)
_ = Z812
__e.TailApply(PrimFunc(symshen_4alpha_1convert), Z812)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp4015, let__2718)
return


} else {
tmp4062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4036 Obj

if True == tmp4062 {
tmp4060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp4060)
}
__typedArg0 := symlet
__typedArg1 := tmp4060
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4038 Obj

if True == tmp4061 {
tmp4058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4058)
}
__typedArg0 := tmp4058
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4040 Obj

if True == tmp4059 {
tmp4055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4055)
}
__typedArg0 := tmp4055
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4056)
}
__typedArg0 := tmp4056
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4042 Obj

if True == tmp4057 {
tmp4051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4051)
}
__typedArg0 := tmp4051
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4052)
}
__typedArg0 := tmp4052
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4053)
}
__typedArg0 := tmp4053
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4044 Obj

if True == tmp4054 {
tmp4046 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4046)
}
__typedArg0 := tmp4046
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4047)
}
__typedArg0 := tmp4047
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4048)
}
__typedArg0 := tmp4048
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4049)
}
__typedArg0 := Nil
__typedArg1 := tmp4049
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4045 Obj

if True == tmp4050 {
ifres4045 = True


} else {
ifres4045 = False


}

ifres4044 = ifres4045


} else {
ifres4044 = False


}

var ifres4043 Obj

if True == ifres4044 {
ifres4043 = True


} else {
ifres4043 = False


}

ifres4042 = ifres4043


} else {
ifres4042 = False


}

var ifres4041 Obj

if True == ifres4042 {
ifres4041 = True


} else {
ifres4041 = False


}

ifres4040 = ifres4041


} else {
ifres4040 = False


}

var ifres4039 Obj

if True == ifres4040 {
ifres4039 = True


} else {
ifres4039 = False


}

ifres4038 = ifres4039


} else {
ifres4038 = False


}

var ifres4037 Obj

if True == ifres4038 {
ifres4037 = True


} else {
ifres4037 = False


}

ifres4036 = ifres4037


} else {
ifres4036 = False


}

if True == ifres4036 {
tmp4016 := Call(__e, PrimFunc(symgensym), symW)


let__2719 := tmp4016
_ = let__2719

tmp4017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4017)
}
__typedArg0 := tmp4017
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4018)
}
__typedArg0 := tmp4018
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4020)
}
__typedArg0 := tmp4020
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4022)
}
__typedArg0 := tmp4022
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4023)
}
__typedArg0 := tmp4023
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4024)
}
__typedArg0 := tmp4024
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4026 := Call(__e, PrimFunc(symshen_4beta), tmp4021, let__2719, tmp4025)


tmp4027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4026, Nil)
}
__typedArg0 := tmp4026
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4019, tmp4027)
}
__typedArg0 := tmp4019
__typedArg1 := tmp4027
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2719, tmp4028)
}
__typedArg0 := let__2719
__typedArg1 := tmp4028
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp4029)
}
__typedArg0 := symlet
__typedArg1 := tmp4029
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__2720 := tmp4030
_ = let__2720

tmp4031 := MakeNative(func(__e *ControlFlow) {
Z815 := __e.Get(1)
_ = Z815
__e.TailApply(PrimFunc(symshen_4alpha_1convert), Z815)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp4031, let__2720)
return


} else {
tmp4034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V809)
}
__typedArg0 := V809
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp4034 {
tmp4032 := MakeNative(func(__e *ControlFlow) {
Z816 := __e.Get(1)
_ = Z816
__e.TailApply(PrimFunc(symshen_4alpha_1convert), Z816)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp4032, V809)
return


} else {
__e.Return(V809)
return
}


}


}


}, 1)

tmp4084 := Call(__e, ns2_1set, symshen_4alpha_1convert, tmp4004)


_ = tmp4084

tmp4085 := MakeNative(func(__e *ControlFlow) {
V817 := __e.Get(1)
_ = V817
V818 := __e.Get(2)
_ = V818
tmp4086 := MakeNative(func(__e *ControlFlow) {
Z819 := __e.Get(1)
_ = Z819
tmp4087 := Call(__e, PrimFunc(symfst), Z819)


tmp4088 := Call(__e, PrimFunc(symsnd), Z819)


tmp4089 := Call(__e, PrimFunc(symshen_4alpha_1convert), tmp4088)


__e.TailApply(PrimFunc(symshen_4triple_1stack), Nil, tmp4087, V818, tmp4089)
return


}, 1)

__e.TailApply(PrimFunc(symmap), tmp4086, V817)
return


}, 2)

tmp4090 := Call(__e, ns2_1set, symshen_4kl_1body, tmp4085)


_ = tmp4090

tmp4091 := MakeNative(func(__e *ControlFlow) {
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
tmp4240 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V829)
}
__typedArg0 := Nil
__typedArg1 := V829
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4215 Obj

if True == tmp4240 {
tmp4239 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V830)
}
__typedArg0 := Nil
__typedArg1 := V830
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4217 Obj

if True == tmp4239 {
tmp4238 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V831)
}
__typedArg0 := V831
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4219 Obj

if True == tmp4238 {
tmp4236 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V831)
}
__typedArg0 := V831
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp4236)
}
__typedArg0 := symwhere
__typedArg1 := tmp4236
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4221 Obj

if True == tmp4237 {
tmp4234 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V831)
}
__typedArg0 := V831
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4235 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4234)
}
__typedArg0 := tmp4234
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4223 Obj

if True == tmp4235 {
tmp4231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V831)
}
__typedArg0 := V831
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4232 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4231)
}
__typedArg0 := tmp4231
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4232)
}
__typedArg0 := tmp4232
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4225 Obj

if True == tmp4233 {
tmp4227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V831)
}
__typedArg0 := V831
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4228 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4227)
}
__typedArg0 := tmp4227
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4228)
}
__typedArg0 := tmp4228
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4230 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4229)
}
__typedArg0 := Nil
__typedArg1 := tmp4229
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4226 Obj

if True == tmp4230 {
ifres4226 = True


} else {
ifres4226 = False


}

ifres4225 = ifres4226


} else {
ifres4225 = False


}

var ifres4224 Obj

if True == ifres4225 {
ifres4224 = True


} else {
ifres4224 = False


}

ifres4223 = ifres4224


} else {
ifres4223 = False


}

var ifres4222 Obj

if True == ifres4223 {
ifres4222 = True


} else {
ifres4222 = False


}

ifres4221 = ifres4222


} else {
ifres4221 = False


}

var ifres4220 Obj

if True == ifres4221 {
ifres4220 = True


} else {
ifres4220 = False


}

ifres4219 = ifres4220


} else {
ifres4219 = False


}

var ifres4218 Obj

if True == ifres4219 {
ifres4218 = True


} else {
ifres4218 = False


}

ifres4217 = ifres4218


} else {
ifres4217 = False


}

var ifres4216 Obj

if True == ifres4217 {
ifres4216 = True


} else {
ifres4216 = False


}

ifres4215 = ifres4216


} else {
ifres4215 = False


}

if True == ifres4215 {
tmp4092 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V831)
}
__typedArg0 := V831
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4092)
}
__typedArg0 := tmp4092
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4093, V828)
}
__typedArg0 := tmp4093
__typedArg1 := V828
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V831)
}
__typedArg0 := V831
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4095)
}
__typedArg0 := tmp4095
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4096)
}
__typedArg0 := tmp4096
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

if PrimFunc(symshen_4triple_1stack) == __self {
__self1, __self2, __self3, __self4 = tmp4094, Nil, Nil, tmp4097
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp4094, Nil, Nil, tmp4097)
return


} else {
tmp4213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V829)
}
__typedArg0 := Nil
__typedArg1 := V829
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4210 Obj

if True == tmp4213 {
tmp4212 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V830)
}
__typedArg0 := Nil
__typedArg1 := V830
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4211 Obj

if True == tmp4212 {
ifres4211 = True


} else {
ifres4211 = False


}

ifres4210 = ifres4211


} else {
ifres4210 = False


}

if True == ifres4210 {
tmp4098 := Call(__e, PrimFunc(symreverse), V828)


tmp4099 := Call(__e, PrimFunc(symshen_4rectify_1test), tmp4098)


tmp4100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V831, Nil)
}
__typedArg0 := V831
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4099, tmp4100)
}
__typedArg0 := tmp4099
__typedArg1 := tmp4100
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp4208 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4201 Obj

if True == tmp4208 {
tmp4207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4203 Obj

if True == tmp4207 {
tmp4205 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4206 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp4205)
}
__typedArg0 := tmp4205
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres4204 Obj

if True == tmp4206 {
ifres4204 = True


} else {
ifres4204 = False


}

ifres4203 = ifres4204


} else {
ifres4203 = False


}

var ifres4202 Obj

if True == ifres4203 {
ifres4202 = True


} else {
ifres4202 = False


}

ifres4201 = ifres4202


} else {
ifres4201 = False


}

if True == ifres4201 {
tmp4101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4104 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4105 := Call(__e, PrimFunc(symshen_4beta), tmp4103, tmp4104, V831)


if PrimFunc(symshen_4triple_1stack) == __self {
__self1, __self2, __self3, __self4 = V828, tmp4101, tmp4102, tmp4105
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4triple_1stack), V828, tmp4101, tmp4102, tmp4105)
return


} else {
tmp4199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4192 Obj

if True == tmp4199 {
tmp4198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4194 Obj

if True == tmp4198 {
tmp4196 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4197 := Call(__e, PrimFunc(symshen_4custom_1pattern_2), tmp4196)


var ifres4195 Obj

if True == tmp4197 {
ifres4195 = True


} else {
ifres4195 = False


}

ifres4194 = ifres4195


} else {
ifres4194 = False


}

var ifres4193 Obj

if True == ifres4194 {
ifres4193 = True


} else {
ifres4193 = False


}

ifres4192 = ifres4193


} else {
ifres4192 = False


}

if True == ifres4192 {
tmp4106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4107 := Call(__e, PrimFunc(symshen_4custom_1pattern_1body), tmp4106)


tmp4108 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4custom_1pattern_1triple_1stack), V828, tmp4107, tmp4108, tmp4109, tmp4110, V831)
return


} else {
tmp4190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4160 Obj

if True == tmp4190 {
tmp4188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4188)
}
__typedArg0 := tmp4188
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4162 Obj

if True == tmp4189 {
tmp4185 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4186 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4185)
}
__typedArg0 := tmp4185
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4186)
}
__typedArg0 := tmp4186
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4164 Obj

if True == tmp4187 {
tmp4181 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4182 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4181)
}
__typedArg0 := tmp4181
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4183 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4182)
}
__typedArg0 := tmp4182
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4184 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4183)
}
__typedArg0 := tmp4183
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4166 Obj

if True == tmp4184 {
tmp4176 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4177 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4176)
}
__typedArg0 := tmp4176
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4178 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4177)
}
__typedArg0 := tmp4177
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4179 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4178)
}
__typedArg0 := tmp4178
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4180 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4179)
}
__typedArg0 := Nil
__typedArg1 := tmp4179
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4168 Obj

if True == tmp4180 {
tmp4175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4170 Obj

if True == tmp4175 {
tmp4172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4172)
}
__typedArg0 := tmp4172
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4174 := Call(__e, PrimFunc(symshen_4constructor_2), tmp4173)


var ifres4171 Obj

if True == tmp4174 {
ifres4171 = True


} else {
ifres4171 = False


}

ifres4170 = ifres4171


} else {
ifres4170 = False


}

var ifres4169 Obj

if True == ifres4170 {
ifres4169 = True


} else {
ifres4169 = False


}

ifres4168 = ifres4169


} else {
ifres4168 = False


}

var ifres4167 Obj

if True == ifres4168 {
ifres4167 = True


} else {
ifres4167 = False


}

ifres4166 = ifres4167


} else {
ifres4166 = False


}

var ifres4165 Obj

if True == ifres4166 {
ifres4165 = True


} else {
ifres4165 = False


}

ifres4164 = ifres4165


} else {
ifres4164 = False


}

var ifres4163 Obj

if True == ifres4164 {
ifres4163 = True


} else {
ifres4163 = False


}

ifres4162 = ifres4163


} else {
ifres4162 = False


}

var ifres4161 Obj

if True == ifres4162 {
ifres4161 = True


} else {
ifres4161 = False


}

ifres4160 = ifres4161


} else {
ifres4160 = False


}

if True == ifres4160 {
tmp4111 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4112 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4111)
}
__typedArg0 := tmp4111
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4113 := Call(__e, PrimFunc(symshen_4op_1test), tmp4112)


tmp4114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4114, Nil)
}
__typedArg0 := tmp4114
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4113, tmp4115)
}
__typedArg0 := tmp4113
__typedArg1 := tmp4115
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4117 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4116, V828)
}
__typedArg0 := tmp4116
__typedArg1 := V828
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4118 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4119 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4118)
}
__typedArg0 := tmp4118
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4119)
}
__typedArg0 := tmp4119
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4121)
}
__typedArg0 := tmp4121
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4122)
}
__typedArg0 := tmp4122
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4123)
}
__typedArg0 := tmp4123
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4125 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4126 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4124, tmp4125)
}
__typedArg0 := tmp4124
__typedArg1 := tmp4125
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4120, tmp4126)
}
__typedArg0 := tmp4120
__typedArg1 := tmp4126
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

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

tmp4130 := Call(__e, PrimFunc(symshen_4op1), tmp4129)


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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4134)
}
__typedArg0 := tmp4134
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4136 := Call(__e, PrimFunc(symshen_4op2), tmp4135)


tmp4137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4137, Nil)
}
__typedArg0 := tmp4137
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4136, tmp4138)
}
__typedArg0 := tmp4136
__typedArg1 := tmp4138
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4140 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4141 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4139, tmp4140)
}
__typedArg0 := tmp4139
__typedArg1 := tmp4140
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4133, tmp4141)
}
__typedArg0 := tmp4133
__typedArg1 := tmp4141
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4143 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4144 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4145 := Call(__e, PrimFunc(symshen_4beta), tmp4143, tmp4144, V831)


if PrimFunc(symshen_4triple_1stack) == __self {
__self1, __self2, __self3, __self4 = tmp4117, tmp4127, tmp4142, tmp4145
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp4117, tmp4127, tmp4142, tmp4145)
return


} else {
tmp4158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4155 Obj

if True == tmp4158 {
tmp4157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4156 Obj

if True == tmp4157 {
ifres4156 = True


} else {
ifres4156 = False


}

ifres4155 = ifres4156


} else {
ifres4155 = False


}

if True == ifres4155 {
tmp4146 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4147 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4148 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4147, Nil)
}
__typedArg0 := tmp4147
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4149 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4146, tmp4148)
}
__typedArg0 := tmp4146
__typedArg1 := tmp4148
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4150 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp4149)
}
__typedArg0 := sym_a
__typedArg1 := tmp4149
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4151 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4150, V828)
}
__typedArg0 := tmp4150
__typedArg1 := V828
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4152 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V829)
}
__typedArg0 := V829
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4153 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V830)
}
__typedArg0 := V830
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4triple_1stack) == __self {
__self1, __self2, __self3, __self4 = tmp4151, tmp4152, tmp4153, V831
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp4151, tmp4152, tmp4153, V831)
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

tmp4241 := Call(__e, ns2_1set, symshen_4triple_1stack, tmp4091)


_ = tmp4241

tmp4242 := MakeNative(func(__e *ControlFlow) {
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
tmp4243 := Call(__e, PrimFunc(sym_8p), V833, V835)


tmp4244 := Call(__e, PrimFunc(symshen_4custom_1pattern_1reducer), tmp4243)


let__2721 := tmp4244
_ = let__2721

tmp4257 := Call(__e, PrimFunc(symtuple_2), let__2721)


if True == tmp4257 {
tmp4245 := Call(__e, PrimFunc(symfst), let__2721)


let__2722 := tmp4245
_ = let__2722

tmp4246 := Call(__e, PrimFunc(symsnd), let__2721)


let__2723 := tmp4246
_ = let__2723

tmp4247 := Call(__e, PrimFunc(symreverse), let__2722)


tmp4248 := Call(__e, PrimFunc(symappend), tmp4247, V832)


tmp4249 := MakeNative(func(__e *ControlFlow) {
Z841 := __e.Get(1)
_ = Z841
__e.TailApply(PrimFunc(symfst), Z841)
return
}, 1)

tmp4250 := Call(__e, PrimFunc(symmap), tmp4249, let__2723)


tmp4251 := Call(__e, PrimFunc(symappend), tmp4250, V834)


tmp4252 := MakeNative(func(__e *ControlFlow) {
Z842 := __e.Get(1)
_ = Z842
__e.TailApply(PrimFunc(symsnd), Z842)
return
}, 1)

tmp4253 := Call(__e, PrimFunc(symmap), tmp4252, let__2723)


tmp4254 := Call(__e, PrimFunc(symappend), tmp4253, V836)


tmp4255 := Call(__e, PrimFunc(symshen_4beta), V833, V835, V837)


__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp4248, tmp4251, tmp4254, tmp4255)
return


} else {
__e.TailApply(PrimFunc(symshen_4constructor_1error), V833)
return
}


}, 6)

tmp4258 := Call(__e, ns2_1set, symshen_4custom_1pattern_1triple_1stack, tmp4242)


_ = tmp4258

tmp4259 := MakeNative(func(__e *ControlFlow) {
V845 := __e.Get(1)
_ = V845
tmp4278 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V845)
}
__typedArg0 := Nil
__typedArg1 := V845
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4278 {
__e.Return(True)
return
} else {
tmp4276 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V845)
}
__typedArg0 := V845
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4272 Obj

if True == tmp4276 {
tmp4274 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V845)
}
__typedArg0 := V845
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4275 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4274)
}
__typedArg0 := Nil
__typedArg1 := tmp4274
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4273 Obj

if True == tmp4275 {
ifres4273 = True


} else {
ifres4273 = False


}

ifres4272 = ifres4273


} else {
ifres4272 = False


}

if True == ifres4272 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V845)
}
__typedArg0 := V845
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp4270 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V845)
}
__typedArg0 := V845
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4266 Obj

if True == tmp4270 {
tmp4268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V845)
}
__typedArg0 := V845
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4269 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4268)
}
__typedArg0 := tmp4268
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4267 Obj

if True == tmp4269 {
ifres4267 = True


} else {
ifres4267 = False


}

ifres4266 = ifres4267


} else {
ifres4266 = False


}

if True == ifres4266 {
tmp4260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V845)
}
__typedArg0 := V845
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V845)
}
__typedArg0 := V845
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4262 := Call(__e, PrimFunc(symshen_4rectify_1test), tmp4261)


tmp4263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4262, Nil)
}
__typedArg0 := tmp4262
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4260, tmp4263)
}
__typedArg0 := tmp4260
__typedArg1 := tmp4263
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp4264)
}
__typedArg0 := symand
__typedArg1 := tmp4264
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

tmp4279 := Call(__e, ns2_1set, symshen_4rectify_1test, tmp4259)


_ = tmp4279

tmp4280 := MakeNative(func(__e *ControlFlow) {
V855 := __e.Get(1)
_ = V855
V856 := __e.Get(2)
_ = V856
V857 := __e.Get(3)
_ = V857
tmp4357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V855, V857)
}
__typedArg0 := V855
__typedArg1 := V857
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4357 {
__e.Return(V856)
return
} else {
tmp4355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4331 Obj

if True == tmp4355 {
tmp4353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp4353)
}
__typedArg0 := symlambda
__typedArg1 := tmp4353
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4333 Obj

if True == tmp4354 {
tmp4351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4351)
}
__typedArg0 := tmp4351
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4335 Obj

if True == tmp4352 {
tmp4348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4348)
}
__typedArg0 := tmp4348
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4349)
}
__typedArg0 := tmp4349
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4337 Obj

if True == tmp4350 {
tmp4344 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4345)
}
__typedArg0 := tmp4345
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4346)
}
__typedArg0 := Nil
__typedArg1 := tmp4346
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4339 Obj

if True == tmp4347 {
tmp4341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4341)
}
__typedArg0 := tmp4341
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V855, tmp4342)
}
__typedArg0 := V855
__typedArg1 := tmp4342
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4340 Obj

if True == tmp4343 {
ifres4340 = True


} else {
ifres4340 = False


}

ifres4339 = ifres4340


} else {
ifres4339 = False


}

var ifres4338 Obj

if True == ifres4339 {
ifres4338 = True


} else {
ifres4338 = False


}

ifres4337 = ifres4338


} else {
ifres4337 = False


}

var ifres4336 Obj

if True == ifres4337 {
ifres4336 = True


} else {
ifres4336 = False


}

ifres4335 = ifres4336


} else {
ifres4335 = False


}

var ifres4334 Obj

if True == ifres4335 {
ifres4334 = True


} else {
ifres4334 = False


}

ifres4333 = ifres4334


} else {
ifres4333 = False


}

var ifres4332 Obj

if True == ifres4333 {
ifres4332 = True


} else {
ifres4332 = False


}

ifres4331 = ifres4332


} else {
ifres4331 = False


}

if True == ifres4331 {
__e.Return(V857)
return
} else {
tmp4329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4298 Obj

if True == tmp4329 {
tmp4327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4328 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp4327)
}
__typedArg0 := symlet
__typedArg1 := tmp4327
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4300 Obj

if True == tmp4328 {
tmp4325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4326 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4325)
}
__typedArg0 := tmp4325
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4302 Obj

if True == tmp4326 {
tmp4322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4322)
}
__typedArg0 := tmp4322
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4324 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4323)
}
__typedArg0 := tmp4323
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4304 Obj

if True == tmp4324 {
tmp4318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4318)
}
__typedArg0 := tmp4318
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4319)
}
__typedArg0 := tmp4319
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4320)
}
__typedArg0 := tmp4320
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4306 Obj

if True == tmp4321 {
tmp4313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4313)
}
__typedArg0 := tmp4313
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4316)
}
__typedArg0 := Nil
__typedArg1 := tmp4316
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4308 Obj

if True == tmp4317 {
tmp4310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4310)
}
__typedArg0 := tmp4310
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4312 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V855, tmp4311)
}
__typedArg0 := V855
__typedArg1 := tmp4311
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4309 Obj

if True == tmp4312 {
ifres4309 = True


} else {
ifres4309 = False


}

ifres4308 = ifres4309


} else {
ifres4308 = False


}

var ifres4307 Obj

if True == ifres4308 {
ifres4307 = True


} else {
ifres4307 = False


}

ifres4306 = ifres4307


} else {
ifres4306 = False


}

var ifres4305 Obj

if True == ifres4306 {
ifres4305 = True


} else {
ifres4305 = False


}

ifres4304 = ifres4305


} else {
ifres4304 = False


}

var ifres4303 Obj

if True == ifres4304 {
ifres4303 = True


} else {
ifres4303 = False


}

ifres4302 = ifres4303


} else {
ifres4302 = False


}

var ifres4301 Obj

if True == ifres4302 {
ifres4301 = True


} else {
ifres4301 = False


}

ifres4300 = ifres4301


} else {
ifres4300 = False


}

var ifres4299 Obj

if True == ifres4300 {
ifres4299 = True


} else {
ifres4299 = False


}

ifres4298 = ifres4299


} else {
ifres4298 = False


}

if True == ifres4298 {
tmp4281 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4282 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4281)
}
__typedArg0 := tmp4281
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4283 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4283)
}
__typedArg0 := tmp4283
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4286 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4285)
}
__typedArg0 := tmp4285
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4287 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4286)
}
__typedArg0 := tmp4286
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4288 := Call(__e, PrimFunc(symshen_4beta), tmp4284, V856, tmp4287)


tmp4289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4290 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4289)
}
__typedArg0 := tmp4289
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4291 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4290)
}
__typedArg0 := tmp4290
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4292 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4288, tmp4291)
}
__typedArg0 := tmp4288
__typedArg1 := tmp4291
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4282, tmp4292)
}
__typedArg0 := tmp4282
__typedArg1 := tmp4292
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp4293)
}
__typedArg0 := symlet
__typedArg1 := tmp4293
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp4296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V857)
}
__typedArg0 := V857
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp4296 {
tmp4294 := MakeNative(func(__e *ControlFlow) {
Z858 := __e.Get(1)
_ = Z858
__e.TailApply(PrimFunc(symshen_4beta), V855, V856, Z858)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp4294, V857)
return


} else {
__e.Return(V857)
return
}


}


}


}


}, 3)

tmp4358 := Call(__e, ns2_1set, symshen_4beta, tmp4280)


_ = tmp4358

tmp4359 := MakeNative(func(__e *ControlFlow) {
V861 := __e.Get(1)
_ = V861
tmp4367 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, V861)
}
__typedArg0 := symcons
__typedArg1 := V861
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4367 {
__e.Return(symhd)
return
} else {
tmp4365 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8s, V861)
}
__typedArg0 := sym_8s
__typedArg1 := V861
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4365 {
__e.Return(symhdstr)
return
} else {
tmp4363 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8p, V861)
}
__typedArg0 := sym_8p
__typedArg1 := V861
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4363 {
__e.Return(symfst)
return
} else {
tmp4361 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8v, V861)
}
__typedArg0 := sym_8v
__typedArg1 := V861
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4361 {
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

tmp4368 := Call(__e, ns2_1set, symshen_4op1, tmp4359)


_ = tmp4368

tmp4369 := MakeNative(func(__e *ControlFlow) {
V864 := __e.Get(1)
_ = V864
tmp4377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, V864)
}
__typedArg0 := symcons
__typedArg1 := V864
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4377 {
__e.Return(symtl)
return
} else {
tmp4375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8s, V864)
}
__typedArg0 := sym_8s
__typedArg1 := V864
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4375 {
__e.Return(symtlstr)
return
} else {
tmp4373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8p, V864)
}
__typedArg0 := sym_8p
__typedArg1 := V864
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4373 {
__e.Return(symsnd)
return
} else {
tmp4371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8v, V864)
}
__typedArg0 := sym_8v
__typedArg1 := V864
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4371 {
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

tmp4378 := Call(__e, ns2_1set, symshen_4op2, tmp4369)


_ = tmp4378

tmp4379 := MakeNative(func(__e *ControlFlow) {
V867 := __e.Get(1)
_ = V867
tmp4387 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, V867)
}
__typedArg0 := symcons
__typedArg1 := V867
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4387 {
__e.Return(symcons_2)
return
} else {
tmp4385 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8s, V867)
}
__typedArg0 := sym_8s
__typedArg1 := V867
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4385 {
__e.Return(symshen_4_7string_2)
return
} else {
tmp4383 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8p, V867)
}
__typedArg0 := sym_8p
__typedArg1 := V867
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4383 {
__e.Return(symtuple_2)
return
} else {
tmp4381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8v, V867)
}
__typedArg0 := sym_8v
__typedArg1 := V867
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4381 {
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

tmp4388 := Call(__e, ns2_1set, symshen_4op_1test, tmp4379)


_ = tmp4388

tmp4389 := MakeNative(func(__e *ControlFlow) {
V868 := __e.Get(1)
_ = V868
tmp4391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V868)
}
__typedArg0 := MakeString("")
__typedArg1 := V868
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4391 {
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

tmp4392 := Call(__e, ns2_1set, symshen_4_7string_2, tmp4389)


_ = tmp4392

tmp4393 := MakeNative(func(__e *ControlFlow) {
V869 := __e.Get(1)
_ = V869
tmp4395 := Call(__e, PrimFunc(symvector), MakeInteger(0))


tmp4396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V869, tmp4395)
}
__typedArg0 := V869
__typedArg1 := tmp4395
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4396 {
__e.Return(False)
return
} else {
__e.TailApply(PrimFunc(symvector_2), V869)
return
}


}, 1)

tmp4397 := Call(__e, ns2_1set, symshen_4_7vector_2, tmp4393)


_ = tmp4397

tmp4398 := MakeNative(func(__e *ControlFlow) {
V872 := __e.Get(1)
_ = V872
tmp4402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V872)
}
__typedArg0 := sym_7
__typedArg1 := V872
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4402 {
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
tmp4400 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V872)
}
__typedArg0 := sym_1
__typedArg1 := V872
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4400 {
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

tmp4403 := Call(__e, ns2_1set, symfactorise, tmp4398)


_ = tmp4403

tmp4404 := MakeNative(func(__e *ControlFlow) {
V873 := __e.Get(1)
_ = V873
tmp4406 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dfactorise_2_d)
}
__typedArg0 := symshen_4_dfactorise_2_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

if True == tmp4406 {
__e.TailApply(PrimFunc(symshen_4factor), V873)
return
} else {
__e.Return(V873)
return
}


}, 1)

tmp4407 := Call(__e, ns2_1set, symshen_4factorise_1code, tmp4404)


_ = tmp4407

tmp4408 := MakeNative(func(__e *ControlFlow) {
V874 := __e.Get(1)
_ = V874
tmp4465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V874)
}
__typedArg0 := V874
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4424 Obj

if True == tmp4465 {
tmp4463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V874)
}
__typedArg0 := V874
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp4463)
}
__typedArg0 := symdefun
__typedArg1 := tmp4463
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4426 Obj

if True == tmp4464 {
tmp4461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V874)
}
__typedArg0 := V874
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4461)
}
__typedArg0 := tmp4461
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4428 Obj

if True == tmp4462 {
tmp4458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V874)
}
__typedArg0 := V874
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4458)
}
__typedArg0 := tmp4458
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4459)
}
__typedArg0 := tmp4459
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4430 Obj

if True == tmp4460 {
tmp4454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V874)
}
__typedArg0 := V874
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4454)
}
__typedArg0 := tmp4454
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4456)
}
__typedArg0 := tmp4456
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4432 Obj

if True == tmp4457 {
tmp4449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V874)
}
__typedArg0 := V874
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4450 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4449)
}
__typedArg0 := tmp4449
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4450)
}
__typedArg0 := tmp4450
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4452 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4451)
}
__typedArg0 := tmp4451
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4453 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4452)
}
__typedArg0 := tmp4452
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4434 Obj

if True == tmp4453 {
tmp4443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V874)
}
__typedArg0 := V874
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4443)
}
__typedArg0 := tmp4443
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4444)
}
__typedArg0 := tmp4444
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4445)
}
__typedArg0 := tmp4445
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4446)
}
__typedArg0 := tmp4446
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcond, tmp4447)
}
__typedArg0 := symcond
__typedArg1 := tmp4447
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4436 Obj

if True == tmp4448 {
tmp4438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V874)
}
__typedArg0 := V874
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4439)
}
__typedArg0 := tmp4439
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4440)
}
__typedArg0 := tmp4440
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4441)
}
__typedArg0 := Nil
__typedArg1 := tmp4441
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4437 Obj

if True == tmp4442 {
ifres4437 = True


} else {
ifres4437 = False


}

ifres4436 = ifres4437


} else {
ifres4436 = False


}

var ifres4435 Obj

if True == ifres4436 {
ifres4435 = True


} else {
ifres4435 = False


}

ifres4434 = ifres4435


} else {
ifres4434 = False


}

var ifres4433 Obj

if True == ifres4434 {
ifres4433 = True


} else {
ifres4433 = False


}

ifres4432 = ifres4433


} else {
ifres4432 = False


}

var ifres4431 Obj

if True == ifres4432 {
ifres4431 = True


} else {
ifres4431 = False


}

ifres4430 = ifres4431


} else {
ifres4430 = False


}

var ifres4429 Obj

if True == ifres4430 {
ifres4429 = True


} else {
ifres4429 = False


}

ifres4428 = ifres4429


} else {
ifres4428 = False


}

var ifres4427 Obj

if True == ifres4428 {
ifres4427 = True


} else {
ifres4427 = False


}

ifres4426 = ifres4427


} else {
ifres4426 = False


}

var ifres4425 Obj

if True == ifres4426 {
ifres4425 = True


} else {
ifres4425 = False


}

ifres4424 = ifres4425


} else {
ifres4424 = False


}

if True == ifres4424 {
tmp4409 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V874)
}
__typedArg0 := V874
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4410 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4409)
}
__typedArg0 := tmp4409
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4411 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V874)
}
__typedArg0 := V874
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4412 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4411)
}
__typedArg0 := tmp4411
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4413 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4412)
}
__typedArg0 := tmp4412
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4414 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V874)
}
__typedArg0 := V874
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4415 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4414)
}
__typedArg0 := tmp4414
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4415)
}
__typedArg0 := tmp4415
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4416)
}
__typedArg0 := tmp4416
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4417)
}
__typedArg0 := tmp4417
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4419 := Call(__e, PrimFunc(symshen_4factor_1recognisors), tmp4418)


tmp4420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4419, Nil)
}
__typedArg0 := tmp4419
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4413, tmp4420)
}
__typedArg0 := tmp4413
__typedArg1 := tmp4420
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4410, tmp4421)
}
__typedArg0 := tmp4410
__typedArg1 := tmp4421
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp4422)
}
__typedArg0 := symdefun
__typedArg1 := tmp4422
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V874)
return
}


}, 1)

tmp4466 := Call(__e, ns2_1set, symshen_4factor, tmp4408)


_ = tmp4466

tmp4467 := MakeNative(func(__e *ControlFlow) {
V877 := __e.Get(1)
_ = V877
tmp4617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4597 Obj

if True == tmp4617 {
tmp4615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4615)
}
__typedArg0 := tmp4615
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4599 Obj

if True == tmp4616 {
tmp4612 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4612)
}
__typedArg0 := tmp4612
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, tmp4613)
}
__typedArg0 := True
__typedArg1 := tmp4613
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4601 Obj

if True == tmp4614 {
tmp4609 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4610 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4609)
}
__typedArg0 := tmp4609
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4611 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4610)
}
__typedArg0 := tmp4610
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4603 Obj

if True == tmp4611 {
tmp4605 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4606 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4605)
}
__typedArg0 := tmp4605
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4607 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4606)
}
__typedArg0 := tmp4606
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4608 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4607)
}
__typedArg0 := Nil
__typedArg1 := tmp4607
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4604 Obj

if True == tmp4608 {
ifres4604 = True


} else {
ifres4604 = False


}

ifres4603 = ifres4604


} else {
ifres4603 = False


}

var ifres4602 Obj

if True == ifres4603 {
ifres4602 = True


} else {
ifres4602 = False


}

ifres4601 = ifres4602


} else {
ifres4601 = False


}

var ifres4600 Obj

if True == ifres4601 {
ifres4600 = True


} else {
ifres4600 = False


}

ifres4599 = ifres4600


} else {
ifres4599 = False


}

var ifres4598 Obj

if True == ifres4599 {
ifres4598 = True


} else {
ifres4598 = False


}

ifres4597 = ifres4598


} else {
ifres4597 = False


}

if True == ifres4597 {
tmp4468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4468)
}
__typedArg0 := tmp4468
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4469)
}
__typedArg0 := tmp4469
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp4595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4548 Obj

if True == tmp4595 {
tmp4593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4593)
}
__typedArg0 := tmp4593
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4550 Obj

if True == tmp4594 {
tmp4590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4590)
}
__typedArg0 := tmp4590
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4591)
}
__typedArg0 := tmp4591
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4552 Obj

if True == tmp4592 {
tmp4586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4586)
}
__typedArg0 := tmp4586
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4587)
}
__typedArg0 := tmp4587
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symand, tmp4588)
}
__typedArg0 := symand
__typedArg1 := tmp4588
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4554 Obj

if True == tmp4589 {
tmp4582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4582)
}
__typedArg0 := tmp4582
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4583)
}
__typedArg0 := tmp4583
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4584)
}
__typedArg0 := tmp4584
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4556 Obj

if True == tmp4585 {
tmp4577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4577)
}
__typedArg0 := tmp4577
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4578)
}
__typedArg0 := tmp4578
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4579)
}
__typedArg0 := tmp4579
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4580)
}
__typedArg0 := tmp4580
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4558 Obj

if True == tmp4581 {
tmp4571 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4572 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4571)
}
__typedArg0 := tmp4571
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4573 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4572)
}
__typedArg0 := tmp4572
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4574 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4573)
}
__typedArg0 := tmp4573
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4574)
}
__typedArg0 := tmp4574
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4575)
}
__typedArg0 := Nil
__typedArg1 := tmp4575
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4560 Obj

if True == tmp4576 {
tmp4568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4568)
}
__typedArg0 := tmp4568
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4569)
}
__typedArg0 := tmp4569
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4562 Obj

if True == tmp4570 {
tmp4564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4564)
}
__typedArg0 := tmp4564
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4565)
}
__typedArg0 := tmp4565
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4566)
}
__typedArg0 := Nil
__typedArg1 := tmp4566
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4563 Obj

if True == tmp4567 {
ifres4563 = True


} else {
ifres4563 = False


}

ifres4562 = ifres4563


} else {
ifres4562 = False


}

var ifres4561 Obj

if True == ifres4562 {
ifres4561 = True


} else {
ifres4561 = False


}

ifres4560 = ifres4561


} else {
ifres4560 = False


}

var ifres4559 Obj

if True == ifres4560 {
ifres4559 = True


} else {
ifres4559 = False


}

ifres4558 = ifres4559


} else {
ifres4558 = False


}

var ifres4557 Obj

if True == ifres4558 {
ifres4557 = True


} else {
ifres4557 = False


}

ifres4556 = ifres4557


} else {
ifres4556 = False


}

var ifres4555 Obj

if True == ifres4556 {
ifres4555 = True


} else {
ifres4555 = False


}

ifres4554 = ifres4555


} else {
ifres4554 = False


}

var ifres4553 Obj

if True == ifres4554 {
ifres4553 = True


} else {
ifres4553 = False


}

ifres4552 = ifres4553


} else {
ifres4552 = False


}

var ifres4551 Obj

if True == ifres4552 {
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

if True == ifres4548 {
tmp4470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4471 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4470)
}
__typedArg0 := tmp4470
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4472 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4471)
}
__typedArg0 := tmp4471
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4472)
}
__typedArg0 := tmp4472
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4474 := Call(__e, PrimFunc(symshen_4pivot_1on), tmp4473, V877, Nil)


let__2724 := tmp4474
_ = let__2724

tmp4475 := Call(__e, PrimFunc(symfst), let__2724)


let__2725 := tmp4475
_ = let__2725

tmp4519 := Call(__e, PrimFunc(symshen_4bad_1pivot_2), let__2725)


if True == tmp4519 {
tmp4476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, Nil)
}
__typedArg0 := symif
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4479)
}
__typedArg0 := tmp4479
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4480)
}
__typedArg0 := tmp4480
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4482 := Call(__e, PrimFunc(symshen_4recursively_1factor_1selectors), tmp4478, tmp4481)


tmp4483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4484 := Call(__e, PrimFunc(symshen_4factor_1recognisors), tmp4483)


tmp4485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4484, Nil)
}
__typedArg0 := tmp4484
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4486 := Call(__e, PrimFunc(symappend), tmp4482, tmp4485)


__e.TailApply(PrimFunc(symappend), tmp4476, tmp4486)
return


} else {
tmp4487 := Call(__e, PrimFunc(symsnd), let__2724)


let__2726 := tmp4487
_ = let__2726

tmp4488 := Call(__e, PrimFunc(symshen_4factor_1recognisors), let__2726)


let__2727 := tmp4488
_ = let__2727

tmp4489 := Call(__e, PrimFunc(symgensym), symGoTo)


let__2728 := tmp4489
_ = let__2728

tmp4490 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2728, Nil)
}
__typedArg0 := let__2728
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp4490)
}
__typedArg0 := symthaw
__typedArg1 := tmp4490
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4491, Nil)
}
__typedArg0 := tmp4491
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4493 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp4492)
}
__typedArg0 := True
__typedArg1 := tmp4492
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4494 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4493, let__2725)
}
__typedArg0 := tmp4493
__typedArg1 := let__2725
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4495 := Call(__e, PrimFunc(symreverse), tmp4494)


let__2729 := tmp4495
_ = let__2729

tmp4496 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2727, Nil)
}
__typedArg0 := let__2727
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp4496)
}
__typedArg0 := symfreeze
__typedArg1 := tmp4496
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4498)
}
__typedArg0 := tmp4498
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4499)
}
__typedArg0 := tmp4499
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4500)
}
__typedArg0 := tmp4500
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4502)
}
__typedArg0 := tmp4502
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4503)
}
__typedArg0 := tmp4503
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4504)
}
__typedArg0 := tmp4504
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4506 := Call(__e, PrimFunc(symshen_4factor_1recognisors), let__2729)


tmp4507 := Call(__e, PrimFunc(symshen_4factor_1selectors), tmp4505, tmp4506)


tmp4508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2728, Nil)
}
__typedArg0 := let__2728
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp4508)
}
__typedArg0 := symthaw
__typedArg1 := tmp4508
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4509, Nil)
}
__typedArg0 := tmp4509
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4507, tmp4510)
}
__typedArg0 := tmp4507
__typedArg1 := tmp4510
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4501, tmp4511)
}
__typedArg0 := tmp4501
__typedArg1 := tmp4511
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp4512)
}
__typedArg0 := symif
__typedArg1 := tmp4512
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4513, Nil)
}
__typedArg0 := tmp4513
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4497, tmp4514)
}
__typedArg0 := tmp4497
__typedArg1 := tmp4514
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2728, tmp4515)
}
__typedArg0 := let__2728
__typedArg1 := tmp4515
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4517 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp4516)
}
__typedArg0 := symlet
__typedArg1 := tmp4516
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

let__2730 := tmp4517
_ = let__2730

__e.TailApply(PrimFunc(symshen_4remove_1indirection), let__2730)
return


}


} else {
tmp4546 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4531 Obj

if True == tmp4546 {
tmp4544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4545 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4544)
}
__typedArg0 := tmp4544
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4533 Obj

if True == tmp4545 {
tmp4541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4541)
}
__typedArg0 := tmp4541
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4542)
}
__typedArg0 := tmp4542
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4535 Obj

if True == tmp4543 {
tmp4537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4537)
}
__typedArg0 := tmp4537
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4538)
}
__typedArg0 := tmp4538
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4539)
}
__typedArg0 := Nil
__typedArg1 := tmp4539
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4536 Obj

if True == tmp4540 {
ifres4536 = True


} else {
ifres4536 = False


}

ifres4535 = ifres4536


} else {
ifres4535 = False


}

var ifres4534 Obj

if True == ifres4535 {
ifres4534 = True


} else {
ifres4534 = False


}

ifres4533 = ifres4534


} else {
ifres4533 = False


}

var ifres4532 Obj

if True == ifres4533 {
ifres4532 = True


} else {
ifres4532 = False


}

ifres4531 = ifres4532


} else {
ifres4531 = False


}

if True == ifres4531 {
tmp4520 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4520)
}
__typedArg0 := tmp4520
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symhd), __typedArg0)
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
return PrimTail(V877)
}
__typedArg0 := V877
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4526 := Call(__e, PrimFunc(symshen_4factor_1recognisors), tmp4525)


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
return PrimCons(tmp4521, tmp4528)
}
__typedArg0 := tmp4521
__typedArg1 := tmp4528
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp4529)
}
__typedArg0 := symif
__typedArg1 := tmp4529
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4factor_1recognisors)
return
}


}


}


}, 1)

tmp4618 := Call(__e, ns2_1set, symshen_4factor_1recognisors, tmp4467)


_ = tmp4618

tmp4619 := MakeNative(func(__e *ControlFlow) {
V885 := __e.Get(1)
_ = V885
V886 := __e.Get(2)
_ = V886
tmp4691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4665 Obj

if True == tmp4691 {
tmp4689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp4689)
}
__typedArg0 := symlet
__typedArg1 := tmp4689
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4667 Obj

if True == tmp4690 {
tmp4687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4687)
}
__typedArg0 := tmp4687
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4669 Obj

if True == tmp4688 {
tmp4684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4684)
}
__typedArg0 := tmp4684
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4685)
}
__typedArg0 := tmp4685
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4671 Obj

if True == tmp4686 {
tmp4680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4680)
}
__typedArg0 := tmp4680
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4682)
}
__typedArg0 := tmp4682
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4673 Obj

if True == tmp4683 {
tmp4675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4675)
}
__typedArg0 := tmp4675
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4677)
}
__typedArg0 := tmp4677
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4678)
}
__typedArg0 := Nil
__typedArg1 := tmp4678
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4674 Obj

if True == tmp4679 {
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

if True == ifres4665 {
tmp4620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4620)
}
__typedArg0 := tmp4620
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4622)
}
__typedArg0 := tmp4622
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4623)
}
__typedArg0 := tmp4623
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4627)
}
__typedArg0 := tmp4627
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4629 := Call(__e, PrimFunc(symshen_4recursively_1factor_1selectors), tmp4628, V886)


__e.TailApply(PrimFunc(symshen_4restore_1local), tmp4621, tmp4624, tmp4629)
return


} else {
tmp4663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4644 Obj

if True == tmp4663 {
tmp4661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symand, tmp4661)
}
__typedArg0 := symand
__typedArg1 := tmp4661
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4646 Obj

if True == tmp4662 {
tmp4659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4659)
}
__typedArg0 := tmp4659
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4648 Obj

if True == tmp4660 {
tmp4656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4656)
}
__typedArg0 := tmp4656
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4657)
}
__typedArg0 := tmp4657
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4650 Obj

if True == tmp4658 {
tmp4652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4652)
}
__typedArg0 := tmp4652
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4653)
}
__typedArg0 := tmp4653
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4654)
}
__typedArg0 := Nil
__typedArg1 := tmp4654
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4651 Obj

if True == tmp4655 {
ifres4651 = True


} else {
ifres4651 = False


}

ifres4650 = ifres4651


} else {
ifres4650 = False


}

var ifres4649 Obj

if True == ifres4650 {
ifres4649 = True


} else {
ifres4649 = False


}

ifres4648 = ifres4649


} else {
ifres4648 = False


}

var ifres4647 Obj

if True == ifres4648 {
ifres4647 = True


} else {
ifres4647 = False


}

ifres4646 = ifres4647


} else {
ifres4646 = False


}

var ifres4645 Obj

if True == ifres4646 {
ifres4645 = True


} else {
ifres4645 = False


}

ifres4644 = ifres4645


} else {
ifres4644 = False


}

if True == ifres4644 {
tmp4630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4630)
}
__typedArg0 := tmp4630
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4632 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4632)
}
__typedArg0 := tmp4632
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4634)
}
__typedArg0 := tmp4634
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4636 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4635)
}
__typedArg0 := tmp4635
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4637 := Call(__e, PrimFunc(symshen_4factor_1selectors), tmp4633, tmp4636)


tmp4638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V885)
}
__typedArg0 := V885
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4638)
}
__typedArg0 := tmp4638
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4640 := Call(__e, PrimFunc(symshen_4factor_1selectors), tmp4639, V886)


tmp4641 := Call(__e, PrimFunc(symshen_4recursively_1factor_1selectors), tmp4637, tmp4640)


__e.TailApply(PrimFunc(symshen_4restore_1P), tmp4631, tmp4641)
return


} else {
tmp4642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V886, Nil)
}
__typedArg0 := V886
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V885, tmp4642)
}
__typedArg0 := V885
__typedArg1 := tmp4642
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}


}, 2)

tmp4692 := Call(__e, ns2_1set, symshen_4recursively_1factor_1selectors, tmp4619)


_ = tmp4692

tmp4693 := MakeNative(func(__e *ControlFlow) {
V887 := __e.Get(1)
_ = V887
V888 := __e.Get(2)
_ = V888
tmp4709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V888)
}
__typedArg0 := V888
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4700 Obj

if True == tmp4709 {
tmp4707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V888)
}
__typedArg0 := V888
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4707)
}
__typedArg0 := tmp4707
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4702 Obj

if True == tmp4708 {
tmp4704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V888)
}
__typedArg0 := V888
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4705)
}
__typedArg0 := Nil
__typedArg1 := tmp4705
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4703 Obj

if True == tmp4706 {
ifres4703 = True


} else {
ifres4703 = False


}

ifres4702 = ifres4703


} else {
ifres4702 = False


}

var ifres4701 Obj

if True == ifres4702 {
ifres4701 = True


} else {
ifres4701 = False


}

ifres4700 = ifres4701


} else {
ifres4700 = False


}

if True == ifres4700 {
tmp4694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V888)
}
__typedArg0 := V888
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4694, Nil)
}
__typedArg0 := tmp4694
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V887, tmp4695)
}
__typedArg0 := V887
__typedArg1 := tmp4695
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp4696)
}
__typedArg0 := symand
__typedArg1 := tmp4696
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V888)
}
__typedArg0 := V888
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4697, tmp4698)
}
__typedArg0 := tmp4697
__typedArg1 := tmp4698
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4restore_1P)
return
}


}, 2)

tmp4710 := Call(__e, ns2_1set, symshen_4restore_1P, tmp4693)


_ = tmp4710

tmp4711 := MakeNative(func(__e *ControlFlow) {
V889 := __e.Get(1)
_ = V889
V890 := __e.Get(2)
_ = V890
V891 := __e.Get(3)
_ = V891
tmp4728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V891)
}
__typedArg0 := V891
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4719 Obj

if True == tmp4728 {
tmp4726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V891)
}
__typedArg0 := V891
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4726)
}
__typedArg0 := tmp4726
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4721 Obj

if True == tmp4727 {
tmp4723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V891)
}
__typedArg0 := V891
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4723)
}
__typedArg0 := tmp4723
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4724)
}
__typedArg0 := Nil
__typedArg1 := tmp4724
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4722 Obj

if True == tmp4725 {
ifres4722 = True


} else {
ifres4722 = False


}

ifres4721 = ifres4722


} else {
ifres4721 = False


}

var ifres4720 Obj

if True == ifres4721 {
ifres4720 = True


} else {
ifres4720 = False


}

ifres4719 = ifres4720


} else {
ifres4719 = False


}

if True == ifres4719 {
tmp4712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V891)
}
__typedArg0 := V891
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4712, Nil)
}
__typedArg0 := tmp4712
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V890, tmp4713)
}
__typedArg0 := V890
__typedArg1 := tmp4713
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V889, tmp4714)
}
__typedArg0 := V889
__typedArg1 := tmp4714
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp4715)
}
__typedArg0 := symlet
__typedArg1 := tmp4715
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V891)
}
__typedArg0 := V891
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4716, tmp4717)
}
__typedArg0 := tmp4716
__typedArg1 := tmp4717
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4restore_1local)
return
}


}, 3)

tmp4729 := Call(__e, ns2_1set, symshen_4restore_1local, tmp4711)


_ = tmp4729

tmp4730 := MakeNative(func(__e *ControlFlow) {
V896 := __e.Get(1)
_ = V896
tmp4736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V896)
}
__typedArg0 := V896
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4732 Obj

if True == tmp4736 {
tmp4734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V896)
}
__typedArg0 := V896
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4734)
}
__typedArg0 := Nil
__typedArg1 := tmp4734
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4733 Obj

if True == tmp4735 {
ifres4733 = True


} else {
ifres4733 = False


}

ifres4732 = ifres4733


} else {
ifres4732 = False


}

if True == ifres4732 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp4737 := Call(__e, ns2_1set, symshen_4bad_1pivot_2, tmp4730)


_ = tmp4737

tmp4738 := MakeNative(func(__e *ControlFlow) {
V897 := __e.Get(1)
_ = V897
tmp4853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4753 Obj

if True == tmp4853 {
tmp4851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp4851)
}
__typedArg0 := symlet
__typedArg1 := tmp4851
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4755 Obj

if True == tmp4852 {
tmp4849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4849)
}
__typedArg0 := tmp4849
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4757 Obj

if True == tmp4850 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4847)
}
__typedArg0 := tmp4847
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4759 Obj

if True == tmp4848 {
tmp4842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4842)
}
__typedArg0 := tmp4842
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4843)
}
__typedArg0 := tmp4843
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4844)
}
__typedArg0 := tmp4844
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4761 Obj

if True == tmp4845 {
tmp4837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4838 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4837)
}
__typedArg0 := tmp4837
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4838)
}
__typedArg0 := tmp4838
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4840 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4839)
}
__typedArg0 := tmp4839
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfreeze, tmp4840)
}
__typedArg0 := symfreeze
__typedArg1 := tmp4840
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4763 Obj

if True == tmp4841 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4835)
}
__typedArg0 := tmp4835
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4765 Obj

if True == tmp4836 {
tmp4826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4829)
}
__typedArg0 := tmp4829
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4830)
}
__typedArg0 := tmp4830
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4767 Obj

if True == tmp4831 {
tmp4819 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4821)
}
__typedArg0 := tmp4821
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4822)
}
__typedArg0 := tmp4822
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4823)
}
__typedArg0 := tmp4823
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symthaw, tmp4824)
}
__typedArg0 := symthaw
__typedArg1 := tmp4824
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4769 Obj

if True == tmp4825 {
tmp4812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4812)
}
__typedArg0 := tmp4812
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4814 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4813)
}
__typedArg0 := tmp4813
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4814)
}
__typedArg0 := tmp4814
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4815)
}
__typedArg0 := tmp4815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4816)
}
__typedArg0 := tmp4816
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4817)
}
__typedArg0 := tmp4817
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4771 Obj

if True == tmp4818 {
tmp4804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4804)
}
__typedArg0 := tmp4804
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4805)
}
__typedArg0 := tmp4805
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4806)
}
__typedArg0 := tmp4806
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4807)
}
__typedArg0 := tmp4807
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4808)
}
__typedArg0 := tmp4808
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4810)
}
__typedArg0 := Nil
__typedArg1 := tmp4810
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4773 Obj

if True == tmp4811 {
tmp4798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4798)
}
__typedArg0 := tmp4798
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4799)
}
__typedArg0 := tmp4799
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4800)
}
__typedArg0 := tmp4800
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4802)
}
__typedArg0 := Nil
__typedArg1 := tmp4802
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4775 Obj

if True == tmp4803 {
tmp4794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4794)
}
__typedArg0 := tmp4794
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4796 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4795)
}
__typedArg0 := tmp4795
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4797 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4796)
}
__typedArg0 := tmp4796
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4777 Obj

if True == tmp4797 {
tmp4789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4791)
}
__typedArg0 := tmp4791
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4793 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4792)
}
__typedArg0 := Nil
__typedArg1 := tmp4792
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4779 Obj

if True == tmp4793 {
tmp4781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4781)
}
__typedArg0 := tmp4781
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4782)
}
__typedArg0 := tmp4782
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4783)
}
__typedArg0 := tmp4783
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4784)
}
__typedArg0 := tmp4784
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4785)
}
__typedArg0 := tmp4785
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4786)
}
__typedArg0 := tmp4786
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(tmp4787)
}
__typedArg0 := tmp4787
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

var ifres4780 Obj

if True == tmp4788 {
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

var ifres4772 Obj

if True == ifres4773 {
ifres4772 = True


} else {
ifres4772 = False


}

ifres4771 = ifres4772


} else {
ifres4771 = False


}

var ifres4770 Obj

if True == ifres4771 {
ifres4770 = True


} else {
ifres4770 = False


}

ifres4769 = ifres4770


} else {
ifres4769 = False


}

var ifres4768 Obj

if True == ifres4769 {
ifres4768 = True


} else {
ifres4768 = False


}

ifres4767 = ifres4768


} else {
ifres4767 = False


}

var ifres4766 Obj

if True == ifres4767 {
ifres4766 = True


} else {
ifres4766 = False


}

ifres4765 = ifres4766


} else {
ifres4765 = False


}

var ifres4764 Obj

if True == ifres4765 {
ifres4764 = True


} else {
ifres4764 = False


}

ifres4763 = ifres4764


} else {
ifres4763 = False


}

var ifres4762 Obj

if True == ifres4763 {
ifres4762 = True


} else {
ifres4762 = False


}

ifres4761 = ifres4762


} else {
ifres4761 = False


}

var ifres4760 Obj

if True == ifres4761 {
ifres4760 = True


} else {
ifres4760 = False


}

ifres4759 = ifres4760


} else {
ifres4759 = False


}

var ifres4758 Obj

if True == ifres4759 {
ifres4758 = True


} else {
ifres4758 = False


}

ifres4757 = ifres4758


} else {
ifres4757 = False


}

var ifres4756 Obj

if True == ifres4757 {
ifres4756 = True


} else {
ifres4756 = False


}

ifres4755 = ifres4756


} else {
ifres4755 = False


}

var ifres4754 Obj

if True == ifres4755 {
ifres4754 = True


} else {
ifres4754 = False


}

ifres4753 = ifres4754


} else {
ifres4753 = False


}

if True == ifres4753 {
tmp4739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4740 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4739)
}
__typedArg0 := tmp4739
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4741 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4740)
}
__typedArg0 := tmp4740
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4741)
}
__typedArg0 := tmp4741
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4742)
}
__typedArg0 := tmp4742
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4743)
}
__typedArg0 := tmp4743
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4744)
}
__typedArg0 := tmp4744
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4746)
}
__typedArg0 := tmp4746
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V897)
}
__typedArg0 := V897
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4748)
}
__typedArg0 := tmp4748
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4749)
}
__typedArg0 := tmp4749
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4750)
}
__typedArg0 := tmp4750
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symsubst), tmp4745, tmp4747, tmp4751)
return


} else {
__e.Return(V897)
return
}


}, 1)

tmp4854 := Call(__e, ns2_1set, symshen_4remove_1indirection, tmp4738)


_ = tmp4854

tmp4855 := MakeNative(func(__e *ControlFlow) {
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
tmp4954 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4900 Obj

if True == tmp4954 {
tmp4952 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4953 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4952)
}
__typedArg0 := tmp4952
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4902 Obj

if True == tmp4953 {
tmp4949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4949)
}
__typedArg0 := tmp4949
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4951 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4950)
}
__typedArg0 := tmp4950
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4904 Obj

if True == tmp4951 {
tmp4945 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4945)
}
__typedArg0 := tmp4945
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4946)
}
__typedArg0 := tmp4946
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symand, tmp4947)
}
__typedArg0 := symand
__typedArg1 := tmp4947
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4906 Obj

if True == tmp4948 {
tmp4941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4941)
}
__typedArg0 := tmp4941
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4942)
}
__typedArg0 := tmp4942
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4944 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4943)
}
__typedArg0 := tmp4943
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4908 Obj

if True == tmp4944 {
tmp4936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4936)
}
__typedArg0 := tmp4936
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4937)
}
__typedArg0 := tmp4937
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4938)
}
__typedArg0 := tmp4938
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4939)
}
__typedArg0 := tmp4939
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4910 Obj

if True == tmp4940 {
tmp4930 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4931 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4930)
}
__typedArg0 := tmp4930
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4932 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4931)
}
__typedArg0 := tmp4931
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4933 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4932)
}
__typedArg0 := tmp4932
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4933)
}
__typedArg0 := tmp4933
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4934)
}
__typedArg0 := Nil
__typedArg1 := tmp4934
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4912 Obj

if True == tmp4935 {
tmp4927 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4928 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4927)
}
__typedArg0 := tmp4927
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4929 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4928)
}
__typedArg0 := tmp4928
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4914 Obj

if True == tmp4929 {
tmp4923 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4924 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4923)
}
__typedArg0 := tmp4923
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4925 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4924)
}
__typedArg0 := tmp4924
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4926 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4925)
}
__typedArg0 := Nil
__typedArg1 := tmp4925
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4916 Obj

if True == tmp4926 {
tmp4918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4918)
}
__typedArg0 := tmp4918
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4920 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4919)
}
__typedArg0 := tmp4919
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4920)
}
__typedArg0 := tmp4920
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V900, tmp4921)
}
__typedArg0 := V900
__typedArg1 := tmp4921
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4917 Obj

if True == tmp4922 {
ifres4917 = True


} else {
ifres4917 = False


}

ifres4916 = ifres4917


} else {
ifres4916 = False


}

var ifres4915 Obj

if True == ifres4916 {
ifres4915 = True


} else {
ifres4915 = False


}

ifres4914 = ifres4915


} else {
ifres4914 = False


}

var ifres4913 Obj

if True == ifres4914 {
ifres4913 = True


} else {
ifres4913 = False


}

ifres4912 = ifres4913


} else {
ifres4912 = False


}

var ifres4911 Obj

if True == ifres4912 {
ifres4911 = True


} else {
ifres4911 = False


}

ifres4910 = ifres4911


} else {
ifres4910 = False


}

var ifres4909 Obj

if True == ifres4910 {
ifres4909 = True


} else {
ifres4909 = False


}

ifres4908 = ifres4909


} else {
ifres4908 = False


}

var ifres4907 Obj

if True == ifres4908 {
ifres4907 = True


} else {
ifres4907 = False


}

ifres4906 = ifres4907


} else {
ifres4906 = False


}

var ifres4905 Obj

if True == ifres4906 {
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

if True == ifres4900 {
tmp4856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4856)
}
__typedArg0 := tmp4856
return Call(__e, PrimFunc(symhd), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4861)
}
__typedArg0 := tmp4861
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4862)
}
__typedArg0 := tmp4862
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4863)
}
__typedArg0 := tmp4863
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4865 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4864)
}
__typedArg0 := tmp4864
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4866 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4866)
}
__typedArg0 := tmp4866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4865, tmp4867)
}
__typedArg0 := tmp4865
__typedArg1 := tmp4867
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4868, V902)
}
__typedArg0 := tmp4868
__typedArg1 := V902
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4pivot_1on) == __self {
__self1, __self2, __self3 = tmp4859, tmp4860, tmp4869
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4pivot_1on), tmp4859, tmp4860, tmp4869)
return


} else {
tmp4898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4878 Obj

if True == tmp4898 {
tmp4896 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4896)
}
__typedArg0 := tmp4896
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4880 Obj

if True == tmp4897 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4894)
}
__typedArg0 := tmp4894
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4882 Obj

if True == tmp4895 {
tmp4889 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4889)
}
__typedArg0 := tmp4889
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4890)
}
__typedArg0 := tmp4890
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4891)
}
__typedArg0 := Nil
__typedArg1 := tmp4891
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4884 Obj

if True == tmp4892 {
tmp4886 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4886)
}
__typedArg0 := tmp4886
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4888 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V900, tmp4887)
}
__typedArg0 := V900
__typedArg1 := tmp4887
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4885 Obj

if True == tmp4888 {
ifres4885 = True


} else {
ifres4885 = False


}

ifres4884 = ifres4885


} else {
ifres4884 = False


}

var ifres4883 Obj

if True == ifres4884 {
ifres4883 = True


} else {
ifres4883 = False


}

ifres4882 = ifres4883


} else {
ifres4882 = False


}

var ifres4881 Obj

if True == ifres4882 {
ifres4881 = True


} else {
ifres4881 = False


}

ifres4880 = ifres4881


} else {
ifres4880 = False


}

var ifres4879 Obj

if True == ifres4880 {
ifres4879 = True


} else {
ifres4879 = False


}

ifres4878 = ifres4879


} else {
ifres4878 = False


}

if True == ifres4878 {
tmp4870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4870)
}
__typedArg0 := tmp4870
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V901)
}
__typedArg0 := V901
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4873)
}
__typedArg0 := tmp4873
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp4874)
}
__typedArg0 := True
__typedArg1 := tmp4874
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4876 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4875, V902)
}
__typedArg0 := tmp4875
__typedArg1 := V902
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

if PrimFunc(symshen_4pivot_1on) == __self {
__self1, __self2, __self3 = tmp4871, tmp4872, tmp4876
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4pivot_1on), tmp4871, tmp4872, tmp4876)
return


} else {
__e.TailApply(PrimFunc(sym_8p), V902, V901)
return
}


}


}, 3)

tmp4955 := Call(__e, ns2_1set, symshen_4pivot_1on, tmp4855)


_ = tmp4955

tmp4956 := MakeNative(func(__e *ControlFlow) {
V905 := __e.Get(1)
_ = V905
V906 := __e.Get(2)
_ = V906
tmp4979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V905)
}
__typedArg0 := V905
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4970 Obj

if True == tmp4979 {
tmp4977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V905)
}
__typedArg0 := V905
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4977)
}
__typedArg0 := tmp4977
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4972 Obj

if True == tmp4978 {
tmp4974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V905)
}
__typedArg0 := V905
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4974)
}
__typedArg0 := tmp4974
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4975)
}
__typedArg0 := Nil
__typedArg1 := tmp4975
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4973 Obj

if True == tmp4976 {
ifres4973 = True


} else {
ifres4973 = False


}

ifres4972 = ifres4973


} else {
ifres4972 = False


}

var ifres4971 Obj

if True == ifres4972 {
ifres4971 = True


} else {
ifres4971 = False


}

ifres4970 = ifres4971


} else {
ifres4970 = False


}

if True == ifres4970 {
tmp4957 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V905)
}
__typedArg0 := V905
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4958 := Call(__e, PrimFunc(symshen_4op), tmp4957)


let__2731 := tmp4958
_ = let__2731

tmp4968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4skip, let__2731)
}
__typedArg0 := symshen_4skip
__typedArg1 := let__2731
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4968 {
__e.Return(V906)
return
} else {
tmp4959 := Call(__e, PrimFunc(symshen_4op1), let__2731)


tmp4960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V905)
}
__typedArg0 := V905
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4961 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4959, tmp4960)
}
__typedArg0 := tmp4959
__typedArg1 := tmp4960
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4962 := Call(__e, PrimFunc(symshen_4op2), let__2731)


tmp4963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V905)
}
__typedArg0 := V905
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4962, tmp4963)
}
__typedArg0 := tmp4962
__typedArg1 := tmp4963
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4965 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4964, Nil)
}
__typedArg0 := tmp4964
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4966 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4961, tmp4965)
}
__typedArg0 := tmp4961
__typedArg1 := tmp4965
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4factor_1selectors_1h), tmp4966, V906)
return


}


} else {
__e.Return(V906)
return
}


}, 2)

tmp4980 := Call(__e, ns2_1set, symshen_4factor_1selectors, tmp4956)


_ = tmp4980

tmp4981 := MakeNative(func(__e *ControlFlow) {
V910 := __e.Get(1)
_ = V910
tmp4989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons_2, V910)
}
__typedArg0 := symcons_2
__typedArg1 := V910
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4989 {
__e.Return(symcons)
return
} else {
tmp4987 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7string_2, V910)
}
__typedArg0 := symshen_4_7string_2
__typedArg1 := V910
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4987 {
__e.Return(sym_8s)
return
} else {
tmp4985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7vector_2, V910)
}
__typedArg0 := symshen_4_7vector_2
__typedArg1 := V910
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4985 {
__e.Return(sym_8v)
return
} else {
tmp4983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtuple_2, V910)
}
__typedArg0 := symtuple_2
__typedArg1 := V910
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4983 {
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

tmp4990 := Call(__e, ns2_1set, symshen_4op, tmp4981)


_ = tmp4990

tmp4991 := MakeNative(func(__e *ControlFlow) {
__self := __e.Get(0)
__self1 := __e.Get(1)
__self2 := __e.Get(2)
__selftop:
V911 := __self1
_ = V911
V912 := __self2
_ = V912
tmp5009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V911)
}
__typedArg0 := Nil
__typedArg1 := V911
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp5009 {
__e.Return(V912)
return
} else {
tmp5007 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V911)
}
__typedArg0 := V911
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp5007 {
tmp5003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V911)
}
__typedArg0 := V911
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp5004 := Call(__e, PrimFunc(symoccurrences), tmp5003, V912)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(tmp5004)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := tmp5004
__typedArg1 := MakeInteger(1)
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
tmp4992 := Call(__e, PrimFunc(symgensym), symSelect)


let__2732 := tmp4992
_ = let__2732

tmp4993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V911)
}
__typedArg0 := V911
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V911)
}
__typedArg0 := V911
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V911)
}
__typedArg0 := V911
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4996 := Call(__e, PrimFunc(symsubst), let__2732, tmp4995, V912)


tmp4997 := Call(__e, PrimFunc(symshen_4factor_1selectors_1h), tmp4994, tmp4996)


tmp4998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4997, Nil)
}
__typedArg0 := tmp4997
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4993, tmp4998)
}
__typedArg0 := tmp4993
__typedArg1 := tmp4998
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp5000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(let__2732, tmp4999)
}
__typedArg0 := let__2732
__typedArg1 := tmp4999
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp5000)
}
__typedArg0 := symlet
__typedArg1 := tmp5000
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp5001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V911)
}
__typedArg0 := V911
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

if PrimFunc(symshen_4factor_1selectors_1h) == __self {
__self1, __self2 = tmp5001, V912
__e.Tick()
goto __selftop
}
__e.TailApply(PrimFunc(symshen_4factor_1selectors_1h), tmp5001, V912)
return


}


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4factor_1selectors_1h)
return
}


}


}, 2)

__e.TailApply(ns2_1set, symshen_4factor_1selectors_1h, tmp4991)
return




}, 0)

