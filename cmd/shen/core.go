package main

import . "github.com/pyrex41/shen-go/kl"

var CoreMain = MakeNative(func(__e *ControlFlow) {
tmp2611 := MakeNative(func(__e *ControlFlow) {
V528 := __e.Get(1)
_ = V528
tmp2612 := MakeNative(func(__e *ControlFlow) {
W529 := __e.Get(1)
_ = W529
__e.TailApply(PrimFunc(symshen_4record_1and_1evaluate), W529)
return
}, 1)

tmp2613 := Call(__e, PrimFunc(symshen_4shen_1_6kl_1h), V528)


__e.TailApply(tmp2612, tmp2613)
return


}, 1)

tmp2614 := Call(__e, ns2_1set, symshen_4shen_1_6kl, tmp2611)


_ = tmp2614

tmp2615 := MakeNative(func(__e *ControlFlow) {
V530 := __e.Get(1)
_ = V530
tmp2668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V530)
}
__typedArg0 := V530
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2642 Obj

if True == tmp2668 {
tmp2666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V530)
}
__typedArg0 := V530
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp2666)
}
__typedArg0 := symdefun
__typedArg1 := tmp2666
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2644 Obj

if True == tmp2667 {
tmp2664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V530)
}
__typedArg0 := V530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2664)
}
__typedArg0 := tmp2664
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2646 Obj

if True == tmp2665 {
tmp2661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V530)
}
__typedArg0 := V530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2661)
}
__typedArg0 := tmp2661
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2662)
}
__typedArg0 := tmp2662
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2648 Obj

if True == tmp2663 {
tmp2657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V530)
}
__typedArg0 := V530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2657)
}
__typedArg0 := tmp2657
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2658)
}
__typedArg0 := tmp2658
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2659)
}
__typedArg0 := tmp2659
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2650 Obj

if True == tmp2660 {
tmp2652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V530)
}
__typedArg0 := V530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2652)
}
__typedArg0 := tmp2652
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2653)
}
__typedArg0 := tmp2653
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2654)
}
__typedArg0 := tmp2654
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2655)
}
__typedArg0 := Nil
__typedArg1 := tmp2655
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2651 Obj

if True == tmp2656 {
ifres2651 = True


} else {
ifres2651 = False


}

ifres2650 = ifres2651


} else {
ifres2650 = False


}

var ifres2649 Obj

if True == ifres2650 {
ifres2649 = True


} else {
ifres2649 = False


}

ifres2648 = ifres2649


} else {
ifres2648 = False


}

var ifres2647 Obj

if True == ifres2648 {
ifres2647 = True


} else {
ifres2647 = False


}

ifres2646 = ifres2647


} else {
ifres2646 = False


}

var ifres2645 Obj

if True == ifres2646 {
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

if True == ifres2642 {
tmp2616 := MakeNative(func(__e *ControlFlow) {
W531 := __e.Get(1)
_ = W531
tmp2617 := MakeNative(func(__e *ControlFlow) {
W532 := __e.Get(1)
_ = W532
tmp2618 := MakeNative(func(__e *ControlFlow) {
W533 := __e.Get(1)
_ = W533
tmp2619 := MakeNative(func(__e *ControlFlow) {
W534 := __e.Get(1)
_ = W534
tmp2620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V530)
}
__typedArg0 := V530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2620)
}
__typedArg0 := tmp2620
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4fn_1print), tmp2621)
return


}, 1)

tmp2622 := Call(__e, PrimFunc(symeval_1kl), V530)


__e.TailApply(tmp2619, tmp2622)
return


}, 1)

tmp2623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V530)
}
__typedArg0 := V530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2623)
}
__typedArg0 := tmp2623
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2625 := Call(__e, PrimFunc(symshen_4record_1kl), tmp2624, V530)


__e.TailApply(tmp2618, tmp2625)
return


}, 1)

tmp2626 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V530)
}
__typedArg0 := V530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2627 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2626)
}
__typedArg0 := tmp2626
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2628 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V530)
}
__typedArg0 := V530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2628)
}
__typedArg0 := tmp2628
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2630 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2629)
}
__typedArg0 := tmp2629
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2631 := Call(__e, PrimFunc(symlength), tmp2630)


tmp2632 := Call(__e, PrimFunc(symshen_4store_1arity), tmp2627, tmp2631)


__e.TailApply(tmp2617, tmp2632)
return


}, 1)

tmp2638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V530)
}
__typedArg0 := V530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2638)
}
__typedArg0 := tmp2638
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2640 := Call(__e, PrimFunc(symshen_4sysfunc_2), tmp2639)


var ifres2633 Obj

if True == tmp2640 {
tmp2634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V530)
}
__typedArg0 := V530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2635 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2634)
}
__typedArg0 := tmp2634
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2636 := Call(__e, PrimFunc(symshen_4app), tmp2635, MakeString(" is not a legitimate function name\n"), symshen_4a)


tmp2637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp2636)
}
__typedArg0 := tmp2636
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres2633 = tmp2637


} else {
ifres2633 = symshen_4skip


}

__e.TailApply(tmp2616, ifres2633)
return


} else {
__e.Return(V530)
return
}


}, 1)

tmp2669 := Call(__e, ns2_1set, symshen_4record_1and_1evaluate, tmp2615)


_ = tmp2669

tmp2670 := MakeNative(func(__e *ControlFlow) {
V535 := __e.Get(1)
_ = V535
tmp2771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2763 Obj

if True == tmp2771 {
tmp2769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefine, tmp2769)
}
__typedArg0 := symdefine
__typedArg1 := tmp2769
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2765 Obj

if True == tmp2770 {
tmp2767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2767)
}
__typedArg0 := tmp2767
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2766 Obj

if True == tmp2768 {
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

if True == ifres2763 {
tmp2671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2671)
}
__typedArg0 := tmp2671
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2673)
}
__typedArg0 := tmp2673
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4shendef_1_6kldef), tmp2672, tmp2674)
return


} else {
tmp2761 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2735 Obj

if True == tmp2761 {
tmp2759 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2760 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp2759)
}
__typedArg0 := symdefun
__typedArg1 := tmp2759
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2737 Obj

if True == tmp2760 {
tmp2757 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2758 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2757)
}
__typedArg0 := tmp2757
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2739 Obj

if True == tmp2758 {
tmp2754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2754)
}
__typedArg0 := tmp2754
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2756 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2755)
}
__typedArg0 := tmp2755
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2741 Obj

if True == tmp2756 {
tmp2750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2750)
}
__typedArg0 := tmp2750
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2751)
}
__typedArg0 := tmp2751
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2752)
}
__typedArg0 := tmp2752
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2743 Obj

if True == tmp2753 {
tmp2745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2745)
}
__typedArg0 := tmp2745
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2746)
}
__typedArg0 := tmp2746
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2748)
}
__typedArg0 := Nil
__typedArg1 := tmp2748
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2744 Obj

if True == tmp2749 {
ifres2744 = True


} else {
ifres2744 = False


}

ifres2743 = ifres2744


} else {
ifres2743 = False


}

var ifres2742 Obj

if True == ifres2743 {
ifres2742 = True


} else {
ifres2742 = False


}

ifres2741 = ifres2742


} else {
ifres2741 = False


}

var ifres2740 Obj

if True == ifres2741 {
ifres2740 = True


} else {
ifres2740 = False


}

ifres2739 = ifres2740


} else {
ifres2739 = False


}

var ifres2738 Obj

if True == ifres2739 {
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

if True == ifres2735 {
__e.Return(V535)
return
} else {
tmp2733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2714 Obj

if True == tmp2733 {
tmp2731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtype, tmp2731)
}
__typedArg0 := symtype
__typedArg1 := tmp2731
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2716 Obj

if True == tmp2732 {
tmp2729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2729)
}
__typedArg0 := tmp2729
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2718 Obj

if True == tmp2730 {
tmp2726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2726)
}
__typedArg0 := tmp2726
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2727)
}
__typedArg0 := tmp2727
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2720 Obj

if True == tmp2728 {
tmp2722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2722)
}
__typedArg0 := tmp2722
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2723)
}
__typedArg0 := tmp2723
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2724)
}
__typedArg0 := Nil
__typedArg1 := tmp2724
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2721 Obj

if True == tmp2725 {
ifres2721 = True


} else {
ifres2721 = False


}

ifres2720 = ifres2721


} else {
ifres2720 = False


}

var ifres2719 Obj

if True == ifres2720 {
ifres2719 = True


} else {
ifres2719 = False


}

ifres2718 = ifres2719


} else {
ifres2718 = False


}

var ifres2717 Obj

if True == ifres2718 {
ifres2717 = True


} else {
ifres2717 = False


}

ifres2716 = ifres2717


} else {
ifres2716 = False


}

var ifres2715 Obj

if True == ifres2716 {
ifres2715 = True


} else {
ifres2715 = False


}

ifres2714 = ifres2715


} else {
ifres2714 = False


}

if True == ifres2714 {
tmp2675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2675)
}
__typedArg0 := tmp2675
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2677)
}
__typedArg0 := tmp2677
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2678)
}
__typedArg0 := tmp2678
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2680 := Call(__e, PrimFunc(symshen_4rcons__form), tmp2679)


tmp2681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2680, Nil)
}
__typedArg0 := tmp2680
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2676, tmp2681)
}
__typedArg0 := tmp2676
__typedArg1 := tmp2681
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symtype, tmp2682)
}
__typedArg0 := symtype
__typedArg1 := tmp2682
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp2712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2693 Obj

if True == tmp2712 {
tmp2710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(syminput_7, tmp2710)
}
__typedArg0 := syminput_7
__typedArg1 := tmp2710
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2695 Obj

if True == tmp2711 {
tmp2708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2708)
}
__typedArg0 := tmp2708
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2697 Obj

if True == tmp2709 {
tmp2705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2705)
}
__typedArg0 := tmp2705
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2706)
}
__typedArg0 := tmp2706
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2699 Obj

if True == tmp2707 {
tmp2701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2701)
}
__typedArg0 := tmp2701
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2702)
}
__typedArg0 := tmp2702
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2703)
}
__typedArg0 := Nil
__typedArg1 := tmp2703
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2700 Obj

if True == tmp2704 {
ifres2700 = True


} else {
ifres2700 = False


}

ifres2699 = ifres2700


} else {
ifres2699 = False


}

var ifres2698 Obj

if True == ifres2699 {
ifres2698 = True


} else {
ifres2698 = False


}

ifres2697 = ifres2698


} else {
ifres2697 = False


}

var ifres2696 Obj

if True == ifres2697 {
ifres2696 = True


} else {
ifres2696 = False


}

ifres2695 = ifres2696


} else {
ifres2695 = False


}

var ifres2694 Obj

if True == ifres2695 {
ifres2694 = True


} else {
ifres2694 = False


}

ifres2693 = ifres2694


} else {
ifres2693 = False


}

if True == ifres2693 {
tmp2683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2683)
}
__typedArg0 := tmp2683
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2685 := Call(__e, PrimFunc(symshen_4rcons__form), tmp2684)


tmp2686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp2686)
}
__typedArg0 := tmp2686
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2685, tmp2687)
}
__typedArg0 := tmp2685
__typedArg1 := tmp2687
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(syminput_7, tmp2688)
}
__typedArg0 := syminput_7
__typedArg1 := tmp2688
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp2691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V535)
}
__typedArg0 := V535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2691 {
tmp2689 := MakeNative(func(__e *ControlFlow) {
Z536 := __e.Get(1)
_ = Z536
__e.TailApply(PrimFunc(symshen_4shen_1_6kl_1h), Z536)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp2689, V535)
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

tmp2772 := Call(__e, ns2_1set, symshen_4shen_1_6kl_1h, tmp2670)


_ = tmp2772

tmp2773 := MakeNative(func(__e *ControlFlow) {
V537 := __e.Get(1)
_ = V537
V538 := __e.Get(2)
_ = V538
tmp2774 := MakeNative(func(__e *ControlFlow) {
Z539 := __e.Get(1)
_ = Z539
__e.TailApply(PrimFunc(symshen_4_5define_6), Z539)
return
}, 1)

tmp2775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V537, V538)
}
__typedArg0 := V537
__typedArg1 := V538
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symcompile), tmp2774, tmp2775)
return


}, 2)

tmp2776 := Call(__e, ns2_1set, symshen_4shendef_1_6kldef, tmp2773)


_ = tmp2776

tmp2777 := MakeNative(func(__e *ControlFlow) {
V540 := __e.Get(1)
_ = V540
tmp2778 := MakeNative(func(__e *ControlFlow) {
W541 := __e.Get(1)
_ = W541
tmp2801 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W541)


if True == tmp2801 {
tmp2779 := MakeNative(func(__e *ControlFlow) {
W552 := __e.Get(1)
_ = W552
tmp2781 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W552)


if True == tmp2781 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W552)
return
}


}, 1)

tmp2782 := MakeNative(func(__e *ControlFlow) {
W553 := __e.Get(1)
_ = W553
tmp2797 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W553)


if True == tmp2797 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2783 := MakeNative(func(__e *ControlFlow) {
W554 := __e.Get(1)
_ = W554
tmp2784 := MakeNative(func(__e *ControlFlow) {
W555 := __e.Get(1)
_ = W555
tmp2785 := MakeNative(func(__e *ControlFlow) {
W556 := __e.Get(1)
_ = W556
tmp2792 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W556)


if True == tmp2792 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2786 := MakeNative(func(__e *ControlFlow) {
W557 := __e.Get(1)
_ = W557
tmp2787 := MakeNative(func(__e *ControlFlow) {
W558 := __e.Get(1)
_ = W558
tmp2788 := Call(__e, PrimFunc(symshen_4shendef_1_6kldef_1h), W554, W557)


__e.TailApply(PrimFunc(symshen_4comb), W558, tmp2788)
return


}, 1)

tmp2789 := Call(__e, PrimFunc(symshen_4in_1_6), W556)


__e.TailApply(tmp2787, tmp2789)
return


}, 1)

tmp2790 := Call(__e, PrimFunc(symshen_4_5_1out), W556)


__e.TailApply(tmp2786, tmp2790)
return


}


}, 1)

tmp2793 := Call(__e, PrimFunc(symshen_4_5rules_6), W555)


__e.TailApply(tmp2785, tmp2793)
return


}, 1)

tmp2794 := Call(__e, PrimFunc(symshen_4in_1_6), W553)


__e.TailApply(tmp2784, tmp2794)
return


}, 1)

tmp2795 := Call(__e, PrimFunc(symshen_4_5_1out), W553)


__e.TailApply(tmp2783, tmp2795)
return


}


}, 1)

tmp2798 := Call(__e, PrimFunc(symshen_4_5name_6), V540)


tmp2799 := Call(__e, tmp2782, tmp2798)


__e.TailApply(tmp2779, tmp2799)
return


} else {
__e.Return(W541)
return
}


}, 1)

tmp2802 := MakeNative(func(__e *ControlFlow) {
W542 := __e.Get(1)
_ = W542
tmp2831 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W542)


if True == tmp2831 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2803 := MakeNative(func(__e *ControlFlow) {
W543 := __e.Get(1)
_ = W543
tmp2804 := MakeNative(func(__e *ControlFlow) {
W544 := __e.Get(1)
_ = W544
tmp2827 := Call(__e, PrimFunc(symshen_4hds_a_2), W544, sym_i)


if True == tmp2827 {
tmp2805 := MakeNative(func(__e *ControlFlow) {
W545 := __e.Get(1)
_ = W545
tmp2806 := MakeNative(func(__e *ControlFlow) {
W546 := __e.Get(1)
_ = W546
tmp2823 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W546)


if True == tmp2823 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2807 := MakeNative(func(__e *ControlFlow) {
W547 := __e.Get(1)
_ = W547
tmp2820 := Call(__e, PrimFunc(symshen_4hds_a_2), W547, sym_j)


if True == tmp2820 {
tmp2808 := MakeNative(func(__e *ControlFlow) {
W548 := __e.Get(1)
_ = W548
tmp2809 := MakeNative(func(__e *ControlFlow) {
W549 := __e.Get(1)
_ = W549
tmp2816 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W549)


if True == tmp2816 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2810 := MakeNative(func(__e *ControlFlow) {
W550 := __e.Get(1)
_ = W550
tmp2811 := MakeNative(func(__e *ControlFlow) {
W551 := __e.Get(1)
_ = W551
tmp2812 := Call(__e, PrimFunc(symshen_4shendef_1_6kldef_1h), W543, W550)


__e.TailApply(PrimFunc(symshen_4comb), W551, tmp2812)
return


}, 1)

tmp2813 := Call(__e, PrimFunc(symshen_4in_1_6), W549)


__e.TailApply(tmp2811, tmp2813)
return


}, 1)

tmp2814 := Call(__e, PrimFunc(symshen_4_5_1out), W549)


__e.TailApply(tmp2810, tmp2814)
return


}


}, 1)

tmp2817 := Call(__e, PrimFunc(symshen_4_5rules_6), W548)


__e.TailApply(tmp2809, tmp2817)
return


}, 1)

tmp2818 := Call(__e, PrimFunc(symtail), W547)


__e.TailApply(tmp2808, tmp2818)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2821 := Call(__e, PrimFunc(symshen_4in_1_6), W546)


__e.TailApply(tmp2807, tmp2821)
return


}


}, 1)

tmp2824 := Call(__e, PrimFunc(symshen_4_5signature_6), W545)


__e.TailApply(tmp2806, tmp2824)
return


}, 1)

tmp2825 := Call(__e, PrimFunc(symtail), W544)


__e.TailApply(tmp2805, tmp2825)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2828 := Call(__e, PrimFunc(symshen_4in_1_6), W542)


__e.TailApply(tmp2804, tmp2828)
return


}, 1)

tmp2829 := Call(__e, PrimFunc(symshen_4_5_1out), W542)


__e.TailApply(tmp2803, tmp2829)
return


}


}, 1)

tmp2832 := Call(__e, PrimFunc(symshen_4_5name_6), V540)


tmp2833 := Call(__e, tmp2802, tmp2832)


__e.TailApply(tmp2778, tmp2833)
return


}, 1)

tmp2834 := Call(__e, ns2_1set, symshen_4_5define_6, tmp2777)


_ = tmp2834

tmp2835 := MakeNative(func(__e *ControlFlow) {
V559 := __e.Get(1)
_ = V559
V560 := __e.Get(2)
_ = V560
tmp2836 := MakeNative(func(__e *ControlFlow) {
W561 := __e.Get(1)
_ = W561
tmp2837 := MakeNative(func(__e *ControlFlow) {
W563 := __e.Get(1)
_ = W563
tmp2838 := MakeNative(func(__e *ControlFlow) {
W564 := __e.Get(1)
_ = W564
tmp2839 := MakeNative(func(__e *ControlFlow) {
W566 := __e.Get(1)
_ = W566
tmp2840 := MakeNative(func(__e *ControlFlow) {
W567 := __e.Get(1)
_ = W567
__e.Return(W567)
return
}, 1)

tmp2841 := Call(__e, PrimFunc(symshen_4compile_1to_1kl), V559, W566, W563)


tmp2842 := Call(__e, PrimFunc(symshen_4factorise_1code), tmp2841)


__e.TailApply(tmp2840, tmp2842)
return


}, 1)

tmp2843 := Call(__e, PrimFunc(symshen_4unprotect), V560)


__e.TailApply(tmp2839, tmp2843)
return


}, 1)

tmp2844 := MakeNative(func(__e *ControlFlow) {
Z565 := __e.Get(1)
_ = Z565
__e.TailApply(PrimFunc(symshen_4free_1var_1chk), V559, Z565)
return
}, 1)

tmp2845 := Call(__e, PrimFunc(symmap), tmp2844, V560)


__e.TailApply(tmp2838, tmp2845)
return


}, 1)

tmp2846 := Call(__e, PrimFunc(symshen_4arity_1chk), V559, W561)


__e.TailApply(tmp2837, tmp2846)
return


}, 1)

tmp2847 := MakeNative(func(__e *ControlFlow) {
Z562 := __e.Get(1)
_ = Z562
__e.TailApply(PrimFunc(symfst), Z562)
return
}, 1)

tmp2848 := Call(__e, PrimFunc(symmap), tmp2847, V560)


__e.TailApply(tmp2836, tmp2848)
return


}, 2)

tmp2849 := Call(__e, ns2_1set, symshen_4shendef_1_6kldef_1h, tmp2835)


_ = tmp2849

tmp2850 := MakeNative(func(__e *ControlFlow) {
V568 := __e.Get(1)
_ = V568
tmp2876 := Call(__e, PrimFunc(symtuple_2), V568)


if True == tmp2876 {
tmp2851 := Call(__e, PrimFunc(symfst), V568)


tmp2852 := Call(__e, PrimFunc(symshen_4unprotect), tmp2851)


tmp2853 := Call(__e, PrimFunc(symsnd), V568)


tmp2854 := Call(__e, PrimFunc(symshen_4unprotect), tmp2853)


__e.TailApply(PrimFunc(sym_8p), tmp2852, tmp2854)
return


} else {
tmp2874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V568)
}
__typedArg0 := V568
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2861 Obj

if True == tmp2874 {
tmp2872 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V568)
}
__typedArg0 := V568
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symprotect, tmp2872)
}
__typedArg0 := symprotect
__typedArg1 := tmp2872
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2863 Obj

if True == tmp2873 {
tmp2870 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V568)
}
__typedArg0 := V568
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp2870)
}
__typedArg0 := tmp2870
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2865 Obj

if True == tmp2871 {
tmp2867 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V568)
}
__typedArg0 := V568
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp2868)
}
__typedArg0 := Nil
__typedArg1 := tmp2868
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres2866 Obj

if True == tmp2869 {
ifres2866 = True


} else {
ifres2866 = False


}

ifres2865 = ifres2866


} else {
ifres2865 = False


}

