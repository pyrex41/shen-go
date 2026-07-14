package main

import . "github.com/tiancaiamao/shen-go/kl"

var CoreMain = MakeNative(func(__e *ControlFlow) {
tmp2605 := MakeNative(func(__e *ControlFlow) {
V521 := __e.Get(1)
_ = V521
tmp2606 := MakeNative(func(__e *ControlFlow) {
W522 := __e.Get(1)
_ = W522
__e.TailApply(PrimFunc(symshen_4record_1and_1evaluate), W522)
return
}, 1)

tmp2607 := Call(__e, PrimFunc(symshen_4shen_1_6kl_1h), V521)


__e.TailApply(tmp2606, tmp2607)
return


}, 1)

tmp2608 := Call(__e, ns2_1set, symshen_4shen_1_6kl, tmp2605)


_ = tmp2608

tmp2609 := MakeNative(func(__e *ControlFlow) {
V523 := __e.Get(1)
_ = V523
tmp2662 := PrimIsPair(V523)

var ifres2636 Obj

if True == tmp2662 {
tmp2660 := PrimHead(V523)

tmp2661 := PrimEqual(symdefun, tmp2660)

var ifres2638 Obj

if True == tmp2661 {
tmp2658 := PrimTail(V523)

tmp2659 := PrimIsPair(tmp2658)

var ifres2640 Obj

if True == tmp2659 {
tmp2655 := PrimTail(V523)

tmp2656 := PrimTail(tmp2655)

tmp2657 := PrimIsPair(tmp2656)

var ifres2642 Obj

if True == tmp2657 {
tmp2651 := PrimTail(V523)

tmp2652 := PrimTail(tmp2651)

tmp2653 := PrimTail(tmp2652)

tmp2654 := PrimIsPair(tmp2653)

var ifres2644 Obj

if True == tmp2654 {
tmp2646 := PrimTail(V523)

tmp2647 := PrimTail(tmp2646)

tmp2648 := PrimTail(tmp2647)

tmp2649 := PrimTail(tmp2648)

tmp2650 := PrimEqual(Nil, tmp2649)

var ifres2645 Obj

if True == tmp2650 {
ifres2645 = True


} else {
ifres2645 = False


}

ifres2644 = ifres2645


} else {
ifres2644 = False


}

var ifres2643 Obj

if True == ifres2644 {
ifres2643 = True


} else {
ifres2643 = False


}

ifres2642 = ifres2643


} else {
ifres2642 = False


}

var ifres2641 Obj

if True == ifres2642 {
ifres2641 = True


} else {
ifres2641 = False


}

ifres2640 = ifres2641


} else {
ifres2640 = False


}

var ifres2639 Obj

if True == ifres2640 {
ifres2639 = True


} else {
ifres2639 = False


}

ifres2638 = ifres2639


} else {
ifres2638 = False


}

var ifres2637 Obj

if True == ifres2638 {
ifres2637 = True


} else {
ifres2637 = False


}

ifres2636 = ifres2637


} else {
ifres2636 = False


}

if True == ifres2636 {
tmp2610 := MakeNative(func(__e *ControlFlow) {
W524 := __e.Get(1)
_ = W524
tmp2611 := MakeNative(func(__e *ControlFlow) {
W525 := __e.Get(1)
_ = W525
tmp2612 := MakeNative(func(__e *ControlFlow) {
W526 := __e.Get(1)
_ = W526
tmp2613 := MakeNative(func(__e *ControlFlow) {
W527 := __e.Get(1)
_ = W527
tmp2614 := PrimTail(V523)

tmp2615 := PrimHead(tmp2614)

__e.TailApply(PrimFunc(symshen_4fn_1print), tmp2615)
return


}, 1)

tmp2616 := Call(__e, PrimFunc(symeval_1kl), V523)


__e.TailApply(tmp2613, tmp2616)
return


}, 1)

tmp2617 := PrimTail(V523)

tmp2618 := PrimHead(tmp2617)

tmp2619 := Call(__e, PrimFunc(symshen_4record_1kl), tmp2618, V523)


__e.TailApply(tmp2612, tmp2619)
return


}, 1)

tmp2620 := PrimTail(V523)

tmp2621 := PrimHead(tmp2620)

tmp2622 := PrimTail(V523)

tmp2623 := PrimTail(tmp2622)

tmp2624 := PrimHead(tmp2623)

tmp2625 := Call(__e, PrimFunc(symlength), tmp2624)


tmp2626 := Call(__e, PrimFunc(symshen_4store_1arity), tmp2621, tmp2625)


__e.TailApply(tmp2611, tmp2626)
return


}, 1)

tmp2632 := PrimTail(V523)

tmp2633 := PrimHead(tmp2632)

tmp2634 := Call(__e, PrimFunc(symshen_4sysfunc_2), tmp2633)


var ifres2627 Obj

if True == tmp2634 {
tmp2628 := PrimTail(V523)

tmp2629 := PrimHead(tmp2628)

tmp2630 := Call(__e, PrimFunc(symshen_4app), tmp2629, MakeString(" is not a legitimate function name\n"), symshen_4a)


tmp2631 := PrimSimpleError(tmp2630)

ifres2627 = tmp2631


} else {
ifres2627 = symshen_4skip


}

__e.TailApply(tmp2610, ifres2627)
return


} else {
__e.Return(V523)
return
}


}, 1)

tmp2663 := Call(__e, ns2_1set, symshen_4record_1and_1evaluate, tmp2609)


_ = tmp2663

tmp2664 := MakeNative(func(__e *ControlFlow) {
V528 := __e.Get(1)
_ = V528
tmp2765 := PrimIsPair(V528)

var ifres2757 Obj

if True == tmp2765 {
tmp2763 := PrimHead(V528)

tmp2764 := PrimEqual(symdefine, tmp2763)

var ifres2759 Obj

if True == tmp2764 {
tmp2761 := PrimTail(V528)

tmp2762 := PrimIsPair(tmp2761)

var ifres2760 Obj

if True == tmp2762 {
ifres2760 = True


} else {
ifres2760 = False


}

ifres2759 = ifres2760


} else {
ifres2759 = False


}

var ifres2758 Obj

if True == ifres2759 {
ifres2758 = True


} else {
ifres2758 = False


}

ifres2757 = ifres2758


} else {
ifres2757 = False


}

if True == ifres2757 {
tmp2665 := PrimTail(V528)

tmp2666 := PrimHead(tmp2665)

tmp2667 := PrimTail(V528)

tmp2668 := PrimTail(tmp2667)

__e.TailApply(PrimFunc(symshen_4shendef_1_6kldef), tmp2666, tmp2668)
return


} else {
tmp2755 := PrimIsPair(V528)

var ifres2729 Obj

if True == tmp2755 {
tmp2753 := PrimHead(V528)

tmp2754 := PrimEqual(symdefun, tmp2753)

var ifres2731 Obj

if True == tmp2754 {
tmp2751 := PrimTail(V528)

tmp2752 := PrimIsPair(tmp2751)

var ifres2733 Obj

if True == tmp2752 {
tmp2748 := PrimTail(V528)

tmp2749 := PrimTail(tmp2748)

tmp2750 := PrimIsPair(tmp2749)

var ifres2735 Obj

if True == tmp2750 {
tmp2744 := PrimTail(V528)

tmp2745 := PrimTail(tmp2744)

tmp2746 := PrimTail(tmp2745)

tmp2747 := PrimIsPair(tmp2746)

var ifres2737 Obj

if True == tmp2747 {
tmp2739 := PrimTail(V528)

tmp2740 := PrimTail(tmp2739)

tmp2741 := PrimTail(tmp2740)

tmp2742 := PrimTail(tmp2741)

tmp2743 := PrimEqual(Nil, tmp2742)

var ifres2738 Obj

if True == tmp2743 {
ifres2738 = True


} else {
ifres2738 = False


}

ifres2737 = ifres2738


} else {
ifres2737 = False


}

var ifres2736 Obj

if True == ifres2737 {
ifres2736 = True


} else {
ifres2736 = False


}

ifres2735 = ifres2736


} else {
ifres2735 = False


}

var ifres2734 Obj

if True == ifres2735 {
ifres2734 = True


} else {
ifres2734 = False


}

ifres2733 = ifres2734


} else {
ifres2733 = False


}

var ifres2732 Obj

if True == ifres2733 {
ifres2732 = True


} else {
ifres2732 = False


}

ifres2731 = ifres2732


} else {
ifres2731 = False


}

var ifres2730 Obj

if True == ifres2731 {
ifres2730 = True


} else {
ifres2730 = False


}

ifres2729 = ifres2730


} else {
ifres2729 = False


}

if True == ifres2729 {
__e.Return(V528)
return
} else {
tmp2727 := PrimIsPair(V528)

var ifres2708 Obj

if True == tmp2727 {
tmp2725 := PrimHead(V528)

tmp2726 := PrimEqual(symtype, tmp2725)

var ifres2710 Obj

if True == tmp2726 {
tmp2723 := PrimTail(V528)

tmp2724 := PrimIsPair(tmp2723)

var ifres2712 Obj

if True == tmp2724 {
tmp2720 := PrimTail(V528)

tmp2721 := PrimTail(tmp2720)

tmp2722 := PrimIsPair(tmp2721)

var ifres2714 Obj

if True == tmp2722 {
tmp2716 := PrimTail(V528)

tmp2717 := PrimTail(tmp2716)

tmp2718 := PrimTail(tmp2717)

tmp2719 := PrimEqual(Nil, tmp2718)

var ifres2715 Obj

if True == tmp2719 {
ifres2715 = True


} else {
ifres2715 = False


}

ifres2714 = ifres2715


} else {
ifres2714 = False


}

var ifres2713 Obj

if True == ifres2714 {
ifres2713 = True


} else {
ifres2713 = False


}

ifres2712 = ifres2713


} else {
ifres2712 = False


}

var ifres2711 Obj

if True == ifres2712 {
ifres2711 = True


} else {
ifres2711 = False


}

ifres2710 = ifres2711


} else {
ifres2710 = False


}

var ifres2709 Obj

if True == ifres2710 {
ifres2709 = True


} else {
ifres2709 = False


}

ifres2708 = ifres2709


} else {
ifres2708 = False


}

if True == ifres2708 {
tmp2669 := PrimTail(V528)

tmp2670 := PrimHead(tmp2669)

tmp2671 := PrimTail(V528)

tmp2672 := PrimTail(tmp2671)

tmp2673 := PrimHead(tmp2672)

tmp2674 := Call(__e, PrimFunc(symshen_4rcons__form), tmp2673)


tmp2675 := PrimCons(tmp2674, Nil)

tmp2676 := PrimCons(tmp2670, tmp2675)

__e.Return(PrimCons(symtype, tmp2676))
return


} else {
tmp2706 := PrimIsPair(V528)

var ifres2687 Obj

if True == tmp2706 {
tmp2704 := PrimHead(V528)

tmp2705 := PrimEqual(syminput_7, tmp2704)

var ifres2689 Obj

if True == tmp2705 {
tmp2702 := PrimTail(V528)

tmp2703 := PrimIsPair(tmp2702)

var ifres2691 Obj

if True == tmp2703 {
tmp2699 := PrimTail(V528)

tmp2700 := PrimTail(tmp2699)

tmp2701 := PrimIsPair(tmp2700)

var ifres2693 Obj

if True == tmp2701 {
tmp2695 := PrimTail(V528)

tmp2696 := PrimTail(tmp2695)

tmp2697 := PrimTail(tmp2696)

tmp2698 := PrimEqual(Nil, tmp2697)

var ifres2694 Obj

if True == tmp2698 {
ifres2694 = True


} else {
ifres2694 = False


}

ifres2693 = ifres2694


} else {
ifres2693 = False


}

var ifres2692 Obj

if True == ifres2693 {
ifres2692 = True


} else {
ifres2692 = False


}

ifres2691 = ifres2692


} else {
ifres2691 = False


}

var ifres2690 Obj

if True == ifres2691 {
ifres2690 = True


} else {
ifres2690 = False


}

ifres2689 = ifres2690


} else {
ifres2689 = False


}

var ifres2688 Obj

if True == ifres2689 {
ifres2688 = True


} else {
ifres2688 = False


}

ifres2687 = ifres2688


} else {
ifres2687 = False


}

if True == ifres2687 {
tmp2677 := PrimTail(V528)

tmp2678 := PrimHead(tmp2677)

tmp2679 := Call(__e, PrimFunc(symshen_4rcons__form), tmp2678)


tmp2680 := PrimTail(V528)

tmp2681 := PrimTail(tmp2680)

tmp2682 := PrimCons(tmp2679, tmp2681)

__e.Return(PrimCons(syminput_7, tmp2682))
return


} else {
tmp2685 := PrimIsPair(V528)

if True == tmp2685 {
tmp2683 := MakeNative(func(__e *ControlFlow) {
Z529 := __e.Get(1)
_ = Z529
__e.TailApply(PrimFunc(symshen_4shen_1_6kl_1h), Z529)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp2683, V528)
return


} else {
__e.Return(V528)
return
}


}


}


}


}


}, 1)

tmp2766 := Call(__e, ns2_1set, symshen_4shen_1_6kl_1h, tmp2664)


_ = tmp2766

tmp2767 := MakeNative(func(__e *ControlFlow) {
V530 := __e.Get(1)
_ = V530
V531 := __e.Get(2)
_ = V531
tmp2768 := MakeNative(func(__e *ControlFlow) {
Z532 := __e.Get(1)
_ = Z532
__e.TailApply(PrimFunc(symshen_4_5define_6), Z532)
return
}, 1)

tmp2769 := PrimCons(V530, V531)

__e.TailApply(PrimFunc(symcompile), tmp2768, tmp2769)
return


}, 2)

tmp2770 := Call(__e, ns2_1set, symshen_4shendef_1_6kldef, tmp2767)


_ = tmp2770

tmp2771 := MakeNative(func(__e *ControlFlow) {
V533 := __e.Get(1)
_ = V533
tmp2772 := MakeNative(func(__e *ControlFlow) {
W534 := __e.Get(1)
_ = W534
tmp2795 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W534)


if True == tmp2795 {
tmp2773 := MakeNative(func(__e *ControlFlow) {
W545 := __e.Get(1)
_ = W545
tmp2775 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W545)


if True == tmp2775 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W545)
return
}


}, 1)

tmp2776 := MakeNative(func(__e *ControlFlow) {
W546 := __e.Get(1)
_ = W546
tmp2791 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W546)


if True == tmp2791 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2777 := MakeNative(func(__e *ControlFlow) {
W547 := __e.Get(1)
_ = W547
tmp2778 := MakeNative(func(__e *ControlFlow) {
W548 := __e.Get(1)
_ = W548
tmp2779 := MakeNative(func(__e *ControlFlow) {
W549 := __e.Get(1)
_ = W549
tmp2786 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W549)


if True == tmp2786 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2780 := MakeNative(func(__e *ControlFlow) {
W550 := __e.Get(1)
_ = W550
tmp2781 := MakeNative(func(__e *ControlFlow) {
W551 := __e.Get(1)
_ = W551
tmp2782 := Call(__e, PrimFunc(symshen_4shendef_1_6kldef_1h), W547, W550)


__e.TailApply(PrimFunc(symshen_4comb), W551, tmp2782)
return


}, 1)

tmp2783 := Call(__e, PrimFunc(symshen_4in_1_6), W549)


__e.TailApply(tmp2781, tmp2783)
return


}, 1)

tmp2784 := Call(__e, PrimFunc(symshen_4_5_1out), W549)


__e.TailApply(tmp2780, tmp2784)
return


}


}, 1)

tmp2787 := Call(__e, PrimFunc(symshen_4_5rules_6), W548)


__e.TailApply(tmp2779, tmp2787)
return


}, 1)

tmp2788 := Call(__e, PrimFunc(symshen_4in_1_6), W546)


__e.TailApply(tmp2778, tmp2788)
return


}, 1)

tmp2789 := Call(__e, PrimFunc(symshen_4_5_1out), W546)


__e.TailApply(tmp2777, tmp2789)
return


}


}, 1)

tmp2792 := Call(__e, PrimFunc(symshen_4_5name_6), V533)


tmp2793 := Call(__e, tmp2776, tmp2792)


__e.TailApply(tmp2773, tmp2793)
return


} else {
__e.Return(W534)
return
}


}, 1)

tmp2796 := MakeNative(func(__e *ControlFlow) {
W535 := __e.Get(1)
_ = W535
tmp2825 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W535)


if True == tmp2825 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2797 := MakeNative(func(__e *ControlFlow) {
W536 := __e.Get(1)
_ = W536
tmp2798 := MakeNative(func(__e *ControlFlow) {
W537 := __e.Get(1)
_ = W537
tmp2821 := Call(__e, PrimFunc(symshen_4hds_a_2), W537, sym_i)


if True == tmp2821 {
tmp2799 := MakeNative(func(__e *ControlFlow) {
W538 := __e.Get(1)
_ = W538
tmp2800 := MakeNative(func(__e *ControlFlow) {
W539 := __e.Get(1)
_ = W539
tmp2817 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W539)


if True == tmp2817 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2801 := MakeNative(func(__e *ControlFlow) {
W540 := __e.Get(1)
_ = W540
tmp2814 := Call(__e, PrimFunc(symshen_4hds_a_2), W540, sym_j)


if True == tmp2814 {
tmp2802 := MakeNative(func(__e *ControlFlow) {
W541 := __e.Get(1)
_ = W541
tmp2803 := MakeNative(func(__e *ControlFlow) {
W542 := __e.Get(1)
_ = W542
tmp2810 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W542)


if True == tmp2810 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2804 := MakeNative(func(__e *ControlFlow) {
W543 := __e.Get(1)
_ = W543
tmp2805 := MakeNative(func(__e *ControlFlow) {
W544 := __e.Get(1)
_ = W544
tmp2806 := Call(__e, PrimFunc(symshen_4shendef_1_6kldef_1h), W536, W543)


__e.TailApply(PrimFunc(symshen_4comb), W544, tmp2806)
return


}, 1)

tmp2807 := Call(__e, PrimFunc(symshen_4in_1_6), W542)


__e.TailApply(tmp2805, tmp2807)
return


}, 1)

tmp2808 := Call(__e, PrimFunc(symshen_4_5_1out), W542)


__e.TailApply(tmp2804, tmp2808)
return


}


}, 1)

tmp2811 := Call(__e, PrimFunc(symshen_4_5rules_6), W541)


__e.TailApply(tmp2803, tmp2811)
return


}, 1)

tmp2812 := Call(__e, PrimFunc(symtail), W540)


__e.TailApply(tmp2802, tmp2812)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2815 := Call(__e, PrimFunc(symshen_4in_1_6), W539)


__e.TailApply(tmp2801, tmp2815)
return


}


}, 1)

tmp2818 := Call(__e, PrimFunc(symshen_4_5signature_6), W538)


__e.TailApply(tmp2800, tmp2818)
return


}, 1)

tmp2819 := Call(__e, PrimFunc(symtail), W537)


__e.TailApply(tmp2799, tmp2819)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2822 := Call(__e, PrimFunc(symshen_4in_1_6), W535)


__e.TailApply(tmp2798, tmp2822)
return


}, 1)

tmp2823 := Call(__e, PrimFunc(symshen_4_5_1out), W535)


__e.TailApply(tmp2797, tmp2823)
return


}


}, 1)

tmp2826 := Call(__e, PrimFunc(symshen_4_5name_6), V533)


tmp2827 := Call(__e, tmp2796, tmp2826)


__e.TailApply(tmp2772, tmp2827)
return


}, 1)

tmp2828 := Call(__e, ns2_1set, symshen_4_5define_6, tmp2771)


_ = tmp2828

tmp2829 := MakeNative(func(__e *ControlFlow) {
V552 := __e.Get(1)
_ = V552
V553 := __e.Get(2)
_ = V553
tmp2830 := MakeNative(func(__e *ControlFlow) {
W554 := __e.Get(1)
_ = W554
tmp2831 := MakeNative(func(__e *ControlFlow) {
W556 := __e.Get(1)
_ = W556
tmp2832 := MakeNative(func(__e *ControlFlow) {
W557 := __e.Get(1)
_ = W557
tmp2833 := MakeNative(func(__e *ControlFlow) {
W559 := __e.Get(1)
_ = W559
tmp2834 := MakeNative(func(__e *ControlFlow) {
W560 := __e.Get(1)
_ = W560
__e.Return(W560)
return
}, 1)

tmp2835 := Call(__e, PrimFunc(symshen_4compile_1to_1kl), V552, W559, W556)


tmp2836 := Call(__e, PrimFunc(symshen_4factorise_1code), tmp2835)


__e.TailApply(tmp2834, tmp2836)
return


}, 1)

tmp2837 := Call(__e, PrimFunc(symshen_4unprotect), V553)


__e.TailApply(tmp2833, tmp2837)
return


}, 1)

tmp2838 := MakeNative(func(__e *ControlFlow) {
Z558 := __e.Get(1)
_ = Z558
__e.TailApply(PrimFunc(symshen_4free_1var_1chk), V552, Z558)
return
}, 1)

tmp2839 := Call(__e, PrimFunc(symmap), tmp2838, V553)


__e.TailApply(tmp2832, tmp2839)
return


}, 1)

tmp2840 := Call(__e, PrimFunc(symshen_4arity_1chk), V552, W554)


__e.TailApply(tmp2831, tmp2840)
return


}, 1)

tmp2841 := MakeNative(func(__e *ControlFlow) {
Z555 := __e.Get(1)
_ = Z555
__e.TailApply(PrimFunc(symfst), Z555)
return
}, 1)

tmp2842 := Call(__e, PrimFunc(symmap), tmp2841, V553)


__e.TailApply(tmp2830, tmp2842)
return


}, 2)

