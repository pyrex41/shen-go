package main

import . "github.com/pyrex41/shen-go/kl"

var CoreMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp2663 := MakeNative1(func(__e *ControlFlow, V540 Obj) Obj {
tmp2664 := Call1(__e, SymFn(symshen_4shen_1_6kl_1h), V540)


let__2511 := tmp2664
_ = let__2511

return TailCall1(__e, SymFn(symshen_4record_1and_1evaluate), let__2511)


panic("unreachable")
})

tmp2665 := Call2(__e, ns2_1set, symshen_4shen_1_6kl, tmp2663)


_ = tmp2665

tmp2666 := MakeNative1(func(__e *ControlFlow, V542 Obj) Obj {
tmp2710 := GConsP(__e, V542)

var ifres2689 Obj

if True == tmp2710 {
tmp2708 := GHd(__e, V542)

tmp2709 := GEq(__e, symdefun, tmp2708)

var ifres2690 Obj

if True == tmp2709 {
tmp2706 := GTl(__e, V542)

tmp2707 := GConsP(__e, tmp2706)

var ifres2691 Obj

if True == tmp2707 {
tmp2703 := GTl(__e, V542)

tmp2704 := GTl(__e, tmp2703)

tmp2705 := GConsP(__e, tmp2704)

var ifres2692 Obj

if True == tmp2705 {
tmp2699 := GTl(__e, V542)

tmp2700 := GTl(__e, tmp2699)

tmp2701 := GTl(__e, tmp2700)

tmp2702 := GConsP(__e, tmp2701)

var ifres2693 Obj

if True == tmp2702 {
tmp2694 := GTl(__e, V542)

tmp2695 := GTl(__e, tmp2694)

tmp2696 := GTl(__e, tmp2695)

tmp2697 := GTl(__e, tmp2696)

tmp2698 := GEq(__e, Nil, tmp2697)

ifres2693 = tmp2698


} else {
ifres2693 = False


}

ifres2692 = ifres2693


} else {
ifres2692 = False


}

ifres2691 = ifres2692


} else {
ifres2691 = False


}

ifres2690 = ifres2691


} else {
ifres2690 = False


}

ifres2689 = ifres2690


} else {
ifres2689 = False


}

if True == ifres2689 {
tmp2672 := GTl(__e, V542)

tmp2673 := GHd(__e, tmp2672)

tmp2674 := Call1(__e, SymFn(symshen_4sysfunc_2), tmp2673)


var ifres2667 Obj

if True == tmp2674 {
tmp2668 := GTl(__e, V542)

tmp2669 := GHd(__e, tmp2668)

tmp2670 := Call3(__e, SymFn(symshen_4app), tmp2669, MakeString(" is not a legitimate function name\n"), symshen_4a)


tmp2671 := PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp2670)

ifres2667 = tmp2671


} else {
ifres2667 = symshen_4skip


}

let__2512 := ifres2667
_ = let__2512

tmp2675 := GTl(__e, V542)

tmp2676 := GHd(__e, tmp2675)

tmp2677 := GTl(__e, V542)

tmp2678 := GTl(__e, tmp2677)

tmp2679 := GHd(__e, tmp2678)

tmp2680 := Call1(__e, SymFn(symlength), tmp2679)


tmp2681 := Call2(__e, SymFn(symshen_4store_1arity), tmp2676, tmp2680)


let__2513 := tmp2681
_ = let__2513

tmp2682 := GTl(__e, V542)

tmp2683 := GHd(__e, tmp2682)

tmp2684 := Call2(__e, SymFn(symshen_4record_1kl), tmp2683, V542)


let__2514 := tmp2684
_ = let__2514

tmp2685 := Call1(__e, SymFn(symeval_1kl), V542)


let__2515 := tmp2685
_ = let__2515

tmp2686 := GTl(__e, V542)

tmp2687 := GHd(__e, tmp2686)

return TailCall1(__e, SymFn(symshen_4fn_1print), tmp2687)


} else {
return V542
}


panic("unreachable")
})

tmp2711 := Call2(__e, ns2_1set, symshen_4record_1and_1evaluate, tmp2666)


_ = tmp2711

tmp2712 := MakeNative1(func(__e *ControlFlow, V547 Obj) Obj {
tmp2798 := GConsP(__e, V547)

var ifres2792 Obj

if True == tmp2798 {
tmp2796 := GHd(__e, V547)

tmp2797 := GEq(__e, symdefine, tmp2796)

var ifres2793 Obj

if True == tmp2797 {
tmp2794 := GTl(__e, V547)

tmp2795 := GConsP(__e, tmp2794)

ifres2793 = tmp2795


} else {
ifres2793 = False


}

ifres2792 = ifres2793


} else {
ifres2792 = False


}

if True == ifres2792 {
tmp2713 := GTl(__e, V547)

tmp2714 := GHd(__e, tmp2713)

tmp2715 := GTl(__e, V547)

tmp2716 := GTl(__e, tmp2715)

return TailCall2(__e, SymFn(symshen_4shendef_1_6kldef), tmp2714, tmp2716)


} else {
tmp2790 := GConsP(__e, V547)

var ifres2769 Obj

if True == tmp2790 {
tmp2788 := GHd(__e, V547)

tmp2789 := GEq(__e, symdefun, tmp2788)

var ifres2770 Obj

if True == tmp2789 {
tmp2786 := GTl(__e, V547)

tmp2787 := GConsP(__e, tmp2786)

var ifres2771 Obj

if True == tmp2787 {
tmp2783 := GTl(__e, V547)

tmp2784 := GTl(__e, tmp2783)

tmp2785 := GConsP(__e, tmp2784)

var ifres2772 Obj

if True == tmp2785 {
tmp2779 := GTl(__e, V547)

tmp2780 := GTl(__e, tmp2779)

tmp2781 := GTl(__e, tmp2780)

tmp2782 := GConsP(__e, tmp2781)

var ifres2773 Obj

if True == tmp2782 {
tmp2774 := GTl(__e, V547)

tmp2775 := GTl(__e, tmp2774)

tmp2776 := GTl(__e, tmp2775)

tmp2777 := GTl(__e, tmp2776)

tmp2778 := GEq(__e, Nil, tmp2777)

ifres2773 = tmp2778


} else {
ifres2773 = False


}

ifres2772 = ifres2773


} else {
ifres2772 = False


}

ifres2771 = ifres2772


} else {
ifres2771 = False


}

ifres2770 = ifres2771


} else {
ifres2770 = False


}

ifres2769 = ifres2770


} else {
ifres2769 = False


}

if True == ifres2769 {
return V547
} else {
tmp2767 := GConsP(__e, V547)

var ifres2752 Obj

if True == tmp2767 {
tmp2765 := GHd(__e, V547)

tmp2766 := GEq(__e, symtype, tmp2765)

var ifres2753 Obj

if True == tmp2766 {
tmp2763 := GTl(__e, V547)

tmp2764 := GConsP(__e, tmp2763)

var ifres2754 Obj

if True == tmp2764 {
tmp2760 := GTl(__e, V547)

tmp2761 := GTl(__e, tmp2760)

tmp2762 := GConsP(__e, tmp2761)

var ifres2755 Obj

if True == tmp2762 {
tmp2756 := GTl(__e, V547)

tmp2757 := GTl(__e, tmp2756)

tmp2758 := GTl(__e, tmp2757)

tmp2759 := GEq(__e, Nil, tmp2758)

ifres2755 = tmp2759


} else {
ifres2755 = False


}

ifres2754 = ifres2755


} else {
ifres2754 = False


}

ifres2753 = ifres2754


} else {
ifres2753 = False


}

ifres2752 = ifres2753


} else {
ifres2752 = False


}

if True == ifres2752 {
tmp2717 := GTl(__e, V547)

tmp2718 := GHd(__e, tmp2717)

tmp2719 := GTl(__e, V547)

tmp2720 := GTl(__e, tmp2719)

tmp2721 := GHd(__e, tmp2720)

tmp2722 := Call1(__e, SymFn(symshen_4rcons__form), tmp2721)


tmp2723 := GCons(__e, tmp2722, Nil)

tmp2724 := GCons(__e, tmp2718, tmp2723)

return GCons(__e, symtype, tmp2724)


} else {
tmp2750 := GConsP(__e, V547)

var ifres2735 Obj

if True == tmp2750 {
tmp2748 := GHd(__e, V547)

tmp2749 := GEq(__e, syminput_7, tmp2748)

var ifres2736 Obj

if True == tmp2749 {
tmp2746 := GTl(__e, V547)

tmp2747 := GConsP(__e, tmp2746)

var ifres2737 Obj

if True == tmp2747 {
tmp2743 := GTl(__e, V547)

tmp2744 := GTl(__e, tmp2743)

tmp2745 := GConsP(__e, tmp2744)

var ifres2738 Obj

if True == tmp2745 {
tmp2739 := GTl(__e, V547)

tmp2740 := GTl(__e, tmp2739)

tmp2741 := GTl(__e, tmp2740)

tmp2742 := GEq(__e, Nil, tmp2741)

ifres2738 = tmp2742


} else {
ifres2738 = False


}

ifres2737 = ifres2738


} else {
ifres2737 = False


}

ifres2736 = ifres2737


} else {
ifres2736 = False


}

ifres2735 = ifres2736


} else {
ifres2735 = False


}

if True == ifres2735 {
tmp2725 := GTl(__e, V547)

tmp2726 := GHd(__e, tmp2725)

tmp2727 := Call1(__e, SymFn(symshen_4rcons__form), tmp2726)


tmp2728 := GTl(__e, V547)

tmp2729 := GTl(__e, tmp2728)

tmp2730 := GCons(__e, tmp2727, tmp2729)

return GCons(__e, syminput_7, tmp2730)


} else {
tmp2733 := GConsP(__e, V547)

if True == tmp2733 {
tmp2731 := MakeNative1(func(__e *ControlFlow, Z548 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4shen_1_6kl_1h), Z548)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp2731, V547)


} else {
return V547
}


}


}


}


}


panic("unreachable")
})

tmp2799 := Call2(__e, ns2_1set, symshen_4shen_1_6kl_1h, tmp2712)


_ = tmp2799

tmp2800 := MakeNative2(func(__e *ControlFlow, V549, V550 Obj) Obj {
tmp2801 := MakeNative1(func(__e *ControlFlow, Z551 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5define_6), Z551)
panic("unreachable")
})

tmp2802 := GCons(__e, V549, V550)

return TailCall2(__e, SymFn(symcompile), tmp2801, tmp2802)


panic("unreachable")
})

tmp2803 := Call2(__e, ns2_1set, symshen_4shendef_1_6kldef, tmp2800)


_ = tmp2803

tmp2804 := MakeNative1(func(__e *ControlFlow, V552 Obj) Obj {
tmp2805 := Call1(__e, SymFn(symshen_4_5name_6), V552)


let__2517 := tmp2805
_ = let__2517

tmp2831 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2517)


var ifres2806 Obj

if True == tmp2831 {
tmp2807 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2806 = tmp2807


} else {
tmp2808 := Call1(__e, SymFn(symshen_4_5_1out), let__2517)


let__2518 := tmp2808
_ = let__2518

tmp2809 := Call1(__e, SymFn(symshen_4in_1_6), let__2517)


let__2519 := tmp2809
_ = let__2519

tmp2830 := Call2(__e, SymFn(symshen_4hds_a_2), let__2519, sym_i)


var ifres2810 Obj

if True == tmp2830 {
tmp2811 := Call1(__e, SymFn(symtail), let__2519)


let__2520 := tmp2811
_ = let__2520

tmp2812 := Call1(__e, SymFn(symshen_4_5signature_6), let__2520)


let__2521 := tmp2812
_ = let__2521

tmp2828 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2521)


var ifres2813 Obj

if True == tmp2828 {
tmp2814 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2813 = tmp2814


} else {
tmp2815 := Call1(__e, SymFn(symshen_4in_1_6), let__2521)


let__2522 := tmp2815
_ = let__2522

tmp2827 := Call2(__e, SymFn(symshen_4hds_a_2), let__2522, sym_j)


var ifres2816 Obj

if True == tmp2827 {
tmp2817 := Call1(__e, SymFn(symtail), let__2522)


let__2523 := tmp2817
_ = let__2523

tmp2818 := Call1(__e, SymFn(symshen_4_5rules_6), let__2523)


let__2524 := tmp2818
_ = let__2524

tmp2825 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2524)


var ifres2819 Obj

if True == tmp2825 {
tmp2820 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2819 = tmp2820


} else {
tmp2821 := Call1(__e, SymFn(symshen_4_5_1out), let__2524)


let__2525 := tmp2821
_ = let__2525

tmp2822 := Call1(__e, SymFn(symshen_4in_1_6), let__2524)


let__2526 := tmp2822
_ = let__2526

tmp2823 := Call2(__e, SymFn(symshen_4shendef_1_6kldef_1h), let__2518, let__2525)


tmp2824 := Call2(__e, SymFn(symshen_4comb), let__2526, tmp2823)


ifres2819 = tmp2824


}

ifres2816 = ifres2819


} else {
tmp2826 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2816 = tmp2826


}

ifres2813 = ifres2816


}

ifres2810 = ifres2813


} else {
tmp2829 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2810 = tmp2829


}

ifres2806 = ifres2810


}

let__2516 := ifres2806
_ = let__2516

tmp2849 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2516)


if True == tmp2849 {
tmp2832 := Call1(__e, SymFn(symshen_4_5name_6), V552)


let__2528 := tmp2832
_ = let__2528

tmp2845 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2528)


var ifres2833 Obj

if True == tmp2845 {
tmp2834 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2833 = tmp2834


} else {
tmp2835 := Call1(__e, SymFn(symshen_4_5_1out), let__2528)


let__2529 := tmp2835
_ = let__2529

tmp2836 := Call1(__e, SymFn(symshen_4in_1_6), let__2528)


let__2530 := tmp2836
_ = let__2530

tmp2837 := Call1(__e, SymFn(symshen_4_5rules_6), let__2530)


let__2531 := tmp2837
_ = let__2531

tmp2844 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2531)


var ifres2838 Obj

if True == tmp2844 {
tmp2839 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2838 = tmp2839


} else {
tmp2840 := Call1(__e, SymFn(symshen_4_5_1out), let__2531)


let__2532 := tmp2840
_ = let__2532

tmp2841 := Call1(__e, SymFn(symshen_4in_1_6), let__2531)


let__2533 := tmp2841
_ = let__2533

tmp2842 := Call2(__e, SymFn(symshen_4shendef_1_6kldef_1h), let__2529, let__2532)


tmp2843 := Call2(__e, SymFn(symshen_4comb), let__2533, tmp2842)


ifres2838 = tmp2843


}

ifres2833 = ifres2838


}

let__2527 := ifres2833
_ = let__2527

tmp2847 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2527)


if True == tmp2847 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2527
}


} else {
return let__2516
}


panic("unreachable")
})

tmp2850 := Call2(__e, ns2_1set, symshen_4_5define_6, tmp2804)


_ = tmp2850

tmp2851 := MakeNative2(func(__e *ControlFlow, V571, V572 Obj) Obj {
tmp2852 := MakeNative1(func(__e *ControlFlow, Z574 Obj) Obj {
return TailCall1(__e, SymFn(symfst), Z574)
panic("unreachable")
})

tmp2853 := Call2(__e, SymFn(symmap), tmp2852, V572)


let__2534 := tmp2853
_ = let__2534

tmp2854 := Call2(__e, SymFn(symshen_4arity_1chk), V571, let__2534)


let__2535 := tmp2854
_ = let__2535

tmp2855 := MakeNative1(func(__e *ControlFlow, Z577 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4free_1var_1chk), V571, Z577)
panic("unreachable")
})

tmp2856 := Call2(__e, SymFn(symmap), tmp2855, V572)


let__2536 := tmp2856
_ = let__2536

tmp2857 := Call1(__e, SymFn(symshen_4unprotect), V572)


let__2537 := tmp2857
_ = let__2537

tmp2858 := Call3(__e, SymFn(symshen_4compile_1to_1kl), V571, let__2537, let__2535)


tmp2859 := Call1(__e, SymFn(symshen_4factorise_1code), tmp2858)


let__2538 := tmp2859
_ = let__2538

return let__2538


panic("unreachable")
})

tmp2860 := Call2(__e, ns2_1set, symshen_4shendef_1_6kldef_1h, tmp2851)


_ = tmp2860

tmp2861 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V580 := __self1
_ = V580
tmp2884 := Call1(__e, SymFn(symtuple_2), V580)


if True == tmp2884 {
tmp2862 := Call1(__e, SymFn(symfst), V580)


tmp2863 := Call1(__e, SymFn(symshen_4unprotect), tmp2862)


tmp2864 := Call1(__e, SymFn(symsnd), V580)


tmp2865 := Call1(__e, SymFn(symshen_4unprotect), tmp2864)


return TailCall2(__e, SymFn(sym_8p), tmp2863, tmp2865)


} else {
tmp2882 := GConsP(__e, V580)

var ifres2872 Obj

if True == tmp2882 {
tmp2880 := GHd(__e, V580)

tmp2881 := GEq(__e, symprotect, tmp2880)

var ifres2873 Obj

if True == tmp2881 {
tmp2878 := GTl(__e, V580)

tmp2879 := GConsP(__e, tmp2878)

var ifres2874 Obj

if True == tmp2879 {
tmp2875 := GTl(__e, V580)

tmp2876 := GTl(__e, tmp2875)

tmp2877 := GEq(__e, Nil, tmp2876)

ifres2874 = tmp2877


} else {
ifres2874 = False


}

ifres2873 = ifres2874


} else {
ifres2873 = False


}

ifres2872 = ifres2873


} else {
ifres2872 = False


}

if True == ifres2872 {
tmp2866 := GTl(__e, V580)

tmp2867 := GHd(__e, tmp2866)

if SymFn(symshen_4unprotect) == __self {
__self1 = tmp2867
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4unprotect), tmp2867)


} else {
tmp2870 := GConsP(__e, V580)

if True == tmp2870 {
tmp2868 := MakeNative1(func(__e *ControlFlow, Z581 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4unprotect), Z581)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp2868, V580)


} else {
return V580
}


}


}


panic("unreachable")
})
return __self
})()

tmp2885 := Call2(__e, ns2_1set, symshen_4unprotect, tmp2861)


_ = tmp2885

tmp2886 := MakeNative1(func(__e *ControlFlow, V582 Obj) Obj {
tmp2899 := GConsP(__e, V582)

var ifres2887 Obj

if True == tmp2899 {
tmp2888 := Call1(__e, SymFn(symhead), V582)


let__2540 := tmp2888
_ = let__2540

tmp2889 := Call1(__e, SymFn(symtail), V582)


let__2541 := tmp2889
_ = let__2541

tmp2896 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, let__2540)

var ifres2893 Obj

if True == tmp2896 {
tmp2894 := PrimGuard1(__e, symvariable_2, PrimIsVariable, let__2540)

ifres2893 = GNot(__e, tmp2894)


} else {
ifres2893 = False


}

var ifres2890 Obj

if True == ifres2893 {
ifres2890 = let__2540


} else {
tmp2891 := Call3(__e, SymFn(symshen_4app), let__2540, MakeString(" is not a legitimate function name.\n"), symshen_4a)


tmp2892 := PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp2891)

ifres2890 = tmp2892


}

tmp2897 := Call2(__e, SymFn(symshen_4comb), let__2541, ifres2890)


ifres2887 = tmp2897


} else {
tmp2898 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2887 = tmp2898


}

let__2539 := ifres2887
_ = let__2539

tmp2901 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2539)


if True == tmp2901 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2539
}


panic("unreachable")
})

tmp2902 := Call2(__e, ns2_1set, symshen_4_5name_6, tmp2886)


_ = tmp2902

