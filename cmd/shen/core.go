package main

import . "github.com/pyrex41/shen-go/kl"

var CoreMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp2646 := MakeNative1(func(__e *ControlFlow, V540 Obj) Obj {
tmp2647 := Call1(__e, SymFn(symshen_4shen_1_6kl_1h), V540)


let__2494 := tmp2647
_ = let__2494

return TailCall1(__e, SymFn(symshen_4record_1and_1evaluate), let__2494)


panic("unreachable")
})

tmp2648 := Call2(__e, ns2_1set, symshen_4shen_1_6kl, tmp2646)


_ = tmp2648

tmp2649 := MakeNative1(func(__e *ControlFlow, V542 Obj) Obj {
tmp2693 := GConsP(__e, V542)

var ifres2672 Obj

if True == tmp2693 {
tmp2691 := GHd(__e, V542)

tmp2692 := GEq(__e, symdefun, tmp2691)

var ifres2673 Obj

if True == tmp2692 {
tmp2689 := GTl(__e, V542)

tmp2690 := GConsP(__e, tmp2689)

var ifres2674 Obj

if True == tmp2690 {
tmp2686 := GTl(__e, V542)

tmp2687 := GTl(__e, tmp2686)

tmp2688 := GConsP(__e, tmp2687)

var ifres2675 Obj

if True == tmp2688 {
tmp2682 := GTl(__e, V542)

tmp2683 := GTl(__e, tmp2682)

tmp2684 := GTl(__e, tmp2683)

tmp2685 := GConsP(__e, tmp2684)

var ifres2676 Obj

if True == tmp2685 {
tmp2677 := GTl(__e, V542)

tmp2678 := GTl(__e, tmp2677)

tmp2679 := GTl(__e, tmp2678)

tmp2680 := GTl(__e, tmp2679)

tmp2681 := GEq(__e, Nil, tmp2680)

ifres2676 = tmp2681


} else {
ifres2676 = False


}

ifres2675 = ifres2676


} else {
ifres2675 = False


}

ifres2674 = ifres2675


} else {
ifres2674 = False


}

ifres2673 = ifres2674


} else {
ifres2673 = False


}

ifres2672 = ifres2673


} else {
ifres2672 = False


}

if True == ifres2672 {
tmp2655 := GTl(__e, V542)

tmp2656 := GHd(__e, tmp2655)

tmp2657 := Call1(__e, SymFn(symshen_4sysfunc_2), tmp2656)


var ifres2650 Obj

if True == tmp2657 {
tmp2651 := GTl(__e, V542)

tmp2652 := GHd(__e, tmp2651)

tmp2653 := Call3(__e, SymFn(symshen_4app), tmp2652, MakeString(" is not a legitimate function name\n"), symshen_4a)


tmp2654 := PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp2653)

ifres2650 = tmp2654


} else {
ifres2650 = symshen_4skip


}

let__2495 := ifres2650
_ = let__2495

tmp2658 := GTl(__e, V542)

tmp2659 := GHd(__e, tmp2658)

tmp2660 := GTl(__e, V542)

tmp2661 := GTl(__e, tmp2660)

tmp2662 := GHd(__e, tmp2661)

tmp2663 := Call1(__e, SymFn(symlength), tmp2662)


tmp2664 := Call2(__e, SymFn(symshen_4store_1arity), tmp2659, tmp2663)


let__2496 := tmp2664
_ = let__2496

tmp2665 := GTl(__e, V542)

tmp2666 := GHd(__e, tmp2665)

tmp2667 := Call2(__e, SymFn(symshen_4record_1kl), tmp2666, V542)


let__2497 := tmp2667
_ = let__2497

tmp2668 := Call1(__e, SymFn(symeval_1kl), V542)


let__2498 := tmp2668
_ = let__2498

tmp2669 := GTl(__e, V542)

tmp2670 := GHd(__e, tmp2669)

return TailCall1(__e, SymFn(symshen_4fn_1print), tmp2670)


} else {
return V542
}


panic("unreachable")
})

tmp2694 := Call2(__e, ns2_1set, symshen_4record_1and_1evaluate, tmp2649)


_ = tmp2694

tmp2695 := MakeNative1(func(__e *ControlFlow, V547 Obj) Obj {
tmp2781 := GConsP(__e, V547)

var ifres2775 Obj

if True == tmp2781 {
tmp2779 := GHd(__e, V547)

tmp2780 := GEq(__e, symdefine, tmp2779)

var ifres2776 Obj

if True == tmp2780 {
tmp2777 := GTl(__e, V547)

tmp2778 := GConsP(__e, tmp2777)

ifres2776 = tmp2778


} else {
ifres2776 = False


}

ifres2775 = ifres2776


} else {
ifres2775 = False


}

if True == ifres2775 {
tmp2696 := GTl(__e, V547)

tmp2697 := GHd(__e, tmp2696)

tmp2698 := GTl(__e, V547)

tmp2699 := GTl(__e, tmp2698)

return TailCall2(__e, SymFn(symshen_4shendef_1_6kldef), tmp2697, tmp2699)


} else {
tmp2773 := GConsP(__e, V547)

var ifres2752 Obj

if True == tmp2773 {
tmp2771 := GHd(__e, V547)

tmp2772 := GEq(__e, symdefun, tmp2771)

var ifres2753 Obj

if True == tmp2772 {
tmp2769 := GTl(__e, V547)

tmp2770 := GConsP(__e, tmp2769)

var ifres2754 Obj

if True == tmp2770 {
tmp2766 := GTl(__e, V547)

tmp2767 := GTl(__e, tmp2766)

tmp2768 := GConsP(__e, tmp2767)

var ifres2755 Obj

if True == tmp2768 {
tmp2762 := GTl(__e, V547)

tmp2763 := GTl(__e, tmp2762)

tmp2764 := GTl(__e, tmp2763)

tmp2765 := GConsP(__e, tmp2764)

var ifres2756 Obj

if True == tmp2765 {
tmp2757 := GTl(__e, V547)

tmp2758 := GTl(__e, tmp2757)

tmp2759 := GTl(__e, tmp2758)

tmp2760 := GTl(__e, tmp2759)

tmp2761 := GEq(__e, Nil, tmp2760)

ifres2756 = tmp2761


} else {
ifres2756 = False


}

ifres2755 = ifres2756


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
return V547
} else {
tmp2750 := GConsP(__e, V547)

var ifres2735 Obj

if True == tmp2750 {
tmp2748 := GHd(__e, V547)

tmp2749 := GEq(__e, symtype, tmp2748)

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
tmp2700 := GTl(__e, V547)

tmp2701 := GHd(__e, tmp2700)

tmp2702 := GTl(__e, V547)

tmp2703 := GTl(__e, tmp2702)

tmp2704 := GHd(__e, tmp2703)

tmp2705 := Call1(__e, SymFn(symshen_4rcons__form), tmp2704)


tmp2706 := GCons(__e, tmp2705, Nil)

tmp2707 := GCons(__e, tmp2701, tmp2706)

return GCons(__e, symtype, tmp2707)


} else {
tmp2733 := GConsP(__e, V547)

var ifres2718 Obj

if True == tmp2733 {
tmp2731 := GHd(__e, V547)

tmp2732 := GEq(__e, syminput_7, tmp2731)

var ifres2719 Obj

if True == tmp2732 {
tmp2729 := GTl(__e, V547)

tmp2730 := GConsP(__e, tmp2729)

var ifres2720 Obj

if True == tmp2730 {
tmp2726 := GTl(__e, V547)

tmp2727 := GTl(__e, tmp2726)

tmp2728 := GConsP(__e, tmp2727)

var ifres2721 Obj

if True == tmp2728 {
tmp2722 := GTl(__e, V547)

tmp2723 := GTl(__e, tmp2722)

tmp2724 := GTl(__e, tmp2723)

tmp2725 := GEq(__e, Nil, tmp2724)

ifres2721 = tmp2725


} else {
ifres2721 = False


}

ifres2720 = ifres2721


} else {
ifres2720 = False


}

ifres2719 = ifres2720


} else {
ifres2719 = False


}

ifres2718 = ifres2719


} else {
ifres2718 = False


}

if True == ifres2718 {
tmp2708 := GTl(__e, V547)

tmp2709 := GHd(__e, tmp2708)

tmp2710 := Call1(__e, SymFn(symshen_4rcons__form), tmp2709)


tmp2711 := GTl(__e, V547)

tmp2712 := GTl(__e, tmp2711)

tmp2713 := GCons(__e, tmp2710, tmp2712)

return GCons(__e, syminput_7, tmp2713)


} else {
tmp2716 := GConsP(__e, V547)

if True == tmp2716 {
tmp2714 := MakeNative1(func(__e *ControlFlow, Z548 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4shen_1_6kl_1h), Z548)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp2714, V547)


} else {
return V547
}


}


}


}


}


panic("unreachable")
})

tmp2782 := Call2(__e, ns2_1set, symshen_4shen_1_6kl_1h, tmp2695)


_ = tmp2782

tmp2783 := MakeNative2(func(__e *ControlFlow, V549, V550 Obj) Obj {
tmp2784 := MakeNative1(func(__e *ControlFlow, Z551 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5define_6), Z551)
panic("unreachable")
})

tmp2785 := GCons(__e, V549, V550)

return TailCall2(__e, SymFn(symcompile), tmp2784, tmp2785)


panic("unreachable")
})

tmp2786 := Call2(__e, ns2_1set, symshen_4shendef_1_6kldef, tmp2783)


_ = tmp2786

tmp2787 := MakeNative1(func(__e *ControlFlow, V552 Obj) Obj {
tmp2788 := Call1(__e, SymFn(symshen_4_5name_6), V552)


let__2500 := tmp2788
_ = let__2500

tmp2814 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2500)


var ifres2789 Obj

if True == tmp2814 {
tmp2790 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2789 = tmp2790


} else {
tmp2791 := Call1(__e, SymFn(symshen_4_5_1out), let__2500)


let__2501 := tmp2791
_ = let__2501

tmp2792 := Call1(__e, SymFn(symshen_4in_1_6), let__2500)


let__2502 := tmp2792
_ = let__2502

tmp2813 := Call2(__e, SymFn(symshen_4hds_a_2), let__2502, sym_i)


var ifres2793 Obj

if True == tmp2813 {
tmp2794 := Call1(__e, SymFn(symtail), let__2502)


let__2503 := tmp2794
_ = let__2503

tmp2795 := Call1(__e, SymFn(symshen_4_5signature_6), let__2503)


let__2504 := tmp2795
_ = let__2504

tmp2811 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2504)


var ifres2796 Obj

if True == tmp2811 {
tmp2797 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2796 = tmp2797


} else {
tmp2798 := Call1(__e, SymFn(symshen_4in_1_6), let__2504)


let__2505 := tmp2798
_ = let__2505

tmp2810 := Call2(__e, SymFn(symshen_4hds_a_2), let__2505, sym_j)


var ifres2799 Obj

if True == tmp2810 {
tmp2800 := Call1(__e, SymFn(symtail), let__2505)


let__2506 := tmp2800
_ = let__2506

tmp2801 := Call1(__e, SymFn(symshen_4_5rules_6), let__2506)


let__2507 := tmp2801
_ = let__2507

tmp2808 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2507)


var ifres2802 Obj

if True == tmp2808 {
tmp2803 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2802 = tmp2803


} else {
tmp2804 := Call1(__e, SymFn(symshen_4_5_1out), let__2507)


let__2508 := tmp2804
_ = let__2508

tmp2805 := Call1(__e, SymFn(symshen_4in_1_6), let__2507)


let__2509 := tmp2805
_ = let__2509

tmp2806 := Call2(__e, SymFn(symshen_4shendef_1_6kldef_1h), let__2501, let__2508)


tmp2807 := Call2(__e, SymFn(symshen_4comb), let__2509, tmp2806)


ifres2802 = tmp2807


}

ifres2799 = ifres2802


} else {
tmp2809 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2799 = tmp2809


}

ifres2796 = ifres2799


}

ifres2793 = ifres2796


} else {
tmp2812 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2793 = tmp2812


}

ifres2789 = ifres2793


}

let__2499 := ifres2789
_ = let__2499

tmp2832 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2499)


if True == tmp2832 {
tmp2815 := Call1(__e, SymFn(symshen_4_5name_6), V552)


let__2511 := tmp2815
_ = let__2511

tmp2828 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2511)


var ifres2816 Obj

if True == tmp2828 {
tmp2817 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2816 = tmp2817


} else {
tmp2818 := Call1(__e, SymFn(symshen_4_5_1out), let__2511)


let__2512 := tmp2818
_ = let__2512

tmp2819 := Call1(__e, SymFn(symshen_4in_1_6), let__2511)


let__2513 := tmp2819
_ = let__2513

tmp2820 := Call1(__e, SymFn(symshen_4_5rules_6), let__2513)


let__2514 := tmp2820
_ = let__2514

tmp2827 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2514)


var ifres2821 Obj

if True == tmp2827 {
tmp2822 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2821 = tmp2822


} else {
tmp2823 := Call1(__e, SymFn(symshen_4_5_1out), let__2514)


let__2515 := tmp2823
_ = let__2515

tmp2824 := Call1(__e, SymFn(symshen_4in_1_6), let__2514)


let__2516 := tmp2824
_ = let__2516

tmp2825 := Call2(__e, SymFn(symshen_4shendef_1_6kldef_1h), let__2512, let__2515)


tmp2826 := Call2(__e, SymFn(symshen_4comb), let__2516, tmp2825)


ifres2821 = tmp2826


}

ifres2816 = ifres2821


}

let__2510 := ifres2816
_ = let__2510

tmp2830 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2510)


if True == tmp2830 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2510
}


} else {
return let__2499
}


panic("unreachable")
})

tmp2833 := Call2(__e, ns2_1set, symshen_4_5define_6, tmp2787)


_ = tmp2833

tmp2834 := MakeNative2(func(__e *ControlFlow, V571, V572 Obj) Obj {
tmp2835 := MakeNative1(func(__e *ControlFlow, Z574 Obj) Obj {
return TailCall1(__e, SymFn(symfst), Z574)
panic("unreachable")
})

tmp2836 := Call2(__e, SymFn(symmap), tmp2835, V572)


let__2517 := tmp2836
_ = let__2517

tmp2837 := Call2(__e, SymFn(symshen_4arity_1chk), V571, let__2517)


let__2518 := tmp2837
_ = let__2518

tmp2838 := MakeNative1(func(__e *ControlFlow, Z577 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4free_1var_1chk), V571, Z577)
panic("unreachable")
})

tmp2839 := Call2(__e, SymFn(symmap), tmp2838, V572)


let__2519 := tmp2839
_ = let__2519

tmp2840 := Call1(__e, SymFn(symshen_4unprotect), V572)


let__2520 := tmp2840
_ = let__2520

tmp2841 := Call3(__e, SymFn(symshen_4compile_1to_1kl), V571, let__2520, let__2518)


tmp2842 := Call1(__e, SymFn(symshen_4factorise_1code), tmp2841)


let__2521 := tmp2842
_ = let__2521

return let__2521


panic("unreachable")
})

tmp2843 := Call2(__e, ns2_1set, symshen_4shendef_1_6kldef_1h, tmp2834)


_ = tmp2843

tmp2844 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V580 := __self1
_ = V580
tmp2867 := Call1(__e, SymFn(symtuple_2), V580)


if True == tmp2867 {
tmp2845 := Call1(__e, SymFn(symfst), V580)


tmp2846 := Call1(__e, SymFn(symshen_4unprotect), tmp2845)


tmp2847 := Call1(__e, SymFn(symsnd), V580)


tmp2848 := Call1(__e, SymFn(symshen_4unprotect), tmp2847)


return TailCall2(__e, SymFn(sym_8p), tmp2846, tmp2848)


} else {
tmp2865 := GConsP(__e, V580)

var ifres2855 Obj

if True == tmp2865 {
tmp2863 := GHd(__e, V580)

tmp2864 := GEq(__e, symprotect, tmp2863)

var ifres2856 Obj

if True == tmp2864 {
tmp2861 := GTl(__e, V580)

tmp2862 := GConsP(__e, tmp2861)

var ifres2857 Obj

if True == tmp2862 {
tmp2858 := GTl(__e, V580)

tmp2859 := GTl(__e, tmp2858)

tmp2860 := GEq(__e, Nil, tmp2859)

ifres2857 = tmp2860


} else {
ifres2857 = False


}

ifres2856 = ifres2857


} else {
ifres2856 = False


}

ifres2855 = ifres2856


} else {
ifres2855 = False


}

if True == ifres2855 {
tmp2849 := GTl(__e, V580)

tmp2850 := GHd(__e, tmp2849)

if SymFn(symshen_4unprotect) == __self {
__self1 = tmp2850
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4unprotect), tmp2850)


} else {
tmp2853 := GConsP(__e, V580)

if True == tmp2853 {
tmp2851 := MakeNative1(func(__e *ControlFlow, Z581 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4unprotect), Z581)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp2851, V580)


} else {
return V580
}


}


}


panic("unreachable")
})
return __self
})()

tmp2868 := Call2(__e, ns2_1set, symshen_4unprotect, tmp2844)


_ = tmp2868

tmp2869 := MakeNative1(func(__e *ControlFlow, V582 Obj) Obj {
tmp2882 := GConsP(__e, V582)

var ifres2870 Obj

if True == tmp2882 {
tmp2871 := Call1(__e, SymFn(symhead), V582)


let__2523 := tmp2871
_ = let__2523

tmp2872 := Call1(__e, SymFn(symtail), V582)


let__2524 := tmp2872
_ = let__2524

tmp2879 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, let__2523)

var ifres2876 Obj

if True == tmp2879 {
tmp2877 := PrimGuard1(__e, symvariable_2, PrimIsVariable, let__2523)

ifres2876 = GNot(__e, tmp2877)


} else {
ifres2876 = False


}

var ifres2873 Obj

if True == ifres2876 {
ifres2873 = let__2523


} else {
tmp2874 := Call3(__e, SymFn(symshen_4app), let__2523, MakeString(" is not a legitimate function name.\n"), symshen_4a)


tmp2875 := PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp2874)

ifres2873 = tmp2875


}

tmp2880 := Call2(__e, SymFn(symshen_4comb), let__2524, ifres2873)


ifres2870 = tmp2880


} else {
tmp2881 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2870 = tmp2881


}

let__2522 := ifres2870
_ = let__2522

tmp2884 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2522)


if True == tmp2884 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2522
}


panic("unreachable")
})

tmp2885 := Call2(__e, ns2_1set, symshen_4_5name_6, tmp2869)


_ = tmp2885

tmp2886 := MakeNative1(func(__e *ControlFlow, V586 Obj) Obj {
tmp2905 := GConsP(__e, V586)

var ifres2887 Obj

if True == tmp2905 {
tmp2888 := Call1(__e, SymFn(symhead), V586)


let__2526 := tmp2888
_ = let__2526

tmp2889 := Call1(__e, SymFn(symtail), V586)


let__2527 := tmp2889
_ = let__2527

tmp2890 := Call1(__e, SymFn(symshen_4_5signature_6), let__2527)


let__2528 := tmp2890
_ = let__2528

tmp2903 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2528)


var ifres2891 Obj

if True == tmp2903 {
tmp2892 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2891 = tmp2892


} else {
tmp2893 := Call1(__e, SymFn(symshen_4_5_1out), let__2528)


let__2529 := tmp2893
_ = let__2529

tmp2894 := Call1(__e, SymFn(symshen_4in_1_6), let__2528)


let__2530 := tmp2894
_ = let__2530

tmp2899 := GCons(__e, sym_j, Nil)

tmp2900 := GCons(__e, sym_i, tmp2899)

tmp2901 := Call2(__e, SymFn(symelement_2), let__2526, tmp2900)


tmp2902 := GNot(__e, tmp2901)

var ifres2895 Obj

if True == tmp2902 {
tmp2896 := GCons(__e, let__2526, let__2529)

tmp2897 := Call2(__e, SymFn(symshen_4comb), let__2530, tmp2896)


ifres2895 = tmp2897


} else {
tmp2898 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2895 = tmp2898


}

ifres2891 = ifres2895


}

ifres2887 = ifres2891


} else {
tmp2904 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2887 = tmp2904


}

let__2525 := ifres2887
_ = let__2525

tmp2915 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2525)


