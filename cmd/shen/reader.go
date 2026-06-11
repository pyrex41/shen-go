package main

import . "github.com/tiancaiamao/shen-go/kl"

var ReaderMain = MakeNative(func(__e *ControlFlow) {
tmp3767 := MakeNative(func(__e *ControlFlow) {
V2528 := __e.Get(1)
_ = V2528
tmp3768 := MakeNative(func(__e *ControlFlow) {
W2529 := __e.Get(1)
_ = W2529
tmp3769 := MakeNative(func(__e *ControlFlow) {
W2530 := __e.Get(1)
_ = W2530
tmp3770 := MakeNative(func(__e *ControlFlow) {
W2533 := __e.Get(1)
_ = W2533
__e.Return(W2533)
return
}, 1)

tmp3771 := Call(__e, PrimFunc(symshen_4process_1sexprs), W2530)


__e.TailApply(tmp3770, tmp3771)
return


}, 1)

tmp3772 := MakeNative(func(__e *ControlFlow) {
tmp3773 := MakeNative(func(__e *ControlFlow) {
Z2531 := __e.Get(1)
_ = Z2531
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2531)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp3773, W2529)
return


}, 0)

tmp3774 := MakeNative(func(__e *ControlFlow) {
Z2532 := __e.Get(1)
_ = Z2532
tmp3775 := PrimValue(symshen_4_dresidue_d)

__e.TailApply(PrimFunc(symshen_4reader_1error), tmp3775)
return


}, 1)

tmp3776 := Call(__e, try_1catch, tmp3772, tmp3774)


__e.TailApply(tmp3769, tmp3776)
return


}, 1)

tmp3777 := PrimReadFileAsByteList(V2528)

__e.TailApply(tmp3768, tmp3777)
return


}, 1)

tmp3778 := Call(__e, ns2_1set, symread_1file, tmp3767)


_ = tmp3778

tmp3779 := MakeNative(func(__e *ControlFlow) {
V2534 := __e.Get(1)
_ = V2534
tmp3780 := PrimValue(sym_dmaximum_1print_1sequence_1size_d)

tmp3781 := Call(__e, PrimFunc(symshen_4reader_1error_1message), tmp3780, MakeNumber(0), V2534)


tmp3782 := PrimStringConcat(MakeString("reader error near here: "), tmp3781)

tmp3783 := Call(__e, PrimFunc(symshen_4proc_1nl), tmp3782)


__e.Return(PrimSimpleError(tmp3783))
return


}, 1)

tmp3784 := Call(__e, ns2_1set, symshen_4reader_1error, tmp3779)


_ = tmp3784

tmp3785 := MakeNative(func(__e *ControlFlow) {
V2542 := __e.Get(1)
_ = V2542
V2543 := __e.Get(2)
_ = V2543
V2544 := __e.Get(3)
_ = V2544
tmp3796 := PrimEqual(Nil, V2544)

if True == tmp3796 {
__e.Return(MakeString(""))
return
} else {
tmp3794 := PrimEqual(V2542, V2543)

if True == tmp3794 {
__e.Return(MakeString(""))
return
} else {
tmp3792 := PrimIsPair(V2544)

if True == tmp3792 {
tmp3786 := PrimHead(V2544)

tmp3787 := PrimNumberToString(tmp3786)

tmp3788 := PrimNumberAdd(V2543, MakeNumber(1))

tmp3789 := PrimTail(V2544)

tmp3790 := Call(__e, PrimFunc(symshen_4reader_1error_1message), V2542, tmp3788, tmp3789)


__e.Return(PrimStringConcat(tmp3787, tmp3790))
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4reader_1error_1message)
return
}


}


}


}, 3)

tmp3797 := Call(__e, ns2_1set, symshen_4reader_1error_1message, tmp3785)


_ = tmp3797

tmp3798 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dit_d))
return
}, 0)

tmp3799 := Call(__e, ns2_1set, symit, tmp3798)


_ = tmp3799

tmp3800 := MakeNative(func(__e *ControlFlow) {
V2545 := __e.Get(1)
_ = V2545
tmp3801 := MakeNative(func(__e *ControlFlow) {
W2546 := __e.Get(1)
_ = W2546
tmp3802 := MakeNative(func(__e *ControlFlow) {
W2547 := __e.Get(1)
_ = W2547
tmp3803 := MakeNative(func(__e *ControlFlow) {
W2548 := __e.Get(1)
_ = W2548
tmp3804 := MakeNative(func(__e *ControlFlow) {
W2549 := __e.Get(1)
_ = W2549
__e.TailApply(PrimFunc(symreverse), W2548)
return
}, 1)

tmp3805 := PrimCloseStream(W2546)

__e.TailApply(tmp3804, tmp3805)
return


}, 1)

tmp3806 := Call(__e, PrimFunc(symshen_4read_1file_1as_1bytelist_1help), W2546, W2547, Nil)


__e.TailApply(tmp3803, tmp3806)
return


}, 1)

tmp3807 := PrimReadByte(W2546)

__e.TailApply(tmp3802, tmp3807)
return


}, 1)

tmp3808 := PrimOpenStream(V2545, symin)

__e.TailApply(tmp3801, tmp3808)
return


}, 1)

tmp3809 := Call(__e, ns2_1set, symread_1file_1as_1bytelist, tmp3800)


_ = tmp3809

tmp3810 := MakeNative(func(__e *ControlFlow) {
V2550 := __e.Get(1)
_ = V2550
V2551 := __e.Get(2)
_ = V2551
V2552 := __e.Get(3)
_ = V2552
tmp3814 := PrimEqual(MakeNumber(-1), V2551)

if True == tmp3814 {
__e.Return(V2552)
return
} else {
tmp3811 := PrimReadByte(V2550)

tmp3812 := PrimCons(V2551, V2552)

__e.TailApply(PrimFunc(symshen_4read_1file_1as_1bytelist_1help), V2550, tmp3811, tmp3812)
return


}


}, 3)

tmp3815 := Call(__e, ns2_1set, symshen_4read_1file_1as_1bytelist_1help, tmp3810)


_ = tmp3815

tmp3816 := MakeNative(func(__e *ControlFlow) {
V2553 := __e.Get(1)
_ = V2553
tmp3817 := MakeNative(func(__e *ControlFlow) {
W2554 := __e.Get(1)
_ = W2554
tmp3818 := PrimReadByte(W2554)

__e.TailApply(PrimFunc(symshen_4rfas_1h), W2554, tmp3818, MakeString(""))
return


}, 1)

tmp3819 := PrimOpenStream(V2553, symin)

__e.TailApply(tmp3817, tmp3819)
return


}, 1)

tmp3820 := Call(__e, ns2_1set, symread_1file_1as_1string, tmp3816)


_ = tmp3820

tmp3821 := MakeNative(func(__e *ControlFlow) {
V2555 := __e.Get(1)
_ = V2555
V2556 := __e.Get(2)
_ = V2556
V2557 := __e.Get(3)
_ = V2557
tmp3827 := PrimEqual(MakeNumber(-1), V2556)

if True == tmp3827 {
tmp3822 := PrimCloseStream(V2555)

_ = tmp3822

__e.Return(V2557)
return


} else {
tmp3823 := PrimReadByte(V2555)

tmp3824 := PrimNumberToString(V2556)

tmp3825 := PrimStringConcat(V2557, tmp3824)

__e.TailApply(PrimFunc(symshen_4rfas_1h), V2555, tmp3823, tmp3825)
return


}


}, 3)

tmp3828 := Call(__e, ns2_1set, symshen_4rfas_1h, tmp3821)


_ = tmp3828

tmp3829 := MakeNative(func(__e *ControlFlow) {
V2558 := __e.Get(1)
_ = V2558
tmp3830 := Call(__e, PrimFunc(symread), V2558)


__e.TailApply(PrimFunc(symeval_1kl), tmp3830)
return


}, 1)

tmp3831 := Call(__e, ns2_1set, syminput, tmp3829)


_ = tmp3831

tmp3832 := MakeNative(func(__e *ControlFlow) {
V2559 := __e.Get(1)
_ = V2559
V2560 := __e.Get(2)
_ = V2560
tmp3833 := MakeNative(func(__e *ControlFlow) {
W2561 := __e.Get(1)
_ = W2561
tmp3834 := MakeNative(func(__e *ControlFlow) {
W2562 := __e.Get(1)
_ = W2562
tmp3840 := Call(__e, PrimFunc(symshen_4rectify_1type), V2559)


tmp3841 := Call(__e, PrimFunc(symshen_4typecheck), W2562, tmp3840)


tmp3842 := PrimEqual(False, tmp3841)

if True == tmp3842 {
tmp3835 := Call(__e, PrimFunc(symshen_4app), V2559, MakeString("\n"), symshen_4r)


tmp3836 := PrimStringConcat(MakeString(" is not of type "), tmp3835)

tmp3837 := Call(__e, PrimFunc(symshen_4app), W2562, tmp3836, symshen_4r)


tmp3838 := PrimStringConcat(MakeString("type error: "), tmp3837)

__e.Return(PrimSimpleError(tmp3838))
return


} else {
__e.TailApply(PrimFunc(symeval_1kl), W2562)
return
}


}, 1)

tmp3843 := Call(__e, PrimFunc(symread), V2560)


__e.TailApply(tmp3834, tmp3843)
return


}, 1)

tmp3844 := Call(__e, PrimFunc(symshen_4monotype), V2559)


__e.TailApply(tmp3833, tmp3844)
return


}, 2)

tmp3845 := Call(__e, ns2_1set, syminput_7, tmp3832)


_ = tmp3845

tmp3846 := MakeNative(func(__e *ControlFlow) {
V2563 := __e.Get(1)
_ = V2563
tmp3853 := PrimIsPair(V2563)

if True == tmp3853 {
tmp3847 := MakeNative(func(__e *ControlFlow) {
Z2564 := __e.Get(1)
_ = Z2564
__e.TailApply(PrimFunc(symshen_4monotype), Z2564)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp3847, V2563)
return


} else {
tmp3851 := PrimIsVariable(V2563)

if True == tmp3851 {
tmp3848 := Call(__e, PrimFunc(symshen_4app), V2563, MakeString("\n"), symshen_4a)


tmp3849 := PrimStringConcat(MakeString("input+ expects a monotype: not "), tmp3848)

__e.Return(PrimSimpleError(tmp3849))
return


} else {
__e.Return(V2563)
return
}


}


}, 1)

tmp3854 := Call(__e, ns2_1set, symshen_4monotype, tmp3846)


_ = tmp3854

tmp3855 := MakeNative(func(__e *ControlFlow) {
V2565 := __e.Get(1)
_ = V2565
tmp3856 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2565)


tmp3857 := MakeNative(func(__e *ControlFlow) {
Z2566 := __e.Get(1)
_ = Z2566
__e.TailApply(PrimFunc(symshen_4return_2), Z2566)
return
}, 1)

__e.TailApply(PrimFunc(symshen_4read_1loop), V2565, tmp3856, Nil, tmp3857)
return


}, 1)

tmp3858 := Call(__e, ns2_1set, symlineread, tmp3855)


_ = tmp3858

tmp3859 := MakeNative(func(__e *ControlFlow) {
V2567 := __e.Get(1)
_ = V2567
tmp3860 := MakeNative(func(__e *ControlFlow) {
W2568 := __e.Get(1)
_ = W2568
tmp3861 := MakeNative(func(__e *ControlFlow) {
W2569 := __e.Get(1)
_ = W2569
tmp3862 := MakeNative(func(__e *ControlFlow) {
W2571 := __e.Get(1)
_ = W2571
__e.Return(W2571)
return
}, 1)

tmp3863 := Call(__e, PrimFunc(symshen_4process_1sexprs), W2569)


__e.TailApply(tmp3862, tmp3863)
return


}, 1)

tmp3864 := MakeNative(func(__e *ControlFlow) {
Z2570 := __e.Get(1)
_ = Z2570
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2570)
return
}, 1)

tmp3865 := Call(__e, PrimFunc(symcompile), tmp3864, W2568)


__e.TailApply(tmp3861, tmp3865)
return


}, 1)

tmp3866 := Call(__e, PrimFunc(symshen_4str_1_6bytes), V2567)


__e.TailApply(tmp3860, tmp3866)
return


}, 1)

tmp3867 := Call(__e, ns2_1set, symread_1from_1string, tmp3859)


_ = tmp3867

tmp3868 := MakeNative(func(__e *ControlFlow) {
V2572 := __e.Get(1)
_ = V2572
tmp3869 := MakeNative(func(__e *ControlFlow) {
W2573 := __e.Get(1)
_ = W2573
tmp3870 := MakeNative(func(__e *ControlFlow) {
W2574 := __e.Get(1)
_ = W2574
__e.Return(W2574)
return
}, 1)

tmp3871 := MakeNative(func(__e *ControlFlow) {
Z2575 := __e.Get(1)
_ = Z2575
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2575)
return
}, 1)

tmp3872 := Call(__e, PrimFunc(symcompile), tmp3871, W2573)


__e.TailApply(tmp3870, tmp3872)
return


}, 1)

tmp3873 := Call(__e, PrimFunc(symshen_4str_1_6bytes), V2572)


__e.TailApply(tmp3869, tmp3873)
return


}, 1)

tmp3874 := Call(__e, ns2_1set, symread_1from_1string_1unprocessed, tmp3868)


_ = tmp3874

tmp3875 := MakeNative(func(__e *ControlFlow) {
V2576 := __e.Get(1)
_ = V2576
tmp3883 := PrimEqual(MakeString(""), V2576)

if True == tmp3883 {
__e.Return(Nil)
return
} else {
tmp3881 := Call(__e, PrimFunc(symshen_4_7string_2), V2576)


if True == tmp3881 {
tmp3876 := Call(__e, PrimFunc(symhdstr), V2576)


tmp3877 := PrimStringToNumber(tmp3876)

tmp3878 := PrimTailString(V2576)

tmp3879 := Call(__e, PrimFunc(symshen_4str_1_6bytes), tmp3878)


__e.Return(PrimCons(tmp3877, tmp3879))
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4str_1_6bytes)
return
}


}


}, 1)

tmp3884 := Call(__e, ns2_1set, symshen_4str_1_6bytes, tmp3875)


_ = tmp3884

tmp3885 := MakeNative(func(__e *ControlFlow) {
V2577 := __e.Get(1)
_ = V2577
tmp3886 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2577)


tmp3887 := MakeNative(func(__e *ControlFlow) {
Z2578 := __e.Get(1)
_ = Z2578
__e.TailApply(PrimFunc(symshen_4whitespace_2), Z2578)
return
}, 1)

tmp3888 := Call(__e, PrimFunc(symshen_4read_1loop), V2577, tmp3886, Nil, tmp3887)


__e.Return(PrimHead(tmp3888))
return


}, 1)

tmp3889 := Call(__e, ns2_1set, symread, tmp3885)


_ = tmp3889

tmp3890 := MakeNative(func(__e *ControlFlow) {
V2579 := __e.Get(1)
_ = V2579
tmp3893 := Call(__e, PrimFunc(symshen_4char_1stinput_2), V2579)


if True == tmp3893 {
tmp3891 := Call(__e, PrimFunc(symshen_4read_1unit_1string), V2579)


__e.Return(PrimStringToNumber(tmp3891))
return


} else {
__e.Return(PrimReadByte(V2579))
return
}


}, 1)

tmp3894 := Call(__e, ns2_1set, symshen_4my_1read_1byte, tmp3890)


_ = tmp3894

tmp3895 := MakeNative(func(__e *ControlFlow) {
V2584 := __e.Get(1)
_ = V2584
V2585 := __e.Get(2)
_ = V2585
V2586 := __e.Get(3)
_ = V2586
V2587 := __e.Get(4)
_ = V2587
tmp3918 := PrimEqual(MakeNumber(94), V2585)

if True == tmp3918 {
__e.Return(PrimSimpleError(MakeString("read aborted")))
return
} else {
tmp3916 := PrimEqual(MakeNumber(-1), V2585)

if True == tmp3916 {
tmp3898 := Call(__e, PrimFunc(symempty_2), V2586)


if True == tmp3898 {
__e.Return(PrimSimpleError(MakeString("error: empty stream")))
return
} else {
tmp3896 := MakeNative(func(__e *ControlFlow) {
Z2588 := __e.Get(1)
_ = Z2588
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2588)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp3896, V2586)
return


}


} else {
tmp3914 := PrimEqual(MakeNumber(0), V2585)

if True == tmp3914 {
tmp3899 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2584)


__e.TailApply(PrimFunc(symshen_4read_1loop), V2584, tmp3899, V2586, V2587)
return


} else {
tmp3912 := Call(__e, V2587, V2585)


if True == tmp3912 {
tmp3900 := MakeNative(func(__e *ControlFlow) {
W2589 := __e.Get(1)
_ = W2589
tmp3906 := Call(__e, PrimFunc(symshen_4nothing_1doing_2), W2589)


if True == tmp3906 {
tmp3901 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2584)


tmp3902 := PrimCons(V2585, Nil)

tmp3903 := Call(__e, PrimFunc(symappend), V2586, tmp3902)


__e.TailApply(PrimFunc(symshen_4read_1loop), V2584, tmp3901, tmp3903, V2587)
return


} else {
tmp3904 := Call(__e, PrimFunc(symshen_4record_1it), V2586)


_ = tmp3904

__e.Return(W2589)
return


}


}, 1)

tmp3907 := Call(__e, PrimFunc(symshen_4try_1parse), V2586)


__e.TailApply(tmp3900, tmp3907)
return


} else {
tmp3908 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2584)


tmp3909 := PrimCons(V2585, Nil)

tmp3910 := Call(__e, PrimFunc(symappend), V2586, tmp3909)


__e.TailApply(PrimFunc(symshen_4read_1loop), V2584, tmp3908, tmp3910, V2587)
return


}


}


}


}


}, 4)

tmp3919 := Call(__e, ns2_1set, symshen_4read_1loop, tmp3895)


_ = tmp3919

tmp3920 := MakeNative(func(__e *ControlFlow) {
V2590 := __e.Get(1)
_ = V2590
tmp3921 := MakeNative(func(__e *ControlFlow) {
W2591 := __e.Get(1)
_ = W2591
tmp3923 := Call(__e, PrimFunc(symshen_4nothing_1doing_2), W2591)


if True == tmp3923 {
__e.Return(symshen_4i_1failed_b)
return
} else {
__e.TailApply(PrimFunc(symshen_4process_1sexprs), W2591)
return
}


}, 1)

tmp3924 := MakeNative(func(__e *ControlFlow) {
tmp3925 := MakeNative(func(__e *ControlFlow) {
Z2592 := __e.Get(1)
_ = Z2592
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2592)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp3925, V2590)
return


}, 0)

tmp3926 := MakeNative(func(__e *ControlFlow) {
Z2593 := __e.Get(1)
_ = Z2593
__e.Return(symshen_4i_1failed_b)
return
}, 1)

tmp3927 := Call(__e, try_1catch, tmp3924, tmp3926)


__e.TailApply(tmp3921, tmp3927)
return


}, 1)

tmp3928 := Call(__e, ns2_1set, symshen_4try_1parse, tmp3920)


_ = tmp3928

tmp3929 := MakeNative(func(__e *ControlFlow) {
V2596 := __e.Get(1)
_ = V2596
tmp3933 := PrimEqual(symshen_4i_1failed_b, V2596)

if True == tmp3933 {
__e.Return(True)
return
} else {
tmp3931 := PrimEqual(Nil, V2596)

if True == tmp3931 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp3934 := Call(__e, ns2_1set, symshen_4nothing_1doing_2, tmp3929)


_ = tmp3934

tmp3935 := MakeNative(func(__e *ControlFlow) {
V2597 := __e.Get(1)
_ = V2597
tmp3936 := Call(__e, PrimFunc(symshen_4bytes_1_6string), V2597)


__e.Return(PrimSet(symshen_4_dit_d, tmp3936))
return


}, 1)

tmp3937 := Call(__e, ns2_1set, symshen_4record_1it, tmp3935)


_ = tmp3937

tmp3938 := MakeNative(func(__e *ControlFlow) {
V2598 := __e.Get(1)
_ = V2598
tmp3946 := PrimEqual(Nil, V2598)

if True == tmp3946 {
__e.Return(MakeString(""))
return
} else {
tmp3944 := PrimIsPair(V2598)

if True == tmp3944 {
tmp3939 := PrimHead(V2598)

tmp3940 := PrimNumberToString(tmp3939)

tmp3941 := PrimTail(V2598)

tmp3942 := Call(__e, PrimFunc(symshen_4bytes_1_6string), tmp3941)


__e.Return(PrimStringConcat(tmp3940, tmp3942))
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4bytes_1_6string)
return
}


}


}, 1)

tmp3947 := Call(__e, ns2_1set, symshen_4bytes_1_6string, tmp3938)


_ = tmp3947

tmp3948 := MakeNative(func(__e *ControlFlow) {
V2599 := __e.Get(1)
_ = V2599
tmp3949 := MakeNative(func(__e *ControlFlow) {
W2600 := __e.Get(1)
_ = W2600
tmp3950 := MakeNative(func(__e *ControlFlow) {
W2601 := __e.Get(1)
_ = W2601
tmp3951 := MakeNative(func(__e *ControlFlow) {
W2602 := __e.Get(1)
_ = W2602
tmp3952 := MakeNative(func(__e *ControlFlow) {
Z2603 := __e.Get(1)
_ = Z2603
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2603, W2602)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp3952, W2600)
return


}, 1)

tmp3953 := Call(__e, PrimFunc(symshen_4find_1types), W2600)


__e.TailApply(tmp3951, tmp3953)
return


}, 1)

tmp3954 := Call(__e, PrimFunc(symshen_4find_1arities), W2600)


__e.TailApply(tmp3950, tmp3954)
return


}, 1)

tmp3955 := Call(__e, PrimFunc(symshen_4unpackage_emacroexpand), V2599)


__e.TailApply(tmp3949, tmp3955)
return


}, 1)

tmp3956 := Call(__e, ns2_1set, symshen_4process_1sexprs, tmp3948)


_ = tmp3956

tmp3957 := MakeNative(func(__e *ControlFlow) {
V2604 := __e.Get(1)
_ = V2604
tmp3979 := PrimIsPair(V2604)

var ifres3970 Obj

if True == tmp3979 {
tmp3977 := PrimTail(V2604)

tmp3978 := PrimIsPair(tmp3977)

var ifres3972 Obj

if True == tmp3978 {
tmp3974 := PrimHead(V2604)

tmp3975 := PrimIntern(MakeString(":"))

tmp3976 := PrimEqual(tmp3974, tmp3975)

var ifres3973 Obj

if True == tmp3976 {
ifres3973 = True


} else {
ifres3973 = False


}

ifres3972 = ifres3973


} else {
ifres3972 = False


}

var ifres3971 Obj

if True == ifres3972 {
ifres3971 = True


} else {
ifres3971 = False


}

ifres3970 = ifres3971


} else {
ifres3970 = False


}

if True == ifres3970 {
tmp3958 := PrimTail(V2604)

tmp3959 := PrimHead(tmp3958)

tmp3960 := PrimTail(V2604)

tmp3961 := PrimTail(tmp3960)

tmp3962 := Call(__e, PrimFunc(symshen_4find_1types), tmp3961)


__e.Return(PrimCons(tmp3959, tmp3962))
return


} else {
tmp3968 := PrimIsPair(V2604)

if True == tmp3968 {
tmp3963 := PrimHead(V2604)

tmp3964 := Call(__e, PrimFunc(symshen_4find_1types), tmp3963)


tmp3965 := PrimTail(V2604)

tmp3966 := Call(__e, PrimFunc(symshen_4find_1types), tmp3965)


__e.TailApply(PrimFunc(symappend), tmp3964, tmp3966)
return


} else {
__e.Return(Nil)
return
}


}


}, 1)

tmp3980 := Call(__e, ns2_1set, symshen_4find_1types, tmp3957)


_ = tmp3980

tmp3981 := MakeNative(func(__e *ControlFlow) {
V2607 := __e.Get(1)
_ = V2607
tmp4030 := PrimIsPair(V2607)

var ifres4011 Obj

if True == tmp4030 {
tmp4028 := PrimHead(V2607)

tmp4029 := PrimEqual(symdefine, tmp4028)

var ifres4013 Obj

if True == tmp4029 {
tmp4026 := PrimTail(V2607)

tmp4027 := PrimIsPair(tmp4026)

var ifres4015 Obj

if True == tmp4027 {
tmp4023 := PrimTail(V2607)

tmp4024 := PrimTail(tmp4023)

tmp4025 := PrimIsPair(tmp4024)

var ifres4017 Obj

if True == tmp4025 {
tmp4019 := PrimTail(V2607)

tmp4020 := PrimTail(tmp4019)

tmp4021 := PrimHead(tmp4020)

tmp4022 := PrimEqual(sym_i, tmp4021)

var ifres4018 Obj

if True == tmp4022 {
ifres4018 = True


} else {
ifres4018 = False


}

ifres4017 = ifres4018


} else {
ifres4017 = False


}

var ifres4016 Obj

if True == ifres4017 {
ifres4016 = True


} else {
ifres4016 = False


}

ifres4015 = ifres4016


} else {
ifres4015 = False


}

var ifres4014 Obj

if True == ifres4015 {
ifres4014 = True


} else {
ifres4014 = False


}

ifres4013 = ifres4014


} else {
ifres4013 = False


}

var ifres4012 Obj

if True == ifres4013 {
ifres4012 = True


} else {
ifres4012 = False


}

ifres4011 = ifres4012


} else {
ifres4011 = False


}

if True == ifres4011 {
tmp3982 := PrimTail(V2607)

tmp3983 := PrimHead(tmp3982)

tmp3984 := PrimTail(V2607)

tmp3985 := PrimHead(tmp3984)

tmp3986 := PrimTail(V2607)

tmp3987 := PrimTail(tmp3986)

tmp3988 := PrimTail(tmp3987)

tmp3989 := Call(__e, PrimFunc(symshen_4find_1arity), tmp3985, MakeNumber(1), tmp3988)


__e.TailApply(PrimFunc(symshen_4store_1arity), tmp3983, tmp3989)
return


} else {
tmp4009 := PrimIsPair(V2607)

var ifres4001 Obj

if True == tmp4009 {
tmp4007 := PrimHead(V2607)

tmp4008 := PrimEqual(symdefine, tmp4007)

var ifres4003 Obj

if True == tmp4008 {
tmp4005 := PrimTail(V2607)

tmp4006 := PrimIsPair(tmp4005)

var ifres4004 Obj

if True == tmp4006 {
ifres4004 = True


} else {
ifres4004 = False


}

ifres4003 = ifres4004


} else {
ifres4003 = False


}

var ifres4002 Obj

if True == ifres4003 {
ifres4002 = True


} else {
ifres4002 = False


}

ifres4001 = ifres4002


} else {
ifres4001 = False


}

if True == ifres4001 {
tmp3990 := PrimTail(V2607)

tmp3991 := PrimHead(tmp3990)

tmp3992 := PrimTail(V2607)

tmp3993 := PrimHead(tmp3992)

tmp3994 := PrimTail(V2607)

tmp3995 := PrimTail(tmp3994)

tmp3996 := Call(__e, PrimFunc(symshen_4find_1arity), tmp3993, MakeNumber(0), tmp3995)


__e.TailApply(PrimFunc(symshen_4store_1arity), tmp3991, tmp3996)
return


} else {
tmp3999 := PrimIsPair(V2607)

if True == tmp3999 {
tmp3997 := MakeNative(func(__e *ControlFlow) {
Z2608 := __e.Get(1)
_ = Z2608
__e.TailApply(PrimFunc(symshen_4find_1arities), Z2608)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp3997, V2607)
return


} else {
__e.Return(symshen_4skip)
return
}


}


}


}, 1)

tmp4031 := Call(__e, ns2_1set, symshen_4find_1arities, tmp3981)


_ = tmp4031

tmp4032 := MakeNative(func(__e *ControlFlow) {
V2609 := __e.Get(1)
_ = V2609
V2610 := __e.Get(2)
_ = V2610
tmp4033 := MakeNative(func(__e *ControlFlow) {
W2611 := __e.Get(1)
_ = W2611
tmp4044 := PrimEqual(W2611, MakeNumber(-1))

if True == tmp4044 {
__e.TailApply(PrimFunc(symshen_4execute_1store_1arity), V2609, V2610)
return
} else {
tmp4042 := PrimEqual(W2611, V2610)

if True == tmp4042 {
__e.Return(symshen_4skip)
return
} else {
tmp4040 := Call(__e, PrimFunc(symshen_4sysfunc_2), V2609)


if True == tmp4040 {
tmp4034 := Call(__e, PrimFunc(symshen_4app), V2609, MakeString(" is a system function\n"), symshen_4a)


__e.Return(PrimSimpleError(tmp4034))
return


} else {
tmp4035 := Call(__e, PrimFunc(symshen_4app), V2609, MakeString(" may cause errors\n"), symshen_4a)


tmp4036 := PrimStringConcat(MakeString("changing the arity of "), tmp4035)

tmp4037 := Call(__e, PrimFunc(symstoutput))


tmp4038 := Call(__e, PrimFunc(sympr), tmp4036, tmp4037)


_ = tmp4038

__e.TailApply(PrimFunc(symshen_4execute_1store_1arity), V2609, V2610)
return


}


}


}


}, 1)

tmp4045 := Call(__e, PrimFunc(symarity), V2609)


__e.TailApply(tmp4033, tmp4045)
return


}, 2)

tmp4046 := Call(__e, ns2_1set, symshen_4store_1arity, tmp4032)


_ = tmp4046

tmp4047 := MakeNative(func(__e *ControlFlow) {
V2612 := __e.Get(1)
_ = V2612
V2613 := __e.Get(2)
_ = V2613
tmp4052 := PrimEqual(MakeNumber(0), V2613)

if True == tmp4052 {
tmp4048 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V2612, symarity, MakeNumber(0), tmp4048)
return


} else {
tmp4049 := PrimValue(sym_dproperty_1vector_d)

tmp4050 := Call(__e, PrimFunc(symput), V2612, symarity, V2613, tmp4049)


_ = tmp4050

__e.TailApply(PrimFunc(symshen_4update_1lambdatable), V2612, V2613)
return


}


}, 2)

tmp4053 := Call(__e, ns2_1set, symshen_4execute_1store_1arity, tmp4047)


_ = tmp4053

tmp4054 := MakeNative(func(__e *ControlFlow) {
V2614 := __e.Get(1)
_ = V2614
V2615 := __e.Get(2)
_ = V2615
tmp4055 := MakeNative(func(__e *ControlFlow) {
W2616 := __e.Get(1)
_ = W2616
tmp4056 := MakeNative(func(__e *ControlFlow) {
W2617 := __e.Get(1)
_ = W2617
__e.Return(V2615)
return
}, 1)

tmp4057 := PrimCons(V2614, W2616)

tmp4058 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp4057)


__e.TailApply(tmp4056, tmp4058)
return


}, 1)

tmp4059 := PrimCons(V2614, Nil)

tmp4060 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp4059, V2615)


tmp4061 := Call(__e, PrimFunc(symeval_1kl), tmp4060)


__e.TailApply(tmp4055, tmp4061)
return


}, 2)

tmp4062 := Call(__e, ns2_1set, symshen_4update_1lambdatable, tmp4054)


_ = tmp4062

tmp4063 := MakeNative(func(__e *ControlFlow) {
V2620 := __e.Get(1)
_ = V2620
V2621 := __e.Get(2)
_ = V2621
tmp4081 := PrimEqual(MakeNumber(0), V2621)

if True == tmp4081 {
__e.Return(symshen_4skip)
return
} else {
tmp4079 := PrimEqual(MakeNumber(1), V2621)

if True == tmp4079 {
tmp4064 := MakeNative(func(__e *ControlFlow) {
W2622 := __e.Get(1)
_ = W2622
tmp4065 := PrimCons(W2622, Nil)

tmp4066 := Call(__e, PrimFunc(symappend), V2620, tmp4065)


tmp4067 := PrimCons(tmp4066, Nil)

tmp4068 := PrimCons(W2622, tmp4067)

__e.Return(PrimCons(symlambda, tmp4068))
return


}, 1)

tmp4069 := Call(__e, PrimFunc(symgensym), symY)


__e.TailApply(tmp4064, tmp4069)
return


} else {
tmp4070 := MakeNative(func(__e *ControlFlow) {
W2623 := __e.Get(1)
_ = W2623
tmp4071 := PrimCons(W2623, Nil)

tmp4072 := Call(__e, PrimFunc(symappend), V2620, tmp4071)


tmp4073 := PrimNumberSubtract(V2621, MakeNumber(1))

tmp4074 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp4072, tmp4073)


tmp4075 := PrimCons(tmp4074, Nil)

tmp4076 := PrimCons(W2623, tmp4075)

__e.Return(PrimCons(symlambda, tmp4076))
return


}, 1)

tmp4077 := Call(__e, PrimFunc(symgensym), symY)


__e.TailApply(tmp4070, tmp4077)
return


}


}


}, 2)

tmp4082 := Call(__e, ns2_1set, symshen_4lambda_1function, tmp4063)


_ = tmp4082