tmp2903 := MakeNative1(func(__e *ControlFlow, V586 Obj) Obj {
tmp2922 := GConsP(__e, V586)

var ifres2904 Obj

if True == tmp2922 {
tmp2905 := Call1(__e, SymFn(symhead), V586)


let__2543 := tmp2905
_ = let__2543

tmp2906 := Call1(__e, SymFn(symtail), V586)


let__2544 := tmp2906
_ = let__2544

tmp2907 := Call1(__e, SymFn(symshen_4_5signature_6), let__2544)


let__2545 := tmp2907
_ = let__2545

tmp2920 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2545)


var ifres2908 Obj

if True == tmp2920 {
tmp2909 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2908 = tmp2909


} else {
tmp2910 := Call1(__e, SymFn(symshen_4_5_1out), let__2545)


let__2546 := tmp2910
_ = let__2546

tmp2911 := Call1(__e, SymFn(symshen_4in_1_6), let__2545)


let__2547 := tmp2911
_ = let__2547

tmp2916 := GCons(__e, sym_j, Nil)

tmp2917 := GCons(__e, sym_i, tmp2916)

tmp2918 := Call2(__e, SymFn(symelement_2), let__2543, tmp2917)


tmp2919 := GNot(__e, tmp2918)

var ifres2912 Obj

if True == tmp2919 {
tmp2913 := GCons(__e, let__2543, let__2546)

tmp2914 := Call2(__e, SymFn(symshen_4comb), let__2547, tmp2913)


ifres2912 = tmp2914


} else {
tmp2915 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2912 = tmp2915


}

ifres2908 = ifres2912


}

ifres2904 = ifres2908


} else {
tmp2921 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2904 = tmp2921


}

let__2542 := ifres2904
_ = let__2542

tmp2932 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2542)


if True == tmp2932 {
tmp2923 := Call1(__e, SymFn(sym_5e_6), V586)


let__2549 := tmp2923
_ = let__2549

tmp2928 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2549)


var ifres2924 Obj

if True == tmp2928 {
tmp2925 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2924 = tmp2925


} else {
tmp2926 := Call1(__e, SymFn(symshen_4in_1_6), let__2549)


let__2550 := tmp2926
_ = let__2550

tmp2927 := Call2(__e, SymFn(symshen_4comb), let__2550, Nil)


ifres2924 = tmp2927


}

let__2548 := ifres2924
_ = let__2548

tmp2930 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2548)


if True == tmp2930 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2548
}


} else {
return let__2542
}


panic("unreachable")
})

tmp2933 := Call2(__e, ns2_1set, symshen_4_5signature_6, tmp2903)


_ = tmp2933

tmp2934 := MakeNative1(func(__e *ControlFlow, V596 Obj) Obj {
tmp2935 := Call1(__e, SymFn(symshen_4_5rule_6), V596)


let__2552 := tmp2935
_ = let__2552

tmp2949 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2552)


var ifres2936 Obj

if True == tmp2949 {
tmp2937 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2936 = tmp2937


} else {
tmp2938 := Call1(__e, SymFn(symshen_4_5_1out), let__2552)


let__2553 := tmp2938
_ = let__2553

tmp2939 := Call1(__e, SymFn(symshen_4in_1_6), let__2552)


let__2554 := tmp2939
_ = let__2554

tmp2940 := Call1(__e, SymFn(symshen_4_5rules_6), let__2554)


let__2555 := tmp2940
_ = let__2555

tmp2948 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2555)


var ifres2941 Obj

if True == tmp2948 {
tmp2942 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2941 = tmp2942


} else {
tmp2943 := Call1(__e, SymFn(symshen_4_5_1out), let__2555)


let__2556 := tmp2943
_ = let__2556

tmp2944 := Call1(__e, SymFn(symshen_4in_1_6), let__2555)


let__2557 := tmp2944
_ = let__2557

tmp2945 := Call1(__e, SymFn(symshen_4linearise), let__2553)


tmp2946 := GCons(__e, tmp2945, let__2556)

tmp2947 := Call2(__e, SymFn(symshen_4comb), let__2557, tmp2946)


ifres2941 = tmp2947


}

ifres2936 = ifres2941


}

let__2551 := ifres2936
_ = let__2551

tmp2965 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2551)


if True == tmp2965 {
tmp2950 := Call1(__e, SymFn(sym_5_b_6), V596)


let__2559 := tmp2950
_ = let__2559

tmp2961 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2559)


var ifres2951 Obj

if True == tmp2961 {
tmp2952 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2951 = tmp2952


} else {
tmp2953 := Call1(__e, SymFn(symshen_4_5_1out), let__2559)


let__2560 := tmp2953
_ = let__2560

tmp2954 := Call1(__e, SymFn(symshen_4in_1_6), let__2559)


let__2561 := tmp2954
_ = let__2561

tmp2959 := Call1(__e, SymFn(symempty_2), let__2560)


var ifres2955 Obj

if True == tmp2959 {
ifres2955 = Nil


} else {
tmp2956 := Call3(__e, SymFn(symshen_4app), let__2560, MakeString("\n ..."), symshen_4r)


tmp2958 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("Shen syntax error here:\n "), tmp2956))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("Shen syntax error here:\n "), tmp2956)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()

ifres2955 = tmp2958


}

tmp2960 := Call2(__e, SymFn(symshen_4comb), let__2561, ifres2955)


ifres2951 = tmp2960


}

let__2558 := ifres2951
_ = let__2558

tmp2963 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2558)


if True == tmp2963 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2558
}


} else {
return let__2551
}


panic("unreachable")
})

tmp2966 := Call2(__e, ns2_1set, symshen_4_5rules_6, tmp2934)


_ = tmp2966

tmp2967 := MakeNative1(func(__e *ControlFlow, V610 Obj) Obj {
tmp2972 := Call1(__e, SymFn(symtuple_2), V610)


if True == tmp2972 {
tmp2968 := Call1(__e, SymFn(symfst), V610)


tmp2969 := Call1(__e, SymFn(symfst), V610)


tmp2970 := Call1(__e, SymFn(symsnd), V610)


return TailCall4(__e, SymFn(symshen_4linearise_1h), tmp2968, tmp2969, Nil, tmp2970)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.linearise"))
}


panic("unreachable")
})

tmp2973 := Call2(__e, ns2_1set, symshen_4linearise, tmp2967)


_ = tmp2973

tmp2974 := (func() Obj {
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
tmp3009 := GEq(__e, Nil, V623)

if True == tmp3009 {
return TailCall2(__e, SymFn(sym_8p), V624, V626)
} else {
tmp3007 := GConsP(__e, V623)

var ifres3004 Obj

if True == tmp3007 {
tmp3005 := GHd(__e, V623)

tmp3006 := GConsP(__e, tmp3005)

ifres3004 = tmp3006


} else {
ifres3004 = False


}

if True == ifres3004 {
tmp2975 := GHd(__e, V623)

tmp2976 := GTl(__e, V623)

tmp2977 := Call2(__e, SymFn(symappend), tmp2975, tmp2976)


if SymFn(symshen_4linearise_1h) == __self {
__self1, __self2, __self3, __self4 = tmp2977, V624, V625, V626
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4linearise_1h), tmp2977, V624, V625, V626)


} else {
tmp3002 := GConsP(__e, V623)

var ifres2999 Obj

if True == tmp3002 {
tmp3000 := GHd(__e, V623)

tmp3001 := PrimGuard1(__e, symvariable_2, PrimIsVariable, tmp3000)

ifres2999 = tmp3001


} else {
ifres2999 = False


}

if True == ifres2999 {
tmp2993 := GHd(__e, V623)

tmp2994 := Call2(__e, SymFn(symelement_2), tmp2993, V625)


if True == tmp2994 {
tmp2978 := Call1(__e, SymFn(symgensym), symV)


let__2562 := tmp2978
_ = let__2562

tmp2979 := GTl(__e, V623)

tmp2980 := GHd(__e, V623)

tmp2981 := Call3(__e, SymFn(symshen_4rep_1X), tmp2980, let__2562, V624)


tmp2982 := GHd(__e, V623)

tmp2983 := GCons(__e, tmp2982, Nil)

tmp2984 := GCons(__e, let__2562, tmp2983)

tmp2985 := GCons(__e, sym_a, tmp2984)

tmp2986 := GCons(__e, V626, Nil)

tmp2987 := GCons(__e, tmp2985, tmp2986)

tmp2988 := GCons(__e, symwhere, tmp2987)

if SymFn(symshen_4linearise_1h) == __self {
__self1, __self2, __self3, __self4 = tmp2979, tmp2981, V625, tmp2988
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4linearise_1h), tmp2979, tmp2981, V625, tmp2988)


} else {
tmp2989 := GTl(__e, V623)

tmp2990 := GHd(__e, V623)

tmp2991 := GCons(__e, tmp2990, V625)

if SymFn(symshen_4linearise_1h) == __self {
__self1, __self2, __self3, __self4 = tmp2989, V624, tmp2991, V626
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4linearise_1h), tmp2989, V624, tmp2991, V626)


}


} else {
tmp2997 := GConsP(__e, V623)

if True == tmp2997 {
tmp2995 := GTl(__e, V623)

if SymFn(symshen_4linearise_1h) == __self {
__self1, __self2, __self3, __self4 = tmp2995, V624, V625, V626
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4linearise_1h), tmp2995, V624, V625, V626)


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

tmp3010 := Call2(__e, ns2_1set, symshen_4linearise_1h, tmp2974)


_ = tmp3010

tmp3011 := MakeNative1(func(__e *ControlFlow, V628 Obj) Obj {
tmp3012 := Call1(__e, SymFn(symshen_4_5patterns_6), V628)


let__2564 := tmp3012
_ = let__2564

tmp3040 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2564)


var ifres3013 Obj

if True == tmp3040 {
tmp3014 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3013 = tmp3014


} else {
tmp3015 := Call1(__e, SymFn(symshen_4_5_1out), let__2564)


let__2565 := tmp3015
_ = let__2565

tmp3016 := Call1(__e, SymFn(symshen_4in_1_6), let__2564)


let__2566 := tmp3016
_ = let__2566

tmp3039 := Call2(__e, SymFn(symshen_4hds_a_2), let__2566, sym_1_6)


var ifres3017 Obj

if True == tmp3039 {
tmp3018 := Call1(__e, SymFn(symtail), let__2566)


let__2567 := tmp3018
_ = let__2567

tmp3037 := GConsP(__e, let__2567)

var ifres3019 Obj

if True == tmp3037 {
tmp3020 := Call1(__e, SymFn(symhead), let__2567)


let__2568 := tmp3020
_ = let__2568

tmp3021 := Call1(__e, SymFn(symtail), let__2567)


let__2569 := tmp3021
_ = let__2569

tmp3035 := Call2(__e, SymFn(symshen_4hds_a_2), let__2569, symwhere)


var ifres3022 Obj

if True == tmp3035 {
tmp3023 := Call1(__e, SymFn(symtail), let__2569)


let__2570 := tmp3023
_ = let__2570

tmp3033 := GConsP(__e, let__2570)

var ifres3024 Obj

if True == tmp3033 {
tmp3025 := Call1(__e, SymFn(symhead), let__2570)


let__2571 := tmp3025
_ = let__2571

tmp3026 := Call1(__e, SymFn(symtail), let__2570)


let__2572 := tmp3026
_ = let__2572

tmp3027 := GCons(__e, let__2568, Nil)

tmp3028 := GCons(__e, let__2571, tmp3027)

tmp3029 := GCons(__e, symwhere, tmp3028)

tmp3030 := Call2(__e, SymFn(sym_8p), let__2565, tmp3029)


tmp3031 := Call2(__e, SymFn(symshen_4comb), let__2572, tmp3030)


ifres3024 = tmp3031


} else {
tmp3032 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3024 = tmp3032


}

ifres3022 = ifres3024


} else {
tmp3034 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3022 = tmp3034


}

ifres3019 = ifres3022


} else {
tmp3036 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3019 = tmp3036


}

ifres3017 = ifres3019


} else {
tmp3038 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3017 = tmp3038


}

ifres3013 = ifres3017


}

let__2563 := ifres3013
_ = let__2563

tmp3115 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2563)


if True == tmp3115 {
tmp3041 := Call1(__e, SymFn(symshen_4_5patterns_6), V628)


let__2574 := tmp3041
_ = let__2574

tmp3057 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2574)


var ifres3042 Obj

if True == tmp3057 {
tmp3043 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3042 = tmp3043


} else {
tmp3044 := Call1(__e, SymFn(symshen_4_5_1out), let__2574)


let__2575 := tmp3044
_ = let__2575

tmp3045 := Call1(__e, SymFn(symshen_4in_1_6), let__2574)


let__2576 := tmp3045
_ = let__2576

tmp3056 := Call2(__e, SymFn(symshen_4hds_a_2), let__2576, sym_1_6)


var ifres3046 Obj

if True == tmp3056 {
tmp3047 := Call1(__e, SymFn(symtail), let__2576)


let__2577 := tmp3047
_ = let__2577

tmp3054 := GConsP(__e, let__2577)

var ifres3048 Obj

if True == tmp3054 {
tmp3049 := Call1(__e, SymFn(symhead), let__2577)


let__2578 := tmp3049
_ = let__2578

tmp3050 := Call1(__e, SymFn(symtail), let__2577)


let__2579 := tmp3050
_ = let__2579

tmp3051 := Call2(__e, SymFn(sym_8p), let__2575, let__2578)


tmp3052 := Call2(__e, SymFn(symshen_4comb), let__2579, tmp3051)


ifres3048 = tmp3052


} else {
tmp3053 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3048 = tmp3053


}

ifres3046 = ifres3048


} else {
tmp3055 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3046 = tmp3055


}

ifres3042 = ifres3046


}

let__2573 := ifres3042
_ = let__2573

tmp3113 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2573)


if True == tmp3113 {
tmp3058 := Call1(__e, SymFn(symshen_4_5patterns_6), V628)


let__2581 := tmp3058
_ = let__2581

tmp3088 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2581)


var ifres3059 Obj

if True == tmp3088 {
tmp3060 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3059 = tmp3060


} else {
tmp3061 := Call1(__e, SymFn(symshen_4_5_1out), let__2581)


let__2582 := tmp3061
_ = let__2582

tmp3062 := Call1(__e, SymFn(symshen_4in_1_6), let__2581)


let__2583 := tmp3062
_ = let__2583

tmp3087 := Call2(__e, SymFn(symshen_4hds_a_2), let__2583, sym_5_1)


var ifres3063 Obj

if True == tmp3087 {
tmp3064 := Call1(__e, SymFn(symtail), let__2583)


let__2584 := tmp3064
_ = let__2584

tmp3085 := GConsP(__e, let__2584)

var ifres3065 Obj

if True == tmp3085 {
tmp3066 := Call1(__e, SymFn(symhead), let__2584)


let__2585 := tmp3066
_ = let__2585

tmp3067 := Call1(__e, SymFn(symtail), let__2584)


let__2586 := tmp3067
_ = let__2586

tmp3083 := Call2(__e, SymFn(symshen_4hds_a_2), let__2586, symwhere)


var ifres3068 Obj

if True == tmp3083 {
tmp3069 := Call1(__e, SymFn(symtail), let__2586)


let__2587 := tmp3069
_ = let__2587

tmp3081 := GConsP(__e, let__2587)

var ifres3070 Obj

if True == tmp3081 {
tmp3071 := Call1(__e, SymFn(symhead), let__2587)


let__2588 := tmp3071
_ = let__2588

tmp3072 := Call1(__e, SymFn(symtail), let__2587)


let__2589 := tmp3072
_ = let__2589

tmp3073 := GCons(__e, let__2585, Nil)

tmp3074 := GCons(__e, symshen_4choicepoint_b, tmp3073)

tmp3075 := GCons(__e, tmp3074, Nil)

tmp3076 := GCons(__e, let__2588, tmp3075)

tmp3077 := GCons(__e, symwhere, tmp3076)

tmp3078 := Call2(__e, SymFn(sym_8p), let__2582, tmp3077)


tmp3079 := Call2(__e, SymFn(symshen_4comb), let__2589, tmp3078)


ifres3070 = tmp3079


} else {
tmp3080 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3070 = tmp3080


}

ifres3068 = ifres3070


} else {
tmp3082 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3068 = tmp3082


}

ifres3065 = ifres3068


} else {
tmp3084 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3065 = tmp3084


}

ifres3063 = ifres3065


} else {
tmp3086 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3063 = tmp3086


}

ifres3059 = ifres3063


}

let__2580 := ifres3059
_ = let__2580

tmp3111 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2580)


if True == tmp3111 {
tmp3089 := Call1(__e, SymFn(symshen_4_5patterns_6), V628)


let__2591 := tmp3089
_ = let__2591

tmp3107 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2591)


var ifres3090 Obj

if True == tmp3107 {
tmp3091 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3090 = tmp3091


} else {
tmp3092 := Call1(__e, SymFn(symshen_4_5_1out), let__2591)


let__2592 := tmp3092
_ = let__2592

tmp3093 := Call1(__e, SymFn(symshen_4in_1_6), let__2591)


let__2593 := tmp3093
_ = let__2593

tmp3106 := Call2(__e, SymFn(symshen_4hds_a_2), let__2593, sym_5_1)


var ifres3094 Obj

if True == tmp3106 {
tmp3095 := Call1(__e, SymFn(symtail), let__2593)


let__2594 := tmp3095
_ = let__2594

tmp3104 := GConsP(__e, let__2594)

var ifres3096 Obj

if True == tmp3104 {
tmp3097 := Call1(__e, SymFn(symhead), let__2594)


let__2595 := tmp3097
_ = let__2595

tmp3098 := Call1(__e, SymFn(symtail), let__2594)


let__2596 := tmp3098
_ = let__2596

tmp3099 := GCons(__e, let__2595, Nil)

tmp3100 := GCons(__e, symshen_4choicepoint_b, tmp3099)

tmp3101 := Call2(__e, SymFn(sym_8p), let__2592, tmp3100)


tmp3102 := Call2(__e, SymFn(symshen_4comb), let__2596, tmp3101)


ifres3096 = tmp3102


} else {
tmp3103 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3096 = tmp3103


}

ifres3094 = ifres3096


} else {
tmp3105 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3094 = tmp3105


}

ifres3090 = ifres3094


}

let__2590 := ifres3090
_ = let__2590

tmp3109 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2590)


if True == tmp3109 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2590
}


} else {
return let__2580
}


} else {
return let__2573
}


} else {
return let__2563
}


panic("unreachable")
})

tmp3116 := Call2(__e, ns2_1set, symshen_4_5rule_6, tmp3011)


_ = tmp3116

tmp3117 := MakeNative1(func(__e *ControlFlow, V663 Obj) Obj {
tmp3118 := Call1(__e, SymFn(symshen_4_5pattern_6), V663)


let__2598 := tmp3118
_ = let__2598

tmp3131 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2598)


var ifres3119 Obj

if True == tmp3131 {
tmp3120 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3119 = tmp3120


} else {
tmp3121 := Call1(__e, SymFn(symshen_4_5_1out), let__2598)


let__2599 := tmp3121
_ = let__2599

tmp3122 := Call1(__e, SymFn(symshen_4in_1_6), let__2598)


let__2600 := tmp3122
_ = let__2600

tmp3123 := Call1(__e, SymFn(symshen_4_5patterns_6), let__2600)


let__2601 := tmp3123
_ = let__2601

tmp3130 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2601)


var ifres3124 Obj

if True == tmp3130 {
tmp3125 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3124 = tmp3125


} else {
tmp3126 := Call1(__e, SymFn(symshen_4_5_1out), let__2601)


let__2602 := tmp3126
_ = let__2602

tmp3127 := Call1(__e, SymFn(symshen_4in_1_6), let__2601)


let__2603 := tmp3127
_ = let__2603

tmp3128 := GCons(__e, let__2599, let__2602)

tmp3129 := Call2(__e, SymFn(symshen_4comb), let__2603, tmp3128)


ifres3124 = tmp3129


}

ifres3119 = ifres3124


}