tmp2843 := Call(__e, ns2_1set, symshen_4shendef_1_6kldef_1h, tmp2829)


_ = tmp2843

tmp2844 := MakeNative(func(__e *ControlFlow) {
V561 := __e.Get(1)
_ = V561
tmp2870 := Call(__e, PrimFunc(symtuple_2), V561)


if True == tmp2870 {
tmp2845 := Call(__e, PrimFunc(symfst), V561)


tmp2846 := Call(__e, PrimFunc(symshen_4unprotect), tmp2845)


tmp2847 := Call(__e, PrimFunc(symsnd), V561)


tmp2848 := Call(__e, PrimFunc(symshen_4unprotect), tmp2847)


__e.TailApply(PrimFunc(sym_8p), tmp2846, tmp2848)
return


} else {
tmp2868 := PrimIsPair(V561)

var ifres2855 Obj

if True == tmp2868 {
tmp2866 := PrimHead(V561)

tmp2867 := PrimEqual(symprotect, tmp2866)

var ifres2857 Obj

if True == tmp2867 {
tmp2864 := PrimTail(V561)

tmp2865 := PrimIsPair(tmp2864)

var ifres2859 Obj

if True == tmp2865 {
tmp2861 := PrimTail(V561)

tmp2862 := PrimTail(tmp2861)

tmp2863 := PrimEqual(Nil, tmp2862)

var ifres2860 Obj

if True == tmp2863 {
ifres2860 = True


} else {
ifres2860 = False


}

ifres2859 = ifres2860


} else {
ifres2859 = False


}

var ifres2858 Obj

if True == ifres2859 {
ifres2858 = True


} else {
ifres2858 = False


}

ifres2857 = ifres2858


} else {
ifres2857 = False


}

var ifres2856 Obj

if True == ifres2857 {
ifres2856 = True


} else {
ifres2856 = False


}

ifres2855 = ifres2856


} else {
ifres2855 = False


}

if True == ifres2855 {
tmp2849 := PrimTail(V561)

tmp2850 := PrimHead(tmp2849)

__e.TailApply(PrimFunc(symshen_4unprotect), tmp2850)
return


} else {
tmp2853 := PrimIsPair(V561)

if True == tmp2853 {
tmp2851 := MakeNative(func(__e *ControlFlow) {
Z562 := __e.Get(1)
_ = Z562
__e.TailApply(PrimFunc(symshen_4unprotect), Z562)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp2851, V561)
return


} else {
__e.Return(V561)
return
}


}


}


}, 1)

tmp2871 := Call(__e, ns2_1set, symshen_4unprotect, tmp2844)


_ = tmp2871

tmp2872 := MakeNative(func(__e *ControlFlow) {
V563 := __e.Get(1)
_ = V563
tmp2873 := MakeNative(func(__e *ControlFlow) {
W564 := __e.Get(1)
_ = W564
tmp2875 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W564)


if True == tmp2875 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W564)
return
}


}, 1)

tmp2891 := PrimIsPair(V563)

var ifres2876 Obj

if True == tmp2891 {
tmp2877 := MakeNative(func(__e *ControlFlow) {
W565 := __e.Get(1)
_ = W565
tmp2878 := MakeNative(func(__e *ControlFlow) {
W566 := __e.Get(1)
_ = W566
tmp2886 := PrimIsSymbol(W565)

var ifres2882 Obj

if True == tmp2886 {
tmp2884 := PrimIsVariable(W565)

tmp2885 := PrimNot(tmp2884)

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

var ifres2879 Obj

if True == ifres2882 {
ifres2879 = W565


} else {
tmp2880 := Call(__e, PrimFunc(symshen_4app), W565, MakeString(" is not a legitimate function name.\n"), symshen_4a)


tmp2881 := PrimSimpleError(tmp2880)

ifres2879 = tmp2881


}

__e.TailApply(PrimFunc(symshen_4comb), W566, ifres2879)
return


}, 1)

tmp2887 := Call(__e, PrimFunc(symtail), V563)


__e.TailApply(tmp2878, tmp2887)
return


}, 1)

tmp2888 := Call(__e, PrimFunc(symhead), V563)


tmp2889 := Call(__e, tmp2877, tmp2888)


ifres2876 = tmp2889


} else {
tmp2890 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2876 = tmp2890


}

__e.TailApply(tmp2873, ifres2876)
return


}, 1)

tmp2892 := Call(__e, ns2_1set, symshen_4_5name_6, tmp2872)


_ = tmp2892

tmp2893 := MakeNative(func(__e *ControlFlow) {
V567 := __e.Get(1)
_ = V567
tmp2894 := MakeNative(func(__e *ControlFlow) {
W568 := __e.Get(1)
_ = W568
tmp2906 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W568)


if True == tmp2906 {
tmp2895 := MakeNative(func(__e *ControlFlow) {
W574 := __e.Get(1)
_ = W574
tmp2897 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W574)


if True == tmp2897 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W574)
return
}


}, 1)

tmp2898 := MakeNative(func(__e *ControlFlow) {
W575 := __e.Get(1)
_ = W575
tmp2902 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W575)


if True == tmp2902 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2899 := MakeNative(func(__e *ControlFlow) {
W576 := __e.Get(1)
_ = W576
__e.TailApply(PrimFunc(symshen_4comb), W576, Nil)
return
}, 1)

tmp2900 := Call(__e, PrimFunc(symshen_4in_1_6), W575)


__e.TailApply(tmp2899, tmp2900)
return


}


}, 1)

tmp2903 := Call(__e, PrimFunc(sym_5e_6), V567)


tmp2904 := Call(__e, tmp2898, tmp2903)


__e.TailApply(tmp2895, tmp2904)
return


} else {
__e.Return(W568)
return
}


}, 1)

tmp2928 := PrimIsPair(V567)

var ifres2907 Obj

if True == tmp2928 {
tmp2908 := MakeNative(func(__e *ControlFlow) {
W569 := __e.Get(1)
_ = W569
tmp2909 := MakeNative(func(__e *ControlFlow) {
W570 := __e.Get(1)
_ = W570
tmp2910 := MakeNative(func(__e *ControlFlow) {
W571 := __e.Get(1)
_ = W571
tmp2922 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W571)


if True == tmp2922 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2911 := MakeNative(func(__e *ControlFlow) {
W572 := __e.Get(1)
_ = W572
tmp2912 := MakeNative(func(__e *ControlFlow) {
W573 := __e.Get(1)
_ = W573
tmp2915 := PrimCons(sym_j, Nil)

tmp2916 := PrimCons(sym_i, tmp2915)

tmp2917 := Call(__e, PrimFunc(symelement_2), W569, tmp2916)


tmp2918 := PrimNot(tmp2917)

if True == tmp2918 {
tmp2913 := PrimCons(W569, W572)

__e.TailApply(PrimFunc(symshen_4comb), W573, tmp2913)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2919 := Call(__e, PrimFunc(symshen_4in_1_6), W571)


__e.TailApply(tmp2912, tmp2919)
return


}, 1)

tmp2920 := Call(__e, PrimFunc(symshen_4_5_1out), W571)


__e.TailApply(tmp2911, tmp2920)
return


}


}, 1)

tmp2923 := Call(__e, PrimFunc(symshen_4_5signature_6), W570)


__e.TailApply(tmp2910, tmp2923)
return


}, 1)

tmp2924 := Call(__e, PrimFunc(symtail), V567)


__e.TailApply(tmp2909, tmp2924)
return


}, 1)

tmp2925 := Call(__e, PrimFunc(symhead), V567)


tmp2926 := Call(__e, tmp2908, tmp2925)


ifres2907 = tmp2926


} else {
tmp2927 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2907 = tmp2927


}

__e.TailApply(tmp2894, ifres2907)
return


}, 1)

tmp2929 := Call(__e, ns2_1set, symshen_4_5signature_6, tmp2893)


_ = tmp2929

tmp2930 := MakeNative(func(__e *ControlFlow) {
V577 := __e.Get(1)
_ = V577
tmp2931 := MakeNative(func(__e *ControlFlow) {
W578 := __e.Get(1)
_ = W578
tmp2950 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W578)


if True == tmp2950 {
tmp2932 := MakeNative(func(__e *ControlFlow) {
W585 := __e.Get(1)
_ = W585
tmp2934 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W585)


if True == tmp2934 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W585)
return
}


}, 1)

tmp2935 := MakeNative(func(__e *ControlFlow) {
W586 := __e.Get(1)
_ = W586
tmp2946 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W586)


if True == tmp2946 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2936 := MakeNative(func(__e *ControlFlow) {
W587 := __e.Get(1)
_ = W587
tmp2937 := MakeNative(func(__e *ControlFlow) {
W588 := __e.Get(1)
_ = W588
tmp2942 := Call(__e, PrimFunc(symempty_2), W587)


var ifres2938 Obj

if True == tmp2942 {
ifres2938 = Nil


} else {
tmp2939 := Call(__e, PrimFunc(symshen_4app), W587, MakeString("\n ..."), symshen_4r)


tmp2940 := PrimStringConcat(MakeString("Shen syntax error here:\n "), tmp2939)

tmp2941 := PrimSimpleError(tmp2940)

ifres2938 = tmp2941


}

__e.TailApply(PrimFunc(symshen_4comb), W588, ifres2938)
return


}, 1)

tmp2943 := Call(__e, PrimFunc(symshen_4in_1_6), W586)


__e.TailApply(tmp2937, tmp2943)
return


}, 1)

tmp2944 := Call(__e, PrimFunc(symshen_4_5_1out), W586)


__e.TailApply(tmp2936, tmp2944)
return


}


}, 1)

tmp2947 := Call(__e, PrimFunc(sym_5_b_6), V577)


tmp2948 := Call(__e, tmp2935, tmp2947)


__e.TailApply(tmp2932, tmp2948)
return


} else {
__e.Return(W578)
return
}


}, 1)

tmp2951 := MakeNative(func(__e *ControlFlow) {
W579 := __e.Get(1)
_ = W579
tmp2967 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W579)


if True == tmp2967 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2952 := MakeNative(func(__e *ControlFlow) {
W580 := __e.Get(1)
_ = W580
tmp2953 := MakeNative(func(__e *ControlFlow) {
W581 := __e.Get(1)
_ = W581
tmp2954 := MakeNative(func(__e *ControlFlow) {
W582 := __e.Get(1)
_ = W582
tmp2962 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W582)


if True == tmp2962 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2955 := MakeNative(func(__e *ControlFlow) {
W583 := __e.Get(1)
_ = W583
tmp2956 := MakeNative(func(__e *ControlFlow) {
W584 := __e.Get(1)
_ = W584
tmp2957 := Call(__e, PrimFunc(symshen_4linearise), W580)


tmp2958 := PrimCons(tmp2957, W583)

__e.TailApply(PrimFunc(symshen_4comb), W584, tmp2958)
return


}, 1)

tmp2959 := Call(__e, PrimFunc(symshen_4in_1_6), W582)


__e.TailApply(tmp2956, tmp2959)
return


}, 1)

tmp2960 := Call(__e, PrimFunc(symshen_4_5_1out), W582)


__e.TailApply(tmp2955, tmp2960)
return


}


}, 1)

tmp2963 := Call(__e, PrimFunc(symshen_4_5rules_6), W581)


__e.TailApply(tmp2954, tmp2963)
return


}, 1)

tmp2964 := Call(__e, PrimFunc(symshen_4in_1_6), W579)


__e.TailApply(tmp2953, tmp2964)
return


}, 1)

tmp2965 := Call(__e, PrimFunc(symshen_4_5_1out), W579)


__e.TailApply(tmp2952, tmp2965)
return


}


}, 1)

tmp2968 := Call(__e, PrimFunc(symshen_4_5rule_6), V577)


tmp2969 := Call(__e, tmp2951, tmp2968)


__e.TailApply(tmp2931, tmp2969)
return


}, 1)

tmp2970 := Call(__e, ns2_1set, symshen_4_5rules_6, tmp2930)


_ = tmp2970

tmp2971 := MakeNative(func(__e *ControlFlow) {
V591 := __e.Get(1)
_ = V591
tmp2976 := Call(__e, PrimFunc(symtuple_2), V591)


if True == tmp2976 {
tmp2972 := Call(__e, PrimFunc(symfst), V591)


tmp2973 := Call(__e, PrimFunc(symfst), V591)


tmp2974 := Call(__e, PrimFunc(symsnd), V591)


__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp2972, tmp2973, Nil, tmp2974)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.linearise")))
return
}


}, 1)

tmp2977 := Call(__e, ns2_1set, symshen_4linearise, tmp2971)


_ = tmp2977

tmp2978 := MakeNative(func(__e *ControlFlow) {
V604 := __e.Get(1)
_ = V604
V605 := __e.Get(2)
_ = V605
V606 := __e.Get(3)
_ = V606
V607 := __e.Get(4)
_ = V607
tmp3016 := PrimEqual(Nil, V604)

if True == tmp3016 {
__e.TailApply(PrimFunc(sym_8p), V605, V607)
return
} else {
tmp3014 := PrimIsPair(V604)

var ifres3010 Obj

if True == tmp3014 {
tmp3012 := PrimHead(V604)

tmp3013 := PrimIsPair(tmp3012)

var ifres3011 Obj

if True == tmp3013 {
ifres3011 = True


} else {
ifres3011 = False


}

ifres3010 = ifres3011


} else {
ifres3010 = False


}

if True == ifres3010 {
tmp2979 := PrimHead(V604)

tmp2980 := PrimTail(V604)

tmp2981 := Call(__e, PrimFunc(symappend), tmp2979, tmp2980)


__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp2981, V605, V606, V607)
return


} else {
tmp3008 := PrimIsPair(V604)

var ifres3004 Obj

if True == tmp3008 {
tmp3006 := PrimHead(V604)

tmp3007 := PrimIsVariable(tmp3006)

var ifres3005 Obj

if True == tmp3007 {
ifres3005 = True


} else {
ifres3005 = False


}

ifres3004 = ifres3005


} else {
ifres3004 = False


}

if True == ifres3004 {
tmp2998 := PrimHead(V604)

tmp2999 := Call(__e, PrimFunc(symelement_2), tmp2998, V606)


if True == tmp2999 {
tmp2982 := MakeNative(func(__e *ControlFlow) {
W608 := __e.Get(1)
_ = W608
tmp2983 := PrimTail(V604)

tmp2984 := PrimHead(V604)

tmp2985 := Call(__e, PrimFunc(symshen_4rep_1X), tmp2984, W608, V605)


tmp2986 := PrimHead(V604)

tmp2987 := PrimCons(tmp2986, Nil)

tmp2988 := PrimCons(W608, tmp2987)

tmp2989 := PrimCons(sym_a, tmp2988)

tmp2990 := PrimCons(V607, Nil)

tmp2991 := PrimCons(tmp2989, tmp2990)

tmp2992 := PrimCons(symwhere, tmp2991)

__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp2983, tmp2985, V606, tmp2992)
return


}, 1)

tmp2993 := Call(__e, PrimFunc(symgensym), symV)


__e.TailApply(tmp2982, tmp2993)
return


} else {
tmp2994 := PrimTail(V604)

tmp2995 := PrimHead(V604)

tmp2996 := PrimCons(tmp2995, V606)

__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp2994, V605, tmp2996, V607)
return


}


} else {
tmp3002 := PrimIsPair(V604)

if True == tmp3002 {
tmp3000 := PrimTail(V604)

__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp3000, V605, V606, V607)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.linearise-h")))
return
}


}


}


}


}, 4)

tmp3017 := Call(__e, ns2_1set, symshen_4linearise_1h, tmp2978)


_ = tmp3017