var ifres2864 Obj

if True == ifres2865 {
ifres2864 = True


} else {
ifres2864 = False


}

ifres2863 = ifres2864


} else {
ifres2863 = False


}

var ifres2862 Obj

if True == ifres2863 {
ifres2862 = True


} else {
ifres2862 = False


}

ifres2861 = ifres2862


} else {
ifres2861 = False


}

if True == ifres2861 {
tmp2855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V568)
}
__typedArg0 := V568
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp2855)
}
__typedArg0 := tmp2855
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4unprotect), tmp2856)
return


} else {
tmp2859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V568)
}
__typedArg0 := V568
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp2859 {
tmp2857 := MakeNative(func(__e *ControlFlow) {
Z569 := __e.Get(1)
_ = Z569
__e.TailApply(PrimFunc(symshen_4unprotect), Z569)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp2857, V568)
return


} else {
__e.Return(V568)
return
}


}


}


}, 1)

tmp2877 := Call(__e, ns2_1set, symshen_4unprotect, tmp2850)


_ = tmp2877

tmp2878 := MakeNative(func(__e *ControlFlow) {
V570 := __e.Get(1)
_ = V570
tmp2879 := MakeNative(func(__e *ControlFlow) {
W571 := __e.Get(1)
_ = W571
tmp2881 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W571)


if True == tmp2881 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W571)
return
}


}, 1)

tmp2897 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V570)
}
__typedArg0 := V570
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2882 Obj

if True == tmp2897 {
tmp2883 := MakeNative(func(__e *ControlFlow) {
W572 := __e.Get(1)
_ = W572
tmp2884 := MakeNative(func(__e *ControlFlow) {
W573 := __e.Get(1)
_ = W573
tmp2892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(W572)
}
__typedArg0 := W572
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

var ifres2888 Obj

if True == tmp2892 {
tmp2890 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(W572)
}
__typedArg0 := W572
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

tmp2891 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp2890)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp2890
return Call(__e, PrimFunc(symnot), __typedArg0)
})()

var ifres2889 Obj

if True == tmp2891 {
ifres2889 = True


} else {
ifres2889 = False


}

ifres2888 = ifres2889


} else {
ifres2888 = False


}

var ifres2885 Obj

if True == ifres2888 {
ifres2885 = W572


} else {
tmp2886 := Call(__e, PrimFunc(symshen_4app), W572, MakeString(" is not a legitimate function name.\n"), symshen_4a)


tmp2887 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp2886)
}
__typedArg0 := tmp2886
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres2885 = tmp2887


}

__e.TailApply(PrimFunc(symshen_4comb), W573, ifres2885)
return


}, 1)

tmp2893 := Call(__e, PrimFunc(symtail), V570)


__e.TailApply(tmp2884, tmp2893)
return


}, 1)

tmp2894 := Call(__e, PrimFunc(symhead), V570)


tmp2895 := Call(__e, tmp2883, tmp2894)


ifres2882 = tmp2895


} else {
tmp2896 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2882 = tmp2896


}

__e.TailApply(tmp2879, ifres2882)
return


}, 1)

tmp2898 := Call(__e, ns2_1set, symshen_4_5name_6, tmp2878)


_ = tmp2898

tmp2899 := MakeNative(func(__e *ControlFlow) {
V574 := __e.Get(1)
_ = V574
tmp2900 := MakeNative(func(__e *ControlFlow) {
W575 := __e.Get(1)
_ = W575
tmp2912 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W575)


if True == tmp2912 {
tmp2901 := MakeNative(func(__e *ControlFlow) {
W581 := __e.Get(1)
_ = W581
tmp2903 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W581)


if True == tmp2903 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W581)
return
}


}, 1)

tmp2904 := MakeNative(func(__e *ControlFlow) {
W582 := __e.Get(1)
_ = W582
tmp2908 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W582)


if True == tmp2908 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2905 := MakeNative(func(__e *ControlFlow) {
W583 := __e.Get(1)
_ = W583
__e.TailApply(PrimFunc(symshen_4comb), W583, Nil)
return
}, 1)

tmp2906 := Call(__e, PrimFunc(symshen_4in_1_6), W582)


__e.TailApply(tmp2905, tmp2906)
return


}


}, 1)

tmp2909 := Call(__e, PrimFunc(sym_5e_6), V574)


tmp2910 := Call(__e, tmp2904, tmp2909)


__e.TailApply(tmp2901, tmp2910)
return


} else {
__e.Return(W575)
return
}


}, 1)

tmp2934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V574)
}
__typedArg0 := V574
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres2913 Obj

if True == tmp2934 {
tmp2914 := MakeNative(func(__e *ControlFlow) {
W576 := __e.Get(1)
_ = W576
tmp2915 := MakeNative(func(__e *ControlFlow) {
W577 := __e.Get(1)
_ = W577
tmp2916 := MakeNative(func(__e *ControlFlow) {
W578 := __e.Get(1)
_ = W578
tmp2928 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W578)


if True == tmp2928 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2917 := MakeNative(func(__e *ControlFlow) {
W579 := __e.Get(1)
_ = W579
tmp2918 := MakeNative(func(__e *ControlFlow) {
W580 := __e.Get(1)
_ = W580
tmp2921 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_j, Nil)
}
__typedArg0 := sym_j
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_i, tmp2921)
}
__typedArg0 := sym_i
__typedArg1 := tmp2921
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2923 := Call(__e, PrimFunc(symelement_2), W576, tmp2922)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp2923)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp2923
return Call(__e, PrimFunc(symnot), __typedArg0)
})() {
tmp2919 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W576, W579)
}
__typedArg0 := W576
__typedArg1 := W579
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W580, tmp2919)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2925 := Call(__e, PrimFunc(symshen_4in_1_6), W578)


__e.TailApply(tmp2918, tmp2925)
return


}, 1)

tmp2926 := Call(__e, PrimFunc(symshen_4_5_1out), W578)


__e.TailApply(tmp2917, tmp2926)
return


}


}, 1)

tmp2929 := Call(__e, PrimFunc(symshen_4_5signature_6), W577)


__e.TailApply(tmp2916, tmp2929)
return


}, 1)

tmp2930 := Call(__e, PrimFunc(symtail), V574)


__e.TailApply(tmp2915, tmp2930)
return


}, 1)

tmp2931 := Call(__e, PrimFunc(symhead), V574)


tmp2932 := Call(__e, tmp2914, tmp2931)


ifres2913 = tmp2932


} else {
tmp2933 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2913 = tmp2933


}

__e.TailApply(tmp2900, ifres2913)
return


}, 1)

tmp2935 := Call(__e, ns2_1set, symshen_4_5signature_6, tmp2899)


_ = tmp2935

tmp2936 := MakeNative(func(__e *ControlFlow) {
V584 := __e.Get(1)
_ = V584
tmp2937 := MakeNative(func(__e *ControlFlow) {
W585 := __e.Get(1)
_ = W585
tmp2956 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W585)


if True == tmp2956 {
tmp2938 := MakeNative(func(__e *ControlFlow) {
W592 := __e.Get(1)
_ = W592
tmp2940 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W592)


if True == tmp2940 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W592)
return
}


}, 1)

tmp2941 := MakeNative(func(__e *ControlFlow) {
W593 := __e.Get(1)
_ = W593
tmp2952 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W593)


if True == tmp2952 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2942 := MakeNative(func(__e *ControlFlow) {
W594 := __e.Get(1)
_ = W594
tmp2943 := MakeNative(func(__e *ControlFlow) {
W595 := __e.Get(1)
_ = W595
tmp2948 := Call(__e, PrimFunc(symempty_2), W594)


var ifres2944 Obj

if True == tmp2948 {
ifres2944 = Nil


} else {
tmp2945 := Call(__e, PrimFunc(symshen_4app), W594, MakeString("\n ..."), symshen_4r)


tmp2947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Shen syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp2945)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Shen syntax error here:\n ")
__typedArg1 := tmp2945
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("Shen syntax error here:\n "))
__typedS1, __typedOK1 := TypedString(tmp2945)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("Shen syntax error here:\n ")
__typedArg1 := tmp2945
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})()

ifres2944 = tmp2947


}

__e.TailApply(PrimFunc(symshen_4comb), W595, ifres2944)
return


}, 1)

tmp2949 := Call(__e, PrimFunc(symshen_4in_1_6), W593)


__e.TailApply(tmp2943, tmp2949)
return


}, 1)

tmp2950 := Call(__e, PrimFunc(symshen_4_5_1out), W593)


__e.TailApply(tmp2942, tmp2950)
return


}


}, 1)

tmp2953 := Call(__e, PrimFunc(sym_5_b_6), V584)


tmp2954 := Call(__e, tmp2941, tmp2953)


__e.TailApply(tmp2938, tmp2954)
return


} else {
__e.Return(W585)
return
}


}, 1)

tmp2957 := MakeNative(func(__e *ControlFlow) {
W586 := __e.Get(1)
_ = W586
tmp2973 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W586)


if True == tmp2973 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2958 := MakeNative(func(__e *ControlFlow) {
W587 := __e.Get(1)
_ = W587
tmp2959 := MakeNative(func(__e *ControlFlow) {
W588 := __e.Get(1)
_ = W588
tmp2960 := MakeNative(func(__e *ControlFlow) {
W589 := __e.Get(1)
_ = W589
tmp2968 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W589)


if True == tmp2968 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2961 := MakeNative(func(__e *ControlFlow) {
W590 := __e.Get(1)
_ = W590
tmp2962 := MakeNative(func(__e *ControlFlow) {
W591 := __e.Get(1)
_ = W591
tmp2963 := Call(__e, PrimFunc(symshen_4linearise), W587)


tmp2964 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2963, W590)
}
__typedArg0 := tmp2963
__typedArg1 := W590
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W591, tmp2964)
return


}, 1)

tmp2965 := Call(__e, PrimFunc(symshen_4in_1_6), W589)


__e.TailApply(tmp2962, tmp2965)
return


}, 1)

tmp2966 := Call(__e, PrimFunc(symshen_4_5_1out), W589)


__e.TailApply(tmp2961, tmp2966)
return


}


}, 1)

tmp2969 := Call(__e, PrimFunc(symshen_4_5rules_6), W588)


__e.TailApply(tmp2960, tmp2969)
return


}, 1)

tmp2970 := Call(__e, PrimFunc(symshen_4in_1_6), W586)


__e.TailApply(tmp2959, tmp2970)
return


}, 1)

tmp2971 := Call(__e, PrimFunc(symshen_4_5_1out), W586)


__e.TailApply(tmp2958, tmp2971)
return


}


}, 1)

tmp2974 := Call(__e, PrimFunc(symshen_4_5rule_6), V584)


tmp2975 := Call(__e, tmp2957, tmp2974)


__e.TailApply(tmp2937, tmp2975)
return


}, 1)

tmp2976 := Call(__e, ns2_1set, symshen_4_5rules_6, tmp2936)


_ = tmp2976

tmp2977 := MakeNative(func(__e *ControlFlow) {
V598 := __e.Get(1)
_ = V598
tmp2982 := Call(__e, PrimFunc(symtuple_2), V598)


if True == tmp2982 {
tmp2978 := Call(__e, PrimFunc(symfst), V598)


tmp2979 := Call(__e, PrimFunc(symfst), V598)


tmp2980 := Call(__e, PrimFunc(symsnd), V598)


__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp2978, tmp2979, Nil, tmp2980)
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

tmp2983 := Call(__e, ns2_1set, symshen_4linearise, tmp2977)


_ = tmp2983

tmp2984 := MakeNative(func(__e *ControlFlow) {
V611 := __e.Get(1)
_ = V611
V612 := __e.Get(2)
_ = V612
V613 := __e.Get(3)
_ = V613
V614 := __e.Get(4)
_ = V614
tmp3022 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V611)
}
__typedArg0 := Nil
__typedArg1 := V611
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp3022 {
__e.TailApply(PrimFunc(sym_8p), V612, V614)
return
} else {
tmp3020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3016 Obj

if True == tmp3020 {
tmp3018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3018)
}
__typedArg0 := tmp3018
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3017 Obj

if True == tmp3019 {
ifres3017 = True


} else {
ifres3017 = False


}

ifres3016 = ifres3017


} else {
ifres3016 = False


}

if True == ifres3016 {
tmp2985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2987 := Call(__e, PrimFunc(symappend), tmp2985, tmp2986)


__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp2987, V612, V613, V614)
return


} else {
tmp3014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3010 Obj

if True == tmp3014 {
tmp3012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3013 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp3012)
}
__typedArg0 := tmp3012
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

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
tmp3004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3005 := Call(__e, PrimFunc(symelement_2), tmp3004, V613)


if True == tmp3005 {
tmp2988 := MakeNative(func(__e *ControlFlow) {
W615 := __e.Get(1)
_ = W615
tmp2989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp2990 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2991 := Call(__e, PrimFunc(symshen_4rep_1X), tmp2990, W615, V612)


tmp2992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp2993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2992, Nil)
}
__typedArg0 := tmp2992
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W615, tmp2993)
}
__typedArg0 := W615
__typedArg1 := tmp2993
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp2994)
}
__typedArg0 := sym_a
__typedArg1 := tmp2994
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V614, Nil)
}
__typedArg0 := V614
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp2995, tmp2996)
}
__typedArg0 := tmp2995
__typedArg1 := tmp2996
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp2998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp2997)
}
__typedArg0 := symwhere
__typedArg1 := tmp2997
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp2989, tmp2991, V613, tmp2998)
return


}, 1)

tmp2999 := Call(__e, PrimFunc(symgensym), symV)


__e.TailApply(tmp2988, tmp2999)
return


} else {
tmp3000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3001, V613)
}
__typedArg0 := tmp3001
__typedArg1 := V613
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp3000, V612, tmp3002, V614)
return


}


} else {
tmp3008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3008 {
tmp3006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V611)
}
__typedArg0 := V611
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp3006, V612, V613, V614)
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

tmp3023 := Call(__e, ns2_1set, symshen_4linearise_1h, tmp2984)


_ = tmp3023

tmp3024 := MakeNative(func(__e *ControlFlow) {
V616 := __e.Get(1)
_ = V616
tmp3025 := MakeNative(func(__e *ControlFlow) {
W617 := __e.Get(1)
_ = W617
tmp3113 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W617)


if True == tmp3113 {
tmp3026 := MakeNative(func(__e *ControlFlow) {
W627 := __e.Get(1)
_ = W627
tmp3091 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W627)


if True == tmp3091 {
tmp3027 := MakeNative(func(__e *ControlFlow) {
W634 := __e.Get(1)
_ = W634
tmp3054 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W634)


if True == tmp3054 {
tmp3028 := MakeNative(func(__e *ControlFlow) {
W644 := __e.Get(1)
_ = W644
tmp3030 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W644)


if True == tmp3030 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W644)
return
}


}, 1)

tmp3031 := MakeNative(func(__e *ControlFlow) {
W645 := __e.Get(1)
_ = W645
tmp3050 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W645)


if True == tmp3050 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3032 := MakeNative(func(__e *ControlFlow) {
W646 := __e.Get(1)
_ = W646
tmp3033 := MakeNative(func(__e *ControlFlow) {
W647 := __e.Get(1)
_ = W647
tmp3046 := Call(__e, PrimFunc(symshen_4hds_a_2), W647, sym_5_1)


if True == tmp3046 {
tmp3034 := MakeNative(func(__e *ControlFlow) {
W648 := __e.Get(1)
_ = W648
tmp3043 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W648)
}
__typedArg0 := W648
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3043 {
tmp3035 := MakeNative(func(__e *ControlFlow) {
W649 := __e.Get(1)
_ = W649
tmp3036 := MakeNative(func(__e *ControlFlow) {
W650 := __e.Get(1)
_ = W650
tmp3037 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W649, Nil)
}
__typedArg0 := W649
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3038 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4choicepoint_b, tmp3037)
}
__typedArg0 := symshen_4choicepoint_b
__typedArg1 := tmp3037
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3039 := Call(__e, PrimFunc(sym_8p), W646, tmp3038)


__e.TailApply(PrimFunc(symshen_4comb), W650, tmp3039)
return


}, 1)

tmp3040 := Call(__e, PrimFunc(symtail), W648)


__e.TailApply(tmp3036, tmp3040)
return


}, 1)

tmp3041 := Call(__e, PrimFunc(symhead), W648)


__e.TailApply(tmp3035, tmp3041)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3044 := Call(__e, PrimFunc(symtail), W647)


__e.TailApply(tmp3034, tmp3044)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3047 := Call(__e, PrimFunc(symshen_4in_1_6), W645)


__e.TailApply(tmp3033, tmp3047)
return


}, 1)

tmp3048 := Call(__e, PrimFunc(symshen_4_5_1out), W645)


__e.TailApply(tmp3032, tmp3048)
return


}


}, 1)

tmp3051 := Call(__e, PrimFunc(symshen_4_5patterns_6), V616)


tmp3052 := Call(__e, tmp3031, tmp3051)


__e.TailApply(tmp3028, tmp3052)
return


} else {
__e.Return(W634)
return
}


}, 1)

tmp3055 := MakeNative(func(__e *ControlFlow) {
W635 := __e.Get(1)
_ = W635
tmp3087 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W635)


if True == tmp3087 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3056 := MakeNative(func(__e *ControlFlow) {
W636 := __e.Get(1)
_ = W636
tmp3057 := MakeNative(func(__e *ControlFlow) {
W637 := __e.Get(1)
_ = W637
tmp3083 := Call(__e, PrimFunc(symshen_4hds_a_2), W637, sym_5_1)


if True == tmp3083 {
tmp3058 := MakeNative(func(__e *ControlFlow) {
W638 := __e.Get(1)
_ = W638
tmp3080 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W638)
}
__typedArg0 := W638
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3080 {
tmp3059 := MakeNative(func(__e *ControlFlow) {
W639 := __e.Get(1)
_ = W639
tmp3060 := MakeNative(func(__e *ControlFlow) {
W640 := __e.Get(1)
_ = W640
tmp3076 := Call(__e, PrimFunc(symshen_4hds_a_2), W640, symwhere)


if True == tmp3076 {
tmp3061 := MakeNative(func(__e *ControlFlow) {
W641 := __e.Get(1)
_ = W641
tmp3073 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W641)
}
__typedArg0 := W641
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3073 {
tmp3062 := MakeNative(func(__e *ControlFlow) {
W642 := __e.Get(1)
_ = W642
tmp3063 := MakeNative(func(__e *ControlFlow) {
W643 := __e.Get(1)
_ = W643
tmp3064 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W639, Nil)
}
__typedArg0 := W639
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3065 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4choicepoint_b, tmp3064)
}
__typedArg0 := symshen_4choicepoint_b
__typedArg1 := tmp3064
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3066 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3065, Nil)
}
__typedArg0 := tmp3065
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3067 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W642, tmp3066)
}
__typedArg0 := W642
__typedArg1 := tmp3066
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symwhere, tmp3067)
}
__typedArg0 := symwhere
__typedArg1 := tmp3067
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3069 := Call(__e, PrimFunc(sym_8p), W636, tmp3068)


__e.TailApply(PrimFunc(symshen_4comb), W643, tmp3069)
return


}, 1)

tmp3070 := Call(__e, PrimFunc(symtail), W641)


__e.TailApply(tmp3063, tmp3070)
return


}, 1)

tmp3071 := Call(__e, PrimFunc(symhead), W641)


__e.TailApply(tmp3062, tmp3071)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3074 := Call(__e, PrimFunc(symtail), W640)


__e.TailApply(tmp3061, tmp3074)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3077 := Call(__e, PrimFunc(symtail), W638)


__e.TailApply(tmp3060, tmp3077)
return


}, 1)

tmp3078 := Call(__e, PrimFunc(symhead), W638)


__e.TailApply(tmp3059, tmp3078)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3081 := Call(__e, PrimFunc(symtail), W637)


__e.TailApply(tmp3058, tmp3081)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3084 := Call(__e, PrimFunc(symshen_4in_1_6), W635)


__e.TailApply(tmp3057, tmp3084)
return


}, 1)

tmp3085 := Call(__e, PrimFunc(symshen_4_5_1out), W635)


__e.TailApply(tmp3056, tmp3085)
return


}


}, 1)

tmp3088 := Call(__e, PrimFunc(symshen_4_5patterns_6), V616)


tmp3089 := Call(__e, tmp3055, tmp3088)


__e.TailApply(tmp3027, tmp3089)
return


} else {
__e.Return(W627)
return
}


}, 1)

tmp3092 := MakeNative(func(__e *ControlFlow) {
W628 := __e.Get(1)
_ = W628
tmp3109 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W628)


if True == tmp3109 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3093 := MakeNative(func(__e *ControlFlow) {
W629 := __e.Get(1)
_ = W629
tmp3094 := MakeNative(func(__e *ControlFlow) {
W630 := __e.Get(1)
_ = W630
tmp3105 := Call(__e, PrimFunc(symshen_4hds_a_2), W630, sym_1_6)


if True == tmp3105 {
tmp3095 := MakeNative(func(__e *ControlFlow) {
W631 := __e.Get(1)
_ = W631
tmp3102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W631)
}
__typedArg0 := W631
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3102 {
tmp3096 := MakeNative(func(__e *ControlFlow) {
W632 := __e.Get(1)
_ = W632
tmp3097 := MakeNative(func(__e *ControlFlow) {
W633 := __e.Get(1)
_ = W633
tmp3098 := Call(__e, PrimFunc(sym_8p), W629, W632)


__e.TailApply(PrimFunc(symshen_4comb), W633, tmp3098)
return


}, 1)

tmp3099 := Call(__e, PrimFunc(symtail), W631)


__e.TailApply(tmp3097, tmp3099)
return


}, 1)

tmp3100 := Call(__e, PrimFunc(symhead), W631)


__e.TailApply(tmp3096, tmp3100)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3103 := Call(__e, PrimFunc(symtail), W630)


__e.TailApply(tmp3095, tmp3103)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3106 := Call(__e, PrimFunc(symshen_4in_1_6), W628)


__e.TailApply(tmp3094, tmp3106)
return


}, 1)

tmp3107 := Call(__e, PrimFunc(symshen_4_5_1out), W628)


__e.TailApply(tmp3093, tmp3107)
return


}


}, 1)

tmp3110 := Call(__e, PrimFunc(symshen_4_5patterns_6), V616)


tmp3111 := Call(__e, tmp3092, tmp3110)


__e.TailApply(tmp3026, tmp3111)
return


} else {
__e.Return(W617)
return
}


}, 1)

