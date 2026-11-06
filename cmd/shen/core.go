package main

import . "github.com/pyrex41/shen-go/kl"

var CoreMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp2733 := MakeNative1(func(__e *ControlFlow, V540 Obj) Obj {
tmp2734 := Call1(__e, SymFn(symshen_4shen_1_6kl_1h), V540)


let__2581 := tmp2734
_ = let__2581

return TailCall1(__e, SymFn(symshen_4record_1and_1evaluate), let__2581)


panic("unreachable")
})

tmp2735 := Call2(__e, ns2_1set, symshen_4shen_1_6kl, tmp2733)


_ = tmp2735

tmp2736 := MakeNative1(func(__e *ControlFlow, V542 Obj) Obj {
tmp2785 := GConsP(__e, V542)

var ifres2759 Obj

if True == tmp2785 {
tmp2783 := GHd(__e, V542)

tmp2784 := GEq(__e, symdefun, tmp2783)

var ifres2761 Obj

if True == tmp2784 {
tmp2781 := GTl(__e, V542)

tmp2782 := GConsP(__e, tmp2781)

var ifres2763 Obj

if True == tmp2782 {
tmp2778 := GTl(__e, V542)

tmp2779 := GTl(__e, tmp2778)

tmp2780 := GConsP(__e, tmp2779)

var ifres2765 Obj

if True == tmp2780 {
tmp2774 := GTl(__e, V542)

tmp2775 := GTl(__e, tmp2774)

tmp2776 := GTl(__e, tmp2775)

tmp2777 := GConsP(__e, tmp2776)

var ifres2767 Obj

if True == tmp2777 {
tmp2769 := GTl(__e, V542)

tmp2770 := GTl(__e, tmp2769)

tmp2771 := GTl(__e, tmp2770)

tmp2772 := GTl(__e, tmp2771)

tmp2773 := GEq(__e, Nil, tmp2772)

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
tmp2742 := GTl(__e, V542)

tmp2743 := GHd(__e, tmp2742)

tmp2744 := Call1(__e, SymFn(symshen_4sysfunc_2), tmp2743)


var ifres2737 Obj

if True == tmp2744 {
tmp2738 := GTl(__e, V542)

tmp2739 := GHd(__e, tmp2738)

tmp2740 := Call3(__e, SymFn(symshen_4app), tmp2739, MakeString(" is not a legitimate function name\n"), symshen_4a)


tmp2741 := PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp2740)

ifres2737 = tmp2741


} else {
ifres2737 = symshen_4skip


}

let__2582 := ifres2737
_ = let__2582

tmp2745 := GTl(__e, V542)

tmp2746 := GHd(__e, tmp2745)

tmp2747 := GTl(__e, V542)

tmp2748 := GTl(__e, tmp2747)

tmp2749 := GHd(__e, tmp2748)

tmp2750 := Call1(__e, SymFn(symlength), tmp2749)


tmp2751 := Call2(__e, SymFn(symshen_4store_1arity), tmp2746, tmp2750)


let__2583 := tmp2751
_ = let__2583

tmp2752 := GTl(__e, V542)

tmp2753 := GHd(__e, tmp2752)

tmp2754 := Call2(__e, SymFn(symshen_4record_1kl), tmp2753, V542)


let__2584 := tmp2754
_ = let__2584

tmp2755 := Call1(__e, SymFn(symeval_1kl), V542)


let__2585 := tmp2755
_ = let__2585

tmp2756 := GTl(__e, V542)

tmp2757 := GHd(__e, tmp2756)

return TailCall1(__e, SymFn(symshen_4fn_1print), tmp2757)


} else {
return V542
}


panic("unreachable")
})

tmp2786 := Call2(__e, ns2_1set, symshen_4record_1and_1evaluate, tmp2736)


_ = tmp2786

tmp2787 := MakeNative1(func(__e *ControlFlow, V547 Obj) Obj {
tmp2888 := GConsP(__e, V547)

var ifres2880 Obj

if True == tmp2888 {
tmp2886 := GHd(__e, V547)

tmp2887 := GEq(__e, symdefine, tmp2886)

var ifres2882 Obj

if True == tmp2887 {
tmp2884 := GTl(__e, V547)

tmp2885 := GConsP(__e, tmp2884)

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
tmp2788 := GTl(__e, V547)

tmp2789 := GHd(__e, tmp2788)

tmp2790 := GTl(__e, V547)

tmp2791 := GTl(__e, tmp2790)

return TailCall2(__e, SymFn(symshen_4shendef_1_6kldef), tmp2789, tmp2791)


} else {
tmp2878 := GConsP(__e, V547)

var ifres2852 Obj

if True == tmp2878 {
tmp2876 := GHd(__e, V547)

tmp2877 := GEq(__e, symdefun, tmp2876)

var ifres2854 Obj

if True == tmp2877 {
tmp2874 := GTl(__e, V547)

tmp2875 := GConsP(__e, tmp2874)

var ifres2856 Obj

if True == tmp2875 {
tmp2871 := GTl(__e, V547)

tmp2872 := GTl(__e, tmp2871)

tmp2873 := GConsP(__e, tmp2872)

var ifres2858 Obj

if True == tmp2873 {
tmp2867 := GTl(__e, V547)

tmp2868 := GTl(__e, tmp2867)

tmp2869 := GTl(__e, tmp2868)

tmp2870 := GConsP(__e, tmp2869)

var ifres2860 Obj

if True == tmp2870 {
tmp2862 := GTl(__e, V547)

tmp2863 := GTl(__e, tmp2862)

tmp2864 := GTl(__e, tmp2863)

tmp2865 := GTl(__e, tmp2864)

tmp2866 := GEq(__e, Nil, tmp2865)

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
return V547
} else {
tmp2850 := GConsP(__e, V547)

var ifres2831 Obj

if True == tmp2850 {
tmp2848 := GHd(__e, V547)

tmp2849 := GEq(__e, symtype, tmp2848)

var ifres2833 Obj

if True == tmp2849 {
tmp2846 := GTl(__e, V547)

tmp2847 := GConsP(__e, tmp2846)

var ifres2835 Obj

if True == tmp2847 {
tmp2843 := GTl(__e, V547)

tmp2844 := GTl(__e, tmp2843)

tmp2845 := GConsP(__e, tmp2844)

var ifres2837 Obj

if True == tmp2845 {
tmp2839 := GTl(__e, V547)

tmp2840 := GTl(__e, tmp2839)

tmp2841 := GTl(__e, tmp2840)

tmp2842 := GEq(__e, Nil, tmp2841)

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
tmp2792 := GTl(__e, V547)

tmp2793 := GHd(__e, tmp2792)

tmp2794 := GTl(__e, V547)

tmp2795 := GTl(__e, tmp2794)

tmp2796 := GHd(__e, tmp2795)

tmp2797 := Call1(__e, SymFn(symshen_4rcons__form), tmp2796)


tmp2798 := GCons(__e, tmp2797, Nil)

tmp2799 := GCons(__e, tmp2793, tmp2798)

return GCons(__e, symtype, tmp2799)


} else {
tmp2829 := GConsP(__e, V547)

var ifres2810 Obj

if True == tmp2829 {
tmp2827 := GHd(__e, V547)

tmp2828 := GEq(__e, syminput_7, tmp2827)

var ifres2812 Obj

if True == tmp2828 {
tmp2825 := GTl(__e, V547)

tmp2826 := GConsP(__e, tmp2825)

var ifres2814 Obj

if True == tmp2826 {
tmp2822 := GTl(__e, V547)

tmp2823 := GTl(__e, tmp2822)

tmp2824 := GConsP(__e, tmp2823)

var ifres2816 Obj

if True == tmp2824 {
tmp2818 := GTl(__e, V547)

tmp2819 := GTl(__e, tmp2818)

tmp2820 := GTl(__e, tmp2819)

tmp2821 := GEq(__e, Nil, tmp2820)

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
tmp2800 := GTl(__e, V547)

tmp2801 := GHd(__e, tmp2800)

tmp2802 := Call1(__e, SymFn(symshen_4rcons__form), tmp2801)


tmp2803 := GTl(__e, V547)

tmp2804 := GTl(__e, tmp2803)

tmp2805 := GCons(__e, tmp2802, tmp2804)

return GCons(__e, syminput_7, tmp2805)


} else {
tmp2808 := GConsP(__e, V547)

if True == tmp2808 {
tmp2806 := MakeNative1(func(__e *ControlFlow, Z548 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4shen_1_6kl_1h), Z548)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp2806, V547)


} else {
return V547
}


}


}


}


}


panic("unreachable")
})

tmp2889 := Call2(__e, ns2_1set, symshen_4shen_1_6kl_1h, tmp2787)


_ = tmp2889

tmp2890 := MakeNative2(func(__e *ControlFlow, V549, V550 Obj) Obj {
tmp2891 := MakeNative1(func(__e *ControlFlow, Z551 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5define_6), Z551)
panic("unreachable")
})

tmp2892 := GCons(__e, V549, V550)

return TailCall2(__e, SymFn(symcompile), tmp2891, tmp2892)


panic("unreachable")
})

tmp2893 := Call2(__e, ns2_1set, symshen_4shendef_1_6kldef, tmp2890)


_ = tmp2893

tmp2894 := MakeNative1(func(__e *ControlFlow, V552 Obj) Obj {
tmp2895 := Call1(__e, SymFn(symshen_4_5name_6), V552)


let__2587 := tmp2895
_ = let__2587

tmp2921 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2587)


var ifres2896 Obj

if True == tmp2921 {
tmp2897 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2896 = tmp2897


} else {
tmp2898 := Call1(__e, SymFn(symshen_4_5_1out), let__2587)


let__2588 := tmp2898
_ = let__2588

tmp2899 := Call1(__e, SymFn(symshen_4in_1_6), let__2587)


let__2589 := tmp2899
_ = let__2589

tmp2920 := Call2(__e, SymFn(symshen_4hds_a_2), let__2589, sym_i)


var ifres2900 Obj

if True == tmp2920 {
tmp2901 := Call1(__e, SymFn(symtail), let__2589)


let__2590 := tmp2901
_ = let__2590

tmp2902 := Call1(__e, SymFn(symshen_4_5signature_6), let__2590)


let__2591 := tmp2902
_ = let__2591

tmp2918 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2591)


var ifres2903 Obj

if True == tmp2918 {
tmp2904 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2903 = tmp2904


} else {
tmp2905 := Call1(__e, SymFn(symshen_4in_1_6), let__2591)


let__2592 := tmp2905
_ = let__2592

tmp2917 := Call2(__e, SymFn(symshen_4hds_a_2), let__2592, sym_j)


var ifres2906 Obj

if True == tmp2917 {
tmp2907 := Call1(__e, SymFn(symtail), let__2592)


let__2593 := tmp2907
_ = let__2593

tmp2908 := Call1(__e, SymFn(symshen_4_5rules_6), let__2593)


let__2594 := tmp2908
_ = let__2594

tmp2915 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2594)


var ifres2909 Obj

if True == tmp2915 {
tmp2910 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2909 = tmp2910


} else {
tmp2911 := Call1(__e, SymFn(symshen_4_5_1out), let__2594)


let__2595 := tmp2911
_ = let__2595

tmp2912 := Call1(__e, SymFn(symshen_4in_1_6), let__2594)


let__2596 := tmp2912
_ = let__2596

tmp2913 := Call2(__e, SymFn(symshen_4shendef_1_6kldef_1h), let__2588, let__2595)


tmp2914 := Call2(__e, SymFn(symshen_4comb), let__2596, tmp2913)


ifres2909 = tmp2914


}

ifres2906 = ifres2909


} else {
tmp2916 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2906 = tmp2916


}

ifres2903 = ifres2906


}

ifres2900 = ifres2903


} else {
tmp2919 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2900 = tmp2919


}

ifres2896 = ifres2900


}

let__2586 := ifres2896
_ = let__2586

tmp2939 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2586)


if True == tmp2939 {
tmp2922 := Call1(__e, SymFn(symshen_4_5name_6), V552)


let__2598 := tmp2922
_ = let__2598

tmp2935 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2598)


var ifres2923 Obj

if True == tmp2935 {
tmp2924 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2923 = tmp2924


} else {
tmp2925 := Call1(__e, SymFn(symshen_4_5_1out), let__2598)


let__2599 := tmp2925
_ = let__2599

tmp2926 := Call1(__e, SymFn(symshen_4in_1_6), let__2598)


let__2600 := tmp2926
_ = let__2600

tmp2927 := Call1(__e, SymFn(symshen_4_5rules_6), let__2600)


let__2601 := tmp2927
_ = let__2601

tmp2934 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2601)


var ifres2928 Obj

if True == tmp2934 {
tmp2929 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2928 = tmp2929


} else {
tmp2930 := Call1(__e, SymFn(symshen_4_5_1out), let__2601)


let__2602 := tmp2930
_ = let__2602

tmp2931 := Call1(__e, SymFn(symshen_4in_1_6), let__2601)


let__2603 := tmp2931
_ = let__2603

tmp2932 := Call2(__e, SymFn(symshen_4shendef_1_6kldef_1h), let__2599, let__2602)


tmp2933 := Call2(__e, SymFn(symshen_4comb), let__2603, tmp2932)


ifres2928 = tmp2933


}

ifres2923 = ifres2928


}

let__2597 := ifres2923
_ = let__2597

tmp2937 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2597)


if True == tmp2937 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2597
}


} else {
return let__2586
}


panic("unreachable")
})

tmp2940 := Call2(__e, ns2_1set, symshen_4_5define_6, tmp2894)


_ = tmp2940

tmp2941 := MakeNative2(func(__e *ControlFlow, V571, V572 Obj) Obj {
tmp2942 := MakeNative1(func(__e *ControlFlow, Z574 Obj) Obj {
return TailCall1(__e, SymFn(symfst), Z574)
panic("unreachable")
})

tmp2943 := Call2(__e, SymFn(symmap), tmp2942, V572)


let__2604 := tmp2943
_ = let__2604

tmp2944 := Call2(__e, SymFn(symshen_4arity_1chk), V571, let__2604)


let__2605 := tmp2944
_ = let__2605

tmp2945 := MakeNative1(func(__e *ControlFlow, Z577 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4free_1var_1chk), V571, Z577)
panic("unreachable")
})

tmp2946 := Call2(__e, SymFn(symmap), tmp2945, V572)


let__2606 := tmp2946
_ = let__2606

tmp2947 := Call1(__e, SymFn(symshen_4unprotect), V572)


let__2607 := tmp2947
_ = let__2607

tmp2948 := Call3(__e, SymFn(symshen_4compile_1to_1kl), V571, let__2607, let__2605)


tmp2949 := Call1(__e, SymFn(symshen_4factorise_1code), tmp2948)


let__2608 := tmp2949
_ = let__2608

return let__2608


panic("unreachable")
})

tmp2950 := Call2(__e, ns2_1set, symshen_4shendef_1_6kldef_1h, tmp2941)


_ = tmp2950

tmp2951 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V580 := __self1
_ = V580
tmp2977 := Call1(__e, SymFn(symtuple_2), V580)


if True == tmp2977 {
tmp2952 := Call1(__e, SymFn(symfst), V580)


tmp2953 := Call1(__e, SymFn(symshen_4unprotect), tmp2952)


tmp2954 := Call1(__e, SymFn(symsnd), V580)


tmp2955 := Call1(__e, SymFn(symshen_4unprotect), tmp2954)


return TailCall2(__e, SymFn(sym_8p), tmp2953, tmp2955)


} else {
tmp2975 := GConsP(__e, V580)

var ifres2962 Obj

if True == tmp2975 {
tmp2973 := GHd(__e, V580)

tmp2974 := GEq(__e, symprotect, tmp2973)

var ifres2964 Obj

if True == tmp2974 {
tmp2971 := GTl(__e, V580)

tmp2972 := GConsP(__e, tmp2971)

var ifres2966 Obj

if True == tmp2972 {
tmp2968 := GTl(__e, V580)

tmp2969 := GTl(__e, tmp2968)

tmp2970 := GEq(__e, Nil, tmp2969)

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
tmp2956 := GTl(__e, V580)

tmp2957 := GHd(__e, tmp2956)

if SymFn(symshen_4unprotect) == __self {
__self1 = tmp2957
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4unprotect), tmp2957)


} else {
tmp2960 := GConsP(__e, V580)

if True == tmp2960 {
tmp2958 := MakeNative1(func(__e *ControlFlow, Z581 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4unprotect), Z581)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp2958, V580)


} else {
return V580
}


}


}


panic("unreachable")
})
return __self
})()

tmp2978 := Call2(__e, ns2_1set, symshen_4unprotect, tmp2951)


_ = tmp2978

tmp2979 := MakeNative1(func(__e *ControlFlow, V582 Obj) Obj {
tmp2993 := GConsP(__e, V582)

var ifres2980 Obj

if True == tmp2993 {
tmp2981 := Call1(__e, SymFn(symhead), V582)


let__2610 := tmp2981
_ = let__2610

tmp2982 := Call1(__e, SymFn(symtail), V582)


let__2611 := tmp2982
_ = let__2611

tmp2990 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, let__2610)

var ifres2986 Obj

if True == tmp2990 {
tmp2988 := PrimGuard1(__e, symvariable_2, PrimIsVariable, let__2610)

tmp2989 := GNot(__e, tmp2988)

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
tmp2984 := Call3(__e, SymFn(symshen_4app), let__2610, MakeString(" is not a legitimate function name.\n"), symshen_4a)


tmp2985 := PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp2984)

ifres2983 = tmp2985


}

tmp2991 := Call2(__e, SymFn(symshen_4comb), let__2611, ifres2983)


ifres2980 = tmp2991


} else {
tmp2992 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2980 = tmp2992


}

let__2609 := ifres2980
_ = let__2609

tmp2995 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2609)


if True == tmp2995 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2609
}


panic("unreachable")
})

tmp2996 := Call2(__e, ns2_1set, symshen_4_5name_6, tmp2979)


_ = tmp2996