if True == tmp2915 {
tmp2906 := Call1(__e, SymFn(sym_5e_6), V586)


let__2532 := tmp2906
_ = let__2532

tmp2911 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2532)


var ifres2907 Obj

if True == tmp2911 {
tmp2908 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2907 = tmp2908


} else {
tmp2909 := Call1(__e, SymFn(symshen_4in_1_6), let__2532)


let__2533 := tmp2909
_ = let__2533

tmp2910 := Call2(__e, SymFn(symshen_4comb), let__2533, Nil)


ifres2907 = tmp2910


}

let__2531 := ifres2907
_ = let__2531

tmp2913 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2531)


if True == tmp2913 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2531
}


} else {
return let__2525
}


panic("unreachable")
})

tmp2916 := Call2(__e, ns2_1set, symshen_4_5signature_6, tmp2886)


_ = tmp2916

tmp2917 := MakeNative1(func(__e *ControlFlow, V596 Obj) Obj {
tmp2918 := Call1(__e, SymFn(symshen_4_5rule_6), V596)


let__2535 := tmp2918
_ = let__2535

tmp2932 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2535)


var ifres2919 Obj

if True == tmp2932 {
tmp2920 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2919 = tmp2920


} else {
tmp2921 := Call1(__e, SymFn(symshen_4_5_1out), let__2535)


let__2536 := tmp2921
_ = let__2536

tmp2922 := Call1(__e, SymFn(symshen_4in_1_6), let__2535)


let__2537 := tmp2922
_ = let__2537

tmp2923 := Call1(__e, SymFn(symshen_4_5rules_6), let__2537)


let__2538 := tmp2923
_ = let__2538

tmp2931 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2538)


var ifres2924 Obj

if True == tmp2931 {
tmp2925 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2924 = tmp2925


} else {
tmp2926 := Call1(__e, SymFn(symshen_4_5_1out), let__2538)


let__2539 := tmp2926
_ = let__2539

tmp2927 := Call1(__e, SymFn(symshen_4in_1_6), let__2538)


let__2540 := tmp2927
_ = let__2540

tmp2928 := Call1(__e, SymFn(symshen_4linearise), let__2536)


tmp2929 := GCons(__e, tmp2928, let__2539)

tmp2930 := Call2(__e, SymFn(symshen_4comb), let__2540, tmp2929)


ifres2924 = tmp2930


}

ifres2919 = ifres2924


}

let__2534 := ifres2919
_ = let__2534

tmp2948 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2534)


if True == tmp2948 {
tmp2933 := Call1(__e, SymFn(sym_5_b_6), V596)


let__2542 := tmp2933
_ = let__2542

tmp2944 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2542)


var ifres2934 Obj

if True == tmp2944 {
tmp2935 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2934 = tmp2935


} else {
tmp2936 := Call1(__e, SymFn(symshen_4_5_1out), let__2542)


let__2543 := tmp2936
_ = let__2543

tmp2937 := Call1(__e, SymFn(symshen_4in_1_6), let__2542)


let__2544 := tmp2937
_ = let__2544

tmp2942 := Call1(__e, SymFn(symempty_2), let__2543)


var ifres2938 Obj

if True == tmp2942 {
ifres2938 = Nil


} else {
tmp2939 := Call3(__e, SymFn(symshen_4app), let__2543, MakeString("\n ..."), symshen_4r)


tmp2941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("Shen syntax error here:\n "), tmp2939))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("Shen syntax error here:\n "), tmp2939)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()

ifres2938 = tmp2941


}

tmp2943 := Call2(__e, SymFn(symshen_4comb), let__2544, ifres2938)


ifres2934 = tmp2943


}

let__2541 := ifres2934
_ = let__2541

tmp2946 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2541)


if True == tmp2946 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2541
}


} else {
return let__2534
}


panic("unreachable")
})

tmp2949 := Call2(__e, ns2_1set, symshen_4_5rules_6, tmp2917)


_ = tmp2949

tmp2950 := MakeNative1(func(__e *ControlFlow, V610 Obj) Obj {
tmp2955 := Call1(__e, SymFn(symtuple_2), V610)


if True == tmp2955 {
tmp2951 := Call1(__e, SymFn(symfst), V610)


tmp2952 := Call1(__e, SymFn(symfst), V610)


tmp2953 := Call1(__e, SymFn(symsnd), V610)


return TailCall4(__e, SymFn(symshen_4linearise_1h), tmp2951, tmp2952, Nil, tmp2953)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.linearise"))
}


panic("unreachable")
})

tmp2956 := Call2(__e, ns2_1set, symshen_4linearise, tmp2950)


_ = tmp2956

tmp2957 := (func() Obj {
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
tmp2992 := GEq(__e, Nil, V623)

if True == tmp2992 {
return TailCall2(__e, SymFn(sym_8p), V624, V626)
} else {
tmp2990 := GConsP(__e, V623)

var ifres2987 Obj

if True == tmp2990 {
tmp2988 := GHd(__e, V623)

tmp2989 := GConsP(__e, tmp2988)

ifres2987 = tmp2989


} else {
ifres2987 = False


}

if True == ifres2987 {
tmp2958 := GHd(__e, V623)

tmp2959 := GTl(__e, V623)

tmp2960 := Call2(__e, SymFn(symappend), tmp2958, tmp2959)


if SymFn(symshen_4linearise_1h) == __self {
__self1, __self2, __self3, __self4 = tmp2960, V624, V625, V626
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4linearise_1h), tmp2960, V624, V625, V626)


} else {
tmp2985 := GConsP(__e, V623)

var ifres2982 Obj

if True == tmp2985 {
tmp2983 := GHd(__e, V623)

tmp2984 := PrimGuard1(__e, symvariable_2, PrimIsVariable, tmp2983)

ifres2982 = tmp2984


} else {
ifres2982 = False


}

if True == ifres2982 {
tmp2976 := GHd(__e, V623)

tmp2977 := Call2(__e, SymFn(symelement_2), tmp2976, V625)


if True == tmp2977 {
tmp2961 := Call1(__e, SymFn(symgensym), symV)


let__2545 := tmp2961
_ = let__2545

tmp2962 := GTl(__e, V623)

tmp2963 := GHd(__e, V623)

tmp2964 := Call3(__e, SymFn(symshen_4rep_1X), tmp2963, let__2545, V624)


tmp2965 := GHd(__e, V623)

tmp2966 := GCons(__e, tmp2965, Nil)

tmp2967 := GCons(__e, let__2545, tmp2966)

tmp2968 := GCons(__e, sym_a, tmp2967)

tmp2969 := GCons(__e, V626, Nil)

tmp2970 := GCons(__e, tmp2968, tmp2969)

tmp2971 := GCons(__e, symwhere, tmp2970)

if SymFn(symshen_4linearise_1h) == __self {
__self1, __self2, __self3, __self4 = tmp2962, tmp2964, V625, tmp2971
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4linearise_1h), tmp2962, tmp2964, V625, tmp2971)


} else {
tmp2972 := GTl(__e, V623)

tmp2973 := GHd(__e, V623)

tmp2974 := GCons(__e, tmp2973, V625)

if SymFn(symshen_4linearise_1h) == __self {
__self1, __self2, __self3, __self4 = tmp2972, V624, tmp2974, V626
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4linearise_1h), tmp2972, V624, tmp2974, V626)


}


} else {
tmp2980 := GConsP(__e, V623)

if True == tmp2980 {
tmp2978 := GTl(__e, V623)

if SymFn(symshen_4linearise_1h) == __self {
__self1, __self2, __self3, __self4 = tmp2978, V624, V625, V626
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4linearise_1h), tmp2978, V624, V625, V626)


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

tmp2993 := Call2(__e, ns2_1set, symshen_4linearise_1h, tmp2957)


_ = tmp2993

tmp2994 := MakeNative1(func(__e *ControlFlow, V628 Obj) Obj {
tmp2995 := Call1(__e, SymFn(symshen_4_5patterns_6), V628)


let__2547 := tmp2995
_ = let__2547

tmp3023 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2547)


var ifres2996 Obj

if True == tmp3023 {
tmp2997 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres2996 = tmp2997


} else {
tmp2998 := Call1(__e, SymFn(symshen_4_5_1out), let__2547)


let__2548 := tmp2998
_ = let__2548

tmp2999 := Call1(__e, SymFn(symshen_4in_1_6), let__2547)


let__2549 := tmp2999
_ = let__2549

tmp3022 := Call2(__e, SymFn(symshen_4hds_a_2), let__2549, sym_1_6)


var ifres3000 Obj

if True == tmp3022 {
tmp3001 := Call1(__e, SymFn(symtail), let__2549)


let__2550 := tmp3001
_ = let__2550

tmp3020 := GConsP(__e, let__2550)

var ifres3002 Obj

if True == tmp3020 {
tmp3003 := Call1(__e, SymFn(symhead), let__2550)


let__2551 := tmp3003
_ = let__2551

tmp3004 := Call1(__e, SymFn(symtail), let__2550)


let__2552 := tmp3004
_ = let__2552

tmp3018 := Call2(__e, SymFn(symshen_4hds_a_2), let__2552, symwhere)


var ifres3005 Obj

if True == tmp3018 {
tmp3006 := Call1(__e, SymFn(symtail), let__2552)


let__2553 := tmp3006
_ = let__2553

tmp3016 := GConsP(__e, let__2553)

var ifres3007 Obj

if True == tmp3016 {
tmp3008 := Call1(__e, SymFn(symhead), let__2553)


let__2554 := tmp3008
_ = let__2554

tmp3009 := Call1(__e, SymFn(symtail), let__2553)


let__2555 := tmp3009
_ = let__2555

tmp3010 := GCons(__e, let__2551, Nil)

tmp3011 := GCons(__e, let__2554, tmp3010)

tmp3012 := GCons(__e, symwhere, tmp3011)

tmp3013 := Call2(__e, SymFn(sym_8p), let__2548, tmp3012)


tmp3014 := Call2(__e, SymFn(symshen_4comb), let__2555, tmp3013)


ifres3007 = tmp3014


} else {
tmp3015 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3007 = tmp3015


}

ifres3005 = ifres3007


} else {
tmp3017 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3005 = tmp3017


}

ifres3002 = ifres3005


} else {
tmp3019 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3002 = tmp3019


}

ifres3000 = ifres3002


} else {
tmp3021 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3000 = tmp3021


}

ifres2996 = ifres3000


}

let__2546 := ifres2996
_ = let__2546

tmp3098 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2546)


if True == tmp3098 {
tmp3024 := Call1(__e, SymFn(symshen_4_5patterns_6), V628)


let__2557 := tmp3024
_ = let__2557

tmp3040 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2557)


var ifres3025 Obj

if True == tmp3040 {
tmp3026 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3025 = tmp3026


} else {
tmp3027 := Call1(__e, SymFn(symshen_4_5_1out), let__2557)


let__2558 := tmp3027
_ = let__2558

tmp3028 := Call1(__e, SymFn(symshen_4in_1_6), let__2557)


let__2559 := tmp3028
_ = let__2559

tmp3039 := Call2(__e, SymFn(symshen_4hds_a_2), let__2559, sym_1_6)


var ifres3029 Obj

if True == tmp3039 {
tmp3030 := Call1(__e, SymFn(symtail), let__2559)


let__2560 := tmp3030
_ = let__2560

tmp3037 := GConsP(__e, let__2560)

var ifres3031 Obj

if True == tmp3037 {
tmp3032 := Call1(__e, SymFn(symhead), let__2560)


let__2561 := tmp3032
_ = let__2561

tmp3033 := Call1(__e, SymFn(symtail), let__2560)


let__2562 := tmp3033
_ = let__2562

tmp3034 := Call2(__e, SymFn(sym_8p), let__2558, let__2561)


tmp3035 := Call2(__e, SymFn(symshen_4comb), let__2562, tmp3034)


ifres3031 = tmp3035


} else {
tmp3036 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3031 = tmp3036


}

ifres3029 = ifres3031


} else {
tmp3038 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3029 = tmp3038


}

ifres3025 = ifres3029


}

let__2556 := ifres3025
_ = let__2556

tmp3096 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2556)


if True == tmp3096 {
tmp3041 := Call1(__e, SymFn(symshen_4_5patterns_6), V628)


let__2564 := tmp3041
_ = let__2564

tmp3071 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2564)


var ifres3042 Obj

if True == tmp3071 {
tmp3043 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3042 = tmp3043


} else {
tmp3044 := Call1(__e, SymFn(symshen_4_5_1out), let__2564)


let__2565 := tmp3044
_ = let__2565

tmp3045 := Call1(__e, SymFn(symshen_4in_1_6), let__2564)


let__2566 := tmp3045
_ = let__2566

tmp3070 := Call2(__e, SymFn(symshen_4hds_a_2), let__2566, sym_5_1)


var ifres3046 Obj

if True == tmp3070 {
tmp3047 := Call1(__e, SymFn(symtail), let__2566)


let__2567 := tmp3047
_ = let__2567

tmp3068 := GConsP(__e, let__2567)

var ifres3048 Obj

if True == tmp3068 {
tmp3049 := Call1(__e, SymFn(symhead), let__2567)


let__2568 := tmp3049
_ = let__2568

tmp3050 := Call1(__e, SymFn(symtail), let__2567)


let__2569 := tmp3050
_ = let__2569

tmp3066 := Call2(__e, SymFn(symshen_4hds_a_2), let__2569, symwhere)


var ifres3051 Obj

if True == tmp3066 {
tmp3052 := Call1(__e, SymFn(symtail), let__2569)


let__2570 := tmp3052
_ = let__2570

tmp3064 := GConsP(__e, let__2570)

var ifres3053 Obj

if True == tmp3064 {
tmp3054 := Call1(__e, SymFn(symhead), let__2570)


let__2571 := tmp3054
_ = let__2571

tmp3055 := Call1(__e, SymFn(symtail), let__2570)


let__2572 := tmp3055
_ = let__2572

tmp3056 := GCons(__e, let__2568, Nil)

tmp3057 := GCons(__e, symshen_4choicepoint_b, tmp3056)

tmp3058 := GCons(__e, tmp3057, Nil)

tmp3059 := GCons(__e, let__2571, tmp3058)

tmp3060 := GCons(__e, symwhere, tmp3059)

tmp3061 := Call2(__e, SymFn(sym_8p), let__2565, tmp3060)


tmp3062 := Call2(__e, SymFn(symshen_4comb), let__2572, tmp3061)


ifres3053 = tmp3062


} else {
tmp3063 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3053 = tmp3063


}

ifres3051 = ifres3053


} else {
tmp3065 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3051 = tmp3065


}

ifres3048 = ifres3051


} else {
tmp3067 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3048 = tmp3067


}

ifres3046 = ifres3048


} else {
tmp3069 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3046 = tmp3069


}

ifres3042 = ifres3046


}

let__2563 := ifres3042
_ = let__2563

tmp3094 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2563)


if True == tmp3094 {
tmp3072 := Call1(__e, SymFn(symshen_4_5patterns_6), V628)


let__2574 := tmp3072
_ = let__2574

tmp3090 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2574)


var ifres3073 Obj

if True == tmp3090 {
tmp3074 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3073 = tmp3074


} else {
tmp3075 := Call1(__e, SymFn(symshen_4_5_1out), let__2574)


let__2575 := tmp3075
_ = let__2575

tmp3076 := Call1(__e, SymFn(symshen_4in_1_6), let__2574)


let__2576 := tmp3076
_ = let__2576

tmp3089 := Call2(__e, SymFn(symshen_4hds_a_2), let__2576, sym_5_1)


var ifres3077 Obj

if True == tmp3089 {
tmp3078 := Call1(__e, SymFn(symtail), let__2576)


let__2577 := tmp3078
_ = let__2577

tmp3087 := GConsP(__e, let__2577)

var ifres3079 Obj

if True == tmp3087 {
tmp3080 := Call1(__e, SymFn(symhead), let__2577)


let__2578 := tmp3080
_ = let__2578

tmp3081 := Call1(__e, SymFn(symtail), let__2577)


let__2579 := tmp3081
_ = let__2579

tmp3082 := GCons(__e, let__2578, Nil)

tmp3083 := GCons(__e, symshen_4choicepoint_b, tmp3082)

tmp3084 := Call2(__e, SymFn(sym_8p), let__2575, tmp3083)


tmp3085 := Call2(__e, SymFn(symshen_4comb), let__2579, tmp3084)


ifres3079 = tmp3085


} else {
tmp3086 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3079 = tmp3086


}

ifres3077 = ifres3079


} else {
tmp3088 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3077 = tmp3088


}

ifres3073 = ifres3077


}

let__2573 := ifres3073
_ = let__2573

tmp3092 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2573)


if True == tmp3092 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2573
}


} else {
return let__2563
}


} else {
return let__2556
}


} else {
return let__2546
}


panic("unreachable")
})

tmp3099 := Call2(__e, ns2_1set, symshen_4_5rule_6, tmp2994)


_ = tmp3099

tmp3100 := MakeNative1(func(__e *ControlFlow, V663 Obj) Obj {
tmp3101 := Call1(__e, SymFn(symshen_4_5pattern_6), V663)


let__2581 := tmp3101
_ = let__2581

tmp3114 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2581)


var ifres3102 Obj

if True == tmp3114 {
tmp3103 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3102 = tmp3103


} else {
tmp3104 := Call1(__e, SymFn(symshen_4_5_1out), let__2581)


let__2582 := tmp3104
_ = let__2582

tmp3105 := Call1(__e, SymFn(symshen_4in_1_6), let__2581)


let__2583 := tmp3105
_ = let__2583

tmp3106 := Call1(__e, SymFn(symshen_4_5patterns_6), let__2583)


let__2584 := tmp3106
_ = let__2584

tmp3113 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2584)


var ifres3107 Obj

if True == tmp3113 {
tmp3108 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3107 = tmp3108


} else {
tmp3109 := Call1(__e, SymFn(symshen_4_5_1out), let__2584)


let__2585 := tmp3109
_ = let__2585

tmp3110 := Call1(__e, SymFn(symshen_4in_1_6), let__2584)


let__2586 := tmp3110
_ = let__2586

tmp3111 := GCons(__e, let__2582, let__2585)

tmp3112 := Call2(__e, SymFn(symshen_4comb), let__2586, tmp3111)


ifres3107 = tmp3112


}

ifres3102 = ifres3107


}

let__2580 := ifres3102
_ = let__2580

tmp3124 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2580)


if True == tmp3124 {
tmp3115 := Call1(__e, SymFn(sym_5e_6), V663)


let__2588 := tmp3115
_ = let__2588

tmp3120 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2588)


var ifres3116 Obj

if True == tmp3120 {
tmp3117 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3116 = tmp3117


} else {
tmp3118 := Call1(__e, SymFn(symshen_4in_1_6), let__2588)


let__2589 := tmp3118
_ = let__2589

tmp3119 := Call2(__e, SymFn(symshen_4comb), let__2589, Nil)


ifres3116 = tmp3119


}