tmp3018 := MakeNative(func(__e *ControlFlow) {
V609 := __e.Get(1)
_ = V609
tmp3019 := MakeNative(func(__e *ControlFlow) {
W610 := __e.Get(1)
_ = W610
tmp3107 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W610)


if True == tmp3107 {
tmp3020 := MakeNative(func(__e *ControlFlow) {
W620 := __e.Get(1)
_ = W620
tmp3085 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W620)


if True == tmp3085 {
tmp3021 := MakeNative(func(__e *ControlFlow) {
W627 := __e.Get(1)
_ = W627
tmp3048 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W627)


if True == tmp3048 {
tmp3022 := MakeNative(func(__e *ControlFlow) {
W637 := __e.Get(1)
_ = W637
tmp3024 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W637)


if True == tmp3024 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W637)
return
}


}, 1)

tmp3025 := MakeNative(func(__e *ControlFlow) {
W638 := __e.Get(1)
_ = W638
tmp3044 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W638)


if True == tmp3044 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3026 := MakeNative(func(__e *ControlFlow) {
W639 := __e.Get(1)
_ = W639
tmp3027 := MakeNative(func(__e *ControlFlow) {
W640 := __e.Get(1)
_ = W640
tmp3040 := Call(__e, PrimFunc(symshen_4hds_a_2), W640, sym_5_1)


if True == tmp3040 {
tmp3028 := MakeNative(func(__e *ControlFlow) {
W641 := __e.Get(1)
_ = W641
tmp3037 := PrimIsPair(W641)

if True == tmp3037 {
tmp3029 := MakeNative(func(__e *ControlFlow) {
W642 := __e.Get(1)
_ = W642
tmp3030 := MakeNative(func(__e *ControlFlow) {
W643 := __e.Get(1)
_ = W643
tmp3031 := PrimCons(W642, Nil)

tmp3032 := PrimCons(symshen_4choicepoint_b, tmp3031)

tmp3033 := Call(__e, PrimFunc(sym_8p), W639, tmp3032)


__e.TailApply(PrimFunc(symshen_4comb), W643, tmp3033)
return


}, 1)

tmp3034 := Call(__e, PrimFunc(symtail), W641)


__e.TailApply(tmp3030, tmp3034)
return


}, 1)

tmp3035 := Call(__e, PrimFunc(symhead), W641)


__e.TailApply(tmp3029, tmp3035)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3038 := Call(__e, PrimFunc(symtail), W640)


__e.TailApply(tmp3028, tmp3038)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3041 := Call(__e, PrimFunc(symshen_4in_1_6), W638)


__e.TailApply(tmp3027, tmp3041)
return


}, 1)

tmp3042 := Call(__e, PrimFunc(symshen_4_5_1out), W638)


__e.TailApply(tmp3026, tmp3042)
return


}


}, 1)

tmp3045 := Call(__e, PrimFunc(symshen_4_5patterns_6), V609)


tmp3046 := Call(__e, tmp3025, tmp3045)


__e.TailApply(tmp3022, tmp3046)
return


} else {
__e.Return(W627)
return
}


}, 1)

tmp3049 := MakeNative(func(__e *ControlFlow) {
W628 := __e.Get(1)
_ = W628
tmp3081 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W628)


if True == tmp3081 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3050 := MakeNative(func(__e *ControlFlow) {
W629 := __e.Get(1)
_ = W629
tmp3051 := MakeNative(func(__e *ControlFlow) {
W630 := __e.Get(1)
_ = W630
tmp3077 := Call(__e, PrimFunc(symshen_4hds_a_2), W630, sym_5_1)


if True == tmp3077 {
tmp3052 := MakeNative(func(__e *ControlFlow) {
W631 := __e.Get(1)
_ = W631
tmp3074 := PrimIsPair(W631)

if True == tmp3074 {
tmp3053 := MakeNative(func(__e *ControlFlow) {
W632 := __e.Get(1)
_ = W632
tmp3054 := MakeNative(func(__e *ControlFlow) {
W633 := __e.Get(1)
_ = W633
tmp3070 := Call(__e, PrimFunc(symshen_4hds_a_2), W633, symwhere)


if True == tmp3070 {
tmp3055 := MakeNative(func(__e *ControlFlow) {
W634 := __e.Get(1)
_ = W634
tmp3067 := PrimIsPair(W634)

if True == tmp3067 {
tmp3056 := MakeNative(func(__e *ControlFlow) {
W635 := __e.Get(1)
_ = W635
tmp3057 := MakeNative(func(__e *ControlFlow) {
W636 := __e.Get(1)
_ = W636
tmp3058 := PrimCons(W632, Nil)

tmp3059 := PrimCons(symshen_4choicepoint_b, tmp3058)

tmp3060 := PrimCons(tmp3059, Nil)

tmp3061 := PrimCons(W635, tmp3060)

tmp3062 := PrimCons(symwhere, tmp3061)

tmp3063 := Call(__e, PrimFunc(sym_8p), W629, tmp3062)


__e.TailApply(PrimFunc(symshen_4comb), W636, tmp3063)
return


}, 1)

tmp3064 := Call(__e, PrimFunc(symtail), W634)


__e.TailApply(tmp3057, tmp3064)
return


}, 1)

tmp3065 := Call(__e, PrimFunc(symhead), W634)


__e.TailApply(tmp3056, tmp3065)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3068 := Call(__e, PrimFunc(symtail), W633)


__e.TailApply(tmp3055, tmp3068)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3071 := Call(__e, PrimFunc(symtail), W631)


__e.TailApply(tmp3054, tmp3071)
return


}, 1)

tmp3072 := Call(__e, PrimFunc(symhead), W631)


__e.TailApply(tmp3053, tmp3072)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3075 := Call(__e, PrimFunc(symtail), W630)


__e.TailApply(tmp3052, tmp3075)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3078 := Call(__e, PrimFunc(symshen_4in_1_6), W628)


__e.TailApply(tmp3051, tmp3078)
return


}, 1)

tmp3079 := Call(__e, PrimFunc(symshen_4_5_1out), W628)


__e.TailApply(tmp3050, tmp3079)
return


}


}, 1)

tmp3082 := Call(__e, PrimFunc(symshen_4_5patterns_6), V609)


tmp3083 := Call(__e, tmp3049, tmp3082)


__e.TailApply(tmp3021, tmp3083)
return


} else {
__e.Return(W620)
return
}


}, 1)

tmp3086 := MakeNative(func(__e *ControlFlow) {
W621 := __e.Get(1)
_ = W621
tmp3103 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W621)


if True == tmp3103 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3087 := MakeNative(func(__e *ControlFlow) {
W622 := __e.Get(1)
_ = W622
tmp3088 := MakeNative(func(__e *ControlFlow) {
W623 := __e.Get(1)
_ = W623
tmp3099 := Call(__e, PrimFunc(symshen_4hds_a_2), W623, sym_1_6)


if True == tmp3099 {
tmp3089 := MakeNative(func(__e *ControlFlow) {
W624 := __e.Get(1)
_ = W624
tmp3096 := PrimIsPair(W624)

if True == tmp3096 {
tmp3090 := MakeNative(func(__e *ControlFlow) {
W625 := __e.Get(1)
_ = W625
tmp3091 := MakeNative(func(__e *ControlFlow) {
W626 := __e.Get(1)
_ = W626
tmp3092 := Call(__e, PrimFunc(sym_8p), W622, W625)


__e.TailApply(PrimFunc(symshen_4comb), W626, tmp3092)
return


}, 1)

tmp3093 := Call(__e, PrimFunc(symtail), W624)


__e.TailApply(tmp3091, tmp3093)
return


}, 1)

tmp3094 := Call(__e, PrimFunc(symhead), W624)


__e.TailApply(tmp3090, tmp3094)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3097 := Call(__e, PrimFunc(symtail), W623)


__e.TailApply(tmp3089, tmp3097)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3100 := Call(__e, PrimFunc(symshen_4in_1_6), W621)


__e.TailApply(tmp3088, tmp3100)
return


}, 1)

tmp3101 := Call(__e, PrimFunc(symshen_4_5_1out), W621)


__e.TailApply(tmp3087, tmp3101)
return


}


}, 1)

tmp3104 := Call(__e, PrimFunc(symshen_4_5patterns_6), V609)


tmp3105 := Call(__e, tmp3086, tmp3104)


__e.TailApply(tmp3020, tmp3105)
return


} else {
__e.Return(W610)
return
}


}, 1)

tmp3108 := MakeNative(func(__e *ControlFlow) {
W611 := __e.Get(1)
_ = W611
tmp3138 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W611)


if True == tmp3138 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3109 := MakeNative(func(__e *ControlFlow) {
W612 := __e.Get(1)
_ = W612
tmp3110 := MakeNative(func(__e *ControlFlow) {
W613 := __e.Get(1)
_ = W613
tmp3134 := Call(__e, PrimFunc(symshen_4hds_a_2), W613, sym_1_6)


if True == tmp3134 {
tmp3111 := MakeNative(func(__e *ControlFlow) {
W614 := __e.Get(1)
_ = W614
tmp3131 := PrimIsPair(W614)

if True == tmp3131 {
tmp3112 := MakeNative(func(__e *ControlFlow) {
W615 := __e.Get(1)
_ = W615
tmp3113 := MakeNative(func(__e *ControlFlow) {
W616 := __e.Get(1)
_ = W616
tmp3127 := Call(__e, PrimFunc(symshen_4hds_a_2), W616, symwhere)


if True == tmp3127 {
tmp3114 := MakeNative(func(__e *ControlFlow) {
W617 := __e.Get(1)
_ = W617
tmp3124 := PrimIsPair(W617)

if True == tmp3124 {
tmp3115 := MakeNative(func(__e *ControlFlow) {
W618 := __e.Get(1)
_ = W618
tmp3116 := MakeNative(func(__e *ControlFlow) {
W619 := __e.Get(1)
_ = W619
tmp3117 := PrimCons(W615, Nil)

tmp3118 := PrimCons(W618, tmp3117)

tmp3119 := PrimCons(symwhere, tmp3118)

tmp3120 := Call(__e, PrimFunc(sym_8p), W612, tmp3119)


__e.TailApply(PrimFunc(symshen_4comb), W619, tmp3120)
return


}, 1)

tmp3121 := Call(__e, PrimFunc(symtail), W617)


__e.TailApply(tmp3116, tmp3121)
return


}, 1)

tmp3122 := Call(__e, PrimFunc(symhead), W617)


__e.TailApply(tmp3115, tmp3122)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3125 := Call(__e, PrimFunc(symtail), W616)


__e.TailApply(tmp3114, tmp3125)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3128 := Call(__e, PrimFunc(symtail), W614)


__e.TailApply(tmp3113, tmp3128)
return


}, 1)

tmp3129 := Call(__e, PrimFunc(symhead), W614)


__e.TailApply(tmp3112, tmp3129)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3132 := Call(__e, PrimFunc(symtail), W613)


__e.TailApply(tmp3111, tmp3132)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3135 := Call(__e, PrimFunc(symshen_4in_1_6), W611)


__e.TailApply(tmp3110, tmp3135)
return


}, 1)

tmp3136 := Call(__e, PrimFunc(symshen_4_5_1out), W611)


__e.TailApply(tmp3109, tmp3136)
return


}


}, 1)

tmp3139 := Call(__e, PrimFunc(symshen_4_5patterns_6), V609)


tmp3140 := Call(__e, tmp3108, tmp3139)


__e.TailApply(tmp3019, tmp3140)
return


}, 1)

tmp3141 := Call(__e, ns2_1set, symshen_4_5rule_6, tmp3018)


_ = tmp3141

tmp3142 := MakeNative(func(__e *ControlFlow) {
V644 := __e.Get(1)
_ = V644
tmp3143 := MakeNative(func(__e *ControlFlow) {
W645 := __e.Get(1)
_ = W645
tmp3155 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W645)


if True == tmp3155 {
tmp3144 := MakeNative(func(__e *ControlFlow) {
W652 := __e.Get(1)
_ = W652
tmp3146 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W652)


if True == tmp3146 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W652)
return
}


}, 1)

tmp3147 := MakeNative(func(__e *ControlFlow) {
W653 := __e.Get(1)
_ = W653
tmp3151 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W653)


if True == tmp3151 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3148 := MakeNative(func(__e *ControlFlow) {
W654 := __e.Get(1)
_ = W654
__e.TailApply(PrimFunc(symshen_4comb), W654, Nil)
return
}, 1)

tmp3149 := Call(__e, PrimFunc(symshen_4in_1_6), W653)


__e.TailApply(tmp3148, tmp3149)
return


}


}, 1)

tmp3152 := Call(__e, PrimFunc(sym_5e_6), V644)


tmp3153 := Call(__e, tmp3147, tmp3152)


__e.TailApply(tmp3144, tmp3153)
return


} else {
__e.Return(W645)
return
}


}, 1)

tmp3156 := MakeNative(func(__e *ControlFlow) {
W646 := __e.Get(1)
_ = W646
tmp3171 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W646)


if True == tmp3171 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3157 := MakeNative(func(__e *ControlFlow) {
W647 := __e.Get(1)
_ = W647
tmp3158 := MakeNative(func(__e *ControlFlow) {
W648 := __e.Get(1)
_ = W648
tmp3159 := MakeNative(func(__e *ControlFlow) {
W649 := __e.Get(1)
_ = W649
tmp3166 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W649)


if True == tmp3166 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3160 := MakeNative(func(__e *ControlFlow) {
W650 := __e.Get(1)
_ = W650
tmp3161 := MakeNative(func(__e *ControlFlow) {
W651 := __e.Get(1)
_ = W651
tmp3162 := PrimCons(W647, W650)

__e.TailApply(PrimFunc(symshen_4comb), W651, tmp3162)
return


}, 1)

tmp3163 := Call(__e, PrimFunc(symshen_4in_1_6), W649)


__e.TailApply(tmp3161, tmp3163)
return


}, 1)

tmp3164 := Call(__e, PrimFunc(symshen_4_5_1out), W649)


__e.TailApply(tmp3160, tmp3164)
return


}


}, 1)

tmp3167 := Call(__e, PrimFunc(symshen_4_5patterns_6), W648)


__e.TailApply(tmp3159, tmp3167)
return


}, 1)

tmp3168 := Call(__e, PrimFunc(symshen_4in_1_6), W646)


__e.TailApply(tmp3158, tmp3168)
return


}, 1)

tmp3169 := Call(__e, PrimFunc(symshen_4_5_1out), W646)


__e.TailApply(tmp3157, tmp3169)
return


}


}, 1)

tmp3172 := Call(__e, PrimFunc(symshen_4_5pattern_6), V644)


tmp3173 := Call(__e, tmp3156, tmp3172)


__e.TailApply(tmp3143, tmp3173)
return


}, 1)

tmp3174 := Call(__e, ns2_1set, symshen_4_5patterns_6, tmp3142)


_ = tmp3174

tmp3175 := MakeNative(func(__e *ControlFlow) {
V655 := __e.Get(1)
_ = V655
tmp3176 := MakeNative(func(__e *ControlFlow) {
W656 := __e.Get(1)
_ = W656
tmp3231 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W656)


if True == tmp3231 {
tmp3177 := MakeNative(func(__e *ControlFlow) {
W670 := __e.Get(1)
_ = W670
tmp3205 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W670)


if True == tmp3205 {
tmp3178 := MakeNative(func(__e *ControlFlow) {
W677 := __e.Get(1)
_ = W677
tmp3192 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W677)


if True == tmp3192 {
tmp3179 := MakeNative(func(__e *ControlFlow) {
W680 := __e.Get(1)
_ = W680
tmp3181 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W680)


if True == tmp3181 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W680)
return
}


}, 1)

tmp3182 := MakeNative(func(__e *ControlFlow) {
W681 := __e.Get(1)
_ = W681
tmp3188 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W681)


if True == tmp3188 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3183 := MakeNative(func(__e *ControlFlow) {
W682 := __e.Get(1)
_ = W682
tmp3184 := MakeNative(func(__e *ControlFlow) {
W683 := __e.Get(1)
_ = W683
__e.TailApply(PrimFunc(symshen_4comb), W683, W682)
return
}, 1)

tmp3185 := Call(__e, PrimFunc(symshen_4in_1_6), W681)


__e.TailApply(tmp3184, tmp3185)
return


}, 1)

tmp3186 := Call(__e, PrimFunc(symshen_4_5_1out), W681)


__e.TailApply(tmp3183, tmp3186)
return


}


}, 1)

tmp3189 := Call(__e, PrimFunc(symshen_4_5simple_1pattern_6), V655)


tmp3190 := Call(__e, tmp3182, tmp3189)


__e.TailApply(tmp3179, tmp3190)
return


} else {
__e.Return(W677)
return
}


}, 1)

tmp3203 := PrimIsPair(V655)

var ifres3193 Obj

if True == tmp3203 {
tmp3194 := MakeNative(func(__e *ControlFlow) {
W678 := __e.Get(1)
_ = W678
tmp3195 := MakeNative(func(__e *ControlFlow) {
W679 := __e.Get(1)
_ = W679
tmp3198 := PrimIsPair(W678)

if True == tmp3198 {
tmp3196 := Call(__e, PrimFunc(symshen_4constructor_1error), W678)


__e.TailApply(PrimFunc(symshen_4comb), W679, tmp3196)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3199 := Call(__e, PrimFunc(symtail), V655)


__e.TailApply(tmp3195, tmp3199)
return


}, 1)

tmp3200 := Call(__e, PrimFunc(symhead), V655)


tmp3201 := Call(__e, tmp3194, tmp3200)


ifres3193 = tmp3201


} else {
tmp3202 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3193 = tmp3202


}

__e.TailApply(tmp3178, ifres3193)
return


} else {
__e.Return(W670)
return
}


}, 1)

tmp3229 := Call(__e, PrimFunc(symshen_4ccons_2), V655)


var ifres3206 Obj

if True == tmp3229 {
tmp3207 := MakeNative(func(__e *ControlFlow) {
W671 := __e.Get(1)
_ = W671
tmp3208 := MakeNative(func(__e *ControlFlow) {
W672 := __e.Get(1)
_ = W672
tmp3224 := Call(__e, PrimFunc(symshen_4hds_a_2), W671, symvector)


if True == tmp3224 {
tmp3209 := MakeNative(func(__e *ControlFlow) {
W673 := __e.Get(1)
_ = W673
tmp3221 := Call(__e, PrimFunc(symshen_4hds_a_2), W673, MakeNumber(0))


if True == tmp3221 {
tmp3210 := MakeNative(func(__e *ControlFlow) {
W674 := __e.Get(1)
_ = W674
tmp3211 := MakeNative(func(__e *ControlFlow) {
W675 := __e.Get(1)
_ = W675
tmp3217 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W675)


if True == tmp3217 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3212 := MakeNative(func(__e *ControlFlow) {
W676 := __e.Get(1)
_ = W676
tmp3213 := PrimCons(MakeNumber(0), Nil)

tmp3214 := PrimCons(symvector, tmp3213)

__e.TailApply(PrimFunc(symshen_4comb), W672, tmp3214)
return


}, 1)

tmp3215 := Call(__e, PrimFunc(symshen_4in_1_6), W675)


__e.TailApply(tmp3212, tmp3215)
return


}


}, 1)

tmp3218 := Call(__e, PrimFunc(sym_5end_6), W674)


__e.TailApply(tmp3211, tmp3218)
return


}, 1)

tmp3219 := Call(__e, PrimFunc(symtail), W673)


__e.TailApply(tmp3210, tmp3219)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3222 := Call(__e, PrimFunc(symtail), W671)


__e.TailApply(tmp3209, tmp3222)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3225 := Call(__e, PrimFunc(symtail), V655)


__e.TailApply(tmp3208, tmp3225)
return


}, 1)

tmp3226 := Call(__e, PrimFunc(symhead), V655)


tmp3227 := Call(__e, tmp3207, tmp3226)


ifres3206 = tmp3227


} else {
tmp3228 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3206 = tmp3228


}

__e.TailApply(tmp3177, ifres3206)
return


} else {
__e.Return(W656)
return
}


}, 1)

tmp3272 := Call(__e, PrimFunc(symshen_4ccons_2), V655)


var ifres3232 Obj

if True == tmp3272 {
tmp3233 := MakeNative(func(__e *ControlFlow) {
W657 := __e.Get(1)
_ = W657
tmp3234 := MakeNative(func(__e *ControlFlow) {
W658 := __e.Get(1)
_ = W658
tmp3235 := MakeNative(func(__e *ControlFlow) {
W659 := __e.Get(1)
_ = W659
tmp3266 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W659)


if True == tmp3266 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3236 := MakeNative(func(__e *ControlFlow) {
W660 := __e.Get(1)
_ = W660
tmp3237 := MakeNative(func(__e *ControlFlow) {
W661 := __e.Get(1)
_ = W661
tmp3238 := MakeNative(func(__e *ControlFlow) {
W662 := __e.Get(1)
_ = W662
tmp3261 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W662)


if True == tmp3261 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3239 := MakeNative(func(__e *ControlFlow) {
W663 := __e.Get(1)
_ = W663
tmp3240 := MakeNative(func(__e *ControlFlow) {
W664 := __e.Get(1)
_ = W664
tmp3241 := MakeNative(func(__e *ControlFlow) {
W665 := __e.Get(1)
_ = W665
tmp3256 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W665)


if True == tmp3256 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3242 := MakeNative(func(__e *ControlFlow) {
W666 := __e.Get(1)
_ = W666
tmp3243 := MakeNative(func(__e *ControlFlow) {
W667 := __e.Get(1)
_ = W667
tmp3244 := MakeNative(func(__e *ControlFlow) {
W668 := __e.Get(1)
_ = W668
tmp3251 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W668)


if True == tmp3251 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3245 := MakeNative(func(__e *ControlFlow) {
W669 := __e.Get(1)
_ = W669
tmp3246 := PrimCons(W666, Nil)

tmp3247 := PrimCons(W663, tmp3246)

tmp3248 := PrimCons(W660, tmp3247)

__e.TailApply(PrimFunc(symshen_4comb), W658, tmp3248)
return


}, 1)

tmp3249 := Call(__e, PrimFunc(symshen_4in_1_6), W668)


__e.TailApply(tmp3245, tmp3249)
return


}


}, 1)

tmp3252 := Call(__e, PrimFunc(sym_5end_6), W667)


__e.TailApply(tmp3244, tmp3252)
return


}, 1)

tmp3253 := Call(__e, PrimFunc(symshen_4in_1_6), W665)


__e.TailApply(tmp3243, tmp3253)
return


}, 1)

tmp3254 := Call(__e, PrimFunc(symshen_4_5_1out), W665)


__e.TailApply(tmp3242, tmp3254)
return


}


}, 1)

tmp3257 := Call(__e, PrimFunc(symshen_4_5pattern2_6), W664)


__e.TailApply(tmp3241, tmp3257)
return


}, 1)

tmp3258 := Call(__e, PrimFunc(symshen_4in_1_6), W662)


__e.TailApply(tmp3240, tmp3258)
return


}, 1)

tmp3259 := Call(__e, PrimFunc(symshen_4_5_1out), W662)


__e.TailApply(tmp3239, tmp3259)
return


}


}, 1)

tmp3262 := Call(__e, PrimFunc(symshen_4_5pattern1_6), W661)


__e.TailApply(tmp3238, tmp3262)
return


}, 1)

tmp3263 := Call(__e, PrimFunc(symshen_4in_1_6), W659)


__e.TailApply(tmp3237, tmp3263)
return


}, 1)

tmp3264 := Call(__e, PrimFunc(symshen_4_5_1out), W659)


__e.TailApply(tmp3236, tmp3264)
return


}


}, 1)

tmp3267 := Call(__e, PrimFunc(symshen_4_5constructor_6), W657)


__e.TailApply(tmp3235, tmp3267)
return


}, 1)

tmp3268 := Call(__e, PrimFunc(symtail), V655)


__e.TailApply(tmp3234, tmp3268)
return


}, 1)

tmp3269 := Call(__e, PrimFunc(symhead), V655)


tmp3270 := Call(__e, tmp3233, tmp3269)


ifres3232 = tmp3270


} else {
tmp3271 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3232 = tmp3271


}

__e.TailApply(tmp3176, ifres3232)
return


}, 1)

tmp3273 := Call(__e, ns2_1set, symshen_4_5pattern_6, tmp3175)


_ = tmp3273

tmp3274 := MakeNative(func(__e *ControlFlow) {
V684 := __e.Get(1)
_ = V684
tmp3275 := MakeNative(func(__e *ControlFlow) {
W685 := __e.Get(1)
_ = W685
tmp3277 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W685)


if True == tmp3277 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W685)
return
}


}, 1)

tmp3287 := PrimIsPair(V684)

var ifres3278 Obj

if True == tmp3287 {
tmp3279 := MakeNative(func(__e *ControlFlow) {
W686 := __e.Get(1)
_ = W686
tmp3280 := MakeNative(func(__e *ControlFlow) {
W687 := __e.Get(1)
_ = W687
tmp3282 := Call(__e, PrimFunc(symshen_4constructor_2), W686)


if True == tmp3282 {
__e.TailApply(PrimFunc(symshen_4comb), W687, W686)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3283 := Call(__e, PrimFunc(symtail), V684)


__e.TailApply(tmp3280, tmp3283)
return


}, 1)

tmp3284 := Call(__e, PrimFunc(symhead), V684)


tmp3285 := Call(__e, tmp3279, tmp3284)


ifres3278 = tmp3285


} else {
tmp3286 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3278 = tmp3286


}

__e.TailApply(tmp3275, ifres3278)
return


}, 1)

tmp3288 := Call(__e, ns2_1set, symshen_4_5constructor_6, tmp3274)


_ = tmp3288

tmp3289 := MakeNative(func(__e *ControlFlow) {
V688 := __e.Get(1)
_ = V688
tmp3290 := PrimCons(sym_8v, Nil)

tmp3291 := PrimCons(sym_8s, tmp3290)

tmp3292 := PrimCons(sym_8p, tmp3291)

tmp3293 := PrimCons(symcons, tmp3292)

__e.TailApply(PrimFunc(symelement_2), V688, tmp3293)
return


}, 1)

tmp3294 := Call(__e, ns2_1set, symshen_4constructor_2, tmp3289)


_ = tmp3294

tmp3295 := MakeNative(func(__e *ControlFlow) {
V689 := __e.Get(1)
_ = V689
tmp3296 := Call(__e, PrimFunc(symshen_4app), V689, MakeString(" is not a legitimate constructor\n"), symshen_4r)


__e.Return(PrimSimpleError(tmp3296))
return


}, 1)

tmp3297 := Call(__e, ns2_1set, symshen_4constructor_1error, tmp3295)


_ = tmp3297

tmp3298 := MakeNative(func(__e *ControlFlow) {
V690 := __e.Get(1)
_ = V690
tmp3299 := MakeNative(func(__e *ControlFlow) {
W691 := __e.Get(1)
_ = W691
tmp3317 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W691)


if True == tmp3317 {
tmp3300 := MakeNative(func(__e *ControlFlow) {
W694 := __e.Get(1)
_ = W694
tmp3302 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W694)


if True == tmp3302 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W694)
return
}


}, 1)

tmp3315 := PrimIsPair(V690)

var ifres3303 Obj

if True == tmp3315 {
tmp3304 := MakeNative(func(__e *ControlFlow) {
W695 := __e.Get(1)
_ = W695
tmp3305 := MakeNative(func(__e *ControlFlow) {
W696 := __e.Get(1)
_ = W696
tmp3307 := PrimCons(sym_5_1, Nil)

tmp3308 := PrimCons(sym_1_6, tmp3307)

tmp3309 := Call(__e, PrimFunc(symelement_2), W695, tmp3308)


tmp3310 := PrimNot(tmp3309)

if True == tmp3310 {
__e.TailApply(PrimFunc(symshen_4comb), W696, W695)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3311 := Call(__e, PrimFunc(symtail), V690)


__e.TailApply(tmp3305, tmp3311)
return


}, 1)

tmp3312 := Call(__e, PrimFunc(symhead), V690)


tmp3313 := Call(__e, tmp3304, tmp3312)


ifres3303 = tmp3313


} else {
tmp3314 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3303 = tmp3314


}

__e.TailApply(tmp3300, ifres3303)
return


} else {
__e.Return(W691)
return
}


}, 1)

tmp3328 := PrimIsPair(V690)

var ifres3318 Obj

if True == tmp3328 {
tmp3319 := MakeNative(func(__e *ControlFlow) {
W692 := __e.Get(1)
_ = W692
tmp3320 := MakeNative(func(__e *ControlFlow) {
W693 := __e.Get(1)
_ = W693
tmp3323 := PrimEqual(W692, sym__)

if True == tmp3323 {
tmp3321 := Call(__e, PrimFunc(symgensym), symY)


__e.TailApply(PrimFunc(symshen_4comb), W693, tmp3321)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3324 := Call(__e, PrimFunc(symtail), V690)


__e.TailApply(tmp3320, tmp3324)
return


}, 1)

tmp3325 := Call(__e, PrimFunc(symhead), V690)


tmp3326 := Call(__e, tmp3319, tmp3325)


ifres3318 = tmp3326


} else {
tmp3327 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3318 = tmp3327


}

__e.TailApply(tmp3299, ifres3318)
return


}, 1)

tmp3329 := Call(__e, ns2_1set, symshen_4_5simple_1pattern_6, tmp3298)


_ = tmp3329

tmp3330 := MakeNative(func(__e *ControlFlow) {
V697 := __e.Get(1)
_ = V697
tmp3331 := MakeNative(func(__e *ControlFlow) {
W698 := __e.Get(1)
_ = W698
tmp3333 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W698)


if True == tmp3333 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W698)
return
}


}, 1)

tmp3334 := MakeNative(func(__e *ControlFlow) {
W699 := __e.Get(1)
_ = W699
tmp3340 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W699)


if True == tmp3340 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3335 := MakeNative(func(__e *ControlFlow) {
W700 := __e.Get(1)
_ = W700
tmp3336 := MakeNative(func(__e *ControlFlow) {
W701 := __e.Get(1)
_ = W701
__e.TailApply(PrimFunc(symshen_4comb), W701, W700)
return
}, 1)

tmp3337 := Call(__e, PrimFunc(symshen_4in_1_6), W699)


__e.TailApply(tmp3336, tmp3337)
return


}, 1)

tmp3338 := Call(__e, PrimFunc(symshen_4_5_1out), W699)


__e.TailApply(tmp3335, tmp3338)
return


}


}, 1)

tmp3341 := Call(__e, PrimFunc(symshen_4_5pattern_6), V697)


tmp3342 := Call(__e, tmp3334, tmp3341)


__e.TailApply(tmp3331, tmp3342)
return


}, 1)