tmp2997 := MakeNative1(func(__e *ControlFlow, V586 Obj) Obj {
tmp3016 := GConsP(__e, V586)

var ifres2998 Obj

if True == tmp3016 {
tmp2999 := Call1(__e, SymFn(symhead), V586)


let__2613 := tmp2999
_ = let__2613

tmp3000 := Call1(__e, SymFn(symtail), V586)


let__2614 := tmp3000
_ = let__2614

tmp3001 := Call1(__e, SymFn(symshen_4_5signature_6), let__2614)


let__2615 := tmp3001
_ = let__2615

tmp3014 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2615)


var ifres3002 Obj

if True == tmp3014 {
tmp3003 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3002 = tmp3003


} else {
tmp3004 := Call1(__e, SymFn(symshen_4_5_1out), let__2615)


let__2616 := tmp3004
_ = let__2616

tmp3005 := Call1(__e, SymFn(symshen_4in_1_6), let__2615)


let__2617 := tmp3005
_ = let__2617

tmp3010 := GCons(__e, sym_j, Nil)

tmp3011 := GCons(__e, sym_i, tmp3010)

tmp3012 := Call2(__e, SymFn(symelement_2), let__2613, tmp3011)


tmp3013 := GNot(__e, tmp3012)

var ifres3006 Obj

if True == tmp3013 {
tmp3007 := GCons(__e, let__2613, let__2616)

tmp3008 := Call2(__e, SymFn(symshen_4comb), let__2617, tmp3007)


ifres3006 = tmp3008


} else {
tmp3009 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3006 = tmp3009


}

ifres3002 = ifres3006


}

ifres2998 = ifres3002


} else {
tmp3015 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2998 = tmp3015


}

let__2612 := ifres2998
_ = let__2612

tmp3026 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2612)


if True == tmp3026 {
tmp3017 := Call1(__e, SymFn(sym_5e_6), V586)


let__2619 := tmp3017
_ = let__2619

tmp3022 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2619)


var ifres3018 Obj

if True == tmp3022 {
tmp3019 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3018 = tmp3019


} else {
tmp3020 := Call1(__e, SymFn(symshen_4in_1_6), let__2619)


let__2620 := tmp3020
_ = let__2620

tmp3021 := Call2(__e, SymFn(symshen_4comb), let__2620, Nil)


ifres3018 = tmp3021


}

let__2618 := ifres3018
_ = let__2618

tmp3024 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2618)


if True == tmp3024 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2618
}


} else {
return let__2612
}


panic("unreachable")
})

tmp3027 := Call2(__e, ns2_1set, symshen_4_5signature_6, tmp2997)


_ = tmp3027

tmp3028 := MakeNative1(func(__e *ControlFlow, V596 Obj) Obj {
tmp3029 := Call1(__e, SymFn(symshen_4_5rule_6), V596)


let__2622 := tmp3029
_ = let__2622

tmp3043 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2622)


var ifres3030 Obj

if True == tmp3043 {
tmp3031 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3030 = tmp3031


} else {
tmp3032 := Call1(__e, SymFn(symshen_4_5_1out), let__2622)


let__2623 := tmp3032
_ = let__2623

tmp3033 := Call1(__e, SymFn(symshen_4in_1_6), let__2622)


let__2624 := tmp3033
_ = let__2624

tmp3034 := Call1(__e, SymFn(symshen_4_5rules_6), let__2624)


let__2625 := tmp3034
_ = let__2625

tmp3042 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2625)


var ifres3035 Obj

if True == tmp3042 {
tmp3036 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3035 = tmp3036


} else {
tmp3037 := Call1(__e, SymFn(symshen_4_5_1out), let__2625)


let__2626 := tmp3037
_ = let__2626

tmp3038 := Call1(__e, SymFn(symshen_4in_1_6), let__2625)


let__2627 := tmp3038
_ = let__2627

tmp3039 := Call1(__e, SymFn(symshen_4linearise), let__2623)


tmp3040 := GCons(__e, tmp3039, let__2626)

tmp3041 := Call2(__e, SymFn(symshen_4comb), let__2627, tmp3040)


ifres3035 = tmp3041


}

ifres3030 = ifres3035


}

let__2621 := ifres3030
_ = let__2621

tmp3059 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2621)


if True == tmp3059 {
tmp3044 := Call1(__e, SymFn(sym_5_b_6), V596)


let__2629 := tmp3044
_ = let__2629

tmp3055 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2629)


var ifres3045 Obj

if True == tmp3055 {
tmp3046 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3045 = tmp3046


} else {
tmp3047 := Call1(__e, SymFn(symshen_4_5_1out), let__2629)


let__2630 := tmp3047
_ = let__2630

tmp3048 := Call1(__e, SymFn(symshen_4in_1_6), let__2629)


let__2631 := tmp3048
_ = let__2631

tmp3053 := Call1(__e, SymFn(symempty_2), let__2630)


var ifres3049 Obj

if True == tmp3053 {
ifres3049 = Nil


} else {
tmp3050 := Call3(__e, SymFn(symshen_4app), let__2630, MakeString("\n ..."), symshen_4r)


tmp3052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("Shen syntax error here:\n "), tmp3050))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("Shen syntax error here:\n "), tmp3050)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()

ifres3049 = tmp3052


}

tmp3054 := Call2(__e, SymFn(symshen_4comb), let__2631, ifres3049)


ifres3045 = tmp3054


}

let__2628 := ifres3045
_ = let__2628

tmp3057 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2628)


if True == tmp3057 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2628
}


} else {
return let__2621
}


panic("unreachable")
})

tmp3060 := Call2(__e, ns2_1set, symshen_4_5rules_6, tmp3028)


_ = tmp3060

tmp3061 := MakeNative1(func(__e *ControlFlow, V610 Obj) Obj {
tmp3066 := Call1(__e, SymFn(symtuple_2), V610)


if True == tmp3066 {
tmp3062 := Call1(__e, SymFn(symfst), V610)


tmp3063 := Call1(__e, SymFn(symfst), V610)


tmp3064 := Call1(__e, SymFn(symsnd), V610)


return TailCall4(__e, SymFn(symshen_4linearise_1h), tmp3062, tmp3063, Nil, tmp3064)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.linearise"))
}


panic("unreachable")
})

tmp3067 := Call2(__e, ns2_1set, symshen_4linearise, tmp3061)


_ = tmp3067

tmp3068 := (func() Obj {
var __self Obj
__self = MakeNative4(func(__e *ControlFlow, __self1, __self2, __self3, __self4 Obj) Obj {
__selftop:
V623 := __self1
_ = V623
V624 := __self2
_ = V624
V625 := __self3
_ = V625
V626 := __self4
_ = V626
tmp3105 := GEq(__e, Nil, V623)

if True == tmp3105 {
return TailCall2(__e, SymFn(sym_8p), V624, V626)
} else {
tmp3103 := GConsP(__e, V623)

var ifres3099 Obj

if True == tmp3103 {
tmp3101 := GHd(__e, V623)

tmp3102 := GConsP(__e, tmp3101)

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
tmp3069 := GHd(__e, V623)

tmp3070 := GTl(__e, V623)

tmp3071 := Call2(__e, SymFn(symappend), tmp3069, tmp3070)


if SymFn(symshen_4linearise_1h) == __self {
__self1, __self2, __self3, __self4 = tmp3071, V624, V625, V626
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4linearise_1h), tmp3071, V624, V625, V626)


} else {
tmp3097 := GConsP(__e, V623)

var ifres3093 Obj

if True == tmp3097 {
tmp3095 := GHd(__e, V623)

tmp3096 := PrimGuard1(__e, symvariable_2, PrimIsVariable, tmp3095)

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
tmp3087 := GHd(__e, V623)

tmp3088 := Call2(__e, SymFn(symelement_2), tmp3087, V625)


if True == tmp3088 {
tmp3072 := Call1(__e, SymFn(symgensym), symV)


let__2632 := tmp3072
_ = let__2632

tmp3073 := GTl(__e, V623)

tmp3074 := GHd(__e, V623)

tmp3075 := Call3(__e, SymFn(symshen_4rep_1X), tmp3074, let__2632, V624)


tmp3076 := GHd(__e, V623)

tmp3077 := GCons(__e, tmp3076, Nil)

tmp3078 := GCons(__e, let__2632, tmp3077)

tmp3079 := GCons(__e, sym_a, tmp3078)

tmp3080 := GCons(__e, V626, Nil)

tmp3081 := GCons(__e, tmp3079, tmp3080)

tmp3082 := GCons(__e, symwhere, tmp3081)

if SymFn(symshen_4linearise_1h) == __self {
__self1, __self2, __self3, __self4 = tmp3073, tmp3075, V625, tmp3082
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4linearise_1h), tmp3073, tmp3075, V625, tmp3082)


} else {
tmp3083 := GTl(__e, V623)

tmp3084 := GHd(__e, V623)

tmp3085 := GCons(__e, tmp3084, V625)

if SymFn(symshen_4linearise_1h) == __self {
__self1, __self2, __self3, __self4 = tmp3083, V624, tmp3085, V626
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4linearise_1h), tmp3083, V624, tmp3085, V626)


}


} else {
tmp3091 := GConsP(__e, V623)

if True == tmp3091 {
tmp3089 := GTl(__e, V623)

if SymFn(symshen_4linearise_1h) == __self {
__self1, __self2, __self3, __self4 = tmp3089, V624, V625, V626
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4linearise_1h), tmp3089, V624, V625, V626)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.linearise-h"))
}


}


}


}


panic("unreachable")
})
return __self
})()

tmp3106 := Call2(__e, ns2_1set, symshen_4linearise_1h, tmp3068)


_ = tmp3106

tmp3107 := MakeNative1(func(__e *ControlFlow, V628 Obj) Obj {
tmp3108 := Call1(__e, SymFn(symshen_4_5patterns_6), V628)


let__2634 := tmp3108
_ = let__2634

tmp3136 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2634)


var ifres3109 Obj

if True == tmp3136 {
tmp3110 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3109 = tmp3110


} else {
tmp3111 := Call1(__e, SymFn(symshen_4_5_1out), let__2634)


let__2635 := tmp3111
_ = let__2635

tmp3112 := Call1(__e, SymFn(symshen_4in_1_6), let__2634)


let__2636 := tmp3112
_ = let__2636

tmp3135 := Call2(__e, SymFn(symshen_4hds_a_2), let__2636, sym_1_6)


var ifres3113 Obj

if True == tmp3135 {
tmp3114 := Call1(__e, SymFn(symtail), let__2636)


let__2637 := tmp3114
_ = let__2637

tmp3133 := GConsP(__e, let__2637)

var ifres3115 Obj

if True == tmp3133 {
tmp3116 := Call1(__e, SymFn(symhead), let__2637)


let__2638 := tmp3116
_ = let__2638

tmp3117 := Call1(__e, SymFn(symtail), let__2637)


let__2639 := tmp3117
_ = let__2639

tmp3131 := Call2(__e, SymFn(symshen_4hds_a_2), let__2639, symwhere)


var ifres3118 Obj

if True == tmp3131 {
tmp3119 := Call1(__e, SymFn(symtail), let__2639)


let__2640 := tmp3119
_ = let__2640

tmp3129 := GConsP(__e, let__2640)

var ifres3120 Obj

if True == tmp3129 {
tmp3121 := Call1(__e, SymFn(symhead), let__2640)


let__2641 := tmp3121
_ = let__2641

tmp3122 := Call1(__e, SymFn(symtail), let__2640)


let__2642 := tmp3122
_ = let__2642

tmp3123 := GCons(__e, let__2638, Nil)

tmp3124 := GCons(__e, let__2641, tmp3123)

tmp3125 := GCons(__e, symwhere, tmp3124)

tmp3126 := Call2(__e, SymFn(sym_8p), let__2635, tmp3125)


tmp3127 := Call2(__e, SymFn(symshen_4comb), let__2642, tmp3126)


ifres3120 = tmp3127


} else {
tmp3128 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3120 = tmp3128


}

ifres3118 = ifres3120


} else {
tmp3130 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3118 = tmp3130


}

ifres3115 = ifres3118


} else {
tmp3132 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3115 = tmp3132


}

ifres3113 = ifres3115


} else {
tmp3134 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3113 = tmp3134


}

ifres3109 = ifres3113


}

let__2633 := ifres3109
_ = let__2633

tmp3211 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2633)


if True == tmp3211 {
tmp3137 := Call1(__e, SymFn(symshen_4_5patterns_6), V628)


let__2644 := tmp3137
_ = let__2644

tmp3153 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2644)


var ifres3138 Obj

if True == tmp3153 {
tmp3139 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3138 = tmp3139


} else {
tmp3140 := Call1(__e, SymFn(symshen_4_5_1out), let__2644)


let__2645 := tmp3140
_ = let__2645

tmp3141 := Call1(__e, SymFn(symshen_4in_1_6), let__2644)


let__2646 := tmp3141
_ = let__2646

tmp3152 := Call2(__e, SymFn(symshen_4hds_a_2), let__2646, sym_1_6)


var ifres3142 Obj

if True == tmp3152 {
tmp3143 := Call1(__e, SymFn(symtail), let__2646)


let__2647 := tmp3143
_ = let__2647

tmp3150 := GConsP(__e, let__2647)

var ifres3144 Obj

if True == tmp3150 {
tmp3145 := Call1(__e, SymFn(symhead), let__2647)


let__2648 := tmp3145
_ = let__2648

tmp3146 := Call1(__e, SymFn(symtail), let__2647)


let__2649 := tmp3146
_ = let__2649

tmp3147 := Call2(__e, SymFn(sym_8p), let__2645, let__2648)


tmp3148 := Call2(__e, SymFn(symshen_4comb), let__2649, tmp3147)


ifres3144 = tmp3148


} else {
tmp3149 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3144 = tmp3149


}

ifres3142 = ifres3144


} else {
tmp3151 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3142 = tmp3151


}

ifres3138 = ifres3142


}

let__2643 := ifres3138
_ = let__2643

tmp3209 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2643)


if True == tmp3209 {
tmp3154 := Call1(__e, SymFn(symshen_4_5patterns_6), V628)


let__2651 := tmp3154
_ = let__2651

tmp3184 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2651)


var ifres3155 Obj

if True == tmp3184 {
tmp3156 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3155 = tmp3156


} else {
tmp3157 := Call1(__e, SymFn(symshen_4_5_1out), let__2651)


let__2652 := tmp3157
_ = let__2652

tmp3158 := Call1(__e, SymFn(symshen_4in_1_6), let__2651)


let__2653 := tmp3158
_ = let__2653

tmp3183 := Call2(__e, SymFn(symshen_4hds_a_2), let__2653, sym_5_1)


var ifres3159 Obj

if True == tmp3183 {
tmp3160 := Call1(__e, SymFn(symtail), let__2653)


let__2654 := tmp3160
_ = let__2654

tmp3181 := GConsP(__e, let__2654)

var ifres3161 Obj

if True == tmp3181 {
tmp3162 := Call1(__e, SymFn(symhead), let__2654)


let__2655 := tmp3162
_ = let__2655

tmp3163 := Call1(__e, SymFn(symtail), let__2654)


let__2656 := tmp3163
_ = let__2656

tmp3179 := Call2(__e, SymFn(symshen_4hds_a_2), let__2656, symwhere)


var ifres3164 Obj

if True == tmp3179 {
tmp3165 := Call1(__e, SymFn(symtail), let__2656)


let__2657 := tmp3165
_ = let__2657

tmp3177 := GConsP(__e, let__2657)

var ifres3166 Obj

if True == tmp3177 {
tmp3167 := Call1(__e, SymFn(symhead), let__2657)


let__2658 := tmp3167
_ = let__2658

tmp3168 := Call1(__e, SymFn(symtail), let__2657)


let__2659 := tmp3168
_ = let__2659

tmp3169 := GCons(__e, let__2655, Nil)

tmp3170 := GCons(__e, symshen_4choicepoint_b, tmp3169)

tmp3171 := GCons(__e, tmp3170, Nil)

tmp3172 := GCons(__e, let__2658, tmp3171)

tmp3173 := GCons(__e, symwhere, tmp3172)

tmp3174 := Call2(__e, SymFn(sym_8p), let__2652, tmp3173)


tmp3175 := Call2(__e, SymFn(symshen_4comb), let__2659, tmp3174)


ifres3166 = tmp3175


} else {
tmp3176 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3166 = tmp3176


}

ifres3164 = ifres3166


} else {
tmp3178 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3164 = tmp3178


}

ifres3161 = ifres3164


} else {
tmp3180 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3161 = tmp3180


}

ifres3159 = ifres3161


} else {
tmp3182 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3159 = tmp3182


}

ifres3155 = ifres3159


}

let__2650 := ifres3155
_ = let__2650

tmp3207 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2650)


if True == tmp3207 {
tmp3185 := Call1(__e, SymFn(symshen_4_5patterns_6), V628)


let__2661 := tmp3185
_ = let__2661

tmp3203 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2661)


var ifres3186 Obj

if True == tmp3203 {
tmp3187 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3186 = tmp3187


} else {
tmp3188 := Call1(__e, SymFn(symshen_4_5_1out), let__2661)


let__2662 := tmp3188
_ = let__2662

tmp3189 := Call1(__e, SymFn(symshen_4in_1_6), let__2661)


let__2663 := tmp3189
_ = let__2663

tmp3202 := Call2(__e, SymFn(symshen_4hds_a_2), let__2663, sym_5_1)


var ifres3190 Obj

if True == tmp3202 {
tmp3191 := Call1(__e, SymFn(symtail), let__2663)


let__2664 := tmp3191
_ = let__2664

tmp3200 := GConsP(__e, let__2664)

var ifres3192 Obj

if True == tmp3200 {
tmp3193 := Call1(__e, SymFn(symhead), let__2664)


let__2665 := tmp3193
_ = let__2665

tmp3194 := Call1(__e, SymFn(symtail), let__2664)


let__2666 := tmp3194
_ = let__2666

tmp3195 := GCons(__e, let__2665, Nil)

tmp3196 := GCons(__e, symshen_4choicepoint_b, tmp3195)

tmp3197 := Call2(__e, SymFn(sym_8p), let__2662, tmp3196)


tmp3198 := Call2(__e, SymFn(symshen_4comb), let__2666, tmp3197)


ifres3192 = tmp3198


} else {
tmp3199 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3192 = tmp3199


}

ifres3190 = ifres3192


} else {
tmp3201 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3190 = tmp3201


}

ifres3186 = ifres3190


}