tmp4083 := MakeNative(func(__e *ControlFlow) {
V2633 := __e.Get(1)
_ = V2633
V2634 := __e.Get(2)
_ = V2634
V2635 := __e.Get(3)
_ = V2635
tmp4106 := PrimEqual(Nil, V2635)

if True == tmp4106 {
tmp4084 := PrimCons(V2633, V2634)

__e.Return(PrimCons(tmp4084, Nil))
return


} else {
tmp4104 := PrimIsPair(V2635)

var ifres4095 Obj

if True == tmp4104 {
tmp4102 := PrimHead(V2635)

tmp4103 := PrimIsPair(tmp4102)

var ifres4097 Obj

if True == tmp4103 {
tmp4099 := PrimHead(V2635)

tmp4100 := PrimHead(tmp4099)

tmp4101 := PrimEqual(V2633, tmp4100)

var ifres4098 Obj

if True == tmp4101 {
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

if True == ifres4095 {
tmp4085 := PrimHead(V2635)

tmp4086 := PrimHead(tmp4085)

tmp4087 := PrimCons(tmp4086, V2634)

tmp4088 := PrimTail(V2635)

__e.Return(PrimCons(tmp4087, tmp4088))
return


} else {
tmp4093 := PrimIsPair(V2635)

if True == tmp4093 {
tmp4089 := PrimHead(V2635)

tmp4090 := PrimTail(V2635)

tmp4091 := Call(__e, PrimFunc(symshen_4assoc_1_6), V2633, V2634, tmp4090)


__e.Return(PrimCons(tmp4089, tmp4091))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.assoc->")))
return
}


}


}


}, 3)

tmp4107 := Call(__e, ns2_1set, symshen_4assoc_1_6, tmp4083)


_ = tmp4107

tmp4108 := MakeNative(func(__e *ControlFlow) {
V2650 := __e.Get(1)
_ = V2650
V2651 := __e.Get(2)
_ = V2651
V2652 := __e.Get(3)
_ = V2652
tmp4155 := PrimEqual(MakeNumber(0), V2651)

var ifres4148 Obj

if True == tmp4155 {
tmp4154 := PrimIsPair(V2652)

var ifres4150 Obj

if True == tmp4154 {
tmp4152 := PrimHead(V2652)

tmp4153 := PrimEqual(tmp4152, sym_1_6)

var ifres4151 Obj

if True == tmp4153 {
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

if True == ifres4148 {
__e.Return(MakeNumber(0))
return
} else {
tmp4146 := PrimEqual(MakeNumber(0), V2651)

var ifres4139 Obj

if True == tmp4146 {
tmp4145 := PrimIsPair(V2652)

var ifres4141 Obj

if True == tmp4145 {
tmp4143 := PrimHead(V2652)

tmp4144 := PrimEqual(tmp4143, sym_5_1)

var ifres4142 Obj

if True == tmp4144 {
ifres4142 = True


} else {
ifres4142 = False


}

ifres4141 = ifres4142


} else {
ifres4141 = False


}

var ifres4140 Obj

if True == ifres4141 {
ifres4140 = True


} else {
ifres4140 = False


}

ifres4139 = ifres4140


} else {
ifres4139 = False


}

if True == ifres4139 {
__e.Return(MakeNumber(0))
return
} else {
tmp4137 := PrimEqual(MakeNumber(0), V2651)

var ifres4134 Obj

if True == tmp4137 {
tmp4136 := PrimIsPair(V2652)

var ifres4135 Obj

if True == tmp4136 {
ifres4135 = True


} else {
ifres4135 = False


}

ifres4134 = ifres4135


} else {
ifres4134 = False


}

if True == ifres4134 {
tmp4109 := PrimTail(V2652)

tmp4110 := Call(__e, PrimFunc(symshen_4find_1arity), V2650, MakeNumber(0), tmp4109)


__e.Return(PrimNumberAdd(MakeNumber(1), tmp4110))
return


} else {
tmp4132 := PrimEqual(MakeNumber(1), V2651)

var ifres4125 Obj

if True == tmp4132 {
tmp4131 := PrimIsPair(V2652)

var ifres4127 Obj

if True == tmp4131 {
tmp4129 := PrimHead(V2652)

tmp4130 := PrimEqual(sym_j, tmp4129)

var ifres4128 Obj

if True == tmp4130 {
ifres4128 = True


} else {
ifres4128 = False


}

ifres4127 = ifres4128


} else {
ifres4127 = False


}

var ifres4126 Obj

if True == ifres4127 {
ifres4126 = True


} else {
ifres4126 = False


}

ifres4125 = ifres4126


} else {
ifres4125 = False


}

if True == ifres4125 {
tmp4111 := PrimTail(V2652)

__e.TailApply(PrimFunc(symshen_4find_1arity), V2650, MakeNumber(0), tmp4111)
return


} else {
tmp4123 := PrimEqual(MakeNumber(1), V2651)

var ifres4120 Obj

if True == tmp4123 {
tmp4122 := PrimIsPair(V2652)

var ifres4121 Obj

if True == tmp4122 {
ifres4121 = True


} else {
ifres4121 = False


}

ifres4120 = ifres4121


} else {
ifres4120 = False


}

if True == ifres4120 {
tmp4112 := PrimTail(V2652)

__e.TailApply(PrimFunc(symshen_4find_1arity), V2650, MakeNumber(1), tmp4112)
return


} else {
tmp4118 := PrimEqual(MakeNumber(1), V2651)

if True == tmp4118 {
tmp4113 := Call(__e, PrimFunc(symshen_4app), V2650, MakeString(" definition: missing }\n"), symshen_4a)


tmp4114 := PrimStringConcat(MakeString("syntax error in "), tmp4113)

__e.Return(PrimSimpleError(tmp4114))
return


} else {
tmp4115 := Call(__e, PrimFunc(symshen_4app), V2650, MakeString(" definition: missing -> or <-\n"), symshen_4a)


tmp4116 := PrimStringConcat(MakeString("syntax error in "), tmp4115)

__e.Return(PrimSimpleError(tmp4116))
return


}


}


}


}


}


}


}, 3)

tmp4156 := Call(__e, ns2_1set, symshen_4find_1arity, tmp4108)


_ = tmp4156

tmp4157 := MakeNative(func(__e *ControlFlow) {
V2653 := __e.Get(1)
_ = V2653
tmp4158 := MakeNative(func(__e *ControlFlow) {
W2654 := __e.Get(1)
_ = W2654
tmp4403 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2654)


if True == tmp4403 {
tmp4159 := MakeNative(func(__e *ControlFlow) {
W2665 := __e.Get(1)
_ = W2665
tmp4371 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2665)


if True == tmp4371 {
tmp4160 := MakeNative(func(__e *ControlFlow) {
W2676 := __e.Get(1)
_ = W2676
tmp4353 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2676)


if True == tmp4353 {
tmp4161 := MakeNative(func(__e *ControlFlow) {
W2682 := __e.Get(1)
_ = W2682
tmp4335 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2682)


if True == tmp4335 {
tmp4162 := MakeNative(func(__e *ControlFlow) {
W2688 := __e.Get(1)
_ = W2688
tmp4317 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2688)


if True == tmp4317 {
tmp4163 := MakeNative(func(__e *ControlFlow) {
W2694 := __e.Get(1)
_ = W2694
tmp4298 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2694)


if True == tmp4298 {
tmp4164 := MakeNative(func(__e *ControlFlow) {
W2700 := __e.Get(1)
_ = W2700
tmp4273 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2700)


if True == tmp4273 {
tmp4165 := MakeNative(func(__e *ControlFlow) {
W2708 := __e.Get(1)
_ = W2708
tmp4254 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2708)


if True == tmp4254 {
tmp4166 := MakeNative(func(__e *ControlFlow) {
W2714 := __e.Get(1)
_ = W2714
tmp4235 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2714)


if True == tmp4235 {
tmp4167 := MakeNative(func(__e *ControlFlow) {
W2720 := __e.Get(1)
_ = W2720
tmp4218 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2720)


if True == tmp4218 {
tmp4168 := MakeNative(func(__e *ControlFlow) {
W2726 := __e.Get(1)
_ = W2726
tmp4198 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2726)


if True == tmp4198 {
tmp4169 := MakeNative(func(__e *ControlFlow) {
W2733 := __e.Get(1)
_ = W2733
tmp4181 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2733)


if True == tmp4181 {
tmp4170 := MakeNative(func(__e *ControlFlow) {
W2739 := __e.Get(1)
_ = W2739
tmp4172 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2739)


if True == tmp4172 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2739)
return
}


}, 1)

tmp4173 := MakeNative(func(__e *ControlFlow) {
W2740 := __e.Get(1)
_ = W2740
tmp4177 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2740)


if True == tmp4177 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4174 := MakeNative(func(__e *ControlFlow) {
W2741 := __e.Get(1)
_ = W2741
__e.TailApply(PrimFunc(symshen_4comb), W2741, Nil)
return
}, 1)

tmp4175 := Call(__e, PrimFunc(symshen_4in_1_6), W2740)


__e.TailApply(tmp4174, tmp4175)
return


}


}, 1)

tmp4178 := Call(__e, PrimFunc(sym_5e_6), V2653)


tmp4179 := Call(__e, tmp4173, tmp4178)


__e.TailApply(tmp4170, tmp4179)
return


} else {
__e.Return(W2733)
return
}


}, 1)

tmp4182 := MakeNative(func(__e *ControlFlow) {
W2734 := __e.Get(1)
_ = W2734
tmp4194 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2734)


if True == tmp4194 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4183 := MakeNative(func(__e *ControlFlow) {
W2735 := __e.Get(1)
_ = W2735
tmp4184 := MakeNative(func(__e *ControlFlow) {
W2736 := __e.Get(1)
_ = W2736
tmp4190 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2736)


if True == tmp4190 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4185 := MakeNative(func(__e *ControlFlow) {
W2737 := __e.Get(1)
_ = W2737
tmp4186 := MakeNative(func(__e *ControlFlow) {
W2738 := __e.Get(1)
_ = W2738
__e.TailApply(PrimFunc(symshen_4comb), W2738, W2737)
return
}, 1)

tmp4187 := Call(__e, PrimFunc(symshen_4in_1_6), W2736)


__e.TailApply(tmp4186, tmp4187)
return


}, 1)

tmp4188 := Call(__e, PrimFunc(symshen_4_5_1out), W2736)


__e.TailApply(tmp4185, tmp4188)
return


}


}, 1)

tmp4191 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2735)


__e.TailApply(tmp4184, tmp4191)
return


}, 1)

tmp4192 := Call(__e, PrimFunc(symshen_4in_1_6), W2734)


__e.TailApply(tmp4183, tmp4192)
return


}


}, 1)

tmp4195 := Call(__e, PrimFunc(symshen_4_5whitespaces_6), V2653)


tmp4196 := Call(__e, tmp4182, tmp4195)


__e.TailApply(tmp4169, tmp4196)
return


} else {
__e.Return(W2726)
return
}


}, 1)

tmp4199 := MakeNative(func(__e *ControlFlow) {
W2727 := __e.Get(1)
_ = W2727
tmp4214 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2727)


if True == tmp4214 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4200 := MakeNative(func(__e *ControlFlow) {
W2728 := __e.Get(1)
_ = W2728
tmp4201 := MakeNative(func(__e *ControlFlow) {
W2729 := __e.Get(1)
_ = W2729
tmp4202 := MakeNative(func(__e *ControlFlow) {
W2730 := __e.Get(1)
_ = W2730
tmp4209 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2730)


if True == tmp4209 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4203 := MakeNative(func(__e *ControlFlow) {
W2731 := __e.Get(1)
_ = W2731
tmp4204 := MakeNative(func(__e *ControlFlow) {
W2732 := __e.Get(1)
_ = W2732
tmp4205 := PrimCons(W2728, W2731)

__e.TailApply(PrimFunc(symshen_4comb), W2732, tmp4205)
return


}, 1)

tmp4206 := Call(__e, PrimFunc(symshen_4in_1_6), W2730)


__e.TailApply(tmp4204, tmp4206)
return


}, 1)

tmp4207 := Call(__e, PrimFunc(symshen_4_5_1out), W2730)


__e.TailApply(tmp4203, tmp4207)
return


}


}, 1)

tmp4210 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2729)


__e.TailApply(tmp4202, tmp4210)
return


}, 1)

tmp4211 := Call(__e, PrimFunc(symshen_4in_1_6), W2727)


__e.TailApply(tmp4201, tmp4211)
return


}, 1)

tmp4212 := Call(__e, PrimFunc(symshen_4_5_1out), W2727)


__e.TailApply(tmp4200, tmp4212)
return


}


}, 1)

tmp4215 := Call(__e, PrimFunc(symshen_4_5atom_6), V2653)


tmp4216 := Call(__e, tmp4199, tmp4215)


__e.TailApply(tmp4168, tmp4216)
return


} else {
__e.Return(W2720)
return
}


}, 1)

tmp4219 := MakeNative(func(__e *ControlFlow) {
W2721 := __e.Get(1)
_ = W2721
tmp4231 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2721)


if True == tmp4231 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4220 := MakeNative(func(__e *ControlFlow) {
W2722 := __e.Get(1)
_ = W2722
tmp4221 := MakeNative(func(__e *ControlFlow) {
W2723 := __e.Get(1)
_ = W2723
tmp4227 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2723)


if True == tmp4227 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4222 := MakeNative(func(__e *ControlFlow) {
W2724 := __e.Get(1)
_ = W2724
tmp4223 := MakeNative(func(__e *ControlFlow) {
W2725 := __e.Get(1)
_ = W2725
__e.TailApply(PrimFunc(symshen_4comb), W2725, W2724)
return
}, 1)

tmp4224 := Call(__e, PrimFunc(symshen_4in_1_6), W2723)


__e.TailApply(tmp4223, tmp4224)
return


}, 1)

tmp4225 := Call(__e, PrimFunc(symshen_4_5_1out), W2723)


__e.TailApply(tmp4222, tmp4225)
return


}


}, 1)

tmp4228 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2722)


__e.TailApply(tmp4221, tmp4228)
return


}, 1)

tmp4229 := Call(__e, PrimFunc(symshen_4in_1_6), W2721)


__e.TailApply(tmp4220, tmp4229)
return


}


}, 1)

tmp4232 := Call(__e, PrimFunc(symshen_4_5comment_6), V2653)


tmp4233 := Call(__e, tmp4219, tmp4232)


__e.TailApply(tmp4167, tmp4233)
return


} else {
__e.Return(W2714)
return
}


}, 1)

tmp4236 := MakeNative(func(__e *ControlFlow) {
W2715 := __e.Get(1)
_ = W2715
tmp4250 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2715)


if True == tmp4250 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4237 := MakeNative(func(__e *ControlFlow) {
W2716 := __e.Get(1)
_ = W2716
tmp4238 := MakeNative(func(__e *ControlFlow) {
W2717 := __e.Get(1)
_ = W2717
tmp4246 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2717)


if True == tmp4246 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4239 := MakeNative(func(__e *ControlFlow) {
W2718 := __e.Get(1)
_ = W2718
tmp4240 := MakeNative(func(__e *ControlFlow) {
W2719 := __e.Get(1)
_ = W2719
tmp4241 := PrimIntern(MakeString(","))

tmp4242 := PrimCons(tmp4241, W2718)

__e.TailApply(PrimFunc(symshen_4comb), W2719, tmp4242)
return


}, 1)

tmp4243 := Call(__e, PrimFunc(symshen_4in_1_6), W2717)


__e.TailApply(tmp4240, tmp4243)
return


}, 1)

tmp4244 := Call(__e, PrimFunc(symshen_4_5_1out), W2717)


__e.TailApply(tmp4239, tmp4244)
return


}


}, 1)

tmp4247 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2716)


__e.TailApply(tmp4238, tmp4247)
return


}, 1)

tmp4248 := Call(__e, PrimFunc(symshen_4in_1_6), W2715)


__e.TailApply(tmp4237, tmp4248)
return


}


}, 1)

tmp4251 := Call(__e, PrimFunc(symshen_4_5comma_6), V2653)


tmp4252 := Call(__e, tmp4236, tmp4251)


__e.TailApply(tmp4166, tmp4252)
return


} else {
__e.Return(W2708)
return
}


}, 1)

tmp4255 := MakeNative(func(__e *ControlFlow) {
W2709 := __e.Get(1)
_ = W2709
tmp4269 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2709)


if True == tmp4269 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4256 := MakeNative(func(__e *ControlFlow) {
W2710 := __e.Get(1)
_ = W2710
tmp4257 := MakeNative(func(__e *ControlFlow) {
W2711 := __e.Get(1)
_ = W2711
tmp4265 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2711)


if True == tmp4265 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4258 := MakeNative(func(__e *ControlFlow) {
W2712 := __e.Get(1)
_ = W2712
tmp4259 := MakeNative(func(__e *ControlFlow) {
W2713 := __e.Get(1)
_ = W2713
tmp4260 := PrimIntern(MakeString(":"))

tmp4261 := PrimCons(tmp4260, W2712)

__e.TailApply(PrimFunc(symshen_4comb), W2713, tmp4261)
return


}, 1)

tmp4262 := Call(__e, PrimFunc(symshen_4in_1_6), W2711)


__e.TailApply(tmp4259, tmp4262)
return


}, 1)

tmp4263 := Call(__e, PrimFunc(symshen_4_5_1out), W2711)


__e.TailApply(tmp4258, tmp4263)
return


}


}, 1)

tmp4266 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2710)


__e.TailApply(tmp4257, tmp4266)
return


}, 1)

tmp4267 := Call(__e, PrimFunc(symshen_4in_1_6), W2709)


__e.TailApply(tmp4256, tmp4267)
return


}


}, 1)

tmp4270 := Call(__e, PrimFunc(symshen_4_5colon_6), V2653)


tmp4271 := Call(__e, tmp4255, tmp4270)


__e.TailApply(tmp4165, tmp4271)
return


} else {
__e.Return(W2700)
return
}


}, 1)

tmp4274 := MakeNative(func(__e *ControlFlow) {
W2701 := __e.Get(1)
_ = W2701
tmp4294 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2701)


if True == tmp4294 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4275 := MakeNative(func(__e *ControlFlow) {
W2702 := __e.Get(1)
_ = W2702
tmp4276 := MakeNative(func(__e *ControlFlow) {
W2703 := __e.Get(1)
_ = W2703
tmp4290 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2703)


if True == tmp4290 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4277 := MakeNative(func(__e *ControlFlow) {
W2704 := __e.Get(1)
_ = W2704
tmp4278 := MakeNative(func(__e *ControlFlow) {
W2705 := __e.Get(1)
_ = W2705
tmp4286 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2705)


if True == tmp4286 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4279 := MakeNative(func(__e *ControlFlow) {
W2706 := __e.Get(1)
_ = W2706
tmp4280 := MakeNative(func(__e *ControlFlow) {
W2707 := __e.Get(1)
_ = W2707
tmp4281 := PrimIntern(MakeString(":="))

tmp4282 := PrimCons(tmp4281, W2706)

__e.TailApply(PrimFunc(symshen_4comb), W2707, tmp4282)
return


}, 1)

tmp4283 := Call(__e, PrimFunc(symshen_4in_1_6), W2705)


__e.TailApply(tmp4280, tmp4283)
return


}, 1)

tmp4284 := Call(__e, PrimFunc(symshen_4_5_1out), W2705)


__e.TailApply(tmp4279, tmp4284)
return


}


}, 1)

tmp4287 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2704)


__e.TailApply(tmp4278, tmp4287)
return


}, 1)

tmp4288 := Call(__e, PrimFunc(symshen_4in_1_6), W2703)


__e.TailApply(tmp4277, tmp4288)
return


}


}, 1)

tmp4291 := Call(__e, PrimFunc(symshen_4_5equal_6), W2702)


__e.TailApply(tmp4276, tmp4291)
return


}, 1)

tmp4292 := Call(__e, PrimFunc(symshen_4in_1_6), W2701)


__e.TailApply(tmp4275, tmp4292)
return


}


}, 1)

tmp4295 := Call(__e, PrimFunc(symshen_4_5colon_6), V2653)


tmp4296 := Call(__e, tmp4274, tmp4295)


__e.TailApply(tmp4164, tmp4296)
return


} else {
__e.Return(W2694)
return
}


}, 1)

tmp4299 := MakeNative(func(__e *ControlFlow) {
W2695 := __e.Get(1)
_ = W2695
tmp4313 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2695)


if True == tmp4313 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4300 := MakeNative(func(__e *ControlFlow) {
W2696 := __e.Get(1)
_ = W2696
tmp4301 := MakeNative(func(__e *ControlFlow) {
W2697 := __e.Get(1)
_ = W2697
tmp4309 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2697)


if True == tmp4309 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4302 := MakeNative(func(__e *ControlFlow) {
W2698 := __e.Get(1)
_ = W2698
tmp4303 := MakeNative(func(__e *ControlFlow) {
W2699 := __e.Get(1)
_ = W2699
tmp4304 := PrimIntern(MakeString(";"))

tmp4305 := PrimCons(tmp4304, W2698)

__e.TailApply(PrimFunc(symshen_4comb), W2699, tmp4305)
return


}, 1)

tmp4306 := Call(__e, PrimFunc(symshen_4in_1_6), W2697)


__e.TailApply(tmp4303, tmp4306)
return


}, 1)

tmp4307 := Call(__e, PrimFunc(symshen_4_5_1out), W2697)


__e.TailApply(tmp4302, tmp4307)
return


}


}, 1)

tmp4310 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2696)


__e.TailApply(tmp4301, tmp4310)
return


}, 1)

tmp4311 := Call(__e, PrimFunc(symshen_4in_1_6), W2695)


__e.TailApply(tmp4300, tmp4311)
return


}


}, 1)

tmp4314 := Call(__e, PrimFunc(symshen_4_5semicolon_6), V2653)


tmp4315 := Call(__e, tmp4299, tmp4314)


__e.TailApply(tmp4163, tmp4315)
return


} else {
__e.Return(W2688)
return
}


}, 1)

tmp4318 := MakeNative(func(__e *ControlFlow) {
W2689 := __e.Get(1)
_ = W2689
tmp4331 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2689)


if True == tmp4331 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4319 := MakeNative(func(__e *ControlFlow) {
W2690 := __e.Get(1)
_ = W2690
tmp4320 := MakeNative(func(__e *ControlFlow) {
W2691 := __e.Get(1)
_ = W2691
tmp4327 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2691)


if True == tmp4327 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4321 := MakeNative(func(__e *ControlFlow) {
W2692 := __e.Get(1)
_ = W2692
tmp4322 := MakeNative(func(__e *ControlFlow) {
W2693 := __e.Get(1)
_ = W2693
tmp4323 := PrimCons(symbar_b, W2692)

__e.TailApply(PrimFunc(symshen_4comb), W2693, tmp4323)
return


}, 1)

tmp4324 := Call(__e, PrimFunc(symshen_4in_1_6), W2691)


__e.TailApply(tmp4322, tmp4324)
return


}, 1)

tmp4325 := Call(__e, PrimFunc(symshen_4_5_1out), W2691)


__e.TailApply(tmp4321, tmp4325)
return


}


}, 1)

tmp4328 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2690)


__e.TailApply(tmp4320, tmp4328)
return


}, 1)

tmp4329 := Call(__e, PrimFunc(symshen_4in_1_6), W2689)


__e.TailApply(tmp4319, tmp4329)
return


}


}, 1)

tmp4332 := Call(__e, PrimFunc(symshen_4_5bar_6), V2653)


tmp4333 := Call(__e, tmp4318, tmp4332)


__e.TailApply(tmp4162, tmp4333)
return


} else {
__e.Return(W2682)
return
}


}, 1)

tmp4336 := MakeNative(func(__e *ControlFlow) {
W2683 := __e.Get(1)
_ = W2683
tmp4349 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2683)


if True == tmp4349 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4337 := MakeNative(func(__e *ControlFlow) {
W2684 := __e.Get(1)
_ = W2684
tmp4338 := MakeNative(func(__e *ControlFlow) {
W2685 := __e.Get(1)
_ = W2685
tmp4345 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2685)


if True == tmp4345 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4339 := MakeNative(func(__e *ControlFlow) {
W2686 := __e.Get(1)
_ = W2686
tmp4340 := MakeNative(func(__e *ControlFlow) {
W2687 := __e.Get(1)
_ = W2687
tmp4341 := PrimCons(sym_j, W2686)

__e.TailApply(PrimFunc(symshen_4comb), W2687, tmp4341)
return


}, 1)

tmp4342 := Call(__e, PrimFunc(symshen_4in_1_6), W2685)


__e.TailApply(tmp4340, tmp4342)
return


}, 1)

tmp4343 := Call(__e, PrimFunc(symshen_4_5_1out), W2685)


__e.TailApply(tmp4339, tmp4343)
return


}


}, 1)

tmp4346 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2684)


__e.TailApply(tmp4338, tmp4346)
return


}, 1)

tmp4347 := Call(__e, PrimFunc(symshen_4in_1_6), W2683)


__e.TailApply(tmp4337, tmp4347)
return


}


}, 1)

tmp4350 := Call(__e, PrimFunc(symshen_4_5rcurly_6), V2653)


tmp4351 := Call(__e, tmp4336, tmp4350)


__e.TailApply(tmp4161, tmp4351)
return


} else {
__e.Return(W2676)
return
}


}, 1)

tmp4354 := MakeNative(func(__e *ControlFlow) {
W2677 := __e.Get(1)
_ = W2677
tmp4367 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2677)


if True == tmp4367 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4355 := MakeNative(func(__e *ControlFlow) {
W2678 := __e.Get(1)
_ = W2678
tmp4356 := MakeNative(func(__e *ControlFlow) {
W2679 := __e.Get(1)
_ = W2679
tmp4363 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2679)


if True == tmp4363 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4357 := MakeNative(func(__e *ControlFlow) {
W2680 := __e.Get(1)
_ = W2680
tmp4358 := MakeNative(func(__e *ControlFlow) {
W2681 := __e.Get(1)
_ = W2681
tmp4359 := PrimCons(sym_i, W2680)

__e.TailApply(PrimFunc(symshen_4comb), W2681, tmp4359)
return


}, 1)

tmp4360 := Call(__e, PrimFunc(symshen_4in_1_6), W2679)


__e.TailApply(tmp4358, tmp4360)
return


}, 1)

tmp4361 := Call(__e, PrimFunc(symshen_4_5_1out), W2679)


__e.TailApply(tmp4357, tmp4361)
return


}


}, 1)

tmp4364 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2678)


__e.TailApply(tmp4356, tmp4364)
return


}, 1)

tmp4365 := Call(__e, PrimFunc(symshen_4in_1_6), W2677)


__e.TailApply(tmp4355, tmp4365)
return


}


}, 1)

tmp4368 := Call(__e, PrimFunc(symshen_4_5lcurly_6), V2653)


tmp4369 := Call(__e, tmp4354, tmp4368)


__e.TailApply(tmp4160, tmp4369)
return


} else {
__e.Return(W2665)
return
}


}, 1)

tmp4372 := MakeNative(func(__e *ControlFlow) {
W2666 := __e.Get(1)
_ = W2666
tmp4399 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2666)


if True == tmp4399 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4373 := MakeNative(func(__e *ControlFlow) {
W2667 := __e.Get(1)
_ = W2667
tmp4374 := MakeNative(func(__e *ControlFlow) {
W2668 := __e.Get(1)
_ = W2668
tmp4395 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2668)


if True == tmp4395 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4375 := MakeNative(func(__e *ControlFlow) {
W2669 := __e.Get(1)
_ = W2669
tmp4376 := MakeNative(func(__e *ControlFlow) {
W2670 := __e.Get(1)
_ = W2670
tmp4377 := MakeNative(func(__e *ControlFlow) {
W2671 := __e.Get(1)
_ = W2671
tmp4390 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2671)


if True == tmp4390 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4378 := MakeNative(func(__e *ControlFlow) {
W2672 := __e.Get(1)
_ = W2672
tmp4379 := MakeNative(func(__e *ControlFlow) {
W2673 := __e.Get(1)
_ = W2673
tmp4386 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2673)


if True == tmp4386 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4380 := MakeNative(func(__e *ControlFlow) {
W2674 := __e.Get(1)
_ = W2674
tmp4381 := MakeNative(func(__e *ControlFlow) {
W2675 := __e.Get(1)
_ = W2675
tmp4382 := Call(__e, PrimFunc(symshen_4add_1sexpr), W2669, W2674)


__e.TailApply(PrimFunc(symshen_4comb), W2675, tmp4382)
return


}, 1)

tmp4383 := Call(__e, PrimFunc(symshen_4in_1_6), W2673)


__e.TailApply(tmp4381, tmp4383)
return


}, 1)

tmp4384 := Call(__e, PrimFunc(symshen_4_5_1out), W2673)


__e.TailApply(tmp4380, tmp4384)
return


}


}, 1)

tmp4387 := Call(__e, PrimFunc(symshen_4_5s_1exprs2_6), W2672)


__e.TailApply(tmp4379, tmp4387)
return


}, 1)

tmp4388 := Call(__e, PrimFunc(symshen_4in_1_6), W2671)


__e.TailApply(tmp4378, tmp4388)
return


}


}, 1)

tmp4391 := Call(__e, PrimFunc(symshen_4_5rrb_6), W2670)


__e.TailApply(tmp4377, tmp4391)
return


}, 1)

tmp4392 := Call(__e, PrimFunc(symshen_4in_1_6), W2668)


__e.TailApply(tmp4376, tmp4392)
return


}, 1)

tmp4393 := Call(__e, PrimFunc(symshen_4_5_1out), W2668)


__e.TailApply(tmp4375, tmp4393)
return


}


}, 1)

tmp4396 := Call(__e, PrimFunc(symshen_4_5s_1exprs1_6), W2667)


__e.TailApply(tmp4374, tmp4396)
return


}, 1)

tmp4397 := Call(__e, PrimFunc(symshen_4in_1_6), W2666)


__e.TailApply(tmp4373, tmp4397)
return


}


}, 1)

tmp4400 := Call(__e, PrimFunc(symshen_4_5lrb_6), V2653)


tmp4401 := Call(__e, tmp4372, tmp4400)


__e.TailApply(tmp4159, tmp4401)
return


} else {
__e.Return(W2654)
return
}


}, 1)

tmp4404 := MakeNative(func(__e *ControlFlow) {
W2655 := __e.Get(1)
_ = W2655
tmp4432 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2655)


if True == tmp4432 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4405 := MakeNative(func(__e *ControlFlow) {
W2656 := __e.Get(1)
_ = W2656
tmp4406 := MakeNative(func(__e *ControlFlow) {
W2657 := __e.Get(1)
_ = W2657
tmp4428 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2657)


if True == tmp4428 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4407 := MakeNative(func(__e *ControlFlow) {
W2658 := __e.Get(1)
_ = W2658
tmp4408 := MakeNative(func(__e *ControlFlow) {
W2659 := __e.Get(1)
_ = W2659
tmp4409 := MakeNative(func(__e *ControlFlow) {
W2660 := __e.Get(1)
_ = W2660
tmp4423 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2660)


if True == tmp4423 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4410 := MakeNative(func(__e *ControlFlow) {
W2661 := __e.Get(1)
_ = W2661
tmp4411 := MakeNative(func(__e *ControlFlow) {
W2662 := __e.Get(1)
_ = W2662
tmp4419 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2662)


if True == tmp4419 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4412 := MakeNative(func(__e *ControlFlow) {
W2663 := __e.Get(1)
_ = W2663
tmp4413 := MakeNative(func(__e *ControlFlow) {
W2664 := __e.Get(1)
_ = W2664
tmp4414 := Call(__e, PrimFunc(symshen_4cons_1form), W2658)


tmp4415 := PrimCons(tmp4414, W2663)

__e.TailApply(PrimFunc(symshen_4comb), W2664, tmp4415)
return


}, 1)

tmp4416 := Call(__e, PrimFunc(symshen_4in_1_6), W2662)


__e.TailApply(tmp4413, tmp4416)
return


}, 1)

tmp4417 := Call(__e, PrimFunc(symshen_4_5_1out), W2662)


__e.TailApply(tmp4412, tmp4417)
return


}


}, 1)

tmp4420 := Call(__e, PrimFunc(symshen_4_5s_1exprs2_6), W2661)


__e.TailApply(tmp4411, tmp4420)
return


}, 1)

tmp4421 := Call(__e, PrimFunc(symshen_4in_1_6), W2660)


__e.TailApply(tmp4410, tmp4421)
return


}


}, 1)

tmp4424 := Call(__e, PrimFunc(symshen_4_5rsb_6), W2659)


__e.TailApply(tmp4409, tmp4424)
return


}, 1)

tmp4425 := Call(__e, PrimFunc(symshen_4in_1_6), W2657)


__e.TailApply(tmp4408, tmp4425)
return


}, 1)

tmp4426 := Call(__e, PrimFunc(symshen_4_5_1out), W2657)


__e.TailApply(tmp4407, tmp4426)
return


}


}, 1)

tmp4429 := Call(__e, PrimFunc(symshen_4_5s_1exprs1_6), W2656)


__e.TailApply(tmp4406, tmp4429)
return


}, 1)

tmp4430 := Call(__e, PrimFunc(symshen_4in_1_6), W2655)


__e.TailApply(tmp4405, tmp4430)
return


}


}, 1)

tmp4433 := Call(__e, PrimFunc(symshen_4_5lsb_6), V2653)


tmp4434 := Call(__e, tmp4404, tmp4433)


__e.TailApply(tmp4158, tmp4434)
return


}, 1)

tmp4435 := Call(__e, ns2_1set, symshen_4_5s_1exprs_6, tmp4157)


_ = tmp4435