tmp3114 := MakeNative(func(__e *ControlFlow) {
W618 := __e.Get(1)
_ = W618
tmp3144 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W618)


if True == tmp3144 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3115 := MakeNative(func(__e *ControlFlow) {
W619 := __e.Get(1)
_ = W619
tmp3116 := MakeNative(func(__e *ControlFlow) {
W620 := __e.Get(1)
_ = W620
tmp3140 := Call(__e, PrimFunc(symshen_4hds_a_2), W620, sym_1_6)


if True == tmp3140 {
tmp3117 := MakeNative(func(__e *ControlFlow) {
W621 := __e.Get(1)
_ = W621
tmp3137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W621)
}
__typedArg0 := W621
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3137 {
tmp3118 := MakeNative(func(__e *ControlFlow) {
W622 := __e.Get(1)
_ = W622
tmp3119 := MakeNative(func(__e *ControlFlow) {
W623 := __e.Get(1)
_ = W623
tmp3133 := Call(__e, PrimFunc(symshen_4hds_a_2), W623, symwhere)


if True == tmp3133 {
tmp3120 := MakeNative(func(__e *ControlFlow) {
W624 := __e.Get(1)
_ = W624
tmp3130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W624)
}
__typedArg0 := W624
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3130 {
tmp3121 := MakeNative(func(__e *ControlFlow) {
W625 := __e.Get(1)
_ = W625
tmp3122 := MakeNative(func(__e *ControlFlow) {
W626 := __e.Get(1)
_ = W626
tmp3123 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W622, Nil)
}
__typedArg0 := W622
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W625, tmp3123)
}
__typedArg0 := W625
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

tmp3126 := Call(__e, PrimFunc(sym_8p), W619, tmp3125)


__e.TailApply(PrimFunc(symshen_4comb), W626, tmp3126)
return


}, 1)

tmp3127 := Call(__e, PrimFunc(symtail), W624)


__e.TailApply(tmp3122, tmp3127)
return


}, 1)

tmp3128 := Call(__e, PrimFunc(symhead), W624)


__e.TailApply(tmp3121, tmp3128)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3131 := Call(__e, PrimFunc(symtail), W623)


__e.TailApply(tmp3120, tmp3131)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3134 := Call(__e, PrimFunc(symtail), W621)


__e.TailApply(tmp3119, tmp3134)
return


}, 1)

tmp3135 := Call(__e, PrimFunc(symhead), W621)


__e.TailApply(tmp3118, tmp3135)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3138 := Call(__e, PrimFunc(symtail), W620)


__e.TailApply(tmp3117, tmp3138)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3141 := Call(__e, PrimFunc(symshen_4in_1_6), W618)


__e.TailApply(tmp3116, tmp3141)
return


}, 1)

tmp3142 := Call(__e, PrimFunc(symshen_4_5_1out), W618)


__e.TailApply(tmp3115, tmp3142)
return


}


}, 1)

tmp3145 := Call(__e, PrimFunc(symshen_4_5patterns_6), V616)


tmp3146 := Call(__e, tmp3114, tmp3145)


__e.TailApply(tmp3025, tmp3146)
return


}, 1)

tmp3147 := Call(__e, ns2_1set, symshen_4_5rule_6, tmp3024)


_ = tmp3147

tmp3148 := MakeNative(func(__e *ControlFlow) {
V651 := __e.Get(1)
_ = V651
tmp3149 := MakeNative(func(__e *ControlFlow) {
W652 := __e.Get(1)
_ = W652
tmp3161 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W652)


if True == tmp3161 {
tmp3150 := MakeNative(func(__e *ControlFlow) {
W659 := __e.Get(1)
_ = W659
tmp3152 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W659)


if True == tmp3152 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W659)
return
}


}, 1)

tmp3153 := MakeNative(func(__e *ControlFlow) {
W660 := __e.Get(1)
_ = W660
tmp3157 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W660)


if True == tmp3157 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3154 := MakeNative(func(__e *ControlFlow) {
W661 := __e.Get(1)
_ = W661
__e.TailApply(PrimFunc(symshen_4comb), W661, Nil)
return
}, 1)

tmp3155 := Call(__e, PrimFunc(symshen_4in_1_6), W660)


__e.TailApply(tmp3154, tmp3155)
return


}


}, 1)

tmp3158 := Call(__e, PrimFunc(sym_5e_6), V651)


tmp3159 := Call(__e, tmp3153, tmp3158)


__e.TailApply(tmp3150, tmp3159)
return


} else {
__e.Return(W652)
return
}


}, 1)

tmp3162 := MakeNative(func(__e *ControlFlow) {
W653 := __e.Get(1)
_ = W653
tmp3177 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W653)


if True == tmp3177 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3163 := MakeNative(func(__e *ControlFlow) {
W654 := __e.Get(1)
_ = W654
tmp3164 := MakeNative(func(__e *ControlFlow) {
W655 := __e.Get(1)
_ = W655
tmp3165 := MakeNative(func(__e *ControlFlow) {
W656 := __e.Get(1)
_ = W656
tmp3172 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W656)


if True == tmp3172 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3166 := MakeNative(func(__e *ControlFlow) {
W657 := __e.Get(1)
_ = W657
tmp3167 := MakeNative(func(__e *ControlFlow) {
W658 := __e.Get(1)
_ = W658
tmp3168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W654, W657)
}
__typedArg0 := W654
__typedArg1 := W657
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W658, tmp3168)
return


}, 1)

tmp3169 := Call(__e, PrimFunc(symshen_4in_1_6), W656)


__e.TailApply(tmp3167, tmp3169)
return


}, 1)

tmp3170 := Call(__e, PrimFunc(symshen_4_5_1out), W656)


__e.TailApply(tmp3166, tmp3170)
return


}


}, 1)

tmp3173 := Call(__e, PrimFunc(symshen_4_5patterns_6), W655)


__e.TailApply(tmp3165, tmp3173)
return


}, 1)

tmp3174 := Call(__e, PrimFunc(symshen_4in_1_6), W653)


__e.TailApply(tmp3164, tmp3174)
return


}, 1)

tmp3175 := Call(__e, PrimFunc(symshen_4_5_1out), W653)


__e.TailApply(tmp3163, tmp3175)
return


}


}, 1)

tmp3178 := Call(__e, PrimFunc(symshen_4_5pattern_6), V651)


tmp3179 := Call(__e, tmp3162, tmp3178)


__e.TailApply(tmp3149, tmp3179)
return


}, 1)

tmp3180 := Call(__e, ns2_1set, symshen_4_5patterns_6, tmp3148)


_ = tmp3180

tmp3181 := MakeNative(func(__e *ControlFlow) {
V662 := __e.Get(1)
_ = V662
tmp3182 := MakeNative(func(__e *ControlFlow) {
W663 := __e.Get(1)
_ = W663
tmp3237 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W663)


if True == tmp3237 {
tmp3183 := MakeNative(func(__e *ControlFlow) {
W677 := __e.Get(1)
_ = W677
tmp3211 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W677)


if True == tmp3211 {
tmp3184 := MakeNative(func(__e *ControlFlow) {
W684 := __e.Get(1)
_ = W684
tmp3198 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W684)


if True == tmp3198 {
tmp3185 := MakeNative(func(__e *ControlFlow) {
W687 := __e.Get(1)
_ = W687
tmp3187 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W687)


if True == tmp3187 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W687)
return
}


}, 1)

tmp3188 := MakeNative(func(__e *ControlFlow) {
W688 := __e.Get(1)
_ = W688
tmp3194 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W688)


if True == tmp3194 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3189 := MakeNative(func(__e *ControlFlow) {
W689 := __e.Get(1)
_ = W689
tmp3190 := MakeNative(func(__e *ControlFlow) {
W690 := __e.Get(1)
_ = W690
__e.TailApply(PrimFunc(symshen_4comb), W690, W689)
return
}, 1)

tmp3191 := Call(__e, PrimFunc(symshen_4in_1_6), W688)


__e.TailApply(tmp3190, tmp3191)
return


}, 1)

tmp3192 := Call(__e, PrimFunc(symshen_4_5_1out), W688)


__e.TailApply(tmp3189, tmp3192)
return


}


}, 1)

tmp3195 := Call(__e, PrimFunc(symshen_4_5simple_1pattern_6), V662)


tmp3196 := Call(__e, tmp3188, tmp3195)


__e.TailApply(tmp3185, tmp3196)
return


} else {
__e.Return(W684)
return
}


}, 1)

tmp3209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V662)
}
__typedArg0 := V662
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3199 Obj

if True == tmp3209 {
tmp3200 := MakeNative(func(__e *ControlFlow) {
W685 := __e.Get(1)
_ = W685
tmp3201 := MakeNative(func(__e *ControlFlow) {
W686 := __e.Get(1)
_ = W686
tmp3204 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(W685)
}
__typedArg0 := W685
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3204 {
tmp3202 := Call(__e, PrimFunc(symshen_4constructor_1error), W685)


__e.TailApply(PrimFunc(symshen_4comb), W686, tmp3202)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3205 := Call(__e, PrimFunc(symtail), V662)


__e.TailApply(tmp3201, tmp3205)
return


}, 1)

tmp3206 := Call(__e, PrimFunc(symhead), V662)


tmp3207 := Call(__e, tmp3200, tmp3206)


ifres3199 = tmp3207


} else {
tmp3208 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3199 = tmp3208


}

__e.TailApply(tmp3184, ifres3199)
return


} else {
__e.Return(W677)
return
}


}, 1)

tmp3235 := Call(__e, PrimFunc(symshen_4ccons_2), V662)


var ifres3212 Obj

if True == tmp3235 {
tmp3213 := MakeNative(func(__e *ControlFlow) {
W678 := __e.Get(1)
_ = W678
tmp3214 := MakeNative(func(__e *ControlFlow) {
W679 := __e.Get(1)
_ = W679
tmp3230 := Call(__e, PrimFunc(symshen_4hds_a_2), W678, symvector)


if True == tmp3230 {
tmp3215 := MakeNative(func(__e *ControlFlow) {
W680 := __e.Get(1)
_ = W680
tmp3227 := Call(__e, PrimFunc(symshen_4hds_a_2), W680, MakeNumber(0))


if True == tmp3227 {
tmp3216 := MakeNative(func(__e *ControlFlow) {
W681 := __e.Get(1)
_ = W681
tmp3217 := MakeNative(func(__e *ControlFlow) {
W682 := __e.Get(1)
_ = W682
tmp3223 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W682)


if True == tmp3223 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3218 := MakeNative(func(__e *ControlFlow) {
W683 := __e.Get(1)
_ = W683
tmp3219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(MakeNumber(0), Nil)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3220 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symvector, tmp3219)
}
__typedArg0 := symvector
__typedArg1 := tmp3219
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W679, tmp3220)
return


}, 1)

tmp3221 := Call(__e, PrimFunc(symshen_4in_1_6), W682)


__e.TailApply(tmp3218, tmp3221)
return


}


}, 1)

tmp3224 := Call(__e, PrimFunc(sym_5end_6), W681)


__e.TailApply(tmp3217, tmp3224)
return


}, 1)

tmp3225 := Call(__e, PrimFunc(symtail), W680)


__e.TailApply(tmp3216, tmp3225)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3228 := Call(__e, PrimFunc(symtail), W678)


__e.TailApply(tmp3215, tmp3228)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3231 := Call(__e, PrimFunc(symtail), V662)


__e.TailApply(tmp3214, tmp3231)
return


}, 1)

tmp3232 := Call(__e, PrimFunc(symhead), V662)


tmp3233 := Call(__e, tmp3213, tmp3232)


ifres3212 = tmp3233


} else {
tmp3234 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3212 = tmp3234


}

__e.TailApply(tmp3183, ifres3212)
return


} else {
__e.Return(W663)
return
}


}, 1)

tmp3278 := Call(__e, PrimFunc(symshen_4ccons_2), V662)


var ifres3238 Obj

if True == tmp3278 {
tmp3239 := MakeNative(func(__e *ControlFlow) {
W664 := __e.Get(1)
_ = W664
tmp3240 := MakeNative(func(__e *ControlFlow) {
W665 := __e.Get(1)
_ = W665
tmp3241 := MakeNative(func(__e *ControlFlow) {
W666 := __e.Get(1)
_ = W666
tmp3272 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W666)


if True == tmp3272 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3242 := MakeNative(func(__e *ControlFlow) {
W667 := __e.Get(1)
_ = W667
tmp3243 := MakeNative(func(__e *ControlFlow) {
W668 := __e.Get(1)
_ = W668
tmp3244 := MakeNative(func(__e *ControlFlow) {
W669 := __e.Get(1)
_ = W669
tmp3267 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W669)


if True == tmp3267 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3245 := MakeNative(func(__e *ControlFlow) {
W670 := __e.Get(1)
_ = W670
tmp3246 := MakeNative(func(__e *ControlFlow) {
W671 := __e.Get(1)
_ = W671
tmp3247 := MakeNative(func(__e *ControlFlow) {
W672 := __e.Get(1)
_ = W672
tmp3262 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W672)


if True == tmp3262 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3248 := MakeNative(func(__e *ControlFlow) {
W673 := __e.Get(1)
_ = W673
tmp3249 := MakeNative(func(__e *ControlFlow) {
W674 := __e.Get(1)
_ = W674
tmp3250 := MakeNative(func(__e *ControlFlow) {
W675 := __e.Get(1)
_ = W675
tmp3257 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W675)


if True == tmp3257 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3251 := MakeNative(func(__e *ControlFlow) {
W676 := __e.Get(1)
_ = W676
tmp3252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W673, Nil)
}
__typedArg0 := W673
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3253 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W670, tmp3252)
}
__typedArg0 := W670
__typedArg1 := tmp3252
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3254 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W667, tmp3253)
}
__typedArg0 := W667
__typedArg1 := tmp3253
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4comb), W665, tmp3254)
return


}, 1)

tmp3255 := Call(__e, PrimFunc(symshen_4in_1_6), W675)


__e.TailApply(tmp3251, tmp3255)
return


}


}, 1)

tmp3258 := Call(__e, PrimFunc(sym_5end_6), W674)


__e.TailApply(tmp3250, tmp3258)
return


}, 1)

tmp3259 := Call(__e, PrimFunc(symshen_4in_1_6), W672)


__e.TailApply(tmp3249, tmp3259)
return


}, 1)

tmp3260 := Call(__e, PrimFunc(symshen_4_5_1out), W672)


__e.TailApply(tmp3248, tmp3260)
return


}


}, 1)

tmp3263 := Call(__e, PrimFunc(symshen_4_5pattern2_6), W671)


__e.TailApply(tmp3247, tmp3263)
return


}, 1)

tmp3264 := Call(__e, PrimFunc(symshen_4in_1_6), W669)


__e.TailApply(tmp3246, tmp3264)
return


}, 1)

tmp3265 := Call(__e, PrimFunc(symshen_4_5_1out), W669)


__e.TailApply(tmp3245, tmp3265)
return


}


}, 1)

tmp3268 := Call(__e, PrimFunc(symshen_4_5pattern1_6), W668)


__e.TailApply(tmp3244, tmp3268)
return


}, 1)

tmp3269 := Call(__e, PrimFunc(symshen_4in_1_6), W666)


__e.TailApply(tmp3243, tmp3269)
return


}, 1)

tmp3270 := Call(__e, PrimFunc(symshen_4_5_1out), W666)


__e.TailApply(tmp3242, tmp3270)
return


}


}, 1)

tmp3273 := Call(__e, PrimFunc(symshen_4_5constructor_6), W664)


__e.TailApply(tmp3241, tmp3273)
return


}, 1)

tmp3274 := Call(__e, PrimFunc(symtail), V662)


__e.TailApply(tmp3240, tmp3274)
return


}, 1)

tmp3275 := Call(__e, PrimFunc(symhead), V662)


tmp3276 := Call(__e, tmp3239, tmp3275)


ifres3238 = tmp3276


} else {
tmp3277 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3238 = tmp3277


}

__e.TailApply(tmp3182, ifres3238)
return


}, 1)

tmp3279 := Call(__e, ns2_1set, symshen_4_5pattern_6, tmp3181)


_ = tmp3279

tmp3280 := MakeNative(func(__e *ControlFlow) {
V691 := __e.Get(1)
_ = V691
tmp3281 := MakeNative(func(__e *ControlFlow) {
W692 := __e.Get(1)
_ = W692
tmp3283 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W692)


if True == tmp3283 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W692)
return
}


}, 1)

tmp3293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V691)
}
__typedArg0 := V691
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3284 Obj

if True == tmp3293 {
tmp3285 := MakeNative(func(__e *ControlFlow) {
W693 := __e.Get(1)
_ = W693
tmp3286 := MakeNative(func(__e *ControlFlow) {
W694 := __e.Get(1)
_ = W694
tmp3288 := Call(__e, PrimFunc(symshen_4constructor_2), W693)


if True == tmp3288 {
__e.TailApply(PrimFunc(symshen_4comb), W694, W693)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3289 := Call(__e, PrimFunc(symtail), V691)


__e.TailApply(tmp3286, tmp3289)
return


}, 1)

tmp3290 := Call(__e, PrimFunc(symhead), V691)


tmp3291 := Call(__e, tmp3285, tmp3290)


ifres3284 = tmp3291


} else {
tmp3292 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3284 = tmp3292


}

__e.TailApply(tmp3281, ifres3284)
return


}, 1)

tmp3294 := Call(__e, ns2_1set, symshen_4_5constructor_6, tmp3280)


_ = tmp3294

tmp3295 := MakeNative(func(__e *ControlFlow) {
V695 := __e.Get(1)
_ = V695
tmp3296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8v, Nil)
}
__typedArg0 := sym_8v
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8s, tmp3296)
}
__typedArg0 := sym_8s
__typedArg1 := tmp3296
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_8p, tmp3297)
}
__typedArg0 := sym_8p
__typedArg1 := tmp3297
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcons, tmp3298)
}
__typedArg0 := symcons
__typedArg1 := tmp3298
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symelement_2), V695, tmp3299)
return


}, 1)

tmp3300 := Call(__e, ns2_1set, symshen_4constructor_2, tmp3295)


_ = tmp3300

tmp3301 := MakeNative(func(__e *ControlFlow) {
V696 := __e.Get(1)
_ = V696
tmp3302 := Call(__e, PrimFunc(symshen_4app), V696, MakeString(" is not a legitimate constructor\n"), symshen_4r)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(tmp3302)
}
__typedArg0 := tmp3302
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}, 1)

tmp3303 := Call(__e, ns2_1set, symshen_4constructor_1error, tmp3301)


_ = tmp3303

tmp3304 := MakeNative(func(__e *ControlFlow) {
V697 := __e.Get(1)
_ = V697
tmp3305 := MakeNative(func(__e *ControlFlow) {
W698 := __e.Get(1)
_ = W698
tmp3323 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W698)


if True == tmp3323 {
tmp3306 := MakeNative(func(__e *ControlFlow) {
W701 := __e.Get(1)
_ = W701
tmp3308 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W701)


if True == tmp3308 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W701)
return
}


}, 1)

tmp3321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V697)
}
__typedArg0 := V697
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3309 Obj

if True == tmp3321 {
tmp3310 := MakeNative(func(__e *ControlFlow) {
W702 := __e.Get(1)
_ = W702
tmp3311 := MakeNative(func(__e *ControlFlow) {
W703 := __e.Get(1)
_ = W703
tmp3313 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_5_1, Nil)
}
__typedArg0 := sym_5_1
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_1_6, tmp3313)
}
__typedArg0 := sym_1_6
__typedArg1 := tmp3313
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3315 := Call(__e, PrimFunc(symelement_2), W702, tmp3314)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp3315)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp3315
return Call(__e, PrimFunc(symnot), __typedArg0)
})() {
__e.TailApply(PrimFunc(symshen_4comb), W703, W702)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3317 := Call(__e, PrimFunc(symtail), V697)


__e.TailApply(tmp3311, tmp3317)
return


}, 1)

tmp3318 := Call(__e, PrimFunc(symhead), V697)


tmp3319 := Call(__e, tmp3310, tmp3318)


ifres3309 = tmp3319


} else {
tmp3320 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3309 = tmp3320


}

__e.TailApply(tmp3306, ifres3309)
return


} else {
__e.Return(W698)
return
}


}, 1)

tmp3334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V697)
}
__typedArg0 := V697
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3324 Obj

if True == tmp3334 {
tmp3325 := MakeNative(func(__e *ControlFlow) {
W699 := __e.Get(1)
_ = W699
tmp3326 := MakeNative(func(__e *ControlFlow) {
W700 := __e.Get(1)
_ = W700
tmp3329 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W699, sym__)
}
__typedArg0 := W699
__typedArg1 := sym__
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp3329 {
tmp3327 := Call(__e, PrimFunc(symgensym), symY)


__e.TailApply(PrimFunc(symshen_4comb), W700, tmp3327)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3330 := Call(__e, PrimFunc(symtail), V697)


__e.TailApply(tmp3326, tmp3330)
return


}, 1)

tmp3331 := Call(__e, PrimFunc(symhead), V697)


tmp3332 := Call(__e, tmp3325, tmp3331)


ifres3324 = tmp3332


} else {
tmp3333 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3324 = tmp3333


}

__e.TailApply(tmp3305, ifres3324)
return


}, 1)