let__2587 := ifres3116
_ = let__2587

tmp3122 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2587)


if True == tmp3122 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2587
}


} else {
return let__2580
}


panic("unreachable")
})

tmp3125 := Call2(__e, ns2_1set, symshen_4_5patterns_6, tmp3100)


_ = tmp3125

tmp3126 := MakeNative1(func(__e *ControlFlow, V674 Obj) Obj {
tmp3147 := Call1(__e, SymFn(symshen_4ccons_2), V674)


var ifres3127 Obj

if True == tmp3147 {
tmp3128 := Call1(__e, SymFn(symhead), V674)


let__2591 := tmp3128
_ = let__2591

tmp3129 := Call1(__e, SymFn(symtail), V674)


let__2592 := tmp3129
_ = let__2592

tmp3145 := Call2(__e, SymFn(symshen_4hds_a_2), let__2591, symvector)


var ifres3130 Obj

if True == tmp3145 {
tmp3131 := Call1(__e, SymFn(symtail), let__2591)


let__2593 := tmp3131
_ = let__2593

tmp3143 := Call2(__e, SymFn(symshen_4hds_a_2), let__2593, MakeInteger(0))


var ifres3132 Obj

if True == tmp3143 {
tmp3133 := Call1(__e, SymFn(symtail), let__2593)


let__2594 := tmp3133
_ = let__2594

tmp3134 := Call1(__e, SymFn(sym_5end_6), let__2594)


let__2595 := tmp3134
_ = let__2595

tmp3141 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2595)


var ifres3135 Obj

if True == tmp3141 {
tmp3136 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3135 = tmp3136


} else {
tmp3137 := Call1(__e, SymFn(symshen_4in_1_6), let__2595)


let__2596 := tmp3137
_ = let__2596

tmp3138 := GCons(__e, MakeInteger(0), Nil)

tmp3139 := GCons(__e, symvector, tmp3138)

tmp3140 := Call2(__e, SymFn(symshen_4comb), let__2592, tmp3139)


ifres3135 = tmp3140


}

ifres3132 = ifres3135


} else {
tmp3142 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3132 = tmp3142


}

ifres3130 = ifres3132


} else {
tmp3144 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3130 = tmp3144


}

ifres3127 = ifres3130


} else {
tmp3146 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3127 = tmp3146


}

let__2590 := ifres3127
_ = let__2590

tmp3170 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2590)


if True == tmp3170 {
tmp3157 := GConsP(__e, V674)

var ifres3148 Obj

if True == tmp3157 {
tmp3149 := Call1(__e, SymFn(symhead), V674)


let__2598 := tmp3149
_ = let__2598

tmp3150 := Call1(__e, SymFn(symtail), V674)


let__2599 := tmp3150
_ = let__2599

tmp3155 := GConsP(__e, let__2598)

var ifres3151 Obj

if True == tmp3155 {
tmp3152 := Call1(__e, SymFn(symshen_4compound_1pattern), let__2598)


tmp3153 := Call2(__e, SymFn(symshen_4comb), let__2599, tmp3152)


ifres3151 = tmp3153


} else {
tmp3154 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3151 = tmp3154


}

ifres3148 = ifres3151


} else {
tmp3156 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3148 = tmp3156


}

let__2597 := ifres3148
_ = let__2597

tmp3168 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2597)


if True == tmp3168 {
tmp3158 := Call1(__e, SymFn(symshen_4_5simple_1pattern_6), V674)


let__2601 := tmp3158
_ = let__2601

tmp3164 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2601)


var ifres3159 Obj

if True == tmp3164 {
tmp3160 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3159 = tmp3160


} else {
tmp3161 := Call1(__e, SymFn(symshen_4_5_1out), let__2601)


let__2602 := tmp3161
_ = let__2602

tmp3162 := Call1(__e, SymFn(symshen_4in_1_6), let__2601)


let__2603 := tmp3162
_ = let__2603

tmp3163 := Call2(__e, SymFn(symshen_4comb), let__2603, let__2602)


ifres3159 = tmp3163


}

let__2600 := ifres3159
_ = let__2600

tmp3166 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2600)


if True == tmp3166 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2600
}


} else {
return let__2597
}


} else {
return let__2590
}


panic("unreachable")
})

tmp3171 := Call2(__e, ns2_1set, symshen_4_5pattern_6, tmp3126)


_ = tmp3171

tmp3172 := MakeNative1(func(__e *ControlFlow, V689 Obj) Obj {
tmp3181 := GConsP(__e, V689)

var ifres3173 Obj

if True == tmp3181 {
tmp3174 := Call1(__e, SymFn(symhead), V689)


let__2605 := tmp3174
_ = let__2605

tmp3175 := Call1(__e, SymFn(symtail), V689)


let__2606 := tmp3175
_ = let__2606

tmp3179 := Call1(__e, SymFn(symshen_4constructor_2), let__2605)


var ifres3176 Obj

if True == tmp3179 {
tmp3177 := Call2(__e, SymFn(symshen_4comb), let__2606, let__2605)


ifres3176 = tmp3177


} else {
tmp3178 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3176 = tmp3178


}

ifres3173 = ifres3176


} else {
tmp3180 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3173 = tmp3180


}

let__2604 := ifres3173
_ = let__2604

tmp3183 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2604)


if True == tmp3183 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2604
}


panic("unreachable")
})

tmp3184 := Call2(__e, ns2_1set, symshen_4_5constructor_6, tmp3172)


_ = tmp3184

tmp3185 := MakeNative1(func(__e *ControlFlow, V693 Obj) Obj {
tmp3186 := GCons(__e, sym_8v, Nil)

tmp3187 := GCons(__e, sym_8s, tmp3186)

tmp3188 := GCons(__e, sym_8p, tmp3187)

tmp3189 := GCons(__e, symcons, tmp3188)

return TailCall2(__e, SymFn(symelement_2), V693, tmp3189)


panic("unreachable")
})

tmp3190 := Call2(__e, ns2_1set, symshen_4constructor_2, tmp3185)


_ = tmp3190

tmp3191 := MakeNative1(func(__e *ControlFlow, V694 Obj) Obj {
tmp3192 := Call3(__e, SymFn(symshen_4app), V694, MakeString(" is not a legitimate constructor\n"), symshen_4r)


return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp3192)


panic("unreachable")
})

tmp3193 := Call2(__e, ns2_1set, symshen_4constructor_1error, tmp3191)


_ = tmp3193

tmp3194 := MakeNative2(func(__e *ControlFlow, V695, V696 Obj) Obj {
tmp3195 := GValue(__e, symshen_4_dcustom_1pattern_1compiler_d)

let__2607 := tmp3195
_ = let__2607

tmp3198 := GEq(__e, let__2607, False)

if True == tmp3198 {
return TailCall1(__e, SymFn(symthaw), V696)
} else {
tmp3196 := Call1(__e, let__2607, V695)


return TailCall1(__e, tmp3196, V696)


}


panic("unreachable")
})

tmp3199 := Call2(__e, ns2_1set, symshen_4custom_1pattern_1compiler, tmp3194)


_ = tmp3199

tmp3200 := MakeNative1(func(__e *ControlFlow, V698 Obj) Obj {
tmp3201 := GValue(__e, symshen_4_dcustom_1pattern_1reducer_d)

let__2608 := tmp3201
_ = let__2608

tmp3203 := GEq(__e, let__2608, False)

if True == tmp3203 {
return TailCall0(__e, SymFn(symfail))
} else {
return TailCall1(__e, let__2608, V698)
}


panic("unreachable")
})

tmp3204 := Call2(__e, ns2_1set, symshen_4custom_1pattern_1reducer, tmp3200)


_ = tmp3204

tmp3205 := MakeNative1(func(__e *ControlFlow, V700 Obj) Obj {
tmp3206 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, SymFn(symshen_4compound_1pattern_1h), V700)
panic("unreachable")
})

return TailCall2(__e, SymFn(symshen_4custom_1pattern_1compiler), V700, tmp3206)


panic("unreachable")
})

tmp3207 := Call2(__e, ns2_1set, symshen_4compound_1pattern, tmp3205)


_ = tmp3207

tmp3208 := MakeNative1(func(__e *ControlFlow, V701 Obj) Obj {
tmp3235 := GConsP(__e, V701)

var ifres3220 Obj

if True == tmp3235 {
tmp3233 := GTl(__e, V701)

tmp3234 := GConsP(__e, tmp3233)

var ifres3221 Obj

if True == tmp3234 {
tmp3230 := GTl(__e, V701)

tmp3231 := GTl(__e, tmp3230)

tmp3232 := GConsP(__e, tmp3231)

var ifres3222 Obj

if True == tmp3232 {
tmp3226 := GTl(__e, V701)

tmp3227 := GTl(__e, tmp3226)

tmp3228 := GTl(__e, tmp3227)

tmp3229 := GEq(__e, Nil, tmp3228)

var ifres3223 Obj

if True == tmp3229 {
tmp3224 := GHd(__e, V701)

tmp3225 := Call1(__e, SymFn(symshen_4constructor_2), tmp3224)


ifres3223 = tmp3225


} else {
ifres3223 = False


}

ifres3222 = ifres3223


} else {
ifres3222 = False


}

ifres3221 = ifres3222


} else {
ifres3221 = False


}

ifres3220 = ifres3221


} else {
ifres3220 = False


}

if True == ifres3220 {
tmp3209 := GHd(__e, V701)

tmp3210 := GTl(__e, V701)

tmp3211 := GHd(__e, tmp3210)

tmp3212 := Call1(__e, SymFn(symshen_4compile_1pattern_1fragment), tmp3211)


tmp3213 := GTl(__e, V701)

tmp3214 := GTl(__e, tmp3213)

tmp3215 := GHd(__e, tmp3214)

tmp3216 := Call1(__e, SymFn(symshen_4compile_1pattern_1fragment), tmp3215)


tmp3217 := GCons(__e, tmp3216, Nil)

tmp3218 := GCons(__e, tmp3212, tmp3217)

return GCons(__e, tmp3209, tmp3218)


} else {
return TailCall1(__e, SymFn(symshen_4constructor_1error), V701)
}


panic("unreachable")
})

tmp3236 := Call2(__e, ns2_1set, symshen_4compound_1pattern_1h, tmp3208)


_ = tmp3236

tmp3237 := MakeNative1(func(__e *ControlFlow, V702 Obj) Obj {
tmp3262 := GConsP(__e, V702)

var ifres3248 Obj

if True == tmp3262 {
tmp3260 := GHd(__e, V702)

tmp3261 := GEq(__e, symvector, tmp3260)

var ifres3249 Obj

if True == tmp3261 {
tmp3258 := GTl(__e, V702)

tmp3259 := GConsP(__e, tmp3258)

var ifres3250 Obj

if True == tmp3259 {
tmp3255 := GTl(__e, V702)

tmp3256 := GHd(__e, tmp3255)

tmp3257 := GEq(__e, MakeInteger(0), tmp3256)

var ifres3251 Obj

if True == tmp3257 {
tmp3252 := GTl(__e, V702)

tmp3253 := GTl(__e, tmp3252)

tmp3254 := GEq(__e, Nil, tmp3253)

ifres3251 = tmp3254


} else {
ifres3251 = False


}

ifres3250 = ifres3251


} else {
ifres3250 = False


}

ifres3249 = ifres3250


} else {
ifres3249 = False


}

ifres3248 = ifres3249


} else {
ifres3248 = False


}

if True == ifres3248 {
return V702
} else {
tmp3246 := GConsP(__e, V702)

if True == tmp3246 {
return TailCall1(__e, SymFn(symshen_4compound_1pattern), V702)
} else {
tmp3244 := GEq(__e, V702, sym__)

if True == tmp3244 {
return TailCall1(__e, SymFn(symgensym), symY)
} else {
tmp3239 := GCons(__e, sym_5_1, Nil)

tmp3240 := GCons(__e, sym_1_6, tmp3239)

tmp3241 := Call2(__e, SymFn(symelement_2), V702, tmp3240)


if True == GNot(__e, tmp3241) {
return V702
} else {
return TailCall1(__e, SymFn(symshen_4constructor_1error), V702)
}


}


}


}


panic("unreachable")
})

tmp3263 := Call2(__e, ns2_1set, symshen_4compile_1pattern_1fragment, tmp3237)


_ = tmp3263

tmp3264 := MakeNative1(func(__e *ControlFlow, V707 Obj) Obj {
tmp3285 := GConsP(__e, V707)

var ifres3266 Obj

if True == tmp3285 {
tmp3283 := GHd(__e, V707)

tmp3284 := GEq(__e, sym_8p, tmp3283)

var ifres3267 Obj

if True == tmp3284 {
tmp3281 := GTl(__e, V707)

tmp3282 := GConsP(__e, tmp3281)

var ifres3268 Obj

if True == tmp3282 {
tmp3278 := GTl(__e, V707)

tmp3279 := GHd(__e, tmp3278)

tmp3280 := GEq(__e, symshen_4custom_1pattern, tmp3279)

var ifres3269 Obj

if True == tmp3280 {
tmp3275 := GTl(__e, V707)

tmp3276 := GTl(__e, tmp3275)

tmp3277 := GConsP(__e, tmp3276)

var ifres3270 Obj

if True == tmp3277 {
tmp3271 := GTl(__e, V707)

tmp3272 := GTl(__e, tmp3271)

tmp3273 := GTl(__e, tmp3272)

tmp3274 := GEq(__e, Nil, tmp3273)

ifres3270 = tmp3274


} else {
ifres3270 = False


}

ifres3269 = ifres3270


} else {
ifres3269 = False


}

ifres3268 = ifres3269


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

if True == ifres3266 {
return True
} else {
return False
}


panic("unreachable")
})

tmp3286 := Call2(__e, ns2_1set, symshen_4custom_1pattern_2, tmp3264)


_ = tmp3286

tmp3287 := MakeNative1(func(__e *ControlFlow, V710 Obj) Obj {
tmp3310 := GConsP(__e, V710)

var ifres3291 Obj

if True == tmp3310 {
tmp3308 := GHd(__e, V710)

tmp3309 := GEq(__e, sym_8p, tmp3308)

var ifres3292 Obj

if True == tmp3309 {
tmp3306 := GTl(__e, V710)

tmp3307 := GConsP(__e, tmp3306)

var ifres3293 Obj

if True == tmp3307 {
tmp3303 := GTl(__e, V710)

tmp3304 := GHd(__e, tmp3303)

tmp3305 := GEq(__e, symshen_4custom_1pattern, tmp3304)

var ifres3294 Obj

if True == tmp3305 {
tmp3300 := GTl(__e, V710)

tmp3301 := GTl(__e, tmp3300)

tmp3302 := GConsP(__e, tmp3301)

var ifres3295 Obj

if True == tmp3302 {
tmp3296 := GTl(__e, V710)

tmp3297 := GTl(__e, tmp3296)

tmp3298 := GTl(__e, tmp3297)

tmp3299 := GEq(__e, Nil, tmp3298)

ifres3295 = tmp3299


} else {
ifres3295 = False


}

ifres3294 = ifres3295


} else {
ifres3294 = False


}

ifres3293 = ifres3294


} else {
ifres3293 = False


}

ifres3292 = ifres3293


} else {
ifres3292 = False


}

ifres3291 = ifres3292


} else {
ifres3291 = False


}

if True == ifres3291 {
tmp3288 := GTl(__e, V710)

tmp3289 := GTl(__e, tmp3288)

return GHd(__e, tmp3289)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.custom-pattern-body"))
}


panic("unreachable")
})

tmp3311 := Call2(__e, ns2_1set, symshen_4custom_1pattern_1body, tmp3287)


_ = tmp3311

tmp3312 := MakeNative1(func(__e *ControlFlow, V711 Obj) Obj {
tmp3322 := GConsP(__e, V711)

var ifres3313 Obj

if True == tmp3322 {
tmp3314 := Call1(__e, SymFn(symhead), V711)


let__2610 := tmp3314
_ = let__2610

tmp3315 := Call1(__e, SymFn(symtail), V711)


let__2611 := tmp3315
_ = let__2611

tmp3320 := GEq(__e, let__2610, sym__)

var ifres3316 Obj

if True == tmp3320 {
tmp3317 := Call1(__e, SymFn(symgensym), symY)


tmp3318 := Call2(__e, SymFn(symshen_4comb), let__2611, tmp3317)


ifres3316 = tmp3318


} else {
tmp3319 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3316 = tmp3319


}

ifres3313 = ifres3316


} else {
tmp3321 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3313 = tmp3321


}

let__2609 := ifres3313
_ = let__2609

tmp3338 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2609)


if True == tmp3338 {
tmp3334 := GConsP(__e, V711)

var ifres3323 Obj

if True == tmp3334 {
tmp3324 := Call1(__e, SymFn(symhead), V711)


let__2613 := tmp3324
_ = let__2613

tmp3325 := Call1(__e, SymFn(symtail), V711)


let__2614 := tmp3325
_ = let__2614

tmp3329 := GCons(__e, sym_5_1, Nil)

tmp3330 := GCons(__e, sym_1_6, tmp3329)

tmp3331 := Call2(__e, SymFn(symelement_2), let__2613, tmp3330)


tmp3332 := GNot(__e, tmp3331)

var ifres3326 Obj

if True == tmp3332 {
tmp3327 := Call2(__e, SymFn(symshen_4comb), let__2614, let__2613)


ifres3326 = tmp3327


} else {
tmp3328 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3326 = tmp3328


}

ifres3323 = ifres3326


} else {
tmp3333 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3323 = tmp3333


}

let__2612 := ifres3323
_ = let__2612

tmp3336 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2612)


if True == tmp3336 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2612
}


} else {
return let__2609
}


panic("unreachable")
})

tmp3339 := Call2(__e, ns2_1set, symshen_4_5simple_1pattern_6, tmp3312)


_ = tmp3339

tmp3340 := MakeNative1(func(__e *ControlFlow, V718 Obj) Obj {
tmp3341 := Call1(__e, SymFn(symshen_4_5pattern_6), V718)


let__2616 := tmp3341
_ = let__2616

tmp3347 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2616)


var ifres3342 Obj

if True == tmp3347 {
tmp3343 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3342 = tmp3343


} else {
tmp3344 := Call1(__e, SymFn(symshen_4_5_1out), let__2616)


let__2617 := tmp3344
_ = let__2617

tmp3345 := Call1(__e, SymFn(symshen_4in_1_6), let__2616)


let__2618 := tmp3345
_ = let__2618

tmp3346 := Call2(__e, SymFn(symshen_4comb), let__2618, let__2617)


ifres3342 = tmp3346


}

let__2615 := ifres3342
_ = let__2615

tmp3349 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2615)


if True == tmp3349 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2615
}


panic("unreachable")
})

tmp3350 := Call2(__e, ns2_1set, symshen_4_5pattern1_6, tmp3340)


_ = tmp3350

tmp3351 := MakeNative1(func(__e *ControlFlow, V723 Obj) Obj {
tmp3352 := Call1(__e, SymFn(symshen_4_5pattern_6), V723)


let__2620 := tmp3352
_ = let__2620

tmp3358 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2620)


var ifres3353 Obj

if True == tmp3358 {
tmp3354 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres3353 = tmp3354


} else {
tmp3355 := Call1(__e, SymFn(symshen_4_5_1out), let__2620)


let__2621 := tmp3355
_ = let__2621

tmp3356 := Call1(__e, SymFn(symshen_4in_1_6), let__2620)


let__2622 := tmp3356
_ = let__2622

tmp3357 := Call2(__e, SymFn(symshen_4comb), let__2622, let__2621)


ifres3353 = tmp3357


}

let__2619 := ifres3353
_ = let__2619

tmp3360 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__2619)