tmp4436 := MakeNative(func(__e *ControlFlow) {
V2742 := __e.Get(1)
_ = V2742
V2743 := __e.Get(2)
_ = V2743
tmp4454 := PrimIsPair(V2742)

var ifres4441 Obj

if True == tmp4454 {
tmp4452 := PrimHead(V2742)

tmp4453 := PrimEqual(sym_3, tmp4452)

var ifres4443 Obj

if True == tmp4453 {
tmp4450 := PrimTail(V2742)

tmp4451 := PrimIsPair(tmp4450)

var ifres4445 Obj

if True == tmp4451 {
tmp4447 := PrimTail(V2742)

tmp4448 := PrimTail(tmp4447)

tmp4449 := PrimEqual(Nil, tmp4448)

var ifres4446 Obj

if True == tmp4449 {
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
tmp4437 := PrimTail(V2742)

tmp4438 := PrimHead(tmp4437)

tmp4439 := Call(__e, PrimFunc(symexplode), tmp4438)


__e.TailApply(PrimFunc(symappend), tmp4439, V2743)
return


} else {
__e.Return(PrimCons(V2742, V2743))
return
}


}, 2)

tmp4455 := Call(__e, ns2_1set, symshen_4add_1sexpr, tmp4436)


_ = tmp4455

tmp4456 := MakeNative(func(__e *ControlFlow) {
V2744 := __e.Get(1)
_ = V2744
tmp4457 := MakeNative(func(__e *ControlFlow) {
W2745 := __e.Get(1)
_ = W2745
tmp4459 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2745)


if True == tmp4459 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2745)
return
}


}, 1)

tmp4465 := Call(__e, PrimFunc(symshen_4hds_a_2), V2744, MakeNumber(91))


var ifres4460 Obj

if True == tmp4465 {
tmp4461 := MakeNative(func(__e *ControlFlow) {
W2746 := __e.Get(1)
_ = W2746
__e.TailApply(PrimFunc(symshen_4comb), W2746, symshen_4skip)
return
}, 1)

tmp4462 := Call(__e, PrimFunc(symtail), V2744)


tmp4463 := Call(__e, tmp4461, tmp4462)


ifres4460 = tmp4463


} else {
tmp4464 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres4460 = tmp4464


}

__e.TailApply(tmp4457, ifres4460)
return


}, 1)

tmp4466 := Call(__e, ns2_1set, symshen_4_5lsb_6, tmp4456)


_ = tmp4466

tmp4467 := MakeNative(func(__e *ControlFlow) {
V2747 := __e.Get(1)
_ = V2747
tmp4468 := MakeNative(func(__e *ControlFlow) {
W2748 := __e.Get(1)
_ = W2748
tmp4470 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2748)


if True == tmp4470 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2748)
return
}


}, 1)

tmp4476 := Call(__e, PrimFunc(symshen_4hds_a_2), V2747, MakeNumber(93))


var ifres4471 Obj

if True == tmp4476 {
tmp4472 := MakeNative(func(__e *ControlFlow) {
W2749 := __e.Get(1)
_ = W2749
__e.TailApply(PrimFunc(symshen_4comb), W2749, symshen_4skip)
return
}, 1)

tmp4473 := Call(__e, PrimFunc(symtail), V2747)


tmp4474 := Call(__e, tmp4472, tmp4473)


ifres4471 = tmp4474


} else {
tmp4475 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres4471 = tmp4475


}

__e.TailApply(tmp4468, ifres4471)
return


}, 1)

tmp4477 := Call(__e, ns2_1set, symshen_4_5rsb_6, tmp4467)


_ = tmp4477

tmp4478 := MakeNative(func(__e *ControlFlow) {
V2750 := __e.Get(1)
_ = V2750
tmp4479 := MakeNative(func(__e *ControlFlow) {
W2751 := __e.Get(1)
_ = W2751
tmp4481 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2751)


if True == tmp4481 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2751)
return
}


}, 1)

tmp4482 := MakeNative(func(__e *ControlFlow) {
W2752 := __e.Get(1)
_ = W2752
tmp4488 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2752)


if True == tmp4488 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4483 := MakeNative(func(__e *ControlFlow) {
W2753 := __e.Get(1)
_ = W2753
tmp4484 := MakeNative(func(__e *ControlFlow) {
W2754 := __e.Get(1)
_ = W2754
__e.TailApply(PrimFunc(symshen_4comb), W2754, W2753)
return
}, 1)

tmp4485 := Call(__e, PrimFunc(symshen_4in_1_6), W2752)


__e.TailApply(tmp4484, tmp4485)
return


}, 1)

tmp4486 := Call(__e, PrimFunc(symshen_4_5_1out), W2752)


__e.TailApply(tmp4483, tmp4486)
return


}


}, 1)

tmp4489 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), V2750)


tmp4490 := Call(__e, tmp4482, tmp4489)


__e.TailApply(tmp4479, tmp4490)
return


}, 1)

tmp4491 := Call(__e, ns2_1set, symshen_4_5s_1exprs1_6, tmp4478)


_ = tmp4491

tmp4492 := MakeNative(func(__e *ControlFlow) {
V2755 := __e.Get(1)
_ = V2755
tmp4493 := MakeNative(func(__e *ControlFlow) {
W2756 := __e.Get(1)
_ = W2756
tmp4495 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2756)


if True == tmp4495 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2756)
return
}


}, 1)

tmp4496 := MakeNative(func(__e *ControlFlow) {
W2757 := __e.Get(1)
_ = W2757
tmp4502 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2757)


if True == tmp4502 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4497 := MakeNative(func(__e *ControlFlow) {
W2758 := __e.Get(1)
_ = W2758
tmp4498 := MakeNative(func(__e *ControlFlow) {
W2759 := __e.Get(1)
_ = W2759
__e.TailApply(PrimFunc(symshen_4comb), W2759, W2758)
return
}, 1)

tmp4499 := Call(__e, PrimFunc(symshen_4in_1_6), W2757)


__e.TailApply(tmp4498, tmp4499)
return


}, 1)

tmp4500 := Call(__e, PrimFunc(symshen_4_5_1out), W2757)


__e.TailApply(tmp4497, tmp4500)
return


}


}, 1)

tmp4503 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), V2755)


tmp4504 := Call(__e, tmp4496, tmp4503)


__e.TailApply(tmp4493, tmp4504)
return


}, 1)

tmp4505 := Call(__e, ns2_1set, symshen_4_5s_1exprs2_6, tmp4492)


_ = tmp4505

tmp4506 := MakeNative(func(__e *ControlFlow) {
V2761 := __e.Get(1)
_ = V2761
tmp4563 := PrimEqual(Nil, V2761)

if True == tmp4563 {
__e.Return(Nil)
return
} else {
tmp4561 := PrimIsPair(V2761)

var ifres4541 Obj

if True == tmp4561 {
tmp4559 := PrimTail(V2761)

tmp4560 := PrimIsPair(tmp4559)

var ifres4543 Obj

if True == tmp4560 {
tmp4556 := PrimTail(V2761)

tmp4557 := PrimTail(tmp4556)

tmp4558 := PrimIsPair(tmp4557)

var ifres4545 Obj

if True == tmp4558 {
tmp4552 := PrimTail(V2761)

tmp4553 := PrimTail(tmp4552)

tmp4554 := PrimTail(tmp4553)

tmp4555 := PrimEqual(Nil, tmp4554)

var ifres4547 Obj

if True == tmp4555 {
tmp4549 := PrimTail(V2761)

tmp4550 := PrimHead(tmp4549)

tmp4551 := PrimEqual(tmp4550, symbar_b)

var ifres4548 Obj

if True == tmp4551 {
ifres4548 = True


} else {
ifres4548 = False


}

ifres4547 = ifres4548


} else {
ifres4547 = False


}

var ifres4546 Obj

if True == ifres4547 {
ifres4546 = True


} else {
ifres4546 = False


}

ifres4545 = ifres4546


} else {
ifres4545 = False


}

var ifres4544 Obj

if True == ifres4545 {
ifres4544 = True


} else {
ifres4544 = False


}

ifres4543 = ifres4544


} else {
ifres4543 = False


}

var ifres4542 Obj

if True == ifres4543 {
ifres4542 = True


} else {
ifres4542 = False


}

ifres4541 = ifres4542


} else {
ifres4541 = False


}

if True == ifres4541 {
tmp4507 := PrimHead(V2761)

tmp4508 := PrimTail(V2761)

tmp4509 := PrimTail(tmp4508)

tmp4510 := PrimCons(tmp4507, tmp4509)

__e.Return(PrimCons(symcons, tmp4510))
return


} else {
tmp4539 := PrimIsPair(V2761)

var ifres4519 Obj

if True == tmp4539 {
tmp4537 := PrimTail(V2761)

tmp4538 := PrimIsPair(tmp4537)

var ifres4521 Obj

if True == tmp4538 {
tmp4534 := PrimTail(V2761)

tmp4535 := PrimTail(tmp4534)

tmp4536 := PrimIsPair(tmp4535)

var ifres4523 Obj

if True == tmp4536 {
tmp4530 := PrimTail(V2761)

tmp4531 := PrimTail(tmp4530)

tmp4532 := PrimTail(tmp4531)

tmp4533 := PrimIsPair(tmp4532)

var ifres4525 Obj

if True == tmp4533 {
tmp4527 := PrimTail(V2761)

tmp4528 := PrimHead(tmp4527)

tmp4529 := PrimEqual(tmp4528, symbar_b)

var ifres4526 Obj

if True == tmp4529 {
ifres4526 = True


} else {
ifres4526 = False


}

ifres4525 = ifres4526


} else {
ifres4525 = False


}

var ifres4524 Obj

if True == ifres4525 {
ifres4524 = True


} else {
ifres4524 = False


}

ifres4523 = ifres4524


} else {
ifres4523 = False


}

var ifres4522 Obj

if True == ifres4523 {
ifres4522 = True


} else {
ifres4522 = False


}

ifres4521 = ifres4522


} else {
ifres4521 = False


}

var ifres4520 Obj

if True == ifres4521 {
ifres4520 = True


} else {
ifres4520 = False


}

ifres4519 = ifres4520


} else {
ifres4519 = False


}

if True == ifres4519 {
__e.Return(PrimSimpleError(MakeString("misapplication of |\n")))
return
} else {
tmp4517 := PrimIsPair(V2761)

if True == tmp4517 {
tmp4511 := PrimHead(V2761)

tmp4512 := PrimTail(V2761)

tmp4513 := Call(__e, PrimFunc(symshen_4cons_1form), tmp4512)


tmp4514 := PrimCons(tmp4513, Nil)

tmp4515 := PrimCons(tmp4511, tmp4514)

__e.Return(PrimCons(symcons, tmp4515))
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4cons_1form)
return
}


}


}


}


}, 1)

tmp4564 := Call(__e, ns2_1set, symshen_4cons_1form, tmp4506)


_ = tmp4564

tmp4565 := MakeNative(func(__e *ControlFlow) {
V2762 := __e.Get(1)
_ = V2762
tmp4566 := MakeNative(func(__e *ControlFlow) {
W2763 := __e.Get(1)
_ = W2763
tmp4568 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2763)


if True == tmp4568 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2763)
return
}


}, 1)

tmp4574 := Call(__e, PrimFunc(symshen_4hds_a_2), V2762, MakeNumber(40))


var ifres4569 Obj

if True == tmp4574 {
tmp4570 := MakeNative(func(__e *ControlFlow) {
W2764 := __e.Get(1)
_ = W2764
__e.TailApply(PrimFunc(symshen_4comb), W2764, symshen_4skip)
return
}, 1)

tmp4571 := Call(__e, PrimFunc(symtail), V2762)


tmp4572 := Call(__e, tmp4570, tmp4571)


ifres4569 = tmp4572


} else {
tmp4573 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres4569 = tmp4573


}

__e.TailApply(tmp4566, ifres4569)
return


}, 1)

tmp4575 := Call(__e, ns2_1set, symshen_4_5lrb_6, tmp4565)


_ = tmp4575

tmp4576 := MakeNative(func(__e *ControlFlow) {
V2765 := __e.Get(1)
_ = V2765
tmp4577 := MakeNative(func(__e *ControlFlow) {
W2766 := __e.Get(1)
_ = W2766
tmp4579 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2766)


if True == tmp4579 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2766)
return
}


}, 1)

tmp4585 := Call(__e, PrimFunc(symshen_4hds_a_2), V2765, MakeNumber(41))


var ifres4580 Obj

if True == tmp4585 {
tmp4581 := MakeNative(func(__e *ControlFlow) {
W2767 := __e.Get(1)
_ = W2767
__e.TailApply(PrimFunc(symshen_4comb), W2767, symshen_4skip)
return
}, 1)

tmp4582 := Call(__e, PrimFunc(symtail), V2765)


tmp4583 := Call(__e, tmp4581, tmp4582)


ifres4580 = tmp4583


} else {
tmp4584 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres4580 = tmp4584


}

__e.TailApply(tmp4577, ifres4580)
return


}, 1)

tmp4586 := Call(__e, ns2_1set, symshen_4_5rrb_6, tmp4576)


_ = tmp4586

tmp4587 := MakeNative(func(__e *ControlFlow) {
V2768 := __e.Get(1)
_ = V2768
tmp4588 := MakeNative(func(__e *ControlFlow) {
W2769 := __e.Get(1)
_ = W2769
tmp4590 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2769)


if True == tmp4590 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2769)
return
}


}, 1)

tmp4596 := Call(__e, PrimFunc(symshen_4hds_a_2), V2768, MakeNumber(123))


var ifres4591 Obj

if True == tmp4596 {
tmp4592 := MakeNative(func(__e *ControlFlow) {
W2770 := __e.Get(1)
_ = W2770
__e.TailApply(PrimFunc(symshen_4comb), W2770, symshen_4skip)
return
}, 1)

tmp4593 := Call(__e, PrimFunc(symtail), V2768)


tmp4594 := Call(__e, tmp4592, tmp4593)


ifres4591 = tmp4594


} else {
tmp4595 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres4591 = tmp4595


}

__e.TailApply(tmp4588, ifres4591)
return


}, 1)

tmp4597 := Call(__e, ns2_1set, symshen_4_5lcurly_6, tmp4587)


_ = tmp4597

tmp4598 := MakeNative(func(__e *ControlFlow) {
V2771 := __e.Get(1)
_ = V2771
tmp4599 := MakeNative(func(__e *ControlFlow) {
W2772 := __e.Get(1)
_ = W2772
tmp4601 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2772)


if True == tmp4601 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2772)
return
}


}, 1)

tmp4607 := Call(__e, PrimFunc(symshen_4hds_a_2), V2771, MakeNumber(125))


var ifres4602 Obj

if True == tmp4607 {
tmp4603 := MakeNative(func(__e *ControlFlow) {
W2773 := __e.Get(1)
_ = W2773
__e.TailApply(PrimFunc(symshen_4comb), W2773, symshen_4skip)
return
}, 1)

tmp4604 := Call(__e, PrimFunc(symtail), V2771)


tmp4605 := Call(__e, tmp4603, tmp4604)


ifres4602 = tmp4605


} else {
tmp4606 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres4602 = tmp4606


}

__e.TailApply(tmp4599, ifres4602)
return


}, 1)

tmp4608 := Call(__e, ns2_1set, symshen_4_5rcurly_6, tmp4598)


_ = tmp4608

tmp4609 := MakeNative(func(__e *ControlFlow) {
V2774 := __e.Get(1)
_ = V2774
tmp4610 := MakeNative(func(__e *ControlFlow) {
W2775 := __e.Get(1)
_ = W2775
tmp4612 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2775)


if True == tmp4612 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2775)
return
}


}, 1)

tmp4618 := Call(__e, PrimFunc(symshen_4hds_a_2), V2774, MakeNumber(124))


var ifres4613 Obj

if True == tmp4618 {
tmp4614 := MakeNative(func(__e *ControlFlow) {
W2776 := __e.Get(1)
_ = W2776
__e.TailApply(PrimFunc(symshen_4comb), W2776, symshen_4skip)
return
}, 1)

tmp4615 := Call(__e, PrimFunc(symtail), V2774)


tmp4616 := Call(__e, tmp4614, tmp4615)


ifres4613 = tmp4616


} else {
tmp4617 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres4613 = tmp4617


}

__e.TailApply(tmp4610, ifres4613)
return


}, 1)

tmp4619 := Call(__e, ns2_1set, symshen_4_5bar_6, tmp4609)


_ = tmp4619

tmp4620 := MakeNative(func(__e *ControlFlow) {
V2777 := __e.Get(1)
_ = V2777
tmp4621 := MakeNative(func(__e *ControlFlow) {
W2778 := __e.Get(1)
_ = W2778
tmp4623 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2778)


if True == tmp4623 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2778)
return
}


}, 1)

tmp4629 := Call(__e, PrimFunc(symshen_4hds_a_2), V2777, MakeNumber(59))


var ifres4624 Obj

if True == tmp4629 {
tmp4625 := MakeNative(func(__e *ControlFlow) {
W2779 := __e.Get(1)
_ = W2779
__e.TailApply(PrimFunc(symshen_4comb), W2779, symshen_4skip)
return
}, 1)

tmp4626 := Call(__e, PrimFunc(symtail), V2777)


tmp4627 := Call(__e, tmp4625, tmp4626)


ifres4624 = tmp4627


} else {
tmp4628 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres4624 = tmp4628


}

__e.TailApply(tmp4621, ifres4624)
return


}, 1)

tmp4630 := Call(__e, ns2_1set, symshen_4_5semicolon_6, tmp4620)


_ = tmp4630

tmp4631 := MakeNative(func(__e *ControlFlow) {
V2780 := __e.Get(1)
_ = V2780
tmp4632 := MakeNative(func(__e *ControlFlow) {
W2781 := __e.Get(1)
_ = W2781
tmp4634 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2781)


if True == tmp4634 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2781)
return
}


}, 1)

tmp4640 := Call(__e, PrimFunc(symshen_4hds_a_2), V2780, MakeNumber(58))


var ifres4635 Obj

if True == tmp4640 {
tmp4636 := MakeNative(func(__e *ControlFlow) {
W2782 := __e.Get(1)
_ = W2782
__e.TailApply(PrimFunc(symshen_4comb), W2782, symshen_4skip)
return
}, 1)

tmp4637 := Call(__e, PrimFunc(symtail), V2780)


tmp4638 := Call(__e, tmp4636, tmp4637)


ifres4635 = tmp4638


} else {
tmp4639 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres4635 = tmp4639


}

__e.TailApply(tmp4632, ifres4635)
return


}, 1)

tmp4641 := Call(__e, ns2_1set, symshen_4_5colon_6, tmp4631)


_ = tmp4641

tmp4642 := MakeNative(func(__e *ControlFlow) {
V2783 := __e.Get(1)
_ = V2783
tmp4643 := MakeNative(func(__e *ControlFlow) {
W2784 := __e.Get(1)
_ = W2784
tmp4645 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2784)


if True == tmp4645 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2784)
return
}


}, 1)

tmp4651 := Call(__e, PrimFunc(symshen_4hds_a_2), V2783, MakeNumber(44))


var ifres4646 Obj

if True == tmp4651 {
tmp4647 := MakeNative(func(__e *ControlFlow) {
W2785 := __e.Get(1)
_ = W2785
__e.TailApply(PrimFunc(symshen_4comb), W2785, symshen_4skip)
return
}, 1)

tmp4648 := Call(__e, PrimFunc(symtail), V2783)


tmp4649 := Call(__e, tmp4647, tmp4648)


ifres4646 = tmp4649


} else {
tmp4650 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres4646 = tmp4650


}

__e.TailApply(tmp4643, ifres4646)
return


}, 1)

tmp4652 := Call(__e, ns2_1set, symshen_4_5comma_6, tmp4642)


_ = tmp4652

tmp4653 := MakeNative(func(__e *ControlFlow) {
V2786 := __e.Get(1)
_ = V2786
tmp4654 := MakeNative(func(__e *ControlFlow) {
W2787 := __e.Get(1)
_ = W2787
tmp4656 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2787)


if True == tmp4656 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2787)
return
}


}, 1)

tmp4662 := Call(__e, PrimFunc(symshen_4hds_a_2), V2786, MakeNumber(61))


var ifres4657 Obj

if True == tmp4662 {
tmp4658 := MakeNative(func(__e *ControlFlow) {
W2788 := __e.Get(1)
_ = W2788
__e.TailApply(PrimFunc(symshen_4comb), W2788, symshen_4skip)
return
}, 1)

tmp4659 := Call(__e, PrimFunc(symtail), V2786)


tmp4660 := Call(__e, tmp4658, tmp4659)


ifres4657 = tmp4660


} else {
tmp4661 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres4657 = tmp4661


}

__e.TailApply(tmp4654, ifres4657)
return


}, 1)

tmp4663 := Call(__e, ns2_1set, symshen_4_5equal_6, tmp4653)


_ = tmp4663

tmp4664 := MakeNative(func(__e *ControlFlow) {
V2789 := __e.Get(1)
_ = V2789
tmp4665 := MakeNative(func(__e *ControlFlow) {
W2790 := __e.Get(1)
_ = W2790
tmp4677 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2790)


if True == tmp4677 {
tmp4666 := MakeNative(func(__e *ControlFlow) {
W2793 := __e.Get(1)
_ = W2793
tmp4668 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2793)


if True == tmp4668 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2793)
return
}


}, 1)

tmp4669 := MakeNative(func(__e *ControlFlow) {
W2794 := __e.Get(1)
_ = W2794
tmp4673 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2794)


if True == tmp4673 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4670 := MakeNative(func(__e *ControlFlow) {
W2795 := __e.Get(1)
_ = W2795
__e.TailApply(PrimFunc(symshen_4comb), W2795, symshen_4skip)
return
}, 1)

tmp4671 := Call(__e, PrimFunc(symshen_4in_1_6), W2794)


__e.TailApply(tmp4670, tmp4671)
return


}


}, 1)

tmp4674 := Call(__e, PrimFunc(symshen_4_5multiline_6), V2789)


tmp4675 := Call(__e, tmp4669, tmp4674)


__e.TailApply(tmp4666, tmp4675)
return


} else {
__e.Return(W2790)
return
}


}, 1)

tmp4678 := MakeNative(func(__e *ControlFlow) {
W2791 := __e.Get(1)
_ = W2791
tmp4682 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2791)


if True == tmp4682 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4679 := MakeNative(func(__e *ControlFlow) {
W2792 := __e.Get(1)
_ = W2792
__e.TailApply(PrimFunc(symshen_4comb), W2792, symshen_4skip)
return
}, 1)

tmp4680 := Call(__e, PrimFunc(symshen_4in_1_6), W2791)


__e.TailApply(tmp4679, tmp4680)
return


}


}, 1)

tmp4683 := Call(__e, PrimFunc(symshen_4_5singleline_6), V2789)


tmp4684 := Call(__e, tmp4678, tmp4683)


__e.TailApply(tmp4665, tmp4684)
return


}, 1)

tmp4685 := Call(__e, ns2_1set, symshen_4_5comment_6, tmp4664)


_ = tmp4685

tmp4686 := MakeNative(func(__e *ControlFlow) {
V2796 := __e.Get(1)
_ = V2796
tmp4687 := MakeNative(func(__e *ControlFlow) {
W2797 := __e.Get(1)
_ = W2797
tmp4689 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2797)


if True == tmp4689 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2797)
return
}


}, 1)

tmp4690 := MakeNative(func(__e *ControlFlow) {
W2798 := __e.Get(1)
_ = W2798
tmp4712 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2798)


if True == tmp4712 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4691 := MakeNative(func(__e *ControlFlow) {
W2799 := __e.Get(1)
_ = W2799
tmp4692 := MakeNative(func(__e *ControlFlow) {
W2800 := __e.Get(1)
_ = W2800
tmp4708 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2800)


if True == tmp4708 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4693 := MakeNative(func(__e *ControlFlow) {
W2801 := __e.Get(1)
_ = W2801
tmp4694 := MakeNative(func(__e *ControlFlow) {
W2802 := __e.Get(1)
_ = W2802
tmp4704 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2802)


if True == tmp4704 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4695 := MakeNative(func(__e *ControlFlow) {
W2803 := __e.Get(1)
_ = W2803
tmp4696 := MakeNative(func(__e *ControlFlow) {
W2804 := __e.Get(1)
_ = W2804
tmp4700 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2804)


if True == tmp4700 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4697 := MakeNative(func(__e *ControlFlow) {
W2805 := __e.Get(1)
_ = W2805
__e.TailApply(PrimFunc(symshen_4comb), W2805, symshen_4skip)
return
}, 1)

tmp4698 := Call(__e, PrimFunc(symshen_4in_1_6), W2804)


__e.TailApply(tmp4697, tmp4698)
return


}


}, 1)

tmp4701 := Call(__e, PrimFunc(symshen_4_5returns_6), W2803)


__e.TailApply(tmp4696, tmp4701)
return


}, 1)

tmp4702 := Call(__e, PrimFunc(symshen_4in_1_6), W2802)


__e.TailApply(tmp4695, tmp4702)
return


}


}, 1)

tmp4705 := Call(__e, PrimFunc(symshen_4_5shortnatters_6), W2801)


__e.TailApply(tmp4694, tmp4705)
return


}, 1)

tmp4706 := Call(__e, PrimFunc(symshen_4in_1_6), W2800)


__e.TailApply(tmp4693, tmp4706)
return


}


}, 1)

tmp4709 := Call(__e, PrimFunc(symshen_4_5backslash_6), W2799)


__e.TailApply(tmp4692, tmp4709)
return


}, 1)

tmp4710 := Call(__e, PrimFunc(symshen_4in_1_6), W2798)


__e.TailApply(tmp4691, tmp4710)
return


}


}, 1)

tmp4713 := Call(__e, PrimFunc(symshen_4_5backslash_6), V2796)


tmp4714 := Call(__e, tmp4690, tmp4713)


__e.TailApply(tmp4687, tmp4714)
return


}, 1)

tmp4715 := Call(__e, ns2_1set, symshen_4_5singleline_6, tmp4686)


_ = tmp4715

tmp4716 := MakeNative(func(__e *ControlFlow) {
V2806 := __e.Get(1)
_ = V2806
tmp4717 := MakeNative(func(__e *ControlFlow) {
W2807 := __e.Get(1)
_ = W2807
tmp4719 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2807)


if True == tmp4719 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2807)
return
}


}, 1)

tmp4725 := Call(__e, PrimFunc(symshen_4hds_a_2), V2806, MakeNumber(92))


var ifres4720 Obj

if True == tmp4725 {
tmp4721 := MakeNative(func(__e *ControlFlow) {
W2808 := __e.Get(1)
_ = W2808
__e.TailApply(PrimFunc(symshen_4comb), W2808, symshen_4skip)
return
}, 1)

tmp4722 := Call(__e, PrimFunc(symtail), V2806)


tmp4723 := Call(__e, tmp4721, tmp4722)


ifres4720 = tmp4723


} else {
tmp4724 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres4720 = tmp4724


}

__e.TailApply(tmp4717, ifres4720)
return


}, 1)

tmp4726 := Call(__e, ns2_1set, symshen_4_5backslash_6, tmp4716)


_ = tmp4726

tmp4727 := MakeNative(func(__e *ControlFlow) {
V2809 := __e.Get(1)
_ = V2809
tmp4728 := MakeNative(func(__e *ControlFlow) {
W2810 := __e.Get(1)
_ = W2810
tmp4740 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2810)


if True == tmp4740 {
tmp4729 := MakeNative(func(__e *ControlFlow) {
W2815 := __e.Get(1)
_ = W2815
tmp4731 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2815)


if True == tmp4731 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2815)
return
}


}, 1)

tmp4732 := MakeNative(func(__e *ControlFlow) {
W2816 := __e.Get(1)
_ = W2816
tmp4736 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2816)


if True == tmp4736 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4733 := MakeNative(func(__e *ControlFlow) {
W2817 := __e.Get(1)
_ = W2817
__e.TailApply(PrimFunc(symshen_4comb), W2817, symshen_4skip)
return
}, 1)

tmp4734 := Call(__e, PrimFunc(symshen_4in_1_6), W2816)


__e.TailApply(tmp4733, tmp4734)
return


}


}, 1)

tmp4737 := Call(__e, PrimFunc(sym_5e_6), V2809)


tmp4738 := Call(__e, tmp4732, tmp4737)


__e.TailApply(tmp4729, tmp4738)
return


} else {
__e.Return(W2810)
return
}


}, 1)

tmp4741 := MakeNative(func(__e *ControlFlow) {
W2811 := __e.Get(1)
_ = W2811
tmp4751 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2811)


if True == tmp4751 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4742 := MakeNative(func(__e *ControlFlow) {
W2812 := __e.Get(1)
_ = W2812
tmp4743 := MakeNative(func(__e *ControlFlow) {
W2813 := __e.Get(1)
_ = W2813
tmp4747 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2813)


if True == tmp4747 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4744 := MakeNative(func(__e *ControlFlow) {
W2814 := __e.Get(1)
_ = W2814
__e.TailApply(PrimFunc(symshen_4comb), W2814, symshen_4skip)
return
}, 1)

tmp4745 := Call(__e, PrimFunc(symshen_4in_1_6), W2813)


__e.TailApply(tmp4744, tmp4745)
return


}


}, 1)

tmp4748 := Call(__e, PrimFunc(symshen_4_5shortnatters_6), W2812)


__e.TailApply(tmp4743, tmp4748)
return


}, 1)

tmp4749 := Call(__e, PrimFunc(symshen_4in_1_6), W2811)


__e.TailApply(tmp4742, tmp4749)
return


}


}, 1)

tmp4752 := Call(__e, PrimFunc(symshen_4_5shortnatter_6), V2809)


tmp4753 := Call(__e, tmp4741, tmp4752)


__e.TailApply(tmp4728, tmp4753)
return


}, 1)

tmp4754 := Call(__e, ns2_1set, symshen_4_5shortnatters_6, tmp4727)


_ = tmp4754

tmp4755 := MakeNative(func(__e *ControlFlow) {
V2818 := __e.Get(1)
_ = V2818
tmp4756 := MakeNative(func(__e *ControlFlow) {
W2819 := __e.Get(1)
_ = W2819
tmp4758 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2819)


if True == tmp4758 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2819)
return
}


}, 1)

tmp4769 := PrimIsPair(V2818)

var ifres4759 Obj

if True == tmp4769 {
tmp4760 := MakeNative(func(__e *ControlFlow) {
W2820 := __e.Get(1)
_ = W2820
tmp4761 := MakeNative(func(__e *ControlFlow) {
W2821 := __e.Get(1)
_ = W2821
tmp4763 := Call(__e, PrimFunc(symshen_4return_2), W2820)


tmp4764 := PrimNot(tmp4763)

if True == tmp4764 {
__e.TailApply(PrimFunc(symshen_4comb), W2821, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp4765 := Call(__e, PrimFunc(symtail), V2818)


__e.TailApply(tmp4761, tmp4765)
return


}, 1)

tmp4766 := Call(__e, PrimFunc(symhead), V2818)


tmp4767 := Call(__e, tmp4760, tmp4766)


ifres4759 = tmp4767


} else {
tmp4768 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres4759 = tmp4768


}

__e.TailApply(tmp4756, ifres4759)
return


}, 1)

tmp4770 := Call(__e, ns2_1set, symshen_4_5shortnatter_6, tmp4755)


_ = tmp4770

tmp4771 := MakeNative(func(__e *ControlFlow) {
V2822 := __e.Get(1)
_ = V2822
tmp4772 := MakeNative(func(__e *ControlFlow) {
W2823 := __e.Get(1)
_ = W2823
tmp4784 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2823)


if True == tmp4784 {
tmp4773 := MakeNative(func(__e *ControlFlow) {
W2828 := __e.Get(1)
_ = W2828
tmp4775 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2828)


if True == tmp4775 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2828)
return
}


}, 1)

tmp4776 := MakeNative(func(__e *ControlFlow) {
W2829 := __e.Get(1)
_ = W2829
tmp4780 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2829)


if True == tmp4780 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4777 := MakeNative(func(__e *ControlFlow) {
W2830 := __e.Get(1)
_ = W2830
__e.TailApply(PrimFunc(symshen_4comb), W2830, symshen_4skip)
return
}, 1)

tmp4778 := Call(__e, PrimFunc(symshen_4in_1_6), W2829)


__e.TailApply(tmp4777, tmp4778)
return


}


}, 1)

tmp4781 := Call(__e, PrimFunc(symshen_4_5return_6), V2822)


tmp4782 := Call(__e, tmp4776, tmp4781)


__e.TailApply(tmp4773, tmp4782)
return


} else {
__e.Return(W2823)
return
}


}, 1)

tmp4785 := MakeNative(func(__e *ControlFlow) {
W2824 := __e.Get(1)
_ = W2824
tmp4795 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2824)


if True == tmp4795 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4786 := MakeNative(func(__e *ControlFlow) {
W2825 := __e.Get(1)
_ = W2825
tmp4787 := MakeNative(func(__e *ControlFlow) {
W2826 := __e.Get(1)
_ = W2826
tmp4791 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2826)


if True == tmp4791 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4788 := MakeNative(func(__e *ControlFlow) {
W2827 := __e.Get(1)
_ = W2827
__e.TailApply(PrimFunc(symshen_4comb), W2827, symshen_4skip)
return
}, 1)

tmp4789 := Call(__e, PrimFunc(symshen_4in_1_6), W2826)


__e.TailApply(tmp4788, tmp4789)
return


}


}, 1)

tmp4792 := Call(__e, PrimFunc(symshen_4_5returns_6), W2825)


__e.TailApply(tmp4787, tmp4792)
return


}, 1)

tmp4793 := Call(__e, PrimFunc(symshen_4in_1_6), W2824)


__e.TailApply(tmp4786, tmp4793)
return


}


}, 1)

tmp4796 := Call(__e, PrimFunc(symshen_4_5return_6), V2822)


tmp4797 := Call(__e, tmp4785, tmp4796)


__e.TailApply(tmp4772, tmp4797)
return


}, 1)

tmp4798 := Call(__e, ns2_1set, symshen_4_5returns_6, tmp4771)


_ = tmp4798

tmp4799 := MakeNative(func(__e *ControlFlow) {
V2831 := __e.Get(1)
_ = V2831
tmp4800 := MakeNative(func(__e *ControlFlow) {
W2832 := __e.Get(1)
_ = W2832
tmp4802 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2832)


if True == tmp4802 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2832)
return
}


}, 1)

tmp4812 := PrimIsPair(V2831)

var ifres4803 Obj