tmp3335 := Call(__e, ns2_1set, symshen_4_5simple_1pattern_6, tmp3304)


_ = tmp3335

tmp3336 := MakeNative(func(__e *ControlFlow) {
V704 := __e.Get(1)
_ = V704
tmp3337 := MakeNative(func(__e *ControlFlow) {
W705 := __e.Get(1)
_ = W705
tmp3339 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W705)


if True == tmp3339 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W705)
return
}


}, 1)

tmp3340 := MakeNative(func(__e *ControlFlow) {
W706 := __e.Get(1)
_ = W706
tmp3346 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W706)


if True == tmp3346 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3341 := MakeNative(func(__e *ControlFlow) {
W707 := __e.Get(1)
_ = W707
tmp3342 := MakeNative(func(__e *ControlFlow) {
W708 := __e.Get(1)
_ = W708
__e.TailApply(PrimFunc(symshen_4comb), W708, W707)
return
}, 1)

tmp3343 := Call(__e, PrimFunc(symshen_4in_1_6), W706)


__e.TailApply(tmp3342, tmp3343)
return


}, 1)

tmp3344 := Call(__e, PrimFunc(symshen_4_5_1out), W706)


__e.TailApply(tmp3341, tmp3344)
return


}


}, 1)

tmp3347 := Call(__e, PrimFunc(symshen_4_5pattern_6), V704)


tmp3348 := Call(__e, tmp3340, tmp3347)


__e.TailApply(tmp3337, tmp3348)
return


}, 1)

tmp3349 := Call(__e, ns2_1set, symshen_4_5pattern1_6, tmp3336)


_ = tmp3349

tmp3350 := MakeNative(func(__e *ControlFlow) {
V709 := __e.Get(1)
_ = V709
tmp3351 := MakeNative(func(__e *ControlFlow) {
W710 := __e.Get(1)
_ = W710
tmp3353 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W710)


if True == tmp3353 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W710)
return
}


}, 1)

tmp3354 := MakeNative(func(__e *ControlFlow) {
W711 := __e.Get(1)
_ = W711
tmp3360 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W711)


if True == tmp3360 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3355 := MakeNative(func(__e *ControlFlow) {
W712 := __e.Get(1)
_ = W712
tmp3356 := MakeNative(func(__e *ControlFlow) {
W713 := __e.Get(1)
_ = W713
__e.TailApply(PrimFunc(symshen_4comb), W713, W712)
return
}, 1)

tmp3357 := Call(__e, PrimFunc(symshen_4in_1_6), W711)


__e.TailApply(tmp3356, tmp3357)
return


}, 1)

tmp3358 := Call(__e, PrimFunc(symshen_4_5_1out), W711)


__e.TailApply(tmp3355, tmp3358)
return


}


}, 1)

tmp3361 := Call(__e, PrimFunc(symshen_4_5pattern_6), V709)


tmp3362 := Call(__e, tmp3354, tmp3361)


__e.TailApply(tmp3351, tmp3362)
return


}, 1)

tmp3363 := Call(__e, ns2_1set, symshen_4_5pattern2_6, tmp3350)


_ = tmp3363

tmp3364 := MakeNative(func(__e *ControlFlow) {
V714 := __e.Get(1)
_ = V714
tmp3365 := MakeNative(func(__e *ControlFlow) {
W715 := __e.Get(1)
_ = W715
tmp3366 := MakeNative(func(__e *ControlFlow) {
W716 := __e.Get(1)
_ = W716
tmp3367 := MakeNative(func(__e *ControlFlow) {
W717 := __e.Get(1)
_ = W717
__e.Return(W717)
return
}, 1)

tmp3368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symstr) {
return PrimStr(V714)
}
__typedArg0 := V714
return Call(__e, PrimFunc(symstr), __typedArg0)
})()

tmp3369 := Call(__e, PrimFunc(sym_8s), tmp3368, MakeString(")"))


tmp3370 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp3369)


tmp3371 := Call(__e, PrimFunc(sym_8s), MakeString("n"), tmp3370)


tmp3372 := Call(__e, PrimFunc(sym_8s), MakeString("f"), tmp3371)


tmp3373 := Call(__e, PrimFunc(sym_8s), MakeString("("), tmp3372)


tmp3374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(W716, MakeNumber(1), tmp3373)
}
__typedArg0 := W716
__typedArg1 := MakeNumber(1)
__typedArg2 := tmp3373
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

__e.TailApply(tmp3367, tmp3374)
return


}, 1)

tmp3375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symaddress_1_6) {
return PrimVectorSet(W715, MakeNumber(0), symshen_4printF)
}
__typedArg0 := W715
__typedArg1 := MakeNumber(0)
__typedArg2 := symshen_4printF
return Call(__e, PrimFunc(symaddress_1_6), __typedArg0, __typedArg1, __typedArg2)
})()

__e.TailApply(tmp3366, tmp3375)
return


}, 1)

tmp3376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symabsvector) {
return PrimAbsvector(MakeNumber(2))
}
__typedArg0 := MakeNumber(2)
return Call(__e, PrimFunc(symabsvector), __typedArg0)
})()

__e.TailApply(tmp3365, tmp3376)
return


}, 1)

tmp3377 := Call(__e, ns2_1set, symshen_4fn_1print, tmp3364)


_ = tmp3377

tmp3378 := MakeNative(func(__e *ControlFlow) {
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

tmp3379 := Call(__e, ns2_1set, symshen_4printF, tmp3378)


_ = tmp3379

tmp3380 := MakeNative(func(__e *ControlFlow) {
V723 := __e.Get(1)
_ = V723
V724 := __e.Get(2)
_ = V724
tmp3404 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V724)
}
__typedArg0 := V724
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3400 Obj

if True == tmp3404 {
tmp3402 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V724)
}
__typedArg0 := V724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3403 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3402)
}
__typedArg0 := Nil
__typedArg1 := tmp3402
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3401 Obj

if True == tmp3403 {
ifres3401 = True


} else {
ifres3401 = False


}

ifres3400 = ifres3401


} else {
ifres3400 = False


}

if True == ifres3400 {
tmp3381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V724)
}
__typedArg0 := V724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symlength), tmp3381)
return


} else {
tmp3398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V724)
}
__typedArg0 := V724
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3386 Obj

if True == tmp3398 {
tmp3396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V724)
}
__typedArg0 := V724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3396)
}
__typedArg0 := tmp3396
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3388 Obj

if True == tmp3397 {
tmp3390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V724)
}
__typedArg0 := V724
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3391 := Call(__e, PrimFunc(symlength), tmp3390)


tmp3392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V724)
}
__typedArg0 := V724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3392)
}
__typedArg0 := tmp3392
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3394 := Call(__e, PrimFunc(symlength), tmp3393)


tmp3395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(tmp3391, tmp3394)
}
__typedArg0 := tmp3391
__typedArg1 := tmp3394
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3389 Obj

if True == tmp3395 {
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

if True == ifres3386 {
tmp3382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V724)
}
__typedArg0 := V724
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4arity_1chk), V723, tmp3382)
return


} else {
tmp3383 := Call(__e, PrimFunc(symshen_4app), V723, MakeString("\n"), symshen_4a)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("arity error in "))
__typedS1, __typedOK1 := TypedString(tmp3383)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("arity error in ")
__typedArg1 := tmp3383
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})())
}
__typedArg0 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("arity error in "))
__typedS1, __typedOK1 := TypedString(tmp3383)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("arity error in ")
__typedArg1 := tmp3383
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()
return Call(__e, PrimFunc(symsimple_1error), __typedArg0)
})())
return


}


}


}, 2)

tmp3405 := Call(__e, ns2_1set, symshen_4arity_1chk, tmp3380)


_ = tmp3405

tmp3406 := MakeNative(func(__e *ControlFlow) {
V725 := __e.Get(1)
_ = V725
V726 := __e.Get(2)
_ = V726
tmp3412 := Call(__e, PrimFunc(symtuple_2), V726)


if True == tmp3412 {
tmp3407 := Call(__e, PrimFunc(symfst), V726)


tmp3408 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp3407)


tmp3409 := Call(__e, PrimFunc(symsnd), V726)


tmp3410 := Call(__e, PrimFunc(symshen_4find_1free_1vars), tmp3408, tmp3409)


__e.TailApply(PrimFunc(symshen_4free_1variable_1error_1message), V725, tmp3410)
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

tmp3413 := Call(__e, ns2_1set, symshen_4free_1var_1chk, tmp3406)


_ = tmp3413

tmp3414 := MakeNative(func(__e *ControlFlow) {
V727 := __e.Get(1)
_ = V727
V728 := __e.Get(2)
_ = V728
tmp3426 := Call(__e, PrimFunc(symempty_2), V728)


if True == tmp3426 {
__e.Return(symshen_4skip)
return
} else {
tmp3415 := Call(__e, PrimFunc(symshen_4app), V727, MakeString(":"), symshen_4a)


tmp3416 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString("free variables in "))
__typedS1, __typedOK1 := TypedString(tmp3415)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString("free variables in ")
__typedArg1 := tmp3415
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp3417 := Call(__e, PrimFunc(symstoutput))


tmp3418 := Call(__e, PrimFunc(sympr), tmp3416, tmp3417)


_ = tmp3418

tmp3419 := MakeNative(func(__e *ControlFlow) {
Z729 := __e.Get(1)
_ = Z729
tmp3420 := Call(__e, PrimFunc(symshen_4app), Z729, MakeString(""), symshen_4a)


tmp3421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcn) {
__typedS0, __typedOK0 := TypedString(MakeString(" "))
__typedS1, __typedOK1 := TypedString(tmp3420)
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(symcn) {
return TypedMaterializeString((__typedS0 + __typedS1))
}}
__typedArg0 := MakeString(" ")
__typedArg1 := tmp3420
return Call(__e, PrimFunc(symcn), __typedArg0, __typedArg1)
})()

tmp3422 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp3421, tmp3422)
return


}, 1)

tmp3423 := Call(__e, PrimFunc(symmap), tmp3419, V728)


_ = tmp3423

tmp3424 := Call(__e, PrimFunc(symnl), MakeNumber(1))


_ = tmp3424

__e.TailApply(PrimFunc(symabort))
return


}


}, 2)

tmp3427 := Call(__e, ns2_1set, symshen_4free_1variable_1error_1message, tmp3414)


_ = tmp3427

tmp3428 := MakeNative(func(__e *ControlFlow) {
V732 := __e.Get(1)
_ = V732
tmp3436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V732)
}
__typedArg0 := V732
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp3436 {
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
tmp3434 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V732)
}
__typedArg0 := V732
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3434 {
tmp3429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V732)
}
__typedArg0 := V732
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3430 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp3429)


tmp3431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V732)
}
__typedArg0 := V732
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3432 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp3431)


__e.TailApply(PrimFunc(symunion), tmp3430, tmp3432)
return


} else {
__e.Return(Nil)
return
}


}


}, 1)

tmp3437 := Call(__e, ns2_1set, symshen_4extract_1vars, tmp3428)


_ = tmp3437

tmp3438 := MakeNative(func(__e *ControlFlow) {
V737 := __e.Get(1)
_ = V737
V738 := __e.Get(2)
_ = V738
tmp3528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3515 Obj

if True == tmp3528 {
tmp3526 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3527 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symprotect, tmp3526)
}
__typedArg0 := symprotect
__typedArg1 := tmp3526
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3517 Obj

if True == tmp3527 {
tmp3524 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3525 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3524)
}
__typedArg0 := tmp3524
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3519 Obj

if True == tmp3525 {
tmp3521 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3522 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3521)
}
__typedArg0 := tmp3521
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3523 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3522)
}
__typedArg0 := Nil
__typedArg1 := tmp3522
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3520 Obj

if True == tmp3523 {
ifres3520 = True


} else {
ifres3520 = False


}

ifres3519 = ifres3520


} else {
ifres3519 = False


}

var ifres3518 Obj

if True == ifres3519 {
ifres3518 = True


} else {
ifres3518 = False


}

ifres3517 = ifres3518


} else {
ifres3517 = False


}

var ifres3516 Obj

if True == ifres3517 {
ifres3516 = True


} else {
ifres3516 = False


}

ifres3515 = ifres3516


} else {
ifres3515 = False


}

if True == ifres3515 {
__e.Return(Nil)
return
} else {
tmp3513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3487 Obj

if True == tmp3513 {
tmp3511 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp3511)
}
__typedArg0 := symlet
__typedArg1 := tmp3511
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3489 Obj

if True == tmp3512 {
tmp3509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3510 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3509)
}
__typedArg0 := tmp3509
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3491 Obj

if True == tmp3510 {
tmp3506 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3507 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3506)
}
__typedArg0 := tmp3506
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3508 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3507)
}
__typedArg0 := tmp3507
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3493 Obj

if True == tmp3508 {
tmp3502 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3503 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3502)
}
__typedArg0 := tmp3502
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3504 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3503)
}
__typedArg0 := tmp3503
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3505 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3504)
}
__typedArg0 := tmp3504
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3495 Obj

if True == tmp3505 {
tmp3497 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3498 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3497)
}
__typedArg0 := tmp3497
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3499 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3498)
}
__typedArg0 := tmp3498
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3500 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3499)
}
__typedArg0 := tmp3499
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3501 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3500)
}
__typedArg0 := Nil
__typedArg1 := tmp3500
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3496 Obj

if True == tmp3501 {
ifres3496 = True


} else {
ifres3496 = False


}

ifres3495 = ifres3496


} else {
ifres3495 = False


}

var ifres3494 Obj

if True == ifres3495 {
ifres3494 = True


} else {
ifres3494 = False


}

ifres3493 = ifres3494


} else {
ifres3493 = False


}

var ifres3492 Obj

if True == ifres3493 {
ifres3492 = True


} else {
ifres3492 = False


}

ifres3491 = ifres3492


} else {
ifres3491 = False


}

var ifres3490 Obj

if True == ifres3491 {
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

if True == ifres3487 {
tmp3439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3439)
}
__typedArg0 := tmp3439
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3440)
}
__typedArg0 := tmp3440
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3442 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V737, tmp3441)


tmp3443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3443)
}
__typedArg0 := tmp3443
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3444, V737)
}
__typedArg0 := tmp3444
__typedArg1 := V737
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3446)
}
__typedArg0 := tmp3446
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3447)
}
__typedArg0 := tmp3447
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3449 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3448)
}
__typedArg0 := tmp3448
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3450 := Call(__e, PrimFunc(symshen_4find_1free_1vars), tmp3445, tmp3449)


__e.TailApply(PrimFunc(symunion), tmp3442, tmp3450)
return


} else {
tmp3485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3466 Obj

if True == tmp3485 {
tmp3483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp3483)
}
__typedArg0 := symlambda
__typedArg1 := tmp3483
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3468 Obj

if True == tmp3484 {
tmp3481 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3482 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3481)
}
__typedArg0 := tmp3481
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3470 Obj

if True == tmp3482 {
tmp3478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3479 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3478)
}
__typedArg0 := tmp3478
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3480 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3479)
}
__typedArg0 := tmp3479
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3472 Obj

if True == tmp3480 {
tmp3474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3474)
}
__typedArg0 := tmp3474
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3476 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3475)
}
__typedArg0 := tmp3475
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3477 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3476)
}
__typedArg0 := Nil
__typedArg1 := tmp3476
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3473 Obj

if True == tmp3477 {
ifres3473 = True


} else {
ifres3473 = False


}

ifres3472 = ifres3473


} else {
ifres3472 = False


}

var ifres3471 Obj

if True == ifres3472 {
ifres3471 = True


} else {
ifres3471 = False


}

ifres3470 = ifres3471


} else {
ifres3470 = False


}

var ifres3469 Obj

if True == ifres3470 {
ifres3469 = True


} else {
ifres3469 = False


}

ifres3468 = ifres3469


} else {
ifres3468 = False


}

var ifres3467 Obj

if True == ifres3468 {
ifres3467 = True


} else {
ifres3467 = False


}

ifres3466 = ifres3467


} else {
ifres3466 = False


}

if True == ifres3466 {
tmp3451 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3452, V737)
}
__typedArg0 := tmp3452
__typedArg1 := V737
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3454 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3455 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3454)
}
__typedArg0 := tmp3454
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3456 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3455)
}
__typedArg0 := tmp3455
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4find_1free_1vars), tmp3453, tmp3456)
return


} else {
tmp3464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3464 {
tmp3457 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3458 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V737, tmp3457)


tmp3459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V738)
}
__typedArg0 := V738
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3460 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V737, tmp3459)


__e.TailApply(PrimFunc(symunion), tmp3458, tmp3460)
return


} else {
tmp3462 := Call(__e, PrimFunc(symshen_4free_1variable_2), V738, V737)


if True == tmp3462 {
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

tmp3529 := Call(__e, ns2_1set, symshen_4find_1free_1vars, tmp3438)


_ = tmp3529

tmp3530 := MakeNative(func(__e *ControlFlow) {
V739 := __e.Get(1)
_ = V739
V740 := __e.Get(2)
_ = V740
tmp3535 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(V739)
}
__typedArg0 := V739
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

if True == tmp3535 {
tmp3532 := Call(__e, PrimFunc(symelement_2), V739, V740)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symnot) {
__typedB0, __typedOK0 := TypedBoolean(tmp3532)
if __typedOK0 && HasCanonicalPrimitiveBinding(symnot) {
return TypedMaterializeBoolean((!__typedB0))
}}
__typedArg0 := tmp3532
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

tmp3536 := Call(__e, ns2_1set, symshen_4free_1variable_2, tmp3530)


_ = tmp3536

tmp3537 := MakeNative(func(__e *ControlFlow) {
V741 := __e.Get(1)
_ = V741
V742 := __e.Get(2)
_ = V742
tmp3538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_duserdefs_d)
}
__typedArg0 := symshen_4_duserdefs_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

tmp3539 := Call(__e, PrimFunc(symadjoin), V741, tmp3538)


tmp3540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symset) {
return PrimSet(symshen_4_duserdefs_d, tmp3539)
}
__typedArg0 := symshen_4_duserdefs_d
__typedArg1 := tmp3539
return Call(__e, PrimFunc(symset), __typedArg0, __typedArg1)
})()

_ = tmp3540

tmp3541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(sym_dproperty_1vector_d)
}
__typedArg0 := sym_dproperty_1vector_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

__e.TailApply(PrimFunc(symput), V741, symshen_4source, V742, tmp3541)
return


}, 2)

tmp3542 := Call(__e, ns2_1set, symshen_4record_1kl, tmp3537)


_ = tmp3542

tmp3543 := MakeNative(func(__e *ControlFlow) {
V743 := __e.Get(1)
_ = V743
V744 := __e.Get(2)
_ = V744
V745 := __e.Get(3)
_ = V745
tmp3544 := MakeNative(func(__e *ControlFlow) {
W746 := __e.Get(1)
_ = W746
tmp3545 := MakeNative(func(__e *ControlFlow) {
W747 := __e.Get(1)
_ = W747
tmp3546 := MakeNative(func(__e *ControlFlow) {
W748 := __e.Get(1)
_ = W748
__e.Return(W748)
return
}, 1)

tmp3547 := Call(__e, PrimFunc(symshen_4cond_1form), W747)


tmp3548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3547, Nil)
}
__typedArg0 := tmp3547
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W746, tmp3548)
}
__typedArg0 := W746
__typedArg1 := tmp3548
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V743, tmp3549)
}
__typedArg0 := V743
__typedArg1 := tmp3549
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp3550)
}
__typedArg0 := symdefun
__typedArg1 := tmp3550
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp3546, tmp3551)
return


}, 1)

tmp3552 := Call(__e, PrimFunc(symshen_4kl_1body), V744, W746)


tmp3553 := Call(__e, PrimFunc(symshen_4scan_1body), V743, tmp3552)


__e.TailApply(tmp3545, tmp3553)
return


}, 1)

tmp3554 := Call(__e, PrimFunc(symshen_4parameters), V745)


__e.TailApply(tmp3544, tmp3554)
return


}, 3)

tmp3555 := Call(__e, ns2_1set, symshen_4compile_1to_1kl, tmp3543)


_ = tmp3555

tmp3556 := MakeNative(func(__e *ControlFlow) {
V749 := __e.Get(1)
_ = V749
tmp3561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeNumber(0), V749)
}
__typedArg0 := MakeNumber(0)
__typedArg1 := V749
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp3561 {
__e.Return(Nil)
return
} else {
tmp3557 := Call(__e, PrimFunc(symgensym), symV)


tmp3559 := Call(__e, PrimFunc(symshen_4parameters), (func() Obj {
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
return PrimCons(tmp3557, tmp3559)
}
__typedArg0 := tmp3557
__typedArg1 := tmp3559
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp3562 := Call(__e, ns2_1set, symshen_4parameters, tmp3556)


_ = tmp3562

tmp3563 := MakeNative(func(__e *ControlFlow) {
V752 := __e.Get(1)
_ = V752
tmp3587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3567 Obj

if True == tmp3587 {
tmp3585 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3586 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3585)
}
__typedArg0 := tmp3585
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3569 Obj

if True == tmp3586 {
tmp3582 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3583 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3582)
}
__typedArg0 := tmp3582
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3584 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, tmp3583)
}
__typedArg0 := True
__typedArg1 := tmp3583
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3571 Obj

if True == tmp3584 {
tmp3579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3579)
}
__typedArg0 := tmp3579
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3580)
}
__typedArg0 := tmp3580
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3573 Obj

if True == tmp3581 {
tmp3575 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3575)
}
__typedArg0 := tmp3575
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3577 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3576)
}
__typedArg0 := tmp3576
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3578 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3577)
}
__typedArg0 := Nil
__typedArg1 := tmp3577
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3574 Obj