let__2660 := ifres3186
_ = let__2660

tmp3205 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2660)


if True == tmp3205 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2660
}


} else {
return let__2650
}


} else {
return let__2643
}


} else {
return let__2633
}


panic("unreachable")
})

tmp3212 := Call2(__e, ns2_1set, symshen_4_5rule_6, tmp3107)


_ = tmp3212

tmp3213 := MakeNative1(func(__e *ControlFlow, V663 Obj) Obj {
tmp3214 := Call1(__e, SymFn(symshen_4_5pattern_6), V663)


let__2668 := tmp3214
_ = let__2668

tmp3227 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2668)


var ifres3215 Obj

if True == tmp3227 {
tmp3216 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3215 = tmp3216


} else {
tmp3217 := Call1(__e, SymFn(symshen_4_5_1out), let__2668)


let__2669 := tmp3217
_ = let__2669

tmp3218 := Call1(__e, SymFn(symshen_4in_1_6), let__2668)


let__2670 := tmp3218
_ = let__2670

tmp3219 := Call1(__e, SymFn(symshen_4_5patterns_6), let__2670)


let__2671 := tmp3219
_ = let__2671

tmp3226 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2671)


var ifres3220 Obj

if True == tmp3226 {
tmp3221 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3220 = tmp3221


} else {
tmp3222 := Call1(__e, SymFn(symshen_4_5_1out), let__2671)


let__2672 := tmp3222
_ = let__2672

tmp3223 := Call1(__e, SymFn(symshen_4in_1_6), let__2671)


let__2673 := tmp3223
_ = let__2673

tmp3224 := GCons(__e, let__2669, let__2672)

tmp3225 := Call2(__e, SymFn(symshen_4comb), let__2673, tmp3224)


ifres3220 = tmp3225


}

ifres3215 = ifres3220


}

let__2667 := ifres3215
_ = let__2667

tmp3237 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2667)


if True == tmp3237 {
tmp3228 := Call1(__e, SymFn(sym_5e_6), V663)


let__2675 := tmp3228
_ = let__2675

tmp3233 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2675)


var ifres3229 Obj

if True == tmp3233 {
tmp3230 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3229 = tmp3230


} else {
tmp3231 := Call1(__e, SymFn(symshen_4in_1_6), let__2675)


let__2676 := tmp3231
_ = let__2676

tmp3232 := Call2(__e, SymFn(symshen_4comb), let__2676, Nil)


ifres3229 = tmp3232


}

let__2674 := ifres3229
_ = let__2674

tmp3235 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2674)


if True == tmp3235 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2674
}


} else {
return let__2667
}


panic("unreachable")
})

tmp3238 := Call2(__e, ns2_1set, symshen_4_5patterns_6, tmp3213)


_ = tmp3238

tmp3239 := MakeNative1(func(__e *ControlFlow, V674 Obj) Obj {
tmp3260 := Call1(__e, SymFn(symshen_4ccons_2), V674)


var ifres3240 Obj

if True == tmp3260 {
tmp3241 := Call1(__e, SymFn(symhead), V674)


let__2678 := tmp3241
_ = let__2678

tmp3242 := Call1(__e, SymFn(symtail), V674)


let__2679 := tmp3242
_ = let__2679

tmp3258 := Call2(__e, SymFn(symshen_4hds_a_2), let__2678, symvector)


var ifres3243 Obj

if True == tmp3258 {
tmp3244 := Call1(__e, SymFn(symtail), let__2678)


let__2680 := tmp3244
_ = let__2680

tmp3256 := Call2(__e, SymFn(symshen_4hds_a_2), let__2680, MakeInteger(0))


var ifres3245 Obj

if True == tmp3256 {
tmp3246 := Call1(__e, SymFn(symtail), let__2680)


let__2681 := tmp3246
_ = let__2681

tmp3247 := Call1(__e, SymFn(sym_5end_6), let__2681)


let__2682 := tmp3247
_ = let__2682

tmp3254 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2682)


var ifres3248 Obj

if True == tmp3254 {
tmp3249 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3248 = tmp3249


} else {
tmp3250 := Call1(__e, SymFn(symshen_4in_1_6), let__2682)


let__2683 := tmp3250
_ = let__2683

tmp3251 := GCons(__e, MakeInteger(0), Nil)

tmp3252 := GCons(__e, symvector, tmp3251)

tmp3253 := Call2(__e, SymFn(symshen_4comb), let__2679, tmp3252)


ifres3248 = tmp3253


}

ifres3245 = ifres3248


} else {
tmp3255 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3245 = tmp3255


}

ifres3243 = ifres3245


} else {
tmp3257 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3243 = tmp3257


}

ifres3240 = ifres3243


} else {
tmp3259 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3240 = tmp3259


}

let__2677 := ifres3240
_ = let__2677

tmp3283 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2677)


if True == tmp3283 {
tmp3270 := GConsP(__e, V674)

var ifres3261 Obj

if True == tmp3270 {
tmp3262 := Call1(__e, SymFn(symhead), V674)


let__2685 := tmp3262
_ = let__2685

tmp3263 := Call1(__e, SymFn(symtail), V674)


let__2686 := tmp3263
_ = let__2686

tmp3268 := GConsP(__e, let__2685)

var ifres3264 Obj

if True == tmp3268 {
tmp3265 := Call1(__e, SymFn(symshen_4compound_1pattern), let__2685)


tmp3266 := Call2(__e, SymFn(symshen_4comb), let__2686, tmp3265)


ifres3264 = tmp3266


} else {
tmp3267 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3264 = tmp3267


}

ifres3261 = ifres3264


} else {
tmp3269 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3261 = tmp3269


}

let__2684 := ifres3261
_ = let__2684

tmp3281 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2684)


if True == tmp3281 {
tmp3271 := Call1(__e, SymFn(symshen_4_5simple_1pattern_6), V674)


let__2688 := tmp3271
_ = let__2688

tmp3277 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2688)


var ifres3272 Obj

if True == tmp3277 {
tmp3273 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3272 = tmp3273


} else {
tmp3274 := Call1(__e, SymFn(symshen_4_5_1out), let__2688)


let__2689 := tmp3274
_ = let__2689

tmp3275 := Call1(__e, SymFn(symshen_4in_1_6), let__2688)


let__2690 := tmp3275
_ = let__2690

tmp3276 := Call2(__e, SymFn(symshen_4comb), let__2690, let__2689)


ifres3272 = tmp3276


}

let__2687 := ifres3272
_ = let__2687

tmp3279 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2687)


if True == tmp3279 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2687
}


} else {
return let__2684
}


} else {
return let__2677
}


panic("unreachable")
})

tmp3284 := Call2(__e, ns2_1set, symshen_4_5pattern_6, tmp3239)


_ = tmp3284

tmp3285 := MakeNative1(func(__e *ControlFlow, V689 Obj) Obj {
tmp3294 := GConsP(__e, V689)

var ifres3286 Obj

if True == tmp3294 {
tmp3287 := Call1(__e, SymFn(symhead), V689)


let__2692 := tmp3287
_ = let__2692

tmp3288 := Call1(__e, SymFn(symtail), V689)


let__2693 := tmp3288
_ = let__2693

tmp3292 := Call1(__e, SymFn(symshen_4constructor_2), let__2692)


var ifres3289 Obj

if True == tmp3292 {
tmp3290 := Call2(__e, SymFn(symshen_4comb), let__2693, let__2692)


ifres3289 = tmp3290


} else {
tmp3291 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3289 = tmp3291


}

ifres3286 = ifres3289


} else {
tmp3293 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3286 = tmp3293


}

let__2691 := ifres3286
_ = let__2691

tmp3296 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2691)


if True == tmp3296 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2691
}


panic("unreachable")
})

tmp3297 := Call2(__e, ns2_1set, symshen_4_5constructor_6, tmp3285)


_ = tmp3297

tmp3298 := MakeNative1(func(__e *ControlFlow, V693 Obj) Obj {
tmp3299 := GCons(__e, sym_8v, Nil)

tmp3300 := GCons(__e, sym_8s, tmp3299)

tmp3301 := GCons(__e, sym_8p, tmp3300)

tmp3302 := GCons(__e, symcons, tmp3301)

return TailCall2(__e, SymFn(symelement_2), V693, tmp3302)


panic("unreachable")
})

tmp3303 := Call2(__e, ns2_1set, symshen_4constructor_2, tmp3298)


_ = tmp3303

tmp3304 := MakeNative1(func(__e *ControlFlow, V694 Obj) Obj {
tmp3305 := Call3(__e, SymFn(symshen_4app), V694, MakeString(" is not a legitimate constructor\n"), symshen_4r)


return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp3305)


panic("unreachable")
})

tmp3306 := Call2(__e, ns2_1set, symshen_4constructor_1error, tmp3304)


_ = tmp3306

tmp3307 := MakeNative2(func(__e *ControlFlow, V695, V696 Obj) Obj {
tmp3308 := GValue(__e, symshen_4_dcustom_1pattern_1compiler_d)

let__2694 := tmp3308
_ = let__2694

tmp3311 := GEq(__e, let__2694, False)

if True == tmp3311 {
return TailCall1(__e, SymFn(symthaw), V696)
} else {
tmp3309 := Call1(__e, let__2694, V695)


return TailCall1(__e, tmp3309, V696)


}


panic("unreachable")
})

tmp3312 := Call2(__e, ns2_1set, symshen_4custom_1pattern_1compiler, tmp3307)


_ = tmp3312

tmp3313 := MakeNative1(func(__e *ControlFlow, V698 Obj) Obj {
tmp3314 := GValue(__e, symshen_4_dcustom_1pattern_1reducer_d)

let__2695 := tmp3314
_ = let__2695

tmp3316 := GEq(__e, let__2695, False)

if True == tmp3316 {
return TailCall0(__e, SymFn(symfail))
} else {
return TailCall1(__e, let__2695, V698)
}


panic("unreachable")
})

tmp3317 := Call2(__e, ns2_1set, symshen_4custom_1pattern_1reducer, tmp3313)


_ = tmp3317

tmp3318 := MakeNative1(func(__e *ControlFlow, V700 Obj) Obj {
tmp3319 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, SymFn(symshen_4compound_1pattern_1h), V700)
panic("unreachable")
})

return TailCall2(__e, SymFn(symshen_4custom_1pattern_1compiler), V700, tmp3319)


panic("unreachable")
})

tmp3320 := Call2(__e, ns2_1set, symshen_4compound_1pattern, tmp3318)


_ = tmp3320