if True == tmp4812 {
tmp4804 := MakeNative(func(__e *ControlFlow) {
W2833 := __e.Get(1)
_ = W2833
tmp4805 := MakeNative(func(__e *ControlFlow) {
W2834 := __e.Get(1)
_ = W2834
tmp4807 := Call(__e, PrimFunc(symshen_4return_2), W2833)


if True == tmp4807 {
__e.TailApply(PrimFunc(symshen_4comb), W2834, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp4808 := Call(__e, PrimFunc(symtail), V2831)


__e.TailApply(tmp4805, tmp4808)
return


}, 1)

tmp4809 := Call(__e, PrimFunc(symhead), V2831)


tmp4810 := Call(__e, tmp4804, tmp4809)


ifres4803 = tmp4810


} else {
tmp4811 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres4803 = tmp4811


}

__e.TailApply(tmp4800, ifres4803)
return


}, 1)

tmp4813 := Call(__e, ns2_1set, symshen_4_5return_6, tmp4799)


_ = tmp4813

tmp4814 := MakeNative(func(__e *ControlFlow) {
V2835 := __e.Get(1)
_ = V2835
tmp4815 := PrimCons(MakeNumber(13), Nil)

tmp4816 := PrimCons(MakeNumber(10), tmp4815)

tmp4817 := PrimCons(MakeNumber(9), tmp4816)

__e.TailApply(PrimFunc(symelement_2), V2835, tmp4817)
return


}, 1)

tmp4818 := Call(__e, ns2_1set, symshen_4return_2, tmp4814)


_ = tmp4818

tmp4819 := MakeNative(func(__e *ControlFlow) {
V2836 := __e.Get(1)
_ = V2836
tmp4820 := MakeNative(func(__e *ControlFlow) {
W2837 := __e.Get(1)
_ = W2837
tmp4822 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2837)


if True == tmp4822 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2837)
return
}


}, 1)

tmp4823 := MakeNative(func(__e *ControlFlow) {
W2838 := __e.Get(1)
_ = W2838
tmp4839 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2838)


if True == tmp4839 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4824 := MakeNative(func(__e *ControlFlow) {
W2839 := __e.Get(1)
_ = W2839
tmp4825 := MakeNative(func(__e *ControlFlow) {
W2840 := __e.Get(1)
_ = W2840
tmp4835 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2840)


if True == tmp4835 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4826 := MakeNative(func(__e *ControlFlow) {
W2841 := __e.Get(1)
_ = W2841
tmp4827 := MakeNative(func(__e *ControlFlow) {
W2842 := __e.Get(1)
_ = W2842
tmp4831 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2842)


if True == tmp4831 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4828 := MakeNative(func(__e *ControlFlow) {
W2843 := __e.Get(1)
_ = W2843
__e.TailApply(PrimFunc(symshen_4comb), W2843, symshen_4skip)
return
}, 1)

tmp4829 := Call(__e, PrimFunc(symshen_4in_1_6), W2842)


__e.TailApply(tmp4828, tmp4829)
return


}


}, 1)

tmp4832 := Call(__e, PrimFunc(symshen_4_5longnatter_6), W2841)


__e.TailApply(tmp4827, tmp4832)
return


}, 1)

tmp4833 := Call(__e, PrimFunc(symshen_4in_1_6), W2840)


__e.TailApply(tmp4826, tmp4833)
return


}


}, 1)

tmp4836 := Call(__e, PrimFunc(symshen_4_5times_6), W2839)


__e.TailApply(tmp4825, tmp4836)
return


}, 1)

tmp4837 := Call(__e, PrimFunc(symshen_4in_1_6), W2838)


__e.TailApply(tmp4824, tmp4837)
return


}


}, 1)

tmp4840 := Call(__e, PrimFunc(symshen_4_5backslash_6), V2836)


tmp4841 := Call(__e, tmp4823, tmp4840)


__e.TailApply(tmp4820, tmp4841)
return


}, 1)

tmp4842 := Call(__e, ns2_1set, symshen_4_5multiline_6, tmp4819)


_ = tmp4842

tmp4843 := MakeNative(func(__e *ControlFlow) {
V2844 := __e.Get(1)
_ = V2844
tmp4844 := MakeNative(func(__e *ControlFlow) {
W2845 := __e.Get(1)
_ = W2845
tmp4846 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2845)


if True == tmp4846 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2845)
return
}


}, 1)

tmp4852 := Call(__e, PrimFunc(symshen_4hds_a_2), V2844, MakeNumber(42))


var ifres4847 Obj

if True == tmp4852 {
tmp4848 := MakeNative(func(__e *ControlFlow) {
W2846 := __e.Get(1)
_ = W2846
__e.TailApply(PrimFunc(symshen_4comb), W2846, symshen_4skip)
return
}, 1)

tmp4849 := Call(__e, PrimFunc(symtail), V2844)


tmp4850 := Call(__e, tmp4848, tmp4849)


ifres4847 = tmp4850


} else {
tmp4851 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres4847 = tmp4851


}

__e.TailApply(tmp4844, ifres4847)
return


}, 1)

tmp4853 := Call(__e, ns2_1set, symshen_4_5times_6, tmp4843)


_ = tmp4853

tmp4854 := MakeNative(func(__e *ControlFlow) {
V2847 := __e.Get(1)
_ = V2847
tmp4855 := MakeNative(func(__e *ControlFlow) {
W2848 := __e.Get(1)
_ = W2848
tmp4888 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2848)


if True == tmp4888 {
tmp4856 := MakeNative(func(__e *ControlFlow) {
W2853 := __e.Get(1)
_ = W2853
tmp4873 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2853)


if True == tmp4873 {
tmp4857 := MakeNative(func(__e *ControlFlow) {
W2858 := __e.Get(1)
_ = W2858
tmp4859 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2858)


if True == tmp4859 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2858)
return
}


}, 1)

tmp4871 := PrimIsPair(V2847)

var ifres4860 Obj

if True == tmp4871 {
tmp4861 := MakeNative(func(__e *ControlFlow) {
W2859 := __e.Get(1)
_ = W2859
tmp4862 := MakeNative(func(__e *ControlFlow) {
W2860 := __e.Get(1)
_ = W2860
tmp4866 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2860)


if True == tmp4866 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4863 := MakeNative(func(__e *ControlFlow) {
W2861 := __e.Get(1)
_ = W2861
__e.TailApply(PrimFunc(symshen_4comb), W2861, symshen_4skip)
return
}, 1)

tmp4864 := Call(__e, PrimFunc(symshen_4in_1_6), W2860)


__e.TailApply(tmp4863, tmp4864)
return


}


}, 1)

tmp4867 := Call(__e, PrimFunc(symshen_4_5longnatter_6), W2859)


__e.TailApply(tmp4862, tmp4867)
return


}, 1)

tmp4868 := Call(__e, PrimFunc(symtail), V2847)


tmp4869 := Call(__e, tmp4861, tmp4868)


ifres4860 = tmp4869


} else {
tmp4870 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres4860 = tmp4870


}

__e.TailApply(tmp4857, ifres4860)
return


} else {
__e.Return(W2853)
return
}


}, 1)

tmp4874 := MakeNative(func(__e *ControlFlow) {
W2854 := __e.Get(1)
_ = W2854
tmp4884 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2854)


if True == tmp4884 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4875 := MakeNative(func(__e *ControlFlow) {
W2855 := __e.Get(1)
_ = W2855
tmp4876 := MakeNative(func(__e *ControlFlow) {
W2856 := __e.Get(1)
_ = W2856
tmp4880 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2856)


if True == tmp4880 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4877 := MakeNative(func(__e *ControlFlow) {
W2857 := __e.Get(1)
_ = W2857
__e.TailApply(PrimFunc(symshen_4comb), W2857, symshen_4skip)
return
}, 1)

tmp4878 := Call(__e, PrimFunc(symshen_4in_1_6), W2856)


__e.TailApply(tmp4877, tmp4878)
return


}


}, 1)

tmp4881 := Call(__e, PrimFunc(symshen_4_5backslash_6), W2855)


__e.TailApply(tmp4876, tmp4881)
return


}, 1)

tmp4882 := Call(__e, PrimFunc(symshen_4in_1_6), W2854)


__e.TailApply(tmp4875, tmp4882)
return


}


}, 1)

tmp4885 := Call(__e, PrimFunc(symshen_4_5times_6), V2847)


tmp4886 := Call(__e, tmp4874, tmp4885)


__e.TailApply(tmp4856, tmp4886)
return


} else {
__e.Return(W2848)
return
}


}, 1)

tmp4889 := MakeNative(func(__e *ControlFlow) {
W2849 := __e.Get(1)
_ = W2849
tmp4899 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2849)


if True == tmp4899 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4890 := MakeNative(func(__e *ControlFlow) {
W2850 := __e.Get(1)
_ = W2850
tmp4891 := MakeNative(func(__e *ControlFlow) {
W2851 := __e.Get(1)
_ = W2851
tmp4895 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2851)


if True == tmp4895 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4892 := MakeNative(func(__e *ControlFlow) {
W2852 := __e.Get(1)
_ = W2852
__e.TailApply(PrimFunc(symshen_4comb), W2852, symshen_4skip)
return
}, 1)

tmp4893 := Call(__e, PrimFunc(symshen_4in_1_6), W2851)


__e.TailApply(tmp4892, tmp4893)
return


}


}, 1)

tmp4896 := Call(__e, PrimFunc(symshen_4_5longnatter_6), W2850)


__e.TailApply(tmp4891, tmp4896)
return


}, 1)

tmp4897 := Call(__e, PrimFunc(symshen_4in_1_6), W2849)


__e.TailApply(tmp4890, tmp4897)
return


}


}, 1)

tmp4900 := Call(__e, PrimFunc(symshen_4_5comment_6), V2847)


tmp4901 := Call(__e, tmp4889, tmp4900)


__e.TailApply(tmp4855, tmp4901)
return


}, 1)

tmp4902 := Call(__e, ns2_1set, symshen_4_5longnatter_6, tmp4854)


_ = tmp4902

tmp4903 := MakeNative(func(__e *ControlFlow) {
V2862 := __e.Get(1)
_ = V2862
tmp4904 := MakeNative(func(__e *ControlFlow) {
W2863 := __e.Get(1)
_ = W2863
tmp4935 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2863)


if True == tmp4935 {
tmp4905 := MakeNative(func(__e *ControlFlow) {
W2867 := __e.Get(1)
_ = W2867
tmp4924 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2867)


if True == tmp4924 {
tmp4906 := MakeNative(func(__e *ControlFlow) {
W2871 := __e.Get(1)
_ = W2871
tmp4908 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2871)


if True == tmp4908 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2871)
return
}


}, 1)

tmp4909 := MakeNative(func(__e *ControlFlow) {
W2872 := __e.Get(1)
_ = W2872
tmp4920 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2872)


if True == tmp4920 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4910 := MakeNative(func(__e *ControlFlow) {
W2873 := __e.Get(1)
_ = W2873
tmp4911 := MakeNative(func(__e *ControlFlow) {
W2874 := __e.Get(1)
_ = W2874
tmp4916 := PrimEqual(W2873, MakeString("<>"))

var ifres4912 Obj

if True == tmp4916 {
tmp4913 := PrimCons(MakeNumber(0), Nil)

tmp4914 := PrimCons(symvector, tmp4913)

ifres4912 = tmp4914


} else {
tmp4915 := PrimIntern(W2873)

ifres4912 = tmp4915


}

__e.TailApply(PrimFunc(symshen_4comb), W2874, ifres4912)
return


}, 1)

tmp4917 := Call(__e, PrimFunc(symshen_4in_1_6), W2872)


__e.TailApply(tmp4911, tmp4917)
return


}, 1)

tmp4918 := Call(__e, PrimFunc(symshen_4_5_1out), W2872)


__e.TailApply(tmp4910, tmp4918)
return


}


}, 1)

tmp4921 := Call(__e, PrimFunc(symshen_4_5sym_6), V2862)


tmp4922 := Call(__e, tmp4909, tmp4921)


__e.TailApply(tmp4906, tmp4922)
return


} else {
__e.Return(W2867)
return
}


}, 1)

tmp4925 := MakeNative(func(__e *ControlFlow) {
W2868 := __e.Get(1)
_ = W2868
tmp4931 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2868)


if True == tmp4931 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4926 := MakeNative(func(__e *ControlFlow) {
W2869 := __e.Get(1)
_ = W2869
tmp4927 := MakeNative(func(__e *ControlFlow) {
W2870 := __e.Get(1)
_ = W2870
__e.TailApply(PrimFunc(symshen_4comb), W2870, W2869)
return
}, 1)

tmp4928 := Call(__e, PrimFunc(symshen_4in_1_6), W2868)


__e.TailApply(tmp4927, tmp4928)
return


}, 1)

tmp4929 := Call(__e, PrimFunc(symshen_4_5_1out), W2868)


__e.TailApply(tmp4926, tmp4929)
return


}


}, 1)

tmp4932 := Call(__e, PrimFunc(symshen_4_5number_6), V2862)


tmp4933 := Call(__e, tmp4925, tmp4932)


__e.TailApply(tmp4905, tmp4933)
return


} else {
__e.Return(W2863)
return
}


}, 1)

tmp4936 := MakeNative(func(__e *ControlFlow) {
W2864 := __e.Get(1)
_ = W2864
tmp4942 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2864)


if True == tmp4942 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4937 := MakeNative(func(__e *ControlFlow) {
W2865 := __e.Get(1)
_ = W2865
tmp4938 := MakeNative(func(__e *ControlFlow) {
W2866 := __e.Get(1)
_ = W2866
__e.TailApply(PrimFunc(symshen_4comb), W2866, W2865)
return
}, 1)

tmp4939 := Call(__e, PrimFunc(symshen_4in_1_6), W2864)


__e.TailApply(tmp4938, tmp4939)
return


}, 1)

tmp4940 := Call(__e, PrimFunc(symshen_4_5_1out), W2864)


__e.TailApply(tmp4937, tmp4940)
return


}


}, 1)

tmp4943 := Call(__e, PrimFunc(symshen_4_5str_6), V2862)


tmp4944 := Call(__e, tmp4936, tmp4943)


__e.TailApply(tmp4904, tmp4944)
return


}, 1)

tmp4945 := Call(__e, ns2_1set, symshen_4_5atom_6, tmp4903)


_ = tmp4945

tmp4946 := MakeNative(func(__e *ControlFlow) {
V2875 := __e.Get(1)
_ = V2875
tmp4947 := MakeNative(func(__e *ControlFlow) {
W2876 := __e.Get(1)
_ = W2876
tmp4949 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2876)


if True == tmp4949 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2876)
return
}


}, 1)

tmp4950 := MakeNative(func(__e *ControlFlow) {
W2877 := __e.Get(1)
_ = W2877
tmp4965 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2877)


if True == tmp4965 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4951 := MakeNative(func(__e *ControlFlow) {
W2878 := __e.Get(1)
_ = W2878
tmp4952 := MakeNative(func(__e *ControlFlow) {
W2879 := __e.Get(1)
_ = W2879
tmp4953 := MakeNative(func(__e *ControlFlow) {
W2880 := __e.Get(1)
_ = W2880
tmp4960 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2880)


if True == tmp4960 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp4954 := MakeNative(func(__e *ControlFlow) {
W2881 := __e.Get(1)
_ = W2881
tmp4955 := MakeNative(func(__e *ControlFlow) {
W2882 := __e.Get(1)
_ = W2882
tmp4956 := PrimStringConcat(W2878, W2881)

__e.TailApply(PrimFunc(symshen_4comb), W2882, tmp4956)
return


}, 1)

tmp4957 := Call(__e, PrimFunc(symshen_4in_1_6), W2880)


__e.TailApply(tmp4955, tmp4957)
return


}, 1)

tmp4958 := Call(__e, PrimFunc(symshen_4_5_1out), W2880)


__e.TailApply(tmp4954, tmp4958)
return


}


}, 1)

tmp4961 := Call(__e, PrimFunc(symshen_4_5alphanums_6), W2879)


__e.TailApply(tmp4953, tmp4961)
return


}, 1)

tmp4962 := Call(__e, PrimFunc(symshen_4in_1_6), W2877)


__e.TailApply(tmp4952, tmp4962)
return


}, 1)

tmp4963 := Call(__e, PrimFunc(symshen_4_5_1out), W2877)


__e.TailApply(tmp4951, tmp4963)
return


}


}, 1)

tmp4966 := Call(__e, PrimFunc(symshen_4_5alpha_6), V2875)


tmp4967 := Call(__e, tmp4950, tmp4966)


__e.TailApply(tmp4947, tmp4967)
return


}, 1)

tmp4968 := Call(__e, ns2_1set, symshen_4_5sym_6, tmp4946)


_ = tmp4968

tmp4969 := MakeNative(func(__e *ControlFlow) {
V2883 := __e.Get(1)
_ = V2883
tmp4970 := MakeNative(func(__e *ControlFlow) {
W2884 := __e.Get(1)
_ = W2884
tmp4972 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2884)


if True == tmp4972 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2884)
return
}


}, 1)

tmp4983 := PrimIsPair(V2883)

var ifres4973 Obj

if True == tmp4983 {
tmp4974 := MakeNative(func(__e *ControlFlow) {
W2885 := __e.Get(1)
_ = W2885
tmp4975 := MakeNative(func(__e *ControlFlow) {
W2886 := __e.Get(1)
_ = W2886
tmp4978 := Call(__e, PrimFunc(symshen_4alpha_2), W2885)


if True == tmp4978 {
tmp4976 := PrimNumberToString(W2885)

__e.TailApply(PrimFunc(symshen_4comb), W2886, tmp4976)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp4979 := Call(__e, PrimFunc(symtail), V2883)


__e.TailApply(tmp4975, tmp4979)
return


}, 1)

tmp4980 := Call(__e, PrimFunc(symhead), V2883)


tmp4981 := Call(__e, tmp4974, tmp4980)


ifres4973 = tmp4981


} else {
tmp4982 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres4973 = tmp4982


}

__e.TailApply(tmp4970, ifres4973)
return


}, 1)

tmp4984 := Call(__e, ns2_1set, symshen_4_5alpha_6, tmp4969)


_ = tmp4984

tmp4985 := MakeNative(func(__e *ControlFlow) {
V2887 := __e.Get(1)
_ = V2887
tmp4992 := Call(__e, PrimFunc(symshen_4lowercase_2), V2887)


if True == tmp4992 {
__e.Return(True)
return
} else {
tmp4990 := Call(__e, PrimFunc(symshen_4uppercase_2), V2887)


var ifres4987 Obj

if True == tmp4990 {
ifres4987 = True


} else {
tmp4989 := Call(__e, PrimFunc(symshen_4misc_2), V2887)


var ifres4988 Obj

if True == tmp4989 {
ifres4988 = True


} else {
ifres4988 = False


}

ifres4987 = ifres4988


}

if True == ifres4987 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp4993 := Call(__e, ns2_1set, symshen_4alpha_2, tmp4985)


_ = tmp4993

tmp4994 := MakeNative(func(__e *ControlFlow) {
V2888 := __e.Get(1)
_ = V2888
tmp4998 := PrimGreatEqual(V2888, MakeNumber(97))

if True == tmp4998 {
tmp4996 := PrimLessEqual(V2888, MakeNumber(122))

if True == tmp4996 {
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

tmp4999 := Call(__e, ns2_1set, symshen_4lowercase_2, tmp4994)


_ = tmp4999

tmp5000 := MakeNative(func(__e *ControlFlow) {
V2889 := __e.Get(1)
_ = V2889
tmp5004 := PrimGreatEqual(V2889, MakeNumber(65))

if True == tmp5004 {
tmp5002 := PrimLessEqual(V2889, MakeNumber(90))

if True == tmp5002 {
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

tmp5005 := Call(__e, ns2_1set, symshen_4uppercase_2, tmp5000)


_ = tmp5005

tmp5006 := MakeNative(func(__e *ControlFlow) {
V2890 := __e.Get(1)
_ = V2890
tmp5007 := PrimCons(MakeNumber(96), Nil)

tmp5008 := PrimCons(MakeNumber(35), tmp5007)

tmp5009 := PrimCons(MakeNumber(39), tmp5008)

tmp5010 := PrimCons(MakeNumber(37), tmp5009)

tmp5011 := PrimCons(MakeNumber(38), tmp5010)

tmp5012 := PrimCons(MakeNumber(60), tmp5011)

tmp5013 := PrimCons(MakeNumber(62), tmp5012)

tmp5014 := PrimCons(MakeNumber(46), tmp5013)

tmp5015 := PrimCons(MakeNumber(126), tmp5014)

tmp5016 := PrimCons(MakeNumber(64), tmp5015)

tmp5017 := PrimCons(MakeNumber(33), tmp5016)

tmp5018 := PrimCons(MakeNumber(36), tmp5017)

tmp5019 := PrimCons(MakeNumber(63), tmp5018)

tmp5020 := PrimCons(MakeNumber(95), tmp5019)

tmp5021 := PrimCons(MakeNumber(43), tmp5020)

tmp5022 := PrimCons(MakeNumber(47), tmp5021)

tmp5023 := PrimCons(MakeNumber(42), tmp5022)

tmp5024 := PrimCons(MakeNumber(45), tmp5023)

tmp5025 := PrimCons(MakeNumber(61), tmp5024)

__e.TailApply(PrimFunc(symelement_2), V2890, tmp5025)
return


}, 1)

tmp5026 := Call(__e, ns2_1set, symshen_4misc_2, tmp5006)


_ = tmp5026

tmp5027 := MakeNative(func(__e *ControlFlow) {
V2891 := __e.Get(1)
_ = V2891
tmp5028 := MakeNative(func(__e *ControlFlow) {
W2892 := __e.Get(1)
_ = W2892
tmp5040 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2892)


if True == tmp5040 {
tmp5029 := MakeNative(func(__e *ControlFlow) {
W2899 := __e.Get(1)
_ = W2899
tmp5031 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2899)


if True == tmp5031 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2899)
return
}


}, 1)

tmp5032 := MakeNative(func(__e *ControlFlow) {
W2900 := __e.Get(1)
_ = W2900
tmp5036 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2900)


if True == tmp5036 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5033 := MakeNative(func(__e *ControlFlow) {
W2901 := __e.Get(1)
_ = W2901
__e.TailApply(PrimFunc(symshen_4comb), W2901, MakeString(""))
return
}, 1)

tmp5034 := Call(__e, PrimFunc(symshen_4in_1_6), W2900)


__e.TailApply(tmp5033, tmp5034)
return


}


}, 1)

tmp5037 := Call(__e, PrimFunc(sym_5e_6), V2891)


tmp5038 := Call(__e, tmp5032, tmp5037)


__e.TailApply(tmp5029, tmp5038)
return


} else {
__e.Return(W2892)
return
}


}, 1)

tmp5041 := MakeNative(func(__e *ControlFlow) {
W2893 := __e.Get(1)
_ = W2893
tmp5056 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2893)


if True == tmp5056 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5042 := MakeNative(func(__e *ControlFlow) {
W2894 := __e.Get(1)
_ = W2894
tmp5043 := MakeNative(func(__e *ControlFlow) {
W2895 := __e.Get(1)
_ = W2895
tmp5044 := MakeNative(func(__e *ControlFlow) {
W2896 := __e.Get(1)
_ = W2896
tmp5051 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2896)


if True == tmp5051 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5045 := MakeNative(func(__e *ControlFlow) {
W2897 := __e.Get(1)
_ = W2897
tmp5046 := MakeNative(func(__e *ControlFlow) {
W2898 := __e.Get(1)
_ = W2898
tmp5047 := PrimStringConcat(W2894, W2897)

__e.TailApply(PrimFunc(symshen_4comb), W2898, tmp5047)
return


}, 1)

tmp5048 := Call(__e, PrimFunc(symshen_4in_1_6), W2896)


__e.TailApply(tmp5046, tmp5048)
return


}, 1)

tmp5049 := Call(__e, PrimFunc(symshen_4_5_1out), W2896)


__e.TailApply(tmp5045, tmp5049)
return


}


}, 1)

tmp5052 := Call(__e, PrimFunc(symshen_4_5alphanums_6), W2895)


__e.TailApply(tmp5044, tmp5052)
return


}, 1)

tmp5053 := Call(__e, PrimFunc(symshen_4in_1_6), W2893)


__e.TailApply(tmp5043, tmp5053)
return


}, 1)

tmp5054 := Call(__e, PrimFunc(symshen_4_5_1out), W2893)


__e.TailApply(tmp5042, tmp5054)
return


}


}, 1)

tmp5057 := Call(__e, PrimFunc(symshen_4_5alphanum_6), V2891)


tmp5058 := Call(__e, tmp5041, tmp5057)


__e.TailApply(tmp5028, tmp5058)
return


}, 1)

tmp5059 := Call(__e, ns2_1set, symshen_4_5alphanums_6, tmp5027)


_ = tmp5059

tmp5060 := MakeNative(func(__e *ControlFlow) {
V2902 := __e.Get(1)
_ = V2902
tmp5061 := MakeNative(func(__e *ControlFlow) {
W2903 := __e.Get(1)
_ = W2903
tmp5075 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2903)


if True == tmp5075 {
tmp5062 := MakeNative(func(__e *ControlFlow) {
W2907 := __e.Get(1)
_ = W2907
tmp5064 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2907)


if True == tmp5064 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2907)
return
}


}, 1)

tmp5065 := MakeNative(func(__e *ControlFlow) {
W2908 := __e.Get(1)
_ = W2908
tmp5071 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2908)


if True == tmp5071 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5066 := MakeNative(func(__e *ControlFlow) {
W2909 := __e.Get(1)
_ = W2909
tmp5067 := MakeNative(func(__e *ControlFlow) {
W2910 := __e.Get(1)
_ = W2910
__e.TailApply(PrimFunc(symshen_4comb), W2910, W2909)
return
}, 1)

tmp5068 := Call(__e, PrimFunc(symshen_4in_1_6), W2908)


__e.TailApply(tmp5067, tmp5068)
return


}, 1)

tmp5069 := Call(__e, PrimFunc(symshen_4_5_1out), W2908)


__e.TailApply(tmp5066, tmp5069)
return


}


}, 1)

tmp5072 := Call(__e, PrimFunc(symshen_4_5numeral_6), V2902)


tmp5073 := Call(__e, tmp5065, tmp5072)


__e.TailApply(tmp5062, tmp5073)
return


} else {
__e.Return(W2903)
return
}


}, 1)

tmp5076 := MakeNative(func(__e *ControlFlow) {
W2904 := __e.Get(1)
_ = W2904
tmp5082 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2904)


if True == tmp5082 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5077 := MakeNative(func(__e *ControlFlow) {
W2905 := __e.Get(1)
_ = W2905
tmp5078 := MakeNative(func(__e *ControlFlow) {
W2906 := __e.Get(1)
_ = W2906
__e.TailApply(PrimFunc(symshen_4comb), W2906, W2905)
return
}, 1)

tmp5079 := Call(__e, PrimFunc(symshen_4in_1_6), W2904)


__e.TailApply(tmp5078, tmp5079)
return


}, 1)

tmp5080 := Call(__e, PrimFunc(symshen_4_5_1out), W2904)


__e.TailApply(tmp5077, tmp5080)
return


}


}, 1)

tmp5083 := Call(__e, PrimFunc(symshen_4_5alpha_6), V2902)


tmp5084 := Call(__e, tmp5076, tmp5083)


__e.TailApply(tmp5061, tmp5084)
return


}, 1)

tmp5085 := Call(__e, ns2_1set, symshen_4_5alphanum_6, tmp5060)


_ = tmp5085

tmp5086 := MakeNative(func(__e *ControlFlow) {
V2911 := __e.Get(1)
_ = V2911
tmp5087 := MakeNative(func(__e *ControlFlow) {
W2912 := __e.Get(1)
_ = W2912
tmp5089 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2912)


if True == tmp5089 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2912)
return
}


}, 1)

tmp5100 := PrimIsPair(V2911)

var ifres5090 Obj

if True == tmp5100 {
tmp5091 := MakeNative(func(__e *ControlFlow) {
W2913 := __e.Get(1)
_ = W2913
tmp5092 := MakeNative(func(__e *ControlFlow) {
W2914 := __e.Get(1)
_ = W2914
tmp5095 := Call(__e, PrimFunc(symshen_4digit_2), W2913)


if True == tmp5095 {
tmp5093 := PrimNumberToString(W2913)

__e.TailApply(PrimFunc(symshen_4comb), W2914, tmp5093)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp5096 := Call(__e, PrimFunc(symtail), V2911)


__e.TailApply(tmp5092, tmp5096)
return


}, 1)

tmp5097 := Call(__e, PrimFunc(symhead), V2911)


tmp5098 := Call(__e, tmp5091, tmp5097)


ifres5090 = tmp5098


} else {
tmp5099 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5090 = tmp5099


}

__e.TailApply(tmp5087, ifres5090)
return


}, 1)

tmp5101 := Call(__e, ns2_1set, symshen_4_5numeral_6, tmp5086)


_ = tmp5101

tmp5102 := MakeNative(func(__e *ControlFlow) {
V2915 := __e.Get(1)
_ = V2915
tmp5106 := PrimGreatEqual(V2915, MakeNumber(48))

if True == tmp5106 {
tmp5104 := PrimLessEqual(V2915, MakeNumber(57))

if True == tmp5104 {
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

tmp5107 := Call(__e, ns2_1set, symshen_4digit_2, tmp5102)


_ = tmp5107

tmp5108 := MakeNative(func(__e *ControlFlow) {
V2916 := __e.Get(1)
_ = V2916
tmp5109 := MakeNative(func(__e *ControlFlow) {
W2917 := __e.Get(1)
_ = W2917
tmp5111 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2917)


if True == tmp5111 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2917)
return
}


}, 1)

tmp5112 := MakeNative(func(__e *ControlFlow) {
W2918 := __e.Get(1)
_ = W2918
tmp5130 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2918)


if True == tmp5130 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5113 := MakeNative(func(__e *ControlFlow) {
W2919 := __e.Get(1)
_ = W2919
tmp5114 := MakeNative(func(__e *ControlFlow) {
W2920 := __e.Get(1)
_ = W2920
tmp5126 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2920)


if True == tmp5126 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5115 := MakeNative(func(__e *ControlFlow) {
W2921 := __e.Get(1)
_ = W2921
tmp5116 := MakeNative(func(__e *ControlFlow) {
W2922 := __e.Get(1)
_ = W2922
tmp5117 := MakeNative(func(__e *ControlFlow) {
W2923 := __e.Get(1)
_ = W2923
tmp5121 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2923)


if True == tmp5121 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5118 := MakeNative(func(__e *ControlFlow) {
W2924 := __e.Get(1)
_ = W2924
__e.TailApply(PrimFunc(symshen_4comb), W2924, W2921)
return
}, 1)

tmp5119 := Call(__e, PrimFunc(symshen_4in_1_6), W2923)


__e.TailApply(tmp5118, tmp5119)
return


}


}, 1)

tmp5122 := Call(__e, PrimFunc(symshen_4_5dbq_6), W2922)


__e.TailApply(tmp5117, tmp5122)
return


}, 1)

tmp5123 := Call(__e, PrimFunc(symshen_4in_1_6), W2920)


__e.TailApply(tmp5116, tmp5123)
return


}, 1)

tmp5124 := Call(__e, PrimFunc(symshen_4_5_1out), W2920)


__e.TailApply(tmp5115, tmp5124)
return


}


}, 1)

tmp5127 := Call(__e, PrimFunc(symshen_4_5strcontents_6), W2919)


__e.TailApply(tmp5114, tmp5127)
return


}, 1)

tmp5128 := Call(__e, PrimFunc(symshen_4in_1_6), W2918)


__e.TailApply(tmp5113, tmp5128)
return


}


}, 1)

tmp5131 := Call(__e, PrimFunc(symshen_4_5dbq_6), V2916)


tmp5132 := Call(__e, tmp5112, tmp5131)


__e.TailApply(tmp5109, tmp5132)
return


}, 1)

tmp5133 := Call(__e, ns2_1set, symshen_4_5str_6, tmp5108)


_ = tmp5133

tmp5134 := MakeNative(func(__e *ControlFlow) {
V2925 := __e.Get(1)
_ = V2925
tmp5135 := MakeNative(func(__e *ControlFlow) {
W2926 := __e.Get(1)
_ = W2926
tmp5137 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2926)


if True == tmp5137 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2926)
return
}


}, 1)

tmp5143 := Call(__e, PrimFunc(symshen_4hds_a_2), V2925, MakeNumber(34))


var ifres5138 Obj

if True == tmp5143 {
tmp5139 := MakeNative(func(__e *ControlFlow) {
W2927 := __e.Get(1)
_ = W2927
__e.TailApply(PrimFunc(symshen_4comb), W2927, symshen_4skip)
return
}, 1)

tmp5140 := Call(__e, PrimFunc(symtail), V2925)


tmp5141 := Call(__e, tmp5139, tmp5140)


ifres5138 = tmp5141


} else {
tmp5142 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5138 = tmp5142


}

__e.TailApply(tmp5135, ifres5138)
return


}, 1)

tmp5144 := Call(__e, ns2_1set, symshen_4_5dbq_6, tmp5134)


_ = tmp5144

tmp5145 := MakeNative(func(__e *ControlFlow) {
V2928 := __e.Get(1)
_ = V2928
tmp5146 := MakeNative(func(__e *ControlFlow) {
W2929 := __e.Get(1)
_ = W2929
tmp5158 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2929)


if True == tmp5158 {
tmp5147 := MakeNative(func(__e *ControlFlow) {
W2936 := __e.Get(1)
_ = W2936
tmp5149 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2936)


if True == tmp5149 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2936)
return
}


}, 1)

tmp5150 := MakeNative(func(__e *ControlFlow) {
W2937 := __e.Get(1)
_ = W2937
tmp5154 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2937)


if True == tmp5154 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5151 := MakeNative(func(__e *ControlFlow) {
W2938 := __e.Get(1)
_ = W2938
__e.TailApply(PrimFunc(symshen_4comb), W2938, MakeString(""))
return
}, 1)

tmp5152 := Call(__e, PrimFunc(symshen_4in_1_6), W2937)


__e.TailApply(tmp5151, tmp5152)
return


}


}, 1)