let__2597 := ifres3119
_ = let__2597

tmp3141 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2597)


if True == tmp3141 {
tmp3132 := Call1(__e, SymFn(sym_5e_6), V663)


let__2605 := tmp3132
_ = let__2605

tmp3137 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2605)


var ifres3133 Obj

if True == tmp3137 {
tmp3134 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3133 = tmp3134


} else {
tmp3135 := Call1(__e, SymFn(symshen_4in_1_6), let__2605)


let__2606 := tmp3135
_ = let__2606

tmp3136 := Call2(__e, SymFn(symshen_4comb), let__2606, Nil)


ifres3133 = tmp3136


}

let__2604 := ifres3133
_ = let__2604

tmp3139 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2604)


if True == tmp3139 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2604
}


} else {
return let__2597
}


panic("unreachable")
})

tmp3142 := Call2(__e, ns2_1set, symshen_4_5patterns_6, tmp3117)


_ = tmp3142

tmp3143 := MakeNative1(func(__e *ControlFlow, V674 Obj) Obj {
tmp3164 := Call1(__e, SymFn(symshen_4ccons_2), V674)


var ifres3144 Obj

if True == tmp3164 {
tmp3145 := Call1(__e, SymFn(symhead), V674)


let__2608 := tmp3145
_ = let__2608

tmp3146 := Call1(__e, SymFn(symtail), V674)


let__2609 := tmp3146
_ = let__2609

tmp3162 := Call2(__e, SymFn(symshen_4hds_a_2), let__2608, symvector)


var ifres3147 Obj

if True == tmp3162 {
tmp3148 := Call1(__e, SymFn(symtail), let__2608)


let__2610 := tmp3148
_ = let__2610

tmp3160 := Call2(__e, SymFn(symshen_4hds_a_2), let__2610, MakeInteger(0))


var ifres3149 Obj

if True == tmp3160 {
tmp3150 := Call1(__e, SymFn(symtail), let__2610)


let__2611 := tmp3150
_ = let__2611

tmp3151 := Call1(__e, SymFn(sym_5end_6), let__2611)


let__2612 := tmp3151
_ = let__2612

tmp3158 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2612)


var ifres3152 Obj

if True == tmp3158 {
tmp3153 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3152 = tmp3153


} else {
tmp3154 := Call1(__e, SymFn(symshen_4in_1_6), let__2612)


let__2613 := tmp3154
_ = let__2613

tmp3155 := GCons(__e, MakeInteger(0), Nil)

tmp3156 := GCons(__e, symvector, tmp3155)

tmp3157 := Call2(__e, SymFn(symshen_4comb), let__2609, tmp3156)


ifres3152 = tmp3157


}

ifres3149 = ifres3152


} else {
tmp3159 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3149 = tmp3159


}

ifres3147 = ifres3149


} else {
tmp3161 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3147 = tmp3161


}

ifres3144 = ifres3147


} else {
tmp3163 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3144 = tmp3163


}

let__2607 := ifres3144
_ = let__2607

tmp3187 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2607)


if True == tmp3187 {
tmp3174 := GConsP(__e, V674)

var ifres3165 Obj

if True == tmp3174 {
tmp3166 := Call1(__e, SymFn(symhead), V674)


let__2615 := tmp3166
_ = let__2615

tmp3167 := Call1(__e, SymFn(symtail), V674)


let__2616 := tmp3167
_ = let__2616

tmp3172 := GConsP(__e, let__2615)

var ifres3168 Obj

if True == tmp3172 {
tmp3169 := Call1(__e, SymFn(symshen_4compound_1pattern), let__2615)


tmp3170 := Call2(__e, SymFn(symshen_4comb), let__2616, tmp3169)


ifres3168 = tmp3170


} else {
tmp3171 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3168 = tmp3171


}

ifres3165 = ifres3168


} else {
tmp3173 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3165 = tmp3173


}

let__2614 := ifres3165
_ = let__2614

tmp3185 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2614)


if True == tmp3185 {
tmp3175 := Call1(__e, SymFn(symshen_4_5simple_1pattern_6), V674)


let__2618 := tmp3175
_ = let__2618

tmp3181 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2618)


var ifres3176 Obj

if True == tmp3181 {
tmp3177 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3176 = tmp3177


} else {
tmp3178 := Call1(__e, SymFn(symshen_4_5_1out), let__2618)


let__2619 := tmp3178
_ = let__2619

tmp3179 := Call1(__e, SymFn(symshen_4in_1_6), let__2618)


let__2620 := tmp3179
_ = let__2620

tmp3180 := Call2(__e, SymFn(symshen_4comb), let__2620, let__2619)


ifres3176 = tmp3180


}

let__2617 := ifres3176
_ = let__2617

tmp3183 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2617)


if True == tmp3183 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2617
}


} else {
return let__2614
}


} else {
return let__2607
}


panic("unreachable")
})

tmp3188 := Call2(__e, ns2_1set, symshen_4_5pattern_6, tmp3143)


_ = tmp3188

tmp3189 := MakeNative1(func(__e *ControlFlow, V689 Obj) Obj {
tmp3198 := GConsP(__e, V689)

var ifres3190 Obj

if True == tmp3198 {
tmp3191 := Call1(__e, SymFn(symhead), V689)


let__2622 := tmp3191
_ = let__2622

tmp3192 := Call1(__e, SymFn(symtail), V689)


let__2623 := tmp3192
_ = let__2623

tmp3196 := Call1(__e, SymFn(symshen_4constructor_2), let__2622)


var ifres3193 Obj

if True == tmp3196 {
tmp3194 := Call2(__e, SymFn(symshen_4comb), let__2623, let__2622)


ifres3193 = tmp3194


} else {
tmp3195 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3193 = tmp3195


}

ifres3190 = ifres3193


} else {
tmp3197 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3190 = tmp3197


}

let__2621 := ifres3190
_ = let__2621

tmp3200 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2621)


if True == tmp3200 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2621
}


panic("unreachable")
})

tmp3201 := Call2(__e, ns2_1set, symshen_4_5constructor_6, tmp3189)


_ = tmp3201

tmp3202 := MakeNative1(func(__e *ControlFlow, V693 Obj) Obj {
tmp3203 := GCons(__e, sym_8v, Nil)

tmp3204 := GCons(__e, sym_8s, tmp3203)

tmp3205 := GCons(__e, sym_8p, tmp3204)

tmp3206 := GCons(__e, symcons, tmp3205)

return TailCall2(__e, SymFn(symelement_2), V693, tmp3206)


panic("unreachable")
})

tmp3207 := Call2(__e, ns2_1set, symshen_4constructor_2, tmp3202)


_ = tmp3207

tmp3208 := MakeNative1(func(__e *ControlFlow, V694 Obj) Obj {
tmp3209 := Call3(__e, SymFn(symshen_4app), V694, MakeString(" is not a legitimate constructor\n"), symshen_4r)


return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp3209)


panic("unreachable")
})

tmp3210 := Call2(__e, ns2_1set, symshen_4constructor_1error, tmp3208)


_ = tmp3210

tmp3211 := MakeNative2(func(__e *ControlFlow, V695, V696 Obj) Obj {
tmp3212 := GValue(__e, symshen_4_dcustom_1pattern_1compiler_d)

let__2624 := tmp3212
_ = let__2624

tmp3215 := GEq(__e, let__2624, False)

if True == tmp3215 {
return TailCall1(__e, SymFn(symthaw), V696)
} else {
tmp3213 := Call1(__e, let__2624, V695)


return TailCall1(__e, tmp3213, V696)


}


panic("unreachable")
})

tmp3216 := Call2(__e, ns2_1set, symshen_4custom_1pattern_1compiler, tmp3211)


_ = tmp3216

tmp3217 := MakeNative1(func(__e *ControlFlow, V698 Obj) Obj {
tmp3218 := GValue(__e, symshen_4_dcustom_1pattern_1reducer_d)

let__2625 := tmp3218
_ = let__2625

tmp3220 := GEq(__e, let__2625, False)

if True == tmp3220 {
return TailCall0(__e, SymFn(symfail))
} else {
return TailCall1(__e, let__2625, V698)
}


panic("unreachable")
})

tmp3221 := Call2(__e, ns2_1set, symshen_4custom_1pattern_1reducer, tmp3217)


_ = tmp3221

tmp3222 := MakeNative1(func(__e *ControlFlow, V700 Obj) Obj {
tmp3223 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, SymFn(symshen_4compound_1pattern_1h), V700)
panic("unreachable")
})

return TailCall2(__e, SymFn(symshen_4custom_1pattern_1compiler), V700, tmp3223)


panic("unreachable")
})

tmp3224 := Call2(__e, ns2_1set, symshen_4compound_1pattern, tmp3222)


_ = tmp3224

tmp3225 := MakeNative1(func(__e *ControlFlow, V701 Obj) Obj {
tmp3252 := GConsP(__e, V701)

var ifres3237 Obj

if True == tmp3252 {
tmp3250 := GTl(__e, V701)

tmp3251 := GConsP(__e, tmp3250)

var ifres3238 Obj

if True == tmp3251 {
tmp3247 := GTl(__e, V701)

tmp3248 := GTl(__e, tmp3247)

tmp3249 := GConsP(__e, tmp3248)

var ifres3239 Obj

if True == tmp3249 {
tmp3243 := GTl(__e, V701)

tmp3244 := GTl(__e, tmp3243)

tmp3245 := GTl(__e, tmp3244)

tmp3246 := GEq(__e, Nil, tmp3245)

var ifres3240 Obj

if True == tmp3246 {
tmp3241 := GHd(__e, V701)

tmp3242 := Call1(__e, SymFn(symshen_4constructor_2), tmp3241)


ifres3240 = tmp3242


} else {
ifres3240 = False


}

ifres3239 = ifres3240


} else {
ifres3239 = False


}

ifres3238 = ifres3239


} else {
ifres3238 = False


}

ifres3237 = ifres3238


} else {
ifres3237 = False


}

if True == ifres3237 {
tmp3226 := GHd(__e, V701)

tmp3227 := GTl(__e, V701)

tmp3228 := GHd(__e, tmp3227)

tmp3229 := Call1(__e, SymFn(symshen_4compile_1pattern_1fragment), tmp3228)


tmp3230 := GTl(__e, V701)

tmp3231 := GTl(__e, tmp3230)

tmp3232 := GHd(__e, tmp3231)

tmp3233 := Call1(__e, SymFn(symshen_4compile_1pattern_1fragment), tmp3232)


tmp3234 := GCons(__e, tmp3233, Nil)

tmp3235 := GCons(__e, tmp3229, tmp3234)

return GCons(__e, tmp3226, tmp3235)


} else {
return TailCall1(__e, SymFn(symshen_4constructor_1error), V701)
}


panic("unreachable")
})

tmp3253 := Call2(__e, ns2_1set, symshen_4compound_1pattern_1h, tmp3225)


_ = tmp3253

tmp3254 := MakeNative1(func(__e *ControlFlow, V702 Obj) Obj {
tmp3279 := GConsP(__e, V702)

var ifres3265 Obj

if True == tmp3279 {
tmp3277 := GHd(__e, V702)

tmp3278 := GEq(__e, symvector, tmp3277)

var ifres3266 Obj

if True == tmp3278 {
tmp3275 := GTl(__e, V702)

tmp3276 := GConsP(__e, tmp3275)

var ifres3267 Obj

if True == tmp3276 {
tmp3272 := GTl(__e, V702)

tmp3273 := GHd(__e, tmp3272)

tmp3274 := GEq(__e, MakeInteger(0), tmp3273)

var ifres3268 Obj

if True == tmp3274 {
tmp3269 := GTl(__e, V702)

tmp3270 := GTl(__e, tmp3269)

tmp3271 := GEq(__e, Nil, tmp3270)

ifres3268 = tmp3271


} else {
ifres3268 = False


}

ifres3267 = ifres3268


} else {
ifres3267 = False


}

ifres3266 = ifres3267


} else {
ifres3266 = False


}

ifres3265 = ifres3266


} else {
ifres3265 = False


}

if True == ifres3265 {
return V702
} else {
tmp3263 := GConsP(__e, V702)

if True == tmp3263 {
return TailCall1(__e, SymFn(symshen_4compound_1pattern), V702)
} else {
tmp3261 := GEq(__e, V702, sym__)

if True == tmp3261 {
return TailCall1(__e, SymFn(symgensym), symY)
} else {
tmp3256 := GCons(__e, sym_5_1, Nil)

tmp3257 := GCons(__e, sym_1_6, tmp3256)

tmp3258 := Call2(__e, SymFn(symelement_2), V702, tmp3257)


if True == GNot(__e, tmp3258) {
return V702
} else {
return TailCall1(__e, SymFn(symshen_4constructor_1error), V702)
}


}


}


}


panic("unreachable")
})

tmp3280 := Call2(__e, ns2_1set, symshen_4compile_1pattern_1fragment, tmp3254)


_ = tmp3280

tmp3281 := MakeNative1(func(__e *ControlFlow, V707 Obj) Obj {
tmp3302 := GConsP(__e, V707)

var ifres3283 Obj

if True == tmp3302 {
tmp3300 := GHd(__e, V707)

tmp3301 := GEq(__e, sym_8p, tmp3300)

var ifres3284 Obj

if True == tmp3301 {
tmp3298 := GTl(__e, V707)

tmp3299 := GConsP(__e, tmp3298)

var ifres3285 Obj

if True == tmp3299 {
tmp3295 := GTl(__e, V707)

tmp3296 := GHd(__e, tmp3295)

tmp3297 := GEq(__e, symshen_4custom_1pattern, tmp3296)

var ifres3286 Obj

if True == tmp3297 {
tmp3292 := GTl(__e, V707)

tmp3293 := GTl(__e, tmp3292)

tmp3294 := GConsP(__e, tmp3293)

var ifres3287 Obj

if True == tmp3294 {
tmp3288 := GTl(__e, V707)

tmp3289 := GTl(__e, tmp3288)

tmp3290 := GTl(__e, tmp3289)

tmp3291 := GEq(__e, Nil, tmp3290)

ifres3287 = tmp3291


} else {
ifres3287 = False


}

ifres3286 = ifres3287


} else {
ifres3286 = False


}

ifres3285 = ifres3286


} else {
ifres3285 = False


}

ifres3284 = ifres3285


} else {
ifres3284 = False


}

ifres3283 = ifres3284


} else {
ifres3283 = False


}

if True == ifres3283 {
return True
} else {
return False
}


panic("unreachable")
})

tmp3303 := Call2(__e, ns2_1set, symshen_4custom_1pattern_2, tmp3281)


_ = tmp3303

tmp3304 := MakeNative1(func(__e *ControlFlow, V710 Obj) Obj {
tmp3327 := GConsP(__e, V710)

var ifres3308 Obj

if True == tmp3327 {
tmp3325 := GHd(__e, V710)

tmp3326 := GEq(__e, sym_8p, tmp3325)

var ifres3309 Obj

if True == tmp3326 {
tmp3323 := GTl(__e, V710)

tmp3324 := GConsP(__e, tmp3323)

var ifres3310 Obj

if True == tmp3324 {
tmp3320 := GTl(__e, V710)

tmp3321 := GHd(__e, tmp3320)

tmp3322 := GEq(__e, symshen_4custom_1pattern, tmp3321)

var ifres3311 Obj

if True == tmp3322 {
tmp3317 := GTl(__e, V710)

tmp3318 := GTl(__e, tmp3317)

tmp3319 := GConsP(__e, tmp3318)

var ifres3312 Obj

if True == tmp3319 {
tmp3313 := GTl(__e, V710)

tmp3314 := GTl(__e, tmp3313)

tmp3315 := GTl(__e, tmp3314)

tmp3316 := GEq(__e, Nil, tmp3315)

ifres3312 = tmp3316


} else {
ifres3312 = False


}

ifres3311 = ifres3312


} else {
ifres3311 = False


}

ifres3310 = ifres3311


} else {
ifres3310 = False


}

ifres3309 = ifres3310


} else {
ifres3309 = False


}

ifres3308 = ifres3309


} else {
ifres3308 = False


}

if True == ifres3308 {
tmp3305 := GTl(__e, V710)

tmp3306 := GTl(__e, tmp3305)

return GHd(__e, tmp3306)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.custom-pattern-body"))
}


panic("unreachable")
})

tmp3328 := Call2(__e, ns2_1set, symshen_4custom_1pattern_1body, tmp3304)


_ = tmp3328

tmp3329 := MakeNative1(func(__e *ControlFlow, V711 Obj) Obj {
tmp3339 := GConsP(__e, V711)

var ifres3330 Obj

if True == tmp3339 {
tmp3331 := Call1(__e, SymFn(symhead), V711)


let__2627 := tmp3331
_ = let__2627

tmp3332 := Call1(__e, SymFn(symtail), V711)


let__2628 := tmp3332
_ = let__2628

tmp3337 := GEq(__e, let__2627, sym__)

var ifres3333 Obj

if True == tmp3337 {
tmp3334 := Call1(__e, SymFn(symgensym), symY)


tmp3335 := Call2(__e, SymFn(symshen_4comb), let__2628, tmp3334)


ifres3333 = tmp3335


} else {
tmp3336 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3333 = tmp3336


}

ifres3330 = ifres3333


} else {
tmp3338 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3330 = tmp3338


}

let__2626 := ifres3330
_ = let__2626

tmp3355 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2626)


if True == tmp3355 {
tmp3351 := GConsP(__e, V711)

var ifres3340 Obj

if True == tmp3351 {
tmp3341 := Call1(__e, SymFn(symhead), V711)


let__2630 := tmp3341
_ = let__2630

tmp3342 := Call1(__e, SymFn(symtail), V711)


let__2631 := tmp3342
_ = let__2631

tmp3346 := GCons(__e, sym_5_1, Nil)

tmp3347 := GCons(__e, sym_1_6, tmp3346)

tmp3348 := Call2(__e, SymFn(symelement_2), let__2630, tmp3347)


tmp3349 := GNot(__e, tmp3348)

var ifres3343 Obj

if True == tmp3349 {
tmp3344 := Call2(__e, SymFn(symshen_4comb), let__2631, let__2630)


ifres3343 = tmp3344


} else {
tmp3345 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3343 = tmp3345


}

ifres3340 = ifres3343


} else {
tmp3350 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3340 = tmp3350


}

let__2629 := ifres3340
_ = let__2629

tmp3353 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2629)


if True == tmp3353 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2629
}


} else {
return let__2626
}


panic("unreachable")
})

tmp3356 := Call2(__e, ns2_1set, symshen_4_5simple_1pattern_6, tmp3329)


_ = tmp3356

tmp3357 := MakeNative1(func(__e *ControlFlow, V718 Obj) Obj {
tmp3358 := Call1(__e, SymFn(symshen_4_5pattern_6), V718)


let__2633 := tmp3358
_ = let__2633

tmp3364 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2633)


var ifres3359 Obj

if True == tmp3364 {
tmp3360 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3359 = tmp3360


} else {
tmp3361 := Call1(__e, SymFn(symshen_4_5_1out), let__2633)


let__2634 := tmp3361
_ = let__2634

tmp3362 := Call1(__e, SymFn(symshen_4in_1_6), let__2633)


let__2635 := tmp3362
_ = let__2635

tmp3363 := Call2(__e, SymFn(symshen_4comb), let__2635, let__2634)


ifres3359 = tmp3363


}

let__2632 := ifres3359
_ = let__2632

tmp3366 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2632)


if True == tmp3366 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2632
}


panic("unreachable")
})

tmp3367 := Call2(__e, ns2_1set, symshen_4_5pattern1_6, tmp3357)