if True == tmp3578 {
ifres3574 = True


} else {
ifres3574 = False


}

ifres3573 = ifres3574


} else {
ifres3573 = False


}

var ifres3572 Obj

if True == ifres3573 {
ifres3572 = True


} else {
ifres3572 = False


}

ifres3571 = ifres3572


} else {
ifres3571 = False


}

var ifres3570 Obj

if True == ifres3571 {
ifres3570 = True


} else {
ifres3570 = False


}

ifres3569 = ifres3570


} else {
ifres3569 = False


}

var ifres3568 Obj

if True == ifres3569 {
ifres3568 = True


} else {
ifres3568 = False


}

ifres3567 = ifres3568


} else {
ifres3567 = False


}

if True == ifres3567 {
tmp3564 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V752)
}
__typedArg0 := V752
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3564)
}
__typedArg0 := tmp3564
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3565)
}
__typedArg0 := tmp3565
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

tmp3588 := Call(__e, ns2_1set, symshen_4cond_1form, tmp3563)


_ = tmp3588

tmp3589 := MakeNative(func(__e *ControlFlow) {
V761 := __e.Get(1)
_ = V761
V762 := __e.Get(2)
_ = V762
tmp3633 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V762)
}
__typedArg0 := Nil
__typedArg1 := V762
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp3633 {
tmp3590 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V761, Nil)
}
__typedArg0 := V761
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3591 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symshen_4f_1error, tmp3590)
}
__typedArg0 := symshen_4f_1error
__typedArg1 := tmp3590
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3591, Nil)
}
__typedArg0 := tmp3591
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp3592)
}
__typedArg0 := True
__typedArg1 := tmp3592
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3593, Nil)
}
__typedArg0 := tmp3593
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp3631 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3627 Obj

if True == tmp3631 {
tmp3629 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3630 := Call(__e, PrimFunc(symshen_4choicepoint_2), tmp3629)


var ifres3628 Obj

if True == tmp3630 {
ifres3628 = True


} else {
ifres3628 = False


}

ifres3627 = ifres3628


} else {
ifres3627 = False


}

if True == ifres3627 {
tmp3594 := Call(__e, PrimFunc(symgensym), symFreeze)


tmp3595 := Call(__e, PrimFunc(symgensym), symResult)


tmp3596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4choicepoint), V761, tmp3594, tmp3595, tmp3596, tmp3597)
return


} else {
tmp3625 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3605 Obj

if True == tmp3625 {
tmp3623 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3624 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3623)
}
__typedArg0 := tmp3623
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3607 Obj

if True == tmp3624 {
tmp3620 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3621 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3620)
}
__typedArg0 := tmp3620
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3622 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, tmp3621)
}
__typedArg0 := True
__typedArg1 := tmp3621
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3609 Obj

if True == tmp3622 {
tmp3617 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3618 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3617)
}
__typedArg0 := tmp3617
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3619 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3618)
}
__typedArg0 := tmp3618
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3611 Obj

if True == tmp3619 {
tmp3613 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3614 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3613)
}
__typedArg0 := tmp3613
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3615 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3614)
}
__typedArg0 := tmp3614
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3616 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3615)
}
__typedArg0 := Nil
__typedArg1 := tmp3615
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3612 Obj

if True == tmp3616 {
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

var ifres3608 Obj

if True == ifres3609 {
ifres3608 = True


} else {
ifres3608 = False


}

ifres3607 = ifres3608


} else {
ifres3607 = False


}

var ifres3606 Obj

if True == ifres3607 {
ifres3606 = True


} else {
ifres3606 = False


}

ifres3605 = ifres3606


} else {
ifres3605 = False


}

if True == ifres3605 {
tmp3598 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3598, Nil)
}
__typedArg0 := tmp3598
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp3603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3603 {
tmp3599 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V762)
}
__typedArg0 := V762
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3601 := Call(__e, PrimFunc(symshen_4scan_1body), V761, tmp3600)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3599, tmp3601)
}
__typedArg0 := tmp3599
__typedArg1 := tmp3601
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

tmp3634 := Call(__e, ns2_1set, symshen_4scan_1body, tmp3589)


_ = tmp3634

tmp3635 := MakeNative(func(__e *ControlFlow) {
V769 := __e.Get(1)
_ = V769
tmp3670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V769)
}
__typedArg0 := V769
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3637 Obj

if True == tmp3670 {
tmp3668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V769)
}
__typedArg0 := V769
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3668)
}
__typedArg0 := tmp3668
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3639 Obj

if True == tmp3669 {
tmp3665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V769)
}
__typedArg0 := V769
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3665)
}
__typedArg0 := tmp3665
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3666)
}
__typedArg0 := tmp3666
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3641 Obj

if True == tmp3667 {
tmp3661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V769)
}
__typedArg0 := V769
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3661)
}
__typedArg0 := tmp3661
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3662)
}
__typedArg0 := tmp3662
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4choicepoint_b, tmp3663)
}
__typedArg0 := symshen_4choicepoint_b
__typedArg1 := tmp3663
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3643 Obj

if True == tmp3664 {
tmp3657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V769)
}
__typedArg0 := V769
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3657)
}
__typedArg0 := tmp3657
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3658)
}
__typedArg0 := tmp3658
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3659)
}
__typedArg0 := tmp3659
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3645 Obj

if True == tmp3660 {
tmp3652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V769)
}
__typedArg0 := V769
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3653 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3652)
}
__typedArg0 := tmp3652
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3654 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3653)
}
__typedArg0 := tmp3653
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3655 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3654)
}
__typedArg0 := tmp3654
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3656 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3655)
}
__typedArg0 := Nil
__typedArg1 := tmp3655
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3647 Obj

if True == tmp3656 {
tmp3649 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V769)
}
__typedArg0 := V769
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3650)
}
__typedArg0 := Nil
__typedArg1 := tmp3650
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3648 Obj

if True == tmp3651 {
ifres3648 = True


} else {
ifres3648 = False


}

ifres3647 = ifres3648


} else {
ifres3647 = False


}

var ifres3646 Obj

if True == ifres3647 {
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

if True == ifres3637 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp3671 := Call(__e, ns2_1set, symshen_4choicepoint_2, tmp3635)


_ = tmp3671

tmp3672 := MakeNative(func(__e *ControlFlow) {
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
tmp3864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3786 Obj

if True == tmp3864 {
tmp3862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3863 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3862)
}
__typedArg0 := tmp3862
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3788 Obj

if True == tmp3863 {
tmp3859 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3860 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3859)
}
__typedArg0 := tmp3859
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3861 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3860)
}
__typedArg0 := tmp3860
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3790 Obj

if True == tmp3861 {
tmp3855 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3855)
}
__typedArg0 := tmp3855
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3857 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3856)
}
__typedArg0 := tmp3856
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3857)
}
__typedArg0 := tmp3857
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3792 Obj

if True == tmp3858 {
tmp3850 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3851 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3850)
}
__typedArg0 := tmp3850
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3851)
}
__typedArg0 := tmp3851
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3852)
}
__typedArg0 := tmp3852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3853)
}
__typedArg0 := tmp3853
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3794 Obj

if True == tmp3854 {
tmp3844 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3845 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3844)
}
__typedArg0 := tmp3844
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3846 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3845)
}
__typedArg0 := tmp3845
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3846)
}
__typedArg0 := tmp3846
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3847)
}
__typedArg0 := tmp3847
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfail_1if, tmp3848)
}
__typedArg0 := symfail_1if
__typedArg1 := tmp3848
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3796 Obj

if True == tmp3849 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3839)
}
__typedArg0 := tmp3839
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3840)
}
__typedArg0 := tmp3840
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3842 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3841)
}
__typedArg0 := tmp3841
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3842)
}
__typedArg0 := tmp3842
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3798 Obj

if True == tmp3843 {
tmp3831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
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

var ifres3800 Obj

if True == tmp3837 {
tmp3823 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3824 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3823)
}
__typedArg0 := tmp3823
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3825 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3824)
}
__typedArg0 := tmp3824
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3826 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3825)
}
__typedArg0 := tmp3825
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3827 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3826)
}
__typedArg0 := tmp3826
return Call(__e, PrimFunc(symtl), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3829)
}
__typedArg0 := Nil
__typedArg1 := tmp3829
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3802 Obj

if True == tmp3830 {
tmp3818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3819)
}
__typedArg0 := tmp3819
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3820)
}
__typedArg0 := tmp3820
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3822 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3821)
}
__typedArg0 := Nil
__typedArg1 := tmp3821
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3804 Obj

if True == tmp3822 {
tmp3815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3815)
}
__typedArg0 := tmp3815
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3816)
}
__typedArg0 := Nil
__typedArg1 := tmp3816
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3806 Obj

if True == tmp3817 {
tmp3808 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3809 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3808)
}
__typedArg0 := tmp3808
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3810 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3809)
}
__typedArg0 := tmp3809
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3811 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3810)
}
__typedArg0 := tmp3810
return Call(__e, PrimFunc(symhd), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V785, tmp3813)
}
__typedArg0 := V785
__typedArg1 := tmp3813
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3807 Obj

if True == tmp3814 {
ifres3807 = True


} else {
ifres3807 = False


}

ifres3806 = ifres3807


} else {
ifres3806 = False


}

var ifres3805 Obj

if True == ifres3806 {
ifres3805 = True


} else {
ifres3805 = False


}

ifres3804 = ifres3805


} else {
ifres3804 = False


}

var ifres3803 Obj

if True == ifres3804 {
ifres3803 = True


} else {
ifres3803 = False


}

ifres3802 = ifres3803


} else {
ifres3802 = False


}

var ifres3801 Obj

if True == ifres3802 {
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

if True == ifres3786 {
tmp3673 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3674 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3673)
}
__typedArg0 := tmp3673
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3674)
}
__typedArg0 := tmp3674
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3675)
}
__typedArg0 := tmp3675
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3677)
}
__typedArg0 := tmp3677
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3679 := Call(__e, PrimFunc(symshen_4scan_1body), tmp3678, V789)


tmp3680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcond, tmp3679)
}
__typedArg0 := symcond
__typedArg1 := tmp3679
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3680, Nil)
}
__typedArg0 := tmp3680
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3682 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp3681)
}
__typedArg0 := symfreeze
__typedArg1 := tmp3681
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3683 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3684 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3684)
}
__typedArg0 := tmp3684
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3686 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3685)
}
__typedArg0 := tmp3685
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3687 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3686)
}
__typedArg0 := tmp3686
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3688 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3687)
}
__typedArg0 := tmp3687
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3688)
}
__typedArg0 := tmp3688
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3689)
}
__typedArg0 := tmp3689
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3691 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3692 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3691)
}
__typedArg0 := tmp3691
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3692)
}
__typedArg0 := tmp3692
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3693)
}
__typedArg0 := tmp3693
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3694)
}
__typedArg0 := tmp3694
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3696 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3695)
}
__typedArg0 := tmp3695
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3697 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V787, Nil)
}
__typedArg0 := V787
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3696, tmp3697)
}
__typedArg0 := tmp3696
__typedArg1 := tmp3697
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V786, Nil)
}
__typedArg0 := V786
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp3699)
}
__typedArg0 := symthaw
__typedArg1 := tmp3699
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V787, Nil)
}
__typedArg0 := V787
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3700, tmp3701)
}
__typedArg0 := tmp3700
__typedArg1 := tmp3701
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3698, tmp3702)
}
__typedArg0 := tmp3698
__typedArg1 := tmp3702
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp3703)
}
__typedArg0 := symif
__typedArg1 := tmp3703
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3704, Nil)
}
__typedArg0 := tmp3704
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3690, tmp3705)
}
__typedArg0 := tmp3690
__typedArg1 := tmp3705
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3707 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V787, tmp3706)
}
__typedArg0 := V787
__typedArg1 := tmp3706
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3708 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp3707)
}
__typedArg0 := symlet
__typedArg1 := tmp3707
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V786, Nil)
}
__typedArg0 := V786
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp3709)
}
__typedArg0 := symthaw
__typedArg1 := tmp3709
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3710, Nil)
}
__typedArg0 := tmp3710
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3708, tmp3711)
}
__typedArg0 := tmp3708
__typedArg1 := tmp3711
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3683, tmp3712)
}
__typedArg0 := tmp3683
__typedArg1 := tmp3712
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp3713)
}
__typedArg0 := symif
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
return PrimCons(tmp3682, tmp3715)
}
__typedArg0 := tmp3682
__typedArg1 := tmp3715
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V786, tmp3716)
}
__typedArg0 := V786
__typedArg1 := tmp3716
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp3717)
}
__typedArg0 := symlet
__typedArg1 := tmp3717
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3718, Nil)
}
__typedArg0 := tmp3718
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp3719)
}
__typedArg0 := True
__typedArg1 := tmp3719
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3720, Nil)
}
__typedArg0 := tmp3720
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp3784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3757 Obj

if True == tmp3784 {
tmp3782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3782)
}
__typedArg0 := tmp3782
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3759 Obj

if True == tmp3783 {
tmp3779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3779)
}
__typedArg0 := tmp3779
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3780)
}
__typedArg0 := tmp3780
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3761 Obj

if True == tmp3781 {
tmp3775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3777)
}
__typedArg0 := tmp3777
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3763 Obj

if True == tmp3778 {
tmp3770 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3770)
}
__typedArg0 := tmp3770
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3771)
}
__typedArg0 := tmp3771
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

var ifres3765 Obj

if True == tmp3774 {
tmp3767 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3768 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3767)
}
__typedArg0 := tmp3767
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3769 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3768)
}
__typedArg0 := Nil
__typedArg1 := tmp3768
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3766 Obj

if True == tmp3769 {
ifres3766 = True


} else {
ifres3766 = False


}

ifres3765 = ifres3766


} else {
ifres3765 = False


}

var ifres3764 Obj

if True == ifres3765 {
ifres3764 = True


} else {
ifres3764 = False


}

ifres3763 = ifres3764


} else {
ifres3763 = False


}

var ifres3762 Obj

if True == ifres3763 {
ifres3762 = True


} else {
ifres3762 = False


}

ifres3761 = ifres3762


} else {
ifres3761 = False


}

var ifres3760 Obj

if True == ifres3761 {
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

if True == ifres3757 {
tmp3721 := Call(__e, PrimFunc(symshen_4scan_1body), V785, V789)


tmp3722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symcond, tmp3721)
}
__typedArg0 := symcond
__typedArg1 := tmp3721
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3722, Nil)
}
__typedArg0 := tmp3722
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3724 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp3723)
}
__typedArg0 := symfreeze
__typedArg1 := tmp3723
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3725 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V788)
}
__typedArg0 := V788
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3726)
}
__typedArg0 := tmp3726
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3727)
}
__typedArg0 := tmp3727
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3728)
}
__typedArg0 := tmp3728
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3730 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfail, Nil)
}
__typedArg0 := symfail
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3731 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3730, Nil)
}
__typedArg0 := tmp3730
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3732 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V787, tmp3731)
}
__typedArg0 := V787
__typedArg1 := tmp3731
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3733 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp3732)
}
__typedArg0 := sym_a
__typedArg1 := tmp3732
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3734 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V786, Nil)
}
__typedArg0 := V786
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3735 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp3734)
}
__typedArg0 := symthaw
__typedArg1 := tmp3734
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3736 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V787, Nil)
}
__typedArg0 := V787
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3737 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3735, tmp3736)
}
__typedArg0 := tmp3735
__typedArg1 := tmp3736
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3738 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3733, tmp3737)
}
__typedArg0 := tmp3733
__typedArg1 := tmp3737
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp3738)
}
__typedArg0 := symif
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
return PrimCons(tmp3729, tmp3740)
}
__typedArg0 := tmp3729
__typedArg1 := tmp3740
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V787, tmp3741)
}
__typedArg0 := V787
__typedArg1 := tmp3741
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp3742)
}
__typedArg0 := symlet
__typedArg1 := tmp3742
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3744 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V786, Nil)
}
__typedArg0 := V786
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3745 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp3744)
}
__typedArg0 := symthaw
__typedArg1 := tmp3744
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3745, Nil)
}
__typedArg0 := tmp3745
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3747 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3743, tmp3746)
}
__typedArg0 := tmp3743
__typedArg1 := tmp3746
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3748 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3725, tmp3747)
}
__typedArg0 := tmp3725
__typedArg1 := tmp3747
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp3748)
}
__typedArg0 := symif
__typedArg1 := tmp3748
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3749, Nil)
}
__typedArg0 := tmp3749
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3724, tmp3750)
}
__typedArg0 := tmp3724
__typedArg1 := tmp3750
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3752 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V786, tmp3751)
}
__typedArg0 := V786
__typedArg1 := tmp3751
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3753 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp3752)
}
__typedArg0 := symlet
__typedArg1 := tmp3752
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3754 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3753, Nil)
}
__typedArg0 := tmp3753
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3755 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp3754)
}
__typedArg0 := True
__typedArg1 := tmp3754
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3755, Nil)
}
__typedArg0 := tmp3755
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

tmp3865 := Call(__e, ns2_1set, symshen_4choicepoint, tmp3672)


_ = tmp3865

tmp3866 := MakeNative(func(__e *ControlFlow) {
V791 := __e.Get(1)
_ = V791
V792 := __e.Get(2)
_ = V792
V793 := __e.Get(3)
_ = V793
tmp3880 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V791, V793)
}
__typedArg0 := V791
__typedArg1 := V793
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp3880 {
__e.Return(V792)
return
} else {
tmp3878 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V793)
}
__typedArg0 := V793
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3878 {
tmp3867 := MakeNative(func(__e *ControlFlow) {
W794 := __e.Get(1)
_ = W794
tmp3873 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V793)
}
__typedArg0 := V793
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3874 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(W794, tmp3873)
}
__typedArg0 := W794
__typedArg1 := tmp3873
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp3874 {
tmp3868 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V793)
}
__typedArg0 := V793
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3869 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V793)
}
__typedArg0 := V793
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3870 := Call(__e, PrimFunc(symshen_4rep_1X), V791, V792, tmp3869)


__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3868, tmp3870)
}
__typedArg0 := tmp3868
__typedArg1 := tmp3870
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp3871 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V793)
}
__typedArg0 := V793
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W794, tmp3871)
}
__typedArg0 := W794
__typedArg1 := tmp3871
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}, 1)

tmp3875 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V793)
}
__typedArg0 := V793
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3876 := Call(__e, PrimFunc(symshen_4rep_1X), V791, V792, tmp3875)


__e.TailApply(tmp3867, tmp3876)
return


} else {
__e.Return(V793)
return
}


}


}, 3)

tmp3881 := Call(__e, ns2_1set, symshen_4rep_1X, tmp3866)


_ = tmp3881

tmp3882 := MakeNative(func(__e *ControlFlow) {
V795 := __e.Get(1)
_ = V795
V796 := __e.Get(2)
_ = V796
tmp3883 := MakeNative(func(__e *ControlFlow) {
Z797 := __e.Get(1)
_ = Z797
tmp3884 := Call(__e, PrimFunc(symfst), Z797)


tmp3885 := Call(__e, PrimFunc(symsnd), Z797)


tmp3886 := Call(__e, PrimFunc(symshen_4alpha_1convert), tmp3885)


__e.TailApply(PrimFunc(symshen_4triple_1stack), Nil, tmp3884, V796, tmp3886)
return


}, 1)

__e.TailApply(PrimFunc(symmap), tmp3883, V795)
return


}, 2)

tmp3887 := Call(__e, ns2_1set, symshen_4kl_1body, tmp3882)


_ = tmp3887

tmp3888 := MakeNative(func(__e *ControlFlow) {
V798 := __e.Get(1)
_ = V798
tmp3971 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3952 Obj

if True == tmp3971 {
tmp3969 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3970 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp3969)
}
__typedArg0 := symlambda
__typedArg1 := tmp3969
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3954 Obj

if True == tmp3970 {
tmp3967 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3968 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3967)
}
__typedArg0 := tmp3967
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3956 Obj

if True == tmp3968 {
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3965)
}
__typedArg0 := tmp3965
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3958 Obj

if True == tmp3966 {
tmp3960 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V798)
}
__typedArg0 := V798
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3961)
}
__typedArg0 := tmp3961
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3963 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3962)
}
__typedArg0 := Nil
__typedArg1 := tmp3962
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3959 Obj

if True == tmp3963 {
ifres3959 = True


} else {
ifres3959 = False


}

ifres3958 = ifres3959


} else {
ifres3958 = False


}

var ifres3957 Obj

if True == ifres3958 {
ifres3957 = True


} else {
ifres3957 = False


}

ifres3956 = ifres3957


} else {
ifres3956 = False


}

var ifres3955 Obj

if True == ifres3956 {
ifres3955 = True


} else {
ifres3955 = False


}

ifres3954 = ifres3955


} else {
ifres3954 = False


}

var ifres3953 Obj

if True == ifres3954 {
ifres3953 = True


} else {
ifres3953 = False


}

ifres3952 = ifres3953


} else {
ifres3952 = False


}

if True == ifres3952 {
tmp3889 := MakeNative(func(__e *ControlFlow) {
W799 := __e.Get(1)
_ = W799
tmp3890 := MakeNative(func(__e *ControlFlow) {
W800 := __e.Get(1)
_ = W800
tmp3891 := MakeNative(func(__e *ControlFlow) {
Z801 := __e.Get(1)
_ = Z801
__e.TailApply(PrimFunc(symshen_4alpha_1convert), Z801)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp3891, W800)
return


}, 1)

tmp3892 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V798)
}
__typedArg0 := V798
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
return PrimTail(V798)
}
__typedArg0 := V798
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3895)
}
__typedArg0 := tmp3895
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3897 := Call(__e, PrimFunc(symshen_4beta), tmp3893, W799, tmp3896)