tmp3343 := Call(__e, ns2_1set, symshen_4_5pattern1_6, tmp3330)


_ = tmp3343

tmp3344 := MakeNative(func(__e *ControlFlow) {
V702 := __e.Get(1)
_ = V702
tmp3345 := MakeNative(func(__e *ControlFlow) {
W703 := __e.Get(1)
_ = W703
tmp3347 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W703)


if True == tmp3347 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W703)
return
}


}, 1)

tmp3348 := MakeNative(func(__e *ControlFlow) {
W704 := __e.Get(1)
_ = W704
tmp3354 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W704)


if True == tmp3354 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3349 := MakeNative(func(__e *ControlFlow) {
W705 := __e.Get(1)
_ = W705
tmp3350 := MakeNative(func(__e *ControlFlow) {
W706 := __e.Get(1)
_ = W706
__e.TailApply(PrimFunc(symshen_4comb), W706, W705)
return
}, 1)

tmp3351 := Call(__e, PrimFunc(symshen_4in_1_6), W704)


__e.TailApply(tmp3350, tmp3351)
return


}, 1)

tmp3352 := Call(__e, PrimFunc(symshen_4_5_1out), W704)


__e.TailApply(tmp3349, tmp3352)
return


}


}, 1)

tmp3355 := Call(__e, PrimFunc(symshen_4_5pattern_6), V702)


tmp3356 := Call(__e, tmp3348, tmp3355)


__e.TailApply(tmp3345, tmp3356)
return


}, 1)

tmp3357 := Call(__e, ns2_1set, symshen_4_5pattern2_6, tmp3344)


_ = tmp3357

tmp3358 := MakeNative(func(__e *ControlFlow) {
V707 := __e.Get(1)
_ = V707
tmp3359 := MakeNative(func(__e *ControlFlow) {
W708 := __e.Get(1)
_ = W708
tmp3360 := MakeNative(func(__e *ControlFlow) {
W709 := __e.Get(1)
_ = W709
tmp3361 := MakeNative(func(__e *ControlFlow) {
W710 := __e.Get(1)
_ = W710
__e.Return(W710)
return
}, 1)

tmp3362 := PrimStr(V707)

tmp3363 := Call(__e, PrimFunc(sym_8s), tmp3362, MakeString(")"))


tmp3364 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp3363)


tmp3365 := Call(__e, PrimFunc(sym_8s), MakeString("n"), tmp3364)


tmp3366 := Call(__e, PrimFunc(sym_8s), MakeString("f"), tmp3365)


tmp3367 := Call(__e, PrimFunc(sym_8s), MakeString("("), tmp3366)


tmp3368 := PrimVectorSet(W709, MakeNumber(1), tmp3367)

__e.TailApply(tmp3361, tmp3368)
return


}, 1)

tmp3369 := PrimVectorSet(W708, MakeNumber(0), symshen_4printF)

__e.TailApply(tmp3360, tmp3369)
return


}, 1)

tmp3370 := PrimAbsvector(MakeNumber(2))

__e.TailApply(tmp3359, tmp3370)
return


}, 1)

tmp3371 := Call(__e, ns2_1set, symshen_4fn_1print, tmp3358)


_ = tmp3371

tmp3372 := MakeNative(func(__e *ControlFlow) {
V711 := __e.Get(1)
_ = V711
__e.Return(PrimVectorGet(V711, MakeNumber(1)))
return
}, 1)

tmp3373 := Call(__e, ns2_1set, symshen_4printF, tmp3372)


_ = tmp3373

tmp3374 := MakeNative(func(__e *ControlFlow) {
V716 := __e.Get(1)
_ = V716
V717 := __e.Get(2)
_ = V717
tmp3398 := PrimIsPair(V717)

var ifres3394 Obj

if True == tmp3398 {
tmp3396 := PrimTail(V717)

tmp3397 := PrimEqual(Nil, tmp3396)

var ifres3395 Obj

if True == tmp3397 {
ifres3395 = True


} else {
ifres3395 = False


}

ifres3394 = ifres3395


} else {
ifres3394 = False


}

if True == ifres3394 {
tmp3375 := PrimHead(V717)

__e.TailApply(PrimFunc(symlength), tmp3375)
return


} else {
tmp3392 := PrimIsPair(V717)

var ifres3380 Obj

if True == tmp3392 {
tmp3390 := PrimTail(V717)

tmp3391 := PrimIsPair(tmp3390)

var ifres3382 Obj

if True == tmp3391 {
tmp3384 := PrimHead(V717)

tmp3385 := Call(__e, PrimFunc(symlength), tmp3384)


tmp3386 := PrimTail(V717)

tmp3387 := PrimHead(tmp3386)

tmp3388 := Call(__e, PrimFunc(symlength), tmp3387)


tmp3389 := PrimEqual(tmp3385, tmp3388)

var ifres3383 Obj

if True == tmp3389 {
ifres3383 = True


} else {
ifres3383 = False


}

ifres3382 = ifres3383


} else {
ifres3382 = False


}

var ifres3381 Obj

if True == ifres3382 {
ifres3381 = True


} else {
ifres3381 = False


}

ifres3380 = ifres3381


} else {
ifres3380 = False


}

if True == ifres3380 {
tmp3376 := PrimTail(V717)

__e.TailApply(PrimFunc(symshen_4arity_1chk), V716, tmp3376)
return


} else {
tmp3377 := Call(__e, PrimFunc(symshen_4app), V716, MakeString("\n"), symshen_4a)


tmp3378 := PrimStringConcat(MakeString("arity error in "), tmp3377)

__e.Return(PrimSimpleError(tmp3378))
return


}


}


}, 2)

tmp3399 := Call(__e, ns2_1set, symshen_4arity_1chk, tmp3374)


_ = tmp3399

tmp3400 := MakeNative(func(__e *ControlFlow) {
V718 := __e.Get(1)
_ = V718
V719 := __e.Get(2)
_ = V719
tmp3406 := Call(__e, PrimFunc(symtuple_2), V719)


if True == tmp3406 {
tmp3401 := Call(__e, PrimFunc(symfst), V719)


tmp3402 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp3401)


tmp3403 := Call(__e, PrimFunc(symsnd), V719)


tmp3404 := Call(__e, PrimFunc(symshen_4find_1free_1vars), tmp3402, tmp3403)


__e.TailApply(PrimFunc(symshen_4free_1variable_1error_1message), V718, tmp3404)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.free-var-chk")))
return
}


}, 2)

tmp3407 := Call(__e, ns2_1set, symshen_4free_1var_1chk, tmp3400)


_ = tmp3407

tmp3408 := MakeNative(func(__e *ControlFlow) {
V720 := __e.Get(1)
_ = V720
V721 := __e.Get(2)
_ = V721
tmp3420 := Call(__e, PrimFunc(symempty_2), V721)


if True == tmp3420 {
__e.Return(symshen_4skip)
return
} else {
tmp3409 := Call(__e, PrimFunc(symshen_4app), V720, MakeString(":"), symshen_4a)


tmp3410 := PrimStringConcat(MakeString("free variables in "), tmp3409)

tmp3411 := Call(__e, PrimFunc(symstoutput))


tmp3412 := Call(__e, PrimFunc(sympr), tmp3410, tmp3411)


_ = tmp3412

tmp3413 := MakeNative(func(__e *ControlFlow) {
Z722 := __e.Get(1)
_ = Z722
tmp3414 := Call(__e, PrimFunc(symshen_4app), Z722, MakeString(""), symshen_4a)


tmp3415 := PrimStringConcat(MakeString(" "), tmp3414)

tmp3416 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp3415, tmp3416)
return


}, 1)

tmp3417 := Call(__e, PrimFunc(symmap), tmp3413, V721)


_ = tmp3417

tmp3418 := Call(__e, PrimFunc(symnl), MakeNumber(1))


_ = tmp3418

__e.TailApply(PrimFunc(symabort))
return


}


}, 2)

tmp3421 := Call(__e, ns2_1set, symshen_4free_1variable_1error_1message, tmp3408)


_ = tmp3421

tmp3422 := MakeNative(func(__e *ControlFlow) {
V725 := __e.Get(1)
_ = V725
tmp3430 := PrimIsVariable(V725)

if True == tmp3430 {
__e.Return(PrimCons(V725, Nil))
return
} else {
tmp3428 := PrimIsPair(V725)

if True == tmp3428 {
tmp3423 := PrimHead(V725)

tmp3424 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp3423)


tmp3425 := PrimTail(V725)

tmp3426 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp3425)


__e.TailApply(PrimFunc(symunion), tmp3424, tmp3426)
return


} else {
__e.Return(Nil)
return
}


}


}, 1)

tmp3431 := Call(__e, ns2_1set, symshen_4extract_1vars, tmp3422)


_ = tmp3431