_ = tmp3367

tmp3368 := MakeNative1(func(__e *ControlFlow, V723 Obj) Obj {
tmp3369 := Call1(__e, SymFn(symshen_4_5pattern_6), V723)


let__2637 := tmp3369
_ = let__2637

tmp3375 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2637)


var ifres3370 Obj

if True == tmp3375 {
tmp3371 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3370 = tmp3371


} else {
tmp3372 := Call1(__e, SymFn(symshen_4_5_1out), let__2637)


let__2638 := tmp3372
_ = let__2638

tmp3373 := Call1(__e, SymFn(symshen_4in_1_6), let__2637)


let__2639 := tmp3373
_ = let__2639

tmp3374 := Call2(__e, SymFn(symshen_4comb), let__2639, let__2638)


ifres3370 = tmp3374


}

let__2636 := ifres3370
_ = let__2636

tmp3377 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2636)


if True == tmp3377 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2636
}


panic("unreachable")
})

tmp3378 := Call2(__e, ns2_1set, symshen_4_5pattern2_6, tmp3368)


_ = tmp3378

tmp3379 := MakeNative1(func(__e *ControlFlow, V728 Obj) Obj {
tmp3380 := PrimGuard1(__e, symabsvector, PrimAbsvector, MakeInteger(2))

let__2640 := tmp3380
_ = let__2640

tmp3381 := GVecSet(__e, let__2640, MakeInteger(0), symshen_4printF)

let__2641 := tmp3381
_ = let__2641

tmp3382 := PrimGuard1(__e, symstr, PrimStr, V728)

tmp3383 := Call2(__e, SymFn(sym_8s), tmp3382, MakeString(")"))


tmp3384 := Call2(__e, SymFn(sym_8s), MakeString(" "), tmp3383)


tmp3385 := Call2(__e, SymFn(sym_8s), MakeString("n"), tmp3384)


tmp3386 := Call2(__e, SymFn(sym_8s), MakeString("f"), tmp3385)


tmp3387 := Call2(__e, SymFn(sym_8s), MakeString("("), tmp3386)


tmp3388 := GVecSet(__e, let__2641, MakeInteger(1), tmp3387)

let__2642 := tmp3388
_ = let__2642

return let__2642


panic("unreachable")
})

tmp3389 := Call2(__e, ns2_1set, symshen_4fn_1print, tmp3379)


_ = tmp3389

tmp3390 := MakeNative1(func(__e *ControlFlow, V732 Obj) Obj {
return GVecGet(__e, V732, MakeInteger(1))
panic("unreachable")
})

tmp3391 := Call2(__e, ns2_1set, symshen_4printF, tmp3390)


_ = tmp3391

tmp3392 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V737 := __self1
_ = V737
V738 := __self2
_ = V738
tmp3413 := GConsP(__e, V738)

var ifres3410 Obj

if True == tmp3413 {
tmp3411 := GTl(__e, V738)

tmp3412 := GEq(__e, Nil, tmp3411)

ifres3410 = tmp3412


} else {
ifres3410 = False


}

if True == ifres3410 {
tmp3393 := GHd(__e, V738)

return TailCall1(__e, SymFn(symlength), tmp3393)


} else {
tmp3408 := GConsP(__e, V738)

var ifres3398 Obj

if True == tmp3408 {
tmp3406 := GTl(__e, V738)

tmp3407 := GConsP(__e, tmp3406)

var ifres3399 Obj

if True == tmp3407 {
tmp3400 := GHd(__e, V738)

tmp3401 := Call1(__e, SymFn(symlength), tmp3400)


tmp3402 := GTl(__e, V738)

tmp3403 := GHd(__e, tmp3402)

tmp3404 := Call1(__e, SymFn(symlength), tmp3403)


tmp3405 := GEq(__e, tmp3401, tmp3404)

ifres3399 = tmp3405


} else {
ifres3399 = False


}

ifres3398 = ifres3399


} else {
ifres3398 = False


}

if True == ifres3398 {
tmp3394 := GTl(__e, V738)

if SymFn(symshen_4arity_1chk) == __self {
__self1, __self2 = V737, tmp3394
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4arity_1chk), V737, tmp3394)


} else {
tmp3395 := Call3(__e, SymFn(symshen_4app), V737, MakeString("\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("arity error in "), tmp3395))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("arity error in "), tmp3395)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


}


panic("unreachable")
})
return __self
})()

tmp3414 := Call2(__e, ns2_1set, symshen_4arity_1chk, tmp3392)


_ = tmp3414

tmp3415 := MakeNative2(func(__e *ControlFlow, V739, V740 Obj) Obj {
tmp3421 := Call1(__e, SymFn(symtuple_2), V740)


if True == tmp3421 {
tmp3416 := Call1(__e, SymFn(symfst), V740)


tmp3417 := Call1(__e, SymFn(symshen_4extract_1vars), tmp3416)


tmp3418 := Call1(__e, SymFn(symsnd), V740)


tmp3419 := Call2(__e, SymFn(symshen_4find_1free_1vars), tmp3417, tmp3418)


return TailCall2(__e, SymFn(symshen_4free_1variable_1error_1message), V739, tmp3419)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4free_1var_1chk)
}


panic("unreachable")
})

tmp3422 := Call2(__e, ns2_1set, symshen_4free_1var_1chk, tmp3415)


_ = tmp3422

tmp3423 := MakeNative2(func(__e *ControlFlow, V741, V742 Obj) Obj {
tmp3435 := Call1(__e, SymFn(symempty_2), V742)


if True == tmp3435 {
return symshen_4skip
} else {
tmp3424 := Call3(__e, SymFn(symshen_4app), V741, MakeString(":"), symshen_4a)


tmp3425 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("free variables in "), tmp3424)

tmp3426 := Call0(__e, SymFn(symstoutput))


tmp3427 := Call2(__e, SymFn(sympr), tmp3425, tmp3426)


_ = tmp3427

tmp3428 := MakeNative1(func(__e *ControlFlow, Z743 Obj) Obj {
tmp3429 := Call3(__e, SymFn(symshen_4app), Z743, MakeString(""), symshen_4a)


tmp3430 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" "), tmp3429)

tmp3431 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp3430, tmp3431)


panic("unreachable")
})

tmp3432 := Call2(__e, SymFn(symshen_4for_1each), tmp3428, V742)


_ = tmp3432

tmp3433 := Call1(__e, SymFn(symnl), MakeInteger(1))


_ = tmp3433

return TailCall0(__e, SymFn(symabort))


}


panic("unreachable")
})

tmp3436 := Call2(__e, ns2_1set, symshen_4free_1variable_1error_1message, tmp3423)


_ = tmp3436

tmp3437 := MakeNative1(func(__e *ControlFlow, V746 Obj) Obj {
tmp3445 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V746)

if True == tmp3445 {
return GCons(__e, V746, Nil)
} else {
tmp3443 := GConsP(__e, V746)

if True == tmp3443 {
tmp3438 := GHd(__e, V746)

tmp3439 := Call1(__e, SymFn(symshen_4extract_1vars), tmp3438)


tmp3440 := GTl(__e, V746)

tmp3441 := Call1(__e, SymFn(symshen_4extract_1vars), tmp3440)


return TailCall2(__e, SymFn(symunion), tmp3439, tmp3441)


} else {
return Nil
}


}


panic("unreachable")
})

tmp3446 := Call2(__e, ns2_1set, symshen_4extract_1vars, tmp3437)


_ = tmp3446

tmp3447 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V751 := __self1
_ = V751
V752 := __self2
_ = V752
tmp3525 := GConsP(__e, V752)

var ifres3515 Obj

if True == tmp3525 {
tmp3523 := GHd(__e, V752)

tmp3524 := GEq(__e, symprotect, tmp3523)

var ifres3516 Obj

if True == tmp3524 {
tmp3521 := GTl(__e, V752)

tmp3522 := GConsP(__e, tmp3521)

var ifres3517 Obj

if True == tmp3522 {
tmp3518 := GTl(__e, V752)

tmp3519 := GTl(__e, tmp3518)

tmp3520 := GEq(__e, Nil, tmp3519)

ifres3517 = tmp3520


} else {
ifres3517 = False


}

ifres3516 = ifres3517


} else {
ifres3516 = False


}

ifres3515 = ifres3516


} else {
ifres3515 = False


}