tmp5155 := Call(__e, PrimFunc(sym_5e_6), V2928)


tmp5156 := Call(__e, tmp5150, tmp5155)


__e.TailApply(tmp5147, tmp5156)
return


} else {
__e.Return(W2929)
return
}


}, 1)

tmp5159 := MakeNative(func(__e *ControlFlow) {
W2930 := __e.Get(1)
_ = W2930
tmp5174 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2930)


if True == tmp5174 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5160 := MakeNative(func(__e *ControlFlow) {
W2931 := __e.Get(1)
_ = W2931
tmp5161 := MakeNative(func(__e *ControlFlow) {
W2932 := __e.Get(1)
_ = W2932
tmp5162 := MakeNative(func(__e *ControlFlow) {
W2933 := __e.Get(1)
_ = W2933
tmp5169 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2933)


if True == tmp5169 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5163 := MakeNative(func(__e *ControlFlow) {
W2934 := __e.Get(1)
_ = W2934
tmp5164 := MakeNative(func(__e *ControlFlow) {
W2935 := __e.Get(1)
_ = W2935
tmp5165 := PrimStringConcat(W2931, W2934)

__e.TailApply(PrimFunc(symshen_4comb), W2935, tmp5165)
return


}, 1)

tmp5166 := Call(__e, PrimFunc(symshen_4in_1_6), W2933)


__e.TailApply(tmp5164, tmp5166)
return


}, 1)

tmp5167 := Call(__e, PrimFunc(symshen_4_5_1out), W2933)


__e.TailApply(tmp5163, tmp5167)
return


}


}, 1)

tmp5170 := Call(__e, PrimFunc(symshen_4_5strcontents_6), W2932)


__e.TailApply(tmp5162, tmp5170)
return


}, 1)

tmp5171 := Call(__e, PrimFunc(symshen_4in_1_6), W2930)


__e.TailApply(tmp5161, tmp5171)
return


}, 1)

tmp5172 := Call(__e, PrimFunc(symshen_4_5_1out), W2930)


__e.TailApply(tmp5160, tmp5172)
return


}


}, 1)

tmp5175 := Call(__e, PrimFunc(symshen_4_5strc_6), V2928)


tmp5176 := Call(__e, tmp5159, tmp5175)


__e.TailApply(tmp5146, tmp5176)
return


}, 1)

tmp5177 := Call(__e, ns2_1set, symshen_4_5strcontents_6, tmp5145)


_ = tmp5177

tmp5178 := MakeNative(func(__e *ControlFlow) {
V2939 := __e.Get(1)
_ = V2939
tmp5179 := MakeNative(func(__e *ControlFlow) {
W2940 := __e.Get(1)
_ = W2940
tmp5193 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2940)


if True == tmp5193 {
tmp5180 := MakeNative(func(__e *ControlFlow) {
W2944 := __e.Get(1)
_ = W2944
tmp5182 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2944)


if True == tmp5182 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2944)
return
}


}, 1)

tmp5183 := MakeNative(func(__e *ControlFlow) {
W2945 := __e.Get(1)
_ = W2945
tmp5189 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2945)


if True == tmp5189 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5184 := MakeNative(func(__e *ControlFlow) {
W2946 := __e.Get(1)
_ = W2946
tmp5185 := MakeNative(func(__e *ControlFlow) {
W2947 := __e.Get(1)
_ = W2947
__e.TailApply(PrimFunc(symshen_4comb), W2947, W2946)
return
}, 1)

tmp5186 := Call(__e, PrimFunc(symshen_4in_1_6), W2945)


__e.TailApply(tmp5185, tmp5186)
return


}, 1)

tmp5187 := Call(__e, PrimFunc(symshen_4_5_1out), W2945)


__e.TailApply(tmp5184, tmp5187)
return


}


}, 1)

tmp5190 := Call(__e, PrimFunc(symshen_4_5notdbq_6), V2939)


tmp5191 := Call(__e, tmp5183, tmp5190)


__e.TailApply(tmp5180, tmp5191)
return


} else {
__e.Return(W2940)
return
}


}, 1)

tmp5194 := MakeNative(func(__e *ControlFlow) {
W2941 := __e.Get(1)
_ = W2941
tmp5200 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2941)


if True == tmp5200 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5195 := MakeNative(func(__e *ControlFlow) {
W2942 := __e.Get(1)
_ = W2942
tmp5196 := MakeNative(func(__e *ControlFlow) {
W2943 := __e.Get(1)
_ = W2943
__e.TailApply(PrimFunc(symshen_4comb), W2943, W2942)
return
}, 1)

tmp5197 := Call(__e, PrimFunc(symshen_4in_1_6), W2941)


__e.TailApply(tmp5196, tmp5197)
return


}, 1)

tmp5198 := Call(__e, PrimFunc(symshen_4_5_1out), W2941)


__e.TailApply(tmp5195, tmp5198)
return


}


}, 1)

tmp5201 := Call(__e, PrimFunc(symshen_4_5control_6), V2939)


tmp5202 := Call(__e, tmp5194, tmp5201)


__e.TailApply(tmp5179, tmp5202)
return


}, 1)

tmp5203 := Call(__e, ns2_1set, symshen_4_5strc_6, tmp5178)


_ = tmp5203

tmp5204 := MakeNative(func(__e *ControlFlow) {
V2948 := __e.Get(1)
_ = V2948
tmp5205 := MakeNative(func(__e *ControlFlow) {
W2949 := __e.Get(1)
_ = W2949
tmp5207 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2949)


if True == tmp5207 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2949)
return
}


}, 1)

tmp5208 := MakeNative(func(__e *ControlFlow) {
W2950 := __e.Get(1)
_ = W2950
tmp5233 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2950)


if True == tmp5233 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5209 := MakeNative(func(__e *ControlFlow) {
W2951 := __e.Get(1)
_ = W2951
tmp5210 := MakeNative(func(__e *ControlFlow) {
W2952 := __e.Get(1)
_ = W2952
tmp5229 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2952)


if True == tmp5229 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5211 := MakeNative(func(__e *ControlFlow) {
W2953 := __e.Get(1)
_ = W2953
tmp5212 := MakeNative(func(__e *ControlFlow) {
W2954 := __e.Get(1)
_ = W2954
tmp5225 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2954)


if True == tmp5225 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5213 := MakeNative(func(__e *ControlFlow) {
W2955 := __e.Get(1)
_ = W2955
tmp5214 := MakeNative(func(__e *ControlFlow) {
W2956 := __e.Get(1)
_ = W2956
tmp5215 := MakeNative(func(__e *ControlFlow) {
W2957 := __e.Get(1)
_ = W2957
tmp5220 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2957)


if True == tmp5220 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5216 := MakeNative(func(__e *ControlFlow) {
W2958 := __e.Get(1)
_ = W2958
tmp5217 := PrimNumberToString(W2955)

__e.TailApply(PrimFunc(symshen_4comb), W2958, tmp5217)
return


}, 1)

tmp5218 := Call(__e, PrimFunc(symshen_4in_1_6), W2957)


__e.TailApply(tmp5216, tmp5218)
return


}


}, 1)

tmp5221 := Call(__e, PrimFunc(symshen_4_5semicolon_6), W2956)


__e.TailApply(tmp5215, tmp5221)
return


}, 1)

tmp5222 := Call(__e, PrimFunc(symshen_4in_1_6), W2954)


__e.TailApply(tmp5214, tmp5222)
return


}, 1)

tmp5223 := Call(__e, PrimFunc(symshen_4_5_1out), W2954)


__e.TailApply(tmp5213, tmp5223)
return


}


}, 1)

tmp5226 := Call(__e, PrimFunc(symshen_4_5integer_6), W2953)


__e.TailApply(tmp5212, tmp5226)
return


}, 1)

tmp5227 := Call(__e, PrimFunc(symshen_4in_1_6), W2952)


__e.TailApply(tmp5211, tmp5227)
return


}


}, 1)

tmp5230 := Call(__e, PrimFunc(symshen_4_5hash_6), W2951)


__e.TailApply(tmp5210, tmp5230)
return


}, 1)

tmp5231 := Call(__e, PrimFunc(symshen_4in_1_6), W2950)


__e.TailApply(tmp5209, tmp5231)
return


}


}, 1)

tmp5234 := Call(__e, PrimFunc(symshen_4_5lowC_6), V2948)


tmp5235 := Call(__e, tmp5208, tmp5234)


__e.TailApply(tmp5205, tmp5235)
return


}, 1)

tmp5236 := Call(__e, ns2_1set, symshen_4_5control_6, tmp5204)


_ = tmp5236

tmp5237 := MakeNative(func(__e *ControlFlow) {
V2959 := __e.Get(1)
_ = V2959
tmp5238 := MakeNative(func(__e *ControlFlow) {
W2960 := __e.Get(1)
_ = W2960
tmp5240 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2960)


if True == tmp5240 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2960)
return
}


}, 1)

tmp5252 := PrimIsPair(V2959)

var ifres5241 Obj

if True == tmp5252 {
tmp5242 := MakeNative(func(__e *ControlFlow) {
W2961 := __e.Get(1)
_ = W2961
tmp5243 := MakeNative(func(__e *ControlFlow) {
W2962 := __e.Get(1)
_ = W2962
tmp5246 := PrimEqual(W2961, MakeNumber(34))

tmp5247 := PrimNot(tmp5246)

if True == tmp5247 {
tmp5244 := PrimNumberToString(W2961)

__e.TailApply(PrimFunc(symshen_4comb), W2962, tmp5244)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp5248 := Call(__e, PrimFunc(symtail), V2959)


__e.TailApply(tmp5243, tmp5248)
return


}, 1)

tmp5249 := Call(__e, PrimFunc(symhead), V2959)


tmp5250 := Call(__e, tmp5242, tmp5249)


ifres5241 = tmp5250


} else {
tmp5251 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5241 = tmp5251


}

__e.TailApply(tmp5238, ifres5241)
return


}, 1)

tmp5253 := Call(__e, ns2_1set, symshen_4_5notdbq_6, tmp5237)


_ = tmp5253

tmp5254 := MakeNative(func(__e *ControlFlow) {
V2963 := __e.Get(1)
_ = V2963
tmp5255 := MakeNative(func(__e *ControlFlow) {
W2964 := __e.Get(1)
_ = W2964
tmp5257 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2964)


if True == tmp5257 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2964)
return
}


}, 1)

tmp5263 := Call(__e, PrimFunc(symshen_4hds_a_2), V2963, MakeNumber(99))


var ifres5258 Obj

if True == tmp5263 {
tmp5259 := MakeNative(func(__e *ControlFlow) {
W2965 := __e.Get(1)
_ = W2965
__e.TailApply(PrimFunc(symshen_4comb), W2965, symshen_4skip)
return
}, 1)

tmp5260 := Call(__e, PrimFunc(symtail), V2963)


tmp5261 := Call(__e, tmp5259, tmp5260)


ifres5258 = tmp5261


} else {
tmp5262 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5258 = tmp5262


}

__e.TailApply(tmp5255, ifres5258)
return


}, 1)

tmp5264 := Call(__e, ns2_1set, symshen_4_5lowC_6, tmp5254)


_ = tmp5264

tmp5265 := MakeNative(func(__e *ControlFlow) {
V2966 := __e.Get(1)
_ = V2966
tmp5266 := MakeNative(func(__e *ControlFlow) {
W2967 := __e.Get(1)
_ = W2967
tmp5268 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2967)


if True == tmp5268 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2967)
return
}


}, 1)

tmp5274 := Call(__e, PrimFunc(symshen_4hds_a_2), V2966, MakeNumber(35))


var ifres5269 Obj

if True == tmp5274 {
tmp5270 := MakeNative(func(__e *ControlFlow) {
W2968 := __e.Get(1)
_ = W2968
__e.TailApply(PrimFunc(symshen_4comb), W2968, symshen_4skip)
return
}, 1)

tmp5271 := Call(__e, PrimFunc(symtail), V2966)


tmp5272 := Call(__e, tmp5270, tmp5271)


ifres5269 = tmp5272


} else {
tmp5273 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5269 = tmp5273


}

__e.TailApply(tmp5266, ifres5269)
return


}, 1)

tmp5275 := Call(__e, ns2_1set, symshen_4_5hash_6, tmp5265)


_ = tmp5275

tmp5276 := MakeNative(func(__e *ControlFlow) {
V2969 := __e.Get(1)
_ = V2969
tmp5277 := MakeNative(func(__e *ControlFlow) {
W2970 := __e.Get(1)
_ = W2970
tmp5333 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2970)


if True == tmp5333 {
tmp5278 := MakeNative(func(__e *ControlFlow) {
W2976 := __e.Get(1)
_ = W2976
tmp5316 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2976)


if True == tmp5316 {
tmp5279 := MakeNative(func(__e *ControlFlow) {
W2982 := __e.Get(1)
_ = W2982
tmp5305 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2982)


if True == tmp5305 {
tmp5280 := MakeNative(func(__e *ControlFlow) {
W2986 := __e.Get(1)
_ = W2986
tmp5294 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2986)


if True == tmp5294 {
tmp5281 := MakeNative(func(__e *ControlFlow) {
W2990 := __e.Get(1)
_ = W2990
tmp5283 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2990)


if True == tmp5283 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2990)
return
}


}, 1)

tmp5284 := MakeNative(func(__e *ControlFlow) {
W2991 := __e.Get(1)
_ = W2991
tmp5290 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2991)


if True == tmp5290 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5285 := MakeNative(func(__e *ControlFlow) {
W2992 := __e.Get(1)
_ = W2992
tmp5286 := MakeNative(func(__e *ControlFlow) {
W2993 := __e.Get(1)
_ = W2993
__e.TailApply(PrimFunc(symshen_4comb), W2993, W2992)
return
}, 1)

tmp5287 := Call(__e, PrimFunc(symshen_4in_1_6), W2991)


__e.TailApply(tmp5286, tmp5287)
return


}, 1)

tmp5288 := Call(__e, PrimFunc(symshen_4_5_1out), W2991)


__e.TailApply(tmp5285, tmp5288)
return


}


}, 1)

tmp5291 := Call(__e, PrimFunc(symshen_4_5integer_6), V2969)


tmp5292 := Call(__e, tmp5284, tmp5291)


__e.TailApply(tmp5281, tmp5292)
return


} else {
__e.Return(W2986)
return
}


}, 1)

tmp5295 := MakeNative(func(__e *ControlFlow) {
W2987 := __e.Get(1)
_ = W2987
tmp5301 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2987)


if True == tmp5301 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5296 := MakeNative(func(__e *ControlFlow) {
W2988 := __e.Get(1)
_ = W2988
tmp5297 := MakeNative(func(__e *ControlFlow) {
W2989 := __e.Get(1)
_ = W2989
__e.TailApply(PrimFunc(symshen_4comb), W2989, W2988)
return
}, 1)

tmp5298 := Call(__e, PrimFunc(symshen_4in_1_6), W2987)


__e.TailApply(tmp5297, tmp5298)
return


}, 1)

tmp5299 := Call(__e, PrimFunc(symshen_4_5_1out), W2987)


__e.TailApply(tmp5296, tmp5299)
return


}


}, 1)

tmp5302 := Call(__e, PrimFunc(symshen_4_5float_6), V2969)


tmp5303 := Call(__e, tmp5295, tmp5302)


__e.TailApply(tmp5280, tmp5303)
return


} else {
__e.Return(W2982)
return
}


}, 1)

tmp5306 := MakeNative(func(__e *ControlFlow) {
W2983 := __e.Get(1)
_ = W2983
tmp5312 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2983)


if True == tmp5312 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5307 := MakeNative(func(__e *ControlFlow) {
W2984 := __e.Get(1)
_ = W2984
tmp5308 := MakeNative(func(__e *ControlFlow) {
W2985 := __e.Get(1)
_ = W2985
__e.TailApply(PrimFunc(symshen_4comb), W2985, W2984)
return
}, 1)

tmp5309 := Call(__e, PrimFunc(symshen_4in_1_6), W2983)


__e.TailApply(tmp5308, tmp5309)
return


}, 1)

tmp5310 := Call(__e, PrimFunc(symshen_4_5_1out), W2983)


__e.TailApply(tmp5307, tmp5310)
return


}


}, 1)

tmp5313 := Call(__e, PrimFunc(symshen_4_5e_1number_6), V2969)


tmp5314 := Call(__e, tmp5306, tmp5313)


__e.TailApply(tmp5279, tmp5314)
return


} else {
__e.Return(W2976)
return
}


}, 1)

tmp5317 := MakeNative(func(__e *ControlFlow) {
W2977 := __e.Get(1)
_ = W2977
tmp5329 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2977)


if True == tmp5329 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5318 := MakeNative(func(__e *ControlFlow) {
W2978 := __e.Get(1)
_ = W2978
tmp5319 := MakeNative(func(__e *ControlFlow) {
W2979 := __e.Get(1)
_ = W2979
tmp5325 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2979)


if True == tmp5325 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5320 := MakeNative(func(__e *ControlFlow) {
W2980 := __e.Get(1)
_ = W2980
tmp5321 := MakeNative(func(__e *ControlFlow) {
W2981 := __e.Get(1)
_ = W2981
__e.TailApply(PrimFunc(symshen_4comb), W2981, W2980)
return
}, 1)

tmp5322 := Call(__e, PrimFunc(symshen_4in_1_6), W2979)


__e.TailApply(tmp5321, tmp5322)
return


}, 1)

tmp5323 := Call(__e, PrimFunc(symshen_4_5_1out), W2979)


__e.TailApply(tmp5320, tmp5323)
return


}


}, 1)

tmp5326 := Call(__e, PrimFunc(symshen_4_5number_6), W2978)


__e.TailApply(tmp5319, tmp5326)
return


}, 1)

tmp5327 := Call(__e, PrimFunc(symshen_4in_1_6), W2977)


__e.TailApply(tmp5318, tmp5327)
return


}


}, 1)

tmp5330 := Call(__e, PrimFunc(symshen_4_5plus_6), V2969)


tmp5331 := Call(__e, tmp5317, tmp5330)


__e.TailApply(tmp5278, tmp5331)
return


} else {
__e.Return(W2970)
return
}


}, 1)

tmp5334 := MakeNative(func(__e *ControlFlow) {
W2971 := __e.Get(1)
_ = W2971
tmp5347 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2971)


if True == tmp5347 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5335 := MakeNative(func(__e *ControlFlow) {
W2972 := __e.Get(1)
_ = W2972
tmp5336 := MakeNative(func(__e *ControlFlow) {
W2973 := __e.Get(1)
_ = W2973
tmp5343 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2973)


if True == tmp5343 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5337 := MakeNative(func(__e *ControlFlow) {
W2974 := __e.Get(1)
_ = W2974
tmp5338 := MakeNative(func(__e *ControlFlow) {
W2975 := __e.Get(1)
_ = W2975
tmp5339 := PrimNumberSubtract(MakeNumber(0), W2974)

__e.TailApply(PrimFunc(symshen_4comb), W2975, tmp5339)
return


}, 1)

tmp5340 := Call(__e, PrimFunc(symshen_4in_1_6), W2973)


__e.TailApply(tmp5338, tmp5340)
return


}, 1)

tmp5341 := Call(__e, PrimFunc(symshen_4_5_1out), W2973)


__e.TailApply(tmp5337, tmp5341)
return


}


}, 1)

tmp5344 := Call(__e, PrimFunc(symshen_4_5number_6), W2972)


__e.TailApply(tmp5336, tmp5344)
return


}, 1)

tmp5345 := Call(__e, PrimFunc(symshen_4in_1_6), W2971)


__e.TailApply(tmp5335, tmp5345)
return


}


}, 1)

tmp5348 := Call(__e, PrimFunc(symshen_4_5minus_6), V2969)


tmp5349 := Call(__e, tmp5334, tmp5348)


__e.TailApply(tmp5277, tmp5349)
return


}, 1)

tmp5350 := Call(__e, ns2_1set, symshen_4_5number_6, tmp5276)


_ = tmp5350

tmp5351 := MakeNative(func(__e *ControlFlow) {
V2994 := __e.Get(1)
_ = V2994
tmp5352 := MakeNative(func(__e *ControlFlow) {
W2995 := __e.Get(1)
_ = W2995
tmp5354 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2995)


if True == tmp5354 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2995)
return
}


}, 1)

tmp5360 := Call(__e, PrimFunc(symshen_4hds_a_2), V2994, MakeNumber(45))


var ifres5355 Obj

if True == tmp5360 {
tmp5356 := MakeNative(func(__e *ControlFlow) {
W2996 := __e.Get(1)
_ = W2996
__e.TailApply(PrimFunc(symshen_4comb), W2996, symshen_4skip)
return
}, 1)

tmp5357 := Call(__e, PrimFunc(symtail), V2994)


tmp5358 := Call(__e, tmp5356, tmp5357)


ifres5355 = tmp5358


} else {
tmp5359 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5355 = tmp5359


}

__e.TailApply(tmp5352, ifres5355)
return


}, 1)

tmp5361 := Call(__e, ns2_1set, symshen_4_5minus_6, tmp5351)


_ = tmp5361

tmp5362 := MakeNative(func(__e *ControlFlow) {
V2997 := __e.Get(1)
_ = V2997
tmp5363 := MakeNative(func(__e *ControlFlow) {
W2998 := __e.Get(1)
_ = W2998
tmp5365 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2998)


if True == tmp5365 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2998)
return
}


}, 1)

tmp5371 := Call(__e, PrimFunc(symshen_4hds_a_2), V2997, MakeNumber(43))


var ifres5366 Obj

if True == tmp5371 {
tmp5367 := MakeNative(func(__e *ControlFlow) {
W2999 := __e.Get(1)
_ = W2999
__e.TailApply(PrimFunc(symshen_4comb), W2999, symshen_4skip)
return
}, 1)

tmp5368 := Call(__e, PrimFunc(symtail), V2997)


tmp5369 := Call(__e, tmp5367, tmp5368)


ifres5366 = tmp5369


} else {
tmp5370 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5366 = tmp5370


}

__e.TailApply(tmp5363, ifres5366)
return


}, 1)

tmp5372 := Call(__e, ns2_1set, symshen_4_5plus_6, tmp5362)


_ = tmp5372

tmp5373 := MakeNative(func(__e *ControlFlow) {
V3000 := __e.Get(1)
_ = V3000
tmp5374 := MakeNative(func(__e *ControlFlow) {
W3001 := __e.Get(1)
_ = W3001
tmp5376 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3001)


if True == tmp5376 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3001)
return
}


}, 1)

tmp5377 := MakeNative(func(__e *ControlFlow) {
W3002 := __e.Get(1)
_ = W3002
tmp5384 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3002)


if True == tmp5384 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5378 := MakeNative(func(__e *ControlFlow) {
W3003 := __e.Get(1)
_ = W3003
tmp5379 := MakeNative(func(__e *ControlFlow) {
W3004 := __e.Get(1)
_ = W3004
tmp5380 := Call(__e, PrimFunc(symshen_4compute_1integer), W3003)


__e.TailApply(PrimFunc(symshen_4comb), W3004, tmp5380)
return


}, 1)

tmp5381 := Call(__e, PrimFunc(symshen_4in_1_6), W3002)


__e.TailApply(tmp5379, tmp5381)
return


}, 1)

tmp5382 := Call(__e, PrimFunc(symshen_4_5_1out), W3002)


__e.TailApply(tmp5378, tmp5382)
return


}


}, 1)

tmp5385 := Call(__e, PrimFunc(symshen_4_5digits_6), V3000)


tmp5386 := Call(__e, tmp5377, tmp5385)


__e.TailApply(tmp5374, tmp5386)
return


}, 1)

tmp5387 := Call(__e, ns2_1set, symshen_4_5integer_6, tmp5373)


_ = tmp5387

tmp5388 := MakeNative(func(__e *ControlFlow) {
V3005 := __e.Get(1)
_ = V3005
tmp5389 := MakeNative(func(__e *ControlFlow) {
W3006 := __e.Get(1)
_ = W3006
tmp5404 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3006)


if True == tmp5404 {
tmp5390 := MakeNative(func(__e *ControlFlow) {
W3013 := __e.Get(1)
_ = W3013
tmp5392 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3013)


if True == tmp5392 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3013)
return
}


}, 1)

tmp5393 := MakeNative(func(__e *ControlFlow) {
W3014 := __e.Get(1)
_ = W3014
tmp5400 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3014)


if True == tmp5400 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5394 := MakeNative(func(__e *ControlFlow) {
W3015 := __e.Get(1)
_ = W3015
tmp5395 := MakeNative(func(__e *ControlFlow) {
W3016 := __e.Get(1)
_ = W3016
tmp5396 := PrimCons(W3015, Nil)

__e.TailApply(PrimFunc(symshen_4comb), W3016, tmp5396)
return


}, 1)

tmp5397 := Call(__e, PrimFunc(symshen_4in_1_6), W3014)


__e.TailApply(tmp5395, tmp5397)
return


}, 1)

tmp5398 := Call(__e, PrimFunc(symshen_4_5_1out), W3014)


__e.TailApply(tmp5394, tmp5398)
return


}


}, 1)

tmp5401 := Call(__e, PrimFunc(symshen_4_5digit_6), V3005)


tmp5402 := Call(__e, tmp5393, tmp5401)


__e.TailApply(tmp5390, tmp5402)
return


} else {
__e.Return(W3006)
return
}


}, 1)

tmp5405 := MakeNative(func(__e *ControlFlow) {
W3007 := __e.Get(1)
_ = W3007
tmp5420 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3007)


if True == tmp5420 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5406 := MakeNative(func(__e *ControlFlow) {
W3008 := __e.Get(1)
_ = W3008
tmp5407 := MakeNative(func(__e *ControlFlow) {
W3009 := __e.Get(1)
_ = W3009
tmp5408 := MakeNative(func(__e *ControlFlow) {
W3010 := __e.Get(1)
_ = W3010
tmp5415 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3010)


if True == tmp5415 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5409 := MakeNative(func(__e *ControlFlow) {
W3011 := __e.Get(1)
_ = W3011
tmp5410 := MakeNative(func(__e *ControlFlow) {
W3012 := __e.Get(1)
_ = W3012
tmp5411 := PrimCons(W3008, W3011)

__e.TailApply(PrimFunc(symshen_4comb), W3012, tmp5411)
return


}, 1)

tmp5412 := Call(__e, PrimFunc(symshen_4in_1_6), W3010)


__e.TailApply(tmp5410, tmp5412)
return


}, 1)

tmp5413 := Call(__e, PrimFunc(symshen_4_5_1out), W3010)


__e.TailApply(tmp5409, tmp5413)
return


}


}, 1)

tmp5416 := Call(__e, PrimFunc(symshen_4_5digits_6), W3009)


__e.TailApply(tmp5408, tmp5416)
return


}, 1)

tmp5417 := Call(__e, PrimFunc(symshen_4in_1_6), W3007)


__e.TailApply(tmp5407, tmp5417)
return


}, 1)

tmp5418 := Call(__e, PrimFunc(symshen_4_5_1out), W3007)


__e.TailApply(tmp5406, tmp5418)
return


}


}, 1)

tmp5421 := Call(__e, PrimFunc(symshen_4_5digit_6), V3005)


tmp5422 := Call(__e, tmp5405, tmp5421)


__e.TailApply(tmp5389, tmp5422)
return


}, 1)

tmp5423 := Call(__e, ns2_1set, symshen_4_5digits_6, tmp5388)


_ = tmp5423

tmp5424 := MakeNative(func(__e *ControlFlow) {
V3017 := __e.Get(1)
_ = V3017
tmp5425 := MakeNative(func(__e *ControlFlow) {
W3018 := __e.Get(1)
_ = W3018
tmp5427 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3018)


if True == tmp5427 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3018)
return
}


}, 1)

tmp5438 := PrimIsPair(V3017)

var ifres5428 Obj

if True == tmp5438 {
tmp5429 := MakeNative(func(__e *ControlFlow) {
W3019 := __e.Get(1)
_ = W3019
tmp5430 := MakeNative(func(__e *ControlFlow) {
W3020 := __e.Get(1)
_ = W3020
tmp5433 := Call(__e, PrimFunc(symshen_4digit_2), W3019)


if True == tmp5433 {
tmp5431 := Call(__e, PrimFunc(symshen_4byte_1_6digit), W3019)


__e.TailApply(PrimFunc(symshen_4comb), W3020, tmp5431)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp5434 := Call(__e, PrimFunc(symtail), V3017)


__e.TailApply(tmp5430, tmp5434)
return


}, 1)

tmp5435 := Call(__e, PrimFunc(symhead), V3017)


tmp5436 := Call(__e, tmp5429, tmp5435)


ifres5428 = tmp5436


} else {
tmp5437 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5428 = tmp5437


}

__e.TailApply(tmp5425, ifres5428)
return


}, 1)

tmp5439 := Call(__e, ns2_1set, symshen_4_5digit_6, tmp5424)


_ = tmp5439

tmp5440 := MakeNative(func(__e *ControlFlow) {
V3021 := __e.Get(1)
_ = V3021
__e.Return(PrimNumberSubtract(V3021, MakeNumber(48)))
return
}, 1)

tmp5441 := Call(__e, ns2_1set, symshen_4byte_1_6digit, tmp5440)


_ = tmp5441

tmp5442 := MakeNative(func(__e *ControlFlow) {
V3022 := __e.Get(1)
_ = V3022
tmp5443 := Call(__e, PrimFunc(symreverse), V3022)


__e.TailApply(PrimFunc(symshen_4compute_1integer_1h), tmp5443, MakeNumber(0))
return


}, 1)

tmp5444 := Call(__e, ns2_1set, symshen_4compute_1integer, tmp5442)


_ = tmp5444

tmp5445 := MakeNative(func(__e *ControlFlow) {
V3025 := __e.Get(1)
_ = V3025
V3026 := __e.Get(2)
_ = V3026
tmp5455 := PrimEqual(Nil, V3025)

if True == tmp5455 {
__e.Return(MakeNumber(0))
return
} else {
tmp5453 := PrimIsPair(V3025)

if True == tmp5453 {
tmp5446 := Call(__e, PrimFunc(symshen_4expt), MakeNumber(10), V3026)


tmp5447 := PrimHead(V3025)

tmp5448 := PrimNumberMultiply(tmp5446, tmp5447)

tmp5449 := PrimTail(V3025)

tmp5450 := PrimNumberAdd(V3026, MakeNumber(1))

tmp5451 := Call(__e, PrimFunc(symshen_4compute_1integer_1h), tmp5449, tmp5450)


__e.Return(PrimNumberAdd(tmp5448, tmp5451))
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4compute_1integer_1h)
return
}


}


}, 2)

tmp5456 := Call(__e, ns2_1set, symshen_4compute_1integer_1h, tmp5445)


_ = tmp5456

tmp5457 := MakeNative(func(__e *ControlFlow) {
V3029 := __e.Get(1)
_ = V3029
V3030 := __e.Get(2)
_ = V3030
tmp5465 := PrimEqual(MakeNumber(0), V3030)

if True == tmp5465 {
__e.Return(MakeNumber(1))
return
} else {
tmp5463 := PrimGreatThan(V3030, MakeNumber(0))

if True == tmp5463 {
tmp5458 := PrimNumberSubtract(V3030, MakeNumber(1))

tmp5459 := Call(__e, PrimFunc(symshen_4expt), V3029, tmp5458)


__e.Return(PrimNumberMultiply(V3029, tmp5459))
return


} else {
tmp5460 := PrimNumberAdd(V3030, MakeNumber(1))

tmp5461 := Call(__e, PrimFunc(symshen_4expt), V3029, tmp5460)


__e.Return(PrimNumberDivide(tmp5461, V3029))
return


}


}


}, 2)

tmp5466 := Call(__e, ns2_1set, symshen_4expt, tmp5457)


_ = tmp5466

tmp5467 := MakeNative(func(__e *ControlFlow) {
V3031 := __e.Get(1)
_ = V3031
tmp5468 := MakeNative(func(__e *ControlFlow) {
W3032 := __e.Get(1)
_ = W3032
tmp5488 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3032)


if True == tmp5488 {
tmp5469 := MakeNative(func(__e *ControlFlow) {
W3041 := __e.Get(1)
_ = W3041
tmp5471 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3041)


if True == tmp5471 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3041)
return
}


}, 1)

tmp5472 := MakeNative(func(__e *ControlFlow) {
W3042 := __e.Get(1)
_ = W3042
tmp5484 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3042)


if True == tmp5484 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5473 := MakeNative(func(__e *ControlFlow) {
W3043 := __e.Get(1)
_ = W3043
tmp5474 := MakeNative(func(__e *ControlFlow) {
W3044 := __e.Get(1)
_ = W3044
tmp5480 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3044)


if True == tmp5480 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5475 := MakeNative(func(__e *ControlFlow) {
W3045 := __e.Get(1)
_ = W3045
tmp5476 := MakeNative(func(__e *ControlFlow) {
W3046 := __e.Get(1)
_ = W3046
__e.TailApply(PrimFunc(symshen_4comb), W3046, W3045)
return
}, 1)

tmp5477 := Call(__e, PrimFunc(symshen_4in_1_6), W3044)


__e.TailApply(tmp5476, tmp5477)
return


}, 1)

tmp5478 := Call(__e, PrimFunc(symshen_4_5_1out), W3044)


__e.TailApply(tmp5475, tmp5478)
return


}


}, 1)

tmp5481 := Call(__e, PrimFunc(symshen_4_5fraction_6), W3043)


__e.TailApply(tmp5474, tmp5481)
return


}, 1)

tmp5482 := Call(__e, PrimFunc(symshen_4in_1_6), W3042)


__e.TailApply(tmp5473, tmp5482)
return


}


}, 1)

tmp5485 := Call(__e, PrimFunc(symshen_4_5stop_6), V3031)