if True == tmp3360 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__2619
}


panic("unreachable")
})

tmp3361 := Call2(__e, ns2_1set, symshen_4_5pattern2_6, tmp3351)


_ = tmp3361

tmp3362 := MakeNative1(func(__e *ControlFlow, V728 Obj) Obj {
tmp3363 := PrimGuard1(__e, symabsvector, PrimAbsvector, MakeInteger(2))

let__2623 := tmp3363
_ = let__2623

tmp3364 := GVecSet(__e, let__2623, MakeInteger(0), symshen_4printF)

let__2624 := tmp3364
_ = let__2624

tmp3365 := PrimGuard1(__e, symstr, PrimStr, V728)

tmp3366 := Call2(__e, SymFn(sym_8s), tmp3365, MakeString(")"))


tmp3367 := Call2(__e, SymFn(sym_8s), MakeString(" "), tmp3366)


tmp3368 := Call2(__e, SymFn(sym_8s), MakeString("n"), tmp3367)


tmp3369 := Call2(__e, SymFn(sym_8s), MakeString("f"), tmp3368)


tmp3370 := Call2(__e, SymFn(sym_8s), MakeString("("), tmp3369)


tmp3371 := GVecSet(__e, let__2624, MakeInteger(1), tmp3370)

let__2625 := tmp3371
_ = let__2625

return let__2625


panic("unreachable")
})

tmp3372 := Call2(__e, ns2_1set, symshen_4fn_1print, tmp3362)


_ = tmp3372

tmp3373 := MakeNative1(func(__e *ControlFlow, V732 Obj) Obj {
return GVecGet(__e, V732, MakeInteger(1))
panic("unreachable")
})

tmp3374 := Call2(__e, ns2_1set, symshen_4printF, tmp3373)


_ = tmp3374

tmp3375 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V737 := __self1
_ = V737
V738 := __self2
_ = V738
tmp3396 := GConsP(__e, V738)

var ifres3393 Obj

if True == tmp3396 {
tmp3394 := GTl(__e, V738)

tmp3395 := GEq(__e, Nil, tmp3394)

ifres3393 = tmp3395


} else {
ifres3393 = False


}

if True == ifres3393 {
tmp3376 := GHd(__e, V738)

return TailCall1(__e, SymFn(symlength), tmp3376)


} else {
tmp3391 := GConsP(__e, V738)

var ifres3381 Obj

if True == tmp3391 {
tmp3389 := GTl(__e, V738)

tmp3390 := GConsP(__e, tmp3389)

var ifres3382 Obj

if True == tmp3390 {
tmp3383 := GHd(__e, V738)

tmp3384 := Call1(__e, SymFn(symlength), tmp3383)


tmp3385 := GTl(__e, V738)

tmp3386 := GHd(__e, tmp3385)

tmp3387 := Call1(__e, SymFn(symlength), tmp3386)


tmp3388 := GEq(__e, tmp3384, tmp3387)

ifres3382 = tmp3388


} else {
ifres3382 = False


}

ifres3381 = ifres3382


} else {
ifres3381 = False


}

if True == ifres3381 {
tmp3377 := GTl(__e, V738)

if SymFn(symshen_4arity_1chk) == __self {
__self1, __self2 = V737, tmp3377
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4arity_1chk), V737, tmp3377)


} else {
tmp3378 := Call3(__e, SymFn(symshen_4app), V737, MakeString("\n"), symshen_4a)


return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("arity error in "), tmp3378))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("arity error in "), tmp3378)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()


}


}


panic("unreachable")
})
return __self
})()

tmp3397 := Call2(__e, ns2_1set, symshen_4arity_1chk, tmp3375)


_ = tmp3397

tmp3398 := MakeNative2(func(__e *ControlFlow, V739, V740 Obj) Obj {
tmp3404 := Call1(__e, SymFn(symtuple_2), V740)


if True == tmp3404 {
tmp3399 := Call1(__e, SymFn(symfst), V740)


tmp3400 := Call1(__e, SymFn(symshen_4extract_1vars), tmp3399)


tmp3401 := Call1(__e, SymFn(symsnd), V740)


tmp3402 := Call2(__e, SymFn(symshen_4find_1free_1vars), tmp3400, tmp3401)


return TailCall2(__e, SymFn(symshen_4free_1variable_1error_1message), V739, tmp3402)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4free_1var_1chk)
}


panic("unreachable")
})

tmp3405 := Call2(__e, ns2_1set, symshen_4free_1var_1chk, tmp3398)


_ = tmp3405

tmp3406 := MakeNative2(func(__e *ControlFlow, V741, V742 Obj) Obj {
tmp3418 := Call1(__e, SymFn(symempty_2), V742)


if True == tmp3418 {
return symshen_4skip
} else {
tmp3407 := Call3(__e, SymFn(symshen_4app), V741, MakeString(":"), symshen_4a)


tmp3408 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("free variables in "), tmp3407)

tmp3409 := Call0(__e, SymFn(symstoutput))


tmp3410 := Call2(__e, SymFn(sympr), tmp3408, tmp3409)


_ = tmp3410

tmp3411 := MakeNative1(func(__e *ControlFlow, Z743 Obj) Obj {
tmp3412 := Call3(__e, SymFn(symshen_4app), Z743, MakeString(""), symshen_4a)


tmp3413 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString(" "), tmp3412)

tmp3414 := Call0(__e, SymFn(symstoutput))


return TailCall2(__e, SymFn(sympr), tmp3413, tmp3414)


panic("unreachable")
})

tmp3415 := Call2(__e, SymFn(symshen_4for_1each), tmp3411, V742)


_ = tmp3415

tmp3416 := Call1(__e, SymFn(symnl), MakeInteger(1))


_ = tmp3416

return TailCall0(__e, SymFn(symabort))


}


panic("unreachable")
})

tmp3419 := Call2(__e, ns2_1set, symshen_4free_1variable_1error_1message, tmp3406)


_ = tmp3419

tmp3420 := MakeNative1(func(__e *ControlFlow, V746 Obj) Obj {
tmp3428 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V746)

if True == tmp3428 {
return GCons(__e, V746, Nil)
} else {
tmp3426 := GConsP(__e, V746)

if True == tmp3426 {
tmp3421 := GHd(__e, V746)

tmp3422 := Call1(__e, SymFn(symshen_4extract_1vars), tmp3421)


tmp3423 := GTl(__e, V746)

tmp3424 := Call1(__e, SymFn(symshen_4extract_1vars), tmp3423)


return TailCall2(__e, SymFn(symunion), tmp3422, tmp3424)


} else {
return Nil
}


}


panic("unreachable")
})

tmp3429 := Call2(__e, ns2_1set, symshen_4extract_1vars, tmp3420)


_ = tmp3429

tmp3430 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V751 := __self1
_ = V751
V752 := __self2
_ = V752
tmp3508 := GConsP(__e, V752)

var ifres3498 Obj

if True == tmp3508 {
tmp3506 := GHd(__e, V752)

tmp3507 := GEq(__e, symprotect, tmp3506)

var ifres3499 Obj

if True == tmp3507 {
tmp3504 := GTl(__e, V752)

tmp3505 := GConsP(__e, tmp3504)

var ifres3500 Obj

if True == tmp3505 {
tmp3501 := GTl(__e, V752)

tmp3502 := GTl(__e, tmp3501)

tmp3503 := GEq(__e, Nil, tmp3502)

ifres3500 = tmp3503


} else {
ifres3500 = False


}

ifres3499 = ifres3500


} else {
ifres3499 = False


}

ifres3498 = ifres3499


} else {
ifres3498 = False


}