if True == ifres3515 {
return Nil
} else {
tmp3513 := GConsP(__e, V752)

var ifres3492 Obj

if True == tmp3513 {
tmp3511 := GHd(__e, V752)

tmp3512 := GEq(__e, symlet, tmp3511)

var ifres3493 Obj

if True == tmp3512 {
tmp3509 := GTl(__e, V752)

tmp3510 := GConsP(__e, tmp3509)

var ifres3494 Obj

if True == tmp3510 {
tmp3506 := GTl(__e, V752)

tmp3507 := GTl(__e, tmp3506)

tmp3508 := GConsP(__e, tmp3507)

var ifres3495 Obj

if True == tmp3508 {
tmp3502 := GTl(__e, V752)

tmp3503 := GTl(__e, tmp3502)

tmp3504 := GTl(__e, tmp3503)

tmp3505 := GConsP(__e, tmp3504)

var ifres3496 Obj

if True == tmp3505 {
tmp3497 := GTl(__e, V752)

tmp3498 := GTl(__e, tmp3497)

tmp3499 := GTl(__e, tmp3498)

tmp3500 := GTl(__e, tmp3499)

tmp3501 := GEq(__e, Nil, tmp3500)

ifres3496 = tmp3501


} else {
ifres3496 = False


}

ifres3495 = ifres3496


} else {
ifres3495 = False


}

ifres3494 = ifres3495


} else {
ifres3494 = False


}

ifres3493 = ifres3494


} else {
ifres3493 = False


}

ifres3492 = ifres3493


} else {
ifres3492 = False


}

if True == ifres3492 {
tmp3448 := GTl(__e, V752)

tmp3449 := GTl(__e, tmp3448)

tmp3450 := GHd(__e, tmp3449)

tmp3451 := Call2(__e, SymFn(symshen_4find_1free_1vars), V751, tmp3450)


tmp3452 := GTl(__e, V752)

tmp3453 := GHd(__e, tmp3452)

tmp3454 := GCons(__e, tmp3453, V751)

tmp3455 := GTl(__e, V752)

tmp3456 := GTl(__e, tmp3455)

tmp3457 := GTl(__e, tmp3456)

tmp3458 := GHd(__e, tmp3457)

tmp3459 := Call2(__e, SymFn(symshen_4find_1free_1vars), tmp3454, tmp3458)


return TailCall2(__e, SymFn(symunion), tmp3451, tmp3459)


} else {
tmp3490 := GConsP(__e, V752)

var ifres3475 Obj

if True == tmp3490 {
tmp3488 := GHd(__e, V752)

tmp3489 := GEq(__e, symlambda, tmp3488)

var ifres3476 Obj

if True == tmp3489 {
tmp3486 := GTl(__e, V752)

tmp3487 := GConsP(__e, tmp3486)

var ifres3477 Obj

if True == tmp3487 {
tmp3483 := GTl(__e, V752)

tmp3484 := GTl(__e, tmp3483)

tmp3485 := GConsP(__e, tmp3484)

var ifres3478 Obj

if True == tmp3485 {
tmp3479 := GTl(__e, V752)

tmp3480 := GTl(__e, tmp3479)

tmp3481 := GTl(__e, tmp3480)

tmp3482 := GEq(__e, Nil, tmp3481)

ifres3478 = tmp3482


} else {
ifres3478 = False


}

ifres3477 = ifres3478


} else {
ifres3477 = False


}

ifres3476 = ifres3477


} else {
ifres3476 = False


}

ifres3475 = ifres3476


} else {
ifres3475 = False


}

if True == ifres3475 {
tmp3460 := GTl(__e, V752)

tmp3461 := GHd(__e, tmp3460)

tmp3462 := GCons(__e, tmp3461, V751)

tmp3463 := GTl(__e, V752)

tmp3464 := GTl(__e, tmp3463)

tmp3465 := GHd(__e, tmp3464)

if SymFn(symshen_4find_1free_1vars) == __self {
__self1, __self2 = tmp3462, tmp3465
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4find_1free_1vars), tmp3462, tmp3465)


} else {
tmp3473 := GConsP(__e, V752)

if True == tmp3473 {
tmp3466 := GHd(__e, V752)

tmp3467 := Call2(__e, SymFn(symshen_4find_1free_1vars), V751, tmp3466)


tmp3468 := GTl(__e, V752)

tmp3469 := Call2(__e, SymFn(symshen_4find_1free_1vars), V751, tmp3468)


return TailCall2(__e, SymFn(symunion), tmp3467, tmp3469)


} else {
tmp3471 := Call2(__e, SymFn(symshen_4free_1variable_2), V752, V751)


if True == tmp3471 {
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

tmp3526 := Call2(__e, ns2_1set, symshen_4find_1free_1vars, tmp3447)


_ = tmp3526

tmp3527 := MakeNative2(func(__e *ControlFlow, V753, V754 Obj) Obj {
tmp3530 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V753)

if True == tmp3530 {
tmp3528 := Call2(__e, SymFn(symelement_2), V753, V754)


return GNot(__e, tmp3528)


} else {
return False
}


panic("unreachable")
})

tmp3531 := Call2(__e, ns2_1set, symshen_4free_1variable_2, tmp3527)


_ = tmp3531

tmp3532 := MakeNative2(func(__e *ControlFlow, V755, V756 Obj) Obj {
tmp3533 := GValue(__e, symshen_4_duserdefs_d)

tmp3534 := Call2(__e, SymFn(symadjoin), V755, tmp3533)


tmp3535 := PrimGuard2(__e, symset, PrimSet, symshen_4_duserdefs_d, tmp3534)

_ = tmp3535

tmp3536 := GValue(__e, sym_dproperty_1vector_d)

return TailCall4(__e, SymFn(symput), V755, symshen_4source, V756, tmp3536)


panic("unreachable")
})

tmp3537 := Call2(__e, ns2_1set, symshen_4record_1kl, tmp3532)


_ = tmp3537

tmp3538 := MakeNative3(func(__e *ControlFlow, V757, V758, V759 Obj) Obj {
tmp3539 := Call1(__e, SymFn(symshen_4parameters), V759)


let__2643 := tmp3539
_ = let__2643

tmp3540 := Call2(__e, SymFn(symshen_4kl_1body), V758, let__2643)


tmp3541 := Call2(__e, SymFn(symshen_4scan_1body), V757, tmp3540)


let__2644 := tmp3541
_ = let__2644

tmp3542 := Call1(__e, SymFn(symshen_4cond_1form), let__2644)


tmp3543 := GCons(__e, tmp3542, Nil)

tmp3544 := GCons(__e, let__2643, tmp3543)

tmp3545 := GCons(__e, V757, tmp3544)

tmp3546 := GCons(__e, symdefun, tmp3545)

let__2645 := tmp3546
_ = let__2645

return let__2645


panic("unreachable")
})

tmp3547 := Call2(__e, ns2_1set, symshen_4compile_1to_1kl, tmp3538)


_ = tmp3547

tmp3548 := MakeNative1(func(__e *ControlFlow, V763 Obj) Obj {
tmp3553 := GEq(__e, MakeInteger(0), V763)

if True == tmp3553 {
return Nil
} else {
tmp3549 := Call1(__e, SymFn(symgensym), symV)


tmp3551 := Call1(__e, SymFn(symshen_4parameters), GSub(__e, V763, MakeInteger(1)))


return GCons(__e, tmp3549, tmp3551)


}


panic("unreachable")
})

tmp3554 := Call2(__e, ns2_1set, symshen_4parameters, tmp3548)


_ = tmp3554

tmp3555 := MakeNative1(func(__e *ControlFlow, V766 Obj) Obj {
tmp3575 := GConsP(__e, V766)

var ifres3559 Obj

if True == tmp3575 {
tmp3573 := GHd(__e, V766)

tmp3574 := GConsP(__e, tmp3573)

var ifres3560 Obj

if True == tmp3574 {
tmp3570 := GHd(__e, V766)

tmp3571 := GHd(__e, tmp3570)

tmp3572 := GEq(__e, True, tmp3571)

var ifres3561 Obj

if True == tmp3572 {
tmp3567 := GHd(__e, V766)

tmp3568 := GTl(__e, tmp3567)

tmp3569 := GConsP(__e, tmp3568)

var ifres3562 Obj

if True == tmp3569 {
tmp3563 := GHd(__e, V766)

tmp3564 := GTl(__e, tmp3563)

tmp3565 := GTl(__e, tmp3564)

tmp3566 := GEq(__e, Nil, tmp3565)

ifres3562 = tmp3566


} else {
ifres3562 = False


}

ifres3561 = ifres3562


} else {
ifres3561 = False


}

ifres3560 = ifres3561


} else {
ifres3560 = False


}

ifres3559 = ifres3560


} else {
ifres3559 = False


}

if True == ifres3559 {
tmp3556 := GHd(__e, V766)

tmp3557 := GTl(__e, tmp3556)

return GHd(__e, tmp3557)


} else {
return GCons(__e, symcond, V766)
}


panic("unreachable")
})

tmp3576 := Call2(__e, ns2_1set, symshen_4cond_1form, tmp3555)


_ = tmp3576

tmp3577 := MakeNative2(func(__e *ControlFlow, V775, V776 Obj) Obj {
tmp3616 := GEq(__e, Nil, V776)

if True == tmp3616 {
tmp3578 := GCons(__e, V775, Nil)

tmp3579 := GCons(__e, symshen_4f_1error, tmp3578)

tmp3580 := GCons(__e, tmp3579, Nil)

tmp3581 := GCons(__e, True, tmp3580)

return GCons(__e, tmp3581, Nil)


} else {
tmp3614 := GConsP(__e, V776)

var ifres3611 Obj

if True == tmp3614 {
tmp3612 := GHd(__e, V776)

tmp3613 := Call1(__e, SymFn(symshen_4choicepoint_2), tmp3612)


ifres3611 = tmp3613


} else {
ifres3611 = False


}

if True == ifres3611 {
tmp3582 := Call1(__e, SymFn(symgensym), symFreeze)


tmp3583 := Call1(__e, SymFn(symgensym), symResult)


tmp3584 := GHd(__e, V776)

tmp3585 := GTl(__e, V776)

return TailCall5(__e, SymFn(symshen_4choicepoint), V775, tmp3582, tmp3583, tmp3584, tmp3585)


} else {
tmp3609 := GConsP(__e, V776)

var ifres3593 Obj

if True == tmp3609 {
tmp3607 := GHd(__e, V776)

tmp3608 := GConsP(__e, tmp3607)

var ifres3594 Obj

if True == tmp3608 {
tmp3604 := GHd(__e, V776)

tmp3605 := GHd(__e, tmp3604)

tmp3606 := GEq(__e, True, tmp3605)

var ifres3595 Obj

if True == tmp3606 {
tmp3601 := GHd(__e, V776)

tmp3602 := GTl(__e, tmp3601)

tmp3603 := GConsP(__e, tmp3602)

var ifres3596 Obj

if True == tmp3603 {
tmp3597 := GHd(__e, V776)

tmp3598 := GTl(__e, tmp3597)

tmp3599 := GTl(__e, tmp3598)

tmp3600 := GEq(__e, Nil, tmp3599)

ifres3596 = tmp3600


} else {
ifres3596 = False


}

ifres3595 = ifres3596


} else {
ifres3595 = False


}

ifres3594 = ifres3595


} else {
ifres3594 = False


}

ifres3593 = ifres3594


} else {
ifres3593 = False


}

if True == ifres3593 {
tmp3586 := GHd(__e, V776)

return GCons(__e, tmp3586, Nil)


} else {
tmp3591 := GConsP(__e, V776)

if True == tmp3591 {
tmp3587 := GHd(__e, V776)

tmp3588 := GTl(__e, V776)

tmp3589 := Call2(__e, SymFn(symshen_4scan_1body), V775, tmp3588)


return GCons(__e, tmp3587, tmp3589)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.scan-body"))
}


}


}


}


panic("unreachable")
})

tmp3617 := Call2(__e, ns2_1set, symshen_4scan_1body, tmp3577)


_ = tmp3617

tmp3618 := MakeNative1(func(__e *ControlFlow, V783 Obj) Obj {
tmp3647 := GConsP(__e, V783)

var ifres3620 Obj

if True == tmp3647 {
tmp3645 := GTl(__e, V783)

tmp3646 := GConsP(__e, tmp3645)

var ifres3621 Obj

if True == tmp3646 {
tmp3642 := GTl(__e, V783)

tmp3643 := GHd(__e, tmp3642)

tmp3644 := GConsP(__e, tmp3643)

var ifres3622 Obj

if True == tmp3644 {
tmp3638 := GTl(__e, V783)

tmp3639 := GHd(__e, tmp3638)

tmp3640 := GHd(__e, tmp3639)

tmp3641 := GEq(__e, symshen_4choicepoint_b, tmp3640)

var ifres3623 Obj

if True == tmp3641 {
tmp3634 := GTl(__e, V783)

tmp3635 := GHd(__e, tmp3634)

tmp3636 := GTl(__e, tmp3635)

tmp3637 := GConsP(__e, tmp3636)

var ifres3624 Obj

if True == tmp3637 {
tmp3629 := GTl(__e, V783)

tmp3630 := GHd(__e, tmp3629)

tmp3631 := GTl(__e, tmp3630)

tmp3632 := GTl(__e, tmp3631)

tmp3633 := GEq(__e, Nil, tmp3632)

var ifres3625 Obj

if True == tmp3633 {
tmp3626 := GTl(__e, V783)

tmp3627 := GTl(__e, tmp3626)

tmp3628 := GEq(__e, Nil, tmp3627)

ifres3625 = tmp3628


} else {
ifres3625 = False


}

ifres3624 = ifres3625


} else {
ifres3624 = False


}

ifres3623 = ifres3624


} else {
ifres3623 = False


}

ifres3622 = ifres3623


} else {
ifres3622 = False


}

ifres3621 = ifres3622


} else {
ifres3621 = False


}

ifres3620 = ifres3621


} else {
ifres3620 = False


}

if True == ifres3620 {
return True
} else {
return False
}


panic("unreachable")
})

tmp3648 := Call2(__e, ns2_1set, symshen_4choicepoint_2, tmp3618)


_ = tmp3648

tmp3649 := MakeNative5(func(__e *ControlFlow, V799, V800, V801, V802, V803 Obj) Obj {
tmp3825 := GConsP(__e, V802)

var ifres3758 Obj

if True == tmp3825 {
tmp3823 := GTl(__e, V802)

tmp3824 := GConsP(__e, tmp3823)

var ifres3759 Obj

if True == tmp3824 {
tmp3820 := GTl(__e, V802)

tmp3821 := GHd(__e, tmp3820)

tmp3822 := GConsP(__e, tmp3821)

var ifres3760 Obj

if True == tmp3822 {
tmp3816 := GTl(__e, V802)

tmp3817 := GHd(__e, tmp3816)

tmp3818 := GTl(__e, tmp3817)

tmp3819 := GConsP(__e, tmp3818)

var ifres3761 Obj

if True == tmp3819 {
tmp3811 := GTl(__e, V802)

tmp3812 := GHd(__e, tmp3811)

tmp3813 := GTl(__e, tmp3812)

tmp3814 := GHd(__e, tmp3813)

tmp3815 := GConsP(__e, tmp3814)

var ifres3762 Obj

if True == tmp3815 {
tmp3805 := GTl(__e, V802)

tmp3806 := GHd(__e, tmp3805)

tmp3807 := GTl(__e, tmp3806)

tmp3808 := GHd(__e, tmp3807)

tmp3809 := GHd(__e, tmp3808)

tmp3810 := GEq(__e, symfail_1if, tmp3809)

var ifres3763 Obj

if True == tmp3810 {
tmp3799 := GTl(__e, V802)

tmp3800 := GHd(__e, tmp3799)

tmp3801 := GTl(__e, tmp3800)

tmp3802 := GHd(__e, tmp3801)

tmp3803 := GTl(__e, tmp3802)

tmp3804 := GConsP(__e, tmp3803)

var ifres3764 Obj

if True == tmp3804 {
tmp3792 := GTl(__e, V802)

tmp3793 := GHd(__e, tmp3792)

tmp3794 := GTl(__e, tmp3793)

tmp3795 := GHd(__e, tmp3794)

tmp3796 := GTl(__e, tmp3795)

tmp3797 := GTl(__e, tmp3796)

tmp3798 := GConsP(__e, tmp3797)

var ifres3765 Obj

if True == tmp3798 {
tmp3784 := GTl(__e, V802)

tmp3785 := GHd(__e, tmp3784)

tmp3786 := GTl(__e, tmp3785)

tmp3787 := GHd(__e, tmp3786)

tmp3788 := GTl(__e, tmp3787)

tmp3789 := GTl(__e, tmp3788)

tmp3790 := GTl(__e, tmp3789)

tmp3791 := GEq(__e, Nil, tmp3790)

var ifres3766 Obj

if True == tmp3791 {
tmp3779 := GTl(__e, V802)

tmp3780 := GHd(__e, tmp3779)

tmp3781 := GTl(__e, tmp3780)

tmp3782 := GTl(__e, tmp3781)

tmp3783 := GEq(__e, Nil, tmp3782)

var ifres3767 Obj

if True == tmp3783 {
tmp3776 := GTl(__e, V802)

tmp3777 := GTl(__e, tmp3776)

tmp3778 := GEq(__e, Nil, tmp3777)

var ifres3768 Obj

if True == tmp3778 {
tmp3769 := GTl(__e, V802)

tmp3770 := GHd(__e, tmp3769)

tmp3771 := GTl(__e, tmp3770)

tmp3772 := GHd(__e, tmp3771)

tmp3773 := GTl(__e, tmp3772)

tmp3774 := GHd(__e, tmp3773)

tmp3775 := GEq(__e, V799, tmp3774)

ifres3768 = tmp3775


} else {
ifres3768 = False


}

ifres3767 = ifres3768


} else {
ifres3767 = False


}

ifres3766 = ifres3767


} else {
ifres3766 = False


}

ifres3765 = ifres3766


} else {
ifres3765 = False


}

ifres3764 = ifres3765


} else {
ifres3764 = False


}

ifres3763 = ifres3764


} else {
ifres3763 = False


}

ifres3762 = ifres3763


} else {
ifres3762 = False


}

ifres3761 = ifres3762


} else {
ifres3761 = False


}

ifres3760 = ifres3761


} else {
ifres3760 = False


}

ifres3759 = ifres3760


} else {
ifres3759 = False


}

ifres3758 = ifres3759


} else {
ifres3758 = False


}

if True == ifres3758 {
tmp3650 := GTl(__e, V802)

tmp3651 := GHd(__e, tmp3650)

tmp3652 := GTl(__e, tmp3651)

tmp3653 := GHd(__e, tmp3652)

tmp3654 := GTl(__e, tmp3653)

tmp3655 := GHd(__e, tmp3654)

tmp3656 := Call2(__e, SymFn(symshen_4scan_1body), tmp3655, V803)


tmp3657 := GCons(__e, symcond, tmp3656)

tmp3658 := GCons(__e, tmp3657, Nil)

tmp3659 := GCons(__e, symfreeze, tmp3658)

tmp3660 := GHd(__e, V802)

tmp3661 := GTl(__e, V802)

tmp3662 := GHd(__e, tmp3661)

tmp3663 := GTl(__e, tmp3662)

tmp3664 := GHd(__e, tmp3663)

tmp3665 := GTl(__e, tmp3664)

tmp3666 := GTl(__e, tmp3665)

tmp3667 := GHd(__e, tmp3666)

tmp3668 := GTl(__e, V802)

tmp3669 := GHd(__e, tmp3668)

tmp3670 := GTl(__e, tmp3669)

tmp3671 := GHd(__e, tmp3670)

tmp3672 := GTl(__e, tmp3671)

tmp3673 := GHd(__e, tmp3672)

tmp3674 := GCons(__e, V801, Nil)

tmp3675 := GCons(__e, tmp3673, tmp3674)

tmp3676 := GCons(__e, V800, Nil)

tmp3677 := GCons(__e, symthaw, tmp3676)

tmp3678 := GCons(__e, V801, Nil)

tmp3679 := GCons(__e, tmp3677, tmp3678)

tmp3680 := GCons(__e, tmp3675, tmp3679)

tmp3681 := GCons(__e, symif, tmp3680)

tmp3682 := GCons(__e, tmp3681, Nil)

tmp3683 := GCons(__e, tmp3667, tmp3682)

tmp3684 := GCons(__e, V801, tmp3683)

tmp3685 := GCons(__e, symlet, tmp3684)

tmp3686 := GCons(__e, V800, Nil)

tmp3687 := GCons(__e, symthaw, tmp3686)

tmp3688 := GCons(__e, tmp3687, Nil)

tmp3689 := GCons(__e, tmp3685, tmp3688)

tmp3690 := GCons(__e, tmp3660, tmp3689)

tmp3691 := GCons(__e, symif, tmp3690)

tmp3692 := GCons(__e, tmp3691, Nil)

tmp3693 := GCons(__e, tmp3659, tmp3692)

tmp3694 := GCons(__e, V800, tmp3693)

tmp3695 := GCons(__e, symlet, tmp3694)

tmp3696 := GCons(__e, tmp3695, Nil)

tmp3697 := GCons(__e, True, tmp3696)

return GCons(__e, tmp3697, Nil)


} else {
tmp3756 := GConsP(__e, V802)

var ifres3734 Obj

if True == tmp3756 {
tmp3754 := GTl(__e, V802)

tmp3755 := GConsP(__e, tmp3754)

var ifres3735 Obj

if True == tmp3755 {
tmp3751 := GTl(__e, V802)

tmp3752 := GHd(__e, tmp3751)

tmp3753 := GConsP(__e, tmp3752)

var ifres3736 Obj

if True == tmp3753 {
tmp3747 := GTl(__e, V802)

tmp3748 := GHd(__e, tmp3747)

tmp3749 := GTl(__e, tmp3748)

tmp3750 := GConsP(__e, tmp3749)

var ifres3737 Obj

if True == tmp3750 {
tmp3742 := GTl(__e, V802)

tmp3743 := GHd(__e, tmp3742)

tmp3744 := GTl(__e, tmp3743)

tmp3745 := GTl(__e, tmp3744)

tmp3746 := GEq(__e, Nil, tmp3745)

var ifres3738 Obj

if True == tmp3746 {
tmp3739 := GTl(__e, V802)

tmp3740 := GTl(__e, tmp3739)

tmp3741 := GEq(__e, Nil, tmp3740)

ifres3738 = tmp3741


} else {
ifres3738 = False


}

ifres3737 = ifres3738


} else {
ifres3737 = False


}

ifres3736 = ifres3737


} else {
ifres3736 = False


}

ifres3735 = ifres3736


} else {
ifres3735 = False


}

ifres3734 = ifres3735


} else {
ifres3734 = False


}

if True == ifres3734 {
tmp3698 := Call2(__e, SymFn(symshen_4scan_1body), V799, V803)


tmp3699 := GCons(__e, symcond, tmp3698)

tmp3700 := GCons(__e, tmp3699, Nil)

tmp3701 := GCons(__e, symfreeze, tmp3700)

tmp3702 := GHd(__e, V802)

tmp3703 := GTl(__e, V802)

tmp3704 := GHd(__e, tmp3703)

tmp3705 := GTl(__e, tmp3704)

tmp3706 := GHd(__e, tmp3705)

tmp3707 := GCons(__e, symfail, Nil)

tmp3708 := GCons(__e, tmp3707, Nil)

tmp3709 := GCons(__e, V801, tmp3708)

tmp3710 := GCons(__e, sym_a, tmp3709)

tmp3711 := GCons(__e, V800, Nil)

tmp3712 := GCons(__e, symthaw, tmp3711)

tmp3713 := GCons(__e, V801, Nil)

tmp3714 := GCons(__e, tmp3712, tmp3713)

tmp3715 := GCons(__e, tmp3710, tmp3714)

tmp3716 := GCons(__e, symif, tmp3715)

tmp3717 := GCons(__e, tmp3716, Nil)

tmp3718 := GCons(__e, tmp3706, tmp3717)

tmp3719 := GCons(__e, V801, tmp3718)

tmp3720 := GCons(__e, symlet, tmp3719)

tmp3721 := GCons(__e, V800, Nil)

tmp3722 := GCons(__e, symthaw, tmp3721)

tmp3723 := GCons(__e, tmp3722, Nil)

tmp3724 := GCons(__e, tmp3720, tmp3723)

tmp3725 := GCons(__e, tmp3702, tmp3724)

tmp3726 := GCons(__e, symif, tmp3725)

tmp3727 := GCons(__e, tmp3726, Nil)

tmp3728 := GCons(__e, tmp3701, tmp3727)

tmp3729 := GCons(__e, V800, tmp3728)

tmp3730 := GCons(__e, symlet, tmp3729)

tmp3731 := GCons(__e, tmp3730, Nil)

tmp3732 := GCons(__e, True, tmp3731)

return GCons(__e, tmp3732, Nil)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.choicepoint"))
}


}


panic("unreachable")
})

tmp3826 := Call2(__e, ns2_1set, symshen_4choicepoint, tmp3649)


_ = tmp3826

tmp3827 := MakeNative3(func(__e *ControlFlow, V805, V806, V807 Obj) Obj {
tmp3840 := GEq(__e, V805, V807)

if True == tmp3840 {
return V806
} else {
tmp3838 := GConsP(__e, V807)

if True == tmp3838 {
tmp3828 := GHd(__e, V807)

tmp3829 := Call3(__e, SymFn(symshen_4rep_1X), V805, V806, tmp3828)


let__2646 := tmp3829
_ = let__2646

tmp3835 := GHd(__e, V807)

tmp3836 := GEq(__e, let__2646, tmp3835)

if True == tmp3836 {
tmp3830 := GHd(__e, V807)

tmp3831 := GTl(__e, V807)

tmp3832 := Call3(__e, SymFn(symshen_4rep_1X), V805, V806, tmp3831)


return GCons(__e, tmp3830, tmp3832)


} else {
tmp3833 := GTl(__e, V807)

return GCons(__e, let__2646, tmp3833)


}


} else {
return V807
}


}


panic("unreachable")
})

tmp3841 := Call2(__e, ns2_1set, symshen_4rep_1X, tmp3827)


_ = tmp3841

tmp3842 := MakeNative1(func(__e *ControlFlow, V809 Obj) Obj {
tmp3912 := GConsP(__e, V809)

var ifres3897 Obj

if True == tmp3912 {
tmp3910 := GHd(__e, V809)

tmp3911 := GEq(__e, symlambda, tmp3910)

var ifres3898 Obj

if True == tmp3911 {
tmp3908 := GTl(__e, V809)

tmp3909 := GConsP(__e, tmp3908)

var ifres3899 Obj

if True == tmp3909 {
tmp3905 := GTl(__e, V809)

tmp3906 := GTl(__e, tmp3905)

tmp3907 := GConsP(__e, tmp3906)

var ifres3900 Obj

if True == tmp3907 {
tmp3901 := GTl(__e, V809)

tmp3902 := GTl(__e, tmp3901)

tmp3903 := GTl(__e, tmp3902)

tmp3904 := GEq(__e, Nil, tmp3903)

ifres3900 = tmp3904


} else {
ifres3900 = False


}

ifres3899 = ifres3900


} else {
ifres3899 = False


}

ifres3898 = ifres3899


} else {
ifres3898 = False


}

ifres3897 = ifres3898


} else {
ifres3897 = False


}

if True == ifres3897 {
tmp3843 := Call1(__e, SymFn(symgensym), symZ)


let__2647 := tmp3843
_ = let__2647

tmp3844 := GTl(__e, V809)

tmp3845 := GHd(__e, tmp3844)

tmp3846 := GTl(__e, V809)

tmp3847 := GTl(__e, tmp3846)

tmp3848 := GHd(__e, tmp3847)

tmp3849 := Call3(__e, SymFn(symshen_4beta), tmp3845, let__2647, tmp3848)


tmp3850 := GCons(__e, tmp3849, Nil)

tmp3851 := GCons(__e, let__2647, tmp3850)

tmp3852 := GCons(__e, symlambda, tmp3851)

let__2648 := tmp3852
_ = let__2648

tmp3853 := MakeNative1(func(__e *ControlFlow, Z812 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4alpha_1convert), Z812)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp3853, let__2648)


} else {
tmp3895 := GConsP(__e, V809)

var ifres3874 Obj

if True == tmp3895 {
tmp3893 := GHd(__e, V809)

tmp3894 := GEq(__e, symlet, tmp3893)

var ifres3875 Obj

if True == tmp3894 {
tmp3891 := GTl(__e, V809)

tmp3892 := GConsP(__e, tmp3891)

var ifres3876 Obj

if True == tmp3892 {
tmp3888 := GTl(__e, V809)

tmp3889 := GTl(__e, tmp3888)

tmp3890 := GConsP(__e, tmp3889)

var ifres3877 Obj

if True == tmp3890 {
tmp3884 := GTl(__e, V809)

tmp3885 := GTl(__e, tmp3884)

tmp3886 := GTl(__e, tmp3885)

tmp3887 := GConsP(__e, tmp3886)

var ifres3878 Obj

if True == tmp3887 {
tmp3879 := GTl(__e, V809)

tmp3880 := GTl(__e, tmp3879)

tmp3881 := GTl(__e, tmp3880)

tmp3882 := GTl(__e, tmp3881)

tmp3883 := GEq(__e, Nil, tmp3882)

ifres3878 = tmp3883


} else {
ifres3878 = False


}

ifres3877 = ifres3878


} else {
ifres3877 = False


}

ifres3876 = ifres3877


} else {
ifres3876 = False


}

ifres3875 = ifres3876


} else {
ifres3875 = False


}

ifres3874 = ifres3875


} else {
ifres3874 = False


}

if True == ifres3874 {
tmp3854 := Call1(__e, SymFn(symgensym), symW)


let__2649 := tmp3854
_ = let__2649

tmp3855 := GTl(__e, V809)

tmp3856 := GTl(__e, tmp3855)

tmp3857 := GHd(__e, tmp3856)

tmp3858 := GTl(__e, V809)

tmp3859 := GHd(__e, tmp3858)

tmp3860 := GTl(__e, V809)

tmp3861 := GTl(__e, tmp3860)

tmp3862 := GTl(__e, tmp3861)

tmp3863 := GHd(__e, tmp3862)

tmp3864 := Call3(__e, SymFn(symshen_4beta), tmp3859, let__2649, tmp3863)


tmp3865 := GCons(__e, tmp3864, Nil)

tmp3866 := GCons(__e, tmp3857, tmp3865)

tmp3867 := GCons(__e, let__2649, tmp3866)

tmp3868 := GCons(__e, symlet, tmp3867)

let__2650 := tmp3868
_ = let__2650

tmp3869 := MakeNative1(func(__e *ControlFlow, Z815 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4alpha_1convert), Z815)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp3869, let__2650)


} else {
tmp3872 := GConsP(__e, V809)

if True == tmp3872 {
tmp3870 := MakeNative1(func(__e *ControlFlow, Z816 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4alpha_1convert), Z816)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp3870, V809)


} else {
return V809
}


}


}


panic("unreachable")
})

tmp3913 := Call2(__e, ns2_1set, symshen_4alpha_1convert, tmp3842)