tmp5486 := Call(__e, tmp5472, tmp5485)


__e.TailApply(tmp5469, tmp5486)
return


} else {
__e.Return(W3032)
return
}


}, 1)

tmp5489 := MakeNative(func(__e *ControlFlow) {
W3033 := __e.Get(1)
_ = W3033
tmp5510 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3033)


if True == tmp5510 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5490 := MakeNative(func(__e *ControlFlow) {
W3034 := __e.Get(1)
_ = W3034
tmp5491 := MakeNative(func(__e *ControlFlow) {
W3035 := __e.Get(1)
_ = W3035
tmp5492 := MakeNative(func(__e *ControlFlow) {
W3036 := __e.Get(1)
_ = W3036
tmp5505 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3036)


if True == tmp5505 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5493 := MakeNative(func(__e *ControlFlow) {
W3037 := __e.Get(1)
_ = W3037
tmp5494 := MakeNative(func(__e *ControlFlow) {
W3038 := __e.Get(1)
_ = W3038
tmp5501 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3038)


if True == tmp5501 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5495 := MakeNative(func(__e *ControlFlow) {
W3039 := __e.Get(1)
_ = W3039
tmp5496 := MakeNative(func(__e *ControlFlow) {
W3040 := __e.Get(1)
_ = W3040
tmp5497 := PrimNumberAdd(W3034, W3039)

__e.TailApply(PrimFunc(symshen_4comb), W3040, tmp5497)
return


}, 1)

tmp5498 := Call(__e, PrimFunc(symshen_4in_1_6), W3038)


__e.TailApply(tmp5496, tmp5498)
return


}, 1)

tmp5499 := Call(__e, PrimFunc(symshen_4_5_1out), W3038)


__e.TailApply(tmp5495, tmp5499)
return


}


}, 1)

tmp5502 := Call(__e, PrimFunc(symshen_4_5fraction_6), W3037)


__e.TailApply(tmp5494, tmp5502)
return


}, 1)

tmp5503 := Call(__e, PrimFunc(symshen_4in_1_6), W3036)


__e.TailApply(tmp5493, tmp5503)
return


}


}, 1)

tmp5506 := Call(__e, PrimFunc(symshen_4_5stop_6), W3035)


__e.TailApply(tmp5492, tmp5506)
return


}, 1)

tmp5507 := Call(__e, PrimFunc(symshen_4in_1_6), W3033)


__e.TailApply(tmp5491, tmp5507)
return


}, 1)

tmp5508 := Call(__e, PrimFunc(symshen_4_5_1out), W3033)


__e.TailApply(tmp5490, tmp5508)
return


}


}, 1)

tmp5511 := Call(__e, PrimFunc(symshen_4_5integer_6), V3031)


tmp5512 := Call(__e, tmp5489, tmp5511)


__e.TailApply(tmp5468, tmp5512)
return


}, 1)

tmp5513 := Call(__e, ns2_1set, symshen_4_5float_6, tmp5467)


_ = tmp5513

tmp5514 := MakeNative(func(__e *ControlFlow) {
V3047 := __e.Get(1)
_ = V3047
tmp5515 := MakeNative(func(__e *ControlFlow) {
W3048 := __e.Get(1)
_ = W3048
tmp5517 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3048)


if True == tmp5517 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3048)
return
}


}, 1)

tmp5523 := Call(__e, PrimFunc(symshen_4hds_a_2), V3047, MakeNumber(46))


var ifres5518 Obj

if True == tmp5523 {
tmp5519 := MakeNative(func(__e *ControlFlow) {
W3049 := __e.Get(1)
_ = W3049
__e.TailApply(PrimFunc(symshen_4comb), W3049, symshen_4skip)
return
}, 1)

tmp5520 := Call(__e, PrimFunc(symtail), V3047)


tmp5521 := Call(__e, tmp5519, tmp5520)


ifres5518 = tmp5521


} else {
tmp5522 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5518 = tmp5522


}

__e.TailApply(tmp5515, ifres5518)
return


}, 1)

tmp5524 := Call(__e, ns2_1set, symshen_4_5stop_6, tmp5514)


_ = tmp5524

tmp5525 := MakeNative(func(__e *ControlFlow) {
V3050 := __e.Get(1)
_ = V3050
tmp5526 := MakeNative(func(__e *ControlFlow) {
W3051 := __e.Get(1)
_ = W3051
tmp5528 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3051)


if True == tmp5528 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3051)
return
}


}, 1)

tmp5529 := MakeNative(func(__e *ControlFlow) {
W3052 := __e.Get(1)
_ = W3052
tmp5536 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3052)


if True == tmp5536 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5530 := MakeNative(func(__e *ControlFlow) {
W3053 := __e.Get(1)
_ = W3053
tmp5531 := MakeNative(func(__e *ControlFlow) {
W3054 := __e.Get(1)
_ = W3054
tmp5532 := Call(__e, PrimFunc(symshen_4compute_1fraction), W3053)


__e.TailApply(PrimFunc(symshen_4comb), W3054, tmp5532)
return


}, 1)

tmp5533 := Call(__e, PrimFunc(symshen_4in_1_6), W3052)


__e.TailApply(tmp5531, tmp5533)
return


}, 1)

tmp5534 := Call(__e, PrimFunc(symshen_4_5_1out), W3052)


__e.TailApply(tmp5530, tmp5534)
return


}


}, 1)

tmp5537 := Call(__e, PrimFunc(symshen_4_5digits_6), V3050)


tmp5538 := Call(__e, tmp5529, tmp5537)


__e.TailApply(tmp5526, tmp5538)
return


}, 1)

tmp5539 := Call(__e, ns2_1set, symshen_4_5fraction_6, tmp5525)


_ = tmp5539

tmp5540 := MakeNative(func(__e *ControlFlow) {
V3055 := __e.Get(1)
_ = V3055
__e.TailApply(PrimFunc(symshen_4compute_1fraction_1h), V3055, MakeNumber(-1))
return
}, 1)

tmp5541 := Call(__e, ns2_1set, symshen_4compute_1fraction, tmp5540)


_ = tmp5541

tmp5542 := MakeNative(func(__e *ControlFlow) {
V3058 := __e.Get(1)
_ = V3058
V3059 := __e.Get(2)
_ = V3059
tmp5552 := PrimEqual(Nil, V3058)

if True == tmp5552 {
__e.Return(MakeNumber(0))
return
} else {
tmp5550 := PrimIsPair(V3058)

if True == tmp5550 {
tmp5543 := Call(__e, PrimFunc(symshen_4expt), MakeNumber(10), V3059)


tmp5544 := PrimHead(V3058)

tmp5545 := PrimNumberMultiply(tmp5543, tmp5544)

tmp5546 := PrimTail(V3058)

tmp5547 := PrimNumberSubtract(V3059, MakeNumber(1))

tmp5548 := Call(__e, PrimFunc(symshen_4compute_1fraction_1h), tmp5546, tmp5547)


__e.Return(PrimNumberAdd(tmp5545, tmp5548))
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4compute_1fraction_1h)
return
}


}


}, 2)

tmp5553 := Call(__e, ns2_1set, symshen_4compute_1fraction_1h, tmp5542)


_ = tmp5553

tmp5554 := MakeNative(func(__e *ControlFlow) {
V3060 := __e.Get(1)
_ = V3060
tmp5555 := MakeNative(func(__e *ControlFlow) {
W3061 := __e.Get(1)
_ = W3061
tmp5584 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3061)


if True == tmp5584 {
tmp5556 := MakeNative(func(__e *ControlFlow) {
W3070 := __e.Get(1)
_ = W3070
tmp5558 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3070)


if True == tmp5558 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3070)
return
}


}, 1)

tmp5559 := MakeNative(func(__e *ControlFlow) {
W3071 := __e.Get(1)
_ = W3071
tmp5580 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3071)


if True == tmp5580 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5560 := MakeNative(func(__e *ControlFlow) {
W3072 := __e.Get(1)
_ = W3072
tmp5561 := MakeNative(func(__e *ControlFlow) {
W3073 := __e.Get(1)
_ = W3073
tmp5562 := MakeNative(func(__e *ControlFlow) {
W3074 := __e.Get(1)
_ = W3074
tmp5575 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3074)


if True == tmp5575 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5563 := MakeNative(func(__e *ControlFlow) {
W3075 := __e.Get(1)
_ = W3075
tmp5564 := MakeNative(func(__e *ControlFlow) {
W3076 := __e.Get(1)
_ = W3076
tmp5571 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3076)


if True == tmp5571 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5565 := MakeNative(func(__e *ControlFlow) {
W3077 := __e.Get(1)
_ = W3077
tmp5566 := MakeNative(func(__e *ControlFlow) {
W3078 := __e.Get(1)
_ = W3078
tmp5567 := Call(__e, PrimFunc(symshen_4compute_1E), W3072, W3077)


__e.TailApply(PrimFunc(symshen_4comb), W3078, tmp5567)
return


}, 1)

tmp5568 := Call(__e, PrimFunc(symshen_4in_1_6), W3076)


__e.TailApply(tmp5566, tmp5568)
return


}, 1)

tmp5569 := Call(__e, PrimFunc(symshen_4_5_1out), W3076)


__e.TailApply(tmp5565, tmp5569)
return


}


}, 1)

tmp5572 := Call(__e, PrimFunc(symshen_4_5log10_6), W3075)


__e.TailApply(tmp5564, tmp5572)
return


}, 1)

tmp5573 := Call(__e, PrimFunc(symshen_4in_1_6), W3074)


__e.TailApply(tmp5563, tmp5573)
return


}


}, 1)

tmp5576 := Call(__e, PrimFunc(symshen_4_5lowE_6), W3073)


__e.TailApply(tmp5562, tmp5576)
return


}, 1)

tmp5577 := Call(__e, PrimFunc(symshen_4in_1_6), W3071)


__e.TailApply(tmp5561, tmp5577)
return


}, 1)

tmp5578 := Call(__e, PrimFunc(symshen_4_5_1out), W3071)


__e.TailApply(tmp5560, tmp5578)
return


}


}, 1)

tmp5581 := Call(__e, PrimFunc(symshen_4_5integer_6), V3060)


tmp5582 := Call(__e, tmp5559, tmp5581)


__e.TailApply(tmp5556, tmp5582)
return


} else {
__e.Return(W3061)
return
}


}, 1)

tmp5585 := MakeNative(func(__e *ControlFlow) {
W3062 := __e.Get(1)
_ = W3062
tmp5606 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3062)


if True == tmp5606 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5586 := MakeNative(func(__e *ControlFlow) {
W3063 := __e.Get(1)
_ = W3063
tmp5587 := MakeNative(func(__e *ControlFlow) {
W3064 := __e.Get(1)
_ = W3064
tmp5588 := MakeNative(func(__e *ControlFlow) {
W3065 := __e.Get(1)
_ = W3065
tmp5601 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3065)


if True == tmp5601 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5589 := MakeNative(func(__e *ControlFlow) {
W3066 := __e.Get(1)
_ = W3066
tmp5590 := MakeNative(func(__e *ControlFlow) {
W3067 := __e.Get(1)
_ = W3067
tmp5597 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3067)


if True == tmp5597 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5591 := MakeNative(func(__e *ControlFlow) {
W3068 := __e.Get(1)
_ = W3068
tmp5592 := MakeNative(func(__e *ControlFlow) {
W3069 := __e.Get(1)
_ = W3069
tmp5593 := Call(__e, PrimFunc(symshen_4compute_1E), W3063, W3068)


__e.TailApply(PrimFunc(symshen_4comb), W3069, tmp5593)
return


}, 1)

tmp5594 := Call(__e, PrimFunc(symshen_4in_1_6), W3067)


__e.TailApply(tmp5592, tmp5594)
return


}, 1)

tmp5595 := Call(__e, PrimFunc(symshen_4_5_1out), W3067)


__e.TailApply(tmp5591, tmp5595)
return


}


}, 1)

tmp5598 := Call(__e, PrimFunc(symshen_4_5log10_6), W3066)


__e.TailApply(tmp5590, tmp5598)
return


}, 1)

tmp5599 := Call(__e, PrimFunc(symshen_4in_1_6), W3065)


__e.TailApply(tmp5589, tmp5599)
return


}


}, 1)

tmp5602 := Call(__e, PrimFunc(symshen_4_5lowE_6), W3064)


__e.TailApply(tmp5588, tmp5602)
return


}, 1)

tmp5603 := Call(__e, PrimFunc(symshen_4in_1_6), W3062)


__e.TailApply(tmp5587, tmp5603)
return


}, 1)

tmp5604 := Call(__e, PrimFunc(symshen_4_5_1out), W3062)


__e.TailApply(tmp5586, tmp5604)
return


}


}, 1)

tmp5607 := Call(__e, PrimFunc(symshen_4_5float_6), V3060)


tmp5608 := Call(__e, tmp5585, tmp5607)


__e.TailApply(tmp5555, tmp5608)
return


}, 1)

tmp5609 := Call(__e, ns2_1set, symshen_4_5e_1number_6, tmp5554)


_ = tmp5609

tmp5610 := MakeNative(func(__e *ControlFlow) {
V3079 := __e.Get(1)
_ = V3079
tmp5611 := MakeNative(func(__e *ControlFlow) {
W3080 := __e.Get(1)
_ = W3080
tmp5644 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3080)


if True == tmp5644 {
tmp5612 := MakeNative(func(__e *ControlFlow) {
W3086 := __e.Get(1)
_ = W3086
tmp5626 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3086)


if True == tmp5626 {
tmp5613 := MakeNative(func(__e *ControlFlow) {
W3092 := __e.Get(1)
_ = W3092
tmp5615 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3092)


if True == tmp5615 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3092)
return
}


}, 1)

tmp5616 := MakeNative(func(__e *ControlFlow) {
W3093 := __e.Get(1)
_ = W3093
tmp5622 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3093)


if True == tmp5622 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5617 := MakeNative(func(__e *ControlFlow) {
W3094 := __e.Get(1)
_ = W3094
tmp5618 := MakeNative(func(__e *ControlFlow) {
W3095 := __e.Get(1)
_ = W3095
__e.TailApply(PrimFunc(symshen_4comb), W3095, W3094)
return
}, 1)

tmp5619 := Call(__e, PrimFunc(symshen_4in_1_6), W3093)


__e.TailApply(tmp5618, tmp5619)
return


}, 1)

tmp5620 := Call(__e, PrimFunc(symshen_4_5_1out), W3093)


__e.TailApply(tmp5617, tmp5620)
return


}


}, 1)

tmp5623 := Call(__e, PrimFunc(symshen_4_5integer_6), V3079)


tmp5624 := Call(__e, tmp5616, tmp5623)


__e.TailApply(tmp5613, tmp5624)
return


} else {
__e.Return(W3086)
return
}


}, 1)

tmp5627 := MakeNative(func(__e *ControlFlow) {
W3087 := __e.Get(1)
_ = W3087
tmp5640 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3087)


if True == tmp5640 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5628 := MakeNative(func(__e *ControlFlow) {
W3088 := __e.Get(1)
_ = W3088
tmp5629 := MakeNative(func(__e *ControlFlow) {
W3089 := __e.Get(1)
_ = W3089
tmp5636 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3089)


if True == tmp5636 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5630 := MakeNative(func(__e *ControlFlow) {
W3090 := __e.Get(1)
_ = W3090
tmp5631 := MakeNative(func(__e *ControlFlow) {
W3091 := __e.Get(1)
_ = W3091
tmp5632 := PrimNumberSubtract(MakeNumber(0), W3090)

__e.TailApply(PrimFunc(symshen_4comb), W3091, tmp5632)
return


}, 1)

tmp5633 := Call(__e, PrimFunc(symshen_4in_1_6), W3089)


__e.TailApply(tmp5631, tmp5633)
return


}, 1)

tmp5634 := Call(__e, PrimFunc(symshen_4_5_1out), W3089)


__e.TailApply(tmp5630, tmp5634)
return


}


}, 1)

tmp5637 := Call(__e, PrimFunc(symshen_4_5log10_6), W3088)


__e.TailApply(tmp5629, tmp5637)
return


}, 1)

tmp5638 := Call(__e, PrimFunc(symshen_4in_1_6), W3087)


__e.TailApply(tmp5628, tmp5638)
return


}


}, 1)

tmp5641 := Call(__e, PrimFunc(symshen_4_5minus_6), V3079)


tmp5642 := Call(__e, tmp5627, tmp5641)


__e.TailApply(tmp5612, tmp5642)
return


} else {
__e.Return(W3080)
return
}


}, 1)

tmp5645 := MakeNative(func(__e *ControlFlow) {
W3081 := __e.Get(1)
_ = W3081
tmp5657 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3081)


if True == tmp5657 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5646 := MakeNative(func(__e *ControlFlow) {
W3082 := __e.Get(1)
_ = W3082
tmp5647 := MakeNative(func(__e *ControlFlow) {
W3083 := __e.Get(1)
_ = W3083
tmp5653 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3083)


if True == tmp5653 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5648 := MakeNative(func(__e *ControlFlow) {
W3084 := __e.Get(1)
_ = W3084
tmp5649 := MakeNative(func(__e *ControlFlow) {
W3085 := __e.Get(1)
_ = W3085
__e.TailApply(PrimFunc(symshen_4comb), W3085, W3084)
return
}, 1)

tmp5650 := Call(__e, PrimFunc(symshen_4in_1_6), W3083)


__e.TailApply(tmp5649, tmp5650)
return


}, 1)

tmp5651 := Call(__e, PrimFunc(symshen_4_5_1out), W3083)


__e.TailApply(tmp5648, tmp5651)
return


}


}, 1)

tmp5654 := Call(__e, PrimFunc(symshen_4_5log10_6), W3082)


__e.TailApply(tmp5647, tmp5654)
return


}, 1)

tmp5655 := Call(__e, PrimFunc(symshen_4in_1_6), W3081)


__e.TailApply(tmp5646, tmp5655)
return


}


}, 1)

tmp5658 := Call(__e, PrimFunc(symshen_4_5plus_6), V3079)


tmp5659 := Call(__e, tmp5645, tmp5658)


__e.TailApply(tmp5611, tmp5659)
return


}, 1)

tmp5660 := Call(__e, ns2_1set, symshen_4_5log10_6, tmp5610)


_ = tmp5660

tmp5661 := MakeNative(func(__e *ControlFlow) {
V3096 := __e.Get(1)
_ = V3096
tmp5662 := MakeNative(func(__e *ControlFlow) {
W3097 := __e.Get(1)
_ = W3097
tmp5664 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3097)


if True == tmp5664 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3097)
return
}


}, 1)

tmp5670 := Call(__e, PrimFunc(symshen_4hds_a_2), V3096, MakeNumber(101))


var ifres5665 Obj

if True == tmp5670 {
tmp5666 := MakeNative(func(__e *ControlFlow) {
W3098 := __e.Get(1)
_ = W3098
__e.TailApply(PrimFunc(symshen_4comb), W3098, symshen_4skip)
return
}, 1)

tmp5667 := Call(__e, PrimFunc(symtail), V3096)


tmp5668 := Call(__e, tmp5666, tmp5667)


ifres5665 = tmp5668


} else {
tmp5669 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5665 = tmp5669


}

__e.TailApply(tmp5662, ifres5665)
return


}, 1)

tmp5671 := Call(__e, ns2_1set, symshen_4_5lowE_6, tmp5661)


_ = tmp5671

tmp5672 := MakeNative(func(__e *ControlFlow) {
V3099 := __e.Get(1)
_ = V3099
V3100 := __e.Get(2)
_ = V3100
tmp5673 := Call(__e, PrimFunc(symshen_4expt), MakeNumber(10), V3100)


__e.Return(PrimNumberMultiply(V3099, tmp5673))
return


}, 2)

tmp5674 := Call(__e, ns2_1set, symshen_4compute_1E, tmp5672)


_ = tmp5674

tmp5675 := MakeNative(func(__e *ControlFlow) {
V3101 := __e.Get(1)
_ = V3101
tmp5676 := MakeNative(func(__e *ControlFlow) {
W3102 := __e.Get(1)
_ = W3102
tmp5688 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3102)


if True == tmp5688 {
tmp5677 := MakeNative(func(__e *ControlFlow) {
W3107 := __e.Get(1)
_ = W3107
tmp5679 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3107)


if True == tmp5679 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3107)
return
}


}, 1)

tmp5680 := MakeNative(func(__e *ControlFlow) {
W3108 := __e.Get(1)
_ = W3108
tmp5684 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3108)


if True == tmp5684 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5681 := MakeNative(func(__e *ControlFlow) {
W3109 := __e.Get(1)
_ = W3109
__e.TailApply(PrimFunc(symshen_4comb), W3109, symshen_4skip)
return
}, 1)

tmp5682 := Call(__e, PrimFunc(symshen_4in_1_6), W3108)


__e.TailApply(tmp5681, tmp5682)
return


}


}, 1)

tmp5685 := Call(__e, PrimFunc(symshen_4_5whitespace_6), V3101)


tmp5686 := Call(__e, tmp5680, tmp5685)


__e.TailApply(tmp5677, tmp5686)
return


} else {
__e.Return(W3102)
return
}


}, 1)

tmp5689 := MakeNative(func(__e *ControlFlow) {
W3103 := __e.Get(1)
_ = W3103
tmp5699 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3103)


if True == tmp5699 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5690 := MakeNative(func(__e *ControlFlow) {
W3104 := __e.Get(1)
_ = W3104
tmp5691 := MakeNative(func(__e *ControlFlow) {
W3105 := __e.Get(1)
_ = W3105
tmp5695 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3105)


if True == tmp5695 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5692 := MakeNative(func(__e *ControlFlow) {
W3106 := __e.Get(1)
_ = W3106
__e.TailApply(PrimFunc(symshen_4comb), W3106, symshen_4skip)
return
}, 1)

tmp5693 := Call(__e, PrimFunc(symshen_4in_1_6), W3105)


__e.TailApply(tmp5692, tmp5693)
return


}


}, 1)

tmp5696 := Call(__e, PrimFunc(symshen_4_5whitespaces_6), W3104)


__e.TailApply(tmp5691, tmp5696)
return


}, 1)

tmp5697 := Call(__e, PrimFunc(symshen_4in_1_6), W3103)


__e.TailApply(tmp5690, tmp5697)
return


}


}, 1)

tmp5700 := Call(__e, PrimFunc(symshen_4_5whitespace_6), V3101)


tmp5701 := Call(__e, tmp5689, tmp5700)


__e.TailApply(tmp5676, tmp5701)
return


}, 1)

tmp5702 := Call(__e, ns2_1set, symshen_4_5whitespaces_6, tmp5675)


_ = tmp5702

tmp5703 := MakeNative(func(__e *ControlFlow) {
V3110 := __e.Get(1)
_ = V3110
tmp5704 := MakeNative(func(__e *ControlFlow) {
W3111 := __e.Get(1)
_ = W3111
tmp5706 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W3111)


if True == tmp5706 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W3111)
return
}


}, 1)

tmp5716 := PrimIsPair(V3110)

var ifres5707 Obj