if True == ifres3498 {
return Nil
} else {
tmp3496 := GConsP(__e, V752)

var ifres3475 Obj

if True == tmp3496 {
tmp3494 := GHd(__e, V752)

tmp3495 := GEq(__e, symlet, tmp3494)

var ifres3476 Obj

if True == tmp3495 {
tmp3492 := GTl(__e, V752)

tmp3493 := GConsP(__e, tmp3492)

var ifres3477 Obj

if True == tmp3493 {
tmp3489 := GTl(__e, V752)

tmp3490 := GTl(__e, tmp3489)

tmp3491 := GConsP(__e, tmp3490)

var ifres3478 Obj

if True == tmp3491 {
tmp3485 := GTl(__e, V752)

tmp3486 := GTl(__e, tmp3485)

tmp3487 := GTl(__e, tmp3486)

tmp3488 := GConsP(__e, tmp3487)

var ifres3479 Obj

if True == tmp3488 {
tmp3480 := GTl(__e, V752)

tmp3481 := GTl(__e, tmp3480)

tmp3482 := GTl(__e, tmp3481)

tmp3483 := GTl(__e, tmp3482)

tmp3484 := GEq(__e, Nil, tmp3483)

ifres3479 = tmp3484


} else {
ifres3479 = False


}

ifres3478 = ifres3479


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
tmp3431 := GTl(__e, V752)

tmp3432 := GTl(__e, tmp3431)

tmp3433 := GHd(__e, tmp3432)

tmp3434 := Call2(__e, SymFn(symshen_4find_1free_1vars), V751, tmp3433)


tmp3435 := GTl(__e, V752)

tmp3436 := GHd(__e, tmp3435)

tmp3437 := GCons(__e, tmp3436, V751)

tmp3438 := GTl(__e, V752)

tmp3439 := GTl(__e, tmp3438)

tmp3440 := GTl(__e, tmp3439)

tmp3441 := GHd(__e, tmp3440)

tmp3442 := Call2(__e, SymFn(symshen_4find_1free_1vars), tmp3437, tmp3441)


return TailCall2(__e, SymFn(symunion), tmp3434, tmp3442)


} else {
tmp3473 := GConsP(__e, V752)

var ifres3458 Obj

if True == tmp3473 {
tmp3471 := GHd(__e, V752)

tmp3472 := GEq(__e, symlambda, tmp3471)

var ifres3459 Obj

if True == tmp3472 {
tmp3469 := GTl(__e, V752)

tmp3470 := GConsP(__e, tmp3469)

var ifres3460 Obj

if True == tmp3470 {
tmp3466 := GTl(__e, V752)

tmp3467 := GTl(__e, tmp3466)

tmp3468 := GConsP(__e, tmp3467)

var ifres3461 Obj

if True == tmp3468 {
tmp3462 := GTl(__e, V752)

tmp3463 := GTl(__e, tmp3462)

tmp3464 := GTl(__e, tmp3463)

tmp3465 := GEq(__e, Nil, tmp3464)

ifres3461 = tmp3465


} else {
ifres3461 = False


}

ifres3460 = ifres3461


} else {
ifres3460 = False


}

ifres3459 = ifres3460


} else {
ifres3459 = False


}

ifres3458 = ifres3459


} else {
ifres3458 = False


}

if True == ifres3458 {
tmp3443 := GTl(__e, V752)

tmp3444 := GHd(__e, tmp3443)

tmp3445 := GCons(__e, tmp3444, V751)

tmp3446 := GTl(__e, V752)

tmp3447 := GTl(__e, tmp3446)

tmp3448 := GHd(__e, tmp3447)

if SymFn(symshen_4find_1free_1vars) == __self {
__self1, __self2 = tmp3445, tmp3448
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4find_1free_1vars), tmp3445, tmp3448)


} else {
tmp3456 := GConsP(__e, V752)

if True == tmp3456 {
tmp3449 := GHd(__e, V752)

tmp3450 := Call2(__e, SymFn(symshen_4find_1free_1vars), V751, tmp3449)


tmp3451 := GTl(__e, V752)

tmp3452 := Call2(__e, SymFn(symshen_4find_1free_1vars), V751, tmp3451)


return TailCall2(__e, SymFn(symunion), tmp3450, tmp3452)


} else {
tmp3454 := Call2(__e, SymFn(symshen_4free_1variable_2), V752, V751)


if True == tmp3454 {
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

tmp3509 := Call2(__e, ns2_1set, symshen_4find_1free_1vars, tmp3430)


_ = tmp3509

tmp3510 := MakeNative2(func(__e *ControlFlow, V753, V754 Obj) Obj {
tmp3513 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V753)

if True == tmp3513 {
tmp3511 := Call2(__e, SymFn(symelement_2), V753, V754)


return GNot(__e, tmp3511)


} else {
return False
}


panic("unreachable")
})

tmp3514 := Call2(__e, ns2_1set, symshen_4free_1variable_2, tmp3510)


_ = tmp3514

tmp3515 := MakeNative2(func(__e *ControlFlow, V755, V756 Obj) Obj {
tmp3516 := GValue(__e, symshen_4_duserdefs_d)

tmp3517 := Call2(__e, SymFn(symadjoin), V755, tmp3516)


tmp3518 := PrimGuard2(__e, symset, PrimSet, symshen_4_duserdefs_d, tmp3517)

_ = tmp3518

tmp3519 := GValue(__e, sym_dproperty_1vector_d)

return TailCall4(__e, SymFn(symput), V755, symshen_4source, V756, tmp3519)


panic("unreachable")
})

tmp3520 := Call2(__e, ns2_1set, symshen_4record_1kl, tmp3515)


_ = tmp3520

tmp3521 := MakeNative3(func(__e *ControlFlow, V757, V758, V759 Obj) Obj {
tmp3522 := Call1(__e, SymFn(symshen_4parameters), V759)


let__2626 := tmp3522
_ = let__2626

tmp3523 := Call2(__e, SymFn(symshen_4kl_1body), V758, let__2626)


tmp3524 := Call2(__e, SymFn(symshen_4scan_1body), V757, tmp3523)


let__2627 := tmp3524
_ = let__2627

tmp3525 := Call1(__e, SymFn(symshen_4cond_1form), let__2627)


tmp3526 := GCons(__e, tmp3525, Nil)

tmp3527 := GCons(__e, let__2626, tmp3526)

tmp3528 := GCons(__e, V757, tmp3527)

tmp3529 := GCons(__e, symdefun, tmp3528)

let__2628 := tmp3529
_ = let__2628

return let__2628


panic("unreachable")
})

tmp3530 := Call2(__e, ns2_1set, symshen_4compile_1to_1kl, tmp3521)


_ = tmp3530

tmp3531 := MakeNative1(func(__e *ControlFlow, V763 Obj) Obj {
tmp3536 := GEq(__e, MakeInteger(0), V763)

if True == tmp3536 {
return Nil
} else {
tmp3532 := Call1(__e, SymFn(symgensym), symV)


tmp3534 := Call1(__e, SymFn(symshen_4parameters), GSub(__e, V763, MakeInteger(1)))


return GCons(__e, tmp3532, tmp3534)


}


panic("unreachable")
})

tmp3537 := Call2(__e, ns2_1set, symshen_4parameters, tmp3531)


_ = tmp3537

tmp3538 := MakeNative1(func(__e *ControlFlow, V766 Obj) Obj {
tmp3558 := GConsP(__e, V766)

var ifres3542 Obj

if True == tmp3558 {
tmp3556 := GHd(__e, V766)

tmp3557 := GConsP(__e, tmp3556)

var ifres3543 Obj

if True == tmp3557 {
tmp3553 := GHd(__e, V766)

tmp3554 := GHd(__e, tmp3553)

tmp3555 := GEq(__e, True, tmp3554)

var ifres3544 Obj

if True == tmp3555 {
tmp3550 := GHd(__e, V766)

tmp3551 := GTl(__e, tmp3550)

tmp3552 := GConsP(__e, tmp3551)

var ifres3545 Obj

if True == tmp3552 {
tmp3546 := GHd(__e, V766)

tmp3547 := GTl(__e, tmp3546)

tmp3548 := GTl(__e, tmp3547)

tmp3549 := GEq(__e, Nil, tmp3548)

ifres3545 = tmp3549


} else {
ifres3545 = False


}

ifres3544 = ifres3545


} else {
ifres3544 = False


}

ifres3543 = ifres3544


} else {
ifres3543 = False


}

ifres3542 = ifres3543


} else {
ifres3542 = False


}

if True == ifres3542 {
tmp3539 := GHd(__e, V766)

tmp3540 := GTl(__e, tmp3539)

return GHd(__e, tmp3540)


} else {
return GCons(__e, symcond, V766)
}


panic("unreachable")
})

tmp3559 := Call2(__e, ns2_1set, symshen_4cond_1form, tmp3538)


_ = tmp3559

tmp3560 := MakeNative2(func(__e *ControlFlow, V775, V776 Obj) Obj {
tmp3599 := GEq(__e, Nil, V776)

if True == tmp3599 {
tmp3561 := GCons(__e, V775, Nil)

tmp3562 := GCons(__e, symshen_4f_1error, tmp3561)

tmp3563 := GCons(__e, tmp3562, Nil)

tmp3564 := GCons(__e, True, tmp3563)

return GCons(__e, tmp3564, Nil)


} else {
tmp3597 := GConsP(__e, V776)

var ifres3594 Obj

if True == tmp3597 {
tmp3595 := GHd(__e, V776)

tmp3596 := Call1(__e, SymFn(symshen_4choicepoint_2), tmp3595)


ifres3594 = tmp3596


} else {
ifres3594 = False


}

if True == ifres3594 {
tmp3565 := Call1(__e, SymFn(symgensym), symFreeze)


tmp3566 := Call1(__e, SymFn(symgensym), symResult)


tmp3567 := GHd(__e, V776)

tmp3568 := GTl(__e, V776)

return TailCall5(__e, SymFn(symshen_4choicepoint), V775, tmp3565, tmp3566, tmp3567, tmp3568)


} else {
tmp3592 := GConsP(__e, V776)

var ifres3576 Obj

if True == tmp3592 {
tmp3590 := GHd(__e, V776)

tmp3591 := GConsP(__e, tmp3590)

var ifres3577 Obj

if True == tmp3591 {
tmp3587 := GHd(__e, V776)

tmp3588 := GHd(__e, tmp3587)

tmp3589 := GEq(__e, True, tmp3588)

var ifres3578 Obj

if True == tmp3589 {
tmp3584 := GHd(__e, V776)

tmp3585 := GTl(__e, tmp3584)

tmp3586 := GConsP(__e, tmp3585)

var ifres3579 Obj

if True == tmp3586 {
tmp3580 := GHd(__e, V776)

tmp3581 := GTl(__e, tmp3580)

tmp3582 := GTl(__e, tmp3581)

tmp3583 := GEq(__e, Nil, tmp3582)

ifres3579 = tmp3583


} else {
ifres3579 = False


}

ifres3578 = ifres3579


} else {
ifres3578 = False


}

ifres3577 = ifres3578


} else {
ifres3577 = False


}

ifres3576 = ifres3577


} else {
ifres3576 = False


}

if True == ifres3576 {
tmp3569 := GHd(__e, V776)

return GCons(__e, tmp3569, Nil)


} else {
tmp3574 := GConsP(__e, V776)

if True == tmp3574 {
tmp3570 := GHd(__e, V776)

tmp3571 := GTl(__e, V776)

tmp3572 := Call2(__e, SymFn(symshen_4scan_1body), V775, tmp3571)


return GCons(__e, tmp3570, tmp3572)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.scan-body"))
}


}


}


}


panic("unreachable")
})

tmp3600 := Call2(__e, ns2_1set, symshen_4scan_1body, tmp3560)


_ = tmp3600

tmp3601 := MakeNative1(func(__e *ControlFlow, V783 Obj) Obj {
tmp3630 := GConsP(__e, V783)

var ifres3603 Obj

if True == tmp3630 {
tmp3628 := GTl(__e, V783)

tmp3629 := GConsP(__e, tmp3628)

var ifres3604 Obj

if True == tmp3629 {
tmp3625 := GTl(__e, V783)

tmp3626 := GHd(__e, tmp3625)

tmp3627 := GConsP(__e, tmp3626)

var ifres3605 Obj

if True == tmp3627 {
tmp3621 := GTl(__e, V783)

tmp3622 := GHd(__e, tmp3621)

tmp3623 := GHd(__e, tmp3622)

tmp3624 := GEq(__e, symshen_4choicepoint_b, tmp3623)

var ifres3606 Obj

if True == tmp3624 {
tmp3617 := GTl(__e, V783)

tmp3618 := GHd(__e, tmp3617)

tmp3619 := GTl(__e, tmp3618)

tmp3620 := GConsP(__e, tmp3619)

var ifres3607 Obj

if True == tmp3620 {
tmp3612 := GTl(__e, V783)

tmp3613 := GHd(__e, tmp3612)

tmp3614 := GTl(__e, tmp3613)

tmp3615 := GTl(__e, tmp3614)

tmp3616 := GEq(__e, Nil, tmp3615)

var ifres3608 Obj

if True == tmp3616 {
tmp3609 := GTl(__e, V783)

tmp3610 := GTl(__e, tmp3609)

tmp3611 := GEq(__e, Nil, tmp3610)

ifres3608 = tmp3611


} else {
ifres3608 = False


}

ifres3607 = ifres3608


} else {
ifres3607 = False


}

ifres3606 = ifres3607


} else {
ifres3606 = False


}

ifres3605 = ifres3606


} else {
ifres3605 = False


}

ifres3604 = ifres3605


} else {
ifres3604 = False


}

ifres3603 = ifres3604


} else {
ifres3603 = False


}

if True == ifres3603 {
return True
} else {
return False
}


panic("unreachable")
})

tmp3631 := Call2(__e, ns2_1set, symshen_4choicepoint_2, tmp3601)


_ = tmp3631

tmp3632 := MakeNative5(func(__e *ControlFlow, V799, V800, V801, V802, V803 Obj) Obj {
tmp3808 := GConsP(__e, V802)

var ifres3741 Obj

if True == tmp3808 {
tmp3806 := GTl(__e, V802)

tmp3807 := GConsP(__e, tmp3806)

var ifres3742 Obj

if True == tmp3807 {
tmp3803 := GTl(__e, V802)

tmp3804 := GHd(__e, tmp3803)

tmp3805 := GConsP(__e, tmp3804)

var ifres3743 Obj

if True == tmp3805 {
tmp3799 := GTl(__e, V802)

tmp3800 := GHd(__e, tmp3799)

tmp3801 := GTl(__e, tmp3800)

tmp3802 := GConsP(__e, tmp3801)

var ifres3744 Obj

if True == tmp3802 {
tmp3794 := GTl(__e, V802)

tmp3795 := GHd(__e, tmp3794)

tmp3796 := GTl(__e, tmp3795)

tmp3797 := GHd(__e, tmp3796)

tmp3798 := GConsP(__e, tmp3797)

var ifres3745 Obj

if True == tmp3798 {
tmp3788 := GTl(__e, V802)

tmp3789 := GHd(__e, tmp3788)

tmp3790 := GTl(__e, tmp3789)

tmp3791 := GHd(__e, tmp3790)

tmp3792 := GHd(__e, tmp3791)

tmp3793 := GEq(__e, symfail_1if, tmp3792)

var ifres3746 Obj

if True == tmp3793 {
tmp3782 := GTl(__e, V802)

tmp3783 := GHd(__e, tmp3782)

tmp3784 := GTl(__e, tmp3783)

tmp3785 := GHd(__e, tmp3784)

tmp3786 := GTl(__e, tmp3785)

tmp3787 := GConsP(__e, tmp3786)

var ifres3747 Obj

if True == tmp3787 {
tmp3775 := GTl(__e, V802)

tmp3776 := GHd(__e, tmp3775)

tmp3777 := GTl(__e, tmp3776)

tmp3778 := GHd(__e, tmp3777)

tmp3779 := GTl(__e, tmp3778)

tmp3780 := GTl(__e, tmp3779)

tmp3781 := GConsP(__e, tmp3780)

var ifres3748 Obj

if True == tmp3781 {
tmp3767 := GTl(__e, V802)

tmp3768 := GHd(__e, tmp3767)

tmp3769 := GTl(__e, tmp3768)

tmp3770 := GHd(__e, tmp3769)

tmp3771 := GTl(__e, tmp3770)

tmp3772 := GTl(__e, tmp3771)

tmp3773 := GTl(__e, tmp3772)

tmp3774 := GEq(__e, Nil, tmp3773)

var ifres3749 Obj

if True == tmp3774 {
tmp3762 := GTl(__e, V802)

tmp3763 := GHd(__e, tmp3762)

tmp3764 := GTl(__e, tmp3763)

tmp3765 := GTl(__e, tmp3764)

tmp3766 := GEq(__e, Nil, tmp3765)

var ifres3750 Obj

if True == tmp3766 {
tmp3759 := GTl(__e, V802)

tmp3760 := GTl(__e, tmp3759)

tmp3761 := GEq(__e, Nil, tmp3760)

var ifres3751 Obj

if True == tmp3761 {
tmp3752 := GTl(__e, V802)

tmp3753 := GHd(__e, tmp3752)

tmp3754 := GTl(__e, tmp3753)

tmp3755 := GHd(__e, tmp3754)

tmp3756 := GTl(__e, tmp3755)

tmp3757 := GHd(__e, tmp3756)

tmp3758 := GEq(__e, V799, tmp3757)

ifres3751 = tmp3758


} else {
ifres3751 = False


}

ifres3750 = ifres3751


} else {
ifres3750 = False


}

ifres3749 = ifres3750


} else {
ifres3749 = False


}

ifres3748 = ifres3749


} else {
ifres3748 = False


}

ifres3747 = ifres3748


} else {
ifres3747 = False


}

ifres3746 = ifres3747


} else {
ifres3746 = False


}

ifres3745 = ifres3746


} else {
ifres3745 = False


}

ifres3744 = ifres3745


} else {
ifres3744 = False


}

ifres3743 = ifres3744


} else {
ifres3743 = False


}

ifres3742 = ifres3743


} else {
ifres3742 = False


}

ifres3741 = ifres3742


} else {
ifres3741 = False


}

if True == ifres3741 {
tmp3633 := GTl(__e, V802)

tmp3634 := GHd(__e, tmp3633)

tmp3635 := GTl(__e, tmp3634)

tmp3636 := GHd(__e, tmp3635)

tmp3637 := GTl(__e, tmp3636)

tmp3638 := GHd(__e, tmp3637)

tmp3639 := Call2(__e, SymFn(symshen_4scan_1body), tmp3638, V803)


tmp3640 := GCons(__e, symcond, tmp3639)

tmp3641 := GCons(__e, tmp3640, Nil)

tmp3642 := GCons(__e, symfreeze, tmp3641)

tmp3643 := GHd(__e, V802)

tmp3644 := GTl(__e, V802)

tmp3645 := GHd(__e, tmp3644)

tmp3646 := GTl(__e, tmp3645)

tmp3647 := GHd(__e, tmp3646)

tmp3648 := GTl(__e, tmp3647)

tmp3649 := GTl(__e, tmp3648)

tmp3650 := GHd(__e, tmp3649)

tmp3651 := GTl(__e, V802)

tmp3652 := GHd(__e, tmp3651)

tmp3653 := GTl(__e, tmp3652)

tmp3654 := GHd(__e, tmp3653)

tmp3655 := GTl(__e, tmp3654)

tmp3656 := GHd(__e, tmp3655)

tmp3657 := GCons(__e, V801, Nil)

tmp3658 := GCons(__e, tmp3656, tmp3657)

tmp3659 := GCons(__e, V800, Nil)

tmp3660 := GCons(__e, symthaw, tmp3659)

tmp3661 := GCons(__e, V801, Nil)

tmp3662 := GCons(__e, tmp3660, tmp3661)

tmp3663 := GCons(__e, tmp3658, tmp3662)

tmp3664 := GCons(__e, symif, tmp3663)

tmp3665 := GCons(__e, tmp3664, Nil)

tmp3666 := GCons(__e, tmp3650, tmp3665)

tmp3667 := GCons(__e, V801, tmp3666)

tmp3668 := GCons(__e, symlet, tmp3667)

tmp3669 := GCons(__e, V800, Nil)

tmp3670 := GCons(__e, symthaw, tmp3669)

tmp3671 := GCons(__e, tmp3670, Nil)

tmp3672 := GCons(__e, tmp3668, tmp3671)

tmp3673 := GCons(__e, tmp3643, tmp3672)

tmp3674 := GCons(__e, symif, tmp3673)

tmp3675 := GCons(__e, tmp3674, Nil)

tmp3676 := GCons(__e, tmp3642, tmp3675)

tmp3677 := GCons(__e, V800, tmp3676)

tmp3678 := GCons(__e, symlet, tmp3677)

tmp3679 := GCons(__e, tmp3678, Nil)

tmp3680 := GCons(__e, True, tmp3679)

return GCons(__e, tmp3680, Nil)


} else {
tmp3739 := GConsP(__e, V802)

var ifres3717 Obj

if True == tmp3739 {
tmp3737 := GTl(__e, V802)

tmp3738 := GConsP(__e, tmp3737)

var ifres3718 Obj

if True == tmp3738 {
tmp3734 := GTl(__e, V802)

tmp3735 := GHd(__e, tmp3734)

tmp3736 := GConsP(__e, tmp3735)

var ifres3719 Obj

if True == tmp3736 {
tmp3730 := GTl(__e, V802)

tmp3731 := GHd(__e, tmp3730)

tmp3732 := GTl(__e, tmp3731)

tmp3733 := GConsP(__e, tmp3732)

var ifres3720 Obj

if True == tmp3733 {
tmp3725 := GTl(__e, V802)

tmp3726 := GHd(__e, tmp3725)

tmp3727 := GTl(__e, tmp3726)

tmp3728 := GTl(__e, tmp3727)

tmp3729 := GEq(__e, Nil, tmp3728)

var ifres3721 Obj

if True == tmp3729 {
tmp3722 := GTl(__e, V802)

tmp3723 := GTl(__e, tmp3722)

tmp3724 := GEq(__e, Nil, tmp3723)

ifres3721 = tmp3724


} else {
ifres3721 = False


}

ifres3720 = ifres3721


} else {
ifres3720 = False


}

ifres3719 = ifres3720


} else {
ifres3719 = False


}

ifres3718 = ifres3719


} else {
ifres3718 = False


}

ifres3717 = ifres3718


} else {
ifres3717 = False


}

if True == ifres3717 {
tmp3681 := Call2(__e, SymFn(symshen_4scan_1body), V799, V803)


tmp3682 := GCons(__e, symcond, tmp3681)

tmp3683 := GCons(__e, tmp3682, Nil)

tmp3684 := GCons(__e, symfreeze, tmp3683)

tmp3685 := GHd(__e, V802)

tmp3686 := GTl(__e, V802)

tmp3687 := GHd(__e, tmp3686)

tmp3688 := GTl(__e, tmp3687)

tmp3689 := GHd(__e, tmp3688)

tmp3690 := GCons(__e, symfail, Nil)

tmp3691 := GCons(__e, tmp3690, Nil)

tmp3692 := GCons(__e, V801, tmp3691)

tmp3693 := GCons(__e, sym_a, tmp3692)

tmp3694 := GCons(__e, V800, Nil)

tmp3695 := GCons(__e, symthaw, tmp3694)

tmp3696 := GCons(__e, V801, Nil)

tmp3697 := GCons(__e, tmp3695, tmp3696)

tmp3698 := GCons(__e, tmp3693, tmp3697)

tmp3699 := GCons(__e, symif, tmp3698)

tmp3700 := GCons(__e, tmp3699, Nil)

tmp3701 := GCons(__e, tmp3689, tmp3700)

tmp3702 := GCons(__e, V801, tmp3701)

tmp3703 := GCons(__e, symlet, tmp3702)

tmp3704 := GCons(__e, V800, Nil)

tmp3705 := GCons(__e, symthaw, tmp3704)

tmp3706 := GCons(__e, tmp3705, Nil)

tmp3707 := GCons(__e, tmp3703, tmp3706)

tmp3708 := GCons(__e, tmp3685, tmp3707)

tmp3709 := GCons(__e, symif, tmp3708)

tmp3710 := GCons(__e, tmp3709, Nil)

tmp3711 := GCons(__e, tmp3684, tmp3710)

tmp3712 := GCons(__e, V800, tmp3711)

tmp3713 := GCons(__e, symlet, tmp3712)

tmp3714 := GCons(__e, tmp3713, Nil)

tmp3715 := GCons(__e, True, tmp3714)

return GCons(__e, tmp3715, Nil)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.choicepoint"))
}


}


panic("unreachable")
})

tmp3809 := Call2(__e, ns2_1set, symshen_4choicepoint, tmp3632)


_ = tmp3809

tmp3810 := MakeNative3(func(__e *ControlFlow, V805, V806, V807 Obj) Obj {
tmp3823 := GEq(__e, V805, V807)

if True == tmp3823 {
return V806
} else {
tmp3821 := GConsP(__e, V807)

if True == tmp3821 {
tmp3811 := GHd(__e, V807)

tmp3812 := Call3(__e, SymFn(symshen_4rep_1X), V805, V806, tmp3811)


let__2629 := tmp3812
_ = let__2629

tmp3818 := GHd(__e, V807)

tmp3819 := GEq(__e, let__2629, tmp3818)

if True == tmp3819 {
tmp3813 := GHd(__e, V807)

tmp3814 := GTl(__e, V807)

tmp3815 := Call3(__e, SymFn(symshen_4rep_1X), V805, V806, tmp3814)


return GCons(__e, tmp3813, tmp3815)


} else {
tmp3816 := GTl(__e, V807)

return GCons(__e, let__2629, tmp3816)


}


} else {
return V807
}


}


panic("unreachable")
})

tmp3824 := Call2(__e, ns2_1set, symshen_4rep_1X, tmp3810)


_ = tmp3824

tmp3825 := MakeNative1(func(__e *ControlFlow, V809 Obj) Obj {
tmp3895 := GConsP(__e, V809)

var ifres3880 Obj

if True == tmp3895 {
tmp3893 := GHd(__e, V809)

tmp3894 := GEq(__e, symlambda, tmp3893)

var ifres3881 Obj

if True == tmp3894 {
tmp3891 := GTl(__e, V809)

tmp3892 := GConsP(__e, tmp3891)

var ifres3882 Obj

if True == tmp3892 {
tmp3888 := GTl(__e, V809)

tmp3889 := GTl(__e, tmp3888)

tmp3890 := GConsP(__e, tmp3889)

var ifres3883 Obj

if True == tmp3890 {
tmp3884 := GTl(__e, V809)

tmp3885 := GTl(__e, tmp3884)

tmp3886 := GTl(__e, tmp3885)

tmp3887 := GEq(__e, Nil, tmp3886)

ifres3883 = tmp3887


} else {
ifres3883 = False


}

ifres3882 = ifres3883


} else {
ifres3882 = False


}

ifres3881 = ifres3882


} else {
ifres3881 = False


}

ifres3880 = ifres3881


} else {
ifres3880 = False


}

if True == ifres3880 {
tmp3826 := Call1(__e, SymFn(symgensym), symZ)


let__2630 := tmp3826
_ = let__2630

tmp3827 := GTl(__e, V809)

tmp3828 := GHd(__e, tmp3827)

tmp3829 := GTl(__e, V809)

tmp3830 := GTl(__e, tmp3829)

tmp3831 := GHd(__e, tmp3830)

tmp3832 := Call3(__e, SymFn(symshen_4beta), tmp3828, let__2630, tmp3831)


tmp3833 := GCons(__e, tmp3832, Nil)

tmp3834 := GCons(__e, let__2630, tmp3833)

tmp3835 := GCons(__e, symlambda, tmp3834)

let__2631 := tmp3835
_ = let__2631

tmp3836 := MakeNative1(func(__e *ControlFlow, Z812 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4alpha_1convert), Z812)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp3836, let__2631)


} else {
tmp3878 := GConsP(__e, V809)

var ifres3857 Obj

if True == tmp3878 {
tmp3876 := GHd(__e, V809)

tmp3877 := GEq(__e, symlet, tmp3876)

var ifres3858 Obj

if True == tmp3877 {
tmp3874 := GTl(__e, V809)

tmp3875 := GConsP(__e, tmp3874)

var ifres3859 Obj

if True == tmp3875 {
tmp3871 := GTl(__e, V809)

tmp3872 := GTl(__e, tmp3871)

tmp3873 := GConsP(__e, tmp3872)

var ifres3860 Obj

if True == tmp3873 {
tmp3867 := GTl(__e, V809)

tmp3868 := GTl(__e, tmp3867)

tmp3869 := GTl(__e, tmp3868)

tmp3870 := GConsP(__e, tmp3869)

var ifres3861 Obj

if True == tmp3870 {
tmp3862 := GTl(__e, V809)

tmp3863 := GTl(__e, tmp3862)

tmp3864 := GTl(__e, tmp3863)

tmp3865 := GTl(__e, tmp3864)

tmp3866 := GEq(__e, Nil, tmp3865)

ifres3861 = tmp3866


} else {
ifres3861 = False


}

ifres3860 = ifres3861


} else {
ifres3860 = False


}

ifres3859 = ifres3860


} else {
ifres3859 = False


}

ifres3858 = ifres3859


} else {
ifres3858 = False


}

ifres3857 = ifres3858


} else {
ifres3857 = False


}

if True == ifres3857 {
tmp3837 := Call1(__e, SymFn(symgensym), symW)


let__2632 := tmp3837
_ = let__2632

tmp3838 := GTl(__e, V809)

tmp3839 := GTl(__e, tmp3838)

tmp3840 := GHd(__e, tmp3839)

tmp3841 := GTl(__e, V809)

tmp3842 := GHd(__e, tmp3841)

tmp3843 := GTl(__e, V809)

tmp3844 := GTl(__e, tmp3843)

tmp3845 := GTl(__e, tmp3844)

tmp3846 := GHd(__e, tmp3845)

tmp3847 := Call3(__e, SymFn(symshen_4beta), tmp3842, let__2632, tmp3846)


tmp3848 := GCons(__e, tmp3847, Nil)

tmp3849 := GCons(__e, tmp3840, tmp3848)

tmp3850 := GCons(__e, let__2632, tmp3849)

tmp3851 := GCons(__e, symlet, tmp3850)

let__2633 := tmp3851
_ = let__2633

tmp3852 := MakeNative1(func(__e *ControlFlow, Z815 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4alpha_1convert), Z815)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp3852, let__2633)


} else {
tmp3855 := GConsP(__e, V809)

if True == tmp3855 {
tmp3853 := MakeNative1(func(__e *ControlFlow, Z816 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4alpha_1convert), Z816)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp3853, V809)


} else {
return V809
}


}


}


panic("unreachable")
})

tmp3896 := Call2(__e, ns2_1set, symshen_4alpha_1convert, tmp3825)


_ = tmp3896

tmp3897 := MakeNative2(func(__e *ControlFlow, V817, V818 Obj) Obj {
tmp3898 := MakeNative1(func(__e *ControlFlow, Z819 Obj) Obj {
tmp3899 := Call1(__e, SymFn(symfst), Z819)


tmp3900 := Call1(__e, SymFn(symsnd), Z819)


tmp3901 := Call1(__e, SymFn(symshen_4alpha_1convert), tmp3900)


return TailCall4(__e, SymFn(symshen_4triple_1stack), Nil, tmp3899, V818, tmp3901)


panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp3898, V817)


panic("unreachable")
})

tmp3902 := Call2(__e, ns2_1set, symshen_4kl_1body, tmp3897)


_ = tmp3902

tmp3903 := (func() Obj {
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
tmp4034 := GEq(__e, Nil, V829)

var ifres4015 Obj

if True == tmp4034 {
tmp4033 := GEq(__e, Nil, V830)

var ifres4016 Obj

if True == tmp4033 {
tmp4032 := GConsP(__e, V831)

var ifres4017 Obj

if True == tmp4032 {
tmp4030 := GHd(__e, V831)

tmp4031 := GEq(__e, symwhere, tmp4030)

var ifres4018 Obj

if True == tmp4031 {
tmp4028 := GTl(__e, V831)

tmp4029 := GConsP(__e, tmp4028)

var ifres4019 Obj

if True == tmp4029 {
tmp4025 := GTl(__e, V831)

tmp4026 := GTl(__e, tmp4025)

tmp4027 := GConsP(__e, tmp4026)

var ifres4020 Obj

if True == tmp4027 {
tmp4021 := GTl(__e, V831)

tmp4022 := GTl(__e, tmp4021)

tmp4023 := GTl(__e, tmp4022)

tmp4024 := GEq(__e, Nil, tmp4023)

ifres4020 = tmp4024


} else {
ifres4020 = False


}

ifres4019 = ifres4020


} else {
ifres4019 = False


}

ifres4018 = ifres4019


} else {
ifres4018 = False


}

ifres4017 = ifres4018


} else {
ifres4017 = False


}

ifres4016 = ifres4017


} else {
ifres4016 = False


}

ifres4015 = ifres4016


} else {
ifres4015 = False


}

if True == ifres4015 {
tmp3904 := GTl(__e, V831)

tmp3905 := GHd(__e, tmp3904)

tmp3906 := GCons(__e, tmp3905, V828)

tmp3907 := GTl(__e, V831)

tmp3908 := GTl(__e, tmp3907)

tmp3909 := GHd(__e, tmp3908)

if SymFn(symshen_4triple_1stack) == __self {
__self1, __self2, __self3, __self4 = tmp3906, Nil, Nil, tmp3909
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4triple_1stack), tmp3906, Nil, Nil, tmp3909)


} else {
tmp4013 := GEq(__e, Nil, V829)

var ifres4011 Obj

if True == tmp4013 {
tmp4012 := GEq(__e, Nil, V830)

ifres4011 = tmp4012


} else {
ifres4011 = False


}

if True == ifres4011 {
tmp3910 := Call1(__e, SymFn(symreverse), V828)


tmp3911 := Call1(__e, SymFn(symshen_4rectify_1test), tmp3910)


tmp3912 := GCons(__e, V831, Nil)

return GCons(__e, tmp3911, tmp3912)


} else {
tmp4009 := GConsP(__e, V829)

var ifres4004 Obj

if True == tmp4009 {
tmp4008 := GConsP(__e, V830)

var ifres4005 Obj

if True == tmp4008 {
tmp4006 := GHd(__e, V829)

tmp4007 := PrimGuard1(__e, symvariable_2, PrimIsVariable, tmp4006)

ifres4005 = tmp4007


} else {
ifres4005 = False


}

ifres4004 = ifres4005


} else {
ifres4004 = False


}

if True == ifres4004 {
tmp3913 := GTl(__e, V829)

tmp3914 := GTl(__e, V830)

tmp3915 := GHd(__e, V829)

tmp3916 := GHd(__e, V830)

tmp3917 := Call3(__e, SymFn(symshen_4beta), tmp3915, tmp3916, V831)


if SymFn(symshen_4triple_1stack) == __self {
__self1, __self2, __self3, __self4 = V828, tmp3913, tmp3914, tmp3917
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4triple_1stack), V828, tmp3913, tmp3914, tmp3917)


} else {
tmp4002 := GConsP(__e, V829)

var ifres3997 Obj

if True == tmp4002 {
tmp4001 := GConsP(__e, V830)

var ifres3998 Obj

if True == tmp4001 {
tmp3999 := GHd(__e, V829)

tmp4000 := Call1(__e, SymFn(symshen_4custom_1pattern_2), tmp3999)


ifres3998 = tmp4000


} else {
ifres3998 = False


}

ifres3997 = ifres3998


} else {
ifres3997 = False


}

if True == ifres3997 {
tmp3918 := GHd(__e, V829)

tmp3919 := Call1(__e, SymFn(symshen_4custom_1pattern_1body), tmp3918)


tmp3920 := GTl(__e, V829)

tmp3921 := GHd(__e, V830)

tmp3922 := GTl(__e, V830)

return TailCall6(__e, SymFn(symshen_4custom_1pattern_1triple_1stack), V828, tmp3919, tmp3920, tmp3921, tmp3922, V831)


} else {
tmp3995 := GConsP(__e, V829)

var ifres3971 Obj

if True == tmp3995 {
tmp3993 := GHd(__e, V829)

tmp3994 := GConsP(__e, tmp3993)

var ifres3972 Obj

if True == tmp3994 {
tmp3990 := GHd(__e, V829)

tmp3991 := GTl(__e, tmp3990)

tmp3992 := GConsP(__e, tmp3991)

var ifres3973 Obj

if True == tmp3992 {
tmp3986 := GHd(__e, V829)

tmp3987 := GTl(__e, tmp3986)

tmp3988 := GTl(__e, tmp3987)

tmp3989 := GConsP(__e, tmp3988)

var ifres3974 Obj

if True == tmp3989 {
tmp3981 := GHd(__e, V829)

tmp3982 := GTl(__e, tmp3981)

tmp3983 := GTl(__e, tmp3982)

tmp3984 := GTl(__e, tmp3983)

tmp3985 := GEq(__e, Nil, tmp3984)

var ifres3975 Obj

if True == tmp3985 {
tmp3980 := GConsP(__e, V830)

var ifres3976 Obj

if True == tmp3980 {
tmp3977 := GHd(__e, V829)

tmp3978 := GHd(__e, tmp3977)

tmp3979 := Call1(__e, SymFn(symshen_4constructor_2), tmp3978)


ifres3976 = tmp3979


} else {
ifres3976 = False


}

ifres3975 = ifres3976


} else {
ifres3975 = False


}

ifres3974 = ifres3975


} else {
ifres3974 = False


}

ifres3973 = ifres3974


} else {
ifres3973 = False


}

ifres3972 = ifres3973


} else {
ifres3972 = False


}

ifres3971 = ifres3972


} else {
ifres3971 = False


}

if True == ifres3971 {
tmp3923 := GHd(__e, V829)

tmp3924 := GHd(__e, tmp3923)

tmp3925 := Call1(__e, SymFn(symshen_4op_1test), tmp3924)


tmp3926 := GHd(__e, V830)

tmp3927 := GCons(__e, tmp3926, Nil)

tmp3928 := GCons(__e, tmp3925, tmp3927)

tmp3929 := GCons(__e, tmp3928, V828)

tmp3930 := GHd(__e, V829)

tmp3931 := GTl(__e, tmp3930)

tmp3932 := GHd(__e, tmp3931)

tmp3933 := GHd(__e, V829)

tmp3934 := GTl(__e, tmp3933)

tmp3935 := GTl(__e, tmp3934)

tmp3936 := GHd(__e, tmp3935)

tmp3937 := GTl(__e, V829)

tmp3938 := GCons(__e, tmp3936, tmp3937)

tmp3939 := GCons(__e, tmp3932, tmp3938)

tmp3940 := GHd(__e, V829)

tmp3941 := GHd(__e, tmp3940)

tmp3942 := Call1(__e, SymFn(symshen_4op1), tmp3941)


tmp3943 := GHd(__e, V830)

tmp3944 := GCons(__e, tmp3943, Nil)

tmp3945 := GCons(__e, tmp3942, tmp3944)

tmp3946 := GHd(__e, V829)

tmp3947 := GHd(__e, tmp3946)

tmp3948 := Call1(__e, SymFn(symshen_4op2), tmp3947)


tmp3949 := GHd(__e, V830)

tmp3950 := GCons(__e, tmp3949, Nil)

tmp3951 := GCons(__e, tmp3948, tmp3950)

tmp3952 := GTl(__e, V830)

tmp3953 := GCons(__e, tmp3951, tmp3952)

tmp3954 := GCons(__e, tmp3945, tmp3953)

tmp3955 := GHd(__e, V829)

tmp3956 := GHd(__e, V830)

tmp3957 := Call3(__e, SymFn(symshen_4beta), tmp3955, tmp3956, V831)


if SymFn(symshen_4triple_1stack) == __self {
__self1, __self2, __self3, __self4 = tmp3929, tmp3939, tmp3954, tmp3957
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4triple_1stack), tmp3929, tmp3939, tmp3954, tmp3957)


} else {
tmp3969 := GConsP(__e, V829)

var ifres3967 Obj

if True == tmp3969 {
tmp3968 := GConsP(__e, V830)

ifres3967 = tmp3968


} else {
ifres3967 = False


}

if True == ifres3967 {
tmp3958 := GHd(__e, V829)

tmp3959 := GHd(__e, V830)

tmp3960 := GCons(__e, tmp3959, Nil)

tmp3961 := GCons(__e, tmp3958, tmp3960)

tmp3962 := GCons(__e, sym_a, tmp3961)

tmp3963 := GCons(__e, tmp3962, V828)

tmp3964 := GTl(__e, V829)

tmp3965 := GTl(__e, V830)

if SymFn(symshen_4triple_1stack) == __self {
__self1, __self2, __self3, __self4 = tmp3963, tmp3964, tmp3965, V831
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4triple_1stack), tmp3963, tmp3964, tmp3965, V831)


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

tmp4035 := Call2(__e, ns2_1set, symshen_4triple_1stack, tmp3903)


_ = tmp4035

tmp4036 := MakeNative6(func(__e *ControlFlow, V832, V833, V834, V835, V836, V837 Obj) Obj {
tmp4037 := Call2(__e, SymFn(sym_8p), V833, V835)


tmp4038 := Call1(__e, SymFn(symshen_4custom_1pattern_1reducer), tmp4037)


let__2634 := tmp4038
_ = let__2634

tmp4051 := Call1(__e, SymFn(symtuple_2), let__2634)


if True == tmp4051 {
tmp4039 := Call1(__e, SymFn(symfst), let__2634)


let__2635 := tmp4039
_ = let__2635

tmp4040 := Call1(__e, SymFn(symsnd), let__2634)


let__2636 := tmp4040
_ = let__2636

tmp4041 := Call1(__e, SymFn(symreverse), let__2635)


tmp4042 := Call2(__e, SymFn(symappend), tmp4041, V832)


tmp4043 := MakeNative1(func(__e *ControlFlow, Z841 Obj) Obj {
return TailCall1(__e, SymFn(symfst), Z841)
panic("unreachable")
})

tmp4044 := Call2(__e, SymFn(symmap), tmp4043, let__2636)


tmp4045 := Call2(__e, SymFn(symappend), tmp4044, V834)


tmp4046 := MakeNative1(func(__e *ControlFlow, Z842 Obj) Obj {
return TailCall1(__e, SymFn(symsnd), Z842)
panic("unreachable")
})

tmp4047 := Call2(__e, SymFn(symmap), tmp4046, let__2636)


tmp4048 := Call2(__e, SymFn(symappend), tmp4047, V836)


tmp4049 := Call3(__e, SymFn(symshen_4beta), V833, V835, V837)


return TailCall4(__e, SymFn(symshen_4triple_1stack), tmp4042, tmp4045, tmp4048, tmp4049)


} else {
return TailCall1(__e, SymFn(symshen_4constructor_1error), V833)
}


panic("unreachable")
})

tmp4052 := Call2(__e, ns2_1set, symshen_4custom_1pattern_1triple_1stack, tmp4036)


_ = tmp4052

tmp4053 := MakeNative1(func(__e *ControlFlow, V845 Obj) Obj {
tmp4070 := GEq(__e, Nil, V845)

if True == tmp4070 {
return True
} else {
tmp4068 := GConsP(__e, V845)

var ifres4065 Obj

if True == tmp4068 {
tmp4066 := GTl(__e, V845)

tmp4067 := GEq(__e, Nil, tmp4066)

ifres4065 = tmp4067


} else {
ifres4065 = False


}

if True == ifres4065 {
return GHd(__e, V845)
} else {
tmp4063 := GConsP(__e, V845)

var ifres4060 Obj

if True == tmp4063 {
tmp4061 := GTl(__e, V845)

tmp4062 := GConsP(__e, tmp4061)

ifres4060 = tmp4062


} else {
ifres4060 = False


}

if True == ifres4060 {
tmp4054 := GHd(__e, V845)

tmp4055 := GTl(__e, V845)

tmp4056 := Call1(__e, SymFn(symshen_4rectify_1test), tmp4055)


tmp4057 := GCons(__e, tmp4056, Nil)

tmp4058 := GCons(__e, tmp4054, tmp4057)

return GCons(__e, symand, tmp4058)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.rectify-test"))
}


}


}


panic("unreachable")
})

tmp4071 := Call2(__e, ns2_1set, symshen_4rectify_1test, tmp4053)


_ = tmp4071

tmp4072 := MakeNative3(func(__e *ControlFlow, V855, V856, V857 Obj) Obj {
tmp4138 := GEq(__e, V855, V857)

if True == tmp4138 {
return V856
} else {
tmp4136 := GConsP(__e, V857)

var ifres4117 Obj

if True == tmp4136 {
tmp4134 := GHd(__e, V857)

tmp4135 := GEq(__e, symlambda, tmp4134)

var ifres4118 Obj

if True == tmp4135 {
tmp4132 := GTl(__e, V857)

tmp4133 := GConsP(__e, tmp4132)

var ifres4119 Obj

if True == tmp4133 {
tmp4129 := GTl(__e, V857)

tmp4130 := GTl(__e, tmp4129)

tmp4131 := GConsP(__e, tmp4130)

var ifres4120 Obj

if True == tmp4131 {
tmp4125 := GTl(__e, V857)

tmp4126 := GTl(__e, tmp4125)

tmp4127 := GTl(__e, tmp4126)

tmp4128 := GEq(__e, Nil, tmp4127)

var ifres4121 Obj

if True == tmp4128 {
tmp4122 := GTl(__e, V857)

tmp4123 := GHd(__e, tmp4122)

tmp4124 := GEq(__e, V855, tmp4123)

ifres4121 = tmp4124


} else {
ifres4121 = False


}

ifres4120 = ifres4121


} else {
ifres4120 = False


}

ifres4119 = ifres4120


} else {
ifres4119 = False


}

ifres4118 = ifres4119


} else {
ifres4118 = False


}

ifres4117 = ifres4118


} else {
ifres4117 = False


}

if True == ifres4117 {
return V857
} else {
tmp4115 := GConsP(__e, V857)

var ifres4090 Obj

if True == tmp4115 {
tmp4113 := GHd(__e, V857)

tmp4114 := GEq(__e, symlet, tmp4113)

var ifres4091 Obj

if True == tmp4114 {
tmp4111 := GTl(__e, V857)

tmp4112 := GConsP(__e, tmp4111)

var ifres4092 Obj

if True == tmp4112 {
tmp4108 := GTl(__e, V857)

tmp4109 := GTl(__e, tmp4108)

tmp4110 := GConsP(__e, tmp4109)

var ifres4093 Obj

if True == tmp4110 {
tmp4104 := GTl(__e, V857)

tmp4105 := GTl(__e, tmp4104)

tmp4106 := GTl(__e, tmp4105)

tmp4107 := GConsP(__e, tmp4106)

var ifres4094 Obj

if True == tmp4107 {
tmp4099 := GTl(__e, V857)

tmp4100 := GTl(__e, tmp4099)

tmp4101 := GTl(__e, tmp4100)

tmp4102 := GTl(__e, tmp4101)

tmp4103 := GEq(__e, Nil, tmp4102)

var ifres4095 Obj

if True == tmp4103 {
tmp4096 := GTl(__e, V857)

tmp4097 := GHd(__e, tmp4096)

tmp4098 := GEq(__e, V855, tmp4097)

ifres4095 = tmp4098


} else {
ifres4095 = False


}

ifres4094 = ifres4095


} else {
ifres4094 = False


}

ifres4093 = ifres4094


} else {
ifres4093 = False


}

ifres4092 = ifres4093


} else {
ifres4092 = False


}

ifres4091 = ifres4092


} else {
ifres4091 = False


}

ifres4090 = ifres4091


} else {
ifres4090 = False


}

if True == ifres4090 {
tmp4073 := GTl(__e, V857)

tmp4074 := GHd(__e, tmp4073)

tmp4075 := GTl(__e, V857)

tmp4076 := GHd(__e, tmp4075)

tmp4077 := GTl(__e, V857)

tmp4078 := GTl(__e, tmp4077)

tmp4079 := GHd(__e, tmp4078)

tmp4080 := Call3(__e, SymFn(symshen_4beta), tmp4076, V856, tmp4079)


tmp4081 := GTl(__e, V857)

tmp4082 := GTl(__e, tmp4081)

tmp4083 := GTl(__e, tmp4082)

tmp4084 := GCons(__e, tmp4080, tmp4083)

tmp4085 := GCons(__e, tmp4074, tmp4084)

return GCons(__e, symlet, tmp4085)


} else {
tmp4088 := GConsP(__e, V857)

if True == tmp4088 {
tmp4086 := MakeNative1(func(__e *ControlFlow, Z858 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4beta), V855, V856, Z858)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp4086, V857)


} else {
return V857
}


}


}


}


panic("unreachable")
})

tmp4139 := Call2(__e, ns2_1set, symshen_4beta, tmp4072)


_ = tmp4139

tmp4140 := MakeNative1(func(__e *ControlFlow, V861 Obj) Obj {
tmp4148 := GEq(__e, symcons, V861)

if True == tmp4148 {
return symhd
} else {
tmp4146 := GEq(__e, sym_8s, V861)

if True == tmp4146 {
return symhdstr
} else {
tmp4144 := GEq(__e, sym_8p, V861)

if True == tmp4144 {
return symfst
} else {
tmp4142 := GEq(__e, sym_8v, V861)

if True == tmp4142 {
return symhdv
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.op1"))
}


}


}


}


panic("unreachable")
})

tmp4149 := Call2(__e, ns2_1set, symshen_4op1, tmp4140)


_ = tmp4149

tmp4150 := MakeNative1(func(__e *ControlFlow, V864 Obj) Obj {
tmp4158 := GEq(__e, symcons, V864)

if True == tmp4158 {
return symtl
} else {
tmp4156 := GEq(__e, sym_8s, V864)

if True == tmp4156 {
return symtlstr
} else {
tmp4154 := GEq(__e, sym_8p, V864)

if True == tmp4154 {
return symsnd
} else {
tmp4152 := GEq(__e, sym_8v, V864)

if True == tmp4152 {
return symtlv
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.op2"))
}


}


}


}


panic("unreachable")
})

tmp4159 := Call2(__e, ns2_1set, symshen_4op2, tmp4150)


_ = tmp4159

tmp4160 := MakeNative1(func(__e *ControlFlow, V867 Obj) Obj {
tmp4168 := GEq(__e, symcons, V867)

if True == tmp4168 {
return symcons_2
} else {
tmp4166 := GEq(__e, sym_8s, V867)

if True == tmp4166 {
return symshen_4_7string_2
} else {
tmp4164 := GEq(__e, sym_8p, V867)

if True == tmp4164 {
return symtuple_2
} else {
tmp4162 := GEq(__e, sym_8v, V867)

if True == tmp4162 {
return symshen_4_7vector_2
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.op-test"))
}


}


}


}


panic("unreachable")
})

tmp4169 := Call2(__e, ns2_1set, symshen_4op_1test, tmp4160)


_ = tmp4169

tmp4170 := MakeNative1(func(__e *ControlFlow, V868 Obj) Obj {
tmp4172 := GEq(__e, MakeString(""), V868)

if True == tmp4172 {
return False
} else {
return PrimGuard1(__e, symstring_2, PrimIsString, V868)
}


panic("unreachable")
})

tmp4173 := Call2(__e, ns2_1set, symshen_4_7string_2, tmp4170)


_ = tmp4173

tmp4174 := MakeNative1(func(__e *ControlFlow, V869 Obj) Obj {
tmp4176 := Call1(__e, SymFn(symvector), MakeInteger(0))


tmp4177 := GEq(__e, V869, tmp4176)

if True == tmp4177 {
return False
} else {
return TailCall1(__e, SymFn(symvector_2), V869)
}


panic("unreachable")
})

tmp4178 := Call2(__e, ns2_1set, symshen_4_7vector_2, tmp4174)


_ = tmp4178

tmp4179 := MakeNative1(func(__e *ControlFlow, V872 Obj) Obj {
tmp4183 := GEq(__e, sym_7, V872)

if True == tmp4183 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dfactorise_2_d, True)
} else {
tmp4181 := GEq(__e, sym_1, V872)

if True == tmp4181 {
return PrimGuard2(__e, symset, PrimSet, symshen_4_dfactorise_2_d, False)
} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("factorise expects a + or a -\n"))
}


}


panic("unreachable")
})

tmp4184 := Call2(__e, ns2_1set, symfactorise, tmp4179)


_ = tmp4184

tmp4185 := MakeNative1(func(__e *ControlFlow, V873 Obj) Obj {
tmp4187 := GValue(__e, symshen_4_dfactorise_2_d)

if True == tmp4187 {
return TailCall1(__e, SymFn(symshen_4factor), V873)
} else {
return V873
}


panic("unreachable")
})

tmp4188 := Call2(__e, ns2_1set, symshen_4factorise_1code, tmp4185)


_ = tmp4188

tmp4189 := MakeNative1(func(__e *ControlFlow, V874 Obj) Obj {
tmp4239 := GConsP(__e, V874)

var ifres4205 Obj

if True == tmp4239 {
tmp4237 := GHd(__e, V874)

tmp4238 := GEq(__e, symdefun, tmp4237)

var ifres4206 Obj

if True == tmp4238 {
tmp4235 := GTl(__e, V874)

tmp4236 := GConsP(__e, tmp4235)

var ifres4207 Obj

if True == tmp4236 {
tmp4232 := GTl(__e, V874)

tmp4233 := GTl(__e, tmp4232)

tmp4234 := GConsP(__e, tmp4233)

var ifres4208 Obj

if True == tmp4234 {
tmp4228 := GTl(__e, V874)

tmp4229 := GTl(__e, tmp4228)

tmp4230 := GTl(__e, tmp4229)

tmp4231 := GConsP(__e, tmp4230)

var ifres4209 Obj

if True == tmp4231 {
tmp4223 := GTl(__e, V874)

tmp4224 := GTl(__e, tmp4223)

tmp4225 := GTl(__e, tmp4224)

tmp4226 := GHd(__e, tmp4225)

tmp4227 := GConsP(__e, tmp4226)

var ifres4210 Obj

if True == tmp4227 {
tmp4217 := GTl(__e, V874)

tmp4218 := GTl(__e, tmp4217)

tmp4219 := GTl(__e, tmp4218)

tmp4220 := GHd(__e, tmp4219)

tmp4221 := GHd(__e, tmp4220)

tmp4222 := GEq(__e, symcond, tmp4221)

var ifres4211 Obj

if True == tmp4222 {
tmp4212 := GTl(__e, V874)

tmp4213 := GTl(__e, tmp4212)

tmp4214 := GTl(__e, tmp4213)

tmp4215 := GTl(__e, tmp4214)

tmp4216 := GEq(__e, Nil, tmp4215)

ifres4211 = tmp4216


} else {
ifres4211 = False


}

ifres4210 = ifres4211


} else {
ifres4210 = False


}

ifres4209 = ifres4210


} else {
ifres4209 = False


}

ifres4208 = ifres4209


} else {
ifres4208 = False


}

ifres4207 = ifres4208


} else {
ifres4207 = False


}

ifres4206 = ifres4207


} else {
ifres4206 = False


}

ifres4205 = ifres4206


} else {
ifres4205 = False


}

if True == ifres4205 {
tmp4190 := GTl(__e, V874)

tmp4191 := GHd(__e, tmp4190)

tmp4192 := GTl(__e, V874)

tmp4193 := GTl(__e, tmp4192)

tmp4194 := GHd(__e, tmp4193)

tmp4195 := GTl(__e, V874)

tmp4196 := GTl(__e, tmp4195)

tmp4197 := GTl(__e, tmp4196)

tmp4198 := GHd(__e, tmp4197)

tmp4199 := GTl(__e, tmp4198)

tmp4200 := Call1(__e, SymFn(symshen_4factor_1recognisors), tmp4199)


tmp4201 := GCons(__e, tmp4200, Nil)

tmp4202 := GCons(__e, tmp4194, tmp4201)

tmp4203 := GCons(__e, tmp4191, tmp4202)

return GCons(__e, symdefun, tmp4203)


} else {
return V874
}


panic("unreachable")
})

tmp4240 := Call2(__e, ns2_1set, symshen_4factor, tmp4189)


_ = tmp4240

tmp4241 := MakeNative1(func(__e *ControlFlow, V877 Obj) Obj {
tmp4376 := GConsP(__e, V877)

var ifres4360 Obj

if True == tmp4376 {
tmp4374 := GHd(__e, V877)

tmp4375 := GConsP(__e, tmp4374)

var ifres4361 Obj

if True == tmp4375 {
tmp4371 := GHd(__e, V877)

tmp4372 := GHd(__e, tmp4371)

tmp4373 := GEq(__e, True, tmp4372)

var ifres4362 Obj

if True == tmp4373 {
tmp4368 := GHd(__e, V877)

tmp4369 := GTl(__e, tmp4368)

tmp4370 := GConsP(__e, tmp4369)

var ifres4363 Obj

if True == tmp4370 {
tmp4364 := GHd(__e, V877)

tmp4365 := GTl(__e, tmp4364)

tmp4366 := GTl(__e, tmp4365)

tmp4367 := GEq(__e, Nil, tmp4366)

ifres4363 = tmp4367


} else {
ifres4363 = False


}

ifres4362 = ifres4363


} else {
ifres4362 = False


}

ifres4361 = ifres4362


} else {
ifres4361 = False


}

ifres4360 = ifres4361


} else {
ifres4360 = False


}

if True == ifres4360 {
tmp4242 := GHd(__e, V877)

tmp4243 := GTl(__e, tmp4242)

return GHd(__e, tmp4243)


} else {
tmp4358 := GConsP(__e, V877)

var ifres4319 Obj

if True == tmp4358 {
tmp4356 := GHd(__e, V877)

tmp4357 := GConsP(__e, tmp4356)

var ifres4320 Obj

if True == tmp4357 {
tmp4353 := GHd(__e, V877)

tmp4354 := GHd(__e, tmp4353)

tmp4355 := GConsP(__e, tmp4354)

var ifres4321 Obj

if True == tmp4355 {
tmp4349 := GHd(__e, V877)

tmp4350 := GHd(__e, tmp4349)

tmp4351 := GHd(__e, tmp4350)

tmp4352 := GEq(__e, symand, tmp4351)

var ifres4322 Obj

if True == tmp4352 {
tmp4345 := GHd(__e, V877)

tmp4346 := GHd(__e, tmp4345)

tmp4347 := GTl(__e, tmp4346)

tmp4348 := GConsP(__e, tmp4347)

var ifres4323 Obj

if True == tmp4348 {
tmp4340 := GHd(__e, V877)

tmp4341 := GHd(__e, tmp4340)

tmp4342 := GTl(__e, tmp4341)

tmp4343 := GTl(__e, tmp4342)

tmp4344 := GConsP(__e, tmp4343)

var ifres4324 Obj

if True == tmp4344 {
tmp4334 := GHd(__e, V877)

tmp4335 := GHd(__e, tmp4334)

tmp4336 := GTl(__e, tmp4335)

tmp4337 := GTl(__e, tmp4336)

tmp4338 := GTl(__e, tmp4337)

tmp4339 := GEq(__e, Nil, tmp4338)

var ifres4325 Obj

if True == tmp4339 {
tmp4331 := GHd(__e, V877)

tmp4332 := GTl(__e, tmp4331)

tmp4333 := GConsP(__e, tmp4332)

var ifres4326 Obj

if True == tmp4333 {
tmp4327 := GHd(__e, V877)

tmp4328 := GTl(__e, tmp4327)

tmp4329 := GTl(__e, tmp4328)

tmp4330 := GEq(__e, Nil, tmp4329)

ifres4326 = tmp4330


} else {
ifres4326 = False


}

ifres4325 = ifres4326


} else {
ifres4325 = False


}

ifres4324 = ifres4325


} else {
ifres4324 = False


}

ifres4323 = ifres4324


} else {
ifres4323 = False


}

ifres4322 = ifres4323


} else {
ifres4322 = False


}

ifres4321 = ifres4322


} else {
ifres4321 = False


}

ifres4320 = ifres4321


} else {
ifres4320 = False


}

ifres4319 = ifres4320


} else {
ifres4319 = False


}

if True == ifres4319 {
tmp4244 := GHd(__e, V877)

tmp4245 := GHd(__e, tmp4244)

tmp4246 := GTl(__e, tmp4245)

tmp4247 := GHd(__e, tmp4246)

tmp4248 := Call3(__e, SymFn(symshen_4pivot_1on), tmp4247, V877, Nil)


let__2637 := tmp4248
_ = let__2637

tmp4249 := Call1(__e, SymFn(symfst), let__2637)


let__2638 := tmp4249
_ = let__2638

tmp4293 := Call1(__e, SymFn(symshen_4bad_1pivot_2), let__2638)


if True == tmp4293 {
tmp4250 := GCons(__e, symif, Nil)

tmp4251 := GHd(__e, V877)

tmp4252 := GHd(__e, tmp4251)

tmp4253 := GHd(__e, V877)

tmp4254 := GTl(__e, tmp4253)

tmp4255 := GHd(__e, tmp4254)

tmp4256 := Call2(__e, SymFn(symshen_4recursively_1factor_1selectors), tmp4252, tmp4255)


tmp4257 := GTl(__e, V877)

tmp4258 := Call1(__e, SymFn(symshen_4factor_1recognisors), tmp4257)


tmp4259 := GCons(__e, tmp4258, Nil)

tmp4260 := Call2(__e, SymFn(symappend), tmp4256, tmp4259)


return TailCall2(__e, SymFn(symappend), tmp4250, tmp4260)


} else {
tmp4261 := Call1(__e, SymFn(symsnd), let__2637)


let__2639 := tmp4261
_ = let__2639

tmp4262 := Call1(__e, SymFn(symshen_4factor_1recognisors), let__2639)


let__2640 := tmp4262
_ = let__2640

tmp4263 := Call1(__e, SymFn(symgensym), symGoTo)


let__2641 := tmp4263
_ = let__2641

tmp4264 := GCons(__e, let__2641, Nil)

tmp4265 := GCons(__e, symthaw, tmp4264)

tmp4266 := GCons(__e, tmp4265, Nil)

tmp4267 := GCons(__e, True, tmp4266)

tmp4268 := GCons(__e, tmp4267, let__2638)

tmp4269 := Call1(__e, SymFn(symreverse), tmp4268)


let__2642 := tmp4269
_ = let__2642

tmp4270 := GCons(__e, let__2640, Nil)

tmp4271 := GCons(__e, symfreeze, tmp4270)

tmp4272 := GHd(__e, V877)

tmp4273 := GHd(__e, tmp4272)

tmp4274 := GTl(__e, tmp4273)

tmp4275 := GHd(__e, tmp4274)

tmp4276 := GHd(__e, V877)

tmp4277 := GHd(__e, tmp4276)

tmp4278 := GTl(__e, tmp4277)

tmp4279 := GHd(__e, tmp4278)

tmp4280 := Call1(__e, SymFn(symshen_4factor_1recognisors), let__2642)


tmp4281 := Call2(__e, SymFn(symshen_4factor_1selectors), tmp4279, tmp4280)


tmp4282 := GCons(__e, let__2641, Nil)

tmp4283 := GCons(__e, symthaw, tmp4282)

tmp4284 := GCons(__e, tmp4283, Nil)

tmp4285 := GCons(__e, tmp4281, tmp4284)

tmp4286 := GCons(__e, tmp4275, tmp4285)

tmp4287 := GCons(__e, symif, tmp4286)

tmp4288 := GCons(__e, tmp4287, Nil)

tmp4289 := GCons(__e, tmp4271, tmp4288)

tmp4290 := GCons(__e, let__2641, tmp4289)

tmp4291 := GCons(__e, symlet, tmp4290)

let__2643 := tmp4291
_ = let__2643

return TailCall1(__e, SymFn(symshen_4remove_1indirection), let__2643)


}


} else {
tmp4317 := GConsP(__e, V877)

var ifres4305 Obj

if True == tmp4317 {
tmp4315 := GHd(__e, V877)

tmp4316 := GConsP(__e, tmp4315)

var ifres4306 Obj

if True == tmp4316 {
tmp4312 := GHd(__e, V877)

tmp4313 := GTl(__e, tmp4312)

tmp4314 := GConsP(__e, tmp4313)

var ifres4307 Obj

if True == tmp4314 {
tmp4308 := GHd(__e, V877)

tmp4309 := GTl(__e, tmp4308)

tmp4310 := GTl(__e, tmp4309)

tmp4311 := GEq(__e, Nil, tmp4310)

ifres4307 = tmp4311


} else {
ifres4307 = False


}

ifres4306 = ifres4307


} else {
ifres4306 = False


}

ifres4305 = ifres4306


} else {
ifres4305 = False


}

if True == ifres4305 {
tmp4294 := GHd(__e, V877)

tmp4295 := GHd(__e, tmp4294)

tmp4296 := GHd(__e, V877)

tmp4297 := GTl(__e, tmp4296)

tmp4298 := GHd(__e, tmp4297)

tmp4299 := GTl(__e, V877)

tmp4300 := Call1(__e, SymFn(symshen_4factor_1recognisors), tmp4299)


tmp4301 := GCons(__e, tmp4300, Nil)

tmp4302 := GCons(__e, tmp4298, tmp4301)

tmp4303 := GCons(__e, tmp4295, tmp4302)

return GCons(__e, symif, tmp4303)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4factor_1recognisors)
}


}


}


panic("unreachable")
})

tmp4377 := Call2(__e, ns2_1set, symshen_4factor_1recognisors, tmp4241)


_ = tmp4377

tmp4378 := MakeNative2(func(__e *ControlFlow, V885, V886 Obj) Obj {
tmp4441 := GConsP(__e, V885)

var ifres4420 Obj

if True == tmp4441 {
tmp4439 := GHd(__e, V885)

tmp4440 := GEq(__e, symlet, tmp4439)

var ifres4421 Obj

if True == tmp4440 {
tmp4437 := GTl(__e, V885)

tmp4438 := GConsP(__e, tmp4437)

var ifres4422 Obj

if True == tmp4438 {
tmp4434 := GTl(__e, V885)

tmp4435 := GTl(__e, tmp4434)

tmp4436 := GConsP(__e, tmp4435)

var ifres4423 Obj

if True == tmp4436 {
tmp4430 := GTl(__e, V885)

tmp4431 := GTl(__e, tmp4430)

tmp4432 := GTl(__e, tmp4431)

tmp4433 := GConsP(__e, tmp4432)

var ifres4424 Obj

if True == tmp4433 {
tmp4425 := GTl(__e, V885)

tmp4426 := GTl(__e, tmp4425)

tmp4427 := GTl(__e, tmp4426)

tmp4428 := GTl(__e, tmp4427)

tmp4429 := GEq(__e, Nil, tmp4428)

ifres4424 = tmp4429


} else {
ifres4424 = False


}

ifres4423 = ifres4424


} else {
ifres4423 = False


}

ifres4422 = ifres4423


} else {
ifres4422 = False


}

ifres4421 = ifres4422


} else {
ifres4421 = False


}

ifres4420 = ifres4421


} else {
ifres4420 = False


}

if True == ifres4420 {
tmp4379 := GTl(__e, V885)

tmp4380 := GHd(__e, tmp4379)

tmp4381 := GTl(__e, V885)

tmp4382 := GTl(__e, tmp4381)

tmp4383 := GHd(__e, tmp4382)

tmp4384 := GTl(__e, V885)

tmp4385 := GTl(__e, tmp4384)

tmp4386 := GTl(__e, tmp4385)

tmp4387 := GHd(__e, tmp4386)

tmp4388 := Call2(__e, SymFn(symshen_4recursively_1factor_1selectors), tmp4387, V886)


return TailCall3(__e, SymFn(symshen_4restore_1local), tmp4380, tmp4383, tmp4388)


} else {
tmp4418 := GConsP(__e, V885)

var ifres4403 Obj

if True == tmp4418 {
tmp4416 := GHd(__e, V885)

tmp4417 := GEq(__e, symand, tmp4416)

var ifres4404 Obj

if True == tmp4417 {
tmp4414 := GTl(__e, V885)

tmp4415 := GConsP(__e, tmp4414)

var ifres4405 Obj

if True == tmp4415 {
tmp4411 := GTl(__e, V885)

tmp4412 := GTl(__e, tmp4411)

tmp4413 := GConsP(__e, tmp4412)

var ifres4406 Obj

if True == tmp4413 {
tmp4407 := GTl(__e, V885)

tmp4408 := GTl(__e, tmp4407)

tmp4409 := GTl(__e, tmp4408)

tmp4410 := GEq(__e, Nil, tmp4409)

ifres4406 = tmp4410


} else {
ifres4406 = False


}

ifres4405 = ifres4406


} else {
ifres4405 = False


}

ifres4404 = ifres4405


} else {
ifres4404 = False


}

ifres4403 = ifres4404


} else {
ifres4403 = False


}

if True == ifres4403 {
tmp4389 := GTl(__e, V885)

tmp4390 := GHd(__e, tmp4389)

tmp4391 := GTl(__e, V885)

tmp4392 := GHd(__e, tmp4391)

tmp4393 := GTl(__e, V885)

tmp4394 := GTl(__e, tmp4393)

tmp4395 := GHd(__e, tmp4394)

tmp4396 := Call2(__e, SymFn(symshen_4factor_1selectors), tmp4392, tmp4395)


tmp4397 := GTl(__e, V885)

tmp4398 := GHd(__e, tmp4397)

tmp4399 := Call2(__e, SymFn(symshen_4factor_1selectors), tmp4398, V886)


tmp4400 := Call2(__e, SymFn(symshen_4recursively_1factor_1selectors), tmp4396, tmp4399)


return TailCall2(__e, SymFn(symshen_4restore_1P), tmp4390, tmp4400)


} else {
tmp4401 := GCons(__e, V886, Nil)

return GCons(__e, V885, tmp4401)


}


}


panic("unreachable")
})

tmp4442 := Call2(__e, ns2_1set, symshen_4recursively_1factor_1selectors, tmp4378)


_ = tmp4442

tmp4443 := MakeNative2(func(__e *ControlFlow, V887, V888 Obj) Obj {
tmp4457 := GConsP(__e, V888)

var ifres4450 Obj

if True == tmp4457 {
tmp4455 := GTl(__e, V888)

tmp4456 := GConsP(__e, tmp4455)

var ifres4451 Obj

if True == tmp4456 {
tmp4452 := GTl(__e, V888)

tmp4453 := GTl(__e, tmp4452)

tmp4454 := GEq(__e, Nil, tmp4453)

ifres4451 = tmp4454


} else {
ifres4451 = False


}

ifres4450 = ifres4451


} else {
ifres4450 = False


}

if True == ifres4450 {
tmp4444 := GHd(__e, V888)

tmp4445 := GCons(__e, tmp4444, Nil)

tmp4446 := GCons(__e, V887, tmp4445)

tmp4447 := GCons(__e, symand, tmp4446)

tmp4448 := GTl(__e, V888)

return GCons(__e, tmp4447, tmp4448)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4restore_1P)
}


panic("unreachable")
})

tmp4458 := Call2(__e, ns2_1set, symshen_4restore_1P, tmp4443)


_ = tmp4458

tmp4459 := MakeNative3(func(__e *ControlFlow, V889, V890, V891 Obj) Obj {
tmp4474 := GConsP(__e, V891)

var ifres4467 Obj

if True == tmp4474 {
tmp4472 := GTl(__e, V891)

tmp4473 := GConsP(__e, tmp4472)

var ifres4468 Obj

if True == tmp4473 {
tmp4469 := GTl(__e, V891)

tmp4470 := GTl(__e, tmp4469)

tmp4471 := GEq(__e, Nil, tmp4470)

ifres4468 = tmp4471


} else {
ifres4468 = False


}

ifres4467 = ifres4468


} else {
ifres4467 = False


}

if True == ifres4467 {
tmp4460 := GHd(__e, V891)

tmp4461 := GCons(__e, tmp4460, Nil)

tmp4462 := GCons(__e, V890, tmp4461)

tmp4463 := GCons(__e, V889, tmp4462)

tmp4464 := GCons(__e, symlet, tmp4463)

tmp4465 := GTl(__e, V891)

return GCons(__e, tmp4464, tmp4465)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4restore_1local)
}


panic("unreachable")
})