tmp3321 := MakeNative1(func(__e *ControlFlow, V701 Obj) Obj {
tmp3352 := GConsP(__e, V701)

var ifres3333 Obj

if True == tmp3352 {
tmp3350 := GTl(__e, V701)

tmp3351 := GConsP(__e, tmp3350)

var ifres3335 Obj

if True == tmp3351 {
tmp3347 := GTl(__e, V701)

tmp3348 := GTl(__e, tmp3347)

tmp3349 := GConsP(__e, tmp3348)

var ifres3337 Obj

if True == tmp3349 {
tmp3343 := GTl(__e, V701)

tmp3344 := GTl(__e, tmp3343)

tmp3345 := GTl(__e, tmp3344)

tmp3346 := GEq(__e, Nil, tmp3345)

var ifres3339 Obj

if True == tmp3346 {
tmp3341 := GHd(__e, V701)

tmp3342 := Call1(__e, SymFn(symshen_4constructor_2), tmp3341)


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
tmp3322 := GHd(__e, V701)

tmp3323 := GTl(__e, V701)

tmp3324 := GHd(__e, tmp3323)

tmp3325 := Call1(__e, SymFn(symshen_4compile_1pattern_1fragment), tmp3324)


tmp3326 := GTl(__e, V701)

tmp3327 := GTl(__e, tmp3326)

tmp3328 := GHd(__e, tmp3327)

tmp3329 := Call1(__e, SymFn(symshen_4compile_1pattern_1fragment), tmp3328)


tmp3330 := GCons(__e, tmp3329, Nil)

tmp3331 := GCons(__e, tmp3325, tmp3330)

return GCons(__e, tmp3322, tmp3331)


} else {
return TailCall1(__e, SymFn(symshen_4constructor_1error), V701)
}


panic("unreachable")
})

tmp3353 := Call2(__e, ns2_1set, symshen_4compound_1pattern_1h, tmp3321)


_ = tmp3353

tmp3354 := MakeNative1(func(__e *ControlFlow, V702 Obj) Obj {
tmp3383 := GConsP(__e, V702)

var ifres3365 Obj

if True == tmp3383 {
tmp3381 := GHd(__e, V702)

tmp3382 := GEq(__e, symvector, tmp3381)

var ifres3367 Obj

if True == tmp3382 {
tmp3379 := GTl(__e, V702)

tmp3380 := GConsP(__e, tmp3379)

var ifres3369 Obj

if True == tmp3380 {
tmp3376 := GTl(__e, V702)

tmp3377 := GHd(__e, tmp3376)

tmp3378 := GEq(__e, MakeInteger(0), tmp3377)

var ifres3371 Obj

if True == tmp3378 {
tmp3373 := GTl(__e, V702)

tmp3374 := GTl(__e, tmp3373)

tmp3375 := GEq(__e, Nil, tmp3374)

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
return V702
} else {
tmp3363 := GConsP(__e, V702)

if True == tmp3363 {
return TailCall1(__e, SymFn(symshen_4compound_1pattern), V702)
} else {
tmp3361 := GEq(__e, V702, sym__)

if True == tmp3361 {
return TailCall1(__e, SymFn(symgensym), symY)
} else {
tmp3356 := GCons(__e, sym_5_1, Nil)

tmp3357 := GCons(__e, sym_1_6, tmp3356)

tmp3358 := Call2(__e, SymFn(symelement_2), V702, tmp3357)


if True == GNot(__e, tmp3358) {
return V702
} else {
return TailCall1(__e, SymFn(symshen_4constructor_1error), V702)
}


}


}


}


panic("unreachable")
})

tmp3384 := Call2(__e, ns2_1set, symshen_4compile_1pattern_1fragment, tmp3354)


_ = tmp3384

tmp3385 := MakeNative1(func(__e *ControlFlow, V707 Obj) Obj {
tmp3411 := GConsP(__e, V707)

var ifres3387 Obj

if True == tmp3411 {
tmp3409 := GHd(__e, V707)

tmp3410 := GEq(__e, sym_8p, tmp3409)

var ifres3389 Obj

if True == tmp3410 {
tmp3407 := GTl(__e, V707)

tmp3408 := GConsP(__e, tmp3407)

var ifres3391 Obj

if True == tmp3408 {
tmp3404 := GTl(__e, V707)

tmp3405 := GHd(__e, tmp3404)

tmp3406 := GEq(__e, symshen_4custom_1pattern, tmp3405)

var ifres3393 Obj

if True == tmp3406 {
tmp3401 := GTl(__e, V707)

tmp3402 := GTl(__e, tmp3401)

tmp3403 := GConsP(__e, tmp3402)

var ifres3395 Obj

if True == tmp3403 {
tmp3397 := GTl(__e, V707)

tmp3398 := GTl(__e, tmp3397)

tmp3399 := GTl(__e, tmp3398)

tmp3400 := GEq(__e, Nil, tmp3399)

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
return True
} else {
return False
}


panic("unreachable")
})

tmp3412 := Call2(__e, ns2_1set, symshen_4custom_1pattern_2, tmp3385)


_ = tmp3412

tmp3413 := MakeNative1(func(__e *ControlFlow, V710 Obj) Obj {
tmp3441 := GConsP(__e, V710)

var ifres3417 Obj

if True == tmp3441 {
tmp3439 := GHd(__e, V710)

tmp3440 := GEq(__e, sym_8p, tmp3439)

var ifres3419 Obj

if True == tmp3440 {
tmp3437 := GTl(__e, V710)

tmp3438 := GConsP(__e, tmp3437)

var ifres3421 Obj

if True == tmp3438 {
tmp3434 := GTl(__e, V710)

tmp3435 := GHd(__e, tmp3434)

tmp3436 := GEq(__e, symshen_4custom_1pattern, tmp3435)

var ifres3423 Obj

if True == tmp3436 {
tmp3431 := GTl(__e, V710)

tmp3432 := GTl(__e, tmp3431)

tmp3433 := GConsP(__e, tmp3432)

var ifres3425 Obj

if True == tmp3433 {
tmp3427 := GTl(__e, V710)

tmp3428 := GTl(__e, tmp3427)

tmp3429 := GTl(__e, tmp3428)

tmp3430 := GEq(__e, Nil, tmp3429)

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
tmp3414 := GTl(__e, V710)

tmp3415 := GTl(__e, tmp3414)

return GHd(__e, tmp3415)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.custom-pattern-body"))
}


panic("unreachable")
})

tmp3442 := Call2(__e, ns2_1set, symshen_4custom_1pattern_1body, tmp3413)


_ = tmp3442

tmp3443 := MakeNative1(func(__e *ControlFlow, V711 Obj) Obj {
tmp3453 := GConsP(__e, V711)

var ifres3444 Obj

if True == tmp3453 {
tmp3445 := Call1(__e, SymFn(symhead), V711)


let__2697 := tmp3445
_ = let__2697

tmp3446 := Call1(__e, SymFn(symtail), V711)


let__2698 := tmp3446
_ = let__2698

tmp3451 := GEq(__e, let__2697, sym__)

var ifres3447 Obj

if True == tmp3451 {
tmp3448 := Call1(__e, SymFn(symgensym), symY)


tmp3449 := Call2(__e, SymFn(symshen_4comb), let__2698, tmp3448)


ifres3447 = tmp3449


} else {
tmp3450 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3447 = tmp3450


}

ifres3444 = ifres3447


} else {
tmp3452 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3444 = tmp3452


}

let__2696 := ifres3444
_ = let__2696

tmp3469 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2696)


if True == tmp3469 {
tmp3465 := GConsP(__e, V711)

var ifres3454 Obj

if True == tmp3465 {
tmp3455 := Call1(__e, SymFn(symhead), V711)


let__2700 := tmp3455
_ = let__2700

tmp3456 := Call1(__e, SymFn(symtail), V711)


let__2701 := tmp3456
_ = let__2701

tmp3460 := GCons(__e, sym_5_1, Nil)

tmp3461 := GCons(__e, sym_1_6, tmp3460)

tmp3462 := Call2(__e, SymFn(symelement_2), let__2700, tmp3461)


tmp3463 := GNot(__e, tmp3462)

var ifres3457 Obj

if True == tmp3463 {
tmp3458 := Call2(__e, SymFn(symshen_4comb), let__2701, let__2700)


ifres3457 = tmp3458


} else {
tmp3459 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3457 = tmp3459


}

ifres3454 = ifres3457


} else {
tmp3464 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3454 = tmp3464


}

let__2699 := ifres3454
_ = let__2699

tmp3467 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2699)


if True == tmp3467 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2699
}


} else {
return let__2696
}


panic("unreachable")
})

tmp3470 := Call2(__e, ns2_1set, symshen_4_5simple_1pattern_6, tmp3443)


_ = tmp3470

tmp3471 := MakeNative1(func(__e *ControlFlow, V718 Obj) Obj {
tmp3472 := Call1(__e, SymFn(symshen_4_5pattern_6), V718)


let__2703 := tmp3472
_ = let__2703

tmp3478 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2703)


var ifres3473 Obj

if True == tmp3478 {
tmp3474 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3473 = tmp3474


} else {
tmp3475 := Call1(__e, SymFn(symshen_4_5_1out), let__2703)


let__2704 := tmp3475
_ = let__2704

tmp3476 := Call1(__e, SymFn(symshen_4in_1_6), let__2703)


let__2705 := tmp3476
_ = let__2705

tmp3477 := Call2(__e, SymFn(symshen_4comb), let__2705, let__2704)


ifres3473 = tmp3477


}

let__2702 := ifres3473
_ = let__2702

tmp3480 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2702)


if True == tmp3480 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2702
}


panic("unreachable")
})

tmp3481 := Call2(__e, ns2_1set, symshen_4_5pattern1_6, tmp3471)


_ = tmp3481

tmp3482 := MakeNative1(func(__e *ControlFlow, V723 Obj) Obj {
tmp3483 := Call1(__e, SymFn(symshen_4_5pattern_6), V723)


let__2707 := tmp3483
_ = let__2707

tmp3489 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2707)


var ifres3484 Obj

if True == tmp3489 {
tmp3485 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3484 = tmp3485


} else {
tmp3486 := Call1(__e, SymFn(symshen_4_5_1out), let__2707)


let__2708 := tmp3486
_ = let__2708

tmp3487 := Call1(__e, SymFn(symshen_4in_1_6), let__2707)


let__2709 := tmp3487
_ = let__2709

tmp3488 := Call2(__e, SymFn(symshen_4comb), let__2709, let__2708)


ifres3484 = tmp3488


}

let__2706 := ifres3484
_ = let__2706

tmp3491 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2706)


if True == tmp3491 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2706
}


panic("unreachable")
})

tmp3492 := Call2(__e, ns2_1set, symshen_4_5pattern2_6, tmp3482)


_ = tmp3492

tmp3493 := MakeNative1(func(__e *ControlFlow, V728 Obj) Obj {
tmp3494 := PrimGuard1(__e, symabsvector, PrimAbsvector, MakeInteger(2))

let__2710 := tmp3494
_ = let__2710

tmp3495 := GVecSet(__e, let__2710, MakeInteger(0), symshen_4printF)

let__2711 := tmp3495
_ = let__2711

tmp3496 := PrimGuard1(__e, symstr, PrimStr, V728)

tmp3497 := Call2(__e, SymFn(sym_8s), tmp3496, MakeString(")"))


tmp3498 := Call2(__e, SymFn(sym_8s), MakeString(" "), tmp3497)


tmp3499 := Call2(__e, SymFn(sym_8s), MakeString("n"), tmp3498)


tmp3500 := Call2(__e, SymFn(sym_8s), MakeString("f"), tmp3499)


tmp3501 := Call2(__e, SymFn(sym_8s), MakeString("("), tmp3500)


tmp3502 := GVecSet(__e, let__2711, MakeInteger(1), tmp3501)

let__2712 := tmp3502
_ = let__2712

return let__2712


panic("unreachable")
})

tmp3503 := Call2(__e, ns2_1set, symshen_4fn_1print, tmp3493)


_ = tmp3503

tmp3504 := MakeNative1(func(__e *ControlFlow, V732 Obj) Obj {
return GVecGet(__e, V732, MakeInteger(1))
panic("unreachable")
})

tmp3505 := Call2(__e, ns2_1set, symshen_4printF, tmp3504)


_ = tmp3505

tmp3506 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V737 := __self1
_ = V737
V738 := __self2
_ = V738
tmp3530 := GConsP(__e, V738)

var ifres3526 Obj

if True == tmp3530 {
tmp3528 := GTl(__e, V738)

tmp3529 := GEq(__e, Nil, tmp3528)

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
tmp3507 := GHd(__e, V738)

return TailCall1(__e, SymFn(symlength), tmp3507)


} else {
tmp3524 := GConsP(__e, V738)

var ifres3512 Obj

if True == tmp3524 {
tmp3522 := GTl(__e, V738)

tmp3523 := GConsP(__e, tmp3522)

var ifres3514 Obj

if True == tmp3523 {
tmp3516 := GHd(__e, V738)

tmp3517 := Call1(__e, SymFn(symlength), tmp3516)


tmp3518 := GTl(__e, V738)

tmp3519 := GHd(__e, tmp3518)

tmp3520 := Call1(__e, SymFn(symlength), tmp3519)


tmp3521 := GEq(__e, tmp3517, tmp3520)

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
tmp3508 := GTl(__e, V738)

if SymFn(symshen_4arity_1chk) == __self {
__self1, __self2 = V737, tmp3508
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4arity_1chk), V737, tmp3508)


} else {
tmp3509 := Call3(__e, SymFn(symshen_4app), V737, MakeString("\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("arity error in "), tmp3509))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("arity error in "), tmp3509)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


}


panic("unreachable")
})
return __self
})()

tmp3531 := Call2(__e, ns2_1set, symshen_4arity_1chk, tmp3506)


_ = tmp3531

tmp3532 := MakeNative2(func(__e *ControlFlow, V739, V740 Obj) Obj {
tmp3538 := Call1(__e, SymFn(symtuple_2), V740)


if True == tmp3538 {
tmp3533 := Call1(__e, SymFn(symfst), V740)


tmp3534 := Call1(__e, SymFn(symshen_4extract_1vars), tmp3533)


tmp3535 := Call1(__e, SymFn(symsnd), V740)


tmp3536 := Call2(__e, SymFn(symshen_4find_1free_1vars), tmp3534, tmp3535)


return TailCall2(__e, SymFn(symshen_4free_1variable_1error_1message), V739, tmp3536)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4free_1var_1chk)
}


panic("unreachable")
})

tmp3539 := Call2(__e, ns2_1set, symshen_4free_1var_1chk, tmp3532)


_ = tmp3539

tmp3540 := MakeNative2(func(__e *ControlFlow, V741, V742 Obj) Obj {
tmp3552 := Call1(__e, SymFn(symempty_2), V742)


if True == tmp3552 {
return symshen_4skip
} else {
tmp3541 := Call3(__e, SymFn(symshen_4app), V741, MakeString(":"), symshen_4a)


tmp3542 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("free variables in "), tmp3541)

tmp3543 := Call0(__e, SymFn(symstoutput))


tmp3544 := Call2(__e, SymFn(sympr), tmp3542, tmp3543)


_ = tmp3544

tmp3545 := MakeNative1(func(__e *ControlFlow, Z743 Obj) Obj {
tmp3546 := Call3(__e, SymFn(symshen_4app), Z743, MakeString(""), symshen_4a)


tmp3547 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" "), tmp3546)

tmp3548 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp3547, tmp3548)


panic("unreachable")
})

tmp3549 := Call2(__e, SymFn(symshen_4for_1each), tmp3545, V742)


_ = tmp3549

tmp3550 := Call1(__e, SymFn(symnl), MakeInteger(1))


_ = tmp3550

return TailCall0(__e, SymFn(symabort))


}


panic("unreachable")
})

tmp3553 := Call2(__e, ns2_1set, symshen_4free_1variable_1error_1message, tmp3540)


_ = tmp3553

tmp3554 := MakeNative1(func(__e *ControlFlow, V746 Obj) Obj {
tmp3562 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V746)

if True == tmp3562 {
return GCons(__e, V746, Nil)
} else {
tmp3560 := GConsP(__e, V746)

if True == tmp3560 {
tmp3555 := GHd(__e, V746)

tmp3556 := Call1(__e, SymFn(symshen_4extract_1vars), tmp3555)


tmp3557 := GTl(__e, V746)

tmp3558 := Call1(__e, SymFn(symshen_4extract_1vars), tmp3557)


return TailCall2(__e, SymFn(symunion), tmp3556, tmp3558)


} else {
return Nil
}


}


panic("unreachable")
})

tmp3563 := Call2(__e, ns2_1set, symshen_4extract_1vars, tmp3554)


_ = tmp3563

tmp3564 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V751 := __self1
_ = V751
V752 := __self2
_ = V752
tmp3654 := GConsP(__e, V752)

var ifres3641 Obj