_ = tmp3913

tmp3914 := MakeNative2(func(__e *ControlFlow, V817, V818 Obj) Obj {
tmp3915 := MakeNative1(func(__e *ControlFlow, Z819 Obj) Obj {
tmp3916 := Call1(__e, SymFn(symfst), Z819)


tmp3917 := Call1(__e, SymFn(symsnd), Z819)


tmp3918 := Call1(__e, SymFn(symshen_4alpha_1convert), tmp3917)


return TailCall4(__e, SymFn(symshen_4triple_1stack), Nil, tmp3916, V818, tmp3918)


panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp3915, V817)


panic("unreachable")
})

tmp3919 := Call2(__e, ns2_1set, symshen_4kl_1body, tmp3914)


_ = tmp3919

tmp3920 := (func() Obj {
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
tmp4051 := GEq(__e, Nil, V829)

var ifres4032 Obj

if True == tmp4051 {
tmp4050 := GEq(__e, Nil, V830)

var ifres4033 Obj

if True == tmp4050 {
tmp4049 := GConsP(__e, V831)

var ifres4034 Obj

if True == tmp4049 {
tmp4047 := GHd(__e, V831)

tmp4048 := GEq(__e, symwhere, tmp4047)

var ifres4035 Obj

if True == tmp4048 {
tmp4045 := GTl(__e, V831)

tmp4046 := GConsP(__e, tmp4045)

var ifres4036 Obj

if True == tmp4046 {
tmp4042 := GTl(__e, V831)

tmp4043 := GTl(__e, tmp4042)

tmp4044 := GConsP(__e, tmp4043)

var ifres4037 Obj

if True == tmp4044 {
tmp4038 := GTl(__e, V831)

tmp4039 := GTl(__e, tmp4038)

tmp4040 := GTl(__e, tmp4039)

tmp4041 := GEq(__e, Nil, tmp4040)

ifres4037 = tmp4041


} else {
ifres4037 = False


}

ifres4036 = ifres4037


} else {
ifres4036 = False


}

ifres4035 = ifres4036


} else {
ifres4035 = False


}

ifres4034 = ifres4035


} else {
ifres4034 = False


}

ifres4033 = ifres4034


} else {
ifres4033 = False


}

ifres4032 = ifres4033


} else {
ifres4032 = False


}

if True == ifres4032 {
tmp3921 := GTl(__e, V831)

tmp3922 := GHd(__e, tmp3921)

tmp3923 := GCons(__e, tmp3922, V828)

tmp3924 := GTl(__e, V831)

tmp3925 := GTl(__e, tmp3924)

tmp3926 := GHd(__e, tmp3925)

if SymFn(symshen_4triple_1stack) == __self {
__self1, __self2, __self3, __self4 = tmp3923, Nil, Nil, tmp3926
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4triple_1stack), tmp3923, Nil, Nil, tmp3926)


} else {
tmp4030 := GEq(__e, Nil, V829)

var ifres4028 Obj

if True == tmp4030 {
tmp4029 := GEq(__e, Nil, V830)

ifres4028 = tmp4029


} else {
ifres4028 = False


}

if True == ifres4028 {
tmp3927 := Call1(__e, SymFn(symreverse), V828)


tmp3928 := Call1(__e, SymFn(symshen_4rectify_1test), tmp3927)


tmp3929 := GCons(__e, V831, Nil)

return GCons(__e, tmp3928, tmp3929)


} else {
tmp4026 := GConsP(__e, V829)

var ifres4021 Obj

if True == tmp4026 {
tmp4025 := GConsP(__e, V830)

var ifres4022 Obj

if True == tmp4025 {
tmp4023 := GHd(__e, V829)

tmp4024 := PrimGuard1(__e, symvariable_2, PrimIsVariable, tmp4023)

ifres4022 = tmp4024


} else {
ifres4022 = False


}

ifres4021 = ifres4022


} else {
ifres4021 = False


}

if True == ifres4021 {
tmp3930 := GTl(__e, V829)

tmp3931 := GTl(__e, V830)

tmp3932 := GHd(__e, V829)

tmp3933 := GHd(__e, V830)

tmp3934 := Call3(__e, SymFn(symshen_4beta), tmp3932, tmp3933, V831)


if SymFn(symshen_4triple_1stack) == __self {
__self1, __self2, __self3, __self4 = V828, tmp3930, tmp3931, tmp3934
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4triple_1stack), V828, tmp3930, tmp3931, tmp3934)


} else {
tmp4019 := GConsP(__e, V829)

var ifres4014 Obj

if True == tmp4019 {
tmp4018 := GConsP(__e, V830)

var ifres4015 Obj

if True == tmp4018 {
tmp4016 := GHd(__e, V829)

tmp4017 := Call1(__e, SymFn(symshen_4custom_1pattern_2), tmp4016)


ifres4015 = tmp4017


} else {
ifres4015 = False


}

ifres4014 = ifres4015


} else {
ifres4014 = False


}

if True == ifres4014 {
tmp3935 := GHd(__e, V829)

tmp3936 := Call1(__e, SymFn(symshen_4custom_1pattern_1body), tmp3935)


tmp3937 := GTl(__e, V829)

tmp3938 := GHd(__e, V830)

tmp3939 := GTl(__e, V830)

return TailCall6(__e, SymFn(symshen_4custom_1pattern_1triple_1stack), V828, tmp3936, tmp3937, tmp3938, tmp3939, V831)


} else {
tmp4012 := GConsP(__e, V829)

var ifres3988 Obj

if True == tmp4012 {
tmp4010 := GHd(__e, V829)

tmp4011 := GConsP(__e, tmp4010)

var ifres3989 Obj

if True == tmp4011 {
tmp4007 := GHd(__e, V829)

tmp4008 := GTl(__e, tmp4007)

tmp4009 := GConsP(__e, tmp4008)

var ifres3990 Obj

if True == tmp4009 {
tmp4003 := GHd(__e, V829)

tmp4004 := GTl(__e, tmp4003)

tmp4005 := GTl(__e, tmp4004)

tmp4006 := GConsP(__e, tmp4005)

var ifres3991 Obj

if True == tmp4006 {
tmp3998 := GHd(__e, V829)

tmp3999 := GTl(__e, tmp3998)

tmp4000 := GTl(__e, tmp3999)

tmp4001 := GTl(__e, tmp4000)

tmp4002 := GEq(__e, Nil, tmp4001)

var ifres3992 Obj

if True == tmp4002 {
tmp3997 := GConsP(__e, V830)

var ifres3993 Obj

if True == tmp3997 {
tmp3994 := GHd(__e, V829)

tmp3995 := GHd(__e, tmp3994)

tmp3996 := Call1(__e, SymFn(symshen_4constructor_2), tmp3995)


ifres3993 = tmp3996


} else {
ifres3993 = False


}

ifres3992 = ifres3993


} else {
ifres3992 = False


}

ifres3991 = ifres3992


} else {
ifres3991 = False


}

ifres3990 = ifres3991


} else {
ifres3990 = False


}

ifres3989 = ifres3990


} else {
ifres3989 = False


}

ifres3988 = ifres3989


} else {
ifres3988 = False


}

if True == ifres3988 {
tmp3940 := GHd(__e, V829)

tmp3941 := GHd(__e, tmp3940)

tmp3942 := Call1(__e, SymFn(symshen_4op_1test), tmp3941)


tmp3943 := GHd(__e, V830)

tmp3944 := GCons(__e, tmp3943, Nil)

tmp3945 := GCons(__e, tmp3942, tmp3944)

tmp3946 := GCons(__e, tmp3945, V828)

tmp3947 := GHd(__e, V829)

tmp3948 := GTl(__e, tmp3947)

tmp3949 := GHd(__e, tmp3948)

tmp3950 := GHd(__e, V829)

tmp3951 := GTl(__e, tmp3950)

tmp3952 := GTl(__e, tmp3951)

tmp3953 := GHd(__e, tmp3952)

tmp3954 := GTl(__e, V829)

tmp3955 := GCons(__e, tmp3953, tmp3954)

tmp3956 := GCons(__e, tmp3949, tmp3955)

tmp3957 := GHd(__e, V829)

tmp3958 := GHd(__e, tmp3957)

tmp3959 := Call1(__e, SymFn(symshen_4op1), tmp3958)


tmp3960 := GHd(__e, V830)

tmp3961 := GCons(__e, tmp3960, Nil)

tmp3962 := GCons(__e, tmp3959, tmp3961)

tmp3963 := GHd(__e, V829)

tmp3964 := GHd(__e, tmp3963)

tmp3965 := Call1(__e, SymFn(symshen_4op2), tmp3964)


tmp3966 := GHd(__e, V830)

tmp3967 := GCons(__e, tmp3966, Nil)

tmp3968 := GCons(__e, tmp3965, tmp3967)

tmp3969 := GTl(__e, V830)

tmp3970 := GCons(__e, tmp3968, tmp3969)

tmp3971 := GCons(__e, tmp3962, tmp3970)

tmp3972 := GHd(__e, V829)

tmp3973 := GHd(__e, V830)

tmp3974 := Call3(__e, SymFn(symshen_4beta), tmp3972, tmp3973, V831)


if SymFn(symshen_4triple_1stack) == __self {
__self1, __self2, __self3, __self4 = tmp3946, tmp3956, tmp3971, tmp3974
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4triple_1stack), tmp3946, tmp3956, tmp3971, tmp3974)


} else {
tmp3986 := GConsP(__e, V829)

var ifres3984 Obj

if True == tmp3986 {
tmp3985 := GConsP(__e, V830)

ifres3984 = tmp3985


} else {
ifres3984 = False


}

if True == ifres3984 {
tmp3975 := GHd(__e, V829)

tmp3976 := GHd(__e, V830)

tmp3977 := GCons(__e, tmp3976, Nil)

tmp3978 := GCons(__e, tmp3975, tmp3977)

tmp3979 := GCons(__e, sym_a, tmp3978)

tmp3980 := GCons(__e, tmp3979, V828)

tmp3981 := GTl(__e, V829)

tmp3982 := GTl(__e, V830)

if SymFn(symshen_4triple_1stack) == __self {
__self1, __self2, __self3, __self4 = tmp3980, tmp3981, tmp3982, V831
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4triple_1stack), tmp3980, tmp3981, tmp3982, V831)


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

tmp4052 := Call2(__e, ns2_1set, symshen_4triple_1stack, tmp3920)


_ = tmp4052

tmp4053 := MakeNative6(func(__e *ControlFlow, V832, V833, V834, V835, V836, V837 Obj) Obj {
tmp4054 := Call2(__e, SymFn(sym_8p), V833, V835)


tmp4055 := Call1(__e, SymFn(symshen_4custom_1pattern_1reducer), tmp4054)


let__2651 := tmp4055
_ = let__2651

tmp4068 := Call1(__e, SymFn(symtuple_2), let__2651)


if True == tmp4068 {
tmp4056 := Call1(__e, SymFn(symfst), let__2651)


let__2652 := tmp4056
_ = let__2652

tmp4057 := Call1(__e, SymFn(symsnd), let__2651)


let__2653 := tmp4057
_ = let__2653

tmp4058 := Call1(__e, SymFn(symreverse), let__2652)


tmp4059 := Call2(__e, SymFn(symappend), tmp4058, V832)


tmp4060 := MakeNative1(func(__e *ControlFlow, Z841 Obj) Obj {
return TailCall1(__e, SymFn(symfst), Z841)
panic("unreachable")
})

tmp4061 := Call2(__e, SymFn(symmap), tmp4060, let__2653)


tmp4062 := Call2(__e, SymFn(symappend), tmp4061, V834)


tmp4063 := MakeNative1(func(__e *ControlFlow, Z842 Obj) Obj {
return TailCall1(__e, SymFn(symsnd), Z842)
panic("unreachable")
})

tmp4064 := Call2(__e, SymFn(symmap), tmp4063, let__2653)


tmp4065 := Call2(__e, SymFn(symappend), tmp4064, V836)


tmp4066 := Call3(__e, SymFn(symshen_4beta), V833, V835, V837)


return TailCall4(__e, SymFn(symshen_4triple_1stack), tmp4059, tmp4062, tmp4065, tmp4066)


} else {
return TailCall1(__e, SymFn(symshen_4constructor_1error), V833)
}


panic("unreachable")
})

tmp4069 := Call2(__e, ns2_1set, symshen_4custom_1pattern_1triple_1stack, tmp4053)


_ = tmp4069

tmp4070 := MakeNative1(func(__e *ControlFlow, V845 Obj) Obj {
tmp4087 := GEq(__e, Nil, V845)

if True == tmp4087 {
return True
} else {
tmp4085 := GConsP(__e, V845)

var ifres4082 Obj

if True == tmp4085 {
tmp4083 := GTl(__e, V845)

tmp4084 := GEq(__e, Nil, tmp4083)

ifres4082 = tmp4084


} else {
ifres4082 = False


}

if True == ifres4082 {
return GHd(__e, V845)
} else {
tmp4080 := GConsP(__e, V845)

var ifres4077 Obj

if True == tmp4080 {
tmp4078 := GTl(__e, V845)

tmp4079 := GConsP(__e, tmp4078)

ifres4077 = tmp4079


} else {
ifres4077 = False


}

if True == ifres4077 {
tmp4071 := GHd(__e, V845)

tmp4072 := GTl(__e, V845)

tmp4073 := Call1(__e, SymFn(symshen_4rectify_1test), tmp4072)


tmp4074 := GCons(__e, tmp4073, Nil)

tmp4075 := GCons(__e, tmp4071, tmp4074)

return GCons(__e, symand, tmp4075)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.rectify-test"))
}


}


}


panic("unreachable")
})

tmp4088 := Call2(__e, ns2_1set, symshen_4rectify_1test, tmp4070)


_ = tmp4088

tmp4089 := MakeNative3(func(__e *ControlFlow, V855, V856, V857 Obj) Obj {
tmp4155 := GEq(__e, V855, V857)

if True == tmp4155 {
return V856
} else {
tmp4153 := GConsP(__e, V857)

var ifres4134 Obj

if True == tmp4153 {
tmp4151 := GHd(__e, V857)

tmp4152 := GEq(__e, symlambda, tmp4151)

var ifres4135 Obj

if True == tmp4152 {
tmp4149 := GTl(__e, V857)

tmp4150 := GConsP(__e, tmp4149)

var ifres4136 Obj

if True == tmp4150 {
tmp4146 := GTl(__e, V857)

tmp4147 := GTl(__e, tmp4146)

tmp4148 := GConsP(__e, tmp4147)

var ifres4137 Obj

if True == tmp4148 {
tmp4142 := GTl(__e, V857)

tmp4143 := GTl(__e, tmp4142)

tmp4144 := GTl(__e, tmp4143)

tmp4145 := GEq(__e, Nil, tmp4144)

var ifres4138 Obj

if True == tmp4145 {
tmp4139 := GTl(__e, V857)

tmp4140 := GHd(__e, tmp4139)

tmp4141 := GEq(__e, V855, tmp4140)

ifres4138 = tmp4141


} else {
ifres4138 = False


}

ifres4137 = ifres4138


} else {
ifres4137 = False


}

ifres4136 = ifres4137


} else {
ifres4136 = False


}

ifres4135 = ifres4136


} else {
ifres4135 = False


}

ifres4134 = ifres4135


} else {
ifres4134 = False


}

if True == ifres4134 {
return V857
} else {
tmp4132 := GConsP(__e, V857)

var ifres4107 Obj

if True == tmp4132 {
tmp4130 := GHd(__e, V857)

tmp4131 := GEq(__e, symlet, tmp4130)

var ifres4108 Obj

if True == tmp4131 {
tmp4128 := GTl(__e, V857)

tmp4129 := GConsP(__e, tmp4128)

var ifres4109 Obj

if True == tmp4129 {
tmp4125 := GTl(__e, V857)

tmp4126 := GTl(__e, tmp4125)

tmp4127 := GConsP(__e, tmp4126)

var ifres4110 Obj

if True == tmp4127 {
tmp4121 := GTl(__e, V857)

tmp4122 := GTl(__e, tmp4121)

tmp4123 := GTl(__e, tmp4122)

tmp4124 := GConsP(__e, tmp4123)

var ifres4111 Obj

if True == tmp4124 {
tmp4116 := GTl(__e, V857)

tmp4117 := GTl(__e, tmp4116)

tmp4118 := GTl(__e, tmp4117)

tmp4119 := GTl(__e, tmp4118)

tmp4120 := GEq(__e, Nil, tmp4119)

var ifres4112 Obj

if True == tmp4120 {
tmp4113 := GTl(__e, V857)

tmp4114 := GHd(__e, tmp4113)

tmp4115 := GEq(__e, V855, tmp4114)

ifres4112 = tmp4115


} else {
ifres4112 = False


}

ifres4111 = ifres4112


} else {
ifres4111 = False


}

ifres4110 = ifres4111


} else {
ifres4110 = False


}

ifres4109 = ifres4110


} else {
ifres4109 = False


}

ifres4108 = ifres4109


} else {
ifres4108 = False


}

ifres4107 = ifres4108


} else {
ifres4107 = False


}

if True == ifres4107 {
tmp4090 := GTl(__e, V857)

tmp4091 := GHd(__e, tmp4090)

tmp4092 := GTl(__e, V857)

tmp4093 := GHd(__e, tmp4092)

tmp4094 := GTl(__e, V857)

tmp4095 := GTl(__e, tmp4094)

tmp4096 := GHd(__e, tmp4095)

tmp4097 := Call3(__e, SymFn(symshen_4beta), tmp4093, V856, tmp4096)


tmp4098 := GTl(__e, V857)

tmp4099 := GTl(__e, tmp4098)

tmp4100 := GTl(__e, tmp4099)

tmp4101 := GCons(__e, tmp4097, tmp4100)

tmp4102 := GCons(__e, tmp4091, tmp4101)

return GCons(__e, symlet, tmp4102)


} else {
tmp4105 := GConsP(__e, V857)

if True == tmp4105 {
tmp4103 := MakeNative1(func(__e *ControlFlow, Z858 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4beta), V855, V856, Z858)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp4103, V857)


} else {
return V857
}


}


}


}


panic("unreachable")
})

tmp4156 := Call2(__e, ns2_1set, symshen_4beta, tmp4089)


_ = tmp4156

tmp4157 := MakeNative1(func(__e *ControlFlow, V861 Obj) Obj {
tmp4165 := GEq(__e, symcons, V861)

if True == tmp4165 {
return symhd
} else {
tmp4163 := GEq(__e, sym_8s, V861)

if True == tmp4163 {
return symhdstr
} else {
tmp4161 := GEq(__e, sym_8p, V861)

if True == tmp4161 {
return symfst
} else {
tmp4159 := GEq(__e, sym_8v, V861)

if True == tmp4159 {
return symhdv
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.op1"))
}


}


}


}


panic("unreachable")
})

tmp4166 := Call2(__e, ns2_1set, symshen_4op1, tmp4157)


_ = tmp4166

tmp4167 := MakeNative1(func(__e *ControlFlow, V864 Obj) Obj {
tmp4175 := GEq(__e, symcons, V864)

if True == tmp4175 {
return symtl
} else {
tmp4173 := GEq(__e, sym_8s, V864)

if True == tmp4173 {
return symtlstr
} else {
tmp4171 := GEq(__e, sym_8p, V864)

if True == tmp4171 {
return symsnd
} else {
tmp4169 := GEq(__e, sym_8v, V864)

if True == tmp4169 {
return symtlv
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.op2"))
}


}


}


}


panic("unreachable")
})

tmp4176 := Call2(__e, ns2_1set, symshen_4op2, tmp4167)


_ = tmp4176

tmp4177 := MakeNative1(func(__e *ControlFlow, V867 Obj) Obj {
tmp4185 := GEq(__e, symcons, V867)

if True == tmp4185 {
return symcons_2
} else {
tmp4183 := GEq(__e, sym_8s, V867)

if True == tmp4183 {
return symshen_4_7string_2
} else {
tmp4181 := GEq(__e, sym_8p, V867)

if True == tmp4181 {
return symtuple_2
} else {
tmp4179 := GEq(__e, sym_8v, V867)

if True == tmp4179 {
return symshen_4_7vector_2
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.op-test"))
}


}


}


}


panic("unreachable")
})

tmp4186 := Call2(__e, ns2_1set, symshen_4op_1test, tmp4177)


_ = tmp4186

tmp4187 := MakeNative1(func(__e *ControlFlow, V868 Obj) Obj {
tmp4189 := GEq(__e, MakeString(""), V868)

if True == tmp4189 {
return False
} else {
return PrimGuard1(__e, symstring_2, PrimIsString, V868)
}


panic("unreachable")
})

tmp4190 := Call2(__e, ns2_1set, symshen_4_7string_2, tmp4187)


_ = tmp4190

tmp4191 := MakeNative1(func(__e *ControlFlow, V869 Obj) Obj {
tmp4193 := Call1(__e, SymFn(symvector), MakeInteger(0))


tmp4194 := GEq(__e, V869, tmp4193)

if True == tmp4194 {
return False
} else {
return TailCall1(__e, SymFn(symvector_2), V869)
}


panic("unreachable")
})

tmp4195 := Call2(__e, ns2_1set, symshen_4_7vector_2, tmp4191)


_ = tmp4195

tmp4196 := MakeNative1(func(__e *ControlFlow, V872 Obj) Obj {
tmp4200 := GEq(__e, sym_7, V872)

if True == tmp4200 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dfactorise_2_d, True)
} else {
tmp4198 := GEq(__e, sym_1, V872)

if True == tmp4198 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dfactorise_2_d, False)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("factorise expects a + or a -\n"))
}


}


panic("unreachable")
})