tmp4475 := Call2(__e, ns2_1set, symshen_4restore_1local, tmp4459)


_ = tmp4475

tmp4476 := MakeNative1(func(__e *ControlFlow, V896 Obj) Obj {
tmp4481 := GConsP(__e, V896)

var ifres4478 Obj

if True == tmp4481 {
tmp4479 := GTl(__e, V896)

tmp4480 := GEq(__e, Nil, tmp4479)

ifres4478 = tmp4480


} else {
ifres4478 = False


}

if True == ifres4478 {
return True
} else {
return False
}


panic("unreachable")
})

tmp4482 := Call2(__e, ns2_1set, symshen_4bad_1pivot_2, tmp4476)


_ = tmp4482

tmp4483 := MakeNative1(func(__e *ControlFlow, V897 Obj) Obj {
tmp4584 := GConsP(__e, V897)

var ifres4498 Obj

if True == tmp4584 {
tmp4582 := GHd(__e, V897)

tmp4583 := GEq(__e, symlet, tmp4582)

var ifres4499 Obj

if True == tmp4583 {
tmp4580 := GTl(__e, V897)

tmp4581 := GConsP(__e, tmp4580)

var ifres4500 Obj

if True == tmp4581 {
tmp4577 := GTl(__e, V897)

tmp4578 := GTl(__e, tmp4577)

tmp4579 := GConsP(__e, tmp4578)

var ifres4501 Obj

if True == tmp4579 {
tmp4573 := GTl(__e, V897)

tmp4574 := GTl(__e, tmp4573)

tmp4575 := GHd(__e, tmp4574)

tmp4576 := GConsP(__e, tmp4575)

var ifres4502 Obj

if True == tmp4576 {
tmp4568 := GTl(__e, V897)

tmp4569 := GTl(__e, tmp4568)

tmp4570 := GHd(__e, tmp4569)

tmp4571 := GHd(__e, tmp4570)

tmp4572 := GEq(__e, symfreeze, tmp4571)

var ifres4503 Obj

if True == tmp4572 {
tmp4563 := GTl(__e, V897)

tmp4564 := GTl(__e, tmp4563)

tmp4565 := GHd(__e, tmp4564)

tmp4566 := GTl(__e, tmp4565)

tmp4567 := GConsP(__e, tmp4566)

var ifres4504 Obj

if True == tmp4567 {
tmp4557 := GTl(__e, V897)

tmp4558 := GTl(__e, tmp4557)

tmp4559 := GHd(__e, tmp4558)

tmp4560 := GTl(__e, tmp4559)

tmp4561 := GHd(__e, tmp4560)

tmp4562 := GConsP(__e, tmp4561)

var ifres4505 Obj

if True == tmp4562 {
tmp4550 := GTl(__e, V897)

tmp4551 := GTl(__e, tmp4550)

tmp4552 := GHd(__e, tmp4551)

tmp4553 := GTl(__e, tmp4552)

tmp4554 := GHd(__e, tmp4553)

tmp4555 := GHd(__e, tmp4554)

tmp4556 := GEq(__e, symthaw, tmp4555)

var ifres4506 Obj

if True == tmp4556 {
tmp4543 := GTl(__e, V897)

tmp4544 := GTl(__e, tmp4543)

tmp4545 := GHd(__e, tmp4544)

tmp4546 := GTl(__e, tmp4545)

tmp4547 := GHd(__e, tmp4546)

tmp4548 := GTl(__e, tmp4547)

tmp4549 := GConsP(__e, tmp4548)

var ifres4507 Obj

if True == tmp4549 {
tmp4535 := GTl(__e, V897)

tmp4536 := GTl(__e, tmp4535)

tmp4537 := GHd(__e, tmp4536)

tmp4538 := GTl(__e, tmp4537)

tmp4539 := GHd(__e, tmp4538)

tmp4540 := GTl(__e, tmp4539)

tmp4541 := GTl(__e, tmp4540)

tmp4542 := GEq(__e, Nil, tmp4541)

var ifres4508 Obj

if True == tmp4542 {
tmp4529 := GTl(__e, V897)

tmp4530 := GTl(__e, tmp4529)

tmp4531 := GHd(__e, tmp4530)

tmp4532 := GTl(__e, tmp4531)

tmp4533 := GTl(__e, tmp4532)

tmp4534 := GEq(__e, Nil, tmp4533)

var ifres4509 Obj

if True == tmp4534 {
tmp4525 := GTl(__e, V897)

tmp4526 := GTl(__e, tmp4525)

tmp4527 := GTl(__e, tmp4526)

tmp4528 := GConsP(__e, tmp4527)

var ifres4510 Obj

if True == tmp4528 {
tmp4520 := GTl(__e, V897)

tmp4521 := GTl(__e, tmp4520)

tmp4522 := GTl(__e, tmp4521)

tmp4523 := GTl(__e, tmp4522)

tmp4524 := GEq(__e, Nil, tmp4523)

var ifres4511 Obj

if True == tmp4524 {
tmp4512 := GTl(__e, V897)

tmp4513 := GTl(__e, tmp4512)

tmp4514 := GHd(__e, tmp4513)

tmp4515 := GTl(__e, tmp4514)

tmp4516 := GHd(__e, tmp4515)

tmp4517 := GTl(__e, tmp4516)

tmp4518 := GHd(__e, tmp4517)

tmp4519 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, tmp4518)

ifres4511 = tmp4519


} else {
ifres4511 = False


}

ifres4510 = ifres4511


} else {
ifres4510 = False


}

ifres4509 = ifres4510


} else {
ifres4509 = False


}

ifres4508 = ifres4509


} else {
ifres4508 = False


}

ifres4507 = ifres4508


} else {
ifres4507 = False


}

ifres4506 = ifres4507


} else {
ifres4506 = False


}

ifres4505 = ifres4506


} else {
ifres4505 = False


}

ifres4504 = ifres4505


} else {
ifres4504 = False


}

ifres4503 = ifres4504


} else {
ifres4503 = False


}

ifres4502 = ifres4503


} else {
ifres4502 = False


}

ifres4501 = ifres4502


} else {
ifres4501 = False


}

ifres4500 = ifres4501


} else {
ifres4500 = False


}

ifres4499 = ifres4500


} else {
ifres4499 = False


}

ifres4498 = ifres4499


} else {
ifres4498 = False


}

if True == ifres4498 {
tmp4484 := GTl(__e, V897)

tmp4485 := GTl(__e, tmp4484)

tmp4486 := GHd(__e, tmp4485)

tmp4487 := GTl(__e, tmp4486)

tmp4488 := GHd(__e, tmp4487)

tmp4489 := GTl(__e, tmp4488)

tmp4490 := GHd(__e, tmp4489)

tmp4491 := GTl(__e, V897)

tmp4492 := GHd(__e, tmp4491)

tmp4493 := GTl(__e, V897)

tmp4494 := GTl(__e, tmp4493)

tmp4495 := GTl(__e, tmp4494)

tmp4496 := GHd(__e, tmp4495)

return TailCall3(__e, SymFn(symsubst), tmp4490, tmp4492, tmp4496)


} else {
return V897
}


panic("unreachable")
})