tmp3898 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3897, Nil)
}
__typedArg0 := tmp3897
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3899 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W799, tmp3898)
}
__typedArg0 := W799
__typedArg1 := tmp3898
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3900 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlambda, tmp3899)
}
__typedArg0 := symlambda
__typedArg1 := tmp3899
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp3890, tmp3900)
return


}, 1)

tmp3901 := Call(__e, PrimFunc(symgensym), symZ)


__e.TailApply(tmp3889, tmp3901)
return


} else {
tmp3950 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3924 Obj

if True == tmp3950 {
tmp3948 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3949 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp3948)
}
__typedArg0 := symlet
__typedArg1 := tmp3948
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3926 Obj

if True == tmp3949 {
tmp3946 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3947 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3946)
}
__typedArg0 := tmp3946
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3928 Obj

if True == tmp3947 {
tmp3943 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V798)
}
__typedArg0 := V798
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3944)
}
__typedArg0 := tmp3944
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3930 Obj

if True == tmp3945 {
tmp3939 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3940 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3939)
}
__typedArg0 := tmp3939
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3941 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3940)
}
__typedArg0 := tmp3940
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3942 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp3941)
}
__typedArg0 := tmp3941
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres3932 Obj

if True == tmp3942 {
tmp3934 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3935 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3934)
}
__typedArg0 := tmp3934
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3936 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3935)
}
__typedArg0 := tmp3935
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3937 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3936)
}
__typedArg0 := tmp3936
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3938 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp3937)
}
__typedArg0 := Nil
__typedArg1 := tmp3937
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres3933 Obj

if True == tmp3938 {
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

if True == ifres3924 {
tmp3902 := MakeNative(func(__e *ControlFlow) {
W802 := __e.Get(1)
_ = W802
tmp3903 := MakeNative(func(__e *ControlFlow) {
W803 := __e.Get(1)
_ = W803
tmp3904 := MakeNative(func(__e *ControlFlow) {
Z804 := __e.Get(1)
_ = Z804
__e.TailApply(PrimFunc(symshen_4alpha_1convert), Z804)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp3904, W803)
return


}, 1)

tmp3905 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3906 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3905)
}
__typedArg0 := tmp3905
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3907 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3906)
}
__typedArg0 := tmp3906
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3908 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3909 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3908)
}
__typedArg0 := tmp3908
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3910 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3911 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3910)
}
__typedArg0 := tmp3910
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3912 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3911)
}
__typedArg0 := tmp3911
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3913 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3912)
}
__typedArg0 := tmp3912
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3914 := Call(__e, PrimFunc(symshen_4beta), tmp3909, W802, tmp3913)


tmp3915 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3914, Nil)
}
__typedArg0 := tmp3914
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3916 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3907, tmp3915)
}
__typedArg0 := tmp3907
__typedArg1 := tmp3915
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3917 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W802, tmp3916)
}
__typedArg0 := W802
__typedArg1 := tmp3916
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3918 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp3917)
}
__typedArg0 := symlet
__typedArg1 := tmp3917
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp3903, tmp3918)
return


}, 1)

tmp3919 := Call(__e, PrimFunc(symgensym), symW)


__e.TailApply(tmp3902, tmp3919)
return


} else {
tmp3922 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V798)
}
__typedArg0 := V798
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp3922 {
tmp3920 := MakeNative(func(__e *ControlFlow) {
Z805 := __e.Get(1)
_ = Z805
__e.TailApply(PrimFunc(symshen_4alpha_1convert), Z805)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp3920, V798)
return


} else {
__e.Return(V798)
return
}


}


}


}, 1)

tmp3972 := Call(__e, ns2_1set, symshen_4alpha_1convert, tmp3888)


_ = tmp3972

tmp3973 := MakeNative(func(__e *ControlFlow) {
V814 := __e.Get(1)
_ = V814
V815 := __e.Get(2)
_ = V815
V816 := __e.Get(3)
_ = V816
V817 := __e.Get(4)
_ = V817
tmp4103 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V815)
}
__typedArg0 := Nil
__typedArg1 := V815
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4078 Obj

if True == tmp4103 {
tmp4102 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V816)
}
__typedArg0 := Nil
__typedArg1 := V816
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4080 Obj

if True == tmp4102 {
tmp4101 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V817)
}
__typedArg0 := V817
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4082 Obj

if True == tmp4101 {
tmp4099 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V817)
}
__typedArg0 := V817
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4100 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symwhere, tmp4099)
}
__typedArg0 := symwhere
__typedArg1 := tmp4099
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4084 Obj

if True == tmp4100 {
tmp4097 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V817)
}
__typedArg0 := V817
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4098 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4097)
}
__typedArg0 := tmp4097
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4086 Obj

if True == tmp4098 {
tmp4094 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V817)
}
__typedArg0 := V817
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4095 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4094)
}
__typedArg0 := tmp4094
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4096 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4095)
}
__typedArg0 := tmp4095
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4088 Obj

if True == tmp4096 {
tmp4090 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V817)
}
__typedArg0 := V817
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4091)
}
__typedArg0 := tmp4091
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4093 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4092)
}
__typedArg0 := Nil
__typedArg1 := tmp4092
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4089 Obj

if True == tmp4093 {
ifres4089 = True


} else {
ifres4089 = False


}

ifres4088 = ifres4089


} else {
ifres4088 = False


}

var ifres4087 Obj

if True == ifres4088 {
ifres4087 = True


} else {
ifres4087 = False


}

ifres4086 = ifres4087


} else {
ifres4086 = False


}

var ifres4085 Obj

if True == ifres4086 {
ifres4085 = True


} else {
ifres4085 = False


}

ifres4084 = ifres4085


} else {
ifres4084 = False


}

var ifres4083 Obj

if True == ifres4084 {
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

if True == ifres4078 {
tmp3974 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V817)
}
__typedArg0 := V817
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3975 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3974)
}
__typedArg0 := tmp3974
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3976 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3975, V814)
}
__typedArg0 := tmp3975
__typedArg1 := V814
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3977 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V817)
}
__typedArg0 := V817
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3978 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3977)
}
__typedArg0 := tmp3977
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3979 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3978)
}
__typedArg0 := tmp3978
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp3976, Nil, Nil, tmp3979)
return


} else {
tmp4076 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V815)
}
__typedArg0 := Nil
__typedArg1 := V815
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4073 Obj

if True == tmp4076 {
tmp4075 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V816)
}
__typedArg0 := Nil
__typedArg1 := V816
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4074 Obj

if True == tmp4075 {
ifres4074 = True


} else {
ifres4074 = False


}

ifres4073 = ifres4074


} else {
ifres4073 = False


}

if True == ifres4073 {
tmp3980 := Call(__e, PrimFunc(symreverse), V814)


tmp3981 := Call(__e, PrimFunc(symshen_4rectify_1test), tmp3980)


tmp3982 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V817, Nil)
}
__typedArg0 := V817
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3981, tmp3982)
}
__typedArg0 := tmp3981
__typedArg1 := tmp3982
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp4071 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4064 Obj

if True == tmp4071 {
tmp4070 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V816)
}
__typedArg0 := V816
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4066 Obj

if True == tmp4070 {
tmp4068 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4069 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvariable_2) {
return PrimIsVariable(tmp4068)
}
__typedArg0 := tmp4068
return Call(__e, PrimFunc(symvariable_2), __typedArg0)
})()

var ifres4067 Obj

if True == tmp4069 {
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
tmp3983 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3984 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V816)
}
__typedArg0 := V816
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3985 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3986 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V816)
}
__typedArg0 := V816
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3987 := Call(__e, PrimFunc(symshen_4beta), tmp3985, tmp3986, V817)


__e.TailApply(PrimFunc(symshen_4triple_1stack), V814, tmp3983, tmp3984, tmp3987)
return


} else {
tmp4062 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4037 Obj

if True == tmp4062 {
tmp4060 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4061 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4060)
}
__typedArg0 := tmp4060
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4039 Obj

if True == tmp4061 {
tmp4057 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4058 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4057)
}
__typedArg0 := tmp4057
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4059 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4058)
}
__typedArg0 := tmp4058
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4041 Obj

if True == tmp4059 {
tmp4053 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4054 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4053)
}
__typedArg0 := tmp4053
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4055 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4054)
}
__typedArg0 := tmp4054
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4056 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4055)
}
__typedArg0 := tmp4055
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4043 Obj

if True == tmp4056 {
tmp4048 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4049 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4048)
}
__typedArg0 := tmp4048
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4050 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4049)
}
__typedArg0 := tmp4049
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4051 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4050)
}
__typedArg0 := tmp4050
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4052 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4051)
}
__typedArg0 := Nil
__typedArg1 := tmp4051
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4045 Obj

if True == tmp4052 {
tmp4047 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V816)
}
__typedArg0 := V816
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4046 Obj

if True == tmp4047 {
ifres4046 = True


} else {
ifres4046 = False


}

ifres4045 = ifres4046


} else {
ifres4045 = False


}

var ifres4044 Obj

if True == ifres4045 {
ifres4044 = True


} else {
ifres4044 = False


}

ifres4043 = ifres4044


} else {
ifres4043 = False


}

var ifres4042 Obj

if True == ifres4043 {
ifres4042 = True


} else {
ifres4042 = False


}

ifres4041 = ifres4042


} else {
ifres4041 = False


}

var ifres4040 Obj

if True == ifres4041 {
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

if True == ifres4037 {
tmp3988 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3989 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3988)
}
__typedArg0 := tmp3988
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3990 := Call(__e, PrimFunc(symshen_4op_1test), tmp3989)


tmp3991 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V816)
}
__typedArg0 := V816
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3992 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3991, Nil)
}
__typedArg0 := tmp3991
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3993 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3990, tmp3992)
}
__typedArg0 := tmp3990
__typedArg1 := tmp3992
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3994 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3993, V814)
}
__typedArg0 := tmp3993
__typedArg1 := V814
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp3995 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3996 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3995)
}
__typedArg0 := tmp3995
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp3997 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp3996)
}
__typedArg0 := tmp3996
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3998 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp3999 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3998)
}
__typedArg0 := tmp3998
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4000 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp3999)
}
__typedArg0 := tmp3999
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4001 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4000)
}
__typedArg0 := tmp4000
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4002 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4003 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4001, tmp4002)
}
__typedArg0 := tmp4001
__typedArg1 := tmp4002
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4004 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp3997, tmp4003)
}
__typedArg0 := tmp3997
__typedArg1 := tmp4003
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4005 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4006 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4005)
}
__typedArg0 := tmp4005
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4007 := Call(__e, PrimFunc(symshen_4op1), tmp4006)


tmp4008 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V816)
}
__typedArg0 := V816
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4009 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4008, Nil)
}
__typedArg0 := tmp4008
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4010 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4007, tmp4009)
}
__typedArg0 := tmp4007
__typedArg1 := tmp4009
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4011 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4012 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4011)
}
__typedArg0 := tmp4011
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4013 := Call(__e, PrimFunc(symshen_4op2), tmp4012)


tmp4014 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V816)
}
__typedArg0 := V816
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4015 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4014, Nil)
}
__typedArg0 := tmp4014
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4016 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4013, tmp4015)
}
__typedArg0 := tmp4013
__typedArg1 := tmp4015
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4017 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V816)
}
__typedArg0 := V816
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4018 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4016, tmp4017)
}
__typedArg0 := tmp4016
__typedArg1 := tmp4017
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4019 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4010, tmp4018)
}
__typedArg0 := tmp4010
__typedArg1 := tmp4018
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4020 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4021 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V816)
}
__typedArg0 := V816
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4022 := Call(__e, PrimFunc(symshen_4beta), tmp4020, tmp4021, V817)


__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp3994, tmp4004, tmp4019, tmp4022)
return


} else {
tmp4035 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4032 Obj

if True == tmp4035 {
tmp4034 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V816)
}
__typedArg0 := V816
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4033 Obj

if True == tmp4034 {
ifres4033 = True


} else {
ifres4033 = False


}

ifres4032 = ifres4033


} else {
ifres4032 = False


}

if True == ifres4032 {
tmp4023 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4024 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V816)
}
__typedArg0 := V816
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4025 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4024, Nil)
}
__typedArg0 := tmp4024
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4026 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4023, tmp4025)
}
__typedArg0 := tmp4023
__typedArg1 := tmp4025
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4027 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(sym_a, tmp4026)
}
__typedArg0 := sym_a
__typedArg1 := tmp4026
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4028 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4027, V814)
}
__typedArg0 := tmp4027
__typedArg1 := V814
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4029 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V815)
}
__typedArg0 := V815
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4030 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V816)
}
__typedArg0 := V816
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp4028, tmp4029, tmp4030, V817)
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

tmp4104 := Call(__e, ns2_1set, symshen_4triple_1stack, tmp3973)


_ = tmp4104

tmp4105 := MakeNative(func(__e *ControlFlow) {
V820 := __e.Get(1)
_ = V820
tmp4124 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V820)
}
__typedArg0 := Nil
__typedArg1 := V820
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4124 {
__e.Return(True)
return
} else {
tmp4122 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V820)
}
__typedArg0 := V820
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4118 Obj

if True == tmp4122 {
tmp4120 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V820)
}
__typedArg0 := V820
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4121 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4120)
}
__typedArg0 := Nil
__typedArg1 := tmp4120
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4119 Obj

if True == tmp4121 {
ifres4119 = True


} else {
ifres4119 = False


}

ifres4118 = ifres4119


} else {
ifres4118 = False


}

if True == ifres4118 {
__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V820)
}
__typedArg0 := V820
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return
} else {
tmp4116 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V820)
}
__typedArg0 := V820
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4112 Obj

if True == tmp4116 {
tmp4114 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V820)
}
__typedArg0 := V820
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4115 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4114)
}
__typedArg0 := tmp4114
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

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
tmp4106 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V820)
}
__typedArg0 := V820
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4107 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V820)
}
__typedArg0 := V820
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4108 := Call(__e, PrimFunc(symshen_4rectify_1test), tmp4107)


tmp4109 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4108, Nil)
}
__typedArg0 := tmp4108
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4110 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4106, tmp4109)
}
__typedArg0 := tmp4106
__typedArg1 := tmp4109
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp4110)
}
__typedArg0 := symand
__typedArg1 := tmp4110
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

tmp4125 := Call(__e, ns2_1set, symshen_4rectify_1test, tmp4105)


_ = tmp4125

tmp4126 := MakeNative(func(__e *ControlFlow) {
V830 := __e.Get(1)
_ = V830
V831 := __e.Get(2)
_ = V831
V832 := __e.Get(3)
_ = V832
tmp4203 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V830, V832)
}
__typedArg0 := V830
__typedArg1 := V832
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4203 {
__e.Return(V831)
return
} else {
tmp4201 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4177 Obj

if True == tmp4201 {
tmp4199 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4200 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlambda, tmp4199)
}
__typedArg0 := symlambda
__typedArg1 := tmp4199
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4179 Obj

if True == tmp4200 {
tmp4197 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4198 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4197)
}
__typedArg0 := tmp4197
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4181 Obj

if True == tmp4198 {
tmp4194 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4195)
}
__typedArg0 := tmp4195
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4183 Obj

if True == tmp4196 {
tmp4190 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4191)
}
__typedArg0 := tmp4191
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4193 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4192)
}
__typedArg0 := Nil
__typedArg1 := tmp4192
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4185 Obj

if True == tmp4193 {
tmp4187 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4188 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4187)
}
__typedArg0 := tmp4187
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4189 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V830, tmp4188)
}
__typedArg0 := V830
__typedArg1 := tmp4188
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4186 Obj

if True == tmp4189 {
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
__e.Return(V832)
return
} else {
tmp4175 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4144 Obj

if True == tmp4175 {
tmp4173 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4174 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp4173)
}
__typedArg0 := symlet
__typedArg1 := tmp4173
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4146 Obj

if True == tmp4174 {
tmp4171 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4172 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4171)
}
__typedArg0 := tmp4171
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4148 Obj

if True == tmp4172 {
tmp4168 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4169 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4168)
}
__typedArg0 := tmp4168
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4170 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4169)
}
__typedArg0 := tmp4169
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4150 Obj

if True == tmp4170 {
tmp4164 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4165 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4164)
}
__typedArg0 := tmp4164
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4166 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4165)
}
__typedArg0 := tmp4165
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4167 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4166)
}
__typedArg0 := tmp4166
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4152 Obj

if True == tmp4167 {
tmp4159 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4160 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4159)
}
__typedArg0 := tmp4159
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4161 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4160)
}
__typedArg0 := tmp4160
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4162 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4161)
}
__typedArg0 := tmp4161
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4163 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4162)
}
__typedArg0 := Nil
__typedArg1 := tmp4162
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4154 Obj

if True == tmp4163 {
tmp4156 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4157 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4156)
}
__typedArg0 := tmp4156
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4158 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V830, tmp4157)
}
__typedArg0 := V830
__typedArg1 := tmp4157
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4155 Obj

if True == tmp4158 {
ifres4155 = True


} else {
ifres4155 = False


}

ifres4154 = ifres4155


} else {
ifres4154 = False


}

var ifres4153 Obj

if True == ifres4154 {
ifres4153 = True


} else {
ifres4153 = False


}

ifres4152 = ifres4153


} else {
ifres4152 = False


}

var ifres4151 Obj

if True == ifres4152 {
ifres4151 = True


} else {
ifres4151 = False


}

ifres4150 = ifres4151


} else {
ifres4150 = False


}

var ifres4149 Obj

if True == ifres4150 {
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

if True == ifres4144 {
tmp4127 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4128 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4127)
}
__typedArg0 := tmp4127
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4129 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4130 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4129)
}
__typedArg0 := tmp4129
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4131 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4132 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4131)
}
__typedArg0 := tmp4131
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4133 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4132)
}
__typedArg0 := tmp4132
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4134 := Call(__e, PrimFunc(symshen_4beta), tmp4130, V831, tmp4133)


tmp4135 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4136 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4135)
}
__typedArg0 := tmp4135
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4137 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4136)
}
__typedArg0 := tmp4136
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4138 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4134, tmp4137)
}
__typedArg0 := tmp4134
__typedArg1 := tmp4137
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4139 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4128, tmp4138)
}
__typedArg0 := tmp4128
__typedArg1 := tmp4138
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp4139)
}
__typedArg0 := symlet
__typedArg1 := tmp4139
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
tmp4142 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V832)
}
__typedArg0 := V832
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp4142 {
tmp4140 := MakeNative(func(__e *ControlFlow) {
Z833 := __e.Get(1)
_ = Z833
__e.TailApply(PrimFunc(symshen_4beta), V830, V831, Z833)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp4140, V832)
return


} else {
__e.Return(V832)
return
}


}


}


}


}, 3)

tmp4204 := Call(__e, ns2_1set, symshen_4beta, tmp4126)


_ = tmp4204

tmp4205 := MakeNative(func(__e *ControlFlow) {
V836 := __e.Get(1)
_ = V836
tmp4213 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, V836)
}
__typedArg0 := symcons
__typedArg1 := V836
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4213 {
__e.Return(symhd)
return
} else {
tmp4211 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8s, V836)
}
__typedArg0 := sym_8s
__typedArg1 := V836
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4211 {
__e.Return(symhdstr)
return
} else {
tmp4209 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8p, V836)
}
__typedArg0 := sym_8p
__typedArg1 := V836
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4209 {
__e.Return(symfst)
return
} else {
tmp4207 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8v, V836)
}
__typedArg0 := sym_8v
__typedArg1 := V836
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4207 {
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

tmp4214 := Call(__e, ns2_1set, symshen_4op1, tmp4205)


_ = tmp4214

tmp4215 := MakeNative(func(__e *ControlFlow) {
V839 := __e.Get(1)
_ = V839
tmp4223 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, V839)
}
__typedArg0 := symcons
__typedArg1 := V839
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4223 {
__e.Return(symtl)
return
} else {
tmp4221 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8s, V839)
}
__typedArg0 := sym_8s
__typedArg1 := V839
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4221 {
__e.Return(symtlstr)
return
} else {
tmp4219 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8p, V839)
}
__typedArg0 := sym_8p
__typedArg1 := V839
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4219 {
__e.Return(symsnd)
return
} else {
tmp4217 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8v, V839)
}
__typedArg0 := sym_8v
__typedArg1 := V839
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4217 {
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

tmp4224 := Call(__e, ns2_1set, symshen_4op2, tmp4215)


_ = tmp4224

tmp4225 := MakeNative(func(__e *ControlFlow) {
V842 := __e.Get(1)
_ = V842
tmp4233 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons, V842)
}
__typedArg0 := symcons
__typedArg1 := V842
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4233 {
__e.Return(symcons_2)
return
} else {
tmp4231 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8s, V842)
}
__typedArg0 := sym_8s
__typedArg1 := V842
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4231 {
__e.Return(symshen_4_7string_2)
return
} else {
tmp4229 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8p, V842)
}
__typedArg0 := sym_8p
__typedArg1 := V842
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4229 {
__e.Return(symtuple_2)
return
} else {
tmp4227 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_8v, V842)
}
__typedArg0 := sym_8v
__typedArg1 := V842
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4227 {
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

tmp4234 := Call(__e, ns2_1set, symshen_4op_1test, tmp4225)


_ = tmp4234

tmp4235 := MakeNative(func(__e *ControlFlow) {
V843 := __e.Get(1)
_ = V843
tmp4237 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(MakeString(""), V843)
}
__typedArg0 := MakeString("")
__typedArg1 := V843
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4237 {
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

tmp4238 := Call(__e, ns2_1set, symshen_4_7string_2, tmp4235)


_ = tmp4238

tmp4239 := MakeNative(func(__e *ControlFlow) {
V844 := __e.Get(1)
_ = V844
tmp4241 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp4242 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V844, tmp4241)
}
__typedArg0 := V844
__typedArg1 := tmp4241
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4242 {
__e.Return(False)
return
} else {
__e.TailApply(PrimFunc(symvector_2), V844)
return
}


}, 1)

tmp4243 := Call(__e, ns2_1set, symshen_4_7vector_2, tmp4239)


_ = tmp4243

tmp4244 := MakeNative(func(__e *ControlFlow) {
V847 := __e.Get(1)
_ = V847
tmp4248 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_7, V847)
}
__typedArg0 := sym_7
__typedArg1 := V847
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4248 {
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
tmp4246 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(sym_1, V847)
}
__typedArg0 := sym_1
__typedArg1 := V847
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4246 {
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

tmp4249 := Call(__e, ns2_1set, symfactorise, tmp4244)


_ = tmp4249

tmp4250 := MakeNative(func(__e *ControlFlow) {
V848 := __e.Get(1)
_ = V848
tmp4252 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symvalue) {
return PrimValue(symshen_4_dfactorise_2_d)
}
__typedArg0 := symshen_4_dfactorise_2_d
return Call(__e, PrimFunc(symvalue), __typedArg0)
})()

if True == tmp4252 {
__e.TailApply(PrimFunc(symshen_4factor), V848)
return
} else {
__e.Return(V848)
return
}


}, 1)