tmp4201 := Call2(__e, ns2_1set, symfactorise, tmp4196)


_ = tmp4201

tmp4202 := MakeNative1(func(__e *ControlFlow, V873 Obj) Obj {
tmp4204 := GValue(__e, symshen_4_dfactorise_2_d)

if True == tmp4204 {
return TailCall1(__e, SymFn(symshen_4factor), V873)
} else {
return V873
}


panic("unreachable")
})

tmp4205 := Call2(__e, ns2_1set, symshen_4factorise_1code, tmp4202)


_ = tmp4205

tmp4206 := MakeNative1(func(__e *ControlFlow, V874 Obj) Obj {
tmp4256 := GConsP(__e, V874)

var ifres4222 Obj

if True == tmp4256 {
tmp4254 := GHd(__e, V874)

tmp4255 := GEq(__e, symdefun, tmp4254)

var ifres4223 Obj

if True == tmp4255 {
tmp4252 := GTl(__e, V874)

tmp4253 := GConsP(__e, tmp4252)

var ifres4224 Obj

if True == tmp4253 {
tmp4249 := GTl(__e, V874)

tmp4250 := GTl(__e, tmp4249)

tmp4251 := GConsP(__e, tmp4250)

var ifres4225 Obj

if True == tmp4251 {
tmp4245 := GTl(__e, V874)

tmp4246 := GTl(__e, tmp4245)

tmp4247 := GTl(__e, tmp4246)

tmp4248 := GConsP(__e, tmp4247)

var ifres4226 Obj

if True == tmp4248 {
tmp4240 := GTl(__e, V874)

tmp4241 := GTl(__e, tmp4240)

tmp4242 := GTl(__e, tmp4241)

tmp4243 := GHd(__e, tmp4242)

tmp4244 := GConsP(__e, tmp4243)

var ifres4227 Obj

if True == tmp4244 {
tmp4234 := GTl(__e, V874)

tmp4235 := GTl(__e, tmp4234)

tmp4236 := GTl(__e, tmp4235)

tmp4237 := GHd(__e, tmp4236)

tmp4238 := GHd(__e, tmp4237)

tmp4239 := GEq(__e, symcond, tmp4238)

var ifres4228 Obj

if True == tmp4239 {
tmp4229 := GTl(__e, V874)

tmp4230 := GTl(__e, tmp4229)

tmp4231 := GTl(__e, tmp4230)

tmp4232 := GTl(__e, tmp4231)

tmp4233 := GEq(__e, Nil, tmp4232)

ifres4228 = tmp4233


} else {
ifres4228 = False


}

ifres4227 = ifres4228


} else {
ifres4227 = False


}

ifres4226 = ifres4227


} else {
ifres4226 = False


}

ifres4225 = ifres4226


} else {
ifres4225 = False


}

ifres4224 = ifres4225


} else {
ifres4224 = False


}

ifres4223 = ifres4224


} else {
ifres4223 = False


}

ifres4222 = ifres4223


} else {
ifres4222 = False


}

if True == ifres4222 {
tmp4207 := GTl(__e, V874)

tmp4208 := GHd(__e, tmp4207)

tmp4209 := GTl(__e, V874)

tmp4210 := GTl(__e, tmp4209)

tmp4211 := GHd(__e, tmp4210)

tmp4212 := GTl(__e, V874)

tmp4213 := GTl(__e, tmp4212)

tmp4214 := GTl(__e, tmp4213)

tmp4215 := GHd(__e, tmp4214)

tmp4216 := GTl(__e, tmp4215)

tmp4217 := Call1(__e, SymFn(symshen_4factor_1recognisors), tmp4216)


tmp4218 := GCons(__e, tmp4217, Nil)

tmp4219 := GCons(__e, tmp4211, tmp4218)

tmp4220 := GCons(__e, tmp4208, tmp4219)

return GCons(__e, symdefun, tmp4220)


} else {
return V874
}


panic("unreachable")
})

tmp4257 := Call2(__e, ns2_1set, symshen_4factor, tmp4206)


_ = tmp4257

tmp4258 := MakeNative1(func(__e *ControlFlow, V877 Obj) Obj {
tmp4396 := GConsP(__e, V877)

var ifres4380 Obj

if True == tmp4396 {
tmp4394 := GHd(__e, V877)

tmp4395 := GConsP(__e, tmp4394)

var ifres4381 Obj

if True == tmp4395 {
tmp4391 := GHd(__e, V877)

tmp4392 := GHd(__e, tmp4391)

tmp4393 := GEq(__e, True, tmp4392)

var ifres4382 Obj

if True == tmp4393 {
tmp4388 := GHd(__e, V877)

tmp4389 := GTl(__e, tmp4388)

tmp4390 := GConsP(__e, tmp4389)

var ifres4383 Obj

if True == tmp4390 {
tmp4384 := GHd(__e, V877)

tmp4385 := GTl(__e, tmp4384)

tmp4386 := GTl(__e, tmp4385)

tmp4387 := GEq(__e, Nil, tmp4386)

ifres4383 = tmp4387


} else {
ifres4383 = False


}

ifres4382 = ifres4383


} else {
ifres4382 = False


}

ifres4381 = ifres4382


} else {
ifres4381 = False


}

ifres4380 = ifres4381


} else {
ifres4380 = False


}

if True == ifres4380 {
tmp4259 := GHd(__e, V877)

tmp4260 := GTl(__e, tmp4259)

return GHd(__e, tmp4260)


} else {
tmp4378 := GConsP(__e, V877)

var ifres4339 Obj

if True == tmp4378 {
tmp4376 := GHd(__e, V877)

tmp4377 := GConsP(__e, tmp4376)

var ifres4340 Obj

if True == tmp4377 {
tmp4373 := GHd(__e, V877)

tmp4374 := GHd(__e, tmp4373)

tmp4375 := GConsP(__e, tmp4374)

var ifres4341 Obj

if True == tmp4375 {
tmp4369 := GHd(__e, V877)

tmp4370 := GHd(__e, tmp4369)

tmp4371 := GHd(__e, tmp4370)

tmp4372 := GEq(__e, symand, tmp4371)

var ifres4342 Obj

if True == tmp4372 {
tmp4365 := GHd(__e, V877)

tmp4366 := GHd(__e, tmp4365)

tmp4367 := GTl(__e, tmp4366)

tmp4368 := GConsP(__e, tmp4367)

var ifres4343 Obj

if True == tmp4368 {
tmp4360 := GHd(__e, V877)

tmp4361 := GHd(__e, tmp4360)

tmp4362 := GTl(__e, tmp4361)

tmp4363 := GTl(__e, tmp4362)

tmp4364 := GConsP(__e, tmp4363)

var ifres4344 Obj

if True == tmp4364 {
tmp4354 := GHd(__e, V877)

tmp4355 := GHd(__e, tmp4354)

tmp4356 := GTl(__e, tmp4355)

tmp4357 := GTl(__e, tmp4356)

tmp4358 := GTl(__e, tmp4357)

tmp4359 := GEq(__e, Nil, tmp4358)

var ifres4345 Obj

if True == tmp4359 {
tmp4351 := GHd(__e, V877)

tmp4352 := GTl(__e, tmp4351)

tmp4353 := GConsP(__e, tmp4352)

var ifres4346 Obj

if True == tmp4353 {
tmp4347 := GHd(__e, V877)

tmp4348 := GTl(__e, tmp4347)

tmp4349 := GTl(__e, tmp4348)

tmp4350 := GEq(__e, Nil, tmp4349)

ifres4346 = tmp4350


} else {
ifres4346 = False


}

ifres4345 = ifres4346


} else {
ifres4345 = False


}

ifres4344 = ifres4345


} else {
ifres4344 = False


}

ifres4343 = ifres4344


} else {
ifres4343 = False


}

ifres4342 = ifres4343


} else {
ifres4342 = False


}

ifres4341 = ifres4342


} else {
ifres4341 = False


}

ifres4340 = ifres4341


} else {
ifres4340 = False


}

ifres4339 = ifres4340


} else {
ifres4339 = False


}

if True == ifres4339 {
tmp4261 := GHd(__e, V877)

tmp4262 := GHd(__e, tmp4261)

tmp4263 := GTl(__e, tmp4262)

tmp4264 := GHd(__e, tmp4263)

tmp4265 := Call3(__e, SymFn(symshen_4pivot_1on), tmp4264, V877, Nil)


let__2654 := tmp4265
_ = let__2654

tmp4266 := Call1(__e, SymFn(symfst), let__2654)


let__2655 := tmp4266
_ = let__2655

tmp4310 := Call1(__e, SymFn(symshen_4bad_1pivot_2), let__2655)


if True == tmp4310 {
tmp4267 := GCons(__e, symif, Nil)

tmp4268 := GHd(__e, V877)

tmp4269 := GHd(__e, tmp4268)

tmp4270 := GHd(__e, V877)

tmp4271 := GTl(__e, tmp4270)

tmp4272 := GHd(__e, tmp4271)

tmp4273 := Call2(__e, SymFn(symshen_4recursively_1factor_1selectors), tmp4269, tmp4272)


tmp4274 := GTl(__e, V877)

tmp4275 := Call1(__e, SymFn(symshen_4factor_1recognisors), tmp4274)


tmp4276 := GCons(__e, tmp4275, Nil)

tmp4277 := Call2(__e, SymFn(symappend), tmp4273, tmp4276)


return TailCall2(__e, SymFn(symappend), tmp4267, tmp4277)


} else {
tmp4278 := Call1(__e, SymFn(symsnd), let__2654)


let__2656 := tmp4278
_ = let__2656

tmp4279 := Call1(__e, SymFn(symshen_4factor_1recognisors), let__2656)


let__2657 := tmp4279
_ = let__2657

tmp4280 := Call1(__e, SymFn(symgensym), symGoTo)


let__2658 := tmp4280
_ = let__2658

tmp4281 := GCons(__e, let__2658, Nil)

tmp4282 := GCons(__e, symthaw, tmp4281)

tmp4283 := GCons(__e, tmp4282, Nil)

tmp4284 := GCons(__e, True, tmp4283)

tmp4285 := GCons(__e, tmp4284, let__2655)

tmp4286 := Call1(__e, SymFn(symreverse), tmp4285)


let__2659 := tmp4286
_ = let__2659

tmp4287 := GCons(__e, let__2657, Nil)

tmp4288 := GCons(__e, symfreeze, tmp4287)

tmp4289 := GHd(__e, V877)

tmp4290 := GHd(__e, tmp4289)

tmp4291 := GTl(__e, tmp4290)

tmp4292 := GHd(__e, tmp4291)

tmp4293 := GHd(__e, V877)

tmp4294 := GHd(__e, tmp4293)

tmp4295 := GTl(__e, tmp4294)

tmp4296 := GHd(__e, tmp4295)

tmp4297 := Call1(__e, SymFn(symshen_4factor_1recognisors), let__2659)


tmp4298 := Call2(__e, SymFn(symshen_4factor_1selectors), tmp4296, tmp4297)


tmp4299 := GCons(__e, let__2658, Nil)

tmp4300 := GCons(__e, symthaw, tmp4299)

tmp4301 := GCons(__e, tmp4300, Nil)

tmp4302 := GCons(__e, tmp4298, tmp4301)

tmp4303 := GCons(__e, tmp4292, tmp4302)

tmp4304 := GCons(__e, symif, tmp4303)

tmp4305 := GCons(__e, tmp4304, Nil)

tmp4306 := GCons(__e, tmp4288, tmp4305)

tmp4307 := GCons(__e, let__2658, tmp4306)

tmp4308 := GCons(__e, symlet, tmp4307)

let__2660 := tmp4308
_ = let__2660

return TailCall1(__e, SymFn(symshen_4remove_1indirection), let__2660)


}


} else {
tmp4337 := GConsP(__e, V877)

var ifres4325 Obj

if True == tmp4337 {
tmp4335 := GHd(__e, V877)

tmp4336 := GConsP(__e, tmp4335)

var ifres4326 Obj

if True == tmp4336 {
tmp4332 := GHd(__e, V877)

tmp4333 := GTl(__e, tmp4332)

tmp4334 := GConsP(__e, tmp4333)

var ifres4327 Obj

if True == tmp4334 {
tmp4328 := GHd(__e, V877)

tmp4329 := GTl(__e, tmp4328)

tmp4330 := GTl(__e, tmp4329)

tmp4331 := GEq(__e, Nil, tmp4330)

ifres4327 = tmp4331


} else {
ifres4327 = False


}

ifres4326 = ifres4327


} else {
ifres4326 = False


}

ifres4325 = ifres4326


} else {
ifres4325 = False


}

if True == ifres4325 {
tmp4311 := GHd(__e, V877)

tmp4312 := GHd(__e, tmp4311)

tmp4313 := GHd(__e, V877)

tmp4314 := GTl(__e, tmp4313)

tmp4315 := GHd(__e, tmp4314)

tmp4316 := GTl(__e, V877)

tmp4317 := Call1(__e, SymFn(symshen_4factor_1recognisors), tmp4316)


tmp4318 := GCons(__e, tmp4317, Nil)

tmp4319 := GCons(__e, tmp4315, tmp4318)

tmp4320 := GCons(__e, tmp4312, tmp4319)

return GCons(__e, symif, tmp4320)


} else {
tmp4323 := GEq(__e, Nil, V877)

if True == tmp4323 {
tmp4321 := GCons(__e, MakeString("condition failure"), Nil)

return GCons(__e, symsimple_1error, tmp4321)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4factor_1recognisors)
}


}


}


}


panic("unreachable")
})

tmp4397 := Call2(__e, ns2_1set, symshen_4factor_1recognisors, tmp4258)


_ = tmp4397

tmp4398 := MakeNative2(func(__e *ControlFlow, V885, V886 Obj) Obj {
tmp4461 := GConsP(__e, V885)

var ifres4440 Obj

if True == tmp4461 {
tmp4459 := GHd(__e, V885)

tmp4460 := GEq(__e, symlet, tmp4459)

var ifres4441 Obj

if True == tmp4460 {
tmp4457 := GTl(__e, V885)

tmp4458 := GConsP(__e, tmp4457)

var ifres4442 Obj

if True == tmp4458 {
tmp4454 := GTl(__e, V885)

tmp4455 := GTl(__e, tmp4454)

tmp4456 := GConsP(__e, tmp4455)

var ifres4443 Obj

if True == tmp4456 {
tmp4450 := GTl(__e, V885)

tmp4451 := GTl(__e, tmp4450)

tmp4452 := GTl(__e, tmp4451)

tmp4453 := GConsP(__e, tmp4452)

var ifres4444 Obj

if True == tmp4453 {
tmp4445 := GTl(__e, V885)

tmp4446 := GTl(__e, tmp4445)

tmp4447 := GTl(__e, tmp4446)

tmp4448 := GTl(__e, tmp4447)

tmp4449 := GEq(__e, Nil, tmp4448)

ifres4444 = tmp4449


} else {
ifres4444 = False


}

ifres4443 = ifres4444


} else {
ifres4443 = False


}

ifres4442 = ifres4443


} else {
ifres4442 = False


}

ifres4441 = ifres4442


} else {
ifres4441 = False


}

ifres4440 = ifres4441


} else {
ifres4440 = False


}

if True == ifres4440 {
tmp4399 := GTl(__e, V885)

tmp4400 := GHd(__e, tmp4399)

tmp4401 := GTl(__e, V885)

tmp4402 := GTl(__e, tmp4401)

tmp4403 := GHd(__e, tmp4402)

tmp4404 := GTl(__e, V885)

tmp4405 := GTl(__e, tmp4404)

tmp4406 := GTl(__e, tmp4405)

tmp4407 := GHd(__e, tmp4406)

tmp4408 := Call2(__e, SymFn(symshen_4recursively_1factor_1selectors), tmp4407, V886)


return TailCall3(__e, SymFn(symshen_4restore_1local), tmp4400, tmp4403, tmp4408)


} else {
tmp4438 := GConsP(__e, V885)

var ifres4423 Obj

if True == tmp4438 {
tmp4436 := GHd(__e, V885)

tmp4437 := GEq(__e, symand, tmp4436)

var ifres4424 Obj

if True == tmp4437 {
tmp4434 := GTl(__e, V885)

tmp4435 := GConsP(__e, tmp4434)

var ifres4425 Obj

if True == tmp4435 {
tmp4431 := GTl(__e, V885)

tmp4432 := GTl(__e, tmp4431)

tmp4433 := GConsP(__e, tmp4432)

var ifres4426 Obj

if True == tmp4433 {
tmp4427 := GTl(__e, V885)

tmp4428 := GTl(__e, tmp4427)

tmp4429 := GTl(__e, tmp4428)

tmp4430 := GEq(__e, Nil, tmp4429)

ifres4426 = tmp4430


} else {
ifres4426 = False


}

ifres4425 = ifres4426


} else {
ifres4425 = False


}

ifres4424 = ifres4425


} else {
ifres4424 = False


}

ifres4423 = ifres4424


} else {
ifres4423 = False


}

if True == ifres4423 {
tmp4409 := GTl(__e, V885)

tmp4410 := GHd(__e, tmp4409)

tmp4411 := GTl(__e, V885)

tmp4412 := GHd(__e, tmp4411)

tmp4413 := GTl(__e, V885)

tmp4414 := GTl(__e, tmp4413)

tmp4415 := GHd(__e, tmp4414)

tmp4416 := Call2(__e, SymFn(symshen_4factor_1selectors), tmp4412, tmp4415)


tmp4417 := GTl(__e, V885)

tmp4418 := GHd(__e, tmp4417)

tmp4419 := Call2(__e, SymFn(symshen_4factor_1selectors), tmp4418, V886)


tmp4420 := Call2(__e, SymFn(symshen_4recursively_1factor_1selectors), tmp4416, tmp4419)


return TailCall2(__e, SymFn(symshen_4restore_1P), tmp4410, tmp4420)


} else {
tmp4421 := GCons(__e, V886, Nil)

return GCons(__e, V885, tmp4421)


}


}


panic("unreachable")
})