tmp4585 := Call2(__e, ns2_1set, symshen_4remove_1indirection, tmp4483)


_ = tmp4585

tmp4586 := (func() Obj {
var __self Obj
__self = MakeNative3(func(__e *ControlFlow, __self1, __self2, __self3 Obj) Obj {
__selftop:
V900 := __self1
_ = V900
V901 := __self2
_ = V901
V902 := __self3
_ = V902
tmp4672 := GConsP(__e, V901)

var ifres4627 Obj

if True == tmp4672 {
tmp4670 := GHd(__e, V901)

tmp4671 := GConsP(__e, tmp4670)

var ifres4628 Obj

if True == tmp4671 {
tmp4667 := GHd(__e, V901)

tmp4668 := GHd(__e, tmp4667)

tmp4669 := GConsP(__e, tmp4668)

var ifres4629 Obj

if True == tmp4669 {
tmp4663 := GHd(__e, V901)

tmp4664 := GHd(__e, tmp4663)

tmp4665 := GHd(__e, tmp4664)

tmp4666 := GEq(__e, symand, tmp4665)

var ifres4630 Obj

if True == tmp4666 {
tmp4659 := GHd(__e, V901)

tmp4660 := GHd(__e, tmp4659)

tmp4661 := GTl(__e, tmp4660)

tmp4662 := GConsP(__e, tmp4661)

var ifres4631 Obj

if True == tmp4662 {
tmp4654 := GHd(__e, V901)

tmp4655 := GHd(__e, tmp4654)

tmp4656 := GTl(__e, tmp4655)

tmp4657 := GTl(__e, tmp4656)

tmp4658 := GConsP(__e, tmp4657)

var ifres4632 Obj

if True == tmp4658 {
tmp4648 := GHd(__e, V901)

tmp4649 := GHd(__e, tmp4648)

tmp4650 := GTl(__e, tmp4649)

tmp4651 := GTl(__e, tmp4650)

tmp4652 := GTl(__e, tmp4651)

tmp4653 := GEq(__e, Nil, tmp4652)

var ifres4633 Obj

if True == tmp4653 {
tmp4645 := GHd(__e, V901)

tmp4646 := GTl(__e, tmp4645)

tmp4647 := GConsP(__e, tmp4646)

var ifres4634 Obj

if True == tmp4647 {
tmp4641 := GHd(__e, V901)

tmp4642 := GTl(__e, tmp4641)

tmp4643 := GTl(__e, tmp4642)

tmp4644 := GEq(__e, Nil, tmp4643)

var ifres4635 Obj

if True == tmp4644 {
tmp4636 := GHd(__e, V901)

tmp4637 := GHd(__e, tmp4636)

tmp4638 := GTl(__e, tmp4637)

tmp4639 := GHd(__e, tmp4638)

tmp4640 := GEq(__e, V900, tmp4639)

ifres4635 = tmp4640


} else {
ifres4635 = False


}

ifres4634 = ifres4635


} else {
ifres4634 = False


}

ifres4633 = ifres4634


} else {
ifres4633 = False


}

ifres4632 = ifres4633


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

ifres4628 = ifres4629


} else {
ifres4628 = False


}

ifres4627 = ifres4628


} else {
ifres4627 = False


}

if True == ifres4627 {
tmp4587 := GHd(__e, V901)

tmp4588 := GHd(__e, tmp4587)

tmp4589 := GTl(__e, tmp4588)

tmp4590 := GHd(__e, tmp4589)

tmp4591 := GTl(__e, V901)

tmp4592 := GHd(__e, V901)

tmp4593 := GHd(__e, tmp4592)

tmp4594 := GTl(__e, tmp4593)

tmp4595 := GTl(__e, tmp4594)

tmp4596 := GHd(__e, tmp4595)

tmp4597 := GHd(__e, V901)

tmp4598 := GTl(__e, tmp4597)

tmp4599 := GCons(__e, tmp4596, tmp4598)

tmp4600 := GCons(__e, tmp4599, V902)

if SymFn(symshen_4pivot_1on) == __self {
__self1, __self2, __self3 = tmp4590, tmp4591, tmp4600
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4pivot_1on), tmp4590, tmp4591, tmp4600)


} else {
tmp4625 := GConsP(__e, V901)

var ifres4609 Obj

if True == tmp4625 {
tmp4623 := GHd(__e, V901)

tmp4624 := GConsP(__e, tmp4623)

var ifres4610 Obj

if True == tmp4624 {
tmp4620 := GHd(__e, V901)

tmp4621 := GTl(__e, tmp4620)

tmp4622 := GConsP(__e, tmp4621)

var ifres4611 Obj

if True == tmp4622 {
tmp4616 := GHd(__e, V901)

tmp4617 := GTl(__e, tmp4616)

tmp4618 := GTl(__e, tmp4617)

tmp4619 := GEq(__e, Nil, tmp4618)

var ifres4612 Obj

if True == tmp4619 {
tmp4613 := GHd(__e, V901)

tmp4614 := GHd(__e, tmp4613)

tmp4615 := GEq(__e, V900, tmp4614)

ifres4612 = tmp4615


} else {
ifres4612 = False


}

ifres4611 = ifres4612


} else {
ifres4611 = False


}

ifres4610 = ifres4611


} else {
ifres4610 = False


}

ifres4609 = ifres4610


} else {
ifres4609 = False


}

if True == ifres4609 {
tmp4601 := GHd(__e, V901)

tmp4602 := GHd(__e, tmp4601)

tmp4603 := GTl(__e, V901)

tmp4604 := GHd(__e, V901)

tmp4605 := GTl(__e, tmp4604)

tmp4606 := GCons(__e, True, tmp4605)

tmp4607 := GCons(__e, tmp4606, V902)

if SymFn(symshen_4pivot_1on) == __self {
__self1, __self2, __self3 = tmp4602, tmp4603, tmp4607
__e.Tick()
goto __selftop
}
return TailCall3(__e, SymFn(symshen_4pivot_1on), tmp4602, tmp4603, tmp4607)


} else {
return TailCall2(__e, SymFn(sym_8p), V902, V901)
}


}


panic("unreachable")
})
return __self
})()

tmp4673 := Call2(__e, ns2_1set, symshen_4pivot_1on, tmp4586)


_ = tmp4673

tmp4674 := MakeNative2(func(__e *ControlFlow, V905, V906 Obj) Obj {
tmp4695 := GConsP(__e, V905)

var ifres4688 Obj

if True == tmp4695 {
tmp4693 := GTl(__e, V905)

tmp4694 := GConsP(__e, tmp4693)

var ifres4689 Obj

if True == tmp4694 {
tmp4690 := GTl(__e, V905)

tmp4691 := GTl(__e, tmp4690)

tmp4692 := GEq(__e, Nil, tmp4691)

ifres4689 = tmp4692


} else {
ifres4689 = False


}

ifres4688 = ifres4689


} else {
ifres4688 = False


}

if True == ifres4688 {
tmp4675 := GHd(__e, V905)

tmp4676 := Call1(__e, SymFn(symshen_4op), tmp4675)


let__2644 := tmp4676
_ = let__2644

tmp4686 := GEq(__e, symshen_4skip, let__2644)

if True == tmp4686 {
return V906
} else {
tmp4677 := Call1(__e, SymFn(symshen_4op1), let__2644)


tmp4678 := GTl(__e, V905)

tmp4679 := GCons(__e, tmp4677, tmp4678)

tmp4680 := Call1(__e, SymFn(symshen_4op2), let__2644)


tmp4681 := GTl(__e, V905)

tmp4682 := GCons(__e, tmp4680, tmp4681)

tmp4683 := GCons(__e, tmp4682, Nil)

tmp4684 := GCons(__e, tmp4679, tmp4683)

return TailCall2(__e, SymFn(symshen_4factor_1selectors_1h), tmp4684, V906)


}


} else {
return V906
}


panic("unreachable")
})

tmp4696 := Call2(__e, ns2_1set, symshen_4factor_1selectors, tmp4674)


_ = tmp4696

tmp4697 := MakeNative1(func(__e *ControlFlow, V910 Obj) Obj {
tmp4705 := GEq(__e, symcons_2, V910)

if True == tmp4705 {
return symcons
} else {
tmp4703 := GEq(__e, symshen_4_7string_2, V910)

if True == tmp4703 {
return sym_8s
} else {
tmp4701 := GEq(__e, symshen_4_7vector_2, V910)

if True == tmp4701 {
return sym_8v
} else {
tmp4699 := GEq(__e, symtuple_2, V910)

if True == tmp4699 {
return sym_8p
} else {
return symshen_4skip
}


}


}


}


panic("unreachable")
})

tmp4706 := Call2(__e, ns2_1set, symshen_4op, tmp4697)


_ = tmp4706

tmp4707 := (func() Obj {
var __self Obj
__self = MakeNative2(func(__e *ControlFlow, __self1, __self2 Obj) Obj {
__selftop:
V911 := __self1
_ = V911
V912 := __self2
_ = V912
tmp4725 := GEq(__e, Nil, V911)

if True == tmp4725 {
return V912
} else {
tmp4723 := GConsP(__e, V911)

if True == tmp4723 {
tmp4719 := GHd(__e, V911)

tmp4720 := Call2(__e, SymFn(symoccurrences), tmp4719, V912)


if True == GGt(__e, tmp4720, MakeInteger(1)) {
tmp4708 := Call1(__e, SymFn(symgensym), symSelect)


let__2645 := tmp4708
_ = let__2645

tmp4709 := GHd(__e, V911)

tmp4710 := GTl(__e, V911)

tmp4711 := GHd(__e, V911)

tmp4712 := Call3(__e, SymFn(symsubst), let__2645, tmp4711, V912)


tmp4713 := Call2(__e, SymFn(symshen_4factor_1selectors_1h), tmp4710, tmp4712)


tmp4714 := GCons(__e, tmp4713, Nil)

tmp4715 := GCons(__e, tmp4709, tmp4714)

tmp4716 := GCons(__e, let__2645, tmp4715)

return GCons(__e, symlet, tmp4716)


} else {
tmp4717 := GTl(__e, V911)

if SymFn(symshen_4factor_1selectors_1h) == __self {
__self1, __self2 = tmp4717, V912
__e.Tick()
goto __selftop
}
return TailCall2(__e, SymFn(symshen_4factor_1selectors_1h), tmp4717, V912)


}


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4factor_1selectors_1h)
}


}


panic("unreachable")
})
return __self
})()

return TailCall2(__e, ns2_1set, symshen_4factor_1selectors_1h, tmp4707)




panic("unreachable")
})