tmp4253 := Call(__e, ns2_1set, symshen_4factorise_1code, tmp4250)


_ = tmp4253

tmp4254 := MakeNative(func(__e *ControlFlow) {
V849 := __e.Get(1)
_ = V849
tmp4311 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V849)
}
__typedArg0 := V849
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4270 Obj

if True == tmp4311 {
tmp4309 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V849)
}
__typedArg0 := V849
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4310 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symdefun, tmp4309)
}
__typedArg0 := symdefun
__typedArg1 := tmp4309
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4272 Obj

if True == tmp4310 {
tmp4307 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V849)
}
__typedArg0 := V849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4308 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4307)
}
__typedArg0 := tmp4307
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4274 Obj

if True == tmp4308 {
tmp4304 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V849)
}
__typedArg0 := V849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4305 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4304)
}
__typedArg0 := tmp4304
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4306 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4305)
}
__typedArg0 := tmp4305
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4276 Obj

if True == tmp4306 {
tmp4300 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V849)
}
__typedArg0 := V849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4301 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4300)
}
__typedArg0 := tmp4300
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4302 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4301)
}
__typedArg0 := tmp4301
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4303 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4302)
}
__typedArg0 := tmp4302
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4278 Obj

if True == tmp4303 {
tmp4295 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V849)
}
__typedArg0 := V849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4296 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4295)
}
__typedArg0 := tmp4295
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4297 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4296)
}
__typedArg0 := tmp4296
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4298 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4297)
}
__typedArg0 := tmp4297
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4299 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4298)
}
__typedArg0 := tmp4298
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4280 Obj

if True == tmp4299 {
tmp4289 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V849)
}
__typedArg0 := V849
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4291)
}
__typedArg0 := tmp4291
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4293 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4292)
}
__typedArg0 := tmp4292
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4294 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcond, tmp4293)
}
__typedArg0 := symcond
__typedArg1 := tmp4293
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4282 Obj

if True == tmp4294 {
tmp4284 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V849)
}
__typedArg0 := V849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4285 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4284)
}
__typedArg0 := tmp4284
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4286)
}
__typedArg0 := tmp4286
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4288 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4287)
}
__typedArg0 := Nil
__typedArg1 := tmp4287
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4283 Obj

if True == tmp4288 {
ifres4283 = True


} else {
ifres4283 = False


}

ifres4282 = ifres4283


} else {
ifres4282 = False


}

var ifres4281 Obj

if True == ifres4282 {
ifres4281 = True


} else {
ifres4281 = False


}

ifres4280 = ifres4281


} else {
ifres4280 = False


}

var ifres4279 Obj

if True == ifres4280 {
ifres4279 = True


} else {
ifres4279 = False


}

ifres4278 = ifres4279


} else {
ifres4278 = False


}

var ifres4277 Obj

if True == ifres4278 {
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

if True == ifres4270 {
tmp4255 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V849)
}
__typedArg0 := V849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4256 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4255)
}
__typedArg0 := tmp4255
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4257 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V849)
}
__typedArg0 := V849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4258 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4257)
}
__typedArg0 := tmp4257
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4259 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4258)
}
__typedArg0 := tmp4258
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4260 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V849)
}
__typedArg0 := V849
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4261 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4260)
}
__typedArg0 := tmp4260
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4262 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4261)
}
__typedArg0 := tmp4261
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4263 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4262)
}
__typedArg0 := tmp4262
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4264 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4263)
}
__typedArg0 := tmp4263
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4265 := Call(__e, PrimFunc(symshen_4factor_1recognisors), tmp4264)


tmp4266 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4265, Nil)
}
__typedArg0 := tmp4265
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4267 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4259, tmp4266)
}
__typedArg0 := tmp4259
__typedArg1 := tmp4266
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4268 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4256, tmp4267)
}
__typedArg0 := tmp4256
__typedArg1 := tmp4267
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symdefun, tmp4268)
}
__typedArg0 := symdefun
__typedArg1 := tmp4268
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


} else {
__e.Return(V849)
return
}


}, 1)

tmp4312 := Call(__e, ns2_1set, symshen_4factor, tmp4254)


_ = tmp4312

tmp4313 := MakeNative(func(__e *ControlFlow) {
V852 := __e.Get(1)
_ = V852
tmp4470 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4450 Obj

if True == tmp4470 {
tmp4468 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4469 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4468)
}
__typedArg0 := tmp4468
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4452 Obj

if True == tmp4469 {
tmp4465 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4466 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4465)
}
__typedArg0 := tmp4465
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4467 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(True, tmp4466)
}
__typedArg0 := True
__typedArg1 := tmp4466
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4454 Obj

if True == tmp4467 {
tmp4462 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4463 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4462)
}
__typedArg0 := tmp4462
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4464 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4463)
}
__typedArg0 := tmp4463
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4456 Obj

if True == tmp4464 {
tmp4458 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4459 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4458)
}
__typedArg0 := tmp4458
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4460 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4459)
}
__typedArg0 := tmp4459
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4461 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4460)
}
__typedArg0 := Nil
__typedArg1 := tmp4460
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4457 Obj

if True == tmp4461 {
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

if True == ifres4450 {
tmp4314 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4315 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4314)
}
__typedArg0 := tmp4314
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4315)
}
__typedArg0 := tmp4315
return Call(__e, PrimFunc(symhd), __typedArg0)
})())
return


} else {
tmp4448 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4401 Obj

if True == tmp4448 {
tmp4446 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4447 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4446)
}
__typedArg0 := tmp4446
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4403 Obj

if True == tmp4447 {
tmp4443 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4444 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4443)
}
__typedArg0 := tmp4443
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4445 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4444)
}
__typedArg0 := tmp4444
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4405 Obj

if True == tmp4445 {
tmp4439 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4440 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4439)
}
__typedArg0 := tmp4439
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4441 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4440)
}
__typedArg0 := tmp4440
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4442 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symand, tmp4441)
}
__typedArg0 := symand
__typedArg1 := tmp4441
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4407 Obj

if True == tmp4442 {
tmp4435 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4436 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4435)
}
__typedArg0 := tmp4435
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4437 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4436)
}
__typedArg0 := tmp4436
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4438 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4437)
}
__typedArg0 := tmp4437
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4409 Obj

if True == tmp4438 {
tmp4430 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4431 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4430)
}
__typedArg0 := tmp4430
return Call(__e, PrimFunc(symhd), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4433)
}
__typedArg0 := tmp4433
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4411 Obj

if True == tmp4434 {
tmp4424 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4425 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4424)
}
__typedArg0 := tmp4424
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4426 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4425)
}
__typedArg0 := tmp4425
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4427 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4426)
}
__typedArg0 := tmp4426
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4428 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4427)
}
__typedArg0 := tmp4427
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4429 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4428)
}
__typedArg0 := Nil
__typedArg1 := tmp4428
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4413 Obj

if True == tmp4429 {
tmp4421 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4422 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4421)
}
__typedArg0 := tmp4421
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4423 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4422)
}
__typedArg0 := tmp4422
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4415 Obj

if True == tmp4423 {
tmp4417 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4418 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4417)
}
__typedArg0 := tmp4417
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4419 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4418)
}
__typedArg0 := tmp4418
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4420 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4419)
}
__typedArg0 := Nil
__typedArg1 := tmp4419
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4416 Obj

if True == tmp4420 {
ifres4416 = True


} else {
ifres4416 = False


}

ifres4415 = ifres4416


} else {
ifres4415 = False


}

var ifres4414 Obj

if True == ifres4415 {
ifres4414 = True


} else {
ifres4414 = False


}

ifres4413 = ifres4414


} else {
ifres4413 = False


}

var ifres4412 Obj

if True == ifres4413 {
ifres4412 = True


} else {
ifres4412 = False


}

ifres4411 = ifres4412


} else {
ifres4411 = False


}

var ifres4410 Obj

if True == ifres4411 {
ifres4410 = True


} else {
ifres4410 = False


}

ifres4409 = ifres4410


} else {
ifres4409 = False


}

var ifres4408 Obj

if True == ifres4409 {
ifres4408 = True


} else {
ifres4408 = False


}

ifres4407 = ifres4408


} else {
ifres4407 = False


}

var ifres4406 Obj

if True == ifres4407 {
ifres4406 = True


} else {
ifres4406 = False


}

ifres4405 = ifres4406


} else {
ifres4405 = False


}

var ifres4404 Obj

if True == ifres4405 {
ifres4404 = True


} else {
ifres4404 = False


}

ifres4403 = ifres4404


} else {
ifres4403 = False


}

var ifres4402 Obj

if True == ifres4403 {
ifres4402 = True


} else {
ifres4402 = False


}

ifres4401 = ifres4402


} else {
ifres4401 = False


}

if True == ifres4401 {
tmp4316 := MakeNative(func(__e *ControlFlow) {
W853 := __e.Get(1)
_ = W853
tmp4317 := MakeNative(func(__e *ControlFlow) {
W854 := __e.Get(1)
_ = W854
tmp4366 := Call(__e, PrimFunc(symshen_4bad_1pivot_2), W854)


if True == tmp4366 {
tmp4318 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, Nil)
}
__typedArg0 := symif
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4319 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4320 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4319)
}
__typedArg0 := tmp4319
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4321 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4322 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4321)
}
__typedArg0 := tmp4321
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4323 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4322)
}
__typedArg0 := tmp4322
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4324 := Call(__e, PrimFunc(symshen_4recursively_1factor_1selectors), tmp4320, tmp4323)


tmp4325 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4326 := Call(__e, PrimFunc(symshen_4factor_1recognisors), tmp4325)


tmp4327 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4326, Nil)
}
__typedArg0 := tmp4326
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4328 := Call(__e, PrimFunc(symappend), tmp4324, tmp4327)


__e.TailApply(PrimFunc(symappend), tmp4318, tmp4328)
return


} else {
tmp4329 := MakeNative(func(__e *ControlFlow) {
W855 := __e.Get(1)
_ = W855
tmp4330 := MakeNative(func(__e *ControlFlow) {
W856 := __e.Get(1)
_ = W856
tmp4331 := MakeNative(func(__e *ControlFlow) {
W857 := __e.Get(1)
_ = W857
tmp4332 := MakeNative(func(__e *ControlFlow) {
W858 := __e.Get(1)
_ = W858
tmp4333 := MakeNative(func(__e *ControlFlow) {
W859 := __e.Get(1)
_ = W859
__e.TailApply(PrimFunc(symshen_4remove_1indirection), W859)
return
}, 1)

tmp4334 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W856, Nil)
}
__typedArg0 := W856
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4335 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symfreeze, tmp4334)
}
__typedArg0 := symfreeze
__typedArg1 := tmp4334
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4336 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4337 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4336)
}
__typedArg0 := tmp4336
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4338 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4337)
}
__typedArg0 := tmp4337
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4339 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4338)
}
__typedArg0 := tmp4338
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4340 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4341 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4340)
}
__typedArg0 := tmp4340
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4342 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4341)
}
__typedArg0 := tmp4341
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4343 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4342)
}
__typedArg0 := tmp4342
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4344 := Call(__e, PrimFunc(symshen_4factor_1recognisors), W858)


tmp4345 := Call(__e, PrimFunc(symshen_4factor_1selectors), tmp4343, tmp4344)


tmp4346 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W857, Nil)
}
__typedArg0 := W857
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4347 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp4346)
}
__typedArg0 := symthaw
__typedArg1 := tmp4346
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4348 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4347, Nil)
}
__typedArg0 := tmp4347
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4349 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4345, tmp4348)
}
__typedArg0 := tmp4345
__typedArg1 := tmp4348
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4350 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4339, tmp4349)
}
__typedArg0 := tmp4339
__typedArg1 := tmp4349
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4351 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp4350)
}
__typedArg0 := symif
__typedArg1 := tmp4350
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4352 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4351, Nil)
}
__typedArg0 := tmp4351
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4353 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4335, tmp4352)
}
__typedArg0 := tmp4335
__typedArg1 := tmp4352
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4354 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W857, tmp4353)
}
__typedArg0 := W857
__typedArg1 := tmp4353
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4355 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp4354)
}
__typedArg0 := symlet
__typedArg1 := tmp4354
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(tmp4333, tmp4355)
return


}, 1)

tmp4356 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W857, Nil)
}
__typedArg0 := W857
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4357 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symthaw, tmp4356)
}
__typedArg0 := symthaw
__typedArg1 := tmp4356
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4358 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4357, Nil)
}
__typedArg0 := tmp4357
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4359 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp4358)
}
__typedArg0 := True
__typedArg1 := tmp4358
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4360 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4359, W854)
}
__typedArg0 := tmp4359
__typedArg1 := W854
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4361 := Call(__e, PrimFunc(symreverse), tmp4360)


__e.TailApply(tmp4332, tmp4361)
return


}, 1)

tmp4362 := Call(__e, PrimFunc(symgensym), symGoTo)


__e.TailApply(tmp4331, tmp4362)
return


}, 1)

tmp4363 := Call(__e, PrimFunc(symshen_4factor_1recognisors), W855)


__e.TailApply(tmp4330, tmp4363)
return


}, 1)

tmp4364 := Call(__e, PrimFunc(symsnd), W853)


__e.TailApply(tmp4329, tmp4364)
return


}


}, 1)

tmp4367 := Call(__e, PrimFunc(symfst), W853)


__e.TailApply(tmp4317, tmp4367)
return


}, 1)

tmp4368 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4369 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4368)
}
__typedArg0 := tmp4368
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4370 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4369)
}
__typedArg0 := tmp4369
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4371 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4370)
}
__typedArg0 := tmp4370
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4372 := Call(__e, PrimFunc(symshen_4pivot_1on), tmp4371, V852, Nil)


__e.TailApply(tmp4316, tmp4372)
return


} else {
tmp4399 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4384 Obj

if True == tmp4399 {
tmp4397 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4398 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4397)
}
__typedArg0 := tmp4397
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4386 Obj

if True == tmp4398 {
tmp4394 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4395 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4394)
}
__typedArg0 := tmp4394
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4396 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4395)
}
__typedArg0 := tmp4395
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4388 Obj

if True == tmp4396 {
tmp4390 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4391 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4390)
}
__typedArg0 := tmp4390
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4392 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4391)
}
__typedArg0 := tmp4391
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4393 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4392)
}
__typedArg0 := Nil
__typedArg1 := tmp4392
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4389 Obj

if True == tmp4393 {
ifres4389 = True


} else {
ifres4389 = False


}

ifres4388 = ifres4389


} else {
ifres4388 = False


}

var ifres4387 Obj

if True == ifres4388 {
ifres4387 = True


} else {
ifres4387 = False


}

ifres4386 = ifres4387


} else {
ifres4386 = False


}

var ifres4385 Obj

if True == ifres4386 {
ifres4385 = True


} else {
ifres4385 = False


}

ifres4384 = ifres4385


} else {
ifres4384 = False


}

if True == ifres4384 {
tmp4373 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4374 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4373)
}
__typedArg0 := tmp4373
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4375 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4376 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4375)
}
__typedArg0 := tmp4375
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4377 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4376)
}
__typedArg0 := tmp4376
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4378 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V852)
}
__typedArg0 := V852
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4379 := Call(__e, PrimFunc(symshen_4factor_1recognisors), tmp4378)


tmp4380 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4379, Nil)
}
__typedArg0 := tmp4379
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4381 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4377, tmp4380)
}
__typedArg0 := tmp4377
__typedArg1 := tmp4380
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4382 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4374, tmp4381)
}
__typedArg0 := tmp4374
__typedArg1 := tmp4381
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symif, tmp4382)
}
__typedArg0 := symif
__typedArg1 := tmp4382
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

tmp4471 := Call(__e, ns2_1set, symshen_4factor_1recognisors, tmp4313)


_ = tmp4471

tmp4472 := MakeNative(func(__e *ControlFlow) {
V860 := __e.Get(1)
_ = V860
V861 := __e.Get(2)
_ = V861
tmp4544 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4518 Obj

if True == tmp4544 {
tmp4542 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4543 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp4542)
}
__typedArg0 := symlet
__typedArg1 := tmp4542
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4520 Obj

if True == tmp4543 {
tmp4540 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4541 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4540)
}
__typedArg0 := tmp4540
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4522 Obj

if True == tmp4541 {
tmp4537 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4538 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4537)
}
__typedArg0 := tmp4537
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4539 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4538)
}
__typedArg0 := tmp4538
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4524 Obj

if True == tmp4539 {
tmp4533 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4534 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4533)
}
__typedArg0 := tmp4533
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4535)
}
__typedArg0 := tmp4535
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4526 Obj

if True == tmp4536 {
tmp4528 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4529 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4528)
}
__typedArg0 := tmp4528
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4530 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4529)
}
__typedArg0 := tmp4529
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4531 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4530)
}
__typedArg0 := tmp4530
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4532 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4531)
}
__typedArg0 := Nil
__typedArg1 := tmp4531
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4527 Obj

if True == tmp4532 {
ifres4527 = True


} else {
ifres4527 = False


}

ifres4526 = ifres4527


} else {
ifres4526 = False


}

var ifres4525 Obj

if True == ifres4526 {
ifres4525 = True


} else {
ifres4525 = False


}

ifres4524 = ifres4525


} else {
ifres4524 = False


}

var ifres4523 Obj

if True == ifres4524 {
ifres4523 = True


} else {
ifres4523 = False


}

ifres4522 = ifres4523


} else {
ifres4522 = False


}

var ifres4521 Obj

if True == ifres4522 {
ifres4521 = True


} else {
ifres4521 = False


}

ifres4520 = ifres4521


} else {
ifres4520 = False


}

var ifres4519 Obj

if True == ifres4520 {
ifres4519 = True


} else {
ifres4519 = False


}

ifres4518 = ifres4519


} else {
ifres4518 = False


}

if True == ifres4518 {
tmp4473 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4474 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4473)
}
__typedArg0 := tmp4473
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4475 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4476)
}
__typedArg0 := tmp4476
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4478 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4480)
}
__typedArg0 := tmp4480
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4482 := Call(__e, PrimFunc(symshen_4recursively_1factor_1selectors), tmp4481, V861)


__e.TailApply(PrimFunc(symshen_4restore_1local), tmp4474, tmp4477, tmp4482)
return


} else {
tmp4516 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4497 Obj

if True == tmp4516 {
tmp4514 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4515 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symand, tmp4514)
}
__typedArg0 := symand
__typedArg1 := tmp4514
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4499 Obj

if True == tmp4515 {
tmp4512 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4513 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4512)
}
__typedArg0 := tmp4512
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4501 Obj

if True == tmp4513 {
tmp4509 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
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
tmp4483 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4484 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4483)
}
__typedArg0 := tmp4483
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4485 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4486 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4485)
}
__typedArg0 := tmp4485
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4487 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4488 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4487)
}
__typedArg0 := tmp4487
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4489 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4488)
}
__typedArg0 := tmp4488
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4490 := Call(__e, PrimFunc(symshen_4factor_1selectors), tmp4486, tmp4489)


tmp4491 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V860)
}
__typedArg0 := V860
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4492 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4491)
}
__typedArg0 := tmp4491
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4493 := Call(__e, PrimFunc(symshen_4factor_1selectors), tmp4492, V861)


tmp4494 := Call(__e, PrimFunc(symshen_4recursively_1factor_1selectors), tmp4490, tmp4493)


__e.TailApply(PrimFunc(symshen_4restore_1P), tmp4484, tmp4494)
return


} else {
tmp4495 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V861, Nil)
}
__typedArg0 := V861
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V860, tmp4495)
}
__typedArg0 := V860
__typedArg1 := tmp4495
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}


}


}, 2)

tmp4545 := Call(__e, ns2_1set, symshen_4recursively_1factor_1selectors, tmp4472)


_ = tmp4545

tmp4546 := MakeNative(func(__e *ControlFlow) {
V862 := __e.Get(1)
_ = V862
V863 := __e.Get(2)
_ = V863
tmp4562 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V863)
}
__typedArg0 := V863
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4553 Obj

if True == tmp4562 {
tmp4560 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V863)
}
__typedArg0 := V863
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4561 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4560)
}
__typedArg0 := tmp4560
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4555 Obj

if True == tmp4561 {
tmp4557 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V863)
}
__typedArg0 := V863
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4558 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4557)
}
__typedArg0 := tmp4557
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4559 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4558)
}
__typedArg0 := Nil
__typedArg1 := tmp4558
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4556 Obj