tmp3432 := MakeNative(func(__e *ControlFlow) {
V730 := __e.Get(1)
_ = V730
V731 := __e.Get(2)
_ = V731
tmp3522 := PrimIsPair(V731)

var ifres3509 Obj

if True == tmp3522 {
tmp3520 := PrimHead(V731)

tmp3521 := PrimEqual(symprotect, tmp3520)

var ifres3511 Obj

if True == tmp3521 {
tmp3518 := PrimTail(V731)

tmp3519 := PrimIsPair(tmp3518)

var ifres3513 Obj

if True == tmp3519 {
tmp3515 := PrimTail(V731)

tmp3516 := PrimTail(tmp3515)

tmp3517 := PrimEqual(Nil, tmp3516)

var ifres3514 Obj

if True == tmp3517 {
ifres3514 = True


} else {
ifres3514 = False


}

ifres3513 = ifres3514


} else {
ifres3513 = False


}

var ifres3512 Obj

if True == ifres3513 {
ifres3512 = True


} else {
ifres3512 = False


}

ifres3511 = ifres3512


} else {
ifres3511 = False


}

var ifres3510 Obj

if True == ifres3511 {
ifres3510 = True


} else {
ifres3510 = False


}

ifres3509 = ifres3510


} else {
ifres3509 = False


}

if True == ifres3509 {
__e.Return(Nil)
return
} else {
tmp3507 := PrimIsPair(V731)

var ifres3481 Obj

if True == tmp3507 {
tmp3505 := PrimHead(V731)

tmp3506 := PrimEqual(symlet, tmp3505)

var ifres3483 Obj

if True == tmp3506 {
tmp3503 := PrimTail(V731)

tmp3504 := PrimIsPair(tmp3503)

var ifres3485 Obj

if True == tmp3504 {
tmp3500 := PrimTail(V731)

tmp3501 := PrimTail(tmp3500)

tmp3502 := PrimIsPair(tmp3501)

var ifres3487 Obj

if True == tmp3502 {
tmp3496 := PrimTail(V731)

tmp3497 := PrimTail(tmp3496)

tmp3498 := PrimTail(tmp3497)

tmp3499 := PrimIsPair(tmp3498)

var ifres3489 Obj

if True == tmp3499 {
tmp3491 := PrimTail(V731)

tmp3492 := PrimTail(tmp3491)

tmp3493 := PrimTail(tmp3492)

tmp3494 := PrimTail(tmp3493)

tmp3495 := PrimEqual(Nil, tmp3494)

var ifres3490 Obj

if True == tmp3495 {
ifres3490 = True


} else {
ifres3490 = False


}

ifres3489 = ifres3490


} else {
ifres3489 = False


}

var ifres3488 Obj

if True == ifres3489 {
ifres3488 = True


} else {
ifres3488 = False


}

ifres3487 = ifres3488


} else {
ifres3487 = False


}

var ifres3486 Obj

if True == ifres3487 {
ifres3486 = True


} else {
ifres3486 = False


}

ifres3485 = ifres3486


} else {
ifres3485 = False


}

var ifres3484 Obj

if True == ifres3485 {
ifres3484 = True


} else {
ifres3484 = False


}

ifres3483 = ifres3484


} else {
ifres3483 = False


}

var ifres3482 Obj

if True == ifres3483 {
ifres3482 = True


} else {
ifres3482 = False


}

ifres3481 = ifres3482


} else {
ifres3481 = False


}

if True == ifres3481 {
tmp3433 := PrimTail(V731)

tmp3434 := PrimTail(tmp3433)

tmp3435 := PrimHead(tmp3434)

tmp3436 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V730, tmp3435)


tmp3437 := PrimTail(V731)

tmp3438 := PrimHead(tmp3437)

tmp3439 := PrimCons(tmp3438, V730)

tmp3440 := PrimTail(V731)

tmp3441 := PrimTail(tmp3440)

tmp3442 := PrimTail(tmp3441)

tmp3443 := PrimHead(tmp3442)

tmp3444 := Call(__e, PrimFunc(symshen_4find_1free_1vars), tmp3439, tmp3443)


__e.TailApply(PrimFunc(symunion), tmp3436, tmp3444)
return


} else {
tmp3479 := PrimIsPair(V731)

var ifres3460 Obj

if True == tmp3479 {
tmp3477 := PrimHead(V731)

tmp3478 := PrimEqual(symlambda, tmp3477)

var ifres3462 Obj

if True == tmp3478 {
tmp3475 := PrimTail(V731)

tmp3476 := PrimIsPair(tmp3475)

var ifres3464 Obj

if True == tmp3476 {
tmp3472 := PrimTail(V731)

tmp3473 := PrimTail(tmp3472)

tmp3474 := PrimIsPair(tmp3473)

var ifres3466 Obj

if True == tmp3474 {
tmp3468 := PrimTail(V731)

tmp3469 := PrimTail(tmp3468)

tmp3470 := PrimTail(tmp3469)

tmp3471 := PrimEqual(Nil, tmp3470)

var ifres3467 Obj

if True == tmp3471 {
ifres3467 = True


} else {
ifres3467 = False


}

ifres3466 = ifres3467


} else {
ifres3466 = False


}

var ifres3465 Obj

if True == ifres3466 {
ifres3465 = True


} else {
ifres3465 = False


}

ifres3464 = ifres3465


} else {
ifres3464 = False


}

var ifres3463 Obj

if True == ifres3464 {
ifres3463 = True


} else {
ifres3463 = False


}

ifres3462 = ifres3463


} else {
ifres3462 = False


}

var ifres3461 Obj

if True == ifres3462 {
ifres3461 = True


} else {
ifres3461 = False


}

ifres3460 = ifres3461


} else {
ifres3460 = False


}

if True == ifres3460 {
tmp3445 := PrimTail(V731)

tmp3446 := PrimHead(tmp3445)

tmp3447 := PrimCons(tmp3446, V730)

tmp3448 := PrimTail(V731)

tmp3449 := PrimTail(tmp3448)

tmp3450 := PrimHead(tmp3449)

__e.TailApply(PrimFunc(symshen_4find_1free_1vars), tmp3447, tmp3450)
return


} else {
tmp3458 := PrimIsPair(V731)

if True == tmp3458 {
tmp3451 := PrimHead(V731)

tmp3452 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V730, tmp3451)


tmp3453 := PrimTail(V731)

tmp3454 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V730, tmp3453)


__e.TailApply(PrimFunc(symunion), tmp3452, tmp3454)
return


} else {
tmp3456 := Call(__e, PrimFunc(symshen_4free_1variable_2), V731, V730)


if True == tmp3456 {
__e.Return(PrimCons(V731, Nil))
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

tmp3523 := Call(__e, ns2_1set, symshen_4find_1free_1vars, tmp3432)


_ = tmp3523

tmp3524 := MakeNative(func(__e *ControlFlow) {
V732 := __e.Get(1)
_ = V732
V733 := __e.Get(2)
_ = V733
tmp3529 := PrimIsVariable(V732)

if True == tmp3529 {
tmp3526 := Call(__e, PrimFunc(symelement_2), V732, V733)


tmp3527 := PrimNot(tmp3526)

if True == tmp3527 {
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

tmp3530 := Call(__e, ns2_1set, symshen_4free_1variable_2, tmp3524)


_ = tmp3530

tmp3531 := MakeNative(func(__e *ControlFlow) {
V734 := __e.Get(1)
_ = V734
V735 := __e.Get(2)
_ = V735
tmp3532 := PrimValue(symshen_4_duserdefs_d)

tmp3533 := Call(__e, PrimFunc(symadjoin), V734, tmp3532)


tmp3534 := PrimSet(symshen_4_duserdefs_d, tmp3533)

_ = tmp3534

tmp3535 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V734, symshen_4source, V735, tmp3535)
return


}, 2)

tmp3536 := Call(__e, ns2_1set, symshen_4record_1kl, tmp3531)


_ = tmp3536

tmp3537 := MakeNative(func(__e *ControlFlow) {
V736 := __e.Get(1)
_ = V736
V737 := __e.Get(2)
_ = V737
V738 := __e.Get(3)
_ = V738
tmp3538 := MakeNative(func(__e *ControlFlow) {
W739 := __e.Get(1)
_ = W739
tmp3539 := MakeNative(func(__e *ControlFlow) {
W740 := __e.Get(1)
_ = W740
tmp3540 := MakeNative(func(__e *ControlFlow) {
W741 := __e.Get(1)
_ = W741
__e.Return(W741)
return
}, 1)

tmp3541 := Call(__e, PrimFunc(symshen_4cond_1form), W740)


tmp3542 := PrimCons(tmp3541, Nil)

tmp3543 := PrimCons(W739, tmp3542)

tmp3544 := PrimCons(V736, tmp3543)

tmp3545 := PrimCons(symdefun, tmp3544)

__e.TailApply(tmp3540, tmp3545)
return


}, 1)

tmp3546 := Call(__e, PrimFunc(symshen_4kl_1body), V737, W739)


tmp3547 := Call(__e, PrimFunc(symshen_4scan_1body), V736, tmp3546)


__e.TailApply(tmp3539, tmp3547)
return


}, 1)

tmp3548 := Call(__e, PrimFunc(symshen_4parameters), V738)


__e.TailApply(tmp3538, tmp3548)
return


}, 3)

tmp3549 := Call(__e, ns2_1set, symshen_4compile_1to_1kl, tmp3537)


_ = tmp3549

tmp3550 := MakeNative(func(__e *ControlFlow) {
V742 := __e.Get(1)
_ = V742
tmp3555 := PrimEqual(MakeNumber(0), V742)

if True == tmp3555 {
__e.Return(Nil)
return
} else {
tmp3551 := Call(__e, PrimFunc(symgensym), symV)


tmp3552 := PrimNumberSubtract(V742, MakeNumber(1))

tmp3553 := Call(__e, PrimFunc(symshen_4parameters), tmp3552)


__e.Return(PrimCons(tmp3551, tmp3553))
return


}


}, 1)

tmp3556 := Call(__e, ns2_1set, symshen_4parameters, tmp3550)


_ = tmp3556

tmp3557 := MakeNative(func(__e *ControlFlow) {
V745 := __e.Get(1)
_ = V745
tmp3581 := PrimIsPair(V745)

var ifres3561 Obj

if True == tmp3581 {
tmp3579 := PrimHead(V745)

tmp3580 := PrimIsPair(tmp3579)

var ifres3563 Obj

if True == tmp3580 {
tmp3576 := PrimHead(V745)

tmp3577 := PrimHead(tmp3576)

tmp3578 := PrimEqual(True, tmp3577)

var ifres3565 Obj

if True == tmp3578 {
tmp3573 := PrimHead(V745)

tmp3574 := PrimTail(tmp3573)

tmp3575 := PrimIsPair(tmp3574)

var ifres3567 Obj

if True == tmp3575 {
tmp3569 := PrimHead(V745)

tmp3570 := PrimTail(tmp3569)

tmp3571 := PrimTail(tmp3570)

tmp3572 := PrimEqual(Nil, tmp3571)

var ifres3568 Obj

if True == tmp3572 {
ifres3568 = True


} else {
ifres3568 = False


}

ifres3567 = ifres3568


} else {
ifres3567 = False


}

var ifres3566 Obj

if True == ifres3567 {
ifres3566 = True


} else {
ifres3566 = False


}

ifres3565 = ifres3566


} else {
ifres3565 = False


}

var ifres3564 Obj

if True == ifres3565 {
ifres3564 = True


} else {
ifres3564 = False


}

ifres3563 = ifres3564


} else {
ifres3563 = False


}

var ifres3562 Obj

if True == ifres3563 {
ifres3562 = True


} else {
ifres3562 = False


}

ifres3561 = ifres3562


} else {
ifres3561 = False


}

if True == ifres3561 {
tmp3558 := PrimHead(V745)

tmp3559 := PrimTail(tmp3558)

__e.Return(PrimHead(tmp3559))
return


} else {
__e.Return(PrimCons(symcond, V745))
return
}


}, 1)

tmp3582 := Call(__e, ns2_1set, symshen_4cond_1form, tmp3557)


_ = tmp3582

tmp3583 := MakeNative(func(__e *ControlFlow) {
V754 := __e.Get(1)
_ = V754
V755 := __e.Get(2)
_ = V755
tmp3627 := PrimEqual(Nil, V755)

if True == tmp3627 {
tmp3584 := PrimCons(V754, Nil)

tmp3585 := PrimCons(symshen_4f_1error, tmp3584)

tmp3586 := PrimCons(tmp3585, Nil)

tmp3587 := PrimCons(True, tmp3586)

__e.Return(PrimCons(tmp3587, Nil))
return


} else {
tmp3625 := PrimIsPair(V755)

var ifres3621 Obj

if True == tmp3625 {
tmp3623 := PrimHead(V755)

tmp3624 := Call(__e, PrimFunc(symshen_4choicepoint_2), tmp3623)


var ifres3622 Obj

if True == tmp3624 {
ifres3622 = True


} else {
ifres3622 = False


}

ifres3621 = ifres3622


} else {
ifres3621 = False


}

if True == ifres3621 {
tmp3588 := Call(__e, PrimFunc(symgensym), symFreeze)


tmp3589 := Call(__e, PrimFunc(symgensym), symResult)


tmp3590 := PrimHead(V755)

tmp3591 := PrimTail(V755)

__e.TailApply(PrimFunc(symshen_4choicepoint), V754, tmp3588, tmp3589, tmp3590, tmp3591)
return


} else {
tmp3619 := PrimIsPair(V755)

var ifres3599 Obj

if True == tmp3619 {
tmp3617 := PrimHead(V755)

tmp3618 := PrimIsPair(tmp3617)

var ifres3601 Obj

if True == tmp3618 {
tmp3614 := PrimHead(V755)

tmp3615 := PrimHead(tmp3614)

tmp3616 := PrimEqual(True, tmp3615)

var ifres3603 Obj

if True == tmp3616 {
tmp3611 := PrimHead(V755)

tmp3612 := PrimTail(tmp3611)

tmp3613 := PrimIsPair(tmp3612)

var ifres3605 Obj

if True == tmp3613 {
tmp3607 := PrimHead(V755)

tmp3608 := PrimTail(tmp3607)

tmp3609 := PrimTail(tmp3608)

tmp3610 := PrimEqual(Nil, tmp3609)

var ifres3606 Obj

if True == tmp3610 {
ifres3606 = True


} else {
ifres3606 = False


}

ifres3605 = ifres3606


} else {
ifres3605 = False


}

var ifres3604 Obj

if True == ifres3605 {
ifres3604 = True


} else {
ifres3604 = False


}

ifres3603 = ifres3604


} else {
ifres3603 = False


}

var ifres3602 Obj

if True == ifres3603 {
ifres3602 = True


} else {
ifres3602 = False


}

ifres3601 = ifres3602


} else {
ifres3601 = False


}

var ifres3600 Obj

if True == ifres3601 {
ifres3600 = True


} else {
ifres3600 = False


}

ifres3599 = ifres3600


} else {
ifres3599 = False


}

if True == ifres3599 {
tmp3592 := PrimHead(V755)

__e.Return(PrimCons(tmp3592, Nil))
return


} else {
tmp3597 := PrimIsPair(V755)

if True == tmp3597 {
tmp3593 := PrimHead(V755)

tmp3594 := PrimTail(V755)

tmp3595 := Call(__e, PrimFunc(symshen_4scan_1body), V754, tmp3594)


__e.Return(PrimCons(tmp3593, tmp3595))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.scan-body")))
return
}


}


}


}


}, 2)

tmp3628 := Call(__e, ns2_1set, symshen_4scan_1body, tmp3583)


_ = tmp3628

tmp3629 := MakeNative(func(__e *ControlFlow) {
V762 := __e.Get(1)
_ = V762
tmp3664 := PrimIsPair(V762)

var ifres3631 Obj

if True == tmp3664 {
tmp3662 := PrimTail(V762)

tmp3663 := PrimIsPair(tmp3662)

var ifres3633 Obj

if True == tmp3663 {
tmp3659 := PrimTail(V762)

tmp3660 := PrimHead(tmp3659)

tmp3661 := PrimIsPair(tmp3660)

var ifres3635 Obj

if True == tmp3661 {
tmp3655 := PrimTail(V762)

tmp3656 := PrimHead(tmp3655)

tmp3657 := PrimHead(tmp3656)

tmp3658 := PrimEqual(symshen_4choicepoint_b, tmp3657)

var ifres3637 Obj

if True == tmp3658 {
tmp3651 := PrimTail(V762)

tmp3652 := PrimHead(tmp3651)

tmp3653 := PrimTail(tmp3652)

tmp3654 := PrimIsPair(tmp3653)

var ifres3639 Obj

if True == tmp3654 {
tmp3646 := PrimTail(V762)

tmp3647 := PrimHead(tmp3646)

tmp3648 := PrimTail(tmp3647)

tmp3649 := PrimTail(tmp3648)

tmp3650 := PrimEqual(Nil, tmp3649)

var ifres3641 Obj

if True == tmp3650 {
tmp3643 := PrimTail(V762)

tmp3644 := PrimTail(tmp3643)

tmp3645 := PrimEqual(Nil, tmp3644)

var ifres3642 Obj

if True == tmp3645 {
ifres3642 = True


} else {
ifres3642 = False


}

ifres3641 = ifres3642


} else {
ifres3641 = False


}

var ifres3640 Obj

if True == ifres3641 {
ifres3640 = True


} else {
ifres3640 = False


}

ifres3639 = ifres3640


} else {
ifres3639 = False


}

var ifres3638 Obj

if True == ifres3639 {
ifres3638 = True


} else {
ifres3638 = False


}

ifres3637 = ifres3638


} else {
ifres3637 = False


}

var ifres3636 Obj

if True == ifres3637 {
ifres3636 = True


} else {
ifres3636 = False


}

ifres3635 = ifres3636


} else {
ifres3635 = False


}

var ifres3634 Obj

if True == ifres3635 {
ifres3634 = True


} else {
ifres3634 = False


}

ifres3633 = ifres3634


} else {
ifres3633 = False


}

var ifres3632 Obj

if True == ifres3633 {
ifres3632 = True


} else {
ifres3632 = False


}

ifres3631 = ifres3632


} else {
ifres3631 = False


}

if True == ifres3631 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp3665 := Call(__e, ns2_1set, symshen_4choicepoint_2, tmp3629)


_ = tmp3665

tmp3666 := MakeNative(func(__e *ControlFlow) {
V778 := __e.Get(1)
_ = V778
V779 := __e.Get(2)
_ = V779
V780 := __e.Get(3)
_ = V780
V781 := __e.Get(4)
_ = V781
V782 := __e.Get(5)
_ = V782
tmp3858 := PrimIsPair(V781)

var ifres3780 Obj

if True == tmp3858 {
tmp3856 := PrimTail(V781)

tmp3857 := PrimIsPair(tmp3856)

var ifres3782 Obj

if True == tmp3857 {
tmp3853 := PrimTail(V781)

tmp3854 := PrimHead(tmp3853)

tmp3855 := PrimIsPair(tmp3854)

var ifres3784 Obj

if True == tmp3855 {
tmp3849 := PrimTail(V781)

tmp3850 := PrimHead(tmp3849)

tmp3851 := PrimTail(tmp3850)

tmp3852 := PrimIsPair(tmp3851)

var ifres3786 Obj

if True == tmp3852 {
tmp3844 := PrimTail(V781)

tmp3845 := PrimHead(tmp3844)

tmp3846 := PrimTail(tmp3845)

tmp3847 := PrimHead(tmp3846)

tmp3848 := PrimIsPair(tmp3847)

var ifres3788 Obj

if True == tmp3848 {
tmp3838 := PrimTail(V781)

tmp3839 := PrimHead(tmp3838)

tmp3840 := PrimTail(tmp3839)

tmp3841 := PrimHead(tmp3840)

tmp3842 := PrimHead(tmp3841)

tmp3843 := PrimEqual(symfail_1if, tmp3842)

var ifres3790 Obj

if True == tmp3843 {
tmp3832 := PrimTail(V781)

tmp3833 := PrimHead(tmp3832)

tmp3834 := PrimTail(tmp3833)

tmp3835 := PrimHead(tmp3834)

tmp3836 := PrimTail(tmp3835)

tmp3837 := PrimIsPair(tmp3836)

var ifres3792 Obj

if True == tmp3837 {
tmp3825 := PrimTail(V781)

tmp3826 := PrimHead(tmp3825)

tmp3827 := PrimTail(tmp3826)

tmp3828 := PrimHead(tmp3827)

tmp3829 := PrimTail(tmp3828)

tmp3830 := PrimTail(tmp3829)

tmp3831 := PrimIsPair(tmp3830)

var ifres3794 Obj

if True == tmp3831 {
tmp3817 := PrimTail(V781)

tmp3818 := PrimHead(tmp3817)

tmp3819 := PrimTail(tmp3818)

tmp3820 := PrimHead(tmp3819)

tmp3821 := PrimTail(tmp3820)

tmp3822 := PrimTail(tmp3821)

tmp3823 := PrimTail(tmp3822)

tmp3824 := PrimEqual(Nil, tmp3823)

var ifres3796 Obj

if True == tmp3824 {
tmp3812 := PrimTail(V781)

tmp3813 := PrimHead(tmp3812)

tmp3814 := PrimTail(tmp3813)

tmp3815 := PrimTail(tmp3814)

tmp3816 := PrimEqual(Nil, tmp3815)

var ifres3798 Obj

if True == tmp3816 {
tmp3809 := PrimTail(V781)

tmp3810 := PrimTail(tmp3809)

tmp3811 := PrimEqual(Nil, tmp3810)

var ifres3800 Obj

if True == tmp3811 {
tmp3802 := PrimTail(V781)

tmp3803 := PrimHead(tmp3802)

tmp3804 := PrimTail(tmp3803)

tmp3805 := PrimHead(tmp3804)

tmp3806 := PrimTail(tmp3805)

tmp3807 := PrimHead(tmp3806)

tmp3808 := PrimEqual(V778, tmp3807)

var ifres3801 Obj

if True == tmp3808 {
ifres3801 = True


} else {
ifres3801 = False


}

ifres3800 = ifres3801


} else {
ifres3800 = False


}

var ifres3799 Obj

if True == ifres3800 {
ifres3799 = True


} else {
ifres3799 = False


}

ifres3798 = ifres3799


} else {
ifres3798 = False


}

var ifres3797 Obj

if True == ifres3798 {
ifres3797 = True


} else {
ifres3797 = False


}

ifres3796 = ifres3797


} else {
ifres3796 = False


}

var ifres3795 Obj

if True == ifres3796 {
ifres3795 = True


} else {
ifres3795 = False


}

ifres3794 = ifres3795


} else {
ifres3794 = False


}

var ifres3793 Obj

if True == ifres3794 {
ifres3793 = True


} else {
ifres3793 = False


}

ifres3792 = ifres3793


} else {
ifres3792 = False


}

var ifres3791 Obj

if True == ifres3792 {
ifres3791 = True


} else {
ifres3791 = False


}

ifres3790 = ifres3791


} else {
ifres3790 = False


}

var ifres3789 Obj

if True == ifres3790 {
ifres3789 = True


} else {
ifres3789 = False


}

ifres3788 = ifres3789


} else {
ifres3788 = False


}

var ifres3787 Obj

if True == ifres3788 {
ifres3787 = True


} else {
ifres3787 = False


}

ifres3786 = ifres3787


} else {
ifres3786 = False


}

var ifres3785 Obj

if True == ifres3786 {
ifres3785 = True


} else {
ifres3785 = False


}

ifres3784 = ifres3785


} else {
ifres3784 = False


}

var ifres3783 Obj

if True == ifres3784 {
ifres3783 = True


} else {
ifres3783 = False


}

ifres3782 = ifres3783


} else {
ifres3782 = False


}

var ifres3781 Obj

if True == ifres3782 {
ifres3781 = True


} else {
ifres3781 = False


}

ifres3780 = ifres3781


} else {
ifres3780 = False


}

if True == ifres3780 {
tmp3667 := PrimTail(V781)

tmp3668 := PrimHead(tmp3667)

tmp3669 := PrimTail(tmp3668)

tmp3670 := PrimHead(tmp3669)

tmp3671 := PrimTail(tmp3670)

tmp3672 := PrimHead(tmp3671)

tmp3673 := Call(__e, PrimFunc(symshen_4scan_1body), tmp3672, V782)


tmp3674 := PrimCons(symcond, tmp3673)

tmp3675 := PrimCons(tmp3674, Nil)

tmp3676 := PrimCons(symfreeze, tmp3675)

tmp3677 := PrimHead(V781)

tmp3678 := PrimTail(V781)

tmp3679 := PrimHead(tmp3678)

tmp3680 := PrimTail(tmp3679)

tmp3681 := PrimHead(tmp3680)

tmp3682 := PrimTail(tmp3681)

tmp3683 := PrimTail(tmp3682)

tmp3684 := PrimHead(tmp3683)

tmp3685 := PrimTail(V781)

tmp3686 := PrimHead(tmp3685)

tmp3687 := PrimTail(tmp3686)

tmp3688 := PrimHead(tmp3687)

tmp3689 := PrimTail(tmp3688)

tmp3690 := PrimHead(tmp3689)

tmp3691 := PrimCons(V780, Nil)

tmp3692 := PrimCons(tmp3690, tmp3691)

tmp3693 := PrimCons(V779, Nil)

tmp3694 := PrimCons(symthaw, tmp3693)

tmp3695 := PrimCons(V780, Nil)

tmp3696 := PrimCons(tmp3694, tmp3695)

tmp3697 := PrimCons(tmp3692, tmp3696)

tmp3698 := PrimCons(symif, tmp3697)

tmp3699 := PrimCons(tmp3698, Nil)

tmp3700 := PrimCons(tmp3684, tmp3699)

tmp3701 := PrimCons(V780, tmp3700)

tmp3702 := PrimCons(symlet, tmp3701)

tmp3703 := PrimCons(V779, Nil)

tmp3704 := PrimCons(symthaw, tmp3703)

tmp3705 := PrimCons(tmp3704, Nil)

tmp3706 := PrimCons(tmp3702, tmp3705)

tmp3707 := PrimCons(tmp3677, tmp3706)

tmp3708 := PrimCons(symif, tmp3707)

tmp3709 := PrimCons(tmp3708, Nil)

tmp3710 := PrimCons(tmp3676, tmp3709)

tmp3711 := PrimCons(V779, tmp3710)

tmp3712 := PrimCons(symlet, tmp3711)

tmp3713 := PrimCons(tmp3712, Nil)

tmp3714 := PrimCons(True, tmp3713)

__e.Return(PrimCons(tmp3714, Nil))
return


} else {
tmp3778 := PrimIsPair(V781)

var ifres3751 Obj

if True == tmp3778 {
tmp3776 := PrimTail(V781)

tmp3777 := PrimIsPair(tmp3776)

var ifres3753 Obj

if True == tmp3777 {
tmp3773 := PrimTail(V781)

tmp3774 := PrimHead(tmp3773)

tmp3775 := PrimIsPair(tmp3774)

var ifres3755 Obj

if True == tmp3775 {
tmp3769 := PrimTail(V781)

tmp3770 := PrimHead(tmp3769)

tmp3771 := PrimTail(tmp3770)

tmp3772 := PrimIsPair(tmp3771)

var ifres3757 Obj

if True == tmp3772 {
tmp3764 := PrimTail(V781)

tmp3765 := PrimHead(tmp3764)

tmp3766 := PrimTail(tmp3765)

tmp3767 := PrimTail(tmp3766)

tmp3768 := PrimEqual(Nil, tmp3767)

var ifres3759 Obj

if True == tmp3768 {
tmp3761 := PrimTail(V781)

tmp3762 := PrimTail(tmp3761)

tmp3763 := PrimEqual(Nil, tmp3762)

var ifres3760 Obj

if True == tmp3763 {
ifres3760 = True


} else {
ifres3760 = False


}

ifres3759 = ifres3760


} else {
ifres3759 = False


}

var ifres3758 Obj

if True == ifres3759 {
ifres3758 = True


} else {
ifres3758 = False


}

ifres3757 = ifres3758


} else {
ifres3757 = False


}

var ifres3756 Obj

if True == ifres3757 {
ifres3756 = True


} else {
ifres3756 = False


}

ifres3755 = ifres3756


} else {
ifres3755 = False


}

var ifres3754 Obj

if True == ifres3755 {
ifres3754 = True


} else {
ifres3754 = False


}

ifres3753 = ifres3754


} else {
ifres3753 = False


}

var ifres3752 Obj

if True == ifres3753 {
ifres3752 = True


} else {
ifres3752 = False


}

ifres3751 = ifres3752


} else {
ifres3751 = False


}

if True == ifres3751 {
tmp3715 := Call(__e, PrimFunc(symshen_4scan_1body), V778, V782)


tmp3716 := PrimCons(symcond, tmp3715)

tmp3717 := PrimCons(tmp3716, Nil)

tmp3718 := PrimCons(symfreeze, tmp3717)

tmp3719 := PrimHead(V781)

tmp3720 := PrimTail(V781)

tmp3721 := PrimHead(tmp3720)

tmp3722 := PrimTail(tmp3721)

tmp3723 := PrimHead(tmp3722)

tmp3724 := PrimCons(symfail, Nil)

tmp3725 := PrimCons(tmp3724, Nil)

tmp3726 := PrimCons(V780, tmp3725)

tmp3727 := PrimCons(sym_a, tmp3726)

tmp3728 := PrimCons(V779, Nil)

tmp3729 := PrimCons(symthaw, tmp3728)

tmp3730 := PrimCons(V780, Nil)

tmp3731 := PrimCons(tmp3729, tmp3730)

tmp3732 := PrimCons(tmp3727, tmp3731)

tmp3733 := PrimCons(symif, tmp3732)

tmp3734 := PrimCons(tmp3733, Nil)

tmp3735 := PrimCons(tmp3723, tmp3734)

tmp3736 := PrimCons(V780, tmp3735)

tmp3737 := PrimCons(symlet, tmp3736)

tmp3738 := PrimCons(V779, Nil)

tmp3739 := PrimCons(symthaw, tmp3738)

tmp3740 := PrimCons(tmp3739, Nil)

tmp3741 := PrimCons(tmp3737, tmp3740)

tmp3742 := PrimCons(tmp3719, tmp3741)

tmp3743 := PrimCons(symif, tmp3742)

tmp3744 := PrimCons(tmp3743, Nil)

tmp3745 := PrimCons(tmp3718, tmp3744)

tmp3746 := PrimCons(V779, tmp3745)

tmp3747 := PrimCons(symlet, tmp3746)

tmp3748 := PrimCons(tmp3747, Nil)

tmp3749 := PrimCons(True, tmp3748)

__e.Return(PrimCons(tmp3749, Nil))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.choicepoint")))
return
}


}


}, 5)

tmp3859 := Call(__e, ns2_1set, symshen_4choicepoint, tmp3666)


_ = tmp3859

tmp3860 := MakeNative(func(__e *ControlFlow) {
V784 := __e.Get(1)
_ = V784
V785 := __e.Get(2)
_ = V785
V786 := __e.Get(3)
_ = V786
tmp3874 := PrimEqual(V784, V786)

if True == tmp3874 {
__e.Return(V785)
return
} else {
tmp3872 := PrimIsPair(V786)

if True == tmp3872 {
tmp3861 := MakeNative(func(__e *ControlFlow) {
W787 := __e.Get(1)
_ = W787
tmp3867 := PrimHead(V786)

tmp3868 := PrimEqual(W787, tmp3867)

if True == tmp3868 {
tmp3862 := PrimHead(V786)

tmp3863 := PrimTail(V786)

tmp3864 := Call(__e, PrimFunc(symshen_4rep_1X), V784, V785, tmp3863)


__e.Return(PrimCons(tmp3862, tmp3864))
return


} else {
tmp3865 := PrimTail(V786)

__e.Return(PrimCons(W787, tmp3865))
return


}


}, 1)

tmp3869 := PrimHead(V786)

tmp3870 := Call(__e, PrimFunc(symshen_4rep_1X), V784, V785, tmp3869)


__e.TailApply(tmp3861, tmp3870)
return


} else {
__e.Return(V786)
return
}


}


}, 3)

tmp3875 := Call(__e, ns2_1set, symshen_4rep_1X, tmp3860)


_ = tmp3875

tmp3876 := MakeNative(func(__e *ControlFlow) {
V788 := __e.Get(1)
_ = V788
V789 := __e.Get(2)
_ = V789
tmp3877 := MakeNative(func(__e *ControlFlow) {
Z790 := __e.Get(1)
_ = Z790
tmp3878 := Call(__e, PrimFunc(symfst), Z790)


tmp3879 := Call(__e, PrimFunc(symsnd), Z790)


tmp3880 := Call(__e, PrimFunc(symshen_4alpha_1convert), tmp3879)


__e.TailApply(PrimFunc(symshen_4triple_1stack), Nil, tmp3878, V789, tmp3880)
return


}, 1)

__e.TailApply(PrimFunc(symmap), tmp3877, V788)
return


}, 2)

tmp3881 := Call(__e, ns2_1set, symshen_4kl_1body, tmp3876)


_ = tmp3881

tmp3882 := MakeNative(func(__e *ControlFlow) {
V791 := __e.Get(1)
_ = V791
tmp3965 := PrimIsPair(V791)

var ifres3946 Obj

if True == tmp3965 {
tmp3963 := PrimHead(V791)

tmp3964 := PrimEqual(symlambda, tmp3963)

var ifres3948 Obj

if True == tmp3964 {
tmp3961 := PrimTail(V791)

tmp3962 := PrimIsPair(tmp3961)

var ifres3950 Obj

if True == tmp3962 {
tmp3958 := PrimTail(V791)

tmp3959 := PrimTail(tmp3958)

tmp3960 := PrimIsPair(tmp3959)

var ifres3952 Obj

if True == tmp3960 {
tmp3954 := PrimTail(V791)

tmp3955 := PrimTail(tmp3954)

tmp3956 := PrimTail(tmp3955)

tmp3957 := PrimEqual(Nil, tmp3956)

var ifres3953 Obj

if True == tmp3957 {
ifres3953 = True


} else {
ifres3953 = False


}

ifres3952 = ifres3953


} else {
ifres3952 = False


}

var ifres3951 Obj

if True == ifres3952 {
ifres3951 = True


} else {
ifres3951 = False


}

ifres3950 = ifres3951


} else {
ifres3950 = False


}

var ifres3949 Obj

if True == ifres3950 {
ifres3949 = True


} else {
ifres3949 = False


}

ifres3948 = ifres3949


} else {
ifres3948 = False


}

var ifres3947 Obj

if True == ifres3948 {
ifres3947 = True


} else {
ifres3947 = False


}

ifres3946 = ifres3947


} else {
ifres3946 = False


}

if True == ifres3946 {
tmp3883 := MakeNative(func(__e *ControlFlow) {
W792 := __e.Get(1)
_ = W792
tmp3884 := MakeNative(func(__e *ControlFlow) {
W793 := __e.Get(1)
_ = W793
tmp3885 := MakeNative(func(__e *ControlFlow) {
Z794 := __e.Get(1)
_ = Z794
__e.TailApply(PrimFunc(symshen_4alpha_1convert), Z794)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp3885, W793)
return


}, 1)

tmp3886 := PrimTail(V791)

tmp3887 := PrimHead(tmp3886)

tmp3888 := PrimTail(V791)

tmp3889 := PrimTail(tmp3888)

tmp3890 := PrimHead(tmp3889)

tmp3891 := Call(__e, PrimFunc(symshen_4beta), tmp3887, W792, tmp3890)


tmp3892 := PrimCons(tmp3891, Nil)

tmp3893 := PrimCons(W792, tmp3892)

tmp3894 := PrimCons(symlambda, tmp3893)

__e.TailApply(tmp3884, tmp3894)
return


}, 1)

tmp3895 := Call(__e, PrimFunc(symgensym), symZ)


__e.TailApply(tmp3883, tmp3895)
return


} else {
tmp3944 := PrimIsPair(V791)

var ifres3918 Obj

if True == tmp3944 {
tmp3942 := PrimHead(V791)

tmp3943 := PrimEqual(symlet, tmp3942)

var ifres3920 Obj

if True == tmp3943 {
tmp3940 := PrimTail(V791)

tmp3941 := PrimIsPair(tmp3940)

var ifres3922 Obj

if True == tmp3941 {
tmp3937 := PrimTail(V791)

tmp3938 := PrimTail(tmp3937)

tmp3939 := PrimIsPair(tmp3938)

var ifres3924 Obj

if True == tmp3939 {
tmp3933 := PrimTail(V791)

tmp3934 := PrimTail(tmp3933)

tmp3935 := PrimTail(tmp3934)

tmp3936 := PrimIsPair(tmp3935)

var ifres3926 Obj

if True == tmp3936 {
tmp3928 := PrimTail(V791)

tmp3929 := PrimTail(tmp3928)

tmp3930 := PrimTail(tmp3929)

tmp3931 := PrimTail(tmp3930)

tmp3932 := PrimEqual(Nil, tmp3931)

var ifres3927 Obj

if True == tmp3932 {
ifres3927 = True


} else {
ifres3927 = False


}

ifres3926 = ifres3927


} else {
ifres3926 = False


}

var ifres3925 Obj

if True == ifres3926 {
ifres3925 = True


} else {
ifres3925 = False


}

ifres3924 = ifres3925


} else {
ifres3924 = False


}

var ifres3923 Obj

if True == ifres3924 {
ifres3923 = True


} else {
ifres3923 = False


}

ifres3922 = ifres3923


} else {
ifres3922 = False


}

var ifres3921 Obj

if True == ifres3922 {
ifres3921 = True


} else {
ifres3921 = False


}

ifres3920 = ifres3921


} else {
ifres3920 = False


}

var ifres3919 Obj

if True == ifres3920 {
ifres3919 = True


} else {
ifres3919 = False


}

ifres3918 = ifres3919


} else {
ifres3918 = False


}

if True == ifres3918 {
tmp3896 := MakeNative(func(__e *ControlFlow) {
W795 := __e.Get(1)
_ = W795
tmp3897 := MakeNative(func(__e *ControlFlow) {
W796 := __e.Get(1)
_ = W796
tmp3898 := MakeNative(func(__e *ControlFlow) {
Z797 := __e.Get(1)
_ = Z797
__e.TailApply(PrimFunc(symshen_4alpha_1convert), Z797)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp3898, W796)
return


}, 1)

tmp3899 := PrimTail(V791)

tmp3900 := PrimTail(tmp3899)

tmp3901 := PrimHead(tmp3900)

tmp3902 := PrimTail(V791)

tmp3903 := PrimHead(tmp3902)

tmp3904 := PrimTail(V791)

tmp3905 := PrimTail(tmp3904)

tmp3906 := PrimTail(tmp3905)

tmp3907 := PrimHead(tmp3906)

tmp3908 := Call(__e, PrimFunc(symshen_4beta), tmp3903, W795, tmp3907)


tmp3909 := PrimCons(tmp3908, Nil)

tmp3910 := PrimCons(tmp3901, tmp3909)

tmp3911 := PrimCons(W795, tmp3910)

tmp3912 := PrimCons(symlet, tmp3911)

__e.TailApply(tmp3897, tmp3912)
return


}, 1)

tmp3913 := Call(__e, PrimFunc(symgensym), symW)


__e.TailApply(tmp3896, tmp3913)
return


} else {
tmp3916 := PrimIsPair(V791)

if True == tmp3916 {
tmp3914 := MakeNative(func(__e *ControlFlow) {
Z798 := __e.Get(1)
_ = Z798
__e.TailApply(PrimFunc(symshen_4alpha_1convert), Z798)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp3914, V791)
return


} else {
__e.Return(V791)
return
}


}


}


}, 1)

tmp3966 := Call(__e, ns2_1set, symshen_4alpha_1convert, tmp3882)


_ = tmp3966

tmp3967 := MakeNative(func(__e *ControlFlow) {
V807 := __e.Get(1)
_ = V807
V808 := __e.Get(2)
_ = V808
V809 := __e.Get(3)
_ = V809
V810 := __e.Get(4)
_ = V810
tmp4097 := PrimEqual(Nil, V808)

var ifres4072 Obj

if True == tmp4097 {
tmp4096 := PrimEqual(Nil, V809)

var ifres4074 Obj

if True == tmp4096 {
tmp4095 := PrimIsPair(V810)

var ifres4076 Obj

if True == tmp4095 {
tmp4093 := PrimHead(V810)

tmp4094 := PrimEqual(symwhere, tmp4093)

var ifres4078 Obj

if True == tmp4094 {
tmp4091 := PrimTail(V810)

tmp4092 := PrimIsPair(tmp4091)

var ifres4080 Obj

if True == tmp4092 {
tmp4088 := PrimTail(V810)

tmp4089 := PrimTail(tmp4088)

tmp4090 := PrimIsPair(tmp4089)

var ifres4082 Obj

if True == tmp4090 {
tmp4084 := PrimTail(V810)

tmp4085 := PrimTail(tmp4084)

tmp4086 := PrimTail(tmp4085)

tmp4087 := PrimEqual(Nil, tmp4086)

var ifres4083 Obj

if True == tmp4087 {
ifres4083 = True


} else {
ifres4083 = False


}

ifres4082 = ifres4083


} else {
ifres4082 = False


}

var ifres4081 Obj

if True == ifres4082 {
ifres4081 = True


} else {
ifres4081 = False


}

ifres4080 = ifres4081


} else {
ifres4080 = False


}

var ifres4079 Obj

if True == ifres4080 {
ifres4079 = True


} else {
ifres4079 = False


}

ifres4078 = ifres4079


} else {
ifres4078 = False


}

var ifres4077 Obj

if True == ifres4078 {
ifres4077 = True


} else {
ifres4077 = False


}

ifres4076 = ifres4077


} else {
ifres4076 = False


}

var ifres4075 Obj

if True == ifres4076 {
ifres4075 = True


} else {
ifres4075 = False


}

ifres4074 = ifres4075


} else {
ifres4074 = False


}

var ifres4073 Obj

if True == ifres4074 {
ifres4073 = True


} else {
ifres4073 = False


}

ifres4072 = ifres4073


} else {
ifres4072 = False


}

if True == ifres4072 {
tmp3968 := PrimTail(V810)

tmp3969 := PrimHead(tmp3968)

tmp3970 := PrimCons(tmp3969, V807)

tmp3971 := PrimTail(V810)

tmp3972 := PrimTail(tmp3971)

tmp3973 := PrimHead(tmp3972)

__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp3970, Nil, Nil, tmp3973)
return


} else {
tmp4070 := PrimEqual(Nil, V808)

var ifres4067 Obj

if True == tmp4070 {
tmp4069 := PrimEqual(Nil, V809)

var ifres4068 Obj

if True == tmp4069 {
ifres4068 = True


} else {
ifres4068 = False


}

ifres4067 = ifres4068


} else {
ifres4067 = False


}

if True == ifres4067 {
tmp3974 := Call(__e, PrimFunc(symreverse), V807)


tmp3975 := Call(__e, PrimFunc(symshen_4rectify_1test), tmp3974)


tmp3976 := PrimCons(V810, Nil)

__e.Return(PrimCons(tmp3975, tmp3976))
return


} else {
tmp4065 := PrimIsPair(V808)

var ifres4058 Obj

if True == tmp4065 {
tmp4064 := PrimIsPair(V809)

var ifres4060 Obj

if True == tmp4064 {
tmp4062 := PrimHead(V808)

tmp4063 := PrimIsVariable(tmp4062)

var ifres4061 Obj

if True == tmp4063 {
ifres4061 = True


} else {
ifres4061 = False


}

ifres4060 = ifres4061


} else {
ifres4060 = False


}

var ifres4059 Obj

if True == ifres4060 {
ifres4059 = True


} else {
ifres4059 = False


}

ifres4058 = ifres4059


} else {
ifres4058 = False


}

if True == ifres4058 {
tmp3977 := PrimTail(V808)

tmp3978 := PrimTail(V809)

tmp3979 := PrimHead(V808)

tmp3980 := PrimHead(V809)

tmp3981 := Call(__e, PrimFunc(symshen_4beta), tmp3979, tmp3980, V810)


__e.TailApply(PrimFunc(symshen_4triple_1stack), V807, tmp3977, tmp3978, tmp3981)
return


} else {
tmp4056 := PrimIsPair(V808)

var ifres4031 Obj

if True == tmp4056 {
tmp4054 := PrimHead(V808)

tmp4055 := PrimIsPair(tmp4054)

var ifres4033 Obj

if True == tmp4055 {
tmp4051 := PrimHead(V808)

tmp4052 := PrimTail(tmp4051)

tmp4053 := PrimIsPair(tmp4052)

var ifres4035 Obj

if True == tmp4053 {
tmp4047 := PrimHead(V808)

tmp4048 := PrimTail(tmp4047)

tmp4049 := PrimTail(tmp4048)

tmp4050 := PrimIsPair(tmp4049)

var ifres4037 Obj

if True == tmp4050 {
tmp4042 := PrimHead(V808)

tmp4043 := PrimTail(tmp4042)

tmp4044 := PrimTail(tmp4043)

tmp4045 := PrimTail(tmp4044)

tmp4046 := PrimEqual(Nil, tmp4045)

var ifres4039 Obj

if True == tmp4046 {
tmp4041 := PrimIsPair(V809)

var ifres4040 Obj

if True == tmp4041 {
ifres4040 = True


} else {
ifres4040 = False


}

ifres4039 = ifres4040


} else {
ifres4039 = False


}

var ifres4038 Obj

if True == ifres4039 {
ifres4038 = True


} else {
ifres4038 = False


}

ifres4037 = ifres4038


} else {
ifres4037 = False


}

var ifres4036 Obj

if True == ifres4037 {
ifres4036 = True


} else {
ifres4036 = False


}

ifres4035 = ifres4036


} else {
ifres4035 = False


}

var ifres4034 Obj

if True == ifres4035 {
ifres4034 = True


} else {
ifres4034 = False


}

ifres4033 = ifres4034


} else {
ifres4033 = False


}

var ifres4032 Obj

if True == ifres4033 {
ifres4032 = True


} else {
ifres4032 = False


}

ifres4031 = ifres4032


} else {
ifres4031 = False


}

if True == ifres4031 {
tmp3982 := PrimHead(V808)

tmp3983 := PrimHead(tmp3982)

tmp3984 := Call(__e, PrimFunc(symshen_4op_1test), tmp3983)


tmp3985 := PrimHead(V809)

tmp3986 := PrimCons(tmp3985, Nil)

tmp3987 := PrimCons(tmp3984, tmp3986)

tmp3988 := PrimCons(tmp3987, V807)

tmp3989 := PrimHead(V808)

tmp3990 := PrimTail(tmp3989)

tmp3991 := PrimHead(tmp3990)

tmp3992 := PrimHead(V808)

tmp3993 := PrimTail(tmp3992)

tmp3994 := PrimTail(tmp3993)

tmp3995 := PrimHead(tmp3994)

tmp3996 := PrimTail(V808)

tmp3997 := PrimCons(tmp3995, tmp3996)

tmp3998 := PrimCons(tmp3991, tmp3997)

tmp3999 := PrimHead(V808)

tmp4000 := PrimHead(tmp3999)

tmp4001 := Call(__e, PrimFunc(symshen_4op1), tmp4000)


tmp4002 := PrimHead(V809)

tmp4003 := PrimCons(tmp4002, Nil)

tmp4004 := PrimCons(tmp4001, tmp4003)

tmp4005 := PrimHead(V808)

tmp4006 := PrimHead(tmp4005)

tmp4007 := Call(__e, PrimFunc(symshen_4op2), tmp4006)


tmp4008 := PrimHead(V809)

tmp4009 := PrimCons(tmp4008, Nil)

tmp4010 := PrimCons(tmp4007, tmp4009)

tmp4011 := PrimTail(V809)

tmp4012 := PrimCons(tmp4010, tmp4011)

tmp4013 := PrimCons(tmp4004, tmp4012)

tmp4014 := PrimHead(V808)

tmp4015 := PrimHead(V809)

tmp4016 := Call(__e, PrimFunc(symshen_4beta), tmp4014, tmp4015, V810)


__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp3988, tmp3998, tmp4013, tmp4016)
return


} else {
tmp4029 := PrimIsPair(V808)

var ifres4026 Obj

if True == tmp4029 {
tmp4028 := PrimIsPair(V809)

var ifres4027 Obj

if True == tmp4028 {
ifres4027 = True


} else {
ifres4027 = False


}

ifres4026 = ifres4027


} else {
ifres4026 = False


}

if True == ifres4026 {
tmp4017 := PrimHead(V808)

tmp4018 := PrimHead(V809)

tmp4019 := PrimCons(tmp4018, Nil)

tmp4020 := PrimCons(tmp4017, tmp4019)

tmp4021 := PrimCons(sym_a, tmp4020)

tmp4022 := PrimCons(tmp4021, V807)

tmp4023 := PrimTail(V808)

tmp4024 := PrimTail(V809)

__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp4022, tmp4023, tmp4024, V810)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.triple-stack")))
return
}


}


}


}


}


}, 4)

tmp4098 := Call(__e, ns2_1set, symshen_4triple_1stack, tmp3967)


_ = tmp4098

tmp4099 := MakeNative(func(__e *ControlFlow) {
V813 := __e.Get(1)
_ = V813
tmp4118 := PrimEqual(Nil, V813)

if True == tmp4118 {
__e.Return(True)
return
} else {
tmp4116 := PrimIsPair(V813)

var ifres4112 Obj

if True == tmp4116 {
tmp4114 := PrimTail(V813)

tmp4115 := PrimEqual(Nil, tmp4114)

var ifres4113 Obj

if True == tmp4115 {
ifres4113 = True


} else {
ifres4113 = False


}

ifres4112 = ifres4113


} else {
ifres4112 = False


}

if True == ifres4112 {
__e.Return(PrimHead(V813))
return
} else {
tmp4110 := PrimIsPair(V813)

var ifres4106 Obj

if True == tmp4110 {
tmp4108 := PrimTail(V813)

tmp4109 := PrimIsPair(tmp4108)

var ifres4107 Obj

if True == tmp4109 {
ifres4107 = True


} else {
ifres4107 = False


}

ifres4106 = ifres4107


} else {
ifres4106 = False


}

if True == ifres4106 {
tmp4100 := PrimHead(V813)

tmp4101 := PrimTail(V813)

tmp4102 := Call(__e, PrimFunc(symshen_4rectify_1test), tmp4101)


tmp4103 := PrimCons(tmp4102, Nil)

tmp4104 := PrimCons(tmp4100, tmp4103)

__e.Return(PrimCons(symand, tmp4104))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.rectify-test")))
return
}


}


}


}, 1)

tmp4119 := Call(__e, ns2_1set, symshen_4rectify_1test, tmp4099)


_ = tmp4119

tmp4120 := MakeNative(func(__e *ControlFlow) {
V823 := __e.Get(1)
_ = V823
V824 := __e.Get(2)
_ = V824
V825 := __e.Get(3)
_ = V825
tmp4197 := PrimEqual(V823, V825)

if True == tmp4197 {
__e.Return(V824)
return
} else {
tmp4195 := PrimIsPair(V825)

var ifres4171 Obj

if True == tmp4195 {
tmp4193 := PrimHead(V825)

tmp4194 := PrimEqual(symlambda, tmp4193)

var ifres4173 Obj

if True == tmp4194 {
tmp4191 := PrimTail(V825)

tmp4192 := PrimIsPair(tmp4191)

var ifres4175 Obj

if True == tmp4192 {
tmp4188 := PrimTail(V825)

tmp4189 := PrimTail(tmp4188)

tmp4190 := PrimIsPair(tmp4189)

var ifres4177 Obj

if True == tmp4190 {
tmp4184 := PrimTail(V825)

tmp4185 := PrimTail(tmp4184)

tmp4186 := PrimTail(tmp4185)

tmp4187 := PrimEqual(Nil, tmp4186)

var ifres4179 Obj

if True == tmp4187 {
tmp4181 := PrimTail(V825)

tmp4182 := PrimHead(tmp4181)

tmp4183 := PrimEqual(V823, tmp4182)

var ifres4180 Obj

if True == tmp4183 {
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

var ifres4176 Obj

if True == ifres4177 {
ifres4176 = True


} else {
ifres4176 = False


}

ifres4175 = ifres4176


} else {
ifres4175 = False


}

var ifres4174 Obj

if True == ifres4175 {
ifres4174 = True


} else {
ifres4174 = False


}

ifres4173 = ifres4174


} else {
ifres4173 = False


}

var ifres4172 Obj

if True == ifres4173 {
ifres4172 = True


} else {
ifres4172 = False


}

ifres4171 = ifres4172


} else {
ifres4171 = False


}

if True == ifres4171 {
__e.Return(V825)
return
} else {
tmp4169 := PrimIsPair(V825)

var ifres4138 Obj

if True == tmp4169 {
tmp4167 := PrimHead(V825)

tmp4168 := PrimEqual(symlet, tmp4167)

var ifres4140 Obj

if True == tmp4168 {
tmp4165 := PrimTail(V825)

tmp4166 := PrimIsPair(tmp4165)

var ifres4142 Obj

if True == tmp4166 {
tmp4162 := PrimTail(V825)

tmp4163 := PrimTail(tmp4162)

tmp4164 := PrimIsPair(tmp4163)

var ifres4144 Obj

if True == tmp4164 {
tmp4158 := PrimTail(V825)

tmp4159 := PrimTail(tmp4158)

tmp4160 := PrimTail(tmp4159)

tmp4161 := PrimIsPair(tmp4160)

var ifres4146 Obj

if True == tmp4161 {
tmp4153 := PrimTail(V825)

tmp4154 := PrimTail(tmp4153)

tmp4155 := PrimTail(tmp4154)

tmp4156 := PrimTail(tmp4155)

tmp4157 := PrimEqual(Nil, tmp4156)

var ifres4148 Obj

if True == tmp4157 {
tmp4150 := PrimTail(V825)

tmp4151 := PrimHead(tmp4150)

tmp4152 := PrimEqual(V823, tmp4151)

var ifres4149 Obj

if True == tmp4152 {
ifres4149 = True


} else {
ifres4149 = False


}

ifres4148 = ifres4149


} else {
ifres4148 = False


}

var ifres4147 Obj

if True == ifres4148 {
ifres4147 = True


} else {
ifres4147 = False


}

ifres4146 = ifres4147


} else {
ifres4146 = False


}

var ifres4145 Obj

if True == ifres4146 {
ifres4145 = True


} else {
ifres4145 = False


}

ifres4144 = ifres4145


} else {
ifres4144 = False


}

var ifres4143 Obj

if True == ifres4144 {
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

if True == ifres4138 {
tmp4121 := PrimTail(V825)

tmp4122 := PrimHead(tmp4121)

tmp4123 := PrimTail(V825)

tmp4124 := PrimHead(tmp4123)

tmp4125 := PrimTail(V825)

tmp4126 := PrimTail(tmp4125)

tmp4127 := PrimHead(tmp4126)

tmp4128 := Call(__e, PrimFunc(symshen_4beta), tmp4124, V824, tmp4127)


tmp4129 := PrimTail(V825)

tmp4130 := PrimTail(tmp4129)

tmp4131 := PrimTail(tmp4130)

tmp4132 := PrimCons(tmp4128, tmp4131)

tmp4133 := PrimCons(tmp4122, tmp4132)

__e.Return(PrimCons(symlet, tmp4133))
return


} else {
tmp4136 := PrimIsPair(V825)

if True == tmp4136 {
tmp4134 := MakeNative(func(__e *ControlFlow) {
Z826 := __e.Get(1)
_ = Z826
__e.TailApply(PrimFunc(symshen_4beta), V823, V824, Z826)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp4134, V825)
return


} else {
__e.Return(V825)
return
}


}


}


}


}, 3)

tmp4198 := Call(__e, ns2_1set, symshen_4beta, tmp4120)


_ = tmp4198

tmp4199 := MakeNative(func(__e *ControlFlow) {
V829 := __e.Get(1)
_ = V829
tmp4207 := PrimEqual(symcons, V829)

if True == tmp4207 {
__e.Return(symhd)
return
} else {
tmp4205 := PrimEqual(sym_8s, V829)

if True == tmp4205 {
__e.Return(symhdstr)
return
} else {
tmp4203 := PrimEqual(sym_8p, V829)

if True == tmp4203 {
__e.Return(symfst)
return
} else {
tmp4201 := PrimEqual(sym_8v, V829)

if True == tmp4201 {
__e.Return(symhdv)
return
} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.op1")))
return
}


}


}


}


}, 1)

tmp4208 := Call(__e, ns2_1set, symshen_4op1, tmp4199)


_ = tmp4208

tmp4209 := MakeNative(func(__e *ControlFlow) {
V832 := __e.Get(1)
_ = V832
tmp4217 := PrimEqual(symcons, V832)

if True == tmp4217 {
__e.Return(symtl)
return
} else {
tmp4215 := PrimEqual(sym_8s, V832)

if True == tmp4215 {
__e.Return(symtlstr)
return
} else {
tmp4213 := PrimEqual(sym_8p, V832)

if True == tmp4213 {
__e.Return(symsnd)
return
} else {
tmp4211 := PrimEqual(sym_8v, V832)

if True == tmp4211 {
__e.Return(symtlv)
return
} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.op2")))
return
}


}


}


}


}, 1)

tmp4218 := Call(__e, ns2_1set, symshen_4op2, tmp4209)


_ = tmp4218

tmp4219 := MakeNative(func(__e *ControlFlow) {
V835 := __e.Get(1)
_ = V835
tmp4227 := PrimEqual(symcons, V835)

if True == tmp4227 {
__e.Return(symcons_2)
return
} else {
tmp4225 := PrimEqual(sym_8s, V835)

if True == tmp4225 {
__e.Return(symshen_4_7string_2)
return
} else {
tmp4223 := PrimEqual(sym_8p, V835)

if True == tmp4223 {
__e.Return(symtuple_2)
return
} else {
tmp4221 := PrimEqual(sym_8v, V835)

if True == tmp4221 {
__e.Return(symshen_4_7vector_2)
return
} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.op-test")))
return
}


}


}


}


}, 1)

tmp4228 := Call(__e, ns2_1set, symshen_4op_1test, tmp4219)


_ = tmp4228

tmp4229 := MakeNative(func(__e *ControlFlow) {
V836 := __e.Get(1)
_ = V836
tmp4231 := PrimEqual(MakeString(""), V836)

if True == tmp4231 {
__e.Return(False)
return
} else {
__e.Return(PrimIsString(V836))
return
}


}, 1)

tmp4232 := Call(__e, ns2_1set, symshen_4_7string_2, tmp4229)


_ = tmp4232

tmp4233 := MakeNative(func(__e *ControlFlow) {
V837 := __e.Get(1)
_ = V837
tmp4235 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp4236 := PrimEqual(V837, tmp4235)

if True == tmp4236 {
__e.Return(False)
return
} else {
__e.TailApply(PrimFunc(symvector_2), V837)
return
}


}, 1)

tmp4237 := Call(__e, ns2_1set, symshen_4_7vector_2, tmp4233)


_ = tmp4237

tmp4238 := MakeNative(func(__e *ControlFlow) {
V840 := __e.Get(1)
_ = V840
tmp4242 := PrimEqual(sym_7, V840)

if True == tmp4242 {
__e.Return(PrimSet(symshen_4_dfactorise_2_d, True))
return
} else {
tmp4240 := PrimEqual(sym_1, V840)

if True == tmp4240 {
__e.Return(PrimSet(symshen_4_dfactorise_2_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("factorise expects a + or a -\n")))
return
}


}


}, 1)

tmp4243 := Call(__e, ns2_1set, symfactorise, tmp4238)


_ = tmp4243

tmp4244 := MakeNative(func(__e *ControlFlow) {
V841 := __e.Get(1)
_ = V841
tmp4246 := PrimValue(symshen_4_dfactorise_2_d)

if True == tmp4246 {
__e.TailApply(PrimFunc(symshen_4factor), V841)
return
} else {
__e.Return(V841)
return
}


}, 1)

tmp4247 := Call(__e, ns2_1set, symshen_4factorise_1code, tmp4244)


_ = tmp4247

tmp4248 := MakeNative(func(__e *ControlFlow) {
V842 := __e.Get(1)
_ = V842
tmp4305 := PrimIsPair(V842)

var ifres4264 Obj

if True == tmp4305 {
tmp4303 := PrimHead(V842)

tmp4304 := PrimEqual(symdefun, tmp4303)

var ifres4266 Obj

if True == tmp4304 {
tmp4301 := PrimTail(V842)

tmp4302 := PrimIsPair(tmp4301)

var ifres4268 Obj

if True == tmp4302 {
tmp4298 := PrimTail(V842)

tmp4299 := PrimTail(tmp4298)

tmp4300 := PrimIsPair(tmp4299)

var ifres4270 Obj

if True == tmp4300 {
tmp4294 := PrimTail(V842)

tmp4295 := PrimTail(tmp4294)

tmp4296 := PrimTail(tmp4295)

tmp4297 := PrimIsPair(tmp4296)

var ifres4272 Obj

if True == tmp4297 {
tmp4289 := PrimTail(V842)

tmp4290 := PrimTail(tmp4289)

tmp4291 := PrimTail(tmp4290)

tmp4292 := PrimHead(tmp4291)

tmp4293 := PrimIsPair(tmp4292)

var ifres4274 Obj

if True == tmp4293 {
tmp4283 := PrimTail(V842)

tmp4284 := PrimTail(tmp4283)

tmp4285 := PrimTail(tmp4284)

tmp4286 := PrimHead(tmp4285)

tmp4287 := PrimHead(tmp4286)

tmp4288 := PrimEqual(symcond, tmp4287)

var ifres4276 Obj

if True == tmp4288 {
tmp4278 := PrimTail(V842)

tmp4279 := PrimTail(tmp4278)

tmp4280 := PrimTail(tmp4279)

tmp4281 := PrimTail(tmp4280)

tmp4282 := PrimEqual(Nil, tmp4281)

var ifres4277 Obj

if True == tmp4282 {
ifres4277 = True


} else {
ifres4277 = False


}

ifres4276 = ifres4277


} else {
ifres4276 = False


}

var ifres4275 Obj

if True == ifres4276 {
ifres4275 = True


} else {
ifres4275 = False


}

ifres4274 = ifres4275


} else {
ifres4274 = False


}

var ifres4273 Obj

if True == ifres4274 {
ifres4273 = True


} else {
ifres4273 = False


}

ifres4272 = ifres4273


} else {
ifres4272 = False


}

var ifres4271 Obj

if True == ifres4272 {
ifres4271 = True


} else {
ifres4271 = False


}

ifres4270 = ifres4271


} else {
ifres4270 = False


}

var ifres4269 Obj

if True == ifres4270 {
ifres4269 = True


} else {
ifres4269 = False


}

ifres4268 = ifres4269


} else {
ifres4268 = False


}

var ifres4267 Obj

if True == ifres4268 {
ifres4267 = True


} else {
ifres4267 = False


}

ifres4266 = ifres4267


} else {
ifres4266 = False


}

var ifres4265 Obj

if True == ifres4266 {
ifres4265 = True


} else {
ifres4265 = False


}

ifres4264 = ifres4265


} else {
ifres4264 = False


}

if True == ifres4264 {
tmp4249 := PrimTail(V842)

tmp4250 := PrimHead(tmp4249)

tmp4251 := PrimTail(V842)

tmp4252 := PrimTail(tmp4251)

tmp4253 := PrimHead(tmp4252)

tmp4254 := PrimTail(V842)

tmp4255 := PrimTail(tmp4254)

tmp4256 := PrimTail(tmp4255)

tmp4257 := PrimHead(tmp4256)

tmp4258 := PrimTail(tmp4257)

tmp4259 := Call(__e, PrimFunc(symshen_4factor_1recognisors), tmp4258)


tmp4260 := PrimCons(tmp4259, Nil)

tmp4261 := PrimCons(tmp4253, tmp4260)

tmp4262 := PrimCons(tmp4250, tmp4261)

__e.Return(PrimCons(symdefun, tmp4262))
return


} else {
__e.Return(V842)
return
}


}, 1)

tmp4306 := Call(__e, ns2_1set, symshen_4factor, tmp4248)


_ = tmp4306

tmp4307 := MakeNative(func(__e *ControlFlow) {
V845 := __e.Get(1)
_ = V845
tmp4463 := PrimIsPair(V845)

var ifres4443 Obj

if True == tmp4463 {
tmp4461 := PrimHead(V845)

tmp4462 := PrimIsPair(tmp4461)

var ifres4445 Obj

if True == tmp4462 {
tmp4458 := PrimHead(V845)

tmp4459 := PrimHead(tmp4458)

tmp4460 := PrimEqual(True, tmp4459)

var ifres4447 Obj

if True == tmp4460 {
tmp4455 := PrimHead(V845)

tmp4456 := PrimTail(tmp4455)

tmp4457 := PrimIsPair(tmp4456)

var ifres4449 Obj

if True == tmp4457 {
tmp4451 := PrimHead(V845)

tmp4452 := PrimTail(tmp4451)

tmp4453 := PrimTail(tmp4452)

tmp4454 := PrimEqual(Nil, tmp4453)

var ifres4450 Obj

if True == tmp4454 {
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

if True == ifres4443 {
tmp4308 := PrimHead(V845)

tmp4309 := PrimTail(tmp4308)

__e.Return(PrimHead(tmp4309))
return


} else {
tmp4441 := PrimIsPair(V845)

var ifres4394 Obj

if True == tmp4441 {
tmp4439 := PrimHead(V845)

tmp4440 := PrimIsPair(tmp4439)

var ifres4396 Obj

if True == tmp4440 {
tmp4436 := PrimHead(V845)

tmp4437 := PrimHead(tmp4436)

tmp4438 := PrimIsPair(tmp4437)

var ifres4398 Obj

if True == tmp4438 {
tmp4432 := PrimHead(V845)

tmp4433 := PrimHead(tmp4432)

tmp4434 := PrimHead(tmp4433)

tmp4435 := PrimEqual(symand, tmp4434)

var ifres4400 Obj

if True == tmp4435 {
tmp4428 := PrimHead(V845)

tmp4429 := PrimHead(tmp4428)

tmp4430 := PrimTail(tmp4429)

tmp4431 := PrimIsPair(tmp4430)

var ifres4402 Obj

if True == tmp4431 {
tmp4423 := PrimHead(V845)

tmp4424 := PrimHead(tmp4423)

tmp4425 := PrimTail(tmp4424)

tmp4426 := PrimTail(tmp4425)

tmp4427 := PrimIsPair(tmp4426)

var ifres4404 Obj

if True == tmp4427 {
tmp4417 := PrimHead(V845)

tmp4418 := PrimHead(tmp4417)

tmp4419 := PrimTail(tmp4418)

tmp4420 := PrimTail(tmp4419)

tmp4421 := PrimTail(tmp4420)

tmp4422 := PrimEqual(Nil, tmp4421)

var ifres4406 Obj

if True == tmp4422 {
tmp4414 := PrimHead(V845)

tmp4415 := PrimTail(tmp4414)

tmp4416 := PrimIsPair(tmp4415)

var ifres4408 Obj

if True == tmp4416 {
tmp4410 := PrimHead(V845)

tmp4411 := PrimTail(tmp4410)

tmp4412 := PrimTail(tmp4411)

tmp4413 := PrimEqual(Nil, tmp4412)

var ifres4409 Obj

if True == tmp4413 {
ifres4409 = True


} else {
ifres4409 = False


}

ifres4408 = ifres4409


} else {
ifres4408 = False


}

var ifres4407 Obj

if True == ifres4408 {
ifres4407 = True


} else {
ifres4407 = False


}

ifres4406 = ifres4407


} else {
ifres4406 = False


}

var ifres4405 Obj

if True == ifres4406 {
ifres4405 = True


} else {
ifres4405 = False


}

ifres4404 = ifres4405


} else {
ifres4404 = False


}

var ifres4403 Obj

if True == ifres4404 {
ifres4403 = True


} else {
ifres4403 = False


}

ifres4402 = ifres4403


} else {
ifres4402 = False


}

var ifres4401 Obj

if True == ifres4402 {
ifres4401 = True


} else {
ifres4401 = False


}

ifres4400 = ifres4401


} else {
ifres4400 = False


}

var ifres4399 Obj

if True == ifres4400 {
ifres4399 = True


} else {
ifres4399 = False


}

ifres4398 = ifres4399


} else {
ifres4398 = False


}

var ifres4397 Obj

if True == ifres4398 {
ifres4397 = True


} else {
ifres4397 = False


}

ifres4396 = ifres4397


} else {
ifres4396 = False


}

var ifres4395 Obj

if True == ifres4396 {
ifres4395 = True


} else {
ifres4395 = False


}

ifres4394 = ifres4395


} else {
ifres4394 = False


}

if True == ifres4394 {
tmp4310 := MakeNative(func(__e *ControlFlow) {
W846 := __e.Get(1)
_ = W846
tmp4311 := MakeNative(func(__e *ControlFlow) {
W847 := __e.Get(1)
_ = W847
tmp4359 := Call(__e, PrimFunc(symshen_4bad_1pivot_2), W847)


if True == tmp4359 {
tmp4312 := PrimHead(V845)

tmp4313 := PrimHead(tmp4312)

tmp4314 := PrimHead(V845)

tmp4315 := PrimTail(tmp4314)

tmp4316 := PrimHead(tmp4315)

tmp4317 := PrimTail(V845)

tmp4318 := Call(__e, PrimFunc(symshen_4factor_1recognisors), tmp4317)


tmp4319 := PrimCons(tmp4318, Nil)

tmp4320 := PrimCons(tmp4316, tmp4319)

tmp4321 := PrimCons(tmp4313, tmp4320)

__e.Return(PrimCons(symif, tmp4321))
return


} else {
tmp4322 := MakeNative(func(__e *ControlFlow) {
W848 := __e.Get(1)
_ = W848
tmp4323 := MakeNative(func(__e *ControlFlow) {
W849 := __e.Get(1)
_ = W849
tmp4324 := MakeNative(func(__e *ControlFlow) {
W850 := __e.Get(1)
_ = W850
tmp4325 := MakeNative(func(__e *ControlFlow) {
W851 := __e.Get(1)
_ = W851
tmp4326 := MakeNative(func(__e *ControlFlow) {
W852 := __e.Get(1)
_ = W852
__e.TailApply(PrimFunc(symshen_4remove_1indirection), W852)
return
}, 1)

tmp4327 := PrimCons(W849, Nil)

tmp4328 := PrimCons(symfreeze, tmp4327)

tmp4329 := PrimHead(V845)

tmp4330 := PrimHead(tmp4329)

tmp4331 := PrimTail(tmp4330)

tmp4332 := PrimHead(tmp4331)

tmp4333 := PrimHead(V845)

tmp4334 := PrimHead(tmp4333)

tmp4335 := PrimTail(tmp4334)

tmp4336 := PrimHead(tmp4335)

tmp4337 := Call(__e, PrimFunc(symshen_4factor_1recognisors), W851)


tmp4338 := Call(__e, PrimFunc(symshen_4factor_1selectors), tmp4336, tmp4337)


tmp4339 := PrimCons(W850, Nil)

tmp4340 := PrimCons(symthaw, tmp4339)

tmp4341 := PrimCons(tmp4340, Nil)

tmp4342 := PrimCons(tmp4338, tmp4341)

tmp4343 := PrimCons(tmp4332, tmp4342)

tmp4344 := PrimCons(symif, tmp4343)

tmp4345 := PrimCons(tmp4344, Nil)

tmp4346 := PrimCons(tmp4328, tmp4345)

tmp4347 := PrimCons(W850, tmp4346)

tmp4348 := PrimCons(symlet, tmp4347)

__e.TailApply(tmp4326, tmp4348)
return


}, 1)

tmp4349 := PrimCons(W850, Nil)

tmp4350 := PrimCons(symthaw, tmp4349)

tmp4351 := PrimCons(tmp4350, Nil)

tmp4352 := PrimCons(True, tmp4351)

tmp4353 := PrimCons(tmp4352, W847)

tmp4354 := Call(__e, PrimFunc(symreverse), tmp4353)


__e.TailApply(tmp4325, tmp4354)
return


}, 1)

tmp4355 := Call(__e, PrimFunc(symgensym), symGoTo)


__e.TailApply(tmp4324, tmp4355)
return


}, 1)

tmp4356 := Call(__e, PrimFunc(symshen_4factor_1recognisors), W848)


__e.TailApply(tmp4323, tmp4356)
return


}, 1)

tmp4357 := Call(__e, PrimFunc(symsnd), W846)


__e.TailApply(tmp4322, tmp4357)
return


}


}, 1)

tmp4360 := Call(__e, PrimFunc(symfst), W846)


__e.TailApply(tmp4311, tmp4360)
return


}, 1)

tmp4361 := PrimHead(V845)

tmp4362 := PrimHead(tmp4361)

tmp4363 := PrimTail(tmp4362)

tmp4364 := PrimHead(tmp4363)

tmp4365 := Call(__e, PrimFunc(symshen_4pivot_1on), tmp4364, V845, Nil)


__e.TailApply(tmp4310, tmp4365)
return


} else {
tmp4392 := PrimIsPair(V845)

var ifres4377 Obj

if True == tmp4392 {
tmp4390 := PrimHead(V845)

tmp4391 := PrimIsPair(tmp4390)

var ifres4379 Obj

if True == tmp4391 {
tmp4387 := PrimHead(V845)

tmp4388 := PrimTail(tmp4387)

tmp4389 := PrimIsPair(tmp4388)

var ifres4381 Obj

if True == tmp4389 {
tmp4383 := PrimHead(V845)

tmp4384 := PrimTail(tmp4383)

tmp4385 := PrimTail(tmp4384)

tmp4386 := PrimEqual(Nil, tmp4385)

var ifres4382 Obj

if True == tmp4386 {
ifres4382 = True


} else {
ifres4382 = False


}

ifres4381 = ifres4382


} else {
ifres4381 = False


}

var ifres4380 Obj

if True == ifres4381 {
ifres4380 = True


} else {
ifres4380 = False


}

ifres4379 = ifres4380


} else {
ifres4379 = False


}

var ifres4378 Obj

if True == ifres4379 {
ifres4378 = True


} else {
ifres4378 = False


}

ifres4377 = ifres4378


} else {
ifres4377 = False


}

if True == ifres4377 {
tmp4366 := PrimHead(V845)

tmp4367 := PrimHead(tmp4366)

tmp4368 := PrimHead(V845)

tmp4369 := PrimTail(tmp4368)

tmp4370 := PrimHead(tmp4369)

tmp4371 := PrimTail(V845)

tmp4372 := Call(__e, PrimFunc(symshen_4factor_1recognisors), tmp4371)


tmp4373 := PrimCons(tmp4372, Nil)

tmp4374 := PrimCons(tmp4370, tmp4373)

tmp4375 := PrimCons(tmp4367, tmp4374)

__e.Return(PrimCons(symif, tmp4375))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.factor-recognisors")))
return
}


}


}


}, 1)

tmp4464 := Call(__e, ns2_1set, symshen_4factor_1recognisors, tmp4307)


_ = tmp4464

tmp4465 := MakeNative(func(__e *ControlFlow) {
V857 := __e.Get(1)
_ = V857
tmp4471 := PrimIsPair(V857)

var ifres4467 Obj

if True == tmp4471 {
tmp4469 := PrimTail(V857)

tmp4470 := PrimEqual(Nil, tmp4469)

var ifres4468 Obj

if True == tmp4470 {
ifres4468 = True


} else {
ifres4468 = False


}

ifres4467 = ifres4468


} else {
ifres4467 = False


}

if True == ifres4467 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp4472 := Call(__e, ns2_1set, symshen_4bad_1pivot_2, tmp4465)


_ = tmp4472

tmp4473 := MakeNative(func(__e *ControlFlow) {
V858 := __e.Get(1)
_ = V858
tmp4588 := PrimIsPair(V858)

var ifres4488 Obj

if True == tmp4588 {
tmp4586 := PrimHead(V858)

tmp4587 := PrimEqual(symlet, tmp4586)

var ifres4490 Obj

if True == tmp4587 {
tmp4584 := PrimTail(V858)

tmp4585 := PrimIsPair(tmp4584)

var ifres4492 Obj

if True == tmp4585 {
tmp4581 := PrimTail(V858)

tmp4582 := PrimTail(tmp4581)

tmp4583 := PrimIsPair(tmp4582)

var ifres4494 Obj

if True == tmp4583 {
tmp4577 := PrimTail(V858)

tmp4578 := PrimTail(tmp4577)

tmp4579 := PrimHead(tmp4578)

tmp4580 := PrimIsPair(tmp4579)

var ifres4496 Obj

if True == tmp4580 {
tmp4572 := PrimTail(V858)

tmp4573 := PrimTail(tmp4572)

tmp4574 := PrimHead(tmp4573)

tmp4575 := PrimHead(tmp4574)

tmp4576 := PrimEqual(symfreeze, tmp4575)

var ifres4498 Obj

if True == tmp4576 {
tmp4567 := PrimTail(V858)

tmp4568 := PrimTail(tmp4567)

tmp4569 := PrimHead(tmp4568)

tmp4570 := PrimTail(tmp4569)

tmp4571 := PrimIsPair(tmp4570)

var ifres4500 Obj

if True == tmp4571 {
tmp4561 := PrimTail(V858)

tmp4562 := PrimTail(tmp4561)

tmp4563 := PrimHead(tmp4562)

tmp4564 := PrimTail(tmp4563)

tmp4565 := PrimHead(tmp4564)

tmp4566 := PrimIsPair(tmp4565)

var ifres4502 Obj

if True == tmp4566 {
tmp4554 := PrimTail(V858)

tmp4555 := PrimTail(tmp4554)

tmp4556 := PrimHead(tmp4555)

tmp4557 := PrimTail(tmp4556)

tmp4558 := PrimHead(tmp4557)

tmp4559 := PrimHead(tmp4558)

tmp4560 := PrimEqual(symthaw, tmp4559)

var ifres4504 Obj

if True == tmp4560 {
tmp4547 := PrimTail(V858)

tmp4548 := PrimTail(tmp4547)

tmp4549 := PrimHead(tmp4548)

tmp4550 := PrimTail(tmp4549)

tmp4551 := PrimHead(tmp4550)

tmp4552 := PrimTail(tmp4551)

tmp4553 := PrimIsPair(tmp4552)

var ifres4506 Obj

if True == tmp4553 {
tmp4539 := PrimTail(V858)

tmp4540 := PrimTail(tmp4539)

tmp4541 := PrimHead(tmp4540)

tmp4542 := PrimTail(tmp4541)

tmp4543 := PrimHead(tmp4542)

tmp4544 := PrimTail(tmp4543)

tmp4545 := PrimTail(tmp4544)

tmp4546 := PrimEqual(Nil, tmp4545)

var ifres4508 Obj

if True == tmp4546 {
tmp4533 := PrimTail(V858)

tmp4534 := PrimTail(tmp4533)

tmp4535 := PrimHead(tmp4534)

tmp4536 := PrimTail(tmp4535)

tmp4537 := PrimTail(tmp4536)

tmp4538 := PrimEqual(Nil, tmp4537)

var ifres4510 Obj

if True == tmp4538 {
tmp4529 := PrimTail(V858)

tmp4530 := PrimTail(tmp4529)

tmp4531 := PrimTail(tmp4530)

tmp4532 := PrimIsPair(tmp4531)

var ifres4512 Obj

if True == tmp4532 {
tmp4524 := PrimTail(V858)

tmp4525 := PrimTail(tmp4524)

tmp4526 := PrimTail(tmp4525)

tmp4527 := PrimTail(tmp4526)

tmp4528 := PrimEqual(Nil, tmp4527)

var ifres4514 Obj

if True == tmp4528 {
tmp4516 := PrimTail(V858)

tmp4517 := PrimTail(tmp4516)

tmp4518 := PrimHead(tmp4517)

tmp4519 := PrimTail(tmp4518)

tmp4520 := PrimHead(tmp4519)

tmp4521 := PrimTail(tmp4520)

tmp4522 := PrimHead(tmp4521)

tmp4523 := PrimIsSymbol(tmp4522)

var ifres4515 Obj

if True == tmp4523 {
ifres4515 = True


} else {
ifres4515 = False


}

ifres4514 = ifres4515


} else {
ifres4514 = False


}

var ifres4513 Obj

if True == ifres4514 {
ifres4513 = True


} else {
ifres4513 = False


}

ifres4512 = ifres4513


} else {
ifres4512 = False


}

var ifres4511 Obj

if True == ifres4512 {
ifres4511 = True


} else {
ifres4511 = False


}

ifres4510 = ifres4511


} else {
ifres4510 = False


}

var ifres4509 Obj

if True == ifres4510 {
ifres4509 = True


} else {
ifres4509 = False


}

ifres4508 = ifres4509


} else {
ifres4508 = False


}

var ifres4507 Obj

if True == ifres4508 {
ifres4507 = True


} else {
ifres4507 = False


}

ifres4506 = ifres4507


} else {
ifres4506 = False


}

var ifres4505 Obj

if True == ifres4506 {
ifres4505 = True


} else {
ifres4505 = False


}

ifres4504 = ifres4505


} else {
ifres4504 = False


}

var ifres4503 Obj

if True == ifres4504 {
ifres4503 = True


} else {
ifres4503 = False


}

ifres4502 = ifres4503


} else {
ifres4502 = False


}

var ifres4501 Obj

if True == ifres4502 {
ifres4501 = True


} else {
ifres4501 = False


}

ifres4500 = ifres4501


} else {
ifres4500 = False


}

var ifres4499 Obj

if True == ifres4500 {
ifres4499 = True


} else {
ifres4499 = False


}

ifres4498 = ifres4499


} else {
ifres4498 = False


}

var ifres4497 Obj

if True == ifres4498 {
ifres4497 = True


} else {
ifres4497 = False


}

ifres4496 = ifres4497


} else {
ifres4496 = False


}

var ifres4495 Obj

if True == ifres4496 {
ifres4495 = True


} else {
ifres4495 = False


}

ifres4494 = ifres4495


} else {
ifres4494 = False


}

var ifres4493 Obj

if True == ifres4494 {
ifres4493 = True


} else {
ifres4493 = False


}

ifres4492 = ifres4493


} else {
ifres4492 = False


}

var ifres4491 Obj

if True == ifres4492 {
ifres4491 = True


} else {
ifres4491 = False


}

ifres4490 = ifres4491


} else {
ifres4490 = False


}

var ifres4489 Obj

if True == ifres4490 {
ifres4489 = True


} else {
ifres4489 = False


}

ifres4488 = ifres4489


} else {
ifres4488 = False


}

if True == ifres4488 {
tmp4474 := PrimTail(V858)

tmp4475 := PrimTail(tmp4474)

tmp4476 := PrimHead(tmp4475)

tmp4477 := PrimTail(tmp4476)

tmp4478 := PrimHead(tmp4477)

tmp4479 := PrimTail(tmp4478)

tmp4480 := PrimHead(tmp4479)

tmp4481 := PrimTail(V858)

tmp4482 := PrimHead(tmp4481)

tmp4483 := PrimTail(V858)

tmp4484 := PrimTail(tmp4483)

tmp4485 := PrimTail(tmp4484)

tmp4486 := PrimHead(tmp4485)

__e.TailApply(PrimFunc(symsubst), tmp4480, tmp4482, tmp4486)
return


} else {
__e.Return(V858)
return
}


}, 1)

tmp4589 := Call(__e, ns2_1set, symshen_4remove_1indirection, tmp4473)


_ = tmp4589

tmp4590 := MakeNative(func(__e *ControlFlow) {
V861 := __e.Get(1)
_ = V861
V862 := __e.Get(2)
_ = V862
V863 := __e.Get(3)
_ = V863
tmp4689 := PrimIsPair(V862)

var ifres4635 Obj

if True == tmp4689 {
tmp4687 := PrimHead(V862)

tmp4688 := PrimIsPair(tmp4687)

var ifres4637 Obj

if True == tmp4688 {
tmp4684 := PrimHead(V862)

tmp4685 := PrimHead(tmp4684)

tmp4686 := PrimIsPair(tmp4685)

var ifres4639 Obj

if True == tmp4686 {
tmp4680 := PrimHead(V862)

tmp4681 := PrimHead(tmp4680)

tmp4682 := PrimHead(tmp4681)

tmp4683 := PrimEqual(symand, tmp4682)

var ifres4641 Obj

if True == tmp4683 {
tmp4676 := PrimHead(V862)

tmp4677 := PrimHead(tmp4676)

tmp4678 := PrimTail(tmp4677)

tmp4679 := PrimIsPair(tmp4678)

var ifres4643 Obj

if True == tmp4679 {
tmp4671 := PrimHead(V862)

tmp4672 := PrimHead(tmp4671)

tmp4673 := PrimTail(tmp4672)

tmp4674 := PrimTail(tmp4673)

tmp4675 := PrimIsPair(tmp4674)

var ifres4645 Obj

if True == tmp4675 {
tmp4665 := PrimHead(V862)

tmp4666 := PrimHead(tmp4665)

tmp4667 := PrimTail(tmp4666)

tmp4668 := PrimTail(tmp4667)

tmp4669 := PrimTail(tmp4668)

tmp4670 := PrimEqual(Nil, tmp4669)

var ifres4647 Obj

if True == tmp4670 {
tmp4662 := PrimHead(V862)

tmp4663 := PrimTail(tmp4662)

tmp4664 := PrimIsPair(tmp4663)

var ifres4649 Obj

if True == tmp4664 {
tmp4658 := PrimHead(V862)

tmp4659 := PrimTail(tmp4658)

tmp4660 := PrimTail(tmp4659)

tmp4661 := PrimEqual(Nil, tmp4660)

var ifres4651 Obj

if True == tmp4661 {
tmp4653 := PrimHead(V862)

tmp4654 := PrimHead(tmp4653)

tmp4655 := PrimTail(tmp4654)

tmp4656 := PrimHead(tmp4655)

tmp4657 := PrimEqual(V861, tmp4656)

var ifres4652 Obj

if True == tmp4657 {
ifres4652 = True


} else {
ifres4652 = False


}

ifres4651 = ifres4652


} else {
ifres4651 = False


}

var ifres4650 Obj

if True == ifres4651 {
ifres4650 = True


} else {
ifres4650 = False


}

ifres4649 = ifres4650


} else {
ifres4649 = False


}

var ifres4648 Obj

if True == ifres4649 {
ifres4648 = True


} else {
ifres4648 = False


}

ifres4647 = ifres4648


} else {
ifres4647 = False


}

var ifres4646 Obj

if True == ifres4647 {
ifres4646 = True


} else {
ifres4646 = False


}

ifres4645 = ifres4646


} else {
ifres4645 = False


}

var ifres4644 Obj

if True == ifres4645 {
ifres4644 = True


} else {
ifres4644 = False


}

ifres4643 = ifres4644


} else {
ifres4643 = False


}

var ifres4642 Obj

if True == ifres4643 {
ifres4642 = True


} else {
ifres4642 = False


}

ifres4641 = ifres4642


} else {
ifres4641 = False


}

var ifres4640 Obj

if True == ifres4641 {
ifres4640 = True


} else {
ifres4640 = False


}

ifres4639 = ifres4640


} else {
ifres4639 = False


}

var ifres4638 Obj

if True == ifres4639 {
ifres4638 = True


} else {
ifres4638 = False


}

ifres4637 = ifres4638


} else {
ifres4637 = False


}

var ifres4636 Obj

if True == ifres4637 {
ifres4636 = True


} else {
ifres4636 = False


}

ifres4635 = ifres4636


} else {
ifres4635 = False


}

if True == ifres4635 {
tmp4591 := PrimHead(V862)

tmp4592 := PrimHead(tmp4591)

tmp4593 := PrimTail(tmp4592)

tmp4594 := PrimHead(tmp4593)

tmp4595 := PrimTail(V862)

tmp4596 := PrimHead(V862)

tmp4597 := PrimHead(tmp4596)

tmp4598 := PrimTail(tmp4597)

tmp4599 := PrimTail(tmp4598)

tmp4600 := PrimHead(tmp4599)

tmp4601 := PrimHead(V862)

tmp4602 := PrimTail(tmp4601)

tmp4603 := PrimCons(tmp4600, tmp4602)

tmp4604 := PrimCons(tmp4603, V863)

__e.TailApply(PrimFunc(symshen_4pivot_1on), tmp4594, tmp4595, tmp4604)
return


} else {
tmp4633 := PrimIsPair(V862)

var ifres4613 Obj

if True == tmp4633 {
tmp4631 := PrimHead(V862)

tmp4632 := PrimIsPair(tmp4631)

var ifres4615 Obj

if True == tmp4632 {
tmp4628 := PrimHead(V862)

tmp4629 := PrimTail(tmp4628)

tmp4630 := PrimIsPair(tmp4629)

var ifres4617 Obj

if True == tmp4630 {
tmp4624 := PrimHead(V862)

tmp4625 := PrimTail(tmp4624)

tmp4626 := PrimTail(tmp4625)

tmp4627 := PrimEqual(Nil, tmp4626)

var ifres4619 Obj

if True == tmp4627 {
tmp4621 := PrimHead(V862)

tmp4622 := PrimHead(tmp4621)

tmp4623 := PrimEqual(V861, tmp4622)

var ifres4620 Obj

if True == tmp4623 {
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

var ifres4616 Obj

if True == ifres4617 {
ifres4616 = True


} else {
ifres4616 = False


}

ifres4615 = ifres4616


} else {
ifres4615 = False


}

var ifres4614 Obj

if True == ifres4615 {
ifres4614 = True


} else {
ifres4614 = False


}

ifres4613 = ifres4614


} else {
ifres4613 = False


}

if True == ifres4613 {
tmp4605 := PrimHead(V862)

tmp4606 := PrimHead(tmp4605)

tmp4607 := PrimTail(V862)

tmp4608 := PrimHead(V862)

tmp4609 := PrimTail(tmp4608)

tmp4610 := PrimCons(True, tmp4609)

tmp4611 := PrimCons(tmp4610, V863)

__e.TailApply(PrimFunc(symshen_4pivot_1on), tmp4606, tmp4607, tmp4611)
return


} else {
__e.TailApply(PrimFunc(sym_8p), V863, V862)
return
}


}


}, 3)

tmp4690 := Call(__e, ns2_1set, symshen_4pivot_1on, tmp4590)


_ = tmp4690

tmp4691 := MakeNative(func(__e *ControlFlow) {
V866 := __e.Get(1)
_ = V866
V867 := __e.Get(2)
_ = V867
tmp4715 := PrimIsPair(V866)

var ifres4706 Obj

if True == tmp4715 {
tmp4713 := PrimTail(V866)

tmp4714 := PrimIsPair(tmp4713)

var ifres4708 Obj

if True == tmp4714 {
tmp4710 := PrimTail(V866)

tmp4711 := PrimTail(tmp4710)

tmp4712 := PrimEqual(Nil, tmp4711)

var ifres4709 Obj

if True == tmp4712 {
ifres4709 = True


} else {
ifres4709 = False


}

ifres4708 = ifres4709


} else {
ifres4708 = False


}

var ifres4707 Obj

if True == ifres4708 {
ifres4707 = True


} else {
ifres4707 = False


}

ifres4706 = ifres4707


} else {
ifres4706 = False


}

if True == ifres4706 {
tmp4692 := MakeNative(func(__e *ControlFlow) {
W868 := __e.Get(1)
_ = W868
tmp4702 := PrimEqual(symshen_4skip, W868)

if True == tmp4702 {
__e.Return(V867)
return
} else {
tmp4693 := Call(__e, PrimFunc(symshen_4op1), W868)


tmp4694 := PrimTail(V866)

tmp4695 := PrimCons(tmp4693, tmp4694)

tmp4696 := Call(__e, PrimFunc(symshen_4op2), W868)


tmp4697 := PrimTail(V866)

tmp4698 := PrimCons(tmp4696, tmp4697)

tmp4699 := PrimCons(tmp4698, Nil)

tmp4700 := PrimCons(tmp4695, tmp4699)

__e.TailApply(PrimFunc(symshen_4factor_1selectors_1h), tmp4700, V867)
return


}


}, 1)

tmp4703 := PrimHead(V866)

tmp4704 := Call(__e, PrimFunc(symshen_4op), tmp4703)


__e.TailApply(tmp4692, tmp4704)
return


} else {
__e.Return(V867)
return
}


}, 2)

tmp4716 := Call(__e, ns2_1set, symshen_4factor_1selectors, tmp4691)


_ = tmp4716

tmp4717 := MakeNative(func(__e *ControlFlow) {
V871 := __e.Get(1)
_ = V871
tmp4725 := PrimEqual(symcons_2, V871)

if True == tmp4725 {
__e.Return(symcons)
return
} else {
tmp4723 := PrimEqual(symshen_4_7string_2, V871)

if True == tmp4723 {
__e.Return(sym_8s)
return
} else {
tmp4721 := PrimEqual(symshen_4_7vector_2, V871)

if True == tmp4721 {
__e.Return(sym_8v)
return
} else {
tmp4719 := PrimEqual(symtuple_2, V871)

if True == tmp4719 {
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

tmp4726 := Call(__e, ns2_1set, symshen_4op, tmp4717)


_ = tmp4726

tmp4727 := MakeNative(func(__e *ControlFlow) {
V872 := __e.Get(1)
_ = V872
V873 := __e.Get(2)
_ = V873
tmp4746 := PrimEqual(Nil, V872)

if True == tmp4746 {
__e.Return(V873)
return
} else {
tmp4744 := PrimIsPair(V872)

if True == tmp4744 {
tmp4740 := PrimHead(V872)

tmp4741 := Call(__e, PrimFunc(symoccurrences), tmp4740, V873)


tmp4742 := PrimGreatThan(tmp4741, MakeNumber(1))

if True == tmp4742 {
tmp4728 := MakeNative(func(__e *ControlFlow) {
W874 := __e.Get(1)
_ = W874
tmp4729 := PrimHead(V872)

tmp4730 := PrimTail(V872)

tmp4731 := PrimHead(V872)

tmp4732 := Call(__e, PrimFunc(symsubst), W874, tmp4731, V873)


tmp4733 := Call(__e, PrimFunc(symshen_4factor_1selectors_1h), tmp4730, tmp4732)


tmp4734 := PrimCons(tmp4733, Nil)

tmp4735 := PrimCons(tmp4729, tmp4734)

tmp4736 := PrimCons(W874, tmp4735)

__e.Return(PrimCons(symlet, tmp4736))
return


}, 1)

tmp4737 := Call(__e, PrimFunc(symgensym), symSelect)


__e.TailApply(tmp4728, tmp4737)
return


} else {
tmp4738 := PrimTail(V872)

__e.TailApply(PrimFunc(symshen_4factor_1selectors_1h), tmp4738, V873)
return


}


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.factor-selectors-h")))
return
}


}


}, 2)

__e.TailApply(ns2_1set, symshen_4factor_1selectors_1h, tmp4727)
return




}, 0)