tmp4462 := Call2(__e, ns2_1set, symshen_4recursively_1factor_1selectors, tmp4398)


_ = tmp4462

tmp4463 := MakeNative2(func(__e *ControlFlow, V887, V888 Obj) Obj {
tmp4477 := GConsP(__e, V888)

var ifres4470 Obj

if True == tmp4477 {
tmp4475 := GTl(__e, V888)

tmp4476 := GConsP(__e, tmp4475)

var ifres4471 Obj

if True == tmp4476 {
tmp4472 := GTl(__e, V888)

tmp4473 := GTl(__e, tmp4472)

tmp4474 := GEq(__e, Nil, tmp4473)

ifres4471 = tmp4474


} else {
ifres4471 = False


}

ifres4470 = ifres4471


} else {
ifres4470 = False


}

if True == ifres4470 {
tmp4464 := GHd(__e, V888)

tmp4465 := GCons(__e, tmp4464, Nil)

tmp4466 := GCons(__e, V887, tmp4465)

tmp4467 := GCons(__e, symand, tmp4466)

tmp4468 := GTl(__e, V888)

return GCons(__e, tmp4467, tmp4468)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4restore_1P)
}


panic("unreachable")
})

tmp4478 := Call2(__e, ns2_1set, symshen_4restore_1P, tmp4463)


_ = tmp4478

tmp4479 := MakeNative3(func(__e *ControlFlow, V889, V890, V891 Obj) Obj {
tmp4494 := GConsP(__e, V891)

var ifres4487 Obj

if True == tmp4494 {
tmp4492 := GTl(__e, V891)

tmp4493 := GConsP(__e, tmp4492)

var ifres4488 Obj

if True == tmp4493 {
tmp4489 := GTl(__e, V891)

tmp4490 := GTl(__e, tmp4489)

tmp4491 := GEq(__e, Nil, tmp4490)

ifres4488 = tmp4491


} else {
ifres4488 = False


}

ifres4487 = ifres4488


} else {
ifres4487 = False


}

if True == ifres4487 {
tmp4480 := GHd(__e, V891)

tmp4481 := GCons(__e, tmp4480, Nil)

tmp4482 := GCons(__e, V890, tmp4481)

tmp4483 := GCons(__e, V889, tmp4482)

tmp4484 := GCons(__e, symlet, tmp4483)

tmp4485 := GTl(__e, V891)

return GCons(__e, tmp4484, tmp4485)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4restore_1local)
}


panic("unreachable")
})

tmp4495 := Call2(__e, ns2_1set, symshen_4restore_1local, tmp4479)


_ = tmp4495

tmp4496 := MakeNative1(func(__e *ControlFlow, V896 Obj) Obj {
tmp4501 := GConsP(__e, V896)

var ifres4498 Obj

if True == tmp4501 {
tmp4499 := GTl(__e, V896)

tmp4500 := GEq(__e, Nil, tmp4499)

ifres4498 = tmp4500


} else {
ifres4498 = False


}

if True == ifres4498 {
return True
} else {
return False
}


panic("unreachable")
})

tmp4502 := Call2(__e, ns2_1set, symshen_4bad_1pivot_2, tmp4496)


_ = tmp4502

tmp4503 := MakeNative1(func(__e *ControlFlow, V897 Obj) Obj {
tmp4604 := GConsP(__e, V897)

var ifres4518 Obj

if True == tmp4604 {
tmp4602 := GHd(__e, V897)

tmp4603 := GEq(__e, symlet, tmp4602)

var ifres4519 Obj

if True == tmp4603 {
tmp4600 := GTl(__e, V897)

tmp4601 := GConsP(__e, tmp4600)

var ifres4520 Obj

if True == tmp4601 {
tmp4597 := GTl(__e, V897)

tmp4598 := GTl(__e, tmp4597)

tmp4599 := GConsP(__e, tmp4598)

var ifres4521 Obj

if True == tmp4599 {
tmp4593 := GTl(__e, V897)

tmp4594 := GTl(__e, tmp4593)

tmp4595 := GHd(__e, tmp4594)

tmp4596 := GConsP(__e, tmp4595)

var ifres4522 Obj

if True == tmp4596 {
tmp4588 := GTl(__e, V897)

tmp4589 := GTl(__e, tmp4588)

tmp4590 := GHd(__e, tmp4589)

tmp4591 := GHd(__e, tmp4590)

tmp4592 := GEq(__e, symfreeze, tmp4591)

var ifres4523 Obj

if True == tmp4592 {
tmp4583 := GTl(__e, V897)

tmp4584 := GTl(__e, tmp4583)

tmp4585 := GHd(__e, tmp4584)

tmp4586 := GTl(__e, tmp4585)

tmp4587 := GConsP(__e, tmp4586)

var ifres4524 Obj

if True == tmp4587 {
tmp4577 := GTl(__e, V897)

tmp4578 := GTl(__e, tmp4577)

tmp4579 := GHd(__e, tmp4578)

tmp4580 := GTl(__e, tmp4579)

tmp4581 := GHd(__e, tmp4580)

tmp4582 := GConsP(__e, tmp4581)

var ifres4525 Obj

if True == tmp4582 {
tmp4570 := GTl(__e, V897)

tmp4571 := GTl(__e, tmp4570)

tmp4572 := GHd(__e, tmp4571)

tmp4573 := GTl(__e, tmp4572)

tmp4574 := GHd(__e, tmp4573)

tmp4575 := GHd(__e, tmp4574)

tmp4576 := GEq(__e, symthaw, tmp4575)

var ifres4526 Obj

if True == tmp4576 {
tmp4563 := GTl(__e, V897)

tmp4564 := GTl(__e, tmp4563)

tmp4565 := GHd(__e, tmp4564)

tmp4566 := GTl(__e, tmp4565)

tmp4567 := GHd(__e, tmp4566)

tmp4568 := GTl(__e, tmp4567)

tmp4569 := GConsP(__e, tmp4568)

var ifres4527 Obj

if True == tmp4569 {
tmp4555 := GTl(__e, V897)

tmp4556 := GTl(__e, tmp4555)

tmp4557 := GHd(__e, tmp4556)

tmp4558 := GTl(__e, tmp4557)

tmp4559 := GHd(__e, tmp4558)

tmp4560 := GTl(__e, tmp4559)

tmp4561 := GTl(__e, tmp4560)

tmp4562 := GEq(__e, Nil, tmp4561)

var ifres4528 Obj

if True == tmp4562 {
tmp4549 := GTl(__e, V897)

tmp4550 := GTl(__e, tmp4549)

tmp4551 := GHd(__e, tmp4550)

tmp4552 := GTl(__e, tmp4551)

tmp4553 := GTl(__e, tmp4552)

tmp4554 := GEq(__e, Nil, tmp4553)

var ifres4529 Obj

if True == tmp4554 {
tmp4545 := GTl(__e, V897)

tmp4546 := GTl(__e, tmp4545)

tmp4547 := GTl(__e, tmp4546)

tmp4548 := GConsP(__e, tmp4547)

var ifres4530 Obj

if True == tmp4548 {
tmp4540 := GTl(__e, V897)

tmp4541 := GTl(__e, tmp4540)

tmp4542 := GTl(__e, tmp4541)

tmp4543 := GTl(__e, tmp4542)

tmp4544 := GEq(__e, Nil, tmp4543)

var ifres4531 Obj

if True == tmp4544 {
tmp4532 := GTl(__e, V897)

tmp4533 := GTl(__e, tmp4532)

tmp4534 := GHd(__e, tmp4533)

tmp4535 := GTl(__e, tmp4534)

tmp4536 := GHd(__e, tmp4535)

tmp4537 := GTl(__e, tmp4536)

tmp4538 := GHd(__e, tmp4537)

tmp4539 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, tmp4538)

ifres4531 = tmp4539


} else {
ifres4531 = False


}

ifres4530 = ifres4531


} else {
ifres4530 = False


}

ifres4529 = ifres4530


} else {
ifres4529 = False


}

ifres4528 = ifres4529


} else {
ifres4528 = False


}

ifres4527 = ifres4528


} else {
ifres4527 = False


}

ifres4526 = ifres4527


} else {
ifres4526 = False


}

ifres4525 = ifres4526


} else {
ifres4525 = False


}

ifres4524 = ifres4525


} else {
ifres4524 = False


}

ifres4523 = ifres4524


} else {
ifres4523 = False


}

ifres4522 = ifres4523


} else {
ifres4522 = False


}

ifres4521 = ifres4522


} else {
ifres4521 = False


}

ifres4520 = ifres4521


} else {
ifres4520 = False


}

ifres4519 = ifres4520


} else {
ifres4519 = False


}

ifres4518 = ifres4519


} else {
ifres4518 = False


}

if True == ifres4518 {
tmp4504 := GTl(__e, V897)

tmp4505 := GTl(__e, tmp4504)

tmp4506 := GHd(__e, tmp4505)

tmp4507 := GTl(__e, tmp4506)

tmp4508 := GHd(__e, tmp4507)

tmp4509 := GTl(__e, tmp4508)

tmp4510 := GHd(__e, tmp4509)

tmp4511 := GTl(__e, V897)

tmp4512 := GHd(__e, tmp4511)

tmp4513 := GTl(__e, V897)

tmp4514 := GTl(__e, tmp4513)

tmp4515 := GTl(__e, tmp4514)

tmp4516 := GHd(__e, tmp4515)

return TailCall3(__e, SymFn(symsubst), tmp4510, tmp4512, tmp4516)


} else {
return V897
}


panic("unreachable")
})

tmp4605 := Call2(__e, ns2_1set, symshen_4remove_1indirection, tmp4503)


_ = tmp4605

tmp4606 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V900 := __self1
_ = V900
V901 := __self2
_ = V901
V902 := __self3
_ = V902
tmp4692 := GConsP(__e, V901)

var ifres4647 Obj

if True == tmp4692 {
tmp4690 := GHd(__e, V901)

tmp4691 := GConsP(__e, tmp4690)

var ifres4648 Obj

if True == tmp4691 {
tmp4687 := GHd(__e, V901)

tmp4688 := GHd(__e, tmp4687)

tmp4689 := GConsP(__e, tmp4688)

var ifres4649 Obj

if True == tmp4689 {
tmp4683 := GHd(__e, V901)

tmp4684 := GHd(__e, tmp4683)

tmp4685 := GHd(__e, tmp4684)

tmp4686 := GEq(__e, symand, tmp4685)

var ifres4650 Obj

if True == tmp4686 {
tmp4679 := GHd(__e, V901)

tmp4680 := GHd(__e, tmp4679)

tmp4681 := GTl(__e, tmp4680)

tmp4682 := GConsP(__e, tmp4681)

var ifres4651 Obj

if True == tmp4682 {
tmp4674 := GHd(__e, V901)

tmp4675 := GHd(__e, tmp4674)

tmp4676 := GTl(__e, tmp4675)

tmp4677 := GTl(__e, tmp4676)

tmp4678 := GConsP(__e, tmp4677)

var ifres4652 Obj

if True == tmp4678 {
tmp4668 := GHd(__e, V901)

tmp4669 := GHd(__e, tmp4668)

tmp4670 := GTl(__e, tmp4669)

tmp4671 := GTl(__e, tmp4670)

tmp4672 := GTl(__e, tmp4671)

tmp4673 := GEq(__e, Nil, tmp4672)

var ifres4653 Obj

if True == tmp4673 {
tmp4665 := GHd(__e, V901)

tmp4666 := GTl(__e, tmp4665)

tmp4667 := GConsP(__e, tmp4666)

var ifres4654 Obj

if True == tmp4667 {
tmp4661 := GHd(__e, V901)

tmp4662 := GTl(__e, tmp4661)

tmp4663 := GTl(__e, tmp4662)

tmp4664 := GEq(__e, Nil, tmp4663)

var ifres4655 Obj

if True == tmp4664 {
tmp4656 := GHd(__e, V901)

tmp4657 := GHd(__e, tmp4656)

tmp4658 := GTl(__e, tmp4657)

tmp4659 := GHd(__e, tmp4658)

tmp4660 := GEq(__e, V900, tmp4659)

ifres4655 = tmp4660


} else {
ifres4655 = False


}

ifres4654 = ifres4655


} else {
ifres4654 = False


}

ifres4653 = ifres4654


} else {
ifres4653 = False


}

ifres4652 = ifres4653


} else {
ifres4652 = False


}

ifres4651 = ifres4652


} else {
ifres4651 = False


}

ifres4650 = ifres4651


} else {
ifres4650 = False


}

ifres4649 = ifres4650


} else {
ifres4649 = False


}

ifres4648 = ifres4649


} else {
ifres4648 = False


}

ifres4647 = ifres4648


} else {
ifres4647 = False


}

if True == ifres4647 {
tmp4607 := GHd(__e, V901)

tmp4608 := GHd(__e, tmp4607)

tmp4609 := GTl(__e, tmp4608)

tmp4610 := GHd(__e, tmp4609)

tmp4611 := GTl(__e, V901)

tmp4612 := GHd(__e, V901)

tmp4613 := GHd(__e, tmp4612)

tmp4614 := GTl(__e, tmp4613)

tmp4615 := GTl(__e, tmp4614)

tmp4616 := GHd(__e, tmp4615)

tmp4617 := GHd(__e, V901)

tmp4618 := GTl(__e, tmp4617)

tmp4619 := GCons(__e, tmp4616, tmp4618)

tmp4620 := GCons(__e, tmp4619, V902)

if SymFn(symshen_4pivot_1on) == __self {
__self1, __self2, __self3 = tmp4610, tmp4611, tmp4620
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4pivot_1on), tmp4610, tmp4611, tmp4620)


} else {
tmp4645 := GConsP(__e, V901)

var ifres4629 Obj

if True == tmp4645 {
tmp4643 := GHd(__e, V901)

tmp4644 := GConsP(__e, tmp4643)

var ifres4630 Obj

if True == tmp4644 {
tmp4640 := GHd(__e, V901)

tmp4641 := GTl(__e, tmp4640)

tmp4642 := GConsP(__e, tmp4641)

var ifres4631 Obj

if True == tmp4642 {
tmp4636 := GHd(__e, V901)

tmp4637 := GTl(__e, tmp4636)

tmp4638 := GTl(__e, tmp4637)

tmp4639 := GEq(__e, Nil, tmp4638)

var ifres4632 Obj

if True == tmp4639 {
tmp4633 := GHd(__e, V901)

tmp4634 := GHd(__e, tmp4633)

tmp4635 := GEq(__e, V900, tmp4634)

ifres4632 = tmp4635


} else {
ifres4632 = False


}

ifres4631 = ifres4632


} else {
ifres4631 = False


}

ifres4630 = ifres4631


} else {
ifres4630 = False


}

ifres4629 = ifres4630


} else {
ifres4629 = False


}

if True == ifres4629 {
tmp4621 := GHd(__e, V901)

tmp4622 := GHd(__e, tmp4621)

tmp4623 := GTl(__e, V901)

tmp4624 := GHd(__e, V901)

tmp4625 := GTl(__e, tmp4624)

tmp4626 := GCons(__e, True, tmp4625)

tmp4627 := GCons(__e, tmp4626, V902)

if SymFn(symshen_4pivot_1on) == __self {
__self1, __self2, __self3 = tmp4622, tmp4623, tmp4627
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4pivot_1on), tmp4622, tmp4623, tmp4627)


} else {
return TailCall2(__e, SymFn(sym_8p), V902, V901)
}


}


panic("unreachable")
})
return __self
})()

tmp4693 := Call2(__e, ns2_1set, symshen_4pivot_1on, tmp4606)


_ = tmp4693

tmp4694 := MakeNative2(func(__e *ControlFlow, V905, V906 Obj) Obj {
tmp4715 := GConsP(__e, V905)

var ifres4708 Obj

if True == tmp4715 {
tmp4713 := GTl(__e, V905)

tmp4714 := GConsP(__e, tmp4713)

var ifres4709 Obj

if True == tmp4714 {
tmp4710 := GTl(__e, V905)

tmp4711 := GTl(__e, tmp4710)

tmp4712 := GEq(__e, Nil, tmp4711)

ifres4709 = tmp4712


} else {
ifres4709 = False


}

ifres4708 = ifres4709


} else {
ifres4708 = False


}

if True == ifres4708 {
tmp4695 := GHd(__e, V905)

tmp4696 := Call1(__e, SymFn(symshen_4op), tmp4695)


let__2661 := tmp4696
_ = let__2661

tmp4706 := GEq(__e, symshen_4skip, let__2661)

if True == tmp4706 {
return V906
} else {
tmp4697 := Call1(__e, SymFn(symshen_4op1), let__2661)


tmp4698 := GTl(__e, V905)

tmp4699 := GCons(__e, tmp4697, tmp4698)

tmp4700 := Call1(__e, SymFn(symshen_4op2), let__2661)


tmp4701 := GTl(__e, V905)

tmp4702 := GCons(__e, tmp4700, tmp4701)

tmp4703 := GCons(__e, tmp4702, Nil)

tmp4704 := GCons(__e, tmp4699, tmp4703)

return TailCall2(__e, SymFn(symshen_4factor_1selectors_1h), tmp4704, V906)


}


} else {
return V906
}


panic("unreachable")
})

tmp4716 := Call2(__e, ns2_1set, symshen_4factor_1selectors, tmp4694)


_ = tmp4716

tmp4717 := MakeNative1(func(__e *ControlFlow, V910 Obj) Obj {
tmp4725 := GEq(__e, symcons_2, V910)

if True == tmp4725 {
return symcons
} else {
tmp4723 := GEq(__e, symshen_4_7string_2, V910)

if True == tmp4723 {
return sym_8s
} else {
tmp4721 := GEq(__e, symshen_4_7vector_2, V910)

if True == tmp4721 {
return sym_8v
} else {
tmp4719 := GEq(__e, symtuple_2, V910)

if True == tmp4719 {
return sym_8p
} else {
return symshen_4skip
}


}


}


}


panic("unreachable")
})

tmp4726 := Call2(__e, ns2_1set, symshen_4op, tmp4717)


_ = tmp4726

tmp4727 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V911 := __self1
_ = V911
V912 := __self2
_ = V912
tmp4745 := GEq(__e, Nil, V911)

if True == tmp4745 {
return V912
} else {
tmp4743 := GConsP(__e, V911)

if True == tmp4743 {
tmp4739 := GHd(__e, V911)

tmp4740 := Call2(__e, SymFn(symoccurrences), tmp4739, V912)


if True == GGt(__e, tmp4740, MakeInteger(1)) {
tmp4728 := Call1(__e, SymFn(symgensym), symSelect)


let__2662 := tmp4728
_ = let__2662

tmp4729 := GHd(__e, V911)

tmp4730 := GTl(__e, V911)

tmp4731 := GHd(__e, V911)

tmp4732 := Call3(__e, SymFn(symsubst), let__2662, tmp4731, V912)


tmp4733 := Call2(__e, SymFn(symshen_4factor_1selectors_1h), tmp4730, tmp4732)


tmp4734 := GCons(__e, tmp4733, Nil)

tmp4735 := GCons(__e, tmp4729, tmp4734)

tmp4736 := GCons(__e, let__2662, tmp4735)

return GCons(__e, symlet, tmp4736)


} else {
tmp4737 := GTl(__e, V911)

if SymFn(symshen_4factor_1selectors_1h) == __self {
__self1, __self2 = tmp4737, V912
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4factor_1selectors_1h), tmp4737, V912)


}


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4factor_1selectors_1h)
}


}


panic("unreachable")
})
return __self
})()

return TailCall2(__e, ns2_1set, symshen_4factor_1selectors_1h, tmp4727)




panic("unreachable")
})