if True == tmp4559 {
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

if True == ifres4553 {
tmp4547 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V863)
}
__typedArg0 := V863
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4548 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4547, Nil)
}
__typedArg0 := tmp4547
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4549 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V862, tmp4548)
}
__typedArg0 := V862
__typedArg1 := tmp4548
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4550 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symand, tmp4549)
}
__typedArg0 := symand
__typedArg1 := tmp4549
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4551 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V863)
}
__typedArg0 := V863
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4550, tmp4551)
}
__typedArg0 := tmp4550
__typedArg1 := tmp4551
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

tmp4563 := Call(__e, ns2_1set, symshen_4restore_1P, tmp4546)


_ = tmp4563

tmp4564 := MakeNative(func(__e *ControlFlow) {
V864 := __e.Get(1)
_ = V864
V865 := __e.Get(2)
_ = V865
V866 := __e.Get(3)
_ = V866
tmp4581 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V866)
}
__typedArg0 := V866
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4572 Obj

if True == tmp4581 {
tmp4579 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V866)
}
__typedArg0 := V866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4580 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4579)
}
__typedArg0 := tmp4579
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4574 Obj

if True == tmp4580 {
tmp4576 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V866)
}
__typedArg0 := V866
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4577)
}
__typedArg0 := Nil
__typedArg1 := tmp4577
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4575 Obj

if True == tmp4578 {
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

if True == ifres4572 {
tmp4565 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V866)
}
__typedArg0 := V866
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4566 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4565, Nil)
}
__typedArg0 := tmp4565
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4567 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V865, tmp4566)
}
__typedArg0 := V865
__typedArg1 := tmp4566
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4568 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(V864, tmp4567)
}
__typedArg0 := V864
__typedArg1 := tmp4567
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4569 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp4568)
}
__typedArg0 := symlet
__typedArg1 := tmp4568
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4570 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V866)
}
__typedArg0 := V866
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4569, tmp4570)
}
__typedArg0 := tmp4569
__typedArg1 := tmp4570
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

tmp4582 := Call(__e, ns2_1set, symshen_4restore_1local, tmp4564)


_ = tmp4582

tmp4583 := MakeNative(func(__e *ControlFlow) {
V871 := __e.Get(1)
_ = V871
tmp4589 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V871)
}
__typedArg0 := V871
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4585 Obj

if True == tmp4589 {
tmp4587 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V871)
}
__typedArg0 := V871
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4588 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4587)
}
__typedArg0 := Nil
__typedArg1 := tmp4587
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4586 Obj

if True == tmp4588 {
ifres4586 = True


} else {
ifres4586 = False


}

ifres4585 = ifres4586


} else {
ifres4585 = False


}

if True == ifres4585 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp4590 := Call(__e, ns2_1set, symshen_4bad_1pivot_2, tmp4583)


_ = tmp4590

tmp4591 := MakeNative(func(__e *ControlFlow) {
V872 := __e.Get(1)
_ = V872
tmp4706 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4606 Obj

if True == tmp4706 {
tmp4704 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4705 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symlet, tmp4704)
}
__typedArg0 := symlet
__typedArg1 := tmp4704
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4608 Obj

if True == tmp4705 {
tmp4702 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4703 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4702)
}
__typedArg0 := tmp4702
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4610 Obj

if True == tmp4703 {
tmp4699 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4700 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4699)
}
__typedArg0 := tmp4699
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4701 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4700)
}
__typedArg0 := tmp4700
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4612 Obj

if True == tmp4701 {
tmp4695 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4696)
}
__typedArg0 := tmp4696
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4698 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4697)
}
__typedArg0 := tmp4697
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4614 Obj

if True == tmp4698 {
tmp4690 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4691)
}
__typedArg0 := tmp4691
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4693 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4692)
}
__typedArg0 := tmp4692
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symfreeze, tmp4693)
}
__typedArg0 := symfreeze
__typedArg1 := tmp4693
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4616 Obj

if True == tmp4694 {
tmp4685 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4687)
}
__typedArg0 := tmp4687
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4689 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4688)
}
__typedArg0 := tmp4688
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4618 Obj

if True == tmp4689 {
tmp4679 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4680 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4679)
}
__typedArg0 := tmp4679
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4681 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4680)
}
__typedArg0 := tmp4680
return Call(__e, PrimFunc(symhd), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4683)
}
__typedArg0 := tmp4683
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4620 Obj

if True == tmp4684 {
tmp4672 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4673)
}
__typedArg0 := tmp4673
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4675 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4674)
}
__typedArg0 := tmp4674
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4676 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4675)
}
__typedArg0 := tmp4675
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4677 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4676)
}
__typedArg0 := tmp4676
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4678 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symthaw, tmp4677)
}
__typedArg0 := symthaw
__typedArg1 := tmp4677
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4622 Obj

if True == tmp4678 {
tmp4665 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4666 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4665)
}
__typedArg0 := tmp4665
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4667 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4666)
}
__typedArg0 := tmp4666
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4668 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4667)
}
__typedArg0 := tmp4667
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4669 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4668)
}
__typedArg0 := tmp4668
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4670 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4669)
}
__typedArg0 := tmp4669
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4671 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4670)
}
__typedArg0 := tmp4670
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4624 Obj

if True == tmp4671 {
tmp4657 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4658 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4657)
}
__typedArg0 := tmp4657
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4659 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4658)
}
__typedArg0 := tmp4658
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4660 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4659)
}
__typedArg0 := tmp4659
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4661 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4660)
}
__typedArg0 := tmp4660
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4662 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4661)
}
__typedArg0 := tmp4661
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4663 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4662)
}
__typedArg0 := tmp4662
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4664 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4663)
}
__typedArg0 := Nil
__typedArg1 := tmp4663
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4626 Obj

if True == tmp4664 {
tmp4651 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4652 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4651)
}
__typedArg0 := tmp4651
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
return PrimTail(tmp4653)
}
__typedArg0 := tmp4653
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4655)
}
__typedArg0 := Nil
__typedArg1 := tmp4655
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4628 Obj

if True == tmp4656 {
tmp4647 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4648 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4647)
}
__typedArg0 := tmp4647
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4649)
}
__typedArg0 := tmp4649
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4630 Obj

if True == tmp4650 {
tmp4642 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4643)
}
__typedArg0 := tmp4643
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4645 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4644)
}
__typedArg0 := tmp4644
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4646 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4645)
}
__typedArg0 := Nil
__typedArg1 := tmp4645
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4632 Obj

if True == tmp4646 {
tmp4634 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
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

tmp4637 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4636)
}
__typedArg0 := tmp4636
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4638 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4637)
}
__typedArg0 := tmp4637
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4639 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4638)
}
__typedArg0 := tmp4638
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4640 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4639)
}
__typedArg0 := tmp4639
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4641 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsymbol_2) {
return PrimIsSymbol(tmp4640)
}
__typedArg0 := tmp4640
return Call(__e, PrimFunc(symsymbol_2), __typedArg0)
})()

var ifres4633 Obj

if True == tmp4641 {
ifres4633 = True


} else {
ifres4633 = False


}

ifres4632 = ifres4633


} else {
ifres4632 = False


}

var ifres4631 Obj

if True == ifres4632 {
ifres4631 = True


} else {
ifres4631 = False


}

ifres4630 = ifres4631


} else {
ifres4630 = False


}

var ifres4629 Obj

if True == ifres4630 {
ifres4629 = True


} else {
ifres4629 = False


}

ifres4628 = ifres4629


} else {
ifres4628 = False


}

var ifres4627 Obj

if True == ifres4628 {
ifres4627 = True


} else {
ifres4627 = False


}

ifres4626 = ifres4627


} else {
ifres4626 = False


}

var ifres4625 Obj

if True == ifres4626 {
ifres4625 = True


} else {
ifres4625 = False


}

ifres4624 = ifres4625


} else {
ifres4624 = False


}

var ifres4623 Obj

if True == ifres4624 {
ifres4623 = True


} else {
ifres4623 = False


}

ifres4622 = ifres4623


} else {
ifres4622 = False


}

var ifres4621 Obj

if True == ifres4622 {
ifres4621 = True


} else {
ifres4621 = False


}

ifres4620 = ifres4621


} else {
ifres4620 = False


}

var ifres4619 Obj

if True == ifres4620 {
ifres4619 = True


} else {
ifres4619 = False


}

ifres4618 = ifres4619


} else {
ifres4618 = False


}

var ifres4617 Obj

if True == ifres4618 {
ifres4617 = True


} else {
ifres4617 = False


}

ifres4616 = ifres4617


} else {
ifres4616 = False


}

var ifres4615 Obj

if True == ifres4616 {
ifres4615 = True


} else {
ifres4615 = False


}

ifres4614 = ifres4615


} else {
ifres4614 = False


}

var ifres4613 Obj

if True == ifres4614 {
ifres4613 = True


} else {
ifres4613 = False


}

ifres4612 = ifres4613


} else {
ifres4612 = False


}

var ifres4611 Obj

if True == ifres4612 {
ifres4611 = True


} else {
ifres4611 = False


}

ifres4610 = ifres4611


} else {
ifres4610 = False


}

var ifres4609 Obj

if True == ifres4610 {
ifres4609 = True


} else {
ifres4609 = False


}

ifres4608 = ifres4609


} else {
ifres4608 = False


}

var ifres4607 Obj

if True == ifres4608 {
ifres4607 = True


} else {
ifres4607 = False


}

ifres4606 = ifres4607


} else {
ifres4606 = False


}

if True == ifres4606 {
tmp4592 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4593 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4592)
}
__typedArg0 := tmp4592
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4594 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4593)
}
__typedArg0 := tmp4593
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4595 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4594)
}
__typedArg0 := tmp4594
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4596 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4595)
}
__typedArg0 := tmp4595
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4597 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4596)
}
__typedArg0 := tmp4596
return Call(__e, PrimFunc(symtl), __typedArg0)
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
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4600 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4599)
}
__typedArg0 := tmp4599
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4601 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V872)
}
__typedArg0 := V872
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4602 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4601)
}
__typedArg0 := tmp4601
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4603 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4602)
}
__typedArg0 := tmp4602
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4604 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4603)
}
__typedArg0 := tmp4603
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

__e.TailApply(PrimFunc(symsubst), tmp4598, tmp4600, tmp4604)
return


} else {
__e.Return(V872)
return
}


}, 1)

tmp4707 := Call(__e, ns2_1set, symshen_4remove_1indirection, tmp4591)


_ = tmp4707

tmp4708 := MakeNative(func(__e *ControlFlow) {
V875 := __e.Get(1)
_ = V875
V876 := __e.Get(2)
_ = V876
V877 := __e.Get(3)
_ = V877
tmp4807 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4753 Obj

if True == tmp4807 {
tmp4805 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4806 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4805)
}
__typedArg0 := tmp4805
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4755 Obj

if True == tmp4806 {
tmp4802 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4803 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4802)
}
__typedArg0 := tmp4802
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4804 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4803)
}
__typedArg0 := tmp4803
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4757 Obj

if True == tmp4804 {
tmp4798 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4799 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4798)
}
__typedArg0 := tmp4798
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4800 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4799)
}
__typedArg0 := tmp4799
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4801 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symand, tmp4800)
}
__typedArg0 := symand
__typedArg1 := tmp4800
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4759 Obj

if True == tmp4801 {
tmp4794 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4795 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4794)
}
__typedArg0 := tmp4794
return Call(__e, PrimFunc(symhd), __typedArg0)
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

var ifres4761 Obj

if True == tmp4797 {
tmp4789 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4790 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4789)
}
__typedArg0 := tmp4789
return Call(__e, PrimFunc(symhd), __typedArg0)
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4792)
}
__typedArg0 := tmp4792
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4763 Obj

if True == tmp4793 {
tmp4783 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4784 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4783)
}
__typedArg0 := tmp4783
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4785 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4784)
}
__typedArg0 := tmp4784
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4786 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4785)
}
__typedArg0 := tmp4785
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4787 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4786)
}
__typedArg0 := tmp4786
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4788 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4787)
}
__typedArg0 := Nil
__typedArg1 := tmp4787
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4765 Obj

if True == tmp4788 {
tmp4780 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4781 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4780)
}
__typedArg0 := tmp4780
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4782 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4781)
}
__typedArg0 := tmp4781
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4767 Obj

if True == tmp4782 {
tmp4776 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4777 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4776)
}
__typedArg0 := tmp4776
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4778 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4777)
}
__typedArg0 := tmp4777
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4779 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4778)
}
__typedArg0 := Nil
__typedArg1 := tmp4778
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4769 Obj

if True == tmp4779 {
tmp4771 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4772 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4771)
}
__typedArg0 := tmp4771
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4773 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4772)
}
__typedArg0 := tmp4772
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4774 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4773)
}
__typedArg0 := tmp4773
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4775 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(V875, tmp4774)
}
__typedArg0 := V875
__typedArg1 := tmp4774
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4770 Obj

if True == tmp4775 {
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
tmp4709 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4710 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4709)
}
__typedArg0 := tmp4709
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4711 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4710)
}
__typedArg0 := tmp4710
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4712 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4711)
}
__typedArg0 := tmp4711
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4713 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4714 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4715 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4714)
}
__typedArg0 := tmp4714
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4716 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4715)
}
__typedArg0 := tmp4715
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4717 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4716)
}
__typedArg0 := tmp4716
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4718 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(tmp4717)
}
__typedArg0 := tmp4717
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4719 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4720 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4719)
}
__typedArg0 := tmp4719
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4721 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4718, tmp4720)
}
__typedArg0 := tmp4718
__typedArg1 := tmp4720
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4722 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4721, V877)
}
__typedArg0 := tmp4721
__typedArg1 := V877
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4pivot_1on), tmp4712, tmp4713, tmp4722)
return


} else {
tmp4751 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4731 Obj

if True == tmp4751 {
tmp4749 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4750 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4749)
}
__typedArg0 := tmp4749
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4733 Obj

if True == tmp4750 {
tmp4746 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
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

var ifres4735 Obj

if True == tmp4748 {
tmp4742 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4743 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4742)
}
__typedArg0 := tmp4742
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

var ifres4737 Obj

if True == tmp4745 {
tmp4739 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
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
return PrimEqual(V875, tmp4740)
}
__typedArg0 := V875
__typedArg1 := tmp4740
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4738 Obj

if True == tmp4741 {
ifres4738 = True


} else {
ifres4738 = False


}

ifres4737 = ifres4738


} else {
ifres4737 = False


}

var ifres4736 Obj

if True == ifres4737 {
ifres4736 = True


} else {
ifres4736 = False


}

ifres4735 = ifres4736


} else {
ifres4735 = False


}

var ifres4734 Obj

if True == ifres4735 {
ifres4734 = True


} else {
ifres4734 = False


}

ifres4733 = ifres4734


} else {
ifres4733 = False


}

var ifres4732 Obj

if True == ifres4733 {
ifres4732 = True


} else {
ifres4732 = False


}

ifres4731 = ifres4732


} else {
ifres4731 = False


}

if True == ifres4731 {
tmp4723 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
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
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4726 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V876)
}
__typedArg0 := V876
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4727 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4726)
}
__typedArg0 := tmp4726
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4728 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(True, tmp4727)
}
__typedArg0 := True
__typedArg1 := tmp4727
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4729 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4728, V877)
}
__typedArg0 := tmp4728
__typedArg1 := V877
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4pivot_1on), tmp4724, tmp4725, tmp4729)
return


} else {
__e.TailApply(PrimFunc(sym_8p), V877, V876)
return
}


}


}, 3)

tmp4808 := Call(__e, ns2_1set, symshen_4pivot_1on, tmp4708)


_ = tmp4808

tmp4809 := MakeNative(func(__e *ControlFlow) {
V880 := __e.Get(1)
_ = V880
V881 := __e.Get(2)
_ = V881
tmp4833 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V880)
}
__typedArg0 := V880
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4824 Obj

if True == tmp4833 {
tmp4831 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V880)
}
__typedArg0 := V880
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4832 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(tmp4831)
}
__typedArg0 := tmp4831
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

var ifres4826 Obj

if True == tmp4832 {
tmp4828 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V880)
}
__typedArg0 := V880
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4829 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(tmp4828)
}
__typedArg0 := tmp4828
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4830 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, tmp4829)
}
__typedArg0 := Nil
__typedArg1 := tmp4829
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

var ifres4827 Obj

if True == tmp4830 {
ifres4827 = True


} else {
ifres4827 = False


}

ifres4826 = ifres4827


} else {
ifres4826 = False


}

var ifres4825 Obj

if True == ifres4826 {
ifres4825 = True


} else {
ifres4825 = False


}

ifres4824 = ifres4825


} else {
ifres4824 = False


}

if True == ifres4824 {
tmp4810 := MakeNative(func(__e *ControlFlow) {
W882 := __e.Get(1)
_ = W882
tmp4820 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4skip, W882)
}
__typedArg0 := symshen_4skip
__typedArg1 := W882
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4820 {
__e.Return(V881)
return
} else {
tmp4811 := Call(__e, PrimFunc(symshen_4op1), W882)


tmp4812 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V880)
}
__typedArg0 := V880
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4813 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4811, tmp4812)
}
__typedArg0 := tmp4811
__typedArg1 := tmp4812
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4814 := Call(__e, PrimFunc(symshen_4op2), W882)


tmp4815 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V880)
}
__typedArg0 := V880
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4816 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4814, tmp4815)
}
__typedArg0 := tmp4814
__typedArg1 := tmp4815
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4817 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4816, Nil)
}
__typedArg0 := tmp4816
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4818 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4813, tmp4817)
}
__typedArg0 := tmp4813
__typedArg1 := tmp4817
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.TailApply(PrimFunc(symshen_4factor_1selectors_1h), tmp4818, V881)
return


}


}, 1)

tmp4821 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V880)
}
__typedArg0 := V880
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4822 := Call(__e, PrimFunc(symshen_4op), tmp4821)


__e.TailApply(tmp4810, tmp4822)
return


} else {
__e.Return(V881)
return
}


}, 2)

tmp4834 := Call(__e, ns2_1set, symshen_4factor_1selectors, tmp4809)


_ = tmp4834

tmp4835 := MakeNative(func(__e *ControlFlow) {
V885 := __e.Get(1)
_ = V885
tmp4843 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symcons_2, V885)
}
__typedArg0 := symcons_2
__typedArg1 := V885
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4843 {
__e.Return(symcons)
return
} else {
tmp4841 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7string_2, V885)
}
__typedArg0 := symshen_4_7string_2
__typedArg1 := V885
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4841 {
__e.Return(sym_8s)
return
} else {
tmp4839 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symshen_4_7vector_2, V885)
}
__typedArg0 := symshen_4_7vector_2
__typedArg1 := V885
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4839 {
__e.Return(sym_8v)
return
} else {
tmp4837 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(symtuple_2, V885)
}
__typedArg0 := symtuple_2
__typedArg1 := V885
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4837 {
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

tmp4844 := Call(__e, ns2_1set, symshen_4op, tmp4835)


_ = tmp4844

tmp4845 := MakeNative(func(__e *ControlFlow) {
V886 := __e.Get(1)
_ = V886
V887 := __e.Get(2)
_ = V887
tmp4864 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_a) {
return PrimEqual(Nil, V886)
}
__typedArg0 := Nil
__typedArg1 := V886
return Call(__e, PrimFunc(sym_a), __typedArg0, __typedArg1)
})()

if True == tmp4864 {
__e.Return(V887)
return
} else {
tmp4862 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons_2) {
return PrimIsPair(V886)
}
__typedArg0 := V886
return Call(__e, PrimFunc(symcons_2), __typedArg0)
})()

if True == tmp4862 {
tmp4858 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V886)
}
__typedArg0 := V886
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4859 := Call(__e, PrimFunc(symoccurrences), tmp4858, V887)


if True == (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym_6) {
__typedN0, __typedOK0 := TypedFloat64(tmp4859)
__typedN1, __typedOK1 := TypedFloat64(MakeNumber(1))
if __typedOK0 && __typedOK1 && HasCanonicalPrimitiveBinding(sym_6) {
return TypedMaterializeBoolean((__typedN0 > __typedN1))
}}
__typedArg0 := tmp4859
__typedArg1 := MakeNumber(1)
return Call(__e, PrimFunc(sym_6), __typedArg0, __typedArg1)
})() {
tmp4846 := MakeNative(func(__e *ControlFlow) {
W888 := __e.Get(1)
_ = W888
tmp4847 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V886)
}
__typedArg0 := V886
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4848 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V886)
}
__typedArg0 := V886
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

tmp4849 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symhd) {
return PrimHead(V886)
}
__typedArg0 := V886
return Call(__e, PrimFunc(symhd), __typedArg0)
})()

tmp4850 := Call(__e, PrimFunc(symsubst), W888, tmp4849, V887)


tmp4851 := Call(__e, PrimFunc(symshen_4factor_1selectors_1h), tmp4848, tmp4850)


tmp4852 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4851, Nil)
}
__typedArg0 := tmp4851
__typedArg1 := Nil
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4853 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(tmp4847, tmp4852)
}
__typedArg0 := tmp4847
__typedArg1 := tmp4852
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

tmp4854 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(W888, tmp4853)
}
__typedArg0 := W888
__typedArg1 := tmp4853
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})()

__e.Return((func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symcons) {
return PrimCons(symlet, tmp4854)
}
__typedArg0 := symlet
__typedArg1 := tmp4854
return Call(__e, PrimFunc(symcons), __typedArg0, __typedArg1)
})())
return


}, 1)

tmp4855 := Call(__e, PrimFunc(symgensym), symSelect)


__e.TailApply(tmp4846, tmp4855)
return


} else {
tmp4856 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symtl) {
return PrimTail(V886)
}
__typedArg0 := V886
return Call(__e, PrimFunc(symtl), __typedArg0)
})()

__e.TailApply(PrimFunc(symshen_4factor_1selectors_1h), tmp4856, V887)
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

__e.TailApply(ns2_1set, symshen_4factor_1selectors_1h, tmp4845)
return




}, 0)