if True == tmp5716 {
tmp5708 := MakeNative(func(__e *ControlFlow) {
W3112 := __e.Get(1)
_ = W3112
tmp5709 := MakeNative(func(__e *ControlFlow) {
W3113 := __e.Get(1)
_ = W3113
tmp5711 := Call(__e, PrimFunc(symshen_4whitespace_2), W3112)


if True == tmp5711 {
__e.TailApply(PrimFunc(symshen_4comb), W3113, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp5712 := Call(__e, PrimFunc(symtail), V3110)


__e.TailApply(tmp5709, tmp5712)
return


}, 1)

tmp5713 := Call(__e, PrimFunc(symhead), V3110)


tmp5714 := Call(__e, tmp5708, tmp5713)


ifres5707 = tmp5714


} else {
tmp5715 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5707 = tmp5715


}

__e.TailApply(tmp5704, ifres5707)
return


}, 1)

tmp5717 := Call(__e, ns2_1set, symshen_4_5whitespace_6, tmp5703)


_ = tmp5717

tmp5718 := MakeNative(func(__e *ControlFlow) {
V3116 := __e.Get(1)
_ = V3116
tmp5726 := PrimEqual(MakeNumber(32), V3116)

if True == tmp5726 {
__e.Return(True)
return
} else {
tmp5724 := PrimEqual(MakeNumber(13), V3116)

if True == tmp5724 {
__e.Return(True)
return
} else {
tmp5722 := PrimEqual(MakeNumber(10), V3116)

if True == tmp5722 {
__e.Return(True)
return
} else {
tmp5720 := PrimEqual(MakeNumber(9), V3116)

if True == tmp5720 {
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

tmp5727 := Call(__e, ns2_1set, symshen_4whitespace_2, tmp5718)


_ = tmp5727

tmp5728 := MakeNative(func(__e *ControlFlow) {
V3117 := __e.Get(1)
_ = V3117
tmp5751 := PrimEqual(Nil, V3117)

if True == tmp5751 {
__e.Return(Nil)
return
} else {
tmp5749 := PrimIsPair(V3117)

var ifres5745 Obj

if True == tmp5749 {
tmp5747 := PrimHead(V3117)

tmp5748 := Call(__e, PrimFunc(symshen_4packaged_2), tmp5747)


var ifres5746 Obj

if True == tmp5748 {
ifres5746 = True


} else {
ifres5746 = False


}

ifres5745 = ifres5746


} else {
ifres5745 = False


}

if True == ifres5745 {
tmp5729 := PrimHead(V3117)

tmp5730 := Call(__e, PrimFunc(symshen_4unpackage), tmp5729)


tmp5731 := PrimTail(V3117)

tmp5732 := Call(__e, PrimFunc(symappend), tmp5730, tmp5731)


__e.TailApply(PrimFunc(symshen_4unpackage_emacroexpand), tmp5732)
return


} else {
tmp5743 := PrimIsPair(V3117)

if True == tmp5743 {
tmp5733 := MakeNative(func(__e *ControlFlow) {
W3118 := __e.Get(1)
_ = W3118
tmp5739 := Call(__e, PrimFunc(symshen_4packaged_2), W3118)


if True == tmp5739 {
tmp5734 := PrimTail(V3117)

tmp5735 := PrimCons(W3118, tmp5734)

__e.TailApply(PrimFunc(symshen_4unpackage_emacroexpand), tmp5735)
return


} else {
tmp5736 := PrimTail(V3117)

tmp5737 := Call(__e, PrimFunc(symshen_4unpackage_emacroexpand), tmp5736)


__e.Return(PrimCons(W3118, tmp5737))
return


}


}, 1)

tmp5740 := PrimHead(V3117)

tmp5741 := Call(__e, PrimFunc(symmacroexpand), tmp5740)


__e.TailApply(tmp5733, tmp5741)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4unpackage_emacroexpand)
return
}


}


}


}, 1)

tmp5752 := Call(__e, ns2_1set, symshen_4unpackage_emacroexpand, tmp5728)


_ = tmp5752

tmp5753 := MakeNative(func(__e *ControlFlow) {
V3121 := __e.Get(1)
_ = V3121
tmp5768 := PrimIsPair(V3121)

var ifres5755 Obj

if True == tmp5768 {
tmp5766 := PrimHead(V3121)

tmp5767 := PrimEqual(sympackage, tmp5766)

var ifres5757 Obj

if True == tmp5767 {
tmp5764 := PrimTail(V3121)

tmp5765 := PrimIsPair(tmp5764)

var ifres5759 Obj

if True == tmp5765 {
tmp5761 := PrimTail(V3121)

tmp5762 := PrimTail(tmp5761)

tmp5763 := PrimIsPair(tmp5762)

var ifres5760 Obj

if True == tmp5763 {
ifres5760 = True


} else {
ifres5760 = False


}

ifres5759 = ifres5760


} else {
ifres5759 = False


}

var ifres5758 Obj

if True == ifres5759 {
ifres5758 = True


} else {
ifres5758 = False


}

ifres5757 = ifres5758


} else {
ifres5757 = False


}

var ifres5756 Obj

if True == ifres5757 {
ifres5756 = True


} else {
ifres5756 = False


}

ifres5755 = ifres5756


} else {
ifres5755 = False


}

if True == ifres5755 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp5769 := Call(__e, ns2_1set, symshen_4packaged_2, tmp5753)


_ = tmp5769

tmp5770 := MakeNative(func(__e *ControlFlow) {
V3124 := __e.Get(1)
_ = V3124
tmp5831 := PrimIsPair(V3124)

var ifres5813 Obj

if True == tmp5831 {
tmp5829 := PrimHead(V3124)

tmp5830 := PrimEqual(sympackage, tmp5829)

var ifres5815 Obj

if True == tmp5830 {
tmp5827 := PrimTail(V3124)

tmp5828 := PrimIsPair(tmp5827)

var ifres5817 Obj

if True == tmp5828 {
tmp5824 := PrimTail(V3124)

tmp5825 := PrimHead(tmp5824)

tmp5826 := PrimEqual(symnull, tmp5825)

var ifres5819 Obj

if True == tmp5826 {
tmp5821 := PrimTail(V3124)

tmp5822 := PrimTail(tmp5821)

tmp5823 := PrimIsPair(tmp5822)

var ifres5820 Obj

if True == tmp5823 {
ifres5820 = True


} else {
ifres5820 = False


}

ifres5819 = ifres5820


} else {
ifres5819 = False


}

var ifres5818 Obj

if True == ifres5819 {
ifres5818 = True


} else {
ifres5818 = False


}

ifres5817 = ifres5818


} else {
ifres5817 = False


}

var ifres5816 Obj

if True == ifres5817 {
ifres5816 = True


} else {
ifres5816 = False


}

ifres5815 = ifres5816


} else {
ifres5815 = False


}

var ifres5814 Obj

if True == ifres5815 {
ifres5814 = True


} else {
ifres5814 = False


}

ifres5813 = ifres5814


} else {
ifres5813 = False


}

if True == ifres5813 {
tmp5771 := PrimTail(V3124)

tmp5772 := PrimTail(tmp5771)

__e.Return(PrimTail(tmp5772))
return


} else {
tmp5811 := PrimIsPair(V3124)

var ifres5798 Obj

if True == tmp5811 {
tmp5809 := PrimHead(V3124)

tmp5810 := PrimEqual(sympackage, tmp5809)

var ifres5800 Obj

if True == tmp5810 {
tmp5807 := PrimTail(V3124)

tmp5808 := PrimIsPair(tmp5807)

var ifres5802 Obj

if True == tmp5808 {
tmp5804 := PrimTail(V3124)

tmp5805 := PrimTail(tmp5804)

tmp5806 := PrimIsPair(tmp5805)

var ifres5803 Obj

if True == tmp5806 {
ifres5803 = True


} else {
ifres5803 = False


}

ifres5802 = ifres5803


} else {
ifres5802 = False


}

var ifres5801 Obj

if True == ifres5802 {
ifres5801 = True


} else {
ifres5801 = False


}

ifres5800 = ifres5801


} else {
ifres5800 = False


}

var ifres5799 Obj

if True == ifres5800 {
ifres5799 = True


} else {
ifres5799 = False


}

ifres5798 = ifres5799


} else {
ifres5798 = False


}

if True == ifres5798 {
tmp5773 := MakeNative(func(__e *ControlFlow) {
W3125 := __e.Get(1)
_ = W3125
tmp5774 := MakeNative(func(__e *ControlFlow) {
W3126 := __e.Get(1)
_ = W3126
tmp5775 := MakeNative(func(__e *ControlFlow) {
W3127 := __e.Get(1)
_ = W3127
tmp5776 := MakeNative(func(__e *ControlFlow) {
W3128 := __e.Get(1)
_ = W3128
__e.Return(W3126)
return
}, 1)

tmp5777 := PrimTail(V3124)

tmp5778 := PrimHead(tmp5777)

tmp5779 := PrimTail(V3124)

tmp5780 := PrimTail(tmp5779)

tmp5781 := PrimTail(tmp5780)

tmp5782 := Call(__e, PrimFunc(symshen_4record_1internal), tmp5778, W3125, tmp5781)


__e.TailApply(tmp5776, tmp5782)
return


}, 1)

tmp5783 := PrimTail(V3124)

tmp5784 := PrimHead(tmp5783)

tmp5785 := Call(__e, PrimFunc(symshen_4record_1external), tmp5784, W3125)


__e.TailApply(tmp5775, tmp5785)
return


}, 1)

tmp5786 := PrimTail(V3124)

tmp5787 := PrimHead(tmp5786)

tmp5788 := PrimStr(tmp5787)

tmp5789 := PrimTail(V3124)

tmp5790 := PrimTail(tmp5789)

tmp5791 := PrimTail(tmp5790)

tmp5792 := Call(__e, PrimFunc(symshen_4package_1symbols), tmp5788, W3125, tmp5791)


__e.TailApply(tmp5774, tmp5792)
return


}, 1)

tmp5793 := PrimTail(V3124)

tmp5794 := PrimTail(tmp5793)

tmp5795 := PrimHead(tmp5794)

tmp5796 := Call(__e, PrimFunc(symeval), tmp5795)


__e.TailApply(tmp5773, tmp5796)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4unpackage)
return
}


}


}, 1)

tmp5832 := Call(__e, ns2_1set, symshen_4unpackage, tmp5770)


_ = tmp5832

tmp5833 := MakeNative(func(__e *ControlFlow) {
V3129 := __e.Get(1)
_ = V3129
V3130 := __e.Get(2)
_ = V3130
V3131 := __e.Get(3)
_ = V3131
tmp5834 := MakeNative(func(__e *ControlFlow) {
W3132 := __e.Get(1)
_ = W3132
tmp5835 := MakeNative(func(__e *ControlFlow) {
W3134 := __e.Get(1)
_ = W3134
tmp5836 := Call(__e, PrimFunc(symunion), W3134, W3132)


tmp5837 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V3129, symshen_4internal_1symbols, tmp5836, tmp5837)
return


}, 1)

tmp5838 := PrimStr(V3129)

tmp5839 := Call(__e, PrimFunc(symshen_4internal_1symbols), tmp5838, V3130, V3131)


__e.TailApply(tmp5835, tmp5839)
return


}, 1)

tmp5840 := MakeNative(func(__e *ControlFlow) {
tmp5841 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V3129, symshen_4internal_1symbols, tmp5841)
return


}, 0)

tmp5842 := MakeNative(func(__e *ControlFlow) {
Z3133 := __e.Get(1)
_ = Z3133
__e.Return(Nil)
return
}, 1)

tmp5843 := Call(__e, try_1catch, tmp5840, tmp5842)


__e.TailApply(tmp5834, tmp5843)
return


}, 3)

tmp5844 := Call(__e, ns2_1set, symshen_4record_1internal, tmp5833)


_ = tmp5844

tmp5845 := MakeNative(func(__e *ControlFlow) {
V3141 := __e.Get(1)
_ = V3141
V3142 := __e.Get(2)
_ = V3142
V3143 := __e.Get(3)
_ = V3143
tmp5854 := PrimIsPair(V3143)

if True == tmp5854 {
tmp5846 := PrimHead(V3143)

tmp5847 := Call(__e, PrimFunc(symshen_4internal_1symbols), V3141, V3142, tmp5846)


tmp5848 := PrimTail(V3143)

tmp5849 := Call(__e, PrimFunc(symshen_4internal_1symbols), V3141, V3142, tmp5848)


__e.TailApply(PrimFunc(symunion), tmp5847, tmp5849)
return


} else {
tmp5852 := Call(__e, PrimFunc(symshen_4internal_2), V3143, V3141, V3142)


if True == tmp5852 {
tmp5850 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V3141, V3143)


__e.Return(PrimCons(tmp5850, Nil))
return


} else {
__e.Return(Nil)
return
}


}


}, 3)

tmp5855 := Call(__e, ns2_1set, symshen_4internal_1symbols, tmp5845)


_ = tmp5855

tmp5856 := MakeNative(func(__e *ControlFlow) {
V3144 := __e.Get(1)
_ = V3144
V3145 := __e.Get(2)
_ = V3145
tmp5857 := MakeNative(func(__e *ControlFlow) {
W3146 := __e.Get(1)
_ = W3146
tmp5858 := Call(__e, PrimFunc(symunion), V3145, W3146)


tmp5859 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V3144, symshen_4external_1symbols, tmp5858, tmp5859)
return


}, 1)

tmp5860 := MakeNative(func(__e *ControlFlow) {
tmp5861 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V3144, symshen_4external_1symbols, tmp5861)
return


}, 0)

tmp5862 := MakeNative(func(__e *ControlFlow) {
Z3147 := __e.Get(1)
_ = Z3147
__e.Return(Nil)
return
}, 1)

tmp5863 := Call(__e, try_1catch, tmp5860, tmp5862)


__e.TailApply(tmp5857, tmp5863)
return


}, 2)

tmp5864 := Call(__e, ns2_1set, symshen_4record_1external, tmp5856)


_ = tmp5864

tmp5865 := MakeNative(func(__e *ControlFlow) {
V3152 := __e.Get(1)
_ = V3152
V3153 := __e.Get(2)
_ = V3153
V3154 := __e.Get(3)
_ = V3154
tmp5870 := PrimIsPair(V3154)

if True == tmp5870 {
tmp5866 := MakeNative(func(__e *ControlFlow) {
Z3155 := __e.Get(1)
_ = Z3155
__e.TailApply(PrimFunc(symshen_4package_1symbols), V3152, V3153, Z3155)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp5866, V3154)
return


} else {
tmp5868 := Call(__e, PrimFunc(symshen_4internal_2), V3154, V3152, V3153)


if True == tmp5868 {
__e.TailApply(PrimFunc(symshen_4intern_1in_1package), V3152, V3154)
return
} else {
__e.Return(V3154)
return
}


}


}, 3)

tmp5871 := Call(__e, ns2_1set, symshen_4package_1symbols, tmp5865)


_ = tmp5871

tmp5872 := MakeNative(func(__e *ControlFlow) {
V3156 := __e.Get(1)
_ = V3156
V3157 := __e.Get(2)
_ = V3157
tmp5873 := PrimStr(V3157)

tmp5874 := Call(__e, PrimFunc(sym_8s), MakeString("."), tmp5873)


tmp5875 := Call(__e, PrimFunc(sym_8s), V3156, tmp5874)


__e.Return(PrimIntern(tmp5875))
return


}, 2)

tmp5876 := Call(__e, ns2_1set, symshen_4intern_1in_1package, tmp5872)


_ = tmp5876

tmp5877 := MakeNative(func(__e *ControlFlow) {
V3158 := __e.Get(1)
_ = V3158
V3159 := __e.Get(2)
_ = V3159
V3160 := __e.Get(3)
_ = V3160
tmp5907 := Call(__e, PrimFunc(symelement_2), V3158, V3160)


tmp5908 := PrimNot(tmp5907)

if True == tmp5908 {
tmp5904 := Call(__e, PrimFunc(symshen_4sng_2), V3158)


tmp5905 := PrimNot(tmp5904)

var ifres5879 Obj

if True == tmp5905 {
tmp5902 := Call(__e, PrimFunc(symshen_4dbl_2), V3158)


tmp5903 := PrimNot(tmp5902)

var ifres5881 Obj

if True == tmp5903 {
tmp5901 := PrimIsSymbol(V3158)

var ifres5883 Obj

if True == tmp5901 {
tmp5899 := Call(__e, PrimFunc(symshen_4sysfunc_2), V3158)


tmp5900 := PrimNot(tmp5899)

var ifres5885 Obj

if True == tmp5900 {
tmp5897 := PrimIsVariable(V3158)

tmp5898 := PrimNot(tmp5897)

var ifres5887 Obj

if True == tmp5898 {
tmp5894 := PrimStr(V3158)

tmp5895 := Call(__e, PrimFunc(symshen_4internal_1to_1shen_2), tmp5894)


tmp5896 := PrimNot(tmp5895)

var ifres5889 Obj

if True == tmp5896 {
tmp5891 := PrimStr(V3158)

tmp5892 := Call(__e, PrimFunc(symshen_4internal_1to_1P_2), V3159, tmp5891)


tmp5893 := PrimNot(tmp5892)

var ifres5890 Obj

if True == tmp5893 {
ifres5890 = True


} else {
ifres5890 = False


}

ifres5889 = ifres5890


} else {
ifres5889 = False


}

var ifres5888 Obj

if True == ifres5889 {
ifres5888 = True


} else {
ifres5888 = False


}

ifres5887 = ifres5888


} else {
ifres5887 = False


}

var ifres5886 Obj

if True == ifres5887 {
ifres5886 = True


} else {
ifres5886 = False


}

ifres5885 = ifres5886


} else {
ifres5885 = False


}

var ifres5884 Obj

if True == ifres5885 {
ifres5884 = True


} else {
ifres5884 = False


}

ifres5883 = ifres5884


} else {
ifres5883 = False


}

var ifres5882 Obj

if True == ifres5883 {
ifres5882 = True


} else {
ifres5882 = False


}

ifres5881 = ifres5882


} else {
ifres5881 = False


}

var ifres5880 Obj

if True == ifres5881 {
ifres5880 = True


} else {
ifres5880 = False


}

ifres5879 = ifres5880


} else {
ifres5879 = False


}

if True == ifres5879 {
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

tmp5909 := Call(__e, ns2_1set, symshen_4internal_2, tmp5877)


_ = tmp5909

tmp5910 := MakeNative(func(__e *ControlFlow) {
V3165 := __e.Get(1)
_ = V3165
tmp5964 := Call(__e, PrimFunc(symshen_4_7string_2), V3165)


var ifres5912 Obj

if True == tmp5964 {
tmp5962 := Call(__e, PrimFunc(symhdstr), V3165)


tmp5963 := PrimEqual(MakeString("s"), tmp5962)

var ifres5914 Obj

if True == tmp5963 {
tmp5960 := PrimTailString(V3165)

tmp5961 := Call(__e, PrimFunc(symshen_4_7string_2), tmp5960)


var ifres5916 Obj

if True == tmp5961 {
tmp5957 := PrimTailString(V3165)

tmp5958 := Call(__e, PrimFunc(symhdstr), tmp5957)


tmp5959 := PrimEqual(MakeString("h"), tmp5958)

var ifres5918 Obj

if True == tmp5959 {
tmp5954 := PrimTailString(V3165)

tmp5955 := PrimTailString(tmp5954)

tmp5956 := Call(__e, PrimFunc(symshen_4_7string_2), tmp5955)


var ifres5920 Obj

if True == tmp5956 {
tmp5950 := PrimTailString(V3165)

tmp5951 := PrimTailString(tmp5950)

tmp5952 := Call(__e, PrimFunc(symhdstr), tmp5951)


tmp5953 := PrimEqual(MakeString("e"), tmp5952)

var ifres5922 Obj

if True == tmp5953 {
tmp5946 := PrimTailString(V3165)

tmp5947 := PrimTailString(tmp5946)

tmp5948 := PrimTailString(tmp5947)

tmp5949 := Call(__e, PrimFunc(symshen_4_7string_2), tmp5948)


var ifres5924 Obj

if True == tmp5949 {
tmp5941 := PrimTailString(V3165)

tmp5942 := PrimTailString(tmp5941)

tmp5943 := PrimTailString(tmp5942)

tmp5944 := Call(__e, PrimFunc(symhdstr), tmp5943)


tmp5945 := PrimEqual(MakeString("n"), tmp5944)

var ifres5926 Obj

if True == tmp5945 {
tmp5936 := PrimTailString(V3165)

tmp5937 := PrimTailString(tmp5936)

tmp5938 := PrimTailString(tmp5937)

tmp5939 := PrimTailString(tmp5938)

tmp5940 := Call(__e, PrimFunc(symshen_4_7string_2), tmp5939)


var ifres5928 Obj

if True == tmp5940 {
tmp5930 := PrimTailString(V3165)

tmp5931 := PrimTailString(tmp5930)

tmp5932 := PrimTailString(tmp5931)

tmp5933 := PrimTailString(tmp5932)

tmp5934 := Call(__e, PrimFunc(symhdstr), tmp5933)


tmp5935 := PrimEqual(MakeString("."), tmp5934)

var ifres5929 Obj

if True == tmp5935 {
ifres5929 = True


} else {
ifres5929 = False


}

ifres5928 = ifres5929


} else {
ifres5928 = False


}

var ifres5927 Obj

if True == ifres5928 {
ifres5927 = True


} else {
ifres5927 = False


}

ifres5926 = ifres5927


} else {
ifres5926 = False


}

var ifres5925 Obj

if True == ifres5926 {
ifres5925 = True


} else {
ifres5925 = False


}

ifres5924 = ifres5925


} else {
ifres5924 = False


}

var ifres5923 Obj

if True == ifres5924 {
ifres5923 = True


} else {
ifres5923 = False


}

ifres5922 = ifres5923


} else {
ifres5922 = False


}

var ifres5921 Obj

if True == ifres5922 {
ifres5921 = True


} else {
ifres5921 = False


}

ifres5920 = ifres5921


} else {
ifres5920 = False


}

var ifres5919 Obj

if True == ifres5920 {
ifres5919 = True


} else {
ifres5919 = False


}

ifres5918 = ifres5919


} else {
ifres5918 = False


}

var ifres5917 Obj

if True == ifres5918 {
ifres5917 = True


} else {
ifres5917 = False


}

ifres5916 = ifres5917


} else {
ifres5916 = False


}

var ifres5915 Obj

if True == ifres5916 {
ifres5915 = True


} else {
ifres5915 = False


}

ifres5914 = ifres5915


} else {
ifres5914 = False


}

var ifres5913 Obj

if True == ifres5914 {
ifres5913 = True


} else {
ifres5913 = False


}

ifres5912 = ifres5913


} else {
ifres5912 = False


}

if True == ifres5912 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp5965 := Call(__e, ns2_1set, symshen_4internal_1to_1shen_2, tmp5910)


_ = tmp5965

tmp5966 := MakeNative(func(__e *ControlFlow) {
V3166 := __e.Get(1)
_ = V3166
tmp5967 := PrimValue(sym_dproperty_1vector_d)

tmp5968 := Call(__e, PrimFunc(symget), symshen, symshen_4external_1symbols, tmp5967)


__e.TailApply(PrimFunc(symelement_2), V3166, tmp5968)
return


}, 1)

tmp5969 := Call(__e, ns2_1set, symshen_4sysfunc_2, tmp5966)


_ = tmp5969

tmp5970 := MakeNative(func(__e *ControlFlow) {
V3174 := __e.Get(1)
_ = V3174
V3175 := __e.Get(2)
_ = V3175
tmp5991 := PrimEqual(MakeString(""), V3174)

var ifres5984 Obj

if True == tmp5991 {
tmp5990 := Call(__e, PrimFunc(symshen_4_7string_2), V3175)


var ifres5986 Obj

if True == tmp5990 {
tmp5988 := Call(__e, PrimFunc(symhdstr), V3175)


tmp5989 := PrimEqual(MakeString("."), tmp5988)

var ifres5987 Obj

if True == tmp5989 {
ifres5987 = True


} else {
ifres5987 = False


}

ifres5986 = ifres5987


} else {
ifres5986 = False


}

var ifres5985 Obj

if True == ifres5986 {
ifres5985 = True


} else {
ifres5985 = False


}

ifres5984 = ifres5985


} else {
ifres5984 = False


}

if True == ifres5984 {
__e.Return(True)
return
} else {
tmp5982 := Call(__e, PrimFunc(symshen_4_7string_2), V3174)


var ifres5974 Obj

if True == tmp5982 {
tmp5981 := Call(__e, PrimFunc(symshen_4_7string_2), V3175)


var ifres5976 Obj

if True == tmp5981 {
tmp5978 := Call(__e, PrimFunc(symhdstr), V3174)


tmp5979 := Call(__e, PrimFunc(symhdstr), V3175)


tmp5980 := PrimEqual(tmp5978, tmp5979)

var ifres5977 Obj

if True == tmp5980 {
ifres5977 = True


} else {
ifres5977 = False


}

ifres5976 = ifres5977


} else {
ifres5976 = False


}

var ifres5975 Obj

if True == ifres5976 {
ifres5975 = True


} else {
ifres5975 = False


}

ifres5974 = ifres5975


} else {
ifres5974 = False


}

if True == ifres5974 {
tmp5971 := PrimTailString(V3174)

tmp5972 := PrimTailString(V3175)

__e.TailApply(PrimFunc(symshen_4internal_1to_1P_2), tmp5971, tmp5972)
return


} else {
__e.Return(False)
return
}


}


}, 2)

tmp5992 := Call(__e, ns2_1set, symshen_4internal_1to_1P_2, tmp5970)


_ = tmp5992

tmp5993 := MakeNative(func(__e *ControlFlow) {
V3178 := __e.Get(1)
_ = V3178
V3179 := __e.Get(2)
_ = V3179
tmp6014 := Call(__e, PrimFunc(symelement_2), V3178, V3179)


if True == tmp6014 {
__e.Return(V3178)
return
} else {
tmp6012 := PrimIsPair(V3178)

var ifres6008 Obj

if True == tmp6012 {
tmp6010 := PrimHead(V3178)

tmp6011 := PrimEqual(symcond, tmp6010)

var ifres6009 Obj

if True == tmp6011 {
ifres6009 = True


} else {
ifres6009 = False


}

ifres6008 = ifres6009


} else {
ifres6008 = False


}

if True == ifres6008 {
tmp5994 := PrimTail(V3178)

tmp5995 := Call(__e, PrimFunc(symshen_4process_1cond_1clauses), tmp5994, V3179)


__e.Return(PrimCons(symcond, tmp5995))
return


} else {
tmp6006 := PrimIsPair(V3178)

var ifres6002 Obj

if True == tmp6006 {
tmp6004 := PrimHead(V3178)

tmp6005 := Call(__e, PrimFunc(symshen_4non_1application_2), tmp6004)


var ifres6003 Obj

if True == tmp6005 {
ifres6003 = True


} else {
ifres6003 = False


}

ifres6002 = ifres6003


} else {
ifres6002 = False


}

if True == ifres6002 {
tmp5996 := PrimHead(V3178)

__e.TailApply(PrimFunc(symshen_4special_1case), tmp5996, V3178, V3179)
return


} else {
tmp6000 := PrimIsPair(V3178)

if True == tmp6000 {
tmp5997 := MakeNative(func(__e *ControlFlow) {
Z3180 := __e.Get(1)
_ = Z3180
__e.TailApply(PrimFunc(symshen_4process_1applications), Z3180, V3179)
return
}, 1)

tmp5998 := Call(__e, PrimFunc(symmap), tmp5997, V3178)


__e.TailApply(PrimFunc(symshen_4process_1application), tmp5998, V3179)
return


} else {
__e.Return(V3178)
return
}


}


}


}


}, 2)

tmp6015 := Call(__e, ns2_1set, symshen_4process_1applications, tmp5993)


_ = tmp6015

tmp6016 := MakeNative(func(__e *ControlFlow) {
V3183 := __e.Get(1)
_ = V3183
tmp6026 := PrimEqual(symdefine, V3183)

if True == tmp6026 {
__e.Return(True)
return
} else {
tmp6024 := PrimEqual(symdefun, V3183)

if True == tmp6024 {
__e.Return(True)
return
} else {
tmp6022 := PrimEqual(symsynonyms, V3183)

if True == tmp6022 {
__e.Return(True)
return
} else {
tmp6020 := Call(__e, PrimFunc(symshen_4special_2), V3183)


if True == tmp6020 {
__e.Return(True)
return
} else {
tmp6018 := Call(__e, PrimFunc(symshen_4extraspecial_2), V3183)


if True == tmp6018 {
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

tmp6027 := Call(__e, ns2_1set, symshen_4non_1application_2, tmp6016)


_ = tmp6027

tmp6028 := MakeNative(func(__e *ControlFlow) {
V3188 := __e.Get(1)
_ = V3188
V3189 := __e.Get(2)
_ = V3189
V3190 := __e.Get(3)
_ = V3190
tmp6270 := PrimEqual(symlambda, V3188)

var ifres6248 Obj

if True == tmp6270 {
tmp6269 := PrimIsPair(V3189)

var ifres6250 Obj

if True == tmp6269 {
tmp6267 := PrimHead(V3189)

tmp6268 := PrimEqual(symlambda, tmp6267)

var ifres6252 Obj

if True == tmp6268 {
tmp6265 := PrimTail(V3189)

tmp6266 := PrimIsPair(tmp6265)

var ifres6254 Obj

if True == tmp6266 {
tmp6262 := PrimTail(V3189)

tmp6263 := PrimTail(tmp6262)

tmp6264 := PrimIsPair(tmp6263)

var ifres6256 Obj

if True == tmp6264 {
tmp6258 := PrimTail(V3189)

tmp6259 := PrimTail(tmp6258)

tmp6260 := PrimTail(tmp6259)

tmp6261 := PrimEqual(Nil, tmp6260)

var ifres6257 Obj

if True == tmp6261 {
ifres6257 = True


} else {
ifres6257 = False


}

ifres6256 = ifres6257


} else {
ifres6256 = False


}

var ifres6255 Obj

if True == ifres6256 {
ifres6255 = True


} else {
ifres6255 = False


}

ifres6254 = ifres6255


} else {
ifres6254 = False


}

var ifres6253 Obj

if True == ifres6254 {
ifres6253 = True


} else {
ifres6253 = False


}

ifres6252 = ifres6253


} else {
ifres6252 = False


}

var ifres6251 Obj

if True == ifres6252 {
ifres6251 = True


} else {
ifres6251 = False


}

ifres6250 = ifres6251


} else {
ifres6250 = False


}

var ifres6249 Obj

if True == ifres6250 {
ifres6249 = True


} else {
ifres6249 = False


}

ifres6248 = ifres6249


} else {
ifres6248 = False


}

if True == ifres6248 {
tmp6029 := PrimTail(V3189)

tmp6030 := PrimHead(tmp6029)

tmp6031 := PrimTail(V3189)

tmp6032 := PrimTail(tmp6031)

tmp6033 := PrimHead(tmp6032)

tmp6034 := Call(__e, PrimFunc(symshen_4process_1applications), tmp6033, V3190)


tmp6035 := PrimCons(tmp6034, Nil)

tmp6036 := PrimCons(tmp6030, tmp6035)

__e.Return(PrimCons(symlambda, tmp6036))
return


} else {
tmp6246 := PrimEqual(symlet, V3188)

var ifres6217 Obj

if True == tmp6246 {
tmp6245 := PrimIsPair(V3189)

var ifres6219 Obj

if True == tmp6245 {
tmp6243 := PrimHead(V3189)

tmp6244 := PrimEqual(symlet, tmp6243)

var ifres6221 Obj

if True == tmp6244 {
tmp6241 := PrimTail(V3189)

tmp6242 := PrimIsPair(tmp6241)

var ifres6223 Obj

if True == tmp6242 {
tmp6238 := PrimTail(V3189)

tmp6239 := PrimTail(tmp6238)

tmp6240 := PrimIsPair(tmp6239)

var ifres6225 Obj

if True == tmp6240 {
tmp6234 := PrimTail(V3189)

tmp6235 := PrimTail(tmp6234)

tmp6236 := PrimTail(tmp6235)

tmp6237 := PrimIsPair(tmp6236)

var ifres6227 Obj

if True == tmp6237 {
tmp6229 := PrimTail(V3189)

tmp6230 := PrimTail(tmp6229)

tmp6231 := PrimTail(tmp6230)

tmp6232 := PrimTail(tmp6231)

tmp6233 := PrimEqual(Nil, tmp6232)

var ifres6228 Obj

if True == tmp6233 {
ifres6228 = True


} else {
ifres6228 = False


}

ifres6227 = ifres6228


} else {
ifres6227 = False


}

var ifres6226 Obj

if True == ifres6227 {
ifres6226 = True


} else {
ifres6226 = False


}

ifres6225 = ifres6226


} else {
ifres6225 = False


}

var ifres6224 Obj

if True == ifres6225 {
ifres6224 = True


} else {
ifres6224 = False


}

ifres6223 = ifres6224


} else {
ifres6223 = False


}

var ifres6222 Obj

if True == ifres6223 {
ifres6222 = True


} else {
ifres6222 = False


}

ifres6221 = ifres6222


} else {
ifres6221 = False


}

var ifres6220 Obj

if True == ifres6221 {
ifres6220 = True


} else {
ifres6220 = False


}

ifres6219 = ifres6220


} else {
ifres6219 = False


}

var ifres6218 Obj

if True == ifres6219 {
ifres6218 = True


} else {
ifres6218 = False


}

ifres6217 = ifres6218


} else {
ifres6217 = False


}

if True == ifres6217 {
tmp6037 := PrimTail(V3189)

tmp6038 := PrimHead(tmp6037)

tmp6039 := PrimTail(V3189)

tmp6040 := PrimTail(tmp6039)

tmp6041 := PrimHead(tmp6040)

tmp6042 := Call(__e, PrimFunc(symshen_4process_1applications), tmp6041, V3190)


tmp6043 := PrimTail(V3189)

tmp6044 := PrimTail(tmp6043)

tmp6045 := PrimTail(tmp6044)

tmp6046 := PrimHead(tmp6045)

tmp6047 := Call(__e, PrimFunc(symshen_4process_1applications), tmp6046, V3190)


tmp6048 := PrimCons(tmp6047, Nil)

tmp6049 := PrimCons(tmp6042, tmp6048)

tmp6050 := PrimCons(tmp6038, tmp6049)

__e.Return(PrimCons(symlet, tmp6050))
return


} else {
tmp6215 := PrimEqual(symdefun, V3188)

var ifres6186 Obj

if True == tmp6215 {
tmp6214 := PrimIsPair(V3189)

var ifres6188 Obj

if True == tmp6214 {
tmp6212 := PrimHead(V3189)

tmp6213 := PrimEqual(symdefun, tmp6212)

var ifres6190 Obj

if True == tmp6213 {
tmp6210 := PrimTail(V3189)

tmp6211 := PrimIsPair(tmp6210)

var ifres6192 Obj

if True == tmp6211 {
tmp6207 := PrimTail(V3189)

tmp6208 := PrimTail(tmp6207)

tmp6209 := PrimIsPair(tmp6208)

var ifres6194 Obj

if True == tmp6209 {
tmp6203 := PrimTail(V3189)

tmp6204 := PrimTail(tmp6203)

tmp6205 := PrimTail(tmp6204)

tmp6206 := PrimIsPair(tmp6205)

var ifres6196 Obj

if True == tmp6206 {
tmp6198 := PrimTail(V3189)

tmp6199 := PrimTail(tmp6198)

tmp6200 := PrimTail(tmp6199)

tmp6201 := PrimTail(tmp6200)

tmp6202 := PrimEqual(Nil, tmp6201)

var ifres6197 Obj

if True == tmp6202 {
ifres6197 = True


} else {
ifres6197 = False


}

ifres6196 = ifres6197


} else {
ifres6196 = False


}

var ifres6195 Obj

if True == ifres6196 {
ifres6195 = True


} else {
ifres6195 = False


}

ifres6194 = ifres6195


} else {
ifres6194 = False


}

var ifres6193 Obj

if True == ifres6194 {
ifres6193 = True


} else {
ifres6193 = False


}

ifres6192 = ifres6193


} else {
ifres6192 = False


}

var ifres6191 Obj

if True == ifres6192 {
ifres6191 = True


} else {
ifres6191 = False


}

ifres6190 = ifres6191


} else {
ifres6190 = False


}

var ifres6189 Obj

if True == ifres6190 {
ifres6189 = True


} else {
ifres6189 = False


}

ifres6188 = ifres6189


} else {
ifres6188 = False


}

var ifres6187 Obj

if True == ifres6188 {
ifres6187 = True


} else {
ifres6187 = False


}

ifres6186 = ifres6187


} else {
ifres6186 = False


}

if True == ifres6186 {
__e.Return(V3189)
return
} else {
tmp6184 := PrimEqual(symdefine, V3188)

var ifres6162 Obj

if True == tmp6184 {
tmp6183 := PrimIsPair(V3189)

var ifres6164 Obj

if True == tmp6183 {
tmp6181 := PrimHead(V3189)

tmp6182 := PrimEqual(symdefine, tmp6181)

var ifres6166 Obj

if True == tmp6182 {
tmp6179 := PrimTail(V3189)

tmp6180 := PrimIsPair(tmp6179)

var ifres6168 Obj

if True == tmp6180 {
tmp6176 := PrimTail(V3189)

tmp6177 := PrimTail(tmp6176)

tmp6178 := PrimIsPair(tmp6177)

var ifres6170 Obj

if True == tmp6178 {
tmp6172 := PrimTail(V3189)

tmp6173 := PrimTail(tmp6172)

tmp6174 := PrimHead(tmp6173)

tmp6175 := PrimEqual(sym_i, tmp6174)

var ifres6171 Obj

if True == tmp6175 {
ifres6171 = True


} else {
ifres6171 = False


}

ifres6170 = ifres6171


} else {
ifres6170 = False


}

var ifres6169 Obj

if True == ifres6170 {
ifres6169 = True


} else {
ifres6169 = False


}

ifres6168 = ifres6169


} else {
ifres6168 = False


}

var ifres6167 Obj

if True == ifres6168 {
ifres6167 = True


} else {
ifres6167 = False


}

ifres6166 = ifres6167


} else {
ifres6166 = False


}

var ifres6165 Obj

if True == ifres6166 {
ifres6165 = True


} else {
ifres6165 = False


}

ifres6164 = ifres6165


} else {
ifres6164 = False


}

var ifres6163 Obj

if True == ifres6164 {
ifres6163 = True


} else {
ifres6163 = False


}

ifres6162 = ifres6163


} else {
ifres6162 = False


}

if True == ifres6162 {
tmp6051 := PrimTail(V3189)

tmp6052 := PrimHead(tmp6051)

tmp6053 := PrimTail(V3189)

tmp6054 := PrimHead(tmp6053)

tmp6055 := PrimTail(V3189)

tmp6056 := PrimTail(tmp6055)

tmp6057 := PrimTail(tmp6056)

tmp6058 := Call(__e, PrimFunc(symshen_4process_1after_1type), tmp6054, tmp6057, V3190)


tmp6059 := PrimCons(sym_i, tmp6058)

tmp6060 := PrimCons(tmp6052, tmp6059)

__e.Return(PrimCons(symdefine, tmp6060))
return


} else {
tmp6160 := PrimEqual(symdefine, V3188)

var ifres6149 Obj

if True == tmp6160 {
tmp6159 := PrimIsPair(V3189)

var ifres6151 Obj

if True == tmp6159 {
tmp6157 := PrimHead(V3189)

tmp6158 := PrimEqual(symdefine, tmp6157)

var ifres6153 Obj

if True == tmp6158 {
tmp6155 := PrimTail(V3189)

tmp6156 := PrimIsPair(tmp6155)

var ifres6154 Obj

if True == tmp6156 {
ifres6154 = True


} else {
ifres6154 = False


}

ifres6153 = ifres6154


} else {
ifres6153 = False


}

var ifres6152 Obj

if True == ifres6153 {
ifres6152 = True


} else {
ifres6152 = False


}

ifres6151 = ifres6152


} else {
ifres6151 = False


}

var ifres6150 Obj

if True == ifres6151 {
ifres6150 = True


} else {
ifres6150 = False


}

ifres6149 = ifres6150


} else {
ifres6149 = False


}

if True == ifres6149 {
tmp6061 := PrimTail(V3189)

tmp6062 := PrimHead(tmp6061)

tmp6063 := MakeNative(func(__e *ControlFlow) {
Z3191 := __e.Get(1)
_ = Z3191
__e.TailApply(PrimFunc(symshen_4process_1applications), Z3191, V3190)
return
}, 1)

tmp6064 := PrimTail(V3189)

tmp6065 := PrimTail(tmp6064)

tmp6066 := Call(__e, PrimFunc(symmap), tmp6063, tmp6065)


tmp6067 := PrimCons(tmp6062, tmp6066)

__e.Return(PrimCons(symdefine, tmp6067))
return


} else {
tmp6147 := PrimEqual(symsynonyms, V3188)

if True == tmp6147 {
__e.Return(PrimCons(symsynonyms, V3189))
return
} else {
tmp6145 := PrimEqual(symtype, V3188)

var ifres6123 Obj

if True == tmp6145 {
tmp6144 := PrimIsPair(V3189)

var ifres6125 Obj

if True == tmp6144 {
tmp6142 := PrimHead(V3189)

tmp6143 := PrimEqual(symtype, tmp6142)

var ifres6127 Obj

if True == tmp6143 {
tmp6140 := PrimTail(V3189)

tmp6141 := PrimIsPair(tmp6140)

var ifres6129 Obj

if True == tmp6141 {
tmp6137 := PrimTail(V3189)

tmp6138 := PrimTail(tmp6137)

tmp6139 := PrimIsPair(tmp6138)

var ifres6131 Obj

if True == tmp6139 {
tmp6133 := PrimTail(V3189)

tmp6134 := PrimTail(tmp6133)

tmp6135 := PrimTail(tmp6134)

tmp6136 := PrimEqual(Nil, tmp6135)

var ifres6132 Obj

if True == tmp6136 {
ifres6132 = True


} else {
ifres6132 = False


}

ifres6131 = ifres6132


} else {
ifres6131 = False


}

var ifres6130 Obj

if True == ifres6131 {
ifres6130 = True


} else {
ifres6130 = False


}

ifres6129 = ifres6130


} else {
ifres6129 = False


}

var ifres6128 Obj

if True == ifres6129 {
ifres6128 = True


} else {
ifres6128 = False


}

ifres6127 = ifres6128


} else {
ifres6127 = False


}

var ifres6126 Obj

if True == ifres6127 {
ifres6126 = True


} else {
ifres6126 = False


}

ifres6125 = ifres6126


} else {
ifres6125 = False


}

var ifres6124 Obj

if True == ifres6125 {
ifres6124 = True


} else {
ifres6124 = False


}

ifres6123 = ifres6124


} else {
ifres6123 = False


}

if True == ifres6123 {
tmp6068 := PrimTail(V3189)

tmp6069 := PrimHead(tmp6068)

tmp6070 := Call(__e, PrimFunc(symshen_4process_1applications), tmp6069, V3190)


tmp6071 := PrimTail(V3189)

tmp6072 := PrimTail(tmp6071)

tmp6073 := PrimCons(tmp6070, tmp6072)

__e.Return(PrimCons(symtype, tmp6073))
return


} else {
tmp6121 := PrimEqual(syminput_7, V3188)

var ifres6099 Obj

if True == tmp6121 {
tmp6120 := PrimIsPair(V3189)

var ifres6101 Obj

if True == tmp6120 {
tmp6118 := PrimHead(V3189)

tmp6119 := PrimEqual(syminput_7, tmp6118)

var ifres6103 Obj

if True == tmp6119 {
tmp6116 := PrimTail(V3189)

tmp6117 := PrimIsPair(tmp6116)

var ifres6105 Obj

if True == tmp6117 {
tmp6113 := PrimTail(V3189)

tmp6114 := PrimTail(tmp6113)

tmp6115 := PrimIsPair(tmp6114)

var ifres6107 Obj

if True == tmp6115 {
tmp6109 := PrimTail(V3189)

tmp6110 := PrimTail(tmp6109)

tmp6111 := PrimTail(tmp6110)

tmp6112 := PrimEqual(Nil, tmp6111)

var ifres6108 Obj

if True == tmp6112 {
ifres6108 = True


} else {
ifres6108 = False


}

ifres6107 = ifres6108


} else {
ifres6107 = False


}

var ifres6106 Obj

if True == ifres6107 {
ifres6106 = True


} else {
ifres6106 = False


}

ifres6105 = ifres6106


} else {
ifres6105 = False


}

var ifres6104 Obj

if True == ifres6105 {
ifres6104 = True


} else {
ifres6104 = False


}

ifres6103 = ifres6104


} else {
ifres6103 = False


}

var ifres6102 Obj

if True == ifres6103 {
ifres6102 = True


} else {
ifres6102 = False


}

ifres6101 = ifres6102


} else {
ifres6101 = False


}

var ifres6100 Obj

if True == ifres6101 {
ifres6100 = True


} else {
ifres6100 = False


}

ifres6099 = ifres6100


} else {
ifres6099 = False


}

if True == ifres6099 {
tmp6074 := PrimTail(V3189)

tmp6075 := PrimHead(tmp6074)

tmp6076 := PrimTail(V3189)

tmp6077 := PrimTail(tmp6076)

tmp6078 := PrimHead(tmp6077)

tmp6079 := Call(__e, PrimFunc(symshen_4process_1applications), tmp6078, V3190)


tmp6080 := PrimCons(tmp6079, Nil)

tmp6081 := PrimCons(tmp6075, tmp6080)

__e.Return(PrimCons(syminput_7, tmp6081))
return


} else {
tmp6097 := PrimIsPair(V3189)

var ifres6093 Obj

if True == tmp6097 {
tmp6095 := PrimHead(V3189)

tmp6096 := Call(__e, PrimFunc(symshen_4special_2), tmp6095)


var ifres6094 Obj

if True == tmp6096 {
ifres6094 = True


} else {
ifres6094 = False


}

ifres6093 = ifres6094


} else {
ifres6093 = False


}

if True == ifres6093 {
tmp6082 := PrimHead(V3189)

tmp6083 := MakeNative(func(__e *ControlFlow) {
Z3192 := __e.Get(1)
_ = Z3192
__e.TailApply(PrimFunc(symshen_4process_1applications), Z3192, V3190)
return
}, 1)

tmp6084 := PrimTail(V3189)

tmp6085 := Call(__e, PrimFunc(symmap), tmp6083, tmp6084)


__e.Return(PrimCons(tmp6082, tmp6085))
return


} else {
tmp6091 := PrimIsPair(V3189)

var ifres6087 Obj

if True == tmp6091 {
tmp6089 := PrimHead(V3189)

tmp6090 := Call(__e, PrimFunc(symshen_4extraspecial_2), tmp6089)


var ifres6088 Obj

if True == tmp6090 {
ifres6088 = True


} else {
ifres6088 = False


}

ifres6087 = ifres6088


} else {
ifres6087 = False


}

if True == ifres6087 {
__e.Return(V3189)
return
} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4special_1case)
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

tmp6271 := Call(__e, ns2_1set, symshen_4special_1case, tmp6028)


_ = tmp6271

tmp6272 := MakeNative(func(__e *ControlFlow) {
V3195 := __e.Get(1)
_ = V3195
V3196 := __e.Get(2)
_ = V3196
tmp6302 := PrimEqual(Nil, V3195)

if True == tmp6302 {
__e.Return(Nil)
return
} else {
tmp6300 := PrimIsPair(V3195)

var ifres6285 Obj

if True == tmp6300 {
tmp6298 := PrimHead(V3195)

tmp6299 := PrimIsPair(tmp6298)

var ifres6287 Obj

if True == tmp6299 {
tmp6295 := PrimHead(V3195)

tmp6296 := PrimTail(tmp6295)

tmp6297 := PrimIsPair(tmp6296)

var ifres6289 Obj

if True == tmp6297 {
tmp6291 := PrimHead(V3195)

tmp6292 := PrimTail(tmp6291)

tmp6293 := PrimTail(tmp6292)

tmp6294 := PrimEqual(Nil, tmp6293)

var ifres6290 Obj

if True == tmp6294 {
ifres6290 = True


} else {
ifres6290 = False


}

ifres6289 = ifres6290


} else {
ifres6289 = False


}

var ifres6288 Obj

if True == ifres6289 {
ifres6288 = True


} else {
ifres6288 = False


}

ifres6287 = ifres6288


} else {
ifres6287 = False


}

var ifres6286 Obj

if True == ifres6287 {
ifres6286 = True


} else {
ifres6286 = False


}

ifres6285 = ifres6286


} else {
ifres6285 = False


}

if True == ifres6285 {
tmp6273 := PrimHead(V3195)

tmp6274 := PrimHead(tmp6273)

tmp6275 := Call(__e, PrimFunc(symshen_4process_1applications), tmp6274, V3196)


tmp6276 := PrimHead(V3195)

tmp6277 := PrimTail(tmp6276)

tmp6278 := PrimHead(tmp6277)

tmp6279 := Call(__e, PrimFunc(symshen_4process_1applications), tmp6278, V3196)


tmp6280 := PrimCons(tmp6279, Nil)

tmp6281 := PrimCons(tmp6275, tmp6280)

tmp6282 := PrimTail(V3195)

tmp6283 := Call(__e, PrimFunc(symshen_4process_1cond_1clauses), tmp6282, V3196)


__e.Return(PrimCons(tmp6281, tmp6283))
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4process_1cond_1clauses)
return
}


}


}, 2)

tmp6303 := Call(__e, ns2_1set, symshen_4process_1cond_1clauses, tmp6272)


_ = tmp6303

tmp6304 := MakeNative(func(__e *ControlFlow) {
V3199 := __e.Get(1)
_ = V3199
V3200 := __e.Get(2)
_ = V3200
V3201 := __e.Get(3)
_ = V3201
tmp6320 := PrimIsPair(V3200)

var ifres6316 Obj

if True == tmp6320 {
tmp6318 := PrimHead(V3200)

tmp6319 := PrimEqual(sym_j, tmp6318)

var ifres6317 Obj

if True == tmp6319 {
ifres6317 = True


} else {
ifres6317 = False


}

ifres6316 = ifres6317


} else {
ifres6316 = False


}

if True == ifres6316 {
tmp6305 := MakeNative(func(__e *ControlFlow) {
Z3202 := __e.Get(1)
_ = Z3202
__e.TailApply(PrimFunc(symshen_4process_1applications), Z3202, V3201)
return
}, 1)

tmp6306 := PrimTail(V3200)

tmp6307 := Call(__e, PrimFunc(symmap), tmp6305, tmp6306)


__e.Return(PrimCons(sym_j, tmp6307))
return


} else {
tmp6314 := PrimIsPair(V3200)

if True == tmp6314 {
tmp6308 := PrimHead(V3200)

tmp6309 := PrimTail(V3200)

tmp6310 := Call(__e, PrimFunc(symshen_4process_1after_1type), V3199, tmp6309, V3201)


__e.Return(PrimCons(tmp6308, tmp6310))
return


} else {
tmp6311 := Call(__e, PrimFunc(symshen_4app), V3199, MakeString("\n"), symshen_4a)


tmp6312 := PrimStringConcat(MakeString("missing } in "), tmp6311)

__e.Return(PrimSimpleError(tmp6312))
return


}


}


}, 3)

tmp6321 := Call(__e, ns2_1set, symshen_4process_1after_1type, tmp6304)


_ = tmp6321

tmp6322 := MakeNative(func(__e *ControlFlow) {
V3203 := __e.Get(1)
_ = V3203
V3204 := __e.Get(2)
_ = V3204
tmp6367 := PrimIsPair(V3203)

if True == tmp6367 {
tmp6323 := MakeNative(func(__e *ControlFlow) {
W3205 := __e.Get(1)
_ = W3205
tmp6324 := MakeNative(func(__e *ControlFlow) {
W3206 := __e.Get(1)
_ = W3206
tmp6361 := Call(__e, PrimFunc(symelement_2), V3203, V3204)


if True == tmp6361 {
__e.Return(V3203)
return
} else {
tmp6358 := PrimHead(V3203)

tmp6359 := Call(__e, PrimFunc(symshen_4shen_1call_2), tmp6358)


if True == tmp6359 {
__e.Return(V3203)
return
} else {
tmp6356 := Call(__e, PrimFunc(symshen_4foreign_2), V3203)


if True == tmp6356 {
__e.TailApply(PrimFunc(symshen_4unpack_1foreign), V3203)
return
} else {
tmp6354 := Call(__e, PrimFunc(symshen_4fn_1call_2), V3203)


if True == tmp6354 {
__e.TailApply(PrimFunc(symshen_4fn_1call), V3203)
return
} else {
tmp6352 := Call(__e, PrimFunc(symshen_4zero_1place_2), V3203)


if True == tmp6352 {
__e.Return(V3203)
return
} else {
tmp6349 := PrimHead(V3203)

tmp6350 := Call(__e, PrimFunc(symshen_4undefined_1f_2), tmp6349, W3205)


if True == tmp6350 {
tmp6325 := PrimHead(V3203)

tmp6326 := PrimCons(tmp6325, Nil)

tmp6327 := PrimCons(symfn, tmp6326)

tmp6328 := PrimTail(V3203)

tmp6329 := PrimCons(tmp6327, tmp6328)

__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp6329)
return


} else {
tmp6346 := PrimHead(V3203)

tmp6347 := PrimIsVariable(tmp6346)

if True == tmp6347 {
__e.TailApply(PrimFunc(symshen_4simple_1curry), V3203)
return
} else {
tmp6343 := PrimHead(V3203)

tmp6344 := Call(__e, PrimFunc(symshen_4application_2), tmp6343)


if True == tmp6344 {
__e.TailApply(PrimFunc(symshen_4simple_1curry), V3203)
return
} else {
tmp6340 := PrimHead(V3203)

tmp6341 := Call(__e, PrimFunc(symshen_4partial_1application_d_2), tmp6340, W3205, W3206)


if True == tmp6341 {
tmp6330 := PrimNumberSubtract(W3205, W3206)

__e.TailApply(PrimFunc(symshen_4lambda_1function), V3203, tmp6330)
return


} else {
tmp6337 := PrimHead(V3203)

tmp6338 := Call(__e, PrimFunc(symshen_4overapplication_2), tmp6337, W3205, W3206)


if True == tmp6338 {
tmp6331 := PrimHead(V3203)

tmp6332 := PrimCons(tmp6331, Nil)

tmp6333 := PrimCons(symfn, tmp6332)

tmp6334 := PrimTail(V3203)

tmp6335 := PrimCons(tmp6333, tmp6334)

__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp6335)
return


} else {
__e.Return(V3203)
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

tmp6362 := PrimTail(V3203)

tmp6363 := Call(__e, PrimFunc(symlength), tmp6362)


__e.TailApply(tmp6324, tmp6363)
return


}, 1)

tmp6364 := PrimHead(V3203)

tmp6365 := Call(__e, PrimFunc(symarity), tmp6364)


__e.TailApply(tmp6323, tmp6365)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4process_1application)
return
}


}, 2)

tmp6368 := Call(__e, ns2_1set, symshen_4process_1application, tmp6322)


_ = tmp6368

tmp6369 := MakeNative(func(__e *ControlFlow) {
V3207 := __e.Get(1)
_ = V3207
tmp6395 := PrimIsPair(V3207)

var ifres6375 Obj

if True == tmp6395 {
tmp6393 := PrimHead(V3207)

tmp6394 := PrimIsPair(tmp6393)

var ifres6377 Obj

if True == tmp6394 {
tmp6390 := PrimHead(V3207)

tmp6391 := PrimHead(tmp6390)

tmp6392 := PrimEqual(symforeign, tmp6391)

var ifres6379 Obj

if True == tmp6392 {
tmp6387 := PrimHead(V3207)

tmp6388 := PrimTail(tmp6387)

tmp6389 := PrimIsPair(tmp6388)

var ifres6381 Obj

if True == tmp6389 {
tmp6383 := PrimHead(V3207)

tmp6384 := PrimTail(tmp6383)

tmp6385 := PrimTail(tmp6384)

tmp6386 := PrimEqual(Nil, tmp6385)

var ifres6382 Obj

if True == tmp6386 {
ifres6382 = True


} else {
ifres6382 = False


}

ifres6381 = ifres6382


} else {
ifres6381 = False


}

var ifres6380 Obj

if True == ifres6381 {
ifres6380 = True


} else {
ifres6380 = False


}

ifres6379 = ifres6380


} else {
ifres6379 = False


}

var ifres6378 Obj

if True == ifres6379 {
ifres6378 = True


} else {
ifres6378 = False


}

ifres6377 = ifres6378


} else {
ifres6377 = False


}

var ifres6376 Obj

if True == ifres6377 {
ifres6376 = True


} else {
ifres6376 = False


}

ifres6375 = ifres6376


} else {
ifres6375 = False


}

if True == ifres6375 {
tmp6370 := PrimHead(V3207)

tmp6371 := PrimTail(tmp6370)

tmp6372 := PrimHead(tmp6371)

tmp6373 := PrimTail(V3207)

__e.Return(PrimCons(tmp6372, tmp6373))
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4unpack_1foreign)
return
}


}, 1)

tmp6396 := Call(__e, ns2_1set, symshen_4unpack_1foreign, tmp6369)


_ = tmp6396

tmp6397 := MakeNative(func(__e *ControlFlow) {
V3210 := __e.Get(1)
_ = V3210
tmp6419 := PrimIsPair(V3210)

var ifres6399 Obj

if True == tmp6419 {
tmp6417 := PrimHead(V3210)

tmp6418 := PrimIsPair(tmp6417)

var ifres6401 Obj

if True == tmp6418 {
tmp6414 := PrimHead(V3210)

tmp6415 := PrimHead(tmp6414)

tmp6416 := PrimEqual(symforeign, tmp6415)

var ifres6403 Obj

if True == tmp6416 {
tmp6411 := PrimHead(V3210)

tmp6412 := PrimTail(tmp6411)

tmp6413 := PrimIsPair(tmp6412)

var ifres6405 Obj

if True == tmp6413 {
tmp6407 := PrimHead(V3210)

tmp6408 := PrimTail(tmp6407)

tmp6409 := PrimTail(tmp6408)

tmp6410 := PrimEqual(Nil, tmp6409)

var ifres6406 Obj

if True == tmp6410 {
ifres6406 = True


} else {
ifres6406 = False


}

ifres6405 = ifres6406


} else {
ifres6405 = False


}

var ifres6404 Obj

if True == ifres6405 {
ifres6404 = True


} else {
ifres6404 = False


}

ifres6403 = ifres6404


} else {
ifres6403 = False


}

var ifres6402 Obj

if True == ifres6403 {
ifres6402 = True


} else {
ifres6402 = False


}

ifres6401 = ifres6402


} else {
ifres6401 = False


}

var ifres6400 Obj

if True == ifres6401 {
ifres6400 = True


} else {
ifres6400 = False


}

ifres6399 = ifres6400


} else {
ifres6399 = False


}

if True == ifres6399 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp6420 := Call(__e, ns2_1set, symshen_4foreign_2, tmp6397)


_ = tmp6420

tmp6421 := MakeNative(func(__e *ControlFlow) {
V3213 := __e.Get(1)
_ = V3213
tmp6427 := PrimIsPair(V3213)

var ifres6423 Obj

if True == tmp6427 {
tmp6425 := PrimTail(V3213)

tmp6426 := PrimEqual(Nil, tmp6425)

var ifres6424 Obj

if True == tmp6426 {
ifres6424 = True


} else {
ifres6424 = False


}

ifres6423 = ifres6424


} else {
ifres6423 = False


}

if True == ifres6423 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp6428 := Call(__e, ns2_1set, symshen_4zero_1place_2, tmp6421)


_ = tmp6428

tmp6429 := MakeNative(func(__e *ControlFlow) {
V3214 := __e.Get(1)
_ = V3214
tmp6434 := PrimIsSymbol(V3214)

if True == tmp6434 {
tmp6431 := PrimStr(V3214)

tmp6432 := Call(__e, PrimFunc(symshen_4internal_1to_1shen_2), tmp6431)


if True == tmp6432 {
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

tmp6435 := Call(__e, ns2_1set, symshen_4shen_1call_2, tmp6429)


_ = tmp6435

tmp6436 := MakeNative(func(__e *ControlFlow) {
V3219 := __e.Get(1)
_ = V3219
tmp6466 := PrimIsPair(V3219)

var ifres6453 Obj

if True == tmp6466 {
tmp6464 := PrimHead(V3219)

tmp6465 := PrimEqual(symprotect, tmp6464)

var ifres6455 Obj

if True == tmp6465 {
tmp6462 := PrimTail(V3219)

tmp6463 := PrimIsPair(tmp6462)

var ifres6457 Obj

if True == tmp6463 {
tmp6459 := PrimTail(V3219)

tmp6460 := PrimTail(tmp6459)

tmp6461 := PrimEqual(Nil, tmp6460)

var ifres6458 Obj

if True == tmp6461 {
ifres6458 = True


} else {
ifres6458 = False


}

ifres6457 = ifres6458


} else {
ifres6457 = False


}

var ifres6456 Obj

if True == ifres6457 {
ifres6456 = True


} else {
ifres6456 = False


}

ifres6455 = ifres6456


} else {
ifres6455 = False


}

var ifres6454 Obj

if True == ifres6455 {
ifres6454 = True


} else {
ifres6454 = False


}

ifres6453 = ifres6454


} else {
ifres6453 = False


}

if True == ifres6453 {
__e.Return(False)
return
} else {
tmp6451 := PrimIsPair(V3219)

var ifres6438 Obj

if True == tmp6451 {
tmp6449 := PrimHead(V3219)

tmp6450 := PrimEqual(symforeign, tmp6449)

var ifres6440 Obj

if True == tmp6450 {
tmp6447 := PrimTail(V3219)

tmp6448 := PrimIsPair(tmp6447)

var ifres6442 Obj

if True == tmp6448 {
tmp6444 := PrimTail(V3219)

tmp6445 := PrimTail(tmp6444)

tmp6446 := PrimEqual(Nil, tmp6445)

var ifres6443 Obj

if True == tmp6446 {
ifres6443 = True


} else {
ifres6443 = False


}

ifres6442 = ifres6443


} else {
ifres6442 = False


}

var ifres6441 Obj

if True == ifres6442 {
ifres6441 = True


} else {
ifres6441 = False


}

ifres6440 = ifres6441


} else {
ifres6440 = False


}

var ifres6439 Obj

if True == ifres6440 {
ifres6439 = True


} else {
ifres6439 = False


}

ifres6438 = ifres6439


} else {
ifres6438 = False


}

if True == ifres6438 {
__e.Return(False)
return
} else {
__e.Return(PrimIsPair(V3219))
return
}


}


}, 1)

tmp6467 := Call(__e, ns2_1set, symshen_4application_2, tmp6436)


_ = tmp6467

tmp6468 := MakeNative(func(__e *ControlFlow) {
V3224 := __e.Get(1)
_ = V3224
V3225 := __e.Get(2)
_ = V3225
tmp6476 := PrimEqual(MakeNumber(-1), V3225)

if True == tmp6476 {
tmp6474 := Call(__e, PrimFunc(symshen_4lowercase_1symbol_2), V3224)


if True == tmp6474 {
tmp6470 := Call(__e, PrimFunc(symexternal), symshen)


tmp6471 := Call(__e, PrimFunc(symelement_2), V3224, tmp6470)


tmp6472 := PrimNot(tmp6471)

if True == tmp6472 {
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

tmp6477 := Call(__e, ns2_1set, symshen_4undefined_1f_2, tmp6468)


_ = tmp6477

tmp6478 := MakeNative(func(__e *ControlFlow) {
V3226 := __e.Get(1)
_ = V3226
tmp6483 := PrimIsSymbol(V3226)

if True == tmp6483 {
tmp6480 := PrimIsVariable(V3226)

tmp6481 := PrimNot(tmp6480)

if True == tmp6481 {
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

tmp6484 := Call(__e, ns2_1set, symshen_4lowercase_1symbol_2, tmp6478)


_ = tmp6484

tmp6485 := MakeNative(func(__e *ControlFlow) {
V3227 := __e.Get(1)
_ = V3227
tmp6515 := PrimIsPair(V3227)

var ifres6506 Obj

if True == tmp6515 {
tmp6513 := PrimTail(V3227)

tmp6514 := PrimIsPair(tmp6513)

var ifres6508 Obj

if True == tmp6514 {
tmp6510 := PrimTail(V3227)

tmp6511 := PrimTail(tmp6510)

tmp6512 := PrimEqual(Nil, tmp6511)

var ifres6509 Obj

if True == tmp6512 {
ifres6509 = True


} else {
ifres6509 = False


}

ifres6508 = ifres6509


} else {
ifres6508 = False


}

var ifres6507 Obj

if True == ifres6508 {
ifres6507 = True


} else {
ifres6507 = False


}

ifres6506 = ifres6507


} else {
ifres6506 = False


}

if True == ifres6506 {
__e.Return(V3227)
return
} else {
tmp6504 := PrimIsPair(V3227)

var ifres6495 Obj

if True == tmp6504 {
tmp6502 := PrimTail(V3227)

tmp6503 := PrimIsPair(tmp6502)

var ifres6497 Obj

if True == tmp6503 {
tmp6499 := PrimTail(V3227)

tmp6500 := PrimTail(tmp6499)

tmp6501 := PrimIsPair(tmp6500)

var ifres6498 Obj

if True == tmp6501 {
ifres6498 = True


} else {
ifres6498 = False


}

ifres6497 = ifres6498


} else {
ifres6497 = False


}

var ifres6496 Obj

if True == ifres6497 {
ifres6496 = True


} else {
ifres6496 = False


}

ifres6495 = ifres6496


} else {
ifres6495 = False


}

if True == ifres6495 {
tmp6486 := PrimHead(V3227)

tmp6487 := PrimTail(V3227)

tmp6488 := PrimHead(tmp6487)

tmp6489 := PrimCons(tmp6488, Nil)

tmp6490 := PrimCons(tmp6486, tmp6489)

tmp6491 := PrimTail(V3227)

tmp6492 := PrimTail(tmp6491)

tmp6493 := PrimCons(tmp6490, tmp6492)

__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp6493)
return


} else {
__e.Return(V3227)
return
}


}


}, 1)

tmp6516 := Call(__e, ns2_1set, symshen_4simple_1curry, tmp6485)


_ = tmp6516

tmp6517 := MakeNative(func(__e *ControlFlow) {
V3228 := __e.Get(1)
_ = V3228
__e.TailApply(PrimFunc(symfn), V3228)
return
}, 1)

tmp6518 := Call(__e, ns2_1set, symfunction, tmp6517)


_ = tmp6518

tmp6519 := MakeNative(func(__e *ControlFlow) {
V3229 := __e.Get(1)
_ = V3229
tmp6526 := Call(__e, PrimFunc(symarity), V3229)


tmp6527 := PrimEqual(tmp6526, MakeNumber(0))

if True == tmp6527 {
__e.TailApply(V3229)
return
} else {
tmp6520 := MakeNative(func(__e *ControlFlow) {
tmp6521 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V3229, symshen_4lambda_1form, tmp6521)
return


}, 0)

tmp6522 := MakeNative(func(__e *ControlFlow) {
Z3230 := __e.Get(1)
_ = Z3230
tmp6523 := Call(__e, PrimFunc(symshen_4app), V3229, MakeString(" is undefined\n"), symshen_4a)


tmp6524 := PrimStringConcat(MakeString("fn: "), tmp6523)

__e.Return(PrimSimpleError(tmp6524))
return


}, 1)

__e.TailApply(try_1catch, tmp6520, tmp6522)
return


}


}, 1)

tmp6528 := Call(__e, ns2_1set, symfn, tmp6519)


_ = tmp6528

tmp6529 := MakeNative(func(__e *ControlFlow) {
V3233 := __e.Get(1)
_ = V3233
tmp6559 := PrimIsPair(V3233)

var ifres6546 Obj

if True == tmp6559 {
tmp6557 := PrimHead(V3233)

tmp6558 := PrimEqual(symfn, tmp6557)

var ifres6548 Obj

if True == tmp6558 {
tmp6555 := PrimTail(V3233)

tmp6556 := PrimIsPair(tmp6555)

var ifres6550 Obj

if True == tmp6556 {
tmp6552 := PrimTail(V3233)

tmp6553 := PrimTail(tmp6552)

tmp6554 := PrimEqual(Nil, tmp6553)

var ifres6551 Obj

if True == tmp6554 {
ifres6551 = True


} else {
ifres6551 = False


}

ifres6550 = ifres6551


} else {
ifres6550 = False


}

var ifres6549 Obj

if True == ifres6550 {
ifres6549 = True


} else {
ifres6549 = False


}

ifres6548 = ifres6549


} else {
ifres6548 = False


}

var ifres6547 Obj

if True == ifres6548 {
ifres6547 = True


} else {
ifres6547 = False


}

ifres6546 = ifres6547


} else {
ifres6546 = False


}

if True == ifres6546 {
__e.Return(True)
return
} else {
tmp6544 := PrimIsPair(V3233)

var ifres6531 Obj

if True == tmp6544 {
tmp6542 := PrimHead(V3233)

tmp6543 := PrimEqual(symfunction, tmp6542)

var ifres6533 Obj

if True == tmp6543 {
tmp6540 := PrimTail(V3233)

tmp6541 := PrimIsPair(tmp6540)

var ifres6535 Obj

if True == tmp6541 {
tmp6537 := PrimTail(V3233)

tmp6538 := PrimTail(tmp6537)

tmp6539 := PrimEqual(Nil, tmp6538)

var ifres6536 Obj

if True == tmp6539 {
ifres6536 = True


} else {
ifres6536 = False


}

ifres6535 = ifres6536


} else {
ifres6535 = False


}

var ifres6534 Obj

if True == ifres6535 {
ifres6534 = True


} else {
ifres6534 = False


}

ifres6533 = ifres6534


} else {
ifres6533 = False


}

var ifres6532 Obj

if True == ifres6533 {
ifres6532 = True


} else {
ifres6532 = False


}

ifres6531 = ifres6532


} else {
ifres6531 = False


}

if True == ifres6531 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp6560 := Call(__e, ns2_1set, symshen_4fn_1call_2, tmp6529)


_ = tmp6560

tmp6561 := MakeNative(func(__e *ControlFlow) {
V3234 := __e.Get(1)
_ = V3234
tmp6602 := PrimIsPair(V3234)

var ifres6589 Obj

if True == tmp6602 {
tmp6600 := PrimHead(V3234)

tmp6601 := PrimEqual(symfunction, tmp6600)

var ifres6591 Obj

if True == tmp6601 {
tmp6598 := PrimTail(V3234)

tmp6599 := PrimIsPair(tmp6598)

var ifres6593 Obj

if True == tmp6599 {
tmp6595 := PrimTail(V3234)

tmp6596 := PrimTail(tmp6595)

tmp6597 := PrimEqual(Nil, tmp6596)

var ifres6594 Obj

if True == tmp6597 {
ifres6594 = True


} else {
ifres6594 = False


}

ifres6593 = ifres6594


} else {
ifres6593 = False


}

var ifres6592 Obj

if True == ifres6593 {
ifres6592 = True


} else {
ifres6592 = False


}

ifres6591 = ifres6592


} else {
ifres6591 = False


}

var ifres6590 Obj

if True == ifres6591 {
ifres6590 = True


} else {
ifres6590 = False


}

ifres6589 = ifres6590


} else {
ifres6589 = False


}

if True == ifres6589 {
tmp6562 := PrimTail(V3234)

tmp6563 := PrimCons(symfn, tmp6562)

__e.TailApply(PrimFunc(symshen_4fn_1call), tmp6563)
return


} else {
tmp6587 := PrimIsPair(V3234)

var ifres6574 Obj

if True == tmp6587 {
tmp6585 := PrimHead(V3234)

tmp6586 := PrimEqual(symfn, tmp6585)

var ifres6576 Obj

if True == tmp6586 {
tmp6583 := PrimTail(V3234)

tmp6584 := PrimIsPair(tmp6583)

var ifres6578 Obj

if True == tmp6584 {
tmp6580 := PrimTail(V3234)

tmp6581 := PrimTail(tmp6580)

tmp6582 := PrimEqual(Nil, tmp6581)

var ifres6579 Obj

if True == tmp6582 {
ifres6579 = True


} else {
ifres6579 = False


}

ifres6578 = ifres6579


} else {
ifres6578 = False


}

var ifres6577 Obj

if True == ifres6578 {
ifres6577 = True


} else {
ifres6577 = False


}

ifres6576 = ifres6577


} else {
ifres6576 = False


}

var ifres6575 Obj

if True == ifres6576 {
ifres6575 = True


} else {
ifres6575 = False


}

ifres6574 = ifres6575


} else {
ifres6574 = False


}

if True == ifres6574 {
tmp6564 := MakeNative(func(__e *ControlFlow) {
W3235 := __e.Get(1)
_ = W3235
tmp6569 := PrimEqual(W3235, MakeNumber(-1))

if True == tmp6569 {
__e.Return(V3234)
return
} else {
tmp6567 := PrimEqual(W3235, MakeNumber(0))

if True == tmp6567 {
__e.Return(PrimTail(V3234))
return
} else {
tmp6565 := PrimTail(V3234)

__e.TailApply(PrimFunc(symshen_4lambda_1function), tmp6565, W3235)
return


}


}


}, 1)

tmp6570 := PrimTail(V3234)

tmp6571 := PrimHead(tmp6570)

tmp6572 := Call(__e, PrimFunc(symarity), tmp6571)


__e.TailApply(tmp6564, tmp6572)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4fn_1call)
return
}


}


}, 1)

tmp6603 := Call(__e, ns2_1set, symshen_4fn_1call, tmp6561)


_ = tmp6603

tmp6604 := MakeNative(func(__e *ControlFlow) {
V3236 := __e.Get(1)
_ = V3236
V3237 := __e.Get(2)
_ = V3237
V3238 := __e.Get(3)
_ = V3238
tmp6605 := MakeNative(func(__e *ControlFlow) {
W3239 := __e.Get(1)
_ = W3239
tmp6606 := MakeNative(func(__e *ControlFlow) {
W3240 := __e.Get(1)
_ = W3240
__e.Return(W3239)
return
}, 1)

var ifres6612 Obj

if True == W3239 {
tmp6620 := Call(__e, PrimFunc(symshen_4loading_2))


var ifres6614 Obj

if True == tmp6620 {
tmp6616 := PrimCons(sym_1, Nil)

tmp6617 := PrimCons(sym_7, tmp6616)

tmp6618 := Call(__e, PrimFunc(symelement_2), V3236, tmp6617)


tmp6619 := PrimNot(tmp6618)

var ifres6615 Obj

if True == tmp6619 {
ifres6615 = True


} else {
ifres6615 = False


}

ifres6614 = ifres6615


} else {
ifres6614 = False


}

var ifres6613 Obj

if True == ifres6614 {
ifres6613 = True


} else {
ifres6613 = False


}

ifres6612 = ifres6613


} else {
ifres6612 = False


}

var ifres6607 Obj

if True == ifres6612 {
tmp6608 := Call(__e, PrimFunc(symshen_4app), V3236, MakeString("\n"), symshen_4a)


tmp6609 := PrimStringConcat(MakeString("partial application of "), tmp6608)

tmp6610 := Call(__e, PrimFunc(symstoutput))


tmp6611 := Call(__e, PrimFunc(sympr), tmp6609, tmp6610)


ifres6607 = tmp6611


} else {
ifres6607 = symshen_4skip


}

__e.TailApply(tmp6606, ifres6607)
return


}, 1)

tmp6621 := PrimGreatThan(V3237, V3238)

__e.TailApply(tmp6605, tmp6621)
return


}, 3)

tmp6622 := Call(__e, ns2_1set, symshen_4partial_1application_d_2, tmp6604)


_ = tmp6622

tmp6623 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dloading_2_d))
return
}, 0)

tmp6624 := Call(__e, ns2_1set, symshen_4loading_2, tmp6623)


_ = tmp6624

tmp6625 := MakeNative(func(__e *ControlFlow) {
V3245 := __e.Get(1)
_ = V3245
V3246 := __e.Get(2)
_ = V3246
V3247 := __e.Get(3)
_ = V3247
tmp6643 := PrimEqual(MakeNumber(-1), V3246)

if True == tmp6643 {
__e.Return(False)
return
} else {
tmp6626 := MakeNative(func(__e *ControlFlow) {
W3248 := __e.Get(1)
_ = W3248
tmp6627 := MakeNative(func(__e *ControlFlow) {
W3249 := __e.Get(1)
_ = W3249
__e.Return(W3248)
return
}, 1)

var ifres6638 Obj

if True == W3248 {
tmp6640 := Call(__e, PrimFunc(symshen_4loading_2))


var ifres6639 Obj

if True == tmp6640 {
ifres6639 = True


} else {
ifres6639 = False


}

ifres6638 = ifres6639


} else {
ifres6638 = False


}

var ifres6628 Obj

if True == ifres6638 {
tmp6630 := PrimEqual(V3247, MakeNumber(1))

var ifres6629 Obj

if True == tmp6630 {
ifres6629 = MakeString("")


} else {
ifres6629 = MakeString("s")


}

tmp6631 := Call(__e, PrimFunc(symshen_4app), ifres6629, MakeString("\n"), symshen_4a)


tmp6632 := PrimStringConcat(MakeString(" argument"), tmp6631)

tmp6633 := Call(__e, PrimFunc(symshen_4app), V3247, tmp6632, symshen_4a)


tmp6634 := PrimStringConcat(MakeString(" might not like "), tmp6633)

tmp6635 := Call(__e, PrimFunc(symshen_4app), V3245, tmp6634, symshen_4a)


tmp6636 := Call(__e, PrimFunc(symstoutput))


tmp6637 := Call(__e, PrimFunc(sympr), tmp6635, tmp6636)


ifres6628 = tmp6637


} else {
ifres6628 = symshen_4skip


}

__e.TailApply(tmp6627, ifres6628)
return


}, 1)

tmp6641 := PrimLessThan(V3246, V3247)

__e.TailApply(tmp6626, tmp6641)
return


}


}, 3)

__e.TailApply(ns2_1set, symshen_4overapplication_2, tmp6625)
return




}, 0)