if True == tmp3654 {
tmp3652 := GHd(__e, V752)

tmp3653 := GEq(__e, symprotect, tmp3652)

var ifres3643 Obj

if True == tmp3653 {
tmp3650 := GTl(__e, V752)

tmp3651 := GConsP(__e, tmp3650)

var ifres3645 Obj

if True == tmp3651 {
tmp3647 := GTl(__e, V752)

tmp3648 := GTl(__e, tmp3647)

tmp3649 := GEq(__e, Nil, tmp3648)

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
return Nil
} else {
tmp3639 := GConsP(__e, V752)

var ifres3613 Obj

if True == tmp3639 {
tmp3637 := GHd(__e, V752)

tmp3638 := GEq(__e, symlet, tmp3637)

var ifres3615 Obj

if True == tmp3638 {
tmp3635 := GTl(__e, V752)

tmp3636 := GConsP(__e, tmp3635)

var ifres3617 Obj

if True == tmp3636 {
tmp3632 := GTl(__e, V752)

tmp3633 := GTl(__e, tmp3632)

tmp3634 := GConsP(__e, tmp3633)

var ifres3619 Obj

if True == tmp3634 {
tmp3628 := GTl(__e, V752)

tmp3629 := GTl(__e, tmp3628)

tmp3630 := GTl(__e, tmp3629)

tmp3631 := GConsP(__e, tmp3630)

var ifres3621 Obj

if True == tmp3631 {
tmp3623 := GTl(__e, V752)

tmp3624 := GTl(__e, tmp3623)

tmp3625 := GTl(__e, tmp3624)

tmp3626 := GTl(__e, tmp3625)

tmp3627 := GEq(__e, Nil, tmp3626)

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
tmp3565 := GTl(__e, V752)

tmp3566 := GTl(__e, tmp3565)

tmp3567 := GHd(__e, tmp3566)

tmp3568 := Call2(__e, SymFn(symshen_4find_1free_1vars), V751, tmp3567)


tmp3569 := GTl(__e, V752)

tmp3570 := GHd(__e, tmp3569)

tmp3571 := GCons(__e, tmp3570, V751)

tmp3572 := GTl(__e, V752)

tmp3573 := GTl(__e, tmp3572)

tmp3574 := GTl(__e, tmp3573)

tmp3575 := GHd(__e, tmp3574)

tmp3576 := Call2(__e, SymFn(symshen_4find_1free_1vars), tmp3571, tmp3575)


return TailCall2(__e, SymFn(symunion), tmp3568, tmp3576)


} else {
tmp3611 := GConsP(__e, V752)

var ifres3592 Obj

if True == tmp3611 {
tmp3609 := GHd(__e, V752)

tmp3610 := GEq(__e, symlambda, tmp3609)

var ifres3594 Obj

if True == tmp3610 {
tmp3607 := GTl(__e, V752)

tmp3608 := GConsP(__e, tmp3607)

var ifres3596 Obj

if True == tmp3608 {
tmp3604 := GTl(__e, V752)

tmp3605 := GTl(__e, tmp3604)

tmp3606 := GConsP(__e, tmp3605)

var ifres3598 Obj

if True == tmp3606 {
tmp3600 := GTl(__e, V752)

tmp3601 := GTl(__e, tmp3600)

tmp3602 := GTl(__e, tmp3601)

tmp3603 := GEq(__e, Nil, tmp3602)

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
tmp3577 := GTl(__e, V752)

tmp3578 := GHd(__e, tmp3577)

tmp3579 := GCons(__e, tmp3578, V751)

tmp3580 := GTl(__e, V752)

tmp3581 := GTl(__e, tmp3580)

tmp3582 := GHd(__e, tmp3581)

if SymFn(symshen_4find_1free_1vars) == __self {
__self1, __self2 = tmp3579, tmp3582
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4find_1free_1vars), tmp3579, tmp3582)


} else {
tmp3590 := GConsP(__e, V752)

if True == tmp3590 {
tmp3583 := GHd(__e, V752)

tmp3584 := Call2(__e, SymFn(symshen_4find_1free_1vars), V751, tmp3583)


tmp3585 := GTl(__e, V752)

tmp3586 := Call2(__e, SymFn(symshen_4find_1free_1vars), V751, tmp3585)


return TailCall2(__e, SymFn(symunion), tmp3584, tmp3586)


} else {
tmp3588 := Call2(__e, SymFn(symshen_4free_1variable_2), V752, V751)


if True == tmp3588 {
return GCons(__e, V752, Nil)
} else {
return Nil
}


}


}


}


}


panic("unreachable")
})
return __self
})()

tmp3655 := Call2(__e, ns2_1set, symshen_4find_1free_1vars, tmp3564)


_ = tmp3655

tmp3656 := MakeNative2(func(__e *ControlFlow, V753, V754 Obj) Obj {
tmp3661 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V753)

if True == tmp3661 {
tmp3658 := Call2(__e, SymFn(symelement_2), V753, V754)


if True == GNot(__e, tmp3658) {
return True
} else {
return False
}


} else {
return False
}


panic("unreachable")
})

tmp3662 := Call2(__e, ns2_1set, symshen_4free_1variable_2, tmp3656)


_ = tmp3662

tmp3663 := MakeNative2(func(__e *ControlFlow, V755, V756 Obj) Obj {
tmp3664 := GValue(__e, symshen_4_duserdefs_d)

tmp3665 := Call2(__e, SymFn(symadjoin), V755, tmp3664)


tmp3666 := PrimGuard2(__e, symset, PrimSet, symshen_4_duserdefs_d, tmp3665)

_ = tmp3666

tmp3667 := GValue(__e, sym_dproperty_1vector_d)

return TailCall4(__e, SymFn(symput), V755, symshen_4source, V756, tmp3667)


panic("unreachable")
})

tmp3668 := Call2(__e, ns2_1set, symshen_4record_1kl, tmp3663)


_ = tmp3668

tmp3669 := MakeNative3(func(__e *ControlFlow, V757, V758, V759 Obj) Obj {
tmp3670 := Call1(__e, SymFn(symshen_4parameters), V759)


let__2713 := tmp3670
_ = let__2713

tmp3671 := Call2(__e, SymFn(symshen_4kl_1body), V758, let__2713)


tmp3672 := Call2(__e, SymFn(symshen_4scan_1body), V757, tmp3671)


let__2714 := tmp3672
_ = let__2714

tmp3673 := Call1(__e, SymFn(symshen_4cond_1form), let__2714)


tmp3674 := GCons(__e, tmp3673, Nil)

tmp3675 := GCons(__e, let__2713, tmp3674)

tmp3676 := GCons(__e, V757, tmp3675)

tmp3677 := GCons(__e, symdefun, tmp3676)

let__2715 := tmp3677
_ = let__2715

return let__2715


panic("unreachable")
})

tmp3678 := Call2(__e, ns2_1set, symshen_4compile_1to_1kl, tmp3669)


_ = tmp3678

tmp3679 := MakeNative1(func(__e *ControlFlow, V763 Obj) Obj {
tmp3684 := GEq(__e, MakeInteger(0), V763)

if True == tmp3684 {
return Nil
} else {
tmp3680 := Call1(__e, SymFn(symgensym), symV)


tmp3682 := Call1(__e, SymFn(symshen_4parameters), GSub(__e, V763, MakeInteger(1)))


return GCons(__e, tmp3680, tmp3682)


}


panic("unreachable")
})

tmp3685 := Call2(__e, ns2_1set, symshen_4parameters, tmp3679)


_ = tmp3685

tmp3686 := MakeNative1(func(__e *ControlFlow, V766 Obj) Obj {
tmp3710 := GConsP(__e, V766)

var ifres3690 Obj

if True == tmp3710 {
tmp3708 := GHd(__e, V766)

tmp3709 := GConsP(__e, tmp3708)

var ifres3692 Obj

if True == tmp3709 {
tmp3705 := GHd(__e, V766)

tmp3706 := GHd(__e, tmp3705)

tmp3707 := GEq(__e, True, tmp3706)

var ifres3694 Obj

if True == tmp3707 {
tmp3702 := GHd(__e, V766)

tmp3703 := GTl(__e, tmp3702)

tmp3704 := GConsP(__e, tmp3703)

var ifres3696 Obj

if True == tmp3704 {
tmp3698 := GHd(__e, V766)

tmp3699 := GTl(__e, tmp3698)

tmp3700 := GTl(__e, tmp3699)

tmp3701 := GEq(__e, Nil, tmp3700)

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
tmp3687 := GHd(__e, V766)

tmp3688 := GTl(__e, tmp3687)

return GHd(__e, tmp3688)


} else {
return GCons(__e, symcond, V766)
}


panic("unreachable")
})

tmp3711 := Call2(__e, ns2_1set, symshen_4cond_1form, tmp3686)


_ = tmp3711

tmp3712 := MakeNative2(func(__e *ControlFlow, V775, V776 Obj) Obj {
tmp3756 := GEq(__e, Nil, V776)

if True == tmp3756 {
tmp3713 := GCons(__e, V775, Nil)

tmp3714 := GCons(__e, symshen_4f_1error, tmp3713)

tmp3715 := GCons(__e, tmp3714, Nil)

tmp3716 := GCons(__e, True, tmp3715)

return GCons(__e, tmp3716, Nil)


} else {
tmp3754 := GConsP(__e, V776)

var ifres3750 Obj

if True == tmp3754 {
tmp3752 := GHd(__e, V776)

tmp3753 := Call1(__e, SymFn(symshen_4choicepoint_2), tmp3752)


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
tmp3717 := Call1(__e, SymFn(symgensym), symFreeze)


tmp3718 := Call1(__e, SymFn(symgensym), symResult)


tmp3719 := GHd(__e, V776)

tmp3720 := GTl(__e, V776)

return TailCall5(__e, SymFn(symshen_4choicepoint), V775, tmp3717, tmp3718, tmp3719, tmp3720)


} else {
tmp3748 := GConsP(__e, V776)

var ifres3728 Obj

if True == tmp3748 {
tmp3746 := GHd(__e, V776)

tmp3747 := GConsP(__e, tmp3746)

var ifres3730 Obj

if True == tmp3747 {
tmp3743 := GHd(__e, V776)

tmp3744 := GHd(__e, tmp3743)

tmp3745 := GEq(__e, True, tmp3744)

var ifres3732 Obj

if True == tmp3745 {
tmp3740 := GHd(__e, V776)

tmp3741 := GTl(__e, tmp3740)

tmp3742 := GConsP(__e, tmp3741)

var ifres3734 Obj

if True == tmp3742 {
tmp3736 := GHd(__e, V776)

tmp3737 := GTl(__e, tmp3736)

tmp3738 := GTl(__e, tmp3737)

tmp3739 := GEq(__e, Nil, tmp3738)

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
tmp3721 := GHd(__e, V776)

return GCons(__e, tmp3721, Nil)


} else {
tmp3726 := GConsP(__e, V776)

if True == tmp3726 {
tmp3722 := GHd(__e, V776)

tmp3723 := GTl(__e, V776)

tmp3724 := Call2(__e, SymFn(symshen_4scan_1body), V775, tmp3723)


return GCons(__e, tmp3722, tmp3724)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.scan-body"))
}


}


}


}


panic("unreachable")
})

tmp3757 := Call2(__e, ns2_1set, symshen_4scan_1body, tmp3712)


_ = tmp3757

tmp3758 := MakeNative1(func(__e *ControlFlow, V783 Obj) Obj {
tmp3793 := GConsP(__e, V783)

var ifres3760 Obj

if True == tmp3793 {
tmp3791 := GTl(__e, V783)

tmp3792 := GConsP(__e, tmp3791)

var ifres3762 Obj

if True == tmp3792 {
tmp3788 := GTl(__e, V783)

tmp3789 := GHd(__e, tmp3788)

tmp3790 := GConsP(__e, tmp3789)

var ifres3764 Obj

if True == tmp3790 {
tmp3784 := GTl(__e, V783)

tmp3785 := GHd(__e, tmp3784)

tmp3786 := GHd(__e, tmp3785)

tmp3787 := GEq(__e, symshen_4choicepoint_b, tmp3786)

var ifres3766 Obj

if True == tmp3787 {
tmp3780 := GTl(__e, V783)

tmp3781 := GHd(__e, tmp3780)

tmp3782 := GTl(__e, tmp3781)

tmp3783 := GConsP(__e, tmp3782)

var ifres3768 Obj

if True == tmp3783 {
tmp3775 := GTl(__e, V783)

tmp3776 := GHd(__e, tmp3775)

tmp3777 := GTl(__e, tmp3776)

tmp3778 := GTl(__e, tmp3777)

tmp3779 := GEq(__e, Nil, tmp3778)

var ifres3770 Obj

if True == tmp3779 {
tmp3772 := GTl(__e, V783)

tmp3773 := GTl(__e, tmp3772)

tmp3774 := GEq(__e, Nil, tmp3773)

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
return True
} else {
return False
}


panic("unreachable")
})

tmp3794 := Call2(__e, ns2_1set, symshen_4choicepoint_2, tmp3758)


_ = tmp3794

tmp3795 := MakeNative5(func(__e *ControlFlow, V799, V800, V801, V802, V803 Obj) Obj {
tmp3987 := GConsP(__e, V802)

var ifres3909 Obj

if True == tmp3987 {
tmp3985 := GTl(__e, V802)

tmp3986 := GConsP(__e, tmp3985)

var ifres3911 Obj

if True == tmp3986 {
tmp3982 := GTl(__e, V802)

tmp3983 := GHd(__e, tmp3982)

tmp3984 := GConsP(__e, tmp3983)

var ifres3913 Obj

if True == tmp3984 {
tmp3978 := GTl(__e, V802)

tmp3979 := GHd(__e, tmp3978)

tmp3980 := GTl(__e, tmp3979)

tmp3981 := GConsP(__e, tmp3980)

var ifres3915 Obj

if True == tmp3981 {
tmp3973 := GTl(__e, V802)

tmp3974 := GHd(__e, tmp3973)

tmp3975 := GTl(__e, tmp3974)

tmp3976 := GHd(__e, tmp3975)

tmp3977 := GConsP(__e, tmp3976)

var ifres3917 Obj

if True == tmp3977 {
tmp3967 := GTl(__e, V802)

tmp3968 := GHd(__e, tmp3967)

tmp3969 := GTl(__e, tmp3968)

tmp3970 := GHd(__e, tmp3969)

tmp3971 := GHd(__e, tmp3970)

tmp3972 := GEq(__e, symfail_1if, tmp3971)

var ifres3919 Obj

if True == tmp3972 {
tmp3961 := GTl(__e, V802)

tmp3962 := GHd(__e, tmp3961)

tmp3963 := GTl(__e, tmp3962)

tmp3964 := GHd(__e, tmp3963)

tmp3965 := GTl(__e, tmp3964)

tmp3966 := GConsP(__e, tmp3965)

var ifres3921 Obj

if True == tmp3966 {
tmp3954 := GTl(__e, V802)

tmp3955 := GHd(__e, tmp3954)

tmp3956 := GTl(__e, tmp3955)

tmp3957 := GHd(__e, tmp3956)

tmp3958 := GTl(__e, tmp3957)

tmp3959 := GTl(__e, tmp3958)

tmp3960 := GConsP(__e, tmp3959)

var ifres3923 Obj

if True == tmp3960 {
tmp3946 := GTl(__e, V802)

tmp3947 := GHd(__e, tmp3946)

tmp3948 := GTl(__e, tmp3947)

tmp3949 := GHd(__e, tmp3948)

tmp3950 := GTl(__e, tmp3949)

tmp3951 := GTl(__e, tmp3950)

tmp3952 := GTl(__e, tmp3951)

tmp3953 := GEq(__e, Nil, tmp3952)

var ifres3925 Obj

if True == tmp3953 {
tmp3941 := GTl(__e, V802)

tmp3942 := GHd(__e, tmp3941)

tmp3943 := GTl(__e, tmp3942)

tmp3944 := GTl(__e, tmp3943)

tmp3945 := GEq(__e, Nil, tmp3944)

var ifres3927 Obj

if True == tmp3945 {
tmp3938 := GTl(__e, V802)

tmp3939 := GTl(__e, tmp3938)

tmp3940 := GEq(__e, Nil, tmp3939)

var ifres3929 Obj

if True == tmp3940 {
tmp3931 := GTl(__e, V802)

tmp3932 := GHd(__e, tmp3931)

tmp3933 := GTl(__e, tmp3932)

tmp3934 := GHd(__e, tmp3933)

tmp3935 := GTl(__e, tmp3934)

tmp3936 := GHd(__e, tmp3935)

tmp3937 := GEq(__e, V799, tmp3936)

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
tmp3796 := GTl(__e, V802)

tmp3797 := GHd(__e, tmp3796)

tmp3798 := GTl(__e, tmp3797)

tmp3799 := GHd(__e, tmp3798)

tmp3800 := GTl(__e, tmp3799)

tmp3801 := GHd(__e, tmp3800)

tmp3802 := Call2(__e, SymFn(symshen_4scan_1body), tmp3801, V803)


tmp3803 := GCons(__e, symcond, tmp3802)

tmp3804 := GCons(__e, tmp3803, Nil)

tmp3805 := GCons(__e, symfreeze, tmp3804)

tmp3806 := GHd(__e, V802)

tmp3807 := GTl(__e, V802)

tmp3808 := GHd(__e, tmp3807)

tmp3809 := GTl(__e, tmp3808)

tmp3810 := GHd(__e, tmp3809)

tmp3811 := GTl(__e, tmp3810)

tmp3812 := GTl(__e, tmp3811)

tmp3813 := GHd(__e, tmp3812)

tmp3814 := GTl(__e, V802)

tmp3815 := GHd(__e, tmp3814)

tmp3816 := GTl(__e, tmp3815)

tmp3817 := GHd(__e, tmp3816)

tmp3818 := GTl(__e, tmp3817)

tmp3819 := GHd(__e, tmp3818)

tmp3820 := GCons(__e, V801, Nil)

tmp3821 := GCons(__e, tmp3819, tmp3820)

tmp3822 := GCons(__e, V800, Nil)

tmp3823 := GCons(__e, symthaw, tmp3822)

tmp3824 := GCons(__e, V801, Nil)

tmp3825 := GCons(__e, tmp3823, tmp3824)

tmp3826 := GCons(__e, tmp3821, tmp3825)

tmp3827 := GCons(__e, symif, tmp3826)

tmp3828 := GCons(__e, tmp3827, Nil)

tmp3829 := GCons(__e, tmp3813, tmp3828)

tmp3830 := GCons(__e, V801, tmp3829)

tmp3831 := GCons(__e, symlet, tmp3830)

tmp3832 := GCons(__e, V800, Nil)

tmp3833 := GCons(__e, symthaw, tmp3832)

tmp3834 := GCons(__e, tmp3833, Nil)

tmp3835 := GCons(__e, tmp3831, tmp3834)

tmp3836 := GCons(__e, tmp3806, tmp3835)

tmp3837 := GCons(__e, symif, tmp3836)

tmp3838 := GCons(__e, tmp3837, Nil)

tmp3839 := GCons(__e, tmp3805, tmp3838)

tmp3840 := GCons(__e, V800, tmp3839)

tmp3841 := GCons(__e, symlet, tmp3840)

tmp3842 := GCons(__e, tmp3841, Nil)

tmp3843 := GCons(__e, True, tmp3842)

return GCons(__e, tmp3843, Nil)


} else {
tmp3907 := GConsP(__e, V802)

var ifres3880 Obj

if True == tmp3907 {
tmp3905 := GTl(__e, V802)

tmp3906 := GConsP(__e, tmp3905)

var ifres3882 Obj

if True == tmp3906 {
tmp3902 := GTl(__e, V802)

tmp3903 := GHd(__e, tmp3902)

tmp3904 := GConsP(__e, tmp3903)

var ifres3884 Obj

if True == tmp3904 {
tmp3898 := GTl(__e, V802)

tmp3899 := GHd(__e, tmp3898)

tmp3900 := GTl(__e, tmp3899)

tmp3901 := GConsP(__e, tmp3900)

var ifres3886 Obj

if True == tmp3901 {
tmp3893 := GTl(__e, V802)

tmp3894 := GHd(__e, tmp3893)

tmp3895 := GTl(__e, tmp3894)

tmp3896 := GTl(__e, tmp3895)

tmp3897 := GEq(__e, Nil, tmp3896)

var ifres3888 Obj

if True == tmp3897 {
tmp3890 := GTl(__e, V802)

tmp3891 := GTl(__e, tmp3890)

tmp3892 := GEq(__e, Nil, tmp3891)

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
tmp3844 := Call2(__e, SymFn(symshen_4scan_1body), V799, V803)


tmp3845 := GCons(__e, symcond, tmp3844)

tmp3846 := GCons(__e, tmp3845, Nil)

tmp3847 := GCons(__e, symfreeze, tmp3846)

tmp3848 := GHd(__e, V802)

tmp3849 := GTl(__e, V802)

tmp3850 := GHd(__e, tmp3849)

tmp3851 := GTl(__e, tmp3850)

tmp3852 := GHd(__e, tmp3851)

tmp3853 := GCons(__e, symfail, Nil)

tmp3854 := GCons(__e, tmp3853, Nil)

tmp3855 := GCons(__e, V801, tmp3854)

tmp3856 := GCons(__e, sym_a, tmp3855)

tmp3857 := GCons(__e, V800, Nil)

tmp3858 := GCons(__e, symthaw, tmp3857)

tmp3859 := GCons(__e, V801, Nil)

tmp3860 := GCons(__e, tmp3858, tmp3859)

tmp3861 := GCons(__e, tmp3856, tmp3860)

tmp3862 := GCons(__e, symif, tmp3861)

tmp3863 := GCons(__e, tmp3862, Nil)

tmp3864 := GCons(__e, tmp3852, tmp3863)

tmp3865 := GCons(__e, V801, tmp3864)

tmp3866 := GCons(__e, symlet, tmp3865)

tmp3867 := GCons(__e, V800, Nil)

tmp3868 := GCons(__e, symthaw, tmp3867)

tmp3869 := GCons(__e, tmp3868, Nil)

tmp3870 := GCons(__e, tmp3866, tmp3869)

tmp3871 := GCons(__e, tmp3848, tmp3870)

tmp3872 := GCons(__e, symif, tmp3871)

tmp3873 := GCons(__e, tmp3872, Nil)

tmp3874 := GCons(__e, tmp3847, tmp3873)

tmp3875 := GCons(__e, V800, tmp3874)

tmp3876 := GCons(__e, symlet, tmp3875)

tmp3877 := GCons(__e, tmp3876, Nil)

tmp3878 := GCons(__e, True, tmp3877)

return GCons(__e, tmp3878, Nil)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.choicepoint"))
}


}


panic("unreachable")
})

tmp3988 := Call2(__e, ns2_1set, symshen_4choicepoint, tmp3795)


_ = tmp3988

tmp3989 := MakeNative3(func(__e *ControlFlow, V805, V806, V807 Obj) Obj {
tmp4002 := GEq(__e, V805, V807)

if True == tmp4002 {
return V806
} else {
tmp4000 := GConsP(__e, V807)

if True == tmp4000 {
tmp3990 := GHd(__e, V807)

tmp3991 := Call3(__e, SymFn(symshen_4rep_1X), V805, V806, tmp3990)


let__2716 := tmp3991
_ = let__2716

tmp3997 := GHd(__e, V807)

tmp3998 := GEq(__e, let__2716, tmp3997)

if True == tmp3998 {
tmp3992 := GHd(__e, V807)

tmp3993 := GTl(__e, V807)

tmp3994 := Call3(__e, SymFn(symshen_4rep_1X), V805, V806, tmp3993)


return GCons(__e, tmp3992, tmp3994)


} else {
tmp3995 := GTl(__e, V807)

return GCons(__e, let__2716, tmp3995)


}


} else {
return V807
}


}


panic("unreachable")
})

tmp4003 := Call2(__e, ns2_1set, symshen_4rep_1X, tmp3989)


_ = tmp4003

tmp4004 := MakeNative1(func(__e *ControlFlow, V809 Obj) Obj {
tmp4083 := GConsP(__e, V809)

var ifres4064 Obj

if True == tmp4083 {
tmp4081 := GHd(__e, V809)

tmp4082 := GEq(__e, symlambda, tmp4081)

var ifres4066 Obj

if True == tmp4082 {
tmp4079 := GTl(__e, V809)

tmp4080 := GConsP(__e, tmp4079)

var ifres4068 Obj

if True == tmp4080 {
tmp4076 := GTl(__e, V809)

tmp4077 := GTl(__e, tmp4076)

tmp4078 := GConsP(__e, tmp4077)

var ifres4070 Obj

if True == tmp4078 {
tmp4072 := GTl(__e, V809)

tmp4073 := GTl(__e, tmp4072)

tmp4074 := GTl(__e, tmp4073)

tmp4075 := GEq(__e, Nil, tmp4074)

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
tmp4005 := Call1(__e, SymFn(symgensym), symZ)


let__2717 := tmp4005
_ = let__2717

tmp4006 := GTl(__e, V809)

tmp4007 := GHd(__e, tmp4006)

tmp4008 := GTl(__e, V809)

tmp4009 := GTl(__e, tmp4008)

tmp4010 := GHd(__e, tmp4009)

tmp4011 := Call3(__e, SymFn(symshen_4beta), tmp4007, let__2717, tmp4010)


tmp4012 := GCons(__e, tmp4011, Nil)

tmp4013 := GCons(__e, let__2717, tmp4012)

tmp4014 := GCons(__e, symlambda, tmp4013)

let__2718 := tmp4014
_ = let__2718

tmp4015 := MakeNative1(func(__e *ControlFlow, Z812 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4alpha_1convert), Z812)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp4015, let__2718)


} else {
tmp4062 := GConsP(__e, V809)

var ifres4036 Obj

if True == tmp4062 {
tmp4060 := GHd(__e, V809)

tmp4061 := GEq(__e, symlet, tmp4060)

var ifres4038 Obj

if True == tmp4061 {
tmp4058 := GTl(__e, V809)

tmp4059 := GConsP(__e, tmp4058)

var ifres4040 Obj

if True == tmp4059 {
tmp4055 := GTl(__e, V809)

tmp4056 := GTl(__e, tmp4055)

tmp4057 := GConsP(__e, tmp4056)

var ifres4042 Obj

if True == tmp4057 {
tmp4051 := GTl(__e, V809)

tmp4052 := GTl(__e, tmp4051)

tmp4053 := GTl(__e, tmp4052)

tmp4054 := GConsP(__e, tmp4053)

var ifres4044 Obj

if True == tmp4054 {
tmp4046 := GTl(__e, V809)

tmp4047 := GTl(__e, tmp4046)

tmp4048 := GTl(__e, tmp4047)

tmp4049 := GTl(__e, tmp4048)

tmp4050 := GEq(__e, Nil, tmp4049)

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
tmp4016 := Call1(__e, SymFn(symgensym), symW)


let__2719 := tmp4016
_ = let__2719

tmp4017 := GTl(__e, V809)

tmp4018 := GTl(__e, tmp4017)

tmp4019 := GHd(__e, tmp4018)

tmp4020 := GTl(__e, V809)

tmp4021 := GHd(__e, tmp4020)

tmp4022 := GTl(__e, V809)

tmp4023 := GTl(__e, tmp4022)

tmp4024 := GTl(__e, tmp4023)

tmp4025 := GHd(__e, tmp4024)

tmp4026 := Call3(__e, SymFn(symshen_4beta), tmp4021, let__2719, tmp4025)


tmp4027 := GCons(__e, tmp4026, Nil)

tmp4028 := GCons(__e, tmp4019, tmp4027)

tmp4029 := GCons(__e, let__2719, tmp4028)

tmp4030 := GCons(__e, symlet, tmp4029)

let__2720 := tmp4030
_ = let__2720

tmp4031 := MakeNative1(func(__e *ControlFlow, Z815 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4alpha_1convert), Z815)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp4031, let__2720)


} else {
tmp4034 := GConsP(__e, V809)

if True == tmp4034 {
tmp4032 := MakeNative1(func(__e *ControlFlow, Z816 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4alpha_1convert), Z816)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp4032, V809)


} else {
return V809
}


}


}


panic("unreachable")
})

tmp4084 := Call2(__e, ns2_1set, symshen_4alpha_1convert, tmp4004)


_ = tmp4084

tmp4085 := MakeNative2(func(__e *ControlFlow, V817, V818 Obj) Obj {
tmp4086 := MakeNative1(func(__e *ControlFlow, Z819 Obj) Obj {
tmp4087 := Call1(__e, SymFn(symfst), Z819)


tmp4088 := Call1(__e, SymFn(symsnd), Z819)


tmp4089 := Call1(__e, SymFn(symshen_4alpha_1convert), tmp4088)


return TailCall4(__e, SymFn(symshen_4triple_1stack), Nil, tmp4087, V818, tmp4089)


panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp4086, V817)


panic("unreachable")
})

tmp4090 := Call2(__e, ns2_1set, symshen_4kl_1body, tmp4085)


_ = tmp4090

tmp4091 := (func() Obj {
var __self Obj
__self = MakeNative4(func(__e *ControlFlow, __self1, __self2, __self3, __self4 Obj) Obj {
__selftop:
V828 := __self1
_ = V828
V829 := __self2
_ = V829
V830 := __self3
_ = V830
V831 := __self4
_ = V831
tmp4240 := GEq(__e, Nil, V829)

var ifres4215 Obj

if True == tmp4240 {
tmp4239 := GEq(__e, Nil, V830)

var ifres4217 Obj

if True == tmp4239 {
tmp4238 := GConsP(__e, V831)

var ifres4219 Obj

if True == tmp4238 {
tmp4236 := GHd(__e, V831)

tmp4237 := GEq(__e, symwhere, tmp4236)

var ifres4221 Obj

if True == tmp4237 {
tmp4234 := GTl(__e, V831)

tmp4235 := GConsP(__e, tmp4234)

var ifres4223 Obj

if True == tmp4235 {
tmp4231 := GTl(__e, V831)

tmp4232 := GTl(__e, tmp4231)

tmp4233 := GConsP(__e, tmp4232)

var ifres4225 Obj

if True == tmp4233 {
tmp4227 := GTl(__e, V831)

tmp4228 := GTl(__e, tmp4227)

tmp4229 := GTl(__e, tmp4228)

tmp4230 := GEq(__e, Nil, tmp4229)

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
tmp4092 := GTl(__e, V831)

tmp4093 := GHd(__e, tmp4092)

tmp4094 := GCons(__e, tmp4093, V828)

tmp4095 := GTl(__e, V831)

tmp4096 := GTl(__e, tmp4095)

tmp4097 := GHd(__e, tmp4096)

if SymFn(symshen_4triple_1stack) == __self {
__self1, __self2, __self3, __self4 = tmp4094, Nil, Nil, tmp4097
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4triple_1stack), tmp4094, Nil, Nil, tmp4097)


} else {
tmp4213 := GEq(__e, Nil, V829)

var ifres4210 Obj

if True == tmp4213 {
tmp4212 := GEq(__e, Nil, V830)

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
tmp4098 := Call1(__e, SymFn(symreverse), V828)


tmp4099 := Call1(__e, SymFn(symshen_4rectify_1test), tmp4098)


tmp4100 := GCons(__e, V831, Nil)

return GCons(__e, tmp4099, tmp4100)


} else {
tmp4208 := GConsP(__e, V829)

var ifres4201 Obj

if True == tmp4208 {
tmp4207 := GConsP(__e, V830)

var ifres4203 Obj

if True == tmp4207 {
tmp4205 := GHd(__e, V829)

tmp4206 := PrimGuard1(__e, symvariable_2, PrimIsVariable, tmp4205)

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
tmp4101 := GTl(__e, V829)

tmp4102 := GTl(__e, V830)

tmp4103 := GHd(__e, V829)

tmp4104 := GHd(__e, V830)

tmp4105 := Call3(__e, SymFn(symshen_4beta), tmp4103, tmp4104, V831)


if SymFn(symshen_4triple_1stack) == __self {
__self1, __self2, __self3, __self4 = V828, tmp4101, tmp4102, tmp4105
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4triple_1stack), V828, tmp4101, tmp4102, tmp4105)


} else {
tmp4199 := GConsP(__e, V829)

var ifres4192 Obj

if True == tmp4199 {
tmp4198 := GConsP(__e, V830)

var ifres4194 Obj

if True == tmp4198 {
tmp4196 := GHd(__e, V829)

tmp4197 := Call1(__e, SymFn(symshen_4custom_1pattern_2), tmp4196)


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
tmp4106 := GHd(__e, V829)

tmp4107 := Call1(__e, SymFn(symshen_4custom_1pattern_1body), tmp4106)


tmp4108 := GTl(__e, V829)

tmp4109 := GHd(__e, V830)

tmp4110 := GTl(__e, V830)

return TailCall6(__e, SymFn(symshen_4custom_1pattern_1triple_1stack), V828, tmp4107, tmp4108, tmp4109, tmp4110, V831)


} else {
tmp4190 := GConsP(__e, V829)

var ifres4160 Obj

if True == tmp4190 {
tmp4188 := GHd(__e, V829)

tmp4189 := GConsP(__e, tmp4188)

var ifres4162 Obj

if True == tmp4189 {
tmp4185 := GHd(__e, V829)

tmp4186 := GTl(__e, tmp4185)

tmp4187 := GConsP(__e, tmp4186)

var ifres4164 Obj

if True == tmp4187 {
tmp4181 := GHd(__e, V829)

tmp4182 := GTl(__e, tmp4181)

tmp4183 := GTl(__e, tmp4182)

tmp4184 := GConsP(__e, tmp4183)

var ifres4166 Obj

if True == tmp4184 {
tmp4176 := GHd(__e, V829)

tmp4177 := GTl(__e, tmp4176)

tmp4178 := GTl(__e, tmp4177)

tmp4179 := GTl(__e, tmp4178)

tmp4180 := GEq(__e, Nil, tmp4179)

var ifres4168 Obj

if True == tmp4180 {
tmp4175 := GConsP(__e, V830)

var ifres4170 Obj

if True == tmp4175 {
tmp4172 := GHd(__e, V829)

tmp4173 := GHd(__e, tmp4172)

tmp4174 := Call1(__e, SymFn(symshen_4constructor_2), tmp4173)


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
tmp4111 := GHd(__e, V829)

tmp4112 := GHd(__e, tmp4111)

tmp4113 := Call1(__e, SymFn(symshen_4op_1test), tmp4112)


tmp4114 := GHd(__e, V830)

tmp4115 := GCons(__e, tmp4114, Nil)

tmp4116 := GCons(__e, tmp4113, tmp4115)

tmp4117 := GCons(__e, tmp4116, V828)

tmp4118 := GHd(__e, V829)

tmp4119 := GTl(__e, tmp4118)

tmp4120 := GHd(__e, tmp4119)

tmp4121 := GHd(__e, V829)

tmp4122 := GTl(__e, tmp4121)

tmp4123 := GTl(__e, tmp4122)

tmp4124 := GHd(__e, tmp4123)

tmp4125 := GTl(__e, V829)

tmp4126 := GCons(__e, tmp4124, tmp4125)

tmp4127 := GCons(__e, tmp4120, tmp4126)

tmp4128 := GHd(__e, V829)

tmp4129 := GHd(__e, tmp4128)

tmp4130 := Call1(__e, SymFn(symshen_4op1), tmp4129)


tmp4131 := GHd(__e, V830)

tmp4132 := GCons(__e, tmp4131, Nil)

tmp4133 := GCons(__e, tmp4130, tmp4132)

tmp4134 := GHd(__e, V829)

tmp4135 := GHd(__e, tmp4134)

tmp4136 := Call1(__e, SymFn(symshen_4op2), tmp4135)


tmp4137 := GHd(__e, V830)

tmp4138 := GCons(__e, tmp4137, Nil)

tmp4139 := GCons(__e, tmp4136, tmp4138)

tmp4140 := GTl(__e, V830)

tmp4141 := GCons(__e, tmp4139, tmp4140)

tmp4142 := GCons(__e, tmp4133, tmp4141)

tmp4143 := GHd(__e, V829)

tmp4144 := GHd(__e, V830)

tmp4145 := Call3(__e, SymFn(symshen_4beta), tmp4143, tmp4144, V831)


if SymFn(symshen_4triple_1stack) == __self {
__self1, __self2, __self3, __self4 = tmp4117, tmp4127, tmp4142, tmp4145
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4triple_1stack), tmp4117, tmp4127, tmp4142, tmp4145)


} else {
tmp4158 := GConsP(__e, V829)

var ifres4155 Obj

if True == tmp4158 {
tmp4157 := GConsP(__e, V830)

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
tmp4146 := GHd(__e, V829)

tmp4147 := GHd(__e, V830)

tmp4148 := GCons(__e, tmp4147, Nil)

tmp4149 := GCons(__e, tmp4146, tmp4148)

tmp4150 := GCons(__e, sym_a, tmp4149)

tmp4151 := GCons(__e, tmp4150, V828)

tmp4152 := GTl(__e, V829)

tmp4153 := GTl(__e, V830)

if SymFn(symshen_4triple_1stack) == __self {
__self1, __self2, __self3, __self4 = tmp4151, tmp4152, tmp4153, V831
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4triple_1stack), tmp4151, tmp4152, tmp4153, V831)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.triple-stack"))
}


}


}


}


}


}


panic("unreachable")
})
return __self
})()

tmp4241 := Call2(__e, ns2_1set, symshen_4triple_1stack, tmp4091)


_ = tmp4241

tmp4242 := MakeNative6(func(__e *ControlFlow, V832, V833, V834, V835, V836, V837 Obj) Obj {
tmp4243 := Call2(__e, SymFn(sym_8p), V833, V835)


tmp4244 := Call1(__e, SymFn(symshen_4custom_1pattern_1reducer), tmp4243)


let__2721 := tmp4244
_ = let__2721

tmp4257 := Call1(__e, SymFn(symtuple_2), let__2721)


if True == tmp4257 {
tmp4245 := Call1(__e, SymFn(symfst), let__2721)


let__2722 := tmp4245
_ = let__2722

tmp4246 := Call1(__e, SymFn(symsnd), let__2721)


let__2723 := tmp4246
_ = let__2723

tmp4247 := Call1(__e, SymFn(symreverse), let__2722)


tmp4248 := Call2(__e, SymFn(symappend), tmp4247, V832)


tmp4249 := MakeNative1(func(__e *ControlFlow, Z841 Obj) Obj {
return TailCall1(__e, SymFn(symfst), Z841)
panic("unreachable")
})

tmp4250 := Call2(__e, SymFn(symmap), tmp4249, let__2723)


tmp4251 := Call2(__e, SymFn(symappend), tmp4250, V834)


tmp4252 := MakeNative1(func(__e *ControlFlow, Z842 Obj) Obj {
return TailCall1(__e, SymFn(symsnd), Z842)
panic("unreachable")
})

tmp4253 := Call2(__e, SymFn(symmap), tmp4252, let__2723)


tmp4254 := Call2(__e, SymFn(symappend), tmp4253, V836)


tmp4255 := Call3(__e, SymFn(symshen_4beta), V833, V835, V837)


return TailCall4(__e, SymFn(symshen_4triple_1stack), tmp4248, tmp4251, tmp4254, tmp4255)


} else {
return TailCall1(__e, SymFn(symshen_4constructor_1error), V833)
}


panic("unreachable")
})

tmp4258 := Call2(__e, ns2_1set, symshen_4custom_1pattern_1triple_1stack, tmp4242)


_ = tmp4258

tmp4259 := MakeNative1(func(__e *ControlFlow, V845 Obj) Obj {
tmp4278 := GEq(__e, Nil, V845)

if True == tmp4278 {
return True
} else {
tmp4276 := GConsP(__e, V845)

var ifres4272 Obj

if True == tmp4276 {
tmp4274 := GTl(__e, V845)

tmp4275 := GEq(__e, Nil, tmp4274)

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
return GHd(__e, V845)
} else {
tmp4270 := GConsP(__e, V845)

var ifres4266 Obj

if True == tmp4270 {
tmp4268 := GTl(__e, V845)

tmp4269 := GConsP(__e, tmp4268)

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
tmp4260 := GHd(__e, V845)

tmp4261 := GTl(__e, V845)

tmp4262 := Call1(__e, SymFn(symshen_4rectify_1test), tmp4261)


tmp4263 := GCons(__e, tmp4262, Nil)

tmp4264 := GCons(__e, tmp4260, tmp4263)

return GCons(__e, symand, tmp4264)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.rectify-test"))
}


}


}


panic("unreachable")
})

tmp4279 := Call2(__e, ns2_1set, symshen_4rectify_1test, tmp4259)


_ = tmp4279

tmp4280 := MakeNative3(func(__e *ControlFlow, V855, V856, V857 Obj) Obj {
tmp4357 := GEq(__e, V855, V857)

if True == tmp4357 {
return V856
} else {
tmp4355 := GConsP(__e, V857)

var ifres4331 Obj

if True == tmp4355 {
tmp4353 := GHd(__e, V857)

tmp4354 := GEq(__e, symlambda, tmp4353)

var ifres4333 Obj

if True == tmp4354 {
tmp4351 := GTl(__e, V857)

tmp4352 := GConsP(__e, tmp4351)

var ifres4335 Obj

if True == tmp4352 {
tmp4348 := GTl(__e, V857)

tmp4349 := GTl(__e, tmp4348)

tmp4350 := GConsP(__e, tmp4349)

var ifres4337 Obj

if True == tmp4350 {
tmp4344 := GTl(__e, V857)

tmp4345 := GTl(__e, tmp4344)

tmp4346 := GTl(__e, tmp4345)

tmp4347 := GEq(__e, Nil, tmp4346)

var ifres4339 Obj

if True == tmp4347 {
tmp4341 := GTl(__e, V857)

tmp4342 := GHd(__e, tmp4341)

tmp4343 := GEq(__e, V855, tmp4342)

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
return V857
} else {
tmp4329 := GConsP(__e, V857)

var ifres4298 Obj

if True == tmp4329 {
tmp4327 := GHd(__e, V857)

tmp4328 := GEq(__e, symlet, tmp4327)

var ifres4300 Obj

if True == tmp4328 {
tmp4325 := GTl(__e, V857)

tmp4326 := GConsP(__e, tmp4325)

var ifres4302 Obj

if True == tmp4326 {
tmp4322 := GTl(__e, V857)

tmp4323 := GTl(__e, tmp4322)

tmp4324 := GConsP(__e, tmp4323)

var ifres4304 Obj

if True == tmp4324 {
tmp4318 := GTl(__e, V857)

tmp4319 := GTl(__e, tmp4318)

tmp4320 := GTl(__e, tmp4319)

tmp4321 := GConsP(__e, tmp4320)

var ifres4306 Obj

if True == tmp4321 {
tmp4313 := GTl(__e, V857)

tmp4314 := GTl(__e, tmp4313)

tmp4315 := GTl(__e, tmp4314)

tmp4316 := GTl(__e, tmp4315)

tmp4317 := GEq(__e, Nil, tmp4316)

var ifres4308 Obj

if True == tmp4317 {
tmp4310 := GTl(__e, V857)

tmp4311 := GHd(__e, tmp4310)

tmp4312 := GEq(__e, V855, tmp4311)

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
tmp4281 := GTl(__e, V857)

tmp4282 := GHd(__e, tmp4281)

tmp4283 := GTl(__e, V857)

tmp4284 := GHd(__e, tmp4283)

tmp4285 := GTl(__e, V857)

tmp4286 := GTl(__e, tmp4285)

tmp4287 := GHd(__e, tmp4286)

tmp4288 := Call3(__e, SymFn(symshen_4beta), tmp4284, V856, tmp4287)


tmp4289 := GTl(__e, V857)

tmp4290 := GTl(__e, tmp4289)

tmp4291 := GTl(__e, tmp4290)

tmp4292 := GCons(__e, tmp4288, tmp4291)

tmp4293 := GCons(__e, tmp4282, tmp4292)

return GCons(__e, symlet, tmp4293)


} else {
tmp4296 := GConsP(__e, V857)

if True == tmp4296 {
tmp4294 := MakeNative1(func(__e *ControlFlow, Z858 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4beta), V855, V856, Z858)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp4294, V857)


} else {
return V857
}


}


}


}


panic("unreachable")
})

tmp4358 := Call2(__e, ns2_1set, symshen_4beta, tmp4280)


_ = tmp4358

tmp4359 := MakeNative1(func(__e *ControlFlow, V861 Obj) Obj {
tmp4367 := GEq(__e, symcons, V861)

if True == tmp4367 {
return symhd
} else {
tmp4365 := GEq(__e, sym_8s, V861)

if True == tmp4365 {
return symhdstr
} else {
tmp4363 := GEq(__e, sym_8p, V861)

if True == tmp4363 {
return symfst
} else {
tmp4361 := GEq(__e, sym_8v, V861)

if True == tmp4361 {
return symhdv
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.op1"))
}


}


}


}


panic("unreachable")
})

tmp4368 := Call2(__e, ns2_1set, symshen_4op1, tmp4359)


_ = tmp4368

tmp4369 := MakeNative1(func(__e *ControlFlow, V864 Obj) Obj {
tmp4377 := GEq(__e, symcons, V864)

if True == tmp4377 {
return symtl
} else {
tmp4375 := GEq(__e, sym_8s, V864)

if True == tmp4375 {
return symtlstr
} else {
tmp4373 := GEq(__e, sym_8p, V864)

if True == tmp4373 {
return symsnd
} else {
tmp4371 := GEq(__e, sym_8v, V864)

if True == tmp4371 {
return symtlv
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.op2"))
}


}


}


}


panic("unreachable")
})

tmp4378 := Call2(__e, ns2_1set, symshen_4op2, tmp4369)


_ = tmp4378

tmp4379 := MakeNative1(func(__e *ControlFlow, V867 Obj) Obj {
tmp4387 := GEq(__e, symcons, V867)

if True == tmp4387 {
return symcons_2
} else {
tmp4385 := GEq(__e, sym_8s, V867)

if True == tmp4385 {
return symshen_4_7string_2
} else {
tmp4383 := GEq(__e, sym_8p, V867)

if True == tmp4383 {
return symtuple_2
} else {
tmp4381 := GEq(__e, sym_8v, V867)

if True == tmp4381 {
return symshen_4_7vector_2
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.op-test"))
}


}


}


}


panic("unreachable")
})

tmp4388 := Call2(__e, ns2_1set, symshen_4op_1test, tmp4379)


_ = tmp4388

tmp4389 := MakeNative1(func(__e *ControlFlow, V868 Obj) Obj {
tmp4391 := GEq(__e, MakeString(""), V868)

if True == tmp4391 {
return False
} else {
return PrimGuard1(__e, symstring_2, PrimIsString, V868)
}


panic("unreachable")
})

tmp4392 := Call2(__e, ns2_1set, symshen_4_7string_2, tmp4389)


_ = tmp4392

tmp4393 := MakeNative1(func(__e *ControlFlow, V869 Obj) Obj {
tmp4395 := Call1(__e, SymFn(symvector), MakeInteger(0))


tmp4396 := GEq(__e, V869, tmp4395)

if True == tmp4396 {
return False
} else {
return TailCall1(__e, SymFn(symvector_2), V869)
}


panic("unreachable")
})

tmp4397 := Call2(__e, ns2_1set, symshen_4_7vector_2, tmp4393)


_ = tmp4397

tmp4398 := MakeNative1(func(__e *ControlFlow, V872 Obj) Obj {
tmp4402 := GEq(__e, sym_7, V872)

if True == tmp4402 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dfactorise_2_d, True)
} else {
tmp4400 := GEq(__e, sym_1, V872)

if True == tmp4400 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dfactorise_2_d, False)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("factorise expects a + or a -\n"))
}


}


panic("unreachable")
})

tmp4403 := Call2(__e, ns2_1set, symfactorise, tmp4398)


_ = tmp4403

tmp4404 := MakeNative1(func(__e *ControlFlow, V873 Obj) Obj {
tmp4406 := GValue(__e, symshen_4_dfactorise_2_d)

if True == tmp4406 {
return TailCall1(__e, SymFn(symshen_4factor), V873)
} else {
return V873
}


panic("unreachable")
})

tmp4407 := Call2(__e, ns2_1set, symshen_4factorise_1code, tmp4404)


_ = tmp4407

tmp4408 := MakeNative1(func(__e *ControlFlow, V874 Obj) Obj {
tmp4465 := GConsP(__e, V874)

var ifres4424 Obj

if True == tmp4465 {
tmp4463 := GHd(__e, V874)

tmp4464 := GEq(__e, symdefun, tmp4463)

var ifres4426 Obj

if True == tmp4464 {
tmp4461 := GTl(__e, V874)

tmp4462 := GConsP(__e, tmp4461)

var ifres4428 Obj

if True == tmp4462 {
tmp4458 := GTl(__e, V874)

tmp4459 := GTl(__e, tmp4458)

tmp4460 := GConsP(__e, tmp4459)

var ifres4430 Obj

if True == tmp4460 {
tmp4454 := GTl(__e, V874)

tmp4455 := GTl(__e, tmp4454)

tmp4456 := GTl(__e, tmp4455)

tmp4457 := GConsP(__e, tmp4456)

var ifres4432 Obj

if True == tmp4457 {
tmp4449 := GTl(__e, V874)

tmp4450 := GTl(__e, tmp4449)

tmp4451 := GTl(__e, tmp4450)

tmp4452 := GHd(__e, tmp4451)

tmp4453 := GConsP(__e, tmp4452)

var ifres4434 Obj

if True == tmp4453 {
tmp4443 := GTl(__e, V874)

tmp4444 := GTl(__e, tmp4443)

tmp4445 := GTl(__e, tmp4444)

tmp4446 := GHd(__e, tmp4445)

tmp4447 := GHd(__e, tmp4446)

tmp4448 := GEq(__e, symcond, tmp4447)

var ifres4436 Obj

if True == tmp4448 {
tmp4438 := GTl(__e, V874)

tmp4439 := GTl(__e, tmp4438)

tmp4440 := GTl(__e, tmp4439)

tmp4441 := GTl(__e, tmp4440)

tmp4442 := GEq(__e, Nil, tmp4441)

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
tmp4409 := GTl(__e, V874)

tmp4410 := GHd(__e, tmp4409)

tmp4411 := GTl(__e, V874)

tmp4412 := GTl(__e, tmp4411)

tmp4413 := GHd(__e, tmp4412)

tmp4414 := GTl(__e, V874)

tmp4415 := GTl(__e, tmp4414)

tmp4416 := GTl(__e, tmp4415)

tmp4417 := GHd(__e, tmp4416)

tmp4418 := GTl(__e, tmp4417)

tmp4419 := Call1(__e, SymFn(symshen_4factor_1recognisors), tmp4418)


tmp4420 := GCons(__e, tmp4419, Nil)

tmp4421 := GCons(__e, tmp4413, tmp4420)

tmp4422 := GCons(__e, tmp4410, tmp4421)

return GCons(__e, symdefun, tmp4422)


} else {
return V874
}


panic("unreachable")
})

tmp4466 := Call2(__e, ns2_1set, symshen_4factor, tmp4408)


_ = tmp4466

tmp4467 := MakeNative1(func(__e *ControlFlow, V877 Obj) Obj {
tmp4617 := GConsP(__e, V877)

var ifres4597 Obj

if True == tmp4617 {
tmp4615 := GHd(__e, V877)

tmp4616 := GConsP(__e, tmp4615)

var ifres4599 Obj

if True == tmp4616 {
tmp4612 := GHd(__e, V877)

tmp4613 := GHd(__e, tmp4612)

tmp4614 := GEq(__e, True, tmp4613)

var ifres4601 Obj

if True == tmp4614 {
tmp4609 := GHd(__e, V877)

tmp4610 := GTl(__e, tmp4609)

tmp4611 := GConsP(__e, tmp4610)

var ifres4603 Obj

if True == tmp4611 {
tmp4605 := GHd(__e, V877)

tmp4606 := GTl(__e, tmp4605)

tmp4607 := GTl(__e, tmp4606)

tmp4608 := GEq(__e, Nil, tmp4607)

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
tmp4468 := GHd(__e, V877)

tmp4469 := GTl(__e, tmp4468)

return GHd(__e, tmp4469)


} else {
tmp4595 := GConsP(__e, V877)

var ifres4548 Obj

if True == tmp4595 {
tmp4593 := GHd(__e, V877)

tmp4594 := GConsP(__e, tmp4593)

var ifres4550 Obj

if True == tmp4594 {
tmp4590 := GHd(__e, V877)

tmp4591 := GHd(__e, tmp4590)

tmp4592 := GConsP(__e, tmp4591)

var ifres4552 Obj

if True == tmp4592 {
tmp4586 := GHd(__e, V877)

tmp4587 := GHd(__e, tmp4586)

tmp4588 := GHd(__e, tmp4587)

tmp4589 := GEq(__e, symand, tmp4588)

var ifres4554 Obj

if True == tmp4589 {
tmp4582 := GHd(__e, V877)

tmp4583 := GHd(__e, tmp4582)

tmp4584 := GTl(__e, tmp4583)

tmp4585 := GConsP(__e, tmp4584)

var ifres4556 Obj

if True == tmp4585 {
tmp4577 := GHd(__e, V877)

tmp4578 := GHd(__e, tmp4577)

tmp4579 := GTl(__e, tmp4578)

tmp4580 := GTl(__e, tmp4579)

tmp4581 := GConsP(__e, tmp4580)

var ifres4558 Obj

if True == tmp4581 {
tmp4571 := GHd(__e, V877)

tmp4572 := GHd(__e, tmp4571)

tmp4573 := GTl(__e, tmp4572)

tmp4574 := GTl(__e, tmp4573)

tmp4575 := GTl(__e, tmp4574)

tmp4576 := GEq(__e, Nil, tmp4575)

var ifres4560 Obj

if True == tmp4576 {
tmp4568 := GHd(__e, V877)

tmp4569 := GTl(__e, tmp4568)

tmp4570 := GConsP(__e, tmp4569)

var ifres4562 Obj

if True == tmp4570 {
tmp4564 := GHd(__e, V877)

tmp4565 := GTl(__e, tmp4564)

tmp4566 := GTl(__e, tmp4565)

tmp4567 := GEq(__e, Nil, tmp4566)

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
tmp4470 := GHd(__e, V877)

tmp4471 := GHd(__e, tmp4470)

tmp4472 := GTl(__e, tmp4471)

tmp4473 := GHd(__e, tmp4472)

tmp4474 := Call3(__e, SymFn(symshen_4pivot_1on), tmp4473, V877, Nil)


let__2724 := tmp4474
_ = let__2724

tmp4475 := Call1(__e, SymFn(symfst), let__2724)


let__2725 := tmp4475
_ = let__2725

tmp4519 := Call1(__e, SymFn(symshen_4bad_1pivot_2), let__2725)


if True == tmp4519 {
tmp4476 := GCons(__e, symif, Nil)

tmp4477 := GHd(__e, V877)

tmp4478 := GHd(__e, tmp4477)

tmp4479 := GHd(__e, V877)

tmp4480 := GTl(__e, tmp4479)

tmp4481 := GHd(__e, tmp4480)

tmp4482 := Call2(__e, SymFn(symshen_4recursively_1factor_1selectors), tmp4478, tmp4481)


tmp4483 := GTl(__e, V877)

tmp4484 := Call1(__e, SymFn(symshen_4factor_1recognisors), tmp4483)


tmp4485 := GCons(__e, tmp4484, Nil)

tmp4486 := Call2(__e, SymFn(symappend), tmp4482, tmp4485)


return TailCall2(__e, SymFn(symappend), tmp4476, tmp4486)


} else {
tmp4487 := Call1(__e, SymFn(symsnd), let__2724)


let__2726 := tmp4487
_ = let__2726

tmp4488 := Call1(__e, SymFn(symshen_4factor_1recognisors), let__2726)


let__2727 := tmp4488
_ = let__2727

tmp4489 := Call1(__e, SymFn(symgensym), symGoTo)


let__2728 := tmp4489
_ = let__2728

tmp4490 := GCons(__e, let__2728, Nil)

tmp4491 := GCons(__e, symthaw, tmp4490)

tmp4492 := GCons(__e, tmp4491, Nil)

tmp4493 := GCons(__e, True, tmp4492)

tmp4494 := GCons(__e, tmp4493, let__2725)

tmp4495 := Call1(__e, SymFn(symreverse), tmp4494)


let__2729 := tmp4495
_ = let__2729

tmp4496 := GCons(__e, let__2727, Nil)

tmp4497 := GCons(__e, symfreeze, tmp4496)

tmp4498 := GHd(__e, V877)

tmp4499 := GHd(__e, tmp4498)

tmp4500 := GTl(__e, tmp4499)

tmp4501 := GHd(__e, tmp4500)

tmp4502 := GHd(__e, V877)

tmp4503 := GHd(__e, tmp4502)

tmp4504 := GTl(__e, tmp4503)

tmp4505 := GHd(__e, tmp4504)

tmp4506 := Call1(__e, SymFn(symshen_4factor_1recognisors), let__2729)


tmp4507 := Call2(__e, SymFn(symshen_4factor_1selectors), tmp4505, tmp4506)


tmp4508 := GCons(__e, let__2728, Nil)

tmp4509 := GCons(__e, symthaw, tmp4508)

tmp4510 := GCons(__e, tmp4509, Nil)

tmp4511 := GCons(__e, tmp4507, tmp4510)

tmp4512 := GCons(__e, tmp4501, tmp4511)

tmp4513 := GCons(__e, symif, tmp4512)

tmp4514 := GCons(__e, tmp4513, Nil)

tmp4515 := GCons(__e, tmp4497, tmp4514)

tmp4516 := GCons(__e, let__2728, tmp4515)

tmp4517 := GCons(__e, symlet, tmp4516)

let__2730 := tmp4517
_ = let__2730

return TailCall1(__e, SymFn(symshen_4remove_1indirection), let__2730)


}


} else {
tmp4546 := GConsP(__e, V877)

var ifres4531 Obj

if True == tmp4546 {
tmp4544 := GHd(__e, V877)

tmp4545 := GConsP(__e, tmp4544)

var ifres4533 Obj

if True == tmp4545 {
tmp4541 := GHd(__e, V877)

tmp4542 := GTl(__e, tmp4541)

tmp4543 := GConsP(__e, tmp4542)

var ifres4535 Obj

if True == tmp4543 {
tmp4537 := GHd(__e, V877)

tmp4538 := GTl(__e, tmp4537)

tmp4539 := GTl(__e, tmp4538)

tmp4540 := GEq(__e, Nil, tmp4539)

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
tmp4520 := GHd(__e, V877)

tmp4521 := GHd(__e, tmp4520)

tmp4522 := GHd(__e, V877)

tmp4523 := GTl(__e, tmp4522)

tmp4524 := GHd(__e, tmp4523)

tmp4525 := GTl(__e, V877)

tmp4526 := Call1(__e, SymFn(symshen_4factor_1recognisors), tmp4525)


tmp4527 := GCons(__e, tmp4526, Nil)

tmp4528 := GCons(__e, tmp4524, tmp4527)

tmp4529 := GCons(__e, tmp4521, tmp4528)

return GCons(__e, symif, tmp4529)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4factor_1recognisors)
}


}


}


panic("unreachable")
})

tmp4618 := Call2(__e, ns2_1set, symshen_4factor_1recognisors, tmp4467)


_ = tmp4618

tmp4619 := MakeNative2(func(__e *ControlFlow, V885, V886 Obj) Obj {
tmp4691 := GConsP(__e, V885)

var ifres4665 Obj

if True == tmp4691 {
tmp4689 := GHd(__e, V885)

tmp4690 := GEq(__e, symlet, tmp4689)

var ifres4667 Obj

if True == tmp4690 {
tmp4687 := GTl(__e, V885)

tmp4688 := GConsP(__e, tmp4687)

var ifres4669 Obj

if True == tmp4688 {
tmp4684 := GTl(__e, V885)

tmp4685 := GTl(__e, tmp4684)

tmp4686 := GConsP(__e, tmp4685)

var ifres4671 Obj

if True == tmp4686 {
tmp4680 := GTl(__e, V885)

tmp4681 := GTl(__e, tmp4680)

tmp4682 := GTl(__e, tmp4681)

tmp4683 := GConsP(__e, tmp4682)

var ifres4673 Obj

if True == tmp4683 {
tmp4675 := GTl(__e, V885)

tmp4676 := GTl(__e, tmp4675)

tmp4677 := GTl(__e, tmp4676)

tmp4678 := GTl(__e, tmp4677)

tmp4679 := GEq(__e, Nil, tmp4678)

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
tmp4620 := GTl(__e, V885)

tmp4621 := GHd(__e, tmp4620)

tmp4622 := GTl(__e, V885)

tmp4623 := GTl(__e, tmp4622)

tmp4624 := GHd(__e, tmp4623)

tmp4625 := GTl(__e, V885)

tmp4626 := GTl(__e, tmp4625)

tmp4627 := GTl(__e, tmp4626)

tmp4628 := GHd(__e, tmp4627)

tmp4629 := Call2(__e, SymFn(symshen_4recursively_1factor_1selectors), tmp4628, V886)


return TailCall3(__e, SymFn(symshen_4restore_1local), tmp4621, tmp4624, tmp4629)


} else {
tmp4663 := GConsP(__e, V885)

var ifres4644 Obj

if True == tmp4663 {
tmp4661 := GHd(__e, V885)

tmp4662 := GEq(__e, symand, tmp4661)

var ifres4646 Obj

if True == tmp4662 {
tmp4659 := GTl(__e, V885)

tmp4660 := GConsP(__e, tmp4659)

var ifres4648 Obj

if True == tmp4660 {
tmp4656 := GTl(__e, V885)

tmp4657 := GTl(__e, tmp4656)

tmp4658 := GConsP(__e, tmp4657)

var ifres4650 Obj

if True == tmp4658 {
tmp4652 := GTl(__e, V885)

tmp4653 := GTl(__e, tmp4652)

tmp4654 := GTl(__e, tmp4653)

tmp4655 := GEq(__e, Nil, tmp4654)

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
tmp4630 := GTl(__e, V885)

tmp4631 := GHd(__e, tmp4630)

tmp4632 := GTl(__e, V885)

tmp4633 := GHd(__e, tmp4632)

tmp4634 := GTl(__e, V885)

tmp4635 := GTl(__e, tmp4634)

tmp4636 := GHd(__e, tmp4635)

tmp4637 := Call2(__e, SymFn(symshen_4factor_1selectors), tmp4633, tmp4636)


tmp4638 := GTl(__e, V885)

tmp4639 := GHd(__e, tmp4638)

tmp4640 := Call2(__e, SymFn(symshen_4factor_1selectors), tmp4639, V886)


tmp4641 := Call2(__e, SymFn(symshen_4recursively_1factor_1selectors), tmp4637, tmp4640)


return TailCall2(__e, SymFn(symshen_4restore_1P), tmp4631, tmp4641)


} else {
tmp4642 := GCons(__e, V886, Nil)

return GCons(__e, V885, tmp4642)


}


}


panic("unreachable")
})

tmp4692 := Call2(__e, ns2_1set, symshen_4recursively_1factor_1selectors, tmp4619)


_ = tmp4692

tmp4693 := MakeNative2(func(__e *ControlFlow, V887, V888 Obj) Obj {
tmp4709 := GConsP(__e, V888)

var ifres4700 Obj

if True == tmp4709 {
tmp4707 := GTl(__e, V888)

tmp4708 := GConsP(__e, tmp4707)

var ifres4702 Obj

if True == tmp4708 {
tmp4704 := GTl(__e, V888)

tmp4705 := GTl(__e, tmp4704)

tmp4706 := GEq(__e, Nil, tmp4705)

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
tmp4694 := GHd(__e, V888)

tmp4695 := GCons(__e, tmp4694, Nil)

tmp4696 := GCons(__e, V887, tmp4695)

tmp4697 := GCons(__e, symand, tmp4696)

tmp4698 := GTl(__e, V888)

return GCons(__e, tmp4697, tmp4698)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4restore_1P)
}


panic("unreachable")
})

tmp4710 := Call2(__e, ns2_1set, symshen_4restore_1P, tmp4693)


_ = tmp4710

tmp4711 := MakeNative3(func(__e *ControlFlow, V889, V890, V891 Obj) Obj {
tmp4728 := GConsP(__e, V891)

var ifres4719 Obj

if True == tmp4728 {
tmp4726 := GTl(__e, V891)

tmp4727 := GConsP(__e, tmp4726)

var ifres4721 Obj

if True == tmp4727 {
tmp4723 := GTl(__e, V891)

tmp4724 := GTl(__e, tmp4723)

tmp4725 := GEq(__e, Nil, tmp4724)

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
tmp4712 := GHd(__e, V891)

tmp4713 := GCons(__e, tmp4712, Nil)

tmp4714 := GCons(__e, V890, tmp4713)

tmp4715 := GCons(__e, V889, tmp4714)

tmp4716 := GCons(__e, symlet, tmp4715)

tmp4717 := GTl(__e, V891)

return GCons(__e, tmp4716, tmp4717)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4restore_1local)
}


panic("unreachable")
})

tmp4729 := Call2(__e, ns2_1set, symshen_4restore_1local, tmp4711)


_ = tmp4729

tmp4730 := MakeNative1(func(__e *ControlFlow, V896 Obj) Obj {
tmp4736 := GConsP(__e, V896)

var ifres4732 Obj

if True == tmp4736 {
tmp4734 := GTl(__e, V896)

tmp4735 := GEq(__e, Nil, tmp4734)

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
return True
} else {
return False
}


panic("unreachable")
})

tmp4737 := Call2(__e, ns2_1set, symshen_4bad_1pivot_2, tmp4730)


_ = tmp4737

tmp4738 := MakeNative1(func(__e *ControlFlow, V897 Obj) Obj {
tmp4853 := GConsP(__e, V897)

var ifres4753 Obj

if True == tmp4853 {
tmp4851 := GHd(__e, V897)

tmp4852 := GEq(__e, symlet, tmp4851)

var ifres4755 Obj

if True == tmp4852 {
tmp4849 := GTl(__e, V897)

tmp4850 := GConsP(__e, tmp4849)

var ifres4757 Obj

if True == tmp4850 {
tmp4846 := GTl(__e, V897)

tmp4847 := GTl(__e, tmp4846)

tmp4848 := GConsP(__e, tmp4847)

var ifres4759 Obj

if True == tmp4848 {
tmp4842 := GTl(__e, V897)

tmp4843 := GTl(__e, tmp4842)

tmp4844 := GHd(__e, tmp4843)

tmp4845 := GConsP(__e, tmp4844)

var ifres4761 Obj

if True == tmp4845 {
tmp4837 := GTl(__e, V897)

tmp4838 := GTl(__e, tmp4837)

tmp4839 := GHd(__e, tmp4838)

tmp4840 := GHd(__e, tmp4839)

tmp4841 := GEq(__e, symfreeze, tmp4840)

var ifres4763 Obj

if True == tmp4841 {
tmp4832 := GTl(__e, V897)

tmp4833 := GTl(__e, tmp4832)

tmp4834 := GHd(__e, tmp4833)

tmp4835 := GTl(__e, tmp4834)

tmp4836 := GConsP(__e, tmp4835)

var ifres4765 Obj

if True == tmp4836 {
tmp4826 := GTl(__e, V897)

tmp4827 := GTl(__e, tmp4826)

tmp4828 := GHd(__e, tmp4827)

tmp4829 := GTl(__e, tmp4828)

tmp4830 := GHd(__e, tmp4829)

tmp4831 := GConsP(__e, tmp4830)

var ifres4767 Obj

if True == tmp4831 {
tmp4819 := GTl(__e, V897)

tmp4820 := GTl(__e, tmp4819)

tmp4821 := GHd(__e, tmp4820)

tmp4822 := GTl(__e, tmp4821)

tmp4823 := GHd(__e, tmp4822)

tmp4824 := GHd(__e, tmp4823)

tmp4825 := GEq(__e, symthaw, tmp4824)

var ifres4769 Obj

if True == tmp4825 {
tmp4812 := GTl(__e, V897)

tmp4813 := GTl(__e, tmp4812)

tmp4814 := GHd(__e, tmp4813)

tmp4815 := GTl(__e, tmp4814)

tmp4816 := GHd(__e, tmp4815)

tmp4817 := GTl(__e, tmp4816)

tmp4818 := GConsP(__e, tmp4817)

var ifres4771 Obj

if True == tmp4818 {
tmp4804 := GTl(__e, V897)

tmp4805 := GTl(__e, tmp4804)

tmp4806 := GHd(__e, tmp4805)

tmp4807 := GTl(__e, tmp4806)

tmp4808 := GHd(__e, tmp4807)

tmp4809 := GTl(__e, tmp4808)

tmp4810 := GTl(__e, tmp4809)

tmp4811 := GEq(__e, Nil, tmp4810)

var ifres4773 Obj

if True == tmp4811 {
tmp4798 := GTl(__e, V897)

tmp4799 := GTl(__e, tmp4798)

tmp4800 := GHd(__e, tmp4799)

tmp4801 := GTl(__e, tmp4800)

tmp4802 := GTl(__e, tmp4801)

tmp4803 := GEq(__e, Nil, tmp4802)

var ifres4775 Obj

if True == tmp4803 {
tmp4794 := GTl(__e, V897)

tmp4795 := GTl(__e, tmp4794)

tmp4796 := GTl(__e, tmp4795)

tmp4797 := GConsP(__e, tmp4796)

var ifres4777 Obj

if True == tmp4797 {
tmp4789 := GTl(__e, V897)

tmp4790 := GTl(__e, tmp4789)

tmp4791 := GTl(__e, tmp4790)

tmp4792 := GTl(__e, tmp4791)

tmp4793 := GEq(__e, Nil, tmp4792)

var ifres4779 Obj

if True == tmp4793 {
tmp4781 := GTl(__e, V897)

tmp4782 := GTl(__e, tmp4781)

tmp4783 := GHd(__e, tmp4782)

tmp4784 := GTl(__e, tmp4783)

tmp4785 := GHd(__e, tmp4784)

tmp4786 := GTl(__e, tmp4785)

tmp4787 := GHd(__e, tmp4786)

tmp4788 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, tmp4787)

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
tmp4739 := GTl(__e, V897)

tmp4740 := GTl(__e, tmp4739)

tmp4741 := GHd(__e, tmp4740)

tmp4742 := GTl(__e, tmp4741)

tmp4743 := GHd(__e, tmp4742)

tmp4744 := GTl(__e, tmp4743)

tmp4745 := GHd(__e, tmp4744)

tmp4746 := GTl(__e, V897)

tmp4747 := GHd(__e, tmp4746)

tmp4748 := GTl(__e, V897)

tmp4749 := GTl(__e, tmp4748)

tmp4750 := GTl(__e, tmp4749)

tmp4751 := GHd(__e, tmp4750)

return TailCall3(__e, SymFn(symsubst), tmp4745, tmp4747, tmp4751)


} else {
return V897
}


panic("unreachable")
})

tmp4854 := Call2(__e, ns2_1set, symshen_4remove_1indirection, tmp4738)


_ = tmp4854

tmp4855 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V900 := __self1
_ = V900
V901 := __self2
_ = V901
V902 := __self3
_ = V902
tmp4954 := GConsP(__e, V901)

var ifres4900 Obj

if True == tmp4954 {
tmp4952 := GHd(__e, V901)

tmp4953 := GConsP(__e, tmp4952)

var ifres4902 Obj

if True == tmp4953 {
tmp4949 := GHd(__e, V901)

tmp4950 := GHd(__e, tmp4949)

tmp4951 := GConsP(__e, tmp4950)

var ifres4904 Obj

if True == tmp4951 {
tmp4945 := GHd(__e, V901)

tmp4946 := GHd(__e, tmp4945)

tmp4947 := GHd(__e, tmp4946)

tmp4948 := GEq(__e, symand, tmp4947)

var ifres4906 Obj

if True == tmp4948 {
tmp4941 := GHd(__e, V901)

tmp4942 := GHd(__e, tmp4941)

tmp4943 := GTl(__e, tmp4942)

tmp4944 := GConsP(__e, tmp4943)

var ifres4908 Obj

if True == tmp4944 {
tmp4936 := GHd(__e, V901)

tmp4937 := GHd(__e, tmp4936)

tmp4938 := GTl(__e, tmp4937)

tmp4939 := GTl(__e, tmp4938)

tmp4940 := GConsP(__e, tmp4939)

var ifres4910 Obj

if True == tmp4940 {
tmp4930 := GHd(__e, V901)

tmp4931 := GHd(__e, tmp4930)

tmp4932 := GTl(__e, tmp4931)

tmp4933 := GTl(__e, tmp4932)

tmp4934 := GTl(__e, tmp4933)

tmp4935 := GEq(__e, Nil, tmp4934)

var ifres4912 Obj

if True == tmp4935 {
tmp4927 := GHd(__e, V901)

tmp4928 := GTl(__e, tmp4927)

tmp4929 := GConsP(__e, tmp4928)

var ifres4914 Obj

if True == tmp4929 {
tmp4923 := GHd(__e, V901)

tmp4924 := GTl(__e, tmp4923)

tmp4925 := GTl(__e, tmp4924)

tmp4926 := GEq(__e, Nil, tmp4925)

var ifres4916 Obj

if True == tmp4926 {
tmp4918 := GHd(__e, V901)

tmp4919 := GHd(__e, tmp4918)

tmp4920 := GTl(__e, tmp4919)

tmp4921 := GHd(__e, tmp4920)

tmp4922 := GEq(__e, V900, tmp4921)

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
tmp4856 := GHd(__e, V901)

tmp4857 := GHd(__e, tmp4856)

tmp4858 := GTl(__e, tmp4857)

tmp4859 := GHd(__e, tmp4858)

tmp4860 := GTl(__e, V901)

tmp4861 := GHd(__e, V901)

tmp4862 := GHd(__e, tmp4861)

tmp4863 := GTl(__e, tmp4862)

tmp4864 := GTl(__e, tmp4863)

tmp4865 := GHd(__e, tmp4864)

tmp4866 := GHd(__e, V901)

tmp4867 := GTl(__e, tmp4866)

tmp4868 := GCons(__e, tmp4865, tmp4867)

tmp4869 := GCons(__e, tmp4868, V902)

if SymFn(symshen_4pivot_1on) == __self {
__self1, __self2, __self3 = tmp4859, tmp4860, tmp4869
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4pivot_1on), tmp4859, tmp4860, tmp4869)


} else {
tmp4898 := GConsP(__e, V901)

var ifres4878 Obj

if True == tmp4898 {
tmp4896 := GHd(__e, V901)

tmp4897 := GConsP(__e, tmp4896)

var ifres4880 Obj

if True == tmp4897 {
tmp4893 := GHd(__e, V901)

tmp4894 := GTl(__e, tmp4893)

tmp4895 := GConsP(__e, tmp4894)

var ifres4882 Obj

if True == tmp4895 {
tmp4889 := GHd(__e, V901)

tmp4890 := GTl(__e, tmp4889)

tmp4891 := GTl(__e, tmp4890)

tmp4892 := GEq(__e, Nil, tmp4891)

var ifres4884 Obj

if True == tmp4892 {
tmp4886 := GHd(__e, V901)

tmp4887 := GHd(__e, tmp4886)

tmp4888 := GEq(__e, V900, tmp4887)

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
tmp4870 := GHd(__e, V901)

tmp4871 := GHd(__e, tmp4870)

tmp4872 := GTl(__e, V901)

tmp4873 := GHd(__e, V901)

tmp4874 := GTl(__e, tmp4873)

tmp4875 := GCons(__e, True, tmp4874)

tmp4876 := GCons(__e, tmp4875, V902)

if SymFn(symshen_4pivot_1on) == __self {
__self1, __self2, __self3 = tmp4871, tmp4872, tmp4876
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4pivot_1on), tmp4871, tmp4872, tmp4876)


} else {
return TailCall2(__e, SymFn(sym_8p), V902, V901)
}


}


panic("unreachable")
})
return __self
})()

tmp4955 := Call2(__e, ns2_1set, symshen_4pivot_1on, tmp4855)


_ = tmp4955

tmp4956 := MakeNative2(func(__e *ControlFlow, V905, V906 Obj) Obj {
tmp4979 := GConsP(__e, V905)

var ifres4970 Obj

if True == tmp4979 {
tmp4977 := GTl(__e, V905)

tmp4978 := GConsP(__e, tmp4977)

var ifres4972 Obj

if True == tmp4978 {
tmp4974 := GTl(__e, V905)

tmp4975 := GTl(__e, tmp4974)

tmp4976 := GEq(__e, Nil, tmp4975)

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
tmp4957 := GHd(__e, V905)

tmp4958 := Call1(__e, SymFn(symshen_4op), tmp4957)


let__2731 := tmp4958
_ = let__2731

tmp4968 := GEq(__e, symshen_4skip, let__2731)

if True == tmp4968 {
return V906
} else {
tmp4959 := Call1(__e, SymFn(symshen_4op1), let__2731)


tmp4960 := GTl(__e, V905)

tmp4961 := GCons(__e, tmp4959, tmp4960)

tmp4962 := Call1(__e, SymFn(symshen_4op2), let__2731)


tmp4963 := GTl(__e, V905)

tmp4964 := GCons(__e, tmp4962, tmp4963)

tmp4965 := GCons(__e, tmp4964, Nil)

tmp4966 := GCons(__e, tmp4961, tmp4965)

return TailCall2(__e, SymFn(symshen_4factor_1selectors_1h), tmp4966, V906)


}


} else {
return V906
}


panic("unreachable")
})

tmp4980 := Call2(__e, ns2_1set, symshen_4factor_1selectors, tmp4956)


_ = tmp4980

tmp4981 := MakeNative1(func(__e *ControlFlow, V910 Obj) Obj {
tmp4989 := GEq(__e, symcons_2, V910)

if True == tmp4989 {
return symcons
} else {
tmp4987 := GEq(__e, symshen_4_7string_2, V910)

if True == tmp4987 {
return sym_8s
} else {
tmp4985 := GEq(__e, symshen_4_7vector_2, V910)

if True == tmp4985 {
return sym_8v
} else {
tmp4983 := GEq(__e, symtuple_2, V910)

if True == tmp4983 {
return sym_8p
} else {
return symshen_4skip
}


}


}


}


panic("unreachable")
})

tmp4990 := Call2(__e, ns2_1set, symshen_4op, tmp4981)


_ = tmp4990

tmp4991 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V911 := __self1
_ = V911
V912 := __self2
_ = V912
tmp5009 := GEq(__e, Nil, V911)

if True == tmp5009 {
return V912
} else {
tmp5007 := GConsP(__e, V911)

if True == tmp5007 {
tmp5003 := GHd(__e, V911)

tmp5004 := Call2(__e, SymFn(symoccurrences), tmp5003, V912)


if True == GGt(__e, tmp5004, MakeInteger(1)) {
tmp4992 := Call1(__e, SymFn(symgensym), symSelect)


let__2732 := tmp4992
_ = let__2732

tmp4993 := GHd(__e, V911)

tmp4994 := GTl(__e, V911)

tmp4995 := GHd(__e, V911)

tmp4996 := Call3(__e, SymFn(symsubst), let__2732, tmp4995, V912)


tmp4997 := Call2(__e, SymFn(symshen_4factor_1selectors_1h), tmp4994, tmp4996)


tmp4998 := GCons(__e, tmp4997, Nil)

tmp4999 := GCons(__e, tmp4993, tmp4998)

tmp5000 := GCons(__e, let__2732, tmp4999)

return GCons(__e, symlet, tmp5000)


} else {
tmp5001 := GTl(__e, V911)

if SymFn(symshen_4factor_1selectors_1h) == __self {
__self1, __self2 = tmp5001, V912
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4factor_1selectors_1h), tmp5001, V912)


}


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4factor_1selectors_1h)
}


}


panic("unreachable")
})
return __self
})()

return TailCall2(__e, ns2_1set, symshen_4factor_1selectors_1h, tmp4991)




panic("unreachable")
})

